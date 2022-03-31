# GO-BOOKSTORE

Go bookstore is a Go CRUD API service using Postgres and gorilla mux.

## TO RUN/BUILD APP

 Navigate to cmd/main
 ```sh
 To Build the app:  go build
 To run the application : go run main
 ```
## To use the bot in Slack

Type below message in chat, in same workspace <br/>
@AgeBot my yob is "year"<br/>
eg : @AgeBot my yob is 1997<br/>

## Technical Dependencies:

* Using gorilla/mux for routes <br/>
* Using postgres as the Database <br/>
* To install mux: go get "http://github.com/gorilla/mux" <br/>
* To install gorm and postgres: go get "http://github.com/jinzhu/gorm" and go get "http://github.com/jinzhu/gorm/dialects/postgres"<br/>
* DB connection is to the local postgres database. I've given my default connection credentials in pkg/config/app.go - Connect() function. <br/>
* The postman collection file is provided in the root of the repo : [BookStore-Collection](https://github.com/Clint-Mathews/GO-BOOKSTORE/blob/main/Book_Store.postman_collection.json) <br/>
* This collection can be directly imported to postman and used to perform the operations in app.
