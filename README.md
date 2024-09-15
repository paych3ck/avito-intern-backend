# avito-intern-backend
Тестовое задание на стажировку в Avito. Суть API - реализация системы тендеров.

## API
### Проверка работоспособности
- **URL**: `/api/ping`
- **Метод**: `GET`
- **Ответ**:
  ```json
    "ok"
  ```

### Получение списка тендеров
- **URL**: `/api/tenders`
- **Метод**: `GET`
- **Ответ**:
  ```json
  [
    {
        "id": "d0e9a8a0-26f4-4f8f-9dc1-848fa7bac9f1",
        "name": "NameTender1",
        "description": "DescriptionTender1",
        "status": "Created",
        "serviceType": "Delivery",
        "verstion": 1,
        "createdAt": "2000-01-01T11:01:01Z07:00"
    },
    {
        "id": "8b7167b4-488d-4ef8-ae5e-5dd86e9de60d",
        "name": "NameTender2",
        "description": "DescriptionTender2",
        "status": "Created",
        "serviceType": "Delivery",
        "verstion": 1,
        "createdAt": "2000-02-02T12:02:02Z07:00"
    }
  ]
  ```

### Создание тендера
- **URL**: `/api/tenders/new`
- **Метод**: `POST`
- **Тело запроса**:
  ```json
  {
    "name": "NewTenderName",
    "description": "NewTenderDescription",
    "serviceType": "Construction",
    "status": "Created",
    "organizationId": "5e982aad-619c-4bdf-a2f2-8a93ace132b9",
    "creatorUsername": "CreatorUsername"
  }
  ```
- **Ответ**:
  ```json
  {
    "id": "83a37566-07bb-4d4b-8916-3d5c5e85b96d",
    "name": "NewTenderName",
    "description": "NewTenderDescription",
    "serviceType": "Construction",
    "status": "Created",
    "organizationId": "5e982aad-619c-4bdf-a2f2-8a93ace132b9",
    "verstion": 1,
    "createdAt": "2024-01-02T15:15:15Z07:00"
  }
  ```

### Получение списка тендеров текущего пользователя
### Пример: GET /tenders/my?username=CreatorUsername&limit=10&offset=0
- **URL**: `/api/tenders/my`
- **Метод**: `GET`
- **Ответ**:
  ```json
  [
    {
      "id": "83a37566-07bb-4d4b-8916-3d5c5e85b96d",
      "name": "NewTenderName",
      "description": "NewTenderDescription",
      "serviceType": "Construction",
      "status": "Created",
      "organizationId": "5e982aad-619c-4bdf-a2f2-8a93ace132b9",
      "verstion": 1,
      "createdAt": "2024-01-02T15:15:15Z07:00"
    },
    {
      "id": "a0e9a8a0-36f4-5f8f-9dc1-848fa7bac9f2",
      "name": "AnotherTenderName",
      "description": "AnotherTenderDescription",
      "serviceType": "Construction",
      "status": "Created",
      "organizationId": "5e982aad-619c-4bdf-a2f2-8a93ace132b9",
      "verstion": 1,
      "createdAt": "2024-02-02T15:15:15Z07:00"
    }
  ```

### Получение статуса тендера
- **URL**: `/api/tenders/{tenderId}/status`
- **Метод**: `GET`
- **Ответ**:
  ```json
    "Created"
  ```

### Изменение статуса тендера
### Пример: ```PUT /api/tenders/83a37566-07bb-4d4b-8916-3d5c5e85b96d?status=Published&username=TestUser```
- **URL**: `/api/tenders/{tenderId}/status`
- **Метод**: `PUT`
- **Ответ**:
  ```json
  {
    "id": "83a37566-07bb-4d4b-8916-3d5c5e85b96d",
    "name": "NewTenderName",
    "description": "NewTenderDescription",
    "serviceType": "Construction",
    "status": "Published",
    "organizationId": "5e982aad-619c-4bdf-a2f2-8a93ace132b9",
    "verstion": 1,
    "createdAt": "2024-01-02T15:15:15Z07:00"
  }
  ```

### Изменение параметров тендера
### Пример: ```PATCH /api/tenders/83a37566-07bb-4d4b-8916-3d5c5e85b96d/edit?username=TestUser```
- **URL**: `/api/tenders/{tenderId}/edit`
- **Метод**: `PATCH`
- **Тело запроса**:
  ```json
  {
    "name": "UpdatedTenderName",
    "version": 2
  }
  ```
- **Ответ**:
  ```json
  {
    "id": "83a37566-07bb-4d4b-8916-3d5c5e85b96d",
    "name": "UpdatedTenderName",
    "description": "NewTenderDescription",
    "serviceType": "Construction",
    "status": "Created",
    "organizationId": "5e982aad-619c-4bdf-a2f2-8a93ace132b9",
    "verstion": 3,
    "createdAt": "2024-01-02T15:15:15Z07:00"
  }
  ```

### Откат версии тендера
### Пример: ```PUT /api/tenders/83a37566-07bb-4d4b-8916-3d5c5e85b96d/rollback/1?username=TestUser```
- **URL**: `/api/tenders/{tenderId}/rollback/{version}`
- **Метод**: `PUT`
- **Ответ**:
  ```json
  {
    "id": "83a37566-07bb-4d4b-8916-3d5c5e85b96d",
    "name": "UpdatedTenderName",
    "description": "NewTenderDescription",
    "serviceType": "Construction",
    "status": "Created",
    "organizationId": "5e982aad-619c-4bdf-a2f2-8a93ace132b9",
    "verstion": 2,
    "createdAt": "2024-01-02T15:15:15Z07:00"
  }
  ```

## Запуск проекта
1. Клонирование репозитория
```bash
git clone https://github.com/paych3ck/avito-intern-backend
```
2. Находясь в корневой директории
```bash
docker-compose up --build
```
3. Сервер запущен на ```http://localhost:8080```
