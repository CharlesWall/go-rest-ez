package app

import (
	uuid "github.com/nu7hatch/gouuid"
)

type Service struct {
	db DB
}

type Document map[string]interface{}

func applyResourceToDocument(resource Resource, document *Document) {
	(*document)["id"] = resource.id
	(*document)["path"] = resource.path
	for key, value := range resource.props {
		(*document)[key] = value
	}
}

func (service Service) postDocument(resource Resource, document *Document) (*Document, int) {
	applyResourceToDocument(resource, document)
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, 500
	}
	resource.id = uuid.String()
	(*document)["id"] = resource.id
	resource.path = resource.path + "/" + resource.id
	(*document)["path"] = resource.path
	service.db.setItem(resource.collection, resource.path, document)
	return document, 200
}

func (service Service) getDocument(resource Resource) (*Document, int) {
	document := service.db.getItem(resource.collection, resource.path)
	if document == nil {
		return nil, 404
	}
	return document, 200
}

func (service Service) putDocument(resource Resource, document *Document) (*Document, int) {
	service.db.setItem(resource.collection, resource.path, document)
	return document, 200
}

func (service Service) patchDocument(resource Resource, document *Document) (*Document, int) {
	item := service.db.getItem(resource.collection, resource.path)
	for k, v := range *document {
		(*item)[k] = v
	}
	service.db.setItem(resource.collection, resource.path, item)
	return item, 200
}

func (service Service) deleteDocument(resource Resource) (*Document, int) {
	service.db.deleteItem(resource.collection, resource.path)
	return nil, 200
}

func (service Service) listDocuments(resource Resource, query *map[string]string) ([]*Document, int) {
	return service.db.listItems(resource.collection), 200
}
