package flags

type Command struct {
	*Group

	Name   string
	Active *Command

	commands []*Command

	hasBuiltinHelpGroup bool
}

// The command interface should be implemented by any command added in the
// options. When implemented, the Execute method will be called for the last
// specified (sub)command providing the remaining command line arguments.
type Commander interface {
	// Execute will be called for the last active (sub)command. The
	// args argument contains the remaining command line arguments. The
	// error that Execute returns will be eventually passed out of the
	// Parse method of the Parser.
	Execute(args []string) error
}

// The usage interface can be implemented to show a custom usage string in
// the help message shown for a command.
type Usage interface {
	Usage() string
}

// AddCommand adds a new command to the parser with the given name and data. The
// data needs to be a pointer to a struct from which the fields indicate which
// options are in the command. The provided data can implement the Command and
// Usage interfaces.
func (c *Command) AddCommand(command string, shortDescription string, longDescription string, data interface{}) (*Command, error) {
	cmd := newCommand(command, shortDescription, longDescription, data)

	if err := cmd.scan(); err != nil {
		return nil, err
	}

	c.commands = append(c.commands, cmd)
	return cmd, nil
}

// Get a list of subcommands of this command.
func (c *Command) Commands() []*Command {
	return c.commands
}