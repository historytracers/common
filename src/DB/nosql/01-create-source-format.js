db.createCollection("source_format", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["sfo_id", "sfo_name", "sfo_description"],
      properties: {
        _id: {
          bsonType: "objectId",
          description: "MongoDB internal identifier"
        },
        sfo_id: {
          bsonType: "binData",
          description: "Unique binary identifier for the format"
        },
        sfo_name: {
          bsonType: "string",
          description: "Format name"
        },
        sfo_description: {
          bsonType: "string",
          description: "Format description"
        }
      }
    }
  }
});
