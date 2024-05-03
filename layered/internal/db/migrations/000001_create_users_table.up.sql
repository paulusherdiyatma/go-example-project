CREATE TABLE public."users" (
    "id" VARCHAR PRIMARY KEY,
    "username" VARCHAR,
    "email" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);