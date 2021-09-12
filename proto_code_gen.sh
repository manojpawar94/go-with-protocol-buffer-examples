export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
protoc -I=src/ --go_out=src/ src/sample/sample.proto
protoc -I=src/ --go_out=src/ src/enum_example/enum_message.proto