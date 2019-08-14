# Go Rest Api Boilerplate

If you are new to the Go language you should first check out the [docs](https://golang.org/doc/) and get familiar with
 the language. Next, you should clone this repo **outside** your go path (we use go modules for package management).
 Once you are done this, you can run `./setup.sh` to configure the boilerplate to your project. Don't forget to
 to initialize a new `git` repo once you have run the command.
 
This boilerplate aims to give a basic rest api structure in go based on the [go-swagger](https://github.com/go-swagger/go-swagger)
 server generation command, and the [Gorm](https://github.com/jinzhu/gorm) ORM for PostgreSql. It also uses
 [Go-Migrate](https://github.com/golang-migrate/migrate) to handle database migrations.
 
## When To Use Go
- Need for greater speed in data processing (but not for database access and network requests)
- High usage of concurrency
- NOT when there is a package that does what you need (ecosystem is far less developed than Node)
- When you want to write statically typed code that will be cleaner and are willing to put in a bit more time and effort

 
## Getting Started

Clone the repo outside of your `$GOBASE` and then run the command `./setup`, it will guide you thru the process of
 setting up your project. You should then modify the `swagger.yml` file to match the needed api spec. Note that this
 file can be modified as much as possible and the command to re-generate the generated code can be re-run as needed.
 It will only overwrite code in the `./gen` directory so you are safe to write whatever you need in the `internal`
 directory.
 
Once you have set up your project, you should commit the changes.
 
##Commands

### Install
```
go get -u github.com/adeptmind/adept-go-postgres-api-boilerplate
```

### Build
Builds a local binary
```
go build -o main
```

### Run
To run the server from the build:
```
./main
```
To run the server in development mode:
```
go run main.go
```

### Generate code from the swagger file
Used to generate the stub and helper code from `swagger.yml` into `./gen`:
```
mkdir -p ./gen && swagger generate server -A boilerplate -f ./swagger.yml --exclude-main -t ./gen
```

### Database Migrations
We use go-migrate for db migrations.

Install migrate tool (call from outside the project):
```
go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
```
or `brew install go-migrate`

Add a migration:
```
migrate create -ext .sql -dir db/migrations/ MIGRATION_NAME
```

Run a migration (note that if using ssl, the ssl_disabled option should be removed):
```
migrate --path db/migrations/ --database "postgres://postgres@localhost:5432/boilerplate?sslmode=disable" up
migrate --path db/migrations/ --database "postgres://postgres@localhost:5432/boilerplate?sslmode=disable" down
```

### Testing
Go has a builtin test tool that can run tests and performance benchmarks with or without coverage (more details in docs):

```
go test ./...
```

### Formatting
Go has a builtin formatter:
```
go fmt ./...
```