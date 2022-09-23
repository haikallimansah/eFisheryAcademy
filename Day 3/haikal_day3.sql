-- buat database
CREATE DATABASE tugas_day3;

-- buat tabel
CREATE TABLE customers
(
    id            INT      NOT NULL PRIMARY KEY,
    customer_name CHAR(50) NOT NULL
);

CREATE TABLE products
(
    id   INT      NOT NULL PRIMARY KEY,
    product_name CHAR(50) NOT NULL
);

CREATE TABLE orders
(
    order_id    INT   NOT NULL PRIMARY KEY,
    customer_id INT   NOT NULL,
    product_id  INT   NOT NULL,
    order_date  DATE  NOT NULL,
    total       FLOAT NOT NULL,
    CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customers (id),
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (id)
);

-- insert
INSERT INTO customers (id, customer_name)
VALUES (1, 'Haikal'),
       (2, 'Persib'),
       (3, 'Bandung')
;

INSERT INTO products (id, product_name)
VALUES (1, 'Pulpen'),
       (2, 'Pensil'),
       (3, 'Penghapus')
;

INSERT INTO orders (order_id, customer_id, product_id, order_date, total)
VALUES (1, 1, 1, '1945-08-15', 100000),
       (2, 2, 2, '1945-08-16', 200000),
       (3, 3, 3, '1945-08-17', 300000)
;

-- update
UPDATE customers
SET customer_name = 'Ronaldo'
WHERE id = 1;

UPDATE products
SET name = 'Penggaris'
WHERE id = 2;

UPDATE orders
SET total = 999999
WHERE order_id = 3;

-- Delete Data
DELETE
FROM orders
WHERE order_id = 1;

DELETE
FROM customers
WHERE id = 1;

DELETE
FROM products
WHERE id = 1;


-- Show Data
SELECT id, customer_name
FROM customers;

SELECT id, product_name
FROM products;

SELECT order_id, customer_id, product_id, order_date, total
FROM orders;

-- Show Data (where)
SELECT customer_name
FROM customers
WHERE id = 1;

SELECT product_name
FROM products
WHERE id = 1;

SELECT order_date, total
FROM orders
WHERE order_id = 1
  AND customer_id = 1;

SELECT order_date, total
FROM orders
WHERE order_id = 1
   OR order_date = '1945-08-15';

-- Inner Join
SELECT c.customer_name AS "Customer Name",
       p.product_name  AS "Product Name",
       o.total         AS "Total Price",
       o.order_date    AS "Order Date"
FROM customers c
         INNER JOIN orders o ON c.id = o.customer_id
         INNER JOIN products p ON p.id = o.product_id
;

SELECT c.customer_name AS "Customer Name",
       p.name          AS "Product Name",
       o.total         AS "Total Price",
       o.order_date    AS "Order Date"
FROM customers c
         INNER JOIN orders o ON c.id = o.customer_id
         INNER JOIN products p ON p.id = o.product_id
WHERE order_date = '1945-08-16'
;

SELECT c.customer_name AS "Customer Name",
       p.name          AS "Product Name",
       o.total         AS "Total Price",
       o.order_date    AS "Order Date"
FROM customers c
         INNER JOIN orders o ON c.id = o.customer_id
         INNER JOIN products p ON p.id = o.product_id
ORDER BY o.total DESC
LIMIT 1;

-- Fungsi
SELECT MIN(total) AS "Sales terkecil"
FROM orders;

SELECT MAX(total) AS "Sales terbesar"
FROM orders;

SELECT AVG(total) AS "Rata-rata total sales"
FROM orders;

SELECT SUM(total) AS "Total Sales"
FROM orders;

SELECT COUNT(total) AS "Hitung order"
FROM orders;