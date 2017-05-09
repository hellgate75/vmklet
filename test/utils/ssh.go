package utils

func SShGoLangInstall(user string, ip string, port int, identityFilePath string) {
	if port == 0 {
		port = 22
	}
	//commander := procedures.SSHCommandStream{
	//	User:         user,
	//	IP:           ip,
	//	Port:         port,
	//	IdentityFile: identityFilePath,
	//}
	//
	//command := []string{
	//	"apt-get",
	//	"install",
	//	"-y",
	//	"golang-go",
	//}

	//cmd := commander.executeOne(command...)
	//reader, err := cmd.StdoutPipe()
	//defer  reader.Close()
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, "There was an error running SSH command: ", err)
	//	os.Exit(1)
	//}
	//
	//scanner := bufio.NewScanner(reader)
	//
	//for scanner.Scan() {
	//	err = scanner.Err()
	//	if err == nil {
	//		fmt.Printf("%s\n", string(scanner.Bytes()))
	//	} else {
	//		fmt.Fprintln(os.Stderr, "Unable to read SSH output: ", err)
	//	}
	//}
	//fmt.Fprintln(os.Stdout, fmt.Sprintf("SSH Connection to %s:%d with user %s closed!!", ip, port, user))
}
