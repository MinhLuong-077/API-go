# stock-exchange-backend-202107

# go.mod
name: stockexchange.com (go mod init stockechage.come)

# require package
go get direct-packages (download all packages in require block of go.mod)

# firebase config (config/firebase.go)
set the path to the service-account-file.json (firebase.go line 15)

# install GORM and mySQL driver
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

# Initialize database for Testing
## Install Polluter
```
go get -u github.com/romanyx/polluter
```
## Create Table and Sample Data
```
go test
```