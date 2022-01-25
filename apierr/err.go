package apierr

type BizErr struct {
	Code int
	Msg  string
}

func (b *BizErr) Error() string {
	return b.Msg
}

func ErrCode(err error) int {
	switch t := err.(type) {
	case *BizErr:
		return t.Code
	default:
		return Param.Code
	}
}
