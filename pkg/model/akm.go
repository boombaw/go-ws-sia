package model

// Akm is the model for AKM
type Akm struct {
	ID       string `json:"id" gorm:"id"` // ID is the primary key
	NPM      string `json:"npm" gorm:"npm"`
	Name     string `json:"name" gorm:"name"`
	Ipk      string `json:"ipk" gorm:"ipk"`
	Ips      string `json:"ips" gorm:"ips"`
	Sks      string `json:"sks" gorm:"sks"`
	TotalSks string `json:"total_sks" gorm:"total_sks"`
	Biaya    string `json:"biaya" gorm:"biaya"`
}

type MhsNA struct {
	Name string `json:"name" db:"name"`
	NPM  string `json:"npm" db:"npm"`
}

type MhsDO struct {
	ID          string `json:"id" db:"id"`
	Skep        string `json:"skep" db:"skep"`
	TglSkep     string `json:"tgl_skep" db:"tgl_skep"`
	NPM         string `json:"npm" db:"npm"`
	Name        string `json:"name" db:"name"`
	Alasan      string `json:"alasan" db:"alasan"`
	TahunAjaran string `json:"tahun_ajaran" db:"tahun_ajaran"`
}

type FeederAKM struct {
	IDRegistrasiMahasiswa string `json:"id_registrasi_mahasiswa"`
	IDMahasiswa           string `json:"id_mahasiswa"`
	IDSemester            string `json:"id_semester"`
	NamaSemester          string `json:"nama_semester"`
	Nim                   string `json:"nim"`
	NamaMahasiswa         string `json:"nama_mahasiswa"`
	Angkatan              string `json:"angkatan"`
	IDProdi               string `json:"id_prodi"`
	NamaProgramStudi      string `json:"nama_program_studi"`
	IDStatusMahasiswa     string `json:"id_status_mahasiswa"`
	NamaStatusMahasiswa   string `json:"nama_status_mahasiswa"`
	Ips                   string `json:"ips"`
	Ipk                   string `json:"ipk"`
	SksSemester           string `json:"sks_semester"`
	SksTotal              string `json:"sks_total"`
	BiayaKuliahSmt        string `json:"biaya_kuliah_smt"`
}
