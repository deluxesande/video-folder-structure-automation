# Download and install video-structure on Windows
$version = "v1.0.0"
$os = "windows"
$arch = "amd64"
$url = "https://github.com/deluxesande/video-folder-structure-automation/releases/download/$version/video-structure.exe"
$installPath = "$env:USERPROFILE\video-structure"

# Create installation directory silently
try {
    New-Item -ItemType Directory -Force -Path $installPath | Out-Null
} catch {
    Write-Host "Failed to create directory: $installPath"
    exit 1
}

# Download the binary
Write-Host "Downloading video-structure..."
Invoke-WebRequest -Uri $url -OutFile "$installPath\video-structure.exe"

# Add to PATH
Write-Host "Adding video-structure to PATH..."
$env:Path += ";$installPath"
[Environment]::SetEnvironmentVariable("Path", $env:Path, [EnvironmentVariableTarget]::User)

Write-Host "video-structure installed successfully! You can now run 'video-structure' from anywhere."