# Pushing SoloOps CLI to GitHub

Follow these steps to publish your repository to GitHub.

## Prerequisites

1. GitHub account
2. Git installed locally
3. GitHub CLI (optional but recommended)

## Option 1: Using GitHub CLI (Recommended)

### Step 1: Install GitHub CLI

```bash
# Windows
winget install --id GitHub.cli

# macOS
brew install gh

# Linux
sudo apt install gh
```

### Step 2: Authenticate

```bash
gh auth login
```

### Step 3: Create and Push Repository

```bash
cd /path/to/soloops-cli

# Initialize git (if not done)
git init

# Create repository on GitHub and push
gh repo create soloops-cli --public --source=. --remote=origin --push

# Or for private repository
gh repo create soloops-cli --private --source=. --remote=origin --push
```

## Option 2: Using Git + GitHub Web Interface

### Step 1: Create Repository on GitHub

1. Go to https://github.com/new
2. Repository name: `soloops-cli`
3. Description: "Infrastructure blueprint management CLI"
4. Choose Public or Private
5. **DO NOT** initialize with README, .gitignore, or license (we already have these)
6. Click "Create repository"

### Step 2: Initialize Git Locally

```bash
cd c:\Users\apipl\projects\soloops-cli

# Initialize git repository
git init

# Add all files
git add .

# Create initial commit
git commit -m "Initial commit: Complete SoloOps CLI implementation

- Implemented CLI with 7 commands (init, validate, generate, preview, apply, destroy, version)
- Added config package for YAML parsing and validation
- Added generator package for Terraform code generation
- Implemented AWS blueprints: Serverless API, Static Site, Budget Alerts
- Added comprehensive test suite
- Created Makefile and Dockerfile for builds
- Set up GitHub Actions CI/CD pipeline
- Added complete documentation (README, CONTRIBUTING, TESTING, QUICKSTART)
- Included example code and templates
- Licensed under Apache 2.0"
```

### Step 3: Add Remote and Push

```bash
# Add GitHub remote (replace YOUR_USERNAME with your GitHub username)
git remote add origin https://github.com/YOUR_USERNAME/soloops-cli.git

# Set main as default branch
git branch -M main

# Push to GitHub
git push -u origin main
```

## Option 3: Using SSH

### Step 1: Set up SSH Key

```bash
# Generate SSH key (if you don't have one)
ssh-keygen -t ed25519 -C "your_email@example.com"

# Start ssh-agent
eval "$(ssh-agent -s)"

# Add SSH key
ssh-add ~/.ssh/id_ed25519

# Copy public key
cat ~/.ssh/id_ed25519.pub
```

### Step 2: Add SSH Key to GitHub

1. Go to https://github.com/settings/keys
2. Click "New SSH key"
3. Paste your public key
4. Click "Add SSH key"

### Step 3: Push Repository

```bash
cd c:\Users\apipl\projects\soloops-cli

git init
git add .
git commit -m "Initial commit: Complete SoloOps CLI implementation"

# Add SSH remote (replace YOUR_USERNAME)
git remote add origin git@github.com:YOUR_USERNAME/soloops-cli.git

git branch -M main
git push -u origin main
```

## Post-Push Setup

### 1. Configure Repository Settings

Go to repository settings on GitHub:

**About Section:**
- Description: "Infrastructure blueprint management CLI - scaffold, validate, and deploy cloud resources"
- Website: (your docs site if any)
- Topics: `cli`, `infrastructure`, `terraform`, `aws`, `golang`, `iac`, `devops`

**General Settings:**
- ✅ Issues
- ✅ Discussions
- ✅ Projects
- ✅ Wiki (optional)

**Branches:**
- Set `main` as default branch
- Add branch protection rules (optional):
  - Require pull request reviews
  - Require status checks to pass
  - Require branches to be up to date

### 2. Configure Secrets for GitHub Actions

If you want to enable automated releases:

1. Go to Settings → Secrets and variables → Actions
2. Add secrets as needed (e.g., for Docker Hub):
   - `DOCKERHUB_USERNAME`
   - `DOCKERHUB_TOKEN`

### 3. Create First Release

```bash
# Tag the release
git tag -a v0.1.0 -m "Initial release: SoloOps CLI v0.1.0

Features:
- Complete CLI with 7 commands
- AWS blueprints (Serverless API, Static Site, Budget Alerts)
- Terraform code generation
- Multi-platform support
- Comprehensive documentation"

# Push the tag
git push origin v0.1.0
```

Or use GitHub CLI:
```bash
gh release create v0.1.0 --title "v0.1.0 - Initial Release" --notes "First release of SoloOps CLI"
```

### 4. Enable GitHub Pages (Optional)

For documentation:

1. Go to Settings → Pages
2. Source: Deploy from branch
3. Branch: `main`, folder: `/docs` (or create a `gh-pages` branch)
4. Save

### 5. Add Social Preview

1. Go to repository main page
2. Click Settings (repository settings, not profile)
3. Scroll to "Social preview"
4. Upload a preview image (1280x640px recommended)

## Repository Description Template

```
Infrastructure blueprint management CLI - scaffold, validate, and deploy cloud resources with best practices built-in.

🚀 Features:
- Declarative YAML configuration
- Terraform code generation
- AWS blueprints (Lambda, S3, CloudFront)
- Multi-environment support
- Budget monitoring
- Security best practices

📖 Quick Start: https://github.com/YOUR_USERNAME/soloops-cli#quick-start
📚 Docs: https://github.com/YOUR_USERNAME/soloops-cli/tree/main/infra-templates
```

## Verify Push

After pushing, verify everything is correct:

```bash
# Check remote
git remote -v

# Check branch
git branch -a

# Check status
git status

# View on GitHub
gh repo view --web
```

## Common Issues

### Authentication Failed

```bash
# Use personal access token
# Generate at: https://github.com/settings/tokens
# Use token as password when pushing
```

### Large Files

```bash
# Check file sizes
find . -type f -size +100M

# Use Git LFS for large files
git lfs track "*.zip"
git lfs track "*.tar.gz"
```

### Windows Line Endings

```bash
# Configure git to handle line endings
git config --global core.autocrlf true
```

## Next Steps

1. ✅ Repository pushed to GitHub
2. 📝 Add badges to README (build status, license, etc.)
3. 🏷️ Create first release (v0.1.0)
4. 📢 Announce on social media
5. 📦 Publish to package registries (optional)
6. 🌟 Star your own repository!
7. 👥 Invite collaborators (if any)

## Quick Reference

```bash
# Clone from GitHub
git clone https://github.com/YOUR_USERNAME/soloops-cli.git

# Pull latest changes
git pull origin main

# Create new branch
git checkout -b feature/new-feature

# Push branch
git push origin feature/new-feature

# Create PR
gh pr create --title "Add new feature" --body "Description"
```

## Support

If you encounter issues:
- GitHub Docs: https://docs.github.com
- Git Docs: https://git-scm.com/doc
- GitHub CLI: https://cli.github.com/manual/