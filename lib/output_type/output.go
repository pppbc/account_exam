package output_type

import (
	"errors"
	"account_exam/proto"
)

func JsonArrTo(js interface{}) (info interface{}, err error) {
	//---------------------------------------------
	//根据类型，判断执行怎样操作
	switch js.(type) {
	case []proto.Staffs:
		info, err = staffs(js.([]proto.Staffs))
	case []proto.Department:

	default:
		err = errors.New("No Such Handler")
	}
	return
}

func staffs(staff []proto.Staffs) (info []proto.StuffOutput, err error) {
	for k, _ := range staff {
		//存output类型
		var in proto.StuffOutput

		in.Avatar = staff[k].Avatar
		in.ID = staff[k].ID
		in.PlantID = staff[k].PlantID
		in.UID = staff[k].UID
		in.Name = staff[k].Name
		in.Sex = staff[k].Sex
		in.CreatedAt = staff[k].CreatedAt
		in.UpdatedAt = staff[k].UpdatedAt
		in.Deleted = staff[k].Deleted
		in.IPCEnabled = staff[k].IPCEnabled
		in.LoginEnabled = staff[k].LoginEnabled
		in.JobNumber = staff[k].JobNumber

		info = append(info, in)

		if staff[k].Department != nil {
			//err = json.Unmarshal(staff[k].Department.([]byte), &info[k].Department)
		}
		if staff[k].Post != nil {
			//err = json.Unmarshal(staff[k].Post.([]byte), &info[k].Post)
		}
		if staff[k].User != nil {
			//err = json.Unmarshal(staff[k].User.([]byte), &info[k].User)
		}
		return
	}
	return
}
