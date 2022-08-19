## chapter01

### go编译和运行说明

- go源文件，通过编译器将其编译成机器可识别的二进制码文件
- 在该源文件目录下，通过go build 对 hello.go 文件进行编译。可以指定生成的可执行文件名，在windows下，必须是.exe后缀
- 如果程序无误会生成一个二进制可执行文件
- 如果程序有误会报那一行出现的错误
- 运行的两种形式
	- 执行运行生成的go程序，hello.exe
	- 通过运行工具 go run 对源代码文件进行运行

## chapter02
### go 语言 转义字符 （escape char)

| 转义字符 | 含义                       |
| -------- | -------------------------- |
| \t       | 一个制表位，实现对齐的功能 |
| \n       | 换行符                     |
| `\\`     | 一个\                      |
| `\"`     | 一个"                      |
| \r       | 回车符                     |

```go
func main() {
	// 演示转义字符的使用 \t
	fmt.Println("tom\tjack")

	// 如果希望一次性注释 ctrl + / 第一次表示注释，第二次表示取消注释
	fmt.Println("hello\nworld")
	fmt.Println("C:\\Users\\Administrator\\Desktop\\Go语言核心编程课程\\资料")
	fmt.Println("tom说\"i love you\"")

	// \r 回车,从当前行的最前面开始输出，覆盖掉以前内容
	fmt.Println("天龙八部雪山飞狐\r张飞厉害")

	fmt.Println("helloworldhelloworldhelloworldhelloworl\n",
		"dhelloworldhelloworldhelloworldhelloworldhelloworldhellowor\n",
		"ldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworl\n",
		"dhelloworldhelloworldhelloworldhelloworldhelloworldhelloworldhel\n",
		"loworldhelloworldhelloworldhelloworldhelloworldhelloworldhellowor\n",
		"ldhelloworldhelloworldhelloworldhelloworldhelloworldhelloworld")

	// var num = 2 + 4 * 5
}
```

### 注释



#### 介绍注释

用于注解说明程序的文字就是注释，注释提高了代码的阅读性

注释是一个程序员必须要具有的良好编程习惯，将自己的思想通过注释先整理出来，再用代码体现

#### 行注释

```go
// 行注释
```



#### 块注释(多行注释)

```go
/*
块注释
*/
```

### 规范的代码风格

#### 正确的注释和注释风格

- go官方推荐使用行注释来注释整个方法和语句

#### 正确的缩进和空白

- 使用一次tab操作，实现缩进默认整体向右边移动，使用shift + tab 整体向左移
- 使用 gofmt 来进行格式化
- 

### golang 标准库 API文档

