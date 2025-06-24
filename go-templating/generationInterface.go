package main

import (
	"path/filepath"
)

// Inputs used across all generation types.
type sharedInput struct {
	// The destination folder where generation takes place.
	DestinationDirPrefix string `yaml:"destinationDirPrefix" example:"~/projects/external"`
	// Variables required that are shared by resource, pipeline, and service templates.
	RequiredInputs map[string]interface{} `yaml:"requiredInputs"`
}

// Inputs for resource generation. Variables are shared with service and pipeline module.
type resourceInput struct {
	DirectoryName string `yaml:"directoryName" example:"shared-resources"`
	// Only used by internal templating code to determine which template to use.
	TemplateName string `yaml:"templateName" example:"resourcesTemplate"`
}

// Inputs for pipeline generation. Variables are shared with service and resource module.
type pipelineInput struct {
	// Used by both internal go code and required for templates.
	// Directory name shared as both shared-resources and service module reference it
	// to access pipeline files.
	DirectoryName string `yaml:"directoryName" example:"pipeline-files"`
	// Variables required only by the pipeline templates.
	RequiredInputs map[string]interface{} `yaml:"requiredInputs"`
	// Only used by internal templating code to determine which template to use.
	TemplateName string `yaml:"templateName" example:"pipelineTemplate"`
}

// Inputs for service generation. Variables are shared with pipeline module.
type serviceInput struct {
	// Directory name is also used by pipeline module to access service pipeline files.
	DirectoryName string `yaml:"directoryName" example:"mygreeterv3"`
	ServiceName   string `yaml:"serviceName" example:"MyGreeter"`
	// Used by pipeline module to determine whether or not the service needs to be
	// run in general overarching pipeline.
	RunPipeline bool `yaml:"runPipeline" example:"true"`
	// Variables required only by the pipeline templates.
	RequiredInputs map[string]interface{} `yaml:"requiredInputs"`
	// Only used by internal templating code to determine which template to use.
	TemplateName string `yaml:"templateName" example:"mygreeterGoTemplate"`
}

type allInput struct {
	SharedInput   sharedInput   `yaml:"sharedInput"`
	ResourceInput resourceInput `yaml:"resourceInput"`
	PipelineInput pipelineInput `yaml:"pipelineInput"`
	// Inputs generated/outputted by terraform resource creation
	EnvInformation map[string]string `yaml:"envInformation"`
	ServiceInput   serviceInput      `yaml:"serviceInput"`

	// The following variables have tags to maintain consistency
	// in variable naming when structs are converted to maps to pass into
	// template executions.

	// Used to store service information generation.
	ServiceDirectoryNameToRunPipeline map[string]bool   `name:"serviceDirectoryNameToRunPipeline"`
	ExtraInputs                       map[string]string `name:"extraInputs"`
	User                              string            `name:"user"`
	// Only used before templating execution, inaccessible to template files.
	GeneratorInputs map[string]string
}

// This function generates the variables used for templating. It returns
// the destination string and the template name based on generationType
func getTemplateVars(allInput allInput, generationType string) (string, string) {
	switch generationType {
	case serviceStr:
		return filepath.Join(allInput.SharedInput.DestinationDirPrefix, allInput.ServiceInput.DirectoryName), allInput.ServiceInput.TemplateName
	case resourceStr:
		return filepath.Join(allInput.SharedInput.DestinationDirPrefix, allInput.ResourceInput.DirectoryName), allInput.ResourceInput.TemplateName
	case pipelineStr:
		return filepath.Join(allInput.SharedInput.DestinationDirPrefix, allInput.PipelineInput.DirectoryName), allInput.PipelineInput.TemplateName
	default:
		return "", ""
	}
}
