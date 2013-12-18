package data

type Item struct{}

type DataManager struct {
	items  []*Item
	lastID int64
}

func NewDataManager() *DataManager {
	return &DataManager{}
}

func (m *DataManager) Save(item *Item) error { return nil }

func (m *DataManager) All() []*Item {
	return m.items
}

func cloneItem(i *Item) *Item {
	c := *i
	return &c

}
