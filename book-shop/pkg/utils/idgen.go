package utils

import (
	"github.com/bwmarrin/snowflake"
)

func GenerateID() (int64, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}

	id := node.Generate().Int64()
	return id, nil
}
