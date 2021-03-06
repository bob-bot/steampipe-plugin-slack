package slack

import (
	"context"
	"errors"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/slack-go/slack"

	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func connect(_ context.Context) (*slack.Client, error) {
	token, ok := os.LookupEnv("SLACK_TOKEN")
	if !ok || token == "" {
		return nil, errors.New("SLACK_TOKEN environment variable must be set")
	}
	api := slack.New(token, slack.OptionDebug(false))
	return api, nil
}

func stringFloatToTime(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	timeFloat, err := strconv.ParseFloat(d.Value.(string), 64)
	if err != nil {
		return nil, err
	}
	if timeFloat == 0 {
		return nil, nil
	}
	sec, dec := math.Modf(timeFloat)
	t := time.Unix(int64(sec), int64(dec*(1e9)))
	return t, nil
}

func intToTime(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	i := int64(d.Value.(int))
	// Assume zero value (1970-01-01 00:00:00) means not set (null).
	if i == 0 {
		return nil, nil
	}
	return time.Unix(i, 0), nil
}

func jsonTimeToTime(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	jt := d.Value.(slack.JSONTime)
	// Assume zero value (1970-01-01 00:00:00) means not set (null).
	if jt == 0 {
		return nil, nil
	}
	return jt.Time(), nil
}
