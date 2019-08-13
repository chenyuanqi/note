
### 为什么是 JWT
因为 HTTP 是无状态的协议（即无法通过 HTTP 来标识出不同的用户），所以判断是不是特定用户就需要用户认证了。  

常见的用户认证有 cookie、session 和 JWT（Token）这三种形式。  
cookie 是早期的用户认证机制，不过由于安全系数太低而几近消失。session 就是解决安全性的问题而生，session 的用户认证会在服务器记录用户信息而把 session_id 抛给客户端（存在过期时间）；但是，session 对服务器的开销不小，并且扩展性差，即便有过期时间也无法确保足够的安全。  

JWT（Json Web Token）是一个开放标准，它规定通信双方使用 JSON 对象的形式进行安全通信，它可以使用非对称加密算法对数据进行加密，保证数据的安全。JWT 方式将用户状态分散到了客户端中，可以明显减轻服务端的内存压力。  

### JWT 的组成结构
JWT 由三个部分组成，分别是 Header 头部、Payload 载荷、Signature 签名。  
```
# Header
# 头部一般包含两个部分，typ 定义 JWT 的类型，alg 定义 JWT 使用的加密算法
base64enc({
    "alg": "HS256",
    "typ": "JWT"
});

# Payload（用来放信息的地方，常见的是非敏感信息）
# iss 定义 Token 的签发者，iat 定义签发时间，exp 定义过期时间，过期时间必须在当前时间之后，aud 定义接受方，sub 定义 Token 的主题
base64({
    "iss": "toptal.com",
    "exp": 11111111111,
    "company": "Toptal",
    "awesome": true
});

# Signature
# 签名目的是为了保证 JWT 没有被人篡改过，保证信息的安全性和可靠性
HMACSHA256(base64enc(header) + "." + base64enc(payload), secretKey);
```

### JWT 使用流程
![jwt-usage](https://s2.ax1x.com/2019/08/13/m9FSAI.png)  

客户端收到服务器返回的 JWT，可以储存在 Cookie 里面，也可以储存在 localStorage。之后，客户端每次与服务器通信，都要带上这个 JWT（更好的做法是放在 HTTP 请求的头信息 Authorization 字段里面）。
```
Authorization: Bearer <token>
```

### JWT 的特点
1）JWT 默认是不加密，但也是可以加密的。生成原始 Token 以后，可以用密钥再加密一次。

2）JWT 不加密的情况下，不能将秘密数据写入 JWT。

3）JWT 不仅可以用于认证，也可以用于交换信息。有效使用 JWT，可以降低服务器查询数据库的次数。

4）JWT 的最大缺点是，由于服务器不保存 session 状态，因此无法在使用过程中废止某个 token，或者更改 token 的权限。也就是说，一旦 JWT 签发了，在到期之前就会始终有效，除非服务器部署额外的逻辑。

5）JWT 本身包含了认证信息，一旦泄露，任何人都可以获得该令牌的所有权限。为了减少盗用，JWT 的有效期应该设置得比较短。对于一些比较重要的权限，使用时应该再次对用户进行认证。

6）为了减少盗用，JWT 不应该使用 HTTP 协议明码传输，要使用 HTTPS 协议传输。