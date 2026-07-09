db.createCollection("sources", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["src_id", "sfo_id", "src_citation", "src_date", "src_publish_date", "src_url"],
      properties: {
        _id: {
          bsonType: "objectId",
          description: "MongoDB internal identifier"
        },
        src_id: {
          bsonType: "binData",
          description: "Unique binary identifier (maps to HTSourceElement.ID)"
        },
        sfo_id: {
          bsonType: "binData",
          description: "UUID of citation format (maps to HTSourceElement.SfoID)"
        },
        src_citation: {
          bsonType: "string",
          description: "Source citation text (maps to HTSourceElement.Citation)"
        },
        src_date: {
          bsonType: "string",
          description: "Date associated with the source (maps to HTSourceElement.Date)"
        },
        src_publish_date: {
          bsonType: "string",
          description: "Publication date of the source (maps to HTSourceElement.PublishDate)"
        },
        src_url: {
          bsonType: "string",
          description: "Source URL (maps to HTSourceElement.URL)"
        }
      }
    }
  }
});
