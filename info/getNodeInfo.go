package info

import (
	"context"
	"log"

	iotago "github.com/iotaledger/iota.go/v2"
)

type NodeInfo struct {
	Name              string    `json:"name"`
	NetworkId         string    `json:"network_id"`
	Version           string    `json:"version"`
	MessagesPerSecond float64   `json:"messages_per_second"`
	Milestone         Milestone `json:"milestone"`
	IsHealthy         bool      `json:"is_healthy"`
}

type Milestone struct {
	ConfirmedMilestoneIndex  uint32 `json:"confirmed_milestone_index"`
	LatestMilestoneIndex     uint32 `json:"latest_milestone_index"`
	LatestMilestoneTimestamp int64  `json:"latest_milestone_timestamp"`
}

// Get Tangle Hornet Network node information.
func GetNodeInfo(nodeUrl string) NodeInfo {
	node := iotago.NewNodeHTTPAPIClient(nodeUrl)

	info, err := node.Info(context.Background())
	if err != nil {
		log.Fatal("Unable to get node information.")
	}

	milestone := &Milestone{
		ConfirmedMilestoneIndex:  info.ConfirmedMilestoneIndex,
		LatestMilestoneIndex:     info.LatestMilestoneIndex,
		LatestMilestoneTimestamp: info.LatestMilestoneTimestamp,
	}

	return NodeInfo{
		Name:              info.Name,
		NetworkId:         info.NetworkID,
		Version:           info.Version,
		MessagesPerSecond: info.MessagesPerSecond,
		Milestone:         *milestone,
		IsHealthy:         info.IsHealthy,
	}
}
