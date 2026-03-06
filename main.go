package main

import (
	"SysTrace_Server/services/web"
)

func main() {
	server := web.Server{}
	server.Start()
}
