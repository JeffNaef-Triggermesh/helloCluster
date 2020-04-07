//gcloud builds submit --tag gcr.io/loadtester-271418/gotst

package main

import (
	"context"
	"fmt"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func main() {
	ctx := context.Background()
	p, err := cloudevents.NewHTTP()
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}

	c, err := cloudevents.NewClient(p)
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Printf("will listen on :8080\n")
	log.Fatalf("failed to start receiver: %s", c.StartReceiver(ctx, receive))
}

func receive(ctx context.Context, event cloudevents.Event) {
	fmt.Printf("%s", event)
}

// func handler(w http.ResponseWriter, r *http.Request) {

// 	recipientEmail := r.Header.Get("RecipientEmail")
// 	if recipientEmail == "" {
// 		fmt.Println("no recipient email found in header file")
// 		fmt.Fprintf(w, "no recipient email found in header file \n %s!", r.URL.Path[1:])
// 	}

// 	recipientName = r.Header.Get("RecipientName")
// 	if recipientName == "" {
// 		fmt.Println("no recipient name found in header")
// 		fmt.Fprintf(w, "no recipient name found in email \n %s!", r.URL.Path[1:])
// 		voidCall = 1
// 	}

// 	msg := r.Header.Get("Message")
// 	if msg == "" {
// 		fmt.Println("no message in header")
// 		fmt.Fprintf(w, "no message found in header \n %s!", r.URL.Path[1:])
// 		voidCall = 1
// 	}

// 	senderEmail := r.Header.Get("SenderEmail")
// 	if senderEmail == "" {
// 		fmt.Println("no sender email found in header")
// 		fmt.Fprintf(w, "no sender email in header \n  %s!", r.URL.Path[1:])
// 		voidCall = 1
// 	}

// 	senderName = r.Header.Get("SenderName")
// 	if senderName == "" {
// 		fmt.Println("no sender name found in header")
// 		fmt.Fprintf(w, "no sender name found in email \n %s!", r.URL.Path[1:])
// 		voidCall = 1
// 	}

// 	//fmt.Println(senderEmail)
// 	if voidCall == 0 {
// 		sendGridStatusCodeResponse = sendEmail(msg, senderEmail, senderName, recipientEmail, recipientName)
// 		fmt.Println("sending email")
// 		fmt.Fprintf(w, "SendGrid API responded with status code : %d", sendGridStatusCodeResponse)

// 	}
// 	//fmt.Fprintf(w, "Helloss %s!", r.URL.Path[1:])
// 	fmt.Println("RESTfulServ. on:8093, Controller:", r.URL.Path[1:])
// 	voidCall = 0
// }


// func errorHandler(err error) {
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }

// func sendEmail(msg string, fromEmail string, fromName string, toEmail string, toName string) int {
// 	from := mail.NewEmail(fromName, fromEmail)
// 	subject := "Notification from GO Rest API"
// 	to := mail.NewEmail(toName, toEmail)
// 	plainTextContent := msg
// 	htmlContent := "<strong>Sent with GO rest API</strong>"
// 	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
// 	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
// 	response, err := client.Send(message)
// 	if err != nil {
// 		log.Println(err)

// 	} else {
// 		fmt.Println("Recieved response from sendGrid API")
// 		fmt.Println("------response.StatusCode------")
// 		fmt.Println(response.StatusCode)
// 		fmt.Println("------response.Body------")
// 		fmt.Println(response.Body)
// 		fmt.Println("------response.Headers------")
// 		fmt.Println(response.Headers)
// 		var resp int = response.StatusCode
// 		return resp
// 	}
// 	return 0
// }

// func main() {
// 	from := mail.NewEmail("Sendgrid Test", "jeff@triggermesh.com")
// 	subject := "Sending with SendGrid is Fun"
// 	to := mail.NewEmail("Jeff", "jeffthenaef@gmail.com")
// 	plainTextContent := "and easy to do anywhere, even with Go"
// 	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
// 	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
// 	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
// 	response, err := client.Send(message)
// 	if err != nil {
// 		log.Println(err)
// 	} else {
// 		fmt.Println(response.StatusCode)
// 		fmt.Println(response.Body)
// 		fmt.Println(response.Headers)
// 	}
// }
