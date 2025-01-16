package linearmath

import "math"

type Vec3 struct {
	X, Y, Z float32
}

func CreateVec3(x float32, y float32, z float32) *Vec3 {
	return &Vec3{x, y, z}
}

func (this *Vec3) Add(other *Vec3) *Vec3 {
	return CreateVec3(
		this.X + other.X,
		this.Y + other.Y,
		this.Z + other.Z)
}

func (this *Vec3) Sub(other *Vec3) *Vec3 {
	return CreateVec3(
		this.X - other.X,
		this.Y - other.Y,
		this.Z - other.Z)
}

func (this *Vec3) Mul(f float32) *Vec3 {
	return CreateVec3(
		this.X * f,
		this.Y * f,
		this.Z * f)
}

func (this *Vec3) Div(f float32) *Vec3 {
	return CreateVec3(
		this.X / f,
		this.Y / f,
		this.Z / f)
}

func (this *Vec3) Length() float32 {
	return float32(math.Sqrt(float64(this.X * this.X + this.Y * this.Y + this.Z * this.Z)))
}

func (this *Vec3) Normalize() *Vec3 {
	return this.Div(this.Length())
}

func (this *Vec3) Dot(other *Vec3) float32 {
	return this.X * other.X + this.Y * other.Y + this.Z * other.Z
}

func (this *Vec3) Angle(other *Vec3) float32 {
	var d float32 = this.Dot(other)
	var al float32 = this.Length()
	var bl float32 = other.Length()
	return float32(math.Acos(float64(d / (al * bl))))
}

func (this *Vec3) Cross(other *Vec3) *Vec3 {
	return CreateVec3(
		this.Y * other.Z - this.Z * other.Y,
		this.Z * other.X - this.X * other.Z,
		this.X * other.Y - this.Y * other.X)
}

func (this *Vec3) ToArray() []float32 {
	var temp []float32
	temp = append(temp, this.X)
	temp = append(temp, this.Y)
	temp = append(temp, this.Z)
	return temp
}