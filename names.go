package names

import "fmt"

var MyName string = "Bartholomew"

func init() {
  fmt.Println(MyName)
}

func GetName() string {
  return MyName
}
