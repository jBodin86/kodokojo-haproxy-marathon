package utils

import (
	"github.com/kodokojo/kodokojo-haproxy-marathon/commons"
	"regexp"
)

const projectNameRegexp string = "/(?P<projectName>.*)/(?P<entityType>.*)"

type ServiceLocator interface {
	LocateServiceByType(projectName string, entityType string) ([]commons.Service, bool)

	LocateAllService() ([]commons.Service, bool)

	LocateServiceByProject(projectName string) ([]commons.Service, bool)
}

func GetAppIdMatchKodokojoProjectName(appId string) (projectName, entityName string) {

	r := regexp.MustCompile(projectNameRegexp)
	namesRegexp := r.SubexpNames()
	submatch := r.FindAllStringSubmatch(appId, -1)
	if submatch != nil {
		result := submatch[0]

		group := map[string]string{}
		for i, value := range result {
			group[namesRegexp[i]] = value
		}
		if len(group) >= 2 {
			projectName = group["projectName"]
			entityName = group["entityType"]
		}
	}
	return
}