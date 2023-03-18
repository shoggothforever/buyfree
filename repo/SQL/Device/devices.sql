
-- ----------------------------
-- Table structure for devices
-- ----------------------------
DROP TABLE IF EXISTS "public"."devices";
CREATE TABLE "public"."devices" (
  "id" text COLLATE "pg_catalog"."default" NOT NULL,
  "owner_id" uuid,
  "platform_id" uuid,
  "is_activated" bool,
  "activated_time" timestamptz(6),
  "updated_time" timestamptz(6),
  "is_online" bool,
  "profit" numeric
)
;
COMMENT ON COLUMN "public"."devices"."owner_id" IS '车主ID';
COMMENT ON COLUMN "public"."devices"."is_activated" IS '1为激活，0为未激活';
COMMENT ON COLUMN "public"."devices"."activated_time" IS '激活时间';
COMMENT ON COLUMN "public"."devices"."updated_time" IS '更新时间';
COMMENT ON COLUMN "public"."devices"."is_online" IS '1为上线，0为未上线';
COMMENT ON COLUMN "public"."devices"."profit" IS '收益额';

-- ----------------------------
-- Primary Key structure for table devices
-- ----------------------------
ALTER TABLE "public"."devices" ADD CONSTRAINT "devices_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table devices
-- ----------------------------
ALTER TABLE "public"."devices" ADD CONSTRAINT "fk_drivers_devices" FOREIGN KEY ("owner_id") REFERENCES "public"."drivers" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."devices" ADD CONSTRAINT "fk_platforms_devices" FOREIGN KEY ("platform_id") REFERENCES "public"."platforms" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
