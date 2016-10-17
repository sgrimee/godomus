package godomus

import "encoding/json"

type CategoryClassId string

type Category struct {
	CatClsId     CategoryClassId `json:"category_clsid"`
	Label        string          `json:"label"`
	PictureKey   PictureKey      `json:"picture_key"`
	DevicesCount int             `json:"devices_count,string"`
}

// CategoriesInRoom returns the list of categories in the given roomId
func (d *Domus) CategoriesInRoom(rk RoomKey) ([]Category, error) {
	queries := map[string]string{
		"room_key": string(rk),
	}
	resp, err := d.GetWithSession("/Mobile/GetCategories", queries)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body struct {
		Categories []Category `json:"category"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	return []Category(body.Categories), nil
}
