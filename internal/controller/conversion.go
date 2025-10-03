package controller

import (
	workshopv1alpha1 "golab.io/kubedredger/api/v1alpha1"
	"golab.io/kubedredger/internal/configfile"
	"k8s.io/utils/ptr"
)

const (
	ConditionAvailable   = "Available"
	ConditionProgressing = "Progressing"
	ConditionDegraded    = "Degraded"
)

const (
	ConditionReasonAsExpected      = "AsExpected"
	ConditionReasonUpToDate        = "UpToDate"
	ConditionReasonWriteError      = "WriteError"
	ConditionReasonUpdatingContent = "UpdatingContent"
	ConditionReasonUpdatingLabels  = "UpdatingLabels"
)

func configurationRequestFromSpec(desired workshopv1alpha1.ConfigurationSpec) configfile.ConfigRequest {
	res := configfile.ConfigRequest{
		Filename: desired.Filename,
		Content:  desired.Content,
		Create:   desired.Create,
	}
	if desired.Permission != nil {
		res.Permission = ptr.To(*desired.Permission)
	}
	return res
}

func statusFromConfStatus(desired workshopv1alpha1.ConfigurationSpec, confStatus configfile.ConfigurationStatus, labelErr error) workshopv1alpha1.ConfigurationStatus {
	// TODO: exercise: add code here
	return workshopv1alpha1.ConfigurationStatus{}
}
