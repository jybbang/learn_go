package main

import "fmt"

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkCode(incomingcode int) error {
	if s.code != incomingcode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode verified")
	return nil
}
