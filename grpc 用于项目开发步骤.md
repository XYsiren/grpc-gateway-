1.安装protoc编译器，根据protobuf定义的文件生成grpc客户端与服务端存根

2.安装proto-gen-go插件，生成message

3.安装proto-gen-go-grpc插件，生成grpc服务的客户端与服务存根

4.protobuf定义message与service

5.通过protoc生成grpc服务代码

6.基于服务端存根实现grpc server端

7.将存根共享给客户端，以便客户端rpc调用server端