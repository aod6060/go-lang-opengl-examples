package linearmath

type Mat3 struct {
	M00, M01, M02 float32
	M10, M11, M12 float32
	M20, M21, M22 float32
}

func CreateMat3(
	m00 float32, m01 float32, m02 float32,
	m10 float32, m11 float32, m12 float32, 
	m20 float32, m21 float32, m22 float32) *Mat3 {
	return &Mat3{
		m00, m01, m02,
		m10, m11, m12,
		m20, m21, m22}
}

func CreateMat3Identity() *Mat3 {
	return CreateMat3(
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0)
}

func CreateMat3Zero() *Mat3 {
	return CreateMat3(
		0.0, 0.0, 0.0,
		0.0, 0.0, 0.0,
		0.0, 0.0, 0.0)
}

func (this *Mat3) Add(other *Mat3) *Mat3 {
	return CreateMat3(
		this.M00 + other.M00,
		this.M01 + other.M01,
		this.M02 + other.M02,
		this.M10 + other.M10,
		this.M11 + other.M11,
		this.M12 + other.M12,
		this.M20 + other.M20,
		this.M21 + other.M21,
		this.M22 + other.M22)
}

func (this *Mat3) Sub(other *Mat3) *Mat3 {
	return CreateMat3(
		this.M00 - other.M00,
		this.M01 - other.M01,
		this.M02 - other.M02,
		this.M10 - other.M10,
		this.M11 - other.M11,
		this.M12 - other.M12,
		this.M20 - other.M20,
		this.M21 - other.M21,
		this.M22 - other.M22)
}

func (this *Mat3) MulScalar(f float32) *Mat3 {
	return CreateMat3(
		this.M00 * f,
		this.M01 * f,
		this.M02 * f,
		this.M10 * f,
		this.M11 * f,
		this.M12 * f,
		this.M20 * f,
		this.M21 * f,
		this.M22 * f)
}

func (this *Mat3) DivScalar(f float32) *Mat3 {
	return CreateMat3(
		this.M00 / f,
		this.M01 / f,
		this.M02 / f,
		this.M10 / f,
		this.M11 / f,
		this.M12 / f,
		this.M20 / f,
		this.M21 / f,
		this.M22 / f)
}

func (this *Mat3) Mul(other *Mat3) *Mat3 {
	return CreateMat3(
		this.M00 * other.M00 + this.M01 * other.M10 + this.M02 * other.M20,
		this.M00 * other.M01 + this.M01 * other.M11 + this.M02 * other.M21,
		this.M00 * other.M02 + this.M01 * other.M12 + this.M02 * other.M22,
		this.M10 * other.M00 + this.M11 * other.M10 + this.M12 * other.M20,
		this.M10 * other.M01 + this.M11 * other.M11 + this.M12 * other.M21,
		this.M10 * other.M02 + this.M11 * other.M12 + this.M12 * other.M22,
		this.M20 * other.M00 + this.M21 * other.M10 + this.M22 * other.M20,
		this.M20 * other.M01 + this.M21 * other.M11 + this.M22 * other.M21,
		this.M20 * other.M02 + this.M21 * other.M12 + this.M22 * other.M22)
}

func (this *Mat3) Transpose() *Mat3 {
	return CreateMat3(
		this.M00, this.M10, this.M20,
		this.M01, this.M11, this.M21,
		this.M02, this.M12, this.M22)
}

func (this *Mat3) Trace() float32 {
	return this.M00 + this.M11 + this.M22
}

func (this *Mat3) Det() float32 {
	var ma, mb, mc *Mat2

	ma = CreateMat2(
		this.M11, this.M12,
		this.M21, this.M22)

	mb = CreateMat2(
		this.M10, this.M12,
		this.M20, this.M22)

	mc = CreateMat2(
		this.M10, this.M11,
		this.M20, this.M21)

	return this.M00 * ma.Det() - this.M01 * mb.Det() + this.M02 * mc.Det()
}


func (this *Mat3) Inverse() *Mat3 {
	return CreateMat3(
		CreateMat2(this.M11, this.M12, this.M21, this.M22).Det(),
		-CreateMat2(this.M10, this.M12, this.M20, this.M22).Det(),
		CreateMat2(this.M10, this.M11, this.M20, this.M21).Det(),
		
		-CreateMat2(this.M01, this.M02, this.M21, this.M22).Det(),
		CreateMat2(this.M00, this.M02, this.M20, this.M22).Det(),
		-CreateMat2(this.M00, this.M01, this.M20, this.M21).Det(),
		
		CreateMat2(this.M01, this.M02, this.M11, this.M12).Det(),
		-CreateMat2(this.M00, this.M02, this.M10, this.M12).Det(),
		CreateMat2(this.M00, this.M01, this.M10, this.M11).Det(),
	).Transpose().DivScalar(this.Det())
}

func (this *Mat3) MulVec3(v *Vec3) *Vec3 {
	return CreateVec3(
		this.M00 * v.X + this.M01 * v.Y + this.M02 * v.Z,
		this.M10 * v.X + this.M11 * v.Y + this.M12 * v.Z,
		this.M20 * v.X + this.M21 * v.Y + this.M22 * v.Z)
}

func (this *Mat3) ToArray() []float32 {
	var temp []float32
	temp = append(temp, this.M00)
	temp = append(temp, this.M01)
	temp = append(temp, this.M02)

	temp = append(temp, this.M10)
	temp = append(temp, this.M11)
	temp = append(temp, this.M12)
	
	temp = append(temp, this.M20)
	temp = append(temp, this.M21)
	temp = append(temp, this.M22)
	return temp;
}