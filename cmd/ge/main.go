package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

func main() {
	for i, a := range os.Args[1:] {
		fmt.Printf("%d:%s ", i, a)
	}

	if len(os.Args) >= 2 {
		switch arg1 := os.Args[1]; arg1 {
		case "web":
			//web()
		case "db":
			//dbstuff()
		case "tmpl":
			//TmplTest()
		default:
			edit()
		}
	} else {
		edit()
	}

}

func edit() {
	mycolor := tcell.ColorValues[tcell.ColorBlack]
	fmt.Printf("edit...%v\n", mycolor)
}

func usage() {
	fmt.Println()
	fmt.Printf("ge [file]\n")
	fmt.Println()
}
