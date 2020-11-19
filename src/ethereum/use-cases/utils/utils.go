package utils

import "fmt"

func ComputeKey(address, namespace string) string {
	path := fmt.Sprintf("%s/ethereum/accounts/%s", namespace, address)
	if namespace != "" {
		path = fmt.Sprintf("%s/%s", namespace, path)
	}

	return path
}
