package mysql


type Link struct {
	ID        int64  `json:"id" xorm:"pk autoincr 'id'"`
	Name      string  `json:"name" xorm:"notnull unique"`
	Url       string  `json:"url" xorm:"notnull unique"`
}

func DaoGetLinks() (links []Link ,err error) {
	links = make([]Link, 0)
	err = engine.Find(&links)
	return links,err
}


func DaoAddLink(link *Link) (err error) {
	_, err = engine.Insert(link)
	return err
}

func DaoDeleteLink(id int64) error {
	link := new(Link)
	_, err := engine.Id(id).Delete(link)
	return err
}

func DaoGetLink(id int64) (link *Link, err error) {
	link = &Link{}
	exist,err := engine.ID(id).Get(link)
	if !exist {
		link = nil
	}
	return link, err
}

func DaoUpdateLinkById(id int64 , link *Link) (err error) {
	_, err = engine.ID(id).Update(link)
	return err
}
