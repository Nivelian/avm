package utils

import "reflect"

func Transpose (matrix [][]interface{}) [][]interface{} {
  xl := len(matrix[0])
  yl := len(matrix)
  result := make([][]interface{}, xl)
  for i := 0; i < xl; i++ {
    result[i] = make([]interface{}, yl)
    for j := 0; j < yl; j++ {
        result[i][j] = matrix[j][i]
    }
  }
  return result
}

func Prepend (x interface{}, xs []interface{}) []interface{} {
  return append([]interface{}{x}, xs...)}

func SetFn (ss []string) func (string) bool {     // Set simulation
  o := map[string]struct{}{}
  for _, s := range ss {o[s] = struct{}{}}
  return func (s string) bool {
    _, ok := o[s]
    return ok
  }
}

func Distinct (l interface{}) (res []interface{}) {
  o := map[interface{}]struct{}{}
  xs := reflect.ValueOf(l)
  for i := 0; i < xs.Len(); i++ {
    x := xs.Index(i).Interface()
    _, ok := o[x]
    if !ok {
      o[x] = struct{}{}
      res = append(res, x)
    }
  }
  return
}
