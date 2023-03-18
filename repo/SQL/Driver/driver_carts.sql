
-- ----------------------------
-- Table structure for driver_carts
-- ----------------------------
DROP TABLE IF EXISTS "public"."driver_carts";
CREATE TABLE "public"."driver_carts" (
  "driver_id" uuid,
  "factory_name" text COLLATE "pg_catalog"."default",
  "distance" int8,
  "cart_id" text COLLATE "pg_catalog"."default" NOT NULL,
  "total_count" int8,
  "total_amount" numeric
)
;
COMMENT ON COLUMN "public"."driver_carts"."factory_name" IS '购物场站名称';
COMMENT ON COLUMN "public"."driver_carts"."distance" IS '距离场站距离';
COMMENT ON COLUMN "public"."driver_carts"."total_count" IS '全选金额';
COMMENT ON COLUMN "public"."driver_carts"."total_amount" IS '全部商品数量';

-- ----------------------------
-- Primary Key structure for table driver_carts
-- ----------------------------
ALTER TABLE "public"."driver_carts" ADD CONSTRAINT "driver_carts_pkey" PRIMARY KEY ("cart_id");

-- ----------------------------
-- Foreign Keys structure for table driver_carts
-- ----------------------------
ALTER TABLE "public"."driver_carts" ADD CONSTRAINT "fk_drivers_cart" FOREIGN KEY ("driver_id") REFERENCES "public"."drivers" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
