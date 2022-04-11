// Copyright 2022 foo Authors
// SPDX-License-Identifier: Apache-2.0

// Package foo provides the Bonzai command branch of the same name.
package foo

import (
	"log"
	"text/template"

	"github.com/rwxrob/bonzai/comp"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/config"
	foo "github.com/rwxrob/foo/pkg"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{

	Name:      `foo`,
	Summary:   `just a sample foo command`,
	Version:   `v2.3.1`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Site:      `rwxrob.tv`,
	Source:    `git@github.com:rwxrob/foo.git`,
	Issues:    `github.com/rwxrob/foo/issues`,

	Commands: []*Z.Cmd{help.Cmd, config.Cmd, Bar, own, pkgfoo},

	Dynamic: template.FuncMap{
		"uname": func(_ *Z.Cmd) string { return Z.Out("uname", "-a") },
		"ls":    func() string { return Z.Out("ls", "-l", "-h") },
	},

	Description: `
		The foo commands do foo stuff. You can start the description here
		and wrap it to look nice and it will just work. Descriptions are
		written in BonzaiMark™, a simplified combination of CommonMark and
		Go doc that allows the inclusion of Go template markup that uses the
		Cmd itself as a data source. There are four block types and four
		span types in BonzaiMark:

		Spans

		    Plain
		    *Italic*
		    **Bold**
		    ***BoldItalic***
		    <Under> (brackets remain)

		Note that on most terminals italic is rendered as underlining and
		depending on how old the terminal, other formatting might not appear
		as expected. If you know how to set LESS_TERMCAP_* variables they
		will be observed when output is to the terminal.

		Blocks

		1. Paragraph
		2. Verbatim (block begins with '    ', never first)
		3. Numbered (block begins with '* ')
		4. Bulleted (block begins with '1. ')

		Currently, a verbatim block must never be first because of the
		stripping of initial white space.

		Templates

		Anything from Cmd that fulfills the requirement to be included in
		a Go text/template may be used. This includes {{ "{{ .Name }}" }}
		and the rest. A number of builtin template functions have also been
		added (such as {{ "indent" }}) which can receive piped input. You
		can add your own functions (or overwrite existing ones) by adding
		your own Dynamic template.FuncMap (see text/template for more about
		Go templates). Note that verbatim blocks will need to indented to work:

		    {{ "{{ ls | indent 4 }}" }}

		Produces a nice verbatim block:

		{{ ls | indent 4 }}

		Note this is different for every user and their specific system. The
		ability to incorporate dynamic data into any help documentation is
		a game-changer not only for creating very consumable tools, but
		creating intelligent, interactive training and education materials
		as well.

		The help documentation can scan the state of the system and give
		specific pointers and instruction based on elements of the host
		system that are missing or misconfigured.  Such was *never* possible
		with simple "man" pages and still is not possible with Cobra,
		urfave/cli, or any other commander framework in use today. In fact,
		Bonzai branch commands can be considered portable, dynamic web
		servers (once the planned support for embedded fs assets is
		added).`,

	Other: []Z.Section{
		{`Custom Sections`, `
			Additional sections can be added to the Other field.

			A Z.Section is just a Title and Body and can be assigned using
			composite notation (without the key names) for cleaner, in-code
			documentation.

			The Title will be capitalized for terminal output if using the
			common help.Cmd, but should use a suitable case for appearing in
			a book for other output renderers later (HTML, PDF, etc.)`,
		},
	},

	// no Call since has Commands, if had Call would only call if
	// commands didn't match
}

// Commands can be grouped into the same file or separately, public or
// private. Public let's others compose specific subcommands (foo.Bar),
// private just keeps it composed and only available within this Bonzai
// command.

// Aliases are not commands but will be replaced by their target names.

var Bar = &Z.Cmd{
	Name:     `bar`,
	Aliases:  []string{"B", "notbar"}, // to make a point
	Commands: []*Z.Cmd{help.Cmd, file},

	// Call first-class functions can be highly detailed, refer to an
	// existing function someplace else, or can call high-level package
	// library functions. Developers are encouraged to consider well where
	// they maintain the core logic of their applications. Often, it will
	// not be here within the Z.Cmd definition. One use case for
	// decoupled first-class Call functions is when creating multiple
	// binaries for different target languages. In such cases this
	// Z.Cmd definition is essentially just a wrapper for
	// documentation and other language-specific embedded assets.

	Call: func(_ *Z.Cmd, _ ...string) error { // note conventional _
		log.Printf("would bar stuff")
		return nil
	},
}

// Different completion methods are be set including the expected
// standard ones from bash and other shells. Not that these completion
// methods only work if the shell supports completion (including
// the Bonzai Shell, which can be set as the default Cmd to provide rich
// shell interactivity where normally no shell is available, such as in
// FROM SCRATCH containers that use a Bonzai tree as the core binary).

var file = &Z.Cmd{
	Name:      `file`,
	Commands:  []*Z.Cmd{help.Cmd},
	Completer: comp.File,
	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		log.Printf("would show file information about %v", args[0])
		return nil
	},
}

// When combining a high-level package library with a Bonzai command it
// is customary to create a pkg directory to avoid cyclical package
// import dependencies.

var pkgfoo = &Z.Cmd{
	Name: `pkgfoo`,
	Call: func(_ *Z.Cmd, _ ...string) error {
		foo.Foo()
		return nil
	},
}
