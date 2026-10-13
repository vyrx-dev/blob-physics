# blob-physics

A 2D liquid blob physics prototype written in Go using [raylib-go](https://github.com/gen2brain/raylib-go).

This is the sandbox I built to prototype liquid blob simulation and lag smoothing before building my game **[Tez](https://github.com/vyrx-dev/tez)**. I wanted to isolate the math, test the visual smear feel, and benchmark the frame step in isolation before building typing mechanics, game states, and audio on top of it.

Inspired by [Rexim's blob](https://gist.github.com/rexim/c1e3c1b99cb824ab6c817c3f45d75564), rewritten and tested in Go.

---

## Evolution to Tez

* **In this prototype (`blob-physics`)**: Proved out the two-point trailing concept (a head chasing the cursor and a tail trailing behind at 45% speed) with a stretched semicircle dome drawn between them.
* **In the game (`tez`)**: Took the verified blob logic and upgraded the first-order `dt * speed` step to frame-rate independent exponential smoothing (`1 - exp(-speed * dt)`), added angle damping across the 360° seam, and built the full typing game around it.

---

## How It Works

1. **Head (`ball`)**: Chases the target point (mouse cursor) each frame.
2. **Tail (`smear`)**: Chases the head with artificial lag (`smearChaseSpeed = ballChaseSpeed * 0.45`).
3. **Stretched Dome**: Calculates the distance and angle between head and tail, stretching the top half of a circle texture between them to give it an elastic, liquid feel.

---

## Running Locally

Requires Go 1.24+. You do **not** need a C compiler or system Raylib libraries installed.

### Standard (Linux, macOS, Windows)
Uses pure Go via `purego` with an embedded Raylib binary:
```sh
CGO_ENABLED=0 go run .
```

If you have Raylib installed on your system and prefer CGo:
```sh
go run .
```

### On NixOS
A `shell.nix` is included to provide `libffi`, OpenGL, and Wayland/X11 libraries:
```sh
nix-shell --run "CGO_ENABLED=0 go run ."
```
Or enter the environment once:
```sh
nix-shell
CGO_ENABLED=0 go run .
```

### Controls
* **Left Click / Hold & Drag**: Move the target point.
* **Esc**: Quit.

---

## Testing & Benchmarks

The core math (`distance` and `moveToward`) lives in `logic.go` completely decoupled from Raylib, allowing unit tests and benchmarks to run headlessly without a display or GPU:

```sh
# Run unit tests
go test -v -tags test ./...

# Run benchmarks
go test -bench=. -tags test ./...
```

*Note: `-tags test` excludes the window loop (`main.go`), so tests run cleanly in headless CI and terminal environments.*

---

## Structure

* [`main.go`](main.go) — Window loop, input capture, and texture rendering.
* [`logic.go`](logic.go) — Vector math and frame smoothing step.
* [`logic_test.go`](logic_test.go) — Unit tests and benchmarks.
* [`shell.nix`](shell.nix) — Reproducible development shell for NixOS.
