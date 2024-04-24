package bidder

import (
	"reflect"
	"testing"
)

func TestBidder_VerifyIfBidderCanIncrementItsCurrentBid(t *testing.T) {
	type fields struct {
		Name            string
		InitialBid      int
		MaxBid          int
		IncrementAmount int
		currentBid      int
		timesIncreased  int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Bidder can increment its current bid",
			fields: fields{
				InitialBid:      1_00,
				MaxBid:          2_00,
				IncrementAmount: 50,
				currentBid:      1_00,
				timesIncreased:  0,
			},
			want: true,
		},
		{
			name: "Bidder can't increment its current bid, because the currentBid + incrementAmount is greater than MaxBid",
			fields: fields{
				InitialBid:      1_00,
				MaxBid:          2_00,
				IncrementAmount: 50_00,
				currentBid:      1_00,
				timesIncreased:  0,
			},
			want: false,
		},
		{
			name: "Bidder can't increment its current bid, because the currentBid is equal MaxBid",
			fields: fields{
				InitialBid:      1_00,
				MaxBid:          2_00,
				IncrementAmount: 1,
				currentBid:      2_00,
				timesIncreased:  0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bidder{
				Name:            tt.fields.Name,
				InitialBid:      tt.fields.InitialBid,
				MaxBid:          tt.fields.MaxBid,
				IncrementAmount: tt.fields.IncrementAmount,
				currentBid:      tt.fields.currentBid,
				timesIncreased:  tt.fields.timesIncreased,
			}
			if got := b.CanIncrement(); got != tt.want {
				t.Errorf("CanIncrement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBidder_BidderTryingToIncrementCurrentBid(t *testing.T) {
	type fields struct {
		Name            string
		InitialBid      int
		MaxBid          int
		IncrementAmount int
		currentBid      int
		timesIncreased  int
	}
	tests := []struct {
		name               string
		fields             fields
		expectedCurrentBid int
	}{
		{
			name: "Bidder increments the current bid with success",
			fields: fields{
				InitialBid:      1_00,
				MaxBid:          2_00,
				IncrementAmount: 1,
				currentBid:      1_00,
				timesIncreased:  0,
			},
			expectedCurrentBid: 1_01,
		},
		{
			name: "Bidder doesn't increment the current bid because reached the max bid",
			fields: fields{
				InitialBid:      1_00,
				MaxBid:          2_00,
				IncrementAmount: 1,
				currentBid:      2_00,
				timesIncreased:  100,
			},
			expectedCurrentBid: 2_00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bidder{
				Name:            tt.fields.Name,
				InitialBid:      tt.fields.InitialBid,
				MaxBid:          tt.fields.MaxBid,
				IncrementAmount: tt.fields.IncrementAmount,
				currentBid:      tt.fields.currentBid,
				timesIncreased:  tt.fields.timesIncreased,
			}
			b.IncrementCurrentBid()

			if got := b.CurrentBid(); got != tt.expectedCurrentBid {
				t.Errorf("CurrentBid() = %v, want %v", got, tt.expectedCurrentBid)
			}
		})
	}
}

func TestBidder_CreatingNewBidder(t *testing.T) {
	type args struct {
		name            string
		initialBid      int
		maxBid          int
		incrementAmount int
	}
	tests := []struct {
		name    string
		args    args
		want    *Bidder
		wantErr bool
		err     error
	}{
		{
			name: "Should create a new bidder with success",
			args: args{
				initialBid:      1_00,
				maxBid:          10_00,
				incrementAmount: 1,
			},
			want: &Bidder{
				InitialBid:      1_00,
				MaxBid:          10_00,
				IncrementAmount: 1,
				currentBid:      1_00,
				timesIncreased:  0,
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Shouldn't create a new bidder because the initial Bid is greater than the max Bid",
			args: args{
				initialBid:      11_00,
				maxBid:          10_00,
				incrementAmount: 1,
			},
			want:    nil,
			wantErr: true,
			err:     ErrInitialBidIsGreaterThanMaxBid,
		},
		{
			name: "Shouldn't create a new bidder because every bid value should be greater than zero",
			args: args{
				initialBid:      1_00,
				maxBid:          10_00,
				incrementAmount: -1,
			},
			want:    nil,
			wantErr: true,
			err:     ErrInvalidBidValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.name, tt.args.initialBid, tt.args.maxBid, tt.args.incrementAmount)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if (err != nil) && err != tt.err {
				t.Errorf("New() error = %v, expectedErr %v", err, tt.err)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
