
-- ----------------------------
-- Table structure for device_products
-- ----------------------------
DROP TABLE IF EXISTS "public"."device_products";
CREATE TABLE "public"."device_products" (
  "id" int8 NOT NULL ,
  "monthly_sales" int8,
  "factory_refer" text COLLATE "pg_catalog"."default",
  "sku" text COLLATE "pg_catalog"."default",
  "name" text COLLATE "pg_catalog"."default",
  "type" text COLLATE "pg_catalog"."default",
  "buy_prize" numeric,
  "supply_prize" numeric,
  "device_id" text COLLATE "pg_catalog"."default",
  CONSTRAINT "device_products_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "fk_devices_products" FOREIGN KEY ("device_id") REFERENCES "public"."devices" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
)
;
ALTER TABLE "public"."device_products"
  OWNER TO "bf";
COMMENT ON COLUMN "public"."device_products"."monthly_sales" IS '月销';
COMMENT ON COLUMN "public"."device_products"."sku" IS '库存控制最小可用单位';
COMMENT ON COLUMN "public"."device_products"."name" IS '产品名称';
COMMENT ON COLUMN "public"."device_products"."type" IS '型号';
COMMENT ON COLUMN "public"."device_products"."buy_prize" IS '销售价';
COMMENT ON COLUMN "public"."device_products"."supply_prize" IS '批发价';
COMMENT ON COLUMN "public"."device_products"."device_id" IS '售货机编号';
