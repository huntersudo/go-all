
生成代码变化
```
protoc --go_out=plugins=grpc:. *.proto
```
在自动生成的go代码程序当中，每一个流模式对应的服务接口，都会自动生成对应的单独的client和server程序，以及对应的结构体实现
