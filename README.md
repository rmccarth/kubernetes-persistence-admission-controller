### Malicious Webhook Admission Controller

This is a 'malicious' webhook admission controller that I developed to learn how to write admission controllers in Go. 

The controller allows an attacker to maintain persistence in a Kubernetes cluster through manifest re-writes when they are submitted to the cluster.

The controller will add 'slixperi says this could have been a total yaml re-write' as a label to any pod being created (with the label field in the original pod.yaml) in the webhookdemo namespace (sorry skripties). 

Thanks to `https://pet2cattle.com/2021/08/kubernetes-mutating-webhook` for the wonderful demonstration in the fundamentals and template files.  


Future work: I'll prolly make this into something sneakier once I learn a bit more about how Kubernetes works under the hood. For now its discoverable through any sort of AdmissionController monitoring or `kubectl describe pod [pod name]`.