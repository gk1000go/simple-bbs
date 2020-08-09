# 论坛主题
DROP TABLE IF EXISTS bbs_thread;
CREATE TABLE bbs_thread (
    fid int NOT NULL default '0',			# 版块 id
    tid int unsigned NOT NULL auto_increment,		# 主题id
    top int NOT NULL default '0',			# 置顶级别: 0: 普通主题, 1-3 置顶的顺序
    uid int unsigned NOT NULL default '0',		# 用户id
    userip int unsigned NOT NULL default '0',		# 发帖时用户ip ip2long()，主要用来清理
    subject char(128) NOT NULL default '',		# 主题
    create_date int unsigned NOT NULL default '0',	# 发帖时间
    last_date int unsigned NOT NULL default '0',	# 最后回复时间

    views int unsigned NOT NULL default '0',		# 查看次数, 剥离出去，单独的服务，避免 cache 失效
    posts int unsigned NOT NULL default '0',		# 回帖数
    images int NOT NULL default '0',		# 附件中包含的图片数
    files int NOT NULL default '0',		# 附件中包含的文件数
    mods int NOT NULL default '0',			# 预留：版主操作次数，如果 > 0, 则查询 modlog，显示斑竹的评分
    closed int unsigned NOT NULL default '0',	# 预留：是否关闭，关闭以后不能再回帖、编辑。
    firstpid int unsigned NOT NULL default '0',	# 首贴 pid
    lastuid int unsigned NOT NULL default '0',	# 最近参与的 uid
    lastpid int unsigned NOT NULL default '0',	# 最后回复的 pid
#     created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
#     updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
#     deleted_at datetime,
#     KEY deleted_at (deleted_at),
    PRIMARY KEY (tid),					# 主键
    KEY (lastpid),					# 最后回复排序
    KEY (fid, tid),					# 发帖时间排序，正序。数据量大时可以考虑建立小表，对小表进行分区优化，只有数据量达到千万级以上时才需要。
    KEY (fid, lastpid)					# 顶贴时间排序，倒序
) ;
