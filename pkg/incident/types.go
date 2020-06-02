package incident

type (
	createIncident struct {
		Email string `json:"email"`
		Issue string `json:"issue"`
	}

	updateIncident struct {
		Email string `json:"email"`
		Issue string `json:"issue"`
	}
)
