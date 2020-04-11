CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "events" (
"id" TEXT PRIMARY KEY,
"f_with" TEXT NOT NULL,
"f_what" TEXT NOT NULL,
"f_when" DATETIME NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
, "f_where" TEXT NOT NULL DEFAULT '');
CREATE TABLE IF NOT EXISTS "votes" (
"id" TEXT PRIMARY KEY,
"author" TEXT NOT NULL,
"vote" smallint NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
, name TEXT NOT NULL DEFAULT '');
CREATE TABLE IF NOT EXISTS "users" (
"id" TEXT PRIMARY KEY,
"email" TEXT NOT NULL,
"password_hash" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
