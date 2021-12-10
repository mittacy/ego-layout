package restyHttp

import (
	"testing"
	"time"
)

type Tag struct {
	Id      int64  `json:"id" mapstructure:"id"`
	TagName string `json:"tag_name" mapstructure:"tag_name"`
}

type Tags struct {
	CommunityTags []Tag `json:"community_tags" mapstructure:"community_tags"`
	RecentlyTags  []Tag `json:"recently_tags" mapstructure:"recently_tags"`
}

func TestGet(t *testing.T) {
	host := ""
	res, _, err := Get(host, "", time.Second*5)

	if err != nil {
		t.Error(err)
	}

	tags := Tags{}
	if err := Decode(res, &tags, nil); err != nil {
		t.Error(err)
	}

	t.Logf("%+v\n", tags)
}

func TestGetParams(t *testing.T) {
	host := ""
	uri := ""
	params := map[string]string{
		"type":         "1",
		"app_id":       "appmfaz41ol6225",
		"community_id": "c_60bede024a515_VnCNlkQA2383",
	}

	res, _, err := GetParams(host, uri, params, time.Second*5)

	if err != nil {
		t.Error(err)
	}

	tags := Tags{}
	if err := Decode(res, &tags, nil); err != nil {
		t.Error(err)
	}

	t.Logf("%+v\n", tags)
}
