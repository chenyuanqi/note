
### 文件下载
前端文件下载：  

1）form 表单提交  
为一个下载按钮添加 click 事件，点击时动态生成一个表单，利用表单提交的功能来实现文件的下载
```js
/**
 * 下载文件
 * @param {String} path - 请求的地址
 * @param {String} fileName - 文件名
 */
function downloadFile (downloadUrl, fileName) {
    // 创建表单
    const formObj = document.createElement('form');
    formObj.action = downloadUrl;
    formObj.method = 'get';
    formObj.style.display = 'none';
    // 创建input，主要是起传参作用
    const formItem = document.createElement('input');
    formItem.value = fileName; // 传参的值
    formItem.name = 'fileName'; // 传参的字段名
    // 插入到网页中
    formObj.appendChild(formItem);
    document.body.appendChild(formObj);
    formObj.submit(); // 发送请求
    document.body.removeChild(formObj); // 发送完清除掉
}
```

2）a 标签的 download 属性   
```html
<a href="example.jpg" download>点击下载</a>

<a href="example.jpg" download="test">点击下载</a> // 指定文件名

<script type="text/javascript">
	// 检测浏览器是否支持download属性
    const isSupport = 'download' in document.createElement('a');
</script>
```

3）open或location.href  
本质上和 a 标签访问下载链接一样
```js
window.open('downloadFile.zip');
location.href = 'downloadFile.zip';
```

4）Blob 对象  
调用 api，将文件流转为 Blob 二进制对象。思路是发请求获取二进制数据，转化为 Blob 对象，利用 URL.createObjectUrl 生成 url 地址，赋值在 a 标签的 href 属性上，结合 download 进行下载。
`注：IE10以下不支持。`
```js
/**
 * 下载文件
 * @param {String} path - 下载地址/下载请求地址。
 * @param {String} name - 下载文件的名字/重命名（考虑到兼容性问题，最好加上后缀名）
 */
downloadFile (path, name) {
    const xhr = new XMLHttpRequest();
    xhr.open('get', path);
    xhr.responseType = 'blob';
    xhr.send();
    xhr.onload = function () {
        if (this.status === 200 || this.status === 304) {
            // 如果是IE10及以上，不支持download属性，采用msSaveOrOpenBlob方法，但是IE10以下也不支持msSaveOrOpenBlob
            if ('msSaveOrOpenBlob' in navigator) {
                navigator.msSaveOrOpenBlob(this.response, name);
                return;
            }
            /* 
              如果发送请求时不设置xhr.responseType = 'blob'，
              默认ajax请求会返回DOMString类型的数据，即字符串。
              此时需要使用两处注释的代码，对返回的文本转化为Blob对象，然后创建blob url，
              此时需要注释掉原本的const url = URL.createObjectURL(target.response)。
            */
            /* 
            const blob = new Blob([this.response], { type: xhr.getResponseHeader('Content-Type') });
            const url = URL.createObjectURL(blob);
            */
            const url = URL.createObjectURL(this.response); // 使用上面则注释此行
            const a = document.createElement('a');
            a.style.display = 'none';
            a.href = url;
            a.download = name;
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
            URL.revokeObjectURL(url);
        }
    };
}

// 上面方法本地测试有时会有跨域问题，下面使用axios重写一下
// 已经配置好proxy
downloadFile (path, name) {
    axios.get({
      url: path,
      method: 'get'
    }).then(res => {
      const blob = new Blob([res.data], { type: res.headers['content-type'] });
      const url = URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.style.display = 'none';
      a.href = url;
      a.download = name;
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
      URL.revokeObjectURL(url);
    })
}
```
该方法不能缺少 a 标签的 download 属性的设置。  
因为发请求时已设置返回数据类型为 Blob 类型(xhr.responseType = 'blob')，所以 target.response 就是一个 Blob 对象，打印出来会看到两个属性 size 和 type。虽然 type 属性已指定了文件的类型，但是为了稳妥起见，还是在 download 属性值里指定后缀名，如 Firefox 不指定下载下来的文件就会不识别类型。

5）利用 Base64  
用法跟上面用Blob大同小异，基本上思路一样。  
不同点： 上面是利用 Blob 对象生成 Blob URL， 这里则是生成 Data URL，即 base64 编码后的 url 形式。  
```js
/**
 * 下载文件
 * @param {String} path - 下载地址/下载请求地址。
 * @param {String} name - 下载文件的名字（考虑到兼容性问题，最好加上后缀名）
 */
downloadFile (path, name) {
    const xhr = new XMLHttpRequest();
    xhr.open('get', path);
    xhr.responseType = 'blob';
    xhr.send();
    xhr.onload = function () {
        if (this.status === 200 || this.status === 304) {
            const fileReader = new FileReader();
            fileReader.readAsDataURL(this.response);
            fileReader.onload = function () {
                const a = document.createElement('a');
                a.style.display = 'none';
                a.href = this.result;
                a.download = name;
                document.body.appendChild(a);
                a.click();
                document.body.removeChild(a);
            };
        }
    };
}
```

### 如何获取文件名
返回文件流的时候，在浏览器上观察接口返回的信息，会看到有这么一个 header：Content-Disposition，其中包含了文件名：filename= 和 filename\*= 可以截取这段字符串中的这两个字段值了。
```js
// xhr是XMLHttpRequest对象
const content = xhr.getResponseHeader('content-disposition'); // 注意是全小写，自定义的header也是全小写
if (content) {
    let name1 = content.match(/filename=(.*);/)[1]; // 获取filename的值
    let name2 = content.match(/filename*=(.*)/)[1]; // 获取filename*的值
    name1 = decodeURIComponent(name1);
    name2 = decodeURIComponent(name2.substring(6)); // 这个下标6就是UTF-8''
}
```



