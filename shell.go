package main

import (
   "bufio"
   "fmt"
   "os"
   "go_shell/exec"
   "go_shell/colors"
   "github.com/google/shlex"
)

func main() {
   // cls := "\033[2J \033[H"
   // print(cls)

   for {
      cmd := prompt();
      if (cmd == "exit") {
         fmt.Println(cmd);
         break;
      }
      if len(cmd) == 0 {
         fmt.Print("\n");
         continue;
      }
      args, err := shlex.Split(cmd);
      if err != nil {
         panic(err);
      }
      go exec.Exec(args...);
   }

   fmt.Println(prompt)
}

func prompt() string {
   scanner := bufio.NewScanner(os.Stdin);
   fmt.Print(colors.GRN + "(go_shell)" + colors.RED + " $ " + colors.RESET);
   scanner.Scan();
   prompt := scanner.Text();
   return prompt;
}