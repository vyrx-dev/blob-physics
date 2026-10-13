package main

import "math"

// distance returns the Euclidean distance between two 2D points.
func distance(x1, y1, x2, y2 float32) float32 {
	dx := x1 - x2
	dy := y1 - y2
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// moveToward steps current toward target by one frame tick.
// Uses first-order frame smoothing: current + dt * (target - current) * speed.
func moveToward(current, target, deltaTime, speed float32) float32 {
	return current + deltaTime*(target-current)*speed
}
