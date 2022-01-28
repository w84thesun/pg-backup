create table if not exists users
(
    id                 bigserial
        primary key,
    created_at         timestamp with time zone not null,
    updated_at         timestamp with time zone not null,
    name               text                     not null,
    address            text                     not null,
    active             boolean default false    not null
);