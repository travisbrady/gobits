/* Connect to redis just to play with Go a little bit
 * You'd need to have Redis running on localhost on port 6379
 * You'd also need to have a computer that can run programs
 * You should probably also get food and shelter too.
 * Actually, just check this out:
 * http://en.wikipedia.org/wiki/Maslow%27s_hierarchy_of_needs
*/
package main

import (
    "os";
    "net";
    "fmt";
    "strings";
)

func sendThenReceive(conn net.Conn, cmd string) []byte {
    _,err := conn.Write(strings.Bytes(cmd));
    if err != nil {
        fmt.Println("uh oh");
    }
    buf := make([]byte, 1000);
    conn.Read(buf);
    return buf;
}

func main() {
	conn, err := net.Dial("tcp", "", "127.0.0.1:6379");
	if err != nil {
        fmt.Println("Connect Error");
        os.Exit(1);
	}
    pingResult := sendThenReceive(conn, "PING\r\n");
    fmt.Printf("Ping result: %s", pingResult);

    setResult := sendThenReceive(conn, "SET key 3\r\nval\r\n");
    fmt.Printf("Set Result: %s", setResult);

    getResult := sendThenReceive(conn, "GET key\r\n");
    fmt.Printf("Get Result: %s\n", getResult);
}
