![Bubbler-BANNER](.assets/Bubbler-BANNER.png)

# Bubbler

[English](README.md) | 简体中文

Bubbler 是一个专为物联网设备优化的协议生成器。它编译 `.bb` 协议文件，并以指定的目标语言生成输出。

Bubbler 的协议功能强大，而且可以是非字节对齐的，这对于资源有限的物联网设备非常有用。下面会详细解释。

另外，你可能需要 `.bb` 文件的语法高亮，可以查看 [bubbler-vscode](https://github.com/xaxys/bubbler-vscode)，或者从 [VSCode Marketplace](https://marketplace.visualstudio.com/items?itemName=xaxys.Bubbler) 安装。

警告：Bubbler 仍在开发中，尚未准备好用于生产。

## 安装

```sh
git clone https://github.com/xaxys/bubbler.git
cd bubbler
go build
```

## 使用方法

```sh
bubbler [options] <input file>
```

### 选项

- `-t <target>`: 目标语言
- `-o <output>`: 输出路径
- `-rmpath`: 移除路径前缀（生成文件时移除输出文件路径的路径前缀）
  通常用于生成 Go 目标。例如，如果 `.bb` 文件的 `go_package` 选项设置为 `github.com/xaxys/bubbler/proto/rpc`，则文件将在 `output/github.com/xaxys/bubbler/proto/rpc` 目录中生成。如果要移除路径前缀 `github.com/xaxys/bubbler/proto/rpc`，可以将此选项设置为 `github.com/xaxys/bubbler/proto`。然后生成的文件将在 `output/rpc` 目录中生成。
- `-inner`: 生成内部类（嵌套结构体）
- `-single`: 生成单个文件（将所有定义合并到一个文件中，而不是每个源文件生成一个文件）
- `-minimal`: 生成最小代码（通常不包含默认的getter/setter方法）
- `-decnum`: 强制生成十进制格式的常量值（将 `0xFF` 翻译为 `255`, `0b1111` 翻译为 `15` 等）
- `-memcpy`: 启用字段的内存复制（解码时复制 `string` 和 `bytes` 字段的内容，而不是直接引用原始解码缓冲区）
- `-signext <method>`: 用于整数字段的符号扩展方法（选项: `shift`, `arith`）

### 示例

```sh
bubbler -t c -minimal -o output/ example.bb
bubbler -t c -single -o gen.hpp example.bb
bubbler -t py -decnum -signext=arith -o output example.bb
bubbler -t go -rmpath github.com/xaxys/bubbler/proto -o output example.bb
```

### 目标语言

运行 `bubbler` 命令查看支持的目标语言列表。

```text
Targets:
  c
  cpp
  csharp [cs]
  commonjs [cjs]
  go
  java
  python [py]

```

当选择目标语言时，可以使用 `[]` 中的别名。例如，`python` 可以缩写为 `py`。

- `dump`：输出 `.bb` 文件的解析树（中间表示）。

- `c`：C 语言，为每个 `.bb` 文件输出一个 `.bb.h` 文件和一个 `.bb.c` 文件。
  - 使用 `-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由`-o`选项确定。
  - 使用 `-minimal`：不为字段生成默认的getter/setter方法函数。
  - 使用 `-memcpy`：将使用 `malloc` 为 `string` 和 `bytes` 字段在堆上分配内存，并从原始缓冲区复制内容。
  - 不使用 `-memcpy`：`string` 和 `bytes` 字段的指针将直接引用原始解码缓冲区。提供零拷贝和零堆分配。

- `cpp`：C++ 语言，为每个 `.bb` 文件输出一个 `.bb.hpp` 文件和一个 `.bb.cpp` 文件。
  - 使用 `-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由 `-o` 选项确定。生成的文件夹结构不会受到 `cpp_namespace` 选项的影响。
  - 使用 `-minimal`：不为字段生成默认的getter/setter方法函数。
  - 使用 `-memcpy`：使用 `std::shared_ptr<uint8_t[]>` 为 `bytes` 字段在堆上分配内存，并从原始缓冲区复制内容。`string` 字段将始终使用 `std::string` 并在每次复制时复制。
  - 不使用 `-memcpy`：使用 `std::shared_ptr<uint8_t[]>` 和空删除器引用 `bytes` 字段的原始缓冲区。`string` 字段将始终使用 `std::string` 并在每次复制时复制。

- `csharp`：C# 语言，为每个 `.bb` 文件输出一个 `.bb.cs` 文件。
  - 使用 `-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由 `-o` 选项确定。生成的文件夹结构不会受到 `csharp_namespace` 选项的影响。
  - 使用 `-memcpy`：使用 `byte[]` 作为 `bytes` 字段的类型。编码和解码方法将仅兼容 `byte[]` 类型参数。旧版 .NET Framework 应使用此选项。
  - 不使用 `-memcpy`：使用 `Memory<byte>` 作为 `bytes` 字段的类型。编码和解码方法将兼容 `byte[]`、`Memory<byte>` 和 `Span<byte>`（仅编码）类型参数。此情况下需要 `System.Memory` 包。

- `commonjs`：CommonJS模块，为每个 `.bb` 文件输出一个 `.bb.js` 文件。（请注意，`int64` 和 `uint64` 字段使用了 `BigInt`，在某些环境中可能不支持）
  - 使用 `-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由 `-o` 选项确定。
  - 强制启用：`-memcpy`。

- `go`：Go 语言，为每个 `.bb` 文件输出一个 `.bb.go` 文件。生成的文件夹结构将受到 `go_package` 选项的影响。例如，`github.com/xaxys/bubbler` 将在 `github.com/xaxys/bubbler` 目录中生成。
  - 使用 `-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由 `-o` 选项确定。包名由输入 `.bb` 文件的包名声明确定。
  - 使用 `-memcpy`：解码时复制 `bytes` 字段。`string` 字段将始终被复制。
  - 不使用 `-memcpy`：将原始缓冲区的切片分配给 `bytes` 字段。`string` 字段将始终被复制。

- `java`：Java 语言，为每个 `.bb` 文件中定义的每个数据结构生成一个 `.java` 文件。生成的文件夹结构将受到 `java_package` 选项的影响。例如，`com.example.rovlink` 将在 `com/example/rovlink` 目录中生成。
  - 强制启用：`-memcpy`。

- `python`：Python 语言，为每个 `.bb` 文件输出一个 `_bb.py` 文件。
  - 使用`-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由 `-o` 选项确定。
  - 强制启用：`-memcpy`。

## 协议语法

Bubbler 使用简洁的语法来定义数据结构和枚举类型。

在 [example](example/) 目录中查看示例。

### 包名声明

使用 `package` 关键字来定义包名。例如：

使用 `import` 关键字导入其他 Bubbler 协议文件。例如：

```protobuf
package com.example.rovlink;
```

包名用于生成输出文件名。例如，如果包名为 `com.example.rovlink`，则输出文件名为 `rovlink.xxx`，并放置在 `${Output Path}/com/example/` 目录中。

在 `.bb` 文件中只允许有一个包名声明，并且包名不能在全局范围内重复。

### 选项声明

使用 `option` 关键字来定义选项。例如：

```protobuf
option omit_empty = true;
option go_package = "example.com/rovlink";
option cpp_namespace = "com::example::rovlink";
option csharp_namespace = "Example.Rovlink";
option java_package = "com.example.rovlink";
```

在 `.bb` 文件中，选项语句不能重复。

如果选项未知，将会产生编译器警告。

#### 支持的选项

##### `omit_empty`

如果将 `omit_empty` 设置为 `true`，不含有任何类型定义的 `.bb` 文件将不会生成任何文件。

```protobuf
package all;

option omit_empty = true;

import "rovlink.bb";
import "control.bb";
import "excomponent.bb";
import "excontrol.bb";
import "exdata.bb";
import "host.bb";
import "mode.bb";
import "sensor.bb";
```

在这个例子中，`omit_empty` 选项被设置为 `true`，这个 `.bb` 文件将不会生成名为 `all.xxx` 文件。

您可以使用这个选项一次生成多个 `.bb` 文件，而无需编写外部脚本，来运行多次 `bubbler` 命令以生成多个文件。

##### `go_package`

如果设置了 `go_package`，生成的代码将在生成的 Go 代码中使用指定的包名。

##### `cpp_namespace`

如果设置了 `cpp_namespace`，生成的代码将在生成的 C++ 代码中使用指定的命名空间。

##### `csharp_namespace`

如果设置了 `csharp_namespace`，生成的代码将在生成的 C# 代码中使用指定的命名空间，但是文件夹结构不会受到影响。

##### `java_package`

如果设置了 `java_package`，生成的代码将在生成的 Java 代码中使用指定的包名，并且生成的文件夹结构将会根据该包名生成。

### 导入语句

使用 `import` 关键字导入其他 Bubbler 协议文件。例如：

```python
import "control.bb";
import "a.bb";
```

### 枚举类型

使用 `enum` 关键字定义枚举类型。枚举类型的定义包括枚举名称和枚举值。例如：

```c
enum FrameType[1] {
    SENSOR_PRESS = 0x00,
    SENSOR_HUMID = 0x01,
    CURRENT_SERVO_A = 0xA0,
    CURRENT_SERVO_B = 0xA1,
};
```

在这个例子中，`FrameType` 是一个枚举类型，它有四个枚举值：`SENSOR_PRESS`、`SENSOR_HUMID`、`CURRENT_SERVO_A` 和 `CURRENT_SERVO_B`。

枚举值不能为负数（暂定），不填写值时，枚举值的默认值为前一个枚举值加 1。

枚举类型名后面的方括号中的数字表示枚举类型的宽度，例如 `[1]` 表示 1 字节。也可以使用 `#` 符号表示字节和比特，例如 `#1` 表示 1 比特，`#2` 表示 2 比特。也可以混合使用，例如 `1#4` 表示 1 字节 4 比特，即 12 比特。

推荐使用 **PascalCase** 作为枚举类型名称。但是只有首字母大写是强制要求。

推荐使用 **ALLCAP_CASE** 作为枚举值。但是只有首字母大写是强制要求。

### 数据结构

使用 `struct` 关键字定义数据结构。数据结构的定义包括结构名称和一系列的字段。例如：

```c
struct Frame[20] {
    FrameType opcode;
    struct SomeEmbed[1] {
        bool valid[#1];
        bool error[#1];
        uint8 source[#3];
        uint8 target[#3];
    };
    uint8<18> payload;
};
```

在这个例子中，`Frame` 是一个数据结构，它有三个字段：`opcode`、`SomeEmbed` 和 `payload`。`opcode` 的类型是 `FrameType`，`SomeEmbed` 是一个匿名内嵌的数据结构，`payload` 的类型是 `uint8`。

请注意，Bubbler中并无作用域的概念（为了适应 C 语言），所以 `Frame` 和 `SomeEmbed` 作为数据结构名称，在全局都不允许重名，哪怕 `SomeEmbed` 是一个匿名内嵌的数据结构。

推荐使用 **PascalCase** 作为数据结构名称。但是只有首字母大写是强制要求。

推荐使用 **snake_case** 作为字段名称。但是只有首字母小写是强制要求。

### 字段类型

Bubbler 协议支持四种字段类型：普通字段、匿名内嵌字段、常量字段和空字段。

- 普通字段：由类型名、字段名和字段宽度（可选）构成。
- 匿名内嵌字段：一个匿名的字段，可以是struct定义或已定义的struct名称，其内部子字段会被提升并展开到父结构体中。
- 常量字段：一个固定值的字段，其值在定义时就已经确定，不能被修改。字段名可选，如果有字段名，会生成对应字段。编码时，常量字段的值会被忽略。解码时，常量字段的值会被检查，如果不匹配，会报错。
- 空字段：一个没有名字和类型的字段，只有宽度，用于占位。

#### 普通字段

普通字段由类型名、字段名和字段宽度构成。例如：

```c
struct Frame {
    RovlinkFrameType opcode;
};
```

在这个例子中，`opcode` 是一个普通字段，其类型为 `RovlinkFrameType`。

字段宽度可选，如果不填写宽度，则字段宽度为类型的宽度。

字段宽度可以小于类型的宽度，例如：

```c
struct Frame[20] {
    int64 my_int48[6];
};
```

在这个例子中，`my_int48` 是一个 6 字节的字段，其类型为 `int64`，但是它的宽度为 6 字节，因此它编码时只会占用 6 字节的空间。

但是，对于`struct`类型的字段，字段宽度必须等于类型的宽度

### 匿名内嵌字段

匿名内嵌字段是一个没有名字的数据结构，它可以包含多个子字段。例如：

```c
struct Frame {
    int64 my_int48[6];
    struct SomeEmbed[1] {
        bool valid[#1];
        bool error[#1];
        uint8 source[#3];
        uint8 target[#3];
    };
};
```

在这个例子中，`SomeEmbed` 是一个匿名内嵌字段，它包含了四个子字段：`valid`、`error`、`source` 和 `target`。

匿名内嵌字段的子字段会被提升并展开到父结构体中。生成的结构如下：

```c
struct Frame {
    int64_t my_int48;
    bool valid;
    bool error;
    uint8_t source;
    uint8_t target;
};
```

匿名内嵌字段也可以是一个已定义的数据结构，例如：

```c
struct AnotherTest {
    int8<2> arr;
}

struct Frame {
    int64 my_int48[6];
    AnotherTest;
    uint8<18> payload;
};
```

这样，生成的结构如下：

```c
struct Frame {
    int64_t my_int48;
    int8_t arr[2];
    uint8_t payload;
};
```

### 常量字段

常量字段是一个固定值的字段，它的值在定义时就已经确定，不能被修改。例如：

```c
struct Frame {
    uint8 FRAME_HEADER = 0xAA;
};
```

在这个例子中，`FRAME_HEADER` 是一个常量字段，其值为 `0xAA`。

或者你可以使用之前定义的枚举类型的枚举值作为常量值：

```c
enum FrameType[1] {
    FRAME_KEEPALIVE = 0x00,
    FRAME_DATA = 0x01,
};

struct Frame {
    FrameType opcode = FRAME_DATA;
    bytes data;
};
```

常量字段的值在编码时会被忽略，解码时会被检查，如果不匹配，会报错。

### 空字段

空字段是一个没有名字和类型的字段，它只有宽度。空字段通常用于填充或对齐数据结构。例如：

```c
struct Frame {
    void [#2];
};
```

在这个例子中，`void [#2]` 是一个空字段，它占用了 2 比特的空间。

### 字段选项

字段选项用于指定字段的额外属性。例如，可以使用 `order` 选项指定数组的字节顺序：

```c
struct AnotherTest {
    int8<2> arr [order = "big"];
}
```

在这个例子中，`arr` 字段的字节顺序被设置为大端序。

> 小贴士：大小端序的设置对于浮点类型同样有效，但是目前浮点解读时一律按照小端序解读，即最高位存放符号位，然后是指数位，最后是尾数位。

### 自定义 getter/setter

可以为字段定义自定义的 getter 和 setter 方法，用于在读取或写入字段值时执行特定的操作。例如：

```bubbler
struct SensorTemperatureData {
    uint16 temperature[2] {
        get temperature_display(float64): value / 10 - 40;
        set temperature_display(float64): value == 0 ? 0 : (value + 40) * 10;
        set another_custom_setter(uint8): value == 0 ? 0 : (value + 40) * 10;
    };
}
```

在这个例子中，`temperature` 字段有一个自定义的 getter 方法和两个自定义的 setter 方法。

自定义getter名为 `temperature_display`, 返回`float64` 类型，并根据 `value / 10 - 40` 计算结果返回。其中 `value` 被填充为字段的值，是uint16类型。

自定义setter名为 `temperature_display`, 接收`float64` 类型的参数，并根据 `value == 0 ? 0 : (value + 40) * 10` 计算结果并以此设置字段的值。其中 `value` 被填充为参数的值，是 `float64` 类型。

自定义setter名为 `another_custom_setter`，`uint8`是参数类型。并根据 `value == 0 ? 0 : (value + 40) * 10` 计算结果并以此设置字段的值。其中 `value` 被填充为参数的值，是 `uint8` 类型。

请注意，自定义的 getter 和 setter 方法的名不可以与任何字段名相同，并且同名的 getter 和 setter 方法必须返回和接收相同的类型。

推荐使用 **snake_case** 作为自定义 getter/setter 方法名称。但是只有首字母小写是强制要求。

## 贡献

欢迎为 Bubbler 做出贡献。

## 许可证

MIT 许可证

## 相关仓库

- [CoralReefPlayer](https://github.com/DawningW/CoralReefPlayer) - 珊瑚礁播放器，一款现代化跨平台低延迟流媒体播放器库。
- [OpenFinNAV](https://github.com/redlightASl/OpenFinNAV) - 鳍航，针对水下机器人（ROV/AUV）设计的飞控固件库。
