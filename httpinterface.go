package main

import (
	"fmt"
	"net/http"
	"os"
)

func http_resp_interface() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("An error occured: %s", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
