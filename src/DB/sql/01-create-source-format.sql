CREATE TABLE IF NOT EXISTS source_format (
    sfo_id          BINARY(16)    NOT NULL PRIMARY KEY,
    sfo_name        TEXT          NOT NULL,
    sfo_description TEXT          NOT NULL
);
