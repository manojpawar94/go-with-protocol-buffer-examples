export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
protoc -I=src/ --go_out=src/ src/simple_example/simple_message.proto
protoc -I=src/ --go_out=src/ src/enum_example/enum_message.proto
protoc -I=src/ --go_out=src/ src/comment_example/comments.proto