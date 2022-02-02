package async_config

import (
	"github.com/mittacy/ego-layout/app/job/job_payload"
	"github.com/mittacy/ego-layout/app/job/job_process"
	"github.com/mittacy/ego/library/async"
)

func Jobs() []async.Job {
	return []async.Job{
		{
			job_payload.ExampleTypeName,
			job_process.NewExample(),
		},
	}
}
