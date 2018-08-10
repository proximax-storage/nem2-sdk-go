
package ed25519
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //keyanalyzer"
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //publickey"
/**
 * Implementation of the key analyzer for Ed25519.
 */
type Ed25519KeyAnalyzer struct {
    KeyAnalyzer
  
    COMPRESSED_KEY_SIZE = 32 int // private static
// @Override
   func (ref *Ed25519KeyAnalyzer) IsKeyCompressed(PublicKey publicKey ) bool {

       Return COMPRESSED_KEY_SIZE == publicKey.Raw.length 
}
} /* Ed25519KeyAnalyzer */ 

}

