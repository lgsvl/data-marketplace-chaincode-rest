# Chaincode REST 
This repository contains the needed code that exposes the chaincode (Hyperledger Fabric smart contract). The project is written in [Go](https://golang.org/).
To run this component correctly, you should be familiar with the [Data marketplace](https://github.com/lgsvl/data-marketplace) components because there is a particular dependency between the components.

# Build prerequisites
  * Install [golang](https://golang.org/).
  * Install [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).
  * Configure go. GOPATH environment variable must be set correctly before starting the build process.

### Download and build source code

```bash
mkdir -p $HOME/workspace
export GOPATH=$HOME/workspace
mkdir -p $GOPATH/src/github.com/lgsvl
cd $GOPATH/src/github.com/lgsvl
git clone git@github.com:lgsvl/data-marketplace-chaincode-rest.git
cd data-marketplace-chaincode-rest
./scripts/build
```

### Docker Build

You can use the dockerfile to build a docker image:
```
docker build -t chaincode-rest .
docker run -p 9090:9090 chaincode-rest
```

You can see the API spec on [localhost:9090/swaggerui](localhost:9090/swaggerui).

### Kubernetes deployment

The [deployment](./deployment) folder contains the deployment and service/ingress manifests to deploy this component.
We assume that you have a running Hyperledger Fabric network running with the [Data marketplace Chaincode](https://github.com/lgsvl/data-marketplace-chaincode) deployed on top of it.
We assume also that the deployment has access to a shared volume (PVC) named `sharedvolume` that contains information about the chaincode, used channels and certificates (probably created for hyperledger Fabric setup already).

Once deployed, the service API could be accessed at the `<service url>:9090/swaggerui`.

# Running the Unit Tests

Run the tests:
```bash
./scripts/run_units.sh
```