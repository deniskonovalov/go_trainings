package handlers

import (
	"errors"
	"lesson_12/internal/helpers"
	"net"

	rsp "lesson_12/cmd/server/responses"
	ds "lesson_12/internal/documentstore"
)

type Handler struct {
	Conn  net.Conn
	store *ds.Store
	coll  *ds.Collection
}

func NewHandler(store *ds.Store, conn net.Conn, coll *ds.Collection) *Handler {
	return &Handler{
		Conn:  conn,
		store: store,
		coll:  coll,
	}
}

func (h *Handler) HandleCreateCollection(inputJson string) rsp.ServerResponse {
	m, err := helpers.JsonToMap(inputJson)
	if err != nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, err.Error())
	}

	pkey, ok := m["primary_key"]
	if !ok {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, ds.ErrMissingPrimaryKey.Error())
	}

	name, ok := m["name"].(string)
	if !ok {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, ds.ErrMissingNameString.Error())
	}

	cnf := ds.CollectionConfig{
		PrimaryKey: pkey.(string),
	}

	_, err = h.store.CreateCollection(name, &cnf)
	if err != nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, err.Error())
	}

	return rsp.NewServerResponse(rsp.StatusOk, nil, "")
}

func (h *Handler) HandleSelectCollection(inputJson string) rsp.ServerResponse {
	m, err := helpers.JsonToMap(inputJson)
	if err != nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, err.Error())
	}

	name, ok := m["name"].(string)
	if !ok {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, ds.ErrMissingNameString.Error())
	}

	col, err := h.store.GetCollection(name)
	if err != nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, err.Error())
	}

	h.coll = col

	return rsp.NewServerResponse(rsp.StatusOk, nil, "")
}

func (h *Handler) HandleDeleteCollection(inputJson string) rsp.ServerResponse {
	m, err := helpers.JsonToMap(inputJson)
	if err != nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, err.Error())
	}

	name, ok := m["name"].(string)
	if !ok {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, ds.ErrMissingNameString.Error())
	}

	err = h.store.DeleteCollection(name)
	if err != nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, err.Error())
	}

	return rsp.NewServerResponse(rsp.StatusOk, nil, "")
}

func (h *Handler) HandleListCollections() rsp.ServerResponse {
	return rsp.NewServerResponse(rsp.StatusOk, h.store.ListCollections(), "")
}

func (h *Handler) HandlePutDocument(inputJson string) rsp.ServerResponse {
	d, err := ds.MakeDocument([]byte(inputJson))
	if err != nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, err.Error())
	}

	if h.coll == nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, ds.ErrCollectionIsNotSelected.Error())
	}

	err = h.coll.Put(*d)
	if err != nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, err.Error())
	}

	return rsp.NewServerResponse(rsp.StatusOk, d, "")
}

func (h *Handler) HandleGetDocument(inputJson string) rsp.ServerResponse {
	m, err := helpers.JsonToMap(inputJson)
	if err != nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, err.Error())
	}

	if h.coll == nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, ds.ErrCollectionIsNotSelected.Error())
	}

	key, ok := m[h.coll.GetPrimaryKey()].(string)
	if !ok {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, errors.New("missing '"+h.coll.GetPrimaryKey()+"' parameter").Error())
	}

	d, err := h.coll.Get(key)
	if err != nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, err.Error())
	}

	return rsp.NewServerResponse(rsp.StatusOk, d, "")
}

func (h *Handler) HandleDeleteDocument(inputJson string) rsp.ServerResponse {
	if h.coll == nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, ds.ErrCollectionIsNotSelected.Error())
	}

	m, err := helpers.JsonToMap(inputJson)
	if err != nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, err.Error())
	}
	key, ok := m[h.coll.GetPrimaryKey()].(string)
	if !ok {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, errors.New("missing '"+h.coll.GetPrimaryKey()+"' parameter").Error())
	}
	err = h.coll.Delete(key)
	if err != nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, err.Error())
	}

	return rsp.NewServerResponse(rsp.StatusOk, nil, "")
}

func (h *Handler) HandleListDocuments() rsp.ServerResponse {
	if h.coll == nil {
		return rsp.NewServerResponse(rsp.StatusFailed, nil, ds.ErrCollectionIsNotSelected.Error())
	}

	return rsp.NewServerResponse(rsp.StatusOk, h.coll.List(), "")
}
