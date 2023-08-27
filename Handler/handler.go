package handler

func Handle(err error) {
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}
}
