package main

import "fmt"

func sendEmail(emailChan chan string) {
	for email := range emailChan {
		fmt.Println("send email to email", email)
	}
}

func main() {

	emailChan := make(chan string, 100)

	done := make(chan bool)

	for i := 0; i < 30; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com")

	}

	<-done

	// emailChan <- "1@sanim.gmail"

	// emailChan <- "2@sanim.gmail"

	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)
}
