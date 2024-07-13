package kube

import (
	"github.com/ultary/monokube/kluster/pkg/k8s"
	"github.com/ultary/monokube/kluster/pkg/kube/apps/system/otlp"
)

type system struct {
	client *k8s.Client
}

func NewSystem(client *k8s.Client) *system {
	return &system{
		client: client,
	}
}

func (s *system) Initialize() {
}

func (s *system) EnableOpenTelemetryCollector() {
	otlp.Enable(s.client)
}

func (s *system) DisableOpenTelemetryCollector() {
	otlp.Disable(s.client)
}

func (s *system) UpdateOpenTelemetryCollector() {
	otlp.Update(s.client)
}