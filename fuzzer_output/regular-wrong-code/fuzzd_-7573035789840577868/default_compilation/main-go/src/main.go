// Dafny program main.dfy compiled into Go
package main

import (
	_System "System_"
	_dafny "dafny"
	os "os"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ _System.Dummy__

// Definition of class Default__
type Default__ struct {
	dummy byte
}

func New_Default___() *Default__ {
	_this := Default__{}

	return &_this
}

type CompanionStruct_Default___ struct {
}

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "_module.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Abs(x _dafny.Int) _dafny.Int {
	if (x).Sign() == -1 {
		return (_dafny.IntOfInt64(-1)).Times(x)
	} else {
		return x
	}
}
func (_static *CompanionStruct_Default___) SafeIndex(x _dafny.Int, length _dafny.Int) _dafny.Int {
	if (x).Sign() == -1 {
		return _dafny.Zero
	} else if (x).Cmp(length) >= 0 {
		return (x).Modulo(length)
	} else {
		return x
	}
}
func (_static *CompanionStruct_Default___) SafeDivisionInt(x1 _dafny.Int, x2 _dafny.Int) _dafny.Int {
	if (x2).Sign() == 0 {
		return x1
	} else {
		return (x1).DivBy(x2)
	}
}
func (_static *CompanionStruct_Default___) SafeModuloInt(x1 _dafny.Int, x2 _dafny.Int) _dafny.Int {
	if (x2).Sign() == 0 {
		return x1
	} else {
		return (x1).Modulo(x2)
	}
}
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) bool {
	return !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(Companion_D0_.Create_DC1_(false), Companion_D0_.Create_DC1_(!(true))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D0_.Create_DC1_(false)), _dafny.SeqOf(Companion_D0_.Create_DC1_(!(!(true)))))))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC5_(_dafny.UnicodeSeqOfUtf8Bytes("hu"), (_dafny.Zero).Minus((func() _dafny.Int {
		if false {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bgppqsue")).Cardinality())
		}
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(850), _dafny.IntOfInt64(630)))).Cardinality())
	})()), _dafny.IntOfInt64(-415))
}
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC6_((_dafny.Zero).Minus((Companion_D11_.Create_DC30_(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(false), false, true, false, true)).Cardinality()))).Cardinality()), true)).Dtor_cf53()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qipvo")).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	return (((_dafny.Zero).Minus((_dafny.SetOf(true)).Cardinality())).Plus(_dafny.IntOfInt64(384))).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kifiqbht")).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	var _source0 D0 = Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(true)))))
	_ = _source0
	if _source0.Is_DC1() {
		var _0___mcc_h0 bool = _source0.Get_().(D0_DC1).Cf1
		_ = _0___mcc_h0
		var _1_cf1 bool = _0___mcc_h0
		_ = _1_cf1
		return ((_dafny.Zero).Minus(_dafny.IntOfInt64(-50))).Minus(_dafny.IntOfInt64(435))
	} else if _source0.Is_DC0() {
		var _2___mcc_h1 bool = _source0.Get_().(D0_DC0).Cf0
		_ = _2___mcc_h1
		var _3_cf0 bool = _2___mcc_h1
		_ = _3_cf0
		return (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(95))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))).Cardinality())))
	} else {
		var _5___mcc_h2 D0 = _source0.Get_().(D0_DC2).Cf2
		_ = _5___mcc_h2
		var _6_cf2 D0 = _5___mcc_h2
		_ = _6_cf2
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(false))).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	return ((func() _dafny.Int {
		if true {
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rch")).Cardinality()))
		}
		return (_dafny.Zero).Minus((_dafny.MultiSetOf(true)).Cardinality())
	})()).Minus(((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(161), false)).Cardinality(), _dafny.IntOfInt64(240), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(!(true), false)).Cardinality())).Cardinality())).Union(_dafny.SetOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xpetoj")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(216))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	}))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(90)))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(387), _dafny.IntOfInt64(476))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _8_v0 _dafny.Int
			_8_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(387)).Cmp(_8_v0) <= 0) && ((_8_v0).Cmp(_dafny.IntOfInt64(476)) < 0) {
				_coll0.Add((_8_v0).Times(_dafny.IntOfInt64(853)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(758))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg2 _dafny.Int) interface{} {
						return coer2(arg2)
					}
				}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				}))).Cardinality()))
			}
		}
		return _coll0.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(187))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_10_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	}))).Cardinality())), _dafny.IntOfInt64(784)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(572), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(249))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bkuqdvcwu")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_(false))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) _dafny.CodePoint {
	var _source1 D0 = Companion_D0_.Create_DC0_(true)
	_ = _source1
	if _source1.Is_DC1() {
		var _11___mcc_h0 bool = _source1.Get_().(D0_DC1).Cf1
		_ = _11___mcc_h0
		var _12_cf1 bool = _11___mcc_h0
		_ = _12_cf1
		return _dafny.CodePoint('a')
	} else if _source1.Is_DC0() {
		var _13___mcc_h1 bool = _source1.Get_().(D0_DC0).Cf0
		_ = _13___mcc_h1
		var _14_cf0 bool = _13___mcc_h1
		_ = _14_cf0
		return _dafny.CodePoint('r')
	} else {
		var _15___mcc_h2 D0 = _source1.Get_().(D0_DC2).Cf2
		_ = _15___mcc_h2
		var _16_cf2 D0 = _15___mcc_h2
		_ = _16_cf2
		return _dafny.CodePoint('b')
	}
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-955))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_17_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(372)
	})), (Companion_D9_.Create_DC24_(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tqeuo")).Cardinality()))).Cardinality()), (_dafny.MultiSetOf(true)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(true)).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(456), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dst")).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-802)), _dafny.IntOfInt64(587))).Cardinality())).Cardinality())).Cardinality()))).Dtor_cf42()), _dafny.SeqOf((_dafny.MultiSetOf(true)).Cardinality(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-557), _dafny.IntOfInt64(-934))).Cardinality()), _dafny.IntOfInt64(811), (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(-612), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality()))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wggx")).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	if true {
		if !(false) {
			return _dafny.UnicodeSeqOfUtf8Bytes("lr")
		} else {
			return _dafny.UnicodeSeqOfUtf8Bytes("mrdlm")
		}
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mtnbe"), _dafny.UnicodeSeqOfUtf8Bytes("sunbm"))
	}
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(990), !((false) == (true)))
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(510))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _18_v0 _dafny.Int
			_18_v0 = interface{}(_compr_1).(_dafny.Int)
			if (_dafny.MultiSetOf(_dafny.IntOfInt64(510))).Contains(_18_v0) {
				_coll1.Add((_18_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))
			}
		}
		return _coll1.ToSet()
	}()).Difference(func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(958), _dafny.IntOfInt64(174))).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _19_v1 _dafny.Int
			_19_v1 = interface{}(_compr_2).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(958), _dafny.IntOfInt64(174))).Contains(_19_v1) {
				_coll2.Add((_19_v1).Minus(_dafny.IntOfInt64(688)))
			}
		}
		return _coll2.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((func() _dafny.Int {
		if true {
			return _dafny.IntOfInt64(779)
		}
		return _dafny.IntOfInt64(725)
	})(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), _dafny.IntOfInt64(-617))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm20(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.IntOfInt64(-355)).Cmp(_dafny.IntOfInt64(-391)) >= 0, true)
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC5_(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("b"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(524))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_20_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	}))), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xqjjx"), _dafny.UnicodeSeqOfUtf8Bytes("uvlppdqpi"))).Cardinality()), ((_dafny.Zero).Minus((_dafny.MultiSetOf(true)).Cardinality())).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(712), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(867), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_21_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	}))).Cardinality()), _dafny.CodePoint('u'))).Cardinality())).Cardinality()), (func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _22_v0 _dafny.Int
			_22_v0 = interface{}(_compr_3).(_dafny.Int)
			if (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()))).Contains(_22_v0) {
				_coll3.Add((_22_v0).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("wcolpleqj"))).Cardinality()), !(true))
			}
		}
		return _coll3.ToMap()
	}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false, !(false))).Cardinality()))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false)).IsSubsetOf(_dafny.SetOf(false)), !((Companion_D12_.Create_DC33_(_dafny.IntOfInt64(-467), false, _dafny.MultiSetOf(_dafny.IntOfInt64(-662)), false, true)).Dtor_cf61()))
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, p1 _dafny.Map, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-322))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_23_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))).Cardinality()), _dafny.IntOfInt64(119), _dafny.IntOfInt64(382), _dafny.IntOfInt64(970), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(false))).Cardinality())), false)).Cardinality())).Cardinality())))).Difference(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-465), false)).Cardinality(), _dafny.IntOfInt64(-733), _dafny.IntOfInt64(919)))
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	if !(func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(_dafny.CodePoint('n'))).Cardinality())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fgnsevrhf")).Cardinality()), _dafny.IntOfInt64(120))).Cardinality(), true)).Cardinality())).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _24_v0 _dafny.Int
			_24_v0 = interface{}(_compr_4).(_dafny.Int)
			if (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(_dafny.CodePoint('n'))).Cardinality())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fgnsevrhf")).Cardinality()), _dafny.IntOfInt64(120))).Cardinality(), true)).Cardinality())).Contains(_24_v0) {
				_coll4.Add(Companion_Default___.SafeDivisionInt(_24_v0, _dafny.IntOfInt64(-51)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uw")).Cardinality()))
			}
		}
		return _coll4.ToMap()
	}()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-855), (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(221), (_dafny.Zero).Minus((func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(783)))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _25_v1 _dafny.Map
			_25_v1 = interface{}(_compr_5).(_dafny.Map)
			if (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(783)))).Contains(_25_v1) {
				_coll5.Add(_25_v1, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(365))).Cardinality()))
			}
		}
		return _coll5.ToMap()
	}()).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(535), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))).Cardinality())) {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(437))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(785))).Cardinality())))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eceq")).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(773), _dafny.IntOfInt64(283))); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _26_v2 _dafny.Int
				_26_v2 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(773)).Cmp(_26_v2) <= 0) && ((_26_v2).Cmp(_dafny.IntOfInt64(283)) < 0) {
					_coll6.Add(Companion_Default___.SafeDivisionInt(_26_v2, _dafny.IntOfUint32((_dafny.SeqOf(!(false), true)).Cardinality())), true)
				}
			}
			return _coll6.ToMap()
		}()).Cardinality(), _dafny.IntOfInt64(537))).Cardinality())))
	}
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tusbtf"), _dafny.UnicodeSeqOfUtf8Bytes("hbix")), _dafny.IntOfInt64(972))
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, p1 _dafny.CodePoint, p2 bool, p3 bool, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC20_(!((_dafny.IntOfInt64(654)).Cmp((_dafny.SetOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-259))).Cardinality()))).Cardinality())).Cardinality()) > 0), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(422))).Cardinality(), false)
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(true, false, false))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(false, !(true))).Difference(_dafny.MultiSetOf(true))).Intersection(_dafny.MultiSetOf(true, true))
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate((func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(292), _dafny.IntOfInt64(523))); ; {
					_compr_9, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _27_v2 _dafny.Int
					_27_v2 = interface{}(_compr_9).(_dafny.Int)
					if ((_dafny.IntOfInt64(292)).Cmp(_27_v2) <= 0) && ((_27_v2).Cmp(_dafny.IntOfInt64(523)) < 0) {
						_coll9.Add((_27_v2).Minus(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
							var _coll10 = _dafny.NewBuilder()
							_ = _coll10
							for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(10), _dafny.IntOfInt64(160))); ; {
								_compr_10, _ok10 := _iter10()
								if !_ok10 {
									break
								}
								var _28_v3 _dafny.Int
								_28_v3 = interface{}(_compr_10).(_dafny.Int)
								if ((_dafny.IntOfInt64(10)).Cmp(_28_v3) <= 0) && ((_28_v3).Cmp(_dafny.IntOfInt64(160)) < 0) {
									_coll10.Add((_28_v3).Minus(_dafny.IntOfInt64(428)))
								}
							}
							return _coll10.ToSet()
						}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality()))).Cardinality())), true)
					}
				}
				return _coll9.ToMap()
			}()), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)).Cardinality()))).Keys().Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _29_v1 _dafny.Map
				_29_v1 = interface{}(_compr_8).(_dafny.Map)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					var _coll11 = _dafny.NewMapBuilder()
					_ = _coll11
					for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(292), _dafny.IntOfInt64(523))); ; {
						_compr_11, _ok11 := _iter11()
						if !_ok11 {
							break
						}
						var _27_v2 _dafny.Int
						_27_v2 = interface{}(_compr_11).(_dafny.Int)
						if ((_dafny.IntOfInt64(292)).Cmp(_27_v2) <= 0) && ((_27_v2).Cmp(_dafny.IntOfInt64(523)) < 0) {
							_coll11.Add((_27_v2).Minus(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
								var _coll12 = _dafny.NewBuilder()
								_ = _coll12
								for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(10), _dafny.IntOfInt64(160))); ; {
									_compr_12, _ok12 := _iter12()
									if !_ok12 {
										break
									}
									var _30_v3 _dafny.Int
									_30_v3 = interface{}(_compr_12).(_dafny.Int)
									if ((_dafny.IntOfInt64(10)).Cmp(_30_v3) <= 0) && ((_30_v3).Cmp(_dafny.IntOfInt64(160)) < 0) {
										_coll12.Add((_30_v3).Minus(_dafny.IntOfInt64(428)))
									}
								}
								return _coll12.ToSet()
							}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality()))).Cardinality())), true)
						}
					}
					return _coll11.ToMap()
				}()), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)).Cardinality()))).Contains(_29_v1) {
					_coll8.Add(_29_v1, _dafny.IntOfInt64(669))
				}
			}
			return _coll8.ToMap()
		}()).Keys().Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _31_v0 _dafny.Map
			_31_v0 = interface{}(_compr_7).(_dafny.Map)
			if (func() _dafny.Map {
				var _coll13 = _dafny.NewMapBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					var _coll14 = _dafny.NewMapBuilder()
					_ = _coll14
					for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(292), _dafny.IntOfInt64(523))); ; {
						_compr_14, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _27_v2 _dafny.Int
						_27_v2 = interface{}(_compr_14).(_dafny.Int)
						if ((_dafny.IntOfInt64(292)).Cmp(_27_v2) <= 0) && ((_27_v2).Cmp(_dafny.IntOfInt64(523)) < 0) {
							_coll14.Add((_27_v2).Minus(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
								var _coll15 = _dafny.NewBuilder()
								_ = _coll15
								for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(10), _dafny.IntOfInt64(160))); ; {
									_compr_15, _ok15 := _iter15()
									if !_ok15 {
										break
									}
									var _32_v3 _dafny.Int
									_32_v3 = interface{}(_compr_15).(_dafny.Int)
									if ((_dafny.IntOfInt64(10)).Cmp(_32_v3) <= 0) && ((_32_v3).Cmp(_dafny.IntOfInt64(160)) < 0) {
										_coll15.Add((_32_v3).Minus(_dafny.IntOfInt64(428)))
									}
								}
								return _coll15.ToSet()
							}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality()))).Cardinality())), true)
						}
					}
					return _coll14.ToMap()
				}()), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)).Cardinality()))).Keys().Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _29_v1 _dafny.Map
					_29_v1 = interface{}(_compr_13).(_dafny.Map)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
						var _coll16 = _dafny.NewMapBuilder()
						_ = _coll16
						for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(292), _dafny.IntOfInt64(523))); ; {
							_compr_16, _ok16 := _iter16()
							if !_ok16 {
								break
							}
							var _27_v2 _dafny.Int
							_27_v2 = interface{}(_compr_16).(_dafny.Int)
							if ((_dafny.IntOfInt64(292)).Cmp(_27_v2) <= 0) && ((_27_v2).Cmp(_dafny.IntOfInt64(523)) < 0) {
								_coll16.Add((_27_v2).Minus(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
									var _coll17 = _dafny.NewBuilder()
									_ = _coll17
									for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(10), _dafny.IntOfInt64(160))); ; {
										_compr_17, _ok17 := _iter17()
										if !_ok17 {
											break
										}
										var _33_v3 _dafny.Int
										_33_v3 = interface{}(_compr_17).(_dafny.Int)
										if ((_dafny.IntOfInt64(10)).Cmp(_33_v3) <= 0) && ((_33_v3).Cmp(_dafny.IntOfInt64(160)) < 0) {
											_coll17.Add((_33_v3).Minus(_dafny.IntOfInt64(428)))
										}
									}
									return _coll17.ToSet()
								}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality()))).Cardinality())), true)
							}
						}
						return _coll16.ToMap()
					}()), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)).Cardinality()))).Contains(_29_v1) {
						_coll13.Add(_29_v1, _dafny.IntOfInt64(669))
					}
				}
				return _coll13.ToMap()
			}()).Contains(_31_v0) {
				_coll7.Add(_31_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lukayiqqc")).Cardinality()))).Cardinality())))
			}
		}
		return _coll7.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, true, !(false), false)).Cardinality(), true), _dafny.IntOfInt64(502)))
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	if true {
		return (func() _dafny.Map {
			var _coll18 = _dafny.NewMapBuilder()
			_ = _coll18
			for _iter18 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(!(true), true), _dafny.SeqOf(false, !(true), false, false))).Elements()); ; {
				_compr_18, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _34_v0 _dafny.Sequence
				_34_v0 = interface{}(_compr_18).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(!(true), true), _dafny.SeqOf(false, !(true), false, false)), _34_v0) {
					_coll18.Add(_34_v0, false)
				}
			}
			return _coll18.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, !(!(false))), false))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), false)
	}
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, p1 _dafny.Int, globalState *GlobalState) D10 {
	return Companion_D10_.Create_DC27_((_dafny.MultiSetOf(false)).Union(_dafny.MultiSetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm32(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC7_(_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.SeqOf(false))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 bool, globalState *GlobalState) D10 {
	return Companion_D10_.Create_DC28_(false, (_dafny.IntOfInt64(59)).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(933), true)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.SeqOf(false, true, true, false), _dafny.SeqOf(true, false), _dafny.SeqOf(!(true)))).Intersection(func() _dafny.Set {
		var _coll19 = _dafny.NewBuilder()
		_ = _coll19
		for _iter19 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(609))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_35_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf(false, false, !(false))
		}))).Elements()); ; {
			_compr_19, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _36_v0 _dafny.Sequence
			_36_v0 = interface{}(_compr_19).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(609))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}(func(_35_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(false, false, !(false))
			})), _36_v0) {
				_coll19.Add(_36_v0)
			}
		}
		return _coll19.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) M9(p0 bool, p1 bool, globalState *GlobalState) (bool, bool, _dafny.Int) {
	var r0 bool = false
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	var _37_v0 _dafny.Sequence
	_ = _37_v0
	_37_v0 = _dafny.UnicodeSeqOfUtf8Bytes("dcrcctt")
	var _38_v1 _dafny.Int
	_ = _38_v1
	_38_v1 = _dafny.IntOfInt64(-916)
	var _39_v2 _dafny.Map
	_ = _39_v2
	_39_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_37_v0, Companion_Default___.Fm12(_38_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(255))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_40_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	})), _37_v0, globalState))
	var _41_v3 _dafny.Sequence
	_ = _41_v3
	_41_v3 = _dafny.SeqOf(p1)
	var _42_v4 _dafny.Set
	_ = _42_v4
	_42_v4 = _dafny.SetOf((_41_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.IntOfUint32((_41_v3).Cardinality()))).Uint32()).(bool), p1, p1)
	var _hi0 _dafny.Int = (_42_v4).Cardinality()
	_ = _hi0
	for _43_i0 := (_39_v2).Cardinality(); _43_i0.Cmp(_hi0) < 0; _43_i0 = _43_i0.Plus(_dafny.One) {
		(globalState).F17 = p1
		var _44_v5 _dafny.Sequence
		_ = _44_v5
		_44_v5 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p1), (Companion_Default___.SafeIndex(_38_v1, _dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()))).Uint32(), p0), _41_v3)
		if !(!(!_dafny.Companion_Sequence_.Contains((_44_v5).Select((Companion_Default___.SafeIndex(_38_v1, _dafny.IntOfUint32((_44_v5).Cardinality()))).Uint32()).(_dafny.Sequence), !(false) || (p1)))) {
			var _45_v6 D0
			_ = _45_v6
			_45_v6 = Companion_D0_.Create_DC1_(p0)
			var _46_v7 *C2
			_ = _46_v7
			var _nw0 *C2 = New_C2_()
			_ = _nw0
			_nw0.Ctor__(_45_v6, p1)
			_46_v7 = _nw0
			var _47_v8 *C5
			_ = _47_v8
			var _nw1 *C5 = New_C5_()
			_ = _nw1
			_nw1.Ctor__(_43_i0)
			_47_v8 = _nw1
			_37_v0 = _37_v0
			var _48_v9 _dafny.Array
			_ = _48_v9
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_0
			var _nw2 _dafny.Array
			_ = _nw2
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw2 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Set = (func(_49_v8 *C5) func(_dafny.Int) _dafny.Set {
					return func(_50_i2 _dafny.Int) _dafny.Set {
						return _dafny.SetOf((_49_v8).F18())
					}
				})(_47_v8)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw2 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw2).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw2).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_48_v9 = _nw2
			var _51_v10 _dafny.CodePoint
			_ = _51_v10
			_51_v10 = _dafny.CodePoint('e')
			var _52_v11 _dafny.Array
			_ = _52_v11
			var _nwElement0_0 bool = false
			_ = _nwElement0_0
			var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(25))
			_ = _nw3
			(_nw3).ArraySet1(_nwElement0_0, 0)
			(_nw3).ArraySet1(false, 1)
			(_nw3).ArraySet1((_46_v7).F21(), 2)
			(_nw3).ArraySet1((_46_v7).F21(), 3)
			(_nw3).ArraySet1(p1, 4)
			(_nw3).ArraySet1(!((_41_v3).Select((Companion_Default___.SafeIndex(_43_i0, _dafny.IntOfUint32((_41_v3).Cardinality()))).Uint32()).(bool)), 5)
			(_nw3).ArraySet1(p0, 6)
			(_nw3).ArraySet1((_46_v7).F21(), 7)
			(_nw3).ArraySet1((_46_v7).F21(), 8)
			(_nw3).ArraySet1(false, 9)
			(_nw3).ArraySet1((_46_v7).F21(), 10)
			(_nw3).ArraySet1(false, 11)
			(_nw3).ArraySet1(p1, 12)
			(_nw3).ArraySet1(p1, 13)
			(_nw3).ArraySet1(p0, 14)
			(_nw3).ArraySet1(p1, 15)
			(_nw3).ArraySet1(p0, 16)
			(_nw3).ArraySet1((_46_v7).F21(), 17)
			(_nw3).ArraySet1(p0, 18)
			(_nw3).ArraySet1((_46_v7).F21(), 19)
			(_nw3).ArraySet1(true, 20)
			(_nw3).ArraySet1(Companion_Default___.Fm0(p1, _51_v10, (_46_v7).F21(), globalState), 21)
			(_nw3).ArraySet1(p0, 22)
			(_nw3).ArraySet1((_46_v7).F21(), 23)
			(_nw3).ArraySet1(p0, 24)
			_52_v11 = _nw3
			var _53_v12 D1
			_ = _53_v12
			_53_v12 = Companion_D1_.Create_DC4_((_47_v8).F18(), p1, _52_v11)
			var _rhs0 _dafny.Sequence = _37_v0
			_ = _rhs0
			var _rhs1 _dafny.Array = _48_v9
			_ = _rhs1
			var _rhs2 D1 = _53_v12
			_ = _rhs2
			var _rhs3 _dafny.Int = _43_i0
			_ = _rhs3
			_37_v0 = _rhs0
			_48_v9 = _rhs1
			_53_v12 = _rhs2
			_38_v1 = _rhs3
			var _54_v13 T0
			_ = _54_v13
			var _nw4 *C0 = New_C0_()
			_ = _nw4
			_nw4.Ctor__(p1)
			_54_v13 = _nw4
			var _55_v14 *C4
			_ = _55_v14
			var _nw5 *C4 = New_C4_()
			_ = _nw5
			_nw5.Ctor__((_46_v7).F21())
			_55_v14 = _nw5
			var _56_v15 _dafny.Map
			_ = _56_v15
			_56_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_54_v13, _54_v13, _54_v13)).Cardinality())), _55_v14)
			r2 = Companion_Default___.Fm3(((_56_v15).Update(_43_i0, _55_v14)).Cardinality(), (_dafny.Zero).Minus(_43_i0), p1, globalState)
		} else {
			var _57_v16 _dafny.Array
			_ = _57_v16
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_1
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Sequence = (func(_58_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_59_i3 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hi"), _58_v0)
					}
				})(_37_v0)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw6 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw6).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw6).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_57_v16 = _nw6
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_57_v16), 0))
			_ = _index0
			(_57_v16).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_37_v0, _37_v0), (_index0).Int())
			var _60_v17 D0
			_ = _60_v17
			_60_v17 = Companion_D0_.Create_DC1_(p1)
			var _61_v18 *C2
			_ = _61_v18
			var _nw7 *C2 = New_C2_()
			_ = _nw7
			_nw7.Ctor__(_60_v17, p1)
			_61_v18 = _nw7
			var _62_v19 _dafny.Array
			_ = _62_v19
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_2
			var _nw8 _dafny.Array
			_ = _nw8
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw8 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Sequence = (func(_63_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_64_i4 _dafny.Int) _dafny.Sequence {
						return _63_v3
					}
				})(_41_v3)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw8 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw8).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw8).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_62_v19 = _nw8
			var _65_v20 D14
			_ = _65_v20
			_65_v20 = Companion_D14_.Create_DC41_(_61_v18, _dafny.UnicodeSeqOfUtf8Bytes("sos"), _62_v19)
			var _pat_let_tv0 = _62_v19
			_ = _pat_let_tv0
			var _pat_let_tv1 = _61_v18
			_ = _pat_let_tv1
			var _66_v21 _dafny.Map
			_ = _66_v21
			_66_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func(_pat_let0_0 D14) D14 {
				return func(_67_dt__update__tmp_h0 D14) D14 {
					return func(_pat_let1_0 _dafny.Array) D14 {
						return func(_68_dt__update_hcf74_h0 _dafny.Array) D14 {
							return func(_pat_let2_0 *C2) D14 {
								return func(_69_dt__update_hcf72_h0 *C2) D14 {
									return Companion_D14_.Create_DC41_(_69_dt__update_hcf72_h0, (_67_dt__update__tmp_h0).Dtor_cf73(), _68_dt__update_hcf74_h0)
								}(_pat_let2_0)
							}(_pat_let_tv1)
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(_65_v20)).Dtor_cf73(), _37_v0)
			var _70_v22 _dafny.Set
			_ = _70_v22
			_70_v22 = _dafny.SetOf(_38_v1)
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_57_v16), 0))
			_ = _index1
			var _rhs4 _dafny.Sequence = (func() _dafny.Sequence {
				if (_66_v21).Contains(_37_v0) {
					return (_66_v21).Get(_37_v0).(_dafny.Sequence)
				}
				return _37_v0
			})()
			_ = _rhs4
			var _rhs5 bool = (_61_v18).F21()
			_ = _rhs5
			var _rhs6 bool = (_61_v18).F21()
			_ = _rhs6
			var _rhs7 bool = ((_70_v22).Union(_70_v22)).IsSubsetOf(_70_v22)
			_ = _rhs7
			var _lhs0 _dafny.Array = _57_v16
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_57_v16), 0))
			_ = _lhs1
			var _lhs2 *GlobalState = globalState
			_ = _lhs2
			var _lhs3 *GlobalState = globalState
			_ = _lhs3
			var _lhs4 *GlobalState = globalState
			_ = _lhs4
			(_lhs0).ArraySet1(_rhs4, (_lhs1).Int())
			_lhs2.F17 = _rhs5
			_lhs3.F11 = _rhs6
			_lhs4.F15 = _rhs7
			var _71_v23 _dafny.Map
			_ = _71_v23
			_71_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, false)
			var _72_v24 _dafny.Map
			_ = _72_v24
			_72_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _71_v23)
			_72_v24 = _72_v24
			var _rhs8 _dafny.Int = Companion_Default___.Fm3(Companion_Default___.SafeDivisionInt(_43_i0, _38_v1), (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(148))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}(func(_73_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			}))).Cardinality())).Minus(_38_v1), p0, globalState)
			_ = _rhs8
			var _rhs9 _dafny.Int = _43_i0
			_ = _rhs9
			r2 = _rhs8
			r2 = _rhs9
			var _74_v25 *C4
			_ = _74_v25
			var _nw9 *C4 = New_C4_()
			_ = _nw9
			_nw9.Ctor__((_61_v18).F21())
			_74_v25 = _nw9
			var _75_v26 bool
			_ = _75_v26
			var _out0 bool
			_ = _out0
			_out0 = (_61_v18).M5((p0) && (false), _41_v3, globalState)
			_75_v26 = _out0
		}
		_38_v1 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-352), _dafny.IntOfUint32((_dafny.SeqOf(_43_i0)).Cardinality()))
		var _76_v27 _dafny.Array
		_ = _76_v27
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_3
		var _nw10 _dafny.Array
		_ = _nw10
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw10 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = (func(_77_p1 bool) func(_dafny.Int) bool {
				return func(_78_i6 _dafny.Int) bool {
					return _77_p1
				}
			})(p1)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw10 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw10).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw10).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_76_v27 = _nw10
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_76_v27), 0))
		_ = _index2
		(_76_v27).ArraySet1(p0, (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_76_v27), 0))
		_ = _index3
		(_76_v27).ArraySet1(p0, (_index3).Int())
	}
	var _79_v28 _dafny.MultiSet
	_ = _79_v28
	_79_v28 = _dafny.MultiSetOf(p0)
	if (_79_v28).IsSubsetOf((_dafny.MultiSetOf(p0)).Intersection(_79_v28)) {
		var _80_v29 _dafny.Array
		_ = _80_v29
		var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
		_ = _nw11
		_80_v29 = _nw11
		var _81_v30 _dafny.MultiSet
		_ = _81_v30
		_81_v30 = _dafny.MultiSetOf(_38_v1)
		var _82_v31 _dafny.Map
		_ = _82_v31
		_82_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v29, ((_dafny.Zero).Minus((_81_v30).Cardinality())).Times(_38_v1))
		_82_v31 = (_82_v31).Update(_80_v29, _38_v1)
		var _83_v32 *C0
		_ = _83_v32
		var _nw12 *C0 = New_C0_()
		_ = _nw12
		_nw12.Ctor__(p0)
		_83_v32 = _nw12
		var _84_v33 _dafny.Set
		_ = _84_v33
		_84_v33 = _dafny.SetOf(_83_v32)
		_38_v1 = (_84_v33).Cardinality()
		var _85_v34 _dafny.Array
		_ = _85_v34
		var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
		_ = _nw13
		_85_v34 = _nw13
		var _86_v35 T0
		_ = _86_v35
		var _nw14 *C4 = New_C4_()
		_ = _nw14
		_nw14.Ctor__((_83_v32).F23())
		_86_v35 = _nw14
		var _87_v36 _dafny.Sequence
		_ = _87_v36
		_87_v36 = _dafny.SeqOf(_86_v35, _86_v35, _86_v35)
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_85_v34), 0))
		_ = _index4
		(_85_v34).ArraySet1(_87_v36, (_index4).Int())
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_85_v34), 0))
		_ = _index5
		var _rhs10 bool = (_38_v1).Cmp((func() _dafny.Int {
			if true {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gvcj")).Cardinality())
			}
			return (_dafny.Zero).Minus(_38_v1)
		})()) > 0
		_ = _rhs10
		var _rhs11 bool = (_dafny.IntOfInt64(429)).Cmp(_38_v1) != 0
		_ = _rhs11
		var _rhs12 _dafny.Sequence = _87_v36
		_ = _rhs12
		var _lhs5 *GlobalState = globalState
		_ = _lhs5
		var _lhs6 *GlobalState = globalState
		_ = _lhs6
		var _lhs7 _dafny.Array = _85_v34
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_85_v34), 0))
		_ = _lhs8
		_lhs5.F17 = _rhs10
		_lhs6.F11 = _rhs11
		(_lhs7).ArraySet1(_rhs12, (_lhs8).Int())
		var _88_v37 D12
		_ = _88_v37
		_88_v37 = Companion_D12_.Create_DC34_()
		var _source2 D12 = (func() D12 {
			if (_83_v32).F23() {
				return _88_v37
			}
			return _88_v37
		})()
		_ = _source2
		if _source2.Is_DC32() {
			var _89___mcc_h0 _dafny.Int = _source2.Get_().(D12_DC32).Cf56
			_ = _89___mcc_h0
			var _90_cf56 _dafny.Int = _89___mcc_h0
			_ = _90_cf56
			var _91_v38 D0
			_ = _91_v38
			_91_v38 = Companion_D0_.Create_DC1_(p1)
			var _92_v39 *C2
			_ = _92_v39
			var _nw15 *C2 = New_C2_()
			_ = _nw15
			_nw15.Ctor__(_91_v38, (_83_v32).F23())
			_92_v39 = _nw15
			r2 = _90_cf56
			(globalState).F11 = true
			var _93_v40 _dafny.Map
			_ = _93_v40
			_93_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_92_v39).F21(), _90_cf56)
			var _94_v41 _dafny.Array
			_ = _94_v41
			var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw16
			_94_v41 = _nw16
			var _95_v42 D3
			_ = _95_v42
			_95_v42 = Companion_D3_.Create_DC11_(_93_v40, _90_cf56, _94_v41)
			var _rhs13 _dafny.Int = _38_v1
			_ = _rhs13
			var _rhs14 _dafny.Int = (_dafny.Zero).Minus((((_dafny.MultiSetOf((_83_v32).F23(), (_83_v32).F23())).Intersection(_dafny.MultiSetOf((_83_v32).F23(), p0))).Cardinality()).Plus((_42_v4).Cardinality()))
			_ = _rhs14
			var _rhs15 D3 = _95_v42
			_ = _rhs15
			_90_cf56 = _rhs13
			_38_v1 = _rhs14
			_95_v42 = _rhs15
		} else if _source2.Is_DC33() {
			var _96___mcc_h1 _dafny.Int = _source2.Get_().(D12_DC33).Cf57
			_ = _96___mcc_h1
			var _97___mcc_h2 bool = _source2.Get_().(D12_DC33).Cf58
			_ = _97___mcc_h2
			var _98___mcc_h3 _dafny.MultiSet = _source2.Get_().(D12_DC33).Cf59
			_ = _98___mcc_h3
			var _99___mcc_h4 bool = _source2.Get_().(D12_DC33).Cf60
			_ = _99___mcc_h4
			var _100___mcc_h5 bool = _source2.Get_().(D12_DC33).Cf61
			_ = _100___mcc_h5
			var _101_cf61 bool = _100___mcc_h5
			_ = _101_cf61
			var _102_cf60 bool = _99___mcc_h4
			_ = _102_cf60
			var _103_cf59 _dafny.MultiSet = _98___mcc_h3
			_ = _103_cf59
			var _104_cf58 bool = _97___mcc_h2
			_ = _104_cf58
			var _105_cf57 _dafny.Int = _96___mcc_h1
			_ = _105_cf57
			r2 = _38_v1
			var _106_v43 _dafny.Map
			_ = _106_v43
			_106_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v32, Companion_Default___.Fm20(globalState))
			_42_v4 = ((_42_v4).Difference(_42_v4)).Union((func() _dafny.Set {
				if (_106_v43).Contains(_83_v32) {
					return (_106_v43).Get(_83_v32).(_dafny.Set)
				}
				return _42_v4
			})())
			var _107_v44 _dafny.CodePoint
			_ = _107_v44
			_107_v44 = _dafny.CodePoint('o')
			var _108_v45 _dafny.Map
			_ = _108_v45
			_108_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v44, _107_v44)
			_108_v45 = _108_v45
			var _109_v46 D11
			_ = _109_v46
			_109_v46 = Companion_D11_.Create_DC29_(_42_v4)
			var _110_v47 _dafny.Sequence
			_ = _110_v47
			_110_v47 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_cf57, ((_109_v46).Dtor_cf51()).Cardinality()))
			var _111_v48 D0
			_ = _111_v48
			_111_v48 = Companion_D0_.Create_DC1_(_102_cf60)
			var _112_v49 *C2
			_ = _112_v49
			var _nw17 *C2 = New_C2_()
			_ = _nw17
			_nw17.Ctor__(_111_v48, _104_cf58)
			_112_v49 = _nw17
			var _113_v50 _dafny.Array
			_ = _113_v50
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_4
			var _nw18 _dafny.Array
			_ = _nw18
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw18 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Sequence = (func(_114_cf60 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_115_i7 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_114_cf60)
					}
				})(_102_cf60)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw18 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw18).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw18).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_113_v50 = _nw18
			var _116_v51 _dafny.MultiSet
			_ = _116_v51
			_116_v51 = _dafny.MultiSetOf(Companion_D14_.Create_DC41_(_112_v49, _37_v0, _113_v50))
			var _117_v52 _dafny.Map
			_ = _117_v52
			_117_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm27(_105_cf57, _105_cf57, ((_110_v47).Select((Companion_Default___.SafeIndex((_116_v51).Cardinality(), _dafny.IntOfUint32((_110_v47).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), globalState)).Cardinality()), _dafny.IntOfInt64(533))
			var _118_v53 _dafny.Set
			_ = _118_v53
			_118_v53 = _dafny.SetOf(_117_v52)
			var _119_v54 D6
			_ = _119_v54
			_119_v54 = Companion_D6_.Create_DC20_((_83_v32).F23(), _105_cf57, p0)
			var _120_v55 _dafny.Map
			_ = _120_v55
			_120_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(72))
			var _121_v56 _dafny.Set
			_ = _121_v56
			_121_v56 = _dafny.SetOf(_38_v1, (Companion_Default___.Fm23((_83_v32).F23(), _120_v55, _38_v1, _107_v44, globalState)).Cardinality())
			var _122_v57 _dafny.MultiSet
			_ = _122_v57
			_122_v57 = _dafny.MultiSetOf(_120_v55)
			var _123_v58 _dafny.Sequence
			_ = _123_v58
			_123_v58 = _dafny.SeqOf(_dafny.IntOfInt64(949))
			var _124_v59 _dafny.Array
			_ = _124_v59
			var _nwElement0_1 _dafny.Int = (_83_v32).Fm14((_118_v53).Cardinality(), globalState)
			_ = _nwElement0_1
			var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(26))
			_ = _nw19
			(_nw19).ArraySet1(_nwElement0_1, 0)
			(_nw19).ArraySet1(_105_cf57, 1)
			(_nw19).ArraySet1(_38_v1, 2)
			(_nw19).ArraySet1(_105_cf57, 3)
			(_nw19).ArraySet1(_38_v1, 4)
			(_nw19).ArraySet1((_119_v54).Dtor_cf37(), 5)
			(_nw19).ArraySet1(_105_cf57, 6)
			(_nw19).ArraySet1(Companion_Default___.SafeDivisionInt(_38_v1, _38_v1), 7)
			(_nw19).ArraySet1((func() _dafny.Int {
				if (_83_v32).F23() {
					return _38_v1
				}
				return _38_v1
			})(), 8)
			(_nw19).ArraySet1(_105_cf57, 9)
			(_nw19).ArraySet1(_105_cf57, 10)
			(_nw19).ArraySet1((Companion_D4_.Create_DC14_(p0, _38_v1, _121_v56, _107_v44)).Dtor_cf24(), 11)
			(_nw19).ArraySet1(_105_cf57, 12)
			(_nw19).ArraySet1(_105_cf57, 13)
			(_nw19).ArraySet1(_38_v1, 14)
			(_nw19).ArraySet1(_105_cf57, 15)
			(_nw19).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_112_v49).F21() {
					return _dafny.UnicodeSeqOfUtf8Bytes("kuj")
				}
				return _37_v0
			})()).Cardinality()), 16)
			(_nw19).ArraySet1(_dafny.IntOfInt64(12), 17)
			(_nw19).ArraySet1((((_122_v57).Update(_120_v55, Companion_Default___.Abs(_dafny.IntOfUint32((_123_v58).Cardinality())))).Union(_122_v57)).Cardinality(), 18)
			(_nw19).ArraySet1((_dafny.Zero).Minus((_79_v28).Cardinality()), 19)
			(_nw19).ArraySet1(_38_v1, 20)
			(_nw19).ArraySet1(_dafny.IntOfUint32((_37_v0).Cardinality()), 21)
			(_nw19).ArraySet1(_105_cf57, 22)
			(_nw19).ArraySet1(_105_cf57, 23)
			(_nw19).ArraySet1((_dafny.Zero).Minus(_105_cf57), 24)
			(_nw19).ArraySet1(_38_v1, 25)
			_124_v59 = _nw19
			var _125_v60 _dafny.Map
			_ = _125_v60
			_125_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_105_cf57).Plus(_38_v1), _124_v59)
			var _rhs16 _dafny.Int = _38_v1
			_ = _rhs16
			var _rhs17 _dafny.Array = (func() _dafny.Array {
				if (_125_v60).Contains((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_105_cf57, _105_cf57))) {
					return (_125_v60).Get((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_105_cf57, _105_cf57))).(_dafny.Array)
				}
				return _124_v59
			})()
			_ = _rhs17
			r2 = _rhs16
			_124_v59 = _rhs17
		} else if _source2.Is_DC34() {
			var _126_v61 _dafny.Map
			_ = _126_v61
			_126_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _38_v1)
			var _127_v62 _dafny.Array
			_ = _127_v62
			var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw20
			_127_v62 = _nw20
			var _128_v63 D3
			_ = _128_v63
			_128_v63 = Companion_D3_.Create_DC11_(_126_v61, _dafny.IntOfUint32((_37_v0).Cardinality()), _127_v62)
			var _129_v64 *C1
			_ = _129_v64
			var _nw21 *C1 = New_C1_()
			_ = _nw21
			_nw21.Ctor__((_128_v63).Dtor_cf20())
			_129_v64 = _nw21
			var _130_v65 _dafny.Sequence
			_ = _130_v65
			_130_v65 = _dafny.SeqOf(_129_v64)
			(globalState).F15 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_130_v65, _130_v65)).Cardinality())).Cmp(_38_v1) > 0
			var _131_v66 _dafny.Sequence
			_ = _131_v66
			_131_v66 = _dafny.SeqOf(_38_v1)
			var _132_v67 _dafny.Sequence
			_ = _132_v67
			_132_v67 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_133_p1 bool, _134_v32 *C0) func(_dafny.Int) _dafny.Int {
				return func(_135_i8 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_p1, (_134_v32).F23()))).Cardinality()))
				}
			})(p1, _83_v32))), (Companion_Default___.SafeIndex(_38_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_136_p1 bool, _137_v32 *C0) func(_dafny.Int) _dafny.Int {
				return func(_138_i8 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_136_p1, (_137_v32).F23()))).Cardinality()))
				}
			})(p1, _83_v32)))).Cardinality()))).Uint32(), _38_v1), _131_v66, _131_v66)
			var _139_v68 _dafny.MultiSet
			_ = _139_v68
			_139_v68 = _dafny.MultiSetOf((_132_v67).Select((Companion_Default___.SafeIndex(_38_v1, _dafny.IntOfUint32((_132_v67).Cardinality()))).Uint32()).(_dafny.Sequence), _131_v66)
			_139_v68 = ((_139_v68).Union(_139_v68)).Difference((_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(648))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_140_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_141_i9 _dafny.Int) _dafny.Int {
					return _140_v1
				}
			})(_38_v1))))).Union(_139_v68))
			var _142_v69 _dafny.Array
			_ = _142_v69
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_5
			var _nw22 _dafny.Array
			_ = _nw22
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw22 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) bool = (func(_143_v32 *C0) func(_dafny.Int) bool {
					return func(_144_i10 _dafny.Int) bool {
						return ((_143_v32).F23()) || ((_143_v32).F23())
					}
				})(_83_v32)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw22 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw22).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw22).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_142_v69 = _nw22
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_142_v69), 0))
			_ = _index6
			(_142_v69).ArraySet1((_83_v32).F23(), (_index6).Int())
			var _145_v70 _dafny.Set
			_ = _145_v70
			_145_v70 = _dafny.SetOf(_41_v3)
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_142_v69), 0))
			_ = _index7
			var _rhs18 T0 = _86_v35
			_ = _rhs18
			var _rhs19 bool = (Companion_Default___.Fm34((_131_v66).Select((Companion_Default___.SafeIndex(_38_v1, _dafny.IntOfUint32((_131_v66).Cardinality()))).Uint32()).(_dafny.Int), _38_v1, true, globalState)).IsSubsetOf(_145_v70)
			_ = _rhs19
			var _rhs20 bool = p0
			_ = _rhs20
			var _lhs9 _dafny.Array = _142_v69
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_142_v69), 0))
			_ = _lhs10
			var _lhs11 *GlobalState = globalState
			_ = _lhs11
			_86_v35 = _rhs18
			(_lhs9).ArraySet1(_rhs19, (_lhs10).Int())
			_lhs11.F11 = _rhs20
			var _146_v71 D12
			_ = _146_v71
			_146_v71 = Companion_D12_.Create_DC33_(_38_v1, false, _81_v30, p1, p0)
			_38_v1 = (_146_v71).Dtor_cf57()
		} else if _source2.Is_DC31() {
			var _147___mcc_h6 _dafny.Set = _source2.Get_().(D12_DC31).Cf55
			_ = _147___mcc_h6
			var _148_cf55 _dafny.Set = _147___mcc_h6
			_ = _148_cf55
			var _149_v72 _dafny.Array
			_ = _149_v72
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_6
			var _nw23 _dafny.Array
			_ = _nw23
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw23 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Int = (func(_150_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_151_i11 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_151_i11, _150_v1)
					}
				})(_38_v1)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw23 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw23).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw23).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_149_v72 = _nw23
			var _152_v73 *C1
			_ = _152_v73
			var _nw24 *C1 = New_C1_()
			_ = _nw24
			_nw24.Ctor__((func() _dafny.Array {
				if p1 {
					return _149_v72
				}
				return _149_v72
			})())
			_152_v73 = _nw24
			var _153_v74 D0
			_ = _153_v74
			_153_v74 = Companion_D0_.Create_DC1_((_83_v32).F23())
			var _154_v75 *C2
			_ = _154_v75
			var _nw25 *C2 = New_C2_()
			_ = _nw25
			_nw25.Ctor__(_153_v74, ((_83_v32).F23()) == (p1))
			_154_v75 = _nw25
			_154_v75 = _154_v75
			(globalState).F11 = (_83_v32).F23()
			var _155_v76 D12
			_ = _155_v76
			_155_v76 = Companion_D12_.Create_DC32_(_38_v1)
			var _156_v77 _dafny.Map
			_ = _156_v77
			_156_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v76, _149_v72)
			_149_v72 = (func() _dafny.Array {
				if (_156_v77).Contains(_155_v76) {
					return (_156_v77).Get(_155_v76).(_dafny.Array)
				}
				return (_152_v73).F22()
			})()
		} else {
			var _157___mcc_h7 D12 = _source2.Get_().(D12_DC35).Cf62
			_ = _157___mcc_h7
			var _158_cf62 D12 = _157___mcc_h7
			_ = _158_cf62
			(globalState).F11 = !(p0) || ((_83_v32).F23())
			var _159_v78 _dafny.Array
			_ = _159_v78
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw26
			_159_v78 = _nw26
			var _160_v79 *C1
			_ = _160_v79
			var _nw27 *C1 = New_C1_()
			_ = _nw27
			_nw27.Ctor__(_159_v78)
			_160_v79 = _nw27
			var _161_v80 _dafny.Map
			_ = _161_v80
			_161_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v79, (_dafny.Zero).Minus(_38_v1))
			_38_v1 = (_dafny.Zero).Minus(((_161_v80).Merge((_161_v80).Update(_160_v79, _dafny.IntOfInt64(161)))).Cardinality())
			var _162_v81 _dafny.CodePoint
			_ = _162_v81
			_162_v81 = _dafny.CodePoint('x')
			_162_v81 = _162_v81
			var _163_v82 _dafny.Map
			_ = _163_v82
			_163_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _38_v1)
			_163_v82 = (_163_v82).Update((_83_v32).F23(), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_38_v1, _38_v1)))
		}
		(globalState).F11 = p1
	} else {
		var _164_v83 _dafny.Array
		_ = _164_v83
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
		_ = _nw28
		_164_v83 = _nw28
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_164_v83), 0))
		_ = _index8
		(_164_v83).ArraySet1(p1, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_164_v83), 0))
		_ = _index9
		(_164_v83).ArraySet1(!((_38_v1).Cmp(_38_v1) < 0), (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_164_v83), 0))
		_ = _index10
		(_164_v83).ArraySet1((_164_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_164_v83), 0))).Int()).(bool), (_index10).Int())
		var _165_v84 _dafny.Sequence
		_ = _165_v84
		_165_v84 = _dafny.SeqOf(_37_v0, _37_v0, _37_v0, _37_v0)
		var _166_v85 _dafny.Sequence
		_ = _166_v85
		_166_v85 = _dafny.SeqOf(_165_v84)
		_165_v84 = (_166_v85).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_37_v0).Cardinality()), _dafny.IntOfUint32((_166_v85).Cardinality()))).Uint32()).(_dafny.Sequence)
		var _167_v86 _dafny.Sequence
		_ = _167_v86
		_167_v86 = _dafny.SeqOf((_dafny.Zero).Minus(_38_v1), _dafny.IntOfInt64(256), _dafny.IntOfUint32((_37_v0).Cardinality()), _38_v1)
		(globalState).F11 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm19(_dafny.SeqOf(_38_v1), p1, globalState), (Companion_Default___.SafeIndex(_38_v1, _dafny.IntOfUint32((Companion_Default___.Fm19(_dafny.SeqOf(_38_v1), p1, globalState)).Cardinality()))).Uint32(), _38_v1), _167_v86)
		var _168_v87 _dafny.Map
		_ = _168_v87
		_168_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _38_v1)
		_168_v87 = (_168_v87).Merge((_168_v87).Merge(_168_v87))
	}
	_41_v3 = _41_v3
	(globalState).F15 = p0
	var _169_v88 _dafny.Array
	_ = _169_v88
	var _nw29 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(29))
	_ = _nw29
	_169_v88 = _nw29
	var _170_v89 _dafny.Array
	_ = _170_v89
	var _nwElement0_2 _dafny.Int = _38_v1
	_ = _nwElement0_2
	var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(8))
	_ = _nw30
	(_nw30).ArraySet1(_nwElement0_2, 0)
	(_nw30).ArraySet1(_38_v1, 1)
	(_nw30).ArraySet1(_38_v1, 2)
	(_nw30).ArraySet1(_38_v1, 3)
	(_nw30).ArraySet1(_38_v1, 4)
	(_nw30).ArraySet1(_38_v1, 5)
	(_nw30).ArraySet1(_38_v1, 6)
	(_nw30).ArraySet1(_38_v1, 7)
	_170_v89 = _nw30
	var _171_v90 D13
	_ = _171_v90
	_171_v90 = Companion_D13_.Create_DC38_(Companion_Default___.Fm12(_dafny.IntOfInt64(-248), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-943))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_172_i12 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	})), _37_v0, globalState), p0, (_dafny.Zero).Minus(_38_v1), _170_v89)
	var _173_v91 *C1
	_ = _173_v91
	var _nw31 *C1 = New_C1_()
	_ = _nw31
	_nw31.Ctor__((_171_v90).Dtor_cf69())
	_173_v91 = _nw31
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_169_v88), 0))
	_ = _index11
	(_169_v88).ArraySet1(_173_v91, (_index11).Int())
	var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_169_v88), 0))
	_ = _index12
	(_169_v88).ArraySet1(_173_v91, (_index12).Int())
	(globalState).F11 = (_38_v1).Cmp((_dafny.IntOfUint32((_dafny.SeqOf(false, p1)).Cardinality())).Times(_38_v1)) == 0
	r0 = !((_38_v1).Cmp(_38_v1) < 0)
	r1 = p1
	r2 = _38_v1
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _174_v0 bool
	_ = _174_v0
	_174_v0 = true
	var _175_v1 _dafny.Sequence
	_ = _175_v1
	_175_v1 = _dafny.SeqOf(_174_v0, _174_v0)
	var _176_v2 _dafny.Int
	_ = _176_v2
	_176_v2 = _dafny.IntOfInt64(644)
	var _177_v3 _dafny.Array
	_ = _177_v3
	var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
	_ = _nw32
	_177_v3 = _nw32
	var _178_v4 _dafny.Array
	_ = _178_v4
	var _nwElement0_3 _dafny.Array = _177_v3
	_ = _nwElement0_3
	var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(28))
	_ = _nw33
	(_nw33).ArraySet1(_nwElement0_3, 0)
	(_nw33).ArraySet1(_177_v3, 1)
	(_nw33).ArraySet1(_177_v3, 2)
	(_nw33).ArraySet1(_177_v3, 3)
	(_nw33).ArraySet1(_177_v3, 4)
	(_nw33).ArraySet1(_177_v3, 5)
	(_nw33).ArraySet1(_177_v3, 6)
	(_nw33).ArraySet1(_177_v3, 7)
	(_nw33).ArraySet1(_177_v3, 8)
	(_nw33).ArraySet1(_177_v3, 9)
	(_nw33).ArraySet1(_177_v3, 10)
	(_nw33).ArraySet1(_177_v3, 11)
	(_nw33).ArraySet1(_177_v3, 12)
	(_nw33).ArraySet1(_177_v3, 13)
	(_nw33).ArraySet1(_177_v3, 14)
	(_nw33).ArraySet1(_177_v3, 15)
	(_nw33).ArraySet1(_177_v3, 16)
	(_nw33).ArraySet1(_177_v3, 17)
	(_nw33).ArraySet1(_177_v3, 18)
	(_nw33).ArraySet1(_177_v3, 19)
	(_nw33).ArraySet1(_177_v3, 20)
	(_nw33).ArraySet1(_177_v3, 21)
	(_nw33).ArraySet1(_177_v3, 22)
	(_nw33).ArraySet1(_177_v3, 23)
	(_nw33).ArraySet1(_177_v3, 24)
	(_nw33).ArraySet1(_177_v3, 25)
	(_nw33).ArraySet1(_177_v3, 26)
	(_nw33).ArraySet1(_177_v3, 27)
	_178_v4 = _nw33
	var _179_v5 D0
	_ = _179_v5
	_179_v5 = Companion_D0_.Create_DC1_(_174_v0)
	var _180_v6 _dafny.Sequence
	_ = _180_v6
	_180_v6 = _dafny.UnicodeSeqOfUtf8Bytes("df")
	var _181_v7 _dafny.Map
	_ = _181_v7
	_181_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v2, _176_v2)
	var _182_globalState *GlobalState
	_ = _182_globalState
	var _nw34 *GlobalState = New_GlobalState_()
	_ = _nw34
	_nw34.Ctor__(_dafny.CodePoint('x'), _dafny.Companion_Sequence_.Update(_175_v1, (Companion_Default___.SafeIndex(_176_v2, _dafny.IntOfUint32((_175_v1).Cardinality()))).Uint32(), _174_v0), _178_v4, false, _dafny.Companion_Sequence_.Concatenate(_175_v1, _dafny.SeqOf(_174_v0)), _dafny.CodePoint('e'), _dafny.IntOfInt64(-835), false, (func() _dafny.Sequence {
		if (_179_v5).Dtor_cf1() {
			return _180_v6
		}
		return _180_v6
	})(), false, true, false, _181_v7, true, _dafny.IntOfInt64(-305), true, _dafny.Companion_Sequence_.Concatenate(_175_v1, _175_v1), false)
	_182_globalState = _nw34
	var _183_v8 _dafny.CodePoint
	_ = _183_v8
	_183_v8 = _dafny.CodePoint('c')
	if Companion_Default___.Fm0(_174_v0, _183_v8, _174_v0, _182_globalState) {
		var _184_v9 _dafny.Map
		_ = _184_v9
		_184_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _174_v0)
		var _185_v10 _dafny.Sequence
		_ = _185_v10
		_185_v10 = _dafny.SeqOf(_176_v2)
		var _186_v11 D0
		_ = _186_v11
		_186_v11 = Companion_D0_.Create_DC0_(_174_v0)
		var _187_v12 _dafny.Array
		_ = _187_v12
		var _nwElement0_4 bool = (_175_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.IntOfUint32((_175_v1).Cardinality()))).Uint32()).(bool)
		_ = _nwElement0_4
		var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(15))
		_ = _nw35
		(_nw35).ArraySet1(_nwElement0_4, 0)
		(_nw35).ArraySet1(_174_v0, 1)
		(_nw35).ArraySet1(_174_v0, 2)
		(_nw35).ArraySet1(false, 3)
		(_nw35).ArraySet1(false, 4)
		(_nw35).ArraySet1(_174_v0, 5)
		(_nw35).ArraySet1(!(_174_v0) || (_174_v0), 6)
		(_nw35).ArraySet1(_174_v0, 7)
		(_nw35).ArraySet1((_184_v9).Contains(true), 8)
		(_nw35).ArraySet1(_174_v0, 9)
		(_nw35).ArraySet1(_174_v0, 10)
		(_nw35).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_185_v10, _185_v10), 11)
		(_nw35).ArraySet1(_174_v0, 12)
		(_nw35).ArraySet1((Companion_D0_.Create_DC1_((_186_v11).Dtor_cf0())).Dtor_cf1(), 13)
		(_nw35).ArraySet1(_174_v0, 14)
		_187_v12 = _nw35
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_187_v12), 0))
		_ = _index13
		(_187_v12).ArraySet1(_174_v0, (_index13).Int())
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_187_v12), 0))
		_ = _index14
		(_187_v12).ArraySet1(_174_v0, (_index14).Int())
		var _rhs21 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(829))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_188_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_189_i0 _dafny.Int) _dafny.CodePoint {
				return _188_v8
			}
		})(_183_v8)))
		_ = _rhs21
		var _rhs22 _dafny.Int = _dafny.IntOfInt64(-916)
		_ = _rhs22
		_180_v6 = _rhs21
		_176_v2 = _rhs22
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_187_v12), 0))
		_ = _index15
		(_187_v12).ArraySet1((_174_v0) || ((_187_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_187_v12), 0))).Int()).(bool)), (_index15).Int())
		var _190_v13 *C3
		_ = _190_v13
		var _nw36 *C3 = New_C3_()
		_ = _nw36
		_nw36.Ctor__()
		_190_v13 = _nw36
		var _191_v14 D6
		_ = _191_v14
		_191_v14 = Companion_D6_.Create_DC20_(_174_v0, _176_v2, _174_v0)
		var _192_v15 T0
		_ = _192_v15
		var _nw37 *C2 = New_C2_()
		_ = _nw37
		_nw37.Ctor__(_179_v5, (_191_v14).Dtor_cf36())
		_192_v15 = _nw37
		var _193_v16 _dafny.Array
		_ = _193_v16
		var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
		_ = _nw38
		_193_v16 = _nw38
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_193_v16), 0))
		_ = _index16
		(_193_v16).ArraySet1((func() _dafny.Array {
			if _174_v0 {
				return _187_v12
			}
			return _187_v12
		})(), (_index16).Int())
		var _194_v17 _dafny.Array
		_ = _194_v17
		var _nw39 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
		_ = _nw39
		_194_v17 = _nw39
		var _195_v18 T1
		_ = _195_v18
		var _nw40 *C4 = New_C4_()
		_ = _nw40
		_nw40.Ctor__((_187_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_187_v12), 0))).Int()).(bool))
		_195_v18 = _nw40
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_194_v17), 0))
		_ = _index17
		(_194_v17).ArraySet1(_195_v18, (_index17).Int())
		var _196_v19 _dafny.Sequence
		_ = _196_v19
		_196_v19 = _dafny.SeqOf(_192_v15, _192_v15, _192_v15)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_193_v16), 0))
		_ = _index18
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_194_v17), 0))
		_ = _index19
		var _rhs23 _dafny.Array = _187_v12
		_ = _rhs23
		var _rhs24 T0 = (_196_v19).Select((Companion_Default___.SafeIndex(_176_v2, _dafny.IntOfUint32((_196_v19).Cardinality()))).Uint32()).(T0)
		_ = _rhs24
		var _rhs25 _dafny.Array = _187_v12
		_ = _rhs25
		var _rhs26 bool = _174_v0
		_ = _rhs26
		var _rhs27 T1 = _195_v18
		_ = _rhs27
		var _lhs12 _dafny.Array = _193_v16
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_193_v16), 0))
		_ = _lhs13
		var _lhs14 *GlobalState = _182_globalState
		_ = _lhs14
		var _lhs15 _dafny.Array = _194_v17
		_ = _lhs15
		var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_194_v17), 0))
		_ = _lhs16
		_187_v12 = _rhs23
		_192_v15 = _rhs24
		(_lhs12).ArraySet1(_rhs25, (_lhs13).Int())
		_lhs14.F17 = _rhs26
		(_lhs15).ArraySet1(_rhs27, (_lhs16).Int())
	} else {
		var _197_v20 _dafny.Array
		_ = _197_v20
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_7
		var _nw41 _dafny.Array
		_ = _nw41
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw41 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.CodePoint = (func(_198_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_199_i1 _dafny.Int) _dafny.CodePoint {
					return _198_v8
				}
			})(_183_v8)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw41 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw41).ArraySet1CodePoint(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw41).ArraySet1CodePoint(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_197_v20 = _nw41
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_197_v20), 0))
		_ = _index20
		(_197_v20).ArraySet1CodePoint(_183_v8, (_index20).Int())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_197_v20), 0))
		_ = _index21
		(_197_v20).ArraySet1CodePoint(_183_v8, (_index21).Int())
		var _200_v21 _dafny.Map
		_ = _200_v21
		_200_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_174_v0, (_dafny.Zero).Minus(_176_v2))
		var _201_v22 D13
		_ = _201_v22
		_201_v22 = Companion_D13_.Create_DC38_((_197_v20).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_197_v20), 0))).Int()), ((_200_v21).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(191))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_202_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_203_i2 _dafny.Int) _dafny.CodePoint {
				return _202_v8
			}
		})(_183_v8))), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(191))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_204_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_205_i2 _dafny.Int) _dafny.CodePoint {
				return _204_v8
			}
		})(_183_v8)))).Cardinality()))).Uint32(), (_197_v20).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_197_v20), 0))).Int()))).Cardinality())) > 0, (_176_v2).Times((_dafny.Zero).Minus(_176_v2)), _177_v3)
		_201_v22 = _201_v22
		var _206_v23 _dafny.Sequence
		_ = _206_v23
		_206_v23 = _dafny.SeqOf((_176_v2).Minus((_200_v21).Cardinality()), (_176_v2).Minus(_dafny.IntOfInt64(409)), Companion_Default___.SafeModuloInt(_176_v2, _176_v2), (_176_v2).Minus(Companion_Default___.Fm9(_176_v2, _176_v2, false, _182_globalState)))
		_206_v23 = _dafny.Companion_Sequence_.Concatenate(_206_v23, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(963))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}((func(_207_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_208_i3 _dafny.Int) _dafny.Int {
				return _207_v2
			}
		})(_176_v2))), _dafny.SeqOf(_176_v2)))
		var _209_v24 _dafny.Set
		_ = _209_v24
		_209_v24 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_180_v6, _180_v6), (Companion_Default___.SafeIndex(_176_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_180_v6, _180_v6)).Cardinality()))).Uint32(), (_197_v20).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_197_v20), 0))).Int()))).Cardinality()), _176_v2, (_dafny.IntOfUint32((_206_v23).Cardinality())).Times(_176_v2), _176_v2, (Companion_D3_.Create_DC11_(_200_v21, _176_v2, _177_v3)).Dtor_cf19())
		var _210_v25 _dafny.Sequence
		_ = _210_v25
		_210_v25 = _dafny.SeqOf(_209_v24)
		_209_v24 = (_210_v25).Select((Companion_Default___.SafeIndex(_176_v2, _dafny.IntOfUint32((_210_v25).Cardinality()))).Uint32()).(_dafny.Set)
		var _211_v26 *C2
		_ = _211_v26
		var _nw42 *C2 = New_C2_()
		_ = _nw42
		_nw42.Ctor__(_179_v5, _174_v0)
		_211_v26 = _nw42
		var _212_v27 _dafny.Array
		_ = _212_v27
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_8
		var _nw43 _dafny.Array
		_ = _nw43
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw43 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) _dafny.Sequence = (func(_213_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_214_i4 _dafny.Int) _dafny.Sequence {
					return _213_v1
				}
			})(_175_v1)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw43 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw43).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw43).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_212_v27 = _nw43
		var _215_v28 D14
		_ = _215_v28
		_215_v28 = Companion_D14_.Create_DC41_(_211_v26, _180_v6, _212_v27)
		var _source3 D14 = _215_v28
		_ = _source3
		if _source3.Is_DC41() {
			var _216___mcc_h0 *C2 = _source3.Get_().(D14_DC41).Cf72
			_ = _216___mcc_h0
			var _217___mcc_h1 _dafny.Sequence = _source3.Get_().(D14_DC41).Cf73
			_ = _217___mcc_h1
			var _218___mcc_h2 _dafny.Array = _source3.Get_().(D14_DC41).Cf74
			_ = _218___mcc_h2
			var _219_cf74 _dafny.Array = _218___mcc_h2
			_ = _219_cf74
			var _220_cf73 _dafny.Sequence = _217___mcc_h1
			_ = _220_cf73
			var _221_cf72 *C2 = _216___mcc_h0
			_ = _221_cf72
			var _222_v29 _dafny.Array
			_ = _222_v29
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(21))
			_ = _nw44
			_222_v29 = _nw44
			var _223_v31 D4
			_ = _223_v31
			_223_v31 = Companion_D4_.Create_DC14_((_221_cf72).F21(), _176_v2, func() _dafny.Set {
				var _coll20 = _dafny.NewBuilder()
				_ = _coll20
				for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(602), _dafny.IntOfInt64(158))); ; {
					_compr_20, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _224_v30 _dafny.Int
					_224_v30 = interface{}(_compr_20).(_dafny.Int)
					if ((_dafny.IntOfInt64(602)).Cmp(_224_v30) <= 0) && ((_224_v30).Cmp(_dafny.IntOfInt64(158)) < 0) {
						_coll20.Add((_224_v30).Plus(_176_v2))
					}
				}
				return _coll20.ToSet()
			}(), (_197_v20).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_197_v20), 0))).Int()))
			var _225_v32 D10
			_ = _225_v32
			_225_v32 = Companion_D10_.Create_DC28_((_211_v26).F21(), (_200_v21).Cardinality())
			var _226_v33 _dafny.Map
			_ = _226_v33
			_226_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_225_v32).Dtor_cf49(), (_211_v26).F21())
			var _227_v34 _dafny.Map
			_ = _227_v34
			_227_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_221_cf72).F21(), (func() bool {
				if (_226_v33).Contains(Companion_Default___.Fm0((_211_v26).F21(), (_197_v20).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_197_v20), 0))).Int()), (_211_v26).F21(), _182_globalState)) {
					return (_226_v33).Get(Companion_Default___.Fm0((_211_v26).F21(), (_197_v20).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_197_v20), 0))).Int()), (_211_v26).F21(), _182_globalState)).(bool)
				}
				return !((_211_v26).F21())
			})())
			var _228_v35 _dafny.Array
			_ = _228_v35
			var _nwElement0_5 bool = !((_223_v31).Dtor_cf23())
			_ = _nwElement0_5
			var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(19))
			_ = _nw45
			(_nw45).ArraySet1(_nwElement0_5, 0)
			(_nw45).ArraySet1((_221_cf72).F21(), 1)
			(_nw45).ArraySet1((_221_cf72).F21(), 2)
			(_nw45).ArraySet1(true, 3)
			(_nw45).ArraySet1(true, 4)
			(_nw45).ArraySet1((_211_v26).F21(), 5)
			(_nw45).ArraySet1((_211_v26).F21(), 6)
			(_nw45).ArraySet1(!((func() bool {
				if (_227_v34).Contains((_211_v26).F21()) {
					return (_227_v34).Get((_211_v26).F21()).(bool)
				}
				return (_211_v26).F21()
			})()), 7)
			(_nw45).ArraySet1(!((_221_cf72).F21()), 8)
			(_nw45).ArraySet1((_211_v26).F21(), 9)
			(_nw45).ArraySet1(_174_v0, 10)
			(_nw45).ArraySet1(false, 11)
			(_nw45).ArraySet1((_221_cf72).F21(), 12)
			(_nw45).ArraySet1(_174_v0, 13)
			(_nw45).ArraySet1((_211_v26).F21(), 14)
			(_nw45).ArraySet1(true, 15)
			(_nw45).ArraySet1(_174_v0, 16)
			(_nw45).ArraySet1(!((_211_v26).F21()), 17)
			(_nw45).ArraySet1(!((_211_v26).F21()), 18)
			_228_v35 = _nw45
			var _229_v36 _dafny.MultiSet
			_ = _229_v36
			_229_v36 = _dafny.MultiSetOf(_228_v35, _228_v35, _228_v35)
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_222_v29), 0))
			_ = _index22
			(_222_v29).ArraySet1(_229_v36, (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_222_v29), 0))
			_ = _index23
			(_222_v29).ArraySet1(_229_v36, (_index23).Int())
			var _230_v37 D3
			_ = _230_v37
			_230_v37 = Companion_D3_.Create_DC10_(_176_v2)
			var _231_v38 _dafny.MultiSet
			_ = _231_v38
			_231_v38 = _dafny.MultiSetOf(_230_v37)
			var _232_v39 _dafny.MultiSet
			_ = _232_v39
			_232_v39 = _dafny.MultiSetOf(!(true), (_211_v26).F21(), false, !(_231_v38).Equals(_231_v38))
			_232_v39 = _232_v39
			var _233_v40 bool
			_ = _233_v40
			var _234_v41 bool
			_ = _234_v41
			var _out1 bool
			_ = _out1
			var _out2 bool
			_ = _out2
			_out1, _out2 = (_221_cf72).M1((_223_v31).Dtor_cf23(), (func() _dafny.Int {
				if (_200_v21).Contains((_211_v26).F21()) {
					return (_200_v21).Get((_211_v26).F21()).(_dafny.Int)
				}
				return _dafny.IntOfInt64(-498)
			})(), false, (Companion_D1_.Create_DC4_(_dafny.IntOfInt64(812), (_211_v26).F21(), _228_v35)).Dtor_cf4(), _182_globalState)
			_233_v40 = _out1
			_234_v41 = _out2
			var _235_v42 _dafny.Array
			_ = _235_v42
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_9
			var _nw46 _dafny.Array
			_ = _nw46
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw46 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.MultiSet = (func(_236_cf72 *C2) func(_dafny.Int) _dafny.MultiSet {
					return func(_237_i5 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf((_236_cf72).F21())
					}
				})(_221_cf72)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw46 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw46).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw46).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_235_v42 = _nw46
			var _rhs28 bool = _234_v41
			_ = _rhs28
			var _rhs29 _dafny.CodePoint = _dafny.CodePoint('j')
			_ = _rhs29
			var _rhs30 _dafny.Array = _235_v42
			_ = _rhs30
			var _lhs17 *GlobalState = _182_globalState
			_ = _lhs17
			_lhs17.F15 = _rhs28
			_183_v8 = _rhs29
			_235_v42 = _rhs30
		} else {
			var _238___mcc_h3 _dafny.Array = _source3.Get_().(D14_DC40).Cf71
			_ = _238___mcc_h3
			var _239_cf71 _dafny.Array = _238___mcc_h3
			_ = _239_cf71
			var _240_v43 T1
			_ = _240_v43
			var _nw47 *C2 = New_C2_()
			_ = _nw47
			_nw47.Ctor__(_179_v5, (_211_v26).F21())
			_240_v43 = _nw47
			_240_v43 = (func() T1 {
				if false {
					return _240_v43
				}
				return _240_v43
			})()
			var _241_v44 _dafny.Array
			_ = _241_v44
			var _nw48 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw48
			_241_v44 = _nw48
			_241_v44 = _241_v44
			var _242_v45 bool
			_ = _242_v45
			var _243_v46 bool
			_ = _243_v46
			var _244_v47 bool
			_ = _244_v47
			var _245_v48 _dafny.Sequence
			_ = _245_v48
			var _out3 bool
			_ = _out3
			var _out4 bool
			_ = _out4
			var _out5 bool
			_ = _out5
			var _out6 _dafny.Sequence
			_ = _out6
			_out3, _out4, _out5, _out6 = (_211_v26).M6(_176_v2, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-465))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_246_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_247_i6 _dafny.Int) _dafny.CodePoint {
					return _246_v8
				}
			})(_183_v8))), _180_v6), (func() bool {
				if (_211_v26).F21() {
					return (_211_v26).F21()
				}
				return (_211_v26).F21()
			})(), _182_globalState)
			_242_v45 = _out3
			_243_v46 = _out4
			_244_v47 = _out5
			_245_v48 = _out6
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_177_v3), 0))
			_ = _index24
			(_177_v3).ArraySet1(_176_v2, (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_177_v3), 0))
			_ = _index25
			(_177_v3).ArraySet1(_176_v2, (_index25).Int())
		}
	}
	var _248_v49 D12
	_ = _248_v49
	_248_v49 = Companion_D12_.Create_DC35_(Companion_D12_.Create_DC34_())
	var _249_v50 D12
	_ = _249_v50
	_249_v50 = Companion_D12_.Create_DC35_(_248_v49)
	var _250_v51 _dafny.Map
	_ = _250_v51
	_250_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_249_v50, _176_v2)
	_250_v51 = (_250_v51).Update(_249_v50, _176_v2)
	var _251_v52 _dafny.MultiSet
	_ = _251_v52
	_251_v52 = _dafny.MultiSetOf(!(_174_v0), _174_v0)
	var _252_v53 _dafny.Sequence
	_ = _252_v53
	_252_v53 = _dafny.SeqOf(Companion_D10_.Create_DC27_(_251_v52))
	var _253_v54 _dafny.Sequence
	_ = _253_v54
	_253_v54 = _dafny.SeqOf(_176_v2)
	var _hi1 _dafny.Int = (_253_v54).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.IntOfUint32((_253_v54).Cardinality()))).Uint32()).(_dafny.Int)
	_ = _hi1
	for _254_i7 := (_dafny.Zero).Minus(_dafny.IntOfUint32((_252_v53).Cardinality())); _254_i7.Cmp(_hi1) < 0; _254_i7 = _254_i7.Plus(_dafny.One) {
		var _255_v55 _dafny.Map
		_ = _255_v55
		_255_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_174_v0, _183_v8)
		(_182_globalState).F11 = (func() bool {
			if _174_v0 {
				return Companion_Default___.Fm0(_174_v0, Companion_Default___.Fm12((_255_v55).Cardinality(), _180_v6, _180_v6, _182_globalState), _174_v0, _182_globalState)
			}
			return _174_v0
		})()
		var _256_v56 bool
		_ = _256_v56
		var _257_v57 bool
		_ = _257_v57
		var _258_v58 _dafny.Int
		_ = _258_v58
		var _out7 bool
		_ = _out7
		var _out8 bool
		_ = _out8
		var _out9 _dafny.Int
		_ = _out9
		_out7, _out8, _out9 = Companion_Default___.M9(_174_v0, (func() bool {
			if _174_v0 {
				return false
			}
			return false
		})(), _182_globalState)
		_256_v56 = _out7
		_257_v57 = _out8
		_258_v58 = _out9
		_183_v8 = (_180_v6).Select((Companion_Default___.SafeIndex(_258_v58, _dafny.IntOfUint32((_180_v6).Cardinality()))).Uint32()).(_dafny.CodePoint)
		var _259_v59 _dafny.Array
		_ = _259_v59
		var _nw49 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw49
		_259_v59 = _nw49
		var _260_v60 _dafny.Sequence
		_ = _260_v60
		_260_v60 = _dafny.SeqOf(_259_v59)
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_260_v60, _dafny.Companion_Sequence_.Concatenate(_260_v60, _260_v60)) {
			var _261_v61 _dafny.Map
			_ = _261_v61
			_261_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v8, Companion_Default___.Fm24(_256_v56, _258_v58, _182_globalState))
			var _262_v62 _dafny.Map
			_ = _262_v62
			_262_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_257_v57, _254_i7)
			var _263_v63 D3
			_ = _263_v63
			_263_v63 = Companion_D3_.Create_DC11_(_262_v62, _dafny.IntOfInt64(-703), _177_v3)
			var _264_v64 _dafny.Array
			_ = _264_v64
			var _nwElement0_6 _dafny.Map = (func() _dafny.Map {
				if (_261_v61).Contains(_183_v8) {
					return (_261_v61).Get(_183_v8).(_dafny.Map)
				}
				return (_263_v63).Dtor_cf18()
			})()
			_ = _nwElement0_6
			var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(4))
			_ = _nw50
			(_nw50).ArraySet1(_nwElement0_6, 0)
			(_nw50).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_174_v0, _176_v2), 1)
			(_nw50).ArraySet1(_262_v62, 2)
			(_nw50).ArraySet1(_262_v62, 3)
			_264_v64 = _nw50
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_264_v64), 0))
			_ = _index26
			(_264_v64).ArraySet1(_262_v62, (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_264_v64), 0))
			_ = _index27
			(_264_v64).ArraySet1(_262_v62, (_index27).Int())
			var _265_v65 *C0
			_ = _265_v65
			var _nw51 *C0 = New_C0_()
			_ = _nw51
			_nw51.Ctor__(Companion_Default___.Fm0(_174_v0, _183_v8, _257_v57, _182_globalState))
			_265_v65 = _nw51
			var _266_v66 _dafny.Map
			_ = _266_v66
			_266_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_251_v52, _180_v6)
			var _267_v67 bool
			_ = _267_v67
			var _268_v68 bool
			_ = _268_v68
			var _269_v69 _dafny.Int
			_ = _269_v69
			var _out10 bool
			_ = _out10
			var _out11 bool
			_ = _out11
			var _out12 _dafny.Int
			_ = _out12
			_out10, _out11, _out12 = Companion_Default___.M9((_265_v65).Fm15(_256_v56, _257_v57, (_dafny.Zero).Minus(_258_v58), _262_v62, _182_globalState), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_266_v66).Contains(_dafny.MultiSetOf(!(_257_v57))) {
					return (_266_v66).Get(_dafny.MultiSetOf(!(_257_v57))).(_dafny.Sequence)
				}
				return _180_v6
			})(), (Companion_Default___.SafeIndex(_254_i7, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_266_v66).Contains(_dafny.MultiSetOf(!(_257_v57))) {
					return (_266_v66).Get(_dafny.MultiSetOf(!(_257_v57))).(_dafny.Sequence)
				}
				return _180_v6
			})()).Cardinality()))).Uint32(), _183_v8), _180_v6), _182_globalState)
			_267_v67 = _out10
			_268_v68 = _out11
			_269_v69 = _out12
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_177_v3), 0))
			_ = _index28
			(_177_v3).ArraySet1((_258_v58).Times(_258_v58), (_index28).Int())
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_177_v3), 0))
			_ = _index29
			(_177_v3).ArraySet1(_176_v2, (_index29).Int())
			var _270_v70 bool
			_ = _270_v70
			var _271_v71 bool
			_ = _271_v71
			var _272_v72 _dafny.Int
			_ = _272_v72
			var _out13 bool
			_ = _out13
			var _out14 bool
			_ = _out14
			var _out15 _dafny.Int
			_ = _out15
			_out13, _out14, _out15 = Companion_Default___.M9((_265_v65).F23(), _dafny.Companion_Sequence_.IsPrefixOf(_180_v6, _180_v6), _182_globalState)
			_270_v70 = _out13
			_271_v71 = _out14
			_272_v72 = _out15
		} else {
			_176_v2 = (_254_i7).Plus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(873), _254_i7))
			_174_v0 = true
			var _273_v73 _dafny.Map
			_ = _273_v73
			_273_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_254_i7).Minus((_dafny.Zero).Minus(_176_v2)), ((_dafny.Zero).Minus(_258_v58)).Cmp(_258_v58) < 0)
			_273_v73 = _273_v73
			(_182_globalState).F15 = _256_v56
			var _274_v74 bool
			_ = _274_v74
			var _275_v75 bool
			_ = _275_v75
			var _276_v76 _dafny.Int
			_ = _276_v76
			var _out16 bool
			_ = _out16
			var _out17 bool
			_ = _out17
			var _out18 _dafny.Int
			_ = _out18
			_out16, _out17, _out18 = Companion_Default___.M9(_174_v0, _256_v56, _182_globalState)
			_274_v74 = _out16
			_275_v75 = _out17
			_276_v76 = _out18
		}
	}
	var _277_v77 *C0
	_ = _277_v77
	var _nw52 *C0 = New_C0_()
	_ = _nw52
	_nw52.Ctor__(_174_v0)
	_277_v77 = _nw52
	var _278_v78 _dafny.MultiSet
	_ = _278_v78
	_278_v78 = _dafny.MultiSetOf(_277_v77)
	var _rhs31 bool = (_174_v0) == ((_277_v77).F23())
	_ = _rhs31
	var _rhs32 _dafny.MultiSet = _278_v78
	_ = _rhs32
	var _lhs18 *GlobalState = _182_globalState
	_ = _lhs18
	_lhs18.F11 = _rhs31
	_278_v78 = _rhs32
	var _279_v79 _dafny.Map
	_ = _279_v79
	_279_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v77).F23(), _176_v2)
	if (_277_v77).Fm15(_174_v0, _174_v0, _176_v2, _279_v79, _182_globalState) {
		var _280_v80 _dafny.Array
		_ = _280_v80
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_10
		var _nw53 _dafny.Array
		_ = _nw53
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw53 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.CodePoint = (func(_281_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_282_i8 _dafny.Int) _dafny.CodePoint {
					return _281_v8
				}
			})(_183_v8)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw53 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw53).ArraySet1CodePoint(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw53).ArraySet1CodePoint(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_280_v80 = _nw53
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_280_v80), 0))
		_ = _index30
		(_280_v80).ArraySet1CodePoint(_183_v8, (_index30).Int())
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_280_v80), 0))
		_ = _index31
		(_280_v80).ArraySet1CodePoint(_183_v8, (_index31).Int())
		var _283_v81 *C2
		_ = _283_v81
		var _nw54 *C2 = New_C2_()
		_ = _nw54
		_nw54.Ctor__(_179_v5, (_277_v77).F23())
		_283_v81 = _nw54
		var _284_v82 bool
		_ = _284_v82
		var _285_v83 bool
		_ = _285_v83
		var _286_v84 bool
		_ = _286_v84
		var _287_v85 _dafny.Sequence
		_ = _287_v85
		var _out19 bool
		_ = _out19
		var _out20 bool
		_ = _out20
		var _out21 bool
		_ = _out21
		var _out22 _dafny.Sequence
		_ = _out22
		_out19, _out20, _out21, _out22 = (_283_v81).M6(Companion_Default___.SafeDivisionInt(_176_v2, _176_v2), _180_v6, (_277_v77).F23(), _182_globalState)
		_284_v82 = _out19
		_285_v83 = _out20
		_286_v84 = _out21
		_287_v85 = _out22
		var _288_v86 _dafny.Map
		_ = _288_v86
		_288_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), (_dafny.Zero).Minus(_176_v2))
		if ((_288_v86).Contains(_183_v8)) || ((_285_v83) || (_284_v82)) {
			var _289_v87 D12
			_ = _289_v87
			_289_v87 = Companion_D12_.Create_DC34_()
			var _290_v88 _dafny.Map
			_ = _290_v88
			_290_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_289_v87, _277_v77)
			_290_v88 = (_290_v88).Update(_289_v87, _277_v77)
			var _291_v89 _dafny.Set
			_ = _291_v89
			_291_v89 = _dafny.SetOf(_253_v54, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-160))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_292_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_293_i9 _dafny.Int) _dafny.Int {
					return _292_v2
				}
			})(_176_v2))))
			var _294_v90 _dafny.Sequence
			_ = _294_v90
			_294_v90 = _dafny.SeqOf(_291_v89, _291_v89, _291_v89)
			(_182_globalState).F17 = (_291_v89).IsProperSubsetOf((_294_v90).Select((Companion_Default___.SafeIndex(_176_v2, _dafny.IntOfUint32((_294_v90).Cardinality()))).Uint32()).(_dafny.Set))
			_176_v2 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_280_v80).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_280_v80), 0))).Int()), _251_v52)).Cardinality()
			var _295_v91 _dafny.Array
			_ = _295_v91
			var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw55
			_295_v91 = _nw55
			var _296_v92 _dafny.MultiSet
			_ = _296_v92
			_296_v92 = _dafny.MultiSetOf(_176_v2, _176_v2)
			var _297_v93 _dafny.Sequence
			_ = _297_v93
			_297_v93 = _dafny.SeqOf(_253_v54)
			var _298_v94 D8
			_ = _298_v94
			_298_v94 = Companion_D8_.Create_DC23_(_296_v92)
			var _pat_let_tv2 = _296_v92
			_ = _pat_let_tv2
			var _299_v95 _dafny.Array
			_ = _299_v95
			var _nwElement0_7 _dafny.MultiSet = _296_v92
			_ = _nwElement0_7
			var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(5))
			_ = _nw56
			(_nw56).ArraySet1(_nwElement0_7, 0)
			(_nw56).ArraySet1(_dafny.MultiSetFromSeq((_297_v93).Select((Companion_Default___.SafeIndex(_176_v2, _dafny.IntOfUint32((_297_v93).Cardinality()))).Uint32()).(_dafny.Sequence)), 1)
			(_nw56).ArraySet1((func(_pat_let3_0 D8) D8 {
				return func(_300_dt__update__tmp_h0 D8) D8 {
					return func(_pat_let4_0 _dafny.MultiSet) D8 {
						return func(_301_dt__update_hcf41_h0 _dafny.MultiSet) D8 {
							return Companion_D8_.Create_DC23_(_301_dt__update_hcf41_h0)
						}(_pat_let4_0)
					}(_pat_let_tv2)
				}(_pat_let3_0)
			}(_298_v94)).Dtor_cf41(), 2)
			(_nw56).ArraySet1(_296_v92, 3)
			(_nw56).ArraySet1(_296_v92, 4)
			_299_v95 = _nw56
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_295_v91), 0))
			_ = _index32
			(_295_v91).ArraySet1(_299_v95, (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_295_v91), 0))
			_ = _index33
			(_295_v91).ArraySet1(_299_v95, (_index33).Int())
			(_182_globalState).F15 = _284_v82
		} else {
			var _302_v96 bool
			_ = _302_v96
			var _out23 bool
			_ = _out23
			_out23 = (_283_v81).M5(!(!(true) || ((_277_v77).F23())), _175_v1, _182_globalState)
			_302_v96 = _out23
			var _303_v97 D9
			_ = _303_v97
			_303_v97 = Companion_D9_.Create_DC25_(_176_v2, (_283_v81).F21(), _176_v2, _277_v77)
			var _304_v98 _dafny.Map
			_ = _304_v98
			_304_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v2, _303_v97)
			var _305_v99 *C4
			_ = _305_v99
			var _nw57 *C4 = New_C4_()
			_ = _nw57
			_nw57.Ctor__((_283_v81).F21())
			_305_v99 = _nw57
			var _306_v100 _dafny.Map
			_ = _306_v100
			_306_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_305_v99, _285_v83)
			var _307_v101 _dafny.Array
			_ = _307_v101
			var _nw58 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw58
			_307_v101 = _nw58
			var _308_v102 _dafny.Map
			_ = _308_v102
			_308_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_307_v101, (_277_v77).F23())
			var _309_v103 bool
			_ = _309_v103
			var _310_v104 bool
			_ = _310_v104
			var _311_v105 bool
			_ = _311_v105
			var _312_v106 _dafny.Sequence
			_ = _312_v106
			var _out24 bool
			_ = _out24
			var _out25 bool
			_ = _out25
			var _out26 bool
			_ = _out26
			var _out27 _dafny.Sequence
			_ = _out27
			_out24, _out25, _out26, _out27 = (_283_v81).M6(((_304_v98).Cardinality()).Plus((_306_v100).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_287_v85, _dafny.UnicodeSeqOfUtf8Bytes("yoiqoa")), (Companion_D1_.Create_DC4_((_308_v102).Cardinality(), (_277_v77).Fm15((_283_v81).F21(), (_283_v81).F21(), _dafny.IntOfInt64(-511), _279_v79, _182_globalState), _307_v101)).Dtor_cf5(), _182_globalState)
			_309_v103 = _out24
			_310_v104 = _out25
			_311_v105 = _out26
			_312_v106 = _out27
			var _313_v107 T1
			_ = _313_v107
			var _nw59 *C2 = New_C2_()
			_ = _nw59
			_nw59.Ctor__(_283_v81.F20, false)
			_313_v107 = _nw59
			var _314_v108 *C2
			_ = _314_v108
			var _nw60 *C2 = New_C2_()
			_ = _nw60
			_nw60.Ctor__(_179_v5, (_dafny.MultiSetOf(_313_v107, _313_v107)).Contains(_313_v107))
			_314_v108 = _nw60
			_284_v82 = Companion_Default___.Fm0(_285_v83, _183_v8, _302_v96, _182_globalState)
			var _315_v109 *C5
			_ = _315_v109
			var _nw61 *C5 = New_C5_()
			_ = _nw61
			_nw61.Ctor__(Companion_Default___.Fm5(_dafny.IntOfUint32((_dafny.SeqOf((_283_v81).F21())).Cardinality()), _182_globalState))
			_315_v109 = _nw61
		}
		var _316_v110 _dafny.Set
		_ = _316_v110
		_316_v110 = _dafny.SetOf(_285_v83, _285_v83)
		var _317_v111 _dafny.Sequence
		_ = _317_v111
		_317_v111 = _dafny.SeqOf(_316_v110)
		var _318_v112 bool
		_ = _318_v112
		var _319_v113 bool
		_ = _319_v113
		var _320_v114 bool
		_ = _320_v114
		var _321_v115 _dafny.Sequence
		_ = _321_v115
		var _out28 bool
		_ = _out28
		var _out29 bool
		_ = _out29
		var _out30 bool
		_ = _out30
		var _out31 _dafny.Sequence
		_ = _out31
		_out28, _out29, _out30, _out31 = (_283_v81).M6(_dafny.IntOfUint32((_317_v111).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-938))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}(func(_322_i10 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})), (_283_v81).F21(), _182_globalState)
		_318_v112 = _out28
		_319_v113 = _out29
		_320_v114 = _out30
		_321_v115 = _out31
	} else {
		var _323_v116 _dafny.Array
		_ = _323_v116
		var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
		_ = _nw62
		_323_v116 = _nw62
		_323_v116 = _323_v116
		(_182_globalState).F17 = !((_277_v77).F23())
		var _324_v117 _dafny.Map
		_ = _324_v117
		_324_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v2, _279_v79)
		_324_v117 = (_324_v117).Update((_176_v2).Plus(_176_v2), (_279_v79).Merge(_279_v79))
		_174_v0 = (_277_v77).F23()
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(480), _dafny.ArrayLen((_177_v3), 0))
		_ = _index34
		(_177_v3).ArraySet1(_176_v2, (_index34).Int())
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(480), _dafny.ArrayLen((_177_v3), 0))
		_ = _index35
		(_177_v3).ArraySet1(_176_v2, (_index35).Int())
	}
	var _325_v118 bool
	_ = _325_v118
	var _326_v119 bool
	_ = _326_v119
	var _327_v120 _dafny.Int
	_ = _327_v120
	var _out32 bool
	_ = _out32
	var _out33 bool
	_ = _out33
	var _out34 _dafny.Int
	_ = _out34
	_out32, _out33, _out34 = Companion_Default___.M9(_174_v0, _174_v0, _182_globalState)
	_325_v118 = _out32
	_326_v119 = _out33
	_327_v120 = _out34
	var _328_v121 _dafny.Map
	_ = _328_v121
	_328_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _325_v118)
	var _329_v122 D13
	_ = _329_v122
	_329_v122 = Companion_D13_.Create_DC36_(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_328_v121, _328_v121, _328_v121, _328_v121, _328_v121), (Companion_Default___.SafeIndex(_176_v2, _dafny.IntOfUint32((_dafny.SeqOf(_328_v121, _328_v121, _328_v121, _328_v121, _328_v121)).Cardinality()))).Uint32(), _328_v121))
	var _pat_let_tv3 = _183_v8
	_ = _pat_let_tv3
	var _pat_let_tv4 = _183_v8
	_ = _pat_let_tv4
	var _pat_let_tv5 = _277_v77
	_ = _pat_let_tv5
	var _pat_let_tv6 = _327_v120
	_ = _pat_let_tv6
	var _pat_let_tv7 = _327_v120
	_ = _pat_let_tv7
	var _pat_let_tv8 = _183_v8
	_ = _pat_let_tv8
	var _pat_let_tv9 = _183_v8
	_ = _pat_let_tv9
	_183_v8 = func(_source4 D13) _dafny.CodePoint {
		if _source4.Is_DC37() {
			var _330___mcc_h4 _dafny.Int = _source4.Get_().(D13_DC37).Cf64
			_ = _330___mcc_h4
			var _331___mcc_h5 _dafny.Sequence = _source4.Get_().(D13_DC37).Cf65
			_ = _331___mcc_h5
			var _332_cf65 _dafny.Sequence = _331___mcc_h5
			_ = _332_cf65
			var _333_cf64 _dafny.Int = _330___mcc_h4
			_ = _333_cf64
			return _pat_let_tv3
		} else if _source4.Is_DC38() {
			var _334___mcc_h6 _dafny.CodePoint = _source4.Get_().(D13_DC38).Cf66
			_ = _334___mcc_h6
			var _335___mcc_h7 bool = _source4.Get_().(D13_DC38).Cf67
			_ = _335___mcc_h7
			var _336___mcc_h8 _dafny.Int = _source4.Get_().(D13_DC38).Cf68
			_ = _336___mcc_h8
			var _337___mcc_h9 _dafny.Array = _source4.Get_().(D13_DC38).Cf69
			_ = _337___mcc_h9
			var _338_cf69 _dafny.Array = _337___mcc_h9
			_ = _338_cf69
			var _339_cf68 _dafny.Int = _336___mcc_h8
			_ = _339_cf68
			var _340_cf67 bool = _335___mcc_h7
			_ = _340_cf67
			var _341_cf66 _dafny.CodePoint = _334___mcc_h6
			_ = _341_cf66
			return _pat_let_tv4
		} else if _source4.Is_DC36() {
			var _342___mcc_h10 _dafny.Sequence = _source4.Get_().(D13_DC36).Cf63
			_ = _342___mcc_h10
			var _343_cf63 _dafny.Sequence = _342___mcc_h10
			_ = _343_cf63
			return (Companion_D4_.Create_DC14_((_pat_let_tv5).F23(), _pat_let_tv6, _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mdb")).Cardinality()), _pat_let_tv7), _pat_let_tv8)).Dtor_cf26()
		} else {
			var _344___mcc_h11 D13 = _source4.Get_().(D13_DC39).Cf70
			_ = _344___mcc_h11
			var _345_cf70 D13 = _344___mcc_h11
			_ = _345_cf70
			return _pat_let_tv9
		}
	}(_329_v122)
	var _hi2 _dafny.Int = _176_v2
	_ = _hi2
	for _346_i11 := _327_v120; _346_i11.Cmp(_hi2) < 0; _346_i11 = _346_i11.Plus(_dafny.One) {
		var _347_v123 _dafny.Array
		_ = _347_v123
		var _nw63 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
		_ = _nw63
		_347_v123 = _nw63
		var _348_v124 *C2
		_ = _348_v124
		var _nw64 *C2 = New_C2_()
		_ = _nw64
		_nw64.Ctor__(_179_v5, true)
		_348_v124 = _nw64
		var _349_v125 _dafny.Map
		_ = _349_v125
		_349_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_i11, _348_v124)
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_347_v123), 0))
		_ = _index36
		(_347_v123).ArraySet1((func() *C2 {
			if (_349_v125).Contains(_176_v2) {
				return (_349_v125).Get(_176_v2).(*C2)
			}
			return _348_v124
		})(), (_index36).Int())
		var _350_v126 _dafny.Sequence
		_ = _350_v126
		_350_v126 = _dafny.SeqOf(_348_v124, _348_v124, _348_v124, _348_v124)
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_347_v123), 0))
		_ = _index37
		(_347_v123).ArraySet1((_350_v126).Select((Companion_Default___.SafeIndex(_346_i11, _dafny.IntOfUint32((_350_v126).Cardinality()))).Uint32()).(*C2), (_index37).Int())
		var _351_v127 _dafny.Sequence
		_ = _351_v127
		_351_v127 = _dafny.SeqOf(_329_v122)
		var _352_v128 _dafny.Sequence
		_ = _352_v128
		_352_v128 = _dafny.SeqOf(_351_v127)
		if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(112))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}((func(_353_v122 D13) func(_dafny.Int) _dafny.Sequence {
			return func(_354_i12 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(856))).Uint32(), func(coer24 func(_dafny.Int) D13) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}((func(_355_v122 D13) func(_dafny.Int) D13 {
					return func(_356_i13 _dafny.Int) D13 {
						return _355_v122
					}
				})(_353_v122)))
			}
		})(_329_v122))), _352_v128), _352_v128) {
			_174_v0 = (true) || (_325_v118)
			var _357_v129 _dafny.Set
			_ = _357_v129
			_357_v129 = _dafny.SetOf(_176_v2)
			_357_v129 = (_dafny.SetOf(_dafny.IntOfUint32((_180_v6).Cardinality()))).Intersection(_357_v129)
			_327_v120 = _dafny.IntOfInt64(371)
			var _358_v130 bool
			_ = _358_v130
			var _359_v131 bool
			_ = _359_v131
			var _360_v132 bool
			_ = _360_v132
			var _361_v133 _dafny.Sequence
			_ = _361_v133
			var _out35 bool
			_ = _out35
			var _out36 bool
			_ = _out36
			var _out37 bool
			_ = _out37
			var _out38 _dafny.Sequence
			_ = _out38
			_out35, _out36, _out37, _out38 = (_348_v124).M6(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_180_v6).Cardinality()), (_279_v79).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_180_v6, _dafny.UnicodeSeqOfUtf8Bytes("cqmm")), (_277_v77).F23(), _182_globalState)
			_358_v130 = _out35
			_359_v131 = _out36
			_360_v132 = _out37
			_361_v133 = _out38
			var _362_v134 D9
			_ = _362_v134
			_362_v134 = Companion_D9_.Create_DC24_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-737))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}((func(_363_v120 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_364_i14 _dafny.Int) _dafny.Int {
					return _363_v120
				}
			})(_327_v120))))
			_253_v54 = (_362_v134).Dtor_cf42()
		} else {
			var _365_v135 bool
			_ = _365_v135
			var _366_v136 bool
			_ = _366_v136
			var _367_v137 _dafny.Int
			_ = _367_v137
			var _out39 bool
			_ = _out39
			var _out40 bool
			_ = _out40
			var _out41 _dafny.Int
			_ = _out41
			_out39, _out40, _out41 = Companion_Default___.M9((_348_v124).F21(), (_277_v77).F23(), _182_globalState)
			_365_v135 = _out39
			_366_v136 = _out40
			_367_v137 = _out41
			_180_v6 = _dafny.Companion_Sequence_.Update(_180_v6, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_367_v137), _dafny.IntOfUint32((_180_v6).Cardinality()))).Uint32(), _183_v8)
			var _368_v138 _dafny.Set
			_ = _368_v138
			_368_v138 = _dafny.SetOf(_180_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-434))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_369_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_370_i15 _dafny.Int) _dafny.CodePoint {
					return _369_v8
				}
			})(_183_v8))))
			_327_v120 = (func() _dafny.Int {
				if (_277_v77).F23() {
					return _346_i11
				}
				return Companion_Default___.SafeModuloInt((_368_v138).Cardinality(), _176_v2)
			})()
			(_182_globalState).F15 = _dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
				if _365_v135 {
					return _180_v6
				}
				return _180_v6
			})(), _dafny.Companion_Sequence_.Concatenate(_180_v6, _180_v6))
			var _371_v139 *C0
			_ = _371_v139
			var _nw65 *C0 = New_C0_()
			_ = _nw65
			_nw65.Ctor__(true)
			_371_v139 = _nw65
		}
		var _372_v140 T1
		_ = _372_v140
		var _nw66 *C4 = New_C4_()
		_ = _nw66
		_nw66.Ctor__(_325_v118)
		_372_v140 = _nw66
		var _373_v141 _dafny.Set
		_ = _373_v141
		_373_v141 = _dafny.SetOf(_372_v140, _372_v140, _372_v140)
		var _374_v142 D12
		_ = _374_v142
		_374_v142 = Companion_D12_.Create_DC31_(_373_v141)
		var _pat_let_tv10 = _373_v141
		_ = _pat_let_tv10
		var _source5 D12 = func(_pat_let5_0 D12) D12 {
			return func(_375_dt__update__tmp_h1 D12) D12 {
				return func(_pat_let6_0 _dafny.Set) D12 {
					return func(_376_dt__update_hcf55_h0 _dafny.Set) D12 {
						return Companion_D12_.Create_DC31_(_376_dt__update_hcf55_h0)
					}(_pat_let6_0)
				}(_pat_let_tv10)
			}(_pat_let5_0)
		}(_374_v142)
		_ = _source5
		if _source5.Is_DC32() {
			var _377___mcc_h12 _dafny.Int = _source5.Get_().(D12_DC32).Cf56
			_ = _377___mcc_h12
			var _378_cf56 _dafny.Int = _377___mcc_h12
			_ = _378_cf56
			var _379_v143 _dafny.Sequence
			_ = _379_v143
			_379_v143 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_348_v124).F21(), _325_v118), _328_v121, _328_v121)
			var _380_v144 D13
			_ = _380_v144
			_380_v144 = Companion_D13_.Create_DC36_(_379_v143)
			var _381_v145 D13
			_ = _381_v145
			_381_v145 = Companion_D13_.Create_DC39_(_380_v144)
			_381_v145 = _381_v145
			var _382_v146 D13
			_ = _382_v146
			_382_v146 = Companion_D13_.Create_DC38_(_183_v8, (_277_v77).F23(), _346_i11, _177_v3)
			var _383_v147 bool
			_ = _383_v147
			var _384_v148 bool
			_ = _384_v148
			var _385_v149 _dafny.Int
			_ = _385_v149
			var _out42 bool
			_ = _out42
			var _out43 bool
			_ = _out43
			var _out44 _dafny.Int
			_ = _out44
			_out42, _out43, _out44 = Companion_Default___.M9((func() bool {
				if (_277_v77).F23() {
					return _174_v0
				}
				return (_348_v124).F21()
			})(), (_382_v146).Dtor_cf67(), _182_globalState)
			_383_v147 = _out42
			_384_v148 = _out43
			_385_v149 = _out44
			var _386_v150 _dafny.Set
			_ = _386_v150
			_386_v150 = _dafny.SetOf(_180_v6)
			_325_v118 = (_386_v150).IsSubsetOf(_386_v150)
			var _387_v151 *C1
			_ = _387_v151
			var _nw67 *C1 = New_C1_()
			_ = _nw67
			_nw67.Ctor__(_177_v3)
			_387_v151 = _nw67
		} else if _source5.Is_DC33() {
			var _388___mcc_h13 _dafny.Int = _source5.Get_().(D12_DC33).Cf57
			_ = _388___mcc_h13
			var _389___mcc_h14 bool = _source5.Get_().(D12_DC33).Cf58
			_ = _389___mcc_h14
			var _390___mcc_h15 _dafny.MultiSet = _source5.Get_().(D12_DC33).Cf59
			_ = _390___mcc_h15
			var _391___mcc_h16 bool = _source5.Get_().(D12_DC33).Cf60
			_ = _391___mcc_h16
			var _392___mcc_h17 bool = _source5.Get_().(D12_DC33).Cf61
			_ = _392___mcc_h17
			var _393_cf61 bool = _392___mcc_h17
			_ = _393_cf61
			var _394_cf60 bool = _391___mcc_h16
			_ = _394_cf60
			var _395_cf59 _dafny.MultiSet = _390___mcc_h15
			_ = _395_cf59
			var _396_cf58 bool = _389___mcc_h14
			_ = _396_cf58
			var _397_cf57 _dafny.Int = _388___mcc_h13
			_ = _397_cf57
			var _398_v152 _dafny.Map
			_ = _398_v152
			_398_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(727))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_399_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_400_i16 _dafny.Int) _dafny.CodePoint {
					return _399_v8
				}
			})(_183_v8))), false)
			_397_cf57 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(Companion_Default___.Fm3(_dafny.IntOfInt64(-604), _327_v120, !((func() bool {
				if (_398_v152).Contains(_180_v6) {
					return (_398_v152).Get(_180_v6).(bool)
				}
				return false
			})()), _182_globalState)), _176_v2)
			var _401_v153 *C4
			_ = _401_v153
			var _nw68 *C4 = New_C4_()
			_ = _nw68
			_nw68.Ctor__((_348_v124).F21())
			_401_v153 = _nw68
			_372_v140 = _372_v140
			(_182_globalState).F11 = (_277_v77).F23()
		} else if _source5.Is_DC34() {
			var _402_v154 _dafny.Map
			_ = _402_v154
			_402_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v2, _174_v0)
			_402_v154 = (_402_v154).Update(_dafny.IntOfInt64(918), ((_dafny.MultiSetFromSeq(_175_v1)).Cardinality()).Cmp(_346_i11) < 0)
			var _403_v155 D13
			_ = _403_v155
			_403_v155 = Companion_D13_.Create_DC37_(_dafny.IntOfInt64(-142), _dafny.SeqOf(_253_v54))
			var _404_v156 _dafny.Set
			_ = _404_v156
			_404_v156 = _dafny.SetOf(_176_v2, _176_v2, _dafny.IntOfInt64(368))
			var _405_v157 D13
			_ = _405_v157
			_405_v157 = Companion_D13_.Create_DC38_(_183_v8, _326_v119, _346_i11, _177_v3)
			var _406_v158 _dafny.Array
			_ = _406_v158
			var _nwElement0_8 _dafny.CodePoint = _183_v8
			_ = _nwElement0_8
			var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(26))
			_ = _nw69
			(_nw69).ArraySet1CodePoint(_nwElement0_8, 0)
			(_nw69).ArraySet1CodePoint((func() _dafny.CodePoint {
				if !(true) {
					return (_180_v6).Select((Companion_Default___.SafeIndex((_404_v156).Cardinality(), _dafny.IntOfUint32((_180_v6).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
				return _183_v8
			})(), 1)
			(_nw69).ArraySet1CodePoint(_183_v8, 2)
			(_nw69).ArraySet1CodePoint(_183_v8, 3)
			(_nw69).ArraySet1CodePoint(_183_v8, 4)
			(_nw69).ArraySet1CodePoint((_405_v157).Dtor_cf66(), 5)
			(_nw69).ArraySet1CodePoint(_183_v8, 6)
			(_nw69).ArraySet1CodePoint(_183_v8, 7)
			(_nw69).ArraySet1CodePoint(_183_v8, 8)
			(_nw69).ArraySet1CodePoint(_183_v8, 9)
			(_nw69).ArraySet1CodePoint(_183_v8, 10)
			(_nw69).ArraySet1CodePoint(_183_v8, 11)
			(_nw69).ArraySet1CodePoint(_183_v8, 12)
			(_nw69).ArraySet1CodePoint(_183_v8, 13)
			(_nw69).ArraySet1CodePoint((_180_v6).Select((Companion_Default___.SafeIndex(_346_i11, _dafny.IntOfUint32((_180_v6).Cardinality()))).Uint32()).(_dafny.CodePoint), 14)
			(_nw69).ArraySet1CodePoint(_183_v8, 15)
			(_nw69).ArraySet1CodePoint(_183_v8, 16)
			(_nw69).ArraySet1CodePoint(_183_v8, 17)
			(_nw69).ArraySet1CodePoint(_183_v8, 18)
			(_nw69).ArraySet1CodePoint(_183_v8, 19)
			(_nw69).ArraySet1CodePoint(Companion_Default___.Fm12(_327_v120, _180_v6, _180_v6, _182_globalState), 20)
			(_nw69).ArraySet1CodePoint(_183_v8, 21)
			(_nw69).ArraySet1CodePoint(_183_v8, 22)
			(_nw69).ArraySet1CodePoint(_183_v8, 23)
			(_nw69).ArraySet1CodePoint(_183_v8, 24)
			(_nw69).ArraySet1CodePoint(Companion_Default___.Fm12(_dafny.IntOfUint32((_175_v1).Cardinality()), _180_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(454))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}((func(_407_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_408_i17 _dafny.Int) _dafny.CodePoint {
					return _407_v8
				}
			})(_183_v8))), _182_globalState), 25)
			_406_v158 = _nw69
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_406_v158), 0))
			_ = _index38
			(_406_v158).ArraySet1CodePoint(_183_v8, (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_406_v158), 0))
			_ = _index39
			var _rhs33 D13 = _403_v155
			_ = _rhs33
			var _rhs34 _dafny.Set = _404_v156
			_ = _rhs34
			var _rhs35 _dafny.Int = _327_v120
			_ = _rhs35
			var _rhs36 _dafny.Int = _176_v2
			_ = _rhs36
			var _rhs37 _dafny.CodePoint = _183_v8
			_ = _rhs37
			var _lhs19 _dafny.Array = _406_v158
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_406_v158), 0))
			_ = _lhs20
			_403_v155 = _rhs33
			_404_v156 = _rhs34
			_176_v2 = _rhs35
			_327_v120 = _rhs36
			(_lhs19).ArraySet1CodePoint(_rhs37, (_lhs20).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_177_v3), 0))
			_ = _index40
			(_177_v3).ArraySet1(_176_v2, (_index40).Int())
			var _409_v159 _dafny.MultiSet
			_ = _409_v159
			_409_v159 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt(_346_i11, (_402_v154).Cardinality()))
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_177_v3), 0))
			_ = _index41
			(_177_v3).ArraySet1((_409_v159).Cardinality(), (_index41).Int())
			var _410_v160 _dafny.Map
			_ = _410_v160
			_410_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _175_v1)
			var _411_v161 _dafny.Sequence
			_ = _411_v161
			_411_v161 = _dafny.SeqOf((func() _dafny.Sequence {
				if (_410_v160).Contains(true) {
					return (_410_v160).Get(true).(_dafny.Sequence)
				}
				return _175_v1
			})())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_177_v3), 0))
			_ = _index42
			(_177_v3).ArraySet1((_176_v2).Plus(_dafny.IntOfUint32(((_411_v161).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm5(_327_v120, _182_globalState), _dafny.IntOfUint32((_411_v161).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())), (_index42).Int())
		} else if _source5.Is_DC31() {
			var _412___mcc_h18 _dafny.Set = _source5.Get_().(D12_DC31).Cf55
			_ = _412___mcc_h18
			var _413_cf55 _dafny.Set = _412___mcc_h18
			_ = _413_cf55
			var _414_v162 bool
			_ = _414_v162
			var _415_v163 bool
			_ = _415_v163
			var _out45 bool
			_ = _out45
			var _out46 bool
			_ = _out46
			_out45, _out46 = (_372_v140).M1((_348_v124).F21(), _dafny.IntOfUint32((_180_v6).Cardinality()), !(_325_v118), Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_279_v79).Contains(_325_v118) {
					return (_279_v79).Get(_325_v118).(_dafny.Int)
				}
				return _dafny.IntOfInt64(-122)
			})(), _dafny.IntOfInt64(-560)), _182_globalState)
			_414_v162 = _out45
			_415_v163 = _out46
			var _rhs38 bool = (_346_i11).Cmp((_346_i11).Minus(_dafny.IntOfUint32((Companion_Default___.Fm19(_253_v54, _414_v162, _182_globalState)).Cardinality()))) <= 0
			_ = _rhs38
			var _rhs39 _dafny.Int = (func() _dafny.Int {
				if (_348_v124).F21() {
					return _dafny.IntOfInt64(-772)
				}
				return (_dafny.Zero).Minus(_176_v2)
			})()
			_ = _rhs39
			var _rhs40 bool = _326_v119
			_ = _rhs40
			var _rhs41 bool = (_348_v124).F21()
			_ = _rhs41
			var _lhs21 *GlobalState = _182_globalState
			_ = _lhs21
			var _lhs22 *GlobalState = _182_globalState
			_ = _lhs22
			_415_v163 = _rhs38
			_176_v2 = _rhs39
			_lhs21.F15 = _rhs40
			_lhs22.F15 = _rhs41
			var _416_v164 D3
			_ = _416_v164
			_416_v164 = Companion_D3_.Create_DC11_(_279_v79, (_251_v52).Cardinality(), _177_v3)
			var _rhs42 _dafny.Map = ((_416_v164).Dtor_cf18()).Merge(_279_v79)
			_ = _rhs42
			var _rhs43 bool = false
			_ = _rhs43
			_279_v79 = _rhs42
			_415_v163 = _rhs43
			_180_v6 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if _415_v163 {
					return _180_v6
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("esrnngllh")
			})(), _180_v6)
		} else {
			var _417___mcc_h19 D12 = _source5.Get_().(D12_DC35).Cf62
			_ = _417___mcc_h19
			var _418_cf62 D12 = _417___mcc_h19
			_ = _418_cf62
			_176_v2 = _327_v120
			_249_v50 = _249_v50
			(_182_globalState).F1 = _dafny.SeqOf((_348_v124).F21())
			var _419_v165 _dafny.Set
			_ = _419_v165
			_419_v165 = _dafny.SetOf(_177_v3)
			var _420_v166 D15
			_ = _420_v166
			_420_v166 = Companion_D15_.Create_DC42_(_372_v140)
			var _421_v167 _dafny.Sequence
			_ = _421_v167
			_421_v167 = _dafny.SeqOf(_372_v140, _372_v140)
			var _422_v168 _dafny.Array
			_ = _422_v168
			var _nwElement0_9 T1 = _372_v140
			_ = _nwElement0_9
			var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(27))
			_ = _nw70
			(_nw70).ArraySet1(_nwElement0_9, 0)
			(_nw70).ArraySet1((Companion_D15_.Create_DC42_(_372_v140)).Dtor_cf75(), 1)
			(_nw70).ArraySet1(_372_v140, 2)
			(_nw70).ArraySet1((_420_v166).Dtor_cf75(), 3)
			(_nw70).ArraySet1(_372_v140, 4)
			(_nw70).ArraySet1(_372_v140, 5)
			(_nw70).ArraySet1(_372_v140, 6)
			(_nw70).ArraySet1(_372_v140, 7)
			(_nw70).ArraySet1(_372_v140, 8)
			(_nw70).ArraySet1(_372_v140, 9)
			(_nw70).ArraySet1(_372_v140, 10)
			(_nw70).ArraySet1(_372_v140, 11)
			(_nw70).ArraySet1(_372_v140, 12)
			(_nw70).ArraySet1(_372_v140, 13)
			(_nw70).ArraySet1(_372_v140, 14)
			(_nw70).ArraySet1(_372_v140, 15)
			(_nw70).ArraySet1(_372_v140, 16)
			(_nw70).ArraySet1(_372_v140, 17)
			(_nw70).ArraySet1(_372_v140, 18)
			(_nw70).ArraySet1(_372_v140, 19)
			(_nw70).ArraySet1(_372_v140, 20)
			(_nw70).ArraySet1(_372_v140, 21)
			(_nw70).ArraySet1(_372_v140, 22)
			(_nw70).ArraySet1(_372_v140, 23)
			(_nw70).ArraySet1((_421_v167).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm19(_253_v54, (_277_v77).F23(), _182_globalState)).Cardinality()), _dafny.IntOfUint32((_421_v167).Cardinality()))).Uint32()).(T1), 24)
			(_nw70).ArraySet1(_372_v140, 25)
			(_nw70).ArraySet1(_372_v140, 26)
			_422_v168 = _nw70
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_422_v168), 0))
			_ = _index43
			(_422_v168).ArraySet1(_372_v140, (_index43).Int())
			var _423_v169 _dafny.Sequence
			_ = _423_v169
			_423_v169 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("tetmmpkc"), _180_v6, _180_v6)
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_422_v168), 0))
			_ = _index44
			var _rhs44 _dafny.Set = _419_v165
			_ = _rhs44
			var _rhs45 T1 = _372_v140
			_ = _rhs45
			var _rhs46 _dafny.Sequence = (_423_v169).Select((Companion_Default___.SafeIndex(_327_v120, _dafny.IntOfUint32((_423_v169).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs46
			var _lhs23 _dafny.Array = _422_v168
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_422_v168), 0))
			_ = _lhs24
			_419_v165 = _rhs44
			(_lhs23).ArraySet1(_rhs45, (_lhs24).Int())
			_180_v6 = _rhs46
		}
		_176_v2 = _327_v120
	}
	var _424_v170 _dafny.Array
	_ = _424_v170
	var _len0_11 _dafny.Int = _dafny.IntOfInt64(28)
	_ = _len0_11
	var _nw71 _dafny.Array
	_ = _nw71
	if _len0_11.Cmp(_dafny.Zero) == 0 {
		_nw71 = _dafny.NewArray(_len0_11)
	} else {
		var _init11 func(_dafny.Int) bool = (func(_425_v1 _dafny.Sequence) func(_dafny.Int) bool {
			return func(_426_i19 _dafny.Int) bool {
				return _dafny.Companion_Sequence_.IsPrefixOf(_425_v1, _425_v1)
			}
		})(_175_v1)
		_ = _init11
		var _element0_11 = _init11(_dafny.Zero)
		_ = _element0_11
		_nw71 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
		(_nw71).ArraySet1(_element0_11, 0)
		var _nativeLen0_11 = (_len0_11).Int()
		_ = _nativeLen0_11
		for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
			(_nw71).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
		}
	}
	_424_v170 = _nw71
	for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_424_v170), 0))); ; {
		_guard_loop_0, _ok21 := _iter21()
		if !_ok21 {
			break
		}
		var _427_i18 _dafny.Int
		_427_i18 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_427_i18).Sign() != -1) && ((_427_i18).Cmp(_dafny.ArrayLen((_424_v170), 0)) < 0)) {
			(_424_v170).ArraySet1(false, (_427_i18).Int())
		}
	}
	var _428_v171 _dafny.Map
	_ = _428_v171
	_428_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v2, _174_v0)
	var _hi3 _dafny.Int = ((func() _dafny.Int {
		if (_279_v79).Contains((_277_v77).F23()) {
			return (_279_v79).Get((_277_v77).F23()).(_dafny.Int)
		}
		return _176_v2
	})()).Minus(_327_v120)
	_ = _hi3
	for _429_i20 := ((_277_v77).Fm14((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_428_v171, _327_v120)).Cardinality(), _182_globalState)).Times((_dafny.Zero).Minus(_327_v120)); _429_i20.Cmp(_hi3) < 0; _429_i20 = _429_i20.Plus(_dafny.One) {
		_327_v120 = (_dafny.Zero).Minus(_327_v120)
		var _430_v172 D10
		_ = _430_v172
		_430_v172 = Companion_D10_.Create_DC27_(_251_v52)
		var _source6 D10 = _430_v172
		_ = _source6
		if _source6.Is_DC28() {
			var _431___mcc_h20 bool = _source6.Get_().(D10_DC28).Cf49
			_ = _431___mcc_h20
			var _432___mcc_h21 _dafny.Int = _source6.Get_().(D10_DC28).Cf50
			_ = _432___mcc_h21
			var _433_cf50 _dafny.Int = _432___mcc_h21
			_ = _433_cf50
			var _434_cf49 bool = _431___mcc_h20
			_ = _434_cf49
			var _435_v173 _dafny.Sequence
			_ = _435_v173
			_435_v173 = _dafny.SeqOf(_253_v54)
			var _436_v174 D13
			_ = _436_v174
			_436_v174 = Companion_D13_.Create_DC39_(Companion_D13_.Create_DC37_(_176_v2, _435_v173))
			var _437_v175 _dafny.Map
			_ = _437_v175
			_437_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _436_v174)
			_437_v175 = _437_v175
			var _438_v176 _dafny.Set
			_ = _438_v176
			_438_v176 = _dafny.SetOf(_177_v3, _177_v3)
			var _439_v177 _dafny.Map
			_ = _439_v177
			_439_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_327_v120, _438_v176)
			var _440_v178 T0
			_ = _440_v178
			var _nw72 *C2 = New_C2_()
			_ = _nw72
			_nw72.Ctor__(_179_v5, (_277_v77).F23())
			_440_v178 = _nw72
			var _441_v179 _dafny.Map
			_ = _441_v179
			_441_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_424_v170, _440_v178)
			_174_v0 = (_277_v77).Fm15((_438_v176).IsSubsetOf((func() _dafny.Set {
				if (_439_v177).Contains(_176_v2) {
					return (_439_v177).Get(_176_v2).(_dafny.Set)
				}
				return _438_v176
			})()), _326_v119, Companion_Default___.SafeDivisionInt((_441_v179).Cardinality(), _327_v120), _279_v79, _182_globalState)
			_326_v119 = _434_cf49
			var _442_v180 *C5
			_ = _442_v180
			var _nw73 *C5 = New_C5_()
			_ = _nw73
			_nw73.Ctor__(_429_i20)
			_442_v180 = _nw73
		} else {
			var _443___mcc_h22 _dafny.MultiSet = _source6.Get_().(D10_DC27).Cf48
			_ = _443___mcc_h22
			var _444_cf48 _dafny.MultiSet = _443___mcc_h22
			_ = _444_cf48
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_424_v170), 0))
			_ = _index45
			(_424_v170).ArraySet1(true, (_index45).Int())
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_424_v170), 0))
			_ = _index46
			(_424_v170).ArraySet1(_326_v119, (_index46).Int())
			_181_v7 = (_181_v7).Update(_327_v120, _dafny.IntOfUint32((_180_v6).Cardinality()))
			_327_v120 = _176_v2
			var _445_v181 _dafny.Array
			_ = _445_v181
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_12
			var _nw74 _dafny.Array
			_ = _nw74
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw74 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) bool = (func(_446_v0 bool) func(_dafny.Int) bool {
					return func(_447_i21 _dafny.Int) bool {
						return _446_v0
					}
				})(_174_v0)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw74 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw74).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw74).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_445_v181 = _nw74
			var _448_v182 D1
			_ = _448_v182
			_448_v182 = Companion_D1_.Create_DC4_((_dafny.Zero).Minus(_327_v120), (_277_v77).F23(), _445_v181)
			var _449_v183 _dafny.Sequence
			_ = _449_v183
			_449_v183 = _dafny.SeqOf(_175_v1, _175_v1, _dafny.SeqOf(false, _326_v119))
			var _rhs47 bool = (_277_v77).F23()
			_ = _rhs47
			var _rhs48 bool = (_327_v120).Cmp((_448_v182).Dtor_cf4()) == 0
			_ = _rhs48
			var _rhs49 bool = _326_v119
			_ = _rhs49
			var _rhs50 _dafny.Int = (Companion_Default___.SafeDivisionInt(_176_v2, _327_v120)).Minus(Companion_Default___.SafeModuloInt(_327_v120, _dafny.IntOfUint32(((_449_v183).Select((Companion_Default___.SafeIndex(_429_i20, _dafny.IntOfUint32((_449_v183).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())))
			_ = _rhs50
			var _lhs25 *GlobalState = _182_globalState
			_ = _lhs25
			var _lhs26 *GlobalState = _182_globalState
			_ = _lhs26
			_lhs25.F11 = _rhs47
			_326_v119 = _rhs48
			_lhs26.F11 = _rhs49
			_176_v2 = _rhs50
		}
		var _450_v184 *C1
		_ = _450_v184
		var _nw75 *C1 = New_C1_()
		_ = _nw75
		_nw75.Ctor__(_177_v3)
		_450_v184 = _nw75
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen(((_450_v184).F22()), 0))
		_ = _index47
		((_450_v184).F22()).ArraySet1(_327_v120, (_index47).Int())
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen(((_450_v184).F22()), 0))
		_ = _index48
		((_450_v184).F22()).ArraySet1(_327_v120, (_index48).Int())
	}
	var _hi4 _dafny.Int = _176_v2
	_ = _hi4
	for _451_i22 := _327_v120; _451_i22.Cmp(_hi4) < 0; _451_i22 = _451_i22.Plus(_dafny.One) {
		var _452_v185 bool
		_ = _452_v185
		var _453_v186 bool
		_ = _453_v186
		var _454_v187 _dafny.Int
		_ = _454_v187
		var _out47 bool
		_ = _out47
		var _out48 bool
		_ = _out48
		var _out49 _dafny.Int
		_ = _out49
		_out47, _out48, _out49 = Companion_Default___.M9((_277_v77).F23(), false, _182_globalState)
		_452_v185 = _out47
		_453_v186 = _out48
		_454_v187 = _out49
		var _455_v188 *C4
		_ = _455_v188
		var _nw76 *C4 = New_C4_()
		_ = _nw76
		_nw76.Ctor__(((_277_v77).Fm15(false, (_277_v77).F23(), _454_v187, _279_v79, _182_globalState)) || (_452_v185))
		_455_v188 = _nw76
		var _456_v189 *C2
		_ = _456_v189
		var _nw77 *C2 = New_C2_()
		_ = _nw77
		_nw77.Ctor__(_179_v5, (_277_v77).F23())
		_456_v189 = _nw77
		_181_v7 = (_181_v7).Update(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bmdrnblv")).Cardinality()), _454_v187), _176_v2)
	}
	var _457_v190 bool
	_ = _457_v190
	var _458_v191 bool
	_ = _458_v191
	var _459_v192 _dafny.Int
	_ = _459_v192
	var _out50 bool
	_ = _out50
	var _out51 bool
	_ = _out51
	var _out52 _dafny.Int
	_ = _out52
	_out50, _out51, _out52 = Companion_Default___.M9(_174_v0, true, _182_globalState)
	_457_v190 = _out50
	_458_v191 = _out51
	_459_v192 = _out52
	_328_v121 = (_328_v121).Update(_174_v0, !(_326_v119))
	var _460_v193 *C4
	_ = _460_v193
	var _nw78 *C4 = New_C4_()
	_ = _nw78
	_nw78.Ctor__(_326_v119)
	_460_v193 = _nw78
	var _461_v194 _dafny.MultiSet
	_ = _461_v194
	_461_v194 = _dafny.MultiSetOf(_424_v170)
	var _462_v195 _dafny.MultiSet
	_ = _462_v195
	_462_v195 = _dafny.MultiSetOf(_176_v2)
	var _rhs51 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_180_v6, (Companion_Default___.SafeIndex((_459_v192).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_253_v54, (Companion_Default___.SafeIndex(_459_v192, _dafny.IntOfUint32((_253_v54).Cardinality()))).Uint32(), (func() _dafny.Int {
		if (_461_v194).Contains(_424_v170) {
			return (_461_v194).Multiplicity(_424_v170)
		}
		return _176_v2
	})())).Cardinality())), _dafny.IntOfUint32((_180_v6).Cardinality()))).Uint32(), _dafny.CodePoint('w'))
	_ = _rhs51
	var _rhs52 _dafny.Int = ((_462_v195).Union(_462_v195)).Cardinality()
	_ = _rhs52
	var _rhs53 *C4 = _460_v193
	_ = _rhs53
	_180_v6 = _rhs51
	_459_v192 = _rhs52
	_460_v193 = _rhs53
	(_182_globalState).F1 = _dafny.SeqOf(true)
	var _463_v196 *C1
	_ = _463_v196
	var _nw79 *C1 = New_C1_()
	_ = _nw79
	_nw79.Ctor__(_177_v3)
	_463_v196 = _nw79
	_dafny.Print(_174_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_175_v1, _dafny.SeqOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_176_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v5).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_180_v6, _dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_181_v7).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(644), _dafny.IntOfInt64(644))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_182_globalState.F1, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_182_globalState).F4(), _dafny.SeqOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F8().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_182_globalState.F11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F12()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(644), _dafny.IntOfInt64(644))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_182_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_182_globalState.F16, _dafny.SeqOf(true, true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_182_globalState.F17)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_183_v8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v51).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_251_v52).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_252_v53, _dafny.SeqOf(Companion_D10_.Create_DC27_(_dafny.MultiSetOf(false, true)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_253_v54, _dafny.SeqOf(_dafny.IntOfInt64(-916))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v77).F23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_278_v78).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_279_v79).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-916))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_325_v118)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_326_v119)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_327_v120)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_328_v121).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_329_v122).Dtor_cf63(), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v170).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_428_v171).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_457_v190)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_458_v191)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_459_v192)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_460_v193).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_461_v194).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_462_v195).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
}

// End of class Default__

// Definition of datatype D0
type D0 struct {
	Data_D0_
}

func (_this D0) Get_() Data_D0_ {
	return _this.Data_D0_
}

type Data_D0_ interface {
	isD0()
}

type CompanionStruct_D0_ struct {
}

var Companion_D0_ = CompanionStruct_D0_{}

type D0_DC1 struct {
	Cf1 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 bool
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 bool) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC2 struct {
	Cf2 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 D0) D0 {
	return D0{D0_DC2{Cf2}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf2() D0 {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D0) Equals(other D0) bool {
	switch data1 := _this.Get_().(type) {
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf1 == data2.Cf1
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2.Equals(data2.Cf2)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D0) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D0)
	return ok && _this.Equals(typed)
}

func Type_D0_() _dafny.TypeDescriptor {
	return type_D0_{}
}

type type_D0_ struct {
}

func (_this type_D0_) Default() interface{} {
	return Companion_D0_.Default()
}

func (_this type_D0_) String() string {
	return "main.D0"
}
func (_this D0) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D0{}

// End of datatype D0

// Definition of datatype D1
type D1 struct {
	Data_D1_
}

func (_this D1) Get_() Data_D1_ {
	return _this.Data_D1_
}

type Data_D1_ interface {
	isD1()
}

type CompanionStruct_D1_ struct {
}

var Companion_D1_ = CompanionStruct_D1_{}

type D1_DC4 struct {
	Cf4 _dafny.Int
	Cf5 bool
	Cf6 _dafny.Array
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf4 _dafny.Int, Cf5 bool, Cf6 _dafny.Array) D1 {
	return D1{D1_DC4{Cf4, Cf5, Cf6}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf7 _dafny.Sequence
	Cf8 _dafny.Int
	Cf9 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf7 _dafny.Sequence, Cf8 _dafny.Int, Cf9 _dafny.Int) D1 {
	return D1{D1_DC5{Cf7, Cf8, Cf9}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
	Cf10 _dafny.Int
	Cf11 _dafny.Int
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf10 _dafny.Int, Cf11 _dafny.Int) D1 {
	return D1{D1_DC6{Cf10, Cf11}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC3 struct {
	Cf3 _dafny.Array
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 _dafny.Array) D1 {
	return D1{D1_DC3{Cf3}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.Zero, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Array {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Sequence {
	return _this.Get_().(D1_DC5).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf11
}

func (_this D1) Dtor_cf3() _dafny.Array {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + data.Cf7.VerbatimString(true) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5 == data2.Cf5 && data1.Cf6 == data2.Cf6
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf7.Equals(data2.Cf7) && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf3 == data2.Cf3
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D1) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D1)
	return ok && _this.Equals(typed)
}

func Type_D1_() _dafny.TypeDescriptor {
	return type_D1_{}
}

type type_D1_ struct {
}

func (_this type_D1_) Default() interface{} {
	return Companion_D1_.Default()
}

func (_this type_D1_) String() string {
	return "main.D1"
}
func (_this D1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D1{}

// End of datatype D1

// Definition of datatype D2
type D2 struct {
	Data_D2_
}

func (_this D2) Get_() Data_D2_ {
	return _this.Data_D2_
}

type Data_D2_ interface {
	isD2()
}

type CompanionStruct_D2_ struct {
}

var Companion_D2_ = CompanionStruct_D2_{}

type D2_DC8 struct {
	Cf13 _dafny.Array
	Cf14 _dafny.Int
	Cf15 _dafny.Int
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf13 _dafny.Array, Cf14 _dafny.Int, Cf15 _dafny.Int) D2 {
	return D2{D2_DC8{Cf13, Cf14, Cf15}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC7 struct {
	Cf12 _dafny.MultiSet
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf12 _dafny.MultiSet) D2 {
	return D2{D2_DC7{Cf12}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf13() _dafny.Array {
	return _this.Get_().(D2_DC8).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf14
}

func (_this D2) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf15
}

func (_this D2) Dtor_cf12() _dafny.MultiSet {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15.Cmp(data2.Cf15) == 0
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf12.Equals(data2.Cf12)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D2) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D2)
	return ok && _this.Equals(typed)
}

func Type_D2_() _dafny.TypeDescriptor {
	return type_D2_{}
}

type type_D2_ struct {
}

func (_this type_D2_) Default() interface{} {
	return Companion_D2_.Default()
}

func (_this type_D2_) String() string {
	return "main.D2"
}
func (_this D2) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D2{}

// End of datatype D2

// Definition of datatype D3
type D3 struct {
	Data_D3_
}

func (_this D3) Get_() Data_D3_ {
	return _this.Data_D3_
}

type Data_D3_ interface {
	isD3()
}

type CompanionStruct_D3_ struct {
}

var Companion_D3_ = CompanionStruct_D3_{}

type D3_DC10 struct {
	Cf17 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf17 _dafny.Int) D3 {
	return D3{D3_DC10{Cf17}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC11 struct {
	Cf18 _dafny.Map
	Cf19 _dafny.Int
	Cf20 _dafny.Array
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf18 _dafny.Map, Cf19 _dafny.Int, Cf20 _dafny.Array) D3 {
	return D3{D3_DC11{Cf18, Cf19, Cf20}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC9 struct {
	Cf16 _dafny.Set
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf16 _dafny.Set) D3 {
	return D3{D3_DC9{Cf16}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(_dafny.Zero)
}

func (_this D3) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Map {
	return _this.Get_().(D3_DC11).Cf18
}

func (_this D3) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf19
}

func (_this D3) Dtor_cf20() _dafny.Array {
	return _this.Get_().(D3_DC11).Cf20
}

func (_this D3) Dtor_cf16() _dafny.Set {
	return _this.Get_().(D3_DC9).Cf16
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf16) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf18.Equals(data2.Cf18) && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20 == data2.Cf20
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf16.Equals(data2.Cf16)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D3) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D3)
	return ok && _this.Equals(typed)
}

func Type_D3_() _dafny.TypeDescriptor {
	return type_D3_{}
}

type type_D3_ struct {
}

func (_this type_D3_) Default() interface{} {
	return Companion_D3_.Default()
}

func (_this type_D3_) String() string {
	return "main.D3"
}
func (_this D3) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D3{}

// End of datatype D3

// Definition of datatype D4
type D4 struct {
	Data_D4_
}

func (_this D4) Get_() Data_D4_ {
	return _this.Data_D4_
}

type Data_D4_ interface {
	isD4()
}

type CompanionStruct_D4_ struct {
}

var Companion_D4_ = CompanionStruct_D4_{}

type D4_DC13 struct {
	Cf22 _dafny.Int
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf22 _dafny.Int) D4 {
	return D4{D4_DC13{Cf22}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC14 struct {
	Cf23 bool
	Cf24 _dafny.Int
	Cf25 _dafny.Set
	Cf26 _dafny.CodePoint
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf23 bool, Cf24 _dafny.Int, Cf25 _dafny.Set, Cf26 _dafny.CodePoint) D4 {
	return D4{D4_DC14{Cf23, Cf24, Cf25, Cf26}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC12 struct {
	Cf21 T0
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf21 T0) D4 {
	return D4{D4_DC12{Cf21}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC15 struct {
	Cf27 D4
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf27 D4) D4 {
	return D4{D4_DC15{Cf27}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_(_dafny.Zero)
}

func (_this D4) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf22
}

func (_this D4) Dtor_cf23() bool {
	return _this.Get_().(D4_DC14).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf24
}

func (_this D4) Dtor_cf25() _dafny.Set {
	return _this.Get_().(D4_DC14).Cf25
}

func (_this D4) Dtor_cf26() _dafny.CodePoint {
	return _this.Get_().(D4_DC14).Cf26
}

func (_this D4) Dtor_cf21() T0 {
	return _this.Get_().(D4_DC12).Cf21
}

func (_this D4) Dtor_cf27() D4 {
	return _this.Get_().(D4_DC15).Cf27
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf22.Cmp(data2.Cf22) == 0
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf23 == data2.Cf23 && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25.Equals(data2.Cf25) && data1.Cf26 == data2.Cf26
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && _dafny.AreEqual(data1.Cf21, data2.Cf21)
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf27.Equals(data2.Cf27)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D4) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D4)
	return ok && _this.Equals(typed)
}

func Type_D4_() _dafny.TypeDescriptor {
	return type_D4_{}
}

type type_D4_ struct {
}

func (_this type_D4_) Default() interface{} {
	return Companion_D4_.Default()
}

func (_this type_D4_) String() string {
	return "main.D4"
}
func (_this D4) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D4{}

// End of datatype D4

// Definition of datatype D5
type D5 struct {
	Data_D5_
}

func (_this D5) Get_() Data_D5_ {
	return _this.Data_D5_
}

type Data_D5_ interface {
	isD5()
}

type CompanionStruct_D5_ struct {
}

var Companion_D5_ = CompanionStruct_D5_{}

type D5_DC17 struct {
	Cf29 bool
	Cf30 D2
	Cf31 bool
	Cf32 bool
	Cf33 _dafny.Int
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf29 bool, Cf30 D2, Cf31 bool, Cf32 bool, Cf33 _dafny.Int) D5 {
	return D5{D5_DC17{Cf29, Cf30, Cf31, Cf32, Cf33}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC16 struct {
	Cf28 *C1
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf28 *C1) D5 {
	return D5{D5_DC16{Cf28}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC18 struct {
	Cf34 D5
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf34 D5) D5 {
	return D5{D5_DC18{Cf34}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC17_(false, Companion_D2_.Default(), false, false, _dafny.Zero)
}

func (_this D5) Dtor_cf29() bool {
	return _this.Get_().(D5_DC17).Cf29
}

func (_this D5) Dtor_cf30() D2 {
	return _this.Get_().(D5_DC17).Cf30
}

func (_this D5) Dtor_cf31() bool {
	return _this.Get_().(D5_DC17).Cf31
}

func (_this D5) Dtor_cf32() bool {
	return _this.Get_().(D5_DC17).Cf32
}

func (_this D5) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D5_DC17).Cf33
}

func (_this D5) Dtor_cf28() *C1 {
	return _this.Get_().(D5_DC16).Cf28
}

func (_this D5) Dtor_cf34() D5 {
	return _this.Get_().(D5_DC18).Cf34
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf28) + ")"
		}
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf34) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf29 == data2.Cf29 && data1.Cf30.Equals(data2.Cf30) && data1.Cf31 == data2.Cf31 && data1.Cf32 == data2.Cf32 && data1.Cf33.Cmp(data2.Cf33) == 0
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf28 == data2.Cf28
		}
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
			return ok && data1.Cf34.Equals(data2.Cf34)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D5) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D5)
	return ok && _this.Equals(typed)
}

func Type_D5_() _dafny.TypeDescriptor {
	return type_D5_{}
}

type type_D5_ struct {
}

func (_this type_D5_) Default() interface{} {
	return Companion_D5_.Default()
}

func (_this type_D5_) String() string {
	return "main.D5"
}
func (_this D5) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D5{}

// End of datatype D5

// Definition of datatype D6
type D6 struct {
	Data_D6_
}

func (_this D6) Get_() Data_D6_ {
	return _this.Data_D6_
}

type Data_D6_ interface {
	isD6()
}

type CompanionStruct_D6_ struct {
}

var Companion_D6_ = CompanionStruct_D6_{}

type D6_DC20 struct {
	Cf36 bool
	Cf37 _dafny.Int
	Cf38 bool
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_(Cf36 bool, Cf37 _dafny.Int, Cf38 bool) D6 {
	return D6{D6_DC20{Cf36, Cf37, Cf38}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

type D6_DC19 struct {
	Cf35 _dafny.MultiSet
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf35 _dafny.MultiSet) D6 {
	return D6{D6_DC19{Cf35}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC20_(false, _dafny.Zero, false)
}

func (_this D6) Dtor_cf36() bool {
	return _this.Get_().(D6_DC20).Cf36
}

func (_this D6) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D6_DC20).Cf37
}

func (_this D6) Dtor_cf38() bool {
	return _this.Get_().(D6_DC20).Cf38
}

func (_this D6) Dtor_cf35() _dafny.MultiSet {
	return _this.Get_().(D6_DC19).Cf35
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC20:
		{
			return "D6.DC20" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf35) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC20:
		{
			data2, ok := other.Get_().(D6_DC20)
			return ok && data1.Cf36 == data2.Cf36 && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38 == data2.Cf38
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf35.Equals(data2.Cf35)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D6) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D6)
	return ok && _this.Equals(typed)
}

func Type_D6_() _dafny.TypeDescriptor {
	return type_D6_{}
}

type type_D6_ struct {
}

func (_this type_D6_) Default() interface{} {
	return Companion_D6_.Default()
}

func (_this type_D6_) String() string {
	return "main.D6"
}
func (_this D6) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D6{}

// End of datatype D6

// Definition of datatype D7
type D7 struct {
	Data_D7_
}

func (_this D7) Get_() Data_D7_ {
	return _this.Data_D7_
}

type Data_D7_ interface {
	isD7()
}

type CompanionStruct_D7_ struct {
}

var Companion_D7_ = CompanionStruct_D7_{}

type D7_DC21 struct {
	Cf39 _dafny.Array
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf39 _dafny.Array) D7 {
	return D7{D7_DC21{Cf39}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D7) Dtor_cf39() _dafny.Array {
	return _this.Get_().(D7_DC21).Cf39
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf39) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf39 == data2.Cf39
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D7) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D7)
	return ok && _this.Equals(typed)
}

func Type_D7_() _dafny.TypeDescriptor {
	return type_D7_{}
}

type type_D7_ struct {
}

func (_this type_D7_) Default() interface{} {
	return Companion_D7_.Default()
}

func (_this type_D7_) String() string {
	return "main.D7"
}
func (_this D7) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D7{}

// End of datatype D7

// Definition of datatype D8
type D8 struct {
	Data_D8_
}

func (_this D8) Get_() Data_D8_ {
	return _this.Data_D8_
}

type Data_D8_ interface {
	isD8()
}

type CompanionStruct_D8_ struct {
}

var Companion_D8_ = CompanionStruct_D8_{}

type D8_DC23 struct {
	Cf41 _dafny.MultiSet
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf41 _dafny.MultiSet) D8 {
	return D8{D8_DC23{Cf41}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC22 struct {
	Cf40 _dafny.Array
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf40 _dafny.Array) D8 {
	return D8{D8_DC22{Cf40}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC23_(_dafny.EmptyMultiSet)
}

func (_this D8) Dtor_cf41() _dafny.MultiSet {
	return _this.Get_().(D8_DC23).Cf41
}

func (_this D8) Dtor_cf40() _dafny.Array {
	return _this.Get_().(D8_DC22).Cf40
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf41) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf41.Equals(data2.Cf41)
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf40 == data2.Cf40
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D8) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D8)
	return ok && _this.Equals(typed)
}

func Type_D8_() _dafny.TypeDescriptor {
	return type_D8_{}
}

type type_D8_ struct {
}

func (_this type_D8_) Default() interface{} {
	return Companion_D8_.Default()
}

func (_this type_D8_) String() string {
	return "main.D8"
}
func (_this D8) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D8{}

// End of datatype D8

// Definition of datatype D9
type D9 struct {
	Data_D9_
}

func (_this D9) Get_() Data_D9_ {
	return _this.Data_D9_
}

type Data_D9_ interface {
	isD9()
}

type CompanionStruct_D9_ struct {
}

var Companion_D9_ = CompanionStruct_D9_{}

type D9_DC25 struct {
	Cf43 _dafny.Int
	Cf44 bool
	Cf45 _dafny.Int
	Cf46 *C0
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf43 _dafny.Int, Cf44 bool, Cf45 _dafny.Int, Cf46 *C0) D9 {
	return D9{D9_DC25{Cf43, Cf44, Cf45, Cf46}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC24 struct {
	Cf42 _dafny.Sequence
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf42 _dafny.Sequence) D9 {
	return D9{D9_DC24{Cf42}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC26 struct {
	Cf47 D9
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf47 D9) D9 {
	return D9{D9_DC26{Cf47}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC25_(_dafny.Zero, false, _dafny.Zero, (*C0)(nil))
}

func (_this D9) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf43
}

func (_this D9) Dtor_cf44() bool {
	return _this.Get_().(D9_DC25).Cf44
}

func (_this D9) Dtor_cf45() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf45
}

func (_this D9) Dtor_cf46() *C0 {
	return _this.Get_().(D9_DC25).Cf46
}

func (_this D9) Dtor_cf42() _dafny.Sequence {
	return _this.Get_().(D9_DC24).Cf42
}

func (_this D9) Dtor_cf47() D9 {
	return _this.Get_().(D9_DC26).Cf47
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf47) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf43.Cmp(data2.Cf43) == 0 && data1.Cf44 == data2.Cf44 && data1.Cf45.Cmp(data2.Cf45) == 0 && data1.Cf46 == data2.Cf46
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf47.Equals(data2.Cf47)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D9) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D9)
	return ok && _this.Equals(typed)
}

func Type_D9_() _dafny.TypeDescriptor {
	return type_D9_{}
}

type type_D9_ struct {
}

func (_this type_D9_) Default() interface{} {
	return Companion_D9_.Default()
}

func (_this type_D9_) String() string {
	return "main.D9"
}
func (_this D9) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D9{}

// End of datatype D9

// Definition of datatype D10
type D10 struct {
	Data_D10_
}

func (_this D10) Get_() Data_D10_ {
	return _this.Data_D10_
}

type Data_D10_ interface {
	isD10()
}

type CompanionStruct_D10_ struct {
}

var Companion_D10_ = CompanionStruct_D10_{}

type D10_DC28 struct {
	Cf49 bool
	Cf50 _dafny.Int
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf49 bool, Cf50 _dafny.Int) D10 {
	return D10{D10_DC28{Cf49, Cf50}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

type D10_DC27 struct {
	Cf48 _dafny.MultiSet
}

func (D10_DC27) isD10() {}

func (CompanionStruct_D10_) Create_DC27_(Cf48 _dafny.MultiSet) D10 {
	return D10{D10_DC27{Cf48}}
}

func (_this D10) Is_DC27() bool {
	_, ok := _this.Get_().(D10_DC27)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC28_(false, _dafny.Zero)
}

func (_this D10) Dtor_cf49() bool {
	return _this.Get_().(D10_DC28).Cf49
}

func (_this D10) Dtor_cf50() _dafny.Int {
	return _this.Get_().(D10_DC28).Cf50
}

func (_this D10) Dtor_cf48() _dafny.MultiSet {
	return _this.Get_().(D10_DC27).Cf48
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf48) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf49 == data2.Cf49 && data1.Cf50.Cmp(data2.Cf50) == 0
		}
	case D10_DC27:
		{
			data2, ok := other.Get_().(D10_DC27)
			return ok && data1.Cf48.Equals(data2.Cf48)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D10) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D10)
	return ok && _this.Equals(typed)
}

func Type_D10_() _dafny.TypeDescriptor {
	return type_D10_{}
}

type type_D10_ struct {
}

func (_this type_D10_) Default() interface{} {
	return Companion_D10_.Default()
}

func (_this type_D10_) String() string {
	return "main.D10"
}
func (_this D10) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D10{}

// End of datatype D10

// Definition of datatype D11
type D11 struct {
	Data_D11_
}

func (_this D11) Get_() Data_D11_ {
	return _this.Data_D11_
}

type Data_D11_ interface {
	isD11()
}

type CompanionStruct_D11_ struct {
}

var Companion_D11_ = CompanionStruct_D11_{}

type D11_DC30 struct {
	Cf52 bool
	Cf53 _dafny.Int
	Cf54 bool
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf52 bool, Cf53 _dafny.Int, Cf54 bool) D11 {
	return D11{D11_DC30{Cf52, Cf53, Cf54}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC29 struct {
	Cf51 _dafny.Set
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf51 _dafny.Set) D11 {
	return D11{D11_DC29{Cf51}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC30_(false, _dafny.Zero, false)
}

func (_this D11) Dtor_cf52() bool {
	return _this.Get_().(D11_DC30).Cf52
}

func (_this D11) Dtor_cf53() _dafny.Int {
	return _this.Get_().(D11_DC30).Cf53
}

func (_this D11) Dtor_cf54() bool {
	return _this.Get_().(D11_DC30).Cf54
}

func (_this D11) Dtor_cf51() _dafny.Set {
	return _this.Get_().(D11_DC29).Cf51
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf52 == data2.Cf52 && data1.Cf53.Cmp(data2.Cf53) == 0 && data1.Cf54 == data2.Cf54
		}
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
			return ok && data1.Cf51.Equals(data2.Cf51)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D11) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D11)
	return ok && _this.Equals(typed)
}

func Type_D11_() _dafny.TypeDescriptor {
	return type_D11_{}
}

type type_D11_ struct {
}

func (_this type_D11_) Default() interface{} {
	return Companion_D11_.Default()
}

func (_this type_D11_) String() string {
	return "main.D11"
}
func (_this D11) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D11{}

// End of datatype D11

// Definition of datatype D12
type D12 struct {
	Data_D12_
}

func (_this D12) Get_() Data_D12_ {
	return _this.Data_D12_
}

type Data_D12_ interface {
	isD12()
}

type CompanionStruct_D12_ struct {
}

var Companion_D12_ = CompanionStruct_D12_{}

type D12_DC32 struct {
	Cf56 _dafny.Int
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf56 _dafny.Int) D12 {
	return D12{D12_DC32{Cf56}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

type D12_DC33 struct {
	Cf57 _dafny.Int
	Cf58 bool
	Cf59 _dafny.MultiSet
	Cf60 bool
	Cf61 bool
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf57 _dafny.Int, Cf58 bool, Cf59 _dafny.MultiSet, Cf60 bool, Cf61 bool) D12 {
	return D12{D12_DC33{Cf57, Cf58, Cf59, Cf60, Cf61}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

type D12_DC34 struct {
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_() D12 {
	return D12{D12_DC34{}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

type D12_DC31 struct {
	Cf55 _dafny.Set
}

func (D12_DC31) isD12() {}

func (CompanionStruct_D12_) Create_DC31_(Cf55 _dafny.Set) D12 {
	return D12{D12_DC31{Cf55}}
}

func (_this D12) Is_DC31() bool {
	_, ok := _this.Get_().(D12_DC31)
	return ok
}

type D12_DC35 struct {
	Cf62 D12
}

func (D12_DC35) isD12() {}

func (CompanionStruct_D12_) Create_DC35_(Cf62 D12) D12 {
	return D12{D12_DC35{Cf62}}
}

func (_this D12) Is_DC35() bool {
	_, ok := _this.Get_().(D12_DC35)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC32_(_dafny.Zero)
}

func (_this D12) Dtor_cf56() _dafny.Int {
	return _this.Get_().(D12_DC32).Cf56
}

func (_this D12) Dtor_cf57() _dafny.Int {
	return _this.Get_().(D12_DC33).Cf57
}

func (_this D12) Dtor_cf58() bool {
	return _this.Get_().(D12_DC33).Cf58
}

func (_this D12) Dtor_cf59() _dafny.MultiSet {
	return _this.Get_().(D12_DC33).Cf59
}

func (_this D12) Dtor_cf60() bool {
	return _this.Get_().(D12_DC33).Cf60
}

func (_this D12) Dtor_cf61() bool {
	return _this.Get_().(D12_DC33).Cf61
}

func (_this D12) Dtor_cf55() _dafny.Set {
	return _this.Get_().(D12_DC31).Cf55
}

func (_this D12) Dtor_cf62() D12 {
	return _this.Get_().(D12_DC35).Cf62
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf56) + ")"
		}
	case D12_DC33:
		{
			return "D12.DC33" + "(" + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ")"
		}
	case D12_DC34:
		{
			return "D12.DC34"
		}
	case D12_DC31:
		{
			return "D12.DC31" + "(" + _dafny.String(data.Cf55) + ")"
		}
	case D12_DC35:
		{
			return "D12.DC35" + "(" + _dafny.String(data.Cf62) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf56.Cmp(data2.Cf56) == 0
		}
	case D12_DC33:
		{
			data2, ok := other.Get_().(D12_DC33)
			return ok && data1.Cf57.Cmp(data2.Cf57) == 0 && data1.Cf58 == data2.Cf58 && data1.Cf59.Equals(data2.Cf59) && data1.Cf60 == data2.Cf60 && data1.Cf61 == data2.Cf61
		}
	case D12_DC34:
		{
			_, ok := other.Get_().(D12_DC34)
			return ok
		}
	case D12_DC31:
		{
			data2, ok := other.Get_().(D12_DC31)
			return ok && data1.Cf55.Equals(data2.Cf55)
		}
	case D12_DC35:
		{
			data2, ok := other.Get_().(D12_DC35)
			return ok && data1.Cf62.Equals(data2.Cf62)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D12) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D12)
	return ok && _this.Equals(typed)
}

func Type_D12_() _dafny.TypeDescriptor {
	return type_D12_{}
}

type type_D12_ struct {
}

func (_this type_D12_) Default() interface{} {
	return Companion_D12_.Default()
}

func (_this type_D12_) String() string {
	return "main.D12"
}
func (_this D12) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D12{}

// End of datatype D12

// Definition of datatype D13
type D13 struct {
	Data_D13_
}

func (_this D13) Get_() Data_D13_ {
	return _this.Data_D13_
}

type Data_D13_ interface {
	isD13()
}

type CompanionStruct_D13_ struct {
}

var Companion_D13_ = CompanionStruct_D13_{}

type D13_DC37 struct {
	Cf64 _dafny.Int
	Cf65 _dafny.Sequence
}

func (D13_DC37) isD13() {}

func (CompanionStruct_D13_) Create_DC37_(Cf64 _dafny.Int, Cf65 _dafny.Sequence) D13 {
	return D13{D13_DC37{Cf64, Cf65}}
}

func (_this D13) Is_DC37() bool {
	_, ok := _this.Get_().(D13_DC37)
	return ok
}

type D13_DC38 struct {
	Cf66 _dafny.CodePoint
	Cf67 bool
	Cf68 _dafny.Int
	Cf69 _dafny.Array
}

func (D13_DC38) isD13() {}

func (CompanionStruct_D13_) Create_DC38_(Cf66 _dafny.CodePoint, Cf67 bool, Cf68 _dafny.Int, Cf69 _dafny.Array) D13 {
	return D13{D13_DC38{Cf66, Cf67, Cf68, Cf69}}
}

func (_this D13) Is_DC38() bool {
	_, ok := _this.Get_().(D13_DC38)
	return ok
}

type D13_DC36 struct {
	Cf63 _dafny.Sequence
}

func (D13_DC36) isD13() {}

func (CompanionStruct_D13_) Create_DC36_(Cf63 _dafny.Sequence) D13 {
	return D13{D13_DC36{Cf63}}
}

func (_this D13) Is_DC36() bool {
	_, ok := _this.Get_().(D13_DC36)
	return ok
}

type D13_DC39 struct {
	Cf70 D13
}

func (D13_DC39) isD13() {}

func (CompanionStruct_D13_) Create_DC39_(Cf70 D13) D13 {
	return D13{D13_DC39{Cf70}}
}

func (_this D13) Is_DC39() bool {
	_, ok := _this.Get_().(D13_DC39)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC37_(_dafny.Zero, _dafny.EmptySeq)
}

func (_this D13) Dtor_cf64() _dafny.Int {
	return _this.Get_().(D13_DC37).Cf64
}

func (_this D13) Dtor_cf65() _dafny.Sequence {
	return _this.Get_().(D13_DC37).Cf65
}

func (_this D13) Dtor_cf66() _dafny.CodePoint {
	return _this.Get_().(D13_DC38).Cf66
}

func (_this D13) Dtor_cf67() bool {
	return _this.Get_().(D13_DC38).Cf67
}

func (_this D13) Dtor_cf68() _dafny.Int {
	return _this.Get_().(D13_DC38).Cf68
}

func (_this D13) Dtor_cf69() _dafny.Array {
	return _this.Get_().(D13_DC38).Cf69
}

func (_this D13) Dtor_cf63() _dafny.Sequence {
	return _this.Get_().(D13_DC36).Cf63
}

func (_this D13) Dtor_cf70() D13 {
	return _this.Get_().(D13_DC39).Cf70
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC37:
		{
			return "D13.DC37" + "(" + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ")"
		}
	case D13_DC38:
		{
			return "D13.DC38" + "(" + _dafny.String(data.Cf66) + ", " + _dafny.String(data.Cf67) + ", " + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ")"
		}
	case D13_DC36:
		{
			return "D13.DC36" + "(" + _dafny.String(data.Cf63) + ")"
		}
	case D13_DC39:
		{
			return "D13.DC39" + "(" + _dafny.String(data.Cf70) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC37:
		{
			data2, ok := other.Get_().(D13_DC37)
			return ok && data1.Cf64.Cmp(data2.Cf64) == 0 && data1.Cf65.Equals(data2.Cf65)
		}
	case D13_DC38:
		{
			data2, ok := other.Get_().(D13_DC38)
			return ok && data1.Cf66 == data2.Cf66 && data1.Cf67 == data2.Cf67 && data1.Cf68.Cmp(data2.Cf68) == 0 && data1.Cf69 == data2.Cf69
		}
	case D13_DC36:
		{
			data2, ok := other.Get_().(D13_DC36)
			return ok && data1.Cf63.Equals(data2.Cf63)
		}
	case D13_DC39:
		{
			data2, ok := other.Get_().(D13_DC39)
			return ok && data1.Cf70.Equals(data2.Cf70)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D13) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D13)
	return ok && _this.Equals(typed)
}

func Type_D13_() _dafny.TypeDescriptor {
	return type_D13_{}
}

type type_D13_ struct {
}

func (_this type_D13_) Default() interface{} {
	return Companion_D13_.Default()
}

func (_this type_D13_) String() string {
	return "main.D13"
}
func (_this D13) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D13{}

// End of datatype D13

// Definition of datatype D14
type D14 struct {
	Data_D14_
}

func (_this D14) Get_() Data_D14_ {
	return _this.Data_D14_
}

type Data_D14_ interface {
	isD14()
}

type CompanionStruct_D14_ struct {
}

var Companion_D14_ = CompanionStruct_D14_{}

type D14_DC41 struct {
	Cf72 *C2
	Cf73 _dafny.Sequence
	Cf74 _dafny.Array
}

func (D14_DC41) isD14() {}

func (CompanionStruct_D14_) Create_DC41_(Cf72 *C2, Cf73 _dafny.Sequence, Cf74 _dafny.Array) D14 {
	return D14{D14_DC41{Cf72, Cf73, Cf74}}
}

func (_this D14) Is_DC41() bool {
	_, ok := _this.Get_().(D14_DC41)
	return ok
}

type D14_DC40 struct {
	Cf71 _dafny.Array
}

func (D14_DC40) isD14() {}

func (CompanionStruct_D14_) Create_DC40_(Cf71 _dafny.Array) D14 {
	return D14{D14_DC40{Cf71}}
}

func (_this D14) Is_DC40() bool {
	_, ok := _this.Get_().(D14_DC40)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC41_((*C2)(nil), _dafny.EmptySeq, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D14) Dtor_cf72() *C2 {
	return _this.Get_().(D14_DC41).Cf72
}

func (_this D14) Dtor_cf73() _dafny.Sequence {
	return _this.Get_().(D14_DC41).Cf73
}

func (_this D14) Dtor_cf74() _dafny.Array {
	return _this.Get_().(D14_DC41).Cf74
}

func (_this D14) Dtor_cf71() _dafny.Array {
	return _this.Get_().(D14_DC40).Cf71
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC41:
		{
			return "D14.DC41" + "(" + _dafny.String(data.Cf72) + ", " + data.Cf73.VerbatimString(true) + ", " + _dafny.String(data.Cf74) + ")"
		}
	case D14_DC40:
		{
			return "D14.DC40" + "(" + _dafny.String(data.Cf71) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC41:
		{
			data2, ok := other.Get_().(D14_DC41)
			return ok && data1.Cf72 == data2.Cf72 && data1.Cf73.Equals(data2.Cf73) && data1.Cf74 == data2.Cf74
		}
	case D14_DC40:
		{
			data2, ok := other.Get_().(D14_DC40)
			return ok && data1.Cf71 == data2.Cf71
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D14) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D14)
	return ok && _this.Equals(typed)
}

func Type_D14_() _dafny.TypeDescriptor {
	return type_D14_{}
}

type type_D14_ struct {
}

func (_this type_D14_) Default() interface{} {
	return Companion_D14_.Default()
}

func (_this type_D14_) String() string {
	return "main.D14"
}
func (_this D14) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D14{}

// End of datatype D14

// Definition of datatype D15
type D15 struct {
	Data_D15_
}

func (_this D15) Get_() Data_D15_ {
	return _this.Data_D15_
}

type Data_D15_ interface {
	isD15()
}

type CompanionStruct_D15_ struct {
}

var Companion_D15_ = CompanionStruct_D15_{}

type D15_DC43 struct {
	Cf76 _dafny.Set
	Cf77 bool
}

func (D15_DC43) isD15() {}

func (CompanionStruct_D15_) Create_DC43_(Cf76 _dafny.Set, Cf77 bool) D15 {
	return D15{D15_DC43{Cf76, Cf77}}
}

func (_this D15) Is_DC43() bool {
	_, ok := _this.Get_().(D15_DC43)
	return ok
}

type D15_DC44 struct {
	Cf78 bool
	Cf79 _dafny.Int
}

func (D15_DC44) isD15() {}

func (CompanionStruct_D15_) Create_DC44_(Cf78 bool, Cf79 _dafny.Int) D15 {
	return D15{D15_DC44{Cf78, Cf79}}
}

func (_this D15) Is_DC44() bool {
	_, ok := _this.Get_().(D15_DC44)
	return ok
}

type D15_DC42 struct {
	Cf75 T1
}

func (D15_DC42) isD15() {}

func (CompanionStruct_D15_) Create_DC42_(Cf75 T1) D15 {
	return D15{D15_DC42{Cf75}}
}

func (_this D15) Is_DC42() bool {
	_, ok := _this.Get_().(D15_DC42)
	return ok
}

type D15_DC45 struct {
	Cf80 D15
}

func (D15_DC45) isD15() {}

func (CompanionStruct_D15_) Create_DC45_(Cf80 D15) D15 {
	return D15{D15_DC45{Cf80}}
}

func (_this D15) Is_DC45() bool {
	_, ok := _this.Get_().(D15_DC45)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC43_(_dafny.EmptySet, false)
}

func (_this D15) Dtor_cf76() _dafny.Set {
	return _this.Get_().(D15_DC43).Cf76
}

func (_this D15) Dtor_cf77() bool {
	return _this.Get_().(D15_DC43).Cf77
}

func (_this D15) Dtor_cf78() bool {
	return _this.Get_().(D15_DC44).Cf78
}

func (_this D15) Dtor_cf79() _dafny.Int {
	return _this.Get_().(D15_DC44).Cf79
}

func (_this D15) Dtor_cf75() T1 {
	return _this.Get_().(D15_DC42).Cf75
}

func (_this D15) Dtor_cf80() D15 {
	return _this.Get_().(D15_DC45).Cf80
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC43:
		{
			return "D15.DC43" + "(" + _dafny.String(data.Cf76) + ", " + _dafny.String(data.Cf77) + ")"
		}
	case D15_DC44:
		{
			return "D15.DC44" + "(" + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ")"
		}
	case D15_DC42:
		{
			return "D15.DC42" + "(" + _dafny.String(data.Cf75) + ")"
		}
	case D15_DC45:
		{
			return "D15.DC45" + "(" + _dafny.String(data.Cf80) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC43:
		{
			data2, ok := other.Get_().(D15_DC43)
			return ok && data1.Cf76.Equals(data2.Cf76) && data1.Cf77 == data2.Cf77
		}
	case D15_DC44:
		{
			data2, ok := other.Get_().(D15_DC44)
			return ok && data1.Cf78 == data2.Cf78 && data1.Cf79.Cmp(data2.Cf79) == 0
		}
	case D15_DC42:
		{
			data2, ok := other.Get_().(D15_DC42)
			return ok && _dafny.AreEqual(data1.Cf75, data2.Cf75)
		}
	case D15_DC45:
		{
			data2, ok := other.Get_().(D15_DC45)
			return ok && data1.Cf80.Equals(data2.Cf80)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D15) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D15)
	return ok && _this.Equals(typed)
}

func Type_D15_() _dafny.TypeDescriptor {
	return type_D15_{}
}

type type_D15_ struct {
}

func (_this type_D15_) Default() interface{} {
	return Companion_D15_.Default()
}

func (_this type_D15_) String() string {
	return "main.D15"
}
func (_this D15) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D15{}

// End of datatype D15

// Definition of datatype D16
type D16 struct {
	Data_D16_
}

func (_this D16) Get_() Data_D16_ {
	return _this.Data_D16_
}

type Data_D16_ interface {
	isD16()
}

type CompanionStruct_D16_ struct {
}

var Companion_D16_ = CompanionStruct_D16_{}

type D16_DC46 struct {
	Cf81 _dafny.Map
}

func (D16_DC46) isD16() {}

func (CompanionStruct_D16_) Create_DC46_(Cf81 _dafny.Map) D16 {
	return D16{D16_DC46{Cf81}}
}

func (_this D16) Is_DC46() bool {
	_, ok := _this.Get_().(D16_DC46)
	return ok
}

func (CompanionStruct_D16_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D16) Dtor_cf81() _dafny.Map {
	return _this.Get_().(D16_DC46).Cf81
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC46:
		{
			return "D16.DC46" + "(" + _dafny.String(data.Cf81) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC46:
		{
			data2, ok := other.Get_().(D16_DC46)
			return ok && data1.Cf81.Equals(data2.Cf81)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D16) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D16)
	return ok && _this.Equals(typed)
}

func Type_D16_() _dafny.TypeDescriptor {
	return type_D16_{}
}

type type_D16_ struct {
}

func (_this type_D16_) Default() interface{} {
	return Companion_D16_.Default()
}

func (_this type_D16_) String() string {
	return "main.D16"
}
func (_this D16) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D16{}

// End of datatype D16

// Definition of trait T0
type T0 interface {
	String() string
}
type CompanionStruct_T0_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T0_ = CompanionStruct_T0_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T0_) CastTo_(x interface{}) T0 {
	var t T0
	t, _ = x.(T0)
	return t
}

// End of trait T0

// Definition of trait T1
type T1 interface {
	String() string
	M1(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) (bool, bool)
}
type CompanionStruct_T1_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T1_ = CompanionStruct_T1_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T1_) CastTo_(x interface{}) T1 {
	var t T1
	t, _ = x.(T1)
	return t
}

// End of trait T1

// Definition of class GlobalState
type GlobalState struct {
	F1   _dafny.Sequence
	F2   _dafny.Array
	F11  bool
	F15  bool
	F16  _dafny.Sequence
	F17  bool
	_f0  _dafny.CodePoint
	_f3  bool
	_f4  _dafny.Sequence
	_f5  _dafny.CodePoint
	_f6  _dafny.Int
	_f7  bool
	_f8  _dafny.Sequence
	_f9  bool
	_f10 bool
	_f12 _dafny.Map
	_f13 bool
	_f14 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.EmptySeq
	_this.F2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F11 = false
	_this.F15 = false
	_this.F16 = _dafny.EmptySeq
	_this.F17 = false
	_this._f0 = _dafny.CodePoint('D')
	_this._f3 = false
	_this._f4 = _dafny.EmptySeq
	_this._f5 = _dafny.CodePoint('D')
	_this._f6 = _dafny.Zero
	_this._f7 = false
	_this._f8 = _dafny.EmptySeq
	_this._f9 = false
	_this._f10 = false
	_this._f12 = _dafny.EmptyMap
	_this._f13 = false
	_this._f14 = _dafny.Zero
	return &_this
}

type CompanionStruct_GlobalState_ struct {
}

var Companion_GlobalState_ = CompanionStruct_GlobalState_{}

func (_this *GlobalState) Equals(other *GlobalState) bool {
	return _this == other
}

func (_this *GlobalState) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*GlobalState)
	return ok && _this.Equals(other)
}

func (*GlobalState) String() string {
	return "_module.GlobalState"
}

func Type_GlobalState_() _dafny.TypeDescriptor {
	return type_GlobalState_{}
}

type type_GlobalState_ struct {
}

func (_this type_GlobalState_) Default() interface{} {
	return (*GlobalState)(nil)
}

func (_this type_GlobalState_) String() string {
	return "main.GlobalState"
}
func (_this *GlobalState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &GlobalState{}

func (_this *GlobalState) Ctor__(f0 _dafny.CodePoint, f1 _dafny.Sequence, f2 _dafny.Array, f3 bool, f4 _dafny.Sequence, f5 _dafny.CodePoint, f6 _dafny.Int, f7 bool, f8 _dafny.Sequence, f9 bool, f10 bool, f11 bool, f12 _dafny.Map, f13 bool, f14 _dafny.Int, f15 bool, f16 _dafny.Sequence, f17 bool) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
		(_this).F16 = f16
		(_this).F17 = f17
	}
}
func (_this *GlobalState) F0() _dafny.CodePoint {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Sequence {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.CodePoint {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Sequence {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() bool {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F12() _dafny.Map {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() bool {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() _dafny.Int {
	{
		return _this._f14
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f23 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f23 = false
	return &_this
}

type CompanionStruct_C0_ struct {
}

var Companion_C0_ = CompanionStruct_C0_{}

func (_this *C0) Equals(other *C0) bool {
	return _this == other
}

func (_this *C0) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C0)
	return ok && _this.Equals(other)
}

func (*C0) String() string {
	return "_module.C0"
}

func Type_C0_() _dafny.TypeDescriptor {
	return type_C0_{}
}

type type_C0_ struct {
}

func (_this type_C0_) Default() interface{} {
	return (*C0)(nil)
}

func (_this type_C0_) String() string {
	return "main.C0"
}
func (_this *C0) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) Ctor__(f23 bool) {
	{
		(_this)._f23 = f23
	}
}
func (_this *C0) Fm14(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		if (func() bool {
			if (_this).F23() {
				return (_this).F23()
			}
			return false
		})() {
			return (_dafny.IntOfUint32((_dafny.SeqOf((_this).F23())).Cardinality())).Minus(_dafny.IntOfInt64(493))
		} else {
			return (_dafny.IntOfInt64(969)).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("opfaxvbk"))).Cardinality())
		}
	}
}
func (_this *C0) Fm15(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) bool {
	{
		return !((_this).F23()) || ((Companion_D0_.Create_DC0_((_this).F23())).Dtor_cf0())
	}
}
func (_this *C0) F23() bool {
	{
		return _this._f23
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f22 _dafny.Array
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f22 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	return &_this
}

type CompanionStruct_C1_ struct {
}

var Companion_C1_ = CompanionStruct_C1_{}

func (_this *C1) Equals(other *C1) bool {
	return _this == other
}

func (_this *C1) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C1)
	return ok && _this.Equals(other)
}

func (*C1) String() string {
	return "_module.C1"
}

func Type_C1_() _dafny.TypeDescriptor {
	return type_C1_{}
}

type type_C1_ struct {
}

func (_this type_C1_) Default() interface{} {
	return (*C1)(nil)
}

func (_this type_C1_) String() string {
	return "main.C1"
}
func (_this *C1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__(f22 _dafny.Array) {
	{
		(_this)._f22 = f22
	}
}
func (_this *C1) M7(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) {
	{
		var _464_v0 _dafny.Sequence
		_ = _464_v0
		_464_v0 = _dafny.SeqOf(true)
		var _465_v1 bool
		_ = _465_v1
		_465_v1 = true
		var _466_v2 _dafny.Sequence
		_ = _466_v2
		_466_v2 = _dafny.UnicodeSeqOfUtf8Bytes("nlrnbcijs")
		var _hi5 _dafny.Int = p1
		_ = _hi5
		for _467_i0 := (Companion_Default___.Fm9(p0, _dafny.IntOfUint32((_464_v0).Cardinality()), _465_v1, globalState)).Plus((Companion_D1_.Create_DC5_(_466_v2, p0, _dafny.IntOfUint32((_466_v2).Cardinality()))).Dtor_cf9()); _467_i0.Cmp(_hi5) < 0; _467_i0 = _467_i0.Plus(_dafny.One) {
			var _468_v3 _dafny.Int
			_ = _468_v3
			_468_v3 = _dafny.IntOfInt64(596)
			_468_v3 = Companion_Default___.SafeModuloInt(p1, _468_v3)
			var _469_v5 _dafny.CodePoint
			_ = _469_v5
			_469_v5 = _dafny.CodePoint('j')
			var _470_v6 _dafny.Set
			_ = _470_v6
			_470_v6 = _dafny.SetOf(_469_v5)
			var _471_v7 _dafny.Map
			_ = _471_v7
			_471_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_465_v1, _465_v1, _465_v1), (_this).F22())
			var _472_v8 _dafny.Map
			_ = _472_v8
			_472_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
				var _coll21 = _dafny.NewMapBuilder()
				_ = _coll21
				for _iter22 := _dafny.Iterate((_470_v6).Elements()); ; {
					_compr_21, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _473_v4 _dafny.CodePoint
					_473_v4 = interface{}(_compr_21).(_dafny.CodePoint)
					if (_470_v6).Contains(_473_v4) {
						_coll21.Add(_473_v4, _467_i0)
					}
				}
				return _coll21.ToMap()
			}(), (_471_v7).Cardinality())
			var _474_v9 _dafny.Map
			_ = _474_v9
			_474_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), _dafny.IntOfInt64(706))
			_472_v8 = (_472_v8).Update(_474_v9, p1)
			var _475_v10 _dafny.Array
			_ = _475_v10
			var _nw80 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
			_ = _nw80
			_475_v10 = _nw80
			var _476_v11 D1
			_ = _476_v11
			_476_v11 = Companion_D1_.Create_DC4_(p1, _465_v1, _475_v10)
			var _source7 D1 = _476_v11
			_ = _source7
			if _source7.Is_DC4() {
				var _477___mcc_h0 _dafny.Int = _source7.Get_().(D1_DC4).Cf4
				_ = _477___mcc_h0
				var _478___mcc_h1 bool = _source7.Get_().(D1_DC4).Cf5
				_ = _478___mcc_h1
				var _479___mcc_h2 _dafny.Array = _source7.Get_().(D1_DC4).Cf6
				_ = _479___mcc_h2
				var _480_cf6 _dafny.Array = _479___mcc_h2
				_ = _480_cf6
				var _481_cf5 bool = _478___mcc_h1
				_ = _481_cf5
				var _482_cf4 _dafny.Int = _477___mcc_h0
				_ = _482_cf4
				var _483_v12 _dafny.Set
				_ = _483_v12
				_483_v12 = _dafny.SetOf(p1, p1)
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_475_v10), 0))
				_ = _index49
				(_475_v10).ArraySet1((_dafny.SetOf(p0, _468_v3, _468_v3)).IsSubsetOf(_483_v12), (_index49).Int())
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_475_v10), 0))
				_ = _index50
				(_475_v10).ArraySet1(!(_465_v1), (_index50).Int())
				var _484_v13 _dafny.MultiSet
				_ = _484_v13
				_484_v13 = _dafny.MultiSetOf(p1)
				var _485_v14 D2
				_ = _485_v14
				_485_v14 = Companion_D2_.Create_DC7_(_484_v13)
				_465_v1 = (((_485_v14).Dtor_cf12()).Cardinality()).Cmp((p1).Plus(p1)) == 0
				var _486_v15 _dafny.Map
				_ = _486_v15
				_486_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_466_v2).Cardinality()))
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_475_v10), 0))
				_ = _index51
				(_475_v10).ArraySet1((Companion_Default___.SafeModuloInt(_482_cf4, (_486_v15).Cardinality())).Cmp((_468_v3).Plus(p1)) == 0, (_index51).Int())
				var _487_v16 _dafny.Map
				_ = _487_v16
				_487_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(624))
				var _488_v17 D0
				_ = _488_v17
				_488_v17 = Companion_D0_.Create_DC1_((_475_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_475_v10), 0))).Int()).(bool))
				var _489_v18 D0
				_ = _489_v18
				_489_v18 = Companion_D0_.Create_DC2_(_488_v17)
				var _pat_let_tv11 = _465_v1
				_ = _pat_let_tv11
				var _pat_let_tv12 = _475_v10
				_ = _pat_let_tv12
				var _pat_let_tv13 = _475_v10
				_ = _pat_let_tv13
				var _pat_let_tv14 = globalState
				_ = _pat_let_tv14
				var _490_v19 _dafny.Map
				_ = _490_v19
				_490_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let7_0 D0) D0 {
					return func(_491_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let8_0 D0) D0 {
							return func(_492_dt__update_hcf2_h0 D0) D0 {
								return Companion_D0_.Create_DC2_(_492_dt__update_hcf2_h0)
							}(_pat_let8_0)
						}(Companion_Default___.Fm11(_467_i0, _pat_let_tv11, (_pat_let_tv13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_pat_let_tv12), 0))).Int()).(bool), _467_i0, _pat_let_tv14))
					}(_pat_let7_0)
				}(_489_v18), _489_v18)
				var _nwElement0_10 bool = false
				_ = _nwElement0_10
				var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(15))
				_ = _nw81
				(_nw81).ArraySet1(_nwElement0_10, 0)
				(_nw81).ArraySet1((_468_v3).Cmp(_467_i0) > 0, 1)
				(_nw81).ArraySet1((_475_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_475_v10), 0))).Int()).(bool), 2)
				(_nw81).ArraySet1((!(_465_v1)) == ((_475_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_475_v10), 0))).Int()).(bool)), 3)
				(_nw81).ArraySet1((_dafny.IntOfInt64(-192)).Cmp(p0) > 0, 4)
				(_nw81).ArraySet1(((_475_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_475_v10), 0))).Int()).(bool)) == (_481_cf5), 5)
				(_nw81).ArraySet1((_475_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_475_v10), 0))).Int()).(bool), 6)
				(_nw81).ArraySet1(!(_487_v16).Equals(Companion_Default___.Fm10(globalState)), 7)
				(_nw81).ArraySet1(true, 8)
				(_nw81).ArraySet1(_465_v1, 9)
				(_nw81).ArraySet1((Companion_Default___.Fm9(_482_cf4, _467_i0, _481_cf5, globalState)).Cmp((_490_v19).Cardinality()) < 0, 10)
				(_nw81).ArraySet1(_481_cf5, 11)
				(_nw81).ArraySet1((_475_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_475_v10), 0))).Int()).(bool), 12)
				(_nw81).ArraySet1(((_475_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_475_v10), 0))).Int()).(bool)) || (_465_v1), 13)
				(_nw81).ArraySet1((func() bool {
					if _465_v1 {
						return _481_cf5
					}
					return _481_cf5
				})(), 14)
				_480_cf6 = _nw81
			} else if _source7.Is_DC5() {
				var _493___mcc_h3 _dafny.Sequence = _source7.Get_().(D1_DC5).Cf7
				_ = _493___mcc_h3
				var _494___mcc_h4 _dafny.Int = _source7.Get_().(D1_DC5).Cf8
				_ = _494___mcc_h4
				var _495___mcc_h5 _dafny.Int = _source7.Get_().(D1_DC5).Cf9
				_ = _495___mcc_h5
				var _496_cf9 _dafny.Int = _495___mcc_h5
				_ = _496_cf9
				var _497_cf8 _dafny.Int = _494___mcc_h4
				_ = _497_cf8
				var _498_cf7 _dafny.Sequence = _493___mcc_h3
				_ = _498_cf7
				_469_v5 = _469_v5
				var _499_v20 _dafny.Map
				_ = _499_v20
				_499_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_465_v1, _464_v0)
				_499_v20 = (_499_v20).Update(Companion_Default___.Fm0(_465_v1, _469_v5, _465_v1, globalState), _464_v0)
				_498_cf7 = _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ss"), (Companion_Default___.SafeIndex(_467_i0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ss")).Cardinality()))).Uint32(), _469_v5)
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_475_v10), 0))
				_ = _index52
				(_475_v10).ArraySet1(_465_v1, (_index52).Int())
				var _500_v21 _dafny.Set
				_ = _500_v21
				_500_v21 = _dafny.SetOf(p0)
				var _501_v22 D0
				_ = _501_v22
				_501_v22 = Companion_D0_.Create_DC0_(_465_v1)
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_475_v10), 0))
				_ = _index53
				(_475_v10).ArraySet1((func() bool {
					if (_500_v21).IsProperSubsetOf(_500_v21) {
						return (_501_v22).Dtor_cf0()
					}
					return _465_v1
				})(), (_index53).Int())
			} else if _source7.Is_DC6() {
				var _502___mcc_h6 _dafny.Int = _source7.Get_().(D1_DC6).Cf10
				_ = _502___mcc_h6
				var _503___mcc_h7 _dafny.Int = _source7.Get_().(D1_DC6).Cf11
				_ = _503___mcc_h7
				var _504_cf11 _dafny.Int = _503___mcc_h7
				_ = _504_cf11
				var _505_cf10 _dafny.Int = _502___mcc_h6
				_ = _505_cf10
				var _506_v23 _dafny.Set
				_ = _506_v23
				_506_v23 = _dafny.SetOf(_465_v1)
				var _507_v24 _dafny.Map
				_ = _507_v24
				_507_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_469_v5, (_506_v23).IsSubsetOf(_506_v23))
				var _508_v25 _dafny.MultiSet
				_ = _508_v25
				_508_v25 = _dafny.MultiSetOf(_dafny.IntOfInt64(69), _504_cf11)
				_507_v24 = (_507_v24).Update(_469_v5, (_dafny.MultiSetOf(_504_cf11)).IsSubsetOf(_508_v25))
				(globalState).F11 = _465_v1
				var _509_v26 _dafny.Sequence
				_ = _509_v26
				_509_v26 = _dafny.SeqOf(_466_v2)
				var _510_v27 _dafny.Map
				_ = _510_v27
				_510_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _467_i0)
				_469_v5 = Companion_Default___.Fm12((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p1, p1))), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_466_v2, (_509_v26).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_468_v3), _dafny.IntOfUint32((_509_v26).Cardinality()))).Uint32()).(_dafny.Sequence)), (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_510_v27).Contains((_this).F22()) {
						return (_510_v27).Get((_this).F22()).(_dafny.Int)
					}
					return _504_cf11
				})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_466_v2, (_509_v26).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_468_v3), _dafny.IntOfUint32((_509_v26).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()))).Uint32(), _469_v5), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_467_i0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_466_v2, (_509_v26).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_468_v3), _dafny.IntOfUint32((_509_v26).Cardinality()))).Uint32()).(_dafny.Sequence)), (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_510_v27).Contains((_this).F22()) {
						return (_510_v27).Get((_this).F22()).(_dafny.Int)
					}
					return _504_cf11
				})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_466_v2, (_509_v26).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_468_v3), _dafny.IntOfUint32((_509_v26).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()))).Uint32(), _469_v5)).Cardinality()))).Uint32(), _469_v5), (_509_v26).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_509_v26).Cardinality()))).Uint32()).(_dafny.Sequence), globalState)
				(globalState).F11 = _465_v1
			} else {
				var _511___mcc_h8 _dafny.Array = _source7.Get_().(D1_DC3).Cf3
				_ = _511___mcc_h8
				var _512_cf3 _dafny.Array = _511___mcc_h8
				_ = _512_cf3
				var _513_v28 _dafny.Array
				_ = _513_v28
				var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
				_ = _nw82
				_513_v28 = _nw82
				var _514_v29 _dafny.Sequence
				_ = _514_v29
				_514_v29 = _dafny.SeqOf(_467_i0)
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_513_v28), 0))
				_ = _index54
				(_513_v28).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_468_v3, p1), _514_v29), (_index54).Int())
				var _515_v30 _dafny.Array
				_ = _515_v30
				var _len0_13 _dafny.Int = _dafny.One
				_ = _len0_13
				var _nw83 _dafny.Array
				_ = _nw83
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw83 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Set = (func(_516_v1 bool) func(_dafny.Int) _dafny.Set {
						return func(_517_i1 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(_516_v1, _516_v1)
						}
					})(_465_v1)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw83 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw83).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw83).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_515_v30 = _nw83
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_513_v28), 0))
				_ = _index55
				var _rhs54 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(266))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg29 _dafny.Int) interface{} {
						return coer29(arg29)
					}
				}((func(_518_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_519_i2 _dafny.Int) _dafny.CodePoint {
						return _518_v5
					}
				})(_469_v5)))
				_ = _rhs54
				var _rhs55 _dafny.Sequence = _514_v29
				_ = _rhs55
				var _rhs56 _dafny.Array = _515_v30
				_ = _rhs56
				var _rhs57 _dafny.Int = (_dafny.IntOfUint32((_466_v2).Cardinality())).Plus((func() _dafny.Int {
					if _465_v1 {
						return _468_v3
					}
					return _dafny.IntOfInt64(578)
				})())
				_ = _rhs57
				var _lhs27 _dafny.Array = _513_v28
				_ = _lhs27
				var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_513_v28), 0))
				_ = _lhs28
				_466_v2 = _rhs54
				(_lhs27).ArraySet1(_rhs55, (_lhs28).Int())
				_515_v30 = _rhs56
				_468_v3 = _rhs57
				var _520_v31 _dafny.Sequence
				_ = _520_v31
				_520_v31 = _dafny.SeqOf((_513_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_513_v28), 0))).Int()).(_dafny.Sequence), (_513_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_513_v28), 0))).Int()).(_dafny.Sequence))
				var _521_v32 _dafny.Map
				_ = _521_v32
				_521_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_520_v31, Companion_Default___.SafeDivisionInt(p1, Companion_Default___.Fm9(p0, _dafny.IntOfUint32((_466_v2).Cardinality()), _465_v1, globalState)))
				var _522_v33 _dafny.Set
				_ = _522_v33
				_522_v33 = _dafny.SetOf(_dafny.IntOfInt64(464))
				_521_v32 = (_521_v32).Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_513_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_513_v28), 0))).Int()).(_dafny.Sequence), _514_v29), (Companion_Default___.SafeIndex((_522_v33).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_513_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_513_v28), 0))).Int()).(_dafny.Sequence), _514_v29)).Cardinality()))).Uint32(), (_513_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_513_v28), 0))).Int()).(_dafny.Sequence)), p0)
				_468_v3 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm13(_467_i0, _467_i0, globalState), _520_v31)).Cardinality())
				var _523_v34 _dafny.Map
				_ = _523_v34
				_523_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_v3, _465_v1)
				var _524_v35 _dafny.Set
				_ = _524_v35
				_524_v35 = _dafny.SetOf(_523_v34, _523_v34, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-625), _465_v1))
				var _525_v37 D3
				_ = _525_v37
				_525_v37 = Companion_D3_.Create_DC9_(_dafny.SetOf(func() _dafny.Map {
					var _coll22 = _dafny.NewMapBuilder()
					_ = _coll22
					for _iter23 := _dafny.Iterate((_514_v29).Elements()); ; {
						_compr_22, _ok23 := _iter23()
						if !_ok23 {
							break
						}
						var _526_v36 _dafny.Int
						_526_v36 = interface{}(_compr_22).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_514_v29, _526_v36) {
							_coll22.Add((_526_v36).Times(p0), _465_v1)
						}
					}
					return _coll22.ToMap()
				}(), _523_v34, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_465_v1)).Cardinality(), _465_v1)))
				(globalState).F15 = ((_524_v35).Intersection(_524_v35)).IsSubsetOf((_525_v37).Dtor_cf16())
			}
			_468_v3 = p1
		}
		_465_v1 = true
		var _527_v38 _dafny.Map
		_ = _527_v38
		_527_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _465_v1)
		var _528_v39 _dafny.MultiSet
		_ = _528_v39
		_528_v39 = _dafny.MultiSetOf(_465_v1)
		var _529_v40 _dafny.Sequence
		_ = _529_v40
		_529_v40 = _dafny.SeqOf(_528_v39)
		_527_v38 = (_527_v38).Update(p1, (p1).Cmp(((_529_v40).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_529_v40).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()) > 0)
		var _530_v41 _dafny.Array
		_ = _530_v41
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_14
		var _nw84 _dafny.Array
		_ = _nw84
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw84 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) bool = (func(_531_p0 _dafny.Int) func(_dafny.Int) bool {
				return func(_532_i3 _dafny.Int) bool {
					return (_531_p0).Cmp(_dafny.IntOfInt64(-797)) >= 0
				}
			})(p0)
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw84 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw84).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw84).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_530_v41 = _nw84
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_530_v41), 0))
		_ = _index56
		(_530_v41).ArraySet1(!(_465_v1) || (_465_v1), (_index56).Int())
		var _533_v42 _dafny.MultiSet
		_ = _533_v42
		_533_v42 = _dafny.MultiSetOf(p1)
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_530_v41), 0))
		_ = _index57
		(_530_v41).ArraySet1(!(_533_v42).Contains(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-892), p1)), (_index57).Int())
		(globalState).F15 = false
		var _534_i4 _dafny.Int
		_ = _534_i4
		_534_i4 = _dafny.Zero
		{
			for (_530_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_530_v41), 0))).Int()).(bool) {
				{
					if (_534_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_534_i4 = (_534_i4).Plus(_dafny.One)
					var _535_v43 *C0
					_ = _535_v43
					var _nw85 *C0 = New_C0_()
					_ = _nw85
					_nw85.Ctor__(_465_v1)
					_535_v43 = _nw85
					var _536_v44 _dafny.Int
					_ = _536_v44
					_536_v44 = _dafny.IntOfInt64(-941)
					var _537_v45 _dafny.Sequence
					_ = _537_v45
					_537_v45 = _dafny.SeqOf(p0)
					_536_v44 = _dafny.IntOfUint32((_537_v45).Cardinality())
					_536_v44 = p0
					_465_v1 = _465_v1
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
	}
}
func (_this *C1) M8(p0 _dafny.Int, p1 bool, globalState *GlobalState) (_dafny.Set, bool, _dafny.Int, _dafny.Map) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Map = _dafny.EmptyMap
		_ = r3
		if p1 {
			var _538_v0 _dafny.Sequence
			_ = _538_v0
			_538_v0 = _dafny.SeqOf(p0, p0, p0)
			var _539_v1 _dafny.Map
			_ = _539_v1
			_539_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(798))
			var _540_v2 _dafny.Set
			_ = _540_v2
			_540_v2 = _dafny.SetOf(p0, (func() _dafny.Int {
				if (_539_v1).Contains(false) {
					return (_539_v1).Get(false).(_dafny.Int)
				}
				return p0
			})())
			var _541_v3 _dafny.Sequence
			_ = _541_v3
			_541_v3 = _dafny.SeqOf(p1)
			var _542_v4 *C0
			_ = _542_v4
			var _nw86 *C0 = New_C0_()
			_ = _nw86
			_nw86.Ctor__(p1)
			_542_v4 = _nw86
			var _543_v5 _dafny.Map
			_ = _543_v5
			_543_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _542_v4)
			var _544_v6 _dafny.CodePoint
			_ = _544_v6
			_544_v6 = _dafny.CodePoint('m')
			var _545_v7 _dafny.Array
			_ = _545_v7
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_15
			var _nw87 _dafny.Array
			_ = _nw87
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw87 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Sequence = (func(_546_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_547_i0 _dafny.Int) _dafny.Sequence {
						return _546_v3
					}
				})(_541_v3)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw87 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw87).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw87).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_545_v7 = _nw87
			var _548_v8 D1
			_ = _548_v8
			_548_v8 = Companion_D1_.Create_DC3_(_545_v7)
			var _549_v9 _dafny.Map
			_ = _549_v9
			_549_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_548_v8, p1)
			var _550_v10 _dafny.Map
			_ = _550_v10
			_550_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_543_v5).Update((_542_v4).F23(), _542_v4), Companion_Default___.Fm0((_542_v4).F23(), _544_v6, (func() bool {
				if (_549_v9).Contains(Companion_D1_.Create_DC3_(_545_v7)) {
					return (_549_v9).Get(Companion_D1_.Create_DC3_(_545_v7)).(bool)
				}
				return (_542_v4).F23()
			})(), globalState))
			var _551_v11 _dafny.Sequence
			_ = _551_v11
			_551_v11 = _dafny.UnicodeSeqOfUtf8Bytes("ufps")
			var _552_v12 _dafny.MultiSet
			_ = _552_v12
			_552_v12 = _dafny.MultiSetOf((_542_v4).Fm14(p0, globalState))
			var _553_v13 _dafny.Sequence
			_ = _553_v13
			_553_v13 = _dafny.SeqOf(_552_v12)
			var _554_v14 _dafny.Array
			_ = _554_v14
			var _nwElement0_11 bool = p1
			_ = _nwElement0_11
			var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(24))
			_ = _nw88
			(_nw88).ArraySet1(_nwElement0_11, 0)
			(_nw88).ArraySet1(p1, 1)
			(_nw88).ArraySet1((p1) && (p1), 2)
			(_nw88).ArraySet1(p1, 3)
			(_nw88).ArraySet1(true, 4)
			(_nw88).ArraySet1(!(_dafny.Companion_Sequence_.Contains(_538_v0, p0)), 5)
			(_nw88).ArraySet1(p1, 6)
			(_nw88).ArraySet1(p1, 7)
			(_nw88).ArraySet1(p1, 8)
			(_nw88).ArraySet1(!(_540_v2).Equals(_dafny.SetOf(p0, p0, p0)), 9)
			(_nw88).ArraySet1(p1, 10)
			(_nw88).ArraySet1(p1, 11)
			(_nw88).ArraySet1((_541_v3).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_541_v3).Cardinality()))).Uint32()).(bool), 12)
			(_nw88).ArraySet1(p1, 13)
			(_nw88).ArraySet1(p1, 14)
			(_nw88).ArraySet1((func() bool {
				if (_550_v10).Contains(_543_v5) {
					return (_550_v10).Get(_543_v5).(bool)
				}
				return p1
			})(), 15)
			(_nw88).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_538_v0, _538_v0), 16)
			(_nw88).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(514))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}((func(_555_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_556_i1 _dafny.Int) _dafny.CodePoint {
					return _555_v6
				}
			})(_544_v6))), _544_v6), 17)
			(_nw88).ArraySet1(!((_dafny.IntOfUint32((_551_v11).Cardinality())).Cmp(_dafny.IntOfUint32((_553_v13).Cardinality())) <= 0), 18)
			(_nw88).ArraySet1(p1, 19)
			(_nw88).ArraySet1(true, 20)
			(_nw88).ArraySet1((Companion_Default___.Fm0((_542_v4).F23(), (_551_v11).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_551_v11).Cardinality()))).Uint32()).(_dafny.CodePoint), p1, globalState)) == ((_542_v4).F23()), 21)
			(_nw88).ArraySet1(p1, 22)
			(_nw88).ArraySet1(p1, 23)
			_554_v14 = _nw88
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_554_v14), 0))
			_ = _index58
			(_554_v14).ArraySet1((_542_v4).F23(), (_index58).Int())
			var _557_v15 _dafny.Map
			_ = _557_v15
			_557_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_542_v4).F23())
			var _558_v16 _dafny.Set
			_ = _558_v16
			_558_v16 = _dafny.SetOf(_540_v2)
			var _559_v17 _dafny.Set
			_ = _559_v17
			_559_v17 = _dafny.SetOf(p1)
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_554_v14), 0))
			_ = _index59
			(_554_v14).ArraySet1(!((_542_v4).F23()) || ((_dafny.SetOf(p1, (func() bool {
				if (_557_v15).Contains((_558_v16).Cardinality()) {
					return (_557_v15).Get((_558_v16).Cardinality()).(bool)
				}
				return false
			})())).IsDisjointFrom(_559_v17)), (_index59).Int())
			var _560_v18 D0
			_ = _560_v18
			_560_v18 = Companion_D0_.Create_DC0_(false)
			_560_v18 = _560_v18
			var _561_v19 *C0
			_ = _561_v19
			var _nw89 *C0 = New_C0_()
			_ = _nw89
			_nw89.Ctor__((func() bool {
				if p1 {
					return !(p1)
				}
				return (_554_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_554_v14), 0))).Int()).(bool)
			})())
			_561_v19 = _nw89
			_544_v6 = _544_v6
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen(((_this).F22()), 0))
			_ = _index60
			((_this).F22()).ArraySet1(_dafny.IntOfUint32((_541_v3).Cardinality()), (_index60).Int())
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen(((_this).F22()), 0))
			_ = _index61
			((_this).F22()).ArraySet1((p0).Plus(_dafny.IntOfUint32((_dafny.SeqOf(p0, _dafny.IntOfUint32((_dafny.SeqOf((_554_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_554_v14), 0))).Int()).(bool), (_542_v4).F23())).Cardinality()))).Cardinality())), (_index61).Int())
		} else {
			var _562_v20 _dafny.Map
			_ = _562_v20
			_562_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _563_v21 _dafny.Set
			_ = _563_v21
			_563_v21 = _dafny.SetOf(_562_v20)
			var _source8 D3 = Companion_D3_.Create_DC9_((_563_v21).Union(_563_v21))
			_ = _source8
			if _source8.Is_DC10() {
				var _564___mcc_h0 _dafny.Int = _source8.Get_().(D3_DC10).Cf17
				_ = _564___mcc_h0
				var _565_cf17 _dafny.Int = _564___mcc_h0
				_ = _565_cf17
				var _566_v22 _dafny.Sequence
				_ = _566_v22
				_566_v22 = _dafny.UnicodeSeqOfUtf8Bytes("cnnnwq")
				var _567_v23 _dafny.MultiSet
				_ = _567_v23
				_567_v23 = _dafny.MultiSetOf(_566_v22)
				var _568_v24 _dafny.Sequence
				_ = _568_v24
				_568_v24 = _dafny.SeqOf(_566_v22)
				r2 = ((_dafny.Zero).Minus((_565_cf17).Plus((_567_v23).Cardinality()))).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_568_v24, _568_v24)).Cardinality()))
				var _569_v25 _dafny.CodePoint
				_ = _569_v25
				_569_v25 = _dafny.CodePoint('r')
				_562_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(871), Companion_Default___.Fm0(p1, _569_v25, p1, globalState))
				var _570_v26 D1
				_ = _570_v26
				_570_v26 = Companion_D1_.Create_DC5_(_566_v22, _565_cf17, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-709))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}((func(_571_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_572_i2 _dafny.Int) _dafny.CodePoint {
						return _571_v25
					}
				})(_569_v25)))).Cardinality()))
				_565_cf17 = (_570_v26).Dtor_cf8()
				var _573_v27 _dafny.Sequence
				_ = _573_v27
				_573_v27 = _dafny.SeqOf(!(p1))
				var _574_v28 _dafny.Array
				_ = _574_v28
				var _nwElement0_12 _dafny.Sequence = _573_v27
				_ = _nwElement0_12
				var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(18))
				_ = _nw90
				(_nw90).ArraySet1(_nwElement0_12, 0)
				(_nw90).ArraySet1(_573_v27, 1)
				(_nw90).ArraySet1(_573_v27, 2)
				(_nw90).ArraySet1(_573_v27, 3)
				(_nw90).ArraySet1(_573_v27, 4)
				(_nw90).ArraySet1(_dafny.Companion_Sequence_.Update(_573_v27, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_573_v27).Cardinality()))).Uint32(), p1), 5)
				(_nw90).ArraySet1(_dafny.SeqOf(!(p1)), 6)
				(_nw90).ArraySet1(_573_v27, 7)
				(_nw90).ArraySet1(_573_v27, 8)
				(_nw90).ArraySet1(_573_v27, 9)
				(_nw90).ArraySet1(_573_v27, 10)
				(_nw90).ArraySet1(_573_v27, 11)
				(_nw90).ArraySet1(_573_v27, 12)
				(_nw90).ArraySet1(_573_v27, 13)
				(_nw90).ArraySet1(_573_v27, 14)
				(_nw90).ArraySet1(_573_v27, 15)
				(_nw90).ArraySet1(_573_v27, 16)
				(_nw90).ArraySet1(_573_v27, 17)
				_574_v28 = _nw90
				var _575_v29 D1
				_ = _575_v29
				_575_v29 = Companion_D1_.Create_DC3_(_574_v28)
				var _pat_let_tv15 = _574_v28
				_ = _pat_let_tv15
				var _pat_let_tv16 = _574_v28
				_ = _pat_let_tv16
				var _rhs58 D1 = func(_pat_let9_0 D1) D1 {
					return func(_576_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let10_0 _dafny.Array) D1 {
							return func(_577_dt__update_hcf3_h0 _dafny.Array) D1 {
								return Companion_D1_.Create_DC3_(_577_dt__update_hcf3_h0)
							}(_pat_let10_0)
						}((func() _dafny.Array {
							if false {
								return _pat_let_tv15
							}
							return _pat_let_tv16
						})())
					}(_pat_let9_0)
				}(_575_v29)
				_ = _rhs58
				var _rhs59 bool = p1
				_ = _rhs59
				var _rhs60 _dafny.Int = (p0).Minus((_565_cf17).Times(_565_cf17))
				_ = _rhs60
				var _rhs61 bool = p1
				_ = _rhs61
				var _rhs62 _dafny.Sequence = _566_v22
				_ = _rhs62
				var _lhs29 *GlobalState = globalState
				_ = _lhs29
				var _lhs30 *GlobalState = globalState
				_ = _lhs30
				_575_v29 = _rhs58
				_lhs29.F15 = _rhs59
				_565_cf17 = _rhs60
				_lhs30.F17 = _rhs61
				_566_v22 = _rhs62
			} else if _source8.Is_DC11() {
				var _578___mcc_h1 _dafny.Map = _source8.Get_().(D3_DC11).Cf18
				_ = _578___mcc_h1
				var _579___mcc_h2 _dafny.Int = _source8.Get_().(D3_DC11).Cf19
				_ = _579___mcc_h2
				var _580___mcc_h3 _dafny.Array = _source8.Get_().(D3_DC11).Cf20
				_ = _580___mcc_h3
				var _581_cf20 _dafny.Array = _580___mcc_h3
				_ = _581_cf20
				var _582_cf19 _dafny.Int = _579___mcc_h2
				_ = _582_cf19
				var _583_cf18 _dafny.Map = _578___mcc_h1
				_ = _583_cf18
				r1 = p1
				var _584_v30 _dafny.CodePoint
				_ = _584_v30
				_584_v30 = _dafny.CodePoint('x')
				var _585_v31 _dafny.Array
				_ = _585_v31
				var _nwElement0_13 _dafny.CodePoint = _584_v30
				_ = _nwElement0_13
				var _nw91 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(9))
				_ = _nw91
				(_nw91).ArraySet1CodePoint(_nwElement0_13, 0)
				(_nw91).ArraySet1CodePoint(_584_v30, 1)
				(_nw91).ArraySet1CodePoint(_584_v30, 2)
				(_nw91).ArraySet1CodePoint(_584_v30, 3)
				(_nw91).ArraySet1CodePoint(_584_v30, 4)
				(_nw91).ArraySet1CodePoint(_584_v30, 5)
				(_nw91).ArraySet1CodePoint(_584_v30, 6)
				(_nw91).ArraySet1CodePoint(_584_v30, 7)
				(_nw91).ArraySet1CodePoint(_584_v30, 8)
				_585_v31 = _nw91
				var _586_v32 _dafny.Sequence
				_ = _586_v32
				_586_v32 = _dafny.UnicodeSeqOfUtf8Bytes("urvmnj")
				var _587_v33 _dafny.Map
				_ = _587_v33
				_587_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_581_cf20, _586_v32)
				var _588_v34 _dafny.MultiSet
				_ = _588_v34
				_588_v34 = _dafny.MultiSetOf(_582_cf19)
				var _nwElement0_14 _dafny.CodePoint = _dafny.CodePoint('i')
				_ = _nwElement0_14
				var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(24))
				_ = _nw92
				(_nw92).ArraySet1CodePoint(_nwElement0_14, 0)
				(_nw92).ArraySet1CodePoint(_584_v30, 1)
				(_nw92).ArraySet1CodePoint(_584_v30, 2)
				(_nw92).ArraySet1CodePoint(_584_v30, 3)
				(_nw92).ArraySet1CodePoint(_584_v30, 4)
				(_nw92).ArraySet1CodePoint(Companion_Default___.Fm12(_582_cf19, (func() _dafny.Sequence {
					if (_587_v33).Contains((_this).F22()) {
						return (_587_v33).Get((_this).F22()).(_dafny.Sequence)
					}
					return _586_v32
				})(), Companion_Default___.Fm16(false, (_588_v34).Cardinality(), globalState), globalState), 5)
				(_nw92).ArraySet1CodePoint(_584_v30, 6)
				(_nw92).ArraySet1CodePoint(_584_v30, 7)
				(_nw92).ArraySet1CodePoint(_584_v30, 8)
				(_nw92).ArraySet1CodePoint(_dafny.CodePoint('m'), 9)
				(_nw92).ArraySet1CodePoint(_584_v30, 10)
				(_nw92).ArraySet1CodePoint(_584_v30, 11)
				(_nw92).ArraySet1CodePoint(_584_v30, 12)
				(_nw92).ArraySet1CodePoint(_584_v30, 13)
				(_nw92).ArraySet1CodePoint(_584_v30, 14)
				(_nw92).ArraySet1CodePoint(_584_v30, 15)
				(_nw92).ArraySet1CodePoint(_584_v30, 16)
				(_nw92).ArraySet1CodePoint(_584_v30, 17)
				(_nw92).ArraySet1CodePoint(_584_v30, 18)
				(_nw92).ArraySet1CodePoint(_584_v30, 19)
				(_nw92).ArraySet1CodePoint(_584_v30, 20)
				(_nw92).ArraySet1CodePoint(_584_v30, 21)
				(_nw92).ArraySet1CodePoint(_584_v30, 22)
				(_nw92).ArraySet1CodePoint(_584_v30, 23)
				_585_v31 = _nw92
				var _589_v35 _dafny.Array
				_ = _589_v35
				var _nwElement0_15 bool = p1
				_ = _nwElement0_15
				var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(12))
				_ = _nw93
				(_nw93).ArraySet1(_nwElement0_15, 0)
				(_nw93).ArraySet1(p1, 1)
				(_nw93).ArraySet1(p1, 2)
				(_nw93).ArraySet1(p1, 3)
				(_nw93).ArraySet1(p1, 4)
				(_nw93).ArraySet1(true, 5)
				(_nw93).ArraySet1(p1, 6)
				(_nw93).ArraySet1(p1, 7)
				(_nw93).ArraySet1(p1, 8)
				(_nw93).ArraySet1(p1, 9)
				(_nw93).ArraySet1(p1, 10)
				(_nw93).ArraySet1(p1, 11)
				_589_v35 = _nw93
				var _590_v36 D1
				_ = _590_v36
				_590_v36 = Companion_D1_.Create_DC5_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-656))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}((func(_591_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_592_i3 _dafny.Int) _dafny.CodePoint {
						return _591_v30
					}
				})(_584_v30))), _582_cf19, _582_cf19)
				var _593_v37 _dafny.Sequence
				_ = _593_v37
				_593_v37 = _dafny.SeqOf(_586_v32, _586_v32, _586_v32, _586_v32, _586_v32)
				var _594_v38 _dafny.MultiSet
				_ = _594_v38
				_594_v38 = _dafny.MultiSetOf(p1)
				var _595_v39 _dafny.Sequence
				_ = _595_v39
				_595_v39 = _dafny.SeqOf(true)
				var _596_v40 _dafny.Map
				_ = _596_v40
				_596_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), _595_v39)
				var _597_v41 _dafny.Array
				_ = _597_v41
				var _nwElement0_16 bool = p1
				_ = _nwElement0_16
				var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(26))
				_ = _nw94
				(_nw94).ArraySet1(_nwElement0_16, 0)
				(_nw94).ArraySet1(p1, 1)
				(_nw94).ArraySet1(p1, 2)
				(_nw94).ArraySet1(!(Companion_Default___.Fm0((Companion_D1_.Create_DC4_(_582_cf19, p1, _589_v35)).Dtor_cf5(), _584_v30, p1, globalState)), 3)
				(_nw94).ArraySet1(p1, 4)
				(_nw94).ArraySet1(!(p1) || (p1), 5)
				(_nw94).ArraySet1(false, 6)
				(_nw94).ArraySet1(p1, 7)
				(_nw94).ArraySet1(false, 8)
				(_nw94).ArraySet1(false, 9)
				(_nw94).ArraySet1((_dafny.IntOfUint32(((_590_v36).Dtor_cf7()).Cardinality())).Cmp(_582_cf19) > 0, 10)
				(_nw94).ArraySet1(p1, 11)
				(_nw94).ArraySet1(!(p1), 12)
				(_nw94).ArraySet1(Companion_Default___.Fm0(p1, _584_v30, Companion_Default___.Fm0(p1, _584_v30, p1, globalState), globalState), 13)
				(_nw94).ArraySet1(p1, 14)
				(_nw94).ArraySet1(_dafny.Companion_Sequence_.Equal(_593_v37, _593_v37), 15)
				(_nw94).ArraySet1(p1, 16)
				(_nw94).ArraySet1(p1, 17)
				(_nw94).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-100))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}((func(_598_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_599_i4 _dafny.Int) _dafny.CodePoint {
						return _598_v30
					}
				})(_584_v30))), _586_v32), 18)
				(_nw94).ArraySet1(p1, 19)
				(_nw94).ArraySet1((_594_v38).Equals(_594_v38), 20)
				(_nw94).ArraySet1(p1, 21)
				(_nw94).ArraySet1(p1, 22)
				(_nw94).ArraySet1(!(p1), 23)
				(_nw94).ArraySet1(p1, 24)
				(_nw94).ArraySet1(_dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
					if (_596_v40).Contains(p1) {
						return (_596_v40).Get(p1).(_dafny.Sequence)
					}
					return _595_v39
				})(), _dafny.Companion_Sequence_.Update(_595_v39, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_595_v39).Cardinality()))).Uint32(), false)), 25)
				_597_v41 = _nw94
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_597_v41), 0))
				_ = _index62
				(_597_v41).ArraySet1(p1, (_index62).Int())
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_597_v41), 0))
				_ = _index63
				(_597_v41).ArraySet1(Companion_Default___.Fm0(p1, _584_v30, p1, globalState), (_index63).Int())
				r2 = Companion_Default___.Fm9(_dafny.IntOfInt64(-77), _582_cf19, (_597_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_597_v41), 0))).Int()).(bool), globalState)
			} else {
				var _600___mcc_h4 _dafny.Set = _source8.Get_().(D3_DC9).Cf16
				_ = _600___mcc_h4
				var _601_cf16 _dafny.Set = _600___mcc_h4
				_ = _601_cf16
				var _602_v42 _dafny.Array
				_ = _602_v42
				var _nw95 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw95
				_602_v42 = _nw95
				_602_v42 = (func() _dafny.Array {
					if p1 {
						return _602_v42
					}
					return _602_v42
				})()
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen(((_this).F22()), 0))
				_ = _index64
				((_this).F22()).ArraySet1(p0, (_index64).Int())
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen(((_this).F22()), 0))
				_ = _index65
				((_this).F22()).ArraySet1(_dafny.IntOfInt64(276), (_index65).Int())
				var _603_v43 _dafny.CodePoint
				_ = _603_v43
				_603_v43 = _dafny.CodePoint('a')
				(globalState).F15 = Companion_Default___.Fm0((func() bool {
					if false {
						return p1
					}
					return p1
				})(), _603_v43, p1, globalState)
				var _604_v44 _dafny.Sequence
				_ = _604_v44
				_604_v44 = _dafny.SeqOf(((_this).F22()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen(((_this).F22()), 0))).Int()).(_dafny.Int))
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen(((_this).F22()), 0))
				_ = _index66
				((_this).F22()).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if p1 {
						return (_604_v44).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_604_v44).Cardinality()))).Uint32()).(_dafny.Int)
					}
					return (_dafny.Zero).Minus(p0)
				})(), ((_this).F22()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen(((_this).F22()), 0))).Int()).(_dafny.Int)), (_index66).Int())
			}
			var _605_v45 _dafny.Array
			_ = _605_v45
			var _nwElement0_17 bool = (p0).Cmp(_dafny.IntOfInt64(667)) <= 0
			_ = _nwElement0_17
			var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(3))
			_ = _nw96
			(_nw96).ArraySet1(_nwElement0_17, 0)
			(_nw96).ArraySet1(!((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-190))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}(func(_606_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			}))).Cardinality())).Cmp(p0) >= 0), 1)
			(_nw96).ArraySet1(p1, 2)
			_605_v45 = _nw96
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_605_v45), 0))
			_ = _index67
			(_605_v45).ArraySet1(true, (_index67).Int())
			var _607_v46 _dafny.Sequence
			_ = _607_v46
			_607_v46 = _dafny.SeqOf(_dafny.IntOfUint32((Companion_Default___.Fm16(!(p1), p0, globalState)).Cardinality()))
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_605_v45), 0))
			_ = _index68
			(_605_v45).ArraySet1((_dafny.Companion_Sequence_.Contains(_607_v46, _dafny.IntOfInt64(543))) == (p1), (_index68).Int())
			var _608_v47 _dafny.Array
			_ = _608_v47
			var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
			_ = _nw97
			_608_v47 = _nw97
			_608_v47 = _608_v47
			var _609_v48 _dafny.Map
			_ = _609_v48
			_609_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), (_605_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_605_v45), 0))).Int()).(bool))
			var _610_v49 D0
			_ = _610_v49
			_610_v49 = Companion_D0_.Create_DC0_((p0).Cmp(p0) <= 0)
			var _pat_let_tv17 = _605_v45
			_ = _pat_let_tv17
			var _pat_let_tv18 = _605_v45
			_ = _pat_let_tv18
			var _rhs63 _dafny.Int = p0
			_ = _rhs63
			var _rhs64 _dafny.Map = (_609_v48).Merge(_609_v48)
			_ = _rhs64
			var _rhs65 D0 = func(_pat_let11_0 D0) D0 {
				return func(_611_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let12_0 bool) D0 {
						return func(_612_dt__update_hcf0_h0 bool) D0 {
							return Companion_D0_.Create_DC0_(_612_dt__update_hcf0_h0)
						}(_pat_let12_0)
					}((_pat_let_tv18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_pat_let_tv17), 0))).Int()).(bool))
				}(_pat_let11_0)
			}(Companion_D0_.Create_DC0_(p1))
			_ = _rhs65
			var _rhs66 _dafny.Int = p0
			_ = _rhs66
			var _rhs67 _dafny.Int = p0
			_ = _rhs67
			r2 = _rhs63
			_609_v48 = _rhs64
			_610_v49 = _rhs65
			r2 = _rhs66
			r2 = _rhs67
			var _613_v50 _dafny.Sequence
			_ = _613_v50
			_613_v50 = _dafny.SeqOf((_605_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_605_v45), 0))).Int()).(bool), p1)
			var _614_v51 _dafny.MultiSet
			_ = _614_v51
			_614_v51 = _dafny.MultiSetOf(_613_v50, _dafny.SeqOf((_605_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_605_v45), 0))).Int()).(bool)))
			r2 = (func() _dafny.Int {
				if ((_614_v51).Update(_613_v50, Companion_Default___.Abs(p0))).IsSubsetOf((_614_v51).Update(_613_v50, Companion_Default___.Abs(p0))) {
					return p0
				}
				return p0
			})()
		}
		var _615_v52 *C0
		_ = _615_v52
		var _nw98 *C0 = New_C0_()
		_ = _nw98
		_nw98.Ctor__(!(p1) || (p1))
		_615_v52 = _nw98
		var _616_v53 _dafny.Sequence
		_ = _616_v53
		_616_v53 = _dafny.UnicodeSeqOfUtf8Bytes("fnymofdaa")
		var _617_v54 _dafny.MultiSet
		_ = _617_v54
		_617_v54 = _dafny.MultiSetOf(p1, p1)
		var _618_v55 D1
		_ = _618_v55
		_618_v55 = Companion_D1_.Create_DC6_(p0, (_617_v54).Cardinality())
		var _pat_let_tv19 = _616_v53
		_ = _pat_let_tv19
		var _pat_let_tv20 = _616_v53
		_ = _pat_let_tv20
		var _pat_let_tv21 = _616_v53
		_ = _pat_let_tv21
		var _pat_let_tv22 = _616_v53
		_ = _pat_let_tv22
		var _pat_let_tv23 = _616_v53
		_ = _pat_let_tv23
		_616_v53 = func(_source9 D1) _dafny.Sequence {
			if _source9.Is_DC4() {
				var _619___mcc_h5 _dafny.Int = _source9.Get_().(D1_DC4).Cf4
				_ = _619___mcc_h5
				var _620___mcc_h6 bool = _source9.Get_().(D1_DC4).Cf5
				_ = _620___mcc_h6
				var _621___mcc_h7 _dafny.Array = _source9.Get_().(D1_DC4).Cf6
				_ = _621___mcc_h7
				var _622_cf6 _dafny.Array = _621___mcc_h7
				_ = _622_cf6
				var _623_cf5 bool = _620___mcc_h6
				_ = _623_cf5
				var _624_cf4 _dafny.Int = _619___mcc_h5
				_ = _624_cf4
				return _dafny.Companion_Sequence_.Concatenate(_pat_let_tv19, _dafny.UnicodeSeqOfUtf8Bytes("wgsdhwirx"))
			} else if _source9.Is_DC5() {
				var _625___mcc_h8 _dafny.Sequence = _source9.Get_().(D1_DC5).Cf7
				_ = _625___mcc_h8
				var _626___mcc_h9 _dafny.Int = _source9.Get_().(D1_DC5).Cf8
				_ = _626___mcc_h9
				var _627___mcc_h10 _dafny.Int = _source9.Get_().(D1_DC5).Cf9
				_ = _627___mcc_h10
				var _628_cf9 _dafny.Int = _627___mcc_h10
				_ = _628_cf9
				var _629_cf8 _dafny.Int = _626___mcc_h9
				_ = _629_cf8
				var _630_cf7 _dafny.Sequence = _625___mcc_h8
				_ = _630_cf7
				return _pat_let_tv20
			} else if _source9.Is_DC6() {
				var _631___mcc_h11 _dafny.Int = _source9.Get_().(D1_DC6).Cf10
				_ = _631___mcc_h11
				var _632___mcc_h12 _dafny.Int = _source9.Get_().(D1_DC6).Cf11
				_ = _632___mcc_h12
				var _633_cf11 _dafny.Int = _632___mcc_h12
				_ = _633_cf11
				var _634_cf10 _dafny.Int = _631___mcc_h11
				_ = _634_cf10
				return _dafny.Companion_Sequence_.Concatenate(_pat_let_tv21, _pat_let_tv22)
			} else {
				var _635___mcc_h13 _dafny.Array = _source9.Get_().(D1_DC3).Cf3
				_ = _635___mcc_h13
				var _636_cf3 _dafny.Array = _635___mcc_h13
				_ = _636_cf3
				return _pat_let_tv23
			}
		}(_618_v55)
		var _637_v56 _dafny.Map
		_ = _637_v56
		_637_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _616_v53)
		var _638_v57 D0
		_ = _638_v57
		_638_v57 = Companion_D0_.Create_DC1_(p1)
		var _639_v58 _dafny.Set
		_ = _639_v58
		_639_v58 = _dafny.SetOf(_638_v57)
		var _640_v59 D1
		_ = _640_v59
		_640_v59 = Companion_D1_.Create_DC5_((func() _dafny.Sequence {
			if (_637_v56).Contains(true) {
				return (_637_v56).Get(true).(_dafny.Sequence)
			}
			return _616_v53
		})(), p0, (_dafny.Zero).Minus((_639_v58).Cardinality()))
		var _641_v60 D3
		_ = _641_v60
		_641_v60 = Companion_D3_.Create_DC10_(p0)
		var _642_v61 _dafny.Map
		_ = _642_v61
		_642_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F22())
		_640_v59 = Companion_D1_.Create_DC5_(_616_v53, (_641_v60).Dtor_cf17(), (p0).Minus((_642_v61).Cardinality()))
		var _643_v62 _dafny.Array
		_ = _643_v62
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_16
		var _nw99 _dafny.Array
		_ = _nw99
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw99 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) _dafny.Int = (func(_644_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_645_i7 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_645_i7, (_dafny.MultiSetOf(_644_p0, _644_p0, _644_p0, _644_p0, _644_p0)).Cardinality())
				}
			})(p0)
			_ = _init16
			var _element0_16 = _init16(_dafny.Zero)
			_ = _element0_16
			_nw99 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
			(_nw99).ArraySet1(_element0_16, 0)
			var _nativeLen0_16 = (_len0_16).Int()
			_ = _nativeLen0_16
			for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
				(_nw99).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
			}
		}
		_643_v62 = _nw99
		for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_643_v62), 0))); ; {
			_guard_loop_1, _ok24 := _iter24()
			if !_ok24 {
				break
			}
			var _646_i6 _dafny.Int
			_646_i6 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_646_i6).Sign() != -1) && ((_646_i6).Cmp(_dafny.ArrayLen((_643_v62), 0)) < 0)) {
				(_643_v62).ArraySet1(Companion_Default___.SafeDivisionInt(_646_i6, (_617_v54).Cardinality()), (_646_i6).Int())
			}
		}
		var _647_v63 _dafny.CodePoint
		_ = _647_v63
		_647_v63 = _dafny.CodePoint('i')
		var _648_v64 _dafny.Map
		_ = _648_v64
		_648_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _647_v63)
		_648_v64 = (_648_v64).Update(p0, _647_v63)
		var _649_v65 _dafny.Set
		_ = _649_v65
		_649_v65 = _dafny.SetOf((_dafny.Zero).Minus(p0))
		r0 = _649_v65
		var _650_v66 _dafny.Map
		_ = _650_v66
		_650_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_615_v52).F23())
		var _651_v67 _dafny.Map
		_ = _651_v67
		_651_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
		var _652_v68 _dafny.Map
		_ = _652_v68
		_652_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() bool {
			if (_650_v66).Contains(p1) {
				return (_650_v66).Get(p1).(bool)
			}
			return (_615_v52).Fm15(false, (_615_v52).F23(), p0, _651_v67, globalState)
		})())
		r1 = (func() bool {
			if (_652_v68).Contains(((func() _dafny.Int {
				if (_651_v67).Contains(p1) {
					return (_651_v67).Get(p1).(_dafny.Int)
				}
				return p0
			})()).Plus(_dafny.IntOfInt64(462))) {
				return (_652_v68).Get(((func() _dafny.Int {
					if (_651_v67).Contains(p1) {
						return (_651_v67).Get(p1).(_dafny.Int)
					}
					return p0
				})()).Plus(_dafny.IntOfInt64(462))).(bool)
			}
			return p1
		})()
		r2 = p0
		var _653_v70 T0
		_ = _653_v70
		var _nw100 *C0 = New_C0_()
		_ = _nw100
		_nw100.Ctor__((_615_v52).F23())
		_653_v70 = _nw100
		var _654_v71 _dafny.Sequence
		_ = _654_v71
		_654_v71 = _dafny.SeqOf(_653_v70)
		var _655_v73 _dafny.MultiSet
		_ = _655_v73
		_655_v73 = _dafny.MultiSetOf(p0)
		var _656_v74 _dafny.Sequence
		_ = _656_v74
		_656_v74 = _dafny.SeqOf(p0, (_655_v73).Cardinality(), p0)
		var _657_v75 _dafny.Sequence
		_ = _657_v75
		_657_v75 = _dafny.SeqOf(_dafny.IntOfUint32((_654_v71).Cardinality()), p0, (func() _dafny.Map {
			var _coll23 = _dafny.NewMapBuilder()
			_ = _coll23
			for _iter25 := _dafny.Iterate((_656_v74).Elements()); ; {
				_compr_23, _ok25 := _iter25()
				if !_ok25 {
					break
				}
				var _658_v72 _dafny.Int
				_658_v72 = interface{}(_compr_23).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_656_v74, _658_v72) {
					_coll23.Add((_658_v72).Times(_dafny.IntOfUint32((_dafny.SeqOf((_615_v52).F23())).Cardinality())), _dafny.IntOfInt64(96))
				}
			}
			return _coll23.ToMap()
		}()).Cardinality(), p0)
		r3 = func() _dafny.Map {
			var _coll24 = _dafny.NewMapBuilder()
			_ = _coll24
			for _iter26 := _dafny.Iterate((_dafny.MultiSetFromSeq(_657_v75)).Elements()); ; {
				_compr_24, _ok26 := _iter26()
				if !_ok26 {
					break
				}
				var _659_v69 _dafny.Int
				_659_v69 = interface{}(_compr_24).(_dafny.Int)
				if (_dafny.MultiSetFromSeq(_657_v75)).Contains(_659_v69) {
					_coll24.Add(Companion_Default___.SafeDivisionInt(_659_v69, p0), (_615_v52).F23())
				}
			}
			return _coll24.ToMap()
		}()
		return r0, r1, r2, r3
	}
}
func (_this *C1) F22() _dafny.Array {
	{
		return _this._f22
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	F20  D0
	_f21 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this.F20 = Companion_D0_.Default()
	_this._f21 = false
	return &_this
}

type CompanionStruct_C2_ struct {
}

var Companion_C2_ = CompanionStruct_C2_{}

func (_this *C2) Equals(other *C2) bool {
	return _this == other
}

func (_this *C2) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C2)
	return ok && _this.Equals(other)
}

func (*C2) String() string {
	return "_module.C2"
}

func Type_C2_() _dafny.TypeDescriptor {
	return type_C2_{}
}

type type_C2_ struct {
}

func (_this type_C2_) Default() interface{} {
	return (*C2)(nil)
}

func (_this type_C2_) String() string {
	return "main.C2"
}
func (_this *C2) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C2{}
var _ T1 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) Ctor__(f20 D0, f21 bool) {
	{
		(_this).F20 = f20
		(_this)._f21 = f21
	}
}
func (_this *C2) Fm8(p0 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qr"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tglyx"), _dafny.UnicodeSeqOfUtf8Bytes("x")))
	}
}
func (_this *C2) M1(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _660_i0 _dafny.Int
		_ = _660_i0
		_660_i0 = _dafny.Zero
		{
			for p2 {
				{
					if (_660_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_660_i0 = (_660_i0).Plus(_dafny.One)
					(globalState).F11 = !(p0)
					var _661_v0 _dafny.Sequence
					_ = _661_v0
					_661_v0 = _dafny.UnicodeSeqOfUtf8Bytes("qletf")
					var _rhs68 bool = p2
					_ = _rhs68
					var _rhs69 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_661_v0, _661_v0), _661_v0)
					_ = _rhs69
					var _lhs31 *GlobalState = globalState
					_ = _lhs31
					var _lhs32 *GlobalState = globalState
					_ = _lhs32
					_lhs31.F15 = _rhs68
					_lhs32.F15 = _rhs69
					var _662_v1 _dafny.MultiSet
					_ = _662_v1
					_662_v1 = _dafny.MultiSetOf(_dafny.IntOfInt64(328))
					var _663_v2 _dafny.Sequence
					_ = _663_v2
					_663_v2 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(456))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg35 _dafny.Int) interface{} {
							return coer35(arg35)
						}
					}(func(_664_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('h')
					})))
					(globalState).F17 = ((_662_v1).Update(p3, Companion_Default___.Abs(_dafny.IntOfUint32((_663_v2).Cardinality())))).IsDisjointFrom(_662_v1)
					var _665_v3 _dafny.Array
					_ = _665_v3
					var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
					_ = _nw101
					_665_v3 = _nw101
					var _666_v4 T0
					_ = _666_v4
					var _nw102 *C1 = New_C1_()
					_ = _nw102
					_nw102.Ctor__(_665_v3)
					_666_v4 = _nw102
					var _667_v5 D4
					_ = _667_v5
					_667_v5 = Companion_D4_.Create_DC12_(_666_v4)
					var _668_v6 _dafny.Sequence
					_ = _668_v6
					_668_v6 = _dafny.SeqOf(_666_v4, _666_v4)
					var _669_v7 _dafny.Map
					_ = _669_v7
					_669_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _665_v3)
					var _670_v8 _dafny.Map
					_ = _670_v8
					_670_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
						if (_669_v7).Contains(p1) {
							return (_669_v7).Get(p1).(_dafny.Array)
						}
						return _665_v3
					})(), p1)
					var _671_v9 _dafny.Array
					_ = _671_v9
					var _nwElement0_18 T0 = (func() T0 {
						if (_this).F21() {
							return _666_v4
						}
						return (_667_v5).Dtor_cf21()
					})()
					_ = _nwElement0_18
					var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(12))
					_ = _nw103
					(_nw103).ArraySet1(_nwElement0_18, 0)
					(_nw103).ArraySet1(_666_v4, 1)
					(_nw103).ArraySet1(_666_v4, 2)
					(_nw103).ArraySet1(_666_v4, 3)
					(_nw103).ArraySet1(_666_v4, 4)
					(_nw103).ArraySet1(_666_v4, 5)
					(_nw103).ArraySet1((_668_v6).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_670_v8).Contains(_665_v3) {
							return (_670_v8).Get(_665_v3).(_dafny.Int)
						}
						return p1
					})(), _dafny.IntOfUint32((_668_v6).Cardinality()))).Uint32()).(T0), 6)
					(_nw103).ArraySet1(_666_v4, 7)
					(_nw103).ArraySet1((func() T0 {
						if (_this).F21() {
							return _666_v4
						}
						return _666_v4
					})(), 8)
					(_nw103).ArraySet1(_666_v4, 9)
					(_nw103).ArraySet1(_666_v4, 10)
					(_nw103).ArraySet1(_666_v4, 11)
					_671_v9 = _nw103
					var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_671_v9), 0))
					_ = _index69
					(_671_v9).ArraySet1((_668_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.IntOfUint32((_668_v6).Cardinality()))).Uint32()).(T0), (_index69).Int())
					var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_671_v9), 0))
					_ = _index70
					(_671_v9).ArraySet1(_666_v4, (_index70).Int())
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _672_v10 _dafny.Int
		_ = _672_v10
		_672_v10 = _dafny.IntOfInt64(-235)
		var _673_v11 _dafny.Sequence
		_ = _673_v11
		_673_v11 = _dafny.UnicodeSeqOfUtf8Bytes("p")
		var _674_v12 _dafny.Map
		_ = _674_v12
		_674_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F21())
		_672_v10 = Companion_Default___.Fm9((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_673_v11, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(278))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg36 _dafny.Int) interface{} {
				return coer36(arg36)
			}
		}(func(_675_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('l')
		})))).Cardinality())), Companion_Default___.SafeModuloInt(p3, _672_v10), (func() bool {
			if (_674_v12).Contains(p3) {
				return (_674_v12).Get(p3).(bool)
			}
			return p0
		})(), globalState)
		if !(!(!(!(p2)))) || ((func() bool {
			if (_this).F21() {
				return (_this).F21()
			}
			return (_this).F21()
		})()) {
			var _676_v13 _dafny.Sequence
			_ = _676_v13
			_676_v13 = _dafny.SeqOf(_672_v10)
			_672_v10 = (p3).Times((_dafny.IntOfUint32((_676_v13).Cardinality())).Minus(p3))
			_672_v10 = Companion_Default___.Fm9(_672_v10, p3, p0, globalState)
			(globalState).F11 = !((_this).F21())
			var _677_v14 _dafny.Array
			_ = _677_v14
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_17
			var _nw104 _dafny.Array
			_ = _nw104
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw104 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) bool = func(_678_i3 _dafny.Int) bool {
					return (_this).F21()
				}
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw104 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw104).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw104).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_677_v14 = _nw104
			var _679_v15 _dafny.Sequence
			_ = _679_v15
			_679_v15 = _dafny.SeqOf(_677_v14)
			_679_v15 = _679_v15
			_672_v10 = (_dafny.Zero).Minus(p3)
		} else {
			var _680_v16 _dafny.Array
			_ = _680_v16
			var _nw105 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw105
			_680_v16 = _nw105
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_680_v16), 0))
			_ = _index71
			(_680_v16).ArraySet1((_dafny.IntOfUint32((_673_v11).Cardinality())).Cmp(_672_v10) > 0, (_index71).Int())
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_680_v16), 0))
			_ = _index72
			(_680_v16).ArraySet1(p2, (_index72).Int())
			_680_v16 = _680_v16
			if p2 {
				var _681_v17 _dafny.Array
				_ = _681_v17
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_18
				var _nw106 _dafny.Array
				_ = _nw106
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw106 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Map = (func(_682_v12 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_683_i4 _dafny.Int) _dafny.Map {
							return _682_v12
						}
					})(_674_v12)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw106 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw106).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw106).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_681_v17 = _nw106
				var _684_v18 _dafny.Sequence
				_ = _684_v18
				_684_v18 = _dafny.SeqOf(!(p2), p0)
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_681_v17), 0))
				_ = _index73
				(_681_v17).ArraySet1(Companion_Default___.Fm17((_684_v18).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_684_v18).Cardinality()))).Uint32()).(bool), (_this).F21(), globalState), (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_681_v17), 0))
				_ = _index74
				(_681_v17).ArraySet1(_674_v12, (_index74).Int())
				var _685_v19 _dafny.Array
				_ = _685_v19
				var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
				_ = _nw107
				_685_v19 = _nw107
				var _686_v20 _dafny.Set
				_ = _686_v20
				_686_v20 = _dafny.SetOf((_this).F21())
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_685_v19), 0))
				_ = _index75
				(_685_v19).ArraySet1(((_686_v20).Union(_686_v20)).Cardinality(), (_index75).Int())
				var _687_v21 _dafny.CodePoint
				_ = _687_v21
				_687_v21 = _dafny.CodePoint('g')
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_685_v19), 0))
				_ = _index76
				(_685_v19).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_673_v11, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_673_v11).Cardinality()))).Uint32(), _687_v21), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pddyfcmi"), _673_v11))).Cardinality()), (_index76).Int())
				var _688_v22 _dafny.MultiSet
				_ = _688_v22
				_688_v22 = _dafny.MultiSetOf(p1, p3)
				(globalState).F17 = (((_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32(((Companion_D1_.Create_DC5_(_673_v11, _672_v10, (_685_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_685_v19), 0))).Int()).(_dafny.Int))).Dtor_cf7()).Cardinality()))), p1)).Union(_688_v22)).Cardinality()).Cmp((_685_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_685_v19), 0))).Int()).(_dafny.Int)) < 0
				(globalState).F15 = (p1).Cmp(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if !((_680_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_680_v16), 0))).Int()).(bool)) {
						return _673_v11
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(84))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg37 _dafny.Int) interface{} {
							return coer37(arg37)
						}
					}((func(_689_v21 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_690_i5 _dafny.Int) _dafny.CodePoint {
							return _689_v21
						}
					})(_687_v21)))
				})()).Cardinality())) != 0
				(_this).F20 = _this.F20
			} else {
				var _691_v23 _dafny.Sequence
				_ = _691_v23
				_691_v23 = _dafny.SeqOf((func() bool {
					if (_674_v12).Contains(p3) {
						return (_674_v12).Get(p3).(bool)
					}
					return (_this).F21()
				})(), p2)
				var _692_v24 D4
				_ = _692_v24
				_692_v24 = Companion_D4_.Create_DC14_((_691_v23).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_691_v23).Cardinality()))).Uint32()).(bool), (_dafny.Zero).Minus(_672_v10), _dafny.SetOf((_dafny.MultiSetOf(p1)).Cardinality(), _dafny.IntOfInt64(-869), (_674_v12).Cardinality()), _dafny.CodePoint('l'))
				var _693_v25 _dafny.Set
				_ = _693_v25
				_693_v25 = _dafny.SetOf(_680_v16)
				var _694_v26 _dafny.Map
				_ = _694_v26
				_694_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_680_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_680_v16), 0))).Int()).(bool), _693_v25)
				var _rhs70 D4 = _692_v24
				_ = _rhs70
				var _rhs71 _dafny.Int = _672_v10
				_ = _rhs71
				var _rhs72 _dafny.Set = ((func() _dafny.Set {
					if p0 {
						return _693_v25
					}
					return (func() _dafny.Set {
						if (_694_v26).Contains((_680_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_680_v16), 0))).Int()).(bool)) {
							return (_694_v26).Get((_680_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_680_v16), 0))).Int()).(bool)).(_dafny.Set)
						}
						return _693_v25
					})()
				})()).Union((func() _dafny.Set {
					if p2 {
						return _693_v25
					}
					return _693_v25
				})())
				_ = _rhs72
				_692_v24 = _rhs70
				_672_v10 = _rhs71
				_693_v25 = _rhs72
				var _695_v27 _dafny.CodePoint
				_ = _695_v27
				_695_v27 = _dafny.CodePoint('w')
				var _696_v28 _dafny.MultiSet
				_ = _696_v28
				_696_v28 = _dafny.MultiSetOf(_695_v27)
				_696_v28 = _696_v28
				(globalState).F17 = p2
				_672_v10 = Companion_Default___.SafeDivisionInt(_672_v10, ((func() _dafny.Map {
					var _coll25 = _dafny.NewMapBuilder()
					_ = _coll25
					for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(447), _dafny.IntOfInt64(-428))); ; {
						_compr_25, _ok27 := _iter27()
						if !_ok27 {
							break
						}
						var _697_v29 _dafny.Int
						_697_v29 = interface{}(_compr_25).(_dafny.Int)
						if ((_dafny.IntOfInt64(447)).Cmp(_697_v29) <= 0) && ((_697_v29).Cmp(_dafny.IntOfInt64(-428)) < 0) {
							_coll25.Add(Companion_Default___.SafeModuloInt(_697_v29, _dafny.IntOfUint32((_dafny.SeqOf((_680_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_680_v16), 0))).Int()).(bool))).Cardinality())), _691_v23)
						}
					}
					return _coll25.ToMap()
				}()).Cardinality()).Times(p1))
				var _698_v30 _dafny.Set
				_ = _698_v30
				_698_v30 = _dafny.SetOf(_dafny.IntOfInt64(-433))
				var _699_v31 _dafny.Set
				_ = _699_v31
				_699_v31 = _dafny.SetOf(p3, (_698_v30).Cardinality(), p3, Companion_Default___.Fm9(_672_v10, p1, p0, globalState))
				(globalState).F11 = (_699_v31).IsSubsetOf(_dafny.SetOf(Companion_Default___.Fm9(_672_v10, p1, p0, globalState)))
			}
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_680_v16), 0))
			_ = _index77
			(_680_v16).ArraySet1((p3).Cmp(p1) != 0, (_index77).Int())
			var _700_v32 *C0
			_ = _700_v32
			var _nw108 *C0 = New_C0_()
			_ = _nw108
			_nw108.Ctor__((_680_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_680_v16), 0))).Int()).(bool))
			_700_v32 = _nw108
		}
		if p2 {
			var _701_v33 _dafny.Array
			_ = _701_v33
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_19
			var _nw109 _dafny.Array
			_ = _nw109
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw109 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) _dafny.Int = (func(_702_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_703_i6 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_703_i6, _702_p3)
					}
				})(p3)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw109 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw109).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw109).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_701_v33 = _nw109
			var _704_v34 _dafny.Sequence
			_ = _704_v34
			_704_v34 = _dafny.SeqOf(p1, _672_v10)
			var _705_v35 _dafny.Map
			_ = _705_v35
			_705_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F21())
			var _706_v36 _dafny.Map
			_ = _706_v36
			_706_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v34, _705_v35)
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_701_v33), 0))
			_ = _index78
			(_701_v33).ArraySet1((_706_v36).Cardinality(), (_index78).Int())
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_701_v33), 0))
			_ = _index79
			(_701_v33).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p3, (_dafny.Zero).Minus(_672_v10))), (func() _dafny.Int {
				if !((_this).F21()) {
					return _672_v10
				}
				return _672_v10
			})()), (_index79).Int())
			var _707_v37 _dafny.Array
			_ = _707_v37
			var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(10))
			_ = _nw110
			_707_v37 = _nw110
			var _708_v38 _dafny.MultiSet
			_ = _708_v38
			_708_v38 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("yyioc"))
			var _709_v39 _dafny.Set
			_ = _709_v39
			_709_v39 = _dafny.SetOf(p0)
			var _710_v40 D4
			_ = _710_v40
			_710_v40 = Companion_D4_.Create_DC13_(_672_v10)
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_701_v33), 0))
			_ = _index80
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_701_v33), 0))
			_ = _index81
			var _rhs73 _dafny.Int = p1
			_ = _rhs73
			var _rhs74 _dafny.Array = _707_v37
			_ = _rhs74
			var _rhs75 _dafny.Int = ((func() _dafny.Int {
				if (_708_v38).Contains(_673_v11) {
					return (_708_v38).Multiplicity(_673_v11)
				}
				return (_709_v39).Cardinality()
			})()).Minus((_710_v40).Dtor_cf22())
			_ = _rhs75
			var _rhs76 _dafny.Sequence = _673_v11
			_ = _rhs76
			var _lhs33 _dafny.Array = _701_v33
			_ = _lhs33
			var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_701_v33), 0))
			_ = _lhs34
			var _lhs35 _dafny.Array = _701_v33
			_ = _lhs35
			var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_701_v33), 0))
			_ = _lhs36
			(_lhs33).ArraySet1(_rhs73, (_lhs34).Int())
			_707_v37 = _rhs74
			(_lhs35).ArraySet1(_rhs75, (_lhs36).Int())
			_673_v11 = _rhs76
			var _711_v41 _dafny.Array
			_ = _711_v41
			var _len0_20 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_20
			var _nw111 _dafny.Array
			_ = _nw111
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw111 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) _dafny.Sequence = (func(_712_v11 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_713_i7 _dafny.Int) _dafny.Sequence {
						return _712_v11
					}
				})(_673_v11)
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw111 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw111).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw111).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_711_v41 = _nw111
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_711_v41), 0))
			_ = _index82
			(_711_v41).ArraySet1(_673_v11, (_index82).Int())
			var _714_v42 _dafny.Map
			_ = _714_v42
			_714_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_672_v10, p1)
			var _715_v43 _dafny.Set
			_ = _715_v43
			_715_v43 = _dafny.SetOf((func() _dafny.Int {
				if (_714_v42).Contains(p1) {
					return (_714_v42).Get(p1).(_dafny.Int)
				}
				return _dafny.IntOfInt64(615)
			})())
			var _716_v44 _dafny.CodePoint
			_ = _716_v44
			_716_v44 = _dafny.CodePoint('v')
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_711_v41), 0))
			_ = _index83
			var _rhs77 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_673_v11, _673_v11), (Companion_Default___.SafeIndex(Companion_Default___.Fm9(p3, p3, false, globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_673_v11, _673_v11)).Cardinality()))).Uint32(), _716_v44)
			_ = _rhs77
			var _rhs78 _dafny.Set = _dafny.SetOf((p1).Plus((_701_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_701_v33), 0))).Int()).(_dafny.Int)), p1, _dafny.IntOfUint32((_704_v34).Cardinality()))
			_ = _rhs78
			var _lhs37 _dafny.Array = _711_v41
			_ = _lhs37
			var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_711_v41), 0))
			_ = _lhs38
			(_lhs37).ArraySet1(_rhs77, (_lhs38).Int())
			_715_v43 = _rhs78
			var _717_v45 D1
			_ = _717_v45
			_717_v45 = Companion_D1_.Create_DC6_(p3, (_701_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_701_v33), 0))).Int()).(_dafny.Int))
			_717_v45 = Companion_D1_.Create_DC6_(p3, p3)
			var _718_v46 *C0
			_ = _718_v46
			var _nw112 *C0 = New_C0_()
			_ = _nw112
			_nw112.Ctor__((_this).F21())
			_718_v46 = _nw112
		} else {
			_672_v10 = _dafny.IntOfInt64(458)
			_672_v10 = (_672_v10).Minus(p3)
			var _719_v47 _dafny.Array
			_ = _719_v47
			var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw113
			_719_v47 = _nw113
			var _720_v48 *C1
			_ = _720_v48
			var _nw114 *C1 = New_C1_()
			_ = _nw114
			_nw114.Ctor__(_719_v47)
			_720_v48 = _nw114
			var _721_v49 _dafny.Map
			_ = _721_v49
			_721_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_720_v48, p0)
			var _722_v50 _dafny.Map
			_ = _722_v50
			_722_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_721_v49).Contains(_720_v48) {
					return (_721_v49).Get(_720_v48).(bool)
				}
				return (_this).F21()
			})(), (_this).F21())
			_722_v50 = (_722_v50).Update((func() bool {
				if (_674_v12).Contains(_672_v10) {
					return (_674_v12).Get(_672_v10).(bool)
				}
				return p0
			})(), p2)
			_672_v10 = p1
			var _723_v51 _dafny.Sequence
			_ = _723_v51
			_723_v51 = _dafny.SeqOf((_this).F21(), !_dafny.Companion_Sequence_.Equal(_673_v11, _673_v11), p0)
			var _rhs79 bool = !((_723_v51).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_723_v51).Cardinality()))).Uint32()).(bool))
			_ = _rhs79
			var _rhs80 bool = p0
			_ = _rhs80
			var _rhs81 bool = p2
			_ = _rhs81
			var _lhs39 *GlobalState = globalState
			_ = _lhs39
			var _lhs40 *GlobalState = globalState
			_ = _lhs40
			var _lhs41 *GlobalState = globalState
			_ = _lhs41
			_lhs39.F15 = _rhs79
			_lhs40.F11 = _rhs80
			_lhs41.F15 = _rhs81
		}
		if !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-184))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg38 _dafny.Int) interface{} {
				return coer38(arg38)
			}
		}(func(_724_i8 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})), _673_v11) {
			var _725_v52 _dafny.Set
			_ = _725_v52
			_725_v52 = _dafny.SetOf(p3)
			_672_v10 = (func() _dafny.Int {
				if (_725_v52).Equals(_dafny.SetOf((_674_v12).Cardinality())) {
					return p1
				}
				return (_672_v10).Plus(p3)
			})()
			_674_v12 = (_674_v12).Update(_dafny.IntOfInt64(707), (Companion_Default___.Fm18((_this).F21(), _672_v10, p3, globalState)).IsSubsetOf(_725_v52))
			_672_v10 = _672_v10
			_672_v10 = _672_v10
			var _726_v53 _dafny.Map
			_ = _726_v53
			_726_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), p1)
			var _727_v54 _dafny.Sequence
			_ = _727_v54
			_727_v54 = _dafny.SeqOf(p2, (_this).F21())
			var _728_v55 _dafny.Sequence
			_ = _728_v55
			_728_v55 = _dafny.SeqOf(p1)
			var _729_v56 _dafny.Array
			_ = _729_v56
			var _nwElement0_19 _dafny.Int = _dafny.IntOfUint32((_727_v54).Cardinality())
			_ = _nwElement0_19
			var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(20))
			_ = _nw115
			(_nw115).ArraySet1(_nwElement0_19, 0)
			(_nw115).ArraySet1(p3, 1)
			(_nw115).ArraySet1(_672_v10, 2)
			(_nw115).ArraySet1(p1, 3)
			(_nw115).ArraySet1(p3, 4)
			(_nw115).ArraySet1((_728_v55).Select((Companion_Default___.SafeIndex((_dafny.SetOf((_this).F21(), true, false, p2)).Cardinality(), _dafny.IntOfUint32((_728_v55).Cardinality()))).Uint32()).(_dafny.Int), 5)
			(_nw115).ArraySet1((_dafny.Zero).Minus(p1), 6)
			(_nw115).ArraySet1(_672_v10, 7)
			(_nw115).ArraySet1(_672_v10, 8)
			(_nw115).ArraySet1(p1, 9)
			(_nw115).ArraySet1((_728_v55).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_728_v55).Cardinality()))).Uint32()).(_dafny.Int), 10)
			(_nw115).ArraySet1((_674_v12).Cardinality(), 11)
			(_nw115).ArraySet1(p1, 12)
			(_nw115).ArraySet1(p3, 13)
			(_nw115).ArraySet1(_dafny.IntOfUint32((_728_v55).Cardinality()), 14)
			(_nw115).ArraySet1(p1, 15)
			(_nw115).ArraySet1((_dafny.Zero).Minus(p3), 16)
			(_nw115).ArraySet1(p3, 17)
			(_nw115).ArraySet1(_672_v10, 18)
			(_nw115).ArraySet1(_672_v10, 19)
			_729_v56 = _nw115
			var _730_v57 _dafny.MultiSet
			_ = _730_v57
			_730_v57 = _dafny.MultiSetOf(false, p2, p2, p2)
			var _731_v58 D4
			_ = _731_v58
			_731_v58 = Companion_D4_.Create_DC14_((_this).F21(), _672_v10, _725_v52, (_673_v11).Select((Companion_Default___.SafeIndex((_730_v57).Cardinality(), _dafny.IntOfUint32((_673_v11).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _732_v59 _dafny.Map
			_ = _732_v59
			_732_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC11_(_726_v53, _dafny.IntOfInt64(919), _729_v56), _731_v58)
			_672_v10 = (_732_v59).Cardinality()
		} else {
			var _733_v60 _dafny.Array
			_ = _733_v60
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_21
			var _nw116 _dafny.Array
			_ = _nw116
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw116 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) bool = (func(_734_v10 _dafny.Int, _735_p3 _dafny.Int, _736_p2 bool) func(_dafny.Int) bool {
					return func(_737_i9 _dafny.Int) bool {
						return (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_734_v10, _dafny.CodePoint('s'))).Cardinality(), _735_p3, _dafny.IntOfUint32((_dafny.SeqOf(_736_p2)).Cardinality()), _735_p3)).IsDisjointFrom(_dafny.SetOf(_735_p3))
					}
				})(_672_v10, p3, p2)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw116 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw116).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw116).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_733_v60 = _nw116
			_733_v60 = _733_v60
			(globalState).F15 = (p0) && (p2)
			if (_672_v10).Cmp(_672_v10) < 0 {
				var _738_v61 _dafny.CodePoint
				_ = _738_v61
				_738_v61 = _dafny.CodePoint('x')
				(globalState).F17 = !_dafny.Companion_Sequence_.Contains(_673_v11, _738_v61)
				var _739_v62 _dafny.Array
				_ = _739_v62
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_22
				var _nw117 _dafny.Array
				_ = _nw117
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw117 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) D4 = (func(_740_p3 _dafny.Int) func(_dafny.Int) D4 {
						return func(_741_i10 _dafny.Int) D4 {
							return Companion_D4_.Create_DC13_(_740_p3)
						}
					})(p3)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw117 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw117).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw117).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_739_v62 = _nw117
				var _742_v63 _dafny.Map
				_ = _742_v63
				_742_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _739_v62)
				var _743_v64 _dafny.Array
				_ = _743_v64
				var _nwElement0_20 _dafny.Array = _739_v62
				_ = _nwElement0_20
				var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(21))
				_ = _nw118
				(_nw118).ArraySet1(_nwElement0_20, 0)
				(_nw118).ArraySet1(_739_v62, 1)
				(_nw118).ArraySet1(_739_v62, 2)
				(_nw118).ArraySet1(_739_v62, 3)
				(_nw118).ArraySet1(_739_v62, 4)
				(_nw118).ArraySet1(_739_v62, 5)
				(_nw118).ArraySet1((func() _dafny.Array {
					if p0 {
						return _739_v62
					}
					return _739_v62
				})(), 6)
				(_nw118).ArraySet1((func() _dafny.Array {
					if (_this).F21() {
						return _739_v62
					}
					return _739_v62
				})(), 7)
				(_nw118).ArraySet1(_739_v62, 8)
				(_nw118).ArraySet1(_739_v62, 9)
				(_nw118).ArraySet1(_739_v62, 10)
				(_nw118).ArraySet1(_739_v62, 11)
				(_nw118).ArraySet1(_739_v62, 12)
				(_nw118).ArraySet1(_739_v62, 13)
				(_nw118).ArraySet1(_739_v62, 14)
				(_nw118).ArraySet1(_739_v62, 15)
				(_nw118).ArraySet1(_739_v62, 16)
				(_nw118).ArraySet1(_739_v62, 17)
				(_nw118).ArraySet1(_739_v62, 18)
				(_nw118).ArraySet1(_739_v62, 19)
				(_nw118).ArraySet1((func() _dafny.Array {
					if (_742_v63).Contains(p1) {
						return (_742_v63).Get(p1).(_dafny.Array)
					}
					return _739_v62
				})(), 20)
				_743_v64 = _nw118
				_743_v64 = (func() _dafny.Array {
					if p0 {
						return _743_v64
					}
					return _743_v64
				})()
				var _744_v65 T0
				_ = _744_v65
				var _nw119 *C0 = New_C0_()
				_ = _nw119
				_nw119.Ctor__(p2)
				_744_v65 = _nw119
				var _745_v66 _dafny.Map
				_ = _745_v66
				_745_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _744_v65)
				var _746_v67 D4
				_ = _746_v67
				_746_v67 = Companion_D4_.Create_DC12_((func() T0 {
					if (_745_v66).Contains((_this).F21()) {
						return (_745_v66).Get((_this).F21()).(T0)
					}
					return _744_v65
				})())
				var _pat_let_tv24 = _744_v65
				_ = _pat_let_tv24
				_746_v67 = (func() D4 {
					if p2 {
						return _746_v67
					}
					return func(_pat_let13_0 D4) D4 {
						return func(_747_dt__update__tmp_h0 D4) D4 {
							return func(_pat_let14_0 T0) D4 {
								return func(_748_dt__update_hcf21_h0 T0) D4 {
									return Companion_D4_.Create_DC12_(_748_dt__update_hcf21_h0)
								}(_pat_let14_0)
							}(_pat_let_tv24)
						}(_pat_let13_0)
					}(_746_v67)
				})()
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_733_v60), 0))
				_ = _index84
				(_733_v60).ArraySet1(p0, (_index84).Int())
				var _749_v68 _dafny.Array
				_ = _749_v68
				var _nwElement0_21 _dafny.Int = _672_v10
				_ = _nwElement0_21
				var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(4))
				_ = _nw120
				(_nw120).ArraySet1(_nwElement0_21, 0)
				(_nw120).ArraySet1(p1, 1)
				(_nw120).ArraySet1(_dafny.IntOfInt64(9), 2)
				(_nw120).ArraySet1(p3, 3)
				_749_v68 = _nw120
				var _750_v69 _dafny.Map
				_ = _750_v69
				_750_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_733_v60, _749_v68)
				var _751_v70 _dafny.Sequence
				_ = _751_v70
				_751_v70 = _dafny.SeqOf(p1)
				var _752_v71 _dafny.Set
				_ = _752_v71
				_752_v71 = _dafny.SetOf(true)
				var _753_v72 _dafny.Array
				_ = _753_v72
				var _nwElement0_22 _dafny.Int = _dafny.IntOfInt64(505)
				_ = _nwElement0_22
				var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(24))
				_ = _nw121
				(_nw121).ArraySet1(_nwElement0_22, 0)
				(_nw121).ArraySet1(p1, 1)
				(_nw121).ArraySet1(p3, 2)
				(_nw121).ArraySet1((_750_v69).Cardinality(), 3)
				(_nw121).ArraySet1(_dafny.IntOfInt64(499), 4)
				(_nw121).ArraySet1(_dafny.IntOfUint32((_751_v70).Cardinality()), 5)
				(_nw121).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(634))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_754_v61 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_755_i11 _dafny.Int) _dafny.CodePoint {
						return _754_v61
					}
				})(_738_v61)))).Cardinality()), 6)
				(_nw121).ArraySet1(_672_v10, 7)
				(_nw121).ArraySet1(_672_v10, 8)
				(_nw121).ArraySet1(_672_v10, 9)
				(_nw121).ArraySet1((_752_v71).Cardinality(), 10)
				(_nw121).ArraySet1(p1, 11)
				(_nw121).ArraySet1(p1, 12)
				(_nw121).ArraySet1(p1, 13)
				(_nw121).ArraySet1(_dafny.IntOfInt64(235), 14)
				(_nw121).ArraySet1(Companion_Default___.Fm9(_dafny.IntOfUint32((_673_v11).Cardinality()), p3, true, globalState), 15)
				(_nw121).ArraySet1(_672_v10, 16)
				(_nw121).ArraySet1(p3, 17)
				(_nw121).ArraySet1(_dafny.IntOfInt64(181), 18)
				(_nw121).ArraySet1(_dafny.IntOfInt64(772), 19)
				(_nw121).ArraySet1(_dafny.IntOfUint32((_673_v11).Cardinality()), 20)
				(_nw121).ArraySet1(_672_v10, 21)
				(_nw121).ArraySet1(p3, 22)
				(_nw121).ArraySet1(_dafny.IntOfInt64(-85), 23)
				_753_v72 = _nw121
				var _756_v73 _dafny.Array
				_ = _756_v73
				var _nwElement0_23 _dafny.Array = _753_v72
				_ = _nwElement0_23
				var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(3))
				_ = _nw122
				(_nw122).ArraySet1(_nwElement0_23, 0)
				(_nw122).ArraySet1(_749_v68, 1)
				(_nw122).ArraySet1(_753_v72, 2)
				_756_v73 = _nw122
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_733_v60), 0))
				_ = _index85
				var _rhs82 _dafny.Array = _756_v73
				_ = _rhs82
				var _rhs83 bool = p2
				_ = _rhs83
				var _lhs42 *GlobalState = globalState
				_ = _lhs42
				var _lhs43 _dafny.Array = _733_v60
				_ = _lhs43
				var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_733_v60), 0))
				_ = _lhs44
				_lhs42.F2 = _rhs82
				(_lhs43).ArraySet1(_rhs83, (_lhs44).Int())
				var _757_v74 _dafny.Map
				_ = _757_v74
				_757_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), p3)
				_757_v74 = (_757_v74).Update(p2, p1)
			} else {
				_672_v10 = (_dafny.Zero).Minus(p3)
				(globalState).F11 = !_dafny.Companion_Sequence_.Equal(_673_v11, _673_v11)
				(globalState).F17 = !((_this).F21())
				var _758_v75 _dafny.CodePoint
				_ = _758_v75
				_758_v75 = _dafny.CodePoint('b')
				var _759_v76 _dafny.Array
				_ = _759_v76
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_23
				var _nw123 _dafny.Array
				_ = _nw123
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw123 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Int = (func(_760_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_761_i12 _dafny.Int) _dafny.Int {
							return (_761_i12).Times(_760_p1)
						}
					})(p1)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw123 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw123).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw123).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_759_v76 = _nw123
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_759_v76), 0))
				_ = _index86
				(_759_v76).ArraySet1((func() _dafny.Int {
					if p0 {
						return p1
					}
					return _dafny.IntOfInt64(185)
				})(), (_index86).Int())
				var _762_v77 _dafny.Set
				_ = _762_v77
				_762_v77 = _dafny.SetOf(p2)
				var _763_v78 _dafny.Sequence
				_ = _763_v78
				_763_v78 = _dafny.SeqOf(p0, p0, p2, p2, p0)
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_759_v76), 0))
				_ = _index87
				var _rhs84 bool = false
				_ = _rhs84
				var _rhs85 _dafny.CodePoint = _dafny.CodePoint('q')
				_ = _rhs85
				var _rhs86 _dafny.Int = _672_v10
				_ = _rhs86
				var _rhs87 _dafny.Int = Companion_Default___.Fm9((Companion_Default___.Fm9(_dafny.IntOfUint32((_673_v11).Cardinality()), p1, p2, globalState)).Plus((_762_v77).Cardinality()), Companion_Default___.SafeDivisionInt(_672_v10, _dafny.IntOfInt64(636)), (p3).Cmp(_dafny.IntOfUint32((_763_v78).Cardinality())) != 0, globalState)
				_ = _rhs87
				var _rhs88 _dafny.Int = _dafny.IntOfInt64(958)
				_ = _rhs88
				var _lhs45 *GlobalState = globalState
				_ = _lhs45
				var _lhs46 _dafny.Array = _759_v76
				_ = _lhs46
				var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_759_v76), 0))
				_ = _lhs47
				_lhs45.F17 = _rhs84
				_758_v75 = _rhs85
				_672_v10 = _rhs86
				(_lhs46).ArraySet1(_rhs87, (_lhs47).Int())
				_672_v10 = _rhs88
				(globalState).F15 = p0
			}
			var _764_v79 _dafny.Array
			_ = _764_v79
			var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw124
			_764_v79 = _nw124
			var _765_v80 *C1
			_ = _765_v80
			var _nw125 *C1 = New_C1_()
			_ = _nw125
			_nw125.Ctor__(_764_v79)
			_765_v80 = _nw125
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_733_v60), 0))
			_ = _index88
			(_733_v60).ArraySet1(!((_this).F21()), (_index88).Int())
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_733_v60), 0))
			_ = _index89
			(_733_v60).ArraySet1(!((_this).F21()), (_index89).Int())
		}
		var _766_v81 _dafny.MultiSet
		_ = _766_v81
		_766_v81 = _dafny.MultiSetOf(true)
		var _hi6 _dafny.Int = (_766_v81).Cardinality()
		_ = _hi6
		for _767_i13 := _dafny.IntOfInt64(-878); _767_i13.Cmp(_hi6) < 0; _767_i13 = _767_i13.Plus(_dafny.One) {
			var _768_v82 _dafny.Map
			_ = _768_v82
			_768_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F21())).Cardinality())
			var _769_v83 _dafny.Set
			_ = _769_v83
			_769_v83 = _dafny.SetOf(p3)
			var _770_v84 _dafny.CodePoint
			_ = _770_v84
			_770_v84 = _dafny.CodePoint('p')
			var _771_v85 D4
			_ = _771_v85
			_771_v85 = Companion_D4_.Create_DC14_((_this).F21(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(477))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg40 _dafny.Int) interface{} {
					return coer40(arg40)
				}
			}(func(_772_i14 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			}))).Cardinality()), _767_i13, _672_v10)).Cardinality()), _769_v83, _770_v84)
			var _773_v86 _dafny.Map
			_ = _773_v86
			_773_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_768_v82).Contains(p2) {
					return (_768_v82).Get(p2).(_dafny.Int)
				}
				return (_771_v85).Dtor_cf24()
			})(), _dafny.IntOfInt64(264))
			_672_v10 = (_672_v10).Minus((((_773_v86).Update(_dafny.IntOfInt64(603), p3)).Merge(_773_v86)).Cardinality())
			_674_v12 = (_674_v12).Update(p1, !(_773_v86).Contains(_672_v10))
			var _774_v87 _dafny.Set
			_ = _774_v87
			_774_v87 = _dafny.SetOf(p0)
			var _775_v88 _dafny.Sequence
			_ = _775_v88
			_775_v88 = _dafny.SeqOf(_774_v87)
			_672_v10 = (p1).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_775_v88, _775_v88)).Cardinality()))
			var _776_v89 _dafny.Array
			_ = _776_v89
			var _nw126 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
			_ = _nw126
			_776_v89 = _nw126
			var _777_v90 _dafny.Map
			_ = _777_v90
			_777_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F21())
			var _778_v91 _dafny.Map
			_ = _778_v91
			_778_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_767_i13, p2), false)
			var _779_v93 _dafny.Sequence
			_ = _779_v93
			_779_v93 = _dafny.SeqOf((_dafny.SetOf(_dafny.CodePoint('k'))).Cardinality())
			var _780_v94 _dafny.Array
			_ = _780_v94
			var _nwElement0_24 bool = p0
			_ = _nwElement0_24
			var _nw127 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(29))
			_ = _nw127
			(_nw127).ArraySet1(_nwElement0_24, 0)
			(_nw127).ArraySet1(p2, 1)
			(_nw127).ArraySet1(p0, 2)
			(_nw127).ArraySet1((_this).F21(), 3)
			(_nw127).ArraySet1(p0, 4)
			(_nw127).ArraySet1(p0, 5)
			(_nw127).ArraySet1(!((_this).F21()), 6)
			(_nw127).ArraySet1(false, 7)
			(_nw127).ArraySet1((_this).F21(), 8)
			(_nw127).ArraySet1(p0, 9)
			(_nw127).ArraySet1(false, 10)
			(_nw127).ArraySet1(!(p0), 11)
			(_nw127).ArraySet1((_this).F21(), 12)
			(_nw127).ArraySet1(p2, 13)
			(_nw127).ArraySet1(p0, 14)
			(_nw127).ArraySet1(p0, 15)
			(_nw127).ArraySet1(Companion_Default___.Fm0(p0, _770_v84, (_this).F21(), globalState), 16)
			(_nw127).ArraySet1(Companion_Default___.Fm0((_this).F21(), _770_v84, (_this).F21(), globalState), 17)
			(_nw127).ArraySet1(p2, 18)
			(_nw127).ArraySet1(p2, 19)
			(_nw127).ArraySet1((_this).F21(), 20)
			(_nw127).ArraySet1(!((func() bool {
				if (_777_v90).Contains((_this).F21()) {
					return (_777_v90).Get((_this).F21()).(bool)
				}
				return (func() bool {
					if (_778_v91).Contains(func() _dafny.Map {
						var _coll26 = _dafny.NewMapBuilder()
						_ = _coll26
						for _iter28 := _dafny.Iterate((_779_v93).Elements()); ; {
							_compr_26, _ok28 := _iter28()
							if !_ok28 {
								break
							}
							var _781_v92 _dafny.Int
							_781_v92 = interface{}(_compr_26).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_779_v93, _781_v92) {
								_coll26.Add(Companion_Default___.SafeModuloInt(_781_v92, _767_i13), (_this).F21())
							}
						}
						return _coll26.ToMap()
					}()) {
						return (_778_v91).Get(func() _dafny.Map {
							var _coll27 = _dafny.NewMapBuilder()
							_ = _coll27
							for _iter29 := _dafny.Iterate((_779_v93).Elements()); ; {
								_compr_27, _ok29 := _iter29()
								if !_ok29 {
									break
								}
								var _782_v92 _dafny.Int
								_782_v92 = interface{}(_compr_27).(_dafny.Int)
								if _dafny.Companion_Sequence_.Contains(_779_v93, _782_v92) {
									_coll27.Add(Companion_Default___.SafeModuloInt(_782_v92, _767_i13), (_this).F21())
								}
							}
							return _coll27.ToMap()
						}()).(bool)
					}
					return p0
				})()
			})()), 21)
			(_nw127).ArraySet1(p2, 22)
			(_nw127).ArraySet1(true, 23)
			(_nw127).ArraySet1(p2, 24)
			(_nw127).ArraySet1(p0, 25)
			(_nw127).ArraySet1(p0, 26)
			(_nw127).ArraySet1((_this).F21(), 27)
			(_nw127).ArraySet1((_this).F21(), 28)
			_780_v94 = _nw127
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_776_v89), 0))
			_ = _index90
			(_776_v89).ArraySet1(_780_v94, (_index90).Int())
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_776_v89), 0))
			_ = _index91
			(_776_v89).ArraySet1((Companion_D1_.Create_DC4_(_672_v10, (_this).F21(), _780_v94)).Dtor_cf6(), (_index91).Int())
		}
		r0 = p2
		var _783_v95 _dafny.Map
		_ = _783_v95
		_783_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
		var _784_v96 _dafny.CodePoint
		_ = _784_v96
		_784_v96 = _dafny.CodePoint('h')
		var _785_v97 D3
		_ = _785_v97
		_785_v97 = Companion_D3_.Create_DC10_(_672_v10)
		var _786_v98 _dafny.Map
		_ = _786_v98
		_786_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), Companion_Default___.Fm11(_672_v10, (func() bool {
			if (_783_v95).Contains(Companion_Default___.Fm0(p0, _784_v96, p2, globalState)) {
				return (_783_v95).Get(Companion_Default___.Fm0(p0, _784_v96, p2, globalState)).(bool)
			}
			return (_this).F21()
		})(), p2, (_785_v97).Dtor_cf17(), globalState))
		r1 = !(((_dafny.Zero).Minus(p3)).Cmp(((_786_v98).Cardinality()).Minus(p3)) > 0)
		return r0, r1
	}
}
func (_this *C2) M5(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _787_v0 D4
		_ = _787_v0
		_787_v0 = Companion_D4_.Create_DC13_(_dafny.IntOfUint32((_dafny.SeqOf((_this).F21())).Cardinality()))
		var _788_v1 _dafny.Int
		_ = _788_v1
		_788_v1 = _dafny.IntOfInt64(992)
		var _789_v2 _dafny.CodePoint
		_ = _789_v2
		_789_v2 = _dafny.CodePoint('e')
		var _790_v3 D4
		_ = _790_v3
		_790_v3 = Companion_D4_.Create_DC14_(false, (_dafny.Zero).Minus(_788_v1), _dafny.SetOf(_788_v1), _789_v2)
		var _791_v4 _dafny.Sequence
		_ = _791_v4
		_791_v4 = _dafny.SeqOf(_788_v1)
		var _792_v5 D4
		_ = _792_v5
		_792_v5 = Companion_D4_.Create_DC14_(!((_this).F21()), _788_v1, _dafny.SetOf(_788_v1, _788_v1, (_790_v3).Dtor_cf24(), _dafny.IntOfUint32((Companion_Default___.Fm19(_791_v4, true, globalState)).Cardinality())), _789_v2)
		var _pat_let_tv25 = _787_v0
		_ = _pat_let_tv25
		var _pat_let_tv26 = _788_v1
		_ = _pat_let_tv26
		var _pat_let_tv27 = _787_v0
		_ = _pat_let_tv27
		var _pat_let_tv28 = _787_v0
		_ = _pat_let_tv28
		_787_v0 = func(_source10 D4) D4 {
			if _source10.Is_DC13() {
				var _793___mcc_h0 _dafny.Int = _source10.Get_().(D4_DC13).Cf22
				_ = _793___mcc_h0
				var _794_cf22 _dafny.Int = _793___mcc_h0
				_ = _794_cf22
				var _795_dt__update__tmp_h0 D4 = _pat_let_tv25
				_ = _795_dt__update__tmp_h0
				var _796_dt__update_hcf22_h0 _dafny.Int = _pat_let_tv26
				_ = _796_dt__update_hcf22_h0
				return Companion_D4_.Create_DC13_(_796_dt__update_hcf22_h0)
			} else if _source10.Is_DC14() {
				var _797___mcc_h1 bool = _source10.Get_().(D4_DC14).Cf23
				_ = _797___mcc_h1
				var _798___mcc_h2 _dafny.Int = _source10.Get_().(D4_DC14).Cf24
				_ = _798___mcc_h2
				var _799___mcc_h3 _dafny.Set = _source10.Get_().(D4_DC14).Cf25
				_ = _799___mcc_h3
				var _800___mcc_h4 _dafny.CodePoint = _source10.Get_().(D4_DC14).Cf26
				_ = _800___mcc_h4
				var _801_cf26 _dafny.CodePoint = _800___mcc_h4
				_ = _801_cf26
				var _802_cf25 _dafny.Set = _799___mcc_h3
				_ = _802_cf25
				var _803_cf24 _dafny.Int = _798___mcc_h2
				_ = _803_cf24
				var _804_cf23 bool = _797___mcc_h1
				_ = _804_cf23
				return Companion_D4_.Create_DC13_(_803_cf24)
			} else if _source10.Is_DC12() {
				var _805___mcc_h5 T0 = _source10.Get_().(D4_DC12).Cf21
				_ = _805___mcc_h5
				var _806_cf21 T0 = _805___mcc_h5
				_ = _806_cf21
				return _pat_let_tv27
			} else {
				var _807___mcc_h6 D4 = _source10.Get_().(D4_DC15).Cf27
				_ = _807___mcc_h6
				var _808_cf27 D4 = _807___mcc_h6
				_ = _808_cf27
				return _pat_let_tv28
			}
		}(_790_v3)
		if (_this).F21() {
			var _809_v6 _dafny.Set
			_ = _809_v6
			_809_v6 = _dafny.SetOf(_dafny.IntOfUint32((p1).Cardinality()), _788_v1, _788_v1)
			var _810_v7 _dafny.Map
			_ = _810_v7
			_810_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !((_809_v6).IsProperSubsetOf(_809_v6)))
			_788_v1 = (_810_v7).Cardinality()
			var _811_v9 _dafny.Array
			_ = _811_v9
			var _nw128 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw128
			_811_v9 = _nw128
			var _812_v10 D1
			_ = _812_v10
			_812_v10 = Companion_D1_.Create_DC4_((_dafny.Zero).Minus(_788_v1), false, _811_v9)
			var _813_v11 _dafny.Map
			_ = _813_v11
			_813_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v1, (_812_v10).Dtor_cf5())
			var _814_v12 D3
			_ = _814_v12
			_814_v12 = Companion_D3_.Create_DC9_(_dafny.SetOf(func() _dafny.Map {
				var _coll28 = _dafny.NewMapBuilder()
				_ = _coll28
				for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(527), _dafny.IntOfInt64(468))); ; {
					_compr_28, _ok30 := _iter30()
					if !_ok30 {
						break
					}
					var _815_v8 _dafny.Int
					_815_v8 = interface{}(_compr_28).(_dafny.Int)
					if ((_dafny.IntOfInt64(527)).Cmp(_815_v8) <= 0) && ((_815_v8).Cmp(_dafny.IntOfInt64(468)) < 0) {
						_coll28.Add((_815_v8).Plus(_788_v1), (_this).F21())
					}
				}
				return _coll28.ToMap()
			}(), _813_v11))
			var _816_v13 _dafny.Set
			_ = _816_v13
			_816_v13 = _dafny.SetOf(_813_v11, _813_v11)
			_814_v12 = Companion_D3_.Create_DC9_(_816_v13)
			var _817_v14 _dafny.Array
			_ = _817_v14
			var _nwElement0_25 _dafny.Int = _788_v1
			_ = _nwElement0_25
			var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(5))
			_ = _nw129
			(_nw129).ArraySet1(_nwElement0_25, 0)
			(_nw129).ArraySet1(_dafny.IntOfInt64(418), 1)
			(_nw129).ArraySet1(_788_v1, 2)
			(_nw129).ArraySet1(_788_v1, 3)
			(_nw129).ArraySet1(_788_v1, 4)
			_817_v14 = _nw129
			var _818_v15 *C1
			_ = _818_v15
			var _nw130 *C1 = New_C1_()
			_ = _nw130
			_nw130.Ctor__(_817_v14)
			_818_v15 = _nw130
			var _819_v16 _dafny.Map
			_ = _819_v16
			_819_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v1, _788_v1)
			var _820_v17 _dafny.Sequence
			_ = _820_v17
			_820_v17 = _dafny.UnicodeSeqOfUtf8Bytes("csvwaexvd")
			_813_v11 = (_813_v11).Update((_dafny.Zero).Minus((func() _dafny.Int {
				if (_819_v16).Contains(_788_v1) {
					return (_819_v16).Get(_788_v1).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_820_v17).Cardinality())
			})()), !(p0))
			var _821_v18 _dafny.Set
			_ = _821_v18
			_821_v18 = _dafny.SetOf(p0)
			_821_v18 = _821_v18
		} else {
			var _822_v19 _dafny.Map
			_ = _822_v19
			_822_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v1, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _788_v1)).Update(p0, _788_v1)).Cardinality())
			var _823_v20 _dafny.Array
			_ = _823_v20
			var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
			_ = _nw131
			_823_v20 = _nw131
			var _824_v21 _dafny.Sequence
			_ = _824_v21
			_824_v21 = _dafny.SeqOf(_823_v20)
			var _825_v22 D1
			_ = _825_v22
			_825_v22 = Companion_D1_.Create_DC3_((_824_v21).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm20(globalState)).Cardinality(), _dafny.IntOfUint32((_824_v21).Cardinality()))).Uint32()).(_dafny.Array))
			var _826_v23 _dafny.MultiSet
			_ = _826_v23
			_826_v23 = _dafny.MultiSetOf(_825_v22)
			_822_v19 = (_822_v19).Update(Companion_Default___.SafeModuloInt(_788_v1, _788_v1), ((_826_v23).Cardinality()).Minus(_788_v1))
			_788_v1 = ((_dafny.Zero).Minus(_788_v1)).Plus((_788_v1).Minus(_dafny.IntOfUint32((p1).Cardinality())))
			var _827_v24 _dafny.Sequence
			_ = _827_v24
			_827_v24 = _dafny.SeqOf(_789_v2, _789_v2, _789_v2)
			var _828_v25 T0
			_ = _828_v25
			var _nw132 *C0 = New_C0_()
			_ = _nw132
			_nw132.Ctor__(_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("qgaehayk"), (_827_v24).Select((Companion_Default___.SafeIndex(_788_v1, _dafny.IntOfUint32((_827_v24).Cardinality()))).Uint32()).(_dafny.CodePoint)))
			_828_v25 = _nw132
			_828_v25 = _828_v25
			_827_v24 = _dafny.Companion_Sequence_.Concatenate(_827_v24, _827_v24)
			var _rhs89 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm9(_788_v1, _dafny.IntOfUint32((p1).Cardinality()), (_this).F21(), globalState), Companion_Default___.Fm9(_788_v1, (_791_v4).Select((Companion_Default___.SafeIndex(_788_v1, _dafny.IntOfUint32((_791_v4).Cardinality()))).Uint32()).(_dafny.Int), p0, globalState))
			_ = _rhs89
			var _rhs90 _dafny.Sequence = (func() _dafny.Sequence {
				if p0 {
					return _827_v24
				}
				return _827_v24
			})()
			_ = _rhs90
			_788_v1 = _rhs89
			_827_v24 = _rhs90
		}
		var _829_i0 _dafny.Int
		_ = _829_i0
		_829_i0 = _dafny.Zero
		{
			for (_this).F21() {
				{
					if (_829_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_829_i0 = (_829_i0).Plus(_dafny.One)
					if (func() bool {
						if p0 {
							return (_this).F21()
						}
						return !((_this).F21())
					})() {
						(globalState).F17 = !((_this).F21())
						var _830_v26 _dafny.Array
						_ = _830_v26
						var _len0_24 _dafny.Int = _dafny.IntOfInt64(23)
						_ = _len0_24
						var _nw133 _dafny.Array
						_ = _nw133
						if _len0_24.Cmp(_dafny.Zero) == 0 {
							_nw133 = _dafny.NewArray(_len0_24)
						} else {
							var _init24 func(_dafny.Int) bool = (func(_831_p0 bool) func(_dafny.Int) bool {
								return func(_832_i1 _dafny.Int) bool {
									return _831_p0
								}
							})(p0)
							_ = _init24
							var _element0_24 = _init24(_dafny.Zero)
							_ = _element0_24
							_nw133 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
							(_nw133).ArraySet1(_element0_24, 0)
							var _nativeLen0_24 = (_len0_24).Int()
							_ = _nativeLen0_24
							for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
								(_nw133).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
							}
						}
						_830_v26 = _nw133
						var _833_v27 _dafny.Map
						_ = _833_v27
						_833_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_830_v26, Companion_Default___.Fm9((_dafny.Zero).Minus(_788_v1), _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), p0, globalState))
						_833_v27 = (_833_v27).Update(_830_v26, (_788_v1).Minus(_788_v1))
						var _834_v28 _dafny.Array
						_ = _834_v28
						var _nw134 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
						_ = _nw134
						_834_v28 = _nw134
						_834_v28 = _834_v28
						var _835_v29 _dafny.Set
						_ = _835_v29
						_835_v29 = _dafny.SetOf(_dafny.IntOfUint32((_791_v4).Cardinality()), _788_v1)
						_788_v1 = Companion_Default___.Fm9(((_dafny.SetOf(_788_v1)).Union(_835_v29)).Cardinality(), Companion_Default___.SafeDivisionInt(_788_v1, _788_v1), !((_this).F21()), globalState)
						(globalState).F15 = (_this).F21()
					} else {
						var _836_v30 _dafny.Array
						_ = _836_v30
						var _nw135 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
						_ = _nw135
						_836_v30 = _nw135
						var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_836_v30), 0))
						_ = _index92
						(_836_v30).ArraySet1(p0, (_index92).Int())
						var _837_v31 _dafny.MultiSet
						_ = _837_v31
						_837_v31 = _dafny.MultiSetOf(_836_v30)
						var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_836_v30), 0))
						_ = _index93
						(_836_v30).ArraySet1(Companion_Default___.Fm0(!((_this).F21()), _789_v2, (_837_v31).IsSubsetOf(_837_v31), globalState), (_index93).Int())
						var _838_v32 _dafny.Sequence
						_ = _838_v32
						_838_v32 = _dafny.UnicodeSeqOfUtf8Bytes("cb")
						_838_v32 = _dafny.Companion_Sequence_.Concatenate(_838_v32, _dafny.Companion_Sequence_.Concatenate(_838_v32, _838_v32))
						_788_v1 = _788_v1
						var _839_v33 *C0
						_ = _839_v33
						var _nw136 *C0 = New_C0_()
						_ = _nw136
						_nw136.Ctor__((_this).F21())
						_839_v33 = _nw136
						var _840_v34 _dafny.Array
						_ = _840_v34
						var _nw137 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
						_ = _nw137
						_840_v34 = _nw137
						var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_840_v34), 0))
						_ = _index94
						(_840_v34).ArraySet1(_788_v1, (_index94).Int())
						var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_840_v34), 0))
						_ = _index95
						(_840_v34).ArraySet1(_788_v1, (_index95).Int())
					}
					if (p0) && ((_this).F21()) {
						var _841_v35 *C0
						_ = _841_v35
						var _nw138 *C0 = New_C0_()
						_ = _nw138
						_nw138.Ctor__((_this).F21())
						_841_v35 = _nw138
						(globalState).F17 = (_841_v35).F23()
						_789_v2 = _789_v2
						var _842_v36 _dafny.Array
						_ = _842_v36
						var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
						_ = _nw139
						_842_v36 = _nw139
						var _843_v37 _dafny.Sequence
						_ = _843_v37
						_843_v37 = _dafny.UnicodeSeqOfUtf8Bytes("nbewbhblg")
						var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_842_v36), 0))
						_ = _index96
						(_842_v36).ArraySet1(_843_v37, (_index96).Int())
						var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_842_v36), 0))
						_ = _index97
						(_842_v36).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jksnhhduo"), (_index97).Int())
						var _844_v38 _dafny.Set
						_ = _844_v38
						_844_v38 = _dafny.SetOf(_dafny.IntOfInt64(478))
						_788_v1 = (((_dafny.SetOf(_dafny.IntOfUint32(((_842_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_842_v36), 0))).Int()).(_dafny.Sequence)).Cardinality()), (_dafny.Zero).Minus((_dafny.Zero).Minus(_788_v1)))).Difference(_844_v38)).Intersection(func() _dafny.Set {
							var _coll29 = _dafny.NewBuilder()
							_ = _coll29
							for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-957), _dafny.IntOfInt64(919))); ; {
								_compr_29, _ok31 := _iter31()
								if !_ok31 {
									break
								}
								var _845_v39 _dafny.Int
								_845_v39 = interface{}(_compr_29).(_dafny.Int)
								if ((_dafny.IntOfInt64(-957)).Cmp(_845_v39) <= 0) && ((_845_v39).Cmp(_dafny.IntOfInt64(919)) < 0) {
									_coll29.Add((_845_v39).Times(_788_v1))
								}
							}
							return _coll29.ToSet()
						}())).Cardinality()
					} else {
						var _846_v40 _dafny.Sequence
						_ = _846_v40
						_846_v40 = _dafny.UnicodeSeqOfUtf8Bytes("p")
						var _847_v41 D1
						_ = _847_v41
						_847_v41 = Companion_D1_.Create_DC5_(_846_v40, (_dafny.Zero).Minus(_788_v1), _788_v1)
						_788_v1 = _dafny.IntOfUint32(((func() _dafny.Sequence {
							if p0 {
								return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(165))).Uint32(), func(coer41 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
									return func(arg41 _dafny.Int) interface{} {
										return coer41(arg41)
									}
								}((func(_848_v1 _dafny.Int) func(_dafny.Int) D1 {
									return func(_849_i2 _dafny.Int) D1 {
										return Companion_D1_.Create_DC5_(_dafny.UnicodeSeqOfUtf8Bytes("dyabgd"), _848_v1, _848_v1)
									}
								})(_788_v1)))
							}
							return _dafny.SeqOf(_847_v41, _847_v41)
						})()).Cardinality())
						_788_v1 = (_dafny.IntOfUint32((_791_v4).Cardinality())).Times(_788_v1)
						var _850_v42 *C0
						_ = _850_v42
						var _nw140 *C0 = New_C0_()
						_ = _nw140
						_nw140.Ctor__((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_846_v40, (Companion_Default___.SafeIndex(_788_v1, _dafny.IntOfUint32((_846_v40).Cardinality()))).Uint32(), _789_v2)).Cardinality())).Cmp(_788_v1) < 0)
						_850_v42 = _nw140
						var _851_v43 _dafny.Sequence
						_ = _851_v43
						_851_v43 = _dafny.SeqOf(_846_v40)
						_788_v1 = _dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_850_v42).F23() {
								return _dafny.Companion_Sequence_.Concatenate(_846_v40, _dafny.Companion_Sequence_.Update((_851_v43).Select((Companion_Default___.SafeIndex(_788_v1, _dafny.IntOfUint32((_851_v43).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_788_v1, _dafny.IntOfUint32(((_851_v43).Select((Companion_Default___.SafeIndex(_788_v1, _dafny.IntOfUint32((_851_v43).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _789_v2))
							}
							return _846_v40
						})()).Cardinality())
						var _rhs91 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_788_v1), _dafny.IntOfUint32(((_this).Fm8(p0, globalState)).Cardinality()))
						_ = _rhs91
						var _rhs92 bool = (p1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool)
						_ = _rhs92
						var _rhs93 _dafny.Int = (func() _dafny.Int {
							if (_850_v42).F23() {
								return _dafny.IntOfInt64(-637)
							}
							return _788_v1
						})()
						_ = _rhs93
						var _lhs48 *GlobalState = globalState
						_ = _lhs48
						_788_v1 = _rhs91
						_lhs48.F15 = _rhs92
						_788_v1 = _rhs93
					}
					var _852_v44 _dafny.Array
					_ = _852_v44
					var _len0_25 _dafny.Int = _dafny.IntOfInt64(22)
					_ = _len0_25
					var _nw141 _dafny.Array
					_ = _nw141
					if _len0_25.Cmp(_dafny.Zero) == 0 {
						_nw141 = _dafny.NewArray(_len0_25)
					} else {
						var _init25 func(_dafny.Int) bool = func(_853_i3 _dafny.Int) bool {
							return (_this).F21()
						}
						_ = _init25
						var _element0_25 = _init25(_dafny.Zero)
						_ = _element0_25
						_nw141 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
						(_nw141).ArraySet1(_element0_25, 0)
						var _nativeLen0_25 = (_len0_25).Int()
						_ = _nativeLen0_25
						for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
							(_nw141).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
						}
					}
					_852_v44 = _nw141
					_852_v44 = _852_v44
					(globalState).F11 = (_this).F21()
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _854_i4 _dafny.Int
		_ = _854_i4
		_854_i4 = _dafny.Zero
		{
			for (_788_v1).Cmp((_dafny.MultiSetOf(_788_v1, _788_v1, _788_v1)).Cardinality()) != 0 {
				{
					if (_854_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_854_i4 = (_854_i4).Plus(_dafny.One)
					(globalState).F17 = (_this).F21()
					var _855_v45 _dafny.Array
					_ = _855_v45
					var _len0_26 _dafny.Int = _dafny.IntOfInt64(13)
					_ = _len0_26
					var _nw142 _dafny.Array
					_ = _nw142
					if _len0_26.Cmp(_dafny.Zero) == 0 {
						_nw142 = _dafny.NewArray(_len0_26)
					} else {
						var _init26 func(_dafny.Int) _dafny.Sequence = (func(_856_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_857_i5 _dafny.Int) _dafny.Sequence {
								return _856_v4
							}
						})(_791_v4)
						_ = _init26
						var _element0_26 = _init26(_dafny.Zero)
						_ = _element0_26
						_nw142 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
						(_nw142).ArraySet1(_element0_26, 0)
						var _nativeLen0_26 = (_len0_26).Int()
						_ = _nativeLen0_26
						for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
							(_nw142).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
						}
					}
					_855_v45 = _nw142
					var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_855_v45), 0))
					_ = _index98
					(_855_v45).ArraySet1(Companion_Default___.Fm19(_791_v4, p0, globalState), (_index98).Int())
					var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_855_v45), 0))
					_ = _index99
					(_855_v45).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_791_v4, _791_v4), (_index99).Int())
					var _rhs94 bool = !_dafny.Companion_Sequence_.Contains((_855_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_855_v45), 0))).Int()).(_dafny.Sequence), _788_v1)
					_ = _rhs94
					var _rhs95 _dafny.Int = _dafny.IntOfInt64(459)
					_ = _rhs95
					var _lhs49 *GlobalState = globalState
					_ = _lhs49
					_lhs49.F11 = _rhs94
					_788_v1 = _rhs95
					_788_v1 = Companion_Default___.Fm9(_788_v1, (func() _dafny.Int {
						if (_this).F21() {
							return _788_v1
						}
						return _788_v1
					})(), p0, globalState)
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _858_v46 _dafny.Array
		_ = _858_v46
		var _nw143 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
		_ = _nw143
		_858_v46 = _nw143
		var _859_v47 D1
		_ = _859_v47
		_859_v47 = Companion_D1_.Create_DC3_(_858_v46)
		var _860_v48 _dafny.Array
		_ = _860_v48
		var _nwElement0_26 _dafny.Array = _858_v46
		_ = _nwElement0_26
		var _nw144 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(7))
		_ = _nw144
		(_nw144).ArraySet1(_nwElement0_26, 0)
		(_nw144).ArraySet1(_858_v46, 1)
		(_nw144).ArraySet1(_858_v46, 2)
		(_nw144).ArraySet1(_858_v46, 3)
		(_nw144).ArraySet1(_858_v46, 4)
		(_nw144).ArraySet1(_858_v46, 5)
		(_nw144).ArraySet1((_859_v47).Dtor_cf3(), 6)
		_860_v48 = _nw144
		var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_860_v48), 0))
		_ = _index100
		(_860_v48).ArraySet1(_858_v46, (_index100).Int())
		var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_860_v48), 0))
		_ = _index101
		(_860_v48).ArraySet1(_858_v46, (_index101).Int())
		if (_this).F21() {
			var _861_v49 _dafny.Sequence
			_ = _861_v49
			_861_v49 = _dafny.UnicodeSeqOfUtf8Bytes("trlsw")
			_788_v1 = (_dafny.IntOfUint32((_861_v49).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_791_v4, _791_v4)).Cardinality()))
			if !((p1).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm9(_788_v1, _dafny.IntOfUint32((p1).Cardinality()), !((_this).F21()), globalState), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool)) {
				var _862_v50 _dafny.Array
				_ = _862_v50
				var _nw145 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
				_ = _nw145
				_862_v50 = _nw145
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_862_v50), 0))
				_ = _index102
				(_862_v50).ArraySet1(p0, (_index102).Int())
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_862_v50), 0))
				_ = _index103
				var _rhs96 bool = (_this).F21()
				_ = _rhs96
				var _rhs97 _dafny.Int = _788_v1
				_ = _rhs97
				var _lhs50 _dafny.Array = _862_v50
				_ = _lhs50
				var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_862_v50), 0))
				_ = _lhs51
				(_lhs50).ArraySet1(_rhs96, (_lhs51).Int())
				_788_v1 = _rhs97
				_788_v1 = (_dafny.Zero).Minus(_788_v1)
				_862_v50 = _862_v50
				var _863_v51 _dafny.MultiSet
				_ = _863_v51
				_863_v51 = _dafny.MultiSetOf(p0, (_862_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_862_v50), 0))).Int()).(bool), (_this).F21())
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_862_v50), 0))
				_ = _index104
				(_862_v50).ArraySet1((_863_v51).IsSubsetOf(_863_v51), (_index104).Int())
				_788_v1 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(468), _788_v1)
			} else {
				_861_v49 = (Companion_Default___.Fm21(globalState)).Dtor_cf7()
				var _864_v52 _dafny.Array
				_ = _864_v52
				var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
				_ = _nw146
				_864_v52 = _nw146
				var _865_v53 _dafny.MultiSet
				_ = _865_v53
				_865_v53 = _dafny.MultiSetOf(_864_v52)
				(globalState).F15 = (func() bool {
					if (_dafny.MultiSetOf(_864_v52, _864_v52)).IsSubsetOf(_865_v53) {
						return (_this).F21()
					}
					return (_this).F21()
				})()
				var _866_v54 _dafny.Set
				_ = _866_v54
				_866_v54 = _dafny.SetOf(_dafny.SeqOf(p0, p0))
				var _867_v55 _dafny.MultiSet
				_ = _867_v55
				_867_v55 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("qn"))
				var _868_v56 _dafny.Sequence
				_ = _868_v56
				_868_v56 = _dafny.SeqOf(_867_v55)
				var _869_v57 _dafny.Array
				_ = _869_v57
				var _nwElement0_27 bool = p0
				_ = _nwElement0_27
				var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(20))
				_ = _nw147
				(_nw147).ArraySet1(_nwElement0_27, 0)
				(_nw147).ArraySet1((_this).F21(), 1)
				(_nw147).ArraySet1(false, 2)
				(_nw147).ArraySet1((_866_v54).IsProperSubsetOf(_866_v54), 3)
				(_nw147).ArraySet1(Companion_Default___.Fm0((_this).F21(), _789_v2, (_this).F21(), globalState), 4)
				(_nw147).ArraySet1(p0, 5)
				(_nw147).ArraySet1(p0, 6)
				(_nw147).ArraySet1((_this).F21(), 7)
				(_nw147).ArraySet1((_867_v55).IsDisjointFrom((_868_v56).Select((Companion_Default___.SafeIndex(_788_v1, _dafny.IntOfUint32((_868_v56).Cardinality()))).Uint32()).(_dafny.MultiSet)), 8)
				(_nw147).ArraySet1(!(!(false)), 9)
				(_nw147).ArraySet1((_this).F21(), 10)
				(_nw147).ArraySet1(p0, 11)
				(_nw147).ArraySet1(p0, 12)
				(_nw147).ArraySet1((_this).F21(), 13)
				(_nw147).ArraySet1(true, 14)
				(_nw147).ArraySet1(p0, 15)
				(_nw147).ArraySet1(p0, 16)
				(_nw147).ArraySet1(!((_this).F21()), 17)
				(_nw147).ArraySet1(Companion_Default___.Fm0((_this).F21(), _789_v2, p0, globalState), 18)
				(_nw147).ArraySet1(true, 19)
				_869_v57 = _nw147
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_869_v57), 0))
				_ = _index105
				(_869_v57).ArraySet1(p0, (_index105).Int())
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_869_v57), 0))
				_ = _index106
				(_869_v57).ArraySet1(((_dafny.Zero).Minus(_788_v1)).Cmp((_788_v1).Plus(_dafny.IntOfInt64(-440))) >= 0, (_index106).Int())
				var _870_v58 _dafny.Map
				_ = _870_v58
				_870_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), p0)
				var _871_v59 _dafny.Map
				_ = _871_v59
				_871_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_870_v58).Merge((_870_v58).Update((_869_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_869_v57), 0))).Int()).(bool), (_this).F21())))
				_871_v59 = (_871_v59).Update((_this).F21(), (Companion_Default___.Fm22(globalState)).Update(p0, (_this).F21()))
				var _872_v60 _dafny.Map
				_ = _872_v60
				_872_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v1, _dafny.IntOfUint32((_861_v49).Cardinality()))
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_869_v57), 0))
				_ = _index107
				(_869_v57).ArraySet1(!(!(_872_v60).Contains(_788_v1)), (_index107).Int())
			}
			(globalState).F15 = Companion_Default___.Fm0((_788_v1).Cmp(_788_v1) != 0, _789_v2, (_this).F21(), globalState)
			var _873_v61 _dafny.Map
			_ = _873_v61
			_873_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _788_v1)
			_873_v61 = (_873_v61).Update(_dafny.Companion_Sequence_.Concatenate(p1, p1), (_788_v1).Plus(_788_v1))
			var _874_v62 _dafny.MultiSet
			_ = _874_v62
			_874_v62 = _dafny.MultiSetOf(_788_v1)
			var _875_v63 _dafny.Set
			_ = _875_v63
			_875_v63 = _dafny.SetOf(_874_v62, _dafny.MultiSetFromSeq(_791_v4), _874_v62, _874_v62)
			var _876_v64 _dafny.Sequence
			_ = _876_v64
			_876_v64 = _dafny.SeqOf(_dafny.SetOf(_874_v62, _874_v62), _875_v63, _875_v63, _875_v63, _875_v63)
			var _877_v65 _dafny.MultiSet
			_ = _877_v65
			_877_v65 = _dafny.MultiSetOf(((_876_v64).Select((Companion_Default___.SafeIndex(_788_v1, _dafny.IntOfUint32((_876_v64).Cardinality()))).Uint32()).(_dafny.Set)).IsProperSubsetOf(_875_v63), (_this).F21(), p0, (_this).F21())
			_877_v65 = _877_v65
		} else {
			if ((Companion_Default___.Fm9(_788_v1, _788_v1, (_this).F21(), globalState)).Minus(_788_v1)).Cmp((func() _dafny.Int {
				if (_this).F21() {
					return _788_v1
				}
				return _788_v1
			})()) != 0 {
				var _878_v66 _dafny.Sequence
				_ = _878_v66
				_878_v66 = _dafny.UnicodeSeqOfUtf8Bytes("upxiueqmn")
				var _879_v67 _dafny.Map
				_ = _879_v67
				_879_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_792_v5, _878_v66)
				_879_v67 = (_879_v67).Update(_792_v5, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(581))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_880_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_881_i6 _dafny.Int) _dafny.CodePoint {
						return _880_v2
					}
				})(_789_v2))))
				var _882_v68 _dafny.Map
				_ = _882_v68
				_882_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _788_v1)
				var _883_v69 _dafny.MultiSet
				_ = _883_v69
				_883_v69 = _dafny.MultiSetOf(_788_v1)
				var _884_v70 _dafny.Sequence
				_ = _884_v70
				_884_v70 = _dafny.SeqOf(_dafny.MultiSetOf(_788_v1), _dafny.MultiSetOf(_dafny.IntOfUint32((_878_v66).Cardinality())), _883_v69, _883_v69, _dafny.MultiSetFromSeq(_791_v4))
				_882_v68 = (_882_v68).Update((Companion_Default___.Fm23((_this).F21(), _882_v68, Companion_Default___.Fm9(_788_v1, _788_v1, p0, globalState), _789_v2, globalState)).IsProperSubsetOf((_884_v70).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_883_v69).Cardinality()), _dafny.IntOfUint32((_884_v70).Cardinality()))).Uint32()).(_dafny.MultiSet)), _788_v1)
				(globalState).F11 = Companion_Default___.Fm0((_this).F21(), _dafny.CodePoint('x'), false, globalState)
				(globalState).F11 = p0
				var _885_v71 _dafny.Array
				_ = _885_v71
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_27
				var _nw148 _dafny.Array
				_ = _nw148
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw148 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) _dafny.Map = (func(_886_v1 _dafny.Int, _887_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
						return func(_888_i7 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.SeqOf(_886_v1), _887_v4)).Cardinality(), _886_v1)
						}
					})(_788_v1, _791_v4)
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw148 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw148).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw148).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_885_v71 = _nw148
				var _889_v72 _dafny.Map
				_ = _889_v72
				_889_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v1, (_883_v69).Cardinality())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_885_v71), 0))
				_ = _index108
				(_885_v71).ArraySet1(_889_v72, (_index108).Int())
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_885_v71), 0))
				_ = _index109
				(_885_v71).ArraySet1(_889_v72, (_index109).Int())
			} else {
				var _890_v73 _dafny.Array
				_ = _890_v73
				var _nw149 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw149
				_890_v73 = _nw149
				var _891_v74 _dafny.Set
				_ = _891_v74
				_891_v74 = _dafny.SetOf(_890_v73)
				var _892_v75 _dafny.Sequence
				_ = _892_v75
				_892_v75 = _dafny.SeqOf(_890_v73, _890_v73)
				var _893_v76 _dafny.Array
				_ = _893_v76
				var _nwElement0_28 _dafny.Set = _891_v74
				_ = _nwElement0_28
				var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(16))
				_ = _nw150
				(_nw150).ArraySet1(_nwElement0_28, 0)
				(_nw150).ArraySet1((_891_v74).Intersection(_891_v74), 1)
				(_nw150).ArraySet1(_891_v74, 2)
				(_nw150).ArraySet1(_891_v74, 3)
				(_nw150).ArraySet1(_891_v74, 4)
				(_nw150).ArraySet1(_891_v74, 5)
				(_nw150).ArraySet1(_dafny.SetOf(_890_v73, _890_v73, _890_v73, _890_v73, (_892_v75).Select((Companion_Default___.SafeIndex(_788_v1, _dafny.IntOfUint32((_892_v75).Cardinality()))).Uint32()).(_dafny.Array)), 6)
				(_nw150).ArraySet1(_891_v74, 7)
				(_nw150).ArraySet1((_891_v74).Difference(_891_v74), 8)
				(_nw150).ArraySet1(_891_v74, 9)
				(_nw150).ArraySet1(_891_v74, 10)
				(_nw150).ArraySet1(_dafny.SetOf(_890_v73), 11)
				(_nw150).ArraySet1(_891_v74, 12)
				(_nw150).ArraySet1(_891_v74, 13)
				(_nw150).ArraySet1(_891_v74, 14)
				(_nw150).ArraySet1(_dafny.SetOf(_890_v73), 15)
				_893_v76 = _nw150
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_893_v76), 0))
				_ = _index110
				(_893_v76).ArraySet1(_891_v74, (_index110).Int())
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_893_v76), 0))
				_ = _index111
				(_893_v76).ArraySet1(_891_v74, (_index111).Int())
				_788_v1 = _788_v1
				var _894_v77 *C0
				_ = _894_v77
				var _nw151 *C0 = New_C0_()
				_ = _nw151
				_nw151.Ctor__(false)
				_894_v77 = _nw151
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((_890_v73), 0))
				_ = _index112
				(_890_v73).ArraySet1((_this).F21(), (_index112).Int())
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_890_v73), 0))
				_ = _index113
				(_890_v73).ArraySet1(true, (_index113).Int())
				var _895_v78 _dafny.Array
				_ = _895_v78
				var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
				_ = _nw152
				_895_v78 = _nw152
				var _896_v79 _dafny.MultiSet
				_ = _896_v79
				_896_v79 = _dafny.MultiSetOf(_788_v1, _788_v1)
				var _897_v80 _dafny.Sequence
				_ = _897_v80
				_897_v80 = _dafny.SeqOf(_896_v79, _896_v79)
				var _898_v81 _dafny.Map
				_ = _898_v81
				_898_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v1, p0)
				var _899_v82 _dafny.Set
				_ = _899_v82
				_899_v82 = _dafny.SetOf(_788_v1, ((_896_v79).Update((_791_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_791_v4).Cardinality()), _dafny.IntOfUint32((_791_v4).Cardinality()))).Uint32()).(_dafny.Int), Companion_Default___.Abs(_788_v1))).Cardinality(), ((_896_v79).Intersection((_897_v80).Select((Companion_Default___.SafeIndex((_898_v81).Cardinality(), _dafny.IntOfUint32((_897_v80).Cardinality()))).Uint32()).(_dafny.MultiSet))).Cardinality())
				var _900_v83 _dafny.MultiSet
				_ = _900_v83
				_900_v83 = _dafny.MultiSetOf(p1, p1, p1)
				var _901_v84 _dafny.Map
				_ = _901_v84
				_901_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_900_v83, _788_v1)
				var _902_v85 _dafny.Map
				_ = _902_v85
				_902_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _788_v1)
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((_890_v73), 0))
				_ = _index114
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_890_v73), 0))
				_ = _index115
				var _rhs98 bool = p0
				_ = _rhs98
				var _rhs99 bool = Companion_Default___.Fm0(!((_894_v77).Fm15((_894_v77).F23(), p0, (func() _dafny.Int {
					if (_901_v84).Contains((_900_v83).Update(_dafny.SeqOf(p0), Companion_Default___.Abs(_788_v1))) {
						return (_901_v84).Get((_900_v83).Update(_dafny.SeqOf(p0), Companion_Default___.Abs(_788_v1))).(_dafny.Int)
					}
					return _788_v1
				})(), Companion_Default___.Fm24((_894_v77).Fm15(p0, p0, _dafny.IntOfUint32((_791_v4).Cardinality()), _902_v85, globalState), (_899_v82).Cardinality(), globalState), globalState)), (_792_v5).Dtor_cf26(), true, globalState)
				_ = _rhs99
				var _rhs100 _dafny.Array = _895_v78
				_ = _rhs100
				var _rhs101 _dafny.Set = (_899_v82).Union((func() _dafny.Set {
					var _coll30 = _dafny.NewBuilder()
					_ = _coll30
					for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(961), _dafny.IntOfInt64(442))); ; {
						_compr_30, _ok32 := _iter32()
						if !_ok32 {
							break
						}
						var _903_v86 _dafny.Int
						_903_v86 = interface{}(_compr_30).(_dafny.Int)
						if ((_dafny.IntOfInt64(961)).Cmp(_903_v86) <= 0) && ((_903_v86).Cmp(_dafny.IntOfInt64(442)) < 0) {
							_coll30.Add((_903_v86).Times(_788_v1))
						}
					}
					return _coll30.ToSet()
				}()).Difference(_dafny.SetOf(_788_v1, _dafny.IntOfInt64(282), _788_v1)))
				_ = _rhs101
				var _lhs52 _dafny.Array = _890_v73
				_ = _lhs52
				var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((_890_v73), 0))
				_ = _lhs53
				var _lhs54 _dafny.Array = _890_v73
				_ = _lhs54
				var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_890_v73), 0))
				_ = _lhs55
				(_lhs52).ArraySet1(_rhs98, (_lhs53).Int())
				(_lhs54).ArraySet1(_rhs99, (_lhs55).Int())
				_895_v78 = _rhs100
				_899_v82 = _rhs101
				var _904_v87 _dafny.Map
				_ = _904_v87
				_904_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_789_v2), _788_v1)
				var _905_v88 _dafny.Sequence
				_ = _905_v88
				_905_v88 = _dafny.SeqOf(_789_v2, _789_v2)
				_904_v87 = (_904_v87).Update(_905_v88, _788_v1)
			}
			var _906_v89 _dafny.Map
			_ = _906_v89
			_906_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v1, (_this).F21())
			var _907_v90 _dafny.Sequence
			_ = _907_v90
			_907_v90 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v1, (_this).F21()))
			var _908_v91 _dafny.Set
			_ = _908_v91
			_908_v91 = _dafny.SetOf(_906_v89, (_907_v90).Select((Companion_Default___.SafeIndex(_788_v1, _dafny.IntOfUint32((_907_v90).Cardinality()))).Uint32()).(_dafny.Map))
			var _909_v92 D3
			_ = _909_v92
			_909_v92 = Companion_D3_.Create_DC9_((_908_v91).Difference(_908_v91))
			_909_v92 = _909_v92
			var _910_v93 *C0
			_ = _910_v93
			var _nw153 *C0 = New_C0_()
			_ = _nw153
			_nw153.Ctor__((_this).F21())
			_910_v93 = _nw153
			var _911_v94 _dafny.Map
			_ = _911_v94
			_911_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus(_788_v1))
			var _912_v95 _dafny.Array
			_ = _912_v95
			var _nwElement0_29 bool = (_this).F21()
			_ = _nwElement0_29
			var _nw154 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(12))
			_ = _nw154
			(_nw154).ArraySet1(_nwElement0_29, 0)
			(_nw154).ArraySet1(p0, 1)
			(_nw154).ArraySet1(p0, 2)
			(_nw154).ArraySet1(((_910_v93).Fm14((_dafny.Zero).Minus(_788_v1), globalState)).Cmp(_788_v1) == 0, 3)
			(_nw154).ArraySet1(p0, 4)
			(_nw154).ArraySet1((_this).F21(), 5)
			(_nw154).ArraySet1(p0, 6)
			(_nw154).ArraySet1(false, 7)
			(_nw154).ArraySet1((_this).F21(), 8)
			(_nw154).ArraySet1((_this).F21(), 9)
			(_nw154).ArraySet1((_910_v93).F23(), 10)
			(_nw154).ArraySet1((_910_v93).Fm15(!((_910_v93).F23()), !((_910_v93).F23()), (_910_v93).Fm14(_788_v1, globalState), _911_v94, globalState), 11)
			_912_v95 = _nw154
			_912_v95 = (func() _dafny.Array {
				if (_this).F21() {
					return _912_v95
				}
				return (func() _dafny.Array {
					if p0 {
						return _912_v95
					}
					return _912_v95
				})()
			})()
			var _source11 D0 = _this.F20
			_ = _source11
			if _source11.Is_DC1() {
				var _913___mcc_h7 bool = _source11.Get_().(D0_DC1).Cf1
				_ = _913___mcc_h7
				var _914_cf1 bool = _913___mcc_h7
				_ = _914_cf1
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_912_v95), 0))
				_ = _index116
				(_912_v95).ArraySet1(p0, (_index116).Int())
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_912_v95), 0))
				_ = _index117
				(_912_v95).ArraySet1((_910_v93).Fm15((_910_v93).F23(), ((_this).F21()) || ((_910_v93).F23()), _dafny.IntOfInt64(395), _911_v94, globalState), (_index117).Int())
				var _915_v96 _dafny.Sequence
				_ = _915_v96
				_915_v96 = _dafny.UnicodeSeqOfUtf8Bytes("aln")
				var _916_v97 _dafny.Map
				_ = _916_v97
				_916_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v1, _915_v96)
				(globalState).F15 = _dafny.Companion_Sequence_.IsPrefixOf(_915_v96, (func() _dafny.Sequence {
					if (_916_v97).Contains(_788_v1) {
						return (_916_v97).Get(_788_v1).(_dafny.Sequence)
					}
					return _915_v96
				})())
				var _917_v98 _dafny.Array
				_ = _917_v98
				var _nw155 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
				_ = _nw155
				_917_v98 = _nw155
				_917_v98 = _917_v98
				var _918_v99 _dafny.Array
				_ = _918_v99
				var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw156
				_918_v99 = _nw156
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_918_v99), 0))
				_ = _index118
				(_918_v99).ArraySet1(_788_v1, (_index118).Int())
				var _919_v100 _dafny.MultiSet
				_ = _919_v100
				_919_v100 = _dafny.MultiSetOf(_788_v1, _788_v1, _788_v1)
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_918_v99), 0))
				_ = _index119
				var _rhs102 _dafny.Int = Companion_Default___.SafeModuloInt(_788_v1, _788_v1)
				_ = _rhs102
				var _rhs103 _dafny.Int = _788_v1
				_ = _rhs103
				var _rhs104 _dafny.MultiSet = (_919_v100).Difference(_919_v100)
				_ = _rhs104
				var _lhs56 _dafny.Array = _918_v99
				_ = _lhs56
				var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_918_v99), 0))
				_ = _lhs57
				(_lhs56).ArraySet1(_rhs102, (_lhs57).Int())
				_788_v1 = _rhs103
				_919_v100 = _rhs104
			} else if _source11.Is_DC0() {
				var _920___mcc_h8 bool = _source11.Get_().(D0_DC0).Cf0
				_ = _920___mcc_h8
				var _921_cf0 bool = _920___mcc_h8
				_ = _921_cf0
				var _922_v101 _dafny.Array
				_ = _922_v101
				var _nw157 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
				_ = _nw157
				_922_v101 = _nw157
				var _923_v102 _dafny.Sequence
				_ = _923_v102
				_923_v102 = _dafny.UnicodeSeqOfUtf8Bytes("uajufq")
				var _924_v103 _dafny.Array
				_ = _924_v103
				var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw158
				_924_v103 = _nw158
				var _925_v104 _dafny.Map
				_ = _925_v104
				_925_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_923_v102, _924_v103)
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_922_v101), 0))
				_ = _index120
				(_922_v101).ArraySet1((func() _dafny.Array {
					if (_925_v104).Contains(_923_v102) {
						return (_925_v104).Get(_923_v102).(_dafny.Array)
					}
					return _924_v103
				})(), (_index120).Int())
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_922_v101), 0))
				_ = _index121
				(_922_v101).ArraySet1(_924_v103, (_index121).Int())
				(globalState).F17 = (_792_v5).Dtor_cf23()
				(globalState).F15 = (_910_v93).F23()
				_788_v1 = _788_v1
			} else {
				var _926___mcc_h9 D0 = _source11.Get_().(D0_DC2).Cf2
				_ = _926___mcc_h9
				var _927_cf2 D0 = _926___mcc_h9
				_ = _927_cf2
				r0 = !(((_this).F21()) && (!((_910_v93).F23())))
				var _928_v105 _dafny.Array
				_ = _928_v105
				var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw159
				_928_v105 = _nw159
				var _929_v106 D3
				_ = _929_v106
				_929_v106 = Companion_D3_.Create_DC11_(_911_v94, _788_v1, _928_v105)
				var _930_v107 _dafny.MultiSet
				_ = _930_v107
				_930_v107 = _dafny.MultiSetOf(Companion_D3_.Create_DC11_(_911_v94, _788_v1, _928_v105), _929_v106, _929_v106, _929_v106)
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_912_v95), 0))
				_ = _index122
				(_912_v95).ArraySet1((_930_v107).IsProperSubsetOf(_930_v107), (_index122).Int())
				var _931_v108 _dafny.Sequence
				_ = _931_v108
				_931_v108 = _dafny.UnicodeSeqOfUtf8Bytes("vpii")
				var _932_v109 _dafny.Set
				_ = _932_v109
				_932_v109 = _dafny.SetOf((_910_v93).F23(), Companion_Default___.Fm0(p0, _789_v2, (_910_v93).F23(), globalState))
				var _933_v110 _dafny.Map
				_ = _933_v110
				_933_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_931_v108, _932_v109)
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_912_v95), 0))
				_ = _index123
				(_912_v95).ArraySet1(((_933_v110).Merge(_933_v110)).Equals(_933_v110), (_index123).Int())
				var _934_v111 D1
				_ = _934_v111
				_934_v111 = Companion_D1_.Create_DC5_(_931_v108, _788_v1, (_dafny.Zero).Minus(Companion_Default___.Fm9(_788_v1, _788_v1, p0, globalState)))
				var _935_v113 _dafny.MultiSet
				_ = _935_v113
				_935_v113 = _dafny.MultiSetOf(_dafny.IntOfInt64(740))
				var _936_v114 _dafny.MultiSet
				_ = _936_v114
				_936_v114 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt(_788_v1, (func() _dafny.Map {
					var _coll31 = _dafny.NewMapBuilder()
					_ = _coll31
					for _iter33 := _dafny.Iterate((_935_v113).Elements()); ; {
						_compr_31, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _937_v112 _dafny.Int
						_937_v112 = interface{}(_compr_31).(_dafny.Int)
						if (_935_v113).Contains(_937_v112) {
							_coll31.Add((_937_v112).Times(_788_v1), p0)
						}
					}
					return _coll31.ToMap()
				}()).Cardinality()), _788_v1)
				var _rhs105 D1 = _934_v111
				_ = _rhs105
				var _rhs106 _dafny.Int = _788_v1
				_ = _rhs106
				var _rhs107 _dafny.MultiSet = _935_v113
				_ = _rhs107
				var _rhs108 _dafny.Int = (func() _dafny.Int {
					if true {
						return _dafny.IntOfInt64(446)
					}
					return (_935_v113).Cardinality()
				})()
				_ = _rhs108
				_934_v111 = _rhs105
				_788_v1 = _rhs106
				_935_v113 = _rhs107
				_788_v1 = _rhs108
				var _rhs109 _dafny.CodePoint = _789_v2
				_ = _rhs109
				var _rhs110 _dafny.Int = (_dafny.Zero).Minus(_788_v1)
				_ = _rhs110
				var _rhs111 _dafny.Int = _788_v1
				_ = _rhs111
				var _rhs112 _dafny.CodePoint = _789_v2
				_ = _rhs112
				var _rhs113 _dafny.Array = _928_v105
				_ = _rhs113
				_789_v2 = _rhs109
				_788_v1 = _rhs110
				_788_v1 = _rhs111
				_789_v2 = _rhs112
				_928_v105 = _rhs113
			}
		}
		r0 = !(!((_this).F21()))
		return r0
	}
}
func (_this *C2) M6(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) (bool, bool, bool, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _938_i0 _dafny.Int
		_ = _938_i0
		_938_i0 = _dafny.Zero
		{
			for p2 {
				{
					if (_938_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_938_i0 = (_938_i0).Plus(_dafny.One)
					if (_this).F21() {
						var _939_v0 D3
						_ = _939_v0
						_939_v0 = Companion_D3_.Create_DC10_(p0)
						var _940_v1 _dafny.Sequence
						_ = _940_v1
						_940_v1 = _dafny.SeqOf(false, (_this).F21())
						var _941_v2 _dafny.Map
						_ = _941_v2
						_941_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _940_v1)
						var _942_v3 _dafny.Map
						_ = _942_v3
						_942_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p1).Cardinality()), _941_v2)
						var _943_v4 _dafny.Map
						_ = _943_v4
						_943_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm9((_939_v0).Dtor_cf17(), p0, p2, globalState), (func() _dafny.Map {
							if (_942_v3).Contains(p0) {
								return (_942_v3).Get(p0).(_dafny.Map)
							}
							return _941_v2
						})())
						var _944_v5 _dafny.Sequence
						_ = _944_v5
						_944_v5 = _dafny.SeqOf(p0)
						var _945_v6 _dafny.Set
						_ = _945_v6
						_945_v6 = _dafny.SetOf((_dafny.SetOf(_dafny.IntOfInt64(154))).Cardinality())
						var _946_v7 _dafny.Map
						_ = _946_v7
						_946_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_945_v6).Cardinality())
						var _947_v8 _dafny.Map
						_ = _947_v8
						_947_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
						var _948_v9 D1
						_ = _948_v9
						_948_v9 = Companion_D1_.Create_DC5_(p1, p0, _dafny.IntOfUint32((p1).Cardinality()))
						var _949_v10 _dafny.Array
						_ = _949_v10
						var _nwElement0_30 _dafny.Int = p0
						_ = _nwElement0_30
						var _nw160 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(22))
						_ = _nw160
						(_nw160).ArraySet1(_nwElement0_30, 0)
						(_nw160).ArraySet1(p0, 1)
						(_nw160).ArraySet1(p0, 2)
						(_nw160).ArraySet1(((func() _dafny.Map {
							if (_942_v3).Contains(_dafny.IntOfUint32((p1).Cardinality())) {
								return (_942_v3).Get(_dafny.IntOfUint32((p1).Cardinality())).(_dafny.Map)
							}
							return _941_v2
						})()).Cardinality(), 3)
						(_nw160).ArraySet1((_dafny.MultiSetOf((_this).F21())).Cardinality(), 4)
						(_nw160).ArraySet1(p0, 5)
						(_nw160).ArraySet1(p0, 6)
						(_nw160).ArraySet1(p0, 7)
						(_nw160).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 8)
						(_nw160).ArraySet1(p0, 9)
						(_nw160).ArraySet1(p0, 10)
						(_nw160).ArraySet1(p0, 11)
						(_nw160).ArraySet1((_dafny.Zero).Minus((_944_v5).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm9((_946_v7).Cardinality(), p0, (_this).F21(), globalState), _dafny.IntOfUint32((_944_v5).Cardinality()))).Uint32()).(_dafny.Int)), 12)
						(_nw160).ArraySet1(p0, 13)
						(_nw160).ArraySet1(_dafny.IntOfUint32((_940_v1).Cardinality()), 14)
						(_nw160).ArraySet1(p0, 15)
						(_nw160).ArraySet1((_947_v8).Cardinality(), 16)
						(_nw160).ArraySet1(p0, 17)
						(_nw160).ArraySet1((_948_v9).Dtor_cf9(), 18)
						(_nw160).ArraySet1(p0, 19)
						(_nw160).ArraySet1(p0, 20)
						(_nw160).ArraySet1(_dafny.IntOfInt64(659), 21)
						_949_v10 = _nw160
						var _950_v11 *C1
						_ = _950_v11
						var _nw161 *C1 = New_C1_()
						_ = _nw161
						_nw161.Ctor__(_949_v10)
						_950_v11 = _nw161
						var _951_v12 _dafny.Array
						_ = _951_v12
						var _nw162 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(18))
						_ = _nw162
						_951_v12 = _nw162
						var _952_v13 _dafny.Array
						_ = _952_v13
						var _len0_28 _dafny.Int = _dafny.IntOfInt64(14)
						_ = _len0_28
						var _nw163 _dafny.Array
						_ = _nw163
						if _len0_28.Cmp(_dafny.Zero) == 0 {
							_nw163 = _dafny.NewArray(_len0_28)
						} else {
							var _init28 func(_dafny.Int) bool = (func(_953_p2 bool) func(_dafny.Int) bool {
								return func(_954_i1 _dafny.Int) bool {
									return _953_p2
								}
							})(p2)
							_ = _init28
							var _element0_28 = _init28(_dafny.Zero)
							_ = _element0_28
							_nw163 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
							(_nw163).ArraySet1(_element0_28, 0)
							var _nativeLen0_28 = (_len0_28).Int()
							_ = _nativeLen0_28
							for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
								(_nw163).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
							}
						}
						_952_v13 = _nw163
						var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_952_v13), 0))
						_ = _index124
						(_952_v13).ArraySet1((_this).F21(), (_index124).Int())
						var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_952_v13), 0))
						_ = _index125
						var _rhs114 _dafny.Array = _951_v12
						_ = _rhs114
						var _rhs115 bool = (_this).F21()
						_ = _rhs115
						var _lhs58 _dafny.Array = _952_v13
						_ = _lhs58
						var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_952_v13), 0))
						_ = _lhs59
						_951_v12 = _rhs114
						(_lhs58).ArraySet1(_rhs115, (_lhs59).Int())
						var _955_v14 *C1
						_ = _955_v14
						var _nw164 *C1 = New_C1_()
						_ = _nw164
						_nw164.Ctor__((_950_v11).F22())
						_955_v14 = _nw164
						_952_v13 = _952_v13
						r2 = p2
					} else {
						var _956_v15 _dafny.Int
						_ = _956_v15
						_956_v15 = _dafny.IntOfInt64(714)
						var _957_v16 _dafny.Sequence
						_ = _957_v16
						_957_v16 = _dafny.SeqOf(!(p2))
						var _958_v17 _dafny.Sequence
						_ = _958_v17
						_958_v17 = _dafny.SeqOf(_957_v16)
						var _959_v18 _dafny.Sequence
						_ = _959_v18
						_959_v18 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(857))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg43 _dafny.Int) interface{} {
								return coer43(arg43)
							}
						}((func(_960_v16 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_961_i2 _dafny.Int) _dafny.Sequence {
								return _960_v16
							}
						})(_957_v16))), _958_v17)
						var _962_v19 _dafny.Map
						_ = _962_v19
						_962_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(222))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg44 _dafny.Int) interface{} {
								return coer44(arg44)
							}
						}(func(_963_i3 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('r')
						})))
						var _964_v20 _dafny.CodePoint
						_ = _964_v20
						_964_v20 = _dafny.CodePoint('c')
						var _965_v21 _dafny.MultiSet
						_ = _965_v21
						_965_v21 = _dafny.MultiSetOf(_956_v15)
						var _966_v22 _dafny.Sequence
						_ = _966_v22
						_966_v22 = _dafny.SeqOf(p0, _956_v15)
						var _967_v23 _dafny.Map
						_ = _967_v23
						_967_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_956_v15, _dafny.IntOfInt64(300))
						var _968_v24 _dafny.Map
						_ = _968_v24
						_968_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_967_v23).Cardinality(), _956_v15)
						var _969_v25 _dafny.Array
						_ = _969_v25
						var _nwElement0_31 _dafny.Int = p0
						_ = _nwElement0_31
						var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(22))
						_ = _nw165
						(_nw165).ArraySet1(_nwElement0_31, 0)
						(_nw165).ArraySet1(p0, 1)
						(_nw165).ArraySet1(_956_v15, 2)
						(_nw165).ArraySet1(_956_v15, 3)
						(_nw165).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm16(Companion_Default___.Fm0(Companion_Default___.Fm0(p2, _964_v20, false, globalState), _964_v20, p2, globalState), p0, globalState)).Cardinality()), 4)
						(_nw165).ArraySet1(Companion_Default___.Fm9(p0, _dafny.IntOfUint32((_957_v16).Cardinality()), p2, globalState), 5)
						(_nw165).ArraySet1(p0, 6)
						(_nw165).ArraySet1((_965_v21).Cardinality(), 7)
						(_nw165).ArraySet1(p0, 8)
						(_nw165).ArraySet1(_956_v15, 9)
						(_nw165).ArraySet1(_dafny.IntOfInt64(418), 10)
						(_nw165).ArraySet1(p0, 11)
						(_nw165).ArraySet1(_956_v15, 12)
						(_nw165).ArraySet1(p0, 13)
						(_nw165).ArraySet1(_dafny.IntOfUint32((_966_v22).Cardinality()), 14)
						(_nw165).ArraySet1(_956_v15, 15)
						(_nw165).ArraySet1(_dafny.IntOfInt64(594), 16)
						(_nw165).ArraySet1((_968_v24).Cardinality(), 17)
						(_nw165).ArraySet1(p0, 18)
						(_nw165).ArraySet1(_dafny.IntOfInt64(284), 19)
						(_nw165).ArraySet1(_956_v15, 20)
						(_nw165).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), p0)).Cardinality(), 21)
						_969_v25 = _nw165
						var _970_v26 _dafny.Sequence
						_ = _970_v26
						_970_v26 = _dafny.SeqOf(p1)
						var _971_v27 _dafny.Map
						_ = _971_v27
						_971_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(214))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg45 _dafny.Int) interface{} {
								return coer45(arg45)
							}
						}((func(_972_v20 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_973_i4 _dafny.Int) _dafny.CodePoint {
								return _972_v20
							}
						})(_964_v20)))).Cardinality()))
						var _974_v28 _dafny.Array
						_ = _974_v28
						var _nwElement0_32 _dafny.Int = _dafny.IntOfUint32((_957_v16).Cardinality())
						_ = _nwElement0_32
						var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(22))
						_ = _nw166
						(_nw166).ArraySet1(_nwElement0_32, 0)
						(_nw166).ArraySet1(_956_v15, 1)
						(_nw166).ArraySet1(p0, 2)
						(_nw166).ArraySet1((_dafny.Zero).Minus(p0), 3)
						(_nw166).ArraySet1(p0, 4)
						(_nw166).ArraySet1(p0, 5)
						(_nw166).ArraySet1(p0, 6)
						(_nw166).ArraySet1(_956_v15, 7)
						(_nw166).ArraySet1(p0, 8)
						(_nw166).ArraySet1(_dafny.IntOfUint32(((_959_v18).Select((Companion_Default___.SafeIndex(_956_v15, _dafny.IntOfUint32((_959_v18).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), 9)
						(_nw166).ArraySet1((_962_v19).Cardinality(), 10)
						(_nw166).ArraySet1(p0, 11)
						(_nw166).ArraySet1(_956_v15, 12)
						(_nw166).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 13)
						(_nw166).ArraySet1(_956_v15, 14)
						(_nw166).ArraySet1(Companion_Default___.Fm9(_dafny.IntOfInt64(83), _956_v15, (_this).F21(), globalState), 15)
						(_nw166).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_969_v25, _969_v25)).Cardinality()), 16)
						(_nw166).ArraySet1(_dafny.IntOfUint32(((_970_v26).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_970_v26).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), 17)
						(_nw166).ArraySet1(_956_v15, 18)
						(_nw166).ArraySet1(_dafny.IntOfUint32((_957_v16).Cardinality()), 19)
						(_nw166).ArraySet1((_971_v27).Cardinality(), 20)
						(_nw166).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 21)
						_974_v28 = _nw166
						_956_v15 = (Companion_D3_.Create_DC11_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _956_v15), _956_v15, _969_v25)).Dtor_cf19()
						var _975_v29 _dafny.Map
						_ = _975_v29
						_975_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), _956_v15)
						_956_v15 = _dafny.IntOfUint32((_dafny.SeqOf(_966_v22, _dafny.SeqOf(_956_v15, _956_v15), _dafny.SeqOf(_dafny.IntOfInt64(544), (_975_v29).Cardinality()))).Cardinality())
						(globalState).F15 = (_this).F21()
						_956_v15 = Companion_Default___.SafeDivisionInt((_966_v22).Select((Companion_Default___.SafeIndex(_956_v15, _dafny.IntOfUint32((_966_v22).Cardinality()))).Uint32()).(_dafny.Int), p0)
						var _976_v30 _dafny.Set
						_ = _976_v30
						_976_v30 = _dafny.SetOf((_this).F21())
						_956_v15 = ((_976_v30).Cardinality()).Times(_956_v15)
					}
					var _977_v31 _dafny.Array
					_ = _977_v31
					var _nw167 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(29))
					_ = _nw167
					_977_v31 = _nw167
					var _978_v32 _dafny.Map
					_ = _978_v32
					_978_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _977_v31)
					var _979_v33 _dafny.Map
					_ = _979_v33
					_979_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(p1, p1), _978_v32)
					_979_v33 = (_979_v33).Update(_dafny.UnicodeSeqOfUtf8Bytes("ubtj"), _978_v32)
					var _980_v34 _dafny.Array
					_ = _980_v34
					var _len0_29 _dafny.Int = _dafny.IntOfInt64(28)
					_ = _len0_29
					var _nw168 _dafny.Array
					_ = _nw168
					if _len0_29.Cmp(_dafny.Zero) == 0 {
						_nw168 = _dafny.NewArray(_len0_29)
					} else {
						var _init29 func(_dafny.Int) _dafny.Int = (func(_981_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_982_i5 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeModuloInt(_982_i5, _981_p0)
							}
						})(p0)
						_ = _init29
						var _element0_29 = _init29(_dafny.Zero)
						_ = _element0_29
						_nw168 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
						(_nw168).ArraySet1(_element0_29, 0)
						var _nativeLen0_29 = (_len0_29).Int()
						_ = _nativeLen0_29
						for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
							(_nw168).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
						}
					}
					_980_v34 = _nw168
					var _983_v35 _dafny.Sequence
					_ = _983_v35
					_983_v35 = _dafny.SeqOf(_980_v34)
					var _984_v36 *C1
					_ = _984_v36
					var _nw169 *C1 = New_C1_()
					_ = _nw169
					_nw169.Ctor__(_980_v34)
					_984_v36 = _nw169
					var _985_v37 D5
					_ = _985_v37
					_985_v37 = Companion_D5_.Create_DC16_(_984_v36)
					var _986_v38 _dafny.Sequence
					_ = _986_v38
					_986_v38 = _dafny.SeqOf((_985_v37).Dtor_cf28())
					var _987_v39 *C1
					_ = _987_v39
					var _nw170 *C1 = New_C1_()
					_ = _nw170
					_nw170.Ctor__((_983_v35).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_986_v38).Cardinality()), _dafny.IntOfUint32((_983_v35).Cardinality()))).Uint32()).(_dafny.Array))
					_987_v39 = _nw170
					var _988_v40 _dafny.Map
					_ = _988_v40
					_988_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _980_v34)
					_988_v40 = (_988_v40).Update(((_this).F21()) && (p2), (_987_v39).F22())
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _989_i6 _dafny.Int
		_ = _989_i6
		_989_i6 = _dafny.Zero
		{
			for p2 {
				{
					if (_989_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_989_i6 = (_989_i6).Plus(_dafny.One)
					var _990_v41 _dafny.Array
					_ = _990_v41
					var _nw171 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
					_ = _nw171
					_990_v41 = _nw171
					var _991_v42 D1
					_ = _991_v42
					_991_v42 = Companion_D1_.Create_DC4_(_dafny.IntOfInt64(451), !((_this).F21()), _990_v41)
					var _source12 D1 = _991_v42
					_ = _source12
					if _source12.Is_DC4() {
						var _992___mcc_h0 _dafny.Int = _source12.Get_().(D1_DC4).Cf4
						_ = _992___mcc_h0
						var _993___mcc_h1 bool = _source12.Get_().(D1_DC4).Cf5
						_ = _993___mcc_h1
						var _994___mcc_h2 _dafny.Array = _source12.Get_().(D1_DC4).Cf6
						_ = _994___mcc_h2
						var _995_cf6 _dafny.Array = _994___mcc_h2
						_ = _995_cf6
						var _996_cf5 bool = _993___mcc_h1
						_ = _996_cf5
						var _997_cf4 _dafny.Int = _992___mcc_h0
						_ = _997_cf4
						r2 = p2
						var _998_v43 *C0
						_ = _998_v43
						var _nw172 *C0 = New_C0_()
						_ = _nw172
						_nw172.Ctor__(p2)
						_998_v43 = _nw172
						var _999_v44 *C0
						_ = _999_v44
						var _nw173 *C0 = New_C0_()
						_ = _nw173
						_nw173.Ctor__(!((_this).F21()))
						_999_v44 = _nw173
						var _1000_v45 _dafny.CodePoint
						_ = _1000_v45
						_1000_v45 = _dafny.CodePoint('b')
						var _1001_v46 _dafny.Map
						_ = _1001_v46
						_1001_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
							if (_999_v44).F23() {
								return p1
							}
							return _dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_997_cf4), _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _1000_v45)
						})(), (_dafny.SetOf((_this).F21(), _996_cf5, (_999_v44).F23(), (_this).F21())).Cardinality())
						var _1002_v47 _dafny.Map
						_ = _1002_v47
						_1002_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_998_v43).F23())
						_1001_v46 = Companion_Default___.Fm25((_998_v43).F23(), (_1002_v47).Equals(_1002_v47), globalState)
					} else if _source12.Is_DC5() {
						var _1003___mcc_h3 _dafny.Sequence = _source12.Get_().(D1_DC5).Cf7
						_ = _1003___mcc_h3
						var _1004___mcc_h4 _dafny.Int = _source12.Get_().(D1_DC5).Cf8
						_ = _1004___mcc_h4
						var _1005___mcc_h5 _dafny.Int = _source12.Get_().(D1_DC5).Cf9
						_ = _1005___mcc_h5
						var _1006_cf9 _dafny.Int = _1005___mcc_h5
						_ = _1006_cf9
						var _1007_cf8 _dafny.Int = _1004___mcc_h4
						_ = _1007_cf8
						var _1008_cf7 _dafny.Sequence = _1003___mcc_h3
						_ = _1008_cf7
						var _1009_v48 _dafny.Sequence
						_ = _1009_v48
						_1009_v48 = _dafny.SeqOf(_dafny.IntOfUint32((p1).Cardinality()))
						var _1010_v49 _dafny.Array
						_ = _1010_v49
						var _nwElement0_33 _dafny.Int = _1007_cf8
						_ = _nwElement0_33
						var _nw174 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(17))
						_ = _nw174
						(_nw174).ArraySet1(_nwElement0_33, 0)
						(_nw174).ArraySet1((p0).Times(_1006_cf9), 1)
						(_nw174).ArraySet1(_1007_cf8, 2)
						(_nw174).ArraySet1(Companion_Default___.SafeDivisionInt(_1006_cf9, p0), 3)
						(_nw174).ArraySet1(_1006_cf9, 4)
						(_nw174).ArraySet1(p0, 5)
						(_nw174).ArraySet1(p0, 6)
						(_nw174).ArraySet1(_1007_cf8, 7)
						(_nw174).ArraySet1(_dafny.IntOfInt64(723), 8)
						(_nw174).ArraySet1(_dafny.IntOfInt64(785), 9)
						(_nw174).ArraySet1(_1006_cf9, 10)
						(_nw174).ArraySet1(p0, 11)
						(_nw174).ArraySet1((_dafny.SetOf((_1009_v48).Select((Companion_Default___.SafeIndex(_1007_cf8, _dafny.IntOfUint32((_1009_v48).Cardinality()))).Uint32()).(_dafny.Int), _1007_cf8, _1006_cf9, p0)).Cardinality(), 12)
						(_nw174).ArraySet1(_1007_cf8, 13)
						(_nw174).ArraySet1((_1009_v48).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1009_v48).Cardinality()))).Uint32()).(_dafny.Int), 14)
						(_nw174).ArraySet1(p0, 15)
						(_nw174).ArraySet1(Companion_Default___.SafeModuloInt(p0, p0), 16)
						_1010_v49 = _nw174
						var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_1010_v49), 0))
						_ = _index126
						(_1010_v49).ArraySet1(_1006_cf9, (_index126).Int())
						var _1011_v50 _dafny.Map
						_ = _1011_v50
						_1011_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1010_v49, _1010_v49)
						var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_1010_v49), 0))
						_ = _index127
						(_1010_v49).ArraySet1((_1011_v50).Cardinality(), (_index127).Int())
						(globalState).F11 = (_this).F21()
						var _1012_v51 _dafny.Sequence
						_ = _1012_v51
						_1012_v51 = _dafny.SeqOf((_this).F21())
						var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_990_v41), 0))
						_ = _index128
						(_990_v41).ArraySet1((_1012_v51).Select((Companion_Default___.SafeIndex((_dafny.SetOf((_this).F21(), (_this).F21(), (_1012_v51).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xlt")).Cardinality()), _dafny.IntOfUint32((_1012_v51).Cardinality()))).Uint32()).(bool))).Cardinality(), _dafny.IntOfUint32((_1012_v51).Cardinality()))).Uint32()).(bool), (_index128).Int())
						var _1013_v52 _dafny.Array
						_ = _1013_v52
						var _nw175 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(3))
						_ = _nw175
						_1013_v52 = _nw175
						var _1014_v53 _dafny.CodePoint
						_ = _1014_v53
						_1014_v53 = _dafny.CodePoint('h')
						var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_1013_v52), 0))
						_ = _index129
						(_1013_v52).ArraySet1CodePoint(_1014_v53, (_index129).Int())
						var _1015_v54 _dafny.Set
						_ = _1015_v54
						_1015_v54 = _dafny.SetOf(!((_this).F21()))
						var _1016_v55 T0
						_ = _1016_v55
						var _nw176 *C0 = New_C0_()
						_ = _nw176
						_nw176.Ctor__((_1015_v54).IsSubsetOf(_1015_v54))
						_1016_v55 = _nw176
						var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_990_v41), 0))
						_ = _index130
						var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_1013_v52), 0))
						_ = _index131
						var _rhs116 bool = p2
						_ = _rhs116
						var _rhs117 _dafny.Int = _1006_cf9
						_ = _rhs117
						var _rhs118 bool = (_1015_v54).IsSubsetOf((_dafny.SetOf((_this).F21())).Difference(_1015_v54))
						_ = _rhs118
						var _rhs119 _dafny.CodePoint = _1014_v53
						_ = _rhs119
						var _rhs120 T0 = _1016_v55
						_ = _rhs120
						var _lhs60 _dafny.Array = _990_v41
						_ = _lhs60
						var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_990_v41), 0))
						_ = _lhs61
						var _lhs62 _dafny.Array = _1013_v52
						_ = _lhs62
						var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_1013_v52), 0))
						_ = _lhs63
						(_lhs60).ArraySet1(_rhs116, (_lhs61).Int())
						_1006_cf9 = _rhs117
						r0 = _rhs118
						(_lhs62).ArraySet1CodePoint(_rhs119, (_lhs63).Int())
						_1016_v55 = _rhs120
						var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_1010_v49), 0))
						_ = _index132
						(_1010_v49).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_1006_cf9), _1007_cf8), (_index132).Int())
					} else if _source12.Is_DC6() {
						var _1017___mcc_h6 _dafny.Int = _source12.Get_().(D1_DC6).Cf10
						_ = _1017___mcc_h6
						var _1018___mcc_h7 _dafny.Int = _source12.Get_().(D1_DC6).Cf11
						_ = _1018___mcc_h7
						var _1019_cf11 _dafny.Int = _1018___mcc_h7
						_ = _1019_cf11
						var _1020_cf10 _dafny.Int = _1017___mcc_h6
						_ = _1020_cf10
						_1020_cf10 = _1019_cf11
						var _1021_v56 _dafny.MultiSet
						_ = _1021_v56
						_1021_v56 = _dafny.MultiSetOf(p0, p0)
						var _1022_v57 _dafny.Map
						_ = _1022_v57
						_1022_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1021_v56, _1020_cf10)
						r1 = (_1022_v57).Equals(_1022_v57)
						var _1023_v58 _dafny.Sequence
						_ = _1023_v58
						_1023_v58 = _dafny.SeqOf(_990_v41, _990_v41, _990_v41, _990_v41, _990_v41)
						r2 = _dafny.Companion_Sequence_.Contains(_1023_v58, _990_v41)
						r3 = p1
					} else {
						var _1024___mcc_h8 _dafny.Array = _source12.Get_().(D1_DC3).Cf3
						_ = _1024___mcc_h8
						var _1025_cf3 _dafny.Array = _1024___mcc_h8
						_ = _1025_cf3
						var _1026_v59 _dafny.Map
						_ = _1026_v59
						_1026_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
						var _1027_v60 _dafny.MultiSet
						_ = _1027_v60
						_1027_v60 = _dafny.MultiSetOf(_1026_v59)
						var _1028_v61 D6
						_ = _1028_v61
						_1028_v61 = Companion_D6_.Create_DC19_(_dafny.MultiSetOf(_1026_v59, _1026_v59))
						_1027_v60 = (_1027_v60).Difference((_1028_v61).Dtor_cf35())
						var _1029_v62 _dafny.Array
						_ = _1029_v62
						var _nw177 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
						_ = _nw177
						_1029_v62 = _nw177
						var _1030_v63 _dafny.Set
						_ = _1030_v63
						_1030_v63 = _dafny.SetOf(_dafny.IntOfUint32((p1).Cardinality()), p0, p0)
						var _1031_v64 _dafny.Sequence
						_ = _1031_v64
						_1031_v64 = _dafny.SeqOf((_1030_v63).Cardinality(), p0)
						var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_1029_v62), 0))
						_ = _index133
						(_1029_v62).ArraySet1(_1031_v64, (_index133).Int())
						var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_1029_v62), 0))
						_ = _index134
						(_1029_v62).ArraySet1(_1031_v64, (_index134).Int())
						var _1032_v65 _dafny.Array
						_ = _1032_v65
						var _nw178 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
						_ = _nw178
						_1032_v65 = _nw178
						var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_1032_v65), 0))
						_ = _index135
						(_1032_v65).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, Companion_Default___.Fm16((_this).F21(), p0, globalState))).Cardinality()), (_index135).Int())
						var _1033_v66 D6
						_ = _1033_v66
						_1033_v66 = Companion_D6_.Create_DC20_((_this).F21(), p0, true)
						var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_1032_v65), 0))
						_ = _index136
						var _rhs121 bool = p2
						_ = _rhs121
						var _rhs122 bool = (_1033_v66).Dtor_cf36()
						_ = _rhs122
						var _rhs123 _dafny.Int = (_dafny.Zero).Minus(p0)
						_ = _rhs123
						var _lhs64 *GlobalState = globalState
						_ = _lhs64
						var _lhs65 *GlobalState = globalState
						_ = _lhs65
						var _lhs66 _dafny.Array = _1032_v65
						_ = _lhs66
						var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_1032_v65), 0))
						_ = _lhs67
						_lhs64.F17 = _rhs121
						_lhs65.F17 = _rhs122
						(_lhs66).ArraySet1(_rhs123, (_lhs67).Int())
						var _1034_v67 _dafny.Sequence
						_ = _1034_v67
						_1034_v67 = _dafny.SeqOf((func() bool {
							if (_this).F21() {
								return p2
							}
							return (_this).F21()
						})(), (_this).F21(), (_this).F21(), p2)
						(globalState).F16 = _dafny.Companion_Sequence_.Update(_1034_v67, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1034_v67).Cardinality()))).Uint32(), (_this).F21())
					}
					var _1035_v68 _dafny.Set
					_ = _1035_v68
					_1035_v68 = _dafny.SetOf(!((_this).F21()), p2, p2)
					var _1036_v69 _dafny.Sequence
					_ = _1036_v69
					_1036_v69 = _dafny.SeqOf(Companion_Default___.Fm20(globalState))
					var _1037_v70 _dafny.Int
					_ = _1037_v70
					_1037_v70 = _dafny.IntOfInt64(616)
					var _rhs124 _dafny.Set = (_1035_v68).Intersection(_dafny.SetOf(true))
					_ = _rhs124
					var _rhs125 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p1, p1)
					_ = _rhs125
					var _rhs126 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1036_v69, _dafny.SeqOf(_1035_v68, _1035_v68))
					_ = _rhs126
					var _rhs127 _dafny.Int = _dafny.IntOfInt64(396)
					_ = _rhs127
					_1035_v68 = _rhs124
					r3 = _rhs125
					_1036_v69 = _rhs126
					_1037_v70 = _rhs127
					var _1038_v71 _dafny.CodePoint
					_ = _1038_v71
					_1038_v71 = _dafny.CodePoint('w')
					var _1039_v72 _dafny.Map
					_ = _1039_v72
					_1039_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1037_v70, Companion_Default___.Fm0((_this).F21(), _1038_v71, p2, globalState))
					var _1040_v73 _dafny.MultiSet
					_ = _1040_v73
					_1040_v73 = _dafny.MultiSetOf((_this).F21())
					var _1041_v74 _dafny.Array
					_ = _1041_v74
					var _nwElement0_34 _dafny.Int = p0
					_ = _nwElement0_34
					var _nw179 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(27))
					_ = _nw179
					(_nw179).ArraySet1(_nwElement0_34, 0)
					(_nw179).ArraySet1(p0, 1)
					(_nw179).ArraySet1(_1037_v70, 2)
					(_nw179).ArraySet1(p0, 3)
					(_nw179).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 4)
					(_nw179).ArraySet1((_dafny.Zero).Minus(p0), 5)
					(_nw179).ArraySet1(p0, 6)
					(_nw179).ArraySet1(_1037_v70, 7)
					(_nw179).ArraySet1(_dafny.IntOfInt64(443), 8)
					(_nw179).ArraySet1(_dafny.IntOfInt64(-109), 9)
					(_nw179).ArraySet1(_1037_v70, 10)
					(_nw179).ArraySet1(_dafny.IntOfInt64(882), 11)
					(_nw179).ArraySet1(_dafny.IntOfInt64(55), 12)
					(_nw179).ArraySet1(p0, 13)
					(_nw179).ArraySet1(_1037_v70, 14)
					(_nw179).ArraySet1(_1037_v70, 15)
					(_nw179).ArraySet1(Companion_Default___.Fm9(p0, (_1035_v68).Cardinality(), (_this).F21(), globalState), 16)
					(_nw179).ArraySet1(_1037_v70, 17)
					(_nw179).ArraySet1((_1039_v72).Cardinality(), 18)
					(_nw179).ArraySet1(p0, 19)
					(_nw179).ArraySet1(_1037_v70, 20)
					(_nw179).ArraySet1(_1037_v70, 21)
					(_nw179).ArraySet1((_1040_v73).Cardinality(), 22)
					(_nw179).ArraySet1(p0, 23)
					(_nw179).ArraySet1(p0, 24)
					(_nw179).ArraySet1(p0, 25)
					(_nw179).ArraySet1(_1037_v70, 26)
					_1041_v74 = _nw179
					var _1042_v75 *C1
					_ = _1042_v75
					var _nw180 *C1 = New_C1_()
					_ = _nw180
					_nw180.Ctor__(_1041_v74)
					_1042_v75 = _nw180
					var _1043_v76 T0
					_ = _1043_v76
					var _nw181 *C0 = New_C0_()
					_ = _nw181
					_nw181.Ctor__((_this).F21())
					_1043_v76 = _nw181
					var _1044_v77 D4
					_ = _1044_v77
					_1044_v77 = Companion_D4_.Create_DC12_(_1043_v76)
					var _1045_v78 _dafny.Map
					_ = _1045_v78
					_1045_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1040_v73).Cardinality(), _1043_v76)
					var _1046_v79 D2
					_ = _1046_v79
					_1046_v79 = Companion_D2_.Create_DC8_(_990_v41, p0, p0)
					var _1047_v80 _dafny.Map
					_ = _1047_v80
					_1047_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1046_v79, _1043_v76)
					var _1048_v81 _dafny.Sequence
					_ = _1048_v81
					_1048_v81 = _dafny.SeqOf(_1037_v70)
					var _1049_v82 _dafny.Map
					_ = _1049_v82
					_1049_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _1048_v81)
					var _1050_v83 _dafny.Array
					_ = _1050_v83
					var _nwElement0_35 T0 = _1043_v76
					_ = _nwElement0_35
					var _nw182 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(25))
					_ = _nw182
					(_nw182).ArraySet1(_nwElement0_35, 0)
					(_nw182).ArraySet1(_1043_v76, 1)
					(_nw182).ArraySet1((_1044_v77).Dtor_cf21(), 2)
					(_nw182).ArraySet1(_1043_v76, 3)
					(_nw182).ArraySet1(_1043_v76, 4)
					(_nw182).ArraySet1((_1044_v77).Dtor_cf21(), 5)
					(_nw182).ArraySet1(_1043_v76, 6)
					(_nw182).ArraySet1(_1043_v76, 7)
					(_nw182).ArraySet1(_1043_v76, 8)
					(_nw182).ArraySet1((func() T0 {
						if (_1045_v78).Contains(p0) {
							return (_1045_v78).Get(p0).(T0)
						}
						return _1043_v76
					})(), 9)
					(_nw182).ArraySet1(_1043_v76, 10)
					(_nw182).ArraySet1(_1043_v76, 11)
					(_nw182).ArraySet1(_1043_v76, 12)
					(_nw182).ArraySet1(_1043_v76, 13)
					(_nw182).ArraySet1(_1043_v76, 14)
					(_nw182).ArraySet1((_1044_v77).Dtor_cf21(), 15)
					(_nw182).ArraySet1(_1043_v76, 16)
					(_nw182).ArraySet1(_1043_v76, 17)
					(_nw182).ArraySet1(_1043_v76, 18)
					(_nw182).ArraySet1(_1043_v76, 19)
					(_nw182).ArraySet1(_1043_v76, 20)
					(_nw182).ArraySet1(_1043_v76, 21)
					(_nw182).ArraySet1(_1043_v76, 22)
					(_nw182).ArraySet1(_1043_v76, 23)
					(_nw182).ArraySet1((func() T0 {
						if (_1047_v80).Contains(Companion_D2_.Create_DC8_(_990_v41, (_1049_v82).Cardinality(), (Companion_Default___.Fm26(false, _dafny.CodePoint('m'), (_this).F21(), p2, globalState)).Dtor_cf37())) {
							return (_1047_v80).Get(Companion_D2_.Create_DC8_(_990_v41, (_1049_v82).Cardinality(), (Companion_Default___.Fm26(false, _dafny.CodePoint('m'), (_this).F21(), p2, globalState)).Dtor_cf37())).(T0)
						}
						return _1043_v76
					})(), 24)
					_1050_v83 = _nw182
					var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_1050_v83), 0))
					_ = _index137
					(_1050_v83).ArraySet1(_1043_v76, (_index137).Int())
					var _1051_v84 _dafny.Sequence
					_ = _1051_v84
					_1051_v84 = _dafny.SeqOf(p2)
					var _1052_v85 _dafny.MultiSet
					_ = _1052_v85
					_1052_v85 = _dafny.MultiSetOf(p0, _dafny.IntOfInt64(60), p0)
					var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_1050_v83), 0))
					_ = _index138
					var _rhs128 *C1 = _1042_v75
					_ = _rhs128
					var _rhs129 bool = (_1037_v70).Cmp(_1037_v70) > 0
					_ = _rhs129
					var _rhs130 T0 = _1043_v76
					_ = _rhs130
					var _rhs131 bool = !((func() bool {
						if true {
							return (_dafny.MultiSetOf(p0, _dafny.IntOfInt64(115), _dafny.IntOfUint32((_1051_v84).Cardinality()))).IsSubsetOf(_1052_v85)
						}
						return (_this).F21()
					})())
					_ = _rhs131
					var _lhs68 *GlobalState = globalState
					_ = _lhs68
					var _lhs69 _dafny.Array = _1050_v83
					_ = _lhs69
					var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_1050_v83), 0))
					_ = _lhs70
					var _lhs71 *GlobalState = globalState
					_ = _lhs71
					_1042_v75 = _rhs128
					_lhs68.F11 = _rhs129
					(_lhs69).ArraySet1(_rhs130, (_lhs70).Int())
					_lhs71.F11 = _rhs131
					var _1053_v86 _dafny.Array
					_ = _1053_v86
					var _nwElement0_36 _dafny.CodePoint = _1038_v71
					_ = _nwElement0_36
					var _nw183 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(2))
					_ = _nw183
					(_nw183).ArraySet1CodePoint(_nwElement0_36, 0)
					(_nw183).ArraySet1CodePoint(_1038_v71, 1)
					_1053_v86 = _nw183
					_1053_v86 = _1053_v86
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		if !(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(p1, p1), p1)) {
			var _1054_v87 _dafny.Set
			_ = _1054_v87
			_1054_v87 = _dafny.SetOf(p0, p0)
			var _1055_v88 _dafny.Map
			_ = _1055_v88
			_1055_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _1054_v87)
			var _1056_v89 _dafny.Sequence
			_ = _1056_v89
			_1056_v89 = _dafny.SeqOf((_this).F21())
			var _1057_v90 bool
			_ = _1057_v90
			var _out53 bool
			_ = _out53
			_out53 = (_this).M5((_1054_v87).Equals((func() _dafny.Set {
				if (_1055_v88).Contains(false) {
					return (_1055_v88).Get(false).(_dafny.Set)
				}
				return _1054_v87
			})()), _1056_v89, globalState)
			_1057_v90 = _out53
			(globalState).F11 = true
			var _1058_v91 _dafny.Array
			_ = _1058_v91
			var _len0_30 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_30
			var _nw184 _dafny.Array
			_ = _nw184
			if _len0_30.Cmp(_dafny.Zero) == 0 {
				_nw184 = _dafny.NewArray(_len0_30)
			} else {
				var _init30 func(_dafny.Int) _dafny.Int = func(_1059_i7 _dafny.Int) _dafny.Int {
					return (_1059_i7).Times(_dafny.IntOfInt64(-579))
				}
				_ = _init30
				var _element0_30 = _init30(_dafny.Zero)
				_ = _element0_30
				_nw184 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
				(_nw184).ArraySet1(_element0_30, 0)
				var _nativeLen0_30 = (_len0_30).Int()
				_ = _nativeLen0_30
				for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
					(_nw184).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
				}
			}
			_1058_v91 = _nw184
			var _1060_v92 *C1
			_ = _1060_v92
			var _nw185 *C1 = New_C1_()
			_ = _nw185
			_nw185.Ctor__(_1058_v91)
			_1060_v92 = _nw185
			var _1061_v93 _dafny.MultiSet
			_ = _1061_v93
			_1061_v93 = _dafny.MultiSetOf(p0, p0, p0)
			var _1062_v94 _dafny.Map
			_ = _1062_v94
			_1062_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
			var _1063_v95 _dafny.CodePoint
			_ = _1063_v95
			_1063_v95 = _dafny.CodePoint('h')
			_1061_v93 = (_1061_v93).Intersection(Companion_Default___.Fm23(_1057_v90, (_1062_v94).Update(p2, p0), p0, _1063_v95, globalState))
			var _1064_v96 _dafny.Map
			_ = _1064_v96
			_1064_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_1061_v93).Contains(p0) {
					return (_1061_v93).Multiplicity(p0)
				}
				return p0
			})(), p0)
			var _1065_v97 _dafny.Sequence
			_ = _1065_v97
			_1065_v97 = _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(Companion_Default___.Fm16(_1057_v90, p0, globalState))).Cardinality())), p0)
			_1064_v96 = (_1064_v96).Update((_1065_v97).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1065_v97).Cardinality()))).Uint32()).(_dafny.Int), Companion_Default___.SafeModuloInt(p0, p0))
		} else {
			var _1066_v98 _dafny.Array
			_ = _1066_v98
			var _nw186 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw186
			_1066_v98 = _nw186
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_1066_v98), 0))
			_ = _index139
			(_1066_v98).ArraySet1(p0, (_index139).Int())
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_1066_v98), 0))
			_ = _index140
			(_1066_v98).ArraySet1(p0, (_index140).Int())
			(globalState).F15 = (_this).F21()
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_1066_v98), 0))
			_ = _index141
			(_1066_v98).ArraySet1(((_dafny.Zero).Minus(p0)).Plus((_1066_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_1066_v98), 0))).Int()).(_dafny.Int)), (_index141).Int())
			var _1067_v99 D6
			_ = _1067_v99
			_1067_v99 = Companion_D6_.Create_DC20_((_this).F21(), (_1066_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_1066_v98), 0))).Int()).(_dafny.Int), (_this).F21())
			(globalState).F17 = (_1067_v99).Dtor_cf36()
			var _1068_v100 _dafny.Array
			_ = _1068_v100
			var _nw187 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw187
			_1068_v100 = _nw187
			var _1069_v101 _dafny.Sequence
			_ = _1069_v101
			_1069_v101 = _dafny.SeqOf(_1068_v100, _1068_v100)
			var _1070_v102 _dafny.Set
			_ = _1070_v102
			_1070_v102 = _dafny.SetOf(_1068_v100, _1068_v100)
			var _1071_v103 _dafny.Map
			_ = _1071_v103
			_1071_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
			var _1072_v104 _dafny.Map
			_ = _1072_v104
			_1072_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1071_v103).Cardinality(), (_1066_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_1066_v98), 0))).Int()).(_dafny.Int))
			var _1073_v106 _dafny.CodePoint
			_ = _1073_v106
			_1073_v106 = _dafny.CodePoint('u')
			var _1074_v107 D4
			_ = _1074_v107
			_1074_v107 = Companion_D4_.Create_DC14_(false, p0, func() _dafny.Set {
				var _coll32 = _dafny.NewBuilder()
				_ = _coll32
				for _iter34 := _dafny.Iterate((_1072_v104).Keys().Elements()); ; {
					_compr_32, _ok34 := _iter34()
					if !_ok34 {
						break
					}
					var _1075_v105 _dafny.Int
					_1075_v105 = interface{}(_compr_32).(_dafny.Int)
					if (_1072_v104).Contains(_1075_v105) {
						_coll32.Add((_1075_v105).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lp")).Cardinality())))
					}
				}
				return _coll32.ToSet()
			}(), _1073_v106)
			var _1076_v108 _dafny.Map
			_ = _1076_v108
			_1076_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1070_v102).IsSubsetOf(_dafny.SetOf((_1069_v101).Select((Companion_Default___.SafeIndex((_1066_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_1066_v98), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1069_v101).Cardinality()))).Uint32()).(_dafny.Array))), Companion_Default___.Fm12((_1074_v107).Dtor_cf24(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-558))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg46 _dafny.Int) interface{} {
					return coer46(arg46)
				}
			}(func(_1077_i8 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(464))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}((func(_1078_v106 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1079_i9 _dafny.Int) _dafny.CodePoint {
					return _1078_v106
				}
			})(_1073_v106))), globalState))
			_1076_v108 = (_1076_v108).Update((_this).F21(), _1073_v106)
		}
		var _1080_i10 _dafny.Int
		_ = _1080_i10
		_1080_i10 = _dafny.Zero
		{
			for (p0).Cmp(p0) != 0 {
				{
					if (_1080_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1080_i10 = (_1080_i10).Plus(_dafny.One)
					var _1081_v109 _dafny.Map
					_ = _1081_v109
					_1081_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F21())
					var _1082_v110 _dafny.Set
					_ = _1082_v110
					_1082_v110 = _dafny.SetOf(_1081_v109)
					var _1083_v111 D3
					_ = _1083_v111
					_1083_v111 = Companion_D3_.Create_DC9_(_1082_v110)
					_1083_v111 = _1083_v111
					var _1084_v112 _dafny.Array
					_ = _1084_v112
					var _nw188 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
					_ = _nw188
					_1084_v112 = _nw188
					var _1085_v113 _dafny.CodePoint
					_ = _1085_v113
					_1085_v113 = _dafny.CodePoint('i')
					var _1086_v114 _dafny.Map
					_ = _1086_v114
					_1086_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p1).Cardinality()), p1)
					var _1087_v115 _dafny.Sequence
					_ = _1087_v115
					_1087_v115 = _dafny.SeqOf(p0, _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))
					var _1088_v116 D3
					_ = _1088_v116
					_1088_v116 = Companion_D3_.Create_DC10_(p0)
					var _1089_v117 _dafny.Map
					_ = _1089_v117
					_1089_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1087_v115).Cardinality()), _1088_v116)
					var _1090_v118 _dafny.Array
					_ = _1090_v118
					var _nwElement0_37 _dafny.Sequence = p1
					_ = _nwElement0_37
					var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(21))
					_ = _nw189
					(_nw189).ArraySet1(_nwElement0_37, 0)
					(_nw189).ArraySet1(p1, 1)
					(_nw189).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("snqcwt"), 2)
					(_nw189).ArraySet1(p1, 3)
					(_nw189).ArraySet1(p1, 4)
					(_nw189).ArraySet1(p1, 5)
					(_nw189).ArraySet1(p1, 6)
					(_nw189).ArraySet1(p1, 7)
					(_nw189).ArraySet1(p1, 8)
					(_nw189).ArraySet1(p1, 9)
					(_nw189).ArraySet1(p1, 10)
					(_nw189).ArraySet1(p1, 11)
					(_nw189).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cx"), 12)
					(_nw189).ArraySet1(_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _1085_v113), 13)
					(_nw189).ArraySet1(p1, 14)
					(_nw189).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ygblryv"), 15)
					(_nw189).ArraySet1((func() _dafny.Sequence {
						if (_1086_v114).Contains((_1089_v117).Cardinality()) {
							return (_1086_v114).Get((_1089_v117).Cardinality()).(_dafny.Sequence)
						}
						return p1
					})(), 16)
					(_nw189).ArraySet1(p1, 17)
					(_nw189).ArraySet1(p1, 18)
					(_nw189).ArraySet1(p1, 19)
					(_nw189).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-367))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg48 _dafny.Int) interface{} {
							return coer48(arg48)
						}
					}((func(_1091_v113 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1092_i11 _dafny.Int) _dafny.CodePoint {
							return _1091_v113
						}
					})(_1085_v113))), 20)
					_1090_v118 = _nw189
					var _1093_v119 _dafny.Array
					_ = _1093_v119
					_1093_v119 = _1090_v118
					var _1094_v120 _dafny.Sequence
					_ = _1094_v120
					_1094_v120 = _dafny.SeqOf((_1090_v118))
					var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_1084_v112), 0))
					_ = _index142
					(_1084_v112).ArraySet1((_1094_v120).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(215), _dafny.IntOfUint32((_1094_v120).Cardinality()))).Uint32()).(_dafny.Array), (_index142).Int())
					var _1095_v121 _dafny.MultiSet
					_ = _1095_v121
					_1095_v121 = _dafny.MultiSetOf(p2)
					var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_1084_v112), 0))
					_ = _index143
					var _rhs132 _dafny.Array = _1090_v118
					_ = _rhs132
					var _rhs133 bool = (_dafny.IntOfInt64(419)).Cmp(Companion_Default___.SafeModuloInt(p0, p0)) >= 0
					_ = _rhs133
					var _rhs134 _dafny.MultiSet = (_1095_v121).Difference((_dafny.MultiSetOf(p2)).Update(false, Companion_Default___.Abs(p0)))
					_ = _rhs134
					var _lhs72 _dafny.Array = _1084_v112
					_ = _lhs72
					var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_1084_v112), 0))
					_ = _lhs73
					var _lhs74 *GlobalState = globalState
					_ = _lhs74
					(_lhs72).ArraySet1(_rhs132, (_lhs73).Int())
					_lhs74.F15 = _rhs133
					_1095_v121 = _rhs134
					var _1096_v122 _dafny.Sequence
					_ = _1096_v122
					_1096_v122 = _dafny.SeqOf(p2)
					var _1097_v123 _dafny.Array
					_ = _1097_v123
					var _nwElement0_38 _dafny.Sequence = _1096_v122
					_ = _nwElement0_38
					var _nw190 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(13))
					_ = _nw190
					(_nw190).ArraySet1(_nwElement0_38, 0)
					(_nw190).ArraySet1(_dafny.SeqOf(p2, p2), 1)
					(_nw190).ArraySet1(_1096_v122, 2)
					(_nw190).ArraySet1((func() _dafny.Sequence {
						if p2 {
							return _1096_v122
						}
						return _1096_v122
					})(), 3)
					(_nw190).ArraySet1(_1096_v122, 4)
					(_nw190).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1096_v122, _dafny.SeqOf(true, (_this).F21())), 5)
					(_nw190).ArraySet1(_1096_v122, 6)
					(_nw190).ArraySet1(_1096_v122, 7)
					(_nw190).ArraySet1(_1096_v122, 8)
					(_nw190).ArraySet1(_1096_v122, 9)
					(_nw190).ArraySet1(_1096_v122, 10)
					(_nw190).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1096_v122, _1096_v122), 11)
					(_nw190).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1096_v122, _1096_v122), 12)
					_1097_v123 = _nw190
					var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_1097_v123), 0))
					_ = _index144
					(_1097_v123).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1096_v122, Companion_Default___.Fm27(p0, p0, (_dafny.Zero).Minus(p0), globalState)), (_index144).Int())
					var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_1097_v123), 0))
					_ = _index145
					(_1097_v123).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p2), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality()))).Uint32(), p2), (_index145).Int())
					var _1098_v124 _dafny.Array
					_ = _1098_v124
					var _len0_31 _dafny.Int = _dafny.IntOfInt64(7)
					_ = _len0_31
					var _nw191 _dafny.Array
					_ = _nw191
					if _len0_31.Cmp(_dafny.Zero) == 0 {
						_nw191 = _dafny.NewArray(_len0_31)
					} else {
						var _init31 func(_dafny.Int) _dafny.Int = (func(_1099_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1100_i12 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeModuloInt(_1100_i12, _1099_p0)
							}
						})(p0)
						_ = _init31
						var _element0_31 = _init31(_dafny.Zero)
						_ = _element0_31
						_nw191 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
						(_nw191).ArraySet1(_element0_31, 0)
						var _nativeLen0_31 = (_len0_31).Int()
						_ = _nativeLen0_31
						for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
							(_nw191).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
						}
					}
					_1098_v124 = _nw191
					var _1101_v125 _dafny.Sequence
					_ = _1101_v125
					_1101_v125 = _dafny.SeqOf(_1098_v124)
					_1101_v125 = _dafny.SeqOf(_1098_v124)
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _hi7 _dafny.Int = p0
		_ = _hi7
		for _1102_i13 := p0; _1102_i13.Cmp(_hi7) < 0; _1102_i13 = _1102_i13.Plus(_dafny.One) {
			var _1103_v126 _dafny.Array
			_ = _1103_v126
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_32
			var _nw192 _dafny.Array
			_ = _nw192
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw192 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) _dafny.Sequence = func(_1104_i14 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_this).F21(), (_this).F21())
				}
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw192 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw192).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw192).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1103_v126 = _nw192
			var _1105_v127 _dafny.Sequence
			_ = _1105_v127
			_1105_v127 = _dafny.SeqOf((_this).F21())
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1103_v126), 0))
			_ = _index146
			(_1103_v126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1105_v127, _1105_v127), (_index146).Int())
			var _1106_v128 _dafny.Array
			_ = _1106_v128
			var _nw193 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
			_ = _nw193
			_1106_v128 = _nw193
			var _1107_v129 _dafny.Int
			_ = _1107_v129
			_1107_v129 = _dafny.IntOfInt64(808)
			var _1108_v130 D8
			_ = _1108_v130
			_1108_v130 = Companion_D8_.Create_DC22_(_1106_v128)
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1103_v126), 0))
			_ = _index147
			var _rhs135 bool = !((Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-12), (_dafny.Zero).Minus(_1107_v129))).Cmp(Companion_Default___.Fm9(p0, p0, false, globalState)) >= 0)
			_ = _rhs135
			var _rhs136 _dafny.Sequence = _1105_v127
			_ = _rhs136
			var _rhs137 _dafny.Array = (_1108_v130).Dtor_cf40()
			_ = _rhs137
			var _rhs138 _dafny.Int = (p0).Minus(_1107_v129)
			_ = _rhs138
			var _rhs139 _dafny.Int = _1102_i13
			_ = _rhs139
			var _lhs75 _dafny.Array = _1103_v126
			_ = _lhs75
			var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1103_v126), 0))
			_ = _lhs76
			r1 = _rhs135
			(_lhs75).ArraySet1(_rhs136, (_lhs76).Int())
			_1106_v128 = _rhs137
			_1107_v129 = _rhs138
			_1107_v129 = _rhs139
			var _1109_v131 _dafny.Map
			_ = _1109_v131
			_1109_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _1102_i13)
			var _1110_v133 _dafny.Array
			_ = _1110_v133
			var _len0_33 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_33
			var _nw194 _dafny.Array
			_ = _nw194
			if _len0_33.Cmp(_dafny.Zero) == 0 {
				_nw194 = _dafny.NewArray(_len0_33)
			} else {
				var _init33 func(_dafny.Int) _dafny.Map = (func(_1111_v131 _dafny.Map, _1112_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_1113_i15 _dafny.Int) _dafny.Map {
						return (func() _dafny.Map {
							var _coll33 = _dafny.NewMapBuilder()
							_ = _coll33
							for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-447), _dafny.IntOfInt64(915))); ; {
								_compr_33, _ok35 := _iter35()
								if !_ok35 {
									break
								}
								var _1114_v132 _dafny.Int
								_1114_v132 = interface{}(_compr_33).(_dafny.Int)
								if ((_dafny.IntOfInt64(-447)).Cmp(_1114_v132) <= 0) && ((_1114_v132).Cmp(_dafny.IntOfInt64(915)) < 0) {
									_coll33.Add((_1114_v132).Minus(_1112_p0), _1112_p0)
								}
							}
							return _coll33.ToMap()
						}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(515), (_1111_v131).Cardinality()))
					}
				})(_1109_v131, p0)
				_ = _init33
				var _element0_33 = _init33(_dafny.Zero)
				_ = _element0_33
				_nw194 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
				(_nw194).ArraySet1(_element0_33, 0)
				var _nativeLen0_33 = (_len0_33).Int()
				_ = _nativeLen0_33
				for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
					(_nw194).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
				}
			}
			_1110_v133 = _nw194
			var _1115_v134 _dafny.MultiSet
			_ = _1115_v134
			_1115_v134 = _dafny.MultiSetOf(p2)
			var _1116_v135 _dafny.Map
			_ = _1116_v135
			_1116_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1115_v134).Cardinality(), _dafny.IntOfUint32((p1).Cardinality()))
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_1110_v133), 0))
			_ = _index148
			(_1110_v133).ArraySet1(_1116_v135, (_index148).Int())
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_1110_v133), 0))
			_ = _index149
			var _rhs140 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, Companion_Default___.Fm9(_dafny.IntOfInt64(682), _dafny.IntOfUint32((p1).Cardinality()), (_this).F21(), globalState)))
			_ = _rhs140
			var _rhs141 _dafny.Map = _1109_v131
			_ = _rhs141
			var _rhs142 _dafny.Int = Companion_Default___.Fm9(_dafny.IntOfInt64(799), p0, (_dafny.MultiSetOf(!(false), (_this).F21())).IsProperSubsetOf(Companion_Default___.Fm28(_1102_i13, p2, globalState)), globalState)
			_ = _rhs142
			var _rhs143 _dafny.Map = Companion_Default___.Fm10(globalState)
			_ = _rhs143
			var _lhs77 _dafny.Array = _1110_v133
			_ = _lhs77
			var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_1110_v133), 0))
			_ = _lhs78
			_1107_v129 = _rhs140
			_1109_v131 = _rhs141
			_1107_v129 = _rhs142
			(_lhs77).ArraySet1(_rhs143, (_lhs78).Int())
			var _1117_v136 _dafny.Map
			_ = _1117_v136
			_1117_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(322))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}(func(_1118_i16 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('y')
			}))).Cardinality()), (_this).F21())
			var _1119_v137 _dafny.Sequence
			_ = _1119_v137
			_1119_v137 = _dafny.SeqOf(_1102_i13)
			var _1120_v138 _dafny.Set
			_ = _1120_v138
			_1120_v138 = _dafny.SetOf((_this).F21(), p2)
			var _1121_v139 _dafny.CodePoint
			_ = _1121_v139
			_1121_v139 = _dafny.CodePoint('q')
			var _1122_v140 _dafny.Map
			_ = _1122_v140
			_1122_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1121_v139, _1121_v139)
			var _1123_v141 _dafny.MultiSet
			_ = _1123_v141
			_1123_v141 = _dafny.MultiSetOf((_1122_v140).Cardinality(), _1107_v129, _1102_i13, _1107_v129, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1107_v129)).Cardinality())
			var _1124_v142 _dafny.Array
			_ = _1124_v142
			var _nwElement0_39 _dafny.Int = p0
			_ = _nwElement0_39
			var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(27))
			_ = _nw195
			(_nw195).ArraySet1(_nwElement0_39, 0)
			(_nw195).ArraySet1(p0, 1)
			(_nw195).ArraySet1(_1102_i13, 2)
			(_nw195).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_1107_v129, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _dafny.CodePoint('t'))).Cardinality()), 3)
			(_nw195).ArraySet1(_dafny.IntOfInt64(-593), 4)
			(_nw195).ArraySet1((_dafny.Zero).Minus(p0), 5)
			(_nw195).ArraySet1(_1107_v129, 6)
			(_nw195).ArraySet1(_1107_v129, 7)
			(_nw195).ArraySet1((_1117_v136).Cardinality(), 8)
			(_nw195).ArraySet1(_dafny.IntOfUint32((_1119_v137).Cardinality()), 9)
			(_nw195).ArraySet1(_1107_v129, 10)
			(_nw195).ArraySet1(_1107_v129, 11)
			(_nw195).ArraySet1(_1107_v129, 12)
			(_nw195).ArraySet1(p0, 13)
			(_nw195).ArraySet1(_1107_v129, 14)
			(_nw195).ArraySet1(_dafny.IntOfUint32((_1119_v137).Cardinality()), 15)
			(_nw195).ArraySet1(Companion_Default___.Fm9(_1107_v129, _dafny.IntOfUint32((p1).Cardinality()), p2, globalState), 16)
			(_nw195).ArraySet1((_dafny.MultiSetFromSeq(p1)).Cardinality(), 17)
			(_nw195).ArraySet1(p0, 18)
			(_nw195).ArraySet1(p0, 19)
			(_nw195).ArraySet1(_1107_v129, 20)
			(_nw195).ArraySet1(p0, 21)
			(_nw195).ArraySet1(_1102_i13, 22)
			(_nw195).ArraySet1((_1120_v138).Cardinality(), 23)
			(_nw195).ArraySet1((func() _dafny.Int {
				if (_1123_v141).Contains(_dafny.IntOfUint32((p1).Cardinality())) {
					return (_1123_v141).Multiplicity(_dafny.IntOfUint32((p1).Cardinality()))
				}
				return _1107_v129
			})(), 24)
			(_nw195).ArraySet1(_1107_v129, 25)
			(_nw195).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F21())).Cardinality(), 26)
			_1124_v142 = _nw195
			var _1125_v143 _dafny.MultiSet
			_ = _1125_v143
			_1125_v143 = _dafny.MultiSetOf(_1124_v142)
			var _1126_v144 _dafny.Sequence
			_ = _1126_v144
			_1126_v144 = _dafny.SeqOf(_1125_v143)
			_1107_v129 = ((_1126_v144).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1126_v144).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()
			var _1127_v145 _dafny.Array
			_ = _1127_v145
			var _nw196 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(5))
			_ = _nw196
			_1127_v145 = _nw196
			var _1128_v146 _dafny.Sequence
			_ = _1128_v146
			_1128_v146 = _dafny.SeqOf(_1127_v145, _1127_v145)
			var _1129_v147 _dafny.Array
			_ = _1129_v147
			var _nwElement0_40 _dafny.Array = _1127_v145
			_ = _nwElement0_40
			var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(26))
			_ = _nw197
			(_nw197).ArraySet1(_nwElement0_40, 0)
			(_nw197).ArraySet1(_1127_v145, 1)
			(_nw197).ArraySet1(_1127_v145, 2)
			(_nw197).ArraySet1(_1127_v145, 3)
			(_nw197).ArraySet1(_1127_v145, 4)
			(_nw197).ArraySet1(_1127_v145, 5)
			(_nw197).ArraySet1(_1127_v145, 6)
			(_nw197).ArraySet1(_1127_v145, 7)
			(_nw197).ArraySet1(_1127_v145, 8)
			(_nw197).ArraySet1(_1127_v145, 9)
			(_nw197).ArraySet1(_1127_v145, 10)
			(_nw197).ArraySet1(_1127_v145, 11)
			(_nw197).ArraySet1(_1127_v145, 12)
			(_nw197).ArraySet1(_1127_v145, 13)
			(_nw197).ArraySet1(_1127_v145, 14)
			(_nw197).ArraySet1((_1128_v146).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1128_v146).Cardinality()))).Uint32()).(_dafny.Array), 15)
			(_nw197).ArraySet1(_1127_v145, 16)
			(_nw197).ArraySet1(_1127_v145, 17)
			(_nw197).ArraySet1(_1127_v145, 18)
			(_nw197).ArraySet1((func() _dafny.Array {
				if (_this).F21() {
					return _1127_v145
				}
				return _1127_v145
			})(), 19)
			(_nw197).ArraySet1(_1127_v145, 20)
			(_nw197).ArraySet1(_1127_v145, 21)
			(_nw197).ArraySet1(_1127_v145, 22)
			(_nw197).ArraySet1(_1127_v145, 23)
			(_nw197).ArraySet1(_1127_v145, 24)
			(_nw197).ArraySet1(_1127_v145, 25)
			_1129_v147 = _nw197
			var _1130_v148 _dafny.Sequence
			_ = _1130_v148
			_1130_v148 = _dafny.SeqOf((_1103_v126).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1103_v126), 0))).Int()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate(_1105_v127, (_1103_v126).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1103_v126), 0))).Int()).(_dafny.Sequence)))
			var _1131_v149 _dafny.Sequence
			_ = _1131_v149
			_1131_v149 = _dafny.SeqOf(_1124_v142)
			var _1132_v150 _dafny.Sequence
			_ = _1132_v150
			_1132_v150 = _dafny.SeqOf(_1124_v142, (_1131_v149).Select((Companion_Default___.SafeIndex(_1107_v129, _dafny.IntOfUint32((_1131_v149).Cardinality()))).Uint32()).(_dafny.Array), _1124_v142, _1124_v142, _1124_v142)
			var _1133_v151 *C1
			_ = _1133_v151
			var _nw198 *C1 = New_C1_()
			_ = _nw198
			_nw198.Ctor__((_1132_v150).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_1132_v150).Cardinality()))).Uint32()).(_dafny.Array))
			_1133_v151 = _nw198
			var _1134_v152 D5
			_ = _1134_v152
			_1134_v152 = Companion_D5_.Create_DC16_(_1133_v151)
			var _1135_v153 _dafny.Array
			_ = _1135_v153
			var _nw199 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw199
			_1135_v153 = _nw199
			var _rhs144 _dafny.Array = _1129_v147
			_ = _rhs144
			var _rhs145 _dafny.Sequence = (func() _dafny.Sequence {
				if p2 {
					return _1130_v148
				}
				return _1130_v148
			})()
			_ = _rhs145
			var _rhs146 D5 = Companion_D5_.Create_DC16_(_1133_v151)
			_ = _rhs146
			var _rhs147 _dafny.Array = _1135_v153
			_ = _rhs147
			_1129_v147 = _rhs144
			_1130_v148 = _rhs145
			_1134_v152 = _rhs146
			_1135_v153 = _rhs147
		}
		var _1136_v154 _dafny.Map
		_ = _1136_v154
		_1136_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), (_this).F21())
		_1136_v154 = (_1136_v154).Update(!(p2), (_this).F21())
		var _1137_v155 _dafny.CodePoint
		_ = _1137_v155
		_1137_v155 = _dafny.CodePoint('e')
		var _1138_v156 _dafny.MultiSet
		_ = _1138_v156
		_1138_v156 = _dafny.MultiSetOf(_1137_v155)
		var _1139_v157 _dafny.Sequence
		_ = _1139_v157
		_1139_v157 = _dafny.SeqOf((_1138_v156).Cardinality())
		var _1140_v158 D0
		_ = _1140_v158
		_1140_v158 = Companion_D0_.Create_DC0_((_this).F21())
		var _1141_v159 D6
		_ = _1141_v159
		_1141_v159 = Companion_D6_.Create_DC20_(p2, _dafny.IntOfInt64(642), p2)
		var _1142_v160 _dafny.Array
		_ = _1142_v160
		var _nwElement0_41 bool = Companion_Default___.Fm0(p2, _1137_v155, true, globalState)
		_ = _nwElement0_41
		var _nw200 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(22))
		_ = _nw200
		(_nw200).ArraySet1(_nwElement0_41, 0)
		(_nw200).ArraySet1(p2, 1)
		(_nw200).ArraySet1((_this).F21(), 2)
		(_nw200).ArraySet1(p2, 3)
		(_nw200).ArraySet1(p2, 4)
		(_nw200).ArraySet1((_this).F21(), 5)
		(_nw200).ArraySet1(false, 6)
		(_nw200).ArraySet1((_this).F21(), 7)
		(_nw200).ArraySet1((_1140_v158).Dtor_cf0(), 8)
		(_nw200).ArraySet1(p2, 9)
		(_nw200).ArraySet1(Companion_Default___.Fm0(false, _1137_v155, (_1141_v159).Dtor_cf36(), globalState), 10)
		(_nw200).ArraySet1((_this).F21(), 11)
		(_nw200).ArraySet1(p2, 12)
		(_nw200).ArraySet1((_this).F21(), 13)
		(_nw200).ArraySet1((_this).F21(), 14)
		(_nw200).ArraySet1(!((_this).F21()), 15)
		(_nw200).ArraySet1((_this).F21(), 16)
		(_nw200).ArraySet1(false, 17)
		(_nw200).ArraySet1(false, 18)
		(_nw200).ArraySet1(false, 19)
		(_nw200).ArraySet1((_this).F21(), 20)
		(_nw200).ArraySet1((_this).F21(), 21)
		_1142_v160 = _nw200
		var _pat_let_tv29 = p0
		_ = _pat_let_tv29
		var _pat_let_tv30 = p0
		_ = _pat_let_tv30
		var _1143_v161 _dafny.Array
		_ = _1143_v161
		var _nwElement0_42 _dafny.Int = p0
		_ = _nwElement0_42
		var _nw201 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(19))
		_ = _nw201
		(_nw201).ArraySet1(_nwElement0_42, 0)
		(_nw201).ArraySet1(p0, 1)
		(_nw201).ArraySet1(p0, 2)
		(_nw201).ArraySet1(p0, 3)
		(_nw201).ArraySet1(p0, 4)
		(_nw201).ArraySet1(p0, 5)
		(_nw201).ArraySet1(_dafny.IntOfInt64(130), 6)
		(_nw201).ArraySet1(p0, 7)
		(_nw201).ArraySet1(p0, 8)
		(_nw201).ArraySet1(p0, 9)
		(_nw201).ArraySet1(_dafny.IntOfInt64(714), 10)
		(_nw201).ArraySet1(_dafny.IntOfUint32((_1139_v157).Cardinality()), 11)
		(_nw201).ArraySet1((func(_pat_let15_0 D2) D2 {
			return func(_1144_dt__update__tmp_h1 D2) D2 {
				return func(_pat_let16_0 _dafny.Int) D2 {
					return func(_1145_dt__update_hcf15_h0 _dafny.Int) D2 {
						return func(_pat_let17_0 _dafny.Int) D2 {
							return func(_1146_dt__update_hcf14_h0 _dafny.Int) D2 {
								return Companion_D2_.Create_DC8_((_1144_dt__update__tmp_h1).Dtor_cf13(), _1146_dt__update_hcf14_h0, _1145_dt__update_hcf15_h0)
							}(_pat_let17_0)
						}(_pat_let_tv30)
					}(_pat_let16_0)
				}((_dafny.Zero).Minus(_pat_let_tv29))
			}(_pat_let15_0)
		}(Companion_D2_.Create_DC8_(_1142_v160, p0, (_1136_v154).Cardinality()))).Dtor_cf15(), 12)
		(_nw201).ArraySet1(p0, 13)
		(_nw201).ArraySet1(p0, 14)
		(_nw201).ArraySet1(p0, 15)
		(_nw201).ArraySet1(_dafny.IntOfInt64(20), 16)
		(_nw201).ArraySet1(p0, 17)
		(_nw201).ArraySet1(p0, 18)
		_1143_v161 = _nw201
		var _1147_v162 *C1
		_ = _1147_v162
		var _nw202 *C1 = New_C1_()
		_ = _nw202
		_nw202.Ctor__(_1143_v161)
		_1147_v162 = _nw202
		var _1148_v163 _dafny.Set
		_ = _1148_v163
		_1148_v163 = _dafny.SetOf(_1147_v162)
		var _1149_v164 _dafny.Map
		_ = _1149_v164
		_1149_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1148_v163)
		r0 = !(((_1148_v163).Equals((func() _dafny.Set {
			if (_1149_v164).Contains(p1) {
				return (_1149_v164).Get(p1).(_dafny.Set)
			}
			return _1148_v163
		})())) || ((_dafny.IntOfInt64(248)).Cmp(p0) >= 0))
		r1 = Companion_Default___.Fm0(p2, _1137_v155, (p0).Cmp(p0) > 0, globalState)
		r2 = ((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p0, p0)))).Cmp(p0) != 0
		r3 = p1
		return r0, r1, r2, r3
	}
}
func (_this *C2) F21() bool {
	{
		return _this._f21
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	dummy byte
}

func New_C3_() *C3 {
	_this := C3{}

	return &_this
}

type CompanionStruct_C3_ struct {
}

var Companion_C3_ = CompanionStruct_C3_{}

func (_this *C3) Equals(other *C3) bool {
	return _this == other
}

func (_this *C3) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C3)
	return ok && _this.Equals(other)
}

func (*C3) String() string {
	return "_module.C3"
}

func Type_C3_() _dafny.TypeDescriptor {
	return type_C3_{}
}

type type_C3_ struct {
}

func (_this type_C3_) Default() interface{} {
	return (*C3)(nil)
}

func (_this type_C3_) String() string {
	return "main.C3"
}
func (_this *C3) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) Ctor__() {
	{
	}
}
func (_this *C3) Fm6(p0 _dafny.Map, globalState *GlobalState) bool {
	{
		return ((_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()).Plus((_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.SetOf((_dafny.SetOf(!(true))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("b")).Cardinality()), _dafny.IntOfInt64(697), _dafny.IntOfInt64(758), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dssq")).Cardinality()))).Cardinality()))).Cardinality()))).Cmp((_dafny.IntOfInt64(414)).Minus(_dafny.IntOfInt64(985))) == 0
	}
}
func (_this *C3) Fm7(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		var _source13 D0 = Companion_D0_.Create_DC0_(true)
		_ = _source13
		if _source13.Is_DC1() {
			var _1150___mcc_h0 bool = _source13.Get_().(D0_DC1).Cf1
			_ = _1150___mcc_h0
			var _1151_cf1 bool = _1150___mcc_h0
			_ = _1151_cf1
			return _1151_cf1
		} else if _source13.Is_DC0() {
			var _1152___mcc_h1 bool = _source13.Get_().(D0_DC0).Cf0
			_ = _1152___mcc_h1
			var _1153_cf0 bool = _1152___mcc_h1
			_ = _1153_cf0
			return _1153_cf0
		} else {
			var _1154___mcc_h2 D0 = _source13.Get_().(D0_DC2).Cf2
			_ = _1154___mcc_h2
			var _1155_cf2 D0 = _1154___mcc_h2
			_ = _1155_cf2
			return false
		}
	}
}
func (_this *C3) M4(p0 _dafny.Int, p1 bool, p2 _dafny.Array, p3 _dafny.CodePoint, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, _dafny.Map, _dafny.Sequence) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		(globalState).F11 = true
		var _1156_i0 _dafny.Int
		_ = _1156_i0
		_1156_i0 = _dafny.Zero
		{
			for p1 {
				{
					if (_1156_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1156_i0 = (_1156_i0).Plus(_dafny.One)
					var _1157_v0 _dafny.Array
					_ = _1157_v0
					var _len0_34 _dafny.Int = _dafny.IntOfInt64(23)
					_ = _len0_34
					var _nw203 _dafny.Array
					_ = _nw203
					if _len0_34.Cmp(_dafny.Zero) == 0 {
						_nw203 = _dafny.NewArray(_len0_34)
					} else {
						var _init34 func(_dafny.Int) _dafny.Int = (func(_1158_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1159_i1 _dafny.Int) _dafny.Int {
								return (_1159_i1).Times(_1158_p0)
							}
						})(p0)
						_ = _init34
						var _element0_34 = _init34(_dafny.Zero)
						_ = _element0_34
						_nw203 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
						(_nw203).ArraySet1(_element0_34, 0)
						var _nativeLen0_34 = (_len0_34).Int()
						_ = _nativeLen0_34
						for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
							(_nw203).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
						}
					}
					_1157_v0 = _nw203
					var _1160_v1 *C1
					_ = _1160_v1
					var _nw204 *C1 = New_C1_()
					_ = _nw204
					_nw204.Ctor__(_1157_v0)
					_1160_v1 = _nw204
					var _1161_v2 _dafny.Sequence
					_ = _1161_v2
					_1161_v2 = _dafny.SeqOf(p2, p2)
					var _1162_v3 _dafny.Set
					_ = _1162_v3
					_1162_v3 = _dafny.SetOf(p0, p0)
					var _1163_v4 _dafny.Sequence
					_ = _1163_v4
					_1163_v4 = _dafny.SeqOf(_1162_v3)
					var _1164_v5 _dafny.Sequence
					_ = _1164_v5
					_1164_v5 = _dafny.SeqOf((_1161_v2).Select((Companion_Default___.SafeIndex(((_1163_v4).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1163_v4).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), _dafny.IntOfUint32((_1161_v2).Cardinality()))).Uint32()).(_dafny.Array), p2)
					var _1165_v6 _dafny.Map
					_ = _1165_v6
					_1165_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-984)), _1164_v5)
					var _1166_v7 _dafny.Sequence
					_ = _1166_v7
					_1166_v7 = _dafny.SeqOf(p1)
					var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))
					_ = _index150
					(_1157_v0).ArraySet1((_dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_1165_v6).Contains(p0) {
							return (_1165_v6).Get(p0).(_dafny.Sequence)
						}
						return _dafny.Companion_Sequence_.Update(_1164_v5, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1164_v5).Cardinality()))).Uint32(), p2)
					})()).Cardinality())).Plus((_dafny.MultiSetFromSeq(_1166_v7)).Cardinality()), (_index150).Int())
					var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))
					_ = _index151
					(_1157_v0).ArraySet1((p0).Times(p0), (_index151).Int())
					var _1167_v8 _dafny.Array
					_ = _1167_v8
					var _nw205 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
					_ = _nw205
					_1167_v8 = _nw205
					var _1168_v9 _dafny.Map
					_ = _1168_v9
					_1168_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1167_v8, !(true))
					if (func() bool {
						if (_1168_v9).Contains(_1167_v8) {
							return (_1168_v9).Get(_1167_v8).(bool)
						}
						return (func() bool {
							if true {
								return p1
							}
							return p1
						})()
					})() {
						(globalState).F15 = (_this).Fm7((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), p1, (p0).Times(_dafny.IntOfInt64(-584)), globalState)
						var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((p2), 0))
						_ = _index152
						(p2).ArraySet1((func() bool {
							if false {
								return p1
							}
							return false
						})(), (_index152).Int())
						var _1169_v10 _dafny.Sequence
						_ = _1169_v10
						_1169_v10 = _dafny.UnicodeSeqOfUtf8Bytes("wngf")
						var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((p2), 0))
						_ = _index153
						var _rhs148 _dafny.Sequence = _1169_v10
						_ = _rhs148
						var _rhs149 bool = p1
						_ = _rhs149
						var _lhs79 _dafny.Array = p2
						_ = _lhs79
						var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((p2), 0))
						_ = _lhs80
						r3 = _rhs148
						(_lhs79).ArraySet1(_rhs149, (_lhs80).Int())
						var _1170_v11 _dafny.Map
						_ = _1170_v11
						_1170_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1169_v10).Cardinality()), p1)
						var _1171_v12 _dafny.Map
						_ = _1171_v12
						_1171_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))
						var _1172_v13 _dafny.Map
						_ = _1172_v13
						_1172_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((p2), 0))).Int()).(bool), (func() bool {
							if (_1170_v11).Contains(p0) {
								return (_1170_v11).Get(p0).(bool)
							}
							return (_this).Fm6(_1171_v12, globalState)
						})())
						var _1173_v14 D6
						_ = _1173_v14
						_1173_v14 = Companion_D6_.Create_DC20_(_dafny.Companion_Sequence_.IsProperPrefixOf(_1169_v10, _dafny.Companion_Sequence_.Update(_1169_v10, (Companion_Default___.SafeIndex((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1169_v10).Cardinality()))).Uint32(), _dafny.CodePoint('v'))), (_dafny.Zero).Minus((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int)), (func() bool {
							if (_1172_v13).Contains((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((p2), 0))).Int()).(bool)) {
								return (_1172_v13).Get((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((p2), 0))).Int()).(bool)).(bool)
							}
							return p1
						})())
						_1173_v14 = Companion_D6_.Create_DC20_(true, Companion_Default___.Fm9(p0, _dafny.IntOfUint32((_1169_v10).Cardinality()), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((p2), 0))).Int()).(bool), globalState), p1)
						(globalState).F11 = p1
						var _1174_v15 _dafny.Array
						_ = _1174_v15
						var _len0_35 _dafny.Int = _dafny.IntOfInt64(25)
						_ = _len0_35
						var _nw206 _dafny.Array
						_ = _nw206
						if _len0_35.Cmp(_dafny.Zero) == 0 {
							_nw206 = _dafny.NewArray(_len0_35)
						} else {
							var _init35 func(_dafny.Int) bool = (func(_1175_p1 bool) func(_dafny.Int) bool {
								return func(_1176_i2 _dafny.Int) bool {
									return _1175_p1
								}
							})(p1)
							_ = _init35
							var _element0_35 = _init35(_dafny.Zero)
							_ = _element0_35
							_nw206 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
							(_nw206).ArraySet1(_element0_35, 0)
							var _nativeLen0_35 = (_len0_35).Int()
							_ = _nativeLen0_35
							for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
								(_nw206).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
							}
						}
						_1174_v15 = _nw206
						var _1177_v16 _dafny.Array
						_ = _1177_v16
						var _len0_36 _dafny.Int = _dafny.IntOfInt64(28)
						_ = _len0_36
						var _nw207 _dafny.Array
						_ = _nw207
						if _len0_36.Cmp(_dafny.Zero) == 0 {
							_nw207 = _dafny.NewArray(_len0_36)
						} else {
							var _init36 func(_dafny.Int) _dafny.Sequence = (func(_1178_v10 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_1179_i3 _dafny.Int) _dafny.Sequence {
									return _dafny.SeqOf(_dafny.IntOfUint32((_1178_v10).Cardinality()))
								}
							})(_1169_v10)
							_ = _init36
							var _element0_36 = _init36(_dafny.Zero)
							_ = _element0_36
							_nw207 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
							(_nw207).ArraySet1(_element0_36, 0)
							var _nativeLen0_36 = (_len0_36).Int()
							_ = _nativeLen0_36
							for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
								(_nw207).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
							}
						}
						_1177_v16 = _nw207
						var _1180_v17 _dafny.MultiSet
						_ = _1180_v17
						_1180_v17 = _dafny.MultiSetOf(p0)
						var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_1177_v16), 0))
						_ = _index154
						(_1177_v16).ArraySet1(_dafny.SeqOf((_1180_v17).Cardinality()), (_index154).Int())
						var _1181_v18 _dafny.Map
						_ = _1181_v18
						_1181_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_Default___.Fm16(p1, p0, globalState), Companion_Default___.Fm16((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((p2), 0))).Int()).(bool), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cfbueqff")).Cardinality()), globalState))), _1174_v15)
						var _1182_v19 _dafny.MultiSet
						_ = _1182_v19
						_1182_v19 = _dafny.MultiSetOf(_1169_v10, _1169_v10, _1169_v10, _1169_v10, _1169_v10)
						var _1183_v20 _dafny.Sequence
						_ = _1183_v20
						_1183_v20 = _dafny.SeqOf((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), p0, (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p0, (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))))
						var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_1177_v16), 0))
						_ = _index155
						var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))
						_ = _index156
						var _rhs150 _dafny.Array = (func() _dafny.Array {
							if (_1181_v18).Contains(_1182_v19) {
								return (_1181_v18).Get(_1182_v19).(_dafny.Array)
							}
							return _1174_v15
						})()
						_ = _rhs150
						var _rhs151 _dafny.Sequence = _1183_v20
						_ = _rhs151
						var _rhs152 _dafny.Int = p0
						_ = _rhs152
						var _lhs81 _dafny.Array = _1177_v16
						_ = _lhs81
						var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_1177_v16), 0))
						_ = _lhs82
						var _lhs83 _dafny.Array = _1157_v0
						_ = _lhs83
						var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))
						_ = _lhs84
						_1174_v15 = _rhs150
						(_lhs81).ArraySet1(_rhs151, (_lhs82).Int())
						(_lhs83).ArraySet1(_rhs152, (_lhs84).Int())
					} else {
						var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))
						_ = _index157
						(_1157_v0).ArraySet1(p0, (_index157).Int())
						var _1184_v21 *C0
						_ = _1184_v21
						var _nw208 *C0 = New_C0_()
						_ = _nw208
						_nw208.Ctor__((func() bool {
							if Companion_Default___.Fm0(p1, p3, p1, globalState) {
								return p1
							}
							return p1
						})())
						_1184_v21 = _nw208
						var _1185_v22 _dafny.Sequence
						_ = _1185_v22
						_1185_v22 = _dafny.UnicodeSeqOfUtf8Bytes("xergafjq")
						r0 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1184_v21).F23(), p1, Companion_Default___.Fm0(p1, (_1185_v22).Select((Companion_Default___.SafeIndex((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1185_v22).Cardinality()))).Uint32()).(_dafny.CodePoint), p1, globalState)), _1166_v7), (Companion_Default___.SafeIndex((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1184_v21).F23(), p1, Companion_Default___.Fm0(p1, (_1185_v22).Select((Companion_Default___.SafeIndex((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1185_v22).Cardinality()))).Uint32()).(_dafny.CodePoint), p1, globalState)), _1166_v7)).Cardinality()))).Uint32(), p1)
						var _1186_v23 _dafny.MultiSet
						_ = _1186_v23
						_1186_v23 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))).Cardinality(), (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), p0)
						r1 = ((func() _dafny.Int {
							if (_1186_v23).Contains((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int)) {
								return (_1186_v23).Multiplicity((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))
							}
							return p0
						})()).Minus((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))
						var _1187_v24 _dafny.Map
						_ = _1187_v24
						_1187_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1160_v1)
						_1187_v24 = (_1187_v24).Update((Companion_D4_.Create_DC14_(p1, (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), _1162_v3, _dafny.CodePoint('s'))).Dtor_cf26(), _1160_v1)
					}
					if p1 {
						var _1188_v25 _dafny.Map
						_ = _1188_v25
						_1188_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1162_v3)
						r1 = ((((func() _dafny.Set {
							if (_1188_v25).Contains((_dafny.Zero).Minus((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))) {
								return (_1188_v25).Get((_dafny.Zero).Minus((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))).(_dafny.Set)
							}
							return _1162_v3
						})()).Intersection(_1162_v3)).Cardinality()).Times((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))
						var _1189_v26 _dafny.MultiSet
						_ = _1189_v26
						_1189_v26 = _dafny.MultiSetOf(p1)
						var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))
						_ = _index158
						var _rhs153 _dafny.MultiSet = _1189_v26
						_ = _rhs153
						var _rhs154 bool = (func() bool {
							if p1 {
								return (p1) && (!(!(p1)))
							}
							return p1
						})()
						_ = _rhs154
						var _rhs155 _dafny.Int = (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int)
						_ = _rhs155
						var _lhs85 *GlobalState = globalState
						_ = _lhs85
						var _lhs86 _dafny.Array = _1157_v0
						_ = _lhs86
						var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))
						_ = _lhs87
						_1189_v26 = _rhs153
						_lhs85.F15 = _rhs154
						(_lhs86).ArraySet1(_rhs155, (_lhs87).Int())
						var _1190_v27 _dafny.MultiSet
						_ = _1190_v27
						_1190_v27 = _dafny.MultiSetOf(p0)
						var _1191_v28 D6
						_ = _1191_v28
						_1191_v28 = Companion_D6_.Create_DC20_(!(!(p1)) || (true), ((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int)).Plus((_1190_v27).Cardinality()), p1)
						var _1192_v29 _dafny.Set
						_ = _1192_v29
						_1192_v29 = _dafny.SetOf(p1)
						var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))
						_ = _index159
						var _rhs156 D6 = (func() D6 {
							if (_1192_v29).IsDisjointFrom(_1192_v29) {
								return _1191_v28
							}
							return _1191_v28
						})()
						_ = _rhs156
						var _rhs157 bool = p1
						_ = _rhs157
						var _rhs158 bool = p1
						_ = _rhs158
						var _rhs159 _dafny.Int = Companion_Default___.SafeDivisionInt((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))
						_ = _rhs159
						var _lhs88 *GlobalState = globalState
						_ = _lhs88
						var _lhs89 *GlobalState = globalState
						_ = _lhs89
						var _lhs90 _dafny.Array = _1157_v0
						_ = _lhs90
						var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))
						_ = _lhs91
						_1191_v28 = _rhs156
						_lhs88.F15 = _rhs157
						_lhs89.F11 = _rhs158
						(_lhs90).ArraySet1(_rhs159, (_lhs91).Int())
						var _1193_v30 _dafny.Sequence
						_ = _1193_v30
						_1193_v30 = _dafny.UnicodeSeqOfUtf8Bytes("gghwthy")
						var _1194_v31 _dafny.Sequence
						_ = _1194_v31
						_1194_v31 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pqjkhiid"), _dafny.Companion_Sequence_.Update(_1193_v30, (Companion_Default___.SafeIndex((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1193_v30).Cardinality()))).Uint32(), p3))
						var _1195_v32 _dafny.Map
						_ = _1195_v32
						_1195_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))
						_1194_v31 = (func() _dafny.Sequence {
							if (_this).Fm6(_1195_v32, globalState) {
								return _1194_v31
							}
							return _dafny.SeqOf(_1193_v30)
						})()
						r1 = _dafny.IntOfInt64(277)
					} else {
						var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))
						_ = _index160
						(_1157_v0).ArraySet1((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), (_index160).Int())
						var _1196_v33 _dafny.Map
						_ = _1196_v33
						_1196_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))
						var _1197_v35 _dafny.Sequence
						_ = _1197_v35
						_1197_v35 = _dafny.UnicodeSeqOfUtf8Bytes("t")
						var _1198_v36 _dafny.MultiSet
						_ = _1198_v36
						_1198_v36 = _dafny.MultiSetOf(_1196_v33, func() _dafny.Map {
							var _coll34 = _dafny.NewMapBuilder()
							_ = _coll34
							for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(617), _dafny.IntOfInt64(957))); ; {
								_compr_34, _ok36 := _iter36()
								if !_ok36 {
									break
								}
								var _1199_v34 _dafny.Int
								_1199_v34 = interface{}(_compr_34).(_dafny.Int)
								if ((_dafny.IntOfInt64(617)).Cmp(_1199_v34) <= 0) && ((_1199_v34).Cmp(_dafny.IntOfInt64(957)) < 0) {
									_coll34.Add((_1199_v34).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uvscymnab")).Cardinality())), p0)
								}
							}
							return _coll34.ToMap()
						}(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1197_v35).Cardinality()), (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))).Merge(_1196_v33))
						var _1200_v37 _dafny.Map
						_ = _1200_v37
						_1200_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))
						var _1201_v38 _dafny.Sequence
						_ = _1201_v38
						_1201_v38 = _dafny.SeqOf((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), p0, Companion_Default___.Fm9((_1200_v37).Cardinality(), p0, p1, globalState))
						var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))
						_ = _index161
						var _rhs160 bool = false
						_ = _rhs160
						var _rhs161 bool = (_1162_v3).IsSubsetOf(_1162_v3)
						_ = _rhs161
						var _rhs162 _dafny.Array = (_1160_v1).F22()
						_ = _rhs162
						var _rhs163 _dafny.Int = (func() _dafny.Int {
							if (_1198_v36).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(54), (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))).Merge(_1196_v33)) {
								return (_1198_v36).Multiplicity((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(54), (_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int))).Merge(_1196_v33))
							}
							return (_1201_v38).Select((Companion_Default___.SafeIndex((_1157_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1201_v38).Cardinality()))).Uint32()).(_dafny.Int)
						})()
						_ = _rhs163
						var _rhs164 _dafny.Int = p0
						_ = _rhs164
						var _lhs92 *GlobalState = globalState
						_ = _lhs92
						var _lhs93 *GlobalState = globalState
						_ = _lhs93
						var _lhs94 _dafny.Array = _1157_v0
						_ = _lhs94
						var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1157_v0), 0))
						_ = _lhs95
						_lhs92.F11 = _rhs160
						_lhs93.F11 = _rhs161
						_1157_v0 = _rhs162
						r1 = _rhs163
						(_lhs94).ArraySet1(_rhs164, (_lhs95).Int())
						var _1202_v39 D9
						_ = _1202_v39
						_1202_v39 = Companion_D9_.Create_DC24_(_1201_v38)
						var _rhs165 bool = (_dafny.IntOfInt64(216)).Cmp((_dafny.IntOfUint32(((_1202_v39).Dtor_cf42()).Cardinality())).Minus(_dafny.IntOfInt64(970))) != 0
						_ = _rhs165
						var _rhs166 bool = !((false) == (false))
						_ = _rhs166
						var _lhs96 *GlobalState = globalState
						_ = _lhs96
						var _lhs97 *GlobalState = globalState
						_ = _lhs97
						_lhs96.F15 = _rhs165
						_lhs97.F11 = _rhs166
						var _1203_v40 _dafny.Map
						_ = _1203_v40
						_1203_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1201_v38, p1)
						_1203_v40 = (_1203_v40).Update(_1201_v38, p1)
						var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((p2), 0))
						_ = _index162
						(p2).ArraySet1(p1, (_index162).Int())
						var _1204_v41 _dafny.Array
						_ = _1204_v41
						var _len0_37 _dafny.Int = _dafny.IntOfInt64(24)
						_ = _len0_37
						var _nw209 _dafny.Array
						_ = _nw209
						if _len0_37.Cmp(_dafny.Zero) == 0 {
							_nw209 = _dafny.NewArray(_len0_37)
						} else {
							var _init37 func(_dafny.Int) _dafny.Sequence = (func(_1205_v35 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_1206_i4 _dafny.Int) _dafny.Sequence {
									return _1205_v35
								}
							})(_1197_v35)
							_ = _init37
							var _element0_37 = _init37(_dafny.Zero)
							_ = _element0_37
							_nw209 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
							(_nw209).ArraySet1(_element0_37, 0)
							var _nativeLen0_37 = (_len0_37).Int()
							_ = _nativeLen0_37
							for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
								(_nw209).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
							}
						}
						_1204_v41 = _nw209
						var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_1204_v41), 0))
						_ = _index163
						(_1204_v41).ArraySet1(Companion_Default___.Fm16(true, p0, globalState), (_index163).Int())
						var _1207_v42 _dafny.Array
						_ = _1207_v42
						var _nwElement0_43 _dafny.Array = (_1160_v1).F22()
						_ = _nwElement0_43
						var _nw210 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(5))
						_ = _nw210
						(_nw210).ArraySet1(_nwElement0_43, 0)
						(_nw210).ArraySet1((_1160_v1).F22(), 1)
						(_nw210).ArraySet1(_1157_v0, 2)
						(_nw210).ArraySet1((_1160_v1).F22(), 3)
						(_nw210).ArraySet1((_1160_v1).F22(), 4)
						_1207_v42 = _nw210
						var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((p2), 0))
						_ = _index164
						var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_1204_v41), 0))
						_ = _index165
						var _rhs167 bool = p1
						_ = _rhs167
						var _rhs168 _dafny.Array = (func() _dafny.Array {
							if p1 {
								return _1207_v42
							}
							return _1207_v42
						})()
						_ = _rhs168
						var _rhs169 _dafny.Sequence = _1197_v35
						_ = _rhs169
						var _lhs98 _dafny.Array = p2
						_ = _lhs98
						var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((p2), 0))
						_ = _lhs99
						var _lhs100 *GlobalState = globalState
						_ = _lhs100
						var _lhs101 _dafny.Array = _1204_v41
						_ = _lhs101
						var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_1204_v41), 0))
						_ = _lhs102
						(_lhs98).ArraySet1(_rhs167, (_lhs99).Int())
						_lhs100.F2 = _rhs168
						(_lhs101).ArraySet1(_rhs169, (_lhs102).Int())
					}
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1208_v43 _dafny.MultiSet
		_ = _1208_v43
		_1208_v43 = _dafny.MultiSetOf(_dafny.IntOfInt64(640))
		_1208_v43 = _1208_v43
		var _1209_v46 _dafny.Array
		_ = _1209_v46
		var _len0_38 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_38
		var _nw211 _dafny.Array
		_ = _nw211
		if _len0_38.Cmp(_dafny.Zero) == 0 {
			_nw211 = _dafny.NewArray(_len0_38)
		} else {
			var _init38 func(_dafny.Int) _dafny.Map = (func(_1210_p0 _dafny.Int, _1211_p1 bool) func(_dafny.Int) _dafny.Map {
				return func(_1212_i5 _dafny.Int) _dafny.Map {
					return (func() _dafny.Map {
						var _coll35 = _dafny.NewMapBuilder()
						_ = _coll35
						for _iter37 := _dafny.Iterate((_dafny.SeqOf(_1210_p0)).Elements()); ; {
							_compr_35, _ok37 := _iter37()
							if !_ok37 {
								break
							}
							var _1213_v44 _dafny.Int
							_1213_v44 = interface{}(_compr_35).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_1210_p0), _1213_v44) {
								_coll35.Add((_1213_v44).Times(_1210_p0), _1211_p1)
							}
						}
						return _coll35.ToMap()
					}()).Merge(func() _dafny.Map {
						var _coll36 = _dafny.NewMapBuilder()
						_ = _coll36
						for _iter38 := _dafny.Iterate((_dafny.MultiSetOf(_1210_p0)).Elements()); ; {
							_compr_36, _ok38 := _iter38()
							if !_ok38 {
								break
							}
							var _1214_v45 _dafny.Int
							_1214_v45 = interface{}(_compr_36).(_dafny.Int)
							if (_dafny.MultiSetOf(_1210_p0)).Contains(_1214_v45) {
								_coll36.Add((_1214_v45).Minus(_1210_p0), _1211_p1)
							}
						}
						return _coll36.ToMap()
					}())
				}
			})(p0, p1)
			_ = _init38
			var _element0_38 = _init38(_dafny.Zero)
			_ = _element0_38
			_nw211 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
			(_nw211).ArraySet1(_element0_38, 0)
			var _nativeLen0_38 = (_len0_38).Int()
			_ = _nativeLen0_38
			for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
				(_nw211).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
			}
		}
		_1209_v46 = _nw211
		var _nw212 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
		_ = _nw212
		_1209_v46 = _nw212
		var _1215_v47 _dafny.Map
		_ = _1215_v47
		_1215_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm9(p0, _dafny.IntOfInt64(383), !(p1), globalState))
		var _1216_v48 _dafny.Map
		_ = _1216_v48
		_1216_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
		_1215_v47 = (_1215_v47).Update(((_1216_v48).Merge(_1216_v48)).Cardinality(), p0)
		var _hi8 _dafny.Int = p0
		_ = _hi8
		for _1217_i6 := p0; _1217_i6.Cmp(_hi8) < 0; _1217_i6 = _1217_i6.Plus(_dafny.One) {
			(globalState).F11 = p1
			(globalState).F11 = (p1) && (p1)
			var _1218_v49 D6
			_ = _1218_v49
			_1218_v49 = Companion_D6_.Create_DC20_(false, p0, p1)
			var _1219_v50 _dafny.Map
			_ = _1219_v50
			_1219_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _pat_let_tv31 = p1
			_ = _pat_let_tv31
			var _pat_let_tv32 = _1219_v50
			_ = _pat_let_tv32
			var _1220_v51 _dafny.Sequence
			_ = _1220_v51
			_1220_v51 = _dafny.SeqOf(_1218_v49, func(_pat_let18_0 D6) D6 {
				return func(_1221_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let19_0 bool) D6 {
						return func(_1222_dt__update_hcf36_h0 bool) D6 {
							return func(_pat_let20_0 _dafny.Int) D6 {
								return func(_1223_dt__update_hcf37_h0 _dafny.Int) D6 {
									return Companion_D6_.Create_DC20_(_1222_dt__update_hcf36_h0, _1223_dt__update_hcf37_h0, (_1221_dt__update__tmp_h0).Dtor_cf38())
								}(_pat_let20_0)
							}((_pat_let_tv32).Cardinality())
						}(_pat_let19_0)
					}(!(_pat_let_tv31))
				}(_pat_let18_0)
			}(_1218_v49))
			var _1224_v52 _dafny.CodePoint
			_ = _1224_v52
			_1224_v52 = _dafny.CodePoint('j')
			var _rhs170 _dafny.Sequence = _1220_v51
			_ = _rhs170
			var _rhs171 _dafny.Int = p0
			_ = _rhs171
			var _rhs172 _dafny.CodePoint = (func() _dafny.CodePoint {
				if true {
					return _1224_v52
				}
				return _1224_v52
			})()
			_ = _rhs172
			_1220_v51 = _rhs170
			r1 = _rhs171
			_1224_v52 = _rhs172
			var _1225_v53 D0
			_ = _1225_v53
			_1225_v53 = Companion_D0_.Create_DC1_(true)
			var _pat_let_tv33 = p1
			_ = _pat_let_tv33
			var _1226_v54 *C2
			_ = _1226_v54
			var _nw213 *C2 = New_C2_()
			_ = _nw213
			_nw213.Ctor__(func(_pat_let21_0 D0) D0 {
				return func(_1227_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let22_0 bool) D0 {
						return func(_1228_dt__update_hcf1_h0 bool) D0 {
							return Companion_D0_.Create_DC1_(_1228_dt__update_hcf1_h0)
						}(_pat_let22_0)
					}(!(_pat_let_tv33))
				}(_pat_let21_0)
			}(_1225_v53), !((p1) && (p1)))
			_1226_v54 = _nw213
		}
		var _1229_v55 _dafny.Sequence
		_ = _1229_v55
		_1229_v55 = _dafny.SeqOf(p1, p1, false, true, p1)
		r0 = _dafny.Companion_Sequence_.Concatenate(_1229_v55, _1229_v55)
		r1 = p0
		var _1230_v56 _dafny.Sequence
		_ = _1230_v56
		_1230_v56 = _dafny.SeqOf(_1216_v48)
		r2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1230_v56).Cardinality()), p0), (p1) == (!(p1)))
		var _1231_v57 _dafny.Sequence
		_ = _1231_v57
		_1231_v57 = _dafny.UnicodeSeqOfUtf8Bytes("suhxbcocc")
		r3 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1231_v57, _1231_v57), _1231_v57)
		return r0, r1, r2, r3
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f19 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f19 = false
	return &_this
}

type CompanionStruct_C4_ struct {
}

var Companion_C4_ = CompanionStruct_C4_{}

func (_this *C4) Equals(other *C4) bool {
	return _this == other
}

func (_this *C4) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C4)
	return ok && _this.Equals(other)
}

func (*C4) String() string {
	return "_module.C4"
}

func Type_C4_() _dafny.TypeDescriptor {
	return type_C4_{}
}

type type_C4_ struct {
}

func (_this type_C4_) Default() interface{} {
	return (*C4)(nil)
}

func (_this type_C4_) String() string {
	return "main.C4"
}
func (_this *C4) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C4{}
var _ T1 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) Ctor__(f19 bool) {
	{
		(_this)._f19 = f19
	}
}
func (_this *C4) Fm4(p0 bool, p1 _dafny.Map, p2 _dafny.CodePoint, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_this).F19()
	}
}
func (_this *C4) M1(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _1232_v0 _dafny.Array
		_ = _1232_v0
		var _nw214 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
		_ = _nw214
		_1232_v0 = _nw214
		var _1233_v1 _dafny.CodePoint
		_ = _1233_v1
		_1233_v1 = _dafny.CodePoint('u')
		var _1234_v2 _dafny.Sequence
		_ = _1234_v2
		_1234_v2 = _dafny.SeqOf(_1233_v1)
		var _1235_v3 _dafny.Map
		_ = _1235_v3
		_1235_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1234_v2, _1232_v0)
		var _1236_v4 _dafny.Sequence
		_ = _1236_v4
		_1236_v4 = _dafny.SeqOf(_1232_v0)
		var _1237_v5 _dafny.Array
		_ = _1237_v5
		var _nwElement0_44 _dafny.Array = _1232_v0
		_ = _nwElement0_44
		var _nw215 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(19))
		_ = _nw215
		(_nw215).ArraySet1(_nwElement0_44, 0)
		(_nw215).ArraySet1(_1232_v0, 1)
		(_nw215).ArraySet1(_1232_v0, 2)
		(_nw215).ArraySet1((func() _dafny.Array {
			if p0 {
				return _1232_v0
			}
			return _1232_v0
		})(), 3)
		(_nw215).ArraySet1(_1232_v0, 4)
		(_nw215).ArraySet1(_1232_v0, 5)
		(_nw215).ArraySet1(_1232_v0, 6)
		(_nw215).ArraySet1(_1232_v0, 7)
		(_nw215).ArraySet1((func() _dafny.Array {
			if (_1235_v3).Contains(_dafny.SeqOf(_1233_v1)) {
				return (_1235_v3).Get(_dafny.SeqOf(_1233_v1)).(_dafny.Array)
			}
			return _1232_v0
		})(), 8)
		(_nw215).ArraySet1(_1232_v0, 9)
		(_nw215).ArraySet1(_1232_v0, 10)
		(_nw215).ArraySet1(_1232_v0, 11)
		(_nw215).ArraySet1(_1232_v0, 12)
		(_nw215).ArraySet1((_1236_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-8), _dafny.IntOfUint32((_1236_v4).Cardinality()))).Uint32()).(_dafny.Array), 13)
		(_nw215).ArraySet1(_1232_v0, 14)
		(_nw215).ArraySet1(_1232_v0, 15)
		(_nw215).ArraySet1((Companion_D1_.Create_DC4_(_dafny.IntOfInt64(51), p2, _1232_v0)).Dtor_cf6(), 16)
		(_nw215).ArraySet1(_1232_v0, 17)
		(_nw215).ArraySet1(_1232_v0, 18)
		_1237_v5 = _nw215
		var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))
		_ = _index166
		(_1237_v5).ArraySet1(_1232_v0, (_index166).Int())
		var _1238_v6 _dafny.Sequence
		_ = _1238_v6
		_1238_v6 = _dafny.SeqOf(_1234_v2)
		var _1239_v7 _dafny.MultiSet
		_ = _1239_v7
		_1239_v7 = _dafny.MultiSetOf(_1233_v1)
		var _1240_v8 _dafny.Map
		_ = _1240_v8
		_1240_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_1239_v7).Contains(_1233_v1) {
				return (_1239_v7).Multiplicity(_1233_v1)
			}
			return p1
		})(), p2)
		var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))
		_ = _index167
		var _rhs173 _dafny.Array = _1232_v0
		_ = _rhs173
		var _rhs174 bool = Companion_Default___.Fm0(!((p1).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1233_v1)).Cardinality()) != 0), _1233_v1, !(_1240_v8).Equals(_1240_v8), globalState)
		_ = _rhs174
		var _rhs175 _dafny.Sequence = _1238_v6
		_ = _rhs175
		var _lhs103 _dafny.Array = _1237_v5
		_ = _lhs103
		var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))
		_ = _lhs104
		var _lhs105 *GlobalState = globalState
		_ = _lhs105
		(_lhs103).ArraySet1(_rhs173, (_lhs104).Int())
		_lhs105.F15 = _rhs174
		_1238_v6 = _rhs175
		if !((_this).F19()) {
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))
			_ = _index168
			(_1237_v5).ArraySet1(_dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int())), (_index168).Int())
			var _1241_v9 _dafny.MultiSet
			_ = _1241_v9
			_1241_v9 = _dafny.MultiSetOf(p1, _dafny.One)
			(globalState).F11 = (_1241_v9).IsDisjointFrom(_1241_v9)
			var _1242_v10 _dafny.Int
			_ = _1242_v10
			_1242_v10 = _dafny.IntOfInt64(938)
			var _1243_v11 _dafny.Sequence
			_ = _1243_v11
			_1243_v11 = _dafny.SeqOf(p2)
			_1242_v10 = Companion_Default___.Fm5(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1243_v11).Cardinality()), p1), globalState)
			_1234_v2 = _1234_v2
			var _1244_v12 _dafny.Map
			_ = _1244_v12
			_1244_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1232_v0)
			var _1245_v13 _dafny.Map
			_ = _1245_v13
			_1245_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
			_1244_v12 = (_1244_v12).Update(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1243_v11, (Companion_Default___.SafeIndex((_1245_v13).Cardinality(), _dafny.IntOfUint32((_1243_v11).Cardinality()))).Uint32(), p0)).Cardinality()), p3), _dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int())))
		} else {
			if (_this).F19() {
				var _1246_v14 _dafny.Sequence
				_ = _1246_v14
				_1246_v14 = _dafny.SeqOf(p1, p1)
				var _1247_v15 _dafny.Array
				_ = _1247_v15
				var _nwElement0_45 _dafny.Int = p3
				_ = _nwElement0_45
				var _nw216 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(11))
				_ = _nw216
				(_nw216).ArraySet1(_nwElement0_45, 0)
				(_nw216).ArraySet1((_1246_v14).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1246_v14).Cardinality()))).Uint32()).(_dafny.Int), 1)
				(_nw216).ArraySet1(p1, 2)
				(_nw216).ArraySet1(p3, 3)
				(_nw216).ArraySet1(p3, 4)
				(_nw216).ArraySet1(p3, 5)
				(_nw216).ArraySet1(p1, 6)
				(_nw216).ArraySet1(p3, 7)
				(_nw216).ArraySet1(p3, 8)
				(_nw216).ArraySet1(p1, 9)
				(_nw216).ArraySet1(_dafny.IntOfUint32((_1238_v6).Cardinality()), 10)
				_1247_v15 = _nw216
				var _1248_v16 _dafny.Map
				_ = _1248_v16
				_1248_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2) || ((_this).F19()), _1247_v15)
				_1248_v16 = (_1248_v16).Merge((_1248_v16).Merge(_1248_v16))
				var _1249_v17 _dafny.Int
				_ = _1249_v17
				_1249_v17 = _dafny.IntOfInt64(919)
				var _1250_v18 _dafny.MultiSet
				_ = _1250_v18
				_1250_v18 = _dafny.MultiSetOf(_1237_v5)
				_1249_v17 = (func() _dafny.Int {
					if (_1250_v18).Contains(_1237_v5) {
						return (_1250_v18).Multiplicity(_1237_v5)
					}
					return p3
				})()
				var _1251_v19 _dafny.Map
				_ = _1251_v19
				_1251_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_dafny.Zero).Minus(p1))
				_1251_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1249_v17, Companion_Default___.SafeModuloInt(p1, _1249_v17))
				var _1252_v20 *C3
				_ = _1252_v20
				var _nw217 *C3 = New_C3_()
				_ = _nw217
				_nw217.Ctor__()
				_1252_v20 = _nw217
				var _1253_v21 _dafny.Sequence
				_ = _1253_v21
				_1253_v21 = _dafny.SeqOf(p0, p2)
				var _rhs176 bool = !(p2)
				_ = _rhs176
				var _rhs177 bool = !((_this).F19()) || (p0)
				_ = _rhs177
				var _rhs178 bool = ((func() _dafny.Int {
					if (_1251_v19).Contains((_dafny.MultiSetFromSeq(_1253_v21)).Cardinality()) {
						return (_1251_v19).Get((_dafny.MultiSetFromSeq(_1253_v21)).Cardinality()).(_dafny.Int)
					}
					return _1249_v17
				})()).Cmp(p3) > 0
				_ = _rhs178
				var _rhs179 bool = p2
				_ = _rhs179
				var _rhs180 _dafny.Int = p1
				_ = _rhs180
				var _lhs106 *GlobalState = globalState
				_ = _lhs106
				var _lhs107 *GlobalState = globalState
				_ = _lhs107
				var _lhs108 *GlobalState = globalState
				_ = _lhs108
				r1 = _rhs176
				_lhs106.F15 = _rhs177
				_lhs107.F17 = _rhs178
				_lhs108.F17 = _rhs179
				_1249_v17 = _rhs180
			} else {
				var _1254_v22 _dafny.Array
				_ = _1254_v22
				var _nw218 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.One)
				_ = _nw218
				_1254_v22 = _nw218
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_1254_v22), 0))
				_ = _index169
				(_1254_v22).ArraySet1(_dafny.MultiSetOf(p0), (_index169).Int())
				var _1255_v23 _dafny.MultiSet
				_ = _1255_v23
				_1255_v23 = _dafny.MultiSetOf(p2)
				var _1256_v24 D10
				_ = _1256_v24
				_1256_v24 = Companion_D10_.Create_DC27_(_1255_v23)
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_1254_v22), 0))
				_ = _index170
				(_1254_v22).ArraySet1((func() _dafny.MultiSet {
					if (_this).F19() {
						return (_1256_v24).Dtor_cf48()
					}
					return _1255_v23
				})(), (_index170).Int())
				var _1257_v25 _dafny.Map
				_ = _1257_v25
				_1257_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1234_v2).Cardinality()), (_dafny.Zero).Minus((_1240_v8).Cardinality())))
				var _1258_v26 _dafny.MultiSet
				_ = _1258_v26
				_1258_v26 = _dafny.MultiSetOf(p1)
				_1257_v25 = (_1257_v25).Update(!(((_1258_v26).Cardinality()).Cmp(p3) <= 0), (_dafny.Zero).Minus((_1257_v25).Cardinality()))
				_1240_v8 = (_1240_v8).Update(p1, p2)
				r0 = (_this).F19()
				var _1259_v27 _dafny.Array
				_ = _1259_v27
				var _nw219 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
				_ = _nw219
				_1259_v27 = _nw219
				_1259_v27 = _1259_v27
			}
			var _1260_v28 _dafny.Map
			_ = _1260_v28
			_1260_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1233_v1, p3)
			var _1261_v29 _dafny.Set
			_ = _1261_v29
			_1261_v29 = _dafny.SetOf(Companion_Default___.Fm9(p1, p3, p2, globalState))
			var _1262_v30 D4
			_ = _1262_v30
			_1262_v30 = Companion_D4_.Create_DC14_(p0, (func() _dafny.Int {
				if (_1260_v28).Contains(_1233_v1) {
					return (_1260_v28).Get(_1233_v1).(_dafny.Int)
				}
				return p3
			})(), _1261_v29, _1233_v1)
			var _1263_v31 D6
			_ = _1263_v31
			_1263_v31 = Companion_D6_.Create_DC20_(Companion_Default___.Fm0(p0, (_1262_v30).Dtor_cf26(), p0, globalState), p1, (_this).F19())
			if !((func(_pat_let23_0 D6) D6 {
				return func(_1264_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let24_0 _dafny.Int) D6 {
						return func(_1265_dt__update_hcf37_h0 _dafny.Int) D6 {
							return Companion_D6_.Create_DC20_((_1264_dt__update__tmp_h0).Dtor_cf36(), _1265_dt__update_hcf37_h0, (_1264_dt__update__tmp_h0).Dtor_cf38())
						}(_pat_let24_0)
					}(_dafny.IntOfInt64(115))
				}(_pat_let23_0)
			}(_1263_v31)).Dtor_cf38()) {
				var _1266_v32 _dafny.Int
				_ = _1266_v32
				_1266_v32 = _dafny.IntOfInt64(203)
				var _1267_v33 _dafny.Sequence
				_ = _1267_v33
				_1267_v33 = _dafny.SeqOf(p3)
				_1266_v32 = (_1267_v33).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1267_v33).Cardinality()))).Uint32()).(_dafny.Int)
				r0 = p2
				(globalState).F15 = !(p2)
				var _1268_v34 _dafny.Array
				_ = _1268_v34
				var _nw220 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw220
				_1268_v34 = _nw220
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1268_v34), 0))
				_ = _index171
				(_1268_v34).ArraySet1(p3, (_index171).Int())
				var _1269_v35 D1
				_ = _1269_v35
				_1269_v35 = Companion_D1_.Create_DC4_((_dafny.Zero).Minus(p1), (_this).F19(), _1232_v0)
				var _1270_v36 _dafny.Set
				_ = _1270_v36
				_1270_v36 = _dafny.SetOf(false, (_this).F19(), (_this).F19(), p2, (_this).F19())
				var _1271_v37 D2
				_ = _1271_v37
				_1271_v37 = Companion_D2_.Create_DC8_((_1269_v35).Dtor_cf6(), _dafny.IntOfInt64(770), (_1270_v36).Cardinality())
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1268_v34), 0))
				_ = _index172
				var _rhs181 _dafny.Int = _1266_v32
				_ = _rhs181
				var _rhs182 _dafny.Array = _dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int()))
				_ = _rhs182
				var _rhs183 _dafny.Array = (_1271_v37).Dtor_cf13()
				_ = _rhs183
				var _lhs109 _dafny.Array = _1268_v34
				_ = _lhs109
				var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1268_v34), 0))
				_ = _lhs110
				(_lhs109).ArraySet1(_rhs181, (_lhs110).Int())
				_1232_v0 = _rhs182
				_1232_v0 = _rhs183
				var _1272_v38 D1
				_ = _1272_v38
				_1272_v38 = Companion_D1_.Create_DC5_(_1234_v2, (_dafny.Zero).Minus(p3), p3)
				var _1273_v39 _dafny.Map
				_ = _1273_v39
				_1273_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _1234_v2)
				var _1274_v40 _dafny.Array
				_ = _1274_v40
				var _nwElement0_46 _dafny.Sequence = _1234_v2
				_ = _nwElement0_46
				var _nw221 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(28))
				_ = _nw221
				(_nw221).ArraySet1(_nwElement0_46, 0)
				(_nw221).ArraySet1(_1234_v2, 1)
				(_nw221).ArraySet1(_1234_v2, 2)
				(_nw221).ArraySet1((func() _dafny.Sequence {
					if true {
						return _1234_v2
					}
					return _1234_v2
				})(), 3)
				(_nw221).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1234_v2, _1234_v2), 4)
				(_nw221).ArraySet1(_1234_v2, 5)
				(_nw221).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1234_v2, (_1272_v38).Dtor_cf7()), 6)
				(_nw221).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1234_v2, _1234_v2), 7)
				(_nw221).ArraySet1((_1238_v6).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1238_v6).Cardinality()))).Uint32()).(_dafny.Sequence), 8)
				(_nw221).ArraySet1((_1272_v38).Dtor_cf7(), 9)
				(_nw221).ArraySet1(_1234_v2, 10)
				(_nw221).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1234_v2, (_1238_v6).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1238_v6).Cardinality()))).Uint32()).(_dafny.Sequence)), 11)
				(_nw221).ArraySet1(_1234_v2, 12)
				(_nw221).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fpbix"), 13)
				(_nw221).ArraySet1(_1234_v2, 14)
				(_nw221).ArraySet1(_1234_v2, 15)
				(_nw221).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(49))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}((func(_1275_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1276_i0 _dafny.Int) _dafny.CodePoint {
						return _1275_v1
					}
				})(_1233_v1))), _1234_v2), 16)
				(_nw221).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(210))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}((func(_1277_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1278_i1 _dafny.Int) _dafny.CodePoint {
						return _1277_v1
					}
				})(_1233_v1))), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(210))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_1279_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1280_i1 _dafny.Int) _dafny.CodePoint {
						return _1279_v1
					}
				})(_1233_v1)))).Cardinality()))).Uint32(), _1233_v1), (func() _dafny.Sequence {
					if (_1273_v39).Contains(p2) {
						return (_1273_v39).Get(p2).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("noveiuir")
				})()), 17)
				(_nw221).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm16(!(p2), (_1268_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1268_v34), 0))).Int()).(_dafny.Int), globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(129))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}((func(_1281_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1282_i2 _dafny.Int) _dafny.CodePoint {
						return _1281_v1
					}
				})(_1233_v1)))), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm16(!(p2), (_1268_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1268_v34), 0))).Int()).(_dafny.Int), globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(129))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}((func(_1283_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1284_i2 _dafny.Int) _dafny.CodePoint {
						return _1283_v1
					}
				})(_1233_v1))))).Cardinality()))).Uint32(), _1233_v1), 18)
				(_nw221).ArraySet1((func() _dafny.Sequence {
					if p2 {
						return _1234_v2
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(898))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg55 _dafny.Int) interface{} {
							return coer55(arg55)
						}
					}((func(_1285_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1286_i3 _dafny.Int) _dafny.CodePoint {
							return _1285_v1
						}
					})(_1233_v1)))
				})(), 19)
				(_nw221).ArraySet1(_1234_v2, 20)
				(_nw221).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vafiqo"), 21)
				(_nw221).ArraySet1(_1234_v2, 22)
				(_nw221).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-200))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}((func(_1287_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1288_i4 _dafny.Int) _dafny.CodePoint {
						return _1287_v1
					}
				})(_1233_v1))), 23)
				(_nw221).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1234_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(717))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}((func(_1289_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1290_i5 _dafny.Int) _dafny.CodePoint {
						return _1289_v1
					}
				})(_1233_v1)))), 24)
				(_nw221).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1234_v2, _1234_v2), 25)
				(_nw221).ArraySet1(_1234_v2, 26)
				(_nw221).ArraySet1(_1234_v2, 27)
				_1274_v40 = _nw221
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1274_v40), 0))
				_ = _index173
				(_1274_v40).ArraySet1(_1234_v2, (_index173).Int())
				var _1291_v41 _dafny.MultiSet
				_ = _1291_v41
				_1291_v41 = _dafny.MultiSetOf(p2, ((_1268_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1268_v34), 0))).Int()).(_dafny.Int)).Cmp(_1266_v32) >= 0, p2)
				var _1292_v42 *C0
				_ = _1292_v42
				var _nw222 *C0 = New_C0_()
				_ = _nw222
				_nw222.Ctor__(p0)
				_1292_v42 = _nw222
				var _1293_v43 D9
				_ = _1293_v43
				_1293_v43 = Companion_D9_.Create_DC25_(p3, p0, p1, _1292_v42)
				var _1294_v44 _dafny.MultiSet
				_ = _1294_v44
				_1294_v44 = _dafny.MultiSetOf((_1293_v43).Dtor_cf45(), _dafny.IntOfUint32((Companion_Default___.Fm16((_1292_v42).F23(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-311)), globalState)).Cardinality()))
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1274_v40), 0))
				_ = _index174
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1268_v34), 0))
				_ = _index175
				var _rhs184 _dafny.Set = _1270_v36
				_ = _rhs184
				var _rhs185 _dafny.Sequence = _1234_v2
				_ = _rhs185
				var _rhs186 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_1294_v44).Cardinality(), _dafny.IntOfUint32((_1234_v2).Cardinality())), (_1261_v29).Cardinality())
				_ = _rhs186
				var _rhs187 _dafny.MultiSet = (_1291_v41).Update((_1270_v36).IsProperSubsetOf(_1270_v36), Companion_Default___.Abs(_1266_v32))
				_ = _rhs187
				var _rhs188 bool = (func() bool {
					if (_1240_v8).Contains(((_1268_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1268_v34), 0))).Int()).(_dafny.Int)).Plus(p1)) {
						return (_1240_v8).Get(((_1268_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1268_v34), 0))).Int()).(_dafny.Int)).Plus(p1)).(bool)
					}
					return p0
				})()
				_ = _rhs188
				var _lhs111 _dafny.Array = _1274_v40
				_ = _lhs111
				var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1274_v40), 0))
				_ = _lhs112
				var _lhs113 _dafny.Array = _1268_v34
				_ = _lhs113
				var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1268_v34), 0))
				_ = _lhs114
				var _lhs115 *GlobalState = globalState
				_ = _lhs115
				_1270_v36 = _rhs184
				(_lhs111).ArraySet1(_rhs185, (_lhs112).Int())
				(_lhs113).ArraySet1(_rhs186, (_lhs114).Int())
				_1291_v41 = _rhs187
				_lhs115.F15 = _rhs188
			} else {
				var _1295_v45 D4
				_ = _1295_v45
				_1295_v45 = Companion_D4_.Create_DC13_(p1)
				var _1296_v46 _dafny.Map
				_ = _1296_v46
				_1296_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _1295_v45)
				_1296_v46 = (_1296_v46).Update(p0, _1295_v45)
				var _1297_v47 *C3
				_ = _1297_v47
				var _nw223 *C3 = New_C3_()
				_ = _nw223
				_nw223.Ctor__()
				_1297_v47 = _nw223
				var _1298_v48 _dafny.Int
				_ = _1298_v48
				_1298_v48 = _dafny.One
				var _1299_v49 T1
				_ = _1299_v49
				var _nw224 *C2 = New_C2_()
				_ = _nw224
				_nw224.Ctor__(Companion_D0_.Create_DC1_(p0), p2)
				_1299_v49 = _nw224
				var _rhs189 _dafny.Int = _1298_v48
				_ = _rhs189
				var _rhs190 T1 = _1299_v49
				_ = _rhs190
				var _rhs191 bool = (_this).F19()
				_ = _rhs191
				var _lhs116 *GlobalState = globalState
				_ = _lhs116
				_1298_v48 = _rhs189
				_1299_v49 = _rhs190
				_lhs116.F17 = _rhs191
				var _1300_v50 _dafny.Set
				_ = _1300_v50
				_1300_v50 = _dafny.SetOf((_this).F19(), p0, p0)
				var _1301_v51 _dafny.Array
				_ = _1301_v51
				var _nwElement0_47 _dafny.Int = _1298_v48
				_ = _nwElement0_47
				var _nw225 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(27))
				_ = _nw225
				(_nw225).ArraySet1(_nwElement0_47, 0)
				(_nw225).ArraySet1(_1298_v48, 1)
				(_nw225).ArraySet1(p1, 2)
				(_nw225).ArraySet1(p1, 3)
				(_nw225).ArraySet1(_1298_v48, 4)
				(_nw225).ArraySet1(_dafny.IntOfInt64(242), 5)
				(_nw225).ArraySet1(p3, 6)
				(_nw225).ArraySet1(p1, 7)
				(_nw225).ArraySet1(_1298_v48, 8)
				(_nw225).ArraySet1(p3, 9)
				(_nw225).ArraySet1(_1298_v48, 10)
				(_nw225).ArraySet1((_dafny.Zero).Minus(_1298_v48), 11)
				(_nw225).ArraySet1(p1, 12)
				(_nw225).ArraySet1(_1298_v48, 13)
				(_nw225).ArraySet1(_1298_v48, 14)
				(_nw225).ArraySet1(p3, 15)
				(_nw225).ArraySet1(_dafny.IntOfUint32((_1234_v2).Cardinality()), 16)
				(_nw225).ArraySet1(_1298_v48, 17)
				(_nw225).ArraySet1((_1300_v50).Cardinality(), 18)
				(_nw225).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p0, false)).Cardinality()), 19)
				(_nw225).ArraySet1(p1, 20)
				(_nw225).ArraySet1(p3, 21)
				(_nw225).ArraySet1(_1298_v48, 22)
				(_nw225).ArraySet1(p1, 23)
				(_nw225).ArraySet1(p3, 24)
				(_nw225).ArraySet1(p1, 25)
				(_nw225).ArraySet1(p3, 26)
				_1301_v51 = _nw225
				var _1302_v52 *C1
				_ = _1302_v52
				var _nw226 *C1 = New_C1_()
				_ = _nw226
				_nw226.Ctor__(_1301_v51)
				_1302_v52 = _nw226
				var _1303_v53 _dafny.Sequence
				_ = _1303_v53
				_1303_v53 = _dafny.SeqOf(_1302_v52)
				_1298_v48 = _dafny.IntOfUint32((_1303_v53).Cardinality())
				_1240_v8 = (_1240_v8).Update(_1298_v48, Companion_Default___.Fm0((_this).F19(), _dafny.CodePoint('a'), (_this).F19(), globalState))
			}
			_1234_v2 = _1234_v2
			_1233_v1 = _dafny.CodePoint('x')
			if ((_dafny.Zero).Minus((_dafny.Zero).Minus(p3))).Cmp(p3) >= 0 {
				var _1304_v54 _dafny.Array
				_ = _1304_v54
				var _len0_39 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_39
				var _nw227 _dafny.Array
				_ = _nw227
				if _len0_39.Cmp(_dafny.Zero) == 0 {
					_nw227 = _dafny.NewArray(_len0_39)
				} else {
					var _init39 func(_dafny.Int) _dafny.CodePoint = (func(_1305_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1306_i6 _dafny.Int) _dafny.CodePoint {
							return _1305_v1
						}
					})(_1233_v1)
					_ = _init39
					var _element0_39 = _init39(_dafny.Zero)
					_ = _element0_39
					_nw227 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
					(_nw227).ArraySet1CodePoint(_element0_39, 0)
					var _nativeLen0_39 = (_len0_39).Int()
					_ = _nativeLen0_39
					for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
						(_nw227).ArraySet1CodePoint(_init39(_dafny.IntOf(_i0_39)), _i0_39)
					}
				}
				_1304_v54 = _nw227
				var _1307_v55 _dafny.Map
				_ = _1307_v55
				_1307_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1304_v54, Companion_Default___.Fm5(p3, globalState))
				var _1308_v56 _dafny.Map
				_ = _1308_v56
				_1308_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p3)
				var _1309_v57 _dafny.Map
				_ = _1309_v57
				_1309_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1308_v56).Contains((_this).F19()) {
						return (_1308_v56).Get((_this).F19()).(_dafny.Int)
					}
					return p1
				})(), _1304_v54)
				var _1310_v58 _dafny.Sequence
				_ = _1310_v58
				_1310_v58 = _dafny.SeqOf(false)
				_1307_v55 = (_1307_v55).Update((func() _dafny.Array {
					if (_1309_v57).Contains(p1) {
						return (_1309_v57).Get(p1).(_dafny.Array)
					}
					return _1304_v54
				})(), _dafny.IntOfUint32((_1310_v58).Cardinality()))
				var _1311_v59 _dafny.Sequence
				_ = _1311_v59
				_1311_v59 = _dafny.SeqOf(p3, p1, p3)
				var _1312_v60 _dafny.Map
				_ = _1312_v60
				_1312_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _dafny.Companion_Sequence_.Update(_1311_v59, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1311_v59).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vbydlk")).Cardinality())))
				var _1313_v61 _dafny.Map
				_ = _1313_v61
				_1313_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(704), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1312_v60).Contains(p2) {
						return (_1312_v60).Get(p2).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(975))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg58 _dafny.Int) interface{} {
							return coer58(arg58)
						}
					}((func(_1314_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_1315_i7 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_1314_v2).Cardinality())
						}
					})(_1234_v2)))
				})()).Cardinality()))
				var _1316_v62 _dafny.Set
				_ = _1316_v62
				_1316_v62 = _dafny.SetOf(p2)
				var _1317_v63 _dafny.Map
				_ = _1317_v63
				_1317_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p2)
				var _1318_v64 D9
				_ = _1318_v64
				_1318_v64 = Companion_D9_.Create_DC24_(Companion_Default___.Fm19(_1311_v59, (func() bool {
					if (_1317_v63).Contains(p0) {
						return (_1317_v63).Get(p0).(bool)
					}
					return (_1310_v58).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1310_v58).Cardinality()))).Uint32()).(bool)
				})(), globalState))
				var _1319_v65 D9
				_ = _1319_v65
				_1319_v65 = Companion_D9_.Create_DC26_(_1318_v64)
				var _1320_v66 _dafny.MultiSet
				_ = _1320_v66
				_1320_v66 = _dafny.MultiSetOf(_1319_v65, _1319_v65)
				var _1321_v67 _dafny.Array
				_ = _1321_v67
				var _nwElement0_48 _dafny.Int = _dafny.IntOfUint32((_1310_v58).Cardinality())
				_ = _nwElement0_48
				var _nw228 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(19))
				_ = _nw228
				(_nw228).ArraySet1(_nwElement0_48, 0)
				(_nw228).ArraySet1(p3, 1)
				(_nw228).ArraySet1(p1, 2)
				(_nw228).ArraySet1(_dafny.IntOfInt64(870), 3)
				(_nw228).ArraySet1(p1, 4)
				(_nw228).ArraySet1(p1, 5)
				(_nw228).ArraySet1((_dafny.Zero).Minus(p3), 6)
				(_nw228).ArraySet1(p3, 7)
				(_nw228).ArraySet1(_dafny.IntOfInt64(738), 8)
				(_nw228).ArraySet1(p3, 9)
				(_nw228).ArraySet1(p3, 10)
				(_nw228).ArraySet1(p1, 11)
				(_nw228).ArraySet1(p3, 12)
				(_nw228).ArraySet1(p3, 13)
				(_nw228).ArraySet1(_dafny.IntOfInt64(-922), 14)
				(_nw228).ArraySet1(p1, 15)
				(_nw228).ArraySet1(p1, 16)
				(_nw228).ArraySet1(p1, 17)
				(_nw228).ArraySet1(p1, 18)
				_1321_v67 = _nw228
				var _1322_v68 _dafny.Map
				_ = _1322_v68
				_1322_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1311_v59).Cardinality()), _1321_v67)
				var _1323_v69 _dafny.Sequence
				_ = _1323_v69
				_1323_v69 = _dafny.SeqOf((func() _dafny.Array {
					if (_1322_v68).Contains(Companion_Default___.Fm9((_1261_v29).Cardinality(), p1, (_this).F19(), globalState)) {
						return (_1322_v68).Get(Companion_Default___.Fm9((_1261_v29).Cardinality(), p1, (_this).F19(), globalState)).(_dafny.Array)
					}
					return _1321_v67
				})())
				var _1324_v70 _dafny.Array
				_ = _1324_v70
				var _nwElement0_49 _dafny.Int = (_dafny.Zero).Minus(((func() _dafny.Int {
					if (_1313_v61).Contains(p1) {
						return (_1313_v61).Get(p1).(_dafny.Int)
					}
					return p3
				})()).Times(_dafny.IntOfInt64(-157)))
				_ = _nwElement0_49
				var _nw229 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(23))
				_ = _nw229
				(_nw229).ArraySet1(_nwElement0_49, 0)
				(_nw229).ArraySet1(p1, 1)
				(_nw229).ArraySet1(((_1261_v29).Cardinality()).Times(p3), 2)
				(_nw229).ArraySet1(p1, 3)
				(_nw229).ArraySet1(Companion_Default___.Fm9(p1, _dafny.IntOfInt64(81), p2, globalState), 4)
				(_nw229).ArraySet1(p1, 5)
				(_nw229).ArraySet1(Companion_Default___.SafeDivisionInt(p3, p3), 6)
				(_nw229).ArraySet1(_dafny.IntOfInt64(-328), 7)
				(_nw229).ArraySet1(p1, 8)
				(_nw229).ArraySet1((func() _dafny.Int {
					if p0 {
						return Companion_Default___.Fm9(p1, p3, p2, globalState)
					}
					return (_1316_v62).Cardinality()
				})(), 9)
				(_nw229).ArraySet1(((func() _dafny.Map {
					if p2 {
						return _1308_v56
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(-174))
				})()).Cardinality(), 10)
				(_nw229).ArraySet1((p1).Minus(p1), 11)
				(_nw229).ArraySet1(p3, 12)
				(_nw229).ArraySet1(p1, 13)
				(_nw229).ArraySet1(_dafny.IntOfInt64(258), 14)
				(_nw229).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1320_v66, p1)).Cardinality()).Times(p1), 15)
				(_nw229).ArraySet1(Companion_Default___.Fm9(p1, (_dafny.Zero).Minus(p3), (_this).F19(), globalState), 16)
				(_nw229).ArraySet1(Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfInt64(589)), 17)
				(_nw229).ArraySet1(p1, 18)
				(_nw229).ArraySet1((_dafny.IntOfInt64(730)).Plus(p1), 19)
				(_nw229).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p1), p3), 20)
				(_nw229).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1323_v69, _1323_v69)).Cardinality()), 21)
				(_nw229).ArraySet1(p1, 22)
				_1324_v70 = _nw229
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_1321_v67), 0))
				_ = _index176
				(_1321_v67).ArraySet1(p1, (_index176).Int())
				var _1325_v71 _dafny.Int
				_ = _1325_v71
				_1325_v71 = _dafny.IntOfInt64(358)
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_1321_v67), 0))
				_ = _index177
				var _rhs192 bool = !(p2)
				_ = _rhs192
				var _rhs193 _dafny.Int = p3
				_ = _rhs193
				var _rhs194 _dafny.Int = _1325_v71
				_ = _rhs194
				var _lhs117 *GlobalState = globalState
				_ = _lhs117
				var _lhs118 _dafny.Array = _1321_v67
				_ = _lhs118
				var _lhs119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_1321_v67), 0))
				_ = _lhs119
				_lhs117.F11 = _rhs192
				(_lhs118).ArraySet1(_rhs193, (_lhs119).Int())
				_1325_v71 = _rhs194
				var _1326_v72 D4
				_ = _1326_v72
				_1326_v72 = Companion_D4_.Create_DC13_((Companion_Default___.Fm20(globalState)).Cardinality())
				var _1327_v73 _dafny.Map
				_ = _1327_v73
				_1327_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1310_v58).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.IntOfUint32((_1310_v58).Cardinality()))).Uint32()).(bool), _1326_v72)
				_1327_v73 = (_1327_v73).Update(true, _1326_v72)
				var _1328_v74 *C1
				_ = _1328_v74
				var _nw230 *C1 = New_C1_()
				_ = _nw230
				_nw230.Ctor__(_1321_v67)
				_1328_v74 = _nw230
				var _1329_v75 _dafny.Sequence
				_ = _1329_v75
				_1329_v75 = _dafny.SeqOf(_1328_v74)
				var _1330_v76 _dafny.MultiSet
				_ = _1330_v76
				_1330_v76 = _dafny.MultiSetOf(!(p2), p2, !(false), false, (_this).F19())
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_1321_v67), 0))
				_ = _index178
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_1321_v67), 0))
				_ = _index179
				var _rhs195 _dafny.Int = (_dafny.Zero).Minus((p3).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1329_v75, _1329_v75)).Cardinality())))
				_ = _rhs195
				var _rhs196 _dafny.Int = ((_dafny.Zero).Minus((func() _dafny.Int {
					if (_1330_v76).Contains(p0) {
						return (_1330_v76).Multiplicity(p0)
					}
					return _1325_v71
				})())).Plus((_1321_v67).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_1321_v67), 0))).Int()).(_dafny.Int))
				_ = _rhs196
				var _lhs120 _dafny.Array = _1321_v67
				_ = _lhs120
				var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_1321_v67), 0))
				_ = _lhs121
				var _lhs122 _dafny.Array = _1321_v67
				_ = _lhs122
				var _lhs123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_1321_v67), 0))
				_ = _lhs123
				(_lhs120).ArraySet1(_rhs195, (_lhs121).Int())
				(_lhs122).ArraySet1(_rhs196, (_lhs123).Int())
				r0 = (p3).Cmp(_1325_v71) == 0
			} else {
				var _1331_v78 _dafny.Map
				_ = _1331_v78
				_1331_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_1234_v2)).Cardinality(), p1)
				var _1332_v79 _dafny.Map
				_ = _1332_v79
				_1332_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int())), func() _dafny.Map {
					var _coll37 = _dafny.NewMapBuilder()
					_ = _coll37
					for _iter39 := _dafny.Iterate((_1331_v78).Keys().Elements()); ; {
						_compr_37, _ok39 := _iter39()
						if !_ok39 {
							break
						}
						var _1333_v77 _dafny.Int
						_1333_v77 = interface{}(_compr_37).(_dafny.Int)
						if (_1331_v78).Contains(_1333_v77) {
							_coll37.Add(Companion_Default___.SafeDivisionInt(_1333_v77, p1), p3)
						}
					}
					return _coll37.ToMap()
				}())
				_1234_v2 = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm16(p2, (_1332_v79).Cardinality(), globalState), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((Companion_Default___.Fm16(p2, (_1332_v79).Cardinality(), globalState)).Cardinality()))).Uint32(), (_1234_v2).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1234_v2).Cardinality()))).Uint32()).(_dafny.CodePoint))
				var _1334_v80 D0
				_ = _1334_v80
				_1334_v80 = Companion_D0_.Create_DC1_(p2)
				var _1335_v81 *C2
				_ = _1335_v81
				var _nw231 *C2 = New_C2_()
				_ = _nw231
				_nw231.Ctor__(_1334_v80, p2)
				_1335_v81 = _nw231
				var _1336_v82 _dafny.Sequence
				_ = _1336_v82
				_1336_v82 = _dafny.SeqOf((_1335_v81).F21())
				var _1337_v83 _dafny.Set
				_ = _1337_v83
				_1337_v83 = _dafny.SetOf(_1336_v82, _1336_v82, _1336_v82)
				var _nw232 *C2 = New_C2_()
				_ = _nw232
				_nw232.Ctor__(_1334_v80, (_1337_v83).IsSubsetOf(_1337_v83))
				_1335_v81 = _nw232
				var _1338_v84 *C0
				_ = _1338_v84
				var _nw233 *C0 = New_C0_()
				_ = _nw233
				_nw233.Ctor__(!((_this).F19()))
				_1338_v84 = _nw233
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1232_v0), 0))
				_ = _index180
				(_1232_v0).ArraySet1((_1335_v81).F21(), (_index180).Int())
				var _1339_v85 _dafny.Int
				_ = _1339_v85
				_1339_v85 = _dafny.IntOfInt64(-626)
				var _1340_v86 _dafny.Set
				_ = _1340_v86
				_1340_v86 = _dafny.SetOf(p0)
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1232_v0), 0))
				_ = _index181
				var _rhs197 bool = ((_1340_v86).Cardinality()).Cmp(Companion_Default___.SafeDivisionInt(p3, _1339_v85)) >= 0
				_ = _rhs197
				var _rhs198 _dafny.Int = p1
				_ = _rhs198
				var _lhs124 _dafny.Array = _1232_v0
				_ = _lhs124
				var _lhs125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1232_v0), 0))
				_ = _lhs125
				(_lhs124).ArraySet1(_rhs197, (_lhs125).Int())
				_1339_v85 = _rhs198
				r0 = ((_dafny.SetOf((_1338_v84).F23())).Union(_1340_v86)).IsSubsetOf((_1340_v86).Difference(_1340_v86))
			}
		}
		var _1341_v87 _dafny.Array
		_ = _1341_v87
		var _nw234 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(26))
		_ = _nw234
		_1341_v87 = _nw234
		for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1341_v87), 0))); ; {
			_guard_loop_2, _ok40 := _iter40()
			if !_ok40 {
				break
			}
			var _1342_i8 _dafny.Int
			_1342_i8 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1342_i8).Sign() != -1) && ((_1342_i8).Cmp(_dafny.ArrayLen((_1341_v87), 0)) < 0)) {
				(_1341_v87).ArraySet1((Companion_D11_.Create_DC29_(_dafny.SetOf((_this).F19()))).Dtor_cf51(), (_1342_i8).Int())
			}
		}
		var _arr0 _dafny.Array = _dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int()))
		_ = _arr0
		var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int()))), 0))
		_ = _index182
		(_arr0).ArraySet1((_dafny.IntOfUint32((_1234_v2).Cardinality())).Cmp(p1) == 0, (_index182).Int())
		var _arr1 _dafny.Array = _dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int()))
		_ = _arr1
		var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int()))), 0))
		_ = _index183
		(_arr1).ArraySet1(p2, (_index183).Int())
		var _hi9 _dafny.Int = (_dafny.IntOfInt64(818)).Plus(_dafny.IntOfInt64(162))
		_ = _hi9
		for _1343_i9 := p3; _1343_i9.Cmp(_hi9) < 0; _1343_i9 = _1343_i9.Plus(_dafny.One) {
			var _1344_v88 _dafny.Array
			_ = _1344_v88
			var _nw235 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
			_ = _nw235
			_1344_v88 = _nw235
			var _1345_v89 _dafny.Sequence
			_ = _1345_v89
			_1345_v89 = _dafny.SeqOf(p0, Companion_Default___.Fm0((_dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int()))), 0))).Int()).(bool), _dafny.CodePoint('a'), p0, globalState))
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1344_v88), 0))
			_ = _index184
			(_1344_v88).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1345_v89).Cardinality()), p3), (_index184).Int())
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1344_v88), 0))
			_ = _index185
			(_1344_v88).ArraySet1((p1).Minus((func() _dafny.Int {
				if p2 {
					return p1
				}
				return _1343_i9
			})()), (_index185).Int())
			var _1346_v90 _dafny.Map
			_ = _1346_v90
			_1346_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1343_i9)
			var _1347_v91 _dafny.Sequence
			_ = _1347_v91
			_1347_v91 = _dafny.SeqOf(p3)
			var _1348_v92 D1
			_ = _1348_v92
			_1348_v92 = Companion_D1_.Create_DC5_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-248))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg59 _dafny.Int) interface{} {
					return coer59(arg59)
				}
			}((func(_1349_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1350_i10 _dafny.Int) _dafny.CodePoint {
					return _1349_v1
				}
			})(_1233_v1))), (_1344_v88).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1344_v88), 0))).Int()).(_dafny.Int), p1)
			var _1351_v93 D3
			_ = _1351_v93
			_1351_v93 = Companion_D3_.Create_DC11_(_1346_v90, Companion_Default___.SafeModuloInt((_1347_v91).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_1348_v92).Dtor_cf7()).Cardinality()), _dafny.IntOfUint32((_1347_v91).Cardinality()))).Uint32()).(_dafny.Int), p1), _1344_v88)
			var _pat_let_tv34 = _1346_v90
			_ = _pat_let_tv34
			_1351_v93 = func(_pat_let25_0 D3) D3 {
				return func(_1352_dt__update__tmp_h1 D3) D3 {
					return func(_pat_let26_0 _dafny.Map) D3 {
						return func(_1353_dt__update_hcf18_h0 _dafny.Map) D3 {
							return Companion_D3_.Create_DC11_(_1353_dt__update_hcf18_h0, (_1352_dt__update__tmp_h1).Dtor_cf19(), (_1352_dt__update__tmp_h1).Dtor_cf20())
						}(_pat_let26_0)
					}(_pat_let_tv34)
				}(_pat_let25_0)
			}(Companion_D3_.Create_DC11_(_1346_v90, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1344_v88).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1344_v88), 0))).Int()).(_dafny.Int))).Cardinality(), _1344_v88))
			_1347_v91 = _1347_v91
			var _1354_v94 _dafny.Map
			_ = _1354_v94
			_1354_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _1233_v1)
			_1354_v94 = _1354_v94
		}
		var _1355_v95 _dafny.Sequence
		_ = _1355_v95
		_1355_v95 = _dafny.SeqOf((_this).F19())
		var _1356_v96 _dafny.Map
		_ = _1356_v96
		_1356_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1355_v95, p1)
		var _1357_i11 _dafny.Int
		_ = _1357_i11
		_1357_i11 = _dafny.Zero
		{
			for !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p0), p3)).Equals(_1356_v96) {
				{
					if (_1357_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1357_i11 = (_1357_i11).Plus(_dafny.One)
					_1355_v95 = _1355_v95
					var _1358_v97 _dafny.Set
					_ = _1358_v97
					_1358_v97 = _dafny.SetOf(p0)
					var _1359_v98 _dafny.MultiSet
					_ = _1359_v98
					_1359_v98 = _dafny.MultiSetOf((_this).F19())
					var _1360_v99 _dafny.Array
					_ = _1360_v99
					var _nwElement0_50 _dafny.Int = p1
					_ = _nwElement0_50
					var _nw236 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(9))
					_ = _nw236
					(_nw236).ArraySet1(_nwElement0_50, 0)
					(_nw236).ArraySet1(p1, 1)
					(_nw236).ArraySet1((_1239_v7).Cardinality(), 2)
					(_nw236).ArraySet1(Companion_Default___.Fm5(p1, globalState), 3)
					(_nw236).ArraySet1(p1, 4)
					(_nw236).ArraySet1((p3).Plus(Companion_Default___.Fm9((_1358_v97).Cardinality(), p1, (_this).F19(), globalState)), 5)
					(_nw236).ArraySet1(((_dafny.MultiSetOf(p0)).Union(_1359_v98)).Cardinality(), 6)
					(_nw236).ArraySet1(p3, 7)
					(_nw236).ArraySet1(p1, 8)
					_1360_v99 = _nw236
					var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1360_v99), 0))
					_ = _index186
					(_1360_v99).ArraySet1((_dafny.IntOfInt64(655)).Minus(p3), (_index186).Int())
					var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_1360_v99), 0))
					_ = _index187
					(_1360_v99).ArraySet1(p3, (_index187).Int())
					var _1361_v100 D5
					_ = _1361_v100
					_1361_v100 = Companion_D5_.Create_DC17_(p2, Companion_D2_.Create_DC8_(_dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int())), p1, _dafny.IntOfInt64(-273)), true, (_this).F19(), p1)
					var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1360_v99), 0))
					_ = _index188
					var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_1360_v99), 0))
					_ = _index189
					var _rhs199 bool = !((_dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_dafny.ArrayCastTo((_1237_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_1237_v5), 0))).Int()))), 0))).Int()).(bool))
					_ = _rhs199
					var _rhs200 _dafny.Int = (_1361_v100).Dtor_cf33()
					_ = _rhs200
					var _rhs201 _dafny.Int = (((_dafny.Zero).Minus(_dafny.IntOfUint32((_1355_v95).Cardinality()))).Times(p1)).Minus(p1)
					_ = _rhs201
					var _lhs126 _dafny.Array = _1360_v99
					_ = _lhs126
					var _lhs127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1360_v99), 0))
					_ = _lhs127
					var _lhs128 _dafny.Array = _1360_v99
					_ = _lhs128
					var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_1360_v99), 0))
					_ = _lhs129
					r1 = _rhs199
					(_lhs126).ArraySet1(_rhs200, (_lhs127).Int())
					(_lhs128).ArraySet1(_rhs201, (_lhs129).Int())
					var _1362_v101 *C3
					_ = _1362_v101
					var _nw237 *C3 = New_C3_()
					_ = _nw237
					_nw237.Ctor__()
					_1362_v101 = _nw237
					var _1363_v102 _dafny.Sequence
					_ = _1363_v102
					_1363_v102 = _dafny.SeqOf(_1360_v99, _1360_v99)
					_1363_v102 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1363_v102, _1363_v102), _1363_v102)
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		r0 = (_this).F19()
		var _1364_v103 *C3
		_ = _1364_v103
		var _nw238 *C3 = New_C3_()
		_ = _nw238
		_nw238.Ctor__()
		_1364_v103 = _nw238
		var _1365_v104 _dafny.Sequence
		_ = _1365_v104
		_1365_v104 = _dafny.SeqOf(_1364_v103)
		r1 = Companion_Default___.Fm0((_dafny.MultiSetOf(_1364_v103)).IsSubsetOf(_dafny.MultiSetFromSeq(_1365_v104)), (func() _dafny.CodePoint {
			if true {
				return _1233_v1
			}
			return Companion_Default___.Fm12(p3, _1234_v2, _1234_v2, globalState)
		})(), p2, globalState)
		return r0, r1
	}
}
func (_this *C4) M2(globalState *GlobalState) {
	{
		var _1366_v0 _dafny.MultiSet
		_ = _1366_v0
		_1366_v0 = _dafny.MultiSetOf((_this).F19())
		var _1367_v1 _dafny.Int
		_ = _1367_v1
		_1367_v1 = _dafny.IntOfInt64(931)
		var _1368_v2 D10
		_ = _1368_v2
		_1368_v2 = Companion_D10_.Create_DC27_((_1366_v0).Union((_1366_v0).Update(!((_this).F19()), Companion_Default___.Abs(_1367_v1))))
		var _1369_v3 _dafny.Sequence
		_ = _1369_v3
		_1369_v3 = _dafny.UnicodeSeqOfUtf8Bytes("yrd")
		var _1370_v4 _dafny.CodePoint
		_ = _1370_v4
		_1370_v4 = _dafny.CodePoint('s')
		var _1371_v5 _dafny.Map
		_ = _1371_v5
		_1371_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("wbgn"), _1370_v4)
		var _1372_v6 _dafny.Sequence
		_ = _1372_v6
		_1372_v6 = _dafny.SeqOf(!(Companion_Default___.Fm0((_this).F19(), _1370_v4, (_this).F19(), globalState)))
		var _1373_v7 _dafny.Array
		_ = _1373_v7
		var _len0_40 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_40
		var _nw239 _dafny.Array
		_ = _nw239
		if _len0_40.Cmp(_dafny.Zero) == 0 {
			_nw239 = _dafny.NewArray(_len0_40)
		} else {
			var _init40 func(_dafny.Int) _dafny.Int = (func(_1374_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1375_i1 _dafny.Int) _dafny.Int {
					return (_1375_i1).Plus(_1374_v1)
				}
			})(_1367_v1)
			_ = _init40
			var _element0_40 = _init40(_dafny.Zero)
			_ = _element0_40
			_nw239 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
			(_nw239).ArraySet1(_element0_40, 0)
			var _nativeLen0_40 = (_len0_40).Int()
			_ = _nativeLen0_40
			for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
				(_nw239).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
			}
		}
		_1373_v7 = _nw239
		var _1376_v8 T0
		_ = _1376_v8
		var _nw240 *C1 = New_C1_()
		_ = _nw240
		_nw240.Ctor__(_1373_v7)
		_1376_v8 = _nw240
		var _1377_v9 _dafny.Sequence
		_ = _1377_v9
		_1377_v9 = _dafny.SeqOf(_1376_v8, _1376_v8)
		var _1378_v10 _dafny.Set
		_ = _1378_v10
		_1378_v10 = _dafny.SetOf((_this).F19())
		var _1379_v11 _dafny.Map
		_ = _1379_v11
		_1379_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1367_v1, _1378_v10)
		var _1380_v12 _dafny.Array
		_ = _1380_v12
		var _nwElement0_51 _dafny.CodePoint = _1370_v4
		_ = _nwElement0_51
		var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(20))
		_ = _nw241
		(_nw241).ArraySet1CodePoint(_nwElement0_51, 0)
		(_nw241).ArraySet1CodePoint(_1370_v4, 1)
		(_nw241).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_this).F19() {
				return (func() _dafny.CodePoint {
					if (_1371_v5).Contains(_dafny.UnicodeSeqOfUtf8Bytes("apnh")) {
						return (_1371_v5).Get(_dafny.UnicodeSeqOfUtf8Bytes("apnh")).(_dafny.CodePoint)
					}
					return _1370_v4
				})()
			}
			return _1370_v4
		})(), 2)
		(_nw241).ArraySet1CodePoint(_1370_v4, 3)
		(_nw241).ArraySet1CodePoint(_1370_v4, 4)
		(_nw241).ArraySet1CodePoint(_1370_v4, 5)
		(_nw241).ArraySet1CodePoint(_1370_v4, 6)
		(_nw241).ArraySet1CodePoint(Companion_Default___.Fm12(_dafny.IntOfUint32((_1372_v6).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(223))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg60 _dafny.Int) interface{} {
				return coer60(arg60)
			}
		}(func(_1381_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		})), _dafny.UnicodeSeqOfUtf8Bytes("jxjyg"), globalState), 7)
		(_nw241).ArraySet1CodePoint(_1370_v4, 8)
		(_nw241).ArraySet1CodePoint((_1369_v3).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm5(_dafny.IntOfUint32((_1377_v9).Cardinality()), globalState), _dafny.IntOfUint32((_1369_v3).Cardinality()))).Uint32()).(_dafny.CodePoint), 9)
		(_nw241).ArraySet1CodePoint(_1370_v4, 10)
		(_nw241).ArraySet1CodePoint(_1370_v4, 11)
		(_nw241).ArraySet1CodePoint(_1370_v4, 12)
		(_nw241).ArraySet1CodePoint(Companion_Default___.Fm12(_1367_v1, _dafny.UnicodeSeqOfUtf8Bytes("vnd"), _1369_v3, globalState), 13)
		(_nw241).ArraySet1CodePoint(_1370_v4, 14)
		(_nw241).ArraySet1CodePoint(_1370_v4, 15)
		(_nw241).ArraySet1CodePoint(_1370_v4, 16)
		(_nw241).ArraySet1CodePoint(Companion_Default___.Fm12(_1367_v1, _dafny.UnicodeSeqOfUtf8Bytes("op"), _1369_v3, globalState), 17)
		(_nw241).ArraySet1CodePoint(_1370_v4, 18)
		(_nw241).ArraySet1CodePoint(Companion_Default___.Fm12((_1379_v11).Cardinality(), _dafny.UnicodeSeqOfUtf8Bytes("hdoqfle"), _1369_v3, globalState), 19)
		_1380_v12 = _nw241
		var _rhs202 D10 = Companion_D10_.Create_DC27_(_1366_v0)
		_ = _rhs202
		var _rhs203 _dafny.Int = (func() _dafny.Int {
			if (_1366_v0).Contains((_this).F19()) {
				return (_1366_v0).Multiplicity((_this).F19())
			}
			return Companion_Default___.Fm5(_dafny.IntOfUint32((_1372_v6).Cardinality()), globalState)
		})()
		_ = _rhs203
		var _rhs204 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1369_v3, _1369_v3)
		_ = _rhs204
		var _rhs205 _dafny.Array = _1380_v12
		_ = _rhs205
		_1368_v2 = _rhs202
		_1367_v1 = _rhs203
		_1369_v3 = _rhs204
		_1380_v12 = _rhs205
		var _1382_v13 D10
		_ = _1382_v13
		_1382_v13 = Companion_D10_.Create_DC28_((_this).F19(), _1367_v1)
		var _1383_v14 _dafny.Sequence
		_ = _1383_v14
		_1383_v14 = _dafny.SeqOf((_dafny.Zero).Minus(_1367_v1), (_1382_v13).Dtor_cf50(), ((_dafny.Zero).Minus(_1367_v1)).Times(_dafny.IntOfUint32((_1369_v3).Cardinality())), _1367_v1)
		_1383_v14 = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_this).F19() {
				return _1383_v14
			}
			return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1383_v14, (Companion_Default___.SafeIndex(_1367_v1, _dafny.IntOfUint32((_1383_v14).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1369_v3).Cardinality())), (Companion_Default___.SafeIndex(_1367_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1383_v14, (Companion_Default___.SafeIndex(_1367_v1, _dafny.IntOfUint32((_1383_v14).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1369_v3).Cardinality()))).Cardinality()))).Uint32(), (Companion_Default___.Fm29(_1370_v4, _dafny.IntOfInt64(420), globalState)).Cardinality()), _1383_v14)
		})(), (Companion_Default___.SafeIndex(_1367_v1, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_this).F19() {
				return _1383_v14
			}
			return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1383_v14, (Companion_Default___.SafeIndex(_1367_v1, _dafny.IntOfUint32((_1383_v14).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1369_v3).Cardinality())), (Companion_Default___.SafeIndex(_1367_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1383_v14, (Companion_Default___.SafeIndex(_1367_v1, _dafny.IntOfUint32((_1383_v14).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1369_v3).Cardinality()))).Cardinality()))).Uint32(), (Companion_Default___.Fm29(_1370_v4, _dafny.IntOfInt64(420), globalState)).Cardinality()), _1383_v14)
		})()).Cardinality()))).Uint32(), _1367_v1)
		var _1384_v15 _dafny.MultiSet
		_ = _1384_v15
		_1384_v15 = _dafny.MultiSetOf(_1367_v1)
		var _1385_v16 D2
		_ = _1385_v16
		_1385_v16 = Companion_D2_.Create_DC7_(_1384_v15)
		var _pat_let_tv35 = _1367_v1
		_ = _pat_let_tv35
		_1367_v1 = (_dafny.Zero).Minus(func(_source14 D2) _dafny.Int {
			if _source14.Is_DC8() {
				var _1386___mcc_h0 _dafny.Array = _source14.Get_().(D2_DC8).Cf13
				_ = _1386___mcc_h0
				var _1387___mcc_h1 _dafny.Int = _source14.Get_().(D2_DC8).Cf14
				_ = _1387___mcc_h1
				var _1388___mcc_h2 _dafny.Int = _source14.Get_().(D2_DC8).Cf15
				_ = _1388___mcc_h2
				var _1389_cf15 _dafny.Int = _1388___mcc_h2
				_ = _1389_cf15
				var _1390_cf14 _dafny.Int = _1387___mcc_h1
				_ = _1390_cf14
				var _1391_cf13 _dafny.Array = _1386___mcc_h0
				_ = _1391_cf13
				if (_this).F19() {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yxsw")).Cardinality())
				} else {
					return _1390_cf14
				}
			} else {
				var _1392___mcc_h3 _dafny.MultiSet = _source14.Get_().(D2_DC7).Cf12
				_ = _1392___mcc_h3
				var _1393_cf12 _dafny.MultiSet = _1392___mcc_h3
				_ = _1393_cf12
				return _pat_let_tv35
			}
		}(_1385_v16))
		var _1394_i2 _dafny.Int
		_ = _1394_i2
		_1394_i2 = _dafny.Zero
		{
			for false {
				{
					if (_1394_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1394_i2 = (_1394_i2).Plus(_dafny.One)
					var _rhs206 bool = ((_this).F19()) && ((_this).F19())
					_ = _rhs206
					var _rhs207 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt(_1367_v1, _1367_v1)).Plus(_1367_v1)))
					_ = _rhs207
					var _rhs208 bool = !((_this).F19()) || (!(Companion_Default___.Fm0((_this).F19(), _dafny.CodePoint('l'), (_this).F19(), globalState)))
					_ = _rhs208
					var _rhs209 bool = (_this).F19()
					_ = _rhs209
					var _rhs210 bool = (_this).F19()
					_ = _rhs210
					var _lhs130 *GlobalState = globalState
					_ = _lhs130
					var _lhs131 *GlobalState = globalState
					_ = _lhs131
					var _lhs132 *GlobalState = globalState
					_ = _lhs132
					var _lhs133 *GlobalState = globalState
					_ = _lhs133
					_lhs130.F11 = _rhs206
					_1367_v1 = _rhs207
					_lhs131.F11 = _rhs208
					_lhs132.F17 = _rhs209
					_lhs133.F15 = _rhs210
					var _1395_v17 *C3
					_ = _1395_v17
					var _nw242 *C3 = New_C3_()
					_ = _nw242
					_nw242.Ctor__()
					_1395_v17 = _nw242
					_1366_v0 = _dafny.MultiSetOf(((_this).F19()) == ((_this).F19()))
					var _1396_v18 _dafny.Array
					_ = _1396_v18
					var _len0_41 _dafny.Int = _dafny.IntOfInt64(11)
					_ = _len0_41
					var _nw243 _dafny.Array
					_ = _nw243
					if _len0_41.Cmp(_dafny.Zero) == 0 {
						_nw243 = _dafny.NewArray(_len0_41)
					} else {
						var _init41 func(_dafny.Int) bool = func(_1397_i3 _dafny.Int) bool {
							return (_this).F19()
						}
						_ = _init41
						var _element0_41 = _init41(_dafny.Zero)
						_ = _element0_41
						_nw243 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
						(_nw243).ArraySet1(_element0_41, 0)
						var _nativeLen0_41 = (_len0_41).Int()
						_ = _nativeLen0_41
						for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
							(_nw243).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
						}
					}
					_1396_v18 = _nw243
					var _1398_v19 _dafny.Map
					_ = _1398_v19
					_1398_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _1367_v1)
					var _1399_v20 _dafny.Map
					_ = _1399_v20
					_1399_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
						if (_1398_v19).Contains((_this).F19()) {
							return (_1398_v19).Get((_this).F19()).(_dafny.Int)
						}
						return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1369_v3, (Companion_Default___.SafeIndex(_1367_v1, _dafny.IntOfUint32((_1369_v3).Cardinality()))).Uint32(), _1370_v4)).Cardinality())
					})(), false)
					var _1400_v21 _dafny.Set
					_ = _1400_v21
					_1400_v21 = _dafny.SetOf(_1367_v1)
					var _1401_v22 D4
					_ = _1401_v22
					_1401_v22 = Companion_D4_.Create_DC14_((_this).F19(), _1367_v1, _1400_v21, _1370_v4)
					var _1402_v23 _dafny.MultiSet
					_ = _1402_v23
					_1402_v23 = _dafny.MultiSetOf((_1401_v22).Dtor_cf26())
					var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1396_v18), 0))
					_ = _index190
					(_1396_v18).ArraySet1((func() bool {
						if (_1399_v20).Contains((_1402_v23).Cardinality()) {
							return (_1399_v20).Get((_1402_v23).Cardinality()).(bool)
						}
						return (_this).F19()
					})(), (_index190).Int())
					var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1396_v18), 0))
					_ = _index191
					(_1396_v18).ArraySet1((_this).Fm4((_this).F19(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1369_v3, _1398_v19), _1370_v4, _1367_v1, globalState), (_index191).Int())
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		(globalState).F17 = !((_this).F19())
		var _1403_v24 D9
		_ = _1403_v24
		_1403_v24 = Companion_D9_.Create_DC24_(_1383_v14)
		var _1404_v25 _dafny.Map
		_ = _1404_v25
		_1404_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), Companion_D9_.Create_DC26_(_1403_v24))
		var _1405_v26 _dafny.Array
		_ = _1405_v26
		var _nwElement0_52 _dafny.Map = _1404_v25
		_ = _nwElement0_52
		var _nw244 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(8))
		_ = _nw244
		(_nw244).ArraySet1(_nwElement0_52, 0)
		(_nw244).ArraySet1(_1404_v25, 1)
		(_nw244).ArraySet1(_1404_v25, 2)
		(_nw244).ArraySet1(_1404_v25, 3)
		(_nw244).ArraySet1(_1404_v25, 4)
		(_nw244).ArraySet1(_1404_v25, 5)
		(_nw244).ArraySet1(_1404_v25, 6)
		(_nw244).ArraySet1((_1404_v25).Merge(_1404_v25), 7)
		_1405_v26 = _nw244
		var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_1405_v26), 0))
		_ = _index192
		(_1405_v26).ArraySet1(_1404_v25, (_index192).Int())
		var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_1405_v26), 0))
		_ = _index193
		(_1405_v26).ArraySet1(_1404_v25, (_index193).Int())
	}
}
func (_this *C4) M3(p0 _dafny.Int, p1 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _1406_v0 _dafny.Map
		_ = _1406_v0
		_1406_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), (_this).F19())
		if (func() bool {
			if (_this).F19() {
				return (func() bool {
					if (_1406_v0).Contains((_this).F19()) {
						return (_1406_v0).Get((_this).F19()).(bool)
					}
					return (func() bool {
						if (_1406_v0).Contains(p1) {
							return (_1406_v0).Get(p1).(bool)
						}
						return (_this).F19()
					})()
				})()
			}
			return p1
		})() {
			var _1407_v1 _dafny.Sequence
			_ = _1407_v1
			_1407_v1 = _dafny.SeqOf(p1)
			var _1408_v2 _dafny.Sequence
			_ = _1408_v2
			_1408_v2 = _dafny.UnicodeSeqOfUtf8Bytes("ghvjg")
			var _1409_v3 _dafny.Map
			_ = _1409_v3
			_1409_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_1408_v2).Cardinality()))
			var _1410_v4 _dafny.Map
			_ = _1410_v4
			_1410_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F19()), _1409_v3)
			var _1411_v5 _dafny.Map
			_ = _1411_v5
			_1411_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1408_v2, (func() _dafny.Map {
				if (_1410_v4).Contains((func() bool {
					if (_1406_v0).Contains((_this).F19()) {
						return (_1406_v0).Get((_this).F19()).(bool)
					}
					return (_this).F19()
				})()) {
					return (_1410_v4).Get((func() bool {
						if (_1406_v0).Contains((_this).F19()) {
							return (_1406_v0).Get((_this).F19()).(bool)
						}
						return (_this).F19()
					})()).(_dafny.Map)
				}
				return _1409_v3
			})())
			var _1412_v6 _dafny.CodePoint
			_ = _1412_v6
			_1412_v6 = _dafny.CodePoint('o')
			var _1413_v7 D1
			_ = _1413_v7
			_1413_v7 = Companion_D1_.Create_DC6_(p0, _dafny.IntOfInt64(943))
			var _1414_v8 D11
			_ = _1414_v8
			_1414_v8 = Companion_D11_.Create_DC30_(p1, (_1413_v7).Dtor_cf11(), (_this).F19())
			var _1415_v9 _dafny.Map
			_ = _1415_v9
			_1415_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1407_v1, p1)
			var _1416_v11 _dafny.MultiSet
			_ = _1416_v11
			_1416_v11 = _dafny.MultiSetOf(_1407_v1)
			var _1417_v14 _dafny.MultiSet
			_ = _1417_v14
			_1417_v14 = _dafny.MultiSetOf(true, (_this).F19())
			var _1418_v15 _dafny.Map
			_ = _1418_v15
			_1418_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1417_v14).Cardinality(), (_this).F19())
			var _1419_v16 _dafny.Map
			_ = _1419_v16
			_1419_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1418_v15)
			var _1420_v17 _dafny.Sequence
			_ = _1420_v17
			_1420_v17 = _dafny.SeqOf(((func() _dafny.Map {
				if (_1419_v16).Contains(p0) {
					return (_1419_v16).Get(p0).(_dafny.Map)
				}
				return _1418_v15
			})()).Cardinality(), p0, Companion_Default___.Fm5((_1417_v14).Cardinality(), globalState), (_dafny.Zero).Minus(p0))
			var _1421_v18 _dafny.Sequence
			_ = _1421_v18
			_1421_v18 = _dafny.SeqOf((_1420_v17).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1420_v17).Cardinality()))).Uint32()).(_dafny.Int), p0, p0)
			var _1422_v20 _dafny.Sequence
			_ = _1422_v20
			_1422_v20 = _dafny.SeqOf(_1407_v1)
			var _1423_v21 _dafny.Array
			_ = _1423_v21
			var _nwElement0_53 _dafny.Map = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1407_v1, p1)).Update(_dafny.SeqOf(p1), (_this).Fm4((_this).F19(), _1411_v5, _1412_v6, p0, globalState))).Update(_1407_v1, (_1414_v8).Dtor_cf52())).Update(_dafny.SeqOf((_this).F19()), (_this).F19())
			_ = _nwElement0_53
			var _nw245 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(15))
			_ = _nw245
			(_nw245).ArraySet1(_nwElement0_53, 0)
			(_nw245).ArraySet1(_1415_v9, 1)
			(_nw245).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, p1), true), 2)
			(_nw245).ArraySet1((func() _dafny.Map {
				var _coll38 = _dafny.NewMapBuilder()
				_ = _coll38
				for _iter41 := _dafny.Iterate((_1416_v11).Elements()); ; {
					_compr_38, _ok41 := _iter41()
					if !_ok41 {
						break
					}
					var _1424_v10 _dafny.Sequence
					_1424_v10 = interface{}(_compr_38).(_dafny.Sequence)
					if (_1416_v11).Contains(_1424_v10) {
						_coll38.Add(_1424_v10, !(p1))
					}
				}
				return _coll38.ToMap()
			}()).Merge(_1415_v9), 3)
			(_nw245).ArraySet1((_1415_v9).Merge(_1415_v9), 4)
			(_nw245).ArraySet1(_1415_v9, 5)
			(_nw245).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm27(p0, p0, p0, globalState), p1), 6)
			(_nw245).ArraySet1(_1415_v9, 7)
			(_nw245).ArraySet1(func() _dafny.Map {
				var _coll39 = _dafny.NewMapBuilder()
				_ = _coll39
				for _iter42 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf((_this).F19(), !(p1)))).Elements()); ; {
					_compr_39, _ok42 := _iter42()
					if !_ok42 {
						break
					}
					var _1425_v12 _dafny.Sequence
					_1425_v12 = interface{}(_compr_39).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf((_this).F19(), !(p1))), _1425_v12) {
						_coll39.Add(_1425_v12, (_this).F19())
					}
				}
				return _coll39.ToMap()
			}(), 8)
			(_nw245).ArraySet1(_1415_v9, 9)
			(_nw245).ArraySet1(func() _dafny.Map {
				var _coll40 = _dafny.NewMapBuilder()
				_ = _coll40
				for _iter43 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p1), (_this).F19())).Keys().Elements()); ; {
					_compr_40, _ok43 := _iter43()
					if !_ok43 {
						break
					}
					var _1426_v13 _dafny.Sequence
					_1426_v13 = interface{}(_compr_40).(_dafny.Sequence)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p1), (_this).F19())).Contains(_1426_v13) {
						_coll40.Add(_1426_v13, p1)
					}
				}
				return _coll40.ToMap()
			}(), 10)
			(_nw245).ArraySet1(Companion_Default___.Fm30(_1421_v18, globalState), 11)
			(_nw245).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1407_v1, Companion_Default___.Fm0((_this).F19(), _1412_v6, p1, globalState))).Merge(_1415_v9), 12)
			(_nw245).ArraySet1(func() _dafny.Map {
				var _coll41 = _dafny.NewMapBuilder()
				_ = _coll41
				for _iter44 := _dafny.Iterate((_1422_v20).Elements()); ; {
					_compr_41, _ok44 := _iter44()
					if !_ok44 {
						break
					}
					var _1427_v19 _dafny.Sequence
					_1427_v19 = interface{}(_compr_41).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_1422_v20, _1427_v19) {
						_coll41.Add(_1427_v19, (_this).F19())
					}
				}
				return _coll41.ToMap()
			}(), 13)
			(_nw245).ArraySet1((_1415_v9).Merge(_1415_v9), 14)
			_1423_v21 = _nw245
			var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_1423_v21), 0))
			_ = _index194
			(_1423_v21).ArraySet1(_1415_v9, (_index194).Int())
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_1423_v21), 0))
			_ = _index195
			(_1423_v21).ArraySet1(_1415_v9, (_index195).Int())
			if _dafny.Companion_Sequence_.IsProperPrefixOf(_1407_v1, _1407_v1) {
				var _1428_v22 _dafny.Array
				_ = _1428_v22
				var _nw246 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
				_ = _nw246
				_1428_v22 = _nw246
				var _1429_v23 _dafny.Map
				_ = _1429_v23
				_1429_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1428_v22, (_1420_v17).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-865), _dafny.IntOfUint32((_1420_v17).Cardinality()))).Uint32()).(_dafny.Int))
				_1429_v23 = (_1429_v23).Update(_1428_v22, p0)
				var _1430_v24 _dafny.Int
				_ = _1430_v24
				_1430_v24 = _dafny.IntOfInt64(47)
				var _1431_v25 _dafny.MultiSet
				_ = _1431_v25
				_1431_v25 = _dafny.MultiSetOf(Companion_D10_.Create_DC27_(_1417_v14))
				_1430_v24 = (func() _dafny.Int {
					if (_1431_v25).Contains(Companion_Default___.Fm31(true, (_1421_v18).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1430_v24), _dafny.IntOfUint32((_1421_v18).Cardinality()))).Uint32()).(_dafny.Int), globalState)) {
						return (_1431_v25).Multiplicity(Companion_Default___.Fm31(true, (_1421_v18).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1430_v24), _dafny.IntOfUint32((_1421_v18).Cardinality()))).Uint32()).(_dafny.Int), globalState))
					}
					return _1430_v24
				})()
				var _1432_v26 _dafny.Array
				_ = _1432_v26
				var _nw247 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
				_ = _nw247
				_1432_v26 = _nw247
				var _1433_v27 _dafny.MultiSet
				_ = _1433_v27
				_1433_v27 = _dafny.MultiSetOf(_1432_v26)
				var _1434_v28 _dafny.Map
				_ = _1434_v28
				_1434_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1433_v27, (_dafny.Zero).Minus((Companion_Default___.Fm5(p0, globalState)).Times(_1430_v24)))
				_1434_v28 = (_1434_v28).Update(_1433_v27, _dafny.IntOfInt64(592))
				var _1435_v29 D9
				_ = _1435_v29
				_1435_v29 = Companion_D9_.Create_DC26_(Companion_D9_.Create_DC24_(_1421_v18))
				var _1436_v30 *C0
				_ = _1436_v30
				var _nw248 *C0 = New_C0_()
				_ = _nw248
				_nw248.Ctor__(p1)
				_1436_v30 = _nw248
				var _1437_v31 D9
				_ = _1437_v31
				_1437_v31 = Companion_D9_.Create_DC26_(Companion_D9_.Create_DC26_(Companion_D9_.Create_DC25_(_1430_v24, (_this).F19(), _1430_v24, _1436_v30)))
				var _pat_let_tv36 = _1437_v31
				_ = _pat_let_tv36
				var _1438_v32 _dafny.Array
				_ = _1438_v32
				var _nwElement0_54 D9 = _1435_v29
				_ = _nwElement0_54
				var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(9))
				_ = _nw249
				(_nw249).ArraySet1(_nwElement0_54, 0)
				(_nw249).ArraySet1(Companion_D9_.Create_DC26_(Companion_D9_.Create_DC26_(_1437_v31)), 1)
				(_nw249).ArraySet1(_1435_v29, 2)
				(_nw249).ArraySet1(_1435_v29, 3)
				(_nw249).ArraySet1(Companion_D9_.Create_DC26_(Companion_D9_.Create_DC24_(_1420_v17)), 4)
				(_nw249).ArraySet1((func() D9 {
					if !(p1) {
						return _1435_v29
					}
					return _1435_v29
				})(), 5)
				(_nw249).ArraySet1(_1435_v29, 6)
				(_nw249).ArraySet1(_1435_v29, 7)
				(_nw249).ArraySet1(func(_pat_let27_0 D9) D9 {
					return func(_1439_dt__update__tmp_h0 D9) D9 {
						return func(_pat_let28_0 D9) D9 {
							return func(_1440_dt__update_hcf47_h0 D9) D9 {
								return Companion_D9_.Create_DC26_(_1440_dt__update_hcf47_h0)
							}(_pat_let28_0)
						}(_pat_let_tv36)
					}(_pat_let27_0)
				}(_1435_v29), 8)
				_1438_v32 = _nw249
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_1438_v32), 0))
				_ = _index196
				(_1438_v32).ArraySet1(_1435_v29, (_index196).Int())
				var _1441_v33 D4
				_ = _1441_v33
				_1441_v33 = Companion_D4_.Create_DC14_((_this).F19(), p0, Companion_Default___.Fm18((_1436_v30).F23(), _dafny.IntOfUint32((_1420_v17).Cardinality()), _dafny.IntOfInt64(-821), globalState), _1412_v6)
				var _1442_v34 _dafny.Set
				_ = _1442_v34
				_1442_v34 = _dafny.SetOf((_1420_v17).Select((Companion_Default___.SafeIndex(_1430_v24, _dafny.IntOfUint32((_1420_v17).Cardinality()))).Uint32()).(_dafny.Int))
				var _1443_v35 _dafny.Set
				_ = _1443_v35
				_1443_v35 = _dafny.SetOf((_1441_v33).Dtor_cf25(), _dafny.SetOf(_dafny.IntOfUint32((_1408_v2).Cardinality())), _1442_v34, _1442_v34, _dafny.SetOf((_1414_v8).Dtor_cf53(), p0))
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_1438_v32), 0))
				_ = _index197
				var _rhs211 _dafny.Int = p0
				_ = _rhs211
				var _rhs212 D9 = Companion_D9_.Create_DC26_(_1437_v31)
				_ = _rhs212
				var _rhs213 bool = !((_1443_v35).Equals(_1443_v35))
				_ = _rhs213
				var _lhs134 _dafny.Array = _1438_v32
				_ = _lhs134
				var _lhs135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_1438_v32), 0))
				_ = _lhs135
				var _lhs136 *GlobalState = globalState
				_ = _lhs136
				_1430_v24 = _rhs211
				(_lhs134).ArraySet1(_rhs212, (_lhs135).Int())
				_lhs136.F15 = _rhs213
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_1428_v22), 0))
				_ = _index198
				(_1428_v22).ArraySet1(Companion_Default___.SafeDivisionInt(_1430_v24, _dafny.IntOfUint32((_1407_v1).Cardinality())), (_index198).Int())
				var _1444_v36 _dafny.MultiSet
				_ = _1444_v36
				_1444_v36 = _dafny.MultiSetOf(_dafny.CodePoint('u'))
				var _1445_v37 D4
				_ = _1445_v37
				_1445_v37 = Companion_D4_.Create_DC13_((_dafny.Zero).Minus(_dafny.IntOfUint32((_1408_v2).Cardinality())))
				var _1446_v38 _dafny.Map
				_ = _1446_v38
				_1446_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1445_v37, (_dafny.Zero).Minus(_1430_v24))
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_1428_v22), 0))
				_ = _index199
				var _rhs214 bool = ((_1436_v30).F23()) || ((_1444_v36).IsProperSubsetOf(_1444_v36))
				_ = _rhs214
				var _rhs215 bool = (_1436_v30).F23()
				_ = _rhs215
				var _rhs216 _dafny.Int = (func() _dafny.Int {
					if (_1446_v38).Contains(_1445_v37) {
						return (_1446_v38).Get(_1445_v37).(_dafny.Int)
					}
					return _1430_v24
				})()
				_ = _rhs216
				var _lhs137 *GlobalState = globalState
				_ = _lhs137
				var _lhs138 *GlobalState = globalState
				_ = _lhs138
				var _lhs139 _dafny.Array = _1428_v22
				_ = _lhs139
				var _lhs140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_1428_v22), 0))
				_ = _lhs140
				_lhs137.F17 = _rhs214
				_lhs138.F17 = _rhs215
				(_lhs139).ArraySet1(_rhs216, (_lhs140).Int())
			} else {
				var _1447_v39 _dafny.Map
				_ = _1447_v39
				_1447_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1414_v8, !(p1))
				var _1448_v40 _dafny.Map
				_ = _1448_v40
				_1448_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1407_v1, (_1447_v39).Update(_1414_v8, p1))
				var _1449_v41 _dafny.Array
				_ = _1449_v41
				var _nwElement0_55 bool = !((_this).F19())
				_ = _nwElement0_55
				var _nw250 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(20))
				_ = _nw250
				(_nw250).ArraySet1(_nwElement0_55, 0)
				(_nw250).ArraySet1((_this).F19(), 1)
				(_nw250).ArraySet1((_this).F19(), 2)
				(_nw250).ArraySet1(!((_this).F19()) || (!(!((_this).F19()))), 3)
				(_nw250).ArraySet1((_this).F19(), 4)
				(_nw250).ArraySet1(!(!(_1448_v40).Contains(_1407_v1)), 5)
				(_nw250).ArraySet1(p1, 6)
				(_nw250).ArraySet1((_this).F19(), 7)
				(_nw250).ArraySet1(p1, 8)
				(_nw250).ArraySet1(p1, 9)
				(_nw250).ArraySet1((func() bool {
					if (_1406_v0).Contains(p1) {
						return (_1406_v0).Get(p1).(bool)
					}
					return true
				})(), 10)
				(_nw250).ArraySet1(true, 11)
				(_nw250).ArraySet1(p1, 12)
				(_nw250).ArraySet1((func() bool {
					if (_this).Fm4((_this).Fm4((_this).F19(), _1411_v5, _1412_v6, (_dafny.Zero).Minus(p0), globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1408_v2, _1409_v3), _1412_v6, _dafny.IntOfUint32((_dafny.SeqOf(_1412_v6)).Cardinality()), globalState) {
						return (_this).F19()
					}
					return true
				})(), 13)
				(_nw250).ArraySet1((_this).F19(), 14)
				(_nw250).ArraySet1((_this).F19(), 15)
				(_nw250).ArraySet1((_this).F19(), 16)
				(_nw250).ArraySet1((func() bool {
					if (_1418_v15).Contains(p0) {
						return (_1418_v15).Get(p0).(bool)
					}
					return false
				})(), 17)
				(_nw250).ArraySet1(p1, 18)
				(_nw250).ArraySet1(p1, 19)
				_1449_v41 = _nw250
				var _1450_v42 D0
				_ = _1450_v42
				_1450_v42 = Companion_D0_.Create_DC1_((_this).F19())
				var _1451_v43 T1
				_ = _1451_v43
				var _nw251 *C2 = New_C2_()
				_ = _nw251
				_nw251.Ctor__(_1450_v42, p1)
				_1451_v43 = _nw251
				var _1452_v44 D12
				_ = _1452_v44
				_1452_v44 = Companion_D12_.Create_DC31_(_dafny.SetOf(_1451_v43, _1451_v43))
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_1449_v41), 0))
				_ = _index200
				(_1449_v41).ArraySet1(!((_1452_v44).Dtor_cf55()).Contains(_1451_v43), (_index200).Int())
				var _1453_v45 _dafny.Int
				_ = _1453_v45
				_1453_v45 = _dafny.IntOfInt64(579)
				var _1454_v46 _dafny.Array
				_ = _1454_v46
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_42
				var _nw252 _dafny.Array
				_ = _nw252
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw252 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) _dafny.Int = (func(_1455_v45 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1456_i0 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1456_i0, _1455_v45)
						}
					})(_1453_v45)
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw252 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw252).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw252).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1454_v46 = _nw252
				var _1457_v47 D3
				_ = _1457_v47
				_1457_v47 = Companion_D3_.Create_DC11_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality())), (_dafny.Zero).Minus(_1453_v45), _1454_v46)
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_1449_v41), 0))
				_ = _index201
				var _rhs217 _dafny.Sequence = _1420_v17
				_ = _rhs217
				var _rhs218 bool = (_1407_v1).Select((Companion_Default___.SafeIndex((_1457_v47).Dtor_cf19(), _dafny.IntOfUint32((_1407_v1).Cardinality()))).Uint32()).(bool)
				_ = _rhs218
				var _rhs219 _dafny.Int = p0
				_ = _rhs219
				var _lhs141 _dafny.Array = _1449_v41
				_ = _lhs141
				var _lhs142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_1449_v41), 0))
				_ = _lhs142
				_1420_v17 = _rhs217
				(_lhs141).ArraySet1(_rhs218, (_lhs142).Int())
				_1453_v45 = _rhs219
				var _1458_v48 _dafny.Set
				_ = _1458_v48
				_1458_v48 = _dafny.SetOf(_1453_v45)
				_1453_v45 = ((_1458_v48).Difference(_1458_v48)).Cardinality()
				_1453_v45 = (_1453_v45).Times(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if !(p1) {
						return _1408_v2
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(271))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg61 _dafny.Int) interface{} {
							return coer61(arg61)
						}
					}((func(_1459_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1460_i1 _dafny.Int) _dafny.CodePoint {
							return _1459_v6
						}
					})(_1412_v6)))
				})()).Cardinality()))
				(globalState).F17 = !((_1449_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_1449_v41), 0))).Int()).(bool))
				var _1461_v49 *C2
				_ = _1461_v49
				var _nw253 *C2 = New_C2_()
				_ = _nw253
				_nw253.Ctor__(_1450_v42, (_1449_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_1449_v41), 0))).Int()).(bool))
				_1461_v49 = _nw253
			}
			if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1408_v2, Companion_Default___.Fm16((_this).F19(), p0, globalState)), _1408_v2) {
				var _1462_v50 _dafny.Array
				_ = _1462_v50
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_43
				var _nw254 _dafny.Array
				_ = _nw254
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw254 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) _dafny.Int = func(_1463_i2 _dafny.Int) _dafny.Int {
						return (_1463_i2).Minus(_dafny.IntOfInt64(-313))
					}
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw254 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw254).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw254).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_1462_v50 = _nw254
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1462_v50), 0))
				_ = _index202
				(_1462_v50).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(626))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_1464_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1465_i3 _dafny.Int) _dafny.CodePoint {
						return _1464_v6
					}
				})(_1412_v6)))).Cardinality()), (_index202).Int())
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_1462_v50), 0))
				_ = _index203
				(_1462_v50).ArraySet1(_dafny.IntOfUint32((_1408_v2).Cardinality()), (_index203).Int())
				var _1466_v51 _dafny.Int
				_ = _1466_v51
				_1466_v51 = _dafny.IntOfInt64(-819)
				var _1467_v52 _dafny.Set
				_ = _1467_v52
				_1467_v52 = _dafny.SetOf(p0, p0)
				var _1468_v53 _dafny.Map
				_ = _1468_v53
				_1468_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1412_v6, (_this).F19())
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1462_v50), 0))
				_ = _index204
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_1462_v50), 0))
				_ = _index205
				var _rhs220 _dafny.Int = (_1406_v0).Cardinality()
				_ = _rhs220
				var _rhs221 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1408_v2, Companion_Default___.Fm16(Companion_Default___.Fm0((_this).F19(), _1412_v6, (_this).F19(), globalState), p0, globalState))
				_ = _rhs221
				var _rhs222 _dafny.Int = p0
				_ = _rhs222
				var _rhs223 _dafny.Int = Companion_Default___.SafeDivisionInt(((_1467_v52).Cardinality()).Plus(_1466_v51), ((_1468_v53).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1412_v6, !(!(!(p1)))))).Cardinality())
				_ = _rhs223
				var _lhs143 _dafny.Array = _1462_v50
				_ = _lhs143
				var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1462_v50), 0))
				_ = _lhs144
				var _lhs145 _dafny.Array = _1462_v50
				_ = _lhs145
				var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_1462_v50), 0))
				_ = _lhs146
				(_lhs143).ArraySet1(_rhs220, (_lhs144).Int())
				_1408_v2 = _rhs221
				(_lhs145).ArraySet1(_rhs222, (_lhs146).Int())
				_1466_v51 = _rhs223
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1462_v50), 0))
				_ = _index206
				(_1462_v50).ArraySet1((func() _dafny.Int {
					if (_1466_v51).Cmp((_1462_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1462_v50), 0))).Int()).(_dafny.Int)) <= 0 {
						return _1466_v51
					}
					return Companion_Default___.Fm5(p0, globalState)
				})(), (_index206).Int())
				var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
				_ = _nw255
				_1462_v50 = _nw255
				var _1469_v54 D1
				_ = _1469_v54
				_1469_v54 = Companion_D1_.Create_DC5_(_dafny.UnicodeSeqOfUtf8Bytes("iteqm"), p0, (_1462_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1462_v50), 0))).Int()).(_dafny.Int))
				var _1470_v55 _dafny.Set
				_ = _1470_v55
				_1470_v55 = _dafny.SetOf(_1408_v2, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fofwaytk"), _1408_v2), (_1469_v54).Dtor_cf7())
				_1470_v55 = _dafny.SetOf(_1408_v2, _1408_v2, _dafny.UnicodeSeqOfUtf8Bytes("frcpgrc"))
				var _1471_v56 D2
				_ = _1471_v56
				_1471_v56 = Companion_D2_.Create_DC7_(_dafny.MultiSetFromSeq(_1420_v17))
				var _1472_v57 _dafny.Map
				_ = _1472_v57
				_1472_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _1412_v6)
				var _1473_v58 D4
				_ = _1473_v58
				_1473_v58 = Companion_D4_.Create_DC14_(p1, _1466_v51, _1467_v52, _1412_v6)
				_1471_v56 = Companion_Default___.Fm32(p1, (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1472_v57).Cardinality(), _dafny.IntOfInt64(905))), (_1462_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1462_v50), 0))).Int()).(_dafny.Int), Companion_Default___.Fm0(!((_this).F19()), (_1473_v58).Dtor_cf26(), p1, globalState), globalState)
			} else {
				var _1474_v59 _dafny.Int
				_ = _1474_v59
				_1474_v59 = _dafny.IntOfInt64(-135)
				_1474_v59 = p0
				var _1475_v60 _dafny.Set
				_ = _1475_v60
				_1475_v60 = _dafny.SetOf(_1474_v59)
				var _1476_v61 D4
				_ = _1476_v61
				_1476_v61 = Companion_D4_.Create_DC14_(true, _1474_v59, _1475_v60, _1412_v6)
				var _1477_v62 _dafny.Set
				_ = _1477_v62
				_1477_v62 = _dafny.SetOf(Companion_Default___.Fm0(!(p1), (_1476_v61).Dtor_cf26(), p1, globalState), (_this).F19(), p1)
				var _1478_v63 _dafny.Array
				_ = _1478_v63
				var _nwElement0_56 _dafny.Int = p0
				_ = _nwElement0_56
				var _nw256 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(2))
				_ = _nw256
				(_nw256).ArraySet1(_nwElement0_56, 0)
				(_nw256).ArraySet1((_1477_v62).Cardinality(), 1)
				_1478_v63 = _nw256
				var _1479_v64 _dafny.Sequence
				_ = _1479_v64
				_1479_v64 = _dafny.SeqOf((func() _dafny.Array {
					if (_this).F19() {
						return _1478_v63
					}
					return _1478_v63
				})(), _1478_v63, _1478_v63)
				var _1480_v65 D3
				_ = _1480_v65
				_1480_v65 = Companion_D3_.Create_DC10_((_1418_v15).Cardinality())
				var _1481_v66 _dafny.MultiSet
				_ = _1481_v66
				_1481_v66 = _dafny.MultiSetOf(_1480_v65, _1480_v65)
				var _rhs224 _dafny.Sequence = _dafny.SeqOf(_1478_v63, _1478_v63, _1478_v63, _1478_v63, _1478_v63)
				_ = _rhs224
				var _rhs225 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_this).F19() {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("je"), _1408_v2)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(816))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg63 _dafny.Int) interface{} {
							return coer63(arg63)
						}
					}(func(_1482_i4 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('m')
					}))
				})()).Cardinality())
				_ = _rhs225
				var _rhs226 bool = (func() bool {
					if p1 {
						return (_1407_v1).Select((Companion_Default___.SafeIndex((_1420_v17).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
							if (_1481_v66).Contains(_1480_v65) {
								return (_1481_v66).Multiplicity(_1480_v65)
							}
							return _1474_v59
						})(), _dafny.IntOfUint32((_1420_v17).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1407_v1).Cardinality()))).Uint32()).(bool)
					}
					return p1
				})()
				_ = _rhs226
				var _lhs147 *GlobalState = globalState
				_ = _lhs147
				_1479_v64 = _rhs224
				_1474_v59 = _rhs225
				_lhs147.F17 = _rhs226
				_1474_v59 = (func() _dafny.Int {
					if !_dafny.Companion_Sequence_.Equal(_1408_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(26))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg64 _dafny.Int) interface{} {
							return coer64(arg64)
						}
					}((func(_1483_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1484_i5 _dafny.Int) _dafny.CodePoint {
							return _1483_v6
						}
					})(_1412_v6)))) {
						return ((_1406_v0).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), true))).Cardinality()
					}
					return _dafny.IntOfInt64(-108)
				})()
				var _1485_v67 _dafny.Map
				_ = _1485_v67
				_1485_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(!(p1), (_this).F19())).IsSubsetOf(Companion_Default___.Fm28(_dafny.IntOfUint32((_1408_v2).Cardinality()), true, globalState)), Companion_Default___.Fm33((_this).F19(), true, globalState))
				var _1486_v68 D10
				_ = _1486_v68
				_1486_v68 = Companion_D10_.Create_DC28_(p1, _1474_v59)
				_1485_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1486_v68)
				_1412_v6 = _1412_v6
			}
			(globalState).F15 = (_this).F19()
			var _1487_v69 D0
			_ = _1487_v69
			_1487_v69 = Companion_D0_.Create_DC1_((_this).F19())
			var _1488_v70 T1
			_ = _1488_v70
			var _nw257 *C2 = New_C2_()
			_ = _nw257
			_nw257.Ctor__(_1487_v69, p1)
			_1488_v70 = _nw257
			var _1489_v71 D12
			_ = _1489_v71
			_1489_v71 = Companion_D12_.Create_DC31_((_dafny.SetOf(_1488_v70, _1488_v70)).Difference(_dafny.SetOf(_1488_v70)))
			_1489_v71 = _1489_v71
		} else {
			var _1490_v72 _dafny.CodePoint
			_ = _1490_v72
			_1490_v72 = _dafny.CodePoint('r')
			_1490_v72 = (func() _dafny.CodePoint {
				if (_this).F19() {
					return _1490_v72
				}
				return _1490_v72
			})()
			var _1491_v73 _dafny.Map
			_ = _1491_v73
			_1491_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
			var _1492_v74 _dafny.Array
			_ = _1492_v74
			var _nw258 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw258
			_1492_v74 = _nw258
			var _1493_v75 D1
			_ = _1493_v75
			_1493_v75 = Companion_D1_.Create_DC4_(p0, (_this).F19(), _1492_v74)
			var _1494_v76 _dafny.Sequence
			_ = _1494_v76
			_1494_v76 = _dafny.UnicodeSeqOfUtf8Bytes("ird")
			var _1495_v77 _dafny.Map
			_ = _1495_v77
			_1495_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1493_v75).Dtor_cf5(), _1494_v76)
			_1491_v73 = (_1491_v73).Update(false, (_1495_v77).Cardinality())
			_1406_v0 = (_1406_v0).Update(p1, (!(Companion_Default___.Fm0((_this).F19(), _1490_v72, p1, globalState))) || (false))
			r0 = ((_this).F19()) || ((_this).F19())
			var _1496_v78 D0
			_ = _1496_v78
			_1496_v78 = Companion_D0_.Create_DC1_(true)
			var _1497_v79 *C2
			_ = _1497_v79
			var _nw259 *C2 = New_C2_()
			_ = _nw259
			_nw259.Ctor__(_1496_v78, (_this).F19())
			_1497_v79 = _nw259
			var _1498_v80 _dafny.Map
			_ = _1498_v80
			_1498_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-116), _1497_v79)
			var _1499_v81 _dafny.Set
			_ = _1499_v81
			_1499_v81 = _dafny.SetOf(_1498_v80)
			var _1500_v82 _dafny.MultiSet
			_ = _1500_v82
			_1500_v82 = _dafny.MultiSetOf(p0)
			var _1501_v83 D8
			_ = _1501_v83
			_1501_v83 = Companion_D8_.Create_DC23_(_1500_v82)
			var _1502_v84 _dafny.Sequence
			_ = _1502_v84
			_1502_v84 = _dafny.SeqOf((_1497_v79).F21(), p1)
			var _1503_v85 _dafny.Map
			_ = _1503_v85
			_1503_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1497_v79).F21())
			var _1504_v86 *C3
			_ = _1504_v86
			var _nw260 *C3 = New_C3_()
			_ = _nw260
			_nw260.Ctor__()
			_1504_v86 = _nw260
			var _1505_v87 _dafny.Sequence
			_ = _1505_v87
			_1505_v87 = _dafny.SeqOf(_1501_v83, Companion_D8_.Create_DC23_((Companion_Default___.Fm23((_1502_v84).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_1502_v84).Cardinality()))).Uint32()).(bool), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1497_v79).F21(), _dafny.IntOfUint32((_1494_v76).Cardinality()))).Update((_1497_v79).F21(), (_1503_v85).Cardinality())).Update(!((_1497_v79).F21()), p0), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1504_v86)).Cardinality()), _1490_v72, globalState)).Update(_dafny.IntOfInt64(-748), Companion_Default___.Abs(p0))), Companion_D8_.Create_DC23_(_1500_v82), _1501_v83)
			var _1506_v88 _dafny.Map
			_ = _1506_v88
			_1506_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1499_v81).IsSubsetOf(_1499_v81), _1505_v87)
			_1506_v88 = (_1506_v88).Merge(_1506_v88)
		}
		var _1507_v89 _dafny.CodePoint
		_ = _1507_v89
		_1507_v89 = _dafny.CodePoint('y')
		var _1508_v90 _dafny.Int
		_ = _1508_v90
		_1508_v90 = _dafny.IntOfInt64(975)
		var _1509_v91 _dafny.Array
		_ = _1509_v91
		var _nw261 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
		_ = _nw261
		_1509_v91 = _nw261
		var _1510_v92 D12
		_ = _1510_v92
		_1510_v92 = Companion_D12_.Create_DC32_(_1508_v90)
		var _1511_v93 _dafny.Sequence
		_ = _1511_v93
		_1511_v93 = _dafny.SeqOf((_this).F19(), Companion_Default___.Fm0((_this).F19(), _1507_v89, (_this).F19(), globalState), (_this).F19(), false, true)
		var _1512_v94 _dafny.Map
		_ = _1512_v94
		_1512_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1508_v90)
		var _1513_v95 _dafny.Sequence
		_ = _1513_v95
		_1513_v95 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()))
		var _1514_v96 _dafny.Sequence
		_ = _1514_v96
		_1514_v96 = _dafny.SeqOf((p0).Times(_1508_v90), p0, (_1510_v92).Dtor_cf56(), (func() _dafny.Int {
			if (_1511_v93).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1512_v94).Contains((_this).F19()) {
					return (_1512_v94).Get((_this).F19()).(_dafny.Int)
				}
				return _dafny.IntOfInt64(589)
			})(), _dafny.IntOfUint32((_1511_v93).Cardinality()))).Uint32()).(bool) {
				return p0
			}
			return _1508_v90
		})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(389))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg65 _dafny.Int) interface{} {
				return coer65(arg65)
			}
		}((func(_1515_v90 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1516_i6 _dafny.Int) _dafny.Int {
				return _1515_v90
			}
		})(_1508_v90))), _1513_v95)).Cardinality()))
		var _1517_v97 _dafny.Array
		_ = _1517_v97
		var _nwElement0_57 _dafny.Sequence = _1511_v93
		_ = _nwElement0_57
		var _nw262 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(2))
		_ = _nw262
		(_nw262).ArraySet1(_nwElement0_57, 0)
		(_nw262).ArraySet1(_1511_v93, 1)
		_1517_v97 = _nw262
		var _1518_v98 D1
		_ = _1518_v98
		_1518_v98 = Companion_D1_.Create_DC3_(_1517_v97)
		var _1519_v99 _dafny.MultiSet
		_ = _1519_v99
		_1519_v99 = _dafny.MultiSetOf(_1518_v98)
		var _1520_v100 _dafny.Map
		_ = _1520_v100
		_1520_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1519_v99)
		var _1521_v101 _dafny.Map
		_ = _1521_v101
		_1521_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1520_v100, Companion_Default___.SafeDivisionInt(p0, _1508_v90))
		var _1522_v102 _dafny.Map
		_ = _1522_v102
		_1522_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus(p0)).Cmp(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(p1, p1)).Cardinality(), _1508_v90)).Cardinality())) != 0, _1509_v91)
		var _1523_v103 _dafny.Map
		_ = _1523_v103
		_1523_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1512_v94).Cardinality(), _1513_v95)
		var _rhs227 _dafny.CodePoint = _1507_v89
		_ = _rhs227
		var _rhs228 _dafny.Int = (func() _dafny.Int {
			if (_1521_v101).Contains(_1520_v100) {
				return (_1521_v101).Get(_1520_v100).(_dafny.Int)
			}
			return p0
		})()
		_ = _rhs228
		var _rhs229 _dafny.Array = (func() _dafny.Array {
			if (_1522_v102).Contains((_this).F19()) {
				return (_1522_v102).Get((_this).F19()).(_dafny.Array)
			}
			return _1509_v91
		})()
		_ = _rhs229
		var _rhs230 bool = p1
		_ = _rhs230
		var _rhs231 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_1523_v103).Contains(_1508_v90) {
				return (_1523_v103).Get(_1508_v90).(_dafny.Sequence)
			}
			return _1514_v96
		})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(906))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg66 _dafny.Int) interface{} {
				return coer66(arg66)
			}
		}((func(_1524_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1525_i7 _dafny.Int) _dafny.Int {
				return _1524_p0
			}
		})(p0))))
		_ = _rhs231
		_1507_v89 = _rhs227
		_1508_v90 = _rhs228
		_1509_v91 = _rhs229
		r0 = _rhs230
		_1514_v96 = _rhs231
		var _1526_v104 _dafny.Sequence
		_ = _1526_v104
		_1526_v104 = _dafny.SeqOf(_1406_v0, (_1406_v0).Merge(_1406_v0))
		var _1527_v105 D13
		_ = _1527_v105
		_1527_v105 = Companion_D13_.Create_DC36_(_dafny.SeqOf(_1406_v0))
		_1526_v104 = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (p0).Cmp(p0) > 0 {
				return (_1527_v105).Dtor_cf63()
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(31))).Uint32(), func(coer67 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg67 _dafny.Int) interface{} {
					return coer67(arg67)
				}
			}((func(_1528_v0 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_1529_i8 _dafny.Int) _dafny.Map {
					return _1528_v0
				}
			})(_1406_v0)))
		})(), (Companion_Default___.SafeIndex((_1508_v90).Plus((_dafny.Zero).Minus(_1508_v90)), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (p0).Cmp(p0) > 0 {
				return (_1527_v105).Dtor_cf63()
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(31))).Uint32(), func(coer68 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg68 _dafny.Int) interface{} {
					return coer68(arg68)
				}
			}((func(_1530_v0 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_1531_i8 _dafny.Int) _dafny.Map {
					return _1530_v0
				}
			})(_1406_v0)))
		})()).Cardinality()))).Uint32(), _1406_v0)
		var _1532_v106 D0
		_ = _1532_v106
		_1532_v106 = Companion_D0_.Create_DC1_((_this).F19())
		var _1533_v107 *C2
		_ = _1533_v107
		var _nw263 *C2 = New_C2_()
		_ = _nw263
		_nw263.Ctor__((func() D0 {
			if !((_this).F19()) {
				return _1532_v106
			}
			return _1532_v106
		})(), (func() bool {
			if !(p1) {
				return true
			}
			return (_this).F19()
		})())
		_1533_v107 = _nw263
		var _1534_v108 _dafny.Array
		_ = _1534_v108
		var _nw264 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
		_ = _nw264
		_1534_v108 = _nw264
		var _1535_v109 D13
		_ = _1535_v109
		_1535_v109 = Companion_D13_.Create_DC38_(_1507_v89, (_1533_v107).F21(), _1508_v90, _1534_v108)
		var _1536_v110 _dafny.Array
		_ = _1536_v110
		var _nwElement0_58 D13 = _1535_v109
		_ = _nwElement0_58
		var _nw265 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(5))
		_ = _nw265
		(_nw265).ArraySet1(_nwElement0_58, 0)
		(_nw265).ArraySet1(_1535_v109, 1)
		(_nw265).ArraySet1(_1535_v109, 2)
		(_nw265).ArraySet1(_1535_v109, 3)
		(_nw265).ArraySet1((func() D13 {
			if (_1533_v107).F21() {
				return Companion_D13_.Create_DC38_(_1507_v89, p1, _1508_v90, _1534_v108)
			}
			return _1535_v109
		})(), 4)
		_1536_v110 = _nw265
		var _1537_v111 D14
		_ = _1537_v111
		_1537_v111 = Companion_D14_.Create_DC40_(_1536_v110)
		_1536_v110 = (_1537_v111).Dtor_cf71()
		var _1538_i9 _dafny.Int
		_ = _1538_i9
		_1538_i9 = _dafny.Zero
		{
			for !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1508_v90, _1508_v90)).Contains(Companion_Default___.SafeModuloInt(p0, _1508_v90)) {
				{
					if (_1538_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1538_i9 = (_1538_i9).Plus(_dafny.One)
					var _1539_v112 _dafny.Map
					_ = _1539_v112
					_1539_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1507_v89, _1507_v89)
					var _1540_v113 _dafny.Sequence
					_ = _1540_v113
					_1540_v113 = _dafny.SeqOf(_1507_v89, _dafny.CodePoint('k'), _1507_v89, _1507_v89, _dafny.CodePoint('e'))
					_1539_v112 = (_1539_v112).Update((_1540_v113).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.IntOfUint32((_1540_v113).Cardinality()))).Uint32()).(_dafny.CodePoint), _1507_v89)
					var _rhs232 bool = !(false) || ((_1533_v107).F21())
					_ = _rhs232
					var _rhs233 *C2 = _1533_v107
					_ = _rhs233
					var _lhs148 *GlobalState = globalState
					_ = _lhs148
					_lhs148.F17 = _rhs232
					_1533_v107 = _rhs233
					var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_1534_v108), 0))
					_ = _index207
					(_1534_v108).ArraySet1(p0, (_index207).Int())
					var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_1509_v91), 0))
					_ = _index208
					(_1509_v91).ArraySet1(p1, (_index208).Int())
					var _1541_v114 _dafny.Set
					_ = _1541_v114
					_1541_v114 = _dafny.SetOf(_1507_v89, _1507_v89, _1507_v89)
					var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_1534_v108), 0))
					_ = _index209
					var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_1509_v91), 0))
					_ = _index210
					var _rhs234 _dafny.Int = (func() _dafny.Int {
						if !((_1541_v114).IsSubsetOf(_1541_v114)) {
							return p0
						}
						return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lxxgp"), _1540_v113), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lxxgp"), _1540_v113)).Cardinality()))).Uint32(), _1507_v89)).Cardinality())
					})()
					_ = _rhs234
					var _rhs235 bool = (_this).F19()
					_ = _rhs235
					var _lhs149 _dafny.Array = _1534_v108
					_ = _lhs149
					var _lhs150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_1534_v108), 0))
					_ = _lhs150
					var _lhs151 _dafny.Array = _1509_v91
					_ = _lhs151
					var _lhs152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_1509_v91), 0))
					_ = _lhs152
					(_lhs149).ArraySet1(_rhs234, (_lhs150).Int())
					(_lhs151).ArraySet1(_rhs235, (_lhs152).Int())
					_1508_v90 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1540_v113, _1540_v113)).Cardinality())
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		r0 = (_1533_v107.F20).Dtor_cf1()
		return r0
	}
}
func (_this *C4) F19() bool {
	{
		return _this._f19
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f18 _dafny.Int
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f18 = _dafny.Zero
	return &_this
}

type CompanionStruct_C5_ struct {
}

var Companion_C5_ = CompanionStruct_C5_{}

func (_this *C5) Equals(other *C5) bool {
	return _this == other
}

func (_this *C5) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C5)
	return ok && _this.Equals(other)
}

func (*C5) String() string {
	return "_module.C5"
}

func Type_C5_() _dafny.TypeDescriptor {
	return type_C5_{}
}

type type_C5_ struct {
}

func (_this type_C5_) Default() interface{} {
	return (*C5)(nil)
}

func (_this type_C5_) String() string {
	return "main.C5"
}
func (_this *C5) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) Ctor__(f18 _dafny.Int) {
	{
		(_this)._f18 = f18
	}
}
func (_this *C5) M0(p0 _dafny.Set, p1 _dafny.Set, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1542_v0 bool
		_ = _1542_v0
		_1542_v0 = true
		var _1543_v1 _dafny.Sequence
		_ = _1543_v1
		_1543_v1 = _dafny.SeqOf(_1542_v0)
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_1543_v1, _1543_v1) {
			var _1544_v2 _dafny.Array
			_ = _1544_v2
			var _len0_44 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_44
			var _nw266 _dafny.Array
			_ = _nw266
			if _len0_44.Cmp(_dafny.Zero) == 0 {
				_nw266 = _dafny.NewArray(_len0_44)
			} else {
				var _init44 func(_dafny.Int) _dafny.Sequence = (func(_1545_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1546_i0 _dafny.Int) _dafny.Sequence {
						return _1545_v1
					}
				})(_1543_v1)
				_ = _init44
				var _element0_44 = _init44(_dafny.Zero)
				_ = _element0_44
				_nw266 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
				(_nw266).ArraySet1(_element0_44, 0)
				var _nativeLen0_44 = (_len0_44).Int()
				_ = _nativeLen0_44
				for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
					(_nw266).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
				}
			}
			_1544_v2 = _nw266
			var _1547_v3 D1
			_ = _1547_v3
			_1547_v3 = Companion_D1_.Create_DC3_(_1544_v2)
			_1544_v2 = (_1547_v3).Dtor_cf3()
			var _1548_v4 _dafny.Array
			_ = _1548_v4
			var _len0_45 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_45
			var _nw267 _dafny.Array
			_ = _nw267
			if _len0_45.Cmp(_dafny.Zero) == 0 {
				_nw267 = _dafny.NewArray(_len0_45)
			} else {
				var _init45 func(_dafny.Int) D1 = func(_1549_i1 _dafny.Int) D1 {
					return Companion_D1_.Create_DC5_(_dafny.UnicodeSeqOfUtf8Bytes("odrkr"), (_this).F18(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tbdrju")).Cardinality()))
				}
				_ = _init45
				var _element0_45 = _init45(_dafny.Zero)
				_ = _element0_45
				_nw267 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
				(_nw267).ArraySet1(_element0_45, 0)
				var _nativeLen0_45 = (_len0_45).Int()
				_ = _nativeLen0_45
				for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
					(_nw267).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
				}
			}
			_1548_v4 = _nw267
			var _1550_v5 _dafny.Sequence
			_ = _1550_v5
			_1550_v5 = _dafny.UnicodeSeqOfUtf8Bytes("bnr")
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1548_v4), 0))
			_ = _index211
			(_1548_v4).ArraySet1(Companion_D1_.Create_DC5_(_1550_v5, (_this).F18(), (_this).F18()), (_index211).Int())
			var _1551_v6 _dafny.Sequence
			_ = _1551_v6
			_1551_v6 = _dafny.SeqOf(_dafny.IntOfInt64(797))
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1548_v4), 0))
			_ = _index212
			var _rhs236 D1 = Companion_Default___.Fm1((_this).F18(), globalState)
			_ = _rhs236
			var _rhs237 bool = (_dafny.MultiSetOf((_this).F18())).IsProperSubsetOf(_dafny.MultiSetFromSeq(_1551_v6))
			_ = _rhs237
			var _rhs238 bool = (p0).IsDisjointFrom(p0)
			_ = _rhs238
			var _rhs239 bool = (_1543_v1).Select((Companion_Default___.SafeIndex(((_this).F18()).Times(_dafny.IntOfUint32((_1551_v6).Cardinality())), _dafny.IntOfUint32((_1543_v1).Cardinality()))).Uint32()).(bool)
			_ = _rhs239
			var _lhs153 _dafny.Array = _1548_v4
			_ = _lhs153
			var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_1548_v4), 0))
			_ = _lhs154
			var _lhs155 *GlobalState = globalState
			_ = _lhs155
			var _lhs156 *GlobalState = globalState
			_ = _lhs156
			(_lhs153).ArraySet1(_rhs236, (_lhs154).Int())
			_1542_v0 = _rhs237
			_lhs155.F15 = _rhs238
			_lhs156.F11 = _rhs239
			var _1552_v7 _dafny.Map
			_ = _1552_v7
			_1552_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_this).F18())
			_1552_v7 = (_1552_v7).Update((_this).F18(), (_dafny.IntOfInt64(166)).Minus((_this).F18()))
			var _1553_v8 _dafny.Array
			_ = _1553_v8
			var _nw268 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw268
			_1553_v8 = _nw268
			var _1554_v9 D1
			_ = _1554_v9
			_1554_v9 = Companion_D1_.Create_DC6_((_this).F18(), Companion_Default___.SafeDivisionInt((_this).F18(), (_this).F18()))
			var _1555_v10 D0
			_ = _1555_v10
			_1555_v10 = Companion_D0_.Create_DC1_(_1542_v0)
			var _1556_v11 D0
			_ = _1556_v11
			_1556_v11 = Companion_D0_.Create_DC2_(_1555_v10)
			var _rhs240 _dafny.Array = _1553_v8
			_ = _rhs240
			var _rhs241 _dafny.Int = (func() _dafny.Int {
				if false {
					return ((_this).F18()).Times(_dafny.IntOfUint32((_1543_v1).Cardinality()))
				}
				return (_this).F18()
			})()
			_ = _rhs241
			var _rhs242 _dafny.Int = (_this).F18()
			_ = _rhs242
			var _rhs243 D1 = Companion_Default___.Fm2(globalState)
			_ = _rhs243
			var _rhs244 D0 = _1556_v11
			_ = _rhs244
			_1553_v8 = _rhs240
			r0 = _rhs241
			r0 = _rhs242
			_1554_v9 = _rhs243
			_1556_v11 = _rhs244
			var _1557_v12 _dafny.Sequence
			_ = _1557_v12
			_1557_v12 = _dafny.SeqOf(_1550_v5, _1550_v5, _1550_v5, _1550_v5, _1550_v5)
			if !(!_dafny.Companion_Sequence_.Equal(_1557_v12, _1557_v12)) {
				var _1558_v13 _dafny.Sequence
				_ = _1558_v13
				_1558_v13 = _dafny.SeqOf(_dafny.SeqOf(_1542_v0))
				var _1559_v14 _dafny.MultiSet
				_ = _1559_v14
				_1559_v14 = _dafny.MultiSetOf(_1542_v0)
				var _1560_v15 _dafny.Map
				_ = _1560_v15
				_1560_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1542_v0, (_1559_v14).Cardinality())
				var _1561_v16 _dafny.MultiSet
				_ = _1561_v16
				_1561_v16 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(228))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg69 _dafny.Int) interface{} {
						return coer69(arg69)
					}
				}((func(_1562_v5 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
					return func(_1563_i2 _dafny.Int) _dafny.CodePoint {
						return (_1562_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.IntOfUint32((_1562_v5).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_1550_v5)))).Cardinality()))
				var _1564_v17 _dafny.Array
				_ = _1564_v17
				var _nwElement0_59 _dafny.Int = (_this).F18()
				_ = _nwElement0_59
				var _nw269 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(23))
				_ = _nw269
				(_nw269).ArraySet1(_nwElement0_59, 0)
				(_nw269).ArraySet1((_this).F18(), 1)
				(_nw269).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((_1558_v13).Select((Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_1558_v13).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), (_1559_v14).Cardinality()), 2)
				(_nw269).ArraySet1((p0).Cardinality(), 3)
				(_nw269).ArraySet1(((_this).F18()).Plus((_this).F18()), 4)
				(_nw269).ArraySet1((_this).F18(), 5)
				(_nw269).ArraySet1((_this).F18(), 6)
				(_nw269).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1551_v6).Cardinality()), _dafny.IntOfInt64(-560)), 7)
				(_nw269).ArraySet1(_dafny.IntOfInt64(871), 8)
				(_nw269).ArraySet1((_dafny.Zero).Minus((_this).F18()), 9)
				(_nw269).ArraySet1(_dafny.IntOfInt64(231), 10)
				(_nw269).ArraySet1(((_dafny.Zero).Minus((_this).F18())).Times((_this).F18()), 11)
				(_nw269).ArraySet1((func() _dafny.Int {
					if (_1560_v15).Contains(_1542_v0) {
						return (_1560_v15).Get(_1542_v0).(_dafny.Int)
					}
					return (_this).F18()
				})(), 12)
				(_nw269).ArraySet1((_dafny.Zero).Minus((_this).F18()), 13)
				(_nw269).ArraySet1((_this).F18(), 14)
				(_nw269).ArraySet1((((_1561_v16).Update((_1552_v7).Cardinality(), Companion_Default___.Abs(_dafny.IntOfInt64(655)))).Intersection(_1561_v16)).Cardinality(), 15)
				(_nw269).ArraySet1((_this).F18(), 16)
				(_nw269).ArraySet1((_this).F18(), 17)
				(_nw269).ArraySet1(Companion_Default___.Fm3((_1551_v6).Select((Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_1551_v6).Cardinality()))).Uint32()).(_dafny.Int), (_this).F18(), _1542_v0, globalState), 18)
				(_nw269).ArraySet1((Companion_Default___.Fm3((_this).F18(), (_this).F18(), _1542_v0, globalState)).Times((_this).F18()), 19)
				(_nw269).ArraySet1((_this).F18(), 20)
				(_nw269).ArraySet1((_this).F18(), 21)
				(_nw269).ArraySet1(_dafny.IntOfUint32((_1550_v5).Cardinality()), 22)
				_1564_v17 = _nw269
				var _1565_v18 _dafny.CodePoint
				_ = _1565_v18
				_1565_v18 = _dafny.CodePoint('j')
				var _1566_v19 _dafny.MultiSet
				_ = _1566_v19
				_1566_v19 = _dafny.MultiSetOf(_1550_v5, _dafny.SeqOf(_1565_v18))
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_1564_v17), 0))
				_ = _index213
				(_1564_v17).ArraySet1(((func() _dafny.Int {
					if (_1566_v19).Contains(_1550_v5) {
						return (_1566_v19).Multiplicity(_1550_v5)
					}
					return (_this).F18()
				})()).Plus((_this).F18()), (_index213).Int())
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_1553_v8), 0))
				_ = _index214
				(_1553_v8).ArraySet1(_1542_v0, (_index214).Int())
				var _1567_v20 D0
				_ = _1567_v20
				_1567_v20 = Companion_D0_.Create_DC1_(_1542_v0)
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_1564_v17), 0))
				_ = _index215
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_1553_v8), 0))
				_ = _index216
				var _rhs245 _dafny.Int = (_dafny.Zero).Minus((_this).F18())
				_ = _rhs245
				var _rhs246 bool = _1542_v0
				_ = _rhs246
				var _rhs247 D0 = _1567_v20
				_ = _rhs247
				var _lhs157 _dafny.Array = _1564_v17
				_ = _lhs157
				var _lhs158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_1564_v17), 0))
				_ = _lhs158
				var _lhs159 _dafny.Array = _1553_v8
				_ = _lhs159
				var _lhs160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_1553_v8), 0))
				_ = _lhs160
				(_lhs157).ArraySet1(_rhs245, (_lhs158).Int())
				(_lhs159).ArraySet1(_rhs246, (_lhs160).Int())
				_1567_v20 = _rhs247
				(globalState).F11 = _dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
					if _1542_v0 {
						return _1543_v1
					}
					return _dafny.SeqOf(true, _1542_v0)
				})(), _1543_v1)
				var _1568_v21 *C0
				_ = _1568_v21
				var _nw270 *C0 = New_C0_()
				_ = _nw270
				_nw270.Ctor__(_1542_v0)
				_1568_v21 = _nw270
				_1550_v5 = _dafny.Companion_Sequence_.Concatenate(_1550_v5, _1550_v5)
				var _1569_v22 T1
				_ = _1569_v22
				var _nw271 *C2 = New_C2_()
				_ = _nw271
				_nw271.Ctor__(Companion_D0_.Create_DC1_(_1542_v0), (func() bool {
					if (_1568_v21).F23() {
						return (_1553_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_1553_v8), 0))).Int()).(bool)
					}
					return (_1553_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_1553_v8), 0))).Int()).(bool)
				})())
				_1569_v22 = _nw271
				var _nw272 *C4 = New_C4_()
				_ = _nw272
				_nw272.Ctor__((_1553_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_1553_v8), 0))).Int()).(bool))
				_1569_v22 = _nw272
			} else {
				(globalState).F17 = (p0).IsSubsetOf(p0)
				var _1570_v23 _dafny.Map
				_ = _1570_v23
				_1570_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_1542_v0), _1542_v0)
				r0 = Companion_Default___.Fm9((_this).F18(), Companion_Default___.Fm5(_dafny.IntOfUint32((_1551_v6).Cardinality()), globalState), (func() bool {
					if (_1570_v23).Contains(Companion_Default___.Fm20(globalState)) {
						return (_1570_v23).Get(Companion_Default___.Fm20(globalState)).(bool)
					}
					return true
				})(), globalState)
				var _1571_v24 _dafny.Map
				_ = _1571_v24
				_1571_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _1542_v0)
				var _1572_v25 _dafny.CodePoint
				_ = _1572_v25
				_1572_v25 = _dafny.CodePoint('n')
				(globalState).F15 = (func() bool {
					if Companion_Default___.Fm0((func() bool {
						if (_1571_v24).Contains((_this).F18()) {
							return (_1571_v24).Get((_this).F18()).(bool)
						}
						return _1542_v0
					})(), _1572_v25, _1542_v0, globalState) {
						return _1542_v0
					}
					return Companion_Default___.Fm0(_1542_v0, _dafny.CodePoint('q'), _1542_v0, globalState)
				})()
				_1542_v0 = ((_this).F18()).Cmp((_dafny.Zero).Minus((_this).F18())) < 0
				(globalState).F16 = _1543_v1
			}
		} else {
			if _1542_v0 {
				var _1573_v26 _dafny.CodePoint
				_ = _1573_v26
				_1573_v26 = _dafny.CodePoint('c')
				var _1574_v27 _dafny.Array
				_ = _1574_v27
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_46
				var _nw273 _dafny.Array
				_ = _nw273
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw273 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) bool = (func(_1575_v0 bool) func(_dafny.Int) bool {
						return func(_1576_i3 _dafny.Int) bool {
							return _1575_v0
						}
					})(_1542_v0)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw273 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw273).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw273).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1574_v27 = _nw273
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1574_v27), 0))
				_ = _index217
				(_1574_v27).ArraySet1(_1542_v0, (_index217).Int())
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1574_v27), 0))
				_ = _index218
				var _rhs248 _dafny.CodePoint = _dafny.CodePoint('o')
				_ = _rhs248
				var _rhs249 bool = !(_1542_v0)
				_ = _rhs249
				var _lhs161 _dafny.Array = _1574_v27
				_ = _lhs161
				var _lhs162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1574_v27), 0))
				_ = _lhs162
				_1573_v26 = _rhs248
				(_lhs161).ArraySet1(_rhs249, (_lhs162).Int())
				r0 = (_this).F18()
				r0 = (_this).F18()
				_1573_v26 = _1573_v26
				var _1577_v28 _dafny.Map
				_ = _1577_v28
				_1577_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_this).F18(), (_this).F18()), (_dafny.Zero).Minus((_this).F18()))
				var _1578_v29 _dafny.Array
				_ = _1578_v29
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_47
				var _nw274 _dafny.Array
				_ = _nw274
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw274 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) _dafny.CodePoint = (func(_1579_v0 bool, _1580_p0 _dafny.Set, _1581_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1582_i4 _dafny.Int) _dafny.CodePoint {
							return (Companion_D4_.Create_DC14_(_1579_v0, (_this).F18(), _1580_p0, _1581_v26)).Dtor_cf26()
						}
					})(_1542_v0, p0, _1573_v26)
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw274 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw274).ArraySet1CodePoint(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw274).ArraySet1CodePoint(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_1578_v29 = _nw274
				var _1583_v30 _dafny.Sequence
				_ = _1583_v30
				_1583_v30 = _dafny.SeqOf(_1578_v29, _1578_v29)
				_1577_v28 = (_1577_v28).Update(Companion_Default___.SafeModuloInt((_this).F18(), _dafny.IntOfUint32((_1583_v30).Cardinality())), (_this).F18())
			} else {
				var _1584_v31 _dafny.Sequence
				_ = _1584_v31
				_1584_v31 = _dafny.UnicodeSeqOfUtf8Bytes("xm")
				var _1585_v32 _dafny.Array
				_ = _1585_v32
				var _nwElement0_60 bool = _1542_v0
				_ = _nwElement0_60
				var _nw275 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(5))
				_ = _nw275
				(_nw275).ArraySet1(_nwElement0_60, 0)
				(_nw275).ArraySet1(_1542_v0, 1)
				(_nw275).ArraySet1(_1542_v0, 2)
				(_nw275).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1584_v31, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(542))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg70 _dafny.Int) interface{} {
						return coer70(arg70)
					}
				}(func(_1586_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				}))), 3)
				(_nw275).ArraySet1(!((_1543_v1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F18()), _dafny.IntOfUint32((_1543_v1).Cardinality()))).Uint32()).(bool)) || (_1542_v0), 4)
				_1585_v32 = _nw275
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_1585_v32), 0))
				_ = _index219
				(_1585_v32).ArraySet1(_1542_v0, (_index219).Int())
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_1585_v32), 0))
				_ = _index220
				(_1585_v32).ArraySet1(true, (_index220).Int())
				var _1587_v33 _dafny.MultiSet
				_ = _1587_v33
				_1587_v33 = _dafny.MultiSetOf((_this).F18())
				var _1588_v34 D2
				_ = _1588_v34
				_1588_v34 = Companion_D2_.Create_DC7_(_1587_v33)
				var _1589_v35 _dafny.Sequence
				_ = _1589_v35
				_1589_v35 = _dafny.SeqOf(_1588_v34, _1588_v34, _1588_v34, _1588_v34)
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_1585_v32), 0))
				_ = _index221
				(_1585_v32).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(149))).Uint32(), func(coer71 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}(func(_1590_i6 _dafny.Int) D2 {
					return Companion_D2_.Create_DC7_(_dafny.MultiSetOf((_this).F18()))
				})), _1589_v35), (_index221).Int())
				r0 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1584_v31, _1584_v31)).Cardinality()))
				var _1591_v36 _dafny.Map
				_ = _1591_v36
				_1591_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_this).F18())
				var _1592_v37 _dafny.Array
				_ = _1592_v37
				var _nwElement0_61 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1584_v31, _dafny.UnicodeSeqOfUtf8Bytes("kkptvhk"))).Cardinality())
				_ = _nwElement0_61
				var _nw276 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(14))
				_ = _nw276
				(_nw276).ArraySet1(_nwElement0_61, 0)
				(_nw276).ArraySet1((_this).F18(), 1)
				(_nw276).ArraySet1((_this).F18(), 2)
				(_nw276).ArraySet1((_1591_v36).Cardinality(), 3)
				(_nw276).ArraySet1((_this).F18(), 4)
				(_nw276).ArraySet1((_this).F18(), 5)
				(_nw276).ArraySet1((_dafny.IntOfUint32((_1584_v31).Cardinality())).Times(_dafny.IntOfInt64(151)), 6)
				(_nw276).ArraySet1(((_this).F18()).Times((_this).F18()), 7)
				(_nw276).ArraySet1((_1591_v36).Cardinality(), 8)
				(_nw276).ArraySet1(_dafny.IntOfInt64(453), 9)
				(_nw276).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ddlyqlq")).Cardinality()), 10)
				(_nw276).ArraySet1((_this).F18(), 11)
				(_nw276).ArraySet1(Companion_Default___.SafeModuloInt((_this).F18(), (_this).F18()), 12)
				(_nw276).ArraySet1((_this).F18(), 13)
				_1592_v37 = _nw276
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_1592_v37), 0))
				_ = _index222
				(_1592_v37).ArraySet1((_this).F18(), (_index222).Int())
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_1592_v37), 0))
				_ = _index223
				(_1592_v37).ArraySet1(_dafny.IntOfInt64(476), (_index223).Int())
				var _1593_v38 _dafny.Map
				_ = _1593_v38
				_1593_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1585_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_1585_v32), 0))).Int()).(bool), _dafny.IntOfInt64(-546))
				var _1594_v39 D3
				_ = _1594_v39
				_1594_v39 = Companion_D3_.Create_DC11_(_1593_v38, (_this).F18(), _1592_v37)
				_1594_v39 = Companion_D3_.Create_DC11_(_1593_v38, (_this).F18(), _1592_v37)
			}
			r0 = (_this).F18()
			var _1595_v40 _dafny.MultiSet
			_ = _1595_v40
			_1595_v40 = _dafny.MultiSetOf(_1542_v0)
			r0 = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm5(((Companion_D10_.Create_DC27_(_1595_v40)).Dtor_cf48()).Cardinality(), globalState), (_this).F18())
			var _1596_v41 _dafny.Sequence
			_ = _1596_v41
			_1596_v41 = _dafny.UnicodeSeqOfUtf8Bytes("datkd")
			_1596_v41 = _dafny.Companion_Sequence_.Concatenate(_1596_v41, _1596_v41)
			var _1597_v42 *C4
			_ = _1597_v42
			var _nw277 *C4 = New_C4_()
			_ = _nw277
			_nw277.Ctor__(_1542_v0)
			_1597_v42 = _nw277
			var _1598_v43 _dafny.Map
			_ = _1598_v43
			_1598_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1597_v42, (_this).F18())
			var _1599_v44 D0
			_ = _1599_v44
			_1599_v44 = Companion_D0_.Create_DC1_(_1542_v0)
			var _1600_v45 *C2
			_ = _1600_v45
			var _nw278 *C2 = New_C2_()
			_ = _nw278
			_nw278.Ctor__(_1599_v44, _1542_v0)
			_1600_v45 = _nw278
			var _1601_v46 _dafny.Array
			_ = _1601_v46
			var _nw279 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
			_ = _nw279
			_1601_v46 = _nw279
			var _1602_v47 D14
			_ = _1602_v47
			_1602_v47 = Companion_D14_.Create_DC41_(_1600_v45, _1596_v41, _1601_v46)
			var _1603_v48 D14
			_ = _1603_v48
			_1603_v48 = Companion_D14_.Create_DC41_((_1602_v47).Dtor_cf72(), _1596_v41, _1601_v46)
			var _1604_v49 _dafny.Map
			_ = _1604_v49
			_1604_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1603_v48, (_this).F18())
			var _1605_v50 _dafny.Sequence
			_ = _1605_v50
			_1605_v50 = _dafny.SeqOf((_this).F18())
			var _1606_v51 _dafny.Map
			_ = _1606_v51
			_1606_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1542_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jdetgprt")).Cardinality()))
			var _1607_v52 _dafny.Map
			_ = _1607_v52
			_1607_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1596_v41, (_1606_v51).Update((_1597_v42).F19(), (_this).F18()))
			var _1608_v53 _dafny.CodePoint
			_ = _1608_v53
			_1608_v53 = _dafny.CodePoint('c')
			var _1609_v54 _dafny.Array
			_ = _1609_v54
			var _nwElement0_62 _dafny.Int = Companion_Default___.Fm5((_this).F18(), globalState)
			_ = _nwElement0_62
			var _nw280 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(20))
			_ = _nw280
			(_nw280).ArraySet1(_nwElement0_62, 0)
			(_nw280).ArraySet1((func() _dafny.Int {
				if (_1598_v43).Contains(_1597_v42) {
					return (_1598_v43).Get(_1597_v42).(_dafny.Int)
				}
				return (_this).F18()
			})(), 1)
			(_nw280).ArraySet1((_this).F18(), 2)
			(_nw280).ArraySet1((_this).F18(), 3)
			(_nw280).ArraySet1(((_this).F18()).Plus((_this).F18()), 4)
			(_nw280).ArraySet1(((p0).Intersection(p0)).Cardinality(), 5)
			(_nw280).ArraySet1((func() _dafny.Int {
				if (_1604_v49).Contains(_1602_v47) {
					return (_1604_v49).Get(_1602_v47).(_dafny.Int)
				}
				return (_this).F18()
			})(), 6)
			(_nw280).ArraySet1((_this).F18(), 7)
			(_nw280).ArraySet1((_this).F18(), 8)
			(_nw280).ArraySet1((_this).F18(), 9)
			(_nw280).ArraySet1(Companion_Default___.Fm5(_dafny.IntOfInt64(28), globalState), 10)
			(_nw280).ArraySet1((_this).F18(), 11)
			(_nw280).ArraySet1((_dafny.IntOfInt64(667)).Plus((_this).F18()), 12)
			(_nw280).ArraySet1((_this).F18(), 13)
			(_nw280).ArraySet1((_this).F18(), 14)
			(_nw280).ArraySet1((_this).F18(), 15)
			(_nw280).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm19(_1605_v50, _1542_v0, globalState)).Cardinality()), 16)
			(_nw280).ArraySet1((_this).F18(), 17)
			(_nw280).ArraySet1((_this).F18(), 18)
			(_nw280).ArraySet1((func() _dafny.Int {
				if (_1597_v42).Fm4((_1597_v42).F19(), _1607_v52, _1608_v53, (_this).F18(), globalState) {
					return (_dafny.Zero).Minus((_this).F18())
				}
				return (_this).F18()
			})(), 19)
			_1609_v54 = _nw280
			var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_1609_v54), 0))
			_ = _index224
			(_1609_v54).ArraySet1(_dafny.IntOfInt64(611), (_index224).Int())
			var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_1609_v54), 0))
			_ = _index225
			(_1609_v54).ArraySet1((_this).F18(), (_index225).Int())
		}
		var _1610_v55 _dafny.Sequence
		_ = _1610_v55
		_1610_v55 = _dafny.UnicodeSeqOfUtf8Bytes("ftlcnhg")
		var _1611_v56 _dafny.Map
		_ = _1611_v56
		_1611_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_this).F18())
		var _1612_v57 D0
		_ = _1612_v57
		_1612_v57 = Companion_D0_.Create_DC1_(_1542_v0)
		var _1613_v58 _dafny.CodePoint
		_ = _1613_v58
		_1613_v58 = _dafny.CodePoint('b')
		var _1614_v59 _dafny.Map
		_ = _1614_v59
		_1614_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1612_v57, _1613_v58)
		var _1615_v60 _dafny.Map
		_ = _1615_v60
		_1615_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1543_v1).Select((Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_1543_v1).Cardinality()))).Uint32()).(bool), (_this).F18())
		var _1616_v61 _dafny.MultiSet
		_ = _1616_v61
		_1616_v61 = _dafny.MultiSetOf(_1542_v0)
		var _1617_v62 _dafny.Set
		_ = _1617_v62
		_1617_v62 = _dafny.SetOf(_1616_v61, _1616_v61)
		var _1618_v63 _dafny.Array
		_ = _1618_v63
		var _nwElement0_63 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg72 _dafny.Int) interface{} {
				return coer72(arg72)
			}
		}(func(_1619_i7 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		}))).Cardinality())
		_ = _nwElement0_63
		var _nw281 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(23))
		_ = _nw281
		(_nw281).ArraySet1(_nwElement0_63, 0)
		(_nw281).ArraySet1((_this).F18(), 1)
		(_nw281).ArraySet1(_dafny.IntOfUint32((_1610_v55).Cardinality()), 2)
		(_nw281).ArraySet1((_this).F18(), 3)
		(_nw281).ArraySet1(_dafny.IntOfInt64(306), 4)
		(_nw281).ArraySet1((_this).F18(), 5)
		(_nw281).ArraySet1((_this).F18(), 6)
		(_nw281).ArraySet1((_this).F18(), 7)
		(_nw281).ArraySet1((_dafny.Zero).Minus((_this).F18()), 8)
		(_nw281).ArraySet1((_this).F18(), 9)
		(_nw281).ArraySet1(((_1611_v56).Update((_this).F18(), (_this).F18())).Cardinality(), 10)
		(_nw281).ArraySet1((_this).F18(), 11)
		(_nw281).ArraySet1((_this).F18(), 12)
		(_nw281).ArraySet1((_dafny.Zero).Minus((_1614_v59).Cardinality()), 13)
		(_nw281).ArraySet1((func() _dafny.Int {
			if (_1615_v60).Contains(false) {
				return (_1615_v60).Get(false).(_dafny.Int)
			}
			return (_this).F18()
		})(), 14)
		(_nw281).ArraySet1((_dafny.Zero).Minus((_1617_v62).Cardinality()), 15)
		(_nw281).ArraySet1((_this).F18(), 16)
		(_nw281).ArraySet1((_this).F18(), 17)
		(_nw281).ArraySet1((_this).F18(), 18)
		(_nw281).ArraySet1(_dafny.IntOfUint32((_1610_v55).Cardinality()), 19)
		(_nw281).ArraySet1((_dafny.Zero).Minus((_this).F18()), 20)
		(_nw281).ArraySet1((_this).F18(), 21)
		(_nw281).ArraySet1((_this).F18(), 22)
		_1618_v63 = _nw281
		var _1620_v64 _dafny.Map
		_ = _1620_v64
		_1620_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1618_v63, _1610_v55)
		_1620_v64 = _1620_v64
		var _1621_i8 _dafny.Int
		_ = _1621_i8
		_1621_i8 = _dafny.Zero
		{
			for _1542_v0 {
				{
					if (_1621_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_1621_i8 = (_1621_i8).Plus(_dafny.One)
					var _1622_v65 _dafny.Map
					_ = _1622_v65
					_1622_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _1542_v0)
					var _1623_v66 _dafny.Set
					_ = _1623_v66
					_1623_v66 = _dafny.SetOf((Companion_Default___.Fm17(_1542_v0, _1542_v0, globalState)).Update((_this).F18(), Companion_Default___.Fm0(_1542_v0, _dafny.CodePoint('i'), _1542_v0, globalState)), _1622_v65)
					var _source15 D3 = Companion_D3_.Create_DC9_(_1623_v66)
					_ = _source15
					if _source15.Is_DC10() {
						var _1624___mcc_h0 _dafny.Int = _source15.Get_().(D3_DC10).Cf17
						_ = _1624___mcc_h0
						var _1625_cf17 _dafny.Int = _1624___mcc_h0
						_ = _1625_cf17
						(globalState).F11 = ((_this).F18()).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_1625_cf17)))) >= 0
						var _1626_v67 _dafny.Array
						_ = _1626_v67
						var _nw282 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
						_ = _nw282
						_1626_v67 = _nw282
						_1626_v67 = _1626_v67
						var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_1626_v67), 0))
						_ = _index226
						(_1626_v67).ArraySet1(_1542_v0, (_index226).Int())
						var _1627_v68 _dafny.Array
						_ = _1627_v68
						var _nwElement0_64 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(508))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg73 _dafny.Int) interface{} {
								return coer73(arg73)
							}
						}((func(_1628_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1629_i9 _dafny.Int) _dafny.CodePoint {
								return _1628_v58
							}
						})(_1613_v58)))
						_ = _nwElement0_64
						var _nw283 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(11))
						_ = _nw283
						(_nw283).ArraySet1(_nwElement0_64, 0)
						(_nw283).ArraySet1(_1610_v55, 1)
						(_nw283).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mk"), 2)
						(_nw283).ArraySet1(_1610_v55, 3)
						(_nw283).ArraySet1(_1610_v55, 4)
						(_nw283).ArraySet1(_1610_v55, 5)
						(_nw283).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qpit"), 6)
						(_nw283).ArraySet1(_1610_v55, 7)
						(_nw283).ArraySet1(_1610_v55, 8)
						(_nw283).ArraySet1(_dafny.Companion_Sequence_.Update(_1610_v55, (Companion_Default___.SafeIndex(_1625_cf17, _dafny.IntOfUint32((_1610_v55).Cardinality()))).Uint32(), _1613_v58), 9)
						(_nw283).ArraySet1(_1610_v55, 10)
						_1627_v68 = _nw283
						var _1630_v69 _dafny.Array
						_ = _1630_v69
						var _nwElement0_65 _dafny.Array = _1627_v68
						_ = _nwElement0_65
						var _nw284 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(17))
						_ = _nw284
						(_nw284).ArraySet1(_nwElement0_65, 0)
						(_nw284).ArraySet1(_1627_v68, 1)
						(_nw284).ArraySet1(_1627_v68, 2)
						(_nw284).ArraySet1(_1627_v68, 3)
						(_nw284).ArraySet1((func() _dafny.Array {
							if _1542_v0 {
								return _1627_v68
							}
							return _1627_v68
						})(), 4)
						(_nw284).ArraySet1(_1627_v68, 5)
						(_nw284).ArraySet1(_1627_v68, 6)
						(_nw284).ArraySet1(_1627_v68, 7)
						(_nw284).ArraySet1(_1627_v68, 8)
						(_nw284).ArraySet1(_1627_v68, 9)
						(_nw284).ArraySet1(_1627_v68, 10)
						(_nw284).ArraySet1(_1627_v68, 11)
						(_nw284).ArraySet1(_1627_v68, 12)
						(_nw284).ArraySet1(_1627_v68, 13)
						(_nw284).ArraySet1(_1627_v68, 14)
						(_nw284).ArraySet1(_1627_v68, 15)
						(_nw284).ArraySet1((func() _dafny.Array {
							if _1542_v0 {
								return _1627_v68
							}
							return _1627_v68
						})(), 16)
						_1630_v69 = _nw284
						var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_1630_v69), 0))
						_ = _index227
						(_1630_v69).ArraySet1(_1627_v68, (_index227).Int())
						var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_1626_v67), 0))
						_ = _index228
						var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_1630_v69), 0))
						_ = _index229
						var _rhs250 bool = (_1543_v1).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_this).F18(), (_this).F18()), _dafny.IntOfUint32((_1543_v1).Cardinality()))).Uint32()).(bool)
						_ = _rhs250
						var _rhs251 _dafny.Array = _1627_v68
						_ = _rhs251
						var _lhs163 _dafny.Array = _1626_v67
						_ = _lhs163
						var _lhs164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_1626_v67), 0))
						_ = _lhs164
						var _lhs165 _dafny.Array = _1630_v69
						_ = _lhs165
						var _lhs166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_1630_v69), 0))
						_ = _lhs166
						(_lhs163).ArraySet1(_rhs250, (_lhs164).Int())
						(_lhs165).ArraySet1(_rhs251, (_lhs166).Int())
						_1611_v56 = (_1611_v56).Update(_1625_cf17, _dafny.IntOfUint32((_1610_v55).Cardinality()))
					} else if _source15.Is_DC11() {
						var _1631___mcc_h1 _dafny.Map = _source15.Get_().(D3_DC11).Cf18
						_ = _1631___mcc_h1
						var _1632___mcc_h2 _dafny.Int = _source15.Get_().(D3_DC11).Cf19
						_ = _1632___mcc_h2
						var _1633___mcc_h3 _dafny.Array = _source15.Get_().(D3_DC11).Cf20
						_ = _1633___mcc_h3
						var _1634_cf20 _dafny.Array = _1633___mcc_h3
						_ = _1634_cf20
						var _1635_cf19 _dafny.Int = _1632___mcc_h2
						_ = _1635_cf19
						var _1636_cf18 _dafny.Map = _1631___mcc_h1
						_ = _1636_cf18
						_1635_cf19 = Companion_Default___.Fm9((_this).F18(), (_this).F18(), _1542_v0, globalState)
						var _1637_v70 _dafny.MultiSet
						_ = _1637_v70
						_1637_v70 = _dafny.MultiSetOf(_1635_cf19)
						var _1638_v71 _dafny.Set
						_ = _1638_v71
						_1638_v71 = _dafny.SetOf(_1637_v70)
						(globalState).F15 = (func() bool {
							if (_1622_v65).Contains(((_this).F18()).Plus(_1635_cf19)) {
								return (_1622_v65).Get(((_this).F18()).Plus(_1635_cf19)).(bool)
							}
							return (_1638_v71).IsSubsetOf(_1638_v71)
						})()
						_1613_v58 = _1613_v58
						_1542_v0 = _1542_v0
					} else {
						var _1639___mcc_h4 _dafny.Set = _source15.Get_().(D3_DC9).Cf16
						_ = _1639___mcc_h4
						var _1640_cf16 _dafny.Set = _1639___mcc_h4
						_ = _1640_cf16
						(globalState).F17 = _1542_v0
						var _1641_v72 T1
						_ = _1641_v72
						var _nw285 *C2 = New_C2_()
						_ = _nw285
						_nw285.Ctor__(_1612_v57, _1542_v0)
						_1641_v72 = _nw285
						var _1642_v73 _dafny.Array
						_ = _1642_v73
						var _nwElement0_66 T1 = _1641_v72
						_ = _nwElement0_66
						var _nw286 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(14))
						_ = _nw286
						(_nw286).ArraySet1(_nwElement0_66, 0)
						(_nw286).ArraySet1(_1641_v72, 1)
						(_nw286).ArraySet1(_1641_v72, 2)
						(_nw286).ArraySet1(_1641_v72, 3)
						(_nw286).ArraySet1(_1641_v72, 4)
						(_nw286).ArraySet1(_1641_v72, 5)
						(_nw286).ArraySet1(_1641_v72, 6)
						(_nw286).ArraySet1(_1641_v72, 7)
						(_nw286).ArraySet1(_1641_v72, 8)
						(_nw286).ArraySet1(_1641_v72, 9)
						(_nw286).ArraySet1(_1641_v72, 10)
						(_nw286).ArraySet1(_1641_v72, 11)
						(_nw286).ArraySet1(_1641_v72, 12)
						(_nw286).ArraySet1(_1641_v72, 13)
						_1642_v73 = _nw286
						var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1642_v73), 0))
						_ = _index230
						(_1642_v73).ArraySet1(_1641_v72, (_index230).Int())
						var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1642_v73), 0))
						_ = _index231
						(_1642_v73).ArraySet1(_1641_v72, (_index231).Int())
						(globalState).F15 = (func() bool {
							if _1542_v0 {
								return _1542_v0
							}
							return _1542_v0
						})()
						var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_1618_v63), 0))
						_ = _index232
						(_1618_v63).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1610_v55, _1610_v55)).Cardinality())), (_index232).Int())
						var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.ArrayLen((_1618_v63), 0))
						_ = _index233
						(_1618_v63).ArraySet1(Companion_Default___.SafeModuloInt((_this).F18(), (_this).F18()), (_index233).Int())
					}
					_1618_v63 = _1618_v63
					var _1643_v74 _dafny.Array
					_ = _1643_v74
					var _nwElement0_67 bool = _1542_v0
					_ = _nwElement0_67
					var _nw287 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(13))
					_ = _nw287
					(_nw287).ArraySet1(_nwElement0_67, 0)
					(_nw287).ArraySet1(_1542_v0, 1)
					(_nw287).ArraySet1(_1542_v0, 2)
					(_nw287).ArraySet1(_1542_v0, 3)
					(_nw287).ArraySet1(_1542_v0, 4)
					(_nw287).ArraySet1(_1542_v0, 5)
					(_nw287).ArraySet1(_1542_v0, 6)
					(_nw287).ArraySet1(_1542_v0, 7)
					(_nw287).ArraySet1(_1542_v0, 8)
					(_nw287).ArraySet1(_1542_v0, 9)
					(_nw287).ArraySet1(_1542_v0, 10)
					(_nw287).ArraySet1(!(_1542_v0), 11)
					(_nw287).ArraySet1(_1542_v0, 12)
					_1643_v74 = _nw287
					var _1644_v75 _dafny.Map
					_ = _1644_v75
					_1644_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
						if (_1616_v61).Contains(_1542_v0) {
							return (_1616_v61).Multiplicity(_1542_v0)
						}
						return (_this).F18()
					})(), _1643_v74)
					var _1645_v76 _dafny.Array
					_ = _1645_v76
					var _nw288 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(29))
					_ = _nw288
					_1645_v76 = _nw288
					var _1646_v77 *C4
					_ = _1646_v77
					var _nw289 *C4 = New_C4_()
					_ = _nw289
					_nw289.Ctor__(_1542_v0)
					_1646_v77 = _nw289
					var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1645_v76), 0))
					_ = _index234
					(_1645_v76).ArraySet1(_1646_v77, (_index234).Int())
					var _1647_v78 _dafny.Map
					_ = _1647_v78
					_1647_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1542_v0, _1613_v58)
					var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1645_v76), 0))
					_ = _index235
					var _rhs252 _dafny.Map = (_1644_v75).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _1643_v74))
					_ = _rhs252
					var _rhs253 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm16(Companion_Default___.Fm0(_1542_v0, (func() _dafny.CodePoint {
						if (_1647_v78).Contains((_1646_v77).F19()) {
							return (_1647_v78).Get((_1646_v77).F19()).(_dafny.CodePoint)
						}
						return _1613_v58
					})(), (_1646_v77).F19(), globalState), (_this).F18(), globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(432))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg74 _dafny.Int) interface{} {
							return coer74(arg74)
						}
					}((func(_1648_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1649_i10 _dafny.Int) _dafny.CodePoint {
							return _1648_v58
						}
					})(_1613_v58))))
					_ = _rhs253
					var _rhs254 _dafny.Int = (_dafny.Zero).Minus((_this).F18())
					_ = _rhs254
					var _rhs255 *C4 = _1646_v77
					_ = _rhs255
					var _lhs167 _dafny.Array = _1645_v76
					_ = _lhs167
					var _lhs168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1645_v76), 0))
					_ = _lhs168
					_1644_v75 = _rhs252
					_1542_v0 = _rhs253
					r0 = _rhs254
					(_lhs167).ArraySet1(_rhs255, (_lhs168).Int())
					var _1650_v79 _dafny.Map
					_ = _1650_v79
					_1650_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D4_.Create_DC13_((_this).F18()), _1542_v0)
					(globalState).F15 = !(!((p0).Union(p0)).Contains((_1650_v79).Cardinality()))
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1618_v63), 0))); ; {
			_guard_loop_3, _ok45 := _iter45()
			if !_ok45 {
				break
			}
			var _1651_i11 _dafny.Int
			_1651_i11 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1651_i11).Sign() != -1) && ((_1651_i11).Cmp(_dafny.ArrayLen((_1618_v63), 0)) < 0)) {
				(_1618_v63).ArraySet1(Companion_Default___.SafeModuloInt(_1651_i11, (Companion_D11_.Create_DC30_(false, (_this).F18(), _1542_v0)).Dtor_cf53()), (_1651_i11).Int())
			}
		}
		_1610_v55 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("up"), _dafny.UnicodeSeqOfUtf8Bytes("calngaa"))
		(globalState).F17 = !(!(_1542_v0)) || (_1542_v0)
		r0 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F18()), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F18(), (_this).F18())))
		return r0
	}
}
func (_this *C5) F18() _dafny.Int {
	{
		return _this._f18
	}
}

// End of class C5
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
