package model

// FeederRiwayatPendidikan
type FeederRiwayatPendidikan struct {
	IDRegistrasiMahasiswa   string      `json:"id_registrasi_mahasiswa"`
	IDMahasiswa             string      `json:"id_mahasiswa"`
	Nim                     string      `json:"nim"`
	NamaMahasiswa           string      `json:"nama_mahasiswa"`
	IDJenisDaftar           string      `json:"id_jenis_daftar"`
	NamaJenisDaftar         string      `json:"nama_jenis_daftar"`
	IDJalurDaftar           string      `json:"id_jalur_daftar"`
	IDPeriodeMasuk          string      `json:"id_periode_masuk"`
	NamaPeriodeMasuk        string      `json:"nama_periode_masuk"`
	IDJenisKeluar           interface{} `json:"id_jenis_keluar"`
	KeteranganKeluar        interface{} `json:"keterangan_keluar"`
	IDPerguruanTinggi       string      `json:"id_perguruan_tinggi"`
	NamaPerguruanTinggi     string      `json:"nama_perguruan_tinggi"`
	IDProdi                 string      `json:"id_prodi"`
	NamaProgramStudi        string      `json:"nama_program_studi"`
	SksDiakui               string      `json:"sks_diakui"`
	IDPerguruanTinggiAsal   interface{} `json:"id_perguruan_tinggi_asal"`
	NamaPerguruanTinggiAsal interface{} `json:"nama_perguruan_tinggi_asal"`
	IDProdiAsal             interface{} `json:"id_prodi_asal"`
	NamaProgramStudiAsal    interface{} `json:"nama_program_studi_asal"`
	JenisKelamin            string      `json:"jenis_kelamin"`
	TanggalDaftar           string      `json:"tanggal_daftar"`
	NamaIbuKandung          string      `json:"nama_ibu_kandung"`
	IDPembiayaan            string      `json:"id_pembiayaan"`
	NamaPembiayaanAwal      string      `json:"nama_pembiayaan_awal"`
	BiayaMasuk              string      `json:"biaya_masuk"`
	IDBidangMinat           interface{} `json:"id_bidang_minat"`
	NmBidangMinat           interface{} `json:"nm_bidang_minat"`
}
