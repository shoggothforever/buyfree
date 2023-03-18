

-- ----------------------------
-- Table structure for platforms
-- ----------------------------
DROP TABLE IF EXISTS "public"."platforms";
CREATE TABLE "public"."platforms" (
  "id" uuid NOT NULL,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "balance" numeric,
  "pic" text COLLATE "pg_catalog"."default",
  "name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "password" text COLLATE "pg_catalog"."default" NOT NULL,
  "mobile" text COLLATE "pg_catalog"."default",
  "id_card" text COLLATE "pg_catalog"."default",
  "role" int8 NOT NULL,
  "level" int8 NOT NULL,
  "password_salt" text COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."platforms"."balance" IS '账户余额';
COMMENT ON COLUMN "public"."platforms"."pic" IS '用户头像';
COMMENT ON COLUMN "public"."platforms"."name" IS '用户昵称';
COMMENT ON COLUMN "public"."platforms"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."platforms"."id_card" IS '身份证';
COMMENT ON COLUMN "public"."platforms"."role" IS '身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 ';
COMMENT ON COLUMN "public"."platforms"."level" IS '用户等级';
COMMENT ON COLUMN "public"."platforms"."password_salt" IS '年销售量';

-- ----------------------------
-- Primary Key structure for table platforms
-- ----------------------------
ALTER TABLE "public"."platforms" ADD CONSTRAINT "platforms_pkey" PRIMARY KEY ("id");
