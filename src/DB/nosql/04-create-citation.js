db.createCollection("citation", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["fil_id", "src_id", "cit_type"],
      properties: {
        _id: {
          bsonType: "objectId",
          description: "MongoDB internal identifier"
        },
        fil_id: {
          bsonType: "binData",
          description: "Unique binary identifier referencing a file"
        },
        src_id: {
          bsonType: "binData",
          description: "Unique binary identifier referencing a source"
        },
        cit_type: {
          bsonType: "int",
          enum: [0, 1, 2, 3],
          description: "Type of citation: 0=Primary Sources, 1=References, 2=Holy, 3=Social Media"
        }
      }
    }
  }
});
