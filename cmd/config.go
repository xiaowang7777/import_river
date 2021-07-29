package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"import_river/config"
)

var Config_CMD = &cobra.Command{
	Use:   "config",
	Short: "设置参数",
	Long: `
设置参数。
`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.New()
		if err != nil {
			logrus.Fatal(err)
		}
		filePath, err := cmd.Flags().GetString("file-path")
		if err != nil {
			logrus.Fatal(err)
		}
		c.FilePath = filePath
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			logrus.Fatal(err)
		}
		c.Host = host
		port, err := cmd.Flags().GetInt16("port")
		if err != nil {
			logrus.Fatal(err)
		}
		c.Port = port
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			logrus.Fatal(err)
		}
		c.Username = username
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			logrus.Fatal(err)
		}
		c.Password = password
		dbname, err := cmd.Flags().GetString("dbname")
		if err != nil {
			logrus.Fatal(err)
		}
		c.DatabaseName = dbname
		tbName, err := cmd.Flags().GetString("tableName")
		if err != nil {
			logrus.Fatal(err)
		}
		c.TableName = tbName
		if err := c.Write(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	Config_CMD.Flags().StringP("file-path", "f", "", "需要导入的文件路径")
	Config_CMD.Flags().StringP("host", "H", "localhost", "导入的数据库地址")
	Config_CMD.Flags().Int16P("port", "p", 5432, "数据库端口")
	Config_CMD.Flags().StringP("username", "u", "postgres", "数据库用户名")
	Config_CMD.Flags().StringP("password", "P", "postgres", "数据库用户密码")
	Config_CMD.Flags().StringP("dbname", "D", "", "数据库名")
	Config_CMD.Flags().StringP("tableName", "T", "", "表名")
}
