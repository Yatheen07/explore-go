package main

import (
	"fmt"
	"runtime"
	"log"
	"os/exec"
	"errors"
	_"os"
)

type tabs struct{
	url,browser string
}

func openInBrowser(url string,browser string){
	var err error

	switch runtime.GOOS {
		case "linux":
			err = errors.New("Linux is not supported.")
		case "windows":
			cmd := exec.Command(browser,url,)
			if err := cmd.Start(); err != nil {
				log.Fatalln("can't open browser", err)
			}
		case "darwin":
			err = errors.New("This OS is not supported.")
		default:
			err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}

func openApplication(path string){
	cmd := exec.Command(path)
	if err:=cmd.Start(); err != nil{
		fmt.Println("Cannot open the application");
	}
}

func main(){

	browsersList := make(map[string]string)
	browsersList["Chrome"] = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	browsersList["Opera"] = "C:\\Users\\prava\\AppData\\Local\\Programs\\Opera\\launcher.exe"

	applicationsList := []string{
		"C:\\Users\\prava\\AppData\\Local\\Programs\\Notion\\Notion.exe",
		"D:\\Program Files\\VS Code\\Microsoft VS Code\\Code.exe",
	}

	tabList := []tabs{
		tabs{"https://google.com",browsersList["Chrome"]},
		tabs{"https://youtube.com",browsersList["Opera"]},
		tabs{"https://mail.zoho.com",browsersList["Chrome"]},
		tabs{"https://people.zoho.com",browsersList["Chrome"]},
	}

	for _,value := range tabList{
		openInBrowser(value.url,value.browser)
	}

	for _,value := range applicationsList{
		openApplication(value)
	}
	
}

