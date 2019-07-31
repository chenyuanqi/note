
### Sass Link
[Sass 官网](https://sass-lang.com/)  
[Sass 中文文档](http://sass.bootcss.com/docs/sass-reference/)  

[Sass 专家](https://www.sassmeister.com/)  

### Sass 是什么
Sass 是一种 CSS 扩展语言，它增加了诸如变量、嵌套、mixin 等功能。 Sass 能帮助开发人员编写易维护和 DRY (Don’t Repeat Yourself) 的代码。  

Sass 有两种不同的语法可选用。SCSS 的语法和 CSS 的相同，但增加了 Sass 的额外功能，扩展名是 .scss；Sass（原来的语法）是以严格语法规则书写，它使用缩进而非大括号和分号，扩展名是 .sass。  
`快速入门请选择 scss`

### Sass 安装和使用
Sass 是基于 ruby 的，要先安装 ruby 环境。[ruby 下载](http://rubyinstaller.org/downloads)， ruby 安装好以后用以下命令安装 sass。  
```bash
gem install sass

# 查看 sass 版本（检查是否安装成功）
sass -v 

# 编译（把 sass 或 scss 转换成 css）
sass input.scss output.css
# 整体监听编译（把 LearnSass 目录下的 scss 全部转换成 css，再保存到 LearnCss 目录下）
sass --watch LearnSass:LearnCss
```

### Sass 变量
sass 用 $ 定义变量，用来存储一些重复的 CSS 值 (在整个站点中一致的值)。  
```scss
$primary-color: #333;

body {
  color: $primary-color;
}
```

### Sass 嵌套
Sass 可以明确的体现 Html 的层次关系。  
```scss
nav{
    ul{
        list-style: none;
    }
    li{
        display: inline-block;
    }
    a{
        display: block;
        padding: 6px 12px;
        text-decoration: none;
    }
}
```

### Sass Partials
Sass 定义了一个规则，用 "\_" 下划线开头的 Sass 文件则不会被转换成 css。
```scss
/* 后缀 scss 可以省略，但是引号是必须的 */
@import '_partials';
```  

### Sass 继承
使用 @extend 命令可以引用一个选择器的样式，这样的写法可以避免 HTML 元素上多写类名，也合理的体现了代码重用。  
```scss
.message {
    border: 1px solid #ccc;
    padding: 10px;
    color: #333;
}

.success {
    @extend .message;
    border-color: green;
}

.error {
    @extend .message;
    border-color: red;
}
```

### Sass 占位符
占位符在创建用于扩展的 CSS 语句时非常有用。如果你想创建一条只通过 @extend 使用的 CSS 语句，你可以利用占位符来实现。  
```scss
%content-window {
  font-size: 14px;
  padding: 10px;
  color: #000;
  border-radius: 4px;
}

.message-window {
  @extend %content-window;
  background-color: #0000ff;
}
```

### Sass 计算
Sass 提供简单的计算，提供标准运算符：+， -， \*， /，和 %。  
```scss
.container { width: 100%; }

.roleHeight{
    height: 100px;
    border: 1px solid;
    margin-top: 20px;
}

article[role="main"] {
    @extend .roleHeight;
    float: left;
    width: 600px / 960px * 100%;
}

aside[role="complementary"] {
    @extend .roleHeight;
    float: right;
    width: 300px / 960px * 100%;
}
```

### Sass 控制命令
Sass 允许你使用 @if, @else, @for, @while 和 @each 来控制你的代码如何编译成 CSS。  
```scss
$debug: true !default;

@mixin debugmode {
    @if $debug {
        @debug "Debug mode enabled";

        display: inline-block;
    }
    @else {
        display: none;
    }
}

.info {
    @include debugmode;
}

@for $i from 1 to 4 {
    div:nth-of-type(#{$i}) {
        left: ($i - 1) * 900 / 3;
    }
}

$columns: 4;
$column-width: 80px;

@while $columns > 0 {
    .col-#{$columns} {
        width: $column-width;
        left: $column-width * ($columns - 1);
    }

    $columns: $columns - 1;
}

/* @each 函数类似 @for, 除了它使用一个列表而不是序列值 */
$social-links: facebook twitter linkedin reddit;

.social-links {
    @each $sm in $social-links {
        .icon-#{$sm} {
            background-image: url("images/#{$sm}.png");
        }
    }
}
```

### Sass Mixin
Mixin 有点像是一个函数，它们的区别是 mixin 最适合于创建 CSS 而函数更适合于处理那些可能在 Sass 代码中使用的逻辑。  
比如生成每个浏览器供应商的前缀。  
```scss
@mixin border-radius($radius) {
  -webkit-border-radius: $radius;
     -moz-border-radius: $radius;
      -ms-border-radius: $radius;
          border-radius: $radius;
}

/* 使用 @include 调用 mixin */
.box { 
    @include border-radius(10px); 
    width: 100px;
    height: 60px;
    border: 1px solid;
}
```

### Sass 函数
Sass 提供的[函数](https://sass-lang.com/documentation/functions)可以用来完成各种各样的任务。  
```scss
.footer {
  background-color: fade_out(#000000, 0.25);
}
```

也可以定义自己的函数。  
```scss
@function calculate-percentage($target-size, $parent-size) {
  @return $target-size / $parent-size * 100%;
}

$main-content: calculate-percentage(600px, 960px);

.main-content {
  width: $main-content;
}
```
