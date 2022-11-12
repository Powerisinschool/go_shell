package exec

import (
	"fmt"
	// "log"
	"os"
	// "os/exec"
	"syscall"
)

func Exec(args ...string) {
	fmt.Printf("%s: pid is %d\n", os.Args[0], os.Getpid());
	_ = syscall.Exec(args[0], args, os.Environ());
	// runner := exec.Command(args[0]);
	// output, err := runner.CombinedOutput();
	// if err != nil {
	// 	log.Fatal(err);
	// }
	// fmt.Print(string(output));
}
