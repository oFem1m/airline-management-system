# Система управления авиакомпанией
# API Документация

---

## Самолеты

### Добавление самолета
**Ubuntu/Linux:**
```bash
curl -X POST http://localhost:8080/api/v1/aircraft \
  -H "Content-Type: application/json" \
  -d '{
    "tail_number": "RA-12345",
    "model": "Boeing 737",
    "capacity": 180,
    "manufacture_year": 2015
  }'
```
**Windows:**
```bash
curl -X POST http://localhost:8080/api/v1/aircraft ^
  -H "Content-Type: application/json" ^
  -d "{\"tail_number\": \"RA-12345\", \"model\": \"Boeing 737\", \"capacity\": 180, \"manufacture_year\": 2015}"
```

### Получение списка самолетов
```bash
curl -X GET http://localhost:8080/api/v1/aircrafts
```

### Получение информации о конкретном самолете
```bash
curl -X GET http://localhost:8080/api/v1/aircraft/1
```

### Удаление самолета
```bash
curl -X DELETE http://localhost:8080/api/v1/aircraft/1
```

---

## Сотрудники

### Добавление сотрудника
**Ubuntu/Linux:**
```bash
curl -X POST http://localhost:8080/api/v1/employee \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "position": "Pilot",
    "hire_date": "2023-03-10T00:00:00Z",
    "salary": 100000,
    "email": "john.doe@example.com",
    "phone": "+12345678901"
  }'
```
**Windows:**
```bash
curl -X POST http://localhost:8080/api/v1/employee ^
  -H "Content-Type: application/json" ^
  -d "{\"first_name\": \"John\", \"last_name\": \"Doe\", \"position\": \"Pilot\", \"hire_date\": \"2023-03-10T00:00:00Z\", \"salary\": 100000, \"email\": \"john.doe@example.com\", \"phone\": \"+12345678901\"}"
```

### Получение списка сотрудников
```bash
curl http://localhost:8080/api/v1/employees
```

### Получение информации о конкретном сотруднике
```bash
curl http://localhost:8080/api/v1/employee/1
```

### Обновление данных сотрудника
**Ubuntu/Linux:**
```bash
curl -X PUT http://localhost:8080/api/v1/employee/1 \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "position": "Senior Pilot",
    "hire_date": "2022-12-01T00:00:00Z",
    "salary": 120000,
    "email": "john.doe@example.com",
    "phone": "+12345678901"
  }'
```
**Windows:**
```bash
curl -X PUT http://localhost:8080/api/v1/employee/1 ^
  -H "Content-Type: application/json" ^
  -d "{\"first_name\": \"John\", \"last_name\": \"Doe\", \"position\": \"Senior Pilot\", \"hire_date\": \"2022-12-01T00:00:00Z\", \"salary\": 120000, \"email\": \"john.doe@example.com\", \"phone\": \"+12345678901\"}"
```

### Удаление сотрудника
```bash
curl -X DELETE http://localhost:8080/api/v1/employee/1
```

---

## Аэропорты

### Добавление аэропорта
**Ubuntu/Linux:**
```bash
curl -X POST http://localhost:8080/api/v1/airport \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Sheremetyevo International Airport",
    "code": "SVO",
    "city": "Moscow",
    "country": "Russia",
    "timezone": "Europe/Moscow"
  }'
```
**Windows:**
```bash
curl -X POST http://localhost:8080/api/v1/airport ^ 
  -H "Content-Type: application/json" ^ 
  -d "{\"name\": \"Sheremetyevo International Airport\", \"code\": \"SVO\", \"city\": \"Moscow\", \"country\": \"Russia\", \"timezone\": \"Europe/Moscow\"}"
```

### Получение информации о конкретном аэропорте
```bash
curl -X GET http://localhost:8080/api/v1/airport/1
```

### Получение списка всех аэропортов
```bash
curl -X GET http://localhost:8080/api/v1/airports
```

### Удаление аэропорта
```bash
curl -X DELETE http://localhost:8080/api/v1/airport/1
```

---

## Маршруты

### Добавление маршрута
**Ubuntu/Linux:**
```bash
curl -X POST http://localhost:8080/api/v1/route \
  -H "Content-Type: application/json" \
  -d '{
    "departure_airport_id": 1,
    "arrival_airport_id": 2,
    "distance": 677,
    "duration_minutes": 94
  }'
```
**Windows:**
```bash
curl -X POST http://localhost:8080/api/v1/route ^ 
  -H "Content-Type: application/json" ^ 
  -d "{\"departure_airport_id\": 1, \"arrival_airport_id\": 2, \"distance\": 677, \"duration_minutes\": 94}"
```

### Получение списка всех маршрутов
```bash
curl -X GET http://localhost:8080/api/v1/routes
```

### Получение маршрута по ID
```bash
curl -X GET http://localhost:8080/api/v1/route/1
```

### Получение маршрутов для конкретного аэропорта
```bash
curl -X GET http://localhost:8080/api/v1/routes/airport/1
```

### Удаление маршрута
```bash
curl -X DELETE http://localhost:8080/api/v1/route/1
```

---
