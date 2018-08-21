package sdk

var abstractTransactionSchemaAttributes = []schemaAttribute{
	newScalarAttribute("size", IntSize),
	newArrayAttribute("signature", ByteSize),
	newArrayAttribute("signer", ByteSize),
	newScalarAttribute("version", ShortSize),
	newScalarAttribute("type", ShortSize),
	newArrayAttribute("fee", IntSize),
	newArrayAttribute("deadline", IntSize),
}

var aggregateTransactionSchema = &schema{
	append(
		abstractTransactionSchemaAttributes,
		[]schemaAttribute{
			newScalarAttribute("transactionsSize", IntSize),
			newArrayAttribute("transactions", ByteSize),
		}...,
	),
}

var mosaicDefinitionTransactionSchema = &schema{
	append(
		abstractTransactionSchemaAttributes,
		[]schemaAttribute{
			newArrayAttribute("parentId", IntSize),
			newArrayAttribute("mosaicId", IntSize),
			newScalarAttribute("mosaicNameLength", ByteSize),
			newScalarAttribute("numOptionalProperties", ByteSize),
			newScalarAttribute("flags", ByteSize),
			newScalarAttribute("divisibility", ByteSize),
			newArrayAttribute("mosaicName", ByteSize),
			newScalarAttribute("indicateDuration", ByteSize),
			newArrayAttribute("duration", IntSize),
		}...,
	),
}

var mosaicSupplyChangeTransactionSchema = &schema{
	append(
		abstractTransactionSchemaAttributes,
		[]schemaAttribute{
			newArrayAttribute("mosaicId", IntSize),
			newScalarAttribute("direction", ByteSize),
			newArrayAttribute("delta", IntSize),
		}...,
	),
}

var transferTransactionSchema = &schema{
	append(
		abstractTransactionSchemaAttributes,
		[]schemaAttribute{
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
		}...,
	),
}

var modifyMultisigAccountTransactionSchema = &schema{
	append(
		abstractTransactionSchemaAttributes,
		[]schemaAttribute{
			newScalarAttribute("minRemovalDelta", ByteSize),
			newScalarAttribute("minApprovalDelta", ByteSize),
			newScalarAttribute("numModifications", ByteSize),
			newTableArrayAttribute("modification", schema{
				[]schemaAttribute{
					newScalarAttribute("type", ByteSize),
					newArrayAttribute("cosignatoryPublicKey", ByteSize),
				},
			}.schemaDefinition),
		}...,
	),
}

var registerNamespaceTransactionSchema = &schema{
	append(
		abstractTransactionSchemaAttributes,
		[]schemaAttribute{
			newArrayAttribute("namespaceType", IntSize),
			newScalarAttribute("durationParentId", IntSize),
			newArrayAttribute("namespaceId", IntSize),
			newArrayAttribute("namespaceNameSize", ByteSize),
			newArrayAttribute("name", ByteSize),
		}...,
	),
}

var lockFundsTransactionSchema = &schema{
	append(
		abstractTransactionSchemaAttributes,
		[]schemaAttribute{
			newArrayAttribute("mosaicId", IntSize),
			newArrayAttribute("mosaicAmount", IntSize),
			newArrayAttribute("duration", IntSize),
			newArrayAttribute("hash", IntSize),
		}...,
	),
}

var secretLockTransactionSchema = &schema{
	append(
		abstractTransactionSchemaAttributes,
		[]schemaAttribute{
			newArrayAttribute("mosaicId", IntSize),
			newArrayAttribute("mosaicAmount", IntSize),
			newArrayAttribute("duration", IntSize),
			newScalarAttribute("hashAlgorithm", IntSize),
			newArrayAttribute("secret", IntSize),
			newArrayAttribute("recipient", IntSize),
		}...,
	),
}

var secretProofTransactionSchema = &schema{
	append(
		abstractTransactionSchemaAttributes,
		[]schemaAttribute{
			newScalarAttribute("hashAlgorithm", ByteSize),
			newArrayAttribute("secret", ByteSize),
			newScalarAttribute("proofSize", ShortSize),
			newArrayAttribute("proof", IntSize),
		}...,
	),
}
