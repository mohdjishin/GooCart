package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	_ "github.com/joho/godotenv/autoload"
	"github.com/twilio/twilio-go"

	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var TWILIO_ACCOUNT_SID string = os.Getenv("TWILIO_ACCOUNT_SID")
var TWILIO_AUTH_TOKEN string = os.Getenv("TWILIO_AUTH_TOKEN")
var VERIFY_SERVICE_SID string = os.Getenv("VERIFY_SERVICE_SID")
var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: TWILIO_ACCOUNT_SID,
	Password: TWILIO_AUTH_TOKEN,
})

func SendOtp(to string) bool {
	params := &openapi.CreateVerificationParams{}

	params.SetTo(to)
	params.SetChannel("sms")
	fmt.Println(params.CustomCode)

	_, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		return true
	}
	return false
}

func CheckOtp(to, code string) bool {

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)

	if err != nil {
		fmt.Println(err.Error())
	} else if *resp.Status == "approved" {
		return true
	} else {
		return false
	}
	return false
}

// Replace AccessKeyID with your AccessKeyID key.
var AccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")

// Replace AccessKeyID with your AccessKeyID key.
var SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")

// Replace us-west-2 with the AWS Region you're using for Amazon SNS.
var AwsRegion = os.Getenv("AWS_REGION")

// OTP := strconv.FormatUint(uint64(uuid.New().ID()), 10)[:6]

func SendSMSOTP(phoneNumber string, otp string) error {
	fmt.Println(phoneNumber)
	message := "GC-" + otp + " is your GoCart verification code"
	// Create Session and assign AccessKeyID and SecretAccessKey
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AwsRegion),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
	},
	)

	// Create SNS service
	svc := sns.New(sess)

	// Pass the phone number and message.
	params := &sns.PublishInput{
		PhoneNumber: aws.String(phoneNumber),
		Message:     aws.String(message),
	}

	// sends a text message (SMS message) directly to a phone number.
	resp, err := svc.Publish(params)

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(resp) // print the response data.

	return nil

}

func WelcomeMsg(phoneNumber string) error {

	message := "Welcome to GoCart! Your OTP verification was successful.. Happy shopping!"
	// Create Session and assign AccessKeyID and SecretAccessKey
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AwsRegion),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
	},
	)

	// Create SNS service
	svc := sns.New(sess)

	// Pass the phone number and message.
	params := &sns.PublishInput{
		PhoneNumber: aws.String(phoneNumber),
		Message:     aws.String(message),
	}

	// sends a text message (SMS message) directly to a phone number.
	resp, err := svc.Publish(params)

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(resp) // print the response data.

	return nil
}
