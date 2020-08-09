# 持久的 key value 数据存储, ttserver, mysql
DROP TABLE IF EXISTS bbs_kv;
CREATE TABLE bbs_kv (
    k char(32) NOT NULL default '',
    v mediumtext NOT NULL,
    expiry int unsigned NOT NULL default '0',		# 过期时间
    PRIMARY KEY(k)
);