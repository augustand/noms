// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeStructType("Tree",
			[]types.Field{
				types.Field{"children", types.MakeCompoundType(types.ListKind, types.MakeType(ref.Ref{}, 0)), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	types.RegisterPackage(&p)
}

// Tree

type Tree struct {
	_children ListOfTree

	ref *ref.Ref
}

func NewTree() Tree {
	return Tree{
		_children: NewListOfTree(),

		ref: &ref.Ref{},
	}
}

type TreeDef struct {
	Children ListOfTreeDef
}

func (def TreeDef) New() Tree {
	return Tree{
		_children: def.Children.New(),
		ref:       &ref.Ref{},
	}
}

func (s Tree) Def() (d TreeDef) {
	d.Children = s._children.Def()
	return
}

var __typeForTree types.Type

func (m Tree) Type() types.Type {
	return __typeForTree
}

func init() {
	__typeForTree = types.MakeType(ref.Parse("sha1-1166998faee80d4112e7b74cf44ad2647c196566"), 0)
	types.RegisterStruct(__typeForTree, builderForTree, readerForTree)
}

func builderForTree(values []types.Value) types.Value {
	i := 0
	s := Tree{ref: &ref.Ref{}}
	s._children = values[i].(ListOfTree)
	i++
	return s
}

func readerForTree(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Tree)
	values = append(values, s._children)
	return values
}

func (s Tree) Equals(other types.Value) bool {
	return other != nil && __typeForTree.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Tree) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Tree) Chunks() (chunks []types.RefBase) {
	chunks = append(chunks, __typeForTree.Chunks()...)
	chunks = append(chunks, s._children.Chunks()...)
	return
}

func (s Tree) ChildValues() (ret []types.Value) {
	ret = append(ret, s._children)
	return
}

func (s Tree) Children() ListOfTree {
	return s._children
}

func (s Tree) SetChildren(val ListOfTree) Tree {
	s._children = val
	s.ref = &ref.Ref{}
	return s
}

// ListOfTree

type ListOfTree struct {
	l   types.List
	ref *ref.Ref
}

func NewListOfTree() ListOfTree {
	return ListOfTree{types.NewTypedList(__typeForListOfTree), &ref.Ref{}}
}

type ListOfTreeDef []TreeDef

func (def ListOfTreeDef) New() ListOfTree {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New()
	}
	return ListOfTree{types.NewTypedList(__typeForListOfTree, l...), &ref.Ref{}}
}

func (l ListOfTree) Def() ListOfTreeDef {
	d := make([]TreeDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(Tree).Def()
	}
	return d
}

func (l ListOfTree) Equals(other types.Value) bool {
	return other != nil && __typeForListOfTree.Equals(other.Type()) && l.Ref() == other.Ref()
}

func (l ListOfTree) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfTree) Chunks() (chunks []types.RefBase) {
	chunks = append(chunks, l.Type().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

func (l ListOfTree) ChildValues() []types.Value {
	return append([]types.Value{}, l.l.ChildValues()...)
}

// A Noms Value that describes ListOfTree.
var __typeForListOfTree types.Type

func (m ListOfTree) Type() types.Type {
	return __typeForListOfTree
}

func init() {
	__typeForListOfTree = types.MakeCompoundType(types.ListKind, types.MakeType(ref.Parse("sha1-1166998faee80d4112e7b74cf44ad2647c196566"), 0))
	types.RegisterValue(__typeForListOfTree, builderForListOfTree, readerForListOfTree)
}

func builderForListOfTree(v types.Value) types.Value {
	return ListOfTree{v.(types.List), &ref.Ref{}}
}

func readerForListOfTree(v types.Value) types.Value {
	return v.(ListOfTree).l
}

func (l ListOfTree) Len() uint64 {
	return l.l.Len()
}

func (l ListOfTree) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfTree) Get(i uint64) Tree {
	return l.l.Get(i).(Tree)
}

func (l ListOfTree) Slice(idx uint64, end uint64) ListOfTree {
	return ListOfTree{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfTree) Set(i uint64, val Tree) ListOfTree {
	return ListOfTree{l.l.Set(i, val), &ref.Ref{}}
}

func (l ListOfTree) Append(v ...Tree) ListOfTree {
	return ListOfTree{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfTree) Insert(idx uint64, v ...Tree) ListOfTree {
	return ListOfTree{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfTree) Remove(idx uint64, end uint64) ListOfTree {
	return ListOfTree{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfTree) RemoveAt(idx uint64) ListOfTree {
	return ListOfTree{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfTree) fromElemSlice(p []Tree) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfTreeIterCallback func(v Tree, i uint64) (stop bool)

func (l ListOfTree) Iter(cb ListOfTreeIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(Tree), i)
	})
}

type ListOfTreeIterAllCallback func(v Tree, i uint64)

func (l ListOfTree) IterAll(cb ListOfTreeIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(Tree), i)
	})
}

func (l ListOfTree) IterAllP(concurrency int, cb ListOfTreeIterAllCallback) {
	l.l.IterAllP(concurrency, func(v types.Value, i uint64) {
		cb(v.(Tree), i)
	})
}

type ListOfTreeFilterCallback func(v Tree, i uint64) (keep bool)

func (l ListOfTree) Filter(cb ListOfTreeFilterCallback) ListOfTree {
	out := l.l.Filter(func(v types.Value, i uint64) bool {
		return cb(v.(Tree), i)
	})
	return ListOfTree{out, &ref.Ref{}}
}
