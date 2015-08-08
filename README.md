# apiPOC

## Install Go:
https://golang.org/doc/install#osx (remember to set your $GOPATH env var)

## Add GOPATH's bin to your path
export PATH=$PATH:$GOPATH/bin

## Clone the repo
Clone inside $GOPATH/src

## Install Godep:
go get github.com/tools/godep

## Install dependencies
cd $GOPATH/src/<reponame> && godep restore

## Install gin for live code reloading
go get github.com/codegangsta/gin

## Run server:
gin run (runs on port 3000)