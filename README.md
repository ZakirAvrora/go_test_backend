# go_test_backend

## Пререквизиты
Требуется `docker-compose`

## Запуск приложения
- Склонировать репозиторий
- Чтобы начать приложение в напишите в командой строке `docker-composer up -d`

# Service 1
- Запускается на порте :8081

Эндпоинты
- POST `/create-user`
- GET `/get-user/{email}`


# Service 2
- Запускается на порте :8000

Эндпоинты
- POST `/generate-salt`
