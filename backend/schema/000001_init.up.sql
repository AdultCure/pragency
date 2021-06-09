-- Файл структуры миграции бд

-- Таблица пользователей
CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null,
    admin         varchar(255)
);

-- Таблица заказов
CREATE TABLE orders
(
    id        serial       not null unique,
    category  varchar(255) not null,
    status    varchar(255) not null,
    date      varchar(255) not null,
    comment   varchar(2000) not null,
    user_id   varchar(255) not null,
    user_name varchar(255) not null
);

-- Связанная таблица зависимости пользователей и заказов
CREATE TABLE users_list
(
    id       serial                                       not null unique,
    user_id  int references users (id) on delete cascade  not null,
    order_id int references orders (id) on delete cascade not null
);

-- Добавление админа при миграции бд
insert into users
values (0, 'Admin', 'admin@pragency.ru',
        '766c6232356876326c6b3532336a356c32336276356c3233b1b3773a05c0ed0176787a4f1574ff0075f7521e', 'admin')