package linearmath


import "math"

type Quaternion struct {
	W, X, Y, Z float32
}

func CreateQuaternion(w float32, x float32, y float32, z float32) *Quaternion{
	return &Quaternion{w, x, y, z}
}

func CreateQuaternionIdentity() *Quaternion {
	return CreateQuaternion(1.0, 0.0, 0.0, 0.0)
}

func CreateQuaternionZero() *Quaternion {
	return CreateQuaternion(0.0, 0.0, 0.0, 0.0)
}

func CreateQuaternionRotateAxis(radian float32, axis *Vec3) *Quaternion {
	return CreateQuaternion(
		float32(math.Cos(float64(radian / 2.0))),
		float32(math.Sin(float64(radian / 2.0))) * axis.X,
		float32(math.Sin(float64(radian / 2.0))) * axis.Y,
		float32(math.Sin(float64(radian / 2.0))) * axis.Z)
}

func (this *Quaternion) Add(other *Quaternion) *Quaternion {
	return CreateQuaternion(
		this.W + other.W,
		this.X + other.X,
		this.Y + other.Y,
		this.Z + other.Z)
}

func (this *Quaternion) Sub(other *Quaternion) *Quaternion {
	return CreateQuaternion(
		this.W - other.W,
		this.X - other.X,
		this.Y - other.Y,
		this.Z - other.Z)
}

func (this *Quaternion) MulScalar(f float32) *Quaternion {
	return CreateQuaternion(
		this.W * f,
		this.X * f,
		this.Y * f,
		this.Z * f)
}

func (this *Quaternion) DivScalar(f float32) *Quaternion {
	return CreateQuaternion(
		this.W / f,
		this.X / f,
		this.Y / f,
		this.Z / f)
}

func (this *Quaternion) Mul(other *Quaternion) *Quaternion {
	/*
	var ai, bi *Vec3
	var ar, br float32

	ar = this.W
	br = other.W

	ai = CreateVec3(this.X, this.Y, this.Z)
	bi = CreateVec3(other.X, other.Y, other.Z)

	var s float32 = ar * br - ai.Dot(bi)

	var part1, part2, part3 *Vec3

	part1 = bi.Mul(ar)
	part2 = ai.Mul(br)
	part3 = ai.Cross(bi)

	var final *Vec3 = part3.Add(part2.Add(part1))

	return CreateQuaternion(
		s,
		final.X,
		final.Y,
		final.Z)
	*/

	return CreateQuaternion(
		this.W * other.W - this.X * other.X - this.Y * other.Y - this.Z * other.Z,
		this.W * other.X + this.X * other.W + this.Y * other.X - this.Z * other.Y,
		this.W * other.Y - this.X * other.Z + this.Y * other.W + this.Z * other.X,
		this.W * other.Z + this.X * other.Y - this.Y * other.X + this.Z * other.W)
}

func (this *Quaternion) Length() float32 {
	return float32(math.Sqrt(float64(this.W * this.W + this.X * this.X + this.Y * this.Y + this.Z * this.Z)))
}

func (this *Quaternion) Normalize() *Quaternion {
	return this.DivScalar(this.Length())
}

func (this *Quaternion) Conjugate() *Quaternion {
	return CreateQuaternion(
		this.W,
		this.X * -1,
		this.Y * -1,
		this.Z * -1,
	)
}

func (this *Quaternion) Inverse() *Quaternion {
	/*
	float av = this.Length()
	av *= av
	av = 1.0 / av
	cong = this.Conjugate()
	return cong.MulScalar()
	*/
	return this.Conjugate().DivScalar(this.Length())
}

func (this *Quaternion) ToMat4() *Mat4 {
	var temp *Mat4 = CreateMat4Identity()

	// q0=w, q1=x, q2=y, q3=z
	temp.M00 = 2 * (this.W * this.W + this.X * this.X) - 1
	temp.M01 = 2 * (this.X * this.Y - this.W * this.Z)
	temp.M02 = 2 * (this.X * this.Z + this.W * this.Y)

	temp.M10 = 2 * (this.X * this.Y + this.W * this.Z)
	temp.M11 = 2 * (this.W * this.W + this.Y * this.Y) - 1
	temp.M12 = 2 * (this.Y * this.Z - this.W * this.X)

	temp.M20 = 2 * (this.X * this.Z - this.W * this.Y)
	temp.M21 = 2 * (this.Y * this.Z + this.W * this.X)
	temp.M22 = 2 * (this.W * this.W + this.Z * this.Z) - 1

	return temp
}

func (this *Quaternion) ToArray() []float32 {
	var temp []float32
	temp = append(temp, this.W)
	temp = append(temp, this.X)
	temp = append(temp, this.Y)
	temp = append(temp, this.Z)
	return temp
}