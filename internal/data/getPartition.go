package data

import (
	models "dormitory/internal/models"
	"fmt"
	"time"
)

type IPartition interface {
	GetPartition() string
}

// Factory
type PartitionQuery struct {
	partition map[int]IPartition
}

func (p *PartitionQuery) SetPartition(year int, partition IPartition) {
	p.partition[year] = partition
}

func (p *PartitionQuery) GetPartition(year int) string {
	partition, ok := p.partition[year]
	if !ok {
		return "contracts"
	}

	return partition.GetPartition()
}

func (p PartitionQuery) GetMonthFromTo(fromMonth, toMonth *time.Time) (from int, to int) {
	if fromMonth != nil && !fromMonth.IsZero() {
		from = int(fromMonth.Month())
	}
	if toMonth != nil && !toMonth.IsZero() {
		to = int(toMonth.Month())
	}

	return from, to
}

func NewPartitionQuery() *PartitionQuery {
	return &PartitionQuery{
		partition: map[int]IPartition{},
	}
}

// Concrete partition 2025
type Partition2025 struct {
	monthFrom int
	monthTo   int
}

func (p Partition2025) GetPartitionTwoTimeline() string {
	var partition string
	for i := p.monthFrom; i <= p.monthTo; i++ {
		partition += fmt.Sprintf("p%02d", i)
		if i < p.monthTo {
			partition += ", "
		}
	}

	return fmt.Sprintf("contracts PARTITION (%s)", partition)
}

func (p Partition2025) GetPartitionOneTimeline() string {
	var timeline int
	if p.monthFrom != 0 {
		timeline = p.monthFrom
	}
	if p.monthTo != 0 {
		timeline = p.monthTo
	}
	return fmt.Sprintf("contracts PARTITION (p%02d)", timeline)
}

func (p Partition2025) GetPartition() string {
	if p.monthFrom == 0 || p.monthTo == 0 {
		return p.GetPartitionOneTimeline()
	}

	return p.GetPartitionTwoTimeline()
}

func NewPartition2025(monthFrom, monthTo int) *Partition2025 {
	return &Partition2025{monthFrom: monthFrom, monthTo: monthTo}
}

func CallPartition2025(filter *models.ContractFilter) string {
	partition := NewPartitionQuery()
	var partitionStr string
	var monthFrom, monthTo int
	if filter != nil {
		monthFrom, monthTo = partition.GetMonthFromTo(filter.RegistryAt.FromTime, filter.RegistryAt.ToTime)
	}
	if monthFrom == 0 && monthTo == 0 {
		return partition.GetPartition(0)
	}
	partition2025 := NewPartition2025(monthFrom, monthTo)
	partition.SetPartition(2025, partition2025)
	partitionStr = partition.GetPartition(2025)

	return partitionStr
}
