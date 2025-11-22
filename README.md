# Julia Set Visualizer

A simple Julia Set visualizer written in Go using the Ebiten game library.

This project renders Julia fractals in real-time and lets you interactively explore the complex plane using your mouse.

---

## Requirements

To run this project from source, you must have:

- **Go 1.22+**
- **Ebiten** (and its dependencies)

Ebiten installation instructions are available here:

🔗 https://ebitengine.org/en/documents/install.html

---

## Running the Project

Clone the repository:

```bash
git clone https://github.com/Abdul-Haseeb-0313/juliaset-visualizer.git
cd juliaset-visualizer
```


Run using Go:

```bash
make run
```

(Or simply go run ./cmd/main.go if you aren’t using Makefile.)

## Prebuilt Executables

A prebuilt executable for linux-amd64 is available in the [Releases](https://github.com/Abdul-Haseeb-0313/juliaset-visualizer/releases) section.

No Go installation or Ebiten setup required.

## Controls

Move your mouse to change the Julia set constant c.

The fractal updates in real time.

## Disclaimer

This project is provided for educational and personal use.  

While it should work on most systems, it comes with no guarantees.  

If something behaves unexpectedlyor if your laptop decides to overheat, crash, or explode that's on your hardware, not on me.  

Feel free to explore, modify, and improve the code however you like.


