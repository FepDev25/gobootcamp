-- 4. Tabla de Encabezado de Ventas
CREATE TABLE IF NOT EXISTS sales (
    sale_id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(customer_id) ON DELETE SET NULL,
    total_amount DECIMAL(10, 2) NOT NULL DEFAULT 0,
    payment_method VARCHAR(50) DEFAULT 'CASH', -- CASH, CARD, QR, etc.
    sale_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 5. Tabla de Detalle de Venta (Qué productos se llevaron en esa venta)
CREATE TABLE IF NOT EXISTS sale_items (
    sale_item_id SERIAL PRIMARY KEY,
    sale_id INT NOT NULL REFERENCES sales(sale_id) ON DELETE CASCADE,
    product_id INT NOT NULL REFERENCES products(product_id),
    quantity INT NOT NULL CHECK (quantity > 0),
    unit_price DECIMAL(10, 2) NOT NULL, -- Se guarda el precio del momento de la venta
    subtotal DECIMAL(10, 2) GENERATED ALWAYS AS (quantity * unit_price) STORED
);
