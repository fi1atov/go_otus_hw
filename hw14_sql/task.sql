CREATE DATABASE test_db;

CREATE USER test_user WITH PASSWORD 'test';

CREATE SCHEMA IF NOT EXISTS test_schema;  -- Если схема не существует, она будет создана

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  password VARCHAR(50) NOT NULL
);

CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  order_date DATE NOT NULL,
  total_amount FLOAT NOT NULL,
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  price FLOAT NOT NULL
);

CREATE TABLE order_products (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL,
  product_id INT NOT NULL,
  CONSTRAINT fk_order FOREIGN KEY(order_id) REFERENCES orders(id),
  CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id)
);

-- Индексы на первичные ключи создались автоматически. Добавляю еще индексы на дополнительные поля

CREATE INDEX ix_user_name ON users (name);
CREATE INDEX ix_user_email ON users (email);

CREATE INDEX ix_product_name ON products (name);

CREATE INDEX ix_order_product_order_id ON order_products (order_id);
CREATE INDEX ix_order_product_product_id ON order_products (product_id);

-- вставка, редактирование и удаление пользователей
INSERT INTO users (name, email, password)
VALUES ('Петя', 'petya@mail.ru', '123');
INSERT INTO users (name, email, password)
VALUES ('Вася', 'vasya@mail.ru', '123');
INSERT INTO users (name, email, password)
VALUES ('Дима', 'dima@mail.ru', '123');

UPDATE users
SET email = 'petya1996@mail.ru'
WHERE name = 'Петя';

DELETE FROM users
WHERE name = 'Вася';

-- вставка, редактирование и удаление продуктов
INSERT INTO products (name, price)
VALUES ('Молоко', 40.0);
INSERT INTO products (name, price)
VALUES ('Хлеб', 20.0);
INSERT INTO products (name, price)
VALUES ('Сыр', 120.0);

UPDATE products
SET price = 45.0
WHERE name = 'Молоко';

DELETE FROM products
WHERE name = 'Сыр';

--сохранение и удаление заказов
INSERT INTO orders (user_id, order_date, total_amount)
VALUES (1, now(), 100);
INSERT INTO orders (user_id, order_date, total_amount)
VALUES (3, now(), 60);
INSERT INTO orders (user_id, order_date, total_amount)
VALUES (3, now(), 50);

insert into order_products (order_id, product_id)
values (1, 1);
insert into order_products (order_id, product_id)
values (1, 2);
insert into order_products (order_id, product_id)
values (2, 2);

DELETE FROM orders
WHERE id = 3;

-- выборка пользователей и выборка товаров
SELECT * FROM users;
SELECT * FROM products;

-- запрос на выборку заказов по пользователю
SELECT o.* 
from orders o
join users u ON o.user_id = u.id
WHERE u.name = 'Петя';

-- запрос на выборку статистики по пользователю (общая сумма заказов/средняя цена товара)
SELECT u.name, SUM(o.total_amount) AS total_amount, AVG(p.price) as avg_price
from orders o
join order_products op ON o.id = op.order_id
join products p ON op.product_id = p.id
join users u ON o.user_id = u.id
WHERE u.name = 'Петя'
group by u.name;
