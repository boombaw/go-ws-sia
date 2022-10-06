package model

type SyncLulusan struct {
	AhirBim             string      `json:"ahir_bim"`
	Dospem1             interface{} `json:"dospem1"`
	Dospem2             interface{} `json:"dospem2"`
	FlagFeeder          int         `json:"flag_feeder"`
	HasSync             interface{} `json:"has_sync"`
	IDLulusan           int         `json:"id_lulusan"`
	Ipk                 float64     `json:"ipk"`
	JdlSkripsi          string      `json:"jdl_skripsi"`
	JdlSkripsiEn        string      `json:"jdl_skripsi_en"`
	MulaiBim            string      `json:"mulai_bim"`
	NoIjazah            string      `json:"no_ijazah" gorm:"type:varchar(100)"`
	NoTranskripAkademik interface{} `json:"no_transkrip_akademik"`
	NpmMahasiswa        string      `json:"npm_mahasiswa"`
	Name                string      `json:"name"`
	Pem1                interface{} `json:"pem1"`
	Pem2                interface{} `json:"pem2"`
	SkYudisium          string      `json:"sk_yudisium" gorm:"type:text"`
	Sks                 int         `json:"sks"`
	TaLulus             int         `json:"ta_lulus"`
	TglLulus            string      `json:"tgl_lulus"`
	TglYudisium         string      `json:"tgl_yudisium" gorm:"type:datetime"`
}
