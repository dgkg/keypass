package model

import "errors"

type Payloadpatch struct {
	Data map[string]interface{}
	Errs []error
}

func (p *Payloadpatch) ToString(fieldReceiver *string, fieldName string) {
	if fieldReceiver == nil {
		return
	}
	val, ok := p.Data[fieldName]
	if !ok {
		return
	}
	newval, ok := val.(string)
	if !ok {
		p.Errs = append(p.Errs, errors.New("cast not possible into string for:"+fieldName))
	}
	*fieldReceiver = newval
}
