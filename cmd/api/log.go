package main

import "log"

func init() {
	log.SetPrefix("LOG : ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}