- API是Golang提供的基本编程接口
- Go语言提供了大量的标准库
- [Golang中文网](https://studygolang.com/pkgdoc)


### chapter02 总结

#### Go语言的SDK是什么？

SDK 就是软件开发工具包，做Go开发，首先需要先安装并配置好sdk

#### Goland环境变量配置及其作用

GOROOT：指定go sdk 安装目录

Path：指令 sdk\bin 目录：go.exe godoc.exe gofmt.exe

GOPATH：就是goland工作目录，所有项目源码都在这个目录下



## chapter03 变量

### 变量的介绍

#### 变量的概念

变量相当于内存中一个数据存储空间的表示，通过变量名可以访问到变量值。

#### 变量的使用步骤

- 声明变量
- 非变量赋值
- 使用变量

#### 变量使用的注意事项

- 变量表示内存的一个存储区域
- 该区域有自己的名称(变量名) 和类型(数据类型)

![image-20220306173822951](imgs\image-20220306173822951.png)

- go 变量使用的三种方式
  - 指定变量类型，声明后若不赋值，使用默认值
  - 根据值自行判断变量类型(类型推导)
  - 省略 var，:= 左侧的变量不应该是已经声明过的，否则会导致编译错误
- 多变量声明
- 该区域的数据值可以在同一个类型范围内不断变化
- 变量在同一个作用域(在一个函数或者代码块) 内不能重名
- 变量 = 变量名 + 值 + 数据类型
- Goland的变量如果没有赋初始值，编译器会使用默认值，比如int默认值0，string默认值为空串，小数默认为0

### 变量的声明，初始化和赋值

#### 变量声明

声明变量的一般形式是使用 var 关键字：

```go
var name type
```

其中，var 是声明变量的关键字，name 是变量名，type 是变量的类型。

**标准格式**

```go
var 变量名 变量类型
```

批量格式

```go
var (
    a int
    b string
    c []float32
    d func() bool
    e struct {
        x int
    }
)
```

简短格式

```go
名字 := 表达式
```

需要注意的是，简短模式（short variable declaration）有以下限制：

- 定义变量，同时显式初始化。
- 不能提供数据类型。
- 只能用在函数内部

#### 初始化变量

```
// 变量名 类型 = 表达式
var hp = 10
hb := 10
```



#### 给变量赋值

```go
var a int = 100
var b int = 200
b, a = a, b
```

#### 匿名变量

```go
func GetData() (int, int) {
    return 100, 200
}
func main(){
    a, _ := GetData()
    _, b := GetData()
    fmt.Println(a, b)
}
```



#### 程序中 + 号的使用

- 当左右两边都是数值型时，则做加法运算
- 当左右两边都是字符串，则做字符串拼接

### 数据类型的基本介绍

![image-20220306175051033](imgs\image-20220306175051033.png)

### 整数类型

#### 整数的各个类型

![image-20220306175157902](imgs\image-20220306175157902.png)

#### int 的 其他类型

![image-20220306175241823](imgs\image-20220306175241823.png)

#### 整型的使用细节

- Golang 各整数类型分: 有符号和无符号， int uint 的大小和系统有关

- Golang 的整型默认声明为 int 型

- 如何在程序查看某个变量的字节大小和数据类型

  ```go
  // unsafe.Sizeof(n1) 是unsafe包的一个函数，可以返回n1变量占用的字节数
  fmt.Printf("n2 的 类型 %T  n2占用的字节数是 %d ", n2, unsafe.Sizeof(n2))
  ```

  

- golang 程序中整型变量在使用时，遵守保小不保大的原则,即在保证程序正确运行下，尽量使用占用内存空间小的数据类型

- bit：计算机中的最小的存储单位。byte：计算机中基本存储单元。

### 小数类型/浮点型

#### 小数类型分类

![image-20220306180336279](imgs\image-20220306180336279.png)

- 关于浮点数在机器中存放形式的简单说明，浮点数 = 符号位 + 指数位 + 尾数位
- 尾数部分可能丢失，造成精度丢失
  - float64的精度比float32要更准确
  - 要保存一个精度更高的数，则应该选用float64
- 浮点型的存储分为三部分：符号位 + 指数位 + 尾数位 在存储过程中，精度有丢失



#### 浮点型使用细节

- golang 浮点类型有固定的范围和字段长度，不受操作系统的影响
- golang 的浮点型默认声明为float64 类型
- 浮点型常量有两种形式
  - 十进制数形式：如：5.12 (必须有小数点)
  - 科学计数法形式 如 5.1234e2 = 5.12 * 10 的2次方 5.12E-2 = 5.12 / 10的2次方
- 推荐使用float64 它比 float32 更精确

### 字符类型

golang 中没有专门的字符类型，如果要存储单个字符，一般使用byte来吧保存

字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。也就是说对于传统的字符串是由字符组成的，而Go的字符串不同，它是由字节组成的。

```go
// 演示golang中字符类型使用
func main() {
	
	var c1 byte = 'a'
	var c2 byte = '0' //字符的0

	// 当我们直接输出byte值，就是输出了的对应的字符的码值
	// 'a' ==> 
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	// 如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	// var c3 byte = '北' //overflow溢出
	var c3 int = '北' //overflow溢出
	fmt.Printf("c3=%c c3对应码值=%d\n", c3, c3)

	// 可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的unicode 字符
	var c4 int = 22269 // 22269 -> '国' 120->'x'
	fmt.Printf("c4=%c\n", c4)

	// 字符类型是可以进行运算的，相当于一个整数,运输时是按照码值运行
	var n1 = 10 + 'a' //  10 + 97 = 107
	fmt.Println("n1=", n1)

}
```

- 保存的字符在ASCLL表的，比如[0-1, a-z, A-Z] 直接可以保存到byte
- 保存的字符对应码值大于255，这时可以考虑使用int类型保存
- 按照字符的方式输出，需要格式化输出，fmt.Printf("c4=%c\n", c4)

#### 字符类型使用细节

- 字符常量是用单引号('') 括起来的单个字符。例如：var c1  byte = 'a'

- go 中允许使用转义 '\' 来将其后的字符转变为特殊字符型常量。 例如： var c3  int = '\n'
- go 语言的字符使用uft-8编码，英文字符 1个字节，汉字 3个字节
- 在go中，字符的本质是一个整数，直接输出时，是该字符对应的utf-8编码的码值。
- 可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的unicode字符

```go
// 可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的unicode 字符
var c4 int = 22269 // 22269 -> '国' 120->'x'
fmt.Printf("c4=%c\n", c4)
```

- 字符类型是可以进行运算的，相当于一个整数

```go
// 字符类型是可以进行运算的，相当于一个整数,运输时是按照码值运行
var n1 = 10 + 'a' //  10 + 97 = 107
fmt.Println("n1=", n1)
```

#### 字符类型本质

- 字符型 存储到 计算机中，需要将字符对应的码值(整数) 找出来

  存储： 字符 --> 对应码值 --> 二进制 --> 存储

  读取： 二进制 --> 码值 --> 字符 --> 读取

- 字符和码值得对应关系是通过字符编码表决定的

- go语言的编码都统一成 utf-8。



### 布尔类型

#### 基本介绍

- 布尔类型即bool类型， bool类型数据只允许取值 true 和 false
- bool 类型占1个字节
- bool 类型适用于逻辑运算，一般用于程序流程控制
  - if 条件控制语句
  - for 循环控制语句

### string 字符串类型

#### 基本介绍

字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。Go语言的字符串的字节使用UTF-8编码表示Unicode文本

#### string使用注意事项和细节

- 字符串一旦赋值了，字符串就不能修改了，在go中字符串是不可变的

- 字符串的两种表示形式

  - 双引号, 会识别转义字符
  - 反引号，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果

- 使用 + 号 拼接字符串

- 当一行字符串太长时可以使用加号拼接

  ```go
  // 当一个拼接的操作很长时，怎么办，可以分行写,但是注意，需要将+保留在上一行.
  str4 := "hello " + "world" + "hello " + "world" + "hello " + 
  "world" + "hello " + "world" + "hello " + "world" + 
  "hello " + "world"
  fmt.Println(str4)
  ```


### 基本数据类型的默认值

#### 基本数据类型的默认值

![image-20220308220859379](imgs\image-20220308220859379.png)

```go
var a int          // 0
var b float32      // 0
var c float64      // 0
var isMarried bool // false
var name string    // ""
```



### 基本数据类型的相互转换

Go在不同类型的变量之间赋值需要显式转化,Golang中数据类型不能自动转换。

#### 基本语法

```go
var i int32 = 100
//希望将 i => float
var n1 float32 = float32(i)
var n2 int8 = int8(i)
var n3 int64 = int64(i) // 低精度->高精度
```

#### 基本数据类型相互转换的注意事项

- Go中，数据类型的转换可以是从表示范围小 --> 表示范围大，也可以为范围大 --> 范围小

- 被转换的是变量存储的数据，变量本身的数据类型并没有变化

- 在转换中，将数据范围值大的转成范围小的要注意数据溢出。

  ```go
  // 在转换中，比如将 int64  转成 int8 【-128---127】 ，编译时不会报错，
  // 只是转换的结果是按溢出处理，和我们希望的结果不一样
  var num1 int64 = 999999
  var num2 int8 = int8(num1) // 
  fmt.Println("num2=", num2) // num2 = 63
  ```

数据转换

```go
//func main() {

	//课堂练习 , tab键
	// var n1 int32 = 12
	// var n2 int64
	// var n3 int8

	// n2 = int64(n1) + 20  //int32 ---> int64 错误
	// n3 = int8(n1) + 20  //int32 ---> int8 错误
	// fmt.Println("n2=", n2, "n3=", n3)


	// var n1 int32 = 12
	// var n3 int8
	// var n4 int8
	// n4 = int8(n1) + 127  //【编译通过，但是 结果 不是127+12 ,按溢出处理】
	// n3 = int8(n1) + 128 //【编译不过】
	// fmt.Println(n4)


//}
```

### 基本数据类型和string类型的转换

#### 基本类型转string类型

- 使用 fmt.Sprintf

  ```go
  func Sprintf(format string, a ...interface{}) string		
  ```

  Sprintf根据format参数生成格式化的字符串并返回该字符串。

  ```go
  var num1 int = 99
  var num2 float64 = 23.456
  var b bool = true
  var myChar byte = 'h'
  var str string //空的str
  
  //使用第一种方式来转换 fmt.Sprintf方法
  
  str = fmt.Sprintf("%d", num1)
  fmt.Printf("str type %T str=%q\n", str, str)
  
  str = fmt.Sprintf("%f", num2)
  fmt.Printf("str type %T str=%q\n", str, str)
  
  str = fmt.Sprintf("%t", b)
  fmt.Printf("str type %T str=%q\n", str, str)
  
  str = fmt.Sprintf("%c", myChar)
  fmt.Printf("str type %T str=%q\n", str, str)
  ```

- 使用 strconv 包 函数

  ```go
  func FormatBool(b bool) string
  func FormatInt(i int64, base int) string
  func FormatUint(i uint64, base int) string
  func FormatFloat(f float64, fmt byte, prec, bitSize int) string
  ```

  ```go
  // 第二种方式 strconv 函数
  	var num3 int = 99
  	var num4 float64 = 23.456
  	var b2 bool = true
  
  	str = strconv.FormatInt(int64(num3), 10)
  	fmt.Printf("str type %T str=%q\n", str, str)
  	
  	// strconv.FormatFloat(num4, 'f', 10, 64)
  	// 说明： 'f' 格式 10：表示小数位保留10位 64 :表示这个小数是float64
  	str = strconv.FormatFloat(num4, 'f', 10, 64)
  	fmt.Printf("str type %T str=%q\n", str, str)
  
  	str = strconv.FormatBool(b2)
  	fmt.Printf("str type %T str=%q\n", str, str)
  
  	//strconv包中有一个函数Itoa
  	var num5 int64 = 4567
  	str = strconv.Itoa(int(num5))
  	fmt.Printf("str type %T str=%q\n", str, str)
  ```

  



#### string 转基本类型

使用 strconv包的函数

```go
func ParseBool(str string) (value bool, err error)
func ParseInt(s string, base int, bitSize int) (i int64, err error)
func ParseUint(s string, base int, bitSize int) (n uint64, err error)
func ParseFloat(s string, bitSize int) (f float64, err error)
```

```go
// 演示golang中string转成基本数据类型
func main() {

	var str string = "true"
	var b bool
	// b, _ = strconv.ParseBool(str)
	// 说明
	// 1. strconv.ParseBool(str) 函数会返回两个值 (value bool, err error)
	// 2. 因为我只想获取到 value bool ,不想获取 err 所以我使用_忽略
	b , _ = strconv.ParseBool(str)
	fmt.Printf("b type %T  b=%v\n", b, b)
	
	var str2 string = "1234590"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T  n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)


	// 注意：
	var str4 string = "hello"
	var n3 int64 = 11
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type %T n3=%v\n", n3, n3)

}
```

#### string转基本数据类型的注意事项

在将string 类型转成 基本数据类型时，要确保String 类型能够转成有效的数据，比如 我们可以把‘123’， 转成一个整数，不能把 ‘hello‘转成一个整数。



### 指针

#### 基本介绍

- 基本数据类型，变量存的就是值，叫做值类型

- 获取变量的地址，用&， 比如： var num int, 获取 num 的地址： &num

  分析基本数据类型在内存的布局

- 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值

  ```go
  var ptr *int = &num
  ```

  指针在内存中的布局

  ![image-20220308225300989](imgs\image-20220308225300989.png)

- 获取指针类型所指向的值，使用：*， 比如 var ptr *int ，`*ptr`获取ptr指向的值

```go
// 演示golang中指针类型
func main() {

	// 基本数据类型在内存布局
	var i int = 10
	// i 的地址是什么,&i
	fmt.Println("i的地址=", &i)
	
	//下面的 var ptr *int = &i
	//1. ptr 是一个指针变量
	//2. ptr 的类型 *int
	//3. ptr 本身的值&i
	var ptr *int = &i 
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr 的地址=%v", &ptr) 
	fmt.Printf("ptr 指向的值=%v", *ptr)

}
```

- ![image-20220308225600230](imgs\image-20220308225600230.png)

- 案例
  - 写一个程序，获取一个int变量num的地址，并显示到终端
  - 将num的地址赋给指针 ptr， 并通过ptr去修改num值

​	![image-20220308230016461](imgs\image-20220308230016461.png)

#### 指针的使用细节

- 值类型，都有对应的指针类型，形式为 *数据类型，比如 int 的对应改的指针就是 *int，float32 对应的指针类型就是 *float32 依次类推。
- 值类型包括：基本数据类型 int系列， float类型， bool，string， 数组和结构体 struct。

### 值类型和引用类型

#### 值类型和引用类型说明

- 值类型：基本数据类型int系列，float系列，bool，string，数组和结构体struct
- 引用类型：指针、slice切片、map、管道chan、interface等都是引用类型

#### 值类型和引用类型的使用特点

- 值类型：变量直接存储值，内存通常在栈中分配

  ![image-20220308233124158](imgs\image-20220308233124158.png)

- 引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据(值)，内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC来回收。

  ![image-20220308233328717](imgs\image-20220308233328717.png)

- 内存的栈区和堆区示意图

  ![image-20220308233412262](imgs\image-20220308233412262.png)

### 标识符的命名规范

#### 标识符概念

- golang对各种变量、方法、函数等命名时使用的字符序列称为标识符
- 凡是起名字的地方都叫标识符

- 标识符的命名规则
  - 由26个英文字母大小写，0-9，_ 组成
  - 数字不可以开头， var num int //ok
  - golang 中严格区分大小写
- 标识符不能包含空格
- 下划线"_" 本身在Go中是一个特殊的标识符，称为空标识符，可以代表任何其他的标识符，但是它对应的值会被忽略，所以仅能被作为占位符使用，不能作为标识符使用。
- 不能使用系统保留关键字作为标识符

#### 标识符命名

- 包名：保持package 的名字和目录保持一致，尽量采取有意义的包名，简短有意义，不要和标准库冲突
- 变量名、函数名、常量名：采用驼峰法
- 如果变量名、函数名、常量名首字母大写，则可以被其他的包访问：如果首字母小写，则只能在本包中使用，首字母大写是公开的，首字母小写是私有的

### 系统保留关键字

![image-20220308234315938](imgs\image-20220308234315938.png)

### 系统的预定义标识符

![image-20220308234404546](imgs\image-20220308234404546.png)

## chapter04 运算符

### 运算符的基本介绍

- 算术运算符
- 赋值运算符
- 比较运算符/关系运算符
- 逻辑运算符
- 位运算符
- 其他运算符

### 算术运算符

算符运算符是对数值类型的变量进行运算

![image-20220308234715256](imgs\image-20220308234715256.png)

+、-、*、/、%、++、--

自增：++

自减：--

#### 算术运算符使用的注意事项

- 对于除号"/"，它的整数除和小数除是有区别的

- 当对于一个数取模时，可以等价 a%b = a - a/b*b

- golang的自增自减只能当做独立语言使用

  ```go
  a = i++ // 错误
  b = i-- // 错误
  ```

- golang的 ++ 和 -- 只能写在变量的后面，不能写在变量的前面

### 关系运算符(比较运算符)

#### 基本介绍

- 关系运算符的结果都是bool型，也就是true，要么是false
- 关系表达式 经常用在 if结构的条件中或循环结构的条件中

![image-20220308235648311](imgs\image-20220308235648311.png)

```go
// 演示关系运算符的使用
var n1 int = 9
var n2 int = 8
fmt.Println(n1 == n2) //false
fmt.Println(n1 != n2) //true
fmt.Println(n1 > n2) //true
fmt.Println(n1 >= n2) //true
fmt.Println(n1 < n2) //flase
fmt.Println(n1 <= n2) //flase
flag := n1 > n2
fmt.Println("flag=", flag)
```

#### 关系运算符的细节说明

- 关系运算符的结果都是bool型，也就是要么是true，要么是false
- 关系运算符组成的表达式，称为关系表达式：a>b
- 比较运算符 == 比较是否相等

### 逻辑运算符

#### 基本介绍

用于连接多个条件，最终结果也是一个bool值

#### 逻辑运算的说明

| 运算符 | 描述                                                         | 实例 |
| ------ | ------------------------------------------------------------ | ---- |
| &&     | 逻辑与运算符，如果两边的操作数都是True，则为True,否则为False |      |
| \|\|   | 逻辑或运算符，如果两边的操作数有一个True,则为True，否则为False |      |
| ！     | 逻辑非运算符，如果条件为True,则逻辑为False,否则为True        |      |

- && 也叫短路与，如果第一个条件为false.则第二个条件不会判断，最终结果为false
- || 也叫短路或，如果第一个条件为true，则第二个条件不会判断，最终结果为true

### 赋值运算符

![image-20220312131530711](imgs\image-20220312131530711.png)



![image-20220312131547001](imgs\image-20220312131547001.png)

#### 赋值运算符的特点

- 运算顺序从右到左
- 赋值运算符的左边，只能是变量，右边是变量、表达式、常量值
- 复合赋值运算符等价效果 a += 3  -> a = a + 3

### 位运算符

![image-20220312131818293](imgs\image-20220312131818293.png)

### 其他运算符

![image-20220312132126219](imgs\image-20220312132126219.png)

### 运算符

![image-20220312132330195](imgs\image-20220312132330195.png)



### 键盘输入语句

- 导入fm包
- 调用 ftm 包的 fmt.Scanln() 或者 fmt.Scanf()

![image-20220312132918612](imgs\image-20220312132918612.png)



###  原码、反码、补码

- 二进制的最高位为符号位：0表示整数，1表示负数
- 整数的原码、反码、补码
- 负数的反码 = 它的原码符号位不变其他位取反
- 负数的补码 = 它的反码 + 1
- 0 的反码，补码都是0

###  位运算符和移位运算符

#### 位运算符

按位与 &

按位或 |

按位异或 ^

#### 移位运算符

`>>` 右移 ：低位溢出，符号位不变，并用符号位补溢处的高位

`<<` 左移：符号位不变，低位补0

## chapter05 程序流程控制

### 程序流程控制介绍

- 顺序控制
- 分支控制
- 循环控制

###  顺序控制

![image-20220309000904976](imgs\image-20220309000904976.png)

### 分支流程

分支流程的三种形式

- 单分支
- 双分支
- 多分支

#### 单分支

单分支流程图

![image-20220309001213911](imgs\image-20220309001213911.png)

```go
// golang支持在if中，直接定义一个变量，比如下面
if age := 20; age > 18 {
    fmt.Println("你年龄大于18,要对自己的行为负责!")
}
```

#### 双分支

双分支流程图

![image-20220309001320399](imgs\image-20220309001320399.png)

#### 多分支

多分支流程图

![image-20220309001405243](imgs\image-20220309001405243.png)

#### 嵌套分支

在一个分支结构中又嵌套另一完整的分支结构，里面的内层分支，外面的分支结构为外层分支。

#### switch分支控制

switch语句用于基于不同条件执行不同动作，每一个case分支都是唯一的，从上到下逐一测试直到匹配为止，不需要添加break。

switch分支路程图

![image-20220309001934071](imgs\image-20220309001934071.png)

- switch的执行流程是，先执行表达式，得到值，然后和case的表达式进行比较，如果相等就匹配到，然后执行对应的case的语句块，然后退出switch控制
- 如果swtich的表达式的值没有和任何case的表达式匹配成功，则执行default的语句块。执行后退出switch控制。
- golang的case后的表达式可以有多个，使用逗号间隔。
- golang中的case语句块不需要写break，因为默认会有，即在默认情况下，当程序执行完case语句块后，就直接退出该switch控制结构。

#### switch的使用的注意事项和细节

- case/switch后是一个表达式（常量值、变量、一个有返回值得函数等）
- case后的各个表达式的值的数据类型，必须和switch的表达式数据类型一致
- case后面可以带多个表达式，使用逗号间隔。比如case 表达式1，表达式2.
- case后面的表达式如果是常量值，则要求不能重复
- case后面不需要带break，程序匹配到一个case后一个case后就会执行对应的代码块，然后退出switch，如果一个都匹配不到，则执行default
- default语句不是必须的
- switch后也可以不带表达式
- switch 穿透 fallthrough,如果在case语句块后增加 fallthrough,则会继续执行下一个case
- Type Switch: switch 语句还可以别用于 type-switch 来判断某个interface 变量中实际指向的变量类型

#### switch 和 if的比较

- 如果判断的具体数值不多，而且符合整数、浮点数、字符、字符串这几种类型。建议使用switch语句。
- 其他情况，对区间判断和结果为bool类型的判断，使用if，if的使用范围更广。

### 循环控制

#### for循环

语法格式

```go
for 循环变量初始化; 循环条件; 循环变量迭代 {
    循环操作(语句)
}
```

语法格式说明

- 循环变量初始化
- 循环条件
- 循环操作即循环体
- 循环变量迭代

#### for循环执行流程分析

![image-20220312122042873](imgs\image-20220312122042873.png)

#### for循环的使用注意事项和细节讨论

- 循环条件是返回一个布尔值的表达式

- for循环的第二种使用方式

  ```go
  for 循环判断条件 {
  	// 循环执行语句
  }
  ```

- for循环的第三种使用方式，等价 for; ; {} 是一个无线循环，通常需要配合break语句使用

  ```go
  for {
      // 循环执行语句
  }
  ```

- golang 提供 for-range 的方式，可以方便字符串和数组

  ```go
  str = "abc~ok上海"
  for index, val := range str {
      fmt.Printf("index=%d, val=%c \n", index, val)
  }
  ```

  

#### while 和 do while 的实现

使用for循环实现while循环

```go
var i int = 1
for {
    if i > 10 { //循环条件
        break // 跳出for循环,结束for循环
    }
    fmt.Println("hello,world", i)
    i++ //循环变量的迭代
}
```



使用for循环实现do while循环

```go
// 使用的do...while实现完成输出10句”hello,ok“
var j int = 1
for {
    fmt.Println("hello,ok", j)
    j++ // 循环变量的迭代
    if j > 10 {
        break //break 就是跳出for循环
    }
}
```

#### 跳转控制语句 break

![image-20220312134612405](imgs\image-20220312134612405.png)

#### break的注意事项

- break语句出现在多层嵌套的语句块中时，可以通过标签指明要终止的是哪一层语句块

```go
// 这里演示一下指定标签的形式来使用 break
lable2: 
for i := 0; i < 4; i++ {
    //lable1: // 设置一个标签
    for j := 0; j < 10; j++ {
        if j == 2 {
            //break // break 默认会跳出最近的for循环
            //break lable1 
            break lable2 // j=0 j=1
        }
        fmt.Println("j=", j) 
    }
}
```



- break 默认会跳出最近的for循环
- break 后面可以指定标签，跳出标签对应的for循环

#### 跳转控制语句 continue

- continue 语句用于结束本次循环，继续执行下一次循环
- continue 语句出现在多层嵌套的循环语句体中时，可以通过标签指明要跳过的是哪一层循环。

![image-20220312135200442](imgs\image-20220312135200442.png)

```go
here:
for i:=0; i<2; i++ {
    for j:=1; j<4; j++ {
        if j==2 {
            continue here
        }
        fmt.Println("i=",i,"j=",j)
    }
}
```

### goto





### 跳转控制语句 return

- 如果 return 是在普通的函数，则表示跳出该函数，即不再执行函数中 return 后面代码，也可以理解成终止函数。
- 如果 return 是在 main函数，表示 main 函数，也就是说终止程序。

## chapter06 函数、包、错误处理

```go
func 函数名 (形参列表) (返回值列表) {
    执行语句 ...
    return 返回值列表
}
```

### 包

#### 包的基本概念

go的每一个文件都是属于一个包的，也就是说go是以包的形式来管理文件和醒目目录结构的

#### 包的三大作用

- 区分相同名字的函数、变量等标识符
- 当程序文件很多时，可以很好的管理项目
- 控制函数、变量等访问范围，即作用域

#### 包的相关说明

打包基本语法

 package 包名

引入包的基本语法

import "包的路径"

#### 包使用的注意事项和细节讨论

- 在给一个文件打包时，该包对应一个文件夹，比如这里的 utils文件夹对应的包名就是utils， 文件的包名通常和文件所在的文件夹名一致，一般为小写字母。

- 当一个文件要使用其他包函数或变量时，需要先引入对应的包

  ```go
  // 方式一
  import "fmt"
  
  // 方式二
  import (
  	"GoCourse/src/chapter06/funcdemo01/utils"
  	"fmt"
  )
  ```

  在 import 包时，路径从 $GOPATH 的 src 下开始，不用带 src,编译器会自动从src下开始引入

- 为了让其他包的文件，可以访问到本包的函数，该函数名的首字母需要大写

- 在访问其他函数，变量时，其语法是 包名.函数名

- 如果包名较长， 支持给包取别名

- 在同一个包下，不能有相同的函数名

### 函数

#### 函数调用过程

![image-20220312141337162](imgs\image-20220312141337162.png)

- 在调用一个函数时，会给该函数分配一个新的空间，编译器会通过自身处理让这个新的空间和其他栈的空间区分开来
- 在每个函数对应的栈中，数据空间是独立的，不会混淆
- 当一个函数调用完毕后，程序会销毁这个函数对应的栈空间

#### return语句

go语言支持返回多个值，这一点是其他编程语言没有的

```
func 函数名(形参列表) (返回值类型列表) {
	语句 ...
	return 返回值列表
}
```

- 如果返回多个值时，在接收时，希望忽略某个返回值，则使用 _ 符号表示占位忽略
- 如果返回值只有一个

#### 递归函数



```go
func test2(n int) {
	if n > 2 {
		n-- //递归必须向退出递归条件逼进，否则就是无限循环调用
		test2(n)
	} else {
		fmt.Println("n=", n)
	}
}
```

![image-20220312142119223](imgs\image-20220312142119223.png)

函数递归需要遵守的重要原则

- 执行一个函数时，就创建一个新的受保护的独立空间
- 函数的局部变量是独立的，不会互相影响
- 递归必须向退出递归条件逼近，否则就是无限逼近
- 当一个函数执行完毕，或者互道return，就会返回。

```go
func fbn(n int) int  {
	if (n == 1 || n == 0) {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func main() {
	fmt.Println(fbn(4))
}
```

#### 函数使用的注意事项和细节讨论

- 函数的形参可以是多个，返回值列表页可以是多个

- 形参列表和返回值列表的数据类型可以是指类型和引用类型

- 函数的命名遵循表示符命名规范，首字母不能是数字，首字母大写可被其他包访问。

-  函数中的变量是局部的，函数外不生效。

- 基本数据类型和数组默认都是值传递的，即进行值拷贝。

- 函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量

  ```go
  // n1 就是 *int 类型
  func test03(n1 *int) {
  	fmt.Printf("n1的地址 %v\n", &n1)
  	*n1 = *n1 + 10
  	fmt.Println("test03() n1= ", *n1) // 30
  }
  ```

  

- go函数不支持函数重载

- 在go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量，通过该变量可以对函数调用

  ```go
  func getSum(n1 int, n2 int) int {
  	return n1 + n2
  }
  
  a := getSum
  fmt.Printf("a的类型%T, getSum类型是%T\n", a, getSum)
  
  res := a(10, 40) // 等价  res := getSum(10, 40)
  fmt.Println("res=", res)
  ```

  

- 函数既然是一种数据类型，在go中，函数可以作为形参

  ```go
  func getSum(n1 int, n2 int) int {
  	return n1 + n2
  }
  
  // 函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
  func myFun2(funvar myFunType, num1 int, num2 int) int {
  	return funvar(num1, num2)
  }
  
  
  //看案例
  res2 := myFun(getSum, 50, 60)
  fmt.Println("res2=", res2)
  ```

  

- go支持自定义数据类型

  基本语法: type 自定义数据类型名 数据类型  // 相当于别名

- 支持对函数返回值命名

  ```go
  // 支持对函数返回值命名
  func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
  	sub = n1 - n2
  	sum = n1 + n2
  	return
  }
  ```

  

- 使用 _ 标识符，忽略返回值

- go 支持可变参数

  ```go
  // 案例演示： 编写一个函数sum ,可以求出  1到多个int的和
  // 可以参数的使用
  func sum(n1 int, args ...int) int {
  	sum := n1
  	//遍历args
  	for i := 0; i < len(args); i++ {
  		sum += args[i] // args[0] 表示取出args切片的第一个元素值，其它依次类推
  	}
  	return sum
  }
  ```

  

#### init 函数

每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被Go运行框架调用，init会在main函数之前被调用

##### init函数使用细节

- 如果一个文件同时包含全局变量定义，init函数和main函数，则执行的流程 全局变量定义 -> init 函数 -> main函数
- init函数最主要的作用，就是完成一些初始化的工作
- main.go 和 utils.go 中都有变量定义和函数初始化

![image-20220312145015943](imgs\image-20220312145015943.png)

#### 匿名函数

匿名函数使用方式1

```go
// 案例演示,求两个数的和， 使用匿名函数的方式完成
res1 := func(n1 int, n2 int) int {
    return n1 + n2
}(10, 20)

fmt.Println("res1=", res1)
```

匿名函数使用方式2

```go
// 将匿名函数func (n1 int, n2 int) int赋给 a变量
// 则a 的数据类型就是函数类型 ，此时,我们可以通过a完成调用
a := func(n1 int, n2 int) int {
    return n1 - n2
}

res2 := a(10, 30)
fmt.Println("res2=", res2)
```

全局匿名函数

```go
// fun1就是一个全局匿名函数
Fun1 = func(n1 int, n2 int) int {
    return n1 * n2
}

func main() {
	// 全局匿名函数的使用
	res4 := Fun1(4, 9)
	fmt.Println("res4=", res4)
}
```

#### 闭包

基本介绍：闭包就是一个函数和与其相关的引用环境组合的一个整体

```go
// 累加器
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(36)        // => 36 = '$'
		fmt.Println("str=", str) // 1. str="hello$" 2. str="hello$$" 3. str="hello$$$"
		return n
	}
}


func main() {

	// 使用前面的代码
	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16
}
```

- AddUpper 是一个函数，返回的数据类型是 fun(int) int

- 闭包的说明

  返回的是一个匿名函数，但是这个匿名函数引用到函数外的n,因此匿名函数和n形成了一个整体，构成闭包。

- 当反复调用f函数时，因为n是初始化一次，因此每调用一次就进行累计

- 弄清楚闭包的关键，就是要分析出返回的函数它使用到哪些变量，因为函数和它引用到的变量共同构成闭包

#### 闭包练习

编写程序要求

- 编写一个函数 makeSuffix(suffix string) 可以接受一个文件后缀名(比如 .jpg), 并返回一个闭包
- 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀（比如.jpg），则返回 文件名.jpg,如果有文件名返回原文件名
- 要求使用闭包的方式完成
- strings.HasSuffix, 该函数可以判断某个字符串是否有指定的后缀

```go
//
// 1)编写一个函数 makeSuffix(suffix string)  可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
// 2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg , 如果已经有.jpg后缀，则返回原文件名。
// 3)要求使用闭包的方式完成
// 4)strings.HasSuffix , 该函数可以判断某个字符串是否有指定的后缀。

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		//如果 name 没有指定后缀，则加上，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

func main() {
    // 试makeSuffix 的使用
	// 返回一个闭包
	f2 := makeSuffix(".jpg")               //如果使用闭包完成，好处是只需要传入一次后缀。
	fmt.Println("文件名处理后=", f2("winter"))   // winter.jgp
	fmt.Println("文件名处理后=", f2("bird.jpg")) // bird.jpg
}
```

- 返回的匿名函数和 makeSuffix(suffix string) 的 stuffix 变量 组成一个闭包，因为 返回的函数引用到 suffix 这个变量
- 传统的方式也可以轻松实现这个功能，但是需要每次传入后缀名







#### 函数的defer

在函数中，程序员经常需要创建资源，为了在函数执行完毕后，及时的释放资源，go的设计者提供了defer(延时机制)


#### defer的使用细节

- go执行一个defer时，不会立即执行defer后的语句，而是将defer后的语句压入到一个栈中，然后继续执行函数
- 当函数执行完毕后，再从defer栈职工，依次从栈顶取出语句

```go
func sum(n1 int, n2 int) int {

	// 当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈(defer栈)
	// 当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1=", n1) //defer 3. ok1 n1 = 10
	defer fmt.Println("ok2 n2=", n2) //defer 2. ok2 n2= 20
	//增加一句话
	n1++                         // n1 = 11
	n2++                         // n2 = 21
	res := n1 + n2               // res = 32
	fmt.Println("ok3 res=", res) // 1. ok3 res= 32
	return res

}
```

defer的应用场景

- 在golang编程中通常的做法是，创建资源后，（比如打开文件、创建数据库连接）可以执行defer file.Close() defer connect.Close()
- 在defer后，可以继续创建资源
- 当函数完毕后，系统会依次从defer栈中，取出语句，关闭资源
- defer机制，程序员不需要在为什么时机关闭资源担心

#### 函数参数传递方式

两种传递方式

- 值传递
- 引用传递

值类型和引用类型

值类型默认是值传递：变量直接存储值，内存通常在栈中分配

![image-20220312222204301](imgs\image-20220312222204301.png)

引用类型默认是引用传递：变量存储的是一个地址，这个地址对应的空间才真正存储数据(值)，内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC来回收。

![image-20220312222403754](imgs\image-20220312222403754.png)

如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量。

![image-20220312222533141](imgs\image-20220312222533141.png)

#### 变量作用域

- 函数内部声明/定义的变量叫做局部变量，作用域仅限于函数内部
- 函数外部声明/定义的变量叫全局变量，作用域在整个包中都有效，如果首字符为大写，则作用域在整个程序有效
- 如果变量在一个代码块，比如 for/if 中，那么变量的作用域就在该代码块中

#### 字符串常用的系统函数

```go
func main() {

	// 1)统计字符串的长度，按字节 len(str)
	// golang的编码统一为utf-8 (ascii的字符(字母和数字) 占一个字节，汉字占用3个字节)
	str := "hello北"
	fmt.Println("str len=", len(str)) // 8

	str2 := "hello北京"
	// 2)字符串遍历，同时处理有中文的问题 r := []rune(str)
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 3)字符串转整数:	 n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果是", n)
	}

	// 4)整数转字符串  str = strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T\n", str, str)

	// 5)字符串 转 []byte:  var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	// 6)[]byte 转 字符串: str = string([]byte{97, 98, 99})
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	// 7)10进制转 2, 8, 16进制:  str = strconv.FormatInt(123, 2),返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制是=%v\n", str)

	// 8)查找子串是否在指定的字符串中: strings.Contains("seafood", "foo") //true
	b := strings.Contains("seafood", "mary")

	fmt.Printf("b=%v\n", b)

	// 9)统计一个字符串有几个指定的子串 ： strings.Count("ceheese", "e") //4
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v\n", num)

	// 10)不区分大小写的字符串比较(==是区分字母大小写的): fmt.Println(strings.EqualFold("abc", "Abc")) // true

	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b) //true

	fmt.Println("结果", "abc" == "Abc") // false //区分字母大小写

	// 11)返回子串在字符串第一次出现的index值，如果没有返回-1 :
	// strings.Index("NLT_abc", "abc") // 4

	index := strings.Index("NLT_abcabcabc", "abc") // 4
	fmt.Printf("index=%v\n", index)

	// 12)返回子串在字符串最后一次出现的index，
	// 如没有返回-1 : strings.LastIndex("go golang", "go")

	index = strings.LastIndex("go golang", "go") //3
	fmt.Printf("index=%v\n", index)

	// 13)将指定的子串替换成 另外一个子串: strings.Replace("go go hello", "go", "go语言", n)
	// n可以指定你希望替换几个，如果n=-1表示全部替换

	str2 = "go go hello"
	str = strings.Replace(str2, "go", "北京", -1)
	fmt.Printf("str=%v str2=%v\n", str, str2)

	// 14)按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：
	// strings.Split("hello,wrold,ok", ",")
	strArr := strings.Split("hello,wrold,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)

	// 15)将字符串的字母进行大小写的转换:
	// strings.ToLower("Go") // go strings.ToUpper("Go") // GO

	str = "goLang Hello"
	str = strings.ToLower(str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str) //golang hello

	// 将字符串左右两边的空格去掉： strings.TrimSpace(" tn a lone gopher ntrn   ")
	str = strings.TrimSpace(" tn a lone gopher ntrn   ")
	fmt.Printf("str=%q\n", str)

	// 17)将字符串左右两边指定的字符去掉 ：
	// strings.Trim("! hello! ", " !")  // ["hello"] //将左右两边 ! 和 " "去掉
	str = strings.Trim("! he!llo! ", " !")
	fmt.Printf("str=%q\n", str)

	// 20)判断字符串是否以指定的字符串开头:
	// strings.HasPrefix("ftp://192.168.10.1", "ftp") // true

	b = strings.HasPrefix("ftp://192.168.10.1", "hsp") //true
	fmt.Printf("b=%v\n", b)
}
```



#### 时间和日期相关函数

- 时间和日期相关函数，需要导入time包

- time.Time 类型，用于表示时间

- 如何获取到其他的日期信息

  ```go
  // 看看日期和时间相关函数和方法使用
  // 1. 获取当前时间
  now := time.Now()
  fmt.Printf("now=%v now type=%T\n", now, now)
  
  // 2.通过now可以获取到年月日，时分秒
  fmt.Printf("年=%v\n", now.Year())
  fmt.Printf("月=%v\n", now.Month())
  fmt.Printf("月=%v\n", int(now.Month()))
  fmt.Printf("日=%v\n", now.Day())
  fmt.Printf("时=%v\n", now.Hour())
  fmt.Printf("分=%v\n", now.Minute())
  fmt.Printf("秒=%v\n", now.Second())
  ```

- 格式化日期时间

  - 使用Printf 或者 Sprintf
  - 使用time.Format() 方法完成

- 时间的常量

  ```go
  const (
  	Nanosecond  Duration = 1
  	Microsecond          = 1000 * Nanosecond
  	Millisecond          = 1000 * Microsecond
  	Second               = 1000 * Millisecond
  	Minute               = 60 * Second
  	Hour                 = 60 * Minute
  )
  ```

  常量的作用：在程序中可用于获取执行时间单位的时间，比如想得到100毫秒 100 * Millsecond

- 结合Sleep来使用一下时间常量，可以查看执行时间

  ```go
  // 需求，每隔1秒中打印一个数字，打印到100时就退出
  // 需求2: 每隔0.1秒中打印一个数字，打印到100时就退出
  i := 0
  for {
      i++
      fmt.Println(i)
      // 休眠
      // time.Sleep(time.Second)
      time.Sleep(time.Millisecond * 100)
      if i == 100 {
          break
      }
  }
  ```

  

- time 的 Unix 和 UnixNano

  ![image-20220312215551787](imgs\image-20220312215551787.png)

  ```go
  // Unix和UnixNano的使用
  fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano())
  // unix时间戳=1647093474 unixnano时间戳=1647093474377446500
  ```

#### 内置函数

golang设计者为了编程方便，提供了一些函数，这些函数可以直接使用，这些称为go的内置函数。

- len：用来求长度，比如string、array、slice、map、channel

- new：用来分配内存，主要用来分配值类型，比如int、float32, struct... 返回的是指针

  ![image-20220312220615981](imgs\image-20220312220615981.png)

- make：用来分配内存，主要用来分配引用类型，比如channel、map、slice。

### 错误处理

```go
num1 := 10
num2 := 0
res := num1 / num2
fmt.Println("res=", res)
```

- 在默认情况下，当发生错误后（panic），程序就会退出
- 当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行。还可以捕获到错误后，给出提示
- 使用错误处理机制

#### 错误机制基本说明 

- go语言追求简洁优雅，不支持try catch finally
- go中引入的处理方式：defer，panic，recover
- 这几个异常的使用场景，go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理

使用 defer + recover来处理错误

```go
func test() {
	// 用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover()内置函数，可以捕获到异常
		if err != nil {  // 说明捕获到错误
			fmt.Println("err=", err)
			// 这里就可以将错误信息发送给管理员....
			fmt.Println("发送邮件给admin@sohu.com~")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}
```

进行错误处理后，程序不会轻易挂掉，如果加入预警代码，就可以让程序更加的健壮。

#### 自定义错误

go程序中，也支持自定义错误，使用errors.New 和 panic 内置函数

- errors.New("错误说明")，会返回一个error类型的值，表示一个错误
- panic内置函数，接收一个interface{}类型的值作为参数。可以接收error类型的变量，输出错误信息，并退出程序

```go
func test02() {

	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行....")
}

```

##  chapter07 数组与切片

### 数组

#### 数组定义和布局

数组的定义	

```go
var 数组名 [数组大小] 数据类型
var a [5] int
```

数组在内存布局

![image-20220312223836189](imgs\image-20220312223836189.png)

- 数组的地址可以通过数组名来获取 &intArr
- 数组的第一个元素的地址，就是数组的首地址
- 数组的各个元素的地址间隔是依据数组的类型决定

#### 数组的使用

数组名[下标] 

四种初始化数组的方式

```go
// 四种初始化数组的方式
var numArr01 [3]int = [3]int{1, 2, 3}
fmt.Println("numArr01=", numArr01)

var numArr02 = [3]int{5, 6, 7}
fmt.Println("numArr02=", numArr02)
// 这里的 [...] 是规定的写法
var numArr03 = [...]int{8, 9, 10}
fmt.Println("numArr03=", numArr03)

var numArr04 = [...]int{1: 800, 0: 900, 2: 999}
fmt.Println("numArr04=", numArr04)

// 类型推导
strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
fmt.Println("strArr05=", strArr05)
```

#### 数组的遍历

- 常规遍历 for循环遍历

  ```go
  var intArr [6]int = [...]int{1, -1, 9, 90, 11, 9000}
  maxVal := intArr[0]
  maxValIndex := 0
  
  for i := 1; i < len(intArr); i++ {
      // 然后从第二个元素开始循环比较，如果发现有更大，则交换
      if maxVal < intArr[i] {
          maxVal = intArr[i]
          maxValIndex = i
      }
  }
  fmt.Printf("maxVal=%v maxValIndex=%v\n\n", maxVal, maxValIndex)
  ```

  

- for-range 结构遍历

```go
strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
fmt.Println("strArr05=", strArr05)
for i, v := range strArr05 {
    fmt.Println(i, v)
}
```

说明

- 第一个返回值index是数组的下标
- 第二个value是在该下标的值
- 他们都是仅在for循环内部可见的局部变量
- 遍历数组元素的时候，如果不想使用下标index，可以直接把下标标为下划线_
- index 和 value 的名称不是固定的

```go
// 请求出一个数组的和和平均值。for-range
// 思路
// 1. 就是声明一个数组  var intArr[5] = [...]int {1, -1, 9, 90, 11}
// 2. 求出和sum
// 3. 求出平均值
//代码
var intArr2 [5]int = [...]int{1, -1, 9, 90, 12}
sum := 0
for _, val := range intArr2 {
    //累计求和
    sum += val
}
```

#### 数组使用的注意事项和细节

- 数组时多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化
- var arr [] int 这时 arr 就是一个切片
- 数组中的元素可以是任何数据类型，包括值引用和引用类型，但是不能混用。
- 数组创建后，如果没有赋值，有默认值
- 使用数组的步骤
  - 声明数组并开辟空间
  - 给数组各个元素赋值
  - 使用数组
- 数组的下标是 从0开始的
- 数组下标必须在指定范围内使用，否则报panic：数组越界
- go的数组属于值类型，在默认情况下是值传递，因此会进行值拷贝。数组间不会互相影响
- 在其他函数职工，去修改原来的数组，可以使用引用传递
- 长度是数组类型的一部分，在传递参数时，需要考虑数组的长度

### 切片

#### 切片的基本介绍

- 切片的英文是slice

- 切片是数组的一个引用，因此切片是引用类型，在进行传递时 ，遵守引用传递机制

- 切片的使用和数组类似，遍历切片、访问切片的元素和切片长度 len(slice)都一样

- 切片的长度是可以变化的，因此切片是一个可以动态变化数组

- 切片定义的基本语法

  var 切片名 [] 类型

#### 切片在内存中形式

![image-20220313112337678](imgs\image-20220313112337678.png)

- slice是一个引用类型

- slice 从底层来说，其实就是一个数据结构(struct 结构体)

  ```go
  type slice struct {
      ptr *[2]int
      len int
      cap int
  }
  ```

#### 切片的使用

第一种方式

```go
// 示切片的基本使用
var intArr [5]int = [...]int{1, 22, 33, 66, 99}
// 声明/定义一个切片
// slice := intArr[1:3]
// 1. slice 就是切片名
// 2. intArr[1:3] 表示 slice 引用到intArr这个数组
// 3. 引用intArr数组的起始下标为 1 , 最后的下标为3(但是不包含3)
slice := intArr[1:3]
fmt.Println("intArr=", intArr)
fmt.Println("slice 的元素是 =", slice)       //  22, 33
fmt.Println("slice 的元素个数 =", len(slice)) // 2
fmt.Println("slice 的容量 =", cap(slice))   // 切片的容量是可以动态变化
```

第二种方式：通过make来创建切片

var 切片名 [] type = make([] type, len, [cap])

参数说明：type：就是数据类型 len：大小 cap：指定切片容量，可选如果分配了cap。则要求cap > len

![image-20220313113212850](imgs\image-20220313113212850.png)

- 通过 make 方式 创建切片可以指定切片的大小和容量
- 如果没有给切片的各个元素赋值，那么就会使用默认值[int, float => 0 string => "" bool => false]
- 通过 make 方式创建的切片对应的数组是由 make1底层维护，对外不可见，即只能通过slice去访问各个元素

第三种方式：定义一个切片，直接就指定具体数组，使用原理类似 make 的方式。

```go
// 方式3
fmt.Println()
// 第3种方式：定义一个切片，直接就指定具体数组，使用原理类似make的方式
var strSlice []string = []string{"tom", "jack", "mary"}
fmt.Println("strSlice=", strSlice)
fmt.Println("strSlice size=", len(strSlice)) //3
fmt.Println("strSlice cap=", cap(strSlice))  // ?
```

#### 切片的遍历

- for循环常规方式遍历

  ```go
  // 使用常规的for循环遍历切片
  var arr [5]int = [...]int{10, 20, 30, 40, 50}
  // slice := arr[1:4] // 20, 30, 40
  slice := arr[1:4]
  for i := 0; i < len(slice); i++ {
      fmt.Printf("slice[%v]=%v ", i, slice[i])
  }
  ```

  

- for-range 结构遍历切片

  ```go
  var arr [5]int = [...]int{10, 20, 30, 40, 50}
  // slice := arr[1:4] // 20, 30, 40
  slice := arr[1:4]
  // 使用for--range 方式遍历切片
  for i, v := range slice {
      fmt.Printf("i=%v v=%v \n", i, v)
  }
  ```

#### 切片的使用的细节

- 切片初始化时， var slice = arr[startIndex:endIndex]]

  从arr数组下标为startIndex， 取到 下标为 endIndex 的元素

