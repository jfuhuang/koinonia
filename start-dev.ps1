# Koinonia Development Setup Script (PowerShell)

Write-Host "🚀 Starting Koinonia Development Environment" -ForegroundColor Cyan
Write-Host "============================================" -ForegroundColor Cyan
Write-Host ""

# Check if Docker is available for database
if (Get-Command docker -ErrorAction SilentlyContinue) {
    Write-Host "📦 Starting PostgreSQL database with Docker..." -ForegroundColor Yellow
    Set-Location backend
    docker-compose up postgres -d
    Write-Host "✅ Database started on localhost:5432" -ForegroundColor Green
    Write-Host ""
    Write-Host "📋 Database Details:" -ForegroundColor Blue
    Write-Host "   Host: localhost"
    Write-Host "   Port: 5432"
    Write-Host "   Database: koinonia"
    Write-Host "   Username: postgres"
    Write-Host "   Password: password"
    Write-Host ""
    Set-Location ..
} else {
    Write-Host "⚠️  Docker not found. Please install Docker or set up PostgreSQL manually." -ForegroundColor Yellow
    Write-Host "   Database connection string: postgres://postgres:password@localhost:5432/koinonia"
    Write-Host ""
}

Write-Host "🔧 To start the backend API server:" -ForegroundColor Magenta
Write-Host "   cd backend; go run main.go"
Write-Host ""

Write-Host "🎨 To start the frontend development server:" -ForegroundColor Magenta
Write-Host "   cd frontend; npm run dev"
Write-Host ""

Write-Host "🌐 Application URLs:" -ForegroundColor Blue
Write-Host "   Frontend: http://localhost:3000"
Write-Host "   Backend API: http://localhost:8080"
Write-Host "   Health Check: http://localhost:8080/health"
Write-Host ""

Write-Host "👤 Sample Login Credentials:" -ForegroundColor Blue
Write-Host "   Admin: admin / admin123"
Write-Host "   User: john_doe / password123"
Write-Host ""

Write-Host "📖 See README.md for detailed setup instructions." -ForegroundColor Green
Write-Host "Happy coding! ✝️" -ForegroundColor Green
