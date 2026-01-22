# Guía de Pruebas de API (API Testing Guide)

A continuación, encontrarás los comandos `curl` para probar todos los endpoints del sistema de Mini Mercado.

**Base URL:** `http://localhost:8080`

## 1. Sistema

### Health Check (Estado del servicio y DB)

```bash
curl -i http://localhost:8080/health
```

---

## 2. Categorías (Categories)

### Crear Categoría

```bash
curl -i -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Bebidas",
    "description": "Refrescos, jugos y aguas"
  }'
```

### Obtener Todas las Categorías

```bash
curl -i http://localhost:8080/categories
```

### Obtener Categoría por ID

```bash
curl -i http://localhost:8080/categories/1
```

### Actualizar Categoría

```bash
curl -i -X PUT http://localhost:8080/categories/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Bebidas & Licores",
    "description": "Refrescos, jugos, aguas y alcohol"
  }'
```

### Eliminar Categoría

```bash
curl -i -X DELETE http://localhost:8080/categories/1
```

---

## 3. Productos (Products)

> **Nota:** Para crear un producto, asegúrate de que el `category_id` exista.

### Crear Producto

```bash
curl -i -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "category_id": 1,
    "barcode": "7791234567890",
    "name": "Cola 1.5L",
    "price": 2.50,
    "stock": 100
  }'
```

### Obtener Todos los Productos

```bash
curl -i http://localhost:8080/products
```

### Obtener Producto por ID

```bash
curl -i http://localhost:8080/products/1
```

### Actualizar Producto

```bash
curl -i -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "category_id": 1,
    "barcode": "7791234567890",
    "name": "Cola 1.5L (Oferta)",
    "price": 2.00,
    "stock": 95,
    "is_active": true
  }'
```

### Eliminar Producto

```bash
curl -i -X DELETE http://localhost:8080/products/1
```

---

## 4. Clientes (Customers)

### Crear Cliente

```bash
curl -i -X POST http://localhost:8080/customers \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "Juan Pérez",
    "email": "juan@example.com",
    "phone": "+5491112345678",
    "address": "Av. Siempreviva 123"
  }'
```

### Obtener Todos los Clientes

```bash
curl -i http://localhost:8080/customers
```

### Obtener Cliente por ID

```bash
curl -i http://localhost:8080/customers/1
```

### Actualizar Cliente

```bash
curl -i -X PUT http://localhost:8080/customers/1 \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "Juan Pérez",
    "email": "juan.perez@email.com",
    "phone": "+5491187654321",
    "address": "Av. Siempreviva 742"
  }'
```

### Eliminar Cliente

```bash
curl -i -X DELETE http://localhost:8080/customers/1
```

---

## 5. Ventas (Sales)

> **Nota:** Para realizar una venta, asegúrate de tener creados los `customer_id` y `product_id` correspondientes, y que los productos tengan `stock` suficiente.

### Crear Venta (Transaccional)

Este endpoint:

1. Verifica stock (bloqueando fila para evitar race conditions).
2. Descuenta stock.
3. Genera la venta y sus detalles.
4. Calcula totales automáticamente.

```bash
curl -i -X POST http://localhost:8080/sales \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": 1,
    "payment_method": "EFECTIVO",
    "items": [
      {
        "product_id": 1,
        "quantity": 2
      }
    ]
  }'
```

### Obtener Venta por ID

Muestra la cabecera de la venta y el detalle de los items.

```bash
curl -i http://localhost:8080/sales/1
```
