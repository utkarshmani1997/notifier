package incident

type IncidentStore interface {
	Create(string, string) (uint, error)
	Get(string) (Incident, error)
	Update(string, string) (Incident, error)
	Delete(string) (uint, error)
}
