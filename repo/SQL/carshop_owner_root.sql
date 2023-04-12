/*
 Navicat Premium Data Transfer
 Source Server         : root
 Source Server Type    : PostgreSQL
 Source Server Version : 150002 (150002)
 Source Host           : localhost:5432
 Source Catalog        : root
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
                                        "user_id" INT8 NOT NULL,
                                        "password" TEXT COLLATE "pg_catalog"."default",
                                        "salt" TEXT COLLATE "pg_catalog"."default",
                                        "jwt" TEXT COLLATE "pg_catalog"."default",
                                        "role" INT8 NOT NULL,
                                        "user_name" TEXT COLLATE "pg_catalog"."default" NOT NULL,
                                        CONSTRAINT "login_infos_pkey" PRIMARY KEY ( "role", "user_name" )
);
ALTER TABLE "public"."login_infos" OWNER TO "root";
COMMENT ON COLUMN "public"."login_infos"."salt" IS '加密盐';
COMMENT ON COLUMN "public"."login_infos"."jwt" IS '鉴权值';

-- ----------------------------
-- Table structure for passengers
-- ----------------------------
DROP TABLE IF EXISTS "public"."passengers";
CREATE TABLE "public"."passengers" (
                                       "id" INT8 NOT NULL,
                                       "created_at" TIMESTAMPTZ ( 6 ) NOT NULL DEFAULT now( ),
                                       "updated_at" TIMESTAMPTZ ( 6 ),
                                       "deleted_at" TIMESTAMPTZ ( 6 ),
                                       "balance" NUMERIC NOT NULL DEFAULT 0,
                                       "pic" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                       "name" VARCHAR ( 32 ) COLLATE "pg_catalog"."default" NOT NULL,
                                       "password" VARCHAR ( 40 ) COLLATE "pg_catalog"."default",
                                       "mobile" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                       "id_card" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                       "role" INT8 NOT NULL DEFAULT 0,
                                       "level" INT8 NOT NULL DEFAULT 0,
                                       "score" INT8 NOT NULL DEFAULT 0,
                                       "password_salt" TEXT COLLATE "pg_catalog"."default",
                                       CONSTRAINT "passengers_pkey" PRIMARY KEY ( "id" ),
                                       CONSTRAINT "passengers_name_key" UNIQUE ( "name" )
);
ALTER TABLE "public"."passengers" OWNER TO "root";
COMMENT ON COLUMN "public"."passengers"."balance" IS '账户余额';
COMMENT ON COLUMN "public"."passengers"."pic" IS '用户头像';
COMMENT ON COLUMN "public"."passengers"."name" IS '用户昵称';
COMMENT ON COLUMN "public"."passengers"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."passengers"."id_card" IS '身份证';
COMMENT ON COLUMN "public"."passengers"."role" IS '身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 ';
COMMENT ON COLUMN "public"."passengers"."level" IS '用户等级';
COMMENT ON COLUMN "public"."passengers"."score" IS '用户积分';
COMMENT ON COLUMN "public"."passengers"."password_salt" IS '密码盐';

-- ----------------------------
-- Table structure for passenger_carts
-- ----------------------------
DROP TABLE IF EXISTS "public"."passenger_carts";
CREATE TABLE "public"."passenger_carts" (
                                            "passenger_id" INT8,
                                            "cart_id" INT8 NOT NULL,
                                            "total_count" INT8 NOT NULL DEFAULT 0,
                                            "total_amount" NUMERIC NOT NULL DEFAULT 0,
                                            CONSTRAINT "passenger_carts_pkey" PRIMARY KEY ( "cart_id" ),
                                            CONSTRAINT "fk_passengers_cart" FOREIGN KEY ( "passenger_id" ) REFERENCES "public"."passengers" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."passenger_carts" OWNER TO "root";
COMMENT ON COLUMN "public"."passenger_carts"."total_count" IS '全选金额';
COMMENT ON COLUMN "public"."passenger_carts"."total_amount" IS '全部商品数量';
-- ----------------------------
-- Table structure for passenger_order_forms
-- ----------------------------
DROP TABLE IF EXISTS "public"."passenger_order_forms";
CREATE TABLE "public"."passenger_order_forms" (
                                                  "passenger_id" INT8,
                                                  "driver_car_id" TEXT COLLATE "pg_catalog"."default",
                                                  "order_id" TEXT COLLATE "pg_catalog"."default" NOT NULL,
                                                  "cost" NUMERIC NOT NULL DEFAULT 0,
                                                  "state" INT2 NOT NULL DEFAULT 0,
                                                  "location" VARCHAR ( 40 ) COLLATE "pg_catalog"."default",
                                                  "place_time" TIMESTAMPTZ ( 6 ) NOT NULL DEFAULT now( ),
                                                  "pay_time" TIMESTAMPTZ ( 6 ),
                                                  CONSTRAINT "passenger_order_forms_pkey" PRIMARY KEY ( "order_id" ),
                                                  CONSTRAINT "fk_passengers_order_forms" FOREIGN KEY ( "passenger_id" ) REFERENCES "public"."passengers" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."passenger_order_forms" OWNER TO "root";
COMMENT ON COLUMN "public"."passenger_order_forms"."driver_car_id" IS '支付时存储车主车牌号';
COMMENT ON COLUMN "public"."passenger_order_forms"."cost" IS '花费';
COMMENT ON COLUMN "public"."passenger_order_forms"."state" IS '订单状态 2-已完成 1-待取货 0-未支付';
COMMENT ON COLUMN "public"."passenger_order_forms"."location" IS '支付时存储位置(购物时获取车主位置）';
COMMENT ON COLUMN "public"."passenger_order_forms"."place_time" IS '下单时间';
COMMENT ON COLUMN "public"."passenger_order_forms"."pay_time" IS '支付时间';
-- ----------------------------
-- Table structure for factories
-- ----------------------------
DROP TABLE IF EXISTS "public"."factories";
CREATE TABLE "public"."factories" (
                                      "id" INT8 NOT NULL,
                                      "created_at" TIMESTAMPTZ ( 6 ) DEFAULT now( ),
                                      "updated_at" TIMESTAMPTZ ( 6 ),
                                      "deleted_at" TIMESTAMPTZ ( 6 ),
                                      "balance" NUMERIC DEFAULT 0,
                                      "pic" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" DEFAULT ' ' :: CHARACTER VARYING,
                                      "name" VARCHAR ( 32 ) COLLATE "pg_catalog"."default" NOT NULL,
                                      "password" VARCHAR COLLATE "pg_catalog"."default",
                                      "mobile" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                      "id_card" VARCHAR ( 20 ) COLLATE "pg_catalog"."default",
                                      "role" INT8 NOT NULL DEFAULT 2,
                                      "level" INT8 NOT NULL DEFAULT 0,
                                      "password_salt" VARCHAR ( 20 ) COLLATE "pg_catalog"."default",
                                      "address" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" NOT NULL,
                                      "longitude" TEXT COLLATE "pg_catalog"."default" NOT NULL,
                                      "latitude" TEXT COLLATE "pg_catalog"."default" NOT NULL,
                                      "description" TEXT COLLATE "pg_catalog"."default",
                                      CONSTRAINT "factories_pkey" PRIMARY KEY ( "id" ),
                                      CONSTRAINT "factories_name_key" UNIQUE ( "name" )
);
ALTER TABLE "public"."factories" OWNER TO "root";
COMMENT ON COLUMN "public"."factories"."balance" IS '账户余额';
COMMENT ON COLUMN "public"."factories"."pic" IS '用户头像';
COMMENT ON COLUMN "public"."factories"."name" IS '用户昵称';
COMMENT ON COLUMN "public"."factories"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."factories"."id_card" IS '身份证';
COMMENT ON COLUMN "public"."factories"."role" IS '身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 ';
COMMENT ON COLUMN "public"."factories"."level" IS '用户等级';
COMMENT ON COLUMN "public"."factories"."password_salt" IS '密码盐';
COMMENT ON COLUMN "public"."factories"."address" IS '场站位置信息';
COMMENT ON COLUMN "public"."factories"."longitude" IS '经度';
COMMENT ON COLUMN "public"."factories"."latitude" IS '纬度';
COMMENT ON COLUMN "public"."factories"."description" IS '场站描述';
-- ----------------------------
-- Table structure for factory_products
-- ----------------------------
DROP TABLE IF EXISTS "public"."factory_products";
CREATE TABLE "public"."factory_products" (
                                             "factory_name" TEXT COLLATE "pg_catalog"."default",
                                             "id" INT8 NOT NULL ,
                                             "factory_id" INT8 NOT NULL,
                                             "sku" VARCHAR COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                             "inventory" INT8 NOT NULL DEFAULT 0,
                                             "name" TEXT COLLATE "pg_catalog"."default" NOT NULL,
                                             "pic" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                             "type" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                             "buy_price" NUMERIC NOT NULL DEFAULT 99999,
                                             "supply_price" NUMERIC NOT NULL DEFAULT 99999,
                                             "daily_sales" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 0,
                                             "weekly_sales" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 0,
                                             "monthly_sales" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 0,
                                             "annually_sales" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 0,
                                             "total_sales" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 0,
                                             "is_on_shelf" BOOL NOT NULL DEFAULT TRUE,
                                             CONSTRAINT "factory_products_pkey" PRIMARY KEY ( "id", "factory_id", "name" ),
                                             CONSTRAINT "fk_factories_products" FOREIGN KEY ( "factory_id" ) REFERENCES "public"."factories" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."factory_products" OWNER TO "root";
COMMENT ON COLUMN "public"."factory_products"."factory_id" IS '指向场站的编号';
COMMENT ON COLUMN "public"."factory_products"."sku" IS '库存控制最小可用单位';
COMMENT ON COLUMN "public"."factory_products"."inventory" IS '存货';
COMMENT ON COLUMN "public"."factory_products"."name" IS '产品名称';
COMMENT ON COLUMN "public"."factory_products"."pic" IS '图片';
COMMENT ON COLUMN "public"."factory_products"."type" IS '型号';
COMMENT ON COLUMN "public"."factory_products"."buy_price" IS '销售价';
COMMENT ON COLUMN "public"."factory_products"."supply_price" IS '批发价';
COMMENT ON COLUMN "public"."factory_products"."daily_sales" IS '日销量';
COMMENT ON COLUMN "public"."factory_products"."weekly_sales" IS '周销量';
COMMENT ON COLUMN "public"."factory_products"."monthly_sales" IS '月销量';
COMMENT ON COLUMN "public"."factory_products"."annually_sales" IS '年销售量';
COMMENT ON COLUMN "public"."factory_products"."total_sales" IS '总销售量';


-- ----------------------------
-- Table structure for platforms
-- ----------------------------
DROP TABLE IF EXISTS "public"."platforms";
CREATE TABLE "public"."platforms" (
                                      "id" INT8 NOT NULL ,
                                      "created_at" TIMESTAMPTZ ( 6 ) NOT NULL DEFAULT now( ),
                                      "updated_at" TIMESTAMPTZ ( 6 ),
                                      "deleted_at" TIMESTAMPTZ ( 6 ),
                                      "balance" NUMERIC NOT NULL DEFAULT 0,
                                      "pic" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                      "name" VARCHAR ( 32 ) COLLATE "pg_catalog"."default" NOT NULL,
                                      "password" VARCHAR ( 40 ) COLLATE "pg_catalog"."default",
                                      "mobile" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" DEFAULT ' ' :: CHARACTER VARYING,
                                      "id_card" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" DEFAULT ' ' :: CHARACTER VARYING,
                                      "role" INT8 NOT NULL DEFAULT 3,
                                      "level" INT8 NOT NULL,
                                      "password_salt" TEXT COLLATE "pg_catalog"."default",
                                      CONSTRAINT "platforms_pkey" PRIMARY KEY ( "id" ),
                                      CONSTRAINT "platforms_name_key" UNIQUE ( "name" )
);
ALTER TABLE "public"."platforms" OWNER TO "root";
COMMENT ON COLUMN "public"."platforms"."balance" IS '账户余额';
COMMENT ON COLUMN "public"."platforms"."pic" IS '用户头像';
COMMENT ON COLUMN "public"."platforms"."name" IS '用户昵称';
COMMENT ON COLUMN "public"."platforms"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."platforms"."id_card" IS '身份证';
COMMENT ON COLUMN "public"."platforms"."role" IS '身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 ';
COMMENT ON COLUMN "public"."platforms"."level" IS '用户等级';
COMMENT ON COLUMN "public"."platforms"."password_salt" IS '年销售量';
-- ----------------------------
-- Table structure for drivers
-- ----------------------------
DROP TABLE IF EXISTS "public"."drivers";
CREATE TABLE "public"."drivers" (
                                    "id" INT8 NOT NULL,
                                    "created_at" TIMESTAMPTZ ( 6 ) NOT NULL DEFAULT now( ),
                                    "updated_at" TIMESTAMPTZ ( 6 ),
                                    "deleted_at" TIMESTAMPTZ ( 6 ),
                                    "balance" NUMERIC,
                                    "pic" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                    "name" VARCHAR ( 32 ) COLLATE "pg_catalog"."default" NOT NULL,
                                    "password" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" DEFAULT 123456,
                                    "mobile" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                    "id_card" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                    "role" INT8 NOT NULL DEFAULT 1,
                                    "level" INT8 NOT NULL DEFAULT 0,
                                    "car_id" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                    "platform_id" INT8 NOT NULL DEFAULT 0,
                                    "is_auth" BOOL NOT NULL DEFAULT FALSE,
                                    "password_salt" VARCHAR ( 20 ) COLLATE "pg_catalog"."default",
                                    "address" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                    "longitude" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 120,
                                    "latitude" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 30,
                                    CONSTRAINT "drivers_pkey" PRIMARY KEY ( "id" ),
                                    CONSTRAINT "drivers_name_key" UNIQUE ( "name" )
);
ALTER TABLE "public"."drivers" OWNER TO "root";
COMMENT ON COLUMN "public"."drivers"."balance" IS '账户余额';
COMMENT ON COLUMN "public"."drivers"."pic" IS '用户头像';
COMMENT ON COLUMN "public"."drivers"."name" IS '用户昵称';
COMMENT ON COLUMN "public"."drivers"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."drivers"."id_card" IS '身份证';
COMMENT ON COLUMN "public"."drivers"."role" IS '身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 ';
COMMENT ON COLUMN "public"."drivers"."level" IS '用户等级';
COMMENT ON COLUMN "public"."drivers"."car_id" IS '车牌号';
COMMENT ON COLUMN "public"."drivers"."is_auth" IS '1为已认证，0为未认证';
COMMENT ON COLUMN "public"."drivers"."password_salt" IS '密码盐';
COMMENT ON COLUMN "public"."drivers"."address" IS '车主位置信息';
COMMENT ON COLUMN "public"."drivers"."longitude" IS '经度';
COMMENT ON COLUMN "public"."drivers"."latitude" IS '纬度';

-- ----------------------------
-- Table structure for order_products
-- ----------------------------
DROP TABLE IF EXISTS "public"."order_products";
CREATE TABLE "public"."order_products" (
                                           "cart_refer" INT8 NOT NULL,
                                           "factory_id" INT8 NOT NULL,
                                           "order_refer" INT8 NOT NULL,
                                           "is_chosen" BOOL NOT NULL DEFAULT TRUE,
                                           "name" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL,
                                           "sku" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL,
                                           "pic" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                           "type" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL,
                                           "count" INT8 NOT NULL DEFAULT 0,
                                           "price" NUMERIC NOT NULL DEFAULT 99999,
                                           CONSTRAINT "order_products_pkey" PRIMARY KEY ( "type", "name", "factory_id", "cart_refer", "order_refer" )
);
ALTER TABLE "public"."order_products" OWNER TO "root";
COMMENT ON COLUMN "public"."order_products"."cart_refer" IS '所属购物车';
COMMENT ON COLUMN "public"."order_products"."factory_id" IS '所属场站';
COMMENT ON COLUMN "public"."order_products"."order_refer" IS '所属订单';
COMMENT ON COLUMN "public"."order_products"."is_chosen" IS '场站是否上线该产品 1-上线 0-下线';
COMMENT ON COLUMN "public"."order_products"."name" IS '商品名称';
COMMENT ON COLUMN "public"."order_products"."sku" IS '库存控制最小可用单位';
COMMENT ON COLUMN "public"."order_products"."pic" IS '图片';
COMMENT ON COLUMN "public"."order_products"."type" IS '商品型号';
COMMENT ON COLUMN "public"."order_products"."count" IS '需求量';
COMMENT ON COLUMN "public"."order_products"."price" IS '价格,根据所属购物车种类赋予不同类型的价格，用户购物车内该值为零售价,车主购物车内该值为批发价';

-- ----------------------------
-- Table structure for cart_products
-- ----------------------------
DROP TABLE IF EXISTS "public"."cart_products";
CREATE TABLE "public"."cart_products" (
                                          "cart_refer" INT8 NOT NULL,
                                          "factory_id" INT8 NOT NULL,
                                          "order_refer" INT8 NOT NULL DEFAULT 0,
                                          "is_chosen" BOOL NOT NULL DEFAULT TRUE,
                                          "name" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                          "sku" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                          "pic" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                          "type" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                          "count" INT8 NOT NULL DEFAULT 0,
                                          "price" NUMERIC NOT NULL DEFAULT 99999,
                                          CONSTRAINT "cart_products_pkey" PRIMARY KEY ( "type", "name", "factory_id", "cart_refer", "order_refer" )
);
ALTER TABLE "public"."cart_products" OWNER TO "root";
COMMENT ON COLUMN "public"."cart_products"."cart_refer" IS '所属购物车';
COMMENT ON COLUMN "public"."cart_products"."factory_id" IS '所属场站';
COMMENT ON COLUMN "public"."cart_products"."order_refer" IS '所属订单';
COMMENT ON COLUMN "public"."cart_products"."is_chosen" IS '场站是否上线该产品 1-上线 0-下线';
COMMENT ON COLUMN "public"."cart_products"."name" IS '商品名称';
COMMENT ON COLUMN "public"."cart_products"."sku" IS '库存控制最小可用单位';
COMMENT ON COLUMN "public"."cart_products"."pic" IS '图片';
COMMENT ON COLUMN "public"."cart_products"."type" IS '商品型号';
COMMENT ON COLUMN "public"."cart_products"."count" IS '需求量';
COMMENT ON COLUMN "public"."cart_products"."price" IS '价格,根据所属购物车种类赋予不同类型的价格，用户购物车内该值为零售价,车主购物车内该值为批发价';

-- ----------------------------
-- Table structure for devices
-- ----------------------------
DROP TABLE IF EXISTS "public"."devices";
CREATE TABLE "public"."devices" (
                                    "id" INT8 NOT NULL,
                                    "owner_id" INT8,
                                    "platform_id" INT8,
                                    "is_activated" BOOL NOT NULL DEFAULT FALSE,
                                    "activated_time" TIMESTAMPTZ ( 6 ) NOT NULL DEFAULT now( ),
                                    "updated_time" TIMESTAMPTZ ( 6 ) NOT NULL DEFAULT now( ),
                                    "is_online" BOOL NOT NULL DEFAULT FALSE,
                                    "profit" NUMERIC NOT NULL DEFAULT 0,
                                    CONSTRAINT "devices_pkey" PRIMARY KEY ( "id" ),
                                    CONSTRAINT "fk_drivers_devices" FOREIGN KEY ( "owner_id" ) REFERENCES "public"."drivers" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION,
                                    CONSTRAINT "fk_platforms_devices" FOREIGN KEY ( "platform_id" ) REFERENCES "public"."platforms" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."devices" OWNER TO "root";
COMMENT ON COLUMN "public"."devices"."owner_id" IS '车主ID';
COMMENT ON COLUMN "public"."devices"."is_activated" IS 'true为激活，false为未激活';
COMMENT ON COLUMN "public"."devices"."activated_time" IS '激活时间';
COMMENT ON COLUMN "public"."devices"."updated_time" IS '更新时间';
COMMENT ON COLUMN "public"."devices"."is_online" IS 'true为上线，false为未上线';
COMMENT ON COLUMN "public"."devices"."profit" IS '收益额';

-- ----------------------------
-- Table structure for driver_carts
-- ----------------------------
DROP TABLE IF EXISTS "public"."driver_carts";
CREATE TABLE "public"."driver_carts" (
                                         "driver_id" INT8,
                                         "cart_id" INT8 NOT NULL,
                                         "total_count" INT8 NOT NULL DEFAULT 0,
                                         "total_amount" NUMERIC NOT NULL DEFAULT 0,
                                         CONSTRAINT "driver_carts_pkey" PRIMARY KEY ( "cart_id" ),
                                         CONSTRAINT "fk_drivers_cart" FOREIGN KEY ( "driver_id" ) REFERENCES "public"."drivers" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."driver_carts" OWNER TO "root";
COMMENT ON COLUMN "public"."driver_carts"."total_count" IS '全选金额';
COMMENT ON COLUMN "public"."driver_carts"."total_amount" IS '全部商品数量';

-- ----------------------------
-- Table structure for driver_order_forms
-- ----------------------------
DROP TABLE IF EXISTS "public"."driver_order_forms";
CREATE TABLE "public"."driver_order_forms" (
                                               "factory_id" INT8,
                                               "factory_name" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                               "driver_id" INT8,
                                               "car_id" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                               "comment" TEXT COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: TEXT,
                                               "get_time" TIMESTAMPTZ ( 6 ) NOT NULL,
                                               "order_id" INT8 NOT NULL,
                                               "cost" NUMERIC ( 64, 0 ) NOT NULL DEFAULT 0,
                                               "state" INT2 NOT NULL DEFAULT 0,
                                               "location" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                               "place_time" TIMESTAMPTZ ( 6 ) NOT NULL DEFAULT now( ),
                                               "pay_time" TIMESTAMPTZ ( 6 ) NOT NULL DEFAULT now( ),
                                               CONSTRAINT "driver_order_forms_pkey" PRIMARY KEY ( "order_id" ),
                                               CONSTRAINT "fk_drivers_driver_order_forms" FOREIGN KEY ( "driver_id" ) REFERENCES "public"."drivers" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION,
                                               CONSTRAINT "fk_factories_order_forms" FOREIGN KEY ( "factory_id" ) REFERENCES "public"."factories" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."driver_order_forms" OWNER TO "root";
COMMENT ON COLUMN "public"."driver_order_forms"."factory_id" IS '指向factory.id';
COMMENT ON COLUMN "public"."driver_order_forms"."factory_name" IS '订单发货场站名';
COMMENT ON COLUMN "public"."driver_order_forms"."comment" IS '备注';
COMMENT ON COLUMN "public"."driver_order_forms"."get_time" IS '自取时间';
COMMENT ON COLUMN "public"."driver_order_forms"."cost" IS '花费';
COMMENT ON COLUMN "public"."driver_order_forms"."state" IS '订单状态 2-已完成 1-待取货 0-未支付';
COMMENT ON COLUMN "public"."driver_order_forms"."location" IS '支付时存储位置(购物时获取车主位置）';
COMMENT ON COLUMN "public"."driver_order_forms"."place_time" IS '下单时间';
COMMENT ON COLUMN "public"."driver_order_forms"."pay_time" IS '支付时间';

-- ----------------------------
-- Table structure for device_products
-- ----------------------------
DROP TABLE IF EXISTS "public"."device_products";
CREATE TABLE "public"."device_products" (
                                            "device_id" INT8,
                                            "id" INT8 NOT NULL,
                                            "factory_id" INT8,
                                            "sku" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                            "inventory" INT8 NOT NULL DEFAULT 0,
                                            "name" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                            "pic" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                            "type" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                            "buy_price" NUMERIC NOT NULL DEFAULT 99999,
                                            "supply_price" NUMERIC NOT NULL DEFAULT 99999,
                                            "daily_sales" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 0,
                                            "weekly_sales" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 0,
                                            "monthly_sales" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 0,
                                            "annually_sales" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 0,
                                            "total_sales" VARCHAR ( 40 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 0,
                                            "driver_id" INT8,
                                            CONSTRAINT "device_products_pkey" PRIMARY KEY ( "id" ),
                                            CONSTRAINT "fk_devices_products" FOREIGN KEY ( "device_id" ) REFERENCES "public"."devices" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."device_products" OWNER TO "root";
COMMENT ON COLUMN "public"."device_products"."device_id" IS '售货机编号';
COMMENT ON COLUMN "public"."device_products"."factory_id" IS '指向场站的编号';
COMMENT ON COLUMN "public"."device_products"."sku" IS '库存控制最小可用单位';
COMMENT ON COLUMN "public"."device_products"."inventory" IS '存货';
COMMENT ON COLUMN "public"."device_products"."name" IS '产品名称';
COMMENT ON COLUMN "public"."device_products"."pic" IS '图片';
COMMENT ON COLUMN "public"."device_products"."type" IS '型号';
COMMENT ON COLUMN "public"."device_products"."buy_price" IS '销售价';
COMMENT ON COLUMN "public"."device_products"."supply_price" IS '批发价';
COMMENT ON COLUMN "public"."device_products"."daily_sales" IS '日销量';
COMMENT ON COLUMN "public"."device_products"."weekly_sales" IS '周销量';
COMMENT ON COLUMN "public"."device_products"."monthly_sales" IS '月销量';
COMMENT ON COLUMN "public"."device_products"."annually_sales" IS '年销售量';
COMMENT ON COLUMN "public"."device_products"."total_sales" IS '总销售量';
COMMENT ON COLUMN "public"."device_products"."driver_id" IS '指向车主的编号';
-- ----------------------------
-- Table structure for advertisements
-- ----------------------------
DROP TABLE IF EXISTS "public"."advertisements";
CREATE TABLE "public"."advertisements" (
                                           "id" INT8 NOT NULL,
                                           "description" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                           "platform_id" INT8,
                                           "expected_play_times" INT8 NOT NULL DEFAULT 0,
                                           "play_times" INT8 NOT NULL DEFAULT 0,
                                           "invest_fund" NUMERIC NOT NULL DEFAULT 0,
                                           "profit" NUMERIC NOT NULL DEFAULT 0,
                                           "video_cover" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                           "ad_owner" VARCHAR ( 20 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                           "play_url" VARCHAR ( 255 ) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ' ' :: CHARACTER VARYING,
                                           "expire_at" TIMESTAMPTZ ( 6 ),
                                           "ad_state" INT8,
                                           CONSTRAINT "advertisements_pkey" PRIMARY KEY ( "id" ),
                                           CONSTRAINT "fk_platforms_advertisements" FOREIGN KEY ( "platform_id" ) REFERENCES "public"."platforms" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."advertisements" OWNER TO "root";
COMMENT ON COLUMN "public"."advertisements"."description" IS '广告描述';
COMMENT ON COLUMN "public"."advertisements"."expected_play_times" IS '预期播放次数';
COMMENT ON COLUMN "public"."advertisements"."play_times" IS '已经播放金额';
COMMENT ON COLUMN "public"."advertisements"."invest_fund" IS '投资金额';
COMMENT ON COLUMN "public"."advertisements"."profit" IS '产生收入';
COMMENT ON COLUMN "public"."advertisements"."video_cover" IS '广告封面地址';
COMMENT ON COLUMN "public"."advertisements"."ad_owner" IS '广告商';
COMMENT ON COLUMN "public"."advertisements"."play_url" IS '广告播放地址';
COMMENT ON COLUMN "public"."advertisements"."expire_at" IS '截止日期';
COMMENT ON COLUMN "public"."advertisements"."ad_state" IS '广告状态 1上线 ， 0下线';


-- ----------------------------
-- Table structure for fund_infos
-- ----------------------------
DROP TABLE IF EXISTS "public"."fund_infos";
CREATE TABLE "public"."fund_infos" (
                                       "user_id" INT8 NOT NULL,
                                       "card_id" INT8,
                                       "bank_name" TEXT COLLATE "pg_catalog"."default",
                                       "cash" NUMERIC,
                                       "bank_fund" NUMERIC,
                                       CONSTRAINT "bank_card_infos_pkey" PRIMARY KEY ( "user_id" ),
                                       CONSTRAINT "bank_card_infos_card_id_key" UNIQUE ( "card_id" )
);
ALTER TABLE "public"."fund_infos" OWNER TO "root";
COMMENT ON COLUMN "public"."fund_infos"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."fund_infos"."cash" IS '账户余额';
COMMENT ON COLUMN "public"."fund_infos"."bank_fund" IS '银行资金';
-- ----------------------------
-- Table structure for ad_devices
-- ----------------------------
DROP TABLE IF EXISTS "public"."ad_devices";
CREATE TABLE "public"."ad_devices" (
                                       "advertisement_id" INT8 NOT NULL,
                                       "device_id" INT8 NOT NULL,
                                       "play_times" INT8 NOT NULL DEFAULT 0,
                                       "profit" NUMERIC NOT NULL DEFAULT 0,
                                       CONSTRAINT "ad_devices_pkey" PRIMARY KEY ( "advertisement_id", "device_id" )
);
ALTER TABLE "public"."ad_devices" OWNER TO "root";
