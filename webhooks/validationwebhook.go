package webhooks

import (
	"context"
	"net/http"

	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"

	batchv2alpha1 "k8s.io/api-v0.20.13/batch/v2alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// implement admission handler
type CronJobValidationHandler struct {
	Client  client.Client
	decoder *admission.Decoder
}

type CronJobValidationHandler_v2alpha1 struct {
	CronJobValidationHandler
}

type CronJobValidationHandler_v1beta1 struct {
	CronJobValidationHandler
}

type CronJobValidationHandler_v1 struct {
	CronJobValidationHandler
}

// admission handler for batch-v2alpha1-cronjob (k8s 1.10, OpenShift 3.10)
func (v *CronJobValidationHandler_v2alpha1) Handle(ctx context.Context, req admission.Request) admission.Response {
	cronJob := &batchv2alpha1.CronJob{}

	// decode cronjob object
	err := v.decoder.Decode(req, cronJob)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	return validate_cronjob(string(cronJob.Spec.ConcurrencyPolicy))
}

// admission handler for batch-v1beta1-cronjob (k8s 1.11, OpenShift 3.11)
func (v *CronJobValidationHandler_v1beta1) Handle(ctx context.Context, req admission.Request) admission.Response {
	cronJob := &batchv1beta1.CronJob{}

	// decode cronjob object
	err := v.decoder.Decode(req, cronJob)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	return validate_cronjob(string(cronJob.Spec.ConcurrencyPolicy))
}

// admission handler for batch-v1-cronjob (k8s 1.22, OpenShift 4.9)
func (v *CronJobValidationHandler_v1) Handle(ctx context.Context, req admission.Request) admission.Response {
	cronJob := &batchv1.CronJob{}

	// decode cronjob object
	err := v.decoder.Decode(req, cronJob)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	return validate_cronjob(string(cronJob.Spec.ConcurrencyPolicy))
}

func validate_cronjob(concurrencyPolicy string) admission.Response {

	// forbid cronjobs with allowed concurrent jobs
	if concurrencyPolicy == "Allow" {
		return admission.Denied("CronJobs with concurrency policy set to 'Allow' are forbidden")
	}

	return admission.Allowed("")
}

func (v *CronJobValidationHandler) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}
