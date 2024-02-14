package services

import (
	"bytes"
	"html/template"
	"log"

	"github.com/bentenison/fastq-server/models"
	"gopkg.in/gomail.v2"
)

type MailConfig struct {
	Host string
	Port int
	From string
	Body string
}

type mailService struct {
	Host string
	Port int
	From string
	Body string
}

func NewMailService(conf *MailConfig) models.MailService {
	return &mailService{
		Host: conf.Host,
		Port: conf.Port,
		From: conf.From,
	}
}

func (ms *mailService) ConfigureMail() error {
	return nil
}

func (ms *mailService) SendMail(to, subject string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", ms.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", ms.Body)

	d := gomail.Dialer{Host: ms.Host, Port: ms.Port}
	if err := d.DialAndSend(m); err != nil {
		log.Println("error sending mail", err)
		return err
	}
	return nil
}

func (ms *mailService) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	// log.Println("mail", buf.String())
	ms.Body = buf.String()

	return nil
}
