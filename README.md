# Project Overview

## Purpose

`Zoneguard` is an IP Geolocator and basic Firewall system designed to identify and manage IP addresses based on their geolocation. It uses a database to map IP addresses to geographic locations and can restrict access based on these locations.

### Directory Structure

The directory structure of the project is as follows:

```bash
└── blackprince001-zoneguard/
    ├── go.sum
    ├── internal/
    │   ├── controllers/
    │   │   ├── ip.go
    │   │   └── root.go
    │   ├── grubber/
    │   │   └── grubber.go
    │   ├── database/
    │   │   ├── db.go
    │   │   ├── models.go
    │   │   └── csv_extractor.go
    │   └── config/
    │       └── config.go
    ├── cmd/
    │   ├── main.go
    │   └── routes.go
    ├── go.mod
    ├── LICENSE
    ├── README.md
    └── templates/
        ├── index.tmpl
        └── ip.tmpl
```

## Dependencies

The project uses several Go packages and dependencies listed in the `go.mod` file. Key dependencies include:

- `gorm.io/gorm` for database operations
- `gorm.io/driver/sqlite` for SQLite database driver
- `github.com/gin-gonic/gin` for web framework
- Various other packages for utility functions and support.

## Configuration

### Database Setup

The project uses a SQLite database to store geolocation data. Here’s how to set it up:

- Ensure the `data.csv` file is downloaded from the specified source in the `README.md`.
- The `CSVtoSqlite` function in `internal/database/csv_extractor.go` is used to import data from the CSV file into the SQLite database.
- The database is initialized in `cmd/main.go` using GORM.

### Environment and Dependencies

- Install Go and the necessary dependencies by running `go mod tidy` in the project directory.
- Ensure the `data.csv` file is in the root directory of the project.

## API Endpoints

The project exposes several API endpoints:

### Root Endpoint

- **Path:** `/`
- **Method:** GET
- **Description:** Displays the main page of the application.
- **Template:** `templates/index.tmpl`

### IP Address Endpoint

- **Path:** `/ip`
- **Method:** GET
- **Description:** Retrieves and displays the client's IP address.
- **Controller:** `controllers.GetIP`

### IP Geolocation Endpoint

- **Path:** `/geo`
- **Method:** GET
- **Description:** Retrieves and displays the geolocation data associated with the client's IP address.
- **Controller:** `controllers.GetIpGeoLocation`

### Address Endpoint

- **Path:** `/address`
- **Method:** GET
- **Description:** Displays the client's IP address in a HTML template.
- **Controller:** `controllers.Address`

## Controllers and Functions

### Controllers

- **ip.go:** Handles IP-related endpoints such as retrieving the client's IP and its geolocation.
- **root.go:** Handles the root endpoint of the application.
- **grubber.go:** Contains functions to extract the client's IP address from various headers.

### Database Functions

- **db.go:** Defines the repository interface and implementation for interacting with the `Iplocator` table.
- **models.go:** Defines the `Iplocator` struct used to represent geolocation data.
- **csv_extractor.go:** Contains functions to import data from a CSV file into the SQLite database.

## Usage

### Running the Application

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/your-repo/blackprince001-zoneguard.git
   ```

2. **Navigate to the Project Directory:**

   ```bash
   cd blackprince001-zoneguard
   ```

3. **Download the CSV File:**
   Download the `data.csv` file from the link provided in the `README.md` and place it in the project root directory.
4. **Install Dependencies:**

   ```bash
   go mod tidy
   ```

5. **Run the Application:**

   ```bash
   go run cmd/main.go
   ```

6. **Access the Application:**
   Open your web browser and navigate to `http://localhost:8000`.

### API Usage

- Use the endpoints as described above to interact with the application.

## Contributing and Development

### Directory Conventions

- **internal:** Contains packages internal to the project, such as controllers, database operations, and config.
- **cmd:** Contains the main application entry points.
- **templates:** Contains HTML templates for the web interface.
- **go.mod and go.sum:** Manage Go dependencies.

### Best Practices

- Follow standard Go coding practices and conventions.
- Ensure unit tests are written for new functionality.
- Use the `internal` directory for project-specific packages and the `pkg` directory for reusable packages if applicable.

## Licensing

The project is licensed under the MIT License. See the `LICENSE` file for details.

This documentation provides a comprehensive overview of the `Zoneguard` project, including its structure, dependencies, configuration, and usage. It serves as a guide for both users and contributors to the project.
