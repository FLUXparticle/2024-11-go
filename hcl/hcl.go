package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	Server struct {
		Port int    `hcl:"port"`
		Host string `hcl:"host"`
	} `hcl:"server,block"`
	Database struct {
		User     string `hcl:"user"`
		Password string `hcl:"password"`
	} `hcl:"database,block"`
}

func main() {
	var config Config

	err := hclsimple.DecodeFile("config.hcl", nil, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server Port: %d\n", config.Server.Port)
	fmt.Printf("Database User: %s\n", config.Database.User)
}
