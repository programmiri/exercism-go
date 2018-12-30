// Package space adds a method for calculation ones age in space
package space

//A Planet represents a planet name.
type Planet string

var planetMap = map[Planet]float64{
  "Earth":   1,
  "Mercury": 0.2408467,
  "Venus":   0.61519726,
  "Mars":    1.8808158,
  "Jupiter": 11.862615,
  "Saturn":  29.447498,
  "Uranus":  84.016846,
  "Neptune": 164.79132,
}

// Age calculates how old someone is on different planets
// based on seconds and a planet given
func Age(sec float64, planet Planet) float64 {
  var base float64 = 31557600
  return sec / (base * planetMap[planet])
}
