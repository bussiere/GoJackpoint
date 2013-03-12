package main

import (
	"log"
	"net/smtp"
)

func sendmail(user string, pass string, server string,port string, message string,recept string[]){
                auth := smtp.PlainAuth(
                "",
                user,
                pass,
                server,
        )
        // Connect to the server, authenticate, set the sender and recipient,
        // and send the email all in one step.
        err := smtp.SendMail(
                "mail.example.com:25",
                auth,
                user,
                recept,
                []byte(message),
        )
        if err != nil {
                log.Fatal(err)
        }
}
}

func main() {
	// Set up authentication information.
}

