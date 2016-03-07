# Gput

Golang port of the tput linux command

#### Install

```Shell
$ go get github.com/maxmclau/qgput
```

#### Usage

```Go
...
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

#### Cursor Movement

| **Function** | **Description** |
|:---|:---|
| `gput.Cup(int x, int y)` | Move cursor to screen location `x`,`y` (top left is 0,0) |
| `gput.Cub(int chars)` | Move `chars` characters left |
| `gput.Cuf(int chars)` | Move `chars` characters right |
| `gput.Cub1()` | Move left one space |
| `gput.Cuf1()` | Non-destructive space (move right one space) |
| `gput.Ll()` | Last line, first column (if no cup) |
| `gput.Cuu1()` | Up one line |
| `gput.Sc()` | Save the cursor position |
| `gput.Rc()` | Restore the cursor position |

#### Return Information

| **Function** | **Description** |
|:---|:---|
| `gput.Lines()(int lines)` | Output the number of lines of the terminal |
| `gput.Cols()(int cols)` | Output the number of columns of the terminal |
| `gput.Colors()(int colors)` | Output the number of colors supported in the terminal |

#### License

[**`MIT`**](LICENSE)
