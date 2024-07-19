package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
)

func main() {

	accountSid := "YOUR_TWILIO_ACCOUNT_SID"
	authToken := "YOUR_TWILIO_AUTH_TOKEN"
	fromPhone := "YOUR_TWILIO_PHONE_NUMBER"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	messageBody := "سلام با آرزوی شاد کامی و برکت و سلامتی در سال پیش رو\nارادتمند شما هادی رضایی"

	fmt.Println("لطفاً شماره‌های تلفن را وارد کنید (هر شماره را در یک خط وارد کنید و برای پایان، فقط Enter را فشار دهید):")

	scanner := bufio.NewScanner(os.Stdin)
	var phoneNumbers []string

	for scanner.Scan() {
		phoneNumber := scanner.Text()
		if phoneNumber == "" {
			break
		}
		phoneNumbers = append(phoneNumbers, phoneNumber)
	}

	for _, toPhone := range phoneNumbers {
		params := &openapi.CreateMessageParams{}
		params.SetTo(toPhone)
		params.SetFrom(fromPhone)
		params.SetBody(messageBody)

		_, err := client.ApiV2010.CreateMessage(params)
		if err != nil {
			fmt.Printf("خطا در ارسال پیام به %s: %s\n", toPhone, err.Error())
		} else {
			fmt.Printf("پیام به شماره %s ارسال شد\n", toPhone)
		}
	}
}
