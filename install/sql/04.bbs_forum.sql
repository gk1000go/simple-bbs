# 板块表，一级, runtime 中存放 forumlist 格式化以后的数据。
DROP TABLE IF EXISTS bbs_forum;
CREATE TABLE bbs_forum (
   fid int unsigned NOT NULL auto_increment,		# fid
# fup int(11) unsigned NOT NULL auto_increment,	# 上一级版块，二级版块作为插件
   name char(16) NOT NULL default '',			# 版块名称
   `rank` tinyint(3) unsigned NOT NULL default '0',	# 显示，倒序，数字越大越靠前
   threads mediumint(8) unsigned NOT NULL default '0',	# 主题数
   todayposts mediumint(8) unsigned NOT NULL default '0',# 今日发帖，计划任务每日凌晨０点清空为０，
   todaythreads mediumint(8) unsigned NOT NULL default '0',# 今日发主题，计划任务每日凌晨０点清空为０
   brief text NOT NULL,					# 版块简介 允许HTML
   announcement text NOT NULL,				# 版块公告 允许HTML
   accesson int unsigned NOT NULL default '0',	# 是否开启权限控制
   orderby int NOT NULL default '0',		# 默认列表排序，0: 顶贴时间 last_date， 1: 发帖时间 tid
   create_date int unsigned NOT NULL default '0',	# 板块创建时间
   icon int unsigned NOT NULL default '0',		# 板块是否有 icon，存放最后更新时间
   moduids char(120) NOT NULL default '',		# 每个版块有多个版主，最多10个： 10*12 = 120，删除用户的时候，如果是版主，则调整后再删除。逗号分隔
   seo_title char(64) NOT NULL default '',		# SEO 标题，如果设置会代替版块名称
   seo_keywords char(64) NOT NULL default '',		# SEO keyword
   PRIMARY KEY (fid)
) ;
INSERT INTO bbs_forum SET fid='1', name='默认版块', brief='默认版块介绍',announcement='no announcement';
#  cache_date int(11) NOT NULL default '0',		# 最后 threadlist 缓存的时间，6种排序前10页结果缓存。如果是前10页，先读缓存，并依据此字段过期。更新条件：发贴
