CREATE TABLE users (
    id text primary key,
    username text not null unique,
    password text not null,
    created_at timestamptz default now() not null
);

CREATE TABLE lists (
    id text primary key,
    title text not null,
    user_id text not null references users(id) on delete cascade,
    created_at timestamptz default now() not null
);

CREATE TABLE todos (
    id text primary key,
    text text not null,
    completed boolean default false not null,
    list_id text not null references lists(id) on delete cascade,
    user_id text not null references users(id) on delete cascade,
    created_at timestamptz default now() not null
);
