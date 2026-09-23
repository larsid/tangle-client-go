package info

import (
	"context"
	"errors"
	"log"

	"github.com/larsid/tangle-client-go/client"
)

type AllNodeInfo struct {
	Name        string            `json:"name"`
	NetworkId   string            `json:"network_id"`
	Version     string            `json:"version"`
	IsHealthy   bool              `json:"is_healthy"`
	MinPowScore float64           `json:"min_pow_score"`
	Features    []string          `json:"features"`
	Bech32Hrp   string            `json:"bech_32_hrp"`
	Milestone   NodeInfoMilestone `json:"milestone"`
	Messages    NodeInfoMessages  `json:"messages"`
}

type NodeInfo struct {
	Name              string            `json:"name"`
	NetworkId         string            `json:"network_id"`
	Version           string            `json:"version"`
	MessagesPerSecond float64           `json:"messages_per_second"`
	Milestone         NodeInfoMilestone `json:"milestone"`
	IsHealthy         bool              `json:"is_healthy"`
}

type NodeInfoMessages struct {
	MessagesPerSecond           float64 `json:"messages_per_second"`
	ReferencedMessagesPerSecond float64 `json:"referenced_messages_per_second"`
	ReferencedRate              float64 `json:"referenced_rate"`
}

type NodeInfoMilestone struct {
	ConfirmedMilestoneIndex  uint32 `json:"confirmed_milestone_index"`
	LatestMilestoneIndex     uint32 `json:"latest_milestone_index"`
	LatestMilestoneTimestamp int64  `json:"latest_milestone_timestamp"`
	PruningIndex             uint32 `json:"pruning_index"`
}

// Get Tangle Hornet Network node information.
func GetNodeInfo(ctx context.Context, nodeUrl string) (NodeInfo, error) {
	node := client.ForURL(nodeUrl)

	log.Printf("[TANGLE-CLIENT] [INFO] GetNodeInfo: consultando nó em %s", nodeUrl)

	info, err := node.Info(ctx)
	if err != nil {
		log.Printf("[TANGLE-CLIENT] [ERROR] GetNodeInfo: falha ao obter info do nó %s: %v", nodeUrl, err)
		return NodeInfo{}, errors.New("unable to get node information")
	}

	log.Printf(
		"[TANGLE-CLIENT] [INFO] GetNodeInfo: nó=%s healthy=%v pruning_index=%d latest_milestone=%d",
		info.Name,
		info.IsHealthy,
		info.PruningIndex,
		info.LatestMilestoneIndex,
	)

	milestone := &NodeInfoMilestone{
		ConfirmedMilestoneIndex:  info.ConfirmedMilestoneIndex,
		LatestMilestoneIndex:     info.LatestMilestoneIndex,
		LatestMilestoneTimestamp: info.LatestMilestoneTimestamp,
		PruningIndex:             info.PruningIndex,
	}

	nodeInfo := &NodeInfo{
		Name:              info.Name,
		NetworkId:         info.NetworkID,
		Version:           info.Version,
		MessagesPerSecond: info.MessagesPerSecond,
		Milestone:         *milestone,
		IsHealthy:         info.IsHealthy,
	}

	return *nodeInfo, nil
}

// Get all Tangle Hornet Network node information.
func GetAllNodeInfo(ctx context.Context, nodeUrl string) (AllNodeInfo, error) {
	node := client.ForURL(nodeUrl)

	log.Printf("[TANGLE-CLIENT] [INFO] GetAllNodeInfo: consultando nó em %s", nodeUrl)

	info, err := node.Info(ctx)
	if err != nil {
		log.Printf("[TANGLE-CLIENT] [ERROR] GetAllNodeInfo: falha ao obter info do nó %s: %v", nodeUrl, err)
		return AllNodeInfo{}, errors.New("unable to get node information")
	}

	log.Printf(
		"[TANGLE-CLIENT] [INFO] GetAllNodeInfo: nó=%s healthy=%v pruning_index=%d latest_milestone=%d",
		info.Name,
		info.IsHealthy,
		info.PruningIndex,
		info.LatestMilestoneIndex,
	)

	milestone := &NodeInfoMilestone{
		ConfirmedMilestoneIndex:  info.ConfirmedMilestoneIndex,
		LatestMilestoneIndex:     info.LatestMilestoneIndex,
		LatestMilestoneTimestamp: info.LatestMilestoneTimestamp,
		PruningIndex:             info.PruningIndex,
	}

	message := &NodeInfoMessages{
		MessagesPerSecond:           info.MessagesPerSecond,
		ReferencedMessagesPerSecond: info.ReferencedMessagesPerSecond,
		ReferencedRate:              info.ReferencedRate,
	}

	allNodeInfo := &AllNodeInfo{
		Name:        info.Name,
		NetworkId:   info.NetworkID,
		Version:     info.Version,
		IsHealthy:   info.IsHealthy,
		MinPowScore: info.MinPowScore,
		Features:    info.Features,
		Bech32Hrp:   info.Bech32HRP,
		Milestone:   *milestone,
		Messages:    *message,
	}

	return *allNodeInfo, nil
}
