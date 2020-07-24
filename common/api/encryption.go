package api

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)
/*
	一共三步加密
	1.通过sha1加密  => res1
	2.将res1 进行sha1 加密，并且拼接 res1  => res2
	3.再将res2 进行md5 加密
*/
// 加密字符串
func Encryption(str string)string{

	/*
	// go 代码

	hash := sha1.New()
	hash.Write([]byte(str))
	res1 := hash.Sum(nil)

	enMd5 := md5.New()
	enMd5.Write(res1)
	res2 := enMd5.Sum(nil)

	return hex.EncodeToString(res2)
	*/



	// 以下代码为了同步nodejs的算法
	// 第一步
	hash := sha1.New()
	hash.Write([]byte(str))
	// node是先将数据转成 16进制 hex 再进行加密，所以我要先将字节数据转成hex 再转成hex的字节数据加密
	res1 := hex.EncodeToString(hash.Sum(nil))


	// 第三步
	enMd5 := md5.New()
	enMd5.Write([]byte(res1))
	res2 := enMd5.Sum(nil)
	return hex.EncodeToString(res2)
}