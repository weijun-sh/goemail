#GORUN = env GO111MODULE=on GOPROXY=https://goproxy.io go

all:
	go build -o email *.go
