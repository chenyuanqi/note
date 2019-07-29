
### 浏览器的 Js
我们知道，浏览器地址栏可以直接运行 JavaScript 代码，做法是以 javascript: 开头后跟要执行的语句。  
```js
// 把浏览器当编辑器
data:text/html, <html contenteditable>

// 编辑网页
javascript:document.body.contentEditable='true'; document.designMode='on'; void 0

// 显示页面所有链接
javascript:a=document.getElementsByTagName("A");newwindow=window.open("newwindow");newwindow.document.open();for(i=0;i<a.length;i++){newwindow.document.write("<a href='"+a.item(i).href+"'>"+a.item(i).innerText+"</a><br>");}newwindow.document.close();void(0)

// 显示网页中的所有图片
javascript:Ai7Mg6P='';for%20(i7M1bQz=0;i7M1bQz<document.images.length;i7M1bQz++){Ai7Mg6P+='<img%20src='+document.images[i7M1bQz].src+'><br>'};if(Ai7Mg6P!=''){document.write('<center>'+Ai7Mg6P+'</center>');void(document.close())}else{alert('No%20images!')}

// 显示网页中除图片的其他
javascript:for(jK6bvW=0;jK6bvW<document.images.length;jK6bvW++){void(document.images[jK6bvW].style.visibility='hidden')}

// 无敌风火轮	
javascript:R=0; x1=.1; y1=.05; x2=.25; y2=.24; x3=1.6; y3=.24; x4=300; y4=200; x5=300; y5=200; DI=document.getElementsByTagName("img"); DIL=DI.length; function A(){for(i=0; i-DIL; i++){DIS=DI[ i ].style; DIS.position='absolute'; DIS.left=(Math.sin(R*x1+i*x2+x3)*x4+x5)+"px"; DIS.top=(Math.cos(R*y1+i*y2+y3)*y4+y5)+"px"}R++}setInterval('A()',5); void(0);
```

### 浏览器控制台的 Js
在浏览器的开发者工具中，console 也是个可以执行 JavaScript 代码的地方
```js
// 页面可编辑
document.body.contentEditable='true';
```