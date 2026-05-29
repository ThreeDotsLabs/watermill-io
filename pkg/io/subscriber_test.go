package io_test

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
	watermillio "github.com/ThreeDotsLabs/watermill-io/pkg/io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubscriber_Subscribe(t *testing.T) {
	t.Run("close channel when nothing is read but context is done", func(t *testing.T) {
		var buf bytes.Buffer

		subscriber, err := watermillio.NewSubscriber(io.NopCloser(&buf), watermillio.SubscriberConfig{
			UnmarshalFunc: watermillio.PayloadUnmarshalFunc,
		}, watermill.NopLogger{})
		require.NoError(t, err)

		ctx, cancel := context.WithCancel(t.Context())
		cancel()

		messages, err := subscriber.Subscribe(ctx, "topic")
		require.NoError(t, err)

		_, ok := <-messages
		assert.False(t, ok, "channel should be closed")
	})
}
