package service

import (
	"github.com/stretchr/testify/suite"
	"log"
	"strconv"
	"sync"
	"testing"
	"universalsdk/models"
	"universalsdk/util"
)

type UsdkServiceSuite struct {
	suite.Suite
	AccountIds []string
	AcctId     string
}

func TestUsdkServiceSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(UsdkServiceSuite))
}

func (suite *UsdkServiceSuite) TestValidateSessionKey() {
	var sessionKeyMap sync.Map
	deviceCheckModel := &models.DeviceCheckDetailsObject{}

	// Test Unique Session Key
	deviceCheckModel.CheckSessionKey = "123654"
	err := validateSessionKey(deviceCheckModel, &sessionKeyMap)
	if err != nil {
		suite.T().Errorf("validate session key failure %s", err.Error())
	}

	// Test Unique Session Key
	deviceCheckModel.CheckSessionKey = "369852"
	err = validateSessionKey(deviceCheckModel, &sessionKeyMap)
	if err != nil {
		suite.T().Errorf("validate session key failure %s", err.Error())
	}

	// Test Duplicate Session Key
	deviceCheckModel.CheckSessionKey = "123654"
	err = validateSessionKey(deviceCheckModel, &sessionKeyMap)
	if err == nil {
		suite.T().Errorf("validate session key expecting failure got none %s", err.Error())
	}
}

func (suite *UsdkServiceSuite) TestValidateActivityData() {
	activityDataMap := make(map[string]bool)

	keyValuePairObject1 := &models.KeyValuePairObject{KvpKey: "ip.address", KvpValue: "1.23.45.123", KvpType: models.EnumKVPType("general.string")}
	keyValuePairObject2 := &models.KeyValuePairObject{KvpKey: "mac.address", KvpValue: "1.23.45.123", KvpType: models.EnumKVPType("general.string")}
	keyValArray := make([]*models.KeyValuePairObject, 0)
	keyValArray = append(keyValArray, keyValuePairObject1, keyValuePairObject2)
	deviceCheckModel := &models.DeviceCheckDetailsObject{ActivityData: keyValArray}

	// Test Unique Key
	err := validateActivityData(deviceCheckModel, activityDataMap)
	if err != nil && len(err) > 0 {
		suite.T().Errorf("validate activity data key uniqueness failure")
	}

	// Test duplicate Key
	keyValuePairObject2.KvpType = "ip.address"
	err = validateActivityData(deviceCheckModel, activityDataMap)
	if err == nil || len(err) <= 0 {
		suite.T().Errorf("validate activity data duplicate key failure")
	}

	// Invalid Data type
	keyValuePairObject2.KvpType = "web"
	keyValuePairObject2.KvpType = models.EnumKVPTypeGeneralInteger
	keyValuePairObject2.KvpValue = "www"
	err = validateActivityData(deviceCheckModel, activityDataMap)
	if err == nil && len(err) <= 0 {
		suite.T().Errorf("validate activity data - invalid data type ")
	}

}

func (suite *UsdkServiceSuite) TestValidateDataType() {

	// Test for valid values
	err := validateDataType("false", models.EnumKVPTypeGeneralBool)
	if err != nil {
		suite.T().Errorf("validate data type failure %s", err.Error())
	}
	err = validateDataType("true", models.EnumKVPTypeGeneralBool)
	if err != nil {
		suite.T().Errorf("validate data type failure %s", err.Error())
	}
	err = validateDataType("12.3326", models.EnumKVPTypeGeneralFloat)
	if err != nil {
		suite.T().Errorf("validate data type failure %s", err.Error())
	}
	err = validateDataType("122365", models.EnumKVPTypeGeneralInteger)
	if err != nil {
		suite.T().Errorf("validate data type failure %s", err.Error())
	}
	err = validateDataType("test", models.EnumKVPTypeGeneralString)
	if err != nil {
		suite.T().Errorf("validate data type failure %s", err.Error())
	}

	// Test with Invalid values
	err = validateDataType("123", models.EnumKVPTypeGeneralBool)
	if err == nil {
		suite.T().Errorf("expecting error, got none %s", err.Error())
	}
	err = validateDataType("test", models.EnumKVPTypeGeneralFloat)
	if err == nil {
		suite.T().Errorf("expecting error, got none %s", err.Error())
	}
	err = validateDataType("test", models.EnumKVPTypeGeneralInteger)
	if err == nil {
		suite.T().Errorf("expecting error, got none %s", err.Error())
	}
}

func (suite *UsdkServiceSuite) TestDeviceCheck() {
	mockRequest := mockRequest()
	usdkService := createService()
	resp, err := usdkService.DeviceCheck(mockRequest)

	if err != nil {
		suite.T().Errorf("Device Check failure %s", err.Error())
	}

	if resp.Puppy != true {
		suite.T().Errorf("Expecting Puppy to be set as true got %t", resp.Puppy)
	}

}

func (suite *UsdkServiceSuite) TestDeviceCheckWithSameSessionRequest() {
	mockRequest := mockSameSessionKeyRequest()
	usdkService := createService()
	_, err := usdkService.DeviceCheck(mockRequest)

	if err == nil {
		suite.T().Errorf("Device Check With Same Session Key Request. Expecting Failure got none")
	}
}

func (suite *UsdkServiceSuite) TestDeviceCheckWithInvalidActivityDataKeyType() {
	mockRequest := mockActivityKeyWithInvalidDataTypeRequest()
	usdkService := createService()
	_, err := usdkService.DeviceCheck(mockRequest)

	if err == nil {
		suite.T().Errorf("Device Check With Invalid ActicityData KeyType. Expecting failure got none")
	}
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

func mockSameSessionKeyRequest() models.DeviceCheckDetailsObjectCollection {
	rand := strconv.Itoa(util.GenerateRandomInRange(1000000, 20000000))
	keyValuePairObject := &models.KeyValuePairObject{KvpKey: "ip.address", KvpValue: "1.23.45.123", KvpType: models.EnumKVPType("general.string")}
	keyValArray := make([]*models.KeyValuePairObject, 0)
	keyValArray = append(keyValArray, keyValuePairObject)

	deviceCheckDetail1 := &models.DeviceCheckDetailsObject{CheckType: "DEVICE", ActivityType: "SIGNUP", CheckSessionKey: rand, ActivityData: keyValArray}
	deviceCheckDetail2 := &models.DeviceCheckDetailsObject{CheckType: "DEVICE", ActivityType: "SIGNUP", CheckSessionKey: rand}
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

func createService() UsdkService {
	var sessionKeyMap sync.Map
	usdkService := NewUsdkService(&sessionKeyMap)
	return usdkService
}

func (suite *UsdkServiceSuite) SetupSuite() {
	log.Println("## In Suite Setup - creating ACCT  ##")

}

func (suite *UsdkServiceSuite) TearDownSuite() {
	log.Println("## In Suite Tear Down ## ")
}
