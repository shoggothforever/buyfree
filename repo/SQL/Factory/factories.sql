
-- ----------------------------
-- Table structure for factories
-- ----------------------------
DROP TABLE IF EXISTS "public"."factories";
CREATE TABLE "public"."factories" (
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
  "password_salt" text COLLATE "pg_catalog"."default",
  "address" text COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."factories"."balance" IS '账户余额';
COMMENT ON COLUMN "public"."factories"."pic" IS '用户头像';
COMMENT ON COLUMN "public"."factories"."name" IS '用户昵称';
COMMENT ON COLUMN "public"."factories"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."factories"."id_card" IS '身份证';
COMMENT ON COLUMN "public"."factories"."role" IS '身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 ';
COMMENT ON COLUMN "public"."factories"."level" IS '用户等级';
COMMENT ON COLUMN "public"."factories"."password_salt" IS '年销售量';
COMMENT ON COLUMN "public"."factories"."address" IS '场站位置信息';

-- ----------------------------
-- Primary Key structure for table factories
-- ----------------------------
ALTER TABLE "public"."factories" ADD CONSTRAINT "factories_pkey" PRIMARY KEY ("id");
