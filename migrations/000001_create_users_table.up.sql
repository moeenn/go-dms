create table
  users (
    id uuid default gen_random_uuid (),
    email text not null,
    password text not null,
    created_at timestamp default now(),
    deleted_at timestamp null,
    primary key (id),
    constraint email_unique unique (email)
  );