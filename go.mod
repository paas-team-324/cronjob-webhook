module github.com/paas-team-324/cronjob-webhook

go 1.16

replace k8s.io/api-v0.20.13 => k8s.io/api v0.20.13

require (
	k8s.io/api v0.22.1
	k8s.io/api-v0.20.13 v0.0.0-00010101000000-000000000000
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
	sigs.k8s.io/controller-runtime v0.10.0
)
