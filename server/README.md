# Restfull api 


# How to run 

- create .env file 

```
DB_USERNAME = 
DB_PASSWORD = 
ATRO_HOST   = 
DB_PORT     = 3306
DB_NAME     = 
JWT_SECRET  = 

IMAGE_SAVE_PATH = 
IMAGE_MAX_SIZE  = 
```

- run
```
go mod init atro
go mod tidy
go run main.go
```

# Pack 

```
github.com/dgrijalva/jwt-go         Jwt 
github.com/go-sql-driver/mysql      Mysql-driver 
github.com/gorilla/mux              Routing 
github.com/jinzhu/gorm              ORM
```