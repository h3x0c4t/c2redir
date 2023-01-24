/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"c2redir/cmd"
	"c2redir/utils/misc"
	"log"
)

func main() {
	log.SetFlags(0)
	misc.CheckSupport()
	cmd.Execute()
}
