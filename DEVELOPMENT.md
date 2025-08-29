# Koinonia Development Guide

## Quick Start

### 1. Start the Database
```bash
# Using Docker (recommended)
cd backend
docker-compose up postgres -d

# OR manually install PostgreSQL and create database 'koinonia'
```

### 2. Start the Backend API
```bash
cd backend
go mod tidy  # Install dependencies
go run main.go  # Start server on :8080
```

### 3. Start the Frontend
```bash
cd frontend
npm install  # Install dependencies (if not done already)
npm run dev  # Start Next.js on :3000
```

### 4. Test the Application
- Visit http://localhost:3000
- Register a new account or use sample credentials:
  - Admin: `admin` / `admin123`
  - User: `john_doe` / `password123`

## Development Workflow

### Adding New Features

1. **Backend Changes:**
   - Add routes in `main.go`
   - Create handlers in `handlers/` package
   - Add database models in `models/` package
   - Test endpoints with curl or Postman

2. **Frontend Changes:**
   - Add pages in `src/app/` directory
   - Create components in `src/components/`
   - Update API client in `src/lib/api.ts`
   - Test UI in browser

### Database Changes

1. Update models in `backend/models/models.go`
2. Run the application - GORM will auto-migrate
3. For production, create proper migration files

### Common Commands

```bash
# Backend
go mod tidy          # Update dependencies
go fmt ./...         # Format code
go run main.go       # Start server

# Frontend  
npm run dev          # Start development server
npm run build        # Build for production
npm run lint         # Run ESLint

# Database
docker-compose up postgres -d     # Start database
docker-compose down               # Stop services
docker-compose logs postgres      # View database logs
```

### Environment Variables

**Backend (.env):**
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=koinonia
PORT=8080
JWT_SECRET=your-secret-key
```

**Frontend (.env.local):**
```
NEXT_PUBLIC_API_URL=http://localhost:8080/api
```

## Troubleshooting

### Common Issues

1. **Database Connection Failed:**
   - Ensure PostgreSQL is running
   - Check connection string in `.env`
   - Verify database exists

2. **Frontend API Errors:**
   - Check backend server is running on :8080
   - Verify CORS settings
   - Check browser network tab for errors

3. **Authentication Issues:**
   - Clear localStorage and try again
   - Check JWT token expiration
   - Verify backend auth middleware

4. **Build Errors:**
   - Run `go mod tidy` for backend
   - Run `npm install` for frontend
   - Check for TypeScript errors

### Useful Tools

- **Database GUI:** pgAdmin, DBeaver, or TablePlus
- **API Testing:** Postman, Insomnia, or VS Code REST Client
- **Browser DevTools:** Network tab for API calls, Console for JS errors

## Project Structure Overview

```
backend/
├── main.go              # Application entry point & routing
├── handlers/            # HTTP request handlers
│   ├── auth.go         # Authentication endpoints
│   ├── middleware.go   # JWT middleware
│   ├── quests.go       # Quest management
│   ├── submissions.go  # Quest submissions
│   └── leaderboard.go  # User rankings
├── models/             # Database models
└── migrations/         # Sample data

frontend/
├── src/app/            # Next.js app router pages
├── src/components/     # Reusable UI components
├── src/lib/           # Utilities and API client
└── public/            # Static assets
```

## Contributing

1. Create feature branch: `git checkout -b feature/your-feature`
2. Make changes and test locally
3. Commit with descriptive message
4. Push and create pull request

Remember: This is a faith-based project - let's build something that encourages and strengthens our community! ✝️
