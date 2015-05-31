package main

import (
  "fmt"
  "os"
  "os/exec"
)

type Player struct {
  x, y int
}

type Level struct {
  width, height int
  data [][]int
}

func main() {
  //player := Player{x: 1, y: 1}
  level := Level{width: 23, height: 11}
  level.data = [][]int{
    // 0: floor
    // 1: wall
    {1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
    {1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
    {1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
    {1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
    {1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
    {1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
    {1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
    {1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
    {1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
    {1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1},
    {1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},
  }

  input()
  draw()
}

func draw(data [][]int) {
  clear()
  for i := 0; i < [][]data; i++ {

  }
  fmt.Println("-----------------------")
  fmt.Println("|                     |")
  fmt.Println("|                     |")
  fmt.Println("|                     |")
  fmt.Println("|                     |")
  fmt.Println("|                     |")
  fmt.Println("|                     |")
  fmt.Println("|                     |")
  fmt.Println("|                     |")
  fmt.Println("|                     |")
  fmt.Println("-----------------------")
}

func clear() {
  cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func input() {
  var response int
  fmt.Scanf("%c", &response)
  fmt.Println(response)
}

func update() {

}
