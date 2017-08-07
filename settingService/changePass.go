package settingService

import (
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
	"github.com/8tomat8/SSU-Golang-252-Chat/database"

)

// ChangePassRequest is a structure for request to changing of password
type ChangePassRequest struct {
	Header RequestHeader `json:"header"`
	Body   json.RawMessage `json:"body"`
}

// UnmarshalChangePassRequest is a function for unmarshaling request for changing of password from []byte JSON
// into map[string]interface{}
func UnmarshalChangePassRequest(byteRequest [] byte) (map[string]interface{}, error) {
	var unmarshaledRequest map[string]interface{}
	err := json.Unmarshal(byteRequest, &unmarshaledRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return nil, err
	}
	return unmarshaledRequest, err
}

// ChangePass perform changing password of users
func ChangePass(changePassRequest [] byte) (bool, error) {
	unmarshaledRequest, err := UnmarshalChangePassRequest(changePassRequest)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	userName := unmarshaledRequest["user_name"]
	OldPass := unmarshaledRequest["old_pass"]
	NewPass := unmarshaledRequest["new_pass"]
	if userName == nil || OldPass == nil || NewPass == nil {
		err := errors.New("Empty field or fields")
		loger.Log.Errorf("Some field or fields are empty: ")
		return false, err
	}
	if OldPass == NewPass {
		loger.Log.Warn("Old and new passwords are the same ")
	}
	db, err := GetStorage() // common gorm-connection from database package
	defer db.Close()
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}
	db.Update("password", NewPass).Where("user_name = ?", userName)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
