package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

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
	ubuntu            string
	createcontainer   string
	pathfile          string
	envfile           string
	author            string
	runcom            string
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

var ubuntuCmd = &cobra.Command{
	Use:  "UbuntuInstall",
	Long: "This will install the Apptainer for the Ubuntu system. You need to define the version that you want to install such as 1.3.4",
	Run:  ubuntuInst,
}

func init() {
	appCmd.Flags().
		StringVarP(&appinstallsuid, "appinstall-suid", "s", "app-suid", "install the apptainer suid")
	appNCmd.Flags().
		StringVarP(&appinstallnonsuid, "appinstall-nonsuid", "n", "app-non-suid", "install the apptainer non-suid")
	ubuntuCmd.Flags().
		StringVarP(&ubuntu, "ubuntu-install", "u", "ubuntu-apptainer", "install the apptainer for ubuntu")
	createCmd.Flags().
		StringVarP(&createcontainer, "container-install", "c", "create-container", "create the container from the specific directories")
	createCmd.Flags().
		StringVarP(&pathfile, "pathcommandfile", "f", "path-to-the-command-file", "path to the command file to be used for %run")
	createCmd.Flags().
		StringVarP(&envfile, "envfile", "e", "path to the env file", "path to the environment file to be loaded")
	createCmd.Flags().
		StringVarP(&runcom, "runcommand", "r", "path to the file containing the run commands", "path to the run commands")
	createCmd.Flags().
		StringVarP(&author, "author", "a", "name of the author", "author for the container/apptainer")
	rootCmd.AddCommand(appCmd)
	rootCmd.AddCommand(appNCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(ubuntuCmd)
}

// struct for the commands
type pathfileStruct struct {
	commandname string
	command     string
}

// environment struct
type envStruct struct {
	env string
}

type runCommand struct {
	command string
}

// functions for the appfunc
func appfunc(cmd *cobra.Command, args []string) {
	if appinstallsuid == "yes" {
		appCi, err := exec.Command("sudo", "dnf", "install", "-y", "apptainer").Output()
		fmt.Println(string(appCi))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// functions for the appNfunc suid

func appNfunc(cmd *cobra.Command, args []string) {
	if appinstallnonsuid == "yes" {
		appSi, err := exec.Command("sudo", "dnf", "install", "-y", "apptainer-suid").Output()
		fmt.Println(string(appSi))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// functions for the ubuntu install

func ubuntuInst(cmd *cobra.Command, args []string) {
	if ubuntu == "yes" {
		install, err := exec.Command("sudo", "add-apt-repository", "-y", "ppa:apptainer/ppa").
			Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(install)
		ubuntInst, err := exec.Command("sudo", "apt", "install", "-y", "apptainer").Output()
		fmt.Println(string(ubuntInst))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// functions for the create containers.
// it reads the struct and stores each command there for later use also.

func createfunc(cmd *cobra.Command, args []string) {
	startline := []string{"Bootsrap: Docker"}
	image := bufio.NewScanner(os.Stdin)
	commandE := []pathfileStruct{}
	envE := []envStruct{}
	runcommand := []runCommand{}

	fOpen, err := os.Open(pathfile)
	if err != nil {
		log.Fatal(err)
	}
	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		commandE = append(commandE, pathfileStruct{
			commandname: strings.Split(string(line), " ")[0],
			command:     strings.Split(string(line), " ")[1],
		})
	}

	enOpen, err := os.Open(envfile)
	if err != nil {
		log.Fatal(err)
	}
	enRead := bufio.NewScanner(enOpen)
	for enRead.Scan() {
		line := enRead.Text()
		envE = append(envE, envStruct{
			env: string(line),
		})
	}

	rOpen, err := os.Open(runcom)
	if err != nil {
		log.Fatal(err)
	}
	rRead := bufio.NewScanner(rOpen)
	for rRead.Scan() {
		line := rRead.Text()
		runcommand = append(runcommand, runCommand{
			command: string(line),
		})
	}
	fmt.Println(startline)
	fmt.Println("From:", image)
	fmt.Println("%post")
	for i := range commandE {
		fmt.Println(commandE[i])
	}
	fmt.Println("%environment")
	for i := range envE {
		fmt.Println(envE[i])
	}
	fmt.Println()
	fmt.Println("%runscript")
	for i := range runcommand {
		fmt.Println(runcommand[i])
	}
	fmt.Println(author)
}
