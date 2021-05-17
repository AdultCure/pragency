CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE users_list
(
    id      serial                                      not null unique,
    user_id int references users (id) on delete cascade not null
);

CREATE TABLE orders
(
    id       serial                                      not null unique,
    category varchar(255)                                not null,
    status   varchar(255)                                not null,
    date     varchar(255)                                not null,
    comment  varchar(255)                                not null,
    user_id  int references users (id) on delete cascade not null
);