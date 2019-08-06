package utils

import (
  "fmt"
  "strings"
  "strconv"
  "net/url"
)

func Str (x interface {}) string {return fmt.Sprintf("%v", x)}

func replace (format string, ss []string) string {
  return strings.NewReplacer(ss...).Replace(format)}
func Fmt (format string, xs ...interface{}) string {
  var args []string
  for i, x := range xs {
    args = append(args, fmt.Sprintf("{%v}", i), Str(x))
  }
  return replace(format, args)
}
func FmtNamed (format string, o interface{}) string {
  var args []string
  for k, x := range StructToMap(o) {
    args = append(args, fmt.Sprintf("{%v}", k), Str(x))
  }
  return replace(format, args)
}

func StrToInt (s string) int {
  if x, err := strconv.Atoi(s); err == nil {
    return x
  } else {
    Error("String to Int convertion failed")
    return 0
  }
}
func BinaryToInt (s string) int {
  if x, err := strconv.ParseInt(s, 2, 64); err == nil {
    return int(x)
  } else {
    Error("Binary to Int convertion failed")
    return 0
  }
}

func Join (sep string, xs ...interface{}) string {
  var ss []string
  for _, x := range xs {
    ss = append(ss, Str(x))
  }
  return strings.Join(ss, sep)
}

func MapToUrl (o map[string]interface{}) string {
  params := url.Values{}
  for k, v := range o {params.Add(k, Str(v))}
  return params.Encode()
}
