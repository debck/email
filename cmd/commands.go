package cmd

import (
	"log"
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/joho/godotenv"
	"gopkg.in/AlecAivazis/survey.v1"
)


var qs = []*survey.Question{
    {
        Name:     "SendTo",
        Prompt:   &survey.Input{Message: "send to ?"},
        Validate: survey.Required,
        Transform: survey.Title,
	},
	{
        Name:     "Subject",
        Prompt:   &survey.Input{Message: "Email Subject?"},
        Validate: survey.Required,
        Transform: survey.Title,
	},
	{
        Name:     "Msg",
        Prompt:   &survey.Input{Message: "Type the message"},
        Validate: survey.Required,
        Transform: survey.Title,
	},
}
	

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send the email",
	Run: func(cmd *cobra.Command, args []string) {

		//load env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		answers := struct {
			SendTo          string
			Subject         string
			Msg 			string          
		}{}

		err = survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		send(answers.SendTo, answers.Subject, answers.Msg);	
	},
}


func send(sendto,sub, Msg string) {

	from := mail.NewEmail(os.Getenv("SENDER_NAME"), os.Getenv("SENDER_MAIL"))
	subject := sub
	to := mail.NewEmail("Coder", sendto)
	plainTextContent := Msg
	htmlContent := "<strong>"+Msg+"</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		if response.StatusCode == 202 {
			fmt.Println("Email Sent Successfully!!!!")
		}
		// fmt.Println(response.StatusCode)
		// fmt.Println(response.Body)
		// fmt.Println(response.Headers)
	}
}

func init() {
	RootCmd.AddCommand(sendCmd)
}
