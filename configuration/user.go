package configuration

type User struct {
   Id string `json:"id,omitempty"`
   Username string `json:"username,omitempty"`
   Environments []Environment `json:"environments"`
   Roles []Role `json:"roles"`
   Organization Organization `json:"tenant,omitempty"`
}