- 切片初始化时，仍然不能越界，范围在 [0 - len(arr)] 之间，但是可以动态增长

  var slice = arr[0:end] 可以简写 var slice = arr[:end]

  var slice = ar[start:len(arr)] 可以简写： var slice = arr[start:]

  var slice = arr[0:len(arr)] 可以简写：var slice = [:]

- cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
  
- 切片定义完后，还不能使用，因为本省是一个空的，需要其引用到一个数组，或者make 一个空间切片
  
- 切片继续切片
  
  ```go
  slice2 := slice[1:2] //  slice [ 20, 30, 40]    [30]
  slice2[0] = 100      // 因为arr , slice 和slice2 指向的数据空间是同一个，因此slice2[0]=100，其它的都变化
  ```
  
- 用 append 内置函数，可以对切片进行动态追加

  ```go
  // 用append内置函数，可以对切片进行动态追加
  var slice3 []int = []int{100, 200, 300}
  // 通过append直接给slice3追加具体的元素
  slice3 = append(slice3, 400, 500, 600)
  fmt.Println("slice3", slice3) //100, 200, 300,400, 500, 600
  ```

  ![image-20220313115635260](imgs\image-20220313115635260.png)

- 切片 append 操作的底层原理分析

  切片 append 操作的本质就是对数组扩容

  go 底层会创建一下新的数组 newArr

  将  slice 原来包含的元素拷贝到新的数组 newArr

  slice 重新引用到 newArr

