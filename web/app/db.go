package app

type Collection map[string]*Document

type DB struct {
	collections map[string]*Collection
}

func NewDB() DB {
	return DB{
		collections: make(map[string]*Collection),
	}
}

func (db *DB) setItem(collectionName string, key string, value *Document) bool {
	println(collectionName, key)
	if collectionName == "" {
		return false
	}
	collection := db.collections[collectionName]
	if collection == nil {
		c := make(Collection)
		println("creating new collection", c, &c)
		db.collections[collectionName] = &c
		collection = &c
	}
	(*collection)[key] = value
	return true
}

func (db DB) getItem(collectionName string, key string) *Document {
	collection := db.collections[collectionName]
	if collection == nil {
		return nil
	}
	return (*collection)[key]
}

func (db *DB) deleteItem(collectionName string, key string) {
	collection := db.collections[collectionName]
	if collection == nil {
		return
	}
	delete((*collection), key)
}

func (db *DB) listItems(collectionName string) []*Document {
	list := []*Document{}
	collection := db.collections[collectionName]

	if collection != nil {
		for _, v := range *collection {
			list = append(list, v)
		}
	}

	return list
}
