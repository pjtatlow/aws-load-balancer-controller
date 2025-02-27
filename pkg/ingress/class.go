package ingress

import (
	elbv2api "github.com/pjtatlow/aws-load-balancer-controller/apis/elbv2/v1beta1"
	networking "k8s.io/api/networking/v1"
)

// ClassifiedIngress is Ingress with it's associated IngressClass Configuration
type ClassifiedIngress struct {
	Ing            *networking.Ingress
	IngClassConfig ClassConfiguration
}

// ClassConfiguration contains configurations for IngressClass
type ClassConfiguration struct {
	// The IngressClass for Ingress if any.
	IngClass *networking.IngressClass

	// The IngressClassParams for Ingress if any.
	IngClassParams *elbv2api.IngressClassParams
}
