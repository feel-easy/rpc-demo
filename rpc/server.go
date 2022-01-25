// golang写RPC程序，必须符合4个基本条件，不然RPC用不了

// 结构体字段首字母要大写，可以别人调用

// 函数名必须首字母大写

// 函数第一参数是接收参数，第二个参数是返回给客户端的参数，必须是指针类型

// 函数还必须有一个返回值error
package rpc

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"

	"github.com/spf13/cobra"
)

//    例题：golang实现RPC程序，实现求矩形面积和周长

type Params struct {
	Width, Height int
}

type Rect struct{}

// RPC服务端方法，求矩形面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	fmt.Println("矩形面积:", p)
	return nil
}

// 周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	fmt.Println("周长：", p)
	return nil
}

// 结构体，用于注册的
type Arith struct{}

// 声明参数结构体
type ArithRequest struct {
	A, B int
}

// 返回给客户端的结果
type ArithResponse struct {
	// 乘积
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}

// 乘法
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 商和余数
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	// 除
	res.Quo = req.A / req.B
	// 取模
	res.Rem = req.A % req.B
	return nil
}

// 主函数
func server() {
	// 1.注册服务
	rect := new(Arith)
	// 注册一个rect的服务
	rpc.Register(rect)
	// 2.服务处理绑定到http协议上
	rpc.HandleHTTP()
	// 3.监听服务
	fmt.Println("服务启动。。。")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panicln(err)
	}
}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "server Short",

	Run: func(cmd *cobra.Command, args []string) {
		server()
		fmt.Fprint(os.Stdout, args)
	},
}

func init() {
	// rootCmd.AddCommand(serverCmd)
}
