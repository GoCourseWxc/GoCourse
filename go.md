

## chapter02
### go 语言 转义字符 （escape char)

| 转义字符 | 含义                       |
| -------- | -------------------------- |
| \t       | 一个制表位，实现对齐的功能 |
| \n       | 换行符                     |
| `\\`     | 一个\                      |
| `\"`     | 一个"                      |
| \r       | 回车符                     |

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

![image-20220306173822951](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220306173822951.png)

- go 变量使用的三种方式
  - 指定变量类型，声明后若不赋值，使用默认值
  - 根据值自行判断变量类型(类型推导)
  - 省略 var，:= 左侧的变量不应该是已经声明过的，否则会导致编译错误
- 多变量声明
- 该区域的数据值可以在同一个类型范围内不断变化
- 变量在同一个作用域(在一个函数或者代码块) 内不能重名
- 变量 = 变量名 + 值 + 数据类型
- Goland的变量如果没有赋初始值，编译器会使用默认值，比如int默认值0，string默认值为空串，小数默认为0

#### 程序中 + 号的使用

- 当左右两边都是数值型时，则做加法运算
- 当左右两边都是字符串，则做字符串拼接

### 数据类型的基本介绍

![image-20220306175051033](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220306175051033.png)

### 整数类型

#### 整数的各个类型

![image-20220306175157902](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220306175157902.png)

#### int 的 其他类型

![image-20220306175241823](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220306175241823.png)

#### 整型的使用细节

- Golang 各整数类型分: 有符号和无符号， int uint 的大小和系统有关

- Golang 的整型默认声明为 int 型

- 如何在程序查看某个变量的字节大小和数据类型

  ```go
  // unsafe.Sizeof(n1) 是unsafe包的一个函数，可以返回n1变量占用的字节数
  fmt.Printf("n2 的 类型 %T  n2占用的字节数是 %d ", n2, unsafe.Sizeof(n2))
  ```

  

- golang 程序中整型变量在使用时，遵守保小不保大的原则

- bit：计算机中的最小的存储单位。byte：计算机中基本存储单元。

### 小数类型/浮点型

#### 小数类型分类

![image-20220306180336279](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220306180336279.png)

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

![image-20220308220859379](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220308220859379.png)

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

  ![image-20220308225300989](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220308225300989.png)

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

- ![image-20220308225600230](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220308225600230.png)

- 案例
  - 写一个程序，获取一个int变量num的地址，并显示到终端
  - 将num的地址赋给指针 ptr， 并通过ptr去修改num值

​	![image-20220308230016461](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220308230016461.png)

#### 指针的使用细节

- 值类型，都有对应的指针类型，形式为 *数据类型，比如 int 的对应改的指针就是 *int，float32 对应的指针类型就是 *float32 依次类推。
- 值类型包括：基本数据类型 int系列， float类型， bool，string， 数组和结构体 struct。

### 值类型和引用类型

#### 值类型和引用类型说明

- 值类型：基本数据类型int系列，float系列，bool，string，数组和结构体struct
- 引用类型：指针、slice切片、map、管道chan、interface等都是引用类型

#### 值类型和引用类型的使用特点

- 值类型：变量直接存储值，内存通常在栈中分配

  ![image-20220308233124158](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220308233124158.png)

- 引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据(值)，内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC来回收。

  ![image-20220308233328717](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220308233328717.png)

- 内存的栈区和堆区示意图

  ![image-20220308233412262](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220308233412262.png)

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

![image-20220308234315938](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220308234315938.png)

### 系统的预定义标识符

![image-20220308234404546](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220308234404546.png)

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

![image-20220308234715256](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220308234715256.png)

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

![image-20220308235648311](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220308235648311.png)

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

![image-20220312131530711](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312131530711.png)



![image-20220312131547001](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312131547001.png)

#### 赋值运算符的特点

- 运算顺序从右到左
- 赋值运算符的左边，只能是变量，右边是变量、表达式、常量值
- 复合赋值运算符等价效果 a += 3  -> a = a + 3

