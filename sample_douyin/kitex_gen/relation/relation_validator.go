package relation

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

func (p *CreateRelationRequest) IsValid() error {
	return nil
}

func (p *DeleteRelationRequest) IsValid() error {
	return nil
}

func (p *ValidIfFollowRequest) IsValid() error {
	return nil
}

func (p *GetFollowListRequest) IsValid() error {
	return nil
}

func (p *GetFollowerListRequest) IsValid() error {
	return nil
}
