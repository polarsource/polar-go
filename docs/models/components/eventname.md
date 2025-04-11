# EventName


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `Name`                                    | *string*                                  | :heavy_check_mark:                        | The name of the event.                    |
| `Occurrences`                             | *int64*                                   | :heavy_check_mark:                        | Number of times the event has occurred.   |
| `FirstSeen`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | The first time the event occurred.        |
| `LastSeen`                                | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | The last time the event occurred.         |