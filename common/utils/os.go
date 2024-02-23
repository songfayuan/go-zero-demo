package utils

import (
	"github.com/denisbrodbeck/machineid"
)

// MachineCode 获取所在电脑机器码
func MachineCode(appID string) (string, error) {
	//id, err := machineid.ID()

	//ID returns the platform specific machine id of the current host OS. Regard the returned id as "confidential" and consider using ProtectedID() instead.
	id, err := machineid.ProtectedID(appID) //推荐
	if err != nil {
		return "", err
	}
	return id, nil
}
