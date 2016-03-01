package configuration

type Organization struct {
   Id string `json:"id,omitempty"`
   Name string `json:"name,omitempty"`
   EntryPoint string `json:"entryPoint,omitempty"`
   Users []User `json:"users"`
   Environments []Environment `json:"environments"`
   Roles []Role `json:"roles"`
}