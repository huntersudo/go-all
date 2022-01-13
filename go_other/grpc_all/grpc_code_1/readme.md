go get -a github.com/golang/protobuf/protoc-gen-go


生成golang的服务代码
```
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```
这个指令支持*.proto模糊匹配。如果有许多文件可以使用helloworld/*.proto 来作为PROTO_FILES


```
protoc --go_out=plugins=grpc:. *.proto
```