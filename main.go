package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
)

func main() {
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
	} else {
		// do nothing
		fmt.Println("no change")
		sendEmail()
	}

	// fmt.Println(string(body))
}

func sendEmail() {
	email := "some_email"
	pass := "some_password"
	host := "smtp.mailtrap.io"

	auth := smtp.PlainAuth("", email, pass, host)

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{email}
	msg := []byte("poocoin listed new feature coins")
	err := smtp.SendMail("smtp.mailtrap.io:25", auth, "piotr@mailtrap.io", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
