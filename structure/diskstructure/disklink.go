package diskstructure

import (
    "bytes"
    "encoding/binary"
    "errors"
    "io"
    "os"
    "strconv"
    "strings"
    "sync"
)

const DBFileName = "link.db"

// do not change this page size if there already have a db file
const PageSize = 4 * 1024

var (
    NotFoundErr        = errors.New("not found")
    AlReadyExistsErr   = errors.New("already exists")
    OffsetOutOfFileErr = errors.New("offset is out of file")
    UnKnowErr          = errors.New("unknow error")
)

type Link interface {
    Get(uint32) (string, error)
    Add(uint32, string) error
    Set(uint32, string) error
    Del(uint32) error
    Keys() ([]uint32, error)
    Print()
    Clear() error
    Invert() error
    Close() error
}

type DLink struct {
    Head   *Head
    DBFile os.File
    mutex  sync.RWMutex
}

type Head struct {
    Length uint32
    Next   uint32
}

type Node struct {
    Key  uint32
    Data string
    Self uint32
    Next uint32
}

type metadata struct {
    DataLen uint32
}

func newNode(key uint32, data string, self uint32, next uint32) *Node {
    return &Node{
        Key:  key,
        Data: data,
        Self: self,
        Next: next,
    }
}

// sequence: metadata, key, data, selef, next
func (n *Node) serialize() (*bytes.Buffer, error) {
    data := []byte(n.Data)
    m := &metadata{
        DataLen: uint32(len(data)),
    }
    buf := bytes.NewBuffer([]byte{})
    if err := binary.Write(buf, binary.LittleEndian, m); err != nil {
        return nil, err
    }
    if err := binary.Write(buf, binary.LittleEndian, n.Key); err != nil {
        return nil, err
    }
    if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
        return nil, err
    }
    if err := binary.Write(buf, binary.LittleEndian, n.Self); err != nil {
        return nil, err
    }
    if err := binary.Write(buf, binary.LittleEndian, n.Next); err != nil {
        return nil, err
    }
    return buf, nil
}

// sequence: metadata, key, data, selef, next
func (n *Node) deSerialize(buf *bytes.Buffer) error {
    m := &metadata{}
    if err := binary.Read(buf, binary.LittleEndian, m); err != nil {
        return err
    }
    if err := binary.Read(buf, binary.LittleEndian, &n.Key); err != nil {
        return err
    }
    dataBytes := make([]byte, int(m.DataLen))
    if err := binary.Read(buf, binary.LittleEndian, dataBytes); err != nil {
        return err
    }
    n.Data = string(dataBytes)
    if err := binary.Read(buf, binary.LittleEndian, &n.Self); err != nil {
        return err
    }
    if err := binary.Read(buf, binary.LittleEndian, &n.Next); err != nil {
        return err
    }
    return nil
}

func (l *DLink) getHead() (*Head, error) {
    pageBytes := make([]byte, PageSize)
    if _, err := l.DBFile.ReadAt(pageBytes, 0); err != nil {
        return nil, err
    }
    head := &Head{}
    buf := bytes.NewBuffer(pageBytes)
    if err := binary.Read(buf, binary.LittleEndian, head); err != nil {
        return nil, err
    }
    return head, nil
}

func (l *DLink) flushHead() error {
    l.mutex.Lock()
    defer l.mutex.Unlock()

    buf := bytes.NewBuffer([]byte{})
    head := l.Head
    if err := binary.Write(buf, binary.LittleEndian, head); err != nil {
        return err
    }
    b := buf.Bytes()
    for len(b) < PageSize {
        b = append(b, 0)
    }
    if _, err := l.DBFile.WriteAt(b, 0); err != nil {
        return err
    }
    return nil
}

func (l *DLink) flushNode(node *Node) error {
    l.mutex.Lock()
    defer l.mutex.Unlock()

    buf, err := node.serialize()
    if err != nil {
        return err
    }
    b := buf.Bytes()
    for len(b) < PageSize {
        b = append(b, 0)
    }
    if _, err := l.DBFile.WriteAt(b, int64(node.Self)); err != nil {
        return err
    }
    return nil
}

func (l *DLink) getNodeByOffset(offset uint32) (*Node, error) {
    l.mutex.RLock()
    defer l.mutex.RUnlock()

    endOffset, _ := l.DBFile.Seek(0, io.SeekEnd)
    if offset > uint32(endOffset) {
        return nil, OffsetOutOfFileErr
    }
    b := make([]byte, PageSize)
    if _, err := l.DBFile.ReadAt(b, int64(offset)); err != nil {
        return nil, err
    }
    buf := bytes.NewBuffer(b)
    n := &Node{}
    if err := n.deSerialize(buf); err != nil {
        return nil, err
    }
    return n, nil
}

func (l *DLink) getNodeByKey(key uint32) (*Node, error) {
    l.mutex.RLock()
    defer l.mutex.RUnlock()

    if l.Head.Next == 0 {
        return nil, NotFoundErr
    }
    var (
        node = &Node{}
        next = l.Head.Next
        err  error
    )
    for node.Key != key {
        if next == 0 {
            return nil, NotFoundErr
        }
        node, err = l.getNodeByOffset(next)
        if err != nil {
            return nil, err
        }
        next = node.Next
    }
    return node, nil
}

