package wns

type raw struct {
	Body string
}

func (r *raw) SetBody(str string) *raw {
	r.Body = str
	return r
}

func (r *raw) GetXml() (string, error) {
	return r.Body, nil
}

func (r *raw) GetWnsType() string {
	return "wns/raw"
}
