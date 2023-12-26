# Strongly Consistent Global Cache 

This a (nearly) strongly consistent caching solution that leverages an external Redis service. This approach is designed to maintain strong consistency across multiple cache instances.


## Explanation

If you use the vanilla version of the [cache](../cache) element. Each instance to maintain its own independent table, without any inter-instance synchronization. However, we enhance consistency by integrating an external Redis service as the unified backend storage.


We achieve strong consistency by using an external Redis service as the backend storage.

## Usage
```bash
# Deploy Redis
kubectl apply -f redis.yaml

# Optional Test redis 
curl http://10.96.99.99:7379/PING # {"PING":[true,"PONG"]}

# Optional: insert and get an item
curl http://10.96.99.99:7379/SET/hello/world # {"SET":[true,"OK"]}
curl http://10.96.99.99:7379/GET/hello # {"GET":"world"}

# Apply the redis filter
kubectl apply -f cache-global-strong.yaml

# The first time you execute this, it will not be cached.
# You can check the proxy log, there will be a cache miss log.
curl http://10.96.88.88:8080/ping-echo?body=redis

# The second time you run this, you will get an error (i.e., a cache hit.)
curl http://10.96.88.88:8080/ping-echo?body=redis
```

This element shows how to issue an HTTP call to a Redis service in Envoy. (Note: we are using https://github.com/nicolasff/webdis)

- Use `dispatch_http_call`` to initiate an asynchronous HTTP call to an external service.
- Implement `on_http_call_response`` to handle the incoming response (decode the result and act accordingly).
- Use `self.resume_http_request()` to resume the request/response processing.

Note: the filter configuration is different from other filters. Check `cache-global-strong.yaml`.

Reference: https://github.com/proxy-wasm/proxy-wasm-rust-sdk/tree/master/examples/http_auth_random