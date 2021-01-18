package web

func PostPseudoHandler(pReq Any, svc func(Any) (Any, error)) func(Filler) (Any, error) {
	return func(filler Filler) (Any, error) {
		err := filler(pReq)

		if err != nil {
			return nil, FillerError{err}
		}

		resp, err := svc(pReq)

		return &resp, err
	}
}
