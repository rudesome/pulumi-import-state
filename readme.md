# pulumi-import-state

- GET private github repositories - by the PAT (Personal Access Token)
    - filter forks
- (WIP) Import github repos to pulumi stack (based on name)

---

## instruction:

- create new pulumi project / stack
- rename sample.env to .env and set your github PAT:

```bash
API_KEY=XXXXXXXXXXXXXXXXX
```

---

## TODO:
[x] Rework to module / package
[ ] Let the user define the pulumi stack (folder?) as argument..
[ ] Add more providers (Azure / AWS / Cloudflare / Kubernetes)
