# GitHub Actions Workflows

This document describes the GitHub Actions workflows configured for the Hardcover CLI project.

## Overview

The project uses several GitHub Actions workflows to ensure code quality, security, and automated deployment:

1. **CI (Continuous Integration)** - Main testing and quality checks
2. **Release** - Automated releases and binary builds
3. **Dependency Update** - Automated dependency management
4. **Nightly Tests** - Scheduled testing with latest dependencies

## Workflows

### 1. CI Workflow (`.github/workflows/ci.yml`)

**Triggers:**
- Push to `main` or `develop` branches
- Pull requests to `main` or `develop` branches

**Jobs:**

#### Test Job
- **Matrix Strategy**: Tests across multiple OS (Ubuntu, Windows, macOS) and Go versions (1.19, 1.20, 1.21)
- **Features**:
  - Go module caching for faster builds
  - Race condition detection
  - Code coverage reporting
  - Codecov integration

#### Lint Job
- **golangci-lint**: Comprehensive code linting
- **Configuration**: Uses `.golangci.yml` for custom rules
- **Features**: Performance, style, and security checks

#### Format Check Job
- **gofmt**: Ensures consistent code formatting
- **Fails if**: Any files are not properly formatted

#### Security Job
- **Gosec**: Security vulnerability scanning
- **SARIF Upload**: Results uploaded to GitHub Security tab

#### Build Job
- **Binary Build**: Ensures application builds successfully
- **CLI Testing**: Basic CLI functionality verification
- **Artifact Upload**: Stores built binary

#### Go Vet Job
- **Static Analysis**: Go's built-in static analysis tool
- **Detects**: Common Go programming errors

#### Mod Tidy Job
- **Dependency Check**: Ensures `go.mod` and `go.sum` are clean
- **Fails if**: `go mod tidy` would make changes

### 2. Release Workflow (`.github/workflows/release.yml`)

**Triggers:**
- Git tags matching `v*` pattern (e.g., `v1.0.0`)

**Jobs:**

#### Test Before Release
- **Full Test Suite**: Ensures all tests pass before release
- **Linting**: Code quality verification

#### Build Matrix
- **Cross-Platform**: Builds for multiple OS/architecture combinations
  - Linux (amd64, arm64)
  - macOS (amd64, arm64)
  - Windows (amd64)
- **Versioning**: Embeds version information in binaries
- **Archive Creation**: Creates `.tar.gz` (Unix) and `.zip` (Windows) archives

#### Release Creation
- **GitHub Release**: Automatically creates GitHub release
- **Asset Upload**: Attaches all platform binaries
- **Release Notes**: Includes installation and usage instructions

#### Homebrew (Optional)
- **Formula Update**: Automatically updates Homebrew formula
- **Requires**: `COMMITTER_TOKEN` secret and Homebrew tap repository

### 3. Dependency Update Workflow (`.github/workflows/dependency-update.yml`)

**Triggers:**
- Weekly schedule (Mondays at 9 AM UTC)
- Manual trigger (`workflow_dispatch`)

**Jobs:**

#### Dependency Updates
- **Automated Updates**: Updates all Go dependencies to latest versions
- **Testing**: Ensures tests still pass with new dependencies
- **Pull Request**: Creates PR with changes if updates are available

#### License Check
- **go-licenses**: Validates dependency licenses
- **Report Generation**: Creates license summary

### 4. Nightly Tests Workflow (`.github/workflows/nightly.yml`)

**Triggers:**
- Daily schedule (2 AM UTC)
- Manual trigger (`workflow_dispatch`)

**Jobs:**

#### Latest Dependencies Test
- **Bleeding Edge**: Tests with latest available dependencies
- **Coverage**: Generates coverage reports

#### Go Version Matrix
- **Compatibility**: Tests across multiple Go versions (1.19-1.22)
- **Future Proofing**: Ensures compatibility with newer Go releases

