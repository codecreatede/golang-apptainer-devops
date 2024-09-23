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
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang-apptainer-devops% \
go run main.go createApp -h
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

- example of apptainer installation suid and non-suid
```
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang/golang-apptainer-devops% \
go run main.go apps -s yes
Updating and loading repositories:
Repositories loaded.
Package                  Arch   Version      Repository      Size
Installing:
 apptainer               x86_64 1.3.4-1.fc40 updates    119.3 MiB
Installing dependencies:
 fakeroot                x86_64 1.36-1.fc40  updates    163.0 KiB
 fakeroot-libs           x86_64 1.36-1.fc40  updates    124.8 KiB

Transaction Summary:
 Installing:        3 packages

Total size of inbound packages is 29 MiB. Need to download 29 MiB.
After this operation 120 MiB will be used (install 120 MiB, remove 0 B).
[1/3] fakeroot-libs-0:1.36-1.fc40.x86_6 100% | 115.1 KiB/s |  39.0 KiB |  00m00s
[2/3] fakeroot-0:1.36-1.fc40.x86_64     100% | 260.9 KiB/s |  99.9 KiB |  00m00s
[3/3] apptainer-0:1.3.4-1.fc40.x86_64   100% |   4.3 MiB/s |  28.7 MiB |  00m07s
--------------------------------------------------------------------------------
[3/3] Total                             100% |   4.1 MiB/s |  28.9 MiB |  00m07s
Running transaction
[1/5] Verify package files              100% |  11.0   B/s |   3.0   B |  00m00s
[2/5] Prepare transaction               100% |   3.0   B/s |   3.0   B |  00m01s
[3/5] Installing fakeroot-libs-0:1.36-1 100% |   2.4 MiB/s | 126.0 KiB |  00m00s
[4/5] Installing fakeroot-0:1.36-1.fc40 100% |   2.7 MiB/s | 167.5 KiB |  00m00s
>>> Running post-install scriptlet: fakeroot-0:1.36-1.fc40.x86_64
>>> Stop post-install scriptlet: fakeroot-0:1.36-1.fc40.x86_64
[5/5] Installing apptainer-0:1.3.4-1.fc 100% |  46.6 MiB/s | 119.4 MiB |  00m03s
>>> Running post-install scriptlet: apptainer-0:1.3.4-1.fc40.x86_64
>>> Stop post-install scriptlet: apptainer-0:1.3.4-1.fc40.x86_64
>>> Running post-transaction scriptlet: apptainer-0:1.3.4-1.fc40.x86_64
>>> Stop post-transaction scriptlet: apptainer-0:1.3.4-1.fc40.x86_64
>>> Running trigger-install scriptlet: glibc-common-0:2.39-22.fc40.x86_64
>>> Stop trigger-install scriptlet: glibc-common-0:2.39-22.fc40.x86_64
>>> Running trigger-install scriptlet: man-db-0:2.12.0-6.fc40.x86_64
>>> Stop trigger-install scriptlet: man-db-0:2.12.0-6.fc40.x86_64
```
- example of suid build install
```
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang/golang-apptainer-devops% \
go run main.go appN -n yes
Updating and loading repositories:
Repositories loaded.
Package         Arch   Version      Repository      Size
Installing:
 apptainer-suid x86_64 1.3.4-1.fc40 updates     17.8 MiB

Transaction Summary:
 Installing:        1 packages

Total size of inbound packages is 6 MiB. Need to download 6 MiB.
After this operation 18 MiB will be used (install 18 MiB, remove 0 B).
[1/1] apptainer-suid-0:1.3.4-1.fc40.x86 100% |   2.4 MiB/s |   5.7 MiB |  00m02s
--------------------------------------------------------------------------------
[1/1] Total                             100% |   2.0 MiB/s |   5.7 MiB |  00m03s
Running transaction
[1/3] Verify package files              100% |  19.0   B/s |   1.0   B |  00m00s
[2/3] Prepare transaction               100% |   1.0   B/s |   1.0   B |  00m01s
[3/3] Installing apptainer-suid-0:1.3.4 100% |  14.4 MiB/s |  17.8 MiB |  00m01s
>>> Running trigger-install scriptlet: glibc-common-0:2.39-22.fc40.x86_64
>>> Stop trigger-install scriptlet: glibc-common-0:2.39-22.fc40.x86_64
```


Gaurav Sablok
