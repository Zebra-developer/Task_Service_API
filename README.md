# Task Service API

REST API сервис для управления задачами с поддержкой повторяющихся задач (recurring tasks).

## Требования

- Go `1.23+`
- Docker и Docker Compose

## Быстрый запуск через Docker Compose

```bash
docker compose up --build
```

После запуска сервис будет доступен по адресу `http://localhost:8080`.

Если `postgres` уже запускался ранее со старой схемой, пересоздай volume:

```bash
docker compose down -v
docker compose up --build
```

Причина в том, что SQL-файл из `migrations/0001_create_tasks.up.sql` монтируется в `docker-entrypoint-initdb.d` и применяется только при инициализации пустого data volume.


## Функционал

- Создание задачи
- Получение списка задач
- Получение задачи по ID
- Обновление задачи
- Удаление задачи
- Поддержка recurring задач
- Валидация входных данных
- Swagger документация
- Docker запуск

---

## Поддерживаемые типы повторений

### `none`
Обычная задача без повторений

Пример:

```json
{
  "recurrence_type": "none",
  "recurrence_value": ""
}
```

---

### `daily`
Повтор каждые N дней

Пример:

```json
{
  "recurrence_type": "daily",
  "recurrence_value": "2"
}
```

Означает: повторять каждые 2 дня

---

### `monthly`
Повтор каждый месяц в определенный день

```json
{
  "recurrence_type": "monthly",
  "recurrence_value": "15"
}
```

Означает: каждый месяц 15 числа

---

### `even_days`
Повтор по четным дням

```json
{
  "recurrence_type": "even_days",
  "recurrence_value": ""
}
```

---

### `odd_days`
Повтор по нечетным дням

```json
{
  "recurrence_type": "odd_days",
  "recurrence_value": ""
}
```

---

### `specific_dates`
Повтор по конкретным датам

```json
{
  "recurrence_type": "specific_dates",
  "recurrence_value": "[\"2026-05-01\", \"2026-05-15\"]"
}
```

---

# Технологии

- Go
- PostgreSQL
- Docker
- Swagger/OpenAPI

---

# Запуск проекта

### 1. Клонировать репозиторий

```bash
git clone https://github.com/YOUR_USERNAME/YOUR_REPOSITORY.git
cd YOUR_REPOSITORY
```

---

### 2. Запустить проект

```bash
docker-compose up --build
```

---

# API endpoints

### Создать задачу

`POST /api/v1/tasks`

---

### Получить список задач

`GET /api/v1/tasks`

---

### Получить задачу по ID

`GET /api/v1/tasks/{id}`

---

### Обновить задачу

`PUT /api/v1/tasks/{id}`

---

### Удалить задачу

`DELETE /api/v1/tasks/{id}`

---

# Пример создания задачи

```json
{
  "title": "Call patients",
  "description": "Daily calls",
  "status": "new",
  "recurrence_type": "daily",
  "recurrence_value": "2"
}
```

---

# Валидация

API возвращает ошибку если:

- пустой title
- некорректный status
- неверный recurrence type
- monthly > 30
- daily <= 0

Пример ошибки:

```json
{
  "error": "invalid task input: invalid recurrence type"
}
```

---


## Swagger

Swagger UI:

```text
http://localhost:8080/swagger/
```

OpenAPI JSON:

```text
http://localhost:8080/swagger/openapi.json
```

## API

Базовый префикс API:

```text
/api/v1
```

Основные маршруты:

- `POST /api/v1/tasks`
- `GET /api/v1/tasks`
- `GET /api/v1/tasks/{id}`
- `PUT /api/v1/tasks/{id}`
- `DELETE /api/v1/tasks/{id}`


# Автор

Zebra_Developer 