# 版块访问规则, forum.accesson 开启时生效, 记录行数： fid * gid
DROP TABLE IF EXISTS bbs_forum_access;
CREATE TABLE bbs_forum_access (				# 字段中文名
  fid int unsigned NOT NULL default '0',		# fid
  gid int unsigned NOT NULL default '0',		# fid
  allowread int unsigned NOT NULL default '0',	# 允许查看
  allowthread int unsigned NOT NULL default '0',	# 允许发主题
  allowpost int unsigned NOT NULL default '0',	# 允许回复
  allowattach int unsigned NOT NULL default '0',	# 允许上传附件
  allowdown int unsigned NOT NULL default '0',	# 允许下载附件
  PRIMARY KEY (fid, gid)
) ;