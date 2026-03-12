CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    level int NOT NULL DEFAULT 1,
    description TEXT
);

INSERT INTO roles (name, level, description) VALUES
('user', 1, 'Regular user role with limited access, can create posts and comments'),
('moderator', 2, 'Moderator role, moderates other users posts'),
('admin', 3, 'Administrator role with full access, can update and delete users posts');
