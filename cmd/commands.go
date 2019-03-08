package cmd

import (
	"log"
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/joho/godotenv"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send the email",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		send();	
	},
}

func init() {
	RootCmd.AddCommand(sendCmd)
}

func send() {
	from := mail.NewEmail("Example User", "b117019@iiit-bh.ac.in")
	subject := "Hello"
	to := mail.NewEmail("Example User", "b117019@iiit-bh.ac.in")
	plainTextContent := "And how are you"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
