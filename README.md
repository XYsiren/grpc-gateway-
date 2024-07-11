# simple introduction

## protoc安装（windows）



protoc就是protobuf的编译器，它把proto文件编译成不同的语言



### 下载安装protoc编译器（protoc）



下载protobuf：[https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-win64.zip](https://link.zhihu.com/?target=https%3A//github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-win64.zip)



解压后，将目录中的 bin 目录的路径添加到系统环境变量，然后打开cmd输入`protoc`查看输出信息，此时则安装成功



### 安装protocbuf的go插件（protoc-gen-go）



由于protobuf并没直接支持go语言需要我们手动安装相关插件



protocol buffer编译器需要一个插件来根据提供的proto文件生成 Go 代码，Go1.16+要使用下面的命令安装插件：



```text
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest  // 目前最新版是v1.3.0
```



### 安装grpc（grpc）



```text
 go get -u -v google.golang.org/grpc@latest    // 目前最新版是v1.53.0
```



### 安装grpc的go插件（protoc-gen-go-grpc）



说明：在**`google.golang.org/protobuf`**中，**`protoc-gen-go`**纯粹用来生成pb序列化相关的文件，不再承载gRPC代码生成功能，所以如果要生成grpc相关的代码需要安装grpc-go相关的插件：**`protoc-gen-go-grpc`**



```text
 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest  // 目前最新版是v1.3.0
```



## protobuf语法





![img](https://pic4.zhimg.com/80/v2-74a16d5d128926144961eab36a57ebc3_720w.webp)



### protobuf语法



- 类型：类型不仅可以是标量类型（**`int`**、**`string`**等），也可以是复合类型（**`enum`**等），也可以是其他**`message`**
- 字段名：字段名比较推荐的是使用下划线/分隔名称
- 字段编号：一个message内每一个字段编号都必须唯一的，在编码后其实传递的是这个编号而不是字段名
- 字段规则：消息字段可以是以下字段之一

`singular`：格式正确的消息可以有零个或一个字段（但不能超过一个）。使用 proto3 语法时，如果未为给定字段指定其他字段规则，则这是默认字段规则

**`optional`**：与 **`singular`** 相同，不过可以检查该值是否明确设置

**`repeated`**：在格式正确的消息中，此字段类型可以重复零次或多次。系统会保留重复值的顺序

**`map`**：这是一个成对的键值对字段



- 保留字段：为了避免再次使用到已移除的字段可以设定保留字段。如果任何未来用户尝试使用这些字段标识符，编译器就会报错



### 简单语法



### proto文件基本语法



```text
 syntax = "proto3";              // 指定版本信息，不指定会报错
 package pb;                     // 后期生成go文件的包名
 // message为关键字，作用为定义一种消息类型
 message Person{
     string name = 1;   // 名字
     int32  age = 2 ;   // 年龄
 }
 
 enum test{
     int32 age = 0;
 }
```



protobuf消息的定义（或者称为描述）通常都写在一个以 .proto 结尾的文件中：



1. 第一行指定正在使用**`proto3`**语法：如果不这样做，协议缓冲区编译器将假定正在使用proto2（这也必须是文件的第一个非空的非注释行）
2. 第二行package指明当前是pb包（生成go文件之后和Go的包名保持一致）
3. message关键字定义一个Person消息体，类似于go语言中的结构体，是包含一系列类型数据的集合。
   许多标准的简单数据类型都可以作为字段类型，包括**`bool`**，**`int32`**， **`float`**，**`double`**，和**`string`**
   也可以使用其他message类型作为字段类型。



在message中有一个字符串类型的value成员，该成员编码时用1代替名字。在json中是通过成员的名字来绑定对应的数据，但是Protobuf编码却是通过成员的唯一编号来绑定对应的数据，因此Protobuf编码后数据的体积会比较小，能够快速传输，缺点是不利于阅读。



### message常见的数据类型与go中类型对比



![img](https://pic2.zhimg.com/80/v2-51e4b2811be2830be450d3dbfd28699d_720w.webp)



### protobuff语法进阶



### message嵌套



messsage除了能放简单数据类型外，还能存放另外的message类型：



```text
 syntax = "proto3";          // 指定版本信息，不指定会报错
 package pb;                 // 后期生成go文件的包名
 // message为关键字，作用为定义一种消息类型
 message Person{
     string name = 1;  // 名字
     int32  age = 2 ;  // 年龄
     // 定义一个message
     message PhoneNumber {
         string number = 1;
         int64 type = 2;
     }
     PhoneNumber phone = 3;
 }
```



message成员编号，可以不从1开始，但是不能重复，不能使用19000 - 19999



### repeated关键字



repeadted关键字类似与go中的切片，编译之后对应的也是go的切片，用法如下：



```text
syntax = "proto3";              // 指定版本信息，不指定会报错
 package pb;                     // 后期生成go文件的包名
 // message为关键字，作用为定义一种消息类型
 message Person{
     string name = 1;   // 名字
     int32  age = 2 ;   // 年龄
     // 定义一个message
     message PhoneNumber {
         string number = 1;
         int64 type = 2;
     }
     repeated PhoneNumber phone = 3;
 }
```



### 默认值



解析数据时，如果编码的消息不包含特定的单数元素，则解析对象对象中的相应字段将设置为该字段的默认值



不同类型的默认值不同，具体如下：



- 对于字符串，默认值为空字符串
- 对于字节，默认值为空字节
- 对于bools，默认值为false
- 对于数字类型，默认值为零
- 对于枚举，默认值是第一个定义的枚举值，该值必须为0。
- repeated字段默认值是空列表
- message字段的默认值为空对象



### enum关键字



在定义消息类型时，可能会希望其中一个字段有一个预定义的值列表



比如说，电话号码字段有个类型，这个类型可以是，home,work,mobile



我们可以通过enum在消息定义中添加每个可能值的常量来非常简单的执行此操作。示例如下：



```text
syntax = "proto3";              // 指定版本信息，不指定会报错
 package pb;                     // 后期生成go文件的包名
 // message为关键字，作用为定义一种消息类型
 message Person{
     string name = 1;   // 名字
     int32  age = 2 ;   // 年龄
     // 定义一个message
     message PhoneNumber {
         string number = 1;
         PhoneType type = 2;
     }
     
     repeated PhoneNumber phone = 3;
 }
 
 // enum为关键字，作用为定义一种枚举类型
 enum PhoneType {
     MOBILE = 0;
     HOME = 1;
     WORK = 2;
 }
```



> 如上，enum的第一个常量映射为0，每个枚举定义必须包含一个映射到零的常量作为其第一个元素。这是因为：

- 必须有一个零值，以便我们可以使用0作为数字默认值。
- 零值必须是第一个元素，以便与proto2语义兼容，其中第一个枚举值始终是默认值。





enum还可以为不同的枚举常量指定相同的值来定义别名。如果想要使用这个功能必须将**`allow_alias`**选项设置为true，负责编译器将报错。示例如下：



```text
syntax = "proto3";              // 指定版本信息，不指定会报错
 package pb;                     // 后期生成go文件的包名
 // message为关键字，作用为定义一种消息类型
 message Person{
     string name = 1;   // 名字
     int32  age = 2 ;   // 年龄
     // 定义一个message
     message PhoneNumber {
         string number = 1;
         PhoneType type = 2;
     }
     repeated PhoneNumber phone = 3;
 }
 
 // enum为关键字，作用为定义一种枚举类型
 enum PhoneType {
     // 如果不设置将报错
     option allow_alias = true;
     MOBILE = 0;
     HOME = 1;
     WORK = 2;
     Personal = 2;
 }
```



### oneof关键字



如果有一个包含许多字段的消息，并且最多只能同时设置其中的一个字段，则可以使用oneof功能，示例如下：



```text
 message Person{
     string name = 1; // 名字
     int32  age = 2 ; // 年龄
     //定义一个message
     message PhoneNumber {
         string number = 1;
         PhoneType type = 2;
     }
 
     repeated PhoneNumber phone = 3;
     oneof data{
         string school = 5;
         int32 score = 6;
     }
 }
```



### 定义RPC服务



如果需要将message与RPC一起使用，则可以在**`.proto`**文件中定义RPC服务接口，protobuf编译器将根据你选择的语言生成RPC接口代码。示例如下：



```text
 //定义RPC服务
 service HelloService {
     rpc Hello (Person)returns (Person);
 }
```



注意：默认protobuf编译期间，不编译服务，如果要想让其编译，需要使用gRPC



## protobuf编译



### 编译器调用



protobuf 编译是通过编译器 protoc 进行的，通过这个编译器，我们可以把 .proto 文件生成 go,Java,Python,C++, Ruby或者C# 代码



可以使用以下命令来通过 .proto 文件生成go代码（以及grpc代码）



```text
 // 将当前目录中的所有 .proto文件进行编译生成go代码
 protoc --go_out=./ --go_opt=paths=source_relative *.proto
```



protobuf 编译器会把 .proto 文件编译成 .pd.go 文件



### --go_out 参数



作用：指定go代码生成的基本路径



1. protocol buffer编译器会将生成的Go代码输出到命令行参数**`go_out`**指定的位置
2. **`go_out`**标志的参数是你希望编译器编写 Go 输出的目录
3. 编译器会为每个**`.proto`** 文件输入创建一个源文件
4. 输出文件的名称是通过将**`.proto`** 扩展名替换为**`.pb.go`** 而创建的



### --go_opt 参数



**`protoc-gen-go`**提供了**`--go_opt`**参数来为其指定参数，可以设置多个：



1. **`paths=import`**：生成的文件会按**`go_package`**路径来生成，当然是在**`--go_out`**目录

- 例如，**`go_out/$go_package/pb_filename.pb.go`**
- 如果未指定路径标志，这就是默认输出模式



1. **`paths=source_relative`**：输出文件与输入文件放在相同的目录中

- 例如，一个**`protos/buzz.proto`**输入文件会产生一个位于**`protos/buzz.pb.go`**的输出文件。



1. **`module=$PREFIX`**：输出文件放在以 Go 包的导入路径命名的目录中，但是从输出文件名中删除了指定的目录前缀。

- 例如，输入文件 **`pros/buzz.proto`**，其导入路径为 **`example.com/project/protos/fizz`** 并指定**`example.com/project`**为**`module`**前缀，结果会产生一个名为 **`pros/fizz/buzz.pb.go`** 的输出文件。
- 在module路径之外生成任何 Go 包都会导致错误，此模式对于将生成的文件直接输出到 Go 模块非常有用。

### --proto_path 参数



**`--proto_path=IMPORT_PATH`**



- IMPORT_PATH是 .proto 文件所在的路径，如果忽略则默认当前目录。
- 如果有多个目录则可以多次调用--proto_path，它们将会顺序的被访问并执行导入。



使用示例：



```text
 protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto
 // 编译器将从 `src` 目录中读取输入文件 `foo.proto` 和 `bar/baz.proto`，并将输出文件 `foo.pb.go` 和 `bar/baz.pb.go` 写入 `out` 目录。如果需要，编译器会自动创建嵌套的输出子目录，但不会创建输出目录本身
```



### 使用grpc的go插件



### 安装proto-gen-go-grpc



在**`google.golang.org/protobuf`**中，**`protoc-gen-go`**纯粹用来生成pb序列化相关的文件，不再承载gRPC代码生成功能。生成gRPC相关代码需要安装grpc-go相关的插件**`protoc-gen-go-grpc`**



```text
 // 安装protoc-gen-go-grpc
 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest  // 目前最新版是v1.3.0
```



生成grpc的go代码：



```text
 // 主要是--go_grpc_out参数会生成go代码
 protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  *.proto
```



### --go-grpc_out 参数



作用：指定grpc go代码生成的基本路径



命令会产生的go文件：



1. **`protoc-gen-go`**：包含所有类型的序列化和反序列化的go代码
2. **`protoc-gen-go-grpc`**：包含service中的用来给client调用的接口定义以及service中的用来给服务端实现的接口定义



### --go-grpc_opt 参数



和**`protoc-gen-go`**类似，**`protoc-gen-go-grpc`**提供 **`--go-grpc_opt`** 来指定参数，并可以设置多个

### `http://github.com/golang/protobuf 和 `

### `http://google.golang.org/protobuf`

### `http://github.com/golang/protobuf`



1. **`github.com/golang/protobuf`** 现在已经废弃
2. 它可以同时生成pb和gRPC相关代码的



用法：

```text
 // 它在--go_out加了plugin关键字，paths参数有两个选项，分别是 import 和 source_relative
 --go_out=plugins=grpc,paths=import:.  *.proto
```

### `http://google.golang.org/protobuf`



1. 它**`github.com/golang/protobuf`**的升级版本，**`v1.4.0`**之后**`github.com/golang/protobuf`**仅是**`google.golang.org/protobuf`**的包装
2. 它纯粹用来生成pb序列化相关的文件，不再承载gRPC代码生成功能，生成gRPC相关代码需要安装grpc-go相关的插件**`protoc-gen-go-grpc`**



用法：



```text
 // 它额外添加了参数--go-grpc_out以调用protoc-gen-go-grpc插件生成grpc代码
  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  *.proto

```

