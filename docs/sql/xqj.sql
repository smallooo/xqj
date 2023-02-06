/*
Navicat MySQL Data Transfer
Source Database       : pss
Target Server Type    : MYSQL
Target Server Version : 50639
File Encoding         : 65001
Date: 2018-03-18 16:52:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------

CREATE TABLE IF NOT EXISTS `xqj_profile` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
    `title` varchar(18) DEFAULT '' COMMENT '标题',
    `desc` varchar(255) DEFAULT '' COMMENT '简述',
    `content` text COMMENT '内容',
    `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `deleted_on` int(10) unsigned DEFAULT '0',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '删除时间',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='相亲对象信息';

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------

CREATE TABLE IF NOT EXISTS `xqj_auth` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(50) DEFAULT '' COMMENT '账号',
    `password` varchar(50) DEFAULT '' COMMENT '密码',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for user_login
-- ----------------------------
CREATE TABLE IF NOT EXISTS `user_login` (
                                            `uid` int unsigned NOT NULL,
                                            `email` varchar(128) NOT NULL DEFAULT '',
    `username` varchar(20) NOT NULL COMMENT '用户名',
    `passcode` char(12) NOT NULL DEFAULT '' COMMENT '加密随机数',
    `passwd` char(32) NOT NULL DEFAULT '' COMMENT 'md5密码',
    `login_ip` varchar(31) NOT NULL DEFAULT '' COMMENT '最后登录 IP',
    `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次登录时间（主动登录或cookie登录）',
    PRIMARY KEY (`uid`),
    UNIQUE KEY (`username`),
    UNIQUE KEY (`email`),
    KEY `logintime` (`login_time`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户登录表';



-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------

CREATE TABLE IF NOT EXISTS `xqj_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '标签名称',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';





-- ----------------------------
-- Table structure for user_info
-- ----------------------------
CREATE TABLE IF NOT EXISTS `user_info` (
                                           `uid` int unsigned NOT NULL AUTO_INCREMENT,
                                           `email` varchar(128) NOT NULL DEFAULT '',
    `open` tinyint NOT NULL DEFAULT 0 COMMENT '邮箱是否公开，默认不公开',
    `username` varchar(20) NOT NULL COMMENT '用户名',
    `name` varchar(20) NOT NULL DEFAULT '' COMMENT '姓名',
    `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像(如果为空，则使用http://www.gravatar.com)',
    `city` varchar(10) NOT NULL DEFAULT '' COMMENT '居住地',
    `company` varchar(63) NOT NULL DEFAULT '' COMMENT '公司',
    `github` varchar(31) NOT NULL DEFAULT '' COMMENT 'Github昵称',
    `gitea` varchar(31) NOT NULL DEFAULT '' COMMENT 'Gitea昵称',
    `weibo` varchar(31) NOT NULL DEFAULT '' COMMENT '微博昵称',
    `website` varchar(63) NOT NULL DEFAULT '' COMMENT '个人主页，博客',
    `monlog` varchar(140) NOT NULL DEFAULT '' COMMENT '个人状态，签名，独白',
    `introduce` varchar(2022) NOT NULL COMMENT '个人简介',
    `unsubscribe` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '是否退订本站邮件，0-否；1-是',
    `is_third` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '是否通过第三方账号注册',
    `balance` int unsigned NOT NULL DEFAULT 0 COMMENT '财富余额（铜币）',
    `dau_auth` int unsigned NOT NULL DEFAULT 0 COMMENT '控制用户权限，如能否发文章等',
    `is_vip` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '是否是VIP付费用户',
    `vip_expire` int unsigned NOT NULL DEFAULT 0 COMMENT 'VIP到期日期，格式20200301',
    `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '用户账号状态。0-默认；1-已审核；2-拒绝；3-冻结；4-停号',
    `is_root` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '是否超级用户，不受权限控制：1-是',
    `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`uid`),
    UNIQUE KEY (`username`),
    UNIQUE KEY (`email`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户信息表';




