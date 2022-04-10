/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50734
 Source Host           : localhost:3306
 Source Schema         : game_sdk

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 10/04/2022 21:38:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for active_records
-- ----------------------------
DROP TABLE IF EXISTS `active_records`;
CREATE TABLE `active_records` (
  `aId` int(10) NOT NULL AUTO_INCREMENT,
  `activeTime` int(10) unsigned NOT NULL COMMENT '点击时间',
  `adId` int(10) NOT NULL DEFAULT '0' COMMENT '广告ID',
  `bId` int(10) NOT NULL DEFAULT '0' COMMENT '渠道包Id',
  `channelId` int(10) NOT NULL DEFAULT '0' COMMENT '渠道Id,包Id',
  `sonChannel` int(10) NOT NULL DEFAULT '0' COMMENT '子渠道Id',
  `gameId` int(10) NOT NULL COMMENT '投放的游戏id',
  `downIp` bigint(10) NOT NULL COMMENT '点击者ip',
  `imeiCode` varchar(64) DEFAULT NULL COMMENT '设备码',
  `isnewImei` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1为新设备',
  PRIMARY KEY (`aId`),
  KEY `imeiCode` (`imeiCode`),
  KEY `activeTime` (`activeTime`),
  KEY `downIp` (`downIp`),
  KEY `gameId` (`gameId`),
  KEY `channelId` (`channelId`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8 COMMENT='近7天游戏激活录入表';

-- ----------------------------
-- Records of active_records
-- ----------------------------
BEGIN;
INSERT INTO `active_records` VALUES (33, 1649597821, 0, 0, 0, 0, 1000000, 123, '12-34-56-78-9100', 1);
COMMIT;

-- ----------------------------
-- Table structure for games
-- ----------------------------
DROP TABLE IF EXISTS `games`;
CREATE TABLE `games` (
  `gameId` bigint(12) NOT NULL COMMENT '游戏Id',
  `gameName` varchar(30) NOT NULL COMMENT '游戏名称',
  `primitiveName` varchar(30) NOT NULL COMMENT '原名游戏',
  `gameShort` varchar(30) DEFAULT NULL COMMENT '游戏名缩写',
  `cpId` varchar(50) NOT NULL COMMENT '合作方Id(厂商)',
  `mainGameId` int(10) NOT NULL DEFAULT '0' COMMENT '同款游戏（标记是同款游戏）',
  `topGameId` int(10) NOT NULL COMMENT '最顶级归属ID',
  `cpLoginKey` varchar(32) NOT NULL COMMENT '合作方登录秘钥',
  `cpPayKey` varchar(32) NOT NULL COMMENT '合作方支付key',
  `cpGameId` int(10) NOT NULL DEFAULT '0' COMMENT 'Cp方的游戏id',
  `payPoint` int(10) NOT NULL DEFAULT '6' COMMENT '游戏支付点，默认6块',
  `payMoneyStr` varchar(255) NOT NULL DEFAULT '0' COMMENT '可充值面额',
  `goldFromRMB` smallint(5) NOT NULL DEFAULT '10' COMMENT '游戏币比率',
  `goldName` varchar(10) NOT NULL COMMENT '游戏币单位',
  `gameUrl` varchar(200) DEFAULT NULL COMMENT '游戏地址',
  `payNotifyUrl` varchar(200) NOT NULL COMMENT '游戏充值回调地址',
  `gameRoleApi` varchar(200) DEFAULT NULL COMMENT '玩家游戏信息查询api，如角色名',
  `onlineCountApi` varchar(200) DEFAULT NULL COMMENT '在线人数统计API',
  `isPayVoucher` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1为可发代金卷',
  `gameState` tinyint(1) NOT NULL DEFAULT '0',
  `gamePower` varchar(100) NOT NULL DEFAULT '0' COMMENT '权限序列化【各渠道推广游戏权限,0表示所有渠道都没权限',
  `addTime` int(10) NOT NULL,
  `upTime` int(10) NOT NULL DEFAULT '0' COMMENT '母包最后一次更新时间',
  `gPx` int(10) NOT NULL DEFAULT '100' COMMENT '权重，越大排前',
  `SDKVersion` varchar(12) NOT NULL COMMENT 'SDK版本号 例如   1.0.3',
  `IosOrAndroid` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1为IOS；0为安卓',
  `isBanPlatformB` tinyint(1) unsigned NOT NULL COMMENT '是否禁止该游戏使用平台币充值1为禁止',
  `isBanShowGold` tinyint(1) unsigned NOT NULL COMMENT '充值禁止显示元宝     1为禁止',
  `ios_product_id` varchar(64) NOT NULL COMMENT 'ios产品ID',
  `ios_buy_id` varchar(8888) NOT NULL COMMENT 'ios内购字符串',
  `isBuy` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1为买量',
  `isServer` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1表示充值数据的区服要转换成我方区服ID',
  `ios_turn_pay_level` varchar(10) NOT NULL COMMENT 'IOS切支付等级',
  `ios_turn_pay_all_money` mediumint(8) unsigned NOT NULL COMMENT 'IOS切支付累充',
  `ios_turn_pay_money_times` tinyint(3) unsigned NOT NULL COMMENT 'IOS切累充笔数',
  `ios_turn_pay_money_day_limit` smallint(5) unsigned NOT NULL COMMENT 'IOS日限单笔金',
  `ios_turn_pay_times_day_limit` tinyint(3) unsigned NOT NULL COMMENT 'IOS日限充值次数',
  `isH5Pay` tinyint(1) unsigned NOT NULL COMMENT '移动支付版SDK切换到H5支付0为默认1为切换',
  `payVoucherDiscount` float NOT NULL DEFAULT '100' COMMENT '代金券折扣，百分比',
  `isBanSubPackage` tinyint(1) unsigned NOT NULL COMMENT '是否禁止分包1为禁止',
  `isiOSTurnPay` tinyint(1) unsigned NOT NULL COMMENT 'iOS游戏切支付标识，为1则直接切，不需要判断条件',
  `isiOSWebPay` tinyint(1) unsigned NOT NULL COMMENT 'ios网页支付',
  `isBanVoucher` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0:未禁止状态 1：禁止代金卷支付状态',
  `buoyState` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1:开启；0关闭',
  `xieyistate` tinyint(1) NOT NULL DEFAULT '1',
  `otherWayState` tinyint(1) NOT NULL DEFAULT '1',
  `logoState` tinyint(1) NOT NULL DEFAULT '1',
  `abroadPay` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'ios海外支付，1为开启，默认0为关闭',
  PRIMARY KEY (`gameId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='游戏信息表';

-- ----------------------------
-- Records of games
-- ----------------------------
BEGIN;
INSERT INTO `games` VALUES (1000000, '测试游戏', '测试游戏', 'cs', '0', 1000000, 1000000, 'b5d6ad09709d5eac958e8fb8ad511c76', 'ec8eb4cdd219dfed6bf3ef8167c88515', 100, 1, '[1,6]', 10, '元宝', 'http://www.777wan.com', 'http://tp.gzjykj.com/?nmyx=1_1', NULL, NULL, 0, 1, 'a:2:{i:1;s:1:\"1\";i:6;s:1:\"6\";}', 1460011561, 0, 0, '1.0.4', 0, 1, 0, 'com.hg6kwan.SDKDemo', 'com.hongguan.ships.t001|6', 0, 0, '', 0, 0, 0, 0, 0, 45, 0, 0, 0, 1, 1, 1, 1, 1, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
