/*
Package guinea is a command line interface library.

Defining commands

This library operates on a tree-like structure of available commands. In the
following example we define a root command with two subcommands. It will most
likely be the best to define the commands as global variables in your package.

	var rootCommand = guinea.Command{
		Run: func(c guinea.Context) error {
			fmt.Println("This is a root command.")
			return nil
		},
		Subcommands: map[string]*guinea.Command{
			"subcommandA": &subCommandA,
			"subcommandB": &subCommandB,
		},
	}

	var subCommandA = guinea.Command{
		Run: func(c guinea.Context) error {
			fmt.Println("This is the first subcommand.")
			return nil
		},
	}

	var subCommandB = guinea.Command{
		Run: func(c guinea.Context) error {
			fmt.Println("This is the second subcommand.")
			return nil
		},
	}

Executing the commands

After defining the commands use the run function to execute them. The library
will read os.Args to determine which command should be executed and to populate
the context passed to it with options and arguments.

	if err := guinea.Run(&rootCommand); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

Using the program

The user can invoke a program in multiple ways.

	$ ./example_program
	$ ./example_program subcommandA
	$ ./example_program subcommandB

Passing options and arguments

To let the user call a command with arguments or options populate the proper
lists in the command struct.

	var parametrizedCommand = guinea.Command{
		Run: func(c guinea.Context) error {
			fmt.Printf("Argument: %s\n", c.Arguments[0])
			fmt.Printf("Option: %d\n", c.Options["myopt"].Int())
			return nil
		},
		Arguments: []guinea.Argument{
			guinea.Argument{
				Name:        "myargument",
				Description: "An argument of a command.",
			},
		},
		Options: []guinea.Option{
			guinea.Option{
				Name:        "myopt",
				Type:        guinea.Int,
				Description: "An option which accepts an integer.",
				Default:     1,
			},
		},
	}

If you wish to parse the arguments in a different way simply don't define any
options or arguments in the command struct and pass the arguments from the
context to your parsing function.

*/
package guinea
