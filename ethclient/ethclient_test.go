// Copyright 2016 The go-sferum Authors
// This file is part of the go-sferum library.
//
// The go-sferum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-sferum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-sferum library. If not, see <http://www.gnu.org/licenses/>.

package ethclient

import "github.com/codesferum/go-sferum"

// Verify that Client implements the sferum interfaces.
var (
	_ = sferum.ChainReader(&Client{})
	_ = sferum.TransactionReader(&Client{})
	_ = sferum.ChainStateReader(&Client{})
	_ = sferum.ChainSyncReader(&Client{})
	_ = sferum.ContractCaller(&Client{})
	_ = sferum.GasEstimator(&Client{})
	_ = sferum.GasPricer(&Client{})
	_ = sferum.LogFilterer(&Client{})
	_ = sferum.PendingStateReader(&Client{})
	// _ = sferum.PendingStateEventer(&Client{})
	_ = sferum.PendingContractCaller(&Client{})
)
