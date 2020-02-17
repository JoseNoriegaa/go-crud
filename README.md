
# SIMPLE GO CLI-CRUD

> NOTE: This is not a professional project. It is only for learning purposes.

## Installation

To install the project, there are the following ways to do it:

1. git clone --single-branch --branch master https://github.com/JoseNoriegaa/go-crud.git $GOPATH/src/github.com/josenoriegaa/go-crud
2. go get github.com/JoseNoriegaa/go-crud

# Setup
1. ```cd $GOPATH/src/github.com/josenoriegaa/go-crud```
3. Create a `.env` file into the project directory and set the `DB_CONNECTION_STRING` variable.
4. Run the following commands to install the dependencies:


```terminal
$ go get http://github.com/jinzhu/gorm
$ go get github.com/jinzhu/gorm/dialects/mysql
$ go get github.com/joho/godotenv
$ go get github.com/briandowns/spinner
$ go get github.com/wayneashleyberry/terminal-dimensions
```

# Usage
1. ```cd $GOPATH/src/github.com/josenoriegaa/go-crud```
2. ```go run main.go```
