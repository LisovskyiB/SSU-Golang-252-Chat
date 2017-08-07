package settingService

import (
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
	"github.com/8tomat8/SSU-Golang-252-Chat/database"

)

// ChangeNickNameRequest is a structure for request to change nick-name of users
type ChangeNickNameRequest struct {
	Header RequestHeader `json:"header"`
	Body   json.RawMessage `json:"body"`
}

// UnmarshalChangeNickNameRequest is a function for unmarshaling request for changing nick-name from []byte JSON
// to map[string]interface{}
func UnmarshalChangeNickNameRequest(changeNickNameRequest [] byte) (map[string]interface{}, error) {
	var unmarshaledRequest map[string]interface{}
	err := json.Unmarshal(changeNickNameRequest, unmarshaledRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return nil, err
	}
	return unmarshaledRequest, nil
}

// ChangeNickName perform changing nick-name of users
func ChangeNickName(changeNickNameRequest [] byte) (bool, error) {
	unmarshaledRequest, err := UnmarshalChangeNickNameRequest(changeNickNameRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return false, err
	}
	nickName := unmarshaledRequest["nick_name"]
	userName := unmarshaledRequest["user_name"]
	if nickName == nil || userName == nil{
		err := errors.New("Empty field or fields")
		loger.Log.Errorf("Some field or fields are empty: ")
		return false, err
	}
	db, err := GetStorage() // common gorm-connection from database package
	defer db.Close()
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}
	db.Update("nick_name", nickName).Where("user_name = ?", userName)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
