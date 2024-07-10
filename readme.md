
### Pages

Basic template contains header. Header contains ref to main page and logo maybe.

#### /
Button to create article
Some articles titles

#### /article/<UUID>
Button edit
Article text

#### /article/edit/<UUID>
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
go build -o ./out/blog.exe ./cmd/main.go

./out/blog.exe


### Potencial problems and TODOs
Make ability to check if this username is valid
