package auction_test

import (
	"fmt"
	"github.com/brendontj/sturdy-umbrella/auction"
	"github.com/brendontj/sturdy-umbrella/bidder"
	"testing"
)

func TestAuctionExecutions(t *testing.T) {
	type testCase struct {
		name               string
		bidders            []bidder.Bidder
		expectedWinnerName string
		expectedError      error
	}

	testCases := []testCase{
		{
			name:               "No bidders",
			bidders:            []bidder.Bidder{},
			expectedWinnerName: "",
			expectedError:      auction.ErrNoBidders,
		},
		{
			name: "One bidder",
			bidders: []bidder.Bidder{
				bidder.New("Bidder 1", 1_00, 10_00, 10),
			},
			expectedWinnerName: "Bidder 1",
			expectedError:      nil,
		},
		{
			name: "Tie between two bidders, increasing the bid by the same amount",
			bidders: []bidder.Bidder{
				bidder.New("Bidder 1", 1_00, 10_00, 9_00),
				bidder.New("Bidder 2", 1_00, 10_00, 9_00),
			},
			expectedWinnerName: "Bidder 2",
			//	Bidder 1 will start the auction with the best offer because it was added first,
			//	then Bidder 2 will increment the bid and reach its maximum. So, Bidder 1 will increment its amount reaching the same offer as Bidder 2 but will lose because Bidder 2 placed its bid first.			expectedError: nil,
			expectedError: nil,
		},
		{
			name: "Tie between two bidders, starting a bidder with the winner offer",
			bidders: []bidder.Bidder{
				bidder.New("Bidder 1", 9_80, 10_00, 9_00),
				bidder.New("Bidder 2", 1_00, 10_00, 4_40),
			},
			expectedWinnerName: "Bidder 1",
			//	Bidder 1 will be the winner because placed the best offer first
			expectedError: nil,
		},
		{
			name: "Case Auction #1",
			bidders: []bidder.Bidder{
				bidder.New("Sasha", 50_00, 80_00, 3_00),
				bidder.New("John", 60_00, 82_00, 2_00),
				bidder.New("Pat", 55_00, 85_00, 5_00),
			},
			expectedWinnerName: "Pat",
			expectedError:      nil,
		},
		{
			name: "Case Auction #2",
			bidders: []bidder.Bidder{
				bidder.New("Riley", 700_00, 725_00, 2_00),
				bidder.New("Morgan", 599_00, 725_00, 15_00),
				bidder.New("Charlie", 625_00, 725_00, 8_00),
			},
			expectedWinnerName: "Riley",
			expectedError:      nil,
		},
		{
			name: "Case Auction #3",
			bidders: []bidder.Bidder{
				bidder.New("Alex", 2500_00, 3000_00, 500_00),
				bidder.New("Jesse", 2800_00, 3100_00, 201_00),
				bidder.New("Drew", 2501_00, 3200_00, 247_00),
			},
			expectedWinnerName: "Jesse",
			expectedError:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := auction.New(tc.bidders)
			err := a.Run()
			if err != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}

			if tc.expectedError == nil {
				if a.GetWinner().Name != tc.expectedWinnerName {
					t.Errorf("Expected winner: %v, got: %v", tc.expectedWinnerName, a.GetWinner().Name)
				}

				t.Logf("Winner: %v, currentBid: %s",
					a.GetWinner().Name,
					fmt.Sprintf("%.2f", float64(a.GetWinner().CurrentBid())/100))
			}
		})
	}
}
