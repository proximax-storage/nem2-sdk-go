package sdk

func transactionStatusDTOsToTransactionStatuses(dtos []*transactionStatusDTO) ([]*TransactionStatus, error) {
	statuses := make([]*TransactionStatus, 0, len(dtos))

	for _, dto := range dtos {
		status, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		statuses = append(statuses, status)
	}

	return statuses, nil
}
