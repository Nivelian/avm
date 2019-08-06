package utils

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "encoding/xml"
  "fmt"
  "strings"
)

func RemoteFile (url string, headers map[string]string, body map[string]interface{}) []byte {
  client := &http.Client{}
  req, err := http.NewRequest("POST", url, strings.NewReader( MapToUrl(body) ))
  PanicIf(err, "Error creating new request")
  if body != nil {req.Header.Add("Content-Type", "application/x-www-form-urlencoded")}
  for k, v := range headers {req.Header.Add(k, v)}

  x, err := client.Do(req)
  PanicIf(err, "HTTP Request error")
  defer x.Body.Close()

  if x.StatusCode != http.StatusOK {
    panic( fmt.Errorf("Response Status error: %v", x.StatusCode) )}

  data, err := ioutil.ReadAll(x.Body)
  PanicIf(err, "Failed to read body")

  return data
}

func LocalFile (path string) []byte {
  res, err := ioutil.ReadFile( Fmt("data/{0}", path) )
  PanicIf(err, "Failed to read file")
  return []byte(res)
}

func FromJson (res interface{}, file []byte) {
  PanicIf(json.Unmarshal(file, res), "Failed to parse json")}

func FromXml (res interface{}, file []byte) {
  PanicIf(xml.Unmarshal(file, res), "Failed to parse xml")}
