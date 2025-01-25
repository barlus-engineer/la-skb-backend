package text

var (
	AppVersion = "0.1"

	HelloWorld = "ສະບາຍດີ, ໂລກ!"

	BindBodyWithJSONError = "ບໍ່ສາມາດແປ່ງ JSON ໄດ້, ຂໍ້ມູນ JSON ອາດຈະບໍ່ຖືກຕ້ອງ"

	JWTRefreshErrorGen = "ມີບັນຫາໃນການສ້າງ token (r)"
	JWTAccessErrorGen  = "ມີບັນຫາໃນການສ້າງ token (a)"

	SignUpUserExists     = "ມີຊື່ຜູ້ໃຊ້ດັັງກ່າວໃນລະບົບແລ້ວ, ກະລຸນາຕັ້ງຊື່ຜູ້ໃຊ້ໃໝ່"
	SignUpError          = "ມີບັນຫາໃນການລົງທະບຽນ"
	SignUpSuccess        = "ລົງທະບຽນສຳເລັດແລ້ວ"
	SignUpIncompleteForm = "ຂໍ້ມູນບໍ່ຄົບ, ກະລຸນາປ້ອນຂໍ້ມູນໃສ່ຟອມໃຫ້ຄົບ"

	SignInUserNotFound      = "ບໍ່ມີບັນຊີຜູ້ໃຊ້ດັງກ່າວໃນລະບົບ"
	SignInServerError       = "ເຊີບເວີມີບັນຫາໃນການເຂົ້າສູ່ລະບົບ"
	SignInPasswordIncorrect = "ລະຫັດຜ່ານບໍ່ຖືກຕ້ອງ"
	SignInSuccess           = "ເຂົ້າສູ່ລະບົບສຳເລັດແລ້ວ"
	SignInIncompleteForm    = "ຂໍ້ມູນບໍ່ຄົບ, ກະລຸນາປ້ອນຂໍ້ມູນໃສ່ຟອມໃຫ້ຄົບ"

	SignOutSuccess = "ອອກຈາກລະບົບສຳເລັດແລ້ວ"
	SignOutError   = "ມີບັນຫາໃນການອອກຈາກລະບົບ"

	DeleteAccountUsernameNotFound  = "ບໍ່ມີບັນຊີຜູ້ໃຊ້ດັງກ່າວໃນລະບົບ"
	DeleteAccountServerError       = "ເຊີບເວີມີບັນຫາໃນການລົບບັນຊີ"
	DeleteAccountPasswordIncorrect = "ລະຫັດຜ່ານບໍ່ຖືກຕ້ອງ"
	DeleteAccountError             = "ມີບັນຫາໃນການລົບບັນຊີຜູ້ໃຊ້"
	DeleteAccountSuccess           = "ລົບບັນຊີຜູ້ໃຊ້ສຳເລັດແລ້ວ"
	DeleteAccountIncompleteForm    = "ຂໍ້ມູນບໍ່ຄົບ, ກະລຸນາປ້ອນຂໍ້ມູນໃສ່ຟອມໃຫ້ຄົບ"
)
