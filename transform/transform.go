package transform

import "math"
import "derf.space/gldemo/linearmath"

func Translate(v *linearmath.Vec3) *linearmath.Mat4 {
	var temp *linearmath.Mat4 = linearmath.CreateMat4Identity()
	temp.M30 = v.X
	temp.M31 = v.Y
	temp.M32 = v.Z
	return temp
}

func RotateX(radian float32) *linearmath.Mat4 {
	var temp *linearmath.Mat4 = linearmath.CreateMat4Identity()
	temp.M11 = float32(math.Cos(float64(radian)))
	temp.M12 = float32(-math.Sin(float64(radian)))
	temp.M21 = float32(math.Sin(float64(radian)))
	temp.M22 = float32(math.Cos(float64(radian)))
	return temp
}

func RotateY(radian float32) *linearmath.Mat4 {
	var temp *linearmath.Mat4 = linearmath.CreateMat4Identity()
	temp.M00 = float32(math.Cos(float64(radian)))
	temp.M02 = float32(math.Sin(float64(radian)))
	temp.M20 = float32(-math.Sin(float64(radian)))
	temp.M22 = float32(math.Cos(float64(radian)))
	return temp
}

func RotateZ(radian float32) *linearmath.Mat4 {
	var temp *linearmath.Mat4 = linearmath.CreateMat4Identity()
	temp.M00 = float32(math.Cos(float64(radian)))
	temp.M01 = float32(-math.Sin(float64(radian)))
	temp.M10 = float32(math.Sin(float64(radian)))
	temp.M11 = float32(math.Cos(float64(radian)))
	return temp
}

func RotateAxis(radian float32, axis *linearmath.Vec3) *linearmath.Mat4 {
	var c float32 = float32(math.Cos(float64(radian)))
	var s float32 = float32(math.Sin(float64(radian)))

	if axis.Length() != 1 {
		axis = axis.Normalize()
	}

	var temp *linearmath.Mat4 = linearmath.CreateMat4Identity()
	temp.M00 = (axis.X * axis.X) * (1 - c) + c
	temp.M01 = (axis.X * axis.Y) * (1 - c) - axis.Z * s
	temp.M02 = (axis.X * axis.Z) * (1 - c) + axis.Y * s
	temp.M10 = (axis.X * axis.Y) * (1 - c) + axis.Z * s
	temp.M11 = (axis.Y * axis.Y) * (1 - c) + c
	temp.M12 = (axis.Y * axis.Z) * (1 - c) - axis.X * s
	temp.M20 = (axis.X * axis.Z) * (1 - c) - axis.Y * s
	temp.M21 = (axis.Y * axis.Z) * (1 - c) + axis.X * s
	temp.M22 = (axis.Z * axis.Z) * (1 - c) + c
	return temp
}

func Scale(v *linearmath.Vec3) *linearmath.Mat4 {
	var temp *linearmath.Mat4 = linearmath.CreateMat4Identity()
	temp.M00 = v.X
	temp.M11 = v.Y
	temp.M22 = v.Z
	return temp
}

func Ortho(left float32, right float32, bottom float32, top float32, zmin float32, zmax float32) *linearmath.Mat4 {
	var temp *linearmath.Mat4 = linearmath.CreateMat4Identity()
	
	var tx float32 = -((right + left) / (right - left))
	var ty float32 = -((top + bottom) / (top - bottom))
	var tz float32 = -((zmax + zmin) / (zmax - zmin))

	temp.M00 = 2.0 / (right - left)
	temp.M11 = 2.0 / (top - bottom)
	temp.M22 = -2.0 / (zmax - zmin)
	temp.M30 = tx
	temp.M31 = ty
	temp.M32 = tz

	return temp
}

func Ortho2D(left float32, right float32, bottom float32, top float32) *linearmath.Mat4 {
	return Ortho(left, right, bottom, top, -1.0, 1.0)
}

func Frustum(left float32, right float32, bottom float32, top float32, zmin float32, zmax float32) *linearmath.Mat4 {
	var temp *linearmath.Mat4 = linearmath.CreateMat4Identity()

	var A float32 = (right + left) / (right - left)
	var B float32 = (top + bottom) / (top - bottom)
	var C float32 = -((zmax + zmin) / (zmax - zmin))
	var D float32 = -((2.0 * zmax * zmin) / (zmax - zmin))

	temp.M00 = (2.0 * zmin) / (right - left)
	temp.M11 = (2.0 * zmin) / (top - bottom)
	temp.M20 = A
	temp.M21 = B
	temp.M22 = C
	temp.M23 = -1
	temp.M32 = D

	return temp
}

func Perspective(fov float32, aspect float32, zmin float32, zmax float32) *linearmath.Mat4 {
	var ymax, xmax float32
	ymax = zmin * float32(math.Tan(float64(fov)))
	xmax = ymax * aspect
	return Frustum(-xmax, xmax, -ymax, ymax, zmin, zmax)
}
