package godomus

type Categories []Category
type CategoryClassId string

type Category struct {
	CatClsId     CategoryClassId `json:"category_clsid"`
	Label        string          `json:"label"`
	PictureKey   string          `json:"picture_key"`
	DevicesCount int             `json:"devices_count,string"`
}
