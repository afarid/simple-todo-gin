CREATE TABLE "todos"
(
    "id"          SERIAL PRIMARY KEY,
    "name"        varchar   NOT NULL,
    "description" varchar   NOT NULL DEFAULT '',
    "username"    varchar   NOT NULL,
    "created_at"  timestamp NOT NULL DEFAULT (now()),
    "deadline"    timestamp NOT NULL
);
