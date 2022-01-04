# CronJob Webhook

Custom CronJob webhooks. Scaffolded with kubebuilder using [the following](https://book.kubebuilder.io/reference/webhook-for-core-types.html) guide.

## Development

Edit the code, then do the following in order to build the image:

```bash
IMAGE=docker.io/paasteam324/cronjob-webhook:<version>

docker login -u paasteam324 docker.io

make docker-build IMG=$IMAGE
make docker-push IMG=$IMAGE
```

## Deployment (OpenShift 4.x)

In OpenShift 4, webhook service utilizes service CA. Once the image is in the registry, do the following in order to deploy:

```bash
IMAGE=docker.io/paasteam324/cronjob-webhook:<version>
NAMESPACE=<cronjob-webhook-namespace>

oc new-project $NAMESPACE
make bundle IMG=$IMAGE NAMESPACE=$NAMESPACE

oc create -f deploy/bundle.yaml
oc create -f examples/
```

## Deployment (OpenShift 3.x)

There is no service CA functionality in OpenShift 3, so the certs must be generated and substituted manually.

```bash
IMAGE=docker.io/paasteam324/cronjob-webhook:<version>
NAMESPACE=<cronjob-webhook-namespace>

# comment out OpenShift 4.x resources and uncomment OpenShift 3.x resources
vi config/webhook/kustomization.yaml

oc new-project $NAMESPACE
make bundle IMG=$IMAGE NAMESPACE=$NAMESPACE

# generate long lasting certificate
make certs NAMESPACE=$NAMESPACE
oc create secret tls webhook-server-cert --cert=deploy/certs/tls.crt --key=deploy/certs/tls.key -n $NAMESPACE

# substitute CA in bundle and create
CA_CERT_B64=$(cat deploy/certs/ca.pem.b64) envsubst < deploy/bundle.yaml | oc create -f -
oc create -f examples/
```
