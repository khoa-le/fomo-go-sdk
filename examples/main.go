package main

import (
	"fmt"

	fomo "github.com/khoa-le/fomo-go-sdk"
)

func main() {
	response := new(interface{})
	client := fomo.New("oH9iiEKegQRG-6HvX4pT_Q")
	//attribute := fomo.Attribute{Key: "count", Value: "200"}
	attributes := map[string]string{"count": "200"}
	event := fomo.EventBasic{
		EventTypeId: 122413,
		Url:         "https://precita.vn/customer/account/create/",
		Title:       "Precita",
		Message:     "Có 5 người đang trên trang này",
		Attributes:  attributes,
	}
	body := make(map[string]fomo.EventBasic)
	body["event"] = event
	err := client.Request("POST", "/api/v1/applications/me/events", nil, body, response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(response)
}
