package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	vaultAddr := os.Getenv("VAULT_ADDR")
	vaultTokenPath := os.Getenv("VAULT_TOKEN_PATH")
	vaultMasterToken := os.Getenv("VAULT_MASTER_TOKEN")
	fmt.Println("VAULT_ADDR:", vaultAddr)
	fmt.Println("VAULT_TOKEN_PATH:", vaultTokenPath)
	fmt.Println("VAULT_MASTER_TOKEN:", vaultMasterToken)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(":80", nil)
}
