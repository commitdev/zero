![zero](https://raw.githubusercontent.com/commitdev/zero/main/docs/img/logo-small.png)

## Real-world Usage Scenarios

### Developing and deploying application changes
1. Clone your git repository.
2. Make a branch, start working on your code.
3. If using the Telepresence dev experience, run the `start-dev-env.sh` script to allow you to use the hybrid cloud environment as you work, to run and test your code in a realistic environment.
3. Commit your finished code, make a PR, have it reviewed. Lightweight tests will run against your branch and prevent merging if they fail.
4. Merge your branch to the main branch. A build will start automatically.
5. The pipeline will build an artifact, run tests, deploy your change to staging, then wait for your input to deploy to production.

### Debugging a problem on production
1. Check the logs of your service:
    - If using cloudwatch, log into the AWS console and go to the [Logs Insights tool](https://us-west-2.console.aws.amazon.com/cloudwatch/home#logsV2:logs-insights). Choose the log group for your production environment ending in `/application` and hit the "Run query" button.
    - If using kibana, make sure you are on the VPN and open the Kibana URL in your browser. Click the "Discover" tab and try searching for logs based on the name of your service.
    - Alternatively, watch the logs in realtime via the CLI using the command `kubectl logs -f -l app=<your service name>` or `stern <your service name>`
2. Check the state of your application pods. Look for strange events or errors from the pods:
```shell
$ kubectl get pods
$ kubectl get events
$ kubectl describe pods
```
3. Exec into your application pod. From here you can check connectivity with `ping` or `nc`, or inspect anything else application-specific.
```shell
$ kubectl get pods
NAME                             READY   STATUS    RESTARTS   AGE
your-service-6c5f6b56b7-2w447    1/1     Running   0          30m49s
$ kubectl exec -it your-service-6c5f6b56b7-2w447 sh
```


### Adding support for a new subdomain or site
1. Check the currently configured ingresses in your cluster:
```shell
$ kubectl get ingress -A
NAMESPACE      NAME               CLASS    HOSTS                   ADDRESS                                                                   PORTS     AGE
your-service   your-service       <none>   api.your-service.dev         abcd1234-1234.us-west-2.elb.amazonaws.com   80, 443   130d
```
2. If this is for a new service entirely, make sure there is an ingress defined in the `kubernetes/` directory of your repo. If you want to add a new domain pointing to an existing service, just go into the file `kubernetes/overlays/<environment>/ingress.yml` and add a section to `spec:` and `tls:`, specifying your new domain.
    - `spec` is where you can define the hostname, any special path rules, and which service you want traffic to be sent to
    - if your hostname is in the `tls` section, a TLS certificate will automatically be provisioned for it using Let's Encrypt
3. A number of things will happen once this is deployed to the cluster:
    - Routing rules will be created to let traffic in to the cluster and send it to the service based on the hostname and path
    - An AWS load balancer will be created if one doesn't already exist and it will be pointed to the cluster
    - In the case of a subdomain, a DNS record will be automatically created for you
    - A certificate will be provisioned using Let's Encrypt for the domain you specified
