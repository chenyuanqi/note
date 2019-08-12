
### 前端数据存储方案
前端数据存储主要是对 token 存储、 代码缓存、 图片存储等等。  
主要的存储方案是 cookie、localStorage， sessionStorage，websql 和 indexeddb。  

#### cookie
Cookie 翻译过来就是 饼干、小甜点 的意思。  
是在 web 端常见的存储方式之一，而且在发起 http 请求的时候会自动被带上。  

cookie 可以手动设置，也可以由服务器产生，当客户端（浏览器）向服务器发送请求，服务器会反馈一些信息给客户端，这些信息的 key/value 值被浏览器作为文件保存在客户端特定的文件夹中。 

cookie 的属性参数说明  
1）name：cookie 的名称  
2）value：cookie 的值  
3）maxAge：cookie 的最大有效期  
4）secure：是否使用安全协议传输  
5）path：cookie 使用的路径，不同路径无法获取到  
6）domain：cookie 的域，和 path 类似，主要防止跨域攻击  
7）comment：该 cookie 的说明  
8）version：cookie 的版本  

```javascript
// 写如 cookie
let setCookie = (name, value, expires) = > {
	let date = new Date()
	// 这里 expires 单位是天
	data.setDate(data.getDate() + expires)
	document.cookie = name + '=' + value + ';expires=' + date
} 
	
// 读取 cookie
let getCookie = (name) => {
	let cookies = document.cookie
	let cookieArr = cookies.split(';')
	for(let i = 0; i < cookieArr.length; i++) {
		let arr = cookieArr[i].split('=')
		if (name == arr[0] ) {
			return arr[1]
		}
	}

	return false
}

// 删除 cookie
let removeCookie = (name) => {
	setCookie(name, '', '-1')
}
```

cookie 的使用注意事项  
1、cookie 是有大小限制的，每个 cookie 所存放的数据不能超过 4k， 如果超过则该属性将返回空字符串  
2、通常我们需要将存放的 value 值进行 escape 编码，在使用的时候再用 unescape () 函数反编码  
3、重要的信息不要存放在 cookie 里  

#### localStorage 
localStorage 是被 W3C 标准化之后的网页存储的一种方式，原本是属于 HTML5 的存储方案，后来被独立出来，适用于 IE 8+，FF 2+，Safari 4+，chrome 4+， Opera 10.50+。  
`注意：IE9 localStorage 不支持本地文件，需要将项目署到服务器才可以支持`  
```javascript
// 判断浏览器是否支持 localStorage
if (window.localStorage) {
    alert('This browser supports localStorage');
} else {
    alert('This browser does NOT support localStorage');
}
```
在浏览器的开发者工具的 Application 下，我们能轻易看到 Storage > Local Storage，那就是 localStorage 的管理界面。  

localStorage 的特点  
1、存储时间：localStorage 的存储周期比 sessionStorage 长，如果用户不清理的话，是可以永久存储的  
2、访问限制：localStorage 的访问域默认设定为设置 localStorage 的当前域，其他域名不可取，localStorage 设定后新开 tab 可以访问  
3、大小限制及检测：标准建议对于每个 domain 的 localStorage 大小为 5M，达到限制时浏览器可以去询问用户是否允许增加存储空间  
4、localStorage 可以提前存放服务器加载的文件，加快页面加载速度并减少服务器压力  
```javascript
// 读取
localStorage.getItem(key)
// 根据索引读取
localStorage.key(index);
// 存储
localStorage.setItem(key,value) 
// 删除
localStorage.removeItem(key)
// 清空
localStorage.clear();
```

#### sessionStorage
sessionStorage 和 localStorage 的区别在于用户关掉了浏览器的当前页，sessionStorage 存储的数据就会被销毁掉。  
sessionStorage 访问存在限制，sessionStorage 是一个 tab 级别的存储，即当前 tab 存储的 sessionStorage，只能在当前 tab 页才能访问。
sessionStorage 一般用于存储当前 tab 页刷新还保留的数据，比如音视频的播放进度。  

```javascript
// 读取
sessionStorage.getItem(key)
// 存储
sessionStorage.setItem(key,value) 
// 删除
sessionStorage.removeItem(key)
// 清空
sessionStorage.clear();
```

#### websql && indexeddb
websql 像是关系型数据库，并且能够使用 sql 语句进行操作。  
虽然现在 websql 已经不再继续维护了，但是因为它的起步早，所以兼容性是非常的好。  

websql 和 indexedDB 在访问权限上和 localStorage 是一致的，都可以跨 tab 访问，并且只能创建数据库下的域名才能访问（不能指定）。  
websql 和 indexedDB 存储时间是永久的，除非用户清除数据。  
IndexedDB 的储存空间比 LocalStorage 大得多，一般来说不少于 250MB，甚至没有上限。  
websql 和 IndexDB 提供搜索功能，允许建立自定义的索引。  
websql 和 indexedDB 主要用于离线应用。  

```javascript
// 打开 websql
var db = openDatabase('mydb', '1.0', 'Test DB', 2 * 1024 * 1024,fn); //openDatabase() 方法对应的五个参数分别为：数据库名称、版本号、描述文本、数据库大小、创建回调  

// 查询
var db = openDatabase('mydb', '1.0', 'Test DB', 2 * 1024 * 1024);
db.transaction(function (tx) {
    tx.executeSql('CREATE TABLE IF NOT EXISTS WIN (id unique, name)');
});

// 写入
var db = openDatabase('mydb', '1.0', 'Test DB', 2 * 1024 * 1024);
db.transaction(function (tx) {
	tx.executeSql('CREATE TABLE IF NOT EXISTS WIN (id unique, name)');
	tx.executeSql('INSERT INTO WIN (id, name) VALUES (1, "winty")');
	tx.executeSql('INSERT INTO WIN (id, name) VALUES (2, "LuckyWinty")');
});

// 读取
db.transaction(function (tx) {
	tx.executeSql('SELECT * FROM WIN', [], function (tx, results) { 
		var len = results.rows.length, i;
	    msg = "<p>查询记录条数: " + len + "</p>"; 
	    document.querySelector('#status').innerHTML += msg; 
	    for (i = 0; i < len; i++){
		    alert(results.rows.item(i).name );
		}
	}, null);
});
```

[websql 的使用参考](https://www.ibm.com/developerworks/cn/web/1108_zhaifeng_websqldb/index.html)  
[indexedDB 的使用参考](https://www.zhangxinxu.com/wordpress/2017/07/html5-indexeddb-js-example/)  
