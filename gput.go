package gput

import (
  "log"
  "strconv"
  "github.com/codeskyblue/go-sh"
)

func Setab(color int) () {
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

func Smul() {
  sh.Command("tput", "smul").Run()
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

func Sc() {
  sh.Command("tput", "sc").Run()
}

func Rc() {
  sh.Command("tput", "rc").Run()
}

func Lines()(string) {
  lines, err := sh.Command("tput", "lines").Output()
  if err != nil {
    log.Fatal(err)
  }
  return string(lines)
}

func Cols()(string) {
  cols, err := sh.Command("tput", "cols").Output()
  if err != nil {
    log.Fatal(err)
  }
  return string(cols)
}

func Cub(dist int) {
  sh.Command("tput", "cub", strconv.Itoa(dist)).Run()
}

func Cuf(dist int) {
  sh.Command("tput", "cuf", strconv.Itoa(dist)).Run()
}

func Cub1() {
  sh.Command("tput", "cub1").Run()
}

func Cuf1() {
  sh.Command("tput", "cuf1").Run()
}

func Ll() {
  sh.Command("tput", "ll").Run()
}

func Cuu1() {
  sh.Command("tput", "cuu1").Run()
}



func Ech(chars int) {
  sh.Command("tput", "ech", strconv.Itoa(chars)).Run()
}

func Clear() {
  sh.Command("tput", "clear").Run()
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
