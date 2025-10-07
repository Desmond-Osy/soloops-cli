#!/bin/bash

# Installation script for SoloOps CLI
# This script downloads and installs the latest release

set -e

# Configuration
REPO="soloops/soloops-cli"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
BINARY_NAME="soloops"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Detect OS and architecture
detect_platform() {
    OS="$(uname -s)"
    ARCH="$(uname -m)"

    case "$OS" in
        Linux*)
            OS="linux"
            ;;
        Darwin*)
            OS="darwin"
            ;;
        *)
            echo -e "${RED}Unsupported OS: $OS${NC}"
            exit 1
            ;;
    esac

    case "$ARCH" in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        arm64|aarch64)
            ARCH="arm64"
            ;;
        *)
            echo -e "${RED}Unsupported architecture: $ARCH${NC}"
            exit 1
            ;;
    esac

    PLATFORM="${OS}-${ARCH}"
    echo -e "${GREEN}Detected platform: ${PLATFORM}${NC}"
}

# Get latest release version
get_latest_release() {
    echo -e "${YELLOW}Fetching latest release...${NC}"

    if command -v curl &> /dev/null; then
        VERSION=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    elif command -v wget &> /dev/null; then
        VERSION=$(wget -qO- "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    else
        echo -e "${RED}Error: Neither curl nor wget is available${NC}"
        exit 1
    fi

    if [ -z "$VERSION" ]; then
        echo -e "${RED}Error: Could not fetch latest version${NC}"
        exit 1
    fi

    echo -e "${GREEN}Latest version: ${VERSION}${NC}"
}

# Download binary
download_binary() {
    BINARY_URL="https://github.com/${REPO}/releases/download/${VERSION}/soloops-${PLATFORM}.tar.gz"
    TMP_DIR=$(mktemp -d)

    echo -e "${YELLOW}Downloading from: ${BINARY_URL}${NC}"

    if command -v curl &> /dev/null; then
        curl -L "${BINARY_URL}" -o "${TMP_DIR}/soloops.tar.gz"
    elif command -v wget &> /dev/null; then
        wget -O "${TMP_DIR}/soloops.tar.gz" "${BINARY_URL}"
    fi

    if [ ! -f "${TMP_DIR}/soloops.tar.gz" ]; then
        echo -e "${RED}Error: Failed to download binary${NC}"
        exit 1
    fi

    # Extract
    echo -e "${YELLOW}Extracting...${NC}"
    tar -xzf "${TMP_DIR}/soloops.tar.gz" -C "${TMP_DIR}"

    BINARY_PATH="${TMP_DIR}/soloops-${PLATFORM}"
}

# Install binary
install_binary() {
    echo -e "${YELLOW}Installing to ${INSTALL_DIR}...${NC}"

    # Check if we need sudo
    if [ -w "$INSTALL_DIR" ]; then
        mv "${BINARY_PATH}" "${INSTALL_DIR}/${BINARY_NAME}"
        chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
    else
        echo -e "${YELLOW}Requesting sudo permission to install to ${INSTALL_DIR}${NC}"
        sudo mv "${BINARY_PATH}" "${INSTALL_DIR}/${BINARY_NAME}"
        sudo chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
    fi

    # Cleanup
    rm -rf "${TMP_DIR}"

    echo -e "${GREEN}✓ SoloOps CLI installed successfully!${NC}"
}

# Verify installation
verify_installation() {
    if command -v ${BINARY_NAME} &> /dev/null; then
        echo -e "${GREEN}✓ Installation verified${NC}"
        echo ""
        ${BINARY_NAME} version
        echo ""
        echo -e "${GREEN}Get started:${NC}"
        echo "  soloops init    # Initialize a new project"
        echo "  soloops --help  # Show all commands"
    else
        echo -e "${YELLOW}Warning: ${BINARY_NAME} not found in PATH${NC}"
        echo "You may need to add ${INSTALL_DIR} to your PATH"
        echo "Add this to your ~/.bashrc or ~/.zshrc:"
        echo "  export PATH=\$PATH:${INSTALL_DIR}"
    fi
}

# Main installation flow
main() {
    echo ""
    echo "======================================"
    echo "  SoloOps CLI Installer"
    echo "======================================"
    echo ""

    detect_platform
    get_latest_release
    download_binary
    install_binary
    verify_installation

    echo ""
    echo "======================================"
    echo "  Installation Complete!"
    echo "======================================"
    echo ""
}

main