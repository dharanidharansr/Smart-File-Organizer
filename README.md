# 🗂️ Smart File Organizer

A modern desktop application that automatically organizes files in any folder into categorized subfolders — built with **Wails v2**, **Go**, **Astro**, and **TailwindCSS**.

![Smart File Organizer](https://img.shields.io/badge/version-1.0.0-white?style=flat-square&labelColor=000000)
![Go](https://img.shields.io/badge/Go-1.22-00ADD8?style=flat-square&logo=go&logoColor=white)
![Wails](https://img.shields.io/badge/Wails-v2-red?style=flat-square)
![Astro](https://img.shields.io/badge/Astro-4.x-FF5D01?style=flat-square&logo=astro&logoColor=white)
![Platform](https://img.shields.io/badge/platform-Windows-0078D6?style=flat-square&logo=windows&logoColor=white)

---

## ✨ Features

- 📁 **One-click folder organization** — select any folder and organize instantly
- 🗃️ **Auto-categorization** — sorts files into `Images`, `Videos`, `Documents`, `Music`, and `Others`
- 📊 **Live statistics** — real-time count of files organized per category
- 📜 **History log** — every file movement recorded with timestamps during the session
- 🎨 **Clean black & white UI** — minimal, distraction-free design
- ⚡ **Native desktop app** — no browser, no server, single `.exe`

---

## 📂 File Categories

| Category   | Extensions |
|------------|-----------|
| 🖼️ Images  | `.jpg` `.jpeg` `.png` `.gif` `.bmp` `.webp` `.svg` `.ico` `.tiff` |
| 🎬 Videos  | `.mp4` `.mkv` `.avi` `.mov` `.wmv` `.flv` `.webm` `.m4v` |
| 📄 Documents | `.pdf` `.doc` `.docx` `.xls` `.xlsx` `.ppt` `.pptx` `.txt` `.csv` `.md` |
| 🎵 Music   | `.mp3` `.wav` `.flac` `.aac` `.ogg` `.wma` `.m4a` |
| 📦 Others  | Everything else |

---

## 🖥️ Tech Stack

| Layer     | Technology |
|-----------|-----------|
| Desktop Framework | [Wails v2](https://wails.io/) |
| Backend   | Go 1.22 |
| Frontend  | [Astro 4](https://astro.build/) + [TailwindCSS 3](https://tailwindcss.com/) |
| IPC       | Wails Go bindings (`window.go.main.App.*`) |
| Runtime   | Windows WebView2 (built into Windows 10/11) |

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- [Node.js 18+](https://nodejs.org/)
- [Wails CLI v2](https://wails.io/docs/gettingstarted/installation)
- Windows 10/11 (WebView2 required — pre-installed on Win11, auto-installs on Win10)

### Install Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### Clone & Run

```bash
git clone https://github.com/YOUR_USERNAME/smart-file-organizer.git
cd smart-file-organizer

# Install frontend dependencies
cd frontend && npm install && cd ..

# Run in development mode (hot reload)
wails dev
```

### Build for Production

```bash
wails build
```

The output executable will be at `build/bin/SmartFileOrganizer.exe`.

---

## 📁 Project Structure

```
smart-file-organizer/
├── main.go            # Wails entry point
├── app.go             # App struct — all exposed Go methods
├── organizer.go       # File scanning & moving logic
├── models.go          # Data types (FileRecord, StatsResponse, etc.)
├── go.mod             # Go module file
├── wails.json         # Wails project config
└── frontend/
    ├── src/
    │   ├── components/
    │   │   ├── FolderSelector.astro   # Folder picker + organize button
    │   │   ├── StatsDashboard.astro   # Per-category file counts
    │   │   └── HistoryTable.astro     # File movement history log
    │   ├── layouts/
    │   │   └── Layout.astro           # Page shell, navbar, toasts
    │   ├── pages/
    │   │   └── index.astro            # Main page
    │   └── styles/
    │       └── global.css             # Global component styles
    ├── astro.config.mjs
    ├── tailwind.config.mjs
    └── package.json
```

---

## 📦 Download

> Pre-built binaries are available on the [Releases](https://github.com/YOUR_USERNAME/smart-file-organizer/releases) page.

No installation required — just download `SmartFileOrganizer.exe` and run it.

---

## 📝 License

MIT © 2026
