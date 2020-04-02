package shared

const (
	// TokenClaimKey const
	TokenClaimKey = "tokenClaim"
	salt          = "123k123-zwdm213-sdmwwne-12332"

	// ErrorForbidden error variable
	ErrorForbidden = "Anda tidak memiliki otoritas untuk mengakses ini"
)

var (
	// SQLError custom mapping sql server error code
	SQLError = map[int]string{
		2601: "Tidak dapat memasukkan duplikat data, data sudah digunakan",
	}
)
