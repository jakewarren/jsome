package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/fatih/color"
	"github.com/nwidger/jsoncolor"
	"github.com/spf13/pflag"
)

func main() {

	displayHelp := pflag.BoolP("help", "h", false, "display help")

	pflag.Parse()

	// override the default usage display
	if *displayHelp {
		displayUsage()
		os.Exit(0)
	}

	//human-friendly CLI output
	log.SetHandler(cli.New(os.Stderr))

	//set the logging level
	log.SetLevel(log.WarnLevel)

	r, err := openStdinOrFile()

	checkError("error acquiring input", err)

	prettyPrint(r)

}

func prettyPrint(data io.Reader) {
	b, _ := ioutil.ReadAll(data)

	// create custom formatter,
	f := jsoncolor.NewFormatter()

	// set custom colors
	f.SpaceColor = color.New(color.FgRed, color.Bold)
	f.CommaColor = color.New(color.FgWhite, color.Bold)
	f.ColonColor = color.New(color.FgYellow)
	f.ObjectColor = color.New(color.FgBlue, color.Bold)
	f.ArrayColor = color.New(color.FgWhite)
	f.FieldColor = color.New(color.FgGreen)
	f.StringColor = color.New(color.FgMagenta, color.Bold)
	f.TrueColor = color.New(color.FgCyan, color.Bold)
	f.FalseColor = color.New(color.FgHiRed)
	f.NumberColor = color.New(color.FgHiYellow)
	f.NullColor = color.New(color.FgWhite, color.Bold)
	f.StringQuoteColor = color.New(color.FgBlue, color.Bold)

	// colorized output is written to dst
	dst := &bytes.Buffer{}
	err := f.Format(dst, b)
	checkError("error colorizing JSON", err)

	// print colorized output to stdout
	fmt.Println(dst.String())
}

// print custom usage instead of the default provided by pflag
func displayUsage() {

	fmt.Printf("Usage: jsome [<flags>] [FILE]\n\n")
	fmt.Printf("Example:\n\tjsome file.json\n")
	fmt.Printf("\tcat file.json | jsome \n\n")
	fmt.Printf("Optional flags:\n\n")
	pflag.PrintDefaults()
}

// openStdinOrFile reads from stdin or a file based on what input the user provides
func openStdinOrFile() (io.Reader, error) {

	// try to open a file if exists
	if len(pflag.Args()) == 1 {
		r, err := os.Open(pflag.Arg(0))
		if err != nil {
			return nil, err
		}
		return r, err
	}

	// no file provided so attempt to read piped data from stdin
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	// ensure stdin is a pipe, bail if not
	if info.Mode()&os.ModeNamedPipe == 0 {
		return nil, fmt.Errorf("no input specified")
	}

	return os.Stdin, err
}

func checkError(message string, err error) {
	if err != nil {
		log.WithError(err).Fatal(message)
	}
}
