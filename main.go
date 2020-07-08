package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public/frontend")))

	if СheckExists("cert.pem") && СheckExists("key.pem") {
		//go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
		fmt.Println("SSL web server up")
		http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	} else {
		fmt.Println("Plain web server up")
		http.ListenAndServe(":8080", nil)
	}
}

// СheckExists - проверяем что файл существует
func СheckExists(filename string) bool {

	if _, err := os.Stat(filename); err == nil {
		return true
	}

	return false
}
