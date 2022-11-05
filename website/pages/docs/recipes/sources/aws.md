# AWS Source Plugin Recipes

Full spec options for AWS Source available [here](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/docs/configuration.md).

```yaml
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "v4.2.0" # latest version of aws plugin
  tables: ["*"]
  destinations: ["YOUR_DESTINATION"]
```