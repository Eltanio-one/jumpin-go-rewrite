module github.com/Eltanio-one/jumpin-go-rewrite

go 1.23

toolchain go1.23.1

require github.com/lib/pq v1.10.9 // direct

require (
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.27.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ezzarghili/recaptcha-go.v4 v4.3.0
)

require github.com/gorilla/sessions v1.4.0

require github.com/gorilla/securecookie v1.1.2 // indirect
