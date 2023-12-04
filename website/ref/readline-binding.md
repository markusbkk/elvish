<!-- toc -->

@module readline-binding

# Introduction

The `readline-binding` module provides readline-like key bindings, such as
binding <kbd>Ctrl-A</kbd> to move the cursor to the start of the line. To use,
put the following in your [`rc.elv`](command.html#rc-file):

```elvish
use readline-binding
```

Note that this will override some of the standard bindings. For example,
<kbd>Ctrl-L</kbd> will be bound to a function that clears the terminal screen
rather than start location mode.

See the
[source code](https://github.com/markusbkk/elvish/pkg/mods/readlinebinding/readline-binding.elv)
for details.
