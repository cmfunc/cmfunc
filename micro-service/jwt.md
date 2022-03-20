JSON Web Token
jwt中可以解码出user_id；
<http://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html>

跨域认证解决方案；用户认证；
 1. 用户向服务器发送用户名和密码；
 2. 服务器验证通过后，在当前对话（session）里面保存相关数据，比如用户角色、登录时间等等；
 3. 服务器向用户返回一个session_id，写入用户的Cookie。
 4. 用户随后的每一次请求，都会通过Cookie，将session_id传回服务器。
 5. 服务器收到session_id，找到前期保存的数据，由此得知用户的身份。

JWT：服务器不保存数据，所有数据都保存在客户端，每次请求都发回服务器。
 
JWT原理
服务认证以后，生成一个JSON对象，发回给用户；
{
 "name":"jobs",
 "role":"admin",
 "deadline":"20990922"
}
之后，用户与服务端通信的时候，都要回传JSON对象。服务器全靠这个对象认证用户身份。为了防止用户篡改数据，服务器在生成这个对象时，会加上签名。

服务器不保存任何session数据，服务器变成无状态。

JWT是一个很长的字符串，中间用  .   分割成三部分。JWT内部没有换行。
JWT三部分组成：Header.Payload.Signature
 • Header（头部）
 • Payload（负载）
 • Signature（签名）

Header部分是一个JSON对象，描述JWT的元数据：
{
  "alg": "HS256",
  "typ": "JWT"
}
alg属性表示签名的算法，默认HMAC SHA256；
typ属性表示这个令牌（token）的类型（type），JWT令牌统一协程（JWT）。
最后将上面的JSON对象使用Base64URL转成字符串。

Payload部分也是一个JSON对象，用来存放实际需要传递的数据。
JWT规定了7个官方字段，供选用：
 • iss（issuer）:签发人
 • exp（expiration time）：过期时间
 • sub（subject）：主题
 • aud（audience）：受众
 • nbf（Not Before）：生效时间
 • iat（Issued At）：签发时间
 • jti（JWT ID）：编号
除了官方字段，还可以自定义字段。
JWT默认是不加密的，任何人都可以读到，所以不能把秘密信息放在这部分。
Payload对象也要使用Base64URL算法转成字符串。

Signature部分是两部分的签名，防止数据篡改。
 1. 指定一个密钥（secret）。这个密钥只有服务器才知道，不能泄露给用户。
 2. 使用Header里的指定签名算法（默认 HMAC SHA256），按照公式进行签名
 HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
 3. 算出签名后，把Header、Payload、Signature三部分拼成一个字符串，每部分用 . 隔开。
 4. 返回给用户

JWT使用方式
客户端收到服务器返回的JWT，可以存储在Cookie和localStorage。
之后，客户端每次都要带上JWT。HTTP请求头的Authorization字段Bearer<token>

JWT特点
 1. JWT默认不加密，生成原始token以后，可以赢秘钥再加密一次；
 2. JWT不加密的情况下，不能将秘密数据写入JWT；
 3. JWT不仅可用于认证，也可用于交换数据，有效使用JWT，可减少服务器查询数据库次数；
 4. JWT最大缺点，由于服务器不保存session状态，因此无法在使用过程中废止某个token，或者更改token的权限。一旦JWT签发，在到期之前就会始终有效，除非服务器部署额外的逻辑。
 5. JWT本身包含认证信息，一旦泄露，任何人都可以获得该令牌的所有权限。为了减少盗用，JWT的有效期应该设置得比较短。对于重要权限，使用时应该再次对用户进行认证。
 6. 为了减少盗用，JWT不应该使用HTTP协议明码传输，要使用HTTPS协议传输。

JWT特点
