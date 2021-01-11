package utils

import "fmt"

func ComputeEthereumKey(address, namespace string) string {
	path := fmt.Sprintf("ethereum/accounts/%s", address)
	if namespace != "" {
		path = fmt.Sprintf("%s/%s", namespace, path)
	}

	return path
}

func ComputeZksKey(address, namespace string) string {
	path := fmt.Sprintf("zk-snarks/accounts/%s", address)
	if namespace != "" {
		path = fmt.Sprintf("%s/%s", namespace, path)
	}

	return path
}
