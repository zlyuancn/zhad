/*
-------------------------------------------------
   Author :       zlyuan
   date：         2019/12/17
   Description :
-------------------------------------------------
*/

package zhad

import "testing"

func EqBytes(a []byte, b []byte) bool {
    if len(a) != len(b) {
        return false
    }

    for i, v := range a {
        if v != b[i] {
            return false
        }
    }

    return true
}

func Test_Header(t *testing.T) {
    h := new(HAD)
    header := []byte("小明")
    if err := h.SetHeader(header); err != nil {
        t.Fatal(err)
    }
    if !EqBytes(h.Header(), header) {
        t.Fatalf("设置的header和获取的header不一致: %s", string(h.Header()))
    }

    data := h.ToData()
    h2, err := New(data)
    if err != nil {
        t.Fatalf("逆向错误: %s", err.Error())
    }

    if !EqBytes(h2.Header(), h.Header()) {
        t.Fatalf("逆向后header发送变化: %s", string(h2.Header()))
    }

    if h2.Size() != h.Size() {
        t.Fatalf("逆向后输出大小 %d 和原始大小 %d 不相等", h2.Size(), h.Size())
    }

    if !EqBytes(h2.ToData(), data) {
        t.Fatalf("逆向后构建的had输出数据和原数据不一致")
    }
}

func Test_Body(t *testing.T) {
    h := new(HAD)
    body := []byte("你好啊")
    body2 := []byte("在做什么")
    if err := h.AddBody(body, body2); err != nil {
        t.Fatal(err)
    }
    if !EqBytes(h.Bodys()[0], body) {
        t.Fatalf("设置的body和获取的body不一致: %s", string(h.Bodys()[0]))
    }
    if !EqBytes(h.Bodys()[1], body2) {
        t.Fatalf("设置的body和获取的body不一致: %s", string(h.Bodys()[1]))
    }

    data := h.ToData()
    h2, err := New(data)
    if err != nil {
        t.Fatalf("逆向错误: %s", err.Error())
    }

    if !EqBytes(h.Bodys()[0], h2.Bodys()[0]) {
        t.Fatalf("逆向后body发送变化: %s", string(h2.Bodys()[0]))
    }
    if !EqBytes(h.Bodys()[1], h2.Bodys()[1]) {
        t.Fatalf("逆向后body发送变化: %s", string(h2.Bodys()[1]))
    }

    if h2.Size() != h.Size() {
        t.Fatalf("逆向后输出大小 %d 和原始大小 %d 不相等", h2.Size(), h.Size())
    }

    if !EqBytes(h2.ToData(), data) {
        t.Fatalf("逆向后构建的had输出数据和原数据不一致")
    }
}
