create table refresh(
    id SERIAL primary key ,
    user_id varchar(250),
    refresh_token varchar(250)
);