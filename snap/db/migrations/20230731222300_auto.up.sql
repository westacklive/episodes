ALTER TABLE "public"."cards" ADD COLUMN "def_id" character varying NOT NULL;
CREATE UNIQUE INDEX "cards_def_id_index" ON "public"."cards" ("def_id");