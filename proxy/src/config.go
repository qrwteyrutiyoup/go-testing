package config

import (
	"os";
	"bufio";
	"strings";
	"./general";
)

func parseLine(line string) ([]string)
{
	line = strings.TrimSpace(line);

	if len(line) > 0
	{
		if line[len(line) - 1] == '\n'
		{
			line = line[0:(len(line) - 1)];	
		}

		ret := strings.Split(line, " ", 2);
		if len(ret) == 2 && len(ret[0]) > 0 && len(ret[1]) > 0
		{
			return ret;
		}
	}

	return nil;
}


func ParseConfigFile(filename string) (map[string]string)
{
	config := make(map[string]string);

	fd, fderr := os.Open(filename, os.O_RDONLY, 0444);

	if fderr != nil
	{
		general.Error(fderr);
		os.Exit(1);
	}
	defer fd.Close();

	reader := bufio.NewReader(fd);

	for 
	{
		line, rerr := reader.ReadString('\n');
		if rerr != nil 
		{
			if rerr != os.EOF
			{
				general.Error(fderr);
			}
			break;
		}

		parsed := parseLine(line);
		if parsed != nil
		{
			config[parsed[0]] = parsed[1];
		}
	}

	return config;
}

