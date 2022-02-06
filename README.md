# Go Domain Drived Design / Service repository pattern

Simple api domain drived design / service repository pattern

## API Overview

every api (for each domain model) is structured like below:

    handler <----rest api endpoint
    |
    service <----service layer, main business logic of the api lives here
    |
    store <---- database logic, operation related to database fetching

## Swagger docs
to access the swagger docs, go to:
$APP_HOST/swagger/index.html

to generate new updated swagger docs, install:
```bash
go get -u github.com/swaggo/swag/cmd/swag
```
to generate docs, in the project root directory, run:

```bash
go swag init -g cmd/APPFOLDER/tngrm-APP-api.go docs/APPFOLDER/
```

if you're having issues like "swag command not found" set the go path variable in project root directory:

```bash
export PATH=$(go env GOPATH)/bin:$PATH
```
### note:
remember to generate swagger docs BEFORE running the ./build.sh script.

for other info about go swag, refer to: https://github.com/swaggo/swag
