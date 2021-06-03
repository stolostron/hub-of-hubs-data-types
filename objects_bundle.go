package datatypes

type ObjectsBundle struct {
	Objects 		[]interface{} `json:"objects"`
	DeletedObjects 	[]interface{} `json:"deletedObjects"`
}

func NewObjectBundle() *ObjectsBundle {
	bundle := ObjectsBundle {
		Objects: make([]interface{}, 0),
		DeletedObjects: make([]interface{}, 0),
	}
	return &bundle
}

func (bundle *ObjectsBundle) AddObject(object interface{}) {
	bundle.Objects = append(bundle.Objects, object)
}

func (bundle *ObjectsBundle) AddDeletedObject(object interface{}) {
	bundle.DeletedObjects = append(bundle.DeletedObjects, object)
}
