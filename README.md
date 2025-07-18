# Golang Blog App 

A RESTful blog application built with Go, featuring user authentication and CRUD operations for blog posts.

## Technologies Used 

### Backend
- **Go** (Golang) - Core programming language
- **http** - Web framework/service
- **GORM** - ORM for database operations
- **PostgreSQL** - Relational database
- **JWT** - Authentication tokens

### Tools
- **godotenv** - Environment variable management
- **uuid** - Unique ID generation
- **Postman** - Testing data retreival for CRUD operations

## Setup Instructions

#### Prerequisites
- Go 1.20+
- PostgreSQL 16+ (PgAdmin4)
- Git

# Clone repository
git clone https://github.com/PanaAnt/Golang-Blog-App.git
cd Golang-Blog-App 

# Install dependencies
go mod download

# Database setup
Create a database in PgAdmin4 
Run this query for the database " CREATE EXTENSION IF NOT EXISTS "uuid-ossp" "

# .env setup
- Create a .env file to hide secrets such as database information needed for connection using your own credentials
- Make sure that the variables match e.g. DB_NAME = XXXXXXXX in the .env file, and in the database file its e.g. os.Getenv("DB_NAME")

# Deployment Instructions
- Start the server and run locally using 'go run main.go' and use Postman to test routes etc.

  OR

- Create a new app on a hosting platform such as Render and connect your Github repository. Next create a database and set up connections/details. Then set the environment variables.
