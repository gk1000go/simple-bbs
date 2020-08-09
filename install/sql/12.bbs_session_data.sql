DROP TABLE IF EXISTS bbs_session_data;
CREATE TABLE bbs_session_data (
  sid char(32) NOT NULL default '0',			#
  last_date int unsigned NOT NULL default '0',	# 上次活动时间
  data text NOT NULL,					# 存超大数据
  PRIMARY KEY (sid)
);