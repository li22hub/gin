package main

import (
	"package/Config"
	"package/Service"
)

var (
	cfgfile = "E:/package/database.yml"
)

func main() {
	//mf := map[int]func() int{
	//	1: func() int { return 10 },
	//	2: func() int { return 20 },
	//	5: func() int { return 50 },
	//}
	//fmt.Println(mf[1])
	//os.Exit(1)
	Config.Initialize(cfgfile)
	service.Start()
}