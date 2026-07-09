CREATE TABLE IF NOT EXISTS citation (
    fil_id          BINARY(16)    NOT NULL,
    src_id          BINARY(16)    NOT NULL,
    cit_type        TINYINT       NOT NULL,
    PRIMARY KEY (fil_id, src_id, cit_type),
    INDEX idx_citation_fil_id (fil_id),
    CONSTRAINT chk_cit_type CHECK (cit_type IN (0, 1, 2, 3))
);
