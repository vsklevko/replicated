package types

import "time"

type KotsChannel struct {
	AdoptionRate               []CustomerAdoption            `json:"adoptionRate,omitempty"`
	AppId                      string                        `json:"appId,omitempty"`
	BuildAirgapAutomatically   bool                          `json:"buildAirgapAutomatically,omitempty"`
	ChannelIcon                string                        `json:"channelIcon,omitempty"`
	ChannelSequence            int32                         `json:"channelSequence,omitempty"`
	ChannelSlug                string                        `json:"channelSlug,omitempty"`
	Created                    time.Time                     `json:"created,omitempty"`
	CurrentVersion             string                        `json:"currentVersion,omitempty"`
	Customers                  *TotalActiveInactiveCustomers `json:"customers,omitempty"`
	Description                string                        `json:"description,omitempty"`
	EnterprisePartnerChannelID string                        `json:"enterprisePartnerChannelID,omitempty"`
	Id                         string                        `json:"id,omitempty"`
	IsArchived                 bool                          `json:"isArchived,omitempty"`
	IsDefault                  bool                          `json:"isDefault,omitempty"`
	Name                       string                        `json:"name,omitempty"`
	NumReleases                int32                         `json:"numReleases,omitempty"`
	ReleaseNotes               string                        `json:"releaseNotes,omitempty"`
	// TODO: set these (see kotsChannelToSchema function)
	ReleaseSequence int32            `json:"releaseSequence,omitempty"`
	Releases        []ChannelRelease `json:"releases,omitempty"`
	Updated         time.Time        `json:"updated,omitempty"`
}

type CustomerAdoption struct {
	ChannelId       string  `json:"channelId,omitempty"`
	Count           int32   `json:"count,omitempty"`
	Percent         float32 `json:"percent,omitempty"`
	ReleaseSequence int32   `json:"releaseSequence,omitempty"`
	Semver          string  `json:"semver,omitempty"`
	TotalOnChannel  int64   `json:"totalOnChannel,omitempty"`
}

type ChannelRelease struct {
	AirgapBuildError  string    `json:"airgapBuildError,omitempty"`
	AirgapBuildStatus string    `json:"airgapBuildStatus,omitempty"`
	ChannelIcon       string    `json:"channelIcon,omitempty"`
	ChannelId         string    `json:"channelId,omitempty"`
	ChannelName       string    `json:"channelName,omitempty"`
	ChannelSequence   int32     `json:"channelSequence,omitempty"`
	Created           time.Time `json:"created,omitempty"`
	RegistrySecret    string    `json:"registrySecret,omitempty"`
	ReleaseNotes      string    `json:"releaseNotes,omitempty"`
	ReleasedAt        time.Time `json:"releasedAt,omitempty"`
	Semver            string    `json:"semver,omitempty"`
	Sequence          int32     `json:"sequence,omitempty"`
	Updated           time.Time `json:"updated,omitempty"`
}

type CreateChannelRequest struct {
	// Description of the channel that is to be created.
	Description string `json:"description,omitempty"`
	// Enterprise Partner Channel Id to be added to channel.
	EnterprisePartnerChannelID string `json:"enterprisePartnerChannelID,omitempty"`
	Name                       string `json:"name"`
}

type UpdateChannelRequest struct {
	// Description of the channel that is to be created.
	Name           string `json:"name"`
	SemverRequired bool   `json:"semverRequired,omitempty"`
}

type Channel struct {
	ID          string
	Name        string
	Description string
	Slug        string

	ReleaseSequence int64
	ReleaseLabel    string

	IsArchived bool `json:"isArchived"`

	InstallCommands *InstallCommands
}

func (c *Channel) Copy() *Channel {
	return &Channel{
		ID:              c.ID,
		Name:            c.Name,
		Description:     c.Description,
		Slug:            c.Slug,
		ReleaseSequence: c.ReleaseSequence,
		ReleaseLabel:    c.ReleaseLabel,
		InstallCommands: c.InstallCommands,
	}
}

type InstallCommands struct {
	Existing string
	Embedded string
	Airgap   string
}
