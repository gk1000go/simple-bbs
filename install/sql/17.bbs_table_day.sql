# 系统表, id
# MAXID 表，几个主要的大表，每天的最大ID，用来削减索引 create_date
# day = 0 表示月； month = 0 AND day = 0 表示年
# 计划任务，1点执行。 不需要太精准，用来作为过滤条件。
# 可以有效的过滤冷热数据
DROP TABLE IF EXISTS `bbs_table_day`;
CREATE TABLE `bbs_table_day` (
 `year` smallint(11) unsigned NOT NULL DEFAULT '0' COMMENT '年',	#
 `month` tinyint(11) unsigned NOT NULL DEFAULT '0' COMMENT '月', 	#
 `day` tinyint(11) unsigned NOT NULL DEFAULT '0' COMMENT '日', 		#
 `create_date` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '时间戳', 	#
 `table` char(16) NOT NULL default '' COMMENT '表名',			#
 `maxid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '最大ID', 	#
 `count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '总数', 		#
 PRIMARY KEY (`year`, `month`, `day`, `table`)
);