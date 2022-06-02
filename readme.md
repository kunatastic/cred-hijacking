# Cred-Hijacking

A very noobish Proof-of-concept of why you should use github ssh way rather than https. In this demonstration, I have a very simple binary executable created using Go that could be utilise in credential hijacking.

So basically, when we are using the github-cli tool and input the credentials it will store the credentials as

```sh
https://<github_username>:<access_token>@github.com
```

in the `~/.git-credentials` folder and when `access_token` can be accessed any person can make operations as being you. According to the update after April 2021, Github treats the access token as the password of the user for some places.

> Also this tool was just for demonstration purpose

## Prerequisites

We have a very limited requirement for this project which is just go any version after 1.18 will get the job done.

### GO

* Linux and Mac-OS

  ```sh
  wget -c https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
  sudo tar -C /usr/local -xvzf go1.15.2.linux-amd64.tar.gz
  export  PATH=$PATH:/usr/local/go/bin
  export GOBIN="$GOPATH/bin"
  ```

* Windows
  Installer [here](https://go.dev/doc/install). Pheww

## Setup

1. Clone the repository :eyes::eyes:

```sh
git clone https://github.com/kunatastic/cred-hijacking
cd cred-hijacking
```

2. Open 2 terminals and run the following commands

* Start the server in the first terminal

```sh
go run server/main.go
```

* Start the application in the second terminal

```sh
go run application/main.go
```

## Building the package

```sh
go build -o build/application application/main.go
go build -o build/server server/main.go
```

Once the build is done, you can run the application and server in the terminal without even requiring Go.

## Ending Note

How to avoid this? I am no expert here but I am just consolidating the points that I have learned so far.

1. To avoid saving the `access token` permanently, you can use `"cache --timeout=<time in seconds>"` flag to store the token in the cache that will get clear after the timeout you have specified. This method could be length everytime you need to create a new token and paste it.

The default is 15 minutes, you can set a longer timeout with:

```sh
git config --global credential.helper "cache --timeout=3600"
```

2. The another solution could be use github with ssh all the process will be same as using https. Just rather than directly storing the token in `~/.git-credentials` now the credentials would be stored as a SSH key which is much more secure. To setup the SSH key, you can use the refer this documentation [Here](https://help.github.com/en/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent).

## Thank you
