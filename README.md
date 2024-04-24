# sturdy-umbrella

The idea of this project is to create a module called auction to calculate the bidder winner in a possible auction.

## Table of Contents

- [Requirements](#requirements)
- [Usage](#Usage)

## Requirements

- Go 1.21 or later

## Usage

Basically to use this module you need to import it in your code and instantiate a new auction passing a list of bidders 
and use the method `Run` to execute the auction and get the winner.

The code utilizes a Heap data structure to store bidders and their bids. This allows for the retrieval of the lowest bid in O(1) time complexity. If possible, the bidder's bid can be increased and stored in the heap again in O(log n) time complexity. If a higher bid cannot be made, the bidder is removed as we already have a better bid. This process continues until there is only one bidder left in the heap, who is declared the winner. In the event of a tie between multiple bids, the order in which the bids were placed is taken into account.


Some tests cases were provided in the `auction_test.go` file to demonstrate the usage of the module.