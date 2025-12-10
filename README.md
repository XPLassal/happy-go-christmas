-----

# 🎄 Happy Go Christmas

High-performance, atmospheric Christmas tree for your terminal, written in Go.
Features smooth animations, procedural snowfall, and embedded 8-bit music.

> **Note:** Best viewed in a terminal with TrueColor support (e.g., Alacritty, Kitty, Windows Terminal).

![Preview](preview.gif)

## ✨ Features

  * **🎵 Embedded Audio:** Uses `embed` standard library to pack 8-bit Jingle Bells directly into the binary. No external dependencies or asset folders required.
  * **❄️ Procedural Snowfall:** "Drifting" snow algorithm that renders snowflakes around the tree with randomized trajectories.
  * **🚀 Optimized Rendering:** Uses `strings.Builder`. Zero-allocation rendering loop where possible.
  * **⚙️ Persistent Config:** Automatically saves user preferences (Tree size, Music, Leaf style) to `os.UserConfigDir()` (`~/.config/happy_go_cristmas/` on Linux).
  * **🖥️ Cross-Platform:** Runs natively on Linux, Windows, and macOS (Intel & Apple Silicon).

## 📥 Installation

### Option 1: Build from Source

Requirements: **Go 1.22+** (due to `math/rand/v2` and `GOEXPERIMENT` features).

1.  Clone the repo:

    ```bash
    git clone https://github.com/XPLassal/happy-go-christmas.git
    cd happy-go-christmas
    ```

2.  Run directly:

    ```bash
    go run .
    ```

3.  Or build using the provided script (support for Windows/Linux/macOS):

    ```bash
    chmod +x build.sh
    ./build.sh
    ```

    *Artifacts will be placed in the `builds/` directory.*

### Option 2: Download Binary

Check the [Releases](https://www.google.com/search?q=https://github.com/XPLassal/happy-go-christmas/releases) page for the latest pre-compiled binaries.

## 🎮 Usage & Configuration

On the first run, the app will ask for your preferences interactively:

```text
Use Music? (y/n): y
Use Leaf ( ⣿", "⣷", "⣯", "⣟", "⣻", "⣽", "⣾" ) instead stars( * )? (y/n): n
Write the size of the tree(more than 5): 15
```

### CLI Arguments

To force a reconfiguration (edit mode), run:

```bash
./happy_new_year --edit
```

## 🛠️ Tech Stack

  * **Language:** Go (Golang)
  * **Audio:** [gopxl/beep/v2](https://github.com/gopxl/beep) (Sound processing)
  * **Build System:** Custom Bash script with `GOEXPERIMENT=greenteagc` for memory optimization.
  * **Assets:** `embed.FS` for single-binary distribution.

## 🤝 Contributing

Feel free to open issues or submit PRs if you want to add new snowflake types, music tracks, or optimize the rendering engine further\!

## 📜 License

MIT License. See [LICENSE](https://www.google.com/search?q=LICENSE) for details.
