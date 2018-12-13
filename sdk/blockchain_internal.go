package sdk

func blockInfoDTOsToBlockInfos(dtos []*blockInfoDTO) ([]*BlockInfo, error) {
	blocks := make([]*BlockInfo, 0, len(dtos))

	for _, dto := range dtos {
		block, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block)
	}

	return blocks, nil
}
