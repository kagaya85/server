// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameOAuthClient = "chii_oauth_clients"

// OAuthClient mapped from table <chii_oauth_clients>
type OAuthClient struct {
	AppID        uint32 `gorm:"column:app_id;type:mediumint(8);primaryKey"`
	ClientID     string `gorm:"column:client_id;type:varchar(80);not null;index:client_id,priority:1"`
	ClientSecret string `gorm:"column:client_secret;type:varchar(80)"`
	RedirectURI  string `gorm:"column:redirect_uri;type:varchar(2000)"`
	GrantTypes   string `gorm:"column:grant_types;type:varchar(80)"`
	Scope        string `gorm:"column:scope;type:varchar(4000)"`
	UserID       string `gorm:"column:user_id;type:varchar(80)"`
	App          App    `gorm:"foreignKey:app_id;references:app_id" json:"app"`
}

// TableName OAuthClient's table name
func (*OAuthClient) TableName() string {
	return TableNameOAuthClient
}
