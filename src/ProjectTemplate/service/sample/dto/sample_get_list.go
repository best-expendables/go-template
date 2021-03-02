package dto

import (
	"fmt"
	"strings"

	"github.com/best-expendables/common-utils/repository/filter"
)

type ${SERVICE_NAME}GetListFilter struct {
	filter.PaginationFilter
}

func New${SERVICE_NAME}GetListFilter() *${SERVICE_NAME}GetListFilter {
	return &${SERVICE_NAME}GetListFilter{
		PaginationFilter: *filter.NewPaginationFilter(),
	}
}

func (f *${SERVICE_NAME}GetListFilter) GetWhere() filter.Where {
	return f.BasicFilter.GetWhere()
}

func (f ${SERVICE_NAME}GetListFilter) GetOrderBy() []string {
	orderByClauses := make([]string, 0)
	originOrderByClauses := f.PaginationFilter.BasicOrder.GetOrderBy()
	for i := 0; i < len(originOrderByClauses); i++ {
		modifiedClause := f.assignTableToOrderByClause(originOrderByClauses[i])
		if modifiedClause == "" {
			continue
		}
		orderByClauses = append(orderByClauses, modifiedClause)
	}

	return orderByClauses
}

func (f ${SERVICE_NAME}GetListFilter) assignTableToOrderByClause(s string) string {
	words := strings.Split(s, " ")
	if len(words) != 2 {
		return ""
	}
	switch words[0] {
	case "id":
		return fmt.Sprintf("%s %s", "id", words[1])
	case "createdAt":
		return fmt.Sprintf("%s %s", "created_at", words[1])
	case "updatedAt":
		return fmt.Sprintf("%s %s", "updated_at", words[1])
	default:
		return ""
	}
}
