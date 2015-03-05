package main

import (
	//"fmt"
	"github.com/bedokuroff/dumper/dumper"
)

func main() {
	dbConfig := dumper.LoadConfig("config.json")
	dumper.RunDump(dbConfig)
	//fmt.Println(dbConfig)
}
