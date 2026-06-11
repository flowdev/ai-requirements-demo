package demo

import (
	"net/http"

	"github.com/flowdev/ai-requirements-demo/data"
	"github.com/flowdev/ai-requirements-demo/db"
)


func convertCreateMIWURequestToData(req http.Request) (data.MIWU, error) {
	panic("unimplemented")
}

func convertMIWUToDb(miwu data.MIWU) db.MIWU {
	panic("unimplemented")
}

func createMIWUInDb(dbMIWU db.MIWU) (db.MIWU, error) {
	panic("unimplemented")
}

func convertDbMIWUToResponse(dbMIWU db.MIWU) http.Response {
	panic("unimplemented")
}

func convertErrorToResponse(err error) http.Response {
	panic("unimplemented")
}

func checkBasicMIWU(miwu data.MIWU) error {
	panic("unimplemented")
}

func checkMIWUFeaturesExist(miwu data.MIWU) error {
	panic("unimplemented")
}

func checkMIWUSetExists(miwu data.MIWU) error {
	panic("unimplemented")
}