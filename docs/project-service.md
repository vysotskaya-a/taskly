# Project Service API Documentation

## Base URL
`http://localhost:8090/api`

---

### 1. Создание проекта

**Endpoint:**  
`POST /projects`

**Описание:**  
Возвращает информацию о проекте по его id.

**Тело запроса:**
```json
{
  "title": "New Project",
  "description": "Project Description",
  "users": ["user1", "user2"]
}
```

**Ответ:**
```json
{
    "id": "12345"
}
```

---

### 2. Получение проекта по ID

**Endpoint:**  
`GET /projects/{project_id}`

**Описание:**  
Возвращает информацию о проекте по его id.

**Ответ:**
```json
{
  "id": "12345",
  "title": "Project Name",
  "description": "Project Description",
  "users": ["user1", "user2"],
  "admin_id": "admin123",
  "notification_subscribers_tg_ids": [123456789],
  "created_at": "2024-06-01T12:00:00Z"
}
```

---

### 3. Обновление проекта

**Endpoint:**  
`PATCH /projects/{project_id}`

**Описание:**  
Обновляет данные проекта.

**Тело запроса:**
```json
{
  "title": "Updated title",
  "description": "Updated description"
}
```

**Ответ:**
```json
{
  "message": "project updated successfully"
}
```

---

### 4. Добавление пользователя в проект

**Endpoint:**  
`POST /projects/{project_id}/add_user`

**Описание:**  
Добавляет пользователя в проект.

**Тело запроса:**
```json
{
  "user_id": "user123"
}
```

**Ответ:**
```json
{
  "message": "User added successfully"
}
```

---


### 5. Подписка на уведомления проекта

**Endpoint:**  
`POST /projects/{project_id}/subscribe_on_notifications`

**Описание:**  
Добавляет пользователя в проект.

**Тело запроса:**
```json
{
  "telegram_id": 123456789
}
```

**Ответ:**
```json
{
  "message": "Subscribed to project notifications"
}
```

---

### 6. Получение списка проектов пользователя

**Endpoint:**  
`GET /projects`

**Описание:**  
Возвращает список всех проектов, к которым имеет доступ текущий пользователь.

**Ответ:**
```json
{
  "projects": [
    {
      "id": "project123",
      "title": "Project Alpha",
      "description": "Description of Project Alpha",
      "users": ["user123", "user456"],
      "admin_id": "user123",
      "notification_subscribers_tg_ids": [123456789],
      "created_at": "2024-06-01T12:00:00Z"
    },
    {
      "id": "project456",
      "title": "Project Beta",
      "description": "Description of Project Beta",
      "users": ["user123"],
      "admin_id": "user123",
      "notification_subscribers_tg_ids": [],
      "created_at": "2024-06-05T12:00:00Z"
    }
  ]
}
```

---

### 6. Удаление проекта

**Endpoint:**  
`DELETE /projects/{project_id}`

**Описание:**  
Удаляет проект по его project_id.


---

### 8. Создание задачи

**Endpoint:**  
`POST /projects/{project_id}/tasks`

**Описание:**  
Создает задачу в проекте.

**Тело запроса:**
```json
{
  "title": "Task Title",
  "description": "Task Description",
  "status": "open",
  "executor": "user123",
  "deadline": "2024-06-10T12:00:00Z"
}
```

**Ответ:**
```json
{
  "id": "task123"
}
```

---

### 9. Получение задачи

**Endpoint:**  
`GET /projects/{project_id}/tasks/{task_id}`

**Описание:**  
Возвращает информацию о задаче.

**Ответ:**
```json
{
  "id": "task123",
  "title": "Task Title",
  "description": "Task Description",
  "status": "open",
  "project_id": "12345",
  "executor_id": "user123",
  "deadline": "2024-06-10T12:00:00Z",
  "created_at": "2024-06-01T12:00:00Z",
  "updated_at": "2024-06-05T12:00:00Z"
}
```

---

### 10. Получение списка задач по проекту

**Endpoint:**  
`GET /projects/{project_id}/tasks`

**Описание:**  
Возвращает список всех задач, связанных с проектом по его project_id.

**Ответ:**
```json
{
  "tasks": [
    {
      "id": "task123",
      "title": "Example Task 1",
      "description": "Description of the first task",
      "status": "In Progress",
      "project_id": "12345",
      "executor_id": "user123",
      "deadline": "2024-06-30T23:59:00Z",
      "created_at": "2024-06-01T12:00:00Z",
      "updated_at": "2024-06-02T12:00:00Z"
    },
    {
      "id": "task124",
      "title": "Example Task 2",
      "description": "Description of the second task",
      "status": "Completed",
      "project_id": "12345",
      "executor_id": "user456",
      "deadline": "2024-06-15T23:59:00Z",
      "created_at": "2024-06-03T12:00:00Z",
      "updated_at": "2024-06-05T12:00:00Z"
    }
  ]
}

```

---

### 11. Обновление задачи

**Endpoint:**  
`PATCH /projects/{project_id}/tasks/{task_id}`

**Описание:**  
Обновляет информацию о задаче.

**Тело запроса:**
```json
{
  "title": "Updated Task Title",
  "description": "Updated Description",
  "deadline": "2024-06-20T12:00:00Z"
}
```

**Ответ:**
```json
{
  "message": "Task updated successfully"
}
```

### 12. Удаление задачи

**Endpoint:**  
`DELETE /projects/{project_id}/tasks/{task_id}`

**Описание:**  
Удаляет задачу из проекта.

---
