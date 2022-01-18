package cmd

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/spf13/cobra"
)

// 主函数
func client() {
	// 1.连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	// params := Params{50, 100}
	// // 2.调用方法
	// // 面积
	// ret := 0
	// err2 := conn.Call("Rect.Area", params, &ret)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	// fmt.Println("面积：", ret)
	// // 周长
	// err3 := conn.Call("Rect.Perimeter", params, &ret)
	// if err3 != nil {
	// 	log.Fatal(err3)
	// }
	// fmt.Println("周长：", ret)

	req := ArithRequest{9, 2}
	var res ArithResponse
	err2 := conn.Call("Arith.Multiply", req, &res)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)
	err3 := conn.Call("Arith.Divide", req, &res)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Printf("%d / %d 商 %d，余数 = %d\n", req.A, req.B, res.Quo, res.Rem)
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client Short",
	Run: func(cmd *cobra.Command, args []string) {
		client()
		fmt.Fprint(os.Stdout, args)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
