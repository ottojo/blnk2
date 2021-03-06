package vector

import "math"

// Here, 1 unit usually represents 1m IRL.

type Vec3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func (v Vec3) Plus(v2 Vec3) Vec3 {
	return Vec3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vec3) Minus(v2 Vec3) Vec3 {
	return v.Plus(v2.Times(-1))
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) Times(m float64) Vec3 {
	return Vec3{v.X * m, v.Y * m, v.Z * m}
}

func (v Vec3) Normalize() Vec3 {
	return v.Times(1 / v.Length())
}

func (v Vec3) ScaleToLength(l float64) Vec3 {
	return v.Normalize().Times(l)
}

func (v Vec3) Phi() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v Vec3) Theta() float64 {
	return math.Acos(v.Z / v.Length())
}
