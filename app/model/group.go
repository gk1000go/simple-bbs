package model

type Group struct {
	Gid             uint
	Name            string // 用户组名称
	Creditsfrom     int    // 积分从
	Creditsto       int    // 积分到
	Allowread       int    // 允许访问
	Allowthread     int    // 允许发主题
	Allowpost       int    // 允许回帖
	Allowattach     int    // 允许上传文件
	Allowdown       int    // 允许下载文件
	Allowtop        int    // 允许置顶
	Allowupdate     int    // 允许编辑
	Allowdelete     int    // 允许删除
	Allowmove       int    // 允许移动
	Allowbanuser    int    // 允许禁止用户
	Allowdeleteuser int    // 允许删除用户
	Allowviewip     uint   // 允许查看用户敏感信息
}
