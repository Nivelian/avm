package utils

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "encoding/xml"
  "fmt"
)

func RemoteFile (url string, headers map[string]string) []byte {
  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  PanicIf(err, "Error creating new request")
  for k, v := range headers {req.Header.Add(k, v)}
  Log(req.Header)

  x, err := client.Do(req)
  PanicIf(err, "HTTP Get error")
  defer x.Body.Close()

  if x.StatusCode != http.StatusOK {
    panic( fmt.Errorf("Response Status error: %v", x.StatusCode) )}

  data, err := ioutil.ReadAll(x.Body)
  PanicIf(err, "Failed to read body")

  return data
}

func LocalFile (path string) []byte {
  res, err := ioutil.ReadFile(path)
  PanicIf(err, "Failed to read file")
  return []byte(res)
}

func FromJson (res interface{}, file []byte) {
  PanicIf(json.Unmarshal(file, res), "Failed to parse json")}

func FromXml (res interface{}, file []byte) {
  PanicIf(xml.Unmarshal(file, res), "Failed to parse xml")}
