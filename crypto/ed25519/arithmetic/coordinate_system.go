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
package arithmetic /*  {packageName}  */
/**
 * Available coordinate systems for a group element.
 */
type CoordinateSystem int 
 /* public  */ 
  const ( 
 firstCoordinateSystem CoordinateSystem = iota 
 secondCoordinateSystem 
)
    /**
     * Affine coordinate system (x, y).
     */
    AFFINE,
    /**
     * Projective coordinate system (X:Y:Z) satisfying x=X/Z, y=Y/Z.
     */
    P2,
    /**
     * Extended projective coordinate system (X:Y:Z:T) satisfying x=X/Z, y=Y/Z, XY=ZT.
     */
    P3,
    /**
     * Completed coordinate system ((X:Z), (Y:T)) satisfying x=X/Z, y=Y/T.
     */
    P1xP1,
    /**
     * Precomputed coordinate system (y+x, y-x, 2dxy).
     */
    PRECOMPUTED,
    /**
     * Cached coordinate system (Y+X, Y-X, Z, 2dT).
     */
    CACHED
}
} /* CoordinateSystem */ 

