package main

import (
	"crypto/tls"
	"log"
)

func main() {
	log.Println("Start Client")

	//cert, err := tls.LoadX509KeyPair("./c7ad737f37-certificate.pem.crt", "./c7ad737f37-private.pem.key")
	// cert, err := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	cert, err := tls.LoadX509KeyPair("./cli.csr", "./cli.key")
	if err != nil {
		log.Panic(err)
		return
	}

	cfg := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	// c, err := tls.Dial("tcp", "127.0.0.1:8080", cfg)
	c, err := tls.Dial("tcp", "34.64.77.83:8080", cfg)
	if err != nil {
		log.Panic(err)
		return
	}

	n, err := c.Write([]byte("hello"))
	log.Println(n, err)

	defer c.Close()
}
