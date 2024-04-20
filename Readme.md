
To load GORM and Sql
go get -u gorm.io/gorm
go get gorm.io/driver/mysql

// to perform routing and web api
To load Gin
go get -u github.com/gin-gonic/gin

//env files to explictly give values to load the variables to connect the database and the server
To Load Env:  go get github.com/joho/godotenv


// used for sql data migration
To Load Cobra CLI
go get github.com/spf13/cobra

install and run command in the beginning to get the root cli
go install github.com/spf13/cobra-cli@latest
cobra-cli init

to migrate code using cli command
run := go run main.go migrate


//used in router
To Load CORS (Cross Origin Resource Sharing) 
go get github.com/gin-contrib/cors

//for password encryption 
go get golang.org/x/crypto/bcrypt

//jwt go
go get github.com/dgrijalva/jwt-go

// cli for cobra cli 


//reflection : meta programming : at runtime if i need to change the data we use reflection : introspect  itself: we use reflection to change the data, manipulate data or type of data at run time.

Login api 
![image](https://github.com/havishhavi/BlogPostApi-s/assets/164078377/dc756c57-2171-45ef-badf-12b683087224)



{
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        "ID": 3,
        "Name": "meena",
        "Email": "meena@gmail.com",
        "Password": "",
        "JwtToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVU0VSX0lEIjoiMyIsIlJPTEVfSUQiOjEsImV4cCI6MTc0NTE1OTY3MywiaWF0IjoxNzEzNTM3MjczLCJpc3MiOiJKV3QgQXV0aG9yaXphdGlvbiJ9.kBshpIji9yxH0rgXBiYjxSNmiP4CPt8d6oP_-gCw7bk",
        "Mobile": 9999999999,
        "Active": 1,
        "Created": "2024-04-19T10:27:52.227-04:00",
        "Updated": "2024-04-19T10:27:52.227-04:00"
    }
}

//middleware:= goes to create post before it goes to middleware the post 
