package cmd

import (
	"fmt"
	"strings"
)

func splitArguments(argumentList string) (map[string]string, error) {
	arguments := make(map[string]string)
	if "" != argumentList {
		argumentPairs := strings.Split(argumentList, ",")
		for _, pair := range argumentPairs {
			splittedPair := strings.Split(pair, "=")
			if _, ok := arguments[splittedPair[0]]; ok {
				return nil, fmt.Errorf("argument doubled: [%s]", splittedPair[0])
			} else {
				if len(splittedPair) == 1 {
					arguments[splittedPair[0]] = ""
				} else {
					arguments[splittedPair[0]] = splittedPair[1]
				}
			}
		}
	}
	return arguments, nil
}
