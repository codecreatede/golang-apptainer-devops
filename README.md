golang-apptainer-devops

- a apptainer installation and configuration management for linux systems.
- creates container and also ports them. 
- add apptainer directories and also creating containers from sanbox diectories. 

```
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang-apptainer-devops% \
go run main.go -h
This is for the installation of the apptainer images across the system devops

Usage:
  use [command]

Available Commands:
  appN
  apps
  completion  Generate the autocompletion script for the specified shell
  createApp
  help        Help about any command

Flags:
  -h, --help   help for use

Use "use [command] --help" for more information about a command.
```

- To create the SIF container follow this 
```
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang-apptainer-devops% go run main.go createApp -h
This will write and create the sif files for the Apptainer

Usage:
  use createApp [flags]

Flags:
  -a, --author string              author for the container/apptainer (default "name of the author")
  -c, --container-install string   create the container from the specific directories (default "create-container")
  -e, --envfile string             path to the environment file to be loaded (default "path to the env file")
  -h, --help                       help for createApp
  -f, --pathcommandfile string     path to the command file to be used for %run (default "path-to-the-command-file")
  -r, --runcommand string          path to the run commands (default "path to the file containing the run commands")
```



Gaurav Sablok
