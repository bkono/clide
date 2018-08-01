package main

import (
	"fmt"

	"github.com/bkono/clide"
)

func main() {
	var (
		config string
		length float64
		age    int
		name   string
		female bool
	)
	rootCmd := clide.Root
	rootCmd.Usage = clide.SimpleUsage(clide.Root, "Hi from a simple usage")
	fooCmd := clide.New("foo", "The foo subcommand")
	bars := fooCmd.Int("bars", 0, "number of bars to foo")
	quuxCmd := clide.New("quux", "The quux subcommand")
	fooBarCmd := clide.New("bar", "the bar subcommand")

	clide.StringVar(&config, "config", "", "help message")
	clide.StringVar(&name, "name", "", "help message")
	clide.IntVar(&age, "age", 0, "help message")
	clide.Float64Var(&length, "length", 0, "help message")
	clide.BoolVar(&female, "female", false, "help message")

	// construct the command hierarchy
	tree := clide.Tree{
		Cmd: rootCmd,
		Sub: []clide.Tree{
			{
				Cmd: fooCmd,
				Sub: []clide.Tree{{Cmd: fooBarCmd}}},
			{Cmd: quuxCmd},
		},
	}

	cmd := clide.ParseTree(tree)
	args := cmd.Args()

	fmt.Println("length:", length)
	fmt.Println("age:", age)
	fmt.Println("name:", name)
	fmt.Println("female:", female)
	fmt.Println("bars: ", *bars)

	switch cmd {
	case rootCmd:
		fmt.Printf("it was root! %v\n", bars)
	case fooCmd:
		fmt.Printf("fooing %v bars\n", *bars)
	case quuxCmd:
		fmt.Println("quux!", args)
	case fooBarCmd:
		fmt.Println("we bar'd the foo!")
	}
}
