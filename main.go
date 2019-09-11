package main

import (
	"fmt"
	"os"
	"strings"

	"metagit.org/blizzlike/wowpasswd/srp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %v <username> <password> <retyped>\n", os.Args[0])
		os.Exit(1)
	}

	username := strings.ToUpper(os.Args[1])
	password := os.Args[2]

	retype := os.Args[3]

	if strings.Compare(password, retype) != 0 {
		fmt.Printf("\nSorry, passwords do not match!\n")
		os.Exit(2)
	}

	identifier := srp.Hash(username, password)

	auth := srp.New()
	auth.GenerateSalt()
	auth.ComputeVerifier(identifier)

	fmt.Printf("INSERT INTO account(username,v,s,joindate) VALUES('%v','%v','%v',NOW())", username, auth.GetSalt(), auth.GetVerifier())
}
