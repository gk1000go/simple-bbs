# 缓存表，用来保存临时数据。
DROP TABLE IF EXISTS bbs_cache;
CREATE TABLE bbs_cache (
   k char(32) NOT NULL default '',
   v mediumtext NOT NULL,
   expiry int unsigned NOT NULL default '0',		# 过期时间
   PRIMARY KEY(k)
);