# golog
Golang Log [![Go Report Card](https://goreportcard.com/badge/github.com/fahrudina/golog)](https://goreportcard.com/report/github.com/fahrudina/golog)

```go
package main

import "github.com/fahrudina/golog"

func main() {
  level := "TRACE"  // 	ERROR ,	WARN,	INFO,	TRACE
  golog.SetLogLevel(level)
  
  golog.LogError("Error")
}
