Go JWT Authentication
This project is a basic JWT authentication system built with Golang. It uses Gin for the HTTP framework, Gorm for the ORM, and MySQL as the database.

Prerequisites
Before you begin, ensure you have Golang installed on your system.

Installation
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

Technology Stack
Golang: Programming language.
Gin: HTTP web framework.
Gorm: ORM library.
JWT: JSON Web Tokens for authentication.
MySQL: Database.

Database Connection

Replace the database connection string in the .env file with your specific database configuration.

Secret Key - You can also change the secret key for JWT in the .env file.


go build -o go-jwt-authentication

Run the executable binary:
./go-jwt-authentication

Testing
Use Postman/Thunder client to test the API endpoints. Remember to add the token only to the cookies. The API endpoints available for testing are:

Signup: POST http://localhost:9090/signup
Login: POST http://localhost:9090/login
Validate: GET http://localhost:9090/validate

