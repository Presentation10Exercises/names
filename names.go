package names

import "fmt"

func init() {
  MyName := "something"
  fmt.Println(MyName)
}

func GetName() string {
  MyName := "something else"
  return MyName
}
