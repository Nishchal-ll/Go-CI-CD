# Go CI/CD Pipeline

[![Go CI](https://github.com/Nishchal-ll/Go-CI-CD/actions/workflows/ci.yml/badge.svg)](https://github.com/Nishchal-ll/Go-CI-CD/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A simple Go web server deployed to **Render** with **Continuous Integration (CI)** using **GitHub Actions**.

---

## 📁 Project Structure

```text
.
├── .github/
│   └── workflows/
│       └── ci.yml        # GitHub Actions CI workflow
├── go.mod                # Go module definition
├── main.go               # HTTP web server listening on PORT
├── main_test.go          # Unit tests
└── README.md             # Project documentation
```

---

## 🚀 Getting Started

### 1. Run Locally
```bash
go run main.go
```
Open [http://localhost:8080](http://localhost:8080) in your browser.

### 2. Run Tests
```bash
go test -v ./...
```

---

## 🌐 Endpoints

- `GET /` $\rightarrow$ Responds with `Hello, Nishchal!` (or `Hello, <name>!` via `/?name=<name>`).

---

## ⚙️ CI/CD Workflow

1. **GitHub Actions (CI)**: Runs unit tests and builds the binary on every push or PR.
2. **Render (CD)**: Automatically pulls the latest code on push, builds `./server`, and runs it live.

---

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
