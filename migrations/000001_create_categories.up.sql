CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    type VARCHAR(20) NOT NULL CHECK (type IN ('income', 'expense')),
    icon VARCHAR(100) NOT NULL,
    color VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_category_name_type UNIQUE (name, type)
);
CREATE INDEX idx_categories_type ON categories(type);
INSERT INTO categories (name, type, icon, color)
VALUES ('Food', 'expense', 'restaurant', '#FF9800'),
    ('Travel', 'expense', 'directions_car', '#03A9F4'),
    ('Shopping', 'expense', 'shopping_bag', '#9C27B0'),
    ('Bills', 'expense', 'receipt_long', '#F44336'),
    ('Entertainment', 'expense', 'movie', '#673AB7'),
    (
        'Healthcare',
        'expense',
        'medical_services',
        '#4CAF50'
    ),
    ('Education', 'expense', 'college', '#3F51B5'),
    ('Other', 'expense', 'category', '#607D8B'),
    ('Salary', 'income', 'payments', '#4CAF50'),
    ('Freelancing', 'income', 'laptop', '#2196F3'),
    ('Bonus', 'income', 'card_giftcard', '#FFC107'),
    ('Interest', 'income', 'savings', '#009688'),
    ('Gift', 'income', 'redeem', '#E91E63'),
    ('Other', 'income', 'category', '#607D8B');