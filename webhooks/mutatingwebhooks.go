package webhooks

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"

	batchv2alpha1 "k8s.io/api-v0.20.13/batch/v2alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// implement admission handler
type CronJobMutationHandler struct {
	Client  client.Client
	decoder *admission.Decoder
}

type CronJobMutationHandler_v2alpha1 struct {
	CronJobMutationHandler
}

type CronJobMutationHandler_v1beta1 struct {
	CronJobMutationHandler
}

type CronJobMutationHandler_v1 struct {
	CronJobMutationHandler
}

// admission handler for batch-v2alpha1-cronjob (k8s 1.10, OpenShift 3.10)
func (v *CronJobMutationHandler_v2alpha1) Handle(ctx context.Context, req admission.Request) admission.Response {
	cronJob := &batchv2alpha1.CronJob{}

	// decode cronjob object
	err := v.decoder.Decode(req, cronJob)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	// mutate cronjob
	marshaledCronJob, err := mutate_cronjob(cronJob)
	if err != nil {
		admission.Errored(http.StatusInternalServerError, err)
	}

	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledCronJob)
}

// admission handler for batch-v1beta1-cronjob (k8s 1.11, OpenShift 3.11)
func (v *CronJobMutationHandler_v1beta1) Handle(ctx context.Context, req admission.Request) admission.Response {
	cronJob := &batchv1beta1.CronJob{}

	// decode cronjob object
	err := v.decoder.Decode(req, cronJob)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	// mutate cronjob
	marshaledCronJob, err := mutate_cronjob(cronJob)
	if err != nil {
		admission.Errored(http.StatusInternalServerError, err)
	}

	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledCronJob)
}

// admission handler for batch-v1-cronjob (k8s 1.22, OpenShift 4.9)
func (v *CronJobMutationHandler_v1) Handle(ctx context.Context, req admission.Request) admission.Response {
	cronJob := &batchv1.CronJob{}

	// decode cronjob object
	err := v.decoder.Decode(req, cronJob)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	// mutate cronjob
	marshaledCronJob, err := mutate_cronjob(cronJob)
	if err != nil {
		admission.Errored(http.StatusInternalServerError, err)
	}

	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledCronJob)
}

func mutate_cronjob(cronJob interface{}) ([]byte, error) {

	// forbid concurrent jobs
	reflect.ValueOf(cronJob).Elem().FieldByName("Spec").FieldByName("ConcurrencyPolicy").SetString("Forbid")

	// unmarshal cronjob as bytes
	marshaledCronJob, err := json.Marshal(cronJob)
	if err != nil {
		return nil, err
	}

	return marshaledCronJob, nil
}

func (v *CronJobMutationHandler) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}
