package crypto

import "math/big"

//Curve  Interface for getting information for a curve.
type Curve interface {

	/**
	 * Gets the name of the curve.
	 *
	 * @return The name of the curve.
	 */
	GetName() string
	/**
	 * Gets the group order.
	 *
	 * @return The group order.
	 */
	GetGroupOrder() *big.Int
	/**
	 * Gets the group order / 2.
	 *
	 * @return The group order / 2.
	 */
	GetHalfGroupOrder() uint64
}
