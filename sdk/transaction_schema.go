package sdk

func transactionSchema(schemaAttributes []schemaAttributeSuper) {
	schemaAttributes = append(schemaAttributes, newScalarAttribute("size", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("signature", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("signer", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("version", SIZEOF_SHORT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("type", SIZEOF_SHORT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("fee", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("deadline", SIZEOF_INT))
}

func aggregateTransactionSchema() *schema {
	schemaAttributes := new(schema).schemaDefinition
	transactionSchema(schemaAttributes)
	schemaAttributes = append(schemaAttributes, newScalarAttribute("transactionsSize", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("transactions", SIZEOF_BYTE))

	schema := newSchema(schemaAttributes)
	return &schema
}

func mosaicDefinitionTransactionSchema() *schema {
	schemaAttributes := new(schema).schemaDefinition
	transactionSchema(schemaAttributes)
	schemaAttributes = append(schemaAttributes, newArrayAttribute("parentId", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("mosaicId", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("mosaicNameLength", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("numOptionalProperties", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("flags", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("divisibility", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("mosaicName", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("indicateDuration", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("duration", SIZEOF_INT))

	schema := newSchema(schemaAttributes)
	return &schema
}

func mosaicSupplyChangeTransaction() *schema {
	schemaAttributes := new(schema).schemaDefinition
	transactionSchema(schemaAttributes)
	schemaAttributes = append(schemaAttributes, newArrayAttribute("mosaicId", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("direction", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("delta", SIZEOF_INT))

	schema := newSchema(schemaAttributes)
	return &schema
}

func transferTransactionSchema() *schema {
	schemaAttributes := new(schema).schemaDefinition
	transactionSchema(schemaAttributes)
	schemaAttributes = append(schemaAttributes, newArrayAttribute("recipient", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("messageSize", SIZEOF_SHORT))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("numMosaics", SIZEOF_BYTE))

	messageAttributes := new(schema).schemaDefinition
	messageAttributes = append(messageAttributes, newScalarAttribute("type", SIZEOF_BYTE))
	messageAttributes = append(messageAttributes, newArrayAttribute("payload", SIZEOF_BYTE))

	schemaAttributes = append(schemaAttributes, newTableAttribute("message", messageAttributes))

	mosaicAttributes := new(schema).schemaDefinition
	mosaicAttributes = append(mosaicAttributes, newArrayAttribute("id", SIZEOF_INT))
	mosaicAttributes = append(mosaicAttributes, newArrayAttribute("amount", SIZEOF_INT))

	schemaAttributes = append(schemaAttributes, newTableAttribute("mosaics", mosaicAttributes))

	schema := newSchema(schemaAttributes)
	return &schema
}

func modifyMultisigAccountTransactionSchema() *schema {
	schemaAttributes := new(schema).schemaDefinition
	transactionSchema(schemaAttributes)
	schemaAttributes = append(schemaAttributes, newScalarAttribute("minRemovalDelta", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("minApprovalDelta", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("numModifications", SIZEOF_BYTE))

	messageAttributes := new(schema).schemaDefinition
	messageAttributes = append(messageAttributes, newScalarAttribute("type", SIZEOF_BYTE))
	messageAttributes = append(messageAttributes, newArrayAttribute("cosignatoryPublicKey", SIZEOF_BYTE))

	schema := newSchema(schemaAttributes)
	return &schema
}

func registerNamespaceTransactionSchema() *schema {
	schemaAttributes := new(schema).schemaDefinition
	transactionSchema(schemaAttributes)
	schemaAttributes = append(schemaAttributes, newArrayAttribute("namespaceType", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("durationParentId", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("namespaceId", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("namespaceNameSize", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("name", SIZEOF_BYTE))

	schema := newSchema(schemaAttributes)
	return &schema
}

func lockFundsTransactionSchema() *schema {
	schemaAttributes := new(schema).schemaDefinition
	transactionSchema(schemaAttributes)
	schemaAttributes = append(schemaAttributes, newArrayAttribute("mosaicId", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("mosaicAmount", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("duration", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("hash", SIZEOF_INT))

	schema := newSchema(schemaAttributes)
	return &schema
}

func secretLockTransactionSchema() *schema {
	schemaAttributes := new(schema).schemaDefinition
	transactionSchema(schemaAttributes)
	schemaAttributes = append(schemaAttributes, newArrayAttribute("mosaicId", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("mosaicAmount", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("duration", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("hashAlgorithm", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("secret", SIZEOF_INT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("recipient", SIZEOF_INT))

	schema := newSchema(schemaAttributes)
	return &schema
}

func secretProofTransactionSchema() *schema {
	schemaAttributes := new(schema).schemaDefinition
	transactionSchema(schemaAttributes)
	schemaAttributes = append(schemaAttributes, newScalarAttribute("secret", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("recipient", SIZEOF_BYTE))
	schemaAttributes = append(schemaAttributes, newScalarAttribute("secret", SIZEOF_SHORT))
	schemaAttributes = append(schemaAttributes, newArrayAttribute("recipient", SIZEOF_INT))

	schema := newSchema(schemaAttributes)
	return &schema
}
