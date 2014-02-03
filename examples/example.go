package main

import (
	"fmt"
	"github.com/johnwesonga/goushahidi/ushahidi"
)

const (
	ushahidiInstance = "http://myushahidi.com"
)

func main() {
	client := client.NewClient(ushahidiInstance, nil)
	fmt.Println("Ushahidi Examples Here")

}
