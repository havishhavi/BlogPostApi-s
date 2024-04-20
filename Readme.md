
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


//ask sir reflection : meta programming : at runtime if i need to change the data we use reflection : introspect  itself: we use reflection to change the data, manipulate data or type of data at run time.


//middleware:= goes to create post before it goes to middleware the post 