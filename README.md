# Cars Go API

Проект Cars Go API представляет собой RESTful API для управления данными об автомобилях. 

## Требования

- База данных PostgreSQL
- Файл конфигурации .env
- Внешний API сервис

## Установка

1. Склонируйте репозиторий.
2. Установите зависимости с помощью команды `go mod tidy`.
3. Настройте базу данных PostgreSQL.
4. Создайте в корне проекта .env файл и установите свои значения:

```
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_SSLMODE=
PORT=
LOG_LEVEL=
EXTERNAL_API_URL=
```

## Внешнее API

Проект взаимодействует с внешним API, описанным Swagger:

```yaml
openapi: 3.0.3
info:
  title: Car info
  version: 0.0.1
paths:
  /info:
    get:
      parameters:
        - name: regNum
          in: query
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Car'
        '400':
          description: Bad request
        '500':
          description: Internal server error
components:
  schemas:
    Car:
      required:
        - regNum
        - mark
        - model
        - owner
      type: object
      properties:
        regNum:
          type: string
          example: X123XX150
        mark:
          type: string
          example: Lada
        model:
          type: string
          example: Vesta
        year:
          type: integer
          example: 2002
        owner:
          $ref: '#/components/schemas/People'
    People:
      required:
        - name
        - surname
      type: object
      properties:
        name:
          type: string
        surname:
          type: string
        patronymic:
          type: string
```

## Конечные точки API

- `GET /cars`: Получить список всех автомобилей.
- `DELETE /cars/{id}`: Удалить автомобиль по ID.
- `PUT /cars/{id}`: Обновить информацию об автомобиле по ID.
- `POST /cars`: Добавить несколько автомобилей по их регистрационным номерам.