- 切片的拷贝操作

  切片使用copy内置函数完成拷贝，

  ```go
  var slice4 []int = []int{1, 2, 3, 4, 5}
  var slice5 = make([]int, 10)
  copy(slice5, slice4)
  fmt.Println("slice4=", slice4) // 1, 2, 3, 4, 5
  fmt.Println("slice5=", slice5) // 1, 2, 3, 4, 5, 0 , 0 ,0,0,0
  ```

- 切片是引用类型，所以在传递时，遵守引用传递机制

  

  

#### string 和 slice

- string底层是一个 byte 数组，因此 string 也可以进行切片处理

- string 和 切片在内存的形式

  ![image-20220313120221213](imgs\image-20220313120221213.png)

- string是不可变的，也就说不能通过 str[0] = 'z' 方式来修改字符串

- 如果需要修改字符串，可以先将 string => [] byte 或者 [] rune => 修改 -> 重写转成 string

```go
// string底层是一个byte数组，因此string也可以进行切片处理
str := "hello@atguigu"
// 使用切片获取到 atguigu
slice := str[6:]
fmt.Println("slice=", slice)

arr1 := []rune(str)
arr1[0] = '北'
str = string(arr1)
fmt.Println("str=", str)
```

斐波那契使用切片实现

```go

func fbn(n int) []uint64 {
	// 声明一个切片，切片大小 n
	fbnSlice := make([]uint64, n)
	// 第一个数和第二个数的 为1
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	//进行for循环来存放的数列
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}

	fmt.Printf("%T\n", fbnSlice)
	return fbnSlice
}

func main() {

	/*
		1)可以接收一个 n int
		2)能够将斐波那契的数列放到切片中
		3)提示, 斐波那契的数列形式:
		arr[0] = 1; arr[1] = 1; arr[2]=2; arr[3] = 3; arr[4]=5; arr[5]=8

		思路
		1. 声明一个函数 fbn(n int) ([]uint64)
		2. 编程fbn(n int) 进行for循环来存放斐波那契的数列  0 =》 1 1 =》 1
	*/

	//测试一把看看是否好用
	fnbSlice := fbn(20)
	fmt.Println("fnbSlice=", fnbSlice)

}
```

