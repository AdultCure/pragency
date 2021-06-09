**docker pull postgres** - Скачивание образа postgres

**docker ps** - просмотреть контейнеры

**(1)docker run --name=pragency-db -e POSTGRES_PASSWORD=yellowcamel -p 5436:5432 -d --rm postgres** - создание и запуск контейнера postgres c флагом уничтожение после остановки
**(2)migrate -path ./schema -database postgres://postgres:yellowcamel@localhost:5436/postgres?sslmode=disable up** - миграция базы данных
**migrate -path ./schema -database postgres://postgres:yellowcamel@localhost:5436/postgres?sslmode=disable down** - остановка базы данных

**docker exec -it "номер контейнера" /bin/bash** - переход в базу данных через докер, требует номер контейнера

**psql -U postgres** - после этой команды можно работать с таблица бд. К примеру select \* from users;

**\d** - показывает структуру бд

**(3)go run cmd/main.go** - запуск приложения

**{
"name": "Artem",
"email": "artemeremenko.m@yandex.ru",
"password": "iloveddos"
}** - пример post запроса регистрации localhost:8000/auth/sign-up

**{
"email": "artemeremenko.m@yandex.ru",
"password": "iloveddos"
}** - пример post запроса авторизации и получения токена


// шаблоны

{
"name": "Denissssdf",
"email": "pupa@yandex.ruu",
"password": "qwersdfsdfty"
}

{
"name": "Denisssddsffsdfssdf",
"email": "pupa@yandex.ru",
"password": "qwerty"
}

{
"category": "12345678",
"status": "1234",
"date": "today",
"comment": "for u ass",
}