package ymd

import "testing"

func TestValidateYmd(t *testing.T) {
	t.Logf("%s", ValidateYmd("29011210"))
	t.Logf("%s", ValidateYmd("99999999"))
	t.Logf("%s", ValidateYmd("29d1121a"))
	t.Logf("%s", ValidateYmd("2901131"))
}
