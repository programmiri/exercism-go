// Package triangle implements methods for triangles
package triangle

import "math"

// Kind represents the kind of triangle
type Kind int

const (
  // Equ is equilateral
  Equ = iota
  // Iso is isosceles
  Iso
  // Sca is scalene
  Sca
  // NaT is not a triangle
  NaT
)

func isNaT(a, b, c float64) bool {
  sides := []float64{a, b, c}
  is := false

  for _, side := range sides {
    if math.IsNaN(side) || math.IsInf(side, 0) || side <= 0 {
      is = true
    }
  }

  if is || (a+b < c || b+c < a || a+c < b) {
    return true
  }
  return false
}

// KindFromSides setermines if a triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
  if isNaT(a, b, c) {
    return NaT
  }

  if a == b && b == c {
    return Equ
  }

  if a == b || a == c || b == c {
    return Iso
  }
  return Sca
}
