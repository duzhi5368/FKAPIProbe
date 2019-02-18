# FKAPIProbe

这是FreeKnight自用的简易Web API压测工具。支持Linux平台和Windows平台。

## USAGE:

FKAPIProbe.exe -h

## VERSION:
   0.1.0

## GLOBAL OPTIONS:
-   --debug           if open the debug mode
-   -c value          number of concurrent connections to use (default: 10)
-   -d value          duration of test in seconds (default: 10)
-   -t value          socket/request timeout in (ms) (default: 1000)
-   --keepalive       if keep-alives
-   --dataFile value  load the par which store in the file
-   -m value          http method, use 'GET' or 'POST'  (default: "GET")
-   -H value          the http headers sent to the target url, e.g:'keyA:ValueA;KeyB:ValueB'
-   --compress        if prevents sending the "Accept-Encoding: gzip" header
-   --help, -h        show help
-   --version, -v     print the version

## OUTPUT:

```
将使用 5 个协程同时调用 https://XXXOOO.com:80/XXX/OOO 持续 2 秒.
压测完成. 收到 5 个协程测试结果.
--------------------------------------
2.04651706s 秒内共发送 90 个请求
每秒平均请求次数:	43.98 次
单协程每秒平均请求次数:	8.80 次
收到总字节数 62.25 KB
平均每秒收到数据:	30.42 KB
平均响应时间:	113.695392ms
最快响应时间:	110.0062ms
最慢响应时间:	130.0075ms
请求异常数量:	0 个
成功请求数量:	90 个，占比100.0%!
请求错误数量:	0 个（服务器内部错误）
--------------------------------------
```