package conf

import (
	"fmt"
	"os"

	Z "github.com/rwxrob/bonzai/z"
	_conf "github.com/rwxrob/conf/pkg"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
)

var conf _conf.C

func init() {
	dir, _ := os.UserConfigDir()
	conf = _conf.C{
		Id:   Z.ExeName,
		Dir:  dir,
		File: `config.yaml`,
	}
	Z.Conf = conf
}

var Cmd = &Z.Cmd{

	Name:      `conf`,
	Summary:   `manage local YAML/JSON configuation`,
	Version:   `v0.5.3`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Commands:  []*Z.Cmd{help.Cmd, data, _init, edit, _file, query},
	Description: `
		The **{{.Name}}** Bonzai branch is for safely managing any
		configuration as single, local YAML/JSON using industry standards
		for local configuration and persistence to a file using system-wide
		semaphores for safety. Use it to add a **{{.Name}}** subcommand to
		any other Bonzai command, or to your root Bonzai tree. Either way,
		the same single configuration file is used, only the path within the
		configuration data is affected by the position of the **{{.Name}}**
		command.

		Querying configuration data can be easily accomplished with the
		**query** command that uses the same selection syntax as the **yq**
		Go utility (the same **yqlib** is used).

		All changes to configuration values are done via the **edit** command
		since configurations may be complex and deeply nested in some cases
		and promoting the automatic changing of configuration values opens
		the possibility that one buggy composed command might overwrite one
		or all the configurations for everything everything else composed
		into the binary.

		CAUTION: Take particular note that all commands composed into
		a single binary, no matter where in the tree, will use the same
		local config file even though the position within the file will be
		qualified by the tree location. Therefore, any composite command can
		read the configurations of any other composite command within the
		same binary. This is by design, but all commands composed together
		should always be vetted for safe practices.

		[The **vars** Bonzai branch is recommended when wanting to persist
		performant local data between command executions.]`,
}

var _init = &Z.Cmd{
	Name:     `init`,
	Aliases:  []string{"i"},
	Summary:  `(re)initializes current configuration`,
	Commands: []*Z.Cmd{help.Cmd},
	ReqConf:  true, // but fulfills at init() above
	Call: func(x *Z.Cmd, _ ...string) error {
		if term.IsInteractive() {
			r := term.Prompt(`Really initialize %v? (y/N) `, conf.DirPath())
			if r != "y" {
				return nil
			}
		}
		return Z.Conf.Init()
	},
}

var _file = &Z.Cmd{
	Name:     `file`,
	Aliases:  []string{"f"},
	Summary:  `outputs path to file ({{execonfdir "config.yaml" }})`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, _ ...string) error {
		fmt.Println(conf.Path())
		return nil
	},
}

var data = &Z.Cmd{
	Name:    `data`,
	Aliases: []string{"d"},
	Summary: `outputs conf data ({{execonfdir "config.yaml" }})`,

	Description: `
			The **{{.Name}}** command prints the entire, unobfuscated contents
			of the YAML configuration file without warning.

			WARNING: Since configuration data regularly includes secrets
			(tokens, keys, etc.) be aware that anyone able to view your screen
			could compromise your security when using this command in front of
			them (presentations, streaming, etc.).`,

	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, _ ...string) error {
		fmt.Print(conf.Data())
		return nil
	},
}

var edit = &Z.Cmd{
	Name:     `edit`,
	Summary:  `edit conf file ({{execonfdir "config.yaml"}}) `,
	Aliases:  []string{"e"},
	Commands: []*Z.Cmd{help.Cmd},

	Description: `
		The **{{.Name}}** command will the configuration file for editing in
		the currently configured editor (in order or priority):

		* $VISUAL
		* $EDITOR
		* vi
		* vim
		* nano

		The edit command hands over control of the currently running process
		to the editor. `,

	Call: func(x *Z.Cmd, _ ...string) error { return conf.Edit() },
}

var query = &Z.Cmd{
	Name:     `query`,
	Summary:  `query conf data using jq/yq style`,
	Usage:    `<dotted>`,
	Aliases:  []string{"q", "get"},
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) == 0 {
			return x.UsageError()
		}
		conf.QueryPrint(args[0])
		return nil
	},
}
