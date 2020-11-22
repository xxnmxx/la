package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/xxnmxx/la"
)

func main() {
	args := os.Args
	// Dot
	dotCmd := flag.NewFlagSet("dot", flag.ExitOnError)
	dotDel := dotCmd.String("d", " ", "set delimiter")
	dotLazy := dotCmd.Bool("l", true, "set lazyquotes")
	// Transpose
	tCmd := flag.NewFlagSet("T", flag.ExitOnError)
	tDel := tCmd.String("d", " ", "set delimiter")
	tLazy := tCmd.Bool("l", true, "set lazyquotes")
	// Add
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addDel := addCmd.String("d", " ", "set delimiter")
	addLazy := addCmd.Bool("l", true, "set lazyquotes")
	// ScalarMultiple
	smlCmd := flag.NewFlagSet("sml", flag.ExitOnError)
	smlDel := smlCmd.String("d", " ", "set delimiter")
	smlLazy := smlCmd.Bool("l", true, "set lazyquotes")
	if len(args) < 3 {
		fmt.Println("subcmds: dot(file, file), T(file), smul(scalar file)")
		os.Exit(1)
	}
	switch args[1] {
	case "dot":
		dotCmd.Parse(args[2:])
		dotArgs := dotCmd.Args()
		if len(dotArgs) != 2 {
			fmt.Println("need 2 files")
			os.Exit(1)
		}
		del := []rune(*dotDel)
		a := la.ImportCsv(del[0], *dotLazy, dotArgs[0])
		b := la.ImportCsv(del[0], *dotLazy, dotArgs[1])
		dot := la.Dot(a, b)
		dot.Output(os.Stdout)
	case "add":
		addCmd.Parse(args[2:])
		addArgs := addCmd.Args()
		if len(addArgs) != 2 {
			fmt.Println("need 2 files")
			os.Exit(1)
		}
		del := []rune(*addDel)
		a := la.ImportCsv(del[0], *addLazy, addArgs[0])
		b := la.ImportCsv(del[0], *addLazy, addArgs[1])
		add := la.Add(a, b)
		add.Output(os.Stdout)
	case "sml":
		smlCmd.Parse(args[2:])
		smlArgs := smlCmd.Args()
		if len(smlArgs) != 2 {
			fmt.Println("need 2 args(float and file)")
			os.Exit(1)
		}
		del := []rune(*smlDel)
		s, err := strconv.ParseFloat(smlArgs[0], 64)
		if err != nil {
			fmt.Println("first arg must be a float")
			os.Exit(1)
		}
		m := la.ImportCsv(del[0], *smlLazy, smlArgs[1])
		sml := la.ScalarMultiple(s, m)
		sml.Output(os.Stdout)
	case "T":
		tCmd.Parse(args[2:])
		tArgs := tCmd.Args()
		if len(tArgs) != 1 {
			fmt.Println("Args must be one")
			os.Exit(1)
		}
		del := []rune(*tDel)
		a := la.ImportCsv(del[0], *tLazy, tArgs[0])
		t := la.Transpose(a)
		t.Output(os.Stdout)
	}
}
