# Contributing to liftoff 🛠️

Thank you for considering contributing to **liftoff**! We welcome all improvements, big or small. This guide will help you get started.

## Table of Contents
- [Contributing to liftoff 🛠️](#contributing-to-liftoff-️)
  - [Table of Contents](#table-of-contents)
  - [Code of Conduct](#code-of-conduct)
  - [How to Contribute](#how-to-contribute)
    - [Reporting Issues](#reporting-issues)
    - [Feature Requests](#feature-requests)
    - [Submitting Pull Requests](#submitting-pull-requests)
  - [Development Setup](#development-setup)
  - [Style Guidelines](#style-guidelines)
  - [Testing](#testing)
  - [Commit Message Format](#commit-message-format)
  - [License](#license)

---

## Code of Conduct

All contributors are expected to follow our [Code of Conduct](CODE_OF_CONDUCT.md). Please read it to ensure a welcoming community for everyone.

---

## How to Contribute

### Reporting Issues

🪲 Found a bug? Please open an issue with:

- **Title**: concise summary of the problem
- **Description**: steps to reproduce, expected vs. actual behavior
- **Environment**: Go version, OS, `gcloud` version
- **Logs**: any error output or logs

### Feature Requests

✨ Have an idea? Open an issue and use the **Feature Request** template. Describe:

- **Use case**: what problem it solves
- **Proposal**: high-level approach or API

### Submitting Pull Requests

1. Fork the repo 🔀
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Commit changes to that branch 📝
4. Push to your fork: `git push origin feature/my-feature`
5. Open a PR against `main` 📬
6. Fill out the PR template and request reviews

---

## Development Setup

1. **Clone**:
```bash
git clone https://github.com/yourorg/liftoff.git
cd liftoff
```
2. **Install dependencies**:
```bash
go mod download
```
3. **Build**:
```bash
go build -o liftoff ./cmd/liftoff
   ```
4. **Run tests**:
```bash
go test ./...
   ```

---

## Style Guidelines

- Follow Go [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` / `goimports`
- Write clear, concise comments
- Emoji sprinkle: keep logs fun 🎉

---

## Testing

- Unit tests live alongside code in `_test.go` files.
- Run `go test ./...` to execute all tests.
- Ensure coverage for new features.

---

## Commit Message Format

Use [Conventional Commits](https://www.conventionalcommits.org/) style:
```
<type>(<scope>): <subject>

<body>

<footer>
``` 

Examples:
```
feat(config): add per-service defaults command
fix(gcloud): handle empty project from gcloud config
``` 

---

## License

By contributing, you agree that your contributions will be licensed under the MIT License.