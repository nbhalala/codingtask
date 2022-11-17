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

Go command that ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.
```
go mod tidy
```

Go command that initializes and writes a new go.mod file in the current directory, in effect creating a new module rooted at the current directory. The go.mod file must not already exist.

```
go mod init
```

Go command with flag "-modcache" causes go clean to remove the entire module cache, including unpacked source code of versioned dependencies.
```
go clean -modcache
```

## Compiling/Building the App
- Execution steps with results:
```
naresh@LNAR-5CG0338GPC:~/codingtask$ docker-compose down
naresh@LNAR-5CG0338GPC:~/codingtask$ docker-compose up
[+] Building 866.6s (16/16) FINISHED
 => [internal] load build definition from Dockerfile                                                                                                     0.0s
 => => transferring dockerfile: 32B                                                                                                                      0.0s
 => [internal] load .dockerignore                                                                                                                        0.0s
 => => transferring context: 2B                                                                                                                          0.0s
 => [internal] load metadata for docker.io/library/golang:alpine                                                                                       846.2s
 => [internal] load metadata for docker.io/library/alpine:latest                                                                                       566.9s
 => [builder 1/5] FROM docker.io/library/golang:alpine@sha256:d171aa333fb386089206252503bc6ab545072670e0286e3d1bbc644362825c6e                           0.0s
 => [internal] load build context                                                                                                                        0.8s
 => => transferring context: 33.96MB                                                                                                                     0.8s
 => [stage-1 1/5] FROM docker.io/library/alpine@sha256:b95359c2505145f16c6aa384f9cc74eeff78eb36d308ca4fd902eeeb0a0b161b                                  0.0s
 => CACHED [stage-1 2/5] RUN adduser -S -D -H -h /app appuser                                                                                            0.0s
 => CACHED [builder 2/5] RUN mkdir /build                                                                                                                0.0s
 => [builder 3/5] ADD . /build/                                                                                                                          0.2s
 => [stage-1 3/5] COPY . /app                                                                                                                            0.3s
 => [builder 4/5] WORKDIR /build                                                                                                                         0.1s
 => [builder 5/5] RUN go build -o main.                                                                                                                 18.5s
 => [stage-1 4/5] COPY --from=builder /build/main /app/                                                                                                  0.1s
 => [stage-1 5/5] WORKDIR /app                                                                                                                           0.1s
 => exporting to image                                                                                                                                   0.3s
 => => exporting layers                                                                                                                                  0.3s
 => => writing image sha256:1867a51eca959f4d80d515295e44dfd47460a5084212012fdfe19c816bc554c6                                                             0.0s
 => => naming to docker.io/library/codingtask-api                                                                                                        0.0s

Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
[+] Running 3/3
 ⠿ Network codingtask_default  Created                                                                                                                   0.8s
 ⠿ Container codingtask-db-1   Created                                                                                                                   0.2s
 ⠿ Container codingtask-api-1  Created                                                                                                                   0.1s
Attaching to codingtask-api-1, codingtask-db-1
codingtask-db-1   | 1:C 17 Nov 2022 02:53:35.008 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
codingtask-db-1   | 1:C 17 Nov 2022 02:53:35.008 # Redis version=7.0.5, bits=64, commit=00000000, modified=0, pid=1, just started
codingtask-db-1   | 1:C 17 Nov 2022 02:53:35.008 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
codingtask-db-1   | 1:M 17 Nov 2022 02:53:35.008 * monotonic clock: POSIX clock_gettime
codingtask-db-1   | 1:M 17 Nov 2022 02:53:35.012 * Running mode=standalone, port=6379.
codingtask-db-1   | 1:M 17 Nov 2022 02:53:35.013 # Server initialized
codingtask-db-1   | 1:M 17 Nov 2022 02:53:35.013 # WARNING overcommit_memory is set to 0! Background save may fail under low memory condition. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
codingtask-db-1   | 1:M 17 Nov 2022 02:53:35.016 * Ready to accept connections
```
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
![image](https://user-images.githubusercontent.com/42828619/202346963-e28819ea-ef0a-4730-a255-6cfc599fbafc.png)

![image](https://user-images.githubusercontent.com/42828619/202348396-bf21e8d8-25e6-46a3-8041-5ef03d444e12.png)

