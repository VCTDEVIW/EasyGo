//package main
package otp

import (
	"fmt"

	"github.com/xlzd/gotp"
)

func main() {
	fmt.Println("Random secret:", gotp.RandomSecret(16))
	DefaultTOTPUsage()
	DefaultHOTPUsage()
}

func InitFn() {
	fmt.Println("Random secret:", gotp.RandomSecret(16))
	DefaultTOTPUsage()
	DefaultHOTPUsage()
}

func DefaultTOTPUsage() {
	otp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")

	fmt.Println("current one-time password is:", otp.Now())
	fmt.Println("one-time password of timestamp 0 is:", otp.At(0))
	fmt.Println(otp.ProvisioningUri("demoAccountName", "issuerName"))

	fmt.Println(otp.Verify("179394", 1524485781))
}

func DefaultHOTPUsage() {
	otp := gotp.NewDefaultHOTP("4S62BZNFXXSZLCRO")

	fmt.Println("one-time password of counter 0 is:", otp.At(0))
	fmt.Println(otp.ProvisioningUri("demoAccountName", "issuerName", 1))

	fmt.Println(otp.Verify("944181", 0))
}