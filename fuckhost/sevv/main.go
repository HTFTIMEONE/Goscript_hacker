package main

import (
	"bufio"
	"fmt"
	"os"
	"sevv/httpd"
	"sync"
)

func main(){
	hosts := httpd.Redfile("./host.txt")
	for _,val:= range hosts{
		dxc(val)
	}
}
func stdio()[]string{
	var caa []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == "bye" {
			break
		}
		caa =append(caa,line)
	}
	return caa
}
func dxc(hostname string){
	ca := stdio()
	wg := sync.WaitGroup{}
	wg.Add(len(ca))
	for i := 0; i < len(ca); i++ {
		go func(i int) {
			b,title := httpd.Reqs(ca[i],hostname)
			if b {
				fmt.Print("[*] Address: "+ca[i]+" title: "+title+"\n")
			}else{
				fmt.Print("[-] "+ca[i]+" no"+"\n")
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

}