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
  `f_fid` char(40) NOT NULL COMMENT '用户谱册关系表主键',
  `f_uid` char(40) DEFAULT NULL COMMENT '用户表主键',
  `f_vid` char(40) DEFAULT NULL COMMENT '谱册表主键',
  `f_create_time` datetime DEFAULT NULL COMMENT '关系创建时间（即收藏该谱册的时间）',
  PRIMARY KEY (`f_fid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `tab_volume` */

DROP TABLE IF EXISTS `tab_volume`;

CREATE TABLE `tab_volume` (
  `f_vid` char(40) NOT NULL COMMENT '主键',
  `f_name` text COMMENT '谱册名',
  `f_cover` text COMMENT '谱册封面url',
  `f_desc` text COMMENT '谱册描述',
  `f_uid` char(40) DEFAULT NULL COMMENT '用户表主键',
  `f_status` int(32) DEFAULT NULL COMMENT '谱册状态，公开、非公开等',
  `f_create_time` datetime DEFAULT NULL COMMENT '谱册创建时间',
  PRIMARY KEY (`f_vid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `tab_volume_sheet` */

DROP TABLE IF EXISTS `tab_volume_sheet`;

CREATE TABLE `tab_volume_sheet` (
  `f_vsid` char(40) NOT NULL COMMENT '谱册和谱子关系表主键',
  `f_vid` char(40) DEFAULT NULL COMMENT '谱册表主键',
  `f_sid` char(40) DEFAULT NULL COMMENT '谱子的id（可能来自于github）',
  `f_create_time` datetime DEFAULT NULL COMMENT '关系的创建时间（即谱子加入谱册的时间）',
  PRIMARY KEY (`f_vsid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
