-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists answers (
    id serial not null primary key,
    lvl integer not null,
    user_id integer not null,
    image text not null,
    foreign key (user_id) references users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists answers;
-- +goose StatementEnd
