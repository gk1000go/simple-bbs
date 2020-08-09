# 我的回帖。大表，需要分区。
DROP TABLE IF EXISTS bbs_mypost;
CREATE TABLE bbs_mypost (
    uid int unsigned NOT NULL default '0',		# uid
    tid int unsigned NOT NULL default '0',		# 用来清理
    pid int unsigned NOT NULL default '0',		#
    KEY (tid),						#
    PRIMARY KEY (uid, pid)				#
);