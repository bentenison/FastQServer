server:
	- go run *.go
tidy:
	- go mod tidy
keys:
	- openssl genpkey -algorithm RSA -out ./rsa_private_dev.pem -pkeyopt rsa_keygen_bits:2048
	- openssl rsa -in ./rsa_private_dev.pem -pubout -out ./rsa_public_dev.pem