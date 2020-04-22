package page

func CaculateTotalPage(totalCount, pageSize float64) int {
	totalPage := float64(totalCount) / float64(pageSize)

	return int(totalPage)
}
