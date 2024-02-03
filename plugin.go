package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

const remoteName = "drone"

func execute(cmd *exec.Cmd) error {
	fmt.Println("+", strings.Join(cmd.Args, " "))
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func pluginExec() error {
	if skipVerify {
		cmd := exec.Command("git", "config", "--global", "http.sslVerify", "false")
		if err := execute(cmd); err != nil {
			return err
		}
	}
	if sshKey == "" {
		return nil
	}
	home := "/root"
	if currentUser, err := user.Current(); err == nil {
		home = currentUser.HomeDir
	}
	sshpath := filepath.Join(home, ".ssh")
	if err := os.MkdirAll(sshpath, 0o700); err != nil {
		return err
	}
	_ = os.WriteFile(filepath.Join(sshpath, "config"), []byte("StrictHostKeyChecking no\n"), 0o700)
	err := os.WriteFile(filepath.Join(sshpath, "id_rsa"), []byte(sshKey), 0o600)
	if err != nil {
		return err
	}
	if remote != "" {
		cmd := exec.Command("git", "remote", "add", remoteName, remote)
		if err := execute(cmd); err != nil {
			return err
		}
	}
	ref := ""
	switch os.Getenv("DRONE_BUILD_EVENT") {
	case "tag":
		ref = os.Getenv("DRONE_TAG")
	case "push":
		ref = os.Getenv("DRONE_BRANCH")
	}
	if ref == "" {
		return errors.New("missing ref")
	}
	cmd := exec.Command("git", "push", remoteName, ref)
	if force {
		cmd.Args = append(cmd.Args, "--force")
	}
	return execute(cmd)
}
