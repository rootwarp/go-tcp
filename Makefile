root_ca_key:
	openssl genrsa -aes256 -out rootCA.key 2048
root_ca_csr:
	openssl req -new -key ./rootCA.key -out rootCA.csr
	openssl x509 -req -days 3650 -in ./rootCA.csr -signkey ./rootCA.key -out rootCA.csr -set_serial 1

gen_cli_ssl:
	openssl genrsa -out cli.key 2048
	openssl req -new -key ./cli.key -out cli.csr
	openssl x509 -req -days 365 -in ./cli.csr -out ./cli.csr -signkey ./rootCA.key -set_serial 1000
