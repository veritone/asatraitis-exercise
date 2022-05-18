package status

const JOB_STATUS_QUERY = `
	query getJobStatus($id: ID!) {
		job(id: $id) {
			id
			targetId
			clusterId
			status
			createdDateTime
				tasks {
					records {
						id
						engine {
							name
						}
						status
					}
				}
		}
	}
`