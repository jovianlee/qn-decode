package cmd

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/luanxuechao/qn-decode/util"
	"github.com/nu11ptr/cmpb"
	"github.com/spf13/cobra"
)

// FilePath file path
var FilePath string
var dirname string

func init() {
	rootCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().StringVarP(&FilePath, "FILE", "f", "", "decode file path")
	decodeCmd.Flags().StringVarP(&dirname, "DIR", "d", "", "decode dir path")
}

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decode music file",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if FilePath == "" && dirname == "" {
			return errors.New("require a file path")
		}
		if FilePath != "" {
			_, err := os.Lstat(FilePath)
			if os.IsNotExist(err) {
				return errors.New("file not found")
			}
		}
		if dirname != "" {
			_, err := os.Lstat(dirname)
			if os.IsNotExist(err) {
				return errors.New("dir not found")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		decode()
	},
}

func decodeFile(filePath string, p *cmpb.Progress) error {
	var strIndex int = strings.LastIndex(filePath, ".")
	var fileFormat = filePath[strIndex+1:]
	_, fileName := filepath.Split(filePath)

	switch fileFormat {
	case "qmcflac":
		return util.DecodeQmcFlac(filePath, fileName, p)
	case "qmc0", "qmc3":
		return util.DecodeQmc0OrQmc3(filePath, fileName, p)
	case "ncm":
		return util.Dump(filePath, fileName, p)
	default:
		return errors.New("the file not support")
	}
}
func decodeDir(p *cmpb.Progress) error {
	log.Println("Decode Dir:", dirname)
	s, err := os.Stat(dirname)
	if err != nil {
		return errors.New("the dir not found")
	}
	if !s.IsDir() {
		return errors.New("the dir is not a folder")
	}
	rd, err := ioutil.ReadDir(dirname)
	if err != nil {
		return errors.New(err.Error())
	}
	var FilePathList []string = make([]string, 0)
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		}
		name := fi.Name()
		var strIndex int = strings.LastIndex(name, ".")
		var fileFormat = name[strIndex+1:]
		if fileFormat != "qmcflac" && fileFormat != "qmc0" && fileFormat != "qmc3" && fileFormat != "ncm" {
			continue
		}
		FilePathList = append(FilePathList, dirname+"/"+fi.Name())
	}

	for _, filePath := range FilePathList {
		err := decodeFile(filePath, p)
		if err != nil {
			log.Println("Decode file error:", err.Error())
			continue
		}
	}
	return nil
}
func decode() error {
	p := cmpb.New()
	colors := new(cmpb.BarColors)

	colors.Post, colors.KeyDiv, colors.LBracket, colors.RBracket =
		color.HiCyanString, color.HiCyanString, color.HiCyanString, color.HiCyanString

	colors.Key = color.HiBlueString
	colors.Msg, colors.Empty = color.HiYellowString, color.HiYellowString
	colors.Full = color.HiGreenString
	colors.Curr = color.GreenString
	colors.PreBar, colors.PostBar = color.HiMagentaString, color.HiMagentaString
	p.SetColors(colors)
	p.Start()
	p.Wait()
	if FilePath != "" {
		return decodeFile(FilePath, p)
	}
	return decodeDir(p)
}
