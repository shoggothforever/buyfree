
-- ----------------------------
-- Table structure for drivers
-- ----------------------------
DROP TABLE IF EXISTS "public"."drivers";
CREATE TABLE "public"."drivers" (
  "platform_id" uuid,
  "location" text COLLATE "pg_catalog"."default",
  "car_id" text COLLATE "pg_catalog"."default",
  "mobile" text COLLATE "pg_catalog"."default",
  "id_card" text COLLATE "pg_catalog"."default",
  "is_auth" bool,
  "id" uuid NOT NULL,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "balance" numeric,
  "pic" text COLLATE "pg_catalog"."default",
  "name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "password" text COLLATE "pg_catalog"."default" NOT NULL,
  "role" int8 NOT NULL,
  "level" int8 NOT NULL
)
;
COMMENT ON COLUMN "public"."drivers"."location" IS '地理位置';
COMMENT ON COLUMN "public"."drivers"."car_id" IS '车牌号';
COMMENT ON COLUMN "public"."drivers"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."drivers"."id_card" IS '身份证';
COMMENT ON COLUMN "public"."drivers"."is_auth" IS '1为已认证，0为未认证';
COMMENT ON COLUMN "public"."drivers"."balance" IS '账户余额';
COMMENT ON COLUMN "public"."drivers"."pic" IS '用户头像';
COMMENT ON COLUMN "public"."drivers"."name" IS '用户昵称';
COMMENT ON COLUMN "public"."drivers"."role" IS '身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 ';
COMMENT ON COLUMN "public"."drivers"."level" IS '用户等级';

-- ----------------------------
-- Primary Key structure for table drivers
-- ----------------------------
ALTER TABLE "public"."drivers" ADD CONSTRAINT "drivers_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table drivers
-- ----------------------------
ALTER TABLE "public"."drivers" ADD CONSTRAINT "fk_platforms_authorized_drivers" FOREIGN KEY ("platform_id") REFERENCES "public"."platforms" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
