package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strings"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var HTTP_OK = []int{200, 201, 202, 204}

func Contains[T comparable](arr []T, x T) bool {
	for _, v := range arr {
		if v == x {
			return true
		}
	}
	return false
}

func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func GetAttr(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	// if curStruct.Kind() != reflect.Struct {
	// 	panic("not struct")
	// }
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	// if !curField.IsValid() {
	// 	panic("not found:" + fieldName)
	// }
	return curField
}

func ConvertString(arr []types.String) []string {
	arrString := make([]string, len(arr))
	for i, v := range arr {
		arrString[i] = v.ValueString()
	}
	return arrString
}

func GetAuthenticator(url string, username string, password string, apiKey string) (core.Authenticator, error) {
	if strings.Contains(url, "cloud.ibm.com") {
		auth := core.NewIamAuthenticatorBuilder()
		auth.ApiKey = apiKey
		err := auth.Validate()
		return auth, err
	} else {
		if username != "" && apiKey != "" {
			auth, err := core.NewCloudPakForDataAuthenticatorUsingAPIKey(fmt.Sprintf("%s/icp4d-api", url), username, apiKey, true, map[string]string{})
			if err != nil {
				return nil, err
			}
			err = auth.Validate()
			return auth, err
		}

		if username != "" && password != "" {
			auth, err := core.NewCloudPakForDataAuthenticatorUsingPassword(fmt.Sprintf("%s/icp4d-api", url), username, password, true, map[string]string{})
			if err != nil {
				return nil, err
			}
			err = auth.Validate()
			return auth, err
		}
	}
	return nil, errors.New("unable to parse credentials")
}

func GetOpenScaleGuid(auth core.Authenticator) (string, error) {
	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(context.Background())
	URL, err := url.Parse("https://resource-controller.cloud.ibm.com/v2/resource_instances")
	if err != nil {
		return "", err
	}
	builder.URL = URL
	builder.AddHeader("Content-Type", "application/json")
	request, err := builder.Build()
	if err != nil {
		return "", err
	}
	err = auth.Authenticate(request)
	if err != nil {
		return "", err
	}

	serviceOptions := &core.ServiceOptions{
		Authenticator: auth,
	}

	service, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return "", err
	}

	response, err := service.Client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var jsonResponse map[string]interface{}
	err = json.Unmarshal(content, &jsonResponse)
	if err != nil {
		panic(err)
	}

	var openScaleGuid string
	resources := jsonResponse["resources"].([]interface{})
	for _, resource := range resources {
		resourceId := resource.(map[string]interface{})["resource_id"]
		if resourceId == "2ad019f3-0fd6-4c25-966d-f3952481a870" {
			openScaleGuid = resource.(map[string]interface{})["guid"].(string)
			break
		}
	}

	return openScaleGuid, nil
}
