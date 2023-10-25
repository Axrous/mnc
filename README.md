# Test MNC Back-End Developer

This is for test mnc

## Overview
- using clean arch by modified PZN x Enigma.
- Using uuid for generate id, logrus for logging, simdb for json file database
- Have 3 Endpoint (Login, Payment, Logout).

## Instalation
this project requires [Go](https://golang.org/) v1.20+ to run.

```bash
# Clone this project
$ git clone https://github.com/Axrous/mnc.git

# Move to project dir
$ cd mnc

# Install dependencies
$ go mod download
# or
$ go mod tidy
```

## Run Application
To run this app, you can just type command ```go run .``` in root project dir.
for url ```http://localhost:8080```

## Try Endpoint
1. /api/v1/login
```bash
body {
"username": "username",
"password": "password",
}
```

![App Screenshot](https://paste.pics/67c1a37f9a9cf637fce5eefd474a94b6)

2. /api/v1/payment

Need jwt token from login, just paste at header Authorization and value is token (no need to add bearer in beginning of token)
```bash
body{
"merchant_id":"string",
"amount":0
}
```
![App Screenshot](https://paste.pics/80b8af860efe4ecf7b5d241dad59a4cf)

![App Screenshot](https://paste.pics/a8607cb9879fe10daaacdf869ebc8375)

3. /api/v1/logout
```bash
Need jwt token from login, just paste at header Authorization and value is token (no need to add bearer in beginning of token)
```
![App Screenshot](https://paste.pics/839cf607d3f8d623220cb732d5b05921)