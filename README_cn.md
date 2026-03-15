![Bubbler-BANNER](.assets/Bubbler-BANNER.png)

# Bubbler

[English](README.md) | 简体中文

Bubbler 是一个专为物联网设备优化的协议生成器。它编译 `.bb` 协议文件，并以指定的目标语言生成输出。

Bubbler 的协议功能强大，而且可以是非字节对齐的，这对于资源有限的物联网设备非常有用。下面会详细解释。

另外，你可能需要 `.bb` 文件的语法高亮，可以查看 [bubbler-vscode](https://github.com/xaxys/bubbler-vscode)，或者从 [VSCode Marketplace](https://marketplace.visualstudio.com/items?itemName=xaxys.Bubbler) 安装。

警告：Bubbler 仍在开发中，尚未准备好用于生产。

## 安装

可以从 [releases page](https://github.com/xaxys/bubbler/releases/) 下载最新版本和预编译的二进制文件。

或者者你也可以从源代码构建：

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
- `-relpath`: 生成相对路径导入（在生成文件引用其他生成文件时，强制使用 `./` 或 `../` 之类的相对路径，而不是依赖默认的绝路路径逻辑或包路径规则）。
- `-inner`: 生成内部类（嵌套结构体）
- `-single`: 生成单个文件（将所有定义合并到一个文件中，而不是每个源文件生成一个文件）
- `-minimal`: 生成最小代码（通常不包含默认的getter/setter方法）
- `-decnum`: 强制生成十进制格式的常量值（将 `0xFF` 翻译为 `255`, `0b1111` 翻译为 `15` 等）
- `-memcpy`: 启用字段的内存复制（解码时复制 `string` 和 `bytes` 字段的内容，而不是直接引用原始解码缓冲区）
- `-signext <method>`: 用于整数字段的符号扩展方法（选项: `shift`, `arith`）
- `-compat`: 生成兼容性代码（在 CommonJS 目标中使用 `Array` 代替 `Uint8Array` 等 typed array 作为缓冲區和 `bytes` 字段的类型。默认使用 `Uint8Array` 以提高性能）

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
  esmodule [javascript, js, mjs, esm]
  go
  java
  python [py]

```

当选择目标语言时，可以使用 `[]` 中的别名。例如，`python` 可以缩写为 `py`。

- `c`：C 语言，为每个 `.bb` 文件输出一个 `.bb.h` 文件和一个 `.bb.c` 文件。
  - 使用 `-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由`-o`选项确定。
  - 使用 `-minimal`：不为字段生成默认的getter/setter方法函数。
  - 使用 `-memcpy`：将使用 `malloc` 为 `string` 和 `bytes` 字段在堆上分配内存，并从原始缓冲区复制内容。
  - 不使用 `-memcpy`：`string` 和 `bytes` 字段的指针将直接引用原始解码缓冲区。提供零拷贝和零堆分配。
  - 使用 `-relpath`：生成相对路径导入（如 `#include "./foo_bb.h"`或`#include "../foo_bb.h"`）。

- `cpp`：C++ 语言，为每个 `.bb` 文件输出一个 `.bb.hpp` 文件和一个 `.bb.cpp` 文件。
  - 使用 `-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由 `-o` 选项确定。生成的文件夹结构不会受到 `cpp_namespace` 选项的影响。
  - 使用 `-minimal`：不为字段生成默认的getter/setter方法函数。
  - 使用 `-memcpy`：使用 `std::shared_ptr<uint8_t[]>` 为 `bytes` 字段在堆上分配内存，并从原始缓冲区复制内容。`string` 字段将始终使用 `std::string` 并在每次复制时复制。
  - 不使用 `-memcpy`：使用 `std::shared_ptr<uint8_t[]>` 和空删除器引用 `bytes` 字段的原始缓冲区。`string` 字段将始终使用 `std::string` 并在每次复制时复制。
  - 使用 `-relpath`：生成相对路径导入（如 `#include "./foo_bb.hpp"`或`#include "../foo_bb.hpp"`）。

- `csharp`：C# 语言，为每个 `.bb` 文件输出一个 `.bb.cs` 文件。
  - 使用 `-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由 `-o` 选项确定。生成的文件夹结构不会受到 `csharp_namespace` 选项的影响。
  - 使用 `-memcpy`：使用 `byte[]` 作为 `bytes` 字段的类型。编码和解码方法将仅兼容 `byte[]` 类型参数。旧版 .NET Framework 应使用此选项。
  - 不使用 `-memcpy`：使用 `Memory<byte>` 作为 `bytes` 字段的类型。编码和解码方法将兼容 `byte[]`、`Memory<byte>` 和 `Span<byte>`（仅编码）类型参数。此情况下需要 `System.Memory` 包。

- `commonjs`：CommonJS模块，为每个 `.bb` 文件输出一个 `.bb.js` 文件。（请注意，`int64` 和 `uint64` 字段使用了 `BigInt`，在某些环境中可能不支持）
  - 使用 `-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由 `-o` 选项确定。
  - 不使用 `-compat`：使用 `Uint8Array` 作为编码缓冲区和 `bytes` 字段的类型，以提高性能（默认）。
  - 使用 `-compat`：使用 `Array` 代替 `Uint8Array` 作为编码缓冲区和 `bytes` 字段的类型，最大化对旧环境的兑容性。
  - 强制启用：`-memcpy`。
  - 强制启用：`-relpath`：生成相对路径引用（如 `require("./foo_bb")`或`require("../foo_bb")`）。

- `esmodule`：ES6模块，为每个 `.bb` 文件输出一个 `.bb.js` 文件。使用原生 ES6 `import`/`export` 语法，适用于现代浏览器和 Node.js ESM。（请注意，`int64` 和 `uint64` 字段使用了 `BigInt`，在某些环境中可能不支持）
  - 使用 `-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由 `-o` 选项确定。
  - 不使用 `-compat`：使用 `Uint8Array` 作为编码缓冲区和 `bytes` 字段的类型，以提高性能（默认）。
  - 使用 `-compat`：使用 `Array` 代替 `Uint8Array` 作为编码缓冲区和 `bytes` 字段的类型，最大化对旧环境的兼容性。
  - 强制启用：`-memcpy`。
  - 强制启用：`-relpath`：生成相对路径导入（如 `import X from "./foo_bb.mjs"`或`import X from "../foo_bb.mjs"`）。

- `go`：Go 语言，为每个 `.bb` 文件输出一个 `.bb.go` 文件。生成的文件夹结构将受到 `go_package` 选项的影响。例如，`github.com/xaxys/bubbler` 将在 `github.com/xaxys/bubbler` 目录中生成。
  - 使用 `-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由 `-o` 选项确定。包名由输入 `.bb` 文件的包名声明确定。
  - 使用 `-memcpy`：解码时复制 `bytes` 字段。`string` 字段将始终被复制。
  - 不使用 `-memcpy`：将原始缓冲区的切片分配给 `bytes` 字段。`string` 字段将始终被复制。
  - 使用 `-relpath`：生成相对路径导入（如 `import "./foo_bb"` 或 `import "../subpkg/foo_bb"`）。对于设置了 `go_package` 的情况，生成的相对路径将基于 `go_package` 定义的包路径进行计算。
  **注意！**生成相对路径时不会考虑 `-rmpath` 选项。即假如 A 的包路径为 `github.com/xaxys/a`，B 的包路径为 `gitlab.com/user/b`，而 `-rmpath=github.com/xaxys`，则 A 的生成路径为 `a`，B 的生成路径为 `gitlab.com/user/b`，但 B 中导入 A 的路径仍然会被计算为 `../../github.com/xaxys/a`，而不是 `../../a`。

- `java`：Java 语言，为每个 `.bb` 文件中定义的每个数据结构生成一个 `.java` 文件。生成的文件夹结构将受到 `java_package` 选项的影响。例如，`com.example.rovlink` 将在 `com/example/rovlink` 目录中生成。
  - 强制启用：`-memcpy`。

- `python`：Python 语言，为每个 `.bb` 文件输出一个 `_bb.py` 文件。
  - 使用`-single`：输出单个文件，其中包含所有 `.bb` 文件的所有定义。输出文件名（包括扩展名）由 `-o` 选项确定。
  - 强制启用：`-memcpy`。
  - 使用 `-relpath`：生成相对路径引用（如 `from .foo_bb import *`或`from ..foo_bb import *`）。

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

如果设置了 `go_package`，生成的代码将在生成的 Go 代码中使用指定的包名，并且生成的文件夹结构将会根据该包名生成。

##### `cpp_namespace`

如果设置了 `cpp_namespace`，生成的代码将在生成的 C++ 代码中使用指定的命名空间，但是文件夹结构不会受到影响。

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

### 变长类型

Bubbler 支持 `string` 和 `bytes` 类型用于变长数据。

#### String

`string` 类型用于文本字符串。在二进制形式中，字符串字段的文本编码为 UTF-8。由于字符串不应包含任何 EOF 字符，因此数据流中字符串的结尾由 `\0` 定义。所以字符串字段的大小是 `str.utf8_length + 1`。

#### Bytes

`bytes` 类型用于以任何形式存储二进制数据。在二进制形式中，bytes 数据存储为 `长度` + `数据` 两部分。

`长度` 部分表示 bytes 字段中包含的字节数。它由几个大小单元组成，每个单元占用 1 个字节，如下所示：

```text
  0   1   2   3   4   5   6   7
+---+---+---+---+---+---+---+---+
| C |       7-bit Size          |
+---+---+---+---+---+---+---+---+
  C = Continue Flag (继续标志)
```

Continue Flag 指示此单元之后是否还有另一个大小单元。如果 Continue Flag 为 0，则它是最后一个大小单元。

7-bit Size 存储小端序的长度。

因此，1 个大小单元可以表示 0-127 字节的数据；2 个大小单元可以表示 128-16383 字节的数据；...

`数据` 部分包含连续形式的原始二进制数据。

**注意**: 变长类型会使结构体大小变为动态的。

### 枚举值作为常量

您可以使用在先前枚举类型中定义的枚举值作为字段或其他枚举值的常量值。

```protobuf
enum FrameType[1] {
    FRAME_DATA = 0x01,
};

struct DataFrame {
    FrameType opcode = FRAME_DATA;
    bytes data;
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

这里还定义了 `Frame` 结构体的宽度为 20 字节。结构体宽度是可选的，可以写作 `struct Frame {`。

如果不填写宽度，则结构体宽度为所有字段宽度之和。但是如果填写了宽度，则结构体宽度必须与所有字段宽度之和完全相等，否则会报错。如果结构体内包含动态大小的字段（如 `string` 或 `bytes`），或者其他动态大小的结构体，则结构体宽度必须被省略。

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
>
> 什么意思呢？例如，对于 `uint32` 类型的字段，假设值是 `0x00123456`，我们可以设置其字段宽度为 3 字节，这样它就变成了 `"uint24"`。然后我们可以按大端序编码这个 `"uint24"` 为 `12 34 56` ，也可以按小端序编码这个 `"uint24"` 为 `56 34 12`。
>
> 但是对于 `float32` 类型的字段，假设符号位为 `1`，我们如果将其字段宽度设为 31 比特，那编码时它的符号位就被扔掉啦。这样它就变成 `"unsigned float31"`。然后虽然我们也可以选择按大端序或小端序编码，但是符号位就不管怎么编码都被丢掉了，所以解码时它的符号位永远是 `0`。

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

## 生成代码 API

每种语言的生成代码都提供了一致的编码和解码 API。

### C

```c
// 编码结构体到缓冲区。返回写入的字节数。
uint64_t <StructName>_encode(struct <StructName>* ptr, void* data);

// 从缓冲区解码结构体。返回读取的字节数，错误时返回 -1。
int64_t <StructName>_decode(const void* data, struct <StructName>* ptr);

// 估算编码后的字节大小。
uint64_t <StructName>_encode_size(struct <StructName>* ptr);

// 返回编码帧大小（>0）；若数据不足，返回负的最小所需长度（<0）。
int64_t <StructName>_decode_size(const void* data, uint64_t size);
```

### C++

```cpp
// 编码到缓冲区。返回写入的字节数。
// 不使用 -compat（C++20）：
uint64_t encode(::std::span<uint8_t> buf) const;
// 不使用 -compat：保留兼容代理（已弃用）
[[deprecated("Use encode(::std::span<uint8_t>) instead")]]
uint64_t encode(void* data) const;
// 使用 -compat：
uint64_t encode(void* data) const;

// 从缓冲区解码。返回读取的字节数，错误时返回 -1。
// 不使用 -compat（C++20）：
int64_t decode(::std::span<const uint8_t> data);
// 不使用 -compat：保留兼容代理（已弃用）
[[deprecated("Use decode(::std::span<const uint8_t>) instead")]]
int64_t decode(const void* data);
// 使用 -compat：
int64_t decode(const void* data);

// 估算编码后的字节大小。
uint64_t encode_size() const;

// 返回编码帧大小（>0）；若数据不足，返回负的最小所需长度（<0）。
// 不使用 -compat（C++20）：
static int64_t decode_size(::std::span<const uint8_t> buf);
// 不使用 -compat：保留兼容代理（已弃用）
[[deprecated("Use decode_size(::std::span<const uint8_t>) instead")]]
static int64_t decode_size(const void* data, uint64_t size);
// 使用 -compat：
static int64_t decode_size(const void* data, uint64_t size);
```

### Go

```go
// 编码并返回新分配的字节数组。
func (s StructName) Encode() []byte

// 编码到缓冲区。返回写入的字节数。
func (s StructName) EncodeTo(data []byte) int

// 从缓冲区解码。返回读取的字节数，错误时返回 -1。
func (s *StructName) Decode(data []byte) int

// 估算编码后的字节大小。
func (s StructName) EncodeSize() int

// 返回编码帧大小（>0）；若数据不足，返回负的最小所需长度（<0）。
func (s *StructName) DecodeSize(data []byte) int
```

### Java

```java
// 编码并返回新分配的字节数组。
public byte[] encode();

// 编码到缓冲区。返回写入的字节数。
public int encode(byte[] data, int start);

// 从缓冲区解码。返回读取的字节数，错误时返回 -1。
public int decode(byte[] data);
public int decode(byte[] data, int start);

// 估算编码后的字节大小。
public int encodeSize();

// 返回编码帧大小（>0）；若数据不足，返回负的最小所需长度（<0）。
public int decodeSize(byte[] data);
public int decodeSize(byte[] data, int start);
```

### Python

```python
# 编码为新字节数组，或编码到传入缓冲区
def encode(self, buffer: Union[None, bytearray, memoryview] = None) -> Union[bytearray, int]:

# 从缓冲区解码
# 成功返回 (True, 已读取字节数)，失败返回 (False, -1)
def decode(self, data: Union[bytes, bytearray, memoryview]) -> Tuple[bool, int]:

# 估算编码后的字节大小
def encode_size(self) -> int:

# 返回编码帧大小（>0）；若数据不足，返回负的最小所需长度（<0）
def decode_size(self, data: Union[bytes, bytearray, memoryview]) -> int:
```

### C\#

```csharp
// 编码并返回新分配的字节数组。
public byte[] Encode();

// 编码到缓冲区。返回写入的字节数。
public int Encode(byte[] data, int start);

// -memcpy=false 时额外提供：
public int Encode(Memory<byte> data);
public int Encode(Span<byte> data);

// 从缓冲区解码。返回读取的字节数，错误时返回 -1。
public int Decode(byte[] data);
public int Decode(byte[] data, int start);

// -memcpy=false 时额外提供：
public int Decode(Memory<byte> memoryData);

// 估算编码后的字节大小。
public int EncodeSize();

// 返回编码帧大小（>0）；若数据不足，返回负的最小所需长度（<0）。
public int DecodeSize(byte[] data);
public int DecodeSize(byte[] data, int start);

// -memcpy=false 时额外提供：
public int DecodeSize(Memory<byte> memoryData);
```

### CommonJS

```javascript
// 静态方法
StructName.encode(obj, buffer, start);
StructName.decode(obj, data, start);
StructName.encode_size(obj);
StructName.decode_size(data, start);

// 实例方法
obj.encode(data, start);
obj.decode(data, start);
obj.encode_size();
obj.decode_size(data, start);
```

按需生成的运行时辅助函数：

```javascript
// 通用
isObj(item);
mergeDeep(target, ...sources);

// 由协议字段特性触发按需生成
createArray(length, init);
floatToUint32Bits(value);
uint32BitsToFloat(value);
doubleToUint64Bits(value);
uint64BitsToDouble(value);
stringToUTF8BytesCount(str);
stringToUTF8Bytes(str, data, start);
stringFromUTF8Bytes(data, start);
```

### ESModule

```javascript
// 静态方法
StructName.encode(obj, buffer, start);
StructName.decode(obj, data, start);
StructName.encode_size(obj);
StructName.decode_size(data, start);

// 实例方法
obj.encode(data, start);
obj.decode(data, start);
obj.encode_size();
obj.decode_size(data, start);
```

按需生成的运行时辅助函数：

```javascript
createArray(length, init);
floatToUint32Bits(value);
uint32BitsToFloat(value);
doubleToUint64Bits(value);
uint64BitsToDouble(value);
stringToUTF8BytesCount(str);
stringToUTF8Bytes(str, data, start);
stringFromUTF8Bytes(data, start);
```

## 贡献

欢迎为 Bubbler 做出贡献。

## 许可证

MIT 许可证

## 相关仓库

- [CoralReefPlayer](https://github.com/DawningW/CoralReefPlayer) - 珊瑚礁播放器，一款现代化跨平台低延迟流媒体播放器库。
- [OpenFinNAV](https://github.com/redlightASl/OpenFinNAV) - 鳍航，针对水下机器人（ROV/AUV）设计的飞控固件库。
