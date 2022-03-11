create table IF NOT EXISTS signed_users (
                                            user_id serial,
                                            firstname varchar,
                                            lastname varchar,
                                            email varchar UNIQUE,
                                            PRIMARY KEY (user_id)

    )
