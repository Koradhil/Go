package utils

import (
  "fmt"
  "time"
  "strconv"
)

const delay = 100

func TimeLogger() {
  fmt.Println("Starting...")
  time.Sleep(delay * time.Second)
  if delay == 1 {
    fmt.Println(strconv.Itoa(delay)+" second passed")
  } else {
    fmt.Println(strconv.Itoa(delay)+" seconds passed")
  }

  for i:=2; true; i++ {
    time.Sleep(delay * time.Second)
    fmt.Println(strconv.Itoa(i*delay)+" seconds passed")
  }
}
