package linearmath

//import "fmt"

type Mat4 struct {
	M00, M01, M02, M03 float32
	M10, M11, M12, M13 float32
	M20, M21, M22, M23 float32
	M30, M31, M32, M33 float32
}

func CreateMat4(
	m00 float32, m01 float32, m02 float32, m03 float32,
	m10 float32, m11 float32, m12 float32, m13 float32,
	m20 float32, m21 float32, m22 float32, m23 float32,
	m30 float32, m31 float32, m32 float32, m33 float32) *Mat4 {
	return &Mat4 {
		m00, m01, m02, m03,
		m10, m11, m12, m13,
		m20, m21, m22, m23,
		m30, m31, m32, m33}
}

func CreateMat4Identity() *Mat4 {
	return CreateMat4(
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0)
}

func CreateMat4Zero() *Mat4 {
	return CreateMat4(
		0.0, 0.0, 0.0, 0.0,
		0.0, 0.0, 0.0, 0.0,
		0.0, 0.0, 0.0, 0.0,
		0.0, 0.0, 0.0, 0.0)
}

func (this *Mat4) Add(other *Mat4) *Mat4 {
	return CreateMat4(
		this.M00 + other.M00,
		this.M01 + other.M01,
		this.M02 + other.M02,
		this.M03 + other.M03,
		this.M10 + other.M10,
		this.M11 + other.M11,
		this.M12 + other.M12,
		this.M13 + other.M13,
		this.M20 + other.M20,
		this.M21 + other.M21,
		this.M22 + other.M22,
		this.M23 + other.M23,
		this.M30 + other.M30,
		this.M31 + other.M31,
		this.M32 + other.M32,
		this.M33 + other.M33)
}

func (this *Mat4) Sub(other *Mat4) *Mat4 {
	return CreateMat4(
		this.M00 - other.M00,
		this.M01 - other.M01,
		this.M02 - other.M02,
		this.M03 - other.M03,
		this.M10 - other.M10,
		this.M11 - other.M11,
		this.M12 - other.M12,
		this.M13 - other.M13,
		this.M20 - other.M20,
		this.M21 - other.M21,
		this.M22 - other.M22,
		this.M23 - other.M23,
		this.M30 - other.M30,
		this.M31 - other.M31,
		this.M32 - other.M32,
		this.M33 - other.M33)
}

func (this *Mat4) MulScalar(f float32) *Mat4 {
	return CreateMat4(
		this.M00 * f,
		this.M01 * f,
		this.M02 * f,
		this.M03 * f,
		this.M10 * f,
		this.M11 * f,
		this.M12 * f,
		this.M13 * f,
		this.M20 * f,
		this.M21 * f,
		this.M22 * f,
		this.M23 * f,
		this.M30 * f,
		this.M31 * f,
		this.M32 * f,
		this.M33 * f)
}

func (this *Mat4) DivScalar(f float32) *Mat4 {
	return CreateMat4(
		this.M00 / f,
		this.M01 / f,
		this.M02 / f,
		this.M03 / f,

		this.M10 / f,
		this.M11 / f,
		this.M12 / f,
		this.M13 / f,

		this.M20 / f,
		this.M21 / f,
		this.M22 / f,
		this.M23 / f,

		this.M30 / f,
		this.M31 / f,
		this.M32 / f,
		this.M33 / f)
}

func (this *Mat4) Mul(other *Mat4) *Mat4 {
	return CreateMat4(
		this.M00 * other.M00 + this.M01 * other.M10 + this.M02 * other.M20 + this.M03 * other.M30,
		this.M00 * other.M01 + this.M01 * other.M11 + this.M02 * other.M21 + this.M03 * other.M31,
		this.M00 * other.M02 + this.M01 * other.M12 + this.M02 * other.M22 + this.M03 * other.M32,
		this.M00 * other.M03 + this.M01 * other.M13 + this.M02 * other.M23 + this.M03 * other.M33,
		
		this.M10 * other.M00 + this.M11 * other.M10 + this.M12 * other.M20 + this.M13 * other.M30,
		this.M10 * other.M01 + this.M11 * other.M11 + this.M12 * other.M21 + this.M13 * other.M31,
		this.M10 * other.M02 + this.M11 * other.M12 + this.M12 * other.M22 + this.M13 * other.M32,
		this.M10 * other.M03 + this.M11 * other.M13 + this.M12 * other.M23 + this.M13 * other.M33,
		
		this.M20 * other.M00 + this.M21 * other.M10 + this.M22 * other.M20 + this.M23 * other.M30,
		this.M20 * other.M01 + this.M21 * other.M11 + this.M22 * other.M21 + this.M23 * other.M31,
		this.M20 * other.M02 + this.M21 * other.M12 + this.M22 * other.M22 + this.M23 * other.M32,
		this.M20 * other.M03 + this.M21 * other.M13 + this.M22 * other.M23 + this.M23 * other.M33,
		
		this.M30 * other.M00 + this.M31 * other.M10 + this.M32 * other.M20 + this.M33 * other.M30,
		this.M30 * other.M01 + this.M31 * other.M11 + this.M32 * other.M21 + this.M33 * other.M31,
		this.M30 * other.M02 + this.M31 * other.M12 + this.M32 * other.M22 + this.M33 * other.M32,
		this.M30 * other.M03 + this.M31 * other.M13 + this.M32 * other.M23 + this.M33 * other.M33,
	)
}

