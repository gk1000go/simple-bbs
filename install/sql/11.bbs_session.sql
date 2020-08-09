# session 表
# 缓存到 runtime 表。 online_0 全局 online_fid 版块。提高遍历效率。
DROP TABLE IF EXISTS bbs_session;
CREATE TABLE bbs_session (
 sid char(32) NOT NULL default '0',			# 随机生成 id 不能重复 uniqueid() 13 位
 uid int unsigned NOT NULL default '0',		# 用户id 未登录为 0，可以重复
 fid int unsigned NOT NULL default '0',		# 所在的版块
 url char(32) NOT NULL default '',			# 当前访问 url
 ip int unsigned NOT NULL default '0',		# 用户ip
 useragent char(128) NOT NULL default '',		# 用户浏览器信息
 data char(255) NOT NULL default '',			# session 数据，超大数据存入大表。
 bigdata int NOT NULL default '0',		# 是否有大数据。
 last_date int unsigned NOT NULL default '0',	# 上次活动时间
 PRIMARY KEY (sid),
 KEY ip (ip),
 KEY fid (fid),
 KEY uid_last_date (uid, last_date)
);