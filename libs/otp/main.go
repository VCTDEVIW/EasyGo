package otp

import (
	."fmt"
	."project/libs/lib_timestamp"

	otp "github.com/xlzd/gotp"
)

func T() {
	Println("**** [DEBUG]: Mod {otp} loaded.")
	Println( GenRanSec(24) )
	Println("Timestamp() => ", Time("beforeHour", 1))
	Println()
}

// "Generate Random Secret" supporting invoker(s): GenRanSec(), GenRanSec(param int)
func GenRanSec(param ...interface{}) string {
	if (len(param) == 0) {
		return otp.RandomSecret(16)
	}

	return otp.RandomSecret(param[0].(int))
}

