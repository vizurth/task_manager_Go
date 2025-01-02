# Task Manager API

## Описание

Task Manager API — это RESTful API, которое предоставляет пользователям возможность управлять задачами. Вы можете выполнять CRUD-операции: создание, чтение, обновление и удаление задач. Также API поддерживает фильтрацию задач по статусу, времени создания или тегам.

## Функциональные возможности

- Создание задач.
- Получение всех задач или конкретной задачи по ID.
- Обновление статуса задачи.
- Удаление задачи.
- Фильтрация задач по:
  - Статусу (`status`).
  - Времени создания (`created_at`).
  - Тегам (`tag`).

## Структура JSON задачи

Каждая задача представлена в следующем формате:

```json
{
  "title": "some question",
  "id": 1,
  "tag": "work",
  "status": "done",
  "created_at": "2025-01-02T12:00:00Z"
}
