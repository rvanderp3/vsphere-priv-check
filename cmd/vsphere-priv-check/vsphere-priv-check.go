package main

import (
	"context"
	"fmt"
	"github.com/rvanderp3/vsphere-priv-check/pkg/util"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var (
	rootOpts struct {
		checkFolder string
	}
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           filepath.Base(os.Args[0]),
		Short:         "Verifies vCenter user account privileges",
		Long:          "",
		Run:           runRootCmd,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.Flags().StringVar(&rootOpts.checkFolder, "check-folder", "", "verify privileges for folder")
	return cmd
}

func runRootCmd(cmd *cobra.Command, args []string) {
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
	log.Printf("folder: %s", rootOpts.checkFolder)
	log.Printf("checking permissions for user %s\n\n", installConfig.Username)
	err = util.ValidatePrivileges(ssn, installConfig, rootOpts.checkFolder)

	if err != nil {
		log.Printf("error while validating required privileges:\n\n%s", err.Error())
	} else {
		log.Printf("no missing privileges found for user")
	}
}

func main() {
	fmt.Printf("OpenShift vSphere Pre-Flight Permissions Validator v1.2\n\n")
	rootCmd := newRootCmd()
	rootCmd.Execute()
}
