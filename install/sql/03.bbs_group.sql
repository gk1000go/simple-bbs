# 用户组
DROP TABLE IF EXISTS `bbs_group`;
CREATE TABLE `bbs_group` (
 gid int unsigned NOT NULL,			#
 name char(20) NOT NULL default '',			# 用户组名称
 creditsfrom int NOT NULL default '0',		# 积分从
 creditsto int NOT NULL default '0',		# 积分到
 allowread int NOT NULL default '0',		# 允许访问
 allowthread int NOT NULL default '0',		# 允许发主题
 allowpost int NOT NULL default '0',		# 允许回帖
 allowattach int NOT NULL default '0',		# 允许上传文件
 allowdown int NOT NULL default '0',		# 允许下载文件
 allowtop int NOT NULL default '0',		# 允许置顶
 allowupdate int NOT NULL default '0',		# 允许编辑
 allowdelete int NOT NULL default '0',		# 允许删除
 allowmove int NOT NULL default '0',		# 允许移动
 allowbanuser int NOT NULL default '0',		# 允许禁止用户
 allowdeleteuser int NOT NULL default '0',		# 允许删除用户
 allowviewip int unsigned NOT NULL default '0',	# 允许查看用户敏感信息
 PRIMARY KEY (gid)
) ;
INSERT INTO `bbs_group` SET gid='0', name="游客组", creditsfrom='0', creditsto='0', allowread='1', allowthread='0', allowpost='1', allowattach='0', allowdown='1', allowtop='0', allowupdate='0', allowdelete='0', allowmove='0', allowbanuser='0', allowdeleteuser='0', allowviewip='0';
INSERT INTO `bbs_group` SET gid='1', name="管理员组", creditsfrom='0', creditsto='0', allowread='1', allowthread='1', allowpost='1', allowattach='1', allowdown='1', allowtop='1', allowupdate='1', allowdelete='1', allowmove='1', allowbanuser='1', allowdeleteuser='1', allowviewip='1';
INSERT INTO `bbs_group` SET gid='2', name="超级版主组", creditsfrom='0', creditsto='0', allowread='1', allowthread='1', allowpost='1', allowattach='1', allowdown='1', allowtop='1', allowupdate='1', allowdelete='1', allowmove='1', allowbanuser='1', allowdeleteuser='1', allowviewip='1';
INSERT INTO `bbs_group` SET gid='4', name="版主组", creditsfrom='0', creditsto='0', allowread='1', allowthread='1', allowpost='1', allowattach='1', allowdown='1', allowtop='1', allowupdate='1', allowdelete='1', allowmove='1', allowbanuser='1', allowdeleteuser='0', allowviewip='1';
INSERT INTO `bbs_group` SET gid='5', name="实习版主组", creditsfrom='0', creditsto='0', allowread='1', allowthread='1', allowpost='1', allowattach='1', allowdown='1', allowtop='1', allowupdate='1', allowdelete='0', allowmove='1', allowbanuser='0', allowdeleteuser='0', allowviewip='0';
INSERT INTO `bbs_group` SET gid='6', name="待验证用户组", creditsfrom='0', creditsto='0', allowread='1', allowthread='0', allowpost='1', allowattach='0', allowdown='1', allowtop='0', allowupdate='0', allowdelete='0', allowmove='0', allowbanuser='0', allowdeleteuser='0', allowviewip='0';
INSERT INTO `bbs_group` SET gid='7', name="禁止用户组", creditsfrom='0', creditsto='0', allowread='0', allowthread='0', allowpost='0', allowattach='0', allowdown='0', allowtop='0', allowupdate='0', allowdelete='0', allowmove='0', allowbanuser='0', allowdeleteuser='0', allowviewip='0';
INSERT INTO `bbs_group` SET gid='101', name="一级用户组", creditsfrom='0', creditsto='50', allowread='1', allowthread='1', allowpost='1', allowattach='1', allowdown='1', allowtop='0', allowupdate='0', allowdelete='0', allowmove='0', allowbanuser='0', allowdeleteuser='0', allowviewip='0';
INSERT INTO `bbs_group` SET gid='102', name="二级用户组", creditsfrom='50', creditsto='200', allowread='1', allowthread='1', allowpost='1', allowattach='1', allowdown='1', allowtop='0', allowupdate='0', allowdelete='0', allowmove='0', allowbanuser='0', allowdeleteuser='0', allowviewip='0';
INSERT INTO `bbs_group` SET gid='103', name="三级用户组", creditsfrom='200', creditsto='1000', allowread='1', allowthread='1', allowpost='1', allowattach='1', allowdown='1', allowtop='0', allowupdate='0', allowdelete='0', allowmove='0', allowbanuser='0', allowdeleteuser='0', allowviewip='0';
INSERT INTO `bbs_group` SET gid='104', name="四级用户组", creditsfrom='1000', creditsto='10000', allowread='1', allowthread='1', allowpost='1', allowattach='1', allowdown='1', allowtop='0', allowupdate='0', allowdelete='0', allowmove='0', allowbanuser='0', allowdeleteuser='0', allowviewip='0';
INSERT INTO `bbs_group` SET gid='105', name="五级用户组", creditsfrom='10000', creditsto='10000000', allowread='1', allowthread='1', allowpost='1', allowattach='1', allowdown='1', allowtop='0', allowupdate='0', allowdelete='0', allowmove='0', allowbanuser='0', allowdeleteuser='0', allowviewip='0';
