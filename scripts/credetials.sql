create table IF NOT EXISTS credentials (
                                            user_id serial,
                                        satl_hash varchar,
                                        FOREIGN KEY (user_id) REFERENCES signed_users(user_id)
)
