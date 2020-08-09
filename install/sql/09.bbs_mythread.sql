DROP TABLE IF EXISTS bbs_mythread;
CREATE TABLE bbs_mythread (
  uid int unsigned NOT NULL default '0',		# uid
  tid int unsigned NOT NULL default '0',		# 用来清理，删除板块的时候需要
  PRIMARY KEY (uid, tid)				# 每一个帖子只能插入一次 unique
);