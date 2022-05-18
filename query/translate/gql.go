package translateJob

const TRANSLATE_JOB_QUERY = `
	mutation create_translate_job($clusterId: ID!, $url: String!, $language: String!) {
		createJob(
			input: {
				target: { status: "downloaded" }
				clusterId: $clusterId
				## Tasks
				tasks: [
					{
						# Chunk engine
						engineId: "8bdb0e3b-ff28-4f6e-a3ba-887bd06e6440"
						payload: {
							url: $url
							ffmpegTemplate: "rawchunk"
						}
						executionPreferences: {
							priority: -92
							parentCompleteBeforeStarting: true
						}
						ioFolders: [{ referenceId: "si-out", mode: chunk, type: output }]
					}
					{
						# Amazon translate
						engineId: "1fc4d3d4-54ab-42d1-882c-cfc9df42f386"
						payload: { target: $language }
						executionPreferences: {
							priority: -92
							parentCompleteBeforeStarting: true
						}
						ioFolders: [
							{ referenceId: "engine-in", mode: chunk, type: input }
							{ referenceId: "engine-out", mode: chunk, type: output }
						]
					}
					{
						# output writer
						engineId: "8eccf9cc-6b6d-4d7d-8cb3-7ebf4950c5f3"
						executionPreferences: {
							priority: -92
							parentCompleteBeforeStarting: true
						}
						ioFolders: [{ referenceId: "ow-in", mode: chunk, type: input }]
					}
				]
				##Routes
				routes: [
					{
						## chunkAudio --> translation
						parentIoFolderReferenceId: "si-out"
						childIoFolderReferenceId: "engine-in"
						options: {}
					}
					{
						## sampleChunkOutputFolderA  --> ow1
						parentIoFolderReferenceId: "engine-out"
						childIoFolderReferenceId: "ow-in"
						options: {}
					}
				]
			}
		) {
			id
			targetId
			createdDateTime
		}
	}
`