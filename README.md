<div align="center">
  <img src="icon.png" alt="Happy Go Christmas Logo" width="180">

  # 🎄 Happy Go Christmas

  **A high-performance, atmospheric Christmas tree for your terminal.**
  <br>
  Written in Go, featuring customizable snow physics, embedded 8-bit music, and persistent settings.

  <br>

  ![Go Version](https://img.shields.io/github/go-mod/go-version/XPLassal/happy-go-christmas?style=for-the-badge&logo=go&color=00ADD8)
  ![AUR Version](https://img.shields.io/aur/version/happy-go-christmas?style=for-the-badge&logo=archlinux&color=1793d1)
  ![License](https://img.shields.io/github/license/XPLassal/happy-go-christmas?style=for-the-badge&color=blue)

  <br>
  <br>
  <img src="preview.gif" alt="Demo" width="600">
</div>

---

## ✨ Features

* 🎵 **Embedded Audio:** Includes an 8-bit "Jingle Bells" track packed directly into the binary using `embed`. No extra files needed!
* ❄️ **Customizable Physics:** Snowflakes drift with the wind, and you can control the **snow density** (from light flurry to heavy blizzard).
* ⚙️ **Smart JSON Config:** Automatically saves your preferences (Tree size, Music, Density, Leaf style) to `~/.config/happy-go-christmas/config.json`.
* 🐧 **Linux Integration:** Native `.desktop` support and application icon.
* 🚀 **High Performance:** Written in Go using efficient `strings.Builder` and VT100 escape codes for flicker-free rendering.

## 📦 Installation

### 🐧 Arch Linux / Manjaro / CachyOS (AUR)
The easiest way to install. Updates are handled by your package manager.

```bash
yay -S happy-go-christmas
# or
paru -S happy-go-christmas
```

### 📥 Binary Download (Windows, macOS, Linux)

Download the latest ready-to-run executable from the [Releases Page](https://www.google.com/search?q=https://github.com/XPLassal/happy-go-christmas/releases).

  * **Windows:** `.exe` comes with a custom pixel-art icon\!
  * **Linux/macOS:** Don't forget to run `chmod +x happy_new_year` after downloading.

### 🛠 Build from Source

Requirements: **Go 1.22+**

```bash
git clone [https://github.com/XPLassal/happy-go-christmas.git](https://github.com/XPLassal/happy-go-christmas.git)
cd happy-go-christmas

# Run directly
go run .

# Or build using the script
chmod +x build.sh
./build.sh
```

## 🎮 Usage & Configuration

Simply run the app in your terminal:

```bash
happy-go-christmas
```

On the first run, it will ask for your preferences interactively:

```text
Use Music? (y/n): y
Use Leaf ( ⣿ ⣷ ⣯ ) instead of stars( * )? (y/n): n
Write the size of the tree(more than 5): 15
Write the snow density in percent(1% < x < 100%): 20
```

**Options:**

  * `--edit`: Force reconfiguration mode to change settings.

<!-- end list -->

```bash
happy-go-christmas --edit
```

## 🏗️ Tech Stack

  * **Language:** Go (Golang)
  * **Audio Engine:** [gopxl/beep/v2](https://github.com/gopxl/beep)
  * **Assets:** `embed.FS` (Single binary distribution)
  * **Build System:** Custom Bash script with `GOEXPERIMENT=greenteagc`

## 📜 License

Distributed under the MIT License. See [LICENSE](https://github.com/XPLassal/happy-go-christmas/blob/main/LICENSE) for more information.

-----
