# guinea
Guinea is a command line interface library.

https://godoc.org/github.com/boreq/guinea

## Description
Programs very often organise the user interface in the form of subcommands. As
an example the `go` command lets the user invoke multiple subcommands such as
`go build` or `go get`. This library lets you nest any numbers of subcommands
(which can be thought of as separate programs) in each other easily building
complex user interfaces.


## Example
This program implements a root command which displays the program version and
two subcommands.

    package main

    import (
    	"fmt"
    	"github.com/boreq/guinea"
    	"os"
    )

    var rootCommand = guinea.Command{
    	Options: []guinea.Option{
    		guinea.Option{
    			Name:        "version",
    			Type:        guinea.Bool,
    			Description: "Display version",
    		},
    	},
    	Run: func(c guinea.Context) error {
    		if c.Options["version"].Bool() {
    			fmt.Println("v0.0.0-dev")
    			return nil
    		}
    		return guinea.ErrInvalidParms
    	},
    	Subcommands: map[string]*guinea.Command{
    		"display_text": &commandDisplayText,
    		"greet":        &commandGreet,
    	},
    	ShortDescription: "an example program using the guinea library",
    	Description:      `This program demonstrates the use of a CLI library.`,
    }

    var commandDisplayText = guinea.Command{
    	Run: func(c guinea.Context) error {
    		fmt.Println("Hello world!")
    		return nil
    	},
    	ShortDescription: "displays text on the screen",
    	Description:      `This is a subcommand that displays "Hello world!" on the screen.`,
    }

    var commandGreet = guinea.Command{
    	Arguments: []guinea.Argument{
    		guinea.Argument{
    			Name:        "person",
    			Multiple:    false,
    			Description: "a person to greet",
    		},
    	},
    	Options: []guinea.Option{
    		guinea.Option{
    			Name:        "times",
    			Type:        guinea.Int,
    			Description: "Number of greetings",
    			Default:     1,
    		},
    	},
    	Run: func(c guinea.Context) error {
    		for i := 0; i < c.Options["times"].Int(); i++ {
    			fmt.Printf("Hello %s!\n", c.Arguments[0])
    		}
    		return nil
    	},
    	ShortDescription: "greets the specified person",
    	Description:      `This is a subcommand that greets the specified person.`,
    }

    func main() {
    	if err := guinea.Run(&rootCommand); err != nil {
    		fmt.Fprintln(os.Stderr, err)
    	}
    }

And here are the example invocations of the program:

    $ ./main --help
    $ ./main --version
    $ ./main display_text
    $ ./main hello --help
    $ ./main hello boreq
    $ ./main hello --times 10 boreq
