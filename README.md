# Kubernauts TK8 EKS CLI

This repo provides terraform modules based on the grate implementation from Wesley Charles Blake:

https://github.com/WesleyCharlesBlake

and the golang sources for tk8eks cli.

## Deploy an EKS Cluster

1. Clone this repo

2. Set the desired region and number of worker nodes in variables.tf file

3. Set the path to your public key in eks-worker-nodes.tf under modules/eks/

4. Download terraform for darwin or linux and put in the root of the cloned repo

5. Install aws-iam-authenticator (heptio-authenticator-aws)

6. Download tk8eks cli from releases  

7. export AWS_SECRET_ACCESS_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

8. export AWS_ACCESS_KEY_ID=xxxxxxxxxxxxxxxxx

9. run "./tk8eks-darwin-amd64 create" or "./tk8eks-linux-amd64 create" to deploy the cluster

10. run "./tk8eks-darwin-amd64 delete" or "./tk8eks-linux-amd64 delete" to delete the cluster

## Build from source

```bash
$ go build .
```

### Get in touch

Join us on Kubernauts Slack Channel:

https://kubernauts-slack-join.herokuapp.com/
