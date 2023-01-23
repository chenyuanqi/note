package controllers

import (
	"strconv"

	"mlj/app/model/demo"
	"mlj/pkg/console"
	"mlj/pkg/excel"

	"github.com/spf13/cobra"
)

var CmdUser = &cobra.Command{
	Use:   "user",
	Short: "Export user information.",
	Run:   runExport,
}

func runExport(cmd *cobra.Command, args []string) {
	console.Success("fetch user information")
	user := &demo.Users{}
	users := user.Find()
	// fmt.Printf("%+v\n", users)

	console.Success("transfer format data from users")
	data := make([]map[string]string, len(users))
	for _, item := range users {
		data = append(data, map[string]string{
			"id":         strconv.FormatInt(item.ID, 10),
			"name":       item.Name,
			"email":      item.Email,
			"created_at": item.CreatedAtDateTime(),
		})
	}
	// fmt.Printf("%+v", data)

	console.Success("export to excel")
	excel.ExportXlsx("用户信息", []string{"id", "name", "email", "created_at"}, []string{"ID", "名称", "邮件", "创建时间"}, data)
}
