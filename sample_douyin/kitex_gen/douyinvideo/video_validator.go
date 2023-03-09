package douyinvideo

import (
	"bytes"
	"fmt"
	"mydouyin/pkg/consts"
	"net/url"
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

func (p *Video) IsValid() error {
	return nil
}

func (p *CreateVideoRequest) IsValid() error {
	if len(p.Title) <= 0 {
		return fmt.Errorf("field Title min_len rule failed, current value: %d", len(p.Title))
	}
	_, err1 := url.ParseRequestURI(consts.CDNURL + p.PlayUrl)
	if err1 != nil {
		return fmt.Errorf("field PlayerUrl failed, current value: %s", p.PlayUrl)
	}
	_, err2 := url.ParseRequestURI(consts.CDNURL + p.CoverUrl)
	if err2 != nil {
		return fmt.Errorf("field CoverUrl failed, current value: %s", p.CoverUrl)
	}
	return nil
}

func (p *CreateVideoResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
