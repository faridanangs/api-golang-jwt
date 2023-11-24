package helper

func PanicError(err error) {

	if err != nil {
		// logrus.WithField("msg", msg).Fatal(err) jika gunakan ini server akan berhenti dan
		// menampilkan errornya di terminal namun jika kita gunakan panic errornya akan di tampilkan di
		// dalam json
		panic(err)
	}
}
