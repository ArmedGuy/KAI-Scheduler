// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package framework

import "github.com/NVIDIA/KAI-scheduler/pkg/scheduler/api"

const (
	Reclaim           api.ActionType = "reclaim"
	Preempt           api.ActionType = "preempt"
	Allocate          api.ActionType = "allocate"
	Consolidation     api.ActionType = "consolidation"
	StaleGangEviction api.ActionType = "stalegangeviction"
)

// Action is the interface of scheduler action.
type Action interface {
	// The unique name of Action.
	Name() api.ActionType

	// Execute allocates the cluster's resources into each queue.
	Execute(ssn *Session)
}

type Plugin interface {
	// The unique name of Plugin.
	Name() string

	OnSessionOpen(ssn *Session)
	OnSessionClose(ssn *Session)
}
