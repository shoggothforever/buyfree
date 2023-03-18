
-- ----------------------------
-- Table structure for order_products
-- ----------------------------
DROP TABLE IF EXISTS "public"."order_products";
CREATE TABLE "public"."order_products" (
  "cart_refer" text COLLATE "pg_catalog"."default",
  "order_refer" text COLLATE "pg_catalog"."default",
  "factory_refer" text COLLATE "pg_catalog"."default",
  "is_chosen" bool,
  "name" text COLLATE "pg_catalog"."default",
  "type" text COLLATE "pg_catalog"."default",
  "count" int8,
  "prize" numeric
)
;
COMMENT ON COLUMN "public"."order_products"."cart_refer" IS '所属购物车';
COMMENT ON COLUMN "public"."order_products"."order_refer" IS '所属订单';
COMMENT ON COLUMN "public"."order_products"."factory_refer" IS '所属场站';
COMMENT ON COLUMN "public"."order_products"."is_chosen" IS '场站是否上线该产品 1-上线 0-下线';
COMMENT ON COLUMN "public"."order_products"."name" IS '商品名称';
COMMENT ON COLUMN "public"."order_products"."type" IS '商品型号';
COMMENT ON COLUMN "public"."order_products"."count" IS '需求量';
COMMENT ON COLUMN "public"."order_products"."prize" IS '价格,根据所属购物车种类赋予不同类型的价格，用户购物车内该值为零售价,车主购物车内该值为批发价';
