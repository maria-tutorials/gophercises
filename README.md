# gophercises-url-shortner
The goal of this exercise is to create an `http.Handler` that will look at the path of any incoming web request and determine if it should redirect the user to a new page, like URL shortener would.

## Part 1
Implement the methods in the handler.go file

## Part 2
- Update the main/main.go source file to accept a YAML file as a flag and then load the YAML from a file rather than from a string.
- Build a JSONHandler that serves the same purpose, but reads from JSON data.

## Extra
- Build a Handler that instead reads from a database (boltDb)
