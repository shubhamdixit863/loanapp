package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func init() {

	godotenv.Load()

}

func SendSMS(n string, m string) (bool, error) {

	accountSid := os.Getenv("accountsid")
	authToken := os.Getenv("authtoken")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(fmt.Sprintf("%q", n))
	params.SetMessagingServiceSid(os.Getenv("messageserviceid"))
	params.SetFrom(os.Getenv("fromNumber"))
	params.SetBody(m)

	message, err := client.Api.CreateMessage(params)

	if err != nil {
		fmt.Println(err)
		return false, err
	} else {
		fmt.Println(message)
		return true, err
	}
}
