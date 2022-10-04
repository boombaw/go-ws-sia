package model

type SyncNilai struct {
	ID           int     `json:"id"`
	TahunAjaran  string  `json:"tahun_ajaran"`
	Prodi        string  `json:"prodi"`
	NPM          string  `json:"npm"`
	Name         string  `json:"name"`
	KdMatakuliah string  `json:"kd_matakuliah"`
	Bobot        float64 `json:"bobot"`
	NilaiIndeks  string  `json:"nilai_indeks"`
	NilaiAkhir   float64 `json:"nilai_akhir"`
}

type Jadwal struct {
	ID           int    `json:"id"`
	KdJadwal     string `json:"kd_jadwal"`
	KdMatakuliah string `json:"kd_matakuliah"`
	KdDosen      string `json:"kd_dosen"`
	Kelas        string `json:"kelas"`
	TahunAjaran  string `json:"tahun_ajaran"`
}

type FeederListKelas struct {
	IDKelasKuliah    string `json:"id_kelas_kuliah"`
	IDProdi          string `json:"id_prodi"`
	NamaProgramStudi string `json:"nama_program_studi"`
	IDSemester       string `json:"id_semester"`
	NamaSemester     string `json:"nama_semester"`
	IDMatkul         string `json:"id_matkul"`
	KodeMataKuliah   string `json:"kode_mata_kuliah"`
	NamaMataKuliah   string `json:"nama_mata_kuliah"`
	NamaKelasKuliah  string `json:"nama_kelas_kuliah"`
	Sks              string `json:"sks"`
	IDDosen          string `json:"id_dosen"`
	NamaDosen        string `json:"nama_dosen"`
	JumlahMahasiswa  string `json:"jumlah_mahasiswa"`
	ApaUntukPditt    string `json:"apa_untuk_pditt"`
}
