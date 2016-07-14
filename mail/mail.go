package mail

import "net/smtp"

// Gmail send an email using a gmail account
func Gmail(from string, password string, to string, subject string, body string) error {
	return smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth(subject, from, password, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte("To: "+to+"\r\n"+
			"Subject: "+subject+"\r\n"+
			"\r\n"+
			body+"\r\n"))
}
