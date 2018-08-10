
package ed25519
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //cryptoengines"
// import org.hamcrest.core.IsEqual
// import org.junit.Assert
// import org.junit.Test
// import java.math.uint64 
type Ed25519CurveTest struct {
      
    GROUP_ORDER = uint64.ONE.shiftLeft(252).add(NewBigInteger("27742317777372353535851937790883648493")) uint64 // private static
// @Test
   func (ref *Ed25519CurveTest) GetNameReturnsCorrectName()    {

        // Assert:
        Assert.assertThat(CryptoEngines.ed25519Engine().Curve.getName(), IsEqual.equalTo(" {package ed25519 /* Name} ")) */
}
} /* Ed25519CurveTest */ 

// @Test
   func (ref *Ed25519CurveTest) GetNameReturnsCorrectGroupOrder()    {

        // Assert:
        Assert.assertThat(CryptoEngines.ed25519Engine().Curve.getGroupOrder(), IsEqual.equalTo(GROUP_ORDER)) 
}

// @Test
   func (ref *Ed25519CurveTest) GetNameReturnsCorrectHalfGroupOrder()    {

        // Arrange:
        halfGroupOrder = GROUP_ORDER.shiftRight(1) uint64
        // Assert:
        Assert.assertThat(CryptoEngines.ed25519Engine().Curve.getHalfGroupOrder(), IsEqual.equalTo(halfGroupOrder)) 
}

}

