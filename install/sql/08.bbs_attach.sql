#论坛附件表  只能按照从上往下的方式查找和删除！ 此表如果大，可以考虑通过 aid 分区。
DROP TABLE IF EXISTS bbs_attach;
CREATE TABLE bbs_attach (
    aid int unsigned NOT NULL auto_increment ,	# 附件id
    tid int NOT NULL default '0',			# 主题id
    pid int NOT NULL default '0',			# 帖子id
    uid int NOT NULL default '0',			# 用户id
    filesize int unsigned NOT NULL default '0',	# 文件尺寸，单位字节
    width int unsigned NOT NULL default '0',	# width > 0 则为图片
    height int unsigned NOT NULL default '0',	# height
    filename char(120) NOT NULL default '',		# 文件名称，会过滤，并且截断，保存后的文件名，不包含URL前缀 upload_url
    orgfilename char(120) NOT NULL default '',		# 上传的原文件名
    filetype char(7) NOT NULL default '',			# 文件类型: image/txt/zip，小图标显示 <i class="icon filetype image"></i>
    create_date int unsigned NOT NULL default '0',	# 文件上传时间 UNIX 时间戳
    comment char(100) NOT NULL default '',		# 文件注释 方便于搜索
    downloads int NOT NULL default '0',		# 下载次数，预留
    credits int NOT NULL default '0',			# 需要的积分，预留
    golds int NOT NULL default '0',			# 需要的金币，预留
    rmbs int NOT NULL default '0',			# 需要的人民币，预留
    isimage int NOT NULL default '0',		# 是否为图片
#     created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
#     updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
#     deleted_at datetime,
#     KEY deleted_at (deleted_at),
    PRIMARY KEY (aid),					# aid
    KEY pid (pid),					# 每个帖子下多个附件
    KEY uid (uid)						# 我的附件，清理数据需要。
);