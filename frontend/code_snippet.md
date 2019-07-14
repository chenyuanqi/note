
### 【调试】显示页面边框
```js
[].forEach.call($$("*"),function(a){ a.style.outline="1px solid #"+(~~(Math.random()*(1<<24))).toString(16) }) 
```

### 获取图片原始宽高
```js
function loadImageAsync(url) {
    return new Promise(function(resolve, reject) {
        var image = new Image();

        image.onload = function() {
            var obj = {
                w: image.naturalWidth,
                h: image.naturalHeight
            }
            resolve(obj);
        };

        image.onerror = function() {
            reject(new Error('Could not load image at ' + url));
        };
        image.src = url;
    });
}
```

### cookie 操作
```js
function getCookie(name) {
    var arr = document.cookie.match(new RegExp("(^| )" + name + "=([^;]*)(;|$)"));
    if (arr != null) return unescape(arr[2]);
    return null
}

function setCookie(name, value, Hours) {
    var d = new Date();
    var offset = 8;
    var utc = d.getTime() + (d.getTimezoneOffset() * 60000);
    var nd = utc + (3600000 * offset);
    var exp = new Date(nd);
    exp.setTime(exp.getTime() + Hours * 60 * 60 * 1000);
    document.cookie = name + "=" + escape(value) + ";path=/;expires=" + exp.toGMTString() + ";domain=360doc.com;"
}
```

### 判断是否数字类型
```js
function isDigit(value) {
    var patrn = /^[0-9]*$/;
    if (patrn.exec(value) == null || value == "") {
        return false
    } else {
        return true
    }
}
```

### 判断日期是否有效
```js
function isValidDate(value, userFormat='mm/dd/yyyy') {
    const delimiter = /[^mdy]/.exec(userFormat)[0];
    const theFormat = userFormat.split(delimiter);
    const theDate = value.split(delimiter);
    function isDate(date, format) {
	let m, d, y, i = 0, len = format.length, f;
	for (i; i < len; i++) {
	    f = format[i];
	    if (/m/.test(f)) m = date[i];
	    if (/d/.test(f)) d = date[i];
	    if (/y/.test(f)) y = date[i];
	}
	return (
	    m > 0 && m < 13 &&
	    y && y.length === 4 &&
	    d > 0 &&
	    d <= (new Date(y, m, 0)).getDate()
	);
    }

    return isDate(theDate, theFormat);
}
```

### 日期格式转换
```js
Date.prototype.Format = function(formatStr) {
    var str = formatStr;
    var Week = ['日', '一', '二', '三', '四', '五', '六'];
    str = str.replace(/yyyy|YYYY/, this.getFullYear());
    str = str.replace(/yy|YY/, (this.getYear() % 100) > 9 ? (this.getYear() % 100).toString() : '0' + (this.getYear() % 100));
    str = str.replace(/MM/, (this.getMonth() + 1) > 9 ? (this.getMonth() + 1).toString() : '0' + (this.getMonth() + 1));
    str = str.replace(/M/g, (this.getMonth() + 1));
    str = str.replace(/w|W/g, Week[this.getDay()]);
    str = str.replace(/dd|DD/, this.getDate() > 9 ? this.getDate().toString() : '0' + this.getDate());
    str = str.replace(/d|D/g, this.getDate());
    str = str.replace(/hh|HH/, this.getHours() > 9 ? this.getHours().toString() : '0' + this.getHours());
    str = str.replace(/h|H/g, this.getHours());
    str = str.replace(/mm/, this.getMinutes() > 9 ? this.getMinutes().toString() : '0' + this.getMinutes());
    str = str.replace(/m/g, this.getMinutes());
    str = str.replace(/ss|SS/, this.getSeconds() > 9 ? this.getSeconds().toString() : '0' + this.getSeconds());
    str = str.replace(/s|S/g, this.getSeconds());
    return str
}
```

### 检验 URL 链接是否有效
```js
function getUrlState(URL){ 
    var xmlhttp = new ActiveXObject("microsoft.xmlhttp"); 
    xmlhttp.Open("GET",URL, false);  
    try{  
            xmlhttp.Send(); 
    }catch(e){
    }finally{ 
        var result = xmlhttp.responseText; 
        if(result){
            if(xmlhttp.Status==200){ 
                return(true); 
             }else{ 
                   return(false); 
             } 
         }else{ 
             return(false); 
         } 
    }
}
```

### 获取当前路径
```js
function getCurrentPageUrl(){
	var currentPageUrl = "";
	if (typeof this.href === "undefined") {
	    currentPageUrl = document.location.toString().toLowerCase();
	}else {
	    currentPageUrl = this.href.toString().toLowerCase();
	}

	return currentPageUrl;
}
```

### 获取 url 中 get 参数
```js
function getParams(){
	querystr = window.location.href.split("?")
    if(querystr[1]){
        GETs = querystr[1].split("&");
        GET = [];
        for(i=0;i<GETs.length;i++){
              tmp_arr = GETs.split("=")
              key=tmp_arr[0]
              GET[key] = tmp_arr[1]
        }
    }
    return querystr[1];
}
```

### 获取页面宽高
```js
function getPageHeight(){
    var g = document, a = g.body, f = g.documentElement, d = g.compatMode == "BackCompat"
                    ? a
                    : g.documentElement;
    return Math.max(f.scrollHeight, a.scrollHeight, d.clientHeight);
}

// 获取可视高度
function getPageViewWidth(){
    var d = document, a = d.compatMode == "BackCompat" ? 
       				   d.body: d.documentElement;
    return a.clientWidth;
}

function getPageWidth(){
    var g = document, a = g.body, f = g.documentElement, d = g.compatMode == "BackCompat"?
    					  a: g.documentElement;
    return Math.max(f.scrollWidth, a.scrollWidth, d.clientWidth);
}
```