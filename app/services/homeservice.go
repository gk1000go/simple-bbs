package services

import (
	"awesomeProject1/app-web/app/dao"
	"awesomeProject1/app-web/app/services/svcfunc"
	"awesomeProject1/app-web/framework"
)

type HomeService struct {
	framework.BaseService
	dao *dao.HomeDAO
}

func NewHomeService() *HomeService {
	return &HomeService{dao:dao.NewHomeDao()}
}
func (this *HomeService) GetGroupList() (framework.ResultMaps) {
	return svcfunc.Group_list_cache()
}
func (this *HomeService) GetUser(uid uint) (framework.BaseDataMap) {
	return svcfunc.User_read(uid)
}

//func (this *HomeService) GetUserIdFromToken() uint {
//	return user_read(uid)
//}






/*
	$user = user__read($uid);
	user_format($user);
 */

/*
	forum.Create_date_fmt = services.IntToTime("2006-01-02",int64(forum.Create_date))
	site := framework.GetApp().Cfg.BaseSite
	if forum.Icon > 0{
		forum.Icon_url = fmt.Sprintf("%sforum/%d.png",site.Upload_url,forum.Fid)
	}else{
		forum.Icon_url = site.View_url + "img/forum.png"
	}
	forum.AccessList = nil
	if forum.Accesson > 0{
		forum.AccessList = this.service.Forum_access_find_by_fid(forum.Fid)
	}
	forum.ModList = nil
 */

