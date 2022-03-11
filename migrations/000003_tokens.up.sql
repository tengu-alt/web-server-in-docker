create table IF NOT EXISTS tokens (
                                      user_id int,
                                      token varchar,
                                      FOREIGN KEY (user_id) REFERENCES signed_users(user_id))


