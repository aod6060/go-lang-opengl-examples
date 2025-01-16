package linearmath

import "math"



type Vec2 struct {
	X, Y float32
}

/*
	For Conviense
*/
func CreateVec2(x float32, y float32) *Vec2 {
	return &Vec2{x, y}
}

func (this *Vec2) Add(other *Vec2) *Vec2 {
	var temp Vec2
	temp.X = this.X + other.X
	temp.Y = this.Y + other.Y
	return &temp
}

func (this *Vec2) Sub(other *Vec2) *Vec2 {
	var temp Vec2
	temp.X = this.X - other.X
	temp.Y = this.Y - other.Y
	return &temp
}

func (this *Vec2) Mul(f float32) *Vec2 {
	var temp Vec2
	temp.X = this.X * f
	temp.Y = this.Y * f
	return &temp
}

func (this *Vec2) Div(f float32) *Vec2 {
	var temp Vec2
	temp.X = this.X / f
	temp.Y = this.Y / f
	return &temp
}

func (this *Vec2) Length() float32 {
	return float32(math.Sqrt(float64(this.X * this.X + this.Y * this.Y)))
}

func (this *Vec2) Normalize() *Vec2 {
	return this.Div(this.Length())
}

func (this *Vec2) Dot(other *Vec2) float32 {
	return this.X * other.X + this.Y * other.Y
}

func (this *Vec2) Angle(other *Vec2) float32 {
	var d = this.Dot(other)
	var al = this.Length()
	var bl = other.Length()
	return float32(math.Acos(float64(d / (al * bl))))
}

func (this *Vec2) ToArray() []float32 {
	var temp []float32
	temp = append(temp, this.X)
	temp = append(temp, this.Y)
	return temp
}