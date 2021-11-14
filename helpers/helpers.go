package helpers

func CompleteAddressTo16(binaryAddress string) string {
	for i := len(binaryAddress); i < 16; i++ {
		binaryAddress = "0" + binaryAddress
	}

	return binaryAddress
}
