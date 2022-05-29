package db

type MockDatabase struct{}

const MockUUID = "b7dbb635-5881-4351-842e-4eb481fbd1ae"

func (database MockDatabase) QueryUserInformation(query QueryType, uuid string) (bool, error) {
	if uuid == MockUUID {
		return true, nil
	}
	return false, nil
}
