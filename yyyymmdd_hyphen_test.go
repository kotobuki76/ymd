package ymd

import "testing"

func TestHyphenToYmd(t *testing.T) {

	src := "1920-02-12"
	want := "19200212"
	got := HyphenYmd(src).ToYmd()

	if want != got.String() {
		t.Errorf("HyphenToYmd got %s want %s", got, want)
	}
}

func TestYmdToHyphen(t *testing.T) {

	want := "1920-02-12"
	src := "19200212"
	got := Ymd(src).ToHyphen()

	if want != got.String() {
		t.Errorf("YmdToHyphen got %s want %s", got, want)
	}
}
