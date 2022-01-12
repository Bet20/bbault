package structures

import "fmt"

type User struct {
	User  string `json:"User"`
	Pass  string `json:"Pass"`
	Items []Item `json:"Items"`
}

func (o User) logObject() {
	fmt.Printf("%v", o)
}
