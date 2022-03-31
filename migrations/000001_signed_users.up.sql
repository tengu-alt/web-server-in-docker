create table IF NOT EXISTS signed_users (
                                            user_id serial,
                                            firstname varchar not null CHECK (firstname <> ''),
                                            lastname varchar not null CHECK (firstname <> ''),
                                            email varchar UNIQUE,
                                            PRIMARY KEY (user_id)

    )
