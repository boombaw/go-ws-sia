package query

var SelectAKM = `SELECT
						takm.id,
						takm.NIMHSTRAKM npm,
						npm_to_nama_mhs(takm.NIMHSTRAKM) name,
						takm.NLIPSTRAKM ips,
						takm.NLIPKTRAKM ipk,
						takm.SKSEMTRAKM sks,
						takm.SKSTTTRAKM total_sks,
						takm.BIAYA biaya 
					FROM
						tbl_aktifitas_kuliah_mahasiswa takm WHERE takm.KDPSTTRAKM  = ? AND takm.THSMSTRAKM = ?`

var SelectNA = `SELECT
					DISTINCT tm.NIMHSMSMHS npm,
					tm.NMMHSMSMHS name
				FROM
					tbl_mahasiswa tm
				WHERE
					tm.KDPSTMSMHS = ?
					AND tm.NIMHSMSMHS NOT IN (SELECT npm_mahasiswa FROM tbl_verifikasi_krs tvk WHERE tvk.tahunajaran = ? AND kd_jurusan = ? )
					AND tm.NIMHSMSMHS NOT IN (SELECT npm FROM tbl_status_mahasiswa tsm WHERE tsm.tahunajaran = ? AND validate = 1 )
					AND tm.STMHSMSMHS IN ('A', 'N')
					AND tm.SMAWLMSMHS >= ? AND tm.SMAWLMSMHS <= ?
					AND tm.NIMHSMSMHS NOT IN ('201810115011','201810115024','201810115026','201810115076','201810115298');`

var SelectLastAKM = `SELECT 
						takm.id,
						takm.NIMHSTRAKM npm,
						npm_to_nama_mhs(takm.NIMHSTRAKM) name,
						takm.NLIPSTRAKM ips,
						takm.NLIPKTRAKM ipk,
						takm.SKSEMTRAKM sks,
						takm.SKSTTTRAKM total_sks,
						takm.BIAYA biaya 
					FROM 
						tbl_aktifitas_kuliah_mahasiswa takm 
					WHERE takm.NIMHSTRAKM = ? ORDER BY THSMSTRAKM DESC LIMIT 1`

var SmsProdi = `SELECT 
					id_sms 
				FROM 
					tbl_jurusan_prodi tjp
				WHERE tjp.kd_prodi = ?`

var UpdateHasSync = `UPDATE
						siakadonline.tbl_aktifitas_kuliah_mahasiswa
					SET
						has_sync = 1
					WHERE
						id = ? `

var GetActYear = `SELECT kode FROM tbl_tahunakademik WHERE status = 1`
