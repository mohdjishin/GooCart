package utils

import (
	"context"
	"fmt"
	"log"
	"path"

	"mime/multipart"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mohdjishin/GoCart/model"
)

type information struct {
	UserId       int
	Name         string
	Username     string
	Phone        string
	Mail         string
	Verification bool
}

type Combine struct {
	UserId       int
	Name         string
	Username     string
	Phone        string
	Mail         string
	Verification bool
	Housename    string
	Street       string
	City         string
	Pin          string
	State        string
}

type Extractaddress struct {
	UserID    uint
	Housename string
	Street    string
	City      string
	Pin       string
	State     string
}

func ExtractPersonalInfo(users []model.Users) []information {
	var newUsers []information
	for _, user := range users {
		newUsers = append(newUsers, information{
			UserId:       int(user.ID),
			Name:         user.Name,
			Username:     user.Username,
			Phone:        user.Phone,
			Mail:         user.Email,
			Verification: user.Verified,
		})
	}
	return newUsers
}
func ExtractAdresses(addr []model.Address) []Extractaddress {
	var newUsers []Extractaddress
	for _, addr := range addr {
		newUsers = append(newUsers, Extractaddress{
			UserID:    addr.UserId,
			Housename: addr.HouseName,
			Street:    addr.Street,
			City:      addr.City,
			Pin:       addr.Pin,
			State:     addr.State,
		})
	}
	return newUsers
}

func Combined(users []information, addresses []Extractaddress) []Combine {
	var combined []Combine
	for _, u := range users {
		for _, a := range addresses {

			if u.UserId == int(a.UserID) {
				fmt.Println("he;p")
				combined = append(combined, Combine{UserId: u.UserId, Username: u.Username, Name: u.Name, Mail: u.Mail, Phone: u.Phone, Verification: u.Verification, Housename: a.Housename, Street: a.Street, Pin: a.Pin, City: a.City, State: a.State})
			}
		}
	}
	fmt.Println(combined)
	return combined
}

func CheckComplexityOFPassword(password string) bool {
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
	hasLower := regexp.MustCompile(`[a-z]`).MatchString
	hasSymbol := regexp.MustCompile(`[!@#\$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString
	return hasNumber(password) && hasUpper(password) && hasLower(password) && hasSymbol(password)
}

func UploadToBucket(file *multipart.FileHeader) (string, bool, string) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return "loading failed", false, ""
	}

	client := s3.NewFromConfig(cfg)

	file.Filename = uuid.New().String() + path.Ext(file.Filename)

	f, err := file.Open()
	if err != nil {
		return "file open failed", false, ""
	}

	uploader := manager.NewUploader(client)
	result, UploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("code-with-jishin"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})
	if UploadErr != nil {
		return "upload error", false, ""
	}

	return result.Location, true, file.Filename

}