## chapter08 排序和查找

### 排序

排序的基本介绍

排序是将一组数据，依指定的顺序进行排列的过程

- 内部排序

  指将需要处理的所有数据都加载到内部存储器中进行排序

  包括(交换排序法、选择式排序法和插入式排序法)

- 外部排序

  数据量过大，无法全部加载到内存中，需要借助外部存储进行排序，包括（合并排序法和直接合并排序法）







#### 冒泡排序
![image-20220313122314503](imgs\image-20220313122314503.png)

```go
// 冒泡排序
func BubbleSort(arr *[5]int) {

	fmt.Println("排序前arr=", (*arr))
	temp := 0 //临时变量(用于做交换)

	//冒泡排序..一步一步推导出来的
	for i := 0; i < len(*arr)-1; i++ {

		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				//交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}

	}

	fmt.Println("排序后arr=", (*arr))

}

func main() {

	// 定义数组
	arr := [5]int{24, 69, 80, 57, 13}
	// 将数组传递给一个函数，完成排序

	BubbleSort(&arr)

	fmt.Println("main arr=", arr) //有序? 是有序的
}
```



### 查找

- 顺序查找
- 二分查找



#### 顺序查找

```go
func main() {
	//有一个数列：白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王
	//猜数游戏：从键盘中任意输入一个名称，判断数列中是否包含此名称【顺序查找】
	//思路
	//1 定义一个数组, 白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王 字符串数组
	//2.从控制台接收一个名字，依次比较，如果发现有，提示

	// 代码
	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Println("请输入要查找的人名...")
	fmt.Scanln(&heroName)

	// 顺序查找:第一种方式
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			fmt.Printf("找到%v , 下标%v \n", heroName, i)
			break
		} else if i == (len(names) - 1) {
			fmt.Printf("没有找到%v \n", heroName)
		}
	}

	//顺序查找:第2种方式
	index := -1

	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i //将找到的值对应的下标赋给 index
			break
		} 
	}
	if index != -1 {
		fmt.Printf("找到%v , 下标%v \n", heroName, index)
	} else {
		fmt.Println("没有找到", heroName)
	}


}
```

#### 二分查找

