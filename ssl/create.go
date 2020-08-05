package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"math/big"

	"log"
	"os"
)

func main() {
	log.Println("Test creating new ssl")

	/*
		pemData, err := ioutil.ReadFile("../rootCA.key")
		if err != nil {
			log.Panic(err)
		}

		pemBlk, _ := pem.Decode(pemData)

		log.Println(pemBlk)
	*/

	// pemData, err := ioutil.ReadFile("../rootCA.csr")

	privKey, err := rsa.GenerateKey(rand.Reader, 2048)

	log.Println(privKey, "\n", privKey.PublicKey, "\n", err)

	f, err := os.Create("temp.key")
	defer f.Close()

	/*
		err = pem.Encode(f, &pem.Block{
			Type: "PRIVATE KEY",
			Bytes: privKey.
		})

	*/
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
	}

	// Create self-signed.
	newCert, err := x509.CreateCertificate(rand.Reader, &template, &template, &privKey.PublicKey, privKey)
	if err != nil {
		log.Panic(err)
	}

	// DER to file.

	f, err = os.Create("temp.csr")
	defer f.Close()

	err = pem.Encode(f, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: newCert,
	})

	log.Println(newCert, err)
}
