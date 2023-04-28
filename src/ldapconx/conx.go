package conx

import "os"
import "fmt"
import "strings"

import "github.com/oppsec/OAO/src/adduser"
import "github.com/fatih/color"
import "github.com/go-ldap/ldap/v3"

func error(err interface{}){
	if err != nil {
		errorMsg := color.New(color.FgRed)
		errorMsg.Println("[-]", err)
		os.Exit(0)
	}
}

func ConnectLDAP(user, password, host, group, module, targetuser string){

	statusMsg := color.New(color.FgYellow)
	successMsg := color.New(color.FgGreen)
	errorMsg := color.New(color.FgRed)


	if len(group) == 0 && len(module) == 0 && len(targetuser) == 0{

		statusMsg.Printf("[!] %s:389\n[!] Username: %s\n", host, user)
		conn, err := ldap.DialURL(fmt.Sprintf("ldap://%s", host))
		error(err)

		defer conn.Close()

		successMsg.Printf("[+] Estabelished TCP connection to %s:389\n", host)

		err = conn.Bind(user, password)
		error(err)

	} else {

		if strings.Contains(module, "add") {
			
			statusMsg.Printf("[!] %s:389\n[!] Username: %s\n", host, user)
			conn, err := ldap.DialURL(fmt.Sprintf("ldap://%s", host))
			error(err)

			defer conn.Close()

			successMsg.Printf("[+] Estabelished TCP connection to %s:389\n", host)
			err = conn.Bind(user, password)
			error(err)
			
			statusMsg.Printf("[!] Adding '%s' to group '%s'\n", targetuser, group)
			oao.GetDomainDN(user, targetuser, group, module, conn)

		} else if strings.Contains(module, "rm") {
			
			statusMsg.Printf("[!] %s:389\n[!] Username: %s\n", host, user)
			conn, err := ldap.DialURL(fmt.Sprintf("ldap://%s", host))
			error(err)

			defer conn.Close()

			successMsg.Printf("[+] Estabelished TCP connection to %s:389\n", host)
			err = conn.Bind(user, password)
			error(err)

			statusMsg.Printf("[!] Deleting '%s' to group '%s'\n", targetuser, group)
			oao.GetDomainDN(user, targetuser, group, module, conn)

		} else {
			if len(targetuser) != 0{
				errorMsg.Println("[-] Specify an user in '--target-user' flag.")
				os.Exit(0)
			}
			errorMsg.Println("[-] Just 'add' or 'rm' to module.")
		}
	}
}