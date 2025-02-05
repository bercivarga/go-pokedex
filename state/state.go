package state

type AppState struct {
	LocationAreas []string
	AreaPage      int
}

func NewAppState() *AppState {
	return &AppState{
		LocationAreas: []string{},
		AreaPage:      0,
	}
}

func (state *AppState) Reset() {
	state.LocationAreas = []string{}
	state.AreaPage = 0
}

func (state *AppState) GetLocationAreas() []string {
	return state.LocationAreas
}

func (state *AppState) GetAreaPage() int {
	return state.AreaPage
}
