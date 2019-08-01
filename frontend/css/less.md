
### Less Link
[less 官网](http://lesscss.org/)  
[less 文档](http://lesscss.cn/features/)  

https://learnxinyminutes.com/docs/zh-cn/less-cn/

### Less 是什么
Less 是一种 CSS 预处理器，它增加了诸如变量、嵌套、mixin 等功能。 Less 能帮助开发人员编写易维护，DRY (Don’t Repeat Yourself) 的代码。  

### Less 安装和使用
Less 安装非常简单
```bash
npm install -g less
```
或者直接使用 cdn
```html
<script src="//cdnjs.cloudflare.com/ajax/libs/less.js/2.5.3/less.min.js"></script>
```

Less 的使用 —— 命令行
```bash
# 编译 less 文件 styles.less 输出到 styles.css
lessc styles.less styles.css

# 编译并进行代码最小化
lessc --clean-css styles.less styles.min.css
```

Less 的使用 —— 代码
```js
var less = require('less');

less.render('.class { width: (1 + 1) }',
{
  paths: ['.', './lib'],  // Specify search paths for @import directives
  filename: 'style.less', // Specify a filename, for better error messages
  compress: true          // Minify CSS output
},
function (e, output) {
   console.log(output.css);
});
```

### Less 注释
```less
// 单行注释，当 Less 被编译成 CSS 后会被删除
/* 多行注释将保留. */
```

### Less 变量
Less 使用 @ 符号来创建一个变量
```less
// 值变量
@primary-color: #a3a4ff;
@secondary-color: #51527f;
@body-font: 'Roboto', sans-serif;
// 选择器变量
@mySelector: #wrap;
@Wrap: wrap;
// 属性变量
@borderStyle: border-style;
@Soild:solid;
// url 变量
@images: "../img";

body {
    background-color: @primary-color;
    color: @secondary-color;
    font-family: @body-font;
}

@{mySelector}{
  color: #999;
  width: 50%;
}
.@{Wrap}{
  color:#ccc;
}

#wrap{
  @{borderStyle}: @Soild;
}

body {
  background: url("@{images}/dog.png");
}
```

### Less 嵌套
Less 允许在选择子中嵌套选择子，& 将被替换成父选择子；也可以嵌套伪类。  
`建议嵌套时不超过 3 层`
```less
ul {
    list-style-type: none;
    margin-top: 2em;

    li {
        background-color: red;

        &:hover {
          background-color: blue;
        }

        a {
          color: white;
        }
    }
}}
```

### Less Mixins
如果要为多个元素编写同样的代码，Mixins 可以实现重用
```less
// .center 会被编译使用，加 () 即可不被编译
.center {
    display: block;
    margin-left: auto;
    margin-right: auto;
    left: 0;
    right: 0;
}

div {
    .center;
    background-color: @primary-color;
}
```

### Less 函数
Less 提供的函数可以用来完成多种任务  
[更多函数参考](http://lesscss.cn/functions/)  
```less
body {
  width: round(10.25px);
}

.header {
    background-color: lighten(#000, 0.5);
}

.footer {
  background-color: fadeout(#000, 0.25)
}
```

也可以定义自己的函数。  
函数非常类似于 mixin，区别是 mixin 最适合用来创建 CSS 而函数更适合于处理那些可能在你的 Less 代码中使用的逻辑。
```less
.average(@x, @y) {
  @average-result: ((@x + @y) / 2);
}

div {
  .average(16px, 50px);        // 调用函数 average
  padding: @average-result;    // 使用它的"返回"值
}
```

### Less 扩展
扩展是在选择子间共享属性的一种方法。  
扩展一条 CSS 语句优于创建一个 mixin，这是由其组合所有共享相同基样式的类的方式决定的。  
如果使用 mixin 完成，其属性将会在调用了该 mixin 的每条语句中重复。虽然它不至会影响你的工作流，但它会在由 Less 编译器生成的的文件中添加不必要的代码。
```less
.display {
  height: 50px;
}

.display-success {
  &:extend(.display);
    border-color: #22df56;
}
```

### Less 片段和导入
Less 允许你创建片段文件，它有助于你的 Less 代码保持模块化。  
片段文件习惯上以 \_ 开头，例如 \_reset.css，并被导入到一个将会被编译成 CSS 的主 less 文件中。Less 提取导入文件并将它们与编译后的代码结合起来。  
```less
// _reset.css
html,
body,
ul,
ol {
  margin: 0;
  padding: 0;
}

// styles.less
@import 'reset';

body {
  font-size: 16px;
  font-family: Helvetica, Arial, Sans-serif;
}
```

### Less 计算
Less 提供以下的运算符: +, \-, \*, /, 和 %。相比于使用你事先手工计算好了的数值，它们对于直接在你的 Less 文件中计算数值很有用。  
```less
@content-area: 960px;
@main-content: 600px;
@sidebar-content: 300px;

@main-size: @main-content / @content-area * 100%;
@sidebar-size: @sidebar-content / @content-area * 100%;
@gutter: 100% - (@main-size + @sidebar-size);

body {
  width: 100%;
}

.main-content {
  width: @main-size;
}

.sidebar {
  width: @sidebar-size;
}

.gutter {
  width: @gutter;
}
```