有序数组进行二分查找 {1, 8, 10,  89, 1000, 1234， 输入一个数查看数组中是否包含此数

![image-20220313123421996](imgs\image-20220313123421996.png)

```go
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {

	// 判断leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	// 先找到 中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		//说明我们要查找的数，应该在  leftIndex --- middel-1
		BinaryFind(arr, leftIndex, middle - 1, findVal)
	} else if (*arr)[middle] < findVal {
		//说明我们要查找的数，应该在  middel+1 --- rightIndex
		BinaryFind(arr, middle + 1, rightIndex, findVal)
	} else {
		//找到了
		fmt.Printf("找到了，下标为%v \n", middle)
	}
}
```

### 二维数组

使用方式1： 先定义，再赋值

var 数组 [大小] [大小] 类型  var arr [2] [3] int  再赋值

```go
// 定义/声明二维数组
var arr [4][6]int
// 赋初值
arr[1][2] = 1
arr[2][1] = 2
arr[2][3] = 3

// 遍历二维数组，按照要求输出图形
for i := 0; i < 4; i++ {
    for j := 0; j < 6; j++ {
        fmt.Print(arr[i][j], " ")
    }
    fmt.Println()
}

```

使用方式2：直接初始化

声明： var 数组名` [大小][大小]` 类型 = [大小] [大小] 类型 {{初值}, {初值}}

赋值 （默认值， 比如int类型的就是0）

```go
arr3  := [2][3]int{{1,2,3}, {4,5,6}}
fmt.Println("arr3=", arr3)
```

#### 二维数组的遍历

- 双层for循环完成遍历

  ```go
  var arr3  = [2][3]int{{1,2,3}, {4,5,6}}
  
  // for循环来遍历
  for i := 0; i < len(arr3); i++ {
     for j := 0; j < len(arr3[i]); j++ {
        fmt.Printf("%v\t", arr3[i][j])
     }
     fmt.Println()
  }	
  ```

-  for - range 方式完成遍历

  ```go
  // for-range来遍历二维数组
  for i, v := range arr3 {
      for j, v2 := range v {
          fmt.Printf("arr3[%v][%v]=%v \t",i, j, v2)
      }
      fmt.Println()	
  }
  ```

  

二维数组的练习

```go
func main() {

	/*
	定义二维数组，用于保存三个班，每个班五名同学成绩，
	并求出每个班级平均分、以及所有班级平均分
	*/

	//1.定义一个二维数组
	var scores [3][5]float64
	//2.循环的输入成绩
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d班的第%d个学生的成绩\n", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	//fmt.Println(scores)

	//3.遍历输出成绩后的二维数组，统计平均分
	totalSum := 0.0 // 定义一个变量，用于累计所有班级的总分
	for i := 0; i < len(scores); i++ {
		sum := 0.0 //定义一个变量，用于累计各个班级的总分
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		totalSum += sum 
		fmt.Printf("第%d班级的总分为%v , 平均分%v\n", i+1, sum, 
			sum / float64(len(scores[i])))
	}

	fmt.Printf("所有班级的总分为%v , 所有班级平均分%v\n", 
		totalSum, totalSum / 15 )
}
```

## chapter09 map

### map 

map 是key-value数据结构，又称为字段或者关联数组

#### map的使用

方式1

```go
// 第一种使用方式
	
var a map[string]string
//在使用map前，需要先make , make的作用就是给map分配数据空间
a = make(map[string]string, 10)
a["no1"] = "宋江" //ok?
a["no2"] = "吴用" //ok?
a["no1"] = "武松" //ok?
a["no3"] = "吴用" //ok?
fmt.Println(a)
```

方式2

```go
// 第二种方式
cities := make(map[string]string)
cities["no1"] = "北京"
cities["no2"] = "天津"
cities["no3"] = "上海"
fmt.Println(cities)
```

方式3

```go
// 第三种方式
heroes := map[string]string{
    "hero1" : "宋江",
    "hero2" : "卢俊义",
    "hero3" : "吴用",
}
heroes["hero4"] = "林冲"
fmt.Println("heroes=", heroes)
```

#### map的增删改查

```go
func main() {
	// 第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	// 因为 no3这个key已经存在，因此下面的这句话就是修改
	cities["no3"] = "上海~"
	fmt.Println(cities)

	// 演示删除
	delete(cities, "no1")
	fmt.Println(cities)
	// 当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

	// 演示map的查找
	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有no1 key 值为%v\n", val)
	} else {
		fmt.Printf("没有no1 key\n")
	}

	// 如果希望一次性删除所有的key
	// 1. 遍历所有的key,如何逐一删除 [遍历]
	// 2. 直接make一个新的空间
	cities = make(map[string]string)
	fmt.Println(cities)

}
```

#### map 遍历

```go
cities := make(map[string]string)
cities["no1"] = "北京"
cities["no2"] = "天津"
cities["no3"] = "上海"

for k, v := range cities {
    fmt.Printf("k=%v v=%v\n", k, v)
}
```



#### map 切片

切片的数据类型如果是map，则称为 slice of map, map切片，这样map的个数就可以改变了

```go
func main() {
	// 演示map切片的使用
	/*
		要求：使用一个map来记录monster的信息 name 和 age, 也就是说一个
		monster对应一个map,并且妖怪的个数可以动态的增加=>map切片
	*/
	// 1. 声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2) // 准备放入两个妖怪
	//2. 增加第一个妖怪的信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}

	// 下面这个写法越界。
	// if monsters[2] == nil {
	// 	monsters[2] = make(map[string]string, 2)
	// 	monsters[2]["name"] = "狐狸精"
	// 	monsters[2]["age"] = "300"
	// }

	// 这里我们需要使用到切片的append函数，可以动态的增加monster
	// 1. 先定义个monster信息
	newMonster := map[string]string{
		"name": "新的妖怪~火云邪神",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)

	fmt.Println(monsters)
}
```

#### map 排序

- golang中没有一个专门的方法针对 map 的 key 进行排序
- golang 中的 map 默认是无序的，添加也是无序的
- golang 中 的 map 的排序，是先将key 进行排序，然后根据key值遍历输出即可

```go
func main() {

	// map的排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	fmt.Println(map1)

	// 如果按照map的key的顺序进行排序输出
	// 1. 先将map的key 放入到 切片中
	// 2. 对切片排序
	// 3. 遍历切片，然后按照key来输出map的值

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	// 排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("map1[%v]=%v \n", k, map1[k])
	}

}
```





#### map 使用细节

- map 是引用类型，遵守引用类型传递的机制，在一个函数接收 map，修改后会直接修改原来的map

- map 的容量达到后，再想 map增加元素，会自动扩容，并不会发生panic，map能动态的增长键值对
- map 的 value 也经常使用 struct 类型

#### map 练习

```go
func main() {
	// 使用for-range遍历map
	// 第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	fmt.Println("cities 有", len(cities), " 对 key-value")

	// 使用for-range遍历一个结构比较复杂的map
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街~"

	studentMap["stu02"] = make(map[string]string, 3) // 这句话不能少!!
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海黄浦江~"

	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
		fmt.Println()
	}
}
```

## chapter10 面向对象上

### 结构体

#### 结构体变量和结构体变量(实例)的区别

结构体是自定义的数据类型，代表一类事物

- 结构体变量是具体的，实际的，代表一个具体变量

  ![image-20220313133608079](imgs\image-20220313133608079.png)

####  声明结构体

```go
type 结构体名称 struct {
    field type
    field2 type
}
```

```go
type Student struct {
    Name string // 字段
    Age int // 字段
    Score float32 
}
```

#### 字段/属性

- 从概念或叫法上看： 结构体字段 = 属性 = field 
- 字段是结构体的一个组成部分，一般是基本数据类型、数组也可以是引用类型。

注意事项和细节说明

- 字段声明语法同变量，示例：字段名 字段类型

- 字段的类可以为：基本类型、数组或引用类型

- 在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值

  指针， slice 和 map 的零值都是 nil，即还没有分配空间

- 不同结构体变量的字段是独立的，互不影响，一个结构体变量字段的更改，不影响另外一个，结构体是值类型

  ![image-20220313134801305](imgs\image-20220313134801305.png)

#### 创建结构体变量和访问结构体字段

```go
func main() {
	// 方式1
	var p1 Person
	p1.Age = 19
	p1.Name = "Mike"
	fmt.Println(p1)
	
	// 方式2
	p2 := Person{"mary", 20}
	// p2.Name = "tom"
	// p2.Age = 18
	fmt.Println(p2)

	// 方式3-&
	//案例: var person *Person = new (Person)

	var p3 *Person= new(Person)
	//因为p3是一个指针，因此标准的给字段赋值方式
	//(*p3).Name = "smith" 也可以这样写 p3.Name = "smith"

	//原因: go的设计者 为了程序员使用方便，底层会对 p3.Name = "smith" 进行处理
	//会给 p3 加上 取值运算 (*p3).Name = "smith"
	(*p3).Name = "smith" 
	p3.Name = "john" //

	(*p3).Age = 30
	p3.Age = 100
	fmt.Println(*p3)

	//方式4-{}
	//案例: var person *Person = &Person{}

	//下面的语句，也可以直接给字符赋值
	//var person *Person = &Person{"mary", 60} 
	var person *Person = &Person{}

	//因为person 是一个指针，因此标准的访问字段的方法
	// (*person).Name = "scott"
	// go的设计者为了程序员使用方便，也可以 person.Name = "scott"
	// 原因和上面一样，底层会对 person.Name = "scott" 进行处理， 会加上 (*person)
	(*person).Name = "scott"
	person.Name = "scott~~"

	(*person).Age = 88
	person.Age = 10
	fmt.Println(*person)

}
```

- 第3种和第4种方式返回的是 结构体指针
- 结构体指针访问字段的标准方式应该是： (*结构体指针).字段名
- go支持简化， 支持结构体.字段名，比如 person.Name = "tom"。更加符合程序员使用的习惯，go 编译器底层对 person.Name 做了转化 (*person).Name

#### struct 类型的内存分配机制



```go
func main() {

	var p1 Person
	p1.Age = 10
	p1.Name = "小明"
	var p2 *Person = &p1 //这里是关键-->画出示意图

	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)
	p2.Name = "tom~"
	fmt.Printf("p2.Name=%v p1.Name=%v \n", p2.Name, p1.Name)    // tom~ tom~
	fmt.Printf("p2.Name=%v p1.Name=%v \n", (*p2).Name, p1.Name) // tom~ tom~

	fmt.Printf("p1的地址%p\n", &p1)
	fmt.Printf("p2的地址%p p2的值%p\n", &p2, p2)

}
```

![image-20220313140557033](imgs\image-20220313140557033.png)

![image-20220313140608620](imgs\image-20220313140608620.png)

#### 结构体使用细节

- 结构的所有字段在内存中是连续的

  ```go
  // 结构体
  type Point struct {
  	x int
  	y int
  }
  
  // 结构体
  type Rect struct {
  	leftUp, rightDown Point
  }
  
  // 结构体
  type Rect2 struct {
  	leftUp, rightDown *Point
  }
  
  func main() {
  
  	r1 := Rect{Point{1, 2}, Point{3, 4}}
  
  	// r1有四个int, 在内存中是连续分布
  	// 打印地址
  	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p \n",
  		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
  
  	// r2有两个 *Point类型，这个两个*Point类型的本身地址也是连续的，
  	// 但是他们指向的地址不一定是连续
  
  	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
  
  	// 打印地址
  	fmt.Printf("r2.leftUp 本身地址=%p r2.rightDown 本身地址=%p \n",
  		&r2.leftUp, &r2.rightDown)
  
  	// 他们指向的地址不一定是连续...， 这个要看系统在运行时是如何分配
  	fmt.Printf("r2.leftUp 指向地址=%p r2.rightDown 指向地址=%p \n",
  		r2.leftUp, r2.rightDown)
  
  }
  ```

  ![image-20220313141713834](imgs\image-20220313141713834.png)

- 结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段

- 结构体进行 type 重新定义，golang认为是新的数据类型但是可以强转

  ```go
  type A struct {
  	Num int
  }
  type B struct {
  	Num int
  }
  
  var a A
  var b B
  a = A(b) // ? 可以转换，但是有要求，就是结构体的的字段要完全一样(包括:名字、个数和类型！)
  fmt.Println(a, b)
  ```

- struct 的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景是序列化和反序列化。

- ![image-20220313142352487](imgs\image-20220313142352487.png)

  ```go
  type Monster struct {
  	Name  string `json:"name"` // `json:"name"` 就是 struct tag
  	Age   int    `json:"age"`
  	Skill string `json:"skill"`
  }
  
  
  // 1. 创建一个Monster变量
  monster := Monster{"牛魔王", 500, "芭蕉扇~"}
  
  // 2. 将monster变量序列化为 json格式字串
  //   json.Marshal 函数中使用反射，这个讲解反射时，我会详细介绍
  jsonStr, err := json.Marshal(monster)
  if err != nil {
      fmt.Println("json 处理错误 ", err)
  }
  fmt.Println("jsonStr", string(jsonStr))
  ```

  ### 方法

  #### 方法的介绍

  golang中的方法 是作用在指定的数据类型上的，因此自定义类型，都可以有方法，而不仅仅是 struct。

  ```go
  type A struct {
      Num int
  }
  
  func (a A) test() {
      fmt.Print(a.NUm)
  }
  ```

  - func (a A) test() {}  表示 A 结构体有一方法，方法名为 test
  - (a A) 体现了 test 方法是和 A 类型绑定的

```go
type Person struct {
	Name string
}

// 给Person类型绑定一方法
func (person Person) test() {
	person.Name = "jack"
	fmt.Println("test() name=", person.Name) // 输出jack
}

var p Person
p.Name = "tom"
p.test()                              // 调用方法
fmt.Println("main() p.Name=", p.Name) // 输出 tom
```

- test方法和 Person类绑定
- test方法只能通过 Person 类型的变量来调用，而不能直接调用，也不能使用其它类型变量来调用
- func(p Person) test {} ... p 表示哪个 Person 变量调用，这个p就是副本。

### 方法

#### 方法的调用和传参机制

方法的调用和传参机制和函数基本一样，不一样的地方是方法调用时，会将调用方法的变量，当做实参也传递给方法。

- 通过一个变量去调用方法时，其调用机制和函数一样
- 不一样的地方是，变量调用方法时，该变量本身也会作为一个参数传递到方法（如果变量是值类型，则进行值拷贝，如果变量是引用类型，则进行地质拷贝）

#### 方法的声明

```
func (recevier type) methodName (参数列表) (返回值列表) {
	方法体
	return 返回值
}
```

- 参数列表：表示方法输入
- recevier type：表示这个方法和type这个类型进行绑定，或者说该方法作用于type类型
- recevier type：type可以是结构体，也可以是其他的自定义类型
- receiver：即使 type 类型的一个变量(实例)，
- 返回值列表：表示返回的值，可以多个
- 方法主体：表示为了实现某一功能代码块
- return 语句不是必须的

#### 方法的注意事项和细节

- 结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式

- 想要修改结构体变量的值，可以通过结构体指针的方式来处理

  ```go
  
  ```

- golang中的方法作用在指定的数据类型上的，因此自定义类型，都可以有方法，而不仅仅是struct。

  ```go
  // 编写一个方法，可以改变i的值
  func (i *integer) change() {
  	*i = *i + 1
  }
  
  var i integer = 10
  i.print()
  i.change()
  ```

  

- 方法的访问范围控制的规则，和函数一样。方法名首字母小写，只能在本包访问，方法首字母大写

- 如果一个类型实现了String() 这个方法，那么 fmt.Println 默认会调用这个变量的 String() 进行输出

  ```go
  // 给*Student实现方法String()
  func (stu *Student) String() string {
  	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
  	return str
  }面向对象
  ```

#### 面向编程应用实例

- 声明结构体，确定结构体名
- 编写结构体的字段
- 编写结构体的方法



#### 工厂模式

定义工厂的构造方法

```go
// 定义结构体
type student struct {
	Name string
	score float64
}

func NewStudent(n string, s float64) *student  {
	return &student{
		Name: n,
		score: s,
	}
}

func (s *student) GetScore()  float64 {
	return s.score
}
```

调用工厂模式



```go
func main() {
	// 通过工厂模式构造方法来访问对象
	student:= model.NewStudent("tom", 88.0)

	fmt.Println(*student)
	fmt.Println("name=", student.Name, "score=", student.GetScore())
}
```




## chapter11 面向对象下

### 面向对象的三大特性-封装

#### 基本介绍

golang仍然有面向对象编程的继承、封装和多态的特性。

#### 封装

封装就是把抽象的字段和对字段的操作封装在一起，数据被保护在内部，程序的其他包只有通过授权的操作，才能对字段进行操作。

#### 封装的好处

- 隐藏实现细节
- 可以对数据进行验证，保证安全合理



#### 如何体现封装

- 对结构体重中的属性进行封装
- 通过方法，包实现封装

#### 封装的实现步骤

- 将结构体、字段的首字母小写

- 将结构体所在包提供一个工厂模式的函数，首字母大写。
- 提供一个首字母大写的Set方法（类似其他语言的public）用于对属性判断并赋值
- 提供一个首字母大写的Get方法，用于获取属性的值

#### 封装案例

```go
type person struct {
	Name string
	age int   // 其它包不能直接访问..
	sal float64
}

// 写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name : name,
	}
}

//为了访问age 和 sal 我们编写一对SetXxx的方法和GetXxx的方法
func (p *person) SetAge(age int) {
	if age >0 && age <150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确..")
		//给程序员给一个默认值
	}
}

func (p *person) GetAge() int {
	return p.age
}


func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确..")
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
```

调用构造体

```go
func main() {

	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())
	
}
```

#### 继承

- 继承可以解决代码复用
- 当多个结构体存在相同的属性(字段)和方法时，可以从这些结构体中抽象出结构体，在结构体定义这些相同的属性和方法
- 其他的结构体不需要重新定义这些属性和方法，只需嵌套一个Student匿名结构体即可

![image-20220315223457559](imgs\image-20220315223457559.png)

在golang中，如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性

#### 继承的深入讨论

```go
type A struct {
	Name string
	age int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string 
}

func (b *B) SayOk() {
	fmt.Println("B SayOk", b.Name)
}
```



- 结构体可以使用嵌套匿名结构体所有的字段和方法

- 匿名结构体字段访问可以简化

  ```go
  var b B
  b.A.Name = "tom"
  b.A.age = 19
  b.A.SayOk()
  b.A.hello()
  
  //上面的写法可以简化
  
  b.Name = "smith"
  b.age = 20
  b.SayOk()
  b.hello()
  ```

  

  - 当直接通过b访问字段或方法时，其执行流程如下比如 b.Name
  - 编译器会先看b对应的类型有没有Name，如果有，则直接调用B类型的Name字段
  - 如果没有就去看B中嵌入的匿名结构体A 有没有声明Name字段，如果有就调用，如果没有继续查找，如果都找不到就报错

- 结构体和匿名结构体有相同字段或者方法时，编译器采用就近访问原则访问，如希望访问匿名结构体的字段和方法，可以通过匿名结构体来区分

- 结构体可以嵌入两个或多个匿名结构体，如两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)，在访问时，就必须明确指定匿名结构体名字，否则编译报错。

- 如果一个struct 嵌套了一个有名结构体，这种模式就是组合，在访问组合结构体的字段或方法时，必须带上结构体的名字

- 嵌套匿名结构体后，也可以在创建结构体变量时，直接指定各个匿名结构体字段的值

```go
package main
import (
	"fmt"
)

type A struct {
	Name string
	age int
}
type B struct {
	Name string
	Score float64
}
type C struct {
	A
	B
	// Name string
}

type D struct {
	a A // 有名结构体
}


type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand	
}

type TV2 struct {
	*Goods
	*Brand	
}

type Monster struct  {
	Name string
	Age int
}

type E struct {
	Monster
	int
	n int
}

func main() {
	var c C
	// 如果c 没有Name字段，而A 和 B有Name, 这时就必须通过指定匿名结构体名字来区分
	// 所以 c.Name 就会包编译错误， 这个规则对方法也是一样的！
	c.A.Name = "tom" // error
	fmt.Println("c")

	// 如果D 中是一个有名结构体，则访问有名结构体的字段时，就必须带上有名结构体的名字
	// 比如 d.a.Name 
	var d D 
	d.a.Name = "jack"


	// 嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	tv := TV{ Goods{"电视机001", 5000.99},  Brand{"海尔", "山东"}, }

	// 演示访问Goods的Name
	fmt.Println(tv.Goods.Name)
	fmt.Println(tv.Price) 

	tv2 := TV{ 
		Goods{
			Price : 5000.99,
			Name : "电视机002", 
		},  
		Brand{
			Name : "夏普", 
			Address :"北京",
		}, 
	}

	fmt.Println("tv", tv)
	fmt.Println("tv2", tv2)

	tv3 := TV2{ &Goods{"电视机003", 7000.99},  &Brand{"创维", "河南"}, }

	tv4 := TV2{ 
			&Goods{
				Name : "电视机004", 
				Price : 9000.99,
			},  
			&Brand{
				Name : "长虹", 
				Address : "四川",
			}, 
		}

	fmt.Println("tv3", *tv3.Goods, *tv3.Brand)
	fmt.Println("tv4", *tv4.Goods, *tv4.Brand)


	// 演示一下匿名字段时基本数据类型的使用
	var e E
	e.Name = "狐狸精"
	e.Age = 300
	e.int = 20
	e.n = 40
	fmt.Println("e=", e)

}
```

#### 多重继承

一个struct 嵌套了多个匿名结构体，那么该结构可以直接访问嵌套的匿名结构体的字段和方法，从而实现了多重继承。

```go
type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand	
}
```

多重继承细节说明

- 嵌入的匿名结构体有相同的字段名或者方法名，则在访问时，需要通过匿名结构体类型名来区分
- 为了保证代码的简洁性，尽量不使用多重继承

### 接口

#### 接口的基本介绍

interface 类型可以定义方法，但是这些方法不需要实现。并且 interface 不能包含任何变量。到某个自定义类型 要使用的时候，再根据具体情况把这些方法写出来。

基本语法

```go
type 接口名 interface {
    method1(参数列表) 返回值列表
    method2(参数列表) 返回值列表
}

func (t 自定义类型) method1(参数列表) 返回值列表 {
	// 方法实现
}
func (t 自定义类型) method2(参数列表) 返回值列表 {
	// 方法实现   
}
```

- 接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低耦合的思想。
- golang中的接口，不需要显式的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口。

####  接口使用的细节

