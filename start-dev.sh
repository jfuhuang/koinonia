#!/bin/bash

# Koinonia Development Setup Script

echo "🚀 Starting Koinonia Development Environment"
echo "============================================"

# Check if Docker is available for database
if command -v docker &> /dev/null; then
    echo "📦 Starting PostgreSQL database with Docker..."
    cd backend
    docker-compose up postgres -d
    echo "✅ Database started on localhost:5432"
    echo ""
    echo "📋 Database Details:"
    echo "   Host: localhost"
    echo "   Port: 5432"
    echo "   Database: koinonia"
    echo "   Username: postgres"
    echo "   Password: password"
    echo ""
    cd ..
else
    echo "⚠️  Docker not found. Please install Docker or set up PostgreSQL manually."
    echo "   Database connection string: postgres://postgres:password@localhost:5432/koinonia"
    echo ""
fi

echo "🔧 To start the backend API server:"
echo "   cd backend && go run main.go"
echo ""

echo "🎨 To start the frontend development server:"
echo "   cd frontend && npm run dev"
echo ""

echo "🌐 Application URLs:"
echo "   Frontend: http://localhost:3000"
echo "   Backend API: http://localhost:8080"
echo "   Health Check: http://localhost:8080/health"
echo ""

echo "👤 Sample Login Credentials:"
echo "   Admin: admin / admin123"
echo "   User: john_doe / password123"
echo ""

echo "📖 See README.md for detailed setup instructions."
echo "Happy coding! ✝️"
