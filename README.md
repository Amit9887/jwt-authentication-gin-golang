Go JWT Authentication
Welcome to the Go JWT Authentication project! This repository contains a robust and secure JWT authentication system built with Golang. It leverages the power of Gin for the HTTP framework, Gorm for ORM, and MySQL as the database. Let's get you started!

✨ Features
JWT Authentication: Secure authentication using JSON Web Tokens.
Gin Framework: Fast and easy-to-use HTTP web framework.
Gorm ORM: Powerful ORM library for Golang.
MySQL Database: Reliable and scalable database solution.
�� Prerequisites
Before you begin, ensure you have the following installed on your system:

Golang
�� Installation
Follow these steps to set up the project on your local machine:

Clone the repository:

    git clone https://github.com/your-repo/go-jwt-authentication.git
    cd go-jwt-authentication
    



Install the required packages:

    go get -u gorm.io/gorm
    go get -u gorm.io/driver/mysql
    go get -u github.com/gin-gonic/gin
    go get -u golang.org/x/crypto/bcrypt
    go get -u github.com/golang-jwt/jwt/v5
    go get github.com/lpernett/godotenv
    go get github.com/githubnemo/CompileDaemon
    



��️ Configuration
Database Connection:

Replace the database connection string in the .env file with your specific database configuration.

Secret Key:

You can also change the secret key for JWT in the .env file.

�� Build and Run
Build the project:

    go build -o go-jwt-authentication
    



Run the executable binary:

    ./go-jwt-authentication
    



�� Testing
Use Postman or Thunder Client to test the API endpoints. Remember to add the token only to the cookies. The API endpoints available for testing are:

Signup:

    POST http://localhost:9090/signup
    
Login:

    POST http://localhost:9090/login

Validate:

    GET http://localhost:9090/validate
    



��️ Technology Stack
Golang: Programming language.
Gin: HTTP web framework.
Gorm: ORM library.
JWT: JSON Web Tokens for authentication.
MySQL: Database.