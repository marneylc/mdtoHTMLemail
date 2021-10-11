package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"flag"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	mail "github.com/xhit/go-simple-mail/v2"
)

// read the md file and convert to html
func mdtohtml(filename string) []byte {
	dat, _ := ioutil.ReadFile(filename)
	//fmt.Print(string(dat))
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	md := []byte(dat)
	htmlBody := markdown.ToHTML(md, parser, nil)
	return htmlBody
}

// email sending
func sendmail(
	filename string,
	htmlBody []byte,
	username string,
	password string,
	smtphost string,
	destination string,
) {
	server := mail.NewSMTPClient()
	server.Host = smtphost
	server.Port = 587
	server.Username = username
	server.Password = password
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("<" + username + ">")
	email.AddTo(destination)
	//	email.AddCc("another_you@example.com")
	subject := strings.TrimSuffix(filename, filepath.Ext(filename))
	email.SetSubject(subject)

	email.SetBody(mail.TextHTML, string(htmlBody))
	//	email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var filename, username, password, smtphost, destination string
	flag.StringVar(&filename,"filename", "foo.md", "a filename")
	flag.StringVar(&username,"username", "johnny@gmail.com", "a username")
	flag.StringVar(&password,"password", "secret", "a password")
	flag.StringVar(&smtphost,"smtphost", "smtp.gmail.com", "the smpt host")
	flag.StringVar(&destination,"destination", "johnnyduderino@gmail.com", "the destination inbox")
	flag.Parse()
	htmlbody := mdtohtml(filename)
	sendmail(filename, htmlbody, username, password, smtphost, destination)
}
