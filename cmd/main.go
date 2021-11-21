package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/damek86/go-windhager-infowin/pkg/api"
)

func main() {
	urlPtr := flag.String("url", "192.168.4.100", "remote address")
	usernamePtr := flag.String("username", api.DefaultCustomerUsername, "username")
	passwordPtr := flag.String("password", api.DefaultCustomerUserPassword, "password")

	flag.Parse()

	client := api.NewClient(*urlPtr, *usernamePtr, *passwordPtr)
	work(client)
}

func work(client api.Client) {
	data, err := client.GetDataPoint("1/15/0/117")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", prettyPrint(data))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
