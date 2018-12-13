package sdk

func accountInfoDTOsToAccountInfos(accountDTOs []*accountInfoDTO) ([]*AccountInfo, error) {
	accountInfos := make([]*AccountInfo, 0, len(accountDTOs))

	for _, dto := range accountDTOs {
		accountInfo, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		accountInfos = append(accountInfos, accountInfo)
	}

	return accountInfos, nil
}
