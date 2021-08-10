package main

import (
	"context"
	"fmt"
	util "github.com/rvanderp/vsphere-perm-check/pkg/util"
	"log"
)

func main() {
	fmt.Printf("OpenShift vSphere Pre-Flight Permissions Validator\n\n")
	installConfig, err := util.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	ssn, err := util.GetSession(context.TODO(), installConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("checking permissions for user %s\n\n", installConfig.Username)
	err = util.ValidatePrivileges(ssn, installConfig)

	if err != nil {
		log.Printf("error while validating required privileges:\n\n%s", err.Error())
	} else {
		log.Printf("no missing privileges found for user")
	}
}
