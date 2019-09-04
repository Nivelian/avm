package overpass

import . "github.com/Nivelian/avm/utils"

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
