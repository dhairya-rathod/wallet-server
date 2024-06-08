-- user
INSERT INTO users (name, email, password)
VALUES (
    'John Doe',
    'john.doe@example.com',
    '$2a$14$xCu/CaUKZ/Pwr7SBRUYIYOCEdYiteB41hVqnO7OhULE5HIfVkInta'
  ),
  (
    'Jane Smith',
    'jane.smith@example.com',
    '$2a$14$xCu/CaUKZ/Pwr7SBRUYIYOCEdYiteB41hVqnO7OhULE5HIfVkInta'
  ),
  (
    'Alice Johnson',
    'alice.johnson@example.com',
    '$2a$14$xCu/CaUKZ/Pwr7SBRUYIYOCEdYiteB41hVqnO7OhULE5HIfVkInta'
  );

-- category
INSERT INTO categories (name, is_active)
VALUES ('Groceries', TRUE),
  ('Utilities', TRUE),
  ('Entertainment', TRUE),
  ('Travel', TRUE),
  ('Healthcare', TRUE);

-- transaction
INSERT INTO transactions (amount, description, type, date, user_id, category_id)
VALUES
(50.75, 'Grocery shopping at Supermart', 'expense', '2024-06-01 10:00:00', 1, 1),
(120.00, 'Monthly electricity bill', 'expense', '2024-06-03 15:30:00', 2, 2),
(15.50, 'Movie tickets', 'expense', '2024-06-05 20:00:00', 3, 3),
(250.00, 'Flight tickets to New York', 'expense', '2024-06-07 08:00:00', 1, 4),
(30.00, 'Doctor consultation', 'expense', '2024-06-08 09:00:00', 2, 5),
(500.00, 'Salary received', 'income', '2024-06-10 09:00:00', 1, 4),
(100.00, 'Freelance project payment', 'income', '2024-06-11 10:00:00', 2, 4);
