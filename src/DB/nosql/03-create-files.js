db.createCollection("files", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["fil_id", "fil_desc"],
      properties: {
        _id: {
          bsonType: "objectId",
          description: "MongoDB internal identifier"
        },
        fil_id: {
          bsonType: "binData",
          description: "Unique binary identifier for the file"
        },
        fil_desc: {
          bsonType: "string",
          description: "Description of the file"
        }
      }
    }
  }
});
