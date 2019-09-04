package overpass

import (
  . "github.com/Nivelian/avm/utils"
)

const URL = "https://lz4.overpass-api.de/api/interpreter"

type OsmTags struct {
  Name    string
  NameEn  string `json:"int_name"`
  NameBy  string `json:"name:be"`
  Colour  string
  Ref     string
}
type OsmMember struct {
  Type string
  Role string
  Ref  int
}
type OsmElement struct {
  Id       int
  Type     string
  Nodes    []int
  Lat      float64
  Lon      float64
  Tags     OsmTags
  Elements []OsmElement
  Members  []OsmMember
}

func LoadData (file string) []OsmElement {
  var res OsmElement
  args := map[string]interface{}{"data": string(LocalFile( Fmt("{0}.ovps", file) ))}
  FromJson(&res, RemoteFile(URL, nil, args))
  return res.Elements
}

func (el OsmElement) reverse () (res OsmElement) {
  res = DeepCopy(el).(OsmElement)
  res.Elements = []OsmElement{}
  for i := len(el.Elements)-1; i >=0; i-- {res.Elements = append(res.Elements, el.Elements[i])}
  return
}

func (el OsmElement) flatten () (res OsmElement) {
  res = DeepCopy(el).(OsmElement)
  var nodes []OsmElement
  for _, way := range el.Elements {nodes = append(nodes, way.Elements...)}
  res.Elements = nodes
  return
}
