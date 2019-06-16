/*
 * Copyright © 2019 Hedzr Yeh.
 */

package cmdr

// WalkAllCommands loops for all commands, starting from root.
func WalkAllCommands(walk func(cmd *Command, index int) (err error)) (err error) {
	command := &rootCommand.Command
	err = walkFromCommand(command, 0, walk)
	return
}

func walkFromCommand(cmd *Command, index int, walk func(cmd *Command, index int) (err error)) (err error) {
	if err = walk(cmd, index); err != nil {
		return
	}
	for ix, cc := range cmd.SubCommands {
		if err = walkFromCommand(cc, ix, walk); err != nil {
			return
		}
	}
	return
}
