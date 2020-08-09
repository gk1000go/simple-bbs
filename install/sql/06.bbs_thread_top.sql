# 置顶主题
DROP TABLE IF EXISTS bbs_thread_top;
CREATE TABLE bbs_thread_top (
    fid int NOT NULL default '0',			# 查找板块置顶
    tid int unsigned NOT NULL default '0',		# tid
    top int unsigned NOT NULL default '0',		# top: 0 是普通最新贴，> 0 置顶贴。
    PRIMARY KEY (tid),					#
    KEY (top, tid),					# 最新贴：top=0 order by tid desc / 全局置顶： top=3
    KEY (fid, top)					# 版块置顶的贴 fid=1 and top=1
);