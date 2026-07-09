CREATE TABLE IF NOT EXISTS source_format (
    sfo_id          BINARY(16)    NOT NULL PRIMARY KEY,
    sfo_name        TEXT          NOT NULL,
    sfo_description TEXT          NOT NULL
);

INSERT INTO source_format (sfo_id, sfo_name, sfo_description)
VALUES (UUID_TO_BIN('a1b2c3d4-0000-4000-8000-000000000001'), 'APA', 'American Psychological Association');
