package main

import (
  "fmt"
)

func main() {
  hexCodes := []byte{0x6d, 0x65, 0x6f, 0x77} // "meow"

  // Repeat 2 times as {2} in regex
  flagContent := ""
  for i := 0; i < 2; i++ {
    for _, b := range hexCodes {
      flagContent += string(b)
    }
  }

  // Add flag format
  flag := fmt.Sprintf("cryptoCTF{%s}", flagContent)
  fmt.Println(flag)
}