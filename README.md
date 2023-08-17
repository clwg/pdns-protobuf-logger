
To complile the protocol buffers you will need to install the compiler ```apt install -y protobuf-compiler protoc-gen-go```


## Generating protobufs

First clone the pdns repo
```git clone https://github.com/PowerDNS/pdns.git```

In order to compile without complaining add the following under the syntax = "proto2"
```option go_package = "github.com/clwg/protos/pdnsmsg";```


protoc -I=./pdns/pdns/ --go_out=. ./pdns/pdns/dnsmessage.proto



protoc --proto_path=src \
  --go_opt=Mprotos/buzz.proto=example.com/project/protos/fizz \
  --go_opt=Mprotos/bar.proto=example.com/project/protos/foo \
  ./pdns/pdns/dnsmessage.proto


