# fuckhost
该工具需要配合httpx 1.0版本一下进行使用

https://github.com/projectdiscovery/httpx/releases

目的是为了匹配出来只能使用绑定的host进行访问的ip

使用办法

echo "192.1680.0.0/24" | ./httpx -silent | ./fuckhost

即可模拟
