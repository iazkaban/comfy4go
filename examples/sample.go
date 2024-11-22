package main

import (
	"fmt"
	"github.com/iazkaban/comfy4go/client"
)

func main() {
	c := client.NewClient("127.0.0.1", 8188)

	userList, err := c.GetUsers()
	if err != nil {
		panic(err)
	}
	fmt.Println("===================用户信息===================")
	fmt.Println(userList.Storage)

	for _, v := range userList.Users {
		fmt.Println(v.UserID, "==>", v.UserName)
		c.SelectUser(v.UserID)
	}

	fmt.Println("===================工作流信息===================")
	wf, err := c.WorkflowList()
	if err != nil {
		panic(err)
	}

	for _, v := range wf {
		fmt.Println("Path:", v.Path, "Size:", v.Size, "Modified:", v.Modified)
	}

	fmt.Println("===================插件信息===================")
	exList, err := c.Extensions()
	if err != nil {
		panic(err)
	}

	for _, v := range exList {
		fmt.Println(v)
	}

	fmt.Println("===================第一个工作流===================")
	w, err := c.Workflow(wf[0].Path)
	if err != nil {
		panic(err)
	}
	fmt.Println(w.Version)
	for _, node := range w.Nodes {
		fmt.Println(node.ID, node.Type)
	}
}
