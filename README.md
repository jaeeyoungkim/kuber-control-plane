linux go install

1. apt update
2. apt upgrade
3. apt install golang-go

4. vim ~/.profile
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

5. source ~/.profile
6. go version
-> go version go1.20.3 linux/amd64

7. install go lib
go mod init example.com/hello
go get k8s.io/client-go/kubernetes
go get k8s.io/client-go/rest
go get k8s.io/client-go/discovery@v0.28.3
