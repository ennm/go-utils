package generate

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "io"
    "time"
    "github.com/ennm/go-utils/util"
)

const DIR = "./handlers"

type RpcArr struct {
    FuncName string `json:"func_name"`
    Req      string `json:"req"`
    Rsp      string `json:"rsp"`
}

func Do(fileName string) error {

    file, err := os.Open("./proto/" + fileName + "/" + fileName + ".proto")

    if err != nil {

        fmt.Println("file open err", err.Error())

        return err
    }

    defer file.Close()

    _, err = file.Stat()

    if err != nil {

        fmt.Println("stat err", err.Error())
    }

    buf := bufio.NewReader(file)

    var funcName, req, rsp string

    rpc := make([]*RpcArr, 0)

    for {
        line, err := buf.ReadString('\n')

        line = strings.TrimSpace(line)

        each := strings.Split(line, " ")

        if each[0] == "rpc" {

            oneArr := strings.Split(each[1], "(")

            funcName = oneArr[0]

            req = util.Trim(oneArr[1])

            rsp = util.Trim(each[3])

            item := new(RpcArr)

            item.FuncName = funcName
            item.Req = req
            item.Rsp = rsp

            rpc = append(rpc, item)
            item = nil
        }

        if err != nil {

            if err == io.EOF {

                fmt.Println(time.Now().Format("2006-01-02 15:04:05"), ">>>>>>读取文件内容成功")

                break
            } else {

                fmt.Println("Read file error!", err)

                return err
            }
        }
    }

    Write(fileName, rpc)

    return nil
}

func Write(fileName string, rpc []*RpcArr) error {

    fileFullName := DIR + "/" + fileName + ".go"

    file, err := os.OpenFile(fileFullName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)

    if err != nil {

        fmt.Println("open file err", err.Error())
    }

    defer file.Close()

    dir, _ := os.Getwd()

    a := strings.Split(dir, "src/")

    pack := "package handlers\n\n"

    io.WriteString(file, pack)

    upperName := util.UcFirst(fileName)

    impo := "import (\n" +
        "    " + "\"context\"\n" +
        "    \"code.aliyun.com/retail/util\"\n" +
        "    " + fmt.Sprintf(fileName) + " \"" + a[1] + "/proto/" + fileName + "\"\n" +
        ")\n\n"

    io.WriteString(file, impo)

    stru := "type " + upperName + " struct{}\n\n"

    io.WriteString(file, stru)

    alias := strings.ToLower(fileName[:1])

    for _, v := range rpc {

        line := "func (" + alias + " *" + upperName + ") " + v.FuncName + "(ctx context.Context, req *" + fileName + "." + v.Req + ", rsp *" + fileName + "." + v.Rsp + ") error {\n\n" +
            "    defer util.HandleError(rsp)()\n\n" +
            "    return nil\n" +
            "}\n\n"

        io.WriteString(file, line)
    }

    fmt.Println(time.Now().Format("2006-01-02 15:04:05"), ">>>>>>生成文件成功")

    return nil
}