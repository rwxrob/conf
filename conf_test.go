package conf_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/conf"
	"github.com/rwxrob/fs/dir"
)

func ExampleC_OverWrite() {

	c := conf.C{Id: `foo`, Dir: `testdata`, File: `config.yaml`}

	thing := struct {
		Some  string
		Other string
	}{"some", "other"}

	if err := c.OverWrite(thing); err != nil {
		fmt.Println(err)
	}

	dir.Create(`testdata/foo`)
	defer os.RemoveAll(`testdata/foo`)

	if err := c.OverWrite(thing); err != nil {
		fmt.Println(err)
	}
	c.Print()

	// Output:
	// some: some
	// other: other
}

func ExampleQuery() {

	c := conf.C{Id: `bar`, Dir: `testdata`, File: `config.yaml`}

	c.QueryPrint(".")
	c.QueryPrint(".some")
	c.QueryPrint(".here")

	// Output:
	// some: thing
	// here: goes
	// command:
	//   path: /here/we/go
	// thing
	// goes

}
