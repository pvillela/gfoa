package web

func PostPseudoHandler(pReq Any, svc func(Any) Any) func(Filler) (Any, error) {
	return func(filler Filler) (Any, error) {
		err := filler(pReq)

		if err != nil {
			return nil, err
		}

		resp := svc(pReq)

		return &resp, nil
	}
}
