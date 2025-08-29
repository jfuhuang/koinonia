import axios from 'axios';

// Base API URL - adjust for your backend
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api';

// Create axios instance
export const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor to add auth token
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor to handle auth errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Token expired or invalid
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

// Types for API responses
export interface User {
  id: number;
  username: string;
  email: string;
  first_name: string;
  last_name: string;
  avatar: string;
  bio: string;
  total_points: number;
  role: string;
  is_active: boolean;
  last_login: string;
  created_at: string;
  updated_at: string;
}

export interface Quest {
  id: number;
  title: string;
  description: string;
  type: 'scripture' | 'side_quest' | 'trivia' | 'encouragement';
  points: number;
  difficulty: string;
  scripture_reference?: string;
  scripture_text?: string;
  trivia_question?: string;
  trivia_options?: string;
  is_active: boolean;
  start_date?: string;
  end_date?: string;
  max_submissions: number;
  created_at: string;
  updated_at: string;
}

export interface Submission {
  id: number;
  user_id: number;
  quest_id: number;
  content: string;
  media_url: string;
  media_type: string;
  status: 'pending' | 'approved' | 'rejected';
  points_awarded: number;
  admin_notes: string;
  reviewed_at?: string;
  reviewed_by_id?: number;
  created_at: string;
  updated_at: string;
  user?: User;
  quest?: Quest;
  reviewed_by?: User;
}

export interface LeaderboardEntry {
  rank: number;
  user_id: number;
  username: string;
  first_name: string;
  last_name: string;
  avatar: string;
  total_points: number;
  quests_completed: number;
}

export interface AuthResponse {
  token: string;
  user: User;
}

// Auth API functions
export const authAPI = {
  register: async (userData: {
    username: string;
    email: string;
    password: string;
    first_name: string;
    last_name: string;
  }): Promise<AuthResponse> => {
    const response = await api.post('/auth/register', userData);
    return response.data;
  },

  login: async (credentials: {
    username: string;
    password: string;
  }): Promise<AuthResponse> => {
    const response = await api.post('/auth/login', credentials);
    return response.data;
  },
};

// User API functions
export const userAPI = {
  getProfile: async (): Promise<User> => {
    const response = await api.get('/profile');
    return response.data;
  },

  updateProfile: async (userData: {
    first_name: string;
    last_name: string;
    bio: string;
    avatar: string;
  }): Promise<User> => {
    const response = await api.put('/profile', userData);
    return response.data;
  },
};

// Quest API functions
export const questAPI = {
  getQuests: async (filters?: { type?: string; difficulty?: string }): Promise<Quest[]> => {
    const params = new URLSearchParams();
    if (filters?.type) params.append('type', filters.type);
    if (filters?.difficulty) params.append('difficulty', filters.difficulty);
    
    const response = await api.get(`/quests?${params.toString()}`);
    return response.data;
  },

  getQuest: async (id: number): Promise<Quest> => {
    const response = await api.get(`/quests/${id}`);
    return response.data;
  },

  submitQuest: async (questId: number, submission: {
    content: string;
    media_url?: string;
    media_type?: string;
  }): Promise<Submission> => {
    const response = await api.post(`/quests/${questId}/submit`, submission);
    return response.data;
  },

  // Admin functions
  createQuest: async (questData: Partial<Quest>): Promise<Quest> => {
    const response = await api.post('/quests', questData);
    return response.data;
  },

  updateQuest: async (id: number, questData: Partial<Quest>): Promise<Quest> => {
    const response = await api.put(`/quests/${id}`, questData);
    return response.data;
  },

  deleteQuest: async (id: number): Promise<void> => {
    await api.delete(`/quests/${id}`);
  },
};

// Submission API functions
export const submissionAPI = {
  getSubmissions: async (filters?: {
    status?: string;
    quest_id?: number;
    user_id?: number;
  }): Promise<Submission[]> => {
    const params = new URLSearchParams();
    if (filters?.status) params.append('status', filters.status);
    if (filters?.quest_id) params.append('quest_id', filters.quest_id.toString());
    if (filters?.user_id) params.append('user_id', filters.user_id.toString());
    
    const response = await api.get(`/submissions?${params.toString()}`);
    return response.data;
  },

  approveSubmission: async (id: number): Promise<Submission> => {
    const response = await api.put(`/submissions/${id}/approve`);
    return response.data;
  },

  rejectSubmission: async (id: number, adminNotes?: string): Promise<Submission> => {
    const response = await api.put(`/submissions/${id}/reject`, { admin_notes: adminNotes });
    return response.data;
  },
};

// Leaderboard API functions
export const leaderboardAPI = {
  getLeaderboard: async (limit?: number): Promise<LeaderboardEntry[]> => {
    const params = limit ? `?limit=${limit}` : '';
    const response = await api.get(`/leaderboard${params}`);
    return response.data;
  },
};
