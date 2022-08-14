// Code generated by thriftgo (0.1.7). DO NOT EDIT.

package base

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type BaseResp struct {
	StatusMessage string            `thrift:"StatusMessage,1" json:"StatusMessage"`
	StatusCode    int32             `thrift:"StatusCode,2" json:"StatusCode"`
	Extra         map[string]string `thrift:"Extra,3" json:"Extra,omitempty"`
}

func NewBaseResp() *BaseResp {
	return &BaseResp{

		StatusMessage: "",
		StatusCode:    0,
	}
}

func (p *BaseResp) GetStatusMessage() (v string) {
	return p.StatusMessage
}

func (p *BaseResp) GetStatusCode() (v int32) {
	return p.StatusCode
}

var BaseResp_Extra_DEFAULT map[string]string

func (p *BaseResp) GetExtra() (v map[string]string) {
	if !p.IsSetExtra() {
		return BaseResp_Extra_DEFAULT
	}
	return p.Extra
}
func (p *BaseResp) SetStatusMessage(val string) {
	p.StatusMessage = val
}
func (p *BaseResp) SetStatusCode(val int32) {
	p.StatusCode = val
}
func (p *BaseResp) SetExtra(val map[string]string) {
	p.Extra = val
}

var fieldIDToName_BaseResp = map[int16]string{
	1: "StatusMessage",
	2: "StatusCode",
	3: "Extra",
}

func (p *BaseResp) IsSetExtra() bool {
	return p.Extra != nil
}

func (p *BaseResp) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.MAP {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_BaseResp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *BaseResp) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.StatusMessage = v
	}
	return nil
}

func (p *BaseResp) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.StatusCode = v
	}
	return nil
}

func (p *BaseResp) ReadField3(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return err
	}
	p.Extra = make(map[string]string, size)
	for i := 0; i < size; i++ {
		var _key string
		if v, err := iprot.ReadString(); err != nil {
			return err
		} else {
			_key = v
		}

		var _val string
		if v, err := iprot.ReadString(); err != nil {
			return err
		} else {
			_val = v
		}

		p.Extra[_key] = _val
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return err
	}
	return nil
}

func (p *BaseResp) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("BaseResp"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *BaseResp) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("StatusMessage", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.StatusMessage); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *BaseResp) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("StatusCode", thrift.I32, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.StatusCode); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *BaseResp) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetExtra() {
		if err = oprot.WriteFieldBegin("Extra", thrift.MAP, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Extra)); err != nil {
			return err
		}
		for k, v := range p.Extra {

			if err := oprot.WriteString(k); err != nil {
				return err
			}

			if err := oprot.WriteString(v); err != nil {
				return err
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *BaseResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BaseResp(%+v)", *p)
}

func (p *BaseResp) DeepEqual(ano *BaseResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.StatusMessage) {
		return false
	}
	if !p.Field2DeepEqual(ano.StatusCode) {
		return false
	}
	if !p.Field3DeepEqual(ano.Extra) {
		return false
	}
	return true
}

func (p *BaseResp) Field1DeepEqual(src string) bool {

	if strings.Compare(p.StatusMessage, src) != 0 {
		return false
	}
	return true
}
func (p *BaseResp) Field2DeepEqual(src int32) bool {

	if p.StatusCode != src {
		return false
	}
	return true
}
func (p *BaseResp) Field3DeepEqual(src map[string]string) bool {

	if len(p.Extra) != len(src) {
		return false
	}
	for k, v := range p.Extra {
		_src := src[k]
		if strings.Compare(v, _src) != 0 {
			return false
		}
	}
	return true
}
