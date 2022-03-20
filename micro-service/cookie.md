cookie是服务器保存在浏览器的一小段信息；

cookie包含的信息：
 Ø Cookie的名字
 Ø Cookie的值（真正的数据写在这里）
 Ø 到期时间
 Ø 所属域名（默认当前域名）
 Ø 生效的路径（默认是当前网址）
根据以上信息，这个Cookie对该域名下的生效路径及子路径有效。
之后，浏览器一旦访问这个路径，浏览器就会附加这段Cookie发送给服务器。

浏览器的同源策略规定，两个网址只要域名相同和端口相同，就可以共享Cookie。

 1. 服务器设置HTTP头Set-Cookie字段。可以设置多个Set-Cookie。一个Set-Cookie中设置多个键值。
 2. 浏览器想服务器发送HTTP请求，HTTP头信息的Cookie字段，可以包含多个Cookie，用；隔开。

服务器收到Cookie时，无法知道的内容：
 Ø cookie的各种属性，过期时间
 Ø 那个域名设置的cookie，是一级域名还是二级域名

Cookie属性：
 · Expires：指定具体的到期时间
 · Max-Age： 倒计时时间
 · Domain：域名
 · Path：路径
 · Secure：指定在加密协议HTTPS下，才能将Cookie发送给浏览器
 · document.cookie
