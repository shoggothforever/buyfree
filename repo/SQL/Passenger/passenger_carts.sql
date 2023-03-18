
-- ----------------------------
-- Table structure for passenger_carts
-- ----------------------------
DROP TABLE IF EXISTS "public"."passenger_carts";
CREATE TABLE "public"."passenger_carts" (
  "passenger_id" uuid,
  "cart_id" text COLLATE "pg_catalog"."default" NOT NULL,
  "total_count" int8,
  "total_amount" numeric
)
;
COMMENT ON COLUMN "public"."passenger_carts"."total_count" IS '全选金额';
COMMENT ON COLUMN "public"."passenger_carts"."total_amount" IS '全部商品数量';

-- ----------------------------
-- Primary Key structure for table passenger_carts
-- ----------------------------
ALTER TABLE "public"."passenger_carts" ADD CONSTRAINT "passenger_carts_pkey" PRIMARY KEY ("cart_id");

-- ----------------------------
-- Foreign Keys structure for table passenger_carts
-- ----------------------------
ALTER TABLE "public"."passenger_carts" ADD CONSTRAINT "fk_passengers_cart" FOREIGN KEY ("passenger_id") REFERENCES "public"."passengers" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
