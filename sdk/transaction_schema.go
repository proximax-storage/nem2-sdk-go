// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

func aggregateTransactionSchema() *schema {
	return &schema{
		[]schemaAttribute{
			newScalarAttribute("size", IntSize),
			newArrayAttribute("signature", ByteSize),
			newArrayAttribute("signer", ByteSize),
			newScalarAttribute("version", ShortSize),
			newScalarAttribute("type", ShortSize),
			newArrayAttribute("fee", IntSize),
			newArrayAttribute("deadline", IntSize),
			newScalarAttribute("transactionsSize", IntSize),
			newArrayAttribute("transactions", ByteSize),
		},
	}
}

func mosaicDefinitionTransactionSchema() *schema {
	return &schema{
		[]schemaAttribute{
			newScalarAttribute("size", IntSize),
			newArrayAttribute("signature", ByteSize),
			newArrayAttribute("signer", ByteSize),
			newScalarAttribute("version", ShortSize),
			newScalarAttribute("type", ShortSize),
			newArrayAttribute("fee", IntSize),
			newArrayAttribute("deadline", IntSize),
			newArrayAttribute("parentId", IntSize),
			newArrayAttribute("mosaicId", IntSize),
			newScalarAttribute("mosaicNameLength", ByteSize),
			newScalarAttribute("numOptionalProperties", ByteSize),
			newScalarAttribute("flags", ByteSize),
			newScalarAttribute("divisibility", ByteSize),
			newArrayAttribute("mosaicName", ByteSize),
			newScalarAttribute("indicateDuration", ByteSize),
			newArrayAttribute("duration", IntSize),
		},
	}
}

func mosaicSupplyChangeTransactionSchema() *schema {
	return &schema{
		[]schemaAttribute{
			newScalarAttribute("size", IntSize),
			newArrayAttribute("signature", ByteSize),
			newArrayAttribute("signer", ByteSize),
			newScalarAttribute("version", ShortSize),
			newScalarAttribute("type", ShortSize),
			newArrayAttribute("fee", IntSize),
			newArrayAttribute("deadline", IntSize),
			newArrayAttribute("mosaicId", IntSize),
			newScalarAttribute("direction", ByteSize),
			newArrayAttribute("delta", IntSize),
		},
	}
}

func transferTransactionSchema() *schema {
	return &schema{
		[]schemaAttribute{
			newScalarAttribute("size", IntSize),
			newArrayAttribute("signature", ByteSize),
			newArrayAttribute("signer", ByteSize),
			newScalarAttribute("version", ShortSize),
			newScalarAttribute("type", ShortSize),
			newArrayAttribute("fee", IntSize),
			newArrayAttribute("deadline", IntSize),
			newArrayAttribute("recipient", ByteSize),
			newScalarAttribute("messageSize", ShortSize),
			newScalarAttribute("numMosaics", ByteSize),
			newTableAttribute("message", schema{
				[]schemaAttribute{
					newScalarAttribute("type", ByteSize),
					newArrayAttribute("payload", ByteSize),
				},
			}.schemaDefinition),
			newTableArrayAttribute("mosaics", schema{
				[]schemaAttribute{
					newArrayAttribute("id", IntSize),
					newArrayAttribute("amount", IntSize),
				},
			}.schemaDefinition),
		},
	}
}
func modifyMultisigAccountTransactionSchema() *schema {
	return &schema{
		[]schemaAttribute{
			newScalarAttribute("size", IntSize),
			newArrayAttribute("signature", ByteSize),
			newArrayAttribute("signer", ByteSize),
			newScalarAttribute("version", ShortSize),
			newScalarAttribute("type", ShortSize),
			newArrayAttribute("fee", IntSize),
			newArrayAttribute("deadline", IntSize),
			newScalarAttribute("minRemovalDelta", ByteSize),
			newScalarAttribute("minApprovalDelta", ByteSize),
			newScalarAttribute("numModifications", ByteSize),
			newTableArrayAttribute("modification", schema{
				[]schemaAttribute{
					newScalarAttribute("type", ByteSize),
					newArrayAttribute("cosignatoryPublicKey", ByteSize),
				},
			}.schemaDefinition),
		},
	}
}

func registerNamespaceTransactionSchema() *schema {
	return &schema{
		[]schemaAttribute{
			newScalarAttribute("size", IntSize),
			newArrayAttribute("signature", ByteSize),
			newArrayAttribute("signer", ByteSize),
			newScalarAttribute("version", ShortSize),
			newScalarAttribute("type", ShortSize),
			newArrayAttribute("fee", IntSize),
			newArrayAttribute("deadline", IntSize),
			newScalarAttribute("namespaceType", ByteSize),
			newArrayAttribute("durationParentId", IntSize),
			newArrayAttribute("namespaceId", IntSize),
			newScalarAttribute("namespaceNameSize", ByteSize),
			newArrayAttribute("name", ByteSize),
		},
	}
}

func lockFundsTransactionSchema() *schema {
	return &schema{
		[]schemaAttribute{
			newScalarAttribute("size", IntSize),
			newArrayAttribute("signature", ByteSize),
			newArrayAttribute("signer", ByteSize),
			newScalarAttribute("version", ShortSize),
			newScalarAttribute("type", ShortSize),
			newArrayAttribute("fee", IntSize),
			newArrayAttribute("deadline", IntSize),
			newArrayAttribute("mosaicId", IntSize),
			newArrayAttribute("mosaicAmount", IntSize),
			newArrayAttribute("duration", IntSize),
			newArrayAttribute("hash", ByteSize),
		},
	}
}

func secretLockTransactionSchema() *schema {
	return &schema{
		[]schemaAttribute{
			newScalarAttribute("size", IntSize),
			newArrayAttribute("signature", ByteSize),
			newArrayAttribute("signer", ByteSize),
			newScalarAttribute("version", ShortSize),
			newScalarAttribute("type", ShortSize),
			newArrayAttribute("fee", IntSize),
			newArrayAttribute("deadline", IntSize),
			newArrayAttribute("mosaicId", IntSize),
			newArrayAttribute("mosaicAmount", IntSize),
			newArrayAttribute("duration", IntSize),
			newScalarAttribute("hashAlgorithm", ByteSize),
			newArrayAttribute("secret", ByteSize),
			newArrayAttribute("recipient", ByteSize),
		},
	}
}

func secretProofTransactionSchema() *schema {
	return &schema{
		[]schemaAttribute{
			newScalarAttribute("size", IntSize),
			newArrayAttribute("signature", ByteSize),
			newArrayAttribute("signer", ByteSize),
			newScalarAttribute("version", ShortSize),
			newScalarAttribute("type", ShortSize),
			newArrayAttribute("fee", IntSize),
			newArrayAttribute("deadline", IntSize),
			newScalarAttribute("hashAlgorithm", ByteSize),
			newArrayAttribute("secret", ByteSize),
			newScalarAttribute("proofSize", ShortSize),
			newArrayAttribute("proof", ByteSize),
		},
	}
}
