# blob-physics

A quick 2D liquid blob prototype in Go using [raylib-go](https://github.com/gen2brain/raylib-go).

This was my initial physics sandbox before turning it into [Tez](https://github.com/vyrx-dev/tez). Big shoutout to [Tsoding](https://github.com/rexim) — I learned a ton from his streams, and this is rewritten and adapted in Go from his [original C blob experiment](https://gist.github.com/rexim/c1e3c1b99cb824ab6c817c3f45d75564).

## How it works

Two points chase the cursor:
1. **Head (`ball`)**: Chases the mouse cursor using a simple frame lerp (`moveToward`).
2. **Tail (`smear`)**: Chases the head at 45% speed to create lag.

A stretched half-circle dome is drawn between head and tail based on their distance and angle, giving it that fluid blob smear. (In Tez, I later upgraded this to framerate-independent exponential smoothing and angle wrapping).

## Running

Requires Go 1.24+. You don't need a C compiler or system Raylib installed:

```sh
CGO_ENABLED=0 go run .
```

On NixOS:
```sh
nix-shell --run "CGO_ENABLED=0 go run ."
```

* **Controls**: Click or hold left mouse and drag to move the blob. `Esc` to quit.

## Tests

The vector math and lerp logic in `logic.go` don't depend on Raylib, so you can test headlessly:

```sh
go test -v -tags test ./...
go test -bench=. -tags test ./...
```
