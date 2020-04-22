package main

import (
	"XE_Currency/api/modules/getXE_Currency"
	"XE_Currency/api/utils"
)

func main() {
	utils.InitViper()
	getXE_Currency.InitTable()
	getXE_Currency.InitJob()
}
