package gutil

import (
	"github.com/mojocn/base64Captcha"
)

func captchaConfig(width, height, length, mode int, isUserSimpleFont bool, complexOfNoiseDot, comlexOfNoiseText int, isShowHollowLine, isShowSlimeLine, isShowSineLIne bool) base64Captcha.ConfigCharacter {
	return base64Captcha.ConfigCharacter{
		Height:             height,
		Width:              width,
		Mode:               mode,
		IsUseSimpleFont:    isUserSimpleFont,
		ComplexOfNoiseDot:  complexOfNoiseDot,
		ComplexOfNoiseText: comlexOfNoiseText,
		IsShowHollowLine:   isShowHollowLine,
		IsShowSlimeLine:    isShowSlimeLine,
		IsShowSineLine:     isShowSineLIne,
		CaptchaLen:         length,
	}
}
func captchaCreate(cfg base64Captcha.ConfigCharacter) (string, string) {
	id, capd := base64Captcha.GenerateCaptcha("", cfg)
	base64 := base64Captcha.CaptchaWriteToBase64Encoding(capd)
	return id, base64
}

// CaptchaSimpleNumber 简单难度数字的验证码
func CaptchaSimpleNumber(width, height, length int) (id string, base64 string) {
	cfg := captchaConfig(width, height, length, base64Captcha.CaptchaModeNumber, true, base64Captcha.CaptchaComplexLower, base64Captcha.CaptchaComplexLower, false, false, false)
	return captchaCreate(cfg)
}

// CaptchaMediumNumber 中等难度数字的验证码
func CaptchaMediumNumber(width, height, length int) (id string, base64 string) {
	cfg := captchaConfig(width, height, length, base64Captcha.CaptchaModeNumber, true, base64Captcha.CaptchaComplexMedium, base64Captcha.CaptchaComplexMedium, false, true, true)
	return captchaCreate(cfg)
}

// CaptchaComplexNumber 复杂难度数字的验证码
func CaptchaComplexNumber(width, height, length int) (id string, base64 string) {
	cfg := captchaConfig(width, height, length, base64Captcha.CaptchaModeNumber, false, base64Captcha.CaptchaComplexHigh, base64Captcha.CaptchaComplexHigh, true, true, false)
	return captchaCreate(cfg)
}

// CaptchaSimpleLetter 简单难度字母的验证码
func CaptchaSimpleLetter(width, height, length int) (id string, base64 string) {
	cfg := captchaConfig(width, height, length, base64Captcha.CaptchaModeAlphabet, true, base64Captcha.CaptchaComplexLower, base64Captcha.CaptchaComplexLower, false, false, false)
	return captchaCreate(cfg)
}

// CaptchaMediumLetter 中等难度字母的验证码
func CaptchaMediumLetter(width, height, length int) (id string, base64 string) {
	cfg := captchaConfig(width, height, length, base64Captcha.CaptchaModeAlphabet, true, base64Captcha.CaptchaComplexMedium, base64Captcha.CaptchaComplexMedium, false, true, true)
	return captchaCreate(cfg)
}

// CaptchaComplexLetter 复杂难度字母的验证码
func CaptchaComplexLetter(width, height, length int) (id string, base64 string) {
	cfg := captchaConfig(width, height, length, base64Captcha.CaptchaModeAlphabet, false, base64Captcha.CaptchaComplexHigh, base64Captcha.CaptchaComplexHigh, true, true, false)
	return captchaCreate(cfg)
}

// CaptchaSimpleNumberAndLetter 简单难度数字字母的验证码
func CaptchaSimpleNumberAndLetter(width, height, length int) (id string, base64 string) {
	cfg := captchaConfig(width, height, length, base64Captcha.CaptchaModeNumberAlphabet, true, base64Captcha.CaptchaComplexLower, base64Captcha.CaptchaComplexLower, false, false, false)
	return captchaCreate(cfg)
}

// CaptchaMediumNumberAndLetter 中等难度数字字母的验证码
func CaptchaMediumNumberAndLetter(width, height, length int) (id string, base64 string) {
	cfg := captchaConfig(width, height, length, base64Captcha.CaptchaModeNumberAlphabet, true, base64Captcha.CaptchaComplexMedium, base64Captcha.CaptchaComplexMedium, false, true, true)
	return captchaCreate(cfg)
}

// CaptchaComplexNumberAndLetter 复杂难度数字字母的验证码
func CaptchaComplexNumberAndLetter(width, height, length int) (id string, base64 string) {
	cfg := captchaConfig(width, height, length, base64Captcha.CaptchaModeNumberAlphabet, false, base64Captcha.CaptchaComplexHigh, base64Captcha.CaptchaComplexHigh, true, true, false)
	return captchaCreate(cfg)
}

// CaptchaValidate 验证码验证
func CaptchaValidate(id, code string) bool {
	return base64Captcha.VerifyCaptcha(id, code)
}
