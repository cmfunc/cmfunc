物联网通信协议，构建于TCP/IP协议上。MQTT最大的优点在于，可以以极少的代码和优先的带宽，为连接远程设备提供实时可靠的消息服务。

https://github.com/eclipse/paho.mqtt.golang

https://www.runoob.com/w3cnote/mqtt-intro.html#:~:text=MQTT%E6%98%AF%E4%B8%80%E4%B8%AA%E5%9F%BA%E4%BA%8E%E5%AE%A2%E6%88%B7,%E8%AE%BE%E5%A4%87%E4%B8%AD%E5%B7%B2%E5%B9%BF%E6%B3%9B%E4%BD%BF%E7%94%A8%E3%80%82

MQTT中有三种身份：发布者（Publish）、代理（Broker）（服务器）、订阅者（Subscribe）。
消息的发布和订阅都是客户端，消息代理是服务器，消息发布者可以同时是订阅者。
MQTT传输的消息分为：主题（Topic）和负载（payload）两部分。
	• topic，消息的类型，订阅者订阅（Subscribe）后，就会收到该主题的消息内容（payload）。
	• payload，消息的内容，指订阅者具体要使用的内容。
	
	
MQTT构建有序、无损、基于字节流的双向传输。
当应用数据通过MQTT网络发送时，MQTT会把与之相关的服务质量（QoS）和主题名（Topic）相关连。
