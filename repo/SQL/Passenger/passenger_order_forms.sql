
-- ----------------------------
-- Table structure for passenger_order_forms
-- ----------------------------
DROP TABLE IF EXISTS "public"."passenger_order_forms";
CREATE TABLE "public"."passenger_order_forms" (
  "passenger_id" uuid,
  "order_id" text COLLATE "pg_catalog"."default" NOT NULL,
  "cost" int8,
  "state" int2,
  "location" text COLLATE "pg_catalog"."default",
  "driver_car_id" text COLLATE "pg_catalog"."default",
  "placetime" timestamptz(6),
  "paytime" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."passenger_order_forms"."cost" IS '花费';
COMMENT ON COLUMN "public"."passenger_order_forms"."state" IS '订单状态 2-已完成 1-待取货 0-未支付';
COMMENT ON COLUMN "public"."passenger_order_forms"."location" IS '支付时存储位置(购物时获取车主位置）';
COMMENT ON COLUMN "public"."passenger_order_forms"."driver_car_id" IS '支付时存储车主车牌号';
COMMENT ON COLUMN "public"."passenger_order_forms"."placetime" IS '下单时间';
COMMENT ON COLUMN "public"."passenger_order_forms"."paytime" IS '支付时间';

-- ----------------------------
-- Primary Key structure for table passenger_order_forms
-- ----------------------------
ALTER TABLE "public"."passenger_order_forms" ADD CONSTRAINT "passenger_order_forms_pkey" PRIMARY KEY ("order_id");

-- ----------------------------
-- Foreign Keys structure for table passenger_order_forms
-- ----------------------------
ALTER TABLE "public"."passenger_order_forms" ADD CONSTRAINT "fk_passengers_order_forms" FOREIGN KEY ("passenger_id") REFERENCES "public"."passengers" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
