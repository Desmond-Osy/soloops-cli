# Contributing to SoloOps

Thank you for your interest in contributing to SoloOps! We welcome contributions from the community.

## Code of Conduct

This project adheres to a [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers.

## How to Contribute

### Reporting Bugs

Before creating a bug report:
- Check the [issue tracker](https://github.com/soloops/soloops-cli/issues) to see if the problem has already been reported
- If you're unable to find an open issue addressing the problem, open a new one

When creating a bug report, include:
- A clear and descriptive title
- Steps to reproduce the problem
- Expected behavior
- Actual behavior
- Your environment (OS, Go version, etc.)
- Any relevant logs or error messages

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, include:
- A clear and descriptive title
- A detailed description of the proposed functionality
- Examples of how the feature would be used
- Why this enhancement would be useful to most SoloOps users

### Pull Requests

1. **Fork the repository** and create your branch from `main`:
   ```bash
   git checkout -b feature/my-new-feature
   ```

2. **Make your changes**:
   - Write clear, commented code
   - Follow the existing code style
   - Add tests for new functionality
   - Update documentation as needed

3. **Test your changes**:
   ```bash
   make test
   make lint
   ```

4. **Commit your changes**:
   - Use clear and meaningful commit messages
   - Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification
   - Example: `feat: add support for GCP Cloud Run`

5. **Push to your fork** and submit a pull request:
   ```bash
   git push origin feature/my-new-feature
   ```

6. **Wait for review**:
   - Maintainers will review your PR
   - Address any feedback or requested changes
   - Once approved, your PR will be merged

### Pull Request Guidelines

- Keep PRs focused on a single feature or bug fix
- Include tests for new functionality
- Update documentation for user-facing changes
- Ensure all tests pass and code is linted
- Add a clear description of what the PR does
- Reference any related issues

## Development Setup

### Prerequisites

- Go 1.21 or later
- Make
- Git

### Setup

1. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/soloops-cli.git
   cd soloops-cli
   ```

2. Add upstream remote:
   ```bash
   git remote add upstream https://github.com/soloops/soloops-cli.git
   ```

3. Install dependencies:
   ```bash
   make deps
   ```

4. Build the project:
   ```bash
   make build
   ```

5. Run tests:
   ```bash
   make test
   ```

### Development Workflow

1. **Stay in sync with upstream**:
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```

2. **Create a feature branch**:
   ```bash
   git checkout -b feature/my-feature
   ```

3. **Make changes and test**:
   ```bash
   # Make your changes
   make test
   make lint
   ```

4. **Commit and push**:
   ```bash
   git add .
   git commit -m "feat: description of changes"
   git push origin feature/my-feature
   ```

## Code Style

- Follow standard Go conventions
- Use `gofmt` for formatting (run `make format`)
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions small and focused

## Testing

- Write unit tests for new functionality
- Ensure tests are deterministic and isolated
- Use table-driven tests where appropriate
- Mock external dependencies
- Aim for high test coverage (>80%)

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific test
go test -v -run TestConfigValidate ./tests
```

## Documentation

- Update README.md for user-facing changes
- Add inline comments for complex code
- Update command help text if adding/modifying commands
- Add examples for new features

## Project Structure

```
soloops-cli/
├── cmd/soloops/          # CLI entry point
├── pkg/
│   ├── cli/              # CLI commands
│   ├── config/           # Configuration parsing
│   ├── generator/        # Terraform generation
│   └── utils/            # Utility functions
├── infra-templates/      # Blueprint templates
├── tests/                # Test files
├── .github/workflows/    # CI/CD workflows
└── docs/                 # Documentation
```

## Adding New Features

### Adding a New Blueprint

1. Create template files in `infra-templates/`
2. Add blueprint type to `pkg/config/config.go`
3. Implement generation logic in `pkg/generator/`
4. Add tests in `tests/`
5. Update documentation

### Adding a New CLI Command

1. Create command file in `pkg/cli/`
2. Register command in `pkg/cli/root.go`
3. Implement command logic
4. Add tests
5. Update README with command documentation

## Release Process

Releases are managed by maintainers. The process includes:

1. Update version in relevant files
2. Update CHANGELOG.md
3. Create and push a git tag
4. GitHub Actions automatically builds and publishes binaries

## Getting Help

- Join our [Discussions](https://github.com/soloops/soloops-cli/discussions)
- Ask questions in issues
- Reach out to maintainers

## Recognition

Contributors will be recognized in:
- The project README
- Release notes
- The GitHub contributors page

Thank you for contributing to SoloOps!