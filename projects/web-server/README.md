# Go Mini Market API

Una API REST robusta y segura para la gestión de un mini mercado, construida con Go (Golang) y PostgreSQL. Este proyecto demuestra prácticas de ingeniería de software sólidas, incluyendo manejo de concurrencia, transacciones, precisión monetaria y arquitectura limpia.

## 🚀 Características Principales

*   **Arquitectura Limpia:** Estructura de proyecto estándar (`cmd/`, `internal/`, `migrations/`).
*   **CRUDs Completos:** Gestión de Categorías, Productos y Clientes.
*   **Gestión de Ventas Transaccional:**
    *   Integridad de datos garantizada con transacciones SQL (`BEGIN`, `COMMIT`, `ROLLBACK`).
    *   Actualización automática de stock.
*   **Manejo Seguro de Dinero:** Uso de la librería `shopspring/decimal` para evitar errores de redondeo de punto flotante.
*   **Prevención de Race Conditions:** Uso de `SELECT ... FOR UPDATE` para bloquear filas de inventario durante la venta, evitando ventas simultáneas que excedan el stock real.
*   **Configuración Segura:** Variables de entorno vía `.env`.
*   **Base de Datos Eficiente:** Uso de `pgx/v5` con pool de conexiones.

## 🛠️ Tech Stack

*   **Lenguaje:** Go 1.22+
*   **Base de Datos:** PostgreSQL
*   **Drivers & Librerías:**
    *   [`pgx/v5`](github.com/jackc/pgx): Driver de PostgreSQL de alto rendimiento.
    *   [`shopspring/decimal`](github.com/shopspring/decimal): Matemáticas decimales de precisión arbitraria.
    *   [`godotenv`](github.com/joho/godotenv): Carga de variables de entorno.
*   **Router:** `net/http` (Go 1.22 standard library router).

## 📂 Estructura del Proyecto

```
.
├── cmd
│   ├── api          # Punto de entrada del servidor (main.go)
│   └── seed         # Script para poblar la base de datos con datos de prueba
├── internal
│   ├── database     # Lógica de acceso a datos (Repositorio/Service)
│   ├── handlers     # Controladores HTTP
│   ├── models       # Estructuras de dominio
│   └── server       # Configuración del servidor HTTP y rutas
├── migrations       # Scripts SQL para crear el esquema
├── API_TESTING.md   # Guía rápida de comandos curl para probar la API
└── go.mod
```

## ⚙️ Configuración y Ejecución

### 1. Prerrequisitos
*   Go 1.22 o superior
*   PostgreSQL corriendo

### 2. Variables de Entorno
Crea un archivo `.env` en la raíz del proyecto:

```env
PORT=8080
APP_ENV=development
DB_CONN_STR=postgres://usuario:password@localhost:5432/web_go?sslmode=disable
```

### 3. Base de Datos
Crea la base de datos `web_go` y ejecuta los scripts de migración en orden:

```bash
psql -U postgres -d web_go -f migrations/001_create_categories.sql
psql -U postgres -d web_go -f migrations/002_create_products.sql
psql -U postgres -d web_go -f migrations/003_create_customers.sql
psql -U postgres -d web_go -f migrations/004_create_sales.sql
```

### 4. Poblar Base de Datos (Seed)
Ejecuta el script de seed para generar categorías, productos, clientes y ventas de prueba:

```bash
go run cmd/seed/main.go
```

### 5. Iniciar el Servidor

```bash
go run cmd/api/main.go
```

El servidor iniciará en `http://localhost:8080`.

## 🧪 Pruebas de API

Consulta el archivo [`API_TESTING.md`](./API_TESTING.md) para ver una lista completa de comandos `curl` listos para copiar y pegar, que cubren todos los endpoints del sistema.

## 📝 Notas de Implementación

### Precisión Monetaria
No usamos `float64` para precios. Usamos `decimal.Decimal` para garantizar que `$10.10 + $20.20` sea exactamente `$30.30` y no `$30.300000004`.

### Concurrencia en Inventario
El sistema es seguro para entornos de alta concurrencia. Si dos usuarios intentan comprar el último item al mismo tiempo, la base de datos serializará las peticiones a nivel de fila, asegurando que el stock nunca sea negativo.
