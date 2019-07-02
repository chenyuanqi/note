
### 什么是 em
em 是 css 的单位长度，相对于固定像素的 px，em 更具灵活性。  

- em 是相对长度单位，更适用于响应式布局  
> 子元素字体大小的 em 是相对于父元素字体大小   
> 元素的 width/height/padding/margin 用 em 的话是相对于该元素的 font-size  
```html
<div>
    我是父元素 div
    <p>
        我是子元素 p
        <span>我是孙元素 span</span>
    </p>
</div>
<style>
div {
  font-size: 40px;
  width: 10em; /* 400px */
  height: 10em;
  border: solid 1px black;
}

p {
  font-size: 0.5em; /* 20px */ 
  width: 10em; /* 200px */
  height: 10em;
  border: solid 1px red;
}

span {
  font-size: 0.5em;  /* 10px，如果是 chrome 则 12px，因为 chrome 设置最小的字体大小是 12px*/
  width: 10em; /* 100px/120px */
  height: 10em;
  border: solid 1px blue;
  display: block;
}
</style>
```
