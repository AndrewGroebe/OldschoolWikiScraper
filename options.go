package main

import (
  "log"
  //"fmt"
  "github.com/BurntSushi/toml"
)

type Options struct {
  Parse_items bool  `toml:"parse_items"`
  Parse_npcs  bool  `toml:"parse_npcs"`
}

var options *Options

func GetOptions() *Options {
  if options == nil {
      if _, err := toml.DecodeFile("./options.toml", &options); err != nil {
        log.Fatal(err)
      }
  }
  return options
}

func main() {
  GetOptions()
}
