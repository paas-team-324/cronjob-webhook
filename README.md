# CronJob Webhook

Custom CronJob webhooks. Scaffoled with kubebuilder based on [the following](https://book.kubebuilder.io/reference/webhook-for-core-types.html) guide.

# Example

```bash
make deploy IMG=docker.io/paasteam324/cronjob-webhook:0.1.0
oc create -f examples/
```