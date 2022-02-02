package main

import "github.com/mittacy/ego/hook"

func RegisterHook() {
	hook.Register(hook.BeforeStart, hook.AddGitCommitMsg)
}
