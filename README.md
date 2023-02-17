# GooCart



GooCart is a high-performance e-commerce platform backend written in Go language. It is designed to handle a large number of requests and transactions efficiently, making it ideal for large-scale e-commerce operations!



## Technologies and tools used

- Language : Go
- Framwork : Fiber
- Database : Postgresql GORM
- JSON Web Token (JWT) authentication for secure user authentication
- Amazon S3 bucket for storing data
- Stripe Payment API for handling transactions
- Twilio API for OTP verification
- Docker



## Run On local machine

clone this project

```
git clone https://github.com/mohdjishin/GooCart
```

open GooCart Directory

```
cd GooCart
```

download dependencies

```
go get
```

run

```
go run *.go
```

app is listening on port 8080



## Run using makefile

clone this project

```
git clone https://github.com/mohdjishin/GooCart
```

open GooCart Directory

```
cd GooCart
```

run makefile
```
make all
```
## Adding .env file
- Create a new file in the root of your project directory and name it .env.
- Add the following information to the file:
```
#port
PORT=8080
#database
DNS= "host=<host> user=<username> password=<password> dbname=<databsename> port=5432 sslmode=disable"

#JWT encryption key
SECRET= <SecretKey>

#Twilio
TWILIO_ACCOUNT_SID=<TWILIO_ACCOUNT_SID>
TWILIO_AUTH_TOKEN = <TWILIO_AUTH_TOKEN >
VERIFY_SERVICE_SID= <VERIFY_SERVICE_SID >

#AWS
AWS_REGION=<AWS_REGION>
AWS_ACCESS_KEY_ID=<AWS_ACCESS_KEY_ID>
AWS_SECRET_ACCESS_KEY=<AWS_SECRET_ACCESS_KEY>

#Stripe
PAYMENT_SEC_KEY=<Stripe PAYMENT_SEC_KEY>
```


## Contributing
We welcome contributions to this project. Please fork the repository and submit a pull request with your changes.
<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />  <img src="https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white"/>
