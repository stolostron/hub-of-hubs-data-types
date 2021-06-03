package hub_of_hubs_data_types

type ObjectsBundle struct {
	Objects 		[]*interface{} `json:"objects"`
	DeletedObjects 	[]*interface{} `json:"deletedObjects"`
}

func (bundle *ObjectsBundle) AddObject(object *interface{}) {
	bundle.Objects = append(bundle.Objects, object)
}

func (bundle *ObjectsBundle) AddDeletedObject(object *interface{}) {
	bundle.DeletedObjects = append(bundle.DeletedObjects, object)
}
