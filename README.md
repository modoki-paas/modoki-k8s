# modoki-k8s
<img src="logo/modoki.svg" width="500" />

- PaaS on Kubernetes

## Components
- APIServer
    - Handle requests for App/UserOrg services
- AuthServer
    - Handle external requests and authorize them
    - Reverse proxy to APIServer, etc...
- YAMLer
    - Generate YAML to apply to a Kubernetes cluster
    - Called from APIServer

## License
- Under the MIT License

## Special Thanks
- Logo: [3c1u](https://github.com/3c1u)
