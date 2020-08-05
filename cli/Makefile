ssl:
	openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out cert.pem

ssl_test:
	openssl genrsa -out cli.key 2048
	openssl req -new -key ./cli.key -out cli.csr
	openssl x509 -req -days 365 \
		-in ./cli.csr \
		-out ./cli.csr \
		-CA ../rootCA.csr \
		-CAkey ../rootCA.key \
		-CAcreateserial


