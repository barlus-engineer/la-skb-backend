# Laos School Knowledge Base (Back-End)

An open-source backend platform for a school knowledge base in Laos, licensed under the MIT License.

## Developer Info
- **Author:** Barlus Engineer
- **Contact:** barluscuda@gmail.com
- **GitHub:** [barlus-engineer](https://github.com/barlus-engineer)
- **Facebook:** Barlathcuda Lorfaichong

## Features

- User-friendly interface for accessing school-related information.
- Optimized for schools and educational institutions in Laos.
- Lightweight and efficient design for performance and accessibility.

## Technologies Used

- **Runtime:** [Go](https://go.dev/)
- **Framework/Library:** Gin, GORM
- **Database:** MariaDB

## Getting Started

Follow these steps to get the project up and running on your local environment.

### Prerequisites

Ensure you have the following installed on your system:

- [Go](https://go.dev/) (version 1.18 or higher recommended)
- [Git](https://git-scm.com/)
- [MariaDB](https://mariadb.org/) (version 10.5 or higher)

### Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/barlus-engineer/la-school-knowledge-base-backend.git
   cd la-school-knowledge-base-backend
   ```

2. **Set Up Environment Variables**:

   Create a `.env` file in the project root directory with the following contents:
   ```env
   IP=localhost
   PORT=3432
   DB_URI="[username]:[password]@tcp(127.0.0.1:3306)/[database]?charset=utf8mb4&parseTime=True&loc=Local"
   ```

3. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

4. **Build the Server**:
   ```bash
   go build cmd/server.go
   ```

5. **Start the Server**:
   ```bash
   cmd/server
   ```

<!-- ### Directory Structure

The repository uses a structured layout for clarity:
```
├── cmd/          # Main application entry points
├── config/       # Configuration settings
├── models/       # Database models
├── routes/       # API route definitions
├── services/     # Business logic
├── .env          # Environment variables (ignored by Git)
├── main.go       # Main application file
└── go.mod        # Go module file
```

### Testing

Run tests using:
```bash
go test ./...
```

### Contributing

Contributions are welcome! Feel free to fork the repository and submit a pull request. -->
