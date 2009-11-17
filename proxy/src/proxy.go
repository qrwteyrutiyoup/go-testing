package proxy

import (
	"os";
	"fmt";
	"net";
	"bufio";
	"strings";
	"strconv";
	"./general";
	"./config";
)

func makeConnection(host string, port int) (net.Conn)
{
	target := fmt.Sprintf("%s:%d", host, port);

	connection, err := net.Dial("tcp", "", target);
	if err != nil
	{
		general.Error(err);
		return nil;
	}

	return connection;
}

func readData(source, destination net.Conn)
{
	defer source.Close();
	reader := bufio.NewReader(source);

	for
	{
		line, err := reader.ReadString('\n');
		if err != nil 
		{
			if err != os.EOF
			{
				fmt.Printf("--> %#v\n", err);
			}
			break;
		}

		fmt.Printf("[%s]\n", line);
		i, err := destination.Write(strings.Bytes(strings.TrimSpace(line) + "\r\n"));
		if err != nil || i < len(line)
		{
			break;
		}
	}

	fmt.Println("goodbye cruel world");
}

func doIt(client net.Conn, config map[string]string)
{
	host := config["targethost"];
	port, err := strconv.Atoi(config["targetport"]);
	if err != nil
	{
		general.Error(err);
		return;
	}

	server := makeConnection(host, port);
	if server == nil
	{
		fmt.Printf("cannot establish a connection with %s:%d\n", host, port);
		return;
	}

	go readData(server, client);
	go readData(client, server);
}

func startServer(config map[string]string)
{
	listening, err := strconv.Atoi(config["listenport"]);
	if err != nil
	{
		general.Error(err);
		return;
	}

	target := fmt.Sprintf("127.0.0.1:%d", listening);	
	address, err := net.ResolveTCPAddr(target);
	if err != nil
	{
		general.Error(err);
		return;
	}

	listener, err := net.ListenTCP("tcp", address);
	if err != nil
	{
		general.Error(err);
		return;
	}



	for 
	{
		var connection net.Conn;

		fmt.Printf("listening for new connections at %s\n", target);
		connection, err := listener.AcceptTCP();
		if err != nil
		{
			general.Error(err);
			return;
		}

		fmt.Printf("dispatching new accepted connection\n");
		go doIt(connection, config);
	}
}

func Start()
{
	config := config.ParseConfigFile("proxy.conf");
	startServer(config);
	fmt.Printf("yay\n");
}
