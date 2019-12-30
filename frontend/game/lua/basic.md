
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
-- 查找字符串
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


