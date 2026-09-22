# Go CI/CD Pipeline

[![Go CI](https://github.com/Nishchal-ll/Go-CI-CD/actions/workflows/ci.yml/badge.svg)](https://github.com/Nishchal-ll/Go-CI-CD/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A lightweight Go web server designed to demonstrate **Continuous Integration (CI)** and **Continuous Delivery (CD)** using **GitHub Actions**.

---

## 📁 Project Structure

```text
.
├── .github/
│   └── workflows/
│       └── ci.yml        # GitHub Actions CI pipeline configuration
├── go.mod                # Go module definition
├── main.go               # HTTP web server & handler logic
├── main_test.go          # Unit tests for endpoints and functions
└── README.md             # Project documentation
```

---

## 🚀 Getting Started

### Prerequisites
- [Go 1.22+](https://golang.org/dl/) installed.
- Git.

### 1. Run Locally
```bash
# Start the server
go run main.go
```
The server will start on `http://localhost:8080`.

### 2. Run Tests
```bash
# Run unit tests with verbose output and code coverage
go test -v -cover ./...
```

---

## 🌐 API Endpoints

| Method | Endpoint | Description | Example Response |
| :--- | :--- | :--- | :--- |
| `GET` | `/` | Returns greeting message | `Hello, World!` or `Hello, Alice!` (`/?name=Alice`) |
| `GET` | `/health` | Health check endpoint | `{"status": "ok"}` |

---

## ⚙️ CI/CD Pipeline (GitHub Actions)

On every `git push` or `pull request` to `main`, GitHub Actions automatically:
1. **Checks out** the repository code.
2. **Sets up** the Go runtime environment.
3. **Runs unit tests** with race detection and coverage checks (`go test -v -race -cover ./...`).
4. **Builds the binary** (`go build -v -o server .`) to ensure no compilation errors.

---

## 🚢 Deployment

To deploy this Go web service for free:
- **[Render](https://render.com)**: Connect this repository $\rightarrow$ Select `Go` runtime $\rightarrow$ Build: `go build -o server .` $\rightarrow$ Start: `./server`.
- **Docker / GHCR**: Containerize using Docker and publish images to GitHub Container Registry.

---

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
