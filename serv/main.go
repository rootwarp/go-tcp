package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	log.Println("hello world")

	cert, err := tls.LoadX509KeyPair("./fullchain.pem", "./privkey.pem")
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	cfg.ClientAuth = tls.RequireAnyClientCert
	cfg.VerifyPeerCertificate = func(certificates [][]byte, _ [][]*x509.Certificate) error {
		certs := make([]*x509.Certificate, len(certificates))
		for i, asn1Data := range certificates {
			cert, err := x509.ParseCertificate(asn1Data)
			if err != nil {
				log.Println(err)
				return err
			}

			log.Println("Issuer ", cert.Issuer.String())
			log.Println("Serial ", cert.SerialNumber)
			log.Println("Subject", cert.Subject.String())

			// TODO: Add client verification code here.
			// TODO: How do I pass identifications?
			certs[i] = cert
		}

		return nil
	}

	l, err := tls.Listen("tcp", "0.0.0.0:8080", cfg)
	if err != nil {
		log.Panic(err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Panic(err)
			continue
		}

		go func(c net.Conn) {
			defer c.Close()

			fmt.Printf("Conn %+v\n", c.(*tls.Conn))
			tlsConn := c.(*tls.Conn)

			tlsConn.Handshake()

			var connState tls.ConnectionState

			for {
				connState = tlsConn.ConnectionState()
				if connState.HandshakeComplete {
					break
				}
				time.Sleep(100 * time.Millisecond)
			}

			// Get serial number.
			fmt.Println(connState.PeerCertificates[0].SerialNumber)
			// TODO: Get device ID.

			recvData := make([]byte, 0)
			count := 0

			for {
				buf := make([]byte, 1024)
				n, err := c.Read(buf)
				if err != nil {
					log.Println("bye ", count, verifyValidity(recvData))
					break
				}

				count += n

				recvData = append(recvData, buf[0:n]...)

				log.Println(n, count)
			}

		}(c)

		// c.Write([]byte("OK"))
	}
}

func verifyValidity(data []byte) bool {
	for i, d := range data {
		if byte(i%256) != d {
			return false
		}
	}

	return true
}
