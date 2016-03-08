# Gput

Golang port of the tput unix command

#### Install

```Shell
$ go get github.com/maxmclau/gput
```

#### Usage

```Go
package main

import (
  "fmt"
  "github.com/maxmclau/gput"
)

func main() {
  gput.Cup(6, 6)
  gput.Setab(128)

  fmt.Printf("%v Gput Colors Available", gput.Colors());
}
```

#### Print Colors

```Go
// Prints all available ANSI colors to terminal

package main

import (
  "fmt"
  "github.com/maxmclau/gput"
)

func main() {
  gput.Clear()
  fmt.Printf("%v Available ANSI Colors: ", gput.Colors())
  gput.Cud1()
  
  for i := 0; i < gput.Colors(); i++ {
		gput.Setab(i)
    fmt.Printf(" %v ", i)
    gput.Sgr0()
	}
  gput.Cud1()
}
```

#### Set Colors

| **Function** | **Description** |
|:---|:---|
| `gput.Setab(int color)` | Set a background to `color` using ANSI escape |
| `gput.Setb(int color)` | Set a background `color` |
| `gput.Setaf(int color)` | Set a foreground `color` using ANSI escape |
| `gput.Setf(int color)` | Set a foreground `color` |

#### Toggle Modes

| **Function** | **Description** |
|:---|:---|
| `gput.Bold()` | Enable bold mode |
| `gput.Dim()` | Enable half-bright mode |
| `gput.Smul()` | Disable underline mode |
| `gput.Rmul()` | Enable underline mode |
| `gput.Rev()` | Disable reverse mode |
| `gput.Blink()` | Enable blink mode |
| `gput.Invis()` | Enable invisible mode |
| `gput.Smso()` | Enable standout mode (bold on rxvt) |
| `gput.Rmso()` | Disable standout mode |
| `gput.Sgr0()` | Turn off all attributes |

#### Cursor Manipulation

| **Function** | **Description** |
|:---|:---|
| `gput.Cup(int x, int y)` | Move cursor to screen location `x`,`y` (top left is 0,0) |
| `gput.Cub(int chars)` | Move `chars` characters left |
| `gput.Cuf(int chars)` | Move `chars` characters right |
| `gput.Cuu(int lines)` | Move `lines` lines up |
| `gput.Cud(int lines)` | Move `lines` lines down |
| `gput.Cub1()` | Move left one space |
| `gput.Cuf1()` | Move right one space (non-destructive space) |
| `gput.Cuu1()` | Move up one line |
| `gput.Cud1()` | Move down one line |
| `gput.Ll()` | Last line, first column (if no cup) |
| `gput.Sc()` | Save the cursor position |
| `gput.Rc()` | Restore the cursor position |
| `gput.Civis()` | Hide the cursor |
| `gput.Cnorm()` | Show the cursor |

#### Character Manipulation

| **Function** | **Description** |
|:---|:---|
| `gput.Clear()` | 	Clear the entire screen and home the cursor |
| `gput.Ich(chars int)` | 	Insert #1 characters |
| `gput.Ech(chars int)` | 	Erase #1 characters |

#### Return Information

| **Function** | **Description** |
|:---|:---|
| `gput.Lines()(int lines)` | Output the number of lines of the terminal |
| `gput.Cols()(int cols)` | Output the number of columns of the terminal |
| `gput.Colors()(int colors)` | Output the number of colors supported in the terminal |

#### ToDo

 - Implement the remaining commands from this [page](https://www.gnu.org/software/termutils/manual/termutils-2.0/html_chapter/tput_1.html).
 - Comment functions in code
 - Remove reliance on go-sh package
 - Generate documentation from GoDoc

#### License

[**`MIT`**](LICENSE)
