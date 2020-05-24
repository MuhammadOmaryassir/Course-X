## Golang REST API - Clean Architecture Principles

![Image of Clean Architecture](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

## Architecture Properties

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. 

## Install Dependencies
```bash
go install
```

## Build
```bash
go build
```

## Run
```bash
go run .
```
```bash
go run *.go
```

## Run Docker container from the app
```bash
# To build the image
    * docker image build -t yourname/gocleanarch .
# To run the container
    * docker run -p 8000:8000 -d yourname/gocleanarch

    * Open the browser and type localhost:8000

```
