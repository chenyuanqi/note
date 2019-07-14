
### Css Link
[css 规范](https://github.com/necolas/idiomatic-css/tree/master/translations/zh-CN)  
[css 魔法](https://github.com/cssmagic/CSS-Secrets)  

[css 布局](http://learnlayout.com/)  
[css 选择器参考](https://css4-selectors.com/selectors/)  

[css 技巧](https://github.com/AllThingsSmitty/css-protips/tree/master/translations/zh-CN)  
[css 常用样式](https://qishaoxuan.github.io/css_tricks/)  

### Css 小技巧
1、不使用 !important 也能让它变绿
```html
<div class="foo">
  <div class="bar"></div>
</div>

<style>
.foo .bar {
  background: red;
}

.bar.bar {
  background: green;
}
</style>
```

2、选择器中特殊字符的处理
```html
<!-- 数字开头的 class 需要加 .\34 才生效 -->
<div class=".\34 404-page"></div>
```

3、实时编辑 Css
```html
<!DOCTYPE html>
<html>
    <body>
        <style style="display:block" contentEditable>
            body { color: blue }
        </style>
    </body>
</html>
```