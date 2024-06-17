package net

import (
	certmanager "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	log "github.com/sirupsen/logrus"
	istio "istio.io/client-go/pkg/apis/networking/v1beta1"
	"sigs.k8s.io/yaml"

	"ultary.co/kluster/pkg/helm"
	"ultary.co/kluster/pkg/k8s"
)

type Gateway struct {
	certificate certmanager.Certificate
	gateway     istio.Gateway
}

func NewGateway(chart *helm.Chart) (retval Gateway) {

	const name = "monokube"

	m := chart.Get("Certificate", name)
	if err := yaml.Unmarshal(m, &retval.certificate); err != nil {
		log.Fatalf("Error unmarshalling YAML to Certificate: %v", err)
	}

	m = chart.Get("Gateway", name)
	if err := yaml.Unmarshal(m, &retval.gateway); err != nil {
		log.Fatalf("Error unmarshalling YAML to Gateway: %v", err)
	}

	return
}

func (g *Gateway) Apply(ctx k8s.Context, namespace string) error {

	if err := k8s.ApplyCertificate(ctx, &g.certificate, namespace); err != nil {
		log.Fatalf("error applying certificate: %v", err)
	}

	if err := k8s.ApplyGateway(ctx, &g.gateway, namespace); err != nil {
		log.Fatalf("error applying gateway: %v", err)
	}

	return nil
}