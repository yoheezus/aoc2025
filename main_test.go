package main

import (
	"testing"
)

func newWheel() Wheel {
	w := Wheel{
		length: 100,
		idx:    50,
	}
	w.NewWheel()
	return w
}

func newWheelIdx(idx int) Wheel {
	w := Wheel{
		length: 100,
		idx:    idx,
	}
	w.NewWheel()
	return w
}

func TestLeftSpin(t *testing.T) {
	w := newWheel()

	want := 45
	result := w.Spin(-5)

	if want != result {
		t.Errorf(`w.Spin(-5) = %q, want %d`, result, want)
	}
}

func TestRightSpin(t *testing.T) {
	w := newWheel()

	want := 55
	result := w.Spin(5)

	if want != result {
		t.Errorf(`w.Spin(-5) = %q, want %d`, result, want)
	}
}

func TestLandsOnZero(t *testing.T) {
	w := newWheel()
	want := 3

	instructions := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	unfiltered := DoSpins(instructions, &w)
	var zero_count int
	for i := range unfiltered {
		if unfiltered[i] == 0 {
			zero_count += 1
		}
	}

	if zero_count != want {
		t.Errorf(`DoSpins(%q) = %q, want %d`, instructions, zero_count, want)
	}
}

func TestLandsAndPassesOnZero(t *testing.T) {
	w := newWheel()

	instructions := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	unfiltered := DoSpins(instructions, &w)
	var zero_count int
	for i := range unfiltered {
		if unfiltered[i] == 0 {
			zero_count += 1
		}
	}

	want_zero_count := 3
	want_zero_passes := 3
	want_total := want_zero_count + want_zero_passes
	t.Errorf(`DoSpins(%v) = %v, want %d sum of zero count and zero passes\n
		total Zero count: %d\n
		total Zero passes %d\n`, instructions, zero_count+w.clicks, want_total, zero_count, w.clicks)
}

func TestStartsonZeroEndsOnZero(t *testing.T) {
	w := newWheelIdx(0)
	instructions := []int{200}
	unfiltered := DoSpins(instructions, &w)
	var zero_count int
	for i := range unfiltered {
		if unfiltered[i] == 0 {
			zero_count += 1
		}
	}

	want_zero_count := 1
	want_zero_passes := 1
	want_total := want_zero_count + want_zero_passes

	if want_total != w.clicks+zero_count {
		t.Errorf(`DoSpins(%v) = %v, want %d sum of zero count and zero passes\n
		total Zero count: %d\n
		total Zero passes %d\n`, instructions, zero_count+w.clicks, want_total, zero_count, w.clicks)

	}

}
