package linearmath

import "math"

type Vec4 struct {
	X, Y, Z, W float32
}


func CreateVec4(x float32, y float32, z float32, w float32) *Vec4 {
	return &Vec4{x, y, z, w}
}

func (this *Vec4) Add(other *Vec4) *Vec4 {
	return CreateVec4(
		this.X + other.X,
		this.Y + other.Y,
		this.Z + other.Z,
		this.W + other.W)
}

func (this *Vec4) Sub(other *Vec4) *Vec4 {
	return CreateVec4(
		this.X - other.X,
		this.Y - other.Y,
		this.Z - other.Z,
		this.W - other.W)
}

func (this *Vec4) Mul(f float32) *Vec4 {
	return CreateVec4(
		this.X * f,
		this.Y * f,
		this.Z * f,
		this.W * f)
}

func (this *Vec4) Div(f float32) *Vec4 {
	return CreateVec4(
		this.X / f,
		this.Y / f,
		this.Z / f,
		this.W / f)
}

func (this *Vec4) Length() float32 {
	return float32(math.Sqrt(float64(this.X * this.X + this.Y * this.Y + this.Z * this.Z + this.W * this.W)))
}

func (this *Vec4) Normalize() *Vec4 {
	return this.Div(this.Length())
}

func (this *Vec4) Dot(other *Vec4) float32 {
	return this.X * other.X + this.Y * other.Y + this.Z * other.Z + this.W * other.Z
}

func (this *Vec4) Angle(other *Vec4) float32 {
	var d float32 = this.Dot(other)
	var al float32 = this.Length()
	var bl float32 = other.Length()
	return float32(math.Acos(float64(d / (al * bl))))
}

func (this *Vec4) ToArray() []float32 {
	var temp []float32
	temp = append(temp, this.X)
	temp = append(temp, this.Y)
	temp = append(temp, this.Z)
	temp = append(temp, this.W)
	return temp
}