
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );


INSERT INTO orders (customer_name, total_amount) VALUES
 ('Maria Eduarda', 99.99),
 ('Nilza Gomes', 49.50),
 ('Alice Silva', 150.00);
