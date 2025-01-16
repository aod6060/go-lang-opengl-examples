package linearmath

const PI float32 = 3.14159

func ToDegree(rad float32) float32 {
	return rad * (180.0 / PI)
}

func ToRadian(deg float32) float32 {
	return deg * (PI / 180.0)
}
