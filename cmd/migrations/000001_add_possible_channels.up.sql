CREATE TABLE streaming_catalog (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL NOT NULL,
    channel_id TEXT NOT NULL REFERENCES channels(id) CHECK (channel_id ~ '^cha_[0-9]{6}$')
);
