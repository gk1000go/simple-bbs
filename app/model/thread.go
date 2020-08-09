package model

type Thread struct {
	Fid         int    // 版块 id
	Tid         uint   // 主题id
	Top         int    // 置顶级别: 0: 普通主题, 1-3 置顶的顺序
	Uid         uint   // 用户id
	Userip      uint   // 发帖时用户ip ip2long()，主要用来清理
	Subject     string // 主题
	Create_date uint   // 发帖时间
	Last_date   uint   // 最后回复时间
	Views    uint // 查看次数, 剥离出去，单独的服务，避免 cache 失效
	Posts    uint // 回帖数
	Images   int  // 附件中包含的图片数
	Files    int  // 附件中包含的文件数
	Mods     int  // 预留：版主操作次数，如果 > 0, 则查询 modlog，显示斑竹的评分
	Closed   uint // 预留：是否关闭，关闭以后不能再回帖、编辑。
	Firstpid uint // 首贴 pid
	Lastuid  uint // 最近参与的 uid
	Lastpid  uint // 最后回复的 pid
}

type Thread_top struct {
	Fid int  // 查找板块置顶
	Tid uint // tid
	Top uint // top: 0 是普通最新贴，> 0 置顶贴。
}

type Mythread struct {
	Uid uint
	Tid uint
}