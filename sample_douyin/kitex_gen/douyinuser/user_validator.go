package douyinuser

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

func (p *User) IsValid() error {
	return nil
}

func (p *CreateUserRequest) IsValid() error {
	if len(p.Username) < int(1) {
		return fmt.Errorf("field Username min_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) < int(1) {
		return fmt.Errorf("field Password min_len rule failed, current value: %d", len(p.Password))
	}
	return nil
}

func (p *CreateUserResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}

func (p *CheckUserRequest) IsValid() error {
	if len(p.Username) < int(1) {
		return fmt.Errorf("field Username min_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) < int(1) {
		return fmt.Errorf("field Password min_len rule failed, current value: %d", len(p.Password))
	}
	return nil
}

func (p *CheckUserResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}

func (p *MGetUserRequest) IsValid() error {
	if len(p.UserIds) < int(1) {
		return fmt.Errorf("field UserIds MinLen rule failed, current value: %v", p.UserIds)
	}
	return nil
}

func (p *MGetUserResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
