package tput

import (
  "github.com/codeskyblue/go-sh"
  "strconv"
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


func Bold(color int) {
  sh.Command("tput", "bold").Run()
}

func Smul(color int) {
  sh.Command("tput", "smul").Run()
}

func Rev(color int) {
  sh.Command("tput", "rev").Run()
}

func Blink(color int) {
  sh.Command("tput", "blink").Run()
}

func Invis(color int) {
  sh.Command("tput", "invis").Run()
}

func Smso(color int) {
  sh.Command("tput", "smso").Run()
}

func Rmso(color int) {
  sh.Command("tput", "rmso").Run()
}

func Sgr0(color int) {
  sh.Command("tput", "sgr0").Run()
}
