package utils

import (
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func SendOtp(msg string, toNumber string) (bool, error) {

	accountSid := os.Getenv("AC50327d87c2943fd0b4d3aed31505df2e")
	authToken := os.Getenv("3ef29f24126b76a90543076aaa74bbb4")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(toNumber)
	params.SetFrom(os.Getenv("+19786435823"))
	params.SetBody(msg)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return false, err
	} else {
		return true, err
	}

}
