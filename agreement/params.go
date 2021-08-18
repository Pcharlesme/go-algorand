// Copyright (C) 2019-2021 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package agreement

import (
	"github.com/algorand/go-algorand/data/basics"
	"github.com/algorand/go-algorand/data/bookkeeping"
)

// ParamsRound returns the round from which consensus parameters
// should be used for agreement on round rnd.
func ParamsRound(rnd basics.Round) basics.Round {
	return rnd.SubSaturate(2)
}

// paramsRoundBranch returns the target round and leaf branch from which consensus parameters
// should be used for agreement on round rnd.
func paramsRoundBranch(rnd round) (basics.Round, bookkeeping.BlockHash) {
	return ParamsRound(rnd.Number), rnd.Branch
}
