-- Add category_id column to products table
ALTER TABLE products ADD COLUMN category_id INTEGER;

-- Add foreign key constraint to categories table
ALTER TABLE products 
ADD CONSTRAINT fk_products_category 
FOREIGN KEY (category_id) REFERENCES categories(id) 
ON DELETE SET NULL;

-- Add index for better query performance
CREATE INDEX idx_products_category_id ON products(category_id);
