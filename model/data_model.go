package data_model

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents {
	{
		ID: "1",
		Title: "Basic Go api application.",
		Description: "Test application."
	}
}
