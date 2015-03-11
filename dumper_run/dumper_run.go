package main

import (
	"github.com/bedokuroff/dumper/dumper"
)

func main() {
	dbConfig := dumper.LoadConfig("config.json")
	dumper.RunDump(dbConfig)
}
