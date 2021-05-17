package db

import (
        "os"
)
import "io/ioutil"


/*
 * 描述: 读取配置文件
 *
 *	name: 路径与文件名
 *
 ****************************************************************/
func ReadConfig( name string)( []byte, error ) {
        jsonFile, err := os.Open( name )
        if err != nil {
                panic("打开文件错误，请查看:" + name )
        }
        defer jsonFile.Close()
        return ioutil.ReadAll(jsonFile)
}


