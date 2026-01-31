-- Drop the index
DROP INDEX IF EXISTS idx_products_category_id;

-- Drop the foreign key constraint
ALTER TABLE products DROP CONSTRAINT IF EXISTS fk_products_category;

-- Remove category_id column from products table
ALTER TABLE products DROP COLUMN IF EXISTS category_id;
