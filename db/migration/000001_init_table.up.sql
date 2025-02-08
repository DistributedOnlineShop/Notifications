CREATE TABLE "notifications" (
  "notification_id" UUID PRIMARY KEY NOT NULL,
  "user_id" UUID NOT NULL,
  "title" VARCHAR,
  "message" TEXT,
  "type" VARCHAR NOT NULL,
  "is_read" BOOLEAN NOT NULL DEFAULT FALSE,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);
