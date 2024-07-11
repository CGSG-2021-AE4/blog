
### Requests

Basic template contains header. Header contains ref to main page and logo maybe.

#### GET /
Button to create article
Some articles titles

#### GET /login/
Username field
Password field
Login button

#### GET /reg/
Username field
Email field
Password field
Register button

#### POST /login/
-> username and password
<- token or error

#### POST /reg/
-> username, email and password
<- complete or error

#### GET /account/
if not authorized -> /login/
username
email
password

#### GET /article/<UUID>
Button edit
Article text

#### GET /article/edit/<UUID>
Button save
Button cancel
Article title in input 
Article text in input

### Structure
cmd/
  main.go
api/
  router/
    articles/
      routes.go
  errors.go
  server.go
  articles.go
internal/
  articles/
    articles.go
    mainpage
web/
  static/
    bin/
      css/
  templates/
    main.html



### Build
#### Backend
go build -o ./out/blog.exe ./cmd/main.go
./out/blog.exe
#### Frontend
node esbuild.js



### Potencial problems and TODOs
Make ability to check if this username is valid
