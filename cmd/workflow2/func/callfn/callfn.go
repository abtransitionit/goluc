package callfn

import (
	"context"
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

func CheckVmSshAccess(ctx context.Context, node []string, logger logx.Logger) (string, error) {
	fmt.Println("called function CheckVmSshAccess with nodes", node)
	return "", nil
}
func DeployLuc(ctx context.Context, node []string, logger logx.Logger) (string, error) {
	fmt.Println("called function DeployLuc with nodes", node)
	return "", nil
}
func UpgradeVmOs(ctx context.Context, node []string, logger logx.Logger) (string, error) {
	fmt.Println("called function UpgradeVmOs with nodes", node)
	return "", nil
}
func UpdateVmOsApp(ctx context.Context, node []string, logger logx.Logger) (string, error) {
	fmt.Println("called function UpdateVmOsApp with nodes", node)
	return "", nil
}

func generic(ctx context.Context, logger logx.Logger) (string, error) {
	fmt.Println("called function generic")
	return "", nil
}

func yoyo(monEntier string) error {
	fmt.Println("called yoyo with param", monEntier)
	return nil
}
