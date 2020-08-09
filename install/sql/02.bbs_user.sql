
### 用户表 ###
DROP TABLE IF EXISTS `bbs_user`;
CREATE TABLE `bbs_user` (
  uid int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  gid int unsigned NOT NULL DEFAULT '0' COMMENT '用户组编号',	# 如果要屏蔽，调整用户组即可
  email char(40) NOT NULL DEFAULT '' COMMENT '邮箱',
  username char(32) NOT NULL DEFAULT '' COMMENT '用户名',	# 不可以重复
  realname char(16) NOT NULL DEFAULT '' COMMENT '用户名',	# 真实姓名，天朝预留
  idnumber char(19) NOT NULL DEFAULT '' COMMENT '用户名',	# 真实身份证号码，天朝预留
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `password_sms` char(16) NOT NULL DEFAULT '' COMMENT '密码',	# 预留，手机发送的 sms 验证码
  salt char(16) NOT NULL DEFAULT '' COMMENT '密码混杂',
  mobile char(11) NOT NULL DEFAULT '' COMMENT '手机号',		# 预留，供二次开发扩展
  qq char(15) NOT NULL DEFAULT '' COMMENT 'QQ',			# 预留，供二次开发扩展，可以弹出QQ直接聊天
  threads int NOT NULL DEFAULT '0' COMMENT '发帖数',		#
  posts int NOT NULL DEFAULT '0' COMMENT '回帖数',		#
  credits int NOT NULL DEFAULT '0' COMMENT '积分',		# 预留，供二次开发扩展
  golds int NOT NULL DEFAULT '0' COMMENT '金币',		# 预留，虚拟币
  rmbs int NOT NULL DEFAULT '0' COMMENT '人民币',		# 预留，人民币
  create_ip int unsigned NOT NULL DEFAULT '0' COMMENT '创建时IP',
  create_date int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  login_ip int unsigned NOT NULL DEFAULT '0' COMMENT '登录时IP',
  login_date int unsigned NOT NULL DEFAULT '0' COMMENT '登录时间',
  logins int unsigned NOT NULL DEFAULT '0' COMMENT '登录次数',
  avatar int unsigned NOT NULL DEFAULT '0' COMMENT '用户最后更新图像时间',
--  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
--  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
--  deleted_at datetime,
--   KEY deleted_at (deleted_at),
  PRIMARY KEY (uid),
  UNIQUE KEY username (username),
  UNIQUE KEY email (email),						# 升级的时候可能为空
  KEY gid (gid)
);
INSERT INTO `bbs_user` SET uid=1, gid=1, email='admin@admin.com', username='admin',`password`='admin',salt='123456';

# ALTER TABLE bbs_group ADD COLUMN deleted_at datetime;
# ALTER TABLE bbs_group ADD INDEX idx_deleted_at(deleted_at);