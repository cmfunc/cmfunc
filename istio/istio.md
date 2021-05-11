https://istio.io/latest/zh/docs/concepts/what-is-istio/

将限流、熔断、降低、监控、链路追踪，都放在一个代理层；
代理层对服务做转发，有点类似python中wsgi接口协议，开发人员不用重新开发server（gunicron、uWSGI服务器、fcgi、bjoern），只用开发application部分的代码。
uwsgi也是通信协议，由uWSGI服务器独占。
https://www.fullstackpython.com/wsgi-servers.html#:~:text=%20WSGI%20servers%20learning%20checklist%20%201%20Understand,requests%20to%20the%20WSGI%20server%20for...%20More%20


