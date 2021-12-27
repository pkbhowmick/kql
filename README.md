# kql

A query server on Kubernetes resources.

### Example curl command:

```bash
# for query single pod resource
curl -g 'http://localhost:8085/kql?query={pod(name:"nginx-ahgt-86888W",namespace:"demo"){node,phase}}'

# for query podList
curl -g 'http://localhost:8085/kql?query={pods{name,namespace,phase}}'
```
