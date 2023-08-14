package main

#Dockerfile: [[{
	Cmd: string | "from" | "copy" | "run" | "volume" | "cmd" | "entrypoint" | "workdir" | "user" | "env" | "expose" | "label" | "arg" | "shell" | "arg"
	Flags: [...]
	JSON:   bool | false
	Stage:  int | 0
	SubCmd: string | ""
	Value: [string]
}, {...
}]]
