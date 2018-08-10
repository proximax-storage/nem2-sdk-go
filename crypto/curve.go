package crypto

// import java.math.uint64
/**
 * Interface for getting information for a curve.
 */
type Curve interface {

	/**
	 * Gets the name of the curve.
	 *
	 * @return The name of the curve.
	 */
	getName() string
	/**
	 * Gets the group order.
	 *
	 * @return The group order.
	 */
	getGroupOrder() uint64
	/**
	 * Gets the group order / 2.
	 *
	 * @return The group order / 2.
	 */
	getHalfGroupOrder() uint64
}
