package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dongri/phonenumber"
	"github.com/ochom/event-utils/helpers"
	gohttp "github.com/ochom/go-http"
)

// Payload payload use to send messages
type Payload struct {
	Phone string
	Text  string
}

// Send ...
func (p *Payload) Send(ctx context.Context) error {

	sendURL := "http://api.eleza.online/v1/sms/send/"

	token, err := helpers.MustGetEnv("ELEZA_SMS_TOKEN")
	if err != nil {
		return err
	}

	productID, err := helpers.MustGetEnv("ELEZA_PRODUCT_ID")
	if err != nil {
		return err
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"X-Token":      token,
	}

	data := map[string]string{
		"msisdn":    phonenumber.Parse(p.Phone, "KE"),
		"sms":       p.Text,
		"productID": productID,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json marshal err: %v", err)
	}

	go func() {
		httpClient := gohttp.New(time.Second * 30)
		status, res, err := httpClient.Post(ctx, sendURL, headers, body)
		if err != nil {
			log.Printf("sending message failed %s", err.Error())
			return
		}

		if status != http.StatusOK {
			log.Printf("message failed status %v", status)
			return
		}

		log.Printf("message sent response %s", string(res))
	}()
	return nil
}
