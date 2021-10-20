# Backend Service Aplikasi Pelelangan Ikan Terintegrasi

## Folder Structure
```
.
|
├── config      # contains config set up
├── constant    # contains constant variable
├── database    # database connection
├── delivery    
|   ├── http    # HTTP delivery (API)
├── entites     # entities 
├── helper      # contains all services
├── middleware  # middleware that used
├── repository  # all database logic
├── services    # service that can be used
├── template    # template for html format
├── usecase     # all business logic
│   |
.   .
```

## Getting Started
### Install dependencies
#### Go
https://golang.org/doc/install/source

### Setup Database
#### Development
- Edit database config in config.json
#### Production
- Set env in server using `export`

### Build
`go build`

### Migrate and Seed Database
- ```./backend-tpi migrate```
- ```./backend-tpi seed```

### Login
- It will create superadmin account with username superadmin and password superadmin
- Login with superadmin account and change the password
