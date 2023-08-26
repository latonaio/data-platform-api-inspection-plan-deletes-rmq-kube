package requests

type Header struct {
	InspectionPlan                 int     `json:"InspectionPlan"`
	IsMarkedForDeletion            bool    `json:"IsMarkedForDeletion"`
}
