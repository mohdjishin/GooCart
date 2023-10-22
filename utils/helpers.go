package utils

import (
	"context"
	"fmt"
	"log"
	"os"
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

type PersonalInformation struct {
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

type ProductInfo struct {
	ID              int
	ProductCategory string
	ProductName     string
	Price           float64
}

type ProductImageInfo struct {
	ProductId  int
	ImageOne   string
	ImageTwo   string
	ImageThree string
}

type CombinedProductInfo struct {
	ID              int     `json:"id"`
	ProductCategory string  `json:"pro_category"`
	ProductName     string  `json:"pro_name"`
	Price           float64 `json:"price"`
	ImageOne        string  `json:"img_one"`
	ImageTwo        string  `json:"img_two"`
	ImageThree      string  `json:"img_three"`
}

func ExtractPersonalInfo(users []model.Users) []PersonalInformation {
	var newUsers []PersonalInformation
	for _, user := range users {
		newUsers = append(newUsers, PersonalInformation{
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

func Combined(users []PersonalInformation, addresses []Extractaddress) []Combine {
	var combined []Combine
	for _, u := range users {
		for _, a := range addresses {

			if u.UserId == int(a.UserID) {

				combined = append(combined, Combine{UserId: u.UserId, Username: u.Username, Name: u.Name, Mail: u.Mail, Phone: u.Phone, Verification: u.Verification, Housename: a.Housename, Street: a.Street, Pin: a.Pin, City: a.City, State: a.State})
			}
		}
	}

	return combined
}

func CheckComplexityOFPassword[T string](password T) bool {
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
	hasLower := regexp.MustCompile(`[a-z]`).MatchString
	hasSymbol := regexp.MustCompile(`[!@#\$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString
	return hasNumber(string(password)) && hasUpper(string(password)) && hasLower(string(password)) && hasSymbol(string(password))
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
	fmt.Println(result.Location)

	return result.Location, true, file.Filename

}

func UploadPDFToS3[T string](filePath, filname string) (bool, string) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return false, ""
	}

	client := s3.NewFromConfig(cfg)
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("error opening file")
		return false, ""

	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("error closing file")
		}

	}()

	uploader := manager.NewUploader(client)
	result, UploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("code-with-jishin"),
		Key:    aws.String(filname),
		Body:   file,
		ACL:    "public-read",
	})
	if UploadErr != nil {
		return false, ""

	}
	fmt.Println(result.Location)

	return true, result.Location
}

func ExtractProductInfo[T []model.Products](product T) []ProductInfo {
	var newpro []ProductInfo
	for _, pro := range product {
		newpro = append(newpro, ProductInfo{
			ID:              int(pro.ID),
			ProductName:     pro.Product_Name,
			ProductCategory: pro.Product_Category,
			Price:           pro.Price,
		})
	}
	return newpro
}
func ExtractProImage[T []model.ProductImage, R []ProductImageInfo](proImg T) R {
	var productImages []ProductImageInfo
	for _, pro := range proImg {
		productImages = append(productImages, ProductImageInfo{
			ProductId:  int(pro.ProductId),
			ImageOne:   pro.ImageOne,
			ImageTwo:   pro.ImgTwo,
			ImageThree: pro.ImgThree,
		})
	}
	return productImages
}
func ExtractOrderInfo[T []model.Order, R []model.OrderRespAdmin](order T) R {
	var orderInfo []model.OrderRespAdmin
	for _, or := range order {
		orderInfo = append(orderInfo, model.OrderRespAdmin{

			OrderID:   or.ID,
			ProductID: or.ProductID,
			UserID:    or.UserID,
			Quantity:  or.Quantity,
			Price:     int(or.Total),
			Status:    or.ShippmentStatus,
		})
	}
	return orderInfo
}

func CombinePRoductAndProductImage(proInfo []ProductInfo, proimageInfo []ProductImageInfo) []CombinedProductInfo {
	var combined []CombinedProductInfo
	for _, proInf := range proInfo {
		for _, proImg := range proimageInfo {

			if proInf.ID == int(proImg.ProductId) {
				fmt.Println("he;p")
				combined = append(combined, CombinedProductInfo{ID: proImg.ProductId, ProductName: proInf.ProductName, ProductCategory: proInf.ProductCategory, Price: proInf.Price, ImageOne: proImg.ImageOne, ImageTwo: proImg.ImageTwo, ImageThree: proImg.ImageThree})
			}
		}
	}

	return combined
}

// func CombinedPRoductOrder(order model.Order, products model.Products) []model.Invoice {
// 	var combined []model.Invoice
// 	for _, u := range order {
// 		for _, a := range products {

// 			if u.ProductID == a.ID {

// 				combined = append(combined, model.Invoice{OrderId: u.OrderId, ProductName: a.Product_Name, Quantity: u.Quantity, Price: a.Price})
// 			}
// 		}
// 	}
// 	fmt.Println(combined)
// 	return combined
// }
