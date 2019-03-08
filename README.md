# Emailer [W I P]

> A Go client to send emails from Command Line.

### Development
* Create a `.env` file in the root and fill with this
```
API_KEY='<api key>'
SENDER_MAIL='<Sender's email id>'
SENDER_NAME='<Your Name>'
```
* See all available command `email --help`
```
Send Emails from your command line

Usage:
  email [command]

Available Commands:
  help        Help about any command
  send        send the email

Flags:
  -h, --help   help for email

Use "email [command] --help" for more information about a command.

```
* Run `email send` to enter all information and press `Enter`

![live](https://user-images.githubusercontent.com/33368759/54043876-a570c400-41f3-11e9-9a7a-ec3269c7e5d0.PNG)

## TODO
* [ ] Avoid sending mails to spam folder
* [ ] Add Test cases
* [ ] Feature to attach files to email
* [ ] ....

> Found it intresting then do :star: it, also suggest feature ( by creating issues ). :tada:
