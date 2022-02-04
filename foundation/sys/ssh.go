package sys

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/crypto/ssh/knownhosts"
)

func SshAgentConnect(user, host string) {

	var hostKeyCallback ssh.HostKeyCallback
	hostKeyCallback, err := getHostKey(host)
	if err != nil {
		log.Fatalf("unable to retrieve host key: %v", err)
	}

	socket := os.Getenv("SSH_AUTH_SOCK")
	conn, err := net.Dial("unix", socket)
	if err != nil {
		log.Fatalf("Failed to open SSH_AUTH_SOCK: %v", err)
	}

	agentClient := agent.NewClient(conn)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// Use a callback rather than PublicKeys so we only consult the
			// agent once the remote server wants it.
			ssh.PublicKeysCallback(agentClient.Signers),
		},
		HostKeyCallback: hostKeyCallback,
	}

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}

func SshPassPhraseConnect(publicKeyFile, user, host string) {
	var hostKeyCallback ssh.HostKeyCallback
	hostKeyCallback, err := getHostKey(host)
	if err != nil {
		log.Fatalf("unable to retrieve host key: %v", err)
	}

	// If you have an encrypted private key, the crypto/x509 package
	// can be used to decrypt it.
	key, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKeyWithPassphrase(key, []byte("XXX"))
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
	}

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}

func getHostKey(host string) (ssh.HostKeyCallback, error) {
	knownHostFile := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")

	hostKeyCallback, err := knownhosts.New(knownHostFile)
	if err != nil {
		log.Fatal(err)
	}
	// scanner := bufio.NewScanner(file)
	// var hostKey ssh.PublicKey
	// for scanner.Scan() {
	// 	fields := strings.Split(scanner.Text(), " ")
	// 	if len(fields) != 3 {
	// 		continue
	// 	}
	// 	if strings.Contains(fields[0], host) {
	// 		var err error
	// 		hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
	// 		if err != nil {
	// 			return nil, errors.New(fmt.Sprintf("error parsing %q: %v", fields[2], err))
	// 		}
	// 		break
	// 	}
	// }

	if hostKeyCallback == nil {
		return nil, fmt.Errorf("no hostkey for %s", host)
	}
	return hostKeyCallback, nil
}
