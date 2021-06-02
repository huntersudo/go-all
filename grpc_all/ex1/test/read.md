https://www.cnblogs.com/chnmig/p/12188604.html


option go_package = "./";

````
protoc --go_out=plugins=grpc:./ ./spider.proto
protoc --go_out=plugins=grpc:./ ./test.proto
  
````


https://www.jianshu.com/p/20ed82218163

准备工作
先安装Protobuf 编译器 protoc，下载地址：https://github.com/google/protobuf/releases
我的是windows，将压缩包bin目录下的exe放到环境PATH目录中即可。

然后获取插件支持库
``` 
 // gRPC运行时接口编解码支持库
  go get -u github.com/golang/protobuf/proto
  // 从 Proto文件(gRPC接口描述文件) 生成 go文件 的编译器插件
  go get -u github.com/golang/protobuf/protoc-gen-go
  
```

获取go的gRPC包(网络问题可参阅https://www.jianshu.com/p/6392cb9dc38f)
``` 
go get google.golang.org/grpc
```

proto文件语法详解参阅：https://blog.csdn.net/u014308482/article/details/52958148


proto文件中追加
```
option go_package = "./";
 ```

然后将proto文件编译为go文件
cd workspace/ws1/src/github.com/go-all/grpc_all/ex1/    
``` 
// protoc --go_out=plugins=grpc:{输出目录}  {proto文件}
protoc --go_out=plugins=grpc:./ ./test.proto
```


