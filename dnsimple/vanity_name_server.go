package dnsimple

import (
	"fmt"
)

// VanityNameServersService handles communication with Vanity Name Servers
// methods of the DNSimple API.
//
// See https://developer.dnsimple.com/v2/vanity/
type VanityNameServersService struct {
	client *Client
}

// VanityNameServer represents data for a single vanity name server
type VanityNameServer struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	IPv4      string `json:"ipv4,omitempty"`
	IPv6      string `json:"ipv6,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// VanityNameServerResponse represents a response for vanity name server enable and disable operations.
type VanityNameServerResponse struct {
	Response
	Data []VanityNameServer `json:"data"`
}

func vanityNameServerPath(accountID string, domainID string) string {
	return fmt.Sprintf("/%v/vanity/%v", accountID, domainID)
}

// Enable Vanity Name Servers for the given domain
//
// See https://developer.dnsimple.com/v2/vanity/#enable
func (s *VanityNameServersService) Enable(accountID string, domainID string) (*VanityNameServerResponse, error) {
	path := versioned(vanityNameServerPath(accountID, domainID))
	vanityNameServerResponse := &VanityNameServerResponse{}

	resp, err := s.client.put(path, nil, vanityNameServerResponse)
	if err != nil {
		return nil, err
	}

	vanityNameServerResponse.HttpResponse = resp
	return vanityNameServerResponse, nil
}

// Disable Vanity Name Servers for the given domain
//
// See https://developer.dnsimple.com/v2/vanity/#disable
func (s *VanityNameServersService) Disable(accountID string, domainID string) (*VanityNameServerResponse, error) {
	path := versioned(vanityNameServerPath(accountID, domainID))
	vanityNameServerResponse := &VanityNameServerResponse{}

	resp, err := s.client.delete(path, nil, nil)
	if err != nil {
		return nil, err
	}

	vanityNameServerResponse.HttpResponse = resp
	return vanityNameServerResponse, nil
}