package ui

import "io/ioutil"
import "github.com/fatih/color"

func GetBanner(){

	bannerMsg := color.New(color.FgCyan)

	arq, err := ioutil.ReadFile("src/ui/banner.txt")
	if err != nil { panic(err) }

	bannerMsg.Println(string(arq))
}