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
Installation & Configuration

Using Windows Powershell (As an administrator)
- Installing Virtual Machine Platform
- Installing Windows Subsystem for Linux (WSL V2)
- Downloading & Installing WSL Kernel
- Downloading & Installing GUI App Support
- Downloading & Installing Ubuntu (20.04 LTS Linux distribution)
```
wsl --install
```

Using Windows Powershell (As an administrator)
- Start WSL (dev/build environment)
```
Windows PowerShell
PS C:\Users\nbhalala> wsl
naresh@LNAR-5CG0338GPC:/mnt/c/Users/nbhalala$
naresh@LNAR-5CG0338GPC:/mnt/c/Users/nbhalala$ cd /home/naresh/codingtask/
naresh@LNAR-5CG0338GPC:~/codingtask$
```

Use MS Visual Studio (IDE) for Coding
- Start IDE
```
naresh@LNAR-5CG0338GPC:~/codingtask$ code . &
```

Install other tools/software
```
sudo apt update
sudo apt install docker.io -y
sudo apt install docker-compose -y
sudo apt install golang-go -y
```

Go Environment (GOPATH, GOROOT, etc)
```
$ go env
$ go env GOPATH
$ go env GOROOT
$ go env GO111MODULE
```

Docker Commands:

Stops containers and removes containers, networks, volumes, and images created by up.
```
docker-compose down
```

Verify docker Configuration.
```
docker-compose config
```


Builds, (re)creates, starts, and attaches to containers for a service.
```
docker-compose up
```

GO Commands:

go mod tidy ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.
```
go mod tidy
```

## Compiling the App
- TBD

## Usage

Using Postman:

```
POST http://localhost:8000/api/v1/
```
  
  Body:
  ```
  {"url":"https://www.capgemini.com"}
  ```

Response:
  ```
  {
    "url": "https://www.capgemini.com"
    "short": "localhost:8000/<ID>"
  }
  ```
