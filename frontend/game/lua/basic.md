
### 什么是 Lua
Lua 是巴西里约热内卢天主教大学（Pontifical Catholic University of Rio de Janeiro）里的一个研究小组，由 Roberto Ierusalimschy、Waldemar Celes 和 Luiz Henrique de Figueiredo 所组成并于 1993 年开发。  
Lua 是一种轻量小巧的脚本语言，用标准 C 语言编写并以源代码形式开放，其设计目的是为了嵌入应用程序中，从而为应用程序提供灵活的扩展和定制功能。

**Lua 特性**  

- 轻量级: 它用标准 C 语言编写并以源代码形式开放，编译后仅仅一百余 K，可以很方便的嵌入别的程序里。
- 可扩展: Lua 提供了非常易于使用的扩展接口和机制：由宿主语言 (通常是 C 或 C++) 提供这些功能，Lua 可以使用它们，就像是本来就内置的功能一样。
- 其它特性:
支持面向过程 (procedure-oriented) 编程和函数式编程 (functional programming)；  
自动内存管理；只提供了一种通用类型的表（table），用它可以实现数组，哈希表，集合，对象；  
语言内置模式匹配；闭包 (closure)；函数也可以看做一个值；提供多线程（协同进程，并非操作系统所支持的线程）支持；  
通过闭包和 table 可以很方便地支持面向对象编程所需要的一些关键机制，比如数据抽象，虚函数，继承和重载等。  

**Lua 应用场景**  

- 游戏开发
- 独立应用脚本
- Web 应用脚本
- 扩展和数据库插件如：MySQL Proxy 和 MySQL WorkBench
- 安全系统，如入侵检测系统

### Lua 环境搭建
Linux & Mac 上安装 Lua 安装非常简单，只需要下载源码包并在终端解压编译即可。  
```bash
curl -R -O http://www.lua.org/ftp/lua-5.3.5.tar.gz
tar zxf lua-5.3.5.tar.gz
cd lua-5.3.5
make linux test # make macosx test
make install
```

