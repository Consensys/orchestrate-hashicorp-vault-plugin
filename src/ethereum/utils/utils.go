package utils

import "fmt"

func ComputeKey(address, namespace string) string {
	if namespace == "" {
		return address
	}

	return fmt.Sprintf("%s%s", namespace, address)
}
