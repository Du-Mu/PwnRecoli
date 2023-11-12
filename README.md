# PwnRecoli

build the manager:
```shell
go build -o ./target/rpwn ./cmd/rpwn/rpwn.go  ./cmd/rpwn/commands.go
```

Then you get the manager `./target/rpwn`, you could set it to your path.

```shell
NAME:
   rpwn - A new cli application

USAGE:
   rpwn [global options] command [command options] [arguments...]

COMMANDS:
   init, i      init a new project
   generate, g  generate a new exploit
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

At the any dir, you could use `rpwn init` to init a new project.

There will be a `config.yaml` file in the dir, you could edit it to set the project info.

```yaml
host: 127.0.0.1
port: 9980
projectDir: /home/nemo/GolandProjects/PwnRecoli
allOutputStr: "aaa"
allOutputByte: !!byte [0xff]
```

All output str and bytes was the output of original awd elf target's output, 
which could be used to deal with the `p.recvuntil`

Then you could use `rpwn generate` to generate the server and client.

scp the client to the remote to replace the original pwn targets, and run the server in the local, 
you could get the output file with hex encode and hexdump



