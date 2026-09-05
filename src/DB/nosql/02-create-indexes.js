// SPDX-License-Identifier: GPL-3.0-or-later

db.sources.createIndex({ src_citation: 1 }, { name: "idx_src_citation" });
db.sources.createIndex({ src_id: 1 }, { unique: true, name: "idx_src_id" });
db.sources.createIndex({ src_url: 1 }, { unique: true, partialFilterExpression: { src_url: { $ne: "" } }, name: "idx_src_url_unique" });
