# CronJob Webhook

Custom CronJob webhooks. Scaffoled with kubebuilder based on [the following](https://book.kubebuilder.io/reference/webhook-for-core-types.html) guide.

# Example

```bash
NAMESPACE=<cronjob-webhook-namespace>

oc new-project $NAMESPACE
make bundle IMG=docker.io/paasteam324/cronjob-webhook:0.1.0 NAMESPACE=$NAMESPACE

oc create -f deploy/bundle.yaml
oc create -f examples/
```