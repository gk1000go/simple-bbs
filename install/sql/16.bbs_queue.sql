# 临时队列，用来保存临时数据。
DROP TABLE IF EXISTS bbs_queue;
CREATE TABLE bbs_queue (
   queueid int(11) unsigned NOT NULL default '0',		# 队列 id
   v int(11) NOT NULL default '0',			# 队列中存放的数据，只能为 int
   expiry int(11) unsigned NOT NULL default '0',		# 过期时间，默认 0，不过期
   UNIQUE KEY(queueid, v),
   KEY(expiry)
);