- 接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量

  ```go
  type Stu struct {
  	Name string
  }
  
  func (stu Stu) Say() {
  	fmt.Println("Stu Say()")
  }
  
  type AInterface interface {
  	Say()
  }
  
  type BInterface interface {
  	Hello()
  }
  
  var stu Stu // 结构体变量，实现了 Say() 实现了 AInterface
  var a AInterface = stu
  a.Say()
  ```

- 接口中所有的方法都没有方法体，即都是没有实现的方法

- 在golang中，一个自定义类型需要将某个接口的所有方法都实现，就说这个自定义类型实现了该接口

- 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋给接口类型

- 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型

  ```go
  type integer int
  
  func (i integer) Say() {
  	fmt.Println("integer Say i =", i)
  }
  
  type AInterface interface {
  	Say()
  }
  
  type BInterface interface {
  	Hello()
  }
  
  var i integer = 10
  var b AInterface = i
  b.Say() // integer Say i = 10
  ```

  

- 一个自定义类型可以实现多个接口

  ```go
  type AInterface interface {
  	Say()
  }
  
  type BInterface interface {
  	Hello()
  }
  
  type Monster struct {
  }
  
  func (m Monster) Hello() {
  	fmt.Println("Monster Hello()~~")
  }
  
  func (m Monster) Say() {
  	fmt.Println("Monster Say()~~")
  }
  
  // Monster实现了AInterface 和 BInterface
  var monster Monster
  var a2 AInterface = monster
  var b2 BInterface = monster
  a2.Say()
  b2.Hello()
  ```

- 接口中不能有任何变量

- 一个接口可以继承多个别的接口，如果要实现A接口，也必定实现B、C接口的所有方法

  ```go
  type BInterface interface {
  	test01()
  }
  
  type CInterface interface {
  	test02()
  }
  
  type AInterface interface {
  	BInterface
  	CInterface
  	test03()
  }
  
  // 如果需要实现AInterface,就需要将BInterface CInterface的方法都实现
  type Stu struct {
  }
  
  func (stu Stu) test01() {
  
  }
  func (stu Stu) test02() {
  
  }
  func (stu Stu) test03() {
  
  }
  var stu Stu
  var a AInterface = stu
  a.test01()
  ```

- interface 类型默认是一个指针(引用类型)，如果没有对 interface初始化就使用，那么会输出nil

- 空接口 interface{}， 没有任何方法， 所以所有类型都可以实现空接口

  ```go
  type T interface {
  }
  
  
  var t T = stu //ok
  fmt.Println(t)
  var t2 interface{} = stu
  var num1 float64 = 8.8
  t2 = num1
  t = num1
  fmt.Println(t2, t)
  ```

  

### 接口 和 继承

接口和继承需要解决的问题不同

- 继承的价值主要在于：解决代码的复用性和可维护性
- 接口的价值主要在于：设计，设计好各种规范，让其它自定义类型去实现这些方法
- 接口比继承更加灵活，继承是满足 is - a 的关系，而接口只需满足 like-a 的关系
- 接口在一定程度上实现代码解耦

```go
// Monkey结构体
type Monkey struct {
	Name string
}

// 声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, " 生来会爬树..")
}

// LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}


// 让LittleMonkey实现BirdAble
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, " 通过学习，会飞翔...")
}

// 让LittleMonkey实现FishAble
func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, " 通过学习，会游泳..")
}

func main() {

	// 创建一个LittleMonkey 实例
	monkey := LittleMonkey{
		Monkey {
			Name : "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()

}
```



#### 多态

在go语言中，多态特征是通过接口实现的，可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态。

#### 接口实现多态的两种形式

- 多态参数

  在传入的参数中，可以根据参数不同实现不同的结构体的具体方法

  ```go
  // 声明/定义一个接口
  type Usb interface {
  	// 声明了两个没有实现的方法
  	Start()
  	Stop()
  }
  
  type Phone struct {
  }
  
  // 让Phone 实现 Usb接口的方法
  func (p Phone) Start() {
  	fmt.Println("手机开始工作。。。")
  }
  func (p Phone) Stop() {
  	fmt.Println("手机停止工作。。。")
  }
  
  type Camera struct {
  }
  
  // 让Camera 实现   Usb接口的方法
  func (c Camera) Start() {
  	fmt.Println("相机开始工作~~~。。。")
  }
  func (c Camera) Stop() {
  	fmt.Println("相机停止工作。。。")
  }
  
  // 计算机
  type Computer struct {
  }
  
  // 编写一个方法Working 方法，接收一个Usb接口类型变量
  // 只要是实现了 Usb接口 （所谓实现Usb接口，就是指实现了 Usb接口声明所有方法）
  func (c Computer) Working(usb Usb) {
  
  	// 通过usb接口变量来调用Start和Stop方法
  	usb.Start()
  	usb.Stop()
  }
  
  func main() {
  
  	// 测试
  	// 先创建结构体变量
  	computer := Computer{}
  	phone := Phone{}
  	camera := Camera{}
  
  	// 关键点
  	computer.Working(phone)
  	computer.Working(camera) //
  }
  
  ```

  

- 多态数组

```go
// 声明/定义一个接口
type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}  

// 让Phone 实现 Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}


type Camera struct {
	name string
}
// 让Camera 实现   Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}



func main() {
	// 定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	// 这里就体现出多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
    
	usbArr[2] = Camera{"尼康"}
	
	fmt.Println(usbArr)

}
```



### 类型断言






## chapter14 文件操作

### 文件操作

#### 输入流和输出流

![image-20220315234555770](imgs\image-20220315234555770.png)

流： 数据在数据源(文件) 和 程序 (内存) 之间经历的路径

输入流：数据从数据源（文件）到程序（内存）的路径

输出流：数据从程序（内存））到数据源（文件）的路径

os.File 封装所有文件相关操作，File是一个结构体

![image-20220315235055114](imgs\image-20220315235055114.png)

#### 打开文件和关闭文件

![image-20220315235250707](imgs\image-20220315235250707.png)

![image-20220315235306279](imgs\image-20220315235306279.png)

```go
func main() {
	// 打开文件
	// 概念说明: file 的叫法
	// 1. file 叫 file对象
	// 2. file 叫 file指针
	// 3. file 叫 file 文件句柄
	file , err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	// 输出下文件，看看文件是什么, 看出file 就是一个指针 *File
	fmt.Printf("file=%v", file)
	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
```

#### 读文件操作应用实例

- 读取文件的内容并显示在终端(带缓冲区的方式)，使用 os.Open()，file.close()，bufio.NewReader(), reader.ReadString 函数和方法

  ```go
  func main() {
  	//打开文件
  	//概念说明: file 的叫法
  	//1. file 叫 file对象
  	//2. file 叫 file指针
  	//3. file 叫 file 文件句柄
  	file, err := os.Open("d:/test.txt")
  	if err != nil {
  		fmt.Println("open file err=", err)
  	}
  
  	//当函数退出时，要及时的关闭file
  	defer file.Close() //要及时关闭file句柄，否则会有内存泄漏.
  
  	// 创建一个 *Reader  ，是带缓冲的
  	/*
  		const (
  		defaultBufSize = 4096 //默认的缓冲区为4096
  		)
  	*/
  	reader := bufio.NewReader(file)
  	//循环的读取文件的内容
  	for {
  		str, err := reader.ReadString('\n') // 读到一个换行就结束
  		if err == io.EOF {                  // io.EOF表示文件的末尾
  			break
  		}
  		//输出内容
  		fmt.Println(str)
  	}
  
  	fmt.Println("文件读取结束...")
  }
  
  ```

- 读取文件的内容并显示在终端(使用ioutil一次将整个文件读入到内存中)，这种适用于文件不大的情况。相关方法和函数(ioutil.ReadFile())

  ```go
  func main() {
  
  	//使用ioutil.ReadFile一次性将文件读取到位
  	file := "d:/test.txt"
  	content, err := ioutil.ReadFile(file)
  	if err != nil {
  		fmt.Printf("read file err=%v", err)
  	}
  	//把读取到的内容显示到终端
  	//fmt.Printf("%v", content) // []byte
  	fmt.Printf("%v", string(content)) // []byte
  
  	//我们没有显式的Open文件，因此也不需要显式的Close文件
  	//因为，文件的Open和Close被封装到 ReadFile 函数内部
  }
  ```





#### 写文件操作应用实例

#### 基本应用案例-方式一

- 创建一个新文件，写入内容5句，"hello, Gardon"
    ```go
    func main() {
    	//创建一个新文件，写入内容 5句 "hello, Gardon"
    	//1 .打开文件 d:/abc.txt
    	filePath := "d:/abc.txt"
    	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
    	if err != nil {
    		fmt.Printf("open file err=%v\n", err)
    		return
    	}
    	//及时关闭file句柄
    	defer file.Close()
    	//准备写入5句 "hello, Gardon"
    	str := "hello,Gardon\r\n" // \r\n 表示换行
    	//写入时，使用带缓存的 *Writer
    	writer := bufio.NewWriter(file)
    	for i := 0; i < 5; i++ {
    		writer.WriteString(str)
    	}
    	//因为writer是带缓存，因此在调用WriterString方法时，其实
    	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
    	//真正写入到文件中， 否则文件中会没有数据!!!
    	writer.Flush()
    }
    ```
- 打开一个存在的文件中，将原来的内容覆盖成新的10句 "你好， 老王"
    ```go
    func main() {
    	//打开一个存在的文件中，将原来的内容覆盖成新的内容10句 "你好，尚硅谷!"
    
    	//创建一个新文件，写入内容 5句 "hello, Gardon"
    	//1 .打开文件已经存在文件, d:/abc.txt
    	filePath := "d:/abc.txt"
    	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
    	if err != nil {
    		fmt.Printf("open file err=%v\n", err)
    		return
    	}
    	//及时关闭file句柄
    	defer file.Close()
    	str := "你好,北京!\r\n" // \r\n 表示换行
    	//写入时，使用带缓存的 *Writer
    	writer := bufio.NewWriter(file)
    	for i := 0; i < 10; i++ {
    		writer.WriteString(str)
    	}
    	//因为writer是带缓存，因此在调用WriterString方法时，其实
    	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
    	//真正写入到文件中， 否则文件中会没有数据!!!
    	writer.Flush()
    
    }
    ```
- 打开一个存在的文件，在原来的内容追加内容 "ABC,ENGLISH!"

    ```go
    func main() {
    
    	//打开一个存在的文件，在原来的内容追加内容 'ABC! ENGLISH!'
    	//1 .打开文件已经存在文件, d:/abc.txt
    	filePath := "d:/abc.txt"
    	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
    	if err != nil {
    		fmt.Printf("open file err=%v\n", err)
    		return
    	}
    	//及时关闭file句柄
    	defer file.Close()
    	//准备写入5句
    	str := "ABC,ENGLISH!\r\n" // \r\n 表示换行
    	//写入时，使用带缓存的 *Writer
    	writer := bufio.NewWriter(file)
    	for i := 0; i < 10; i++ {
    		writer.WriteString(str)
    	}
    	//因为writer是带缓存，因此在调用WriterString方法时，其实
    	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
    	//真正写入到文件中， 否则文件中会没有数据!!!
    	writer.Flush()
    
    }
    ```
- 打开一个存在的文件，将原来的内容读出显示在终端，并追加5句 "hello, 北京"
    ```go
    func main() {
    
    	//打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句"hello,北京!"
    	//1 .打开文件已经存在文件, d:/abc.txt
    	filePath := "d:/abc.txt"
    	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
    	if err != nil {
    		fmt.Printf("open file err=%v\n", err)
    		return
    	}
    	//及时关闭file句柄
    	defer file.Close()
    
    	//先读取原来文件的内容，并显示在终端.
    	reader := bufio.NewReader(file)
    	for {
    		str, err := reader.ReadString('\n')
    		if err == io.EOF { //如果读取到文件的末尾
    			break
    		}
    		//显示到终端
    		fmt.Print(str)
    	}
    
    	//准备写入5句
    	str := "hello,北京!\r\n" // \r\n 表示换行
    	//写入时，使用带缓存的 *Writer
    	writer := bufio.NewWriter(file)
    	for i := 0; i < 5; i++ {
    		writer.WriteString(str)
    	}
    	//因为writer是带缓存，因此在调用WriterString方法时，其实
    	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
    	//真正写入到文件中， 否则文件中会没有数据!!!
    	writer.Flush()
    
    }
    ```
#### 基本应用实例-方式二
编写一个程序，将一个文件的内容，写入到另一个文件。
    ```go
    func main() {
    	//将d:/abc.txt 文件内容导入到  e:/kkk.txt
    	//1. 首先将  d:/abc.txt 内容读取到内存
    	//2. 将读取到的内容写入 e:/kkk.txt
    	file1Path := "d:/abc.txt"
    	file2Path := "e:/kkk.txt"
    	data, err := ioutil.ReadFile(file1Path)
    	if err != nil {
    		//说明读取文件有错误
    		fmt.Printf("read file err=%v\n", err)
    		return
    	}
    	err = ioutil.WriteFile(file2Path, data, 0666)
    	if err != nil {
    		fmt.Printf("write file error=%v\n", err)
    	}
    }

    ```

