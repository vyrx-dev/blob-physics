package main

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		name           string
		x1, y1, x2, y2 float32
		expected       float32
	}{
		{"same point", 0, 0, 0, 0, 0},
		{"3-4-5 triangle", 0, 0, 3, 4, 5},
		{"horizontal line", 0, 5, 10, 5, 10},
		{"vertical line", 7, 0, 7, 24, 24},
		{"negative coordinates", -3, -4, 0, 0, 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := distance(tc.x1, tc.y1, tc.x2, tc.y2)
			if math.Abs(float64(got-tc.expected)) > 1e-5 {
				t.Errorf("distance(%v,%v, %v,%v) = %v, want %v",
					tc.x1, tc.y1, tc.x2, tc.y2, got, tc.expected)
			}
		})
	}
}

func TestDistanceIsSymmetric(t *testing.T) {
	pairs := [][4]float32{
		{1, 2, 5, 6},
		{-10, 3, 4, -7},
		{100, 200, 300, 400},
	}
	for _, pair := range pairs {
		fromAtoB := distance(pair[0], pair[1], pair[2], pair[3])
		fromBtoA := distance(pair[2], pair[3], pair[0], pair[1])
		if fromAtoB != fromBtoA {
			t.Errorf("distance not symmetric for %v: got %v and %v", pair, fromAtoB, fromBtoA)
		}
	}
}

func TestDistanceIsNeverNegative(t *testing.T) {
	pairs := [][4]float32{
		{0, 0, 0, 0},
		{-5, -5, 5, 5},
		{1000, -1000, -1000, 1000},
	}
	for _, pair := range pairs {
		if result := distance(pair[0], pair[1], pair[2], pair[3]); result < 0 {
			t.Errorf("got negative distance %v for %v", result, pair)
		}
	}
}

func TestMoveToward(t *testing.T) {
	tests := []struct {
		name      string
		current   float32
		target    float32
		deltaTime float32
		speed     float32
		expected  float32
	}{
		// already there, should stay put
		{"already at target", 5, 5, 0.016, 15, 5},

		// next = 0 + 0.1 * (10 - 0) * 1.0 = 1.0
		{"simple step", 0, 10, 0.1, 1.0, 1.0},

		// next = 100 + 0.016 * (0 - 100) * 15 = 76
		{"approaching from above", 100, 0, 0.016, 15, 76},

		// next = 0 + 0.016 * (-50 - 0) * 15 = -12
		{"negative target", 0, -50, 0.016, 15, -12},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := moveToward(tc.current, tc.target, tc.deltaTime, tc.speed)
			if math.Abs(float64(got-tc.expected)) > 1e-4 {
				t.Errorf("moveToward(%v, %v, %v, %v) = %v, want %v",
					tc.current, tc.target, tc.deltaTime, tc.speed, got, tc.expected)
			}
		})
	}
}

// simulate 300 frames and make sure the ball never flies past the target
func TestMoveTowardNeverOvershoots(t *testing.T) {
	currentPosition := float32(0)
	targetPosition := float32(400)

	for frame := 0; frame < 300; frame++ {
		next := moveToward(currentPosition, targetPosition, 1.0/60.0, 15.0)
		if next > targetPosition {
			t.Fatalf("frame %d: overshot target %.4f (got %.4f)", frame, targetPosition, next)
		}
		currentPosition = next
	}

	// should be basically at the target after ~5 simulated seconds
	if math.Abs(float64(currentPosition-targetPosition)) > 0.01 {
		t.Errorf("didn't converge: stuck at %v, target was %v", currentPosition, targetPosition)
	}
}

func BenchmarkDistance(b *testing.B) {
	for b.Loop() {
		distance(10, 20, 30, 40)
	}
}

func BenchmarkMoveToward(b *testing.B) {
	pos := float32(0)
	for b.Loop() {
		pos = moveToward(pos, 400, 0.016, 15)
	}
}
