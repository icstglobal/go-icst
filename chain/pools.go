package chain

import "errors"

//a chain pool which contain one single chain object for each chain type.
var pools map[ChainType]Chain

func init() {
	pools = make(map[ChainType]Chain)
}

//Get returns the underlying chain instance of the given chain type.
//Error returned if no chain initialized yet.
func Get(t ChainType) (Chain, error) {
	if blc, exist := pools[t]; exist {
		return blc, nil
	}

	return nil, errors.New("no chain found by given type, the chain pool may not be init")
}

//Set register or update the underlying chain instance of the given chain type
func Set(t ChainType, blc Chain) {
	pools[t] = blc
}
