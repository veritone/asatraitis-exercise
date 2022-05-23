package main

import (
	appdata "translate/appData"
	"translate/cli"
)

func main() {

	appdata.GetAppData()
	cli.Init()

}