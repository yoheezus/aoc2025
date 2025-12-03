package main

import (
	"reflect"
	"testing"
)

func newWheel() Wheel {
	w := Wheel{
		length: 100,
		idx:    50,
	}
	return w
}

func newWheelIdx(idx int) Wheel {
	w := Wheel{
		length: 100,
		idx:    idx,
	}
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
	} else {
		t.Logf(`total Zero count: %d\n
		total Zero passes %d\n`, zero_count, w.clicks)
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

	if want_total != w.clicks+zero_count {
		t.Errorf(`DoSpins(%v) = %v, want %d sum of zero count and zero passes\n
		total Zero count: %d\n
		total Zero passes %d\n`, instructions, zero_count+w.clicks, want_total, zero_count, w.clicks)
	} else {
		t.Logf(`total Zero count: %d\n
		total Zero passes %d\n`, zero_count, w.clicks)
	}
}

func TestStartsonZeroEndsOnZeroRight(t *testing.T) {
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
	want_idx := 0

	if want_total != w.clicks+zero_count && want_idx == w.idx {
		t.Errorf(`DoSpins(%v) = %v, want %d sum of zero count and zero passes\n
		total Zero count: %d\n
		total Zero passes %d\n`, instructions, zero_count+w.clicks, want_total, zero_count, w.clicks)

	} else {
		t.Logf(`total Zero count: %d\n
		total Zero passes %d\n`, zero_count, w.clicks)
	}

}

func TestStartsonZeroEndsOnZeroLeft(t *testing.T) {
	w := newWheelIdx(0)
	instructions := []int{-200}
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
	want_idx := 0

	if want_total != w.clicks+zero_count && want_idx == w.idx {
		t.Errorf(`DoSpins(%v) = %v, want %d sum of zero count and zero passes\n
		total Zero count: %d\n
		total Zero passes %d\n`, instructions, zero_count+w.clicks, want_total, zero_count, w.clicks)

	} else {
		t.Logf(`total Zero count: %d\n
		total Zero passes %d\n`, zero_count, w.clicks)
	}

}

func TestRight1000At50(t *testing.T) {
	w := newWheelIdx(50)
	instructions := []int{1000}
	unfiltered := DoSpins(instructions, &w)
	var zero_count int
	for i := range unfiltered {
		if unfiltered[i] == 0 {
			zero_count += 1
		}
	}

	want_zero_count := 0
	want_zero_passes := 10
	want_total := want_zero_count + want_zero_passes
	want_idx := 50

	if want_total != w.clicks+zero_count && want_idx == w.idx {
		t.Errorf(`DoSpins(%v) = %v, want %d sum of zero count and zero passes\n
		total Zero count: %d\n
		total Zero passes %d\n`, instructions, zero_count+w.clicks, want_total, zero_count, w.clicks)

	} else {
		t.Logf(`total Zero count: %d\n
		total Zero passes %d\n`, zero_count, w.clicks)
	}
}

func TestRight1000At0(t *testing.T) {
	w := newWheelIdx(0)
	instructions := []int{1000}
	unfiltered := DoSpins(instructions, &w)
	var zero_count int
	for i := range unfiltered {
		if unfiltered[i] == 0 {
			zero_count += 1
		}
	}

	want_zero_count := 1
	want_zero_passes := 9
	want_total := want_zero_count + want_zero_passes
	want_idx := 0

	if want_total != w.clicks+zero_count && want_idx == w.idx {
		t.Errorf(`DoSpins(%v) = %v, want %d sum of zero count and zero passes\n
		total Zero count: %d\n
		total Zero passes %d\n`, instructions, zero_count+w.clicks, want_total, zero_count, w.clicks)

	} else {
		t.Logf(`total Zero count: %d\n
		total Zero passes %d\n`, zero_count, w.clicks)
	}
}

func TestPastZeroOnce(t *testing.T) {
	w := newWheel()
	spin := -70

	want_idx := 80
	want_clicks := 1
	result := w.Spin(spin)

	if want_idx != w.idx && w.clicks != want_clicks {
		t.Errorf(`DoSpins(%v) = %v, total clicks: %d, index: %d`, spin, result, w.clicks, w.idx)

	} else {
		t.Logf(`DoSpins(%v) = %v, total clicks: %d, index: %d`, spin, result, w.clicks, w.idx)
	}
}

func TestPast99Once(t *testing.T) {
	w := newWheel()
	spin := 55

	want_idx := 5
	want_clicks := 1
	result := w.Spin(55)

	if want_idx != w.idx && w.clicks != want_clicks {
		t.Errorf(`DoSpins(%v) = %v, total clicks: %d, index: %d`, spin, result, w.clicks, w.idx)

	} else {
		t.Logf(`DoSpins(%v) = %v, total clicks: %d, index: %d`, spin, result, w.clicks, w.idx)
	}
}

func TestPastZeroThree(t *testing.T) {
	w := newWheelIdx(5)
	instructions := []int{-7, 5, -5, 2, 40}

	want_idx := []int{98, 3, 98, 0, 40}
	want_clicks := 6
	result := DoSpins(instructions, &w)

	if !reflect.DeepEqual(want_idx, result) && w.clicks != want_clicks {
		t.Errorf(`DoSpins(%v) = %v, total clicks: %d, index: %d`, instructions, result, w.clicks, w.idx)

	} else {
		t.Logf(`DoSpins(%v) = %v, total clicks: %d, index: %d`, instructions, result, w.clicks, w.idx)
	}
}
