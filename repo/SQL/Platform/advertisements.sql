
-- ----------------------------
-- Table structure for advertisements
-- ----------------------------
DROP TABLE IF EXISTS "public"."advertisements";
CREATE TABLE "public"."advertisements" (
  "id" text COLLATE "pg_catalog"."default" NOT NULL,
  "description" text COLLATE "pg_catalog"."default",
  "platform_id" uuid,
  "expected_play_times" int8,
  "now_play_times" int8,
  "invest_fund" numeric,
  "profie" numeric,
  "ad_owner" text COLLATE "pg_catalog"."default",
  "play_url" text COLLATE "pg_catalog"."default",
  "expire_at" timestamptz(6),
  "ad_state" int8
)
;
COMMENT ON COLUMN "public"."advertisements"."description" IS '广告描述';
COMMENT ON COLUMN "public"."advertisements"."expected_play_times" IS '预期播放次数';
COMMENT ON COLUMN "public"."advertisements"."now_play_times" IS '已经播放金额';
COMMENT ON COLUMN "public"."advertisements"."invest_fund" IS '投资金额';
COMMENT ON COLUMN "public"."advertisements"."profie" IS '产生收入';
COMMENT ON COLUMN "public"."advertisements"."ad_owner" IS '广告商';
COMMENT ON COLUMN "public"."advertisements"."play_url" IS '广告播放地址';
COMMENT ON COLUMN "public"."advertisements"."expire_at" IS '截止日期';
COMMENT ON COLUMN "public"."advertisements"."ad_state" IS '广告状态 1上线 ， 0下线';

-- ----------------------------
-- Primary Key structure for table advertisements
-- ----------------------------
ALTER TABLE "public"."advertisements" ADD CONSTRAINT "advertisements_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table advertisements
-- ----------------------------
ALTER TABLE "public"."advertisements" ADD CONSTRAINT "fk_platforms_advertisements" FOREIGN KEY ("platform_id") REFERENCES "public"."platforms" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
