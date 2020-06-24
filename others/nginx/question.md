

### 跨域问题
有时候会遇到一些接口不支持跨域，这时候可以简单的添加 add_headers 来支持 cors 跨域。
```
server {
  listen 80;
  server_name api.xxx.com;
    
  add_header 'Access-Control-Allow-Origin' '*';
  add_header 'Access-Control-Allow-Credentials' 'true';
  add_header 'Access-Control-Allow-Methods' 'GET,POST,HEAD';
  add_header 'Access-Control-Allow-Headers' 'DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,X-XSRF-TOKEN';

  location / {
    proxy_pass http://127.0.0.1:3000;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host  $http_host;    
  } 
}
```

也可以使用 rewrite 指令重定向 URI 来解决跨域问题
```
upstream test {
  server 127.0.0.1:8080;
  server localhost:8081;
}
server {
  listen 80;
  server_name api.xxx.com;
  location / { 
    root  html;                   #去请求../html文件夹里的文件
    index  index.html index.htm;  #首页响应地址
  }
  # 用于拦截请求，匹配任何以 /api/开头的地址，
  # 匹配符合以后，停止往下搜索正则。
  location ^~/api/{ 
    # 代表重写拦截进来的请求，并且只能对域名后边的除去传递的参数外的字符串起作用，
    # 例如www.a.com/proxy/api/msg?meth=1&par=2重写，只对/proxy/api/msg重写。
    # rewrite后面的参数是一个简单的正则 ^/api/(.*)$，
    # $1代表正则中的第一个()，$2代表第二个()的值，以此类推。
    rewrite ^/api/(.*)$ /$1 break;
    
    # 把请求代理到其他主机 
    # 其中 http://www.b.com/ 写法和 http://www.b.com写法的区别如下
    # 如果你的请求地址是他 http://server/html/test.jsp
    # 配置一： http://www.b.com/ 后面有“/” 
    #         将反向代理成 http://www.b.com/html/test.jsp 访问
    # 配置一： http://www.b.com 后面没有有“/” 
    #         将反向代理成 http://www.b.com/test.jsp 访问
    proxy_pass http://test;

    # 如果 proxy_pass  URL 是 http://a.xx.com/platform/ 这种情况
    # proxy_cookie_path应该设置成 /platform/ / (注意两个斜杠之间有空格)。
    proxy_cookie_path /platfrom/ /;

    # http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass_header
    # 设置 Cookie 头通过
    proxy_pass_header Set-Cookie;
  } 
}
```
