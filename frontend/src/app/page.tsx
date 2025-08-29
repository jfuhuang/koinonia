'use client';

import { useEffect, useState } from 'react';
import Link from 'next/link';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { questAPI, leaderboardAPI, type Quest, type LeaderboardEntry, type User } from '@/lib/api';
import { Target, Trophy, Users, BookOpen, MapPin, Brain, Heart } from 'lucide-react';

export default function Home() {
  const [user, setUser] = useState<User | null>(null);
  const [recentQuests, setRecentQuests] = useState<Quest[]>([]);
  const [topUsers, setTopUsers] = useState<LeaderboardEntry[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // Check if user is logged in
    const userData = localStorage.getItem('user');
    if (userData) {
      setUser(JSON.parse(userData));
    }

    // Load data if user is authenticated
    if (userData) {
      loadDashboardData();
    } else {
      setLoading(false);
    }
  }, []);

  const loadDashboardData = async () => {
    try {
      // Load recent quests and top users
      const [quests, leaderboard] = await Promise.all([
        questAPI.getQuests(),
        leaderboardAPI.getLeaderboard(5)
      ]);

      setRecentQuests(quests.slice(0, 3)); // Show latest 3 quests
      setTopUsers(leaderboard);
    } catch (error) {
      console.error('Failed to load dashboard data:', error);
    } finally {
      setLoading(false);
    }
  };

  const getQuestIcon = (type: string) => {
    switch (type) {
      case 'scripture':
        return <BookOpen className="h-5 w-5" />;
      case 'side_quest':
        return <MapPin className="h-5 w-5" />;
      case 'trivia':
        return <Brain className="h-5 w-5" />;
      case 'encouragement':
        return <Heart className="h-5 w-5" />;
      default:
        return <Target className="h-5 w-5" />;
    }
  };

  const getQuestTypeLabel = (type: string) => {
    switch (type) {
      case 'scripture':
        return 'Scripture Memory';
      case 'side_quest':
        return 'Side Quest';
      case 'trivia':
        return 'Bible Trivia';
      case 'encouragement':
        return 'Encouragement';
      default:
        return type;
    }
  };

  if (!user) {
    // Landing page for non-authenticated users
    return (
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        {/* Hero Section */}
        <div className="text-center mb-16">
          <div className="flex justify-center mb-8">
            <BookOpen className="h-20 w-20 text-blue-600" />
          </div>
          <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-4">
            Welcome to Koinonia
          </h1>
          <p className="text-xl text-gray-600 dark:text-gray-300 mb-8 max-w-2xl mx-auto">
            A Christian fellowship and gamified community app promoting Scripture memory, 
            encouragement, and fellowship among students.
          </p>
          <div className="space-x-4">
            <Link href="/register">
              <Button size="lg">Get Started</Button>
            </Link>
            <Link href="/login">
              <Button variant="outline" size="lg">Sign In</Button>
            </Link>
          </div>
        </div>

        {/* Features Section */}
        <div className="grid md:grid-cols-3 gap-8 mb-16">
          <Card>
            <CardHeader className="text-center">
              <BookOpen className="h-12 w-12 text-blue-600 mx-auto mb-4" />
              <CardTitle>Scripture Memory</CardTitle>
              <CardDescription>
                Memorize Bible verses and earn points while growing in faith
              </CardDescription>
            </CardHeader>
          </Card>

          <Card>
            <CardHeader className="text-center">
              <Target className="h-12 w-12 text-green-600 mx-auto mb-4" />
              <CardTitle>Fun Quests</CardTitle>
              <CardDescription>
                Complete photo challenges, trivia, and encouragement activities
              </CardDescription>
            </CardHeader>
          </Card>

          <Card>
            <CardHeader className="text-center">
              <Trophy className="h-12 w-12 text-yellow-600 mx-auto mb-4" />
              <CardTitle>Leaderboards</CardTitle>
              <CardDescription>
                Compete with friends and track your spiritual growth journey
              </CardDescription>
            </CardHeader>
          </Card>
        </div>

        {/* Call to Action */}
        <div className="text-center bg-blue-50 dark:bg-blue-900/20 rounded-lg p-8">
          <h2 className="text-2xl font-bold text-gray-900 dark:text-white mb-4">
            Ready to Start Your Journey?
          </h2>
          <p className="text-gray-600 dark:text-gray-300 mb-6">
            Join our community of believers growing together in faith and fellowship.
          </p>
          <Link href="/register">
            <Button size="lg">
              Join Koinonia Today
            </Button>
          </Link>
        </div>
      </div>
    );
  }

  if (loading) {
    return (
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="text-center">
          <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-blue-600 mx-auto"></div>
          <p className="mt-4 text-gray-600 dark:text-gray-300">Loading your dashboard...</p>
        </div>
      </div>
    );
  }

  // Dashboard for authenticated users
  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      {/* Welcome Header */}
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900 dark:text-white">
          Welcome back, {user.first_name || user.username}!
        </h1>
        <p className="text-gray-600 dark:text-gray-300 mt-2">
          You have <span className="font-semibold text-blue-600">{user.total_points} points</span>. 
          Keep up the great work!
        </p>
      </div>

      <div className="grid lg:grid-cols-3 gap-8">
        {/* Recent Quests */}
        <div className="lg:col-span-2">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-xl font-semibold text-gray-900 dark:text-white">
              Latest Quests
            </h2>
            <Link href="/quests">
              <Button variant="outline" size="sm">View All</Button>
            </Link>
          </div>
          
          <div className="space-y-4">
            {recentQuests.map((quest) => (
              <Card key={quest.id}>
                <CardHeader className="pb-3">
                  <div className="flex items-start justify-between">
                    <div className="flex items-center space-x-2">
                      {getQuestIcon(quest.type)}
                      <div>
                        <CardTitle className="text-lg">{quest.title}</CardTitle>
                        <CardDescription>
                          {getQuestTypeLabel(quest.type)} • {quest.points} points
                        </CardDescription>
                      </div>
                    </div>
                    <span className={`px-2 py-1 rounded-full text-xs font-medium ${
                      quest.difficulty === 'easy' ? 'bg-green-100 text-green-800' :
                      quest.difficulty === 'medium' ? 'bg-yellow-100 text-yellow-800' :
                      'bg-red-100 text-red-800'
                    }`}>
                      {quest.difficulty}
                    </span>
                  </div>
                </CardHeader>
                <CardContent>
                  <p className="text-gray-600 dark:text-gray-300 mb-4">
                    {quest.description}
                  </p>
                  <Link href={`/quests/${quest.id}`}>
                    <Button size="sm">Start Quest</Button>
                  </Link>
                </CardContent>
              </Card>
            ))}
            
            {recentQuests.length === 0 && (
              <Card>
                <CardContent className="text-center py-8">
                  <Target className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                  <p className="text-gray-500">No quests available yet.</p>
                  <p className="text-sm text-gray-400 mt-1">Check back later for new challenges!</p>
                </CardContent>
              </Card>
            )}
          </div>
        </div>

        {/* Leaderboard Preview */}
        <div>
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-xl font-semibold text-gray-900 dark:text-white">
              Top Contributors
            </h2>
            <Link href="/leaderboard">
              <Button variant="outline" size="sm">View All</Button>
            </Link>
          </div>
          
          <Card>
            <CardContent className="p-0">
              {topUsers.map((entry, index) => (
                <div key={entry.user_id} className="flex items-center justify-between p-4 border-b border-gray-200 dark:border-gray-700 last:border-b-0">
                  <div className="flex items-center space-x-3">
                    <div className={`w-8 h-8 rounded-full flex items-center justify-center text-sm font-bold ${
                      index === 0 ? 'bg-yellow-100 text-yellow-800' :
                      index === 1 ? 'bg-gray-100 text-gray-800' :
                      index === 2 ? 'bg-orange-100 text-orange-800' :
                      'bg-blue-100 text-blue-800'
                    }`}>
                      {entry.rank}
                    </div>
                    <div>
                      <p className="font-medium text-gray-900 dark:text-white">
                        {entry.first_name} {entry.last_name}
                      </p>
                      <p className="text-sm text-gray-500">
                        @{entry.username}
                      </p>
                    </div>
                  </div>
                  <div className="text-right">
                    <p className="font-semibold text-gray-900 dark:text-white">
                      {entry.total_points}
                    </p>
                    <p className="text-sm text-gray-500">points</p>
                  </div>
                </div>
              ))}
              
              {topUsers.length === 0 && (
                <div className="text-center py-8">
                  <Users className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                  <p className="text-gray-500">No rankings yet.</p>
                  <p className="text-sm text-gray-400 mt-1">Be the first to complete a quest!</p>
                </div>
              )}
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
