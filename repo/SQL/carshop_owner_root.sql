/*
 Navicat Premium Data Transfer
 Source Server         : root

 Source Server Type    : PostgreSQL
 Source Server Version : 150002 (150002)
 Source Host           : localhost:5432
 Source Catalog        : root
 db
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
                                        "user_name" TEXT COLLATE "pg_catalog"."default",
                                        CONSTRAINT "login_infos_pkey" PRIMARY KEY ( "user_id", "role" )
);
ALTER TABLE "public"."login_infos" OWNER TO "root
";
COMMENT ON COLUMN "public"."login_infos"."salt" IS '加密盐';
COMMENT ON COLUMN "public"."login_infos"."jwt" IS '鉴权值';
COMMENT ON COLUMN "public"."login_infos"."user_name" IS '用户密码';

-- ----------------------------
-- Table structure for passengers
-- ----------------------------
DROP TABLE IF EXISTS "public"."passengers";
CREATE TABLE "public"."passengers" (
                                       "id" INT8 NOT NULL ,
                                       "created_at" TIMESTAMPTZ ( 6 ),
                                       "updated_at" TIMESTAMPTZ ( 6 ),
                                       "deleted_at" TIMESTAMPTZ ( 6 ),
                                       "balance" NUMERIC,
                                       "pic" TEXT COLLATE "pg_catalog"."default",
                                       "name" VARCHAR ( 32 ) COLLATE "pg_catalog"."default" NOT NULL,
                                       "password" TEXT COLLATE "pg_catalog"."default" NOT NULL,
                                       "password_salt" TEXT COLLATE "pg_catalog"."default",
                                       "mobile" TEXT COLLATE "pg_catalog"."default",
                                       "id_card" TEXT COLLATE "pg_catalog"."default",
                                       "role" INT8 NOT NULL,
                                       "level" INT8 NOT NULL,
                                       "score" INT8,
                                       CONSTRAINT "passengers_pkey" PRIMARY KEY ( "id" )
);
ALTER TABLE "public"."passengers" OWNER TO "root
";
COMMENT ON COLUMN "public"."passengers"."balance" IS '账户余额';
COMMENT ON COLUMN "public"."passengers"."pic" IS '用户头像';
COMMENT ON COLUMN "public"."passengers"."name" IS '用户昵称';
COMMENT ON COLUMN "public"."passengers"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."passengers"."id_card" IS '身份证';
COMMENT ON COLUMN "public"."passengers"."role" IS '身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 ';
COMMENT ON COLUMN "public"."passengers"."level" IS '用户等级';
COMMENT ON COLUMN "public"."passengers"."score" IS '用户积分';

-- ----------------------------
-- Table structure for passenger_carts
-- ----------------------------
DROP TABLE IF EXISTS "public"."passenger_carts";
CREATE TABLE "public"."passenger_carts" (
                                            "passenger_id" INT8,
                                            "cart_id" INT8 NOT NULL ,
                                            "total_count" INT8,
                                            "total_amount" NUMERIC,
                                            CONSTRAINT "passenger_carts_pkey" PRIMARY KEY ( "cart_id" ),
                                            CONSTRAINT "fk_passengers_cart" FOREIGN KEY ( "passenger_id" ) REFERENCES "public"."passengers" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."passenger_carts" OWNER TO "root
";
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
                                                  "cost" INT8,
                                                  "state" INT2,
                                                  "location" TEXT COLLATE "pg_catalog"."default",
                                                  "placetime" TIMESTAMPTZ ( 6 ),
                                                  "paytime" TIMESTAMPTZ ( 6 ),
                                                  CONSTRAINT "passenger_order_forms_pkey" PRIMARY KEY ( "order_id" ),
                                                  CONSTRAINT "fk_passengers_order_forms" FOREIGN KEY ( "passenger_id" ) REFERENCES "public"."passengers" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."passenger_order_forms" OWNER TO "root
";
COMMENT ON COLUMN "public"."passenger_order_forms"."driver_car_id" IS '支付时存储车主车牌号';
COMMENT ON COLUMN "public"."passenger_order_forms"."cost" IS '花费';
COMMENT ON COLUMN "public"."passenger_order_forms"."state" IS '订单状态 2-已完成 1-待取货 0-未支付';
COMMENT ON COLUMN "public"."passenger_order_forms"."location" IS '支付时存储位置(购物时获取车主位置）';
COMMENT ON COLUMN "public"."passenger_order_forms"."placetime" IS '下单时间';
COMMENT ON COLUMN "public"."passenger_order_forms"."paytime" IS '支付时间';
-- ----------------------------
-- Table structure for factories
-- ----------------------------
DROP TABLE IF EXISTS "public"."factories";
CREATE TABLE "public"."factories" (
                                      "id" INT8 NOT NULL,
                                      "created_at" TIMESTAMPTZ ( 6 ),
                                      "updated_at" TIMESTAMPTZ ( 6 ),
                                      "deleted_at" TIMESTAMPTZ ( 6 ),
                                      "balance" NUMERIC,
                                      "pic" TEXT COLLATE "pg_catalog"."default",
                                      "name" VARCHAR ( 32 ) COLLATE "pg_catalog"."default" NOT NULL,
                                      "password" TEXT COLLATE "pg_catalog"."default" NOT NULL,
                                      "mobile" TEXT COLLATE "pg_catalog"."default",
                                      "id_card" TEXT COLLATE "pg_catalog"."default",
                                      "role" INT8 NOT NULL,
                                      "level" INT8 NOT NULL,
                                      "password_salt" TEXT COLLATE "pg_catalog"."default",
                                      "address" TEXT COLLATE "pg_catalog"."default",
                                      "longitude" VARCHAR ( 30 ) COLLATE "pg_catalog"."default",
                                      "description" VARCHAR ( 255 ) COLLATE "pg_catalog"."default",
                                      "latitude" VARCHAR ( 30 ) COLLATE "pg_catalog"."default",
                                      CONSTRAINT "factories_pkey" PRIMARY KEY ( "id" ),
                                      CONSTRAINT "factories_name_key" UNIQUE ( "name" )
);
ALTER TABLE "public"."factories" OWNER TO "root
";
COMMENT ON COLUMN "public"."factories"."balance" IS '账户余额';
COMMENT ON COLUMN "public"."factories"."pic" IS '用户头像';
COMMENT ON COLUMN "public"."factories"."name" IS '用户昵称';
COMMENT ON COLUMN "public"."factories"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."factories"."id_card" IS '身份证';
COMMENT ON COLUMN "public"."factories"."role" IS '身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 ';
COMMENT ON COLUMN "public"."factories"."level" IS '用户等级';
COMMENT ON COLUMN "public"."factories"."password_salt" IS '年销售量';
COMMENT ON COLUMN "public"."factories"."address" IS '场站位置信息';
COMMENT ON COLUMN "public"."factories"."longitude" IS '经度';
COMMENT ON COLUMN "public"."factories"."description" IS '场站描述';
COMMENT ON COLUMN "public"."factories"."latitude" IS '纬度';
-- ----------------------------
-- Table structure for factory_products
-- ----------------------------
DROP TABLE IF EXISTS "public"."factory_products";
CREATE TABLE "public"."factory_products" (
                                             "factory_name" TEXT COLLATE "pg_catalog"."default",
                                             "id" INT8 NOT NULL,
                                             "factory_id" INT8,
                                             "sku" TEXT COLLATE "pg_catalog"."default",
                                             "inventory" INT8,
                                             "name" TEXT COLLATE "pg_catalog"."default",
                                             "pic" TEXT COLLATE "pg_catalog"."default",
                                             "type" TEXT COLLATE "pg_catalog"."default",
                                             "buy_price" NUMERIC,
                                             "supply_price" NUMERIC,
                                             "daily_sales" NUMERIC,
                                             "weekly_sales" NUMERIC,
                                             "monthly_sales" NUMERIC,
                                             "annually_sales" NUMERIC,
                                             "total_sales" NUMERIC,
                                             "is_on_shelf" BOOL,
                                             CONSTRAINT "factory_products_pkey" PRIMARY KEY ( "id" ),
                                             CONSTRAINT "fk_factories_products" FOREIGN KEY ( "factory_id" ) REFERENCES "public"."factories" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."factory_products" OWNER TO "root
";
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
                                      "created_at" TIMESTAMPTZ ( 6 ),
                                      "updated_at" TIMESTAMPTZ ( 6 ),
                                      "deleted_at" TIMESTAMPTZ ( 6 ),
                                      "balance" NUMERIC,
                                      "pic" TEXT COLLATE "pg_catalog"."default",
                                      "name" VARCHAR ( 32 ) COLLATE "pg_catalog"."default" NOT NULL,
                                      "password" TEXT COLLATE "pg_catalog"."default" NOT NULL,
                                      "mobile" TEXT COLLATE "pg_catalog"."default",
                                      "id_card" TEXT COLLATE "pg_catalog"."default",
                                      "role" INT8 NOT NULL,
                                      "level" INT8 NOT NULL,
                                      "password_salt" TEXT COLLATE "pg_catalog"."default",
                                      CONSTRAINT "platforms_pkey" PRIMARY KEY ( "id" ),
                                      CONSTRAINT "platforms_name_key" UNIQUE ("name")
);
ALTER TABLE "public"."platforms" OWNER TO "root
";
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
                                    "created_at" TIMESTAMPTZ ( 6 ),
                                    "updated_at" TIMESTAMPTZ ( 6 ),
                                    "deleted_at" TIMESTAMPTZ ( 6 ),
                                    "balance" NUMERIC,
                                    "pic" TEXT COLLATE "pg_catalog"."default",
                                    "name" VARCHAR ( 32 ) COLLATE "pg_catalog"."default" NOT NULL,
                                    "password" TEXT COLLATE "pg_catalog"."default" NOT NULL,
                                    "mobile" TEXT COLLATE "pg_catalog"."default",
                                    "id_card" TEXT COLLATE "pg_catalog"."default",
                                    "role" INT8 NOT NULL,
                                    "level" INT8 NOT NULL,
                                    "car_id" TEXT COLLATE "pg_catalog"."default",
                                    "platform_id" INT8,
                                    "is_auth" BOOL,
                                    "address" TEXT COLLATE "pg_catalog"."default",
                                    "password_salt" TEXT COLLATE "pg_catalog"."default",
                                    "longitude" VARCHAR COLLATE "pg_catalog"."default",
                                    "latitude" VARCHAR COLLATE "pg_catalog"."default",
                                    CONSTRAINT "drivers_pkey" PRIMARY KEY ( "id" ),
                                    CONSTRAINT "fk_platforms_authorized_drivers" FOREIGN KEY ( "platform_id" ) REFERENCES "public"."platforms" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION,
                                    CONSTRAINT "idx_drivers_name" UNIQUE ( "name" )
);
ALTER TABLE "public"."drivers" OWNER TO "root
";
COMMENT ON COLUMN "public"."drivers"."balance" IS '账户余额';
COMMENT ON COLUMN "public"."drivers"."pic" IS '用户头像';
COMMENT ON COLUMN "public"."drivers"."name" IS '用户昵称';
COMMENT ON COLUMN "public"."drivers"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."drivers"."id_card" IS '身份证';
COMMENT ON COLUMN "public"."drivers"."role" IS '身份 0-乘客 1-司机 2-场站管理员 3-平台管理员 ';
COMMENT ON COLUMN "public"."drivers"."level" IS '用户等级';
COMMENT ON COLUMN "public"."drivers"."car_id" IS '车牌号';
COMMENT ON COLUMN "public"."drivers"."is_auth" IS '1为已认证，0为未认证';
COMMENT ON COLUMN "public"."drivers"."address" IS '地理位置';
COMMENT ON COLUMN "public"."drivers"."password_salt" IS '密码盐';
COMMENT ON COLUMN "public"."drivers"."longitude" IS '经度';
COMMENT ON COLUMN "public"."drivers"."latitude" IS '纬度';

-- ----------------------------
-- Table structure for order_products
-- ----------------------------
DROP TABLE IF EXISTS "public"."order_products";
CREATE TABLE "public"."order_products" (
                                           "cart_refer" INT8,
                                           "factory_id" INT8,
                                           "order_refer" TEXT COLLATE "pg_catalog"."default",
                                           "is_chosen" BOOL,
                                           "name" TEXT COLLATE "pg_catalog"."default" NOT NULL,
                                           "sku" TEXT COLLATE "pg_catalog"."default" ,
                                           "pic" TEXT COLLATE "pg_catalog"."default",
                                           "type" TEXT COLLATE "pg_catalog"."default" NOT NULL,
                                           "count" INT8 NOT NULL,
                                           "price" NUMERIC NOT NULL,
                                           CONSTRAINT "order_products_pkey" PRIMARY KEY ("type","name", "factory_id", "cart_refer")
);
ALTER TABLE "public"."order_products" OWNER TO "root
";
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
-- Table structure for devices
-- ----------------------------
DROP TABLE IF EXISTS "public"."devices";
CREATE TABLE "public"."devices" (
                                    "id" INT8 NOT NULL ,
                                    "owner_id" INT8,
                                    "platform_id" INT8,
                                    "is_activated" BOOL,
                                    "activated_time" TIMESTAMPTZ ( 6 ),
                                    "updated_time" TIMESTAMPTZ ( 6 ),
                                    "is_online" BOOL,
                                    "profit" NUMERIC,
                                    CONSTRAINT "devices_pkey" PRIMARY KEY ( "id" ),
                                    CONSTRAINT "fk_drivers_devices" FOREIGN KEY ( "owner_id" ) REFERENCES "public"."drivers" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION,
                                    CONSTRAINT "fk_platforms_devices" FOREIGN KEY ( "platform_id" ) REFERENCES "public"."platforms" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."devices" OWNER TO "root
";
COMMENT ON COLUMN "public"."devices"."owner_id" IS '车主ID';
COMMENT ON COLUMN "public"."devices"."is_activated" IS '1为激活，0为未激活';
COMMENT ON COLUMN "public"."devices"."activated_time" IS '激活时间';
COMMENT ON COLUMN "public"."devices"."updated_time" IS '更新时间';
COMMENT ON COLUMN "public"."devices"."is_online" IS '1为上线，0为未上线';
COMMENT ON COLUMN "public"."devices"."profit" IS '收益额';

-- ----------------------------
-- Table structure for driver_carts
-- ----------------------------
DROP TABLE IF EXISTS "public"."driver_carts";
CREATE TABLE "public"."driver_carts" (
                                         "driver_id" INT8,
                                         "cart_id" INT8 NOT NULL,
                                         "total_count" INT8,
                                         "total_amount" NUMERIC,
                                         CONSTRAINT "driver_carts_pkey" PRIMARY KEY ( "cart_id" ),
                                         CONSTRAINT "fk_drivers_cart" FOREIGN KEY ( "driver_id" ) REFERENCES "public"."drivers" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."driver_carts" OWNER TO "root
";
COMMENT ON COLUMN "public"."driver_carts"."total_count" IS '全选金额';
COMMENT ON COLUMN "public"."driver_carts"."total_amount" IS '全部商品数量';

-- ----------------------------
-- Table structure for driver_order_forms
-- ----------------------------
DROP TABLE IF EXISTS "public"."driver_order_forms";
CREATE TABLE "public"."driver_order_forms" (
                                               "factory_id" INT8,
                                               "factory_name" TEXT COLLATE "pg_catalog"."default",
                                               "driver_id" INT8,
                                               "car_id" TEXT COLLATE "pg_catalog"."default",
                                               "comment" TEXT COLLATE "pg_catalog"."default",
                                               "get_time" TIMESTAMPTZ ( 6 ),
                                               "order_id" TEXT COLLATE "pg_catalog"."default" NOT NULL,
                                               "cost" INT8,
                                               "state" INT2,
                                               "location" TEXT COLLATE "pg_catalog"."default",
                                               "placetime" TIMESTAMPTZ ( 6 ),
                                               "paytime" TIMESTAMPTZ ( 6 ),
                                               CONSTRAINT "driver_order_forms_pkey" PRIMARY KEY ( "order_id" ),
                                               CONSTRAINT "fk_drivers_driver_order_forms" FOREIGN KEY ( "driver_id" ) REFERENCES "public"."drivers" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION,
                                               CONSTRAINT "fk_factories_order_forms" FOREIGN KEY ( "factory_id" ) REFERENCES "public"."factories" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."driver_order_forms" OWNER TO "root
";
COMMENT ON COLUMN "public"."driver_order_forms"."factory_id" IS '指向factory.id';
COMMENT ON COLUMN "public"."driver_order_forms"."factory_name" IS '订单发货场站名';
COMMENT ON COLUMN "public"."driver_order_forms"."comment" IS '备注';
COMMENT ON COLUMN "public"."driver_order_forms"."get_time" IS '自取时间';
COMMENT ON COLUMN "public"."driver_order_forms"."cost" IS '花费';
COMMENT ON COLUMN "public"."driver_order_forms"."state" IS '订单状态 2-已完成 1-待取货 0-未支付';
COMMENT ON COLUMN "public"."driver_order_forms"."location" IS '支付时存储位置(购物时获取车主位置）';
COMMENT ON COLUMN "public"."driver_order_forms"."placetime" IS '下单时间';
COMMENT ON COLUMN "public"."driver_order_forms"."paytime" IS '支付时间';

-- ----------------------------
-- Table structure for device_products
-- ----------------------------
DROP TABLE IF EXISTS "public"."device_products";
CREATE TABLE "public"."device_products" (
                                            "id" INT8 NOT NULL,
                                            "factory_id" INT8,
                                            "driver_id" int8,
                                            "device_id" INT8,
                                            "sku" TEXT COLLATE "pg_catalog"."default",
                                            "inventory" INT8,
                                            "name" TEXT COLLATE "pg_catalog"."default",
                                            "pic" TEXT COLLATE "pg_catalog"."default",
                                            "type" TEXT COLLATE "pg_catalog"."default",
                                            "buy_price" NUMERIC,
                                            "supply_price" NUMERIC,
                                            "daily_sales" NUMERIC,
                                            "weekly_sales" NUMERIC,
                                            "monthly_sales" NUMERIC,
                                            "annually_sales" NUMERIC,
                                            "total_sales" NUMERIC,
                                            CONSTRAINT "device_products_pkey" PRIMARY KEY ( "id" ),
                                            CONSTRAINT "fk_devices_products" FOREIGN KEY ( "device_id" ) REFERENCES "public"."devices" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."device_products" OWNER TO "root
";
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

-- ----------------------------
-- Table structure for advertisements
-- ----------------------------
DROP TABLE IF EXISTS "public"."advertisements";
CREATE TABLE "public"."advertisements" (
                                           "id" INT8 NOT NULL ,
                                           "description" TEXT COLLATE "pg_catalog"."default",
                                           "platform_id" INT8,
                                           "expected_play_times" INT8,
                                           "play_times" INT8,
                                           "invest_fund" NUMERIC,
                                           "profit" NUMERIC,
                                           "ad_owner" TEXT COLLATE "pg_catalog"."default",
                                           "play_url" TEXT COLLATE "pg_catalog"."default",
                                           "expire_at" TIMESTAMPTZ ( 6 ),
                                           "ad_state" INT8,
                                           CONSTRAINT "advertisements_pkey" PRIMARY KEY ( "id" ),
                                           CONSTRAINT "fk_platforms_advertisements" FOREIGN KEY ( "platform_id" ) REFERENCES "public"."platforms" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."advertisements" OWNER TO "root
";
COMMENT ON COLUMN "public"."advertisements"."description" IS '广告描述';
COMMENT ON COLUMN "public"."advertisements"."expected_play_times" IS '预期播放次数';
COMMENT ON COLUMN "public"."advertisements"."play_times" IS '已经播放金额';
COMMENT ON COLUMN "public"."advertisements"."invest_fund" IS '投资金额';
COMMENT ON COLUMN "public"."advertisements"."profit" IS '产生收入';
COMMENT ON COLUMN "public"."advertisements"."ad_owner" IS '广告商';
COMMENT ON COLUMN "public"."advertisements"."play_url" IS '广告播放地址';
COMMENT ON COLUMN "public"."advertisements"."expire_at" IS '截止日期';
COMMENT ON COLUMN "public"."advertisements"."ad_state" IS '广告状态 1上线 ， 0下线';


-- ----------------------------
-- Table structure for bank_card_infos
-- ----------------------------
DROP TABLE IF EXISTS "public"."bank_card_infos";
CREATE TABLE "public"."bank_card_infos" (
                                            "id" INT8 NOT NULL,
                                            "card_id" INT8,
                                            "bank_name" TEXT COLLATE "pg_catalog"."default",
                                            "password" TEXT COLLATE "pg_catalog"."default",
                                            "cash" NUMERIC,
                                            "bank_fund" NUMERIC,
                                            CONSTRAINT "bank_card_infos_pkey" PRIMARY KEY ( "id" ),
                                            CONSTRAINT "bank_card_infos_card_id_key" UNIQUE ( "card_id" )
);
ALTER TABLE "public"."bank_card_infos" OWNER TO "root
";
COMMENT ON COLUMN "public"."bank_card_infos"."id" IS '用户ID';
COMMENT ON COLUMN "public"."bank_card_infos"."cash" IS '账户余额';
COMMENT ON COLUMN "public"."bank_card_infos"."bank_fund" IS '银行资金';
-- ----------------------------
-- Table structure for ad_devices
-- ----------------------------
DROP TABLE IF EXISTS "public"."ad_devices";
CREATE TABLE "public"."ad_devices" (
                                       "advertisement_id" INT8 NOT NULL,
                                       "device_id" INT8 NOT NULL,
                                       "play_times" INT8 NOT NULL,
                                       "profit" NUMERIC NOT NULL,
                                       CONSTRAINT "ad_devices_pkey" PRIMARY KEY ( "advertisement_id", "device_id" ),
                                       CONSTRAINT "fk_ad_devices_advertisement" FOREIGN KEY ( "advertisement_id" ) REFERENCES "public"."advertisements" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION,
                                       CONSTRAINT "fk_ad_devices_device" FOREIGN KEY ( "device_id" ) REFERENCES "public"."devices" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION
);
ALTER TABLE "public"."ad_devices" OWNER TO "root
";
