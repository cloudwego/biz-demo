package message

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *BaseResp) IsValid() error {
	return nil
}

func (p *Message) IsValid() error {
	return nil
}

func (p *CreateMessageRequest) IsValid() error {
	if len(p.Content) <= 0 {
		return fmt.Errorf("field Title min_len rule failed, current value: %d", len(p.Content))
	}
	return nil
}

func (p *CreateMessageResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
