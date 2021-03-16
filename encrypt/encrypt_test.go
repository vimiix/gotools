package encrypt

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	key := "123456789abcdefg"
	tt := []struct {
		in     string
		expect string
	}{
		{"1", "GcM+2sD0Tw3Nu8yL+PwIhQ=="},
		{"test", "qMmHE08QXslbu+Zy8emBLQ=="},
	}
	for _, c := range tt {
		got, _ := Encrypt(c.in, key)
		if got != c.expect {
			t.Errorf("test encrypt fail. expect \"%s\" got \"%s\"", c.expect, got)
		}
	}
}

func TestDecrypt(t *testing.T) {
	key := "123456789abcdefg"
	tt := []struct {
		in        string
		expect    string
		expectErr string
	}{
		{"GcM+2sD0Tw3Nu8yL+PwIhQ==", "1", ""},
		{"qMmHE08QXslbu+Zy8emBLQ==", "test", ""},
		{"GcM+2sD0Tw3Nu8yLPwIhQ==", "", "crypto/cipher: input not full blocks"},
	}
	for _, c := range tt {
		got, err := Decrypt(c.in, key)
		if err != nil {
			if c.expectErr != err.Error() {
				t.Errorf("test encrypt fail. expectErr: \"%s\", got \"%s\"", c.expectErr, err)
			}
		}
		if got != c.expect {
			t.Errorf("test encrypt fail. expect \"%s\" got \"%s\"", c.expect, got)
		}
	}
}
