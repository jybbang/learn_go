package main

import "fmt"

type email struct {
	otp
}

func (s *email) genRandomOTP(len int) string {
	randomOtp := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOtp)
	return randomOtp
}
func (s *email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: save otp %s to cache\n", otp)
}
func (s *email) getMessage(otp string) string {
	return "EMAIL Otp for login is " + otp
}
func (s *email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email %s\n", message)
	return nil
}
func (s *email) publishMetric() {
	fmt.Printf("EMAIL: spublising metrics\n")
}
