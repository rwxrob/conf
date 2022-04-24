# Design Considerations

* **JSON Output**

  JSON is YAML. But JSON is also much safer to deal with when parsing
  and piping into other things. The `Query` form has been modeled after
  `jq` (which has become something of a standard tool for mining
  information from configuration and other files.

* **No SoftInit at init()**

  It's one thing to set the defaults for Z on import (like a database
  driver would). It's another thing to make potential changes to the
  persistence system before the author has had a chance to change those
  defaults. This means that calling SoftInit will forever be something
  that Cmd authors will have to do themselves, presumably from the
  `init()` of their Cmds.
