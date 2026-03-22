# ght

> A fast, CLI tool for discovering trending GitHub repositories from your terminal.

`ght` is a lightweight command-line tool that brings GitHub's trending page straight to your terminal. Instead of opening a browser, developers can instantly see what's hot on GitHub.

---

## ✨ Features

- **Instant Results** — Fetch and display trending repos in seconds.
- **Language Filtering** — Narrow results to a specific programming language.
- **Time Range Support** — Browse daily, weekly, or monthly trending repos.
- **Cross-Platform Binaries** — Prebuilt releases for macOS and Linux.

---

## 📦 Installation

### Homebrew (macOS/Linux)

```bash
brew tap kwame-Owusu/ght
brew install ght
```

### Manual

Download a prebuilt binary from the [Releases page](https://github.com/kwame-Owusu/ght/releases) and move it to a directory on your `$PATH`.

---

## 🚀 Usage

After installation, run:

```bash
ght --help
```

Which will help you know the commands and flags for ght.

---

## 🛠 Development

### Prerequisites

- Go (latest stable version recommended)

### Clone the Repository

```bash
git clone https://github.com/kwame-Owusu/ght.git
cd ght
```

### Build

```bash
go build -o ght main.go
```

### Run

```bash
./ght
```

---

## 📦 Releases

Prebuilt binaries are available on the [GitHub Releases page](https://github.com/kwame-Owusu/ght/releases).

**Supported platforms:**

- macOS
  - `ght-darwin-amd64`
  - `ght-darwin-arm64`
- Linux
  - `ght-linux-amd64`
  - `ght-linux-arm64`

---