func (l *DLink) getEndNode() (*Node, error) {
    l.mutex.RLock()
    defer l.mutex.RUnlock()

    if l.Head.Next == 0 {
        return nil, NotFoundErr
    }
    var (
        next = l.Head.Next
        node = &Node{}
        err  error
    )
    for next != 0 {
        node, err = l.getNodeByOffset(next)
        if err != nil {
            return nil, err
        }
        next = node.Next
    }
    return node, nil
}

func (l *DLink) Get(key uint32) (string, error) {
    node, err := l.getNodeByKey(key)
    if err != nil {
        return "", err
    }
    return node.Data, nil
}

func (l *DLink) Add(key uint32, data string) error {
    if _, err := l.getNodeByKey(key); err == nil {
        return AlReadyExistsErr
    }
    if l.Head.Next == 0 {
        self := uint32(PageSize)
        node := newNode(key, data, self, 0)
        l.flushNode(node)
        l.Head.Next = PageSize
        l.Head.Length++
        l.flushHead()
        return nil
    }
    node, err := l.getEndNode()
    if err != nil {
        return err
    }
    // is not alway append, eg: after invert
    // newNodeOffset := node.Self + PageSize
    endOffset, err := l.DBFile.Seek(0, io.SeekEnd)
    if err != nil {
        return err
    }
    l.DBFile.Seek(0, 0)
    newNodeOffset := uint32(endOffset)
    newNode := newNode(key, data, newNodeOffset, 0)
    if err := l.flushNode(newNode); err != nil {
        return err
    }
    node.Next = newNodeOffset
    if err := l.flushNode(node); err != nil {
        return err
    }
    l.Head.Length++
    l.flushHead()
    return nil
}

func (l *DLink) Set(key uint32, data string) error {
    node, err := l.getNodeByKey(key)
    if err != nil {
        return err
    }
    node.Data = data
    l.flushNode(node)
    return nil
}

func (l *DLink) Del(key uint32) error {
    headNextNode, _ := l.getNodeByOffset(l.Head.Next)
    node, err := l.getNodeByKey(key)
    if err != nil {
        return err
    }
    if headNextNode.Key == node.Key {
        l.Head.Next = node.Next
        l.Head.Length--
        l.flushHead()
        return nil
    }
    var (
        preNode  = &Node{}
        findNode = &Node{}
        next     = l.Head.Next
    )
    for next != 0 {
        findNode, err = l.getNodeByOffset(next)
        if err != nil {
            return err
        }
        if findNode.Key == key {
            break
        }
        preNode = findNode
        next = findNode.Next
    }
    // already include the end of link (when next is 0)
    preNode.Next = node.Next
    l.flushNode(preNode)
    l.Head.Length--
    l.flushHead()
    return nil
}

func (l *DLink) Invert() error {
    var (
        preNode, _ = l.getNodeByOffset(l.Head.Next)
        next       = preNode.Next
        node       = &Node{}
    )
    for next != 0 {
        // head -> preNode -> node -> nextNode ...
        node, _ = l.getNodeByOffset(next)
        next = node.Next
        node.Next = preNode.Self
        l.flushNode(node)
        preNode = node
    }
    // this first node means the latest node
    firstNode, _ := l.getNodeByOffset(l.Head.Next)
    firstNode.Next = 0
    l.flushNode(firstNode)
    l.Head.Next = node.Self
    l.flushHead()
    return nil
}

func (l *DLink) Keys() (list []uint32, err error) {
    next := l.Head.Next
    node := &Node{}
    for next != 0 {
        node, err = l.getNodeByOffset(next)
        if err != nil {
            return nil, err
        }
        list = append(list, node.Key)
        next = node.Next
    }
    return list, nil
}

func (l *DLink) Print() {
    var build strings.Builder
    build.WriteString("{ ")
    build.WriteString("length: " + strconv.Itoa(int(l.Head.Length)))
    build.WriteString(", ")
    build.WriteString("next: " + strconv.Itoa(int(l.Head.Next)))
    build.WriteString(" }")
    list, err := l.Keys()
    if err != nil {
        return
    }
    for _, k := range list {
        d, err := l.getNodeByKey(k)
        if err != nil {
            return
        }
        build.WriteString(" -> ")
        build.WriteString("{ ")
        build.WriteString("key: " + strconv.Itoa(int(k)))
        build.WriteString(", ")
        build.WriteString("data: " + "\"" + d.Data + "\"")
        build.WriteString(", ")
        build.WriteString("self: " + strconv.Itoa(int(d.Self)))
        build.WriteString(", ")
        build.WriteString("next: " + strconv.Itoa(int(d.Next)))
        build.WriteString(" }")
    }
    print(build.String(), "\n")
}

func (l *DLink) Clear() error {
    l.mutex.Lock()
    defer l.mutex.Unlock()

    if err := l.Close(); err != nil {
        return err
    }
    if err := os.Remove(DBFileName); err != nil {
        return err
    }
    l = NewLink().(*DLink)
    return nil
}

func (l *DLink) Close() error {
    l.mutex.Lock()
    defer l.mutex.Unlock()

    if err := l.DBFile.Sync(); err != nil {
        return err
    }
    if err := l.DBFile.Close(); err != nil {
        return err
    }
    return nil
}

func NewLink() Link {
    l := &DLink{
        mutex: sync.RWMutex{},
    }
    file, err := os.OpenFile(DBFileName, os.O_RDWR|os.O_CREATE, 0600)
    if err != nil {
        panic(err)
    }
    l.DBFile = *file
    h, err := l.getHead()
    if err != nil {
        l.Head = &Head{}
        if err := l.flushHead(); err != nil {
            panic(err)
        }
    } else {
        l.Head = h
    }
    return l
}
