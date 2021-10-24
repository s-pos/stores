package constants

type Code string

const (
	Error = "Terjadi kesalahan pada server, silakan coba beberapa saat lagi"

	SuccessCreateStore Code = "150101"
	FailedCreatedStore Code = "150140"

	ImageNotEncoded   Code = "153080"
	ImageUploadFailed Code = "153081"

	BodyRequired            Code = "154040"
	FailedUpdateFile        Code = "154041"
	StoreNameAlreadyTaken   Code = "154042"
	StoreDomainAlreadyTaken Code = "154043"
	StoreImageNotFound      Code = "154044"

	ErrorInsert Code = "159080"
	ErrorFetch  Code = "159081"
	ErrorUpdate Code = "159082"
	ErrorDelete Code = "159083"

	ErrorMarshal   Code = "159090"
	ErrorUnmarshal Code = "159091"
)

var (
	Message = map[Code]string{
		SuccessCreateStore: "store.create.success",
		FailedCreatedStore: "store.create.failed",
	}

	Reason = map[Code]string{
		ImageNotEncoded:   "Gagal decode foto",
		ImageUploadFailed: "Upload foto gagal",

		BodyRequired:            "Permintaan tidak sesuai",
		FailedUpdateFile:        "Terjadi kesalahan saat unggah berkas, silahkan coba beberapa saat lagi",
		StoreDomainAlreadyTaken: "Username toko sudah digunakan",
		StoreNameAlreadyTaken:   "Nama toko sudah digunakan",
		StoreImageNotFound:      "Foto toko tidak ditemukan",
	}
)
