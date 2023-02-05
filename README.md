# GooCart



GooCart is a high-performance e-commerce platform backend written in Go language. It is designed to handle a large number of requests and transactions efficiently, making it ideal for large-scale e-commerce operations.



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



## Contributing
We welcome contributions to this project. Please fork the repository and submit a pull request with your changes.