package convert

import (
	"lagou_jobs/downloader"
	"lagou_jobs/pipeline"
	"strconv"
	"strings"
	"time"
)

func ToPipelineJobs(dJobs []downloader.Result) []pipeline.LgJob {
	var pJobs []pipeline.LgJob
	for _, v := range dJobs {
		longitude, _ := strconv.ParseFloat(v.Longitude, 64)
		latitude, _ := strconv.ParseFloat(v.Latitude, 64)
		pJobs = append(pJobs, pipeline.LgJob{
			City: 		v.City,
			District: 	v.District,

			CompanyShortName: v.CompanyShortName,
			CompanyFullName:  v.CompanyFullName,
			CompanyLabelList: strings.Join(v.CompanyLabelList,","),
			CompanySize: v.CompanySize,
			FinanceStage: v.FinancesStage,

			PositionName: v.PositionName,
			PositionLables: strings.Join(v.PositionLabels,","),
			PositionAdvantage: v.PositionAdvantage,
			WorkYear: v.WorkYear,
			Education: v.Education,
			Salary: v.Salary,

			IndustryField: v.IndustryField,
			IndustryLables: strings.Join(v.IndustryLabels, ","),

			Longitude: longitude,
			Latitude: latitude,
			Linestaion: v.Linestation,

			CreateTime: MustDateToUnix(v.CreateTime),
			AddTime: time.Now().Unix(),
		})
	}
	return pJobs
}