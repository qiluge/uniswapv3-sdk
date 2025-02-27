package entities

// A data provider for ticks that is backed by an in-memory array of ticks.
type TickListDataProvider struct {
	Ticks []*Tick
}

func NewTickListDataProvider(ticks []*Tick, tickSpacing int) (*TickListDataProvider, error) {
	if err := ValidateList(ticks, tickSpacing); err != nil {
		return nil, err
	}
	return &TickListDataProvider{Ticks: ticks}, nil
}

func (p *TickListDataProvider) GetTick(tick int) *Tick {
	return GetTick(p.Ticks, tick)
}

func (p *TickListDataProvider) NextInitializedTickWithinOneWord(tick int, lte bool, tickSpacing int) (int, bool) {
	return NextInitializedTickWithinOneWord(p.Ticks, tick, lte, tickSpacing)
}
