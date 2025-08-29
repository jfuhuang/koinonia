# Koinonia

**Fellowship through Faith and Fun** ✝️

Koinonia is a community-driven web app designed to promote encouragement, Scripture memory, and fellowship among students. Inspired by the Greek word *koinonia* (meaning fellowship, sharing, and unity), this project combines faith with a sense of adventure — through quests, challenges, and community engagement.

---

## 🌟 Vision
- Build stronger Christian community through fun and faith-centered activities.
- Encourage Scripture memorization as a spiritual discipline.
- Create engaging "side quests" around campus for fellowship and exploration.
- Provide a hub for encouragement, prayer, and accountability.

---

## 🚀 Features (MVP)
- ✅ User signup/login (JWT authentication)
- ✅ Quest system (Scripture memory, trivia, side quests, photo submissions)
- ✅ Admin approval flow for submissions
- ✅ Points system with history tracking
- ✅ Leaderboard for top users
- ✅ Modern UI with Next.js, TailwindCSS, and shadcn/ui
- ✅ Responsive design for mobile and desktop

---

## 🔮 Future Features
- Daily/weekly rotating trivia and quests
- Integration with Scripture Memory website
- Encouragement Wall (post anonymous notes of encouragement)
- Prayer Requests & Prayer tracking
- Badges & achievements (e.g. Scripture Master, Campus Explorer)
- Team quests & seasonal competitions
- PWA (Progressive Web App) for mobile use

---

## 🛠️ Tech Stack
- **Frontend**: Next.js 15 (React 19), TypeScript, TailwindCSS v4, shadcn/ui
- **Backend**: Go with Chi router, GORM, PostgreSQL
- **Database**: PostgreSQL
- **Auth**: JWT tokens with bcrypt password hashing
- **File Storage**: Local dev → S3/Supabase for production
- **Deployment**: Vercel (frontend), Render/Fly.io (backend)

---

## 📋 Project Setup

### Prerequisites
- Node.js 18+ and npm/yarn
- Go 1.21+
- PostgreSQL (local or Docker)
- Git

### Backend Setup (Go API)

1. **Navigate to backend directory:**
   ```bash
   cd backend
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables:**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. **Start PostgreSQL database:**
   ```bash
   # Using Docker Compose (recommended for development)
   docker-compose up postgres -d
   
   # OR install PostgreSQL locally and create database
   createdb koinonia
   ```

5. **Run the server:**
   ```bash
   go run main.go
   ```

The API server will start on `http://localhost:8080`

**Available API Endpoints:**
- `POST /api/auth/register` - User registration
- `POST /api/auth/login` - User login
- `GET /api/quests` - Get all quests
- `POST /api/quests/:id/submit` - Submit quest completion
- `GET /api/leaderboard` - Get leaderboard
- `GET /api/profile` - Get user profile
- Admin endpoints for quest management and submission approval

### Frontend Setup (Next.js)

1. **Navigate to frontend directory:**
   ```bash
   cd frontend
   ```

2. **Install dependencies:**
   ```bash
   npm install
   ```

3. **Set up environment variables:**
   ```bash
   # Create .env.local file
   echo "NEXT_PUBLIC_API_URL=http://localhost:8080/api" > .env.local
   ```

4. **Start the development server:**
   ```bash
   npm run dev
   ```

The frontend will start on `http://localhost:3000`

**Available Pages:**
- `/` - Home dashboard (landing page for guests, dashboard for authenticated users)
- `/login` - User login
- `/register` - User registration
- `/quests` - Quest listing (coming soon)
- `/leaderboard` - User rankings (coming soon)
- `/profile` - User profile (coming soon)

### Database Schema

The application uses the following main models:

**Users Table:**
- Authentication and profile information
- Points tracking
- Role-based access (user/admin)

**Quests Table:**
- Different quest types (scripture, side_quest, trivia, encouragement)
- Difficulty levels and point values
- Content specific to quest type

**Submissions Table:**
- User quest submissions
- Admin approval workflow
- Media attachments support

