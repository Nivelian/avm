package utils

import ("reflect")

func StructToMap (x interface{}) map[string]interface{} {
  res := make(map[string]interface{})
  value := reflect.ValueOf(x)
  for i := 0; i < value.NumField(); i++ {
    res[value.Type().Field(i).Name] = value.Field(i).Interface()
  }
  return res
}
