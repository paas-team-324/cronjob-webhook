package webhooks

import (
	"context"

	"gomodules.xyz/jsonpatch/v2"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// implement admission handler
type CronJobMutationHandler struct {
	Client  client.Client
	decoder *admission.Decoder
}

// admission handler for batch-(v2alpha1|v1beta1|v1)-cronjob (k8s 1.10+, OpenShift 3.10+)
func (v *CronJobMutationHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	return admission.Patched("", jsonpatch.NewOperation("replace", "/spec/concurrencyPolicy", "Forbid"))
}

func (v *CronJobMutationHandler) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}
