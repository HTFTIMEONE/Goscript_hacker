# FTPsmall



### 1，文件搜索
 第一个功能(文件搜索)
	htf.exe -i 1 -wj 123.txt (不指定路径的情况下，默认搜索全硬盘)
	htf.exe -i 1 -wj 123.txt -ispwd c:/aaa/ （去指定目录搜索）

### 2，ftp上传

 第二个功能(ftp上传)
	htf.exe -i 2 -f C:\1.txt -ip FTP地址 -p FTP端口 -user FTP用户 -pwd FTP密码
### 3，socket上传
 第三个功能(socket上传)
	首先cl执行监听
	 	./cl -ip 0.0.0.0 -p 1222
	然后htf.exe去上传
		htf.exe -i 3 -f C:\11231.txt -ip ip -p 端口
### 4，屏幕截图
第四个功能(屏幕截图)
htf.exe -i 4 -f 123.png
直接将截图保存为123.png
