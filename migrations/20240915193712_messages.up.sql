CREATE TABLE request_bodies (
                                id SERIAL PRIMARY KEY,
                                created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                                updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                                deleted_at TIMESTAMP,
                                message TEXT NOT NULL
);
