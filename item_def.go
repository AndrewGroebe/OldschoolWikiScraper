package main

import (
  "encoding/json"
  "log"
)

type ItemDef struct {
    Name      string  `json:"name"`
    Members   bool    `json:"members"`
    QuestItem bool    `json:"questItem"`
    Tradeable bool    `json:"tradeable"`
    Equipable bool    `json:"equipable"`
    Stackable bool    `json:"stackable"`
    Noteable  bool    `json:"noteable"`
    Examine   string  `json:"examine"`
}

func (def *ItemDef) ToJson() string {
  result, err := json.MarshalIndent(def, "", "    ")
  if err != nil {
    log.Fatal(err)
  }
  return string(result)
}
