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

![App Screenshot](https://i2.paste.pics/2ed0aa0ef608936256ad59fa6e391918.png)

2. /api/v1/payment

Need jwt token from login, just paste at header Authorization and value is token (no need to add bearer in beginning of token)
```bash
body{
"merchant_id":"string",
"amount":0
}
```
![App Screenshot](https://i2.paste.pics/f3c1b78c91cbebbcf1711515699e03d2.png)

![App Screenshot](https://i2.paste.pics/2e26141f0e4cde0164ecc58fb37ff341.png)

3. /api/v1/logout
```bash
Need jwt token from login, just paste at header Authorization and value is token (no need to add bearer in beginning of token)
```
![App Screenshot](https://i2.paste.pics/eba7ea70ea317be4e57c5e387530f949.png)