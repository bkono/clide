// Package flagg provides a simple means of constructing CLI command
// hierarchies. A hierarchy is simply a tree of flag.FlagSets; the Parse
// function is then used to determine which command is being invoked.
package clide

import (
	"os"
)

// Root is the default root flag.FlagSet.
var Root = CommandLine

// New returns a new flag.FlagSet. It is equivalent to flag.NewFlagSet(name,
// flag.ExitOnError), and setting f.Usage = SimpleUsage(f, usage)
func New(name, usage string) *FlagSet {
	f := NewFlagSet(name, ExitOnError)
	f.Usage = SimpleUsage(f, usage)
	return f
}

// SimpleUsage returns a func that writes usage to cmd.Output(). If cmd has
// associated flags, the func also calls cmd.PrintDefaults.
func SimpleUsage(cmd *FlagSet, usage string) func() {
	return func() {
		cmd.Output().Write([]byte(usage))
		numFlags := 0
		cmd.VisitAll(func(*Flag) { numFlags++ })
		if numFlags > 0 {
			cmd.Output().Write([]byte("\nFlags:\n"))
			cmd.PrintDefaults()
		}
	}
}

// A Tree is a tree of commands and subcommands.
type Tree struct {
	Cmd *FlagSet
	Sub []Tree
}

// Parse parses os.Args according to the supplied Tree. It returns the
// most deeply-nested flag.FlagSet selected by the args.
func ParseTree(tree Tree) *FlagSet {
	return parseTree(tree, os.Args[1:])
}

func parseTree(tree Tree, args []string) *FlagSet {
	tree.Cmd.Parse(args)
	args = tree.Cmd.Args()
	if len(args) > 0 {
		for _, t := range tree.Sub {
			if t.Cmd.Name() == args[0] {
				return parseTree(t, args[1:])
			}
		}
	}
	return tree.Cmd
}
