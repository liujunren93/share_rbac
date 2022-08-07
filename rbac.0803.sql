/*
 Navicat Premium Data Transfer

 Source Server         : node1
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : node1:3306
 Source Schema         : rbac

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 03/08/2022 10:11:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 2 COMMENT '\'1:已发布,2:未发布\'',
  `namespace` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'命名空间\'',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'配置名\'',
  `group` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'分组:debug',
  `content` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'内容\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_config_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of config
-- ----------------------------

-- ----------------------------
-- Table structure for rbac_admin
-- ----------------------------
DROP TABLE IF EXISTS `rbac_admin`;
CREATE TABLE `rbac_admin`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `domain_id` bigint NOT NULL DEFAULT 0 COMMENT '\'domain_id,omitempty\'',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'\'',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'\'',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'\'',
  `status` bigint NOT NULL DEFAULT 1 COMMENT '\'1:启用,2:禁用\'',
  `pl` bigint NOT NULL DEFAULT 0 COMMENT '\'permission level\'',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_rbac_admin_account`(`account`) USING BTREE,
  INDEX `idx_rbac_admin_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of rbac_admin
-- ----------------------------
INSERT INTO `rbac_admin` VALUES (1, 1, '2022-04-16 21:09:58', '2022-08-02 14:28:08', NULL, 'admin', 'admin', '$2a$10$viSK8m57ZvitkXSmVnG8xeggv7g5kt49sfjg73p1NXxkvz9Zcb0gW', 1, 1);
INSERT INTO `rbac_admin` VALUES (2, 1, '2022-06-06 01:11:54', '2022-08-02 14:24:45', NULL, 'test', 'qqq', '$2a$10$8CbKPZeYQvE.MYidauD/EOxIJliEaV7ILCobPK2stH0o9iQmrn4v2', 1, 2);

-- ----------------------------
-- Table structure for rbac_domain
-- ----------------------------
DROP TABLE IF EXISTS `rbac_domain`;
CREATE TABLE `rbac_domain`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '\'1:启用,2:禁用\'',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'\'',
  `domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'域名\'',
  `domain_id` bigint NOT NULL DEFAULT 0 COMMENT '\'\'',
  `pl` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'level_uid,权限锁管理员只能操作数据pl>管理员pl的数据 和自己锁上的数据\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_rbac_domain_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of rbac_domain
-- ----------------------------
INSERT INTO `rbac_domain` VALUES (1, '2022-04-05 20:29:06', '2022-04-09 22:24:48', NULL, 1, 'sharelife', 'sharelife1', 0, '');
INSERT INTO `rbac_domain` VALUES (2, '2022-04-10 00:33:59', '2022-05-11 23:15:20', NULL, 1, 'sharelife', 'sharelife', 0, '');
INSERT INTO `rbac_domain` VALUES (3, '2022-05-11 23:19:26', '2022-05-11 23:19:26', NULL, 1, 'test', 'test', 0, '');

-- ----------------------------
-- Table structure for rbac_path
-- ----------------------------
DROP TABLE IF EXISTS `rbac_path`;
CREATE TABLE `rbac_path`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'\'',
  `key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'前端权限组\'',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'name\'',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'component\'',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'redirect\'',
  `parent_id` bigint NOT NULL DEFAULT 0 COMMENT '\'parent_id\'',
  `api_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'api:api_path\'',
  `is_menu` tinyint(1) NOT NULL DEFAULT 1 COMMENT '\'1:生成菜单\'',
  `meta` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'meta{permission:}\'',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'path\'',
  `path_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '\'1:前端',
  `method` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'api:method\'',
  `status` bigint NOT NULL DEFAULT 1 COMMENT '\'1:启用,2:禁用\'',
  `target` tinyint(1) NOT NULL DEFAULT 1 COMMENT '\'1:_self',
  `domain_id` bigint NOT NULL DEFAULT -1 COMMENT '\'-1:公共\'',
  `action` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'方法:前端权限标识\'',
  `sort` bigint NOT NULL DEFAULT 9999,
  `pl` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'level_uid,权限锁管理员只能操作数据pl>管理员pl的数据 和自己锁上的数据\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_rbac_path_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_rbac_path_domain_id`(`domain_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 45 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of rbac_path
-- ----------------------------
INSERT INTO `rbac_path` VALUES (1, '2022-04-10 23:49:13', '2022-04-10 23:49:13', '2022-04-11 00:35:48', '', '9bELSAPeu9zLPCgnTTCroz2HMuOvsbhf', '主页', 'component', '/redirect', 0, '', 1, '', 'home', 1, '', 1, 1, -1, '', 9999, '');
INSERT INTO `rbac_path` VALUES (2, '2022-04-11 00:02:00', '2022-08-02 10:51:22', NULL, 'dashboard', '/rbac/rolePermission', 'dashboard', 'RouteView', '', 0, 'dashboard', 1, '{\"icon\":\"dashboard\",\"title\":\"dashboard\"}', '/dashboard', 1, 'GET', 1, 1, 1, 'query', 1, '');
INSERT INTO `rbac_path` VALUES (3, '2022-04-11 00:30:36', '2022-04-28 15:40:43', '2022-08-02 10:51:57', 'home', 'qA78petxYwXrUYof5TtNabSJdOn4VIIq', '主页12', 'component', '/redirect', 0, '', 1, '{\"title\":\"主页12\"}', 'home2', 1, '', 1, 1, 1, '', 9999, '');
INSERT INTO `rbac_path` VALUES (4, '2022-04-11 01:07:17', '2022-04-28 15:42:31', '2022-08-02 10:52:56', 'home', 'SMcamylRQt0nCFBppm5SkrX8TJOwC8XN', 'dingdan ', 'component', '/redirect', 0, '', 1, '{\"title\":\"dingdan \"}', 'home3', 1, '', 1, 1, 1, '', 9999, '');
INSERT INTO `rbac_path` VALUES (5, '2022-04-12 00:52:20', '2022-04-12 00:52:20', '2022-08-02 10:53:02', 'home2', '56feAOwv9LjYCeZ9WOITtraglOmPWmcS', 'dingdan2', 'component', '/redirect', 3, '', 1, '', 'home4', 1, '', 1, 1, 1, '', 9999, '');
INSERT INTO `rbac_path` VALUES (6, '2022-04-12 00:52:28', '2022-04-28 15:44:18', '2022-08-02 10:53:06', 'home33', 'XB6wQpdsHaPadJse7chBgTfYRaAnoeNX', 'dingdan33', 'component', '/redirect', 3, '', 1, '{\"icon\":\"test\",\"title\":\"dingdan33\"}', 'home5', 1, '', 1, 1, 1, '', 9999, '');
INSERT INTO `rbac_path` VALUES (7, '2022-04-12 00:52:52', '2022-04-28 15:42:45', '2022-08-02 10:53:10', 'home33', 'Kz2whCJLVusMbOgBH7PHliS3eN3FnMEo', 'dingdan33', 'component', '/redirect', 6, '', 1, '{\"title\":\"dingdan33\"}', 'home6', 1, '', 1, 1, 1, '', 9999, '');
INSERT INTO `rbac_path` VALUES (8, '2022-04-19 23:42:28', '2022-04-29 01:00:27', NULL, 'rbac', '2ucgA2qbQ75heFXSFU7yhaCA7Wqg0u97', '权限控制', 'RouteView', '/rbac/path', 0, '', 1, '{\"icon\":\"team\",\"title\":\"权限控制\"}', '/rbac', 1, '', 1, 1, -1, '', 9999, '');
INSERT INTO `rbac_path` VALUES (9, '2022-04-19 23:47:12', '2022-07-15 11:56:06', NULL, 'path列表', '/rbac/path/', '路由管理', 'rbac/path/List', '', 8, '/rbac/path', 1, '{\"hideChildrenInMenu\":\"true\",\"title\":\"路由管理\"}', '/rbac/path/list', 1, 'GET', 1, 1, 1, 'query', 9999, '');
INSERT INTO `rbac_path` VALUES (10, '2022-04-22 00:51:55', '2022-07-15 11:57:11', NULL, 'path管理', '/rbac/path', '路由设置', 'rbac/path/Edit', '', 8, '/rbac/path/:id', 1, '{\"hideChildren\":\"true\",\"icon\":\"dashboard\",\"show\":\"false\",\"title\":\"路由设置\"}', '/rbac/path/edit', 1, 'PUT', 1, 1, 1, 'edit', 9999, '');
INSERT INTO `rbac_path` VALUES (11, '2022-04-23 17:30:41', '2022-05-31 16:54:54', NULL, 'role列表', '/rbac/role', '角色管理', 'rbac/role/List', '', 8, '/rbac/role', 1, '{\"title\":\"角色管理\"}', '/rbac/role/list', 1, 'GET', 1, 1, 1, 'query', 9999, '');
INSERT INTO `rbac_path` VALUES (12, '2022-04-27 18:24:57', '2022-06-02 14:59:05', NULL, 'tttt', '/rbac/permission', '权限管理', 'rbac/permission/List', '', 8, '/rbac/permission', 1, '{\"title\":\"权限管理\"}', '/rbac/permission/list', 1, 'GET', 1, 1, -1, 'query', 9999, '');
INSERT INTO `rbac_path` VALUES (13, '2022-04-27 18:54:26', '2022-06-02 14:59:26', NULL, 'ttt', '/rbac/permission', '权限删除', '', '', 12, '/rbac/permission/:id', 1, '{\"icon\":\"true\",\"keepAlive\":\"true\",\"show\":\"false\",\"title\":\"权限编辑\"}', '', 2, 'DELETE', 1, 1, -1, 'delete', 9999, '');
INSERT INTO `rbac_path` VALUES (14, '2022-05-11 00:13:27', '2022-06-04 23:17:34', NULL, '', '/rbac/domain', '域管理', 'rbac/domain/List', '', 8, '/rbac/domain', 1, '{\"title\":\"域管理\"}', '/rbac/domain/list', 1, 'GET', 1, 1, 1, 'query', 9999, '');
INSERT INTO `rbac_path` VALUES (15, '2022-05-13 00:08:47', '2022-05-15 23:54:38', NULL, '', '/rbac/path', '路由详情', '', '', 9, '/rbac/path/:id', 1, '', '', 2, 'GET', 1, 1, -1, 'info', 9999, '');
INSERT INTO `rbac_path` VALUES (16, '2022-05-13 00:37:21', '2022-05-15 23:54:56', NULL, '', '/rbac/permission', '权限添加', '', '', 12, '/rbac/permission', 1, '', '', 2, 'POST', 1, 1, -1, 'add', 9999, '');
INSERT INTO `rbac_path` VALUES (17, '2022-05-13 00:37:50', '2022-05-15 23:55:07', NULL, '', '/rbac/permission', '权限修改', '', '', 12, '/rbac/permission/:id', 1, '', '', 2, 'PUT', 1, 1, -1, 'edit', 9999, '');
INSERT INTO `rbac_path` VALUES (18, '2022-05-13 00:38:07', '2022-05-15 23:55:16', NULL, '', '/rbac/permission', '权限删除', '', '', 12, '/rbac/permission/:id', 1, '', '', 2, 'DELETE', 1, 1, -1, 'delete', 9999, '');
INSERT INTO `rbac_path` VALUES (19, '2022-05-13 00:44:09', '2022-05-13 00:44:09', NULL, '', 'cuUkDIyiTukSaGNsrIxFNKy8l1PN4h94', '权限路由获取', '', '', 12, '/rbac/permissionPath/:id', 1, '', '', 2, 'GET', 1, 1, -1, '', 9999, '');
INSERT INTO `rbac_path` VALUES (20, '2022-05-13 00:44:18', '2022-05-13 00:44:18', NULL, '', 'SNUNcpI4cZWxk0ixPC1GSuzKppruAkRL', '权限路由设置', '', '', 12, '/rbac/permissionPath/:id', 1, '', '', 2, 'PUT', 1, 1, -1, '', 9999, '');
INSERT INTO `rbac_path` VALUES (21, '2022-05-13 00:58:58', '2022-05-15 23:55:28', NULL, '', '/rbac/role', '角色添加', '', '', 11, '/rbac/role', 1, '', '', 2, 'POST', 1, 1, -1, 'add', 9999, '');
INSERT INTO `rbac_path` VALUES (22, '2022-05-13 01:04:32', '2022-05-15 23:55:41', NULL, '', '/rbac/role', '角色删除', '', '', 11, '/rbac/role/:id', 1, '', '', 2, 'DELETE', 1, 1, -1, 'delete', 9999, '');
INSERT INTO `rbac_path` VALUES (23, '2022-05-13 01:04:46', '2022-05-16 00:11:59', NULL, '', '/rbac/role', '角色修改', '', '', 11, '/rbac/role/:id', 1, '', '', 2, 'PUT', 1, 1, -1, 'edit', 9999, '');
INSERT INTO `rbac_path` VALUES (24, '2022-05-13 01:05:21', '2022-05-15 23:56:03', NULL, '', '/rbac/domain', '域添加', '', '', 14, '/rbac/domain', 1, '', '', 2, 'POST', 1, 1, 1, 'add', 9999, '');
INSERT INTO `rbac_path` VALUES (25, '2022-05-13 01:05:43', '2022-05-15 23:56:25', NULL, '', '/rbac/domain', '域修改', '', '', 14, '/rbac/domain/:id', 1, '', '', 2, 'PUT', 1, 1, 1, 'edit', 9999, '');
INSERT INTO `rbac_path` VALUES (26, '2022-05-13 01:05:59', '2022-05-15 23:53:18', NULL, '', '/rbac/domain', '域删除', '', '', 14, '/rbac/domain/:id', 1, '', '', 2, 'DELETE', 1, 1, 1, 'delete', 9999, '');
INSERT INTO `rbac_path` VALUES (27, '2022-05-13 01:07:19', '2022-05-13 01:07:19', NULL, '', '/rbac/rolePermission', '角色权限列表', '', '', 12, '/rbac/rolePermission/:id', 1, '', '', 2, 'GET', 1, 1, -1, '', 9999, '');
INSERT INTO `rbac_path` VALUES (28, '2022-05-13 01:07:37', '2022-05-13 01:07:37', NULL, '', '/rbac/rolePermission', '角色权限设置', '', '', 12, '/rbac/rolePermission/:id', 1, '', '', 2, 'PUT', 1, 1, -1, '', 9999, '');
INSERT INTO `rbac_path` VALUES (29, '2022-05-16 00:34:21', '2022-05-16 00:35:05', NULL, '', '/rbac/path', '路由添加', '', '', 9, '/rbac/path', 1, '', '', 2, 'POST', 1, 1, 1, 'add', 9999, '');
INSERT INTO `rbac_path` VALUES (30, '2022-05-16 01:07:21', '2022-05-16 01:07:21', NULL, '', '/rbac/path', '路由删除', '', '', 9, '/rbac/path/:id', 1, '', '', 2, 'DELETE', 1, 1, 1, 'delete', 9999, '');
INSERT INTO `rbac_path` VALUES (31, '2022-06-04 23:44:37', '2022-06-04 23:44:37', NULL, '', '/rbac/admin', '管理员管理', 'rbac/user/List', '', 8, '/rbac/admin', 1, '{\"title\":\"管理员管理\"}', '/rbac/user/list', 1, 'GET', 1, 1, -1, 'query', 9999, '');
INSERT INTO `rbac_path` VALUES (32, '2022-06-05 01:13:00', '2022-06-05 01:13:00', NULL, '', '/rbac/adminRole', '管理员角色列表', '', '', 31, '/rbac/adminRole/:id', 1, '', '', 2, 'GET', 1, 1, -1, 'query', 9999, '');
INSERT INTO `rbac_path` VALUES (33, '2022-06-05 01:15:12', '2022-06-06 02:06:02', NULL, '', '/rbac/adminRole', '管理员角色设置', '', '', 31, '/rbac/adminRole/:id', 1, '', '', 2, 'PUT', 1, 1, -1, 'add', 9999, '');
INSERT INTO `rbac_path` VALUES (34, '2022-06-05 01:20:16', '2022-06-05 01:20:16', NULL, '', '/rbac/admin', '管理员添加', '', '', 31, '/rbac/admin', 1, '', '', 2, 'POST', 1, 1, -1, 'add', 9999, '');
INSERT INTO `rbac_path` VALUES (35, '2022-06-05 01:20:49', '2022-06-05 01:20:49', NULL, '', '/rbac/admin', '管理员修改', '', '', 31, '/rbac/admin/:id', 1, '', '', 2, 'PUT', 1, 1, -1, 'edit', 9999, '');
INSERT INTO `rbac_path` VALUES (36, '2022-06-05 01:21:08', '2022-06-05 01:21:08', NULL, '', '/rbac/admin', '管理员删除', '', '', 31, '/rbac/admin/:id', 1, '', '', 2, 'DELETE', 1, 1, -1, 'delete', 9999, '');
INSERT INTO `rbac_path` VALUES (37, '2022-07-14 00:07:16', '2022-07-14 23:58:21', NULL, '', '/cluster', '集群管理', 'RouteView', '', 0, '/cluster', 1, '{\"icon\":\"cluster\",\"title\":\"集群管理\"}', '/cluster', 1, 'GET', 1, 1, 1, 'query', 500, '');
INSERT INTO `rbac_path` VALUES (38, '2022-07-14 00:11:26', '2022-07-19 13:50:47', NULL, '', '/configCenter/Config', '配置中心', 'cluster/config/List', '', 37, '/configCenter/Config', 1, '{\"title\":\"配置中心\"}', '/cluster/configCenter', 1, 'GET', 1, 1, 1, 'query', 9999, '');
INSERT INTO `rbac_path` VALUES (39, '2022-07-19 13:50:27', '2022-07-19 13:51:15', NULL, '', '/configCenter/Config', '配置添加', '', '', 38, '/configCenter/Config', 1, '', '', 2, 'POST', 1, 1, 1, 'add', 9999, '');
INSERT INTO `rbac_path` VALUES (40, '2022-07-19 13:55:54', '2022-07-20 21:52:59', NULL, '', '/configCenter/Config', '配置修改', '', '', 38, '/configCenter/Config/:id', 1, '', '', 2, 'PUT', 1, 1, 1, 'edit', 9999, '');
INSERT INTO `rbac_path` VALUES (41, '2022-07-19 13:57:44', '2022-07-21 01:35:19', NULL, '', '/configCenter/Config', '配置删除', '', '', 38, '/configCenter/Config/:id', 1, '', '', 2, 'DELETE', 1, 1, 1, 'delete', 9999, '');
INSERT INTO `rbac_path` VALUES (42, '2022-07-19 14:54:17', '2022-07-21 00:52:29', NULL, '', '/configCenter/Config/Publish', '发布配置', '', '', 38, '/configCenter/Config/Publish', 1, '', '', 2, 'POST', 1, 1, 1, 'publish', 9999, '');
INSERT INTO `rbac_path` VALUES (43, '2022-07-19 15:06:35', '2022-07-19 15:07:19', NULL, '', '/configCenter/Config/SetStatus', '设置配置状态', '', '', 38, '/configCenter/Config/SetStatus', 1, '', '', 2, 'PUT', 1, 1, 1, 'SetStatus', 9999, '');
INSERT INTO `rbac_path` VALUES (44, '2022-07-20 00:36:25', '2022-07-20 00:40:20', NULL, '', '/configCenter/Config', '配置详情', '', '', 38, '/configCenter/Config/:id', 1, '', '', 2, 'GET', 1, 1, 1, 'info', 9999, '');

-- ----------------------------
-- Table structure for rbac_permission
-- ----------------------------
DROP TABLE IF EXISTS `rbac_permission`;
CREATE TABLE `rbac_permission`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'name\'',
  `desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'desc\'',
  `domain_id` bigint NOT NULL DEFAULT 0 COMMENT '\'-1:基础权限，所有域通用\'',
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '\'1:普通权限,2:action',
  `status` bigint NOT NULL DEFAULT 1 COMMENT '\'1:启用,2:禁用\'',
  `pl` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'level_uid,权限锁管理员只能操作数据pl>管理员pl的数据 和自己锁上的数据\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_rbac_permission_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_rbac_permission_domain_id`(`domain_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of rbac_permission
-- ----------------------------
INSERT INTO `rbac_permission` VALUES (1, '2022-04-10 18:52:16', '2022-04-10 23:14:56', '2022-04-10 23:28:24', 'fdsaf', 'desc1', 1, 2, 1, '');
INSERT INTO `rbac_permission` VALUES (2, '2022-04-10 23:15:29', '2022-08-02 09:52:06', NULL, '平台超级所有权限', '平台超级所有权限', 1, 3, 1, '0_1');
INSERT INTO `rbac_permission` VALUES (3, '2022-05-10 17:45:48', '2022-08-02 09:46:58', NULL, '用户超级管理员权限', '用户超级管理员权限', 1, 1, 1, '0_1');

-- ----------------------------
-- Table structure for rbac_permission_path
-- ----------------------------
DROP TABLE IF EXISTS `rbac_permission_path`;
CREATE TABLE `rbac_permission_path`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `permission_id` bigint NOT NULL DEFAULT 0 COMMENT '\'permission_id\'',
  `path_id` bigint NOT NULL DEFAULT 0 COMMENT '\'path_id\'',
  `domain_id` bigint NOT NULL DEFAULT 0 COMMENT '\'domain_id\'',
  `pl` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'level_uid,权限锁管理员只能操作数据pl>管理员pl的数据 和自己锁上的数据\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_rbac_permission_path_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 63 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of rbac_permission_path
-- ----------------------------
INSERT INTO `rbac_permission_path` VALUES (1, '2022-04-11 00:56:37', '2022-04-11 00:56:37', NULL, 1, 1, 0, '');
INSERT INTO `rbac_permission_path` VALUES (2, '2022-04-11 00:56:37', '2022-04-11 00:56:37', NULL, 1, 2, 0, '');
INSERT INTO `rbac_permission_path` VALUES (3, '2022-04-11 00:56:37', '2022-04-11 00:56:37', '2022-08-02 10:51:57', 1, 3, 0, '');
INSERT INTO `rbac_permission_path` VALUES (4, '2022-04-11 01:08:12', '2022-04-11 01:08:12', '2022-08-02 10:52:56', 1, 4, 0, '');
INSERT INTO `rbac_permission_path` VALUES (5, '2022-04-12 00:53:40', '2022-04-12 00:53:40', '2022-05-04 18:41:24', 2, 1, 0, '');
INSERT INTO `rbac_permission_path` VALUES (6, '2022-04-12 00:53:40', '2022-04-12 00:53:40', '2022-05-04 18:41:24', 2, 2, 0, '');
INSERT INTO `rbac_permission_path` VALUES (7, '2022-04-12 00:53:40', '2022-04-12 00:53:40', '2022-08-02 10:51:57', 2, 3, 0, '');
INSERT INTO `rbac_permission_path` VALUES (8, '2022-04-12 00:53:40', '2022-04-12 00:53:40', '2022-08-02 10:52:56', 2, 4, 0, '');
INSERT INTO `rbac_permission_path` VALUES (9, '2022-04-12 00:53:40', '2022-04-12 00:53:40', '2022-08-02 10:53:02', 2, 5, 0, '');
INSERT INTO `rbac_permission_path` VALUES (10, '2022-04-12 00:53:40', '2022-04-12 00:53:40', '2022-08-02 10:53:06', 2, 6, 0, '');
INSERT INTO `rbac_permission_path` VALUES (11, '2022-04-12 00:53:40', '2022-04-12 00:53:40', '2022-08-02 10:53:10', 2, 7, 0, '');
INSERT INTO `rbac_permission_path` VALUES (12, '2022-04-19 23:43:37', '2022-04-19 23:43:37', NULL, 2, 8, 0, '');
INSERT INTO `rbac_permission_path` VALUES (13, '2022-04-19 23:50:42', '2022-04-19 23:50:42', NULL, 2, 9, 0, '');
INSERT INTO `rbac_permission_path` VALUES (14, '2022-04-22 00:52:17', '2022-04-22 00:52:17', NULL, 2, 10, 0, '');
INSERT INTO `rbac_permission_path` VALUES (15, '2022-04-23 21:03:34', '2022-04-23 21:03:34', NULL, 2, 11, 0, '');
INSERT INTO `rbac_permission_path` VALUES (16, '2022-04-29 00:51:24', '2022-04-29 00:51:24', NULL, 2, 12, 0, '');
INSERT INTO `rbac_permission_path` VALUES (17, '2022-04-29 00:51:24', '2022-04-29 00:51:24', NULL, 2, 13, 0, '');
INSERT INTO `rbac_permission_path` VALUES (18, '2022-05-10 16:52:05', '2022-05-10 16:52:05', NULL, 2, 2, 0, '');
INSERT INTO `rbac_permission_path` VALUES (19, '2022-05-10 17:45:48', '2022-05-10 17:45:48', NULL, 0, 2, 0, '');
INSERT INTO `rbac_permission_path` VALUES (20, '2022-05-10 17:52:18', '2022-05-10 17:52:18', NULL, 3, 2, 0, '');
INSERT INTO `rbac_permission_path` VALUES (21, '2022-05-10 17:52:18', '2022-05-10 17:52:18', '2022-08-02 10:53:02', 3, 5, 0, '');
INSERT INTO `rbac_permission_path` VALUES (22, '2022-05-10 17:52:18', '2022-05-10 17:52:18', '2022-08-02 10:53:10', 3, 7, 0, '');
INSERT INTO `rbac_permission_path` VALUES (23, '2022-05-10 17:52:18', '2022-05-10 17:52:18', '2022-08-02 10:53:06', 3, 6, 0, '');
INSERT INTO `rbac_permission_path` VALUES (24, '2022-05-10 17:52:18', '2022-05-10 17:52:18', '2022-08-02 10:51:57', 3, 3, 0, '');
INSERT INTO `rbac_permission_path` VALUES (25, '2022-05-11 22:21:30', '2022-05-11 22:21:30', NULL, 2, 14, 0, '');
INSERT INTO `rbac_permission_path` VALUES (26, '2022-05-13 00:34:09', '2022-05-13 00:34:09', NULL, 2, 15, 0, '');
INSERT INTO `rbac_permission_path` VALUES (27, '2022-05-13 00:44:50', '2022-05-13 00:44:50', NULL, 2, 16, 0, '');
INSERT INTO `rbac_permission_path` VALUES (28, '2022-05-13 00:44:50', '2022-05-13 00:44:50', NULL, 2, 17, 0, '');
INSERT INTO `rbac_permission_path` VALUES (29, '2022-05-13 00:44:50', '2022-05-13 00:44:50', NULL, 2, 18, 0, '');
INSERT INTO `rbac_permission_path` VALUES (30, '2022-05-13 00:44:50', '2022-05-13 00:44:50', NULL, 2, 19, 0, '');
INSERT INTO `rbac_permission_path` VALUES (31, '2022-05-13 00:44:50', '2022-05-13 00:44:50', NULL, 2, 20, 0, '');
INSERT INTO `rbac_permission_path` VALUES (32, '2022-05-15 21:02:59', '2022-05-15 21:02:59', NULL, 2, 24, 0, '');
INSERT INTO `rbac_permission_path` VALUES (33, '2022-05-15 21:02:59', '2022-05-15 21:02:59', NULL, 2, 25, 0, '');
INSERT INTO `rbac_permission_path` VALUES (34, '2022-05-15 21:02:59', '2022-05-15 21:02:59', NULL, 2, 26, 0, '');
INSERT INTO `rbac_permission_path` VALUES (35, '2022-05-15 21:02:59', '2022-05-15 21:02:59', NULL, 2, 21, 0, '');
INSERT INTO `rbac_permission_path` VALUES (36, '2022-05-15 21:02:59', '2022-05-15 21:02:59', NULL, 2, 22, 0, '');
INSERT INTO `rbac_permission_path` VALUES (37, '2022-05-15 21:02:59', '2022-05-15 21:02:59', NULL, 2, 23, 0, '');
INSERT INTO `rbac_permission_path` VALUES (38, '2022-05-16 00:46:51', '2022-05-16 00:46:51', NULL, 2, 29, 0, '');
INSERT INTO `rbac_permission_path` VALUES (39, '2022-05-16 01:07:45', '2022-05-16 01:07:45', NULL, 2, 30, 0, '');
INSERT INTO `rbac_permission_path` VALUES (40, '2022-06-02 15:22:11', '2022-06-02 15:22:11', NULL, 2, 27, 0, '');
INSERT INTO `rbac_permission_path` VALUES (41, '2022-06-02 15:22:11', '2022-06-02 15:22:11', NULL, 2, 28, 0, '');
INSERT INTO `rbac_permission_path` VALUES (42, '2022-06-04 23:49:02', '2022-06-04 23:49:02', NULL, 2, 31, 0, '');
INSERT INTO `rbac_permission_path` VALUES (43, '2022-06-05 01:16:40', '2022-06-05 01:16:40', NULL, 2, 32, 0, '');
INSERT INTO `rbac_permission_path` VALUES (44, '2022-06-05 01:16:40', '2022-06-05 01:16:40', NULL, 2, 33, 0, '');
INSERT INTO `rbac_permission_path` VALUES (45, '2022-06-05 01:21:39', '2022-06-05 01:21:39', NULL, 2, 34, 0, '');
INSERT INTO `rbac_permission_path` VALUES (46, '2022-06-05 01:21:39', '2022-06-05 01:21:39', NULL, 2, 35, 0, '');
INSERT INTO `rbac_permission_path` VALUES (47, '2022-06-05 01:21:39', '2022-06-05 01:21:39', NULL, 2, 36, 0, '');
INSERT INTO `rbac_permission_path` VALUES (48, '2022-07-14 00:07:37', '2022-07-14 00:07:37', NULL, 2, 37, 0, '');
INSERT INTO `rbac_permission_path` VALUES (49, '2022-07-14 00:11:52', '2022-07-14 00:11:52', NULL, 2, 38, 0, '');
INSERT INTO `rbac_permission_path` VALUES (50, '2022-07-19 13:52:00', '2022-07-19 13:52:00', '2022-07-19 14:04:13', 2, 39, 0, '');
INSERT INTO `rbac_permission_path` VALUES (51, '2022-07-19 13:59:41', '2022-07-19 13:59:41', NULL, 2, 40, 0, '');
INSERT INTO `rbac_permission_path` VALUES (52, '2022-07-19 13:59:41', '2022-07-19 13:59:41', NULL, 2, 41, 0, '');
INSERT INTO `rbac_permission_path` VALUES (53, '2022-07-19 14:04:24', '2022-07-19 14:04:24', '2022-07-19 15:22:58', 2, 39, 0, '');
INSERT INTO `rbac_permission_path` VALUES (54, '2022-07-19 15:10:06', '2022-07-19 15:10:06', NULL, 2, 42, 0, '');
INSERT INTO `rbac_permission_path` VALUES (55, '2022-07-19 15:10:06', '2022-07-19 15:10:06', '2022-07-19 15:22:20', 2, 43, 0, '');
INSERT INTO `rbac_permission_path` VALUES (56, '2022-07-19 15:23:25', '2022-07-19 15:23:25', NULL, 2, 39, 0, '');
INSERT INTO `rbac_permission_path` VALUES (57, '2022-07-19 15:23:25', '2022-07-19 15:23:25', NULL, 2, 43, 0, '');
INSERT INTO `rbac_permission_path` VALUES (58, '2022-07-20 00:36:38', '2022-07-20 00:36:38', '2022-07-20 00:55:36', 2, 44, 0, '');
INSERT INTO `rbac_permission_path` VALUES (59, '2022-07-20 00:56:41', '2022-07-20 00:56:41', '2022-07-20 01:04:58', 2, 44, 0, '');
INSERT INTO `rbac_permission_path` VALUES (60, '2022-07-20 01:07:45', '2022-07-20 01:07:45', '2022-07-20 01:44:44', 2, 44, 0, '');
INSERT INTO `rbac_permission_path` VALUES (61, '2022-07-20 01:47:43', '2022-07-20 01:47:43', '2022-07-20 01:49:42', 2, 44, 0, '');
INSERT INTO `rbac_permission_path` VALUES (62, '2022-07-20 01:49:58', '2022-07-20 01:49:58', NULL, 2, 44, 0, '');

-- ----------------------------
-- Table structure for rbac_role
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role`;
CREATE TABLE `rbac_role`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'name\'',
  `desc` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'desc\'',
  `domain_id` bigint NOT NULL DEFAULT 0 COMMENT '\'domain_id\'',
  `status` bigint NOT NULL DEFAULT 1 COMMENT '\'1:启用,2:禁用\'',
  `pl` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'level_uid,权限锁管理员只能操作数据pl>管理员pl的数据 和自己锁上的数据\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_rbac_role_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of rbac_role
-- ----------------------------
INSERT INTO `rbac_role` VALUES (1, '2022-04-11 01:48:06', '2022-06-15 02:09:06', NULL, 'root', 'root', 1, 1, '');
INSERT INTO `rbac_role` VALUES (2, '2022-04-11 01:54:53', '2022-04-11 01:54:53', NULL, 'role', 'role', 1, 1, '');
INSERT INTO `rbac_role` VALUES (3, '2022-04-28 19:47:28', '2022-05-10 17:53:03', NULL, 'test1', 'test', 1, 1, '');

-- ----------------------------
-- Table structure for rbac_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role_permission`;
CREATE TABLE `rbac_role_permission`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `role_id` bigint NOT NULL DEFAULT 0 COMMENT '\'role_id\'',
  `permission_id` bigint NOT NULL DEFAULT 0 COMMENT '\'permission_id\'',
  `status` bigint NOT NULL DEFAULT 1 COMMENT '\'1:启用,2:禁用\'',
  `domain_id` bigint NOT NULL DEFAULT 0 COMMENT '\'domain_id\'',
  `pl` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'level_uid,权限锁管理员只能操作数据pl>管理员pl的数据 和自己锁上的数据\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_rbac_role_permission_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of rbac_role_permission
-- ----------------------------
INSERT INTO `rbac_role_permission` VALUES (1, '2022-04-11 02:07:55', '2022-04-11 02:07:55', '2022-06-15 01:48:52', 1, 1, 1, 0, '');
INSERT INTO `rbac_role_permission` VALUES (2, '2022-04-11 02:07:55', '2022-04-11 02:07:55', NULL, 1, 2, 1, 0, '');
INSERT INTO `rbac_role_permission` VALUES (3, '2022-04-11 02:12:51', '2022-04-11 02:12:51', '2022-06-15 01:48:52', 1, 1, 1, 1, '');
INSERT INTO `rbac_role_permission` VALUES (4, '2022-04-11 02:12:51', '2022-04-11 02:12:51', NULL, 1, 2, 1, 1, '');
INSERT INTO `rbac_role_permission` VALUES (5, '2022-05-09 23:29:21', '2022-05-09 23:29:21', '2022-05-09 23:30:06', 3, 2, 1, 1, '');
INSERT INTO `rbac_role_permission` VALUES (6, '2022-05-09 23:32:44', '2022-05-09 23:32:44', NULL, 3, 2, 1, 1, '');
INSERT INTO `rbac_role_permission` VALUES (7, '2022-05-10 17:53:04', '2022-05-10 17:53:04', NULL, 3, 3, 1, 1, '');

-- ----------------------------
-- Table structure for rbac_role_user
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role_user`;
CREATE TABLE `rbac_role_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `domain_id` bigint NOT NULL DEFAULT 0 COMMENT '\'domain_id\'',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `role_id` bigint NOT NULL DEFAULT 0 COMMENT '\'role_id\'',
  `uid` bigint NOT NULL DEFAULT 0 COMMENT '\'uid\'',
  `status` bigint NOT NULL DEFAULT 1 COMMENT '\'1:启用,2:禁用\'',
  `pl` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '\'level_uid,权限锁管理员只能操作数据pl>管理员pl的数据 和自己锁上的数据\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_rbac_role_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of rbac_role_user
-- ----------------------------
INSERT INTO `rbac_role_user` VALUES (1, 1, '2022-04-11 23:17:01', '2022-04-11 23:17:01', '2022-04-11 23:55:10', 1, 1, 1, '');
INSERT INTO `rbac_role_user` VALUES (2, 1, '2022-04-11 23:17:01', '2022-04-11 23:17:01', '2022-04-11 23:33:21', 2, 1, 1, '');
INSERT INTO `rbac_role_user` VALUES (3, 1, '2022-04-11 23:34:01', '2022-04-11 23:34:01', '2022-04-11 23:55:10', 2, 1, 1, '');
INSERT INTO `rbac_role_user` VALUES (4, 1, '2022-04-11 23:55:51', '2022-04-11 23:55:51', NULL, 1, 1, 1, '');
INSERT INTO `rbac_role_user` VALUES (5, 1, '2022-04-11 23:55:51', '2022-04-11 23:55:51', NULL, 2, 1, 1, '');
INSERT INTO `rbac_role_user` VALUES (6, 1, '2022-06-06 02:09:03', '2022-06-06 02:09:03', NULL, 1, 2, 1, '');

SET FOREIGN_KEY_CHECKS = 1;
