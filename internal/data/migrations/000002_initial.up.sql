BEGIN TRANSACTION;

CREATE TABLE "todolist" (
"id" uuid PRIMARY KEY,
"key" text,
"name" text,
"description" text
);

CREATE TABLE "todolistitem" (
"id" uuid PRIMARY KEY,
"todolist_id" uuid,
"name" text,
"is_complete" bool
);

ALTER TABLE "todolistitem" ADD FOREIGN KEY ("todolist_id") REFERENCES "todolist" ("id");

COMMIT;
