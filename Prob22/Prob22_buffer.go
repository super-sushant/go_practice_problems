package main 

func addEmailToQueue(emails []string ) chan string {
	emailsToSend := make(chan string,len(emails))

	for _,email :=range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

func sendEmails(batchSize int, ch chan string) {
	for i:= 0; i< batchSize; i++ {
		email:= <-ch
		println("Sending email:",email)
	}
}

func main () {
	println("Adding email to queue")

	ch:=addEmailToQueue([]string{"hi","hello","email2","email3"})

	println("Sending emails")
	sendEmails(4,ch)
}