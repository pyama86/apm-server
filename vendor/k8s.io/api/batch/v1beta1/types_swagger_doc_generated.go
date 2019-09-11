/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_CronJob = map[string]string{
	"":         "CronJob represents the configuration of a single cron job.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"spec":     "Specification of the desired behavior of a cron job, including the schedule. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
	"status":   "Current status of a cron job. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
}

func (CronJob) SwaggerDoc() map[string]string {
	return map_CronJob
}

var map_CronJobList = map[string]string{
	"":         "CronJobList is a collection of cron jobs.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"items":    "items is the list of CronJobs.",
}

func (CronJobList) SwaggerDoc() map[string]string {
	return map_CronJobList
}

var map_CronJobSpec = map[string]string{
	"":                           "CronJobSpec describes how the job execution will look like and when it will actually run.",
	"schedule":                   "The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.",
	"startingDeadlineSeconds":    "Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.",
	"concurrencyPolicy":          "Specifies how to treat concurrent executions of a Job. Valid values are: - \"Allow\" (default): allows CronJobs to run concurrently; - \"Forbid\": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - \"Replace\": cancels currently running job and replaces it with a new one",
	"suspend":                    "This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.",
	"jobTemplate":                "Specifies the job that will be created when executing a CronJob.",
	"successfulJobsHistoryLimit": "The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. Defaults to 3.",
	"failedJobsHistoryLimit":     "The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.",
}

func (CronJobSpec) SwaggerDoc() map[string]string {
	return map_CronJobSpec
}

var map_CronJobStatus = map[string]string{
	"":                 "CronJobStatus represents the current state of a cron job.",
	"active":           "A list of pointers to currently running jobs.",
	"lastScheduleTime": "Information when was the last time the job was successfully scheduled.",
}

func (CronJobStatus) SwaggerDoc() map[string]string {
	return map_CronJobStatus
}

var map_JobTemplate = map[string]string{
	"":         "JobTemplate describes a template for creating copies of a predefined pod.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"template": "Defines jobs that will be created from this template. https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
}

func (JobTemplate) SwaggerDoc() map[string]string {
	return map_JobTemplate
}

var map_JobTemplateSpec = map[string]string{
	"":         "JobTemplateSpec describes the data a Job should have when created from a template",
	"metadata": "Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"spec":     "Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
}

func (JobTemplateSpec) SwaggerDoc() map[string]string {
	return map_JobTemplateSpec
}

// AUTO-GENERATED FUNCTIONS END HERE