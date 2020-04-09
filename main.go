package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/vault/api"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	vaultAddr := os.Getenv("VAULT_ADDR")
	vaultTokenPath := os.Getenv("VAULT_TOKEN_PATH")
	fmt.Println("VAULT_ADDR:", vaultAddr)
	fmt.Println("VAULT_TOKEN_PATH:", vaultTokenPath)

	b, err := ioutil.ReadFile(vaultTokenPath)
	if err != nil {
		fmt.Println(err)
	}

	vaultToken := string(b)
	var client *api.Client
	if client, err = api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient}); err != nil {
		fmt.Println(err)
	}
	client.SetToken(vaultToken)
	var vaultSecret []byte
	data, err := client.Logical().Read("secret/data/secret")
	if err != nil {
		fmt.Println(err)
		vaultSecret = []byte(err.Error())
	} else {
		vaultSecret, _ = json.Marshal(data.Data)
	}

	fmt.Println(string(vaultSecret))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Demo valet!<h2>")
		fmt.Fprintf(w, "<h1>VAULT_ADDR: %s<h2>", vaultAddr)
		fmt.Fprintf(w, "<h1>VAULT_TOKEN_PATH: %s<h2>", vaultTokenPath)
		fmt.Fprintf(w, "<h1>VAULT_TOKEN: %s<h2>", vaultToken)
		fmt.Fprintf(w, "<h1>VAULT_SECRET: %s<h2>", vaultSecret)
	})
	http.ListenAndServe(":80", nil)
}