func (this *Mat4) Transpose() *Mat4 {
	return CreateMat4(
		this.M00, this.M10, this.M20, this.M30,
		this.M01, this.M11, this.M21, this.M31,
		this.M02, this.M12, this.M22, this.M32,
		this.M03, this.M13, this.M23, this.M33,
	)
}

func (this *Mat4) Trace() float32 {
	return this.M00 + this.M11 + this.M22 + this.M33
}

func (this *Mat4) Det() float32 {
	var ma, mb, mc, md *Mat3

	ma = CreateMat3(
		this.M11, this.M12, this.M13,
		this.M21, this.M22, this.M23,
		this.M31, this.M32, this.M33)

	mb = CreateMat3(
		this.M10, this.M12, this.M13,
		this.M20, this.M22, this.M23,
		this.M30, this.M32, this.M33)

	mc = CreateMat3(
		this.M10, this.M11, this.M13,
		this.M20, this.M21, this.M23,
		this.M30, this.M31, this.M33)

	md = CreateMat3(
		this.M10, this.M11, this.M12,
		this.M20, this.M21, this.M22,
		this.M30, this.M31, this.M32)

	return this.M00 * ma.Det() - this.M01 * mb.Det() + this.M02 * mc.Det() - this.M03 * md.Det()
}

func (this *Mat4) Inverse() *Mat4 {
	return CreateMat4(
		// Row 1
		CreateMat3(
			this.M11, this.M12, this.M13, 
			this.M21, this.M22, this.M23, 
			this.M31, this.M32, this.M33).Det(),

		-CreateMat3(
			this.M10, this.M12, this.M13, 
			this.M20, this.M22, this.M23, 
			this.M30, this.M32, this.M33).Det(),

		CreateMat3(
			this.M10, this.M11, this.M13, 
			this.M20, this.M21, this.M23, 
			this.M30, this.M31, this.M33).Det(),

		-CreateMat3(
			this.M10, this.M11, this.M12, 
			this.M20, this.M21, this.M22, 
			this.M30, this.M31, this.M32).Det(),

		// Row 2
		-CreateMat3(
			this.M01, this.M02, this.M03, 
			this.M21, this.M22, this.M23, 
			this.M31, this.M32, this.M33).Det(),

		CreateMat3(
			this.M00, this.M02, this.M03, 
			this.M20, this.M22, this.M23, 
			this.M30, this.M32, this.M33).Det(),

		-CreateMat3(
			this.M00, this.M01, this.M03, 
			this.M20, this.M21, this.M23, 
			this.M30, this.M31, this.M33).Det(),

		CreateMat3(
			this.M00, this.M01, this.M02, 
			this.M20, this.M21, this.M22, 
			this.M30, this.M31, this.M32).Det(),
		
		// Row 3
		CreateMat3(
			this.M01, this.M02, this.M03, 
			this.M11, this.M12, this.M13, 
			this.M31, this.M32, this.M33).Det(),

		-CreateMat3(
			this.M00, this.M02, this.M03,
			this.M10, this.M12, this.M13, 
			this.M30, this.M32, this.M33).Det(),

		CreateMat3(
			this.M00, this.M01, this.M03, 
			this.M10, this.M11, this.M13, 
			this.M30, this.M31, this.M33).Det(),

		-CreateMat3(
			this.M00, this.M01, this.M02, 
			this.M10, this.M11, this.M12, 
			this.M30, this.M31, this.M32).Det(),
		
		// Row 4
		-CreateMat3(
			this.M01, this.M02, this.M03, 
			this.M11, this.M12, this.M13, 
			this.M21, this.M22, this.M23).Det(),

		CreateMat3(
			this.M00, this.M02, this.M03, 
			this.M10, this.M12, this.M13, 
			this.M20, this.M22, this.M23).Det(),

		-CreateMat3(
			this.M00, this.M01, this.M03, 
			this.M10, this.M11, this.M13, 
			this.M20, this.M21, this.M23).Det(),

		CreateMat3(
			this.M00, this.M01, this.M02, 
			this.M10, this.M11, this.M12, 
			this.M20, this.M21, this.M22).Det(),

	).Transpose().DivScalar(this.Det())
}

func (this *Mat4) MulVec4(v *Vec4) *Vec4 {
	return CreateVec4(
		this.M00 * v.X + this.M01 * v.Y + this.M02 * v.Z + this.M03 * v.W,
		this.M10 * v.X + this.M11 * v.Y + this.M12 * v.Z + this.M13 * v.W,
		this.M20 * v.X + this.M21 * v.Y + this.M22 * v.Z + this.M23 * v.W,
		this.M30 * v.X + this.M31 * v.Y + this.M32 * v.Z + this.M33 * v.W,
	)
}

func (this *Mat4) ToArray() []float32 {
	var temp []float32
	temp = append(temp, this.M00)
	temp = append(temp, this.M01)
	temp = append(temp, this.M02)
	temp = append(temp, this.M03)

	temp = append(temp, this.M10)
	temp = append(temp, this.M11)
	temp = append(temp, this.M12)
	temp = append(temp, this.M13)

	temp = append(temp, this.M20)
	temp = append(temp, this.M21)
	temp = append(temp, this.M22)
	temp = append(temp, this.M23)

	temp = append(temp, this.M30)
	temp = append(temp, this.M31)
	temp = append(temp, this.M32)
	temp = append(temp, this.M33)

	return temp
}