package subscriber

import (
	"context"
	"log"

	gomail "gopkg.in/gomail.v2"

	common_proto "github.com/Ankr-network/dccn-common/protos/common"
)

var (
	sender = "994336359@qq.com"
	passwd = "espgstzviouubfjh"
)

func SendMail(e *common_proto.MailEvent) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", sender, "Ankr Network")
	m.SetHeader("To",
		m.FormatAddress(e.Address, e.Name),
	)
	m.SetHeader("Subject", "Gomail")
	m.SetBody("text/html", e.Message)

	d := gomail.NewPlainDialer("smtp.qq.com", 465, sender, passwd)
	if err := d.DialAndSend(m); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func Handler(ctx context.Context, e *common_proto.MailEvent) error {
	log.Println("Function Received message: ", e)

	return SendMail(e)
}
