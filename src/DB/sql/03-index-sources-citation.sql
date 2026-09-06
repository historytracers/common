CREATE INDEX IF NOT EXISTS idx_sources_src_citation ON sources (src_citation);

CREATE UNIQUE INDEX IF NOT EXISTS idx_sources_src_url_unique ON sources (src_url) WHERE src_url != '';
