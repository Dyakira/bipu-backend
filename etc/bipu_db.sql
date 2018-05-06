/*
SQLyog Ultimate v11.11 (64 bit)
MySQL - 5.6.24-log : Database - bipubipu
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
/*Table structure for table `tab_favorite` */

DROP TABLE IF EXISTS `tab_favorite`;

CREATE TABLE `tab_favorite` (
  `f_fid` int(32) NOT NULL COMMENT '用户谱册关系表主键',
  `f_uid` int(32) DEFAULT NULL COMMENT '用户表主键',
  `f_vid` int(32) DEFAULT NULL COMMENT '谱册表主键',
  `f_create_time` datetime DEFAULT NULL COMMENT '关系创建时间（即收藏该谱册的时间）',
  PRIMARY KEY (`f_fid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `tab_volume` */

DROP TABLE IF EXISTS `tab_volume`;

CREATE TABLE `tab_volume` (
  `f_vid` int(32) NOT NULL COMMENT '主键',
  `f_name` text COMMENT '谱册名',
  `f_cover` text COMMENT '谱册封面url',
  `f_desc` text COMMENT '谱册描述',
  `f_uid` int(32) DEFAULT NULL COMMENT '用户表主键',
  `f_status` int(32) DEFAULT NULL COMMENT '谱册状态，公开、非公开等',
  `f_create_time` datetime DEFAULT NULL COMMENT '谱册创建时间',
  PRIMARY KEY (`f_vid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `tab_volume_sheet` */

DROP TABLE IF EXISTS `tab_volume_sheet`;

CREATE TABLE `tab_volume_sheet` (
  `f_vsid` int(32) NOT NULL COMMENT '谱册和谱子关系表主键',
  `f_vid` int(32) DEFAULT NULL COMMENT '谱册表主键',
  `f_sid` int(32) DEFAULT NULL COMMENT '谱子的id（可能来自于github）',
  `f_create_time` datetime DEFAULT NULL COMMENT '关系的创建时间（即谱子加入谱册的时间）',
  PRIMARY KEY (`f_vsid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `uid` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `login` varchar(20) DEFAULT NULL COMMENT '登陆账号',
  `password` varchar(50) DEFAULT NULL COMMENT 'MD5密码',
  `name` varchar(10) DEFAULT NULL COMMENT '昵称',
  `avatar_url` varchar(255) DEFAULT NULL COMMENT '头像',
  `tel` char(15) DEFAULT NULL COMMENT '手机',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `sex` tinyint(2) DEFAULT '0' COMMENT '性别',
  `status` tinyint(2) DEFAULT '0' COMMENT '状态',
  `openid` varchar(50) NOT NULL COMMENT '授权id',
  `source` varchar(10) NOT NULL COMMENT '第三方平台(github)',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日期',
  `open_token` varchar(255) DEFAULT NULL COMMENT '授权token',
  `bio` varchar(255) DEFAULT NULL COMMENT '个人简介',
  PRIMARY KEY (`uid`)
) ENGINE=MyISAM AUTO_INCREMENT=123145 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
