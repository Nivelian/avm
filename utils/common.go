package utils

import "github.com/mitchellh/copystructure"

func DeepCopy (x interface {}) interface{} {
  res, err := copystructure.Copy(x)
  PanicIf(err, "Deepcopy error")
  return res
}
