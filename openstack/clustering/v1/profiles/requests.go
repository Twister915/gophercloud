package profiles

import (
	"net/http"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// CreateOptsBuilder for options used for creating a profile.
type CreateOptsBuilder interface {
	ToProfileCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used for creating a profile
type CreateOpts struct {
	Name     string                 `json:"name" required:"true"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Spec     Spec                   `json:"spec" required:"true"`
}

// ToProfileCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToProfileCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "profile")
}

// Create requests the creation of a new profile on the server.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToProfileCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	var result *http.Response
	result, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})

	if r.Err == nil {
		r.Header = result.Header
	}
	return
}

// Get retrieves detail of a single profile. Use Extract to convert its
// result into a Profile.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	var result *http.Response
	result, r.Err = client.Get(getURL(client, id), &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})

	if r.Err == nil {
		r.Header = result.Header
	}
	return
}

// ListOptsBuilder Builder.
type ListOptsBuilder interface {
	ToProfileListQuery() (string, error)
}

// ListOpts params
type ListOpts struct {
	GlobalProject *bool  `q:"global_project"`
	Limit         int    `q:"limit"`
	Marker        string `q:"marker"`
	Name          string `q:"name"`
	Sort          string `q:"sort"`
	Type          string `q:"type"`
}

// ToProfileListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToProfileListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List instructs OpenStack to provide a list of profiles.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToProfileListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ProfilePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToProfileUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts implements Profile's UpdateOpts
type UpdateOpts struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name     string                 `json:"name,omitempty"`
}

// ToProfileUpdateMap assembles a request body based on the contents of
// UpdateOpts.
func (opts UpdateOpts) ToProfileUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "profile")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Update implements profile update request.
func Update(client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToProfileUpdateMap()
	if err != nil {
		r.Err = err
		return r
	}
	var result *http.Response
	result, r.Err = client.Patch(updateURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	if r.Err == nil {
		r.Header = result.Header
	}
	return
}
