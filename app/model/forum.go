package model

type Forum struct {
	Fid uint // fid
	//ForumUp      uint   `gorm:"not null;default:0"`                 // 上一级版块，二级版块作为插件
	Name         string // 版块名称
	Rank         uint   // 显示，倒序，数字越大越靠前
	Threads      uint   // 主题数
	Todayposts   uint   // 今日发帖，计划任务每日凌晨０点清空为０，
	Todaythreads uint   // 今日发主题，计划任务每日凌晨０点清空为０
	Brief        string // 版块简介 允许HTML
	Announcement string // 版块公告 允许HTML
	Accesson     uint   // 是否开启权限控制
	Orderby      int    // 默认列表排序，0: 顶贴时间 last_date， 1: 发帖时间 tid
	Create_date  uint   // 板块创建时间
	Icon         uint   // 板块是否有 icon，存放最后更新时间
	Moduids      string // 每个版块有多个版主，最多10个： 10*12 = 120，删除用户的时候，如果是版主，则调整后再删除。逗号分隔
	Seo_title    string // SEO 标题，如果设置会代替版块名称
	Seo_keywords string // SEO keyword

	Create_date_fmt string
	Icon_url string
	AccessList []*Forum_access
	ModList map[uint]*User
}

type Forum_access struct {
	Fid         uint // fid PK1
	Gid         uint // fid PK2
	Allowread   int  // 允许查看
	Allowthread int  // 允许发主题
	Allowpost   int  // 允许回复
	Allowattach int  // 允许上传附件
	Allowdown   int  // 允许下载附件
}


//func InitForum() []Forum {
//	return []Forum{
//		Forum{Fid: 1,Name: "默认版块",Brief: "默认版块介绍"},
//	}
//}
//INSERT INTO bbs_forum SET fid='1', name='默认版块', brief='默认版块介绍';
//
//type Forum struct {
//	ForumId      uint   `gorm:"primary_key;AUTO_INCREMENT'"`        // fid
//	ForumUp      uint   `gorm:"not null;default:0"`                 // 上一级版块，二级版块作为插件
//	ForumName    string `gorm:"type:char(16);not null;default:''"`  // 版块名称
//	ForumRank    uint   `gorm:"not null;default:0"`                 // 显示，倒序，数字越大越靠前
//	Threads      uint   `gorm:"not null;default:0"`                 // 主题数
//	TodayPosts   uint   `gorm:"not null;default:0"`                 // 今日发帖，计划任务每日凌晨０点清空为０，
//	TodayThreads uint   `gorm:"not null;default:0"`                 // 今日发主题，计划任务每日凌晨０点清空为０
//	Brief        string `gorm:"type:text;not null"`      // 版块简介 允许HTML
//	Announcement string `gorm:"type:text;not null"`      // 版块公告 允许HTML
//	Accession    uint   `gorm:"not null;default:0"`                 // 是否开启权限控制
//	OrderBy      int    `gorm:"not null;default:0"`                 // 默认列表排序，0: 顶贴时间 last_date， 1: 发帖时间 tid
//	CreateDate   uint   `gorm:"not null;default:0"`                 // 板块创建时间
//	Icon         uint   `gorm:"not null;default:0"`                 // 板块是否有 icon，存放最后更新时间
//	Modules      string `gorm:"type:char(120);not null;default:''"` // 每个版块有多个版主，最多10个： 10*12 = 120，删除用户的时候，如果是版主，则调整后再删除。逗号分隔
//	SeoTitle     string `gorm:"type:char(64);not null;default:''"`  // SEO 标题，如果设置会代替版块名称
//	SeoKeywords  string `gorm:"type:char(64);not null;default:''"`  // SEO keyword
//}