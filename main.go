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
	vaultToken := "token"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Demo valet!<h2>")
		fmt.Fprintf(w, "<h1>VAULT_ADDR: %s<h2>", vaultAddr)
		fmt.Fprintf(w, "<h1>VAULT_TOKEN_PATH: %s<h2>", vaultTokenPath)
		fmt.Fprintf(w, "<h1>VAULT_TOKEN: %s<h2>", vaultToken)
		fmt.Fprintf(w, "<h1>VAULT_MASTER_TOKEN: %s<h2>", vaultMasterToken)
	})
	http.ListenAndServe(":80", nil)
}
