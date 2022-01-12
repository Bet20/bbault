package structures

import (
	"encoding/json"
	"fmt"

	"github.com/Bet20/bbault/data"
)

const PATH = "storage/vault.json"

type User struct {
	Usern string `json:"User"`
	Pass  string `json:"Pass"`
	Items []Item `json:"Items"`
}

func (o User) LogObject() {
	fmt.Printf("%v", o)
}

func GetUser(username string, password string) (User, error) {
	bytes := data.AcessStorageFile(PATH)
	fmt.Printf("username: %v, password: %v", username, password)
	var entries Entries
	json.Unmarshal(bytes, &entries)
	for _, user := range entries.Users {
		if user.Usern == username && user.Pass == password {
			PrintItems(user.Items)
			return user, nil
		}
	}
	return User{}, fmt.Errorf("error found! :(")
}

func GetUserByte(username string, password string) ([]byte, error) {
	bytes := data.AcessStorageFile(PATH)
	fmt.Printf("username: %v, password: %v", username, password)
	var entries Entries
	json.Unmarshal(bytes, &entries)
	for _, user := range entries.Users {
		if user.Usern == username && user.Pass == password {
			PrintItems(user.Items)
			us, _ := json.Marshal(user)
			return us, nil
		}
	}
	return []byte{}, fmt.Errorf("error found! :(")
}
