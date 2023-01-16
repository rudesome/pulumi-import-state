# pulumi-import-state

- GET private github repositories - by the PAT (Personal Access Token)
    - filter forks
- (WIP) Import github repos to pulumi stack (based on name)


## instruction:

1. rename sample.env to .env and set your own PAT github api token:

```bash
API_KEY="XXXXXXXXXXXXXXXXX"
```

2. don't forget to set the pulumi code to keep the state as code configured


## TODO:
- Rework to module / pacakage
- Let the user define the pulumi stack as argument..
- Add more providers (Azure / AWS / Cloudflare / Kubernetes)

