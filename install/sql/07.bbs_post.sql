# 论坛帖子数据
DROP TABLE IF EXISTS bbs_post;
CREATE TABLE bbs_post (
  tid int unsigned NOT NULL default '0',		# 主题id
  pid int unsigned NOT NULL auto_increment,		# 帖子id
  uid int unsigned NOT NULL default '0',		# 用户id
  isfirst int unsigned NOT NULL default '0',	# 是否为首帖，与 thread.firstpid 呼应
  create_date int unsigned NOT NULL default '0',	# 发贴时间
  userip int unsigned NOT NULL default '0',		# 发帖时用户ip ip2long()
  images int NOT NULL default '0',		# 附件中包含的图片数
  files int NOT NULL default '0',		# 附件中包含的文件数
  doctype int NOT NULL default '0',		# 类型，0: html, 1: txt; 2: markdown; 3: ubb
  quotepid int NOT NULL default '0',		# 引用哪个 pid，可能不存在
  message longtext NOT NULL,				# 内容，用户提示的原始数据
  message_fmt longtext NOT NULL,			# 内容，存放的过滤后的html内容，可以定期清理，减肥。
#   created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
#   updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
#   deleted_at datetime,
#   KEY deleted_at (deleted_at),
  PRIMARY KEY (pid),
  KEY (tid, pid),
  KEY (uid)						# 我的回帖，清理数据需要
) ;
# 编辑历史