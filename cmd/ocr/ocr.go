package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"log"
)

func main() {
	client := gosseract.NewClient()
	client.Languages = []string{"eng"}

	wlErr := client.SetWhitelist("abcdefghijklmnopqrstuvwxyz0123456789")
	if wlErr != nil {
		log.Fatalf("Unable to set whitelist for tesseract: %v", wlErr)
	}

	setErr := client.SetImage("km_test.png")
	if setErr != nil {
		log.Fatalf("Unable to set image to read: %v", setErr)
	}

	text, toTextErr := client.Text()
	if toTextErr != nil {
		log.Fatalf("Unable to extract text: %v", toTextErr)
	}

	fmt.Println(text)

	client.Close()
}