window 下，可以使用一个叫 "SciTE" 的 IDE 环境来执行 lua 程序，下载地址在[这里](https://github.com/rjpcomputing/luaforwindows/releases)。

### Lua 基础
Lua 的两种编程方式：交互式编程和脚本式编程。  
```bash
# 交互式编程
$ lua -i
> print("Hello, Lua")
```

脚本式编程可以将 Lua 程序代码保存到一个以 lua 结尾的文件，并执行。  
```bash
lua hello.lua
```

也可以在文件中声明解释器，这时候执行文件时就可以不用加 lua。
```lua
#!/usr/local/bin/lua

print("Hello, Lua")
```

**单行注释**  
两个减号是单行注释。  
```lua
-- 这是单行注释
```

**多行注释**  
```lua
--[[
   多行注释
   多行注释
--]]
```

**关键字**  
Lua 的保留关键字，它们不能作为常量或变量或其他用户自定义标示符：  
and、break、do、else  
elseif、end、false、for  
function、if、in、local  
nil、not、or、repeat  
return、then、true、until  
while			 

**标识符与变量**  
Lua 的标识符用于定义一个变量，函数获取其他用户定义的项。  
标示符以一个字母 A 到 Z 或 a 到 z 或下划线 _ 开头后加上 0 个或多个字母，下划线，数字（0 到 9）。Lua 不允许使用特殊字符如 @, $, 和 % 来定义标示符；Lua 是一个区分大小写的编程语言。

`建议：不要使用下划线加大写字母的标示符，因为 Lua 的保留字也是这样的。`

Lua 变量有三种类型：全局变量、局部变量、表中的域。  
在默认情况下，变量总是认为是全局的，除非使用 local 声明；函数内变量与函数的参数默认为局部变量（局部变量的作用域为从声明位置开始到所在语句块结束）。  
全局变量不需要声明，给一个变量赋值后即创建了这个全局变量，访问一个没有初始化的全局变量也不会出错，只不过得到的结果是：nil。  

`注意：应尽可能的使用局部变量，这样可以避免命名冲突，而且访问局部变量的速度比全局变量快。`

```lua
print(b) -- nil
b = 1    -- 全局变量
print(b) -- 1
b = nil  -- 删除全局变量

local b = 5 -- 局部变量

function joke()
    c = 5           -- 全局变量
    local d = 6     -- 局部变量
end
joke()
print(c,d)          -- 5 nil

-- 变量赋值
a = "hello" .. "world"
x = 2
y = 3
-- 交换变量
x, y = y, x 
-- 多变量赋值
a, b = 10, 2*x 
a, b = func()
-- 当变量个数和值的个数不一致时，Lua 会一直以变量个数为基础采取以下策略：
-- 变量个数 > 值的个数，按变量个数补足 nil
-- 变量个数 < 值的个数，多余的值会被忽略
a, b, c = 0, 1
print(a,b,c) -- 0   1   nil
a, b = a+1, b+1, b+2
print(a,b) --> 1   2
```

**数据类型**  
Lua 是动态类型语言，变量不要类型定义，只需要为变量赋值（值可以存储在变量中，作为参数传递或结果返回）。  
Lua 中有 8 个基本类型分别为：nil、boolean、number、string、userdata、function、thread 和 table。  

使用 type 函数测试给定变量或者值的类型。
```lua
print(type(true))  -- boolean
```

nil 类型表示一种没有任何有效值，它只有一个值（即 nil）。  
对于全局变量和 table，nil 还有一个“删除”作用：给全局变量或者 table 表里的变量赋一个 nil 值，等同于把它们删掉。  
```lua
print(type(a))

tab1 = { key1 = "val1", key2 = "val2", "val3" }
for k, v in pairs(tab1) do
    print(k .. " - " .. v)
end

tab1.key1 = nil
for k, v in pairs(tab1) do
    print(k .. " - " .. v)
end
```

boolean 类型只有两个可选值：true（真） 和 false（假），Lua 把 false 和 nil 看作是” 假”，其他的都为” 真”。 
```lua
if type(false) or type(nil) then
    print("false and nil are false")
else
    print("other is true")
end
``` 

Lua 默认只有一种 number 类型 —— double（双精度）类型，有以下几种写法。  
```lua
print(type(2))
print(type(2.2))
print(type(0.2))
print(type(2e+1))
print(type(0.2e-1))
print(type(7.8263692594256e-06))
```

Lua 字符串由一对双引号或单引号来表示，也可以用 2 个方括号 “[[]]” 来表示” 一块” 字符串。  
```lua
string1 = "this is string1"
string2 = 'this is string2'

html = [[
<html>
<head></head>
<body>
    <a href="//www.w3cschool.cn/">w3cschoolW3Cschool教程</a>
</body>
</html>
]]
print(html)

-- 对一个数字字符串上进行算术操作时，Lua 会尝试将这个数字字符串转成一个数字
print("2" + 6) -- 8.0
print("2" + "6") -- 8.0
print("2 + 6") -- 2 + 6
print("-2e2" * "6") -- -100
print("error" + 1) -- 报错

-- 字符串连接
print("a" .. 'b')
print(157 .. 428)

-- 使用 # 来计算字符串的长度
str = "ihaveadream"
print(#str)
print("字符串长度还可以这样：",string.len(str)) 

-- 大小写进行转换
string = "Lua Tutorial"
print(string.upper(str)) -- LUA TUTORIAL
print(string.lower(str)) -- lua tutorial
-- 查找字符串，模式匹配
print(string.find(string,"Tutorial")) -- 5  12
reversedString = string.reverse(string) 
print("反转后的字符串为：",reversedString) -- lairotuT auL

string1 = "Lua"
string2 = "Tutorial"
number1 = 10
number2 = 20
-- 基本字符串格式化
print(string.format("基本格式化 %s %s",string1,string2)) -- 基本格式化 Lua Tutorial
-- 日期格式化
date = 2; month = 1; year = 2014
print(string.format("日期格式化 %02d/%02d/%03d", date, month, year)) -- 日期格式化 02/01/2014
-- 十进制格式化
print(string.format("%.4f",1/3)) -- 0.3333

-- 字符转换
-- 转换第一个字符
print(string.byte("Lua")) -- 76
-- 转换第三个字符
print(string.byte("Lua",3)) -- 97
-- 转换末尾第一个字符
print(string.byte("Lua",-1)) -- 97
-- 第二个字符
print(string.byte("Lua",2)) -- 117
-- 转换末尾第二个字符
print(string.byte("Lua",-2)) -- 117
 -- 整数 ASCII 码转换为字符
print(string.char(97)) -- a

-- 字符串复制 2 次
str = "Lua"
repeatedString = string.rep(str,2) -- LuaLua

-- 截取字符串
local s = "[in brackets]"
print(string.sub(s,2,-1))  -- in brackets]
```

在 Lua 里，表（table）是一种用来帮助我们创建不同的数据类型的数据结构，如：数字、字典等，它的创建是通过” 构造表达式” 来完成，最简单构造表达式是 {}，用来创建一个空表；也可以在表里添加一些数据，直接初始化表。  
Lua 中的表（table）其实是一个” 关联数组”（associative arrays），数组的索引可以是数字或者是字符串（但不能为 nil）。  
Lua table 是不固定大小的，你可以根据自己需要进行扩容。  
Lua 也是通过 table 来解决模块（module）、包（package）和对象（Object）的。 例如 string.format 表示使用”format” 来索引 table string。  
```lua
-- 创建一个空的 table
local tbl1 = {}

-- 直接初始表
local tbl2 = {"apple", "pear", "orange", "grape"}

a = {}
a["key"] = "value"
key = 10
a[key] = 22
a[key] = a[key] + 11
for k, v in pairs(a) do
    print(k .. " : " .. v)
end
-- key : value
-- 10 : 33

-- 注意：不同于其他语言的数组把 0 作为数组的初始索引，在 Lua 里表的默认初始索引一般以 1 开始
local tbl = {"apple", "orange"}
for key, val in pairs(tbl) do
    print("key", key)
end
-- key 1
-- key 2

-- table 不会固定长度大小，有新数据添加时 table 长度会自动增长，没初始的 table 都是 nil
a3 = {}
for i = 1, 10 do
    a3[i] = i
end
a3["key"] = "val"
print(a3["key"]) -- val
print(a3["none"]) -- nil

-- 初始化多维 table
array = {}
for i=1,3 do
   array[i] = {}
      for j=1,3 do
         array[i][j] = i*j
      end
end
-- 访问多维 table
for i=1,3 do
   for j=1,3 do
      print(array[i][j])
   end
end

-- table 操作

fruits = {"banana","orange","apple"}
-- 返回 table 连接后的字符串
print("连接后的字符串 ",table.concat(fruits)) -- 连接后的字符串     bananaorangeapple
-- 指定连接字符
print("连接后的字符串 ",table.concat(fruits,", ")) -- 连接后的字符串     banana, orange, apple
-- 指定索引来连接 table
print("连接后的字符串 ",table.concat(fruits,", ", 2,3)) -- 连接后的字符串     orange, apple

fruits = {"banana","orange","apple"}
-- 在末尾插入
table.insert(fruits,"mango")
print("索引为 4 的元素为 ",fruits[4]) -- 索引为 4 的元素为  mango
-- 在索引为 2 的键处插入
table.insert(fruits,2,"grapes")
print("索引为 2 的元素为 ",fruits[2]) -- 索引为 2 的元素为   grapes
print("最后一个元素为 ",fruits[5]) -- 最后一个元素为     mango
table.remove(fruits)
print("移除后最后一个元素为 ",fruits[5]) -- 移除后最后一个元素为   nil

fruits = {"banana","orange","apple","grapes"}
print("排序前")
for k,v in ipairs(fruits) do
  print(k,v)
end
-- table 排序
table.sort(fruits)
print("排序后")
for k,v in ipairs(fruits) do
   print(k,v)
end

-- 最大键
function table_maxn(t)
    local mn = 0
    for k, v in pairs(t) do
        if mn < k then
            mn = k
        end
    end
    return mn
end
tbl = {[1] = "a", [2] = "b", [3] = "c", [26] = "z"}
print("tbl 长度 ", #tbl) -- tbl 长度    3
print("tbl 最大值 ", table_maxn(tbl)) -- tbl 最大值  26
```

在 Lua 中，函数是被看作是” 第一类值（First-Class Value）”，函数可以存在变量里。  
```lua
function factorial1(n)
    if n == 0 then
        return 1
    else
        return n * factorial1(n - 1)
    end
end
print(factorial1(5)) -- 120
factorial2 = factorial1
print(factorial2(5)) -- 120

-- 以匿名函数（anonymous function）的方式通过参数传递
function anonymous(tab, fun)
    for k, v in pairs(tab) do
        print(fun(k, v))
    end
end
tab = { key1 = "val1", key2 = "val2" }
anonymous(tab, function(key, val)
    return key .. " = " .. val
end)
-- key1 = val1
-- key2 = val2
```

在 Lua 里，最主要的线程是协同程序（coroutine），它跟线程（thread）差不多，拥有自己独立的栈、局部变量和指令指针，可以跟其他协同程序共享全局变量和其他大部分东西。  
`线程跟协程的区别：线程可以同时多个运行，而协程任意时刻只能运行一个，并且处于运行状态的协程只有被挂起（suspend）时才会暂停。`

userdata 是一种用户自定义数据，用于表示一种由应用程序或 C/C++ 语言库所创建的类型，可以将任意 C/C++ 的任意数据类型的数据（通常是 struct 和 指针）存储到 Lua 变量中调用。  

**运算符**  
运算符是一个特殊的符号，用于告诉解释器执行特定的数学或逻辑运算。  
Lua 提供了以下几种运算符类型：算术运算符、关系运算符、逻辑运算符、其他运算符。  

Lua 语言中的常用算术运算符，设定 A 的值为 10，B 的值为 20。  

| 操作符 | 描述 | 示例 |
| :----: | :----: | :----: |  
| +	| 加法 | A + B 输出结果 30 |  
| –	| 减法 | A – B 输出结果 -10 |  
| *	| 乘法 | A * B 输出结果 200 |  
| /	| 除法 | B / A 输出结果 2 |  
| %	| 取余 | B % A 输出结果 0 |  
| ^	| 乘幂 | A^2 输出结果 100 |  
| –	| 负号 | -A 输出结果 -10 |  

Lua 语言中的常用关系运算符，设定 A 的值为 10，B 的值为 2。  

| 操作符 | 描述 | 示例 |
| :----: | :----: | :----: |  
| == | 等于，检测两个值是否相等，相等返回 true，否则返回 false | (A == B) 为 false |  
| ~= | 不等于，检测两个值是否相等，相等返回 false，否则返回 true | (A ~= B) 为 true |  
| >	| 大于，如果左边的值大于右边的值，返回 true，否则返回 false | (A > B) 为 true |  
| < | 小于，如果左边的值大于右边的值，返回 false，否则返回 true | (A < B) 为 false |  
| >= | 大于等于，如果左边的值大于等于右边的值，返回 true，否则返回 false | (A >= B) 为 false |  
| <= | 小于等于， 如果左边的值小于等于右边的值，返回 true，否则返回 false | (A <= B) 为 true |  

Lua 语言中的常用逻辑运算符，设定 A 的值为 true，B 的值为 false。  

| 操作符 | 描述 | 示例 |
| :----: | :----: | :----: |  
| and | 逻辑与操作符，如果两边的操作都为 true 则条件为 true | (A and B) 为 false |  
| or | 逻辑或操作符，如果两边的操作任一个为 true 则条件为 true | (A or B) 为 true |  
| not | 逻辑非操作符，与逻辑运算结果相反，如果条件为 true，逻辑非为 false | not (A and B) 为 true |  

Lua 语言中的连接运算符与计算表或字符串长度的运算符。  

| 操作符 | 描述 | 示例 |
| :----: | :----: | :----: |  
| .. | 连接两个字符串 | a..b ，其中 a 为 “Hello” ， b 为 “World”, 输出结果为 “Hello World” |   
| # | 一元运算符，返回字符串或表的长度 | #”Hello” 返回 5 |   

**运算符优先级**  
从高到低的顺序（除了 ^ 和 .. 外所有的二元运算符都是左连接的，比如 a+i < b/2+1 等同于 (a+i) < ((b/2)+1)）：
```
^
not    -（负号）
*      /
+      -（减号）
..
<      >      <=     >=     ~=     ==
and
or
```

**流程控制**  
Lua 编程语言流程控制语句通过程序设定一个或多个条件语句来设定；在条件为 true 时执行指定程序代码，在条件为 false 时执行其他指定代码。
```lua
if 0 then
    print("0 为真") -- 除了 false 和 nil 都为真
end

if num > 40 then 
  print('over 40') 
elseif s ~= 'walternate' then  -- ~= 表示不等于
  io.write('not over 40\n')  -- 默认标准输出
else 
  -- 默认全局变量
  thisIsGlobal = 5  -- 通常使用驼峰

  -- 如何定义局部变量： 
  local line = io.read()  -- 读取标准输入的下一行

  -- ..操作符用于连接字符串： 
  print('Winter is coming, ' .. line) 
end 
```

**迭代器**  
迭代器（iterator）是一种对象，它能够用来遍历标准模板库容器中的部分或全部元素，每个迭代器对象代表容器中的确定的地址。  
在 Lua 中迭代器是一种支持指针类型的结构，它可以遍历集合的每一个元素。  

泛型 for 在自己内部保存迭代函数，实际上它保存三个值：迭代函数、状态常量、控制变量。
```lua
tab = {"Lua", "Tutorial"}

-- Lua 默认迭代函数 pairs/ipairs
for key,value in ipairs(tab) 
do
   print(key, value)
end
-- Lua
-- Tutorial
```

无状态的迭代器是指不保留任何状态的迭代器，因此在循环中我们可以利用无状态迭代器避免创建闭包花费额外的代价。  
每一次迭代，迭代函数都是用两个变量（状态常量和控制变量）的值作为参数被调用，一个无状态的迭代器只利用这两个值可以获取下一个元素。这种无状态迭代器的典型的简单的例子是 ipairs，他遍历数组的每一个元素。
```lua
-- 简单实现迭代器
function square(iteratorMaxCount,currentNumber)
   if currentNumber<iteratorMaxCount
   then
      currentNumber = currentNumber+1
   return currentNumber, currentNumber*currentNumber
   end
end

for i,n in square,3,0
do
   print(i,n)
end
-- 1   1
-- 2   4
-- 3   9
```

很多情况下，迭代器需要保存多个状态信息而不是简单的状态常量和控制变量，最简单的方法是使用闭包，还有一种方法就是将所有的状态信息封装到 table 内，将 table 作为迭代器的状态常量，因为这种情况下可以将所有的信息存放在 table 内，所以迭代函数通常不需要第二个参数。
```lua
-- 多状态的迭代器
tab = {"Lua", "Tutorial"}

function elementIterator (collection)
   local index = 0
   local count = #collection
   -- 闭包函数
   return function ()
      index = index + 1
      if index <= count
      then
         --  返回迭代器的当前元素
         return collection[index]
      end
   end
end

for element in elementIterator(tab)
do
   print(element)
end
```

**循环**  
循环结构是在一定条件下反复执行某段程序的流程结构，被反复执行的程序被称为循环体。  
```lua
karlSum = 0 
for i = 1, 100 do  -- 范围包含两端，默认步进为 1 
    karlSum = karlSum + i 
    print(karlSum)
end
-- 使用 "100, 1, -1" 表示递减的范围
fredSum = 0 
for j = 100, 1, -1 do 
    fredSum = fredSum + j 
    if j == 50 then
        break -- 跳出循环
    end

    print(fredSum)
end 

num = 10
repeat 
  print('the way of the future') 
  num = num - 1 
until num == 0 

-- 无限循环
while(true)
do
   print("循环将永远执行下去")
end
```

**函数**  
函数是对语句和表达式进行抽象的主要方法，既可以用来处理一些特殊的工作，也可以用来计算一些值。  
Lua 提供了许多的内建函数，你可以很方便的在程序中调用它们，如 print() 函数可以将传入的参数打印在控制台上。  
```lua
-- lua 函数格式
optional_function_scope function function_name( argument1, argument2, argument3..., argumentn)
    function_body
    return result_params_comma_separated
end
-- optional_function_scope 值默认为全局，local 声明局部

function max (num1, num2) 
	if (num1 > num2) then 
		result = num1
	else 
		result = num2
	end 

	return result
end

function fib(n)
  if n < 2 then return n end

  return fib(n - 2) + fib(n - 1)
end

-- 多个返回值
function bar() 
  return 4, 8, 15, 16, 23, 42 
end 

-- 支持闭包及匿名函数
function adder(x) 
  -- 调用adder时，会创建返回的函数，并且会记住 x 的值
  return function (y) return x + y end 
end 
a1 = adder(9) 
print(a1(16)) -- 25
```

**模块与包**  
模块类似于一个封装库，从 Lua 5.1 开始，Lua 加入了标准的模块管理机制，可以把一些公用的代码放在一个文件里，以 API 接口的形式在其他地方调用，有利于代码的重用和降低代码耦合度。  
Lua 的模块是由变量、函数等已知元素组成的 table，因此创建一个模块很简单，就是创建一个 table，然后把需要导出的常量、函数放入其中，最后返回这个 table 就行。  
```lua
-- 文件名为 module.lua
-- 定义一个名为 module 的模块
module = {}

-- 定义一个常量
module.constant = "这是一个常量"

-- 定义一个函数
function module.func1()
    io.write("这是一个公有函数！\n")
end

-- 声明为程序块的局部变量，即表示一个私有函数
local function func2()
    print("这是一个私有函数！")
end

function module.func3()
    func2()
end

return module
```

由上可见，模块的结构就是一个 table 的结构，因此可以像操作调用 table 里的元素那样来操作调用模块里的常量或函数。  

Lua 提供了一个名为 require 的函数用来加载模块。要加载一个模块，只需要简单地调用就可以了。  
```lua
-- test_module.php 文件
-- module 模块为上文提到到 module.lua
require("module") -- 或可定义别名变量 local m = require("module")

print(module.constant) -- 这是一个常量

module.func3() -- 这是一个私有函数！
```

对于自定义的模块，模块文件不是放在哪个文件目录都行，函数 require 有它自己的文件路径加载策略，它会尝试从 Lua 文件或 C 程序库中加载模块。  
require 用于搜索 Lua 文件的路径是存放在全局变量 package.path 中，当 Lua 启动后，会以环境变量 LUA_PATH 的值来初始这个环境变量。如果没有找到该环境变量，则使用一个编译时定义的默认路径来初始化。当然，如果没有 LUA_PATH 这个环境变量，也可以自定义设置，在当前用户根目录下打开 .profile 文件（没有则创建，打开 .bashrc 文件也可以），例如把 “~/lua/” 路径加入 LUA_PATH 环境变量里：  
```
# 文件路径以 “;” 号分隔，最后的 2 个 “;;” 表示新加的路径后面加上原来的默认路径
export LUA_PATH="~/lua/?.lua;;"
```
如果找过目标文件，则会调用 package.loadfile 来加载模块。否则，就会去找 C 程序库。搜索的文件路径是从全局变量 package.cpath 获取，而这个变量则是通过环境变量 LUA_CPATH 来初始。  

Lua 和 C 是很容易结合的，使用 C 为 Lua 写包。与 Lua 中写包不同，C 包在使用以前必须首先加载并连接，在大多数系统中最容易的实现方式是通过动态连接库机制。  
Lua 在一个叫 loadlib 的函数内提供了所有的动态连接的功能。这个函数有两个参数：库的绝对路径和初始化函数。  
```lua
local path = "/usr/local/lua/lib/libluasocket.so"
local f = loadlib(path, "luaopen_socket")
```
loadlib 函数加载指定的库并且连接到 Lua，然而它并不打开库（也就是说没有调用初始化函数），反之他返回初始化函数作为 Lua 的一个函数，这样我们就可以直接在 Lua 中调用他。如果加载动态库或者查找初始化函数时出错，loadlib 将返回 nil 和错误信息。  
```lua
local path = "/usr/local/lua/lib/libluasocket.so"
-- 或者 path = "C:\\windows\\luasocket.dll"，这是 Window 平台下
local f = assert(loadlib(path, "luaopen_socket"))
f()  -- 真正打开库
```
一般情况下，我们期望二进制的发布库包含一个与前面代码段相似的 stub 文件，安装二进制库的时候可以随便放在某个目录，只需要修改 stub 文件对应二进制库的实际路径即可。将 stub 文件所在的目录加入到 LUA_PATH，这样设定后就可以使用 require 函数加载 C 库了。  

**元表（Metatable）**  
在 Lua table 中我们可以访问对应的 key 来得到 value 值，但是却无法对两个 table 进行操作。因此，Lua 提供了元表 (Metatable)，允许我们改变 table 的行为，每个行为关联了对应的元方法。  
元表可以很好的简化我们的代码功能，了解 Lua 的元表可以让我们写出更加简单优秀的 Lua 代码。  

有两个很重要的函数来处理元表：  

* setmetatable(table,metatable): 对指定 table 设置元表 (metatable)，如果元表 (metatable) 中存在__metatable 键值，setmetatable 会失败 。  
* getmetatable(table): 返回对象的元表 (metatable)。

\_\_index 元方法是 metatable 最常用的键。  
当你通过键来访问 table 的时候，如果这个键没有值，那么 Lua 就会寻找该 table 的 metatable（假定有 metatable）中的__index 键。如果__index 包含一个表格，Lua 会在表格中查找相应的键。如果__index 包含一个函数的话，Lua 就会调用那个函数，table 和键会作为参数传递给函数。  
```lua
mytable = setmetatable({key1 = "value1"}, { __index = { key2 = "metatablevalue" } })
print(mytable.key1,mytable.key2)
```

\_\_newindex 元方法用来对表更新，\_\_index 则用来对表访问。  
当你给表的一个缺少的索引赋值，解释器就会查找__newindex 元方法：如果存在则调用这个函数而不进行赋值操作。  
```lua
mytable = setmetatable({key1 = "value1"}, {
  __newindex = function(mytable, key, value)
       rawset(mytable, key, "\""..value.."\"")

  end
})

mytable.key1 = "new value"
mytable.key2 = 4

print(mytable.key1,mytable.key2)
```

\_\_add 键包含在元表中，并进行相加操作。  
```lua
-- 计算表中最大值，table.maxn 在 Lua5.2 以上版本中已无法使用
-- 自定义计算表中最大值函数 table_maxn
function table_maxn(t)
    local mn = 0
    for k, v in pairs(t) do
        if mn < k then
            mn = k
        end
    end
    return mn
end

-- 两表相加操作
mytable = setmetatable({ 1, 2, 3 }, {
  __add = function(mytable, newtable)
    for i = 1, table_maxn(newtable) do
      table.insert(mytable, table_maxn(mytable)+1,newtable[i])
    end
    return mytable
  end
})

secondtable = {4,5,6}

mytable = mytable + secondtable
    for k,v in ipairs(mytable) do
print(k,v)
end
```

为表添加操作符，对应操作列表如下。  

| 模式 | 描述 |  
| :----: | :----: |  
| __add	| 对应的运算符 ‘+’ |  
| __sub	| 对应的运算符 ‘-‘ |  
| __mul	| 对应的运算符 ‘*’ |  
| __div	| 对应的运算符 ‘/’ |  
| __mod	| 对应的运算符 ‘%’ |  
| __unm	| 对应的运算符 ‘-‘ |  
| __concat | 对应的运算符 ‘..’ |  
| __eq	| 对应的运算符 ‘==’ |  
| __lt	| 对应的运算符 ‘<‘ |  
| __le	| 对应的运算符 ‘<=’ |  


\_\_call 元方法在 Lua 调用一个值时调用。  
```lua
-- 计算表中最大值，table.maxn在Lua5.2以上版本中已无法使用
-- 自定义计算表中最大值函数 table_maxn
function table_maxn(t)
    local mn = 0
    for k, v in pairs(t) do
        if mn < k then
            mn = k
        end
    end
    return mn
end

-- 定义元方法__call
mytable = setmetatable({10}, {
  __call = function(mytable, newtable)
    sum = 0
    for i = 1, table_maxn(mytable) do
        sum = sum + mytable[i]
    end
    for i = 1, table_maxn(newtable) do
        sum = sum + newtable[i]
    end
    return sum
  end
})
newtable = {10,20,30}
print(mytable(newtable)) -- 70
```

\_\_tostring 元方法用于修改表的输出行为。  
```lua
mytable = setmetatable({ 10, 20, 30 }, {
  __tostring = function(mytable)
    sum = 0
    for k, v in pairs(mytable) do
        sum = sum + v
 end
    return "表所有元素的和为 " .. sum
  end
})

print(mytable) -- 表所有元素的和为 60
```

**协同程序（coroutine）**  
Lua 协同程序 (coroutine) 与线程比较类似：拥有独立的堆栈，独立的局部变量，独立的指令指针，同时又与其它协同程序共享全局变量和其它大部分东西。  
线程与协同程序的主要区别在于，一个具有多个线程的程序可以同时运行几个线程，而协同程序却需要彼此协作的运行。在任一指定时刻只有一个协同程序在运行，并且这个正在运行的协同程序只有在明确的被要求挂起的时候才会被挂起。协同程序有点类似同步的多线程，在等待同一个线程锁的几个线程有点类似协同。  


| 方法 | 描述 |  
| ---: | :--- |  
| coroutine.create() | 创建 coroutine，返回 coroutine， 参数是一个函数，当和 resume 配合使用的时候就唤醒函数调用 |   
| coroutine.resume() | 重启 coroutine，和 create 配合使用 |  
| coroutine.yield() | 挂起 coroutine，将 coroutine 设置为挂起状态，这个和 resume 配合使用能有很多有用的效果 |  
| coroutine.status() | 查看 coroutine 的状态（`注：coroutine 的状态有三种：dead，suspend，running，具体什么时候有这样的状态请参考下面的程序`） |   
| coroutine.wrap（） | 	创建 coroutine，返回一个函数，一旦你调用这个函数，就进入 coroutine，和 create 功能重复 |  
| coroutine.running() | 返回正在跑的 coroutine，一个 coroutine 就是一个线程，当使用 running 的时候，就是返回一个 corouting 的线程号 |   

coroutine 在底层实现就是一个线程，当 create 一个 coroutine 的时候就是在新线程中注册了一个事件；当使用 resume 触发事件的时候，create 的 coroutine 函数就被执行了，当遇到 yield 的时候就代表挂起当前线程，等候再次 resume 触发事件。  
```lua
function foo (a)
    print("foo 函数输出", a)
    return coroutine.yield(2 * a) -- 返回  2*a 的值
end

co = coroutine.create(function (a , b)
    print("第一次协同程序执行输出", a, b) -- co-body 1 10
    local r = foo(a + 1)

    print("第二次协同程序执行输出", r)
    local r, s = coroutine.yield(a + b, a - b)  -- a，b的值为第一次调用协同程序时传入

    print("第三次协同程序执行输出", r, s)
    return b, "结束协同程序"                   -- b的值为第二次调用协同程序时传入
end)

print("main", coroutine.resume(co, 1, 10)) -- true, 4
print("--分割线----")
print("main", coroutine.resume(co, "r")) -- true 11 -9
print("---分割线---")
print("main", coroutine.resume(co, "x", "y")) -- true 10 end
print("---分割线---")
print("main", coroutine.resume(co, "x", "y")) -- cannot resume dead coroutine
print("---分割线---")

-- 调用 resume，将协同程序唤醒，resume 操作成功返回 true，否则返回 false；
-- 协同程序运行；
-- 运行到 yield 语句；
-- yield 挂起协同程序，第一次 resume 返回；（注意：此处 yield 返回，参数是 resume 的参数）
-- 第二次 resume，再次唤醒协同程序；（注意：此处 resume 的参数中，除了第一个参数，剩下的参数将作为 yield 的参数）
-- yield 返回；
-- 协同程序继续运行；
-- 如果使用的协同程序继续运行完成后继续调用 resumev 方法则输出：cannot resume dead coroutine
```

Lua 的协同程序来完成生产者 - 消费者这一经典问题。  
```lua
local newProductor

function productor()
     local i = 0
     while true do
          i = i + 1
          send(i)     -- 将生产的物品发送给消费者
     end
end

function consumer()
     while true do
          local i = receive()     -- 从生产者那里得到物品
          print(i)
     end
end

function receive()
     local status, value = coroutine.resume(newProductor)
     return value
end

function send(x)
     coroutine.yield(x)     -- x表示需要发送的值，值返回以后，就挂起该协同程序
end

-- 启动程序
newProductor = coroutine.create(productor)
consumer()
```

**文件 IO**  
Lua I/O 库用于读取和处理文件，分为简单模式（和 C 一样）、完全模式。

- 简单模式（simple model）拥有一个当前输入文件和一个当前输出文件，并且提供针对这些文件相关的操作。  
简单模式在做一些简单的文件操作时较为合适。

- 完全模式（complete model） 使用外部的文件句柄来实现。它以一种面对对象的形式，将所有的文件操作定义为文件句柄的方法。  
在进行一些高级的文件操作的时候，简单模式就显得力不从心。例如，同时读取多个文件这样的操作，使用完全模式则较为合适。  
```lua
file = io.open (filename [, mode])
```

mod 值有如下这些：  

| 模式 | 描述 |  
| :---: | :--- |   
| r	| 以只读方式打开文件，该文件必须存在。 |  
| w	| 打开只写文件，若文件存在则文件长度清为 0，即该文件内容会消失。若文件不存在则建立该文件。 |  
| a	| 以附加的方式打开只写文件。若文件不存在，则会建立该文件，如果文件存在，写入的数据会被加到文件尾，即文件原先的内容会被保留。（EOF 符保留） |  
| r+ | 以可读写方式打开文件，该文件必须存在。 |  
| w+ | 打开可读写文件，若文件存在则文件长度清为零，即该文件内容会消失。若文件不存在则建立该文件。 |  
| a+ | 与 a 类似，但此文件可读可写 |  
| b	| 二进制模式，如果文件是二进制文件，可以加上 b |  
| +	| 号表示对文件既可以读也可以写 |  

```lua
-- 简单模式读写文件

-- 以只读方式打开文件
file = io.open("test.lua", "r")

-- 设置默认输入文件为 test.lua
io.input(file)

-- 输出文件第一行
print(io.read())
--[[
io.read 参数如下：
	"*n": 读取一个数字并返回它。例：file.read ("*n")
	"*a": 从当前位置读取整个文件。例：file.read ("*a")
	"*l"（默认）: 读取下一行，在文件尾 (EOF) 处返回 nil。例：file.read ("*l")
	number: 返回一个指定字符个数的字符串，或在 EOF 时返回 nil。例：file.read(5)
]]--


-- 关闭打开的文件
io.close(file)

-- 以附加的方式打开只写文件
file = io.open("test.lua", "a")

-- 设置默认输出文件为 test.lua
io.output(file)

-- 在文件最后一行添加 Lua 注释
io.write("--  test.lua 文件末尾注释")

-- 关闭打开的文件
io.close(file)
```

其他的 io 方法有：

- io.tmpfile(): 返回一个临时文件句柄，该文件以更新模式打开，程序结束时自动删除
- io.type(file): 检测 obj 是否一个可用的文件句柄
- io.flush(): 向文件写入缓冲中的所有数据
- io.lines(optional file name): 返回一个迭代函数，每次调用将获得文件中的一行内容，当到文件尾时，将返回 nil, 但不关闭文件

通常我们需要在同一时间处理多个文件，需要使用 file:function_name 来代替 io.function_name 方法。  
```lua
-- 完全模式读写文件

-- 以只读方式打开文件
file = io.open("test.lua", "r")

-- 输出文件第一行
print(file:read())
-- read 的参数与简单模式一致

-- 关闭打开的文件
file:close()

-- 以附加的方式打开只写文件
file = io.open("test.lua", "a")

-- 在文件最后一行添加 Lua 注释
file:write("--test")

-- 关闭打开的文件
file:close()
```

其他方法：  

- file:seek(optional whence, optional offset): 设置和获取当前文件位置，成功则返回最终的文件位置 (按字节), 失败则返回 nil 加错误信息。  
参数 whence 值可以是:  
"set": 从文件头开始  
"cur": 从当前位置开始 [默认]  
"end": 从文件尾开始  
offset: 默认为 0  
不带参数 file:seek () 则返回当前位置，file:seek (“set”) 则定位到文件头，file:seek (“end”) 则定位到文件尾并返回文件大小

- file:flush(): 向文件写入缓冲中的所有数据  

- io.lines(optional file name): 打开指定的文件 filename 为读模式并返回一个迭代函数，每次调用将获得文件中的一行内容，当到文件尾时，将返回 nil, 并自动关闭文件。  
若不带参数时 io.lines()、io.input():lines(); 读取默认输入设备的内容，但结束时不关闭文件。  
```lua
for line in io.lines("main.lua") do
    print(line)
end
```

以下实例使用了 seek 方法，定位到文件倒数第 25 个位置并使用 read 方法的 \*a 参数，即从当期位置 (倒数第 25 个位置) 读取整个文件。  
```lua
-- 以只读方式打开文件
file = io.open("test.lua", "r")

file:seek("end",-25)
print(file:read("*a"))

-- 关闭打开的文件
file:close()
```

**错误处理**  
程序运行中错误处理是必要的，在我们进行文件操作，数据转移及 web service 调用过程中都会出现不可预期的错误。如果不注重错误信息的处理，就会造成信息泄露，程序无法运行等情况。  
任何程序语言中，都需要错误处理。错误类型有：  

- 语法错误
语法错误通常是由于对程序的组件（如运算符、表达式）使用不当引起的。  
```lua
-- test.lua 文件
a == 2
-- lua: test.lua:2: syntax error near '=='
```

- 运行错误
运行错误是程序可以正常执行，但是会输出报错信息。  
```lua
function add(a,b)
   return a+b
end

add(10)
-- lua: test.lua:2: attempt to perform arithmetic on local 'b' (a nil value)
```

我们可以使用 assert 和 error 来处理错误。  
assert 首先检查第一个参数，若没问题，assert 不做任何事情；否则，assert 以第二个参数作为错误信息抛出。  
```lua
local function add(a,b)
   assert(type(a) == "number", "a 不是一个数字")
   assert(type(b) == "number", "b 不是一个数字")
   return a+b
end

add(10)
-- lua: test.lua:3: b 不是一个数字
```

error 函数会终止正在执行的函数，并返回 message 的内容作为错误信息 (error 函数永远都不会返回)。它的语法格式如下：  
```lua
error (message [, level])
```
通常情况下，error 会附加一些错误位置的信息到 message 头部。  

Level 参数指示获得错误的位置:  
Level=1 [默认]：为调用 error 位置 (文件 + 行号)  
Level=2：指出哪个调用 error 的函数的函数  
Level=0: 不添加错误位置信息  

Lua 中处理错误，可以使用函数 pcall（protected call）来包装需要执行的代码。  
pcall 接收一个函数和要传递个后者的参数，并执行，执行结果：有错误、无错误；返回值 true 或者或 false, errorinfo。  
pcall 以一种” 保护模式” 来调用第一个参数，因此 pcall 可以捕获函数执行中的任何错误。  
```lua
if pcall(function_name, ….) then
-- 没有错误
else
-- 一些错误
end
```

通常在错误发生时，希望获得更多的调试信息，而不只是发生错误的位置；但 pcall 返回时，它已经销毁了调用桟的部分内容。  
Lua 提供了 xpcall 函数，xpcall 接收第二个参数 —— 一个错误处理函数，当错误发生时，Lua 会在调用桟展看（unwind）前调用错误处理函数，于是就可以在这个函数中使用 debug 库来获取关于错误的额外信息了。  
debug 库提供了两个通用的错误处理函数:  
debug.debug：提供一个 Lua 提示符，让用户来价差错误的原因  
debug.traceback：根据调用桟来构建一个扩展的错误消息  
```lua
function myfunction()
   n = n/nil
end

function myerrorhandler(err)
   print("ERROR:", err)
end

status = xpcall(myfunction, myerrorhandler)
print(status)
-- ERROR: test2.lua:2: attempt to perform arithmetic on global 'n' (a nil value)
-- false
```

**调试**  
Lua 提供了 debug 库用于提供创建我们自定义调试器的功能。  
Lua 本身并未有内置的调试器，但很多开发者共享了他们的 Lua 调试器代码。  

Lua 中 debug 库包含以下函数。  

| 方法 | 用途 |  
| ---: | :--- |  
| debug() | 进入一个用户交互模式，运行用户输入的每个字符串。使用简单的命令以及其它调试设置，用户可以检阅全局变量和局部变量， 改变变量的值，计算一些表达式等等。输入一行仅包含 cont 的字符串将结束这个函数， 这样调用者就可以继续向下运行。 |  
| getfenv(object) | 返回对象的环境变量 |  
| gethook(optional thread) | 返回三个表示线程钩子设置的值： 当前钩子函数，当前钩子掩码，当前钩子计数 |  
| getinfo ([thread,] f [, what]) | 返回关于一个函数信息的表。 你可以直接提供该函数， 也可以用一个数字 f 表示该函数。 数字 f 表示运行在指定线程的调用栈对应层次上的函数： 0 层表示当前函数（getinfo 自身）； 1 层表示调用 getinfo 的函数 （除非是尾调用，这种情况不计入栈）；等等。 如果 f 是一个比活动函数数量还大的数字， getinfo 返回 nil |  
| debug.getlocal ([thread,] f, local) | 此函数返回在栈的 f 层处函数的索引为 local 的局部变量 的名字和值。 这个函数不仅用于访问显式定义的局部变量，也包括形参、临时变量等 |  
| getmetatable(value) | 把给定索引指向的值的元表压入堆栈。如果索引无效，或是这个值没有元表，函数将返回 0 并且不会向栈上压任何东西 |  
| getregistry() | 返回注册表表，这是一个预定义出来的表， 可以用来保存任何 C 代码想保存的 Lua 值 |  
| getupvalue(f, up) | 此函数返回函数 f 的第 up 个上值的名字和值。 如果该函数没有那个上值，返回 nil。以 ‘(‘ （开括号）打头的变量名表示没有名字的变量 （去除了调试信息的代码块） |  
| setlocal([thread,] level, local, value) | 这个函数将 value 赋给 栈上第 level 层函数的第 local 个局部变量。 如果没有那个变量，函数返回 nil 。 如果 level 越界，抛出一个错误。 |  
| setmetatable(value, table) | 将 value 的元表设为 table （可以是 nil），返回 value |  
| setupvalue(f, up, value) | 这个函数将 value 设为函数 f 的第 up 个上值。 如果函数没有那个上值，返回 nil 否则，返回该上值的名字 |  
| traceback ([thread,] [message [, level]]) | 如果 message 有，且不是字符串或 nil， 函数不做任何处理直接返回 message。 否则，它返回调用栈的栈回溯信息。 字符串可选项 message 被添加在栈回溯信息的开头。 数字可选项 level 指明从栈的哪一层开始回溯 （默认为 1 ，即调用 traceback 的那里） |  

```lua
function myfunction ()
	print(debug.traceback("Stack trace"))
	print(debug.getinfo(1))
	print("Stack trace end")

    return 10
end

myfunction ()
print(debug.getinfo(1))
--[[
Stack trace
stack traceback:
   test2.lua:2: in function 'myfunction'
   test2.lua:8: in main chunk
    [C]: ?
table: 0054C6C8
Stack trace end
]]--
```

调试类型分命令行调试和图形界面调试。  
命令行调试器有这些：RemDebug、clidebugger、ctrace、xdbLua、LuaInterface – Debugger、Rldb、ModDebug。  
图形界调试器有这些：SciTE、Decoda、ZeroBrane Studio、akdebugger、luaedit。

**垃圾回收**  
Lua 采用了自动内存管理。   
这意味着你不用操心新创建的对象需要的内存如何分配出来，也不用考虑在对象不再被使用后怎样释放它们所占用的内存。  

Lua 运行了一个垃圾收集器来收集所有死对象 （即在 Lua 中不可能再访问到的对象）来完成自动内存管理的工作。 Lua 中所有用到的内存，如：字符串、表、用户数据、函数、线程、 内部结构等，都服从自动管理。  
Lua 实现了一个增量标记 - 扫描收集器。 它使用这两个数字来控制垃圾收集循环： 垃圾收集器间歇率和垃圾收集器步进倍率。 这两个数字都使用百分数为单位 （例如：值 100 在内部表示 1 ）。  

垃圾收集器间歇率控制着收集器需要在开启新的循环前要等待多久。 增大这个值会减少收集器的积极性。 当这个值比 100 小的时候，收集器在开启新的循环前不会有等待。 设置这个值为 200 就会让收集器等到总内存使用量达到 之前的两倍时才开始新的循环。  
垃圾收集器步进倍率控制着收集器运作速度相对于内存分配速度的倍率。 增大这个值不仅会让收集器更加积极，还会增加每个增量步骤的长度。 不要把这个值设得小于 100 ， 那样的话收集器就工作的太慢了以至于永远都干不完一个循环。 默认值是 200 ，这表示收集器以内存分配的” 两倍” 速工作。  
如果你把步进倍率设为一个非常大的数字 （比你的程序可能用到的字节数还大 10% ）， 收集器的行为就像一个 stop-the-world 收集器。 接着你若把间歇率设为 200 ， 收集器的行为就和过去的 Lua 版本一样了： 每次 Lua 使用的内存翻倍时，就做一次完整的收集。  

Lua 提供了以下函数 collectgarbage([opt [, arg]]) 用来控制自动内存管理。  

- collectgarbage("collect"): 做一次完整的垃圾收集循环。通过参数 opt 它提供了一组不同的功能：
- collectgarbage("count"): 以 K 字节数为单位返回 Lua 使用的总内存数。 这个值有小数部分，所以只需要乘上 1024 就能得到 Lua 使用的准确字节数（除非溢出）。  
- collectgarbage("restart"): 重启垃圾收集器的自动运行。
- collectgarbage("setpause"): 将 arg 设为收集器的 间歇率 （参见 §2.5）。 返回 间歇率 的前一个值。
- collectgarbage("setstepmul"): 返回 步进倍率 的前一个值。
- collectgarbage("step"): 单步运行垃圾收集器。 步长” 大小” 由 arg 控制。 传入 0 时，收集器步进（不可分割的）一步。 传入非 0 值， 收集器收集相当于 Lua 分配这些多（K 字节）内存的工作。 如果收集器结束一个循环将返回 true 。
- collectgarbage("stop"): 停止垃圾收集器的运行。 在调用重启前，收集器只会因显式的调用运行。

```lua
mytable = {"apple", "orange", "banana"}

print(collectgarbage("count")) -- 20.9560546875

mytable = nil

print(collectgarbage("count")) -- 20.9853515625 

print(collectgarbage("collect")) -- 0

print(collectgarbage("count")) -- 19.4111328125
```

**面向对象**  
面向对象编程（Object Oriented Programming，OOP）是一种非常流行的计算机编程架构，它有封装、继承、多态、抽象的特征。  

对象由属性和方法组成，Lua 中最基本的结构是 table，所以需要用 table 来描述对象的属性，function 可以用来表示方法。而对于继承，可以通过 metetable 模拟出来（不推荐用，只模拟最基本的对象大部分实现就够用了）。  
Lua 中的表不仅在某种意义上是一种对象。像对象一样，表也有状态（成员变量）；也有与对象的值独立的本性，特别是拥有两个不同值的对象（table）代表两个不同的对象；一个对象在不同的时候也可以有不同的值，但他始终是一个对象；与对象类似，表的生命周期与其由什么创建、在哪创建没有关系。
```lua
-- Meta class
Rectangle = {area = 0, length = 0, breadth = 0}

-- 派生类的方法 new
-- 注意：冒号只是起了省略第一个参数 self 的作用；如果使用点号，则需要传递 self（如下可写成 function Rectangle.new(self,o,length,breadth)）；建议使用冒号
function Rectangle:new(o,length,breadth)
  o = o or {}
  setmetatable(o, self)
  -- 引用参数 self 指向调用者自身
  self.__index = self
  self.length = length or 0
  self.breadth = breadth or 0
  self.area = length*breadth;

  return o
end

-- 派生类的方法 printArea
function Rectangle:printArea()
  print("矩形面积为 ",self.area)
end

-- 创建对象是为类的实例分配内存的过程，每个类都有属于自己的内存并共享公共数据
r = Rectangle:new(nil,10,20)
-- 访问属性
print(r.length) -- 10
-- 访问成员函数，内存在对象初始化时分配
r:printArea()
```

**数据库访问**  
Lua 数据库的操作库是 LuaSQL，它支持的数据库有 ODBC, ADO, Oracle, MySQL, SQLite 和 PostgreSQL。  
LuaSQL 可以使用 LuaRocks 来安装可以根据需要安装你需要的数据库驱动。  
```bash
wget http://luarocks.org/releases/luarocks-2.2.1.tar.gz
tar zxpf luarocks-2.2.1.tar.gz
cd luarocks-2.2.1
./configure; sudo make bootstrap
sudo luarocks install luasocket

# 安装不同数据库驱动
luarocks install luasql-sqlite3
luarocks install luasql-postgres
luarocks install luasql-mysql
luarocks install luasql-sqlite
luarocks install luasql-odbc
```

Lua 连接 MySql 数据库，并之行相应操作。  
```lua
require "luasql.mysql"

-- 创建环境对象
env = luasql.mysql()

-- 连接数据库
conn = env:connect("数据库名","用户名","密码","IP 地址","端口")

-- 设置数据库的编码格式
conn:execute("SET NAMES UTF8")

-- 执行数据库操作
cur = conn:execute("select * from role")
row = cur:fetch({},"a")

-- 文件对象的创建
file = io.open("role.txt","w+");

while row do
    var = string.format("%d %s\n", row.id, row.name)
    print(var)
    file:write(var)

    row = cur:fetch(row,"a")
end


file:close()  -- 关闭文件对象
conn:close()  -- 关闭数据库连接
env:close()   -- 关闭数据库环境
```

### Lua 进阶

**常用操作系统库**  
os.time([table])  
功能：按 table 的内容返回一个时间值 (数字), 若不带参数则返回当前时间（在许多系统中该数值是当前距离某个特定时间的秒数）。  
说明：当为函数调用附加一个特殊的时间表时，该函数就是返回距该表描述的时间的数值。  
```lua
print(os.time{year=1970, month=1, day=1,hour=8}) -- 默认此刻为 0
print(os.time{year=1970, month=1, day=1}) --若未定义时分秒，默认时间为正午（04:00:00）43200
```

os.date([format [, time]])  
功能：返回一个按 format 格式化日期、时间的字串或表  
说明：函数 date，其实是 time 函数的一种 “反函数”。它将一个表示日期和时间的数值，转换成更高级的表现形式。其第一个参数是一个格式化字符串，描述了要返回的时间形式。第二个参数就是时间的数字表示，默认为当前的时间。  
```lua
t = os.date("*t", os.time());
for i, v in pairs(t) do
      print(i,"->",v);
end
-- isdst	->	false
-- min	->	43
-- sec	->	49
-- yday	->	365
-- month	->	12
-- wday	->	3
-- year	->	2019
-- day	->	31
-- hour	->	10

print(os.date("today is %A, in %B")) -- today is Tuesday, in December
print(os.date("%X", 906000490)) -- 02:48:10

print(os.date("%m/%d/%Y", 906000490)) -- 09/17/1998
```

**函数回调技巧**  
```lua
local a = {};
function b()    
    print("Hello World")
    end
    
a["sell"] = {callFunc=b}
a["sell"].callFunc() -- Hello World

-- 使用 lua 自带的 unpack 实现回调
function unpackex(tbl, args)    
    local ret = {}    
    for _,v in ipairs(args) 
    do        
        table.insert(ret, tbl[v])    
    end    
    
    return unpack(ret)
end
print(unpackex({one = {"one", "two", "three"}, two = "T" , three = "TH"},{"one", "two", "three"}))
```

**迭代**  
迭代是 for 语句的一种特殊形式，可以通过 for 语句驱动迭代函数对一个给定集合进行遍历。
```lua
-- 迭代函数 enum
local function enum(array)
    local index = 1
    -- 返回一个匿名的迭代函数
    return function()
        local ret = array[index]
        index = index + 1
        
        return ret
    end
end

local function foreach(array,action)
    -- 每次调用迭代函数 enum 都得到一个值 (通过 element 变量引用)
    -- 若该值为 nil, 则 for 循环结束
    for element in enum(array)do
        action(element)
    end
end

foreach({1,2,3},print)
```

**协作线程**  
通过 coroutine.create 可以创建一个协作线程，该函数接收一个函数类型的参数作为线程的执行体，返回一个线程对象。  
通过 coroutine.resume 可以启动一个线程或者继续一个挂起的线程。该函数接收一个线程对象以及其他需要传递给该线程的参数。线程可以通过线程函数的参数或者 coroutine.yield 调用的返回值来获取这些参数。当线程初次执行时，resume 传递的参数通过线程函数的参数传递给线程，线程从线程函数开始执行；当线程由挂起转为执行时，resume 传递的参数以 yield 调用返回值的形式传递给线程，线程从 yield 调用后继续执行。  
线程调用 coroutine.yield 暂停自己的执行，并把执行权返回给启动 / 继续它的线程；线程还可利用 yield 返回一些值给后者，这些值以 resume 调用的返回值的形式返回。  

```lua
--线程
local function producer()
    return coroutine.create(
	    function(salt)
	        local t = {1,2,3}
	        for i = 1,#t do
	            salt = coroutine.yield(t[i] + salt)
	        end
	    end
    )
end

function consumer(prod)
    local salt = 10
    while true do
        local running ,product = coroutine.resume(prod, salt)
        salt = salt*salt
        if running then
            print(product or "END!")
        else
            break
        end
    end
end

consumer(producer())
```

**常用数据结构**  
Lua 中的 table 不是一种简单的数据结构，它可以作为其它数据结构的基础，如数组、记录、线性表、队列和集合等，在 Lua 中都可以通过 table 来表示。  

在 lua 中通过整数下标访问表中的元素即可简单的实现数组，并且数组不必事先指定大小，大小可以随需要动态的增长。  
`注意：在 Lua 中习惯上数组的下表从 1 开始，Lua 的标准库与此习惯保持一致，因此如果你的数组下标也是从 1 开始你就可以直接使用标准库的函数，否则就无法直接使用。`  
```lua
a = {}
for i = 1,100 do
    a[i] = 0
end
print("The length of array 'a' is " .. #a) -- The length of array 'a' is 100

squares = {1, 4, 9, 16, 25}
print("The length of array 'a' is " .. #squares) -- The length of array 'a' is 5
```

Lua 中表示矩阵（二维数组）的方法，用数组的数组表示，也就是说一个表的元素是另一个表。  
```lua
local N = 3
local M = 3
mt = {}
for i = 1,N do
    mt[i] = {}
    for j = 1,M do
        mt[i][j] = i * j
    end
end

mt = {}
for i = 1, N do
    for j = 1, M do
        mt[(i - 1) * M + j] = i * j
    end
```

Lua 中用 tables 很容易实现链表，每一个节点是一个 table，指针是这个表的一个域，并且指向另一个节点 (table)。  
```lua
-- 要实现一个链表需要两个域：值和指针的基本链表
list = nil
for i = 1, 10 do
    list = { next = list ,value = i}
end

local l = list
while l do 
    --print(l.value)
    l = l.next
end
```

虽然可以使用 Lua 的 table 库提供的 insert 和 remove 操作来实现队列，但这种方式实现的队列针对大数据量时效率太低，有效的方式是使用两个索引下标，一个表示第一个元素，另一个表示最后一个元素。  
```lua
-- 双向队列
List = {}

--创建
function List.new()
    return {first = 0,last = -1}
end

--队列头插入
function List.pushFront(list,value)
    local first = list.first - 1
    list.first = first
    list[first] = value
end

--队列尾插入
function List.popFront(list)
    local first = list.first
    if first > list.last then
        error("List is empty")
    end

    local value = list[first]
    list[first] = nil
    list.first = first + 1
    return value
end

function List.popBack(list)
    local last = list.last
    if list.first > last then
        error("List is empty")
    end
    local value = list[last]
    list[last] = nil
    list.last = last - 1 
    return value
end

--测试代码
local testList = {first = 0,last = -1}
local tableTest = 12

List.pushFront(testList,tableTest)
print( List.popFront(testList))
```

Lua 实现堆栈功能。  
```lua
local stackMng = {}
stackMng.__index = stackMng

function stackMng:new()
    local temp = {}
    setmetatable(temp,stackMng)
    return temp
end

function stackMng:init()
    self.stackList = {}
end

function stackMng:reset()
    self:init()
end

function stackMng:clear()
    self.stackList = {}
end

function stackMng:pop()
    if #self.stackList == 0 then
        return
    end
    if self.stackList[1] then
        print(self.stackList[1])
    end

    return table.remove(self.stackList,1)
end

function stackMng:push(t)
    table.insert(self.stackList,t)
end

function stackMng:Count()
    return #self.stackList
end

--测试代码
object = stackMng:new()
object:init()
object:push(1)
object:pop()
```

Lua 中用 table 实现集合。  
```lua
reserved = {
["while"] = true, 
["end"] = true,
["function"] = true, 
["local"] = true,
}

for k,v in pairs(reserved) do
    print(k,"->",v)
end
```
