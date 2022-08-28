package subtraction

import "GolangMK/maths/addition"

func Maths2() uint32 {
	var FirstUint uint32 = 12
	var AnotherUint uint32 = 0

	AnotherUint = uint32(addition.Maths1()) - FirstUint
	return AnotherUint
}
