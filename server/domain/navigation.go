package domain

type NavigationSidebarProps struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	YoungestChannelID int64  `json:"youngestChannelId"`
}

type NavigationRepository interface {
	GetSidebarProps() ([]NavigationSidebarProps, error)
}
