## API Endpoints

### Authentication
- `POST /api/v1/auth/register` - Register user baru
- `POST /api/v1/auth/login` - Login user
- `POST /api/v1/auth/logout` - Logout user

### Categories (Protected)
- `POST /api/v1/categories` - Buat kategori baru
- `GET /api/v1/categories` - Ambil semua kategori user
- `GET /api/v1/categories/:id` - Ambil kategori berdasarkan ID
- `PUT /api/v1/categories/:id` - Update kategori
- `DELETE /api/v1/categories/:id` - Hapus kategori

### Todos (Protected)
- `POST /api/v1/todos` - Buat todo baru
- `GET /api/v1/todos` - Ambil semua todo user (dengan filter & pagination)
- `GET /api/v1/todos/:id` - Ambil todo berdasarkan ID
- `PUT /api/v1/todos/:id` - Update todo
- `DELETE /api/v1/todos/:id` - Hapus todo
- `PATCH /api/v1/todos/:id/toggle` - Toggle status todo

### Query Parameters untuk GET /api/v1/todos
- `status` - Filter berdasarkan status (todo/done)
- `priority` - Filter berdasarkan prioritas (low/medium/high)
- `category_id` - Filter berdasarkan kategori
- `keyword` - Cari berdasarkan title atau description
- `page` - Halaman (default: 1)
- `limit` - Jumlah item per halaman (default: 10)

## Contoh Request

### Register
```json
POST /api/v1/auth/register
{
    "email": "user@example.com",
    "username": "username",
    "password": "password123"
}
```

### Login
```json
POST /api/v1/auth/login
{
    "email": "user@example.com",
    "password": "password123"
}
```

### Create Todo
```json
POST /api/v1/todos
Authorization: Bearer <jwt_token>
{
    "title": "Complete project",
    "description": "Finish the todo API project",
    "category_id": 1,
    "priority": "high",
    "deadline": "2024-12-31T23:59:59Z"
}
```

### Create Category
```json
POST /api/v1/categories
Authorization: Bearer <jwt_token>
{
    "name": "Work",
    "description": "Work related tasks",
    "color": "#FF5722"
}
```