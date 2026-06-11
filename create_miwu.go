package demo

import (
	"net/http"

	"github.com/flowdev/ai-requirements-demo/data"
)


func CreateMIWU(req http.Request) http.Response {
    miwu, err := convertCreateMIWURequestToData(req)
    if err != nil {
		return convertErrorToResponse(err)
	}
    if err = validateMIWU(miwu); err != nil {
		return convertErrorToResponse(err)
	}
    dbMIWU := convertMIWUToDb(miwu)

    dbMIWU, err = createMIWUInDb(dbMIWU)
    if err != nil {
		return convertErrorToResponse(err)
	}
    return convertDbMIWUToResponse(dbMIWU)
}

func validateMIWU(miwu data.MIWU) error {
    if err := checkBasicMIWU(miwu); err != nil {
		return err
	}
    if err := checkMIWUFeaturesExist(miwu); err != nil {
		return err
	}
    if err := checkMIWUSetExists(miwu); err != nil {
		return err
	}
	return nil
}
