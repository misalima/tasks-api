create table if not exists tasks (
  id serial primary key,
  title text not null,
  description text,
  status text default 'pendente'
);
---- create above / drop below ----

drop table tasks;