### 位运算符

![image-20220312131818293](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312131818293.png)

### 其他运算符

![image-20220312132126219](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312132126219.png)

### 运算符

![image-20220312132330195](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312132330195.png)



### 键盘输入语句

- 导入fm包
- 调用 ftm 包的 fmt.Scanln() 或者 fmt.Scanf()

![image-20220312132918612](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312132918612.png)



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

![image-20220309000904976](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220309000904976.png)

### 分支流程

分支流程的三种形式

- 单分支
- 双分支
- 多分支

#### 单分支

单分支流程图

![image-20220309001213911](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220309001213911.png)

```go
// golang支持在if中，直接定义一个变量，比如下面
if age := 20; age > 18 {
    fmt.Println("你年龄大于18,要对自己的行为负责!")
}
```

#### 双分支

双分支流程图

![image-20220309001320399](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220309001320399.png)

#### 多分支

多分支流程图

![image-20220309001405243](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220309001405243.png)

#### 嵌套分支

在一个分支结构中又嵌套另一完整的分支结构，里面的内层分支，外面的分支结构为外层分支。

#### switch分支控制

switch语句用于基于不同条件执行不同动作，每一个case分支都是唯一的，从上到下逐一测试直到匹配为止，不需要添加break。

switch分支路程图

![image-20220309001934071](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220309001934071.png)

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

![image-20220312122042873](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312122042873.png)

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

![image-20220312134612405](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312134612405.png)

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

![image-20220312135200442](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312135200442.png)

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

![image-20220312141337162](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312141337162.png)

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

![image-20220312142119223](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312142119223.png)

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

-  函数中的变量是局部的。

- 基本数据类型和数组默认都是值传递的，即进行值拷贝。

- 函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量

- go函数不支持函数重载

- 在go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量，通过该变量可以对函数调用

- 函数既然是一种数据类型，在go中，函数可以作为形参

- go支持自定义数据类型

  基本语法: type 自定义数据类型名 数据类型  // 相当于别名

- 支持对函数返回值命名

- 使用 _ 标识符，忽略返回值

- go 支持可变参数

#### init 函数

每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被Go运行框架调用，init会在main函数之前被调用

##### init函数使用细节

- 如果一个文件同时包含全局变量定义，init函数和main函数，则执行的流程 全局变量定义 -> init 函数 -> main函数
- init函数最主要的作用，就是完成一些初始化的工作
- main.go 和 utils.go 中都有变量定义和函数初始化

![image-20220312145015943](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312145015943.png)

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

![image-20220312222204301](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312222204301.png)

引用类型默认是引用传递：变量存储的是一个地址，这个地址对应的空间才真正存储数据(值)，内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC来回收。

![image-20220312222403754](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312222403754.png)

如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量。

![image-20220312222533141](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312222533141.png)

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

  ![image-20220312215551787](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312215551787.png)

  ```go
  // Unix和UnixNano的使用
  fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano())
  // unix时间戳=1647093474 unixnano时间戳=1647093474377446500
  ```

#### 内置函数

golang设计者为了编程方便，提供了一些函数，这些函数可以直接使用，这些称为go的内置函数。

- len：用来求长度，比如string、array、slice、map、channel

- new：用来分配内存，主要用来分配值类型，比如int、float32, struct... 返回的是指针

  ![image-20220312220615981](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312220615981.png)

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

![image-20220312223836189](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220312223836189.png)

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
for index, value := range array01 {
    
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

![image-20220313112337678](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220313112337678.png)

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

![image-20220313113212850](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220313113212850.png)

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

  ![image-20220313115635260](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220313115635260.png)

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

  ![image-20220313120221213](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220313120221213.png)

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
![image-20220313122314503](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220313122314503.png)

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

![image-20220313123421996](C:\Users\ａｄｍｉｎ\AppData\Roaming\Typora\typora-user-images\image-20220313123421996.png)

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

使用
