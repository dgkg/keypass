package model

import "errors"

type Payloadpatch struct {
	Data map[string]interface{}
	Errs []error
}

func (p *Payloadpatch) ToString(fieldName string) string {
	val, ok := p.Data[fieldName]
	if !ok {
		p.Errs = append(p.Errs, errors.New("no values for:"+fieldName))
		return ""
	}
	newval, ok := val.(string)
	if !ok {
		p.Errs = append(p.Errs, errors.New("cast not possible into string for:"+fieldName))
		return ""
	}
	return newval
}
