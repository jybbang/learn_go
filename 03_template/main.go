package main

import "fmt"

func main() {
	smsOTP := &sms{}
	o := otp{
		iOtp: smsOTP,
	}
	o.genAndSendOtp(4)

	fmt.Println()

	emailOTP := &email{}
	o = otp{
		iOtp: emailOTP,
	}
	o.genAndSendOtp(4)
}
