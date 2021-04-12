package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"time"
)

func main() {

	for {
		fmt.Println("fetching updates")

		r, err := http.Get("https://poocoin.app/whitelist1-tokens.json")
		if err != nil {
			// handle error
		}
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		response := string(body)

		lastResponse, err := ioutil.ReadFile("./last_response.txt")
		if err != nil {
			log.Fatal(err)
		}

		if response != string(lastResponse) {
			ioutil.WriteFile("./last_response.txt", []byte(response), 0644)
			//email
			fmt.Println("change")
			sendEmail()
		} else {
			// do nothing
			fmt.Println("no change")
		}

		time.Sleep(time.Minute)
	}
}

func sendEmail() {
	email := "talos.tester.1@gmail.com"
	pass := "Tester123!"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	message := []byte("Poocoin updated their feature list!")

	// add recipient emails here
	to := []string{
		"talos.tester.1@gmail.com",
		"7rossilli7@gmail.com",
		"rrossilli55@gmail.com",
	}

	auth := smtp.PlainAuth("", email, pass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
