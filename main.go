package main

import "flag"

func main() {
	var confPath = flag.String("conf", "config.tml", "config file path")
	flag.Parse()

	wasabi := New(*confPath)
	wasabi.Run()
}
