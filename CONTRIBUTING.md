# Contributing to Expense Tracker API

Thank you for your interest in contributing to the Expense Tracker API! We welcome contributions from the community.

## Getting Started

### Prerequisites

- **Go** (v1.20+ recommended)
- **PostgreSQL/MySQL/SQLite** (as the database backend)
- **Git** (for version control)

### Setup

1. **Fork the repository**
   - Fork this repository to your GitHub account
   - Clone your fork locally

2. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/expense-tracker-api.git
   cd expense-tracker-api
   ```

3. **Install dependencies:**
   ```bash
   go mod download
   ```

4. **Set up the database:**
   - Create a database for the project
   - Copy `.env.example` to `.env` and update the configuration
   ```bash
   cp .env.example .env
   ```
   - Update the `DATABASE_URL` and `JWT_SECRET` in `.env`

5. **Run database migrations:**
   - Execute the SQL files in the `migrations/` directory against your database

6. **Run the server:**
   ```bash
   go run cmd/server/main.go
   ```

## How to Contribute

### Reporting Issues

- Use the GitHub issue tracker to report bugs
- Include as much information as possible (Go version, OS, error messages, steps to reproduce)

### Submitting Changes

1. **Create a new branch** for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following the coding standards below

3. **Test your changes**:
   ```bash
   go build ./cmd/server/
   go test ./...
   ```

4. **Commit your changes** with a clear commit message:
   ```bash
   git commit -m "Add feature: description of what you added"
   ```

5. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```

6. **Create a Pull Request** from your fork to the main repository

### Coding Standards

- Follow Go's standard formatting (`go fmt`)
- Write clear, self-documenting code
- Add comments for complex logic
- Follow existing patterns in the codebase
- Ensure your code builds without warnings

### Environment Variables

Make sure to never commit sensitive information like:
- Database passwords
- JWT secrets
- API keys

Always use environment variables for configuration and update `.env.example` if you add new configuration options.

## Code Structure

```
├── cmd/server/          # Application entry point
├── config/              # Configuration management
├── internal/
│   ├── controllers/     # HTTP request handlers
│   ├── middleware/      # HTTP middleware
│   ├── repository/      # Data access layer
│   ├── routes/          # Route definitions
│   └── services/        # Business logic
├── migrations/          # Database migrations
└── pkg/
    ├── jwt/            # JWT utilities
    ├── models/         # Data models
    └── utils/          # Utility functions
```

## Security

- This project handles sensitive user data
- Always follow security best practices
- Never commit secrets or credentials
- Use parameterized queries to prevent SQL injection
- Validate all input data

## License

By contributing to this project, you agree that your contributions will be licensed under the MIT License.