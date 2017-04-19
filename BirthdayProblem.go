package main

import "fmt"
import "flag"

func main() {
  var days, people, chanceOfNoRepeat float64
  people = 1
  days = 365

  var percentage float64
  flag.Float64Var(&percentage, "p", 0.5, "percent chance of a repeat")
  flag.Parse()
  chanceOfNoRepeat = 1.0

  for chanceOfNoRepeat > (1.0 - percentage) {
    people += 1
    var tmp float64
    tmp = float64((days-people+1) / days)
    //fmt.Println("tmp:", tmp)
    chanceOfNoRepeat = chanceOfNoRepeat * tmp
  }
  fmt.Println("With", people, "people there is a", (percentage*100.0),  "percent chance of 2 people sharing a birthday")
}
