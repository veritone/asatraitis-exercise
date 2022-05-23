package sosearch

const SEARCH_BY_JOB_TITLE = `
	query searchSdo($value: String!) {
		searchMedia(search:{
			index: ["mine"]
			query: {
				operator: "query_string"
				field: "sdo_onboarding_1_aj_6_jng_8848.series.jobTitle.fulltext" # Specific to SDO schema
				value: $value
			}
			offset: 0    
			limit: 10
		}){
			jsondata      
		}
	}
`