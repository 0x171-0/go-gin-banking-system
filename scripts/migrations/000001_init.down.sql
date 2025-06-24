-- Drop indexes first
DROP INDEX IF EXISTS idx_cart_items_book;
DROP INDEX IF EXISTS idx_cart_items_user;
DROP INDEX IF EXISTS idx_order_items_book;
DROP INDEX IF EXISTS idx_order_items_order;
DROP INDEX IF EXISTS idx_orders_user;
DROP INDEX IF EXISTS idx_books_category;

-- Drop tables in reverse order of creation (to handle foreign key dependencies)
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS users;