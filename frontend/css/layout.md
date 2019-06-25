
### 一、布局之道
#### 1.1 居中布局
##### 1.1.1 水平居中
子元素于父元素水平居中且其（子元素与父元素）宽度均可变。  
方案 1：inline-block + text-align 兼容性佳（甚至可以兼容 IE 6 和 IE 7）  
```html
<div class="parent">
  <div class="child">Demo</div>
</div>

<style>
  .child {
    display: inline-block;
  }
  .parent {
    text-align: center;
  }
</style>
```

方案 2：table + margin 无需设置父元素样式 （支持 IE 8 及其以上版本）
```html
<div class="parent">
  <div class="child">Demo</div>
</div>

<style>
  .child {
  	// display: table 在表现上类似 block 元素，但是宽度为内容宽
    display: table;
    margin: 0 auto;
  }
</style>
```

方案 3：absolute + transform 绝对定位脱离文档流，不会对后续元素的布局造成影响；但是 transform 为 CSS3 属性，有兼容性问题
```html
<div class="parent">
  <div class="child">Demo</div>
</div>

<style>
  .parent {
    position: relative;
  }
  .child {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
  }
</style>
```

方案 4：flex + justify-content 只需设置父节点属性，无需设置子元素；但是存在兼容问题
```html
<div class="parent">
  <div class="child">Demo</div>
</div>

<style>
  .parent {
    display: flex;
    justify-content: center;
  }

  /* 或者下面的方法，可以达到一样的效果 */
  .parent {
    display: flex;
  }
  .child {
    margin: 0 auto;
  }
</style>
```

##### 1.1.2 垂直居中
子元素于父元素垂直居中且其（子元素与父元素）高度均可变。  
方案 1：table-cell + vertical-align 兼容性好（支持 IE 8，以下版本需要调整页面结构至 table）
```html
<div class="parent">
  <div class="child">Demo</div>
</div>

<style>
  .parent {
    display: table-cell;
    vertical-align: middle;
  }
</style>
```

方案 2：absolute + transform 绝对定位脱离文档流，不会对后续元素的布局造成影响，但如果绝对定位元素是唯一的元素则父元素也会失去高度；transform 为 CSS3 属性，有兼容性问题  
```html
<div class="parent">
  <div class="child">Demo</div>
</div>

<style>
  .parent {
    position: relative;
  }
  .child {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
  }
</style>
```

方案 3：flex + align-items 只需设置父节点属性，无需设置子元素；有兼容性问题
```html
<div class="parent">
  <div class="child">Demo</div>
</div>

<style>
  .parent {
    display: flex;
    align-items: center;
  }
</style>
```

##### 1.1.3 水平与垂直居中
子元素于父元素垂直及水平居中且其（子元素与父元素）高度宽度均可变。  
方案 1：inline-block + text-align + table-cell + vertical-align 兼容性好
```html
<div class="parent">
  <div class="child">Demo</div>
</div>

<style>
  .parent {
    text-align: center;
    display: table-cell;
    vertical-align: middle;
  }
  .child {
    display: inline-block;
  }
</style>
```

方案 2：absolute + transform 绝对定位脱离文档流，不会对后续元素的布局造成影响；transform 为 CSS3 属性，有兼容性问题
```html
<div class="parent">
  <div class="child">Demo</div>
</div>

<style>
  .parent {
    position: relative;
  }
  .child {
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
  }
</style>
```

方案 3：flex + justify-content + align-items 只需设置父节点属性，无需设置子元素；有兼容性问题
```html
<div class="parent">
  <div class="child">Demo</div>
</div>

<style>
  .parent {
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>
```

#### 1.2 多列布局
多列布局在网页中非常常见（例如两列布局），多列布局可以是两列定宽，一列自适应， 或者多列不定宽一列自适应还有等分布局等。

http://caibaojian.com/fend_note/chapter4/02_layout.html

flex 布局：
https://juejin.im/post/58e3a5a0a0bb9f0069fc16bb

### 二、响应式布局
多屏的环境让我们不得不考虑网络内容在各个尺寸中的表现， 均可正常访问和极佳的用户体验。

响应式布局可以更具屏幕尺子的大小对内容和布局做出适当的调成， 从而提供更好的用户感受。也是因为响应式布局的出现， 开发者也无需对不同尺寸设备而特殊定制不同的页面， 这大大降低了开发成本和缩短了开发时间。

这样的方法也同样存在着缺点。 缺点是同样的资源被加载，但因为展示平台所限并不能全部显示。
```html
<meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=no">
```

针对不同尺寸的屏幕进行开发，因少使用定宽而使用自适应单位。  
需求会更具具体设备而产生变化。  
```css
@media screen and (max-width: 320px) {
  /* 视窗宽度小于等于 320px */
}
@media screen and (min-width: 320px) {
  /* 视窗宽度大于等于 320px */
}
@media screen and (min-width: 320px) and (max-width: 1000px){
  /* 视窗宽度大于等于 320px 且小于等于 1000px */
}
```