#### Performance Benchmarks
- **Benchmark Tests**: Runs performance benchmarks
- **Trend Tracking**: Monitors performance over time
- **Alerts**: Notifies on significant performance regressions

#### Integration Tests
- **CLI Testing**: End-to-end CLI command testing
- **Smoke Tests**: Basic functionality verification

#### Failure Notification
- **Issue Creation**: Automatically creates GitHub issues on failure
- **Comment Updates**: Updates existing issues with new failures

## Dependabot Configuration (`.github/dependabot.yml`)

**Features:**
- **Go Modules**: Weekly dependency updates
- **GitHub Actions**: Weekly action version updates
- **Pull Request Limits**: Prevents overwhelming number of PRs
- **Auto-Assignment**: Assigns PRs to maintainers
- **Labeling**: Automatic categorization

## Required Secrets

### For Basic Workflows
- `GITHUB_TOKEN` - Automatically provided by GitHub

### For Enhanced Features (Optional)
- `CODECOV_TOKEN` - For code coverage reporting
- `COMMITTER_TOKEN` - For Homebrew formula updates

## Configuration Files

### `.golangci.yml`
Comprehensive linting configuration with:
- **Enabled Linters**: 30+ linters for code quality
- **Custom Rules**: Project-specific configurations
- **Test Exclusions**: Relaxed rules for test files
- **Performance Focus**: Optimized for Go best practices

### `.github/dependabot.yml`
Automated dependency management:
- **Schedule**: Weekly updates on Mondays
- **Limits**: Reasonable PR limits to avoid spam
- **Categorization**: Proper labeling and assignment

## Usage

### Running Tests Locally

```bash
# Run all tests
go test ./...

# Run tests with race detection
go test -race ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...

# Run linter
golangci-lint run

# Check formatting
gofmt -s -l .
```

### Creating a Release

1. **Tag the Release**:
   ```bash
   git tag -a v1.0.0 -m "Release version 1.0.0"
   git push origin v1.0.0
   ```

2. **Automatic Process**:
   - Tests run automatically
   - Binaries built for all platforms
   - GitHub release created
   - Assets uploaded

### Manual Workflow Triggers

All workflows support manual triggering via GitHub UI:
1. Go to **Actions** tab
2. Select desired workflow
3. Click **Run workflow**
4. Choose branch and click **Run workflow**

## Best Practices

### For Contributors
1. **Test Locally**: Run tests before pushing
2. **Format Code**: Use `gofmt` or IDE formatting
3. **Check Linting**: Run `golangci-lint` locally
4. **Update Dependencies**: Keep dependencies current

### For Maintainers
1. **Review PRs**: Check automated PR reviews
2. **Monitor Nightly**: Check nightly test results
3. **Security**: Review security scan results
4. **Releases**: Use semantic versioning for tags

## Troubleshooting

### Common Issues

1. **Test Failures**: Check for dependency conflicts or API changes
2. **Lint Errors**: Review `.golangci.yml` configuration
3. **Build Failures**: Verify Go version compatibility
4. **Security Alerts**: Review and update vulnerable dependencies

### Workflow Debugging

1. **Enable Debug Logging**: Add `ACTIONS_STEP_DEBUG=true` secret
2. **Check Logs**: Review workflow run logs in Actions tab
3. **Local Reproduction**: Run commands locally to reproduce issues

## Monitoring

### Key Metrics
- **Test Coverage**: Monitor via Codecov dashboard
- **Security**: Review GitHub Security tab
- **Performance**: Track benchmark trends
- **Dependencies**: Monitor Dependabot PRs

### Alerts
- **Nightly Failures**: Automatic issue creation
- **Security Vulnerabilities**: GitHub Security alerts
- **Performance Regressions**: Benchmark alerts
- **Dependency Updates**: Dependabot notifications

This comprehensive CI/CD setup ensures high code quality, security, and automated maintenance for the Hardcover CLI project.