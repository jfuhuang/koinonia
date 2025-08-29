-- Sample data for Koinonia database

-- Insert sample quests
INSERT INTO quests (title, description, type, points, difficulty, scripture_reference, scripture_text, is_active, created_at, updated_at) VALUES
('Memorize John 3:16', 'Learn the most famous verse about God''s love', 'scripture', 50, 'easy', 'John 3:16', 'For God so loved the world that he gave his one and only Son, that whoever believes in him shall not perish but have eternal life.', true, NOW(), NOW()),

('Memorize Philippians 4:13', 'A verse about strength through Christ', 'scripture', 50, 'easy', 'Philippians 4:13', 'I can do all this through him who gives me strength.', true, NOW(), NOW()),

('Campus Prayer Walk', 'Take a photo while praying at 3 different locations on campus', 'side_quest', 75, 'medium', '', '', true, NOW(), NOW()),

('Bible Trivia: Old Testament', 'Test your knowledge of Old Testament stories', 'trivia', 30, 'easy', '', '', true, NOW(), NOW()),

('Encourage a Friend', 'Send an encouraging message to someone and share a screenshot (blur names for privacy)', 'encouragement', 40, 'easy', '', '', true, NOW(), NOW()),

('Memorize Psalm 23', 'Learn the beloved shepherd psalm', 'scripture', 100, 'hard', 'Psalm 23', 'The Lord is my shepherd, I lack nothing. He makes me lie down in green pastures, he leads me beside quiet waters, he refreshes my soul. He guides me along the right paths for his name''s sake. Even though I walk through the darkest valley, I will fear no evil, for you are with me; your rod and your staff, they comfort me. You prepare a table before me in the presence of my enemies. You anoint my head with oil; my cup overflows. Surely your goodness and love will follow me all the days of my life, and I will dwell in the house of the Lord forever.', true, NOW(), NOW());

-- Create an admin user (password: admin123)
INSERT INTO users (username, email, password, first_name, last_name, role, total_points, is_active, created_at, updated_at) VALUES
('admin', 'admin@koinonia.app', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Admin', 'User', 'admin', 0, true, NOW(), NOW());

-- Create some sample users (password: password123 for all)
INSERT INTO users (username, email, password, first_name, last_name, role, total_points, is_active, created_at, updated_at) VALUES
('john_doe', 'john@example.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'John', 'Doe', 'user', 150, true, NOW(), NOW()),
('jane_smith', 'jane@example.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Jane', 'Smith', 'user', 200, true, NOW(), NOW()),
('mike_wilson', 'mike@example.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Mike', 'Wilson', 'user', 75, true, NOW(), NOW());
