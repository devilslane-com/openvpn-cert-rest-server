package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func generateCertificate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Replace "client1" with your method of dynamically setting client names
	clientName := "client1"

	// Command to source vars and build the client key and certificate
	// Ensure these commands are correct for your Easy-RSA setup
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("cd /etc/openvpn/easy-rsa; "+
		"source vars; "+
		"./build-key --batch %s", clientName))

	if err := cmd.Run(); err != nil {
		log.Printf("Error generating certificate: %s", err)
		http.Error(w, "Failed to generate certificate", http.StatusInternalServerError)
		return
	}

	// Construct file paths
	basePath := "/etc/openvpn/easy-rsa/keys/"
	certPath := basePath + clientName + ".crt"
	keyPath := basePath + clientName + ".key"

	// Read and send the certificate file
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Printf("Error reading certificate file: %s", err)
		http.Error(w, "Failed to read certificate file", http.StatusInternalServerError)
		return
	}

	keyData, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Printf("Error reading key file: %s", err)
		http.Error(w, "Failed to read key file", http.StatusInternalServerError)
		return
	}

	// Respond with the certificate and key
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(certData)
	w.Write([]byte("\n"))
	w.Write(keyData)
}

func main() {
	http.HandleFunc("/generate-cert", generateCertificate)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
