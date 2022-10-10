package repository

type smsProdiRepository struct{}

type SmsProdiRepository interface {
	SmsProdi(arg SmsParams) string
}

type SmsParams struct {
	KdProdi string `json:"kd_prodi"`
}

func NewSmsProdiRepository() *smsProdiRepository {
	return &smsProdiRepository{}
}

func (l *smsProdiRepository) SMSProdi(arg SmsParams) string {

	var sms = getSMSProdi(arg.KdProdi)
	// var db = database.Conn()
	// var data []uint8

	// row := db.Raw(query.SmsProdi, arg.KdProdi).Row()
	// row.Scan(&data)

	// sql, _ := db.DB()
	// defer sql.Close()

	// return string(data)
	return sms
}

func getSMSProdi(kd_prodi string) string {

	var list = make(map[string]string)

	list["62201"] = "a26699d1-68f5-43cd-9eea-183d02932b8a"
	list["73201"] = "683afb04-010d-4a79-bb9e-cae295051ebb"
	list["70201"] = "1be96463-2d80-4490-a4fe-960c1fd2abc2"
	list["25201"] = "eea608d7-8266-4867-9d00-9f2d010b9b75"
	list["24201"] = "01877499-5a53-4ad4-a333-3859d9364371"
	list["55201"] = "6a6afbb7-adc6-48f0-b775-3e3b13815f5a"
	list["26201"] = "dab0ab91-6fde-487c-8249-649a68d9acad"
	list["61201"] = "b87a16a1-81a3-47f4-915c-65210cb18eee"
	list["74201"] = "461199b6-844c-4df3-a301-6f51376f3010"
	list["74101"] = "bb78eb4e-c042-4efd-a11b-6f6563cdb3e3"
	list["61101"] = "345e2677-f4b1-4407-9eae-281e2b1c675d"
	list["32201"] = "afa1bca2-b129-4fa5-99ab-0157a0772f1c"
	list["86206"] = "12a1527b-7f13-4120-ba8d-774b38947ae6"
	list["85202"] = "95b9cd9f-5e77-4b78-a322-85927a25cf66"

	if val, ok := list[kd_prodi]; ok {
		return val
	}

	return ""
}
