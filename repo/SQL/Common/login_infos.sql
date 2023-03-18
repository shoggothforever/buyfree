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

 Date: 18/03/2023 14:48:46
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
