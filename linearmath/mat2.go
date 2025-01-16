package linearmath



type Mat2 struct {
	M00, M01 float32
	M10, M11 float32
}

func CreateMat2(m00 float32, m01 float32, m10 float32, m11 float32) *Mat2 {
	return &Mat2{
		m00, m01,
		m10, m11}
}

func CreateMat2Identity() *Mat2 {
	return CreateMat2(
		1.0, 0.0,
		0.0, 1.0)
}

func CreateMat2Zero() *Mat2 {
	return CreateMat2(
		0.0, 0.0,
		0.0, 0.0)
}

func (this *Mat2) Add(other *Mat2) *Mat2 {
	return CreateMat2(
		this.M00 + other.M00,
		this.M01 + other.M01,
		this.M10 + other.M10,
		this.M11 + other.M11)
}

func (this *Mat2) Sub(other *Mat2) *Mat2 {
	return CreateMat2(
		this.M00 + other.M00,
		this.M01 + other.M01,
		this.M10 + other.M10,
		this.M11 + other.M11)
}

func (this *Mat2) MulScalar(f float32) *Mat2 {
	return CreateMat2(
		this.M00 * f,
		this.M01 * f,
		this.M10 * f,
		this.M11 * f)
}

func (this *Mat2) DivScalar(f float32) *Mat2 {
	return CreateMat2(
		this.M00 / f,
		this.M01 / f,
		this.M10 / f,
		this.M11 / f)
}

func (this *Mat2) Mul(other *Mat2) *Mat2 {
	return CreateMat2(
		this.M00 * other.M00 + this.M01 * other.M10,
		this.M00 * other.M01 + this.M01 * other.M11,
		this.M10 * other.M00 + this.M11 * other.M10,
		this.M10 * other.M01 + this.M11 * other.M11)
}

func (this *Mat2) Transpose() *Mat2 {
	return CreateMat2(
		this.M00, this.M10,
		this.M01, this.M11)
}

func (this *Mat2) Trace() float32 {
	return this.M00 + this.M11
}

func (this *Mat2) Det() float32 {
	return this.M00 * this.M11 - this.M01 * this.M10
}

func (this *Mat2) Inverse() *Mat2 {
	var temp *Mat2 = CreateMat2(
		this.M11, -this.M01,
		-this.M10, this.M00)
	return temp.DivScalar(this.Det())
}

func (this *Mat2) MulVec2(v *Vec2) *Vec2 {
	return CreateVec2(
		this.M00 * v.X + this.M01 * v.Y,
		this.M10 * v.X + this.M11 * v.Y)
}

func (this *Mat2) ToArray() []float32 {
	var temp []float32
	temp = append(temp, this.M00)
	temp = append(temp, this.M01)
	temp = append(temp, this.M10)
	temp = append(temp, this.M11)
	return temp
}