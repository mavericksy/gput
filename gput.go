package gput

import (
  "strconv"
  "log"
  "strings"
  "github.com/codeskyblue/go-sh"
)

func Setab(color int) {
  sh.Command("tput", "setab", strconv.Itoa(color)).Run()
}

func Setb(color int) {
  sh.Command("tput", "setb", strconv.Itoa(color)).Run()
}

func Setaf(color int) {
  sh.Command("tput", "setaf", strconv.Itoa(color)).Run()
}

func Setf(color int) {
  sh.Command("tput", "setf", strconv.Itoa(color)).Run()
}



func Bold() {
  sh.Command("tput", "bold").Run()
}

func Dim() {
  sh.Command("tput", "dim").Run()
}

func Smul() {
  sh.Command("tput", "smul").Run()
}

func Rmul() {
  sh.Command("tput", "rmul").Run()
}

func Rev() {
  sh.Command("tput", "rev").Run()
}

func Blink() {
  sh.Command("tput", "blink").Run()
}

func Invis() {
  sh.Command("tput", "invis").Run()
}

func Smso() {
  sh.Command("tput", "smso").Run()
}

func Rmso() {
  sh.Command("tput", "rmso").Run()
}

func Sgr0() {
  sh.Command("tput", "sgr0").Run()
}



func Cup(x int, y int) {
  sh.Command("tput", "cup", strconv.Itoa(y), strconv.Itoa(x)).Run()
}

func Cub(dist int) {
  sh.Command("tput", "cub", strconv.Itoa(dist)).Run()
}

func Cuf(dist int) {
  sh.Command("tput", "cuf", strconv.Itoa(dist)).Run()
}

func Cuu(dist int) {
  sh.Command("tput", "cuu", strconv.Itoa(dist)).Run()
}

func Cud(dist int) {
  sh.Command("tput", "cud", strconv.Itoa(dist)).Run()
}

func Cub1(dist int) {
  sh.Command("tput", "cub1").Run()
}

func Cuf1() {
  sh.Command("tput", "cuf1").Run()
}

func Cuu1() {
  sh.Command("tput", "cuu1").Run()
}

func Cud1() {
  sh.Command("tput", "cud1").Run()
}

func Ll() {
  sh.Command("tput", "ll").Run()
}

func Sc() {
  sh.Command("tput", "sc").Run()
}

func Rc() {
  sh.Command("tput", "rc").Run()
}

func Civis() {
  sh.Command("tput", "civis").Run()
}

func Cnorm() {
  sh.Command("tput", "cnorm").Run()
}




func Lines()(int) {
  output, err := sh.Command("tput", "lines").Output()
  lines, err := strconv.Atoi(strings.TrimSpace(string(output)))
  if err != nil {
    log.Fatal(err)
  }
  return lines;
}

func Cols()(int) {
  output, err := sh.Command("tput", "cols").Output()
  cols, err := strconv.Atoi(strings.TrimSpace(string(output)))
  if err != nil {
    log.Fatal(err)
  }
  return cols;
}

func Colors()(int) {
  output, err := sh.Command("tput", "colors").Output()
  colors, err := strconv.Atoi(strings.TrimSpace(string(output)))
  if err != nil {
    log.Fatal(err)
  }
  return colors;
}



func Clear() {
  sh.Command("tput", "clear").Run()
}

func Ech(chars int) {
  sh.Command("tput", "ech", strconv.Itoa(chars)).Run()
}

func El1() {
  sh.Command("tput", "el1").Run()
}

func El() {
  sh.Command("tput", "el").Run()
}

func Ed() {
  sh.Command("tput", "ed").Run()
}

func Ich(chars int) {
  sh.Command("tput", "ich", strconv.Itoa(chars)).Run()
}

func Il(lines int) {
  sh.Command("tput", "il", strconv.Itoa(lines)).Run()
}

func Dl(lines int) {
  sh.Command("tput", "dl", strconv.Itoa(lines)).Run()
}
