package model

type User struct {
	Uid          uint //用户编号
	Gid          uint // 如果要屏蔽，调整用户组即可
	Email        string
	Username     string // 不可以重复
	Realname     string // 真实姓名，天朝预留
	Idnumber     string // 真实身份证号码，天朝预留
	Password     string
	Password_sms string // 预留，手机发送的 sms 验证码
	Salt         string
	Mobile       string // 预留，供二次开发扩展
	Qq           string // 预留，供二次开发扩展，可以弹出QQ直接聊天
	Threads      int    //发帖数
	Posts        int    //回帖数
	Credits      int    // 积分 预留，供二次开发扩展
	Golds        int    // 金币 预留，虚拟币
	Rmbs         int    // 人民币 预留，人民币
	Create_ip    uint   //创建时IP
	Create_date  uint   //创建时间
	Login_ip     uint   //登录时IP
	Login_date   uint   //登录时间
	Logins       uint   //登录次数
	Avatar       uint   //用户最后更新图像时间
	//CreatedAt       time.Time `db:"created_at"`
	//UpdatedAt       time.Time `db:"updated_at"`
	//DeletedAt       *time.Time `db:"deleted_at"`
	Create_ip_fmt string
	Create_date_fmt string
	Login_ip_fmt string
	Login_date_fmt string
	Groupname string
	Avatar_url string
	Avatar_path string
	Online_status int
}



//func InitUser() []User {
//	return []User{
//		User{Uid: 1,Gid: 1,Email: "admin@admin.com",Username: "admin",Password: "admin",Salt: "123456"},
//	}
//}
//INSERT INTO `bbs_user` SET uid=1, gid=1, email='admin@admin.com', username='admin',`password`='admin',salt='123456';


//type User struct {
//	UserId          uint   `gorm:"primary_key;AUTO_INCREMENT'"` //用户编号
//	Gid         uint   `gorm:"NOT NULL;default:0;AUTO_INCREMENT:false"`  // 如果要屏蔽，调整用户组即可
//	Email       	string `gorm:"type:char(40);unique;not null;default:''"`
//	UserName        string `gorm:"type:char(32);unique;not null;default:''"` // 不可以重复
//	RealName        string `gorm:"type:char(16);not null;default:''"`        // 真实姓名，天朝预留
//	IdNumber        string `gorm:"type:char(19);not null;default:''"`        // 真实身份证号码，天朝预留
//	UserPassword    string `gorm:"type:char(32);not null;default:''"`
//	UserPasswordSms string `gorm:"type:char(16);not null;default:''"` // 预留，手机发送的 sms 验证码
//	Salt            string `gorm:"type:char(16);not null;default:''"`
//	Mobile          string `gorm:"type:char(11);not null;default:''"` // 预留，供二次开发扩展
//	Qq              string `gorm:"type:char(15);not null;default:''"` // 预留，供二次开发扩展，可以弹出QQ直接聊天
//	Threads         int    `gorm:"column:beast_id`                //发帖数
//	Posts           int    `gorm:"column:beast_id`                //回帖数
//	Credits         int    `gorm:"column:beast_id`                // 积分 预留，供二次开发扩展
//	Golds           int    `gorm:"column:beast_id`                // 金币 预留，虚拟币
//	Rmbs            int    `gorm:"column:beast_id`                // 人民币 预留，人民币
//	CreatePp        string `gorm:"type:char(45);not null;default:''"` //创建时IP
//	CreateDate      uint   `gorm:"column:beast_id`                //创建时间
//	LoginIp         string `gorm:"type:char(45);not null;default:''"` //登录时IP
//	LoginDate       uint   `gorm:"column:beast_id`                //登录时间
//	Logins          uint   `gorm:"column:beast_id`                //登录次数
//	Avatar          uint   `gorm:"column:beast_id`                //用户最后更新图像时间
//}