#### 判断文件是否存在
golang判断文件或文件夹是否存在的方法为使用os.Stat()函数返回的错误值进行判断
- 如果返回的错误为nil, 说明文件或文件夹存在
- 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
- 如果返回的错误为其他类型，则不确定是否再存在
```go


```
### 文件编程应用案例
#### 拷贝文件
func Copy(dst Writer, src Reader) (writer int64, err error)
```go
//自己编写一个函数，接收两个文件路径 srcFileName dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	//通过srcfile ,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	//通过dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)

}

func main() {

	//将d:/flower.jpg 文件拷贝到 e:/abc.jpg

	//调用CopyFile 完成文件拷贝
	srcFile := "d:/flower.jpg"
	dstFile := "e:/abc.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Printf("拷贝完成\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}

}
```
#### 统计英文、数字、空格和其他字符数量
```go

//定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int // 记录英文个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCount int // 记录其它字符的个数
}

func main() {

	//思路: 打开一个文件, 创一个Reader
	//每读取一行，就去统计该行有多少个 英文、数字、空格和其他字符
	//然后将结果保存到一个结构体
	fileName := "e:/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	//定义个CharCount 实例
	var count CharCount
	//创建一个Reader
	reader := bufio.NewReader(file)

	//开始循环的读取fileName的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //读到文件末尾就退出
			break
		}
		//遍历 str ，进行统计
		for _, v := range str {

			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	//输出统计的结果看看是否正确
	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)

}

```
### 命令行参数
#### 基本介绍
os.Args 是一个 string的切片，用来存储所有的命令行参数

```bash
test.exe -u -root -pwd 123456 -h 192.168.1.15 -port 3306
```

```go
func main() {

	fmt.Println("命令行的参数有", len(os.Args))
	//遍历os.Args切片，就可以得到所有的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}
```
### json 基本介绍
https//www.json.cn/ 网站可以验证json格式
#### json的序列化
json序列化是指，将有key-value结构的数据类型(结构体、map、切片)序列化成json字符串的操作。
介绍一下结构体、map和切片的序列化，其他数据类型的序列化类似
```go
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数退出时，要及时的关闭file
	defer file.Close() //要及时关闭file句柄，否则会有内存泄漏.

	// 创建一个 *Reader  ，是带缓冲的
	/*
		const (
		defaultBufSize = 4096 //默认的缓冲区为4096
		)
	*/
	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {                  // io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Println(str)
	}

	fmt.Println("文件读取结束...")
}
```
#### json的反序列化
json反序列化是指，将json字符串反序列化成对应的数据类型
```go
//定义一个结构体
type Monster struct {
	Name     string
	Age      int
	Birthday string //....
	Sal      float64
	Skill    string
}

//演示将json字符串，反序列化成struct
func unmarshalStruct() {
	//说明str 在项目开发中，是通过网络传输获取到.. 或者是读取文件获取到
	str := "{\"Name\":\"牛魔王~~~\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"

	//定义一个Monster实例
	var monster Monster

	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v monster.Name=%v \n", monster, monster.Name)

}

//将map进行序列化
func testMap() string {
	//定义一个map
	var a map[string]interface{}
	//使用map,需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿~~~~~~"
	a["age"] = 30
	a["address"] = "洪崖洞"

	//将a这个map进行序列化
	//将monster 序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	//fmt.Printf("a map 序列化后=%v\n", string(data))
	return string(data)

}

//演示将json字符串，反序列化成map
func unmarshalMap() {
	//str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	str := testMap()
	//定义一个map
	var a map[string]interface{}

	//反序列化
	//注意：反序列化map,不需要make,因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 a=%v\n", a)

}

//演示将json字符串，反序列化成切片
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"

	//定义一个slice
	var slice []map[string]interface{}
	//反序列化，不需要make,因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}

func main() {

	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}

```

## chapter15 单元测试

### 单元测试

####单元测试介绍

go语言中自带有一个轻量级的测试框架testing 和自带的 go test 命令来实现单元测试的和性能测试，testing框架和其他语言中的测试框架框架类似，可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例。

通过单元测试可以解决如下问题

- 确保每个函数时可运行的，并且运行结果是正确的
- 确保写出来的代码性能是好的
- 单元测试能及时的发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决，而性能测试的重点在于发现程序设计上的一些问题，让程序能够在高并发的情况下还能保持稳定。





#### 单元测试快速入门

![image-20220316002048219](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220316002048219.png)

- 测试用例文件名必须以 _test.go 结尾。比如 cal_test.go, cal 不是固定的

- 测试用例函数必须以 Test 开头，一般来说就是 Test + 被测试的函数名，比如 TestAddUpper

- TestAddUpper(t *testing T) 的形参类型必须是 *testing.T

- 一个测试用例文件中，可以有多个测试用例函数

- 运行测试用例指令

  - go test [如果运行正确，无日志，错误时，会输出日志]
  - go test -v [运行正确或是错误，都输出日志]

- 当出现错误时，可以使用 t.Fatalf 来格式化输出错误信息，并退出程序

- t.Logf 方法可以输出相应的日志

- 测试用例函数，并没有放在main函数中

- PASS表示测试用例运行成功，FAIL表示测试用例运行失败

- 测试单个文件，一定要带上被测试的原文件

  - go test -v cal_test.go cal.go

- 测试单个方法

  go test -v test.run TestAddUser

#### 单元测试综合案例

```go
type Monster struct {
	Name string
	Age int
	Skill string
} 

//给Monster绑定方法Store, 可以将一个Monster变量(对象),序列化后保存到文件中

func (this *Monster) Store() bool {

	//先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	} 

	//保存到文件
	filePath := "d:/monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err =", err)
		return false
	}
	return true
}


// 给Monster绑定方法ReStore, 可以将一个序列化的Monster,从文件中读取，
// 并反序列化为Monster对象,检查反序列化，名字正确
func (this *Monster) ReStore() bool {

	//1. 先从文件中，读取序列化的字符串
	filePath := "d:/monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err =", err)
		return false
	}

	//2.使用读取到data []byte ,对反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("UnMarshal err =", err)
		return false
	}
	return true
}

```

```go
// 测试用例,测试 Store 方法
func TestStore(t *testing.T) {

	// 先创建一个Monster 实例
	monster := &Monster{
		Name : "红孩儿",
		Age :10,
		Skill : "吐火.",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误，希望为=%v 实际为=%v", true, res)
	}
	t.Logf("monster.Store() 测试成功!")
}

func TestReStore(t *testing.T) {

	// 测试数据是很多，测试很多次，才确定函数，模块..
	
	// 先创建一个 Monster 实例 ， 不需要指定字段的值
	var monster = &Monster{}
	res := monster.ReStore() 
	if !res {
		t.Fatalf("monster.ReStore() 错误，希望为=%v 实际为=%v", true, res)
	}

	// 进一步判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore() 错误，希望为=%v 实际为=%v", "红孩儿", monster.Name)
	}

	t.Logf("monster.ReStore() 测试成功!") 
}

```

## chatper16 goroutine 和 channel

### goroutine

#### goroutine 基本介绍

- 进程就是程序在操作系统中的一次执行过程，是系统进行资源分配和调度的基本单位
- 线程是进程的一个执行实例，是程序执行的最小单位，它是比进程更小的能独立运行的基本单位
- 一个进程可以创建核销毁多个线程，同一个进程中的多个线程可以并发执行
- 一个程序至少有一个进程，一个进程至少有一个线程

#### 程序、进程和线程的关系示意图

![image-20220316164450197](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220316164450197.png)

#### 并发和并行

- 多线程程序在单核上运行，就是并发
- 多线程程序在多核上运行，就是并行

![image-20220316164656353](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220316164656353.png)

并发：在一个cpu上，比如有10个线程，每个线程执行10毫秒(进行轮询操作)。

并行：在多个cpu上，比如有10个线程，每个线程执行10毫秒(各自在不同cpu上执行)，从某一个时间点看，也同时有10个线程在执行，这就是并行

#### Go协程和Go主线程

Go主线程，一个Go线程上，可以起多个协程，协程是轻量级的线程。

Go协程的特点

- 有独立的栈空间

- 共享程序堆空间

- 调度由用户控制

- 协程是轻量级的线程

  ![image-20220316165242532](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220316165242532.png)

#### goroutine 快速入门

入门案例

- 在主线程中，开启一个 goroutine，该协程每隔1秒输出 "hello,world"
- 在主线程中每一秒输出 "hello,golang", 输出10次后，退出程序
- 要求主线程和goroutine 同时执行
- 画出主线程和协程执行流程图

```go
// 在主线程中，开启一个goroutine,该协程每隔1秒输出 "helo,world"
// 在主线程中每隔一秒输出 “hello，golang",输出10次后，退出程序
// 要求主线程和goroutine同时执行

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main()  {

	go test() // 开启一个协程

	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
```

主线程和协程执行流程图

![image-20220316170555358](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220316170555358.png)

- 主线程是一个物理线程，直接作用在cpu上的。是重量级的，非常耗费资源。
- 协程从主线程开启的，是轻量级的线程，是逻辑态。对资源消耗非常小
- Golang 的协程机制是重要的特点，可以轻松的开启上万个协程。

### go routine 的调度模型

#### MPG模式基本介绍

MPG模式运行的状态1

![image-20220316171110495](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220316171110495.png)

- M：操作系统的主线程
- p：协程执行需要的上下文
- G：协程

![image-20220316171243173](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220316171243173.png)

- 当前程序有三个M，如果三个M都在cpu运行，就是并发，如果在不同的cpu运行就是并行
- M1、M2、M3正在执行一个G，M1的协程队列有三个，M2的协程队列有3个，M3协程队列有两个
- go的协程是轻量级的线程，是逻辑态的，Go可以容易的起上万个协程
- 其他程序 c/java多线程，往往是内核态的，比较重量级。

#### MPG模式运行的状态2



![image-20220316171736323](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220316171736323.png)

1. 分两个部分看
2. 原来的情况是M0 主线程正在执行G0协程，另外有三个协程在队列等待
3. 如果GO协程阻塞，比如读取文件或者数据库等
4. 这时就会创建M1主线程（也可能是从已有的线程池中取出M1），并且将等待的3个协程挂到M1下开始执行，M0的主线程下的GO仍然执行文件io的读写
5. 这样的MPG调度模式，可以既让G0执行，同时也不会让队列的其他协程一直阻塞，仍然可以并发/并行执行
6. 等到GO不阻塞了，M0会被放到空闲的主线程继续执行（从已有的线程池中取），同时G0又不会被唤醒。

#### 管道需求

- 使用 goroutine 来完成，效率高，但是会出现 并发/并行安全问题
- 这里就提出了不同 goroutine 如何通信的问题

```go
// 计算 1 - 200 的各个数的阶乘，并且把各个数的阶乘放入到map中
//
var (
   myMap = make(map[int] int, 10)
)

func test(n int) {
   res := 1
   for i := 1; i <= n; i++ {
      res *= i
   }
   myMap[n] = res
}

func main() {
   for i := 1; i <= 200; i++ {
      go test(i)
   }
   time.Sleep(time.Second)
   for k, v := range myMap {
      fmt.Println("map[%d]=%d", k, v)
   }
}
```

![image-20220318121945414](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220318121945414.png)

#### 全局互斥锁及解决资源竞争

- 因为没有对全局变量 m 加锁，因此会出现资源争夺问题，代码会出现错误，提示 concurrent map writes
- 解决方案：加入互斥锁

```go
// 计算 1 - 200 的各个数的阶乘，并且把各个数的阶乘放入到map中
//
var (
	myMap2 = make(map[int] int, 10)
	// 声明一个全局的互斥锁
	// lock 是一个全局的互斥锁
	// sync 是包：synchorinized 同步
	// Mutex: 是互斥
	lock sync.Mutex
)

func test2(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 加锁
	lock.Lock()
	myMap2[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ {
		go test2(i)
	}
	time.Sleep(time.Second)
	for k, v := range myMap2 {
		fmt.Printf("map[%d]=%d\n", k, v)
	}
}

```

### channel

#### channel 介绍

- channel 本质就是一个数据结构
- 数据是先进先出
- 线程安全，多 goroutine访问时，不需要加锁，就是说 channel 本身就是线程安全的

![image-20220318122701785](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220318122701785.png)

#### 定义/声明channel

```go
var 变量名 chan 数据类型
```



- channel 是引用类型
- channel 必须初始化才能写入数据，即make 后才能使用
- 管道是有类型的， intChan 只能是写入整数int

```go
func main() {
	// 演示管道的使用
	// 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	// 查看intChan
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p \n", intChan, &intChan)

	// 管道内写入数据
	intChan<- 10
	num := 211
	intChan<- num

	// 管道写入数据，不能超过容量

	// 查看管道的长度和容量
	fmt.Printf("channel len= %v cap= %v \n", len(intChan), cap(intChan))

	// 从管道中读取数据
	num2 := <-intChan
	fmt.Printf("num2= %v\n", num2)
	// 在没有使用协程的情况下，如果管道数据已经全部取出，再取出就会报告 deadlock

	num3 := <-intChan
	fmt.Printf("num3= %v\n", num3)
	num4 := <-intChan
	fmt.Printf("num4= %v\n", num4)
}
```



![image-20220318124430437](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220318124430437.png)

## chapter TCP编程

### 网络编程基础介绍

Golang的主要设计目标之一就是面向大规模后端服务程序，网络通信这块是服务端，也是程序必不可少和至关重要的一部分。

网络编程分两种

1. TCP socket编程，是网络编程的主流。之所以叫 Tcp Socket编程，是因为底层是基于tcp/ip协议的
2. b/s 结构的htp编程，我们使用浏览器去访问服务器时
