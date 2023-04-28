package main

import "os"
import "fmt"
import "strings"
import "github.com/oppsec/OAO/src/ui"
import "github.com/oppsec/OAO/src/ldapconx"

import "github.com/fatih/color"
import "github.com/jessevdk/go-flags"

func error(err interface{}){
	if err != nil {
		errorMsg := color.New(color.FgRed)
		errorMsg.Println("[-]", err)
		os.Exit(0)
	}
}

func help(){
	fmt.Println("[!] You need to specify a group and module from arguments.")
	fmt.Println(" > e.g: OAO -u domain.local/username:password@IP -g 'Domain Admins' -m add")
}

func main(){
	ui.GetBanner()

	var opts struct {
		User string `short:"u" long:"username" description:"Definition: Username to authenticate in LDAP" required:"true"`
		TargetUser string `short:"t" long:"target-user" description:"Target user to add/rm from group." required:"false"`
		Group string `short:"g" long:"group" description:"Group to add/rem user" required:"false"`
		Module string `short:"m" long:"module" description:"Add or Remove user from a group" required:"false"`
	}

	fmt.Println(opts.TargetUser)

	_, err := flags.Parse(&opts)
	if err != nil { os.Exit(0) }

	errorMsg := color.New(color.FgRed)
	slashIndex := strings.Index(opts.User, "/")
	colonIndex := strings.Index(opts.User, ":")
	atIndex := strings.LastIndex(opts.User, "@")

	if slashIndex == -1 || colonIndex == -1 || atIndex == -1 {
		errorMsg.Println("[-] Use this format: domain.local/username:password@IP")
		return
	}

	domain := opts.User[:slashIndex]
	user := opts.User[slashIndex+1 : colonIndex]
	password := opts.User[colonIndex+1 : atIndex]
	ip := opts.User[atIndex+1:]

	username := fmt.Sprintf("%s@%s", user, domain)

	if len(opts.Group) == 0 && len(opts.Module) == 0 && len(opts.TargetUser) == 0{
		conx.ConnectLDAP(username, password, ip, opts.Group, opts.Module, opts.TargetUser)
	} else if len(opts.Group) > 0 && len(opts.Module) == 0 {
		help()
	} else if len(opts.Module) > 0 && len(opts.Group) == 0 {
		help()
	} else {
		conx.ConnectLDAP(username, password, ip, opts.Group, opts.Module, opts.TargetUser)
	}
}	