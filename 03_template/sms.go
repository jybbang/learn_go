package main

import "fmt"

type sms struct {
	otp
}

func (s *sms) genRandomOTP(len int) string {
	randomOtp := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOtp)
	return randomOtp
}
func (s *sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: save otp %s to cache\n", otp)
}
func (s *sms) getMessage(otp string) string {
	return "SMS Otp for login is " + otp
}
func (s *sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms %s\n", message)
	return nil
}
func (s *sms) publishMetric() {
	fmt.Printf("SMS: publising metrics\n")
}
