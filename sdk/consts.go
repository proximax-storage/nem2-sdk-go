package sdk

import "regexp"

// const routers path for methods NamespaceService
const (
	pathNamespacesFromAccounts = "/account/namespaces"
	pathNamespace              = "/namespace/"
	pathNamespacenames         = "/namespace/names"
	pathNamespacesFromAccount  = "/account/%s/namespaces"
)

// const routers path for methods MosaicService
const (
	pathMosaic              = "/mosaic/"
	pathMosaicNames         = "/mosaic/names"
	pathMosaicFromNamespace = "/namespace/%s/mosaics/"
)

const tplNamespaceInfo = `"active": %v,
    "index": %d,
    "id": "%s",
	"type": %d,
    "depth": %d,
    "levels": [
      %v
    ],
    "parentId": [
      %v
    ],
    "owner": "%v",
    "ownerAddress": "%s",
    "startHeight": [
      %v
    ],
    "endHeight": [
      %v
    ]
  }
`

type NamespaceType uint8

const (
	Root NamespaceType = iota
	Sub
)

// regValidNamespace check namespace on valid symbols
var regValidNamespace = regexp.MustCompile(`^[a-z0-9][a-z0-9-_]*$`)
