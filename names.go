package names

import "fmt"

var MyName = "Bartholomew"

func init() {
  fmt.Println(MyName)
}

func GetName() string {
  return MyName
}
