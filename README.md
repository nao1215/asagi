# [WIP] asagi - penetration toolbox
The asagi project is for developers (that's me) to learn penetration. I have not set a goal for the project. There are no legal issues with this project because the penetration tool does not yet exist.

## Dependency & Supported OS
- go version 1.18
- gcc version 11.0
- make version 4.3
- Linux

## How to build
```
$ git clone https://github.com/nao1215/asagi.git
$ cd asagi
$ make build
```

## My thoughts :)
Penetration......I don't know anything ;)  
I will develop asagi as a CLI tool consisting of multiple subcommands. The image is to copy the asagi command to a specific system (the system to be infiltrated). If this idea is not a good plan, change it.  

Currently, I have only developed the shared-library that gives asago execution permissions on the system to which the asagi command is sent.

```
$ ls
asagi asagi.so

$ enable -f ./asagi.so executable   ※ Only works with bash.

$ executable    ※ asagi command in the current directory is given execute permission.
```