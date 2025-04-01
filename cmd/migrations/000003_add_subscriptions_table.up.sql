CREATE TABLE SUBSCRIPTIONS (
    id TEXT PRIMARY KEY,
    user_id TEXT REFERENCES users(id),
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    channel_id TEXT REFERENCES streaming_catalog(id),
    inform_date TIMESTAMP,
    note TEXT
);