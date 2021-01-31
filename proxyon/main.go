package main

import (
	_ "Proxyone/initone"
	"Proxyone/proxy"
	"flag"
)

func main(){
	var xz string
	flag.StringVar(&xz,"i","","请选择你要进行的模式")
	flag.Parse()
	switch xz {
	case "1":
		proxy.Xmjtproks()
	}

}