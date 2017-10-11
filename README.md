# 模拟CDN重定向性能测试
- 使用方法：    
1.下载所有脚本到任意文件夹    
2.go build 0_test_demo_redirect-To-lambdaPage.go，该脚本引入了Go语言标准库中的net/http，提供web服务，端口号为5002，可以响应并处理客户端（浏览器）的HTTP请求，此时在浏览器打开http://162.3.111.24:5002/p2，会触发远程lambda函数，获取结果，并重定位到http://114.115.145.32:8081/，模拟CDN服务，0_test_demo_redirect-To-baidu.go脚本可以重定位到www.baidu.com其他不变
3.打开另一个终端，运行脚本0_test_demo_auto_selectTime.sh，该脚本自动统计time curl http://162.3.111.24:5002/p2和／p3的时间并将结果写入文件output.txt