**Sample Data:**
The backend includes sample quests and users. Default admin credentials:
- Username: `admin`
- Password: `admin123`

---

## 🧪 Testing the Application

1. **Start both backend and frontend servers**
2. **Visit `http://localhost:3000`**
3. **Create a new account or use sample credentials:**
   - Sample users have password: `password123`
   - Admin user: `admin` / `admin123`

4. **Test key features:**
   - User registration and login
   - Browse available quests on dashboard
   - View leaderboard with sample data
   - Navigate between pages

---

## 🗂️ Project Structure

```
koinonia/
├── backend/                 # Go API server
│   ├── main.go             # Application entry point
│   ├── models/             # Database models (User, Quest, Submission)
│   ├── handlers/           # HTTP handlers (auth, quests, submissions, leaderboard)
│   ├── migrations/         # Database migrations and sample data
│   ├── docker-compose.yml  # PostgreSQL development setup
│   └── go.mod              # Go dependencies
├── frontend/               # Next.js React app
│   ├── src/
│   │   ├── app/           # Next.js app router pages
│   │   ├── components/    # Reusable UI components
│   │   └── lib/          # Utilities and API client
│   ├── package.json      # Node.js dependencies
│   └── tailwind.config.* # TailwindCSS configuration
└── README.md             # Project documentation
```

---

## 🎯 Development Roadmap

### Phase 1: MVP (Current)
- [x] Authentication system
- [x] Basic quest system
- [x] Points and leaderboard
- [x] Admin functionality
- [x] Responsive UI

### Phase 2: Enhanced Features
- [ ] Quest submission pages
- [ ] Profile management
- [ ] Quest filtering and search
- [ ] Real-time updates
- [ ] Mobile app (PWA)

### Phase 3: Community Features
- [ ] Team competitions
- [ ] Encouragement wall
- [ ] Prayer requests
- [ ] Badges and achievements
- [ ] Social features

---

## 🤝 Contributing

This project is open for collaboration with friends and fellow believers. The goal is to grow together — in fellowship and faith — through code and creativity.

### Getting Started with Contributions:
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

### Code Style:
- **Go**: Follow standard Go conventions with `gofmt`
- **TypeScript/React**: Use ESLint and Prettier
- **Database**: Use descriptive table and column names
- **Comments**: Explain complex business logic

---

## 📝 License

This project is open source and available under the MIT License.

---

## 🙏 Acknowledgments

Built with love for the Christian community, inspired by the desire to grow together in faith and fellowship.

> *"And let us consider how we may spur one another on toward love and good deeds, not giving up meeting together, as some are in the habit of doing, but encouraging one another—and all the more as you see the Day approaching."* - Hebrews 10:24-25

## ✅ To-Do List (MVP Roadmap)

### Setup
- [ ] Initialize Go backend (`/backend`)
- [ ] Initialize Next.js frontend (`/frontend`)
- [ ] Setup Postgres database (local Docker or Supabase/Neon)

---

### Backend (Go)
- [ ] Create `/hello` test endpoint
- [ ] Setup DB models: `users`, `quests`, `submissions`, `points_history`
- [ ] Implement `/auth/register` (store hashed password, return JWT)
- [ ] Implement `/auth/login` (validate credentials, return JWT)
- [ ] Implement `/quests` (list active quests)
- [ ] Implement `/quests/:id/submit` (store submission, mark as pending)
- [ ] Implement `/admin/submissions` (approve/reject submissions → add points)
- [ ] Implement `/leaderboard` (return users sorted by points)

---

### Frontend (Next.js)
- [ ] Build signup/login forms (connect to backend auth)
- [ ] Display current quests (fetched from backend)
- [ ] Build “submit quest” page (upload proof, send to backend)
- [ ] Build leaderboard page (fetch & display rankings)
- [ ] Add Navbar with user profile + points

---

### MVP Complete When:
- [ ] A user can sign up, log in, see quests, submit one, and appear on the leaderboard after admin approval


