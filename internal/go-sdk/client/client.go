package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"terraform-provider-ibmcpd/internal/go-sdk/spacev2"
	"terraform-provider-ibmcpd/internal/go-sdk/watsonmachinelearningv4"
	"terraform-provider-ibmcpd/internal/go-sdk/watsonopenscalev2"
	"terraform-provider-ibmcpd/internal/utils"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-cleanhttp"
)

const (
	WatsonStudioService             = "WatsonStudio"
	WatsonMachineLearningAPIVersion = "2021-12-01"
)

var (
	versionTime, _                      = time.Parse("2006-01-02", WatsonMachineLearningAPIVersion)
	WatsonMachineLearningAPIVersionDate = strfmt.Date(versionTime)
)

type Config struct {
	URL string
}

type Client struct {
	Config        *Config
	space         *spacev2.SpaceV2
	wml           *watsonmachinelearningv4.WatsonMachineLearningV4
	wos           *watsonopenscalev2.WatsonOpenScaleV2
	authenticator core.Authenticator
}

func (c *Client) useDefaultRoundTripper(service *core.BaseService) {
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	transport := cleanhttp.DefaultPooledTransport()
	transport.TLSClientConfig = tlsConfig
	service.Client.Transport = transport
	// service.Client.Transport = c.Config.defaultRoundTripper
}

func (c *Config) IsPublicCloud() bool {
	return strings.Contains(c.URL, "cloud.ibm.com")
}

func NewClient(authenticator core.Authenticator, config *Config) (*Client, error) {
	client := &Client{
		authenticator: authenticator,
		Config:        config,
	}
	return client, nil
}

func (c *Client) SpaceClient(ctx context.Context) (*spacev2.SpaceV2, error) {
	var err error
	if c.space == nil {
		serviceOptions := spacev2.SpaceV2Options{
			URL:           utils.If(strings.Contains(c.Config.URL, "cloud.ibm.com"), "https://api.dataplatform.cloud.ibm.com", c.Config.URL),
			Authenticator: c.authenticator,
		}
		c.space, err = spacev2.NewSpaceV2(&serviceOptions)
		if err != nil {
			return nil, err
		}
		c.useDefaultRoundTripper(c.space.Service)
	}
	return c.space, err
}

func (c *Client) WMLClient(ctx context.Context) (*watsonmachinelearningv4.WatsonMachineLearningV4, error) {
	var err error
	if c.wml == nil {
		serviceOptions := watsonmachinelearningv4.WatsonMachineLearningV4Options{
			URL:           c.Config.URL + "/ml",
			Authenticator: c.authenticator,
			Version:       &WatsonMachineLearningAPIVersionDate,
		}
		c.wml, err = watsonmachinelearningv4.NewWatsonMachineLearningV4(&serviceOptions)
		if err != nil {
			return nil, err
		}
		c.useDefaultRoundTripper(c.wml.Service)
	}
	return c.wml, err
}

func (c *Client) WOSClient(ctx context.Context) (*watsonopenscalev2.WatsonOpenScaleV2, error) {
	var err error
	if c.wos == nil {
		var url string
		if strings.Contains(c.Config.URL, "cloud.ibm.com") {
			guid, err := utils.GetOpenScaleGuid(c.authenticator)
			if err != nil {
				return nil, err
			}
			url = fmt.Sprintf("https://api.aiopenscale.cloud.ibm.com/openscale/%s", guid)
		} else {
			url = fmt.Sprintf("%s/openscale/00000000-0000-0000-0000-000000000000", c.Config.URL)
		}

		serviceOptions := watsonopenscalev2.WatsonOpenScaleV2Options{
			URL:           url,
			Authenticator: c.authenticator,
		}
		c.wos, err = watsonopenscalev2.NewWatsonOpenScaleV2(&serviceOptions)
		if err != nil {
			return nil, err
		}
		c.useDefaultRoundTripper(c.wos.Service)
	}
	return c.wos, err
}
