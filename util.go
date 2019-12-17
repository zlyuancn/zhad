/*
-------------------------------------------------
   Author :       zlyuan
   date：         2019/12/17
   Description :
-------------------------------------------------
*/

package zhad

// uint32转为bytes, 从右边开始写入
func Uint32ToBytes(v uint32) []byte {
    return []byte{
        byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v),
    }
}

// bytes转为uint32, 从右边开始读取
func BytesToUint32(b []byte) uint32 {
    return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

// uint16转为bytes, 从右边开始写入
func Uint16ToBytes(v uint16) []byte {
    return []byte{
        byte(v >> 8), byte(v),
    }
}

// bytes转为uint16, 从右边开始读取
func BytesToUint16(b []byte) uint16 {
    return uint16(b[1]) | uint16(b[0])<<8
}
