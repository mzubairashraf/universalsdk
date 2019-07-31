package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"universalsdk/models"
)

type usdkServiceImpl struct {
	sessionKeyMap *sync.Map
}

func NewUsdkService(sessionKeyMap *sync.Map) UsdkService {
	return usdkServiceImpl{sessionKeyMap: sessionKeyMap}
}

// This service layer function will perform business validations related to session key and activity data
func (u usdkServiceImpl) DeviceCheck(deviceCheckCollection models.DeviceCheckDetailsObjectCollection) (*models.PuppyObject, error) {
	log.Printf("##  Usdk Service ##")

	activityDataMap := make(map[string]bool)

	// iterating deviceCheckCollection to validate session key, activity data 'kvpKey' uniqueness and data type
	for _, elem := range deviceCheckCollection {
		// Validate Session Key
		err := validateSessionKey(elem, u.sessionKeyMap)
		if err != nil {
			return nil, err
		}

		// Validate Activity Data
		errArray := validateActivityData(elem, activityDataMap)
		if errArray != nil && len(errArray) > 0 {
			str := strings.Join(errArray, ",")
			return nil, fmt.Errorf("activity data validation %s", str)
		}
	}

	return &models.PuppyObject{Puppy: true}, nil
}

// The function validates the session key
// Session key must be unique or an error will be returned.
func validateSessionKey(dCheckDetailsObject *models.DeviceCheckDetailsObject, sessionKeyMap *sync.Map) error {

	if dCheckDetailsObject.CheckSessionKey == "" {
		return nil
	}

	_, ok := sessionKeyMap.Load(dCheckDetailsObject.CheckSessionKey)

	if ok {
		return fmt.Errorf("checkSessionKey should be unique")
	}

	sessionKeyMap.Store(dCheckDetailsObject.CheckSessionKey, "Y")

	return nil
}

// The function validate
// * the list of "Keys" in ActivityData are unique to the call (no double-ups)
// * that the Value provided matches the Type specified.
// Should the verification fail, the error message returned will include information for each KVP pair that fails.
func validateActivityData(dCheckDetailsObject *models.DeviceCheckDetailsObject, activityMap map[string]bool) []string {

	var errstrings []string

	if dCheckDetailsObject.ActivityData == nil || len(dCheckDetailsObject.ActivityData) <= 0 {
		log.Print("activity data is empty")
		return errstrings
	}

	for _, elem := range dCheckDetailsObject.ActivityData {
		// Validate uniqueness of KvpKey
		keyResp := activityMap[elem.KvpKey]
		if keyResp {
			errstrings = append(errstrings, fmt.Sprintf("KvpKey %s is not unique", elem.KvpKey))
		}
		activityMap[elem.KvpKey] = true

		// Validate Data Type of Kvp
		err := validateDataType(elem.KvpValue, elem.KvpType)
		if err != nil {
			errstrings = append(errstrings, fmt.Sprintf("KvpKey %s %s", elem.KvpKey, err.Error()))
		}
	}

	return errstrings
}

// This function validates that provided value is of dataType mentioned in EnumKVPType
func validateDataType(value string, dataType models.EnumKVPType) error {
	var err error

	switch dataType {
	case "general.integer":
		_, err = strconv.ParseInt(value, 10, 64)
	case "general.float":
		_, err = strconv.ParseFloat(value, 64)
	case "general.bool":
		_, err = strconv.ParseBool(value)
	case "general.string":

	default:
		err = fmt.Errorf("data type %s invalid", dataType)
	}

	return err

}
