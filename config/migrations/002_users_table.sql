create table if not exists users (
    id uuid primary key not null,
    username text not null,
    password text not null,
    email text not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

ALTER TABLE tasks
ADD COLUMN created_at timestamp default now(),
ADD COLUMN updated_at timestamp default now(),
ADD COLUMN user_id UUID,
ADD CONSTRAINT fk_tasks_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE;

---- create above / drop below ----

drop table if exists users;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
