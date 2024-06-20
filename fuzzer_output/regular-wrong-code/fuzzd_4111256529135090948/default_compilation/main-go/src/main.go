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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	if true {
		return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-138), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())
	} else {
		return _dafny.IntOfInt64(541)
	}
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ko"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(296))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(78))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	})))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('s'))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _2_v0 _dafny.CodePoint
			_2_v0 = interface{}(_compr_0).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('s')), _2_v0) {
				_coll0.Add(_2_v0, _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ac")), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("n"), _dafny.UnicodeSeqOfUtf8Bytes("dubvfe"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(152))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg2 _dafny.Int) interface{} {
						return coer2(arg2)
					}
				}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				})), _dafny.UnicodeSeqOfUtf8Bytes("pt"), _dafny.UnicodeSeqOfUtf8Bytes("k"))))
			}
		}
		return _coll0.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm4(p0 D0, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false), _dafny.SeqOf(false)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true, false, false, true, false))))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Set, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.SetOf(!(!(false))), _dafny.SetOf(true, false, true, true, (Companion_D9_.Create_DC25_(true, _dafny.SetOf(_dafny.CodePoint('d')))).Dtor_cf41()), _dafny.SetOf(false), _dafny.SetOf(true))).Cardinality(), !(_dafny.MultiSetOf(!(true), false)).Equals(_dafny.MultiSetFromSeq(_dafny.SeqOf(true))))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('b')
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), !(false))))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 _dafny.Int, p2 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false, false, false, false, true)).Intersection(_dafny.MultiSetOf(false, false)), _dafny.CodePoint('l'))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	if false {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(671))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		}))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(273), (_dafny.MultiSetOf(!(false), false)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-33), func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(228), _dafny.IntOfInt64(940))); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _5_v0 _dafny.Int
				_5_v0 = interface{}(_compr_1).(_dafny.Int)
				if ((_dafny.IntOfInt64(228)).Cmp(_5_v0) <= 0) && ((_5_v0).Cmp(_dafny.IntOfInt64(940)) < 0) {
					_coll1.Add(Companion_Default___.SafeModuloInt(_5_v0, _dafny.IntOfInt64(160)), _dafny.IntOfInt64(791))
				}
			}
			return _coll1.ToMap()
		}()))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-777)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), _dafny.IntOfInt64(-658)))
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(459)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-480)))).Union(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(true)).Cardinality())), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(207))))))
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, p1 _dafny.Map, globalState *GlobalState) _dafny.MultiSet {
	return ((func() _dafny.MultiSet {
		if true {
			return _dafny.MultiSetOf(_dafny.SeqOf(true))
		}
		return _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqOf(true)))
	})()).Difference(_dafny.MultiSetOf(_dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	if (_dafny.IntOfInt64(988)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()) != 0 {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(478), _dafny.IntOfInt64(394))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(105), _dafny.IntOfInt64(451))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(251), _dafny.IntOfInt64(-748)))
	}
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	if (true) || (false) {
		return _dafny.MultiSetOf(_dafny.CodePoint('s'))
	} else {
		return _dafny.MultiSetOf(_dafny.CodePoint('c'), _dafny.CodePoint('o'), _dafny.CodePoint('i'))
	}
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.MultiSetOf(_dafny.CodePoint('f'), _dafny.CodePoint('q'), _dafny.CodePoint('y'))).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("kwrvsluv"), _dafny.UnicodeSeqOfUtf8Bytes("mtfd"))).Cardinality()), (_dafny.MultiSetFromSeq(_dafny.SeqOf((func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SeqOf(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(633))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_6_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		})))))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _7_v0 D0
			_7_v0 = interface{}(_compr_2).(D0)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(633))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_6_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			}))))), _7_v0) {
				_coll2.Add(_7_v0)
			}
		}
		return _coll2.ToSet()
	}()).Cardinality()))).Cardinality(), (_dafny.SetOf(false)).Cardinality()), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yrcmuol")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.IntOfInt64(175), _dafny.IntOfInt64(652), _dafny.IntOfInt64(-976), _dafny.IntOfInt64(896))).Union(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("b")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm22(p0 D3, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC12_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true), false), _dafny.SeqOf(!(true))))
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC5_((_dafny.Zero).Minus((_dafny.IntOfInt64(182)).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(357))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_8_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(true, !(!(false)))
	}))).Cardinality()))), (_dafny.IntOfInt64(-715)).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lucmkk")).Cardinality())) != 0, _dafny.IntOfInt64(-424), ((func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(673), _dafny.IntOfInt64(245))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _9_v0 _dafny.Int
			_9_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(673)).Cmp(_9_v0) <= 0) && ((_9_v0).Cmp(_dafny.IntOfInt64(245)) < 0) {
				_coll3.Add((_9_v0).Minus(_dafny.IntOfInt64(937)))
			}
		}
		return _coll3.ToSet()
	}()).Union(_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))).Cardinality())).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(329))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_10_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	}))).Cardinality()))).Cardinality()))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, p1 _dafny.Int, globalState *GlobalState) D9 {
	return Companion_D9_.Create_DC26_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("yglvl"), _dafny.UnicodeSeqOfUtf8Bytes("emwcyv"))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfInt64(-269)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(773)))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Set, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(_dafny.IntOfInt64(873), !(true), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(380))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Map {
	if !((_dafny.MultiSetFromSeq(_dafny.SeqOf(!(false), false, !(true)))).IsProperSubsetOf(_dafny.MultiSetOf(false, false))) {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.UnicodeSeqOfUtf8Bytes("ivfd"))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("jbqnyc"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("jjxglfc")))
	}
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), _dafny.IntOfInt64(868))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, false), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality())))).Merge(func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(true))).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _11_v0 _dafny.Sequence
			_11_v0 = interface{}(_compr_4).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.SeqOf(true))).Contains(_11_v0) {
				_coll4.Add(_11_v0, _dafny.IntOfInt64(200))
			}
		}
		return _coll4.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(171))).Uint32(), func(coer8 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_12_i0 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetOf(false, false, true)
	})), _dafny.SeqOf(_dafny.MultiSetOf(true))), _dafny.SeqOf(_dafny.MultiSetOf(!(false)), _dafny.MultiSetFromSeq(_dafny.SeqOf(true, false)), _dafny.MultiSetOf(true), _dafny.MultiSetOf(true), _dafny.MultiSetOf(true, false)))
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Int, p1 D8, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(false)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, globalState *GlobalState) D12 {
	if !(!(true)) {
		return Companion_D12_.Create_DC36_(true, false, _dafny.IntOfInt64(125), _dafny.IntOfInt64(710), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(674))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_13_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-438)
		}))))
	} else if !(true) {
		return Companion_D12_.Create_DC36_(false, true, _dafny.IntOfInt64(770), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(988))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(_dafny.IntOfInt64(735), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()), _dafny.IntOfInt64(-224))))
	} else {
		return Companion_D12_.Create_DC36_(false, !(false), _dafny.IntOfInt64(977), _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(-67), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Cardinality())).Cardinality()))))
	}
}
func (_static *CompanionStruct_Default___) Fm32(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D11_.Create_DC31_(_dafny.IntOfInt64(-732))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(681))).Uint32(), func(coer10 func(_dafny.Int) D11) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_14_i0 _dafny.Int) D11 {
		return Companion_D11_.Create_DC31_(_dafny.IntOfInt64(582))
	}))), _dafny.SeqOf(Companion_D11_.Create_DC31_((_dafny.MultiSetFromSeq(_dafny.SeqOf((Companion_D0_.Create_DC1_((_dafny.Zero).Minus((_dafny.SetOf(true, false)).Cardinality()), true, _dafny.IntOfInt64(333))).Dtor_cf2(), false))).Cardinality()), Companion_D11_.Create_DC31_((_dafny.SetOf(!(true))).Cardinality()), Companion_D11_.Create_DC31_(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), Companion_D11_.Create_DC31_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('l'))).Cardinality()))))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-70))).Uint32(), func(coer11 func(_dafny.Int) D11) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_15_i0 _dafny.Int) D11 {
		return Companion_D11_.Create_DC31_(_dafny.IntOfInt64(807))
	})), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aguorpcjs")).Cardinality()))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_D11_.Create_DC31_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("habj")).Cardinality()))), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nv")).Cardinality())), _dafny.IntOfInt64(200))))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_D11_.Create_DC31_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kdusdcx")).Cardinality()))), _dafny.SeqOf(_dafny.IntOfInt64(582)))).Merge(func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_D11_.Create_DC31_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("to")).Cardinality()))), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bffhds")).Cardinality()))).Keys().Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _16_v0 _dafny.Sequence
			_16_v0 = interface{}(_compr_5).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_D11_.Create_DC31_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("to")).Cardinality()))), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bffhds")).Cardinality()))).Contains(_16_v0) {
				_coll5.Add(_16_v0, _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-691))).Cardinality()))
			}
		}
		return _coll5.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.UnicodeSeqOfUtf8Bytes("mlwjdjdmf"))
}
func (_static *CompanionStruct_Default___) Fm35(p0 bool, p1 bool, p2 _dafny.MultiSet, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(!(!(true)), false)).Union(_dafny.SetOf(true))).Union(_dafny.SetOf(false, true))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Array, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (D0, _dafny.Map, _dafny.Int) {
	var r0 D0 = Companion_D0_.Default()
	_ = r0
	var r1 _dafny.Map = _dafny.EmptyMap
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	r2 = p2
	var _17_v0 _dafny.Array
	_ = _17_v0
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
	_ = _nw0
	_17_v0 = _nw0
	var _18_v1 bool
	_ = _18_v1
	_18_v1 = false
	var _19_v2 _dafny.Map
	_ = _19_v2
	_19_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_17_v0, _18_v1)
	var _20_v3 _dafny.Sequence
	_ = _20_v3
	_20_v3 = _dafny.SeqOf(p2)
	var _21_v4 D8
	_ = _21_v4
	_21_v4 = Companion_D8_.Create_DC20_(true, (_20_v3).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_20_v3).Cardinality()))).Uint32()).(_dafny.Int))
	var _22_v5 *C5
	_ = _22_v5
	var _nw1 *C5 = New_C5_()
	_ = _nw1
	_nw1.Ctor__((_19_v2).Cardinality(), Companion_Default___.SafeModuloInt(p2, p2), (_21_v4).Dtor_cf34())
	_22_v5 = _nw1
	var _23_v6 _dafny.Sequence
	_ = _23_v6
	_23_v6 = _dafny.SeqOf(p1)
	var _24_v7 _dafny.CodePoint
	_ = _24_v7
	_24_v7 = _dafny.CodePoint('x')
	var _25_v8 *C4
	_ = _25_v8
	var _nw2 *C4 = New_C4_()
	_ = _nw2
	_nw2.Ctor__(_17_v0, (_22_v5).F25())
	_25_v8 = _nw2
	var _26_v9 _dafny.Sequence
	_ = _26_v9
	_26_v9 = _dafny.SeqOf(_25_v8, _25_v8)
	var _27_v10 _dafny.MultiSet
	_ = _27_v10
	_27_v10 = _dafny.MultiSetOf((_26_v9).Select((Companion_Default___.SafeIndex((_22_v5).F25(), _dafny.IntOfUint32((_26_v9).Cardinality()))).Uint32()).(*C4))
	var _28_v11 _dafny.Sequence
	_ = _28_v11
	_28_v11 = _dafny.SeqOf(Companion_Default___.Fm1(_dafny.IntOfInt64(895), _18_v1, globalState))
	var _29_v12 _dafny.Set
	_ = _29_v12
	_29_v12 = _dafny.SetOf(p2)
	var _30_v13 _dafny.Map
	_ = _30_v13
	_30_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_22_v5).F25(), (_22_v5).F26())
	var _rhs0 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex((Companion_D8_.Create_DC21_(_18_v1, _18_v1, (_dafny.Zero).Minus(_dafny.IntOfUint32(((_23_v6).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_23_v6).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())))).Dtor_cf37(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _24_v7)).Cardinality())).Cmp(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus((_27_v10).Cardinality())), _dafny.IntOfUint32((_28_v11).Cardinality()))) < 0
	_ = _rhs0
	var _rhs1 _dafny.Int = (_22_v5).F25()
	_ = _rhs1
	var _rhs2 bool = (_29_v12).IsSubsetOf(_29_v12)
	_ = _rhs2
	var _rhs3 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_30_v13).Cardinality()), p2)
	_ = _rhs3
	var _lhs0 *GlobalState = globalState
	_ = _lhs0
	var _lhs1 *GlobalState = globalState
	_ = _lhs1
	_lhs0.F7 = _rhs0
	r2 = _rhs1
	_18_v1 = _rhs2
	_lhs1.F8 = _rhs3
	var _31_v14 _dafny.Map
	_ = _31_v14
	_31_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_18_v1, _18_v1)
	var _32_v15 *C0
	_ = _32_v15
	var _nw3 *C0 = New_C0_()
	_ = _nw3
	_nw3.Ctor__()
	_32_v15 = _nw3
	var _33_v16 _dafny.Map
	_ = _33_v16
	_33_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v15, _dafny.IntOfInt64(797))
	var _34_v17 _dafny.Map
	_ = _34_v17
	_34_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_33_v16, _31_v14)
	var _35_v18 _dafny.Array
	_ = _35_v18
	var _nwElement0_0 _dafny.Map = _31_v14
	_ = _nwElement0_0
	var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(7))
	_ = _nw4
	(_nw4).ArraySet1(_nwElement0_0, 0)
	(_nw4).ArraySet1(_31_v14, 1)
	(_nw4).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_18_v1), _18_v1), 2)
	(_nw4).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_18_v1), _18_v1), 3)
	(_nw4).ArraySet1((func() _dafny.Map {
		if (_34_v17).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v15, (_22_v5).F26())) {
			return (_34_v17).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v15, (_22_v5).F26())).(_dafny.Map)
		}
		return _31_v14
	})(), 4)
	(_nw4).ArraySet1(_31_v14, 5)
	(_nw4).ArraySet1(_31_v14, 6)
	_35_v18 = _nw4
	for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_35_v18), 0))); ; {
		_guard_loop_0, _ok6 := _iter6()
		if !_ok6 {
			break
		}
		var _36_i0 _dafny.Int
		_36_i0 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_36_i0).Sign() != -1) && ((_36_i0).Cmp(_dafny.ArrayLen((_35_v18), 0)) < 0)) {
			(_35_v18).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _18_v1)).Merge((Companion_D19_.Create_DC52_(_31_v14)).Dtor_cf85()), (_36_i0).Int())
		}
	}
	var _hi0 _dafny.Int = p2
	_ = _hi0
	for _37_i1 := (_22_v5).F25(); _37_i1.Cmp(_hi0) < 0; _37_i1 = _37_i1.Plus(_dafny.One) {
		var _38_v19 _dafny.Array
		_ = _38_v19
		var _nw5 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
		_ = _nw5
		_38_v19 = _nw5
		_38_v19 = _38_v19
		r2 = Companion_Default___.Fm0((_22_v5).F25(), _18_v1, globalState)
		(globalState).F7 = _18_v1
		(globalState).F6 = p1
	}
	_18_v1 = _18_v1
	var _39_v20 *C2
	_ = _39_v20
	var _nw6 *C2 = New_C2_()
	_ = _nw6
	_nw6.Ctor__(_30_v13, _18_v1, _18_v1, _20_v3, p2)
	_39_v20 = _nw6
	var _40_v21 _dafny.Array
	_ = _40_v21
	var _nwElement0_1 *C2 = _39_v20
	_ = _nwElement0_1
	var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.One)
	_ = _nw7
	(_nw7).ArraySet1(_nwElement0_1, 0)
	_40_v21 = _nw7
	var _41_v22 _dafny.MultiSet
	_ = _41_v22
	_41_v22 = _dafny.MultiSetOf(_40_v21, _40_v21, _40_v21, _40_v21)
	var _42_v23 D20
	_ = _42_v23
	_42_v23 = Companion_D20_.Create_DC55_(_41_v22)
	var _43_v24 D0
	_ = _43_v24
	_43_v24 = Companion_D0_.Create_DC1_(Companion_Default___.Fm0(p2, false, globalState), (_41_v22).IsSubsetOf((_42_v23).Dtor_cf89()), (_22_v5).F26())
	r0 = _43_v24
	var _44_v25 _dafny.Map
	_ = _44_v25
	_44_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_18_v1, _20_v3)
	var _45_v26 D12
	_ = _45_v26
	_45_v26 = Companion_D12_.Create_DC36_(!((_39_v20).F32()), _18_v1, _dafny.IntOfUint32((p1).Cardinality()), p2, _44_v25)
	var _pat_let_tv0 = _39_v20
	_ = _pat_let_tv0
	var _pat_let_tv1 = _22_v5
	_ = _pat_let_tv1
	var _pat_let_tv2 = _39_v20
	_ = _pat_let_tv2
	var _pat_let_tv3 = _24_v7
	_ = _pat_let_tv3
	var _pat_let_tv4 = _24_v7
	_ = _pat_let_tv4
	var _pat_let_tv5 = _39_v20
	_ = _pat_let_tv5
	var _pat_let_tv6 = _22_v5
	_ = _pat_let_tv6
	var _pat_let_tv7 = p1
	_ = _pat_let_tv7
	var _pat_let_tv8 = _18_v1
	_ = _pat_let_tv8
	var _pat_let_tv9 = _22_v5
	_ = _pat_let_tv9
	var _pat_let_tv10 = _22_v5
	_ = _pat_let_tv10
	var _pat_let_tv11 = _18_v1
	_ = _pat_let_tv11
	r1 = func(_source0 D12) _dafny.Map {
		if _source0.Is_DC35() {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((_pat_let_tv0).F31()).Cardinality()), true)).Update((_pat_let_tv1).F26(), (_pat_let_tv2).F32())
		} else if _source0.Is_DC36() {
			var _46___mcc_h0 bool = _source0.Get_().(D12_DC36).Cf56
			_ = _46___mcc_h0
			var _47___mcc_h1 bool = _source0.Get_().(D12_DC36).Cf57
			_ = _47___mcc_h1
			var _48___mcc_h2 _dafny.Int = _source0.Get_().(D12_DC36).Cf58
			_ = _48___mcc_h2
			var _49___mcc_h3 _dafny.Int = _source0.Get_().(D12_DC36).Cf59
			_ = _49___mcc_h3
			var _50___mcc_h4 _dafny.Map = _source0.Get_().(D12_DC36).Cf60
			_ = _50___mcc_h4
			var _51_cf60 _dafny.Map = _50___mcc_h4
			_ = _51_cf60
			var _52_cf59 _dafny.Int = _49___mcc_h3
			_ = _52_cf59
			var _53_cf58 _dafny.Int = _48___mcc_h2
			_ = _53_cf58
			var _54_cf57 bool = _47___mcc_h1
			_ = _54_cf57
			var _55_cf56 bool = _46___mcc_h0
			_ = _55_cf56
			return (func() _dafny.Map {
				var _coll6 = _dafny.NewMapBuilder()
				_ = _coll6
				for _iter7 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(842))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg12 _dafny.Int) interface{} {
						return coer12(arg12)
					}
				}((func(_56_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_57_i2 _dafny.Int) _dafny.CodePoint {
						return _56_v7
					}
				})(_pat_let_tv3)))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-950))).Cardinality()))).Elements()); ; {
					_compr_6, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _58_v27 _dafny.Int
					_58_v27 = interface{}(_compr_6).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(842))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg13 _dafny.Int) interface{} {
							return coer13(arg13)
						}
					}((func(_59_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_57_i2 _dafny.Int) _dafny.CodePoint {
							return _59_v7
						}
					})(_pat_let_tv4)))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-950))).Cardinality())), _58_v27) {
						_coll6.Add((_58_v27).Times((_pat_let_tv6).F26()), (_pat_let_tv5).F32())
					}
				}
				return _coll6.ToMap()
			}()).Merge((func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-214), _dafny.IntOfInt64(460))); ; {
					_compr_7, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _60_v28 _dafny.Int
					_60_v28 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(-214)).Cmp(_60_v28) <= 0) && ((_60_v28).Cmp(_dafny.IntOfInt64(460)) < 0) {
						_coll7.Add(Companion_Default___.SafeDivisionInt(_60_v28, _53_cf58), _55_cf56)
					}
				}
				return _coll7.ToMap()
			}()).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_pat_let_tv7).Cardinality()))).Cardinality(), _54_cf57))
		} else if _source0.Is_DC34() {
			var _61___mcc_h5 _dafny.Map = _source0.Get_().(D12_DC34).Cf55
			_ = _61___mcc_h5
			var _62_cf55 _dafny.Map = _61___mcc_h5
			_ = _62_cf55
			return func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(989), _dafny.IntOfInt64(320))); ; {
					_compr_8, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _63_v29 _dafny.Int
					_63_v29 = interface{}(_compr_8).(_dafny.Int)
					if ((_dafny.IntOfInt64(989)).Cmp(_63_v29) <= 0) && ((_63_v29).Cmp(_dafny.IntOfInt64(320)) < 0) {
						_coll8.Add(Companion_Default___.SafeModuloInt(_63_v29, (_pat_let_tv9).F25()), _pat_let_tv8)
					}
				}
				return _coll8.ToMap()
			}()
		} else {
			var _64___mcc_h6 D12 = _source0.Get_().(D12_DC37).Cf61
			_ = _64___mcc_h6
			var _65_cf61 D12 = _64___mcc_h6
			_ = _65_cf61
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_pat_let_tv10).F25(), _pat_let_tv11)
		}
	}(_45_v26)
	r2 = (_22_v5).F25()
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _66_v0 _dafny.Sequence
	_ = _66_v0
	_66_v0 = _dafny.UnicodeSeqOfUtf8Bytes("hg")
	var _67_v1 D0
	_ = _67_v1
	_67_v1 = Companion_D0_.Create_DC0_(_66_v0)
	var _68_v2 _dafny.CodePoint
	_ = _68_v2
	_68_v2 = _dafny.CodePoint('t')
	var _69_v3 _dafny.Array
	_ = _69_v3
	var _nwElement0_2 _dafny.Sequence = _66_v0
	_ = _nwElement0_2
	var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(12))
	_ = _nw8
	(_nw8).ArraySet1(_nwElement0_2, 0)
	(_nw8).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mexrhnq"), 1)
	(_nw8).ArraySet1(_66_v0, 2)
	(_nw8).ArraySet1(_66_v0, 3)
	(_nw8).ArraySet1(_66_v0, 4)
	(_nw8).ArraySet1((_67_v1).Dtor_cf0(), 5)
	(_nw8).ArraySet1(_66_v0, 6)
	(_nw8).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qu"), 7)
	(_nw8).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("wggsbwf"), 8)
	(_nw8).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rnjar"), 9)
	(_nw8).ArraySet1(_66_v0, 10)
	(_nw8).ArraySet1(_dafny.Companion_Sequence_.Update(_66_v0, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-282), _dafny.IntOfUint32((_66_v0).Cardinality()))).Uint32(), _68_v2), 11)
	_69_v3 = _nw8
	var _70_v4 bool
	_ = _70_v4
	_70_v4 = true
	var _71_v5 _dafny.MultiSet
	_ = _71_v5
	_71_v5 = _dafny.MultiSetOf(_70_v4, true, _70_v4)
	var _72_v6 _dafny.MultiSet
	_ = _72_v6
	_72_v6 = _dafny.MultiSetOf(_71_v5)
	var _73_v7 _dafny.Array
	_ = _73_v7
	var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
	_ = _nw9
	_73_v7 = _nw9
	var _74_v8 _dafny.Array
	_ = _74_v8
	var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
	_ = _nw10
	_74_v8 = _nw10
	var _75_v9 _dafny.Int
	_ = _75_v9
	_75_v9 = _dafny.IntOfInt64(-20)
	var _76_v10 _dafny.Map
	_ = _76_v10
	_76_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_74_v8, _75_v9)
	var _77_globalState *GlobalState
	_ = _77_globalState
	var _nw11 *GlobalState = New_GlobalState_()
	_ = _nw11
	_nw11.Ctor__(true, _dafny.IntOfInt64(907), _69_v3, _dafny.IntOfInt64(-486), false, _72_v6, _66_v0, false, _dafny.IntOfInt64(-977), false, _dafny.IntOfInt64(632), _dafny.CodePoint('f'), _dafny.IntOfInt64(179), _73_v7, _dafny.IntOfInt64(435), _dafny.IntOfInt64(814), true, _dafny.IntOfInt64(558), (_76_v10).Merge(_76_v10), false, _74_v8, false, _dafny.IntOfInt64(634))
	_77_globalState = _nw11
	var _78_i0 _dafny.Int
	_ = _78_i0
	_78_i0 = _dafny.Zero
	{
		for _70_v4 {
			{
				if (_78_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_78_i0 = (_78_i0).Plus(_dafny.One)
				var _79_v11 _dafny.Array
				_ = _79_v11
				var _nw12 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
				_ = _nw12
				_79_v11 = _nw12
				var _80_v12 D0
				_ = _80_v12
				var _81_v13 _dafny.Map
				_ = _81_v13
				var _82_v14 _dafny.Int
				_ = _82_v14
				var _out0 D0
				_ = _out0
				var _out1 _dafny.Map
				_ = _out1
				var _out2 _dafny.Int
				_ = _out2
				_out0, _out1, _out2 = Companion_Default___.M0(_79_v11, _66_v0, _75_v9, _77_globalState)
				_80_v12 = _out0
				_81_v13 = _out1
				_82_v14 = _out2
				var _83_v15 _dafny.Map
				_ = _83_v15
				_83_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_75_v9, _dafny.IntOfInt64(921))
				var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_74_v8), 0))
				_ = _index0
				(_74_v8).ArraySet1((func() _dafny.Int {
					if (_83_v15).Contains(_dafny.IntOfUint32((_66_v0).Cardinality())) {
						return (_83_v15).Get(_dafny.IntOfUint32((_66_v0).Cardinality())).(_dafny.Int)
					}
					return _dafny.IntOfInt64(973)
				})(), (_index0).Int())
				var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_74_v8), 0))
				_ = _index1
				(_74_v8).ArraySet1(Companion_Default___.Fm0(_75_v9, _70_v4, _77_globalState), (_index1).Int())
				var _84_v16 _dafny.Array
				_ = _84_v16
				var _nwElement0_3 _dafny.Array = _69_v3
				_ = _nwElement0_3
				var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(22))
				_ = _nw13
				(_nw13).ArraySet1(_nwElement0_3, 0)
				(_nw13).ArraySet1(_69_v3, 1)
				(_nw13).ArraySet1(_69_v3, 2)
				(_nw13).ArraySet1(_69_v3, 3)
				(_nw13).ArraySet1(_69_v3, 4)
				(_nw13).ArraySet1((func() _dafny.Array {
					if _70_v4 {
						return _69_v3
					}
					return _69_v3
				})(), 5)
				(_nw13).ArraySet1((func() _dafny.Array {
					if _70_v4 {
						return _69_v3
					}
					return _69_v3
				})(), 6)
				(_nw13).ArraySet1(_69_v3, 7)
				(_nw13).ArraySet1(_69_v3, 8)
				(_nw13).ArraySet1(_69_v3, 9)
				(_nw13).ArraySet1(_69_v3, 10)
				(_nw13).ArraySet1(_69_v3, 11)
				(_nw13).ArraySet1(_69_v3, 12)
				(_nw13).ArraySet1(_69_v3, 13)
				(_nw13).ArraySet1(_69_v3, 14)
				(_nw13).ArraySet1(_69_v3, 15)
				(_nw13).ArraySet1(_69_v3, 16)
				(_nw13).ArraySet1(_69_v3, 17)
				(_nw13).ArraySet1(_69_v3, 18)
				(_nw13).ArraySet1(_69_v3, 19)
				(_nw13).ArraySet1(_69_v3, 20)
				(_nw13).ArraySet1(_69_v3, 21)
				_84_v16 = _nw13
				var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_84_v16), 0))
				_ = _index2
				(_84_v16).ArraySet1(_69_v3, (_index2).Int())
				var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_84_v16), 0))
				_ = _index3
				(_84_v16).ArraySet1(_69_v3, (_index3).Int())
				var _85_v17 _dafny.Sequence
				_ = _85_v17
				_85_v17 = _dafny.SeqOf(_70_v4, (_dafny.IntOfUint32((_66_v0).Cardinality())).Cmp(_82_v14) > 0, _70_v4)
				_85_v17 = _85_v17
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _86_v18 _dafny.Array
	_ = _86_v18
	var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
	_ = _nw14
	_86_v18 = _nw14
	var _87_v19 _dafny.Map
	_ = _87_v19
	_87_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(378), (_dafny.MultiSetOf(_75_v9)).Cardinality())
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_86_v18), 0))
	_ = _index4
	(_86_v18).ArraySet1(_87_v19, (_index4).Int())
	var _88_v21 _dafny.Sequence
	_ = _88_v21
	_88_v21 = _dafny.SeqOf(_dafny.IntOfInt64(467))
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_86_v18), 0))
	_ = _index5
	(_86_v18).ArraySet1(((_87_v19).Update(_dafny.IntOfInt64(688), _75_v9)).Merge((func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter10 := _dafny.Iterate((_88_v21).Elements()); ; {
			_compr_9, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _89_v20 _dafny.Int
			_89_v20 = interface{}(_compr_9).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_88_v21, _89_v20) {
				_coll9.Add((_89_v20).Minus(_dafny.IntOfInt64(732)), _75_v9)
			}
		}
		return _coll9.ToMap()
	}()).Merge(_87_v19)), (_index5).Int())
	var _90_v22 _dafny.Array
	_ = _90_v22
	var _nw15 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
	_ = _nw15
	_90_v22 = _nw15
	var _91_v23 _dafny.Set
	_ = _91_v23
	_91_v23 = _dafny.SetOf(_70_v4, _70_v4)
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))
	_ = _index6
	(_90_v22).ArraySet1((_dafny.SetOf(_70_v4, _70_v4, _70_v4, _70_v4)).IsProperSubsetOf(_91_v23), (_index6).Int())
	var _92_v24 _dafny.Map
	_ = _92_v24
	_92_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v2, Companion_Default___.Fm1(_75_v9, _70_v4, _77_globalState))
	var _93_v25 _dafny.MultiSet
	_ = _93_v25
	_93_v25 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("jxuh"), Companion_Default___.Fm2(_77_globalState), _66_v0)
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))
	_ = _index7
	var _rhs4 bool = !(_70_v4) || (!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v2, !(!(_70_v4)))).Equals(_92_v24))
	_ = _rhs4
	var _rhs5 _dafny.Int = (_75_v9).Times((func() _dafny.Int {
		if (_93_v25).Contains(_66_v0) {
			return (_93_v25).Multiplicity(_66_v0)
		}
		return _75_v9
	})())
	_ = _rhs5
	var _rhs6 _dafny.Int = Companion_Default___.SafeModuloInt(_75_v9, _75_v9)
	_ = _rhs6
	var _rhs7 bool = !(true)
	_ = _rhs7
	var _lhs2 *GlobalState = _77_globalState
	_ = _lhs2
	var _lhs3 *GlobalState = _77_globalState
	_ = _lhs3
	var _lhs4 *GlobalState = _77_globalState
	_ = _lhs4
	var _lhs5 _dafny.Array = _90_v22
	_ = _lhs5
	var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))
	_ = _lhs6
	_lhs2.F7 = _rhs4
	_lhs3.F14 = _rhs5
	_lhs4.F14 = _rhs6
	(_lhs5).ArraySet1(_rhs7, (_lhs6).Int())
	var _94_v26 _dafny.Sequence
	_ = _94_v26
	_94_v26 = _dafny.SeqOf(!(_70_v4))
	var _hi1 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
		if true {
			return _dafny.Companion_Sequence_.Update(_94_v26, (Companion_Default___.SafeIndex(_75_v9, _dafny.IntOfUint32((_94_v26).Cardinality()))).Uint32(), (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool))
		}
		return _94_v26
	})()).Cardinality())
	_ = _hi1
	for _95_i1 := _75_v9; _95_i1.Cmp(_hi1) < 0; _95_i1 = _95_i1.Plus(_dafny.One) {
		(_77_globalState).F14 = _75_v9
		var _96_v27 _dafny.Array
		_ = _96_v27
		var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
		_ = _nw16
		_96_v27 = _nw16
		var _97_v28 _dafny.Map
		_ = _97_v28
		_97_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_95_i1, _96_v27)
		_97_v28 = (_97_v28).Update(Companion_Default___.SafeModuloInt(_95_i1, _75_v9), _96_v27)
		_70_v4 = (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool)
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_74_v8), 0))
		_ = _index8
		(_74_v8).ArraySet1(_75_v9, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_74_v8), 0))
		_ = _index9
		(_74_v8).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.SetOf(_70_v4)).Cardinality(), (_dafny.IntOfUint32((_88_v21).Cardinality())).Plus(_75_v9)), (_index9).Int())
	}
	var _98_i2 _dafny.Int
	_ = _98_i2
	_98_i2 = _dafny.Zero
	{
		for ((_75_v9).Times(_75_v9)).Cmp(Companion_Default___.SafeDivisionInt(_75_v9, (_dafny.Zero).Minus(_75_v9))) != 0 {
			{
				if (_98_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_98_i2 = (_98_i2).Plus(_dafny.One)
				(_77_globalState).F8 = _dafny.IntOfUint32((_dafny.SeqOf((_91_v23).Union(_91_v23), _91_v23, _91_v23, _91_v23)).Cardinality())
				var _99_v29 D0
				_ = _99_v29
				var _100_v30 _dafny.Map
				_ = _100_v30
				var _101_v31 _dafny.Int
				_ = _101_v31
				var _out3 D0
				_ = _out3
				var _out4 _dafny.Map
				_ = _out4
				var _out5 _dafny.Int
				_ = _out5
				_out3, _out4, _out5 = Companion_Default___.M0(_90_v22, _66_v0, _75_v9, _77_globalState)
				_99_v29 = _out3
				_100_v30 = _out4
				_101_v31 = _out5
				var _102_v32 _dafny.Sequence
				_ = _102_v32
				_102_v32 = _dafny.SeqOf(Companion_Default___.Fm3(_75_v9, _77_globalState), _92_v24, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v2, _70_v4), _92_v24, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v2, _70_v4))
				var _103_v33 D0
				_ = _103_v33
				var _104_v34 _dafny.Map
				_ = _104_v34
				var _105_v35 _dafny.Int
				_ = _105_v35
				var _out6 D0
				_ = _out6
				var _out7 _dafny.Map
				_ = _out7
				var _out8 _dafny.Int
				_ = _out8
				_out6, _out7, _out8 = Companion_Default___.M0(_90_v22, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(751))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg14 _dafny.Int) interface{} {
						return coer14(arg14)
					}
				}((func(_106_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_107_i3 _dafny.Int) _dafny.CodePoint {
						return _106_v2
					}
				})(_68_v2))), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_102_v32).Cardinality()), _101_v31)).Cardinality()).Times((Companion_Default___.Fm4(Companion_D0_.Create_DC0_(_dafny.UnicodeSeqOfUtf8Bytes("k")), _dafny.IntOfInt64(998), _94_v26, _68_v2, _77_globalState)).Cardinality()), _77_globalState)
				_103_v33 = _out6
				_104_v34 = _out7
				_105_v35 = _out8
				var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_74_v8), 0))
				_ = _index10
				(_74_v8).ArraySet1(_dafny.IntOfInt64(-24), (_index10).Int())
				var _108_v36 _dafny.Map
				_ = _108_v36
				_108_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool), _74_v8)
				var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_74_v8), 0))
				_ = _index11
				(_74_v8).ArraySet1(Companion_Default___.Fm0(_101_v31, (_75_v9).Cmp((_108_v36).Cardinality()) < 0, _77_globalState), (_index11).Int())
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	(_77_globalState).F6 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ahdth"), _dafny.UnicodeSeqOfUtf8Bytes("syrkxpe"))
	(_77_globalState).F14 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((func() _dafny.Int {
		if _70_v4 {
			return _75_v9
		}
		return _75_v9
	})(), (_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm2(_77_globalState)).Cardinality()))))
	if (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool) {
		(_77_globalState).F7 = _70_v4
		(_77_globalState).F8 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_66_v0).Cardinality()), (_75_v9).Plus(_dafny.IntOfInt64(30)))
		var _109_v37 _dafny.Set
		_ = _109_v37
		_109_v37 = _dafny.SetOf(_75_v9)
		var _110_v38 _dafny.Sequence
		_ = _110_v38
		_110_v38 = _dafny.SeqOf(Companion_Default___.Fm5(_109_v37, _77_globalState))
		(_77_globalState).F8 = Companion_Default___.SafeModuloInt(_75_v9, Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_66_v0).Cardinality())), _dafny.IntOfUint32((_110_v38).Cardinality())))
		var _111_v39 _dafny.Map
		_ = _111_v39
		_111_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v4, _75_v9)
		var _112_v40 D0
		_ = _112_v40
		_112_v40 = Companion_D0_.Create_DC1_(_75_v9, (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool), _75_v9)
		_111_v39 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_66_v0, (Companion_Default___.SafeIndex(_75_v9, _dafny.IntOfUint32((_66_v0).Cardinality()))).Uint32(), _68_v2)).Cardinality()))).Update((_112_v40).Dtor_cf2(), _dafny.IntOfInt64(-355))).Merge(_111_v39)
		_88_v21 = _88_v21
	} else {
		var _113_v41 _dafny.MultiSet
		_ = _113_v41
		_113_v41 = _dafny.MultiSetOf(_75_v9, _dafny.IntOfUint32((_88_v21).Cardinality()))
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))
		_ = _index12
		var _rhs8 _dafny.Int = _75_v9
		_ = _rhs8
		var _rhs9 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_66_v0, (Companion_Default___.SafeIndex(_75_v9, _dafny.IntOfUint32((_66_v0).Cardinality()))).Uint32(), _68_v2), (Companion_Default___.SafeIndex(_75_v9, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_66_v0, (Companion_Default___.SafeIndex(_75_v9, _dafny.IntOfUint32((_66_v0).Cardinality()))).Uint32(), _68_v2)).Cardinality()))).Uint32(), _68_v2)
		_ = _rhs9
		var _rhs10 bool = _70_v4
		_ = _rhs10
		var _rhs11 bool = (_113_v41).IsSubsetOf(_113_v41)
		_ = _rhs11
		var _rhs12 bool = (_94_v26).Select((Companion_Default___.SafeIndex((_dafny.IntOfUint32((_94_v26).Cardinality())).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm2(_77_globalState), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_88_v21).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm2(_77_globalState)).Cardinality()))).Uint32(), _68_v2)).Cardinality())), _dafny.IntOfUint32((_94_v26).Cardinality()))).Uint32()).(bool)
		_ = _rhs12
		var _lhs7 *GlobalState = _77_globalState
		_ = _lhs7
		var _lhs8 *GlobalState = _77_globalState
		_ = _lhs8
		var _lhs9 _dafny.Array = _90_v22
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))
		_ = _lhs10
		var _lhs11 *GlobalState = _77_globalState
		_ = _lhs11
		_lhs7.F14 = _rhs8
		_lhs8.F6 = _rhs9
		(_lhs9).ArraySet1(_rhs10, (_lhs10).Int())
		_70_v4 = _rhs11
		_lhs11.F21 = _rhs12
		(_77_globalState).F11 = Companion_Default___.Fm6(_75_v9, _77_globalState)
		(_77_globalState).F9 = (Companion_D0_.Create_DC1_(_75_v9, (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool), _75_v9)).Dtor_cf2()
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))
		_ = _index13
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))
		_ = _index14
		var _rhs13 bool = (func() bool {
			if (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool) {
				return (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool)
			}
			return _70_v4
		})()
		_ = _rhs13
		var _rhs14 bool = _70_v4
		_ = _rhs14
		var _rhs15 bool = (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool)
		_ = _rhs15
		var _lhs12 *GlobalState = _77_globalState
		_ = _lhs12
		var _lhs13 _dafny.Array = _90_v22
		_ = _lhs13
		var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))
		_ = _lhs14
		var _lhs15 _dafny.Array = _90_v22
		_ = _lhs15
		var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))
		_ = _lhs16
		_lhs12.F7 = _rhs13
		(_lhs13).ArraySet1(_rhs14, (_lhs14).Int())
		(_lhs15).ArraySet1(_rhs15, (_lhs16).Int())
		var _114_v42 D0
		_ = _114_v42
		var _115_v43 _dafny.Map
		_ = _115_v43
		var _116_v44 _dafny.Int
		_ = _116_v44
		var _out9 D0
		_ = _out9
		var _out10 _dafny.Map
		_ = _out10
		var _out11 _dafny.Int
		_ = _out11
		_out9, _out10, _out11 = Companion_Default___.M0(_90_v22, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-195))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}((func(_117_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_118_i4 _dafny.Int) _dafny.CodePoint {
				return _117_v2
			}
		})(_68_v2))), (_dafny.Zero).Minus(_75_v9), _77_globalState)
		_114_v42 = _out9
		_115_v43 = _out10
		_116_v44 = _out11
	}
	var _hi2 _dafny.Int = _75_v9
	_ = _hi2
	for _119_i5 := _75_v9; _119_i5.Cmp(_hi2) < 0; _119_i5 = _119_i5.Plus(_dafny.One) {
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))
		_ = _index15
		(_90_v22).ArraySet1(!_dafny.Companion_Sequence_.Equal(_66_v0, _66_v0), (_index15).Int())
		(_77_globalState).F21 = _70_v4
		(_77_globalState).F8 = (_119_i5).Plus(Companion_Default___.Fm0(_dafny.IntOfInt64(892), _70_v4, _77_globalState))
		var _120_v45 D0
		_ = _120_v45
		var _121_v46 _dafny.Map
		_ = _121_v46
		var _122_v47 _dafny.Int
		_ = _122_v47
		var _out12 D0
		_ = _out12
		var _out13 _dafny.Map
		_ = _out13
		var _out14 _dafny.Int
		_ = _out14
		_out12, _out13, _out14 = Companion_Default___.M0((func() _dafny.Array {
			if (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool) {
				return _90_v22
			}
			return _90_v22
		})(), _66_v0, _119_i5, _77_globalState)
		_120_v45 = _out12
		_121_v46 = _out13
		_122_v47 = _out14
	}
	var _123_i6 _dafny.Int
	_ = _123_i6
	_123_i6 = _dafny.Zero
	{
		for (_75_v9).Cmp(_dafny.IntOfInt64(-886)) >= 0 {
			{
				if (_123_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_123_i6 = (_123_i6).Plus(_dafny.One)
				var _124_v48 _dafny.Map
				_ = _124_v48
				_124_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_75_v9, _70_v4)
				var _rhs16 _dafny.Array = _74_v8
				_ = _rhs16
				var _rhs17 _dafny.Int = ((_124_v48).Cardinality()).Minus((_dafny.Zero).Minus(_75_v9))
				_ = _rhs17
				var _rhs18 bool = (Companion_Default___.Fm0(_75_v9, _70_v4, _77_globalState)).Cmp(_dafny.IntOfInt64(743)) == 0
				_ = _rhs18
				var _rhs19 bool = (func() bool {
					if _70_v4 {
						return _70_v4
					}
					return _70_v4
				})()
				_ = _rhs19
				var _rhs20 _dafny.Int = _75_v9
				_ = _rhs20
				var _lhs17 *GlobalState = _77_globalState
				_ = _lhs17
				var _lhs18 *GlobalState = _77_globalState
				_ = _lhs18
				var _lhs19 *GlobalState = _77_globalState
				_ = _lhs19
				var _lhs20 *GlobalState = _77_globalState
				_ = _lhs20
				_74_v8 = _rhs16
				_lhs17.F14 = _rhs17
				_lhs18.F21 = _rhs18
				_lhs19.F7 = _rhs19
				_lhs20.F14 = _rhs20
				(_77_globalState).F9 = (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool)
				var _125_v49 _dafny.Sequence
				_ = _125_v49
				_125_v49 = _dafny.SeqOf(_dafny.SeqOf((_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool)), _94_v26)
				(_77_globalState).F14 = Companion_Default___.SafeModuloInt(_75_v9, (_75_v9).Minus(_dafny.IntOfUint32(((_125_v49).Select((Companion_Default___.SafeIndex(_75_v9, _dafny.IntOfUint32((_125_v49).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())))
				var _126_v50 _dafny.Array
				_ = _126_v50
				var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(13))
				_ = _nw17
				_126_v50 = _nw17
				var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_126_v50), 0))
				_ = _index16
				(_126_v50).ArraySet1CodePoint(_68_v2, (_index16).Int())
				var _127_v51 _dafny.MultiSet
				_ = _127_v51
				_127_v51 = _dafny.MultiSetOf(_75_v9)
				var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_126_v50), 0))
				_ = _index17
				var _rhs21 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ghjmshpn"), _66_v0)
				_ = _rhs21
				var _rhs22 _dafny.CodePoint = (_66_v0).Select((Companion_Default___.SafeIndex((_127_v51).Cardinality(), _dafny.IntOfUint32((_66_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs22
				var _rhs23 _dafny.CodePoint = _68_v2
				_ = _rhs23
				var _lhs21 *GlobalState = _77_globalState
				_ = _lhs21
				var _lhs22 *GlobalState = _77_globalState
				_ = _lhs22
				var _lhs23 _dafny.Array = _126_v50
				_ = _lhs23
				var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_126_v50), 0))
				_ = _lhs24
				_lhs21.F6 = _rhs21
				_lhs22.F11 = _rhs22
				(_lhs23).ArraySet1CodePoint(_rhs23, (_lhs24).Int())
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _128_v52 D0
	_ = _128_v52
	var _129_v53 _dafny.Map
	_ = _129_v53
	var _130_v54 _dafny.Int
	_ = _130_v54
	var _out15 D0
	_ = _out15
	var _out16 _dafny.Map
	_ = _out16
	var _out17 _dafny.Int
	_ = _out17
	_out15, _out16, _out17 = Companion_Default___.M0(_90_v22, Companion_Default___.Fm2(_77_globalState), _75_v9, _77_globalState)
	_128_v52 = _out15
	_129_v53 = _out16
	_130_v54 = _out17
	var _131_v55 _dafny.Set
	_ = _131_v55
	_131_v55 = _dafny.SetOf(_dafny.IntOfUint32((_66_v0).Cardinality()), _75_v9)
	var _hi3 _dafny.Int = (_131_v55).Cardinality()
	_ = _hi3
	for _132_i7 := ((_71_v5).Difference(_71_v5)).Cardinality(); _132_i7.Cmp(_hi3) < 0; _132_i7 = _132_i7.Plus(_dafny.One) {
		var _rhs24 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm2(_77_globalState)).Cardinality())))
		_ = _rhs24
		var _rhs25 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-737), _132_i7)
		_ = _rhs25
		var _rhs26 _dafny.Array = _74_v8
		_ = _rhs26
		var _lhs25 *GlobalState = _77_globalState
		_ = _lhs25
		_lhs25.F8 = _rhs24
		_130_v54 = _rhs25
		_74_v8 = _rhs26
		var _133_v56 _dafny.Array
		_ = _133_v56
		var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
		_ = _nw18
		_133_v56 = _nw18
		_133_v56 = _133_v56
		var _134_v57 _dafny.Array
		_ = _134_v57
		var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
		_ = _nw19
		_134_v57 = _nw19
		var _135_v58 _dafny.Sequence
		_ = _135_v58
		_135_v58 = _dafny.SeqOf(_74_v8)
		var _136_v59 D1
		_ = _136_v59
		_136_v59 = Companion_D1_.Create_DC3_(_135_v58)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_134_v57), 0))
		_ = _index18
		(_134_v57).ArraySet1((_136_v59).Dtor_cf5(), (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_134_v57), 0))
		_ = _index19
		(_134_v57).ArraySet1(_135_v58, (_index19).Int())
		var _137_v60 D1
		_ = _137_v60
		_137_v60 = Companion_D1_.Create_DC5_(_130_v54, true, (_131_v55).Cardinality(), Companion_Default___.Fm0(_75_v9, _70_v4, _77_globalState))
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_86_v18), 0))
		_ = _index20
		(_86_v18).ArraySet1(((_86_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_86_v18), 0))).Int()).(_dafny.Map)).Update((Companion_Default___.Fm0(Companion_Default___.Fm0(_130_v54, _70_v4, _77_globalState), (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool), _77_globalState)).Plus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_75_v9, _132_i7, (_137_v60).Dtor_cf12(), _132_i7)).Cardinality())))), _75_v9), (_index20).Int())
	}
	var _138_v61 _dafny.MultiSet
	_ = _138_v61
	_138_v61 = _dafny.MultiSetOf(_75_v9, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.IntOfInt64(75)), (Companion_Default___.SafeIndex(_75_v9, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(75))).Cardinality()))).Uint32(), _130_v54)).Cardinality()))
	if (_70_v4) == (((_138_v61).Update((_dafny.SetOf(_67_v1, Companion_D0_.Create_DC0_(_66_v0))).Cardinality(), Companion_Default___.Abs(_dafny.IntOfUint32((_66_v0).Cardinality())))).IsDisjointFrom(_138_v61)) {
		(_77_globalState).F9 = _70_v4
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_74_v8), 0))
		_ = _index21
		(_74_v8).ArraySet1((_dafny.Zero).Minus(_75_v9), (_index21).Int())
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_74_v8), 0))
		_ = _index22
		(_74_v8).ArraySet1(_dafny.IntOfInt64(-38), (_index22).Int())
		var _139_v62 *C0
		_ = _139_v62
		var _nw20 *C0 = New_C0_()
		_ = _nw20
		_nw20.Ctor__()
		_139_v62 = _nw20
		var _140_v63 _dafny.Sequence
		_ = _140_v63
		_140_v63 = _dafny.SeqOf(_71_v5, _dafny.MultiSetOf((_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool)))
		var _141_v64 *C3
		_ = _141_v64
		var _nw21 *C3 = New_C3_()
		_ = _nw21
		_nw21.Ctor__(_140_v63, (_74_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_74_v8), 0))).Int()).(_dafny.Int), !(_70_v4), _dafny.SeqOf(_75_v9))
		_141_v64 = _nw21
		var _142_v65 _dafny.MultiSet
		_ = _142_v65
		_142_v65 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_139_v62, _141_v64))
		var _143_v66 _dafny.Map
		_ = _143_v66
		_143_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_139_v62, _141_v64)
		var _144_v67 *C3
		_ = _144_v67
		var _nw22 *C3 = New_C3_()
		_ = _nw22
		_nw22.Ctor__(Companion_Default___.Fm29(_75_v9, _dafny.SetOf(_130_v54, _75_v9), _77_globalState), (_74_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_74_v8), 0))).Int()).(_dafny.Int), !_dafny.Companion_Sequence_.Contains(_88_v21, (_74_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_74_v8), 0))).Int()).(_dafny.Int)), _dafny.Companion_Sequence_.Update(_88_v21, (Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_142_v65).Contains(_143_v66) {
				return (_142_v65).Multiplicity(_143_v66)
			}
			return _130_v54
		})(), _dafny.IntOfUint32((_88_v21).Cardinality()))).Uint32(), _75_v9))
		_144_v67 = _nw22
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_74_v8), 0))
		_ = _index23
		var _rhs27 _dafny.Array = _90_v22
		_ = _rhs27
		var _rhs28 _dafny.Int = _130_v54
		_ = _rhs28
		var _rhs29 bool = Companion_Default___.Fm1(_75_v9, (func() bool {
			if false {
				return !(true)
			}
			return (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool)
		})(), _77_globalState)
		_ = _rhs29
		var _lhs26 _dafny.Array = _74_v8
		_ = _lhs26
		var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_74_v8), 0))
		_ = _lhs27
		var _lhs28 *GlobalState = _77_globalState
		_ = _lhs28
		_90_v22 = _rhs27
		(_lhs26).ArraySet1(_rhs28, (_lhs27).Int())
		_lhs28.F7 = _rhs29
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))
		_ = _index24
		(_90_v22).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_66_v0, _66_v0), (_index24).Int())
	} else {
		var _145_v68 _dafny.Map
		_ = _145_v68
		_145_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v4, !((_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool)))
		var _146_v69 _dafny.Map
		_ = _146_v69
		_146_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool), (_145_v68).Cardinality())
		_75_v9 = (func() _dafny.Int {
			if (_71_v5).Contains((_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool)) {
				return (_71_v5).Multiplicity((_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool))
			}
			return (func() _dafny.Int {
				if (_146_v69).Contains((_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool)) {
					return (_146_v69).Get((_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool)).(_dafny.Int)
				}
				return _75_v9
			})()
		})()
		var _147_v70 _dafny.Map
		_ = _147_v70
		_147_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(347))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_148_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_149_i8 _dafny.Int) _dafny.CodePoint {
				return _148_v2
			}
		})(_68_v2))), (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool))
		_147_v70 = (_147_v70).Update(_66_v0, _70_v4)
		var _150_v71 _dafny.Sequence
		_ = _150_v71
		_150_v71 = _dafny.SeqOf(_94_v26, _94_v26)
		_94_v26 = _dafny.Companion_Sequence_.Concatenate(_94_v26, _dafny.Companion_Sequence_.Concatenate((_150_v71).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_66_v0).Cardinality())), _dafny.IntOfUint32((_150_v71).Cardinality()))).Uint32()).(_dafny.Sequence), _94_v26))
		var _151_v72 _dafny.Sequence
		_ = _151_v72
		_151_v72 = _dafny.SeqOf(_145_v68)
		_151_v72 = _151_v72
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_74_v8), 0))
		_ = _index25
		(_74_v8).ArraySet1(_130_v54, (_index25).Int())
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_74_v8), 0))
		_ = _index26
		var _rhs30 bool = _70_v4
		_ = _rhs30
		var _rhs31 _dafny.Int = _75_v9
		_ = _rhs31
		var _rhs32 _dafny.Int = _75_v9
		_ = _rhs32
		var _rhs33 _dafny.Int = _75_v9
		_ = _rhs33
		var _lhs29 *GlobalState = _77_globalState
		_ = _lhs29
		var _lhs30 _dafny.Array = _74_v8
		_ = _lhs30
		var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_74_v8), 0))
		_ = _lhs31
		_lhs29.F7 = _rhs30
		(_lhs30).ArraySet1(_rhs31, (_lhs31).Int())
		_75_v9 = _rhs32
		_75_v9 = _rhs33
	}
	var _152_v73 _dafny.Map
	_ = _152_v73
	_152_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_70_v4), (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool))
	_152_v73 = (_152_v73).Update(_70_v4, _70_v4)
	var _153_v74 _dafny.Array
	_ = _153_v74
	var _nw23 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
	_ = _nw23
	_153_v74 = _nw23
	var _154_v75 _dafny.Map
	_ = _154_v75
	_154_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool), _153_v74)
	var _155_v76 T1
	_ = _155_v76
	var _nw24 *C2 = New_C2_()
	_ = _nw24
	_nw24.Ctor__((_86_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_86_v18), 0))).Int()).(_dafny.Map), (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool), _70_v4, Companion_Default___.Fm19(_75_v9, (_90_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_90_v22), 0))).Int()).(bool), _77_globalState), (_154_v75).Cardinality())
	_155_v76 = _nw24
	var _156_v77 _dafny.Array
	_ = _156_v77
	var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(11))
	_ = _nw25
	_156_v77 = _nw25
	var _157_v78 _dafny.Map
	_ = _157_v78
	_157_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v76, _156_v77)
	_157_v78 = (_157_v78).Update(_155_v76, _156_v77)
	(_77_globalState).F7 = (Companion_Default___.SafeModuloInt(_130_v54, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_66_v0, (Companion_Default___.SafeIndex(_75_v9, _dafny.IntOfUint32((_66_v0).Cardinality()))).Uint32(), _68_v2)).Cardinality()))).Cmp(_75_v9) <= 0
	_dafny.Print(_66_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_67_v1).Dtor_cf0().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_68_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v3).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v3).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v3).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v3).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v3).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v3).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v3).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v3).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v3).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v3).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v3).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_70_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v5).Equals(_dafny.MultiSetOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_72_v6).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(true, true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_74_v8).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_74_v8).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_74_v8).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_74_v8).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_75_v9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v10).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_77_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_77_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F2()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F2()).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_77_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_77_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F5()).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(true, true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_77_globalState.F6.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_77_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_77_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_77_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_77_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_77_globalState.F11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_77_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_77_globalState.F14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_77_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_77_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_77_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F18()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_77_globalState).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F20()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_77_globalState).F20()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_77_globalState.F21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_77_globalState).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_78_i0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_86_v18).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(378), _dafny.One).UpdateUnsafe(_dafny.IntOfInt64(688), _dafny.IntOfInt64(-20)).UpdateUnsafe(_dafny.IntOfInt64(-265), _dafny.IntOfInt64(-20)).UpdateUnsafe(_dafny.IntOfInt64(4), _dafny.IntOfInt64(-20))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_87_v19).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(378), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_88_v21, _dafny.SeqOf(_dafny.IntOfInt64(467))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v22).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v23).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_92_v24).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_v25).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("jxuh"), _dafny.UnicodeSeqOfUtf8Bytes("koddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii"), _dafny.UnicodeSeqOfUtf8Bytes("hg"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_94_v26, _dafny.SeqOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_98_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_123_i6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_128_v52).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_128_v52).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_128_v52).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_129_v53).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, false).UpdateUnsafe(_dafny.IntOfInt64(-1), false).UpdateUnsafe(_dafny.One, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_130_v54)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v55).Equals(_dafny.SetOf(_dafny.IntOfInt64(2), _dafny.IntOfInt64(-20))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v61).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-20), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_152_v73).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false).UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_154_v75).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_155_v76).F28())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_155_v76.F29(), _dafny.SeqOf(_dafny.IntOfInt64(3), _dafny.IntOfInt64(-2), _dafny.One, _dafny.One, _dafny.IntOfInt64(7))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_155_v76).F23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_v78).Cardinality())
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
	Cf1 _dafny.Int
	Cf2 bool
	Cf3 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 bool, Cf3 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Sequence
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Sequence) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC2 struct {
	Cf4 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 D0) D0 {
	return D0{D0_DC2{Cf4}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf0() _dafny.Sequence {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf4() D0 {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + data.Cf0.VerbatimString(true) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Equals(data2.Cf0)
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4.Equals(data2.Cf4)
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
	Cf6 _dafny.Map
	Cf7 bool
	Cf8 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf6 _dafny.Map, Cf7 bool, Cf8 bool) D1 {
	return D1{D1_DC4{Cf6, Cf7, Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf9  _dafny.Int
	Cf10 bool
	Cf11 _dafny.Int
	Cf12 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf9 _dafny.Int, Cf10 bool, Cf11 _dafny.Int, Cf12 _dafny.Int) D1 {
	return D1{D1_DC5{Cf9, Cf10, Cf11, Cf12}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
	Cf13 _dafny.Array
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf13 _dafny.Array) D1 {
	return D1{D1_DC6{Cf13}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC3 struct {
	Cf5 _dafny.Sequence
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf5 _dafny.Sequence) D1 {
	return D1{D1_DC3{Cf5}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.EmptyMap, false, false)
}

func (_this D1) Dtor_cf6() _dafny.Map {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf10() bool {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf11
}

func (_this D1) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf12
}

func (_this D1) Dtor_cf13() _dafny.Array {
	return _this.Get_().(D1_DC6).Cf13
}

func (_this D1) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf6.Equals(data2.Cf6) && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10 && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf13 == data2.Cf13
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf5.Equals(data2.Cf5)
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
	Cf15 _dafny.Int
	Cf16 _dafny.Int
	Cf17 _dafny.CodePoint
	Cf18 _dafny.Int
	Cf19 bool
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf15 _dafny.Int, Cf16 _dafny.Int, Cf17 _dafny.CodePoint, Cf18 _dafny.Int, Cf19 bool) D2 {
	return D2{D2_DC8{Cf15, Cf16, Cf17, Cf18, Cf19}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC7 struct {
	Cf14 _dafny.Set
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf14 _dafny.Set) D2 {
	return D2{D2_DC7{Cf14}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(_dafny.Zero, _dafny.Zero, _dafny.CodePoint('D'), _dafny.Zero, false)
}

func (_this D2) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf15
}

func (_this D2) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf16
}

func (_this D2) Dtor_cf17() _dafny.CodePoint {
	return _this.Get_().(D2_DC8).Cf17
}

func (_this D2) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf18
}

func (_this D2) Dtor_cf19() bool {
	return _this.Get_().(D2_DC8).Cf19
}

func (_this D2) Dtor_cf14() _dafny.Set {
	return _this.Get_().(D2_DC7).Cf14
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17 == data2.Cf17 && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19 == data2.Cf19
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf14.Equals(data2.Cf14)
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
	Cf21 bool
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf21 bool) D3 {
	return D3{D3_DC10{Cf21}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC11 struct {
	Cf22 _dafny.Int
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf22 _dafny.Int) D3 {
	return D3{D3_DC11{Cf22}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC12 struct {
	Cf23 _dafny.Sequence
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf23 _dafny.Sequence) D3 {
	return D3{D3_DC12{Cf23}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

type D3_DC9 struct {
	Cf20 _dafny.MultiSet
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf20 _dafny.MultiSet) D3 {
	return D3{D3_DC9{Cf20}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC13 struct {
	Cf24 D3
}

func (D3_DC13) isD3() {}

func (CompanionStruct_D3_) Create_DC13_(Cf24 D3) D3 {
	return D3{D3_DC13{Cf24}}
}

func (_this D3) Is_DC13() bool {
	_, ok := _this.Get_().(D3_DC13)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(false)
}

func (_this D3) Dtor_cf21() bool {
	return _this.Get_().(D3_DC10).Cf21
}

func (_this D3) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf22
}

func (_this D3) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D3_DC12).Cf23
}

func (_this D3) Dtor_cf20() _dafny.MultiSet {
	return _this.Get_().(D3_DC9).Cf20
}

func (_this D3) Dtor_cf24() D3 {
	return _this.Get_().(D3_DC13).Cf24
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf20) + ")"
		}
	case D3_DC13:
		{
			return "D3.DC13" + "(" + _dafny.String(data.Cf24) + ")"
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
			return ok && data1.Cf21 == data2.Cf21
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf22.Cmp(data2.Cf22) == 0
		}
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf23.Equals(data2.Cf23)
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf20.Equals(data2.Cf20)
		}
	case D3_DC13:
		{
			data2, ok := other.Get_().(D3_DC13)
			return ok && data1.Cf24.Equals(data2.Cf24)
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

type D4_DC15 struct {
	Cf26 _dafny.Int
	Cf27 _dafny.CodePoint
	Cf28 _dafny.Int
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf26 _dafny.Int, Cf27 _dafny.CodePoint, Cf28 _dafny.Int) D4 {
	return D4{D4_DC15{Cf26, Cf27, Cf28}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

type D4_DC14 struct {
	Cf25 _dafny.Array
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf25 _dafny.Array) D4 {
	return D4{D4_DC14{Cf25}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC15_(_dafny.Zero, _dafny.CodePoint('D'), _dafny.Zero)
}

func (_this D4) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D4_DC15).Cf26
}

func (_this D4) Dtor_cf27() _dafny.CodePoint {
	return _this.Get_().(D4_DC15).Cf27
}

func (_this D4) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D4_DC15).Cf28
}

func (_this D4) Dtor_cf25() _dafny.Array {
	return _this.Get_().(D4_DC14).Cf25
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf25) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf26.Cmp(data2.Cf26) == 0 && data1.Cf27 == data2.Cf27 && data1.Cf28.Cmp(data2.Cf28) == 0
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf25 == data2.Cf25
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

type D5_DC16 struct {
	Cf29 T1
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf29 T1) D5 {
	return D5{D5_DC16{Cf29}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

func (CompanionStruct_D5_) Default() T1 {
	return (T1)(nil)
}

func (_this D5) Dtor_cf29() T1 {
	return _this.Get_().(D5_DC16).Cf29
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf29) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && _dafny.AreEqual(data1.Cf29, data2.Cf29)
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

type D6_DC17 struct {
	Cf30 _dafny.Sequence
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf30 _dafny.Sequence) D6 {
	return D6{D6_DC17{Cf30}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D6) Dtor_cf30() _dafny.Sequence {
	return _this.Get_().(D6_DC17).Cf30
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf30.Equals(data2.Cf30)
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

type D7_DC18 struct {
	Cf31 _dafny.Map
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf31 _dafny.Map) D7 {
	return D7{D7_DC18{Cf31}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D7) Dtor_cf31() _dafny.Map {
	return _this.Get_().(D7_DC18).Cf31
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf31) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf31.Equals(data2.Cf31)
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

type D8_DC20 struct {
	Cf33 bool
	Cf34 _dafny.Int
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf33 bool, Cf34 _dafny.Int) D8 {
	return D8{D8_DC20{Cf33, Cf34}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

type D8_DC21 struct {
	Cf35 bool
	Cf36 bool
	Cf37 _dafny.Int
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf35 bool, Cf36 bool, Cf37 _dafny.Int) D8 {
	return D8{D8_DC21{Cf35, Cf36, Cf37}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC22 struct {
	Cf38 bool
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf38 bool) D8 {
	return D8{D8_DC22{Cf38}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC19 struct {
	Cf32 _dafny.Sequence
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf32 _dafny.Sequence) D8 {
	return D8{D8_DC19{Cf32}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

type D8_DC23 struct {
	Cf39 D8
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf39 D8) D8 {
	return D8{D8_DC23{Cf39}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC20_(false, _dafny.Zero)
}

func (_this D8) Dtor_cf33() bool {
	return _this.Get_().(D8_DC20).Cf33
}

func (_this D8) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D8_DC20).Cf34
}

func (_this D8) Dtor_cf35() bool {
	return _this.Get_().(D8_DC21).Cf35
}

func (_this D8) Dtor_cf36() bool {
	return _this.Get_().(D8_DC21).Cf36
}

func (_this D8) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf37
}

func (_this D8) Dtor_cf38() bool {
	return _this.Get_().(D8_DC22).Cf38
}

func (_this D8) Dtor_cf32() _dafny.Sequence {
	return _this.Get_().(D8_DC19).Cf32
}

func (_this D8) Dtor_cf39() D8 {
	return _this.Get_().(D8_DC23).Cf39
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf38) + ")"
		}
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf39) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf33 == data2.Cf33 && data1.Cf34.Cmp(data2.Cf34) == 0
		}
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf35 == data2.Cf35 && data1.Cf36 == data2.Cf36 && data1.Cf37.Cmp(data2.Cf37) == 0
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf38 == data2.Cf38
		}
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
			return ok && data1.Cf32.Equals(data2.Cf32)
		}
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf39.Equals(data2.Cf39)
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
	Cf41 bool
	Cf42 _dafny.Set
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf41 bool, Cf42 _dafny.Set) D9 {
	return D9{D9_DC25{Cf41, Cf42}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC26 struct {
	Cf43 _dafny.Int
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf43 _dafny.Int) D9 {
	return D9{D9_DC26{Cf43}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC27 struct {
	Cf44 bool
	Cf45 _dafny.Sequence
	Cf46 _dafny.Array
	Cf47 bool
	Cf48 bool
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf44 bool, Cf45 _dafny.Sequence, Cf46 _dafny.Array, Cf47 bool, Cf48 bool) D9 {
	return D9{D9_DC27{Cf44, Cf45, Cf46, Cf47, Cf48}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

type D9_DC24 struct {
	Cf40 _dafny.Array
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf40 _dafny.Array) D9 {
	return D9{D9_DC24{Cf40}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC28 struct {
	Cf49 D9
}

func (D9_DC28) isD9() {}

func (CompanionStruct_D9_) Create_DC28_(Cf49 D9) D9 {
	return D9{D9_DC28{Cf49}}
}

func (_this D9) Is_DC28() bool {
	_, ok := _this.Get_().(D9_DC28)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC25_(false, _dafny.EmptySet)
}

func (_this D9) Dtor_cf41() bool {
	return _this.Get_().(D9_DC25).Cf41
}

func (_this D9) Dtor_cf42() _dafny.Set {
	return _this.Get_().(D9_DC25).Cf42
}

func (_this D9) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D9_DC26).Cf43
}

func (_this D9) Dtor_cf44() bool {
	return _this.Get_().(D9_DC27).Cf44
}

func (_this D9) Dtor_cf45() _dafny.Sequence {
	return _this.Get_().(D9_DC27).Cf45
}

func (_this D9) Dtor_cf46() _dafny.Array {
	return _this.Get_().(D9_DC27).Cf46
}

func (_this D9) Dtor_cf47() bool {
	return _this.Get_().(D9_DC27).Cf47
}

func (_this D9) Dtor_cf48() bool {
	return _this.Get_().(D9_DC27).Cf48
}

func (_this D9) Dtor_cf40() _dafny.Array {
	return _this.Get_().(D9_DC24).Cf40
}

func (_this D9) Dtor_cf49() D9 {
	return _this.Get_().(D9_DC28).Cf49
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ")"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D9_DC28:
		{
			return "D9.DC28" + "(" + _dafny.String(data.Cf49) + ")"
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
			return ok && data1.Cf41 == data2.Cf41 && data1.Cf42.Equals(data2.Cf42)
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf43.Cmp(data2.Cf43) == 0
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf44 == data2.Cf44 && data1.Cf45.Equals(data2.Cf45) && data1.Cf46 == data2.Cf46 && data1.Cf47 == data2.Cf47 && data1.Cf48 == data2.Cf48
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf40 == data2.Cf40
		}
	case D9_DC28:
		{
			data2, ok := other.Get_().(D9_DC28)
			return ok && data1.Cf49.Equals(data2.Cf49)
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

type D10_DC29 struct {
	Cf50 _dafny.Map
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf50 _dafny.Map) D10 {
	return D10{D10_DC29{Cf50}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D10) Dtor_cf50() _dafny.Map {
	return _this.Get_().(D10_DC29).Cf50
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf50) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC29:
		{
			data2, ok := other.Get_().(D10_DC29)
			return ok && data1.Cf50.Equals(data2.Cf50)
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

type D11_DC31 struct {
	Cf52 _dafny.Int
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf52 _dafny.Int) D11 {
	return D11{D11_DC31{Cf52}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

type D11_DC32 struct {
	Cf53 _dafny.Sequence
}

func (D11_DC32) isD11() {}

func (CompanionStruct_D11_) Create_DC32_(Cf53 _dafny.Sequence) D11 {
	return D11{D11_DC32{Cf53}}
}

func (_this D11) Is_DC32() bool {
	_, ok := _this.Get_().(D11_DC32)
	return ok
}

type D11_DC30 struct {
	Cf51 _dafny.Array
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf51 _dafny.Array) D11 {
	return D11{D11_DC30{Cf51}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC33 struct {
	Cf54 D11
}

func (D11_DC33) isD11() {}

func (CompanionStruct_D11_) Create_DC33_(Cf54 D11) D11 {
	return D11{D11_DC33{Cf54}}
}

func (_this D11) Is_DC33() bool {
	_, ok := _this.Get_().(D11_DC33)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC31_(_dafny.Zero)
}

func (_this D11) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D11_DC31).Cf52
}

func (_this D11) Dtor_cf53() _dafny.Sequence {
	return _this.Get_().(D11_DC32).Cf53
}

func (_this D11) Dtor_cf51() _dafny.Array {
	return _this.Get_().(D11_DC30).Cf51
}

func (_this D11) Dtor_cf54() D11 {
	return _this.Get_().(D11_DC33).Cf54
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf52) + ")"
		}
	case D11_DC32:
		{
			return "D11.DC32" + "(" + data.Cf53.VerbatimString(true) + ")"
		}
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf51) + ")"
		}
	case D11_DC33:
		{
			return "D11.DC33" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf52.Cmp(data2.Cf52) == 0
		}
	case D11_DC32:
		{
			data2, ok := other.Get_().(D11_DC32)
			return ok && data1.Cf53.Equals(data2.Cf53)
		}
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf51 == data2.Cf51
		}
	case D11_DC33:
		{
			data2, ok := other.Get_().(D11_DC33)
			return ok && data1.Cf54.Equals(data2.Cf54)
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

type D12_DC35 struct {
}

func (D12_DC35) isD12() {}

func (CompanionStruct_D12_) Create_DC35_() D12 {
	return D12{D12_DC35{}}
}

func (_this D12) Is_DC35() bool {
	_, ok := _this.Get_().(D12_DC35)
	return ok
}

type D12_DC36 struct {
	Cf56 bool
	Cf57 bool
	Cf58 _dafny.Int
	Cf59 _dafny.Int
	Cf60 _dafny.Map
}

func (D12_DC36) isD12() {}

func (CompanionStruct_D12_) Create_DC36_(Cf56 bool, Cf57 bool, Cf58 _dafny.Int, Cf59 _dafny.Int, Cf60 _dafny.Map) D12 {
	return D12{D12_DC36{Cf56, Cf57, Cf58, Cf59, Cf60}}
}

func (_this D12) Is_DC36() bool {
	_, ok := _this.Get_().(D12_DC36)
	return ok
}

type D12_DC34 struct {
	Cf55 _dafny.Map
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf55 _dafny.Map) D12 {
	return D12{D12_DC34{Cf55}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

type D12_DC37 struct {
	Cf61 D12
}

func (D12_DC37) isD12() {}

func (CompanionStruct_D12_) Create_DC37_(Cf61 D12) D12 {
	return D12{D12_DC37{Cf61}}
}

func (_this D12) Is_DC37() bool {
	_, ok := _this.Get_().(D12_DC37)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC35_()
}

func (_this D12) Dtor_cf56() bool {
	return _this.Get_().(D12_DC36).Cf56
}

func (_this D12) Dtor_cf57() bool {
	return _this.Get_().(D12_DC36).Cf57
}

func (_this D12) Dtor_cf58() _dafny.Int {
	return _this.Get_().(D12_DC36).Cf58
}

func (_this D12) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D12_DC36).Cf59
}

func (_this D12) Dtor_cf60() _dafny.Map {
	return _this.Get_().(D12_DC36).Cf60
}

func (_this D12) Dtor_cf55() _dafny.Map {
	return _this.Get_().(D12_DC34).Cf55
}

func (_this D12) Dtor_cf61() D12 {
	return _this.Get_().(D12_DC37).Cf61
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC35:
		{
			return "D12.DC35"
		}
	case D12_DC36:
		{
			return "D12.DC36" + "(" + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ")"
		}
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf55) + ")"
		}
	case D12_DC37:
		{
			return "D12.DC37" + "(" + _dafny.String(data.Cf61) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC35:
		{
			_, ok := other.Get_().(D12_DC35)
			return ok
		}
	case D12_DC36:
		{
			data2, ok := other.Get_().(D12_DC36)
			return ok && data1.Cf56 == data2.Cf56 && data1.Cf57 == data2.Cf57 && data1.Cf58.Cmp(data2.Cf58) == 0 && data1.Cf59.Cmp(data2.Cf59) == 0 && data1.Cf60.Equals(data2.Cf60)
		}
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
			return ok && data1.Cf55.Equals(data2.Cf55)
		}
	case D12_DC37:
		{
			data2, ok := other.Get_().(D12_DC37)
			return ok && data1.Cf61.Equals(data2.Cf61)
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

type D13_DC39 struct {
	Cf63 _dafny.Int
}

func (D13_DC39) isD13() {}

func (CompanionStruct_D13_) Create_DC39_(Cf63 _dafny.Int) D13 {
	return D13{D13_DC39{Cf63}}
}

func (_this D13) Is_DC39() bool {
	_, ok := _this.Get_().(D13_DC39)
	return ok
}

type D13_DC38 struct {
	Cf62 _dafny.Set
}

func (D13_DC38) isD13() {}

func (CompanionStruct_D13_) Create_DC38_(Cf62 _dafny.Set) D13 {
	return D13{D13_DC38{Cf62}}
}

func (_this D13) Is_DC38() bool {
	_, ok := _this.Get_().(D13_DC38)
	return ok
}

type D13_DC40 struct {
	Cf64 D13
}

func (D13_DC40) isD13() {}

func (CompanionStruct_D13_) Create_DC40_(Cf64 D13) D13 {
	return D13{D13_DC40{Cf64}}
}

func (_this D13) Is_DC40() bool {
	_, ok := _this.Get_().(D13_DC40)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC39_(_dafny.Zero)
}

func (_this D13) Dtor_cf63() _dafny.Int {
	return _this.Get_().(D13_DC39).Cf63
}

func (_this D13) Dtor_cf62() _dafny.Set {
	return _this.Get_().(D13_DC38).Cf62
}

func (_this D13) Dtor_cf64() D13 {
	return _this.Get_().(D13_DC40).Cf64
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC39:
		{
			return "D13.DC39" + "(" + _dafny.String(data.Cf63) + ")"
		}
	case D13_DC38:
		{
			return "D13.DC38" + "(" + _dafny.String(data.Cf62) + ")"
		}
	case D13_DC40:
		{
			return "D13.DC40" + "(" + _dafny.String(data.Cf64) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC39:
		{
			data2, ok := other.Get_().(D13_DC39)
			return ok && data1.Cf63.Cmp(data2.Cf63) == 0
		}
	case D13_DC38:
		{
			data2, ok := other.Get_().(D13_DC38)
			return ok && data1.Cf62.Equals(data2.Cf62)
		}
	case D13_DC40:
		{
			data2, ok := other.Get_().(D13_DC40)
			return ok && data1.Cf64.Equals(data2.Cf64)
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

type D14_DC42 struct {
	Cf66 _dafny.Map
}

func (D14_DC42) isD14() {}

func (CompanionStruct_D14_) Create_DC42_(Cf66 _dafny.Map) D14 {
	return D14{D14_DC42{Cf66}}
}

func (_this D14) Is_DC42() bool {
	_, ok := _this.Get_().(D14_DC42)
	return ok
}

type D14_DC41 struct {
	Cf65 _dafny.Sequence
}

func (D14_DC41) isD14() {}

func (CompanionStruct_D14_) Create_DC41_(Cf65 _dafny.Sequence) D14 {
	return D14{D14_DC41{Cf65}}
}

func (_this D14) Is_DC41() bool {
	_, ok := _this.Get_().(D14_DC41)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC42_(_dafny.EmptyMap)
}

func (_this D14) Dtor_cf66() _dafny.Map {
	return _this.Get_().(D14_DC42).Cf66
}

func (_this D14) Dtor_cf65() _dafny.Sequence {
	return _this.Get_().(D14_DC41).Cf65
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC42:
		{
			return "D14.DC42" + "(" + _dafny.String(data.Cf66) + ")"
		}
	case D14_DC41:
		{
			return "D14.DC41" + "(" + _dafny.String(data.Cf65) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC42:
		{
			data2, ok := other.Get_().(D14_DC42)
			return ok && data1.Cf66.Equals(data2.Cf66)
		}
	case D14_DC41:
		{
			data2, ok := other.Get_().(D14_DC41)
			return ok && data1.Cf65.Equals(data2.Cf65)
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

type D15_DC44 struct {
	Cf68 bool
	Cf69 _dafny.Sequence
	Cf70 _dafny.Int
	Cf71 bool
	Cf72 bool
}

func (D15_DC44) isD15() {}

func (CompanionStruct_D15_) Create_DC44_(Cf68 bool, Cf69 _dafny.Sequence, Cf70 _dafny.Int, Cf71 bool, Cf72 bool) D15 {
	return D15{D15_DC44{Cf68, Cf69, Cf70, Cf71, Cf72}}
}

func (_this D15) Is_DC44() bool {
	_, ok := _this.Get_().(D15_DC44)
	return ok
}

type D15_DC43 struct {
	Cf67 _dafny.Map
}

func (D15_DC43) isD15() {}

func (CompanionStruct_D15_) Create_DC43_(Cf67 _dafny.Map) D15 {
	return D15{D15_DC43{Cf67}}
}

func (_this D15) Is_DC43() bool {
	_, ok := _this.Get_().(D15_DC43)
	return ok
}

type D15_DC45 struct {
	Cf73 D15
}

func (D15_DC45) isD15() {}

func (CompanionStruct_D15_) Create_DC45_(Cf73 D15) D15 {
	return D15{D15_DC45{Cf73}}
}

func (_this D15) Is_DC45() bool {
	_, ok := _this.Get_().(D15_DC45)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC44_(false, _dafny.EmptySeq, _dafny.Zero, false, false)
}

func (_this D15) Dtor_cf68() bool {
	return _this.Get_().(D15_DC44).Cf68
}

func (_this D15) Dtor_cf69() _dafny.Sequence {
	return _this.Get_().(D15_DC44).Cf69
}

func (_this D15) Dtor_cf70() _dafny.Int {
	return _this.Get_().(D15_DC44).Cf70
}

func (_this D15) Dtor_cf71() bool {
	return _this.Get_().(D15_DC44).Cf71
}

func (_this D15) Dtor_cf72() bool {
	return _this.Get_().(D15_DC44).Cf72
}

func (_this D15) Dtor_cf67() _dafny.Map {
	return _this.Get_().(D15_DC43).Cf67
}

func (_this D15) Dtor_cf73() D15 {
	return _this.Get_().(D15_DC45).Cf73
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC44:
		{
			return "D15.DC44" + "(" + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ")"
		}
	case D15_DC43:
		{
			return "D15.DC43" + "(" + _dafny.String(data.Cf67) + ")"
		}
	case D15_DC45:
		{
			return "D15.DC45" + "(" + _dafny.String(data.Cf73) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC44:
		{
			data2, ok := other.Get_().(D15_DC44)
			return ok && data1.Cf68 == data2.Cf68 && data1.Cf69.Equals(data2.Cf69) && data1.Cf70.Cmp(data2.Cf70) == 0 && data1.Cf71 == data2.Cf71 && data1.Cf72 == data2.Cf72
		}
	case D15_DC43:
		{
			data2, ok := other.Get_().(D15_DC43)
			return ok && data1.Cf67.Equals(data2.Cf67)
		}
	case D15_DC45:
		{
			data2, ok := other.Get_().(D15_DC45)
			return ok && data1.Cf73.Equals(data2.Cf73)
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

type D16_DC47 struct {
	Cf75 _dafny.Int
	Cf76 _dafny.Int
	Cf77 *C1
	Cf78 _dafny.Int
}

func (D16_DC47) isD16() {}

func (CompanionStruct_D16_) Create_DC47_(Cf75 _dafny.Int, Cf76 _dafny.Int, Cf77 *C1, Cf78 _dafny.Int) D16 {
	return D16{D16_DC47{Cf75, Cf76, Cf77, Cf78}}
}

func (_this D16) Is_DC47() bool {
	_, ok := _this.Get_().(D16_DC47)
	return ok
}

type D16_DC46 struct {
	Cf74 _dafny.Set
}

func (D16_DC46) isD16() {}

func (CompanionStruct_D16_) Create_DC46_(Cf74 _dafny.Set) D16 {
	return D16{D16_DC46{Cf74}}
}

func (_this D16) Is_DC46() bool {
	_, ok := _this.Get_().(D16_DC46)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC47_(_dafny.Zero, _dafny.Zero, (*C1)(nil), _dafny.Zero)
}

func (_this D16) Dtor_cf75() _dafny.Int {
	return _this.Get_().(D16_DC47).Cf75
}

func (_this D16) Dtor_cf76() _dafny.Int {
	return _this.Get_().(D16_DC47).Cf76
}

func (_this D16) Dtor_cf77() *C1 {
	return _this.Get_().(D16_DC47).Cf77
}

func (_this D16) Dtor_cf78() _dafny.Int {
	return _this.Get_().(D16_DC47).Cf78
}

func (_this D16) Dtor_cf74() _dafny.Set {
	return _this.Get_().(D16_DC46).Cf74
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC47:
		{
			return "D16.DC47" + "(" + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ", " + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ")"
		}
	case D16_DC46:
		{
			return "D16.DC46" + "(" + _dafny.String(data.Cf74) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC47:
		{
			data2, ok := other.Get_().(D16_DC47)
			return ok && data1.Cf75.Cmp(data2.Cf75) == 0 && data1.Cf76.Cmp(data2.Cf76) == 0 && data1.Cf77 == data2.Cf77 && data1.Cf78.Cmp(data2.Cf78) == 0
		}
	case D16_DC46:
		{
			data2, ok := other.Get_().(D16_DC46)
			return ok && data1.Cf74.Equals(data2.Cf74)
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

// Definition of datatype D17
type D17 struct {
	Data_D17_
}

func (_this D17) Get_() Data_D17_ {
	return _this.Data_D17_
}

type Data_D17_ interface {
	isD17()
}

type CompanionStruct_D17_ struct {
}

var Companion_D17_ = CompanionStruct_D17_{}

type D17_DC49 struct {
	Cf80 _dafny.Int
	Cf81 bool
	Cf82 bool
	Cf83 bool
}

func (D17_DC49) isD17() {}

func (CompanionStruct_D17_) Create_DC49_(Cf80 _dafny.Int, Cf81 bool, Cf82 bool, Cf83 bool) D17 {
	return D17{D17_DC49{Cf80, Cf81, Cf82, Cf83}}
}

func (_this D17) Is_DC49() bool {
	_, ok := _this.Get_().(D17_DC49)
	return ok
}

type D17_DC48 struct {
	Cf79 _dafny.MultiSet
}

func (D17_DC48) isD17() {}

func (CompanionStruct_D17_) Create_DC48_(Cf79 _dafny.MultiSet) D17 {
	return D17{D17_DC48{Cf79}}
}

func (_this D17) Is_DC48() bool {
	_, ok := _this.Get_().(D17_DC48)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC49_(_dafny.Zero, false, false, false)
}

func (_this D17) Dtor_cf80() _dafny.Int {
	return _this.Get_().(D17_DC49).Cf80
}

func (_this D17) Dtor_cf81() bool {
	return _this.Get_().(D17_DC49).Cf81
}

func (_this D17) Dtor_cf82() bool {
	return _this.Get_().(D17_DC49).Cf82
}

func (_this D17) Dtor_cf83() bool {
	return _this.Get_().(D17_DC49).Cf83
}

func (_this D17) Dtor_cf79() _dafny.MultiSet {
	return _this.Get_().(D17_DC48).Cf79
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC49:
		{
			return "D17.DC49" + "(" + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ", " + _dafny.String(data.Cf82) + ", " + _dafny.String(data.Cf83) + ")"
		}
	case D17_DC48:
		{
			return "D17.DC48" + "(" + _dafny.String(data.Cf79) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC49:
		{
			data2, ok := other.Get_().(D17_DC49)
			return ok && data1.Cf80.Cmp(data2.Cf80) == 0 && data1.Cf81 == data2.Cf81 && data1.Cf82 == data2.Cf82 && data1.Cf83 == data2.Cf83
		}
	case D17_DC48:
		{
			data2, ok := other.Get_().(D17_DC48)
			return ok && data1.Cf79.Equals(data2.Cf79)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D17) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D17)
	return ok && _this.Equals(typed)
}

func Type_D17_() _dafny.TypeDescriptor {
	return type_D17_{}
}

type type_D17_ struct {
}

func (_this type_D17_) Default() interface{} {
	return Companion_D17_.Default()
}

func (_this type_D17_) String() string {
	return "main.D17"
}
func (_this D17) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D17{}

// End of datatype D17

// Definition of datatype D18
type D18 struct {
	Data_D18_
}

func (_this D18) Get_() Data_D18_ {
	return _this.Data_D18_
}

type Data_D18_ interface {
	isD18()
}

type CompanionStruct_D18_ struct {
}

var Companion_D18_ = CompanionStruct_D18_{}

type D18_DC51 struct {
}

func (D18_DC51) isD18() {}

func (CompanionStruct_D18_) Create_DC51_() D18 {
	return D18{D18_DC51{}}
}

func (_this D18) Is_DC51() bool {
	_, ok := _this.Get_().(D18_DC51)
	return ok
}

type D18_DC50 struct {
	Cf84 _dafny.Map
}

func (D18_DC50) isD18() {}

func (CompanionStruct_D18_) Create_DC50_(Cf84 _dafny.Map) D18 {
	return D18{D18_DC50{Cf84}}
}

func (_this D18) Is_DC50() bool {
	_, ok := _this.Get_().(D18_DC50)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC51_()
}

func (_this D18) Dtor_cf84() _dafny.Map {
	return _this.Get_().(D18_DC50).Cf84
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC51:
		{
			return "D18.DC51"
		}
	case D18_DC50:
		{
			return "D18.DC50" + "(" + _dafny.String(data.Cf84) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC51:
		{
			_, ok := other.Get_().(D18_DC51)
			return ok
		}
	case D18_DC50:
		{
			data2, ok := other.Get_().(D18_DC50)
			return ok && data1.Cf84.Equals(data2.Cf84)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D18) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D18)
	return ok && _this.Equals(typed)
}

func Type_D18_() _dafny.TypeDescriptor {
	return type_D18_{}
}

type type_D18_ struct {
}

func (_this type_D18_) Default() interface{} {
	return Companion_D18_.Default()
}

func (_this type_D18_) String() string {
	return "main.D18"
}
func (_this D18) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D18{}

// End of datatype D18

// Definition of datatype D19
type D19 struct {
	Data_D19_
}

func (_this D19) Get_() Data_D19_ {
	return _this.Data_D19_
}

type Data_D19_ interface {
	isD19()
}

type CompanionStruct_D19_ struct {
}

var Companion_D19_ = CompanionStruct_D19_{}

type D19_DC53 struct {
	Cf86 _dafny.Int
	Cf87 _dafny.Int
}

func (D19_DC53) isD19() {}

func (CompanionStruct_D19_) Create_DC53_(Cf86 _dafny.Int, Cf87 _dafny.Int) D19 {
	return D19{D19_DC53{Cf86, Cf87}}
}

func (_this D19) Is_DC53() bool {
	_, ok := _this.Get_().(D19_DC53)
	return ok
}

type D19_DC52 struct {
	Cf85 _dafny.Map
}

func (D19_DC52) isD19() {}

func (CompanionStruct_D19_) Create_DC52_(Cf85 _dafny.Map) D19 {
	return D19{D19_DC52{Cf85}}
}

func (_this D19) Is_DC52() bool {
	_, ok := _this.Get_().(D19_DC52)
	return ok
}

type D19_DC54 struct {
	Cf88 D19
}

func (D19_DC54) isD19() {}

func (CompanionStruct_D19_) Create_DC54_(Cf88 D19) D19 {
	return D19{D19_DC54{Cf88}}
}

func (_this D19) Is_DC54() bool {
	_, ok := _this.Get_().(D19_DC54)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC53_(_dafny.Zero, _dafny.Zero)
}

func (_this D19) Dtor_cf86() _dafny.Int {
	return _this.Get_().(D19_DC53).Cf86
}

func (_this D19) Dtor_cf87() _dafny.Int {
	return _this.Get_().(D19_DC53).Cf87
}

func (_this D19) Dtor_cf85() _dafny.Map {
	return _this.Get_().(D19_DC52).Cf85
}

func (_this D19) Dtor_cf88() D19 {
	return _this.Get_().(D19_DC54).Cf88
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC53:
		{
			return "D19.DC53" + "(" + _dafny.String(data.Cf86) + ", " + _dafny.String(data.Cf87) + ")"
		}
	case D19_DC52:
		{
			return "D19.DC52" + "(" + _dafny.String(data.Cf85) + ")"
		}
	case D19_DC54:
		{
			return "D19.DC54" + "(" + _dafny.String(data.Cf88) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC53:
		{
			data2, ok := other.Get_().(D19_DC53)
			return ok && data1.Cf86.Cmp(data2.Cf86) == 0 && data1.Cf87.Cmp(data2.Cf87) == 0
		}
	case D19_DC52:
		{
			data2, ok := other.Get_().(D19_DC52)
			return ok && data1.Cf85.Equals(data2.Cf85)
		}
	case D19_DC54:
		{
			data2, ok := other.Get_().(D19_DC54)
			return ok && data1.Cf88.Equals(data2.Cf88)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D19) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D19)
	return ok && _this.Equals(typed)
}

func Type_D19_() _dafny.TypeDescriptor {
	return type_D19_{}
}

type type_D19_ struct {
}

func (_this type_D19_) Default() interface{} {
	return Companion_D19_.Default()
}

func (_this type_D19_) String() string {
	return "main.D19"
}
func (_this D19) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D19{}

// End of datatype D19

// Definition of datatype D20
type D20 struct {
	Data_D20_
}

func (_this D20) Get_() Data_D20_ {
	return _this.Data_D20_
}

type Data_D20_ interface {
	isD20()
}

type CompanionStruct_D20_ struct {
}

var Companion_D20_ = CompanionStruct_D20_{}

type D20_DC56 struct {
	Cf90 bool
	Cf91 _dafny.Int
	Cf92 _dafny.Array
	Cf93 _dafny.CodePoint
	Cf94 _dafny.Int
}

func (D20_DC56) isD20() {}

func (CompanionStruct_D20_) Create_DC56_(Cf90 bool, Cf91 _dafny.Int, Cf92 _dafny.Array, Cf93 _dafny.CodePoint, Cf94 _dafny.Int) D20 {
	return D20{D20_DC56{Cf90, Cf91, Cf92, Cf93, Cf94}}
}

func (_this D20) Is_DC56() bool {
	_, ok := _this.Get_().(D20_DC56)
	return ok
}

type D20_DC57 struct {
}

func (D20_DC57) isD20() {}

func (CompanionStruct_D20_) Create_DC57_() D20 {
	return D20{D20_DC57{}}
}

func (_this D20) Is_DC57() bool {
	_, ok := _this.Get_().(D20_DC57)
	return ok
}

type D20_DC55 struct {
	Cf89 _dafny.MultiSet
}

func (D20_DC55) isD20() {}

func (CompanionStruct_D20_) Create_DC55_(Cf89 _dafny.MultiSet) D20 {
	return D20{D20_DC55{Cf89}}
}

func (_this D20) Is_DC55() bool {
	_, ok := _this.Get_().(D20_DC55)
	return ok
}

type D20_DC58 struct {
	Cf95 D20
}

func (D20_DC58) isD20() {}

func (CompanionStruct_D20_) Create_DC58_(Cf95 D20) D20 {
	return D20{D20_DC58{Cf95}}
}

func (_this D20) Is_DC58() bool {
	_, ok := _this.Get_().(D20_DC58)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC56_(false, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.CodePoint('D'), _dafny.Zero)
}

func (_this D20) Dtor_cf90() bool {
	return _this.Get_().(D20_DC56).Cf90
}

func (_this D20) Dtor_cf91() _dafny.Int {
	return _this.Get_().(D20_DC56).Cf91
}

func (_this D20) Dtor_cf92() _dafny.Array {
	return _this.Get_().(D20_DC56).Cf92
}

func (_this D20) Dtor_cf93() _dafny.CodePoint {
	return _this.Get_().(D20_DC56).Cf93
}

func (_this D20) Dtor_cf94() _dafny.Int {
	return _this.Get_().(D20_DC56).Cf94
}

func (_this D20) Dtor_cf89() _dafny.MultiSet {
	return _this.Get_().(D20_DC55).Cf89
}

func (_this D20) Dtor_cf95() D20 {
	return _this.Get_().(D20_DC58).Cf95
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC56:
		{
			return "D20.DC56" + "(" + _dafny.String(data.Cf90) + ", " + _dafny.String(data.Cf91) + ", " + _dafny.String(data.Cf92) + ", " + _dafny.String(data.Cf93) + ", " + _dafny.String(data.Cf94) + ")"
		}
	case D20_DC57:
		{
			return "D20.DC57"
		}
	case D20_DC55:
		{
			return "D20.DC55" + "(" + _dafny.String(data.Cf89) + ")"
		}
	case D20_DC58:
		{
			return "D20.DC58" + "(" + _dafny.String(data.Cf95) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC56:
		{
			data2, ok := other.Get_().(D20_DC56)
			return ok && data1.Cf90 == data2.Cf90 && data1.Cf91.Cmp(data2.Cf91) == 0 && data1.Cf92 == data2.Cf92 && data1.Cf93 == data2.Cf93 && data1.Cf94.Cmp(data2.Cf94) == 0
		}
	case D20_DC57:
		{
			_, ok := other.Get_().(D20_DC57)
			return ok
		}
	case D20_DC55:
		{
			data2, ok := other.Get_().(D20_DC55)
			return ok && data1.Cf89.Equals(data2.Cf89)
		}
	case D20_DC58:
		{
			data2, ok := other.Get_().(D20_DC58)
			return ok && data1.Cf95.Equals(data2.Cf95)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D20) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D20)
	return ok && _this.Equals(typed)
}

func Type_D20_() _dafny.TypeDescriptor {
	return type_D20_{}
}

type type_D20_ struct {
}

func (_this type_D20_) Default() interface{} {
	return Companion_D20_.Default()
}

func (_this type_D20_) String() string {
	return "main.D20"
}
func (_this D20) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D20{}

// End of datatype D20

// Definition of trait T0
type T0 interface {
	String() string
	M1(p0 _dafny.Array, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Set)
	M2(p0 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.Sequence)
	F23() _dafny.Int
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
	M1(p0 _dafny.Array, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Set)
	M2(p0 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.Sequence)
	F23() _dafny.Int
	F29() _dafny.Sequence
	F29_set_(value _dafny.Sequence)
	F28() bool
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
	F6   _dafny.Sequence
	F7   bool
	F8   _dafny.Int
	F9   bool
	F11  _dafny.CodePoint
	F14  _dafny.Int
	F21  bool
	_f0  bool
	_f1  _dafny.Int
	_f2  _dafny.Array
	_f3  _dafny.Int
	_f4  bool
	_f5  _dafny.MultiSet
	_f10 _dafny.Int
	_f12 _dafny.Int
	_f13 _dafny.Array
	_f15 _dafny.Int
	_f16 bool
	_f17 _dafny.Int
	_f18 _dafny.Map
	_f19 bool
	_f20 _dafny.Array
	_f22 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F6 = _dafny.EmptySeq
	_this.F7 = false
	_this.F8 = _dafny.Zero
	_this.F9 = false
	_this.F11 = _dafny.CodePoint('D')
	_this.F14 = _dafny.Zero
	_this.F21 = false
	_this._f0 = false
	_this._f1 = _dafny.Zero
	_this._f2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f3 = _dafny.Zero
	_this._f4 = false
	_this._f5 = _dafny.EmptyMultiSet
	_this._f10 = _dafny.Zero
	_this._f12 = _dafny.Zero
	_this._f13 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f15 = _dafny.Zero
	_this._f16 = false
	_this._f17 = _dafny.Zero
	_this._f18 = _dafny.EmptyMap
	_this._f19 = false
	_this._f20 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f22 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 _dafny.Array, f3 _dafny.Int, f4 bool, f5 _dafny.MultiSet, f6 _dafny.Sequence, f7 bool, f8 _dafny.Int, f9 bool, f10 _dafny.Int, f11 _dafny.CodePoint, f12 _dafny.Int, f13 _dafny.Array, f14 _dafny.Int, f15 _dafny.Int, f16 bool, f17 _dafny.Int, f18 _dafny.Map, f19 bool, f20 _dafny.Array, f21 bool, f22 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this).F7 = f7
		(_this).F8 = f8
		(_this).F9 = f9
		(_this)._f10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this).F21 = f21
		(_this)._f22 = f22
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Array {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.MultiSet {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Array {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() bool {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Int {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Map {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F19() bool {
	{
		return _this._f19
	}
}
func (_this *GlobalState) F20() _dafny.Array {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F22() _dafny.Int {
	{
		return _this._f22
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	dummy byte
}

func New_C0_() *C0 {
	_this := C0{}

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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) Ctor__() {
	{
	}
}
func (_this *C0) Fm20(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return false
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f29 _dafny.Sequence
	_f23 _dafny.Int
	_f28 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f29 = _dafny.EmptySeq
	_this._f23 = _dafny.Zero
	_this._f28 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C1{}
var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F29() _dafny.Sequence {
	return _this._f29
}
func (_this *C1) F29_set_(value _dafny.Sequence) {
	_this._f29 = value
}
func (_this *C1) F23() _dafny.Int {
	return _this._f23
}
func (_this *C1) F28() bool {
	return _this._f28
}
func (_this *C1) Ctor__(f28 bool, f29 _dafny.Sequence, f23 _dafny.Int) {
	{
		(_this)._f28 = f28
		(_this)._f29 = f29
		(_this)._f23 = f23
	}
}
func (_this *C1) M1(p0 _dafny.Array, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Set) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Set = _dafny.EmptySet
		_ = r1
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
		_ = _index27
		(p0).ArraySet1(p1, (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
		_ = _index28
		(p0).ArraySet1((_this).F28(), (_index28).Int())
		var _158_v0 _dafny.Sequence
		_ = _158_v0
		_158_v0 = _dafny.UnicodeSeqOfUtf8Bytes("pilipchus")
		var _159_v1 _dafny.Set
		_ = _159_v1
		_159_v1 = _dafny.SetOf(_158_v0)
		(_this).F29_set_(Companion_Default___.Fm19(_dafny.IntOfUint32((_158_v0).Cardinality()), !(_159_v1).Contains(_158_v0), globalState))
		var _160_v2 _dafny.CodePoint
		_ = _160_v2
		_160_v2 = _dafny.CodePoint('n')
		var _161_v3 _dafny.MultiSet
		_ = _161_v3
		_161_v3 = _dafny.MultiSetOf(_160_v2)
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
		_ = _index29
		var _rhs34 _dafny.Int = Companion_Default___.Fm0((_this).F23(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState)
		_ = _rhs34
		var _rhs35 bool = (_this).F28()
		_ = _rhs35
		var _rhs36 _dafny.Int = (_this).F23()
		_ = _rhs36
		var _rhs37 _dafny.MultiSet = _161_v3
		_ = _rhs37
		var _rhs38 bool = p1
		_ = _rhs38
		var _lhs32 *GlobalState = globalState
		_ = _lhs32
		var _lhs33 *GlobalState = globalState
		_ = _lhs33
		var _lhs34 _dafny.Array = p0
		_ = _lhs34
		var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
		_ = _lhs35
		_lhs32.F14 = _rhs34
		_lhs33.F7 = _rhs35
		r0 = _rhs36
		_161_v3 = _rhs37
		(_lhs34).ArraySet1(_rhs38, (_lhs35).Int())
		var _162_v5 _dafny.Map
		_ = _162_v5
		_162_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F23())
		var _hi4 _dafny.Int = (_this).F23()
		_ = _hi4
		for _163_i0 := _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter11 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(965))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}(func(_191_i1 _dafny.Int) _dafny.Int {
				return (_this).F23()
			}))).Elements()); ; {
				_compr_10, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _192_v4 _dafny.Int
				_192_v4 = interface{}(_compr_10).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(965))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg20 _dafny.Int) interface{} {
						return coer20(arg20)
					}
				}(func(_191_i1 _dafny.Int) _dafny.Int {
					return (_this).F23()
				})), _192_v4) {
					_coll10.Add((_192_v4).Minus(_dafny.IntOfInt64(783)))
				}
			}
			return _coll10.ToSet()
		}()).Cardinality(), (_this).F23(), (func() _dafny.Int {
			if (_162_v5).Contains(p0) {
				return (_162_v5).Get(p0).(_dafny.Int)
			}
			return (_this).F23()
		})())).Cardinality()); _163_i0.Cmp(_hi4) < 0; _163_i0 = _163_i0.Plus(_dafny.One) {
			if p1 {
				(globalState).F14 = (func() _dafny.Int {
					if _dafny.Companion_Sequence_.IsPrefixOf(_158_v0, _158_v0) {
						return (_163_i0).Times((_this).F23())
					}
					return _163_i0
				})()
				var _164_v6 _dafny.Sequence
				_ = _164_v6
				_164_v6 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p1, (_this).F28(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool), Companion_Default___.Fm1((_this).F23(), (_this).F28(), globalState), p1), (Companion_Default___.SafeIndex(_163_i0, _dafny.IntOfUint32((_dafny.SeqOf(p1, (_this).F28(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool), Companion_Default___.Fm1((_this).F23(), (_this).F28(), globalState), p1)).Cardinality()))).Uint32(), (_this).F28()))
				var _165_v7 _dafny.Map
				_ = _165_v7
				_165_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_i0, (_this).F23())
				var _166_v8 D1
				_ = _166_v8
				_166_v8 = Companion_D1_.Create_DC4_(_165_v7, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool), (_this).F28())
				var _167_v9 _dafny.Map
				_ = _167_v9
				_167_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_164_v6, _dafny.Companion_Sequence_.Update(_164_v6, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_164_v6).Cardinality()))).Uint32(), _dafny.SeqOf(p1, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool)))), (_166_v8).Dtor_cf8())
				_167_v9 = (_167_v9).Update(_164_v6, (_this).F28())
				var _168_v10 _dafny.Map
				_ = _168_v10
				_168_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool), _dafny.IntOfInt64(-375))
				var _169_v11 _dafny.Sequence
				_ = _169_v11
				_169_v11 = _dafny.SeqOf((_this).F28())
				var _rhs39 _dafny.Int = ((_this).F23()).Plus((func() _dafny.Int {
					if (_168_v10).Contains(Companion_Default___.Fm1(_dafny.IntOfUint32((_169_v11).Cardinality()), (_this).F28(), globalState)) {
						return (_168_v10).Get(Companion_Default___.Fm1(_dafny.IntOfUint32((_169_v11).Cardinality()), (_this).F28(), globalState)).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-960))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg17 _dafny.Int) interface{} {
							return coer17(arg17)
						}
					}(func(_170_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('v')
					}))).Cardinality())
				})())
				_ = _rhs39
				var _rhs40 bool = ((_this).F23()).Cmp((_this).F23()) > 0
				_ = _rhs40
				var _lhs36 *GlobalState = globalState
				_ = _lhs36
				var _lhs37 *GlobalState = globalState
				_ = _lhs37
				_lhs36.F14 = _rhs39
				_lhs37.F9 = _rhs40
				var _171_v12 _dafny.Array
				_ = _171_v12
				var _len0_0 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_0
				var _nw26 _dafny.Array
				_ = _nw26
				if _len0_0.Cmp(_dafny.Zero) == 0 {
					_nw26 = _dafny.NewArray(_len0_0)
				} else {
					var _init0 func(_dafny.Int) _dafny.Int = func(_172_i3 _dafny.Int) _dafny.Int {
						return (_172_i3).Times(_dafny.IntOfInt64(635))
					}
					_ = _init0
					var _element0_0 = _init0(_dafny.Zero)
					_ = _element0_0
					_nw26 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
					(_nw26).ArraySet1(_element0_0, 0)
					var _nativeLen0_0 = (_len0_0).Int()
					_ = _nativeLen0_0
					for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
						(_nw26).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
					}
				}
				_171_v12 = _nw26
				var _173_v13 _dafny.Sequence
				_ = _173_v13
				_173_v13 = _dafny.SeqOf(_171_v12)
				var _174_v14 D1
				_ = _174_v14
				_174_v14 = Companion_D1_.Create_DC3_(_dafny.SeqOf(_171_v12, _171_v12, _171_v12))
				var _175_v15 _dafny.Array
				_ = _175_v15
				var _nwElement0_4 D1 = Companion_D1_.Create_DC3_(_173_v13)
				_ = _nwElement0_4
				var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(10))
				_ = _nw27
				(_nw27).ArraySet1(_nwElement0_4, 0)
				(_nw27).ArraySet1(_174_v14, 1)
				(_nw27).ArraySet1(_174_v14, 2)
				(_nw27).ArraySet1(_174_v14, 3)
				(_nw27).ArraySet1(Companion_D1_.Create_DC3_(_173_v13), 4)
				(_nw27).ArraySet1(_174_v14, 5)
				(_nw27).ArraySet1(_174_v14, 6)
				(_nw27).ArraySet1(_174_v14, 7)
				(_nw27).ArraySet1(_174_v14, 8)
				(_nw27).ArraySet1(_174_v14, 9)
				_175_v15 = _nw27
				var _176_v16 _dafny.MultiSet
				_ = _176_v16
				_176_v16 = _dafny.MultiSetOf(_175_v15, _175_v15)
				var _177_v17 D3
				_ = _177_v17
				_177_v17 = Companion_D3_.Create_DC9_(_176_v16)
				var _178_v18 _dafny.Sequence
				_ = _178_v18
				_178_v18 = _dafny.SeqOf(_177_v17, _177_v17, _177_v17, _177_v17)
				_178_v18 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_this).F28() {
						return _178_v18
					}
					return _178_v18
				})(), _178_v18)
				(globalState).F21 = (_this).F28()
			} else {
				var _179_v19 D0
				_ = _179_v19
				var _180_v20 _dafny.Map
				_ = _180_v20
				var _181_v21 _dafny.Int
				_ = _181_v21
				var _out18 D0
				_ = _out18
				var _out19 _dafny.Map
				_ = _out19
				var _out20 _dafny.Int
				_ = _out20
				_out18, _out19, _out20 = Companion_Default___.M0(p0, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm2(globalState), _158_v0), (_this.F29()).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32()).(_dafny.Int), globalState)
				_179_v19 = _out18
				_180_v20 = _out19
				_181_v21 = _out20
				var _182_v22 _dafny.Array
				_ = _182_v22
				var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(25))
				_ = _nw28
				_182_v22 = _nw28
				var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(24))
				_ = _nw29
				_182_v22 = _nw29
				var _183_v23 _dafny.Map
				_ = _183_v23
				_183_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_158_v0).Cardinality()), _181_v21)
				var _184_v24 _dafny.Set
				_ = _184_v24
				_184_v24 = _dafny.SetOf((_183_v23).Cardinality(), (_this).F23())
				var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
				_ = _index30
				var _rhs41 bool = p1
				_ = _rhs41
				var _rhs42 bool = !((func() bool {
					if (_184_v24).IsProperSubsetOf(_184_v24) {
						return true
					}
					return !_dafny.Companion_Sequence_.Equal(_158_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(196))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg18 _dafny.Int) interface{} {
							return coer18(arg18)
						}
					}((func(_185_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_186_i4 _dafny.Int) _dafny.CodePoint {
							return _185_v2
						}
					})(_160_v2))))
				})())
				_ = _rhs42
				var _rhs43 _dafny.Int = _dafny.IntOfInt64(595)
				_ = _rhs43
				var _lhs38 _dafny.Array = p0
				_ = _lhs38
				var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
				_ = _lhs39
				var _lhs40 *GlobalState = globalState
				_ = _lhs40
				var _lhs41 *GlobalState = globalState
				_ = _lhs41
				(_lhs38).ArraySet1(_rhs41, (_lhs39).Int())
				_lhs40.F7 = _rhs42
				_lhs41.F8 = _rhs43
				var _187_v25 _dafny.Map
				_ = _187_v25
				_187_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_this.F29()).Cardinality()), _160_v2)
				(globalState).F11 = (func() _dafny.CodePoint {
					if (_187_v25).Contains((_this).F23()) {
						return (_187_v25).Get((_this).F23()).(_dafny.CodePoint)
					}
					return Companion_Default___.Fm6(_163_i0, globalState)
				})()
				var _188_v26 _dafny.Array
				_ = _188_v26
				var _nw30 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
				_ = _nw30
				_188_v26 = _nw30
				_188_v26 = _188_v26
			}
			var _189_v27 _dafny.Sequence
			_ = _189_v27
			_189_v27 = _dafny.SeqOf((_this).F28())
			var _190_v28 _dafny.Sequence
			_ = _190_v28
			_190_v28 = _dafny.SeqOf(_189_v27, _189_v27)
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
			_ = _index31
			var _rhs44 bool = p1
			_ = _rhs44
			var _rhs45 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_190_v28, _dafny.Companion_Sequence_.Update(_190_v28, (Companion_Default___.SafeIndex(_163_i0, _dafny.IntOfUint32((_190_v28).Cardinality()))).Uint32(), _189_v27))
			_ = _rhs45
			var _rhs46 bool = (_189_v27).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_this).F23(), _163_i0), _dafny.IntOfUint32((_189_v27).Cardinality()))).Uint32()).(bool)
			_ = _rhs46
			var _lhs42 _dafny.Array = p0
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
			_ = _lhs43
			var _lhs44 *GlobalState = globalState
			_ = _lhs44
			(_lhs42).ArraySet1(_rhs44, (_lhs43).Int())
			_190_v28 = _rhs45
			_lhs44.F21 = _rhs46
			(globalState).F8 = (func() _dafny.Int {
				if (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
					return (_dafny.IntOfInt64(950)).Plus((_this).F23())
				}
				return _163_i0
			})()
			(globalState).F21 = (_this).F28()
		}
		var _193_v29 _dafny.Sequence
		_ = _193_v29
		_193_v29 = _dafny.SeqOf(false)
		var _194_v30 _dafny.Map
		_ = _194_v30
		_194_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F23())
		var _195_v31 _dafny.Sequence
		_ = _195_v31
		_195_v31 = _dafny.SeqOf(true, !((_dafny.IntOfUint32((_193_v29).Cardinality())).Cmp((_this).F23()) >= 0), p1, (_this).F28(), (_194_v30).Contains(false))
		if (_195_v31).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_195_v31).Cardinality()))).Uint32()).(bool) {
			var _196_v32 _dafny.Array
			_ = _196_v32
			var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw31
			_196_v32 = _nw31
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))
			_ = _index32
			(_196_v32).ArraySet1(_dafny.IntOfUint32((_this.F29()).Cardinality()), (_index32).Int())
			var _197_v33 _dafny.Map
			_ = _197_v33
			_197_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _196_v32)
			var _198_v34 _dafny.Map
			_ = _198_v34
			_198_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_197_v33, true)
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))
			_ = _index33
			var _rhs47 _dafny.Int = (_dafny.Zero).Minus((_this).F23())
			_ = _rhs47
			var _rhs48 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_this.F29(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(288))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}(func(_199_i5 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(715)
			})))
			_ = _rhs48
			var _rhs49 _dafny.Int = (_dafny.Zero).Minus((_this).F23())
			_ = _rhs49
			var _rhs50 bool = (func() bool {
				if (_198_v34).Contains(_197_v33) {
					return (_198_v34).Get(_197_v33).(bool)
				}
				return Companion_Default___.Fm1((_this).F23(), p1, globalState)
			})()
			_ = _rhs50
			var _rhs51 _dafny.Int = (func() _dafny.Int {
				if !(!(!((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool)))) || ((_this).F28()) {
					return _dafny.IntOfUint32((_195_v31).Cardinality())
				}
				return (_this).F23()
			})()
			_ = _rhs51
			var _lhs45 *GlobalState = globalState
			_ = _lhs45
			var _lhs46 *C1 = _this
			_ = _lhs46
			var _lhs47 _dafny.Array = _196_v32
			_ = _lhs47
			var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))
			_ = _lhs48
			var _lhs49 *GlobalState = globalState
			_ = _lhs49
			_lhs45.F14 = _rhs47
			_lhs46.F29_set_(_rhs48)
			(_lhs47).ArraySet1(_rhs49, (_lhs48).Int())
			_lhs49.F21 = _rhs50
			r0 = _rhs51
			(globalState).F14 = (_196_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))).Int()).(_dafny.Int)
			var _200_v35 _dafny.Sequence
			_ = _200_v35
			_200_v35 = _dafny.SeqOf(_158_v0, Companion_Default___.Fm2(globalState), _158_v0, _158_v0, _158_v0)
			var _201_v37 _dafny.Map
			_ = _201_v37
			_201_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(831), (_this).F28())
			_159_v1 = func() _dafny.Set {
				var _coll11 = _dafny.NewBuilder()
				_ = _coll11
				for _iter12 := _dafny.Iterate(((func() _dafny.Set {
					var _coll12 = _dafny.NewBuilder()
					_ = _coll12
					for _iter13 := _dafny.Iterate((_200_v35).Elements()); ; {
						_compr_12, _ok13 := _iter13()
						if !_ok13 {
							break
						}
						var _202_v36 _dafny.Sequence
						_202_v36 = interface{}(_compr_12).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_200_v35, _202_v36) {
							_coll12.Add(_202_v36)
						}
					}
					return _coll12.ToSet()
				}()).Difference(func() _dafny.Set {
					var _coll13 = _dafny.NewBuilder()
					_ = _coll13
					for _iter14 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v0, (func() bool {
						if (_201_v37).Contains((_194_v30).Cardinality()) {
							return (_201_v37).Get((_194_v30).Cardinality()).(bool)
						}
						return p1
					})())).Keys().Elements()); ; {
						_compr_13, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _203_v38 _dafny.Sequence
						_203_v38 = interface{}(_compr_13).(_dafny.Sequence)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v0, (func() bool {
							if (_201_v37).Contains((_194_v30).Cardinality()) {
								return (_201_v37).Get((_194_v30).Cardinality()).(bool)
							}
							return p1
						})())).Contains(_203_v38) {
							_coll13.Add(_203_v38)
						}
					}
					return _coll13.ToSet()
				}())).Elements()); ; {
					_compr_11, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _204_v39 _dafny.Sequence
					_204_v39 = interface{}(_compr_11).(_dafny.Sequence)
					if ((func() _dafny.Set {
						var _coll14 = _dafny.NewBuilder()
						_ = _coll14
						for _iter15 := _dafny.Iterate((_200_v35).Elements()); ; {
							_compr_14, _ok15 := _iter15()
							if !_ok15 {
								break
							}
							var _205_v36 _dafny.Sequence
							_205_v36 = interface{}(_compr_14).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(_200_v35, _205_v36) {
								_coll14.Add(_205_v36)
							}
						}
						return _coll14.ToSet()
					}()).Difference(func() _dafny.Set {
						var _coll15 = _dafny.NewBuilder()
						_ = _coll15
						for _iter16 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v0, (func() bool {
							if (_201_v37).Contains((_194_v30).Cardinality()) {
								return (_201_v37).Get((_194_v30).Cardinality()).(bool)
							}
							return p1
						})())).Keys().Elements()); ; {
							_compr_15, _ok16 := _iter16()
							if !_ok16 {
								break
							}
							var _206_v38 _dafny.Sequence
							_206_v38 = interface{}(_compr_15).(_dafny.Sequence)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v0, (func() bool {
								if (_201_v37).Contains((_194_v30).Cardinality()) {
									return (_201_v37).Get((_194_v30).Cardinality()).(bool)
								}
								return p1
							})())).Contains(_206_v38) {
								_coll15.Add(_206_v38)
							}
						}
						return _coll15.ToSet()
					}())).Contains(_204_v39) {
						_coll11.Add(_204_v39)
					}
				}
				return _coll11.ToSet()
			}()
			var _207_v41 _dafny.Set
			_ = _207_v41
			_207_v41 = _dafny.SetOf((func() _dafny.Set {
				var _coll16 = _dafny.NewBuilder()
				_ = _coll16
				for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(784), _dafny.IntOfInt64(480))); ; {
					_compr_16, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _208_v40 _dafny.Int
					_208_v40 = interface{}(_compr_16).(_dafny.Int)
					if ((_dafny.IntOfInt64(784)).Cmp(_208_v40) <= 0) && ((_208_v40).Cmp(_dafny.IntOfInt64(480)) < 0) {
						_coll16.Add(Companion_Default___.SafeDivisionInt(_208_v40, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _dafny.MultiSetFromSeq(_this.F29()))).Cardinality()))
					}
				}
				return _coll16.ToSet()
			}()).Cardinality(), (_196_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))).Int()).(_dafny.Int))
			r0 = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0((Companion_Default___.Fm5(_207_v41, globalState)).Cardinality(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState), (func() _dafny.Int {
				if Companion_Default___.Fm1((_196_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))).Int()).(_dafny.Int), p1, globalState) {
					return (_196_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))).Int()).(_dafny.Int)
				}
				return (_this).F23()
			})())
			if (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
				var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
				_ = _index34
				(p0).ArraySet1((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool), (_index34).Int())
				var _209_v42 *C0
				_ = _209_v42
				var _nw32 *C0 = New_C0_()
				_ = _nw32
				_nw32.Ctor__()
				_209_v42 = _nw32
				var _210_v43 _dafny.Array
				_ = _210_v43
				var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
				_ = _nw33
				_210_v43 = _nw33
				var _211_v44 _dafny.Sequence
				_ = _211_v44
				_211_v44 = _dafny.SeqOf(_193_v29, _195_v31, _dafny.SeqOf(p1))
				var _rhs52 bool = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate((_211_v44).Select((Companion_Default___.SafeIndex((_196_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_211_v44).Cardinality()))).Uint32()).(_dafny.Sequence), _193_v29), true)
				_ = _rhs52
				var _rhs53 _dafny.Array = _210_v43
				_ = _rhs53
				var _lhs50 *GlobalState = globalState
				_ = _lhs50
				_lhs50.F9 = _rhs52
				_210_v43 = _rhs53
				(globalState).F9 = p1
				var _212_v45 _dafny.Set
				_ = _212_v45
				_212_v45 = _dafny.SetOf(_160_v2, _160_v2)
				var _213_v46 D2
				_ = _213_v46
				_213_v46 = Companion_D2_.Create_DC7_(_212_v45)
				var _214_v47 _dafny.Map
				_ = _214_v47
				_214_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_213_v46, _158_v0)
				var _215_v48 _dafny.MultiSet
				_ = _215_v48
				_215_v48 = _dafny.MultiSetOf((_209_v42).Fm20(_214_v47, (_196_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_158_v0).Cardinality()), (_196_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))).Int()).(_dafny.Int), globalState))
				var _216_v49 _dafny.Map
				_ = _216_v49
				_216_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v48, _dafny.UnicodeSeqOfUtf8Bytes("dux"))
				var _217_v50 D0
				_ = _217_v50
				_217_v50 = Companion_D0_.Create_DC0_(_158_v0)
				var _218_v51 _dafny.Sequence
				_ = _218_v51
				_218_v51 = _dafny.SeqOf(Companion_Default___.Fm4(_217_v50, (_this).F23(), _dafny.SeqOf((_this).F28()), _160_v2, globalState), _215_v48)
				_216_v49 = (_216_v49).Update((_218_v51).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_218_v51).Cardinality()))).Uint32()).(_dafny.MultiSet), _158_v0)
			} else {
				var _219_v52 _dafny.Map
				_ = _219_v52
				_219_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _dafny.IntOfInt64(275))
				_219_v52 = (_219_v52).Update((_196_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))).Int()).(_dafny.Int), (_196_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))).Int()).(_dafny.Int))
				_207_v41 = Companion_Default___.Fm21(_200_v35, globalState)
				var _220_v53 D0
				_ = _220_v53
				var _221_v54 _dafny.Map
				_ = _221_v54
				var _222_v55 _dafny.Int
				_ = _222_v55
				var _out21 D0
				_ = _out21
				var _out22 _dafny.Map
				_ = _out22
				var _out23 _dafny.Int
				_ = _out23
				_out21, _out22, _out23 = Companion_Default___.M0(p0, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_158_v0, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_158_v0).Cardinality()))).Uint32(), _160_v2), _158_v0), (_this).F23(), globalState)
				_220_v53 = _out21
				_221_v54 = _out22
				_222_v55 = _out23
				var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))
				_ = _index35
				(_196_v32).ArraySet1((_this).F23(), (_index35).Int())
				(globalState).F9 = !((((_196_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))).Int()).(_dafny.Int)).Cmp(_222_v55) != 0) == (((_220_v53).Dtor_cf3()).Cmp((_196_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_196_v32), 0))).Int()).(_dafny.Int)) == 0))
			}
		} else {
			var _223_v56 _dafny.Array
			_ = _223_v56
			var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
			_ = _nw34
			_223_v56 = _nw34
			var _224_v57 _dafny.Array
			_ = _224_v57
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_1
			var _nw35 _dafny.Array
			_ = _nw35
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw35 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Sequence = func(_225_i6 _dafny.Int) _dafny.Sequence {
					return _this.F29()
				}
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw35 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw35).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw35).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_224_v57 = _nw35
			var _226_v58 _dafny.Sequence
			_ = _226_v58
			_226_v58 = _dafny.SeqOf(_224_v57, _224_v57)
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_223_v56), 0))
			_ = _index36
			(_223_v56).ArraySet1((_226_v58).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_226_v58).Cardinality()))).Uint32()).(_dafny.Array), (_index36).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_223_v56), 0))
			_ = _index37
			(_223_v56).ArraySet1(_224_v57, (_index37).Int())
			var _227_v59 _dafny.Array
			_ = _227_v59
			var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw36
			_227_v59 = _nw36
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_227_v59), 0))
			_ = _index38
			(_227_v59).ArraySet1((_this).F23(), (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_227_v59), 0))
			_ = _index39
			(_227_v59).ArraySet1((_this).F23(), (_index39).Int())
			var _228_v60 _dafny.Set
			_ = _228_v60
			_228_v60 = _dafny.SetOf(!((_this).F28()), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool), false, p1)
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
			_ = _index40
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_227_v59), 0))
			_ = _index41
			var _rhs54 bool = Companion_Default___.Fm1(Companion_Default___.Fm0((_dafny.Zero).Minus((_228_v60).Cardinality()), p1, globalState), (_this).F28(), globalState)
			_ = _rhs54
			var _rhs55 bool = !((_this).F28())
			_ = _rhs55
			var _rhs56 _dafny.Int = (_this).F23()
			_ = _rhs56
			var _rhs57 bool = (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool)
			_ = _rhs57
			var _rhs58 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_227_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_227_v59), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_158_v0).Cardinality())), Companion_Default___.SafeDivisionInt((_this).F23(), Companion_Default___.Fm0((_227_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_227_v59), 0))).Int()).(_dafny.Int), p1, globalState)))
			_ = _rhs58
			var _lhs51 *GlobalState = globalState
			_ = _lhs51
			var _lhs52 _dafny.Array = p0
			_ = _lhs52
			var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
			_ = _lhs53
			var _lhs54 *GlobalState = globalState
			_ = _lhs54
			var _lhs55 _dafny.Array = _227_v59
			_ = _lhs55
			var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_227_v59), 0))
			_ = _lhs56
			_lhs51.F9 = _rhs54
			(_lhs52).ArraySet1(_rhs55, (_lhs53).Int())
			r0 = _rhs56
			_lhs54.F9 = _rhs57
			(_lhs55).ArraySet1(_rhs58, (_lhs56).Int())
			var _229_v61 D3
			_ = _229_v61
			_229_v61 = Companion_D3_.Create_DC12_(_195_v31)
			var _230_v62 D3
			_ = _230_v62
			_230_v62 = Companion_D3_.Create_DC10_(true)
			_229_v61 = Companion_Default___.Fm22(_230_v62, globalState)
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_227_v59), 0))
			_ = _index42
			(_227_v59).ArraySet1((_this.F29()).Select((Companion_Default___.SafeIndex((_227_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_227_v59), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32()).(_dafny.Int), (_index42).Int())
		}
		var _hi5 _dafny.Int = (_this).F23()
		_ = _hi5
		for _231_i7 := (_this).F23(); _231_i7.Cmp(_hi5) < 0; _231_i7 = _231_i7.Plus(_dafny.One) {
			var _232_v63 _dafny.Array
			_ = _232_v63
			var _nwElement0_5 bool = (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool)
			_ = _nwElement0_5
			var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.One)
			_ = _nw37
			(_nw37).ArraySet1(_nwElement0_5, 0)
			_232_v63 = _nw37
			_232_v63 = _232_v63
			(globalState).F11 = _160_v2
			_194_v30 = (_194_v30).Merge(_194_v30)
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
			_ = _index43
			var _rhs59 bool = Companion_Default___.Fm1((_this).F23(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState)
			_ = _rhs59
			var _rhs60 bool = (_this).F28()
			_ = _rhs60
			var _rhs61 bool = (_193_v29).Select((Companion_Default___.SafeIndex(_231_i7, _dafny.IntOfUint32((_193_v29).Cardinality()))).Uint32()).(bool)
			_ = _rhs61
			var _lhs57 _dafny.Array = p0
			_ = _lhs57
			var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((p0), 0))
			_ = _lhs58
			var _lhs59 *GlobalState = globalState
			_ = _lhs59
			var _lhs60 *GlobalState = globalState
			_ = _lhs60
			(_lhs57).ArraySet1(_rhs59, (_lhs58).Int())
			_lhs59.F7 = _rhs60
			_lhs60.F7 = _rhs61
		}
		r0 = (_this).F23()
		var _233_v64 _dafny.Set
		_ = _233_v64
		_233_v64 = _dafny.SetOf((_dafny.Zero).Minus(((_this).F23()).Times((_this).F23())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_158_v0).Cardinality())))
		r1 = _233_v64
		return r0, r1
	}
}
func (_this *C1) M2(p0 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var _rhs62 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hoiflwc"), p0)
		_ = _rhs62
		var _rhs63 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p0, p0)
		_ = _rhs63
		var _rhs64 _dafny.Int = _dafny.IntOfInt64(641)
		_ = _rhs64
		var _rhs65 _dafny.Int = ((_this).F23()).Plus((_this).F23())
		_ = _rhs65
		var _lhs61 *GlobalState = globalState
		_ = _lhs61
		var _lhs62 *GlobalState = globalState
		_ = _lhs62
		_lhs61.F6 = _rhs62
		_lhs62.F6 = _rhs63
		r0 = _rhs64
		r0 = _rhs65
		var _234_v0 _dafny.MultiSet
		_ = _234_v0
		_234_v0 = _dafny.MultiSetOf((_this).F28())
		var _235_v1 _dafny.Map
		_ = _235_v1
		_235_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_234_v0).IsSubsetOf(_dafny.MultiSetOf((_this).F28())), (_dafny.Zero).Minus((_this).F23()))
		_235_v1 = (_235_v1).Update(((_this).F28()) || ((_this).F28()), _dafny.IntOfInt64(922))
		(globalState).F14 = (_dafny.Zero).Minus((_this).F23())
		var _236_v2 _dafny.MultiSet
		_ = _236_v2
		_236_v2 = _dafny.MultiSetOf((_this).F23())
		var _237_v3 _dafny.Sequence
		_ = _237_v3
		_237_v3 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(308), _236_v2))
		var _238_v4 _dafny.Map
		_ = _238_v4
		_238_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _236_v2)
		var _239_v5 _dafny.Sequence
		_ = _239_v5
		_239_v5 = _dafny.SeqOf((_237_v3).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_237_v3).Cardinality()))).Uint32()).(_dafny.Map), (_238_v4).Update(_dafny.IntOfInt64(324), _236_v2), _238_v4)
		var _240_v6 _dafny.Map
		_ = _240_v6
		_240_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), ((_239_v5).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F23()), _dafny.IntOfUint32((_239_v5).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())
		var _241_v7 _dafny.CodePoint
		_ = _241_v7
		_241_v7 = _dafny.CodePoint('n')
		var _242_v8 D2
		_ = _242_v8
		_242_v8 = Companion_D2_.Create_DC8_((_this).F23(), (_240_v6).Cardinality(), _241_v7, (_this).F23(), (_this).F28())
		_235_v1 = (_235_v1).Update((_242_v8).Dtor_cf19(), (_dafny.Zero).Minus((_this).F23()))
		var _243_v9 _dafny.Array
		_ = _243_v9
		var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
		_ = _nw38
		_243_v9 = _nw38
		var _hi6 _dafny.Int = (_this).F23()
		_ = _hi6
		for _244_i0 := (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_243_v9)).Cardinality())); _244_i0.Cmp(_hi6) < 0; _244_i0 = _244_i0.Plus(_dafny.One) {
			(globalState).F14 = (_244_i0).Times((_this).F23())
			if !(Companion_Default___.Fm1(_244_i0, (_this).F28(), globalState)) {
				(globalState).F21 = !((_this).F28())
				var _245_v10 _dafny.Map
				_ = _245_v10
				_245_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F28())
				var _246_v11 _dafny.Sequence
				_ = _246_v11
				_246_v11 = _dafny.SeqOf((_this).F28(), (_this).F28())
				_245_v10 = (_245_v10).Update((_this).F28(), (_246_v11).Select((Companion_Default___.SafeIndex(_244_i0, _dafny.IntOfUint32((_246_v11).Cardinality()))).Uint32()).(bool))
				var _247_v12 _dafny.Array
				_ = _247_v12
				var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
				_ = _nw39
				_247_v12 = _nw39
				var _248_v13 _dafny.Array
				_ = _248_v13
				var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(3))
				_ = _nw40
				_248_v13 = _nw40
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_247_v12), 0))
				_ = _index44
				(_247_v12).ArraySet1(_248_v13, (_index44).Int())
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_247_v12), 0))
				_ = _index45
				var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(23))
				_ = _nw41
				(_247_v12).ArraySet1(_nw41, (_index45).Int())
				r0 = (_this).F23()
				var _249_v14 _dafny.Array
				_ = _249_v14
				var _len0_2 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_2
				var _nw42 _dafny.Array
				_ = _nw42
				if _len0_2.Cmp(_dafny.Zero) == 0 {
					_nw42 = _dafny.NewArray(_len0_2)
				} else {
					var _init2 func(_dafny.Int) _dafny.Sequence = (func(_250_p0 _dafny.Sequence, _251_i0 _dafny.Int, _252_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
						return func(_253_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Update(_250_p0, (Companion_Default___.SafeIndex(_251_i0, _dafny.IntOfUint32((_250_p0).Cardinality()))).Uint32(), _252_v7)
						}
					})(p0, _244_i0, _241_v7)
					_ = _init2
					var _element0_2 = _init2(_dafny.Zero)
					_ = _element0_2
					_nw42 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
					(_nw42).ArraySet1(_element0_2, 0)
					var _nativeLen0_2 = (_len0_2).Int()
					_ = _nativeLen0_2
					for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
						(_nw42).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
					}
				}
				_249_v14 = _nw42
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_249_v14), 0))
				_ = _index46
				(_249_v14).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_244_i0, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _241_v7), p0), (_index46).Int())
				var _254_v15 _dafny.Sequence
				_ = _254_v15
				_254_v15 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_244_i0, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _241_v7), (Companion_Default___.SafeIndex(_244_i0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_244_i0, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _241_v7)).Cardinality()))).Uint32(), _241_v7))
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_249_v14), 0))
				_ = _index47
				(_249_v14).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_254_v15).Select((Companion_Default___.SafeIndex(_244_i0, _dafny.IntOfUint32((_254_v15).Cardinality()))).Uint32()).(_dafny.Sequence), p0), (_index47).Int())
				var _255_v16 _dafny.Sequence
				_ = _255_v16
				_255_v16 = _dafny.SeqOf(_234_v0)
				var _256_v17 _dafny.Array
				_ = _256_v17
				var _nwElement0_6 _dafny.MultiSet = _234_v0
				_ = _nwElement0_6
				var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(21))
				_ = _nw43
				(_nw43).ArraySet1(_nwElement0_6, 0)
				(_nw43).ArraySet1(_234_v0, 1)
				(_nw43).ArraySet1(_dafny.MultiSetOf((_this).F28(), (_this).F28(), (_this).F28(), (_this).F28()), 2)
				(_nw43).ArraySet1(_234_v0, 3)
				(_nw43).ArraySet1(_234_v0, 4)
				(_nw43).ArraySet1(_234_v0, 5)
				(_nw43).ArraySet1((((_234_v0).Update((_this).F28(), Companion_Default___.Abs((_this).F23()))).Update((_this).F28(), Companion_Default___.Abs((_this).F23()))).Union(_dafny.MultiSetOf((_this).F28(), true)), 6)
				(_nw43).ArraySet1(_234_v0, 7)
				(_nw43).ArraySet1(_234_v0, 8)
				(_nw43).ArraySet1(_234_v0, 9)
				(_nw43).ArraySet1(_234_v0, 10)
				(_nw43).ArraySet1((_234_v0).Intersection(_234_v0), 11)
				(_nw43).ArraySet1((_234_v0).Update((_this).F28(), Companion_Default___.Abs((_dafny.Zero).Minus((_this).F23()))), 12)
				(_nw43).ArraySet1((_234_v0).Union(_dafny.MultiSetOf((_this).F28())), 13)
				(_nw43).ArraySet1(_234_v0, 14)
				(_nw43).ArraySet1(_234_v0, 15)
				(_nw43).ArraySet1(_234_v0, 16)
				(_nw43).ArraySet1((_234_v0).Difference(_234_v0), 17)
				(_nw43).ArraySet1((_255_v16).Select((Companion_Default___.SafeIndex(_244_i0, _dafny.IntOfUint32((_255_v16).Cardinality()))).Uint32()).(_dafny.MultiSet), 18)
				(_nw43).ArraySet1(_234_v0, 19)
				(_nw43).ArraySet1((_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F28()))).Intersection(_234_v0), 20)
				_256_v17 = _nw43
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_256_v17), 0))
				_ = _index48
				(_256_v17).ArraySet1(_234_v0, (_index48).Int())
				var _257_v18 _dafny.Set
				_ = _257_v18
				_257_v18 = _dafny.SetOf(!((_this).F28()) || (!((_this).F28())))
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_249_v14), 0))
				_ = _index49
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_249_v14), 0))
				_ = _index50
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_256_v17), 0))
				_ = _index51
				var _rhs66 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bshdxkt"), p0)
				_ = _rhs66
				var _rhs67 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(128))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}((func(_258_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_259_i2 _dafny.Int) _dafny.CodePoint {
						return _258_v7
					}
				})(_241_v7))), _dafny.UnicodeSeqOfUtf8Bytes("sgs"))
				_ = _rhs67
				var _rhs68 _dafny.MultiSet = _234_v0
				_ = _rhs68
				var _rhs69 _dafny.Set = _257_v18
				_ = _rhs69
				var _rhs70 _dafny.Sequence = p0
				_ = _rhs70
				var _lhs63 _dafny.Array = _249_v14
				_ = _lhs63
				var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_249_v14), 0))
				_ = _lhs64
				var _lhs65 _dafny.Array = _249_v14
				_ = _lhs65
				var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_249_v14), 0))
				_ = _lhs66
				var _lhs67 _dafny.Array = _256_v17
				_ = _lhs67
				var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_256_v17), 0))
				_ = _lhs68
				var _lhs69 *GlobalState = globalState
				_ = _lhs69
				(_lhs63).ArraySet1(_rhs66, (_lhs64).Int())
				(_lhs65).ArraySet1(_rhs67, (_lhs66).Int())
				(_lhs67).ArraySet1(_rhs68, (_lhs68).Int())
				_257_v18 = _rhs69
				_lhs69.F6 = _rhs70
			} else {
				(globalState).F21 = (_this).F28()
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_243_v9), 0))
				_ = _index52
				(_243_v9).ArraySet1(_244_i0, (_index52).Int())
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_243_v9), 0))
				_ = _index53
				var _rhs71 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("celkv")
				_ = _rhs71
				var _rhs72 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf((_this).F28())).Cardinality())
				_ = _rhs72
				var _lhs70 *GlobalState = globalState
				_ = _lhs70
				var _lhs71 _dafny.Array = _243_v9
				_ = _lhs71
				var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_243_v9), 0))
				_ = _lhs72
				_lhs70.F6 = _rhs71
				(_lhs71).ArraySet1(_rhs72, (_lhs72).Int())
				var _260_v19 D0
				_ = _260_v19
				_260_v19 = Companion_D0_.Create_DC1_((_this).F23(), (_this).F28(), (func() _dafny.Int {
					if (_236_v2).Contains((_this).F23()) {
						return (_236_v2).Multiplicity((_this).F23())
					}
					return (_236_v2).Cardinality()
				})())
				(globalState).F9 = (_260_v19).Dtor_cf2()
				(globalState).F21 = (_this).F28()
				var _261_v20 _dafny.Map
				_ = _261_v20
				_261_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (func() bool {
					if Companion_Default___.Fm1((_dafny.Zero).Minus((_243_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_243_v9), 0))).Int()).(_dafny.Int)), (_this).F28(), globalState) {
						return (_this).F28()
					}
					return (_this).F28()
				})())
				_261_v20 = (_261_v20).Update((_this).F23(), (_this).F28())
			}
			var _262_v21 D4
			_ = _262_v21
			_262_v21 = Companion_D4_.Create_DC15_(_dafny.IntOfInt64(-165), _241_v7, Companion_Default___.SafeModuloInt((_this.F29()).Select((Companion_Default___.SafeIndex((_236_v2).Cardinality(), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32()).(_dafny.Int), (_this).F23()))
			_262_v21 = _262_v21
			var _263_v22 _dafny.Array
			_ = _263_v22
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw44
			_263_v22 = _nw44
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_263_v22), 0))
			_ = _index54
			(_263_v22).ArraySet1((_this).F28(), (_index54).Int())
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_263_v22), 0))
			_ = _index55
			(_263_v22).ArraySet1((_242_v8).Dtor_cf19(), (_index55).Int())
		}
		var _264_i3 _dafny.Int
		_ = _264_i3
		_264_i3 = _dafny.Zero
		{
			for !((func() bool {
				if (_this).F28() {
					return (_dafny.IntOfInt64(628)).Cmp((Companion_D1_.Create_DC5_((_this).F23(), (_this).F28(), (_this).F23(), (_this).F23())).Dtor_cf9()) != 0
				}
				return (_this).F28()
			})()) {
				{
					if (_264_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_264_i3 = (_264_i3).Plus(_dafny.One)
					var _265_v23 _dafny.Sequence
					_ = _265_v23
					_265_v23 = _dafny.SeqOf((_this).F28())
					var _266_v24 D3
					_ = _266_v24
					_266_v24 = Companion_D3_.Create_DC10_((_265_v23).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_265_v23).Cardinality()))).Uint32()).(bool))
					var _267_v25 _dafny.Set
					_ = _267_v25
					_267_v25 = _dafny.SetOf((_this).F28())
					var _268_v26 _dafny.Map
					_ = _268_v26
					_268_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(329), (_this).F28())
					var _269_v27 _dafny.Map
					_ = _269_v27
					_269_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (func() bool {
						if (_268_v26).Contains((_this).F23()) {
							return (_268_v26).Get((_this).F23()).(bool)
						}
						return (_this).F28()
					})())
					var _270_v28 _dafny.Array
					_ = _270_v28
					var _nwElement0_7 bool = (_this).F28()
					_ = _nwElement0_7
					var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(29))
					_ = _nw45
					(_nw45).ArraySet1(_nwElement0_7, 0)
					(_nw45).ArraySet1(Companion_Default___.Fm1((_this).F23(), (_this).F28(), globalState), 1)
					(_nw45).ArraySet1((_this).F28(), 2)
					(_nw45).ArraySet1((func() bool {
						if (_this).F28() {
							return (_this).F28()
						}
						return (_this).F28()
					})(), 3)
					(_nw45).ArraySet1((_this).F28(), 4)
					(_nw45).ArraySet1((_this).F28(), 5)
					(_nw45).ArraySet1((_266_v24).Dtor_cf21(), 6)
					(_nw45).ArraySet1((_this).F28(), 7)
					(_nw45).ArraySet1((_this).F28(), 8)
					(_nw45).ArraySet1((_this).F28(), 9)
					(_nw45).ArraySet1((_this).F28(), 10)
					(_nw45).ArraySet1((_this).F28(), 11)
					(_nw45).ArraySet1(!((_265_v23).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_265_v23).Cardinality()))).Uint32()).(bool)) || (Companion_Default___.Fm1((_this).F23(), (_this).F28(), globalState)), 12)
					(_nw45).ArraySet1((func() bool {
						if (_this).F28() {
							return (_this).F28()
						}
						return (_this).F28()
					})(), 13)
					(_nw45).ArraySet1((_267_v25).Equals(_dafny.SetOf((_this).F28())), 14)
					(_nw45).ArraySet1((_this).F28(), 15)
					(_nw45).ArraySet1((_this).F28(), 16)
					(_nw45).ArraySet1((_this).F28(), 17)
					(_nw45).ArraySet1((_this).F28(), 18)
					(_nw45).ArraySet1((func() bool {
						if (_269_v27).Contains((_this).F28()) {
							return (_269_v27).Get((_this).F28()).(bool)
						}
						return (_this).F28()
					})(), 19)
					(_nw45).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_this.F29(), (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32(), _dafny.IntOfInt64(128)), (_this).F23()), 20)
					(_nw45).ArraySet1(!(((_dafny.Zero).Minus((_this).F23())).Cmp((_this).F23()) < 0), 21)
					(_nw45).ArraySet1((_this).F28(), 22)
					(_nw45).ArraySet1((_this).F28(), 23)
					(_nw45).ArraySet1(((func() _dafny.Int {
						if (_235_v1).Contains((_this).F28()) {
							return (_235_v1).Get((_this).F28()).(_dafny.Int)
						}
						return (_this).F23()
					})()).Cmp(_dafny.IntOfInt64(-581)) > 0, 24)
					(_nw45).ArraySet1((_this).F28(), 25)
					(_nw45).ArraySet1((_265_v23).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_265_v23).Cardinality()))).Uint32()).(bool), 26)
					(_nw45).ArraySet1(!((_this).F28()), 27)
					(_nw45).ArraySet1(((_this).F23()).Cmp((_this).F23()) < 0, 28)
					_270_v28 = _nw45
					var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_270_v28), 0))
					_ = _index56
					(_270_v28).ArraySet1((func() bool {
						if Companion_Default___.Fm1((_this).F23(), (_this).F28(), globalState) {
							return (_this).F28()
						}
						return (_this).F28()
					})(), (_index56).Int())
					var _271_v29 _dafny.MultiSet
					_ = _271_v29
					_271_v29 = _dafny.MultiSetOf(_243_v9, _243_v9, _243_v9)
					var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_270_v28), 0))
					_ = _index57
					var _rhs73 bool = (_dafny.MultiSetOf(_243_v9, _243_v9)).IsSubsetOf(_271_v29)
					_ = _rhs73
					var _rhs74 _dafny.Int = (_this).F23()
					_ = _rhs74
					var _rhs75 bool = Companion_Default___.Fm1((_this).F23(), (_this).F28(), globalState)
					_ = _rhs75
					var _rhs76 bool = !((_this).F28()) || ((_this).F28())
					_ = _rhs76
					var _rhs77 bool = Companion_Default___.Fm1((_this).F23(), (_this).F28(), globalState)
					_ = _rhs77
					var _lhs73 _dafny.Array = _270_v28
					_ = _lhs73
					var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_270_v28), 0))
					_ = _lhs74
					var _lhs75 *GlobalState = globalState
					_ = _lhs75
					var _lhs76 *GlobalState = globalState
					_ = _lhs76
					var _lhs77 *GlobalState = globalState
					_ = _lhs77
					(_lhs73).ArraySet1(_rhs73, (_lhs74).Int())
					r0 = _rhs74
					_lhs75.F21 = _rhs75
					_lhs76.F7 = _rhs76
					_lhs77.F21 = _rhs77
					var _272_v30 _dafny.Map
					_ = _272_v30
					_272_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
						if !(true) {
							return _265_v23
						}
						return _dafny.SeqOf((_this).F28())
					})(), (_this).F28())
					var _273_v31 _dafny.Map
					_ = _273_v31
					_273_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_241_v7, p0)
					_272_v30 = (_272_v30).Update(_265_v23, !((_270_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_270_v28), 0))).Int()).(bool)) || ((_265_v23).Select((Companion_Default___.SafeIndex((_273_v31).Cardinality(), _dafny.IntOfUint32((_265_v23).Cardinality()))).Uint32()).(bool)))
					var _274_v32 *C0
					_ = _274_v32
					var _nw46 *C0 = New_C0_()
					_ = _nw46
					_nw46.Ctor__()
					_274_v32 = _nw46
					var _275_v33 *C0
					_ = _275_v33
					var _nw47 *C0 = New_C0_()
					_ = _nw47
					_nw47.Ctor__()
					_275_v33 = _nw47
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		r0 = (((_this).F23()).Minus((_this).F23())).Plus(_dafny.IntOfInt64(712))
		r1 = _dafny.Companion_Sequence_.Concatenate(_this.F29(), (func() _dafny.Sequence {
			if Companion_Default___.Fm1((_this).F23(), !((_this).F28()), globalState) {
				return _this.F29()
			}
			return _this.F29()
		})())
		return r0, r1
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f29 _dafny.Sequence
	_f23 _dafny.Int
	_f28 bool
	_f31 _dafny.Map
	_f32 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f29 = _dafny.EmptySeq
	_this._f23 = _dafny.Zero
	_this._f28 = false
	_this._f31 = _dafny.EmptyMap
	_this._f32 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C2{}
var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F29() _dafny.Sequence {
	return _this._f29
}
func (_this *C2) F29_set_(value _dafny.Sequence) {
	_this._f29 = value
}
func (_this *C2) F23() _dafny.Int {
	return _this._f23
}
func (_this *C2) F28() bool {
	return _this._f28
}
func (_this *C2) Ctor__(f31 _dafny.Map, f32 bool, f28 bool, f29 _dafny.Sequence, f23 _dafny.Int) {
	{
		(_this)._f31 = f31
		(_this)._f32 = f32
		(_this)._f28 = f28
		(_this)._f29 = f29
		(_this)._f23 = f23
	}
}
func (_this *C2) Fm18(p0 _dafny.Map, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((((func() _dafny.MultiSet {
			if (_this).F28() {
				return _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(779))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg23 _dafny.Int) interface{} {
						return coer23(arg23)
					}
				}(func(_276_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('x')
				})))
			}
			return _dafny.MultiSetOf(_dafny.CodePoint('k'), _dafny.CodePoint('g'))
		})()).Intersection((_dafny.MultiSetOf(_dafny.CodePoint('s'))).Intersection(_dafny.MultiSetOf(_dafny.CodePoint('n'))))).Cardinality())
	}
}
func (_this *C2) M1(p0 _dafny.Array, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Set) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Set = _dafny.EmptySet
		_ = r1
		var _277_v0 _dafny.Sequence
		_ = _277_v0
		_277_v0 = _dafny.UnicodeSeqOfUtf8Bytes("rovoqj")
		var _278_v1 _dafny.Map
		_ = _278_v1
		_278_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_277_v0, true)
		var _279_v2 _dafny.Map
		_ = _279_v2
		_279_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_277_v0, _277_v0), (_278_v1).Merge(_278_v1))
		_279_v2 = (_279_v2).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-392))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}(func(_280_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		})), (_278_v1).Merge(_278_v1))
		var _281_v3 _dafny.Array
		_ = _281_v3
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_3
		var _nw48 _dafny.Array
		_ = _nw48
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw48 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Sequence = (func(_282_p1 bool) func(_dafny.Int) _dafny.Sequence {
				return func(_283_i1 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F32(), _282_p1), _dafny.SeqOf(_282_p1))
				}
			})(p1)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw48 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw48).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw48).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_281_v3 = _nw48
		_281_v3 = _281_v3
		var _284_i2 _dafny.Int
		_ = _284_i2
		_284_i2 = _dafny.Zero
		{
			for ((_this).F23()).Cmp((_this).F23()) <= 0 {
				{
					if (_284_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_284_i2 = (_284_i2).Plus(_dafny.One)
					var _285_v4 _dafny.MultiSet
					_ = _285_v4
					_285_v4 = _dafny.MultiSetOf(_277_v0, _277_v0)
					var _286_v5 T0
					_ = _286_v5
					var _nw49 *C1 = New_C1_()
					_ = _nw49
					_nw49.Ctor__(false, _this.F29(), _dafny.IntOfUint32((_277_v0).Cardinality()))
					_286_v5 = _nw49
					var _287_v6 _dafny.MultiSet
					_ = _287_v6
					_287_v6 = _dafny.MultiSetOf(_286_v5, _286_v5, _286_v5)
					var _288_v7 _dafny.CodePoint
					_ = _288_v7
					_288_v7 = _dafny.CodePoint('f')
					var _289_v8 _dafny.Array
					_ = _289_v8
					var _nwElement0_8 _dafny.Int = (_this).F23()
					_ = _nwElement0_8
					var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(11))
					_ = _nw50
					(_nw50).ArraySet1(_nwElement0_8, 0)
					(_nw50).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
						if (_285_v4).Contains(_277_v0) {
							return (_285_v4).Multiplicity(_277_v0)
						}
						return (_this).F23()
					})()), 1)
					(_nw50).ArraySet1((_dafny.SetOf(_dafny.IntOfInt64(-334), (_287_v6).Cardinality())).Cardinality(), 2)
					(_nw50).ArraySet1(_dafny.IntOfInt64(118), 3)
					(_nw50).ArraySet1((_286_v5).F23(), 4)
					(_nw50).ArraySet1((_this.F29()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32()).(_dafny.Int), 5)
					(_nw50).ArraySet1((_this).F23(), 6)
					(_nw50).ArraySet1((_this).F23(), 7)
					(_nw50).ArraySet1((_this).F23(), 8)
					(_nw50).ArraySet1((_this).F23(), 9)
					(_nw50).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_288_v7)).Cardinality()), 10)
					_289_v8 = _nw50
					var _290_v9 _dafny.Sequence
					_ = _290_v9
					_290_v9 = _dafny.SeqOf(_289_v8, _289_v8, _289_v8, _289_v8, _289_v8)
					var _291_v10 _dafny.Map
					_ = _291_v10
					_291_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v9, (_286_v5).F23())
					_291_v10 = (_291_v10).Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_289_v8), _290_v9), (Companion_Default___.SafeIndex((_286_v5).F23(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_289_v8), _290_v9)).Cardinality()))).Uint32(), _289_v8), (_this).F23())
					var _292_v11 _dafny.Array
					_ = _292_v11
					var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
					_ = _nw51
					_292_v11 = _nw51
					var _293_v12 *C0
					_ = _293_v12
					var _nw52 *C0 = New_C0_()
					_ = _nw52
					_nw52.Ctor__()
					_293_v12 = _nw52
					var _294_v13 _dafny.Sequence
					_ = _294_v13
					_294_v13 = _dafny.SeqOf(_293_v12)
					var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_292_v11), 0))
					_ = _index58
					(_292_v11).ArraySet1(_294_v13, (_index58).Int())
					var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_292_v11), 0))
					_ = _index59
					(_292_v11).ArraySet1((func() _dafny.Sequence {
						if (_this).F28() {
							return _294_v13
						}
						return _294_v13
					})(), (_index59).Int())
					var _295_v14 D1
					_ = _295_v14
					_295_v14 = Companion_D1_.Create_DC4_((_this).F31(), p1, true)
					if (_295_v14).Dtor_cf8() {
						var _296_v15 D1
						_ = _296_v15
						_296_v15 = Companion_D1_.Create_DC3_(_290_v9)
						_296_v15 = _296_v15
						(globalState).F9 = p1
						var _297_v16 _dafny.MultiSet
						_ = _297_v16
						_297_v16 = _dafny.MultiSetOf((_this).F28())
						(globalState).F9 = ((_297_v16).IsSubsetOf(_297_v16)) || (p1)
						var _rhs78 bool = !((Companion_Default___.SafeModuloInt((_286_v5).F23(), (_286_v5).F23())).Cmp((_286_v5).F23()) <= 0)
						_ = _rhs78
						var _rhs79 _dafny.Int = (_dafny.Zero).Minus((_this).F23())
						_ = _rhs79
						var _rhs80 _dafny.Int = (_286_v5).F23()
						_ = _rhs80
						var _lhs78 *GlobalState = globalState
						_ = _lhs78
						var _lhs79 *GlobalState = globalState
						_ = _lhs79
						var _lhs80 *GlobalState = globalState
						_ = _lhs80
						_lhs78.F7 = _rhs78
						_lhs79.F8 = _rhs79
						_lhs80.F8 = _rhs80
						(globalState).F7 = ((_this).F32()) == (_dafny.Companion_Sequence_.IsProperPrefixOf(_277_v0, _dafny.Companion_Sequence_.Update(_277_v0, (Companion_Default___.SafeIndex((_286_v5).F23(), _dafny.IntOfUint32((_277_v0).Cardinality()))).Uint32(), _288_v7)))
					} else {
						var _298_v18 _dafny.Sequence
						_ = _298_v18
						_298_v18 = _dafny.SeqOf(_277_v0, _dafny.UnicodeSeqOfUtf8Bytes("tocc"))
						var _299_v19 _dafny.Map
						_ = _299_v19
						_299_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_286_v5).F23(), (_286_v5).F23()), func() _dafny.Map {
							var _coll17 = _dafny.NewMapBuilder()
							_ = _coll17
							for _iter18 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_298_v18, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.IntOfUint32((_298_v18).Cardinality()))).Uint32(), _277_v0)).Elements()); ; {
								_compr_17, _ok18 := _iter18()
								if !_ok18 {
									break
								}
								var _300_v17 _dafny.Sequence
								_300_v17 = interface{}(_compr_17).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_298_v18, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.IntOfUint32((_298_v18).Cardinality()))).Uint32(), _277_v0), _300_v17) {
									_coll17.Add(_300_v17, (_dafny.MultiSetOf((_this).F28())).Cardinality())
								}
							}
							return _coll17.ToMap()
						}())
						var _301_v20 _dafny.Map
						_ = _301_v20
						_301_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_277_v0, (_dafny.Zero).Minus((_286_v5).F23()))
						_299_v19 = (_299_v19).Update((_this).F23(), _301_v20)
						var _302_v21 _dafny.Map
						_ = _302_v21
						_302_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_286_v5).F23(), p1, globalState), p1)
						_302_v21 = (_302_v21).Update(((_this).F23()).Times((_286_v5).F23()), Companion_Default___.Fm1((_286_v5).F23(), (_this).F28(), globalState))
						(globalState).F14 = (Companion_Default___.Fm0((_286_v5).F23(), p1, globalState)).Times((_dafny.IntOfInt64(-677)).Times((_286_v5).F23()))
						var _303_v22 _dafny.Map
						_ = _303_v22
						_303_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _dafny.IntOfInt64(368))
						_303_v22 = (_303_v22).Update(p1, (_dafny.IntOfUint32((Companion_Default___.Fm2(globalState)).Cardinality())).Minus((_286_v5).F23()))
						var _rhs81 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_277_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-846))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg25 _dafny.Int) interface{} {
								return coer25(arg25)
							}
						}((func(_304_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_305_i3 _dafny.Int) _dafny.CodePoint {
								return _304_v7
							}
						})(_288_v7)))), _277_v0)
						_ = _rhs81
						var _rhs82 _dafny.Int = (_this).F23()
						_ = _rhs82
						var _lhs81 *GlobalState = globalState
						_ = _lhs81
						_277_v0 = _rhs81
						_lhs81.F14 = _rhs82
					}
					var _306_v23 _dafny.Array
					_ = _306_v23
					var _nw53 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(2))
					_ = _nw53
					_306_v23 = _nw53
					var _307_v24 D0
					_ = _307_v24
					_307_v24 = Companion_D0_.Create_DC1_((_this).F23(), (_this).F32(), (_this).F23())
					var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_306_v23), 0))
					_ = _index60
					(_306_v23).ArraySet1(_307_v24, (_index60).Int())
					var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_306_v23), 0))
					_ = _index61
					(_306_v23).ArraySet1(_307_v24, (_index61).Int())
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _308_v25 _dafny.Set
		_ = _308_v25
		_308_v25 = _dafny.SetOf(p1)
		_308_v25 = _308_v25
		var _309_i4 _dafny.Int
		_ = _309_i4
		_309_i4 = _dafny.Zero
		{
			for (_this).F28() {
				{
					if (_309_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_309_i4 = (_309_i4).Plus(_dafny.One)
					var _310_v26 _dafny.Sequence
					_ = _310_v26
					_310_v26 = _dafny.SeqOf((_this).F32(), (_this).F32())
					var _311_v27 _dafny.MultiSet
					_ = _311_v27
					_311_v27 = _dafny.MultiSetOf(Companion_Default___.Fm0((_this).F23(), p1, globalState), (_this).F23())
					var _312_v28 _dafny.Sequence
					_ = _312_v28
					_312_v28 = _dafny.SeqOf(p0)
					var _313_v29 _dafny.MultiSet
					_ = _313_v29
					_313_v29 = _dafny.MultiSetOf(_277_v0)
					var _314_v30 _dafny.Array
					_ = _314_v30
					var _nwElement0_9 _dafny.Int = Companion_Default___.Fm0((_this).F23(), (_this).F32(), globalState)
					_ = _nwElement0_9
					var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(26))
					_ = _nw54
					(_nw54).ArraySet1(_nwElement0_9, 0)
					(_nw54).ArraySet1((_this).F23(), 1)
					(_nw54).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_310_v26).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_310_v26).Cardinality()))).Uint32()).(bool) {
							return _277_v0
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(11))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg26 _dafny.Int) interface{} {
								return coer26(arg26)
							}
						}(func(_315_i5 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('k')
						}))
					})()).Cardinality()), 2)
					(_nw54).ArraySet1((_this).F23(), 3)
					(_nw54).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hpcxes")).Cardinality()), 4)
					(_nw54).ArraySet1((_this).F23(), 5)
					(_nw54).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(481))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg27 _dafny.Int) interface{} {
							return coer27(arg27)
						}
					}(func(_316_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('g')
					})), _277_v0)).Cardinality()), 6)
					(_nw54).ArraySet1((_this).F23(), 7)
					(_nw54).ArraySet1((_this).F23(), 8)
					(_nw54).ArraySet1((_this).F23(), 9)
					(_nw54).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F23(), (_this).F23()), 10)
					(_nw54).ArraySet1((_this).F23(), 11)
					(_nw54).ArraySet1((func() _dafny.Int {
						if p1 {
							return (_this).F23()
						}
						return _dafny.IntOfInt64(222)
					})(), 12)
					(_nw54).ArraySet1((_dafny.Zero).Minus((_this).F23()), 13)
					(_nw54).ArraySet1(_dafny.IntOfInt64(583), 14)
					(_nw54).ArraySet1((_this).F23(), 15)
					(_nw54).ArraySet1((_dafny.IntOfUint32((_277_v0).Cardinality())).Times(((_dafny.MultiSetOf(p1)).Update(p1, Companion_Default___.Abs((_this).F23()))).Cardinality()), 16)
					(_nw54).ArraySet1((_this).F23(), 17)
					(_nw54).ArraySet1((_this).F23(), 18)
					(_nw54).ArraySet1(((_this).F23()).Times((_this).F23()), 19)
					(_nw54).ArraySet1((_308_v25).Cardinality(), 20)
					(_nw54).ArraySet1((_this).F23(), 21)
					(_nw54).ArraySet1((func() _dafny.Int {
						if (_311_v27).Contains(_dafny.IntOfUint32((_312_v28).Cardinality())) {
							return (_311_v27).Multiplicity(_dafny.IntOfUint32((_312_v28).Cardinality()))
						}
						return (_this).F23()
					})(), 22)
					(_nw54).ArraySet1((_313_v29).Cardinality(), 23)
					(_nw54).ArraySet1((_this).F23(), 24)
					(_nw54).ArraySet1((_this).F23(), 25)
					_314_v30 = _nw54
					_314_v30 = _314_v30
					var _317_v31 _dafny.MultiSet
					_ = _317_v31
					_317_v31 = _dafny.MultiSetOf((_this).F32(), false)
					var _318_v32 _dafny.Sequence
					_ = _318_v32
					_318_v32 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("thg"), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("suocs"), (Companion_Default___.SafeIndex((_317_v31).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("suocs")).Cardinality()))).Uint32(), _dafny.CodePoint('p')), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(247))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg28 _dafny.Int) interface{} {
							return coer28(arg28)
						}
					}((func(_319_v0 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
						return func(_320_i7 _dafny.Int) _dafny.CodePoint {
							return (Companion_D4_.Create_DC15_((_this).F23(), _dafny.CodePoint('f'), _dafny.IntOfUint32((_319_v0).Cardinality()))).Dtor_cf27()
						}
					})(_277_v0))), _277_v0, _277_v0)
					var _321_v33 _dafny.Sequence
					_ = _321_v33
					_321_v33 = _dafny.SeqOf(_310_v26)
					var _322_v34 _dafny.Map
					_ = _322_v34
					_322_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_318_v32).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(_dafny.IntOfInt64(348), !((_this).F28()), globalState), _dafny.IntOfUint32((_318_v32).Cardinality()))).Uint32()).(_dafny.Sequence), _321_v33)
					_322_v34 = (_322_v34).Update(_277_v0, _321_v33)
					var _323_v35 _dafny.Map
					_ = _323_v35
					_323_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _314_v30)
					var _324_v36 _dafny.Array
					_ = _324_v36
					var _nwElement0_10 _dafny.Array = (func() _dafny.Array {
						if (_this).F32() {
							return _314_v30
						}
						return _314_v30
					})()
					_ = _nwElement0_10
					var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(13))
					_ = _nw55
					(_nw55).ArraySet1(_nwElement0_10, 0)
					(_nw55).ArraySet1(_314_v30, 1)
					(_nw55).ArraySet1(_314_v30, 2)
					(_nw55).ArraySet1(_314_v30, 3)
					(_nw55).ArraySet1(_314_v30, 4)
					(_nw55).ArraySet1((func() _dafny.Array {
						if (_323_v35).Contains((_this).F23()) {
							return (_323_v35).Get((_this).F23()).(_dafny.Array)
						}
						return _314_v30
					})(), 5)
					(_nw55).ArraySet1(_314_v30, 6)
					(_nw55).ArraySet1(_314_v30, 7)
					(_nw55).ArraySet1(_314_v30, 8)
					(_nw55).ArraySet1(_314_v30, 9)
					(_nw55).ArraySet1(_314_v30, 10)
					(_nw55).ArraySet1(_314_v30, 11)
					(_nw55).ArraySet1(_314_v30, 12)
					_324_v36 = _nw55
					var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_324_v36), 0))
					_ = _index62
					(_324_v36).ArraySet1(_314_v30, (_index62).Int())
					var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_324_v36), 0))
					_ = _index63
					(_324_v36).ArraySet1(_314_v30, (_index63).Int())
					var _325_v37 *C0
					_ = _325_v37
					var _nw56 *C0 = New_C0_()
					_ = _nw56
					_nw56.Ctor__()
					_325_v37 = _nw56
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _326_v38 _dafny.Sequence
		_ = _326_v38
		_326_v38 = _dafny.SeqOf((_this).F32())
		var _327_v39 _dafny.MultiSet
		_ = _327_v39
		_327_v39 = _dafny.MultiSetOf((_this).F23())
		_326_v38 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_326_v38, (Companion_Default___.SafeIndex(Companion_Default___.Fm0((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_327_v39).Cardinality(), _dafny.MultiSetFromSeq(_326_v38))).Cardinality(), !(p1), globalState), _dafny.IntOfUint32((_326_v38).Cardinality()))).Uint32(), Companion_Default___.Fm1((_this).F23(), (_this).F32(), globalState)), _326_v38)
		r0 = (_this).F23()
		var _328_v40 _dafny.Set
		_ = _328_v40
		_328_v40 = _dafny.SetOf((_this).F23(), _dafny.IntOfUint32((_this.F29()).Cardinality()))
		r1 = (_328_v40).Union(_328_v40)
		return r0, r1
	}
}
func (_this *C2) M2(p0 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var _329_v0 _dafny.Array
		_ = _329_v0
		var _nw57 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
		_ = _nw57
		_329_v0 = _nw57
		var _330_v1 _dafny.Sequence
		_ = _330_v1
		_330_v1 = _dafny.SeqOf(p0, p0, _dafny.Companion_Sequence_.Concatenate(p0, p0))
		var _331_v2 D0
		_ = _331_v2
		_331_v2 = Companion_D0_.Create_DC0_(p0)
		var _332_v3 D1
		_ = _332_v3
		_332_v3 = Companion_D1_.Create_DC5_((_this).F23(), (_this).F32(), (_this).F23(), (_this).F23())
		var _pat_let_tv12 = p0
		_ = _pat_let_tv12
		var _pat_let_tv13 = p0
		_ = _pat_let_tv13
		var _pat_let_tv14 = p0
		_ = _pat_let_tv14
		var _pat_let_tv15 = p0
		_ = _pat_let_tv15
		var _pat_let_tv16 = p0
		_ = _pat_let_tv16
		var _pat_let_tv17 = p0
		_ = _pat_let_tv17
		var _pat_let_tv18 = p0
		_ = _pat_let_tv18
		var _pat_let_tv19 = p0
		_ = _pat_let_tv19
		var _pat_let_tv20 = _330_v1
		_ = _pat_let_tv20
		var _pat_let_tv21 = p0
		_ = _pat_let_tv21
		var _pat_let_tv22 = p0
		_ = _pat_let_tv22
		var _pat_let_tv23 = p0
		_ = _pat_let_tv23
		var _pat_let_tv24 = p0
		_ = _pat_let_tv24
		var _pat_let_tv25 = _330_v1
		_ = _pat_let_tv25
		var _rhs83 bool = _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("wcapto"), p0)
		_ = _rhs83
		var _rhs84 _dafny.Array = _329_v0
		_ = _rhs84
		var _rhs85 bool = func(_source1 D0) bool {
			if _source1.Is_DC1() {
				var _333___mcc_h0 _dafny.Int = _source1.Get_().(D0_DC1).Cf1
				_ = _333___mcc_h0
				var _334___mcc_h1 bool = _source1.Get_().(D0_DC1).Cf2
				_ = _334___mcc_h1
				var _335___mcc_h2 _dafny.Int = _source1.Get_().(D0_DC1).Cf3
				_ = _335___mcc_h2
				var _336_cf3 _dafny.Int = _335___mcc_h2
				_ = _336_cf3
				var _337_cf2 bool = _334___mcc_h1
				_ = _337_cf2
				var _338_cf1 _dafny.Int = _333___mcc_h0
				_ = _338_cf1
				if _337_cf2 {
					return _337_cf2
				} else {
					return (_this).F32()
				}
			} else if _source1.Is_DC0() {
				var _339___mcc_h3 _dafny.Sequence = _source1.Get_().(D0_DC0).Cf0
				_ = _339___mcc_h3
				var _340_cf0 _dafny.Sequence = _339___mcc_h3
				_ = _340_cf0
				return (_this).F28()
			} else {
				var _341___mcc_h4 D0 = _source1.Get_().(D0_DC2).Cf4
				_ = _341___mcc_h4
				var _342_cf4 D0 = _341___mcc_h4
				_ = _342_cf4
				return !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_pat_let_tv12, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_pat_let_tv13).Cardinality()))).Uint32(), (Companion_D2_.Create_DC8_((_this).F23(), (_this).F23(), _dafny.CodePoint('o'), _dafny.IntOfUint32((_pat_let_tv14).Cardinality()), (_this).F28())).Dtor_cf17()), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_pat_let_tv15, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_pat_let_tv16).Cardinality()))).Uint32(), _dafny.CodePoint('i')), (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_pat_let_tv17, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_pat_let_tv18).Cardinality()))).Uint32(), _dafny.CodePoint('i'))).Cardinality()))).Uint32(), _dafny.CodePoint('b')))
			}
		}(_331_v2)
		_ = _rhs85
		var _rhs86 _dafny.Sequence = func(_source2 D1) _dafny.Sequence {
			if _source2.Is_DC4() {
				var _343___mcc_h5 _dafny.Map = _source2.Get_().(D1_DC4).Cf6
				_ = _343___mcc_h5
				var _344___mcc_h6 bool = _source2.Get_().(D1_DC4).Cf7
				_ = _344___mcc_h6
				var _345___mcc_h7 bool = _source2.Get_().(D1_DC4).Cf8
				_ = _345___mcc_h7
				var _346_cf8 bool = _345___mcc_h7
				_ = _346_cf8
				var _347_cf7 bool = _344___mcc_h6
				_ = _347_cf7
				var _348_cf6 _dafny.Map = _343___mcc_h5
				_ = _348_cf6
				return _dafny.SeqOf(_pat_let_tv19)
			} else if _source2.Is_DC5() {
				var _349___mcc_h8 _dafny.Int = _source2.Get_().(D1_DC5).Cf9
				_ = _349___mcc_h8
				var _350___mcc_h9 bool = _source2.Get_().(D1_DC5).Cf10
				_ = _350___mcc_h9
				var _351___mcc_h10 _dafny.Int = _source2.Get_().(D1_DC5).Cf11
				_ = _351___mcc_h10
				var _352___mcc_h11 _dafny.Int = _source2.Get_().(D1_DC5).Cf12
				_ = _352___mcc_h11
				var _353_cf12 _dafny.Int = _352___mcc_h11
				_ = _353_cf12
				var _354_cf11 _dafny.Int = _351___mcc_h10
				_ = _354_cf11
				var _355_cf10 bool = _350___mcc_h9
				_ = _355_cf10
				var _356_cf9 _dafny.Int = _349___mcc_h8
				_ = _356_cf9
				return _dafny.Companion_Sequence_.Concatenate((_pat_let_tv20), _dafny.SeqOf(_pat_let_tv21))
			} else if _source2.Is_DC6() {
				var _357___mcc_h12 _dafny.Array = _source2.Get_().(D1_DC6).Cf13
				_ = _357___mcc_h12
				var _358_cf13 _dafny.Array = _357___mcc_h12
				_ = _358_cf13
				return _dafny.SeqOf(_pat_let_tv22, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-44))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg29 _dafny.Int) interface{} {
						return coer29(arg29)
					}
				}(func(_359_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				})), _pat_let_tv23, _pat_let_tv24)
			} else {
				var _360___mcc_h13 _dafny.Sequence = _source2.Get_().(D1_DC3).Cf5
				_ = _360___mcc_h13
				var _361_cf5 _dafny.Sequence = _360___mcc_h13
				_ = _361_cf5
				return _pat_let_tv25
			}
		}(_332_v3)
		_ = _rhs86
		var _lhs82 *GlobalState = globalState
		_ = _lhs82
		var _lhs83 *GlobalState = globalState
		_ = _lhs83
		_lhs82.F7 = _rhs83
		_329_v0 = _rhs84
		_lhs83.F9 = _rhs85
		_330_v1 = _rhs86
		var _362_v4 D1
		_ = _362_v4
		_362_v4 = Companion_D1_.Create_DC6_(_329_v0)
		var _363_v5 _dafny.MultiSet
		_ = _363_v5
		_363_v5 = _dafny.MultiSetOf((_362_v4).Dtor_cf13())
		_363_v5 = (_363_v5).Union(_363_v5)
		(globalState).F21 = !((_this).F32()) || ((_this).F28())
		var _hi7 _dafny.Int = (_this).F23()
		_ = _hi7
		for _364_i1 := ((_this).F23()).Minus((_this).F23()); _364_i1.Cmp(_hi7) < 0; _364_i1 = _364_i1.Plus(_dafny.One) {
			var _365_v6 _dafny.MultiSet
			_ = _365_v6
			_365_v6 = _dafny.MultiSetOf((_this).F23())
			if ((_365_v6).Difference(_365_v6)).Contains(_364_i1) {
				(globalState).F7 = true
				var _366_v7 _dafny.Map
				_ = _366_v7
				_366_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F23()), (_this).F32())
				_366_v7 = (func() _dafny.Map {
					if !(false) {
						return _366_v7
					}
					return _366_v7
				})()
				var _367_v8 _dafny.Map
				_ = _367_v8
				_367_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_364_i1, (_dafny.Zero).Minus(((_this).F23()).Times((_this).F23())))
				_367_v8 = (_367_v8).Update(_364_i1, ((_this).F31()).Cardinality())
				(globalState).F8 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.Companion_Sequence_.Concatenate(p0, p0)), (Companion_Default___.SafeIndex((_364_i1).Minus(((_363_v5).Update(_329_v0, Companion_Default___.Abs((_this).F23()))).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, _dafny.Companion_Sequence_.Concatenate(p0, p0))).Cardinality()))).Uint32(), _dafny.CodePoint('b'))).Cardinality())
				var _368_v9 _dafny.MultiSet
				_ = _368_v9
				_368_v9 = _dafny.MultiSetOf((_this).F32(), (_this).F28())
				var _369_v10 _dafny.Sequence
				_ = _369_v10
				_369_v10 = _dafny.SeqOf((func() bool {
					if (_366_v7).Contains((_this).F23()) {
						return (_366_v7).Get((_this).F23()).(bool)
					}
					return (_this).F28()
				})(), (_this).F32())
				var _370_v11 _dafny.Map
				_ = _370_v11
				_370_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F28()), (_this).F28())
				var _371_v12 _dafny.Sequence
				_ = _371_v12
				_371_v12 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F28()), _370_v11)
				var _372_v13 _dafny.Array
				_ = _372_v13
				var _nwElement0_11 _dafny.Int = Companion_Default___.Fm0(_364_i1, (_this).F32(), globalState)
				_ = _nwElement0_11
				var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(24))
				_ = _nw58
				(_nw58).ArraySet1(_nwElement0_11, 0)
				(_nw58).ArraySet1((_this).F23(), 1)
				(_nw58).ArraySet1(_364_i1, 2)
				(_nw58).ArraySet1(_364_i1, 3)
				(_nw58).ArraySet1(_364_i1, 4)
				(_nw58).ArraySet1((_this).F23(), 5)
				(_nw58).ArraySet1(_dafny.IntOfUint32((_369_v10).Cardinality()), 6)
				(_nw58).ArraySet1(_dafny.IntOfInt64(362), 7)
				(_nw58).ArraySet1(_364_i1, 8)
				(_nw58).ArraySet1(((_371_v12).Select((Companion_Default___.SafeIndex(_364_i1, _dafny.IntOfUint32((_371_v12).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), 9)
				(_nw58).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 10)
				(_nw58).ArraySet1((_this).F23(), 11)
				(_nw58).ArraySet1(((_368_v9).Update(true, Companion_Default___.Abs(_364_i1))).Cardinality(), 12)
				(_nw58).ArraySet1(Companion_Default___.Fm0((_this).F23(), (_this).F28(), globalState), 13)
				(_nw58).ArraySet1((_this).F23(), 14)
				(_nw58).ArraySet1(_dafny.IntOfInt64(384), 15)
				(_nw58).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 16)
				(_nw58).ArraySet1(_364_i1, 17)
				(_nw58).ArraySet1(_dafny.IntOfInt64(501), 18)
				(_nw58).ArraySet1((_368_v9).Cardinality(), 19)
				(_nw58).ArraySet1((_this).F23(), 20)
				(_nw58).ArraySet1(_dafny.IntOfUint32((_369_v10).Cardinality()), 21)
				(_nw58).ArraySet1(_364_i1, 22)
				(_nw58).ArraySet1(_364_i1, 23)
				_372_v13 = _nw58
				var _373_v14 *C1
				_ = _373_v14
				var _nw59 *C1 = New_C1_()
				_ = _nw59
				_nw59.Ctor__((_this).F32(), _this.F29(), (func() _dafny.Int {
					if (_368_v9).Contains(!((_this).F28())) {
						return (_368_v9).Multiplicity(!((_this).F28()))
					}
					return (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _372_v13)).Cardinality())
				})())
				_373_v14 = _nw59
			} else {
				var _374_v15 _dafny.Array
				_ = _374_v15
				var _len0_4 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_4
				var _nw60 _dafny.Array
				_ = _nw60
				if _len0_4.Cmp(_dafny.Zero) == 0 {
					_nw60 = _dafny.NewArray(_len0_4)
				} else {
					var _init4 func(_dafny.Int) _dafny.Int = func(_375_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_375_i2, (_dafny.Zero).Minus((_this).F23()))
					}
					_ = _init4
					var _element0_4 = _init4(_dafny.Zero)
					_ = _element0_4
					_nw60 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
					(_nw60).ArraySet1(_element0_4, 0)
					var _nativeLen0_4 = (_len0_4).Int()
					_ = _nativeLen0_4
					for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
						(_nw60).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
					}
				}
				_374_v15 = _nw60
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_374_v15), 0))
				_ = _index64
				(_374_v15).ArraySet1(_364_i1, (_index64).Int())
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_374_v15), 0))
				_ = _index65
				(_374_v15).ArraySet1((_this).F23(), (_index65).Int())
				(globalState).F21 = ((_this).F32()) || (!((_this).F32()))
				(globalState).F7 = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(p0, p0), Companion_Default___.Fm2(globalState))
				var _376_v16 _dafny.Map
				_ = _376_v16
				_376_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F32()) == ((_this).F28()), _374_v15)
				_376_v16 = (_376_v16).Merge(_376_v16)
				var _377_v17 _dafny.Map
				_ = _377_v17
				_377_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_364_i1, (_dafny.MultiSetOf(true, (_this).F32(), (_this).F32())).IsDisjointFrom(_dafny.MultiSetOf(true)))
				(globalState).F9 = !((func() bool {
					if (_377_v17).Contains(_364_i1) {
						return (_377_v17).Get(_364_i1).(bool)
					}
					return (_this).F32()
				})())
			}
			var _378_v18 _dafny.Map
			_ = _378_v18
			_378_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-469))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}(func(_379_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			})))
			var _380_v19 *C1
			_ = _380_v19
			var _nw61 *C1 = New_C1_()
			_ = _nw61
			_nw61.Ctor__((_this).F32(), _this.F29(), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_378_v18).Contains((_this).F32()) {
					return (_378_v18).Get((_this).F32()).(_dafny.Sequence)
				}
				return p0
			})()).Cardinality()))
			_380_v19 = _nw61
			var _381_v20 _dafny.Sequence
			_ = _381_v20
			_381_v20 = _dafny.SeqOf((_this).F28(), true, (_this).F32(), (_this).F32(), Companion_Default___.Fm1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(131))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}(func(_382_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			}))).Cardinality()), (_this).F32(), globalState))
			var _383_v21 _dafny.Map
			_ = _383_v21
			_383_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), p0)
			var _384_v22 _dafny.Map
			_ = _384_v22
			_384_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_381_v20, _dafny.Companion_Sequence_.Concatenate(p0, (func() _dafny.Sequence {
				if (_383_v21).Contains(_364_i1) {
					return (_383_v21).Get(_364_i1).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("hkctqqp")
			})()))
			_384_v22 = (_384_v22).Update(_381_v20, p0)
			var _385_v23 _dafny.Array
			_ = _385_v23
			var _nw62 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(15))
			_ = _nw62
			_385_v23 = _nw62
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_385_v23), 0))
			_ = _index66
			(_385_v23).ArraySet1(_332_v3, (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_385_v23), 0))
			_ = _index67
			(_385_v23).ArraySet1((func() D1 {
				if !((_this).F32()) || ((_this).F32()) {
					return Companion_Default___.Fm23(globalState)
				}
				return _332_v3
			})(), (_index67).Int())
		}
		(globalState).F14 = (_this).F23()
		var _386_v24 _dafny.Array
		_ = _386_v24
		var _nw63 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(11))
		_ = _nw63
		_386_v24 = _nw63
		for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_386_v24), 0))); ; {
			_guard_loop_1, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _387_i5 _dafny.Int
			_387_i5 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_387_i5).Sign() != -1) && ((_387_i5).Cmp(_dafny.ArrayLen((_386_v24), 0)) < 0)) {
				(_386_v24).ArraySet1(Companion_D1_.Create_DC4_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F23()), (_this).F28(), !((_this).F32())), (_387_i5).Int())
			}
		}
		r0 = (_this).F23()
		r1 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-922))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg32 _dafny.Int) interface{} {
				return coer32(arg32)
			}
		}((func(_388_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_389_i6 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_388_p0).Cardinality())
			}
		})(p0))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(454))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg33 _dafny.Int) interface{} {
				return coer33(arg33)
			}
		}(func(_390_i7 _dafny.Int) _dafny.Int {
			return (_this).F23()
		}))), (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-922))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}((func(_391_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_392_i6 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_391_p0).Cardinality())
			}
		})(p0))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(454))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}(func(_393_i7 _dafny.Int) _dafny.Int {
			return (_this).F23()
		})))).Cardinality()))).Uint32(), (_this).F23())
		return r0, r1
	}
}
func (_this *C2) M6(p0 bool, globalState *GlobalState) (bool, _dafny.Map, D1, _dafny.Array) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var r2 D1 = Companion_D1_.Default()
		_ = r2
		var r3 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r3
		var _394_v0 _dafny.MultiSet
		_ = _394_v0
		_394_v0 = _dafny.MultiSetOf((_this).F28(), (_this).F28(), p0, true)
		var _395_v1 _dafny.Map
		_ = _395_v1
		_395_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_394_v0).Cardinality())
		var _396_v2 _dafny.Sequence
		_ = _396_v2
		_396_v2 = _dafny.UnicodeSeqOfUtf8Bytes("fnnsolgl")
		var _397_v3 D2
		_ = _397_v3
		_397_v3 = Companion_D2_.Create_DC8_((_this).Fm18(_395_v1, _396_v2, (_this).F28(), globalState), (_this).F23(), _dafny.CodePoint('a'), _dafny.IntOfInt64(812), false)
		var _398_v4 D2
		_ = _398_v4
		_398_v4 = Companion_D2_.Create_DC8_((_this).F23(), (_this).F23(), _dafny.CodePoint('t'), (_397_v3).Dtor_cf16(), (_397_v3).Dtor_cf19())
		var _399_v5 _dafny.Set
		_ = _399_v5
		_399_v5 = _dafny.SetOf(_398_v4, _397_v3)
		_399_v5 = _399_v5
		var _400_v6 _dafny.Array
		_ = _400_v6
		var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
		_ = _nw64
		_400_v6 = _nw64
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_400_v6), 0))
		_ = _index68
		(_400_v6).ArraySet1(_dafny.IntOfInt64(528), (_index68).Int())
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_400_v6), 0))
		_ = _index69
		(_400_v6).ArraySet1((_this).F23(), (_index69).Int())
		var _hi8 _dafny.Int = ((_400_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_400_v6), 0))).Int()).(_dafny.Int)).Minus((_400_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_400_v6), 0))).Int()).(_dafny.Int))
		_ = _hi8
		for _401_i0 := (_dafny.IntOfUint32((_this.F29()).Cardinality())).Minus((_this).F23()); _401_i0.Cmp(_hi8) < 0; _401_i0 = _401_i0.Plus(_dafny.One) {
			(globalState).F8 = (_400_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_400_v6), 0))).Int()).(_dafny.Int)
			var _402_v7 _dafny.CodePoint
			_ = _402_v7
			_402_v7 = _dafny.CodePoint('w')
			var _403_v8 _dafny.Map
			_ = _403_v8
			_403_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _402_v7)
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_400_v6), 0))
			_ = _index70
			(_400_v6).ArraySet1((func() _dafny.Int {
				if !((_403_v8).Update((_this).F32(), _402_v7)).Equals(_403_v8) {
					return (func() _dafny.Int {
						if (_395_v1).Contains(p0) {
							return (_395_v1).Get(p0).(_dafny.Int)
						}
						return (_this).F23()
					})()
				}
				return ((_this).F23()).Times(_dafny.IntOfInt64(-826))
			})(), (_index70).Int())
			var _404_v9 *C1
			_ = _404_v9
			var _nw65 *C1 = New_C1_()
			_ = _nw65
			_nw65.Ctor__((_this).F28(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_401_i0), _this.F29()), _401_i0)
			_404_v9 = _nw65
			(_this).F29_set_(_this.F29())
		}
		var _405_v10 _dafny.Sequence
		_ = _405_v10
		_405_v10 = _dafny.SeqOf(p0, (_this).F28())
		var _406_v11 D3
		_ = _406_v11
		_406_v11 = Companion_D3_.Create_DC12_(_405_v10)
		_406_v11 = (func() D3 {
			if (p0) == (p0) {
				return _406_v11
			}
			return _406_v11
		})()
		var _407_v12 _dafny.Array
		_ = _407_v12
		var _nw66 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
		_ = _nw66
		_407_v12 = _nw66
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_407_v12), 0))
		_ = _index71
		(_407_v12).ArraySet1(p0, (_index71).Int())
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_407_v12), 0))
		_ = _index72
		(_407_v12).ArraySet1((_this).F32(), (_index72).Int())
		var _408_i1 _dafny.Int
		_ = _408_i1
		_408_i1 = _dafny.Zero
		{
			for (_407_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_407_v12), 0))).Int()).(bool) {
				{
					if (_408_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_408_i1 = (_408_i1).Plus(_dafny.One)
					r0 = !((_407_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_407_v12), 0))).Int()).(bool))
					var _409_v13 _dafny.Sequence
					_ = _409_v13
					_409_v13 = _dafny.SeqOf(_396_v2, _396_v2)
					(globalState).F21 = !_dafny.Companion_Sequence_.Contains(_409_v13, _396_v2)
					var _410_v14 _dafny.Sequence
					_ = _410_v14
					_410_v14 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_395_v1).Cardinality()))
					(globalState).F14 = (_this).Fm18((_410_v14).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_410_v14).Cardinality()))).Uint32()).(_dafny.Map), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(339))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg36 _dafny.Int) interface{} {
							return coer36(arg36)
						}
					}(func(_411_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('h')
					})), _396_v2), (_407_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_407_v12), 0))).Int()).(bool), globalState)
					_395_v1 = (_395_v1).Update((_this).F32(), (_this).F23())
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		r0 = (_this).F28()
		var _412_v15 _dafny.Map
		_ = _412_v15
		_412_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F23()).Times((_this).F23()), _407_v12)
		r1 = _412_v15
		var _413_v16 D1
		_ = _413_v16
		_413_v16 = Companion_D1_.Create_DC4_((_this).F31(), (_407_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_407_v12), 0))).Int()).(bool), (_407_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_407_v12), 0))).Int()).(bool))
		r2 = _413_v16
		r3 = _407_v12
		return r0, r1, r2, r3
	}
}
func (_this *C2) F31() _dafny.Map {
	{
		return _this._f31
	}
}
func (_this *C2) F32() bool {
	{
		return _this._f32
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f29 _dafny.Sequence
	_f23 _dafny.Int
	_f28 bool
	_f30 _dafny.Sequence
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f29 = _dafny.EmptySeq
	_this._f23 = _dafny.Zero
	_this._f28 = false
	_this._f30 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C3{}
var _ T1 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F29() _dafny.Sequence {
	return _this._f29
}
func (_this *C3) F29_set_(value _dafny.Sequence) {
	_this._f29 = value
}
func (_this *C3) F23() _dafny.Int {
	return _this._f23
}
func (_this *C3) F28() bool {
	return _this._f28
}
func (_this *C3) Ctor__(f30 _dafny.Sequence, f23 _dafny.Int, f28 bool, f29 _dafny.Sequence) {
	{
		(_this)._f30 = f30
		(_this)._f23 = f23
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C3) Fm12(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(((_this).F23()).Plus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-585))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg37 _dafny.Int) interface{} {
				return coer37(arg37)
			}
		}(func(_414_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		}))).Cardinality()), _dafny.IntOfInt64(974))))
	}
}
func (_this *C3) M1(p0 _dafny.Array, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Set) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Set = _dafny.EmptySet
		_ = r1
		var _415_v0 _dafny.Array
		_ = _415_v0
		var _nw67 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
		_ = _nw67
		_415_v0 = _nw67
		for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_415_v0), 0))); ; {
			_guard_loop_2, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _416_i0 _dafny.Int
			_416_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_416_i0).Sign() != -1) && ((_416_i0).Cmp(_dafny.ArrayLen((_415_v0), 0)) < 0)) {
				(_415_v0).ArraySet1((Companion_D2_.Create_DC8_((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F23())))).Cardinality(), (_this).F23(), _dafny.CodePoint('p'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nasts")).Cardinality()), true)).Dtor_cf19(), (_416_i0).Int())
			}
		}
		if (Companion_Default___.SafeDivisionInt((_this).F23(), (_this).F23())).Cmp(((_this).F23()).Times((_this).F23())) < 0 {
			var _417_v1 _dafny.Set
			_ = _417_v1
			_417_v1 = _dafny.SetOf(p1, p1)
			var _rhs87 bool = (_417_v1).IsProperSubsetOf((_417_v1).Difference(_dafny.SetOf((_this).F28())))
			_ = _rhs87
			var _rhs88 bool = (_this).F28()
			_ = _rhs88
			var _lhs84 *GlobalState = globalState
			_ = _lhs84
			var _lhs85 *GlobalState = globalState
			_ = _lhs85
			_lhs84.F21 = _rhs87
			_lhs85.F21 = _rhs88
			var _418_v2 _dafny.Map
			_ = _418_v2
			_418_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), false)
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((p0), 0))
			_ = _index73
			(p0).ArraySet1((func() bool {
				if (_418_v2).Contains(p1) {
					return (_418_v2).Get(p1).(bool)
				}
				return p1
			})(), (_index73).Int())
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((p0), 0))
			_ = _index74
			(p0).ArraySet1((_this).F28(), (_index74).Int())
			var _419_v3 _dafny.Array
			_ = _419_v3
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
			_ = _nw68
			_419_v3 = _nw68
			var _420_v4 _dafny.Map
			_ = _420_v4
			_420_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), true)
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_419_v3), 0))
			_ = _index75
			(_419_v3).ArraySet1(_420_v4, (_index75).Int())
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_419_v3), 0))
			_ = _index76
			(_419_v3).ArraySet1(_420_v4, (_index76).Int())
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((p0), 0))
			_ = _index77
			(p0).ArraySet1(!(_418_v2).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((p0), 0))).Int()).(bool)), (_index77).Int())
			var _421_v5 D1
			_ = _421_v5
			_421_v5 = Companion_D1_.Create_DC5_((_this).F23(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((p0), 0))).Int()).(bool), (_this).F23(), (_this).F23())
			var _422_v6 _dafny.Sequence
			_ = _422_v6
			_422_v6 = _dafny.SeqOf(_421_v5)
			var _423_v7 _dafny.Map
			_ = _423_v7
			_423_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v6, p1)
			var _424_v8 _dafny.Map
			_ = _424_v8
			_424_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-225), (_this).F23()))
			var _425_v9 _dafny.MultiSet
			_ = _425_v9
			_425_v9 = _dafny.MultiSetOf((_this).F28())
			var _426_v10 D3
			_ = _426_v10
			_426_v10 = Companion_D3_.Create_DC11_(Companion_Default___.Fm0((_425_v9).Cardinality(), (_this).F28(), globalState))
			var _427_v11 _dafny.Map
			_ = _427_v11
			_427_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_425_v9).Cardinality(), (_426_v10).Dtor_cf22())
			var _428_v12 _dafny.Set
			_ = _428_v12
			_428_v12 = _dafny.SetOf((_this).F23())
			var _429_v13 D1
			_ = _429_v13
			_429_v13 = Companion_D1_.Create_DC4_(_427_v11, ((_this).F28()) || (p1), ((_428_v12).Cardinality()).Cmp((_this).F23()) != 0)
			var _rhs89 _dafny.Map = _423_v7
			_ = _rhs89
			var _rhs90 _dafny.Int = (_this).F23()
			_ = _rhs90
			var _rhs91 _dafny.Map = Companion_Default___.Fm13(_dafny.CodePoint('q'), Companion_Default___.Fm1((_this).F23(), (_this).F28(), globalState), (func() _dafny.Int {
				if Companion_Default___.Fm1((_this).F23(), p1, globalState) {
					return (_this).F23()
				}
				return (_this).F23()
			})(), (func() bool {
				if p1 {
					return (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((p0), 0))).Int()).(bool)
				}
				return Companion_Default___.Fm1((_this).F23(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState)
			})(), globalState)
			_ = _rhs91
			var _rhs92 D1 = _429_v13
			_ = _rhs92
			var _rhs93 _dafny.Int = _dafny.IntOfInt64(473)
			_ = _rhs93
			var _lhs86 *GlobalState = globalState
			_ = _lhs86
			var _lhs87 *GlobalState = globalState
			_ = _lhs87
			_423_v7 = _rhs89
			_lhs86.F14 = _rhs90
			_424_v8 = _rhs91
			_429_v13 = _rhs92
			_lhs87.F14 = _rhs93
		} else {
			(globalState).F14 = (_this).F23()
			var _430_v14 _dafny.Set
			_ = _430_v14
			_430_v14 = _dafny.SetOf(p1)
			var _431_v15 _dafny.CodePoint
			_ = _431_v15
			_431_v15 = _dafny.CodePoint('h')
			var _432_v16 _dafny.Map
			_ = _432_v16
			_432_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F23())
			var _433_v17 _dafny.MultiSet
			_ = _433_v17
			_433_v17 = _dafny.MultiSetOf((_this).F23(), (_this).Fm12((_this).F23(), (_this).F28(), p1, globalState), (_this).F23())
			var _434_v18 _dafny.Sequence
			_ = _434_v18
			_434_v18 = _dafny.UnicodeSeqOfUtf8Bytes("htcv")
			var _435_v19 _dafny.Array
			_ = _435_v19
			var _nwElement0_12 _dafny.Int = (_this).F23()
			_ = _nwElement0_12
			var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(28))
			_ = _nw69
			(_nw69).ArraySet1(_nwElement0_12, 0)
			(_nw69).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F23(), (_this).F23()), 1)
			(_nw69).ArraySet1((_this).F23(), 2)
			(_nw69).ArraySet1((_this).F23(), 3)
			(_nw69).ArraySet1((_this).F23(), 4)
			(_nw69).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("n")).Cardinality()), 5)
			(_nw69).ArraySet1(((_430_v14).Cardinality()).Times((_this).F23()), 6)
			(_nw69).ArraySet1((_this).F23(), 7)
			(_nw69).ArraySet1((Companion_Default___.Fm14((_this).F23(), _431_v15, (_this).F28(), globalState)).Cardinality(), 8)
			(_nw69).ArraySet1((_this).F23(), 9)
			(_nw69).ArraySet1((_this).F23(), 10)
			(_nw69).ArraySet1((_this).F23(), 11)
			(_nw69).ArraySet1(((_this).Fm12((_this).F23(), Companion_Default___.Fm1((_this).F23(), p1, globalState), !(p1), globalState)).Plus(_dafny.IntOfInt64(447)), 12)
			(_nw69).ArraySet1((_this).F23(), 13)
			(_nw69).ArraySet1(_dafny.IntOfInt64(629), 14)
			(_nw69).ArraySet1((_this).Fm12(Companion_Default___.Fm0((_this).F23(), p1, globalState), (_this).F28(), (_this).F28(), globalState), 15)
			(_nw69).ArraySet1((_this).F23(), 16)
			(_nw69).ArraySet1(((_432_v16).Cardinality()).Times((func() _dafny.Int {
				if (_433_v17).Contains(_dafny.IntOfInt64(876)) {
					return (_433_v17).Multiplicity(_dafny.IntOfInt64(876))
				}
				return (_this).F23()
			})()), 17)
			(_nw69).ArraySet1(((_this).F23()).Plus((_this).Fm12(_dafny.IntOfUint32((_434_v18).Cardinality()), (_this).F28(), p1, globalState)), 18)
			(_nw69).ArraySet1((_dafny.IntOfInt64(547)).Plus((_this).F23()), 19)
			(_nw69).ArraySet1(_dafny.IntOfUint32((_434_v18).Cardinality()), 20)
			(_nw69).ArraySet1((func() _dafny.Int {
				if p1 {
					return _dafny.IntOfUint32((_this.F29()).Cardinality())
				}
				return (_this).F23()
			})(), 21)
			(_nw69).ArraySet1((_this).F23(), 22)
			(_nw69).ArraySet1((_this).F23(), 23)
			(_nw69).ArraySet1(Companion_Default___.SafeModuloInt((_this).F23(), (_this).F23()), 24)
			(_nw69).ArraySet1(((_433_v17).Update((_this).F23(), Companion_Default___.Abs((_this).Fm12((_this).F23(), (_this).F28(), (_this).F28(), globalState)))).Cardinality(), 25)
			(_nw69).ArraySet1(_dafny.IntOfInt64(966), 26)
			(_nw69).ArraySet1((_this).F23(), 27)
			_435_v19 = _nw69
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))
			_ = _index78
			(_435_v19).ArraySet1((_this).F23(), (_index78).Int())
			var _436_v20 _dafny.Map
			_ = _436_v20
			_436_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), ((_this).F23()).Cmp(_dafny.IntOfInt64(141)) <= 0)
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))
			_ = _index79
			(_435_v19).ArraySet1((_dafny.Zero).Minus((_436_v20).Cardinality()), (_index79).Int())
			(globalState).F8 = Companion_Default___.Fm0((_this).F23(), p1, globalState)
			if (func() bool {
				if (_this).F28() {
					return (_this).F28()
				}
				return p1
			})() {
				var _437_v21 _dafny.Sequence
				_ = _437_v21
				_437_v21 = _dafny.SeqOf(false, p1, (_this).F28())
				var _438_v22 _dafny.MultiSet
				_ = _438_v22
				_438_v22 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_437_v21, _437_v21))
				var _439_v23 _dafny.MultiSet
				_ = _439_v23
				_439_v23 = _dafny.MultiSetOf(p1)
				var _440_v24 _dafny.Map
				_ = _440_v24
				_440_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int), _439_v23)
				var _pat_let_tv26 = p1
				_ = _pat_let_tv26
				_438_v22 = Companion_Default___.Fm15((func(_pat_let0_0 D3) D3 {
					return func(_441_dt__update__tmp_h0 D3) D3 {
						return func(_pat_let1_0 bool) D3 {
							return func(_442_dt__update_hcf21_h0 bool) D3 {
								return Companion_D3_.Create_DC10_(_442_dt__update_hcf21_h0)
							}(_pat_let1_0)
						}(_pat_let_tv26)
					}(_pat_let0_0)
				}(Companion_D3_.Create_DC10_(p1))).Dtor_cf21(), _440_v24, globalState)
				var _443_v25 _dafny.Sequence
				_ = _443_v25
				_443_v25 = _dafny.SeqOf(_415_v0, _415_v0)
				var _444_v26 _dafny.Array
				_ = _444_v26
				var _nwElement0_13 _dafny.Array = _415_v0
				_ = _nwElement0_13
				var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(14))
				_ = _nw70
				(_nw70).ArraySet1(_nwElement0_13, 0)
				(_nw70).ArraySet1(p0, 1)
				(_nw70).ArraySet1(p0, 2)
				(_nw70).ArraySet1(_415_v0, 3)
				(_nw70).ArraySet1((func() _dafny.Array {
					if (_this).F28() {
						return p0
					}
					return _415_v0
				})(), 4)
				(_nw70).ArraySet1(_415_v0, 5)
				(_nw70).ArraySet1((_443_v25).Select((Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_this.F29())).Cardinality(), _dafny.IntOfUint32((_443_v25).Cardinality()))).Uint32()).(_dafny.Array), 6)
				(_nw70).ArraySet1(_415_v0, 7)
				(_nw70).ArraySet1((func() _dafny.Array {
					if p1 {
						return p0
					}
					return _415_v0
				})(), 8)
				(_nw70).ArraySet1(p0, 9)
				(_nw70).ArraySet1(p0, 10)
				(_nw70).ArraySet1(_415_v0, 11)
				(_nw70).ArraySet1(p0, 12)
				(_nw70).ArraySet1(p0, 13)
				_444_v26 = _nw70
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_444_v26), 0))
				_ = _index80
				(_444_v26).ArraySet1(p0, (_index80).Int())
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_444_v26), 0))
				_ = _index81
				(_444_v26).ArraySet1(p0, (_index81).Int())
				_430_v14 = _430_v14
				var _445_v27 _dafny.Array
				_ = _445_v27
				var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
				_ = _nw71
				_445_v27 = _nw71
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_445_v27), 0))
				_ = _index82
				(_445_v27).ArraySet1(_435_v19, (_index82).Int())
				var _446_v28 _dafny.Map
				_ = _446_v28
				_446_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int), p1)
				var _447_v29 _dafny.Map
				_ = _447_v29
				_447_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_434_v18, (func() _dafny.Map {
					if p1 {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_437_v21).Select((Companion_Default___.SafeIndex((_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_437_v21).Cardinality()))).Uint32()).(bool), (_439_v23).Cardinality())
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
						if (_446_v28).Contains(_dafny.IntOfUint32((_434_v18).Cardinality())) {
							return (_446_v28).Get(_dafny.IntOfUint32((_434_v18).Cardinality())).(bool)
						}
						return (_this).F28()
					})(), (_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int))
				})())
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_445_v27), 0))
				_ = _index83
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))
				_ = _index84
				var _rhs94 _dafny.Map = (func() _dafny.Map {
					if (_447_v29).Contains(_434_v18) {
						return (_447_v29).Get(_434_v18).(_dafny.Map)
					}
					return _432_v16
				})()
				_ = _rhs94
				var _rhs95 _dafny.Array = _435_v19
				_ = _rhs95
				var _rhs96 _dafny.Int = (_this).Fm12((_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int), false, p1, globalState)
				_ = _rhs96
				var _rhs97 bool = (_this).F28()
				_ = _rhs97
				var _rhs98 _dafny.Int = ((_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int)).Minus((_this).F23())
				_ = _rhs98
				var _lhs88 _dafny.Array = _445_v27
				_ = _lhs88
				var _lhs89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_445_v27), 0))
				_ = _lhs89
				var _lhs90 _dafny.Array = _435_v19
				_ = _lhs90
				var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))
				_ = _lhs91
				var _lhs92 *GlobalState = globalState
				_ = _lhs92
				var _lhs93 *GlobalState = globalState
				_ = _lhs93
				_432_v16 = _rhs94
				(_lhs88).ArraySet1(_rhs95, (_lhs89).Int())
				(_lhs90).ArraySet1(_rhs96, (_lhs91).Int())
				_lhs92.F9 = _rhs97
				_lhs93.F14 = _rhs98
				(globalState).F8 = (_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int)
			} else {
				var _448_v30 _dafny.Map
				_ = _448_v30
				_448_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_415_v0, ((_this).F23()).Times(_dafny.IntOfInt64(434)))
				_448_v30 = (_448_v30).Update(_415_v0, (_this).F23())
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))
				_ = _index85
				(_435_v19).ArraySet1((_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int), (_index85).Int())
				var _449_v31 _dafny.Sequence
				_ = _449_v31
				_449_v31 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_this.F29(), _this.F29()), _this.F29())
				var _450_v32 _dafny.Map
				_ = _450_v32
				_450_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(391), _this.F29())
				_449_v31 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_449_v31, (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_433_v17).Contains((_this).F23()) {
						return (_433_v17).Multiplicity((_this).F23())
					}
					return (_this).F23()
				})(), _dafny.IntOfUint32((_449_v31).Cardinality()))).Uint32(), _this.F29()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_449_v31, (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_433_v17).Contains((_this).F23()) {
						return (_433_v17).Multiplicity((_this).F23())
					}
					return (_this).F23()
				})(), _dafny.IntOfUint32((_449_v31).Cardinality()))).Uint32(), _this.F29())).Cardinality()))).Uint32(), (func() _dafny.Sequence {
					if (_450_v32).Contains((_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int)) {
						return (_450_v32).Get((_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
					}
					return _dafny.SeqOf((_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int))
				})()), _dafny.Companion_Sequence_.Concatenate(_449_v31, _449_v31))
				(_this).M5(globalState)
				var _451_v33 D4
				_ = _451_v33
				_451_v33 = Companion_D4_.Create_DC14_(_435_v19)
				_435_v19 = (_451_v33).Dtor_cf25()
			}
			var _452_v34 _dafny.Map
			_ = _452_v34
			_452_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC1_((_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int), (_this).F28(), (_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int))).Dtor_cf1(), (_this).F23())
			var _453_v35 _dafny.Sequence
			_ = _453_v35
			_453_v35 = _dafny.SeqOf(Companion_Default___.Fm16((_this).F28(), (_this).F23(), (_this).F23(), globalState))
			_452_v34 = (_452_v34).Update((_this).F23(), (((_453_v35).Select((Companion_Default___.SafeIndex((_435_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_435_v19), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_453_v35).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_434_v18).Cardinality()), (_this).F23()))).Cardinality())
		}
		(globalState).F14 = (_this).F23()
		var _454_v36 _dafny.Sequence
		_ = _454_v36
		_454_v36 = _dafny.UnicodeSeqOfUtf8Bytes("tjvi")
		var _455_v37 _dafny.CodePoint
		_ = _455_v37
		_455_v37 = _dafny.CodePoint('y')
		if !(_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("ii"), _dafny.Companion_Sequence_.Update(_454_v36, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_454_v36).Cardinality()))).Uint32(), _455_v37))) {
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_415_v0), 0))
			_ = _index86
			(_415_v0).ArraySet1((_this).F28(), (_index86).Int())
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_415_v0), 0))
			_ = _index87
			(_415_v0).ArraySet1(!((Companion_D3_.Create_DC10_((_this).F28())).Dtor_cf21()), (_index87).Int())
			var _456_v38 _dafny.Array
			_ = _456_v38
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_5
			var _nw72 _dafny.Array
			_ = _nw72
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw72 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.MultiSet = (func(_457_v0 _dafny.Array) func(_dafny.Int) _dafny.MultiSet {
					return func(_458_i1 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf((_457_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_457_v0), 0))).Int()).(bool))
					}
				})(_415_v0)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw72 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw72).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw72).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_456_v38 = _nw72
			_456_v38 = _456_v38
			var _459_v40 D1
			_ = _459_v40
			_459_v40 = Companion_D1_.Create_DC4_(func() _dafny.Map {
				var _coll18 = _dafny.NewMapBuilder()
				_ = _coll18
				for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(57), _dafny.IntOfInt64(984))); ; {
					_compr_18, _ok21 := _iter21()
					if !_ok21 {
						break
					}
					var _460_v39 _dafny.Int
					_460_v39 = interface{}(_compr_18).(_dafny.Int)
					if ((_dafny.IntOfInt64(57)).Cmp(_460_v39) <= 0) && ((_460_v39).Cmp(_dafny.IntOfInt64(984)) < 0) {
						_coll18.Add((_460_v39).Plus((_this).F23()), (_this).F23())
					}
				}
				return _coll18.ToMap()
			}(), (_415_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_415_v0), 0))).Int()).(bool), p1)
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_415_v0), 0))
			_ = _index88
			(_415_v0).ArraySet1((Companion_Default___.Fm1((_this).F23(), false, globalState)) == ((_459_v40).Dtor_cf7()), (_index88).Int())
			(globalState).F6 = _dafny.Companion_Sequence_.Concatenate(_454_v36, _dafny.UnicodeSeqOfUtf8Bytes("l"))
			(globalState).F8 = (_this.F29()).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32()).(_dafny.Int)
		} else {
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_415_v0), 0))
			_ = _index89
			(_415_v0).ArraySet1(p1, (_index89).Int())
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_415_v0), 0))
			_ = _index90
			(_415_v0).ArraySet1(p1, (_index90).Int())
			var _461_v41 _dafny.Array
			_ = _461_v41
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_6
			var _nw73 _dafny.Array
			_ = _nw73
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw73 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Int = func(_462_i2 _dafny.Int) _dafny.Int {
					return (_462_i2).Plus((_this).F23())
				}
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw73 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw73).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw73).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_461_v41 = _nw73
			var _463_v42 _dafny.Sequence
			_ = _463_v42
			_463_v42 = _dafny.SeqOf((_415_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_415_v0), 0))).Int()).(bool), p1)
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_461_v41), 0))
			_ = _index91
			(_461_v41).ArraySet1(_dafny.IntOfUint32((_463_v42).Cardinality()), (_index91).Int())
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_461_v41), 0))
			_ = _index92
			(_461_v41).ArraySet1((_this).Fm12((_this).F23(), (_415_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_415_v0), 0))).Int()).(bool), (_415_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_415_v0), 0))).Int()).(bool), globalState), (_index92).Int())
			var _464_v43 _dafny.Array
			_ = _464_v43
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_7
			var _nw74 _dafny.Array
			_ = _nw74
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw74 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Sequence = (func(_465_v36 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_466_i3 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_465_v36, _465_v36)
					}
				})(_454_v36)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw74 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw74).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw74).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_464_v43 = _nw74
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_464_v43), 0))
			_ = _index93
			(_464_v43).ArraySet1(_454_v36, (_index93).Int())
			var _467_v44 _dafny.MultiSet
			_ = _467_v44
			_467_v44 = _dafny.MultiSetOf(_455_v37)
			var _468_v45 _dafny.Map
			_ = _468_v45
			_468_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_461_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_461_v41), 0))).Int()).(_dafny.Int))
			var _469_v46 _dafny.Sequence
			_ = _469_v46
			_469_v46 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("qkbs"), _dafny.UnicodeSeqOfUtf8Bytes("uxecrdjl"))
			var _470_v47 _dafny.Set
			_ = _470_v47
			_470_v47 = _dafny.SetOf(true, (_this).F28())
			var _471_v48 _dafny.MultiSet
			_ = _471_v48
			_471_v48 = _dafny.MultiSetOf(_470_v47)
			var _472_v49 D4
			_ = _472_v49
			_472_v49 = Companion_D4_.Create_DC15_((_471_v48).Cardinality(), _455_v37, (_this).F23())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_461_v41), 0))
			_ = _index94
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_461_v41), 0))
			_ = _index95
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_464_v43), 0))
			_ = _index96
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_461_v41), 0))
			_ = _index97
			var _rhs99 _dafny.Int = (_461_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_461_v41), 0))).Int()).(_dafny.Int)
			_ = _rhs99
			var _rhs100 _dafny.Int = Companion_Default___.Fm0((_this).F23(), (Companion_Default___.Fm17((_this).F28(), _dafny.IntOfUint32((_454_v36).Cardinality()), globalState)).IsProperSubsetOf((_467_v44).Update(_455_v37, Companion_Default___.Abs((_468_v45).Cardinality()))), globalState)
			_ = _rhs100
			var _rhs101 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_454_v36, _454_v36)
			_ = _rhs101
			var _rhs102 _dafny.Int = (func() _dafny.Int {
				if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(39))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}((func(_473_v36 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_474_i4 _dafny.Int) _dafny.Sequence {
						return _473_v36
					}
				})(_454_v36))), _469_v46) {
					return (func() _dafny.Int {
						if (_415_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_415_v0), 0))).Int()).(bool) {
							return (_472_v49).Dtor_cf28()
						}
						return (_this).Fm12((_this).F23(), (_this).F28(), Companion_Default___.Fm1((_this).F23(), p1, globalState), globalState)
					})()
				}
				return ((_461_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_461_v41), 0))).Int()).(_dafny.Int)).Minus((_this).F23())
			})()
			_ = _rhs102
			var _lhs94 _dafny.Array = _461_v41
			_ = _lhs94
			var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_461_v41), 0))
			_ = _lhs95
			var _lhs96 _dafny.Array = _461_v41
			_ = _lhs96
			var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_461_v41), 0))
			_ = _lhs97
			var _lhs98 _dafny.Array = _464_v43
			_ = _lhs98
			var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_464_v43), 0))
			_ = _lhs99
			var _lhs100 _dafny.Array = _461_v41
			_ = _lhs100
			var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_461_v41), 0))
			_ = _lhs101
			(_lhs94).ArraySet1(_rhs99, (_lhs95).Int())
			(_lhs96).ArraySet1(_rhs100, (_lhs97).Int())
			(_lhs98).ArraySet1(_rhs101, (_lhs99).Int())
			(_lhs100).ArraySet1(_rhs102, (_lhs101).Int())
			(_this).M5(globalState)
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_415_v0), 0))
			_ = _index98
			(_415_v0).ArraySet1(((_461_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_461_v41), 0))).Int()).(_dafny.Int)).Cmp((_468_v45).Cardinality()) == 0, (_index98).Int())
		}
		var _hi9 _dafny.Int = _dafny.IntOfInt64(581)
		_ = _hi9
		for _475_i5 := (_dafny.Zero).Minus((_this).F23()); _475_i5.Cmp(_hi9) < 0; _475_i5 = _475_i5.Plus(_dafny.One) {
			var _476_v50 _dafny.Map
			_ = _476_v50
			_476_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !(p1))
			var _477_v51 _dafny.Map
			_ = _477_v51
			_477_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_476_v50).Cardinality(), p0)
			var _478_v52 _dafny.Array
			_ = _478_v52
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_8
			var _nw75 _dafny.Array
			_ = _nw75
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw75 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Int = func(_479_i6 _dafny.Int) _dafny.Int {
					return (_479_i6).Plus((_dafny.SetOf(_this.F29())).Cardinality())
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw75 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw75).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw75).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_478_v52 = _nw75
			var _480_v53 _dafny.Map
			_ = _480_v53
			_480_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_477_v51).Contains((_this).F23()) {
					return (_477_v51).Get((_this).F23()).(_dafny.Array)
				}
				return p0
			})(), _478_v52)
			var _481_v54 D4
			_ = _481_v54
			_481_v54 = Companion_D4_.Create_DC14_((func() _dafny.Array {
				if (_480_v53).Contains(p0) {
					return (_480_v53).Get(p0).(_dafny.Array)
				}
				return _478_v52
			})())
			_481_v54 = Companion_D4_.Create_DC14_(_478_v52)
			var _482_v55 _dafny.Map
			_ = _482_v55
			_482_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F23())
			var _483_v56 D1
			_ = _483_v56
			_483_v56 = Companion_D1_.Create_DC4_(_482_v55, p1, p1)
			var _source3 D1 = _483_v56
			_ = _source3
			if _source3.Is_DC4() {
				var _484___mcc_h0 _dafny.Map = _source3.Get_().(D1_DC4).Cf6
				_ = _484___mcc_h0
				var _485___mcc_h1 bool = _source3.Get_().(D1_DC4).Cf7
				_ = _485___mcc_h1
				var _486___mcc_h2 bool = _source3.Get_().(D1_DC4).Cf8
				_ = _486___mcc_h2
				var _487_cf8 bool = _486___mcc_h2
				_ = _487_cf8
				var _488_cf7 bool = _485___mcc_h1
				_ = _488_cf7
				var _489_cf6 _dafny.Map = _484___mcc_h0
				_ = _489_cf6
				(globalState).F14 = (_dafny.IntOfUint32((_454_v36).Cardinality())).Times(((_this).F23()).Times((_this).F23()))
				var _490_v57 _dafny.Map
				_ = _490_v57
				_490_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F23())
				var _491_v58 _dafny.MultiSet
				_ = _491_v58
				_491_v58 = _dafny.MultiSetOf((_490_v57).Cardinality())
				var _rhs103 bool = (_491_v58).IsSubsetOf(_491_v58)
				_ = _rhs103
				var _rhs104 D4 = _481_v54
				_ = _rhs104
				var _rhs105 _dafny.CodePoint = _455_v37
				_ = _rhs105
				var _rhs106 _dafny.Int = (_475_i5).Plus((_this).F23())
				_ = _rhs106
				var _rhs107 D4 = _481_v54
				_ = _rhs107
				var _lhs102 *GlobalState = globalState
				_ = _lhs102
				_488_cf7 = _rhs103
				_481_v54 = _rhs104
				_455_v37 = _rhs105
				_lhs102.F14 = _rhs106
				_481_v54 = _rhs107
				var _492_v59 _dafny.Map
				_ = _492_v59
				_492_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(Companion_Default___.Fm1(_dafny.IntOfInt64(631), _487_cf8, globalState)) || (!(_487_cf8)), _476_v50)
				var _493_v60 _dafny.Sequence
				_ = _493_v60
				_493_v60 = _dafny.SeqOf(_476_v50)
				_492_v59 = (_492_v59).Update((_this).F28(), ((_476_v50).Update(p1, _487_cf8)).Merge((_493_v60).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_493_v60).Cardinality()))).Uint32()).(_dafny.Map)))
				var _494_v61 _dafny.Map
				_ = _494_v61
				_494_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_478_v52, _dafny.SetOf(_455_v37, _455_v37)), (_dafny.Zero).Minus((_this).F23()))
				var _495_v62 _dafny.Set
				_ = _495_v62
				_495_v62 = _dafny.SetOf(_455_v37, _dafny.CodePoint('e'), _dafny.CodePoint('a'), _455_v37)
				var _496_v63 _dafny.Map
				_ = _496_v63
				_496_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_478_v52, _495_v62)
				_494_v61 = (_494_v61).Update((_496_v63).Merge(_496_v63), _dafny.IntOfInt64(568))
			} else if _source3.Is_DC5() {
				var _497___mcc_h3 _dafny.Int = _source3.Get_().(D1_DC5).Cf9
				_ = _497___mcc_h3
				var _498___mcc_h4 bool = _source3.Get_().(D1_DC5).Cf10
				_ = _498___mcc_h4
				var _499___mcc_h5 _dafny.Int = _source3.Get_().(D1_DC5).Cf11
				_ = _499___mcc_h5
				var _500___mcc_h6 _dafny.Int = _source3.Get_().(D1_DC5).Cf12
				_ = _500___mcc_h6
				var _501_cf12 _dafny.Int = _500___mcc_h6
				_ = _501_cf12
				var _502_cf11 _dafny.Int = _499___mcc_h5
				_ = _502_cf11
				var _503_cf10 bool = _498___mcc_h4
				_ = _503_cf10
				var _504_cf9 _dafny.Int = _497___mcc_h3
				_ = _504_cf9
				var _505_v64 _dafny.Array
				_ = _505_v64
				var _nw76 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
				_ = _nw76
				_505_v64 = _nw76
				var _506_v65 T1
				_ = _506_v65
				var _nw77 *C2 = New_C2_()
				_ = _nw77
				_nw77.Ctor__(_482_v55, false, _503_cf10, _dafny.SeqOf(_dafny.IntOfInt64(970)), (_this.F29()).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F23()), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32()).(_dafny.Int))
				_506_v65 = _nw77
				var _507_v66 T1
				_ = _507_v66
				_507_v66 = _506_v65
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_505_v64), 0))
				_ = _index99
				(_505_v64).ArraySet1((_507_v66), (_index99).Int())
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_505_v64), 0))
				_ = _index100
				(_505_v64).ArraySet1(_506_v65, (_index100).Int())
				var _508_v67 _dafny.Sequence
				_ = _508_v67
				_508_v67 = _dafny.SeqOf((_this).F28(), p1, _503_cf10, (_this).F28())
				_502_cf11 = (_502_cf11).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_508_v67).Cardinality()), (_this).F23()))
				(globalState).F14 = _501_cf12
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_415_v0), 0))
				_ = _index101
				(_415_v0).ArraySet1(!(true), (_index101).Int())
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_415_v0), 0))
				_ = _index102
				(_415_v0).ArraySet1(_503_cf10, (_index102).Int())
			} else if _source3.Is_DC6() {
				var _509___mcc_h7 _dafny.Array = _source3.Get_().(D1_DC6).Cf13
				_ = _509___mcc_h7
				var _510_cf13 _dafny.Array = _509___mcc_h7
				_ = _510_cf13
				var _511_v68 *C1
				_ = _511_v68
				var _nw78 *C1 = New_C1_()
				_ = _nw78
				_nw78.Ctor__((_this).F28(), _this.F29(), Companion_Default___.Fm0((_this).F23(), (_this).F28(), globalState))
				_511_v68 = _nw78
				_511_v68 = _511_v68
				_476_v50 = (_476_v50).Update(p1, !(p1))
				var _512_v69 _dafny.Map
				_ = _512_v69
				_512_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _455_v37)
				_512_v69 = (_512_v69).Update((_this).F28(), _455_v37)
				var _513_v70 _dafny.Map
				_ = _513_v70
				_513_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _476_v50)
				(globalState).F9 = (_513_v70).Contains((_this).F28())
			} else {
				var _514___mcc_h8 _dafny.Sequence = _source3.Get_().(D1_DC3).Cf5
				_ = _514___mcc_h8
				var _515_cf5 _dafny.Sequence = _514___mcc_h8
				_ = _515_cf5
				var _rhs108 bool = p1
				_ = _rhs108
				var _rhs109 _dafny.Int = _dafny.IntOfInt64(408)
				_ = _rhs109
				var _lhs103 *GlobalState = globalState
				_ = _lhs103
				_lhs103.F7 = _rhs108
				r0 = _rhs109
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_478_v52), 0))
				_ = _index103
				(_478_v52).ArraySet1((_this).F23(), (_index103).Int())
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_478_v52), 0))
				_ = _index104
				(_478_v52).ArraySet1((_this).Fm12((_this).F23(), p1, (_this).F28(), globalState), (_index104).Int())
				var _516_v71 D0
				_ = _516_v71
				var _517_v72 _dafny.Map
				_ = _517_v72
				var _518_v73 _dafny.Int
				_ = _518_v73
				var _out24 D0
				_ = _out24
				var _out25 _dafny.Map
				_ = _out25
				var _out26 _dafny.Int
				_ = _out26
				_out24, _out25, _out26 = Companion_Default___.M0((func() _dafny.Array {
					if (_this).F28() {
						return p0
					}
					return _415_v0
				})(), _dafny.UnicodeSeqOfUtf8Bytes("nplonay"), _475_i5, globalState)
				_516_v71 = _out24
				_517_v72 = _out25
				_518_v73 = _out26
				(globalState).F14 = (_475_i5).Plus(_518_v73)
			}
			(globalState).F14 = (_475_i5).Minus(_dafny.IntOfUint32((_454_v36).Cardinality()))
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_415_v0), 0))
			_ = _index105
			(_415_v0).ArraySet1(p1, (_index105).Int())
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_415_v0), 0))
			_ = _index106
			(_415_v0).ArraySet1(Companion_Default___.Fm1((_this).F23(), !(p1), globalState), (_index106).Int())
		}
		var _519_v74 _dafny.MultiSet
		_ = _519_v74
		_519_v74 = _dafny.MultiSetOf(!(true))
		var _520_v75 _dafny.Set
		_ = _520_v75
		_520_v75 = _dafny.SetOf((_this).F23(), Companion_Default___.Fm0((_this).F23(), (_this).F28(), globalState))
		if ((_dafny.Zero).Minus((_dafny.Zero).Minus(((_519_v74).Cardinality()).Minus((_520_v75).Cardinality())))).Cmp((_this).F23()) == 0 {
			if (_520_v75).IsProperSubsetOf(_520_v75) {
				(globalState).F7 = false
				var _521_v76 _dafny.Array
				_ = _521_v76
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_9
				var _nw79 _dafny.Array
				_ = _nw79
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw79 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.Int = func(_522_i7 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_522_i7, (_this).F23())
					}
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw79 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw79).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw79).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_521_v76 = _nw79
				var _523_v77 _dafny.Array
				_ = _523_v77
				var _nwElement0_14 _dafny.Array = _521_v76
				_ = _nwElement0_14
				var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(20))
				_ = _nw80
				(_nw80).ArraySet1(_nwElement0_14, 0)
				(_nw80).ArraySet1(_521_v76, 1)
				(_nw80).ArraySet1(_521_v76, 2)
				(_nw80).ArraySet1(_521_v76, 3)
				(_nw80).ArraySet1(_521_v76, 4)
				(_nw80).ArraySet1(_521_v76, 5)
				(_nw80).ArraySet1(_521_v76, 6)
				(_nw80).ArraySet1(_521_v76, 7)
				(_nw80).ArraySet1(_521_v76, 8)
				(_nw80).ArraySet1(_521_v76, 9)
				(_nw80).ArraySet1(_521_v76, 10)
				(_nw80).ArraySet1(_521_v76, 11)
				(_nw80).ArraySet1(_521_v76, 12)
				(_nw80).ArraySet1(_521_v76, 13)
				(_nw80).ArraySet1(_521_v76, 14)
				(_nw80).ArraySet1((func() _dafny.Array {
					if (_this).F28() {
						return _521_v76
					}
					return _521_v76
				})(), 15)
				(_nw80).ArraySet1(_521_v76, 16)
				(_nw80).ArraySet1(_521_v76, 17)
				(_nw80).ArraySet1(_521_v76, 18)
				(_nw80).ArraySet1(_521_v76, 19)
				_523_v77 = _nw80
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_523_v77), 0))
				_ = _index107
				(_523_v77).ArraySet1(_521_v76, (_index107).Int())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_523_v77), 0))
				_ = _index108
				(_523_v77).ArraySet1(_521_v76, (_index108).Int())
				var _524_v78 _dafny.MultiSet
				_ = _524_v78
				_524_v78 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt((_this).F23(), (_this).F23()))
				(globalState).F8 = (func() _dafny.Int {
					if (_524_v78).Contains((_this).F23()) {
						return (_524_v78).Multiplicity((_this).F23())
					}
					return (_this).F23()
				})()
				(globalState).F21 = (_this).F28()
				var _525_v79 _dafny.MultiSet
				_ = _525_v79
				_525_v79 = _dafny.MultiSetOf(_521_v76, _521_v76, _521_v76)
				_525_v79 = _525_v79
			} else {
				var _526_v80 D8
				_ = _526_v80
				_526_v80 = Companion_D8_.Create_DC19_(_dafny.Companion_Sequence_.Update(_this.F29(), (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32(), (_this).F23()))
				var _527_v81 _dafny.MultiSet
				_ = _527_v81
				_527_v81 = _dafny.MultiSetOf(_this.F29(), (_526_v80).Dtor_cf32())
				(globalState).F21 = (_527_v81).Equals(_527_v81)
				(globalState).F6 = Companion_Default___.Fm2(globalState)
				_415_v0 = p0
				var _528_v82 _dafny.Array
				_ = _528_v82
				var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw81
				_528_v82 = _nw81
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_528_v82), 0))
				_ = _index109
				(_528_v82).ArraySet1((_this).F23(), (_index109).Int())
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_528_v82), 0))
				_ = _index110
				(_528_v82).ArraySet1((_this).F23(), (_index110).Int())
				var _529_v83 D3
				_ = _529_v83
				_529_v83 = Companion_D3_.Create_DC11_((_this).F23())
				var _530_v84 _dafny.Map
				_ = _530_v84
				_530_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_520_v75, _529_v83)
				var _531_v86 _dafny.Sequence
				_ = _531_v86
				_531_v86 = _dafny.SeqOf(_520_v75)
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_528_v82), 0))
				_ = _index111
				var _rhs110 _dafny.Map = (func() _dafny.Map {
					var _coll19 = _dafny.NewMapBuilder()
					_ = _coll19
					for _iter22 := _dafny.Iterate((_531_v86).Elements()); ; {
						_compr_19, _ok22 := _iter22()
						if !_ok22 {
							break
						}
						var _532_v85 _dafny.Set
						_532_v85 = interface{}(_compr_19).(_dafny.Set)
						if _dafny.Companion_Sequence_.Contains(_531_v86, _532_v85) {
							_coll19.Add(_532_v85, _529_v83)
						}
					}
					return _coll19.ToMap()
				}()).Merge(_530_v84)
				_ = _rhs110
				var _rhs111 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_454_v36, _dafny.UnicodeSeqOfUtf8Bytes("udxyoam"))).Cardinality())), (_528_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_528_v82), 0))).Int()).(_dafny.Int))
				_ = _rhs111
				var _lhs104 _dafny.Array = _528_v82
				_ = _lhs104
				var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_528_v82), 0))
				_ = _lhs105
				_530_v84 = _rhs110
				(_lhs104).ArraySet1(_rhs111, (_lhs105).Int())
			}
			(globalState).F21 = ((_this).F23()).Cmp(_dafny.IntOfInt64(805)) == 0
			var _533_v87 _dafny.Sequence
			_ = _533_v87
			_533_v87 = _dafny.SeqOf(p1)
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((p0), 0))
			_ = _index112
			(p0).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_533_v87, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_533_v87).Cardinality()))).Uint32(), (_this).F28()), _dafny.SeqOf(true, (_this).F28())), (_index112).Int())
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((p0), 0))
			_ = _index113
			(p0).ArraySet1(p1, (_index113).Int())
			var _534_v89 _dafny.Map
			_ = _534_v89
			_534_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, func() _dafny.Map {
				var _coll20 = _dafny.NewMapBuilder()
				_ = _coll20
				for _iter23 := _dafny.Iterate((_454_v36).Elements()); ; {
					_compr_20, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _535_v88 _dafny.CodePoint
					_535_v88 = interface{}(_compr_20).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_454_v36, _535_v88) {
						_coll20.Add(_535_v88, (_this).F23())
					}
				}
				return _coll20.ToMap()
			}())
			var _536_v90 _dafny.Set
			_ = _536_v90
			_536_v90 = _dafny.SetOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((p0), 0))).Int()).(bool))
			var _537_v91 _dafny.Map
			_ = _537_v91
			_537_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_455_v37, (_534_v89).Cardinality())).Cardinality(), true, globalState), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_this.F29(), (Companion_Default___.SafeIndex((_536_v90).Cardinality(), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32(), (_this).F23()), (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_this.F29(), (Companion_Default___.SafeIndex((_536_v90).Cardinality(), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32(), (_this).F23())).Cardinality()))).Uint32(), _dafny.IntOfInt64(943)))
			var _538_v92 _dafny.Array
			_ = _538_v92
			var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw82
			_538_v92 = _nw82
			var _539_v93 _dafny.Map
			_ = _539_v93
			_539_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_538_v92, _this.F29())
			var _540_v94 *C1
			_ = _540_v94
			var _nw83 *C1 = New_C1_()
			_ = _nw83
			_nw83.Ctor__(false, (func() _dafny.Sequence {
				if (_539_v93).Contains(_538_v92) {
					return (_539_v93).Get(_538_v92).(_dafny.Sequence)
				}
				return _this.F29()
			})(), (_this).F23())
			_540_v94 = _nw83
			var _541_v95 _dafny.Map
			_ = _541_v95
			_541_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
				if (_537_v91).Contains((_this).F23()) {
					return (_537_v91).Get((_this).F23()).(_dafny.Sequence)
				}
				return _dafny.SeqOf((_537_v91).Cardinality())
			})(), _540_v94)
			_541_v95 = (_541_v95).Update(_dafny.Companion_Sequence_.Concatenate(_this.F29(), _dafny.SeqOf((_this).F23())), _540_v94)
			(globalState).F21 = true
		} else {
			var _542_v96 _dafny.Sequence
			_ = _542_v96
			_542_v96 = _dafny.SeqOf((_this).F28(), false, p1)
			var _rhs112 _dafny.Int = (_this).F23()
			_ = _rhs112
			var _rhs113 _dafny.Sequence = _542_v96
			_ = _rhs113
			var _rhs114 bool = p1
			_ = _rhs114
			var _lhs106 *GlobalState = globalState
			_ = _lhs106
			r0 = _rhs112
			_542_v96 = _rhs113
			_lhs106.F9 = _rhs114
			var _543_v97 _dafny.Map
			_ = _543_v97
			_543_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F28())
			_543_v97 = (_543_v97).Update(_dafny.IntOfUint32((_454_v36).Cardinality()), !(true))
			var _544_v98 _dafny.Array
			_ = _544_v98
			var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
			_ = _nw84
			_544_v98 = _nw84
			var _545_v99 _dafny.Array
			_ = _545_v99
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_10
			var _nw85 _dafny.Array
			_ = _nw85
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw85 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.CodePoint = (func(_546_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_547_i8 _dafny.Int) _dafny.CodePoint {
						return _546_v37
					}
				})(_455_v37)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw85 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw85).ArraySet1CodePoint(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw85).ArraySet1CodePoint(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_545_v99 = _nw85
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_544_v98), 0))
			_ = _index114
			(_544_v98).ArraySet1(_545_v99, (_index114).Int())
			var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_544_v98), 0))
			_ = _index115
			var _rhs115 _dafny.Int = (_dafny.Zero).Minus((_this.F29()).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs115
			var _rhs116 _dafny.Array = _545_v99
			_ = _rhs116
			var _rhs117 bool = !(p1)
			_ = _rhs117
			var _rhs118 _dafny.Int = (_this).F23()
			_ = _rhs118
			var _lhs107 *GlobalState = globalState
			_ = _lhs107
			var _lhs108 _dafny.Array = _544_v98
			_ = _lhs108
			var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_544_v98), 0))
			_ = _lhs109
			var _lhs110 *GlobalState = globalState
			_ = _lhs110
			var _lhs111 *GlobalState = globalState
			_ = _lhs111
			_lhs107.F14 = _rhs115
			(_lhs108).ArraySet1(_rhs116, (_lhs109).Int())
			_lhs110.F9 = _rhs117
			_lhs111.F14 = _rhs118
			var _548_v100 _dafny.Map
			_ = _548_v100
			_548_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F23())
			var _549_v101 *C2
			_ = _549_v101
			var _nw86 *C2 = New_C2_()
			_ = _nw86
			_nw86.Ctor__(_548_v100, p1, (!(p1)) || (p1), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F23(), (_this).F23()), _this.F29()), (func() _dafny.Int {
				if (_this).F28() {
					return (_this).F23()
				}
				return (_this).F23()
			})())
			_549_v101 = _nw86
			(globalState).F8 = (_this).F23()
		}
		r0 = (_this).F23()
		var _550_v102 _dafny.MultiSet
		_ = _550_v102
		_550_v102 = _dafny.MultiSetOf((_this).F23())
		var _551_v103 _dafny.Map
		_ = _551_v103
		_551_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F23())
		var _552_v104 _dafny.Map
		_ = _552_v104
		_552_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_455_v37, (func() _dafny.Int {
			if (_550_v102).Contains((_551_v103).Cardinality()) {
				return (_550_v102).Multiplicity((_551_v103).Cardinality())
			}
			return (_this).F23()
		})())
		r1 = (_520_v75).Difference(_dafny.SetOf((func() _dafny.Set {
			var _coll21 = _dafny.NewBuilder()
			_ = _coll21
			for _iter24 := _dafny.Iterate((_552_v104).Keys().Elements()); ; {
				_compr_21, _ok24 := _iter24()
				if !_ok24 {
					break
				}
				var _553_v105 _dafny.CodePoint
				_553_v105 = interface{}(_compr_21).(_dafny.CodePoint)
				if (_552_v104).Contains(_553_v105) {
					_coll21.Add(_553_v105)
				}
			}
			return _coll21.ToSet()
		}()).Cardinality(), (_this).F23(), _dafny.IntOfUint32((_454_v36).Cardinality())))
		return r0, r1
	}
}
func (_this *C3) M2(p0 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var _554_v0 _dafny.Array
		_ = _554_v0
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_11
		var _nw87 _dafny.Array
		_ = _nw87
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw87 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Int = func(_555_i0 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_555_i0, (_this).F23())
			}
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw87 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw87).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw87).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_554_v0 = _nw87
		var _556_v1 _dafny.Sequence
		_ = _556_v1
		_556_v1 = _dafny.SeqOf(_554_v0)
		var _557_v2 D1
		_ = _557_v2
		_557_v2 = Companion_D1_.Create_DC3_(_556_v1)
		var _558_v3 _dafny.MultiSet
		_ = _558_v3
		_558_v3 = _dafny.MultiSetOf(_557_v2, _557_v2, _557_v2)
		var _559_v4 _dafny.Map
		_ = _559_v4
		_559_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), _558_v3)
		var _560_v5 _dafny.Sequence
		_ = _560_v5
		_560_v5 = _dafny.SeqOf((_this).F28())
		_559_v4 = (_559_v4).Update(_560_v5, _558_v3)
		var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))
		_ = _index116
		(_554_v0).ArraySet1((_this).F23(), (_index116).Int())
		var _561_v6 _dafny.Sequence
		_ = _561_v6
		_561_v6 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(317))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}(func(_562_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})))
		var _563_v7 _dafny.Sequence
		_ = _563_v7
		_563_v7 = _561_v6
		var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))
		_ = _index117
		(_554_v0).ArraySet1(func(_source4 _dafny.Sequence) _dafny.Int {
			var _564___mcc_h0 _dafny.Sequence = _source4
			_ = _564___mcc_h0
			var _565_cf30 _dafny.Sequence = _564___mcc_h0
			_ = _565_cf30
			return (_this).F23()
		}((func() _dafny.Sequence {
			if (_this).F28() {
				return _563_v7
			}
			return _561_v6
		})()), (_index117).Int())
		var _566_i2 _dafny.Int
		_ = _566_i2
		_566_i2 = _dafny.Zero
		{
			for !((_this).F28()) || (_dafny.Companion_Sequence_.Contains(_this.F29(), (_this).F23())) {
				{
					if (_566_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_566_i2 = (_566_i2).Plus(_dafny.One)
					var _567_v8 _dafny.MultiSet
					_ = _567_v8
					_567_v8 = _dafny.MultiSetOf((_this).F28())
					var _568_v9 _dafny.CodePoint
					_ = _568_v9
					_568_v9 = _dafny.CodePoint('g')
					var _569_v10 D2
					_ = _569_v10
					_569_v10 = Companion_D2_.Create_DC8_((_this).F23(), (_567_v8).Cardinality(), _568_v9, _dafny.IntOfUint32((p0).Cardinality()), (_this).F28())
					var _rhs119 bool = !(!((_569_v10).Dtor_cf19()))
					_ = _rhs119
					var _rhs120 _dafny.Int = ((_this).F23()).Plus((_dafny.IntOfInt64(692)).Plus((_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int)))
					_ = _rhs120
					var _rhs121 bool = !((_567_v8).Update((_this).F28(), Companion_Default___.Abs((_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int)))).Equals(_567_v8)
					_ = _rhs121
					var _lhs112 *GlobalState = globalState
					_ = _lhs112
					var _lhs113 *GlobalState = globalState
					_ = _lhs113
					var _lhs114 *GlobalState = globalState
					_ = _lhs114
					_lhs112.F7 = _rhs119
					_lhs113.F8 = _rhs120
					_lhs114.F21 = _rhs121
					(globalState).F21 = false
					(globalState).F11 = _568_v9
					(globalState).F8 = (_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int)
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _570_v11 _dafny.Array
		_ = _570_v11
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_12
		var _nw88 _dafny.Array
		_ = _nw88
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw88 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.CodePoint = func(_571_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			}
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw88 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw88).ArraySet1CodePoint(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw88).ArraySet1CodePoint(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_570_v11 = _nw88
		var _572_v12 _dafny.Map
		_ = _572_v12
		_572_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_570_v11, (_this).F28())
		_572_v12 = (_572_v12).Update(_570_v11, (_this).F28())
		if false {
			var _573_v14 _dafny.Map
			_ = _573_v14
			_573_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_560_v5).Cardinality()), (_this).F28())
			var _574_v15 _dafny.Sequence
			_ = _574_v15
			_574_v15 = _dafny.SeqOf(_dafny.SeqOf((_573_v14).Cardinality()))
			var _575_v16 *C2
			_ = _575_v16
			var _nw89 *C2 = New_C2_()
			_ = _nw89
			_nw89.Ctor__(func() _dafny.Map {
				var _coll22 = _dafny.NewMapBuilder()
				_ = _coll22
				for _iter25 := _dafny.Iterate(((_574_v15).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_574_v15).Cardinality()))).Uint32()).(_dafny.Sequence)).Elements()); ; {
					_compr_22, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _576_v13 _dafny.Int
					_576_v13 = interface{}(_compr_22).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains((_574_v15).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_574_v15).Cardinality()))).Uint32()).(_dafny.Sequence), _576_v13) {
						_coll22.Add((_576_v13).Minus((_this).F23()), (_this).F23())
					}
				}
				return _coll22.ToMap()
			}(), (_this).F28(), (_this).F28(), _this.F29(), (_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int))
			_575_v16 = _nw89
			var _577_v17 _dafny.MultiSet
			_ = _577_v17
			_577_v17 = _dafny.MultiSetOf((_575_v16).F32(), (_575_v16).F32(), (_this).F28())
			(globalState).F6 = (_561_v6).Select((Companion_Default___.SafeIndex((_577_v17).Cardinality(), _dafny.IntOfUint32((_561_v6).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _578_v18 _dafny.Array
			_ = _578_v18
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_13
			var _nw90 _dafny.Array
			_ = _nw90
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw90 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Sequence = (func(_579_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_580_i4 _dafny.Int) _dafny.Sequence {
						return _579_p0
					}
				})(p0)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw90 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw90).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw90).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_578_v18 = _nw90
			_578_v18 = _578_v18
			var _581_v19 _dafny.Array
			_ = _581_v19
			var _nw91 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw91
			_581_v19 = _nw91
			var _582_v20 _dafny.Map
			_ = _582_v20
			_582_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_581_v19, (_dafny.IntOfInt64(945)).Cmp((_dafny.Zero).Minus(_dafny.IntOfInt64(-555))) >= 0)
			(globalState).F21 = (func() bool {
				if (_582_v20).Contains(_581_v19) {
					return (_582_v20).Get(_581_v19).(bool)
				}
				return (_this).F28()
			})()
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_581_v19), 0))
			_ = _index118
			(_581_v19).ArraySet1((_575_v16).F32(), (_index118).Int())
			var _583_v21 _dafny.Map
			_ = _583_v21
			_583_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_575_v16).F32(), (_575_v16).F32())
			var _584_v22 _dafny.Map
			_ = _584_v22
			_584_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_583_v21).Cardinality())
			var _585_v23 _dafny.MultiSet
			_ = _585_v23
			_585_v23 = _dafny.MultiSetOf(_dafny.IntOfInt64(39), (_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int))
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_581_v19), 0))
			_ = _index119
			(_581_v19).ArraySet1((((_584_v22).Cardinality()).Minus((_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int))).Cmp((func() _dafny.Int {
				if (_585_v23).Contains((_this).F23()) {
					return (_585_v23).Multiplicity((_this).F23())
				}
				return _dafny.IntOfUint32((_560_v5).Cardinality())
			})()) != 0, (_index119).Int())
		} else {
			var _586_v24 _dafny.Map
			_ = _586_v24
			_586_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_561_v6).Cardinality())), _dafny.IntOfInt64(63))
			var _587_v25 T1
			_ = _587_v25
			var _nw92 *C2 = New_C2_()
			_ = _nw92
			_nw92.Ctor__(_586_v24, (_this).F28(), (_this).F28(), _this.F29(), (_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int))
			_587_v25 = _nw92
			var _588_v26 T1
			_ = _588_v26
			_588_v26 = _587_v25
			var _source5 T1 = _588_v26
			_ = _source5
			var _589___mcc_h1 T1 = _source5
			_ = _589___mcc_h1
			var _590_cf29 T1 = _589___mcc_h1
			_ = _590_cf29
			var _591_v27 _dafny.MultiSet
			_ = _591_v27
			_591_v27 = _dafny.MultiSetOf((_587_v25).F23(), (_590_cf29).F23(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_560_v5).Cardinality())))
			var _592_v28 _dafny.MultiSet
			_ = _592_v28
			_592_v28 = _dafny.MultiSetOf(_591_v27, _dafny.MultiSetFromSeq(_this.F29()), (_591_v27).Union(_591_v27), (_591_v27).Update((_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int), Companion_Default___.Abs((_587_v25).F23())), _591_v27)
			var _593_v29 _dafny.MultiSet
			_ = _593_v29
			_593_v29 = _dafny.MultiSetOf(_570_v11, _570_v11, _570_v11)
			var _rhs122 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F23()), ((_this).F23()).Plus(_dafny.IntOfUint32((p0).Cardinality())))
			_ = _rhs122
			var _rhs123 _dafny.MultiSet = _dafny.MultiSetOf((_dafny.MultiSetOf((_587_v25).F23(), (_593_v29).Cardinality(), (_587_v25).F23(), (_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((p0).Cardinality()))).Update((_587_v25).F23(), Companion_Default___.Abs((_590_cf29).F23())), _591_v27, (_591_v27).Intersection(_591_v27), _591_v27)
			_ = _rhs123
			r0 = _rhs122
			_592_v28 = _rhs123
			var _594_v30 _dafny.Array
			_ = _594_v30
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_14
			var _nw93 _dafny.Array
			_ = _nw93
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw93 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.MultiSet = (func(_595_v25 T1, _596_cf29 T1) func(_dafny.Int) _dafny.MultiSet {
					return func(_597_i5 _dafny.Int) _dafny.MultiSet {
						return (_dafny.MultiSetOf(false, !((_595_v25).F28()), (_596_cf29).F28(), (_596_cf29).F28())).Union(_dafny.MultiSetOf((_595_v25).F28(), false, !(false), (_595_v25).F28()))
					}
				})(_587_v25, _590_cf29)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw93 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw93).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw93).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_594_v30 = _nw93
			var _rhs124 bool = (_587_v25).F28()
			_ = _rhs124
			var _rhs125 _dafny.Array = _594_v30
			_ = _rhs125
			var _lhs115 *GlobalState = globalState
			_ = _lhs115
			_lhs115.F21 = _rhs124
			_594_v30 = _rhs125
			var _598_v31 D3
			_ = _598_v31
			_598_v31 = Companion_D3_.Create_DC11_((_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int))
			var _pat_let_tv27 = _587_v25
			_ = _pat_let_tv27
			var _599_v32 _dafny.Map
			_ = _599_v32
			_599_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(func(_pat_let2_0 D3) D3 {
				return func(_600_dt__update__tmp_h0 D3) D3 {
					return func(_pat_let3_0 _dafny.Int) D3 {
						return func(_601_dt__update_hcf22_h0 _dafny.Int) D3 {
							return Companion_D3_.Create_DC11_(_601_dt__update_hcf22_h0)
						}(_pat_let3_0)
					}((_pat_let_tv27).F23())
				}(_pat_let2_0)
			}(_598_v31)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(490))).Uint32(), func(coer40 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
				return func(arg40 _dafny.Int) interface{} {
					return coer40(arg40)
				}
			}((func(_602_v31 D3) func(_dafny.Int) D3 {
				return func(_603_i6 _dafny.Int) D3 {
					return _602_v31
				}
			})(_598_v31)))), _dafny.IntOfInt64(-173))
			var _604_v33 _dafny.Sequence
			_ = _604_v33
			_604_v33 = _dafny.SeqOf(Companion_D3_.Create_DC11_((_590_cf29).F23()), _598_v31)
			_599_v32 = (_599_v32).Update((func() _dafny.Sequence {
				if (_590_cf29).F28() {
					return _604_v33
				}
				return _604_v33
			})(), (_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int))
			var _605_v34 _dafny.MultiSet
			_ = _605_v34
			_605_v34 = _dafny.MultiSetOf(_554_v0, _554_v0, _554_v0)
			(globalState).F7 = ((_dafny.MultiSetOf(_554_v0)).Update((_556_v1).Select((Companion_Default___.SafeIndex((_590_cf29).F23(), _dafny.IntOfUint32((_556_v1).Cardinality()))).Uint32()).(_dafny.Array), Companion_Default___.Abs((_587_v25).F23()))).IsProperSubsetOf(_605_v34)
			var _606_v35 _dafny.Map
			_ = _606_v35
			_606_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_554_v0, _587_v25.F29())
			var _607_v36 _dafny.Map
			_ = _607_v36
			_607_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_606_v35).Update(_554_v0, _this.F29()))
			_607_v36 = (_607_v36).Update((_587_v25).F28(), _606_v35)
			if !(Companion_Default___.Fm1((_587_v25).F23(), (_587_v25).F28(), globalState)) {
				(globalState).F7 = (_587_v25).F28()
				_570_v11 = _570_v11
				(globalState).F9 = false
				_587_v25 = _587_v25
				var _608_v37 _dafny.Array
				_ = _608_v37
				var _nwElement0_15 bool = (_587_v25).F28()
				_ = _nwElement0_15
				var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(5))
				_ = _nw94
				(_nw94).ArraySet1(_nwElement0_15, 0)
				(_nw94).ArraySet1((_587_v25).F28(), 1)
				(_nw94).ArraySet1(Companion_Default___.Fm1((_dafny.Zero).Minus((_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int)), (_587_v25).F28(), globalState), 2)
				(_nw94).ArraySet1((_587_v25).F28(), 3)
				(_nw94).ArraySet1(Companion_Default___.Fm1((_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int), (_587_v25).F28(), globalState), 4)
				_608_v37 = _nw94
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_608_v37), 0))
				_ = _index120
				(_608_v37).ArraySet1((_this).F28(), (_index120).Int())
				var _609_v38 _dafny.CodePoint
				_ = _609_v38
				_609_v38 = _dafny.CodePoint('a')
				var _610_v39 D9
				_ = _610_v39
				_610_v39 = Companion_D9_.Create_DC24_(_570_v11)
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_608_v37), 0))
				_ = _index121
				var _rhs126 bool = !_dafny.Companion_Sequence_.Contains(p0, _609_v38)
				_ = _rhs126
				var _rhs127 _dafny.Int = _dafny.IntOfInt64(-741)
				_ = _rhs127
				var _rhs128 _dafny.Array = (_610_v39).Dtor_cf40()
				_ = _rhs128
				var _rhs129 bool = _dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(729))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}((func(_611_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_612_i7 _dafny.Int) _dafny.CodePoint {
						return _611_v38
					}
				})(_609_v38))), p0)
				_ = _rhs129
				var _rhs130 _dafny.Array = _608_v37
				_ = _rhs130
				var _lhs116 *GlobalState = globalState
				_ = _lhs116
				var _lhs117 *GlobalState = globalState
				_ = _lhs117
				var _lhs118 _dafny.Array = _608_v37
				_ = _lhs118
				var _lhs119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_608_v37), 0))
				_ = _lhs119
				_lhs116.F21 = _rhs126
				_lhs117.F14 = _rhs127
				_570_v11 = _rhs128
				(_lhs118).ArraySet1(_rhs129, (_lhs119).Int())
				_608_v37 = _rhs130
			} else {
				var _613_v40 _dafny.Map
				_ = _613_v40
				_613_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_560_v5, (_this).F28())
				var _614_v42 _dafny.CodePoint
				_ = _614_v42
				_614_v42 = _dafny.CodePoint('w')
				var _615_v43 _dafny.Map
				_ = _615_v43
				_615_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_614_v42, (_587_v25).F28())
				_613_v40 = (func() _dafny.Map {
					if (func() _dafny.Map {
						var _coll23 = _dafny.NewMapBuilder()
						_ = _coll23
						for _iter26 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(770))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg42 _dafny.Int) interface{} {
								return coer42(arg42)
							}
						}(func(_616_i8 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(177))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg43 _dafny.Int) interface{} {
									return coer43(arg43)
								}
							}(func(_617_i9 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('y')
							}))).Cardinality())
						}))).Elements()); ; {
							_compr_23, _ok26 := _iter26()
							if !_ok26 {
								break
							}
							var _618_v41 _dafny.Int
							_618_v41 = interface{}(_compr_23).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(770))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg44 _dafny.Int) interface{} {
									return coer44(arg44)
								}
							}(func(_616_i8 _dafny.Int) _dafny.Int {
								return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(177))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg45 _dafny.Int) interface{} {
										return coer45(arg45)
									}
								}(func(_617_i9 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('y')
								}))).Cardinality())
							})), _618_v41) {
								_coll23.Add(Companion_Default___.SafeModuloInt(_618_v41, (_587_v25).F23()), (_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int))
							}
						}
						return _coll23.ToMap()
					}()).Equals(Companion_Default___.Fm16((_this).F28(), (_615_v43).Cardinality(), (_this).F23(), globalState)) {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_560_v5, (_this).F28())
					}
					return _613_v40
				})()
				var _619_v44 D9
				_ = _619_v44
				_619_v44 = Companion_D9_.Create_DC26_(_dafny.IntOfInt64(548))
				var _pat_let_tv28 = _554_v0
				_ = _pat_let_tv28
				var _pat_let_tv29 = _554_v0
				_ = _pat_let_tv29
				var _620_v45 _dafny.Array
				_ = _620_v45
				var _nwElement0_16 D9 = Companion_D9_.Create_DC26_((_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int))
				_ = _nwElement0_16
				var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(22))
				_ = _nw95
				(_nw95).ArraySet1(_nwElement0_16, 0)
				(_nw95).ArraySet1(_619_v44, 1)
				(_nw95).ArraySet1(_619_v44, 2)
				(_nw95).ArraySet1(_619_v44, 3)
				(_nw95).ArraySet1(_619_v44, 4)
				(_nw95).ArraySet1(_619_v44, 5)
				(_nw95).ArraySet1(_619_v44, 6)
				(_nw95).ArraySet1(_619_v44, 7)
				(_nw95).ArraySet1(_619_v44, 8)
				(_nw95).ArraySet1(_619_v44, 9)
				(_nw95).ArraySet1(_619_v44, 10)
				(_nw95).ArraySet1(_619_v44, 11)
				(_nw95).ArraySet1(_619_v44, 12)
				(_nw95).ArraySet1(Companion_D9_.Create_DC26_(_dafny.IntOfInt64(604)), 13)
				(_nw95).ArraySet1(func(_pat_let4_0 D9) D9 {
					return func(_621_dt__update__tmp_h1 D9) D9 {
						return func(_pat_let5_0 _dafny.Int) D9 {
							return func(_622_dt__update_hcf43_h0 _dafny.Int) D9 {
								return Companion_D9_.Create_DC26_(_622_dt__update_hcf43_h0)
							}(_pat_let5_0)
						}((_this).F23())
					}(_pat_let4_0)
				}(_619_v44), 14)
				(_nw95).ArraySet1(_619_v44, 15)
				(_nw95).ArraySet1(_619_v44, 16)
				(_nw95).ArraySet1(_619_v44, 17)
				(_nw95).ArraySet1(func(_pat_let6_0 D9) D9 {
					return func(_623_dt__update__tmp_h2 D9) D9 {
						return func(_pat_let7_0 _dafny.Int) D9 {
							return func(_624_dt__update_hcf43_h1 _dafny.Int) D9 {
								return Companion_D9_.Create_DC26_(_624_dt__update_hcf43_h1)
							}(_pat_let7_0)
						}((_pat_let_tv29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_pat_let_tv28), 0))).Int()).(_dafny.Int))
					}(_pat_let6_0)
				}(_619_v44), 18)
				(_nw95).ArraySet1(_619_v44, 19)
				(_nw95).ArraySet1(_619_v44, 20)
				(_nw95).ArraySet1(_619_v44, 21)
				_620_v45 = _nw95
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_620_v45), 0))
				_ = _index122
				(_620_v45).ArraySet1(_619_v44, (_index122).Int())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_620_v45), 0))
				_ = _index123
				(_620_v45).ArraySet1(_619_v44, (_index123).Int())
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))
				_ = _index124
				(_554_v0).ArraySet1(_dafny.IntOfInt64(721), (_index124).Int())
				var _625_v46 _dafny.Array
				_ = _625_v46
				var _nw96 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
				_ = _nw96
				_625_v46 = _nw96
				var _626_v47 _dafny.MultiSet
				_ = _626_v47
				_626_v47 = _dafny.MultiSetOf(true)
				var _627_v48 _dafny.Sequence
				_ = _627_v48
				_627_v48 = _dafny.SeqOf(_dafny.SetOf((_this).F28()))
				var _628_v49 D0
				_ = _628_v49
				_628_v49 = Companion_D0_.Create_DC0_(p0)
				var _629_v50 _dafny.MultiSet
				_ = _629_v50
				_629_v50 = _dafny.MultiSetOf(_dafny.IntOfUint32((p0).Cardinality()))
				var _630_v51 _dafny.Array
				_ = _630_v51
				var _nwElement0_17 _dafny.MultiSet = _626_v47
				_ = _nwElement0_17
				var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(22))
				_ = _nw97
				(_nw97).ArraySet1(_nwElement0_17, 0)
				(_nw97).ArraySet1(_626_v47, 1)
				(_nw97).ArraySet1(_626_v47, 2)
				(_nw97).ArraySet1(_626_v47, 3)
				(_nw97).ArraySet1(_626_v47, 4)
				(_nw97).ArraySet1(_626_v47, 5)
				(_nw97).ArraySet1(_626_v47, 6)
				(_nw97).ArraySet1(_dafny.MultiSetOf((_587_v25).F28(), (_this).F28()), 7)
				(_nw97).ArraySet1(_dafny.MultiSetOf((_587_v25).F28(), (_this).F28()), 8)
				(_nw97).ArraySet1(_626_v47, 9)
				(_nw97).ArraySet1(_626_v47, 10)
				(_nw97).ArraySet1(_626_v47, 11)
				(_nw97).ArraySet1((((_this).F30()).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32(((_this).F30()).Cardinality()))).Uint32()).(_dafny.MultiSet)).Update((_this).F28(), Companion_Default___.Abs(_dafny.IntOfUint32((_627_v48).Cardinality()))), 12)
				(_nw97).ArraySet1(_626_v47, 13)
				(_nw97).ArraySet1(_626_v47, 14)
				(_nw97).ArraySet1((_626_v47).Update((_587_v25).F28(), Companion_Default___.Abs((_587_v25).F23())), 15)
				(_nw97).ArraySet1(_dafny.MultiSetOf((_587_v25).F28(), (_587_v25).F28()), 16)
				(_nw97).ArraySet1(Companion_Default___.Fm4(_628_v49, (_this).F23(), _560_v5, _614_v42, globalState), 17)
				(_nw97).ArraySet1(_dafny.MultiSetFromSeq(_560_v5), 18)
				(_nw97).ArraySet1(_dafny.MultiSetOf((_587_v25).F28(), (_this).F28()), 19)
				(_nw97).ArraySet1(_626_v47, 20)
				(_nw97).ArraySet1(Companion_Default___.Fm4(_628_v49, (func() _dafny.Int {
					if (_629_v50).Contains((_626_v47).Cardinality()) {
						return (_629_v50).Multiplicity((_626_v47).Cardinality())
					}
					return (_587_v25).F23()
				})(), _560_v5, _614_v42, globalState), 21)
				_630_v51 = _nw97
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_625_v46), 0))
				_ = _index125
				(_625_v46).ArraySet1(_630_v51, (_index125).Int())
				var _631_v52 _dafny.Map
				_ = _631_v52
				_631_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_587_v25).F23(), (_this).F28())
				var _632_v53 _dafny.Array
				_ = _632_v53
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_15
				var _nw98 _dafny.Array
				_ = _nw98
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw98 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) bool = func(_633_i10 _dafny.Int) bool {
						return false
					}
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw98 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw98).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw98).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_632_v53 = _nw98
				var _634_v54 _dafny.Map
				_ = _634_v54
				_634_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_632_v53, Companion_Default___.Fm1((_587_v25).F23(), (_this).F28(), globalState))
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_625_v46), 0))
				_ = _index126
				var _rhs131 bool = ((_631_v52).Cardinality()).Cmp((_634_v54).Cardinality()) >= 0
				_ = _rhs131
				var _rhs132 _dafny.Int = (_587_v25).F23()
				_ = _rhs132
				var _rhs133 _dafny.Array = _630_v51
				_ = _rhs133
				var _lhs120 *GlobalState = globalState
				_ = _lhs120
				var _lhs121 *GlobalState = globalState
				_ = _lhs121
				var _lhs122 _dafny.Array = _625_v46
				_ = _lhs122
				var _lhs123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_625_v46), 0))
				_ = _lhs123
				_lhs120.F21 = _rhs131
				_lhs121.F8 = _rhs132
				(_lhs122).ArraySet1(_rhs133, (_lhs123).Int())
				var _635_v55 D0
				_ = _635_v55
				var _636_v56 _dafny.Map
				_ = _636_v56
				var _637_v57 _dafny.Int
				_ = _637_v57
				var _out27 D0
				_ = _out27
				var _out28 _dafny.Map
				_ = _out28
				var _out29 _dafny.Int
				_ = _out29
				_out27, _out28, _out29 = Companion_Default___.M0(_632_v53, _dafny.Companion_Sequence_.Concatenate(p0, _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _614_v42)), (_587_v25).F23(), globalState)
				_635_v55 = _out27
				_636_v56 = _out28
				_637_v57 = _out29
			}
			var _638_v58 D8
			_ = _638_v58
			_638_v58 = Companion_D8_.Create_DC22_((_this).F28())
			(globalState).F8 = (func() _dafny.Int {
				if (_638_v58).Dtor_cf38() {
					return (_this).F23()
				}
				return (_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int)
			})()
			(globalState).F14 = (_this).F23()
		}
		var _639_v59 *C0
		_ = _639_v59
		var _nw99 *C0 = New_C0_()
		_ = _nw99
		_nw99.Ctor__()
		_639_v59 = _nw99
		var _640_v60 _dafny.Map
		_ = _640_v60
		_640_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rtk"), (_this).F23())
		var _641_v61 _dafny.Map
		_ = _641_v61
		_641_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(499), _dafny.UnicodeSeqOfUtf8Bytes("gun"))
		var _642_v62 _dafny.CodePoint
		_ = _642_v62
		_642_v62 = _dafny.CodePoint('a')
		var _643_v63 _dafny.Set
		_ = _643_v63
		_643_v63 = _dafny.SetOf((_this).F23(), _dafny.IntOfInt64(-620), (_this).F23(), (_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int), (_this).F23())
		var _644_v64 _dafny.Map
		_ = _644_v64
		_644_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_640_v60).Contains(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_641_v61).Contains(_dafny.IntOfInt64(-191)) {
					return (_641_v61).Get(_dafny.IntOfInt64(-191)).(_dafny.Sequence)
				}
				return p0
			})(), (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_641_v61).Contains(_dafny.IntOfInt64(-191)) {
					return (_641_v61).Get(_dafny.IntOfInt64(-191)).(_dafny.Sequence)
				}
				return p0
			})()).Cardinality()))).Uint32(), _642_v62)) {
				return (_640_v60).Get(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_641_v61).Contains(_dafny.IntOfInt64(-191)) {
						return (_641_v61).Get(_dafny.IntOfInt64(-191)).(_dafny.Sequence)
					}
					return p0
				})(), (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_641_v61).Contains(_dafny.IntOfInt64(-191)) {
						return (_641_v61).Get(_dafny.IntOfInt64(-191)).(_dafny.Sequence)
					}
					return p0
				})()).Cardinality()))).Uint32(), _642_v62)).(_dafny.Int)
			}
			return (_643_v63).Cardinality()
		})(), (_dafny.Zero).Minus((_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int)))
		r0 = (func() _dafny.Int {
			if (_644_v64).Contains((_dafny.Zero).Minus((_this).F23())) {
				return (_644_v64).Get((_dafny.Zero).Minus((_this).F23())).(_dafny.Int)
			}
			return (_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int)
		})()
		r1 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_this.F29(), _this.F29()), (Companion_Default___.SafeIndex((_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F29(), _this.F29())).Cardinality()))).Uint32(), (_554_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_554_v0), 0))).Int()).(_dafny.Int))
		return r0, r1
	}
}
func (_this *C3) M5(globalState *GlobalState) {
	{
		var _645_v0 _dafny.Array
		_ = _645_v0
		var _nw100 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
		_ = _nw100
		_645_v0 = _nw100
		for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_645_v0), 0))); ; {
			_guard_loop_3, _ok27 := _iter27()
			if !_ok27 {
				break
			}
			var _646_i0 _dafny.Int
			_646_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_646_i0).Sign() != -1) && ((_646_i0).Cmp(_dafny.ArrayLen((_645_v0), 0)) < 0)) {
				(_645_v0).ArraySet1(Companion_Default___.SafeModuloInt(_646_i0, (_dafny.IntOfUint32((_this.F29()).Cardinality())).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fgr")).Cardinality()))), (_646_i0).Int())
			}
		}
		var _647_i1 _dafny.Int
		_ = _647_i1
		_647_i1 = _dafny.Zero
		{
			for !((_this).F28()) {
				{
					if (_647_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_647_i1 = (_647_i1).Plus(_dafny.One)
					var _648_v1 *C0
					_ = _648_v1
					var _nw101 *C0 = New_C0_()
					_ = _nw101
					_nw101.Ctor__()
					_648_v1 = _nw101
					var _649_v2 _dafny.Array
					_ = _649_v2
					var _nw102 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
					_ = _nw102
					_649_v2 = _nw102
					var _650_v3 _dafny.Set
					_ = _650_v3
					_650_v3 = _dafny.SetOf((_dafny.Zero).Minus((_this).F23()))
					var _651_v4 _dafny.Map
					_ = _651_v4
					_651_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F23())
					var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_649_v2), 0))
					_ = _index127
					(_649_v2).ArraySet1((_650_v3).IsProperSubsetOf(_dafny.SetOf((func() _dafny.Int {
						if (_651_v4).Contains((_this.F29()).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32()).(_dafny.Int)) {
							return (_651_v4).Get((_this.F29()).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Int)
						}
						return (_this).F23()
					})())), (_index127).Int())
					var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_649_v2), 0))
					_ = _index128
					(_649_v2).ArraySet1(!((func() bool {
						if !((_this).F28()) {
							return (func() _dafny.Set {
								var _coll24 = _dafny.NewBuilder()
								_ = _coll24
								for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(979), _dafny.IntOfInt64(687))); ; {
									_compr_24, _ok28 := _iter28()
									if !_ok28 {
										break
									}
									var _652_v5 _dafny.Int
									_652_v5 = interface{}(_compr_24).(_dafny.Int)
									if ((_dafny.IntOfInt64(979)).Cmp(_652_v5) <= 0) && ((_652_v5).Cmp(_dafny.IntOfInt64(687)) < 0) {
										_coll24.Add((_652_v5).Plus((_this).F23()))
									}
								}
								return _coll24.ToSet()
							}()).IsDisjointFrom(_650_v3)
						}
						return (func() bool {
							if (_this).F28() {
								return (_this).F28()
							}
							return (_this).F28()
						})()
					})()), (_index128).Int())
					var _653_v6 _dafny.Set
					_ = _653_v6
					_653_v6 = _dafny.SetOf(false, (_this).F28())
					var _654_v7 _dafny.Sequence
					_ = _654_v7
					_654_v7 = _dafny.SeqOf(_653_v6)
					var _rhs134 _dafny.Int = (_this).F23()
					_ = _rhs134
					var _rhs135 _dafny.Set = (_654_v7).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_654_v7).Cardinality()))).Uint32()).(_dafny.Set)
					_ = _rhs135
					var _rhs136 _dafny.Int = (_this.F29()).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32()).(_dafny.Int)
					_ = _rhs136
					var _lhs124 *GlobalState = globalState
					_ = _lhs124
					var _lhs125 *GlobalState = globalState
					_ = _lhs125
					_lhs124.F14 = _rhs134
					_653_v6 = _rhs135
					_lhs125.F8 = _rhs136
					var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_649_v2), 0))
					_ = _index129
					(_649_v2).ArraySet1((_this).F28(), (_index129).Int())
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _655_v8 _dafny.Sequence
		_ = _655_v8
		_655_v8 = _dafny.UnicodeSeqOfUtf8Bytes("kgw")
		var _656_v9 _dafny.MultiSet
		_ = _656_v9
		_656_v9 = _dafny.MultiSetOf(_dafny.IntOfUint32((_655_v8).Cardinality()), (_this).F23())
		var _657_v10 _dafny.Map
		_ = _657_v10
		_657_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _this.F29())
		var _hi10 _dafny.Int = (func() _dafny.Int {
			if (_656_v9).Contains(_dafny.IntOfInt64(685)) {
				return (_656_v9).Multiplicity(_dafny.IntOfInt64(685))
			}
			return _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_657_v10).Contains(_dafny.IntOfInt64(260)) {
					return (_657_v10).Get(_dafny.IntOfInt64(260)).(_dafny.Sequence)
				}
				return _dafny.SeqOf((_this).F23(), (_this).F23())
			})()).Cardinality())
		})()
		_ = _hi10
		for _658_i2 := _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((Companion_Default___.Fm17((_this).F28(), Companion_Default___.Fm0((_this).F23(), (_this).F28(), globalState), globalState)).Cardinality(), (_this).F23()), _this.F29())).Cardinality()); _658_i2.Cmp(_hi10) < 0; _658_i2 = _658_i2.Plus(_dafny.One) {
			var _659_v11 _dafny.Sequence
			_ = _659_v11
			_659_v11 = _dafny.SeqOf((_this).F28())
			var _660_v12 _dafny.MultiSet
			_ = _660_v12
			_660_v12 = _dafny.MultiSetOf(!((_this).F28()), (_this).F28())
			var _661_v13 _dafny.Map
			_ = _661_v13
			_661_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)
			var _662_v14 _dafny.Map
			_ = _662_v14
			_662_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _661_v13)
			var _663_v15 _dafny.Set
			_ = _663_v15
			_663_v15 = _dafny.SetOf(_658_i2, Companion_Default___.SafeModuloInt((_this).F23(), (_656_v9).Cardinality()), (_658_i2).Minus(_dafny.IntOfUint32((_659_v11).Cardinality())), (((_660_v12).Update((_this).F28(), Companion_Default___.Abs(_658_i2))).Cardinality()).Times((_this).F23()), (_662_v14).Cardinality())
			_663_v15 = (func() _dafny.Set {
				if false {
					return _663_v15
				}
				return _663_v15
			})()
			var _664_v16 _dafny.CodePoint
			_ = _664_v16
			_664_v16 = _dafny.CodePoint('e')
			var _665_v17 _dafny.Map
			_ = _665_v17
			_665_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_this.F29()).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_655_v8, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_655_v8).Cardinality()))).Uint32(), _664_v16), (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_655_v8, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_655_v8).Cardinality()))).Uint32(), _664_v16)).Cardinality()))).Uint32(), _664_v16), _655_v8)).Cardinality()))
			var _666_v18 D4
			_ = _666_v18
			_666_v18 = Companion_D4_.Create_DC14_(_645_v0)
			var _667_v19 *C2
			_ = _667_v19
			var _nw103 *C2 = New_C2_()
			_ = _nw103
			_nw103.Ctor__(_665_v17, Companion_Default___.Fm1(_658_i2, (_this).F28(), globalState), (_this).F28(), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_this.F29(), (Companion_Default___.SafeIndex(_658_i2, _dafny.IntOfUint32((_this.F29()).Cardinality()))).Uint32(), _658_i2), _this.F29()), _dafny.IntOfUint32((_dafny.SeqOf(_666_v18, _666_v18)).Cardinality()))
			_667_v19 = _nw103
			_665_v17 = (_665_v17).Update(_dafny.IntOfUint32((_655_v8).Cardinality()), (Companion_Default___.Fm24((_667_v19).F32(), _658_i2, globalState)).Dtor_cf43())
			if (_this).F28() {
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_645_v0), 0))
				_ = _index130
				(_645_v0).ArraySet1(_658_i2, (_index130).Int())
				var _668_v20 _dafny.Array
				_ = _668_v20
				var _nwElement0_18 _dafny.CodePoint = _dafny.CodePoint('n')
				_ = _nwElement0_18
				var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(6))
				_ = _nw104
				(_nw104).ArraySet1CodePoint(_nwElement0_18, 0)
				(_nw104).ArraySet1CodePoint(_664_v16, 1)
				(_nw104).ArraySet1CodePoint(_664_v16, 2)
				(_nw104).ArraySet1CodePoint(_664_v16, 3)
				(_nw104).ArraySet1CodePoint(_664_v16, 4)
				(_nw104).ArraySet1CodePoint(_664_v16, 5)
				_668_v20 = _nw104
				var _669_v21 D9
				_ = _669_v21
				_669_v21 = Companion_D9_.Create_DC24_(_668_v20)
				var _670_v22 _dafny.Array
				_ = _670_v22
				var _nwElement0_19 _dafny.Array = _668_v20
				_ = _nwElement0_19
				var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(19))
				_ = _nw105
				(_nw105).ArraySet1(_nwElement0_19, 0)
				(_nw105).ArraySet1(_668_v20, 1)
				(_nw105).ArraySet1(_668_v20, 2)
				(_nw105).ArraySet1(_668_v20, 3)
				(_nw105).ArraySet1(_668_v20, 4)
				(_nw105).ArraySet1(_668_v20, 5)
				(_nw105).ArraySet1(_668_v20, 6)
				(_nw105).ArraySet1(_668_v20, 7)
				(_nw105).ArraySet1(_668_v20, 8)
				(_nw105).ArraySet1(_668_v20, 9)
				(_nw105).ArraySet1(_668_v20, 10)
				(_nw105).ArraySet1(_668_v20, 11)
				(_nw105).ArraySet1(_668_v20, 12)
				(_nw105).ArraySet1(_668_v20, 13)
				(_nw105).ArraySet1(_668_v20, 14)
				(_nw105).ArraySet1(_668_v20, 15)
				(_nw105).ArraySet1((_669_v21).Dtor_cf40(), 16)
				(_nw105).ArraySet1(_668_v20, 17)
				(_nw105).ArraySet1(_668_v20, 18)
				_670_v22 = _nw105
				var _671_v23 _dafny.Array
				_ = _671_v23
				var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
				_ = _nw106
				_671_v23 = _nw106
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_671_v23), 0))
				_ = _index131
				(_671_v23).ArraySet1(_this.F29(), (_index131).Int())
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_645_v0), 0))
				_ = _index132
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_671_v23), 0))
				_ = _index133
				var _rhs137 _dafny.Int = _dafny.IntOfInt64(938)
				_ = _rhs137
				var _rhs138 _dafny.Array = _670_v22
				_ = _rhs138
				var _rhs139 _dafny.Sequence = _this.F29()
				_ = _rhs139
				var _rhs140 bool = (_this).F28()
				_ = _rhs140
				var _lhs126 _dafny.Array = _645_v0
				_ = _lhs126
				var _lhs127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_645_v0), 0))
				_ = _lhs127
				var _lhs128 _dafny.Array = _671_v23
				_ = _lhs128
				var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_671_v23), 0))
				_ = _lhs129
				var _lhs130 *GlobalState = globalState
				_ = _lhs130
				(_lhs126).ArraySet1(_rhs137, (_lhs127).Int())
				_670_v22 = _rhs138
				(_lhs128).ArraySet1(_rhs139, (_lhs129).Int())
				_lhs130.F21 = _rhs140
				var _672_v24 D1
				_ = _672_v24
				_672_v24 = Companion_D1_.Create_DC5_(_658_i2, (_667_v19).F32(), _dafny.IntOfInt64(284), (_this).F23())
				(globalState).F9 = (_672_v24).Dtor_cf10()
				(globalState).F11 = _664_v16
				var _673_v25 _dafny.Map
				_ = _673_v25
				_673_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_667_v19).F32())
				var _674_v26 _dafny.Map
				_ = _674_v26
				_674_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_673_v25).Cardinality())
				var _pat_let_tv30 = _655_v8
				_ = _pat_let_tv30
				var _pat_let_tv31 = _645_v0
				_ = _pat_let_tv31
				var _pat_let_tv32 = _645_v0
				_ = _pat_let_tv32
				var _pat_let_tv33 = _655_v8
				_ = _pat_let_tv33
				var _pat_let_tv34 = _664_v16
				_ = _pat_let_tv34
				(globalState).F14 = (_667_v19).Fm18((func() _dafny.Map {
					if (_this).F28() {
						return _674_v26
					}
					return Companion_Default___.Fm25((_dafny.MultiSetOf(_664_v16)).Cardinality(), (_667_v19).F32(), globalState)
				})(), _dafny.Companion_Sequence_.Concatenate((func(_pat_let8_0 D0) D0 {
					return func(_675_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let9_0 _dafny.Sequence) D0 {
							return func(_676_dt__update_hcf0_h0 _dafny.Sequence) D0 {
								return Companion_D0_.Create_DC0_(_676_dt__update_hcf0_h0)
							}(_pat_let9_0)
						}(_dafny.Companion_Sequence_.Update(_pat_let_tv30, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_pat_let_tv32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_pat_let_tv31), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_pat_let_tv33).Cardinality()))).Uint32(), _pat_let_tv34))
					}(_pat_let8_0)
				}(Companion_D0_.Create_DC0_(_dafny.UnicodeSeqOfUtf8Bytes("lvnjskw")))).Dtor_cf0(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(387))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}((func(_677_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_678_i3 _dafny.Int) _dafny.CodePoint {
						return _677_v16
					}
				})(_664_v16)))), (_667_v19).F32(), globalState)
				(globalState).F7 = (_667_v19).F32()
			} else {
				var _679_v27 _dafny.Map
				_ = _679_v27
				_679_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _658_i2)
				var _680_v28 _dafny.Map
				_ = _680_v28
				_680_v28 = _679_v27
				_680_v28 = _679_v27
				(_this).F29_set_(_dafny.SeqOf((_667_v19).Fm18(_679_v27, _dafny.UnicodeSeqOfUtf8Bytes("epl"), (_659_v11).Select((Companion_Default___.SafeIndex(_658_i2, _dafny.IntOfUint32((_659_v11).Cardinality()))).Uint32()).(bool), globalState), Companion_Default___.Fm0(_658_i2, (_667_v19).F32(), globalState), (_this).F23()))
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_645_v0), 0))
				_ = _index134
				(_645_v0).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pnttsu")).Cardinality())), (_index134).Int())
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_645_v0), 0))
				_ = _index135
				var _rhs141 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_655_v8).Cardinality()), (_this).F23())
				_ = _rhs141
				var _rhs142 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F23(), Companion_Default___.SafeDivisionInt((_this).F23(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bqnxnel")).Cardinality())))
				_ = _rhs142
				var _rhs143 _dafny.Int = (_this).F23()
				_ = _rhs143
				var _rhs144 _dafny.Int = (func() _dafny.Int {
					if Companion_Default___.Fm1(((_667_v19).F31()).Cardinality(), (_667_v19).F32(), globalState) {
						return Companion_Default___.SafeDivisionInt((_this).F23(), _658_i2)
					}
					return _658_i2
				})()
				_ = _rhs144
				var _lhs131 *GlobalState = globalState
				_ = _lhs131
				var _lhs132 _dafny.Array = _645_v0
				_ = _lhs132
				var _lhs133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_645_v0), 0))
				_ = _lhs133
				var _lhs134 *GlobalState = globalState
				_ = _lhs134
				var _lhs135 *GlobalState = globalState
				_ = _lhs135
				_lhs131.F8 = _rhs141
				(_lhs132).ArraySet1(_rhs142, (_lhs133).Int())
				_lhs134.F14 = _rhs143
				_lhs135.F14 = _rhs144
				var _681_v29 _dafny.Array
				_ = _681_v29
				var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
				_ = _nw107
				_681_v29 = _nw107
				var _682_v30 _dafny.Map
				_ = _682_v30
				_682_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_655_v8, (_667_v19).F32())
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_681_v29), 0))
				_ = _index136
				(_681_v29).ArraySet1((func() _dafny.Map {
					if (_667_v19).F32() {
						return _682_v30
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("pinimjpwv"), (_this).F28())
				})(), (_index136).Int())
				var _683_v31 _dafny.Array
				_ = _683_v31
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_16
				var _nw108 _dafny.Array
				_ = _nw108
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw108 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.Int = (func(_684_v0 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_685_i4 _dafny.Int) _dafny.Int {
							return (_685_i4).Times((_684_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_684_v0), 0))).Int()).(_dafny.Int))
						}
					})(_645_v0)
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw108 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw108).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw108).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_683_v31 = _nw108
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_681_v29), 0))
				_ = _index137
				var _rhs145 _dafny.Map = (_682_v30).Update(_655_v8, !((_667_v19).F32()))
				_ = _rhs145
				var _rhs146 _dafny.Array = _683_v31
				_ = _rhs146
				var _rhs147 _dafny.Int = Companion_Default___.SafeDivisionInt(_658_i2, (_dafny.Zero).Minus((_645_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_645_v0), 0))).Int()).(_dafny.Int)))
				_ = _rhs147
				var _lhs136 _dafny.Array = _681_v29
				_ = _lhs136
				var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_681_v29), 0))
				_ = _lhs137
				var _lhs138 *GlobalState = globalState
				_ = _lhs138
				(_lhs136).ArraySet1(_rhs145, (_lhs137).Int())
				_683_v31 = _rhs146
				_lhs138.F14 = _rhs147
				var _686_v32 _dafny.Map
				_ = _686_v32
				_686_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_dafny.IntOfInt64(564), (_this).F28(), globalState), _664_v16)
				var _687_v33 _dafny.Map
				_ = _687_v33
				_687_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_667_v19).F32(), _664_v16)
				var _688_v34 _dafny.MultiSet
				_ = _688_v34
				_688_v34 = _dafny.MultiSetOf((_686_v32).Merge(_686_v32), (_687_v33), (func() _dafny.Map {
					if Companion_Default___.Fm1(_658_i2, (_667_v19).F32(), globalState) {
						return _686_v32
					}
					return _686_v32
				})())
				_688_v34 = _688_v34
			}
		}
		(globalState).F14 = (_this).F23()
		(globalState).F14 = (_this).F23()
		(globalState).F9 = (func() bool {
			if (_this).F28() {
				return (_this).F28()
			}
			return (_this).F28()
		})()
	}
}
func (_this *C3) F30() _dafny.Sequence {
	{
		return _this._f30
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f23 _dafny.Int
	_f27 _dafny.Array
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f23 = _dafny.Zero
	_this._f27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F23() _dafny.Int {
	return _this._f23
}
func (_this *C4) Ctor__(f27 _dafny.Array, f23 _dafny.Int) {
	{
		(_this)._f27 = f27
		(_this)._f23 = f23
	}
}
func (_this *C4) Fm10(p0 _dafny.Sequence, p1 bool, p2 _dafny.Map, p3 bool, globalState *GlobalState) bool {
	{
		return ((_dafny.SetOf(false)).Difference(_dafny.SetOf(true, !(true), false))).IsSubsetOf((_dafny.SetOf(false, false)).Difference(_dafny.SetOf(true, false)))
	}
}
func (_this *C4) Fm11(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), (func() _dafny.Sequence {
			if true {
				return _dafny.SeqOf(false)
			}
			return _dafny.SeqOf(false)
		})())
	}
}
func (_this *C4) M1(p0 _dafny.Array, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Set) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Set = _dafny.EmptySet
		_ = r1
		var _689_v0 _dafny.Array
		_ = _689_v0
		var _nw109 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(15))
		_ = _nw109
		_689_v0 = _nw109
		var _690_v1 _dafny.Sequence
		_ = _690_v1
		_690_v1 = _dafny.SeqOf(_dafny.MultiSetOf(_689_v0))
		var _691_v2 _dafny.Sequence
		_ = _691_v2
		_691_v2 = _dafny.UnicodeSeqOfUtf8Bytes("ru")
		var _692_v3 _dafny.Map
		_ = _692_v3
		_692_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_690_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_691_v2).Cardinality()), _dafny.IntOfUint32((_690_v1).Cardinality()))).Uint32()).(_dafny.MultiSet))
		var _693_v4 _dafny.MultiSet
		_ = _693_v4
		_693_v4 = _dafny.MultiSetOf(_689_v0)
		var _694_v5 D3
		_ = _694_v5
		_694_v5 = Companion_D3_.Create_DC9_(_693_v4)
		var _rhs148 _dafny.Int = ((func() _dafny.MultiSet {
			if (_692_v3).Contains(p1) {
				return (_692_v3).Get(p1).(_dafny.MultiSet)
			}
			return (_694_v5).Dtor_cf20()
		})()).Cardinality()
		_ = _rhs148
		var _rhs149 _dafny.Int = (_this).F23()
		_ = _rhs149
		var _rhs150 _dafny.Int = (_this).F23()
		_ = _rhs150
		var _rhs151 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F23(), (func() _dafny.Int {
			if p1 {
				return Companion_Default___.Fm0((_this).F23(), p1, globalState)
			}
			return (_this).F23()
		})())
		_ = _rhs151
		var _rhs152 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("taruujebt"), _691_v2), _691_v2)
		_ = _rhs152
		var _lhs139 *GlobalState = globalState
		_ = _lhs139
		var _lhs140 *GlobalState = globalState
		_ = _lhs140
		var _lhs141 *GlobalState = globalState
		_ = _lhs141
		r0 = _rhs148
		_lhs139.F8 = _rhs149
		r0 = _rhs150
		_lhs140.F14 = _rhs151
		_lhs141.F6 = _rhs152
		var _695_v6 _dafny.Map
		_ = _695_v6
		_695_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F23())
		var _696_v7 _dafny.Map
		_ = _696_v7
		_696_v7 = _695_v6
		var _697_v8 _dafny.MultiSet
		_ = _697_v8
		_697_v8 = _dafny.MultiSetOf(_696_v7, _696_v7, _696_v7)
		var _698_v9 _dafny.Sequence
		_ = _698_v9
		_698_v9 = _dafny.SeqOf((_697_v8).Cardinality())
		var _699_v10 *C1
		_ = _699_v10
		var _nw110 *C1 = New_C1_()
		_ = _nw110
		_nw110.Ctor__(Companion_Default___.Fm1((_this).F23(), true, globalState), _dafny.Companion_Sequence_.Concatenate(_698_v9, _698_v9), (_this).F23())
		_699_v10 = _nw110
		var _700_v11 _dafny.Array
		_ = _700_v11
		var _len0_17 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_17
		var _nw111 _dafny.Array
		_ = _nw111
		if _len0_17.Cmp(_dafny.Zero) == 0 {
			_nw111 = _dafny.NewArray(_len0_17)
		} else {
			var _init17 func(_dafny.Int) _dafny.Sequence = (func(_701_v9 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_702_i0 _dafny.Int) _dafny.Sequence {
					return _701_v9
				}
			})(_698_v9)
			_ = _init17
			var _element0_17 = _init17(_dafny.Zero)
			_ = _element0_17
			_nw111 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
			(_nw111).ArraySet1(_element0_17, 0)
			var _nativeLen0_17 = (_len0_17).Int()
			_ = _nativeLen0_17
			for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
				(_nw111).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
			}
		}
		_700_v11 = _nw111
		var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_700_v11), 0))
		_ = _index138
		(_700_v11).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F23()), _698_v9), (_index138).Int())
		var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))
		_ = _index139
		(p0).ArraySet1(((_this).F23()).Cmp((_this).F23()) <= 0, (_index139).Int())
		var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_700_v11), 0))
		_ = _index140
		var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))
		_ = _index141
		var _rhs153 _dafny.Sequence = (Companion_D9_.Create_DC27_(p1, _698_v9, p0, p1, p1)).Dtor_cf45()
		_ = _rhs153
		var _rhs154 bool = p1
		_ = _rhs154
		var _rhs155 bool = p1
		_ = _rhs155
		var _rhs156 bool = p1
		_ = _rhs156
		var _rhs157 bool = ((_this).F23()).Cmp(((_this).F23()).Plus((_this).F23())) > 0
		_ = _rhs157
		var _lhs142 _dafny.Array = _700_v11
		_ = _lhs142
		var _lhs143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_700_v11), 0))
		_ = _lhs143
		var _lhs144 *GlobalState = globalState
		_ = _lhs144
		var _lhs145 *GlobalState = globalState
		_ = _lhs145
		var _lhs146 _dafny.Array = p0
		_ = _lhs146
		var _lhs147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))
		_ = _lhs147
		var _lhs148 *GlobalState = globalState
		_ = _lhs148
		(_lhs142).ArraySet1(_rhs153, (_lhs143).Int())
		_lhs144.F7 = _rhs154
		_lhs145.F21 = _rhs155
		(_lhs146).ArraySet1(_rhs156, (_lhs147).Int())
		_lhs148.F21 = _rhs157
		var _703_i1 _dafny.Int
		_ = _703_i1
		_703_i1 = _dafny.Zero
		{
			for (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
				{
					if (_703_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_703_i1 = (_703_i1).Plus(_dafny.One)
					(globalState).F8 = (_this).F23()
					var _704_v12 _dafny.Array
					_ = _704_v12
					var _nw112 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(11))
					_ = _nw112
					_704_v12 = _nw112
					var _705_v13 D3
					_ = _705_v13
					_705_v13 = Companion_D3_.Create_DC10_((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool))
					var _rhs158 _dafny.Array = _704_v12
					_ = _rhs158
					var _rhs159 _dafny.Int = (_this).F23()
					_ = _rhs159
					var _rhs160 D3 = func(_pat_let10_0 D3) D3 {
						return func(_706_dt__update__tmp_h0 D3) D3 {
							return func(_pat_let11_0 bool) D3 {
								return func(_707_dt__update_hcf21_h0 bool) D3 {
									return Companion_D3_.Create_DC10_(_707_dt__update_hcf21_h0)
								}(_pat_let11_0)
							}((_dafny.IntOfInt64(-625)).Cmp(_dafny.IntOfInt64(856)) > 0)
						}(_pat_let10_0)
					}(Companion_D3_.Create_DC10_((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool)))
					_ = _rhs160
					var _lhs149 *GlobalState = globalState
					_ = _lhs149
					_704_v12 = _rhs158
					_lhs149.F8 = _rhs159
					_705_v13 = _rhs160
					var _708_v14 _dafny.Set
					_ = _708_v14
					_708_v14 = _dafny.SetOf((_this).F27())
					var _709_v15 _dafny.Map
					_ = _709_v15
					_709_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_708_v14, p1)
					(globalState).F21 = (func() bool {
						if (_709_v15).Contains(_708_v14) {
							return (_709_v15).Get(_708_v14).(bool)
						}
						return true
					})()
					var _710_v16 _dafny.Array
					_ = _710_v16
					var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
					_ = _nw113
					_710_v16 = _nw113
					_710_v16 = _710_v16
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _711_v17 _dafny.Array
		_ = _711_v17
		var _len0_18 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_18
		var _nw114 _dafny.Array
		_ = _nw114
		if _len0_18.Cmp(_dafny.Zero) == 0 {
			_nw114 = _dafny.NewArray(_len0_18)
		} else {
			var _init18 func(_dafny.Int) _dafny.Map = (func(_712_p1 bool, _713_p0 _dafny.Array) func(_dafny.Int) _dafny.Map {
				return func(_714_i2 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false, _712_p1, (_713_p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_713_p0), 0))).Int()).(bool)), (_this).F23())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(!((_713_p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_713_p0), 0))).Int()).(bool)), (_713_p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_713_p0), 0))).Int()).(bool), true, (_713_p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_713_p0), 0))).Int()).(bool), _712_p1), (_this).F23()))
				}
			})(p1, p0)
			_ = _init18
			var _element0_18 = _init18(_dafny.Zero)
			_ = _element0_18
			_nw114 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
			(_nw114).ArraySet1(_element0_18, 0)
			var _nativeLen0_18 = (_len0_18).Int()
			_ = _nativeLen0_18
			for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
				(_nw114).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
			}
		}
		_711_v17 = _nw114
		var _715_v18 _dafny.Set
		_ = _715_v18
		_715_v18 = _dafny.SetOf(!((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool)))
		var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_711_v17), 0))
		_ = _index142
		(_711_v17).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_715_v18, (_this).F23()), (_index142).Int())
		var _716_v20 _dafny.Map
		_ = _716_v20
		_716_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_715_v18, _698_v9)
		var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_711_v17), 0))
		_ = _index143
		(_711_v17).ArraySet1(func() _dafny.Map {
			var _coll25 = _dafny.NewMapBuilder()
			_ = _coll25
			for _iter29 := _dafny.Iterate((_716_v20).Keys().Elements()); ; {
				_compr_25, _ok29 := _iter29()
				if !_ok29 {
					break
				}
				var _717_v19 _dafny.Set
				_717_v19 = interface{}(_compr_25).(_dafny.Set)
				if (_716_v20).Contains(_717_v19) {
					_coll25.Add(_717_v19, (_this).F23())
				}
			}
			return _coll25.ToMap()
		}(), (_index143).Int())
		if (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
			var _718_v21 _dafny.Map
			_ = _718_v21
			_718_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F23())
			r0 = Companion_Default___.Fm0((func() _dafny.Int {
				if (_718_v21).Contains(_dafny.IntOfUint32((_691_v2).Cardinality())) {
					return (_718_v21).Get(_dafny.IntOfUint32((_691_v2).Cardinality())).(_dafny.Int)
				}
				return (_this).F23()
			})(), false, globalState)
			(globalState).F9 = (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool)
			var _719_v22 _dafny.MultiSet
			_ = _719_v22
			_719_v22 = _dafny.MultiSetOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), !(p1), !(p1))
			var _720_v23 _dafny.Sequence
			_ = _720_v23
			_720_v23 = _dafny.SeqOf(_719_v22)
			var _721_v24 T0
			_ = _721_v24
			var _nw115 *C3 = New_C3_()
			_ = _nw115
			_nw115.Ctor__(_720_v23, Companion_Default___.Fm0((_this).F23(), p1, globalState), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), (_700_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_700_v11), 0))).Int()).(_dafny.Sequence))
			_721_v24 = _nw115
			var _722_v25 _dafny.Map
			_ = _722_v25
			_722_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_721_v24, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool))
			var _723_v26 *C3
			_ = _723_v26
			var _nw116 *C3 = New_C3_()
			_ = _nw116
			_nw116.Ctor__((func() _dafny.Sequence {
				if (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
					return _720_v23
				}
				return _720_v23
			})(), ((_dafny.Zero).Minus((_this).F23())).Minus((_722_v25).Cardinality()), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), (_700_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_700_v11), 0))).Int()).(_dafny.Sequence))
			_723_v26 = _nw116
			var _724_v29 _dafny.Sequence
			_ = _724_v29
			_724_v29 = _dafny.SeqOf(_718_v21, _718_v21)
			var _725_v30 _dafny.Set
			_ = _725_v30
			_725_v30 = _dafny.SetOf((_this).F23(), (_718_v21).Cardinality())
			var _726_v31 _dafny.Array
			_ = _726_v31
			var _nwElement0_20 _dafny.Map = _718_v21
			_ = _nwElement0_20
			var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(16))
			_ = _nw117
			(_nw117).ArraySet1(_nwElement0_20, 0)
			(_nw117).ArraySet1(_718_v21, 1)
			(_nw117).ArraySet1(_718_v21, 2)
			(_nw117).ArraySet1(_718_v21, 3)
			(_nw117).ArraySet1((func() _dafny.Map {
				if Companion_Default___.Fm1((func() _dafny.Int {
					if (_719_v22).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool)) {
						return (_719_v22).Multiplicity((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool))
					}
					return _dafny.IntOfInt64(651)
				})(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState) {
					return _718_v21
				}
				return _718_v21
			})(), 4)
			(_nw117).ArraySet1(_718_v21, 5)
			(_nw117).ArraySet1(_718_v21, 6)
			(_nw117).ArraySet1(_718_v21, 7)
			(_nw117).ArraySet1(_718_v21, 8)
			(_nw117).ArraySet1(_718_v21, 9)
			(_nw117).ArraySet1((_718_v21).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_718_v21).Cardinality(), (_721_v24).F23())), 10)
			(_nw117).ArraySet1(func() _dafny.Map {
				var _coll26 = _dafny.NewMapBuilder()
				_ = _coll26
				for _iter30 := _dafny.Iterate((func() _dafny.Map {
					var _coll27 = _dafny.NewMapBuilder()
					_ = _coll27
					for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(118), _dafny.IntOfInt64(492))); ; {
						_compr_27, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _727_v28 _dafny.Int
						_727_v28 = interface{}(_compr_27).(_dafny.Int)
						if ((_dafny.IntOfInt64(118)).Cmp(_727_v28) <= 0) && ((_727_v28).Cmp(_dafny.IntOfInt64(492)) < 0) {
							_coll27.Add(Companion_Default___.SafeDivisionInt(_727_v28, (_dafny.SetOf(Companion_D3_.Create_DC10_(!(p1)))).Cardinality()), _dafny.MultiSetOf((_715_v18).Cardinality()))
						}
					}
					return _coll27.ToMap()
				}()).Keys().Elements()); ; {
					_compr_26, _ok30 := _iter30()
					if !_ok30 {
						break
					}
					var _728_v27 _dafny.Int
					_728_v27 = interface{}(_compr_26).(_dafny.Int)
					if (func() _dafny.Map {
						var _coll28 = _dafny.NewMapBuilder()
						_ = _coll28
						for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(118), _dafny.IntOfInt64(492))); ; {
							_compr_28, _ok32 := _iter32()
							if !_ok32 {
								break
							}
							var _727_v28 _dafny.Int
							_727_v28 = interface{}(_compr_28).(_dafny.Int)
							if ((_dafny.IntOfInt64(118)).Cmp(_727_v28) <= 0) && ((_727_v28).Cmp(_dafny.IntOfInt64(492)) < 0) {
								_coll28.Add(Companion_Default___.SafeDivisionInt(_727_v28, (_dafny.SetOf(Companion_D3_.Create_DC10_(!(p1)))).Cardinality()), _dafny.MultiSetOf((_715_v18).Cardinality()))
							}
						}
						return _coll28.ToMap()
					}()).Contains(_728_v27) {
						_coll26.Add((_728_v27).Minus((_721_v24).F23()), (_721_v24).F23())
					}
				}
				return _coll26.ToMap()
			}(), 11)
			(_nw117).ArraySet1((_724_v29).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_691_v2).Cardinality()), _dafny.IntOfUint32((_724_v29).Cardinality()))).Uint32()).(_dafny.Map), 12)
			(_nw117).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_721_v24).F23()), (_721_v24).F23())).Update((_725_v30).Cardinality(), (_this).F23()), 13)
			(_nw117).ArraySet1(_718_v21, 14)
			(_nw117).ArraySet1(_718_v21, 15)
			_726_v31 = _nw117
			var _729_v32 _dafny.Sequence
			_ = _729_v32
			_729_v32 = _dafny.SeqOf((_this).Fm10(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(868))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}(func(_730_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			})), false, _695_v6, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState), p1, p1, p1, !((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool)))
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index144
			((_this).F27()).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm0(((_700_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_700_v11), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_729_v32).Cardinality()), _dafny.IntOfUint32(((_700_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_700_v11), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Int), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState)), (_index144).Int())
			var _731_v33 D11
			_ = _731_v33
			_731_v33 = Companion_D11_.Create_DC30_(_726_v31)
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index145
			var _rhs161 _dafny.Int = (_721_v24).F23()
			_ = _rhs161
			var _rhs162 _dafny.Array = (_731_v33).Dtor_cf51()
			_ = _rhs162
			var _rhs163 _dafny.Int = (_this).F23()
			_ = _rhs163
			var _rhs164 _dafny.Int = (_this).F23()
			_ = _rhs164
			var _lhs150 *GlobalState = globalState
			_ = _lhs150
			var _lhs151 _dafny.Array = (_this).F27()
			_ = _lhs151
			var _lhs152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _lhs152
			var _lhs153 *GlobalState = globalState
			_ = _lhs153
			_lhs150.F14 = _rhs161
			_726_v31 = _rhs162
			(_lhs151).ArraySet1(_rhs163, (_lhs152).Int())
			_lhs153.F14 = _rhs164
			(globalState).F14 = _dafny.IntOfUint32((_698_v9).Cardinality())
		} else {
			var _732_v34 _dafny.Map
			_ = _732_v34
			_732_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_this).F23(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState), Companion_Default___.Fm6((_dafny.Zero).Minus((_this).F23()), globalState))
			(globalState).F8 = (_dafny.Zero).Minus((_732_v34).Cardinality())
			var _733_v35 _dafny.Map
			_ = _733_v35
			_733_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F23())
			var _734_v36 *C2
			_ = _734_v36
			var _nw118 *C2 = New_C2_()
			_ = _nw118
			_nw118.Ctor__(_733_v35, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), (_dafny.IntOfUint32((_691_v2).Cardinality())).Cmp((_dafny.Zero).Minus((_this).F23())) != 0, _dafny.Companion_Sequence_.Concatenate(_698_v9, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(397))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg48 _dafny.Int) interface{} {
					return coer48(arg48)
				}
			}(func(_735_i4 _dafny.Int) _dafny.Int {
				return (_this).F23()
			}))), (_this).F23())
			_734_v36 = _nw118
			var _736_v37 _dafny.Array
			_ = _736_v37
			var _nw119 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw119
			_736_v37 = _nw119
			_736_v37 = p0
			var _737_v38 _dafny.Set
			_ = _737_v38
			_737_v38 = _dafny.SetOf(_734_v36)
			var _738_v39 _dafny.Sequence
			_ = _738_v39
			_738_v39 = _dafny.SeqOf(p1, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), Companion_Default___.Fm1((_this).F23(), (_734_v36).F32(), globalState), p1, (_737_v38).IsDisjointFrom(_737_v38))
			if (_738_v39).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F23()), _dafny.IntOfUint32((_738_v39).Cardinality()))).Uint32()).(bool) {
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))
				_ = _index146
				(p0).ArraySet1(((_this).F23()).Cmp((_this).F23()) < 0, (_index146).Int())
				var _739_v40 _dafny.Array
				_ = _739_v40
				var _nwElement0_21 _dafny.Map = _695_v6
				_ = _nwElement0_21
				var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(24))
				_ = _nw120
				(_nw120).ArraySet1(_nwElement0_21, 0)
				(_nw120).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_734_v36).F32(), (_this).F23()), 1)
				(_nw120).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), (_this).F23()), 2)
				(_nw120).ArraySet1((_695_v6).Merge(_695_v6), 3)
				(_nw120).ArraySet1(_695_v6, 4)
				(_nw120).ArraySet1((_695_v6).Update((_734_v36).F32(), (_this).F23()), 5)
				(_nw120).ArraySet1(_695_v6, 6)
				(_nw120).ArraySet1(_695_v6, 7)
				(_nw120).ArraySet1((func() _dafny.Map {
					if (_734_v36).F32() {
						return _695_v6
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_734_v36).F32(), (_this).F23())
				})(), 8)
				(_nw120).ArraySet1(_695_v6, 9)
				(_nw120).ArraySet1((_695_v6).Merge(_695_v6), 10)
				(_nw120).ArraySet1((_695_v6).Merge(_695_v6), 11)
				(_nw120).ArraySet1(((_695_v6).Update(false, (_this).F23())).Merge(_695_v6), 12)
				(_nw120).ArraySet1(_695_v6, 13)
				(_nw120).ArraySet1((_695_v6).Merge(_695_v6), 14)
				(_nw120).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F23())).Update((_734_v36).F32(), _dafny.IntOfUint32((_691_v2).Cardinality())), 15)
				(_nw120).ArraySet1(_695_v6, 16)
				(_nw120).ArraySet1(_695_v6, 17)
				(_nw120).ArraySet1(_695_v6, 18)
				(_nw120).ArraySet1((_695_v6).Update((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), (_695_v6).Cardinality()), 19)
				(_nw120).ArraySet1(_695_v6, 20)
				(_nw120).ArraySet1(_695_v6, 21)
				(_nw120).ArraySet1(_695_v6, 22)
				(_nw120).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), (_this).F23()), 23)
				_739_v40 = _nw120
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_739_v40), 0))
				_ = _index147
				(_739_v40).ArraySet1(_695_v6, (_index147).Int())
				var _740_v41 _dafny.CodePoint
				_ = _740_v41
				_740_v41 = _dafny.CodePoint('v')
				var _741_v42 _dafny.Map
				_ = _741_v42
				_741_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_740_v41, p1)
				var _742_v43 D12
				_ = _742_v43
				_742_v43 = Companion_D12_.Create_DC34_(_741_v42)
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_739_v40), 0))
				_ = _index148
				(_739_v40).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, ((_742_v43).Dtor_cf55()).Cardinality())).Merge(_695_v6), (_index148).Int())
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index149
				((_this).F27()).ArraySet1((_this).F23(), (_index149).Int())
				var _743_v44 _dafny.Map
				_ = _743_v44
				_743_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool))
				var _744_v45 D8
				_ = _744_v45
				_744_v45 = Companion_D8_.Create_DC21_((_734_v36).F32(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), (_743_v44).Cardinality())
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index150
				((_this).F27()).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(798), ((_744_v45).Dtor_cf37()).Times((_this).F23())), (_index150).Int())
				var _745_v46 bool
				_ = _745_v46
				var _746_v47 _dafny.Map
				_ = _746_v47
				var _747_v48 D1
				_ = _747_v48
				var _748_v49 _dafny.Array
				_ = _748_v49
				var _out30 bool
				_ = _out30
				var _out31 _dafny.Map
				_ = _out31
				var _out32 D1
				_ = _out32
				var _out33 _dafny.Array
				_ = _out33
				_out30, _out31, _out32, _out33 = (_734_v36).M6((_dafny.SetOf(true, true, (_734_v36).F32())).IsSubsetOf(_715_v18), globalState)
				_745_v46 = _out30
				_746_v47 = _out31
				_747_v48 = _out32
				_748_v49 = _out33
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))
				_ = _index151
				(p0).ArraySet1(p1, (_index151).Int())
			} else {
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))
				_ = _index152
				(p0).ArraySet1(((_this).F23()).Cmp((_dafny.IntOfInt64(33)).Times(_dafny.IntOfInt64(127))) > 0, (_index152).Int())
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index153
				((_this).F27()).ArraySet1((_this).F23(), (_index153).Int())
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index154
				((_this).F27()).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F23(), Companion_Default___.SafeModuloInt((_this).F23(), (_this).F23())), (_index154).Int())
				var _749_v50 _dafny.CodePoint
				_ = _749_v50
				_749_v50 = _dafny.CodePoint('n')
				(globalState).F21 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_691_v2, (Companion_Default___.SafeIndex(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_691_v2).Cardinality()))).Uint32(), _749_v50), _dafny.Companion_Sequence_.Concatenate(_691_v2, _691_v2))
				var _750_v51 D3
				_ = _750_v51
				_750_v51 = Companion_D3_.Create_DC11_(_dafny.IntOfInt64(453))
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))
				_ = _index155
				(p0).ArraySet1((_738_v39).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((func() _dafny.Int {
					if (_734_v36).F32() {
						return ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int)
					}
					return (_750_v51).Dtor_cf22()
				})()), _dafny.IntOfUint32((_738_v39).Cardinality()))).Uint32()).(bool), (_index155).Int())
				(globalState).F21 = false
			}
			var _751_v52 _dafny.Set
			_ = _751_v52
			_751_v52 = _dafny.SetOf((_this).F23())
			var _752_v53 *C2
			_ = _752_v53
			var _nw121 *C2 = New_C2_()
			_ = _nw121
			_nw121.Ctor__(((_734_v36).F31()).Update((_this).F23(), _dafny.IntOfInt64(797)), false, (_this).Fm10(_691_v2, p1, (_695_v6).Update((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((p0), 0))).Int()).(bool), (_this).F23()), Companion_Default___.Fm1((_dafny.Zero).Minus((_this).F23()), (_734_v36).F32(), globalState), globalState), _698_v9, Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_751_v52).Cardinality()), (_dafny.Zero).Minus((_this).F23())))
			_752_v53 = _nw121
		}
		r0 = (_this).F23()
		var _753_v54 _dafny.Set
		_ = _753_v54
		_753_v54 = _dafny.SetOf((_695_v6).Cardinality(), (_this).F23())
		r1 = (_753_v54).Union((_753_v54).Intersection(_753_v54))
		return r0, r1
	}
}
func (_this *C4) M2(p0 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		(globalState).F8 = (_this).F23()
		var _754_v0 bool
		_ = _754_v0
		_754_v0 = false
		var _755_i0 _dafny.Int
		_ = _755_i0
		_755_i0 = _dafny.Zero
		{
			for ((func() bool {
				if false {
					return (_this).Fm10(p0, true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v0, (_this).F23()), _754_v0, globalState)
				}
				return _754_v0
			})()) || (_754_v0) {
				{
					if (_755_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_755_i0 = (_755_i0).Plus(_dafny.One)
					var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))
					_ = _index156
					((_this).F27()).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F23(), (_this).F23()), (_index156).Int())
					var _756_v1 _dafny.Sequence
					_ = _756_v1
					_756_v1 = _dafny.SeqOf(_754_v0)
					var _757_v2 _dafny.Array
					_ = _757_v2
					var _nwElement0_22 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uovdktobe"), p0)).Cardinality())
					_ = _nwElement0_22
					var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(3))
					_ = _nw122
					(_nw122).ArraySet1(_nwElement0_22, 0)
					(_nw122).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_756_v1).Cardinality()), (_this).F23()), 1)
					(_nw122).ArraySet1((_this).F23(), 2)
					_757_v2 = _nw122
					var _758_v3 _dafny.Sequence
					_ = _758_v3
					_758_v3 = _dafny.SeqOf((_this).F23())
					var _759_v4 _dafny.Map
					_ = _759_v4
					_759_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _dafny.IntOfUint32((_758_v3).Cardinality()))
					var _760_v5 _dafny.Set
					_ = _760_v5
					_760_v5 = _dafny.SetOf(_754_v0)
					var _761_v6 _dafny.Map
					_ = _761_v6
					_761_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_754_v0), ((Companion_D13_.Create_DC38_(_760_v5)).Dtor_cf62()).Cardinality())
					var _762_v7 *C2
					_ = _762_v7
					var _nw123 *C2 = New_C2_()
					_ = _nw123
					_nw123.Ctor__(_759_v4, (_this).Fm10(p0, _754_v0, _761_v6, true, globalState), _754_v0, _758_v3, _dafny.IntOfInt64(465))
					_762_v7 = _nw123
					var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))
					_ = _index157
					var _rhs165 _dafny.Int = ((_this).F23()).Minus((_this).F23())
					_ = _rhs165
					var _rhs166 bool = _754_v0
					_ = _rhs166
					var _rhs167 _dafny.Array = (_this).F27()
					_ = _rhs167
					var _rhs168 *C2 = (func() *C2 {
						if _754_v0 {
							return _762_v7
						}
						return _762_v7
					})()
					_ = _rhs168
					var _lhs154 _dafny.Array = (_this).F27()
					_ = _lhs154
					var _lhs155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))
					_ = _lhs155
					var _lhs156 *GlobalState = globalState
					_ = _lhs156
					(_lhs154).ArraySet1(_rhs165, (_lhs155).Int())
					_lhs156.F21 = _rhs166
					_757_v2 = _rhs167
					_762_v7 = _rhs168
					var _763_v8 _dafny.CodePoint
					_ = _763_v8
					_763_v8 = _dafny.CodePoint('n')
					var _764_v9 _dafny.Set
					_ = _764_v9
					_764_v9 = _dafny.SetOf((Companion_D8_.Create_DC20_(_754_v0, ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int))).Dtor_cf34(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("dxmrtsdjn"), (Companion_Default___.SafeIndex(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dxmrtsdjn")).Cardinality()))).Uint32(), _763_v8), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("dxmrtsdjn"), (Companion_Default___.SafeIndex(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dxmrtsdjn")).Cardinality()))).Uint32(), _763_v8)).Cardinality()))).Uint32(), _763_v8)).Cardinality()))
					var _765_v10 _dafny.Map
					_ = _765_v10
					_765_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v0, _754_v0)
					var _766_v11 _dafny.Map
					_ = _766_v11
					_766_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_764_v9).Cardinality(), (func() bool {
						if (_765_v10).Contains((_762_v7).F32()) {
							return (_765_v10).Get((_762_v7).F32()).(bool)
						}
						return _754_v0
					})())
					var _767_v12 _dafny.Array
					_ = _767_v12
					var _nwElement0_23 bool = true
					_ = _nwElement0_23
					var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(12))
					_ = _nw124
					(_nw124).ArraySet1(_nwElement0_23, 0)
					(_nw124).ArraySet1((_762_v7).F32(), 1)
					(_nw124).ArraySet1(true, 2)
					(_nw124).ArraySet1((_762_v7).F32(), 3)
					(_nw124).ArraySet1(_754_v0, 4)
					(_nw124).ArraySet1(((_762_v7).F32()) || (_754_v0), 5)
					(_nw124).ArraySet1(!_dafny.Companion_Sequence_.Equal(_756_v1, _756_v1), 6)
					(_nw124).ArraySet1(_754_v0, 7)
					(_nw124).ArraySet1((_754_v0) == (_754_v0), 8)
					(_nw124).ArraySet1((func() bool {
						if (_762_v7).F32() {
							return (_762_v7).F32()
						}
						return (_762_v7).F32()
					})(), 9)
					(_nw124).ArraySet1(false, 10)
					(_nw124).ArraySet1((((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_754_v0, _754_v0, (_762_v7).F32(), (func() bool {
						if (_766_v11).Contains((_dafny.Zero).Minus((_this).F23())) {
							return (_766_v11).Get((_dafny.Zero).Minus((_this).F23())).(bool)
						}
						return (_762_v7).F32()
					})())).Cardinality())) == 0, 11)
					_767_v12 = _nw124
					_767_v12 = _767_v12
					var _768_v13 D9
					_ = _768_v13
					_768_v13 = Companion_D9_.Create_DC27_(((_764_v9).Cardinality()).Cmp((_this).F23()) != 0, _758_v3, _767_v12, _754_v0, (_762_v7).F32())
					var _769_v14 _dafny.Map
					_ = _769_v14
					_769_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_762_v7).F32(), _dafny.IntOfInt64(957))
					var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))
					_ = _index158
					var _rhs169 _dafny.Int = Companion_Default___.Fm0(_dafny.IntOfUint32((p0).Cardinality()), (func() bool {
						if (_766_v11).Contains(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int)) {
							return (_766_v11).Get(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int)).(bool)
						}
						return (func() bool {
							if (_765_v10).Contains((_762_v7).F32()) {
								return (_765_v10).Get((_762_v7).F32()).(bool)
							}
							return (_this).Fm10(p0, _754_v0, (_769_v14), _754_v0, globalState)
						})()
					})(), globalState)
					_ = _rhs169
					var _rhs170 D9 = _768_v13
					_ = _rhs170
					var _lhs157 _dafny.Array = (_this).F27()
					_ = _lhs157
					var _lhs158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))
					_ = _lhs158
					(_lhs157).ArraySet1(_rhs169, (_lhs158).Int())
					_768_v13 = _rhs170
					var _770_v15 _dafny.Map
					_ = _770_v15
					_770_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_756_v1, _756_v1), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_756_v1, _756_v1)).Cardinality()))).Uint32(), (_762_v7).F32()))
					var _rhs171 bool = ((_this).F23()).Cmp(((_dafny.MultiSetOf(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int))).Cardinality()).Plus(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int))) == 0
					_ = _rhs171
					var _rhs172 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int), _756_v1)).Merge(_770_v15)).Merge((_770_v15).Update(_dafny.IntOfUint32((_756_v1).Cardinality()), (_this).Fm11(globalState)))
					_ = _rhs172
					var _rhs173 _dafny.Int = (_dafny.Zero).Minus((_this).F23())
					_ = _rhs173
					var _lhs159 *GlobalState = globalState
					_ = _lhs159
					_lhs159.F9 = _rhs171
					_770_v15 = _rhs172
					r0 = _rhs173
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _771_v16 _dafny.Set
		_ = _771_v16
		_771_v16 = _dafny.SetOf(!(_754_v0), !(!(_754_v0)))
		(globalState).F8 = (_dafny.Zero).Minus(func(_source6 D0) _dafny.Int {
			if _source6.Is_DC1() {
				var _772___mcc_h0 _dafny.Int = _source6.Get_().(D0_DC1).Cf1
				_ = _772___mcc_h0
				var _773___mcc_h1 bool = _source6.Get_().(D0_DC1).Cf2
				_ = _773___mcc_h1
				var _774___mcc_h2 _dafny.Int = _source6.Get_().(D0_DC1).Cf3
				_ = _774___mcc_h2
				var _775_cf3 _dafny.Int = _774___mcc_h2
				_ = _775_cf3
				var _776_cf2 bool = _773___mcc_h1
				_ = _776_cf2
				var _777_cf1 _dafny.Int = _772___mcc_h0
				_ = _777_cf1
				return (_775_cf3).Minus(_777_cf1)
			} else if _source6.Is_DC0() {
				var _778___mcc_h3 _dafny.Sequence = _source6.Get_().(D0_DC0).Cf0
				_ = _778___mcc_h3
				var _779_cf0 _dafny.Sequence = _778___mcc_h3
				_ = _779_cf0
				return (_dafny.IntOfInt64(11)).Minus((_this).F23())
			} else {
				var _780___mcc_h4 D0 = _source6.Get_().(D0_DC2).Cf4
				_ = _780___mcc_h4
				var _781_cf4 D0 = _780___mcc_h4
				_ = _781_cf4
				return (Companion_D9_.Create_DC26_((_this).F23())).Dtor_cf43()
			}
		}(Companion_Default___.Fm26(_771_v16, globalState)))
		var _782_i1 _dafny.Int
		_ = _782_i1
		_782_i1 = _dafny.Zero
		{
			for !(_754_v0) {
				{
					if (_782_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_782_i1 = (_782_i1).Plus(_dafny.One)
					if _754_v0 {
						var _783_v17 _dafny.CodePoint
						_ = _783_v17
						_783_v17 = _dafny.CodePoint('s')
						var _784_v18 _dafny.MultiSet
						_ = _784_v18
						_784_v18 = _dafny.MultiSetOf(_754_v0, _754_v0, false, (_754_v0) || (Companion_Default___.Fm1(_dafny.IntOfUint32((_dafny.SeqOf(_783_v17)).Cardinality()), _754_v0, globalState)), ((_this).F23()).Cmp((_this).F23()) >= 0)
						_784_v18 = (_dafny.MultiSetOf(_754_v0, _754_v0)).Difference(_784_v18)
						var _785_v19 _dafny.Map
						_ = _785_v19
						_785_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F23()).Cmp((_this).F23()) == 0, _783_v17)
						var _786_v20 _dafny.Sequence
						_ = _786_v20
						_786_v20 = _dafny.SeqOf((_this).F23(), _dafny.IntOfInt64(419), (_this).F23())
						var _787_v21 _dafny.Map
						_ = _787_v21
						_787_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-174))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg49 _dafny.Int) interface{} {
								return coer49(arg49)
							}
						}(func(_788_i2 _dafny.Int) _dafny.Int {
							return (_this).F23()
						})))
						var _789_v22 D12
						_ = _789_v22
						_789_v22 = Companion_D12_.Create_DC36_(_754_v0, true, (_this).F23(), (_this).F23(), _787_v21)
						var _790_v23 _dafny.Array
						_ = _790_v23
						var _nwElement0_24 bool = (_dafny.IntOfInt64(597)).Cmp((_this).F23()) >= 0
						_ = _nwElement0_24
						var _nw125 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(26))
						_ = _nw125
						(_nw125).ArraySet1(_nwElement0_24, 0)
						(_nw125).ArraySet1(_754_v0, 1)
						(_nw125).ArraySet1(_754_v0, 2)
						(_nw125).ArraySet1(_754_v0, 3)
						(_nw125).ArraySet1(!(_754_v0) || (_754_v0), 4)
						(_nw125).ArraySet1(true, 5)
						(_nw125).ArraySet1(!(_754_v0) || (_754_v0), 6)
						(_nw125).ArraySet1(_754_v0, 7)
						(_nw125).ArraySet1(true, 8)
						(_nw125).ArraySet1(_754_v0, 9)
						(_nw125).ArraySet1(((_this).F23()).Cmp((_dafny.Zero).Minus((_this).F23())) > 0, 10)
						(_nw125).ArraySet1(_754_v0, 11)
						(_nw125).ArraySet1(!(_754_v0), 12)
						(_nw125).ArraySet1(_754_v0, 13)
						(_nw125).ArraySet1(_754_v0, 14)
						(_nw125).ArraySet1(!_dafny.Companion_Sequence_.Contains(p0, _783_v17), 15)
						(_nw125).ArraySet1(_754_v0, 16)
						(_nw125).ArraySet1(_754_v0, 17)
						(_nw125).ArraySet1(!(_dafny.Companion_Sequence_.Equal(_786_v20, _786_v20)), 18)
						(_nw125).ArraySet1((_754_v0) || (_754_v0), 19)
						(_nw125).ArraySet1((_789_v22).Dtor_cf57(), 20)
						(_nw125).ArraySet1(_754_v0, 21)
						(_nw125).ArraySet1(!(true), 22)
						(_nw125).ArraySet1(_754_v0, 23)
						(_nw125).ArraySet1(_754_v0, 24)
						(_nw125).ArraySet1(_754_v0, 25)
						_790_v23 = _nw125
						var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_790_v23), 0))
						_ = _index159
						(_790_v23).ArraySet1(!(_754_v0), (_index159).Int())
						var _791_v24 _dafny.Array
						_ = _791_v24
						var _len0_19 _dafny.Int = _dafny.IntOfInt64(6)
						_ = _len0_19
						var _nw126 _dafny.Array
						_ = _nw126
						if _len0_19.Cmp(_dafny.Zero) == 0 {
							_nw126 = _dafny.NewArray(_len0_19)
						} else {
							var _init19 func(_dafny.Int) _dafny.CodePoint = func(_792_i3 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('o')
							}
							_ = _init19
							var _element0_19 = _init19(_dafny.Zero)
							_ = _element0_19
							_nw126 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
							(_nw126).ArraySet1CodePoint(_element0_19, 0)
							var _nativeLen0_19 = (_len0_19).Int()
							_ = _nativeLen0_19
							for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
								(_nw126).ArraySet1CodePoint(_init19(_dafny.IntOf(_i0_19)), _i0_19)
							}
						}
						_791_v24 = _nw126
						var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_791_v24), 0))
						_ = _index160
						(_791_v24).ArraySet1CodePoint(_783_v17, (_index160).Int())
						var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_790_v23), 0))
						_ = _index161
						var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_791_v24), 0))
						_ = _index162
						var _rhs174 _dafny.Map = (_785_v19).Merge(_785_v19)
						_ = _rhs174
						var _rhs175 bool = !(_754_v0) || (!(_754_v0))
						_ = _rhs175
						var _rhs176 _dafny.Array = _790_v23
						_ = _rhs176
						var _rhs177 _dafny.CodePoint = _783_v17
						_ = _rhs177
						var _lhs160 _dafny.Array = _790_v23
						_ = _lhs160
						var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_790_v23), 0))
						_ = _lhs161
						var _lhs162 _dafny.Array = _791_v24
						_ = _lhs162
						var _lhs163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_791_v24), 0))
						_ = _lhs163
						_785_v19 = _rhs174
						(_lhs160).ArraySet1(_rhs175, (_lhs161).Int())
						_790_v23 = _rhs176
						(_lhs162).ArraySet1CodePoint(_rhs177, (_lhs163).Int())
						var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_790_v23), 0))
						_ = _index163
						(_790_v23).ArraySet1(!(Companion_Default___.Fm1((_this).F23(), !((_790_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_790_v23), 0))).Int()).(bool)), globalState)) || (_754_v0), (_index163).Int())
						var _793_v25 _dafny.Map
						_ = _793_v25
						_793_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v0, false)
						var _794_v26 _dafny.Sequence
						_ = _794_v26
						_794_v26 = _dafny.SeqOf(_784_v18, _784_v18, ((_784_v18).Update((_790_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_790_v23), 0))).Int()).(bool), Companion_Default___.Abs((_793_v25).Cardinality()))).Update((_790_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_790_v23), 0))).Int()).(bool), Companion_Default___.Abs(_dafny.IntOfInt64(678))), _784_v18, _784_v18)
						var _795_v27 *C3
						_ = _795_v27
						var _nw127 *C3 = New_C3_()
						_ = _nw127
						_nw127.Ctor__(_794_v26, (_this).F23(), !((_790_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_790_v23), 0))).Int()).(bool)), _786_v20)
						_795_v27 = _nw127
						(globalState).F21 = _754_v0
					} else {
						var _796_v28 _dafny.CodePoint
						_ = _796_v28
						_796_v28 = _dafny.CodePoint('x')
						(globalState).F11 = _796_v28
						var _797_v29 _dafny.Array
						_ = _797_v29
						var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
						_ = _nw128
						_797_v29 = _nw128
						_797_v29 = _797_v29
						var _798_v31 _dafny.Map
						_ = _798_v31
						_798_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v0, (_this).F23())
						var _799_v32 _dafny.Set
						_ = _799_v32
						_799_v32 = _dafny.SetOf(Companion_Default___.Fm0((_this).F23(), _754_v0, globalState), (_this).F23(), (_this).F23(), Companion_Default___.Fm0((_dafny.Zero).Minus((_this).F23()), (_this).Fm10(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(532))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg50 _dafny.Int) interface{} {
								return coer50(arg50)
							}
						}((func(_800_v28 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_801_i4 _dafny.Int) _dafny.CodePoint {
								return _800_v28
							}
						})(_796_v28))), _754_v0, _798_v31, _754_v0, globalState), globalState), (_this).F23())
						var _802_v33 _dafny.Array
						_ = _802_v33
						var _nwElement0_25 _dafny.Set = (func() _dafny.Set {
							var _coll29 = _dafny.NewBuilder()
							_ = _coll29
							for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(397), _dafny.IntOfInt64(-231))); ; {
								_compr_29, _ok33 := _iter33()
								if !_ok33 {
									break
								}
								var _803_v30 _dafny.Int
								_803_v30 = interface{}(_compr_29).(_dafny.Int)
								if ((_dafny.IntOfInt64(397)).Cmp(_803_v30) <= 0) && ((_803_v30).Cmp(_dafny.IntOfInt64(-231)) < 0) {
									_coll29.Add((_803_v30).Minus((_this).F23()))
								}
							}
							return _coll29.ToSet()
						}()).Union(_dafny.SetOf((_this).F23()))
						_ = _nwElement0_25
						var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(2))
						_ = _nw129
						(_nw129).ArraySet1(_nwElement0_25, 0)
						(_nw129).ArraySet1(_799_v32, 1)
						_802_v33 = _nw129
						var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_802_v33), 0))
						_ = _index164
						(_802_v33).ArraySet1(_799_v32, (_index164).Int())
						var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_802_v33), 0))
						_ = _index165
						(_802_v33).ArraySet1(_799_v32, (_index165).Int())
						var _804_v34 _dafny.Map
						_ = _804_v34
						_804_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(!(_754_v0))).Cardinality()), _754_v0)
						_804_v34 = (_804_v34).Merge(_804_v34)
						var _805_v35 _dafny.Map
						_ = _805_v35
						_805_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F23())
						_805_v35 = (_805_v35).Update(Companion_Default___.Fm0(_dafny.IntOfInt64(-241), !(_754_v0), globalState), (_this).F23())
					}
					var _806_v36 _dafny.Sequence
					_ = _806_v36
					_806_v36 = _dafny.SeqOf((_this).F23())
					var _807_v37 D8
					_ = _807_v37
					_807_v37 = Companion_D8_.Create_DC20_(_754_v0, (_this).F23())
					var _808_v38 _dafny.Sequence
					_ = _808_v38
					_808_v38 = _dafny.SeqOf(_754_v0, _754_v0, (_807_v37).Dtor_cf33(), _754_v0, _754_v0)
					var _809_v39 _dafny.Array
					_ = _809_v39
					var _nwElement0_26 bool = Companion_Default___.Fm1((_this).F23(), true, globalState)
					_ = _nwElement0_26
					var _nw130 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(18))
					_ = _nw130
					(_nw130).ArraySet1(_nwElement0_26, 0)
					(_nw130).ArraySet1(_754_v0, 1)
					(_nw130).ArraySet1((_771_v16).IsSubsetOf(_771_v16), 2)
					(_nw130).ArraySet1(_754_v0, 3)
					(_nw130).ArraySet1(!(_754_v0), 4)
					(_nw130).ArraySet1((_dafny.IntOfUint32((_806_v36).Cardinality())).Cmp((_this).F23()) >= 0, 5)
					(_nw130).ArraySet1(_754_v0, 6)
					(_nw130).ArraySet1((_808_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hm")).Cardinality()), _dafny.IntOfUint32((_808_v38).Cardinality()))).Uint32()).(bool), 7)
					(_nw130).ArraySet1(_754_v0, 8)
					(_nw130).ArraySet1(_754_v0, 9)
					(_nw130).ArraySet1(_754_v0, 10)
					(_nw130).ArraySet1(_754_v0, 11)
					(_nw130).ArraySet1(false, 12)
					(_nw130).ArraySet1(_754_v0, 13)
					(_nw130).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfUint32((p0).Cardinality()), _754_v0, globalState), 14)
					(_nw130).ArraySet1(!(!(!(_754_v0))), 15)
					(_nw130).ArraySet1(!(_754_v0) || (_754_v0), 16)
					(_nw130).ArraySet1(_754_v0, 17)
					_809_v39 = _nw130
					var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))
					_ = _index166
					(_809_v39).ArraySet1(_754_v0, (_index166).Int())
					var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))
					_ = _index167
					(_809_v39).ArraySet1(!(_754_v0), (_index167).Int())
					if _754_v0 {
						var _810_v40 _dafny.Map
						_ = _810_v40
						_810_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_806_v36, _754_v0)
						_810_v40 = ((_810_v40).Merge(_810_v40)).Merge(_810_v40)
						var _811_v41 _dafny.Map
						_ = _811_v41
						_811_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1((_this).F23(), (_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool), globalState), ((Companion_D0_.Create_DC1_(_dafny.IntOfInt64(967), (_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool), (_this).F23())).Dtor_cf3()).Cmp(_dafny.IntOfInt64(872)) <= 0)
						(globalState).F7 = (func() bool {
							if (_811_v41).Contains(((_dafny.Zero).Minus((_this).F23())).Cmp((_this).F23()) <= 0) {
								return (_811_v41).Get(((_dafny.Zero).Minus((_this).F23())).Cmp((_this).F23()) <= 0).(bool)
							}
							return ((_this).F23()).Cmp(_dafny.IntOfInt64(341)) <= 0
						})()
						(globalState).F14 = Companion_Default___.Fm0((_this).F23(), (_dafny.IntOfInt64(80)).Cmp((_this).F23()) <= 0, globalState)
						var _812_v42 _dafny.Map
						_ = _812_v42
						_812_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool), (_this).F23())
						var _813_v43 _dafny.Map
						_ = _813_v43
						_813_v43 = (_812_v42).Merge(_812_v42)
						_813_v43 = _813_v43
						(globalState).F21 = !(true) || (_754_v0)
					} else {
						var _814_v44 D0
						_ = _814_v44
						var _815_v45 _dafny.Map
						_ = _815_v45
						var _816_v46 _dafny.Int
						_ = _816_v46
						var _out34 D0
						_ = _out34
						var _out35 _dafny.Map
						_ = _out35
						var _out36 _dafny.Int
						_ = _out36
						_out34, _out35, _out36 = Companion_Default___.M0(_809_v39, p0, (_this).F23(), globalState)
						_814_v44 = _out34
						_815_v45 = _out35
						_816_v46 = _out36
						var _817_v47 _dafny.Array
						_ = _817_v47
						var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
						_ = _nw131
						_817_v47 = _nw131
						var _818_v48 *C1
						_ = _818_v48
						var _nw132 *C1 = New_C1_()
						_ = _nw132
						_nw132.Ctor__((_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool), _806_v36, (_this).F23())
						_818_v48 = _nw132
						var _819_v49 _dafny.Map
						_ = _819_v49
						_819_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_818_v48, _817_v47)
						var _820_v50 _dafny.Array
						_ = _820_v50
						var _nwElement0_27 _dafny.Array = _817_v47
						_ = _nwElement0_27
						var _nw133 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(23))
						_ = _nw133
						(_nw133).ArraySet1(_nwElement0_27, 0)
						(_nw133).ArraySet1(_817_v47, 1)
						(_nw133).ArraySet1(_817_v47, 2)
						(_nw133).ArraySet1(_817_v47, 3)
						(_nw133).ArraySet1(_817_v47, 4)
						(_nw133).ArraySet1((func() _dafny.Array {
							if (_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool) {
								return _817_v47
							}
							return (func() _dafny.Array {
								if (_819_v49).Contains(_818_v48) {
									return (_819_v49).Get(_818_v48).(_dafny.Array)
								}
								return _817_v47
							})()
						})(), 5)
						(_nw133).ArraySet1(_817_v47, 6)
						(_nw133).ArraySet1(_817_v47, 7)
						(_nw133).ArraySet1(_817_v47, 8)
						(_nw133).ArraySet1(_817_v47, 9)
						(_nw133).ArraySet1(_817_v47, 10)
						(_nw133).ArraySet1(_817_v47, 11)
						(_nw133).ArraySet1(_817_v47, 12)
						(_nw133).ArraySet1(_817_v47, 13)
						(_nw133).ArraySet1(_817_v47, 14)
						(_nw133).ArraySet1(_817_v47, 15)
						(_nw133).ArraySet1(_817_v47, 16)
						(_nw133).ArraySet1(_817_v47, 17)
						(_nw133).ArraySet1(_817_v47, 18)
						(_nw133).ArraySet1(_817_v47, 19)
						(_nw133).ArraySet1(_817_v47, 20)
						(_nw133).ArraySet1(_817_v47, 21)
						(_nw133).ArraySet1(_817_v47, 22)
						_820_v50 = _nw133
						var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_820_v50), 0))
						_ = _index168
						(_820_v50).ArraySet1(_817_v47, (_index168).Int())
						var _821_v51 _dafny.CodePoint
						_ = _821_v51
						_821_v51 = _dafny.CodePoint('r')
						var _822_v52 _dafny.MultiSet
						_ = _822_v52
						_822_v52 = _dafny.MultiSetOf(_821_v51, _dafny.CodePoint('o'), _821_v51)
						var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen(((_this).F27()), 0))
						_ = _index169
						((_this).F27()).ArraySet1((func() _dafny.Int {
							if (_822_v52).Contains(_821_v51) {
								return (_822_v52).Multiplicity(_821_v51)
							}
							return _dafny.IntOfUint32((_dafny.SeqOf((_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool))).Cardinality())
						})(), (_index169).Int())
						var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_820_v50), 0))
						_ = _index170
						var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen(((_this).F27()), 0))
						_ = _index171
						var _rhs178 _dafny.Array = _817_v47
						_ = _rhs178
						var _rhs179 _dafny.Int = ((_dafny.Zero).Minus(_816_v46)).Plus((func() _dafny.Int {
							if (_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool) {
								return (_this).F23()
							}
							return (_this).F23()
						})())
						_ = _rhs179
						var _lhs164 _dafny.Array = _820_v50
						_ = _lhs164
						var _lhs165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_820_v50), 0))
						_ = _lhs165
						var _lhs166 _dafny.Array = (_this).F27()
						_ = _lhs166
						var _lhs167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen(((_this).F27()), 0))
						_ = _lhs167
						(_lhs164).ArraySet1(_rhs178, (_lhs165).Int())
						(_lhs166).ArraySet1(_rhs179, (_lhs167).Int())
						var _823_v53 _dafny.MultiSet
						_ = _823_v53
						_823_v53 = _dafny.MultiSetOf((_this).F23())
						var _824_v54 _dafny.Sequence
						_ = _824_v54
						_824_v54 = _dafny.SeqOf(p0, _dafny.UnicodeSeqOfUtf8Bytes("sxaeugrwk"))
						var _825_v55 _dafny.Map
						_ = _825_v55
						_825_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Int {
							if (_823_v53).Contains((_dafny.MultiSetOf(_816_v46, ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int))).Cardinality()) {
								return (_823_v53).Multiplicity((_dafny.MultiSetOf(_816_v46, ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int))).Cardinality())
							}
							return (_this).F23()
						})()).Minus(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32(((_824_v54).Select((Companion_Default___.SafeIndex(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_824_v54).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
						_825_v55 = _825_v55
						var _826_v56 _dafny.Array
						_ = _826_v56
						var _len0_20 _dafny.Int = _dafny.IntOfInt64(25)
						_ = _len0_20
						var _nw134 _dafny.Array
						_ = _nw134
						if _len0_20.Cmp(_dafny.Zero) == 0 {
							_nw134 = _dafny.NewArray(_len0_20)
						} else {
							var _init20 func(_dafny.Int) _dafny.Map = (func(_827_v45 _dafny.Map, _828_v0 bool) func(_dafny.Int) _dafny.Map {
								return func(_829_i5 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_827_v45, _828_v0)
								}
							})(_815_v45, _754_v0)
							_ = _init20
							var _element0_20 = _init20(_dafny.Zero)
							_ = _element0_20
							_nw134 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
							(_nw134).ArraySet1(_element0_20, 0)
							var _nativeLen0_20 = (_len0_20).Int()
							_ = _nativeLen0_20
							for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
								(_nw134).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
							}
						}
						_826_v56 = _nw134
						var _830_v57 _dafny.Map
						_ = _830_v57
						_830_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_815_v45, false)
						var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_826_v56), 0))
						_ = _index172
						(_826_v56).ArraySet1(_830_v57, (_index172).Int())
						var _831_v58 _dafny.Sequence
						_ = _831_v58
						_831_v58 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(975))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg51 _dafny.Int) interface{} {
								return coer51(arg51)
							}
						}(func(_832_i6 _dafny.Int) _dafny.Sequence {
							return _dafny.UnicodeSeqOfUtf8Bytes("o")
						})))
						var _833_v59 _dafny.Sequence
						_ = _833_v59
						_833_v59 = _dafny.SeqOf(Companion_Default___.Fm21(_dafny.Companion_Sequence_.Update((_831_v58).Select((Companion_Default___.SafeIndex(_816_v46, _dafny.IntOfUint32((_831_v58).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_816_v46, _dafny.IntOfUint32(((_831_v58).Select((Companion_Default___.SafeIndex(_816_v46, _dafny.IntOfUint32((_831_v58).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), p0), globalState))
						var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_826_v56), 0))
						_ = _index173
						var _rhs180 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_815_v45, _754_v0)
						_ = _rhs180
						var _rhs181 bool = false
						_ = _rhs181
						var _rhs182 _dafny.Sequence = _833_v59
						_ = _rhs182
						var _rhs183 _dafny.Int = (_this).F23()
						_ = _rhs183
						var _lhs168 _dafny.Array = _826_v56
						_ = _lhs168
						var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_826_v56), 0))
						_ = _lhs169
						var _lhs170 *GlobalState = globalState
						_ = _lhs170
						var _lhs171 *GlobalState = globalState
						_ = _lhs171
						(_lhs168).ArraySet1(_rhs180, (_lhs169).Int())
						_lhs170.F21 = _rhs181
						_833_v59 = _rhs182
						_lhs171.F8 = _rhs183
						(globalState).F14 = Companion_Default___.Fm0((Companion_D0_.Create_DC1_(_816_v46, !((_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool)), _816_v46)).Dtor_cf3(), !((((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int)).Cmp(_816_v46) <= 0), globalState)
					}
					if true {
						var _834_v60 _dafny.Array
						_ = _834_v60
						var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
						_ = _nw135
						_834_v60 = _nw135
						var _835_v61 D11
						_ = _835_v61
						_835_v61 = Companion_D11_.Create_DC33_(Companion_D11_.Create_DC30_(_834_v60))
						var _836_v62 D11
						_ = _836_v62
						_836_v62 = Companion_D11_.Create_DC30_(_834_v60)
						var _837_v63 _dafny.Array
						_ = _837_v63
						var _nwElement0_28 D11 = _835_v61
						_ = _nwElement0_28
						var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(18))
						_ = _nw136
						(_nw136).ArraySet1(_nwElement0_28, 0)
						(_nw136).ArraySet1(_835_v61, 1)
						(_nw136).ArraySet1(_835_v61, 2)
						(_nw136).ArraySet1(_835_v61, 3)
						(_nw136).ArraySet1(_835_v61, 4)
						(_nw136).ArraySet1(_835_v61, 5)
						(_nw136).ArraySet1(_835_v61, 6)
						(_nw136).ArraySet1(_835_v61, 7)
						(_nw136).ArraySet1(Companion_D11_.Create_DC33_(_836_v62), 8)
						(_nw136).ArraySet1(_835_v61, 9)
						(_nw136).ArraySet1(_835_v61, 10)
						(_nw136).ArraySet1(_835_v61, 11)
						(_nw136).ArraySet1(_835_v61, 12)
						(_nw136).ArraySet1(_835_v61, 13)
						(_nw136).ArraySet1(_835_v61, 14)
						(_nw136).ArraySet1(_835_v61, 15)
						(_nw136).ArraySet1(_835_v61, 16)
						(_nw136).ArraySet1(_835_v61, 17)
						_837_v63 = _nw136
						var _838_v64 _dafny.Map
						_ = _838_v64
						_838_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_806_v36).Cardinality()), (_dafny.Zero).Minus((_806_v36).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.IntOfUint32((_806_v36).Cardinality()))).Uint32()).(_dafny.Int)))), _837_v63)
						_838_v64 = (_838_v64).Update(_dafny.IntOfInt64(885), _837_v63)
						var _839_v65 *C1
						_ = _839_v65
						var _nw137 *C1 = New_C1_()
						_ = _nw137
						_nw137.Ctor__((_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool), _806_v36, (_this).F23())
						_839_v65 = _nw137
						var _840_v66 _dafny.Sequence
						_ = _840_v66
						_840_v66 = _dafny.SeqOf(_839_v65)
						var _841_v67 _dafny.Map
						_ = _841_v67
						_841_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool), Companion_Default___.Fm1(_dafny.IntOfUint32((_840_v66).Cardinality()), Companion_Default___.Fm1((_771_v16).Cardinality(), (_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool), globalState), globalState))
						var _842_v68 _dafny.Map
						_ = _842_v68
						_842_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_841_v67).Cardinality(), _771_v16)
						var _843_v69 _dafny.Map
						_ = _843_v69
						_843_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_842_v68).Update((_this).F23(), _771_v16)).Merge(_842_v68), (_this).F23())
						(globalState).F14 = (func() _dafny.Int {
							if (_843_v69).Contains((func() _dafny.Map {
								if (_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool) {
									return _842_v68
								}
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _771_v16)
							})()) {
								return (_843_v69).Get((func() _dafny.Map {
									if (_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool) {
										return _842_v68
									}
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _771_v16)
								})()).(_dafny.Int)
							}
							return (_this).F23()
						})()
						var _844_v70 _dafny.Array
						_ = _844_v70
						var _nw138 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(10))
						_ = _nw138
						_844_v70 = _nw138
						var _845_v71 _dafny.Map
						_ = _845_v71
						_845_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F23())
						var _846_v72 _dafny.Map
						_ = _846_v72
						_846_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eys")).Cardinality()))
						var _847_v73 _dafny.MultiSet
						_ = _847_v73
						_847_v73 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-251))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg52 _dafny.Int) interface{} {
								return coer52(arg52)
							}
						}(func(_848_i7 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(-445)
						}))).Cardinality()))
						var _849_v74 *C2
						_ = _849_v74
						var _nw139 *C2 = New_C2_()
						_ = _nw139
						_nw139.Ctor__(_845_v71, _754_v0, (_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool), _806_v36, (_dafny.Zero).Minus((func() _dafny.Int {
							if (_846_v72).Contains((_this).Fm10(_dafny.UnicodeSeqOfUtf8Bytes("bms"), _754_v0, _846_v72, true, globalState)) {
								return (_846_v72).Get((_this).Fm10(_dafny.UnicodeSeqOfUtf8Bytes("bms"), _754_v0, _846_v72, true, globalState)).(_dafny.Int)
							}
							return (_847_v73).Cardinality()
						})()))
						_849_v74 = _nw139
						var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_844_v70), 0))
						_ = _index174
						(_844_v70).ArraySet1(_849_v74, (_index174).Int())
						var _850_v75 _dafny.CodePoint
						_ = _850_v75
						_850_v75 = _dafny.CodePoint('o')
						var _851_v76 _dafny.Map
						_ = _851_v76
						_851_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_849_v74, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_754_v0), _850_v75)).Cardinality())
						var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_844_v70), 0))
						_ = _index175
						var _rhs184 bool = !(_754_v0)
						_ = _rhs184
						var _rhs185 bool = (_849_v74).F32()
						_ = _rhs185
						var _rhs186 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F23(), (func() _dafny.Int {
							if (_851_v76).Contains(_849_v74) {
								return (_851_v76).Get(_849_v74).(_dafny.Int)
							}
							return (_this).F23()
						})())
						_ = _rhs186
						var _rhs187 *C2 = _849_v74
						_ = _rhs187
						var _lhs172 *GlobalState = globalState
						_ = _lhs172
						var _lhs173 *GlobalState = globalState
						_ = _lhs173
						var _lhs174 *GlobalState = globalState
						_ = _lhs174
						var _lhs175 _dafny.Array = _844_v70
						_ = _lhs175
						var _lhs176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_844_v70), 0))
						_ = _lhs176
						_lhs172.F9 = _rhs184
						_lhs173.F21 = _rhs185
						_lhs174.F14 = _rhs186
						(_lhs175).ArraySet1(_rhs187, (_lhs176).Int())
						var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))
						_ = _index176
						(_809_v39).ArraySet1(_754_v0, (_index176).Int())
						var _852_v77 _dafny.Map
						_ = _852_v77
						_852_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _839_v65)
						_852_v77 = _852_v77
					} else {
						var _853_v78 _dafny.Map
						_ = _853_v78
						_853_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
							if (_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool) {
								return (_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool)
							}
							return (_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool)
						})(), (_this).F23())
						var _854_v79 _dafny.Map
						_ = _854_v79
						_854_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_this).F23(), (_this).F23()), (_this).F23())
						var _rhs188 bool = _754_v0
						_ = _rhs188
						var _rhs189 _dafny.Int = (func() _dafny.Int {
							if (_854_v79).Contains((_this).F23()) {
								return (_854_v79).Get((_this).F23()).(_dafny.Int)
							}
							return (_this).F23()
						})()
						_ = _rhs189
						var _rhs190 _dafny.Int = ((_853_v78).Merge((_853_v78).Update(_754_v0, _dafny.IntOfInt64(437)))).Cardinality()
						_ = _rhs190
						var _rhs191 _dafny.Map = (_853_v78).Merge(_853_v78)
						_ = _rhs191
						var _lhs177 *GlobalState = globalState
						_ = _lhs177
						var _lhs178 *GlobalState = globalState
						_ = _lhs178
						_lhs177.F7 = _rhs188
						r0 = _rhs189
						_lhs178.F8 = _rhs190
						_853_v78 = _rhs191
						var _855_v80 _dafny.Map
						_ = _855_v80
						_855_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool), p0)
						var _856_v81 D11
						_ = _856_v81
						_856_v81 = Companion_D11_.Create_DC32_(p0)
						var _857_v82 _dafny.Array
						_ = _857_v82
						var _nwElement0_29 _dafny.Map = _855_v80
						_ = _nwElement0_29
						var _nw140 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(11))
						_ = _nw140
						(_nw140).ArraySet1(_nwElement0_29, 0)
						(_nw140).ArraySet1(_855_v80, 1)
						(_nw140).ArraySet1(_855_v80, 2)
						(_nw140).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(Companion_Default___.Fm1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool), _754_v0)).Update((_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool), _754_v0)).Cardinality(), _754_v0, globalState)), _dafny.UnicodeSeqOfUtf8Bytes("lqs")), 3)
						(_nw140).ArraySet1(_855_v80, 4)
						(_nw140).ArraySet1(_855_v80, 5)
						(_nw140).ArraySet1(Companion_Default___.Fm27(globalState), 6)
						(_nw140).ArraySet1(_855_v80, 7)
						(_nw140).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v0, (_856_v81).Dtor_cf53()), 8)
						(_nw140).ArraySet1((_855_v80).Merge(_855_v80), 9)
						(_nw140).ArraySet1(_855_v80, 10)
						_857_v82 = _nw140
						var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_857_v82), 0))
						_ = _index177
						(_857_v82).ArraySet1(_855_v80, (_index177).Int())
						var _858_v83 _dafny.Map
						_ = _858_v83
						_858_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_808_v38, (_855_v80).Merge((_855_v80).Update(_754_v0, p0)))
						var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_857_v82), 0))
						_ = _index178
						(_857_v82).ArraySet1((func() _dafny.Map {
							if (_858_v83).Contains(_dafny.Companion_Sequence_.Concatenate(_808_v38, _808_v38)) {
								return (_858_v83).Get(_dafny.Companion_Sequence_.Concatenate(_808_v38, _808_v38)).(_dafny.Map)
							}
							return Companion_Default___.Fm27(globalState)
						})(), (_index178).Int())
						var _859_v84 D3
						_ = _859_v84
						_859_v84 = Companion_D3_.Create_DC12_(_808_v38)
						var _860_v85 _dafny.Array
						_ = _860_v85
						var _len0_21 _dafny.Int = _dafny.IntOfInt64(6)
						_ = _len0_21
						var _nw141 _dafny.Array
						_ = _nw141
						if _len0_21.Cmp(_dafny.Zero) == 0 {
							_nw141 = _dafny.NewArray(_len0_21)
						} else {
							var _init21 func(_dafny.Int) _dafny.Sequence = (func(_861_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_862_i8 _dafny.Int) _dafny.Sequence {
									return _dafny.SeqOf(_861_p0, _861_p0, _861_p0)
								}
							})(p0)
							_ = _init21
							var _element0_21 = _init21(_dafny.Zero)
							_ = _element0_21
							_nw141 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
							(_nw141).ArraySet1(_element0_21, 0)
							var _nativeLen0_21 = (_len0_21).Int()
							_ = _nativeLen0_21
							for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
								(_nw141).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
							}
						}
						_860_v85 = _nw141
						var _863_v86 _dafny.Sequence
						_ = _863_v86
						_863_v86 = _dafny.SeqOf(Companion_Default___.Fm2(globalState))
						var _864_v87 _dafny.Sequence
						_ = _864_v87
						_864_v87 = _863_v86
						var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_860_v85), 0))
						_ = _index179
						(_860_v85).ArraySet1(_864_v87, (_index179).Int())
						var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_860_v85), 0))
						_ = _index180
						var _rhs192 D3 = _859_v84
						_ = _rhs192
						var _rhs193 _dafny.Sequence = _864_v87
						_ = _rhs193
						var _rhs194 _dafny.Int = (_this).F23()
						_ = _rhs194
						var _lhs179 _dafny.Array = _860_v85
						_ = _lhs179
						var _lhs180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_860_v85), 0))
						_ = _lhs180
						var _lhs181 *GlobalState = globalState
						_ = _lhs181
						_859_v84 = _rhs192
						(_lhs179).ArraySet1(_rhs193, (_lhs180).Int())
						_lhs181.F8 = _rhs194
						(globalState).F14 = (_this).F23()
						var _865_v88 _dafny.Map
						_ = _865_v88
						_865_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_809_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_809_v39), 0))).Int()).(bool))
						_854_v79 = (_854_v79).Update((_865_v88).Cardinality(), _dafny.IntOfInt64(933))
					}
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _866_v89 D13
		_ = _866_v89
		_866_v89 = Companion_D13_.Create_DC38_(_771_v16)
		var _867_v90 D13
		_ = _867_v90
		_867_v90 = Companion_D13_.Create_DC40_(_866_v89)
		var _pat_let_tv35 = _866_v89
		_ = _pat_let_tv35
		var _pat_let_tv36 = _866_v89
		_ = _pat_let_tv36
		var _source7 D13 = func(_pat_let12_0 D13) D13 {
			return func(_870_dt__update__tmp_h1 D13) D13 {
				return func(_pat_let15_0 D13) D13 {
					return func(_871_dt__update_hcf64_h1 D13) D13 {
						return Companion_D13_.Create_DC40_(_871_dt__update_hcf64_h1)
					}(_pat_let15_0)
				}(_pat_let_tv36)
			}(_pat_let12_0)
		}(func(_pat_let13_0 D13) D13 {
			return func(_868_dt__update__tmp_h0 D13) D13 {
				return func(_pat_let14_0 D13) D13 {
					return func(_869_dt__update_hcf64_h0 D13) D13 {
						return Companion_D13_.Create_DC40_(_869_dt__update_hcf64_h0)
					}(_pat_let14_0)
				}(Companion_D13_.Create_DC40_(_pat_let_tv35))
			}(_pat_let13_0)
		}(_867_v90))
		_ = _source7
		if _source7.Is_DC39() {
			var _872___mcc_h5 _dafny.Int = _source7.Get_().(D13_DC39).Cf63
			_ = _872___mcc_h5
			var _873_cf63 _dafny.Int = _872___mcc_h5
			_ = _873_cf63
			var _874_v91 _dafny.Array
			_ = _874_v91
			var _nw142 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw142
			_874_v91 = _nw142
			var _875_v92 _dafny.Set
			_ = _875_v92
			_875_v92 = _dafny.SetOf(_874_v91)
			var _876_v93 _dafny.Sequence
			_ = _876_v93
			_876_v93 = _dafny.SeqOf(_874_v91)
			var _877_v94 _dafny.Sequence
			_ = _877_v94
			_877_v94 = _dafny.SeqOf(_875_v92, _dafny.SetOf(_874_v91))
			var _878_v95 _dafny.Array
			_ = _878_v95
			var _nwElement0_30 _dafny.Set = (func() _dafny.Set {
				if _754_v0 {
					return _875_v92
				}
				return _875_v92
			})()
			_ = _nwElement0_30
			var _nw143 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(19))
			_ = _nw143
			(_nw143).ArraySet1(_nwElement0_30, 0)
			(_nw143).ArraySet1(_875_v92, 1)
			(_nw143).ArraySet1(_875_v92, 2)
			(_nw143).ArraySet1(_dafny.SetOf(_874_v91, (_876_v93).Select((Companion_Default___.SafeIndex(_873_cf63, _dafny.IntOfUint32((_876_v93).Cardinality()))).Uint32()).(_dafny.Array)), 3)
			(_nw143).ArraySet1(_875_v92, 4)
			(_nw143).ArraySet1(_875_v92, 5)
			(_nw143).ArraySet1((_875_v92).Union(_875_v92), 6)
			(_nw143).ArraySet1(_875_v92, 7)
			(_nw143).ArraySet1(_875_v92, 8)
			(_nw143).ArraySet1(_875_v92, 9)
			(_nw143).ArraySet1(_875_v92, 10)
			(_nw143).ArraySet1((_877_v94).Select((Companion_Default___.SafeIndex(_873_cf63, _dafny.IntOfUint32((_877_v94).Cardinality()))).Uint32()).(_dafny.Set), 11)
			(_nw143).ArraySet1(_875_v92, 12)
			(_nw143).ArraySet1(_875_v92, 13)
			(_nw143).ArraySet1(_875_v92, 14)
			(_nw143).ArraySet1(_875_v92, 15)
			(_nw143).ArraySet1(_875_v92, 16)
			(_nw143).ArraySet1(_875_v92, 17)
			(_nw143).ArraySet1(_875_v92, 18)
			_878_v95 = _nw143
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_878_v95), 0))
			_ = _index181
			(_878_v95).ArraySet1(_875_v92, (_index181).Int())
			var _879_v96 _dafny.Array
			_ = _879_v96
			var _nw144 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(23))
			_ = _nw144
			_879_v96 = _nw144
			var _880_v97 _dafny.Sequence
			_ = _880_v97
			_880_v97 = _dafny.SeqOf(Companion_D9_.Create_DC24_(_879_v96))
			var _881_v98 _dafny.Sequence
			_ = _881_v98
			_881_v98 = _dafny.SeqOf(_dafny.IntOfUint32((_880_v97).Cardinality()))
			var _882_v99 D9
			_ = _882_v99
			_882_v99 = Companion_D9_.Create_DC27_(_754_v0, _881_v98, _874_v91, false, true)
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_878_v95), 0))
			_ = _index182
			(_878_v95).ArraySet1(_dafny.SetOf((_882_v99).Dtor_cf46(), _874_v91), (_index182).Int())
			var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index183
			((_this).F27()).ArraySet1((_this).F23(), (_index183).Int())
			var _883_v100 *C1
			_ = _883_v100
			var _nw145 *C1 = New_C1_()
			_ = _nw145
			_nw145.Ctor__(!(_754_v0), _881_v98, (_this).F23())
			_883_v100 = _nw145
			var _884_v101 _dafny.Map
			_ = _884_v101
			_884_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v0, _883_v100)
			var _885_v102 _dafny.Sequence
			_ = _885_v102
			_885_v102 = _dafny.SeqOf(_754_v0)
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))
			_ = _index184
			((_this).F27()).ArraySet1(Companion_Default___.Fm0((_884_v101).Cardinality(), (_885_v102).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_885_v102).Cardinality()))).Uint32()).(bool), globalState), (_index184).Int())
			var _886_v103 _dafny.Map
			_ = _886_v103
			_886_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int), _754_v0, globalState), _754_v0)
			var _887_v104 _dafny.MultiSet
			_ = _887_v104
			_887_v104 = _dafny.MultiSetOf((_this).F23(), (_886_v103).Cardinality(), (_this).F23())
			if (((_887_v104).Cardinality()).Plus(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int))).Cmp(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int)) < 0 {
				_885_v102 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_885_v102, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(878))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}((func(_888_v102 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_889_i9 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_888_v102).Cardinality())
					}
				})(_885_v102)))).Cardinality()), _dafny.IntOfUint32((_885_v102).Cardinality()))).Uint32(), _754_v0), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_754_v0), _885_v102))
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index185
				((_this).F27()).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("chtuw"), p0)), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int), ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("chtuw"), p0))).Cardinality()))).Uint32(), (p0).Select((Companion_Default___.SafeIndex(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality()), (_index185).Int())
				var _890_v105 D3
				_ = _890_v105
				_890_v105 = Companion_D3_.Create_DC11_((func() _dafny.Int {
					if (_887_v104).Contains(_873_cf63) {
						return (_887_v104).Multiplicity(_873_cf63)
					}
					return ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int)
				})())
				_890_v105 = _890_v105
				var _891_v106 _dafny.Map
				_ = _891_v106
				_891_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int), (_this).F23())
				_891_v106 = (func() _dafny.Map {
					var _coll30 = _dafny.NewMapBuilder()
					_ = _coll30
					for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(311), _dafny.IntOfInt64(520))); ; {
						_compr_30, _ok34 := _iter34()
						if !_ok34 {
							break
						}
						var _892_v107 _dafny.Int
						_892_v107 = interface{}(_compr_30).(_dafny.Int)
						if ((_dafny.IntOfInt64(311)).Cmp(_892_v107) <= 0) && ((_892_v107).Cmp(_dafny.IntOfInt64(520)) < 0) {
							_coll30.Add((_892_v107).Plus((_this).F23()), _873_cf63)
						}
					}
					return _coll30.ToMap()
				}()).Merge(_891_v106)
				var _893_v108 *C0
				_ = _893_v108
				var _nw146 *C0 = New_C0_()
				_ = _nw146
				_nw146.Ctor__()
				_893_v108 = _nw146
			} else {
				var _894_v109 _dafny.CodePoint
				_ = _894_v109
				_894_v109 = _dafny.CodePoint('n')
				var _895_v110 _dafny.MultiSet
				_ = _895_v110
				_895_v110 = _dafny.MultiSetOf(_894_v109)
				var _896_v112 D9
				_ = _896_v112
				_896_v112 = Companion_D9_.Create_DC25_(_754_v0, func() _dafny.Set {
					var _coll31 = _dafny.NewBuilder()
					_ = _coll31
					for _iter35 := _dafny.Iterate((_895_v110).Elements()); ; {
						_compr_31, _ok35 := _iter35()
						if !_ok35 {
							break
						}
						var _897_v111 _dafny.CodePoint
						_897_v111 = interface{}(_compr_31).(_dafny.CodePoint)
						if (_895_v110).Contains(_897_v111) {
							_coll31.Add(_897_v111)
						}
					}
					return _coll31.ToSet()
				}())
				var _898_v113 D9
				_ = _898_v113
				_898_v113 = Companion_D9_.Create_DC28_(_896_v112)
				_898_v113 = _898_v113
				var _899_v114 _dafny.Map
				_ = _899_v114
				_899_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v0, ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int))
				var _900_v115 D9
				_ = _900_v115
				_900_v115 = Companion_D9_.Create_DC26_((_this).F23())
				var _rhs195 _dafny.Sequence = p0
				_ = _rhs195
				var _rhs196 bool = ((((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int)).Plus((_899_v114).Cardinality())).Cmp(Companion_Default___.SafeDivisionInt((_this).F23(), _873_cf63)) >= 0
				_ = _rhs196
				var _rhs197 bool = _754_v0
				_ = _rhs197
				var _rhs198 bool = _754_v0
				_ = _rhs198
				var _rhs199 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-409), (_900_v115).Dtor_cf43())
				_ = _rhs199
				var _lhs182 *GlobalState = globalState
				_ = _lhs182
				var _lhs183 *GlobalState = globalState
				_ = _lhs183
				var _lhs184 *GlobalState = globalState
				_ = _lhs184
				var _lhs185 *GlobalState = globalState
				_ = _lhs185
				var _lhs186 *GlobalState = globalState
				_ = _lhs186
				_lhs182.F6 = _rhs195
				_lhs183.F9 = _rhs196
				_lhs184.F9 = _rhs197
				_lhs185.F9 = _rhs198
				_lhs186.F8 = _rhs199
				_754_v0 = !((_this).Fm10(_dafny.UnicodeSeqOfUtf8Bytes("vrkt"), _754_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v0, ((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int)), _754_v0, globalState)) || (_754_v0)
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index186
				((_this).F27()).ArraySet1(((_this).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_this).F27()), 0))).Int()).(_dafny.Int), (_index186).Int())
				var _901_v116 _dafny.MultiSet
				_ = _901_v116
				_901_v116 = _dafny.MultiSetOf(_754_v0)
				var _902_v117 _dafny.Sequence
				_ = _902_v117
				_902_v117 = _dafny.SeqOf(_901_v116, _901_v116)
				var _903_v118 *C3
				_ = _903_v118
				var _nw147 *C3 = New_C3_()
				_ = _nw147
				_nw147.Ctor__(_dafny.Companion_Sequence_.Concatenate(_902_v117, _902_v117), _873_cf63, Companion_Default___.Fm1(_873_cf63, _754_v0, globalState), _881_v98)
				_903_v118 = _nw147
			}
			var _904_v119 D9
			_ = _904_v119
			_904_v119 = Companion_D9_.Create_DC26_((_this).F23())
			var _905_v120 D9
			_ = _905_v120
			_905_v120 = Companion_D9_.Create_DC28_(_904_v119)
			var _906_v121 _dafny.Sequence
			_ = _906_v121
			_906_v121 = _dafny.SeqOf((func() D9 {
				if _754_v0 {
					return _905_v120
				}
				return _905_v120
			})())
			var _907_v122 D14
			_ = _907_v122
			_907_v122 = Companion_D14_.Create_DC41_(_906_v121)
			_906_v121 = _dafny.Companion_Sequence_.Concatenate(_906_v121, (_907_v122).Dtor_cf65())
		} else if _source7.Is_DC38() {
			var _908___mcc_h6 _dafny.Set = _source7.Get_().(D13_DC38).Cf62
			_ = _908___mcc_h6
			var _909_cf62 _dafny.Set = _908___mcc_h6
			_ = _909_cf62
			var _910_v123 _dafny.Array
			_ = _910_v123
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_22
			var _nw148 _dafny.Array
			_ = _nw148
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw148 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) bool = (func(_911_v0 bool) func(_dafny.Int) bool {
					return func(_912_i10 _dafny.Int) bool {
						return _911_v0
					}
				})(_754_v0)
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw148 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw148).ArraySet1(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw148).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_910_v123 = _nw148
			var _913_v124 D0
			_ = _913_v124
			var _914_v125 _dafny.Map
			_ = _914_v125
			var _915_v126 _dafny.Int
			_ = _915_v126
			var _out37 D0
			_ = _out37
			var _out38 _dafny.Map
			_ = _out38
			var _out39 _dafny.Int
			_ = _out39
			_out37, _out38, _out39 = Companion_Default___.M0(_910_v123, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(779))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg54 _dafny.Int) interface{} {
					return coer54(arg54)
				}
			}(func(_916_i11 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lwigmupbn")).Cardinality()), globalState)
			_913_v124 = _out37
			_914_v125 = _out38
			_915_v126 = _out39
			if _754_v0 {
				var _917_v127 _dafny.CodePoint
				_ = _917_v127
				_917_v127 = _dafny.CodePoint('v')
				(globalState).F11 = _917_v127
				var _918_v128 _dafny.Sequence
				_ = _918_v128
				_918_v128 = _dafny.SeqOf(p0, _dafny.UnicodeSeqOfUtf8Bytes("fq"))
				var _919_v129 _dafny.Sequence
				_ = _919_v129
				_919_v129 = _918_v128
				_918_v128 = _dafny.Companion_Sequence_.Concatenate(_918_v128, _dafny.Companion_Sequence_.Concatenate((_919_v129), _918_v128))
				var _920_v130 _dafny.Map
				_ = _920_v130
				_920_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
					if _754_v0 {
						return _910_v123
					}
					return _910_v123
				})(), (func() bool {
					if _754_v0 {
						return _754_v0
					}
					return _754_v0
				})())
				_920_v130 = (_920_v130).Update(_910_v123, !_dafny.Companion_Sequence_.Equal(p0, p0))
				var _921_v131 T1
				_ = _921_v131
				var _nw149 *C1 = New_C1_()
				_ = _nw149
				_nw149.Ctor__(_754_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(771))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}(func(_922_i12 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(-205)
				})), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vfnpyft")).Cardinality()))
				_921_v131 = _nw149
				_921_v131 = _921_v131
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index187
				((_this).F27()).ArraySet1((_this).F23(), (_index187).Int())
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen(((_this).F27()), 0))
				_ = _index188
				((_this).F27()).ArraySet1((_this).F23(), (_index188).Int())
			} else {
				var _923_v132 _dafny.Map
				_ = _923_v132
				_923_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_915_v126, Companion_Default___.Fm6(_dafny.IntOfInt64(962), globalState))
				var _924_v133 _dafny.CodePoint
				_ = _924_v133
				_924_v133 = _dafny.CodePoint('j')
				var _925_v134 _dafny.Map
				_ = _925_v134
				_925_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_924_v133, !(_754_v0))
				var _926_v135 _dafny.Sequence
				_ = _926_v135
				_926_v135 = _dafny.SeqOf((_this).F23())
				var _927_v136 _dafny.Sequence
				_ = _927_v136
				_927_v136 = _dafny.SeqOf((_925_v134).Cardinality(), (_this).F23(), (_926_v135).Select((Companion_Default___.SafeIndex(_915_v126, _dafny.IntOfUint32((_926_v135).Cardinality()))).Uint32()).(_dafny.Int))
				var _rhs200 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (_923_v132).Contains((_this).F23()) {
						return (_923_v132).Get((_this).F23()).(_dafny.CodePoint)
					}
					return _924_v133
				})()
				_ = _rhs200
				var _rhs201 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_926_v135, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(336))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}(func(_928_i13 _dafny.Int) _dafny.Int {
					return (_this).F23()
				})))
				_ = _rhs201
				var _lhs187 *GlobalState = globalState
				_ = _lhs187
				_lhs187.F11 = _rhs200
				r1 = _rhs201
				var _929_v137 _dafny.Sequence
				_ = _929_v137
				_929_v137 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _dafny.IntOfInt64(396)))
				var _930_v138 _dafny.Map
				_ = _930_v138
				_930_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_929_v137, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-831))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}((func(_931_v133 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_932_i14 _dafny.Int) _dafny.CodePoint {
						return _931_v133
					}
				})(_924_v133))))
				_930_v138 = (_930_v138).Update(_dafny.Companion_Sequence_.Concatenate(_929_v137, _929_v137), _dafny.Companion_Sequence_.Concatenate(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(271))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}((func(_933_v133 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_934_i15 _dafny.Int) _dafny.CodePoint {
						return _933_v133
					}
				})(_924_v133)))))
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_910_v123), 0))
				_ = _index189
				(_910_v123).ArraySet1(!(((_dafny.Zero).Minus(_915_v126)).Cmp(_dafny.IntOfInt64(495)) <= 0), (_index189).Int())
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_910_v123), 0))
				_ = _index190
				(_910_v123).ArraySet1(!(_754_v0), (_index190).Int())
				var _935_v139 _dafny.MultiSet
				_ = _935_v139
				_935_v139 = _dafny.MultiSetOf(_dafny.CodePoint('w'), _924_v133, _924_v133)
				var _rhs202 _dafny.Sequence = p0
				_ = _rhs202
				var _rhs203 bool = Companion_Default___.Fm1((_926_v135).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_935_v139).Contains(_924_v133) {
						return (_935_v139).Multiplicity(_924_v133)
					}
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ew")).Cardinality())
				})(), _dafny.IntOfUint32((_926_v135).Cardinality()))).Uint32()).(_dafny.Int), Companion_Default___.Fm1((func() _dafny.Set {
					var _coll32 = _dafny.NewBuilder()
					_ = _coll32
					for _iter36 := _dafny.Iterate((_925_v134).Keys().Elements()); ; {
						_compr_32, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _936_v140 _dafny.CodePoint
						_936_v140 = interface{}(_compr_32).(_dafny.CodePoint)
						if (_925_v134).Contains(_936_v140) {
							_coll32.Add(_936_v140)
						}
					}
					return _coll32.ToSet()
				}()).Cardinality(), true, globalState), globalState)
				_ = _rhs203
				var _rhs204 _dafny.Int = _915_v126
				_ = _rhs204
				var _lhs188 *GlobalState = globalState
				_ = _lhs188
				var _lhs189 *GlobalState = globalState
				_ = _lhs189
				var _lhs190 *GlobalState = globalState
				_ = _lhs190
				_lhs188.F6 = _rhs202
				_lhs189.F7 = _rhs203
				_lhs190.F8 = _rhs204
				var _937_v141 _dafny.Sequence
				_ = _937_v141
				_937_v141 = _dafny.SeqOf((_915_v126).Cmp(_dafny.IntOfInt64(-100)) == 0)
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_910_v123), 0))
				_ = _index191
				(_910_v123).ArraySet1((_937_v141).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_937_v141).Cardinality()))).Uint32()).(bool), (_index191).Int())
			}
			if _754_v0 {
				var _938_v142 _dafny.Sequence
				_ = _938_v142
				_938_v142 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ufjwpiq")).Cardinality()))
				(globalState).F11 = Companion_Default___.Fm6((_dafny.Zero).Minus((_938_v142).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_938_v142).Cardinality()))).Uint32()).(_dafny.Int)), globalState)
				var _939_v143 *C1
				_ = _939_v143
				var _nw150 *C1 = New_C1_()
				_ = _nw150
				_nw150.Ctor__(!(Companion_Default___.Fm1((_this).F23(), !(_754_v0), globalState)) || (_754_v0), _938_v142, (_this).F23())
				_939_v143 = _nw150
				(globalState).F8 = _915_v126
				var _940_v144 _dafny.MultiSet
				_ = _940_v144
				_940_v144 = _dafny.MultiSetOf(!(_754_v0), !(_754_v0) || (_754_v0), _754_v0, _754_v0)
				_940_v144 = _940_v144
				(globalState).F7 = !(_754_v0)
			} else {
				var _941_v145 _dafny.Sequence
				_ = _941_v145
				_941_v145 = _dafny.SeqOf(_915_v126)
				(globalState).F8 = ((_941_v145).Select((Companion_Default___.SafeIndex(_915_v126, _dafny.IntOfUint32((_941_v145).Cardinality()))).Uint32()).(_dafny.Int)).Minus((_this).F23())
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_910_v123), 0))
				_ = _index192
				(_910_v123).ArraySet1(_754_v0, (_index192).Int())
				var _942_v146 _dafny.Sequence
				_ = _942_v146
				_942_v146 = _dafny.SeqOf(_754_v0, _754_v0)
				var _943_v147 _dafny.Map
				_ = _943_v147
				_943_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_942_v146, _dafny.IntOfInt64(267))
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_910_v123), 0))
				_ = _index193
				(_910_v123).ArraySet1(!((func() _dafny.Map {
					if true {
						return Companion_Default___.Fm28(_915_v126, globalState)
					}
					return _943_v147
				})()).Equals(_943_v147), (_index193).Int())
				var _944_v148 _dafny.CodePoint
				_ = _944_v148
				_944_v148 = _dafny.CodePoint('e')
				var _945_v149 D4
				_ = _945_v149
				_945_v149 = Companion_D4_.Create_DC15_((_this).F23(), (p0).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.CodePoint), (_this).F23())
				var _946_v150 _dafny.Set
				_ = _946_v150
				_946_v150 = _dafny.SetOf(_944_v148)
				var _947_v151 _dafny.Map
				_ = _947_v151
				_947_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC7_(_946_v150), _944_v148)
				var _948_v152 D2
				_ = _948_v152
				_948_v152 = Companion_D2_.Create_DC7_(_946_v150)
				var _949_v153 D2
				_ = _949_v153
				_949_v153 = Companion_D2_.Create_DC8_(_915_v126, (_914_v125).Cardinality(), _944_v148, (_dafny.Zero).Minus((_this).F23()), (func() bool {
					if (_914_v125).Contains(_915_v126) {
						return (_914_v125).Get(_915_v126).(bool)
					}
					return (_910_v123).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_910_v123), 0))).Int()).(bool)
				})())
				var _950_v154 _dafny.Array
				_ = _950_v154
				var _nwElement0_31 _dafny.CodePoint = _944_v148
				_ = _nwElement0_31
				var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(19))
				_ = _nw151
				(_nw151).ArraySet1CodePoint(_nwElement0_31, 0)
				(_nw151).ArraySet1CodePoint(_944_v148, 1)
				(_nw151).ArraySet1CodePoint(_944_v148, 2)
				(_nw151).ArraySet1CodePoint(Companion_Default___.Fm6(_dafny.IntOfInt64(531), globalState), 3)
				(_nw151).ArraySet1CodePoint(_944_v148, 4)
				(_nw151).ArraySet1CodePoint((_945_v149).Dtor_cf27(), 5)
				(_nw151).ArraySet1CodePoint(_944_v148, 6)
				(_nw151).ArraySet1CodePoint(_dafny.CodePoint('y'), 7)
				(_nw151).ArraySet1CodePoint(_944_v148, 8)
				(_nw151).ArraySet1CodePoint(_944_v148, 9)
				(_nw151).ArraySet1CodePoint(_944_v148, 10)
				(_nw151).ArraySet1CodePoint(_944_v148, 11)
				(_nw151).ArraySet1CodePoint(Companion_Default___.Fm6(_dafny.IntOfUint32((_942_v146).Cardinality()), globalState), 12)
				(_nw151).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_947_v151).Contains(_948_v152) {
						return (_947_v151).Get(_948_v152).(_dafny.CodePoint)
					}
					return _944_v148
				})(), 13)
				(_nw151).ArraySet1CodePoint(_944_v148, 14)
				(_nw151).ArraySet1CodePoint(_dafny.CodePoint('u'), 15)
				(_nw151).ArraySet1CodePoint(_944_v148, 16)
				(_nw151).ArraySet1CodePoint((_949_v153).Dtor_cf17(), 17)
				(_nw151).ArraySet1CodePoint(_dafny.CodePoint('y'), 18)
				_950_v154 = _nw151
				_950_v154 = _950_v154
				var _951_v155 _dafny.Map
				_ = _951_v155
				_951_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_910_v123).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_910_v123), 0))).Int()).(bool)) == (_754_v0), p0)
				_951_v155 = (_951_v155).Update((_910_v123).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_910_v123), 0))).Int()).(bool), p0)
				var _952_v157 *C3
				_ = _952_v157
				var _nw152 *C3 = New_C3_()
				_ = _nw152
				_nw152.Ctor__(Companion_Default___.Fm29(_915_v126, func() _dafny.Set {
					var _coll33 = _dafny.NewBuilder()
					_ = _coll33
					for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(899), _dafny.IntOfInt64(229))); ; {
						_compr_33, _ok37 := _iter37()
						if !_ok37 {
							break
						}
						var _953_v156 _dafny.Int
						_953_v156 = interface{}(_compr_33).(_dafny.Int)
						if ((_dafny.IntOfInt64(899)).Cmp(_953_v156) <= 0) && ((_953_v156).Cmp(_dafny.IntOfInt64(229)) < 0) {
							_coll33.Add((_953_v156).Minus((_this).F23()))
						}
					}
					return _coll33.ToSet()
				}(), globalState), (_this).F23(), (_910_v123).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_910_v123), 0))).Int()).(bool), _dafny.Companion_Sequence_.Concatenate(_941_v145, _941_v145))
				_952_v157 = _nw152
				_952_v157 = _952_v157
			}
			_754_v0 = !(!(false)) || (_754_v0)
		} else {
			var _954___mcc_h7 D13 = _source7.Get_().(D13_DC40).Cf64
			_ = _954___mcc_h7
			var _955_cf64 D13 = _954___mcc_h7
			_ = _955_cf64
			var _956_v158 _dafny.Array
			_ = _956_v158
			var _nw153 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(25))
			_ = _nw153
			_956_v158 = _nw153
			var _957_v159 _dafny.MultiSet
			_ = _957_v159
			_957_v159 = _dafny.MultiSetOf(_956_v158)
			var _958_v160 _dafny.Sequence
			_ = _958_v160
			_958_v160 = _dafny.SeqOf(Companion_D3_.Create_DC9_(_957_v159))
			var _959_v161 _dafny.MultiSet
			_ = _959_v161
			_959_v161 = _dafny.MultiSetOf(_dafny.IntOfInt64(393))
			var _960_v162 _dafny.Map
			_ = _960_v162
			_960_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_this).F23()), Companion_Default___.Fm1((_this).F23(), _754_v0, globalState))
			var _961_v163 _dafny.Sequence
			_ = _961_v163
			_961_v163 = _dafny.SeqOf(!(_754_v0))
			var _962_v164 _dafny.Set
			_ = _962_v164
			_962_v164 = _dafny.SetOf((_dafny.Zero).Minus((_this).F23()))
			var _963_v165 _dafny.CodePoint
			_ = _963_v165
			_963_v165 = _dafny.CodePoint('p')
			var _964_v166 _dafny.Map
			_ = _964_v166
			_964_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_963_v165, (_this).F23())
			var _965_v167 _dafny.MultiSet
			_ = _965_v167
			_965_v167 = _dafny.MultiSetOf(_754_v0, false)
			var _966_v168 _dafny.Map
			_ = _966_v168
			_966_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v0, _754_v0)
			var _967_v169 _dafny.Map
			_ = _967_v169
			_967_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v0, (_966_v168).Cardinality())
			var _968_v170 _dafny.Array
			_ = _968_v170
			var _nwElement0_32 bool = _754_v0
			_ = _nwElement0_32
			var _nw154 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(27))
			_ = _nw154
			(_nw154).ArraySet1(_nwElement0_32, 0)
			(_nw154).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_958_v160, _958_v160), 1)
			(_nw154).ArraySet1((_959_v161).IsSubsetOf((_959_v161).Update((_this).F23(), Companion_Default___.Abs((_this).F23()))), 2)
			(_nw154).ArraySet1(_754_v0, 3)
			(_nw154).ArraySet1((func() bool {
				if (_960_v162).Contains(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_961_v163).Cardinality())))) {
					return (_960_v162).Get(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_961_v163).Cardinality())))).(bool)
				}
				return _754_v0
			})(), 4)
			(_nw154).ArraySet1(_754_v0, 5)
			(_nw154).ArraySet1((_962_v164).IsProperSubsetOf(_962_v164), 6)
			(_nw154).ArraySet1((_dafny.IntOfInt64(-888)).Cmp((_dafny.Zero).Minus((_this).F23())) != 0, 7)
			(_nw154).ArraySet1(true, 8)
			(_nw154).ArraySet1(((_dafny.Zero).Minus((func() _dafny.Int {
				if (_964_v166).Contains(_963_v165) {
					return (_964_v166).Get(_963_v165).(_dafny.Int)
				}
				return (_959_v161).Cardinality()
			})())).Cmp((_this).F23()) < 0, 9)
			(_nw154).ArraySet1(((_this).F23()).Cmp((_this).F23()) > 0, 10)
			(_nw154).ArraySet1(_754_v0, 11)
			(_nw154).ArraySet1(!((_754_v0) == (_754_v0)), 12)
			(_nw154).ArraySet1(((_this).F23()).Cmp((_this).F23()) <= 0, 13)
			(_nw154).ArraySet1(_754_v0, 14)
			(_nw154).ArraySet1(!(_754_v0), 15)
			(_nw154).ArraySet1(_754_v0, 16)
			(_nw154).ArraySet1(_754_v0, 17)
			(_nw154).ArraySet1(_754_v0, 18)
			(_nw154).ArraySet1(!(_754_v0), 19)
			(_nw154).ArraySet1(((_965_v167).Cardinality()).Cmp((_this).F23()) != 0, 20)
			(_nw154).ArraySet1((_this).Fm10(p0, _754_v0, _967_v169, _754_v0, globalState), 21)
			(_nw154).ArraySet1((_this).Fm10(p0, _754_v0, _967_v169, _754_v0, globalState), 22)
			(_nw154).ArraySet1(_754_v0, 23)
			(_nw154).ArraySet1(_754_v0, 24)
			(_nw154).ArraySet1(_754_v0, 25)
			(_nw154).ArraySet1(_754_v0, 26)
			_968_v170 = _nw154
			_968_v170 = _968_v170
			var _969_v171 _dafny.Array
			_ = _969_v171
			var _len0_23 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_23
			var _nw155 _dafny.Array
			_ = _nw155
			if _len0_23.Cmp(_dafny.Zero) == 0 {
				_nw155 = _dafny.NewArray(_len0_23)
			} else {
				var _init23 func(_dafny.Int) _dafny.Int = func(_970_i16 _dafny.Int) _dafny.Int {
					return (_970_i16).Times((_this).F23())
				}
				_ = _init23
				var _element0_23 = _init23(_dafny.Zero)
				_ = _element0_23
				_nw155 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
				(_nw155).ArraySet1(_element0_23, 0)
				var _nativeLen0_23 = (_len0_23).Int()
				_ = _nativeLen0_23
				for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
					(_nw155).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
				}
			}
			_969_v171 = _nw155
			var _971_v172 _dafny.Map
			_ = _971_v172
			_971_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v0, (_this).F23())
			var _rhs205 _dafny.Int = Companion_Default___.Fm0((_dafny.Zero).Minus((_this).F23()), true, globalState)
			_ = _rhs205
			var _rhs206 _dafny.Array = _969_v171
			_ = _rhs206
			var _rhs207 _dafny.Map = _967_v169
			_ = _rhs207
			var _rhs208 _dafny.Int = (_this).F23()
			_ = _rhs208
			r0 = _rhs205
			_969_v171 = _rhs206
			_971_v172 = _rhs207
			r0 = _rhs208
			var _972_v173 _dafny.Array
			_ = _972_v173
			var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
			_ = _nw156
			_972_v173 = _nw156
			_972_v173 = _972_v173
			var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_968_v170), 0))
			_ = _index194
			(_968_v170).ArraySet1((func() bool {
				if !(_754_v0) {
					return true
				}
				return _754_v0
			})(), (_index194).Int())
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_968_v170), 0))
			_ = _index195
			(_968_v170).ArraySet1(_754_v0, (_index195).Int())
		}
		var _973_v174 D8
		_ = _973_v174
		_973_v174 = Companion_D8_.Create_DC22_(_754_v0)
		var _pat_let_tv37 = _754_v0
		_ = _pat_let_tv37
		(globalState).F7 = (func(_pat_let16_0 D8) D8 {
			return func(_974_dt__update__tmp_h3 D8) D8 {
				return func(_pat_let17_0 bool) D8 {
					return func(_975_dt__update_hcf38_h0 bool) D8 {
						return Companion_D8_.Create_DC22_(_975_dt__update_hcf38_h0)
					}(_pat_let17_0)
				}(_pat_let_tv37)
			}(_pat_let16_0)
		}(_973_v174)).Dtor_cf38()
		r0 = (_this).F23()
		r1 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_this).F23()), (_this).F23())).Cardinality()), ((_this).F23()).Times((_this).F23()), (_this).F23())
		return r0, r1
	}
}
func (_this *C4) F27() _dafny.Array {
	{
		return _this._f27
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f23 _dafny.Int
	_f25 _dafny.Int
	_f26 _dafny.Int
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f23 = _dafny.Zero
	_this._f25 = _dafny.Zero
	_this._f26 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F23() _dafny.Int {
	return _this._f23
}
func (_this *C5) Ctor__(f25 _dafny.Int, f26 _dafny.Int, f23 _dafny.Int) {
	{
		(_this)._f25 = f25
		(_this)._f26 = f26
		(_this)._f23 = f23
	}
}
func (_this *C5) Fm8(p0 D0, p1 _dafny.Set, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			var _coll34 = _dafny.NewMapBuilder()
			_ = _coll34
			for _iter38 := _dafny.Iterate((func() _dafny.Map {
				var _coll35 = _dafny.NewMapBuilder()
				_ = _coll35
				for _iter39 := _dafny.Iterate((_dafny.MultiSetOf((_this).F26(), (_this).F26())).Elements()); ; {
					_compr_35, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _976_v1 _dafny.Int
					_976_v1 = interface{}(_compr_35).(_dafny.Int)
					if (_dafny.MultiSetOf((_this).F26(), (_this).F26())).Contains(_976_v1) {
						_coll35.Add((_976_v1).Minus((func() _dafny.Map {
							var _coll36 = _dafny.NewMapBuilder()
							_ = _coll36
							for _iter40 := _dafny.Iterate(((Companion_D2_.Create_DC7_(func() _dafny.Set {
								var _coll37 = _dafny.NewBuilder()
								_ = _coll37
								for _iter41 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(547))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg59 _dafny.Int) interface{} {
										return coer59(arg59)
									}
								}(func(_977_i0 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('u')
								}))).Elements()); ; {
									_compr_37, _ok41 := _iter41()
									if !_ok41 {
										break
									}
									var _978_v3 _dafny.CodePoint
									_978_v3 = interface{}(_compr_37).(_dafny.CodePoint)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(547))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
										return func(arg60 _dafny.Int) interface{} {
											return coer60(arg60)
										}
									}(func(_977_i0 _dafny.Int) _dafny.CodePoint {
										return _dafny.CodePoint('u')
									})), _978_v3) {
										_coll37.Add(_978_v3)
									}
								}
								return _coll37.ToSet()
							}())).Dtor_cf14()).Elements()); ; {
								_compr_36, _ok40 := _iter40()
								if !_ok40 {
									break
								}
								var _979_v2 _dafny.CodePoint
								_979_v2 = interface{}(_compr_36).(_dafny.CodePoint)
								if ((Companion_D2_.Create_DC7_(func() _dafny.Set {
									var _coll38 = _dafny.NewBuilder()
									_ = _coll38
									for _iter42 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(547))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
										return func(arg61 _dafny.Int) interface{} {
											return coer61(arg61)
										}
									}(func(_977_i0 _dafny.Int) _dafny.CodePoint {
										return _dafny.CodePoint('u')
									}))).Elements()); ; {
										_compr_38, _ok42 := _iter42()
										if !_ok42 {
											break
										}
										var _980_v3 _dafny.CodePoint
										_980_v3 = interface{}(_compr_38).(_dafny.CodePoint)
										if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(547))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
											return func(arg62 _dafny.Int) interface{} {
												return coer62(arg62)
											}
										}(func(_977_i0 _dafny.Int) _dafny.CodePoint {
											return _dafny.CodePoint('u')
										})), _980_v3) {
											_coll38.Add(_980_v3)
										}
									}
									return _coll38.ToSet()
								}())).Dtor_cf14()).Contains(_979_v2) {
									_coll36.Add(_979_v2, true)
								}
							}
							return _coll36.ToMap()
						}()).Cardinality()), false)
					}
				}
				return _coll35.ToMap()
			}()).Keys().Elements()); ; {
				_compr_34, _ok38 := _iter38()
				if !_ok38 {
					break
				}
				var _981_v0 _dafny.Int
				_981_v0 = interface{}(_compr_34).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll39 = _dafny.NewMapBuilder()
					_ = _coll39
					for _iter43 := _dafny.Iterate((_dafny.MultiSetOf((_this).F26(), (_this).F26())).Elements()); ; {
						_compr_39, _ok43 := _iter43()
						if !_ok43 {
							break
						}
						var _976_v1 _dafny.Int
						_976_v1 = interface{}(_compr_39).(_dafny.Int)
						if (_dafny.MultiSetOf((_this).F26(), (_this).F26())).Contains(_976_v1) {
							_coll39.Add((_976_v1).Minus((func() _dafny.Map {
								var _coll40 = _dafny.NewMapBuilder()
								_ = _coll40
								for _iter44 := _dafny.Iterate(((Companion_D2_.Create_DC7_(func() _dafny.Set {
									var _coll41 = _dafny.NewBuilder()
									_ = _coll41
									for _iter45 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(547))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
										return func(arg63 _dafny.Int) interface{} {
											return coer63(arg63)
										}
									}(func(_977_i0 _dafny.Int) _dafny.CodePoint {
										return _dafny.CodePoint('u')
									}))).Elements()); ; {
										_compr_41, _ok45 := _iter45()
										if !_ok45 {
											break
										}
										var _982_v3 _dafny.CodePoint
										_982_v3 = interface{}(_compr_41).(_dafny.CodePoint)
										if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(547))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
											return func(arg64 _dafny.Int) interface{} {
												return coer64(arg64)
											}
										}(func(_977_i0 _dafny.Int) _dafny.CodePoint {
											return _dafny.CodePoint('u')
										})), _982_v3) {
											_coll41.Add(_982_v3)
										}
									}
									return _coll41.ToSet()
								}())).Dtor_cf14()).Elements()); ; {
									_compr_40, _ok44 := _iter44()
									if !_ok44 {
										break
									}
									var _979_v2 _dafny.CodePoint
									_979_v2 = interface{}(_compr_40).(_dafny.CodePoint)
									if ((Companion_D2_.Create_DC7_(func() _dafny.Set {
										var _coll42 = _dafny.NewBuilder()
										_ = _coll42
										for _iter46 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(547))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
											return func(arg65 _dafny.Int) interface{} {
												return coer65(arg65)
											}
										}(func(_977_i0 _dafny.Int) _dafny.CodePoint {
											return _dafny.CodePoint('u')
										}))).Elements()); ; {
											_compr_42, _ok46 := _iter46()
											if !_ok46 {
												break
											}
											var _983_v3 _dafny.CodePoint
											_983_v3 = interface{}(_compr_42).(_dafny.CodePoint)
											if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(547))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
												return func(arg66 _dafny.Int) interface{} {
													return coer66(arg66)
												}
											}(func(_977_i0 _dafny.Int) _dafny.CodePoint {
												return _dafny.CodePoint('u')
											})), _983_v3) {
												_coll42.Add(_983_v3)
											}
										}
										return _coll42.ToSet()
									}())).Dtor_cf14()).Contains(_979_v2) {
										_coll40.Add(_979_v2, true)
									}
								}
								return _coll40.ToMap()
							}()).Cardinality()), false)
						}
					}
					return _coll39.ToMap()
				}()).Contains(_981_v0) {
					_coll34.Add(Companion_Default___.SafeModuloInt(_981_v0, (_this).F25()), false)
				}
			}
			return _coll34.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vdal")).Cardinality()), false)), (true) && (true))
	}
}
func (_this *C5) M1(p0 _dafny.Array, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Set) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Set = _dafny.EmptySet
		_ = r1
		var _984_v0 _dafny.Map
		_ = _984_v0
		_984_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(-317))
		var _rhs209 bool = p1
		_ = _rhs209
		var _rhs210 _dafny.Int = (Companion_Default___.Fm9(p1, (_dafny.Zero).Minus((_this).F25()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F26())).Merge(_984_v0), globalState)).Cardinality()
		_ = _rhs210
		var _lhs191 *GlobalState = globalState
		_ = _lhs191
		_lhs191.F21 = _rhs209
		r0 = _rhs210
		var _985_v1 _dafny.Sequence
		_ = _985_v1
		_985_v1 = _dafny.SeqOf((_this).F23())
		var _986_v2 *C1
		_ = _986_v2
		var _nw157 *C1 = New_C1_()
		_ = _nw157
		_nw157.Ctor__((Companion_Default___.Fm1((_this).F26(), p1, globalState)) || (p1), _dafny.Companion_Sequence_.Concatenate(_985_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(12))).Uint32(), func(coer67 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg67 _dafny.Int) interface{} {
				return coer67(arg67)
			}
		}(func(_987_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-818)
		}))), (_dafny.Zero).Minus((_this).F26()))
		_986_v2 = _nw157
		var _988_v3 _dafny.CodePoint
		_ = _988_v3
		_988_v3 = _dafny.CodePoint('m')
		var _989_v4 _dafny.Map
		_ = _989_v4
		_989_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_988_v3, p0)
		var _990_v5 D2
		_ = _990_v5
		_990_v5 = Companion_D2_.Create_DC8_(_dafny.IntOfInt64(340), _dafny.IntOfInt64(-216), _988_v3, (_this).F26(), p1)
		_989_v4 = (_989_v4).Update((_990_v5).Dtor_cf17(), p0)
		var _991_v6 D1
		_ = _991_v6
		_991_v6 = Companion_D1_.Create_DC4_(Companion_Default___.Fm16(p1, (_this).F23(), Companion_Default___.Fm0((_this).F26(), false, globalState), globalState), p1, p1)
		var _992_i1 _dafny.Int
		_ = _992_i1
		_992_i1 = _dafny.Zero
		{
			for (_991_v6).Dtor_cf7() {
				{
					if (_992_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_992_i1 = (_992_i1).Plus(_dafny.One)
					(globalState).F21 = p1
					var _993_v7 _dafny.Array
					_ = _993_v7
					var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
					_ = _nw158
					_993_v7 = _nw158
					var _994_v8 _dafny.Map
					_ = _994_v8
					_994_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F23())
					var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_993_v7), 0))
					_ = _index196
					(_993_v7).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F23()), (_994_v8).Cardinality()), (_index196).Int())
					var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_993_v7), 0))
					_ = _index197
					(_993_v7).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
						if ((_this).F26()).Cmp((_this).F25()) > 0 {
							return (_this).F23()
						}
						return (_this).F25()
					})()), (_index197).Int())
					var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((p0), 0))
					_ = _index198
					(p0).ArraySet1(p1, (_index198).Int())
					var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((p0), 0))
					_ = _index199
					(p0).ArraySet1(((_this).F26()).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm0((_this).F25(), p1, globalState))))) <= 0, (_index199).Int())
					var _995_v9 _dafny.Sequence
					_ = _995_v9
					_995_v9 = _dafny.UnicodeSeqOfUtf8Bytes("jyx")
					var _996_v10 _dafny.Map
					_ = _996_v10
					_996_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(762), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((p0), 0))).Int()).(bool))
					var _997_v11 _dafny.Map
					_ = _997_v11
					_997_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_995_v9, _996_v10)
					var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_993_v7), 0))
					_ = _index200
					var _rhs211 _dafny.Map = ((_997_v11).Update(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("gxiuegtec"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-119), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gxiuegtec")).Cardinality()))).Uint32(), _988_v3), _996_v10)).Update(Companion_Default___.Fm2(globalState), _996_v10)
					_ = _rhs211
					var _rhs212 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_995_v9).Cardinality()))
					_ = _rhs212
					var _lhs192 _dafny.Array = _993_v7
					_ = _lhs192
					var _lhs193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_993_v7), 0))
					_ = _lhs193
					_997_v11 = _rhs211
					(_lhs192).ArraySet1(_rhs212, (_lhs193).Int())
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
		_984_v0 = (_984_v0).Update(true, (_this).F26())
		var _hi11 _dafny.Int = (_this).F25()
		_ = _hi11
		for _998_i2 := _dafny.IntOfUint32((_985_v1).Cardinality()); _998_i2.Cmp(_hi11) < 0; _998_i2 = _998_i2.Plus(_dafny.One) {
			var _999_v12 _dafny.Sequence
			_ = _999_v12
			_999_v12 = _dafny.SeqOf(((_this).F23()).Cmp((_dafny.Zero).Minus(_dafny.IntOfInt64(-840))) > 0)
			var _1000_v13 D8
			_ = _1000_v13
			_1000_v13 = Companion_D8_.Create_DC19_(_985_v1)
			var _1001_v14 _dafny.Sequence
			_ = _1001_v14
			_1001_v14 = _dafny.UnicodeSeqOfUtf8Bytes("nxqooq")
			var _rhs213 bool = !((_dafny.IntOfInt64(4)).Cmp(((_dafny.Zero).Minus((_this).F23())).Times((_this).F25())) != 0)
			_ = _rhs213
			var _rhs214 bool = (Companion_Default___.SafeModuloInt((_this).F26(), (_this).F23())).Cmp(Companion_Default___.SafeDivisionInt((_this).F26(), (_this).F26())) != 0
			_ = _rhs214
			var _rhs215 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm30((_this).F23(), _1000_v13, _1001_v14, (_this).F25(), globalState), _999_v12), _dafny.Companion_Sequence_.Concatenate(_999_v12, _999_v12))
			_ = _rhs215
			var _rhs216 bool = ((_this).F23()).Cmp((_this).F26()) > 0
			_ = _rhs216
			var _lhs194 *GlobalState = globalState
			_ = _lhs194
			var _lhs195 *GlobalState = globalState
			_ = _lhs195
			var _lhs196 *GlobalState = globalState
			_ = _lhs196
			_lhs194.F7 = _rhs213
			_lhs195.F21 = _rhs214
			_999_v12 = _rhs215
			_lhs196.F21 = _rhs216
			var _1002_v15 D3
			_ = _1002_v15
			_1002_v15 = Companion_D3_.Create_DC10_(p1)
			var _1003_v16 _dafny.Map
			_ = _1003_v16
			_1003_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC10_(p1), Companion_Default___.Fm0((_this).F25(), p1, globalState))
			if (_1003_v16).Contains(_1002_v15) {
				var _1004_v17 _dafny.Map
				_ = _1004_v17
				_1004_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Companion_Sequence_.Update(_985_v1, (Companion_Default___.SafeIndex(_998_i2, _dafny.IntOfUint32((_985_v1).Cardinality()))).Uint32(), (_this).F26()))
				var _1005_v18 D12
				_ = _1005_v18
				_1005_v18 = Companion_D12_.Create_DC36_(false, p1, (_this).F25(), (_this).F25(), _1004_v17)
				var _1006_v19 _dafny.Map
				_ = _1006_v19
				_1006_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1001_v14)).Cardinality(), _988_v3)
				var _1007_v20 _dafny.Sequence
				_ = _1007_v20
				_1007_v20 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _988_v3), _1006_v19)
				var _1008_v21 _dafny.Set
				_ = _1008_v21
				_1008_v21 = _dafny.SetOf((_this).F26())
				var _rhs217 _dafny.Int = (_this).F23()
				_ = _rhs217
				var _rhs218 bool = p1
				_ = _rhs218
				var _rhs219 _dafny.Sequence = _1001_v14
				_ = _rhs219
				var _rhs220 _dafny.Int = Companion_Default___.SafeModuloInt(((_1007_v20).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_1007_v20).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), ((_1008_v21).Union(_1008_v21)).Cardinality())
				_ = _rhs220
				var _rhs221 D12 = Companion_Default___.Fm31((_this).F25(), globalState)
				_ = _rhs221
				var _lhs197 *GlobalState = globalState
				_ = _lhs197
				var _lhs198 *GlobalState = globalState
				_ = _lhs198
				var _lhs199 *GlobalState = globalState
				_ = _lhs199
				var _lhs200 *GlobalState = globalState
				_ = _lhs200
				_lhs197.F8 = _rhs217
				_lhs198.F7 = _rhs218
				_lhs199.F6 = _rhs219
				_lhs200.F8 = _rhs220
				_1005_v18 = _rhs221
				var _1009_v22 _dafny.Map
				_ = _1009_v22
				_1009_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(344), p1)
				var _1010_v23 _dafny.Map
				_ = _1010_v23
				_1010_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _1011_v24 D0
				_ = _1011_v24
				_1011_v24 = Companion_D0_.Create_DC1_((_this).F26(), p1, (_1010_v23).Cardinality())
				_1009_v22 = (_1009_v22).Update((_this).F25(), (_1011_v24).Dtor_cf2())
				var _1012_v25 *C0
				_ = _1012_v25
				var _nw159 *C0 = New_C0_()
				_ = _nw159
				_nw159.Ctor__()
				_1012_v25 = _nw159
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((p0), 0))
				_ = _index201
				(p0).ArraySet1(p1, (_index201).Int())
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((p0), 0))
				_ = _index202
				(p0).ArraySet1(p1, (_index202).Int())
				var _1013_v26 _dafny.Array
				_ = _1013_v26
				var _nw160 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
				_ = _nw160
				_1013_v26 = _nw160
				var _1014_v27 _dafny.Array
				_ = _1014_v27
				var _nw161 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
				_ = _nw161
				_1014_v27 = _nw161
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_1013_v26), 0))
				_ = _index203
				(_1013_v26).ArraySet1(_1014_v27, (_index203).Int())
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_1013_v26), 0))
				_ = _index204
				(_1013_v26).ArraySet1(_1014_v27, (_index204).Int())
			} else {
				var _1015_v28 _dafny.MultiSet
				_ = _1015_v28
				_1015_v28 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1001_v14).Cardinality()))
				(globalState).F7 = (p1) == ((_1015_v28).IsSubsetOf(_dafny.MultiSetOf((_this).F26())))
				var _1016_v29 _dafny.Array
				_ = _1016_v29
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_24
				var _nw162 _dafny.Array
				_ = _nw162
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw162 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.Int = (func(_1017_i2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1018_i3 _dafny.Int) _dafny.Int {
							return (_1018_i3).Times(_1017_i2)
						}
					})(_998_i2)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw162 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw162).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw162).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_1016_v29 = _nw162
				var _1019_v30 _dafny.Map
				_ = _1019_v30
				_1019_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1016_v29)
				var _1020_v31 _dafny.Map
				_ = _1020_v31
				_1020_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_this).F26(), _998_i2), _1019_v30)
				_1020_v31 = (_1020_v31).Update((_this).F25(), _1019_v30)
				var _1021_v32 _dafny.Map
				_ = _1021_v32
				_1021_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.Fm0((_this).F25(), Companion_Default___.Fm1((_this).F23(), p1, globalState), globalState)), (_this).F26())
				var _1022_v33 _dafny.Map
				_ = _1022_v33
				_1022_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_985_v1, p1)
				var _1023_v34 *C2
				_ = _1023_v34
				var _nw163 *C2 = New_C2_()
				_ = _nw163
				_nw163.Ctor__(_1021_v32, (_1022_v33).Equals(_1022_v33), p1, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F25(), (_this).F26()), _985_v1), (_this).F25())
				_1023_v34 = _nw163
				var _1024_v35 _dafny.Array
				_ = _1024_v35
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_25
				var _nw164 _dafny.Array
				_ = _nw164
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw164 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) D12 = func(_1025_i4 _dafny.Int) D12 {
						return Companion_D12_.Create_DC37_(Companion_D12_.Create_DC35_())
					}
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw164 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw164).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw164).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_1024_v35 = _nw164
				var _1026_v36 D12
				_ = _1026_v36
				_1026_v36 = Companion_D12_.Create_DC35_()
				var _1027_v37 D12
				_ = _1027_v37
				_1027_v37 = Companion_D12_.Create_DC37_(_1026_v36)
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_1024_v35), 0))
				_ = _index205
				(_1024_v35).ArraySet1(_1027_v37, (_index205).Int())
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.ArrayLen((_1024_v35), 0))
				_ = _index206
				(_1024_v35).ArraySet1(_1027_v37, (_index206).Int())
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1016_v29), 0))
				_ = _index207
				(_1016_v29).ArraySet1((_this).F25(), (_index207).Int())
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1016_v29), 0))
				_ = _index208
				(_1016_v29).ArraySet1(((_this).F23()).Minus((_this).F26()), (_index208).Int())
			}
			r0 = ((_dafny.Zero).Minus((_this).F26())).Times((_this).F25())
			var _1028_v38 _dafny.Array
			_ = _1028_v38
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_26
			var _nw165 _dafny.Array
			_ = _nw165
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw165 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) _dafny.Int = func(_1029_i5 _dafny.Int) _dafny.Int {
					return (_1029_i5).Minus((_this).F25())
				}
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw165 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw165).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw165).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_1028_v38 = _nw165
			_1028_v38 = _1028_v38
		}
		r0 = (_this).F23()
		var _1030_v39 _dafny.Set
		_ = _1030_v39
		_1030_v39 = _dafny.SetOf((_this).F23())
		r1 = ((_1030_v39).Intersection(_1030_v39)).Intersection(_1030_v39)
		return r0, r1
	}
}
func (_this *C5) M2(p0 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var _1031_v0 bool
		_ = _1031_v0
		_1031_v0 = false
		var _1032_v1 _dafny.Set
		_ = _1032_v1
		_1032_v1 = _dafny.SetOf(_1031_v0, _1031_v0)
		var _1033_v2 _dafny.Map
		_ = _1033_v2
		_1033_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25())
		var _1034_v3 D13
		_ = _1034_v3
		_1034_v3 = Companion_D13_.Create_DC39_((func() _dafny.Int {
			if (_1033_v2).Contains((_this).F23()) {
				return (_1033_v2).Get((_this).F23()).(_dafny.Int)
			}
			return (_this).F23()
		})())
		var _1035_v4 _dafny.Set
		_ = _1035_v4
		_1035_v4 = _dafny.SetOf((_1032_v1).Cardinality())
		var _1036_v5 _dafny.Array
		_ = _1036_v5
		var _nwElement0_33 _dafny.Int = (_this).F23()
		_ = _nwElement0_33
		var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(12))
		_ = _nw166
		(_nw166).ArraySet1(_nwElement0_33, 0)
		(_nw166).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 1)
		(_nw166).ArraySet1((_this).F23(), 2)
		(_nw166).ArraySet1((_this).F23(), 3)
		(_nw166).ArraySet1(_dafny.IntOfInt64(700), 4)
		(_nw166).ArraySet1((_this).F25(), 5)
		(_nw166).ArraySet1((_1034_v3).Dtor_cf63(), 6)
		(_nw166).ArraySet1((_dafny.Zero).Minus((_this).F25()), 7)
		(_nw166).ArraySet1((_this).F26(), 8)
		(_nw166).ArraySet1((_this).F25(), 9)
		(_nw166).ArraySet1((_1035_v4).Cardinality(), 10)
		(_nw166).ArraySet1(_dafny.IntOfInt64(947), 11)
		_1036_v5 = _nw166
		var _1037_v6 _dafny.Map
		_ = _1037_v6
		_1037_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1032_v1, _1036_v5)
		var _1038_v7 D13
		_ = _1038_v7
		_1038_v7 = Companion_D13_.Create_DC38_(_1032_v1)
		var _1039_v8 _dafny.Map
		_ = _1039_v8
		_1039_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1((_this).F23(), !(_1031_v0), globalState), (_1038_v7).Dtor_cf62())
		var _1040_v9 D0
		_ = _1040_v9
		_1040_v9 = Companion_D0_.Create_DC0_(p0)
		var _rhs222 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
			if (_1039_v8).Contains(false) {
				return (_1039_v8).Get(false).(_dafny.Set)
			}
			return _1032_v1
		})(), _1036_v5)
		_ = _rhs222
		var _rhs223 _dafny.Set = (_1032_v1).Difference(_1032_v1)
		_ = _rhs223
		var _rhs224 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_1040_v9).Dtor_cf0(), p0)
		_ = _rhs224
		var _lhs201 *GlobalState = globalState
		_ = _lhs201
		_1037_v6 = _rhs222
		_1032_v1 = _rhs223
		_lhs201.F6 = _rhs224
		_1032_v1 = _1032_v1
		var _1041_v10 _dafny.CodePoint
		_ = _1041_v10
		_1041_v10 = _dafny.CodePoint('k')
		(globalState).F11 = _1041_v10
		r0 = Companion_Default___.SafeModuloInt(((_this).F26()).Minus((_this).F25()), (_this).F26())
		var _1042_v11 _dafny.Array
		_ = _1042_v11
		var _nwElement0_34 _dafny.CodePoint = _1041_v10
		_ = _nwElement0_34
		var _nw167 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(3))
		_ = _nw167
		(_nw167).ArraySet1CodePoint(_nwElement0_34, 0)
		(_nw167).ArraySet1CodePoint(_1041_v10, 1)
		(_nw167).ArraySet1CodePoint(_1041_v10, 2)
		_1042_v11 = _nw167
		var _1043_v12 D9
		_ = _1043_v12
		_1043_v12 = Companion_D9_.Create_DC24_(_1042_v11)
		var _1044_v13 _dafny.MultiSet
		_ = _1044_v13
		_1044_v13 = _dafny.MultiSetOf(_1043_v12)
		var _1045_v14 _dafny.Sequence
		_ = _1045_v14
		_1045_v14 = _dafny.SeqOf(Companion_D9_.Create_DC24_(_1042_v11), _1043_v12, _1043_v12)
		(globalState).F8 = (((_1044_v13).Difference(_1044_v13)).Difference(_dafny.MultiSetFromSeq(_1045_v14))).Cardinality()
		var _1046_v15 _dafny.Map
		_ = _1046_v15
		_1046_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _1031_v0)
		_1046_v15 = (_1046_v15).Update(_1041_v10, _1031_v0)
		r0 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F25()), _dafny.IntOfUint32((p0).Cardinality()))
		r1 = _dafny.SeqOf((_this).F25())
		return r0, r1
	}
}
func (_this *C5) M3(globalState *GlobalState) (bool, _dafny.Sequence, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var _1047_v0 _dafny.Map
		_ = _1047_v0
		_1047_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F26())
		var _1048_v1 bool
		_ = _1048_v1
		_1048_v1 = false
		var _1049_v2 D1
		_ = _1049_v2
		_1049_v2 = Companion_D1_.Create_DC5_(_dafny.IntOfInt64(954), _1048_v1, (_this).F23(), (_this).F26())
		var _1050_v3 _dafny.Sequence
		_ = _1050_v3
		_1050_v3 = _dafny.SeqOf((_this).F25())
		var _1051_v4 _dafny.Map
		_ = _1051_v4
		_1051_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1048_v1, _1048_v1)
		var _1052_v5 T1
		_ = _1052_v5
		var _nw168 *C2 = New_C2_()
		_ = _nw168
		_nw168.Ctor__(_1047_v0, _1048_v1, (_1049_v2).Dtor_cf10(), _1050_v3, (_1051_v4).Cardinality())
		_1052_v5 = _nw168
		var _1053_v6 _dafny.Sequence
		_ = _1053_v6
		_1053_v6 = _dafny.SeqOf(_1052_v5)
		var _1054_v7 _dafny.Sequence
		_ = _1054_v7
		_1054_v7 = _dafny.UnicodeSeqOfUtf8Bytes("uckqsq")
		var _1055_v8 _dafny.Sequence
		_ = _1055_v8
		_1055_v8 = _dafny.SeqOf(_1054_v7)
		var _hi12 _dafny.Int = (func() _dafny.Int {
			if false {
				return _dafny.IntOfUint32(((_1055_v8).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_1055_v8).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
			}
			return (_this).F26()
		})()
		_ = _hi12
		for _1056_i0 := _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1053_v6, _dafny.Companion_Sequence_.Update(_1053_v6, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-858), _dafny.IntOfUint32((_1053_v6).Cardinality()))).Uint32(), _1052_v5)), (Companion_Default___.SafeIndex((_1052_v5).F23(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1053_v6, _dafny.Companion_Sequence_.Update(_1053_v6, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-858), _dafny.IntOfUint32((_1053_v6).Cardinality()))).Uint32(), _1052_v5))).Cardinality()))).Uint32(), _1052_v5)).Cardinality()); _1056_i0.Cmp(_hi12) < 0; _1056_i0 = _1056_i0.Plus(_dafny.One) {
			(globalState).F9 = ((_this).F25()).Cmp((_this).F26()) != 0
			var _1057_v9 _dafny.Array
			_ = _1057_v9
			var _nw169 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw169
			_1057_v9 = _nw169
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_27
			var _nw170 _dafny.Array
			_ = _nw170
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw170 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Int = func(_1058_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1058_i1, (_this).F23())
				}
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw170 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw170).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw170).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1057_v9 = _nw170
			r0 = (func() bool {
				if true {
					return (_1052_v5).F28()
				}
				return (_1052_v5).F28()
			})()
			_1057_v9 = _1057_v9
		}
		if _1048_v1 {
			var _1059_v10 _dafny.Map
			_ = _1059_v10
			_1059_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1054_v7, (_1052_v5).F28())
			_1059_v10 = (_1059_v10).Update(_1054_v7, _1048_v1)
			(_1052_v5).F29_set_(_1050_v3)
			(globalState).F7 = false
			var _1060_v11 _dafny.MultiSet
			_ = _1060_v11
			_1060_v11 = _dafny.MultiSetOf((_1052_v5).F28(), (_1052_v5).F28())
			var _1061_v12 _dafny.Sequence
			_ = _1061_v12
			_1061_v12 = _dafny.SeqOf(_dafny.MultiSetOf((_1052_v5).F28()))
			var _1062_v13 *C3
			_ = _1062_v13
			var _nw171 *C3 = New_C3_()
			_ = _nw171
			_nw171.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1060_v11, _1060_v11), _1061_v12), (_this).F25(), (_1052_v5).F28(), _1050_v3)
			_1062_v13 = _nw171
			var _1063_v14 _dafny.CodePoint
			_ = _1063_v14
			_1063_v14 = _dafny.CodePoint('a')
			var _1064_v15 _dafny.Sequence
			_ = _1064_v15
			_1064_v15 = _dafny.SeqOf(!(_1048_v1))
			var _1065_v16 _dafny.Array
			_ = _1065_v16
			var _nwElement0_35 bool = (_1052_v5).F28()
			_ = _nwElement0_35
			var _nw172 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(27))
			_ = _nw172
			(_nw172).ArraySet1(_nwElement0_35, 0)
			(_nw172).ArraySet1((_1052_v5).F28(), 1)
			(_nw172).ArraySet1((_1052_v5).F28(), 2)
			(_nw172).ArraySet1(true, 3)
			(_nw172).ArraySet1((_1052_v5).F28(), 4)
			(_nw172).ArraySet1((_1052_v5).F28(), 5)
			(_nw172).ArraySet1((_1052_v5).F28(), 6)
			(_nw172).ArraySet1(!(_1048_v1), 7)
			(_nw172).ArraySet1(true, 8)
			(_nw172).ArraySet1((_1052_v5).F28(), 9)
			(_nw172).ArraySet1((_1052_v5).F28(), 10)
			(_nw172).ArraySet1(_1048_v1, 11)
			(_nw172).ArraySet1(_1048_v1, 12)
			(_nw172).ArraySet1((_1064_v15).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_1064_v15).Cardinality()))).Uint32()).(bool), 13)
			(_nw172).ArraySet1(_1048_v1, 14)
			(_nw172).ArraySet1(_1048_v1, 15)
			(_nw172).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfInt64(-246), _1048_v1, globalState), 16)
			(_nw172).ArraySet1(_1048_v1, 17)
			(_nw172).ArraySet1((_1052_v5).F28(), 18)
			(_nw172).ArraySet1((_1052_v5).F28(), 19)
			(_nw172).ArraySet1(false, 20)
			(_nw172).ArraySet1((_1052_v5).F28(), 21)
			(_nw172).ArraySet1(_1048_v1, 22)
			(_nw172).ArraySet1((_1064_v15).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.IntOfUint32((_1064_v15).Cardinality()))).Uint32()).(bool), 23)
			(_nw172).ArraySet1(true, 24)
			(_nw172).ArraySet1((_1052_v5).F28(), 25)
			(_nw172).ArraySet1(_1048_v1, 26)
			_1065_v16 = _nw172
			var _1066_v17 D9
			_ = _1066_v17
			_1066_v17 = Companion_D9_.Create_DC27_((_1052_v5).F28(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(19))).Uint32(), func(coer68 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg68 _dafny.Int) interface{} {
					return coer68(arg68)
				}
			}(func(_1067_i2 _dafny.Int) _dafny.Int {
				return (_this).F23()
			})), _1065_v16, (_1052_v5).F28(), (_1064_v15).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_1064_v15).Cardinality()))).Uint32()).(bool))
			var _1068_v18 D9
			_ = _1068_v18
			_1068_v18 = Companion_D9_.Create_DC28_(_1066_v17)
			var _1069_v19 _dafny.Map
			_ = _1069_v19
			_1069_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1063_v14, _1068_v18)
			_1069_v19 = (_1069_v19).Update(_1063_v14, _1068_v18)
		} else {
			var _1070_v20 *C0
			_ = _1070_v20
			var _nw173 *C0 = New_C0_()
			_ = _nw173
			_nw173.Ctor__()
			_1070_v20 = _nw173
			(globalState).F14 = (_dafny.Zero).Minus(((_1052_v5).F23()).Plus((_1052_v5).F23()))
			var _1071_v21 D11
			_ = _1071_v21
			_1071_v21 = Companion_D11_.Create_DC32_(_1054_v7)
			var _source8 D11 = _1071_v21
			_ = _source8
			if _source8.Is_DC31() {
				var _1072___mcc_h0 _dafny.Int = _source8.Get_().(D11_DC31).Cf52
				_ = _1072___mcc_h0
				var _1073_cf52 _dafny.Int = _1072___mcc_h0
				_ = _1073_cf52
				var _1074_v22 _dafny.Set
				_ = _1074_v22
				_1074_v22 = _dafny.SetOf((_1052_v5).F23(), _dafny.IntOfInt64(335), (_1051_v4).Cardinality(), (_this).F23(), (_1050_v3).Select((Companion_Default___.SafeIndex(_1073_cf52, _dafny.IntOfUint32((_1050_v3).Cardinality()))).Uint32()).(_dafny.Int))
				var _1075_v23 *C3
				_ = _1075_v23
				var _nw174 *C3 = New_C3_()
				_ = _nw174
				_nw174.Ctor__(Companion_Default___.Fm29(Companion_Default___.Fm0(_dafny.IntOfInt64(383), (_1052_v5).F28(), globalState), _1074_v22, globalState), (_1052_v5).F23(), (_1052_v5).F28(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F23()), _1052_v5.F29()))
				_1075_v23 = _nw174
				_1051_v4 = (_1051_v4).Update(Companion_Default___.Fm1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mnssog"), _dafny.IntOfInt64(387))).Cardinality(), (_1052_v5).F28(), globalState), _1048_v1)
				var _1076_v24 _dafny.Set
				_ = _1076_v24
				_1076_v24 = _dafny.SetOf((_1052_v5).F28(), _1048_v1)
				_1076_v24 = _1076_v24
				var _1077_v25 _dafny.Map
				_ = _1077_v25
				_1077_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm24((_1052_v5).F28(), _dafny.IntOfUint32((_1054_v7).Cardinality()), globalState), (_1052_v5).F23())
				var _1078_v26 D9
				_ = _1078_v26
				_1078_v26 = Companion_D9_.Create_DC26_((_1052_v5).F23())
				_1077_v25 = (_1077_v25).Update(_1078_v26, (_1052_v5).F23())
			} else if _source8.Is_DC32() {
				var _1079___mcc_h1 _dafny.Sequence = _source8.Get_().(D11_DC32).Cf53
				_ = _1079___mcc_h1
				var _1080_cf53 _dafny.Sequence = _1079___mcc_h1
				_ = _1080_cf53
				var _rhs225 _dafny.Int = (_1052_v5).F23()
				_ = _rhs225
				var _rhs226 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(695))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg69 _dafny.Int) interface{} {
						return coer69(arg69)
					}
				}(func(_1081_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				}))
				_ = _rhs226
				var _lhs202 *GlobalState = globalState
				_ = _lhs202
				var _lhs203 *GlobalState = globalState
				_ = _lhs203
				_lhs202.F14 = _rhs225
				_lhs203.F6 = _rhs226
				var _1082_v27 _dafny.Sequence
				_ = _1082_v27
				_1082_v27 = _dafny.SeqOf((_1052_v5).F28())
				_1048_v1 = !_dafny.Companion_Sequence_.Contains(_1082_v27, (_1052_v5).F28())
				var _1083_v28 _dafny.Set
				_ = _1083_v28
				_1083_v28 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(915))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg70 _dafny.Int) interface{} {
						return coer70(arg70)
					}
				}(func(_1084_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				})))
				(globalState).F7 = !(_1083_v28).Equals(_dafny.SetOf(_1080_cf53, _1054_v7))
				_1050_v3 = _1052_v5.F29()
			} else if _source8.Is_DC30() {
				var _1085___mcc_h2 _dafny.Array = _source8.Get_().(D11_DC30).Cf51
				_ = _1085___mcc_h2
				var _1086_cf51 _dafny.Array = _1085___mcc_h2
				_ = _1086_cf51
				(globalState).F8 = (Companion_Default___.SafeModuloInt((_this).F26(), (_dafny.Zero).Minus((_1052_v5).F23()))).Times(_dafny.IntOfUint32((_1054_v7).Cardinality()))
				var _1087_v29 _dafny.Set
				_ = _1087_v29
				_1087_v29 = _dafny.SetOf((_1052_v5).F23())
				var _1088_v30 _dafny.Map
				_ = _1088_v30
				_1088_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1048_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1070_v20, (_1052_v5).F28())).Cardinality())
				var _1089_v31 _dafny.Map
				_ = _1089_v31
				_1089_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1088_v30)
				(globalState).F21 = (_1089_v31).Contains(((_dafny.Zero).Minus(Companion_Default___.Fm0((_1087_v29).Cardinality(), !((_1052_v5).F28()), globalState))).Cmp((_this).F23()) != 0)
				(globalState).F7 = (_1052_v5).F28()
				var _1090_v32 _dafny.Map
				_ = _1090_v32
				_1090_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality()), false)
				var _1091_v33 *C2
				_ = _1091_v33
				var _nw175 *C2 = New_C2_()
				_ = _nw175
				_nw175.Ctor__(_1047_v0, !((_1052_v5).F28()), _1048_v1, _1052_v5.F29(), ((_1090_v32).Update((_1052_v5).F23(), (_1052_v5).F28())).Cardinality())
				_1091_v33 = _nw175
				var _1092_v34 _dafny.MultiSet
				_ = _1092_v34
				_1092_v34 = _dafny.MultiSetOf(!(true))
				var _1093_v35 D9
				_ = _1093_v35
				_1093_v35 = Companion_D9_.Create_DC26_((_1052_v5).F23())
				var _1094_v36 _dafny.Set
				_ = _1094_v36
				_1094_v36 = _dafny.SetOf(false)
				var _1095_v37 _dafny.Array
				_ = _1095_v37
				var _nwElement0_36 D9 = Companion_D9_.Create_DC26_((_1052_v5).F23())
				_ = _nwElement0_36
				var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(17))
				_ = _nw176
				(_nw176).ArraySet1(_nwElement0_36, 0)
				(_nw176).ArraySet1(Companion_D9_.Create_DC26_((func() _dafny.Int {
					if (_1092_v34).Contains((_1052_v5).F28()) {
						return (_1092_v34).Multiplicity((_1052_v5).F28())
					}
					return _dafny.IntOfInt64(418)
				})()), 1)
				(_nw176).ArraySet1(_1093_v35, 2)
				(_nw176).ArraySet1(_1093_v35, 3)
				(_nw176).ArraySet1(Companion_Default___.Fm24((_1052_v5).F28(), _dafny.IntOfInt64(159), globalState), 4)
				(_nw176).ArraySet1(Companion_D9_.Create_DC26_((_this).F25()), 5)
				(_nw176).ArraySet1(_1093_v35, 6)
				(_nw176).ArraySet1(_1093_v35, 7)
				(_nw176).ArraySet1(_1093_v35, 8)
				(_nw176).ArraySet1(_1093_v35, 9)
				(_nw176).ArraySet1(_1093_v35, 10)
				(_nw176).ArraySet1(_1093_v35, 11)
				(_nw176).ArraySet1(Companion_D9_.Create_DC26_((_1094_v36).Cardinality()), 12)
				(_nw176).ArraySet1(_1093_v35, 13)
				(_nw176).ArraySet1(_1093_v35, 14)
				(_nw176).ArraySet1(_1093_v35, 15)
				(_nw176).ArraySet1(_1093_v35, 16)
				_1095_v37 = _nw176
				var _1096_v38 _dafny.Map
				_ = _1096_v38
				_1096_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1091_v33, _1095_v37)
				_1096_v38 = (_1096_v38).Update(_1091_v33, _1095_v37)
			} else {
				var _1097___mcc_h3 D11 = _source8.Get_().(D11_DC33).Cf54
				_ = _1097___mcc_h3
				var _1098_cf54 D11 = _1097___mcc_h3
				_ = _1098_cf54
				var _1099_v39 D9
				_ = _1099_v39
				_1099_v39 = Companion_D9_.Create_DC25_(_1048_v1, _dafny.SetOf(_dafny.CodePoint('v')))
				var _1100_v40 D2
				_ = _1100_v40
				_1100_v40 = Companion_D2_.Create_DC7_((_1099_v39).Dtor_cf42())
				var _1101_v41 _dafny.Map
				_ = _1101_v41
				_1101_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1100_v40, _1054_v7)
				var _1102_v42 _dafny.Set
				_ = _1102_v42
				_1102_v42 = _dafny.SetOf(_1070_v20)
				var _1103_v43 _dafny.Set
				_ = _1103_v43
				_1103_v43 = _dafny.SetOf(_1048_v1, (_1052_v5).F28(), (_1052_v5).F28())
				var _1104_v44 _dafny.Sequence
				_ = _1104_v44
				_1104_v44 = _dafny.SeqOf(!((_1070_v20).Fm20(_1101_v41, (_1052_v5).F23(), (_this).F26(), (_1102_v42).Cardinality(), globalState)) || ((_1052_v5).F28()), (_1052_v5).F28(), (func() bool {
					if (_1052_v5).F28() {
						return (_1052_v5).F28()
					}
					return (_1052_v5).F28()
				})(), ((_1103_v43).Cardinality()).Cmp((_1052_v5).F23()) >= 0)
				(globalState).F21 = (_1104_v44).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1052_v5).F28() {
						return (_this).F25()
					}
					return _dafny.IntOfInt64(414)
				})(), _dafny.IntOfUint32((_1104_v44).Cardinality()))).Uint32()).(bool)
				var _1105_v45 _dafny.Array
				_ = _1105_v45
				var _nw177 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
				_ = _nw177
				_1105_v45 = _nw177
				var _1106_v46 _dafny.Sequence
				_ = _1106_v46
				_1106_v46 = _dafny.SeqOf(_dafny.SeqOf((_1052_v5).F23(), (_dafny.Zero).Minus((_1052_v5).F23())))
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_1105_v45), 0))
				_ = _index209
				(_1105_v45).ArraySet1(_1106_v46, (_index209).Int())
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_1105_v45), 0))
				_ = _index210
				(_1105_v45).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(506))).Uint32(), func(coer71 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}((func(_1107_v5 T1) func(_dafny.Int) _dafny.Sequence {
					return func(_1108_i5 _dafny.Int) _dafny.Sequence {
						return _1107_v5.F29()
					}
				})(_1052_v5))), (Companion_Default___.SafeIndex((func() _dafny.Int {
					if _1048_v1 {
						return (_1052_v5).F23()
					}
					return _dafny.IntOfUint32((_1055_v8).Cardinality())
				})(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(506))).Uint32(), func(coer72 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg72 _dafny.Int) interface{} {
						return coer72(arg72)
					}
				}((func(_1109_v5 T1) func(_dafny.Int) _dafny.Sequence {
					return func(_1110_i5 _dafny.Int) _dafny.Sequence {
						return _1109_v5.F29()
					}
				})(_1052_v5)))).Cardinality()))).Uint32(), _1052_v5.F29()), (_index210).Int())
				var _1111_v47 _dafny.Array
				_ = _1111_v47
				var _nw178 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(26))
				_ = _nw178
				_1111_v47 = _nw178
				var _1112_v48 D9
				_ = _1112_v48
				_1112_v48 = Companion_D9_.Create_DC24_(_1111_v47)
				var _1113_v49 D9
				_ = _1113_v49
				_1113_v49 = Companion_D9_.Create_DC28_(_1112_v48)
				_1113_v49 = _1113_v49
				var _1114_v50 _dafny.Array
				_ = _1114_v50
				var _nw179 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(11))
				_ = _nw179
				_1114_v50 = _nw179
				var _1115_v51 _dafny.CodePoint
				_ = _1115_v51
				_1115_v51 = _dafny.CodePoint('l')
				var _1116_v52 _dafny.Map
				_ = _1116_v52
				_1116_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1115_v51, (_1052_v5).F28())
				var _1117_v53 _dafny.Set
				_ = _1117_v53
				_1117_v53 = _dafny.SetOf(_1116_v52)
				var _1118_v54 _dafny.Map
				_ = _1118_v54
				_1118_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1116_v52, (_1052_v5).F23())
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1114_v50), 0))
				_ = _index211
				(_1114_v50).ArraySet1((_1117_v53).Union(func() _dafny.Set {
					var _coll43 = _dafny.NewBuilder()
					_ = _coll43
					for _iter47 := _dafny.Iterate((_1118_v54).Keys().Elements()); ; {
						_compr_43, _ok47 := _iter47()
						if !_ok47 {
							break
						}
						var _1119_v55 _dafny.Map
						_1119_v55 = interface{}(_compr_43).(_dafny.Map)
						if (_1118_v54).Contains(_1119_v55) {
							_coll43.Add(_1119_v55)
						}
					}
					return _coll43.ToSet()
				}()), (_index211).Int())
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1114_v50), 0))
				_ = _index212
				var _rhs227 _dafny.Int = (_1052_v5).F23()
				_ = _rhs227
				var _rhs228 _dafny.Set = (_1117_v53).Difference(_1117_v53)
				_ = _rhs228
				var _rhs229 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1054_v7, (Companion_Default___.SafeIndex((_1051_v4).Cardinality(), _dafny.IntOfUint32((_1054_v7).Cardinality()))).Uint32(), _1115_v51)
				_ = _rhs229
				var _lhs204 *GlobalState = globalState
				_ = _lhs204
				var _lhs205 _dafny.Array = _1114_v50
				_ = _lhs205
				var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1114_v50), 0))
				_ = _lhs206
				var _lhs207 *GlobalState = globalState
				_ = _lhs207
				_lhs204.F14 = _rhs227
				(_lhs205).ArraySet1(_rhs228, (_lhs206).Int())
				_lhs207.F6 = _rhs229
			}
			var _1120_v56 _dafny.Sequence
			_ = _1120_v56
			_1120_v56 = _dafny.SeqOf(_1048_v1, (_1052_v5).F28(), (_1052_v5).F28())
			var _1121_v57 *C1
			_ = _1121_v57
			var _nw180 *C1 = New_C1_()
			_ = _nw180
			_nw180.Ctor__(!(!((_1120_v56).Select((Companion_Default___.SafeIndex((_1052_v5).F23(), _dafny.IntOfUint32((_1120_v56).Cardinality()))).Uint32()).(bool))), _1052_v5.F29(), ((_this).F23()).Plus((_this).F23()))
			_1121_v57 = _nw180
			var _1122_v58 _dafny.CodePoint
			_ = _1122_v58
			_1122_v58 = _dafny.CodePoint('y')
			var _1123_v59 _dafny.Array
			_ = _1123_v59
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_28
			var _nw181 _dafny.Array
			_ = _nw181
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw181 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) _dafny.Int = (func(_1124_v5 T1) func(_dafny.Int) _dafny.Int {
					return func(_1125_i6 _dafny.Int) _dafny.Int {
						return (_1125_i6).Times((_1124_v5).F23())
					}
				})(_1052_v5)
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw181 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw181).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw181).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1123_v59 = _nw181
			var _1126_v60 _dafny.Sequence
			_ = _1126_v60
			_1126_v60 = _dafny.SeqOf(_1123_v59, _1123_v59)
			var _1127_v61 D1
			_ = _1127_v61
			_1127_v61 = Companion_D1_.Create_DC3_(_1126_v60)
			var _1128_v62 _dafny.Map
			_ = _1128_v62
			_1128_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1122_v58, _1127_v61)
			var _1129_v63 _dafny.Set
			_ = _1129_v63
			_1129_v63 = _dafny.SetOf((_1052_v5).F28(), !((_1052_v5).F28()), (_1052_v5).F28())
			var _1130_v64 _dafny.Map
			_ = _1130_v64
			_1130_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1129_v63, (_1128_v62).Merge(_1128_v62))
			_1128_v62 = (func() _dafny.Map {
				if (_1130_v64).Contains((_1129_v63).Difference(_1129_v63)) {
					return (_1130_v64).Get((_1129_v63).Difference(_1129_v63)).(_dafny.Map)
				}
				return _1128_v62
			})()
		}
		(globalState).F7 = (_1052_v5).F28()
		var _hi13 _dafny.Int = ((_1052_v5).F23()).Minus((_1052_v5).F23())
		_ = _hi13
		for _1131_i7 := (_this).F23(); _1131_i7.Cmp(_hi13) < 0; _1131_i7 = _1131_i7.Plus(_dafny.One) {
			var _1132_v65 _dafny.Array
			_ = _1132_v65
			var _len0_29 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_29
			var _nw182 _dafny.Array
			_ = _nw182
			if _len0_29.Cmp(_dafny.Zero) == 0 {
				_nw182 = _dafny.NewArray(_len0_29)
			} else {
				var _init29 func(_dafny.Int) _dafny.Int = func(_1133_i8 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1133_i8, (_this).F23())
				}
				_ = _init29
				var _element0_29 = _init29(_dafny.Zero)
				_ = _element0_29
				_nw182 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
				(_nw182).ArraySet1(_element0_29, 0)
				var _nativeLen0_29 = (_len0_29).Int()
				_ = _nativeLen0_29
				for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
					(_nw182).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
				}
			}
			_1132_v65 = _nw182
			var _1134_v66 *C4
			_ = _1134_v66
			var _nw183 *C4 = New_C4_()
			_ = _nw183
			_nw183.Ctor__(_1132_v65, _dafny.IntOfUint32((_1054_v7).Cardinality()))
			_1134_v66 = _nw183
			var _1135_v67 _dafny.MultiSet
			_ = _1135_v67
			_1135_v67 = _dafny.MultiSetOf(_1134_v66)
			var _1136_v68 D11
			_ = _1136_v68
			_1136_v68 = Companion_D11_.Create_DC32_(_1054_v7)
			var _pat_let_tv38 = _1054_v7
			_ = _pat_let_tv38
			var _1137_v69 _dafny.MultiSet
			_ = _1137_v69
			_1137_v69 = _dafny.MultiSetOf((func(_pat_let18_0 D11) D11 {
				return func(_1138_dt__update__tmp_h0 D11) D11 {
					return func(_pat_let19_0 _dafny.Sequence) D11 {
						return func(_1139_dt__update_hcf53_h0 _dafny.Sequence) D11 {
							return Companion_D11_.Create_DC32_(_1139_dt__update_hcf53_h0)
						}(_pat_let19_0)
					}(_pat_let_tv38)
				}(_pat_let18_0)
			}(_1136_v68)).Dtor_cf53(), _1054_v7)
			if Companion_Default___.Fm1((func() _dafny.Int {
				if (_1135_v67).Contains(_1134_v66) {
					return (_1135_v67).Multiplicity(_1134_v66)
				}
				return (_this).F23()
			})(), (_1137_v69).IsSubsetOf(_1137_v69), globalState) {
				var _1140_v70 _dafny.Array
				_ = _1140_v70
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_30
				var _nw184 _dafny.Array
				_ = _nw184
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw184 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) bool = (func(_1141_v5 T1) func(_dafny.Int) bool {
						return func(_1142_i9 _dafny.Int) bool {
							return ((_1141_v5).F28()) == (!((_1141_v5).F28()))
						}
					})(_1052_v5)
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
				_1140_v70 = _nw184
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_1140_v70), 0))
				_ = _index213
				(_1140_v70).ArraySet1((_1052_v5).F28(), (_index213).Int())
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_1140_v70), 0))
				_ = _index214
				(_1140_v70).ArraySet1((_1052_v5).F28(), (_index214).Int())
				var _1143_v71 _dafny.Array
				_ = _1143_v71
				var _nw185 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
				_ = _nw185
				_1143_v71 = _nw185
				var _1144_v72 _dafny.Map
				_ = _1144_v72
				_1144_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v5).F28(), _1143_v71)
				_1144_v72 = (_1144_v72).Update(!(!(true)) || (true), _1143_v71)
				var _1145_v73 _dafny.CodePoint
				_ = _1145_v73
				_1145_v73 = _dafny.CodePoint('b')
				(globalState).F11 = _1145_v73
				var _1146_v74 _dafny.Array
				_ = _1146_v74
				var _nwElement0_37 _dafny.Sequence = _1054_v7
				_ = _nwElement0_37
				var _nw186 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(7))
				_ = _nw186
				(_nw186).ArraySet1(_nwElement0_37, 0)
				(_nw186).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1054_v7, _1054_v7), 1)
				(_nw186).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("vrrj"), (Companion_Default___.SafeIndex(Companion_Default___.Fm0((_this).F26(), (_1052_v5).F28(), globalState), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vrrj")).Cardinality()))).Uint32(), _1145_v73), 2)
				(_nw186).ArraySet1(_1054_v7, 3)
				(_nw186).ArraySet1(Companion_Default___.Fm2(globalState), 4)
				(_nw186).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(379))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}((func(_1147_v73 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1148_i10 _dafny.Int) _dafny.CodePoint {
						return _1147_v73
					}
				})(_1145_v73))), 5)
				(_nw186).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("etxm"), _1054_v7), 6)
				_1146_v74 = _nw186
				var _1149_v75 _dafny.Map
				_ = _1149_v75
				_1149_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v5).F28(), _1131_i7)
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_1146_v74), 0))
				_ = _index215
				(_1146_v74).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1054_v7, _1054_v7), (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1149_v75).Contains(_1048_v1) {
						return (_1149_v75).Get(_1048_v1).(_dafny.Int)
					}
					return (_1149_v75).Cardinality()
				})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1054_v7, _1054_v7)).Cardinality()))).Uint32(), _1145_v73), (_index215).Int())
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_1146_v74), 0))
				_ = _index216
				(_1146_v74).ArraySet1(_1054_v7, (_index216).Int())
				var _1150_v76 _dafny.MultiSet
				_ = _1150_v76
				_1150_v76 = _dafny.MultiSetOf(_1145_v73)
				var _1151_v77 D8
				_ = _1151_v77
				_1151_v77 = Companion_D8_.Create_DC20_((_1052_v5).F28(), _1131_i7)
				var _1152_v78 *C1
				_ = _1152_v78
				var _nw187 *C1 = New_C1_()
				_ = _nw187
				_nw187.Ctor__((_1052_v5).F28(), _1050_v3, (_1052_v5).F23())
				_1152_v78 = _nw187
				var _1153_v79 _dafny.Array
				_ = _1153_v79
				var _nwElement0_38 *C1 = _1152_v78
				_ = _nwElement0_38
				var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(18))
				_ = _nw188
				(_nw188).ArraySet1(_nwElement0_38, 0)
				(_nw188).ArraySet1(_1152_v78, 1)
				(_nw188).ArraySet1(_1152_v78, 2)
				(_nw188).ArraySet1(_1152_v78, 3)
				(_nw188).ArraySet1(_1152_v78, 4)
				(_nw188).ArraySet1(_1152_v78, 5)
				(_nw188).ArraySet1(_1152_v78, 6)
				(_nw188).ArraySet1(_1152_v78, 7)
				(_nw188).ArraySet1(_1152_v78, 8)
				(_nw188).ArraySet1(_1152_v78, 9)
				(_nw188).ArraySet1(_1152_v78, 10)
				(_nw188).ArraySet1(_1152_v78, 11)
				(_nw188).ArraySet1(_1152_v78, 12)
				(_nw188).ArraySet1(_1152_v78, 13)
				(_nw188).ArraySet1(_1152_v78, 14)
				(_nw188).ArraySet1(_1152_v78, 15)
				(_nw188).ArraySet1(_1152_v78, 16)
				(_nw188).ArraySet1(_1152_v78, 17)
				_1153_v79 = _nw188
				var _1154_v80 _dafny.Map
				_ = _1154_v80
				_1154_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v5).F28(), _1153_v79)
				var _1155_v81 _dafny.Map
				_ = _1155_v81
				_1155_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1140_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_1140_v70), 0))).Int()).(bool)), _1052_v5.F29())
				var _1156_v82 _dafny.MultiSet
				_ = _1156_v82
				_1156_v82 = _dafny.MultiSetOf((_1155_v81).Cardinality())
				(globalState).F14 = Companion_Default___.Fm0((((_dafny.MultiSetFromSeq(_1052_v5.F29())).Update((_1150_v76).Cardinality(), Companion_Default___.Abs((_1151_v77).Dtor_cf34()))).Cardinality()).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm2(globalState)).Cardinality()))), !((_dafny.MultiSetOf((_1154_v80).Cardinality(), (_1052_v5).F23())).IsProperSubsetOf(_1156_v82)), globalState)
			} else {
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen(((_1134_v66).F27()), 0))
				_ = _index217
				((_1134_v66).F27()).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hoqqdqjyb"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-545))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg74 _dafny.Int) interface{} {
						return coer74(arg74)
					}
				}(func(_1157_i11 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				})))).Cardinality()), (_index217).Int())
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen(((_1134_v66).F27()), 0))
				_ = _index218
				((_1134_v66).F27()).ArraySet1(((_1052_v5).F23()).Plus(_dafny.IntOfUint32((_1054_v7).Cardinality())), (_index218).Int())
				var _1158_v83 *C0
				_ = _1158_v83
				var _nw189 *C0 = New_C0_()
				_ = _nw189
				_nw189.Ctor__()
				_1158_v83 = _nw189
				var _1159_v84 _dafny.Sequence
				_ = _1159_v84
				_1159_v84 = _dafny.SeqOf(false)
				var _1160_v85 _dafny.CodePoint
				_ = _1160_v85
				_1160_v85 = _dafny.CodePoint('l')
				var _1161_v86 _dafny.Set
				_ = _1161_v86
				_1161_v86 = _dafny.SetOf(_1160_v85)
				var _1162_v87 D9
				_ = _1162_v87
				_1162_v87 = Companion_D9_.Create_DC25_(true, _1161_v86)
				var _1163_v88 _dafny.Map
				_ = _1163_v88
				_1163_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v5).F28(), (_this).F26())
				var _1164_v89 D8
				_ = _1164_v89
				_1164_v89 = Companion_D8_.Create_DC21_((_1134_v66).Fm10(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(224))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}((func(_1165_v85 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1166_i12 _dafny.Int) _dafny.CodePoint {
						return _1165_v85
					}
				})(_1160_v85))), (_1052_v5).F28(), _1163_v88, (_1159_v84).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ixco")).Cardinality()), _dafny.IntOfUint32((_1159_v84).Cardinality()))).Uint32()).(bool), globalState), (_1052_v5).F28(), (_1052_v5).F23())
				var _1167_v90 _dafny.MultiSet
				_ = _1167_v90
				_1167_v90 = _dafny.MultiSetOf(_dafny.CodePoint('q'))
				var _1168_v92 D2
				_ = _1168_v92
				_1168_v92 = Companion_D2_.Create_DC7_(func() _dafny.Set {
					var _coll44 = _dafny.NewBuilder()
					_ = _coll44
					for _iter48 := _dafny.Iterate((_1167_v90).Elements()); ; {
						_compr_44, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _1169_v91 _dafny.CodePoint
						_1169_v91 = interface{}(_compr_44).(_dafny.CodePoint)
						if (_1167_v90).Contains(_1169_v91) {
							_coll44.Add(_1169_v91)
						}
					}
					return _coll44.ToSet()
				}())
				var _1170_v93 _dafny.Map
				_ = _1170_v93
				_1170_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1168_v92, _1054_v7)
				var _1171_v94 _dafny.Array
				_ = _1171_v94
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_31
				var _nw190 _dafny.Array
				_ = _nw190
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw190 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) bool = (func(_1172_v1 bool) func(_dafny.Int) bool {
						return func(_1173_i13 _dafny.Int) bool {
							return _1172_v1
						}
					})(_1048_v1)
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw190 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw190).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw190).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1171_v94 = _nw190
				var _1174_v95 _dafny.Set
				_ = _1174_v95
				_1174_v95 = _dafny.SetOf(_1050_v3, _1052_v5.F29(), _1050_v3)
				var _1175_v96 _dafny.Array
				_ = _1175_v96
				var _nwElement0_39 bool = _1048_v1
				_ = _nwElement0_39
				var _nw191 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(27))
				_ = _nw191
				(_nw191).ArraySet1(_nwElement0_39, 0)
				(_nw191).ArraySet1(_1048_v1, 1)
				(_nw191).ArraySet1((_1159_v84).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1054_v7).Cardinality()), _dafny.IntOfUint32((_1159_v84).Cardinality()))).Uint32()).(bool), 2)
				(_nw191).ArraySet1(_1048_v1, 3)
				(_nw191).ArraySet1(false, 4)
				(_nw191).ArraySet1((_1162_v87).Dtor_cf41(), 5)
				(_nw191).ArraySet1((func() bool {
					if (_1052_v5).F28() {
						return _1048_v1
					}
					return false
				})(), 6)
				(_nw191).ArraySet1((_1052_v5).F28(), 7)
				(_nw191).ArraySet1((_1052_v5).F28(), 8)
				(_nw191).ArraySet1((_1052_v5).F28(), 9)
				(_nw191).ArraySet1((_1052_v5).F28(), 10)
				(_nw191).ArraySet1((_1052_v5).F28(), 11)
				(_nw191).ArraySet1((_1052_v5).F28(), 12)
				(_nw191).ArraySet1(false, 13)
				(_nw191).ArraySet1(!_dafny.Companion_Sequence_.Equal(_1054_v7, _1054_v7), 14)
				(_nw191).ArraySet1(true, 15)
				(_nw191).ArraySet1((_1052_v5).F28(), 16)
				(_nw191).ArraySet1((_1164_v89).Dtor_cf36(), 17)
				(_nw191).ArraySet1(_1048_v1, 18)
				(_nw191).ArraySet1((_1052_v5).F28(), 19)
				(_nw191).ArraySet1(!(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v5).F28(), _1171_v94)).Contains((_1158_v83).Fm20(_1170_v93, _dafny.IntOfUint32((_1054_v7).Cardinality()), (_1052_v5).F23(), (_1052_v5).F23(), globalState))), 20)
				(_nw191).ArraySet1((_1052_v5).F28(), 21)
				(_nw191).ArraySet1((_1052_v5).F28(), 22)
				(_nw191).ArraySet1((_1052_v5).F28(), 23)
				(_nw191).ArraySet1(true, 24)
				(_nw191).ArraySet1((_1052_v5).F28(), 25)
				(_nw191).ArraySet1((_1174_v95).IsSubsetOf(_1174_v95), 26)
				_1175_v96 = _nw191
				var _1176_v97 _dafny.Sequence
				_ = _1176_v97
				_1176_v97 = _dafny.SeqOf(_1175_v96, _1175_v96, _1175_v96)
				_1175_v96 = (_1176_v97).Select((Companion_Default___.SafeIndex((_1052_v5).F23(), _dafny.IntOfUint32((_1176_v97).Cardinality()))).Uint32()).(_dafny.Array)
				(globalState).F8 = Companion_Default___.Fm0(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(545))).Uint32(), func(coer76 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg76 _dafny.Int) interface{} {
						return coer76(arg76)
					}
				}(func(_1177_i14 _dafny.Int) _dafny.Int {
					return (_this).F26()
				}))).Cardinality()), _1048_v1, globalState)
				var _1178_v98 _dafny.MultiSet
				_ = _1178_v98
				_1178_v98 = _dafny.MultiSetOf(_1159_v84)
				(globalState).F8 = ((func() _dafny.Int {
					if (_1178_v98).Contains(_1159_v84) {
						return (_1178_v98).Multiplicity(_1159_v84)
					}
					return (_this).F23()
				})()).Plus(((_1134_v66).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen(((_1134_v66).F27()), 0))).Int()).(_dafny.Int))
			}
			r0 = _1048_v1
			var _1179_v99 _dafny.CodePoint
			_ = _1179_v99
			_1179_v99 = _dafny.CodePoint('g')
			(globalState).F21 = _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(931))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg77 _dafny.Int) interface{} {
					return coer77(arg77)
				}
			}((func(_1180_v99 _dafny.CodePoint, _1181_v7 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
				return func(_1182_i15 _dafny.Int) _dafny.CodePoint {
					return (Companion_D4_.Create_DC15_((_this).F26(), _1180_v99, _dafny.IntOfUint32((_1181_v7).Cardinality()))).Dtor_cf27()
				}
			})(_1179_v99, _1054_v7))), _1179_v99)
			var _source9 D8 = Companion_D8_.Create_DC22_(_1048_v1)
			_ = _source9
			if _source9.Is_DC20() {
				var _1183___mcc_h4 bool = _source9.Get_().(D8_DC20).Cf33
				_ = _1183___mcc_h4
				var _1184___mcc_h5 _dafny.Int = _source9.Get_().(D8_DC20).Cf34
				_ = _1184___mcc_h5
				var _1185_cf34 _dafny.Int = _1184___mcc_h5
				_ = _1185_cf34
				var _1186_cf33 bool = _1183___mcc_h4
				_ = _1186_cf33
				(globalState).F21 = Companion_Default___.Fm1((_this).F26(), Companion_Default___.Fm1((_1052_v5).F23(), !(_1186_cf33), globalState), globalState)
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen(((_1134_v66).F27()), 0))
				_ = _index219
				((_1134_v66).F27()).ArraySet1(_dafny.IntOfInt64(953), (_index219).Int())
				var _1187_v100 _dafny.Set
				_ = _1187_v100
				_1187_v100 = _dafny.SetOf(_1048_v1)
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen(((_1134_v66).F27()), 0))
				_ = _index220
				var _rhs230 _dafny.CodePoint = (func() _dafny.CodePoint {
					if _1186_cf33 {
						return (_1054_v7).Select((Companion_Default___.SafeIndex(_1131_i7, _dafny.IntOfUint32((_1054_v7).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
					return _dafny.CodePoint('b')
				})()
				_ = _rhs230
				var _rhs231 _dafny.Int = Companion_Default___.Fm0(Companion_Default___.SafeDivisionInt((_1187_v100).Cardinality(), _dafny.IntOfUint32((_1052_v5.F29()).Cardinality())), _1186_cf33, globalState)
				_ = _rhs231
				var _rhs232 _dafny.CodePoint = Companion_Default___.Fm6((_this).F25(), globalState)
				_ = _rhs232
				var _lhs208 _dafny.Array = (_1134_v66).F27()
				_ = _lhs208
				var _lhs209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen(((_1134_v66).F27()), 0))
				_ = _lhs209
				var _lhs210 *GlobalState = globalState
				_ = _lhs210
				_1179_v99 = _rhs230
				(_lhs208).ArraySet1(_rhs231, (_lhs209).Int())
				_lhs210.F11 = _rhs232
				var _1188_v101 _dafny.Array
				_ = _1188_v101
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_32
				var _nw192 _dafny.Array
				_ = _nw192
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw192 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) bool = (func(_1189_v5 T1) func(_dafny.Int) bool {
						return func(_1190_i16 _dafny.Int) bool {
							return (_1189_v5).F28()
						}
					})(_1052_v5)
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
				_1188_v101 = _nw192
				var _1191_v102 _dafny.Map
				_ = _1191_v102
				_1191_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1188_v101, (((_1134_v66).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen(((_1134_v66).F27()), 0))).Int()).(_dafny.Int)).Cmp((_dafny.MultiSetOf(_1048_v1, (_1052_v5).F28())).Cardinality()) <= 0)
				var _1192_v103 _dafny.Map
				_ = _1192_v103
				_1192_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v5).F28(), ((_1134_v66).F27()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen(((_1134_v66).F27()), 0))).Int()).(_dafny.Int))
				_1191_v102 = (_1191_v102).Update(_1188_v101, (_1134_v66).Fm10(_1054_v7, (_1052_v5).F28(), _1192_v103, _1048_v1, globalState))
				var _rhs233 _dafny.CodePoint = _dafny.CodePoint('v')
				_ = _rhs233
				var _rhs234 _dafny.Int = (_dafny.Zero).Minus((_this).F26())
				_ = _rhs234
				var _lhs211 *GlobalState = globalState
				_ = _lhs211
				var _lhs212 *GlobalState = globalState
				_ = _lhs212
				_lhs211.F11 = _rhs233
				_lhs212.F8 = _rhs234
			} else if _source9.Is_DC21() {
				var _1193___mcc_h6 bool = _source9.Get_().(D8_DC21).Cf35
				_ = _1193___mcc_h6
				var _1194___mcc_h7 bool = _source9.Get_().(D8_DC21).Cf36
				_ = _1194___mcc_h7
				var _1195___mcc_h8 _dafny.Int = _source9.Get_().(D8_DC21).Cf37
				_ = _1195___mcc_h8
				var _1196_cf37 _dafny.Int = _1195___mcc_h8
				_ = _1196_cf37
				var _1197_cf36 bool = _1194___mcc_h7
				_ = _1197_cf36
				var _1198_cf35 bool = _1193___mcc_h6
				_ = _1198_cf35
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen(((_1134_v66).F27()), 0))
				_ = _index221
				((_1134_v66).F27()).ArraySet1(_1196_cf37, (_index221).Int())
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen(((_1134_v66).F27()), 0))
				_ = _index222
				((_1134_v66).F27()).ArraySet1((_dafny.Zero).Minus(_1131_i7), (_index222).Int())
				var _1199_v104 _dafny.Array
				_ = _1199_v104
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_33
				var _nw193 _dafny.Array
				_ = _nw193
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw193 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) bool = (func(_1200_v5 T1) func(_dafny.Int) bool {
						return func(_1201_i17 _dafny.Int) bool {
							return (_1200_v5).F28()
						}
					})(_1052_v5)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw193 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw193).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw193).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1199_v104 = _nw193
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_1199_v104), 0))
				_ = _index223
				(_1199_v104).ArraySet1((_1052_v5).F28(), (_index223).Int())
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_1199_v104), 0))
				_ = _index224
				(_1199_v104).ArraySet1(true, (_index224).Int())
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_1199_v104), 0))
				_ = _index225
				(_1199_v104).ArraySet1(!(!_dafny.Companion_Sequence_.Contains(_1054_v7, _1179_v99)), (_index225).Int())
				var _1202_v105 _dafny.Map
				_ = _1202_v105
				_1202_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1199_v104, _1198_cf35)
				_1202_v105 = _1202_v105
			} else if _source9.Is_DC22() {
				var _1203___mcc_h9 bool = _source9.Get_().(D8_DC22).Cf38
				_ = _1203___mcc_h9
				var _1204_cf38 bool = _1203___mcc_h9
				_ = _1204_cf38
				var _1205_v106 _dafny.Array
				_ = _1205_v106
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_34
				var _nw194 _dafny.Array
				_ = _nw194
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw194 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) bool = (func(_1206_v5 T1) func(_dafny.Int) bool {
						return func(_1207_i18 _dafny.Int) bool {
							return (_1206_v5).F28()
						}
					})(_1052_v5)
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw194 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw194).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw194).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1205_v106 = _nw194
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_1205_v106), 0))
				_ = _index226
				(_1205_v106).ArraySet1(_1048_v1, (_index226).Int())
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_1205_v106), 0))
				_ = _index227
				(_1205_v106).ArraySet1(_1204_cf38, (_index227).Int())
				(globalState).F11 = _dafny.CodePoint('l')
				var _1208_v107 *C0
				_ = _1208_v107
				var _nw195 *C0 = New_C0_()
				_ = _nw195
				_nw195.Ctor__()
				_1208_v107 = _nw195
				_1051_v4 = (_1051_v4).Update(!(true), _dafny.Companion_Sequence_.Contains(_1052_v5.F29(), _1131_i7))
			} else if _source9.Is_DC19() {
				var _1209___mcc_h10 _dafny.Sequence = _source9.Get_().(D8_DC19).Cf32
				_ = _1209___mcc_h10
				var _1210_cf32 _dafny.Sequence = _1209___mcc_h10
				_ = _1210_cf32
				var _1211_v108 _dafny.Map
				_ = _1211_v108
				_1211_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vflk"), (_1052_v5).F28())
				var _1212_v109 _dafny.Array
				_ = _1212_v109
				var _nwElement0_40 bool = _1048_v1
				_ = _nwElement0_40
				var _nw196 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(21))
				_ = _nw196
				(_nw196).ArraySet1(_nwElement0_40, 0)
				(_nw196).ArraySet1((_1052_v5).F28(), 1)
				(_nw196).ArraySet1((_1052_v5).F28(), 2)
				(_nw196).ArraySet1((_1052_v5).F28(), 3)
				(_nw196).ArraySet1((_1052_v5).F28(), 4)
				(_nw196).ArraySet1((_1052_v5).F28(), 5)
				(_nw196).ArraySet1(_1048_v1, 6)
				(_nw196).ArraySet1((_1052_v5).F28(), 7)
				(_nw196).ArraySet1((_1052_v5).F28(), 8)
				(_nw196).ArraySet1((_1052_v5).F28(), 9)
				(_nw196).ArraySet1((Companion_D8_.Create_DC20_((_1052_v5).F28(), (_this).F23())).Dtor_cf33(), 10)
				(_nw196).ArraySet1(false, 11)
				(_nw196).ArraySet1(_1048_v1, 12)
				(_nw196).ArraySet1(true, 13)
				(_nw196).ArraySet1(_1048_v1, 14)
				(_nw196).ArraySet1((_1052_v5).F28(), 15)
				(_nw196).ArraySet1((func() bool {
					if (_1211_v108).Contains(_1054_v7) {
						return (_1211_v108).Get(_1054_v7).(bool)
					}
					return (_1052_v5).F28()
				})(), 16)
				(_nw196).ArraySet1(Companion_Default___.Fm1(_1131_i7, _1048_v1, globalState), 17)
				(_nw196).ArraySet1(_1048_v1, 18)
				(_nw196).ArraySet1((_1052_v5).F28(), 19)
				(_nw196).ArraySet1(!((_1052_v5).F28()), 20)
				_1212_v109 = _nw196
				var _1213_v110 _dafny.Map
				_ = _1213_v110
				_1213_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1212_v109, _1179_v99)
				(globalState).F11 = (func() _dafny.CodePoint {
					if (_1213_v110).Contains(_1212_v109) {
						return (_1213_v110).Get(_1212_v109).(_dafny.CodePoint)
					}
					return _1179_v99
				})()
				var _1214_v111 D0
				_ = _1214_v111
				var _1215_v112 _dafny.Map
				_ = _1215_v112
				var _1216_v113 _dafny.Int
				_ = _1216_v113
				var _out40 D0
				_ = _out40
				var _out41 _dafny.Map
				_ = _out41
				var _out42 _dafny.Int
				_ = _out42
				_out40, _out41, _out42 = Companion_Default___.M0(_1212_v109, _dafny.Companion_Sequence_.Concatenate(_1054_v7, _1054_v7), _1131_i7, globalState)
				_1214_v111 = _out40
				_1215_v112 = _out41
				_1216_v113 = _out42
				(globalState).F14 = ((_this).F25()).Minus(_1131_i7)
				var _1217_v114 _dafny.Array
				_ = _1217_v114
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_35
				var _nw197 _dafny.Array
				_ = _nw197
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw197 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) _dafny.CodePoint = (func(_1218_v99 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1219_i19 _dafny.Int) _dafny.CodePoint {
							return _1218_v99
						}
					})(_1179_v99)
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw197 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw197).ArraySet1CodePoint(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw197).ArraySet1CodePoint(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1217_v114 = _nw197
				var _1220_v115 _dafny.Sequence
				_ = _1220_v115
				_1220_v115 = _dafny.SeqOf((_1052_v5).F28())
				var _1221_v116 D3
				_ = _1221_v116
				_1221_v116 = Companion_D3_.Create_DC12_(_1220_v115)
				var _1222_v117 _dafny.Map
				_ = _1222_v117
				_1222_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D9_.Create_DC24_(_1217_v114), _1221_v116)
				_1222_v117 = _1222_v117
			} else {
				var _1223___mcc_h11 D8 = _source9.Get_().(D8_DC23).Cf39
				_ = _1223___mcc_h11
				var _1224_cf39 D8 = _1223___mcc_h11
				_ = _1224_cf39
				var _1225_v118 _dafny.Map
				_ = _1225_v118
				_1225_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v5).F28(), (_1052_v5).F23())
				var _1226_v119 _dafny.Sequence
				_ = _1226_v119
				_1226_v119 = _dafny.SeqOf(_1225_v118)
				_1051_v4 = (_1051_v4).Update((_1134_v66).Fm10(_1054_v7, _1048_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1134_v66).Fm10(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-438))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg78 _dafny.Int) interface{} {
						return coer78(arg78)
					}
				}((func(_1227_v99 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1228_i20 _dafny.Int) _dafny.CodePoint {
						return _1227_v99
					}
				})(_1179_v99))), (_1052_v5).F28(), _1225_v118, true, globalState), _dafny.IntOfInt64(922))).Update((_1052_v5).F28(), Companion_Default___.Fm0(((_1226_v119).Select((Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_1226_v119).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), (_1052_v5).F28(), globalState)), (_1052_v5).F28(), globalState), _1048_v1)
				var _1229_v120 _dafny.Array
				_ = _1229_v120
				var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
				_ = _nw198
				_1229_v120 = _nw198
				var _1230_v121 _dafny.Array
				_ = _1230_v121
				var _nw199 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
				_ = _nw199
				_1230_v121 = _nw199
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_1229_v120), 0))
				_ = _index228
				(_1229_v120).ArraySet1(_1230_v121, (_index228).Int())
				var _1231_v122 _dafny.MultiSet
				_ = _1231_v122
				_1231_v122 = _dafny.MultiSetOf(_1179_v99)
				var _1232_v123 _dafny.Map
				_ = _1232_v123
				_1232_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_1052_v5).F28())
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_1229_v120), 0))
				_ = _index229
				var _rhs235 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1055_v8).Cardinality()), (_this).F25()), _dafny.IntOfInt64(491))
				_ = _rhs235
				var _rhs236 bool = (func() bool {
					if (_1051_v4).Contains(!((_1052_v5).F28())) {
						return (_1051_v4).Get(!((_1052_v5).F28())).(bool)
					}
					return _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_1054_v7, (Companion_Default___.SafeIndex((_1052_v5).F23(), _dafny.IntOfUint32((_1054_v7).Cardinality()))).Uint32(), _1179_v99), _1179_v99)
				})()
				_ = _rhs236
				var _rhs237 _dafny.Array = _1230_v121
				_ = _rhs237
				var _rhs238 _dafny.Int = (func() _dafny.Int {
					if (_1231_v122).Contains(_1179_v99) {
						return (_1231_v122).Multiplicity(_1179_v99)
					}
					return _dafny.IntOfInt64(645)
				})()
				_ = _rhs238
				var _rhs239 _dafny.Map = (((_1232_v123).Update(_dafny.IntOfInt64(-753), (_1052_v5).F28())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1131_i7, (_1052_v5).F28())).Update((_this).F25(), _1048_v1))).Update(Companion_Default___.SafeDivisionInt((func() _dafny.Map {
					var _coll45 = _dafny.NewMapBuilder()
					_ = _coll45
					for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-391), _dafny.IntOfInt64(524))); ; {
						_compr_45, _ok49 := _iter49()
						if !_ok49 {
							break
						}
						var _1233_v124 _dafny.Int
						_1233_v124 = interface{}(_compr_45).(_dafny.Int)
						if ((_dafny.IntOfInt64(-391)).Cmp(_1233_v124) <= 0) && ((_1233_v124).Cmp(_dafny.IntOfInt64(524)) < 0) {
							_coll45.Add(Companion_Default___.SafeModuloInt(_1233_v124, _dafny.IntOfInt64(909)), (_dafny.Zero).Minus((_1052_v5).F23()))
						}
					}
					return _coll45.ToMap()
				}()).Cardinality(), (_1052_v5).F23()), true)
				_ = _rhs239
				var _lhs213 *GlobalState = globalState
				_ = _lhs213
				var _lhs214 _dafny.Array = _1229_v120
				_ = _lhs214
				var _lhs215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_1229_v120), 0))
				_ = _lhs215
				var _lhs216 *GlobalState = globalState
				_ = _lhs216
				_lhs213.F14 = _rhs235
				_1048_v1 = _rhs236
				(_lhs214).ArraySet1(_rhs237, (_lhs215).Int())
				_lhs216.F14 = _rhs238
				r2 = _rhs239
				_1232_v123 = (_1232_v123).Update(_1131_i7, (_1052_v5).F28())
				var _1234_v125 _dafny.Sequence
				_ = _1234_v125
				_1234_v125 = _dafny.SeqOf((_1052_v5).F28())
				var _1235_v126 _dafny.Map
				_ = _1235_v126
				_1235_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1234_v125, ((_this).F23()).Cmp((_this).F23()) < 0)
				var _rhs240 _dafny.Sequence = _1054_v7
				_ = _rhs240
				var _rhs241 bool = (_1052_v5).F28()
				_ = _rhs241
				var _rhs242 _dafny.Map = (_1235_v126).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1234_v125, (_1052_v5).F28()))
				_ = _rhs242
				var _rhs243 bool = !((_1052_v5).F28())
				_ = _rhs243
				var _rhs244 _dafny.Int = ((_dafny.IntOfInt64(-332)).Plus((_this).F23())).Plus((_1052_v5).F23())
				_ = _rhs244
				var _lhs217 *GlobalState = globalState
				_ = _lhs217
				var _lhs218 *GlobalState = globalState
				_ = _lhs218
				var _lhs219 *GlobalState = globalState
				_ = _lhs219
				_lhs217.F6 = _rhs240
				r0 = _rhs241
				_1235_v126 = _rhs242
				_lhs218.F7 = _rhs243
				_lhs219.F8 = _rhs244
			}
		}
		var _hi14 _dafny.Int = Companion_Default___.SafeDivisionInt((_1052_v5).F23(), (_this).F26())
		_ = _hi14
		for _1236_i21 := (_dafny.Zero).Minus((_1052_v5).F23()); _1236_i21.Cmp(_hi14) < 0; _1236_i21 = _1236_i21.Plus(_dafny.One) {
			(globalState).F14 = Companion_Default___.Fm0((_this).F25(), (func() bool {
				if (_1051_v4).Contains((_1052_v5).F28()) {
					return (_1051_v4).Get((_1052_v5).F28()).(bool)
				}
				return (_1052_v5).F28()
			})(), globalState)
			var _1237_v127 D12
			_ = _1237_v127
			_1237_v127 = Companion_D12_.Create_DC35_()
			var _1238_v128 _dafny.Array
			_ = _1238_v128
			var _nw200 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw200
			_1238_v128 = _nw200
			var _1239_v129 _dafny.Map
			_ = _1239_v129
			_1239_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v127, _1238_v128)
			var _1240_v130 _dafny.Int
			_ = _1240_v130
			var _1241_v131 _dafny.Set
			_ = _1241_v131
			var _out43 _dafny.Int
			_ = _out43
			var _out44 _dafny.Set
			_ = _out44
			_out43, _out44 = (_1052_v5).M1((func() _dafny.Array {
				if (_1239_v129).Contains(_1237_v127) {
					return (_1239_v129).Get(_1237_v127).(_dafny.Array)
				}
				return _1238_v128
			})(), ((_1052_v5).F28()) && ((_1052_v5).F28()), globalState)
			_1240_v130 = _out43
			_1241_v131 = _out44
			var _1242_v132 _dafny.Array
			_ = _1242_v132
			var _len0_36 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_36
			var _nw201 _dafny.Array
			_ = _nw201
			if _len0_36.Cmp(_dafny.Zero) == 0 {
				_nw201 = _dafny.NewArray(_len0_36)
			} else {
				var _init36 func(_dafny.Int) _dafny.Int = (func(_1243_v130 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1244_i22 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1244_i22, _1243_v130)
					}
				})(_1240_v130)
				_ = _init36
				var _element0_36 = _init36(_dafny.Zero)
				_ = _element0_36
				_nw201 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
				(_nw201).ArraySet1(_element0_36, 0)
				var _nativeLen0_36 = (_len0_36).Int()
				_ = _nativeLen0_36
				for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
					(_nw201).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
				}
			}
			_1242_v132 = _nw201
			var _1245_v133 _dafny.Sequence
			_ = _1245_v133
			_1245_v133 = _dafny.SeqOf(_1242_v132)
			var _1246_v134 D1
			_ = _1246_v134
			_1246_v134 = Companion_D1_.Create_DC3_(_1245_v133)
			var _source10 D1 = _1246_v134
			_ = _source10
			if _source10.Is_DC4() {
				var _1247___mcc_h12 _dafny.Map = _source10.Get_().(D1_DC4).Cf6
				_ = _1247___mcc_h12
				var _1248___mcc_h13 bool = _source10.Get_().(D1_DC4).Cf7
				_ = _1248___mcc_h13
				var _1249___mcc_h14 bool = _source10.Get_().(D1_DC4).Cf8
				_ = _1249___mcc_h14
				var _1250_cf8 bool = _1249___mcc_h14
				_ = _1250_cf8
				var _1251_cf7 bool = _1248___mcc_h13
				_ = _1251_cf7
				var _1252_cf6 _dafny.Map = _1247___mcc_h12
				_ = _1252_cf6
				var _1253_v135 *C3
				_ = _1253_v135
				var _nw202 *C3 = New_C3_()
				_ = _nw202
				_nw202.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(67))).Uint32(), func(coer79 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg79 _dafny.Int) interface{} {
						return coer79(arg79)
					}
				}(func(_1254_i23 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(true, true)
				})), (_1052_v5).F23(), (_1052_v5).F28(), _1052_v5.F29())
				_1253_v135 = _nw202
				(globalState).F14 = (_this).F23()
				var _1255_v136 *C4
				_ = _1255_v136
				var _nw203 *C4 = New_C4_()
				_ = _nw203
				_nw203.Ctor__(_1242_v132, _1236_i21)
				_1255_v136 = _nw203
				var _1256_v137 D3
				_ = _1256_v137
				_1256_v137 = Companion_D3_.Create_DC12_(_dafny.SeqOf(_1048_v1))
				var _1257_v138 _dafny.Map
				_ = _1257_v138
				_1257_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1256_v137, (_1052_v5).F28())
				var _1258_v139 _dafny.Sequence
				_ = _1258_v139
				_1258_v139 = _dafny.SeqOf(!((_1052_v5).F28()), !(_1250_cf8))
				var _pat_let_tv39 = _1258_v139
				_ = _pat_let_tv39
				var _pat_let_tv40 = _1258_v139
				_ = _pat_let_tv40
				(globalState).F21 = (func() bool {
					if (_1257_v138).Contains(func(_pat_let20_0 D3) D3 {
						return func(_1259_dt__update__tmp_h1 D3) D3 {
							return func(_pat_let21_0 _dafny.Sequence) D3 {
								return func(_1260_dt__update_hcf23_h0 _dafny.Sequence) D3 {
									return Companion_D3_.Create_DC12_(_1260_dt__update_hcf23_h0)
								}(_pat_let21_0)
							}(_pat_let_tv39)
						}(_pat_let20_0)
					}(_1256_v137)) {
						return (_1257_v138).Get(func(_pat_let22_0 D3) D3 {
							return func(_1261_dt__update__tmp_h2 D3) D3 {
								return func(_pat_let23_0 _dafny.Sequence) D3 {
									return func(_1262_dt__update_hcf23_h1 _dafny.Sequence) D3 {
										return Companion_D3_.Create_DC12_(_1262_dt__update_hcf23_h1)
									}(_pat_let23_0)
								}(_pat_let_tv40)
							}(_pat_let22_0)
						}(_1256_v137)).(bool)
					}
					return true
				})()
			} else if _source10.Is_DC5() {
				var _1263___mcc_h15 _dafny.Int = _source10.Get_().(D1_DC5).Cf9
				_ = _1263___mcc_h15
				var _1264___mcc_h16 bool = _source10.Get_().(D1_DC5).Cf10
				_ = _1264___mcc_h16
				var _1265___mcc_h17 _dafny.Int = _source10.Get_().(D1_DC5).Cf11
				_ = _1265___mcc_h17
				var _1266___mcc_h18 _dafny.Int = _source10.Get_().(D1_DC5).Cf12
				_ = _1266___mcc_h18
				var _1267_cf12 _dafny.Int = _1266___mcc_h18
				_ = _1267_cf12
				var _1268_cf11 _dafny.Int = _1265___mcc_h17
				_ = _1268_cf11
				var _1269_cf10 bool = _1264___mcc_h16
				_ = _1269_cf10
				var _1270_cf9 _dafny.Int = _1263___mcc_h15
				_ = _1270_cf9
				(globalState).F14 = (_1236_i21).Minus((_1052_v5).F23())
				var _1271_v140 _dafny.Sequence
				_ = _1271_v140
				_1271_v140 = _dafny.SeqOf(_1048_v1, (_1052_v5).F28())
				(globalState).F21 = (func() bool {
					if _dafny.Companion_Sequence_.IsPrefixOf(_1271_v140, _1271_v140) {
						return _1269_cf10
					}
					return ((_this).F26()).Cmp(_1236_i21) <= 0
				})()
				_1055_v8 = _1055_v8
				var _1272_v141 *C0
				_ = _1272_v141
				var _nw204 *C0 = New_C0_()
				_ = _nw204
				_nw204.Ctor__()
				_1272_v141 = _nw204
			} else if _source10.Is_DC6() {
				var _1273___mcc_h19 _dafny.Array = _source10.Get_().(D1_DC6).Cf13
				_ = _1273___mcc_h19
				var _1274_cf13 _dafny.Array = _1273___mcc_h19
				_ = _1274_cf13
				(globalState).F8 = (_dafny.Zero).Minus((_this).F23())
				(globalState).F9 = _1048_v1
				_1274_cf13 = (func() _dafny.Array {
					if (_1052_v5).F28() {
						return _1274_cf13
					}
					return _1274_cf13
				})()
				var _1275_v142 _dafny.Map
				_ = _1275_v142
				_1275_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1052_v5.F29(), _1236_i21)
				_1275_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm0(_dafny.IntOfInt64(455), _1048_v1, globalState), (_this).F25()), _dafny.SeqOf((_1052_v5).F23())), _dafny.IntOfInt64(993))
			} else {
				var _1276___mcc_h20 _dafny.Sequence = _source10.Get_().(D1_DC3).Cf5
				_ = _1276___mcc_h20
				var _1277_cf5 _dafny.Sequence = _1276___mcc_h20
				_ = _1277_cf5
				var _1278_v143 _dafny.Map
				_ = _1278_v143
				_1278_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v5).F28(), _1240_v130)
				var _1279_v144 _dafny.Sequence
				_ = _1279_v144
				_1279_v144 = _dafny.SeqOf((_1052_v5).F28())
				_1278_v143 = (_1278_v143).Update(_1048_v1, _dafny.IntOfUint32((_1279_v144).Cardinality()))
				var _1280_v145 _dafny.MultiSet
				_ = _1280_v145
				_1280_v145 = _dafny.MultiSetOf((_1052_v5).F28())
				(globalState).F14 = Companion_Default___.SafeDivisionInt((_1280_v145).Cardinality(), ((_dafny.Zero).Minus((_this).F23())).Plus((_dafny.MultiSetOf((_1052_v5).F23())).Cardinality()))
				var _1281_v146 _dafny.Array
				_ = _1281_v146
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_37
				var _nw205 _dafny.Array
				_ = _nw205
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw205 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) _dafny.Sequence = (func(_1282_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1283_i24 _dafny.Int) _dafny.Sequence {
							return _1282_v7
						}
					})(_1054_v7)
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw205 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw205).ArraySet1(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw205).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1281_v146 = _nw205
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_1281_v146), 0))
				_ = _index230
				(_1281_v146).ArraySet1(_1054_v7, (_index230).Int())
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_1281_v146), 0))
				_ = _index231
				(_1281_v146).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-84))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg80 _dafny.Int) interface{} {
						return coer80(arg80)
					}
				}(func(_1284_i25 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('x')
				})), _1054_v7), (_index231).Int())
				_1054_v7 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1054_v7, (_1281_v146).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_1281_v146), 0))).Int()).(_dafny.Sequence)), _1054_v7)
			}
			if _1048_v1 {
				var _1285_v147 _dafny.Array
				_ = _1285_v147
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_38
				var _nw206 _dafny.Array
				_ = _nw206
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw206 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) _dafny.Set = (func(_1286_v131 _dafny.Set) func(_dafny.Int) _dafny.Set {
						return func(_1287_i26 _dafny.Int) _dafny.Set {
							return _1286_v131
						}
					})(_1241_v131)
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw206 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw206).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw206).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1285_v147 = _nw206
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_1285_v147), 0))
				_ = _index232
				(_1285_v147).ArraySet1(func() _dafny.Set {
					var _coll46 = _dafny.NewBuilder()
					_ = _coll46
					for _iter50 := _dafny.Iterate((func() _dafny.Set {
						var _coll47 = _dafny.NewBuilder()
						_ = _coll47
						for _iter51 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(836), _dafny.IntOfInt64(653))); ; {
							_compr_47, _ok51 := _iter51()
							if !_ok51 {
								break
							}
							var _1288_v148 _dafny.Int
							_1288_v148 = interface{}(_compr_47).(_dafny.Int)
							if ((_dafny.IntOfInt64(836)).Cmp(_1288_v148) <= 0) && ((_1288_v148).Cmp(_dafny.IntOfInt64(653)) < 0) {
								_coll47.Add((_1288_v148).Times(_dafny.IntOfInt64(25)))
							}
						}
						return _coll47.ToSet()
					}()).Elements()); ; {
						_compr_46, _ok50 := _iter50()
						if !_ok50 {
							break
						}
						var _1289_v149 _dafny.Int
						_1289_v149 = interface{}(_compr_46).(_dafny.Int)
						if (func() _dafny.Set {
							var _coll48 = _dafny.NewBuilder()
							_ = _coll48
							for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(836), _dafny.IntOfInt64(653))); ; {
								_compr_48, _ok52 := _iter52()
								if !_ok52 {
									break
								}
								var _1290_v148 _dafny.Int
								_1290_v148 = interface{}(_compr_48).(_dafny.Int)
								if ((_dafny.IntOfInt64(836)).Cmp(_1290_v148) <= 0) && ((_1290_v148).Cmp(_dafny.IntOfInt64(653)) < 0) {
									_coll48.Add((_1290_v148).Times(_dafny.IntOfInt64(25)))
								}
							}
							return _coll48.ToSet()
						}()).Contains(_1289_v149) {
							_coll46.Add(Companion_Default___.SafeModuloInt(_1289_v149, (_dafny.MultiSetOf(_dafny.CodePoint('b'), _dafny.CodePoint('e'))).Cardinality()))
						}
					}
					return _coll46.ToSet()
				}(), (_index232).Int())
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_1285_v147), 0))
				_ = _index233
				(_1285_v147).ArraySet1(_1241_v131, (_index233).Int())
				var _1291_v150 _dafny.Int
				_ = _1291_v150
				var _1292_v151 _dafny.Sequence
				_ = _1292_v151
				var _out45 _dafny.Int
				_ = _out45
				var _out46 _dafny.Sequence
				_ = _out46
				_out45, _out46 = (_1052_v5).M2(_dafny.Companion_Sequence_.Concatenate(_1054_v7, _1054_v7), globalState)
				_1291_v150 = _out45
				_1292_v151 = _out46
				var _1293_v152 D0
				_ = _1293_v152
				_1293_v152 = Companion_D0_.Create_DC1_((_dafny.Zero).Minus(_1240_v130), _1048_v1, (_this).F26())
				var _rhs245 _dafny.Int = (_1052_v5).F23()
				_ = _rhs245
				var _rhs246 bool = (_1052_v5).F28()
				_ = _rhs246
				var _rhs247 bool = (_1293_v152).Dtor_cf2()
				_ = _rhs247
				var _lhs220 *GlobalState = globalState
				_ = _lhs220
				var _lhs221 *GlobalState = globalState
				_ = _lhs221
				_lhs220.F14 = _rhs245
				_1048_v1 = _rhs246
				_lhs221.F21 = _rhs247
				var _1294_v153 _dafny.Map
				_ = _1294_v153
				_1294_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v5).F23(), (_1052_v5).F28())
				var _1295_v154 _dafny.Map
				_ = _1295_v154
				_1295_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1048_v1, _1240_v130)
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_1238_v128), 0))
				_ = _index234
				(_1238_v128).ArraySet1(!(((_1294_v153).Cardinality()).Cmp((_dafny.Zero).Minus((func() _dafny.Int {
					if (_1295_v154).Contains((_1052_v5).F28()) {
						return (_1295_v154).Get((_1052_v5).F28()).(_dafny.Int)
					}
					return (_1052_v5).F23()
				})())) > 0), (_index234).Int())
				var _1296_v155 _dafny.CodePoint
				_ = _1296_v155
				_1296_v155 = _dafny.CodePoint('u')
				var _1297_v156 D11
				_ = _1297_v156
				_1297_v156 = Companion_D11_.Create_DC32_(_dafny.UnicodeSeqOfUtf8Bytes("mxdqyef"))
				var _1298_v157 _dafny.Array
				_ = _1298_v157
				var _nwElement0_41 _dafny.Sequence = _1054_v7
				_ = _nwElement0_41
				var _nw207 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(26))
				_ = _nw207
				(_nw207).ArraySet1(_nwElement0_41, 0)
				(_nw207).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1054_v7, _1054_v7), 1)
				(_nw207).ArraySet1(_1054_v7, 2)
				(_nw207).ArraySet1((_1055_v8).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0((_1052_v5).F23(), (_1052_v5).F28(), globalState), _dafny.IntOfUint32((_1055_v8).Cardinality()))).Uint32()).(_dafny.Sequence), 3)
				(_nw207).ArraySet1(_dafny.Companion_Sequence_.Update(_1054_v7, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_1055_v8).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_1055_v8).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_1054_v7).Cardinality()))).Uint32(), _1296_v155), 4)
				(_nw207).ArraySet1(_1054_v7, 5)
				(_nw207).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1054_v7, _dafny.Companion_Sequence_.Update(_1054_v7, (Companion_Default___.SafeIndex(_1291_v150, _dafny.IntOfUint32((_1054_v7).Cardinality()))).Uint32(), _1296_v155)), 6)
				(_nw207).ArraySet1(_1054_v7, 7)
				(_nw207).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(525))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg81 _dafny.Int) interface{} {
						return coer81(arg81)
					}
				}(func(_1299_i27 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				})), 8)
				(_nw207).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tfynwpxj"), _1054_v7), 9)
				(_nw207).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1054_v7, _1054_v7), 10)
				(_nw207).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(62))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg82 _dafny.Int) interface{} {
						return coer82(arg82)
					}
				}((func(_1300_v155 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1301_i28 _dafny.Int) _dafny.CodePoint {
						return _1300_v155
					}
				})(_1296_v155))), 11)
				(_nw207).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-834))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}((func(_1302_v155 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1303_i29 _dafny.Int) _dafny.CodePoint {
						return _1302_v155
					}
				})(_1296_v155))), 12)
				(_nw207).ArraySet1((_1297_v156).Dtor_cf53(), 13)
				(_nw207).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("u"), 14)
				(_nw207).ArraySet1(Companion_Default___.Fm2(globalState), 15)
				(_nw207).ArraySet1(_1054_v7, 16)
				(_nw207).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1054_v7, _1054_v7), 17)
				(_nw207).ArraySet1(_1054_v7, 18)
				(_nw207).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("g"), 19)
				(_nw207).ArraySet1(_1054_v7, 20)
				(_nw207).ArraySet1(_1054_v7, 21)
				(_nw207).ArraySet1(_1054_v7, 22)
				(_nw207).ArraySet1(_1054_v7, 23)
				(_nw207).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("esqnlyvg"), 24)
				(_nw207).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tkq"), _1054_v7), 25)
				_1298_v157 = _nw207
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_1298_v157), 0))
				_ = _index235
				(_1298_v157).ArraySet1((func() _dafny.Sequence {
					if (_1052_v5).F28() {
						return _dafny.UnicodeSeqOfUtf8Bytes("nxukkyffw")
					}
					return _1054_v7
				})(), (_index235).Int())
				var _1304_v158 _dafny.Array
				_ = _1304_v158
				var _nwElement0_42 _dafny.Map = _1047_v0
				_ = _nwElement0_42
				var _nw208 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.One)
				_ = _nw208
				(_nw208).ArraySet1(_nwElement0_42, 0)
				_1304_v158 = _nw208
				var _1305_v159 D11
				_ = _1305_v159
				_1305_v159 = Companion_D11_.Create_DC30_(_1304_v158)
				var _1306_v160 _dafny.Map
				_ = _1306_v160
				_1306_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
					if true {
						return _1296_v155
					}
					return _dafny.CodePoint('e')
				})(), (_1052_v5).F28())
				var _1307_v161 _dafny.Map
				_ = _1307_v161
				_1307_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1054_v7, _1054_v7)
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_1238_v128), 0))
				_ = _index236
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_1298_v157), 0))
				_ = _index237
				var _rhs248 bool = !(((_1052_v5).F23()).Cmp((_1052_v5).F23()) < 0)
				_ = _rhs248
				var _rhs249 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_1307_v161).Contains(_dafny.UnicodeSeqOfUtf8Bytes("culpysq")) {
						return (_1307_v161).Get(_dafny.UnicodeSeqOfUtf8Bytes("culpysq")).(_dafny.Sequence)
					}
					return _1054_v7
				})(), _1054_v7), _dafny.Companion_Sequence_.Concatenate(_1054_v7, _1054_v7))
				_ = _rhs249
				var _rhs250 D11 = _1305_v159
				_ = _rhs250
				var _rhs251 _dafny.Map = (_1306_v160).Merge(_1306_v160)
				_ = _rhs251
				var _lhs222 _dafny.Array = _1238_v128
				_ = _lhs222
				var _lhs223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_1238_v128), 0))
				_ = _lhs223
				var _lhs224 _dafny.Array = _1298_v157
				_ = _lhs224
				var _lhs225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_1298_v157), 0))
				_ = _lhs225
				(_lhs222).ArraySet1(_rhs248, (_lhs223).Int())
				(_lhs224).ArraySet1(_rhs249, (_lhs225).Int())
				_1305_v159 = _rhs250
				_1306_v160 = _rhs251
				_1054_v7 = (_1298_v157).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_1298_v157), 0))).Int()).(_dafny.Sequence)
			} else {
				_1055_v8 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(217))).Uint32(), func(coer84 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}(func(_1308_i30 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("tp")
				}))
				(globalState).F8 = ((_this).F23()).Plus((_1052_v5).F23())
				var _1309_v162 _dafny.Map
				_ = _1309_v162
				_1309_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v5).F28(), (_this).F26())
				var _1310_v163 *C1
				_ = _1310_v163
				var _nw209 *C1 = New_C1_()
				_ = _nw209
				_nw209.Ctor__((_1052_v5).F28(), _dafny.SeqOf((_dafny.MultiSetOf(_1309_v162)).Cardinality()), (_1052_v5).F23())
				_1310_v163 = _nw209
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_1242_v132), 0))
				_ = _index238
				(_1242_v132).ArraySet1((_this).F25(), (_index238).Int())
				var _1311_v164 _dafny.CodePoint
				_ = _1311_v164
				_1311_v164 = _dafny.CodePoint('i')
				var _1312_v165 _dafny.MultiSet
				_ = _1312_v165
				_1312_v165 = _dafny.MultiSetOf(_1311_v164, _1311_v164)
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_1242_v132), 0))
				_ = _index239
				(_1242_v132).ArraySet1((((_this).F26()).Plus((func() _dafny.Int {
					if (_1312_v165).Contains(_1311_v164) {
						return (_1312_v165).Multiplicity(_1311_v164)
					}
					return (_this).F26()
				})())).Plus((_this).F26()), (_index239).Int())
				var _1313_v166 D11
				_ = _1313_v166
				_1313_v166 = Companion_D11_.Create_DC31_(_1236_i21)
				var _1314_v167 _dafny.Sequence
				_ = _1314_v167
				_1314_v167 = _dafny.SeqOf(_1313_v166, _1313_v166, _1313_v166)
				var _1315_v168 _dafny.Map
				_ = _1315_v168
				_1315_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1314_v167, _1052_v5.F29())
				var _1316_v170 _dafny.Set
				_ = _1316_v170
				_1316_v170 = _dafny.SetOf(Companion_Default___.Fm32(globalState))
				var _1317_v172 _dafny.MultiSet
				_ = _1317_v172
				_1317_v172 = _dafny.MultiSetOf(_dafny.SeqOf(_1313_v166), _1314_v167, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1313_v166), (Companion_Default___.SafeIndex((_1052_v5).F23(), _dafny.IntOfUint32((_dafny.SeqOf(_1313_v166)).Cardinality()))).Uint32(), _1313_v166), _1314_v167)
				var _1318_v173 _dafny.Array
				_ = _1318_v173
				var _nwElement0_43 _dafny.Map = (_1315_v168).Update(_1314_v167, _1052_v5.F29())
				_ = _nwElement0_43
				var _nw210 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(26))
				_ = _nw210
				(_nw210).ArraySet1(_nwElement0_43, 0)
				(_nw210).ArraySet1(_1315_v168, 1)
				(_nw210).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1314_v167, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(428))).Uint32(), func(coer85 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg85 _dafny.Int) interface{} {
						return coer85(arg85)
					}
				}((func(_1319_v5 T1, _1320_v130 _dafny.Int, _1321_i21 _dafny.Int, _1322_v132 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_1323_i31 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetOf((_this).F26()), _dafny.MultiSetFromSeq(_1319_v5.F29()), _dafny.MultiSetOf(_1320_v130, _dafny.IntOfInt64(383)), _dafny.MultiSetOf(_1321_i21, (_1322_v132).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_1322_v132), 0))).Int()).(_dafny.Int)))).Cardinality())
					}
				})(_1052_v5, _1240_v130, _1236_i21, _1242_v132)))), 2)
				(_nw210).ArraySet1((_1315_v168).Merge((Companion_D15_.Create_DC43_(_1315_v168)).Dtor_cf67()), 3)
				(_nw210).ArraySet1(func() _dafny.Map {
					var _coll49 = _dafny.NewMapBuilder()
					_ = _coll49
					for _iter53 := _dafny.Iterate((_1316_v170).Elements()); ; {
						_compr_49, _ok53 := _iter53()
						if !_ok53 {
							break
						}
						var _1324_v169 _dafny.Sequence
						_1324_v169 = interface{}(_compr_49).(_dafny.Sequence)
						if (_1316_v170).Contains(_1324_v169) {
							_coll49.Add(_1324_v169, _1052_v5.F29())
						}
					}
					return _coll49.ToMap()
				}(), 4)
				(_nw210).ArraySet1(_1315_v168, 5)
				(_nw210).ArraySet1(_1315_v168, 6)
				(_nw210).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1314_v167, _1052_v5.F29()), 7)
				(_nw210).ArraySet1(_1315_v168, 8)
				(_nw210).ArraySet1(_1315_v168, 9)
				(_nw210).ArraySet1(_1315_v168, 10)
				(_nw210).ArraySet1(_1315_v168, 11)
				(_nw210).ArraySet1(func() _dafny.Map {
					var _coll50 = _dafny.NewMapBuilder()
					_ = _coll50
					for _iter54 := _dafny.Iterate((_1317_v172).Elements()); ; {
						_compr_50, _ok54 := _iter54()
						if !_ok54 {
							break
						}
						var _1325_v171 _dafny.Sequence
						_1325_v171 = interface{}(_compr_50).(_dafny.Sequence)
						if (_1317_v172).Contains(_1325_v171) {
							_coll50.Add(_1325_v171, _1050_v3)
						}
					}
					return _coll50.ToMap()
				}(), 12)
				(_nw210).ArraySet1(_1315_v168, 13)
				(_nw210).ArraySet1((func() _dafny.Map {
					if (_1052_v5).F28() {
						return _1315_v168
					}
					return _1315_v168
				})(), 14)
				(_nw210).ArraySet1(Companion_Default___.Fm33((_dafny.SetOf(_dafny.IntOfUint32((_1052_v5.F29()).Cardinality()), _1236_i21, (_dafny.Zero).Minus((_this).F23()))).Cardinality(), (_1052_v5).F28(), globalState), 15)
				(_nw210).ArraySet1(_1315_v168, 16)
				(_nw210).ArraySet1(_1315_v168, 17)
				(_nw210).ArraySet1(_1315_v168, 18)
				(_nw210).ArraySet1(_1315_v168, 19)
				(_nw210).ArraySet1(_1315_v168, 20)
				(_nw210).ArraySet1((_1315_v168).Merge(_1315_v168), 21)
				(_nw210).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1314_v167, _1050_v3), 22)
				(_nw210).ArraySet1(_1315_v168, 23)
				(_nw210).ArraySet1((_1315_v168).Merge((_1315_v168).Update(_1314_v167, _1052_v5.F29())), 24)
				(_nw210).ArraySet1(_1315_v168, 25)
				_1318_v173 = _nw210
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_1318_v173), 0))
				_ = _index240
				(_1318_v173).ArraySet1(_1315_v168, (_index240).Int())
				var _1326_v174 _dafny.Sequence
				_ = _1326_v174
				_1326_v174 = _dafny.SeqOf((_1052_v5).F28())
				var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_1318_v173), 0))
				_ = _index241
				(_1318_v173).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(836))).Uint32(), func(coer86 func(_dafny.Int) D11) func(_dafny.Int) interface{} {
					return func(arg86 _dafny.Int) interface{} {
						return coer86(arg86)
					}
				}((func(_1327_v166 D11) func(_dafny.Int) D11 {
					return func(_1328_i32 _dafny.Int) D11 {
						return _1327_v166
					}
				})(_1313_v166))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1326_v174).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(836))).Uint32(), func(coer87 func(_dafny.Int) D11) func(_dafny.Int) interface{} {
					return func(arg87 _dafny.Int) interface{} {
						return coer87(arg87)
					}
				}((func(_1329_v166 D11) func(_dafny.Int) D11 {
					return func(_1330_i32 _dafny.Int) D11 {
						return _1329_v166
					}
				})(_1313_v166)))).Cardinality()))).Uint32(), _1313_v166), _dafny.Companion_Sequence_.Concatenate(_1052_v5.F29(), _dafny.SeqOf((_1242_v132).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_1242_v132), 0))).Int()).(_dafny.Int)))), (_index241).Int())
			}
		}
		var _1331_v175 _dafny.Set
		_ = _1331_v175
		_1331_v175 = _dafny.SetOf((func() _dafny.Sequence {
			if (_1052_v5).F28() {
				return _1054_v7
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("seef")
		})())
		var _1332_v176 _dafny.Map
		_ = _1332_v176
		_1332_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _1331_v175)
		var _1333_v177 D16
		_ = _1333_v177
		_1333_v177 = Companion_D16_.Create_DC46_(_1331_v175)
		_1331_v175 = (func() _dafny.Set {
			if (_1332_v176).Contains((_this).F26()) {
				return (_1332_v176).Get((_this).F26()).(_dafny.Set)
			}
			return (func() _dafny.Set {
				if (_1052_v5).F28() {
					return _1331_v175
				}
				return (_1333_v177).Dtor_cf74()
			})()
		})()
		var _1334_v178 _dafny.Map
		_ = _1334_v178
		_1334_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1054_v7, (_1052_v5).F28())
		r0 = !((func() bool {
			if (_1334_v178).Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ulwvqx"), _1054_v7)) {
				return (_1334_v178).Get(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ulwvqx"), _1054_v7)).(bool)
			}
			return (_1052_v5).F28()
		})())
		var _1335_v179 _dafny.Map
		_ = _1335_v179
		_1335_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1048_v1, (_this).F23())
		var _1336_v180 _dafny.Sequence
		_ = _1336_v180
		_1336_v180 = _dafny.SeqOf(_1335_v179, _1335_v179)
		var _1337_v181 _dafny.Set
		_ = _1337_v181
		_1337_v181 = _dafny.SetOf(_dafny.IntOfUint32((_1055_v8).Cardinality()))
		var _1338_v182 _dafny.Set
		_ = _1338_v182
		_1338_v182 = _dafny.SetOf(_dafny.SetOf((_dafny.Zero).Minus((_1337_v181).Cardinality()), _dafny.IntOfInt64(957), _dafny.IntOfInt64(843), (_this).F26()))
		r1 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1336_v180, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(681))).Uint32(), func(coer88 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg88 _dafny.Int) interface{} {
				return coer88(arg88)
			}
		}((func(_1339_v179 _dafny.Map) func(_dafny.Int) _dafny.Map {
			return func(_1340_i33 _dafny.Int) _dafny.Map {
				return _1339_v179
			}
		})(_1335_v179)))), _dafny.Companion_Sequence_.Concatenate(_1336_v180, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v5).F28(), (_this).F25()), _1335_v179, _1335_v179, _1335_v179))), (Companion_Default___.SafeIndex((_1338_v182).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1336_v180, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(681))).Uint32(), func(coer89 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg89 _dafny.Int) interface{} {
				return coer89(arg89)
			}
		}((func(_1341_v179 _dafny.Map) func(_dafny.Int) _dafny.Map {
			return func(_1342_i33 _dafny.Int) _dafny.Map {
				return _1341_v179
			}
		})(_1335_v179)))), _dafny.Companion_Sequence_.Concatenate(_1336_v180, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v5).F28(), (_this).F25()), _1335_v179, _1335_v179, _1335_v179)))).Cardinality()))).Uint32(), _1335_v179)
		r2 = Companion_Default___.Fm5(_1337_v181, globalState)
		return r0, r1, r2
	}
}
func (_this *C5) M4(p0 bool, globalState *GlobalState) (bool, _dafny.Int, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		(globalState).F7 = p0
		var _1343_v0 D11
		_ = _1343_v0
		_1343_v0 = Companion_D11_.Create_DC31_((_this).F25())
		var _1344_v1 _dafny.Sequence
		_ = _1344_v1
		_1344_v1 = _dafny.SeqOf(_1343_v0, _1343_v0)
		var _1345_v2 _dafny.Sequence
		_ = _1345_v2
		_1345_v2 = _dafny.UnicodeSeqOfUtf8Bytes("fvpo")
		var _1346_v3 _dafny.Sequence
		_ = _1346_v3
		_1346_v3 = _dafny.SeqOf(_dafny.IntOfUint32((_1345_v2).Cardinality()), (_this).F25())
		var _1347_v4 _dafny.Map
		_ = _1347_v4
		_1347_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1344_v1, _1346_v3)
		var _1348_v5 D15
		_ = _1348_v5
		_1348_v5 = Companion_D15_.Create_DC43_((_1347_v4).Merge(_1347_v4))
		var _source11 D15 = _1348_v5
		_ = _source11
		if _source11.Is_DC44() {
			var _1349___mcc_h0 bool = _source11.Get_().(D15_DC44).Cf68
			_ = _1349___mcc_h0
			var _1350___mcc_h1 _dafny.Sequence = _source11.Get_().(D15_DC44).Cf69
			_ = _1350___mcc_h1
			var _1351___mcc_h2 _dafny.Int = _source11.Get_().(D15_DC44).Cf70
			_ = _1351___mcc_h2
			var _1352___mcc_h3 bool = _source11.Get_().(D15_DC44).Cf71
			_ = _1352___mcc_h3
			var _1353___mcc_h4 bool = _source11.Get_().(D15_DC44).Cf72
			_ = _1353___mcc_h4
			var _1354_cf72 bool = _1353___mcc_h4
			_ = _1354_cf72
			var _1355_cf71 bool = _1352___mcc_h3
			_ = _1355_cf71
			var _1356_cf70 _dafny.Int = _1351___mcc_h2
			_ = _1356_cf70
			var _1357_cf69 _dafny.Sequence = _1350___mcc_h1
			_ = _1357_cf69
			var _1358_cf68 bool = _1349___mcc_h0
			_ = _1358_cf68
			var _1359_v6 _dafny.Array
			_ = _1359_v6
			var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw211
			_1359_v6 = _nw211
			var _1360_v7 _dafny.Map
			_ = _1360_v7
			_1360_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), p0)
			var _1361_v8 D12
			_ = _1361_v8
			_1361_v8 = Companion_D12_.Create_DC34_(_1360_v7)
			var _pat_let_tv41 = _1360_v7
			_ = _pat_let_tv41
			var _pat_let_tv42 = _1360_v7
			_ = _pat_let_tv42
			var _pat_let_tv43 = _1360_v7
			_ = _pat_let_tv43
			var _1362_v9 _dafny.Array
			_ = _1362_v9
			var _nwElement0_44 D12 = Companion_D12_.Create_DC34_(_1360_v7)
			_ = _nwElement0_44
			var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(21))
			_ = _nw212
			(_nw212).ArraySet1(_nwElement0_44, 0)
			(_nw212).ArraySet1(Companion_D12_.Create_DC34_(_1360_v7), 1)
			(_nw212).ArraySet1(_1361_v8, 2)
			(_nw212).ArraySet1(_1361_v8, 3)
			(_nw212).ArraySet1(func(_pat_let24_0 D12) D12 {
				return func(_1363_dt__update__tmp_h0 D12) D12 {
					return func(_pat_let25_0 _dafny.Map) D12 {
						return func(_1364_dt__update_hcf55_h0 _dafny.Map) D12 {
							return Companion_D12_.Create_DC34_(_1364_dt__update_hcf55_h0)
						}(_pat_let25_0)
					}(_pat_let_tv41)
				}(_pat_let24_0)
			}(_1361_v8), 4)
			(_nw212).ArraySet1(_1361_v8, 5)
			(_nw212).ArraySet1(_1361_v8, 6)
			(_nw212).ArraySet1(func(_pat_let26_0 D12) D12 {
				return func(_1367_dt__update__tmp_h2 D12) D12 {
					return func(_pat_let29_0 _dafny.Map) D12 {
						return func(_1368_dt__update_hcf55_h2 _dafny.Map) D12 {
							return Companion_D12_.Create_DC34_(_1368_dt__update_hcf55_h2)
						}(_pat_let29_0)
					}(_pat_let_tv43)
				}(_pat_let26_0)
			}(func(_pat_let27_0 D12) D12 {
				return func(_1365_dt__update__tmp_h1 D12) D12 {
					return func(_pat_let28_0 _dafny.Map) D12 {
						return func(_1366_dt__update_hcf55_h1 _dafny.Map) D12 {
							return Companion_D12_.Create_DC34_(_1366_dt__update_hcf55_h1)
						}(_pat_let28_0)
					}(_pat_let_tv42)
				}(_pat_let27_0)
			}(_1361_v8)), 7)
			(_nw212).ArraySet1(_1361_v8, 8)
			(_nw212).ArraySet1(Companion_D12_.Create_DC34_(_1360_v7), 9)
			(_nw212).ArraySet1(_1361_v8, 10)
			(_nw212).ArraySet1(_1361_v8, 11)
			(_nw212).ArraySet1(_1361_v8, 12)
			(_nw212).ArraySet1(_1361_v8, 13)
			(_nw212).ArraySet1(Companion_D12_.Create_DC34_(_1360_v7), 14)
			(_nw212).ArraySet1(_1361_v8, 15)
			(_nw212).ArraySet1(_1361_v8, 16)
			(_nw212).ArraySet1(_1361_v8, 17)
			(_nw212).ArraySet1(_1361_v8, 18)
			(_nw212).ArraySet1(_1361_v8, 19)
			(_nw212).ArraySet1(_1361_v8, 20)
			_1362_v9 = _nw212
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_1359_v6), 0))
			_ = _index242
			(_1359_v6).ArraySet1(_1362_v9, (_index242).Int())
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_1359_v6), 0))
			_ = _index243
			(_1359_v6).ArraySet1(_1362_v9, (_index243).Int())
			var _1369_v10 _dafny.Map
			_ = _1369_v10
			_1369_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1354_cf72, _1358_cf68)
			var _1370_v11 _dafny.Sequence
			_ = _1370_v11
			_1370_v11 = _dafny.SeqOf(_1345_v2)
			var _1371_v12 _dafny.Sequence
			_ = _1371_v12
			_1371_v12 = _dafny.SeqOf(_1370_v11)
			var _1372_v13 _dafny.Map
			_ = _1372_v13
			_1372_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _1355_cf71)
			var _1373_v14 _dafny.Set
			_ = _1373_v14
			_1373_v14 = _dafny.SetOf(_1354_cf72)
			var _1374_v15 _dafny.Array
			_ = _1374_v15
			var _nwElement0_45 _dafny.Int = (_1369_v10).Cardinality()
			_ = _nwElement0_45
			var _nw213 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(27))
			_ = _nw213
			(_nw213).ArraySet1(_nwElement0_45, 0)
			(_nw213).ArraySet1((_this).F26(), 1)
			(_nw213).ArraySet1(_dafny.IntOfUint32(((_1371_v12).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(!(_1355_cf71))).Cardinality(), _dafny.IntOfUint32((_1371_v12).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), 2)
			(_nw213).ArraySet1((_1372_v13).Cardinality(), 3)
			(_nw213).ArraySet1((_this).F25(), 4)
			(_nw213).ArraySet1((_this).F23(), 5)
			(_nw213).ArraySet1((_this).F26(), 6)
			(_nw213).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_1373_v14).Cardinality()))).Cardinality()), (_this).F26())).Cardinality(), 7)
			(_nw213).ArraySet1(_dafny.IntOfInt64(353), 8)
			(_nw213).ArraySet1((_this).F25(), 9)
			(_nw213).ArraySet1((_this).F25(), 10)
			(_nw213).ArraySet1(_dafny.IntOfInt64(-646), 11)
			(_nw213).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rocyxmwas")).Cardinality()), 12)
			(_nw213).ArraySet1(_1356_cf70, 13)
			(_nw213).ArraySet1((_1372_v13).Cardinality(), 14)
			(_nw213).ArraySet1(_dafny.IntOfInt64(390), 15)
			(_nw213).ArraySet1((_this).F25(), 16)
			(_nw213).ArraySet1((_this).F23(), 17)
			(_nw213).ArraySet1(_1356_cf70, 18)
			(_nw213).ArraySet1((_this).F23(), 19)
			(_nw213).ArraySet1((_this).F23(), 20)
			(_nw213).ArraySet1((_this).F26(), 21)
			(_nw213).ArraySet1((_this).F25(), 22)
			(_nw213).ArraySet1((_this).F23(), 23)
			(_nw213).ArraySet1((_this).F23(), 24)
			(_nw213).ArraySet1((_this).F23(), 25)
			(_nw213).ArraySet1((_this).F23(), 26)
			_1374_v15 = _nw213
			var _1375_v16 _dafny.Map
			_ = _1375_v16
			_1375_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v15, _1355_cf71)
			var _1376_v17 _dafny.Array
			_ = _1376_v17
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_39
			var _nw214 _dafny.Array
			_ = _nw214
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw214 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) bool = (func(_1377_cf71 bool) func(_dafny.Int) bool {
					return func(_1378_i0 _dafny.Int) bool {
						return _1377_cf71
					}
				})(_1355_cf71)
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw214 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw214).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw214).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1376_v17 = _nw214
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_1376_v17), 0))
			_ = _index244
			(_1376_v17).ArraySet1(_1355_cf71, (_index244).Int())
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_1376_v17), 0))
			_ = _index245
			var _rhs252 _dafny.Map = _1375_v16
			_ = _rhs252
			var _rhs253 bool = (func() bool {
				if (_1369_v10).Contains(_1354_cf72) {
					return (_1369_v10).Get(_1354_cf72).(bool)
				}
				return !(((_this).F25()).Cmp((_this).F23()) <= 0)
			})()
			_ = _rhs253
			var _lhs226 _dafny.Array = _1376_v17
			_ = _lhs226
			var _lhs227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_1376_v17), 0))
			_ = _lhs227
			_1375_v16 = _rhs252
			(_lhs226).ArraySet1(_rhs253, (_lhs227).Int())
			var _1379_v18 _dafny.Map
			_ = _1379_v18
			_1379_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1346_v3, _1358_cf68)
			var _rhs254 bool = _1354_cf72
			_ = _rhs254
			var _rhs255 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1346_v3, _1354_cf72)
			_ = _rhs255
			var _rhs256 _dafny.Int = _1356_cf70
			_ = _rhs256
			_1358_cf68 = _rhs254
			_1379_v18 = _rhs255
			r1 = _rhs256
			(globalState).F6 = Companion_Default___.Fm2(globalState)
		} else if _source11.Is_DC43() {
			var _1380___mcc_h5 _dafny.Map = _source11.Get_().(D15_DC43).Cf67
			_ = _1380___mcc_h5
			var _1381_cf67 _dafny.Map = _1380___mcc_h5
			_ = _1381_cf67
			var _1382_v19 _dafny.MultiSet
			_ = _1382_v19
			_1382_v19 = _dafny.MultiSetOf(p0, p0)
			var _1383_v20 _dafny.Sequence
			_ = _1383_v20
			_1383_v20 = _dafny.SeqOf(_1382_v19)
			var _1384_v21 *C3
			_ = _1384_v21
			var _nw215 *C3 = New_C3_()
			_ = _nw215
			_nw215.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1382_v19), _1383_v20), (_dafny.IntOfInt64(983)).Plus((_this).F23()), p0, _1346_v3)
			_1384_v21 = _nw215
			var _1385_v22 _dafny.Set
			_ = _1385_v22
			_1385_v22 = _dafny.SetOf((_1382_v19).Update(p0, Companion_Default___.Abs(_dafny.IntOfInt64(190))), _1382_v19)
			if (_1385_v22).Equals(func() _dafny.Set {
				var _coll51 = _dafny.NewBuilder()
				_ = _coll51
				for _iter55 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1382_v19, p0)).Keys().Elements()); ; {
					_compr_51, _ok55 := _iter55()
					if !_ok55 {
						break
					}
					var _1386_v23 _dafny.MultiSet
					_1386_v23 = interface{}(_compr_51).(_dafny.MultiSet)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1382_v19, p0)).Contains(_1386_v23) {
						_coll51.Add(_1386_v23)
					}
				}
				return _coll51.ToSet()
			}()) {
				var _1387_v24 _dafny.Array
				_ = _1387_v24
				var _nwElement0_46 bool = p0
				_ = _nwElement0_46
				var _nw216 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(9))
				_ = _nw216
				(_nw216).ArraySet1(_nwElement0_46, 0)
				(_nw216).ArraySet1(true, 1)
				(_nw216).ArraySet1(Companion_Default___.Fm1((_this).F26(), p0, globalState), 2)
				(_nw216).ArraySet1(p0, 3)
				(_nw216).ArraySet1(p0, 4)
				(_nw216).ArraySet1(p0, 5)
				(_nw216).ArraySet1(Companion_Default___.Fm1((_this).F25(), p0, globalState), 6)
				(_nw216).ArraySet1(p0, 7)
				(_nw216).ArraySet1(p0, 8)
				_1387_v24 = _nw216
				var _1388_v25 _dafny.Sequence
				_ = _1388_v25
				_1388_v25 = _dafny.SeqOf(_1387_v24, _1387_v24, _1387_v24, _1387_v24)
				_1388_v25 = _1388_v25
				var _1389_v26 _dafny.CodePoint
				_ = _1389_v26
				_1389_v26 = _dafny.CodePoint('j')
				var _1390_v27 _dafny.Map
				_ = _1390_v27
				_1390_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F26()), _1345_v2)
				var _1391_v28 _dafny.Map
				_ = _1391_v28
				_1391_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _1392_v29 _dafny.Sequence
				_ = _1392_v29
				_1392_v29 = _dafny.SeqOf(_1345_v2)
				var _1393_v30 _dafny.Array
				_ = _1393_v30
				var _nwElement0_47 _dafny.Sequence = _1345_v2
				_ = _nwElement0_47
				var _nw217 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(23))
				_ = _nw217
				(_nw217).ArraySet1(_nwElement0_47, 0)
				(_nw217).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pfojm"), _dafny.UnicodeSeqOfUtf8Bytes("mmwijrbtt")), 1)
				(_nw217).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(172))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg90 _dafny.Int) interface{} {
						return coer90(arg90)
					}
				}(func(_1394_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				})), 2)
				(_nw217).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vabaqredp"), 3)
				(_nw217).ArraySet1(_1345_v2, 4)
				(_nw217).ArraySet1(_dafny.Companion_Sequence_.Update(_1345_v2, (Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_1345_v2).Cardinality()))).Uint32(), _1389_v26), 5)
				(_nw217).ArraySet1((func() _dafny.Sequence {
					if (_1390_v27).Contains((_1391_v28).Cardinality()) {
						return (_1390_v27).Get((_1391_v28).Cardinality()).(_dafny.Sequence)
					}
					return _1345_v2
				})(), 6)
				(_nw217).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_1392_v29).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.IntOfUint32((_1392_v29).Cardinality()))).Uint32()).(_dafny.Sequence), _1345_v2), 7)
				(_nw217).ArraySet1(_1345_v2, 8)
				(_nw217).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("untl"), 9)
				(_nw217).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vfldre"), _1345_v2), 10)
				(_nw217).ArraySet1(_1345_v2, 11)
				(_nw217).ArraySet1(_1345_v2, 12)
				(_nw217).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(349))).Uint32(), func(coer91 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg91 _dafny.Int) interface{} {
						return coer91(arg91)
					}
				}((func(_1395_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1396_i2 _dafny.Int) _dafny.CodePoint {
						return _1395_v26
					}
				})(_1389_v26))), 13)
				(_nw217).ArraySet1(_1345_v2, 14)
				(_nw217).ArraySet1(_1345_v2, 15)
				(_nw217).ArraySet1(_1345_v2, 16)
				(_nw217).ArraySet1(_1345_v2, 17)
				(_nw217).ArraySet1(_1345_v2, 18)
				(_nw217).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1345_v2, (Companion_Default___.Fm34(_dafny.IntOfInt64(500), globalState)).Dtor_cf0()), 19)
				(_nw217).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("k"), 20)
				(_nw217).ArraySet1((func() _dafny.Sequence {
					if p0 {
						return _1345_v2
					}
					return _1345_v2
				})(), 21)
				(_nw217).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1345_v2, _dafny.Companion_Sequence_.Update(_1345_v2, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_1345_v2).Cardinality()))).Uint32(), _1389_v26)), 22)
				_1393_v30 = _nw217
				var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1393_v30), 0))
				_ = _index246
				(_1393_v30).ArraySet1(_1345_v2, (_index246).Int())
				var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1393_v30), 0))
				_ = _index247
				(_1393_v30).ArraySet1(_1345_v2, (_index247).Int())
				(globalState).F6 = _dafny.Companion_Sequence_.Concatenate((_1393_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1393_v30), 0))).Int()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("qwammv"))
				var _1397_v31 *C0
				_ = _1397_v31
				var _nw218 *C0 = New_C0_()
				_ = _nw218
				_nw218.Ctor__()
				_1397_v31 = _nw218
				var _1398_v32 _dafny.Sequence
				_ = _1398_v32
				_1398_v32 = _dafny.SeqOf(_1397_v31, _1397_v31, _1397_v31)
				var _1399_v33 D17
				_ = _1399_v33
				_1399_v33 = Companion_D17_.Create_DC48_(_1382_v19)
				var _1400_v34 _dafny.Set
				_ = _1400_v34
				_1400_v34 = _dafny.SetOf(p0)
				_1398_v32 = (func() _dafny.Sequence {
					if (Companion_Default___.Fm35(p0, false, (_1399_v33).Dtor_cf79(), globalState)).IsDisjointFrom(_1400_v34) {
						return _dafny.Companion_Sequence_.Concatenate(_1398_v32, _1398_v32)
					}
					return _1398_v32
				})()
				var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1393_v30), 0))
				_ = _index248
				(_1393_v30).ArraySet1(_1345_v2, (_index248).Int())
			} else {
				r0 = Companion_Default___.Fm1((_this).F23(), p0, globalState)
				var _1401_v35 _dafny.Map
				_ = _1401_v35
				_1401_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F26())
				var _1402_v36 _dafny.Array
				_ = _1402_v36
				var _nw219 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw219
				_1402_v36 = _nw219
				var _1403_v37 _dafny.Sequence
				_ = _1403_v37
				_1403_v37 = _dafny.SeqOf(_1402_v36, _1402_v36, _1402_v36)
				var _1404_v38 *C1
				_ = _1404_v38
				var _nw220 *C1 = New_C1_()
				_ = _nw220
				_nw220.Ctor__(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-756))).Uint32(), func(coer92 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg92 _dafny.Int) interface{} {
						return coer92(arg92)
					}
				}(func(_1405_i3 _dafny.Int) _dafny.Int {
					return (_this).F26()
				})), Companion_Default___.SafeModuloInt((func() _dafny.Int {
					if (_1401_v35).Contains(_dafny.IntOfUint32((_1403_v37).Cardinality())) {
						return (_1401_v35).Get(_dafny.IntOfUint32((_1403_v37).Cardinality())).(_dafny.Int)
					}
					return (_this).F26()
				})(), (_this).F23()))
				_1404_v38 = _nw220
				r2 = _dafny.IntOfInt64(-370)
				(globalState).F9 = p0
				var _1406_v39 _dafny.Map
				_ = _1406_v39
				_1406_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _1407_v40 _dafny.Sequence
				_ = _1407_v40
				_1407_v40 = _dafny.SeqOf(p0, p0, p0)
				var _1408_v41 _dafny.Sequence
				_ = _1408_v41
				_1408_v41 = _dafny.SeqOf(p0, (_1407_v40).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf((_this).F25())).Cardinality(), _dafny.IntOfUint32((_1407_v40).Cardinality()))).Uint32()).(bool), p0, p0)
				_1406_v39 = (_1406_v39).Update(p0, (_1407_v40).Select((Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_1407_v40).Cardinality()))).Uint32()).(bool))
			}
			r2 = Companion_Default___.SafeDivisionInt((_this).F25(), (_this).F25())
			var _1409_v42 _dafny.Array
			_ = _1409_v42
			var _len0_40 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_40
			var _nw221 _dafny.Array
			_ = _nw221
			if _len0_40.Cmp(_dafny.Zero) == 0 {
				_nw221 = _dafny.NewArray(_len0_40)
			} else {
				var _init40 func(_dafny.Int) bool = func(_1410_i4 _dafny.Int) bool {
					return false
				}
				_ = _init40
				var _element0_40 = _init40(_dafny.Zero)
				_ = _element0_40
				_nw221 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
				(_nw221).ArraySet1(_element0_40, 0)
				var _nativeLen0_40 = (_len0_40).Int()
				_ = _nativeLen0_40
				for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
					(_nw221).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
				}
			}
			_1409_v42 = _nw221
			var _1411_v43 D1
			_ = _1411_v43
			_1411_v43 = Companion_D1_.Create_DC4_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25()), p0, p0)
			var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(604), _dafny.ArrayLen((_1409_v42), 0))
			_ = _index249
			(_1409_v42).ArraySet1((_1411_v43).Dtor_cf8(), (_index249).Int())
			var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(604), _dafny.ArrayLen((_1409_v42), 0))
			_ = _index250
			(_1409_v42).ArraySet1(p0, (_index250).Int())
		} else {
			var _1412___mcc_h6 D15 = _source11.Get_().(D15_DC45).Cf73
			_ = _1412___mcc_h6
			var _1413_cf73 D15 = _1412___mcc_h6
			_ = _1413_cf73
			(globalState).F14 = ((_dafny.Zero).Minus((_this).F25())).Minus((_this).F25())
			var _1414_v44 _dafny.Array
			_ = _1414_v44
			var _nwElement0_48 _dafny.Int = (_this).F25()
			_ = _nwElement0_48
			var _nw222 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(2))
			_ = _nw222
			(_nw222).ArraySet1(_nwElement0_48, 0)
			(_nw222).ArraySet1(_dafny.IntOfInt64(79), 1)
			_1414_v44 = _nw222
			var _1415_v45 _dafny.Sequence
			_ = _1415_v45
			_1415_v45 = _dafny.SeqOf(_1414_v44, _1414_v44, _1414_v44)
			_1415_v45 = _1415_v45
			var _1416_v46 _dafny.Map
			_ = _1416_v46
			_1416_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), Companion_Default___.Fm0((_this).F23(), false, globalState))
			var _1417_v47 *C2
			_ = _1417_v47
			var _nw223 *C2 = New_C2_()
			_ = _nw223
			_nw223.Ctor__(_1416_v46, (func() bool {
				if p0 {
					return p0
				}
				return p0
			})(), (func() bool {
				if true {
					return p0
				}
				return p0
			})(), _dafny.Companion_Sequence_.Concatenate(_1346_v3, _1346_v3), (_this).F23())
			_1417_v47 = _nw223
			var _1418_v48 *C0
			_ = _1418_v48
			var _nw224 *C0 = New_C0_()
			_ = _nw224
			_nw224.Ctor__()
			_1418_v48 = _nw224
		}
		(globalState).F7 = (p0) && (p0)
		(globalState).F9 = !(p0) || (p0)
		var _hi15 _dafny.Int = (_this).F23()
		_ = _hi15
		for _1419_i5 := (_this).F23(); _1419_i5.Cmp(_hi15) < 0; _1419_i5 = _1419_i5.Plus(_dafny.One) {
			var _1420_v49 _dafny.Sequence
			_ = _1420_v49
			_1420_v49 = _dafny.SeqOf(p0)
			var _1421_v50 _dafny.Sequence
			_ = _1421_v50
			_1421_v50 = _dafny.SeqOf(_1420_v49, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(Companion_Default___.Fm1((_dafny.MultiSetFromSeq(_1346_v3)).Cardinality(), p0, globalState)), (Companion_Default___.SafeIndex(Companion_Default___.Fm0((_this).F26(), p0, globalState), _dafny.IntOfUint32((_dafny.SeqOf(Companion_Default___.Fm1((_dafny.MultiSetFromSeq(_1346_v3)).Cardinality(), p0, globalState))).Cardinality()))).Uint32(), p0), _dafny.Companion_Sequence_.Update(_1420_v49, (Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_1420_v49).Cardinality()))).Uint32(), p0))
			var _1422_v51 _dafny.Map
			_ = _1422_v51
			_1422_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_this).F23(), p0, globalState), (_this).F26())
			var _1423_v52 D8
			_ = _1423_v52
			_1423_v52 = Companion_D8_.Create_DC19_(_1346_v3)
			var _1424_v53 _dafny.Array
			_ = _1424_v53
			var _nwElement0_49 _dafny.Sequence = _1420_v49
			_ = _nwElement0_49
			var _nw225 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(27))
			_ = _nw225
			(_nw225).ArraySet1(_nwElement0_49, 0)
			(_nw225).ArraySet1((_1421_v50).Select((Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_1421_v50).Cardinality()))).Uint32()).(_dafny.Sequence), 1)
			(_nw225).ArraySet1(_1420_v49, 2)
			(_nw225).ArraySet1(_1420_v49, 3)
			(_nw225).ArraySet1(_1420_v49, 4)
			(_nw225).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1420_v49, _1420_v49), 5)
			(_nw225).ArraySet1(_1420_v49, 6)
			(_nw225).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1420_v49, _1420_v49), 7)
			(_nw225).ArraySet1(_1420_v49, 8)
			(_nw225).ArraySet1((func() _dafny.Sequence {
				if p0 {
					return _dafny.SeqOf(true)
				}
				return _1420_v49
			})(), 9)
			(_nw225).ArraySet1(_1420_v49, 10)
			(_nw225).ArraySet1(Companion_Default___.Fm30((func() _dafny.Int {
				if (_1422_v51).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1346_v3, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F23()), _dafny.IntOfUint32((_1346_v3).Cardinality()))).Uint32(), _dafny.IntOfInt64(612))).Cardinality())) {
					return (_1422_v51).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1346_v3, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F23()), _dafny.IntOfUint32((_1346_v3).Cardinality()))).Uint32(), _dafny.IntOfInt64(612))).Cardinality())).(_dafny.Int)
				}
				return (_this).F26()
			})(), _1423_v52, _1345_v2, _1419_i5, globalState), 11)
			(_nw225).ArraySet1(_1420_v49, 12)
			(_nw225).ArraySet1(_1420_v49, 13)
			(_nw225).ArraySet1(Companion_Default___.Fm30((_this).F23(), _1423_v52, _1345_v2, (_this).F25(), globalState), 14)
			(_nw225).ArraySet1(_1420_v49, 15)
			(_nw225).ArraySet1((func() _dafny.Sequence {
				if true {
					return _1420_v49
				}
				return _1420_v49
			})(), 16)
			(_nw225).ArraySet1(_1420_v49, 17)
			(_nw225).ArraySet1(_dafny.Companion_Sequence_.Update(_1420_v49, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.IntOfUint32((_1420_v49).Cardinality()))).Uint32(), p0), 18)
			(_nw225).ArraySet1(_1420_v49, 19)
			(_nw225).ArraySet1(_dafny.SeqOf(p0, p0), 20)
			(_nw225).ArraySet1(_1420_v49, 21)
			(_nw225).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1420_v49, _1420_v49), 22)
			(_nw225).ArraySet1(_1420_v49, 23)
			(_nw225).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm30((_this).F23(), _1423_v52, _1345_v2, _1419_i5, globalState), _dafny.SeqOf(p0, true, !(p0))), 24)
			(_nw225).ArraySet1(_dafny.Companion_Sequence_.Update(_1420_v49, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_1345_v2).Cardinality())), _dafny.IntOfUint32((_1420_v49).Cardinality()))).Uint32(), true), 25)
			(_nw225).ArraySet1(_1420_v49, 26)
			_1424_v53 = _nw225
			_1424_v53 = _1424_v53
			var _rhs257 _dafny.Int = (_this).F26()
			_ = _rhs257
			var _rhs258 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1345_v2, _1345_v2), _1345_v2)
			_ = _rhs258
			var _rhs259 bool = p0
			_ = _rhs259
			var _lhs228 *GlobalState = globalState
			_ = _lhs228
			var _lhs229 *GlobalState = globalState
			_ = _lhs229
			var _lhs230 *GlobalState = globalState
			_ = _lhs230
			_lhs228.F8 = _rhs257
			_lhs229.F6 = _rhs258
			_lhs230.F9 = _rhs259
			var _1425_v54 _dafny.Array
			_ = _1425_v54
			var _nw226 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw226
			_1425_v54 = _nw226
			var _1426_v55 D8
			_ = _1426_v55
			_1426_v55 = Companion_D8_.Create_DC20_(true, _1419_i5)
			var _1427_v56 D8
			_ = _1427_v56
			_1427_v56 = Companion_D8_.Create_DC23_(_1426_v55)
			var _1428_v57 T1
			_ = _1428_v57
			var _nw227 *C2 = New_C2_()
			_ = _nw227
			_nw227.Ctor__(_1422_v51, p0, !(p0), _1346_v3, (_this).F25())
			_1428_v57 = _nw227
			var _1429_v58 _dafny.Sequence
			_ = _1429_v58
			_1429_v58 = _dafny.SeqOf(_1428_v57)
			var _1430_v59 _dafny.Sequence
			_ = _1430_v59
			_1430_v59 = _dafny.SeqOf(_1425_v54, _1425_v54)
			var _1431_v60 _dafny.Set
			_ = _1431_v60
			_1431_v60 = _dafny.SetOf((_this).F26(), (_this).F26())
			var _1432_v61 _dafny.Map
			_ = _1432_v61
			_1432_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1428_v57).F28(), _1428_v57)
			var _1433_v62 T1
			_ = _1433_v62
			_1433_v62 = _1428_v57
			var _1434_v63 _dafny.Array
			_ = _1434_v63
			var _nwElement0_50 T1 = _1428_v57
			_ = _nwElement0_50
			var _nw228 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(28))
			_ = _nw228
			(_nw228).ArraySet1(_nwElement0_50, 0)
			(_nw228).ArraySet1((_1429_v58).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1430_v59).Cardinality()), _dafny.IntOfUint32((_1429_v58).Cardinality()))).Uint32()).(T1), 1)
			(_nw228).ArraySet1(_1428_v57, 2)
			(_nw228).ArraySet1(_1428_v57, 3)
			(_nw228).ArraySet1(_1428_v57, 4)
			(_nw228).ArraySet1(_1428_v57, 5)
			(_nw228).ArraySet1(_1428_v57, 6)
			(_nw228).ArraySet1(_1428_v57, 7)
			(_nw228).ArraySet1(_1428_v57, 8)
			(_nw228).ArraySet1((_1429_v58).Select((Companion_Default___.SafeIndex(_1419_i5, _dafny.IntOfUint32((_1429_v58).Cardinality()))).Uint32()).(T1), 9)
			(_nw228).ArraySet1(_1428_v57, 10)
			(_nw228).ArraySet1((_1429_v58).Select((Companion_Default___.SafeIndex((_1431_v60).Cardinality(), _dafny.IntOfUint32((_1429_v58).Cardinality()))).Uint32()).(T1), 11)
			(_nw228).ArraySet1((_1429_v58).Select((Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_1429_v58).Cardinality()))).Uint32()).(T1), 12)
			(_nw228).ArraySet1((func() T1 {
				if (_1432_v61).Contains(!((_1428_v57).F28())) {
					return (_1432_v61).Get(!((_1428_v57).F28())).(T1)
				}
				return _1428_v57
			})(), 13)
			(_nw228).ArraySet1(_1428_v57, 14)
			(_nw228).ArraySet1((_1433_v62), 15)
			(_nw228).ArraySet1(_1428_v57, 16)
			(_nw228).ArraySet1(_1428_v57, 17)
			(_nw228).ArraySet1(_1428_v57, 18)
			(_nw228).ArraySet1(_1428_v57, 19)
			(_nw228).ArraySet1(_1428_v57, 20)
			(_nw228).ArraySet1((_1433_v62), 21)
			(_nw228).ArraySet1(_1428_v57, 22)
			(_nw228).ArraySet1(_1428_v57, 23)
			(_nw228).ArraySet1(_1428_v57, 24)
			(_nw228).ArraySet1(_1428_v57, 25)
			(_nw228).ArraySet1(_1428_v57, 26)
			(_nw228).ArraySet1(_1428_v57, 27)
			_1434_v63 = _nw228
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_1434_v63), 0))
			_ = _index251
			(_1434_v63).ArraySet1(_1428_v57, (_index251).Int())
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_1434_v63), 0))
			_ = _index252
			var _rhs260 _dafny.Int = ((_this).F25()).Times((_1428_v57).F23())
			_ = _rhs260
			var _rhs261 _dafny.Array = _1425_v54
			_ = _rhs261
			var _rhs262 bool = p0
			_ = _rhs262
			var _rhs263 D8 = _1427_v56
			_ = _rhs263
			var _rhs264 T1 = _1428_v57
			_ = _rhs264
			var _lhs231 *GlobalState = globalState
			_ = _lhs231
			var _lhs232 *GlobalState = globalState
			_ = _lhs232
			var _lhs233 _dafny.Array = _1434_v63
			_ = _lhs233
			var _lhs234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_1434_v63), 0))
			_ = _lhs234
			_lhs231.F14 = _rhs260
			_1425_v54 = _rhs261
			_lhs232.F9 = _rhs262
			_1427_v56 = _rhs263
			(_lhs233).ArraySet1(_rhs264, (_lhs234).Int())
			(globalState).F7 = ((_this).F26()).Cmp((func() _dafny.Int {
				if (_1428_v57).F28() {
					return _dafny.IntOfUint32((_1420_v49).Cardinality())
				}
				return (_this).F25()
			})()) == 0
		}
		var _1435_i6 _dafny.Int
		_ = _1435_i6
		_1435_i6 = _dafny.Zero
		{
			for ((_this).F26()).Cmp((_this).F23()) > 0 {
				{
					if (_1435_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L13
					}
					_1435_i6 = (_1435_i6).Plus(_dafny.One)
					var _1436_v64 _dafny.Map
					_ = _1436_v64
					_1436_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F26())
					var _1437_v65 _dafny.Map
					_ = _1437_v65
					_1437_v65 = (_1436_v64).Merge(_1436_v64)
					var _source12 _dafny.Map = _1437_v65
					_ = _source12
					var _1438___mcc_h7 _dafny.Map = _source12
					_ = _1438___mcc_h7
					var _1439_cf31 _dafny.Map = _1438___mcc_h7
					_ = _1439_cf31
					var _1440_v66 _dafny.Sequence
					_ = _1440_v66
					_1440_v66 = _dafny.SeqOf(true)
					var _1441_v67 _dafny.Array
					_ = _1441_v67
					var _nw229 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
					_ = _nw229
					_1441_v67 = _nw229
					var _1442_v68 _dafny.Map
					_ = _1442_v68
					_1442_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1440_v66, _1441_v67)
					var _1443_v69 _dafny.Sequence
					_ = _1443_v69
					_1443_v69 = _dafny.SeqOf(_1442_v68)
					var _1444_v70 _dafny.Map
					_ = _1444_v70
					_1444_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1443_v69).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_1443_v69).Cardinality()))).Uint32()).(_dafny.Map), (_dafny.Zero).Minus((_dafny.IntOfInt64(-861)).Minus((_this).F26())))
					_1444_v70 = (_1444_v70).Update(_1442_v68, (_this).F23())
					(globalState).F21 = true
					var _1445_v71 _dafny.Array
					_ = _1445_v71
					var _nw230 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
					_ = _nw230
					_1445_v71 = _nw230
					var _1446_v72 _dafny.Map
					_ = _1446_v72
					_1446_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(300), !(p0))
					var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_1445_v71), 0))
					_ = _index253
					(_1445_v71).ArraySet1((_1446_v72).Update((_this).F23(), p0), (_index253).Int())
					var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_1445_v71), 0))
					_ = _index254
					(_1445_v71).ArraySet1(_1446_v72, (_index254).Int())
					var _1447_v73 _dafny.Map
					_ = _1447_v73
					_1447_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((_this).F25()).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(705))).Uint32(), func(coer93 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg93 _dafny.Int) interface{} {
							return coer93(arg93)
						}
					}(func(_1448_i7 _dafny.Int) _dafny.Int {
						return (_this).F25()
					}))).Cardinality()))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), !(p0)))
					_1447_v73 = (_1447_v73).Update(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-201), _dafny.IntOfInt64(598))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), p0))
					var _1449_v74 _dafny.Sequence
					_ = _1449_v74
					_1449_v74 = _dafny.SeqOf(_1348_v5, _1348_v5, _1348_v5)
					_1449_v74 = _1449_v74
					var _1450_v75 *C0
					_ = _1450_v75
					var _nw231 *C0 = New_C0_()
					_ = _nw231
					_nw231.Ctor__()
					_1450_v75 = _nw231
					var _1451_v76 _dafny.Map
					_ = _1451_v76
					_1451_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), _1450_v75)
					_1451_v76 = (_1451_v76).Update(Companion_Default___.SafeModuloInt((_this).F23(), (_this).F25()), _1450_v75)
					(globalState).F8 = ((_1346_v3).Select((Companion_Default___.SafeIndex((_this).F26(), _dafny.IntOfUint32((_1346_v3).Cardinality()))).Uint32()).(_dafny.Int)).Plus((func() _dafny.Int {
						if p0 {
							return (_this).F23()
						}
						return ((_1436_v64).Update(p0, (_this).F26())).Cardinality()
					})())
					goto C13
				}
			C13:
			}
			goto L13
		}
	L13:
		r0 = p0
		r1 = (_this).F25()
		var _1452_v77 D0
		_ = _1452_v77
		_1452_v77 = Companion_D0_.Create_DC0_(_1345_v2)
		var _1453_v78 D0
		_ = _1453_v78
		_1453_v78 = Companion_D0_.Create_DC2_(_1452_v77)
		var _pat_let_tv44 = _1346_v3
		_ = _pat_let_tv44
		var _pat_let_tv45 = _1346_v3
		_ = _pat_let_tv45
		var _pat_let_tv46 = _1346_v3
		_ = _pat_let_tv46
		var _pat_let_tv47 = _1346_v3
		_ = _pat_let_tv47
		r2 = func(_source13 D0) _dafny.Int {
			if _source13.Is_DC1() {
				var _1454___mcc_h8 _dafny.Int = _source13.Get_().(D0_DC1).Cf1
				_ = _1454___mcc_h8
				var _1455___mcc_h9 bool = _source13.Get_().(D0_DC1).Cf2
				_ = _1455___mcc_h9
				var _1456___mcc_h10 _dafny.Int = _source13.Get_().(D0_DC1).Cf3
				_ = _1456___mcc_h10
				var _1457_cf3 _dafny.Int = _1456___mcc_h10
				_ = _1457_cf3
				var _1458_cf2 bool = _1455___mcc_h9
				_ = _1458_cf2
				var _1459_cf1 _dafny.Int = _1454___mcc_h8
				_ = _1459_cf1
				return ((_this).F26()).Plus((_this).F23())
			} else if _source13.Is_DC0() {
				var _1460___mcc_h11 _dafny.Sequence = _source13.Get_().(D0_DC0).Cf0
				_ = _1460___mcc_h11
				var _1461_cf0 _dafny.Sequence = _1460___mcc_h11
				_ = _1461_cf0
				return Companion_Default___.SafeModuloInt((func() _dafny.Map {
					var _coll52 = _dafny.NewMapBuilder()
					_ = _coll52
					for _iter56 := _dafny.Iterate((func() _dafny.Set {
						var _coll53 = _dafny.NewBuilder()
						_ = _coll53
						for _iter57 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-298))).Uint32(), func(coer94 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg94 _dafny.Int) interface{} {
								return coer94(arg94)
							}
						}((func(_1462_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_1463_i8 _dafny.Int) _dafny.Sequence {
								return _1462_v3
							}
						})(_pat_let_tv44)))).Elements()); ; {
							_compr_53, _ok57 := _iter57()
							if !_ok57 {
								break
							}
							var _1464_v80 _dafny.Sequence
							_1464_v80 = interface{}(_compr_53).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-298))).Uint32(), func(coer95 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg95 _dafny.Int) interface{} {
									return coer95(arg95)
								}
							}((func(_1465_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_1463_i8 _dafny.Int) _dafny.Sequence {
									return _1465_v3
								}
							})(_pat_let_tv45))), _1464_v80) {
								_coll53.Add(_1464_v80)
							}
						}
						return _coll53.ToSet()
					}()).Elements()); ; {
						_compr_52, _ok56 := _iter56()
						if !_ok56 {
							break
						}
						var _1466_v79 _dafny.Sequence
						_1466_v79 = interface{}(_compr_52).(_dafny.Sequence)
						if (func() _dafny.Set {
							var _coll54 = _dafny.NewBuilder()
							_ = _coll54
							for _iter58 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-298))).Uint32(), func(coer96 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg96 _dafny.Int) interface{} {
									return coer96(arg96)
								}
							}((func(_1467_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_1463_i8 _dafny.Int) _dafny.Sequence {
									return _1467_v3
								}
							})(_pat_let_tv46)))).Elements()); ; {
								_compr_54, _ok58 := _iter58()
								if !_ok58 {
									break
								}
								var _1468_v80 _dafny.Sequence
								_1468_v80 = interface{}(_compr_54).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-298))).Uint32(), func(coer97 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
									return func(arg97 _dafny.Int) interface{} {
										return coer97(arg97)
									}
								}((func(_1469_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
									return func(_1463_i8 _dafny.Int) _dafny.Sequence {
										return _1469_v3
									}
								})(_pat_let_tv47))), _1468_v80) {
									_coll54.Add(_1468_v80)
								}
							}
							return _coll54.ToSet()
						}()).Contains(_1466_v79) {
							_coll52.Add(_1466_v79, (_this).F25())
						}
					}
					return _coll52.ToMap()
				}()).Cardinality(), _dafny.IntOfInt64(362))
			} else {
				var _1470___mcc_h12 D0 = _source13.Get_().(D0_DC2).Cf4
				_ = _1470___mcc_h12
				var _1471_cf4 D0 = _1470___mcc_h12
				_ = _1471_cf4
				return (_this).F26()
			}
		}(_1453_v78)
		return r0, r1, r2
	}
}
func (_this *C5) F25() _dafny.Int {
	{
		return _this._f25
	}
}
func (_this *C5) F26() _dafny.Int {
	{
		return _this._f26
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f23 _dafny.Int
	_f24 _dafny.Int
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f23 = _dafny.Zero
	_this._f24 = _dafny.Zero
	return &_this
}

type CompanionStruct_C6_ struct {
}

var Companion_C6_ = CompanionStruct_C6_{}

func (_this *C6) Equals(other *C6) bool {
	return _this == other
}

func (_this *C6) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C6)
	return ok && _this.Equals(other)
}

func (*C6) String() string {
	return "_module.C6"
}

func Type_C6_() _dafny.TypeDescriptor {
	return type_C6_{}
}

type type_C6_ struct {
}

func (_this type_C6_) Default() interface{} {
	return (*C6)(nil)
}

func (_this type_C6_) String() string {
	return "main.C6"
}
func (_this *C6) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F23() _dafny.Int {
	return _this._f23
}
func (_this *C6) Ctor__(f24 _dafny.Int, f23 _dafny.Int) {
	{
		(_this)._f24 = f24
		(_this)._f23 = f23
	}
}
func (_this *C6) M1(p0 _dafny.Array, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Set) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Set = _dafny.EmptySet
		_ = r1
		var _1472_v0 _dafny.Map
		_ = _1472_v0
		_1472_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_this).F24(), (_this).F23())).Cardinality()), p0)
		var _hi16 _dafny.Int = (_1472_v0).Cardinality()
		_ = _hi16
		for _1473_i0 := (_this).F23(); _1473_i0.Cmp(_hi16) < 0; _1473_i0 = _1473_i0.Plus(_dafny.One) {
			r0 = (_this).F24()
			var _1474_v1 _dafny.Map
			_ = _1474_v1
			_1474_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !(p1))
			var _1475_v2 _dafny.Set
			_ = _1475_v2
			_1475_v2 = _dafny.SetOf((func() _dafny.Map {
				if p1 {
					return _1474_v1
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), p1)
			})(), _1474_v1, _1474_v1, _1474_v1)
			var _1476_v3 _dafny.CodePoint
			_ = _1476_v3
			_1476_v3 = _dafny.CodePoint('j')
			var _1477_v4 _dafny.MultiSet
			_ = _1477_v4
			_1477_v4 = _dafny.MultiSetOf(_dafny.CodePoint('n'), _1476_v3)
			_1475_v2 = _dafny.SetOf((_1474_v1).Update(Companion_Default___.Fm1((_1477_v4).Cardinality(), p1, globalState), p1), (Companion_Default___.Fm7((_this).F24(), p1, p1, globalState)).Merge(_1474_v1), _1474_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Merge(_1474_v1))
			var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p0), 0))
			_ = _index255
			(p0).ArraySet1(p1, (_index255).Int())
			var _1478_v5 _dafny.MultiSet
			_ = _1478_v5
			_1478_v5 = _dafny.MultiSetOf(p1, p1)
			var _1479_v6 _dafny.Sequence
			_ = _1479_v6
			_1479_v6 = _dafny.SeqOf(p1)
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((p0), 0))
			_ = _index256
			(p0).ArraySet1((_1478_v5).IsDisjointFrom(Companion_Default___.Fm4(Companion_D0_.Create_DC0_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(591))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg98 _dafny.Int) interface{} {
					return coer98(arg98)
				}
			}((func(_1480_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1481_i1 _dafny.Int) _dafny.CodePoint {
					return _1480_v3
				}
			})(_1476_v3)))), _1473_i0, _1479_v6, _1476_v3, globalState)), (_index256).Int())
			var _1482_v7 _dafny.Sequence
			_ = _1482_v7
			_1482_v7 = _dafny.UnicodeSeqOfUtf8Bytes("cduaj")
			var _1483_v8 _dafny.Sequence
			_ = _1483_v8
			_1483_v8 = _dafny.SeqOf(_1482_v7)
			(globalState).F14 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_1483_v8).Cardinality()))
		}
		var _1484_v9 _dafny.Sequence
		_ = _1484_v9
		_1484_v9 = _dafny.SeqOf((_this).F23())
		(globalState).F14 = (_1484_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.IntOfUint32((_1484_v9).Cardinality()))).Uint32()).(_dafny.Int)
		var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((p0), 0))
		_ = _index257
		(p0).ArraySet1((Companion_Default___.Fm0((_this).F23(), p1, globalState)).Cmp((_this).F24()) != 0, (_index257).Int())
		var _1485_v10 _dafny.Map
		_ = _1485_v10
		_1485_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F24())
		var _1486_v11 _dafny.Sequence
		_ = _1486_v11
		_1486_v11 = _dafny.UnicodeSeqOfUtf8Bytes("l")
		var _1487_v12 _dafny.Set
		_ = _1487_v12
		_1487_v12 = _dafny.SetOf((_this).F24(), (func() _dafny.Int {
			if (_1485_v10).Contains((_this).F24()) {
				return (_1485_v10).Get((_this).F24()).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_1486_v11).Cardinality())
		})())
		var _1488_v13 _dafny.MultiSet
		_ = _1488_v13
		_1488_v13 = _dafny.MultiSetOf(Companion_Default___.Fm1((_1487_v12).Cardinality(), false, globalState), p1)
		var _1489_v14 _dafny.Set
		_ = _1489_v14
		_1489_v14 = _dafny.SetOf(p1, p1)
		var _1490_v15 T0
		_ = _1490_v15
		var _nw232 *C3 = New_C3_()
		_ = _nw232
		_nw232.Ctor__(_dafny.SeqOf((_1488_v13).Update(true, Companion_Default___.Abs((_1489_v14).Cardinality())), _dafny.MultiSetOf(p1, p1), _1488_v13, _1488_v13, _dafny.MultiSetOf(p1)), (_this).F23(), p1, _1484_v9)
		_1490_v15 = _nw232
		var _1491_v16 _dafny.Sequence
		_ = _1491_v16
		_1491_v16 = _dafny.SeqOf(_1490_v15)
		var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((p0), 0))
		_ = _index258
		(p0).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1491_v16, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Cardinality(), _dafny.IntOfUint32((_1491_v16).Cardinality()))).Uint32(), _1490_v15), _1491_v16), _1490_v15), (_index258).Int())
		var _1492_v17 _dafny.Map
		_ = _1492_v17
		_1492_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(904))
		var _1493_v18 _dafny.Sequence
		_ = _1493_v18
		_1493_v18 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((p0), 0))).Int()).(bool), (_1490_v15).F23()), _1492_v17)
		r0 = (func() _dafny.Int {
			if _dafny.Companion_Sequence_.Contains(_1493_v18, _1492_v17) {
				return ((_this).F23()).Times((_this).F24())
			}
			return _dafny.IntOfInt64(-954)
		})()
		var _hi17 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F24(), (_this).F23())
		_ = _hi17
		for _1494_i2 := (_1490_v15).F23(); _1494_i2.Cmp(_hi17) < 0; _1494_i2 = _1494_i2.Plus(_dafny.One) {
			var _1495_v19 _dafny.Map
			_ = _1495_v19
			_1495_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(!(p1)))
			_1495_v19 = (_1495_v19).Update(p1, Companion_Default___.Fm1((_1490_v15).F23(), p1, globalState))
			var _1496_v20 _dafny.Sequence
			_ = _1496_v20
			_1496_v20 = _dafny.SeqOf(Companion_Default___.Fm1((_this).F23(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState), p1, (_1488_v13).IsDisjointFrom(_1488_v13), p1, (func() bool {
				if p1 {
					return !(p1)
				}
				return true
			})())
			(globalState).F7 = (_1496_v20).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm2(globalState)).Cardinality()), _dafny.IntOfUint32((_1496_v20).Cardinality()))).Uint32()).(bool)
			var _1497_v21 _dafny.CodePoint
			_ = _1497_v21
			_1497_v21 = _dafny.CodePoint('n')
			var _1498_v22 _dafny.Map
			_ = _1498_v22
			_1498_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1497_v21, _dafny.Companion_Sequence_.Concatenate(_1484_v9, _dafny.SeqOf((_1490_v15).F23())))
			_1498_v22 = (_1498_v22).Update((_1486_v11).Select((Companion_Default___.SafeIndex((_1490_v15).F23(), _dafny.IntOfUint32((_1486_v11).Cardinality()))).Uint32()).(_dafny.CodePoint), _1484_v9)
			_1496_v20 = _dafny.Companion_Sequence_.Concatenate(_1496_v20, _1496_v20)
		}
		var _1499_v23 _dafny.Array
		_ = _1499_v23
		var _nw233 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw233
		_1499_v23 = _nw233
		for _iter59 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1499_v23), 0))); ; {
			_guard_loop_4, _ok59 := _iter59()
			if !_ok59 {
				break
			}
			var _1500_i3 _dafny.Int
			_1500_i3 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1500_i3).Sign() != -1) && ((_1500_i3).Cmp(_dafny.ArrayLen((_1499_v23), 0)) < 0)) {
				(_1499_v23).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1486_v11, _1486_v11), (_1500_i3).Int())
			}
		}
		r0 = (_this).F24()
		var _1501_v24 _dafny.MultiSet
		_ = _1501_v24
		_1501_v24 = _dafny.MultiSetOf((_this).F23())
		var _1502_v25 _dafny.Map
		_ = _1502_v25
		_1502_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1501_v24, (_1488_v13).Cardinality())
		var _1503_v26 D18
		_ = _1503_v26
		_1503_v26 = Companion_D18_.Create_DC50_(_1502_v25)
		r1 = _dafny.SetOf((_1490_v15).F23(), (((_1503_v26).Dtor_cf84()).Merge(_1502_v25)).Cardinality(), (_this).F23())
		return r0, r1
	}
}
func (_this *C6) M2(p0 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var _1504_v0 bool
		_ = _1504_v0
		_1504_v0 = true
		var _1505_v1 _dafny.Map
		_ = _1505_v1
		_1505_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1504_v0, (_this).F23())
		var _1506_v2 _dafny.CodePoint
		_ = _1506_v2
		_1506_v2 = _dafny.CodePoint('u')
		var _1507_v3 _dafny.Set
		_ = _1507_v3
		_1507_v3 = _dafny.SetOf(_1506_v2)
		var _1508_v4 _dafny.Map
		_ = _1508_v4
		_1508_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1507_v3, p0)
		var _1509_v5 _dafny.Set
		_ = _1509_v5
		_1509_v5 = _dafny.SetOf(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1508_v4).Contains(_1507_v3) {
				return (_1508_v4).Get(_1507_v3).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _1506_v2), (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _1506_v2)).Cardinality()))).Uint32(), _1506_v2)
		})()).Cardinality()), (_this).F23())
		_1505_v1 = (_1505_v1).Update((_1509_v5).IsProperSubsetOf(_dafny.SetOf((_this).F23())), (_this).F24())
		_1506_v2 = _1506_v2
		if _1504_v0 {
			var _1510_v6 D3
			_ = _1510_v6
			_1510_v6 = Companion_D3_.Create_DC10_(!(_1504_v0))
			var _source14 D3 = _1510_v6
			_ = _source14
			if _source14.Is_DC10() {
				var _1511___mcc_h0 bool = _source14.Get_().(D3_DC10).Cf21
				_ = _1511___mcc_h0
				var _1512_cf21 bool = _1511___mcc_h0
				_ = _1512_cf21
				var _1513_v7 _dafny.Array
				_ = _1513_v7
				var _nw234 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
				_ = _nw234
				_1513_v7 = _nw234
				var _1514_v8 _dafny.Array
				_ = _1514_v8
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_41
				var _nw235 _dafny.Array
				_ = _nw235
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw235 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) bool = (func(_1515_cf21 bool) func(_dafny.Int) bool {
						return func(_1516_i0 _dafny.Int) bool {
							return _1515_cf21
						}
					})(_1512_cf21)
					_ = _init41
					var _element0_41 = _init41(_dafny.Zero)
					_ = _element0_41
					_nw235 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
					(_nw235).ArraySet1(_element0_41, 0)
					var _nativeLen0_41 = (_len0_41).Int()
					_ = _nativeLen0_41
					for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
						(_nw235).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
					}
				}
				_1514_v8 = _nw235
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_1513_v7), 0))
				_ = _index259
				(_1513_v7).ArraySet1(_1514_v8, (_index259).Int())
				var _1517_v9 _dafny.Sequence
				_ = _1517_v9
				_1517_v9 = _dafny.SeqOf((_this).F24())
				var _1518_v10 *C3
				_ = _1518_v10
				var _nw236 *C3 = New_C3_()
				_ = _nw236
				_nw236.Ctor__(Companion_Default___.Fm29(_dafny.IntOfUint32((_1517_v9).Cardinality()), _dafny.SetOf((_this).F23(), (_this).F23()), globalState), (_this).F24(), false, _dafny.Companion_Sequence_.Concatenate(_1517_v9, _1517_v9))
				_1518_v10 = _nw236
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_1513_v7), 0))
				_ = _index260
				var _rhs265 _dafny.Int = (_dafny.Zero).Minus((_this).F23())
				_ = _rhs265
				var _rhs266 _dafny.CodePoint = _1506_v2
				_ = _rhs266
				var _rhs267 _dafny.Array = _1514_v8
				_ = _rhs267
				var _rhs268 *C3 = _1518_v10
				_ = _rhs268
				var _rhs269 _dafny.Array = _1514_v8
				_ = _rhs269
				var _lhs235 *GlobalState = globalState
				_ = _lhs235
				var _lhs236 _dafny.Array = _1513_v7
				_ = _lhs236
				var _lhs237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_1513_v7), 0))
				_ = _lhs237
				_lhs235.F14 = _rhs265
				_1506_v2 = _rhs266
				(_lhs236).ArraySet1(_rhs267, (_lhs237).Int())
				_1518_v10 = _rhs268
				_1514_v8 = _rhs269
				(globalState).F6 = Companion_Default___.Fm2(globalState)
				(globalState).F21 = _1504_v0
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_1514_v8), 0))
				_ = _index261
				(_1514_v8).ArraySet1(_1512_cf21, (_index261).Int())
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_1514_v8), 0))
				_ = _index262
				(_1514_v8).ArraySet1(_1512_cf21, (_index262).Int())
			} else if _source14.Is_DC11() {
				var _1519___mcc_h1 _dafny.Int = _source14.Get_().(D3_DC11).Cf22
				_ = _1519___mcc_h1
				var _1520_cf22 _dafny.Int = _1519___mcc_h1
				_ = _1520_cf22
				var _1521_v11 _dafny.Map
				_ = _1521_v11
				_1521_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1509_v5, _dafny.Companion_Sequence_.IsPrefixOf(p0, _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _1506_v2)))
				_1521_v11 = (_1521_v11).Update(_1509_v5, _1504_v0)
				var _1522_v12 _dafny.MultiSet
				_ = _1522_v12
				_1522_v12 = _dafny.MultiSetOf(_1504_v0)
				var _1523_v13 _dafny.Array
				_ = _1523_v13
				var _nwElement0_51 _dafny.Int = (_this).F24()
				_ = _nwElement0_51
				var _nw237 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(8))
				_ = _nw237
				(_nw237).ArraySet1(_nwElement0_51, 0)
				(_nw237).ArraySet1(_1520_cf22, 1)
				(_nw237).ArraySet1(_1520_cf22, 2)
				(_nw237).ArraySet1((_1505_v1).Cardinality(), 3)
				(_nw237).ArraySet1((_this).F24(), 4)
				(_nw237).ArraySet1((_this).F24(), 5)
				(_nw237).ArraySet1((_1522_v12).Cardinality(), 6)
				(_nw237).ArraySet1((_this).F24(), 7)
				_1523_v13 = _nw237
				var _1524_v14 _dafny.Map
				_ = _1524_v14
				_1524_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1523_v13, _1504_v0)
				var _1525_v15 _dafny.MultiSet
				_ = _1525_v15
				_1525_v15 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1504_v0, (func() bool {
					if (_1524_v14).Contains(_1523_v13) {
						return (_1524_v14).Get(_1523_v13).(bool)
					}
					return _1504_v0
				})())).Cardinality(), _dafny.IntOfUint32((p0).Cardinality()), (_this).F23(), _1520_cf22, _1520_cf22)
				var _1526_v16 _dafny.Array
				_ = _1526_v16
				var _nw238 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw238
				_1526_v16 = _nw238
				var _1527_v17 _dafny.Map
				_ = _1527_v17
				_1527_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1525_v15, _1526_v16)
				var _1528_v18 _dafny.Map
				_ = _1528_v18
				_1528_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F24()), _1527_v17)
				_1528_v18 = (_1528_v18).Update((_this).F24(), _1527_v17)
				var _1529_v19 D2
				_ = _1529_v19
				_1529_v19 = Companion_D2_.Create_DC7_(_1507_v3)
				var _1530_v20 _dafny.Map
				_ = _1530_v20
				_1530_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1529_v19, _1504_v0)
				var _1531_v21 _dafny.Map
				_ = _1531_v21
				_1531_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1530_v20).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1529_v19, _1504_v0)), (_this).F24())
				_1531_v21 = (_1531_v21).Update(_1530_v20, (_this).F24())
				var _1532_v22 *C5
				_ = _1532_v22
				var _nw239 *C5 = New_C5_()
				_ = _nw239
				_nw239.Ctor__(_dafny.IntOfUint32((p0).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality())), (_this).F24())
				_1532_v22 = _nw239
				_1532_v22 = _1532_v22
			} else if _source14.Is_DC12() {
				var _1533___mcc_h2 _dafny.Sequence = _source14.Get_().(D3_DC12).Cf23
				_ = _1533___mcc_h2
				var _1534_cf23 _dafny.Sequence = _1533___mcc_h2
				_ = _1534_cf23
				var _1535_v23 _dafny.Map
				_ = _1535_v23
				_1535_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), (_this).F24())
				var _1536_v24 _dafny.Array
				_ = _1536_v24
				var _nwElement0_52 _dafny.Int = (_this).F23()
				_ = _nwElement0_52
				var _nw240 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(12))
				_ = _nw240
				(_nw240).ArraySet1(_nwElement0_52, 0)
				(_nw240).ArraySet1(Companion_Default___.Fm0((_this).F24(), _1504_v0, globalState), 1)
				(_nw240).ArraySet1((_this).F24(), 2)
				(_nw240).ArraySet1((func() _dafny.Int {
					if false {
						return (_this).F24()
					}
					return (_this).F24()
				})(), 3)
				(_nw240).ArraySet1((func() _dafny.Int {
					if (_1535_v23).Contains((_this).F24()) {
						return (_1535_v23).Get((_this).F24()).(_dafny.Int)
					}
					return (_this).F23()
				})(), 4)
				(_nw240).ArraySet1((_this).F24(), 5)
				(_nw240).ArraySet1(((_this).F23()).Minus(_dafny.IntOfUint32((p0).Cardinality())), 6)
				(_nw240).ArraySet1((_dafny.IntOfUint32((_1534_cf23).Cardinality())).Plus((_dafny.Zero).Minus(Companion_Default___.Fm0((_this).F23(), _1504_v0, globalState))), 7)
				(_nw240).ArraySet1(Companion_Default___.Fm0((_dafny.Zero).Minus((_1535_v23).Cardinality()), _1504_v0, globalState), 8)
				(_nw240).ArraySet1((_this).F24(), 9)
				(_nw240).ArraySet1((_this).F24(), 10)
				(_nw240).ArraySet1((_this).F24(), 11)
				_1536_v24 = _nw240
				_1536_v24 = _1536_v24
				var _1537_v25 _dafny.Sequence
				_ = _1537_v25
				_1537_v25 = _dafny.SeqOf(p0, p0)
				var _1538_v26 _dafny.Sequence
				_ = _1538_v26
				_1538_v26 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1537_v25).Cardinality()), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _1506_v2))
				var _1539_v27 D11
				_ = _1539_v27
				_1539_v27 = Companion_D11_.Create_DC32_((_1538_v26).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1538_v26).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _1540_v28 _dafny.MultiSet
				_ = _1540_v28
				_1540_v28 = _dafny.MultiSetOf(_1539_v27, _1539_v27)
				var _1541_v29 _dafny.Map
				_ = _1541_v29
				_1541_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1540_v28, _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F23(), (_this).F24(), _dafny.IntOfUint32((_1534_cf23).Cardinality()), (_this).F24(), (_this).F24()), (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F23(), (_this).F24(), _dafny.IntOfUint32((_1534_cf23).Cardinality()), (_this).F24(), (_this).F24())).Cardinality()))).Uint32(), Companion_Default___.Fm0((_this).F24(), _1504_v0, globalState)))
				var _1542_v30 _dafny.Sequence
				_ = _1542_v30
				_1542_v30 = _dafny.SeqOf((_this).F23(), (_this).F24(), (_this).F23(), (_this).F23(), (_this).F24())
				_1541_v29 = (_1541_v29).Update(_1540_v28, _dafny.Companion_Sequence_.Concatenate(_1542_v30, _1542_v30))
				var _1543_v31 _dafny.Map
				_ = _1543_v31
				_1543_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1504_v0, _dafny.UnicodeSeqOfUtf8Bytes("bxnn"))
				_1543_v31 = (func() _dafny.Map {
					if !(_1504_v0) {
						return (_1543_v31).Merge(_1543_v31)
					}
					return _1543_v31
				})()
				(globalState).F8 = (_dafny.Zero).Minus((_this).F23())
			} else if _source14.Is_DC9() {
				var _1544___mcc_h3 _dafny.MultiSet = _source14.Get_().(D3_DC9).Cf20
				_ = _1544___mcc_h3
				var _1545_cf20 _dafny.MultiSet = _1544___mcc_h3
				_ = _1545_cf20
				(globalState).F9 = _1504_v0
				(globalState).F8 = (_this).F23()
				(globalState).F6 = p0
				var _1546_v32 _dafny.MultiSet
				_ = _1546_v32
				_1546_v32 = _dafny.MultiSetOf(_1504_v0, true, true)
				var _1547_v33 _dafny.Sequence
				_ = _1547_v33
				_1547_v33 = _dafny.SeqOf(_1504_v0, _1504_v0, _1504_v0)
				var _1548_v34 _dafny.Array
				_ = _1548_v34
				var _nwElement0_53 bool = _1504_v0
				_ = _nwElement0_53
				var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(21))
				_ = _nw241
				(_nw241).ArraySet1(_nwElement0_53, 0)
				(_nw241).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32((p0).Cardinality()))).Cmp((_this).F23()) >= 0, 1)
				(_nw241).ArraySet1(!(_1504_v0), 2)
				(_nw241).ArraySet1(!(_1504_v0), 3)
				(_nw241).ArraySet1(_1504_v0, 4)
				(_nw241).ArraySet1(!(_1504_v0), 5)
				(_nw241).ArraySet1(Companion_Default___.Fm1((_this).F23(), _1504_v0, globalState), 6)
				(_nw241).ArraySet1(_dafny.Companion_Sequence_.Contains(p0, _1506_v2), 7)
				(_nw241).ArraySet1(_1504_v0, 8)
				(_nw241).ArraySet1((_1504_v0) == (!(_1504_v0)), 9)
				(_nw241).ArraySet1(true, 10)
				(_nw241).ArraySet1(Companion_Default___.Fm1((_this).F23(), _1504_v0, globalState), 11)
				(_nw241).ArraySet1(_1504_v0, 12)
				(_nw241).ArraySet1((_1546_v32).IsProperSubsetOf(_1546_v32), 13)
				(_nw241).ArraySet1(_1504_v0, 14)
				(_nw241).ArraySet1(_1504_v0, 15)
				(_nw241).ArraySet1(_1504_v0, 16)
				(_nw241).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("thw"), p0), 17)
				(_nw241).ArraySet1(_1504_v0, 18)
				(_nw241).ArraySet1((_1547_v33).Select((Companion_Default___.SafeIndex((_this).F23(), _dafny.IntOfUint32((_1547_v33).Cardinality()))).Uint32()).(bool), 19)
				(_nw241).ArraySet1(_1504_v0, 20)
				_1548_v34 = _nw241
				_1548_v34 = _1548_v34
			} else {
				var _1549___mcc_h4 D3 = _source14.Get_().(D3_DC13).Cf24
				_ = _1549___mcc_h4
				var _1550_cf24 D3 = _1549___mcc_h4
				_ = _1550_cf24
				var _1551_v35 _dafny.MultiSet
				_ = _1551_v35
				_1551_v35 = _dafny.MultiSetOf(_1504_v0)
				var _1552_v36 _dafny.Set
				_ = _1552_v36
				_1552_v36 = _dafny.SetOf(Companion_Default___.Fm35(_1504_v0, !(_1504_v0), _1551_v35, globalState))
				var _1553_v37 _dafny.Sequence
				_ = _1553_v37
				_1553_v37 = _dafny.SeqOf((_this).F24())
				var _1554_v38 _dafny.Sequence
				_ = _1554_v38
				_1554_v38 = _dafny.SeqOf(_1553_v37)
				var _1555_v39 _dafny.Map
				_ = _1555_v39
				_1555_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1504_v0, _dafny.SeqOf((_this).F23(), (_this).F24()))
				var _1556_v40 D12
				_ = _1556_v40
				_1556_v40 = Companion_D12_.Create_DC36_(true, _1504_v0, (_this).F23(), (_this).F24(), _1555_v39)
				var _1557_v41 _dafny.MultiSet
				_ = _1557_v41
				_1557_v41 = _dafny.MultiSetOf((_this).F24(), (_this).F24(), (_dafny.Zero).Minus((_this).F24()), (_this).F24())
				var _1558_v42 _dafny.Array
				_ = _1558_v42
				var _nwElement0_54 _dafny.Int = (_this).F24()
				_ = _nwElement0_54
				var _nw242 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(17))
				_ = _nw242
				(_nw242).ArraySet1(_nwElement0_54, 0)
				(_nw242).ArraySet1((_this).F23(), 1)
				(_nw242).ArraySet1(((_this).F23()).Plus((_this).F23()), 2)
				(_nw242).ArraySet1((_this).F24(), 3)
				(_nw242).ArraySet1((_this).F23(), 4)
				(_nw242).ArraySet1((_this).F23(), 5)
				(_nw242).ArraySet1(((_this).F23()).Times((_this).F24()), 6)
				(_nw242).ArraySet1((_this).F24(), 7)
				(_nw242).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(196), (_this).F24()), 8)
				(_nw242).ArraySet1((_this).F24(), 9)
				(_nw242).ArraySet1((_this).F23(), 10)
				(_nw242).ArraySet1(_dafny.IntOfUint32((_1554_v38).Cardinality()), 11)
				(_nw242).ArraySet1((_1556_v40).Dtor_cf59(), 12)
				(_nw242).ArraySet1(Companion_Default___.Fm0((func() _dafny.Int {
					if (_1557_v41).Contains((_this).F24()) {
						return (_1557_v41).Multiplicity((_this).F24())
					}
					return _dafny.IntOfInt64(-104)
				})(), _1504_v0, globalState), 13)
				(_nw242).ArraySet1((_this).F23(), 14)
				(_nw242).ArraySet1(((_this).F23()).Times((_dafny.MultiSetOf((_this).F23(), (_this).F24(), (_this).F23(), _dafny.IntOfUint32((p0).Cardinality()))).Cardinality()), 15)
				(_nw242).ArraySet1((_this).F23(), 16)
				_1558_v42 = _nw242
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1558_v42), 0))
				_ = _index263
				(_1558_v42).ArraySet1((_this).F24(), (_index263).Int())
				var _1559_v43 _dafny.Set
				_ = _1559_v43
				_1559_v43 = _dafny.SetOf(_1504_v0)
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1558_v42), 0))
				_ = _index264
				var _rhs270 _dafny.Set = _dafny.SetOf(_1559_v43)
				_ = _rhs270
				var _rhs271 _dafny.Int = _dafny.IntOfUint32((p0).Cardinality())
				_ = _rhs271
				var _rhs272 _dafny.Int = ((_dafny.SetOf(_1504_v0, _1504_v0)).Intersection(_1559_v43)).Cardinality()
				_ = _rhs272
				var _lhs238 _dafny.Array = _1558_v42
				_ = _lhs238
				var _lhs239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1558_v42), 0))
				_ = _lhs239
				_1552_v36 = _rhs270
				(_lhs238).ArraySet1(_rhs271, (_lhs239).Int())
				r0 = _rhs272
				var _1560_v44 _dafny.Array
				_ = _1560_v44
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_42
				var _nw243 _dafny.Array
				_ = _nw243
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw243 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) bool = func(_1561_i1 _dafny.Int) bool {
						return true
					}
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw243 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw243).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw243).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1560_v44 = _nw243
				var _1562_v45 D0
				_ = _1562_v45
				var _1563_v46 _dafny.Map
				_ = _1563_v46
				var _1564_v47 _dafny.Int
				_ = _1564_v47
				var _out47 D0
				_ = _out47
				var _out48 _dafny.Map
				_ = _out48
				var _out49 _dafny.Int
				_ = _out49
				_out47, _out48, _out49 = Companion_Default___.M0(_1560_v44, _dafny.UnicodeSeqOfUtf8Bytes("hqminyem"), (_this).F23(), globalState)
				_1562_v45 = _out47
				_1563_v46 = _out48
				_1564_v47 = _out49
				_1504_v0 = ((_this).F23()).Cmp(((_1558_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1558_v42), 0))).Int()).(_dafny.Int)).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-520))).Uint32(), func(coer99 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg99 _dafny.Int) interface{} {
						return coer99(arg99)
					}
				}((func(_1565_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1566_i2 _dafny.Int) _dafny.CodePoint {
						return _1565_v2
					}
				})(_1506_v2)))).Cardinality()))) != 0
				_1505_v1 = (func() _dafny.Map {
					if (_1557_v41).IsProperSubsetOf(_1557_v41) {
						return (func() _dafny.Map {
							if _1504_v0 {
								return _1505_v1
							}
							return _1505_v1
						})()
					}
					return _1505_v1
				})()
			}
			var _1567_v48 _dafny.Array
			_ = _1567_v48
			var _nwElement0_55 bool = _1504_v0
			_ = _nwElement0_55
			var _nw244 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(2))
			_ = _nw244
			(_nw244).ArraySet1(_nwElement0_55, 0)
			(_nw244).ArraySet1(false, 1)
			_1567_v48 = _nw244
			_1567_v48 = _1567_v48
			var _1568_v49 _dafny.Array
			_ = _1568_v49
			var _nw245 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(25))
			_ = _nw245
			_1568_v49 = _nw245
			var _1569_v50 _dafny.Map
			_ = _1569_v50
			_1569_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1504_v0, _1506_v2)
			var _nwElement0_56 _dafny.CodePoint = _1506_v2
			_ = _nwElement0_56
			var _nw246 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(6))
			_ = _nw246
			(_nw246).ArraySet1CodePoint(_nwElement0_56, 0)
			(_nw246).ArraySet1CodePoint(_1506_v2, 1)
			(_nw246).ArraySet1CodePoint(_1506_v2, 2)
			(_nw246).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_1569_v50).Contains(false) {
					return (_1569_v50).Get(false).(_dafny.CodePoint)
				}
				return Companion_Default___.Fm6((_this).F23(), globalState)
			})(), 3)
			(_nw246).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_1569_v50).Contains(_1504_v0) {
					return (_1569_v50).Get(_1504_v0).(_dafny.CodePoint)
				}
				return _1506_v2
			})(), 4)
			(_nw246).ArraySet1CodePoint(Companion_Default___.Fm6((_this).F23(), globalState), 5)
			_1568_v49 = _nw246
			(globalState).F14 = _dafny.IntOfInt64(558)
			var _rhs273 bool = true
			_ = _rhs273
			var _rhs274 bool = false
			_ = _rhs274
			var _lhs240 *GlobalState = globalState
			_ = _lhs240
			var _lhs241 *GlobalState = globalState
			_ = _lhs241
			_lhs240.F9 = _rhs273
			_lhs241.F7 = _rhs274
		} else {
			(globalState).F8 = (_this).F23()
			var _1570_v51 D13
			_ = _1570_v51
			_1570_v51 = Companion_D13_.Create_DC39_((_dafny.Zero).Minus(Companion_Default___.Fm0((_this).F24(), _1504_v0, globalState)))
			_1570_v51 = (func() D13 {
				if _1504_v0 {
					return _1570_v51
				}
				return _1570_v51
			})()
			var _1571_v52 _dafny.MultiSet
			_ = _1571_v52
			_1571_v52 = _dafny.MultiSetOf(true, _1504_v0)
			var _1572_v53 T1
			_ = _1572_v53
			var _nw247 *C3 = New_C3_()
			_ = _nw247
			_nw247.Ctor__(_dafny.SeqOf(_1571_v52, _1571_v52, _1571_v52, _1571_v52), (_this).F23(), (_1504_v0) == (_1504_v0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(205))).Uint32(), func(coer100 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg100 _dafny.Int) interface{} {
					return coer100(arg100)
				}
			}((func(_1573_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_1574_i3 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_1573_p0).Cardinality())
				}
			})(p0))))
			_1572_v53 = _nw247
			var _rhs275 _dafny.CodePoint = _1506_v2
			_ = _rhs275
			var _rhs276 _dafny.Int = (_1572_v53).F23()
			_ = _rhs276
			var _rhs277 bool = Companion_Default___.Fm1((_this).F23(), (_1572_v53).F28(), globalState)
			_ = _rhs277
			var _rhs278 T1 = (func() T1 {
				if (_1572_v53).F28() {
					return _1572_v53
				}
				return _1572_v53
			})()
			_ = _rhs278
			var _lhs242 *GlobalState = globalState
			_ = _lhs242
			var _lhs243 *GlobalState = globalState
			_ = _lhs243
			var _lhs244 *GlobalState = globalState
			_ = _lhs244
			_lhs242.F11 = _rhs275
			_lhs243.F14 = _rhs276
			_lhs244.F21 = _rhs277
			_1572_v53 = _rhs278
			(globalState).F7 = (_1572_v53).F28()
			var _1575_v54 _dafny.Map
			_ = _1575_v54
			_1575_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1572_v53).F28(), p0)
			(globalState).F6 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(350))).Uint32(), func(coer101 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg101 _dafny.Int) interface{} {
					return coer101(arg101)
				}
			}((func(_1576_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1577_i4 _dafny.Int) _dafny.CodePoint {
					return _1576_v2
				}
			})(_1506_v2))), p0), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(635))).Uint32(), func(coer102 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg102 _dafny.Int) interface{} {
					return coer102(arg102)
				}
			}(func(_1578_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})), (func() _dafny.Sequence {
				if (_1575_v54).Contains(_1504_v0) {
					return (_1575_v54).Get(_1504_v0).(_dafny.Sequence)
				}
				return p0
			})()), (Companion_Default___.SafeIndex((_1572_v53).F23(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(635))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg103 _dafny.Int) interface{} {
					return coer103(arg103)
				}
			}(func(_1579_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})), (func() _dafny.Sequence {
				if (_1575_v54).Contains(_1504_v0) {
					return (_1575_v54).Get(_1504_v0).(_dafny.Sequence)
				}
				return p0
			})())).Cardinality()))).Uint32(), Companion_Default___.Fm6(_dafny.IntOfInt64(-66), globalState)))
		}
		r0 = ((_this).F23()).Times((_this).F23())
		var _1580_v55 _dafny.Array
		_ = _1580_v55
		var _nw248 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
		_ = _nw248
		_1580_v55 = _nw248
		for _iter60 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1580_v55), 0))); ; {
			_guard_loop_5, _ok60 := _iter60()
			if !_ok60 {
				break
			}
			var _1581_i6 _dafny.Int
			_1581_i6 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_1581_i6).Sign() != -1) && ((_1581_i6).Cmp(_dafny.ArrayLen((_1580_v55), 0)) < 0)) {
				(_1580_v55).ArraySet1(Companion_Default___.SafeModuloInt(_1581_i6, ((_this).F24()).Times((_this).F24())), (_1581_i6).Int())
			}
		}
		var _1582_v56 _dafny.Map
		_ = _1582_v56
		_1582_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1504_v0)
		var _1583_v57 _dafny.Map
		_ = _1583_v57
		_1583_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1582_v56, (_this).F24())
		(globalState).F9 = (_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("r"), _1506_v2)) || (((_this).F23()).Cmp((_1583_v57).Cardinality()) < 0)
		r0 = (_dafny.IntOfInt64(166)).Minus((_this).F23())
		var _1584_v58 _dafny.Sequence
		_ = _1584_v58
		_1584_v58 = _dafny.SeqOf((_this).F24(), (_this).F23())
		r1 = _1584_v58
		return r0, r1
	}
}
func (_this *C6) F24() _dafny.Int {
	{
		return _this._f24
	}
}

// End of class C6
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
