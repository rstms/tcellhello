# tcellhello - a sparkly go tcell sparkly demo

Full-screen animated display with realtime interactive keyboard control.  Displays a character in random colors and locations on the console.

Key  | Function 
---- | ------------
 'q' | quit
 'c' | clear screen


To load libraries:

```
go get github.com/urfave/cli
go get github.com/gdamore/tcell
```

```
NAME:
   tcellhello - Hello World for tcell

USAGE:
   tcellhello [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug        
   --glyph value  Character to display (default: ".")
   --help, -h     show help
   --version, -v  print the version
```
