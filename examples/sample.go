package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/iazkaban/comfy4go/client"
	"github.com/iazkaban/comfy4go/client/websocket_message_model"
	"github.com/tidwall/sjson"
	"math/rand/v2"
	"os"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	opt := &client.ClientOption{
		Host:     "127.0.0.1",
		Port:     8818,
		Wg:       wg,
		ClientID: uuid.New().String(),
	}
	fmt.Println("====连接====ComfyUI服务:", opt.Host, ":", opt.Port)
	c, err := client.NewClient(opt)
	if err != nil {
		panic(err)
	}
	fmt.Println("====连接成功并绑定Websocket传递信息处理器====，ClientID:", opt.ClientID)

	once := &sync.Once{}
	c.RegisterMessageProcessor("crystools.monitor", func(message *client.BaseWebsocketMessage) error {
		once.Do(func() {
			fmt.Println("[Websocket传递信息处理器]===>定时CPU/内存/显存/显卡使用率信息收到，只展示一次")
		})
		return nil
	})

	c.RegisterMessageProcessor("status", func(message *client.BaseWebsocketMessage) error {
		rs := &websocket_message_model.Status{}
		err = json.Unmarshal(message.Data, rs)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println("[Websocket传递信息处理器]===>status", "当前队列数量：", rs.Status.ExecInfo.QueueRemaining)
		return nil
	})

	c.RegisterMessageProcessor("execution_start", func(message *client.BaseWebsocketMessage) error {
		rs := &websocket_message_model.ExecutionStart{}
		err = json.Unmarshal(message.Data, rs)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println("[Websocket传递信息处理器]===>execution_start添加执行队列成功，返回的PromptID为：", rs.PromptID)
		return nil
	})

	c.RegisterMessageProcessor("execution_cached", func(message *client.BaseWebsocketMessage) error {
		fmt.Println("[Websocket传递信息处理器]===>execution_cached信息返回，不是重要节点，不做解析")
		return nil
	})

	c.RegisterMessageProcessor("executing", func(message *client.BaseWebsocketMessage) error {
		rs := &websocket_message_model.Executing{}
		err = json.Unmarshal(message.Data, rs)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println("[Websocket传递信息处理器]executing开始进入", rs.PromptID, "===>节点", rs.Node, "执行")
		return nil
	})

	c.RegisterMessageProcessor("progress", func(message *client.BaseWebsocketMessage) error {
		rs := &websocket_message_model.Progress{}
		err = json.Unmarshal(message.Data, rs)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		if rs.Value%10 == 0 {
			fmt.Println("[Websocket传递信息处理器]progress处理进度简化展示：", rs.Value, "/", rs.Max)
		}
		if rs.Value == rs.Max {
			fmt.Println("[Websocket传递信息处理器]progress处理进度简化展示：处理完成")
		}
		return nil
	})

	c.RegisterMessageProcessor("execution_success", func(message *client.BaseWebsocketMessage) error {
		rs := &websocket_message_model.ExecutionSuccess{}
		err = json.Unmarshal(message.Data, rs)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		fmt.Println("[Websocket传递信息处理器]execution_success处理成功通知：", rs.PromptID)
		return nil
	})

	c.RegisterMessageProcessor("executed", func(message *client.BaseWebsocketMessage) error {
		rs := &websocket_message_model.Executed{}
		err = json.Unmarshal(message.Data, rs)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		fmt.Println("[Websocket传递信息处理器]executed输出结果通知,生成图片数量：", len(rs.Output.Images))
		for _, v := range rs.Output.Images {
			_, body, err := c.DownloadImage(v)
			if err != nil {
				fmt.Println("Download Error:", err)
				return nil
			}
			f, err := os.Create("./" + v.Filename)
			if err != nil {
				fmt.Println("Save File Error:", err)
				continue
			}
			_, err = f.Write(body)
			if err != nil {
				fmt.Println("Error:", err)
				return nil
			}
			err = f.Close()
			if err != nil {
				fmt.Println("Error:", err)
				return nil
			}
			fmt.Println("图片类型:", v.Type, "，保存图片", v.Filename, "成功")
		}
		return nil
	})

	userList, err := c.GetUsers()
	if err != nil {
		panic(err)
	}
	fmt.Println("====开始获取用户信息====")
	c.SelectUser(userList.Users[0].UserID)
	fmt.Println("\t用户数量：", len(userList.Users), "选择用户：", userList.Users[0].UserName)

	fmt.Println("====获取工作流信息====")
	wf, err := c.WorkflowList()
	if err != nil {
		panic(err)
	}
	fmt.Println("\t获取到工作流数量：", len(wf))

	fmt.Println("====获取插件信息====")
	exList, err := c.Extensions()
	if err != nil {
		panic(err)
	}
	fmt.Println("\t获取到插件数量：", len(exList))

	fmt.Println("====上传一个图片====")
	filename := "C:/Users/隋龙飞/Developer/comfy4go/examples/ComfyUI_temp_gctxs_00001_.png"

	fmt.Println("\t第一种提供文件上传方式：")
	image, err := c.UploadImageFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("\t\t上传成功，文件名:", image.Name)

	fmt.Println("\t第二种提供文件二进制内容方式：")
	body, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	image1, err := c.UploadImage("第二种方式图片.png", body)
	if err != nil {
		panic(err)
	}
	fmt.Println("\t\t上传成功，文件名:", image1.Name)

	fmt.Println("=====选择一个工作流====")

	fmt.Println("\t获取【美女姿态.json】工作流")

	w, err := c.Workflow("美女姿态.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("\t\t获取到的节点数量:", len(w.Nodes))

	rsw, _ := json.Marshal(w)

	fBody, _ := os.ReadFile("./aaa.json")

	//生成随机Seed
	fBody, _ = sjson.SetBytes(fBody, "1.inputs.seed", rand.Int64())

	req := &client.PromptRequest{}
	req.ClientID = c.ClientID
	req.ExtraData.ExtraPngInfo.WorkFlow = rsw
	req.Prompt = fBody

	for i := 0; i < 3; i++ {
		rs, err := c.Prompt(req)
		if err != nil {
			panic(err)
		}

		fmt.Println("====使用工作流添加执行====执行生成的PromptID:", rs.PromptID)
		time.Sleep(time.Second * 10)
	}

	wg.Wait()
}
