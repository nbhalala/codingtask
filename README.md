# Coding Task

## Author
Naresh Bhalala - naresh.bhalala@gmail.com

## Aim of this coding task
This project is an coding exercise.
We'd like you to write a simple URL shortener microservice, similar to [tinyurl.com](http://tinyurl.com) or [bit.ly](bit.ly). When running your service, we should be able to make a request to "/shorten" with a URL of our choice; this should return us a unique URL on the service domain. Visiting this URL should redirect us to our original link.

## Pre-requisites
1. Windows Subsystem for Linux (WSL) 
2. GoLang
3. Fiber (Go web framework built on top of Fasthttp, the fastest HTTP engine for Go)
4. Docker
5. Docker-Compose
6. Redis
7. Dependencies (Packages)

## Setting up the build environment
- TBD

## Compiling the App
- TBD

## Aim of this coding task
- TBD

## TODO
- TBD


## Building
- TBD
It is expected that this application will be built on go 1.16. I've been using an alpine base.

## Usage

```
Using Postman:
```

POST http://localhost:8000/api/v1/
  
  Body:
  {"url":"https://www.capgemini.com"}

Response:
  {
    "url": "https://www.capgemini.com"
    "short": "localhost:8000/<ID>"
  }
