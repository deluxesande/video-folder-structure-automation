# Download and install video-structure on Windows
$version = "v1.0.0"
$url = "https://github.com/deluxesande/video-folder-structure-automation/releases/download/$version/video-structure_windows.zip"
$installPath = "$env:USERPROFILE\video-structure"

# Create installation directory silently
try {
    New-Item -ItemType Directory -Force -Path $installPath | Out-Null
} catch {
    Write-Host "Failed to create directory: $installPath"
    exit 1
}

# Download the zip
Write-Host "Downloading video-structure..."
Invoke-WebRequest -Uri $url -OutFile "$installPath\video-structure.zip"

# Extract the zip
Write-Host "Extracting video-structure..."
Expand-Archive -Path "$installPath\video-structure.zip" -DestinationPath $installPath -Force

# Optionally rename the binary if needed (adjust the name if your zip contains a different name)
if (Test-Path "$installPath\video-structure.exe") {
    # Already correct name
} elseif (Test-Path "$installPath\video-structure_windows_amd64.exe") {
    Rename-Item -Path "$installPath\video-structure_windows_amd64.exe" -NewName "video-structure.exe"
}

# Add to PATH
Write-Host "Adding video-structure to PATH..."
$env:Path += ";$installPath"
[Environment]::SetEnvironmentVariable("Path", $env:Path, [EnvironmentVariableTarget]::User)

# Clean up zip
Remove-Item "$installPath\video-structure.zip"

Write-Host "video-structure installed successfully! You can now run 'video-structure' from anywhere."