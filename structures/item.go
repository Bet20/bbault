package structures

import "fmt"

type Item struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func (o Item) logObject() {
	fmt.Printf("%v", o)
}
