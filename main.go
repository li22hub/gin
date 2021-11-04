package main

import (
	"package/Config"
	"package/Service"
)

var (
	cfgfile = "E:/package/database.yml"
)

func main() {
	Config.Initialize(cfgfile)
	service.Start()
}