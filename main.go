package main

import (
	"fmt"
    "net/smtp"
	"os"
	"io/ioutil"
	"strings"
)

type Receiver struct {
	Login string
	Name string
	Company string
	Location string
	Email string
}

const sender_name string = "Test user"
const sender_email string = "iterator_sl@qq.com"
const auth_email string = "iterator_sl@qq.com"
const auth_password string = "doubzolyuvyubdai"
const auth_host string = "smtp.qq.com"
const host string = "smtp.qq.com:25"
const subject string = "Test Subject"
const content string = "Hi Test Content"

func sendEmail(receiver Receiver, subject string, content string) error {
	auth := smtp.PlainAuth("", auth_email, auth_password, auth_host)
	receiver_email := []string{receiver.Email}

	text := fmt.Sprintf("From: %s <%s>\r\n" + 
		"To: %s <%s>\r\n" + 
		"Subject: %s\r\n" + 
		"Content-type: text/plain; charset=UTF-8\r\n" + 
		"\r\n" +  
		"%s", sender_name, sender_email, receiver.Login, receiver.Email, subject, content)
	
	return smtp.SendMail("smtp.qq.com:25", auth, sender_email, receiver_email, []byte(text))
}

func getReceivers(path string) ([]Receiver, error) {
	fin,err := os.Open(path) 
	defer fin.Close()       
	if err != nil {
		fmt.Printf("Open %s error: %s\n", path, err)
		return make([]Receiver, 1), err
	}
	contents, err := ioutil.ReadAll(fin)
	receiver_raw := strings.Split(strings.Trim(string(contents), "\n"), "\n")
	receivers := make([]Receiver, 0, 10)
	for _, val := range receiver_raw {
		receiver_info := strings.Split(val, "|")
		r := Receiver{
			Login: receiver_info[0],
			Name : receiver_info[1],
			Company : receiver_info[2],
			Location : receiver_info[3],
			Email : receiver_info[4]}
			receivers = append(receivers, r)
	}

	return receivers, err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify file contains email list")
		return
	}

	receivers, err := getReceivers(os.Args[1])
	if err != nil{
		fmt.Println("err :", err)
		return
	}
	for _, recev := range receivers {
		fmt.Printf("Sending email to %s\n", recev.Email)
		err = sendEmail( recev, subject, content)
		if err!= nil {
			fmt.Printf("Sending email to %s error: %s\n", recev.Email, err)
			return
		}
		fmt.Printf("Sending email to %s Complete\n", recev.Email)
	}
}

