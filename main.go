package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	data "github.com/Bet20/bbault/data"
	str "github.com/Bet20/bbault/structures"
	"github.com/gorilla/mux"
)

const PORT = ":8081"
const DIRNAME = "storage"
const FILENAME = "vault.json"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to your storage.")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	u, _ := str.GetUser(vars["user"], vars["pass"])
	fmt.Fprintf(w, "%v", u)
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func createStorageFile(name string) {
	_, err := os.Stat(DIRNAME)
	if os.IsNotExist(err) {
		err := os.Mkdir(DIRNAME, 0755)
		checkError(err, "failed on creating storage dir.")
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Insert a new user: ")
	user, _ := reader.ReadString('\n')

	fmt.Printf("Insert a password: ")
	pass, _ := reader.ReadString('\n')

	defUser := str.User{
		Usern: data.Higienize(user),
		Pass:  data.Higienize(pass),
		Items: []str.Item{},
	}

	defUser.LogObject()

	file, err := json.Marshal(defUser)
	checkError(err, "")
	err = ioutil.WriteFile(data.AppendPaths([]string{DIRNAME, FILENAME}), file, 0644)
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
			return
		case "nuser":
			return
		case "ls":
			u, err := str.GetUser(data.Higienize(os.Args[2]), data.Higienize(os.Args[3]))
			checkError(err, "")
			for _, o := range u.Items {
				o.LogObject()
			}
		case "run":
		}
	}

	fmt.Printf("%s", "Storage initiated...")

	server()
}

func server() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/get/{user}/{pass}", getUser)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
