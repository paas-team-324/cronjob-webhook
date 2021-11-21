package main

import (
	"context"
	"net/http"

	batchv1 "k8s.io/api/batch/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/validate-batch-v1-cronjob,mutating=false,failurePolicy=fail,groups="batch",resources=cronjobs,verbs=create;update,versions=v1,name=vcronjob.kb.io,admissionReviewVersions=v1beta1,sideEffects=None

// implement admission handler
type cronJobValidationHandler struct {
	Client  client.Client
	decoder *admission.Decoder
}

// admission handler itself
func (v *cronJobValidationHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	cronJob := &batchv1.CronJob{}

	// decode cronjob object
	err := v.decoder.Decode(req, cronJob)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	// forbid cronjobs with allowed concurrent jobs
	if cronJob.Spec.ConcurrencyPolicy == "Allow" {
		return admission.Denied("CronJobs with concurrency policy set to 'Allow' are forbidden")
	}

	return admission.Allowed("")
}

func (v *cronJobValidationHandler) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}
