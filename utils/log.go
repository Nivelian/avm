package utils

import (
  "bufio"
  "fmt"
  "os"
  "time"
)

const FMT = "02-01-2006 15:04:05"  // WHY, GO ??!!
func _print (prefix string, xs ...interface{}) {
  fmt.Println(Prepend(Fmt("[{0}] [{1}]", time.Now().Format(FMT), prefix), xs)...)
}
func Log   (xs ...interface{}) {_print("INFO",  xs...)}
func Warn  (xs ...interface{}) {_print("WARN",  xs...)}
func Error (xs ...interface{}) {_print("ERROR", xs...)}

func PanicIf (x error, msg string) {
  if x != nil {
    Error(msg)
    panic(x)
  }
}

func FilePrint () (print func (...interface{}), finish func ()) {
  file, err := os.Create("./log.txt")
  PanicIf(err, "Error creating file")
  writer := bufio.NewWriter(file)
  print = func (xs ...interface{}) {fmt.Fprintln(writer, xs...)}
  finish = func () {
    writer.Flush()
    file.Close()
  }
  return
}
