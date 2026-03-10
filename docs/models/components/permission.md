# Permission

The permission level to grant. Read more about roles and their permissions on [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).

## Example Usage

```go
import (
	"github.com/polarsource/polar-go/models/components"
)

value := components.PermissionPull
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `PermissionPull`     | pull                 |
| `PermissionTriage`   | triage               |
| `PermissionPush`     | push                 |
| `PermissionMaintain` | maintain             |
| `PermissionAdmin`    | admin                |