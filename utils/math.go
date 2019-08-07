package utils

import "math"

type Location struct {
  Lat float64
  Lng float64
}

// http://en.wikipedia.org/wiki/Haversine_formula
const RADIUS = 6371   // average radius of the Earth in km
func rad (x float64) float64 {return x * (math.Pi/180)}  // radians
func sqr (x float64) float64 {return math.Pow(x, 2)}
func Dist (p1 Location, p2 Location) float64 {  // distance between to points in km
  conv := func (x1 float64, x2 float64) float64 {return sqr(math.Sin( rad(x2-x1) / 2 ))}
  havLat, havLng := conv(p1.Lat, p2.Lat), conv(p1.Lng, p2.Lng)
  haversine := havLat + math.Cos(rad(p1.Lat)) * math.Cos(rad(p2.Lat)) * havLng
  return RADIUS * (2 * math.Atan2(math.Sqrt(haversine), math.Sqrt(1-haversine)))
}
