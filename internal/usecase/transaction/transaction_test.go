package transaction

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/asetriza/transaction-tracker/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestTransactionService_Create(t *testing.T) {
	type fields struct {
		trRepo    TransactionRepository
		acService AccountService
	}
	type args struct {
		ctx         context.Context
		transaction domain.Transaction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   func(t *testing.T, tr domain.Transaction, err error)
	}{
		{
			name: "validation error",
			fields: fields{
				trRepo:    nil,
				acService: nil,
			},
			args: args{
				ctx: context.Background(),
				transaction: domain.Transaction{
					ID:            0,
					TransactionID: "",
					AccountID:     0,
					State:         domain.StateLost,
					Amount:        0.0,
					IsCanceled:    false,
					CreatedAt:     time.Time{},
				},
			},
			want: func(t *testing.T, tr domain.Transaction, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "transaction repository error",
			fields: fields{
				trRepo: MockTransactionRepository{
					CreateFunc: func(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
						return domain.Transaction{}, errors.New("unknown error")
					},
				},
				acService: nil,
			},
			args: args{
				ctx: context.Background(),
				transaction: domain.Transaction{
					ID:            0,
					TransactionID: "some id",
					AccountID:     0,
					State:         domain.StateLost,
					Amount:        1.0,
					IsCanceled:    false,
					CreatedAt:     time.Time{},
				},
			},
			want: func(t *testing.T, tr domain.Transaction, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "success",
			fields: fields{
				trRepo: MockTransactionRepository{
					CreateFunc: func(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
						return domain.Transaction{}, nil
					},
				},
				acService: nil,
			},
			args: args{
				ctx: context.Background(),
				transaction: domain.Transaction{
					ID:            0,
					TransactionID: "some id",
					AccountID:     0,
					State:         domain.StateLost,
					Amount:        1.0,
					IsCanceled:    false,
					CreatedAt:     time.Time{},
				},
			},
			want: func(t *testing.T, tr domain.Transaction, err error) {
				require.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TransactionService{
				trRepo:    tt.fields.trRepo,
				acService: tt.fields.acService,
			}
			got, err := tr.Create(tt.args.ctx, tt.args.transaction)
			tt.want(t, got, err)
		})
	}
}
