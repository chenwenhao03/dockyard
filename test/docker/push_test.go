package docker

import (
	"os/exec"
	"testing"
	//"github.com/containerops/dockyard/test"
)

func TestPushInit(t *testing.T) {
	repoBase := "busybox:latest"

	if err := exec.Command(DockerBinary, "inspect", repoBase).Run(); err != nil {
		cmd := exec.Command(DockerBinary, "pull", repoBase)
		if out, err := ParseCmdCtx(cmd); err != nil {
			t.Fatalf("Push testing preparation is failed: [Info]%v, [Error]%v", out, err)
		}
	}
}

func TestPushRepoWithSingleTag(t *testing.T) {
	var cmd *exec.Cmd
	var err error
	var out string

	reponame := "busybox"
	repotag := "latest"
	repoBase := reponame + ":" + repotag

	repoDest := Domains + "/" + UserName + "/" + repoBase
	cmd = exec.Command(DockerBinary, "tag", "-f", repoBase, repoDest)
	if out, err = ParseCmdCtx(cmd); err != nil {
		t.Fatalf("Tag %v failed: [Info]%v, [Error]%v", repoBase, out, err)
	}

	//push the same repository with specified tag more than once to cover related code processing branch
	for i := 1; i <= 2; i++ {
		cmd = exec.Command(DockerBinary, "push", repoDest)
		if out, err = ParseCmdCtx(cmd); err != nil {
			t.Fatalf("Push %v failed: [Info]%v, [Error]%v", repoDest, out, err)
		}
	}

	cmd = exec.Command(DockerBinary, "rmi", repoDest)
	if out, err = ParseCmdCtx(cmd); err != nil {
		t.Fatalf("Romove image %v failed: [Info]%v, [Error]%v", repoDest, out, err)
	}
}

func TestPushRepoWithMultipleTags(t *testing.T) {
	var cmd *exec.Cmd
	var err error
	var out string

	reponame := "busybox"
	repotags := []string{"latest", "1.0", "2.0"}
	repoBase := reponame + ":" + repotags[0] //pull busybox:latest from docker hub

	repoDest := Domains + "/" + UserName + "/" + reponame
	for _, v := range repotags {
		tag := repoDest + ":" + v
		cmd = exec.Command(DockerBinary, "tag", "-f", repoBase, tag)
		if out, err = ParseCmdCtx(cmd); err != nil {
			t.Fatalf("Tag %v failed: [Info]%v, [Error]%v", repoBase, out, err)
		}
	}

	//push the same repository with multiple tags more than once to cover related code processing branch
	for i := 1; i <= 2; i++ {
		cmd = exec.Command(DockerBinary, "push", repoDest)
		if out, err = ParseCmdCtx(cmd); err != nil {
			t.Fatalf("Push all tags %v failed: [Info]%v, [Error]%v", repoDest, out, err)
		}
	}

	for _, v := range repotags {
		tag := repoDest + ":" + v
		cmd = exec.Command(DockerBinary, "rmi", tag)
		if out, err = ParseCmdCtx(cmd); err != nil {
			t.Fatalf("Romove image %v failed: [Info]%v, [Error]%v", repoDest, out, err)
		}
	}
}
