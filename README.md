# fomo-go-sdk

## Introduction
Golang client for [Fomo API](https://docs.fomo.com/reference).

## Install
Install with `go get`:

```bash
$ go get github.com/khoa-le/fomo-go-sdk
```

## Usage
```go
package main

import (
    "fmt"
    "os"

    "github.com/khoa-le/fomo-go-sdk"
)

const (
    apiKey = "oH123EKegQRG-6HvX4pT_X"
)
func main() {
	response := new(interface{})
	client := fomo.New(apiKey)
	attributes := map[string]string{"count": "200"}
	event := fomo.EventBasic{
		EventTypeId: 122413,
		Url:         "https://precita.vn/customer/account/create/",
		Title:       "Precita",
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
```
