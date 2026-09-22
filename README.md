# Go CI/CD Pipeline

[![Go CI](https://github.com/Nishchal-ll/Go-CI-CD/actions/workflows/ci.yml/badge.svg)](https://github.com/Nishchal-ll/Go-CI-CD/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A simple Go application demonstrating **Continuous Integration (CI)** using **GitHub Actions**.

---

## 📁 Project Structure

```text
.
├── .github/
│   └── workflows/
│       └── ci.yml        # GitHub Actions CI workflow
├── go.mod                # Go module file
├── main.go               # Simple Hello Nishchal program
├── main_test.go          # Unit tests
└── README.md             # Project documentation
```

---

## 🚀 Getting Started

### 1. Run Locally
```bash
go run main.go
```
**Output:**
```text
Hello, Nishchal!
```

### 2. Run Tests
```bash
go test -v ./...
```

---

## ⚙️ How the CI Pipeline Works

On every `git push` or `pull request` to `main`, GitHub Actions automatically:
1. **Checks out** the code onto an Ubuntu runner.
2. **Installs** Go.
3. **Runs unit tests** (`go test -v -race -cover ./...`).
4. **Builds the binary** (`go build -v -o app .`) to verify that the code compiles.

---

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
