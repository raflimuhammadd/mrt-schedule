# MRT Schedule API

A Go backend API for fetching Jakarta MRT station information and schedules. This service retrieves data from the official MRT Jakarta middleware API and provides a clean RESTful interface.

## Features

- **Stations Information**: Get comprehensive details about MRT stations including facilities, schedules, and integration with other transit systems
- **Environment Configuration**: Secure API URL management using `.env` files
- **Standardized Responses**: Consistent JSON response format for all endpoints
- **Modular Architecture**: Clean separation of concerns with service layer pattern
- **RESTful Design**: Following REST principles with proper HTTP status codes

## Tech Stack

- **Go 1.26.0** - Backend language
- **Gin Web Framework** - HTTP router and middleware
- **godotenv** - Environment variable management
- **Sonic** - High-performance JSON handling
- **MongoDB Driver** - Database connectivity (future use)

## Project Structure

```
mrt-schedule/
├── common/                    # Shared utilities
│   ├── client/               # HTTP client utilities
│   └── response/             # Standard API response format
├── modules/                  # Business logic modules
│   └── station/              # Station data module
│       ├── dto.go           # Data Transfer Objects
│       ├── router.go        # Route definitions
│       └── service.go       # Business logic
├── main.go                   # Application entry point
├── go.mod                    # Go module dependencies
├── .env                      # Environment variables (gitignored)
└── README.md                 # This file
```

## Getting Started

### Prerequisites

- Go 1.26.0 or later
- Git


1. **Clone the repository**
   ```bash
   git clone https://github.com/raflimuhammadd/mrt-schedule.git
   cd mrt-schedule
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

   The server will start at `http://localhost:8080`

## API Documentation

### Base URL
```
http://localhost:8080/v1/api
```

### Endpoints

#### Get All Stations
```http
GET /stations
```

**Response Example:**
```json
{
  "success": true,
  "message": "success",
  "data": [
    {
      "id": 6,
      "name": "Bundaran HI Bank Jakarta",
      "description": "Jalan Mohammad Husni Thamrin Kav. 17- 136, Gondangdia, Menteng, 10350, Jakarta Pusat, Indonesia"
    },
    // ... more stations
  ]
}
```

**Error Response:**
```json
{
  "success": false,
  "message": "API_BASE_URL is required",
  "data": null
}
```

### Response Format
All API responses follow this structure:
```json
{
  "success": boolean,    // Indicates if the request was successful
  "message": string,     // Human-readable message
  "data": any           // Response data (array, object, or null)
}
```


## Development

### Running in Development Mode
```bash
go run main.go
```

### Building for Production
```bash
go build -o mrt-schedule main.go
./mrt-schedule
```

### Checking Dependencies
```bash
go list -m all
```

### Code Structure

#### Service Layer Pattern
- **DTO (Data Transfer Object)**: Defines structs for API request/response
- **Service Interface**: Abstract business logic layer
- **Service Implementation**: Concrete implementation with API integration
- **Router**: HTTP route handlers with Gin

#### Adding New Modules
1. Create new module directory under `modules/`
2. Define DTOs, service interface, and implementation
3. Add router functions
4. Register routes in `main.go`

## Troubleshooting

### Common Issues

1. **"API_BASE_URL is required" error**
   - Ensure `.env` file exists in project root
   - Check that `API_BASE_URL` is properly set

2. **Connection refused**
   - Verify the external API is accessible
   - Check network connectivity

3. **Port already in use**
   - Change port in `main.go:30` (default is `:8080`)

### Logs
Application logs are printed to stdout. Check for:
- Environment variable loading status
- Server startup confirmation
- API request/response logs

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Future Enhancements

- [ ] Train schedule endpoints
- [ ] Real-time arrival predictions
- [ ] Station search functionality
- [ ] Multi-language support (English/Indonesian)
- [ ] Caching layer for API responses
- [ ] Rate limiting
- [ ] API documentation with Swagger/OpenAPI
- [ ] Docker containerization
- [ ] CI/CD pipeline
- [ ] Unit and integration tests

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Data provided by MRT Jakarta
- Built with Go and Gin framework
- Inspired by the need for accessible transit information