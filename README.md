# Video Structure Folder Automation

A simple command-line tool to automatically create a standardized folder structure for your video projects on Windows.  
This helps you keep your assets, project files, and exports organized with a single command.

---

## Features

- **Creates a main project folder** (with the current date in the name)
- **Adds subfolders** for project info, assets (with nested folders for footage, music, graphics, etc.), project files, and exports (drafts/final)
- **Customizable project name** via command-line flag
- **Easy installation and uninstallation on Windows**
- **No dependencies required** (except for [Go](https://golang.org/) if building from source)

---

## Folder Structure Example

```
My Project - 1 July 2025/
├── 00 Project Info/
├── 01 Assets/
│   ├── 01 FOOTAGE/
│   ├── 02 ASSETS/
│   ├── 03 SFX/
│   ├── 04 MUSIC/
│   ├── 05 VFX/
│   ├── 06 GRAPHICS/
│   └── 07 IMAGES/
├── 02 Project Files/
└── 03 Exports/
    ├── Drafts/
    └── Final/
```

---

## Usage

After installation, open a terminal or PowerShell and run:

```sh
video-structure -name "My Project"
```

A new folder will be created in your Videos directory with the specified structure.

---

## Installation

### Windows

```powershell
$script = "$env:TEMP\video_structure_install.ps1"; Invoke-WebRequest -Uri 'https://raw.githubusercontent.com/deluxesande/video-folder-structure-automation/main/install.ps1' -OutFile $script; & $script; Remove-Item $script
```

This script will:
- Download the latest release
- Extract the binary to `$env:USERPROFILE\video-structure`
- Add it to your PATH so you can run `video-structure` from anywhere

---

## Uninstallation

To completely remove `video-structure` and its PATH entry, run this one-liner in PowerShell:

```powershell
$installPath = "$env:USERPROFILE\video-structure"; $escapedPath = [regex]::Escape(";$installPath"); [Environment]::SetEnvironmentVariable("Path", ($env:Path -replace $escapedPath, ""), [EnvironmentVariableTarget]::User); Remove-Item -Recurse -Force $installPath; Write-Host "video-structure uninstalled successfully."
```

---

## Building from Source

1. [Install Go](https://golang.org/dl/)
2. Clone this repository:
    ```sh
    git clone https://github.com/deluxesande/video-folder-structure-automation.git
    cd video-folder-structure-automation
    ```
3. Build the executable:
    ```sh
    go build -o video-structure.exe
    ```

---

## License

MIT License

---

## Author

Deluxe Sande ([deluxesande](https://github.com/deluxesande))