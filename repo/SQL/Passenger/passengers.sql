
-- ----------------------------
-- Table structure for passengers
-- ----------------------------
DROP TABLE IF EXISTS "public"."passengers";
CREATE TABLE "public"."passengers" (
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
  "score" int8
)
;
COMMENT ON COLUMN "public"."passengers"."balance" IS '账户余额';
COMMENT ON COLUMN "public"."passengers"."pic" IS '用户头像';
COMMENT ON COLUMN "public"."passengers"."name" IS '用户昵称';
COMMENT ON COLUMN "public"."passengers"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."passengers"."id_card" IS '身份证';
COMMENT ON COLUMN "public"."passengers"."role" IS '身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 ';
COMMENT ON COLUMN "public"."passengers"."level" IS '用户等级';
COMMENT ON COLUMN "public"."passengers"."score" IS '用户积分';

-- ----------------------------
-- Primary Key structure for table passengers
-- ----------------------------
ALTER TABLE "public"."passengers" ADD CONSTRAINT "passengers_pkey" PRIMARY KEY ("id");
