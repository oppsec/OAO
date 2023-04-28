package oao

import "os"
import "fmt"
import "strings"

import "github.com/fatih/color"
import "github.com/go-ldap/ldap/v3"

func error(err interface{}){
	if err != nil {
		errorMsg := color.New(color.FgRed)
		errorMsg.Println("[-]", err)
		os.Exit(0)
	}
}

func adduser(domaindn, user, targetuser, group, module string, conn *ldap.Conn){

	successMsg := color.New(color.FgGreen)
	errorMsg := color.New(color.FgRed)

	userSearchFilter := fmt.Sprintf("(&(objectClass=user)(objectCategory=person)(sAMAccountName=%s))", targetuser)
	userSearchBase := domaindn
	userSearchAttributes := []string{"dn"}

	userSearchRequest := ldap.NewSearchRequest(
		userSearchBase,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		userSearchFilter,
		userSearchAttributes,
		nil,
	)

	userSearchResult, err := conn.Search(userSearchRequest)
	if err != nil || len(userSearchResult.Entries) != 1 {
		errorMsg.Printf("[-] Error to find user '%s' -> %s", targetuser, userSearchFilter)
		os.Exit(0)
	}

	userDN := userSearchResult.Entries[0].DN

	groupSearchFilter := fmt.Sprintf("(&(objectClass=group)(cn=%s))", group)
	groupSearchBase := domaindn

	groupSearchRequest := ldap.NewSearchRequest(
		groupSearchBase,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		groupSearchFilter,
		[]string{"dn"},
		nil,
	)

	groupSearchResult, err := conn.Search(groupSearchRequest)
	if err != nil || len(groupSearchResult.Entries) != 1 {
		errorMsg.Printf("[-] Error to find group '%s'\n", group)
		os.Exit(0)
	}

	groupDN := groupSearchResult.Entries[0].DN

	if module == "add"{
		modify := ldap.NewModifyRequest(groupDN, nil)
		modify.Add("member", []string{userDN})

		err = conn.Modify(modify)
		if err != nil {
			if strings.Contains(string(err.Error()), "Entry Already Exists") == true {
				errorMsg.Println("[-] The user is already in the group.")
			}
			os.Exit(0)
		}

		successMsg.Printf("[+] User '%s' successfuly added to the group %s!\n", targetuser, group)

	} else if module == "rm"{
		modify := ldap.NewModifyRequest(groupDN, nil)
		modify.Delete("member", []string{userDN})

		err = conn.Modify(modify)
		if err != nil {
			if strings.Contains(string(err.Error()), "Unwilling To Perform") == true {
				errorMsg.Println("[-] The user is already outside the group.")
			}
			os.Exit(0)
		}

		successMsg.Printf("[+] User '%s' successfuly removed to the group %s!\n", targetuser, group)
	}

}

func GetDomainDN(user, targetuser, group, module string, conn *ldap.Conn){

	errorMsg := color.New(color.FgRed)
	searchRequest := ldap.NewSearchRequest(
		"",
		ldap.ScopeBaseObject,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(&(objectClass=domain)(objectClass=domainDNS))", // query
		[]string{"defaultNamingContext"},
		nil,
	)

	sr, err := conn.Search(searchRequest)
	error(err)

	if len(sr.Entries) != 1 {
		errorMsg.Println("[-] Unable to retrieve defaultNamingContext")
	}

	domainDN := sr.Entries[0].GetAttributeValue("defaultNamingContext")

	statusMsg := color.New(color.FgYellow)
	statusMsg.Printf("[!] Domain DN: %s\n", domainDN)

	adduser(domainDN, user, targetuser, group, module, conn)
}