# GO-BOOKSTORE

Go bookstore is a Go CRUD app using Postgres and gorilla mux.

## TO RUN/BUILD APP
 Navigate to cmd/main
 To Build the app: go build <br/>
 To run the application :  go run main
 
## Technical Dependencies:

Using gorilla/mux for routes <br/>
Using postgres as the Database <br/>
To install mux: go get "http://github.com/gorilla/mux" <br/>
To install gorm and postgres: go get "http://github.com/jinzhu/gorm" and go get "http://github.com/jinzhu/gorm/dialects/postgres"<br/>
DB connection is to the local postgres database. I've given my default connection credentials in pkg/config/app.go Connect function.
