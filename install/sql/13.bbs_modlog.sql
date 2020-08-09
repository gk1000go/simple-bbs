# 版主操作日志
DROP TABLE IF EXISTS bbs_modlog;
CREATE TABLE bbs_modlog (
    logid int unsigned NOT NULL auto_increment,	# logid
    uid int unsigned NOT NULL default '0',		# 版主 uid
    tid int unsigned NOT NULL default '0',		# 主题id
    pid int unsigned NOT NULL default '0',		# 帖子id
    subject char(32) NOT NULL default '',			# 主题
    comment char(64) NOT NULL default '',			# 版主评价
    rmbs int NOT NULL default '0',			# 加减人民币, 预留
    create_date int unsigned NOT NULL default '0',	# 时间
    action char(16) NOT NULL default '',			# top|delete|untop
    PRIMARY KEY (logid),
    KEY (uid, logid),
    KEY (tid)
);