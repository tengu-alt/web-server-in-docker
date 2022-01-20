create table IF NOT EXISTS signed_users (
                              firstname varchar,
                              lastname varchar,
                              email varchar UNIQUE,
                              user_password varchar
)