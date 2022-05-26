# 🌳 Go YAML/JSON Configuration

![Go
Version](https://img.shields.io/github/go-mod/go-version/rwxrob/conf)
Coding simple !go function to grab current !k8s context[![GoDoc](https://godoc.org/github.com/rwxrob/conf?status.svg)](https://godoc.org/github.com/rwxrob/conf)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

This `conf` Bonzai branch is for safely managing any configuration as
single, local YAML/JSON using industry standards for local configuration
and system-safe writes. Use it to add a `conf` subcommand to any other
Bonzai command, or to your root Bonzai tree (`z`). All commands that use
`conf` that are composed into a single binary, no matter where in the
tree, will use the same local conf file even though the position
within the file will be qualified by the tree location.

By default, importing `conf` will assigned a new implementation of
`bonzai.Configurer` to `Z.Conf` (satisfying any `Z.Cmd` requirement for
configuration) and will use the name of the binary (`Z.ExeName`) as the
directory name within `os.UserConfDir` with a `config.yaml` file name.
To override this behavior, create a new `pkd/conf.Conf` struct assign
`Id`, `Dir` and `File`, and then assign that to `Z.Conf`.

## Install

This command can be installed as a standalone program (for combined use
with shell scripts perhaps) or composed into a Bonzai command tree.

Standalone

```
go install github.com/rwxrob/conf/cmd/conf@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
)

var Cmd = &bonzai.Cmd{
	Name:     `z`,
	Commands: []*Z.Cmd{help.Cmd, conf.Cmd},
}
```

Note conf is designed to be composed only in monolith mode (not
multicall binary).

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C conf conf
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.
