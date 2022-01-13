https://www.cnblogs.com/chnmig/p/12188604.html

微服务相关
使用 GRPC 通讯的 Golang 微服务入门
举例写一个微服务,接收网址发送请求获取返回结果返回
``` 
这是 proto 文件的编译器
先安装Protobuf 编译器 protoc，下载地址：https://github.com/google/protobuf/releases 
我的是windows，将压缩包bin目录下的exe放到环境PATH目录中即可。

安装 golang 的proto工具包             
go get -u github.com/golang/protobuf/proto          
安装 goalng 的proto编译支持                
go get -u github.com/golang/protobuf/protoc-gen-go              
安装 GRPC 包               
go get -u google.golang.org/grpc
 ```


proto文件中追加
```
option go_package = "./";
 ```
生成 .bp.go 文件            
使用刚才下载的 protoc 工具将 proto 文件编译成 golang 可识别的文件
````
protoc --go_out=plugins=grpc:./ ./spider.proto

````
运行后会在当前目录下生成 spider.pb.go 文件
该文件是 server 和 client 的通信协议,业务代码不在这里,所以除非必须改,否则勿动

pb.go#
需要注意的是,在本个 demo 中,客户端与服务端都是 Golang,所以在客户端与服务端都公用一个 pb.go 模板文件(如果是不同的语言生成的pb是对应语言),
可以将 pb.go 文件放置在云上由双方引用,也可以生成两个副本放在两端项目中,本次就使用 COPY 两份的方式
由于 Golang 一个文件夹只有一个 package,而生成的 pb.go 文件 package 为创建 proto 的名字(示例为 spider), 

所以我们在项目内单独建立文件夹 spider将文件放入其中即可正常使用
-- 此处 公用导入  

运行#
需要先启动 server 端监听端口,再启动 client 端向端口发送请求
我们运行后可看到结果已经正常返回并打印
