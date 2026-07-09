CREATE TABLE IF NOT EXISTS sources (
    src_id          BINARY(16)    NOT NULL PRIMARY KEY,
    src_citation    TEXT          NOT NULL,
    src_date        TEXT          NOT NULL,
    src_publish_date TEXT         NOT NULL,
    src_url         TEXT          NOT NULL
);
