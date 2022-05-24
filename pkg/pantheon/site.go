package pantheon

import (
	"encoding/json"
	"fmt"
	"log"
)

// Site is a representation of a deployed pantheon site.
type Site struct {
	Archived bool   `json:"archived"`
	ID       string `json:"id"`
	Key      string `json:"key"`
	Role     string `json:"role"`
	Site     struct {
		Created         jsonInt64 `json:"created"`
		CreatedByUserID string    `json:"created_by_user_id"`
		Framework       string    `json:"framework"`
		Frozen          bool      `json:"frozen"`
		ID              string    `json:"id"`
		LastCodePush    struct {
			Timestamp string      `json:"timestamp"`
			UserUUID  interface{} `json:"user_uuid"`
		} `json:"last_code_push"`
		Name          string    `json:"name"`
		Owner         string    `json:"owner"`
		PhpVersion    jsonInt64 `json:"php_version"`
		PreferredZone string    `json:"preferred_zone"`
		Product       struct {
			ID       string `json:"id"`
			Longname string `json:"longname"`
		} `json:"product"`
		ProductID    string `json:"product_id"`
		ServiceLevel string `json:"service_level"`
		Upstream     struct {
			Branch    string `json:"branch"`
			ProductID string `json:"product_id"`
			URL       string `json:"url"`
		} `json:"upstream"`
	} `json:"site"`
	SiteID string `json:"site_id"`
	UserID string `json:"user_id"`
}

// SiteList represents a grouping of deployed Pantheon sites.
type SiteList struct {
	Sites []Site
}

// NewSiteList creates an SiteList. The user will be specified by which session you use to make the GET request. You are responsible for making the HTTP request.
func NewSiteList() *SiteList {
	return &SiteList{}
}

// Path returns the API endpoint which can be used to get a SiteList for the current user.
func (sl SiteList) Path(method string, auth AuthSession) string {
	userid, err := auth.GetUser()
	if err != nil {
		log.Fatalf("Could not determine user for request: %v", err)
	}

	return fmt.Sprintf("/users/%s/memberships/sites", userid)
}

// JSON prepares the SiteList for HTTP transport.
func (sl SiteList) JSON() ([]byte, error) {
	return json.Marshal(sl.Sites)
}

// Unmarshal is responsible for converting a HTTP response into a SiteList struct.
func (sl *SiteList) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &sl.Sites)
}

// Domain is a representation of a domain for a deployed Pantheon site.
type Domain struct {
	DNSZoneName      string `json:"dns_zone_name"`
	DomainLookupRAX  uint64 `json:"domain_lookup_rax"`
	DomainLookupUSCL uint64 `json:"domain_lookup_usc1"`
	Environment      string `json:"environment"`
	SiteID           string `json:"site_id"`
	Type             string `json:"type"`
	ID               string `json:"id"`
	Key              string `json:"key"`
	Status           string `json:"status"`
	StatusMessage    string `json:"status_message"`
	Deleteable       bool   `json:"deletable"`
}

// SiteDomainList represents a grouping of domains for a deployed Pantheon site.
type SiteDomainList struct {
	SiteID      string
	Environment string
	Domains     []Domain
}

// NewSiteList creates an SiteDomainList. The user will be specified by which session you use to make the GET request. You are responsible for making the HTTP request.
func NewSiteDomainList(siteID string, env string) *SiteDomainList {
	return &SiteDomainList{
		SiteID:      siteID,
		Environment: env,
		Domains:     make([]Domain, 0),
	}
}

// Path returns the API endpoint which can be used to get a SiteDomainList for the current user.
func (sdl SiteDomainList) Path(method string, auth AuthSession) string {
	return fmt.Sprintf("/sites/%s/environments/%s/domains", sdl.SiteID, sdl.Environment)
}

// JSON prepares the SiteList for HTTP transport.
func (sdl SiteDomainList) JSON() ([]byte, error) {
	return json.Marshal(sdl.Domains)
}

// Unmarshal is responsible for converting a HTTP response into a SiteDomainList struct.
func (sdl *SiteDomainList) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &sdl.Domains)
}
