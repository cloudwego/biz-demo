package douyincomment

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

func (p *CreateCommentRequest) IsValid() error {
	if len(p.Content) < int(1) {
		return fmt.Errorf("field Content min_len rule failed, current value: %d", len(p.Content))
	}
	if len(p.CreateDate) < int(1) {
		return fmt.Errorf("field CreateDate min_len rule failed, current value: %d", len(p.CreateDate))
	}
	return nil
}
