package otp

type OTP_FORMAT int

const (
	_                       OTP_FORMAT = iota
	OTP_FORMAT_NUMERIC                 // 1
	OTP_FORMAT_ALPHANUMERIC            // 2
	OTP_FORMAT_ALPHABETIC              // 3
)
