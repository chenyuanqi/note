
### OAuth2.0 是什么
OAuth 2.0 是目前最流行的授权机制，用来授权第三方应用，获取用户数据。数据的所有者告诉系统，同意授权第三方应用进入系统，获取这些数据。系统从而产生一个短期的进入令牌（token），用来代替密码，供第三方应用使用。  

OAuth 主要用来颁发令牌（token）,与密码的区别：  
1、令牌是短期的，到期会自动失效，用户自己无法修改。密码一般长期有效，用户不修改，就不会发生变化。  
2、令牌可以被数据所有者撤销，会立即失效。  
3、令牌有权限范围（scope），对于网络服务来说，只读令牌就比读写令牌更安全；密码一般是完整权限。  
4、OAuth 保证了令牌既可以让第三方应用获得权限，同时又随时可控，不会危及系统安全。  

OAuth 2.0 的标准是 RFC 6749 文件。  
OAuth 的核心就是向第三方应用颁发令牌。OAuth 引入了一个授权层，用来分离两种不同的角色：客户端和资源所有者...资源所有者同意以后，资源服务器可以向客户端颁发令牌。客户端通过令牌，去请求数据。

### OAuth 2.0 的四种方式
OAuth 2.0 规定了四种获得令牌的流程：  

- 授权码（authorization-code）：指的是第三方应用先申请一个授权码，然后再用该码获取令牌。这种方式是最常用的流程，安全性也最高，它适用于那些有后端的 Web 应用。授权码通过前端传送，令牌则是储存在后端，而且所有与资源服务器的通信都在后端完成。这样的前后端分离，可以避免令牌泄漏。  
> 第一步，A 网站提供一个链接，用户点击后就会跳转到 B 网站，授权用户数据给 A 网站使用：https://b.com/oauth/authorize?response_type=code&client_id=CLIENT_ID&redirect_uri=CALLBACK_URL&scope=read（response_type 参数表示要求返回授权码，client_id 参数让 B 知道是谁在请求，redirect_uri 参数是 B 接受或拒绝请求后的跳转网址，scope 参数表示要求的授权范围）  
> 第二步，用户跳转后，B 网站会要求用户登录，然后询问是否同意给予 A 网站授权。用户表示同意，这时 B 网站就会跳回 redirect_uri 参数指定的网址。跳转时，会传回一个授权码，就像下面这样：https://a.com/callback?code=AUTHORIZATION_CODE（code 参数就是授权码）  
> 第三步，A 网站拿到授权码以后，就可以在后端，向 B 网站请求令牌：https://b.com/oauth/token?client_id=CLIENT_ID&client_secret=CLIENT_SECRET&grant_type=authorization_code&code=AUTHORIZATION_CODE&redirect_uri=CALLBACK_URL（client_id 参数和 client_secret 参数用来让 B 确认 A 的身份，grant_type 参数的值是 AUTHORIZATION_CODE，表示采用的授权方式是授权码，code 参数是上一步拿到的授权码，redirect_uri 参数是令牌颁发后的回调网址）  
> 第四步，B 网站收到请求以后，就会颁发令牌。具体做法是向 redirect_uri 指定的网址，发送一段 JSON 数据。JSON 数据中，access_token 字段就是令牌，A 网站在后端拿到了。  
- 隐藏式（implicit）：允许直接向前端颁发令牌，这种方式没有授权码这个中间步骤。这种方式把令牌直接传给前端，是很不安全的。因此，只能用于一些安全要求不高的场景，并且令牌的有效期必须非常短，通常就是会话期间（session）有效，浏览器关掉，令牌就失效了。  
> 第一步，A 网站提供一个链接，要求用户跳转到 B 网站，授权用户数据给 A 网站使用:  
> https://b.com/oauth/authorize?response_type=token&client_id=CLIENT_ID&redirect_uri=CALLBACK_URL&scope=read（response_type 参数为 token，表示要求直接返回令牌）  
> 第二步，用户跳转到 B 网站，登录后同意给予 A 网站授权。这时，B 网站就会跳回 redirect_uri 参数指定的跳转网址，并且把令牌作为 URL 参数，传给 A 网站：https://a.com/callback#token=ACCESS_TOKEN（token 参数就是令牌）  
> 注意，令牌的位置是 URL 锚点（fragment），而不是查询字符串（querystring），这是因为 OAuth 2.0 允许跳转网址是 HTTP 协议，因此存在 "中间人攻击" 的风险，而浏览器跳转时，锚点不会发到服务器，就减少了泄漏令牌的风险。
- 密码式（password）：如果你高度信任某个应用，RFC 6749 也允许用户把用户名和密码，直接告诉该应用。这种方式需要用户给出自己的用户名 / 密码，显然风险很大，因此只适用于其他授权方式都无法采用的情况，而且必须是用户高度信任的应用。  
> 第一步，A 网站要求用户提供 B 网站的用户名和密码。拿到以后，A 就直接向 B 请求令牌:  
> https://oauth.b.com/token?grant_type=password&username=USERNAME&password=PASSWORD&client_id=CLIENT_ID（grant_type 参数是授权方式，这里的 password 表示 "密码式"，username 和 password 是 B 的用户名和密码）  
> 第二步，B 网站验证身份通过后，直接给出令牌。注意，这时不需要跳转，而是把令牌放在 JSON 数据里面，作为 HTTP 回应，A 因此拿到令牌。
- 客户端凭证（client credentials）：适用于没有前端的命令行应用，即在命令行下请求令牌。这种方式给出的令牌，是针对第三方应用的，而不是针对用户的，即有可能多个用户共享同一个令牌。  
> 第一步，A 应用在命令行向 B 发出请求：  
> https://oauth.b.com/token?grant_type=client_credentials&client_id=CLIENT_ID&client_secret=CLIENT_SECRET（grant_type 参数等于 client_credentials 表示采用凭证式，client_id 和 client_secret 用来让 B 确认 A 的身份）  
> 第二步，B 网站验证通过以后，直接返回令牌。

**令牌的使用**  
A 网站拿到令牌以后，就可以向 B 网站的 API 请求数据了。  
此时，每个发到 API 的请求，都必须带有令牌。具体做法是在请求的头信息，加上一个 Authorization 字段，令牌就放在这个字段里面。  
```
Authorization: Bearer ACCESS_TOKEN
```

**更新令牌**  
令牌的有效期到了，如果让用户重新走一遍上面的流程，再申请一个新的令牌，很可能体验不好，而且也没有必要，OAuth 2.0 允许用户自动更新令牌。  
具体方法是，B 网站颁发令牌的时候，一次性颁发两个令牌，一个用于获取数据，另一个用于获取新的令牌（refresh token 字段）。令牌到期前，用户使用 refresh token 发一个请求，去更新令牌。  
```
https://b.com/oauth/token?
  grant_type=refresh_token&
  client_id=CLIENT_ID&
  client_secret=CLIENT_SECRET&
  refresh_token=REFRESH_TOKEN
```
上面 URL 中，grant_type 参数为 refresh_token 表示要求更新令牌，client_id 参数和 client_secret 参数用于确认身份，refresh_token 参数就是用于更新令牌的令牌。  
B 网站验证通过以后，就会颁发新的令牌。  