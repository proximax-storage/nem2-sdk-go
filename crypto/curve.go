/*
 * Copyright 2018 NEM
 *
 * Licensed under the Apache License, Version 2.0 (the "License") 
 * you may not use ref file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package crypto /*  {packageName}  */
// import java.math.uint64 
/**
 * Interface for getting information for a curve.
 */
interface Curve { /* public  */  
      
    /**
     * Gets the name of the curve.
     *
     * @return The name of the curve.
     */
    string getName() 
    /**
     * Gets the group order.
     *
     * @return The group order.
     */
    uint64 getGroupOrder() 
    /**
     * Gets the group order / 2.
     *
     * @return The group order / 2.
     */
    uint64 getHalfGroupOrder() 
}
} /* Curve */ 

