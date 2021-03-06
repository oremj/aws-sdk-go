package elasticbeanstalkiface

import (
	"github.com/awslabs/aws-sdk-go/service/elasticbeanstalk"
)

type ElasticBeanstalkAPI interface {
	CheckDNSAvailability(*elasticbeanstalk.CheckDNSAvailabilityInput) (*elasticbeanstalk.CheckDNSAvailabilityOutput, error)

	CreateApplication(*elasticbeanstalk.CreateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error)

	CreateApplicationVersion(*elasticbeanstalk.CreateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)

	CreateConfigurationTemplate(*elasticbeanstalk.CreateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error)

	CreateEnvironment(*elasticbeanstalk.CreateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)

	CreateStorageLocation(*elasticbeanstalk.CreateStorageLocationInput) (*elasticbeanstalk.CreateStorageLocationOutput, error)

	DeleteApplication(*elasticbeanstalk.DeleteApplicationInput) (*elasticbeanstalk.DeleteApplicationOutput, error)

	DeleteApplicationVersion(*elasticbeanstalk.DeleteApplicationVersionInput) (*elasticbeanstalk.DeleteApplicationVersionOutput, error)

	DeleteConfigurationTemplate(*elasticbeanstalk.DeleteConfigurationTemplateInput) (*elasticbeanstalk.DeleteConfigurationTemplateOutput, error)

	DeleteEnvironmentConfiguration(*elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*elasticbeanstalk.DeleteEnvironmentConfigurationOutput, error)

	DescribeApplicationVersions(*elasticbeanstalk.DescribeApplicationVersionsInput) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error)

	DescribeApplications(*elasticbeanstalk.DescribeApplicationsInput) (*elasticbeanstalk.DescribeApplicationsOutput, error)

	DescribeConfigurationOptions(*elasticbeanstalk.DescribeConfigurationOptionsInput) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error)

	DescribeConfigurationSettings(*elasticbeanstalk.DescribeConfigurationSettingsInput) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error)

	DescribeEnvironmentResources(*elasticbeanstalk.DescribeEnvironmentResourcesInput) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error)

	DescribeEnvironments(*elasticbeanstalk.DescribeEnvironmentsInput) (*elasticbeanstalk.DescribeEnvironmentsOutput, error)

	DescribeEvents(*elasticbeanstalk.DescribeEventsInput) (*elasticbeanstalk.DescribeEventsOutput, error)

	ListAvailableSolutionStacks(*elasticbeanstalk.ListAvailableSolutionStacksInput) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error)

	RebuildEnvironment(*elasticbeanstalk.RebuildEnvironmentInput) (*elasticbeanstalk.RebuildEnvironmentOutput, error)

	RequestEnvironmentInfo(*elasticbeanstalk.RequestEnvironmentInfoInput) (*elasticbeanstalk.RequestEnvironmentInfoOutput, error)

	RestartAppServer(*elasticbeanstalk.RestartAppServerInput) (*elasticbeanstalk.RestartAppServerOutput, error)

	RetrieveEnvironmentInfo(*elasticbeanstalk.RetrieveEnvironmentInfoInput) (*elasticbeanstalk.RetrieveEnvironmentInfoOutput, error)

	SwapEnvironmentCNAMEs(*elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*elasticbeanstalk.SwapEnvironmentCNAMEsOutput, error)

	TerminateEnvironment(*elasticbeanstalk.TerminateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)

	UpdateApplication(*elasticbeanstalk.UpdateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error)

	UpdateApplicationVersion(*elasticbeanstalk.UpdateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)

	UpdateConfigurationTemplate(*elasticbeanstalk.UpdateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error)

	UpdateEnvironment(*elasticbeanstalk.UpdateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)

	ValidateConfigurationSettings(*elasticbeanstalk.ValidateConfigurationSettingsInput) (*elasticbeanstalk.ValidateConfigurationSettingsOutput, error)
}