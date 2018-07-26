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
package sdk

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

type NamespaceService service

type NamespaceId struct { /* public  */

	id       int64
	fullName string
} /* NamespaceId */
func NewNamespaceId(id int64, namespaceName string) (ref *NamespaceId, err error) { /* public  */
	if namespaceName == "" {
		return nil, errors.New("nameSpace is empty")
	}
	if id < 0 {
		id = generateNamespaceId(namespaceName)
	}
	ref = &NamespaceId{id, namespaceName}
	return
}

type NamespaceType int

const (
	RootNamespace NamespaceType = iota
	SubNamespace
) /* NamespaceType */
//NamespaceInfo contains the state information of a namespace.
type NamespaceInfo struct { /* public  */

	active bool

	id int

	metaId string

	typeSpace NamespaceType

	depth int

	levels []NamespaceId /* List */

	parentId NamespaceId

	owner PublicAccount

	startHeight int64

	endHeight int64
} /* NamespaceInfo */
func NewNamespaceInfo(active bool, id int, metaId string, typeSpace NamespaceType, depth int, levels []NamespaceId,
	parentId NamespaceId, owner PublicAccount, startHeight int64, endHeight int64) *NamespaceInfo { /* public  */
	ref := &NamespaceInfo{
		active,
		id,
		metaId,
		typeSpace,
		depth,
		levels,
		parentId,
		owner,
		startHeight,
		endHeight,
	}
	return ref
}
func (ref *NamespaceService) GetNameSpaceInfo(ctx context.Context, nsId int) (nsInfo *NamespaceInfo, resp *http.Response, err error) {
	var req *http.Request

	req, err = ref.client.NewRequest("GET", fmt.Sprintf("/namespace/%d", nsId), nil)

	if err == nil {
		nsInfo = &NamespaceInfo{}
		resp, err = ref.client.Do(ctx, req, nsInfo)

		if err == nil {
			if nsInfo.id == nsId {
				return
			}
			err = errors.New("nod valid id returns")
		}
	}
	//	err occurent
	return nil, nil, err
}
func (ref *NamespaceService) GetAccountNameSpaceInfo(ctx context.Context, nsId int, pageSize int, addresses Addresses) (nsInfo *NamespaceInfo, resp *http.Response, err error) {
	var req *http.Request

	body, _ := addresses.MarshalJSON()

	req, err = ref.client.NewRequest("POST", fmt.Sprintf("account/namespace/%d/%d", nsId, pageSize), &body)

	if err == nil {
		nsInfo = &NamespaceInfo{}
		resp, err = ref.client.Do(ctx, req, nsInfo)

		if err == nil {
			if nsInfo.id == nsId {
				return
			}
			err = errors.New("nod valid id returns")
		}
	}
	//	err occurent
	return
}
func generateNamespaceId(namespaceName string) int64 { /* public static  */

	namespacePath := generateNamespacePath(namespaceName)
	return namespacePath[len(namespacePath)-1]
}
func generateId(name string, parentId int64) int64 { /* public static  */

	//var parentIdBytes byte[8]
	//bytes.Buffer.wrap(parentIdBytes).put(parentId.toByteArray()); // GO
	//ArrayUtils.reverse(parentIdBytes)
	//[]byte bytes = name.getBytes()
	//[]byte result = Hashes.sha3_256(parentIdBytes, bytes)
	//[]byte low = Arrays.copyOfRange(result, 0, 4)
	//[]byte high = Arrays.copyOfRange(result, 4, 8)
	//ArrayUtils.reverse(low)
	//ArrayUtils.reverse(high)
	//[]byte last = ArrayUtils.addAll(high, low)
	return -1 //NewBigInteger(last)
}
func generateNamespacePath(name string) []int64 { /* public static  */

	//[]string parts = name.split(Pattern.quote("."))
	//[]int64 /* List */ path = new Array[]int64 /* List */()
	//if (parts.length == 0) {
	//	panic(IllegalIdentifierException{"invalid namespace name"})
	//}
	//else if (parts.length > 3) {
	//	panic(IllegalIdentifierException{"too many parts"})
	//}
	//
	//int64 namespaceId = int64.valueOf(0)
	//for (int i = 0; i < parts.length; i++) {
	//	if (!parts[i].matches("^[a-z0-9][a-z0-9-_]*$")) {
	//		panic(IllegalIdentifierException{"invalid namespace name"})
	//	}
	//
	//	namespaceId = generateId(parts[i], namespaceId)
	//	path.add(namespaceId)
	//}

	return []int64{0, 1}
}
