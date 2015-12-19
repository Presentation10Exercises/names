package names

import "fmt"

var MyName string = "Bartholomew"

func init() {
  MyName := "something"
  fmt.Println(MyName)
}

func GetName() string {
  return MyName
}
