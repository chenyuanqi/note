
### 一、Css 权重
CSS 权重指的是样式的优先级，有两条或多条样式作用于一个元素，权重高的那条样式对元素起作用，权重相同的，后写的样式会覆盖前面写的样式。

#### 1.1 权重的等级
Css 权重可以把样式的应用方式分为几个等级，按照等级来计算权重   
> 1.!important，加在样式属性值后，权重值为 10000  
> 2.内联样式，如：style=””，权重值为 1000  
> 3.ID 选择器，如：#content，权重值为 100  
> 4.类，伪类和属性选择器，如： content、:hover 权重值为 10  
> 5.标签选择器和伪元素选择器，如：div、p、:before 权重值为 1  
> 6.通用选择器（*）、子选择器（>）、相邻选择器（+）、同胞选择器（~）、权重值为 0  

#### 1.2 权重的计算

### 二、Css 选择器
#### 2.1 标签选择器
标签选择器，此种选择器影响范围大，建议尽量应用在层级选择器中。  
```css
*{margin:0;padding:0}
div{color:red}   
```

#### 2.2 id 选择器
通过 id 名来选择元素，元素的 id 名称不能重复，所以一个样式设置项只能对应于页面上一个元素，不能复用，id 名一般给程序使用，所以不推荐使用 id 作为选择器。  
```css
#box{color:red} 
```

#### 2.3 类选择器
通过类名来选择元素，一个类可应用于多个元素，一个元素上也可以使用多个类，应用灵活，可复用，是 css 中应用最多的一种选择器。  
```css
.red{color:red}
.big{font-size:20px}
.mt10{margin-top:10px} 
```

#### 2.4 层级选择器
主要应用在选择父元素下的子元素，或者子元素下面的子元素，可与标签元素结合使用，减少命名，同时也可以通过层级，防止命名冲突。
```html
<style>
.box span{color:red}
.box .red{color:pink}
.red{color:red}
</style>
<div class="box">
    <span>....</span>
    <a href="#" class="red">....</a>
</div>

<h3 class="red">....</h3>
```

#### 2.5 组选择器
多个选择器，如果有同样的样式设置，可以使用组选择器。  
```css
.box1,.box2,.box3{width:100px;height:100px}
.box1{background:red}
.box2{background:pink}
.box2{background:gold}
```

#### 2.6 伪类及伪元素选择器
常用的伪类选择器有 hover，表示鼠标悬浮在元素上时的状态，伪元素选择器有 before 和 after, 它们可以通过样式在元素中插入内容。
```css
.box1:hover{color:red}
.box2:before{content:'行首文字';}
.box3:after{content:'行尾文字';}
```
