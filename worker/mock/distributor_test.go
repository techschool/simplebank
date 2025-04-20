package mockwk

import (
    "context"
    "testing"
    "github.com/golang/mock/gomock"
    "github.com/techschool/simplebank/worker"
    "github.com/hibiken/asynq"
)


// Test generated using Keploy
func TestDistributeTaskSendVerifyEmail_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDistributor := NewMockTaskDistributor(ctrl)
    ctx := context.Background()
    payload := &worker.PayloadSendVerifyEmail{Username: "testuser"}
    options := []asynq.Option{}

    mockDistributor.EXPECT().
        DistributeTaskSendVerifyEmail(ctx, payload, options).
        Return(nil)

    err := mockDistributor.DistributeTaskSendVerifyEmail(ctx, payload, options...)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
}

// Test generated using Keploy
func TestDistributeTaskSendVerifyEmail_WithOptions(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDistributor := NewMockTaskDistributor(ctrl)
    ctx := context.Background()
    payload := &worker.PayloadSendVerifyEmail{Username: "testuser"}
    option1 := asynq.MaxRetry(5)
    option2 := asynq.Queue("critical")

    mockDistributor.EXPECT().
        DistributeTaskSendVerifyEmail(ctx, payload, option1, option2).
        Return(nil)

    err := mockDistributor.DistributeTaskSendVerifyEmail(ctx, payload, option1, option2)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
}

