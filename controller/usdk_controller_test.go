package controller

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
	"universalsdk/models"
	"universalsdk/service"
	"universalsdk/util"
)

type UsdkControllerSuite struct {
	suite.Suite
	AccountIds []string
	AcctId     string
}

func TestUsdkControllerSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(UsdkControllerSuite))
}

func (suite *UsdkControllerSuite) TestEmptyRequest() {

	mockRequest := &models.DeviceCheckDetailsObjectCollection{}
	jsonAccount, _ := json.Marshal(mockRequest)
	usdkController := createUsdkController()

	req, _ := http.NewRequest("POST", "/isgood", bytes.NewBuffer(jsonAccount))

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(usdkController.DeviceCheck)
	handler.ServeHTTP(response, req)

	checkResponseCode(suite.T(), http.StatusBadRequest, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["message"] == "" {
		suite.T().Errorf("Expected the detail of error. Got '%s'", m["status"])
	}

	suite.T().Log(m["message"])
}

func (suite *UsdkControllerSuite) TestInvalidContentType() {

	mockRequest := &models.DeviceCheckDetailsObjectCollection{}
	jsonAccount, _ := json.Marshal(mockRequest)
	usdkController := createUsdkController()

	req, _ := http.NewRequest("POST", "/isgood", bytes.NewBuffer(jsonAccount))

	req.Header.Set("Content-Type", "application/xml")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(usdkController.DeviceCheck)
	handler.ServeHTTP(response, req)

	checkResponseCode(suite.T(), http.StatusBadRequest, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["message"] == "" {
		suite.T().Errorf("Expected the detail of error. Got '%s'", m["status"])
	}

	suite.T().Log(m["message"])
}

func (suite *UsdkControllerSuite) TestInvalidRequest() {

	mockRequest := mockInvalidRequest()
	jsonAccount, _ := json.Marshal(mockRequest)
	usdkController := createUsdkController()

	req, _ := http.NewRequest("POST", "/isgood", bytes.NewBuffer(jsonAccount))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(usdkController.DeviceCheck)
	handler.ServeHTTP(response, req)

	checkResponseCode(suite.T(), http.StatusBadRequest, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["message"] == "" {
		suite.T().Errorf("Expected the detail of error. Got '%s'", m["status"])
	}

	suite.T().Log(m["message"])
}

func (suite *UsdkControllerSuite) TestSameSessionKey() {

	mockRequest := mockSessionKeyRequest()
	jsonAccount, _ := json.Marshal(mockRequest)
	usdkController := createUsdkController()

	req, _ := http.NewRequest("POST", "/isgood", bytes.NewBuffer(jsonAccount))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(usdkController.DeviceCheck)
	handler.ServeHTTP(response, req)

	checkResponseCode(suite.T(), http.StatusBadRequest, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["message"] == "" {
		suite.T().Errorf("Expected the detail of error. Got '%s'", m["status"])
	}
	suite.T().Log(m["message"])
}

func (suite *UsdkControllerSuite) TestDuplicateActivityKey() {

	mockRequest := mockDuplicateActivityKeyRequest()
	jsonAccount, _ := json.Marshal(mockRequest)
	usdkController := createUsdkController()

	req, _ := http.NewRequest("POST", "/isgood", bytes.NewBuffer(jsonAccount))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(usdkController.DeviceCheck)
	handler.ServeHTTP(response, req)

	checkResponseCode(suite.T(), http.StatusBadRequest, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["message"] == "" {
		suite.T().Errorf("Expected the detail of error. Got '%s'", m["status"])
	}
	suite.T().Log(m["message"])
}

func (suite *UsdkControllerSuite) TestActivityKeyWithInvalidDataType() {

	mockRequest := mockActivityKeyWithInvalidDataTypeRequest()
	jsonAccount, _ := json.Marshal(mockRequest)
	usdkController := createUsdkController()

	req, _ := http.NewRequest("POST", "/isgood", bytes.NewBuffer(jsonAccount))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(usdkController.DeviceCheck)
	handler.ServeHTTP(response, req)

	checkResponseCode(suite.T(), http.StatusBadRequest, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["message"] == "" {
		suite.T().Errorf("Expected the detail of error. Got '%s'", m["status"])
	}
	suite.T().Log(m["message"])
}

func (suite *UsdkControllerSuite) TestVendorActivityType() {

	mockRequest := mockVendorActivityTypeRequest()
	jsonAccount, _ := json.Marshal(mockRequest)
	usdkController := createUsdkController()

	req, _ := http.NewRequest("POST", "/isgood", bytes.NewBuffer(jsonAccount))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(usdkController.DeviceCheck)
	handler.ServeHTTP(response, req)

	checkResponseCode(suite.T(), http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["message"] == "" {
		suite.T().Errorf("Expected the detail of error. Got '%s'", m["status"])
	}
	suite.T().Log(m["message"])
}

func (suite *UsdkControllerSuite) TestInvalidActivityType() {

	mockRequest := mockInvalidActivityTypeRequest()
	jsonAccount, _ := json.Marshal(mockRequest)
	usdkController := createUsdkController()

	req, _ := http.NewRequest("POST", "/isgood", bytes.NewBuffer(jsonAccount))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(usdkController.DeviceCheck)
	handler.ServeHTTP(response, req)

	checkResponseCode(suite.T(), http.StatusBadRequest, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["message"] == "" {
		suite.T().Errorf("Expected the detail of error. Got '%s'", m["status"])
	}
	suite.T().Log(m["message"])
}

func (suite *UsdkControllerSuite) TestInvalidCheckType() {

	mockRequest := mockInvalidCheckTypeRequest()
	jsonAccount, _ := json.Marshal(mockRequest)
	usdkController := createUsdkController()

	req, _ := http.NewRequest("POST", "/isgood", bytes.NewBuffer(jsonAccount))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(usdkController.DeviceCheck)
	handler.ServeHTTP(response, req)

	checkResponseCode(suite.T(), http.StatusBadRequest, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["message"] == "" {
		suite.T().Errorf("Expected the detail of error. Got '%s'", m["status"])
	}
	suite.T().Log(m["message"])
}

func (suite *UsdkControllerSuite) TestSuccessRequest() {

	mockRequest := mockRequest()
	jsonAccount, _ := json.Marshal(mockRequest)
	usdkController := createUsdkController()

	req, _ := http.NewRequest("POST", "/isgood", bytes.NewBuffer(jsonAccount))
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(usdkController.DeviceCheck)
	handler.ServeHTTP(response, req)

	checkResponseCode(suite.T(), http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["puppy"] != true {
		suite.T().Errorf("Expected the 'puppy' in response to be set to 'true'. Got '%s'", m["puppy"])
	}
}

func mockInvalidRequest() models.DeviceCheckDetailsObjectCollection {
	deviceCheckDetail1 := &models.DeviceCheckDetailsObject{CheckType: "DEVICE", ActivityType: "SIGNUP", CheckSessionKey: "123654789"}
	deviceCheckDetail2 := &models.DeviceCheckDetailsObject{CheckType: "DUMMY", ActivityType: "DUMMY"}
	mockDeviceCheckCollection := &models.DeviceCheckDetailsObjectCollection{deviceCheckDetail1, deviceCheckDetail2}
	return *mockDeviceCheckCollection
}

func mockSessionKeyRequest() models.DeviceCheckDetailsObjectCollection {
	rand := strconv.Itoa(util.GenerateRandomInRange(1000000, 20000000))
	keyValuePairObject := &models.KeyValuePairObject{KvpKey: "ip.address", KvpValue: "1.23.45.123", KvpType: models.EnumKVPType("general.string")}
	keyValArray := make([]*models.KeyValuePairObject, 0)
	keyValArray = append(keyValArray, keyValuePairObject)

	deviceCheckDetail1 := &models.DeviceCheckDetailsObject{CheckType: "DEVICE", ActivityType: "SIGNUP", CheckSessionKey: rand, ActivityData: keyValArray}
	deviceCheckDetail2 := &models.DeviceCheckDetailsObject{CheckType: "DEVICE", ActivityType: "SIGNUP", CheckSessionKey: rand}
	mockDeviceCheckCollection := &models.DeviceCheckDetailsObjectCollection{deviceCheckDetail1, deviceCheckDetail2}
	return *mockDeviceCheckCollection
}

func mockDuplicateActivityKeyRequest() models.DeviceCheckDetailsObjectCollection {
	rand := strconv.Itoa(util.GenerateRandomInRange(1000000, 20000000))
	rand2 := strconv.Itoa(util.GenerateRandomInRange(3000000, 40000000))
	keyValuePairObject := &models.KeyValuePairObject{KvpKey: "ip.address", KvpValue: "1.23.45.123", KvpType: models.EnumKVPType("general.string")}
	keyValArray := make([]*models.KeyValuePairObject, 0)
	keyValArray = append(keyValArray, keyValuePairObject)

	deviceCheckDetail1 := &models.DeviceCheckDetailsObject{CheckType: "DEVICE", ActivityType: "SIGNUP", CheckSessionKey: rand, ActivityData: keyValArray}
	deviceCheckDetail2 := &models.DeviceCheckDetailsObject{CheckType: "DEVICE", ActivityType: "SIGNUP", CheckSessionKey: rand2, ActivityData: keyValArray}
	mockDeviceCheckCollection := &models.DeviceCheckDetailsObjectCollection{deviceCheckDetail1, deviceCheckDetail2}
	return *mockDeviceCheckCollection
}

func mockActivityKeyWithInvalidDataTypeRequest() models.DeviceCheckDetailsObjectCollection {
	rand := strconv.Itoa(util.GenerateRandomInRange(1000000, 20000000))
	rand2 := strconv.Itoa(util.GenerateRandomInRange(3000000, 40000000))

	keyValuePairObject := &models.KeyValuePairObject{KvpKey: "ip.address", KvpValue: "1.23.45.123", KvpType: models.EnumKVPType("general.bool")}
	keyValuePairObject2 := &models.KeyValuePairObject{KvpKey: "ip.address", KvpValue: "asdfh", KvpType: models.EnumKVPType("general.float")}
	keyValArray := make([]*models.KeyValuePairObject, 0)
	keyValArray = append(keyValArray, keyValuePairObject)
	keyValArray = append(keyValArray, keyValuePairObject2)

	deviceCheckDetail1 := &models.DeviceCheckDetailsObject{CheckType: "DEVICE", ActivityType: "SIGNUP", CheckSessionKey: rand, ActivityData: keyValArray}
	deviceCheckDetail2 := &models.DeviceCheckDetailsObject{CheckType: "DEVICE", ActivityType: "SIGNUP", CheckSessionKey: rand2, ActivityData: keyValArray}
	mockDeviceCheckCollection := &models.DeviceCheckDetailsObjectCollection{deviceCheckDetail1, deviceCheckDetail2}
	return *mockDeviceCheckCollection
}

func mockVendorActivityTypeRequest() models.DeviceCheckDetailsObjectCollection {
	rand := strconv.Itoa(util.GenerateRandomInRange(1000000, 20000000))
	keyValuePairObject := &models.KeyValuePairObject{KvpKey: "ip.address", KvpValue: "1.23.45.123", KvpType: models.EnumKVPType("general.string")}
	keyValArray := make([]*models.KeyValuePairObject, 0)
	keyValArray = append(keyValArray, keyValuePairObject)

	deviceCheckDetail1 := &models.DeviceCheckDetailsObject{CheckType: "DEVICE", ActivityType: "_SIGNUP", CheckSessionKey: rand, ActivityData: keyValArray}
	mockDeviceCheckCollection := &models.DeviceCheckDetailsObjectCollection{deviceCheckDetail1}
	return *mockDeviceCheckCollection
}

func mockInvalidActivityTypeRequest() models.DeviceCheckDetailsObjectCollection {
	rand := strconv.Itoa(util.GenerateRandomInRange(1000000, 20000000))
	keyValuePairObject := &models.KeyValuePairObject{KvpKey: "ip.address", KvpValue: "1.23.45.123", KvpType: models.EnumKVPType("general.string")}
	keyValArray := make([]*models.KeyValuePairObject, 0)
	keyValArray = append(keyValArray, keyValuePairObject)

	deviceCheckDetail1 := &models.DeviceCheckDetailsObject{CheckType: "DEVICE", ActivityType: "DUMMY", CheckSessionKey: rand, ActivityData: keyValArray}
	mockDeviceCheckCollection := &models.DeviceCheckDetailsObjectCollection{deviceCheckDetail1}
	return *mockDeviceCheckCollection
}

func mockInvalidCheckTypeRequest() models.DeviceCheckDetailsObjectCollection {
	rand := strconv.Itoa(util.GenerateRandomInRange(1000000, 20000000))
	keyValuePairObject := &models.KeyValuePairObject{KvpKey: "ip.address", KvpValue: "1.23.45.123", KvpType: models.EnumKVPType("general.string")}
	keyValArray := make([]*models.KeyValuePairObject, 0)
	keyValArray = append(keyValArray, keyValuePairObject)

	deviceCheckDetail1 := &models.DeviceCheckDetailsObject{CheckType: "DUMMY", ActivityType: "SIGNUP", CheckSessionKey: rand, ActivityData: keyValArray}
	mockDeviceCheckCollection := &models.DeviceCheckDetailsObjectCollection{deviceCheckDetail1}
	return *mockDeviceCheckCollection
}

func mockRequest() models.DeviceCheckDetailsObjectCollection {
	rand := strconv.Itoa(util.GenerateRandomInRange(1000000, 20000000))
	keyValuePairObject := &models.KeyValuePairObject{KvpKey: "ip.address", KvpValue: "1.23.45.123", KvpType: models.EnumKVPType("general.string")}
	keyValArray := make([]*models.KeyValuePairObject, 0)
	keyValArray = append(keyValArray, keyValuePairObject)

	deviceCheckDetail1 := &models.DeviceCheckDetailsObject{CheckType: "DEVICE", ActivityType: "SIGNUP", CheckSessionKey: rand, ActivityData: keyValArray}
	mockDeviceCheckCollection := &models.DeviceCheckDetailsObjectCollection{deviceCheckDetail1}
	return *mockDeviceCheckCollection
}

func createUsdkController() *UsdkController {
	var sessionKeyMap sync.Map
	usdkService := service.NewUsdkService(&sessionKeyMap)
	usdkController := NewUsdkController(usdkService)
	return usdkController
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func (suite *UsdkControllerSuite) SetupSuite() {
	log.Println("## In Suite Setup - creating ACCT  ##")

}

func (suite *UsdkControllerSuite) TearDownSuite() {
	log.Println("## In Suite Tear Down ## ")
}
