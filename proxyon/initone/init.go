package initone

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

type Config struct {
	Xmapi string
	Orderno string
	Secret string
	Zmapi string
}
var InitConfig = &Config{}

func init(){
	bytes,err := ioutil.ReadFile("./unit.config")
	if err != nil {
		fmt.Println("配置文件读取失败 : %s", err)
		return
	}
	Xmapivalue := gjson.GetBytes(bytes,"Xmapi")
	osrderno := gjson.GetBytes(bytes,"orderno")
	secret := gjson.GetBytes(bytes,"secret")
	zmapi := gjson.GetBytes(bytes,"Zmapi")

	InitConfig.Xmapi = Xmapivalue.String()
	InitConfig.Orderno = osrderno.String()
	InitConfig.Secret = secret.String()
	InitConfig.Zmapi = zmapi.String()
	fmt.Println("检测状态中")
	if (InitConfig.Xmapi == ""){
		fmt.Println("熊猫API地址                  [X]")
	}else{
		fmt.Println("熊猫API地址                  [*]")
	}
	if (InitConfig.Orderno == ""){
		fmt.Println("熊猫Orderno地址              [X]")
	}else{
		fmt.Println("熊猫Orderno地址              [*]")
	}
	if (InitConfig.Secret == ""){
		fmt.Println("熊猫Secret地址               [X]")
	}else{
		fmt.Println("熊猫Secret地址               [*]")
	}
	if (InitConfig.Zmapi == ""){
		fmt.Println("芝麻API地址                  [X]")
	}else{
		fmt.Println("芝麻API地址                  [*]")
	}
}
