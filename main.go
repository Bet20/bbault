package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const PORT = ":8081"
const DIRNAME = "storage"
const FILENAME = "vault.json"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to your storage.")
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func acessStorageFile(path string) {

}

func createStorageFile(name string) {
	_, err := os.Stat(DIRNAME)
	if os.IsNotExist(err) {
		err := os.Mkdir(DIRNAME, 0755)
		checkError(err, "failed on creating storage dir.")
	}

	//f, err := os.Create("storage/vault.json")
	//fcheckError(err, "Error creating the vault file.")
	//defer f.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Insert a new user: ")
	user, _ := reader.ReadString('\n')

	fmt.Printf("Insert a password: ")
	pass, _ := reader.ReadString('\n')

	defUser := User{
		User:  higianize(user),
		Pass:  higianize(pass),
		Items: []Item{},
	}

	defUser.logObject()

	file, err := json.Marshal(defUser)

	checkError(err, "")
	err = ioutil.WriteFile(appendPaths([]string{DIRNAME, FILENAME}), file, 0644)
	checkError(err, "")
}

func checkError(e error, message string) {
	if e != nil {
		fmt.Errorf("%s", message)
		panic(e)
	}
}

func main() {

	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "create":
			createStorageFile("storage")
		case "nuser":

		}

	}

	fmt.Printf("%s", "Storage initiated...")

	//server()
}

func server() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(PORT, nil)
}
