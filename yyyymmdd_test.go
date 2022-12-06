package ymd

import (
	"testing"
)

func TestExplodeYmd(t *testing.T) {

	want_y := "1920"
	want_m := "02"
	want_d := "12"
	src := "19200212"
	got_y, got_m, got_d := Ymd(src).Explode()

	if want_y != got_y {
		t.Errorf("ExplodeYmd got %s want %s", got_y, want_y)
	}

	if want_m != got_m {
		t.Errorf("ExplodeYmd got %s want %s", got_m, want_m)
	}

	if want_d != got_d {
		t.Errorf("ExplodeYmd got %s want %s", got_d, want_d)
	}
}

func TestAddDay(t *testing.T) {

	want := "19200214"
	src_c := 2
	src := "19200212"
	got := Ymd(src).AddDay(src_c)

	if want != got.String() {
		t.Errorf("ToAddedYmd got %s want %s", got, want)
	}
}

func TestFirstDateOfMonth(t *testing.T) {

	want := "19200201"
	src := "19200212"
	got := Ymd(src).FirstDateOfMonth()

	if want != got.String() {
		t.Errorf("FirstDateOfMonth got %s want %s", got, want)
	}
}

func TestEndDateOfMonth(t *testing.T) {

	want := "19200229"
	src := "19200212"
	got := Ymd(src).EndDateOfMonth()

	if want != got.String() {
		t.Errorf("EndDateOfMonth got %s want %s", got, want)
	}
}

func TestBetween(t *testing.T) {

	want := 8
	src1 := "19200212"
	src2 := "19200220"
	got := Ymd(src1).Between(Ymd(src2))

	if want != got {
		t.Errorf("Between got %d want %d", got, want)
	}

	t.Log(got)
}

func TestAddYear(t *testing.T) {

	want := "20220212"
	src_c := 2
	src := "20200212"
	got := Ymd(src).AddYear(src_c)

	if want != got.String() {
		t.Errorf("ToAddedYmd got %s want %s", got, want)
	}
}

func TestAddYear2(t *testing.T) {

	want := "20200212"
	src_c := -2
	src := "20220212"
	got := Ymd(src).AddYear(src_c)

	if want != got.String() {
		t.Errorf("ToAddedYmd got %s want %s", got, want)
	}
}

func TestTodayYmdhms(t *testing.T) {
	t.Log(TodayYmdhms())
}
