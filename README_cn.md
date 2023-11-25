# Bubbler

[English](README.md) | 简体中文

Bubbler 是一个专为物联网设备优化的协议生成器。它编译 `.bb` 协议文件，并以指定的目标语言生成输出。

Bubbler 的协议功能强大，而且可以是非字节对齐的，这对于资源有限的物联网设备非常有用。下面会详细解释。

警告：Bubbler 仍在开发中，尚未准备好用于生产。

## 安装

```sh
git clone https://github.com/xaxys/bubbler.git
cd bubbler
go build
```

## 使用

```sh
bubbler [options] <input file>
```

### 选项

- `-t <target>`: 目标语言
- `-o <output>`: 输出文件

### 目标语言

运行 `bubbler` 以查看支持的目标语言列表。

### 示例

```sh
bubbler -t c -o gen.c example.bb
bubbler -t dump example.bb
```

## 协议语法

Bubbler 使用一种简洁的语法来定义数据结构和枚举类型。

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

### 数据结构

使用 `struct` 关键字定义数据结构。数据结构的定义包括结构名称和一系列的字段。例如：

```bubbler
struct Frame[20] {
    FrameType opcode;
    struct some_embed[1] {
        bool valid[#1];
        bool error[#1];
        uint8 source[#3];
        uint8 target[#3];
    };
    uint8<18> payload;
};
```

在这个例子中，`Frame` 是一个数据结构，它有三个字段：`opcode`、`some_embed` 和 `payload`。`opcode` 的类型是 `FrameType`，`some_embed` 是一个匿名内嵌的数据结构，`payload` 的类型是 `uint8`。

### 字段类型

Bubbler 协议支持四种字段类型：普通字段、匿名内嵌字段、常量字段和空字段。

- 普通字段：由类型名、字段名和字段宽度（可选）构成。
- 匿名内嵌字段：一个匿名的字段，可以是struct定义或已定义的struct名称，其内部子字段会被提升并展开到父结构体中。
- 常量字段：一个固定值的字段，其值在定义时就已经确定，不能被修改。字段名可选，如果有字段名，会生成对应字段。编码时，常量字段的值会被忽略。解码时，常量字段的值会被检查，如果不匹配，会报错。
- 空字段：一个没有名字和类型的字段，只有宽度，用于占位。

#### 普通字段

普通字段由类型名、字段名和字段宽度构成。例如：

```bubbler
struct Frame {
    RovlinkFrameType opcode;
};
```

在这个例子中，`opcode` 是一个普通字段，其类型为 `RovlinkFrameType`。

字段宽度可选，如果不填写宽度，则字段宽度为类型的宽度。

字段宽度可以小于类型的宽度，例如：

```c
struct Frame[20] {
    int64 myInt48[6];
};
```

在这个例子中，`myInt48` 是一个 6 字节的字段，其类型为 `int64`，但是它的宽度为 6 字节，因此它编码时只会占用 6 字节的空间。

但是，对于`struct`类型的字段，字段宽度必须等于类型的宽度（暂定）

### 匿名内嵌字段

匿名内嵌字段是一个没有名字的数据结构，它可以包含多个子字段。例如：

```bubbler
struct Frame {
    int64 myInt48[6];
    struct some_embed[1] {
        bool valid[#1];
        bool error[#1];
        uint8 source[#3];
        uint8 target[#3];
    };
};
```

在这个例子中，`some_embed` 是一个匿名内嵌字段，它包含了四个子字段：`valid`、`error`、`source` 和 `target`。

匿名内嵌字段的子字段会被提升并展开到父结构体中。生成的结构如下：

```c
struct Frame {
    int64_t myInt48;
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
    int64 myInt48[6];
    AnotherTest;
    uint8<18> payload;
};
```

这样，生成的结构如下：

```c
struct Frame {
    int64_t myInt48;
    int8_t arr[2];
    uint8_t payload;
};
```

### 常量字段

常量字段是一个固定值的字段，它的值在定义时就已经确定，不能被修改。例如：

```bubbler
struct Frame {
    uint8 FRAME_HEADER = 0xAA;
};
```

在这个例子中，`FRAME_HEADER` 是一个常量字段，其值为 `0xAA`。

常量字段的值在编码时会被忽略，解码时会被检查，如果不匹配，会报错。

### 空字段

空字段是一个没有名字和类型的字段，它只有宽度。空字段通常用于填充或对齐数据结构。例如：

```bubbler
struct Frame {
    void [#2];
};
```

在这个例子中，`void [#2]` 是一个空字段，它占用了 2 比特的空间。


### 字段选项

字段选项用于指定字段的额外属性。例如，可以使用 `order` 选项指定数组的字节顺序：

```bubbler
struct AnotherTest {
    int8<2> arr [order = "big"];
}
```

在这个例子中，`arr` 字段的字节顺序被设置为大端序。

### 自定义 getter/setter

可以为字段定义自定义的 getter 和 setter 方法，用于在读取或写入字段值时执行特定的操作。例如：

```bubbler
struct SensorTemperatureData {
    uint16 temperature[2] {
        get(float64): value / 10 - 40;
        set(float64): value == 0 ? 0 : (value + 40) * 10;
        set AnotherCustomSet(uint8): value == 0 ? 0 : (value + 40) * 10;
    };
}
```

在这个例子中，`temperature` 字段有一个自定义的 getter 方法和两个自定义的 setter 方法。

默认getter返回`float64`类型，并根据`value / 10 - 40`计算结果返回。其中`value`被填充为字段的值，是uint16类型。

默认setter接收`float64`类型的参数，并根据`value == 0 ? 0 : (value + 40) * 10`计算结果并以此设置字段的值。其中`value`被填充为参数的值，是float64类型。

自定义setter名为`AnotherCustomSet`，`uint8`是参数类型。并根据`value == 0 ? 0 : (value + 40) * 10`计算结果并以此设置字段的值。其中`value`被填充为参数的值，是uint8类型。

## 贡献

欢迎为 Bubbler 做出贡献。

## 许可证

MIT 许可证
