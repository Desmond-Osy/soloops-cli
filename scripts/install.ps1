# PowerShell installation script for SoloOps CLI on Windows
# Downloads and installs the latest release

param(
    [string]$InstallDir = "$env:LOCALAPPDATA\soloops",
    [string]$Version = "latest"
)

$ErrorActionPreference = "Stop"

$Repo = "soloops/soloops-cli"
$BinaryName = "soloops.exe"

Write-Host ""
Write-Host "======================================" -ForegroundColor Cyan
Write-Host "  SoloOps CLI Installer" -ForegroundColor Cyan
Write-Host "======================================" -ForegroundColor Cyan
Write-Host ""

# Detect architecture
$Arch = "amd64"
if ([System.Environment]::Is64BitOperatingSystem) {
    $Arch = "amd64"
} else {
    Write-Host "Error: 32-bit Windows is not supported" -ForegroundColor Red
    exit 1
}

$Platform = "windows-$Arch"
Write-Host "Detected platform: $Platform" -ForegroundColor Green

# Get latest version if needed
if ($Version -eq "latest") {
    Write-Host "Fetching latest release..." -ForegroundColor Yellow

    try {
        $ReleaseInfo = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest"
        $Version = $ReleaseInfo.tag_name
        Write-Host "Latest version: $Version" -ForegroundColor Green
    } catch {
        Write-Host "Error: Could not fetch latest version" -ForegroundColor Red
        Write-Host $_.Exception.Message -ForegroundColor Red
        exit 1
    }
}

# Download URL
$BinaryUrl = "https://github.com/$Repo/releases/download/$Version/soloops-$Platform.zip"
$TmpDir = New-Item -ItemType Directory -Path "$env:TEMP\soloops-install-$(Get-Random)" -Force

Write-Host "Downloading from: $BinaryUrl" -ForegroundColor Yellow

try {
    $ZipPath = "$TmpDir\soloops.zip"
    Invoke-WebRequest -Uri $BinaryUrl -OutFile $ZipPath

    # Extract
    Write-Host "Extracting..." -ForegroundColor Yellow
    Expand-Archive -Path $ZipPath -DestinationPath $TmpDir -Force

    # Create install directory
    if (-not (Test-Path $InstallDir)) {
        New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
    }

    # Move binary
    Write-Host "Installing to $InstallDir..." -ForegroundColor Yellow
    $SourceBinary = Get-ChildItem -Path $TmpDir -Filter "soloops*.exe" | Select-Object -First 1
    Copy-Item -Path $SourceBinary.FullName -Destination "$InstallDir\$BinaryName" -Force

    # Cleanup
    Remove-Item -Path $TmpDir -Recurse -Force

    Write-Host ""
    Write-Host "Installation successful!" -ForegroundColor Green

} catch {
    Write-Host "Error during installation: $($_.Exception.Message)" -ForegroundColor Red
    Remove-Item -Path $TmpDir -Recurse -Force -ErrorAction SilentlyContinue
    exit 1
}

# Add to PATH
Write-Host ""
Write-Host "Adding to PATH..." -ForegroundColor Yellow

$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable(
        "Path",
        "$UserPath;$InstallDir",
        "User"
    )
    Write-Host "Added $InstallDir to PATH" -ForegroundColor Green
    Write-Host "Please restart your terminal for PATH changes to take effect" -ForegroundColor Yellow
} else {
    Write-Host "$InstallDir is already in PATH" -ForegroundColor Green
}

# Verify installation
Write-Host ""
Write-Host "Verifying installation..." -ForegroundColor Yellow

if (Test-Path "$InstallDir\$BinaryName") {
    Write-Host "Installation verified!" -ForegroundColor Green
    Write-Host ""
    Write-Host "======================================" -ForegroundColor Cyan
    Write-Host "  Installation Complete!" -ForegroundColor Cyan
    Write-Host "======================================" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "Installed to: $InstallDir\$BinaryName" -ForegroundColor Green
    Write-Host ""
    Write-Host "Get started:" -ForegroundColor Green
    Write-Host "  soloops init    # Initialize a new project" -ForegroundColor White
    Write-Host "  soloops --help  # Show all commands" -ForegroundColor White
    Write-Host ""
    Write-Host "Note: You may need to restart your terminal for the PATH changes to take effect" -ForegroundColor Yellow
    Write-Host ""
} else {
    Write-Host "Error: Installation verification failed" -ForegroundColor Red
    exit 1
}