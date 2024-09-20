package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-20


A golang devops application for developing apptainer images and system wide installation of the appainer images


*/

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	appinstallsuid    string
	appinstallnonsuid string
	createcontainer   string
)

var rootCmd = &cobra.Command{
	Use:  "use",
	Long: "This is for the installation of the apptainer images across the system devops",
}

var appCmd = &cobra.Command{
	Use:  "apps",
	Long: "This will install the apptainer with the suid on your high performance system",
	Run:  appfunc,
}

var appNCmd = &cobra.Command{
	Use:  "appN",
	Long: "This will install the non-suid version of the apptainer on your high performance system",
	Run:  appNfunc,
}

var createCmd = &cobra.Command{
	Use:  "createApp",
	Long: "This will write and create the sif files for the Apptainer",
	Run:  createfunc,
}

func init() {
	appCmd.Flags().
		StringVarP(&appinstallsuid, "appinstall-suid", "as", "app-suid", "install the apptainer suid")
	appNCmd.Flags().
		StringVarP(&appinstallnonsuid, "appinstall-nonsuid", "an", "app-non-suid", "install the apptainer non-suid")
	createCmd.Flags().
		StringVarP(&createcontainer, "container-install", "co", "create-container", "create the container from the specific directories")

	rootCmd.AddCommand(appCmd)
	rootCmd.AddCommand(appNCmd)
	rootCmd.AddCommand(createCmd)
}

func appfunc(cmd *cobra.Command, args []string) {
	if appinstallsuid == "yes" {
		appC := cmd.Execute("sudo", "dnf", "install", "-y", "epel-release")
		err := appC.Run()
		if err != nil {
			log.Fatal(err)
		}
		appCi := cmd.Execute("sudo", "dnf", "install", "-y", "apptainer")
		appErr := appCi.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func appNfunc(cmd *cobra.Command, args []string) {
	if appinstallsuid == "yes" {
		appS := cmd.Execute("sudo", "dnf", "install", "-y", "epel-release")
		err := appS.Run()
		if err != nil {
			log.Fatal(err)
		}
		appSi := cmd.Execute("sudo", "dnf", "install", "apptainer-suid")
		appErr := appSi.Run()
		if appErr != nil {
			log.Fatal(appErr)
		}
	}
	// finished coding the installation and now start to code the sif file creation.
}
