# sports-junior-backend
 
Это тестовое задание, представляющее собой API для управления списком заметок с использованием стандартной библиотеки Go.

## Возможности

- **Создание заметки**
  - POST `/notes`
  - Входные данные: JSON с текстом заметки (content)
  - Ответ: JSON с ID созданной заметки

- **Получение всех заметок**
  - GET `/notes`
  - Ответ: JSON со списком всех заметок

- **Удаление заметки**
  - DELETE `/notes/{id}`
  - Ответ: JSON с подтверждением удаления

## API
### Создание заметки
- **URL**: `/notes`
- **Метод**: `POST`
- **Тело запроса**:
  ```json
  {
    "content": "Note with ID 0"
  }
  ```
- **Ответ**:
  ```json
  {
    "id": 0
  }
  ```

### Получение всех заметок
- **URL**: `/notes`
- **Метод**: `GET`
- **Ответ**:
  ```json
  {
    "notes": [
      {
        "id": 0,
        "content": "Note with ID 0"
      },
      {
        "id": 1,
        "content": "Note with ID 1"
      }
    ]
  }
  ```

### Удаление заметки
- **URL**: `/notes`
- **Метод**: `DELETE`
- **Ответ**:
  ```json
  {
    "status": "success"
  }
  ```

## Запуск проекта
1. Клонирование репозитория
```bash
git clone https://github.com/paych3ck/sports-junior-backend
```
2. Переход в нужную директорию
```bash
cd sports-junior-backend/cmd/sports-junior-backend
```
3. Запуск приложения
```bash
go run main.go
```
4. Сервер запущен на ```http://localhost:8080```

## Запуск тестов
Для запуска тестов необходимо выполнить последовательность команд
```bash
cd test
go test
```
или, находясь в корневой директории
```bash
go test ./...
```
