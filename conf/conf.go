package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type AppConfig struct {
	BaseDb Mysql `toml:"mysql"`
	BaseCache []Cache `toml:"cache"`
	BaseSite SiteConfig `toml:"site_config"`
	BaseAttach AttachConfig `toml:"attach_config"`

	DEBUG int
	APP_PATH string
	ADMIN_PATH string
	XIUNOPHP_PATH string
}

type Mysql struct {
	Master DbConf
	Slaves []DbConf
}

type DbConf struct {
	Host string
	Port string
	User string
	Password string
	Name string
	TablePre string
	Charset string
	Engine	string
}

type Cache struct {
	Host string
	Port string
	CachePre string
	CacheType string
}

type SiteConfig struct {
	Tmp_path string `toml:"tmp_path"`					// 可以配置为 linux 下的 /dev/shm ，通过内存缓存临时文件
	Log_path string `toml:"log_path"`					// 日志目录
	// -------------------> xiuno bbs 4.0 配置
	Tmpl_path string `toml:"tmpl_path"`
	View_url string `toml:"view_url"`					// 可以配置单独的 CDN 域名：比如：http://static.domain.com/view/
	View_path string  `toml:"view_path"`
	Upload_url string `toml:"upload_url"`				// 可以配置单独的 CDN 域名：比如：http://upload.domain.com/upload/
	Upload_path string `toml:"upload_path"`				// 物理路径，可以用 NFS 存入到单独的文件服务器

	Logo_mobile_url string `toml:"logo_mobile_url"`		// 手机的 LOGO URL
	Logo_pc_url string `toml:"logo_pc_url"`				// PC 的 LOGO URL
	Logo_water_url string `toml:"logo_water_url"`		// 水印的 LOGO URL

	Sitename string `toml:"sitename"`
	Sitebrief string `toml:"sitebrief"`
	Timezone string `toml:"timezone"`					// 时区，默认中国
	Lang string `toml:"lang"`
	Runlevel int `toml:"runlevel"`						// 0: 站点关闭; 1: 管理员可读写; 2: 会员可读;  3: 会员可读写; 4：所有人只读; 5: 所有人可读写
	Runlevel_reason string `toml:"runlevel_reason"`

	Keywords string `toml:"keywords"`
	Description string `toml:"description"`
	Title string `toml:"title"`
	Mobile_link string `toml:"mobile_link"`
	Mobile_title string `toml:"mobile_title"`

	Base_href string `toml:"base_href"`

	Cookie_domain string `toml:"cookie_domain"`
	Cookie_path string `toml:"cookie_path"`
	Auth_key string `toml:"auth_key"`

	Pgesize int `toml:"pagesize"`
	Postlist_pagesize int `toml:"postlist_pagesize"`
	Cache_thread_list_pages int `toml:"cache_thread_list_pages"`
	Online_update_span int `toml:"online_update_span"`				// 在线更新频度，大站设置的长一些
	Online_hold_time int `toml:"online_hold_time"`					// 在线的时间
	Session_delay_update int `toml:"session_delay_update"`
	Upload_image_width int `toml:"upload_image_width"`				// 上传图片自动缩略的最大宽度
	Order_default string `toml:"order_default"`
	Attach_dir_save_rule string `toml:"attach_dir_save_rule"`		// 附件存放规则，附件多用：Ymd，附件少：Ym

	Update_views_on int `toml:"update_views_on"`
	User_create_email_on int `toml:"user_create_email_on"`
	User_create_on int `toml:"user_create_on"`
	User_resetpw_on int `toml:"user_resetpw_on"`

	Admin_bind_ip int `toml:"admin_bind_ip"`						// 后台是否绑定 IP

	Cdn_on int `toml:"cdn_on"`

	/* 支持多种 URL 格式：
	0: ?thread-create-1.htm
	1: thread-create-1.htm
	2: ?/thread/create/1  不支持
	3: /thread/create/1   不支持
	*/
	Url_rewrite_on int `toml:"url_rewrite_on"`

	Disabled_plugin int `toml:"disabled_plugin"`				// 禁止插件

	Version string `toml:"version"`
	Static_version string `toml:"static_version"`
	Installed int `toml:"installed"`
}

type AttachConfig struct {
	All []string
	Video []string
	Music []string
	Exe	[]string
	Flash []string
	Image []string
	Office []string
	Pdf []string
	Text []string
	Zip []string
	Book []string
	Torrent []string
	Font []string
	Other []string
}

var DefaultConfig AppConfig
var SiteConf map[string]string

func InitialConfig(fConf string) error {
	if _, err := toml.DecodeFile(fConf, &DefaultConfig); err != nil {
		fmt.Println(err)
		return err
	}
	DefaultConfig.DEBUG = 1
	DefaultConfig.APP_PATH = "/"
	DefaultConfig.ADMIN_PATH = "admin/"
	DefaultConfig.XIUNOPHP_PATH = "xiunophp/"
	DefaultConfig.BaseSite.Version="4.0.4" // 定义版本号！避免手工修改 conf/conf.php

	SiteConf = map[string]string{
		"sitename":DefaultConfig.BaseSite.Sitename,
		"sitebrief":DefaultConfig.BaseSite.Sitebrief,
		"lang":DefaultConfig.BaseSite.Lang,
		"view_url":DefaultConfig.BaseSite.View_url, // static/
		"logo_pc_url":DefaultConfig.BaseSite.Logo_pc_url,
		"logo_mobile_url":DefaultConfig.BaseSite.Logo_mobile_url,
		"logo_water_url":DefaultConfig.BaseSite.Logo_water_url,
		"version":DefaultConfig.BaseSite.Version,
		"static_version":DefaultConfig.BaseSite.Static_version,
		"timezone":DefaultConfig.BaseSite.Timezone,
		"keywords":DefaultConfig.BaseSite.Keywords,
		"description":DefaultConfig.BaseSite.Description,
		"title":DefaultConfig.BaseSite.Title,
		"mobile_link":DefaultConfig.BaseSite.Mobile_link,
		"mobile_title":DefaultConfig.BaseSite.Mobile_title,
		"base_href":DefaultConfig.BaseSite.Base_href,
		"bootstrap_css":"",
		"bootstrap_bbs_css":"",
	}
	return nil
}