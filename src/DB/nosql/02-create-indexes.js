db.sources.createIndex({ src_citation: 1 }, { name: "idx_src_citation" });
db.sources.createIndex({ src_id: 1 }, { unique: true, name: "idx_src_id" });
