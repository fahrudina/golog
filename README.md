# golog
Golang Log

```go
package main

import "github.com/fahrudina/golog"

func main() {
  level := "TRACE"  // 	ERROR ,	WARN,	INFO,	TRACE
  golog.SetLogLevel(level)
  
  golog.LogError("Error")
}
