package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Start..")

	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
		}

		go func(c net.Conn) {
			defer c.Close()

			log.Println("Client connected ", c)

			totalRecvSize := 0

			recvData := make([]byte, 0)

			for {
				buffer := make([]byte, 1024)
				n, err := c.Read(buffer)
				if err != nil {
					log.Println("bye ", err)

					log.Println("Check ", checkValidity(recvData))
					break
				}

				totalRecvSize += n
				recvData = append(recvData, buffer[0:n]...)

				log.Println(n, totalRecvSize)
				c.Write([]byte("OK"))

				// Check.
			}

		}(c)
	}

}

func checkValidity(data []byte) bool {
	log.Println("Len ", len(data))

	for i, data := range data {
		expectValue := byte(i % 256)
		if expectValue != data {
			log.Println("False ", i, expectValue, data)
			return false
		}
	}
	return true
}
