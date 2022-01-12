package structures

import "fmt"

type Item struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func (o Item) LogObject() {
	fmt.Printf("%v", o)
}

func PrintItems(items []Item) {
	fmt.Printf("\n\n--- ITEMS ---\n")
	for _, item := range items {
		fmt.Printf("|-------------\n")
		fmt.Printf("| - NAME: %s\n", item.Name)
		fmt.Printf("| - USERNAME: %s\n", item.Username)
		fmt.Printf("| - PASSWORD: %s\n", item.Password)
	}
	fmt.Printf("| -------------\n")
}
