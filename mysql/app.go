package mysql

//app的模型


type App struct {
	ID         int64  `json:"id" xorm:"pk autoincr 'id'"`
	PlatformID int64  `json:"platform_id" xorm:"notnull"`
	Name       string `json:"name" xorm:"notnull unique"`
	Script     string `json:"script" xorm:"notnull"`
	Links      []Link `json:"links" xorm:"-"`
}

type AppPayload struct {
	//ID         int64  `json:"id" xorm:"pk autoincr 'id'"`
	//PlatformID int64  `json:"platform_id" xorm:"notnull" comment:"平台id"`
	Name       string `json:"name" xorm:"notnull unique" comment:"名称"`
	Script     string `json:"script" xorm:"notnull" comment:"脚本"`
}


func GetAllApps() ([]App, error) {
	var err error
	apps := make([]App, 0)
	err = engine.Find(&apps)
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func GetAppsByPlatformID(platformID int64) ([]App, error) {
	apps := make([]App, 0)
	err := engine.Where("platform_id = ?", platformID).Find(&apps)
	return apps, err
}

//func GetAllLinks() ([]App, error) {
//	apps := make([]App, 0)
//	err := engine.Find(&apps)
//	if err != nil {
//		return nil, err
//	}
//	for idx, app := range apps {
//		links, err := GetLinksByAppID(app.ID)
//		if err != nil {
//			return nil, err
//		}
//		apps[idx].Links = links
//	}
//	return apps, nil
//}

func UpdateAppById(appId int64, app *App) error {
	_, err := engine.ID(appId).Update(app)
	return err
}

func GetPlatformIdById(appId int64) (int64, error) {
	app := App{}
	_, err := engine.ID(appId).Get(&app)
	return app.PlatformID, err
}
