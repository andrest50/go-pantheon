package pantheon

import (
	"encoding/json"
	"fmt"
	"log"
)

// Org is a representation of a Pantheon organization.
type Org struct {
	Archived     bool   `json:"archived"`
	ID           string `json:"id"`
	Key          string `json:"key"`
	Role         string `json:"role"`
	Organization struct {
		Profile struct {
			MachineName      string `json:"machine_name"`
			ChangeServiceURL string `json:"change_service_url"`
			Name             string `json:"name"`
			EmailDomain      string `json:"email_domain"`
			OrgLogoWidth     uint32 `json:"org_logo_width"`
			OrgLogoHeight    uint32 `json:"org_logo_height"`
			BaseDomain       string `json:"base_domain"`
			BillingURL       string `json:"billing_url"`
			TermsOfService   string `json:"terms_of_service"`
			OrgLogo          string `json:"org_logo"`
		} `json:"profile"`
		ID string `json:"id"`
	} `json:"organization"`
	SiteID string `json:"site_id"`
	UserID string `json:"user_id"`
}

// OrgList represents a grouping of Pantheon organizations.
type OrgList struct {
	Orgs []Org
}

// NewOrgList creates an OrgList. The user will be specified by which session you use to make the GET request. You are responsible for making the HTTP request.
func NewOrgList() *OrgList {
	return &OrgList{}
}

// Path returns the API endpoint which can be used to get a OrgList for the current user.
func (ol OrgList) Path(method string, auth AuthSession) string {
	userid, err := auth.GetUser()
	if err != nil {
		log.Fatalf("Could not determine user for request: %v", err)
	}

	return fmt.Sprintf("/users/%s/memberships/organizations", userid)
}

// JSON prepares the OrgList for HTTP transport.
func (ol OrgList) JSON() ([]byte, error) {
	return json.Marshal(ol.Orgs)
}

// Unmarshal is responsible for converting a HTTP response into a OrgList struct.
func (ol *OrgList) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &ol.Orgs)
}

// OrgSite is a representation of a deployed pantheon site for an organization.
type OrgSite struct {
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
	Organization string `json:"organization_id"`
	SiteID       string `json:"site_id"`
}

// OrgSiteList represents a grouping of deployed Pantheon sites for an organization.
type OrgSiteList struct {
	Organization string
	OrgSites     []OrgSite
}

// NewOrgSiteList creates an OrgSiteList. The user will be specified by which session you use to make the GET request. You are responsible for making the HTTP request.
func NewOrgSiteList(org string) *OrgSiteList {
	return &OrgSiteList{
		Organization: org,
		OrgSites:     make([]OrgSite, 0),
	}
}

// Path returns the API endpoint which can be used to get a OrgSiteList for the current user.
func (osl OrgSiteList) Path(method string, auth AuthSession) string {
	return fmt.Sprintf("/organizations/%s/memberships/sites", osl.Organization)
}

// JSON prepares the OrgSiteList for HTTP transport.
func (osl OrgSiteList) JSON() ([]byte, error) {
	return json.Marshal(osl.OrgSites)
}

// Unmarshal is responsible for converting a HTTP response into a OrgSiteList struct.
func (osl *OrgSiteList) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &osl.OrgSites)
}
