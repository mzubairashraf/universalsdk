package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"universalsdk/models"
	"universalsdk/service"
	"universalsdk/util"
)

type UsdkController struct {
	usdkService service.UsdkService
}

func NewUsdkController(service service.UsdkService) *UsdkController {
	return &UsdkController{usdkService: service}
}

// Controller handler function to receive request and parse to json.
// After conversion it will pass request to service layer for further processing
func (x UsdkController) DeviceCheck(w http.ResponseWriter, r *http.Request) {

	// Content Type Validation
	if !util.HasContentType(r, "application/json") {
		log.Print(" ## Invalid Content Type ##")
		errorObj := models.ErrorObject{Code: 1, Message: "Content-type should be application/json"}
		util.RespondWithErrorObject(w, errorObj)
		return
	}

	deviceCheckReq, err := parseAndValidateRequest(r)
	if err != nil {
		log.Print(err)
		errorObj := models.ErrorObject{Code: 2, Message: err.Error()}
		util.RespondWithErrorObject(w, errorObj)
		return
	}

	log.Printf(" Request %#v: ", deviceCheckReq)

	// Calling Service to process the request
	serviceResp, err := x.usdkService.DeviceCheck(*deviceCheckReq)

	if err != nil {
		log.Print(err)
		errorObj := models.ErrorObject{Code: 3, Message: err.Error()}
		util.RespondWithErrorObject(w, errorObj)
		return
	}
	util.RespondWithObject(w, serviceResp)

}

func parseAndValidateRequest(r *http.Request) (*models.DeviceCheckDetailsObjectCollection, error) {

	// Parse request to json
	deviceCheckReq := &models.DeviceCheckDetailsObjectCollection{}
	err := json.NewDecoder(r.Body).Decode(deviceCheckReq)
	if err != nil {
		return nil, err
	}

	if deviceCheckReq == nil || len(*deviceCheckReq) <= 0 {
		return nil, fmt.Errorf("invalid or missing input")
	}

	// Validate Request according to Swagger Schema
	err = deviceCheckReq.Validate(nil)
	if err != nil {
		return nil, err
	}

	return deviceCheckReq, nil
}
