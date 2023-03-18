/*
 Navicat Premium Data Transfer
 Source Server         : bf
 Source Server Type    : PostgreSQL
 Source Server Version : 150002 (150002)
 Source Host           : localhost:5432
 Source Catalog        : bfdb
 Source Schema         : public
 Target Server Type    : PostgreSQL
 Target Server Version : 150002 (150002)
 File Encoding         : 65001
*/


-- ----------------------------
-- Table structure for login_infos
-- ----------------------------
DROP TABLE IF EXISTS "public"."login_infos";
CREATE TABLE "public"."login_infos" (
  "user_id" text COLLATE "pg_catalog"."default",
  "password" text COLLATE "pg_catalog"."default",
  "salt" text COLLATE "pg_catalog"."default",
  "jwt" text COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."login_infos"."salt" IS '加密盐';
COMMENT ON COLUMN "public"."login_infos"."jwt" IS '鉴权值';

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


-- ----------------------------
-- Table structure for devices
-- ----------------------------
DROP TABLE IF EXISTS "public"."devices";
CREATE TABLE "public"."devices" (
  "id" text COLLATE "pg_catalog"."default" NOT NULL,
  "owner_id" uuid,
  "platform_id" uuid,
  "is_activated" bool,
  "activated_time" timestamptz(6),
  "updated_time" timestamptz(6),
  "is_online" bool,
  "profit" numeric
)
;
COMMENT ON COLUMN "public"."devices"."owner_id" IS '车主ID';
COMMENT ON COLUMN "public"."devices"."is_activated" IS '1为激活，0为未激活';
COMMENT ON COLUMN "public"."devices"."activated_time" IS '激活时间';
COMMENT ON COLUMN "public"."devices"."updated_time" IS '更新时间';
COMMENT ON COLUMN "public"."devices"."is_online" IS '1为上线，0为未上线';
COMMENT ON COLUMN "public"."devices"."profit" IS '收益额';

-- ----------------------------
-- Primary Key structure for table devices
-- ----------------------------
ALTER TABLE "public"."devices" ADD CONSTRAINT "devices_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table devices
-- ----------------------------
ALTER TABLE "public"."devices" ADD CONSTRAINT "fk_drivers_devices" FOREIGN KEY ("owner_id") REFERENCES "public"."drivers" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."devices" ADD CONSTRAINT "fk_platforms_devices" FOREIGN KEY ("platform_id") REFERENCES "public"."platforms" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


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


-- ----------------------------
-- Table structure for driver_order_forms
-- ----------------------------
DROP TABLE IF EXISTS "public"."driver_order_forms";
CREATE TABLE "public"."driver_order_forms" (
  "driver_id" uuid,
  "car_id" text COLLATE "pg_catalog"."default",
  "comment" text COLLATE "pg_catalog"."default",
  "get_time" timestamptz(6),
  "order_id" text COLLATE "pg_catalog"."default" NOT NULL,
  "cost" int8,
  "state" int2,
  "location" text COLLATE "pg_catalog"."default",
  "driver_car_id" text COLLATE "pg_catalog"."default",
  "placetime" timestamptz(6),
  "paytime" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."driver_order_forms"."comment" IS '备注';
COMMENT ON COLUMN "public"."driver_order_forms"."get_time" IS '自取时间';
COMMENT ON COLUMN "public"."driver_order_forms"."cost" IS '花费';
COMMENT ON COLUMN "public"."driver_order_forms"."state" IS '订单状态 2-已完成 1-待取货 0-未支付';
COMMENT ON COLUMN "public"."driver_order_forms"."location" IS '支付时存储位置(购物时获取车主位置）';
COMMENT ON COLUMN "public"."driver_order_forms"."driver_car_id" IS '支付时存储车主车牌号';
COMMENT ON COLUMN "public"."driver_order_forms"."placetime" IS '下单时间';
COMMENT ON COLUMN "public"."driver_order_forms"."paytime" IS '支付时间';

-- ----------------------------
-- Primary Key structure for table driver_order_forms
-- ----------------------------
ALTER TABLE "public"."driver_order_forms" ADD CONSTRAINT "driver_order_forms_pkey" PRIMARY KEY ("order_id");

-- ----------------------------
-- Foreign Keys structure for table driver_order_forms
-- ----------------------------
ALTER TABLE "public"."driver_order_forms" ADD CONSTRAINT "fk_drivers_driver_order_forms" FOREIGN KEY ("driver_id") REFERENCES "public"."drivers" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;


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
