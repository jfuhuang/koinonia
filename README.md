# Koinonia

**Fellowship through Faith and Fun** ✝️

Koinonia is a community-driven web app designed to promote encouragement, Scripture memory, and fellowship among students at Iowa State University (and beyond). Inspired by the Greek word *koinonia* (meaning fellowship, sharing, and unity), this project combines faith with a sense of adventure — through quests, challenges, and community engagement.

---

## 🌟 Vision
- Build stronger Christian community through fun and faith-centered activities.
- Encourage Scripture memorization as a spiritual discipline.
- Create engaging "side quests" around campus for fellowship and exploration.
- Provide a hub for encouragement, prayer, and accountability.

---

## 🚀 Features (MVP)
- User signup/login (JWT authentication)
- Quest system (Scripture memory, trivia, side quests, photo submissions)
- Admin approval flow for submissions
- Points system with history tracking
- Leaderboard for top users
- Basic UI with Next.js (quests list, submissions, leaderboard)

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
- **Frontend**: Next.js (React), TailwindCSS, shadcn/ui
- **Backend**: Go (Chi or Fiber framework)
- **Database**: PostgreSQL (via Supabase or Neon)
- **Auth**: JWT tokens
- **File Storage**: Local dev → S3/Supabase for production
- **Deployment**: Vercel (frontend), Render/Fly.io (backend)

---

## 📋 Project Setup

### Backend (Go)
1. Install Go (>=1.22).
2. Initialize project: `go mod init github.com/yourname/koinonia`
3. Add dependencies (Chi/Fiber, GORM, bcrypt, JWT).
4. Run server: `go run main.go`.

### Frontend (Next.js)
1. Create app: `npx create-next-app koinonia-frontend`
2. Add TailwindCSS + shadcn/ui.
3. Connect to backend via API routes.

### Database
- Use Docker for local Postgres OR sign up for Supabase/Neon (free Postgres hosting).
- Create tables: `users`, `quests`, `submissions`, `points_history`.

---

## 🗺️ Roadmap
See [koinonia_mvp_tasks.csv](./koinonia_mvp_tasks.csv) for the full MVP checklist.

---

## 🤝 Contributing
This project is open for collaboration with friends and fellow believers.  
The goal is to grow together — in fellowship and faith — through code and creativity.

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


