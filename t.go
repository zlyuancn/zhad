/*
-------------------------------------------------
   Author :       zlyuan
   date：         2019/12/17
   Description :
-------------------------------------------------
*/

package zhad

import (
    "errors"
    "fmt"
)

type HAD struct {
    header []byte
    bodys  [][]byte

    pre_size int
    out_size int
}

// 从data中加载
func New(data []byte) (*HAD, error) {
    dsize := len(data)
    m := new(HAD)
    offset := 0

    if dsize < offset+2 {
        return nil, errors.New("header描述数据不完整")
    }
    hsize := int(BytesToUint16(data[offset : offset+2]))
    offset += 2

    if dsize < offset+hsize {
        return nil, errors.New("header数据不完整")
    }
    m.header = make([]byte, hsize)
    copy(m.header, data[offset:offset+hsize])
    offset += hsize

    if dsize < offset+1 {
        return nil, errors.New("没有body数量")
    }
    bsize := int(data[offset])
    offset += 1

    out_size := offset + bsize*4
    if dsize < out_size {
        return nil, errors.New("没有body描述数据")
    }

    bsizes := make([]int, bsize)
    for i := 0; i < bsize; i++ {
        size := int(BytesToUint32(data[offset : offset+4]))
        bsizes[i] = size
        offset += 4
        out_size += size
    }

    if dsize < out_size {
        return nil, errors.New("body数据不完整")
    }

    for i := 0; i < bsize; i++ {
        size := bsizes[i]
        body := make([]byte, size)
        copy(body, data[offset:offset+size])
        offset += size
        m.bodys = append(m.bodys, body)
    }

    m.pre_size = m.prefixSize()
    m.out_size = m.calculateSize()
    return m, nil
}

// 设置头
func (m *HAD) SetHeader(h []byte) error {
    if len(h) > HeaderMaxSize {
        return fmt.Errorf("header最大大小超出 %d", HeaderMaxSize)
    }
    m.header = h
    m.pre_size = m.prefixSize()
    m.out_size = m.calculateSize()
    return nil
}

// 设置body
func (m *HAD) SetBody(bodys ...[]byte) error {
    if len(bodys) > BodyMaxCount {
        return fmt.Errorf("body最大数量超出 %d", BodyMaxCount)
    }

    for i, body := range bodys {
        if len(body) > BodyMaxSize {
            return fmt.Errorf("bodys[%d]超出了最大大小 %d", i, BodyMaxSize)
        }
    }
    m.bodys = bodys
    m.pre_size = m.prefixSize()
    m.out_size = m.calculateSize()
    return nil
}

// 添加body
func (m *HAD) AddBody(bodys ...[]byte) error {
    if len(m.bodys)+len(bodys) > BodyMaxCount {
        return fmt.Errorf("添加后body最大数量会超出 %d, 操作已拦截", BodyMaxCount)
    }

    for i, body := range bodys {
        if len(body) > BodyMaxSize {
            return fmt.Errorf("bodys[%d]超出了最大大小 %d", i, BodyMaxSize)
        }
    }

    m.bodys = append(m.bodys, bodys...)
    m.pre_size = m.prefixSize()
    m.out_size = m.calculateSize()
    return nil
}

// 计算前置数据大小
func (m *HAD) prefixSize() int {
    // 头大小 头数据 body数量 每个body大小 body数据
    return 2 + len(m.header) + 1 + 4*len(m.bodys)
}

// 计算大小
func (m *HAD) calculateSize() int {
    size := m.prefixSize()
    for _, body := range m.bodys {
        size += len(body)
    }
    return size
}

// 获取头
func (m *HAD) Header() []byte {
    return m.header
}

// 获取body
func (m *HAD) Bodys() [][]byte {
    return m.bodys
}

// 获取输出大小
func (m *HAD) Size() int {
    return m.out_size
}

// 输出数据
func (m *HAD) ToData() []byte {
    bs := make([]byte, m.out_size)
    offset := 0

    // 头大小
    copy(bs[offset:], Uint16ToBytes(uint16(len(m.header))))
    offset += 2

    // 头数据
    copy(bs[offset:], m.header)
    offset += len(m.header)

    // body数量
    bs[offset] = byte(len(m.bodys))
    offset += 1

    // 每个body大小
    for _, body := range m.bodys {
        copy(bs[offset:], Uint32ToBytes(uint32(len(body))))
        offset += 4
    }

    // body数据
    for _, body := range m.bodys {
        copy(bs[offset:], body)
        offset += len(body)
    }

    return bs
}
