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
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D12_.Create_DC29_(true, _dafny.IntOfInt64(620), false, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), false)).Dtor_cf52(), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	return ((Companion_D9_.Create_DC22_((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(192))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})), _dafny.IntOfInt64(-91))).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Sequence
			_1_v0 = interface{}(_compr_0).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(192))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			})), _dafny.IntOfInt64(-91))).Contains(_1_v0) {
				_coll0.Add(_1_v0, _dafny.IntOfInt64(865))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality())).Dtor_cf38()).Cmp(_dafny.IntOfInt64(103)) <= 0
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.CodePoint, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(true)), _dafny.CodePoint('y'))).Merge(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(895))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_2_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf(false)
		}))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _3_v0 _dafny.Sequence
			_3_v0 = interface{}(_compr_1).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(895))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_2_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(false)
			})), _3_v0) {
				_coll1.Add(_3_v0, _dafny.CodePoint('e'))
			}
		}
		return _coll1.ToMap()
	}())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), _dafny.CodePoint('m')))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Sequence, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(271), !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("s"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-254))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	}))))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.IntOfInt64(810))).Intersection((func() _dafny.Set {
		if true {
			return func() _dafny.Set {
				var _coll2 = _dafny.NewBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(192), _dafny.IntOfInt64(644))); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _5_v0 _dafny.Int
					_5_v0 = interface{}(_compr_2).(_dafny.Int)
					if ((_dafny.IntOfInt64(192)).Cmp(_5_v0) <= 0) && ((_5_v0).Cmp(_dafny.IntOfInt64(644)) < 0) {
						_coll2.Add((_5_v0).Minus(_dafny.IntOfInt64(-269)))
					}
				}
				return _coll2.ToSet()
			}()
		}
		return _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(-526))
	})())
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(716))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_6_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))).Cardinality()))).Cardinality())).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(38))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_7_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(904))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_8_i2 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()
	}))).Cardinality()))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("l")
}
func (_static *CompanionStruct_Default___) Fm17(globalState *GlobalState) bool {
	return !(!(!((func() bool {
		if true {
			return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-780))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_9_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(true)
			})), _dafny.SeqOf(_dafny.SeqOf(false, false)))
		}
		return (_dafny.SetOf(false)).IsProperSubsetOf(_dafny.SetOf(false))
	})())))
}
func (_static *CompanionStruct_Default___) Fm18(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.SetOf(!(!(true))))).Intersection((_dafny.MultiSetOf(_dafny.SetOf(true), _dafny.SetOf(false, true))).Difference((Companion_D26_.Create_DC67_(_dafny.MultiSetOf(_dafny.SetOf(false, !(false), true), _dafny.SetOf(true, false)))).Dtor_cf111()))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 bool, p2 _dafny.Map, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, true, false, false, !(false)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, false, false, false, false), _dafny.SeqOf(false, false)))
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, globalState *GlobalState) _dafny.Int {
	return (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfInt64(291))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetOf(true), _dafny.MultiSetOf(true), _dafny.MultiSetOf(!(true), true))).Cardinality())))).Cardinality()).Plus((_dafny.Zero).Minus(((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("qeutyx"), _dafny.UnicodeSeqOfUtf8Bytes("dpui"))).Difference((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ts"), _dafny.UnicodeSeqOfUtf8Bytes("jkwix"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-889))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_10_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	})))))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) _dafny.Map {
	if (false) == (true) {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(817), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(746), false))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(494), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lno")).Cardinality()), true))
	}
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(910))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_11_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	}))).Cardinality()), _dafny.IntOfInt64(905)), _dafny.IntOfInt64(-703))
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, globalState *GlobalState) _dafny.Sequence {
	if (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hlbbvk")).Cardinality())).Cmp((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), false)).Cardinality())) != 0 {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, !(true), false, !(!(true)), true), _dafny.SeqOf(!(true)))
	} else if true {
		return _dafny.SeqOf(true, !(false), false)
	} else {
		return _dafny.SeqOf(true)
	}
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.IntOfInt64(-72), _dafny.IntOfInt64(-329))
}
func (_static *CompanionStruct_Default___) Fm25(p0 D4, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vy"), _dafny.UnicodeSeqOfUtf8Bytes("y")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mbgtm"), _dafny.UnicodeSeqOfUtf8Bytes("pacqqiw")))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Set, p1 _dafny.Sequence, p2 D0, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return (Companion_D11_.Create_DC24_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("le"), _dafny.CodePoint('s')))).Dtor_cf40()
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	if false {
		return _dafny.CodePoint('m')
	} else {
		return _dafny.CodePoint('t')
	}
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("roi")
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return (func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(89), _dafny.IntOfInt64(359))); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _12_v1 _dafny.Int
				_12_v1 = interface{}(_compr_4).(_dafny.Int)
				if ((_dafny.IntOfInt64(89)).Cmp(_12_v1) <= 0) && ((_12_v1).Cmp(_dafny.IntOfInt64(359)) < 0) {
					_coll4.Add((_12_v1).Minus(_dafny.IntOfInt64(952)), false)
				}
			}
			return _coll4.ToMap()
		}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Keys().Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _13_v0 _dafny.Int
			_13_v0 = interface{}(_compr_3).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(89), _dafny.IntOfInt64(359))); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _12_v1 _dafny.Int
					_12_v1 = interface{}(_compr_5).(_dafny.Int)
					if ((_dafny.IntOfInt64(89)).Cmp(_12_v1) <= 0) && ((_12_v1).Cmp(_dafny.IntOfInt64(359)) < 0) {
						_coll5.Add((_12_v1).Minus(_dafny.IntOfInt64(952)), false)
					}
				}
				return _coll5.ToMap()
			}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Contains(_13_v0) {
				_coll3.Add((_13_v0).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rqy")).Cardinality()))), _dafny.IntOfInt64(326))
			}
		}
		return _coll3.ToMap()
	}()).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC14_(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("f"), _dafny.UnicodeSeqOfUtf8Bytes("csnarsqlx")))
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) D12 {
	if !(((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('i'))).Cardinality())).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(550), _dafny.IntOfInt64(706))).Cardinality())) > 0) {
		return Companion_D12_.Create_DC28_(true, _dafny.IntOfInt64(-395))
	} else {
		return Companion_D12_.Create_DC28_(!(false), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm35(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfInt64(325))
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Set, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _dafny.IntOfInt64(-65))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xaxw")).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(!(false)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("utuis")).Cardinality()))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, globalState *GlobalState) D7 {
	return Companion_D7_.Create_DC19_(_dafny.IntOfInt64(236))
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(true, true)), _dafny.SeqOf(true, false))
}
func (_static *CompanionStruct_Default___) Fm39(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(888), _dafny.IntOfInt64(779))); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _14_v0 _dafny.Int
			_14_v0 = interface{}(_compr_6).(_dafny.Int)
			if ((_dafny.IntOfInt64(888)).Cmp(_14_v0) <= 0) && ((_14_v0).Cmp(_dafny.IntOfInt64(779)) < 0) {
				_coll6.Add(Companion_Default___.SafeModuloInt(_14_v0, _dafny.IntOfInt64(778)), !(!(false)))
			}
		}
		return _coll6.ToMap()
	}()).Merge((Companion_D28_.Create_DC71_(func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(376))).Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _15_v1 _dafny.Int
			_15_v1 = interface{}(_compr_7).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(376)), _15_v1) {
				_coll7.Add((_15_v1).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), false)
			}
		}
		return _coll7.ToMap()
	}())).Dtor_cf116())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(884), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.MultiSetOf(true, false)).Cardinality()), false)))
}
func (_static *CompanionStruct_Default___) Fm40(p0 bool, p1 _dafny.CodePoint, p2 bool, p3 D7, globalState *GlobalState) D12 {
	if true {
		return Companion_D12_.Create_DC27_(_dafny.SetOf(true))
	} else {
		return Companion_D12_.Create_DC27_(_dafny.SetOf(false, true))
	}
}
func (_static *CompanionStruct_Default___) Fm43(p0 D11, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(997), !(false))).Cardinality()).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vwotwnwen")).Cardinality()))), _dafny.UnicodeSeqOfUtf8Bytes("amno"))
}
func (_static *CompanionStruct_Default___) Fm44(p0 bool, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(862)
}
func (_static *CompanionStruct_Default___) Fm45(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) D14 {
	return Companion_D14_.Create_DC35_((_dafny.IntOfInt64(656)).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, true, true)).Cardinality()))).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(65))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_16_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(584))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_17_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	}))), (_dafny.MultiSetOf(true, !(true))).Cardinality(), (_dafny.MultiSetOf(!(true))).IsSubsetOf(_dafny.MultiSetOf(!(true))), (_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("edklwnr")).Cardinality())) >= 0)
}
func (_static *CompanionStruct_Default___) Fm46(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("x"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(493))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_18_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))), _dafny.UnicodeSeqOfUtf8Bytes("sdfbv"))
}
func (_static *CompanionStruct_Default___) Fm47(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.MultiSetOf(_dafny.CodePoint('i'))).Cardinality()), _dafny.IntOfInt64(-789)) {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false)
	}
}
func (_static *CompanionStruct_Default___) Fm48(p0 D7, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(395))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_19_i0 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(false), false, false), _dafny.IntOfInt64(-383))).Cardinality()
	})))).Union((_dafny.MultiSetOf(_dafny.IntOfInt64(817))).Intersection(_dafny.MultiSetOf((func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(958), _dafny.IntOfInt64(817))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _20_v0 _dafny.Int
			_20_v0 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(958)).Cmp(_20_v0) <= 0) && ((_20_v0).Cmp(_dafny.IntOfInt64(817)) < 0) {
				_coll8.Add(Companion_Default___.SafeDivisionInt(_20_v0, _dafny.IntOfInt64(713)), _dafny.IntOfInt64(855))
			}
		}
		return _coll8.ToMap()
	}()).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kjajll")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm49(p0 _dafny.CodePoint, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_((func() _dafny.Int {
		if false {
			return _dafny.IntOfInt64(814)
		}
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetOf(false), _dafny.MultiSetOf(!(false)))).Cardinality())
	})())
}
func (_static *CompanionStruct_Default___) Fm50(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(970), _dafny.IntOfInt64(345))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _21_v0 _dafny.Int
			_21_v0 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(970)).Cmp(_21_v0) <= 0) && ((_21_v0).Cmp(_dafny.IntOfInt64(345)) < 0) {
				_coll9.Add(Companion_Default___.SafeModuloInt(_21_v0, _dafny.IntOfInt64(169)), !(false))
			}
		}
		return _coll9.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm51(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('k')
}
func (_static *CompanionStruct_Default___) Fm52(p0 bool, p1 D15, p2 bool, globalState *GlobalState) _dafny.Map {
	if (func() bool {
		if false {
			return true
		}
		return true
	})() {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), (_dafny.Zero).Minus((func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), (Companion_D20_.Create_DC53_(false)).Dtor_cf94())).Keys().Elements()); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _22_v0 _dafny.CodePoint
				_22_v0 = interface{}(_compr_10).(_dafny.CodePoint)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), (Companion_D20_.Create_DC53_(false)).Dtor_cf94())).Contains(_22_v0) {
					_coll10.Add(_22_v0, _dafny.IntOfInt64(935))
				}
			}
			return _coll10.ToMap()
		}()).Cardinality()))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(910))).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm53(p0 _dafny.Map, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("souhprl")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ywax")).Cardinality()))).Cardinality(), false)).Cardinality(), _dafny.IntOfInt64(426), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("liy")).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()), _dafny.IntOfInt64(-525))).Cardinality()))), false)).Cardinality()).Cmp(_dafny.IntOfInt64(981)) >= 0, (!(!(false))) == (true), !(!(false) || (true)), true)
}
func (_static *CompanionStruct_Default___) Fm54(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	var _source0 D11 = Companion_D11_.Create_DC25_(_dafny.CodePoint('t'), _dafny.IntOfInt64(75), true)
	_ = _source0
	if _source0.Is_DC25() {
		var _23___mcc_h0 _dafny.CodePoint = _source0.Get_().(D11_DC25).Cf41
		_ = _23___mcc_h0
		var _24___mcc_h1 _dafny.Int = _source0.Get_().(D11_DC25).Cf42
		_ = _24___mcc_h1
		var _25___mcc_h2 bool = _source0.Get_().(D11_DC25).Cf43
		_ = _25___mcc_h2
		var _26_cf43 bool = _25___mcc_h2
		_ = _26_cf43
		var _27_cf42 _dafny.Int = _24___mcc_h1
		_ = _27_cf42
		var _28_cf41 _dafny.CodePoint = _23___mcc_h0
		_ = _28_cf41
		if _26_cf43 {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_cf41, _27_cf42)
		} else {
			return func() _dafny.Map {
				var _coll11 = _dafny.NewMapBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate((_dafny.SeqOf(_28_cf41)).Elements()); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _29_v0 _dafny.CodePoint
					_29_v0 = interface{}(_compr_11).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_28_cf41), _29_v0) {
						_coll11.Add(_29_v0, _27_cf42)
					}
				}
				return _coll11.ToMap()
			}()
		}
	} else if _source0.Is_DC24() {
		var _30___mcc_h3 _dafny.Map = _source0.Get_().(D11_DC24).Cf40
		_ = _30___mcc_h3
		var _31_cf40 _dafny.Map = _30___mcc_h3
		_ = _31_cf40
		if false {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), (_dafny.MultiSetOf(false, true)).Cardinality())
		} else {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _dafny.IntOfInt64(997))
		}
	} else {
		var _32___mcc_h4 D11 = _source0.Get_().(D11_DC26).Cf44
		_ = _32___mcc_h4
		var _33_cf44 D11 = _32___mcc_h4
		_ = _33_cf44
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true))).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm55(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.SeqOf(false), _dafny.SeqOf(false), _dafny.SeqOf(false, false, true, true), _dafny.SeqOf(true, false, true), _dafny.SeqOf(!(true), true, false, false, true))).Difference((_dafny.SetOf(_dafny.SeqOf(false), _dafny.SeqOf(!(false), false))).Union(_dafny.SetOf(_dafny.SeqOf(false, false, false))))
}
func (_static *CompanionStruct_Default___) Fm56(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(false)), func() _dafny.Set {
			var _coll13 = _dafny.NewBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('f'), _dafny.CodePoint('o'), _dafny.CodePoint('y'), _dafny.CodePoint('d'))).Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _34_v1 _dafny.CodePoint
				_34_v1 = interface{}(_compr_13).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('f'), _dafny.CodePoint('o'), _dafny.CodePoint('y'), _dafny.CodePoint('d')), _34_v1) {
					_coll13.Add(_34_v1)
				}
			}
			return _coll13.ToSet()
		}())).Keys().Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _35_v0 _dafny.Sequence
			_35_v0 = interface{}(_compr_12).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(false)), func() _dafny.Set {
				var _coll14 = _dafny.NewBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('f'), _dafny.CodePoint('o'), _dafny.CodePoint('y'), _dafny.CodePoint('d'))).Elements()); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _36_v1 _dafny.CodePoint
					_36_v1 = interface{}(_compr_14).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('f'), _dafny.CodePoint('o'), _dafny.CodePoint('y'), _dafny.CodePoint('d')), _36_v1) {
						_coll14.Add(_36_v1)
					}
				}
				return _coll14.ToSet()
			}())).Contains(_35_v0) {
				_coll12.Add(_35_v0, _dafny.IntOfInt64(421))
			}
		}
		return _coll12.ToMap()
	}()).Merge(func() _dafny.Map {
		var _coll15 = _dafny.NewMapBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(!(false)))).Elements()); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _37_v2 _dafny.Sequence
			_37_v2 = interface{}(_compr_15).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.SeqOf(!(false)))).Contains(_37_v2) {
				_coll15.Add(_37_v2, _dafny.IntOfInt64(753))
			}
		}
		return _coll15.ToMap()
	}())).Merge(func() _dafny.Map {
		var _coll16 = _dafny.NewMapBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(false, false))).Elements()); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _38_v3 _dafny.Sequence
			_38_v3 = interface{}(_compr_16).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.SeqOf(false, false))).Contains(_38_v3) {
				_coll16.Add(_38_v3, _dafny.IntOfInt64(-829))
			}
		}
		return _coll16.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm57(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.CodePoint('g'))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('u')))
}
func (_static *CompanionStruct_Default___) Fm58(p0 D16, p1 bool, p2 _dafny.Set, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.CodePoint('o'))
}
func (_static *CompanionStruct_Default___) Fm59(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("wsdwwpuv"))
}
func (_static *CompanionStruct_Default___) Fm60(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll17 = _dafny.NewMapBuilder()
		_ = _coll17
		for _iter17 := _dafny.Iterate(((func() _dafny.Set {
			if !(true) {
				return _dafny.SetOf(Companion_D11_.Create_DC25_(_dafny.CodePoint('o'), _dafny.IntOfInt64(-847), true), Companion_D11_.Create_DC25_(_dafny.CodePoint('d'), _dafny.IntOfInt64(-798), true))
			}
			return _dafny.SetOf(Companion_D11_.Create_DC25_(_dafny.CodePoint('h'), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(500))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}(func(_39_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
			}))).Cardinality()))).Cardinality()), true))
		})()).Elements()); ; {
			_compr_17, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _40_v0 D11
			_40_v0 = interface{}(_compr_17).(D11)
			if ((func() _dafny.Set {
				if !(true) {
					return _dafny.SetOf(Companion_D11_.Create_DC25_(_dafny.CodePoint('o'), _dafny.IntOfInt64(-847), true), Companion_D11_.Create_DC25_(_dafny.CodePoint('d'), _dafny.IntOfInt64(-798), true))
				}
				return _dafny.SetOf(Companion_D11_.Create_DC25_(_dafny.CodePoint('h'), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(500))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg16 _dafny.Int) interface{} {
						return coer16(arg16)
					}
				}(func(_39_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
				}))).Cardinality()))).Cardinality()), true))
			})()).Contains(_40_v0) {
				_coll17.Add(_40_v0, _dafny.IntOfInt64(-489))
			}
		}
		return _coll17.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm61(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D9 {
	if !(true) {
		return Companion_D9_.Create_DC21_(_dafny.MultiSetOf(false))
	} else {
		return Companion_D9_.Create_DC21_(_dafny.MultiSetOf(true, true))
	}
}
func (_static *CompanionStruct_Default___) Fm62(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) D15 {
	return Companion_D15_.Create_DC39_(_dafny.UnicodeSeqOfUtf8Bytes("liockxaf"), Companion_D9_.Create_DC21_(_dafny.MultiSetFromSeq(_dafny.SeqOf(false))))
}
func (_static *CompanionStruct_Default___) Fm63(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(true)).Intersection(_dafny.MultiSetOf(!(true), true))).Intersection((_dafny.MultiSetOf(true, !(!(false)))).Union(_dafny.MultiSetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm64(globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC10_(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(499), _dafny.IntOfInt64(664)), ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(845), (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-856))).Cardinality()), true)).Cardinality())).Cardinality())).Cardinality()))).Minus(_dafny.IntOfInt64(575)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ga"), _dafny.UnicodeSeqOfUtf8Bytes("bywhwnw"))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm65(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(704))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(true, true, true)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mdjiex")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm66(p0 _dafny.Int, p1 bool, globalState *GlobalState) D16 {
	return Companion_D16_.Create_DC42_(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(149), (_dafny.SetOf(_dafny.IntOfInt64(-16))).Cardinality()), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, false), _dafny.SeqOf(true, false))).Cardinality()))), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(318), _dafny.IntOfInt64(-834)))
}
func (_static *CompanionStruct_Default___) Fm67(p0 _dafny.Set, globalState *GlobalState) D11 {
	return Companion_D11_.Create_DC25_(_dafny.CodePoint('a'), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(691))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_41_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(127)
	})), _dafny.SeqOf((_dafny.SetOf(false)).Cardinality(), (_dafny.SetOf(true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vhdvmehq")).Cardinality())))).Cardinality())), true)
}
func (_static *CompanionStruct_Default___) Fm68(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) D14 {
	return Companion_D14_.Create_DC34_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-61), _dafny.UnicodeSeqOfUtf8Bytes("vj")))
}
func (_static *CompanionStruct_Default___) Fm69(p0 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((func() _dafny.Int {
		if !((Companion_D20_.Create_DC54_(_dafny.UnicodeSeqOfUtf8Bytes("irhb"), _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()), true)).Dtor_cf97()) {
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wfcecnt")).Cardinality()))
		}
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cicqjbm")).Cardinality()))).Cardinality())
	})()), Companion_D15_.Create_DC37_(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(484))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_42_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(313)
	})), _dafny.SeqOf(_dafny.IntOfInt64(-802)))))
}
func (_static *CompanionStruct_Default___) Fm70(p0 bool, globalState *GlobalState) _dafny.Set {
	var _source1 D16 = Companion_D16_.Create_DC43_(_dafny.UnicodeSeqOfUtf8Bytes("eqqftomtx"), _dafny.CodePoint('v'), (_dafny.Zero).Minus((func() _dafny.Map {
		var _coll18 = _dafny.NewMapBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(979), _dafny.IntOfInt64(837))); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _43_v0 _dafny.Int
			_43_v0 = interface{}(_compr_18).(_dafny.Int)
			if ((_dafny.IntOfInt64(979)).Cmp(_43_v0) <= 0) && ((_43_v0).Cmp(_dafny.IntOfInt64(837)) < 0) {
				_coll18.Add((_43_v0).Minus(_dafny.IntOfInt64(-243)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('g'))).Cardinality())).Cardinality())
			}
		}
		return _coll18.ToMap()
	}()).Cardinality()))
	_ = _source1
	if _source1.Is_DC42() {
		var _44___mcc_h0 _dafny.Int = _source1.Get_().(D16_DC42).Cf76
		_ = _44___mcc_h0
		var _45___mcc_h1 _dafny.Int = _source1.Get_().(D16_DC42).Cf77
		_ = _45___mcc_h1
		var _46___mcc_h2 _dafny.Int = _source1.Get_().(D16_DC42).Cf78
		_ = _46___mcc_h2
		var _47_cf78 _dafny.Int = _46___mcc_h2
		_ = _47_cf78
		var _48_cf77 _dafny.Int = _45___mcc_h1
		_ = _48_cf77
		var _49_cf76 _dafny.Int = _44___mcc_h0
		_ = _49_cf76
		return func() _dafny.Set {
			var _coll19 = _dafny.NewBuilder()
			_ = _coll19
			for _iter19 := _dafny.Iterate((_dafny.SeqOf(func() _dafny.Map {
				var _coll20 = _dafny.NewMapBuilder()
				_ = _coll20
				for _iter20 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(782))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}(func(_50_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				}))).Elements()); ; {
					_compr_20, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _51_v1 _dafny.CodePoint
					_51_v1 = interface{}(_compr_20).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(782))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg20 _dafny.Int) interface{} {
							return coer20(arg20)
						}
					}(func(_50_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('d')
					})), _51_v1) {
						_coll20.Add(_51_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-369))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg21 _dafny.Int) interface{} {
								return coer21(arg21)
							}
						}(func(_52_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.UnicodeSeqOfUtf8Bytes("ttkmplyt")
						}))).Cardinality()))
					}
				}
				return _coll20.ToMap()
			}())).Elements()); ; {
				_compr_19, _ok19 := _iter19()
				if !_ok19 {
					break
				}
				var _53_v2 _dafny.Map
				_53_v2 = interface{}(_compr_19).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(func() _dafny.Map {
					var _coll21 = _dafny.NewMapBuilder()
					_ = _coll21
					for _iter21 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(782))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg22 _dafny.Int) interface{} {
							return coer22(arg22)
						}
					}(func(_50_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('d')
					}))).Elements()); ; {
						_compr_21, _ok21 := _iter21()
						if !_ok21 {
							break
						}
						var _51_v1 _dafny.CodePoint
						_51_v1 = interface{}(_compr_21).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(782))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg23 _dafny.Int) interface{} {
								return coer23(arg23)
							}
						}(func(_50_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('d')
						})), _51_v1) {
							_coll21.Add(_51_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-369))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg24 _dafny.Int) interface{} {
									return coer24(arg24)
								}
							}(func(_52_i1 _dafny.Int) _dafny.Sequence {
								return _dafny.UnicodeSeqOfUtf8Bytes("ttkmplyt")
							}))).Cardinality()))
						}
					}
					return _coll21.ToMap()
				}()), _53_v2) {
					_coll19.Add(_53_v2)
				}
			}
			return _coll19.ToSet()
		}()
	} else if _source1.Is_DC43() {
		var _54___mcc_h3 _dafny.Sequence = _source1.Get_().(D16_DC43).Cf79
		_ = _54___mcc_h3
		var _55___mcc_h4 _dafny.CodePoint = _source1.Get_().(D16_DC43).Cf80
		_ = _55___mcc_h4
		var _56___mcc_h5 _dafny.Int = _source1.Get_().(D16_DC43).Cf81
		_ = _56___mcc_h5
		var _57_cf81 _dafny.Int = _56___mcc_h5
		_ = _57_cf81
		var _58_cf80 _dafny.CodePoint = _55___mcc_h4
		_ = _58_cf80
		var _59_cf79 _dafny.Sequence = _54___mcc_h3
		_ = _59_cf79
		return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_58_cf80, _dafny.IntOfUint32((_59_cf79).Cardinality())))).Intersection(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_58_cf80, _57_cf81)))
	} else if _source1.Is_DC44() {
		var _60___mcc_h6 _dafny.Int = _source1.Get_().(D16_DC44).Cf82
		_ = _60___mcc_h6
		var _61_cf82 _dafny.Int = _60___mcc_h6
		_ = _61_cf82
		return (func() _dafny.Set {
			var _coll22 = _dafny.NewBuilder()
			_ = _coll22
			for _iter22 := _dafny.Iterate((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), _61_cf82), func() _dafny.Map {
				var _coll23 = _dafny.NewMapBuilder()
				_ = _coll23
				for _iter23 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.CodePoint('u'))).Elements()); ; {
					_compr_23, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _62_v3 _dafny.CodePoint
					_62_v3 = interface{}(_compr_23).(_dafny.CodePoint)
					if (_dafny.MultiSetOf(_dafny.CodePoint('u'))).Contains(_62_v3) {
						_coll23.Add(_62_v3, _61_cf82)
					}
				}
				return _coll23.ToMap()
			}())).Elements()); ; {
				_compr_22, _ok22 := _iter22()
				if !_ok22 {
					break
				}
				var _63_v4 _dafny.Map
				_63_v4 = interface{}(_compr_22).(_dafny.Map)
				if (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), _61_cf82), func() _dafny.Map {
					var _coll24 = _dafny.NewMapBuilder()
					_ = _coll24
					for _iter24 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.CodePoint('u'))).Elements()); ; {
						_compr_24, _ok24 := _iter24()
						if !_ok24 {
							break
						}
						var _62_v3 _dafny.CodePoint
						_62_v3 = interface{}(_compr_24).(_dafny.CodePoint)
						if (_dafny.MultiSetOf(_dafny.CodePoint('u'))).Contains(_62_v3) {
							_coll24.Add(_62_v3, _61_cf82)
						}
					}
					return _coll24.ToMap()
				}())).Contains(_63_v4) {
					_coll22.Add(_63_v4)
				}
			}
			return _coll22.ToSet()
		}()).Intersection(func() _dafny.Set {
			var _coll25 = _dafny.NewBuilder()
			_ = _coll25
			for _iter25 := _dafny.Iterate((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), _dafny.IntOfInt64(815)))).Elements()); ; {
				_compr_25, _ok25 := _iter25()
				if !_ok25 {
					break
				}
				var _64_v5 _dafny.Map
				_64_v5 = interface{}(_compr_25).(_dafny.Map)
				if (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), _dafny.IntOfInt64(815)))).Contains(_64_v5) {
					_coll25.Add(_64_v5)
				}
			}
			return _coll25.ToSet()
		}())
	} else {
		var _65___mcc_h7 _dafny.Array = _source1.Get_().(D16_DC41).Cf75
		_ = _65___mcc_h7
		var _66_cf75 _dafny.Array = _65___mcc_h7
		_ = _66_cf75
		return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("imekcx")).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), _dafny.IntOfInt64(-246)), func() _dafny.Map {
			var _coll26 = _dafny.NewMapBuilder()
			_ = _coll26
			for _iter26 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), false)).Keys().Elements()); ; {
				_compr_26, _ok26 := _iter26()
				if !_ok26 {
					break
				}
				var _67_v6 _dafny.CodePoint
				_67_v6 = interface{}(_compr_26).(_dafny.CodePoint)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), false)).Contains(_67_v6) {
					_coll26.Add(_67_v6, (_dafny.Zero).Minus(_dafny.IntOfInt64(-297)))
				}
			}
			return _coll26.ToMap()
		}())).Union(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), _dafny.IntOfInt64(85)), func() _dafny.Map {
			var _coll27 = _dafny.NewMapBuilder()
			_ = _coll27
			for _iter27 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('q'))).Elements()); ; {
				_compr_27, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _68_v7 _dafny.CodePoint
				_68_v7 = interface{}(_compr_27).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('q')), _68_v7) {
					_coll27.Add(_68_v7, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-771))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg25 _dafny.Int) interface{} {
							return coer25(arg25)
						}
					}(func(_69_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('d')
					}))).Cardinality()))).Cardinality())
				}
			}
			return _coll27.ToMap()
		}()))
	}
}
func (_static *CompanionStruct_Default___) Fm71(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SeqOf(false, !(false)), _dafny.SeqOf(!(true)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, true, true), _dafny.SeqOf(!(true))))
}
func (_static *CompanionStruct_Default___) Fm72(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(395))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg26 _dafny.Int) interface{} {
			return coer26(arg26)
		}
	}(func(_70_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	})), _dafny.UnicodeSeqOfUtf8Bytes("usta")))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _71_globalState *GlobalState
	_ = _71_globalState
	var _nw0 *GlobalState = New_GlobalState_()
	_ = _nw0
	_nw0.Ctor__()
	_71_globalState = _nw0
	var _72_v0 _dafny.Int
	_ = _72_v0
	_72_v0 = _dafny.IntOfInt64(936)
	var _73_v1 *C4
	_ = _73_v1
	var _nw1 *C4 = New_C4_()
	_ = _nw1
	_nw1.Ctor__()
	_73_v1 = _nw1
	var _74_v3 _dafny.Sequence
	_ = _74_v3
	_74_v3 = _dafny.UnicodeSeqOfUtf8Bytes("dtn")
	var _75_v4 _dafny.MultiSet
	_ = _75_v4
	_75_v4 = _dafny.MultiSetOf(_74_v3, _74_v3)
	var _76_v5 _dafny.Array
	_ = _76_v5
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(5)
	_ = _len0_0
	var _nw2 _dafny.Array
	_ = _nw2
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw2 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Set = func(_77_i0 _dafny.Int) _dafny.Set {
			return _dafny.SetOf(true, true)
		}
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
	_76_v5 = _nw2
	var _78_v6 D0
	_ = _78_v6
	_78_v6 = Companion_D0_.Create_DC0_(_76_v5, _72_v0)
	var _79_v7 _dafny.Sequence
	_ = _79_v7
	_79_v7 = _dafny.SeqOf(_72_v0)
	var _80_v8 D0
	_ = _80_v8
	_80_v8 = Companion_D0_.Create_DC1_(_72_v0)
	var _81_v9 *C8
	_ = _81_v9
	var _nw3 *C8 = New_C8_()
	_ = _nw3
	_nw3.Ctor__(Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_72_v0)).Cardinality()), _73_v1)).Cardinality(), (func() _dafny.Map {
		var _coll28 = _dafny.NewMapBuilder()
		_ = _coll28
		for _iter28 := _dafny.Iterate((_75_v4).Elements()); ; {
			_compr_28, _ok28 := _iter28()
			if !_ok28 {
				break
			}
			var _82_v2 _dafny.Sequence
			_82_v2 = interface{}(_compr_28).(_dafny.Sequence)
			if (_75_v4).Contains(_82_v2) {
				_coll28.Add(_82_v2, _dafny.IntOfInt64(-676))
			}
		}
		return _coll28.ToMap()
	}()).Cardinality()), _78_v6, _dafny.Companion_Sequence_.Concatenate(_79_v7, _79_v7), _80_v8)
	_81_v9 = _nw3
	var _83_v10 bool
	_ = _83_v10
	_83_v10 = false
	var _pat_let_tv0 = _72_v0
	_ = _pat_let_tv0
	var _84_v11 _dafny.Int
	_ = _84_v11
	var _85_v12 _dafny.Map
	_ = _85_v12
	var _86_v13 _dafny.Int
	_ = _86_v13
	var _out0 _dafny.Int
	_ = _out0
	var _out1 _dafny.Map
	_ = _out1
	var _out2 _dafny.Int
	_ = _out2
	_out0, _out1, _out2 = (_81_v9).M3(_dafny.IntOfInt64(490), func(_pat_let0_0 D0) D0 {
		return func(_87_dt__update__tmp_h0 D0) D0 {
			return func(_pat_let1_0 _dafny.Int) D0 {
				return func(_88_dt__update_hcf1_h0 _dafny.Int) D0 {
					return Companion_D0_.Create_DC0_((_87_dt__update__tmp_h0).Dtor_cf0(), _88_dt__update_hcf1_h0)
				}(_pat_let1_0)
			}(_pat_let_tv0)
		}(_pat_let0_0)
	}(Companion_D0_.Create_DC0_(_76_v5, _dafny.IntOfUint32((_74_v3).Cardinality()))), _83_v10, _71_globalState)
	_84_v11 = _out0
	_85_v12 = _out1
	_86_v13 = _out2
	var _89_v14 _dafny.Map
	_ = _89_v14
	_89_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v10, _dafny.IntOfInt64(548))
	var _90_v15 _dafny.Map
	_ = _90_v15
	_90_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v14, _74_v3)
	var _91_v16 D18
	_ = _91_v16
	_91_v16 = Companion_D18_.Create_DC47_(_90_v15)
	var _pat_let_tv1 = _91_v16
	_ = _pat_let_tv1
	var _source2 D18 = (func() D18 {
		if _83_v10 {
			return func(_pat_let2_0 D18) D18 {
				return func(_92_dt__update__tmp_h1 D18) D18 {
					return func(_pat_let3_0 _dafny.Map) D18 {
						return func(_93_dt__update_hcf85_h0 _dafny.Map) D18 {
							return Companion_D18_.Create_DC47_(_93_dt__update_hcf85_h0)
						}(_pat_let3_0)
					}((_pat_let_tv1).Dtor_cf85())
				}(_pat_let2_0)
			}(_91_v16)
		}
		return _91_v16
	})()
	_ = _source2
	if _source2.Is_DC48() {
		if _83_v10 {
			var _94_v17 _dafny.Array
			_ = _94_v17
			var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
			_ = _nw4
			_94_v17 = _nw4
			_94_v17 = _94_v17
			var _95_v18 _dafny.Array
			_ = _95_v18
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw5
			_95_v18 = _nw5
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_95_v18), 0))
			_ = _index0
			(_95_v18).ArraySet1(_83_v10, (_index0).Int())
			var _96_v19 _dafny.MultiSet
			_ = _96_v19
			_96_v19 = _dafny.MultiSetOf(_84_v11)
			var _pat_let_tv2 = _81_v9
			_ = _pat_let_tv2
			var _97_v20 T1
			_ = _97_v20
			var _nw6 *C7 = New_C7_()
			_ = _nw6
			_nw6.Ctor__(_96_v19, func(_pat_let4_0 D0) D0 {
				return func(_98_dt__update__tmp_h2 D0) D0 {
					return func(_pat_let5_0 _dafny.Int) D0 {
						return func(_99_dt__update_hcf2_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC1_(_99_dt__update_hcf2_h0)
						}(_pat_let5_0)
					}((_pat_let_tv2).F2())
				}(_pat_let4_0)
			}(_80_v8), _dafny.Companion_Sequence_.Concatenate(_79_v7, _79_v7), Companion_D0_.Create_DC1_(_72_v0))
			_97_v20 = _nw6
			var _100_v21 _dafny.Array
			_ = _100_v21
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw7
			_100_v21 = _nw7
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_100_v21), 0))
			_ = _index1
			(_100_v21).ArraySet1((Companion_Default___.Fm70(_83_v10, _71_globalState)).Cardinality(), (_index1).Int())
			var _101_v22 _dafny.CodePoint
			_ = _101_v22
			_101_v22 = _dafny.CodePoint('e')
			var _102_v23 D16
			_ = _102_v23
			_102_v23 = Companion_D16_.Create_DC43_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(403))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}(func(_103_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('c')
			})), _101_v22, _84_v11)
			var _104_v24 *C0
			_ = _104_v24
			var _nw8 *C0 = New_C0_()
			_ = _nw8
			_nw8.Ctor__((_102_v23).Dtor_cf79())
			_104_v24 = _nw8
			var _105_v25 _dafny.Map
			_ = _105_v25
			_105_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_104_v24, _83_v10)
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_95_v18), 0))
			_ = _index2
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_100_v21), 0))
			_ = _index3
			var _rhs0 bool = _83_v10
			_ = _rhs0
			var _rhs1 T1 = (func() T1 {
				if (func() bool {
					if (_105_v25).Contains(_104_v24) {
						return (_105_v25).Get(_104_v24).(bool)
					}
					return _83_v10
				})() {
					return _97_v20
				}
				return _97_v20
			})()
			_ = _rhs1
			var _rhs2 _dafny.Array = _95_v18
			_ = _rhs2
			var _rhs3 _dafny.Int = (_81_v9).F2()
			_ = _rhs3
			var _rhs4 bool = ((Companion_Default___.Fm44(true, _71_globalState)).Cmp(_86_v13) < 0) && (_83_v10)
			_ = _rhs4
			var _lhs0 _dafny.Array = _95_v18
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_95_v18), 0))
			_ = _lhs1
			var _lhs2 _dafny.Array = _100_v21
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_100_v21), 0))
			_ = _lhs3
			(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
			_97_v20 = _rhs1
			_95_v18 = _rhs2
			(_lhs2).ArraySet1(_rhs3, (_lhs3).Int())
			_83_v10 = _rhs4
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_95_v18), 0))
			_ = _index4
			(_95_v18).ArraySet1(!(!(_89_v14).Contains(_83_v10)) || (true), (_index4).Int())
			var _106_v26 _dafny.Map
			_ = _106_v26
			_106_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_95_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_95_v18), 0))).Int()).(bool), Companion_Default___.Fm17(_71_globalState))
			var _107_v27 _dafny.Sequence
			_ = _107_v27
			_107_v27 = _dafny.SeqOf((_95_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_95_v18), 0))).Int()).(bool), false, _83_v10, _83_v10, (_95_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_95_v18), 0))).Int()).(bool))
			var _108_v28 _dafny.MultiSet
			_ = _108_v28
			_108_v28 = _dafny.MultiSetOf((_107_v27).Select((Companion_Default___.SafeIndex(_86_v13, _dafny.IntOfUint32((_107_v27).Cardinality()))).Uint32()).(bool))
			var _109_v29 _dafny.Map
			_ = _109_v29
			_109_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_81_v9).F2(), ((_dafny.MultiSetOf(!((func() bool {
				if (_106_v26).Contains(_83_v10) {
					return (_106_v26).Get(_83_v10).(bool)
				}
				return (_95_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_95_v18), 0))).Int()).(bool)
			})()))).Difference(_108_v28)).Cardinality())
			_109_v29 = (_109_v29).Update(_86_v13, (_81_v9).F2())
			var _110_v30 bool
			_ = _110_v30
			var _111_v31 bool
			_ = _111_v31
			var _out3 bool
			_ = _out3
			var _out4 bool
			_ = _out4
			_out3, _out4 = (_81_v9).M1((func() _dafny.Int {
				if (_109_v29).Contains((_100_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_100_v21), 0))).Int()).(_dafny.Int)) {
					return (_109_v29).Get((_100_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_100_v21), 0))).Int()).(_dafny.Int)).(_dafny.Int)
				}
				return (_102_v23).Dtor_cf81()
			})(), _71_globalState)
			_110_v30 = _out3
			_111_v31 = _out4
		} else {
			var _112_v32 _dafny.CodePoint
			_ = _112_v32
			_112_v32 = _dafny.CodePoint('c')
			var _113_v33 _dafny.Map
			_ = _113_v33
			_113_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v32, (_dafny.Zero).Minus(Companion_Default___.Fm32((_81_v9).F2(), _71_globalState)))
			var _114_v34 _dafny.Array
			_ = _114_v34
			var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw9
			_114_v34 = _nw9
			var _115_v35 T1
			_ = _115_v35
			var _nw10 *C5 = New_C5_()
			_ = _nw10
			_nw10.Ctor__(_114_v34, _79_v7, _80_v8)
			_115_v35 = _nw10
			var _116_v36 *C0
			_ = _116_v36
			var _nw11 *C0 = New_C0_()
			_ = _nw11
			_nw11.Ctor__(_74_v3)
			_116_v36 = _nw11
			var _117_v37 _dafny.Map
			_ = _117_v37
			_117_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_115_v35, _116_v36)
			var _118_v38 _dafny.Array
			_ = _118_v38
			var _nwElement0_0 bool = false
			_ = _nwElement0_0
			var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(29))
			_ = _nw12
			(_nw12).ArraySet1(_nwElement0_0, 0)
			(_nw12).ArraySet1(_83_v10, 1)
			(_nw12).ArraySet1(_83_v10, 2)
			(_nw12).ArraySet1(_83_v10, 3)
			(_nw12).ArraySet1(_83_v10, 4)
			(_nw12).ArraySet1(_83_v10, 5)
			(_nw12).ArraySet1(_83_v10, 6)
			(_nw12).ArraySet1(_83_v10, 7)
			(_nw12).ArraySet1(_83_v10, 8)
			(_nw12).ArraySet1(_83_v10, 9)
			(_nw12).ArraySet1(_83_v10, 10)
			(_nw12).ArraySet1(_83_v10, 11)
			(_nw12).ArraySet1(_83_v10, 12)
			(_nw12).ArraySet1(true, 13)
			(_nw12).ArraySet1(_83_v10, 14)
			(_nw12).ArraySet1(!(_83_v10), 15)
			(_nw12).ArraySet1(Companion_Default___.Fm17(_71_globalState), 16)
			(_nw12).ArraySet1(true, 17)
			(_nw12).ArraySet1(true, 18)
			(_nw12).ArraySet1(_83_v10, 19)
			(_nw12).ArraySet1(_83_v10, 20)
			(_nw12).ArraySet1(false, 21)
			(_nw12).ArraySet1(_83_v10, 22)
			(_nw12).ArraySet1(_83_v10, 23)
			(_nw12).ArraySet1(_83_v10, 24)
			(_nw12).ArraySet1(true, 25)
			(_nw12).ArraySet1(_83_v10, 26)
			(_nw12).ArraySet1(_83_v10, 27)
			(_nw12).ArraySet1(_83_v10, 28)
			_118_v38 = _nw12
			var _119_v39 bool
			_ = _119_v39
			var _120_v40 _dafny.Int
			_ = _120_v40
			var _out5 bool
			_ = _out5
			var _out6 _dafny.Int
			_ = _out6
			_out5, _out6 = (_73_v1).M10(((_81_v9).F2()).Plus((_81_v9).F2()), ((_113_v33).Update(Companion_Default___.Fm51(_83_v10, (_81_v9).F2(), (_117_v37).Cardinality(), _71_globalState), _72_v0)).Cardinality(), _83_v10, _118_v38, _71_globalState)
			_119_v39 = _out5
			_120_v40 = _out6
			var _121_v41 _dafny.Sequence
			_ = _121_v41
			_121_v41 = _dafny.SeqOf(_83_v10)
			_121_v41 = _dafny.SeqOf((_119_v39) == (_83_v10), _119_v39, false, _83_v10)
			_84_v11 = _dafny.IntOfInt64(-509)
			var _122_v42 _dafny.Sequence
			_ = _122_v42
			_122_v42 = _dafny.SeqOf(_118_v38, _118_v38, _118_v38)
			_83_v10 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_118_v38, _118_v38), _122_v42)
			var _123_v43 _dafny.Map
			_ = _123_v43
			_123_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v10, false)
			var _124_v44 _dafny.Int
			_ = _124_v44
			var _125_v45 _dafny.Map
			_ = _125_v45
			var _126_v46 _dafny.Int
			_ = _126_v46
			var _out7 _dafny.Int
			_ = _out7
			var _out8 _dafny.Map
			_ = _out8
			var _out9 _dafny.Int
			_ = _out9
			_out7, _out8, _out9 = (_81_v9).M3((_dafny.Zero).Minus((_dafny.Zero).Minus(_86_v13)), _78_v6, (func() bool {
				if (_123_v43).Contains(false) {
					return (_123_v43).Get(false).(bool)
				}
				return !(_83_v10)
			})(), _71_globalState)
			_124_v44 = _out7
			_125_v45 = _out8
			_126_v46 = _out9
		}
		var _127_v47 _dafny.Sequence
		_ = _127_v47
		_127_v47 = _dafny.SeqOf(!(_83_v10), _83_v10)
		var _128_v48 _dafny.Set
		_ = _128_v48
		_128_v48 = _dafny.SetOf(_dafny.CodePoint('f'))
		var _129_v49 _dafny.Map
		_ = _129_v49
		_129_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_127_v47, _dafny.Companion_Sequence_.Update(_127_v47, (Companion_Default___.SafeIndex((_128_v48).Cardinality(), _dafny.IntOfUint32((_127_v47).Cardinality()))).Uint32(), _83_v10))
		_129_v49 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_127_v47, _127_v47)).Merge((_129_v49).Merge(_129_v49))
		_83_v10 = _83_v10
		var _130_v50 T0
		_ = _130_v50
		var _nw13 *C4 = New_C4_()
		_ = _nw13
		_nw13.Ctor__()
		_130_v50 = _nw13
		var _131_v51 D20
		_ = _131_v51
		_131_v51 = Companion_D20_.Create_DC53_(_83_v10)
		var _132_v52 _dafny.Map
		_ = _132_v52
		_132_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_130_v50, (_131_v51).Dtor_cf94())
		var _133_v53 T0
		_ = _133_v53
		_133_v53 = _130_v50
		_132_v52 = (_132_v52).Update((_133_v53), true)
	} else if _source2.Is_DC49() {
		var _134___mcc_h0 _dafny.Set = _source2.Get_().(D18_DC49).Cf86
		_ = _134___mcc_h0
		var _135_cf86 _dafny.Set = _134___mcc_h0
		_ = _135_cf86
		if _83_v10 {
			var _136_v54 _dafny.CodePoint
			_ = _136_v54
			_136_v54 = _dafny.CodePoint('n')
			var _137_v55 _dafny.Map
			_ = _137_v55
			_137_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_136_v54, _80_v8)
			var _138_v56 _dafny.Map
			_ = _138_v56
			_138_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_Default___.Fm71((_81_v9).Fm5((_81_v9).F2(), _83_v10, _137_v55, _83_v10, _71_globalState), _71_globalState))
			var _139_v57 _dafny.Sequence
			_ = _139_v57
			_139_v57 = _dafny.SeqOf(_83_v10, _83_v10, _83_v10)
			var _140_v58 _dafny.MultiSet
			_ = _140_v58
			_140_v58 = _dafny.MultiSetOf(_83_v10, _83_v10)
			var _141_v59 _dafny.Sequence
			_ = _141_v59
			_141_v59 = _dafny.SeqOf(_dafny.SeqOf(_83_v10, !(_83_v10)))
			_84_v11 = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_138_v56).Contains((_140_v58).IsProperSubsetOf(_dafny.MultiSetFromSeq(_139_v57))) {
					return (_138_v56).Get((_140_v58).IsProperSubsetOf(_dafny.MultiSetFromSeq(_139_v57))).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Concatenate(_141_v59, _141_v59)
			})()).Cardinality())
			var _142_v60 _dafny.Set
			_ = _142_v60
			_142_v60 = _dafny.SetOf(_83_v10)
			_89_v14 = (_89_v14).Update((_142_v60).IsProperSubsetOf(_142_v60), (_81_v9).F2())
			var _143_v62 _dafny.MultiSet
			_ = _143_v62
			_143_v62 = _dafny.MultiSetOf((func() _dafny.Set {
				var _coll29 = _dafny.NewBuilder()
				_ = _coll29
				for _iter29 := _dafny.Iterate((_79_v7).Elements()); ; {
					_compr_29, _ok29 := _iter29()
					if !_ok29 {
						break
					}
					var _144_v61 _dafny.Int
					_144_v61 = interface{}(_compr_29).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_79_v7, _144_v61) {
						_coll29.Add((_144_v61).Minus(_dafny.IntOfInt64(185)))
					}
				}
				return _coll29.ToSet()
			}()).Cardinality(), _dafny.IntOfInt64(12))
			var _145_v63 *C7
			_ = _145_v63
			var _nw14 *C7 = New_C7_()
			_ = _nw14
			_nw14.Ctor__((((_143_v62).Update((_81_v9).F2(), Companion_Default___.Abs(_72_v0))).Update(_72_v0, Companion_Default___.Abs(_86_v13))).Update(_72_v0, Companion_Default___.Abs(_72_v0)), (func() D0 {
				if _83_v10 {
					return Companion_Default___.Fm49(_136_v54, _71_globalState)
				}
				return _80_v8
			})(), _79_v7, _80_v8)
			_145_v63 = _nw14
			var _146_v64 _dafny.Array
			_ = _146_v64
			var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
			_ = _nw15
			_146_v64 = _nw15
			var _147_v65 *C1
			_ = _147_v65
			var _nw16 *C1 = New_C1_()
			_ = _nw16
			_nw16.Ctor__(_84_v11, _146_v64, _dafny.Companion_Sequence_.Update(_79_v7, (Companion_Default___.SafeIndex((_81_v9).F2(), _dafny.IntOfUint32((_79_v7).Cardinality()))).Uint32(), _72_v0), Companion_D0_.Create_DC1_(_72_v0))
			_147_v65 = _nw16
			var _148_v66 _dafny.Map
			_ = _148_v66
			_148_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_147_v65, _136_v54)
			_148_v66 = (_148_v66).Update(_147_v65, _136_v54)
			var _149_v67 T0
			_ = _149_v67
			var _nw17 *C8 = New_C8_()
			_ = _nw17
			_nw17.Ctor__(_147_v65.F9, _78_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(169))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}((func(_150_v9 *C8) func(_dafny.Int) _dafny.Int {
				return func(_151_i2 _dafny.Int) _dafny.Int {
					return (_150_v9).F2()
				}
			})(_81_v9))), Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_139_v57).Cardinality())))
			_149_v67 = _nw17
			var _152_v68 _dafny.Map
			_ = _152_v68
			_152_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_149_v67, (_81_v9).F2())
			var _153_v69 _dafny.Array
			_ = _153_v69
			var _nw18 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
			_ = _nw18
			_153_v69 = _nw18
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_153_v69), 0))
			_ = _index5
			(_153_v69).ArraySet1(_83_v10, (_index5).Int())
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_153_v69), 0))
			_ = _index6
			var _rhs5 _dafny.Map = (_152_v68).Merge(_152_v68)
			_ = _rhs5
			var _rhs6 bool = _83_v10
			_ = _rhs6
			var _lhs4 _dafny.Array = _153_v69
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_153_v69), 0))
			_ = _lhs5
			_152_v68 = _rhs5
			(_lhs4).ArraySet1(_rhs6, (_lhs5).Int())
		} else {
			var _154_v70 _dafny.Array
			_ = _154_v70
			var _nw19 _dafny.Array = _dafny.NewArrayWithValue(Companion_D17_.Default(), _dafny.One)
			_ = _nw19
			_154_v70 = _nw19
			var _155_v71 _dafny.MultiSet
			_ = _155_v71
			_155_v71 = _dafny.MultiSetOf(_86_v13, (_81_v9).F2())
			var _156_v72 D17
			_ = _156_v72
			_156_v72 = Companion_D17_.Create_DC45_(_dafny.MultiSetOf(_155_v71, _155_v71))
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_154_v70), 0))
			_ = _index7
			(_154_v70).ArraySet1(_156_v72, (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_154_v70), 0))
			_ = _index8
			(_154_v70).ArraySet1(_156_v72, (_index8).Int())
			var _157_v73 *C4
			_ = _157_v73
			_157_v73 = _73_v1
			var _158_v74 _dafny.Array
			_ = _158_v74
			var _nwElement0_1 *C4 = _73_v1
			_ = _nwElement0_1
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(28))
			_ = _nw20
			(_nw20).ArraySet1(_nwElement0_1, 0)
			(_nw20).ArraySet1(_73_v1, 1)
			(_nw20).ArraySet1(_73_v1, 2)
			(_nw20).ArraySet1(_73_v1, 3)
			(_nw20).ArraySet1(_73_v1, 4)
			(_nw20).ArraySet1(_73_v1, 5)
			(_nw20).ArraySet1(_73_v1, 6)
			(_nw20).ArraySet1(_73_v1, 7)
			(_nw20).ArraySet1(_73_v1, 8)
			(_nw20).ArraySet1(_73_v1, 9)
			(_nw20).ArraySet1(_73_v1, 10)
			(_nw20).ArraySet1(_73_v1, 11)
			(_nw20).ArraySet1(_73_v1, 12)
			(_nw20).ArraySet1(_73_v1, 13)
			(_nw20).ArraySet1(_73_v1, 14)
			(_nw20).ArraySet1(_73_v1, 15)
			(_nw20).ArraySet1(_73_v1, 16)
			(_nw20).ArraySet1(_73_v1, 17)
			(_nw20).ArraySet1(_73_v1, 18)
			(_nw20).ArraySet1(_73_v1, 19)
			(_nw20).ArraySet1(_73_v1, 20)
			(_nw20).ArraySet1((_157_v73), 21)
			(_nw20).ArraySet1(_73_v1, 22)
			(_nw20).ArraySet1(_73_v1, 23)
			(_nw20).ArraySet1(_73_v1, 24)
			(_nw20).ArraySet1(_73_v1, 25)
			(_nw20).ArraySet1(_73_v1, 26)
			(_nw20).ArraySet1(_73_v1, 27)
			_158_v74 = _nw20
			var _159_v75 _dafny.Sequence
			_ = _159_v75
			_159_v75 = _dafny.SeqOf(_158_v74, _158_v74, _158_v74, _158_v74)
			var _rhs7 _dafny.Int = (_81_v9).F2()
			_ = _rhs7
			var _rhs8 _dafny.Sequence = (_81_v9).Fm7(_86_v13, _71_globalState)
			_ = _rhs8
			var _rhs9 _dafny.Sequence = (func() _dafny.Sequence {
				if _83_v10 {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_158_v74, _158_v74), _dafny.SeqOf(_158_v74))
				}
				return _159_v75
			})()
			_ = _rhs9
			_84_v11 = _rhs7
			_74_v3 = _rhs8
			_159_v75 = _rhs9
			var _160_v76 _dafny.Array
			_ = _160_v76
			var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
			_ = _nw21
			_160_v76 = _nw21
			var _161_v77 D4
			_ = _161_v77
			_161_v77 = Companion_D4_.Create_DC10_((_dafny.Zero).Minus((_84_v11).Plus((_81_v9).F2())), _86_v13, _dafny.IntOfUint32((_74_v3).Cardinality()))
			var _162_v78 _dafny.Sequence
			_ = _162_v78
			_162_v78 = _dafny.SeqOf(_160_v76, _160_v76, _160_v76, _160_v76)
			var _163_v79 _dafny.Map
			_ = _163_v79
			_163_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v10, (_162_v78).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm44(_83_v10, _71_globalState), _dafny.IntOfUint32((_162_v78).Cardinality()))).Uint32()).(_dafny.Array))
			var _rhs10 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(335))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}(func(_164_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(235))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}(func(_165_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})))).Cardinality())
			_ = _rhs10
			var _rhs11 _dafny.Array = (func() _dafny.Array {
				if (_163_v79).Contains(true) {
					return (_163_v79).Get(true).(_dafny.Array)
				}
				return _160_v76
			})()
			_ = _rhs11
			var _rhs12 D4 = _161_v77
			_ = _rhs12
			var _rhs13 _dafny.Int = (func() _dafny.Int {
				if _83_v10 {
					return (_81_v9).F2()
				}
				return Companion_Default___.Fm32((_dafny.Zero).Minus(_72_v0), _71_globalState)
			})()
			_ = _rhs13
			_84_v11 = _rhs10
			_160_v76 = _rhs11
			_161_v77 = _rhs12
			_84_v11 = _rhs13
			_84_v11 = ((_72_v0).Plus(_dafny.IntOfInt64(-618))).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_74_v3).Cardinality())))
			_84_v11 = (func() _dafny.Int {
				if (_75_v4).Contains(_74_v3) {
					return (_75_v4).Multiplicity(_74_v3)
				}
				return _84_v11
			})()
		}
		_83_v10 = _83_v10
		var _166_v80 D7
		_ = _166_v80
		_166_v80 = Companion_D7_.Create_DC18_()
		_166_v80 = _166_v80
		_83_v10 = (_83_v10) == ((_84_v11).Cmp((_81_v9).F2()) == 0)
	} else if _source2.Is_DC50() {
		var _167___mcc_h1 bool = _source2.Get_().(D18_DC50).Cf87
		_ = _167___mcc_h1
		var _168___mcc_h2 _dafny.Array = _source2.Get_().(D18_DC50).Cf88
		_ = _168___mcc_h2
		var _169___mcc_h3 bool = _source2.Get_().(D18_DC50).Cf89
		_ = _169___mcc_h3
		var _170___mcc_h4 bool = _source2.Get_().(D18_DC50).Cf90
		_ = _170___mcc_h4
		var _171___mcc_h5 _dafny.Sequence = _source2.Get_().(D18_DC50).Cf91
		_ = _171___mcc_h5
		var _172_cf91 _dafny.Sequence = _171___mcc_h5
		_ = _172_cf91
		var _173_cf90 bool = _170___mcc_h4
		_ = _173_cf90
		var _174_cf89 bool = _169___mcc_h3
		_ = _174_cf89
		var _175_cf88 _dafny.Array = _168___mcc_h2
		_ = _175_cf88
		var _176_cf87 bool = _167___mcc_h1
		_ = _176_cf87
		var _177_v81 _dafny.Array
		_ = _177_v81
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
		_ = _nw22
		_177_v81 = _nw22
		var _178_v82 _dafny.Sequence
		_ = _178_v82
		_178_v82 = _dafny.SeqOf(_79_v7)
		var _179_v83 T1
		_ = _179_v83
		var _nw23 *C1 = New_C1_()
		_ = _nw23
		_nw23.Ctor__((_dafny.IntOfInt64(-276)).Plus((_81_v9).F2()), (func() _dafny.Array {
			if _174_cf89 {
				return _177_v81
			}
			return _177_v81
		})(), _dafny.SeqOf(_72_v0, (_dafny.MultiSetOf((_dafny.SetOf((_178_v82).Select((Companion_Default___.SafeIndex((_81_v9).F2(), _dafny.IntOfUint32((_178_v82).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality())).Cardinality()), _80_v8)
		_179_v83 = _nw23
		_179_v83 = _179_v83
		_174_cf89 = !(!(_83_v10))
		var _180_v84 bool
		_ = _180_v84
		var _181_v85 bool
		_ = _181_v85
		var _out10 bool
		_ = _out10
		var _out11 bool
		_ = _out11
		_out10, _out11 = (_81_v9).M1(_86_v13, _71_globalState)
		_180_v84 = _out10
		_181_v85 = _out11
		_84_v11 = (_81_v9).F2()
	} else {
		var _182___mcc_h6 _dafny.Map = _source2.Get_().(D18_DC47).Cf85
		_ = _182___mcc_h6
		var _183_cf85 _dafny.Map = _182___mcc_h6
		_ = _183_cf85
		_83_v10 = _83_v10
		_83_v10 = !((_dafny.IntOfInt64(780)).Cmp((func() _dafny.Int {
			if _83_v10 {
				return _84_v11
			}
			return _dafny.IntOfInt64(770)
		})()) != 0)
		_84_v11 = (_86_v13).Times(Companion_Default___.SafeDivisionInt(_84_v11, _dafny.IntOfInt64(707)))
		var _184_v86 _dafny.Set
		_ = _184_v86
		_184_v86 = _dafny.SetOf(_72_v0)
		var _185_v87 _dafny.Set
		_ = _185_v87
		_185_v87 = _dafny.SetOf(_184_v86)
		var _186_v88 _dafny.Map
		_ = _186_v88
		_186_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v10, (Companion_D23_.Create_DC57_(_185_v87)).Dtor_cf100())
		_186_v88 = (_186_v88).Update(_83_v10, _185_v87)
	}
	var _187_v89 D20
	_ = _187_v89
	_187_v89 = Companion_D20_.Create_DC53_(!(_83_v10))
	if (_187_v89).Dtor_cf94() {
		_84_v11 = (_81_v9).F2()
		var _188_v90 _dafny.Array
		_ = _188_v90
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_1
		var _nw24 _dafny.Array
		_ = _nw24
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw24 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Int = (func(_189_v9 *C8) func(_dafny.Int) _dafny.Int {
				return func(_190_i5 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_190_i5, (_189_v9).F2())
				}
			})(_81_v9)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw24 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw24).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw24).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_188_v90 = _nw24
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_188_v90), 0))
		_ = _index9
		(_188_v90).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_74_v3, _dafny.UnicodeSeqOfUtf8Bytes("bnkl"))).Cardinality()), (_index9).Int())
		var _191_v91 _dafny.Set
		_ = _191_v91
		_191_v91 = _dafny.SetOf(_83_v10)
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_188_v90), 0))
		_ = _index10
		(_188_v90).ArraySet1(((_dafny.SetOf(_83_v10)).Difference(_191_v91)).Cardinality(), (_index10).Int())
		var _192_v92 T0
		_ = _192_v92
		var _nw25 *C2 = New_C2_()
		_ = _nw25
		_nw25.Ctor__()
		_192_v92 = _nw25
		var _193_v93 _dafny.MultiSet
		_ = _193_v93
		_193_v93 = _dafny.MultiSetOf(_dafny.IntOfInt64(394))
		var _194_v94 T0
		_ = _194_v94
		_194_v94 = _192_v92
		var _195_v95 _dafny.Map
		_ = _195_v95
		_195_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_193_v93).Equals(_193_v93)), (_194_v94))
		var _196_v96 _dafny.Map
		_ = _196_v96
		_196_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), _80_v8)
		var _197_v97 _dafny.Map
		_ = _197_v97
		_197_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_83_v10)), (_192_v92).Fm5(_72_v0, _83_v10, _196_v96, _83_v10, _71_globalState))
		_192_v92 = (func() T0 {
			if (_195_v95).Contains((func() bool {
				if (_197_v97).Contains(_83_v10) {
					return (_197_v97).Get(_83_v10).(bool)
				}
				return _83_v10
			})()) {
				return (_195_v95).Get((func() bool {
					if (_197_v97).Contains(_83_v10) {
						return (_197_v97).Get(_83_v10).(bool)
					}
					return _83_v10
				})()).(T0)
			}
			return _192_v92
		})()
		var _198_v98 _dafny.Array
		_ = _198_v98
		var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
		_ = _nw26
		_198_v98 = _nw26
		var _199_v99 _dafny.Array
		_ = _199_v99
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_2
		var _nw27 _dafny.Array
		_ = _nw27
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw27 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) D16 = (func(_200_v9 *C8, _201_v14 _dafny.Map, _202_v0 _dafny.Int) func(_dafny.Int) D16 {
				return func(_203_i6 _dafny.Int) D16 {
					return Companion_D16_.Create_DC42_((_200_v9).F2(), (_201_v14).Cardinality(), _202_v0)
				}
			})(_81_v9, _89_v14, _72_v0)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw27 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw27).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw27).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_199_v99 = _nw27
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_198_v98), 0))
		_ = _index11
		(_198_v98).ArraySet1(_199_v99, (_index11).Int())
		var _204_v100 D24
		_ = _204_v100
		_204_v100 = Companion_D24_.Create_DC59_(_199_v99)
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_198_v98), 0))
		_ = _index12
		(_198_v98).ArraySet1((_204_v100).Dtor_cf101(), (_index12).Int())
		var _205_v101 _dafny.CodePoint
		_ = _205_v101
		_205_v101 = _dafny.CodePoint('a')
		_205_v101 = _205_v101
	} else {
		var _206_v102 _dafny.Array
		_ = _206_v102
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
		_ = _nw28
		_206_v102 = _nw28
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))
		_ = _index13
		(_206_v102).ArraySet1(_83_v10, (_index13).Int())
		var _207_v103 D12
		_ = _207_v103
		_207_v103 = Companion_D12_.Create_DC28_(_83_v10, _dafny.IntOfUint32((_74_v3).Cardinality()))
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))
		_ = _index14
		(_206_v102).ArraySet1((Companion_Default___.SafeModuloInt((_207_v103).Dtor_cf47(), _dafny.IntOfInt64(860))).Cmp(_dafny.IntOfInt64(765)) == 0, (_index14).Int())
		_83_v10 = !(_83_v10) || (_83_v10)
		var _208_v105 _dafny.Set
		_ = _208_v105
		_208_v105 = _dafny.SetOf((_dafny.Zero).Minus(_72_v0), _84_v11, (func() _dafny.Map {
			var _coll30 = _dafny.NewMapBuilder()
			_ = _coll30
			for _iter30 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(163))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}((func(_209_v12 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_210_i7 _dafny.Int) _dafny.Map {
					return _209_v12
				}
			})(_85_v12)))).Elements()); ; {
				_compr_30, _ok30 := _iter30()
				if !_ok30 {
					break
				}
				var _211_v104 _dafny.Map
				_211_v104 = interface{}(_compr_30).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(163))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}((func(_212_v12 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_210_i7 _dafny.Int) _dafny.Map {
						return _212_v12
					}
				})(_85_v12))), _211_v104) {
					_coll30.Add(_211_v104, _86_v13)
				}
			}
			return _coll30.ToMap()
		}()).Cardinality(), _72_v0)
		var _213_v106 _dafny.Sequence
		_ = _213_v106
		_213_v106 = _dafny.SeqOf((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool))
		var _214_v107 _dafny.CodePoint
		_ = _214_v107
		_214_v107 = _dafny.CodePoint('m')
		var _215_v108 _dafny.Map
		_ = _215_v108
		_215_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), _214_v107)
		var _216_v109 _dafny.MultiSet
		_ = _216_v109
		_216_v109 = _dafny.MultiSetOf((_81_v9).F2(), (_215_v108).Cardinality(), _84_v11)
		var _nwElement0_2 bool = (_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool)
		_ = _nwElement0_2
		var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(29))
		_ = _nw29
		(_nw29).ArraySet1(_nwElement0_2, 0)
		(_nw29).ArraySet1(_83_v10, 1)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 2)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 3)
		(_nw29).ArraySet1(!((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool)), 4)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 5)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 6)
		(_nw29).ArraySet1(_83_v10, 7)
		(_nw29).ArraySet1(true, 8)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 9)
		(_nw29).ArraySet1((func() bool {
			if _83_v10 {
				return (_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool)
			}
			return false
		})(), 10)
		(_nw29).ArraySet1(_83_v10, 11)
		(_nw29).ArraySet1(_83_v10, 12)
		(_nw29).ArraySet1((_83_v10) == (_83_v10), 13)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 14)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 15)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 16)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 17)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 18)
		(_nw29).ArraySet1((_86_v13).Cmp(_72_v0) >= 0, 19)
		(_nw29).ArraySet1((_208_v105).IsSubsetOf(_dafny.SetOf(_72_v0)), 20)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 21)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 22)
		(_nw29).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_74_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-976))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg33 _dafny.Int) interface{} {
				return coer33(arg33)
			}
		}(func(_217_i8 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		}))), 23)
		(_nw29).ArraySet1((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), 24)
		(_nw29).ArraySet1(!((_213_v106).Select((Companion_Default___.SafeIndex((_216_v109).Cardinality(), _dafny.IntOfUint32((_213_v106).Cardinality()))).Uint32()).(bool)), 25)
		(_nw29).ArraySet1(((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool)) == (true), 26)
		(_nw29).ArraySet1(_83_v10, 27)
		(_nw29).ArraySet1(((_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool)) && (_83_v10), 28)
		_206_v102 = _nw29
		var _218_v110 _dafny.Map
		_ = _218_v110
		_218_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v10, (func() bool {
			if _83_v10 {
				return (_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool)
			}
			return (_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool)
		})())
		_218_v110 = (_218_v110).Update((_72_v0).Cmp(_86_v13) == 0, _83_v10)
		var _219_v111 _dafny.Map
		_ = _219_v111
		_219_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v13, (_81_v9).Fm6(_72_v0, _dafny.IntOfUint32((_79_v7).Cardinality()), _71_globalState))
		_72_v0 = (func() _dafny.Int {
			if (_219_v111).Contains(_84_v11) {
				return (_219_v111).Get(_84_v11).(_dafny.Int)
			}
			return _72_v0
		})()
	}
	var _220_v112 _dafny.Map
	_ = _220_v112
	_220_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v0, (_81_v9).F2())
	_72_v0 = ((_81_v9).F2()).Times((_dafny.Zero).Minus((_220_v112).Cardinality()))
	var _221_i9 _dafny.Int
	_ = _221_i9
	_221_i9 = _dafny.Zero
	{
		var _pat_let_tv3 = _83_v10
		_ = _pat_let_tv3
		var _pat_let_tv4 = _83_v10
		_ = _pat_let_tv4
		var _pat_let_tv5 = _83_v10
		_ = _pat_let_tv5
		for !(func(_source3 D18) bool {
			if _source3.Is_DC48() {
				return _pat_let_tv3
			} else if _source3.Is_DC49() {
				var _227___mcc_h7 _dafny.Set = _source3.Get_().(D18_DC49).Cf86
				_ = _227___mcc_h7
				var _228_cf86 _dafny.Set = _227___mcc_h7
				_ = _228_cf86
				return _pat_let_tv4
			} else if _source3.Is_DC50() {
				var _229___mcc_h8 bool = _source3.Get_().(D18_DC50).Cf87
				_ = _229___mcc_h8
				var _230___mcc_h9 _dafny.Array = _source3.Get_().(D18_DC50).Cf88
				_ = _230___mcc_h9
				var _231___mcc_h10 bool = _source3.Get_().(D18_DC50).Cf89
				_ = _231___mcc_h10
				var _232___mcc_h11 bool = _source3.Get_().(D18_DC50).Cf90
				_ = _232___mcc_h11
				var _233___mcc_h12 _dafny.Sequence = _source3.Get_().(D18_DC50).Cf91
				_ = _233___mcc_h12
				var _234_cf91 _dafny.Sequence = _233___mcc_h12
				_ = _234_cf91
				var _235_cf90 bool = _232___mcc_h11
				_ = _235_cf90
				var _236_cf89 bool = _231___mcc_h10
				_ = _236_cf89
				var _237_cf88 _dafny.Array = _230___mcc_h9
				_ = _237_cf88
				var _238_cf87 bool = _229___mcc_h8
				_ = _238_cf87
				return _238_cf87
			} else {
				var _239___mcc_h13 _dafny.Map = _source3.Get_().(D18_DC47).Cf85
				_ = _239___mcc_h13
				var _240_cf85 _dafny.Map = _239___mcc_h13
				_ = _240_cf85
				return _pat_let_tv5
			}
		}(_91_v16)) {
			{
				if (_221_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_221_i9 = (_221_i9).Plus(_dafny.One)
				var _222_v113 _dafny.Map
				_ = _222_v113
				_222_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v10, _83_v10)
				var _rhs14 _dafny.Map = Companion_Default___.Fm47(((_81_v9).F2()).Times(_72_v0), ((_81_v9).F2()).Cmp(_72_v0) <= 0, _71_globalState)
				_ = _rhs14
				var _rhs15 bool = (_84_v11).Cmp(_72_v0) > 0
				_ = _rhs15
				_222_v113 = _rhs14
				_83_v10 = _rhs15
				var _223_v114 _dafny.Array
				_ = _223_v114
				var _nw30 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
				_ = _nw30
				_223_v114 = _nw30
				var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_223_v114), 0))
				_ = _index15
				(_223_v114).ArraySet1(_83_v10, (_index15).Int())
				var _224_v115 T0
				_ = _224_v115
				var _nw31 *C2 = New_C2_()
				_ = _nw31
				_nw31.Ctor__()
				_224_v115 = _nw31
				var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_223_v114), 0))
				_ = _index16
				var _rhs16 *C4 = _73_v1
				_ = _rhs16
				var _rhs17 bool = !(!(_83_v10))
				_ = _rhs17
				var _rhs18 _dafny.Array = _223_v114
				_ = _rhs18
				var _rhs19 _dafny.Int = (_86_v13).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_224_v115)).Cardinality()))
				_ = _rhs19
				var _lhs6 _dafny.Array = _223_v114
				_ = _lhs6
				var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_223_v114), 0))
				_ = _lhs7
				_73_v1 = _rhs16
				(_lhs6).ArraySet1(_rhs17, (_lhs7).Int())
				_223_v114 = _rhs18
				_72_v0 = _rhs19
				var _225_v116 *C3
				_ = _225_v116
				var _nw32 *C3 = New_C3_()
				_ = _nw32
				_nw32.Ctor__()
				_225_v116 = _nw32
				var _226_v117 _dafny.Sequence
				_ = _226_v117
				_226_v117 = _dafny.SeqOf(_225_v116)
				_225_v116 = (func() *C3 {
					if _83_v10 {
						return _225_v116
					}
					return (_226_v117).Select((Companion_Default___.SafeIndex((_81_v9).F2(), _dafny.IntOfUint32((_226_v117).Cardinality()))).Uint32()).(*C3)
				})()
				_83_v10 = (_223_v114).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_223_v114), 0))).Int()).(bool)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _241_v118 *C4
	_ = _241_v118
	var _nw33 *C4 = New_C4_()
	_ = _nw33
	_nw33.Ctor__()
	_241_v118 = _nw33
	var _242_v119 _dafny.Array
	_ = _242_v119
	var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
	_ = _nw34
	_242_v119 = _nw34
	var _243_v120 _dafny.Map
	_ = _243_v120
	_243_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v119, _242_v119)
	var _244_v121 _dafny.Map
	_ = _244_v121
	_244_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v10, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v119, _242_v119))
	var _245_v122 _dafny.Array
	_ = _245_v122
	var _nwElement0_3 _dafny.Map = (func() _dafny.Map {
		if _83_v10 {
			return _243_v120
		}
		return _243_v120
	})()
	_ = _nwElement0_3
	var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(16))
	_ = _nw35
	(_nw35).ArraySet1(_nwElement0_3, 0)
	(_nw35).ArraySet1((_243_v120).Merge(_243_v120), 1)
	(_nw35).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v119, _242_v119)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v119, _242_v119)), 2)
	(_nw35).ArraySet1((_243_v120).Merge(_243_v120), 3)
	(_nw35).ArraySet1((_243_v120).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v119, _242_v119)), 4)
	(_nw35).ArraySet1((_243_v120).Update(_242_v119, _242_v119), 5)
	(_nw35).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v119, _242_v119)).Merge(_243_v120), 6)
	(_nw35).ArraySet1((func() _dafny.Map {
		if (_244_v121).Contains(_83_v10) {
			return (_244_v121).Get(_83_v10).(_dafny.Map)
		}
		return _243_v120
	})(), 7)
	(_nw35).ArraySet1(_243_v120, 8)
	(_nw35).ArraySet1(_243_v120, 9)
	(_nw35).ArraySet1((_243_v120).Update(_242_v119, _242_v119), 10)
	(_nw35).ArraySet1(_243_v120, 11)
	(_nw35).ArraySet1(_243_v120, 12)
	(_nw35).ArraySet1(_243_v120, 13)
	(_nw35).ArraySet1(_243_v120, 14)
	(_nw35).ArraySet1(_243_v120, 15)
	_245_v122 = _nw35
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_245_v122), 0))
	_ = _index17
	(_245_v122).ArraySet1(_243_v120, (_index17).Int())
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_245_v122), 0))
	_ = _index18
	(_245_v122).ArraySet1(_243_v120, (_index18).Int())
	_83_v10 = false
	var _246_v123 _dafny.Set
	_ = _246_v123
	_246_v123 = _dafny.SetOf(false, _83_v10, true, _83_v10, true)
	var _247_v124 _dafny.Map
	_ = _247_v124
	_247_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v10, _246_v123)
	var _248_v125 _dafny.Map
	_ = _248_v125
	_248_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v0, _247_v124)
	var _249_v126 _dafny.Sequence
	_ = _249_v126
	_249_v126 = _dafny.SeqOf(false, _83_v10, false)
	var _250_v127 _dafny.Sequence
	_ = _250_v127
	_250_v127 = _dafny.SeqOf(_247_v124, _247_v124, _247_v124)
	var _251_v128 _dafny.Array
	_ = _251_v128
	var _nwElement0_4 _dafny.Map = (func() _dafny.Map {
		if (_248_v125).Contains(_dafny.IntOfUint32((_249_v126).Cardinality())) {
			return (_248_v125).Get(_dafny.IntOfUint32((_249_v126).Cardinality())).(_dafny.Map)
		}
		return _247_v124
	})()
	_ = _nwElement0_4
	var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(28))
	_ = _nw36
	(_nw36).ArraySet1(_nwElement0_4, 0)
	(_nw36).ArraySet1(_247_v124, 1)
	(_nw36).ArraySet1(_247_v124, 2)
	(_nw36).ArraySet1(_247_v124, 3)
	(_nw36).ArraySet1(_247_v124, 4)
	(_nw36).ArraySet1(_247_v124, 5)
	(_nw36).ArraySet1((_247_v124).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _246_v123)), 6)
	(_nw36).ArraySet1((_250_v127).Select((Companion_Default___.SafeIndex(_72_v0, _dafny.IntOfUint32((_250_v127).Cardinality()))).Uint32()).(_dafny.Map), 7)
	(_nw36).ArraySet1(_247_v124, 8)
	(_nw36).ArraySet1(_247_v124, 9)
	(_nw36).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _246_v123)).Merge(_247_v124), 10)
	(_nw36).ArraySet1(_247_v124, 11)
	(_nw36).ArraySet1(_247_v124, 12)
	(_nw36).ArraySet1(_247_v124, 13)
	(_nw36).ArraySet1(_247_v124, 14)
	(_nw36).ArraySet1(_247_v124, 15)
	(_nw36).ArraySet1(_247_v124, 16)
	(_nw36).ArraySet1(_247_v124, 17)
	(_nw36).ArraySet1((func() _dafny.Map {
		if _83_v10 {
			return _247_v124
		}
		return _247_v124
	})(), 18)
	(_nw36).ArraySet1((_247_v124).Merge(_247_v124), 19)
	(_nw36).ArraySet1(_247_v124, 20)
	(_nw36).ArraySet1((func() _dafny.Map {
		if (_248_v125).Contains((_81_v9).F2()) {
			return (_248_v125).Get((_81_v9).F2()).(_dafny.Map)
		}
		return _247_v124
	})(), 21)
	(_nw36).ArraySet1((_247_v124).Update(true, _246_v123), 22)
	(_nw36).ArraySet1(_247_v124, 23)
	(_nw36).ArraySet1(_247_v124, 24)
	(_nw36).ArraySet1(_247_v124, 25)
	(_nw36).ArraySet1(_247_v124, 26)
	(_nw36).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _246_v123)).Merge(_247_v124), 27)
	_251_v128 = _nw36
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_251_v128), 0))
	_ = _index19
	(_251_v128).ArraySet1(_247_v124, (_index19).Int())
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_251_v128), 0))
	_ = _index20
	(_251_v128).ArraySet1(_247_v124, (_index20).Int())
	_83_v10 = _83_v10
	var _252_v129 _dafny.CodePoint
	_ = _252_v129
	_252_v129 = _dafny.CodePoint('a')
	var _253_v130 _dafny.Array
	_ = _253_v130
	var _nwElement0_5 _dafny.CodePoint = _252_v129
	_ = _nwElement0_5
	var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(3))
	_ = _nw37
	(_nw37).ArraySet1CodePoint(_nwElement0_5, 0)
	(_nw37).ArraySet1CodePoint(_252_v129, 1)
	(_nw37).ArraySet1CodePoint(_252_v129, 2)
	_253_v130 = _nw37
	var _254_v131 _dafny.Sequence
	_ = _254_v131
	_254_v131 = _dafny.SeqOf(_253_v130)
	var _255_v132 _dafny.Map
	_ = _255_v132
	_255_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_187_v89).Dtor_cf94(), (_254_v131).Select((Companion_Default___.SafeIndex(_84_v11, _dafny.IntOfUint32((_254_v131).Cardinality()))).Uint32()).(_dafny.Array))
	_255_v132 = (_255_v132).Update(_83_v10, _253_v130)
	if (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("vdgmtsaje"), _dafny.UnicodeSeqOfUtf8Bytes("lhgie"))).Equals(_75_v4) {
		var _256_v133 *C0
		_ = _256_v133
		var _nw38 *C0 = New_C0_()
		_ = _nw38
		_nw38.Ctor__(_74_v3)
		_256_v133 = _nw38
		var _257_v134 _dafny.Map
		_ = _257_v134
		_257_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_256_v133, _83_v10)
		_83_v10 = (func() bool {
			if (func() bool {
				if (_257_v134).Contains(_256_v133) {
					return (_257_v134).Get(_256_v133).(bool)
				}
				return _83_v10
			})() {
				return (func() bool {
					if (_249_v126).Select((Companion_Default___.SafeIndex(_86_v13, _dafny.IntOfUint32((_249_v126).Cardinality()))).Uint32()).(bool) {
						return _83_v10
					}
					return _83_v10
				})()
			}
			return _83_v10
		})()
		var _258_v135 D24
		_ = _258_v135
		_258_v135 = Companion_D24_.Create_DC62_((_81_v9).F2())
		var _pat_let_tv6 = _84_v11
		_ = _pat_let_tv6
		var _259_v136 _dafny.Map
		_ = _259_v136
		_259_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let6_0 D24) D24 {
			return func(_260_dt__update__tmp_h3 D24) D24 {
				return func(_pat_let7_0 _dafny.Int) D24 {
					return func(_261_dt__update_hcf104_h0 _dafny.Int) D24 {
						return Companion_D24_.Create_DC62_(_261_dt__update_hcf104_h0)
					}(_pat_let7_0)
				}(_pat_let_tv6)
			}(_pat_let6_0)
		}(_258_v135), _242_v119)
		var _262_v137 _dafny.Array
		_ = _262_v137
		var _nwElement0_6 _dafny.Array = _242_v119
		_ = _nwElement0_6
		var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(27))
		_ = _nw39
		(_nw39).ArraySet1(_nwElement0_6, 0)
		(_nw39).ArraySet1(_242_v119, 1)
		(_nw39).ArraySet1(_242_v119, 2)
		(_nw39).ArraySet1(_242_v119, 3)
		(_nw39).ArraySet1(_242_v119, 4)
		(_nw39).ArraySet1(_242_v119, 5)
		(_nw39).ArraySet1(_242_v119, 6)
		(_nw39).ArraySet1(_242_v119, 7)
		(_nw39).ArraySet1(_242_v119, 8)
		(_nw39).ArraySet1(_242_v119, 9)
		(_nw39).ArraySet1(_242_v119, 10)
		(_nw39).ArraySet1(_242_v119, 11)
		(_nw39).ArraySet1(_242_v119, 12)
		(_nw39).ArraySet1((func() _dafny.Array {
			if (_259_v136).Contains(_258_v135) {
				return (_259_v136).Get(_258_v135).(_dafny.Array)
			}
			return _242_v119
		})(), 13)
		(_nw39).ArraySet1(_242_v119, 14)
		(_nw39).ArraySet1(_242_v119, 15)
		(_nw39).ArraySet1(_242_v119, 16)
		(_nw39).ArraySet1((func() _dafny.Array {
			if true {
				return _242_v119
			}
			return _242_v119
		})(), 17)
		(_nw39).ArraySet1(_242_v119, 18)
		(_nw39).ArraySet1(_242_v119, 19)
		(_nw39).ArraySet1(_242_v119, 20)
		(_nw39).ArraySet1((func() _dafny.Array {
			if _83_v10 {
				return _242_v119
			}
			return _242_v119
		})(), 21)
		(_nw39).ArraySet1(_242_v119, 22)
		(_nw39).ArraySet1(_242_v119, 23)
		(_nw39).ArraySet1(_242_v119, 24)
		(_nw39).ArraySet1(_242_v119, 25)
		(_nw39).ArraySet1(_242_v119, 26)
		_262_v137 = _nw39
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_262_v137), 0))
		_ = _index21
		(_262_v137).ArraySet1(_242_v119, (_index21).Int())
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_262_v137), 0))
		_ = _index22
		(_262_v137).ArraySet1(_242_v119, (_index22).Int())
		var _263_v138 D15
		_ = _263_v138
		_263_v138 = Companion_D15_.Create_DC38_(_252_v129)
		var _source4 D15 = _263_v138
		_ = _source4
		if _source4.Is_DC38() {
			var _264___mcc_h14 _dafny.CodePoint = _source4.Get_().(D15_DC38).Cf71
			_ = _264___mcc_h14
			var _265_cf71 _dafny.CodePoint = _264___mcc_h14
			_ = _265_cf71
			var _266_v140 _dafny.Set
			_ = _266_v140
			_266_v140 = _dafny.SetOf(_dafny.IntOfUint32((_74_v3).Cardinality()), _84_v11)
			var _267_v141 T0
			_ = _267_v141
			var _nw40 *C8 = New_C8_()
			_ = _nw40
			_nw40.Ctor__(((func() _dafny.Set {
				if _83_v10 {
					return func() _dafny.Set {
						var _coll31 = _dafny.NewBuilder()
						_ = _coll31
						for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(109), _dafny.IntOfInt64(874))); ; {
							_compr_31, _ok31 := _iter31()
							if !_ok31 {
								break
							}
							var _268_v139 _dafny.Int
							_268_v139 = interface{}(_compr_31).(_dafny.Int)
							if ((_dafny.IntOfInt64(109)).Cmp(_268_v139) <= 0) && ((_268_v139).Cmp(_dafny.IntOfInt64(874)) < 0) {
								_coll31.Add((_268_v139).Plus((_81_v9).F2()))
							}
						}
						return _coll31.ToSet()
					}()
				}
				return _266_v140
			})()).Cardinality(), Companion_D0_.Create_DC0_(_76_v5, (_81_v9).F2()), _dafny.Companion_Sequence_.Concatenate(_79_v7, _79_v7), _80_v8)
			_267_v141 = _nw40
			var _nw41 *C2 = New_C2_()
			_ = _nw41
			_nw41.Ctor__()
			_267_v141 = _nw41
			var _269_v142 _dafny.Map
			_ = _269_v142
			var _out12 _dafny.Map
			_ = _out12
			_out12 = (_81_v9).M2(_84_v11, _71_globalState)
			_269_v142 = _out12
			var _270_v143 D16
			_ = _270_v143
			_270_v143 = Companion_D16_.Create_DC42_((_81_v9).F2(), (_81_v9).F2(), _86_v13)
			var _271_v145 _dafny.MultiSet
			_ = _271_v145
			_271_v145 = _dafny.MultiSetOf(_83_v10)
			var _272_v146 _dafny.Sequence
			_ = _272_v146
			_272_v146 = _dafny.SeqOf(_271_v145, _271_v145, _271_v145)
			var _273_v147 _dafny.Sequence
			_ = _273_v147
			_273_v147 = _dafny.SeqOf(_256_v133.F8, _256_v133.F8, (_256_v133).Fm15(_dafny.SeqOf(_86_v13, (func() _dafny.Map {
				var _coll32 = _dafny.NewMapBuilder()
				_ = _coll32
				for _iter32 := _dafny.Iterate((_dafny.MultiSetFromSeq(_272_v146)).Elements()); ; {
					_compr_32, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _274_v144 _dafny.MultiSet
					_274_v144 = interface{}(_compr_32).(_dafny.MultiSet)
					if (_dafny.MultiSetFromSeq(_272_v146)).Contains(_274_v144) {
						_coll32.Add(_274_v144, _265_cf71)
					}
				}
				return _coll32.ToMap()
			}()).Cardinality(), (_81_v9).F2(), _84_v11), _83_v10, _71_globalState), _256_v133.F8, _dafny.UnicodeSeqOfUtf8Bytes("s"))
			var _275_v148 _dafny.Array
			_ = _275_v148
			var _nwElement0_7 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(224))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}((func(_276_v3 _dafny.Sequence, _277_v10 bool) func(_dafny.Int) _dafny.Sequence {
				return func(_278_i10 _dafny.Int) _dafny.Sequence {
					return (Companion_D15_.Create_DC39_(_276_v3, Companion_D9_.Create_DC21_(_dafny.MultiSetOf(false, _277_v10)))).Dtor_cf72()
				}
			})(_74_v3, _83_v10))), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-827))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}((func(_279_v133 *C0) func(_dafny.Int) _dafny.Sequence {
				return func(_280_i11 _dafny.Int) _dafny.Sequence {
					return _279_v133.F8
				}
			})(_256_v133))), (Companion_Default___.SafeIndex((_270_v143).Dtor_cf77(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-827))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}((func(_281_v133 *C0) func(_dafny.Int) _dafny.Sequence {
				return func(_282_i11 _dafny.Int) _dafny.Sequence {
					return _281_v133.F8
				}
			})(_256_v133)))).Cardinality()))).Uint32(), _74_v3))
			_ = _nwElement0_7
			var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(7))
			_ = _nw42
			(_nw42).ArraySet1(_nwElement0_7, 0)
			(_nw42).ArraySet1(_273_v147, 1)
			(_nw42).ArraySet1(_273_v147, 2)
			(_nw42).ArraySet1(Companion_Default___.Fm72(_71_globalState), 3)
			(_nw42).ArraySet1(_273_v147, 4)
			(_nw42).ArraySet1(_dafny.SeqOf(_74_v3, _256_v133.F8, _dafny.UnicodeSeqOfUtf8Bytes("ea"), _256_v133.F8, _256_v133.F8), 5)
			(_nw42).ArraySet1(_dafny.Companion_Sequence_.Update(_273_v147, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.IntOfUint32((_273_v147).Cardinality()))).Uint32(), _256_v133.F8), 6)
			_275_v148 = _nw42
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(350), _dafny.ArrayLen((_275_v148), 0))
			_ = _index23
			(_275_v148).ArraySet1(_dafny.Companion_Sequence_.Update(_273_v147, (Companion_Default___.SafeIndex(_86_v13, _dafny.IntOfUint32((_273_v147).Cardinality()))).Uint32(), _74_v3), (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(350), _dafny.ArrayLen((_275_v148), 0))
			_ = _index24
			(_275_v148).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_273_v147, _dafny.Companion_Sequence_.Concatenate(_273_v147, _273_v147)), (_index24).Int())
			_265_cf71 = _265_cf71
		} else if _source4.Is_DC39() {
			var _283___mcc_h15 _dafny.Sequence = _source4.Get_().(D15_DC39).Cf72
			_ = _283___mcc_h15
			var _284___mcc_h16 D9 = _source4.Get_().(D15_DC39).Cf73
			_ = _284___mcc_h16
			var _285_cf73 D9 = _284___mcc_h16
			_ = _285_cf73
			var _286_cf72 _dafny.Sequence = _283___mcc_h15
			_ = _286_cf72
			var _287_v149 _dafny.Map
			_ = _287_v149
			_287_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_75_v4, ((_220_v112).Cardinality()).Cmp((_81_v9).F2()) >= 0)
			_287_v149 = (_287_v149).Update(_dafny.MultiSetOf(_256_v133.F8, _256_v133.F8, _256_v133.F8), _83_v10)
			var _288_v150 T0
			_ = _288_v150
			var _nw43 *C8 = New_C8_()
			_ = _nw43
			_nw43.Ctor__((_81_v9).F2(), _78_v6, _79_v7, _80_v8)
			_288_v150 = _nw43
			var _289_v151 _dafny.MultiSet
			_ = _289_v151
			_289_v151 = _dafny.MultiSetOf(_288_v150)
			var _290_v152 _dafny.Map
			_ = _290_v152
			_290_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm32((_dafny.Zero).Minus((func() _dafny.Int {
				if (_289_v151).Contains(_288_v150) {
					return (_289_v151).Multiplicity(_288_v150)
				}
				return (_81_v9).F2()
			})()), _71_globalState), _256_v133.F8)
			_290_v152 = (_290_v152).Update(_72_v0, _dafny.Companion_Sequence_.Concatenate(_256_v133.F8, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(915))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}((func(_291_v129 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_292_i12 _dafny.Int) _dafny.CodePoint {
					return _291_v129
				}
			})(_252_v129)))))
			var _293_v153 *C9
			_ = _293_v153
			var _nw44 *C9 = New_C9_()
			_ = _nw44
			_nw44.Ctor__()
			_293_v153 = _nw44
			var _294_v154 _dafny.Array
			_ = _294_v154
			_294_v154 = _dafny.ArrayCastTo((_262_v137).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_262_v137), 0))).Int()))
			var _295_v155 _dafny.MultiSet
			_ = _295_v155
			_295_v155 = _dafny.MultiSetOf(_242_v119, _294_v154, _294_v154)
			_295_v155 = _295_v155
		} else if _source4.Is_DC37() {
			var _296___mcc_h17 _dafny.Set = _source4.Get_().(D15_DC37).Cf70
			_ = _296___mcc_h17
			var _297_cf70 _dafny.Set = _296___mcc_h17
			_ = _297_cf70
			var _298_v156 _dafny.Array
			_ = _298_v156
			var _299_v157 _dafny.Array
			_ = _299_v157
			var _300_v158 bool
			_ = _300_v158
			var _out13 _dafny.Array
			_ = _out13
			var _out14 _dafny.Array
			_ = _out14
			var _out15 bool
			_ = _out15
			_out13, _out14, _out15 = (_73_v1).M9(_71_globalState)
			_298_v156 = _out13
			_299_v157 = _out14
			_300_v158 = _out15
			_84_v11 = ((_81_v9).F2()).Times(Companion_Default___.Fm20(_300_v158, _71_globalState))
			var _301_v159 _dafny.Set
			_ = _301_v159
			_301_v159 = _dafny.SetOf((_81_v9).F2())
			var _302_v160 _dafny.MultiSet
			_ = _302_v160
			_302_v160 = _dafny.MultiSetOf(_300_v158)
			var _303_v161 _dafny.Array
			_ = _303_v161
			var _nwElement0_8 bool = _300_v158
			_ = _nwElement0_8
			var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(22))
			_ = _nw45
			(_nw45).ArraySet1(_nwElement0_8, 0)
			(_nw45).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_256_v133.F8, _dafny.UnicodeSeqOfUtf8Bytes("jtsgpyod")), 1)
			(_nw45).ArraySet1(_300_v158, 2)
			(_nw45).ArraySet1((_301_v159).Equals(_301_v159), 3)
			(_nw45).ArraySet1(_83_v10, 4)
			(_nw45).ArraySet1(_300_v158, 5)
			(_nw45).ArraySet1(!(!(_302_v160).Contains(!(_300_v158))), 6)
			(_nw45).ArraySet1(_83_v10, 7)
			(_nw45).ArraySet1(_300_v158, 8)
			(_nw45).ArraySet1(_300_v158, 9)
			(_nw45).ArraySet1(_83_v10, 10)
			(_nw45).ArraySet1(_83_v10, 11)
			(_nw45).ArraySet1(!(_300_v158) || (_300_v158), 12)
			(_nw45).ArraySet1(_83_v10, 13)
			(_nw45).ArraySet1(_83_v10, 14)
			(_nw45).ArraySet1((_249_v126).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(_84_v11)).Cardinality(), _dafny.IntOfUint32((_249_v126).Cardinality()))).Uint32()).(bool), 15)
			(_nw45).ArraySet1(!(_300_v158), 16)
			(_nw45).ArraySet1(_300_v158, 17)
			(_nw45).ArraySet1(_83_v10, 18)
			(_nw45).ArraySet1(_dafny.Companion_Sequence_.Contains(_256_v133.F8, _252_v129), 19)
			(_nw45).ArraySet1(_300_v158, 20)
			(_nw45).ArraySet1(!(_83_v10) || (_83_v10), 21)
			_303_v161 = _nw45
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_303_v161), 0))
			_ = _index25
			(_303_v161).ArraySet1(!(_83_v10) || (_83_v10), (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_303_v161), 0))
			_ = _index26
			var _rhs20 bool = (_86_v13).Cmp((_84_v11).Plus((_81_v9).F2())) <= 0
			_ = _rhs20
			var _rhs21 _dafny.Int = (_81_v9).F2()
			_ = _rhs21
			var _rhs22 bool = !(_83_v10)
			_ = _rhs22
			var _lhs8 _dafny.Array = _303_v161
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_303_v161), 0))
			_ = _lhs9
			(_lhs8).ArraySet1(_rhs20, (_lhs9).Int())
			_84_v11 = _rhs21
			_300_v158 = _rhs22
			var _304_v162 _dafny.MultiSet
			_ = _304_v162
			_304_v162 = _dafny.MultiSetOf(_86_v13)
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_303_v161), 0))
			_ = _index27
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_303_v161), 0))
			_ = _index28
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_303_v161), 0))
			_ = _index29
			var _rhs23 _dafny.Sequence = _249_v126
			_ = _rhs23
			var _rhs24 bool = true
			_ = _rhs24
			var _rhs25 bool = (_304_v162).IsSubsetOf((_304_v162).Union(Companion_Default___.Fm24(_83_v10, _71_globalState)))
			_ = _rhs25
			var _rhs26 bool = (_303_v161).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_303_v161), 0))).Int()).(bool)
			_ = _rhs26
			var _lhs10 _dafny.Array = _303_v161
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_303_v161), 0))
			_ = _lhs11
			var _lhs12 _dafny.Array = _303_v161
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_303_v161), 0))
			_ = _lhs13
			var _lhs14 _dafny.Array = _303_v161
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_303_v161), 0))
			_ = _lhs15
			_249_v126 = _rhs23
			(_lhs10).ArraySet1(_rhs24, (_lhs11).Int())
			(_lhs12).ArraySet1(_rhs25, (_lhs13).Int())
			(_lhs14).ArraySet1(_rhs26, (_lhs15).Int())
		} else {
			var _305___mcc_h18 D15 = _source4.Get_().(D15_DC40).Cf74
			_ = _305___mcc_h18
			var _306_cf74 D15 = _305___mcc_h18
			_ = _306_cf74
			_86_v13 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_241_v118)).Cardinality()), (_81_v9).F2()), (_81_v9).F2())
			var _307_v163 *C4
			_ = _307_v163
			var _nw46 *C4 = New_C4_()
			_ = _nw46
			_nw46.Ctor__()
			_307_v163 = _nw46
			var _308_v164 _dafny.Map
			_ = _308_v164
			var _out16 _dafny.Map
			_ = _out16
			_out16 = (_241_v118).M2((_81_v9).F2(), _71_globalState)
			_308_v164 = _out16
			var _309_v165 _dafny.Map
			_ = _309_v165
			_309_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_81_v9).F2(), _83_v10)
			var _pat_let_tv7 = _81_v9
			_ = _pat_let_tv7
			var _310_v166 _dafny.Map
			_ = _310_v166
			_310_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_252_v129, func(_pat_let8_0 D0) D0 {
				return func(_311_dt__update__tmp_h4 D0) D0 {
					return func(_pat_let9_0 _dafny.Int) D0 {
						return func(_312_dt__update_hcf2_h1 _dafny.Int) D0 {
							return Companion_D0_.Create_DC1_(_312_dt__update_hcf2_h1)
						}(_pat_let9_0)
					}((_pat_let_tv7).F2())
				}(_pat_let8_0)
			}(_80_v8))
			_83_v10 = (_81_v9).Fm5(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(200), _84_v11), (_dafny.IntOfInt64(993)).Cmp((_309_v165).Cardinality()) < 0, _310_v166, _83_v10, _71_globalState)
		}
		_83_v10 = _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("sdjic"), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("q"), _256_v133.F8), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_249_v126).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("q"), _256_v133.F8)).Cardinality()))).Uint32(), _252_v129))
		var _313_v167 _dafny.Map
		_ = _313_v167
		_313_v167 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_81_v9).F2(), (_81_v9).F2()), _242_v119)
		_313_v167 = _313_v167
	} else {
		var _314_v168 _dafny.Array
		_ = _314_v168
		var _nwElement0_9 _dafny.Array = _242_v119
		_ = _nwElement0_9
		var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(14))
		_ = _nw47
		(_nw47).ArraySet1(_nwElement0_9, 0)
		(_nw47).ArraySet1(_242_v119, 1)
		(_nw47).ArraySet1(_242_v119, 2)
		(_nw47).ArraySet1(_242_v119, 3)
		(_nw47).ArraySet1(_242_v119, 4)
		(_nw47).ArraySet1(_242_v119, 5)
		(_nw47).ArraySet1(_242_v119, 6)
		(_nw47).ArraySet1(_242_v119, 7)
		(_nw47).ArraySet1(_242_v119, 8)
		(_nw47).ArraySet1(_242_v119, 9)
		(_nw47).ArraySet1(_242_v119, 10)
		(_nw47).ArraySet1(_242_v119, 11)
		(_nw47).ArraySet1(_242_v119, 12)
		(_nw47).ArraySet1(_242_v119, 13)
		_314_v168 = _nw47
		var _315_v169 _dafny.MultiSet
		_ = _315_v169
		_315_v169 = _dafny.MultiSetOf(_83_v10, _83_v10, _83_v10)
		var _316_v170 _dafny.Map
		_ = _316_v170
		_316_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_315_v169).Contains(_83_v10) {
				return (_315_v169).Multiplicity(_83_v10)
			}
			return _dafny.IntOfInt64(202)
		})(), _dafny.MultiSetFromSeq(_249_v126))
		var _317_v171 D25
		_ = _317_v171
		_317_v171 = Companion_D25_.Create_DC63_(_316_v170)
		var _318_v172 _dafny.Map
		_ = _318_v172
		_318_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_314_v168, ((_317_v171).Dtor_cf105()).Update(_dafny.IntOfInt64(879), (func() _dafny.MultiSet {
			if (_316_v170).Contains(_72_v0) {
				return (_316_v170).Get(_72_v0).(_dafny.MultiSet)
			}
			return _315_v169
		})()))
		_318_v172 = (_318_v172).Update(_314_v168, _316_v170)
		var _319_v173 *C4
		_ = _319_v173
		var _nw48 *C4 = New_C4_()
		_ = _nw48
		_nw48.Ctor__()
		_319_v173 = _nw48
		var _320_v174 D24
		_ = _320_v174
		_320_v174 = Companion_D24_.Create_DC61_(_84_v11, _dafny.MultiSetOf(_242_v119))
		_320_v174 = _320_v174
		var _321_v175 *C5
		_ = _321_v175
		var _nw49 *C5 = New_C5_()
		_ = _nw49
		_nw49.Ctor__(_242_v119, _dafny.SeqOf(_86_v13), _80_v8)
		_321_v175 = _nw49
		_83_v10 = !_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("xg"), (_74_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_74_v3).Cardinality()), _dafny.IntOfUint32((_74_v3).Cardinality()))).Uint32()).(_dafny.CodePoint))
	}
	var _322_i13 _dafny.Int
	_ = _322_i13
	_322_i13 = _dafny.Zero
	{
		for (func() bool {
			if _83_v10 {
				return _83_v10
			}
			return _83_v10
		})() {
			{
				if (_322_i13).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_322_i13 = (_322_i13).Plus(_dafny.One)
				var _323_v176 D16
				_ = _323_v176
				_323_v176 = Companion_D16_.Create_DC44_(_84_v11)
				var _324_v177 _dafny.Set
				_ = _324_v177
				_324_v177 = _dafny.SetOf(Companion_D16_.Create_DC44_((_81_v9).F2()), _323_v176, _323_v176, _323_v176)
				var _325_v178 D11
				_ = _325_v178
				_325_v178 = Companion_D11_.Create_DC25_(_252_v129, (_81_v9).F2(), _83_v10)
				var _rhs27 bool = ((_324_v177).Intersection(_324_v177)).IsDisjointFrom(_324_v177)
				_ = _rhs27
				var _rhs28 bool = (_325_v178).Dtor_cf43()
				_ = _rhs28
				var _rhs29 _dafny.Int = _dafny.IntOfInt64(559)
				_ = _rhs29
				var _rhs30 bool = _83_v10
				_ = _rhs30
				_83_v10 = _rhs27
				_83_v10 = _rhs28
				_72_v0 = _rhs29
				_83_v10 = _rhs30
				var _326_v179 _dafny.Map
				_ = _326_v179
				_326_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_249_v126, (_81_v9).F2())
				_326_v179 = (_326_v179).Update(_249_v126, _86_v13)
				_84_v11 = (func() _dafny.Int {
					if _83_v10 {
						return _84_v11
					}
					return (_dafny.Zero).Minus((_81_v9).F2())
				})()
				var _327_v180 bool
				_ = _327_v180
				var _328_v181 bool
				_ = _328_v181
				var _out17 bool
				_ = _out17
				var _out18 bool
				_ = _out18
				_out17, _out18 = (_241_v118).M1((_72_v0).Plus(_84_v11), _71_globalState)
				_327_v180 = _out17
				_328_v181 = _out18
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	_86_v13 = _dafny.IntOfInt64(670)
	_74_v3 = _74_v3
	_dafny.Print(_72_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_74_v3.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_v4).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("dtn"), _dafny.UnicodeSeqOfUtf8Bytes("dtn"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_76_v5).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_76_v5).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_76_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_76_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_76_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_78_v6).Dtor_cf0()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_78_v6).Dtor_cf0()).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_78_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_78_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_78_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_78_v6).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_79_v7, _dafny.SeqOf(_dafny.IntOfInt64(936))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v8).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_81_v9).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_81_v9).F3()).Dtor_cf0()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_81_v9).F3()).Dtor_cf0()).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_81_v9).F3()).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_81_v9).F3()).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_81_v9).F3()).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_81_v9).F3()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_83_v10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_84_v11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(936), _dafny.CodePoint('i')).UpdateUnsafe(_dafny.IntOfInt64(3), _dafny.CodePoint('i'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_86_v13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_89_v14).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(548))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v15).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(548)), _dafny.UnicodeSeqOfUtf8Bytes("dtn"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_v16).Dtor_cf85()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(548)), _dafny.UnicodeSeqOfUtf8Bytes("dtn"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v89).Dtor_cf94())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v112).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(936), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_221_i9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_243_v120).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v121).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v122).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_v123).Equals(_dafny.SetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_247_v124).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_248_v125).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(341), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_249_v126, _dafny.SeqOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_250_v127, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf(false, true)).UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true)).UpdateUnsafe(true, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_251_v128).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_252_v129)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v130).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v130).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v130).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_254_v131).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v132).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_322_i13)
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

type D0_DC0 struct {
	Cf0 _dafny.Array
	Cf1 _dafny.Int
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Array, Cf1 _dafny.Int) D0 {
	return D0{D0_DC0{Cf0, Cf1}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
	Cf2 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf2 _dafny.Int) D0 {
	return D0{D0_DC1{Cf2}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D0) Dtor_cf0() _dafny.Array {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf2) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D0) Equals(other D0) bool {
	switch data1 := _this.Get_().(type) {
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0 && data1.Cf1.Cmp(data2.Cf1) == 0
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0
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

type D1_DC2 struct {
	Cf3 bool
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf3 bool) D1 {
	return D1{D1_DC2{Cf3}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() bool {
	return false
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf3) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
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

type D2_DC4 struct {
	Cf5 _dafny.Int
	Cf6 bool
	Cf7 _dafny.Array
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf5 _dafny.Int, Cf6 bool, Cf7 _dafny.Array) D2 {
	return D2{D2_DC4{Cf5, Cf6, Cf7}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC5 struct {
	Cf8  _dafny.Int
	Cf9  _dafny.Sequence
	Cf10 bool
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf8 _dafny.Int, Cf9 _dafny.Sequence, Cf10 bool) D2 {
	return D2{D2_DC5{Cf8, Cf9, Cf10}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC3 struct {
	Cf4 _dafny.Set
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf4 _dafny.Set) D2 {
	return D2{D2_DC3{Cf4}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

type D2_DC6 struct {
	Cf11 D2
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf11 D2) D2 {
	return D2{D2_DC6{Cf11}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC4_(_dafny.Zero, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D2) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf5
}

func (_this D2) Dtor_cf6() bool {
	return _this.Get_().(D2_DC4).Cf6
}

func (_this D2) Dtor_cf7() _dafny.Array {
	return _this.Get_().(D2_DC4).Cf7
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) Dtor_cf10() bool {
	return _this.Get_().(D2_DC5).Cf10
}

func (_this D2) Dtor_cf4() _dafny.Set {
	return _this.Get_().(D2_DC3).Cf4
}

func (_this D2) Dtor_cf11() D2 {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf11) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6 == data2.Cf6 && data1.Cf7 == data2.Cf7
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Equals(data2.Cf9) && data1.Cf10 == data2.Cf10
		}
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf4.Equals(data2.Cf4)
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf11.Equals(data2.Cf11)
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

type D3_DC7 struct {
	Cf12 _dafny.Map
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf12 _dafny.Map) D3 {
	return D3{D3_DC7{Cf12}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D3) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D3_DC7).Cf12
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf12.Equals(data2.Cf12)
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

type D4_DC9 struct {
	Cf14 _dafny.MultiSet
	Cf15 _dafny.CodePoint
	Cf16 _dafny.Int
	Cf17 _dafny.Set
	Cf18 _dafny.Int
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf14 _dafny.MultiSet, Cf15 _dafny.CodePoint, Cf16 _dafny.Int, Cf17 _dafny.Set, Cf18 _dafny.Int) D4 {
	return D4{D4_DC9{Cf14, Cf15, Cf16, Cf17, Cf18}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC10 struct {
	Cf19 _dafny.Int
	Cf20 _dafny.Int
	Cf21 _dafny.Int
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf19 _dafny.Int, Cf20 _dafny.Int, Cf21 _dafny.Int) D4 {
	return D4{D4_DC10{Cf19, Cf20, Cf21}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC11 struct {
	Cf22 _dafny.Int
	Cf23 bool
	Cf24 bool
	Cf25 bool
	Cf26 bool
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf22 _dafny.Int, Cf23 bool, Cf24 bool, Cf25 bool, Cf26 bool) D4 {
	return D4{D4_DC11{Cf22, Cf23, Cf24, Cf25, Cf26}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC8 struct {
	Cf13 _dafny.Sequence
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf13 _dafny.Sequence) D4 {
	return D4{D4_DC8{Cf13}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

type D4_DC12 struct {
	Cf27 D4
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf27 D4) D4 {
	return D4{D4_DC12{Cf27}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC9_(_dafny.EmptyMultiSet, _dafny.CodePoint('D'), _dafny.Zero, _dafny.EmptySet, _dafny.Zero)
}

func (_this D4) Dtor_cf14() _dafny.MultiSet {
	return _this.Get_().(D4_DC9).Cf14
}

func (_this D4) Dtor_cf15() _dafny.CodePoint {
	return _this.Get_().(D4_DC9).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf16
}

func (_this D4) Dtor_cf17() _dafny.Set {
	return _this.Get_().(D4_DC9).Cf17
}

func (_this D4) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf18
}

func (_this D4) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf19
}

func (_this D4) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf20
}

func (_this D4) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf21
}

func (_this D4) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf22
}

func (_this D4) Dtor_cf23() bool {
	return _this.Get_().(D4_DC11).Cf23
}

func (_this D4) Dtor_cf24() bool {
	return _this.Get_().(D4_DC11).Cf24
}

func (_this D4) Dtor_cf25() bool {
	return _this.Get_().(D4_DC11).Cf25
}

func (_this D4) Dtor_cf26() bool {
	return _this.Get_().(D4_DC11).Cf26
}

func (_this D4) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D4_DC8).Cf13
}

func (_this D4) Dtor_cf27() D4 {
	return _this.Get_().(D4_DC12).Cf27
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ")"
		}
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf14.Equals(data2.Cf14) && data1.Cf15 == data2.Cf15 && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17.Equals(data2.Cf17) && data1.Cf18.Cmp(data2.Cf18) == 0
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21.Cmp(data2.Cf21) == 0
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf22.Cmp(data2.Cf22) == 0 && data1.Cf23 == data2.Cf23 && data1.Cf24 == data2.Cf24 && data1.Cf25 == data2.Cf25 && data1.Cf26 == data2.Cf26
		}
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
			return ok && data1.Cf13.Equals(data2.Cf13)
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
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

type D5_DC13 struct {
	Cf28 _dafny.Array
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf28 _dafny.Array) D5 {
	return D5{D5_DC13{Cf28}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D5) Dtor_cf28() _dafny.Array {
	return _this.Get_().(D5_DC13).Cf28
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf28 == data2.Cf28
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

type D6_DC15 struct {
	Cf30 _dafny.Int
	Cf31 _dafny.Int
	Cf32 bool
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf30 _dafny.Int, Cf31 _dafny.Int, Cf32 bool) D6 {
	return D6{D6_DC15{Cf30, Cf31, Cf32}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC14 struct {
	Cf29 _dafny.Sequence
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf29 _dafny.Sequence) D6 {
	return D6{D6_DC14{Cf29}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

type D6_DC16 struct {
	Cf33 D6
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf33 D6) D6 {
	return D6{D6_DC16{Cf33}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC15_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D6) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf30
}

func (_this D6) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf31
}

func (_this D6) Dtor_cf32() bool {
	return _this.Get_().(D6_DC15).Cf32
}

func (_this D6) Dtor_cf29() _dafny.Sequence {
	return _this.Get_().(D6_DC14).Cf29
}

func (_this D6) Dtor_cf33() D6 {
	return _this.Get_().(D6_DC16).Cf33
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D6_DC14:
		{
			return "D6.DC14" + "(" + data.Cf29.VerbatimString(true) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf33) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32 == data2.Cf32
		}
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf29.Equals(data2.Cf29)
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf33.Equals(data2.Cf33)
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
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_() D7 {
	return D7{D7_DC18{}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC19 struct {
	Cf35 _dafny.Int
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf35 _dafny.Int) D7 {
	return D7{D7_DC19{Cf35}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC17 struct {
	Cf34 _dafny.Sequence
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf34 _dafny.Sequence) D7 {
	return D7{D7_DC17{Cf34}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC18_()
}

func (_this D7) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D7_DC19).Cf35
}

func (_this D7) Dtor_cf34() _dafny.Sequence {
	return _this.Get_().(D7_DC17).Cf34
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf35) + ")"
		}
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf34) + ")"
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
			_, ok := other.Get_().(D7_DC18)
			return ok
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf35.Cmp(data2.Cf35) == 0
		}
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf34.Equals(data2.Cf34)
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
	Cf36 _dafny.Set
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf36 _dafny.Set) D8 {
	return D8{D8_DC20{Cf36}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D8) Dtor_cf36() _dafny.Set {
	return _this.Get_().(D8_DC20).Cf36
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf36) + ")"
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
			return ok && data1.Cf36.Equals(data2.Cf36)
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

type D9_DC22 struct {
	Cf38 _dafny.Int
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf38 _dafny.Int) D9 {
	return D9{D9_DC22{Cf38}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

type D9_DC21 struct {
	Cf37 _dafny.MultiSet
}

func (D9_DC21) isD9() {}

func (CompanionStruct_D9_) Create_DC21_(Cf37 _dafny.MultiSet) D9 {
	return D9{D9_DC21{Cf37}}
}

func (_this D9) Is_DC21() bool {
	_, ok := _this.Get_().(D9_DC21)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC22_(_dafny.Zero)
}

func (_this D9) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D9_DC22).Cf38
}

func (_this D9) Dtor_cf37() _dafny.MultiSet {
	return _this.Get_().(D9_DC21).Cf37
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf38) + ")"
		}
	case D9_DC21:
		{
			return "D9.DC21" + "(" + _dafny.String(data.Cf37) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D9_DC21:
		{
			data2, ok := other.Get_().(D9_DC21)
			return ok && data1.Cf37.Equals(data2.Cf37)
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

type D10_DC23 struct {
	Cf39 _dafny.Array
}

func (D10_DC23) isD10() {}

func (CompanionStruct_D10_) Create_DC23_(Cf39 _dafny.Array) D10 {
	return D10{D10_DC23{Cf39}}
}

func (_this D10) Is_DC23() bool {
	_, ok := _this.Get_().(D10_DC23)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D10) Dtor_cf39() _dafny.Array {
	return _this.Get_().(D10_DC23).Cf39
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC23:
		{
			return "D10.DC23" + "(" + _dafny.String(data.Cf39) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC23:
		{
			data2, ok := other.Get_().(D10_DC23)
			return ok && data1.Cf39 == data2.Cf39
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

type D11_DC25 struct {
	Cf41 _dafny.CodePoint
	Cf42 _dafny.Int
	Cf43 bool
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_(Cf41 _dafny.CodePoint, Cf42 _dafny.Int, Cf43 bool) D11 {
	return D11{D11_DC25{Cf41, Cf42, Cf43}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

type D11_DC24 struct {
	Cf40 _dafny.Map
}

func (D11_DC24) isD11() {}

func (CompanionStruct_D11_) Create_DC24_(Cf40 _dafny.Map) D11 {
	return D11{D11_DC24{Cf40}}
}

func (_this D11) Is_DC24() bool {
	_, ok := _this.Get_().(D11_DC24)
	return ok
}

type D11_DC26 struct {
	Cf44 D11
}

func (D11_DC26) isD11() {}

func (CompanionStruct_D11_) Create_DC26_(Cf44 D11) D11 {
	return D11{D11_DC26{Cf44}}
}

func (_this D11) Is_DC26() bool {
	_, ok := _this.Get_().(D11_DC26)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC25_(_dafny.CodePoint('D'), _dafny.Zero, false)
}

func (_this D11) Dtor_cf41() _dafny.CodePoint {
	return _this.Get_().(D11_DC25).Cf41
}

func (_this D11) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D11_DC25).Cf42
}

func (_this D11) Dtor_cf43() bool {
	return _this.Get_().(D11_DC25).Cf43
}

func (_this D11) Dtor_cf40() _dafny.Map {
	return _this.Get_().(D11_DC24).Cf40
}

func (_this D11) Dtor_cf44() D11 {
	return _this.Get_().(D11_DC26).Cf44
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC25:
		{
			return "D11.DC25" + "(" + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ")"
		}
	case D11_DC24:
		{
			return "D11.DC24" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D11_DC26:
		{
			return "D11.DC26" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC25:
		{
			data2, ok := other.Get_().(D11_DC25)
			return ok && data1.Cf41 == data2.Cf41 && data1.Cf42.Cmp(data2.Cf42) == 0 && data1.Cf43 == data2.Cf43
		}
	case D11_DC24:
		{
			data2, ok := other.Get_().(D11_DC24)
			return ok && data1.Cf40.Equals(data2.Cf40)
		}
	case D11_DC26:
		{
			data2, ok := other.Get_().(D11_DC26)
			return ok && data1.Cf44.Equals(data2.Cf44)
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

type D12_DC28 struct {
	Cf46 bool
	Cf47 _dafny.Int
}

func (D12_DC28) isD12() {}

func (CompanionStruct_D12_) Create_DC28_(Cf46 bool, Cf47 _dafny.Int) D12 {
	return D12{D12_DC28{Cf46, Cf47}}
}

func (_this D12) Is_DC28() bool {
	_, ok := _this.Get_().(D12_DC28)
	return ok
}

type D12_DC29 struct {
	Cf48 bool
	Cf49 _dafny.Int
	Cf50 bool
	Cf51 _dafny.Int
	Cf52 bool
}

func (D12_DC29) isD12() {}

func (CompanionStruct_D12_) Create_DC29_(Cf48 bool, Cf49 _dafny.Int, Cf50 bool, Cf51 _dafny.Int, Cf52 bool) D12 {
	return D12{D12_DC29{Cf48, Cf49, Cf50, Cf51, Cf52}}
}

func (_this D12) Is_DC29() bool {
	_, ok := _this.Get_().(D12_DC29)
	return ok
}

type D12_DC27 struct {
	Cf45 _dafny.Set
}

func (D12_DC27) isD12() {}

func (CompanionStruct_D12_) Create_DC27_(Cf45 _dafny.Set) D12 {
	return D12{D12_DC27{Cf45}}
}

func (_this D12) Is_DC27() bool {
	_, ok := _this.Get_().(D12_DC27)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC28_(false, _dafny.Zero)
}

func (_this D12) Dtor_cf46() bool {
	return _this.Get_().(D12_DC28).Cf46
}

func (_this D12) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D12_DC28).Cf47
}

func (_this D12) Dtor_cf48() bool {
	return _this.Get_().(D12_DC29).Cf48
}

func (_this D12) Dtor_cf49() _dafny.Int {
	return _this.Get_().(D12_DC29).Cf49
}

func (_this D12) Dtor_cf50() bool {
	return _this.Get_().(D12_DC29).Cf50
}

func (_this D12) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D12_DC29).Cf51
}

func (_this D12) Dtor_cf52() bool {
	return _this.Get_().(D12_DC29).Cf52
}

func (_this D12) Dtor_cf45() _dafny.Set {
	return _this.Get_().(D12_DC27).Cf45
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC28:
		{
			return "D12.DC28" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ")"
		}
	case D12_DC29:
		{
			return "D12.DC29" + "(" + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ")"
		}
	case D12_DC27:
		{
			return "D12.DC27" + "(" + _dafny.String(data.Cf45) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC28:
		{
			data2, ok := other.Get_().(D12_DC28)
			return ok && data1.Cf46 == data2.Cf46 && data1.Cf47.Cmp(data2.Cf47) == 0
		}
	case D12_DC29:
		{
			data2, ok := other.Get_().(D12_DC29)
			return ok && data1.Cf48 == data2.Cf48 && data1.Cf49.Cmp(data2.Cf49) == 0 && data1.Cf50 == data2.Cf50 && data1.Cf51.Cmp(data2.Cf51) == 0 && data1.Cf52 == data2.Cf52
		}
	case D12_DC27:
		{
			data2, ok := other.Get_().(D12_DC27)
			return ok && data1.Cf45.Equals(data2.Cf45)
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

type D13_DC31 struct {
	Cf54 bool
	Cf55 _dafny.Int
	Cf56 _dafny.Int
}

func (D13_DC31) isD13() {}

func (CompanionStruct_D13_) Create_DC31_(Cf54 bool, Cf55 _dafny.Int, Cf56 _dafny.Int) D13 {
	return D13{D13_DC31{Cf54, Cf55, Cf56}}
}

func (_this D13) Is_DC31() bool {
	_, ok := _this.Get_().(D13_DC31)
	return ok
}

type D13_DC32 struct {
	Cf57 _dafny.Sequence
	Cf58 bool
	Cf59 _dafny.Int
	Cf60 _dafny.Array
}

func (D13_DC32) isD13() {}

func (CompanionStruct_D13_) Create_DC32_(Cf57 _dafny.Sequence, Cf58 bool, Cf59 _dafny.Int, Cf60 _dafny.Array) D13 {
	return D13{D13_DC32{Cf57, Cf58, Cf59, Cf60}}
}

func (_this D13) Is_DC32() bool {
	_, ok := _this.Get_().(D13_DC32)
	return ok
}

type D13_DC30 struct {
	Cf53 _dafny.Array
}

func (D13_DC30) isD13() {}

func (CompanionStruct_D13_) Create_DC30_(Cf53 _dafny.Array) D13 {
	return D13{D13_DC30{Cf53}}
}

func (_this D13) Is_DC30() bool {
	_, ok := _this.Get_().(D13_DC30)
	return ok
}

type D13_DC33 struct {
	Cf61 D13
}

func (D13_DC33) isD13() {}

func (CompanionStruct_D13_) Create_DC33_(Cf61 D13) D13 {
	return D13{D13_DC33{Cf61}}
}

func (_this D13) Is_DC33() bool {
	_, ok := _this.Get_().(D13_DC33)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC31_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D13) Dtor_cf54() bool {
	return _this.Get_().(D13_DC31).Cf54
}

func (_this D13) Dtor_cf55() _dafny.Int {
	return _this.Get_().(D13_DC31).Cf55
}

func (_this D13) Dtor_cf56() _dafny.Int {
	return _this.Get_().(D13_DC31).Cf56
}

func (_this D13) Dtor_cf57() _dafny.Sequence {
	return _this.Get_().(D13_DC32).Cf57
}

func (_this D13) Dtor_cf58() bool {
	return _this.Get_().(D13_DC32).Cf58
}

func (_this D13) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D13_DC32).Cf59
}

func (_this D13) Dtor_cf60() _dafny.Array {
	return _this.Get_().(D13_DC32).Cf60
}

func (_this D13) Dtor_cf53() _dafny.Array {
	return _this.Get_().(D13_DC30).Cf53
}

func (_this D13) Dtor_cf61() D13 {
	return _this.Get_().(D13_DC33).Cf61
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC31:
		{
			return "D13.DC31" + "(" + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ")"
		}
	case D13_DC32:
		{
			return "D13.DC32" + "(" + data.Cf57.VerbatimString(true) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ")"
		}
	case D13_DC30:
		{
			return "D13.DC30" + "(" + _dafny.String(data.Cf53) + ")"
		}
	case D13_DC33:
		{
			return "D13.DC33" + "(" + _dafny.String(data.Cf61) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC31:
		{
			data2, ok := other.Get_().(D13_DC31)
			return ok && data1.Cf54 == data2.Cf54 && data1.Cf55.Cmp(data2.Cf55) == 0 && data1.Cf56.Cmp(data2.Cf56) == 0
		}
	case D13_DC32:
		{
			data2, ok := other.Get_().(D13_DC32)
			return ok && data1.Cf57.Equals(data2.Cf57) && data1.Cf58 == data2.Cf58 && data1.Cf59.Cmp(data2.Cf59) == 0 && data1.Cf60 == data2.Cf60
		}
	case D13_DC30:
		{
			data2, ok := other.Get_().(D13_DC30)
			return ok && data1.Cf53 == data2.Cf53
		}
	case D13_DC33:
		{
			data2, ok := other.Get_().(D13_DC33)
			return ok && data1.Cf61.Equals(data2.Cf61)
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

type D14_DC35 struct {
	Cf63 _dafny.Int
	Cf64 _dafny.Sequence
	Cf65 _dafny.Int
	Cf66 bool
	Cf67 bool
}

func (D14_DC35) isD14() {}

func (CompanionStruct_D14_) Create_DC35_(Cf63 _dafny.Int, Cf64 _dafny.Sequence, Cf65 _dafny.Int, Cf66 bool, Cf67 bool) D14 {
	return D14{D14_DC35{Cf63, Cf64, Cf65, Cf66, Cf67}}
}

func (_this D14) Is_DC35() bool {
	_, ok := _this.Get_().(D14_DC35)
	return ok
}

type D14_DC36 struct {
	Cf68 _dafny.Map
	Cf69 _dafny.Map
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf68 _dafny.Map, Cf69 _dafny.Map) D14 {
	return D14{D14_DC36{Cf68, Cf69}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

type D14_DC34 struct {
	Cf62 _dafny.Map
}

func (D14_DC34) isD14() {}

func (CompanionStruct_D14_) Create_DC34_(Cf62 _dafny.Map) D14 {
	return D14{D14_DC34{Cf62}}
}

func (_this D14) Is_DC34() bool {
	_, ok := _this.Get_().(D14_DC34)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC35_(_dafny.Zero, _dafny.EmptySeq, _dafny.Zero, false, false)
}

func (_this D14) Dtor_cf63() _dafny.Int {
	return _this.Get_().(D14_DC35).Cf63
}

func (_this D14) Dtor_cf64() _dafny.Sequence {
	return _this.Get_().(D14_DC35).Cf64
}

func (_this D14) Dtor_cf65() _dafny.Int {
	return _this.Get_().(D14_DC35).Cf65
}

func (_this D14) Dtor_cf66() bool {
	return _this.Get_().(D14_DC35).Cf66
}

func (_this D14) Dtor_cf67() bool {
	return _this.Get_().(D14_DC35).Cf67
}

func (_this D14) Dtor_cf68() _dafny.Map {
	return _this.Get_().(D14_DC36).Cf68
}

func (_this D14) Dtor_cf69() _dafny.Map {
	return _this.Get_().(D14_DC36).Cf69
}

func (_this D14) Dtor_cf62() _dafny.Map {
	return _this.Get_().(D14_DC34).Cf62
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC35:
		{
			return "D14.DC35" + "(" + _dafny.String(data.Cf63) + ", " + data.Cf64.VerbatimString(true) + ", " + _dafny.String(data.Cf65) + ", " + _dafny.String(data.Cf66) + ", " + _dafny.String(data.Cf67) + ")"
		}
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ")"
		}
	case D14_DC34:
		{
			return "D14.DC34" + "(" + _dafny.String(data.Cf62) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC35:
		{
			data2, ok := other.Get_().(D14_DC35)
			return ok && data1.Cf63.Cmp(data2.Cf63) == 0 && data1.Cf64.Equals(data2.Cf64) && data1.Cf65.Cmp(data2.Cf65) == 0 && data1.Cf66 == data2.Cf66 && data1.Cf67 == data2.Cf67
		}
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && data1.Cf68.Equals(data2.Cf68) && data1.Cf69.Equals(data2.Cf69)
		}
	case D14_DC34:
		{
			data2, ok := other.Get_().(D14_DC34)
			return ok && data1.Cf62.Equals(data2.Cf62)
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

type D15_DC38 struct {
	Cf71 _dafny.CodePoint
}

func (D15_DC38) isD15() {}

func (CompanionStruct_D15_) Create_DC38_(Cf71 _dafny.CodePoint) D15 {
	return D15{D15_DC38{Cf71}}
}

func (_this D15) Is_DC38() bool {
	_, ok := _this.Get_().(D15_DC38)
	return ok
}

type D15_DC39 struct {
	Cf72 _dafny.Sequence
	Cf73 D9
}

func (D15_DC39) isD15() {}

func (CompanionStruct_D15_) Create_DC39_(Cf72 _dafny.Sequence, Cf73 D9) D15 {
	return D15{D15_DC39{Cf72, Cf73}}
}

func (_this D15) Is_DC39() bool {
	_, ok := _this.Get_().(D15_DC39)
	return ok
}

type D15_DC37 struct {
	Cf70 _dafny.Set
}

func (D15_DC37) isD15() {}

func (CompanionStruct_D15_) Create_DC37_(Cf70 _dafny.Set) D15 {
	return D15{D15_DC37{Cf70}}
}

func (_this D15) Is_DC37() bool {
	_, ok := _this.Get_().(D15_DC37)
	return ok
}

type D15_DC40 struct {
	Cf74 D15
}

func (D15_DC40) isD15() {}

func (CompanionStruct_D15_) Create_DC40_(Cf74 D15) D15 {
	return D15{D15_DC40{Cf74}}
}

func (_this D15) Is_DC40() bool {
	_, ok := _this.Get_().(D15_DC40)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC38_(_dafny.CodePoint('D'))
}

func (_this D15) Dtor_cf71() _dafny.CodePoint {
	return _this.Get_().(D15_DC38).Cf71
}

func (_this D15) Dtor_cf72() _dafny.Sequence {
	return _this.Get_().(D15_DC39).Cf72
}

func (_this D15) Dtor_cf73() D9 {
	return _this.Get_().(D15_DC39).Cf73
}

func (_this D15) Dtor_cf70() _dafny.Set {
	return _this.Get_().(D15_DC37).Cf70
}

func (_this D15) Dtor_cf74() D15 {
	return _this.Get_().(D15_DC40).Cf74
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC38:
		{
			return "D15.DC38" + "(" + _dafny.String(data.Cf71) + ")"
		}
	case D15_DC39:
		{
			return "D15.DC39" + "(" + data.Cf72.VerbatimString(true) + ", " + _dafny.String(data.Cf73) + ")"
		}
	case D15_DC37:
		{
			return "D15.DC37" + "(" + _dafny.String(data.Cf70) + ")"
		}
	case D15_DC40:
		{
			return "D15.DC40" + "(" + _dafny.String(data.Cf74) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC38:
		{
			data2, ok := other.Get_().(D15_DC38)
			return ok && data1.Cf71 == data2.Cf71
		}
	case D15_DC39:
		{
			data2, ok := other.Get_().(D15_DC39)
			return ok && data1.Cf72.Equals(data2.Cf72) && data1.Cf73.Equals(data2.Cf73)
		}
	case D15_DC37:
		{
			data2, ok := other.Get_().(D15_DC37)
			return ok && data1.Cf70.Equals(data2.Cf70)
		}
	case D15_DC40:
		{
			data2, ok := other.Get_().(D15_DC40)
			return ok && data1.Cf74.Equals(data2.Cf74)
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

type D16_DC42 struct {
	Cf76 _dafny.Int
	Cf77 _dafny.Int
	Cf78 _dafny.Int
}

func (D16_DC42) isD16() {}

func (CompanionStruct_D16_) Create_DC42_(Cf76 _dafny.Int, Cf77 _dafny.Int, Cf78 _dafny.Int) D16 {
	return D16{D16_DC42{Cf76, Cf77, Cf78}}
}

func (_this D16) Is_DC42() bool {
	_, ok := _this.Get_().(D16_DC42)
	return ok
}

type D16_DC43 struct {
	Cf79 _dafny.Sequence
	Cf80 _dafny.CodePoint
	Cf81 _dafny.Int
}

func (D16_DC43) isD16() {}

func (CompanionStruct_D16_) Create_DC43_(Cf79 _dafny.Sequence, Cf80 _dafny.CodePoint, Cf81 _dafny.Int) D16 {
	return D16{D16_DC43{Cf79, Cf80, Cf81}}
}

func (_this D16) Is_DC43() bool {
	_, ok := _this.Get_().(D16_DC43)
	return ok
}

type D16_DC44 struct {
	Cf82 _dafny.Int
}

func (D16_DC44) isD16() {}

func (CompanionStruct_D16_) Create_DC44_(Cf82 _dafny.Int) D16 {
	return D16{D16_DC44{Cf82}}
}

func (_this D16) Is_DC44() bool {
	_, ok := _this.Get_().(D16_DC44)
	return ok
}

type D16_DC41 struct {
	Cf75 _dafny.Array
}

func (D16_DC41) isD16() {}

func (CompanionStruct_D16_) Create_DC41_(Cf75 _dafny.Array) D16 {
	return D16{D16_DC41{Cf75}}
}

func (_this D16) Is_DC41() bool {
	_, ok := _this.Get_().(D16_DC41)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC42_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D16) Dtor_cf76() _dafny.Int {
	return _this.Get_().(D16_DC42).Cf76
}

func (_this D16) Dtor_cf77() _dafny.Int {
	return _this.Get_().(D16_DC42).Cf77
}

func (_this D16) Dtor_cf78() _dafny.Int {
	return _this.Get_().(D16_DC42).Cf78
}

func (_this D16) Dtor_cf79() _dafny.Sequence {
	return _this.Get_().(D16_DC43).Cf79
}

func (_this D16) Dtor_cf80() _dafny.CodePoint {
	return _this.Get_().(D16_DC43).Cf80
}

func (_this D16) Dtor_cf81() _dafny.Int {
	return _this.Get_().(D16_DC43).Cf81
}

func (_this D16) Dtor_cf82() _dafny.Int {
	return _this.Get_().(D16_DC44).Cf82
}

func (_this D16) Dtor_cf75() _dafny.Array {
	return _this.Get_().(D16_DC41).Cf75
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC42:
		{
			return "D16.DC42" + "(" + _dafny.String(data.Cf76) + ", " + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ")"
		}
	case D16_DC43:
		{
			return "D16.DC43" + "(" + data.Cf79.VerbatimString(true) + ", " + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ")"
		}
	case D16_DC44:
		{
			return "D16.DC44" + "(" + _dafny.String(data.Cf82) + ")"
		}
	case D16_DC41:
		{
			return "D16.DC41" + "(" + _dafny.String(data.Cf75) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC42:
		{
			data2, ok := other.Get_().(D16_DC42)
			return ok && data1.Cf76.Cmp(data2.Cf76) == 0 && data1.Cf77.Cmp(data2.Cf77) == 0 && data1.Cf78.Cmp(data2.Cf78) == 0
		}
	case D16_DC43:
		{
			data2, ok := other.Get_().(D16_DC43)
			return ok && data1.Cf79.Equals(data2.Cf79) && data1.Cf80 == data2.Cf80 && data1.Cf81.Cmp(data2.Cf81) == 0
		}
	case D16_DC44:
		{
			data2, ok := other.Get_().(D16_DC44)
			return ok && data1.Cf82.Cmp(data2.Cf82) == 0
		}
	case D16_DC41:
		{
			data2, ok := other.Get_().(D16_DC41)
			return ok && data1.Cf75 == data2.Cf75
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

type D17_DC46 struct {
	Cf84 _dafny.Int
}

func (D17_DC46) isD17() {}

func (CompanionStruct_D17_) Create_DC46_(Cf84 _dafny.Int) D17 {
	return D17{D17_DC46{Cf84}}
}

func (_this D17) Is_DC46() bool {
	_, ok := _this.Get_().(D17_DC46)
	return ok
}

type D17_DC45 struct {
	Cf83 _dafny.MultiSet
}

func (D17_DC45) isD17() {}

func (CompanionStruct_D17_) Create_DC45_(Cf83 _dafny.MultiSet) D17 {
	return D17{D17_DC45{Cf83}}
}

func (_this D17) Is_DC45() bool {
	_, ok := _this.Get_().(D17_DC45)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC46_(_dafny.Zero)
}

func (_this D17) Dtor_cf84() _dafny.Int {
	return _this.Get_().(D17_DC46).Cf84
}

func (_this D17) Dtor_cf83() _dafny.MultiSet {
	return _this.Get_().(D17_DC45).Cf83
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC46:
		{
			return "D17.DC46" + "(" + _dafny.String(data.Cf84) + ")"
		}
	case D17_DC45:
		{
			return "D17.DC45" + "(" + _dafny.String(data.Cf83) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC46:
		{
			data2, ok := other.Get_().(D17_DC46)
			return ok && data1.Cf84.Cmp(data2.Cf84) == 0
		}
	case D17_DC45:
		{
			data2, ok := other.Get_().(D17_DC45)
			return ok && data1.Cf83.Equals(data2.Cf83)
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

type D18_DC48 struct {
}

func (D18_DC48) isD18() {}

func (CompanionStruct_D18_) Create_DC48_() D18 {
	return D18{D18_DC48{}}
}

func (_this D18) Is_DC48() bool {
	_, ok := _this.Get_().(D18_DC48)
	return ok
}

type D18_DC49 struct {
	Cf86 _dafny.Set
}

func (D18_DC49) isD18() {}

func (CompanionStruct_D18_) Create_DC49_(Cf86 _dafny.Set) D18 {
	return D18{D18_DC49{Cf86}}
}

func (_this D18) Is_DC49() bool {
	_, ok := _this.Get_().(D18_DC49)
	return ok
}

type D18_DC50 struct {
	Cf87 bool
	Cf88 _dafny.Array
	Cf89 bool
	Cf90 bool
	Cf91 _dafny.Sequence
}

func (D18_DC50) isD18() {}

func (CompanionStruct_D18_) Create_DC50_(Cf87 bool, Cf88 _dafny.Array, Cf89 bool, Cf90 bool, Cf91 _dafny.Sequence) D18 {
	return D18{D18_DC50{Cf87, Cf88, Cf89, Cf90, Cf91}}
}

func (_this D18) Is_DC50() bool {
	_, ok := _this.Get_().(D18_DC50)
	return ok
}

type D18_DC47 struct {
	Cf85 _dafny.Map
}

func (D18_DC47) isD18() {}

func (CompanionStruct_D18_) Create_DC47_(Cf85 _dafny.Map) D18 {
	return D18{D18_DC47{Cf85}}
}

func (_this D18) Is_DC47() bool {
	_, ok := _this.Get_().(D18_DC47)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC48_()
}

func (_this D18) Dtor_cf86() _dafny.Set {
	return _this.Get_().(D18_DC49).Cf86
}

func (_this D18) Dtor_cf87() bool {
	return _this.Get_().(D18_DC50).Cf87
}

func (_this D18) Dtor_cf88() _dafny.Array {
	return _this.Get_().(D18_DC50).Cf88
}

func (_this D18) Dtor_cf89() bool {
	return _this.Get_().(D18_DC50).Cf89
}

func (_this D18) Dtor_cf90() bool {
	return _this.Get_().(D18_DC50).Cf90
}

func (_this D18) Dtor_cf91() _dafny.Sequence {
	return _this.Get_().(D18_DC50).Cf91
}

func (_this D18) Dtor_cf85() _dafny.Map {
	return _this.Get_().(D18_DC47).Cf85
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC48:
		{
			return "D18.DC48"
		}
	case D18_DC49:
		{
			return "D18.DC49" + "(" + _dafny.String(data.Cf86) + ")"
		}
	case D18_DC50:
		{
			return "D18.DC50" + "(" + _dafny.String(data.Cf87) + ", " + _dafny.String(data.Cf88) + ", " + _dafny.String(data.Cf89) + ", " + _dafny.String(data.Cf90) + ", " + data.Cf91.VerbatimString(true) + ")"
		}
	case D18_DC47:
		{
			return "D18.DC47" + "(" + _dafny.String(data.Cf85) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC48:
		{
			_, ok := other.Get_().(D18_DC48)
			return ok
		}
	case D18_DC49:
		{
			data2, ok := other.Get_().(D18_DC49)
			return ok && data1.Cf86.Equals(data2.Cf86)
		}
	case D18_DC50:
		{
			data2, ok := other.Get_().(D18_DC50)
			return ok && data1.Cf87 == data2.Cf87 && data1.Cf88 == data2.Cf88 && data1.Cf89 == data2.Cf89 && data1.Cf90 == data2.Cf90 && data1.Cf91.Equals(data2.Cf91)
		}
	case D18_DC47:
		{
			data2, ok := other.Get_().(D18_DC47)
			return ok && data1.Cf85.Equals(data2.Cf85)
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

type D19_DC51 struct {
	Cf92 *C5
}

func (D19_DC51) isD19() {}

func (CompanionStruct_D19_) Create_DC51_(Cf92 *C5) D19 {
	return D19{D19_DC51{Cf92}}
}

func (_this D19) Is_DC51() bool {
	_, ok := _this.Get_().(D19_DC51)
	return ok
}

func (CompanionStruct_D19_) Default() *C5 {
	return (*C5)(nil)
}

func (_this D19) Dtor_cf92() *C5 {
	return _this.Get_().(D19_DC51).Cf92
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC51:
		{
			return "D19.DC51" + "(" + _dafny.String(data.Cf92) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC51:
		{
			data2, ok := other.Get_().(D19_DC51)
			return ok && data1.Cf92 == data2.Cf92
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

type D20_DC53 struct {
	Cf94 bool
}

func (D20_DC53) isD20() {}

func (CompanionStruct_D20_) Create_DC53_(Cf94 bool) D20 {
	return D20{D20_DC53{Cf94}}
}

func (_this D20) Is_DC53() bool {
	_, ok := _this.Get_().(D20_DC53)
	return ok
}

type D20_DC54 struct {
	Cf95 _dafny.Sequence
	Cf96 _dafny.Int
	Cf97 bool
}

func (D20_DC54) isD20() {}

func (CompanionStruct_D20_) Create_DC54_(Cf95 _dafny.Sequence, Cf96 _dafny.Int, Cf97 bool) D20 {
	return D20{D20_DC54{Cf95, Cf96, Cf97}}
}

func (_this D20) Is_DC54() bool {
	_, ok := _this.Get_().(D20_DC54)
	return ok
}

type D20_DC52 struct {
	Cf93 _dafny.Array
}

func (D20_DC52) isD20() {}

func (CompanionStruct_D20_) Create_DC52_(Cf93 _dafny.Array) D20 {
	return D20{D20_DC52{Cf93}}
}

func (_this D20) Is_DC52() bool {
	_, ok := _this.Get_().(D20_DC52)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC53_(false)
}

func (_this D20) Dtor_cf94() bool {
	return _this.Get_().(D20_DC53).Cf94
}

func (_this D20) Dtor_cf95() _dafny.Sequence {
	return _this.Get_().(D20_DC54).Cf95
}

func (_this D20) Dtor_cf96() _dafny.Int {
	return _this.Get_().(D20_DC54).Cf96
}

func (_this D20) Dtor_cf97() bool {
	return _this.Get_().(D20_DC54).Cf97
}

func (_this D20) Dtor_cf93() _dafny.Array {
	return _this.Get_().(D20_DC52).Cf93
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC53:
		{
			return "D20.DC53" + "(" + _dafny.String(data.Cf94) + ")"
		}
	case D20_DC54:
		{
			return "D20.DC54" + "(" + data.Cf95.VerbatimString(true) + ", " + _dafny.String(data.Cf96) + ", " + _dafny.String(data.Cf97) + ")"
		}
	case D20_DC52:
		{
			return "D20.DC52" + "(" + _dafny.String(data.Cf93) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC53:
		{
			data2, ok := other.Get_().(D20_DC53)
			return ok && data1.Cf94 == data2.Cf94
		}
	case D20_DC54:
		{
			data2, ok := other.Get_().(D20_DC54)
			return ok && data1.Cf95.Equals(data2.Cf95) && data1.Cf96.Cmp(data2.Cf96) == 0 && data1.Cf97 == data2.Cf97
		}
	case D20_DC52:
		{
			data2, ok := other.Get_().(D20_DC52)
			return ok && data1.Cf93 == data2.Cf93
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

// Definition of datatype D21
type D21 struct {
	Data_D21_
}

func (_this D21) Get_() Data_D21_ {
	return _this.Data_D21_
}

type Data_D21_ interface {
	isD21()
}

type CompanionStruct_D21_ struct {
}

var Companion_D21_ = CompanionStruct_D21_{}

type D21_DC55 struct {
	Cf98 T0
}

func (D21_DC55) isD21() {}

func (CompanionStruct_D21_) Create_DC55_(Cf98 T0) D21 {
	return D21{D21_DC55{Cf98}}
}

func (_this D21) Is_DC55() bool {
	_, ok := _this.Get_().(D21_DC55)
	return ok
}

func (CompanionStruct_D21_) Default() T0 {
	return (T0)(nil)
}

func (_this D21) Dtor_cf98() T0 {
	return _this.Get_().(D21_DC55).Cf98
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC55:
		{
			return "D21.DC55" + "(" + _dafny.String(data.Cf98) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC55:
		{
			data2, ok := other.Get_().(D21_DC55)
			return ok && _dafny.AreEqual(data1.Cf98, data2.Cf98)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D21) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D21)
	return ok && _this.Equals(typed)
}

func Type_D21_() _dafny.TypeDescriptor {
	return type_D21_{}
}

type type_D21_ struct {
}

func (_this type_D21_) Default() interface{} {
	return Companion_D21_.Default()
}

func (_this type_D21_) String() string {
	return "main.D21"
}
func (_this D21) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D21{}

// End of datatype D21

// Definition of datatype D22
type D22 struct {
	Data_D22_
}

func (_this D22) Get_() Data_D22_ {
	return _this.Data_D22_
}

type Data_D22_ interface {
	isD22()
}

type CompanionStruct_D22_ struct {
}

var Companion_D22_ = CompanionStruct_D22_{}

type D22_DC56 struct {
	Cf99 *C4
}

func (D22_DC56) isD22() {}

func (CompanionStruct_D22_) Create_DC56_(Cf99 *C4) D22 {
	return D22{D22_DC56{Cf99}}
}

func (_this D22) Is_DC56() bool {
	_, ok := _this.Get_().(D22_DC56)
	return ok
}

func (CompanionStruct_D22_) Default() *C4 {
	return (*C4)(nil)
}

func (_this D22) Dtor_cf99() *C4 {
	return _this.Get_().(D22_DC56).Cf99
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC56:
		{
			return "D22.DC56" + "(" + _dafny.String(data.Cf99) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC56:
		{
			data2, ok := other.Get_().(D22_DC56)
			return ok && data1.Cf99 == data2.Cf99
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D22) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D22)
	return ok && _this.Equals(typed)
}

func Type_D22_() _dafny.TypeDescriptor {
	return type_D22_{}
}

type type_D22_ struct {
}

func (_this type_D22_) Default() interface{} {
	return Companion_D22_.Default()
}

func (_this type_D22_) String() string {
	return "main.D22"
}
func (_this D22) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D22{}

// End of datatype D22

// Definition of datatype D23
type D23 struct {
	Data_D23_
}

func (_this D23) Get_() Data_D23_ {
	return _this.Data_D23_
}

type Data_D23_ interface {
	isD23()
}

type CompanionStruct_D23_ struct {
}

var Companion_D23_ = CompanionStruct_D23_{}

type D23_DC58 struct {
}

func (D23_DC58) isD23() {}

func (CompanionStruct_D23_) Create_DC58_() D23 {
	return D23{D23_DC58{}}
}

func (_this D23) Is_DC58() bool {
	_, ok := _this.Get_().(D23_DC58)
	return ok
}

type D23_DC57 struct {
	Cf100 _dafny.Set
}

func (D23_DC57) isD23() {}

func (CompanionStruct_D23_) Create_DC57_(Cf100 _dafny.Set) D23 {
	return D23{D23_DC57{Cf100}}
}

func (_this D23) Is_DC57() bool {
	_, ok := _this.Get_().(D23_DC57)
	return ok
}

func (CompanionStruct_D23_) Default() D23 {
	return Companion_D23_.Create_DC58_()
}

func (_this D23) Dtor_cf100() _dafny.Set {
	return _this.Get_().(D23_DC57).Cf100
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC58:
		{
			return "D23.DC58"
		}
	case D23_DC57:
		{
			return "D23.DC57" + "(" + _dafny.String(data.Cf100) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC58:
		{
			_, ok := other.Get_().(D23_DC58)
			return ok
		}
	case D23_DC57:
		{
			data2, ok := other.Get_().(D23_DC57)
			return ok && data1.Cf100.Equals(data2.Cf100)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D23) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D23)
	return ok && _this.Equals(typed)
}

func Type_D23_() _dafny.TypeDescriptor {
	return type_D23_{}
}

type type_D23_ struct {
}

func (_this type_D23_) Default() interface{} {
	return Companion_D23_.Default()
}

func (_this type_D23_) String() string {
	return "main.D23"
}
func (_this D23) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D23{}

// End of datatype D23

// Definition of datatype D24
type D24 struct {
	Data_D24_
}

func (_this D24) Get_() Data_D24_ {
	return _this.Data_D24_
}

type Data_D24_ interface {
	isD24()
}

type CompanionStruct_D24_ struct {
}

var Companion_D24_ = CompanionStruct_D24_{}

type D24_DC60 struct {
}

func (D24_DC60) isD24() {}

func (CompanionStruct_D24_) Create_DC60_() D24 {
	return D24{D24_DC60{}}
}

func (_this D24) Is_DC60() bool {
	_, ok := _this.Get_().(D24_DC60)
	return ok
}

type D24_DC61 struct {
	Cf102 _dafny.Int
	Cf103 _dafny.MultiSet
}

func (D24_DC61) isD24() {}

func (CompanionStruct_D24_) Create_DC61_(Cf102 _dafny.Int, Cf103 _dafny.MultiSet) D24 {
	return D24{D24_DC61{Cf102, Cf103}}
}

func (_this D24) Is_DC61() bool {
	_, ok := _this.Get_().(D24_DC61)
	return ok
}

type D24_DC62 struct {
	Cf104 _dafny.Int
}

func (D24_DC62) isD24() {}

func (CompanionStruct_D24_) Create_DC62_(Cf104 _dafny.Int) D24 {
	return D24{D24_DC62{Cf104}}
}

func (_this D24) Is_DC62() bool {
	_, ok := _this.Get_().(D24_DC62)
	return ok
}

type D24_DC59 struct {
	Cf101 _dafny.Array
}

func (D24_DC59) isD24() {}

func (CompanionStruct_D24_) Create_DC59_(Cf101 _dafny.Array) D24 {
	return D24{D24_DC59{Cf101}}
}

func (_this D24) Is_DC59() bool {
	_, ok := _this.Get_().(D24_DC59)
	return ok
}

func (CompanionStruct_D24_) Default() D24 {
	return Companion_D24_.Create_DC60_()
}

func (_this D24) Dtor_cf102() _dafny.Int {
	return _this.Get_().(D24_DC61).Cf102
}

func (_this D24) Dtor_cf103() _dafny.MultiSet {
	return _this.Get_().(D24_DC61).Cf103
}

func (_this D24) Dtor_cf104() _dafny.Int {
	return _this.Get_().(D24_DC62).Cf104
}

func (_this D24) Dtor_cf101() _dafny.Array {
	return _this.Get_().(D24_DC59).Cf101
}

func (_this D24) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D24_DC60:
		{
			return "D24.DC60"
		}
	case D24_DC61:
		{
			return "D24.DC61" + "(" + _dafny.String(data.Cf102) + ", " + _dafny.String(data.Cf103) + ")"
		}
	case D24_DC62:
		{
			return "D24.DC62" + "(" + _dafny.String(data.Cf104) + ")"
		}
	case D24_DC59:
		{
			return "D24.DC59" + "(" + _dafny.String(data.Cf101) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D24) Equals(other D24) bool {
	switch data1 := _this.Get_().(type) {
	case D24_DC60:
		{
			_, ok := other.Get_().(D24_DC60)
			return ok
		}
	case D24_DC61:
		{
			data2, ok := other.Get_().(D24_DC61)
			return ok && data1.Cf102.Cmp(data2.Cf102) == 0 && data1.Cf103.Equals(data2.Cf103)
		}
	case D24_DC62:
		{
			data2, ok := other.Get_().(D24_DC62)
			return ok && data1.Cf104.Cmp(data2.Cf104) == 0
		}
	case D24_DC59:
		{
			data2, ok := other.Get_().(D24_DC59)
			return ok && data1.Cf101 == data2.Cf101
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D24) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D24)
	return ok && _this.Equals(typed)
}

func Type_D24_() _dafny.TypeDescriptor {
	return type_D24_{}
}

type type_D24_ struct {
}

func (_this type_D24_) Default() interface{} {
	return Companion_D24_.Default()
}

func (_this type_D24_) String() string {
	return "main.D24"
}
func (_this D24) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D24{}

// End of datatype D24

// Definition of datatype D25
type D25 struct {
	Data_D25_
}

func (_this D25) Get_() Data_D25_ {
	return _this.Data_D25_
}

type Data_D25_ interface {
	isD25()
}

type CompanionStruct_D25_ struct {
}

var Companion_D25_ = CompanionStruct_D25_{}

type D25_DC64 struct {
	Cf106 _dafny.MultiSet
}

func (D25_DC64) isD25() {}

func (CompanionStruct_D25_) Create_DC64_(Cf106 _dafny.MultiSet) D25 {
	return D25{D25_DC64{Cf106}}
}

func (_this D25) Is_DC64() bool {
	_, ok := _this.Get_().(D25_DC64)
	return ok
}

type D25_DC65 struct {
	Cf107 _dafny.Int
	Cf108 _dafny.Array
	Cf109 _dafny.Int
}

func (D25_DC65) isD25() {}

func (CompanionStruct_D25_) Create_DC65_(Cf107 _dafny.Int, Cf108 _dafny.Array, Cf109 _dafny.Int) D25 {
	return D25{D25_DC65{Cf107, Cf108, Cf109}}
}

func (_this D25) Is_DC65() bool {
	_, ok := _this.Get_().(D25_DC65)
	return ok
}

type D25_DC63 struct {
	Cf105 _dafny.Map
}

func (D25_DC63) isD25() {}

func (CompanionStruct_D25_) Create_DC63_(Cf105 _dafny.Map) D25 {
	return D25{D25_DC63{Cf105}}
}

func (_this D25) Is_DC63() bool {
	_, ok := _this.Get_().(D25_DC63)
	return ok
}

type D25_DC66 struct {
	Cf110 D25
}

func (D25_DC66) isD25() {}

func (CompanionStruct_D25_) Create_DC66_(Cf110 D25) D25 {
	return D25{D25_DC66{Cf110}}
}

func (_this D25) Is_DC66() bool {
	_, ok := _this.Get_().(D25_DC66)
	return ok
}

func (CompanionStruct_D25_) Default() D25 {
	return Companion_D25_.Create_DC64_(_dafny.EmptyMultiSet)
}

func (_this D25) Dtor_cf106() _dafny.MultiSet {
	return _this.Get_().(D25_DC64).Cf106
}

func (_this D25) Dtor_cf107() _dafny.Int {
	return _this.Get_().(D25_DC65).Cf107
}

func (_this D25) Dtor_cf108() _dafny.Array {
	return _this.Get_().(D25_DC65).Cf108
}

func (_this D25) Dtor_cf109() _dafny.Int {
	return _this.Get_().(D25_DC65).Cf109
}

func (_this D25) Dtor_cf105() _dafny.Map {
	return _this.Get_().(D25_DC63).Cf105
}

func (_this D25) Dtor_cf110() D25 {
	return _this.Get_().(D25_DC66).Cf110
}

func (_this D25) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D25_DC64:
		{
			return "D25.DC64" + "(" + _dafny.String(data.Cf106) + ")"
		}
	case D25_DC65:
		{
			return "D25.DC65" + "(" + _dafny.String(data.Cf107) + ", " + _dafny.String(data.Cf108) + ", " + _dafny.String(data.Cf109) + ")"
		}
	case D25_DC63:
		{
			return "D25.DC63" + "(" + _dafny.String(data.Cf105) + ")"
		}
	case D25_DC66:
		{
			return "D25.DC66" + "(" + _dafny.String(data.Cf110) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D25) Equals(other D25) bool {
	switch data1 := _this.Get_().(type) {
	case D25_DC64:
		{
			data2, ok := other.Get_().(D25_DC64)
			return ok && data1.Cf106.Equals(data2.Cf106)
		}
	case D25_DC65:
		{
			data2, ok := other.Get_().(D25_DC65)
			return ok && data1.Cf107.Cmp(data2.Cf107) == 0 && data1.Cf108 == data2.Cf108 && data1.Cf109.Cmp(data2.Cf109) == 0
		}
	case D25_DC63:
		{
			data2, ok := other.Get_().(D25_DC63)
			return ok && data1.Cf105.Equals(data2.Cf105)
		}
	case D25_DC66:
		{
			data2, ok := other.Get_().(D25_DC66)
			return ok && data1.Cf110.Equals(data2.Cf110)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D25) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D25)
	return ok && _this.Equals(typed)
}

func Type_D25_() _dafny.TypeDescriptor {
	return type_D25_{}
}

type type_D25_ struct {
}

func (_this type_D25_) Default() interface{} {
	return Companion_D25_.Default()
}

func (_this type_D25_) String() string {
	return "main.D25"
}
func (_this D25) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D25{}

// End of datatype D25

// Definition of datatype D26
type D26 struct {
	Data_D26_
}

func (_this D26) Get_() Data_D26_ {
	return _this.Data_D26_
}

type Data_D26_ interface {
	isD26()
}

type CompanionStruct_D26_ struct {
}

var Companion_D26_ = CompanionStruct_D26_{}

type D26_DC68 struct {
	Cf112 bool
}

func (D26_DC68) isD26() {}

func (CompanionStruct_D26_) Create_DC68_(Cf112 bool) D26 {
	return D26{D26_DC68{Cf112}}
}

func (_this D26) Is_DC68() bool {
	_, ok := _this.Get_().(D26_DC68)
	return ok
}

type D26_DC69 struct {
	Cf113 _dafny.Sequence
	Cf114 bool
}

func (D26_DC69) isD26() {}

func (CompanionStruct_D26_) Create_DC69_(Cf113 _dafny.Sequence, Cf114 bool) D26 {
	return D26{D26_DC69{Cf113, Cf114}}
}

func (_this D26) Is_DC69() bool {
	_, ok := _this.Get_().(D26_DC69)
	return ok
}

type D26_DC67 struct {
	Cf111 _dafny.MultiSet
}

func (D26_DC67) isD26() {}

func (CompanionStruct_D26_) Create_DC67_(Cf111 _dafny.MultiSet) D26 {
	return D26{D26_DC67{Cf111}}
}

func (_this D26) Is_DC67() bool {
	_, ok := _this.Get_().(D26_DC67)
	return ok
}

func (CompanionStruct_D26_) Default() D26 {
	return Companion_D26_.Create_DC68_(false)
}

func (_this D26) Dtor_cf112() bool {
	return _this.Get_().(D26_DC68).Cf112
}

func (_this D26) Dtor_cf113() _dafny.Sequence {
	return _this.Get_().(D26_DC69).Cf113
}

func (_this D26) Dtor_cf114() bool {
	return _this.Get_().(D26_DC69).Cf114
}

func (_this D26) Dtor_cf111() _dafny.MultiSet {
	return _this.Get_().(D26_DC67).Cf111
}

func (_this D26) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D26_DC68:
		{
			return "D26.DC68" + "(" + _dafny.String(data.Cf112) + ")"
		}
	case D26_DC69:
		{
			return "D26.DC69" + "(" + data.Cf113.VerbatimString(true) + ", " + _dafny.String(data.Cf114) + ")"
		}
	case D26_DC67:
		{
			return "D26.DC67" + "(" + _dafny.String(data.Cf111) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D26) Equals(other D26) bool {
	switch data1 := _this.Get_().(type) {
	case D26_DC68:
		{
			data2, ok := other.Get_().(D26_DC68)
			return ok && data1.Cf112 == data2.Cf112
		}
	case D26_DC69:
		{
			data2, ok := other.Get_().(D26_DC69)
			return ok && data1.Cf113.Equals(data2.Cf113) && data1.Cf114 == data2.Cf114
		}
	case D26_DC67:
		{
			data2, ok := other.Get_().(D26_DC67)
			return ok && data1.Cf111.Equals(data2.Cf111)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D26) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D26)
	return ok && _this.Equals(typed)
}

func Type_D26_() _dafny.TypeDescriptor {
	return type_D26_{}
}

type type_D26_ struct {
}

func (_this type_D26_) Default() interface{} {
	return Companion_D26_.Default()
}

func (_this type_D26_) String() string {
	return "main.D26"
}
func (_this D26) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D26{}

// End of datatype D26

// Definition of datatype D27
type D27 struct {
	Data_D27_
}

func (_this D27) Get_() Data_D27_ {
	return _this.Data_D27_
}

type Data_D27_ interface {
	isD27()
}

type CompanionStruct_D27_ struct {
}

var Companion_D27_ = CompanionStruct_D27_{}

type D27_DC70 struct {
	Cf115 _dafny.MultiSet
}

func (D27_DC70) isD27() {}

func (CompanionStruct_D27_) Create_DC70_(Cf115 _dafny.MultiSet) D27 {
	return D27{D27_DC70{Cf115}}
}

func (_this D27) Is_DC70() bool {
	_, ok := _this.Get_().(D27_DC70)
	return ok
}

func (CompanionStruct_D27_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D27) Dtor_cf115() _dafny.MultiSet {
	return _this.Get_().(D27_DC70).Cf115
}

func (_this D27) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D27_DC70:
		{
			return "D27.DC70" + "(" + _dafny.String(data.Cf115) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D27) Equals(other D27) bool {
	switch data1 := _this.Get_().(type) {
	case D27_DC70:
		{
			data2, ok := other.Get_().(D27_DC70)
			return ok && data1.Cf115.Equals(data2.Cf115)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D27) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D27)
	return ok && _this.Equals(typed)
}

func Type_D27_() _dafny.TypeDescriptor {
	return type_D27_{}
}

type type_D27_ struct {
}

func (_this type_D27_) Default() interface{} {
	return Companion_D27_.Default()
}

func (_this type_D27_) String() string {
	return "main.D27"
}
func (_this D27) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D27{}

// End of datatype D27

// Definition of datatype D28
type D28 struct {
	Data_D28_
}

func (_this D28) Get_() Data_D28_ {
	return _this.Data_D28_
}

type Data_D28_ interface {
	isD28()
}

type CompanionStruct_D28_ struct {
}

var Companion_D28_ = CompanionStruct_D28_{}

type D28_DC72 struct {
	Cf117 _dafny.Int
	Cf118 _dafny.Int
	Cf119 bool
	Cf120 bool
}

func (D28_DC72) isD28() {}

func (CompanionStruct_D28_) Create_DC72_(Cf117 _dafny.Int, Cf118 _dafny.Int, Cf119 bool, Cf120 bool) D28 {
	return D28{D28_DC72{Cf117, Cf118, Cf119, Cf120}}
}

func (_this D28) Is_DC72() bool {
	_, ok := _this.Get_().(D28_DC72)
	return ok
}

type D28_DC71 struct {
	Cf116 _dafny.Map
}

func (D28_DC71) isD28() {}

func (CompanionStruct_D28_) Create_DC71_(Cf116 _dafny.Map) D28 {
	return D28{D28_DC71{Cf116}}
}

func (_this D28) Is_DC71() bool {
	_, ok := _this.Get_().(D28_DC71)
	return ok
}

func (CompanionStruct_D28_) Default() D28 {
	return Companion_D28_.Create_DC72_(_dafny.Zero, _dafny.Zero, false, false)
}

func (_this D28) Dtor_cf117() _dafny.Int {
	return _this.Get_().(D28_DC72).Cf117
}

func (_this D28) Dtor_cf118() _dafny.Int {
	return _this.Get_().(D28_DC72).Cf118
}

func (_this D28) Dtor_cf119() bool {
	return _this.Get_().(D28_DC72).Cf119
}

func (_this D28) Dtor_cf120() bool {
	return _this.Get_().(D28_DC72).Cf120
}

func (_this D28) Dtor_cf116() _dafny.Map {
	return _this.Get_().(D28_DC71).Cf116
}

func (_this D28) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D28_DC72:
		{
			return "D28.DC72" + "(" + _dafny.String(data.Cf117) + ", " + _dafny.String(data.Cf118) + ", " + _dafny.String(data.Cf119) + ", " + _dafny.String(data.Cf120) + ")"
		}
	case D28_DC71:
		{
			return "D28.DC71" + "(" + _dafny.String(data.Cf116) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D28) Equals(other D28) bool {
	switch data1 := _this.Get_().(type) {
	case D28_DC72:
		{
			data2, ok := other.Get_().(D28_DC72)
			return ok && data1.Cf117.Cmp(data2.Cf117) == 0 && data1.Cf118.Cmp(data2.Cf118) == 0 && data1.Cf119 == data2.Cf119 && data1.Cf120 == data2.Cf120
		}
	case D28_DC71:
		{
			data2, ok := other.Get_().(D28_DC71)
			return ok && data1.Cf116.Equals(data2.Cf116)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D28) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D28)
	return ok && _this.Equals(typed)
}

func Type_D28_() _dafny.TypeDescriptor {
	return type_D28_{}
}

type type_D28_ struct {
}

func (_this type_D28_) Default() interface{} {
	return Companion_D28_.Default()
}

func (_this type_D28_) String() string {
	return "main.D28"
}
func (_this D28) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D28{}

// End of datatype D28

// Definition of trait T0
type T0 interface {
	String() string
	Fm5(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 bool, globalState *GlobalState) bool
	M1(p0 _dafny.Int, globalState *GlobalState) (bool, bool)
	M2(p0 _dafny.Int, globalState *GlobalState) _dafny.Map
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
	F1() D0
	F1_set_(value D0)
	Fm6(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int
	F0() _dafny.Sequence
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
	dummy byte
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

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

func (_this *GlobalState) Ctor__() {
	{
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F8 _dafny.Sequence
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F8 = _dafny.EmptySeq
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

func (_this *C0) Ctor__(f8 _dafny.Sequence) {
	{
		(_this).F8 = f8
	}
}
func (_this *C0) Fm14(globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('w')
	}
}
func (_this *C0) Fm15(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(5))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg38 _dafny.Int) interface{} {
				return coer38(arg38)
			}
		}(func(_329_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		})), _this.F8)
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f1  D0
	_f0  _dafny.Sequence
	F9   _dafny.Int
	_f10 _dafny.Array
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f1 = Companion_D0_.Default()
	_this._f0 = _dafny.EmptySeq
	_this.F9 = _dafny.Zero
	_this._f10 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_}
}

var _ T1 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F1() D0 {
	return _this._f1
}
func (_this *C1) F1_set_(value D0) {
	_this._f1 = value
}
func (_this *C1) F0() _dafny.Sequence {
	return _this._f0
}
func (_this *C1) Ctor__(f9 _dafny.Int, f10 _dafny.Array, f0 _dafny.Sequence, f1 D0) {
	{
		(_this).F9 = f9
		(_this)._f10 = f10
		(_this)._f0 = f0
		(_this)._f1 = f1
	}
}
func (_this *C1) Fm6(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.MultiSetOf(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(603), _this.F9), _this.F9, _this.F9, ((func() _dafny.Set {
			var _coll33 = _dafny.NewBuilder()
			_ = _coll33
			for _iter33 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("drwqyvqsw")).Cardinality()))).Elements()); ; {
				_compr_33, _ok33 := _iter33()
				if !_ok33 {
					break
				}
				var _330_v0 _dafny.Int
				_330_v0 = interface{}(_compr_33).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("drwqyvqsw")).Cardinality()))).Contains(_330_v0) {
					_coll33.Add((_330_v0).Plus(_dafny.IntOfInt64(613)))
				}
			}
			return _coll33.ToSet()
		}()).Difference(_dafny.SetOf((_dafny.Zero).Minus(_this.F9), _this.F9, (_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(797))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}(func(_331_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})), _dafny.SeqOf(_dafny.CodePoint('h'), _dafny.CodePoint('c')))).Cardinality()))).Cardinality()
	}
}
func (_this *C1) Fm27(p0 _dafny.Map, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(!(false))), _dafny.SeqOf(true)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(false))))
	}
}
func (_this *C1) Fm28(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return _this.F9
	}
}
func (_this *C1) M13(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		(_this).F9 = _this.F9
		r0 = p1
		if p1 {
			var _332_v0 _dafny.Sequence
			_ = _332_v0
			_332_v0 = _dafny.UnicodeSeqOfUtf8Bytes("dmxvlol")
			var _333_v1 *C0
			_ = _333_v1
			var _nw50 *C0 = New_C0_()
			_ = _nw50
			_nw50.Ctor__(_332_v0)
			_333_v1 = _nw50
			var _334_v2 *C0
			_ = _334_v2
			var _nw51 *C0 = New_C0_()
			_ = _nw51
			_nw51.Ctor__(_dafny.Companion_Sequence_.Concatenate(_332_v0, _333_v1.F8))
			_334_v2 = _nw51
			var _rhs31 bool = true
			_ = _rhs31
			var _rhs32 bool = p1
			_ = _rhs32
			var _rhs33 _dafny.Int = Companion_Default___.SafeDivisionInt((_this.F9).Minus(_this.F9), (_this.F9).Minus(_this.F9))
			_ = _rhs33
			var _lhs16 *C1 = _this
			_ = _lhs16
			r0 = _rhs31
			r0 = _rhs32
			_lhs16.F9 = _rhs33
			var _335_v3 _dafny.MultiSet
			_ = _335_v3
			_335_v3 = _dafny.MultiSetOf(p1, p1)
			var _336_v4 _dafny.Sequence
			_ = _336_v4
			_336_v4 = _dafny.SeqOf(p1)
			var _337_v5 D7
			_ = _337_v5
			_337_v5 = Companion_D7_.Create_DC17_(_336_v4)
			var _338_v6 _dafny.MultiSet
			_ = _338_v6
			_338_v6 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Equal(_332_v0, _334_v2.F8), p1, p1, !((_dafny.MultiSetFromSeq((_337_v5).Dtor_cf34())).IsProperSubsetOf(_335_v3)))
			_335_v3 = ((_dafny.MultiSetOf(p1)).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(p1, p1)))).Difference(_335_v3)
			var _339_v7 *C0
			_ = _339_v7
			var _nw52 *C0 = New_C0_()
			_ = _nw52
			_nw52.Ctor__(_dafny.Companion_Sequence_.Concatenate(_332_v0, _333_v1.F8))
			_339_v7 = _nw52
		} else {
			var _340_v8 _dafny.Array
			_ = _340_v8
			var _nw53 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw53
			_340_v8 = _nw53
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_340_v8), 0))
			_ = _index30
			(_340_v8).ArraySet1(p1, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_340_v8), 0))
			_ = _index31
			(_340_v8).ArraySet1(p1, (_index31).Int())
			var _341_v9 _dafny.Array
			_ = _341_v9
			var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw54
			_341_v9 = _nw54
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_341_v9), 0))
			_ = _index32
			(_341_v9).ArraySet1(_this.F9, (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_341_v9), 0))
			_ = _index33
			(_341_v9).ArraySet1(Companion_Default___.SafeModuloInt(_this.F9, _this.F9), (_index33).Int())
			_340_v8 = _340_v8
			var _342_v10 _dafny.Array
			_ = _342_v10
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_3
			var _nw55 _dafny.Array
			_ = _nw55
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw55 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.CodePoint = (func(_343_p0 _dafny.CodePoint, _344_p1 bool) func(_dafny.Int) _dafny.CodePoint {
					return func(_345_i0 _dafny.Int) _dafny.CodePoint {
						return (Companion_D11_.Create_DC25_(_343_p0, (_dafny.MultiSetOf(Companion_D6_.Create_DC14_(_dafny.UnicodeSeqOfUtf8Bytes("pqjlcv")), Companion_D6_.Create_DC14_(_dafny.UnicodeSeqOfUtf8Bytes("lcnl")))).Cardinality(), _344_p1)).Dtor_cf41()
					}
				})(p0, p1)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw55 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw55).ArraySet1CodePoint(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw55).ArraySet1CodePoint(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_342_v10 = _nw55
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_342_v10), 0))
			_ = _index34
			(_342_v10).ArraySet1CodePoint(Companion_Default___.Fm29((_340_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_340_v8), 0))).Int()).(bool), globalState), (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_342_v10), 0))
			_ = _index35
			var _rhs34 _dafny.CodePoint = p0
			_ = _rhs34
			var _rhs35 bool = (_340_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_340_v8), 0))).Int()).(bool)
			_ = _rhs35
			var _lhs17 _dafny.Array = _342_v10
			_ = _lhs17
			var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_342_v10), 0))
			_ = _lhs18
			(_lhs17).ArraySet1CodePoint(_rhs34, (_lhs18).Int())
			r0 = _rhs35
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_341_v9), 0))
			_ = _index36
			(_341_v9).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.IntOfInt64(-65)).Plus(_this.F9), ((_341_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_341_v9), 0))).Int()).(_dafny.Int)).Times((_341_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_341_v9), 0))).Int()).(_dafny.Int)))), (_index36).Int())
		}
		var _346_v11 _dafny.Array
		_ = _346_v11
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_4
		var _nw56 _dafny.Array
		_ = _nw56
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw56 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.CodePoint = (func(_347_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_348_i2 _dafny.Int) _dafny.CodePoint {
					return _347_p0
				}
			})(p0)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw56 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw56).ArraySet1CodePoint(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw56).ArraySet1CodePoint(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_346_v11 = _nw56
		var _349_v12 _dafny.Sequence
		_ = _349_v12
		_349_v12 = _dafny.SeqOf(_346_v11, _346_v11, _346_v11)
		var _hi0 _dafny.Int = _dafny.IntOfUint32((_349_v12).Cardinality())
		_ = _hi0
		for _350_i1 := _this.F9; _350_i1.Cmp(_hi0) < 0; _350_i1 = _350_i1.Plus(_dafny.One) {
			var _351_v13 _dafny.Array
			_ = _351_v13
			var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
			_ = _nw57
			_351_v13 = _nw57
			var _352_v14 D2
			_ = _352_v14
			_352_v14 = Companion_D2_.Create_DC4_(_this.F9, p1, _351_v13)
			var _353_v15 _dafny.MultiSet
			_ = _353_v15
			_353_v15 = _dafny.MultiSetOf((_this.F9).Times((_352_v14).Dtor_cf5()))
			_353_v15 = _353_v15
			var _354_v16 _dafny.Array
			_ = _354_v16
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_5
			var _nw58 _dafny.Array
			_ = _nw58
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw58 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) bool = (func(_355_p1 bool) func(_dafny.Int) bool {
					return func(_356_i3 _dafny.Int) bool {
						return (_355_p1) && (_355_p1)
					}
				})(p1)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw58 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw58).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw58).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_354_v16 = _nw58
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_354_v16), 0))
			_ = _index37
			(_354_v16).ArraySet1((p1) && (p1), (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_354_v16), 0))
			_ = _index38
			(_354_v16).ArraySet1(p1, (_index38).Int())
			var _357_v17 _dafny.Array
			_ = _357_v17
			var _nw59 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(17))
			_ = _nw59
			_357_v17 = _nw59
			var _358_v18 _dafny.Sequence
			_ = _358_v18
			_358_v18 = _dafny.SeqOf(p1)
			var _359_v19 D6
			_ = _359_v19
			_359_v19 = Companion_D6_.Create_DC15_(_350_i1, _this.F9, Companion_Default___.Fm30(_dafny.IntOfUint32((_358_v18).Cardinality()), (_354_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_354_v16), 0))).Int()).(bool), (_354_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_354_v16), 0))).Int()).(bool), globalState))
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_357_v17), 0))
			_ = _index39
			(_357_v17).ArraySet1(_359_v19, (_index39).Int())
			var _360_v20 _dafny.Set
			_ = _360_v20
			_360_v20 = _dafny.SetOf(_this.F9)
			var _361_v21 _dafny.Array
			_ = _361_v21
			var _nwElement0_10 _dafny.Set = _dafny.SetOf(_this.F9)
			_ = _nwElement0_10
			var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(4))
			_ = _nw60
			(_nw60).ArraySet1(_nwElement0_10, 0)
			(_nw60).ArraySet1((_360_v20).Difference(_360_v20), 1)
			(_nw60).ArraySet1(_360_v20, 2)
			(_nw60).ArraySet1(_360_v20, 3)
			_361_v21 = _nw60
			var _362_v22 _dafny.Array
			_ = _362_v22
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_6
			var _nw61 _dafny.Array
			_ = _nw61
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw61 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Int = (func(_363_i1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_364_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_364_i4, _363_i1)
					}
				})(_350_i1)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw61 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw61).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw61).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_362_v22 = _nw61
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_362_v22), 0))
			_ = _index40
			(_362_v22).ArraySet1(Companion_Default___.SafeModuloInt(_this.F9, _this.F9), (_index40).Int())
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_357_v17), 0))
			_ = _index41
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_354_v16), 0))
			_ = _index42
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_362_v22), 0))
			_ = _index43
			var _rhs36 D6 = _359_v19
			_ = _rhs36
			var _rhs37 _dafny.Array = _361_v21
			_ = _rhs37
			var _rhs38 bool = (_354_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_354_v16), 0))).Int()).(bool)
			_ = _rhs38
			var _rhs39 _dafny.Int = (func() _dafny.Int {
				if p1 {
					return (_dafny.Zero).Minus(_350_i1)
				}
				return (_this.F9).Minus(_this.F9)
			})()
			_ = _rhs39
			var _lhs19 _dafny.Array = _357_v17
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_357_v17), 0))
			_ = _lhs20
			var _lhs21 _dafny.Array = _354_v16
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_354_v16), 0))
			_ = _lhs22
			var _lhs23 _dafny.Array = _362_v22
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_362_v22), 0))
			_ = _lhs24
			(_lhs19).ArraySet1(_rhs36, (_lhs20).Int())
			_361_v21 = _rhs37
			(_lhs21).ArraySet1(_rhs38, (_lhs22).Int())
			(_lhs23).ArraySet1(_rhs39, (_lhs24).Int())
			var _365_v23 _dafny.Sequence
			_ = _365_v23
			_365_v23 = _dafny.UnicodeSeqOfUtf8Bytes("rarjd")
			var _366_v24 D6
			_ = _366_v24
			_366_v24 = Companion_D6_.Create_DC14_(_dafny.Companion_Sequence_.Concatenate(_365_v23, _dafny.UnicodeSeqOfUtf8Bytes("jmypydeh")))
			var _pat_let_tv8 = _365_v23
			_ = _pat_let_tv8
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_362_v22), 0))
			_ = _index44
			var _rhs40 D6 = func(_pat_let10_0 D6) D6 {
				return func(_367_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let11_0 _dafny.Sequence) D6 {
						return func(_368_dt__update_hcf29_h0 _dafny.Sequence) D6 {
							return Companion_D6_.Create_DC14_(_368_dt__update_hcf29_h0)
						}(_pat_let11_0)
					}(_pat_let_tv8)
				}(_pat_let10_0)
			}(Companion_D6_.Create_DC14_(_365_v23))
			_ = _rhs40
			var _rhs41 _dafny.Sequence = _dafny.SeqOf((_354_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_354_v16), 0))).Int()).(bool), p1, p1)
			_ = _rhs41
			var _rhs42 _dafny.Int = _this.F9
			_ = _rhs42
			var _lhs25 _dafny.Array = _362_v22
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_362_v22), 0))
			_ = _lhs26
			_366_v24 = _rhs40
			_358_v18 = _rhs41
			(_lhs25).ArraySet1(_rhs42, (_lhs26).Int())
		}
		var _hi1 _dafny.Int = _this.F9
		_ = _hi1
		for _369_i5 := _this.F9; _369_i5.Cmp(_hi1) < 0; _369_i5 = _369_i5.Plus(_dafny.One) {
			var _370_v25 _dafny.Array
			_ = _370_v25
			var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
			_ = _nw62
			_370_v25 = _nw62
			var _371_v26 _dafny.Sequence
			_ = _371_v26
			_371_v26 = _dafny.SeqOf(p1, p1)
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_370_v25), 0))
			_ = _index45
			(_370_v25).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_371_v26, _dafny.Companion_Sequence_.Update(_371_v26, (Companion_Default___.SafeIndex(_this.F9, _dafny.IntOfUint32((_371_v26).Cardinality()))).Uint32(), p1)), (_index45).Int())
			var _372_v27 D7
			_ = _372_v27
			_372_v27 = Companion_D7_.Create_DC17_(_371_v26)
			var _373_v28 _dafny.Map
			_ = _373_v28
			_373_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _371_v26)
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_370_v25), 0))
			_ = _index46
			(_370_v25).ArraySet1(_dafny.Companion_Sequence_.Update(((func() D7 {
				if p1 {
					return _372_v27
				}
				return Companion_D7_.Create_DC17_((func() _dafny.Sequence {
					if (_373_v28).Contains(p1) {
						return (_373_v28).Get(p1).(_dafny.Sequence)
					}
					return _371_v26
				})())
			})()).Dtor_cf34(), (Companion_Default___.SafeIndex(_this.F9, _dafny.IntOfUint32((((func() D7 {
				if p1 {
					return _372_v27
				}
				return Companion_D7_.Create_DC17_((func() _dafny.Sequence {
					if (_373_v28).Contains(p1) {
						return (_373_v28).Get(p1).(_dafny.Sequence)
					}
					return _371_v26
				})())
			})()).Dtor_cf34()).Cardinality()))).Uint32(), (Companion_D12_.Create_DC29_(p1, _369_i5, p1, _dafny.IntOfInt64(897), p1)).Dtor_cf50()), (_index46).Int())
			var _374_v29 _dafny.Sequence
			_ = _374_v29
			_374_v29 = _dafny.UnicodeSeqOfUtf8Bytes("nyg")
			_374_v29 = _dafny.Companion_Sequence_.Concatenate(_374_v29, _374_v29)
			(_this).F9 = Companion_Default___.SafeDivisionInt(_this.F9, _this.F9)
			var _375_v30 _dafny.Map
			_ = _375_v30
			_375_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_374_v29).Cardinality()), p1)
			var _376_v31 _dafny.Array
			_ = _376_v31
			var _nw63 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
			_ = _nw63
			_376_v31 = _nw63
			var _377_v32 _dafny.MultiSet
			_ = _377_v32
			_377_v32 = _dafny.MultiSetOf(_376_v31, _376_v31)
			var _378_v33 _dafny.MultiSet
			_ = _378_v33
			_378_v33 = _dafny.MultiSetOf(p0)
			var _379_v34 _dafny.MultiSet
			_ = _379_v34
			_379_v34 = _dafny.MultiSetOf(_369_i5, _369_i5, (_378_v33).Cardinality())
			var _380_v35 _dafny.Set
			_ = _380_v35
			_380_v35 = _dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_369_i5)), _379_v34, _dafny.MultiSetFromSeq((_this).F0()), _dafny.MultiSetOf(_dafny.IntOfUint32((Companion_Default___.Fm31(_this.F9, p1, _this.F9, globalState)).Cardinality())))
			var _381_v36 _dafny.Array
			_ = _381_v36
			var _nwElement0_11 _dafny.Int = ((_dafny.Zero).Minus(_369_i5)).Minus(_369_i5)
			_ = _nwElement0_11
			var _nw64 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(22))
			_ = _nw64
			(_nw64).ArraySet1(_nwElement0_11, 0)
			(_nw64).ArraySet1((_375_v30).Cardinality(), 1)
			(_nw64).ArraySet1((_this.F9).Plus(_369_i5), 2)
			(_nw64).ArraySet1(_369_i5, 3)
			(_nw64).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_369_i5), _369_i5), 4)
			(_nw64).ArraySet1(Companion_Default___.SafeDivisionInt(_this.F9, _369_i5), 5)
			(_nw64).ArraySet1(_369_i5, 6)
			(_nw64).ArraySet1((_this.F9).Times(_this.F9), 7)
			(_nw64).ArraySet1(_dafny.IntOfUint32(((_this).F0()).Cardinality()), 8)
			(_nw64).ArraySet1((_dafny.Zero).Minus((_this).Fm28(_dafny.IntOfInt64(-80), p1, globalState)), 9)
			(_nw64).ArraySet1((func() _dafny.Int {
				if (_377_v32).Contains(_376_v31) {
					return (_377_v32).Multiplicity(_376_v31)
				}
				return (_380_v35).Cardinality()
			})(), 10)
			(_nw64).ArraySet1(_this.F9, 11)
			(_nw64).ArraySet1(_369_i5, 12)
			(_nw64).ArraySet1(_369_i5, 13)
			(_nw64).ArraySet1(_this.F9, 14)
			(_nw64).ArraySet1(_dafny.IntOfUint32(((_this).F0()).Cardinality()), 15)
			(_nw64).ArraySet1(_369_i5, 16)
			(_nw64).ArraySet1((_369_i5).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm31(_this.F9, p1, (_dafny.Zero).Minus(_dafny.IntOfUint32((_374_v29).Cardinality())), globalState)).Cardinality()))), 17)
			(_nw64).ArraySet1((_this.F9).Plus(_dafny.IntOfInt64(345)), 18)
			(_nw64).ArraySet1(_369_i5, 19)
			(_nw64).ArraySet1(_this.F9, 20)
			(_nw64).ArraySet1(_this.F9, 21)
			_381_v36 = _nw64
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_381_v36), 0))
			_ = _index47
			(_381_v36).ArraySet1(_dafny.IntOfInt64(882), (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_381_v36), 0))
			_ = _index48
			(_381_v36).ArraySet1((func() _dafny.Int {
				if (_377_v32).Contains(_376_v31) {
					return (_377_v32).Multiplicity(_376_v31)
				}
				return _this.F9
			})(), (_index48).Int())
		}
		var _382_v37 D12
		_ = _382_v37
		_382_v37 = Companion_D12_.Create_DC29_(p1, (_this).Fm28(_this.F9, !(!(p1)), globalState), p1, _this.F9, !(p1))
		var _383_v38 _dafny.Map
		_ = _383_v38
		_383_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_382_v37).Dtor_cf49(), _this.F9)
		var _384_v39 _dafny.Map
		_ = _384_v39
		_384_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9, p1)
		var _385_v40 D11
		_ = _385_v40
		_385_v40 = Companion_D11_.Create_DC25_(p0, _dafny.IntOfInt64(390), (func() bool {
			if (_384_v39).Contains(_this.F9) {
				return (_384_v39).Get(_this.F9).(bool)
			}
			return p1
		})())
		var _386_v41 _dafny.Sequence
		_ = _386_v41
		_386_v41 = _dafny.UnicodeSeqOfUtf8Bytes("bhh")
		var _hi2 _dafny.Int = _dafny.IntOfUint32((_386_v41).Cardinality())
		_ = _hi2
		for _387_i6 := (_dafny.Zero).Minus((func() _dafny.Int {
			if (_383_v38).Contains((_385_v40).Dtor_cf42()) {
				return (_383_v38).Get((_385_v40).Dtor_cf42()).(_dafny.Int)
			}
			return (_dafny.MultiSetOf(p0)).Cardinality()
		})()); _387_i6.Cmp(_hi2) < 0; _387_i6 = _387_i6.Plus(_dafny.One) {
			r0 = (func() bool {
				if p1 {
					return !(p1)
				}
				return p1
			})()
			var _388_v42 _dafny.MultiSet
			_ = _388_v42
			_388_v42 = _dafny.MultiSetOf(_387_i6, _387_i6, (_this).Fm6(_this.F9, _this.F9, globalState), _387_i6, _this.F9)
			(_this).F9 = Companion_Default___.SafeModuloInt(((_388_v42).Cardinality()).Plus(_387_i6), ((_this).Fm6(_this.F9, _this.F9, globalState)).Plus(_this.F9))
			var _389_v43 _dafny.CodePoint
			_ = _389_v43
			_389_v43 = _dafny.CodePoint('q')
			_389_v43 = _389_v43
			var _390_v44 _dafny.Array
			_ = _390_v44
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_7
			var _nw65 _dafny.Array
			_ = _nw65
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw65 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Sequence = func(_391_i7 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("htrgdatxk")
				}
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw65 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw65).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw65).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_390_v44 = _nw65
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_390_v44), 0))
			_ = _index49
			(_390_v44).ArraySet1(_386_v41, (_index49).Int())
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_390_v44), 0))
			_ = _index50
			(_390_v44).ArraySet1((func() _dafny.Sequence {
				if p1 {
					return _386_v41
				}
				return _386_v41
			})(), (_index50).Int())
		}
		r0 = (func() bool {
			if _dafny.Companion_Sequence_.IsPrefixOf(_386_v41, _dafny.UnicodeSeqOfUtf8Bytes("ly")) {
				return !_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm31(_this.F9, p1, (_dafny.MultiSetOf(p0, p0, p0)).Cardinality(), globalState), p0)
			}
			return p1
		})()
		return r0
	}
}
func (_this *C1) F10() _dafny.Array {
	{
		return _this._f10
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	dummy byte
}

func New_C2_() *C2 {
	_this := C2{}

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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) Ctor__() {
	{
	}
}
func (_this *C2) Fm5(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 bool, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C2) Fm41(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return _dafny.MultiSetOf(false)
	}
}
func (_this *C2) Fm42(p0 bool, globalState *GlobalState) bool {
	{
		return ((func() _dafny.Map {
			var _coll34 = _dafny.NewMapBuilder()
			_ = _coll34
			for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-558), _dafny.IntOfInt64(560))); ; {
				_compr_34, _ok34 := _iter34()
				if !_ok34 {
					break
				}
				var _392_v0 _dafny.Int
				_392_v0 = interface{}(_compr_34).(_dafny.Int)
				if ((_dafny.IntOfInt64(-558)).Cmp(_392_v0) <= 0) && ((_392_v0).Cmp(_dafny.IntOfInt64(560)) < 0) {
					_coll34.Add(Companion_Default___.SafeDivisionInt(_392_v0, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SeqOf(Companion_D2_.Create_DC5_(_dafny.IntOfInt64(-832), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-101))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg40 _dafny.Int) interface{} {
							return coer40(arg40)
						}
					}(func(_393_i0 _dafny.Int) _dafny.Int {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.CodePoint('p'))).Cardinality()
					})), false), Companion_D2_.Create_DC5_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tam")).Cardinality()), _dafny.CodePoint('k'))).Cardinality(), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, !(false), !(true))).Cardinality()), (_dafny.SetOf(true)).Cardinality()), true)))
				}
			}
			return _coll34.ToMap()
		}()).Cardinality()).Cmp(_dafny.IntOfInt64(281)) < 0
	}
}
func (_this *C2) M1(p0 _dafny.Int, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _394_v0 _dafny.MultiSet
		_ = _394_v0
		_394_v0 = _dafny.MultiSetOf(p0, p0)
		var _395_v1 _dafny.Map
		_ = _395_v1
		_395_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_394_v0).Cardinality(), p0)
		var _396_v2 _dafny.Array
		_ = _396_v2
		var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.One)
		_ = _nw66
		_396_v2 = _nw66
		var _397_v3 _dafny.Sequence
		_ = _397_v3
		_397_v3 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kpmro")).Cardinality()))
		var _398_v4 T1
		_ = _398_v4
		var _nw67 *C1 = New_C1_()
		_ = _nw67
		_nw67.Ctor__(p0, _396_v2, _397_v3, Companion_D0_.Create_DC1_(p0))
		_398_v4 = _nw67
		var _399_v5 _dafny.Map
		_ = _399_v5
		_399_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_398_v4, p0)
		var _400_v6 *C0
		_ = _400_v6
		var _nw68 *C0 = New_C0_()
		_ = _nw68
		_nw68.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("us"))
		_400_v6 = _nw68
		var _401_v7 _dafny.Map
		_ = _401_v7
		_401_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _400_v6)
		var _402_v8 _dafny.Sequence
		_ = _402_v8
		_402_v8 = _dafny.SeqOf(_395_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_399_v5).Cardinality(), (_401_v7).Cardinality()))
		var _403_v9 _dafny.Int
		_ = _403_v9
		_403_v9 = _dafny.IntOfInt64(156)
		var _rhs43 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_402_v8, _402_v8)
		_ = _rhs43
		var _rhs44 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-425))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg41 _dafny.Int) interface{} {
				return coer41(arg41)
			}
		}((func(_404_v9 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_405_i0 _dafny.Int) _dafny.Int {
				return _404_v9
			}
		})(_403_v9)))).Cardinality())
		_ = _rhs44
		var _rhs45 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("fvuiwig")
		_ = _rhs45
		var _lhs27 *C0 = _400_v6
		_ = _lhs27
		_402_v8 = _rhs43
		_403_v9 = _rhs44
		_lhs27.F8 = _rhs45
		var _406_v10 _dafny.Array
		_ = _406_v10
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_8
		var _nw69 _dafny.Array
		_ = _nw69
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw69 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) bool = func(_407_i2 _dafny.Int) bool {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.MultiSetOf(false))).Contains(false)
			}
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw69 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw69).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw69).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_406_v10 = _nw69
		for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_406_v10), 0))); ; {
			_guard_loop_0, _ok35 := _iter35()
			if !_ok35 {
				break
			}
			var _408_i1 _dafny.Int
			_408_i1 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_408_i1).Sign() != -1) && ((_408_i1).Cmp(_dafny.ArrayLen((_406_v10), 0)) < 0)) {
				(_406_v10).ArraySet1((_403_v9).Cmp(((_398_v4).F0()).Select((Companion_Default___.SafeIndex(_403_v9, _dafny.IntOfUint32(((_398_v4).F0()).Cardinality()))).Uint32()).(_dafny.Int)) == 0, (_408_i1).Int())
			}
		}
		var _409_v11 bool
		_ = _409_v11
		_409_v11 = true
		var _410_v12 _dafny.Map
		_ = _410_v12
		_410_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_409_v11, _403_v9)
		var _411_v13 _dafny.Sequence
		_ = _411_v13
		_411_v13 = _dafny.SeqOf(_400_v6.F8, _400_v6.F8, _400_v6.F8)
		var _412_v14 _dafny.Map
		_ = _412_v14
		_412_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_410_v12, (_411_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.IntOfUint32((_411_v13).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _413_v16 _dafny.Sequence
		_ = _413_v16
		_413_v16 = _dafny.SeqOf(_410_v12, _410_v12)
		var _414_v17 _dafny.Int
		_ = _414_v17
		var _out19 _dafny.Int
		_ = _out19
		_out19 = (_this).M14(_409_v11, _409_v11, (_412_v14).Merge(func() _dafny.Map {
			var _coll35 = _dafny.NewMapBuilder()
			_ = _coll35
			for _iter36 := _dafny.Iterate((_413_v16).Elements()); ; {
				_compr_35, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _415_v15 _dafny.Map
				_415_v15 = interface{}(_compr_35).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_413_v16, _415_v15) {
					_coll35.Add(_415_v15, _400_v6.F8)
				}
			}
			return _coll35.ToMap()
		}()), p0, globalState)
		_414_v17 = _out19
		r0 = _409_v11
		var _416_v18 _dafny.CodePoint
		_ = _416_v18
		_416_v18 = _dafny.CodePoint('x')
		var _417_i3 _dafny.Int
		_ = _417_i3
		_417_i3 = _dafny.Zero
		{
			for (func() _dafny.Set {
				var _coll37 = _dafny.NewBuilder()
				_ = _coll37
				for _iter38 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v9, _416_v18)).Keys().Elements()); ; {
					_compr_37, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _427_v19 _dafny.Int
					_427_v19 = interface{}(_compr_37).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v9, _416_v18)).Contains(_427_v19) {
						_coll37.Add(Companion_Default___.SafeModuloInt(_427_v19, (_dafny.Zero).Minus(_dafny.IntOfInt64(-404))))
					}
				}
				return _coll37.ToSet()
			}()).IsSubsetOf(_dafny.SetOf(p0)) {
				{
					if (_417_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_417_i3 = (_417_i3).Plus(_dafny.One)
					var _418_v20 D6
					_ = _418_v20
					_418_v20 = Companion_D6_.Create_DC15_((_398_v4).Fm6(_414_v17, _414_v17, globalState), p0, _409_v11)
					var _419_v21 D6
					_ = _419_v21
					_419_v21 = Companion_D6_.Create_DC16_(_418_v20)
					var _420_v22 D6
					_ = _420_v22
					_420_v22 = Companion_D6_.Create_DC16_(_418_v20)
					_420_v22 = _420_v22
					r1 = _409_v11
					var _421_v24 _dafny.Map
					_ = _421_v24
					_421_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_400_v6.F8, _416_v18)
					var _422_v25 D11
					_ = _422_v25
					_422_v25 = Companion_D11_.Create_DC24_(_421_v24)
					var _423_v26 D11
					_ = _423_v26
					_423_v26 = Companion_D11_.Create_DC26_(_422_v25)
					var _424_v27 _dafny.Sequence
					_ = _424_v27
					_424_v27 = _dafny.SeqOf(_409_v11, _409_v11)
					var _425_v28 D14
					_ = _425_v28
					_425_v28 = Companion_D14_.Create_DC34_(Companion_Default___.Fm43(_423_v26, _409_v11, _dafny.IntOfUint32((_424_v27).Cardinality()), globalState))
					_409_v11 = !((func() _dafny.Map {
						var _coll36 = _dafny.NewMapBuilder()
						_ = _coll36
						for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(442), _dafny.IntOfInt64(-928))); ; {
							_compr_36, _ok37 := _iter37()
							if !_ok37 {
								break
							}
							var _426_v23 _dafny.Int
							_426_v23 = interface{}(_compr_36).(_dafny.Int)
							if ((_dafny.IntOfInt64(442)).Cmp(_426_v23) <= 0) && ((_426_v23).Cmp(_dafny.IntOfInt64(-928)) < 0) {
								_coll36.Add((_426_v23).Plus(_414_v17), _dafny.UnicodeSeqOfUtf8Bytes("qkbtfupqy"))
							}
						}
						return _coll36.ToMap()
					}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _400_v6.F8))).Equals((_425_v28).Dtor_cf62())
					var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_406_v10), 0))
					_ = _index51
					(_406_v10).ArraySet1(_409_v11, (_index51).Int())
					var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_406_v10), 0))
					_ = _index52
					(_406_v10).ArraySet1(_409_v11, (_index52).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		r1 = _409_v11
		r0 = _409_v11
		r1 = _409_v11
		return r0, r1
	}
}
func (_this *C2) M2(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _428_v0 _dafny.Sequence
		_ = _428_v0
		_428_v0 = _dafny.UnicodeSeqOfUtf8Bytes("jgwejo")
		_428_v0 = _dafny.Companion_Sequence_.Concatenate(_428_v0, _428_v0)
		var _429_i0 _dafny.Int
		_ = _429_i0
		_429_i0 = _dafny.Zero
		{
			for true {
				{
					if (_429_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_429_i0 = (_429_i0).Plus(_dafny.One)
					var _430_v1 _dafny.Array
					_ = _430_v1
					var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
					_ = _nw70
					_430_v1 = _nw70
					var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_430_v1), 0))
					_ = _index53
					(_430_v1).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(54))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg42 _dafny.Int) interface{} {
							return coer42(arg42)
						}
					}(func(_431_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('i')
					})), (_index53).Int())
					var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_430_v1), 0))
					_ = _index54
					(_430_v1).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("pekucbt"), (_index54).Int())
					var _432_v2 bool
					_ = _432_v2
					_432_v2 = true
					var _433_v3 D0
					_ = _433_v3
					_433_v3 = Companion_D0_.Create_DC1_(p0)
					var _434_v4 _dafny.Sequence
					_ = _434_v4
					_434_v4 = _dafny.SeqOf(_432_v2)
					var _435_v5 _dafny.Map
					_ = _435_v5
					_435_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_433_v3, _dafny.IntOfUint32((_434_v4).Cardinality()))
					var _436_v6 _dafny.Sequence
					_ = _436_v6
					_436_v6 = _dafny.SeqOf(_435_v5)
					var _437_v7 _dafny.Sequence
					_ = _437_v7
					_437_v7 = _dafny.SeqOf(p0, p0, p0)
					var _438_v8 _dafny.Map
					_ = _438_v8
					_438_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_437_v7).Cardinality()), _432_v2)
					var _439_v9 _dafny.Map
					_ = _439_v9
					_439_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_436_v6)).Cardinality()), ((func() bool {
						if (_438_v8).Contains(_dafny.IntOfInt64(432)) {
							return (_438_v8).Get(_dafny.IntOfInt64(432)).(bool)
						}
						return _432_v2
					})()) && (_432_v2))
					var _440_v10 _dafny.Set
					_ = _440_v10
					_440_v10 = _dafny.SetOf(p0)
					_432_v2 = (func() bool {
						if (_438_v8).Contains(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_440_v10).Cardinality(), false)).Merge(_438_v8)).Cardinality()) {
							return (_438_v8).Get(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_440_v10).Cardinality(), false)).Merge(_438_v8)).Cardinality()).(bool)
						}
						return (_432_v2) == (_432_v2)
					})()
					var _441_v11 _dafny.Array
					_ = _441_v11
					var _nw71 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
					_ = _nw71
					_441_v11 = _nw71
					var _442_v12 _dafny.Set
					_ = _442_v12
					_442_v12 = _dafny.SetOf(_441_v11)
					var _443_v13 _dafny.Set
					_ = _443_v13
					_443_v13 = _442_v12
					var _444_v14 _dafny.Map
					_ = _444_v14
					_444_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_432_v2, _443_v13)
					_444_v14 = (_444_v14).Update(_432_v2, _442_v12)
					var _445_v15 _dafny.Array
					_ = _445_v15
					var _len0_9 _dafny.Int = _dafny.IntOfInt64(24)
					_ = _len0_9
					var _nw72 _dafny.Array
					_ = _nw72
					if _len0_9.Cmp(_dafny.Zero) == 0 {
						_nw72 = _dafny.NewArray(_len0_9)
					} else {
						var _init9 func(_dafny.Int) _dafny.Int = (func(_446_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_447_i2 _dafny.Int) _dafny.Int {
								return (_447_i2).Plus(_446_p0)
							}
						})(p0)
						_ = _init9
						var _element0_9 = _init9(_dafny.Zero)
						_ = _element0_9
						_nw72 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
						(_nw72).ArraySet1(_element0_9, 0)
						var _nativeLen0_9 = (_len0_9).Int()
						_ = _nativeLen0_9
						for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
							(_nw72).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
						}
					}
					_445_v15 = _nw72
					var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_445_v15), 0))
					_ = _index55
					(_445_v15).ArraySet1(p0, (_index55).Int())
					var _448_v16 _dafny.Array
					_ = _448_v16
					var _nw73 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
					_ = _nw73
					_448_v16 = _nw73
					var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_448_v16), 0))
					_ = _index56
					(_448_v16).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_432_v2, _432_v2), (_index56).Int())
					var _449_v17 _dafny.Map
					_ = _449_v17
					_449_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_430_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_430_v1), 0))).Int()).(_dafny.Sequence))
					var _450_v18 _dafny.CodePoint
					_ = _450_v18
					_450_v18 = _dafny.CodePoint('i')
					var _451_v19 _dafny.Map
					_ = _451_v19
					_451_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_434_v4).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_434_v4).Cardinality()))).Uint32()).(bool)), _432_v2)
					var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_445_v15), 0))
					_ = _index57
					var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_448_v16), 0))
					_ = _index58
					var _rhs46 _dafny.Int = Companion_Default___.Fm44(!((p0).Cmp((_dafny.SetOf(_432_v2)).Cardinality()) < 0), globalState)
					_ = _rhs46
					var _rhs47 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
						if _432_v2 {
							return (func() _dafny.Sequence {
								if (_449_v17).Contains(p0) {
									return (_449_v17).Get(p0).(_dafny.Sequence)
								}
								return _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(120))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg43 _dafny.Int) interface{} {
										return coer43(arg43)
									}
								}(func(_452_i3 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('s')
								})), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(120))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg44 _dafny.Int) interface{} {
										return coer44(arg44)
									}
								}(func(_453_i3 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('s')
								}))).Cardinality()))).Uint32(), _450_v18)
							})()
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("fhkvhgy")
					})(), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((func() _dafny.Sequence {
						if _432_v2 {
							return (func() _dafny.Sequence {
								if (_449_v17).Contains(p0) {
									return (_449_v17).Get(p0).(_dafny.Sequence)
								}
								return _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(120))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg45 _dafny.Int) interface{} {
										return coer45(arg45)
									}
								}(func(_454_i3 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('s')
								})), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(120))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg46 _dafny.Int) interface{} {
										return coer46(arg46)
									}
								}(func(_455_i3 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('s')
								}))).Cardinality()))).Uint32(), _450_v18)
							})()
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("fhkvhgy")
					})()).Cardinality()))).Uint32(), _450_v18), _428_v0)
					_ = _rhs47
					var _rhs48 bool = _432_v2
					_ = _rhs48
					var _rhs49 _dafny.Map = _451_v19
					_ = _rhs49
					var _rhs50 _dafny.Sequence = _437_v7
					_ = _rhs50
					var _lhs28 _dafny.Array = _445_v15
					_ = _lhs28
					var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_445_v15), 0))
					_ = _lhs29
					var _lhs30 _dafny.Array = _448_v16
					_ = _lhs30
					var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_448_v16), 0))
					_ = _lhs31
					(_lhs28).ArraySet1(_rhs46, (_lhs29).Int())
					_428_v0 = _rhs47
					_432_v2 = _rhs48
					(_lhs30).ArraySet1(_rhs49, (_lhs31).Int())
					_437_v7 = _rhs50
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _456_v20 bool
		_ = _456_v20
		_456_v20 = true
		var _457_v21 _dafny.Sequence
		_ = _457_v21
		_457_v21 = _dafny.SeqOf(p0)
		var _source5 D14 = Companion_Default___.Fm45(_456_v20, _456_v20, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_457_v21).Cardinality()), p0), globalState)
		_ = _source5
		if _source5.Is_DC35() {
			var _458___mcc_h0 _dafny.Int = _source5.Get_().(D14_DC35).Cf63
			_ = _458___mcc_h0
			var _459___mcc_h1 _dafny.Sequence = _source5.Get_().(D14_DC35).Cf64
			_ = _459___mcc_h1
			var _460___mcc_h2 _dafny.Int = _source5.Get_().(D14_DC35).Cf65
			_ = _460___mcc_h2
			var _461___mcc_h3 bool = _source5.Get_().(D14_DC35).Cf66
			_ = _461___mcc_h3
			var _462___mcc_h4 bool = _source5.Get_().(D14_DC35).Cf67
			_ = _462___mcc_h4
			var _463_cf67 bool = _462___mcc_h4
			_ = _463_cf67
			var _464_cf66 bool = _461___mcc_h3
			_ = _464_cf66
			var _465_cf65 _dafny.Int = _460___mcc_h2
			_ = _465_cf65
			var _466_cf64 _dafny.Sequence = _459___mcc_h1
			_ = _466_cf64
			var _467_cf63 _dafny.Int = _458___mcc_h0
			_ = _467_cf63
			var _468_v22 _dafny.Map
			_ = _468_v22
			_468_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_467_cf63, _465_cf65)
			var _469_v23 D12
			_ = _469_v23
			_469_v23 = Companion_D12_.Create_DC29_(_456_v20, (func() _dafny.Int {
				if (_468_v22).Contains(p0) {
					return (_468_v22).Get(p0).(_dafny.Int)
				}
				return _465_cf65
			})(), _456_v20, (_dafny.Zero).Minus((_465_cf65).Plus(_467_cf63)), _464_cf66)
			var _source6 D12 = _469_v23
			_ = _source6
			if _source6.Is_DC28() {
				var _470___mcc_h8 bool = _source6.Get_().(D12_DC28).Cf46
				_ = _470___mcc_h8
				var _471___mcc_h9 _dafny.Int = _source6.Get_().(D12_DC28).Cf47
				_ = _471___mcc_h9
				var _472_cf47 _dafny.Int = _471___mcc_h9
				_ = _472_cf47
				var _473_cf46 bool = _470___mcc_h8
				_ = _473_cf46
				_472_cf47 = (p0).Plus(_472_cf47)
				var _474_v24 _dafny.Sequence
				_ = _474_v24
				_474_v24 = _dafny.SeqOf(((_457_v21).Select((Companion_Default___.SafeIndex(_467_cf63, _dafny.IntOfUint32((_457_v21).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_472_cf47) <= 0)
				_474_v24 = _474_v24
				_467_cf63 = _465_cf65
				_465_cf65 = _472_cf47
			} else if _source6.Is_DC29() {
				var _475___mcc_h10 bool = _source6.Get_().(D12_DC29).Cf48
				_ = _475___mcc_h10
				var _476___mcc_h11 _dafny.Int = _source6.Get_().(D12_DC29).Cf49
				_ = _476___mcc_h11
				var _477___mcc_h12 bool = _source6.Get_().(D12_DC29).Cf50
				_ = _477___mcc_h12
				var _478___mcc_h13 _dafny.Int = _source6.Get_().(D12_DC29).Cf51
				_ = _478___mcc_h13
				var _479___mcc_h14 bool = _source6.Get_().(D12_DC29).Cf52
				_ = _479___mcc_h14
				var _480_cf52 bool = _479___mcc_h14
				_ = _480_cf52
				var _481_cf51 _dafny.Int = _478___mcc_h13
				_ = _481_cf51
				var _482_cf50 bool = _477___mcc_h12
				_ = _482_cf50
				var _483_cf49 _dafny.Int = _476___mcc_h11
				_ = _483_cf49
				var _484_cf48 bool = _475___mcc_h10
				_ = _484_cf48
				var _485_v25 _dafny.Array
				_ = _485_v25
				var _nw74 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
				_ = _nw74
				_485_v25 = _nw74
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_485_v25), 0))
				_ = _index59
				(_485_v25).ArraySet1((_465_cf65).Minus(p0), (_index59).Int())
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_485_v25), 0))
				_ = _index60
				(_485_v25).ArraySet1(_483_cf49, (_index60).Int())
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_485_v25), 0))
				_ = _index61
				(_485_v25).ArraySet1((_dafny.IntOfUint32((_428_v0).Cardinality())).Minus(_467_cf63), (_index61).Int())
				var _486_v26 D6
				_ = _486_v26
				_486_v26 = Companion_D6_.Create_DC14_(_428_v0)
				var _487_v27 _dafny.CodePoint
				_ = _487_v27
				_487_v27 = _dafny.CodePoint('d')
				var _488_v28 _dafny.Sequence
				_ = _488_v28
				_488_v28 = _dafny.SeqOf(_466_cf64)
				var _489_v29 _dafny.Array
				_ = _489_v29
				var _nwElement0_12 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_486_v26).Dtor_cf29(), (Companion_Default___.SafeIndex(Companion_Default___.Fm44(_484_cf48, globalState), _dafny.IntOfUint32(((_486_v26).Dtor_cf29()).Cardinality()))).Uint32(), _487_v27), _428_v0)
				_ = _nwElement0_12
				var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(28))
				_ = _nw75
				(_nw75).ArraySet1(_nwElement0_12, 0)
				(_nw75).ArraySet1(_428_v0, 1)
				(_nw75).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_466_cf64, Companion_Default___.Fm46(true, _dafny.IntOfUint32((_dafny.SeqOf(_482_cf50)).Cardinality()), _484_cf48, globalState)), 2)
				(_nw75).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("y"), 3)
				(_nw75).ArraySet1(_428_v0, 4)
				(_nw75).ArraySet1(_466_cf64, 5)
				(_nw75).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rc"), 6)
				(_nw75).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_466_cf64, _428_v0), 7)
				(_nw75).ArraySet1(_428_v0, 8)
				(_nw75).ArraySet1(_466_cf64, 9)
				(_nw75).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ovqriexd"), 10)
				(_nw75).ArraySet1(Companion_Default___.Fm46(true, Companion_Default___.Fm44(_484_cf48, globalState), _482_cf50, globalState), 11)
				(_nw75).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(134))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_490_v27 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_491_i4 _dafny.Int) _dafny.CodePoint {
						return _490_v27
					}
				})(_487_v27))), 12)
				(_nw75).ArraySet1(_428_v0, 13)
				(_nw75).ArraySet1((_486_v26).Dtor_cf29(), 14)
				(_nw75).ArraySet1(_428_v0, 15)
				(_nw75).ArraySet1(_466_cf64, 16)
				(_nw75).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(552))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}(func(_492_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				})), 17)
				(_nw75).ArraySet1(_428_v0, 18)
				(_nw75).ArraySet1(_428_v0, 19)
				(_nw75).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-733))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}((func(_493_v27 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_494_i6 _dafny.Int) _dafny.CodePoint {
						return _493_v27
					}
				})(_487_v27))), 20)
				(_nw75).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_466_cf64, (_488_v28).Select((Companion_Default___.SafeIndex(_465_cf65, _dafny.IntOfUint32((_488_v28).Cardinality()))).Uint32()).(_dafny.Sequence)), 21)
				(_nw75).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("crcuft"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(599))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}((func(_495_v27 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_496_i7 _dafny.Int) _dafny.CodePoint {
						return _495_v27
					}
				})(_487_v27)))), 22)
				(_nw75).ArraySet1((Companion_D6_.Create_DC14_(_466_cf64)).Dtor_cf29(), 23)
				(_nw75).ArraySet1(_428_v0, 24)
				(_nw75).ArraySet1(_428_v0, 25)
				(_nw75).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("nys"), 26)
				(_nw75).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jijhsve"), 27)
				_489_v29 = _nw75
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_489_v29), 0))
				_ = _index62
				(_489_v29).ArraySet1(_466_cf64, (_index62).Int())
				var _497_v30 _dafny.MultiSet
				_ = _497_v30
				_497_v30 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("sfisdg"), (Companion_Default___.SafeIndex(_467_cf63, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sfisdg")).Cardinality()))).Uint32(), _487_v27), _466_cf64)
				var _498_v31 _dafny.Map
				_ = _498_v31
				_498_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_464_cf66, _497_v30)
				var _499_v32 _dafny.Map
				_ = _499_v32
				_499_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _484_cf48)
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_489_v29), 0))
				_ = _index63
				var _rhs51 _dafny.Array = _485_v25
				_ = _rhs51
				var _rhs52 _dafny.Sequence = _428_v0
				_ = _rhs52
				var _rhs53 _dafny.Int = (_dafny.Zero).Minus(_465_cf65)
				_ = _rhs53
				var _rhs54 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-480), ((func() _dafny.MultiSet {
					if (_498_v31).Contains((func() bool {
						if (_499_v32).Contains((_485_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_485_v25), 0))).Int()).(_dafny.Int)) {
							return (_499_v32).Get((_485_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_485_v25), 0))).Int()).(_dafny.Int)).(bool)
						}
						return _484_cf48
					})()) {
						return (_498_v31).Get((func() bool {
							if (_499_v32).Contains((_485_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_485_v25), 0))).Int()).(_dafny.Int)) {
								return (_499_v32).Get((_485_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_485_v25), 0))).Int()).(_dafny.Int)).(bool)
							}
							return _484_cf48
						})()).(_dafny.MultiSet)
					}
					return _497_v30
				})()).Cardinality())
				_ = _rhs54
				var _lhs32 _dafny.Array = _489_v29
				_ = _lhs32
				var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_489_v29), 0))
				_ = _lhs33
				_485_v25 = _rhs51
				(_lhs32).ArraySet1(_rhs52, (_lhs33).Int())
				_467_cf63 = _rhs53
				_467_cf63 = _rhs54
				_463_cf67 = _484_cf48
			} else {
				var _500___mcc_h15 _dafny.Set = _source6.Get_().(D12_DC27).Cf45
				_ = _500___mcc_h15
				var _501_cf45 _dafny.Set = _500___mcc_h15
				_ = _501_cf45
				var _502_v33 _dafny.Array
				_ = _502_v33
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_10
				var _nw76 _dafny.Array
				_ = _nw76
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw76 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Int = func(_503_i8 _dafny.Int) _dafny.Int {
						return (_503_i8).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))
					}
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw76 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw76).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw76).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_502_v33 = _nw76
				var _504_v34 _dafny.Array
				_ = _504_v34
				_504_v34 = _502_v33
				_504_v34 = _504_v34
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_502_v33), 0))
				_ = _index64
				(_502_v33).ArraySet1(_467_cf63, (_index64).Int())
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_502_v33), 0))
				_ = _index65
				(_502_v33).ArraySet1((_dafny.Zero).Minus(_465_cf65), (_index65).Int())
				var _505_v35 _dafny.CodePoint
				_ = _505_v35
				_505_v35 = _dafny.CodePoint('h')
				var _506_v36 D0
				_ = _506_v36
				_506_v36 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_466_cf64).Cardinality()))
				var _507_v37 _dafny.Map
				_ = _507_v37
				_507_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_505_v35, _506_v36)
				var _508_v38 _dafny.Map
				_ = _508_v38
				_508_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_464_cf66), _463_cf67)
				var _509_v39 _dafny.Array
				_ = _509_v39
				var _nwElement0_13 bool = _456_v20
				_ = _nwElement0_13
				var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(28))
				_ = _nw77
				(_nw77).ArraySet1(_nwElement0_13, 0)
				(_nw77).ArraySet1(_464_cf66, 1)
				(_nw77).ArraySet1(false, 2)
				(_nw77).ArraySet1(_463_cf67, 3)
				(_nw77).ArraySet1(_464_cf66, 4)
				(_nw77).ArraySet1(_463_cf67, 5)
				(_nw77).ArraySet1(_463_cf67, 6)
				(_nw77).ArraySet1(_463_cf67, 7)
				(_nw77).ArraySet1(_464_cf66, 8)
				(_nw77).ArraySet1(_464_cf66, 9)
				(_nw77).ArraySet1(_463_cf67, 10)
				(_nw77).ArraySet1((_465_cf65).Cmp((_dafny.Zero).Minus(_467_cf63)) < 0, 11)
				(_nw77).ArraySet1(!(!(!(_456_v20))), 12)
				(_nw77).ArraySet1((_this).Fm42(_464_cf66, globalState), 13)
				(_nw77).ArraySet1(_456_v20, 14)
				(_nw77).ArraySet1(_464_cf66, 15)
				(_nw77).ArraySet1((_this).Fm42(_464_cf66, globalState), 16)
				(_nw77).ArraySet1(_463_cf67, 17)
				(_nw77).ArraySet1((_465_cf65).Cmp(_465_cf65) <= 0, 18)
				(_nw77).ArraySet1(_463_cf67, 19)
				(_nw77).ArraySet1(_463_cf67, 20)
				(_nw77).ArraySet1(!(_464_cf66), 21)
				(_nw77).ArraySet1((_this).Fm5((_502_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_502_v33), 0))).Int()).(_dafny.Int), _464_cf66, _507_v37, (_this).Fm5(p0, true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_505_v35, _506_v36), _464_cf66, globalState), globalState), 22)
				(_nw77).ArraySet1(false, 23)
				(_nw77).ArraySet1((_464_cf66) == (_463_cf67), 24)
				(_nw77).ArraySet1((func() bool {
					if (_508_v38).Contains((_this).Fm42(_456_v20, globalState)) {
						return (_508_v38).Get((_this).Fm42(_456_v20, globalState)).(bool)
					}
					return _456_v20
				})(), 25)
				(_nw77).ArraySet1(_456_v20, 26)
				(_nw77).ArraySet1(_464_cf66, 27)
				_509_v39 = _nw77
				var _510_v40 D2
				_ = _510_v40
				_510_v40 = Companion_D2_.Create_DC5_(_465_cf65, _457_v21, _456_v20)
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_509_v39), 0))
				_ = _index66
				(_509_v39).ArraySet1((func() bool {
					if _464_cf66 {
						return (_510_v40).Dtor_cf10()
					}
					return _463_cf67
				})(), (_index66).Int())
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_509_v39), 0))
				_ = _index67
				(_509_v39).ArraySet1((_463_cf67) || (_464_cf66), (_index67).Int())
				_502_v33 = _502_v33
			}
			var _511_v41 _dafny.Array
			_ = _511_v41
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_11
			var _nw78 _dafny.Array
			_ = _nw78
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw78 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) bool = (func(_512_v20 bool) func(_dafny.Int) bool {
					return func(_513_i9 _dafny.Int) bool {
						return _512_v20
					}
				})(_456_v20)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw78 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw78).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw78).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_511_v41 = _nw78
			var _514_v42 _dafny.Sequence
			_ = _514_v42
			_514_v42 = _dafny.SeqOf(true, false)
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))
			_ = _index68
			(_511_v41).ArraySet1((_514_v42).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_514_v42).Cardinality()))).Uint32()).(bool), (_index68).Int())
			var _515_v43 _dafny.Array
			_ = _515_v43
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_12
			var _nw79 _dafny.Array
			_ = _nw79
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw79 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Int = (func(_516_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_517_i10 _dafny.Int) _dafny.Int {
						return (_517_i10).Times(_516_p0)
					}
				})(p0)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw79 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw79).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw79).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_515_v43 = _nw79
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_515_v43), 0))
			_ = _index69
			(_515_v43).ArraySet1(_465_cf65, (_index69).Int())
			var _518_v44 _dafny.CodePoint
			_ = _518_v44
			_518_v44 = _dafny.CodePoint('t')
			var _519_v45 D0
			_ = _519_v45
			_519_v45 = Companion_D0_.Create_DC1_(p0)
			var _520_v46 _dafny.Map
			_ = _520_v46
			_520_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_518_v44, _519_v45)
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))
			_ = _index70
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_515_v43), 0))
			_ = _index71
			var _rhs55 bool = _dafny.Companion_Sequence_.Equal(_428_v0, _428_v0)
			_ = _rhs55
			var _rhs56 _dafny.Int = Companion_Default___.Fm44((_this).Fm5(_465_cf65, _464_cf66, _520_v46, _464_cf66, globalState), globalState)
			_ = _rhs56
			var _lhs34 _dafny.Array = _511_v41
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))
			_ = _lhs35
			var _lhs36 _dafny.Array = _515_v43
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_515_v43), 0))
			_ = _lhs37
			(_lhs34).ArraySet1(_rhs55, (_lhs35).Int())
			(_lhs36).ArraySet1(_rhs56, (_lhs37).Int())
			if _456_v20 {
				var _521_v47 _dafny.Map
				_ = _521_v47
				_521_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _464_cf66)
				var _522_v48 _dafny.Set
				_ = _522_v48
				_522_v48 = _dafny.SetOf((_521_v47).Update((_515_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_515_v43), 0))).Int()).(_dafny.Int), (_511_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))).Int()).(bool)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm47(_dafny.IntOfUint32((_428_v0).Cardinality()), (_511_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))).Int()).(bool), globalState)).Cardinality(), (_511_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))).Int()).(bool)))
				var _523_v49 D13
				_ = _523_v49
				_523_v49 = Companion_D13_.Create_DC31_(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v20, _465_cf65)).Cardinality(), (_515_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_515_v43), 0))).Int()).(_dafny.Int))
				var _524_v50 _dafny.MultiSet
				_ = _524_v50
				_524_v50 = _dafny.MultiSetOf(_467_cf63, (_522_v48).Cardinality(), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_514_v42).Cardinality()), (_523_v49).Dtor_cf56()))
				var _525_v51 D7
				_ = _525_v51
				_525_v51 = Companion_D7_.Create_DC18_()
				_524_v50 = (_dafny.MultiSetOf((_dafny.Zero).Minus(Companion_Default___.Fm44(false, globalState)), _465_cf65)).Union(Companion_Default___.Fm48(_525_v51, (_515_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_515_v43), 0))).Int()).(_dafny.Int), Companion_Default___.Fm44(true, globalState), _dafny.IntOfInt64(-50), globalState))
				var _526_v52 _dafny.Map
				_ = _526_v52
				_526_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_518_v44, !((_511_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))).Int()).(bool))), _465_cf65)
				_526_v52 = ((func() _dafny.Map {
					if _456_v20 {
						return _526_v52
					}
					return _526_v52
				})()).Merge(_526_v52)
				var _527_v53 _dafny.Array
				_ = _527_v53
				var _nw80 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
				_ = _nw80
				_527_v53 = _nw80
				var _528_v54 *C1
				_ = _528_v54
				var _nw81 *C1 = New_C1_()
				_ = _nw81
				_nw81.Ctor__(_467_cf63, _527_v53, _dafny.Companion_Sequence_.Concatenate(_457_v21, _457_v21), _519_v45)
				_528_v54 = _nw81
				var _529_v55 _dafny.Map
				_ = _529_v55
				_529_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm46((_511_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))).Int()).(bool), (_524_v50).Cardinality(), false, globalState), (_511_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))).Int()).(bool))
				var _530_v56 _dafny.Map
				_ = _530_v56
				_530_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_457_v21, _466_cf64)
				var _531_v57 _dafny.Sequence
				_ = _531_v57
				_531_v57 = _dafny.SeqOf(_428_v0, (func() _dafny.Sequence {
					if (_530_v56).Contains(_457_v21) {
						return (_530_v56).Get(_457_v21).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("mmjnwlr")
				})(), _466_cf64, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("l"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality()))).Uint32(), _518_v44))
				var _532_v58 _dafny.Set
				_ = _532_v58
				_532_v58 = _dafny.SetOf((_511_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))).Int()).(bool), _456_v20, true)
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))
				_ = _index72
				(_511_v41).ArraySet1((func() bool {
					if (_529_v55).Contains((_531_v57).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.IntOfUint32((_531_v57).Cardinality()))).Uint32()).(_dafny.Sequence)) {
						return (_529_v55).Get((_531_v57).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.IntOfUint32((_531_v57).Cardinality()))).Uint32()).(_dafny.Sequence)).(bool)
					}
					return (_532_v58).IsProperSubsetOf(_dafny.SetOf((_511_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))).Int()).(bool)))
				})(), (_index72).Int())
				var _533_v59 _dafny.Map
				_ = _533_v59
				_533_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_518_v44, _528_v54.F9)
				var _534_v60 _dafny.Map
				_ = _534_v60
				_534_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_533_v59).Merge(_533_v59), _518_v44)
				_534_v60 = _534_v60
			} else {
				var _535_v61 *C0
				_ = _535_v61
				var _nw82 *C0 = New_C0_()
				_ = _nw82
				_nw82.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("acttpcc"))
				_535_v61 = _nw82
				var _536_v62 _dafny.Array
				_ = _536_v62
				var _nwElement0_14 _dafny.Sequence = _514_v42
				_ = _nwElement0_14
				var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.One)
				_ = _nw83
				(_nw83).ArraySet1(_nwElement0_14, 0)
				_536_v62 = _nw83
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_536_v62), 0))
				_ = _index73
				(_536_v62).ArraySet1(_514_v42, (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_536_v62), 0))
				_ = _index74
				(_536_v62).ArraySet1(_514_v42, (_index74).Int())
				_456_v20 = _463_cf67
				_465_cf65 = p0
				var _537_v63 _dafny.Array
				_ = _537_v63
				var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(14))
				_ = _nw84
				_537_v63 = _nw84
				var _538_v64 *C1
				_ = _538_v64
				var _nw85 *C1 = New_C1_()
				_ = _nw85
				_nw85.Ctor__((_515_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_515_v43), 0))).Int()).(_dafny.Int), _537_v63, _dafny.Companion_Sequence_.Update(_457_v21, (Companion_Default___.SafeIndex(_465_cf65, _dafny.IntOfUint32((_457_v21).Cardinality()))).Uint32(), _465_cf65), _519_v45)
				_538_v64 = _nw85
			}
			if _463_cf67 {
				var _539_v65 _dafny.Map
				_ = _539_v65
				_539_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.Companion_Sequence_.IsPrefixOf(_428_v0, _dafny.SeqOf(_518_v44)))
				_539_v65 = (_539_v65).Update((_515_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_515_v43), 0))).Int()).(_dafny.Int), true)
				_456_v20 = !((_511_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))).Int()).(bool))
				var _540_v66 _dafny.Set
				_ = _540_v66
				_540_v66 = _dafny.SetOf(_457_v21)
				var _541_v67 D15
				_ = _541_v67
				_541_v67 = Companion_D15_.Create_DC37_(_540_v66)
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))
				_ = _index75
				(_511_v41).ArraySet1((_540_v66).IsSubsetOf((_541_v67).Dtor_cf70()), (_index75).Int())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))
				_ = _index76
				(_511_v41).ArraySet1(_464_cf66, (_index76).Int())
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_515_v43), 0))
				_ = _index77
				(_515_v43).ArraySet1(Companion_Default___.SafeDivisionInt(p0, (_515_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_515_v43), 0))).Int()).(_dafny.Int)), (_index77).Int())
			} else {
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_515_v43), 0))
				_ = _index78
				(_515_v43).ArraySet1(_465_cf65, (_index78).Int())
				var _542_v68 _dafny.Set
				_ = _542_v68
				_542_v68 = _dafny.SetOf((_511_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))).Int()).(bool), (_this).Fm5(p0, true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_428_v0).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_467_cf63, (_511_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))).Int()).(bool))).Cardinality(), _dafny.IntOfUint32((_428_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), _519_v45), false, globalState))
				var _543_v69 _dafny.Map
				_ = _543_v69
				_543_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_542_v68, _456_v20)
				_465_cf65 = ((_515_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_515_v43), 0))).Int()).(_dafny.Int)).Minus((_543_v69).Cardinality())
				_463_cf67 = _463_cf67
				var _544_v70 _dafny.Array
				_ = _544_v70
				var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
				_ = _nw86
				_544_v70 = _nw86
				var _pat_let_tv9 = _467_cf63
				_ = _pat_let_tv9
				var _545_v71 *C1
				_ = _545_v71
				var _nw87 *C1 = New_C1_()
				_ = _nw87
				_nw87.Ctor__(_dafny.IntOfInt64(-834), _544_v70, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-92))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}((func(_546_v20 bool) func(_dafny.Int) _dafny.Int {
					return func(_547_i11 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus((_dafny.MultiSetOf(_546_v20)).Cardinality())
					}
				})(_456_v20))), func(_pat_let12_0 D0) D0 {
					return func(_548_dt__update__tmp_h1 D0) D0 {
						return func(_pat_let13_0 _dafny.Int) D0 {
							return func(_549_dt__update_hcf2_h0 _dafny.Int) D0 {
								return Companion_D0_.Create_DC1_(_549_dt__update_hcf2_h0)
							}(_pat_let13_0)
						}(_pat_let_tv9)
					}(_pat_let12_0)
				}(_519_v45))
				_545_v71 = _nw87
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_511_v41), 0))
				_ = _index79
				(_511_v41).ArraySet1(false, (_index79).Int())
			}
		} else if _source5.Is_DC36() {
			var _550___mcc_h5 _dafny.Map = _source5.Get_().(D14_DC36).Cf68
			_ = _550___mcc_h5
			var _551___mcc_h6 _dafny.Map = _source5.Get_().(D14_DC36).Cf69
			_ = _551___mcc_h6
			var _552_cf69 _dafny.Map = _551___mcc_h6
			_ = _552_cf69
			var _553_cf68 _dafny.Map = _550___mcc_h5
			_ = _553_cf68
			var _554_v72 _dafny.Array
			_ = _554_v72
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_13
			var _nw88 _dafny.Array
			_ = _nw88
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw88 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Int = (func(_555_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_556_i12 _dafny.Int) _dafny.Int {
						return (_556_i12).Minus(_555_p0)
					}
				})(p0)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw88 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw88).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw88).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_554_v72 = _nw88
			var _557_v73 _dafny.Map
			_ = _557_v73
			_557_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _554_v72)
			var _558_v74 D12
			_ = _558_v74
			_558_v74 = Companion_D12_.Create_DC28_((_557_v73).Equals((_557_v73).Update(_456_v20, _554_v72)), p0)
			_558_v74 = _558_v74
			var _559_v75 _dafny.Map
			_ = _559_v75
			_559_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), !(_456_v20) || (_456_v20))
			_559_v75 = (_559_v75).Update(!(_552_cf69).Contains(_456_v20), _456_v20)
			var _560_v76 _dafny.Map
			_ = _560_v76
			_560_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _456_v20)
			var _561_v77 _dafny.Map
			_ = _561_v77
			_561_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v20, _560_v76)
			_552_cf69 = (_552_cf69).Update(!(_561_v77).Contains(_456_v20), p0)
			if _456_v20 {
				var _562_v78 _dafny.Int
				_ = _562_v78
				_562_v78 = _dafny.IntOfInt64(633)
				_562_v78 = p0
				var _563_v79 _dafny.Set
				_ = _563_v79
				_563_v79 = _dafny.SetOf(true, _456_v20)
				var _564_v80 _dafny.MultiSet
				_ = _564_v80
				_564_v80 = _dafny.MultiSetOf(_563_v79)
				var _565_v81 _dafny.Map
				_ = _565_v81
				_565_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v20, _564_v80)
				_564_v80 = ((_564_v80).Update(_563_v79, Companion_Default___.Abs(_562_v78))).Union((func() _dafny.MultiSet {
					if (_565_v81).Contains(_456_v20) {
						return (_565_v81).Get(_456_v20).(_dafny.MultiSet)
					}
					return _dafny.MultiSetOf(_563_v79)
				})())
				_557_v73 = (_557_v73).Update((func() bool {
					if (_559_v75).Contains(_456_v20) {
						return (_559_v75).Get(_456_v20).(bool)
					}
					return _456_v20
				})(), _554_v72)
				var _566_v82 _dafny.Sequence
				_ = _566_v82
				_566_v82 = _dafny.SeqOf(!(_456_v20), _456_v20, _456_v20)
				var _567_v83 _dafny.Sequence
				_ = _567_v83
				_567_v83 = _dafny.SeqOf(_566_v82, _566_v82)
				var _568_v84 _dafny.Map
				_ = _568_v84
				_568_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_566_v82).Select((Companion_Default___.SafeIndex(_562_v78, _dafny.IntOfUint32((_566_v82).Cardinality()))).Uint32()).(bool), (_567_v83).Select((Companion_Default___.SafeIndex((_457_v21).Select((Companion_Default___.SafeIndex(_562_v78, _dafny.IntOfUint32((_457_v21).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_567_v83).Cardinality()))).Uint32()).(_dafny.Sequence))
				_568_v84 = _568_v84
				var _569_v85 *C0
				_ = _569_v85
				var _nw89 *C0 = New_C0_()
				_ = _nw89
				_nw89.Ctor__(_428_v0)
				_569_v85 = _nw89
			} else {
				var _570_v86 _dafny.Array
				_ = _570_v86
				var _nw90 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw90
				_570_v86 = _nw90
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_570_v86), 0))
				_ = _index80
				(_570_v86).ArraySet1(_456_v20, (_index80).Int())
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_570_v86), 0))
				_ = _index81
				var _rhs57 _dafny.Sequence = _457_v21
				_ = _rhs57
				var _rhs58 bool = !(_456_v20) || (_456_v20)
				_ = _rhs58
				var _lhs38 _dafny.Array = _570_v86
				_ = _lhs38
				var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_570_v86), 0))
				_ = _lhs39
				_457_v21 = _rhs57
				(_lhs38).ArraySet1(_rhs58, (_lhs39).Int())
				var _571_v87 D6
				_ = _571_v87
				_571_v87 = Companion_D6_.Create_DC15_(p0, Companion_Default___.Fm44((_570_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_570_v86), 0))).Int()).(bool), globalState), !(_456_v20))
				_571_v87 = _571_v87
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_570_v86), 0))
				_ = _index82
				(_570_v86).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_428_v0, _428_v0), (_index82).Int())
				var _572_v88 _dafny.Map
				_ = _572_v88
				_572_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_552_cf69, _428_v0)
				var _573_v89 _dafny.Int
				_ = _573_v89
				var _out20 _dafny.Int
				_ = _out20
				_out20 = (_this).M14(false, ((_570_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_570_v86), 0))).Int()).(bool)) && ((_570_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_570_v86), 0))).Int()).(bool)), (_572_v88).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_552_cf69, _428_v0)), (_dafny.Zero).Minus(_dafny.IntOfUint32((_428_v0).Cardinality())), globalState)
				_573_v89 = _out20
				var _574_v90 _dafny.CodePoint
				_ = _574_v90
				_574_v90 = _dafny.CodePoint('s')
				var _575_v91 _dafny.Map
				_ = _575_v91
				_575_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _574_v90)
				_575_v91 = (_575_v91).Update((_570_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_570_v86), 0))).Int()).(bool), _574_v90)
			}
		} else {
			var _576___mcc_h7 _dafny.Map = _source5.Get_().(D14_DC34).Cf62
			_ = _576___mcc_h7
			var _577_cf62 _dafny.Map = _576___mcc_h7
			_ = _577_cf62
			var _578_v92 D12
			_ = _578_v92
			_578_v92 = Companion_D12_.Create_DC28_(!(_456_v20) || (_456_v20), p0)
			_578_v92 = _578_v92
			_456_v20 = _456_v20
			var _579_v93 _dafny.Array
			_ = _579_v93
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_14
			var _nw91 _dafny.Array
			_ = _nw91
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw91 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Set = (func(_580_v20 bool) func(_dafny.Int) _dafny.Set {
					return func(_581_i13 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf(_580_v20, _580_v20)).Difference(_dafny.SetOf(true))
					}
				})(_456_v20)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw91 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw91).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw91).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_579_v93 = _nw91
			var _582_v94 _dafny.Set
			_ = _582_v94
			_582_v94 = _dafny.SetOf(_456_v20)
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_579_v93), 0))
			_ = _index83
			(_579_v93).ArraySet1(_582_v94, (_index83).Int())
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_579_v93), 0))
			_ = _index84
			(_579_v93).ArraySet1(_582_v94, (_index84).Int())
			var _583_v95 bool
			_ = _583_v95
			_583_v95 = _456_v20
			if !(_583_v95) {
				var _584_v96 _dafny.CodePoint
				_ = _584_v96
				_584_v96 = _dafny.CodePoint('c')
				var _585_v97 _dafny.Map
				_ = _585_v97
				_585_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm44(_456_v20, globalState)).Times(p0), p0)
				var _rhs59 _dafny.CodePoint = (func() _dafny.CodePoint {
					if _456_v20 {
						return _584_v96
					}
					return _584_v96
				})()
				_ = _rhs59
				var _rhs60 bool = (_456_v20) && (_456_v20)
				_ = _rhs60
				var _rhs61 _dafny.Map = _585_v97
				_ = _rhs61
				var _rhs62 bool = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus(p0))).Equals((_585_v97).Merge(_585_v97))
				_ = _rhs62
				_584_v96 = _rhs59
				_456_v20 = _rhs60
				_585_v97 = _rhs61
				_456_v20 = _rhs62
				var _586_v98 _dafny.Map
				_ = _586_v98
				_586_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Plus(p0), ((_579_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_579_v93), 0))).Int()).(_dafny.Set)).Contains(_456_v20))
				_586_v98 = (_586_v98).Update(p0, true)
				_456_v20 = _456_v20
				var _587_v99 _dafny.MultiSet
				_ = _587_v99
				_587_v99 = _dafny.MultiSetOf(true, _456_v20, false)
				var _588_v100 _dafny.Map
				_ = _588_v100
				_588_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _587_v99)
				var _589_v101 _dafny.Array
				_ = _589_v101
				var _nwElement0_15 _dafny.MultiSet = _587_v99
				_ = _nwElement0_15
				var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(19))
				_ = _nw92
				(_nw92).ArraySet1(_nwElement0_15, 0)
				(_nw92).ArraySet1(_587_v99, 1)
				(_nw92).ArraySet1((_587_v99).Union(_587_v99), 2)
				(_nw92).ArraySet1(_dafny.MultiSetOf(_456_v20), 3)
				(_nw92).ArraySet1(_587_v99, 4)
				(_nw92).ArraySet1((_587_v99).Union(_587_v99), 5)
				(_nw92).ArraySet1(_dafny.MultiSetOf(_456_v20, _456_v20), 6)
				(_nw92).ArraySet1(_587_v99, 7)
				(_nw92).ArraySet1(_587_v99, 8)
				(_nw92).ArraySet1(_587_v99, 9)
				(_nw92).ArraySet1(_587_v99, 10)
				(_nw92).ArraySet1(_587_v99, 11)
				(_nw92).ArraySet1((func() _dafny.MultiSet {
					if (_588_v100).Contains(p0) {
						return (_588_v100).Get(p0).(_dafny.MultiSet)
					}
					return _587_v99
				})(), 12)
				(_nw92).ArraySet1(_587_v99, 13)
				(_nw92).ArraySet1(_587_v99, 14)
				(_nw92).ArraySet1(_587_v99, 15)
				(_nw92).ArraySet1((_dafny.MultiSetOf(_456_v20, _456_v20)).Update(_456_v20, Companion_Default___.Abs(p0)), 16)
				(_nw92).ArraySet1(_587_v99, 17)
				(_nw92).ArraySet1(_587_v99, 18)
				_589_v101 = _nw92
				var _590_v102 _dafny.Sequence
				_ = _590_v102
				_590_v102 = _dafny.SeqOf(_dafny.MultiSetOf((_578_v92).Dtor_cf46()))
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_589_v101), 0))
				_ = _index85
				(_589_v101).ArraySet1((_590_v102).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_590_v102).Cardinality()))).Uint32()).(_dafny.MultiSet), (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_589_v101), 0))
				_ = _index86
				(_589_v101).ArraySet1(_587_v99, (_index86).Int())
				_456_v20 = (p0).Cmp(p0) >= 0
			} else {
				var _591_v103 _dafny.Array
				_ = _591_v103
				var _nw93 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
				_ = _nw93
				_591_v103 = _nw93
				var _592_v104 D16
				_ = _592_v104
				_592_v104 = Companion_D16_.Create_DC41_(_591_v103)
				_591_v103 = (_592_v104).Dtor_cf75()
				var _593_v105 _dafny.Int
				_ = _593_v105
				_593_v105 = _dafny.IntOfInt64(795)
				var _594_v106 _dafny.Map
				_ = _594_v106
				_594_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_593_v105, _456_v20)
				var _595_v107 D0
				_ = _595_v107
				_595_v107 = Companion_D0_.Create_DC1_(_593_v105)
				_593_v105 = Companion_Default___.SafeModuloInt((p0).Times(Companion_Default___.Fm44((func() bool {
					if (_594_v106).Contains(p0) {
						return (_594_v106).Get(p0).(bool)
					}
					return _456_v20
				})(), globalState)), ((_595_v107).Dtor_cf2()).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_456_v20)).Cardinality())))
				_456_v20 = !(_456_v20)
				_456_v20 = (p0).Cmp((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_428_v0).Cardinality()), p0))) != 0
				var _596_v108 _dafny.Array
				_ = _596_v108
				var _nw94 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(13))
				_ = _nw94
				_596_v108 = _nw94
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen((_596_v108), 0))
				_ = _index87
				(_596_v108).ArraySet1(_595_v107, (_index87).Int())
				var _597_v109 _dafny.CodePoint
				_ = _597_v109
				_597_v109 = _dafny.CodePoint('l')
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen((_596_v108), 0))
				_ = _index88
				(_596_v108).ArraySet1(Companion_Default___.Fm49(_597_v109, globalState), (_index88).Int())
			}
		}
		if !(_456_v20) {
			var _598_v110 _dafny.MultiSet
			_ = _598_v110
			_598_v110 = _dafny.MultiSetOf(_456_v20)
			var _599_v111 _dafny.Map
			_ = _599_v111
			_599_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_598_v110).Cardinality(), _456_v20)
			var _600_v112 _dafny.Sequence
			_ = _600_v112
			_600_v112 = _dafny.SeqOf(_599_v111)
			var _601_v113 _dafny.Array
			_ = _601_v113
			var _nwElement0_16 _dafny.Map = _599_v111
			_ = _nwElement0_16
			var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(5))
			_ = _nw95
			(_nw95).ArraySet1(_nwElement0_16, 0)
			(_nw95).ArraySet1(_599_v111, 1)
			(_nw95).ArraySet1(Companion_Default___.Fm50(_456_v20, _456_v20, p0, globalState), 2)
			(_nw95).ArraySet1((_600_v112).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_600_v112).Cardinality()))).Uint32()).(_dafny.Map), 3)
			(_nw95).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _456_v20)).Update(_dafny.IntOfInt64(-789), _456_v20), 4)
			_601_v113 = _nw95
			var _602_v114 _dafny.Map
			_ = _602_v114
			_602_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v20, _601_v113)
			var _603_v115 *C0
			_ = _603_v115
			var _nw96 *C0 = New_C0_()
			_ = _nw96
			_nw96.Ctor__(_428_v0)
			_603_v115 = _nw96
			var _604_v116 _dafny.Sequence
			_ = _604_v116
			_604_v116 = _dafny.SeqOf(_603_v115)
			_602_v114 = (_602_v114).Update((_dafny.IntOfUint32((_604_v116).Cardinality())).Cmp(p0) >= 0, _601_v113)
			var _605_v117 _dafny.Array
			_ = _605_v117
			var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw97
			_605_v117 = _nw97
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_605_v117), 0))
			_ = _index89
			(_605_v117).ArraySet1(Companion_Default___.Fm44(!(_456_v20), globalState), (_index89).Int())
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_605_v117), 0))
			_ = _index90
			(_605_v117).ArraySet1(p0, (_index90).Int())
			_457_v21 = _dafny.Companion_Sequence_.Concatenate(_457_v21, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-305))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}((func(_606_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_607_i14 _dafny.Int) _dafny.Int {
					return _606_p0
				}
			})(p0))))
			_456_v20 = _456_v20
			var _608_v118 _dafny.Set
			_ = _608_v118
			_608_v118 = _dafny.SetOf(p0, _dafny.IntOfInt64(151), (_605_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_605_v117), 0))).Int()).(_dafny.Int))
			_608_v118 = (_608_v118).Intersection(_608_v118)
		} else {
			_456_v20 = _456_v20
			var _609_v119 _dafny.Array
			_ = _609_v119
			var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw98
			_609_v119 = _nw98
			var _610_v120 _dafny.Set
			_ = _610_v120
			_610_v120 = _dafny.SetOf(_456_v20)
			var _611_v121 _dafny.Sequence
			_ = _611_v121
			_611_v121 = _dafny.SeqOf(_610_v120, _610_v120)
			var _612_v122 D12
			_ = _612_v122
			_612_v122 = Companion_D12_.Create_DC27_(_610_v120)
			var _613_v123 _dafny.Array
			_ = _613_v123
			var _nwElement0_17 _dafny.Set = _610_v120
			_ = _nwElement0_17
			var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(18))
			_ = _nw99
			(_nw99).ArraySet1(_nwElement0_17, 0)
			(_nw99).ArraySet1(_610_v120, 1)
			(_nw99).ArraySet1((_611_v121).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_611_v121).Cardinality()))).Uint32()).(_dafny.Set), 2)
			(_nw99).ArraySet1(_610_v120, 3)
			(_nw99).ArraySet1(_610_v120, 4)
			(_nw99).ArraySet1(_610_v120, 5)
			(_nw99).ArraySet1(_610_v120, 6)
			(_nw99).ArraySet1(_610_v120, 7)
			(_nw99).ArraySet1(_610_v120, 8)
			(_nw99).ArraySet1(_610_v120, 9)
			(_nw99).ArraySet1(_610_v120, 10)
			(_nw99).ArraySet1((_612_v122).Dtor_cf45(), 11)
			(_nw99).ArraySet1(_610_v120, 12)
			(_nw99).ArraySet1(_610_v120, 13)
			(_nw99).ArraySet1(_610_v120, 14)
			(_nw99).ArraySet1(_610_v120, 15)
			(_nw99).ArraySet1(_610_v120, 16)
			(_nw99).ArraySet1(_610_v120, 17)
			_613_v123 = _nw99
			var _614_v124 D0
			_ = _614_v124
			_614_v124 = Companion_D0_.Create_DC0_(_613_v123, _dafny.IntOfInt64(596))
			var _615_v125 D15
			_ = _615_v125
			_615_v125 = Companion_D15_.Create_DC38_(Companion_Default___.Fm51(_456_v20, p0, p0, globalState))
			var _616_v126 _dafny.Map
			_ = _616_v126
			_616_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_614_v124, (_615_v125).Dtor_cf71())
			var _617_v127 _dafny.Sequence
			_ = _617_v127
			_617_v127 = _dafny.SeqOf(_456_v20)
			var _618_v128 _dafny.Sequence
			_ = _618_v128
			_618_v128 = _dafny.SeqOf(_457_v21)
			var _619_v129 _dafny.Array
			_ = _619_v129
			var _nwElement0_18 _dafny.Int = p0
			_ = _nwElement0_18
			var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(23))
			_ = _nw100
			(_nw100).ArraySet1(_nwElement0_18, 0)
			(_nw100).ArraySet1(p0, 1)
			(_nw100).ArraySet1(_dafny.IntOfInt64(550), 2)
			(_nw100).ArraySet1(_dafny.IntOfUint32((_428_v0).Cardinality()), 3)
			(_nw100).ArraySet1(Companion_Default___.Fm44(_456_v20, globalState), 4)
			(_nw100).ArraySet1(p0, 5)
			(_nw100).ArraySet1(p0, 6)
			(_nw100).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), 7)
			(_nw100).ArraySet1(p0, 8)
			(_nw100).ArraySet1(p0, 9)
			(_nw100).ArraySet1(p0, 10)
			(_nw100).ArraySet1(_dafny.IntOfUint32((_428_v0).Cardinality()), 11)
			(_nw100).ArraySet1(p0, 12)
			(_nw100).ArraySet1((_616_v126).Cardinality(), 13)
			(_nw100).ArraySet1(_dafny.IntOfUint32((_617_v127).Cardinality()), 14)
			(_nw100).ArraySet1(_dafny.IntOfUint32((_618_v128).Cardinality()), 15)
			(_nw100).ArraySet1(p0, 16)
			(_nw100).ArraySet1(p0, 17)
			(_nw100).ArraySet1((_dafny.Zero).Minus(p0), 18)
			(_nw100).ArraySet1(_dafny.IntOfInt64(259), 19)
			(_nw100).ArraySet1(p0, 20)
			(_nw100).ArraySet1(p0, 21)
			(_nw100).ArraySet1(_dafny.IntOfInt64(-31), 22)
			_619_v129 = _nw100
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_609_v119), 0))
			_ = _index91
			(_609_v119).ArraySet1(_619_v129, (_index91).Int())
			var _620_v130 _dafny.Array
			_ = _620_v130
			var _nw101 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw101
			_620_v130 = _nw101
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_620_v130), 0))
			_ = _index92
			(_620_v130).ArraySet1(_456_v20, (_index92).Int())
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_619_v129), 0))
			_ = _index93
			(_619_v129).ArraySet1(Companion_Default___.Fm44(_456_v20, globalState), (_index93).Int())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_619_v129), 0))
			_ = _index94
			(_619_v129).ArraySet1(p0, (_index94).Int())
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_609_v119), 0))
			_ = _index95
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_620_v130), 0))
			_ = _index96
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_619_v129), 0))
			_ = _index97
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_619_v129), 0))
			_ = _index98
			var _rhs63 _dafny.Array = _619_v129
			_ = _rhs63
			var _rhs64 bool = _456_v20
			_ = _rhs64
			var _rhs65 _dafny.Int = p0
			_ = _rhs65
			var _rhs66 _dafny.Int = Companion_Default___.SafeModuloInt(p0, (_457_v21).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_457_v21).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs66
			var _lhs40 _dafny.Array = _609_v119
			_ = _lhs40
			var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_609_v119), 0))
			_ = _lhs41
			var _lhs42 _dafny.Array = _620_v130
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_620_v130), 0))
			_ = _lhs43
			var _lhs44 _dafny.Array = _619_v129
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_619_v129), 0))
			_ = _lhs45
			var _lhs46 _dafny.Array = _619_v129
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_619_v129), 0))
			_ = _lhs47
			(_lhs40).ArraySet1(_rhs63, (_lhs41).Int())
			(_lhs42).ArraySet1(_rhs64, (_lhs43).Int())
			(_lhs44).ArraySet1(_rhs65, (_lhs45).Int())
			(_lhs46).ArraySet1(_rhs66, (_lhs47).Int())
			_456_v20 = (_this).Fm42(_456_v20, globalState)
			var _621_v131 _dafny.Set
			_ = _621_v131
			_621_v131 = _dafny.SetOf(_620_v130)
			var _622_v132 _dafny.Set
			_ = _622_v132
			_622_v132 = _621_v131
			var _623_v133 _dafny.Set
			_ = _623_v133
			_623_v133 = (_622_v132)
			var _source7 _dafny.Set = _623_v133
			_ = _source7
			var _624___mcc_h16 _dafny.Set = _source7
			_ = _624___mcc_h16
			var _625_cf36 _dafny.Set = _624___mcc_h16
			_ = _625_cf36
			var _626_v134 _dafny.Map
			_ = _626_v134
			_626_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(518))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}(func(_627_i15 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})), (_dafny.Zero).Minus(p0))
			var _628_v135 _dafny.Map
			_ = _628_v135
			_628_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_626_v134).Cardinality(), (_620_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_620_v130), 0))).Int()).(bool))
			var _629_v137 _dafny.CodePoint
			_ = _629_v137
			_629_v137 = _dafny.CodePoint('f')
			_456_v20 = (_this).Fm5((_628_v135).Cardinality(), !(_456_v20) || (_456_v20), func() _dafny.Map {
				var _coll38 = _dafny.NewMapBuilder()
				_ = _coll38
				for _iter39 := _dafny.Iterate((_dafny.MultiSetOf(_629_v137, _629_v137, _629_v137, _629_v137, _629_v137)).Elements()); ; {
					_compr_38, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _630_v136 _dafny.CodePoint
					_630_v136 = interface{}(_compr_38).(_dafny.CodePoint)
					if (_dafny.MultiSetOf(_629_v137, _629_v137, _629_v137, _629_v137, _629_v137)).Contains(_630_v136) {
						_coll38.Add(_630_v136, Companion_D0_.Create_DC1_(p0))
					}
				}
				return _coll38.ToMap()
			}(), _456_v20, globalState)
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_613_v123), 0))
			_ = _index99
			(_613_v123).ArraySet1(_610_v120, (_index99).Int())
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_613_v123), 0))
			_ = _index100
			(_613_v123).ArraySet1(_610_v120, (_index100).Int())
			var _631_v139 _dafny.Array
			_ = _631_v139
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_15
			var _nw102 _dafny.Array
			_ = _nw102
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw102 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.MultiSet = (func(_632_v130 _dafny.Array, _633_v129 _dafny.Array, _634_p0 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
					return func(_635_i16 _dafny.Int) _dafny.MultiSet {
						return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_632_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_632_v130), 0))).Int()).(bool), (_633_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_633_v129), 0))).Int()).(_dafny.Int)))).Union(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_632_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_632_v130), 0))).Int()).(bool), (func() _dafny.Set {
							var _coll39 = _dafny.NewBuilder()
							_ = _coll39
							for _iter40 := _dafny.Iterate((_dafny.SetOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_632_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_632_v130), 0))).Int()).(bool), (_632_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_632_v130), 0))).Int()).(bool))).Cardinality(), _634_p0))).Elements()); ; {
								_compr_39, _ok40 := _iter40()
								if !_ok40 {
									break
								}
								var _636_v138 _dafny.MultiSet
								_636_v138 = interface{}(_compr_39).(_dafny.MultiSet)
								if (_dafny.SetOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_632_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_632_v130), 0))).Int()).(bool), (_632_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_632_v130), 0))).Int()).(bool))).Cardinality(), _634_p0))).Contains(_636_v138) {
									_coll39.Add(_636_v138)
								}
							}
							return _coll39.ToSet()
						}()).Cardinality())))
					}
				})(_620_v130, _619_v129, p0)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw102 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw102).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw102).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_631_v139 = _nw102
			var _637_v140 _dafny.Map
			_ = _637_v140
			_637_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_619_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_619_v129), 0))).Int()).(_dafny.Int))
			var _638_v141 _dafny.MultiSet
			_ = _638_v141
			_638_v141 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_620_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_620_v130), 0))).Int()).(bool)), (_637_v140).Cardinality()))
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_631_v139), 0))
			_ = _index101
			(_631_v139).ArraySet1(_638_v141, (_index101).Int())
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(64), _dafny.ArrayLen((_631_v139), 0))
			_ = _index102
			(_631_v139).ArraySet1(_638_v141, (_index102).Int())
			var _639_v142 D13
			_ = _639_v142
			_639_v142 = Companion_D13_.Create_DC32_(_428_v0, (_620_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_620_v130), 0))).Int()).(bool), (_619_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_619_v129), 0))).Int()).(_dafny.Int), _dafny.ArrayCastTo((_609_v119).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_609_v119), 0))).Int())))
			var _640_v143 _dafny.Map
			_ = _640_v143
			_640_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_639_v142).Dtor_cf58()), p0)
			var _641_v144 _dafny.MultiSet
			_ = _641_v144
			_641_v144 = _dafny.MultiSetOf(_623_v133)
			var _642_v145 _dafny.Map
			_ = _642_v145
			_642_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_640_v143, (func() _dafny.Int {
				if (_641_v144).Contains(_dafny.SetOf(_620_v130, _620_v130)) {
					return (_641_v144).Multiplicity(_dafny.SetOf(_620_v130, _620_v130))
				}
				return _dafny.IntOfUint32((_428_v0).Cardinality())
			})())
			var _643_v146 D0
			_ = _643_v146
			_643_v146 = Companion_D0_.Create_DC1_((_619_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_619_v129), 0))).Int()).(_dafny.Int))
			var _644_v147 _dafny.Map
			_ = _644_v147
			_644_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_629_v137, _643_v146)
			_642_v145 = (func() _dafny.Map {
				if (_456_v20) && ((_this).Fm5(_dafny.IntOfInt64(630), _456_v20, _644_v147, (_620_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_620_v130), 0))).Int()).(bool), globalState)) {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v20, p0), (func() _dafny.Int {
						if (_640_v143).Contains(true) {
							return (_640_v143).Get(true).(_dafny.Int)
						}
						return p0
					})())
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_620_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_620_v130), 0))).Int()).(bool), _dafny.IntOfInt64(-451)), p0)
			})()
			var _645_v148 _dafny.Array
			_ = _645_v148
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_16
			var _nw103 _dafny.Array
			_ = _nw103
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw103 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) _dafny.CodePoint = func(_646_i17 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw103 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw103).ArraySet1CodePoint(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw103).ArraySet1CodePoint(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_645_v148 = _nw103
			var _647_v149 _dafny.Sequence
			_ = _647_v149
			_647_v149 = _dafny.SeqOf(_645_v148)
			_647_v149 = _647_v149
		}
		var _648_v150 _dafny.Map
		_ = _648_v150
		_648_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_428_v0, _456_v20)
		_648_v150 = (_648_v150).Update(_dafny.Companion_Sequence_.Concatenate(_428_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(680))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg54 _dafny.Int) interface{} {
				return coer54(arg54)
			}
		}(func(_649_i18 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		}))), _456_v20)
		if !(false) {
			var _650_v151 _dafny.Array
			_ = _650_v151
			var _nw104 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
			_ = _nw104
			_650_v151 = _nw104
			var _651_v152 _dafny.Array
			_ = _651_v152
			_651_v152 = _650_v151
			_651_v152 = _651_v152
			var _652_v153 _dafny.Int
			_ = _652_v153
			_652_v153 = _dafny.IntOfInt64(404)
			_652_v153 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-90))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg55 _dafny.Int) interface{} {
					return coer55(arg55)
				}
			}(func(_653_i19 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			})), _428_v0)).Cardinality())
			var _654_v154 _dafny.CodePoint
			_ = _654_v154
			_654_v154 = _dafny.CodePoint('g')
			var _655_v155 _dafny.Map
			_ = _655_v155
			_655_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_654_v154, p0)
			_655_v155 = (_655_v155).Update(_654_v154, _652_v153)
			_654_v154 = _654_v154
			_652_v153 = (_dafny.Zero).Minus(p0)
		} else {
			var _656_v156 D4
			_ = _656_v156
			_656_v156 = Companion_D4_.Create_DC10_(_dafny.IntOfUint32((_428_v0).Cardinality()), (_dafny.Zero).Minus(p0), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v20, false)).Cardinality())
			var _657_v157 _dafny.CodePoint
			_ = _657_v157
			_657_v157 = _dafny.CodePoint('a')
			var _658_v158 _dafny.Map
			_ = _658_v158
			_658_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v20, _657_v157)
			var _659_v159 _dafny.Set
			_ = _659_v159
			_659_v159 = _dafny.SetOf((_658_v158).Cardinality())
			var _660_v160 _dafny.Array
			_ = _660_v160
			var _nwElement0_19 _dafny.Int = Companion_Default___.Fm44(_456_v20, globalState)
			_ = _nwElement0_19
			var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(24))
			_ = _nw105
			(_nw105).ArraySet1(_nwElement0_19, 0)
			(_nw105).ArraySet1(p0, 1)
			(_nw105).ArraySet1(_dafny.IntOfInt64(714), 2)
			(_nw105).ArraySet1(p0, 3)
			(_nw105).ArraySet1(p0, 4)
			(_nw105).ArraySet1(p0, 5)
			(_nw105).ArraySet1((_656_v156).Dtor_cf20(), 6)
			(_nw105).ArraySet1(p0, 7)
			(_nw105).ArraySet1(p0, 8)
			(_nw105).ArraySet1(p0, 9)
			(_nw105).ArraySet1(p0, 10)
			(_nw105).ArraySet1(p0, 11)
			(_nw105).ArraySet1(p0, 12)
			(_nw105).ArraySet1((_659_v159).Cardinality(), 13)
			(_nw105).ArraySet1(p0, 14)
			(_nw105).ArraySet1(p0, 15)
			(_nw105).ArraySet1(p0, 16)
			(_nw105).ArraySet1(p0, 17)
			(_nw105).ArraySet1(_dafny.IntOfUint32((_428_v0).Cardinality()), 18)
			(_nw105).ArraySet1(p0, 19)
			(_nw105).ArraySet1(p0, 20)
			(_nw105).ArraySet1(p0, 21)
			(_nw105).ArraySet1(p0, 22)
			(_nw105).ArraySet1(p0, 23)
			_660_v160 = _nw105
			var _661_v161 _dafny.Array
			_ = _661_v161
			var _nwElement0_20 _dafny.Array = _660_v160
			_ = _nwElement0_20
			var _nw106 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(12))
			_ = _nw106
			(_nw106).ArraySet1(_nwElement0_20, 0)
			(_nw106).ArraySet1(_660_v160, 1)
			(_nw106).ArraySet1(_660_v160, 2)
			(_nw106).ArraySet1(_660_v160, 3)
			(_nw106).ArraySet1(_660_v160, 4)
			(_nw106).ArraySet1(_660_v160, 5)
			(_nw106).ArraySet1(_660_v160, 6)
			(_nw106).ArraySet1(_660_v160, 7)
			(_nw106).ArraySet1(_660_v160, 8)
			(_nw106).ArraySet1(_660_v160, 9)
			(_nw106).ArraySet1(_660_v160, 10)
			(_nw106).ArraySet1(_660_v160, 11)
			_661_v161 = _nw106
			var _662_v162 D2
			_ = _662_v162
			_662_v162 = Companion_D2_.Create_DC4_(p0, _456_v20, _661_v161)
			var _source8 D2 = _662_v162
			_ = _source8
			if _source8.Is_DC4() {
				var _663___mcc_h17 _dafny.Int = _source8.Get_().(D2_DC4).Cf5
				_ = _663___mcc_h17
				var _664___mcc_h18 bool = _source8.Get_().(D2_DC4).Cf6
				_ = _664___mcc_h18
				var _665___mcc_h19 _dafny.Array = _source8.Get_().(D2_DC4).Cf7
				_ = _665___mcc_h19
				var _666_cf7 _dafny.Array = _665___mcc_h19
				_ = _666_cf7
				var _667_cf6 bool = _664___mcc_h18
				_ = _667_cf6
				var _668_cf5 _dafny.Int = _663___mcc_h17
				_ = _668_cf5
				_456_v20 = false
				var _669_v163 _dafny.Array
				_ = _669_v163
				var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(16))
				_ = _nw107
				_669_v163 = _nw107
				var _670_v164 _dafny.MultiSet
				_ = _670_v164
				_670_v164 = _dafny.MultiSetOf(_456_v20)
				var _671_v165 D9
				_ = _671_v165
				_671_v165 = Companion_D9_.Create_DC21_(_670_v164)
				var _672_v166 D15
				_ = _672_v166
				_672_v166 = Companion_D15_.Create_DC39_(_428_v0, _671_v165)
				var _673_v167 _dafny.Sequence
				_ = _673_v167
				_673_v167 = _dafny.SeqOf(_428_v0, (_672_v166).Dtor_cf72())
				var _674_v168 *C1
				_ = _674_v168
				var _nw108 *C1 = New_C1_()
				_ = _nw108
				_nw108.Ctor__(p0, _669_v163, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(378))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}((func(_675_v21 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_676_i20 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_675_v21).Cardinality())
					}
				})(_457_v21))), _457_v21), Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_673_v167).Cardinality())))
				_674_v168 = _nw108
				_660_v160 = _660_v160
				_428_v0 = _428_v0
			} else if _source8.Is_DC5() {
				var _677___mcc_h20 _dafny.Int = _source8.Get_().(D2_DC5).Cf8
				_ = _677___mcc_h20
				var _678___mcc_h21 _dafny.Sequence = _source8.Get_().(D2_DC5).Cf9
				_ = _678___mcc_h21
				var _679___mcc_h22 bool = _source8.Get_().(D2_DC5).Cf10
				_ = _679___mcc_h22
				var _680_cf10 bool = _679___mcc_h22
				_ = _680_cf10
				var _681_cf9 _dafny.Sequence = _678___mcc_h21
				_ = _681_cf9
				var _682_cf8 _dafny.Int = _677___mcc_h20
				_ = _682_cf8
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_660_v160), 0))
				_ = _index103
				(_660_v160).ArraySet1(_dafny.IntOfInt64(711), (_index103).Int())
				var _683_v169 _dafny.MultiSet
				_ = _683_v169
				_683_v169 = _dafny.MultiSetOf(true, _456_v20)
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_660_v160), 0))
				_ = _index104
				(_660_v160).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_682_cf8, _dafny.IntOfUint32((_428_v0).Cardinality())), (func() _dafny.Int {
					if (_683_v169).Contains(_456_v20) {
						return (_683_v169).Multiplicity(_456_v20)
					}
					return p0
				})()), (_index104).Int())
				var _684_v170 _dafny.Sequence
				_ = _684_v170
				_684_v170 = _dafny.SeqOf(_428_v0, _428_v0)
				var _685_v171 _dafny.Map
				_ = _685_v171
				_685_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_684_v170, p0)
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_660_v160), 0))
				_ = _index105
				(_660_v160).ArraySet1((func() _dafny.Int {
					if (_685_v171).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-291))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg57 _dafny.Int) interface{} {
							return coer57(arg57)
						}
					}((func(_686_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_687_i21 _dafny.Int) _dafny.Sequence {
							return _686_v0
						}
					})(_428_v0)))) {
						return (_685_v171).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-291))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg58 _dafny.Int) interface{} {
								return coer58(arg58)
							}
						}((func(_688_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_689_i21 _dafny.Int) _dafny.Sequence {
								return _688_v0
							}
						})(_428_v0)))).(_dafny.Int)
					}
					return ((_dafny.Zero).Minus(_682_cf8)).Minus(_682_cf8)
				})(), (_index105).Int())
				var _690_v172 _dafny.Array
				_ = _690_v172
				var _nw109 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
				_ = _nw109
				_690_v172 = _nw109
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_690_v172), 0))
				_ = _index106
				(_690_v172).ArraySet1((_680_cf10) == (true), (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_690_v172), 0))
				_ = _index107
				(_690_v172).ArraySet1(_456_v20, (_index107).Int())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_660_v160), 0))
				_ = _index108
				(_660_v160).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_660_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_660_v160), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-912)), (_dafny.IntOfUint32((_428_v0).Cardinality())).Times(p0)), (_index108).Int())
			} else if _source8.Is_DC3() {
				var _691___mcc_h23 _dafny.Set = _source8.Get_().(D2_DC3).Cf4
				_ = _691___mcc_h23
				var _692_cf4 _dafny.Set = _691___mcc_h23
				_ = _692_cf4
				var _693_v173 _dafny.Int
				_ = _693_v173
				_693_v173 = _dafny.IntOfInt64(602)
				var _694_v174 D15
				_ = _694_v174
				_694_v174 = Companion_D15_.Create_DC38_((func() _dafny.CodePoint {
					if _456_v20 {
						return Companion_Default___.Fm51(_456_v20, _693_v173, _693_v173, globalState)
					}
					return _657_v157
				})())
				var _695_v175 _dafny.Map
				_ = _695_v175
				_695_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v20, _457_v21)
				var _rhs67 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
					if _456_v20 {
						return (func() _dafny.Sequence {
							if (_695_v175).Contains(_456_v20) {
								return (_695_v175).Get(_456_v20).(_dafny.Sequence)
							}
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(10))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg59 _dafny.Int) interface{} {
									return coer59(arg59)
								}
							}((func(_696_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_697_i22 _dafny.Int) _dafny.Int {
									return _696_p0
								}
							})(p0)))
						})()
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-820))).Uint32(), func(coer60 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg60 _dafny.Int) interface{} {
							return coer60(arg60)
						}
					}((func(_698_v173 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_699_i23 _dafny.Int) _dafny.Int {
							return _698_v173
						}
					})(_693_v173)))
				})()).Cardinality())
				_ = _rhs67
				var _rhs68 _dafny.Set = (_692_cf4).Intersection(_692_cf4)
				_ = _rhs68
				var _rhs69 D15 = _694_v174
				_ = _rhs69
				var _rhs70 bool = _456_v20
				_ = _rhs70
				_693_v173 = _rhs67
				_692_cf4 = _rhs68
				_694_v174 = _rhs69
				_456_v20 = _rhs70
				_693_v173 = ((p0).Plus((_dafny.Zero).Minus(p0))).Plus(_693_v173)
				var _700_v176 _dafny.Array
				_ = _700_v176
				var _nw110 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw110
				_700_v176 = _nw110
				var _701_v177 _dafny.MultiSet
				_ = _701_v177
				_701_v177 = _dafny.MultiSetOf(_456_v20)
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_700_v176), 0))
				_ = _index109
				(_700_v176).ArraySet1((_701_v177).IsProperSubsetOf(_dafny.MultiSetOf(_456_v20, _456_v20)), (_index109).Int())
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_700_v176), 0))
				_ = _index110
				(_700_v176).ArraySet1(_456_v20, (_index110).Int())
				var _702_v178 _dafny.Map
				_ = _702_v178
				_702_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _456_v20)
				_702_v178 = _702_v178
			} else {
				var _703___mcc_h24 D2 = _source8.Get_().(D2_DC6).Cf11
				_ = _703___mcc_h24
				var _704_cf11 D2 = _703___mcc_h24
				_ = _704_cf11
				var _705_v179 _dafny.Sequence
				_ = _705_v179
				_705_v179 = _dafny.SeqOf(_456_v20, _456_v20)
				var _706_v180 _dafny.Map
				_ = _706_v180
				_706_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_657_v157, _dafny.IntOfInt64(68))
				var _707_v181 D15
				_ = _707_v181
				_707_v181 = Companion_D15_.Create_DC38_(_657_v157)
				var _708_v182 _dafny.MultiSet
				_ = _708_v182
				_708_v182 = _dafny.MultiSetOf(p0, _dafny.IntOfInt64(317))
				var _709_v183 _dafny.Array
				_ = _709_v183
				var _nwElement0_21 _dafny.Map = _706_v180
				_ = _nwElement0_21
				var _nw111 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(9))
				_ = _nw111
				(_nw111).ArraySet1(_nwElement0_21, 0)
				(_nw111).ArraySet1(Companion_Default___.Fm52(true, _707_v181, _456_v20, globalState), 1)
				(_nw111).ArraySet1(_706_v180, 2)
				(_nw111).ArraySet1(_706_v180, 3)
				(_nw111).ArraySet1(_706_v180, 4)
				(_nw111).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_657_v157, Companion_Default___.Fm44(_456_v20, globalState)), 5)
				(_nw111).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_657_v157, (func() _dafny.Int {
					if (_708_v182).Contains(p0) {
						return (_708_v182).Multiplicity(p0)
					}
					return p0
				})()), 6)
				(_nw111).ArraySet1(_706_v180, 7)
				(_nw111).ArraySet1(_706_v180, 8)
				_709_v183 = _nw111
				var _710_v184 T1
				_ = _710_v184
				var _nw112 *C1 = New_C1_()
				_ = _nw112
				_nw112.Ctor__(_dafny.IntOfUint32((_705_v179).Cardinality()), _709_v183, _457_v21, Companion_D0_.Create_DC1_(p0))
				_710_v184 = _nw112
				var _711_v185 _dafny.Map
				_ = _711_v185
				_711_v185 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_710_v184, _dafny.CodePoint('b'))
				_711_v185 = (_711_v185).Update(_710_v184, _dafny.CodePoint('e'))
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_660_v160), 0))
				_ = _index111
				(_660_v160).ArraySet1(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32((_705_v179).Cardinality())), (_index111).Int())
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_660_v160), 0))
				_ = _index112
				(_660_v160).ArraySet1((_dafny.IntOfInt64(602)).Minus(_dafny.IntOfInt64(98)), (_index112).Int())
				var _712_v186 *C0
				_ = _712_v186
				var _nw113 *C0 = New_C0_()
				_ = _nw113
				_nw113.Ctor__(_428_v0)
				_712_v186 = _nw113
				var _713_v187 _dafny.Array
				_ = _713_v187
				var _nw114 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw114
				_713_v187 = _nw114
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_713_v187), 0))
				_ = _index113
				(_713_v187).ArraySet1((func() bool {
					if _456_v20 {
						return _456_v20
					}
					return _456_v20
				})(), (_index113).Int())
				var _714_v188 _dafny.MultiSet
				_ = _714_v188
				_714_v188 = _dafny.MultiSetOf(_705_v179)
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_660_v160), 0))
				_ = _index114
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_713_v187), 0))
				_ = _index115
				var _rhs71 _dafny.Int = (_708_v182).Cardinality()
				_ = _rhs71
				var _rhs72 bool = _456_v20
				_ = _rhs72
				var _rhs73 _dafny.MultiSet = _714_v188
				_ = _rhs73
				var _lhs48 _dafny.Array = _660_v160
				_ = _lhs48
				var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_660_v160), 0))
				_ = _lhs49
				var _lhs50 _dafny.Array = _713_v187
				_ = _lhs50
				var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_713_v187), 0))
				_ = _lhs51
				(_lhs48).ArraySet1(_rhs71, (_lhs49).Int())
				(_lhs50).ArraySet1(_rhs72, (_lhs51).Int())
				_714_v188 = _rhs73
			}
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))
			_ = _index116
			(_660_v160).ArraySet1(p0, (_index116).Int())
			var _715_v189 _dafny.Map
			_ = _715_v189
			_715_v189 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0)
			var _716_v190 _dafny.Sequence
			_ = _716_v190
			_716_v190 = _dafny.SeqOf(_715_v189, _715_v189, _715_v189, _715_v189)
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))
			_ = _index117
			(_660_v160).ArraySet1((p0).Plus((Companion_Default___.Fm53((_716_v190).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_716_v190).Cardinality()))).Uint32()).(_dafny.Map), globalState)).Cardinality()), (_index117).Int())
			var _717_v191 _dafny.Map
			_ = _717_v191
			_717_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _662_v162)
			var _718_v192 _dafny.Map
			_ = _718_v192
			_718_v192 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v20, _717_v191)
			var _719_v193 _dafny.Map
			_ = _719_v193
			_719_v193 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_660_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))).Int()).(_dafny.Int))
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))
			_ = _index118
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))
			_ = _index119
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))
			_ = _index120
			var _rhs74 _dafny.Int = (_660_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))).Int()).(_dafny.Int)
			_ = _rhs74
			var _rhs75 _dafny.Int = (func() _dafny.Int {
				if !(_715_v189).Equals(_715_v189) {
					return (_dafny.IntOfInt64(52)).Plus(_dafny.IntOfInt64(499))
				}
				return p0
			})()
			_ = _rhs75
			var _rhs76 bool = ((_660_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))).Int()).(_dafny.Int)).Cmp((_660_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))).Int()).(_dafny.Int)) > 0
			_ = _rhs76
			var _rhs77 bool = !(_717_v191).Equals((func() _dafny.Map {
				if (_718_v192).Contains(_456_v20) {
					return (_718_v192).Get(_456_v20).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_660_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))).Int()).(_dafny.Int), Companion_D2_.Create_DC4_(_dafny.IntOfInt64(438), _456_v20, _661_v161))
			})())
			_ = _rhs77
			var _rhs78 _dafny.Int = (p0).Times(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_719_v193).Contains((_dafny.MultiSetFromSeq(_457_v21)).Cardinality()) {
					return (_719_v193).Get((_dafny.MultiSetFromSeq(_457_v21)).Cardinality()).(_dafny.Int)
				}
				return p0
			})(), (_660_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))).Int()).(_dafny.Int)))
			_ = _rhs78
			var _lhs52 _dafny.Array = _660_v160
			_ = _lhs52
			var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))
			_ = _lhs53
			var _lhs54 _dafny.Array = _660_v160
			_ = _lhs54
			var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))
			_ = _lhs55
			var _lhs56 _dafny.Array = _660_v160
			_ = _lhs56
			var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))
			_ = _lhs57
			(_lhs52).ArraySet1(_rhs74, (_lhs53).Int())
			(_lhs54).ArraySet1(_rhs75, (_lhs55).Int())
			_456_v20 = _rhs76
			_456_v20 = _rhs77
			(_lhs56).ArraySet1(_rhs78, (_lhs57).Int())
			var _720_v194 _dafny.Map
			_ = _720_v194
			_720_v194 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm44(!(_456_v20), globalState), _456_v20)
			var _721_v195 _dafny.Set
			_ = _721_v195
			_721_v195 = _dafny.SetOf(_657_v157)
			_456_v20 = (func() bool {
				if (_720_v194).Contains(p0) {
					return (_720_v194).Get(p0).(bool)
				}
				return !((_dafny.SetOf(_657_v157, _657_v157, _657_v157, _657_v157)).IsProperSubsetOf(_721_v195))
			})()
			var _722_v196 _dafny.Sequence
			_ = _722_v196
			_722_v196 = _dafny.SeqOf(_456_v20)
			var _723_v197 _dafny.Set
			_ = _723_v197
			_723_v197 = _dafny.SetOf(_456_v20)
			var _724_v198 _dafny.MultiSet
			_ = _724_v198
			_724_v198 = _dafny.MultiSetOf(Companion_Default___.Fm44(_456_v20, globalState))
			var _725_v199 _dafny.Sequence
			_ = _725_v199
			_725_v199 = _dafny.SeqOf(Companion_D13_.Create_DC32_(_428_v0, _456_v20, (_660_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))).Int()).(_dafny.Int), _660_v160))
			var _726_v200 _dafny.Array
			_ = _726_v200
			var _nwElement0_22 bool = (func() bool {
				if _456_v20 {
					return _456_v20
				}
				return _456_v20
			})()
			_ = _nwElement0_22
			var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(20))
			_ = _nw115
			(_nw115).ArraySet1(_nwElement0_22, 0)
			(_nw115).ArraySet1(true, 1)
			(_nw115).ArraySet1((_722_v196).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_660_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_722_v196).Cardinality()))).Uint32()).(bool), 2)
			(_nw115).ArraySet1(!(_456_v20), 3)
			(_nw115).ArraySet1(_456_v20, 4)
			(_nw115).ArraySet1(false, 5)
			(_nw115).ArraySet1((_723_v197).IsDisjointFrom(_dafny.SetOf(_456_v20, _456_v20)), 6)
			(_nw115).ArraySet1(false, 7)
			(_nw115).ArraySet1(_456_v20, 8)
			(_nw115).ArraySet1(_456_v20, 9)
			(_nw115).ArraySet1((_dafny.MultiSetOf((func() _dafny.Int {
				if (_715_v189).Contains(_456_v20) {
					return (_715_v189).Get(_456_v20).(_dafny.Int)
				}
				return p0
			})(), _dafny.IntOfUint32((_725_v199).Cardinality()))).IsProperSubsetOf(_724_v198), 10)
			(_nw115).ArraySet1(_456_v20, 11)
			(_nw115).ArraySet1((p0).Cmp((_660_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))).Int()).(_dafny.Int)) >= 0, 12)
			(_nw115).ArraySet1((func() bool {
				if (_720_v194).Contains((_660_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))).Int()).(_dafny.Int)) {
					return (_720_v194).Get((_660_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_660_v160), 0))).Int()).(_dafny.Int)).(bool)
				}
				return _456_v20
			})(), 13)
			(_nw115).ArraySet1((_this).Fm42(_456_v20, globalState), 14)
			(_nw115).ArraySet1((func() bool {
				if _456_v20 {
					return _456_v20
				}
				return _456_v20
			})(), 15)
			(_nw115).ArraySet1((func() bool {
				if _456_v20 {
					return _456_v20
				}
				return _456_v20
			})(), 16)
			(_nw115).ArraySet1(!(_456_v20), 17)
			(_nw115).ArraySet1(_456_v20, 18)
			(_nw115).ArraySet1(_dafny.Companion_Sequence_.Contains(_428_v0, _657_v157), 19)
			_726_v200 = _nw115
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_726_v200), 0))
			_ = _index121
			(_726_v200).ArraySet1(_456_v20, (_index121).Int())
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_726_v200), 0))
			_ = _index122
			(_726_v200).ArraySet1((p0).Cmp(p0) == 0, (_index122).Int())
		}
		var _727_v201 D4
		_ = _727_v201
		_727_v201 = Companion_D4_.Create_DC10_(p0, p0, p0)
		var _pat_let_tv10 = _456_v20
		_ = _pat_let_tv10
		var _pat_let_tv11 = _457_v21
		_ = _pat_let_tv11
		var _pat_let_tv12 = _456_v20
		_ = _pat_let_tv12
		var _pat_let_tv13 = _457_v21
		_ = _pat_let_tv13
		var _pat_let_tv14 = _456_v20
		_ = _pat_let_tv14
		var _pat_let_tv15 = _456_v20
		_ = _pat_let_tv15
		var _pat_let_tv16 = _456_v20
		_ = _pat_let_tv16
		var _pat_let_tv17 = _456_v20
		_ = _pat_let_tv17
		var _pat_let_tv18 = _456_v20
		_ = _pat_let_tv18
		var _pat_let_tv19 = p0
		_ = _pat_let_tv19
		var _pat_let_tv20 = _456_v20
		_ = _pat_let_tv20
		var _pat_let_tv21 = _456_v20
		_ = _pat_let_tv21
		var _pat_let_tv22 = p0
		_ = _pat_let_tv22
		var _pat_let_tv23 = _456_v20
		_ = _pat_let_tv23
		var _pat_let_tv24 = p0
		_ = _pat_let_tv24
		var _pat_let_tv25 = p0
		_ = _pat_let_tv25
		r0 = func(_source9 D4) _dafny.Map {
			if _source9.Is_DC9() {
				var _728___mcc_h25 _dafny.MultiSet = _source9.Get_().(D4_DC9).Cf14
				_ = _728___mcc_h25
				var _729___mcc_h26 _dafny.CodePoint = _source9.Get_().(D4_DC9).Cf15
				_ = _729___mcc_h26
				var _730___mcc_h27 _dafny.Int = _source9.Get_().(D4_DC9).Cf16
				_ = _730___mcc_h27
				var _731___mcc_h28 _dafny.Set = _source9.Get_().(D4_DC9).Cf17
				_ = _731___mcc_h28
				var _732___mcc_h29 _dafny.Int = _source9.Get_().(D4_DC9).Cf18
				_ = _732___mcc_h29
				var _733_cf18 _dafny.Int = _732___mcc_h29
				_ = _733_cf18
				var _734_cf17 _dafny.Set = _731___mcc_h28
				_ = _734_cf17
				var _735_cf16 _dafny.Int = _730___mcc_h27
				_ = _735_cf16
				var _736_cf15 _dafny.CodePoint = _729___mcc_h26
				_ = _736_cf15
				var _737_cf14 _dafny.MultiSet = _728___mcc_h25
				_ = _737_cf14
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv10, _dafny.MultiSetFromSeq(_pat_let_tv11))
			} else if _source9.Is_DC10() {
				var _738___mcc_h30 _dafny.Int = _source9.Get_().(D4_DC10).Cf19
				_ = _738___mcc_h30
				var _739___mcc_h31 _dafny.Int = _source9.Get_().(D4_DC10).Cf20
				_ = _739___mcc_h31
				var _740___mcc_h32 _dafny.Int = _source9.Get_().(D4_DC10).Cf21
				_ = _740___mcc_h32
				var _741_cf21 _dafny.Int = _740___mcc_h32
				_ = _741_cf21
				var _742_cf20 _dafny.Int = _739___mcc_h31
				_ = _742_cf20
				var _743_cf19 _dafny.Int = _738___mcc_h30
				_ = _743_cf19
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv12, _dafny.MultiSetFromSeq((Companion_D2_.Create_DC5_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(842))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}(func(_744_i24 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				}))).Cardinality()), _pat_let_tv13, _pat_let_tv14)).Dtor_cf9()))
			} else if _source9.Is_DC11() {
				var _745___mcc_h33 _dafny.Int = _source9.Get_().(D4_DC11).Cf22
				_ = _745___mcc_h33
				var _746___mcc_h34 bool = _source9.Get_().(D4_DC11).Cf23
				_ = _746___mcc_h34
				var _747___mcc_h35 bool = _source9.Get_().(D4_DC11).Cf24
				_ = _747___mcc_h35
				var _748___mcc_h36 bool = _source9.Get_().(D4_DC11).Cf25
				_ = _748___mcc_h36
				var _749___mcc_h37 bool = _source9.Get_().(D4_DC11).Cf26
				_ = _749___mcc_h37
				var _750_cf26 bool = _749___mcc_h37
				_ = _750_cf26
				var _751_cf25 bool = _748___mcc_h36
				_ = _751_cf25
				var _752_cf24 bool = _747___mcc_h35
				_ = _752_cf24
				var _753_cf23 bool = _746___mcc_h34
				_ = _753_cf23
				var _754_cf22 _dafny.Int = _745___mcc_h33
				_ = _754_cf22
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_750_cf26, _dafny.MultiSetOf(_754_cf22))
			} else if _source9.Is_DC8() {
				var _755___mcc_h38 _dafny.Sequence = _source9.Get_().(D4_DC8).Cf13
				_ = _755___mcc_h38
				var _756_cf13 _dafny.Sequence = _755___mcc_h38
				_ = _756_cf13
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv15, _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), _pat_let_tv16), _pat_let_tv17)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv18, _dafny.MultiSetOf(_pat_let_tv19, _dafny.IntOfUint32((_dafny.SeqOf(_pat_let_tv20, _pat_let_tv21)).Cardinality()))))
			} else {
				var _757___mcc_h39 D4 = _source9.Get_().(D4_DC12).Cf27
				_ = _757___mcc_h39
				var _758_cf27 D4 = _757___mcc_h39
				_ = _758_cf27
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.MultiSetOf(_pat_let_tv22))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv23, _dafny.MultiSetOf(_pat_let_tv24)))
			}
		}((func() D4 {
			if _456_v20 {
				return func(_pat_let14_0 D4) D4 {
					return func(_759_dt__update__tmp_h4 D4) D4 {
						return func(_pat_let15_0 _dafny.Int) D4 {
							return func(_760_dt__update_hcf21_h0 _dafny.Int) D4 {
								return Companion_D4_.Create_DC10_((_759_dt__update__tmp_h4).Dtor_cf19(), (_759_dt__update__tmp_h4).Dtor_cf20(), _760_dt__update_hcf21_h0)
							}(_pat_let15_0)
						}(_pat_let_tv25)
					}(_pat_let14_0)
				}(_727_v201)
			}
			return _727_v201
		})())
		return r0
	}
}
func (_this *C2) M14(p0 bool, p1 bool, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		r0 = Companion_Default___.Fm44(p1, globalState)
		var _761_v0 _dafny.Sequence
		_ = _761_v0
		_761_v0 = _dafny.UnicodeSeqOfUtf8Bytes("wnm")
		var _762_v1 _dafny.Map
		_ = _762_v1
		_762_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, true)
		var _hi3 _dafny.Int = (_dafny.IntOfInt64(293)).Plus((_762_v1).Cardinality())
		_ = _hi3
		for _763_i0 := (_dafny.Zero).Minus(_dafny.IntOfUint32((_761_v0).Cardinality())); _763_i0.Cmp(_hi3) < 0; _763_i0 = _763_i0.Plus(_dafny.One) {
			var _764_v2 _dafny.Array
			_ = _764_v2
			var _nw116 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw116
			_764_v2 = _nw116
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_764_v2), 0))
			_ = _index123
			(_764_v2).ArraySet1(p1, (_index123).Int())
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_764_v2), 0))
			_ = _index124
			(_764_v2).ArraySet1((p1) || ((func() bool {
				if p0 {
					return false
				}
				return p1
			})()), (_index124).Int())
			var _765_v3 _dafny.Set
			_ = _765_v3
			_765_v3 = _dafny.SetOf(p0)
			var _766_v4 _dafny.Array
			_ = _766_v4
			var _nwElement0_23 _dafny.Set = _765_v3
			_ = _nwElement0_23
			var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(18))
			_ = _nw117
			(_nw117).ArraySet1(_nwElement0_23, 0)
			(_nw117).ArraySet1(_765_v3, 1)
			(_nw117).ArraySet1(_765_v3, 2)
			(_nw117).ArraySet1(_765_v3, 3)
			(_nw117).ArraySet1(_765_v3, 4)
			(_nw117).ArraySet1(_765_v3, 5)
			(_nw117).ArraySet1(_765_v3, 6)
			(_nw117).ArraySet1(_765_v3, 7)
			(_nw117).ArraySet1(_dafny.SetOf((_764_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_764_v2), 0))).Int()).(bool), p0), 8)
			(_nw117).ArraySet1(_dafny.SetOf(p0), 9)
			(_nw117).ArraySet1(_765_v3, 10)
			(_nw117).ArraySet1(_765_v3, 11)
			(_nw117).ArraySet1(_765_v3, 12)
			(_nw117).ArraySet1(_dafny.SetOf((_764_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_764_v2), 0))).Int()).(bool)), 13)
			(_nw117).ArraySet1(_765_v3, 14)
			(_nw117).ArraySet1(_765_v3, 15)
			(_nw117).ArraySet1(_765_v3, 16)
			(_nw117).ArraySet1(_765_v3, 17)
			_766_v4 = _nw117
			var _767_v5 D0
			_ = _767_v5
			_767_v5 = Companion_D0_.Create_DC0_(_766_v4, _dafny.IntOfInt64(2))
			var _768_v6 _dafny.Sequence
			_ = _768_v6
			_768_v6 = _dafny.SeqOf(_767_v5, _767_v5)
			var _769_v7 D4
			_ = _769_v7
			_769_v7 = Companion_D4_.Create_DC8_(_768_v6)
			_769_v7 = _769_v7
			var _770_v8 _dafny.CodePoint
			_ = _770_v8
			_770_v8 = _dafny.CodePoint('n')
			var _771_v9 _dafny.Map
			_ = _771_v9
			_771_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_770_v8, p3)
			var _772_v11 _dafny.Map
			_ = _772_v11
			_772_v11 = (func() _dafny.Map {
				if p1 {
					return _771_v9
				}
				return func() _dafny.Map {
					var _coll40 = _dafny.NewMapBuilder()
					_ = _coll40
					for _iter41 := _dafny.Iterate((_761_v0).Elements()); ; {
						_compr_40, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _773_v10 _dafny.CodePoint
						_773_v10 = interface{}(_compr_40).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_761_v0, _773_v10) {
							_coll40.Add(_773_v10, p3)
						}
					}
					return _coll40.ToMap()
				}()
			})()
			_772_v11 = Companion_Default___.Fm54((_764_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_764_v2), 0))).Int()).(bool), _dafny.SeqOf(p1, (_764_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_764_v2), 0))).Int()).(bool)), globalState)
			r0 = (_dafny.Zero).Minus(p3)
		}
		var _774_v12 bool
		_ = _774_v12
		_774_v12 = false
		_774_v12 = p0
		var _hi4 _dafny.Int = (func() _dafny.Int {
			if p1 {
				return p3
			}
			return p3
		})()
		_ = _hi4
		for _775_i1 := p3; _775_i1.Cmp(_hi4) < 0; _775_i1 = _775_i1.Plus(_dafny.One) {
			var _776_v13 _dafny.CodePoint
			_ = _776_v13
			_776_v13 = _dafny.CodePoint('g')
			_776_v13 = (func() _dafny.CodePoint {
				if _774_v12 {
					return _776_v13
				}
				return _776_v13
			})()
			_774_v12 = _dafny.Companion_Sequence_.IsPrefixOf(_761_v0, _761_v0)
			_761_v0 = (func() _dafny.Sequence {
				if p1 {
					return _dafny.UnicodeSeqOfUtf8Bytes("w")
				}
				return _761_v0
			})()
			var _777_v14 _dafny.Array
			_ = _777_v14
			var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw118
			_777_v14 = _nw118
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(783), _dafny.ArrayLen((_777_v14), 0))
			_ = _index125
			(_777_v14).ArraySet1(p3, (_index125).Int())
			var _778_v15 _dafny.Sequence
			_ = _778_v15
			_778_v15 = _dafny.SeqOf(false)
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(783), _dafny.ArrayLen((_777_v14), 0))
			_ = _index126
			var _rhs79 _dafny.Int = (((_762_v1).Merge(_762_v1)).Cardinality()).Minus(_775_i1)
			_ = _rhs79
			var _rhs80 _dafny.Int = _775_i1
			_ = _rhs80
			var _rhs81 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_778_v15, _dafny.SeqOf((Companion_D14_.Create_DC35_(p3, _dafny.UnicodeSeqOfUtf8Bytes("sia"), _775_i1, _774_v12, p0)).Dtor_cf67(), p0, p1, _774_v12))
			_ = _rhs81
			var _lhs58 _dafny.Array = _777_v14
			_ = _lhs58
			var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(783), _dafny.ArrayLen((_777_v14), 0))
			_ = _lhs59
			(_lhs58).ArraySet1(_rhs79, (_lhs59).Int())
			r0 = _rhs80
			_778_v15 = _rhs81
		}
		var _779_v16 _dafny.Set
		_ = _779_v16
		_779_v16 = _dafny.SetOf(p3, p3)
		var _780_v17 D6
		_ = _780_v17
		_780_v17 = Companion_D6_.Create_DC15_((_779_v16).Cardinality(), p3, p0)
		var _781_v18 D6
		_ = _781_v18
		_781_v18 = Companion_D6_.Create_DC16_(_780_v17)
		var _782_v19 _dafny.MultiSet
		_ = _782_v19
		_782_v19 = _dafny.MultiSetOf(_781_v18, Companion_D6_.Create_DC16_((Companion_D6_.Create_DC16_(_780_v17)).Dtor_cf33()))
		_774_v12 = !(!_dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
			if !(p1) {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(204))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_783_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_784_i2 _dafny.Int) _dafny.Int {
						return _783_p3
					}
				})(p3)))
			}
			return _dafny.SeqOf((_782_v19).Cardinality(), p3, p3, p3)
		})(), (func() _dafny.Int {
			if p0 {
				return p3
			}
			return p3
		})()))
		_761_v0 = _761_v0
		r0 = p3
		return r0
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) Ctor__() {
	{
	}
}
func (_this *C3) Fm5(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 bool, globalState *GlobalState) bool {
	{
		return !(!(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hi")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('x'), _dafny.CodePoint('f'))).Cardinality()))).Cardinality())), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fqgtkbdg")).Cardinality()), (_dafny.SetOf(_dafny.IntOfInt64(-285), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), _dafny.CodePoint('k'))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-394), _dafny.IntOfInt64(772))).Cardinality())).Cardinality())), (_dafny.IntOfInt64(622)).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-779)))))))
	}
}
func (_this *C3) M1(p0 _dafny.Int, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _785_v0 bool
		_ = _785_v0
		_785_v0 = false
		if _785_v0 {
			var _786_v1 bool
			_ = _786_v1
			var _787_v2 _dafny.Int
			_ = _787_v2
			var _788_v3 _dafny.Int
			_ = _788_v3
			var _789_v4 bool
			_ = _789_v4
			var _out21 bool
			_ = _out21
			var _out22 _dafny.Int
			_ = _out22
			var _out23 _dafny.Int
			_ = _out23
			var _out24 bool
			_ = _out24
			_out21, _out22, _out23, _out24 = (_this).M12(!(_785_v0) || (_785_v0), globalState)
			_786_v1 = _out21
			_787_v2 = _out22
			_788_v3 = _out23
			_789_v4 = _out24
			if _785_v0 {
				var _790_v5 _dafny.Array
				_ = _790_v5
				var _nw119 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(26))
				_ = _nw119
				_790_v5 = _nw119
				var _791_v6 _dafny.CodePoint
				_ = _791_v6
				_791_v6 = _dafny.CodePoint('g')
				var _792_v7 D0
				_ = _792_v7
				_792_v7 = Companion_D0_.Create_DC1_(p0)
				var _793_v8 _dafny.Sequence
				_ = _793_v8
				_793_v8 = _dafny.UnicodeSeqOfUtf8Bytes("rnnc")
				var _794_v9 _dafny.Map
				_ = _794_v9
				_794_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm5(p0, false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_791_v6, _792_v7), _789_v4, globalState), _793_v8)
				var _795_v10 _dafny.Set
				_ = _795_v10
				_795_v10 = _dafny.SetOf(p0, _788_v3)
				var _796_v11 _dafny.MultiSet
				_ = _796_v11
				_796_v11 = _dafny.MultiSetOf(p0)
				var _797_v12 _dafny.Sequence
				_ = _797_v12
				_797_v12 = _dafny.SeqOf((_794_v9).Cardinality(), (_795_v10).Cardinality(), (_796_v11).Cardinality(), Companion_Default___.Fm32(p0, globalState), _788_v3)
				var _798_v13 T1
				_ = _798_v13
				var _nw120 *C1 = New_C1_()
				_ = _nw120
				_nw120.Ctor__((_dafny.Zero).Minus(_dafny.IntOfInt64(-771)), _790_v5, _797_v12, _792_v7)
				_798_v13 = _nw120
				var _799_v14 _dafny.Sequence
				_ = _799_v14
				_799_v14 = _dafny.SeqOf(_798_v13)
				_799_v14 = _799_v14
				var _800_v15 _dafny.Array
				_ = _800_v15
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_17
				var _nw121 _dafny.Array
				_ = _nw121
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw121 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.Int = (func(_801_v0 bool) func(_dafny.Int) _dafny.Int {
						return func(_802_i0 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_802_i0, _dafny.IntOfUint32((_dafny.SeqOf(!(_801_v0))).Cardinality()))
						}
					})(_785_v0)
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw121 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw121).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw121).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_800_v15 = _nw121
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_800_v15), 0))
				_ = _index127
				(_800_v15).ArraySet1(_787_v2, (_index127).Int())
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_800_v15), 0))
				_ = _index128
				(_800_v15).ArraySet1(_787_v2, (_index128).Int())
				var _803_v16 _dafny.Set
				_ = _803_v16
				_803_v16 = _dafny.SetOf(_789_v4)
				var _804_v17 D12
				_ = _804_v17
				_804_v17 = Companion_D12_.Create_DC27_(_803_v16)
				var _805_v18 _dafny.Map
				_ = _805_v18
				_805_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_789_v4, _803_v16)
				_804_v17 = Companion_D12_.Create_DC27_((func() _dafny.Set {
					if (_805_v18).Contains(_785_v0) {
						return (_805_v18).Get(_785_v0).(_dafny.Set)
					}
					return _dafny.SetOf(_785_v0)
				})())
				_800_v15 = _800_v15
				var _806_v19 D6
				_ = _806_v19
				_806_v19 = Companion_D6_.Create_DC14_(_793_v8)
				var _807_v20 _dafny.MultiSet
				_ = _807_v20
				_807_v20 = _dafny.MultiSetOf(_786_v1)
				var _808_v21 _dafny.Map
				_ = _808_v21
				_808_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v3, (func() _dafny.Int {
					if (_807_v20).Contains(_789_v4) {
						return (_807_v20).Multiplicity(_789_v4)
					}
					return _dafny.IntOfInt64(262)
				})())
				var _809_v22 _dafny.Array
				_ = _809_v22
				var _nwElement0_24 D6 = _806_v19
				_ = _nwElement0_24
				var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(12))
				_ = _nw122
				(_nw122).ArraySet1(_nwElement0_24, 0)
				(_nw122).ArraySet1(_806_v19, 1)
				(_nw122).ArraySet1(func(_pat_let16_0 D6) D6 {
					return func(_811_dt__update__tmp_h0 D6) D6 {
						return func(_pat_let17_0 _dafny.Sequence) D6 {
							return func(_813_dt__update_hcf29_h0 _dafny.Sequence) D6 {
								return Companion_D6_.Create_DC14_(_813_dt__update_hcf29_h0)
							}(_pat_let17_0)
						}(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(287))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg64 _dafny.Int) interface{} {
								return coer64(arg64)
							}
						}(func(_812_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('i')
						})))
					}(_pat_let16_0)
				}(Companion_Default___.Fm33((func() _dafny.Int {
					if (_808_v21).Contains(_788_v3) {
						return (_808_v21).Get(_788_v3).(_dafny.Int)
					}
					return p0
				})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(911))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}(func(_810_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				})), globalState)), 2)
				(_nw122).ArraySet1(_806_v19, 3)
				(_nw122).ArraySet1(_806_v19, 4)
				(_nw122).ArraySet1(_806_v19, 5)
				(_nw122).ArraySet1(_806_v19, 6)
				(_nw122).ArraySet1(_806_v19, 7)
				(_nw122).ArraySet1(Companion_D6_.Create_DC14_(_793_v8), 8)
				(_nw122).ArraySet1(_806_v19, 9)
				(_nw122).ArraySet1(Companion_D6_.Create_DC14_(_793_v8), 10)
				(_nw122).ArraySet1(Companion_Default___.Fm33((_800_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_800_v15), 0))).Int()).(_dafny.Int), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(190))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}((func(_814_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_815_i3 _dafny.Int) _dafny.CodePoint {
						return _814_v6
					}
				})(_791_v6))), globalState), 11)
				_809_v22 = _nw122
				var _pat_let_tv26 = _793_v8
				_ = _pat_let_tv26
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_809_v22), 0))
				_ = _index129
				(_809_v22).ArraySet1(func(_pat_let18_0 D6) D6 {
					return func(_816_dt__update__tmp_h1 D6) D6 {
						return func(_pat_let19_0 _dafny.Sequence) D6 {
							return func(_817_dt__update_hcf29_h1 _dafny.Sequence) D6 {
								return Companion_D6_.Create_DC14_(_817_dt__update_hcf29_h1)
							}(_pat_let19_0)
						}(_pat_let_tv26)
					}(_pat_let18_0)
				}(_806_v19), (_index129).Int())
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_809_v22), 0))
				_ = _index130
				(_809_v22).ArraySet1(_806_v19, (_index130).Int())
			} else {
				var _818_v23 _dafny.Array
				_ = _818_v23
				var _nw123 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
				_ = _nw123
				_818_v23 = _nw123
				_818_v23 = _818_v23
				_788_v3 = (_dafny.Zero).Minus(p0)
				r1 = _789_v4
				var _819_v24 _dafny.Array
				_ = _819_v24
				var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
				_ = _nw124
				_819_v24 = _nw124
				var _820_v25 _dafny.Sequence
				_ = _820_v25
				_820_v25 = _dafny.UnicodeSeqOfUtf8Bytes("gvm")
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_819_v24), 0))
				_ = _index131
				(_819_v24).ArraySet1(_820_v25, (_index131).Int())
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_819_v24), 0))
				_ = _index132
				(_819_v24).ArraySet1(_820_v25, (_index132).Int())
				var _821_v26 _dafny.Map
				_ = _821_v26
				_821_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_789_v4) == (_785_v0), _785_v0)
				_821_v26 = (_821_v26).Merge(_821_v26)
			}
			var _822_v27 _dafny.Sequence
			_ = _822_v27
			_822_v27 = _dafny.UnicodeSeqOfUtf8Bytes("opafnutak")
			var _823_v28 _dafny.Sequence
			_ = _823_v28
			_823_v28 = _dafny.SeqOf(_822_v27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(620))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg66 _dafny.Int) interface{} {
					return coer66(arg66)
				}
			}(func(_824_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			})), _822_v27, _822_v27, _822_v27)
			var _825_v29 *C0
			_ = _825_v29
			var _nw125 *C0 = New_C0_()
			_ = _nw125
			_nw125.Ctor__((_823_v28).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.IntOfUint32((_823_v28).Cardinality()))).Uint32()).(_dafny.Sequence))
			_825_v29 = _nw125
			if true {
				var _826_v30 _dafny.Sequence
				_ = _826_v30
				_826_v30 = _dafny.SeqOf(_789_v4)
				var _827_v31 _dafny.Sequence
				_ = _827_v31
				_827_v31 = _dafny.SeqOf(_826_v30, _826_v30)
				var _828_v32 _dafny.Map
				_ = _828_v32
				_828_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _788_v3)
				var _829_v33 _dafny.Sequence
				_ = _829_v33
				_829_v33 = _dafny.SeqOf(_dafny.IntOfUint32((_822_v27).Cardinality()), _dafny.IntOfInt64(553), (func() _dafny.Int {
					if (_828_v32).Contains(_dafny.IntOfInt64(-361)) {
						return (_828_v32).Get(_dafny.IntOfInt64(-361)).(_dafny.Int)
					}
					return p0
				})(), (_dafny.Zero).Minus(_788_v3), _787_v2)
				var _830_v34 _dafny.Set
				_ = _830_v34
				_830_v34 = _dafny.SetOf(p0)
				var _831_v35 _dafny.Array
				_ = _831_v35
				var _nwElement0_25 _dafny.Int = _788_v3
				_ = _nwElement0_25
				var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(28))
				_ = _nw126
				(_nw126).ArraySet1(_nwElement0_25, 0)
				(_nw126).ArraySet1(_787_v2, 1)
				(_nw126).ArraySet1(_788_v3, 2)
				(_nw126).ArraySet1(_787_v2, 3)
				(_nw126).ArraySet1(_787_v2, 4)
				(_nw126).ArraySet1(p0, 5)
				(_nw126).ArraySet1(p0, 6)
				(_nw126).ArraySet1(p0, 7)
				(_nw126).ArraySet1(_787_v2, 8)
				(_nw126).ArraySet1(p0, 9)
				(_nw126).ArraySet1(p0, 10)
				(_nw126).ArraySet1(_787_v2, 11)
				(_nw126).ArraySet1(_787_v2, 12)
				(_nw126).ArraySet1(_dafny.IntOfInt64(356), 13)
				(_nw126).ArraySet1(_787_v2, 14)
				(_nw126).ArraySet1(_788_v3, 15)
				(_nw126).ArraySet1(_dafny.IntOfInt64(-723), 16)
				(_nw126).ArraySet1((_dafny.MultiSetFromSeq((_827_v31).Select((Companion_Default___.SafeIndex(_788_v3, _dafny.IntOfUint32((_827_v31).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality(), 17)
				(_nw126).ArraySet1((_829_v33).Select((Companion_Default___.SafeIndex(_788_v3, _dafny.IntOfUint32((_829_v33).Cardinality()))).Uint32()).(_dafny.Int), 18)
				(_nw126).ArraySet1(_788_v3, 19)
				(_nw126).ArraySet1(p0, 20)
				(_nw126).ArraySet1(_788_v3, 21)
				(_nw126).ArraySet1(_787_v2, 22)
				(_nw126).ArraySet1((_830_v34).Cardinality(), 23)
				(_nw126).ArraySet1(_788_v3, 24)
				(_nw126).ArraySet1(p0, 25)
				(_nw126).ArraySet1(_787_v2, 26)
				(_nw126).ArraySet1(p0, 27)
				_831_v35 = _nw126
				var _832_v36 _dafny.Map
				_ = _832_v36
				_832_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _831_v35)
				var _833_v37 D12
				_ = _833_v37
				_833_v37 = Companion_D12_.Create_DC28_(_785_v0, _787_v2)
				var _834_v39 _dafny.Array
				_ = _834_v39
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_18
				var _nw127 _dafny.Array
				_ = _nw127
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw127 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Map = (func(_835_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
						return func(_836_i5 _dafny.Int) _dafny.Map {
							return func() _dafny.Map {
								var _coll41 = _dafny.NewMapBuilder()
								_ = _coll41
								for _iter42 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('s'))).Elements()); ; {
									_compr_41, _ok42 := _iter42()
									if !_ok42 {
										break
									}
									var _837_v38 _dafny.CodePoint
									_837_v38 = interface{}(_compr_41).(_dafny.CodePoint)
									if (_dafny.SetOf(_dafny.CodePoint('s'))).Contains(_837_v38) {
										_coll41.Add(_837_v38, _835_p0)
									}
								}
								return _coll41.ToMap()
							}()
						}
					})(p0)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw127 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw127).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw127).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_834_v39 = _nw127
				var _838_v40 D0
				_ = _838_v40
				_838_v40 = Companion_D0_.Create_DC1_(p0)
				var _839_v41 *C1
				_ = _839_v41
				var _nw128 *C1 = New_C1_()
				_ = _nw128
				_nw128.Ctor__(p0, _834_v39, _829_v33, _838_v40)
				_839_v41 = _nw128
				var _840_v42 _dafny.Map
				_ = _840_v42
				_840_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_839_v41, false)
				var _841_v43 _dafny.Set
				_ = _841_v43
				_841_v43 = _dafny.SetOf(!(_786_v1))
				var _842_v44 _dafny.Array
				_ = _842_v44
				var _nwElement0_26 bool = _789_v4
				_ = _nwElement0_26
				var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(20))
				_ = _nw129
				(_nw129).ArraySet1(_nwElement0_26, 0)
				(_nw129).ArraySet1((_788_v3).Cmp(_788_v3) < 0, 1)
				(_nw129).ArraySet1(_786_v1, 2)
				(_nw129).ArraySet1((_787_v2).Cmp((_832_v36).Cardinality()) < 0, 3)
				(_nw129).ArraySet1(false, 4)
				(_nw129).ArraySet1(!(_785_v0), 5)
				(_nw129).ArraySet1(!(_786_v1), 6)
				(_nw129).ArraySet1(_786_v1, 7)
				(_nw129).ArraySet1((func() bool {
					if _789_v4 {
						return (_833_v37).Dtor_cf46()
					}
					return _786_v1
				})(), 8)
				(_nw129).ArraySet1(_786_v1, 9)
				(_nw129).ArraySet1(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_840_v42, _841_v43)).Contains(_840_v42), 10)
				(_nw129).ArraySet1(_786_v1, 11)
				(_nw129).ArraySet1(_789_v4, 12)
				(_nw129).ArraySet1(!(false) || (_789_v4), 13)
				(_nw129).ArraySet1(_786_v1, 14)
				(_nw129).ArraySet1((_788_v3).Cmp(_839_v41.F9) >= 0, 15)
				(_nw129).ArraySet1(false, 16)
				(_nw129).ArraySet1((_830_v34).IsSubsetOf(_830_v34), 17)
				(_nw129).ArraySet1(!((func() bool {
					if _789_v4 {
						return _786_v1
					}
					return _785_v0
				})()), 18)
				(_nw129).ArraySet1(_786_v1, 19)
				_842_v44 = _nw129
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_842_v44), 0))
				_ = _index133
				(_842_v44).ArraySet1(_789_v4, (_index133).Int())
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_842_v44), 0))
				_ = _index134
				(_842_v44).ArraySet1(_785_v0, (_index134).Int())
				_788_v3 = _787_v2
				var _843_v45 _dafny.Map
				_ = _843_v45
				_843_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v0, (_839_v41.F9).Plus((Companion_Default___.Fm34(_786_v1, _785_v0, Companion_Default___.Fm30(_839_v41.F9, (_842_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_842_v44), 0))).Int()).(bool), (_842_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_842_v44), 0))).Int()).(bool), globalState), _dafny.IntOfInt64(369), globalState)).Dtor_cf47()))
				_788_v3 = (_843_v45).Cardinality()
				(_839_v41).F9 = _dafny.IntOfUint32((_825_v29.F8).Cardinality())
				var _844_v46 *C0
				_ = _844_v46
				var _nw130 *C0 = New_C0_()
				_ = _nw130
				_nw130.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-8))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}((func(_845_v2 _dafny.Int, _846_v30 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
					return func(_847_i6 _dafny.Int) _dafny.CodePoint {
						return (Companion_D11_.Create_DC25_(_dafny.CodePoint('w'), _845_v2, (_846_v30).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.IntOfUint32((_846_v30).Cardinality()))).Uint32()).(bool))).Dtor_cf41()
					}
				})(_787_v2, _826_v30))))
				_844_v46 = _nw130
			} else {
				var _848_v47 _dafny.MultiSet
				_ = _848_v47
				_848_v47 = _dafny.MultiSetOf(_786_v1)
				var _849_v48 _dafny.MultiSet
				_ = _849_v48
				_849_v48 = _dafny.MultiSetOf((_848_v47).Cardinality())
				var _850_v49 _dafny.Array
				_ = _850_v49
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_19
				var _nw131 _dafny.Array
				_ = _nw131
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw131 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) bool = (func(_851_v1 bool) func(_dafny.Int) bool {
						return func(_852_i7 _dafny.Int) bool {
							return _851_v1
						}
					})(_786_v1)
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw131 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw131).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw131).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_850_v49 = _nw131
				var _853_v50 D12
				_ = _853_v50
				_853_v50 = Companion_D12_.Create_DC28_(_786_v1, _dafny.IntOfInt64(524))
				var _854_v51 _dafny.CodePoint
				_ = _854_v51
				_854_v51 = _dafny.CodePoint('y')
				var _855_v52 _dafny.Map
				_ = _855_v52
				_855_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v0, _dafny.CodePoint('c'))
				var _856_v53 D0
				_ = _856_v53
				_856_v53 = Companion_D0_.Create_DC1_((_855_v52).Cardinality())
				var _857_v54 _dafny.Map
				_ = _857_v54
				_857_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_854_v51, _856_v53)
				var _858_v55 _dafny.Map
				_ = _858_v55
				_858_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_786_v1, Companion_Default___.Fm30((_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), (_this).Fm5(_787_v2, _789_v4, _857_v54, _785_v0, globalState), _789_v4, globalState))
				var _859_v56 _dafny.Map
				_ = _859_v56
				_859_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_787_v2), _786_v1)
				var _860_v57 _dafny.Map
				_ = _860_v57
				_860_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v3, _786_v1)).Cardinality())
				var _861_v58 _dafny.Array
				_ = _861_v58
				var _nwElement0_27 _dafny.Int = _787_v2
				_ = _nwElement0_27
				var _nw132 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(22))
				_ = _nw132
				(_nw132).ArraySet1(_nwElement0_27, 0)
				(_nw132).ArraySet1(_787_v2, 1)
				(_nw132).ArraySet1(_788_v3, 2)
				(_nw132).ArraySet1(_787_v2, 3)
				(_nw132).ArraySet1((Companion_Default___.Fm32(p0, globalState)).Plus(p0), 4)
				(_nw132).ArraySet1(Companion_Default___.SafeDivisionInt(_787_v2, (_dafny.Zero).Minus(_787_v2)), 5)
				(_nw132).ArraySet1(_787_v2, 6)
				(_nw132).ArraySet1((_dafny.Zero).Minus(p0), 7)
				(_nw132).ArraySet1(_788_v3, 8)
				(_nw132).ArraySet1((_849_v48).Cardinality(), 9)
				(_nw132).ArraySet1(p0, 10)
				(_nw132).ArraySet1((_dafny.MultiSetOf(_850_v49, _850_v49, _850_v49)).Cardinality(), 11)
				(_nw132).ArraySet1(p0, 12)
				(_nw132).ArraySet1((_853_v50).Dtor_cf47(), 13)
				(_nw132).ArraySet1((_858_v55).Cardinality(), 14)
				(_nw132).ArraySet1((_787_v2).Plus((_859_v56).Cardinality()), 15)
				(_nw132).ArraySet1(p0, 16)
				(_nw132).ArraySet1(_787_v2, 17)
				(_nw132).ArraySet1((_860_v57).Cardinality(), 18)
				(_nw132).ArraySet1(p0, 19)
				(_nw132).ArraySet1((_dafny.Zero).Minus(p0), 20)
				(_nw132).ArraySet1((_787_v2).Minus(_787_v2), 21)
				_861_v58 = _nw132
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_861_v58), 0))
				_ = _index135
				(_861_v58).ArraySet1((func() _dafny.Int {
					if !(_786_v1) {
						return _788_v3
					}
					return _787_v2
				})(), (_index135).Int())
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_861_v58), 0))
				_ = _index136
				(_861_v58).ArraySet1((Companion_Default___.Fm32(_dafny.IntOfInt64(174), globalState)).Plus(p0), (_index136).Int())
				var _862_v59 _dafny.Array
				_ = _862_v59
				var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(8))
				_ = _nw133
				_862_v59 = _nw133
				var _863_v60 _dafny.MultiSet
				_ = _863_v60
				_863_v60 = _dafny.MultiSetOf(_854_v51)
				var _rhs82 _dafny.Int = p0
				_ = _rhs82
				var _rhs83 _dafny.Array = _862_v59
				_ = _rhs83
				var _rhs84 _dafny.MultiSet = _863_v60
				_ = _rhs84
				_788_v3 = _rhs82
				_862_v59 = _rhs83
				_863_v60 = _rhs84
				var _864_v61 _dafny.Sequence
				_ = _864_v61
				_864_v61 = _dafny.SeqOf(_786_v1, _789_v4, _785_v0)
				_864_v61 = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if _786_v1 {
						return _864_v61
					}
					return _864_v61
				})(), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if _786_v1 {
						return _864_v61
					}
					return _864_v61
				})()).Cardinality()))).Uint32(), ((_859_v56).Cardinality()).Cmp((_861_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_861_v58), 0))).Int()).(_dafny.Int)) >= 0)
				var _865_v62 _dafny.Map
				_ = _865_v62
				_865_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(689), _854_v51)
				var _866_v63 _dafny.Map
				_ = _866_v63
				_866_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_865_v62, _787_v2)
				_866_v63 = (_866_v63).Update(_865_v62, (_dafny.Zero).Minus(_787_v2))
				_822_v27 = _822_v27
			}
			var _867_v64 *C0
			_ = _867_v64
			var _nw134 *C0 = New_C0_()
			_ = _nw134
			_nw134.Ctor__(_822_v27)
			_867_v64 = _nw134
		} else {
			var _868_v65 _dafny.Sequence
			_ = _868_v65
			_868_v65 = _dafny.SeqOf(_785_v0)
			var _869_v66 _dafny.MultiSet
			_ = _869_v66
			_869_v66 = _dafny.MultiSetOf(_785_v0, true, _785_v0, _785_v0, (_868_v65).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_868_v65).Cardinality()))).Uint32()).(bool))
			var _870_v67 _dafny.Map
			_ = _870_v67
			_870_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0)).Cardinality(), Companion_Default___.Fm35(_785_v0, false, globalState))
			var _871_v69 _dafny.Sequence
			_ = _871_v69
			_871_v69 = _dafny.SeqOf(p0, p0, p0)
			_785_v0 = ((func() _dafny.Int {
				if (_869_v66).Contains(!(_785_v0)) {
					return (_869_v66).Multiplicity(!(_785_v0))
				}
				return p0
			})()).Cmp(((_870_v67).Merge(func() _dafny.Map {
				var _coll42 = _dafny.NewMapBuilder()
				_ = _coll42
				for _iter43 := _dafny.Iterate((_871_v69).Elements()); ; {
					_compr_42, _ok43 := _iter43()
					if !_ok43 {
						break
					}
					var _872_v68 _dafny.Int
					_872_v68 = interface{}(_compr_42).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_871_v69, _872_v68) {
						_coll42.Add((_872_v68).Minus((_dafny.Zero).Minus(p0)), _871_v69)
					}
				}
				return _coll42.ToMap()
			}())).Cardinality()) <= 0
			var _873_v70 _dafny.Sequence
			_ = _873_v70
			_873_v70 = _dafny.UnicodeSeqOfUtf8Bytes("is")
			var _874_v71 _dafny.CodePoint
			_ = _874_v71
			_874_v71 = _dafny.CodePoint('v')
			var _875_v72 _dafny.Sequence
			_ = _875_v72
			_875_v72 = _dafny.SeqOf((_873_v70).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_873_v70).Cardinality()))).Uint32()).(_dafny.CodePoint), (_873_v70).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_873_v70).Cardinality()))).Uint32()).(_dafny.CodePoint), _874_v71, _874_v71)
			var _876_v73 _dafny.Set
			_ = _876_v73
			_876_v73 = _dafny.SetOf(_785_v0)
			var _877_v74 D11
			_ = _877_v74
			_877_v74 = Companion_D11_.Create_DC25_((_873_v70).Select((Companion_Default___.SafeIndex((_876_v73).Cardinality(), _dafny.IntOfUint32((_873_v70).Cardinality()))).Uint32()).(_dafny.CodePoint), p0, _785_v0)
			var _878_v75 _dafny.Map
			_ = _878_v75
			_878_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v0, p0)
			var _879_v76 D12
			_ = _879_v76
			_879_v76 = Companion_D12_.Create_DC29_(_785_v0, p0, false, (_dafny.Zero).Minus((_dafny.Zero).Minus((_878_v75).Cardinality())), _785_v0)
			var _source10 D12 = (func() D12 {
				if (_785_v0) == ((_877_v74).Dtor_cf43()) {
					return _879_v76
				}
				return _879_v76
			})()
			_ = _source10
			if _source10.Is_DC28() {
				var _880___mcc_h0 bool = _source10.Get_().(D12_DC28).Cf46
				_ = _880___mcc_h0
				var _881___mcc_h1 _dafny.Int = _source10.Get_().(D12_DC28).Cf47
				_ = _881___mcc_h1
				var _882_cf47 _dafny.Int = _881___mcc_h1
				_ = _882_cf47
				var _883_cf46 bool = _880___mcc_h0
				_ = _883_cf46
				var _884_v77 _dafny.Map
				_ = _884_v77
				_884_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), _874_v71)
				var _885_v78 _dafny.Array
				_ = _885_v78
				var _nwElement0_28 _dafny.CodePoint = _dafny.CodePoint('c')
				_ = _nwElement0_28
				var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(22))
				_ = _nw135
				(_nw135).ArraySet1CodePoint(_nwElement0_28, 0)
				(_nw135).ArraySet1CodePoint(_874_v71, 1)
				(_nw135).ArraySet1CodePoint(_874_v71, 2)
				(_nw135).ArraySet1CodePoint(_874_v71, 3)
				(_nw135).ArraySet1CodePoint(_874_v71, 4)
				(_nw135).ArraySet1CodePoint(_874_v71, 5)
				(_nw135).ArraySet1CodePoint(_874_v71, 6)
				(_nw135).ArraySet1CodePoint(_874_v71, 7)
				(_nw135).ArraySet1CodePoint(_dafny.CodePoint('i'), 8)
				(_nw135).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_884_v77).Contains(_dafny.IntOfInt64(509)) {
						return (_884_v77).Get(_dafny.IntOfInt64(509)).(_dafny.CodePoint)
					}
					return _874_v71
				})(), 9)
				(_nw135).ArraySet1CodePoint(_874_v71, 10)
				(_nw135).ArraySet1CodePoint(_874_v71, 11)
				(_nw135).ArraySet1CodePoint(_874_v71, 12)
				(_nw135).ArraySet1CodePoint(_874_v71, 13)
				(_nw135).ArraySet1CodePoint(_874_v71, 14)
				(_nw135).ArraySet1CodePoint(_874_v71, 15)
				(_nw135).ArraySet1CodePoint(_dafny.CodePoint('h'), 16)
				(_nw135).ArraySet1CodePoint(_dafny.CodePoint('b'), 17)
				(_nw135).ArraySet1CodePoint(Companion_Default___.Fm29(_883_cf46, globalState), 18)
				(_nw135).ArraySet1CodePoint(_874_v71, 19)
				(_nw135).ArraySet1CodePoint(_874_v71, 20)
				(_nw135).ArraySet1CodePoint(_874_v71, 21)
				_885_v78 = _nw135
				var _886_v79 _dafny.Map
				_ = _886_v79
				_886_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_885_v78, _876_v73)
				_886_v79 = _886_v79
				var _887_v80 _dafny.Array
				_ = _887_v80
				var _nw136 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
				_ = _nw136
				_887_v80 = _nw136
				var _888_v81 _dafny.MultiSet
				_ = _888_v81
				_888_v81 = _dafny.MultiSetOf(_dafny.IntOfInt64(-804))
				var _889_v82 *C0
				_ = _889_v82
				var _nw137 *C0 = New_C0_()
				_ = _nw137
				_nw137.Ctor__(_875_v72)
				_889_v82 = _nw137
				var _890_v83 _dafny.Set
				_ = _890_v83
				_890_v83 = _dafny.SetOf(_889_v82, _889_v82, _889_v82, _889_v82)
				var _891_v84 D4
				_ = _891_v84
				_891_v84 = Companion_D4_.Create_DC9_(_888_v81, _874_v71, _882_cf47, _890_v83, _dafny.IntOfInt64(712))
				var _892_v85 _dafny.Map
				_ = _892_v85
				_892_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D4_.Create_DC9_(_888_v81, _874_v71, _dafny.IntOfInt64(207), _890_v83, p0))).Update(_785_v0, _891_v84), _dafny.MultiSetFromSeq(_871_v69))
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_887_v80), 0))
				_ = _index137
				(_887_v80).ArraySet1((_892_v85).Merge(_892_v85), (_index137).Int())
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_887_v80), 0))
				_ = _index138
				(_887_v80).ArraySet1(_892_v85, (_index138).Int())
				var _893_v86 _dafny.Array
				_ = _893_v86
				var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(25))
				_ = _nw138
				_893_v86 = _nw138
				var _894_v87 D0
				_ = _894_v87
				_894_v87 = Companion_D0_.Create_DC0_(_893_v86, _882_cf47)
				var _895_v88 _dafny.Map
				_ = _895_v88
				_895_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm32(p0, globalState), _785_v0)
				var _896_v89 _dafny.Map
				_ = _896_v89
				_896_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_874_v71, (_895_v88).Cardinality())
				var _897_v91 _dafny.Array
				_ = _897_v91
				var _nwElement0_29 _dafny.Map = _896_v89
				_ = _nwElement0_29
				var _nw139 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(17))
				_ = _nw139
				(_nw139).ArraySet1(_nwElement0_29, 0)
				(_nw139).ArraySet1((_896_v89).Update(_874_v71, _dafny.IntOfUint32((_875_v72).Cardinality())), 1)
				(_nw139).ArraySet1(Companion_Default___.Fm36(_874_v71, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_878_v75, _874_v71)).Cardinality()), Companion_Default___.Fm32((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v0, _882_cf47)).Cardinality(), globalState), (Companion_D12_.Create_DC27_(_876_v73)).Dtor_cf45(), globalState), 2)
				(_nw139).ArraySet1(_896_v89, 3)
				(_nw139).ArraySet1(_896_v89, 4)
				(_nw139).ArraySet1(_896_v89, 5)
				(_nw139).ArraySet1(_896_v89, 6)
				(_nw139).ArraySet1(_896_v89, 7)
				(_nw139).ArraySet1(_896_v89, 8)
				(_nw139).ArraySet1(func() _dafny.Map {
					var _coll43 = _dafny.NewMapBuilder()
					_ = _coll43
					for _iter44 := _dafny.Iterate((_873_v70).Elements()); ; {
						_compr_43, _ok44 := _iter44()
						if !_ok44 {
							break
						}
						var _898_v90 _dafny.CodePoint
						_898_v90 = interface{}(_compr_43).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_873_v70, _898_v90) {
							_coll43.Add(_898_v90, _882_cf47)
						}
					}
					return _coll43.ToMap()
				}(), 9)
				(_nw139).ArraySet1(_896_v89, 10)
				(_nw139).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_874_v71, _dafny.IntOfInt64(774)), 11)
				(_nw139).ArraySet1(_896_v89, 12)
				(_nw139).ArraySet1(_896_v89, 13)
				(_nw139).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_874_v71, p0), 14)
				(_nw139).ArraySet1(_896_v89, 15)
				(_nw139).ArraySet1((_896_v89).Update(_874_v71, p0), 16)
				_897_v91 = _nw139
				var _899_v92 D0
				_ = _899_v92
				_899_v92 = Companion_D0_.Create_DC1_((_895_v88).Cardinality())
				var _900_v93 *C1
				_ = _900_v93
				var _nw140 *C1 = New_C1_()
				_ = _nw140
				_nw140.Ctor__((_894_v87).Dtor_cf1(), _897_v91, _dafny.SeqOf(p0, (_dafny.Zero).Minus(p0)), _899_v92)
				_900_v93 = _nw140
				var _901_v94 _dafny.Array
				_ = _901_v94
				var _nwElement0_30 bool = _785_v0
				_ = _nwElement0_30
				var _nw141 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(22))
				_ = _nw141
				(_nw141).ArraySet1(_nwElement0_30, 0)
				(_nw141).ArraySet1(_785_v0, 1)
				(_nw141).ArraySet1(_883_cf46, 2)
				(_nw141).ArraySet1(_785_v0, 3)
				(_nw141).ArraySet1(true, 4)
				(_nw141).ArraySet1(_785_v0, 5)
				(_nw141).ArraySet1(_785_v0, 6)
				(_nw141).ArraySet1(false, 7)
				(_nw141).ArraySet1(_785_v0, 8)
				(_nw141).ArraySet1(_883_cf46, 9)
				(_nw141).ArraySet1(_883_cf46, 10)
				(_nw141).ArraySet1(!(_883_cf46), 11)
				(_nw141).ArraySet1(_883_cf46, 12)
				(_nw141).ArraySet1(_883_cf46, 13)
				(_nw141).ArraySet1(_785_v0, 14)
				(_nw141).ArraySet1(!(_883_cf46), 15)
				(_nw141).ArraySet1(_785_v0, 16)
				(_nw141).ArraySet1(_883_cf46, 17)
				(_nw141).ArraySet1(_785_v0, 18)
				(_nw141).ArraySet1(_785_v0, 19)
				(_nw141).ArraySet1(false, 20)
				(_nw141).ArraySet1(_785_v0, 21)
				_901_v94 = _nw141
				var _902_v95 _dafny.Map
				_ = _902_v95
				var _903_v96 bool
				_ = _903_v96
				var _904_v97 bool
				_ = _904_v97
				var _905_v98 _dafny.Set
				_ = _905_v98
				var _out25 _dafny.Map
				_ = _out25
				var _out26 bool
				_ = _out26
				var _out27 bool
				_ = _out27
				var _out28 _dafny.Set
				_ = _out28
				_out25, _out26, _out27, _out28 = (_this).M11(_874_v71, _901_v94, _869_v66, Companion_Default___.SafeDivisionInt(_900_v93.F9, _900_v93.F9), globalState)
				_902_v95 = _out25
				_903_v96 = _out26
				_904_v97 = _out27
				_905_v98 = _out28
			} else if _source10.Is_DC29() {
				var _906___mcc_h2 bool = _source10.Get_().(D12_DC29).Cf48
				_ = _906___mcc_h2
				var _907___mcc_h3 _dafny.Int = _source10.Get_().(D12_DC29).Cf49
				_ = _907___mcc_h3
				var _908___mcc_h4 bool = _source10.Get_().(D12_DC29).Cf50
				_ = _908___mcc_h4
				var _909___mcc_h5 _dafny.Int = _source10.Get_().(D12_DC29).Cf51
				_ = _909___mcc_h5
				var _910___mcc_h6 bool = _source10.Get_().(D12_DC29).Cf52
				_ = _910___mcc_h6
				var _911_cf52 bool = _910___mcc_h6
				_ = _911_cf52
				var _912_cf51 _dafny.Int = _909___mcc_h5
				_ = _912_cf51
				var _913_cf50 bool = _908___mcc_h4
				_ = _913_cf50
				var _914_cf49 _dafny.Int = _907___mcc_h3
				_ = _914_cf49
				var _915_cf48 bool = _906___mcc_h2
				_ = _915_cf48
				var _916_v99 _dafny.Array
				_ = _916_v99
				var _nw142 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(26))
				_ = _nw142
				_916_v99 = _nw142
				var _917_v100 D0
				_ = _917_v100
				_917_v100 = Companion_D0_.Create_DC1_(_914_cf49)
				var _918_v101 *C1
				_ = _918_v101
				var _nw143 *C1 = New_C1_()
				_ = _nw143
				_nw143.Ctor__(_914_cf49, _916_v99, _871_v69, _917_v100)
				_918_v101 = _nw143
				_912_cf51 = _dafny.IntOfInt64(852)
				_914_cf49 = (func() _dafny.Int {
					if _913_cf50 {
						return _dafny.IntOfInt64(501)
					}
					return Companion_Default___.SafeDivisionInt(_912_cf51, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(436))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg68 _dafny.Int) interface{} {
							return coer68(arg68)
						}
					}((func(_919_v71 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_920_i8 _dafny.Int) _dafny.CodePoint {
							return _919_v71
						}
					})(_874_v71)))).Cardinality()))
				})()
				var _921_v102 *C0
				_ = _921_v102
				var _nw144 *C0 = New_C0_()
				_ = _nw144
				_nw144.Ctor__(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_873_v70, _875_v72), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_873_v70, _875_v72)).Cardinality()))).Uint32(), _dafny.CodePoint('p')))
				_921_v102 = _nw144
			} else {
				var _922___mcc_h7 _dafny.Set = _source10.Get_().(D12_DC27).Cf45
				_ = _922___mcc_h7
				var _923_cf45 _dafny.Set = _922___mcc_h7
				_ = _923_cf45
				r1 = (p0).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("oggyllotw"), (Companion_Default___.SafeIndex((_878_v75).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oggyllotw")).Cardinality()))).Uint32(), _874_v71)).Cardinality())) < 0
				var _924_v103 _dafny.Int
				_ = _924_v103
				_924_v103 = _dafny.IntOfInt64(345)
				_924_v103 = p0
				_874_v71 = _874_v71
				_924_v103 = Companion_Default___.SafeModuloInt((_dafny.SetOf(_785_v0, _785_v0, _785_v0, _785_v0)).Cardinality(), _dafny.IntOfInt64(469))
			}
			if false {
				var _925_v104 _dafny.Array
				_ = _925_v104
				var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
				_ = _nw145
				_925_v104 = _nw145
				var _926_v105 _dafny.Map
				_ = _926_v105
				_926_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _785_v0)
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_925_v104), 0))
				_ = _index139
				(_925_v104).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uxuolybl")).Cardinality()), (_926_v105).Cardinality()), (_index139).Int())
				var _927_v106 _dafny.MultiSet
				_ = _927_v106
				_927_v106 = _dafny.MultiSetOf(_871_v69)
				var _928_v107 D7
				_ = _928_v107
				_928_v107 = Companion_D7_.Create_DC19_((_927_v106).Cardinality())
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_925_v104), 0))
				_ = _index140
				(_925_v104).ArraySet1(Companion_Default___.SafeDivisionInt((p0).Plus((func() _dafny.Set {
					var _coll44 = _dafny.NewBuilder()
					_ = _coll44
					for _iter45 := _dafny.Iterate((_dafny.SeqOf(Companion_D7_.Create_DC19_((_dafny.Zero).Minus(p0)), Companion_Default___.Fm37(true, globalState), _928_v107, _928_v107, _928_v107)).Elements()); ; {
						_compr_44, _ok45 := _iter45()
						if !_ok45 {
							break
						}
						var _929_v108 D7
						_929_v108 = interface{}(_compr_44).(D7)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D7_.Create_DC19_((_dafny.Zero).Minus(p0)), Companion_Default___.Fm37(true, globalState), _928_v107, _928_v107, _928_v107), _929_v108) {
							_coll44.Add(_929_v108)
						}
					}
					return _coll44.ToSet()
				}()).Cardinality()), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v0, p0)).Merge(_878_v75)).Cardinality()), (_index140).Int())
				var _930_v109 _dafny.MultiSet
				_ = _930_v109
				_930_v109 = _dafny.MultiSetOf(p0)
				var _931_v110 *C0
				_ = _931_v110
				var _nw146 *C0 = New_C0_()
				_ = _nw146
				_nw146.Ctor__(_875_v72)
				_931_v110 = _nw146
				var _932_v111 _dafny.Set
				_ = _932_v111
				_932_v111 = _dafny.SetOf(_931_v110)
				var _933_v112 D4
				_ = _933_v112
				_933_v112 = Companion_D4_.Create_DC9_(_dafny.MultiSetOf(p0), (_875_v72).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_875_v72).Cardinality()))).Uint32()).(_dafny.CodePoint), _dafny.IntOfInt64(748), _932_v111, _dafny.IntOfInt64(83))
				var _934_v113 _dafny.Array
				_ = _934_v113
				var _nwElement0_31 bool = (func() bool {
					if (_926_v105).Contains(_dafny.IntOfInt64(-785)) {
						return (_926_v105).Get(_dafny.IntOfInt64(-785)).(bool)
					}
					return _785_v0
				})()
				_ = _nwElement0_31
				var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(17))
				_ = _nw147
				(_nw147).ArraySet1(_nwElement0_31, 0)
				(_nw147).ArraySet1(_785_v0, 1)
				(_nw147).ArraySet1(_785_v0, 2)
				(_nw147).ArraySet1(!_dafny.Companion_Sequence_.Contains(_873_v70, _874_v71), 3)
				(_nw147).ArraySet1((_930_v109).IsDisjointFrom((_933_v112).Dtor_cf14()), 4)
				(_nw147).ArraySet1(_785_v0, 5)
				(_nw147).ArraySet1(_785_v0, 6)
				(_nw147).ArraySet1((_785_v0) == (!(false)), 7)
				(_nw147).ArraySet1(_785_v0, 8)
				(_nw147).ArraySet1((_785_v0) && (_785_v0), 9)
				(_nw147).ArraySet1(_785_v0, 10)
				(_nw147).ArraySet1(_785_v0, 11)
				(_nw147).ArraySet1(_785_v0, 12)
				(_nw147).ArraySet1(false, 13)
				(_nw147).ArraySet1(_785_v0, 14)
				(_nw147).ArraySet1(_785_v0, 15)
				(_nw147).ArraySet1(_785_v0, 16)
				_934_v113 = _nw147
				var _nw148 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw148
				_934_v113 = _nw148
				var _935_v114 _dafny.Array
				_ = _935_v114
				var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
				_ = _nw149
				_935_v114 = _nw149
				var _936_v115 _dafny.Sequence
				_ = _936_v115
				_936_v115 = _dafny.SeqOf(_875_v72, _875_v72, _875_v72)
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_935_v114), 0))
				_ = _index141
				(_935_v114).ArraySet1(_936_v115, (_index141).Int())
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_935_v114), 0))
				_ = _index142
				(_935_v114).ArraySet1(_936_v115, (_index142).Int())
				(_931_v110).F8 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kmpvievok"), _875_v72), _931_v110.F8)
				var _937_v116 _dafny.Map
				_ = _937_v116
				_937_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_931_v110).Fm15(_871_v69, (_868_v65).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.IntOfUint32((_868_v65).Cardinality()))).Uint32()).(bool), globalState)).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_878_v75).Contains(!(_785_v0)) {
						return (_878_v75).Get(!(_785_v0)).(_dafny.Int)
					}
					return p0
				})(), _875_v72))
				var _938_v117 _dafny.Map
				_ = _938_v117
				_938_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_925_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_925_v104), 0))).Int()).(_dafny.Int), _931_v110.F8)
				_937_v116 = (_937_v116).Update(p0, _938_v117)
			} else {
				_873_v70 = _873_v70
				_785_v0 = _785_v0
				_869_v66 = _869_v66
				r1 = _dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm35(_785_v0, _785_v0, globalState), _871_v69)
				r0 = _785_v0
			}
			var _939_v118 _dafny.Int
			_ = _939_v118
			_939_v118 = _dafny.IntOfInt64(128)
			_939_v118 = p0
			r1 = _785_v0
		}
		var _940_v119 _dafny.Int
		_ = _940_v119
		_940_v119 = _dafny.IntOfInt64(-72)
		_940_v119 = Companion_Default___.SafeModuloInt(_940_v119, p0)
		var _941_v120 _dafny.Sequence
		_ = _941_v120
		_941_v120 = _dafny.SeqOf(p0, _940_v119, _940_v119, p0, p0)
		var _942_v121 _dafny.MultiSet
		_ = _942_v121
		_942_v121 = _dafny.MultiSetOf(p0, _dafny.IntOfInt64(348))
		var _943_v122 *C0
		_ = _943_v122
		var _nw150 *C0 = New_C0_()
		_ = _nw150
		_nw150.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("jctntdft"))
		_943_v122 = _nw150
		var _944_v123 _dafny.Set
		_ = _944_v123
		_944_v123 = _dafny.SetOf(_943_v122)
		var _hi5 _dafny.Int = _940_v119
		_ = _hi5
		for _945_i9 := (_941_v120).Select((Companion_Default___.SafeIndex((Companion_D4_.Create_DC9_(_942_v121, _dafny.CodePoint('n'), p0, _944_v123, p0)).Dtor_cf16(), _dafny.IntOfUint32((_941_v120).Cardinality()))).Uint32()).(_dafny.Int); _945_i9.Cmp(_hi5) < 0; _945_i9 = _945_i9.Plus(_dafny.One) {
			var _946_v124 _dafny.Array
			_ = _946_v124
			var _nw151 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw151
			_946_v124 = _nw151
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_946_v124), 0))
			_ = _index143
			(_946_v124).ArraySet1(_940_v119, (_index143).Int())
			var _947_v125 _dafny.Array
			_ = _947_v125
			var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
			_ = _nw152
			_947_v125 = _nw152
			var _948_v126 _dafny.Map
			_ = _948_v126
			_948_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_945_i9, p0)
			var _949_v127 D13
			_ = _949_v127
			_949_v127 = Companion_D13_.Create_DC30_(_947_v125)
			var _pat_let_tv27 = _947_v125
			_ = _pat_let_tv27
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_946_v124), 0))
			_ = _index144
			var _rhs85 _dafny.Int = (func() _dafny.Int {
				if (_948_v126).Contains((_dafny.Zero).Minus(_945_i9)) {
					return (_948_v126).Get((_dafny.Zero).Minus(_945_i9)).(_dafny.Int)
				}
				return (func() _dafny.Int {
					if (_942_v121).Contains(_945_i9) {
						return (_942_v121).Multiplicity(_945_i9)
					}
					return (_dafny.Zero).Minus(_940_v119)
				})()
			})()
			_ = _rhs85
			var _rhs86 _dafny.Int = p0
			_ = _rhs86
			var _rhs87 _dafny.Array = (func(_pat_let20_0 D13) D13 {
				return func(_950_dt__update__tmp_h2 D13) D13 {
					return func(_pat_let21_0 _dafny.Array) D13 {
						return func(_951_dt__update_hcf53_h0 _dafny.Array) D13 {
							return Companion_D13_.Create_DC30_(_951_dt__update_hcf53_h0)
						}(_pat_let21_0)
					}(_pat_let_tv27)
				}(_pat_let20_0)
			}(_949_v127)).Dtor_cf53()
			_ = _rhs87
			var _lhs60 _dafny.Array = _946_v124
			_ = _lhs60
			var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_946_v124), 0))
			_ = _lhs61
			_940_v119 = _rhs85
			(_lhs60).ArraySet1(_rhs86, (_lhs61).Int())
			_947_v125 = _rhs87
			r1 = !((func() bool {
				if !(_785_v0) || (_785_v0) {
					return _785_v0
				}
				return _785_v0
			})())
			(_943_v122).F8 = _943_v122.F8
			var _952_v128 _dafny.Sequence
			_ = _952_v128
			_952_v128 = _dafny.SeqOf(_941_v120, _941_v120)
			if (func() bool {
				if (_dafny.IntOfUint32(((_952_v128).Select((Companion_Default___.SafeIndex(_940_v119, _dafny.IntOfUint32((_952_v128).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())).Cmp(_dafny.IntOfUint32((_941_v120).Cardinality())) <= 0 {
					return (_dafny.IntOfInt64(-986)).Cmp(_dafny.IntOfInt64(792)) != 0
				}
				return _785_v0
			})() {
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_946_v124), 0))
				_ = _index145
				(_946_v124).ArraySet1((Companion_Default___.SafeDivisionInt(_940_v119, (_946_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_946_v124), 0))).Int()).(_dafny.Int))).Times(_dafny.IntOfUint32((_dafny.SeqOf(_785_v0, _785_v0)).Cardinality())), (_index145).Int())
				var _953_v129 _dafny.Map
				_ = _953_v129
				_953_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v0, _946_v124)
				_953_v129 = (_953_v129).Update(_785_v0, _946_v124)
				(_943_v122).F8 = _dafny.Companion_Sequence_.Concatenate((_943_v122).Fm15(_941_v120, !(_785_v0), globalState), _943_v122.F8)
				_949_v127 = _949_v127
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_946_v124), 0))
				_ = _index146
				var _rhs88 _dafny.MultiSet = _942_v121
				_ = _rhs88
				var _rhs89 _dafny.Sequence = _943_v122.F8
				_ = _rhs89
				var _rhs90 _dafny.Int = _945_i9
				_ = _rhs90
				var _rhs91 _dafny.Sequence = (func() _dafny.Sequence {
					if _785_v0 {
						return _dafny.Companion_Sequence_.Concatenate(_941_v120, _941_v120)
					}
					return _941_v120
				})()
				_ = _rhs91
				var _lhs62 *C0 = _943_v122
				_ = _lhs62
				var _lhs63 _dafny.Array = _946_v124
				_ = _lhs63
				var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_946_v124), 0))
				_ = _lhs64
				_942_v121 = _rhs88
				_lhs62.F8 = _rhs89
				(_lhs63).ArraySet1(_rhs90, (_lhs64).Int())
				_941_v120 = _rhs91
			} else {
				(_943_v122).F8 = _943_v122.F8
				var _954_v130 _dafny.Sequence
				_ = _954_v130
				_954_v130 = _dafny.SeqOf(_943_v122.F8)
				var _955_v131 _dafny.Sequence
				_ = _955_v131
				_955_v131 = _dafny.SeqOf(!(_785_v0))
				var _rhs92 bool = _785_v0
				_ = _rhs92
				var _rhs93 bool = _785_v0
				_ = _rhs93
				var _rhs94 _dafny.Sequence = (_954_v130).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_955_v131, (Companion_Default___.SafeIndex(_945_i9, _dafny.IntOfUint32((_955_v131).Cardinality()))).Uint32(), _785_v0)).Cardinality()), _dafny.IntOfUint32((_954_v130).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs94
				var _lhs65 *C0 = _943_v122
				_ = _lhs65
				_785_v0 = _rhs92
				_785_v0 = _rhs93
				_lhs65.F8 = _rhs94
				var _956_v132 _dafny.Array
				_ = _956_v132
				var _nw153 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw153
				_956_v132 = _nw153
				var _957_v133 _dafny.Array
				_ = _957_v133
				var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
				_ = _nw154
				_957_v133 = _nw154
				var _958_v134 D2
				_ = _958_v134
				_958_v134 = Companion_D2_.Create_DC4_(_dafny.IntOfInt64(289), _785_v0, _957_v133)
				var _959_v135 _dafny.CodePoint
				_ = _959_v135
				_959_v135 = _dafny.CodePoint('c')
				var _960_v136 _dafny.MultiSet
				_ = _960_v136
				_960_v136 = _dafny.MultiSetOf(_785_v0)
				var _961_v137 D0
				_ = _961_v137
				_961_v137 = Companion_D0_.Create_DC1_((_960_v136).Cardinality())
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_956_v132), 0))
				_ = _index147
				(_956_v132).ArraySet1((_this).Fm5((_958_v134).Dtor_cf5(), _785_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_959_v135, _961_v137), _785_v0, globalState), (_index147).Int())
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_956_v132), 0))
				_ = _index148
				(_956_v132).ArraySet1((func() bool {
					if _785_v0 {
						return ((_960_v136).Update(_785_v0, Companion_Default___.Abs((_946_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_946_v124), 0))).Int()).(_dafny.Int)))).IsProperSubsetOf(_960_v136)
					}
					return (func() bool {
						if _785_v0 {
							return !(!(_785_v0))
						}
						return _785_v0
					})()
				})(), (_index148).Int())
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_946_v124), 0))
				_ = _index149
				(_946_v124).ArraySet1((_946_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_946_v124), 0))).Int()).(_dafny.Int), (_index149).Int())
				var _962_v138 *C0
				_ = _962_v138
				var _nw155 *C0 = New_C0_()
				_ = _nw155
				_nw155.Ctor__(_943_v122.F8)
				_962_v138 = _nw155
			}
		}
		var _963_v139 *C0
		_ = _963_v139
		var _nw156 *C0 = New_C0_()
		_ = _nw156
		_nw156.Ctor__(_943_v122.F8)
		_963_v139 = _nw156
		var _964_v140 _dafny.Map
		_ = _964_v140
		_964_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v0, _940_v119)
		var _hi6 _dafny.Int = (func() _dafny.Int {
			if _785_v0 {
				return _940_v119
			}
			return (_964_v140).Cardinality()
		})()
		_ = _hi6
		for _965_i10 := (_964_v140).Cardinality(); _965_i10.Cmp(_hi6) < 0; _965_i10 = _965_i10.Plus(_dafny.One) {
			r1 = _785_v0
			r1 = true
			(_963_v139).F8 = _963_v139.F8
			var _966_v141 _dafny.Array
			_ = _966_v141
			var _nw157 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw157
			_966_v141 = _nw157
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_966_v141), 0))
			_ = _index150
			(_966_v141).ArraySet1(!(_785_v0), (_index150).Int())
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_966_v141), 0))
			_ = _index151
			(_966_v141).ArraySet1(_785_v0, (_index151).Int())
		}
		var _967_v142 _dafny.Array
		_ = _967_v142
		var _len0_20 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_20
		var _nw158 _dafny.Array
		_ = _nw158
		if _len0_20.Cmp(_dafny.Zero) == 0 {
			_nw158 = _dafny.NewArray(_len0_20)
		} else {
			var _init20 func(_dafny.Int) _dafny.Int = (func(_968_v139 *C0, _969_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_970_i11 _dafny.Int) _dafny.Int {
					return (_970_i11).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_968_v139.F8, _969_p0)).Cardinality())
				}
			})(_963_v139, p0)
			_ = _init20
			var _element0_20 = _init20(_dafny.Zero)
			_ = _element0_20
			_nw158 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
			(_nw158).ArraySet1(_element0_20, 0)
			var _nativeLen0_20 = (_len0_20).Int()
			_ = _nativeLen0_20
			for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
				(_nw158).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
			}
		}
		_967_v142 = _nw158
		var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_967_v142), 0))
		_ = _index152
		(_967_v142).ArraySet1(_940_v119, (_index152).Int())
		var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_967_v142), 0))
		_ = _index153
		(_967_v142).ArraySet1(_940_v119, (_index153).Int())
		r0 = _785_v0
		r1 = _785_v0
		return r0, r1
	}
}
func (_this *C3) M2(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _971_v0 bool
		_ = _971_v0
		_971_v0 = true
		var _972_v1 _dafny.MultiSet
		_ = _972_v1
		_972_v1 = _dafny.MultiSetOf(_971_v0, _971_v0)
		var _973_v2 _dafny.Map
		_ = _973_v2
		_973_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (Companion_D9_.Create_DC21_(_972_v1)).Dtor_cf37())
		_973_v2 = (_973_v2).Update(p0, _972_v1)
		var _974_v3 D7
		_ = _974_v3
		_974_v3 = Companion_D7_.Create_DC18_()
		var _975_v4 _dafny.Array
		_ = _975_v4
		var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(4))
		_ = _nw159
		_975_v4 = _nw159
		var _976_v5 _dafny.Array
		_ = _976_v5
		var _nwElement0_32 bool = _971_v0
		_ = _nwElement0_32
		var _nw160 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(19))
		_ = _nw160
		(_nw160).ArraySet1(_nwElement0_32, 0)
		(_nw160).ArraySet1(_971_v0, 1)
		(_nw160).ArraySet1(_971_v0, 2)
		(_nw160).ArraySet1(_971_v0, 3)
		(_nw160).ArraySet1(_971_v0, 4)
		(_nw160).ArraySet1(_971_v0, 5)
		(_nw160).ArraySet1(_971_v0, 6)
		(_nw160).ArraySet1(false, 7)
		(_nw160).ArraySet1(_971_v0, 8)
		(_nw160).ArraySet1(_971_v0, 9)
		(_nw160).ArraySet1(_971_v0, 10)
		(_nw160).ArraySet1(_971_v0, 11)
		(_nw160).ArraySet1(Companion_Default___.Fm30(p0, _971_v0, _971_v0, globalState), 12)
		(_nw160).ArraySet1(true, 13)
		(_nw160).ArraySet1(_971_v0, 14)
		(_nw160).ArraySet1(_971_v0, 15)
		(_nw160).ArraySet1(_971_v0, 16)
		(_nw160).ArraySet1(_971_v0, 17)
		(_nw160).ArraySet1(_971_v0, 18)
		_976_v5 = _nw160
		var _977_v6 _dafny.Array
		_ = _977_v6
		var _nwElement0_33 _dafny.Array = _976_v5
		_ = _nwElement0_33
		var _nw161 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(10))
		_ = _nw161
		(_nw161).ArraySet1(_nwElement0_33, 0)
		(_nw161).ArraySet1(_976_v5, 1)
		(_nw161).ArraySet1(_976_v5, 2)
		(_nw161).ArraySet1(_976_v5, 3)
		(_nw161).ArraySet1(_976_v5, 4)
		(_nw161).ArraySet1(_976_v5, 5)
		(_nw161).ArraySet1(_976_v5, 6)
		(_nw161).ArraySet1(_976_v5, 7)
		(_nw161).ArraySet1(_976_v5, 8)
		(_nw161).ArraySet1(_976_v5, 9)
		_977_v6 = _nw161
		var _rhs95 D7 = _974_v3
		_ = _rhs95
		var _rhs96 _dafny.Array = _975_v4
		_ = _rhs96
		var _rhs97 _dafny.Array = _977_v6
		_ = _rhs97
		_974_v3 = _rhs95
		_975_v4 = _rhs96
		_977_v6 = _rhs97
		var _hi7 _dafny.Int = p0
		_ = _hi7
		for _978_i0 := Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfInt64(-664)); _978_i0.Cmp(_hi7) < 0; _978_i0 = _978_i0.Plus(_dafny.One) {
			var _979_v7 _dafny.Int
			_ = _979_v7
			_979_v7 = _dafny.IntOfInt64(373)
			var _980_v8 _dafny.Sequence
			_ = _980_v8
			_980_v8 = _dafny.UnicodeSeqOfUtf8Bytes("fbkmmu")
			var _981_v9 _dafny.CodePoint
			_ = _981_v9
			_981_v9 = _dafny.CodePoint('a')
			_979_v7 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-424))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg69 _dafny.Int) interface{} {
					return coer69(arg69)
				}
			}(func(_982_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			})), _dafny.Companion_Sequence_.Update(_980_v8, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_980_v8).Cardinality()), _dafny.IntOfUint32((_980_v8).Cardinality()))).Uint32(), _981_v9))).Cardinality())
			_981_v9 = _981_v9
			var _983_v10 *C0
			_ = _983_v10
			var _nw162 *C0 = New_C0_()
			_ = _nw162
			_nw162.Ctor__(_dafny.Companion_Sequence_.Concatenate(_980_v8, _980_v8))
			_983_v10 = _nw162
			var _984_v11 _dafny.Array
			_ = _984_v11
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_21
			var _nw163 _dafny.Array
			_ = _nw163
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw163 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) _dafny.Int = (func(_985_v7 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_986_i2 _dafny.Int) _dafny.Int {
						return (_986_i2).Plus(_985_v7)
					}
				})(_979_v7)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw163 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw163).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw163).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_984_v11 = _nw163
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_984_v11), 0))
			_ = _index154
			(_984_v11).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(625))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg70 _dafny.Int) interface{} {
					return coer70(arg70)
				}
			}((func(_987_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_988_i3 _dafny.Int) _dafny.Int {
					return _987_p0
				}
			})(p0)))).Cardinality())).Plus(_978_i0), (_index154).Int())
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_984_v11), 0))
			_ = _index155
			(_984_v11).ArraySet1((func() _dafny.Int {
				if _971_v0 {
					return _979_v7
				}
				return Companion_Default___.SafeDivisionInt(_978_i0, _978_i0)
			})(), (_index155).Int())
		}
		var _989_v12 _dafny.Array
		_ = _989_v12
		var _nw164 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(14))
		_ = _nw164
		_989_v12 = _nw164
		for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_989_v12), 0))); ; {
			_guard_loop_1, _ok46 := _iter46()
			if !_ok46 {
				break
			}
			var _990_i4 _dafny.Int
			_990_i4 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_990_i4).Sign() != -1) && ((_990_i4).Cmp(_dafny.ArrayLen((_989_v12), 0)) < 0)) {
				(_989_v12).ArraySet1(Companion_D2_.Create_DC5_(p0, _dafny.SeqOf(p0, p0), _971_v0), (_990_i4).Int())
			}
		}
		var _991_v13 bool
		_ = _991_v13
		var _992_v14 _dafny.Int
		_ = _992_v14
		var _993_v15 _dafny.Int
		_ = _993_v15
		var _994_v16 bool
		_ = _994_v16
		var _out29 bool
		_ = _out29
		var _out30 _dafny.Int
		_ = _out30
		var _out31 _dafny.Int
		_ = _out31
		var _out32 bool
		_ = _out32
		_out29, _out30, _out31, _out32 = (_this).M12(_971_v0, globalState)
		_991_v13 = _out29
		_992_v14 = _out30
		_993_v15 = _out31
		_994_v16 = _out32
		var _995_v17 _dafny.Map
		_ = _995_v17
		_995_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_994_v16, false)
		var _hi8 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, _993_v15)
		_ = _hi8
		for _996_i5 := ((_995_v17).Merge(_995_v17)).Cardinality(); _996_i5.Cmp(_hi8) < 0; _996_i5 = _996_i5.Plus(_dafny.One) {
			var _997_v18 D9
			_ = _997_v18
			_997_v18 = Companion_D9_.Create_DC22_((_993_v15).Minus(_992_v14))
			var _source11 D9 = _997_v18
			_ = _source11
			if _source11.Is_DC22() {
				var _998___mcc_h0 _dafny.Int = _source11.Get_().(D9_DC22).Cf38
				_ = _998___mcc_h0
				var _999_cf38 _dafny.Int = _998___mcc_h0
				_ = _999_cf38
				var _1000_v19 _dafny.Sequence
				_ = _1000_v19
				_1000_v19 = _dafny.SeqOf(_994_v16)
				var _1001_v20 _dafny.MultiSet
				_ = _1001_v20
				_1001_v20 = _dafny.MultiSetOf(_1000_v19, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_971_v0), _1000_v19), _1000_v19)
				_1001_v20 = ((_1001_v20).Union(_1001_v20)).Union(_1001_v20)
				_999_cf38 = (p0).Plus((_992_v14).Plus(_993_v15))
				var _1002_v21 _dafny.Array
				_ = _1002_v21
				var _nw165 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
				_ = _nw165
				_1002_v21 = _nw165
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_1002_v21), 0))
				_ = _index156
				(_1002_v21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm38(_994_v16, _992_v14, _994_v16, globalState), _1000_v19), (_index156).Int())
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_1002_v21), 0))
				_ = _index157
				(_1002_v21).ArraySet1(_1000_v19, (_index157).Int())
				var _1003_v22 D2
				_ = _1003_v22
				_1003_v22 = Companion_D2_.Create_DC5_(_999_cf38, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(24))).Uint32(), func(coer71 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}((func(_1004_v14 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1005_i6 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_1004_v14)
					}
				})(_992_v14))), _994_v16)
				var _1006_v23 D2
				_ = _1006_v23
				_1006_v23 = Companion_D2_.Create_DC6_(_1003_v22)
				var _1007_v24 _dafny.Map
				_ = _1007_v24
				_1007_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_v13, p0)
				var _1008_v25 _dafny.Sequence
				_ = _1008_v25
				_1008_v25 = _dafny.SeqOf(_1007_v24)
				var _1009_v26 _dafny.MultiSet
				_ = _1009_v26
				_1009_v26 = _dafny.MultiSetOf(_995_v17, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _991_v13), (_995_v17).Update(false, _971_v0), _995_v17, _995_v17)
				var _1010_v27 _dafny.Sequence
				_ = _1010_v27
				_1010_v27 = _dafny.SeqOf(_996_i5, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm38(_991_v13, ((_1008_v25).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1009_v26).Contains(_995_v17) {
						return (_1009_v26).Multiplicity(_995_v17)
					}
					return (_dafny.Zero).Minus(Companion_Default___.Fm32(_993_v15, globalState))
				})(), _dafny.IntOfUint32((_1008_v25).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _994_v16, globalState), (Companion_Default___.SafeIndex(_999_cf38, _dafny.IntOfUint32((Companion_Default___.Fm38(_991_v13, ((_1008_v25).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1009_v26).Contains(_995_v17) {
						return (_1009_v26).Multiplicity(_995_v17)
					}
					return (_dafny.Zero).Minus(Companion_Default___.Fm32(_993_v15, globalState))
				})(), _dafny.IntOfUint32((_1008_v25).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _994_v16, globalState)).Cardinality()))).Uint32(), _971_v0)).Cardinality()), _993_v15, Companion_Default___.Fm32(_dafny.IntOfUint32(((_1002_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_1002_v21), 0))).Int()).(_dafny.Sequence)).Cardinality()), globalState), _999_cf38)
				var _1011_v28 _dafny.Map
				_ = _1011_v28
				_1011_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1006_v23, !_dafny.Companion_Sequence_.Equal(_1010_v27, _1010_v27))
				_971_v0 = (func() bool {
					if (_1011_v28).Contains(_1006_v23) {
						return (_1011_v28).Get(_1006_v23).(bool)
					}
					return _994_v16
				})()
			} else {
				var _1012___mcc_h1 _dafny.MultiSet = _source11.Get_().(D9_DC21).Cf37
				_ = _1012___mcc_h1
				var _1013_cf37 _dafny.MultiSet = _1012___mcc_h1
				_ = _1013_cf37
				var _1014_v29 _dafny.Map
				_ = _1014_v29
				_1014_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_992_v14, _971_v0)
				var _1015_v30 _dafny.Sequence
				_ = _1015_v30
				_1015_v30 = _dafny.UnicodeSeqOfUtf8Bytes("cmpc")
				var _1016_v31 _dafny.Set
				_ = _1016_v31
				_1016_v31 = _dafny.SetOf(_994_v16)
				var _1017_v32 _dafny.Map
				_ = _1017_v32
				_1017_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1015_v30).Cardinality()), _1016_v31)
				var _1018_v33 D12
				_ = _1018_v33
				_1018_v33 = Companion_D12_.Create_DC27_((func() _dafny.Set {
					if (_1017_v32).Contains(_993_v15) {
						return (_1017_v32).Get(_993_v15).(_dafny.Set)
					}
					return _1016_v31
				})())
				var _1019_v34 _dafny.Map
				_ = _1019_v34
				_1019_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1014_v29).Update(_996_i5, _971_v0), _1018_v33)
				var _pat_let_tv28 = _1016_v31
				_ = _pat_let_tv28
				_1019_v34 = (_1019_v34).Update(_1014_v29, func(_pat_let22_0 D12) D12 {
					return func(_1020_dt__update__tmp_h0 D12) D12 {
						return func(_pat_let23_0 _dafny.Set) D12 {
							return func(_1021_dt__update_hcf45_h0 _dafny.Set) D12 {
								return Companion_D12_.Create_DC27_(_1021_dt__update_hcf45_h0)
							}(_pat_let23_0)
						}(_pat_let_tv28)
					}(_pat_let22_0)
				}(_1018_v33))
				var _1022_v35 _dafny.Set
				_ = _1022_v35
				_1022_v35 = _dafny.SetOf(_976_v5)
				var _1023_v36 _dafny.Set
				_ = _1023_v36
				_1023_v36 = _1022_v35
				var _1024_v37 _dafny.Array
				_ = _1024_v37
				var _nwElement0_34 _dafny.Set = _1023_v36
				_ = _nwElement0_34
				var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(22))
				_ = _nw166
				(_nw166).ArraySet1(_nwElement0_34, 0)
				(_nw166).ArraySet1(_1022_v35, 1)
				(_nw166).ArraySet1(_dafny.SetOf(_976_v5), 2)
				(_nw166).ArraySet1(_1023_v36, 3)
				(_nw166).ArraySet1(_1022_v35, 4)
				(_nw166).ArraySet1(_1023_v36, 5)
				(_nw166).ArraySet1(_1023_v36, 6)
				(_nw166).ArraySet1(_1023_v36, 7)
				(_nw166).ArraySet1(_1023_v36, 8)
				(_nw166).ArraySet1(_1022_v35, 9)
				(_nw166).ArraySet1(_1023_v36, 10)
				(_nw166).ArraySet1(_1023_v36, 11)
				(_nw166).ArraySet1(_1023_v36, 12)
				(_nw166).ArraySet1((func() _dafny.Set {
					if !(_971_v0) {
						return _1023_v36
					}
					return _1022_v35
				})(), 13)
				(_nw166).ArraySet1(_1023_v36, 14)
				(_nw166).ArraySet1(_1023_v36, 15)
				(_nw166).ArraySet1(_1023_v36, 16)
				(_nw166).ArraySet1(_1022_v35, 17)
				(_nw166).ArraySet1(_1023_v36, 18)
				(_nw166).ArraySet1(_1023_v36, 19)
				(_nw166).ArraySet1(_1023_v36, 20)
				(_nw166).ArraySet1(_1023_v36, 21)
				_1024_v37 = _nw166
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_1024_v37), 0))
				_ = _index158
				(_1024_v37).ArraySet1(_1023_v36, (_index158).Int())
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_1024_v37), 0))
				_ = _index159
				(_1024_v37).ArraySet1(_1023_v36, (_index159).Int())
				var _1025_v38 _dafny.CodePoint
				_ = _1025_v38
				_1025_v38 = _dafny.CodePoint('f')
				var _1026_v39 _dafny.Sequence
				_ = _1026_v39
				_1026_v39 = _dafny.SeqOf(_994_v16)
				var _1027_v40 _dafny.Map
				_ = _1027_v40
				_1027_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1025_v38, _993_v15)
				var _1028_v42 _dafny.Set
				_ = _1028_v42
				_1028_v42 = _dafny.SetOf(_1025_v38, _1025_v38, _1025_v38)
				var _1029_v43 _dafny.Sequence
				_ = _1029_v43
				_1029_v43 = _dafny.SeqOf(_1027_v40, _1027_v40)
				var _1030_v44 _dafny.Array
				_ = _1030_v44
				var _nwElement0_35 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1025_v38, _dafny.IntOfUint32((_1026_v39).Cardinality()))
				_ = _nwElement0_35
				var _nw167 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(25))
				_ = _nw167
				(_nw167).ArraySet1(_nwElement0_35, 0)
				(_nw167).ArraySet1(_1027_v40, 1)
				(_nw167).ArraySet1(_1027_v40, 2)
				(_nw167).ArraySet1(_1027_v40, 3)
				(_nw167).ArraySet1(func() _dafny.Map {
					var _coll45 = _dafny.NewMapBuilder()
					_ = _coll45
					for _iter47 := _dafny.Iterate((_1028_v42).Elements()); ; {
						_compr_45, _ok47 := _iter47()
						if !_ok47 {
							break
						}
						var _1031_v41 _dafny.CodePoint
						_1031_v41 = interface{}(_compr_45).(_dafny.CodePoint)
						if (_1028_v42).Contains(_1031_v41) {
							_coll45.Add(_1031_v41, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-584))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg72 _dafny.Int) interface{} {
									return coer72(arg72)
								}
							}((func(_1032_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1033_i7 _dafny.Int) _dafny.CodePoint {
									return _1032_v38
								}
							})(_1025_v38)))).Cardinality()))
						}
					}
					return _coll45.ToMap()
				}(), 4)
				(_nw167).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm29(_971_v0, globalState), _993_v15), 5)
				(_nw167).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1025_v38, Companion_Default___.Fm32(_993_v15, globalState)), 6)
				(_nw167).ArraySet1(_1027_v40, 7)
				(_nw167).ArraySet1(_1027_v40, 8)
				(_nw167).ArraySet1(_1027_v40, 9)
				(_nw167).ArraySet1((_1029_v43).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1029_v43).Cardinality()))).Uint32()).(_dafny.Map), 10)
				(_nw167).ArraySet1((_1029_v43).Select((Companion_Default___.SafeIndex(_992_v14, _dafny.IntOfUint32((_1029_v43).Cardinality()))).Uint32()).(_dafny.Map), 11)
				(_nw167).ArraySet1((_1027_v40), 12)
				(_nw167).ArraySet1(_1027_v40, 13)
				(_nw167).ArraySet1(_1027_v40, 14)
				(_nw167).ArraySet1(_1027_v40, 15)
				(_nw167).ArraySet1(_1027_v40, 16)
				(_nw167).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1025_v38, _996_i5), 17)
				(_nw167).ArraySet1(_1027_v40, 18)
				(_nw167).ArraySet1(_1027_v40, 19)
				(_nw167).ArraySet1(_1027_v40, 20)
				(_nw167).ArraySet1(_1027_v40, 21)
				(_nw167).ArraySet1(_1027_v40, 22)
				(_nw167).ArraySet1(_1027_v40, 23)
				(_nw167).ArraySet1(_1027_v40, 24)
				_1030_v44 = _nw167
				var _1034_v45 _dafny.Sequence
				_ = _1034_v45
				_1034_v45 = _dafny.SeqOf(_992_v14)
				var _1035_v46 D0
				_ = _1035_v46
				_1035_v46 = Companion_D0_.Create_DC1_(p0)
				var _1036_v47 *C1
				_ = _1036_v47
				var _nw168 *C1 = New_C1_()
				_ = _nw168
				_nw168.Ctor__((_996_i5).Plus(_996_i5), _1030_v44, _1034_v45, _1035_v46)
				_1036_v47 = _nw168
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_976_v5), 0))
				_ = _index160
				(_976_v5).ArraySet1(_991_v13, (_index160).Int())
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_976_v5), 0))
				_ = _index161
				(_976_v5).ArraySet1((Companion_Default___.Fm39(_991_v13, globalState)).Contains(_993_v15), (_index161).Int())
			}
			_993_v15 = _996_i5
			var _1037_v48 _dafny.CodePoint
			_ = _1037_v48
			_1037_v48 = _dafny.CodePoint('t')
			var _1038_v50 D7
			_ = _1038_v50
			_1038_v50 = Companion_D7_.Create_DC19_((func() _dafny.Set {
				var _coll46 = _dafny.NewBuilder()
				_ = _coll46
				for _iter48 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(195), _dafny.IntOfInt64(383))); ; {
					_compr_46, _ok48 := _iter48()
					if !_ok48 {
						break
					}
					var _1039_v49 _dafny.Int
					_1039_v49 = interface{}(_compr_46).(_dafny.Int)
					if ((_dafny.IntOfInt64(195)).Cmp(_1039_v49) <= 0) && ((_1039_v49).Cmp(_dafny.IntOfInt64(383)) < 0) {
						_coll46.Add(Companion_Default___.SafeModuloInt(_1039_v49, _996_i5))
					}
				}
				return _coll46.ToSet()
			}()).Cardinality())
			var _source12 D12 = Companion_Default___.Fm40(_991_v13, _1037_v48, _994_v16, _1038_v50, globalState)
			_ = _source12
			if _source12.Is_DC28() {
				var _1040___mcc_h2 bool = _source12.Get_().(D12_DC28).Cf46
				_ = _1040___mcc_h2
				var _1041___mcc_h3 _dafny.Int = _source12.Get_().(D12_DC28).Cf47
				_ = _1041___mcc_h3
				var _1042_cf47 _dafny.Int = _1041___mcc_h3
				_ = _1042_cf47
				var _1043_cf46 bool = _1040___mcc_h2
				_ = _1043_cf46
				_993_v15 = (_dafny.Zero).Minus(_996_i5)
				var _1044_v51 _dafny.Array
				_ = _1044_v51
				var _nw169 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
				_ = _nw169
				_1044_v51 = _nw169
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1044_v51), 0))
				_ = _index162
				(_1044_v51).ArraySet1(p0, (_index162).Int())
				var _1045_v52 _dafny.Map
				_ = _1045_v52
				_1045_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_996_i5, _1044_v51)
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_1044_v51), 0))
				_ = _index163
				(_1044_v51).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_1045_v52).Cardinality())).Cardinality()), (_index163).Int())
				var _1046_v53 _dafny.Sequence
				_ = _1046_v53
				_1046_v53 = _dafny.UnicodeSeqOfUtf8Bytes("ruavycy")
				var _1047_v54 D13
				_ = _1047_v54
				_1047_v54 = Companion_D13_.Create_DC31_(_1043_cf46, _dafny.IntOfUint32((_1046_v53).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-504))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}((func(_1048_v48 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1049_i8 _dafny.Int) _dafny.CodePoint {
						return _1048_v48
					}
				})(_1037_v48)))).Cardinality()))
				var _1050_v55 _dafny.Array
				_ = _1050_v55
				var _nwElement0_36 D13 = func(_pat_let24_0 D13) D13 {
					return func(_1051_dt__update__tmp_h4 D13) D13 {
						return func(_pat_let25_0 _dafny.Int) D13 {
							return func(_1052_dt__update_hcf56_h0 _dafny.Int) D13 {
								return Companion_D13_.Create_DC31_((_1051_dt__update__tmp_h4).Dtor_cf54(), (_1051_dt__update__tmp_h4).Dtor_cf55(), _1052_dt__update_hcf56_h0)
							}(_pat_let25_0)
						}(_996_i5)
					}(_pat_let24_0)
				}(_1047_v54)
				_ = _nwElement0_36
				var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.One)
				_ = _nw170
				(_nw170).ArraySet1(_nwElement0_36, 0)
				_1050_v55 = _nw170
				var _1053_v56 _dafny.Map
				_ = _1053_v56
				_1053_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm30(_996_i5, _991_v13, true, globalState), _1050_v55)
				var _1054_v57 _dafny.Array
				_ = _1054_v57
				var _nwElement0_37 _dafny.Array = _1050_v55
				_ = _nwElement0_37
				var _nw171 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(3))
				_ = _nw171
				(_nw171).ArraySet1(_nwElement0_37, 0)
				(_nw171).ArraySet1((func() _dafny.Array {
					if (_1053_v56).Contains(!(_1043_cf46)) {
						return (_1053_v56).Get(!(_1043_cf46)).(_dafny.Array)
					}
					return _1050_v55
				})(), 1)
				(_nw171).ArraySet1(_1050_v55, 2)
				_1054_v57 = _nw171
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1054_v57), 0))
				_ = _index164
				(_1054_v57).ArraySet1(_1050_v55, (_index164).Int())
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1054_v57), 0))
				_ = _index165
				(_1054_v57).ArraySet1(_1050_v55, (_index165).Int())
				var _1055_v58 _dafny.Sequence
				_ = _1055_v58
				_1055_v58 = _dafny.SeqOf(_1042_cf47)
				var _1056_v59 _dafny.Map
				_ = _1056_v59
				_1056_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_971_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(977))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg74 _dafny.Int) interface{} {
						return coer74(arg74)
					}
				}((func(_1057_v15 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1058_i9 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_1057_v15)
					}
				})(_993_v15)))).Cardinality())))
				var _rhs98 bool = _994_v16
				_ = _rhs98
				var _rhs99 _dafny.Sequence = _1046_v53
				_ = _rhs99
				var _rhs100 _dafny.Sequence = (func() _dafny.Sequence {
					if (_dafny.IntOfInt64(226)).Cmp(_992_v14) == 0 {
						return _dafny.SeqOf((_1047_v54).Dtor_cf55())
					}
					return _dafny.SeqOf(_996_i5, _1042_cf47, Companion_Default___.Fm32(_996_i5, globalState), (_1056_v59).Cardinality())
				})()
				_ = _rhs100
				var _rhs101 bool = _994_v16
				_ = _rhs101
				_1043_cf46 = _rhs98
				_1046_v53 = _rhs99
				_1055_v58 = _rhs100
				_991_v13 = _rhs101
			} else if _source12.Is_DC29() {
				var _1059___mcc_h4 bool = _source12.Get_().(D12_DC29).Cf48
				_ = _1059___mcc_h4
				var _1060___mcc_h5 _dafny.Int = _source12.Get_().(D12_DC29).Cf49
				_ = _1060___mcc_h5
				var _1061___mcc_h6 bool = _source12.Get_().(D12_DC29).Cf50
				_ = _1061___mcc_h6
				var _1062___mcc_h7 _dafny.Int = _source12.Get_().(D12_DC29).Cf51
				_ = _1062___mcc_h7
				var _1063___mcc_h8 bool = _source12.Get_().(D12_DC29).Cf52
				_ = _1063___mcc_h8
				var _1064_cf52 bool = _1063___mcc_h8
				_ = _1064_cf52
				var _1065_cf51 _dafny.Int = _1062___mcc_h7
				_ = _1065_cf51
				var _1066_cf50 bool = _1061___mcc_h6
				_ = _1066_cf50
				var _1067_cf49 _dafny.Int = _1060___mcc_h5
				_ = _1067_cf49
				var _1068_cf48 bool = _1059___mcc_h4
				_ = _1068_cf48
				var _1069_v60 _dafny.Sequence
				_ = _1069_v60
				_1069_v60 = _dafny.SeqOf(_1068_cf48, _1064_cf52)
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_976_v5), 0))
				_ = _index166
				(_976_v5).ArraySet1((_1069_v60).Select((Companion_Default___.SafeIndex(_996_i5, _dafny.IntOfUint32((_1069_v60).Cardinality()))).Uint32()).(bool), (_index166).Int())
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_976_v5), 0))
				_ = _index167
				(_976_v5).ArraySet1((_994_v16) || (!(true) || (_991_v13)), (_index167).Int())
				var _1070_v61 _dafny.Sequence
				_ = _1070_v61
				_1070_v61 = _dafny.UnicodeSeqOfUtf8Bytes("wnk")
				var _1071_v62 D6
				_ = _1071_v62
				_1071_v62 = Companion_D6_.Create_DC14_(_1070_v61)
				_1071_v62 = _1071_v62
				var _1072_v63 _dafny.Map
				_ = _1072_v63
				_1072_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1037_v48, _1065_cf51)
				var _1073_v65 _dafny.Sequence
				_ = _1073_v65
				_1073_v65 = _dafny.SeqOf((_1072_v63).Update(_1037_v48, p0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1037_v48, _992_v14))
				var _1074_v66 _dafny.Array
				_ = _1074_v66
				var _nwElement0_38 _dafny.Map = _1072_v63
				_ = _nwElement0_38
				var _nw172 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(18))
				_ = _nw172
				(_nw172).ArraySet1(_nwElement0_38, 0)
				(_nw172).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1037_v48, (_dafny.Zero).Minus((_972_v1).Cardinality())), 1)
				(_nw172).ArraySet1(_1072_v63, 2)
				(_nw172).ArraySet1(_1072_v63, 3)
				(_nw172).ArraySet1(_1072_v63, 4)
				(_nw172).ArraySet1(func() _dafny.Map {
					var _coll47 = _dafny.NewMapBuilder()
					_ = _coll47
					for _iter49 := _dafny.Iterate(((_1072_v63).Update(_1037_v48, _1065_cf51)).Keys().Elements()); ; {
						_compr_47, _ok49 := _iter49()
						if !_ok49 {
							break
						}
						var _1075_v64 _dafny.CodePoint
						_1075_v64 = interface{}(_compr_47).(_dafny.CodePoint)
						if ((_1072_v63).Update(_1037_v48, _1065_cf51)).Contains(_1075_v64) {
							_coll47.Add(_1075_v64, _dafny.IntOfInt64(813))
						}
					}
					return _coll47.ToMap()
				}(), 5)
				(_nw172).ArraySet1(_1072_v63, 6)
				(_nw172).ArraySet1(_1072_v63, 7)
				(_nw172).ArraySet1(_1072_v63, 8)
				(_nw172).ArraySet1(_1072_v63, 9)
				(_nw172).ArraySet1(_1072_v63, 10)
				(_nw172).ArraySet1(_1072_v63, 11)
				(_nw172).ArraySet1(_1072_v63, 12)
				(_nw172).ArraySet1(_1072_v63, 13)
				(_nw172).ArraySet1((_1072_v63).Update(_dafny.CodePoint('k'), _996_i5), 14)
				(_nw172).ArraySet1(_1072_v63, 15)
				(_nw172).ArraySet1((_1073_v65).Select((Companion_Default___.SafeIndex(_993_v15, _dafny.IntOfUint32((_1073_v65).Cardinality()))).Uint32()).(_dafny.Map), 16)
				(_nw172).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1037_v48, (func() _dafny.Int {
					if (_972_v1).Contains(_1066_cf50) {
						return (_972_v1).Multiplicity(_1066_cf50)
					}
					return _992_v14
				})()), 17)
				_1074_v66 = _nw172
				var _1076_v67 _dafny.Sequence
				_ = _1076_v67
				_1076_v67 = _dafny.SeqOf(p0)
				var _1077_v68 *C1
				_ = _1077_v68
				var _nw173 *C1 = New_C1_()
				_ = _nw173
				_nw173.Ctor__((_dafny.Zero).Minus(_1065_cf51), (func() _dafny.Array {
					if _994_v16 {
						return _1074_v66
					}
					return _1074_v66
				})(), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1076_v67, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.IntOfUint32((_1076_v67).Cardinality()))).Uint32(), _996_i5), _1076_v67), Companion_D0_.Create_DC1_(Companion_Default___.Fm32(_dafny.IntOfInt64(-714), globalState)))
				_1077_v68 = _nw173
				var _1078_v69 _dafny.Map
				_ = _1078_v69
				_1078_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_972_v1).Difference(_972_v1)).Cardinality(), _dafny.IntOfInt64(111))
				_1078_v69 = (_1078_v69).Update(p0, (_dafny.Zero).Minus((_996_i5).Minus(_1065_cf51)))
			} else {
				var _1079___mcc_h9 _dafny.Set = _source12.Get_().(D12_DC27).Cf45
				_ = _1079___mcc_h9
				var _1080_cf45 _dafny.Set = _1079___mcc_h9
				_ = _1080_cf45
				var _1081_v70 _dafny.Array
				_ = _1081_v70
				var _nw174 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
				_ = _nw174
				_1081_v70 = _nw174
				var _1082_v71 T0
				_ = _1082_v71
				var _nw175 *C2 = New_C2_()
				_ = _nw175
				_nw175.Ctor__()
				_1082_v71 = _nw175
				var _1083_v72 _dafny.Map
				_ = _1083_v72
				_1083_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1082_v71, true)
				var _1084_v73 _dafny.Sequence
				_ = _1084_v73
				_1084_v73 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(474))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}((func(_1085_v48 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1086_i10 _dafny.Int) _dafny.CodePoint {
						return _1085_v48
					}
				})(_1037_v48)))).Cardinality()), (_1083_v72).Cardinality())
				var _1087_v74 D0
				_ = _1087_v74
				_1087_v74 = Companion_D0_.Create_DC1_(_993_v15)
				var _1088_v75 *C1
				_ = _1088_v75
				var _nw176 *C1 = New_C1_()
				_ = _nw176
				_nw176.Ctor__(_993_v15, _1081_v70, _1084_v73, _1087_v74)
				_1088_v75 = _nw176
				var _1089_v76 _dafny.Sequence
				_ = _1089_v76
				_1089_v76 = _dafny.UnicodeSeqOfUtf8Bytes("dkjicxbn")
				var _1090_v77 *C0
				_ = _1090_v77
				var _nw177 *C0 = New_C0_()
				_ = _nw177
				_nw177.Ctor__(_dafny.Companion_Sequence_.Concatenate(_1089_v76, _1089_v76))
				_1090_v77 = _nw177
				var _1091_v78 _dafny.Array
				_ = _1091_v78
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_22
				var _nw178 _dafny.Array
				_ = _nw178
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw178 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Int = (func(_1092_v75 *C1) func(_dafny.Int) _dafny.Int {
						return func(_1093_i11 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1093_i11, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(_1092_v75.F9))).Cardinality()))
						}
					})(_1088_v75)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw178 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw178).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw178).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_1091_v78 = _nw178
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_1091_v78), 0))
				_ = _index168
				(_1091_v78).ArraySet1(_dafny.IntOfUint32((_1090_v77.F8).Cardinality()), (_index168).Int())
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_1091_v78), 0))
				_ = _index169
				(_1091_v78).ArraySet1(_996_i5, (_index169).Int())
				var _1094_v79 _dafny.Set
				_ = _1094_v79
				_1094_v79 = _dafny.SetOf(_996_i5)
				_1094_v79 = _1094_v79
			}
			_1037_v48 = _1037_v48
		}
		var _1095_v80 _dafny.Sequence
		_ = _1095_v80
		_1095_v80 = _dafny.UnicodeSeqOfUtf8Bytes("ohddunlw")
		var _1096_v81 _dafny.Set
		_ = _1096_v81
		_1096_v81 = _dafny.SetOf(_994_v16, _971_v0)
		var _1097_v82 _dafny.Sequence
		_ = _1097_v82
		_1097_v82 = _dafny.SeqOf(_993_v15, _993_v15, (_1096_v81).Cardinality())
		var _1098_v83 _dafny.Map
		_ = _1098_v83
		_1098_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf(_1095_v80, _1095_v80), _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1097_v82, _1097_v82)))
		r0 = _1098_v83
		return r0
	}
}
func (_this *C3) M11(p0 _dafny.CodePoint, p1 _dafny.Array, p2 _dafny.MultiSet, p3 _dafny.Int, globalState *GlobalState) (_dafny.Map, bool, bool, _dafny.Set) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Set = _dafny.EmptySet
		_ = r3
		var _1099_v0 bool
		_ = _1099_v0
		_1099_v0 = true
		var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))
		_ = _index170
		(p1).ArraySet1(_1099_v0, (_index170).Int())
		var _1100_v1 _dafny.Sequence
		_ = _1100_v1
		_1100_v1 = _dafny.SeqOf(p3)
		var _1101_v2 _dafny.Sequence
		_ = _1101_v2
		_1101_v2 = _dafny.UnicodeSeqOfUtf8Bytes("h")
		var _1102_v3 _dafny.Set
		_ = _1102_v3
		_1102_v3 = _dafny.SetOf(_dafny.MultiSetOf(p3))
		var _1103_v4 _dafny.Map
		_ = _1103_v4
		_1103_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1101_v2).Cardinality()), _1102_v3)
		var _1104_v5 _dafny.Array
		_ = _1104_v5
		var _nwElement0_39 _dafny.Int = p3
		_ = _nwElement0_39
		var _nw179 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(3))
		_ = _nw179
		(_nw179).ArraySet1(_nwElement0_39, 0)
		(_nw179).ArraySet1(p3, 1)
		(_nw179).ArraySet1(p3, 2)
		_1104_v5 = _nw179
		var _1105_v6 _dafny.Map
		_ = _1105_v6
		_1105_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1104_v5, p3)
		var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))
		_ = _index171
		var _rhs102 bool = true
		_ = _rhs102
		var _rhs103 bool = !(((func() _dafny.Set {
			if (_1103_v4).Contains(_dafny.IntOfInt64(-931)) {
				return (_1103_v4).Get(_dafny.IntOfInt64(-931)).(_dafny.Set)
			}
			return _1102_v3
		})()).IsDisjointFrom(_1102_v3))
		_ = _rhs103
		var _rhs104 _dafny.Sequence = _dafny.SeqOf((_1105_v6).Cardinality())
		_ = _rhs104
		var _rhs105 bool = _1099_v0
		_ = _rhs105
		var _lhs66 _dafny.Array = p1
		_ = _lhs66
		var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))
		_ = _lhs67
		(_lhs66).ArraySet1(_rhs102, (_lhs67).Int())
		_1099_v0 = _rhs103
		_1100_v1 = _rhs104
		r1 = _rhs105
		if _1099_v0 {
			_1099_v0 = _1099_v0
			_1101_v2 = _1101_v2
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))
			_ = _index172
			(p1).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool), (_index172).Int())
			var _1106_v7 _dafny.Map
			_ = _1106_v7
			_1106_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1099_v0, (func() _dafny.Sequence {
				if _1099_v0 {
					return _dafny.SeqOf(p0, p0)
				}
				return _1101_v2
			})())
			var _1107_v8 _dafny.Int
			_ = _1107_v8
			_1107_v8 = _dafny.IntOfInt64(-951)
			var _1108_v9 _dafny.Array
			_ = _1108_v9
			var _len0_23 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_23
			var _nw180 _dafny.Array
			_ = _nw180
			if _len0_23.Cmp(_dafny.Zero) == 0 {
				_nw180 = _dafny.NewArray(_len0_23)
			} else {
				var _init23 func(_dafny.Int) _dafny.CodePoint = (func(_1109_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1110_i0 _dafny.Int) _dafny.CodePoint {
						return _1109_p0
					}
				})(p0)
				_ = _init23
				var _element0_23 = _init23(_dafny.Zero)
				_ = _element0_23
				_nw180 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
				(_nw180).ArraySet1CodePoint(_element0_23, 0)
				var _nativeLen0_23 = (_len0_23).Int()
				_ = _nativeLen0_23
				for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
					(_nw180).ArraySet1CodePoint(_init23(_dafny.IntOf(_i0_23)), _i0_23)
				}
			}
			_1108_v9 = _nw180
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1108_v9), 0))
			_ = _index173
			(_1108_v9).ArraySet1CodePoint(_dafny.CodePoint('v'), (_index173).Int())
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1108_v9), 0))
			_ = _index174
			var _rhs106 _dafny.Map = (_1106_v7).Merge(_1106_v7)
			_ = _rhs106
			var _rhs107 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(803))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg76 _dafny.Int) interface{} {
					return coer76(arg76)
				}
			}((func(_1111_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1112_i1 _dafny.Int) _dafny.CodePoint {
					return _1111_p0
				}
			})(p0)))).Cardinality()), (_dafny.IntOfInt64(491)).Minus(_1107_v8))
			_ = _rhs107
			var _rhs108 _dafny.CodePoint = p0
			_ = _rhs108
			var _rhs109 _dafny.Int = _1107_v8
			_ = _rhs109
			var _lhs68 _dafny.Array = _1108_v9
			_ = _lhs68
			var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1108_v9), 0))
			_ = _lhs69
			_1106_v7 = _rhs106
			_1107_v8 = _rhs107
			(_lhs68).ArraySet1CodePoint(_rhs108, (_lhs69).Int())
			_1107_v8 = _rhs109
			r1 = (_1107_v8).Cmp(_1107_v8) == 0
		} else {
			var _1113_v10 _dafny.CodePoint
			_ = _1113_v10
			_1113_v10 = _dafny.CodePoint('c')
			var _1114_v11 _dafny.Int
			_ = _1114_v11
			_1114_v11 = _dafny.IntOfInt64(388)
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1104_v5), 0))
			_ = _index175
			(_1104_v5).ArraySet1((p3).Times(Companion_Default___.Fm32(_1114_v11, globalState)), (_index175).Int())
			var _1115_v12 D7
			_ = _1115_v12
			_1115_v12 = Companion_D7_.Create_DC18_()
			var _1116_v13 _dafny.Map
			_ = _1116_v13
			_1116_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool), _1115_v12)
			var _1117_v14 _dafny.Sequence
			_ = _1117_v14
			_1117_v14 = _dafny.SeqOf(_1116_v13, _1116_v13)
			var _1118_v15 D13
			_ = _1118_v15
			_1118_v15 = Companion_D13_.Create_DC31_((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool), _1114_v11, p3)
			var _1119_v16 _dafny.MultiSet
			_ = _1119_v16
			_1119_v16 = _dafny.MultiSetOf(Companion_Default___.Fm44(true, globalState), _1114_v11, _1114_v11, Companion_Default___.Fm44(_1099_v0, globalState))
			var _1120_v17 _dafny.MultiSet
			_ = _1120_v17
			_1120_v17 = _dafny.MultiSetOf(_1119_v16, _1119_v16, (_1119_v16).Intersection(_dafny.MultiSetFromSeq(_1100_v1)))
			var _1121_v18 _dafny.Sequence
			_ = _1121_v18
			_1121_v18 = _dafny.SeqOf(_1099_v0, _1099_v0)
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1104_v5), 0))
			_ = _index176
			var _rhs110 _dafny.CodePoint = Companion_Default___.Fm51(_1099_v0, p3, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1101_v2).Cardinality()), (_1100_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1117_v14).Cardinality()), _dafny.IntOfUint32((_1100_v1).Cardinality()))).Uint32()).(_dafny.Int)), globalState)
			_ = _rhs110
			var _rhs111 bool = ((func() D13 {
				if _1099_v0 {
					return _1118_v15
				}
				return Companion_D13_.Create_DC31_(_1099_v0, p3, p3)
			})()).Dtor_cf54()
			_ = _rhs111
			var _rhs112 _dafny.Int = (func() _dafny.Int {
				if (_1120_v17).Contains(((_1119_v16).Update(_1114_v11, Companion_Default___.Abs(_dafny.IntOfUint32((_1121_v18).Cardinality())))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(816)))) {
					return (_1120_v17).Multiplicity(((_1119_v16).Update(_1114_v11, Companion_Default___.Abs(_dafny.IntOfUint32((_1121_v18).Cardinality())))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(816))))
				}
				return _1114_v11
			})()
			_ = _rhs112
			var _rhs113 bool = !((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool))
			_ = _rhs113
			var _rhs114 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ghsv")).Cardinality())
			_ = _rhs114
			var _lhs70 _dafny.Array = _1104_v5
			_ = _lhs70
			var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1104_v5), 0))
			_ = _lhs71
			_1113_v10 = _rhs110
			r2 = _rhs111
			_1114_v11 = _rhs112
			r2 = _rhs113
			(_lhs70).ArraySet1(_rhs114, (_lhs71).Int())
			if ((_dafny.Zero).Minus((_1104_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1104_v5), 0))).Int()).(_dafny.Int))).Cmp((_dafny.Zero).Minus(_dafny.IntOfInt64(-737))) <= 0 {
				var _1122_v19 _dafny.Set
				_ = _1122_v19
				_1122_v19 = _dafny.SetOf(_1113_v10, p0, p0)
				var _1123_v20 _dafny.Set
				_ = _1123_v20
				_1123_v20 = _dafny.SetOf(_1121_v18, _dafny.SeqOf(_1099_v0))
				var _1124_v21 _dafny.Map
				_ = _1124_v21
				_1124_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1122_v19).Cardinality(), (_1123_v20).IsSubsetOf(Companion_Default___.Fm55(p3, globalState)))
				var _1125_v22 _dafny.Set
				_ = _1125_v22
				_1125_v22 = _dafny.SetOf(_dafny.IntOfInt64(830))
				_1124_v21 = (_1124_v21).Update(((_dafny.SetOf((_1104_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1104_v5), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_1114_v11), (_1104_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1104_v5), 0))).Int()).(_dafny.Int))).Difference(_1125_v22)).Cardinality(), (Companion_D11_.Create_DC25_(_1113_v10, _1114_v11, false)).Dtor_cf43())
				var _1126_v23 _dafny.Array
				_ = _1126_v23
				var _nw181 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
				_ = _nw181
				_1126_v23 = _nw181
				_1126_v23 = _1126_v23
				var _1127_v24 _dafny.Array
				_ = _1127_v24
				var _nw182 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
				_ = _nw182
				_1127_v24 = _nw182
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_1127_v24), 0))
				_ = _index177
				(_1127_v24).ArraySet1(_1100_v1, (_index177).Int())
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_1127_v24), 0))
				_ = _index178
				(_1127_v24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1100_v1, _1100_v1), (_index178).Int())
				var _1128_v25 _dafny.MultiSet
				_ = _1128_v25
				_1128_v25 = _dafny.MultiSetOf(true)
				_1128_v25 = (_1128_v25).Update((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool), Companion_Default___.Abs(_1114_v11))
				var _1129_v26 _dafny.Sequence
				_ = _1129_v26
				_1129_v26 = _dafny.SeqOf(p1, p1, p1, (func() _dafny.Array {
					if (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool) {
						return p1
					}
					return p1
				})(), p1)
				_1129_v26 = _1129_v26
			} else {
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))
				_ = _index179
				(p1).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1101_v2, _1101_v2), (_index179).Int())
				var _1130_v27 _dafny.Map
				_ = _1130_v27
				_1130_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool))
				_1130_v27 = (_1130_v27).Update(true, !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_1099_v0), _1121_v18))
				r1 = _1099_v0
				var _1131_v28 _dafny.Set
				_ = _1131_v28
				_1131_v28 = _dafny.SetOf(_1099_v0)
				var _1132_v29 _dafny.Map
				_ = _1132_v29
				_1132_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1099_v0, p3)
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))
				_ = _index180
				(p1).ArraySet1((_1131_v28).IsDisjointFrom(Companion_Default___.Fm53(_1132_v29, globalState)), (_index180).Int())
				var _1133_v30 _dafny.Map
				_ = _1133_v30
				_1133_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1099_v0)
				var _rhs115 bool = ((_dafny.Zero).Minus(_1114_v11)).Cmp(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1114_v11, _1099_v0)).Cardinality()).Times((_1133_v30).Cardinality())) >= 0
				_ = _rhs115
				var _rhs116 _dafny.Sequence = _1101_v2
				_ = _rhs116
				var _rhs117 _dafny.Int = (Companion_Default___.Fm44((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool), globalState)).Times(p3)
				_ = _rhs117
				r1 = _rhs115
				_1101_v2 = _rhs116
				_1114_v11 = _rhs117
			}
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1104_v5), 0))
			_ = _index181
			(_1104_v5).ArraySet1((_1104_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1104_v5), 0))).Int()).(_dafny.Int), (_index181).Int())
			var _1134_v31 D12
			_ = _1134_v31
			_1134_v31 = Companion_D12_.Create_DC29_(false, p3, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool), _1114_v11, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool))
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1104_v5), 0))
			_ = _index182
			(_1104_v5).ArraySet1((_1134_v31).Dtor_cf49(), (_index182).Int())
			var _1135_v32 D12
			_ = _1135_v32
			_1135_v32 = Companion_D12_.Create_DC28_((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool), (_dafny.Zero).Minus(_1114_v11))
			var _1136_v33 _dafny.Set
			_ = _1136_v33
			_1136_v33 = _dafny.SetOf((_dafny.Zero).Minus(p3), _1114_v11)
			var _pat_let_tv29 = _1136_v33
			_ = _pat_let_tv29
			var _1137_v34 _dafny.MultiSet
			_ = _1137_v34
			_1137_v34 = _dafny.MultiSetOf(Companion_D12_.Create_DC28_((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool), _1114_v11), _1135_v32, _1135_v32, func(_pat_let26_0 D12) D12 {
				return func(_1138_dt__update__tmp_h0 D12) D12 {
					return func(_pat_let27_0 _dafny.Int) D12 {
						return func(_1139_dt__update_hcf47_h0 _dafny.Int) D12 {
							return Companion_D12_.Create_DC28_((_1138_dt__update__tmp_h0).Dtor_cf46(), _1139_dt__update_hcf47_h0)
						}(_pat_let27_0)
					}((_pat_let_tv29).Cardinality())
				}(_pat_let26_0)
			}(_1135_v32), _1135_v32)
			var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_1104_v5), 0))
			_ = _index183
			(_1104_v5).ArraySet1((_1137_v34).Cardinality(), (_index183).Int())
		}
		var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_1104_v5), 0))
		_ = _index184
		(_1104_v5).ArraySet1(_dafny.IntOfUint32((_1101_v2).Cardinality()), (_index184).Int())
		var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_1104_v5), 0))
		_ = _index185
		(_1104_v5).ArraySet1(p3, (_index185).Int())
		_1099_v0 = !(!(!(_1099_v0)) || ((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool))) || (_1099_v0)
		var _1140_v35 bool
		_ = _1140_v35
		var _1141_v36 _dafny.Int
		_ = _1141_v36
		var _1142_v37 _dafny.Int
		_ = _1142_v37
		var _1143_v38 bool
		_ = _1143_v38
		var _out33 bool
		_ = _out33
		var _out34 _dafny.Int
		_ = _out34
		var _out35 _dafny.Int
		_ = _out35
		var _out36 bool
		_ = _out36
		_out33, _out34, _out35, _out36 = (_this).M12((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool), globalState)
		_1140_v35 = _out33
		_1141_v36 = _out34
		_1142_v37 = _out35
		_1143_v38 = _out36
		var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_1104_v5), 0))
		_ = _index186
		(_1104_v5).ArraySet1(_1141_v36, (_index186).Int())
		var _1144_v39 _dafny.Map
		_ = _1144_v39
		_1144_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1143_v38, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool))
		var _1145_v40 D16
		_ = _1145_v40
		_1145_v40 = Companion_D16_.Create_DC42_(_dafny.IntOfInt64(173), _1141_v36, _1141_v36)
		var _1146_v41 _dafny.Map
		_ = _1146_v41
		_1146_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm30(_1142_v37, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool), _1140_v35, globalState), (_1104_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_1104_v5), 0))).Int()).(_dafny.Int))
		var _pat_let_tv30 = _1146_v41
		_ = _pat_let_tv30
		r0 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1104_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_1104_v5), 0))).Int()).(_dafny.Int), _1144_v39)).Update((func(_pat_let28_0 D16) D16 {
			return func(_1147_dt__update__tmp_h1 D16) D16 {
				return func(_pat_let29_0 _dafny.Int) D16 {
					return func(_1148_dt__update_hcf76_h0 _dafny.Int) D16 {
						return Companion_D16_.Create_DC42_(_1148_dt__update_hcf76_h0, (_1147_dt__update__tmp_h1).Dtor_cf77(), (_1147_dt__update__tmp_h1).Dtor_cf78())
					}(_pat_let29_0)
				}((_pat_let_tv30).Cardinality())
			}(_pat_let28_0)
		}(_1145_v40)).Dtor_cf78(), _1144_v39)
		r1 = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool)
		var _1149_v42 D0
		_ = _1149_v42
		_1149_v42 = Companion_D0_.Create_DC1_(_1141_v36)
		var _1150_v43 _dafny.Sequence
		_ = _1150_v43
		_1150_v43 = _dafny.SeqOf(_1140_v35, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool))
		var _pat_let_tv31 = _1104_v5
		_ = _pat_let_tv31
		var _pat_let_tv32 = _1104_v5
		_ = _pat_let_tv32
		r2 = (_this).Fm5((_dafny.MultiSetOf((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool))).Cardinality(), _1140_v35, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, func(_pat_let30_0 D0) D0 {
			return func(_1151_dt__update__tmp_h2 D0) D0 {
				return func(_pat_let31_0 _dafny.Int) D0 {
					return func(_1152_dt__update_hcf2_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC1_(_1152_dt__update_hcf2_h0)
					}(_pat_let31_0)
				}((_pat_let_tv32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_pat_let_tv31), 0))).Int()).(_dafny.Int))
			}(_pat_let30_0)
		}(_1149_v42))).Update(p0, Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_1150_v43).Cardinality()))), true, globalState)
		var _1153_v44 _dafny.Set
		_ = _1153_v44
		_1153_v44 = _dafny.SetOf(_1146_v41, _1146_v41, _1146_v41, _1146_v41)
		r3 = (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((p1), 0))).Int()).(bool), (_1104_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_1104_v5), 0))).Int()).(_dafny.Int)))).Intersection(_1153_v44)
		return r0, r1, r2, r3
	}
}
func (_this *C3) M12(p0 bool, globalState *GlobalState) (bool, _dafny.Int, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 bool = false
		_ = r3
		var _1154_v0 _dafny.Int
		_ = _1154_v0
		_1154_v0 = _dafny.IntOfInt64(886)
		var _1155_v1 _dafny.Sequence
		_ = _1155_v1
		_1155_v1 = _dafny.SeqOf(_1154_v0)
		var _rhs118 bool = false
		_ = _rhs118
		var _rhs119 _dafny.Int = _1154_v0
		_ = _rhs119
		var _rhs120 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1154_v0), (Companion_Default___.SafeIndex(_1154_v0, _dafny.IntOfUint32((_dafny.SeqOf(_1154_v0)).Cardinality()))).Uint32(), _dafny.IntOfInt64(667))
		_ = _rhs120
		r3 = _rhs118
		r1 = _rhs119
		_1155_v1 = _rhs120
		var _hi9 _dafny.Int = _1154_v0
		_ = _hi9
		for _1156_i0 := _1154_v0; _1156_i0.Cmp(_hi9) < 0; _1156_i0 = _1156_i0.Plus(_dafny.One) {
			var _1157_v2 _dafny.Sequence
			_ = _1157_v2
			_1157_v2 = _dafny.UnicodeSeqOfUtf8Bytes("sxccix")
			if (_1154_v0).Cmp(_dafny.IntOfUint32((_1157_v2).Cardinality())) != 0 {
				var _rhs121 bool = !(p0)
				_ = _rhs121
				var _rhs122 bool = p0
				_ = _rhs122
				r3 = _rhs121
				r3 = _rhs122
				var _1158_v3 _dafny.Array
				_ = _1158_v3
				var _nw183 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
				_ = _nw183
				_1158_v3 = _nw183
				var _1159_v4 _dafny.Map
				_ = _1159_v4
				_1159_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), p0)
				var _1160_v5 _dafny.Set
				_ = _1160_v5
				_1160_v5 = _dafny.SetOf((_1159_v4).Cardinality(), _dafny.IntOfInt64(415))
				var _1161_v7 _dafny.Map
				_ = _1161_v7
				_1161_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1156_i0)
				var _1162_v8 D14
				_ = _1162_v8
				_1162_v8 = Companion_D14_.Create_DC36_(func() _dafny.Map {
					var _coll48 = _dafny.NewMapBuilder()
					_ = _coll48
					for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(381), _dafny.IntOfInt64(529))); ; {
						_compr_48, _ok50 := _iter50()
						if !_ok50 {
							break
						}
						var _1163_v6 _dafny.Int
						_1163_v6 = interface{}(_compr_48).(_dafny.Int)
						if ((_dafny.IntOfInt64(381)).Cmp(_1163_v6) <= 0) && ((_1163_v6).Cmp(_dafny.IntOfInt64(529)) < 0) {
							_coll48.Add((_1163_v6).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _1156_i0)
						}
					}
					return _coll48.ToMap()
				}(), _1161_v7)
				var _1164_v9 _dafny.Map
				_ = _1164_v9
				_1164_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1157_v2).Cardinality()), _1156_i0)
				var _1165_v10 _dafny.Array
				_ = _1165_v10
				var _nwElement0_40 D14 = _1162_v8
				_ = _nwElement0_40
				var _nw184 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(6))
				_ = _nw184
				(_nw184).ArraySet1(_nwElement0_40, 0)
				(_nw184).ArraySet1(Companion_D14_.Create_DC36_(_1164_v9, _1161_v7), 1)
				(_nw184).ArraySet1(_1162_v8, 2)
				(_nw184).ArraySet1(_1162_v8, 3)
				(_nw184).ArraySet1(_1162_v8, 4)
				(_nw184).ArraySet1(_1162_v8, 5)
				_1165_v10 = _nw184
				var _1166_v11 _dafny.Sequence
				_ = _1166_v11
				_1166_v11 = _dafny.SeqOf(_1165_v10, _1165_v10)
				var _1167_v14 _dafny.Array
				_ = _1167_v14
				var _nwElement0_41 _dafny.Set = _1160_v5
				_ = _nwElement0_41
				var _nw185 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(25))
				_ = _nw185
				(_nw185).ArraySet1(_nwElement0_41, 0)
				(_nw185).ArraySet1(_1160_v5, 1)
				(_nw185).ArraySet1(_1160_v5, 2)
				(_nw185).ArraySet1(_1160_v5, 3)
				(_nw185).ArraySet1(_1160_v5, 4)
				(_nw185).ArraySet1(_dafny.SetOf(_dafny.IntOfUint32((_1166_v11).Cardinality()), _1156_i0), 5)
				(_nw185).ArraySet1(_dafny.SetOf((_dafny.Zero).Minus(_1156_i0), _1156_i0), 6)
				(_nw185).ArraySet1(_1160_v5, 7)
				(_nw185).ArraySet1(_1160_v5, 8)
				(_nw185).ArraySet1(_1160_v5, 9)
				(_nw185).ArraySet1(_1160_v5, 10)
				(_nw185).ArraySet1(_dafny.SetOf(Companion_Default___.Fm44(p0, globalState)), 11)
				(_nw185).ArraySet1(_1160_v5, 12)
				(_nw185).ArraySet1(_1160_v5, 13)
				(_nw185).ArraySet1(func() _dafny.Set {
					var _coll49 = _dafny.NewBuilder()
					_ = _coll49
					for _iter51 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfInt64(289))); ; {
						_compr_49, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _1168_v12 _dafny.Int
						_1168_v12 = interface{}(_compr_49).(_dafny.Int)
						if ((_1168_v12).Sign() != -1) && ((_1168_v12).Cmp(_dafny.IntOfInt64(289)) < 0) {
							_coll49.Add((_1168_v12).Plus(_dafny.IntOfUint32((_1155_v1).Cardinality())))
						}
					}
					return _coll49.ToSet()
				}(), 14)
				(_nw185).ArraySet1(_1160_v5, 15)
				(_nw185).ArraySet1(_1160_v5, 16)
				(_nw185).ArraySet1(_1160_v5, 17)
				(_nw185).ArraySet1(_1160_v5, 18)
				(_nw185).ArraySet1(_1160_v5, 19)
				(_nw185).ArraySet1(_1160_v5, 20)
				(_nw185).ArraySet1(_dafny.SetOf(_1154_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-591))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg77 _dafny.Int) interface{} {
						return coer77(arg77)
					}
				}(func(_1169_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				}))).Cardinality()), _1156_i0, _1156_i0), 21)
				(_nw185).ArraySet1(_1160_v5, 22)
				(_nw185).ArraySet1(_1160_v5, 23)
				(_nw185).ArraySet1(func() _dafny.Set {
					var _coll50 = _dafny.NewBuilder()
					_ = _coll50
					for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(578), _dafny.IntOfInt64(846))); ; {
						_compr_50, _ok52 := _iter52()
						if !_ok52 {
							break
						}
						var _1170_v13 _dafny.Int
						_1170_v13 = interface{}(_compr_50).(_dafny.Int)
						if ((_dafny.IntOfInt64(578)).Cmp(_1170_v13) <= 0) && ((_1170_v13).Cmp(_dafny.IntOfInt64(846)) < 0) {
							_coll50.Add((_1170_v13).Minus((_1159_v4).Cardinality()))
						}
					}
					return _coll50.ToSet()
				}(), 24)
				_1167_v14 = _nw185
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_1158_v3), 0))
				_ = _index187
				(_1158_v3).ArraySet1(_1167_v14, (_index187).Int())
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_1158_v3), 0))
				_ = _index188
				(_1158_v3).ArraySet1(_1167_v14, (_index188).Int())
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_24
				var _nw186 _dafny.Array
				_ = _nw186
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw186 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.Set = (func(_1171_v5 _dafny.Set) func(_dafny.Int) _dafny.Set {
						return func(_1172_i2 _dafny.Int) _dafny.Set {
							return _1171_v5
						}
					})(_1160_v5)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw186 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw186).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw186).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_1167_v14 = _nw186
				var _1173_v15 _dafny.CodePoint
				_ = _1173_v15
				_1173_v15 = _dafny.CodePoint('f')
				var _1174_v16 *C0
				_ = _1174_v16
				var _nw187 *C0 = New_C0_()
				_ = _nw187
				_nw187.Ctor__(_dafny.Companion_Sequence_.Concatenate(_1157_v2, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(790))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg78 _dafny.Int) interface{} {
						return coer78(arg78)
					}
				}(func(_1175_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				})), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1157_v2).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(790))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg79 _dafny.Int) interface{} {
						return coer79(arg79)
					}
				}(func(_1176_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}))).Cardinality()))).Uint32(), _1173_v15)))
				_1174_v16 = _nw187
				var _1177_v17 D0
				_ = _1177_v17
				_1177_v17 = Companion_D0_.Create_DC1_((_1161_v7).Cardinality())
				var _1178_v18 _dafny.Map
				_ = _1178_v18
				_1178_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1173_v15, _1177_v17)
				_1159_v4 = (_1159_v4).Update((_this).Fm5(_1156_i0, p0, _1178_v18, p0, globalState), !_dafny.Companion_Sequence_.Equal(_1157_v2, _1157_v2))
			} else {
				var _1179_v19 D7
				_ = _1179_v19
				_1179_v19 = Companion_D7_.Create_DC18_()
				_1179_v19 = _1179_v19
				r3 = p0
				var _1180_v20 _dafny.MultiSet
				_ = _1180_v20
				_1180_v20 = _dafny.MultiSetOf(true, p0)
				var _1181_v21 _dafny.Map
				_ = _1181_v21
				_1181_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1156_i0, p0)
				var _1182_v22 _dafny.Sequence
				_ = _1182_v22
				_1182_v22 = _dafny.SeqOf(p0)
				var _1183_v23 _dafny.Map
				_ = _1183_v23
				_1183_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_1182_v22), p0)
				var _1184_v24 _dafny.Sequence
				_ = _1184_v24
				_1184_v24 = _dafny.SeqOf(p0, Companion_Default___.Fm30(_1156_i0, (func() bool {
					if (_1181_v21).Contains((_1183_v23).Cardinality()) {
						return (_1181_v21).Get((_1183_v23).Cardinality()).(bool)
					}
					return p0
				})(), true, globalState), p0, p0)
				var _1185_v25 D4
				_ = _1185_v25
				_1185_v25 = Companion_D4_.Create_DC11_(_1154_v0, p0, p0, p0, p0)
				r3 = !((func() bool {
					if p0 {
						return (_1180_v20).IsDisjointFrom(_dafny.MultiSetFromSeq(_1182_v22))
					}
					return (_1185_v25).Dtor_cf26()
				})())
				var _1186_v26 _dafny.CodePoint
				_ = _1186_v26
				_1186_v26 = _dafny.CodePoint('k')
				_1186_v26 = (func() _dafny.CodePoint {
					if p0 {
						return _1186_v26
					}
					return (_1157_v2).Select((Companion_Default___.SafeIndex(_1154_v0, _dafny.IntOfUint32((_1157_v2).Cardinality()))).Uint32()).(_dafny.CodePoint)
				})()
				var _1187_v27 _dafny.Array
				_ = _1187_v27
				var _nwElement0_42 _dafny.Int = _1154_v0
				_ = _nwElement0_42
				var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(2))
				_ = _nw188
				(_nw188).ArraySet1(_nwElement0_42, 0)
				(_nw188).ArraySet1(_1154_v0, 1)
				_1187_v27 = _nw188
				var _1188_v28 _dafny.Set
				_ = _1188_v28
				_1188_v28 = _dafny.SetOf(_1187_v27)
				var _1189_v29 _dafny.MultiSet
				_ = _1189_v29
				_1189_v29 = _dafny.MultiSetOf(_1156_i0, (_1188_v28).Cardinality())
				var _1190_v30 _dafny.Map
				_ = _1190_v30
				_1190_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1189_v29, p0)
				r3 = ((func() bool {
					if (_1190_v30).Contains(_1189_v29) {
						return (_1190_v30).Get(_1189_v29).(bool)
					}
					return true
				})()) == (p0)
			}
			r0 = p0
			var _1191_v31 _dafny.CodePoint
			_ = _1191_v31
			_1191_v31 = _dafny.CodePoint('a')
			var _1192_v32 _dafny.Map
			_ = _1192_v32
			_1192_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1156_i0)
			var _1193_v33 _dafny.Map
			_ = _1193_v33
			_1193_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1191_v31, (_1192_v32).Cardinality())
			var _1194_v34 _dafny.Map
			_ = _1194_v34
			_1194_v34 = _1193_v33
			var _1195_v35 D11
			_ = _1195_v35
			_1195_v35 = Companion_D11_.Create_DC25_(_1191_v31, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1194_v34, _1154_v0)).Update(_1194_v34, _1156_i0)).Cardinality(), p0)
			var _source13 D11 = _1195_v35
			_ = _source13
			if _source13.Is_DC25() {
				var _1196___mcc_h0 _dafny.CodePoint = _source13.Get_().(D11_DC25).Cf41
				_ = _1196___mcc_h0
				var _1197___mcc_h1 _dafny.Int = _source13.Get_().(D11_DC25).Cf42
				_ = _1197___mcc_h1
				var _1198___mcc_h2 bool = _source13.Get_().(D11_DC25).Cf43
				_ = _1198___mcc_h2
				var _1199_cf43 bool = _1198___mcc_h2
				_ = _1199_cf43
				var _1200_cf42 _dafny.Int = _1197___mcc_h1
				_ = _1200_cf42
				var _1201_cf41 _dafny.CodePoint = _1196___mcc_h0
				_ = _1201_cf41
				var _1202_v36 _dafny.Map
				_ = _1202_v36
				_1202_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1154_v0, _1156_i0)
				var _1203_v37 _dafny.Map
				_ = _1203_v37
				_1203_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1156_i0, p0)
				var _1204_v38 _dafny.Sequence
				_ = _1204_v38
				_1204_v38 = _dafny.SeqOf(_1199_cf43)
				var _1205_v39 _dafny.Set
				_ = _1205_v39
				_1205_v39 = _dafny.SetOf(_1204_v38)
				var _1206_v40 _dafny.Array
				_ = _1206_v40
				var _nwElement0_43 _dafny.Int = _1156_i0
				_ = _nwElement0_43
				var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(28))
				_ = _nw189
				(_nw189).ArraySet1(_nwElement0_43, 0)
				(_nw189).ArraySet1(_1154_v0, 1)
				(_nw189).ArraySet1(_1156_i0, 2)
				(_nw189).ArraySet1((_1154_v0).Times((func() _dafny.Int {
					if (_1202_v36).Contains(_1156_i0) {
						return (_1202_v36).Get(_1156_i0).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(413))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg80 _dafny.Int) interface{} {
							return coer80(arg80)
						}
					}((func(_1207_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1208_i4 _dafny.Int) _dafny.CodePoint {
							return _1207_v31
						}
					})(_1191_v31)))).Cardinality())
				})()), 3)
				(_nw189).ArraySet1(_1156_i0, 4)
				(_nw189).ArraySet1(Companion_Default___.Fm32(_1156_i0, globalState), 5)
				(_nw189).ArraySet1((_dafny.IntOfUint32((_1157_v2).Cardinality())).Plus(_1154_v0), 6)
				(_nw189).ArraySet1(_1156_i0, 7)
				(_nw189).ArraySet1(_1156_i0, 8)
				(_nw189).ArraySet1(((_1155_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("y")).Cardinality()), _dafny.IntOfUint32((_1155_v1).Cardinality()))).Uint32()).(_dafny.Int)).Minus((_1203_v37).Cardinality()), 9)
				(_nw189).ArraySet1(((_1205_v39).Intersection(_1205_v39)).Cardinality(), 10)
				(_nw189).ArraySet1(_1200_cf42, 11)
				(_nw189).ArraySet1(_1156_i0, 12)
				(_nw189).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(793))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg81 _dafny.Int) interface{} {
						return coer81(arg81)
					}
				}((func(_1209_cf41 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1210_i5 _dafny.Int) _dafny.CodePoint {
						return _1209_cf41
					}
				})(_1201_cf41)))).Cardinality()), 13)
				(_nw189).ArraySet1(_1156_i0, 14)
				(_nw189).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(487), _1200_cf42), 15)
				(_nw189).ArraySet1((_1156_i0).Times(_dafny.IntOfUint32((_1157_v2).Cardinality())), 16)
				(_nw189).ArraySet1(_1200_cf42, 17)
				(_nw189).ArraySet1((_dafny.Zero).Minus((_1200_cf42).Plus(_1156_i0)), 18)
				(_nw189).ArraySet1(_1200_cf42, 19)
				(_nw189).ArraySet1(_1200_cf42, 20)
				(_nw189).ArraySet1((_dafny.Zero).Minus(_1154_v0), 21)
				(_nw189).ArraySet1((_1200_cf42).Times(_dafny.IntOfUint32((_1157_v2).Cardinality())), 22)
				(_nw189).ArraySet1(_1156_i0, 23)
				(_nw189).ArraySet1((func() _dafny.Int {
					if p0 {
						return _1156_i0
					}
					return _dafny.IntOfInt64(940)
				})(), 24)
				(_nw189).ArraySet1(Companion_Default___.SafeDivisionInt(_1154_v0, _1156_i0), 25)
				(_nw189).ArraySet1((_dafny.IntOfInt64(727)).Minus((_dafny.Zero).Minus((_dafny.SetOf(_1201_cf41)).Cardinality())), 26)
				(_nw189).ArraySet1(_1156_i0, 27)
				_1206_v40 = _nw189
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_1206_v40), 0))
				_ = _index189
				(_1206_v40).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.Fm32(_1154_v0, globalState), _dafny.IntOfInt64(824)), (_index189).Int())
				var _1211_v41 _dafny.Map
				_ = _1211_v41
				_1211_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1206_v40)
				var _1212_v42 _dafny.Map
				_ = _1212_v42
				_1212_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1211_v41, _1200_cf42)
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_1206_v40), 0))
				_ = _index190
				(_1206_v40).ArraySet1((func() _dafny.Int {
					if (_1212_v42).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1206_v40)) {
						return (_1212_v42).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1206_v40)).(_dafny.Int)
					}
					return _1154_v0
				})(), (_index190).Int())
				var _1213_v43 _dafny.Array
				_ = _1213_v43
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_25
				var _nw190 _dafny.Array
				_ = _nw190
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw190 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) bool = (func(_1214_cf43 bool) func(_dafny.Int) bool {
						return func(_1215_i6 _dafny.Int) bool {
							return _1214_cf43
						}
					})(_1199_cf43)
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw190 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw190).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw190).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_1213_v43 = _nw190
				var _1216_v44 _dafny.Map
				_ = _1216_v44
				_1216_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1202_v36).Cardinality(), _1213_v43)
				var _1217_v45 _dafny.Sequence
				_ = _1217_v45
				_1217_v45 = _dafny.SeqOf((func() _dafny.Array {
					if (_1216_v44).Contains((_1206_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_1206_v40), 0))).Int()).(_dafny.Int)) {
						return (_1216_v44).Get((_1206_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_1206_v40), 0))).Int()).(_dafny.Int)).(_dafny.Array)
					}
					return _1213_v43
				})())
				var _1218_v46 _dafny.MultiSet
				_ = _1218_v46
				_1218_v46 = _dafny.MultiSetOf(_1213_v43)
				var _1219_v47 D16
				_ = _1219_v47
				_1219_v47 = Companion_D16_.Create_DC41_((_1217_v45).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1218_v46).Contains(_1213_v43) {
						return (_1218_v46).Multiplicity(_1213_v43)
					}
					return _dafny.IntOfUint32((_1204_v38).Cardinality())
				})(), _dafny.IntOfUint32((_1217_v45).Cardinality()))).Uint32()).(_dafny.Array))
				_1219_v47 = _1219_v47
				var _1220_v48 _dafny.Set
				_ = _1220_v48
				_1220_v48 = _dafny.SetOf(_1200_cf42, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(956))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg82 _dafny.Int) interface{} {
						return coer82(arg82)
					}
				}(func(_1221_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}))).Cardinality()))
				var _rhs123 _dafny.Set = _1220_v48
				_ = _rhs123
				var _rhs124 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(411), _1154_v0)
				_ = _rhs124
				var _rhs125 _dafny.Int = (_1154_v0).Times(_1200_cf42)
				_ = _rhs125
				var _rhs126 bool = (func() bool {
					if p0 {
						return false
					}
					return _1199_cf43
				})()
				_ = _rhs126
				var _rhs127 _dafny.Int = (func() _dafny.Int {
					if !(p0) {
						return (_1206_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_1206_v40), 0))).Int()).(_dafny.Int)
					}
					return Companion_Default___.SafeDivisionInt(_1200_cf42, _1156_i0)
				})()
				_ = _rhs127
				_1220_v48 = _rhs123
				r1 = _rhs124
				r1 = _rhs125
				r3 = _rhs126
				_1154_v0 = _rhs127
				var _1222_v49 _dafny.Sequence
				_ = _1222_v49
				_1222_v49 = _dafny.SeqOf(_1204_v38)
				var _1223_v50 D0
				_ = _1223_v50
				_1223_v50 = Companion_D0_.Create_DC1_((_1206_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_1206_v40), 0))).Int()).(_dafny.Int))
				var _1224_v51 _dafny.Map
				_ = _1224_v51
				_1224_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1201_cf41, _1223_v50)
				r3 = (_this).Fm5((_dafny.Zero).Minus((_1156_i0).Plus(_dafny.IntOfUint32((_1222_v49).Cardinality()))), p0, _1224_v51, p0, globalState)
			} else if _source13.Is_DC24() {
				var _1225___mcc_h3 _dafny.Map = _source13.Get_().(D11_DC24).Cf40
				_ = _1225___mcc_h3
				var _1226_cf40 _dafny.Map = _1225___mcc_h3
				_ = _1226_cf40
				r3 = p0
				var _1227_v52 _dafny.Sequence
				_ = _1227_v52
				_1227_v52 = _dafny.SeqOf(p0, p0)
				var _1228_v53 D0
				_ = _1228_v53
				_1228_v53 = Companion_D0_.Create_DC1_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1191_v31, _1156_i0)).Cardinality())
				var _1229_v54 _dafny.Map
				_ = _1229_v54
				_1229_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p0, (_1227_v52).Select((Companion_Default___.SafeIndex(_1156_i0, _dafny.IntOfUint32((_1227_v52).Cardinality()))).Uint32()).(bool), (_this).Fm5(_1156_i0, p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1191_v31, _1228_v53), p0, globalState), p0, p0), _1154_v0)
				var _1230_v55 _dafny.Set
				_ = _1230_v55
				_1230_v55 = _dafny.SetOf(p0)
				var _1231_v56 _dafny.Sequence
				_ = _1231_v56
				_1231_v56 = _dafny.SeqOf(_dafny.SetOf(p0, p0))
				_1229_v54 = (func() _dafny.Map {
					if ((_1231_v56).Select((Companion_Default___.SafeIndex(_1156_i0, _dafny.IntOfUint32((_1231_v56).Cardinality()))).Uint32()).(_dafny.Set)).IsProperSubsetOf(_1230_v55) {
						return Companion_Default___.Fm56(false, globalState)
					}
					return (func() _dafny.Map {
						if p0 {
							return _1229_v54
						}
						return _1229_v54
					})()
				})()
				var _1232_v57 _dafny.Array
				_ = _1232_v57
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_26
				var _nw191 _dafny.Array
				_ = _nw191
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw191 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) bool = (func(_1233_p0 bool) func(_dafny.Int) bool {
						return func(_1234_i8 _dafny.Int) bool {
							return _1233_p0
						}
					})(p0)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw191 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw191).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw191).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_1232_v57 = _nw191
				var _1235_v58 _dafny.MultiSet
				_ = _1235_v58
				_1235_v58 = _dafny.MultiSetOf(p0)
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_1232_v57), 0))
				_ = _index191
				(_1232_v57).ArraySet1(!((_1235_v58).IsProperSubsetOf(_1235_v58)), (_index191).Int())
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_1232_v57), 0))
				_ = _index192
				(_1232_v57).ArraySet1(!((_1227_v52).Select((Companion_Default___.SafeIndex((_1154_v0).Plus((Companion_Default___.Fm57(p0, p0, p0, globalState)).Cardinality()), _dafny.IntOfUint32((_1227_v52).Cardinality()))).Uint32()).(bool)), (_index192).Int())
				var _1236_v59 _dafny.Array
				_ = _1236_v59
				var _nw192 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
				_ = _nw192
				_1236_v59 = _nw192
				_1236_v59 = _1236_v59
			} else {
				var _1237___mcc_h4 D11 = _source13.Get_().(D11_DC26).Cf44
				_ = _1237___mcc_h4
				var _1238_cf44 D11 = _1237___mcc_h4
				_ = _1238_cf44
				var _1239_v60 _dafny.Map
				_ = _1239_v60
				_1239_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1156_i0, (_1154_v0).Cmp(_1156_i0) == 0)
				_1239_v60 = _1239_v60
				var _1240_v61 *C2
				_ = _1240_v61
				var _nw193 *C2 = New_C2_()
				_ = _nw193
				_nw193.Ctor__()
				_1240_v61 = _nw193
				var _1241_v62 _dafny.MultiSet
				_ = _1241_v62
				_1241_v62 = _dafny.MultiSetOf(p0)
				r0 = ((_1241_v62).IsDisjointFrom(_1241_v62)) || (p0)
				var _1242_v63 _dafny.Array
				_ = _1242_v63
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_27
				var _nw194 _dafny.Array
				_ = _nw194
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw194 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) bool = (func(_1243_p0 bool) func(_dafny.Int) bool {
						return func(_1244_i9 _dafny.Int) bool {
							return !(_1243_p0)
						}
					})(p0)
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw194 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw194).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw194).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_1242_v63 = _nw194
				var _1245_v64 _dafny.Map
				_ = _1245_v64
				_1245_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1154_v0, _1242_v63)
				var _1246_v65 _dafny.Map
				_ = _1246_v65
				_1246_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1157_v2, _1157_v2), _1245_v64)
				_1246_v65 = (_1246_v65).Update(_1157_v2, _1245_v64)
			}
			r3 = p0
		}
		var _1247_v66 _dafny.Sequence
		_ = _1247_v66
		_1247_v66 = _dafny.SeqOf(p0)
		var _1248_v67 _dafny.Set
		_ = _1248_v67
		_1248_v67 = _dafny.SetOf(_dafny.IntOfInt64(980))
		var _1249_v69 _dafny.Array
		_ = _1249_v69
		var _nwElement0_44 bool = p0
		_ = _nwElement0_44
		var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(8))
		_ = _nw195
		(_nw195).ArraySet1(_nwElement0_44, 0)
		(_nw195).ArraySet1(p0, 1)
		(_nw195).ArraySet1(p0, 2)
		(_nw195).ArraySet1(true, 3)
		(_nw195).ArraySet1(false, 4)
		(_nw195).ArraySet1(p0, 5)
		(_nw195).ArraySet1(p0, 6)
		(_nw195).ArraySet1(p0, 7)
		_1249_v69 = _nw195
		var _1250_v70 _dafny.MultiSet
		_ = _1250_v70
		_1250_v70 = _dafny.MultiSetOf(_1249_v69)
		var _1251_v71 D4
		_ = _1251_v71
		_1251_v71 = Companion_D4_.Create_DC11_(_1154_v0, p0, p0, p0, !(p0))
		var _rhs128 bool = !(((_1248_v67).Intersection(func() _dafny.Set {
			var _coll51 = _dafny.NewBuilder()
			_ = _coll51
			for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(445), _dafny.IntOfInt64(117))); ; {
				_compr_51, _ok53 := _iter53()
				if !_ok53 {
					break
				}
				var _1252_v68 _dafny.Int
				_1252_v68 = interface{}(_compr_51).(_dafny.Int)
				if ((_dafny.IntOfInt64(445)).Cmp(_1252_v68) <= 0) && ((_1252_v68).Cmp(_dafny.IntOfInt64(117)) < 0) {
					_coll51.Add((_1252_v68).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yjfgxccy")).Cardinality())))
				}
			}
			return _coll51.ToSet()
		}())).IsSubsetOf(_1248_v67))
		_ = _rhs128
		var _rhs129 bool = p0
		_ = _rhs129
		var _rhs130 bool = !(_1250_v70).Equals((_1250_v70).Intersection(_1250_v70))
		_ = _rhs130
		var _rhs131 bool = (!(p0) || ((_1251_v71).Dtor_cf23())) && (p0)
		_ = _rhs131
		var _rhs132 _dafny.Sequence = Companion_Default___.Fm38(p0, _1154_v0, p0, globalState)
		_ = _rhs132
		r0 = _rhs128
		r0 = _rhs129
		r3 = _rhs130
		r3 = _rhs131
		_1247_v66 = _rhs132
		var _1253_v72 bool
		_ = _1253_v72
		_1253_v72 = p0
		var _1254_v73 _dafny.MultiSet
		_ = _1254_v73
		_1254_v73 = _dafny.MultiSetOf(p0)
		var _1255_v74 _dafny.Map
		_ = _1255_v74
		_1255_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1253_v72, _1254_v73)
		_1255_v74 = (_1255_v74).Update(_1253_v72, (_dafny.MultiSetOf(p0, p0)).Intersection(_1254_v73))
		var _1256_v75 _dafny.Map
		_ = _1256_v75
		_1256_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1154_v0).Cmp(_1154_v0) == 0, _dafny.IntOfUint32((_1247_v66).Cardinality()))
		_1256_v75 = _1256_v75
		var _1257_v76 _dafny.Array
		_ = _1257_v76
		var _len0_28 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_28
		var _nw196 _dafny.Array
		_ = _nw196
		if _len0_28.Cmp(_dafny.Zero) == 0 {
			_nw196 = _dafny.NewArray(_len0_28)
		} else {
			var _init28 func(_dafny.Int) _dafny.Int = (func(_1258_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1259_i10 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_1259_i10, _1258_v0)
				}
			})(_1154_v0)
			_ = _init28
			var _element0_28 = _init28(_dafny.Zero)
			_ = _element0_28
			_nw196 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
			(_nw196).ArraySet1(_element0_28, 0)
			var _nativeLen0_28 = (_len0_28).Int()
			_ = _nativeLen0_28
			for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
				(_nw196).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
			}
		}
		_1257_v76 = _nw196
		_1257_v76 = _1257_v76
		var _1260_v77 _dafny.Sequence
		_ = _1260_v77
		_1260_v77 = _dafny.UnicodeSeqOfUtf8Bytes("qtblq")
		var _1261_v78 D9
		_ = _1261_v78
		_1261_v78 = Companion_D9_.Create_DC21_(_1254_v73)
		var _1262_v79 D15
		_ = _1262_v79
		_1262_v79 = Companion_D15_.Create_DC39_(_1260_v77, _1261_v78)
		var _pat_let_tv33 = p0
		_ = _pat_let_tv33
		var _pat_let_tv34 = p0
		_ = _pat_let_tv34
		var _pat_let_tv35 = p0
		_ = _pat_let_tv35
		var _pat_let_tv36 = p0
		_ = _pat_let_tv36
		var _pat_let_tv37 = _1261_v78
		_ = _pat_let_tv37
		r0 = func(_source14 D15) bool {
			if _source14.Is_DC38() {
				var _1263___mcc_h5 _dafny.CodePoint = _source14.Get_().(D15_DC38).Cf71
				_ = _1263___mcc_h5
				var _1264_cf71 _dafny.CodePoint = _1263___mcc_h5
				_ = _1264_cf71
				return _pat_let_tv33
			} else if _source14.Is_DC39() {
				var _1265___mcc_h6 _dafny.Sequence = _source14.Get_().(D15_DC39).Cf72
				_ = _1265___mcc_h6
				var _1266___mcc_h7 D9 = _source14.Get_().(D15_DC39).Cf73
				_ = _1266___mcc_h7
				var _1267_cf73 D9 = _1266___mcc_h7
				_ = _1267_cf73
				var _1268_cf72 _dafny.Sequence = _1265___mcc_h6
				_ = _1268_cf72
				return _pat_let_tv34
			} else if _source14.Is_DC37() {
				var _1269___mcc_h8 _dafny.Set = _source14.Get_().(D15_DC37).Cf70
				_ = _1269___mcc_h8
				var _1270_cf70 _dafny.Set = _1269___mcc_h8
				_ = _1270_cf70
				return _pat_let_tv35
			} else {
				var _1271___mcc_h9 D15 = _source14.Get_().(D15_DC40).Cf74
				_ = _1271___mcc_h9
				var _1272_cf74 D15 = _1271___mcc_h9
				_ = _1272_cf74
				return _pat_let_tv36
			}
		}(func(_pat_let32_0 D15) D15 {
			return func(_1273_dt__update__tmp_h0 D15) D15 {
				return func(_pat_let33_0 D9) D15 {
					return func(_1274_dt__update_hcf73_h0 D9) D15 {
						return Companion_D15_.Create_DC39_((_1273_dt__update__tmp_h0).Dtor_cf72(), _1274_dt__update_hcf73_h0)
					}(_pat_let33_0)
				}(_pat_let_tv37)
			}(_pat_let32_0)
		}(_1262_v79))
		r1 = _1154_v0
		r2 = _1154_v0
		r3 = p0
		return r0, r1, r2, r3
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	dummy byte
}

func New_C4_() *C4 {
	_this := C4{}

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

func (_this *C4) Ctor__() {
	{
	}
}
func (_this *C4) Fm5(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 bool, globalState *GlobalState) bool {
	{
		return !(!(false))
	}
}
func (_this *C4) M1(p0 _dafny.Int, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _1275_v0 _dafny.Array
		_ = _1275_v0
		var _len0_29 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_29
		var _nw197 _dafny.Array
		_ = _nw197
		if _len0_29.Cmp(_dafny.Zero) == 0 {
			_nw197 = _dafny.NewArray(_len0_29)
		} else {
			var _init29 func(_dafny.Int) bool = func(_1276_i0 _dafny.Int) bool {
				return true
			}
			_ = _init29
			var _element0_29 = _init29(_dafny.Zero)
			_ = _element0_29
			_nw197 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
			(_nw197).ArraySet1(_element0_29, 0)
			var _nativeLen0_29 = (_len0_29).Int()
			_ = _nativeLen0_29
			for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
				(_nw197).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
			}
		}
		_1275_v0 = _nw197
		var _1277_v1 bool
		_ = _1277_v1
		_1277_v1 = true
		var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
		_ = _index193
		(_1275_v0).ArraySet1(_1277_v1, (_index193).Int())
		var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
		_ = _index194
		(_1275_v0).ArraySet1(true, (_index194).Int())
		if _1277_v1 {
			var _1278_v2 _dafny.Set
			_ = _1278_v2
			_1278_v2 = _dafny.SetOf((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
			var _1279_v3 _dafny.Array
			_ = _1279_v3
			var _nwElement0_45 _dafny.Set = _1278_v2
			_ = _nwElement0_45
			var _nw198 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(14))
			_ = _nw198
			(_nw198).ArraySet1(_nwElement0_45, 0)
			(_nw198).ArraySet1(_1278_v2, 1)
			(_nw198).ArraySet1(_dafny.SetOf((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), true, _1277_v1), 2)
			(_nw198).ArraySet1(_1278_v2, 3)
			(_nw198).ArraySet1(_1278_v2, 4)
			(_nw198).ArraySet1(_1278_v2, 5)
			(_nw198).ArraySet1(_1278_v2, 6)
			(_nw198).ArraySet1(_1278_v2, 7)
			(_nw198).ArraySet1(_1278_v2, 8)
			(_nw198).ArraySet1(_dafny.SetOf(_1277_v1), 9)
			(_nw198).ArraySet1(_1278_v2, 10)
			(_nw198).ArraySet1(_dafny.SetOf((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)), 11)
			(_nw198).ArraySet1(_1278_v2, 12)
			(_nw198).ArraySet1(_1278_v2, 13)
			_1279_v3 = _nw198
			var _1280_v4 D0
			_ = _1280_v4
			_1280_v4 = Companion_D0_.Create_DC0_(_1279_v3, p0)
			var _1281_v5 _dafny.Map
			_ = _1281_v5
			_1281_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), _1280_v4)
			_1281_v5 = (_1281_v5).Update(_1277_v1, Companion_D0_.Create_DC0_(_1279_v3, Companion_Default___.Fm20(!((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)), globalState)))
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
			_ = _index195
			(_1275_v0).ArraySet1((Companion_Default___.Fm20(_1277_v1, globalState)).Cmp((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), _1277_v1)).Cardinality())) != 0, (_index195).Int())
			var _1282_v6 _dafny.Map
			_ = _1282_v6
			_1282_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
			var _1283_v7 _dafny.Map
			_ = _1283_v7
			_1283_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1282_v6)
			_1283_v7 = (_1283_v7).Update((true) == (!((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))), (_1282_v6).Merge(Companion_Default___.Fm21(globalState)))
			var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
			_ = _index196
			(_1275_v0).ArraySet1(_1277_v1, (_index196).Int())
			var _1284_v8 _dafny.Array
			_ = _1284_v8
			var _nw199 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw199
			_1284_v8 = _nw199
			var _1285_v9 _dafny.MultiSet
			_ = _1285_v9
			_1285_v9 = _dafny.MultiSetOf(_1284_v8, _1284_v8)
			var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
			_ = _index197
			(_1275_v0).ArraySet1((func() bool {
				if !((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)) {
					return (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)
				}
				return !((p0).Cmp((_1285_v9).Cardinality()) <= 0)
			})(), (_index197).Int())
		} else {
			var _1286_v10 _dafny.Array
			_ = _1286_v10
			var _nw200 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
			_ = _nw200
			_1286_v10 = _nw200
			var _len0_30 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_30
			var _nw201 _dafny.Array
			_ = _nw201
			if _len0_30.Cmp(_dafny.Zero) == 0 {
				_nw201 = _dafny.NewArray(_len0_30)
			} else {
				var _init30 func(_dafny.Int) _dafny.Sequence = (func(_1287_p0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_1288_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_1287_p0, _1287_p0)
					}
				})(p0)
				_ = _init30
				var _element0_30 = _init30(_dafny.Zero)
				_ = _element0_30
				_nw201 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
				(_nw201).ArraySet1(_element0_30, 0)
				var _nativeLen0_30 = (_len0_30).Int()
				_ = _nativeLen0_30
				for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
					(_nw201).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
				}
			}
			_1286_v10 = _nw201
			var _1289_v11 _dafny.Array
			_ = _1289_v11
			var _nw202 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw202
			_1289_v11 = _nw202
			var _1290_v12 D2
			_ = _1290_v12
			_1290_v12 = Companion_D2_.Create_DC4_(p0, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), _1289_v11)
			var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
			_ = _index198
			var _rhs133 bool = _1277_v1
			_ = _rhs133
			var _rhs134 bool = (_1290_v12).Dtor_cf6()
			_ = _rhs134
			var _lhs72 _dafny.Array = _1275_v0
			_ = _lhs72
			var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
			_ = _lhs73
			(_lhs72).ArraySet1(_rhs133, (_lhs73).Int())
			_1277_v1 = _rhs134
			var _1291_v13 _dafny.MultiSet
			_ = _1291_v13
			_1291_v13 = _dafny.MultiSetOf(_dafny.CodePoint('t'))
			var _1292_v14 _dafny.Map
			_ = _1292_v14
			_1292_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1291_v13, p0)
			_1292_v14 = (_1292_v14).Update((_1291_v13).Difference(_1291_v13), _dafny.IntOfInt64(474))
			var _1293_v15 *C0
			_ = _1293_v15
			var _nw203 *C0 = New_C0_()
			_ = _nw203
			_nw203.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("o"))
			_1293_v15 = _nw203
			var _1294_v16 _dafny.Set
			_ = _1294_v16
			_1294_v16 = _dafny.SetOf((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
			var _1295_v17 _dafny.Map
			_ = _1295_v17
			_1295_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1275_v0, _1294_v16)
			var _1296_v18 _dafny.Array
			_ = _1296_v18
			var _nwElement0_46 _dafny.Set = (_1294_v16).Difference(_1294_v16)
			_ = _nwElement0_46
			var _nw204 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(16))
			_ = _nw204
			(_nw204).ArraySet1(_nwElement0_46, 0)
			(_nw204).ArraySet1(_1294_v16, 1)
			(_nw204).ArraySet1((_dafny.SetOf(true)).Intersection(_1294_v16), 2)
			(_nw204).ArraySet1(_1294_v16, 3)
			(_nw204).ArraySet1(_1294_v16, 4)
			(_nw204).ArraySet1(_1294_v16, 5)
			(_nw204).ArraySet1(_1294_v16, 6)
			(_nw204).ArraySet1(_1294_v16, 7)
			(_nw204).ArraySet1(_1294_v16, 8)
			(_nw204).ArraySet1(_dafny.SetOf((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), true), 9)
			(_nw204).ArraySet1(_1294_v16, 10)
			(_nw204).ArraySet1((_1294_v16).Difference(_1294_v16), 11)
			(_nw204).ArraySet1(_1294_v16, 12)
			(_nw204).ArraySet1((func() _dafny.Set {
				if true {
					return _dafny.SetOf(true, _1277_v1, _1277_v1, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
				}
				return (func() _dafny.Set {
					if (_1295_v17).Contains(_1275_v0) {
						return (_1295_v17).Get(_1275_v0).(_dafny.Set)
					}
					return _1294_v16
				})()
			})(), 13)
			(_nw204).ArraySet1((_1294_v16).Intersection(_1294_v16), 14)
			(_nw204).ArraySet1(_1294_v16, 15)
			_1296_v18 = _nw204
			_1296_v18 = _1296_v18
		}
		var _1297_v19 _dafny.Array
		_ = _1297_v19
		var _len0_31 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_31
		var _nw205 _dafny.Array
		_ = _nw205
		if _len0_31.Cmp(_dafny.Zero) == 0 {
			_nw205 = _dafny.NewArray(_len0_31)
		} else {
			var _init31 func(_dafny.Int) _dafny.Set = (func(_1298_v1 bool, _1299_v0 _dafny.Array) func(_dafny.Int) _dafny.Set {
				return func(_1300_i2 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_1298_v1, _1298_v1, (_1299_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1299_v0), 0))).Int()).(bool), _1298_v1)
				}
			})(_1277_v1, _1275_v0)
			_ = _init31
			var _element0_31 = _init31(_dafny.Zero)
			_ = _element0_31
			_nw205 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
			(_nw205).ArraySet1(_element0_31, 0)
			var _nativeLen0_31 = (_len0_31).Int()
			_ = _nativeLen0_31
			for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
				(_nw205).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
			}
		}
		_1297_v19 = _nw205
		var _1301_v20 D0
		_ = _1301_v20
		_1301_v20 = Companion_D0_.Create_DC0_(_1297_v19, p0)
		var _1302_v21 _dafny.Sequence
		_ = _1302_v21
		_1302_v21 = _dafny.SeqOf(_1301_v20)
		var _1303_v22 D4
		_ = _1303_v22
		_1303_v22 = Companion_D4_.Create_DC8_(_1302_v21)
		if _dafny.Companion_Sequence_.IsPrefixOf((_1303_v22).Dtor_cf13(), _1302_v21) {
			var _1304_v23 _dafny.Map
			_ = _1304_v23
			_1304_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(689), _1277_v1)
			var _1305_v24 _dafny.Sequence
			_ = _1305_v24
			_1305_v24 = _dafny.UnicodeSeqOfUtf8Bytes("wuqe")
			var _1306_v25 _dafny.Sequence
			_ = _1306_v25
			_1306_v25 = _dafny.SeqOf(_dafny.IntOfUint32((_1305_v24).Cardinality()))
			var _1307_v26 _dafny.Map
			_ = _1307_v26
			_1307_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1277_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1306_v25).Cardinality()), p0)).Cardinality())
			var _1308_v27 _dafny.MultiSet
			_ = _1308_v27
			_1308_v27 = _dafny.MultiSetOf(_1277_v1)
			var _1309_v28 _dafny.MultiSet
			_ = _1309_v28
			_1309_v28 = _dafny.MultiSetOf(p0)
			var _1310_v29 _dafny.Map
			_ = _1310_v29
			_1310_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), _dafny.IntOfInt64(159))
			var _1311_v30 _dafny.Sequence
			_ = _1311_v30
			_1311_v30 = _dafny.SeqOf(_1310_v29)
			var _1312_v31 _dafny.MultiSet
			_ = _1312_v31
			_1312_v31 = _dafny.MultiSetOf(_1275_v0)
			var _1313_v32 _dafny.Array
			_ = _1313_v32
			var _nwElement0_47 _dafny.Int = p0
			_ = _nwElement0_47
			var _nw206 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(22))
			_ = _nw206
			(_nw206).ArraySet1(_nwElement0_47, 0)
			(_nw206).ArraySet1((_1304_v23).Cardinality(), 1)
			(_nw206).ArraySet1(p0, 2)
			(_nw206).ArraySet1((_1307_v26).Cardinality(), 3)
			(_nw206).ArraySet1(_dafny.IntOfUint32((_1305_v24).Cardinality()), 4)
			(_nw206).ArraySet1(p0, 5)
			(_nw206).ArraySet1(p0, 6)
			(_nw206).ArraySet1(_dafny.IntOfUint32((_1305_v24).Cardinality()), 7)
			(_nw206).ArraySet1((_1308_v27).Cardinality(), 8)
			(_nw206).ArraySet1(p0, 9)
			(_nw206).ArraySet1((_1306_v25).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1306_v25).Cardinality()))).Uint32()).(_dafny.Int), 10)
			(_nw206).ArraySet1(_dafny.IntOfInt64(-17), 11)
			(_nw206).ArraySet1((_1306_v25).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1306_v25).Cardinality()))).Uint32()).(_dafny.Int), 12)
			(_nw206).ArraySet1(p0, 13)
			(_nw206).ArraySet1((_1309_v28).Cardinality(), 14)
			(_nw206).ArraySet1(p0, 15)
			(_nw206).ArraySet1(((_1311_v30).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1311_v30).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), 16)
			(_nw206).ArraySet1((func() _dafny.Int {
				if (_1312_v31).Contains(_1275_v0) {
					return (_1312_v31).Multiplicity(_1275_v0)
				}
				return p0
			})(), 17)
			(_nw206).ArraySet1(p0, 18)
			(_nw206).ArraySet1(p0, 19)
			(_nw206).ArraySet1((Companion_Default___.Fm22(false, p0, (_1308_v27).Cardinality(), globalState)).Cardinality(), 20)
			(_nw206).ArraySet1(p0, 21)
			_1313_v32 = _nw206
			var _1314_v33 _dafny.Sequence
			_ = _1314_v33
			_1314_v33 = _dafny.SeqOf(_1313_v32, _1313_v32, _1313_v32, _1313_v32, _1313_v32)
			var _1315_v34 _dafny.Map
			_ = _1315_v34
			_1315_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1314_v33).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1314_v33).Cardinality()))).Uint32()).(_dafny.Array), p0)
			_1315_v34 = (_1315_v34).Update((func() _dafny.Array {
				if _1277_v1 {
					return _1313_v32
				}
				return _1313_v32
			})(), p0)
			var _1316_v35 *C0
			_ = _1316_v35
			var _nw207 *C0 = New_C0_()
			_ = _nw207
			_nw207.Ctor__(_1305_v24)
			_1316_v35 = _nw207
			var _1317_v36 _dafny.Array
			_ = _1317_v36
			var _nwElement0_48 _dafny.Sequence = _1314_v33
			_ = _nwElement0_48
			var _nw208 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(3))
			_ = _nw208
			(_nw208).ArraySet1(_nwElement0_48, 0)
			(_nw208).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1314_v33, _1314_v33), 1)
			(_nw208).ArraySet1(_1314_v33, 2)
			_1317_v36 = _nw208
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1317_v36), 0))
			_ = _index199
			(_1317_v36).ArraySet1(_dafny.SeqOf(_1313_v32), (_index199).Int())
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1317_v36), 0))
			_ = _index200
			(_1317_v36).ArraySet1(_dafny.SeqOf(_1313_v32), (_index200).Int())
			var _1318_v37 _dafny.Map
			_ = _1318_v37
			_1318_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1277_v1, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
			var _source15 D4 = Companion_D4_.Create_DC11_((_dafny.IntOfUint32((_1306_v25).Cardinality())).Plus(p0), (!(false)) || ((func() bool {
				if (_1318_v37).Contains(true) {
					return (_1318_v37).Get(true).(bool)
				}
				return false
			})()), _1277_v1, (p0).Cmp(_dafny.IntOfInt64(-703)) >= 0, _1277_v1)
			_ = _source15
			if _source15.Is_DC9() {
				var _1319___mcc_h0 _dafny.MultiSet = _source15.Get_().(D4_DC9).Cf14
				_ = _1319___mcc_h0
				var _1320___mcc_h1 _dafny.CodePoint = _source15.Get_().(D4_DC9).Cf15
				_ = _1320___mcc_h1
				var _1321___mcc_h2 _dafny.Int = _source15.Get_().(D4_DC9).Cf16
				_ = _1321___mcc_h2
				var _1322___mcc_h3 _dafny.Set = _source15.Get_().(D4_DC9).Cf17
				_ = _1322___mcc_h3
				var _1323___mcc_h4 _dafny.Int = _source15.Get_().(D4_DC9).Cf18
				_ = _1323___mcc_h4
				var _1324_cf18 _dafny.Int = _1323___mcc_h4
				_ = _1324_cf18
				var _1325_cf17 _dafny.Set = _1322___mcc_h3
				_ = _1325_cf17
				var _1326_cf16 _dafny.Int = _1321___mcc_h2
				_ = _1326_cf16
				var _1327_cf15 _dafny.CodePoint = _1320___mcc_h1
				_ = _1327_cf15
				var _1328_cf14 _dafny.MultiSet = _1319___mcc_h0
				_ = _1328_cf14
				_1324_cf18 = ((_1326_cf16).Plus((_dafny.Zero).Minus(p0))).Plus((func() _dafny.Int {
					if false {
						return _1326_cf16
					}
					return _1324_cf18
				})())
				var _1329_v38 _dafny.Array
				_ = _1329_v38
				var _nw209 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
				_ = _nw209
				_1329_v38 = _nw209
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1329_v38), 0))
				_ = _index201
				(_1329_v38).ArraySet1(_1316_v35, (_index201).Int())
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1329_v38), 0))
				_ = _index202
				(_1329_v38).ArraySet1(_1316_v35, (_index202).Int())
				var _1330_v39 _dafny.Map
				_ = _1330_v39
				_1330_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), Companion_D0_.Create_DC1_(_dafny.IntOfInt64(486)))
				r1 = (_this).Fm5(_1326_cf16, _1277_v1, (_1330_v39).Merge(_1330_v39), (func() bool {
					if (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool) {
						return (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)
					}
					return !((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
				})(), globalState)
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
				_ = _index203
				(_1275_v0).ArraySet1((func() bool {
					if (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool) {
						return !((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
					}
					return (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)
				})(), (_index203).Int())
			} else if _source15.Is_DC10() {
				var _1331___mcc_h5 _dafny.Int = _source15.Get_().(D4_DC10).Cf19
				_ = _1331___mcc_h5
				var _1332___mcc_h6 _dafny.Int = _source15.Get_().(D4_DC10).Cf20
				_ = _1332___mcc_h6
				var _1333___mcc_h7 _dafny.Int = _source15.Get_().(D4_DC10).Cf21
				_ = _1333___mcc_h7
				var _1334_cf21 _dafny.Int = _1333___mcc_h7
				_ = _1334_cf21
				var _1335_cf20 _dafny.Int = _1332___mcc_h6
				_ = _1335_cf20
				var _1336_cf19 _dafny.Int = _1331___mcc_h5
				_ = _1336_cf19
				var _1337_v40 _dafny.CodePoint
				_ = _1337_v40
				_1337_v40 = _dafny.CodePoint('k')
				_1307_v26 = (_1307_v26).Update(!_dafny.Companion_Sequence_.Contains(_1316_v35.F8, _1337_v40), (_1336_cf19).Plus(_dafny.IntOfUint32((_1306_v25).Cardinality())))
				var _1338_v41 _dafny.Set
				_ = _1338_v41
				_1338_v41 = _dafny.SetOf(_1305_v24, _dafny.SeqOf(_1337_v40, _1337_v40))
				var _1339_v42 D2
				_ = _1339_v42
				_1339_v42 = Companion_D2_.Create_DC3_(_1338_v41)
				var _1340_v43 D2
				_ = _1340_v43
				_1340_v43 = Companion_D2_.Create_DC6_(_1339_v42)
				var _1341_v44 _dafny.Map
				_ = _1341_v44
				_1341_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1340_v43, _1316_v35.F8)
				_1341_v44 = _1341_v44
				var _1342_v45 D0
				_ = _1342_v45
				_1342_v45 = Companion_D0_.Create_DC1_((_1308_v27).Cardinality())
				var _1343_v46 _dafny.Map
				_ = _1343_v46
				_1343_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1337_v40, _1342_v45)
				_1334_cf21 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_1277_v1, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), (func() bool {
					if (_1318_v37).Contains((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)) {
						return (_1318_v37).Get((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)).(bool)
					}
					return _1277_v1
				})())).Cardinality()), _1335_cf20), (func() _dafny.Int {
					if (_1307_v26).Contains((_this).Fm5(_1334_cf21, _1277_v1, _1343_v46, _1277_v1, globalState)) {
						return (_1307_v26).Get((_this).Fm5(_1334_cf21, _1277_v1, _1343_v46, _1277_v1, globalState)).(_dafny.Int)
					}
					return (_dafny.SetOf((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))).Cardinality()
				})())
				_1307_v26 = ((_1307_v26).Merge(_1307_v26)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), _dafny.IntOfInt64(-533))).Merge(_1307_v26))
			} else if _source15.Is_DC11() {
				var _1344___mcc_h8 _dafny.Int = _source15.Get_().(D4_DC11).Cf22
				_ = _1344___mcc_h8
				var _1345___mcc_h9 bool = _source15.Get_().(D4_DC11).Cf23
				_ = _1345___mcc_h9
				var _1346___mcc_h10 bool = _source15.Get_().(D4_DC11).Cf24
				_ = _1346___mcc_h10
				var _1347___mcc_h11 bool = _source15.Get_().(D4_DC11).Cf25
				_ = _1347___mcc_h11
				var _1348___mcc_h12 bool = _source15.Get_().(D4_DC11).Cf26
				_ = _1348___mcc_h12
				var _1349_cf26 bool = _1348___mcc_h12
				_ = _1349_cf26
				var _1350_cf25 bool = _1347___mcc_h11
				_ = _1350_cf25
				var _1351_cf24 bool = _1346___mcc_h10
				_ = _1351_cf24
				var _1352_cf23 bool = _1345___mcc_h9
				_ = _1352_cf23
				var _1353_cf22 _dafny.Int = _1344___mcc_h8
				_ = _1353_cf22
				var _1354_v47 _dafny.Sequence
				_ = _1354_v47
				_1354_v47 = _dafny.SeqOf(_1351_cf24)
				var _1355_v48 _dafny.Sequence
				_ = _1355_v48
				_1355_v48 = _dafny.SeqOf(_1354_v47, Companion_Default___.Fm23(_1277_v1, globalState), _dafny.Companion_Sequence_.Update(_1354_v47, (Companion_Default___.SafeIndex(_1353_cf22, _dafny.IntOfUint32((_1354_v47).Cardinality()))).Uint32(), _1350_cf25))
				_1355_v48 = _1355_v48
				var _1356_v49 _dafny.Sequence
				_ = _1356_v49
				_1356_v49 = _dafny.SeqOf(_1275_v0)
				var _1357_v50 bool
				_ = _1357_v50
				var _1358_v51 _dafny.Int
				_ = _1358_v51
				var _out37 bool
				_ = _out37
				var _out38 _dafny.Int
				_ = _out38
				_out37, _out38 = (_this).M10(_1353_cf22, _1353_cf22, !(!(_1277_v1) || ((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))), (_1356_v49).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1356_v49).Cardinality()))).Uint32()).(_dafny.Array), globalState)
				_1357_v50 = _out37
				_1358_v51 = _out38
				_1353_cf22 = _1353_cf22
				var _1359_v52 _dafny.Array
				_ = _1359_v52
				var _nw210 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
				_ = _nw210
				_1359_v52 = _nw210
				_1359_v52 = _1359_v52
			} else if _source15.Is_DC8() {
				var _1360___mcc_h13 _dafny.Sequence = _source15.Get_().(D4_DC8).Cf13
				_ = _1360___mcc_h13
				var _1361_cf13 _dafny.Sequence = _1360___mcc_h13
				_ = _1361_cf13
				var _1362_v53 _dafny.Int
				_ = _1362_v53
				_1362_v53 = _dafny.IntOfInt64(-892)
				_1362_v53 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(266))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}(func(_1363_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('m')
				}))).Cardinality())))
				var _1364_v54 _dafny.CodePoint
				_ = _1364_v54
				_1364_v54 = _dafny.CodePoint('f')
				_1364_v54 = (func() _dafny.CodePoint {
					if _1277_v1 {
						return _1364_v54
					}
					return _1364_v54
				})()
				var _1365_v55 D2
				_ = _1365_v55
				_1365_v55 = Companion_D2_.Create_DC6_(Companion_D2_.Create_DC5_((_dafny.Zero).Minus(_1362_v53), _1306_v25, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)))
				var _1366_v56 _dafny.Sequence
				_ = _1366_v56
				_1366_v56 = _dafny.SeqOf(_1365_v55)
				var _1367_v57 bool
				_ = _1367_v57
				_1367_v57 = _dafny.Companion_Sequence_.Contains(_1366_v56, _1365_v55)
				var _1368_v58 _dafny.Map
				_ = _1368_v58
				_1368_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1364_v54, _1275_v0)
				var _1369_v59 _dafny.Set
				_ = _1369_v59
				_1369_v59 = _dafny.SetOf(_1316_v35, _1316_v35)
				var _1370_v60 D4
				_ = _1370_v60
				_1370_v60 = Companion_D4_.Create_DC9_((Companion_Default___.Fm24((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), globalState)).Update((_dafny.Zero).Minus((_dafny.Zero).Minus(_1362_v53)), Companion_Default___.Abs(_1362_v53)), _dafny.CodePoint('d'), ((_dafny.SetOf(_1277_v1)).Cardinality()).Plus((_dafny.Zero).Minus((_1304_v23).Cardinality())), _1369_v59, (_1362_v53).Minus(_1362_v53))
				var _1371_v61 D0
				_ = _1371_v61
				_1371_v61 = Companion_D0_.Create_DC1_(p0)
				var _1372_v62 _dafny.Map
				_ = _1372_v62
				_1372_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1364_v54, _1371_v61)
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
				_ = _index204
				var _rhs135 bool = _1367_v57
				_ = _rhs135
				var _rhs136 bool = (true) || ((_this).Fm5(_1362_v53, true, _1372_v62, !(!((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))), globalState))
				_ = _rhs136
				var _rhs137 _dafny.Map = _1368_v58
				_ = _rhs137
				var _rhs138 D4 = _1370_v60
				_ = _rhs138
				var _lhs74 _dafny.Array = _1275_v0
				_ = _lhs74
				var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
				_ = _lhs75
				_1367_v57 = _rhs135
				(_lhs74).ArraySet1(_rhs136, (_lhs75).Int())
				_1368_v58 = _rhs137
				_1370_v60 = _rhs138
				var _1373_v63 _dafny.Array
				_ = _1373_v63
				var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
				_ = _nw211
				_1373_v63 = _nw211
				var _1374_v64 _dafny.Array
				_ = _1374_v64
				var _nwElement0_49 _dafny.Array = _1313_v32
				_ = _nwElement0_49
				var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(17))
				_ = _nw212
				(_nw212).ArraySet1(_nwElement0_49, 0)
				(_nw212).ArraySet1(_1313_v32, 1)
				(_nw212).ArraySet1(_1313_v32, 2)
				(_nw212).ArraySet1(_1313_v32, 3)
				(_nw212).ArraySet1(_1313_v32, 4)
				(_nw212).ArraySet1(_1313_v32, 5)
				(_nw212).ArraySet1(_1313_v32, 6)
				(_nw212).ArraySet1(_1313_v32, 7)
				(_nw212).ArraySet1(_1313_v32, 8)
				(_nw212).ArraySet1(_1313_v32, 9)
				(_nw212).ArraySet1(_1313_v32, 10)
				(_nw212).ArraySet1(_1313_v32, 11)
				(_nw212).ArraySet1(_1313_v32, 12)
				(_nw212).ArraySet1(_1313_v32, 13)
				(_nw212).ArraySet1(_1313_v32, 14)
				(_nw212).ArraySet1(_1313_v32, 15)
				(_nw212).ArraySet1(_1313_v32, 16)
				_1374_v64 = _nw212
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(432), _dafny.ArrayLen((_1373_v63), 0))
				_ = _index205
				(_1373_v63).ArraySet1(_1374_v64, (_index205).Int())
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(432), _dafny.ArrayLen((_1373_v63), 0))
				_ = _index206
				(_1373_v63).ArraySet1(_1374_v64, (_index206).Int())
			} else {
				var _1375___mcc_h14 D4 = _source15.Get_().(D4_DC12).Cf27
				_ = _1375___mcc_h14
				var _1376_cf27 D4 = _1375___mcc_h14
				_ = _1376_cf27
				_1277_v1 = (func() bool {
					if _1277_v1 {
						return ((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)) || ((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
					}
					return (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)
				})()
				_1313_v32 = _1313_v32
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
				_ = _index207
				(_1275_v0).ArraySet1(_1277_v1, (_index207).Int())
				var _1377_v65 _dafny.Int
				_ = _1377_v65
				_1377_v65 = _dafny.IntOfInt64(226)
				var _1378_v66 _dafny.Array
				_ = _1378_v66
				var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
				_ = _nw213
				_1378_v66 = _nw213
				var _1379_v67 _dafny.Map
				_ = _1379_v67
				_1379_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1277_v1, _1275_v0)
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_1378_v66), 0))
				_ = _index208
				(_1378_v66).ArraySet1(_1379_v67, (_index208).Int())
				var _1380_v68 _dafny.Array
				_ = _1380_v68
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_32
				var _nw214 _dafny.Array
				_ = _nw214
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw214 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) _dafny.Map = (func(_1381_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
						return func(_1382_i4 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), _1381_p0)
						}
					})(p0)
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw214 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw214).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw214).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1380_v68 = _nw214
				var _1383_v69 _dafny.Map
				_ = _1383_v69
				_1383_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1377_v65, _1380_v68)
				var _1384_v70 _dafny.Array
				_ = _1384_v70
				_1384_v70 = _1380_v68
				var _1385_v71 _dafny.Array
				_ = _1385_v71
				var _nwElement0_50 _dafny.Array = _1380_v68
				_ = _nwElement0_50
				var _nw215 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(26))
				_ = _nw215
				(_nw215).ArraySet1(_nwElement0_50, 0)
				(_nw215).ArraySet1(_1380_v68, 1)
				(_nw215).ArraySet1(_1380_v68, 2)
				(_nw215).ArraySet1(_1380_v68, 3)
				(_nw215).ArraySet1(_1380_v68, 4)
				(_nw215).ArraySet1(_1380_v68, 5)
				(_nw215).ArraySet1(_1380_v68, 6)
				(_nw215).ArraySet1(_1380_v68, 7)
				(_nw215).ArraySet1((func() _dafny.Array {
					if (_1383_v69).Contains(_dafny.IntOfUint32((_1306_v25).Cardinality())) {
						return (_1383_v69).Get(_dafny.IntOfUint32((_1306_v25).Cardinality())).(_dafny.Array)
					}
					return _1380_v68
				})(), 8)
				(_nw215).ArraySet1((_1384_v70), 9)
				(_nw215).ArraySet1(_1380_v68, 10)
				(_nw215).ArraySet1(_1380_v68, 11)
				(_nw215).ArraySet1(_1380_v68, 12)
				(_nw215).ArraySet1(_1380_v68, 13)
				(_nw215).ArraySet1(_1380_v68, 14)
				(_nw215).ArraySet1(_1380_v68, 15)
				(_nw215).ArraySet1(_1380_v68, 16)
				(_nw215).ArraySet1(_1380_v68, 17)
				(_nw215).ArraySet1(_1380_v68, 18)
				(_nw215).ArraySet1(_1380_v68, 19)
				(_nw215).ArraySet1(_1380_v68, 20)
				(_nw215).ArraySet1(_1380_v68, 21)
				(_nw215).ArraySet1(_1380_v68, 22)
				(_nw215).ArraySet1(_1380_v68, 23)
				(_nw215).ArraySet1(_1380_v68, 24)
				(_nw215).ArraySet1(_1380_v68, 25)
				_1385_v71 = _nw215
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1385_v71), 0))
				_ = _index209
				(_1385_v71).ArraySet1(_1380_v68, (_index209).Int())
				var _1386_v72 _dafny.CodePoint
				_ = _1386_v72
				_1386_v72 = _dafny.CodePoint('p')
				var _1387_v73 _dafny.Set
				_ = _1387_v73
				_1387_v73 = _dafny.SetOf(_1277_v1, _1277_v1, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
				var _1388_v74 _dafny.Sequence
				_ = _1388_v74
				_1388_v74 = _dafny.SeqOf(_1387_v73)
				var _1389_v75 _dafny.Sequence
				_ = _1389_v75
				_1389_v75 = _dafny.SeqOf(_1379_v67)
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_1378_v66), 0))
				_ = _index210
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1385_v71), 0))
				_ = _index211
				var _rhs139 bool = !(false)
				_ = _rhs139
				var _rhs140 _dafny.Int = _dafny.IntOfUint32((_1388_v74).Cardinality())
				_ = _rhs140
				var _rhs141 _dafny.Map = (_1379_v67).Merge((_1389_v75).Select((Companion_Default___.SafeIndex(_1377_v65, _dafny.IntOfUint32((_1389_v75).Cardinality()))).Uint32()).(_dafny.Map))
				_ = _rhs141
				var _rhs142 _dafny.Array = _1380_v68
				_ = _rhs142
				var _rhs143 _dafny.CodePoint = _dafny.CodePoint('d')
				_ = _rhs143
				var _lhs76 _dafny.Array = _1378_v66
				_ = _lhs76
				var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_1378_v66), 0))
				_ = _lhs77
				var _lhs78 _dafny.Array = _1385_v71
				_ = _lhs78
				var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1385_v71), 0))
				_ = _lhs79
				r0 = _rhs139
				_1377_v65 = _rhs140
				(_lhs76).ArraySet1(_rhs141, (_lhs77).Int())
				(_lhs78).ArraySet1(_rhs142, (_lhs79).Int())
				_1386_v72 = _rhs143
			}
			var _1390_v76 _dafny.Array
			_ = _1390_v76
			var _1391_v77 _dafny.Array
			_ = _1391_v77
			var _1392_v78 bool
			_ = _1392_v78
			var _out39 _dafny.Array
			_ = _out39
			var _out40 _dafny.Array
			_ = _out40
			var _out41 bool
			_ = _out41
			_out39, _out40, _out41 = (_this).M9(globalState)
			_1390_v76 = _out39
			_1391_v77 = _out40
			_1392_v78 = _out41
		} else {
			var _1393_v79 _dafny.Array
			_ = _1393_v79
			var _1394_v80 _dafny.Array
			_ = _1394_v80
			var _1395_v81 bool
			_ = _1395_v81
			var _out42 _dafny.Array
			_ = _out42
			var _out43 _dafny.Array
			_ = _out43
			var _out44 bool
			_ = _out44
			_out42, _out43, _out44 = (_this).M9(globalState)
			_1393_v79 = _out42
			_1394_v80 = _out43
			_1395_v81 = _out44
			var _1396_v83 _dafny.Set
			_ = _1396_v83
			_1396_v83 = _dafny.SetOf(_dafny.IntOfInt64(-990), p0)
			var _1397_v84 _dafny.Set
			_ = _1397_v84
			_1397_v84 = _dafny.SetOf((_1396_v83).Cardinality(), p0)
			var _1398_v85 _dafny.MultiSet
			_ = _1398_v85
			_1398_v85 = _dafny.MultiSetOf(p0, (_1396_v83).Cardinality())
			var _1399_v86 _dafny.CodePoint
			_ = _1399_v86
			_1399_v86 = _dafny.CodePoint('a')
			var _1400_v87 _dafny.Map
			_ = _1400_v87
			_1400_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1398_v85).Cardinality(), _1399_v86)
			var _1401_v88 _dafny.Sequence
			_ = _1401_v88
			_1401_v88 = _dafny.SeqOf((func() _dafny.Map {
				var _coll52 = _dafny.NewMapBuilder()
				_ = _coll52
				for _iter54 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-592), _dafny.IntOfInt64(863))); ; {
					_compr_52, _ok54 := _iter54()
					if !_ok54 {
						break
					}
					var _1402_v82 _dafny.Int
					_1402_v82 = interface{}(_compr_52).(_dafny.Int)
					if ((_dafny.IntOfInt64(-592)).Cmp(_1402_v82) <= 0) && ((_1402_v82).Cmp(_dafny.IntOfInt64(863)) < 0) {
						_coll52.Add((_1402_v82).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dmetdhseu")).Cardinality())), _dafny.CodePoint('y'))
					}
				}
				return _coll52.ToMap()
			}()).Merge(_1400_v87))
			_1401_v88 = _1401_v88
			var _1403_v89 _dafny.Int
			_ = _1403_v89
			_1403_v89 = _dafny.IntOfInt64(768)
			var _1404_v90 _dafny.Sequence
			_ = _1404_v90
			_1404_v90 = _dafny.SeqOf(!(_1277_v1))
			var _1405_v91 _dafny.Sequence
			_ = _1405_v91
			_1405_v91 = _dafny.SeqOf(_1404_v90)
			var _1406_v92 _dafny.Sequence
			_ = _1406_v92
			_1406_v92 = _dafny.SeqOf(_dafny.IntOfInt64(732), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1404_v90, (_1405_v91).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1405_v91).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()), p0, p0, _1403_v89)
			var _1407_v93 _dafny.Map
			_ = _1407_v93
			_1407_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1395_v81, p0)
			var _1408_v94 _dafny.Sequence
			_ = _1408_v94
			_1408_v94 = _dafny.UnicodeSeqOfUtf8Bytes("wvfkivcdy")
			var _rhs144 _dafny.Int = _dafny.IntOfInt64(665)
			_ = _rhs144
			var _rhs145 _dafny.Int = (func() _dafny.Int {
				if (_1407_v93).Contains(_1277_v1) {
					return (_1407_v93).Get(_1277_v1).(_dafny.Int)
				}
				return _1403_v89
			})()
			_ = _rhs145
			var _rhs146 bool = !(_dafny.Companion_Sequence_.IsPrefixOf(_1408_v94, _1408_v94)) || ((_dafny.IntOfInt64(890)).Cmp((func() _dafny.Int {
				if (_1398_v85).Contains(p0) {
					return (_1398_v85).Multiplicity(p0)
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)).Cardinality()
			})()) == 0)
			_ = _rhs146
			var _rhs147 _dafny.Sequence = _1406_v92
			_ = _rhs147
			_1403_v89 = _rhs144
			_1403_v89 = _rhs145
			_1395_v81 = _rhs146
			_1406_v92 = _rhs147
			var _1409_v95 _dafny.Sequence
			_ = _1409_v95
			_1409_v95 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(967))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg84 _dafny.Int) interface{} {
					return coer84(arg84)
				}
			}((func(_1410_v86 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1411_i5 _dafny.Int) _dafny.CodePoint {
					return _1410_v86
				}
			})(_1399_v86))), _dafny.Companion_Sequence_.Update((Companion_D6_.Create_DC14_(_dafny.UnicodeSeqOfUtf8Bytes("erond"))).Dtor_cf29(), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((Companion_D6_.Create_DC14_(_dafny.UnicodeSeqOfUtf8Bytes("erond"))).Dtor_cf29()).Cardinality()))).Uint32(), _dafny.CodePoint('f')), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(523))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg85 _dafny.Int) interface{} {
					return coer85(arg85)
				}
			}((func(_1412_v86 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1413_i6 _dafny.Int) _dafny.CodePoint {
					return _1412_v86
				}
			})(_1399_v86))))
			var _1414_v96 *C0
			_ = _1414_v96
			var _nw216 *C0 = New_C0_()
			_ = _nw216
			_nw216.Ctor__(_dafny.Companion_Sequence_.Concatenate(_1408_v94, (_1409_v95).Select((Companion_Default___.SafeIndex(_1403_v89, _dafny.IntOfUint32((_1409_v95).Cardinality()))).Uint32()).(_dafny.Sequence)))
			_1414_v96 = _nw216
			var _1415_v97 _dafny.Map
			_ = _1415_v97
			_1415_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), _1395_v81)
			var _1416_v98 D2
			_ = _1416_v98
			_1416_v98 = Companion_D2_.Create_DC5_((_1415_v97).Cardinality(), _1406_v92, _1395_v81)
			_1277_v1 = (func() bool {
				if _1395_v81 {
					return (_1416_v98).Dtor_cf10()
				}
				return !((_1404_v90).Select((Companion_Default___.SafeIndex(_1403_v89, _dafny.IntOfUint32((_1404_v90).Cardinality()))).Uint32()).(bool))
			})()
		}
		var _1417_v99 _dafny.Set
		_ = _1417_v99
		_1417_v99 = _dafny.SetOf(p0, p0, p0)
		var _1418_v100 _dafny.Map
		_ = _1418_v100
		_1418_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), (_1417_v99).Cardinality())
		var _1419_v101 _dafny.CodePoint
		_ = _1419_v101
		_1419_v101 = _dafny.CodePoint('h')
		var _1420_v102 D0
		_ = _1420_v102
		_1420_v102 = Companion_D0_.Create_DC1_(p0)
		var _1421_v103 _dafny.Map
		_ = _1421_v103
		_1421_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1419_v101, _1420_v102)
		if !((_this).Fm5((_1418_v100).Cardinality(), _1277_v1, _1421_v103, _1277_v1, globalState)) || ((p0).Cmp(p0) != 0) {
			var _1422_v104 _dafny.Array
			_ = _1422_v104
			var _nw217 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw217
			_1422_v104 = _nw217
			var _1423_v105 _dafny.Sequence
			_ = _1423_v105
			_1423_v105 = _dafny.UnicodeSeqOfUtf8Bytes("awvnycoec")
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))
			_ = _index212
			(_1422_v104).ArraySet1(_dafny.IntOfUint32((_1423_v105).Cardinality()), (_index212).Int())
			var _1424_v106 _dafny.Int
			_ = _1424_v106
			_1424_v106 = _dafny.IntOfInt64(125)
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_1422_v104), 0))
			_ = _index213
			(_1422_v104).ArraySet1(p0, (_index213).Int())
			var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))
			_ = _index214
			var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_1422_v104), 0))
			_ = _index215
			var _rhs148 _dafny.Int = _1424_v106
			_ = _rhs148
			var _rhs149 _dafny.Int = (_1424_v106).Minus(_1424_v106)
			_ = _rhs149
			var _rhs150 _dafny.Int = p0
			_ = _rhs150
			var _lhs80 _dafny.Array = _1422_v104
			_ = _lhs80
			var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))
			_ = _lhs81
			var _lhs82 _dafny.Array = _1422_v104
			_ = _lhs82
			var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_1422_v104), 0))
			_ = _lhs83
			(_lhs80).ArraySet1(_rhs148, (_lhs81).Int())
			_1424_v106 = _rhs149
			(_lhs82).ArraySet1(_rhs150, (_lhs83).Int())
			var _1425_v107 _dafny.MultiSet
			_ = _1425_v107
			_1425_v107 = _dafny.MultiSetOf(_1275_v0, _1275_v0)
			var _1426_v108 D4
			_ = _1426_v108
			_1426_v108 = Companion_D4_.Create_DC10_(p0, _1424_v106, _1424_v106)
			var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))
			_ = _index216
			var _rhs151 _dafny.Int = Companion_Default___.SafeModuloInt((((_dafny.MultiSetOf(_1275_v0)).Update(_1275_v0, Companion_Default___.Abs(p0))).Union(_1425_v107)).Cardinality(), _1424_v106)
			_ = _rhs151
			var _rhs152 _dafny.Sequence = Companion_Default___.Fm25(_1426_v108, _1423_v105, globalState)
			_ = _rhs152
			var _lhs84 _dafny.Array = _1422_v104
			_ = _lhs84
			var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))
			_ = _lhs85
			(_lhs84).ArraySet1(_rhs151, (_lhs85).Int())
			_1423_v105 = _rhs152
			var _1427_v109 _dafny.Map
			_ = _1427_v109
			_1427_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1422_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))).Int()).(_dafny.Int))
			var _1428_v110 _dafny.Array
			_ = _1428_v110
			var _nw218 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
			_ = _nw218
			_1428_v110 = _nw218
			var _1429_v111 D2
			_ = _1429_v111
			_1429_v111 = Companion_D2_.Create_DC4_((func() _dafny.Int {
				if (_1427_v109).Contains(p0) {
					return (_1427_v109).Get(p0).(_dafny.Int)
				}
				return _dafny.IntOfInt64(-269)
			})(), false, _1428_v110)
			if (Companion_Default___.Fm24((_1429_v111).Dtor_cf6(), globalState)).IsSubsetOf(_dafny.MultiSetOf(p0, p0)) {
				var _1430_v112 _dafny.Sequence
				_ = _1430_v112
				_1430_v112 = _dafny.SeqOf((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
				var _1431_v113 D7
				_ = _1431_v113
				_1431_v113 = Companion_D7_.Create_DC17_(_1430_v112)
				var _1432_v114 _dafny.Array
				_ = _1432_v114
				var _nwElement0_51 _dafny.Sequence = _dafny.SeqOf((_1429_v111).Dtor_cf6(), false)
				_ = _nwElement0_51
				var _nw219 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(20))
				_ = _nw219
				(_nw219).ArraySet1(_nwElement0_51, 0)
				(_nw219).ArraySet1(_1430_v112, 1)
				(_nw219).ArraySet1(_1430_v112, 2)
				(_nw219).ArraySet1(_1430_v112, 3)
				(_nw219).ArraySet1(_1430_v112, 4)
				(_nw219).ArraySet1(_dafny.SeqOf((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)), 5)
				(_nw219).ArraySet1(_1430_v112, 6)
				(_nw219).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), _1277_v1, _1277_v1, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)), (_1431_v113).Dtor_cf34()), 7)
				(_nw219).ArraySet1(Companion_Default___.Fm23((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), globalState), 8)
				(_nw219).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1430_v112, _dafny.SeqOf(_1277_v1)), 9)
				(_nw219).ArraySet1(_1430_v112, 10)
				(_nw219).ArraySet1(_1430_v112, 11)
				(_nw219).ArraySet1(_1430_v112, 12)
				(_nw219).ArraySet1(_1430_v112, 13)
				(_nw219).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1430_v112, _1430_v112), 14)
				(_nw219).ArraySet1(_dafny.SeqOf(true, _1277_v1, false), 15)
				(_nw219).ArraySet1(_dafny.Companion_Sequence_.Update(_1430_v112, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.IntOfUint32((_1430_v112).Cardinality()))).Uint32(), (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)), 16)
				(_nw219).ArraySet1(_dafny.SeqOf(!(!((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))), (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)), 17)
				(_nw219).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1277_v1, _1277_v1, _1277_v1), _1430_v112), 18)
				(_nw219).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1430_v112, _1430_v112), 19)
				_1432_v114 = _nw219
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_1432_v114), 0))
				_ = _index217
				(_1432_v114).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1430_v112, _1430_v112), (_index217).Int())
				var _1433_v115 _dafny.Sequence
				_ = _1433_v115
				_1433_v115 = _dafny.SeqOf((_1422_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1430_v112).Cardinality()), _1424_v106)
				var _1434_v116 _dafny.Map
				_ = _1434_v116
				_1434_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1433_v115, (_1422_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))).Int()).(_dafny.Int))
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
				_ = _index218
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))
				_ = _index219
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_1432_v114), 0))
				_ = _index220
				var _rhs153 bool = !((_1424_v106).Cmp(Companion_Default___.Fm20(_1277_v1, globalState)) == 0)
				_ = _rhs153
				var _rhs154 _dafny.Int = (func() _dafny.Int {
					if (_1434_v116).Contains(_1433_v115) {
						return (_1434_v116).Get(_1433_v115).(_dafny.Int)
					}
					return (_1422_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))).Int()).(_dafny.Int)
				})()
				_ = _rhs154
				var _rhs155 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1277_v1, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)), _dafny.SeqOf(false))
				_ = _rhs155
				var _lhs86 _dafny.Array = _1275_v0
				_ = _lhs86
				var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
				_ = _lhs87
				var _lhs88 _dafny.Array = _1422_v104
				_ = _lhs88
				var _lhs89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))
				_ = _lhs89
				var _lhs90 _dafny.Array = _1432_v114
				_ = _lhs90
				var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_1432_v114), 0))
				_ = _lhs91
				(_lhs86).ArraySet1(_rhs153, (_lhs87).Int())
				(_lhs88).ArraySet1(_rhs154, (_lhs89).Int())
				(_lhs90).ArraySet1(_rhs155, (_lhs91).Int())
				var _1435_v117 *C0
				_ = _1435_v117
				var _nw220 *C0 = New_C0_()
				_ = _nw220
				_nw220.Ctor__(_1423_v105)
				_1435_v117 = _nw220
				var _1436_v118 _dafny.Map
				_ = _1436_v118
				_1436_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1435_v117, !((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)))
				_1436_v118 = _1436_v118
				var _1437_v119 *C0
				_ = _1437_v119
				var _nw221 *C0 = New_C0_()
				_ = _nw221
				_nw221.Ctor__(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm25(_1426_v108, _1423_v105, globalState), _1423_v105))
				_1437_v119 = _nw221
				var _1438_v120 _dafny.Set
				_ = _1438_v120
				_1438_v120 = _dafny.SetOf(_1277_v1, !((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)))
				var _1439_v121 _dafny.Sequence
				_ = _1439_v121
				_1439_v121 = _dafny.SeqOf(_dafny.SeqOf((_1438_v120).Cardinality(), _dafny.IntOfInt64(767)))
				_1439_v121 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1439_v121, _1439_v121), (Companion_Default___.SafeIndex(_1424_v106, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1439_v121, _1439_v121)).Cardinality()))).Uint32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(790))).Uint32(), func(coer86 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg86 _dafny.Int) interface{} {
						return coer86(arg86)
					}
				}((func(_1440_v109 _dafny.Map) func(_dafny.Int) _dafny.Int {
					return func(_1441_i7 _dafny.Int) _dafny.Int {
						return (_1440_v109).Cardinality()
					}
				})(_1427_v109))))
				var _1442_v122 _dafny.Map
				_ = _1442_v122
				_1442_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1419_v101, Companion_Default___.Fm20((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), globalState))
				_1442_v122 = (_1442_v122).Update(_1419_v101, p0)
			} else {
				var _1443_v123 _dafny.Map
				_ = _1443_v123
				_1443_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1277_v1, _1275_v0)
				var _1444_v124 _dafny.Set
				_ = _1444_v124
				_1444_v124 = _dafny.SetOf(_1275_v0, (func() _dafny.Array {
					if (_1443_v123).Contains(!((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))) {
						return (_1443_v123).Get(!((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))).(_dafny.Array)
					}
					return _1275_v0
				})())
				var _1445_v125 _dafny.Array
				_ = _1445_v125
				var _nwElement0_52 _dafny.Set = _1444_v124
				_ = _nwElement0_52
				var _nw222 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(4))
				_ = _nw222
				(_nw222).ArraySet1(_nwElement0_52, 0)
				(_nw222).ArraySet1(_1444_v124, 1)
				(_nw222).ArraySet1(_dafny.SetOf(_1275_v0, _1275_v0), 2)
				(_nw222).ArraySet1((_1444_v124).Union(_dafny.SetOf(_1275_v0)), 3)
				_1445_v125 = _nw222
				var _1446_v126 _dafny.Set
				_ = _1446_v126
				_1446_v126 = _1444_v124
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_1445_v125), 0))
				_ = _index221
				(_1445_v125).ArraySet1((func() _dafny.Set {
					if _1277_v1 {
						return (_1446_v126)
					}
					return _1444_v124
				})(), (_index221).Int())
				var _1447_v127 _dafny.Map
				_ = _1447_v127
				_1447_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1444_v124).Union(_1444_v124))
				var _1448_v128 _dafny.Map
				_ = _1448_v128
				_1448_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1277_v1, _1277_v1)
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_1445_v125), 0))
				_ = _index222
				(_1445_v125).ArraySet1((func() _dafny.Set {
					if (_1447_v127).Contains((func() bool {
						if _1277_v1 {
							return (func() bool {
								if (_1448_v128).Contains(_1277_v1) {
									return (_1448_v128).Get(_1277_v1).(bool)
								}
								return false
							})()
						}
						return !((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
					})()) {
						return (_1447_v127).Get((func() bool {
							if _1277_v1 {
								return (func() bool {
									if (_1448_v128).Contains(_1277_v1) {
										return (_1448_v128).Get(_1277_v1).(bool)
									}
									return false
								})()
							}
							return !((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
						})()).(_dafny.Set)
					}
					return (_1444_v124).Difference(_1444_v124)
				})(), (_index222).Int())
				var _1449_v129 *C0
				_ = _1449_v129
				var _nw223 *C0 = New_C0_()
				_ = _nw223
				_nw223.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(445))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg87 _dafny.Int) interface{} {
						return coer87(arg87)
					}
				}((func(_1450_v101 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1451_i8 _dafny.Int) _dafny.CodePoint {
						return _1450_v101
					}
				})(_1419_v101))))
				_1449_v129 = _nw223
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))
				_ = _index223
				(_1422_v104).ArraySet1((_dafny.MultiSetOf(((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)) && ((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)), (_this).Fm5(_1424_v106, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), _1421_v103, false, globalState))).Cardinality(), (_index223).Int())
				var _1452_v130 _dafny.MultiSet
				_ = _1452_v130
				_1452_v130 = _dafny.MultiSetOf(_1424_v106, _1424_v106)
				var _1453_v131 D4
				_ = _1453_v131
				_1453_v131 = Companion_D4_.Create_DC9_(_1452_v130, _1419_v101, _1424_v106, _dafny.SetOf(_1449_v129), _1424_v106)
				var _1454_v132 *C0
				_ = _1454_v132
				var _nw224 *C0 = New_C0_()
				_ = _nw224
				_nw224.Ctor__(_dafny.Companion_Sequence_.Concatenate(_1423_v105, _dafny.Companion_Sequence_.Update(_1449_v129.F8, (Companion_Default___.SafeIndex((_1422_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1449_v129.F8).Cardinality()))).Uint32(), (_1453_v131).Dtor_cf15())))
				_1454_v132 = _nw224
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))
				_ = _index224
				(_1422_v104).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)), p0)).Cardinality(), (_index224).Int())
			}
			var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1422_v104), 0))
			_ = _index225
			(_1422_v104).ArraySet1(_1424_v106, (_index225).Int())
			_1424_v106 = p0
		} else {
			r0 = !(_1277_v1)
			_1277_v1 = (func() bool {
				if (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool) {
					return (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)
				}
				return !((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
			})()
			var _1455_v133 _dafny.Int
			_ = _1455_v133
			_1455_v133 = _dafny.IntOfInt64(58)
			_1455_v133 = p0
			var _1456_v134 _dafny.Sequence
			_ = _1456_v134
			_1456_v134 = _dafny.UnicodeSeqOfUtf8Bytes("gv")
			var _1457_v135 *C0
			_ = _1457_v135
			var _nw225 *C0 = New_C0_()
			_ = _nw225
			_nw225.Ctor__(_dafny.Companion_Sequence_.Concatenate(_1456_v134, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-868))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg88 _dafny.Int) interface{} {
					return coer88(arg88)
				}
			}((func(_1458_v101 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1459_i9 _dafny.Int) _dafny.CodePoint {
					return _1458_v101
				}
			})(_1419_v101)))))
			_1457_v135 = _nw225
			var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
			_ = _index226
			(_1275_v0).ArraySet1(_1277_v1, (_index226).Int())
		}
		var _1460_v136 _dafny.Sequence
		_ = _1460_v136
		_1460_v136 = _dafny.UnicodeSeqOfUtf8Bytes("ng")
		var _1461_v137 _dafny.Sequence
		_ = _1461_v137
		_1461_v137 = _dafny.SeqOf(_1277_v1, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
		var _1462_v138 _dafny.Sequence
		_ = _1462_v138
		_1462_v138 = _dafny.SeqOf(_1461_v137, _1461_v137)
		var _1463_v139 _dafny.Sequence
		_ = _1463_v139
		_1463_v139 = _dafny.SeqOf(_1461_v137, (_1462_v138).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1462_v138).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _1464_v140 D6
		_ = _1464_v140
		_1464_v140 = Companion_D6_.Create_DC15_(p0, _dafny.IntOfUint32((_1460_v136).Cardinality()), _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(Companion_Default___.Fm23(_1277_v1, globalState), _1461_v137), _1462_v138))
		var _source16 D6 = _1464_v140
		_ = _source16
		if _source16.Is_DC15() {
			var _1465___mcc_h15 _dafny.Int = _source16.Get_().(D6_DC15).Cf30
			_ = _1465___mcc_h15
			var _1466___mcc_h16 _dafny.Int = _source16.Get_().(D6_DC15).Cf31
			_ = _1466___mcc_h16
			var _1467___mcc_h17 bool = _source16.Get_().(D6_DC15).Cf32
			_ = _1467___mcc_h17
			var _1468_cf32 bool = _1467___mcc_h17
			_ = _1468_cf32
			var _1469_cf31 _dafny.Int = _1466___mcc_h16
			_ = _1469_cf31
			var _1470_cf30 _dafny.Int = _1465___mcc_h15
			_ = _1470_cf30
			_1420_v102 = Companion_D0_.Create_DC1_((_1469_cf31).Minus(_dafny.IntOfInt64(894)))
			_1461_v137 = _1461_v137
			var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
			_ = _index227
			(_1275_v0).ArraySet1(_1468_cf32, (_index227).Int())
			var _1471_v141 _dafny.Array
			_ = _1471_v141
			var _nw226 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw226
			_1471_v141 = _nw226
			var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_1471_v141), 0))
			_ = _index228
			(_1471_v141).ArraySet1((_dafny.Zero).Minus(_1470_cf30), (_index228).Int())
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_1471_v141), 0))
			_ = _index229
			(_1471_v141).ArraySet1(_dafny.IntOfInt64(335), (_index229).Int())
		} else if _source16.Is_DC14() {
			var _1472___mcc_h18 _dafny.Sequence = _source16.Get_().(D6_DC14).Cf29
			_ = _1472___mcc_h18
			var _1473_cf29 _dafny.Sequence = _1472___mcc_h18
			_ = _1473_cf29
			var _1474_v142 _dafny.Int
			_ = _1474_v142
			_1474_v142 = _dafny.IntOfInt64(308)
			_1474_v142 = (_dafny.Zero).Minus(p0)
			var _1475_v143 *C0
			_ = _1475_v143
			var _nw227 *C0 = New_C0_()
			_ = _nw227
			_nw227.Ctor__(_1473_cf29)
			_1475_v143 = _nw227
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
			_ = _index230
			var _rhs156 *C0 = _1475_v143
			_ = _rhs156
			var _rhs157 _dafny.Int = _1474_v142
			_ = _rhs157
			var _rhs158 bool = (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)
			_ = _rhs158
			var _lhs92 _dafny.Array = _1275_v0
			_ = _lhs92
			var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
			_ = _lhs93
			_1475_v143 = _rhs156
			_1474_v142 = _rhs157
			(_lhs92).ArraySet1(_rhs158, (_lhs93).Int())
			var _rhs159 D6 = _1464_v140
			_ = _rhs159
			var _rhs160 _dafny.Int = (_1474_v142).Plus((_1418_v100).Cardinality())
			_ = _rhs160
			var _rhs161 *C0 = _1475_v143
			_ = _rhs161
			_1464_v140 = _rhs159
			_1474_v142 = _rhs160
			_1475_v143 = _rhs161
			var _1476_v144 _dafny.Array
			_ = _1476_v144
			var _nw228 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
			_ = _nw228
			_1476_v144 = _nw228
			var _1477_v145 _dafny.MultiSet
			_ = _1477_v145
			_1477_v145 = _dafny.MultiSetOf(_1277_v1)
			var _1478_v146 D9
			_ = _1478_v146
			_1478_v146 = Companion_D9_.Create_DC21_(_1477_v145)
			var _pat_let_tv38 = _1461_v137
			_ = _pat_let_tv38
			var _1479_v147 _dafny.Array
			_ = _1479_v147
			var _nwElement0_53 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1277_v1, p0)).Cardinality()
			_ = _nwElement0_53
			var _nw229 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(23))
			_ = _nw229
			(_nw229).ArraySet1(_nwElement0_53, 0)
			(_nw229).ArraySet1((_1420_v102).Dtor_cf2(), 1)
			(_nw229).ArraySet1(p0, 2)
			(_nw229).ArraySet1((_dafny.Zero).Minus(p0), 3)
			(_nw229).ArraySet1(p0, 4)
			(_nw229).ArraySet1(_1474_v142, 5)
			(_nw229).ArraySet1(p0, 6)
			(_nw229).ArraySet1(p0, 7)
			(_nw229).ArraySet1((_dafny.Zero).Minus(p0), 8)
			(_nw229).ArraySet1(p0, 9)
			(_nw229).ArraySet1(_1474_v142, 10)
			(_nw229).ArraySet1((_dafny.Zero).Minus(((func(_pat_let34_0 D9) D9 {
				return func(_1480_dt__update__tmp_h0 D9) D9 {
					return func(_pat_let35_0 _dafny.MultiSet) D9 {
						return func(_1481_dt__update_hcf37_h0 _dafny.MultiSet) D9 {
							return Companion_D9_.Create_DC21_(_1481_dt__update_hcf37_h0)
						}(_pat_let35_0)
					}(_dafny.MultiSetFromSeq(_pat_let_tv38))
				}(_pat_let34_0)
			}(_1478_v146)).Dtor_cf37()).Cardinality()), 11)
			(_nw229).ArraySet1(_1474_v142, 12)
			(_nw229).ArraySet1(p0, 13)
			(_nw229).ArraySet1(_1474_v142, 14)
			(_nw229).ArraySet1(_1474_v142, 15)
			(_nw229).ArraySet1(p0, 16)
			(_nw229).ArraySet1(_1474_v142, 17)
			(_nw229).ArraySet1(p0, 18)
			(_nw229).ArraySet1(_dafny.IntOfInt64(70), 19)
			(_nw229).ArraySet1(p0, 20)
			(_nw229).ArraySet1(p0, 21)
			(_nw229).ArraySet1(_1474_v142, 22)
			_1479_v147 = _nw229
			var _1482_v148 _dafny.Array
			_ = _1482_v148
			var _nwElement0_54 _dafny.Array = _1479_v147
			_ = _nwElement0_54
			var _nw230 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(8))
			_ = _nw230
			(_nw230).ArraySet1(_nwElement0_54, 0)
			(_nw230).ArraySet1(_1479_v147, 1)
			(_nw230).ArraySet1(_1479_v147, 2)
			(_nw230).ArraySet1(_1479_v147, 3)
			(_nw230).ArraySet1(_1479_v147, 4)
			(_nw230).ArraySet1(_1479_v147, 5)
			(_nw230).ArraySet1(_1479_v147, 6)
			(_nw230).ArraySet1(_1479_v147, 7)
			_1482_v148 = _nw230
			var _1483_v149 _dafny.Map
			_ = _1483_v149
			_1483_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1476_v144, _1482_v148)
			var _1484_v150 _dafny.Sequence
			_ = _1484_v150
			_1484_v150 = _dafny.SeqOf(_1476_v144)
			var _1485_v151 _dafny.Sequence
			_ = _1485_v151
			_1485_v151 = _dafny.SeqOf((func() _dafny.Array {
				if (_1483_v149).Contains((_1484_v150).Select((Companion_Default___.SafeIndex(_1474_v142, _dafny.IntOfUint32((_1484_v150).Cardinality()))).Uint32()).(_dafny.Array)) {
					return (_1483_v149).Get((_1484_v150).Select((Companion_Default___.SafeIndex(_1474_v142, _dafny.IntOfUint32((_1484_v150).Cardinality()))).Uint32()).(_dafny.Array)).(_dafny.Array)
				}
				return _1482_v148
			})())
			var _1486_v152 _dafny.Array
			_ = _1486_v152
			var _nwElement0_55 _dafny.Array = (_1485_v151).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1485_v151).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _nwElement0_55
			var _nw231 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(28))
			_ = _nw231
			(_nw231).ArraySet1(_nwElement0_55, 0)
			(_nw231).ArraySet1(_1482_v148, 1)
			(_nw231).ArraySet1(_1482_v148, 2)
			(_nw231).ArraySet1(_1482_v148, 3)
			(_nw231).ArraySet1((func() _dafny.Array {
				if (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool) {
					return _1482_v148
				}
				return _1482_v148
			})(), 4)
			(_nw231).ArraySet1(_1482_v148, 5)
			(_nw231).ArraySet1(_1482_v148, 6)
			(_nw231).ArraySet1(_1482_v148, 7)
			(_nw231).ArraySet1(_1482_v148, 8)
			(_nw231).ArraySet1(_1482_v148, 9)
			(_nw231).ArraySet1(_1482_v148, 10)
			(_nw231).ArraySet1(_1482_v148, 11)
			(_nw231).ArraySet1(_1482_v148, 12)
			(_nw231).ArraySet1(_1482_v148, 13)
			(_nw231).ArraySet1(_1482_v148, 14)
			(_nw231).ArraySet1(_1482_v148, 15)
			(_nw231).ArraySet1(_1482_v148, 16)
			(_nw231).ArraySet1(_1482_v148, 17)
			(_nw231).ArraySet1(_1482_v148, 18)
			(_nw231).ArraySet1(_1482_v148, 19)
			(_nw231).ArraySet1(_1482_v148, 20)
			(_nw231).ArraySet1(_1482_v148, 21)
			(_nw231).ArraySet1(_1482_v148, 22)
			(_nw231).ArraySet1((func() _dafny.Array {
				if true {
					return _1482_v148
				}
				return _1482_v148
			})(), 23)
			(_nw231).ArraySet1((func() _dafny.Array {
				if _1277_v1 {
					return _1482_v148
				}
				return _1482_v148
			})(), 24)
			(_nw231).ArraySet1(_1482_v148, 25)
			(_nw231).ArraySet1(_1482_v148, 26)
			(_nw231).ArraySet1(_1482_v148, 27)
			_1486_v152 = _nw231
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_1486_v152), 0))
			_ = _index231
			(_1486_v152).ArraySet1(_1482_v148, (_index231).Int())
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1479_v147), 0))
			_ = _index232
			(_1479_v147).ArraySet1(((_1418_v100).Merge(_1418_v100)).Cardinality(), (_index232).Int())
			var _1487_v153 _dafny.MultiSet
			_ = _1487_v153
			_1487_v153 = _dafny.MultiSetOf(Companion_Default___.Fm20(true, globalState))
			var _1488_v154 _dafny.MultiSet
			_ = _1488_v154
			_1488_v154 = _dafny.MultiSetOf(_1487_v153, _1487_v153)
			var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_1486_v152), 0))
			_ = _index233
			var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
			_ = _index234
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1479_v147), 0))
			_ = _index235
			var _rhs162 _dafny.Array = _1482_v148
			_ = _rhs162
			var _rhs163 bool = (_dafny.MultiSetOf(_1487_v153, _1487_v153, Companion_Default___.Fm24((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), globalState))).IsSubsetOf(_1488_v154)
			_ = _rhs163
			var _rhs164 _dafny.Int = (func() _dafny.Int {
				if (_this).Fm5(p0, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), _1421_v103, _1277_v1, globalState) {
					return _dafny.IntOfInt64(502)
				}
				return p0
			})()
			_ = _rhs164
			var _rhs165 _dafny.Int = _1474_v142
			_ = _rhs165
			var _rhs166 bool = _1277_v1
			_ = _rhs166
			var _lhs94 _dafny.Array = _1486_v152
			_ = _lhs94
			var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_1486_v152), 0))
			_ = _lhs95
			var _lhs96 _dafny.Array = _1275_v0
			_ = _lhs96
			var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))
			_ = _lhs97
			var _lhs98 _dafny.Array = _1479_v147
			_ = _lhs98
			var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1479_v147), 0))
			_ = _lhs99
			(_lhs94).ArraySet1(_rhs162, (_lhs95).Int())
			(_lhs96).ArraySet1(_rhs163, (_lhs97).Int())
			_1474_v142 = _rhs164
			(_lhs98).ArraySet1(_rhs165, (_lhs99).Int())
			r0 = _rhs166
		} else {
			var _1489___mcc_h19 D6 = _source16.Get_().(D6_DC16).Cf33
			_ = _1489___mcc_h19
			var _1490_cf33 D6 = _1489___mcc_h19
			_ = _1490_cf33
			var _1491_v155 _dafny.Map
			_ = _1491_v155
			_1491_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)))
			var _1492_v156 _dafny.Sequence
			_ = _1492_v156
			_1492_v156 = _dafny.SeqOf(p0, p0, p0, (_1491_v155).Cardinality())
			var _1493_v157 _dafny.Map
			_ = _1493_v157
			_1493_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_1492_v156), Companion_Default___.Fm20((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), globalState))
			var _1494_v158 _dafny.Set
			_ = _1494_v158
			_1494_v158 = _dafny.SetOf((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), (_1461_v137).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.IntOfUint32((_1461_v137).Cardinality()))).Uint32()).(bool), _1277_v1)
			var _1495_v159 _dafny.Array
			_ = _1495_v159
			var _nwElement0_56 _dafny.Int = p0
			_ = _nwElement0_56
			var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(29))
			_ = _nw232
			(_nw232).ArraySet1(_nwElement0_56, 0)
			(_nw232).ArraySet1((p0).Times(p0), 1)
			(_nw232).ArraySet1((_dafny.Zero).Minus(p0), 2)
			(_nw232).ArraySet1(p0, 3)
			(_nw232).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), 4)
			(_nw232).ArraySet1(p0, 5)
			(_nw232).ArraySet1((_dafny.Zero).Minus(p0), 6)
			(_nw232).ArraySet1(Companion_Default___.Fm20(_1277_v1, globalState), 7)
			(_nw232).ArraySet1(p0, 8)
			(_nw232).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_1418_v100, _1418_v100)).Cardinality()), 9)
			(_nw232).ArraySet1(Companion_Default___.Fm20(!(false), globalState), 10)
			(_nw232).ArraySet1(p0, 11)
			(_nw232).ArraySet1(p0, 12)
			(_nw232).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(_1277_v1, false, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))).Cardinality()), (_1493_v157).Cardinality()), 13)
			(_nw232).ArraySet1(p0, 14)
			(_nw232).ArraySet1(((_1494_v158).Union(_1494_v158)).Cardinality(), 15)
			(_nw232).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1461_v137).Cardinality()), p0), 16)
			(_nw232).ArraySet1(p0, 17)
			(_nw232).ArraySet1(p0, 18)
			(_nw232).ArraySet1((_dafny.IntOfInt64(199)).Times(p0), 19)
			(_nw232).ArraySet1(((_dafny.Zero).Minus(p0)).Plus(p0), 20)
			(_nw232).ArraySet1(p0, 21)
			(_nw232).ArraySet1(p0, 22)
			(_nw232).ArraySet1((func() _dafny.Int {
				if (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool) {
					return _dafny.IntOfUint32((_1461_v137).Cardinality())
				}
				return p0
			})(), 23)
			(_nw232).ArraySet1(_dafny.IntOfInt64(-208), 24)
			(_nw232).ArraySet1(p0, 25)
			(_nw232).ArraySet1(Companion_Default___.Fm20((_this).Fm5(p0, _1277_v1, _1421_v103, _1277_v1, globalState), globalState), 26)
			(_nw232).ArraySet1(_dafny.IntOfInt64(-406), 27)
			(_nw232).ArraySet1(p0, 28)
			_1495_v159 = _nw232
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_1495_v159), 0))
			_ = _index236
			(_1495_v159).ArraySet1((_dafny.Zero).Minus(p0), (_index236).Int())
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_1495_v159), 0))
			_ = _index237
			(_1495_v159).ArraySet1(p0, (_index237).Int())
			var _1496_v160 _dafny.Sequence
			_ = _1496_v160
			_1496_v160 = _dafny.SeqOf(_1460_v136)
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_1495_v159), 0))
			_ = _index238
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_1495_v159), 0))
			_ = _index239
			var _rhs167 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_1496_v160).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1496_v160).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("wqfk"))).Cardinality())).Plus(p0)
			_ = _rhs167
			var _rhs168 _dafny.Int = p0
			_ = _rhs168
			var _lhs100 _dafny.Array = _1495_v159
			_ = _lhs100
			var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_1495_v159), 0))
			_ = _lhs101
			var _lhs102 _dafny.Array = _1495_v159
			_ = _lhs102
			var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_1495_v159), 0))
			_ = _lhs103
			(_lhs100).ArraySet1(_rhs167, (_lhs101).Int())
			(_lhs102).ArraySet1(_rhs168, (_lhs103).Int())
			var _1497_v161 _dafny.MultiSet
			_ = _1497_v161
			_1497_v161 = _dafny.MultiSetOf((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), _1277_v1)
			var _1498_v162 _dafny.Map
			_ = _1498_v162
			_1498_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1460_v136, _1497_v161)
			var _1499_v163 _dafny.MultiSet
			_ = _1499_v163
			_1499_v163 = _dafny.MultiSetOf(_1498_v162, _1498_v162)
			var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_1495_v159), 0))
			_ = _index240
			(_1495_v159).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_1497_v161).Cardinality(), ((func() _dafny.Int {
				if (_1499_v163).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1460_v136, _1497_v161)) {
					return (_1499_v163).Multiplicity(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1460_v136, _1497_v161))
				}
				return p0
			})()).Plus((_1495_v159).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_1495_v159), 0))).Int()).(_dafny.Int)), p0)).Cardinality()), (_index240).Int())
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_1495_v159), 0))
			_ = _index241
			(_1495_v159).ArraySet1(_dafny.IntOfInt64(468), (_index241).Int())
			var _1500_v164 _dafny.Map
			_ = _1500_v164
			_1500_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1277_v1, _1460_v136)
			var _1501_v165 bool
			_ = _1501_v165
			var _1502_v166 _dafny.Int
			_ = _1502_v166
			var _out45 bool
			_ = _out45
			var _out46 _dafny.Int
			_ = _out46
			_out45, _out46 = (_this).M10(((_1500_v164).Cardinality()).Plus(p0), (_1495_v159).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_1495_v159), 0))).Int()).(_dafny.Int), !((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)), _1275_v0, globalState)
			_1501_v165 = _out45
			_1502_v166 = _out46
		}
		var _1503_v167 _dafny.Array
		_ = _1503_v167
		var _len0_33 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_33
		var _nw233 _dafny.Array
		_ = _nw233
		if _len0_33.Cmp(_dafny.Zero) == 0 {
			_nw233 = _dafny.NewArray(_len0_33)
		} else {
			var _init33 func(_dafny.Int) _dafny.MultiSet = func(_1504_i10 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(true)
			}
			_ = _init33
			var _element0_33 = _init33(_dafny.Zero)
			_ = _element0_33
			_nw233 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
			(_nw233).ArraySet1(_element0_33, 0)
			var _nativeLen0_33 = (_len0_33).Int()
			_ = _nativeLen0_33
			for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
				(_nw233).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
			}
		}
		_1503_v167 = _nw233
		var _nw234 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(20))
		_ = _nw234
		_1503_v167 = _nw234
		r0 = !((_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool)) || ((_this).Fm5(p0, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool), _1421_v103, false, globalState))
		var _1505_v168 _dafny.Map
		_ = _1505_v168
		_1505_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1275_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1275_v0), 0))).Int()).(bool))
		r1 = (_1505_v168).Equals(_1505_v168)
		return r0, r1
	}
}
func (_this *C4) M2(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _1506_v0 bool
		_ = _1506_v0
		_1506_v0 = false
		var _1507_v1 D6
		_ = _1507_v1
		_1507_v1 = Companion_D6_.Create_DC14_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(491))).Uint32(), func(coer89 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg89 _dafny.Int) interface{} {
				return coer89(arg89)
			}
		}(func(_1508_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})))
		var _1509_v2 *C0
		_ = _1509_v2
		var _nw235 *C0 = New_C0_()
		_ = _nw235
		_nw235.Ctor__((func() _dafny.Sequence {
			if _1506_v0 {
				return (_1507_v1).Dtor_cf29()
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("gpmds")
		})())
		_1509_v2 = _nw235
		var _1510_v3 _dafny.Map
		_ = _1510_v3
		_1510_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(751))).Uint32(), func(coer90 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg90 _dafny.Int) interface{} {
				return coer90(arg90)
			}
		}((func(_1511_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1512_i1 _dafny.Int) _dafny.Int {
				return _1511_p0
			}
		})(p0)))).Cardinality()), _1506_v0)
		var _1513_v4 _dafny.MultiSet
		_ = _1513_v4
		_1513_v4 = _dafny.MultiSetOf(_1506_v0)
		var _1514_v5 _dafny.MultiSet
		_ = _1514_v5
		_1514_v5 = _dafny.MultiSetOf(p0, p0, p0)
		var _1515_v6 _dafny.Set
		_ = _1515_v6
		_1515_v6 = _dafny.SetOf(p0, p0)
		var _1516_v7 _dafny.Map
		_ = _1516_v7
		_1516_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1514_v5, (_1515_v6).Cardinality())
		var _1517_v8 _dafny.Map
		_ = _1517_v8
		_1517_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1516_v7).Cardinality(), _dafny.IntOfInt64(828))
		var _1518_v9 _dafny.Map
		_ = _1518_v9
		_1518_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1517_v8).Cardinality())
		var _1519_v10 _dafny.Array
		_ = _1519_v10
		var _nwElement0_57 _dafny.Int = p0
		_ = _nwElement0_57
		var _nw236 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(20))
		_ = _nw236
		(_nw236).ArraySet1(_nwElement0_57, 0)
		(_nw236).ArraySet1(_dafny.IntOfUint32((_1509_v2.F8).Cardinality()), 1)
		(_nw236).ArraySet1(_dafny.IntOfInt64(928), 2)
		(_nw236).ArraySet1(p0, 3)
		(_nw236).ArraySet1(p0, 4)
		(_nw236).ArraySet1(p0, 5)
		(_nw236).ArraySet1(p0, 6)
		(_nw236).ArraySet1(p0, 7)
		(_nw236).ArraySet1(p0, 8)
		(_nw236).ArraySet1((_1513_v4).Cardinality(), 9)
		(_nw236).ArraySet1(p0, 10)
		(_nw236).ArraySet1(p0, 11)
		(_nw236).ArraySet1((_1518_v9).Cardinality(), 12)
		(_nw236).ArraySet1(p0, 13)
		(_nw236).ArraySet1(p0, 14)
		(_nw236).ArraySet1(p0, 15)
		(_nw236).ArraySet1(p0, 16)
		(_nw236).ArraySet1((_dafny.Zero).Minus(p0), 17)
		(_nw236).ArraySet1(p0, 18)
		(_nw236).ArraySet1(p0, 19)
		_1519_v10 = _nw236
		var _1520_v11 _dafny.Map
		_ = _1520_v11
		_1520_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1519_v10, p0)
		var _1521_v12 _dafny.Map
		_ = _1521_v12
		_1521_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1520_v11, _1519_v10)
		var _1522_v13 _dafny.Array
		_ = _1522_v13
		var _nwElement0_58 _dafny.Array = _1519_v10
		_ = _nwElement0_58
		var _nw237 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(15))
		_ = _nw237
		(_nw237).ArraySet1(_nwElement0_58, 0)
		(_nw237).ArraySet1(_1519_v10, 1)
		(_nw237).ArraySet1((func() _dafny.Array {
			if (_1521_v12).Contains(_1520_v11) {
				return (_1521_v12).Get(_1520_v11).(_dafny.Array)
			}
			return _1519_v10
		})(), 2)
		(_nw237).ArraySet1(_1519_v10, 3)
		(_nw237).ArraySet1(_1519_v10, 4)
		(_nw237).ArraySet1(_1519_v10, 5)
		(_nw237).ArraySet1(_1519_v10, 6)
		(_nw237).ArraySet1(_1519_v10, 7)
		(_nw237).ArraySet1(_1519_v10, 8)
		(_nw237).ArraySet1(_1519_v10, 9)
		(_nw237).ArraySet1(_1519_v10, 10)
		(_nw237).ArraySet1(_1519_v10, 11)
		(_nw237).ArraySet1(_1519_v10, 12)
		(_nw237).ArraySet1(_1519_v10, 13)
		(_nw237).ArraySet1(_1519_v10, 14)
		_1522_v13 = _nw237
		var _1523_v14 D2
		_ = _1523_v14
		_1523_v14 = Companion_D2_.Create_DC4_(p0, _1506_v0, _1522_v13)
		var _1524_v15 _dafny.Sequence
		_ = _1524_v15
		_1524_v15 = _dafny.SeqOf((func() _dafny.Int {
			if (_1518_v9).Contains(_1506_v0) {
				return (_1518_v9).Get(_1506_v0).(_dafny.Int)
			}
			return p0
		})(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uotirlywr")).Cardinality()))
		var _1525_v16 _dafny.Array
		_ = _1525_v16
		_1525_v16 = _1519_v10
		var _1526_v17 _dafny.Map
		_ = _1526_v17
		_1526_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1525_v16))
		var _1527_v18 _dafny.Sequence
		_ = _1527_v18
		_1527_v18 = _dafny.SeqOf(_1509_v2)
		var _1528_v19 _dafny.Array
		_ = _1528_v19
		var _nwElement0_59 _dafny.Int = _dafny.IntOfUint32((_1509_v2.F8).Cardinality())
		_ = _nwElement0_59
		var _nw238 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(19))
		_ = _nw238
		(_nw238).ArraySet1(_nwElement0_59, 0)
		(_nw238).ArraySet1(p0, 1)
		(_nw238).ArraySet1((_1510_v3).Cardinality(), 2)
		(_nw238).ArraySet1((p0).Plus(p0), 3)
		(_nw238).ArraySet1((p0).Plus(p0), 4)
		(_nw238).ArraySet1((_dafny.Zero).Minus(p0), 5)
		(_nw238).ArraySet1(p0, 6)
		(_nw238).ArraySet1((p0).Times(p0), 7)
		(_nw238).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(Companion_Default___.Fm20(_1506_v0, globalState))).Cardinality()), 8)
		(_nw238).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfInt64(85)).Plus(_dafny.IntOfUint32((_1509_v2.F8).Cardinality()))), 9)
		(_nw238).ArraySet1(p0, 10)
		(_nw238).ArraySet1(p0, 11)
		(_nw238).ArraySet1((_1523_v14).Dtor_cf5(), 12)
		(_nw238).ArraySet1(p0, 13)
		(_nw238).ArraySet1((_1524_v15).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1524_v15).Cardinality()))).Uint32()).(_dafny.Int), 14)
		(_nw238).ArraySet1(p0, 15)
		(_nw238).ArraySet1((_dafny.Zero).Minus((_1526_v17).Cardinality()), 16)
		(_nw238).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1527_v18).Cardinality()), p0), 17)
		(_nw238).ArraySet1(p0, 18)
		_1528_v19 = _nw238
		var _1529_v20 _dafny.MultiSet
		_ = _1529_v20
		_1529_v20 = _dafny.MultiSetOf((_1507_v1).Dtor_cf29(), _1509_v2.F8)
		var _nwElement0_60 _dafny.Int = (_1523_v14).Dtor_cf5()
		_ = _nwElement0_60
		var _nw239 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(12))
		_ = _nw239
		(_nw239).ArraySet1(_nwElement0_60, 0)
		(_nw239).ArraySet1((_1524_v15).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1524_v15).Cardinality()))).Uint32()).(_dafny.Int), 1)
		(_nw239).ArraySet1(p0, 2)
		(_nw239).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p0), p0), 3)
		(_nw239).ArraySet1(p0, 4)
		(_nw239).ArraySet1((p0).Plus((_1529_v20).Cardinality()), 5)
		(_nw239).ArraySet1(p0, 6)
		(_nw239).ArraySet1((p0).Plus(_dafny.IntOfInt64(75)), 7)
		(_nw239).ArraySet1((_dafny.Zero).Minus((p0).Plus(p0)), 8)
		(_nw239).ArraySet1((p0).Plus(p0), 9)
		(_nw239).ArraySet1(p0, 10)
		(_nw239).ArraySet1((p0).Minus(p0), 11)
		_1528_v19 = _nw239
		var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))
		_ = _index242
		(_1519_v10).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1509_v2.F8).Cardinality()), p0), (_index242).Int())
		var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))
		_ = _index243
		(_1519_v10).ArraySet1(p0, (_index243).Int())
		var _1530_i2 _dafny.Int
		_ = _1530_i2
		_1530_i2 = _dafny.Zero
		{
			for !(false) {
				{
					if (_1530_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_1530_i2 = (_1530_i2).Plus(_dafny.One)
					var _1531_v21 D4
					_ = _1531_v21
					_1531_v21 = Companion_D4_.Create_DC10_(p0, p0, (_1519_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))).Int()).(_dafny.Int))
					var _1532_v22 _dafny.CodePoint
					_ = _1532_v22
					_1532_v22 = _dafny.CodePoint('t')
					var _1533_v23 _dafny.Map
					_ = _1533_v23
					_1533_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm25(_1531_v21, _1509_v2.F8, globalState), _1532_v22)
					var _1534_v24 _dafny.Map
					_ = _1534_v24
					_1534_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1532_v22, _1533_v23)
					var _1535_v25 _dafny.Sequence
					_ = _1535_v25
					_1535_v25 = _dafny.SeqOf(_1513_v4)
					var _1536_v27 _dafny.Set
					_ = _1536_v27
					_1536_v27 = _dafny.SetOf(_1509_v2.F8)
					var _1537_v28 _dafny.Map
					_ = _1537_v28
					_1537_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), Companion_D2_.Create_DC3_(_1536_v27))
					var _1538_v29 _dafny.Map
					_ = _1538_v29
					_1538_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1509_v2.F8, _1537_v28)
					var _1539_v30 _dafny.Set
					_ = _1539_v30
					_1539_v30 = _dafny.SetOf(_1506_v0)
					var _1540_v31 _dafny.Map
					_ = _1540_v31
					_1540_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqOf(_1506_v0))
					var _1541_v32 _dafny.Map
					_ = _1541_v32
					_1541_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1509_v2, p0)
					var _1542_v33 _dafny.Sequence
					_ = _1542_v33
					_1542_v33 = _dafny.SeqOf(_1506_v0, _1506_v0)
					var _1543_v34 D0
					_ = _1543_v34
					_1543_v34 = Companion_D0_.Create_DC1_((_1519_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))).Int()).(_dafny.Int))
					var _1544_v37 _dafny.Map
					_ = _1544_v37
					_1544_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1509_v2.F8, _1506_v0)
					var _1545_v40 _dafny.Map
					_ = _1545_v40
					_1545_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm25(_1531_v21, _1509_v2.F8, globalState), (_1519_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))).Int()).(_dafny.Int))
					var _1546_v41 D12
					_ = _1546_v41
					_1546_v41 = Companion_D12_.Create_DC27_(_1539_v30)
					var _1547_v42 D11
					_ = _1547_v42
					_1547_v42 = Companion_D11_.Create_DC24_(Companion_Default___.Fm26((_1546_v41).Dtor_cf45(), _1542_v33, _1543_v34, _1532_v22, globalState))
					var _1548_v43 _dafny.Array
					_ = _1548_v43
					var _nwElement0_61 _dafny.Map = _1533_v23
					_ = _nwElement0_61
					var _nw240 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(29))
					_ = _nw240
					(_nw240).ArraySet1(_nwElement0_61, 0)
					(_nw240).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1509_v2.F8, _1532_v22), 1)
					(_nw240).ArraySet1(_1533_v23, 2)
					(_nw240).ArraySet1(_1533_v23, 3)
					(_nw240).ArraySet1(_1533_v23, 4)
					(_nw240).ArraySet1((_1533_v23).Update(_1509_v2.F8, _1532_v22), 5)
					(_nw240).ArraySet1((func() _dafny.Map {
						if (_1534_v24).Contains(_1532_v22) {
							return (_1534_v24).Get(_1532_v22).(_dafny.Map)
						}
						return _1533_v23
					})(), 6)
					(_nw240).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1509_v2).Fm15(_dafny.SeqOf(_dafny.IntOfUint32((_1535_v25).Cardinality())), _1506_v0, globalState), _dafny.CodePoint('s')), 7)
					(_nw240).ArraySet1(_1533_v23, 8)
					(_nw240).ArraySet1(func() _dafny.Map {
						var _coll53 = _dafny.NewMapBuilder()
						_ = _coll53
						for _iter55 := _dafny.Iterate((_1538_v29).Keys().Elements()); ; {
							_compr_53, _ok55 := _iter55()
							if !_ok55 {
								break
							}
							var _1549_v26 _dafny.Sequence
							_1549_v26 = interface{}(_compr_53).(_dafny.Sequence)
							if (_1538_v29).Contains(_1549_v26) {
								_coll53.Add(_1549_v26, _dafny.CodePoint('t'))
							}
						}
						return _coll53.ToMap()
					}(), 9)
					(_nw240).ArraySet1(_1533_v23, 10)
					(_nw240).ArraySet1(_1533_v23, 11)
					(_nw240).ArraySet1(_1533_v23, 12)
					(_nw240).ArraySet1(_1533_v23, 13)
					(_nw240).ArraySet1(Companion_Default___.Fm26(_1539_v30, (func() _dafny.Sequence {
						if (_1540_v31).Contains((func() _dafny.Int {
							if (_1541_v32).Contains(_1509_v2) {
								return (_1541_v32).Get(_1509_v2).(_dafny.Int)
							}
							return (_1519_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))).Int()).(_dafny.Int)
						})()) {
							return (_1540_v31).Get((func() _dafny.Int {
								if (_1541_v32).Contains(_1509_v2) {
									return (_1541_v32).Get(_1509_v2).(_dafny.Int)
								}
								return (_1519_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))).Int()).(_dafny.Int)
							})()).(_dafny.Sequence)
						}
						return _1542_v33
					})(), _1543_v34, _1532_v22, globalState), 14)
					(_nw240).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1509_v2.F8, (_1509_v2).Fm14(globalState))).Merge((func() _dafny.Map {
						var _coll54 = _dafny.NewMapBuilder()
						_ = _coll54
						for _iter56 := _dafny.Iterate((_1529_v20).Elements()); ; {
							_compr_54, _ok56 := _iter56()
							if !_ok56 {
								break
							}
							var _1550_v35 _dafny.Sequence
							_1550_v35 = interface{}(_compr_54).(_dafny.Sequence)
							if (_1529_v20).Contains(_1550_v35) {
								_coll54.Add(_1550_v35, _1532_v22)
							}
						}
						return _coll54.ToMap()
					}()).Update(_dafny.UnicodeSeqOfUtf8Bytes("u"), _1532_v22)), 15)
					(_nw240).ArraySet1(_1533_v23, 16)
					(_nw240).ArraySet1((_1533_v23).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1509_v2.F8, (_1509_v2.F8).Select((Companion_Default___.SafeIndex((_1519_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1509_v2.F8).Cardinality()))).Uint32()).(_dafny.CodePoint))), 17)
					(_nw240).ArraySet1(func() _dafny.Map {
						var _coll55 = _dafny.NewMapBuilder()
						_ = _coll55
						for _iter57 := _dafny.Iterate((_1544_v37).Keys().Elements()); ; {
							_compr_55, _ok57 := _iter57()
							if !_ok57 {
								break
							}
							var _1551_v36 _dafny.Sequence
							_1551_v36 = interface{}(_compr_55).(_dafny.Sequence)
							if (_1544_v37).Contains(_1551_v36) {
								_coll55.Add(_1551_v36, _1532_v22)
							}
						}
						return _coll55.ToMap()
					}(), 18)
					(_nw240).ArraySet1(_1533_v23, 19)
					(_nw240).ArraySet1(_1533_v23, 20)
					(_nw240).ArraySet1(_1533_v23, 21)
					(_nw240).ArraySet1(func() _dafny.Map {
						var _coll56 = _dafny.NewMapBuilder()
						_ = _coll56
						for _iter58 := _dafny.Iterate((func() _dafny.Map {
							var _coll57 = _dafny.NewMapBuilder()
							_ = _coll57
							for _iter59 := _dafny.Iterate((_1545_v40).Keys().Elements()); ; {
								_compr_57, _ok59 := _iter59()
								if !_ok59 {
									break
								}
								var _1552_v39 _dafny.Sequence
								_1552_v39 = interface{}(_compr_57).(_dafny.Sequence)
								if (_1545_v40).Contains(_1552_v39) {
									_coll57.Add(_1552_v39, _1506_v0)
								}
							}
							return _coll57.ToMap()
						}()).Keys().Elements()); ; {
							_compr_56, _ok58 := _iter58()
							if !_ok58 {
								break
							}
							var _1553_v38 _dafny.Sequence
							_1553_v38 = interface{}(_compr_56).(_dafny.Sequence)
							if (func() _dafny.Map {
								var _coll58 = _dafny.NewMapBuilder()
								_ = _coll58
								for _iter60 := _dafny.Iterate((_1545_v40).Keys().Elements()); ; {
									_compr_58, _ok60 := _iter60()
									if !_ok60 {
										break
									}
									var _1552_v39 _dafny.Sequence
									_1552_v39 = interface{}(_compr_58).(_dafny.Sequence)
									if (_1545_v40).Contains(_1552_v39) {
										_coll58.Add(_1552_v39, _1506_v0)
									}
								}
								return _coll58.ToMap()
							}()).Contains(_1553_v38) {
								_coll56.Add(_1553_v38, _1532_v22)
							}
						}
						return _coll56.ToMap()
					}(), 22)
					(_nw240).ArraySet1(_1533_v23, 23)
					(_nw240).ArraySet1(_1533_v23, 24)
					(_nw240).ArraySet1((_1547_v42).Dtor_cf40(), 25)
					(_nw240).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(890))).Uint32(), func(coer91 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg91 _dafny.Int) interface{} {
							return coer91(arg91)
						}
					}((func(_1554_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1555_i3 _dafny.Int) _dafny.CodePoint {
							return _1554_v22
						}
					})(_1532_v22))), _1532_v22), 26)
					(_nw240).ArraySet1((_1533_v23).Merge(_1533_v23), 27)
					(_nw240).ArraySet1(_1533_v23, 28)
					_1548_v43 = _nw240
					var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_1548_v43), 0))
					_ = _index244
					(_1548_v43).ArraySet1(_1533_v23, (_index244).Int())
					var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_1548_v43), 0))
					_ = _index245
					(_1548_v43).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(136))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg92 _dafny.Int) interface{} {
							return coer92(arg92)
						}
					}((func(_1556_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1557_i4 _dafny.Int) _dafny.CodePoint {
							return _1556_v22
						}
					})(_1532_v22))), _1532_v22), (_index245).Int())
					var _1558_v44 _dafny.Array
					_ = _1558_v44
					var _nw241 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
					_ = _nw241
					_1558_v44 = _nw241
					var _1559_v45 _dafny.Map
					_ = _1559_v45
					_1559_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1532_v22, (_1515_v6).Cardinality())
					var _1560_v46 _dafny.MultiSet
					_ = _1560_v46
					_1560_v46 = _dafny.MultiSetOf(_1559_v45, _1559_v45)
					var _1561_v47 T0
					_ = _1561_v47
					var _nw242 *C2 = New_C2_()
					_ = _nw242
					_nw242.Ctor__()
					_1561_v47 = _nw242
					var _1562_v48 _dafny.Map
					_ = _1562_v48
					_1562_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1561_v47)
					var _1563_v49 _dafny.Map
					_ = _1563_v49
					_1563_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1560_v46, _1562_v48)
					var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.ArrayLen((_1558_v44), 0))
					_ = _index246
					(_1558_v44).ArraySet1((func() _dafny.Map {
						if _1506_v0 {
							return _1563_v49
						}
						return _1563_v49
					})(), (_index246).Int())
					var _1564_v50 _dafny.Sequence
					_ = _1564_v50
					_1564_v50 = _dafny.SeqOf(_1509_v2.F8, _1509_v2.F8)
					var _1565_v51 _dafny.Map
					_ = _1565_v51
					_1565_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1506_v0, _1562_v48)
					var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.ArrayLen((_1558_v44), 0))
					_ = _index247
					var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))
					_ = _index248
					var _rhs169 bool = _1506_v0
					_ = _rhs169
					var _rhs170 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("xone"), (_1564_v50).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1564_v50).Cardinality()))).Uint32()).(_dafny.Sequence))
					_ = _rhs170
					var _rhs171 _dafny.Map = ((func() _dafny.Map {
						if _1506_v0 {
							return _1563_v49
						}
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_1559_v45), _1562_v48)
					})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1560_v46, (func() _dafny.Map {
						if (_1565_v51).Contains(_1506_v0) {
							return (_1565_v51).Get(_1506_v0).(_dafny.Map)
						}
						return _1562_v48
					})()))
					_ = _rhs171
					var _rhs172 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
						if !(_1510_v3).Equals(_1510_v3) {
							return (p0).Plus(Companion_Default___.Fm20(_1506_v0, globalState))
						}
						return ((_1519_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))).Int()).(_dafny.Int)).Plus(p0)
					})())
					_ = _rhs172
					var _lhs104 _dafny.Array = _1558_v44
					_ = _lhs104
					var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.ArrayLen((_1558_v44), 0))
					_ = _lhs105
					var _lhs106 _dafny.Array = _1519_v10
					_ = _lhs106
					var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))
					_ = _lhs107
					_1506_v0 = _rhs169
					_1506_v0 = _rhs170
					(_lhs104).ArraySet1(_rhs171, (_lhs105).Int())
					(_lhs106).ArraySet1(_rhs172, (_lhs107).Int())
					var _1566_v52 _dafny.Map
					_ = _1566_v52
					_1566_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), _1543_v34)
					var _1567_v53 D2
					_ = _1567_v53
					_1567_v53 = Companion_D2_.Create_DC5_((_1519_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))).Int()).(_dafny.Int), _1524_v15, false)
					if (_this).Fm5((_1519_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))).Int()).(_dafny.Int), !(((_1519_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))).Int()).(_dafny.Int)).Cmp(p0) <= 0), _1566_v52, (_1567_v53).Dtor_cf10(), globalState) {
						var _1568_v54 _dafny.Array
						_ = _1568_v54
						var _len0_34 _dafny.Int = _dafny.IntOfInt64(12)
						_ = _len0_34
						var _nw243 _dafny.Array
						_ = _nw243
						if _len0_34.Cmp(_dafny.Zero) == 0 {
							_nw243 = _dafny.NewArray(_len0_34)
						} else {
							var _init34 func(_dafny.Int) bool = func(_1569_i5 _dafny.Int) bool {
								return true
							}
							_ = _init34
							var _element0_34 = _init34(_dafny.Zero)
							_ = _element0_34
							_nw243 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
							(_nw243).ArraySet1(_element0_34, 0)
							var _nativeLen0_34 = (_len0_34).Int()
							_ = _nativeLen0_34
							for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
								(_nw243).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
							}
						}
						_1568_v54 = _nw243
						_1568_v54 = _1568_v54
						var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_1522_v13), 0))
						_ = _index249
						(_1522_v13).ArraySet1(_1528_v19, (_index249).Int())
						var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_1522_v13), 0))
						_ = _index250
						(_1522_v13).ArraySet1(_1528_v19, (_index250).Int())
						var _1570_v55 bool
						_ = _1570_v55
						var _1571_v56 bool
						_ = _1571_v56
						var _out47 bool
						_ = _out47
						var _out48 bool
						_ = _out48
						_out47, _out48 = (_1561_v47).M1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(5))), globalState)
						_1570_v55 = _out47
						_1571_v56 = _out48
						var _1572_v57 _dafny.Map
						_ = _1572_v57
						_1572_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1568_v54, true)
						_1571_v56 = (_1572_v57).Equals(_1572_v57)
						var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))
						_ = _index251
						(_1519_v10).ArraySet1((_1519_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))).Int()).(_dafny.Int), (_index251).Int())
					} else {
						var _1573_v58 D7
						_ = _1573_v58
						_1573_v58 = Companion_D7_.Create_DC17_(_1542_v33)
						var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))
						_ = _index252
						(_1519_v10).ArraySet1(_dafny.IntOfUint32(((_1573_v58).Dtor_cf34()).Cardinality()), (_index252).Int())
						(_1509_v2).F8 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(324))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg93 _dafny.Int) interface{} {
								return coer93(arg93)
							}
						}((func(_1574_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1575_i6 _dafny.Int) _dafny.CodePoint {
								return _1574_v22
							}
						})(_1532_v22)))
						var _1576_v59 D16
						_ = _1576_v59
						_1576_v59 = Companion_D16_.Create_DC44_((_1519_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))).Int()).(_dafny.Int))
						var _pat_let_tv39 = _1524_v15
						_ = _pat_let_tv39
						var _pat_let_tv40 = _1518_v9
						_ = _pat_let_tv40
						var _pat_let_tv41 = _1524_v15
						_ = _pat_let_tv41
						_1576_v59 = func(_pat_let36_0 D16) D16 {
							return func(_1577_dt__update__tmp_h0 D16) D16 {
								return func(_pat_let37_0 _dafny.Int) D16 {
									return func(_1578_dt__update_hcf82_h0 _dafny.Int) D16 {
										return Companion_D16_.Create_DC44_(_1578_dt__update_hcf82_h0)
									}(_pat_let37_0)
								}((_pat_let_tv39).Select((Companion_Default___.SafeIndex((_pat_let_tv40).Cardinality(), _dafny.IntOfUint32((_pat_let_tv41).Cardinality()))).Uint32()).(_dafny.Int))
							}(_pat_let36_0)
						}(_1576_v59)
						_1544_v37 = (_1544_v37).Update(_1509_v2.F8, _1506_v0)
						_1514_v5 = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(311))).Uint32(), func(coer94 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg94 _dafny.Int) interface{} {
								return coer94(arg94)
							}
						}(func(_1579_i7 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(-76)
						})), _1524_v15))).Intersection((_1514_v5).Union(_1514_v5))
					}
					var _1580_v60 D9
					_ = _1580_v60
					_1580_v60 = Companion_D9_.Create_DC21_((Companion_D9_.Create_DC21_(_1513_v4)).Dtor_cf37())
					var _1581_v61 _dafny.Sequence
					_ = _1581_v61
					_1581_v61 = _dafny.SeqOf(_1580_v60, _1580_v60, _1580_v60)
					_1581_v61 = _dafny.Companion_Sequence_.Concatenate(_1581_v61, _1581_v61)
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _1582_v62 *C2
		_ = _1582_v62
		var _nw244 *C2 = New_C2_()
		_ = _nw244
		_nw244.Ctor__()
		_1582_v62 = _nw244
		var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1519_v10), 0))
		_ = _index253
		(_1519_v10).ArraySet1(p0, (_index253).Int())
		var _1583_v63 _dafny.Map
		_ = _1583_v63
		_1583_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1506_v0), _1514_v5)
		r0 = _1583_v63
		return r0
	}
}
func (_this *C4) M9(globalState *GlobalState) (_dafny.Array, _dafny.Array, bool) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 bool = false
		_ = r2
		var _1584_v0 _dafny.CodePoint
		_ = _1584_v0
		_1584_v0 = _dafny.CodePoint('x')
		var _1585_v1 _dafny.Sequence
		_ = _1585_v1
		_1585_v1 = _dafny.UnicodeSeqOfUtf8Bytes("jvypfkdr")
		var _1586_v2 bool
		_ = _1586_v2
		_1586_v2 = false
		var _1587_v3 _dafny.Int
		_ = _1587_v3
		_1587_v3 = _dafny.IntOfInt64(-189)
		var _1588_v4 _dafny.Sequence
		_ = _1588_v4
		_1588_v4 = _dafny.SeqOf(_1586_v2)
		var _1589_v5 _dafny.MultiSet
		_ = _1589_v5
		_1589_v5 = _dafny.MultiSetOf(true, _1586_v2)
		var _1590_v6 D0
		_ = _1590_v6
		_1590_v6 = Companion_D0_.Create_DC1_(_1587_v3)
		var _1591_v7 _dafny.Map
		_ = _1591_v7
		_1591_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1584_v0, _1590_v6)
		var _1592_v8 _dafny.Set
		_ = _1592_v8
		_1592_v8 = _dafny.SetOf(_1586_v2)
		var _1593_v9 _dafny.Array
		_ = _1593_v9
		var _nwElement0_62 bool = _dafny.Companion_Sequence_.Contains(_1585_v1, _1584_v0)
		_ = _nwElement0_62
		var _nw245 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(28))
		_ = _nw245
		(_nw245).ArraySet1(_nwElement0_62, 0)
		(_nw245).ArraySet1(_1586_v2, 1)
		(_nw245).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(343))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg95 _dafny.Int) interface{} {
				return coer95(arg95)
			}
		}((func(_1594_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1595_i0 _dafny.Int) _dafny.CodePoint {
				return _1594_v0
			}
		})(_1584_v0))), _1584_v0), 2)
		(_nw245).ArraySet1(_dafny.Companion_Sequence_.Equal(_1585_v1, Companion_Default___.Fm31(_1587_v3, _1586_v2, _dafny.IntOfUint32((_1588_v4).Cardinality()), globalState)), 3)
		(_nw245).ArraySet1(_1586_v2, 4)
		(_nw245).ArraySet1((_1589_v5).IsDisjointFrom(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).Fm5(_dafny.IntOfUint32((_1585_v1).Cardinality()), _1586_v2, _1591_v7, _1586_v2, globalState)))), 5)
		(_nw245).ArraySet1(_1586_v2, 6)
		(_nw245).ArraySet1(_1586_v2, 7)
		(_nw245).ArraySet1(_1586_v2, 8)
		(_nw245).ArraySet1((_1587_v3).Cmp(_1587_v3) >= 0, 9)
		(_nw245).ArraySet1((_1586_v2) && (_1586_v2), 10)
		(_nw245).ArraySet1(_1586_v2, 11)
		(_nw245).ArraySet1((_1586_v2) && (true), 12)
		(_nw245).ArraySet1(_1586_v2, 13)
		(_nw245).ArraySet1(_1586_v2, 14)
		(_nw245).ArraySet1(_1586_v2, 15)
		(_nw245).ArraySet1(_1586_v2, 16)
		(_nw245).ArraySet1((_1592_v8).IsDisjointFrom(_dafny.SetOf(_1586_v2)), 17)
		(_nw245).ArraySet1(_1586_v2, 18)
		(_nw245).ArraySet1((func() bool {
			if _1586_v2 {
				return !(_1586_v2)
			}
			return true
		})(), 19)
		(_nw245).ArraySet1(_1586_v2, 20)
		(_nw245).ArraySet1(!((func() bool {
			if _1586_v2 {
				return _1586_v2
			}
			return true
		})()), 21)
		(_nw245).ArraySet1(false, 22)
		(_nw245).ArraySet1(_1586_v2, 23)
		(_nw245).ArraySet1((_1587_v3).Cmp(_1587_v3) < 0, 24)
		(_nw245).ArraySet1((_1589_v5).IsSubsetOf(_1589_v5), 25)
		(_nw245).ArraySet1(_1586_v2, 26)
		(_nw245).ArraySet1(_1586_v2, 27)
		_1593_v9 = _nw245
		var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))
		_ = _index254
		(_1593_v9).ArraySet1(_1586_v2, (_index254).Int())
		var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))
		_ = _index255
		(_1593_v9).ArraySet1(_1586_v2, (_index255).Int())
		var _1596_v10 _dafny.Sequence
		_ = _1596_v10
		_1596_v10 = _dafny.SeqOf(_1587_v3)
		var _1597_v11 _dafny.Sequence
		_ = _1597_v11
		_1597_v11 = _dafny.SeqOf(_1596_v10, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(185))).Uint32(), func(coer96 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg96 _dafny.Int) interface{} {
				return coer96(arg96)
			}
		}((func(_1598_v10 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_1599_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_1598_v10).Cardinality())
			}
		})(_1596_v10))))
		var _1600_v12 _dafny.Set
		_ = _1600_v12
		_1600_v12 = _dafny.SetOf(_dafny.SeqOf(_1587_v3), (_1597_v11).Select((Companion_Default___.SafeIndex(_1587_v3, _dafny.IntOfUint32((_1597_v11).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _1601_v13 D15
		_ = _1601_v13
		_1601_v13 = Companion_D15_.Create_DC37_(_1600_v12)
		_1601_v13 = _1601_v13
		var _1602_v14 *C2
		_ = _1602_v14
		var _nw246 *C2 = New_C2_()
		_ = _nw246
		_nw246.Ctor__()
		_1602_v14 = _nw246
		var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))
		_ = _index256
		var _rhs173 *C2 = _1602_v14
		_ = _rhs173
		var _rhs174 _dafny.Int = _dafny.IntOfInt64(689)
		_ = _rhs174
		var _rhs175 _dafny.Int = _1587_v3
		_ = _rhs175
		var _rhs176 bool = !((_1593_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))).Int()).(bool)) || ((_1593_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))).Int()).(bool))
		_ = _rhs176
		var _lhs108 _dafny.Array = _1593_v9
		_ = _lhs108
		var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))
		_ = _lhs109
		_1602_v14 = _rhs173
		_1587_v3 = _rhs174
		_1587_v3 = _rhs175
		(_lhs108).ArraySet1(_rhs176, (_lhs109).Int())
		var _1603_v15 _dafny.Map
		_ = _1603_v15
		_1603_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1587_v3, (_1593_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))).Int()).(bool))
		_1603_v15 = (_1603_v15).Update(_1587_v3, (_dafny.SetOf(_1586_v2, _1586_v2, (_1593_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))).Int()).(bool))).Equals(_1592_v8))
		var _hi10 _dafny.Int = _1587_v3
		_ = _hi10
		for _1604_i2 := (_1587_v3).Minus(_1587_v3); _1604_i2.Cmp(_hi10) < 0; _1604_i2 = _1604_i2.Plus(_dafny.One) {
			_1585_v1 = _1585_v1
			var _rhs177 bool = true
			_ = _rhs177
			var _rhs178 bool = _dafny.Companion_Sequence_.Contains(_1596_v10, Companion_Default___.SafeModuloInt(_1587_v3, _1587_v3))
			_ = _rhs178
			r2 = _rhs177
			_1586_v2 = _rhs178
			var _1605_v16 _dafny.Array
			_ = _1605_v16
			var _nw247 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
			_ = _nw247
			_1605_v16 = _nw247
			var _1606_v17 _dafny.Array
			_ = _1606_v17
			var _nwElement0_63 _dafny.Sequence = _1585_v1
			_ = _nwElement0_63
			var _nw248 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(20))
			_ = _nw248
			(_nw248).ArraySet1(_nwElement0_63, 0)
			(_nw248).ArraySet1(_1585_v1, 1)
			(_nw248).ArraySet1(_1585_v1, 2)
			(_nw248).ArraySet1(_1585_v1, 3)
			(_nw248).ArraySet1(_dafny.Companion_Sequence_.Update(_1585_v1, (Companion_Default___.SafeIndex(_1604_i2, _dafny.IntOfUint32((_1585_v1).Cardinality()))).Uint32(), _1584_v0), 4)
			(_nw248).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("oe"), 5)
			(_nw248).ArraySet1(_1585_v1, 6)
			(_nw248).ArraySet1(_1585_v1, 7)
			(_nw248).ArraySet1(_1585_v1, 8)
			(_nw248).ArraySet1(_1585_v1, 9)
			(_nw248).ArraySet1(_1585_v1, 10)
			(_nw248).ArraySet1(_1585_v1, 11)
			(_nw248).ArraySet1(_1585_v1, 12)
			(_nw248).ArraySet1(_1585_v1, 13)
			(_nw248).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("diut"), 14)
			(_nw248).ArraySet1(_1585_v1, 15)
			(_nw248).ArraySet1(_1585_v1, 16)
			(_nw248).ArraySet1(_1585_v1, 17)
			(_nw248).ArraySet1(_1585_v1, 18)
			(_nw248).ArraySet1(_1585_v1, 19)
			_1606_v17 = _nw248
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_1605_v16), 0))
			_ = _index257
			(_1605_v16).ArraySet1(_1606_v17, (_index257).Int())
			var _1607_v18 _dafny.Sequence
			_ = _1607_v18
			_1607_v18 = _dafny.SeqOf(_1606_v17)
			var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_1605_v16), 0))
			_ = _index258
			(_1605_v16).ArraySet1((_1607_v18).Select((Companion_Default___.SafeIndex(_1587_v3, _dafny.IntOfUint32((_1607_v18).Cardinality()))).Uint32()).(_dafny.Array), (_index258).Int())
			var _1608_v19 _dafny.Set
			_ = _1608_v19
			_1608_v19 = _dafny.SetOf(_1593_v9)
			var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))
			_ = _index259
			(_1593_v9).ArraySet1((_1608_v19).IsProperSubsetOf(_1608_v19), (_index259).Int())
		}
		var _1609_v20 _dafny.Set
		_ = _1609_v20
		_1609_v20 = _dafny.SetOf(_1593_v9, _1593_v9, _1593_v9)
		var _1610_v21 _dafny.Set
		_ = _1610_v21
		_1610_v21 = _1609_v20
		var _1611_v22 _dafny.Set
		_ = _1611_v22
		_1611_v22 = (_1610_v21).Difference(_1609_v20)
		var _source17 _dafny.Set = _1610_v21
		_ = _source17
		var _1612___mcc_h0 _dafny.Set = _source17
		_ = _1612___mcc_h0
		var _1613_cf36 _dafny.Set = _1612___mcc_h0
		_ = _1613_cf36
		_1587_v3 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_1587_v3, _1587_v3), Companion_Default___.SafeDivisionInt(_1587_v3, _dafny.IntOfInt64(317))))
		var _1614_v23 _dafny.Array
		_ = _1614_v23
		var _nw249 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
		_ = _nw249
		_1614_v23 = _nw249
		_1614_v23 = _1614_v23
		var _1615_v25 _dafny.Map
		_ = _1615_v25
		_1615_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1584_v0, _1587_v3)
		var _1616_v27 _dafny.Set
		_ = _1616_v27
		_1616_v27 = _dafny.SetOf(_1584_v0)
		var _1617_v28 _dafny.Map
		_ = _1617_v28
		_1617_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1584_v0, _dafny.SeqOf((_1593_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))).Int()).(bool), (_1593_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))).Int()).(bool)))
		var _1618_v30 D6
		_ = _1618_v30
		_1618_v30 = Companion_D6_.Create_DC14_(_1585_v1)
		var _1619_v31 _dafny.Map
		_ = _1619_v31
		_1619_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1584_v0, _1618_v30)
		var _1620_v32 _dafny.Map
		_ = _1620_v32
		_1620_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1587_v3, _1617_v28)
		var _1621_v33 _dafny.Map
		_ = _1621_v33
		_1621_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1587_v3, _1587_v3)
		var _1622_v35 _dafny.Array
		_ = _1622_v35
		var _nwElement0_64 _dafny.Map = func() _dafny.Map {
			var _coll59 = _dafny.NewMapBuilder()
			_ = _coll59
			for _iter61 := _dafny.Iterate((_1615_v25).Keys().Elements()); ; {
				_compr_59, _ok61 := _iter61()
				if !_ok61 {
					break
				}
				var _1623_v24 _dafny.CodePoint
				_1623_v24 = interface{}(_compr_59).(_dafny.CodePoint)
				if (_1615_v25).Contains(_1623_v24) {
					_coll59.Add(_1623_v24, _1588_v4)
				}
			}
			return _coll59.ToMap()
		}()
		_ = _nwElement0_64
		var _nw250 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(23))
		_ = _nw250
		(_nw250).ArraySet1(_nwElement0_64, 0)
		(_nw250).ArraySet1(func() _dafny.Map {
			var _coll60 = _dafny.NewMapBuilder()
			_ = _coll60
			for _iter62 := _dafny.Iterate((_1616_v27).Elements()); ; {
				_compr_60, _ok62 := _iter62()
				if !_ok62 {
					break
				}
				var _1624_v26 _dafny.CodePoint
				_1624_v26 = interface{}(_compr_60).(_dafny.CodePoint)
				if (_1616_v27).Contains(_1624_v26) {
					_coll60.Add(_1624_v26, _1588_v4)
				}
			}
			return _coll60.ToMap()
		}(), 1)
		(_nw250).ArraySet1((func() _dafny.Map {
			if (_1593_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))).Int()).(bool) {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1584_v0, _1588_v4)
			}
			return _1617_v28
		})(), 2)
		(_nw250).ArraySet1(_1617_v28, 3)
		(_nw250).ArraySet1(func() _dafny.Map {
			var _coll61 = _dafny.NewMapBuilder()
			_ = _coll61
			for _iter63 := _dafny.Iterate((_1619_v31).Keys().Elements()); ; {
				_compr_61, _ok63 := _iter63()
				if !_ok63 {
					break
				}
				var _1625_v29 _dafny.CodePoint
				_1625_v29 = interface{}(_compr_61).(_dafny.CodePoint)
				if (_1619_v31).Contains(_1625_v29) {
					_coll61.Add(_1625_v29, _1588_v4)
				}
			}
			return _coll61.ToMap()
		}(), 4)
		(_nw250).ArraySet1((func() _dafny.Map {
			if (_1620_v32).Contains((func() _dafny.Int {
				if (_1621_v33).Contains(_1587_v3) {
					return (_1621_v33).Get(_1587_v3).(_dafny.Int)
				}
				return _1587_v3
			})()) {
				return (_1620_v32).Get((func() _dafny.Int {
					if (_1621_v33).Contains(_1587_v3) {
						return (_1621_v33).Get(_1587_v3).(_dafny.Int)
					}
					return _1587_v3
				})()).(_dafny.Map)
			}
			return _1617_v28
		})(), 5)
		(_nw250).ArraySet1(_1617_v28, 6)
		(_nw250).ArraySet1(_1617_v28, 7)
		(_nw250).ArraySet1(_1617_v28, 8)
		(_nw250).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1584_v0, _1588_v4), 9)
		(_nw250).ArraySet1(_1617_v28, 10)
		(_nw250).ArraySet1(_1617_v28, 11)
		(_nw250).ArraySet1(_1617_v28, 12)
		(_nw250).ArraySet1(_1617_v28, 13)
		(_nw250).ArraySet1(_1617_v28, 14)
		(_nw250).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm51(false, _1587_v3, _1587_v3, globalState), _dafny.SeqOf(!((_1593_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))).Int()).(bool)))), 15)
		(_nw250).ArraySet1((_1617_v28).Merge(_1617_v28), 16)
		(_nw250).ArraySet1((_1617_v28).Merge(_1617_v28), 17)
		(_nw250).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1584_v0, _1588_v4), 18)
		(_nw250).ArraySet1(_1617_v28, 19)
		(_nw250).ArraySet1(_1617_v28, 20)
		(_nw250).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1584_v0, _1588_v4), 21)
		(_nw250).ArraySet1((func() _dafny.Map {
			var _coll62 = _dafny.NewMapBuilder()
			_ = _coll62
			for _iter64 := _dafny.Iterate((_dafny.SeqOf(_1584_v0)).Elements()); ; {
				_compr_62, _ok64 := _iter64()
				if !_ok64 {
					break
				}
				var _1626_v34 _dafny.CodePoint
				_1626_v34 = interface{}(_compr_62).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_1584_v0), _1626_v34) {
					_coll62.Add(_1626_v34, _1588_v4)
				}
			}
			return _coll62.ToMap()
		}()).Merge((_1617_v28).Update(Companion_Default___.Fm51((_1593_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))).Int()).(bool), _1587_v3, _dafny.IntOfUint32((_dafny.SeqOf(_1587_v3)).Cardinality()), globalState), _1588_v4)), 22)
		_1622_v35 = _nw250
		var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_1622_v35), 0))
		_ = _index260
		(_1622_v35).ArraySet1(func() _dafny.Map {
			var _coll63 = _dafny.NewMapBuilder()
			_ = _coll63
			for _iter65 := _dafny.Iterate((_1617_v28).Keys().Elements()); ; {
				_compr_63, _ok65 := _iter65()
				if !_ok65 {
					break
				}
				var _1627_v36 _dafny.CodePoint
				_1627_v36 = interface{}(_compr_63).(_dafny.CodePoint)
				if (_1617_v28).Contains(_1627_v36) {
					_coll63.Add(_1627_v36, _dafny.SeqOf(_1586_v2, (_1593_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1593_v9), 0))).Int()).(bool)))
				}
			}
			return _coll63.ToMap()
		}(), (_index260).Int())
		var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_1622_v35), 0))
		_ = _index261
		(_1622_v35).ArraySet1(((_1617_v28).Merge(func() _dafny.Map {
			var _coll64 = _dafny.NewMapBuilder()
			_ = _coll64
			for _iter66 := _dafny.Iterate((_1585_v1).Elements()); ; {
				_compr_64, _ok66 := _iter66()
				if !_ok66 {
					break
				}
				var _1628_v37 _dafny.CodePoint
				_1628_v37 = interface{}(_compr_64).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_1585_v1, _1628_v37) {
					_coll64.Add(_1628_v37, _1588_v4)
				}
			}
			return _coll64.ToMap()
		}())).Merge(_1617_v28), (_index261).Int())
		var _1629_v38 _dafny.Array
		_ = _1629_v38
		var _nw251 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
		_ = _nw251
		_1629_v38 = _nw251
		var _1630_v39 *C1
		_ = _1630_v39
		var _nw252 *C1 = New_C1_()
		_ = _nw252
		_nw252.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1585_v1, _1585_v1)).Cardinality()), _1629_v38, _1596_v10, _1590_v6)
		_1630_v39 = _nw252
		var _1631_v40 _dafny.Array
		_ = _1631_v40
		var _len0_35 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_35
		var _nw253 _dafny.Array
		_ = _nw253
		if _len0_35.Cmp(_dafny.Zero) == 0 {
			_nw253 = _dafny.NewArray(_len0_35)
		} else {
			var _init35 func(_dafny.Int) D0 = (func(_1632_v6 D0) func(_dafny.Int) D0 {
				return func(_1633_i3 _dafny.Int) D0 {
					return _1632_v6
				}
			})(_1590_v6)
			_ = _init35
			var _element0_35 = _init35(_dafny.Zero)
			_ = _element0_35
			_nw253 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
			(_nw253).ArraySet1(_element0_35, 0)
			var _nativeLen0_35 = (_len0_35).Int()
			_ = _nativeLen0_35
			for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
				(_nw253).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
			}
		}
		_1631_v40 = _nw253
		r0 = _1631_v40
		var _1634_v41 _dafny.Array
		_ = _1634_v41
		var _len0_36 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_36
		var _nw254 _dafny.Array
		_ = _nw254
		if _len0_36.Cmp(_dafny.Zero) == 0 {
			_nw254 = _dafny.NewArray(_len0_36)
		} else {
			var _init36 func(_dafny.Int) _dafny.Int = (func(_1635_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1636_i4 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1636_i4, (_dafny.Zero).Minus(_1635_v3))
				}
			})(_1587_v3)
			_ = _init36
			var _element0_36 = _init36(_dafny.Zero)
			_ = _element0_36
			_nw254 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
			(_nw254).ArraySet1(_element0_36, 0)
			var _nativeLen0_36 = (_len0_36).Int()
			_ = _nativeLen0_36
			for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
				(_nw254).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
			}
		}
		_1634_v41 = _nw254
		var _1637_v42 D13
		_ = _1637_v42
		_1637_v42 = Companion_D13_.Create_DC32_(_1585_v1, _1586_v2, (_dafny.MultiSetOf(_1586_v2, false)).Cardinality(), _1634_v41)
		var _1638_v43 _dafny.Sequence
		_ = _1638_v43
		_1638_v43 = _dafny.SeqOf(_1585_v1)
		var _1639_v44 _dafny.Map
		_ = _1639_v44
		_1639_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_1638_v43).Select((Companion_Default___.SafeIndex(_1587_v3, _dafny.IntOfUint32((_1638_v43).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _1634_v41)
		var _1640_v45 _dafny.Array
		_ = _1640_v45
		var _nwElement0_65 _dafny.Array = _1634_v41
		_ = _nwElement0_65
		var _nw255 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(27))
		_ = _nw255
		(_nw255).ArraySet1(_nwElement0_65, 0)
		(_nw255).ArraySet1((func() _dafny.Array {
			if _1586_v2 {
				return _1634_v41
			}
			return _1634_v41
		})(), 1)
		(_nw255).ArraySet1(_1634_v41, 2)
		(_nw255).ArraySet1(_1634_v41, 3)
		(_nw255).ArraySet1(_1634_v41, 4)
		(_nw255).ArraySet1(_1634_v41, 5)
		(_nw255).ArraySet1((_1637_v42).Dtor_cf60(), 6)
		(_nw255).ArraySet1(_1634_v41, 7)
		(_nw255).ArraySet1(_1634_v41, 8)
		(_nw255).ArraySet1(_1634_v41, 9)
		(_nw255).ArraySet1(_1634_v41, 10)
		(_nw255).ArraySet1(_1634_v41, 11)
		(_nw255).ArraySet1(_1634_v41, 12)
		(_nw255).ArraySet1(_1634_v41, 13)
		(_nw255).ArraySet1(_1634_v41, 14)
		(_nw255).ArraySet1(_1634_v41, 15)
		(_nw255).ArraySet1(_1634_v41, 16)
		(_nw255).ArraySet1(_1634_v41, 17)
		(_nw255).ArraySet1(_1634_v41, 18)
		(_nw255).ArraySet1(_1634_v41, 19)
		(_nw255).ArraySet1(_1634_v41, 20)
		(_nw255).ArraySet1((func() _dafny.Array {
			if (_1639_v44).Contains(_1587_v3) {
				return (_1639_v44).Get(_1587_v3).(_dafny.Array)
			}
			return _1634_v41
		})(), 21)
		(_nw255).ArraySet1(_1634_v41, 22)
		(_nw255).ArraySet1(_1634_v41, 23)
		(_nw255).ArraySet1(_1634_v41, 24)
		(_nw255).ArraySet1(_1634_v41, 25)
		(_nw255).ArraySet1(_1634_v41, 26)
		_1640_v45 = _nw255
		r1 = _1640_v45
		r2 = (_dafny.IntOfUint32((_1585_v1).Cardinality())).Cmp(_dafny.IntOfInt64(396)) > 0
		return r0, r1, r2
	}
}
func (_this *C4) M10(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Array, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1641_v0 _dafny.CodePoint
		_ = _1641_v0
		_1641_v0 = _dafny.CodePoint('u')
		var _1642_v1 _dafny.Map
		_ = _1642_v1
		_1642_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1641_v0, p1)
		var _1643_v2 _dafny.Map
		_ = _1643_v2
		_1643_v2 = _1642_v1
		_1643_v2 = _1642_v1
		var _hi11 _dafny.Int = p0
		_ = _hi11
		for _1644_i0 := _dafny.IntOfInt64(-403); _1644_i0.Cmp(_hi11) < 0; _1644_i0 = _1644_i0.Plus(_dafny.One) {
			var _1645_v3 _dafny.Map
			_ = _1645_v3
			_1645_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
			var _1646_v4 _dafny.MultiSet
			_ = _1646_v4
			_1646_v4 = _dafny.MultiSetOf(_1644_i0)
			var _1647_v5 D0
			_ = _1647_v5
			_1647_v5 = Companion_D0_.Create_DC1_(p1)
			var _1648_v6 _dafny.Map
			_ = _1648_v6
			_1648_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1641_v0, _1647_v5)
			var _1649_v7 _dafny.Sequence
			_ = _1649_v7
			_1649_v7 = _dafny.SeqOf(!((_this).Fm5(_1644_i0, p2, _1648_v6, p2, globalState)))
			var _1650_v8 _dafny.Sequence
			_ = _1650_v8
			_1650_v8 = _dafny.SeqOf((_1645_v3).Cardinality())
			var _1651_v9 _dafny.Array
			_ = _1651_v9
			var _nw256 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw256
			_1651_v9 = _nw256
			var _1652_v10 _dafny.Map
			_ = _1652_v10
			_1652_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_1651_v9)).Cardinality()), p1)
			var _1653_v11 _dafny.Sequence
			_ = _1653_v11
			_1653_v11 = _dafny.SeqOf(_1646_v4)
			var _1654_v12 _dafny.Array
			_ = _1654_v12
			var _nwElement0_66 _dafny.Int = (_dafny.Zero).Minus(((func() _dafny.Map {
				if p2 {
					return _1645_v3
				}
				return _1645_v3
			})()).Cardinality())
			_ = _nwElement0_66
			var _nw257 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(23))
			_ = _nw257
			(_nw257).ArraySet1(_nwElement0_66, 0)
			(_nw257).ArraySet1((_1646_v4).Cardinality(), 1)
			(_nw257).ArraySet1((p0).Times(_dafny.IntOfUint32((_1649_v7).Cardinality())), 2)
			(_nw257).ArraySet1(p0, 3)
			(_nw257).ArraySet1(_dafny.IntOfUint32((_1650_v8).Cardinality()), 4)
			(_nw257).ArraySet1(_dafny.IntOfInt64(205), 5)
			(_nw257).ArraySet1(_1644_i0, 6)
			(_nw257).ArraySet1((_dafny.Zero).Minus(p0), 7)
			(_nw257).ArraySet1(p0, 8)
			(_nw257).ArraySet1(p1, 9)
			(_nw257).ArraySet1(p1, 10)
			(_nw257).ArraySet1(p1, 11)
			(_nw257).ArraySet1(p1, 12)
			(_nw257).ArraySet1(_dafny.IntOfInt64(-828), 13)
			(_nw257).ArraySet1((func() _dafny.Int {
				if (_1652_v10).Contains((_dafny.Zero).Minus(_1644_i0)) {
					return (_1652_v10).Get((_dafny.Zero).Minus(_1644_i0)).(_dafny.Int)
				}
				return Companion_Default___.Fm20(false, globalState)
			})(), 14)
			(_nw257).ArraySet1(Companion_Default___.SafeDivisionInt(p0, _1644_i0), 15)
			(_nw257).ArraySet1(_dafny.IntOfInt64(892), 16)
			(_nw257).ArraySet1(_1644_i0, 17)
			(_nw257).ArraySet1((p1).Times(_1644_i0), 18)
			(_nw257).ArraySet1(p1, 19)
			(_nw257).ArraySet1(p1, 20)
			(_nw257).ArraySet1(_1644_i0, 21)
			(_nw257).ArraySet1((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(779), _dafny.IntOfUint32((_1653_v11).Cardinality()))).Cardinality()), 22)
			_1654_v12 = _nw257
			var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_1654_v12), 0))
			_ = _index262
			(_1654_v12).ArraySet1(p1, (_index262).Int())
			var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_1654_v12), 0))
			_ = _index263
			(_1654_v12).ArraySet1((_dafny.Zero).Minus(p1), (_index263).Int())
			_1641_v0 = _1641_v0
			var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_1654_v12), 0))
			_ = _index264
			(_1654_v12).ArraySet1((_dafny.Zero).Minus(p0), (_index264).Int())
			var _1655_v13 _dafny.MultiSet
			_ = _1655_v13
			_1655_v13 = _dafny.MultiSetOf(p2)
			r0 = (_1655_v13).IsSubsetOf((_1655_v13).Intersection(_1655_v13))
		}
		var _1656_v14 _dafny.Array
		_ = _1656_v14
		var _len0_37 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_37
		var _nw258 _dafny.Array
		_ = _nw258
		if _len0_37.Cmp(_dafny.Zero) == 0 {
			_nw258 = _dafny.NewArray(_len0_37)
		} else {
			var _init37 func(_dafny.Int) _dafny.MultiSet = (func(_1657_p0 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
				return func(_1658_i1 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf((_dafny.Zero).Minus(_1657_p0))
				}
			})(p0)
			_ = _init37
			var _element0_37 = _init37(_dafny.Zero)
			_ = _element0_37
			_nw258 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
			(_nw258).ArraySet1(_element0_37, 0)
			var _nativeLen0_37 = (_len0_37).Int()
			_ = _nativeLen0_37
			for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
				(_nw258).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
			}
		}
		_1656_v14 = _nw258
		var _1659_v15 _dafny.MultiSet
		_ = _1659_v15
		_1659_v15 = _dafny.MultiSetOf(p0)
		var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1656_v14), 0))
		_ = _index265
		(_1656_v14).ArraySet1(_1659_v15, (_index265).Int())
		var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_1656_v14), 0))
		_ = _index266
		(_1656_v14).ArraySet1((func() _dafny.MultiSet {
			if p2 {
				return _1659_v15
			}
			return (_1659_v15).Difference(_1659_v15)
		})(), (_index266).Int())
		var _1660_v16 _dafny.Sequence
		_ = _1660_v16
		_1660_v16 = _dafny.SeqOf(_1641_v0)
		var _1661_v17 _dafny.Array
		_ = _1661_v17
		var _len0_38 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_38
		var _nw259 _dafny.Array
		_ = _nw259
		if _len0_38.Cmp(_dafny.Zero) == 0 {
			_nw259 = _dafny.NewArray(_len0_38)
		} else {
			var _init38 func(_dafny.Int) _dafny.Map = (func(_1662_v1 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_1663_i2 _dafny.Int) _dafny.Map {
					return _1662_v1
				}
			})(_1642_v1)
			_ = _init38
			var _element0_38 = _init38(_dafny.Zero)
			_ = _element0_38
			_nw259 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
			(_nw259).ArraySet1(_element0_38, 0)
			var _nativeLen0_38 = (_len0_38).Int()
			_ = _nativeLen0_38
			for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
				(_nw259).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
			}
		}
		_1661_v17 = _nw259
		var _1664_v18 _dafny.Set
		_ = _1664_v18
		_1664_v18 = _dafny.SetOf(p3, p3, p3)
		var _1665_v19 _dafny.Sequence
		_ = _1665_v19
		_1665_v19 = _dafny.SeqOf(p0, (_1664_v18).Cardinality(), p0, p1, p1)
		var _1666_v20 *C1
		_ = _1666_v20
		var _nw260 *C1 = New_C1_()
		_ = _nw260
		_nw260.Ctor__((p1).Plus(_dafny.IntOfUint32((_1660_v16).Cardinality())), _1661_v17, _dafny.Companion_Sequence_.Concatenate(_1665_v19, _dafny.Companion_Sequence_.Update(_1665_v19, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1665_v19).Cardinality()))).Uint32(), _dafny.IntOfInt64(661))), Companion_D0_.Create_DC1_(p1))
		_1666_v20 = _nw260
		var _1667_v21 *C2
		_ = _1667_v21
		var _nw261 *C2 = New_C2_()
		_ = _nw261
		_nw261.Ctor__()
		_1667_v21 = _nw261
		var _1668_v22 D0
		_ = _1668_v22
		_1668_v22 = Companion_D0_.Create_DC1_(_1666_v20.F9)
		var _1669_v23 _dafny.Sequence
		_ = _1669_v23
		_1669_v23 = _dafny.SeqOf(_1668_v22)
		var _1670_v24 _dafny.Map
		_ = _1670_v24
		_1670_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(320), _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_1669_v23, (Companion_Default___.SafeIndex(_1666_v20.F9, _dafny.IntOfUint32((_1669_v23).Cardinality()))).Uint32(), _1668_v22), _1669_v23))
		_1670_v24 = _1670_v24
		r0 = p2
		r1 = _1666_v20.F9
		return r0, r1
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f1 D0
	_f0 _dafny.Sequence
	_f7 _dafny.Array
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f1 = Companion_D0_.Default()
	_this._f0 = _dafny.EmptySeq
	_this._f7 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_}
}

var _ T1 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F1() D0 {
	return _this._f1
}
func (_this *C5) F1_set_(value D0) {
	_this._f1 = value
}
func (_this *C5) F0() _dafny.Sequence {
	return _this._f0
}
func (_this *C5) Ctor__(f7 _dafny.Array, f0 _dafny.Sequence, f1 D0) {
	{
		(_this)._f7 = f7
		(_this)._f0 = f0
		(_this)._f1 = f1
	}
}
func (_this *C5) Fm6(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(487)
	}
}
func (_this *C5) M7(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Array, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Set, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Set = _dafny.EmptySet
		_ = r1
		var r2 bool = false
		_ = r2
		var _1671_v0 _dafny.Sequence
		_ = _1671_v0
		_1671_v0 = _dafny.UnicodeSeqOfUtf8Bytes("ibpiei")
		var _1672_v1 _dafny.CodePoint
		_ = _1672_v1
		_1672_v1 = _dafny.CodePoint('c')
		var _1673_v2 *C0
		_ = _1673_v2
		var _nw262 *C0 = New_C0_()
		_ = _nw262
		_nw262.Ctor__(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm16(_1671_v0, p0, globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm16(_1671_v0, p0, globalState)).Cardinality()))).Uint32(), _1672_v1))
		_1673_v2 = _nw262
		if p3 {
			var _1674_v3 _dafny.Set
			_ = _1674_v3
			_1674_v3 = _dafny.SetOf(p3)
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
			_ = _index267
			((_this).F7()).ArraySet1(((_this).F0()).Select((Companion_Default___.SafeIndex((_1674_v3).Cardinality(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int), (_index267).Int())
			var _1675_v4 _dafny.Map
			_ = _1675_v4
			_1675_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
			_ = _index268
			((_this).F7()).ArraySet1(((func() _dafny.Int {
				if (_1675_v4).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wsieh")).Cardinality())) {
					return (_1675_v4).Get(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wsieh")).Cardinality())).(_dafny.Int)
				}
				return p0
			})()).Minus(_dafny.IntOfInt64(333)), (_index268).Int())
			if Companion_Default___.Fm17(globalState) {
				r2 = p3
				var _1676_v5 _dafny.Map
				_ = _1676_v5
				_1676_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int), p3)
				var _1677_v6 _dafny.Sequence
				_ = _1677_v6
				_1677_v6 = _dafny.SeqOf(p3, (func() bool {
					if (_1676_v5).Contains(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int)) {
						return (_1676_v5).Get(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int)).(bool)
					}
					return p3
				})())
				var _1678_v7 _dafny.Array
				_ = _1678_v7
				var _len0_39 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_39
				var _nw263 _dafny.Array
				_ = _nw263
				if _len0_39.Cmp(_dafny.Zero) == 0 {
					_nw263 = _dafny.NewArray(_len0_39)
				} else {
					var _init39 func(_dafny.Int) bool = (func(_1679_p3 bool) func(_dafny.Int) bool {
						return func(_1680_i0 _dafny.Int) bool {
							return _1679_p3
						}
					})(p3)
					_ = _init39
					var _element0_39 = _init39(_dafny.Zero)
					_ = _element0_39
					_nw263 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
					(_nw263).ArraySet1(_element0_39, 0)
					var _nativeLen0_39 = (_len0_39).Int()
					_ = _nativeLen0_39
					for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
						(_nw263).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
					}
				}
				_1678_v7 = _nw263
				var _1681_v8 _dafny.Map
				_ = _1681_v8
				_1681_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1678_v7, _1677_v6)
				_1677_v6 = _dafny.Companion_Sequence_.Concatenate(_1677_v6, (func() _dafny.Sequence {
					if (_1681_v8).Contains(_1678_v7) {
						return (_1681_v8).Get(_1678_v7).(_dafny.Sequence)
					}
					return _1677_v6
				})())
				r0 = Companion_Default___.SafeDivisionInt((_dafny.IntOfInt64(-93)).Plus(_dafny.IntOfInt64(-786)), ((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int))
				var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
				_ = _index269
				var _rhs179 _dafny.Int = p0
				_ = _rhs179
				var _rhs180 _dafny.Int = (((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int)).Minus(p0)
				_ = _rhs180
				var _lhs110 _dafny.Array = (_this).F7()
				_ = _lhs110
				var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
				_ = _lhs111
				(_lhs110).ArraySet1(_rhs179, (_lhs111).Int())
				r0 = _rhs180
				r0 = ((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int)
			} else {
				var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
				_ = _index270
				((_this).F7()).ArraySet1((((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int)).Plus(p0), (_index270).Int())
				var _1682_v9 _dafny.Set
				_ = _1682_v9
				_1682_v9 = _dafny.SetOf(_dafny.SetOf(p3), _1674_v3)
				var _1683_v10 _dafny.Array
				_ = _1683_v10
				var _nw264 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw264
				_1683_v10 = _nw264
				var _rhs181 bool = Companion_Default___.Fm17(globalState)
				_ = _rhs181
				var _rhs182 _dafny.Set = func() _dafny.Set {
					var _coll65 = _dafny.NewBuilder()
					_ = _coll65
					for _iter67 := _dafny.Iterate((Companion_Default___.Fm18(globalState)).Elements()); ; {
						_compr_65, _ok67 := _iter67()
						if !_ok67 {
							break
						}
						var _1684_v11 _dafny.Set
						_1684_v11 = interface{}(_compr_65).(_dafny.Set)
						if (Companion_Default___.Fm18(globalState)).Contains(_1684_v11) {
							_coll65.Add(_1684_v11)
						}
					}
					return _coll65.ToSet()
				}()
				_ = _rhs182
				var _rhs183 _dafny.Array = _1683_v10
				_ = _rhs183
				var _rhs184 bool = p3
				_ = _rhs184
				var _rhs185 bool = !(p3)
				_ = _rhs185
				r2 = _rhs181
				_1682_v9 = _rhs182
				_1683_v10 = _rhs183
				r2 = _rhs184
				r2 = _rhs185
				r2 = p3
				var _1685_v12 _dafny.Map
				_ = _1685_v12
				_1685_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
				_1685_v12 = (_1685_v12).Update((_this).Fm6(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int), p0, globalState), p3)
				_1671_v0 = _dafny.Companion_Sequence_.Concatenate(_1673_v2.F8, _dafny.Companion_Sequence_.Concatenate(_1671_v0, _1673_v2.F8))
			}
			var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
			_ = _index271
			((_this).F7()).ArraySet1(p0, (_index271).Int())
			var _1686_v13 _dafny.Sequence
			_ = _1686_v13
			_1686_v13 = _dafny.SeqOf((_this).F0())
			r2 = !_dafny.Companion_Sequence_.Contains(_1686_v13, (_this).F0())
			if !(p3) {
				_1671_v0 = _1673_v2.F8
				(_this).F1_set_(_this.F1())
				r0 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(202), p0)
				var _1687_v14 _dafny.Map
				_ = _1687_v14
				_1687_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1672_v1, p0)
				var _1688_v15 _dafny.Map
				_ = _1688_v15
				_1688_v15 = _1687_v14
				var _1689_v16 _dafny.Map
				_ = _1689_v16
				_1689_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p3)
				var _1690_v17 _dafny.Map
				_ = _1690_v17
				_1690_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm19(Companion_Default___.Fm17(globalState), p3, (_1688_v15), p3, globalState), (func() bool {
					if (_1689_v16).Contains(p3) {
						return (_1689_v16).Get(p3).(bool)
					}
					return !(false)
				})())
				_1690_v17 = ((_1690_v17).Merge(_1690_v17)).Merge(_1690_v17)
				var _1691_v18 _dafny.MultiSet
				_ = _1691_v18
				_1691_v18 = _dafny.MultiSetOf(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int))
				var _1692_v19 _dafny.MultiSet
				_ = _1692_v19
				_1692_v19 = _dafny.MultiSetOf(p3, p3, p3, !(p3), p3)
				var _1693_v20 _dafny.Sequence
				_ = _1693_v20
				_1693_v20 = _dafny.SeqOf((((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int)).Cmp(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int)) >= 0, p3, (_1692_v19).IsSubsetOf(_1692_v19))
				var _1694_v21 T0
				_ = _1694_v21
				var _nw265 *C4 = New_C4_()
				_ = _nw265
				_nw265.Ctor__()
				_1694_v21 = _nw265
				var _1695_v22 _dafny.Sequence
				_ = _1695_v22
				_1695_v22 = _dafny.SeqOf(_1693_v20)
				var _1696_v23 _dafny.Sequence
				_ = _1696_v23
				_1696_v23 = _dafny.SeqOf((_1695_v22).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1695_v22).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
				_ = _index272
				var _rhs186 _dafny.Int = ((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int)
				_ = _rhs186
				var _rhs187 _dafny.MultiSet = (_1691_v18).Difference(_1691_v18)
				_ = _rhs187
				var _rhs188 _dafny.Sequence = (_1695_v22).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm32(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int), globalState), _dafny.IntOfUint32((_1695_v22).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs188
				var _rhs189 T0 = _1694_v21
				_ = _rhs189
				var _rhs190 _dafny.Int = ((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int)
				_ = _rhs190
				var _lhs112 _dafny.Array = (_this).F7()
				_ = _lhs112
				var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
				_ = _lhs113
				r0 = _rhs186
				_1691_v18 = _rhs187
				_1693_v20 = _rhs188
				_1694_v21 = _rhs189
				(_lhs112).ArraySet1(_rhs190, (_lhs113).Int())
			} else {
				var _1697_v24 *C2
				_ = _1697_v24
				var _nw266 *C2 = New_C2_()
				_ = _nw266
				_nw266.Ctor__()
				_1697_v24 = _nw266
				var _1698_v25 _dafny.Map
				_ = _1698_v25
				_1698_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int), p3)
				var _1699_v26 _dafny.Sequence
				_ = _1699_v26
				_1699_v26 = _dafny.SeqOf(_1698_v25)
				var _1700_v27 _dafny.Map
				_ = _1700_v27
				_1700_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1699_v26).Select((Companion_Default___.SafeIndex(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1699_v26).Cardinality()))).Uint32()).(_dafny.Map), p0)
				var _1701_v28 _dafny.Map
				_ = _1701_v28
				_1701_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1672_v1, _this.F1())
				var _1702_v29 _dafny.Map
				_ = _1702_v29
				_1702_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_1697_v24).Fm5(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int), false, _1701_v28, p3, globalState))
				var _1703_v30 _dafny.Set
				_ = _1703_v30
				_1703_v30 = _dafny.SetOf((_this).F0(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-806))).Uint32(), func(coer97 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg97 _dafny.Int) interface{} {
						return coer97(arg97)
					}
				}(func(_1704_i1 _dafny.Int) _dafny.Int {
					return ((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int)
				})))
				var _1705_v31 _dafny.MultiSet
				_ = _1705_v31
				_1705_v31 = _dafny.MultiSetOf(p3, true, p3, p3)
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
				_ = _index273
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
				_ = _index274
				var _rhs191 _dafny.Int = ((p0).Plus(_dafny.IntOfInt64(174))).Minus((func() _dafny.Int {
					if (_1700_v27).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int), p3)) {
						return (_1700_v27).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int), p3)).(_dafny.Int)
					}
					return (_1702_v29).Cardinality()
				})())
				_ = _rhs191
				var _rhs192 _dafny.Sequence = _1673_v2.F8
				_ = _rhs192
				var _rhs193 bool = !((func() bool {
					if p3 {
						return (Companion_Default___.Fm58(Companion_D16_.Create_DC42_(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int), p0, p0), true, _1703_v30, globalState)).IsSubsetOf(_dafny.MultiSetOf(_1672_v1, _1672_v1))
					}
					return p3
				})())
				_ = _rhs193
				var _rhs194 _dafny.Int = ((_1705_v31).Cardinality()).Times(p0)
				_ = _rhs194
				var _rhs195 bool = p3
				_ = _rhs195
				var _lhs114 _dafny.Array = (_this).F7()
				_ = _lhs114
				var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
				_ = _lhs115
				var _lhs116 _dafny.Array = (_this).F7()
				_ = _lhs116
				var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
				_ = _lhs117
				(_lhs114).ArraySet1(_rhs191, (_lhs115).Int())
				_1671_v0 = _rhs192
				r2 = _rhs193
				(_lhs116).ArraySet1(_rhs194, (_lhs117).Int())
				r2 = _rhs195
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))
				_ = _index275
				((_this).F7()).ArraySet1(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int), (_index275).Int())
				r2 = !_dafny.Companion_Sequence_.Equal(_1671_v0, _1673_v2.F8)
				var _1706_v32 _dafny.Array
				_ = _1706_v32
				var _len0_40 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_40
				var _nw267 _dafny.Array
				_ = _nw267
				if _len0_40.Cmp(_dafny.Zero) == 0 {
					_nw267 = _dafny.NewArray(_len0_40)
				} else {
					var _init40 func(_dafny.Int) bool = (func(_1707_p3 bool) func(_dafny.Int) bool {
						return func(_1708_i2 _dafny.Int) bool {
							return (_1707_p3) || (false)
						}
					})(p3)
					_ = _init40
					var _element0_40 = _init40(_dafny.Zero)
					_ = _element0_40
					_nw267 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
					(_nw267).ArraySet1(_element0_40, 0)
					var _nativeLen0_40 = (_len0_40).Int()
					_ = _nativeLen0_40
					for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
						(_nw267).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
					}
				}
				_1706_v32 = _nw267
				var _1709_v33 _dafny.Map
				_ = _1709_v33
				_1709_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1672_v1, _dafny.IntOfInt64(583))
				var _1710_v34 _dafny.Set
				_ = _1710_v34
				_1710_v34 = _dafny.SetOf(_1709_v33, _1709_v33)
				var _1711_v36 _dafny.Map
				_ = _1711_v36
				_1711_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1672_v1, true)
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_1706_v32), 0))
				_ = _index276
				(_1706_v32).ArraySet1((_1710_v34).Equals(_dafny.SetOf(_1709_v33, (func() _dafny.Map {
					var _coll66 = _dafny.NewMapBuilder()
					_ = _coll66
					for _iter68 := _dafny.Iterate((_1711_v36).Keys().Elements()); ; {
						_compr_66, _ok68 := _iter68()
						if !_ok68 {
							break
						}
						var _1712_v35 _dafny.CodePoint
						_1712_v35 = interface{}(_compr_66).(_dafny.CodePoint)
						if (_1711_v36).Contains(_1712_v35) {
							_coll66.Add(_1712_v35, ((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int))
						}
					}
					return _coll66.ToMap()
				}()).Update(_1672_v1, (_this).Fm6(((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int), ((_this).F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen(((_this).F7()), 0))).Int()).(_dafny.Int), globalState)))), (_index276).Int())
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_1706_v32), 0))
				_ = _index277
				(_1706_v32).ArraySet1(p3, (_index277).Int())
			}
		} else {
			r2 = (p0).Cmp(p0) >= 0
			var _1713_v37 _dafny.Sequence
			_ = _1713_v37
			_1713_v37 = _dafny.SeqOf(_1673_v2.F8)
			var _1714_v38 _dafny.Sequence
			_ = _1714_v38
			_1714_v38 = _dafny.SeqOf((_1713_v37).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1713_v37).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _1715_v39 _dafny.Map
			_ = _1715_v39
			_1715_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.MultiSetFromSeq(_dafny.SeqOf(p0, p0))).Cardinality()).Minus(_dafny.IntOfUint32((_1713_v37).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(((_this).F0()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F0()).Cardinality()), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int), p0, p0, p0, (p1).Cardinality()), (_this).F0()))
			_1715_v39 = (_1715_v39).Update(p0, (_this).F0())
			r2 = p3
			r2 = !((p0).Cmp((_dafny.Zero).Minus(p0)) != 0)
			var _1716_v40 _dafny.Set
			_ = _1716_v40
			_1716_v40 = _dafny.SetOf(p0)
			_1716_v40 = _1716_v40
		}
		var _1717_v41 _dafny.Sequence
		_ = _1717_v41
		_1717_v41 = _dafny.SeqOf(p3)
		var _rhs196 bool = p3
		_ = _rhs196
		var _rhs197 bool = (func() bool {
			if p3 {
				return p3
			}
			return !(_dafny.Companion_Sequence_.IsProperPrefixOf(_1717_v41, Companion_Default___.Fm19(p3, !(false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1672_v1, _dafny.IntOfUint32((_1673_v2.F8).Cardinality())), p3, globalState)))
		})()
		_ = _rhs197
		var _rhs198 _dafny.Int = p0
		_ = _rhs198
		var _rhs199 _dafny.Int = (_dafny.Zero).Minus((p0).Times(p0))
		_ = _rhs199
		r2 = _rhs196
		r2 = _rhs197
		r0 = _rhs198
		r0 = _rhs199
		var _1718_v42 _dafny.Map
		_ = _1718_v42
		_1718_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), p3)
		_1718_v42 = (_1718_v42).Merge(_1718_v42)
		var _1719_v43 D12
		_ = _1719_v43
		_1719_v43 = Companion_D12_.Create_DC28_(p3, p0)
		_1719_v43 = _1719_v43
		var _hi12 _dafny.Int = p0
		_ = _hi12
		for _1720_i3 := p0; _1720_i3.Cmp(_hi12) < 0; _1720_i3 = _1720_i3.Plus(_dafny.One) {
			(_1673_v2).F8 = _1671_v0
			r0 = p0
			var _1721_v44 _dafny.Map
			_ = _1721_v44
			_1721_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
			_1721_v44 = (_1721_v44).Update((p0).Cmp(p0) < 0, p3)
			var _1722_v45 D11
			_ = _1722_v45
			_1722_v45 = Companion_D11_.Create_DC25_(_1672_v1, _dafny.IntOfUint32((_1673_v2.F8).Cardinality()), true)
			var _1723_v46 _dafny.Array
			_ = _1723_v46
			var _nwElement0_67 bool = p3
			_ = _nwElement0_67
			var _nw268 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(21))
			_ = _nw268
			(_nw268).ArraySet1(_nwElement0_67, 0)
			(_nw268).ArraySet1(p3, 1)
			(_nw268).ArraySet1(p3, 2)
			(_nw268).ArraySet1(p3, 3)
			(_nw268).ArraySet1(p3, 4)
			(_nw268).ArraySet1(p3, 5)
			(_nw268).ArraySet1(p3, 6)
			(_nw268).ArraySet1(p3, 7)
			(_nw268).ArraySet1(true, 8)
			(_nw268).ArraySet1(p3, 9)
			(_nw268).ArraySet1(p3, 10)
			(_nw268).ArraySet1((_1722_v45).Dtor_cf43(), 11)
			(_nw268).ArraySet1(p3, 12)
			(_nw268).ArraySet1(p3, 13)
			(_nw268).ArraySet1(p3, 14)
			(_nw268).ArraySet1(!(p3), 15)
			(_nw268).ArraySet1(p3, 16)
			(_nw268).ArraySet1(p3, 17)
			(_nw268).ArraySet1(p3, 18)
			(_nw268).ArraySet1(true, 19)
			(_nw268).ArraySet1(p3, 20)
			_1723_v46 = _nw268
			var _1724_v47 _dafny.Set
			_ = _1724_v47
			_1724_v47 = _dafny.SetOf(_1723_v46)
			var _1725_v48 _dafny.MultiSet
			_ = _1725_v48
			_1725_v48 = _dafny.MultiSetOf((_1724_v47).Cardinality(), _1720_i3, _dafny.IntOfInt64(239), p0, _1720_i3)
			var _1726_v49 _dafny.Set
			_ = _1726_v49
			_1726_v49 = _dafny.SetOf(p3, (_1725_v48).Equals(_1725_v48))
			_1726_v49 = _1726_v49
		}
		var _1727_v50 _dafny.MultiSet
		_ = _1727_v50
		_1727_v50 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1671_v0).Cardinality()), (_dafny.Zero).Minus(p0))
		r0 = ((_1727_v50).Union((_1727_v50).Union(_1727_v50))).Cardinality()
		r1 = Companion_Default___.Fm59(globalState)
		r2 = true
		return r0, r1, r2
	}
}
func (_this *C5) M8(p0 bool, p1 bool, p2 _dafny.MultiSet, p3 bool, globalState *GlobalState) (bool, _dafny.Map, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var _1728_v0 _dafny.Sequence
		_ = _1728_v0
		_1728_v0 = _dafny.UnicodeSeqOfUtf8Bytes("fi")
		var _1729_v1 *C0
		_ = _1729_v1
		var _nw269 *C0 = New_C0_()
		_ = _nw269
		_nw269.Ctor__(_1728_v0)
		_1729_v1 = _nw269
		var _1730_v2 _dafny.CodePoint
		_ = _1730_v2
		_1730_v2 = _dafny.CodePoint('j')
		var _1731_v3 _dafny.Int
		_ = _1731_v3
		_1731_v3 = _dafny.IntOfInt64(-151)
		var _1732_v4 D16
		_ = _1732_v4
		_1732_v4 = Companion_D16_.Create_DC43_(_dafny.UnicodeSeqOfUtf8Bytes("ekkqo"), _1730_v2, _1731_v3)
		var _pat_let_tv42 = p1
		_ = _pat_let_tv42
		var _pat_let_tv43 = p1
		_ = _pat_let_tv43
		var _pat_let_tv44 = p1
		_ = _pat_let_tv44
		var _pat_let_tv45 = p3
		_ = _pat_let_tv45
		r0 = !(func(_source18 D16) bool {
			if _source18.Is_DC42() {
				var _1733___mcc_h0 _dafny.Int = _source18.Get_().(D16_DC42).Cf76
				_ = _1733___mcc_h0
				var _1734___mcc_h1 _dafny.Int = _source18.Get_().(D16_DC42).Cf77
				_ = _1734___mcc_h1
				var _1735___mcc_h2 _dafny.Int = _source18.Get_().(D16_DC42).Cf78
				_ = _1735___mcc_h2
				var _1736_cf78 _dafny.Int = _1735___mcc_h2
				_ = _1736_cf78
				var _1737_cf77 _dafny.Int = _1734___mcc_h1
				_ = _1737_cf77
				var _1738_cf76 _dafny.Int = _1733___mcc_h0
				_ = _1738_cf76
				return _pat_let_tv42
			} else if _source18.Is_DC43() {
				var _1739___mcc_h3 _dafny.Sequence = _source18.Get_().(D16_DC43).Cf79
				_ = _1739___mcc_h3
				var _1740___mcc_h4 _dafny.CodePoint = _source18.Get_().(D16_DC43).Cf80
				_ = _1740___mcc_h4
				var _1741___mcc_h5 _dafny.Int = _source18.Get_().(D16_DC43).Cf81
				_ = _1741___mcc_h5
				var _1742_cf81 _dafny.Int = _1741___mcc_h5
				_ = _1742_cf81
				var _1743_cf80 _dafny.CodePoint = _1740___mcc_h4
				_ = _1743_cf80
				var _1744_cf79 _dafny.Sequence = _1739___mcc_h3
				_ = _1744_cf79
				return _pat_let_tv43
			} else if _source18.Is_DC44() {
				var _1745___mcc_h6 _dafny.Int = _source18.Get_().(D16_DC44).Cf82
				_ = _1745___mcc_h6
				var _1746_cf82 _dafny.Int = _1745___mcc_h6
				_ = _1746_cf82
				return _pat_let_tv44
			} else {
				var _1747___mcc_h7 _dafny.Array = _source18.Get_().(D16_DC41).Cf75
				_ = _1747___mcc_h7
				var _1748_cf75 _dafny.Array = _1747___mcc_h7
				_ = _1748_cf75
				return _pat_let_tv45
			}
		}(_1732_v4))
		var _hi13 _dafny.Int = Companion_Default___.Fm44(p0, globalState)
		_ = _hi13
		for _1749_i0 := (_dafny.Zero).Minus(_dafny.IntOfUint32((_1728_v0).Cardinality())); _1749_i0.Cmp(_hi13) < 0; _1749_i0 = _1749_i0.Plus(_dafny.One) {
			var _1750_v5 *C0
			_ = _1750_v5
			var _nw270 *C0 = New_C0_()
			_ = _nw270
			_nw270.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("m"))
			_1750_v5 = _nw270
			_1731_v3 = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if p1 {
					return _1750_v5.F8
				}
				return _dafny.Companion_Sequence_.Concatenate(_1729_v1.F8, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(386))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg98 _dafny.Int) interface{} {
						return coer98(arg98)
					}
				}((func(_1751_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1752_i1 _dafny.Int) _dafny.CodePoint {
						return _1751_v2
					}
				})(_1730_v2))))
			})()).Cardinality())
			r0 = false
			var _1753_v6 _dafny.Sequence
			_ = _1753_v6
			_1753_v6 = _dafny.SeqOf(true)
			var _1754_v7 D6
			_ = _1754_v7
			_1754_v7 = Companion_D6_.Create_DC15_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-836))).Uint32(), func(coer99 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg99 _dafny.Int) interface{} {
					return coer99(arg99)
				}
			}((func(_1755_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1756_i2 _dafny.Int) _dafny.CodePoint {
					return _1755_v2
				}
			})(_1730_v2)))).Cardinality()), _dafny.IntOfUint32((_1753_v6).Cardinality()), p3)
			var _1757_v8 _dafny.Map
			_ = _1757_v8
			_1757_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1754_v7)
			var _source19 D6 = Companion_D6_.Create_DC16_((func() D6 {
				if (_1757_v8).Contains(p3) {
					return (_1757_v8).Get(p3).(D6)
				}
				return _1754_v7
			})())
			_ = _source19
			if _source19.Is_DC15() {
				var _1758___mcc_h8 _dafny.Int = _source19.Get_().(D6_DC15).Cf30
				_ = _1758___mcc_h8
				var _1759___mcc_h9 _dafny.Int = _source19.Get_().(D6_DC15).Cf31
				_ = _1759___mcc_h9
				var _1760___mcc_h10 bool = _source19.Get_().(D6_DC15).Cf32
				_ = _1760___mcc_h10
				var _1761_cf32 bool = _1760___mcc_h10
				_ = _1761_cf32
				var _1762_cf31 _dafny.Int = _1759___mcc_h9
				_ = _1762_cf31
				var _1763_cf30 _dafny.Int = _1758___mcc_h8
				_ = _1763_cf30
				var _rhs200 bool = ((_1763_cf30).Times(_1763_cf30)).Cmp((_1762_cf31).Times(_dafny.IntOfInt64(263))) > 0
				_ = _rhs200
				var _rhs201 bool = (true) || (p0)
				_ = _rhs201
				var _rhs202 _dafny.Int = (_dafny.IntOfInt64(191)).Minus(_1731_v3)
				_ = _rhs202
				var _rhs203 _dafny.Int = (_dafny.Zero).Minus(_1763_cf30)
				_ = _rhs203
				r0 = _rhs200
				r0 = _rhs201
				_1731_v3 = _rhs202
				_1763_cf30 = _rhs203
				var _1764_v9 _dafny.Array
				_ = _1764_v9
				var _nw271 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw271
				_1764_v9 = _nw271
				var _1765_v13 _dafny.Array
				_ = _1765_v13
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_41
				var _nw272 _dafny.Array
				_ = _nw272
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw272 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) bool = (func(_1766_i0 _dafny.Int, _1767_cf31 _dafny.Int) func(_dafny.Int) bool {
						return func(_1768_i3 _dafny.Int) bool {
							return !(func() _dafny.Set {
								var _coll67 = _dafny.NewBuilder()
								_ = _coll67
								for _iter69 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(892), _dafny.IntOfInt64(820))); ; {
									_compr_67, _ok69 := _iter69()
									if !_ok69 {
										break
									}
									var _1769_v10 _dafny.Int
									_1769_v10 = interface{}(_compr_67).(_dafny.Int)
									if ((_dafny.IntOfInt64(892)).Cmp(_1769_v10) <= 0) && ((_1769_v10).Cmp(_dafny.IntOfInt64(820)) < 0) {
										_coll67.Add((_1769_v10).Minus(_1766_i0))
									}
								}
								return _coll67.ToSet()
							}()).Equals(func() _dafny.Set {
								var _coll68 = _dafny.NewBuilder()
								_ = _coll68
								for _iter70 := _dafny.Iterate((func() _dafny.Map {
									var _coll69 = _dafny.NewMapBuilder()
									_ = _coll69
									for _iter71 := _dafny.Iterate((_dafny.SetOf(_1766_i0)).Elements()); ; {
										_compr_69, _ok71 := _iter71()
										if !_ok71 {
											break
										}
										var _1770_v11 _dafny.Int
										_1770_v11 = interface{}(_compr_69).(_dafny.Int)
										if (_dafny.SetOf(_1766_i0)).Contains(_1770_v11) {
											_coll69.Add(Companion_Default___.SafeDivisionInt(_1770_v11, _1767_cf31), _dafny.IntOfInt64(49))
										}
									}
									return _coll69.ToMap()
								}()).Keys().Elements()); ; {
									_compr_68, _ok70 := _iter70()
									if !_ok70 {
										break
									}
									var _1771_v12 _dafny.Int
									_1771_v12 = interface{}(_compr_68).(_dafny.Int)
									if (func() _dafny.Map {
										var _coll70 = _dafny.NewMapBuilder()
										_ = _coll70
										for _iter72 := _dafny.Iterate((_dafny.SetOf(_1766_i0)).Elements()); ; {
											_compr_70, _ok72 := _iter72()
											if !_ok72 {
												break
											}
											var _1770_v11 _dafny.Int
											_1770_v11 = interface{}(_compr_70).(_dafny.Int)
											if (_dafny.SetOf(_1766_i0)).Contains(_1770_v11) {
												_coll70.Add(Companion_Default___.SafeDivisionInt(_1770_v11, _1767_cf31), _dafny.IntOfInt64(49))
											}
										}
										return _coll70.ToMap()
									}()).Contains(_1771_v12) {
										_coll68.Add(Companion_Default___.SafeDivisionInt(_1771_v12, _dafny.IntOfInt64(793)))
									}
								}
								return _coll68.ToSet()
							}())
						}
					})(_1749_i0, _1762_cf31)
					_ = _init41
					var _element0_41 = _init41(_dafny.Zero)
					_ = _element0_41
					_nw272 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
					(_nw272).ArraySet1(_element0_41, 0)
					var _nativeLen0_41 = (_len0_41).Int()
					_ = _nativeLen0_41
					for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
						(_nw272).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
					}
				}
				_1765_v13 = _nw272
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_1765_v13), 0))
				_ = _index278
				(_1765_v13).ArraySet1(Companion_Default___.Fm17(globalState), (_index278).Int())
				var _1772_v14 _dafny.Set
				_ = _1772_v14
				_1772_v14 = _dafny.SetOf(_1731_v3, _dafny.IntOfInt64(-459))
				var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_1765_v13), 0))
				_ = _index279
				var _rhs204 bool = (_dafny.IntOfInt64(459)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1761_cf32, !(p0), _1761_cf32), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1772_v14).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_1761_cf32, !(p0), _1761_cf32)).Cardinality()))).Uint32(), p3)).Cardinality())) == 0
				_ = _rhs204
				var _rhs205 _dafny.Int = _1731_v3
				_ = _rhs205
				var _rhs206 _dafny.Array = _1764_v9
				_ = _rhs206
				var _rhs207 bool = p3
				_ = _rhs207
				var _lhs118 _dafny.Array = _1765_v13
				_ = _lhs118
				var _lhs119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_1765_v13), 0))
				_ = _lhs119
				_1761_cf32 = _rhs204
				_1731_v3 = _rhs205
				_1764_v9 = _rhs206
				(_lhs118).ArraySet1(_rhs207, (_lhs119).Int())
				_1761_cf32 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F0()), (Companion_Default___.SafeIndex(Companion_Default___.Fm20((_1765_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_1765_v13), 0))).Int()).(bool), globalState), _dafny.IntOfUint32((_dafny.SeqOf((_this).F0())).Cardinality()))).Uint32(), (_this).F0()), (Companion_Default___.SafeIndex(_1762_cf31, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F0()), (Companion_Default___.SafeIndex(Companion_Default___.Fm20((_1765_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_1765_v13), 0))).Int()).(bool), globalState), _dafny.IntOfUint32((_dafny.SeqOf((_this).F0())).Cardinality()))).Uint32(), (_this).F0())).Cardinality()))).Uint32(), _dafny.SeqOf(_1749_i0)), _dafny.SeqOf(_dafny.SeqOf(_1731_v3), _dafny.Companion_Sequence_.Update((_this).F0(), (Companion_Default___.SafeIndex(_1731_v3, _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(906))).Uint32(), func(coer100 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}((func(_1773_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1774_i4 _dafny.Int) _dafny.CodePoint {
						return _1773_v2
					}
				})(_1730_v2)))).Cardinality())))), _dafny.Companion_Sequence_.Concatenate((_this).F0(), (_this).F0()))
				r0 = _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-884))).Uint32(), func(coer101 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg101 _dafny.Int) interface{} {
						return coer101(arg101)
					}
				}((func(_1775_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1776_i5 _dafny.Int) _dafny.CodePoint {
						return _1775_v2
					}
				})(_1730_v2))), _1730_v2)
			} else if _source19.Is_DC14() {
				var _1777___mcc_h11 _dafny.Sequence = _source19.Get_().(D6_DC14).Cf29
				_ = _1777___mcc_h11
				var _1778_cf29 _dafny.Sequence = _1777___mcc_h11
				_ = _1778_cf29
				r2 = _dafny.Companion_Sequence_.Concatenate((_this).F0(), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_this).F0(), (_this).F0()), (Companion_Default___.SafeIndex(((_this).F0()).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm44(Companion_Default___.Fm17(globalState), globalState), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F0(), (_this).F0())).Cardinality()))).Uint32(), _1731_v3))
				var _1779_v15 *C4
				_ = _1779_v15
				var _nw273 *C4 = New_C4_()
				_ = _nw273
				_nw273.Ctor__()
				_1779_v15 = _nw273
				_1730_v2 = _1730_v2
				var _1780_v16 *C0
				_ = _1780_v16
				var _nw274 *C0 = New_C0_()
				_ = _nw274
				_nw274.Ctor__(_1728_v0)
				_1780_v16 = _nw274
			} else {
				var _1781___mcc_h12 D6 = _source19.Get_().(D6_DC16).Cf33
				_ = _1781___mcc_h12
				var _1782_cf33 D6 = _1781___mcc_h12
				_ = _1782_cf33
				var _1783_v17 _dafny.Array
				_ = _1783_v17
				var _nw275 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
				_ = _nw275
				_1783_v17 = _nw275
				var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_1783_v17), 0))
				_ = _index280
				(_1783_v17).ArraySet1(p1, (_index280).Int())
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_1783_v17), 0))
				_ = _index281
				(_1783_v17).ArraySet1(p3, (_index281).Int())
				_1731_v3 = (_dafny.IntOfInt64(-536)).Plus(_1731_v3)
				var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_1783_v17), 0))
				_ = _index282
				(_1783_v17).ArraySet1((_dafny.SetOf(_dafny.IntOfInt64(-289))).Contains((_dafny.Zero).Minus(Companion_Default___.Fm44(p1, globalState))), (_index282).Int())
				_1783_v17 = _1783_v17
			}
		}
		var _1784_i6 _dafny.Int
		_ = _1784_i6
		_1784_i6 = _dafny.Zero
		{
			for false {
				{
					if (_1784_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1784_i6 = (_1784_i6).Plus(_dafny.One)
					r0 = p0
					if !(p0) {
						var _1785_v18 _dafny.Array
						_ = _1785_v18
						var _nw276 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
						_ = _nw276
						_1785_v18 = _nw276
						_1785_v18 = (_this).F7()
						var _1786_v19 _dafny.Set
						_ = _1786_v19
						_1786_v19 = _dafny.SetOf(_1731_v3)
						r0 = (_1786_v19).IsProperSubsetOf(_1786_v19)
						_1731_v3 = (_dafny.Zero).Minus((Companion_Default___.Fm44(p1, globalState)).Plus(_1731_v3))
						var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_1785_v18), 0))
						_ = _index283
						(_1785_v18).ArraySet1(_1731_v3, (_index283).Int())
						var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_1785_v18), 0))
						_ = _index284
						(_1785_v18).ArraySet1(Companion_Default___.Fm20(p1, globalState), (_index284).Int())
						var _1787_v20 _dafny.Array
						_ = _1787_v20
						var _len0_42 _dafny.Int = _dafny.IntOfInt64(16)
						_ = _len0_42
						var _nw277 _dafny.Array
						_ = _nw277
						if _len0_42.Cmp(_dafny.Zero) == 0 {
							_nw277 = _dafny.NewArray(_len0_42)
						} else {
							var _init42 func(_dafny.Int) _dafny.Map = (func(_1788_v2 _dafny.CodePoint, _1789_v3 _dafny.Int) func(_dafny.Int) _dafny.Map {
								return func(_1790_i8 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1788_v2, _1789_v3)
								}
							})(_1730_v2, _1731_v3)
							_ = _init42
							var _element0_42 = _init42(_dafny.Zero)
							_ = _element0_42
							_nw277 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
							(_nw277).ArraySet1(_element0_42, 0)
							var _nativeLen0_42 = (_len0_42).Int()
							_ = _nativeLen0_42
							for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
								(_nw277).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
							}
						}
						_1787_v20 = _nw277
						var _1791_v21 *C1
						_ = _1791_v21
						var _nw278 *C1 = New_C1_()
						_ = _nw278
						_nw278.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(966))).Uint32(), func(coer102 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg102 _dafny.Int) interface{} {
								return coer102(arg102)
							}
						}((func(_1792_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1793_i7 _dafny.Int) _dafny.CodePoint {
								return _1792_v2
							}
						})(_1730_v2))), _dafny.UnicodeSeqOfUtf8Bytes("skvawdoq"))).Cardinality()), _1787_v20, (_this).F0(), _this.F1())
						_1791_v21 = _nw278
					} else {
						var _1794_v22 _dafny.Array
						_ = _1794_v22
						var _nw279 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
						_ = _nw279
						_1794_v22 = _nw279
						var _1795_v23 _dafny.Array
						_ = _1795_v23
						var _nw280 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
						_ = _nw280
						_1795_v23 = _nw280
						var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_1794_v22), 0))
						_ = _index285
						(_1794_v22).ArraySet1(_1795_v23, (_index285).Int())
						var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_1794_v22), 0))
						_ = _index286
						(_1794_v22).ArraySet1(_1795_v23, (_index286).Int())
						var _1796_v24 _dafny.Set
						_ = _1796_v24
						_1796_v24 = _dafny.SetOf(_1731_v3)
						var _1797_v25 _dafny.Map
						_ = _1797_v25
						_1797_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1796_v24, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(408))).Uint32(), func(coer103 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg103 _dafny.Int) interface{} {
								return coer103(arg103)
							}
						}((func(_1798_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1799_i9 _dafny.Int) _dafny.Int {
								return _1798_v3
							}
						})(_1731_v3))), (_this).F0()))
						var _1800_v26 _dafny.Map
						_ = _1800_v26
						_1800_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1731_v3, _1796_v24)
						var _1801_v27 _dafny.Map
						_ = _1801_v27
						_1801_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1731_v3, p1)
						_1797_v25 = (_1797_v25).Update((func() _dafny.Set {
							if (_1800_v26).Contains((_1801_v27).Cardinality()) {
								return (_1800_v26).Get((_1801_v27).Cardinality()).(_dafny.Set)
							}
							return func() _dafny.Set {
								var _coll71 = _dafny.NewBuilder()
								_ = _coll71
								for _iter73 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-23), _dafny.IntOfInt64(-358))); ; {
									_compr_71, _ok73 := _iter73()
									if !_ok73 {
										break
									}
									var _1802_v28 _dafny.Int
									_1802_v28 = interface{}(_compr_71).(_dafny.Int)
									if ((_dafny.IntOfInt64(-23)).Cmp(_1802_v28) <= 0) && ((_1802_v28).Cmp(_dafny.IntOfInt64(-358)) < 0) {
										_coll71.Add(Companion_Default___.SafeModuloInt(_1802_v28, _1731_v3))
									}
								}
								return _coll71.ToSet()
							}()
						})(), (_this).F0())
						var _1803_v29 _dafny.Array
						_ = _1803_v29
						var _nw281 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(22))
						_ = _nw281
						_1803_v29 = _nw281
						var _1804_v30 _dafny.Set
						_ = _1804_v30
						_1804_v30 = _dafny.SetOf(p1, p0)
						var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_1803_v29), 0))
						_ = _index287
						(_1803_v29).ArraySet1(_1804_v30, (_index287).Int())
						var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_1803_v29), 0))
						_ = _index288
						(_1803_v29).ArraySet1(_1804_v30, (_index288).Int())
						var _1805_v31 _dafny.Map
						_ = _1805_v31
						_1805_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1731_v3)
						var _1806_v32 _dafny.Set
						_ = _1806_v32
						_1806_v32 = _dafny.SetOf(_1805_v31, _1805_v31)
						_1806_v32 = (func() _dafny.Set {
							if p3 {
								return _1806_v32
							}
							return _1806_v32
						})()
						var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen(((_this).F7()), 0))
						_ = _index289
						((_this).F7()).ArraySet1(_dafny.IntOfInt64(136), (_index289).Int())
						var _1807_v33 D11
						_ = _1807_v33
						_1807_v33 = Companion_D11_.Create_DC25_(_1730_v2, _1731_v3, !(p1))
						var _1808_v34 _dafny.Map
						_ = _1808_v34
						_1808_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1807_v33, _dafny.IntOfInt64(-795))
						var _1809_v36 _dafny.Map
						_ = _1809_v36
						_1809_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1807_v33, p1)
						var _1810_v37 _dafny.MultiSet
						_ = _1810_v37
						_1810_v37 = _dafny.MultiSetOf((_1808_v34).Merge(_1808_v34), (_1808_v34).Merge(func() _dafny.Map {
							var _coll72 = _dafny.NewMapBuilder()
							_ = _coll72
							for _iter74 := _dafny.Iterate((_1809_v36).Keys().Elements()); ; {
								_compr_72, _ok74 := _iter74()
								if !_ok74 {
									break
								}
								var _1811_v35 D11
								_1811_v35 = interface{}(_compr_72).(D11)
								if (_1809_v36).Contains(_1811_v35) {
									_coll72.Add(_1811_v35, _1731_v3)
								}
							}
							return _coll72.ToMap()
						}()))
						var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen(((_this).F7()), 0))
						_ = _index290
						((_this).F7()).ArraySet1((func() _dafny.Int {
							if (_1810_v37).Contains((Companion_Default___.Fm60(_1731_v3, p1, globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1807_v33, _1731_v3))) {
								return (_1810_v37).Multiplicity((Companion_Default___.Fm60(_1731_v3, p1, globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1807_v33, _1731_v3)))
							}
							return _dafny.IntOfUint32((_1728_v0).Cardinality())
						})(), (_index290).Int())
					}
					_1731_v3 = (_1731_v3).Minus(Companion_Default___.SafeDivisionInt(_1731_v3, _1731_v3))
					var _1812_v38 _dafny.Array
					_ = _1812_v38
					var _nw282 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
					_ = _nw282
					_1812_v38 = _nw282
					var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_1812_v38), 0))
					_ = _index291
					(_1812_v38).ArraySet1(p0, (_index291).Int())
					var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_1812_v38), 0))
					_ = _index292
					(_1812_v38).ArraySet1((func() bool {
						if true {
							return p1
						}
						return p3
					})(), (_index292).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _1813_v39 _dafny.Sequence
		_ = _1813_v39
		_1813_v39 = _dafny.SeqOf(Companion_Default___.Fm30(_1731_v3, p1, p1, globalState))
		var _1814_v40 _dafny.Map
		_ = _1814_v40
		_1814_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1813_v39).Cardinality()), _1730_v2)
		_1814_v40 = (_1814_v40).Update((_1731_v3).Plus((_dafny.Zero).Minus(_1731_v3)), (func() _dafny.CodePoint {
			if true {
				return _1730_v2
			}
			return _1730_v2
		})())
		var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(934), _dafny.ArrayLen(((_this).F7()), 0))
		_ = _index293
		((_this).F7()).ArraySet1((_1731_v3).Times(_dafny.IntOfInt64(446)), (_index293).Int())
		var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(934), _dafny.ArrayLen(((_this).F7()), 0))
		_ = _index294
		((_this).F7()).ArraySet1(_1731_v3, (_index294).Int())
		var _1815_v41 D2
		_ = _1815_v41
		_1815_v41 = Companion_D2_.Create_DC5_(_dafny.IntOfInt64(600), (_this).F0(), false)
		r0 = (_1815_v41).Dtor_cf10()
		r1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("fhwdeveiq"), p1)
		r2 = (_this).F0()
		return r0, r1, r2
	}
}
func (_this *C5) F7() _dafny.Array {
	{
		return _this._f7
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f6 _dafny.Int
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f6 = _dafny.Zero
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) Ctor__(f6 _dafny.Int) {
	{
		(_this)._f6 = f6
	}
}
func (_this *C6) Fm11(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(261)
	}
}
func (_this *C6) Fm12(p0 bool, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(409)
	}
}
func (_this *C6) M6(p0 D0, p1 _dafny.Sequence, p2 D0, p3 _dafny.Int, globalState *GlobalState) (bool, bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _pat_let_tv46 = p1
		_ = _pat_let_tv46
		var _source20 D0 = func(_pat_let38_0 D0) D0 {
			return func(_1816_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let39_0 _dafny.Int) D0 {
					return func(_1817_dt__update_hcf2_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC1_(_1817_dt__update_hcf2_h0)
					}(_pat_let39_0)
				}((_dafny.SetOf((_this).F6(), _dafny.IntOfUint32((_pat_let_tv46).Cardinality()))).Cardinality())
			}(_pat_let38_0)
		}(Companion_D0_.Create_DC1_(p3))
		_ = _source20
		if _source20.Is_DC0() {
			var _1818___mcc_h0 _dafny.Array = _source20.Get_().(D0_DC0).Cf0
			_ = _1818___mcc_h0
			var _1819___mcc_h1 _dafny.Int = _source20.Get_().(D0_DC0).Cf1
			_ = _1819___mcc_h1
			var _1820_cf1 _dafny.Int = _1819___mcc_h1
			_ = _1820_cf1
			var _1821_cf0 _dafny.Array = _1818___mcc_h0
			_ = _1821_cf0
			var _1822_v0 bool
			_ = _1822_v0
			_1822_v0 = false
			if _1822_v0 {
				var _1823_v1 _dafny.Sequence
				_ = _1823_v1
				_1823_v1 = _dafny.SeqOf((func() bool {
					if _1822_v0 {
						return _1822_v0
					}
					return _1822_v0
				})(), _1822_v0, _dafny.Companion_Sequence_.IsPrefixOf(p1, p1), _1822_v0)
				_1823_v1 = _1823_v1
				var _1824_v2 _dafny.Sequence
				_ = _1824_v2
				_1824_v2 = _dafny.SeqOf(_1820_cf1)
				r3 = (_1824_v2).Select((Companion_Default___.SafeIndex((_this).Fm12(_1822_v0, _1822_v0, globalState), _dafny.IntOfUint32((_1824_v2).Cardinality()))).Uint32()).(_dafny.Int)
				_1820_cf1 = p3
				var _1825_v3 _dafny.MultiSet
				_ = _1825_v3
				_1825_v3 = _dafny.MultiSetOf(((_this).F6()).Times(_dafny.IntOfUint32((_1823_v1).Cardinality())))
				var _1826_v4 _dafny.CodePoint
				_ = _1826_v4
				_1826_v4 = _dafny.CodePoint('r')
				var _1827_v5 _dafny.Map
				_ = _1827_v5
				_1827_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F6()), _1822_v0)
				var _1828_v6 _dafny.MultiSet
				_ = _1828_v6
				_1828_v6 = _dafny.MultiSetOf(false)
				var _1829_v7 _dafny.Map
				_ = _1829_v7
				_1829_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1828_v6).Cardinality(), (_this).F6())
				var _1830_v8 _dafny.Sequence
				_ = _1830_v8
				_1830_v8 = _dafny.SeqOf(_1829_v7, _1829_v7)
				_1825_v3 = (Companion_Default___.Fm13(_1826_v4, !((func() bool {
					if (_1827_v5).Contains(p3) {
						return (_1827_v5).Get(p3).(bool)
					}
					return _1822_v0
				})()), globalState)).Update(_dafny.IntOfInt64(111), Companion_Default___.Abs((p3).Minus(_dafny.IntOfUint32((_1830_v8).Cardinality()))))
				var _1831_v9 *C2
				_ = _1831_v9
				var _nw283 *C2 = New_C2_()
				_ = _nw283
				_nw283.Ctor__()
				_1831_v9 = _nw283
			} else {
				var _1832_v10 _dafny.Array
				_ = _1832_v10
				var _nw284 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
				_ = _nw284
				_1832_v10 = _nw284
				var _1833_v11 _dafny.Map
				_ = _1833_v11
				_1833_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1822_v0), _1822_v0)
				var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_1832_v10), 0))
				_ = _index295
				(_1832_v10).ArraySet1((_1833_v11).Merge(_1833_v11), (_index295).Int())
				var _1834_v12 _dafny.Sequence
				_ = _1834_v12
				_1834_v12 = _dafny.SeqOf(_1822_v0, _1822_v0)
				var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_1832_v10), 0))
				_ = _index296
				(_1832_v10).ArraySet1((_1833_v11).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1834_v12).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1834_v12).Cardinality()))).Uint32()).(bool), _1822_v0)), (_index296).Int())
				var _1835_v13 _dafny.Map
				_ = _1835_v13
				_1835_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1822_v0, (_this).F6())
				_1835_v13 = (_1835_v13).Update((_1834_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.IntOfUint32((_1834_v12).Cardinality()))).Uint32()).(bool), _1820_cf1)
				var _1836_v14 _dafny.Array
				_ = _1836_v14
				var _nw285 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(5))
				_ = _nw285
				_1836_v14 = _nw285
				var _1837_v15 _dafny.CodePoint
				_ = _1837_v15
				_1837_v15 = _dafny.CodePoint('s')
				var _1838_v16 _dafny.MultiSet
				_ = _1838_v16
				_1838_v16 = _dafny.MultiSetOf(_dafny.CodePoint('t'), _1837_v15)
				var _1839_v17 _dafny.Sequence
				_ = _1839_v17
				_1839_v17 = _dafny.SeqOf(p3, (_dafny.Zero).Minus((_1838_v16).Cardinality()), (_this).F6())
				var _1840_v18 *C1
				_ = _1840_v18
				var _nw286 *C1 = New_C1_()
				_ = _nw286
				_nw286.Ctor__((_1835_v13).Cardinality(), _1836_v14, _1839_v17, p0)
				_1840_v18 = _nw286
				var _1841_v19 _dafny.Map
				_ = _1841_v19
				_1841_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1840_v18, _1840_v18.F9)
				_1820_cf1 = ((_this).Fm12(_1822_v0, !(_1822_v0), globalState)).Plus((func() _dafny.Int {
					if (_1841_v19).Contains(_1840_v18) {
						return (_1841_v19).Get(_1840_v18).(_dafny.Int)
					}
					return (_this).F6()
				})())
				(_1840_v18).F9 = (_1840_v18).Fm6(p3, _1820_cf1, globalState)
				var _1842_v20 _dafny.Array
				_ = _1842_v20
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_43
				var _nw287 _dafny.Array
				_ = _nw287
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw287 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) bool = (func(_1843_v18 *C1, _1844_p1 _dafny.Sequence, _1845_v0 bool) func(_dafny.Int) bool {
						return func(_1846_i0 _dafny.Int) bool {
							return ((_dafny.MultiSetOf(_1843_v18.F9, (_this).F6())).Cardinality()).Cmp((Companion_D4_.Create_DC11_(_dafny.IntOfUint32((_1844_p1).Cardinality()), false, _1845_v0, _1845_v0, !(_1845_v0))).Dtor_cf22()) < 0
						}
					})(_1840_v18, p1, _1822_v0)
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw287 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw287).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw287).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_1842_v20 = _nw287
				var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(276), _dafny.ArrayLen((_1842_v20), 0))
				_ = _index297
				(_1842_v20).ArraySet1(_1822_v0, (_index297).Int())
				var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(276), _dafny.ArrayLen((_1842_v20), 0))
				_ = _index298
				(_1842_v20).ArraySet1(_1822_v0, (_index298).Int())
			}
			var _1847_v21 _dafny.Map
			_ = _1847_v21
			_1847_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F6())
			_1847_v21 = (_1847_v21).Update(p0, _dafny.IntOfInt64(509))
			var _1848_v22 _dafny.CodePoint
			_ = _1848_v22
			_1848_v22 = _dafny.CodePoint('i')
			var _1849_v23 D16
			_ = _1849_v23
			_1849_v23 = Companion_D16_.Create_DC43_(p1, _dafny.CodePoint('u'), _dafny.IntOfInt64(-656))
			var _1850_v24 _dafny.Sequence
			_ = _1850_v24
			_1850_v24 = _dafny.SeqOf(p1, p1, (_1849_v23).Dtor_cf79())
			var _1851_v25 _dafny.Sequence
			_ = _1851_v25
			_1851_v25 = _dafny.SeqOf((_1850_v24).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1850_v24).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _1852_v26 _dafny.Map
			_ = _1852_v26
			_1852_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1822_v0)
			var _1853_v27 _dafny.Map
			_ = _1853_v27
			_1853_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1, p1, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("nydykbd"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-926), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nydykbd")).Cardinality()))).Uint32(), _1848_v22), p1, p1), _1851_v25), _1852_v26)
			_1853_v27 = (_1853_v27).Update(_1850_v24, _1852_v26)
			var _1854_v28 *C3
			_ = _1854_v28
			var _nw288 *C3 = New_C3_()
			_ = _nw288
			_nw288.Ctor__()
			_1854_v28 = _nw288
		} else {
			var _1855___mcc_h2 _dafny.Int = _source20.Get_().(D0_DC1).Cf2
			_ = _1855___mcc_h2
			var _1856_cf2 _dafny.Int = _1855___mcc_h2
			_ = _1856_cf2
			var _1857_v29 *C0
			_ = _1857_v29
			var _nw289 *C0 = New_C0_()
			_ = _nw289
			_nw289.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("jvlu"))
			_1857_v29 = _nw289
			var _1858_v30 _dafny.Map
			_ = _1858_v30
			_1858_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(461), true)
			var _1859_v31 _dafny.Map
			_ = _1859_v31
			_1859_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), !((func() bool {
				if (_1858_v30).Contains(_dafny.IntOfInt64(207)) {
					return (_1858_v30).Get(_dafny.IntOfInt64(207)).(bool)
				}
				return true
			})()))
			var _1860_v32 bool
			_ = _1860_v32
			_1860_v32 = true
			r1 = ((_1858_v30).Equals(_1859_v31)) == (_1860_v32)
			r1 = !((_1860_v32) && (_1860_v32)) || (_1860_v32)
			r2 = _1860_v32
		}
		var _1861_v33 bool
		_ = _1861_v33
		_1861_v33 = true
		var _1862_v34 _dafny.MultiSet
		_ = _1862_v34
		_1862_v34 = _dafny.MultiSetOf(_1861_v33)
		var _1863_v35 _dafny.Sequence
		_ = _1863_v35
		_1863_v35 = _dafny.SeqOf(_1862_v34, _1862_v34)
		var _1864_v36 _dafny.Map
		_ = _1864_v36
		_1864_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1861_v33)
		var _1865_v37 _dafny.MultiSet
		_ = _1865_v37
		_1865_v37 = _dafny.MultiSetOf((_this).F6())
		var _1866_v38 _dafny.Set
		_ = _1866_v38
		_1866_v38 = _dafny.SetOf(_1861_v33)
		var _1867_v39 _dafny.Array
		_ = _1867_v39
		var _nwElement0_68 _dafny.Int = _dafny.IntOfUint32((_1863_v35).Cardinality())
		_ = _nwElement0_68
		var _nw290 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(29))
		_ = _nw290
		(_nw290).ArraySet1(_nwElement0_68, 0)
		(_nw290).ArraySet1(p3, 1)
		(_nw290).ArraySet1((_this).F6(), 2)
		(_nw290).ArraySet1(_dafny.IntOfInt64(-178), 3)
		(_nw290).ArraySet1((p3).Minus(p3), 4)
		(_nw290).ArraySet1((_dafny.Zero).Minus((_1864_v36).Cardinality()), 5)
		(_nw290).ArraySet1(((_dafny.Zero).Minus((_this).F6())).Minus((_dafny.Zero).Minus((_this).F6())), 6)
		(_nw290).ArraySet1(Companion_Default___.SafeModuloInt((_this).F6(), _dafny.IntOfInt64(762)), 7)
		(_nw290).ArraySet1((p3).Minus((_this).F6()), 8)
		(_nw290).ArraySet1((_this).F6(), 9)
		(_nw290).ArraySet1(_dafny.IntOfInt64(747), 10)
		(_nw290).ArraySet1(_dafny.IntOfInt64(-715), 11)
		(_nw290).ArraySet1((_this).F6(), 12)
		(_nw290).ArraySet1(((_this).F6()).Plus((((_1862_v34).Update(_1861_v33, Companion_Default___.Abs((_this).F6()))).Update(_1861_v33, Companion_Default___.Abs((_1865_v37).Cardinality()))).Cardinality()), 13)
		(_nw290).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F6())), 14)
		(_nw290).ArraySet1(p3, 15)
		(_nw290).ArraySet1(p3, 16)
		(_nw290).ArraySet1((_this).F6(), 17)
		(_nw290).ArraySet1((_1865_v37).Cardinality(), 18)
		(_nw290).ArraySet1((_this).F6(), 19)
		(_nw290).ArraySet1((p3).Plus((_this).F6()), 20)
		(_nw290).ArraySet1((_1866_v38).Cardinality(), 21)
		(_nw290).ArraySet1(_dafny.IntOfInt64(97), 22)
		(_nw290).ArraySet1((_dafny.Zero).Minus((p3).Times(p3)), 23)
		(_nw290).ArraySet1((_this).F6(), 24)
		(_nw290).ArraySet1((_this).F6(), 25)
		(_nw290).ArraySet1(_dafny.IntOfInt64(296), 26)
		(_nw290).ArraySet1((_this).F6(), 27)
		(_nw290).ArraySet1((_this).F6(), 28)
		_1867_v39 = _nw290
		for _iter75 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1867_v39), 0))); ; {
			_guard_loop_2, _ok75 := _iter75()
			if !_ok75 {
				break
			}
			var _1868_i1 _dafny.Int
			_1868_i1 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1868_i1).Sign() != -1) && ((_1868_i1).Cmp(_dafny.ArrayLen((_1867_v39), 0)) < 0)) {
				(_1867_v39).ArraySet1(Companion_Default___.SafeDivisionInt(_1868_i1, (_this).F6()), (_1868_i1).Int())
			}
		}
		var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))
		_ = _index299
		(_1867_v39).ArraySet1((_this).F6(), (_index299).Int())
		var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))
		_ = _index300
		(_1867_v39).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((p3).Plus(p3), ((_this).F6()).Plus((_dafny.Zero).Minus((_this).F6())))), (_index300).Int())
		var _1869_v40 _dafny.CodePoint
		_ = _1869_v40
		_1869_v40 = _dafny.CodePoint('l')
		var _1870_v41 D15
		_ = _1870_v41
		_1870_v41 = Companion_D15_.Create_DC38_(_1869_v40)
		var _1871_v42 _dafny.Set
		_ = _1871_v42
		_1871_v42 = _dafny.SetOf((_1870_v41).Dtor_cf71())
		var _1872_v43 _dafny.Sequence
		_ = _1872_v43
		_1872_v43 = _dafny.SeqOf((_this).F6(), (_1871_v42).Cardinality())
		var _1873_i2 _dafny.Int
		_ = _1873_i2
		_1873_i2 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.IsPrefixOf(_1872_v43, _1872_v43) {
				{
					if (_1873_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1873_i2 = (_1873_i2).Plus(_dafny.One)
					_1861_v33 = (_1861_v33) && ((func() bool {
						if _1861_v33 {
							return _1861_v33
						}
						return _1861_v33
					})())
					if _1861_v33 {
						var _1874_v44 D4
						_ = _1874_v44
						_1874_v44 = Companion_D4_.Create_DC10_(p3, (_this).F6(), p3)
						var _1875_v45 D15
						_ = _1875_v45
						_1875_v45 = Companion_D15_.Create_DC39_(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm25(_1874_v44, p1, globalState), p1), Companion_Default___.Fm61(_dafny.IntOfInt64(-243), (_this).F6(), globalState))
						var _1876_v46 _dafny.Map
						_ = _1876_v46
						_1876_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1869_v40, (_this).F6())
						var _1877_v47 *C0
						_ = _1877_v47
						var _nw291 *C0 = New_C0_()
						_ = _nw291
						_nw291.Ctor__(Companion_Default___.Fm16(Companion_Default___.Fm46(_1861_v33, p3, _1861_v33, globalState), (_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int), globalState))
						_1877_v47 = _nw291
						var _1878_v48 _dafny.Set
						_ = _1878_v48
						_1878_v48 = _dafny.SetOf(_1877_v47)
						var _1879_v49 D4
						_ = _1879_v49
						_1879_v49 = Companion_D4_.Create_DC9_(_1865_v37, _1869_v40, p3, _1878_v48, (_1872_v43).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1872_v43).Cardinality()))).Uint32()).(_dafny.Int))
						var _1880_v50 _dafny.Sequence
						_ = _1880_v50
						_1880_v50 = _dafny.SeqOf((_1876_v46).Update((_1879_v49).Dtor_cf15(), p3))
						var _1881_v51 D7
						_ = _1881_v51
						_1881_v51 = Companion_D7_.Create_DC19_(_dafny.IntOfUint32((_1877_v47.F8).Cardinality()))
						var _1882_v52 _dafny.Map
						_ = _1882_v52
						_1882_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1881_v51, p3)
						var _1883_v53 _dafny.Map
						_ = _1883_v53
						_1883_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1882_v52, _1861_v33)
						_1875_v45 = Companion_Default___.Fm62(_1880_v50, (func() bool {
							if (_1883_v53).Contains(_1882_v52) {
								return (_1883_v53).Get(_1882_v52).(bool)
							}
							return _1861_v33
						})(), (_this).F6(), (_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int), globalState)
						var _1884_v54 D13
						_ = _1884_v54
						_1884_v54 = Companion_D13_.Create_DC32_(p1, true, (_this).F6(), _1867_v39)
						_1867_v39 = (_1884_v54).Dtor_cf60()
						var _1885_v55 _dafny.Sequence
						_ = _1885_v55
						_1885_v55 = _dafny.SeqOf(true)
						_1861_v33 = (_1885_v55).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1885_v55).Cardinality()), _dafny.IntOfUint32((_1885_v55).Cardinality()))).Uint32()).(bool)
						r2 = _1861_v33
						r3 = (((Companion_Default___.Fm50(_1861_v33, _1861_v33, _dafny.IntOfUint32((p1).Cardinality()), globalState)).Cardinality()).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1877_v47.F8, (Companion_Default___.SafeIndex((_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1877_v47.F8).Cardinality()))).Uint32(), _1869_v40)).Cardinality()))).Plus((_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int))
					} else {
						var _1886_v56 _dafny.Array
						_ = _1886_v56
						var _nw292 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
						_ = _nw292
						_1886_v56 = _nw292
						var _1887_v57 _dafny.Sequence
						_ = _1887_v57
						_1887_v57 = _dafny.SeqOf(_1861_v33, _1861_v33)
						var _1888_v58 _dafny.Array
						_ = _1888_v58
						var _nwElement0_69 bool = _1861_v33
						_ = _nwElement0_69
						var _nw293 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(18))
						_ = _nw293
						(_nw293).ArraySet1(_nwElement0_69, 0)
						(_nw293).ArraySet1(_1861_v33, 1)
						(_nw293).ArraySet1(_1861_v33, 2)
						(_nw293).ArraySet1(true, 3)
						(_nw293).ArraySet1(Companion_Default___.Fm17(globalState), 4)
						(_nw293).ArraySet1(!(!(_1861_v33)), 5)
						(_nw293).ArraySet1(Companion_Default___.Fm17(globalState), 6)
						(_nw293).ArraySet1(true, 7)
						(_nw293).ArraySet1(!(false), 8)
						(_nw293).ArraySet1(true, 9)
						(_nw293).ArraySet1(_1861_v33, 10)
						(_nw293).ArraySet1(_1861_v33, 11)
						(_nw293).ArraySet1(_1861_v33, 12)
						(_nw293).ArraySet1(false, 13)
						(_nw293).ArraySet1(!((_1887_v57).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1887_v57).Cardinality()))).Uint32()).(bool)), 14)
						(_nw293).ArraySet1(_1861_v33, 15)
						(_nw293).ArraySet1(_1861_v33, 16)
						(_nw293).ArraySet1(!(_1861_v33), 17)
						_1888_v58 = _nw293
						var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_1886_v56), 0))
						_ = _index301
						(_1886_v56).ArraySet1(_1888_v58, (_index301).Int())
						var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_1886_v56), 0))
						_ = _index302
						(_1886_v56).ArraySet1(_1888_v58, (_index302).Int())
						_1864_v36 = (_1864_v36).Update((_dafny.Zero).Minus(p3), (_1861_v33) && (Companion_Default___.Fm30((_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int), (_1887_v57).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F6()), _dafny.IntOfUint32((_1887_v57).Cardinality()))).Uint32()).(bool), !(_1861_v33), globalState)))
						r0 = _1861_v33
						var _1889_v59 _dafny.Map
						_ = _1889_v59
						_1889_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1869_v40, (_1864_v36).Cardinality())
						_1889_v59 = (_1889_v59).Update((func() _dafny.CodePoint {
							if _1861_v33 {
								return _1869_v40
							}
							return Companion_Default___.Fm29(!(_1861_v33), globalState)
						})(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yyicitvk")).Cardinality()))
						var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))
						_ = _index303
						(_1867_v39).ArraySet1((_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int), (_index303).Int())
					}
					var _1890_v60 _dafny.Sequence
					_ = _1890_v60
					_1890_v60 = _dafny.SeqOf(_1861_v33, _1861_v33, Companion_Default___.Fm17(globalState), _1861_v33, _1861_v33)
					var _1891_v61 D12
					_ = _1891_v61
					_1891_v61 = Companion_D12_.Create_DC28_((_1890_v60).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1890_v60).Cardinality()))).Uint32()).(bool), p3)
					var _1892_v62 _dafny.Map
					_ = _1892_v62
					_1892_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p3)
					var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))
					_ = _index304
					(_1867_v39).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p3, p3), _dafny.SeqOf(_dafny.IntOfUint32((p1).Cardinality()), _dafny.IntOfUint32((_1890_v60).Cardinality()), (_1891_v61).Dtor_cf47(), (_1892_v62).Cardinality()))).Cardinality()), (_index304).Int())
					_1864_v36 = (_1864_v36).Update(p3, _1861_v33)
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		if ((_dafny.IntOfUint32((p1).Cardinality())).Minus((_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int))).Cmp((_this).F6()) < 0 {
			var _1893_v63 D12
			_ = _1893_v63
			_1893_v63 = Companion_D12_.Create_DC29_(!(true), (_this).F6(), _1861_v33, Companion_Default___.SafeDivisionInt((_this).F6(), (_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int)), !(_1861_v33))
			_1893_v63 = Companion_D12_.Create_DC29_(_1861_v33, p3, _1861_v33, ((_this).F6()).Times((_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int)), !(_1861_v33))
			var _1894_v64 _dafny.Sequence
			_ = _1894_v64
			_1894_v64 = _dafny.SeqOf(!(_1866_v38).Equals(_1866_v38))
			_1894_v64 = _1894_v64
			var _1895_v65 D4
			_ = _1895_v65
			_1895_v65 = Companion_D4_.Create_DC8_(_dafny.SeqOf(p2, p2))
			var _1896_v66 _dafny.Map
			_ = _1896_v66
			_1896_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1861_v33, _1895_v65)
			var _1897_v67 _dafny.Sequence
			_ = _1897_v67
			_1897_v67 = _dafny.SeqOf(p2, p2)
			_1896_v66 = (_1896_v66).Update(!(_1861_v33), Companion_D4_.Create_DC8_(_1897_v67))
			var _1898_v68 *C3
			_ = _1898_v68
			var _nw294 *C3 = New_C3_()
			_ = _nw294
			_nw294.Ctor__()
			_1898_v68 = _nw294
			var _1899_v69 _dafny.Sequence
			_ = _1899_v69
			_1899_v69 = _dafny.UnicodeSeqOfUtf8Bytes("knsogq")
			_1899_v69 = p1
		} else {
			var _1900_v70 _dafny.Array
			_ = _1900_v70
			var _nwElement0_70 bool = _1861_v33
			_ = _nwElement0_70
			var _nw295 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(12))
			_ = _nw295
			(_nw295).ArraySet1(_nwElement0_70, 0)
			(_nw295).ArraySet1(Companion_Default___.Fm17(globalState), 1)
			(_nw295).ArraySet1(Companion_Default___.Fm30((_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int), true, !(_1861_v33), globalState), 2)
			(_nw295).ArraySet1(_1861_v33, 3)
			(_nw295).ArraySet1(true, 4)
			(_nw295).ArraySet1(false, 5)
			(_nw295).ArraySet1(((_this).F6()).Cmp((_dafny.Zero).Minus((func() _dafny.Int {
				if (_1865_v37).Contains(_dafny.IntOfInt64(272)) {
					return (_1865_v37).Multiplicity(_dafny.IntOfInt64(272))
				}
				return p3
			})())) > 0, 6)
			(_nw295).ArraySet1(_1861_v33, 7)
			(_nw295).ArraySet1(_1861_v33, 8)
			(_nw295).ArraySet1(_1861_v33, 9)
			(_nw295).ArraySet1(_1861_v33, 10)
			(_nw295).ArraySet1(_1861_v33, 11)
			_1900_v70 = _nw295
			var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_1900_v70), 0))
			_ = _index305
			(_1900_v70).ArraySet1(_1861_v33, (_index305).Int())
			var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_1900_v70), 0))
			_ = _index306
			(_1900_v70).ArraySet1(_1861_v33, (_index306).Int())
			var _1901_v71 _dafny.Sequence
			_ = _1901_v71
			_1901_v71 = _dafny.UnicodeSeqOfUtf8Bytes("ydabfju")
			var _1902_v72 _dafny.MultiSet
			_ = _1902_v72
			_1902_v72 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1872_v43).Cardinality()), _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _1869_v40), _dafny.UnicodeSeqOfUtf8Bytes("ayhpyhe"))
			var _1903_v73 _dafny.Set
			_ = _1903_v73
			_1903_v73 = _dafny.SetOf((_this).F6(), (_this).F6())
			var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_1900_v70), 0))
			_ = _index307
			var _rhs208 bool = ((func() _dafny.MultiSet {
				if _1861_v33 {
					return _1902_v72
				}
				return _1902_v72
			})()).IsSubsetOf(_1902_v72)
			_ = _rhs208
			var _rhs209 _dafny.Int = (_this).F6()
			_ = _rhs209
			var _rhs210 _dafny.Int = (func() _dafny.Int {
				if _1861_v33 {
					return (_this).F6()
				}
				return Companion_Default___.SafeModuloInt((_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int), (_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int))
			})()
			_ = _rhs210
			var _rhs211 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("su"), _1901_v71)
			_ = _rhs211
			var _rhs212 bool = (_1903_v73).IsProperSubsetOf(_1903_v73)
			_ = _rhs212
			var _lhs120 _dafny.Array = _1900_v70
			_ = _lhs120
			var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_1900_v70), 0))
			_ = _lhs121
			r2 = _rhs208
			r3 = _rhs209
			r3 = _rhs210
			_1901_v71 = _rhs211
			(_lhs120).ArraySet1(_rhs212, (_lhs121).Int())
			var _1904_v74 D12
			_ = _1904_v74
			_1904_v74 = Companion_D12_.Create_DC28_(_1861_v33, p3)
			var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))
			_ = _index308
			(_1867_v39).ArraySet1((_1904_v74).Dtor_cf47(), (_index308).Int())
			r3 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F6()), (_this).Fm12((_1900_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_1900_v70), 0))).Int()).(bool), (_1900_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_1900_v70), 0))).Int()).(bool), globalState))
			_1861_v33 = (_1900_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_1900_v70), 0))).Int()).(bool)
		}
		var _1905_v75 _dafny.Sequence
		_ = _1905_v75
		_1905_v75 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fm"), p1), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1861_v33)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fm"), p1)).Cardinality()))).Uint32(), Companion_Default___.Fm51(_1861_v33, (_1867_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1867_v39), 0))).Int()).(_dafny.Int), p3, globalState)), p1, p1, p1, p1)
		_1905_v75 = _1905_v75
		r0 = _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-746))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg104 _dafny.Int) interface{} {
				return coer104(arg104)
			}
		}((func(_1906_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1907_i3 _dafny.Int) _dafny.CodePoint {
				return _1906_v40
			}
		})(_1869_v40))), _1869_v40)
		r1 = !(Companion_Default___.Fm17(globalState))
		r2 = _1861_v33
		var _1908_v76 _dafny.Array
		_ = _1908_v76
		var _len0_44 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_44
		var _nw296 _dafny.Array
		_ = _nw296
		if _len0_44.Cmp(_dafny.Zero) == 0 {
			_nw296 = _dafny.NewArray(_len0_44)
		} else {
			var _init44 func(_dafny.Int) bool = (func(_1909_v33 bool, _1910_v36 _dafny.Map, _1911_v39 _dafny.Array) func(_dafny.Int) bool {
				return func(_1912_i4 _dafny.Int) bool {
					return (func() bool {
						if (_1910_v36).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_1909_v33)).Cardinality())) {
							return (_1910_v36).Get(_dafny.IntOfUint32((_dafny.SeqOf(_1909_v33)).Cardinality())).(bool)
						}
						return (func() bool {
							if (_1910_v36).Contains((_1911_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1911_v39), 0))).Int()).(_dafny.Int)) {
								return (_1910_v36).Get((_1911_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1911_v39), 0))).Int()).(_dafny.Int)).(bool)
							}
							return _1909_v33
						})()
					})()
				}
			})(_1861_v33, _1864_v36, _1867_v39)
			_ = _init44
			var _element0_44 = _init44(_dafny.Zero)
			_ = _element0_44
			_nw296 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
			(_nw296).ArraySet1(_element0_44, 0)
			var _nativeLen0_44 = (_len0_44).Int()
			_ = _nativeLen0_44
			for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
				(_nw296).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
			}
		}
		_1908_v76 = _nw296
		var _1913_v77 _dafny.Set
		_ = _1913_v77
		_1913_v77 = _dafny.SetOf(_1908_v76)
		var _1914_v78 _dafny.Set
		_ = _1914_v78
		_1914_v78 = _1913_v77
		var _1915_v79 _dafny.Map
		_ = _1915_v79
		_1915_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1861_v33, _1914_v78)
		var _1916_v80 bool
		_ = _1916_v80
		_1916_v80 = _1861_v33
		r3 = Companion_Default___.SafeDivisionInt(((_1915_v79).Merge((_1915_v79).Update((_1916_v80), _1913_v77))).Cardinality(), _dafny.IntOfInt64(93))
		return r0, r1, r2, r3
	}
}
func (_this *C6) F6() _dafny.Int {
	{
		return _this._f6
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f1 D0
	_f0 _dafny.Sequence
	_f4 _dafny.MultiSet
	_f5 D0
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f1 = Companion_D0_.Default()
	_this._f0 = _dafny.EmptySeq
	_this._f4 = _dafny.EmptyMultiSet
	_this._f5 = Companion_D0_.Default()
	return &_this
}

type CompanionStruct_C7_ struct {
}

var Companion_C7_ = CompanionStruct_C7_{}

func (_this *C7) Equals(other *C7) bool {
	return _this == other
}

func (_this *C7) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C7)
	return ok && _this.Equals(other)
}

func (*C7) String() string {
	return "_module.C7"
}

func Type_C7_() _dafny.TypeDescriptor {
	return type_C7_{}
}

type type_C7_ struct {
}

func (_this type_C7_) Default() interface{} {
	return (*C7)(nil)
}

func (_this type_C7_) String() string {
	return "main.C7"
}
func (_this *C7) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T1_.TraitID_}
}

var _ T1 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) F1() D0 {
	return _this._f1
}
func (_this *C7) F1_set_(value D0) {
	_this._f1 = value
}
func (_this *C7) F0() _dafny.Sequence {
	return _this._f0
}
func (_this *C7) Ctor__(f4 _dafny.MultiSet, f5 D0, f0 _dafny.Sequence, f1 D0) {
	{
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f0 = f0
		(_this)._f1 = f1
	}
}
func (_this *C7) Fm6(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfInt64(-76)).Minus((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("af")).Cardinality())).Minus(_dafny.IntOfInt64(-575)))
	}
}
func (_this *C7) Fm8(p0 _dafny.Map, globalState *GlobalState) bool {
	{
		if (!(false)) && (false) {
			return !((_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(440))).Uint32(), func(coer105 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}(func(_1917_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			}))).Cardinality())) != 0)
		} else {
			return (!(!(!(!(!(!(false))))))) == (false)
		}
	}
}
func (_this *C7) Fm9(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Map, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(606), _dafny.IntOfInt64(-633))).Merge(func() _dafny.Map {
			var _coll73 = _dafny.NewMapBuilder()
			_ = _coll73
			for _iter76 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(834), _dafny.IntOfInt64(708))); ; {
				_compr_73, _ok76 := _iter76()
				if !_ok76 {
					break
				}
				var _1918_v0 _dafny.Int
				_1918_v0 = interface{}(_compr_73).(_dafny.Int)
				if ((_dafny.IntOfInt64(834)).Cmp(_1918_v0) <= 0) && ((_1918_v0).Cmp(_dafny.IntOfInt64(708)) < 0) {
					_coll73.Add((_1918_v0).Times(_dafny.IntOfInt64(597)), _dafny.IntOfInt64(52))
				}
			}
			return _coll73.ToMap()
		}())).Merge((func() _dafny.Map {
			var _coll74 = _dafny.NewMapBuilder()
			_ = _coll74
			for _iter77 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(329), true)).Keys().Elements()); ; {
				_compr_74, _ok77 := _iter77()
				if !_ok77 {
					break
				}
				var _1919_v1 _dafny.Int
				_1919_v1 = interface{}(_compr_74).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(329), true)).Contains(_1919_v1) {
					_coll74.Add(Companion_Default___.SafeModuloInt(_1919_v1, _dafny.IntOfInt64(809)), _dafny.IntOfInt64(743))
				}
			}
			return _coll74.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(312))).Uint32(), func(coer106 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg106 _dafny.Int) interface{} {
				return coer106(arg106)
			}
		}(func(_1920_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		}))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dx")).Cardinality())))))
	}
}
func (_this *C7) M5(p0 _dafny.Int, p1 _dafny.Array, globalState *GlobalState) {
	{
		var _1921_v0 bool
		_ = _1921_v0
		_1921_v0 = true
		var _1922_i0 _dafny.Int
		_ = _1922_i0
		_1922_i0 = _dafny.Zero
		{
			for _1921_v0 {
				{
					if (_1922_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1922_i0 = (_1922_i0).Plus(_dafny.One)
					_1921_v0 = _1921_v0
					var _1923_v1 _dafny.Array
					_ = _1923_v1
					var _nw297 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(4))
					_ = _nw297
					_1923_v1 = _nw297
					var _len0_45 _dafny.Int = _dafny.IntOfInt64(12)
					_ = _len0_45
					var _nw298 _dafny.Array
					_ = _nw298
					if _len0_45.Cmp(_dafny.Zero) == 0 {
						_nw298 = _dafny.NewArray(_len0_45)
					} else {
						var _init45 func(_dafny.Int) _dafny.CodePoint = func(_1924_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('c')
						}
						_ = _init45
						var _element0_45 = _init45(_dafny.Zero)
						_ = _element0_45
						_nw298 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
						(_nw298).ArraySet1CodePoint(_element0_45, 0)
						var _nativeLen0_45 = (_len0_45).Int()
						_ = _nativeLen0_45
						for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
							(_nw298).ArraySet1CodePoint(_init45(_dafny.IntOf(_i0_45)), _i0_45)
						}
					}
					_1923_v1 = _nw298
					var _1925_v2 _dafny.Array
					_ = _1925_v2
					var _nw299 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
					_ = _nw299
					_1925_v2 = _nw299
					var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1925_v2), 0))
					_ = _index309
					(_1925_v2).ArraySet1((_this).F0(), (_index309).Int())
					var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1925_v2), 0))
					_ = _index310
					(_1925_v2).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F0(), (_this).F0()), (_index310).Int())
					var _1926_v3 _dafny.Array
					_ = _1926_v3
					var _len0_46 _dafny.Int = _dafny.IntOfInt64(25)
					_ = _len0_46
					var _nw300 _dafny.Array
					_ = _nw300
					if _len0_46.Cmp(_dafny.Zero) == 0 {
						_nw300 = _dafny.NewArray(_len0_46)
					} else {
						var _init46 func(_dafny.Int) _dafny.Sequence = func(_1927_i2 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(737))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg107 _dafny.Int) interface{} {
									return coer107(arg107)
								}
							}(func(_1928_i3 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('u')
							}))
						}
						_ = _init46
						var _element0_46 = _init46(_dafny.Zero)
						_ = _element0_46
						_nw300 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
						(_nw300).ArraySet1(_element0_46, 0)
						var _nativeLen0_46 = (_len0_46).Int()
						_ = _nativeLen0_46
						for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
							(_nw300).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
						}
					}
					_1926_v3 = _nw300
					_1926_v3 = _1926_v3
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1929_v4 _dafny.Array
		_ = _1929_v4
		var _nw301 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
		_ = _nw301
		_1929_v4 = _nw301
		var _1930_v5 D0
		_ = _1930_v5
		_1930_v5 = Companion_D0_.Create_DC0_(_1929_v4, _dafny.IntOfInt64(600))
		var _1931_v6 _dafny.MultiSet
		_ = _1931_v6
		_1931_v6 = _dafny.MultiSetOf(_1930_v5)
		var _1932_v7 _dafny.Set
		_ = _1932_v7
		_1932_v7 = _dafny.SetOf((_this).Fm6(((_1931_v6).Update(_1930_v5, Companion_Default___.Abs(p0))).Cardinality(), p0, globalState))
		var _1933_v8 _dafny.Sequence
		_ = _1933_v8
		_1933_v8 = _dafny.SeqOf(_1932_v7, _1932_v7, _1932_v7, _1932_v7, _1932_v7)
		if !((Companion_Default___.Fm10(true, globalState)).IsSubsetOf((_1933_v8).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1933_v8).Cardinality()))).Uint32()).(_dafny.Set))) {
			_1932_v7 = _dafny.SetOf((func() _dafny.Int {
				if _1921_v0 {
					return p0
				}
				return p0
			})(), p0)
			var _1934_v9 _dafny.Array
			_ = _1934_v9
			var _nw302 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw302
			_1934_v9 = _nw302
			_1934_v9 = _1934_v9
			var _1935_v10 _dafny.CodePoint
			_ = _1935_v10
			_1935_v10 = _dafny.CodePoint('c')
			var _1936_v11 _dafny.Sequence
			_ = _1936_v11
			_1936_v11 = _dafny.SeqOf(_1935_v10, _1935_v10, _1935_v10, _1935_v10)
			var _1937_v12 _dafny.Set
			_ = _1937_v12
			_1937_v12 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_1936_v11, _1936_v11))
			var _1938_v13 D2
			_ = _1938_v13
			_1938_v13 = Companion_D2_.Create_DC3_(_1937_v12)
			var _1939_v14 _dafny.Sequence
			_ = _1939_v14
			_1939_v14 = _dafny.SeqOf(_1936_v11)
			var _pat_let_tv47 = _1939_v14
			_ = _pat_let_tv47
			var _pat_let_tv48 = _1939_v14
			_ = _pat_let_tv48
			_1937_v12 = (func(_pat_let40_0 D2) D2 {
				return func(_1940_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let41_0 _dafny.Set) D2 {
						return func(_1942_dt__update_hcf4_h0 _dafny.Set) D2 {
							return Companion_D2_.Create_DC3_(_1942_dt__update_hcf4_h0)
						}(_pat_let41_0)
					}(func() _dafny.Set {
						var _coll75 = _dafny.NewBuilder()
						_ = _coll75
						for _iter78 := _dafny.Iterate((_pat_let_tv47).Elements()); ; {
							_compr_75, _ok78 := _iter78()
							if !_ok78 {
								break
							}
							var _1941_v15 _dafny.Sequence
							_1941_v15 = interface{}(_compr_75).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(_pat_let_tv48, _1941_v15) {
								_coll75.Add(_1941_v15)
							}
						}
						return _coll75.ToSet()
					}())
				}(_pat_let40_0)
			}(_1938_v13)).Dtor_cf4()
			var _1943_v16 _dafny.Array
			_ = _1943_v16
			var _nw303 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
			_ = _nw303
			_1943_v16 = _nw303
			var _1944_v17 D2
			_ = _1944_v17
			_1944_v17 = Companion_D2_.Create_DC4_(p0, _1921_v0, _1943_v16)
			if (_1944_v17).Dtor_cf6() {
				var _1945_v18 _dafny.Map
				_ = _1945_v18
				_1945_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1935_v10, _1936_v11)
				var _1946_v19 _dafny.MultiSet
				_ = _1946_v19
				_1946_v19 = _dafny.MultiSetOf((func() _dafny.Sequence {
					if (_1945_v18).Contains(_1935_v10) {
						return (_1945_v18).Get(_1935_v10).(_dafny.Sequence)
					}
					return _1936_v11
				})(), _1936_v11, _dafny.Companion_Sequence_.Concatenate(_1936_v11, _1936_v11), _1936_v11)
				_1946_v19 = (func() _dafny.MultiSet {
					if _1921_v0 {
						return (_1946_v19).Union(_1946_v19)
					}
					return _1946_v19
				})()
				var _1947_v20 _dafny.Int
				_ = _1947_v20
				_1947_v20 = _dafny.IntOfInt64(221)
				_1947_v20 = _1947_v20
				var _1948_v21 _dafny.Array
				_ = _1948_v21
				var _nw304 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
				_ = _nw304
				_1948_v21 = _nw304
				var _1949_v22 *C1
				_ = _1949_v22
				var _nw305 *C1 = New_C1_()
				_ = _nw305
				_nw305.Ctor__(Companion_Default___.Fm32(_1947_v20, globalState), _1948_v21, (_this).F0(), Companion_Default___.Fm49(_1935_v10, globalState))
				_1949_v22 = _nw305
				var _1950_v23 *C3
				_ = _1950_v23
				var _nw306 *C3 = New_C3_()
				_ = _nw306
				_nw306.Ctor__()
				_1950_v23 = _nw306
				var _1951_v24 _dafny.Map
				_ = _1951_v24
				_1951_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1947_v20, _1947_v20)
				var _1952_v25 _dafny.Map
				_ = _1952_v25
				_1952_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1951_v24).Contains(_1947_v20) {
						return (_1951_v24).Get(_1947_v20).(_dafny.Int)
					}
					return (_dafny.Zero).Minus(_1947_v20)
				})(), (_dafny.Zero).Minus(_1947_v20))
				var _1953_v26 _dafny.Sequence
				_ = _1953_v26
				_1953_v26 = _dafny.SeqOf(_1921_v0, _1921_v0)
				_1952_v25 = (_1952_v25).Update(p0, _dafny.IntOfUint32((_1953_v26).Cardinality()))
			} else {
				var _1954_v27 _dafny.Map
				_ = _1954_v27
				_1954_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1921_v0, p0)
				var _1955_v28 _dafny.Map
				_ = _1955_v28
				_1955_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1954_v27)
				var _1956_v29 _dafny.MultiSet
				_ = _1956_v29
				_1956_v29 = _dafny.MultiSetOf(_1921_v0)
				_1921_v0 = !(Companion_Default___.Fm63((func() _dafny.Map {
					if (_1955_v28).Contains(p0) {
						return (_1955_v28).Get(p0).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1921_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1921_v0), p0)).Cardinality())
				})(), p0, globalState)).Equals(_1956_v29)
				var _1957_v30 _dafny.Int
				_ = _1957_v30
				_1957_v30 = _dafny.IntOfInt64(838)
				_1957_v30 = _1957_v30
				_1957_v30 = _1957_v30
				var _1958_v31 D12
				_ = _1958_v31
				_1958_v31 = Companion_D12_.Create_DC28_(false, _1957_v30)
				var _1959_v32 _dafny.Map
				_ = _1959_v32
				_1959_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1958_v31).Dtor_cf47(), (_this).F0())
				var _1960_v33 _dafny.Map
				_ = _1960_v33
				_1960_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1921_v0, false)
				_1936_v11 = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if _1921_v0 {
						return _dafny.Companion_Sequence_.Update(_1936_v11, (Companion_Default___.SafeIndex((_1959_v32).Cardinality(), _dafny.IntOfUint32((_1936_v11).Cardinality()))).Uint32(), Companion_Default___.Fm51(_1921_v0, (_1960_v33).Cardinality(), _1957_v30, globalState))
					}
					return _dafny.Companion_Sequence_.Concatenate(_1936_v11, _dafny.SeqOf(_1935_v10, _1935_v10, _1935_v10))
				})(), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_1957_v30, p0), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if _1921_v0 {
						return _dafny.Companion_Sequence_.Update(_1936_v11, (Companion_Default___.SafeIndex((_1959_v32).Cardinality(), _dafny.IntOfUint32((_1936_v11).Cardinality()))).Uint32(), Companion_Default___.Fm51(_1921_v0, (_1960_v33).Cardinality(), _1957_v30, globalState))
					}
					return _dafny.Companion_Sequence_.Concatenate(_1936_v11, _dafny.SeqOf(_1935_v10, _1935_v10, _1935_v10))
				})()).Cardinality()))).Uint32(), _1935_v10)
				var _1961_v34 _dafny.Sequence
				_ = _1961_v34
				_1961_v34 = _dafny.SeqOf(false, _1921_v0, (p0).Cmp(_dafny.IntOfInt64(281)) != 0)
				_1961_v34 = _dafny.Companion_Sequence_.Concatenate(_1961_v34, _1961_v34)
			}
			var _1962_v35 _dafny.Array
			_ = _1962_v35
			var _len0_47 _dafny.Int = _dafny.One
			_ = _len0_47
			var _nw307 _dafny.Array
			_ = _nw307
			if _len0_47.Cmp(_dafny.Zero) == 0 {
				_nw307 = _dafny.NewArray(_len0_47)
			} else {
				var _init47 func(_dafny.Int) _dafny.Map = (func(_1963_v10 _dafny.CodePoint, _1964_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_1965_i4 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1963_v10, _1964_p0)
					}
				})(_1935_v10, p0)
				_ = _init47
				var _element0_47 = _init47(_dafny.Zero)
				_ = _element0_47
				_nw307 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
				(_nw307).ArraySet1(_element0_47, 0)
				var _nativeLen0_47 = (_len0_47).Int()
				_ = _nativeLen0_47
				for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
					(_nw307).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
				}
			}
			_1962_v35 = _nw307
			var _1966_v36 *C1
			_ = _1966_v36
			var _nw308 *C1 = New_C1_()
			_ = _nw308
			_nw308.Ctor__(p0, _1962_v35, (_this).F0(), (_this).F5())
			_1966_v36 = _nw308
			var _1967_v37 _dafny.Set
			_ = _1967_v37
			_1967_v37 = _dafny.SetOf(_1966_v36, _1966_v36, _1966_v36)
			var _1968_v38 _dafny.Map
			_ = _1968_v38
			_1968_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), (_1967_v37).Cardinality())
			var _1969_v39 _dafny.Map
			_ = _1969_v39
			_1969_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1921_v0, (_1932_v7).Cardinality())
			var _1970_v40 _dafny.Map
			_ = _1970_v40
			_1970_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1935_v10, _1966_v36.F9)
			var _1971_v41 _dafny.Map
			_ = _1971_v41
			_1971_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm19(_1921_v0, _1921_v0, _1970_v40, false, globalState), p0)
			var _1972_v42 _dafny.Sequence
			_ = _1972_v42
			_1972_v42 = _dafny.SeqOf(_1968_v38, (_1968_v38).Update((_1969_v39).Cardinality(), (_1971_v41).Cardinality()), _1968_v38, (_1968_v38).Merge(_1968_v38))
			var _1973_v43 _dafny.Map
			_ = _1973_v43
			_1973_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-919))).Uint32(), func(coer108 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg108 _dafny.Int) interface{} {
					return coer108(arg108)
				}
			}((func(_1974_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1975_i5 _dafny.Int) _dafny.Int {
					return _1974_p0
				}
			})(p0))), _1972_v42)
			var _1976_v44 _dafny.MultiSet
			_ = _1976_v44
			_1976_v44 = _dafny.MultiSetOf(_1921_v0)
			var _1977_v45 _dafny.Sequence
			_ = _1977_v45
			_1977_v45 = _dafny.SeqOf((_this).F5(), Companion_D0_.Create_DC1_(p0), _this.F1(), Companion_D0_.Create_DC1_(_1966_v36.F9), (_this).F5())
			_1972_v42 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (p0).Cmp(_1966_v36.F9) == 0 {
					return (func() _dafny.Sequence {
						if (_1973_v43).Contains(_dafny.Companion_Sequence_.Update((_this).F0(), (Companion_Default___.SafeIndex((_1976_v44).Cardinality(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32(), _1966_v36.F9)) {
							return (_1973_v43).Get(_dafny.Companion_Sequence_.Update((_this).F0(), (Companion_Default___.SafeIndex((_1976_v44).Cardinality(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32(), _1966_v36.F9)).(_dafny.Sequence)
						}
						return _1972_v42
					})()
				}
				return _dafny.Companion_Sequence_.Concatenate(_1972_v42, _1972_v42)
			})(), (Companion_Default___.SafeIndex((_1966_v36).Fm6(p0, _1966_v36.F9, globalState), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (p0).Cmp(_1966_v36.F9) == 0 {
					return (func() _dafny.Sequence {
						if (_1973_v43).Contains(_dafny.Companion_Sequence_.Update((_this).F0(), (Companion_Default___.SafeIndex((_1976_v44).Cardinality(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32(), _1966_v36.F9)) {
							return (_1973_v43).Get(_dafny.Companion_Sequence_.Update((_this).F0(), (Companion_Default___.SafeIndex((_1976_v44).Cardinality(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32(), _1966_v36.F9)).(_dafny.Sequence)
						}
						return _1972_v42
					})()
				}
				return _dafny.Companion_Sequence_.Concatenate(_1972_v42, _1972_v42)
			})()).Cardinality()))).Uint32(), _1968_v38), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (p0).Cmp(_1966_v36.F9) == 0 {
					return (func() _dafny.Sequence {
						if (_1973_v43).Contains(_dafny.Companion_Sequence_.Update((_this).F0(), (Companion_Default___.SafeIndex((_1976_v44).Cardinality(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32(), _1966_v36.F9)) {
							return (_1973_v43).Get(_dafny.Companion_Sequence_.Update((_this).F0(), (Companion_Default___.SafeIndex((_1976_v44).Cardinality(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32(), _1966_v36.F9)).(_dafny.Sequence)
						}
						return _1972_v42
					})()
				}
				return _dafny.Companion_Sequence_.Concatenate(_1972_v42, _1972_v42)
			})(), (Companion_Default___.SafeIndex((_1966_v36).Fm6(p0, _1966_v36.F9, globalState), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (p0).Cmp(_1966_v36.F9) == 0 {
					return (func() _dafny.Sequence {
						if (_1973_v43).Contains(_dafny.Companion_Sequence_.Update((_this).F0(), (Companion_Default___.SafeIndex((_1976_v44).Cardinality(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32(), _1966_v36.F9)) {
							return (_1973_v43).Get(_dafny.Companion_Sequence_.Update((_this).F0(), (Companion_Default___.SafeIndex((_1976_v44).Cardinality(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32(), _1966_v36.F9)).(_dafny.Sequence)
						}
						return _1972_v42
					})()
				}
				return _dafny.Companion_Sequence_.Concatenate(_1972_v42, _1972_v42)
			})()).Cardinality()))).Uint32(), _1968_v38)).Cardinality()))).Uint32(), (_this).Fm9(_1977_v45, p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SetOf(true)), _dafny.SeqOf(true), globalState))
		} else {
			var _1978_v46 _dafny.Int
			_ = _1978_v46
			_1978_v46 = _dafny.IntOfInt64(842)
			_1978_v46 = ((_dafny.Zero).Minus(p0)).Minus(_1978_v46)
			var _1979_v47 D4
			_ = _1979_v47
			_1979_v47 = Companion_D4_.Create_DC10_((_1978_v46).Plus(p0), _1978_v46, _1978_v46)
			_1979_v47 = _1979_v47
			var _1980_v48 *C2
			_ = _1980_v48
			var _nw309 *C2 = New_C2_()
			_ = _nw309
			_nw309.Ctor__()
			_1980_v48 = _nw309
			var _1981_v49 D16
			_ = _1981_v49
			_1981_v49 = Companion_D16_.Create_DC44_((_dafny.Zero).Minus(p0))
			var _1982_v50 T0
			_ = _1982_v50
			var _nw310 *C2 = New_C2_()
			_ = _nw310
			_nw310.Ctor__()
			_1982_v50 = _nw310
			var _rhs213 D16 = _1981_v49
			_ = _rhs213
			var _rhs214 T0 = (func() T0 {
				if _1921_v0 {
					return _1982_v50
				}
				return _1982_v50
			})()
			_ = _rhs214
			_1981_v49 = _rhs213
			_1982_v50 = _rhs214
			var _1983_v51 _dafny.Sequence
			_ = _1983_v51
			_1983_v51 = _dafny.UnicodeSeqOfUtf8Bytes("ru")
			var _1984_v52 D14
			_ = _1984_v52
			_1984_v52 = Companion_D14_.Create_DC35_(Companion_Default___.SafeDivisionInt(p0, _1978_v46), (func() _dafny.Sequence {
				if true {
					return _1983_v51
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-598))).Uint32(), func(coer109 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg109 _dafny.Int) interface{} {
						return coer109(arg109)
					}
				}(func(_1985_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				}))
			})(), (func() _dafny.Int {
				if ((_this).F4()).Contains(_1978_v46) {
					return ((_this).F4()).Multiplicity(_1978_v46)
				}
				return _1978_v46
			})(), true, (p0).Cmp(p0) >= 0)
			var _source21 D14 = _1984_v52
			_ = _source21
			if _source21.Is_DC35() {
				var _1986___mcc_h0 _dafny.Int = _source21.Get_().(D14_DC35).Cf63
				_ = _1986___mcc_h0
				var _1987___mcc_h1 _dafny.Sequence = _source21.Get_().(D14_DC35).Cf64
				_ = _1987___mcc_h1
				var _1988___mcc_h2 _dafny.Int = _source21.Get_().(D14_DC35).Cf65
				_ = _1988___mcc_h2
				var _1989___mcc_h3 bool = _source21.Get_().(D14_DC35).Cf66
				_ = _1989___mcc_h3
				var _1990___mcc_h4 bool = _source21.Get_().(D14_DC35).Cf67
				_ = _1990___mcc_h4
				var _1991_cf67 bool = _1990___mcc_h4
				_ = _1991_cf67
				var _1992_cf66 bool = _1989___mcc_h3
				_ = _1992_cf66
				var _1993_cf65 _dafny.Int = _1988___mcc_h2
				_ = _1993_cf65
				var _1994_cf64 _dafny.Sequence = _1987___mcc_h1
				_ = _1994_cf64
				var _1995_cf63 _dafny.Int = _1986___mcc_h0
				_ = _1995_cf63
				_1921_v0 = !(!(_1991_cf67) || (!(_1992_cf66) || (true)))
				var _1996_v53 *C0
				_ = _1996_v53
				var _nw311 *C0 = New_C0_()
				_ = _nw311
				_nw311.Ctor__(_1994_cf64)
				_1996_v53 = _nw311
				_1995_cf63 = ((_this).F0()).Select((Companion_Default___.SafeIndex(_1995_cf63, _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int)
				var _1997_v54 _dafny.CodePoint
				_ = _1997_v54
				_1997_v54 = _dafny.CodePoint('m')
				_1997_v54 = _1997_v54
			} else if _source21.Is_DC36() {
				var _1998___mcc_h5 _dafny.Map = _source21.Get_().(D14_DC36).Cf68
				_ = _1998___mcc_h5
				var _1999___mcc_h6 _dafny.Map = _source21.Get_().(D14_DC36).Cf69
				_ = _1999___mcc_h6
				var _2000_cf69 _dafny.Map = _1999___mcc_h6
				_ = _2000_cf69
				var _2001_cf68 _dafny.Map = _1998___mcc_h5
				_ = _2001_cf68
				_1921_v0 = _1921_v0
				var _2002_v55 _dafny.Set
				_ = _2002_v55
				_2002_v55 = _dafny.SetOf(_1921_v0)
				_1921_v0 = (_2002_v55).IsDisjointFrom(_dafny.SetOf(_1921_v0))
				var _2003_v56 *C4
				_ = _2003_v56
				var _nw312 *C4 = New_C4_()
				_ = _nw312
				_nw312.Ctor__()
				_2003_v56 = _nw312
				var _2004_v57 _dafny.Sequence
				_ = _2004_v57
				_2004_v57 = _dafny.SeqOf(true, _1921_v0)
				var _2005_v58 _dafny.Map
				_ = _2005_v58
				_2005_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
					if (_2004_v57).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_2001_cf68).Contains(p0) {
							return (_2001_cf68).Get(p0).(_dafny.Int)
						}
						return p0
					})(), _dafny.IntOfUint32((_2004_v57).Cardinality()))).Uint32()).(bool) {
						return _1932_v7
					}
					return _dafny.SetOf(_1978_v46)
				})(), _1921_v0)
				var _2006_v59 _dafny.Map
				_ = _2006_v59
				_2006_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1921_v0, false)
				_2005_v58 = (_2005_v58).Update(_dafny.SetOf(_dafny.IntOfInt64(602), Companion_Default___.Fm20(!(_1921_v0), globalState), _1978_v46, (_2006_v59).Cardinality(), p0), _1921_v0)
			} else {
				var _2007___mcc_h7 _dafny.Map = _source21.Get_().(D14_DC34).Cf62
				_ = _2007___mcc_h7
				var _2008_cf62 _dafny.Map = _2007___mcc_h7
				_ = _2008_cf62
				var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((p1), 0))
				_ = _index311
				(p1).ArraySet1(_1978_v46, (_index311).Int())
				var _2009_v60 _dafny.CodePoint
				_ = _2009_v60
				_2009_v60 = _dafny.CodePoint('y')
				var _2010_v61 _dafny.Array
				_ = _2010_v61
				var _len0_48 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_48
				var _nw313 _dafny.Array
				_ = _nw313
				if _len0_48.Cmp(_dafny.Zero) == 0 {
					_nw313 = _dafny.NewArray(_len0_48)
				} else {
					var _init48 func(_dafny.Int) bool = (func(_2011_v0 bool) func(_dafny.Int) bool {
						return func(_2012_i7 _dafny.Int) bool {
							return _2011_v0
						}
					})(_1921_v0)
					_ = _init48
					var _element0_48 = _init48(_dafny.Zero)
					_ = _element0_48
					_nw313 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
					(_nw313).ArraySet1(_element0_48, 0)
					var _nativeLen0_48 = (_len0_48).Int()
					_ = _nativeLen0_48
					for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
						(_nw313).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
					}
				}
				_2010_v61 = _nw313
				var _2013_v62 D7
				_ = _2013_v62
				_2013_v62 = Companion_D7_.Create_DC18_()
				var _2014_v63 _dafny.Map
				_ = _2014_v63
				_2014_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1921_v0)
				var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((p1), 0))
				_ = _index312
				var _rhs215 bool = !(!(!((Companion_Default___.Fm48(_2013_v62, (_2014_v63).Cardinality(), p0, p0, globalState)).IsSubsetOf((_this).F4()))))
				_ = _rhs215
				var _rhs216 _dafny.Int = _1978_v46
				_ = _rhs216
				var _rhs217 bool = _1921_v0
				_ = _rhs217
				var _rhs218 _dafny.CodePoint = _2009_v60
				_ = _rhs218
				var _rhs219 _dafny.Array = _2010_v61
				_ = _rhs219
				var _lhs122 _dafny.Array = p1
				_ = _lhs122
				var _lhs123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((p1), 0))
				_ = _lhs123
				_1921_v0 = _rhs215
				(_lhs122).ArraySet1(_rhs216, (_lhs123).Int())
				_1921_v0 = _rhs217
				_2009_v60 = _rhs218
				_2010_v61 = _rhs219
				var _2015_v64 *C6
				_ = _2015_v64
				var _nw314 *C6 = New_C6_()
				_ = _nw314
				_nw314.Ctor__(_1978_v46)
				_2015_v64 = _nw314
				var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_2010_v61), 0))
				_ = _index313
				(_2010_v61).ArraySet1(_1921_v0, (_index313).Int())
				var _2016_v65 _dafny.MultiSet
				_ = _2016_v65
				_2016_v65 = _dafny.MultiSetOf(!(_1921_v0))
				var _2017_v66 _dafny.Sequence
				_ = _2017_v66
				_2017_v66 = _dafny.SeqOf(_1983_v51)
				var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_2010_v61), 0))
				_ = _index314
				var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((p1), 0))
				_ = _index315
				var _rhs220 bool = ((_2016_v65).Union(_2016_v65)).Equals(_2016_v65)
				_ = _rhs220
				var _rhs221 _dafny.Int = (_2015_v64).F6()
				_ = _rhs221
				var _rhs222 _dafny.Array = _2010_v61
				_ = _rhs222
				var _rhs223 _dafny.Sequence = (_2017_v66).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2017_v66).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs223
				var _lhs124 _dafny.Array = _2010_v61
				_ = _lhs124
				var _lhs125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_2010_v61), 0))
				_ = _lhs125
				var _lhs126 _dafny.Array = p1
				_ = _lhs126
				var _lhs127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((p1), 0))
				_ = _lhs127
				(_lhs124).ArraySet1(_rhs220, (_lhs125).Int())
				(_lhs126).ArraySet1(_rhs221, (_lhs127).Int())
				_2010_v61 = _rhs222
				_1983_v51 = _rhs223
				var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((p1), 0))
				_ = _index316
				(p1).ArraySet1(_dafny.IntOfInt64(-723), (_index316).Int())
			}
		}
		var _hi14 _dafny.Int = p0
		_ = _hi14
		for _2018_i8 := (p0).Plus(p0); _2018_i8.Cmp(_hi14) < 0; _2018_i8 = _2018_i8.Plus(_dafny.One) {
			var _2019_v67 _dafny.Map
			_ = _2019_v67
			_2019_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2018_i8, _1921_v0)
			var _2020_v68 _dafny.Map
			_ = _2020_v68
			_2020_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2019_v67)
			var _2021_v69 _dafny.Int
			_ = _2021_v69
			_2021_v69 = _dafny.IntOfInt64(611)
			var _rhs224 _dafny.Map = _2020_v68
			_ = _rhs224
			var _rhs225 _dafny.Int = (p0).Minus((_2018_i8).Plus((_dafny.Zero).Minus(p0)))
			_ = _rhs225
			_2020_v68 = _rhs224
			_2021_v69 = _rhs225
			if _1921_v0 {
				var _2022_v70 _dafny.Sequence
				_ = _2022_v70
				_2022_v70 = _dafny.UnicodeSeqOfUtf8Bytes("ashcuj")
				var _2023_v71 *C0
				_ = _2023_v71
				var _nw315 *C0 = New_C0_()
				_ = _nw315
				_nw315.Ctor__(_2022_v70)
				_2023_v71 = _nw315
				_2021_v69 = _2018_i8
				var _2024_v72 _dafny.Array
				_ = _2024_v72
				var _nw316 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw316
				_2024_v72 = _nw316
				var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2024_v72), 0))
				_ = _index317
				(_2024_v72).ArraySet1(_1921_v0, (_index317).Int())
				var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2024_v72), 0))
				_ = _index318
				(_2024_v72).ArraySet1(_1921_v0, (_index318).Int())
				_2021_v69 = Companion_Default___.SafeModuloInt(p0, _2021_v69)
				var _2025_v73 _dafny.Map
				_ = _2025_v73
				_2025_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0)
				var _2026_v74 _dafny.Sequence
				_ = _2026_v74
				_2026_v74 = _dafny.SeqOf(((_2025_v73).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2024_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2024_v72), 0))).Int()).(bool), _2018_i8))).Cardinality(), _2021_v69)
				_2026_v74 = _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(11))).Uint32(), func(coer110 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg110 _dafny.Int) interface{} {
						return coer110(arg110)
					}
				}((func(_2027_i8 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2028_i9 _dafny.Int) _dafny.Int {
						return _2027_i8
					}
				})(_2018_i8))), (Companion_Default___.SafeIndex(_2021_v69, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(11))).Uint32(), func(coer111 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg111 _dafny.Int) interface{} {
						return coer111(arg111)
					}
				}((func(_2029_i8 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2030_i9 _dafny.Int) _dafny.Int {
						return _2029_i8
					}
				})(_2018_i8)))).Cardinality()))).Uint32(), _2021_v69)
			} else {
				var _2031_v75 _dafny.CodePoint
				_ = _2031_v75
				_2031_v75 = _dafny.CodePoint('m')
				var _2032_v76 _dafny.Sequence
				_ = _2032_v76
				_2032_v76 = _dafny.UnicodeSeqOfUtf8Bytes("unqvelq")
				var _2033_v77 *C0
				_ = _2033_v77
				var _nw317 *C0 = New_C0_()
				_ = _nw317
				_nw317.Ctor__(_2032_v76)
				_2033_v77 = _nw317
				var _2034_v78 _dafny.Set
				_ = _2034_v78
				_2034_v78 = _dafny.SetOf(_2033_v77, _2033_v77, _2033_v77)
				var _2035_v79 _dafny.Array
				_ = _2035_v79
				var _nw318 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
				_ = _nw318
				_2035_v79 = _nw318
				var _2036_v80 D2
				_ = _2036_v80
				_2036_v80 = Companion_D2_.Create_DC4_(_dafny.IntOfUint32((_2033_v77.F8).Cardinality()), _1921_v0, _2035_v79)
				var _2037_v81 D2
				_ = _2037_v81
				_2037_v81 = Companion_D2_.Create_DC6_(_2036_v80)
				var _2038_v82 _dafny.Map
				_ = _2038_v82
				_2038_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2032_v76).Cardinality()), _2037_v81)
				var _2039_v83 D4
				_ = _2039_v83
				_2039_v83 = Companion_D4_.Create_DC9_((_this).F4(), _2031_v75, _2021_v69, _2034_v78, (_2038_v82).Cardinality())
				var _2040_v84 _dafny.Map
				_ = _2040_v84
				_2040_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1921_v0, _2033_v77.F8)
				var _2041_v85 D12
				_ = _2041_v85
				_2041_v85 = Companion_D12_.Create_DC29_(_1921_v0, p0, _1921_v0, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_2040_v84).Contains(_1921_v0) {
						return (_2040_v84).Get(_1921_v0).(_dafny.Sequence)
					}
					return _2032_v76
				})()).Cardinality()), _1921_v0)
				var _pat_let_tv49 = _2033_v77
				_ = _pat_let_tv49
				var _2042_v86 _dafny.MultiSet
				_ = _2042_v86
				_2042_v86 = _dafny.MultiSetOf(_2041_v85, _2041_v85, func(_pat_let42_0 D12) D12 {
					return func(_2043_dt__update__tmp_h1 D12) D12 {
						return func(_pat_let43_0 _dafny.Int) D12 {
							return func(_2044_dt__update_hcf49_h0 _dafny.Int) D12 {
								return Companion_D12_.Create_DC29_((_2043_dt__update__tmp_h1).Dtor_cf48(), _2044_dt__update_hcf49_h0, (_2043_dt__update__tmp_h1).Dtor_cf50(), (_2043_dt__update__tmp_h1).Dtor_cf51(), (_2043_dt__update__tmp_h1).Dtor_cf52())
							}(_pat_let43_0)
						}(_dafny.IntOfUint32((_pat_let_tv49.F8).Cardinality()))
					}(_pat_let42_0)
				}(_2041_v85))
				var _rhs226 bool = _1921_v0
				_ = _rhs226
				var _rhs227 _dafny.Int = p0
				_ = _rhs227
				var _rhs228 D4 = Companion_D4_.Create_DC9_((_this).F4(), _2031_v75, (func() _dafny.Int {
					if (_2042_v86).Contains(_2041_v85) {
						return (_2042_v86).Multiplicity(_2041_v85)
					}
					return p0
				})(), _2034_v78, Companion_Default___.SafeDivisionInt(_2021_v69, (_dafny.Zero).Minus(_2018_i8)))
				_ = _rhs228
				var _rhs229 bool = _1921_v0
				_ = _rhs229
				_1921_v0 = _rhs226
				_2021_v69 = _rhs227
				_2039_v83 = _rhs228
				_1921_v0 = _rhs229
				_2021_v69 = _2021_v69
				var _2045_v90 D17
				_ = _2045_v90
				_2045_v90 = Companion_D17_.Create_DC45_(_dafny.MultiSetOf((_this).F4()))
				var _rhs230 _dafny.Sequence = _2032_v76
				_ = _rhs230
				var _rhs231 _dafny.Int = Companion_Default___.SafeDivisionInt(((func() _dafny.Map {
					var _coll76 = _dafny.NewMapBuilder()
					_ = _coll76
					for _iter79 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-905))).Uint32(), func(coer112 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
						return func(arg112 _dafny.Int) interface{} {
							return coer112(arg112)
						}
					}(func(_2046_i10 _dafny.Int) _dafny.MultiSet {
						return (_this).F4()
					}))).Elements()); ; {
						_compr_76, _ok79 := _iter79()
						if !_ok79 {
							break
						}
						var _2047_v87 _dafny.MultiSet
						_2047_v87 = interface{}(_compr_76).(_dafny.MultiSet)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-905))).Uint32(), func(coer113 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
							return func(arg113 _dafny.Int) interface{} {
								return coer113(arg113)
							}
						}(func(_2046_i10 _dafny.Int) _dafny.MultiSet {
							return (_this).F4()
						})), _2047_v87) {
							_coll76.Add(_2047_v87, _1921_v0)
						}
					}
					return _coll76.ToMap()
				}()).Merge(func() _dafny.Map {
					var _coll77 = _dafny.NewMapBuilder()
					_ = _coll77
					for _iter80 := _dafny.Iterate((func() _dafny.Map {
						var _coll78 = _dafny.NewMapBuilder()
						_ = _coll78
						for _iter81 := _dafny.Iterate(((_2045_v90).Dtor_cf83()).Elements()); ; {
							_compr_78, _ok81 := _iter81()
							if !_ok81 {
								break
							}
							var _2048_v89 _dafny.MultiSet
							_2048_v89 = interface{}(_compr_78).(_dafny.MultiSet)
							if ((_2045_v90).Dtor_cf83()).Contains(_2048_v89) {
								_coll78.Add(_2048_v89, _dafny.SetOf(_1921_v0, _1921_v0))
							}
						}
						return _coll78.ToMap()
					}()).Keys().Elements()); ; {
						_compr_77, _ok80 := _iter80()
						if !_ok80 {
							break
						}
						var _2049_v88 _dafny.MultiSet
						_2049_v88 = interface{}(_compr_77).(_dafny.MultiSet)
						if (func() _dafny.Map {
							var _coll79 = _dafny.NewMapBuilder()
							_ = _coll79
							for _iter82 := _dafny.Iterate(((_2045_v90).Dtor_cf83()).Elements()); ; {
								_compr_79, _ok82 := _iter82()
								if !_ok82 {
									break
								}
								var _2048_v89 _dafny.MultiSet
								_2048_v89 = interface{}(_compr_79).(_dafny.MultiSet)
								if ((_2045_v90).Dtor_cf83()).Contains(_2048_v89) {
									_coll79.Add(_2048_v89, _dafny.SetOf(_1921_v0, _1921_v0))
								}
							}
							return _coll79.ToMap()
						}()).Contains(_2049_v88) {
							_coll77.Add(_2049_v88, _1921_v0)
						}
					}
					return _coll77.ToMap()
				}())).Cardinality(), p0)
				_ = _rhs231
				var _rhs232 _dafny.Int = p0
				_ = _rhs232
				_2032_v76 = _rhs230
				_2021_v69 = _rhs231
				_2021_v69 = _rhs232
				var _2050_v91 _dafny.Sequence
				_ = _2050_v91
				_2050_v91 = _dafny.SeqOf((func() bool {
					if (_2019_v67).Contains(_dafny.IntOfUint32(((_this).F0()).Cardinality())) {
						return (_2019_v67).Get(_dafny.IntOfUint32(((_this).F0()).Cardinality())).(bool)
					}
					return _1921_v0
				})(), _1921_v0, _1921_v0)
				_1921_v0 = ((func() _dafny.Int {
					if _1921_v0 {
						return (_this).Fm6(_2021_v69, (Companion_Default___.Fm24(_1921_v0, globalState)).Cardinality(), globalState)
					}
					return _dafny.IntOfUint32((_2050_v91).Cardinality())
				})()).Cmp(p0) != 0
				var _2051_v92 _dafny.Array
				_ = _2051_v92
				var _len0_49 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_49
				var _nw319 _dafny.Array
				_ = _nw319
				if _len0_49.Cmp(_dafny.Zero) == 0 {
					_nw319 = _dafny.NewArray(_len0_49)
				} else {
					var _init49 func(_dafny.Int) bool = (func(_2052_v77 *C0, _2053_v76 _dafny.Sequence) func(_dafny.Int) bool {
						return func(_2054_i11 _dafny.Int) bool {
							return !_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("tvmhkbqyi"), (_2052_v77.F8).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2053_v76).Cardinality()), _dafny.IntOfUint32((_2052_v77.F8).Cardinality()))).Uint32()).(_dafny.CodePoint))
						}
					})(_2033_v77, _2032_v76)
					_ = _init49
					var _element0_49 = _init49(_dafny.Zero)
					_ = _element0_49
					_nw319 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
					(_nw319).ArraySet1(_element0_49, 0)
					var _nativeLen0_49 = (_len0_49).Int()
					_ = _nativeLen0_49
					for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
						(_nw319).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
					}
				}
				_2051_v92 = _nw319
				var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_2051_v92), 0))
				_ = _index319
				(_2051_v92).ArraySet1(_1921_v0, (_index319).Int())
				var _2055_v93 _dafny.Array
				_ = _2055_v93
				var _nw320 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
				_ = _nw320
				_2055_v93 = _nw320
				var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_2051_v92), 0))
				_ = _index320
				var _rhs233 bool = false
				_ = _rhs233
				var _rhs234 bool = _1921_v0
				_ = _rhs234
				var _rhs235 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("yyco")
				_ = _rhs235
				var _rhs236 bool = Companion_Default___.Fm30(Companion_Default___.SafeModuloInt(p0, _2018_i8), _1921_v0, (true) && (_1921_v0), globalState)
				_ = _rhs236
				var _rhs237 _dafny.Array = _2055_v93
				_ = _rhs237
				var _lhs128 _dafny.Array = _2051_v92
				_ = _lhs128
				var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_2051_v92), 0))
				_ = _lhs129
				var _lhs130 *C0 = _2033_v77
				_ = _lhs130
				(_lhs128).ArraySet1(_rhs233, (_lhs129).Int())
				_1921_v0 = _rhs234
				_lhs130.F8 = _rhs235
				_1921_v0 = _rhs236
				_2055_v93 = _rhs237
			}
			var _2056_v94 D4
			_ = _2056_v94
			_2056_v94 = Companion_D4_.Create_DC8_(_dafny.SeqOf(_1930_v5))
			_2056_v94 = _2056_v94
			_1921_v0 = _1921_v0
		}
		var _2057_v95 _dafny.Map
		_ = _2057_v95
		_2057_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1921_v0, p0)
		var _2058_v96 _dafny.Map
		_ = _2058_v96
		_2058_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2057_v95, _dafny.UnicodeSeqOfUtf8Bytes("fakxfloyk"))
		var _2059_v97 _dafny.Array
		_ = _2059_v97
		var _nwElement0_71 _dafny.Map = _2058_v96
		_ = _nwElement0_71
		var _nw321 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(3))
		_ = _nw321
		(_nw321).ArraySet1(_nwElement0_71, 0)
		(_nw321).ArraySet1((_2058_v96).Merge(_2058_v96), 1)
		(_nw321).ArraySet1((_2058_v96).Merge(_2058_v96), 2)
		_2059_v97 = _nw321
		var _2060_v98 D18
		_ = _2060_v98
		_2060_v98 = Companion_D18_.Create_DC47_(_2058_v96)
		var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_2059_v97), 0))
		_ = _index321
		(_2059_v97).ArraySet1((_2060_v98).Dtor_cf85(), (_index321).Int())
		var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_2059_v97), 0))
		_ = _index322
		(_2059_v97).ArraySet1(_2058_v96, (_index322).Int())
		var _2061_v99 _dafny.Map
		_ = _2061_v99
		_2061_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1921_v0)
		var _2062_v100 _dafny.Map
		_ = _2062_v100
		_2062_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(982), p0)
		var _2063_v101 _dafny.Set
		_ = _2063_v101
		_2063_v101 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_2061_v99).Cardinality()), _2062_v100, _2062_v100)
		var _2064_v102 _dafny.Map
		_ = _2064_v102
		_2064_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(816), _1921_v0)
		var _2065_v103 _dafny.Map
		_ = _2065_v103
		_2065_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((func() bool {
			if (_2064_v102).Contains(p0) {
				return (_2064_v102).Get(p0).(bool)
			}
			return _1921_v0
		})())).Cardinality(), _1921_v0)
		var _rhs238 bool = (_2063_v101).IsProperSubsetOf(_2063_v101)
		_ = _rhs238
		var _rhs239 bool = (_2065_v103).Contains((p0).Minus((_dafny.Zero).Minus(p0)))
		_ = _rhs239
		var _rhs240 bool = (func() bool {
			if (_1921_v0) && (_1921_v0) {
				return _1921_v0
			}
			return true
		})()
		_ = _rhs240
		_1921_v0 = _rhs238
		_1921_v0 = _rhs239
		_1921_v0 = _rhs240
		var _2066_i12 _dafny.Int
		_ = _2066_i12
		_2066_i12 = _dafny.Zero
		{
			for _1921_v0 {
				{
					if (_2066_i12).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_2066_i12 = (_2066_i12).Plus(_dafny.One)
					var _2067_v104 _dafny.Sequence
					_ = _2067_v104
					_2067_v104 = _dafny.SeqOf(_1921_v0)
					var _2068_v105 D7
					_ = _2068_v105
					_2068_v105 = Companion_D7_.Create_DC17_(_2067_v104)
					var _2069_v106 D4
					_ = _2069_v106
					_2069_v106 = Companion_D4_.Create_DC11_(p0, !_dafny.Companion_Sequence_.Equal((_2068_v105).Dtor_cf34(), _2067_v104), true, !(_1921_v0), _1921_v0)
					_2069_v106 = _2069_v106
					var _2070_v107 _dafny.Int
					_ = _2070_v107
					_2070_v107 = _dafny.IntOfInt64(399)
					_2070_v107 = _2070_v107
					var _2071_v108 D2
					_ = _2071_v108
					_2071_v108 = Companion_D2_.Create_DC5_(_2070_v107, (_this).F0(), _1921_v0)
					var _2072_v109 T1
					_ = _2072_v109
					var _nw322 *C5 = New_C5_()
					_ = _nw322
					_nw322.Ctor__(p1, (_2071_v108).Dtor_cf9(), _this.F1())
					_2072_v109 = _nw322
					var _2073_v110 D12
					_ = _2073_v110
					_2073_v110 = Companion_D12_.Create_DC29_(_1921_v0, _2070_v107, _1921_v0, _dafny.IntOfInt64(331), false)
					var _2074_v111 _dafny.Sequence
					_ = _2074_v111
					_2074_v111 = _dafny.SeqOf(_2072_v109, _2072_v109, _2072_v109, (func() T1 {
						if (_2073_v110).Dtor_cf52() {
							return _2072_v109
						}
						return _2072_v109
					})())
					_2072_v109 = (_2074_v111).Select((Companion_Default___.SafeIndex(((_this).Fm6(_2070_v107, _2070_v107, globalState)).Times(_2070_v107), _dafny.IntOfUint32((_2074_v111).Cardinality()))).Uint32()).(T1)
					var _2075_v112 *C2
					_ = _2075_v112
					var _nw323 *C2 = New_C2_()
					_ = _nw323
					_nw323.Ctor__()
					_2075_v112 = _nw323
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
	}
}
func (_this *C7) F4() _dafny.MultiSet {
	{
		return _this._f4
	}
}
func (_this *C7) F5() D0 {
	{
		return _this._f5
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	_f1 D0
	_f0 _dafny.Sequence
	_f2 _dafny.Int
	_f3 D0
}

func New_C8_() *C8 {
	_this := C8{}

	_this._f1 = Companion_D0_.Default()
	_this._f0 = _dafny.EmptySeq
	_this._f2 = _dafny.Zero
	_this._f3 = Companion_D0_.Default()
	return &_this
}

type CompanionStruct_C8_ struct {
}

var Companion_C8_ = CompanionStruct_C8_{}

func (_this *C8) Equals(other *C8) bool {
	return _this == other
}

func (_this *C8) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C8)
	return ok && _this.Equals(other)
}

func (*C8) String() string {
	return "_module.C8"
}

func Type_C8_() _dafny.TypeDescriptor {
	return type_C8_{}
}

type type_C8_ struct {
}

func (_this type_C8_) Default() interface{} {
	return (*C8)(nil)
}

func (_this type_C8_) String() string {
	return "main.C8"
}
func (_this *C8) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C8{}
var _ T1 = &C8{}
var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) F1() D0 {
	return _this._f1
}
func (_this *C8) F1_set_(value D0) {
	_this._f1 = value
}
func (_this *C8) F0() _dafny.Sequence {
	return _this._f0
}
func (_this *C8) Ctor__(f2 _dafny.Int, f3 D0, f0 _dafny.Sequence, f1 D0) {
	{
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f0 = f0
		(_this)._f1 = f1
	}
}
func (_this *C8) Fm5(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 bool, globalState *GlobalState) bool {
	{
		return !((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(191))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg114 _dafny.Int) interface{} {
				return coer114(arg114)
			}
		}(func(_2076_i0 _dafny.Int) _dafny.Int {
			return (_this).F2()
		})))).Contains(((_this).F0()).Select((Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int)))
	}
}
func (_this *C8) Fm6(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		if (_dafny.IntOfUint32(((_this).F0()).Cardinality())).Cmp((_this).F2()) >= 0 {
			return (_this).F2()
		} else {
			return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F2()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_this).F2()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F2()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F2()))).Cardinality())
		}
	}
}
func (_this *C8) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uyousyj"), _dafny.UnicodeSeqOfUtf8Bytes("b"))
	}
}
func (_this *C8) M1(p0 _dafny.Int, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _2077_v0 _dafny.CodePoint
		_ = _2077_v0
		_2077_v0 = _dafny.CodePoint('d')
		var _2078_v1 _dafny.Map
		_ = _2078_v1
		_2078_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2077_v0, p0)
		var _2079_v2 bool
		_ = _2079_v2
		_2079_v2 = true
		var _2080_v3 _dafny.Sequence
		_ = _2080_v3
		_2080_v3 = _dafny.UnicodeSeqOfUtf8Bytes("smiohn")
		var _2081_v5 _dafny.Map
		_ = _2081_v5
		_2081_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2077_v0, _2079_v2)
		var _2082_v6 _dafny.Array
		_ = _2082_v6
		var _nwElement0_72 _dafny.Map = (_2078_v1).Update(_2077_v0, (_this).F2())
		_ = _nwElement0_72
		var _nw324 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(25))
		_ = _nw324
		(_nw324).ArraySet1(_nwElement0_72, 0)
		(_nw324).ArraySet1(_2078_v1, 1)
		(_nw324).ArraySet1(_2078_v1, 2)
		(_nw324).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2077_v0, _dafny.IntOfInt64(409)), 3)
		(_nw324).ArraySet1(Companion_Default___.Fm52(_2079_v2, Companion_D15_.Create_DC38_(Companion_Default___.Fm51(true, _dafny.IntOfUint32((_2080_v3).Cardinality()), _dafny.IntOfInt64(535), globalState)), _2079_v2, globalState), 4)
		(_nw324).ArraySet1(_2078_v1, 5)
		(_nw324).ArraySet1(_2078_v1, 6)
		(_nw324).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2077_v0, _dafny.IntOfInt64(237)), 7)
		(_nw324).ArraySet1(_2078_v1, 8)
		(_nw324).ArraySet1(_2078_v1, 9)
		(_nw324).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2077_v0, (_this).F2()), 10)
		(_nw324).ArraySet1(_2078_v1, 11)
		(_nw324).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), (_this).F2())).Update(_2077_v0, (_dafny.Zero).Minus(p0)), 12)
		(_nw324).ArraySet1(_2078_v1, 13)
		(_nw324).ArraySet1(_2078_v1, 14)
		(_nw324).ArraySet1(_2078_v1, 15)
		(_nw324).ArraySet1(_2078_v1, 16)
		(_nw324).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2077_v0, p0), 17)
		(_nw324).ArraySet1(_2078_v1, 18)
		(_nw324).ArraySet1(func() _dafny.Map {
			var _coll80 = _dafny.NewMapBuilder()
			_ = _coll80
			for _iter83 := _dafny.Iterate((_2081_v5).Keys().Elements()); ; {
				_compr_80, _ok83 := _iter83()
				if !_ok83 {
					break
				}
				var _2083_v4 _dafny.CodePoint
				_2083_v4 = interface{}(_compr_80).(_dafny.CodePoint)
				if (_2081_v5).Contains(_2083_v4) {
					_coll80.Add(_2083_v4, p0)
				}
			}
			return _coll80.ToMap()
		}(), 19)
		(_nw324).ArraySet1(_2078_v1, 20)
		(_nw324).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm29(_2079_v2, globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(565))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg115 _dafny.Int) interface{} {
				return coer115(arg115)
			}
		}((func(_2084_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2085_i0 _dafny.Int) _dafny.CodePoint {
				return _2084_v0
			}
		})(_2077_v0))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(565))).Uint32(), func(coer116 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg116 _dafny.Int) interface{} {
				return coer116(arg116)
			}
		}((func(_2086_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2087_i0 _dafny.Int) _dafny.CodePoint {
				return _2086_v0
			}
		})(_2077_v0)))).Cardinality()))).Uint32(), _2077_v0)).Cardinality())), 21)
		(_nw324).ArraySet1((_2078_v1).Update(_2077_v0, (_this).F2()), 22)
		(_nw324).ArraySet1(_2078_v1, 23)
		(_nw324).ArraySet1(_2078_v1, 24)
		_2082_v6 = _nw324
		var _2088_v7 *C1
		_ = _2088_v7
		var _nw325 *C1 = New_C1_()
		_ = _nw325
		_nw325.Ctor__((_this).F2(), _2082_v6, (func() _dafny.Sequence {
			if _2079_v2 {
				return (_this).F0()
			}
			return (_this).F0()
		})(), _this.F1())
		_2088_v7 = _nw325
		var _2089_v8 _dafny.Array
		_ = _2089_v8
		var _len0_50 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_50
		var _nw326 _dafny.Array
		_ = _nw326
		if _len0_50.Cmp(_dafny.Zero) == 0 {
			_nw326 = _dafny.NewArray(_len0_50)
		} else {
			var _init50 func(_dafny.Int) _dafny.Int = (func(_2090_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2091_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_2091_i1, _2090_p0)
				}
			})(p0)
			_ = _init50
			var _element0_50 = _init50(_dafny.Zero)
			_ = _element0_50
			_nw326 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
			(_nw326).ArraySet1(_element0_50, 0)
			var _nativeLen0_50 = (_len0_50).Int()
			_ = _nativeLen0_50
			for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
				(_nw326).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
			}
		}
		_2089_v8 = _nw326
		var _2092_v9 _dafny.Map
		_ = _2092_v9
		_2092_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2080_v3, _2089_v8)
		var _2093_v10 _dafny.Map
		_ = _2093_v10
		_2093_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
			if (_2092_v9).Contains(_2080_v3) {
				return (_2092_v9).Get(_2080_v3).(_dafny.Array)
			}
			return _2089_v8
		})(), (p0).Plus(p0))
		_2093_v10 = (_2093_v10).Update(_2089_v8, (_2081_v5).Cardinality())
		var _2094_v11 _dafny.Sequence
		_ = _2094_v11
		_2094_v11 = _dafny.SeqOf(_2079_v2, false)
		var _2095_v12 _dafny.Sequence
		_ = _2095_v12
		_2095_v12 = _dafny.SeqOf(_2094_v11)
		_2094_v11 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2094_v11, (_2095_v12).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2095_v12).Cardinality()))).Uint32()).(_dafny.Sequence)), _2094_v11)
		var _2096_v13 _dafny.Map
		_ = _2096_v13
		_2096_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2079_v2)
		var _2097_v14 _dafny.Map
		_ = _2097_v14
		_2097_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2096_v13).Cardinality(), _2088_v7.F9)
		var _2098_v15 _dafny.Set
		_ = _2098_v15
		_2098_v15 = _dafny.SetOf((_this).F2())
		var _2099_v16 _dafny.Map
		_ = _2099_v16
		_2099_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2079_v2, _dafny.MultiSetOf((_this).F2(), (_2097_v14).Cardinality(), (_dafny.Zero).Minus((_2098_v15).Cardinality())))
		var _2100_v17 _dafny.MultiSet
		_ = _2100_v17
		_2100_v17 = _dafny.MultiSetOf(Companion_Default___.Fm44(_2079_v2, globalState))
		var _hi15 _dafny.Int = (Companion_Default___.Fm44(_2079_v2, globalState)).Plus(p0)
		_ = _hi15
		for _2101_i2 := ((func() _dafny.MultiSet {
			if (_2099_v16).Contains(_2079_v2) {
				return (_2099_v16).Get(_2079_v2).(_dafny.MultiSet)
			}
			return _2100_v17
		})()).Cardinality(); _2101_i2.Cmp(_hi15) < 0; _2101_i2 = _2101_i2.Plus(_dafny.One) {
			_2079_v2 = false
			if _2079_v2 {
				var _2102_v18 _dafny.Array
				_ = _2102_v18
				var _nwElement0_73 _dafny.Sequence = _2080_v3
				_ = _nwElement0_73
				var _nw327 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(3))
				_ = _nw327
				(_nw327).ArraySet1(_nwElement0_73, 0)
				(_nw327).ArraySet1(_2080_v3, 1)
				(_nw327).ArraySet1(_2080_v3, 2)
				_2102_v18 = _nw327
				var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_2102_v18), 0))
				_ = _index323
				(_2102_v18).ArraySet1(_2080_v3, (_index323).Int())
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_2102_v18), 0))
				_ = _index324
				(_2102_v18).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2080_v3, _2080_v3), (_index324).Int())
				var _2103_v19 _dafny.Map
				_ = _2103_v19
				_2103_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2079_v2, !(_2079_v2))
				(_2088_v7).F9 = (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2079_v2, true), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_2079_v2), _2079_v2)).Update(true, _2079_v2)).Merge(_2103_v19), _2103_v19, _2103_v19)).Cardinality())
				var _2104_v20 _dafny.Array
				_ = _2104_v20
				var _nwElement0_74 _dafny.Map = _2096_v13
				_ = _nwElement0_74
				var _nw328 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(6))
				_ = _nw328
				(_nw328).ArraySet1(_nwElement0_74, 0)
				(_nw328).ArraySet1((_2096_v13).Update((_2088_v7).Fm28(_2088_v7.F9, _2079_v2, globalState), _2079_v2), 1)
				(_nw328).ArraySet1(_2096_v13, 2)
				(_nw328).ArraySet1(_2096_v13, 3)
				(_nw328).ArraySet1(_2096_v13, 4)
				(_nw328).ArraySet1(Companion_Default___.Fm50(false, _2079_v2, _2101_i2, globalState), 5)
				_2104_v20 = _nw328
				var _2105_v21 _dafny.Array
				_ = _2105_v21
				var _nwElement0_75 _dafny.Array = _2104_v20
				_ = _nwElement0_75
				var _nw329 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(8))
				_ = _nw329
				(_nw329).ArraySet1(_nwElement0_75, 0)
				(_nw329).ArraySet1(_2104_v20, 1)
				(_nw329).ArraySet1(_2104_v20, 2)
				(_nw329).ArraySet1(_2104_v20, 3)
				(_nw329).ArraySet1(_2104_v20, 4)
				(_nw329).ArraySet1(_2104_v20, 5)
				(_nw329).ArraySet1(_2104_v20, 6)
				(_nw329).ArraySet1(_2104_v20, 7)
				_2105_v21 = _nw329
				var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_2105_v21), 0))
				_ = _index325
				(_2105_v21).ArraySet1(_2104_v20, (_index325).Int())
				var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_2105_v21), 0))
				_ = _index326
				(_2105_v21).ArraySet1(_2104_v20, (_index326).Int())
				r0 = _2079_v2
				var _2106_v22 _dafny.Set
				_ = _2106_v22
				_2106_v22 = _dafny.SetOf(_2079_v2)
				var _2107_v23 _dafny.Map
				_ = _2107_v23
				_2107_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2079_v2, (_2106_v22).Cardinality())
				r0 = !(_2107_v23).Contains(_2079_v2)
			} else {
				var _2108_v24 _dafny.Array
				_ = _2108_v24
				var _nw330 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw330
				_2108_v24 = _nw330
				var _2109_v25 _dafny.Array
				_ = _2109_v25
				var _len0_51 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_51
				var _nw331 _dafny.Array
				_ = _nw331
				if _len0_51.Cmp(_dafny.Zero) == 0 {
					_nw331 = _dafny.NewArray(_len0_51)
				} else {
					var _init51 func(_dafny.Int) bool = (func(_2110_v2 bool) func(_dafny.Int) bool {
						return func(_2111_i3 _dafny.Int) bool {
							return _2110_v2
						}
					})(_2079_v2)
					_ = _init51
					var _element0_51 = _init51(_dafny.Zero)
					_ = _element0_51
					_nw331 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
					(_nw331).ArraySet1(_element0_51, 0)
					var _nativeLen0_51 = (_len0_51).Int()
					_ = _nativeLen0_51
					for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
						(_nw331).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
					}
				}
				_2109_v25 = _nw331
				var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_2108_v24), 0))
				_ = _index327
				(_2108_v24).ArraySet1(_2109_v25, (_index327).Int())
				var _2112_v26 _dafny.Array
				_ = _2112_v26
				var _nw332 _dafny.Array = _dafny.NewArrayWithValue(Companion_D14_.Default(), _dafny.IntOfInt64(2))
				_ = _nw332
				_2112_v26 = _nw332
				var _2113_v27 D14
				_ = _2113_v27
				_2113_v27 = Companion_D14_.Create_DC35_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(630))).Uint32(), func(coer117 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg117 _dafny.Int) interface{} {
						return coer117(arg117)
					}
				}((func(_2114_v7 *C1) func(_dafny.Int) _dafny.Int {
					return func(_2115_i4 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_2114_v7.F9)
					}
				})(_2088_v7)))).Cardinality()), _2080_v3, _2101_i2, true, _2079_v2)
				var _pat_let_tv50 = _2077_v0
				_ = _pat_let_tv50
				var _pat_let_tv51 = _2096_v13
				_ = _pat_let_tv51
				var _pat_let_tv52 = _2088_v7
				_ = _pat_let_tv52
				var _pat_let_tv53 = _2096_v13
				_ = _pat_let_tv53
				var _pat_let_tv54 = _2088_v7
				_ = _pat_let_tv54
				var _pat_let_tv55 = _2079_v2
				_ = _pat_let_tv55
				var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_2112_v26), 0))
				_ = _index328
				(_2112_v26).ArraySet1(func(_pat_let44_0 D14) D14 {
					return func(_2116_dt__update__tmp_h0 D14) D14 {
						return func(_pat_let45_0 _dafny.Sequence) D14 {
							return func(_2119_dt__update_hcf64_h0 _dafny.Sequence) D14 {
								return func(_pat_let46_0 bool) D14 {
									return func(_2120_dt__update_hcf67_h0 bool) D14 {
										return Companion_D14_.Create_DC35_((_2116_dt__update__tmp_h0).Dtor_cf63(), _2119_dt__update_hcf64_h0, (_2116_dt__update__tmp_h0).Dtor_cf65(), (_2116_dt__update__tmp_h0).Dtor_cf66(), _2120_dt__update_hcf67_h0)
									}(_pat_let46_0)
								}((func() bool {
									if (_pat_let_tv51).Contains(_pat_let_tv52.F9) {
										return (_pat_let_tv53).Get(_pat_let_tv54.F9).(bool)
									}
									return _pat_let_tv55
								})())
							}(_pat_let45_0)
						}(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(892))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg118 _dafny.Int) interface{} {
								return coer118(arg118)
							}
						}((func(_2117_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_2118_i5 _dafny.Int) _dafny.CodePoint {
								return _2117_v0
							}
						})(_pat_let_tv50))))
					}(_pat_let44_0)
				}(_2113_v27), (_index328).Int())
				var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_2109_v25), 0))
				_ = _index329
				(_2109_v25).ArraySet1(_2079_v2, (_index329).Int())
				var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_2108_v24), 0))
				_ = _index330
				var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_2112_v26), 0))
				_ = _index331
				var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_2109_v25), 0))
				_ = _index332
				var _rhs241 _dafny.Array = _2109_v25
				_ = _rhs241
				var _rhs242 D14 = _2113_v27
				_ = _rhs242
				var _rhs243 bool = _2079_v2
				_ = _rhs243
				var _rhs244 _dafny.Int = (_this).F2()
				_ = _rhs244
				var _rhs245 bool = Companion_Default___.Fm17(globalState)
				_ = _rhs245
				var _lhs131 _dafny.Array = _2108_v24
				_ = _lhs131
				var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_2108_v24), 0))
				_ = _lhs132
				var _lhs133 _dafny.Array = _2112_v26
				_ = _lhs133
				var _lhs134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_2112_v26), 0))
				_ = _lhs134
				var _lhs135 _dafny.Array = _2109_v25
				_ = _lhs135
				var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_2109_v25), 0))
				_ = _lhs136
				var _lhs137 *C1 = _2088_v7
				_ = _lhs137
				(_lhs131).ArraySet1(_rhs241, (_lhs132).Int())
				(_lhs133).ArraySet1(_rhs242, (_lhs134).Int())
				(_lhs135).ArraySet1(_rhs243, (_lhs136).Int())
				_lhs137.F9 = _rhs244
				r1 = _rhs245
				(_2088_v7).F9 = (func() _dafny.Int {
					if _2079_v2 {
						return ((Companion_Default___.Fm13(_2077_v0, true, globalState)).Intersection(_2100_v17)).Cardinality()
					}
					return _2088_v7.F9
				})()
				var _2121_v28 _dafny.Int
				_ = _2121_v28
				var _2122_v29 _dafny.Map
				_ = _2122_v29
				var _2123_v30 _dafny.Int
				_ = _2123_v30
				var _out49 _dafny.Int
				_ = _out49
				var _out50 _dafny.Map
				_ = _out50
				var _out51 _dafny.Int
				_ = _out51
				_out49, _out50, _out51 = (_this).M3(_2101_i2, (_this).F3(), Companion_Default___.Fm30(p0, (_2109_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_2109_v25), 0))).Int()).(bool), (Companion_D18_.Create_DC50_((_2109_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_2109_v25), 0))).Int()).(bool), _2109_v25, _2079_v2, (_2109_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_2109_v25), 0))).Int()).(bool), _2080_v3)).Dtor_cf90(), globalState), globalState)
				_2121_v28 = _out49
				_2122_v29 = _out50
				_2123_v30 = _out51
				(_2088_v7).F9 = Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(811)).Minus(((_this).F0()).Select((Companion_Default___.SafeIndex(_2101_i2, _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int)), (func() _dafny.Int {
					if (_2100_v17).Contains(_2121_v28) {
						return (_2100_v17).Multiplicity(_2121_v28)
					}
					return _2088_v7.F9
				})())
				(_2088_v7).F9 = (_this).Fm6(_2101_i2, (_dafny.Zero).Minus((func() _dafny.Int {
					if _2079_v2 {
						return p0
					}
					return _dafny.IntOfInt64(877)
				})()), globalState)
			}
			var _2124_v31 _dafny.Array
			_ = _2124_v31
			var _len0_52 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_52
			var _nw333 _dafny.Array
			_ = _nw333
			if _len0_52.Cmp(_dafny.Zero) == 0 {
				_nw333 = _dafny.NewArray(_len0_52)
			} else {
				var _init52 func(_dafny.Int) _dafny.CodePoint = (func(_2125_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2126_i6 _dafny.Int) _dafny.CodePoint {
						return _2125_v0
					}
				})(_2077_v0)
				_ = _init52
				var _element0_52 = _init52(_dafny.Zero)
				_ = _element0_52
				_nw333 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
				(_nw333).ArraySet1CodePoint(_element0_52, 0)
				var _nativeLen0_52 = (_len0_52).Int()
				_ = _nativeLen0_52
				for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
					(_nw333).ArraySet1CodePoint(_init52(_dafny.IntOf(_i0_52)), _i0_52)
				}
			}
			_2124_v31 = _nw333
			var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_2124_v31), 0))
			_ = _index333
			(_2124_v31).ArraySet1CodePoint(_2077_v0, (_index333).Int())
			var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_2124_v31), 0))
			_ = _index334
			(_2124_v31).ArraySet1CodePoint(_2077_v0, (_index334).Int())
			var _2127_v32 _dafny.Map
			_ = _2127_v32
			_2127_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2079_v2, _2080_v3)
			var _2128_v33 _dafny.Sequence
			_ = _2128_v33
			_2128_v33 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_2127_v32).Contains(_2079_v2) {
					return (_2127_v32).Get(_2079_v2).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(526))).Uint32(), func(coer119 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg119 _dafny.Int) interface{} {
						return coer119(arg119)
					}
				}(func(_2129_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				}))
			})(), (Companion_Default___.SafeIndex(_2088_v7.F9, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_2127_v32).Contains(_2079_v2) {
					return (_2127_v32).Get(_2079_v2).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(526))).Uint32(), func(coer120 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg120 _dafny.Int) interface{} {
						return coer120(arg120)
					}
				}(func(_2130_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				}))
			})()).Cardinality()))).Uint32(), _2077_v0))
			_2128_v33 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_2080_v3, _dafny.UnicodeSeqOfUtf8Bytes("chsubpa")))
		}
		var _2131_v34 _dafny.Map
		_ = _2131_v34
		_2131_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2079_v2, _2077_v0)
		_2131_v34 = (_2131_v34).Update(false, _2077_v0)
		var _2132_v36 _dafny.Map
		_ = _2132_v36
		_2132_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2079_v2, p0)
		var _2133_v37 D18
		_ = _2133_v37
		_2133_v37 = Companion_D18_.Create_DC47_(func() _dafny.Map {
			var _coll81 = _dafny.NewMapBuilder()
			_ = _coll81
			for _iter84 := _dafny.Iterate((_dafny.SeqOf(_2132_v36, _2132_v36)).Elements()); ; {
				_compr_81, _ok84 := _iter84()
				if !_ok84 {
					break
				}
				var _2134_v35 _dafny.Map
				_2134_v35 = interface{}(_compr_81).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_2132_v36, _2132_v36), _2134_v35) {
					_coll81.Add(_2134_v35, _dafny.UnicodeSeqOfUtf8Bytes("w"))
				}
			}
			return _coll81.ToMap()
		}())
		var _pat_let_tv56 = _2079_v2
		_ = _pat_let_tv56
		var _pat_let_tv57 = p0
		_ = _pat_let_tv57
		var _pat_let_tv58 = _2098_v15
		_ = _pat_let_tv58
		var _pat_let_tv59 = _2098_v15
		_ = _pat_let_tv59
		var _rhs246 bool = func(_source22 D18) bool {
			if _source22.Is_DC48() {
				return _pat_let_tv56
			} else if _source22.Is_DC49() {
				var _2135___mcc_h0 _dafny.Set = _source22.Get_().(D18_DC49).Cf86
				_ = _2135___mcc_h0
				var _2136_cf86 _dafny.Set = _2135___mcc_h0
				_ = _2136_cf86
				return (_pat_let_tv57).Cmp((_this).F2()) > 0
			} else if _source22.Is_DC50() {
				var _2137___mcc_h1 bool = _source22.Get_().(D18_DC50).Cf87
				_ = _2137___mcc_h1
				var _2138___mcc_h2 _dafny.Array = _source22.Get_().(D18_DC50).Cf88
				_ = _2138___mcc_h2
				var _2139___mcc_h3 bool = _source22.Get_().(D18_DC50).Cf89
				_ = _2139___mcc_h3
				var _2140___mcc_h4 bool = _source22.Get_().(D18_DC50).Cf90
				_ = _2140___mcc_h4
				var _2141___mcc_h5 _dafny.Sequence = _source22.Get_().(D18_DC50).Cf91
				_ = _2141___mcc_h5
				var _2142_cf91 _dafny.Sequence = _2141___mcc_h5
				_ = _2142_cf91
				var _2143_cf90 bool = _2140___mcc_h4
				_ = _2143_cf90
				var _2144_cf89 bool = _2139___mcc_h3
				_ = _2144_cf89
				var _2145_cf88 _dafny.Array = _2138___mcc_h2
				_ = _2145_cf88
				var _2146_cf87 bool = _2137___mcc_h1
				_ = _2146_cf87
				return _2143_cf90
			} else {
				var _2147___mcc_h6 _dafny.Map = _source22.Get_().(D18_DC47).Cf85
				_ = _2147___mcc_h6
				var _2148_cf85 _dafny.Map = _2147___mcc_h6
				_ = _2148_cf85
				return (_pat_let_tv58).IsProperSubsetOf(_pat_let_tv59)
			}
		}(_2133_v37)
		_ = _rhs246
		var _rhs247 bool = !((p0).Cmp((_this).F2()) > 0)
		_ = _rhs247
		r1 = _rhs246
		r1 = _rhs247
		r0 = (func() bool {
			if _2079_v2 {
				return false
			}
			return _2079_v2
		})()
		r1 = Companion_Default___.Fm17(globalState)
		return r0, r1
	}
}
func (_this *C8) M2(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _2149_v0 bool
		_ = _2149_v0
		_2149_v0 = true
		_2149_v0 = _2149_v0
		var _2150_v1 _dafny.Array
		_ = _2150_v1
		var _nw334 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
		_ = _nw334
		_2150_v1 = _nw334
		var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_2150_v1), 0))
		_ = _index335
		(_2150_v1).ArraySet1(_2149_v0, (_index335).Int())
		var _2151_v2 bool
		_ = _2151_v2
		_2151_v2 = _2149_v0
		var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_2150_v1), 0))
		_ = _index336
		(_2150_v1).ArraySet1(_2151_v2, (_index336).Int())
		if (p0).Cmp(p0) <= 0 {
			var _2152_v3 _dafny.Int
			_ = _2152_v3
			_2152_v3 = _dafny.IntOfInt64(14)
			var _2153_v4 _dafny.Array
			_ = _2153_v4
			var _nw335 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
			_ = _nw335
			_2153_v4 = _nw335
			var _2154_v5 _dafny.MultiSet
			_ = _2154_v5
			_2154_v5 = _dafny.MultiSetOf(_2153_v4)
			var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_2153_v4), 0))
			_ = _index337
			(_2153_v4).ArraySet1(!(!(_2149_v0)), (_index337).Int())
			var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_2153_v4), 0))
			_ = _index338
			var _rhs248 _dafny.Int = (_this).F2()
			_ = _rhs248
			var _rhs249 _dafny.MultiSet = _2154_v5
			_ = _rhs249
			var _rhs250 bool = !((_2152_v3).Cmp(_2152_v3) == 0)
			_ = _rhs250
			var _rhs251 bool = !(!(_2149_v0))
			_ = _rhs251
			var _lhs138 _dafny.Array = _2153_v4
			_ = _lhs138
			var _lhs139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_2153_v4), 0))
			_ = _lhs139
			_2152_v3 = _rhs248
			_2154_v5 = _rhs249
			(_lhs138).ArraySet1(_rhs250, (_lhs139).Int())
			_2149_v0 = _rhs251
			var _2155_v6 _dafny.Map
			_ = _2155_v6
			_2155_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_2153_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_2153_v4), 0))).Int()).(bool))
			var _2156_v9 _dafny.Map
			_ = _2156_v9
			_2156_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2149_v0, (_2153_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_2153_v4), 0))).Int()).(bool))
			var _2157_v10 _dafny.Set
			_ = _2157_v10
			_2157_v10 = _dafny.SetOf(_2155_v6, _2155_v6, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2156_v9).Cardinality(), _2149_v0))
			var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_2153_v4), 0))
			_ = _index339
			(_2153_v4).ArraySet1((_dafny.SetOf(_2155_v6, func() _dafny.Map {
				var _coll82 = _dafny.NewMapBuilder()
				_ = _coll82
				for _iter85 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(567), _dafny.IntOfInt64(-187))); ; {
					_compr_82, _ok85 := _iter85()
					if !_ok85 {
						break
					}
					var _2158_v7 _dafny.Int
					_2158_v7 = interface{}(_compr_82).(_dafny.Int)
					if ((_dafny.IntOfInt64(567)).Cmp(_2158_v7) <= 0) && ((_2158_v7).Cmp(_dafny.IntOfInt64(-187)) < 0) {
						_coll82.Add((_2158_v7).Minus((_this).F2()), _2149_v0)
					}
				}
				return _coll82.ToMap()
			}(), func() _dafny.Map {
				var _coll83 = _dafny.NewMapBuilder()
				_ = _coll83
				for _iter86 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-911), _dafny.IntOfInt64(126))); ; {
					_compr_83, _ok86 := _iter86()
					if !_ok86 {
						break
					}
					var _2159_v8 _dafny.Int
					_2159_v8 = interface{}(_compr_83).(_dafny.Int)
					if ((_dafny.IntOfInt64(-911)).Cmp(_2159_v8) <= 0) && ((_2159_v8).Cmp(_dafny.IntOfInt64(126)) < 0) {
						_coll83.Add((_2159_v8).Minus((_this).F2()), !((_2153_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_2153_v4), 0))).Int()).(bool)))
					}
				}
				return _coll83.ToMap()
			}())).IsDisjointFrom(_2157_v10), (_index339).Int())
			var _2160_v11 _dafny.Map
			_ = _2160_v11
			_2160_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((_2154_v5).Union(_2154_v5)).Cardinality()), p0)
			_2160_v11 = (_2160_v11).Update((_dafny.IntOfInt64(-793)).Plus((_dafny.SetOf((_this).F2())).Cardinality()), (_this).F2())
			var _2161_v12 _dafny.Sequence
			_ = _2161_v12
			_2161_v12 = _dafny.UnicodeSeqOfUtf8Bytes("i")
			var _2162_v13 _dafny.Map
			_ = _2162_v13
			_2162_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2152_v3).Cmp(_2152_v3) >= 0, _2161_v12)
			var _2163_v14 _dafny.CodePoint
			_ = _2163_v14
			_2163_v14 = _dafny.CodePoint('u')
			_2162_v13 = (_2162_v13).Update(_2149_v0, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_2161_v12, (Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32((_2161_v12).Cardinality()))).Uint32(), _2163_v14), _2161_v12))
			var _2164_v15 _dafny.Sequence
			_ = _2164_v15
			_2164_v15 = _dafny.SeqOf(_2149_v0)
			var _2165_v16 _dafny.MultiSet
			_ = _2165_v16
			_2165_v16 = _dafny.MultiSetOf(_2152_v3)
			var _2166_v17 _dafny.Set
			_ = _2166_v17
			_2166_v17 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-438))).Uint32(), func(coer121 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg121 _dafny.Int) interface{} {
					return coer121(arg121)
				}
			}((func(_2167_v14 _dafny.CodePoint) func(_dafny.Int) _dafny.Int {
				return func(_2168_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(492))).Uint32(), func(coer122 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg122 _dafny.Int) interface{} {
							return coer122(arg122)
						}
					}((func(_2169_v14 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2170_i1 _dafny.Int) _dafny.CodePoint {
							return _2169_v14
						}
					})(_2167_v14)))).Cardinality())
				}
			})(_2163_v14)))).Cardinality()), _2152_v3, (_this).F2())
			var _2171_v18 _dafny.Set
			_ = _2171_v18
			_2171_v18 = _dafny.SetOf((_2166_v17).Cardinality(), (_this).F2())
			var _2172_v19 _dafny.Array
			_ = _2172_v19
			var _nwElement0_76 _dafny.Int = p0
			_ = _nwElement0_76
			var _nw336 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(27))
			_ = _nw336
			(_nw336).ArraySet1(_nwElement0_76, 0)
			(_nw336).ArraySet1(_2152_v3, 1)
			(_nw336).ArraySet1(p0, 2)
			(_nw336).ArraySet1((_this).F2(), 3)
			(_nw336).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_2152_v3, p0)).Cardinality()), 4)
			(_nw336).ArraySet1(_dafny.IntOfUint32((_2164_v15).Cardinality()), 5)
			(_nw336).ArraySet1(_2152_v3, 6)
			(_nw336).ArraySet1(_dafny.IntOfInt64(292), 7)
			(_nw336).ArraySet1((_this).F2(), 8)
			(_nw336).ArraySet1((_dafny.Zero).Minus(p0), 9)
			(_nw336).ArraySet1(p0, 10)
			(_nw336).ArraySet1((_this).F2(), 11)
			(_nw336).ArraySet1(p0, 12)
			(_nw336).ArraySet1(p0, 13)
			(_nw336).ArraySet1(_2152_v3, 14)
			(_nw336).ArraySet1((_this).F2(), 15)
			(_nw336).ArraySet1((_this).F2(), 16)
			(_nw336).ArraySet1(p0, 17)
			(_nw336).ArraySet1(_dafny.IntOfUint32((_2161_v12).Cardinality()), 18)
			(_nw336).ArraySet1((_dafny.MultiSetFromSeq(_2164_v15)).Cardinality(), 19)
			(_nw336).ArraySet1(_2152_v3, 20)
			(_nw336).ArraySet1(_2152_v3, 21)
			(_nw336).ArraySet1((_2160_v11).Cardinality(), 22)
			(_nw336).ArraySet1((_2165_v16).Cardinality(), 23)
			(_nw336).ArraySet1((_2166_v17).Cardinality(), 24)
			(_nw336).ArraySet1(_2152_v3, 25)
			(_nw336).ArraySet1(_2152_v3, 26)
			_2172_v19 = _nw336
			var _pat_let_tv60 = _2152_v3
			_ = _pat_let_tv60
			var _2173_v20 *C5
			_ = _2173_v20
			var _nw337 *C5 = New_C5_()
			_ = _nw337
			_nw337.Ctor__(_2172_v19, (Companion_D2_.Create_DC5_(p0, (_this).F0(), (_2153_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_2153_v4), 0))).Int()).(bool))).Dtor_cf9(), func(_pat_let47_0 D0) D0 {
				return func(_2174_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let48_0 _dafny.Int) D0 {
						return func(_2175_dt__update_hcf2_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC1_(_2175_dt__update_hcf2_h0)
						}(_pat_let48_0)
					}((_dafny.Zero).Minus((_dafny.Zero).Minus(_pat_let_tv60)))
				}(_pat_let47_0)
			}(_this.F1()))
			_2173_v20 = _nw337
			var _2176_v21 _dafny.MultiSet
			_ = _2176_v21
			_2176_v21 = _dafny.MultiSetOf(_2173_v20, _2173_v20)
			_2176_v21 = _2176_v21
		} else {
			var _2177_v22 _dafny.Int
			_ = _2177_v22
			_2177_v22 = _dafny.IntOfInt64(150)
			_2177_v22 = (func() _dafny.Int {
				if !(true) {
					return _2177_v22
				}
				return (_this).F2()
			})()
			var _2178_v23 _dafny.Sequence
			_ = _2178_v23
			_2178_v23 = _dafny.UnicodeSeqOfUtf8Bytes("uefbomho")
			var _2179_v24 D6
			_ = _2179_v24
			_2179_v24 = Companion_D6_.Create_DC14_(_2178_v23)
			var _source23 D6 = _2179_v24
			_ = _source23
			if _source23.Is_DC15() {
				var _2180___mcc_h0 _dafny.Int = _source23.Get_().(D6_DC15).Cf30
				_ = _2180___mcc_h0
				var _2181___mcc_h1 _dafny.Int = _source23.Get_().(D6_DC15).Cf31
				_ = _2181___mcc_h1
				var _2182___mcc_h2 bool = _source23.Get_().(D6_DC15).Cf32
				_ = _2182___mcc_h2
				var _2183_cf32 bool = _2182___mcc_h2
				_ = _2183_cf32
				var _2184_cf31 _dafny.Int = _2181___mcc_h1
				_ = _2184_cf31
				var _2185_cf30 _dafny.Int = _2180___mcc_h0
				_ = _2185_cf30
				var _2186_v25 _dafny.Array
				_ = _2186_v25
				var _len0_53 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_53
				var _nw338 _dafny.Array
				_ = _nw338
				if _len0_53.Cmp(_dafny.Zero) == 0 {
					_nw338 = _dafny.NewArray(_len0_53)
				} else {
					var _init53 func(_dafny.Int) _dafny.Set = (func(_2187_cf32 bool) func(_dafny.Int) _dafny.Set {
						return func(_2188_i2 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(!(_2187_cf32), _2187_cf32)
						}
					})(_2183_cf32)
					_ = _init53
					var _element0_53 = _init53(_dafny.Zero)
					_ = _element0_53
					_nw338 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
					(_nw338).ArraySet1(_element0_53, 0)
					var _nativeLen0_53 = (_len0_53).Int()
					_ = _nativeLen0_53
					for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
						(_nw338).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
					}
				}
				_2186_v25 = _nw338
				var _2189_v26 _dafny.Map
				_ = _2189_v26
				_2189_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if false {
						return _2177_v22
					}
					return _2177_v22
				})(), _2186_v25)
				_2189_v26 = (_2189_v26).Update(_2177_v22, _2186_v25)
				_2183_cf32 = true
				var _2190_v27 _dafny.Array
				_ = _2190_v27
				var _nw339 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
				_ = _nw339
				_2190_v27 = _nw339
				var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_2190_v27), 0))
				_ = _index340
				(_2190_v27).ArraySet1(_2149_v0, (_index340).Int())
				var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_2190_v27), 0))
				_ = _index341
				(_2190_v27).ArraySet1(_2183_cf32, (_index341).Int())
				var _2191_v28 _dafny.Sequence
				_ = _2191_v28
				_2191_v28 = _dafny.SeqOf(_2183_cf32, _2149_v0, _2183_cf32, _2183_cf32)
				var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_2190_v27), 0))
				_ = _index342
				var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_2190_v27), 0))
				_ = _index343
				var _rhs252 bool = (func() bool {
					if (((_this).F0()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2178_v23).Cardinality()), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(p0) >= 0 {
						return _2149_v0
					}
					return (_2191_v28).Select((Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32((_2191_v28).Cardinality()))).Uint32()).(bool)
				})()
				_ = _rhs252
				var _rhs253 bool = _2149_v0
				_ = _rhs253
				var _rhs254 bool = _2183_cf32
				_ = _rhs254
				var _rhs255 bool = true
				_ = _rhs255
				var _rhs256 bool = _2149_v0
				_ = _rhs256
				var _lhs140 _dafny.Array = _2190_v27
				_ = _lhs140
				var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_2190_v27), 0))
				_ = _lhs141
				var _lhs142 _dafny.Array = _2190_v27
				_ = _lhs142
				var _lhs143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_2190_v27), 0))
				_ = _lhs143
				(_lhs140).ArraySet1(_rhs252, (_lhs141).Int())
				(_lhs142).ArraySet1(_rhs253, (_lhs143).Int())
				_2149_v0 = _rhs254
				_2149_v0 = _rhs255
				_2149_v0 = _rhs256
				var _2192_v29 _dafny.MultiSet
				_ = _2192_v29
				_2192_v29 = _dafny.MultiSetOf(_2149_v0)
				(_this).M4(_2149_v0, _dafny.IntOfInt64(105), _2192_v29, p0, globalState)
			} else if _source23.Is_DC14() {
				var _2193___mcc_h3 _dafny.Sequence = _source23.Get_().(D6_DC14).Cf29
				_ = _2193___mcc_h3
				var _2194_cf29 _dafny.Sequence = _2193___mcc_h3
				_ = _2194_cf29
				var _2195_v30 _dafny.Array
				_ = _2195_v30
				var _nw340 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw340
				_2195_v30 = _nw340
				var _2196_v31 D18
				_ = _2196_v31
				_2196_v31 = Companion_D18_.Create_DC50_(!(_2149_v0), _2195_v30, false, false, _2194_cf29)
				var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))
				_ = _index344
				(_2195_v30).ArraySet1((_2196_v31).Dtor_cf87(), (_index344).Int())
				var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))
				_ = _index345
				(_2195_v30).ArraySet1(_2149_v0, (_index345).Int())
				var _2197_v32 *C3
				_ = _2197_v32
				var _nw341 *C3 = New_C3_()
				_ = _nw341
				_nw341.Ctor__()
				_2197_v32 = _nw341
				var _2198_v33 _dafny.CodePoint
				_ = _2198_v33
				_2198_v33 = _dafny.CodePoint('m')
				var _2199_v34 _dafny.MultiSet
				_ = _2199_v34
				_2199_v34 = _dafny.MultiSetOf(_2198_v33, _2198_v33)
				var _2200_v35 _dafny.Sequence
				_ = _2200_v35
				_2200_v35 = _dafny.SeqOf(_2199_v34)
				var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))
				_ = _index346
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))
				_ = _index347
				var _rhs257 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_2200_v35, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.IntOfUint32((_2200_v35).Cardinality()))).Uint32(), _2199_v34), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(127))).Uint32(), func(coer123 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg123 _dafny.Int) interface{} {
						return coer123(arg123)
					}
				}((func(_2201_v34 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_2202_i3 _dafny.Int) _dafny.MultiSet {
						return _2201_v34
					}
				})(_2199_v34))))
				_ = _rhs257
				var _rhs258 bool = (_2195_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))).Int()).(bool)
				_ = _rhs258
				var _rhs259 bool = !((((_this).F2()).Cmp(_dafny.IntOfUint32((_2194_cf29).Cardinality())) > 0) && (!((_2195_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))).Int()).(bool))))
				_ = _rhs259
				var _rhs260 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).Fm6((_this).F2(), (_this).F2(), globalState)), _dafny.IntOfInt64(833))).Minus((_this).F2())))
				_ = _rhs260
				var _lhs144 _dafny.Array = _2195_v30
				_ = _lhs144
				var _lhs145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))
				_ = _lhs145
				var _lhs146 _dafny.Array = _2195_v30
				_ = _lhs146
				var _lhs147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))
				_ = _lhs147
				(_lhs144).ArraySet1(_rhs257, (_lhs145).Int())
				(_lhs146).ArraySet1(_rhs258, (_lhs147).Int())
				_2149_v0 = _rhs259
				_2177_v22 = _rhs260
				var _2203_v36 _dafny.Set
				_ = _2203_v36
				_2203_v36 = _dafny.SetOf(_dafny.CodePoint('s'))
				var _2204_v37 _dafny.Set
				_ = _2204_v37
				_2204_v37 = _dafny.SetOf(_2177_v22, (_this).F2(), _dafny.IntOfInt64(25))
				var _2205_v38 _dafny.Set
				_ = _2205_v38
				_2205_v38 = _dafny.SetOf(_2149_v0)
				var _2206_v39 _dafny.Map
				_ = _2206_v39
				_2206_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_2177_v22), (_dafny.Zero).Minus((_this).F2()))
				var _2207_v40 _dafny.Map
				_ = _2207_v40
				_2207_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2198_v33, _2198_v33)
				var _2208_v41 _dafny.Sequence
				_ = _2208_v41
				_2208_v41 = _dafny.SeqOf(_2207_v40, _2207_v40)
				var _2209_v42 _dafny.MultiSet
				_ = _2209_v42
				_2209_v42 = _dafny.MultiSetOf(_2149_v0)
				var _2210_v43 _dafny.Sequence
				_ = _2210_v43
				_2210_v43 = _dafny.SeqOf(false)
				var _2211_v44 _dafny.Array
				_ = _2211_v44
				var _nwElement0_77 _dafny.Int = _dafny.IntOfInt64(994)
				_ = _nwElement0_77
				var _nw342 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(25))
				_ = _nw342
				(_nw342).ArraySet1(_nwElement0_77, 0)
				(_nw342).ArraySet1(_2177_v22, 1)
				(_nw342).ArraySet1((_this).F2(), 2)
				(_nw342).ArraySet1(_2177_v22, 3)
				(_nw342).ArraySet1(_dafny.IntOfInt64(314), 4)
				(_nw342).ArraySet1((_2205_v38).Cardinality(), 5)
				(_nw342).ArraySet1((_this).F2(), 6)
				(_nw342).ArraySet1(_2177_v22, 7)
				(_nw342).ArraySet1((func() _dafny.Int {
					if (_2206_v39).Contains(p0) {
						return (_2206_v39).Get(p0).(_dafny.Int)
					}
					return Companion_Default___.Fm44(_2149_v0, globalState)
				})(), 8)
				(_nw342).ArraySet1(p0, 9)
				(_nw342).ArraySet1(_dafny.IntOfUint32(((_this).F0()).Cardinality()), 10)
				(_nw342).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32(((_this).F0()).Cardinality())), 11)
				(_nw342).ArraySet1((_this).F2(), 12)
				(_nw342).ArraySet1(((_2208_v41).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_2209_v42).Cardinality()), _dafny.IntOfUint32((_2208_v41).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), 13)
				(_nw342).ArraySet1(_dafny.IntOfUint32((_2194_cf29).Cardinality()), 14)
				(_nw342).ArraySet1(_2177_v22, 15)
				(_nw342).ArraySet1((_this).F2(), 16)
				(_nw342).ArraySet1(_dafny.IntOfUint32(((_this).F0()).Cardinality()), 17)
				(_nw342).ArraySet1(p0, 18)
				(_nw342).ArraySet1(_2177_v22, 19)
				(_nw342).ArraySet1(_dafny.IntOfInt64(266), 20)
				(_nw342).ArraySet1(_2177_v22, 21)
				(_nw342).ArraySet1((_dafny.MultiSetFromSeq(_2210_v43)).Cardinality(), 22)
				(_nw342).ArraySet1((_this).F2(), 23)
				(_nw342).ArraySet1(p0, 24)
				_2211_v44 = _nw342
				var _2212_v45 _dafny.Sequence
				_ = _2212_v45
				_2212_v45 = _dafny.SeqOf(_2211_v44, _2211_v44, _2211_v44)
				var _2213_v46 _dafny.Map
				_ = _2213_v46
				_2213_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2195_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))).Int()).(bool), p0)
				var _2214_v47 _dafny.Map
				_ = _2214_v47
				_2214_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2212_v45).Select((Companion_Default___.SafeIndex((_2213_v46).Cardinality(), _dafny.IntOfUint32((_2212_v45).Cardinality()))).Uint32()).(_dafny.Array), (_this).F2())
				var _2215_v48 _dafny.MultiSet
				_ = _2215_v48
				_2215_v48 = _dafny.MultiSetOf(_dafny.IntOfUint32((Companion_Default___.Fm25(Companion_Default___.Fm64(globalState), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm16(_2178_v23, (_dafny.MultiSetOf((_this).F2(), _dafny.IntOfUint32((_2210_v43).Cardinality()), _2177_v22, (_this).F2(), (_2204_v37).Cardinality())).Cardinality(), globalState), (Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32((Companion_Default___.Fm16(_2178_v23, (_dafny.MultiSetOf((_this).F2(), _dafny.IntOfUint32((_2210_v43).Cardinality()), _2177_v22, (_this).F2(), (_2204_v37).Cardinality())).Cardinality(), globalState)).Cardinality()))).Uint32(), _2198_v33), globalState)).Cardinality()), _2177_v22)
				var _2216_v49 _dafny.Map
				_ = _2216_v49
				_2216_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2198_v33, p0)
				var _2217_v50 _dafny.Map
				_ = _2217_v50
				_2217_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_2195_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))).Int()).(bool))
				var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))
				_ = _index348
				var _rhs261 _dafny.Set = _dafny.SetOf(_2198_v33, _2198_v33, _2198_v33, _2198_v33, _2198_v33)
				_ = _rhs261
				var _rhs262 bool = _2149_v0
				_ = _rhs262
				var _rhs263 _dafny.Set = (_dafny.SetOf((_2215_v48).Cardinality(), _dafny.IntOfInt64(12), _dafny.IntOfInt64(535), (_2216_v49).Cardinality())).Difference((func() _dafny.Set {
					var _coll84 = _dafny.NewBuilder()
					_ = _coll84
					for _iter87 := _dafny.Iterate((_2217_v50).Keys().Elements()); ; {
						_compr_84, _ok87 := _iter87()
						if !_ok87 {
							break
						}
						var _2218_v51 _dafny.Int
						_2218_v51 = interface{}(_compr_84).(_dafny.Int)
						if (_2217_v50).Contains(_2218_v51) {
							_coll84.Add(Companion_Default___.SafeDivisionInt(_2218_v51, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32(((Companion_D2_.Create_DC5_(_dafny.IntOfUint32((_dafny.SeqOf(true, false, true, true)).Cardinality()), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(671), _dafny.IntOfInt64(-9))).Cardinality()), true)).Dtor_cf9()).Cardinality()))).Cardinality())))
						}
					}
					return _coll84.ToSet()
				}()).Intersection(_2204_v37))
				_ = _rhs263
				var _rhs264 _dafny.Map = (_2214_v47).Merge(_2214_v47)
				_ = _rhs264
				var _rhs265 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p0, (func() _dafny.Int {
					if (_2195_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))).Int()).(bool) {
						return _2177_v22
					}
					return (_this).F2()
				})()))
				_ = _rhs265
				var _lhs148 _dafny.Array = _2195_v30
				_ = _lhs148
				var _lhs149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_2195_v30), 0))
				_ = _lhs149
				_2203_v36 = _rhs261
				(_lhs148).ArraySet1(_rhs262, (_lhs149).Int())
				_2204_v37 = _rhs263
				_2214_v47 = _rhs264
				_2177_v22 = _rhs265
			} else {
				var _2219___mcc_h4 D6 = _source23.Get_().(D6_DC16).Cf33
				_ = _2219___mcc_h4
				var _2220_cf33 D6 = _2219___mcc_h4
				_ = _2220_cf33
				var _2221_v52 *C4
				_ = _2221_v52
				var _nw343 *C4 = New_C4_()
				_ = _nw343
				_nw343.Ctor__()
				_2221_v52 = _nw343
				_2149_v0 = _2149_v0
				var _2222_v53 _dafny.Array
				_ = _2222_v53
				var _nw344 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
				_ = _nw344
				_2222_v53 = _nw344
				var _2223_v54 _dafny.Sequence
				_ = _2223_v54
				_2223_v54 = _dafny.SeqOf(true)
				var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_2222_v53), 0))
				_ = _index349
				(_2222_v53).ArraySet1(_2223_v54, (_index349).Int())
				var _2224_v55 _dafny.CodePoint
				_ = _2224_v55
				_2224_v55 = _dafny.CodePoint('g')
				var _2225_v56 _dafny.Map
				_ = _2225_v56
				_2225_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2224_v55, _2223_v54)
				var _2226_v57 D15
				_ = _2226_v57
				_2226_v57 = Companion_D15_.Create_DC38_(_2224_v55)
				var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_2222_v53), 0))
				_ = _index350
				(_2222_v53).ArraySet1((func() _dafny.Sequence {
					if (_2225_v56).Contains((_2226_v57).Dtor_cf71()) {
						return (_2225_v56).Get((_2226_v57).Dtor_cf71()).(_dafny.Sequence)
					}
					return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2223_v54, _2223_v54), (Companion_Default___.SafeIndex(_2177_v22, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2223_v54, _2223_v54)).Cardinality()))).Uint32(), _2149_v0)
				})(), (_index350).Int())
				var _2227_v58 _dafny.Array
				_ = _2227_v58
				var _nw345 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
				_ = _nw345
				_2227_v58 = _nw345
				var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_2227_v58), 0))
				_ = _index351
				(_2227_v58).ArraySet1((_this).F2(), (_index351).Int())
				var _2228_v59 _dafny.MultiSet
				_ = _2228_v59
				_2228_v59 = _dafny.MultiSetOf(_dafny.CodePoint('i'))
				var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_2227_v58), 0))
				_ = _index352
				(_2227_v58).ArraySet1((_2228_v59).Cardinality(), (_index352).Int())
				var _2229_v60 _dafny.Sequence
				_ = _2229_v60
				_2229_v60 = _dafny.SeqOf(_2227_v58, _2227_v58, _2227_v58, _2227_v58)
				var _2230_v61 _dafny.MultiSet
				_ = _2230_v61
				_2230_v61 = _dafny.MultiSetOf(false)
				var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_2227_v58), 0))
				_ = _index353
				var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_2227_v58), 0))
				_ = _index354
				var _rhs266 _dafny.Int = _dafny.IntOfUint32((_2229_v60).Cardinality())
				_ = _rhs266
				var _rhs267 _dafny.Int = (_2230_v61).Cardinality()
				_ = _rhs267
				var _lhs150 _dafny.Array = _2227_v58
				_ = _lhs150
				var _lhs151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_2227_v58), 0))
				_ = _lhs151
				var _lhs152 _dafny.Array = _2227_v58
				_ = _lhs152
				var _lhs153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_2227_v58), 0))
				_ = _lhs153
				(_lhs150).ArraySet1(_rhs266, (_lhs151).Int())
				(_lhs152).ArraySet1(_rhs267, (_lhs153).Int())
			}
			var _pat_let_tv61 = _2177_v22
			_ = _pat_let_tv61
			var _2231_v62 _dafny.Int
			_ = _2231_v62
			var _2232_v63 _dafny.Map
			_ = _2232_v63
			var _2233_v64 _dafny.Int
			_ = _2233_v64
			var _out52 _dafny.Int
			_ = _out52
			var _out53 _dafny.Map
			_ = _out53
			var _out54 _dafny.Int
			_ = _out54
			_out52, _out53, _out54 = (_this).M3((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2178_v23).Cardinality()), _2149_v0)).Cardinality(), func(_pat_let49_0 D0) D0 {
				return func(_2234_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let50_0 _dafny.Int) D0 {
						return func(_2235_dt__update_hcf1_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC0_((_2234_dt__update__tmp_h1).Dtor_cf0(), _2235_dt__update_hcf1_h0)
						}(_pat_let50_0)
					}(_pat_let_tv61)
				}(_pat_let49_0)
			}((_this).F3()), true, globalState)
			_2231_v62 = _out52
			_2232_v63 = _out53
			_2233_v64 = _out54
			var _2236_v65 D7
			_ = _2236_v65
			_2236_v65 = Companion_D7_.Create_DC19_((_this).F2())
			var _source24 D7 = _2236_v65
			_ = _source24
			if _source24.Is_DC18() {
				var _2237_v66 _dafny.Map
				_ = _2237_v66
				_2237_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), _this.F1())
				_2149_v0 = (Companion_Default___.Fm44(!((_this).Fm5((_this).F2(), !(_2149_v0), _2237_v66, _2149_v0, globalState)), globalState)).Cmp(_dafny.IntOfInt64(-859)) == 0
				var _2238_v67 _dafny.Sequence
				_ = _2238_v67
				_2238_v67 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2231_v62, _2149_v0))
				var _2239_v68 _dafny.Sequence
				_ = _2239_v68
				_2239_v68 = _dafny.SeqOf(((_2238_v67).Select((Companion_Default___.SafeIndex(_2231_v62, _dafny.IntOfUint32((_2238_v67).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _dafny.IntOfInt64(-681))
				var _2240_v69 _dafny.Sequence
				_ = _2240_v69
				_2240_v69 = _dafny.SeqOf(_2149_v0)
				var _2241_v70 _dafny.Array
				_ = _2241_v70
				var _nw346 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw346
				_2241_v70 = _nw346
				var _2242_v71 D13
				_ = _2242_v71
				_2242_v71 = Companion_D13_.Create_DC32_(_2178_v23, (_2240_v69).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.IntOfUint32((_2240_v69).Cardinality()))).Uint32()).(bool), _2177_v22, _2241_v70)
				var _rhs268 bool = (_2242_v71).Dtor_cf58()
				_ = _rhs268
				var _rhs269 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2239_v68, _2239_v68), _2239_v68)
				_ = _rhs269
				_2149_v0 = _rhs268
				_2239_v68 = _rhs269
				var _2243_v72 _dafny.Array
				_ = _2243_v72
				var _nw347 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
				_ = _nw347
				_2243_v72 = _nw347
				_2243_v72 = _2243_v72
				_2149_v0 = _2149_v0
			} else if _source24.Is_DC19() {
				var _2244___mcc_h5 _dafny.Int = _source24.Get_().(D7_DC19).Cf35
				_ = _2244___mcc_h5
				var _2245_cf35 _dafny.Int = _2244___mcc_h5
				_ = _2245_cf35
				_2149_v0 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_2178_v23, _2178_v23), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(588))).Uint32(), func(coer124 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg124 _dafny.Int) interface{} {
						return coer124(arg124)
					}
				}(func(_2246_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				})), _2178_v23))
				_2233_v64 = Companion_Default___.Fm44(true, globalState)
				_2149_v0 = _2149_v0
				_2233_v64 = (_this).F2()
			} else {
				var _2247___mcc_h6 _dafny.Sequence = _source24.Get_().(D7_DC17).Cf34
				_ = _2247___mcc_h6
				var _2248_cf34 _dafny.Sequence = _2247___mcc_h6
				_ = _2248_cf34
				var _2249_v73 _dafny.MultiSet
				_ = _2249_v73
				_2249_v73 = _dafny.MultiSetOf(_2231_v62, _2177_v22, (_this).Fm6(Companion_Default___.Fm32(_dafny.IntOfInt64(990), globalState), p0, globalState), p0, _2233_v64)
				_2149_v0 = (_2249_v73).IsSubsetOf((_2249_v73).Difference(_dafny.MultiSetOf((_this).F2())))
				var _2250_v74 _dafny.Array
				_ = _2250_v74
				var _nw348 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
				_ = _nw348
				_2250_v74 = _nw348
				var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_2250_v74), 0))
				_ = _index355
				(_2250_v74).ArraySet1(_dafny.IntOfInt64(341), (_index355).Int())
				var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_2250_v74), 0))
				_ = _index356
				(_2250_v74).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32(((_this).F0()).Cardinality()))).Minus(_dafny.IntOfInt64(-205)), (_index356).Int())
				var _2251_v75 _dafny.CodePoint
				_ = _2251_v75
				_2251_v75 = _dafny.CodePoint('x')
				var _2252_v76 _dafny.MultiSet
				_ = _2252_v76
				_2252_v76 = _dafny.MultiSetOf(_2251_v75, _2251_v75)
				var _2253_v77 _dafny.Int
				_ = _2253_v77
				var _2254_v78 _dafny.Map
				_ = _2254_v78
				var _2255_v79 _dafny.Int
				_ = _2255_v79
				var _out55 _dafny.Int
				_ = _out55
				var _out56 _dafny.Map
				_ = _out56
				var _out57 _dafny.Int
				_ = _out57
				_out55, _out56, _out57 = (_this).M3(_dafny.IntOfInt64(147), (_this).F3(), (_2252_v76).IsSubsetOf(_2252_v76), globalState)
				_2253_v77 = _out55
				_2254_v78 = _out56
				_2255_v79 = _out57
				var _2256_v80 _dafny.Map
				_ = _2256_v80
				_2256_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2149_v0, (_2248_cf34).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2178_v23).Cardinality()), _dafny.IntOfUint32((_2248_cf34).Cardinality()))).Uint32()).(bool))
				_2149_v0 = (func() bool {
					if (_2256_v80).Contains(_2149_v0) {
						return (_2256_v80).Get(_2149_v0).(bool)
					}
					return _2149_v0
				})()
			}
			var _2257_v81 _dafny.Sequence
			_ = _2257_v81
			_2257_v81 = _dafny.SeqOf(_2149_v0, _2149_v0, Companion_Default___.Fm17(globalState))
			var _2258_v82 _dafny.Map
			_ = _2258_v82
			_2258_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm7(_2233_v64, globalState), _2149_v0)
			var _2259_v83 _dafny.Map
			_ = _2259_v83
			_2259_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2177_v22, _2149_v0)
			var _rhs270 _dafny.Sequence = _dafny.SeqOf((func() bool {
				if (_2258_v82).Contains(_2178_v23) {
					return (_2258_v82).Get(_2178_v23).(bool)
				}
				return (func() bool {
					if (_2259_v83).Contains(_2233_v64) {
						return (_2259_v83).Get(_2233_v64).(bool)
					}
					return _2149_v0
				})()
			})())
			_ = _rhs270
			var _rhs271 _dafny.Int = _2233_v64
			_ = _rhs271
			_2257_v81 = _rhs270
			_2231_v62 = _rhs271
		}
		var _2260_v84 _dafny.Array
		_ = _2260_v84
		var _len0_54 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_54
		var _nw349 _dafny.Array
		_ = _nw349
		if _len0_54.Cmp(_dafny.Zero) == 0 {
			_nw349 = _dafny.NewArray(_len0_54)
		} else {
			var _init54 func(_dafny.Int) bool = (func(_2261_v0 bool) func(_dafny.Int) bool {
				return func(_2262_i5 _dafny.Int) bool {
					return _2261_v0
				}
			})(_2149_v0)
			_ = _init54
			var _element0_54 = _init54(_dafny.Zero)
			_ = _element0_54
			_nw349 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
			(_nw349).ArraySet1(_element0_54, 0)
			var _nativeLen0_54 = (_len0_54).Int()
			_ = _nativeLen0_54
			for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
				(_nw349).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
			}
		}
		_2260_v84 = _nw349
		_2260_v84 = (func() _dafny.Array {
			if _2149_v0 {
				return _2260_v84
			}
			return _2260_v84
		})()
		var _2263_v85 _dafny.Int
		_ = _2263_v85
		_2263_v85 = _dafny.IntOfInt64(844)
		_2263_v85 = (_dafny.Zero).Minus(p0)
		var _2264_v86 _dafny.Sequence
		_ = _2264_v86
		_2264_v86 = _dafny.SeqOf(_2149_v0, !(_2149_v0) || (_2149_v0))
		var _2265_i6 _dafny.Int
		_ = _2265_i6
		_2265_i6 = _dafny.Zero
		{
			for (_2264_v86).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if _2149_v0 {
					return _dafny.IntOfUint32((_2264_v86).Cardinality())
				}
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(73))).Uint32(), func(coer125 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg125 _dafny.Int) interface{} {
						return coer125(arg125)
					}
				}(func(_2293_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('q')
				}))).Cardinality())
			})(), _dafny.IntOfUint32((_2264_v86).Cardinality()))).Uint32()).(bool) {
				{
					if (_2265_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_2265_i6 = (_2265_i6).Plus(_dafny.One)
					_2263_v85 = _dafny.IntOfInt64(724)
					if !((Companion_D6_.Create_DC15_(p0, Companion_Default___.Fm20(_2149_v0, globalState), !(_2149_v0))).Dtor_cf32()) {
						var _2266_v87 _dafny.Sequence
						_ = _2266_v87
						_2266_v87 = _dafny.SeqOf((func() _dafny.Array {
							if _2149_v0 {
								return _2260_v84
							}
							return _2260_v84
						})(), _2260_v84, _2260_v84)
						_2266_v87 = _2266_v87
						var _2267_v88 _dafny.Array
						_ = _2267_v88
						var _len0_55 _dafny.Int = _dafny.IntOfInt64(18)
						_ = _len0_55
						var _nw350 _dafny.Array
						_ = _nw350
						if _len0_55.Cmp(_dafny.Zero) == 0 {
							_nw350 = _dafny.NewArray(_len0_55)
						} else {
							var _init55 func(_dafny.Int) _dafny.Sequence = (func(_2268_v86 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_2269_i8 _dafny.Int) _dafny.Sequence {
									return _2268_v86
								}
							})(_2264_v86)
							_ = _init55
							var _element0_55 = _init55(_dafny.Zero)
							_ = _element0_55
							_nw350 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
							(_nw350).ArraySet1(_element0_55, 0)
							var _nativeLen0_55 = (_len0_55).Int()
							_ = _nativeLen0_55
							for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
								(_nw350).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
							}
						}
						_2267_v88 = _nw350
						var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_2267_v88), 0))
						_ = _index357
						(_2267_v88).ArraySet1(_2264_v86, (_index357).Int())
						var _2270_v89 D16
						_ = _2270_v89
						_2270_v89 = Companion_D16_.Create_DC44_((_this).F2())
						var _2271_v90 _dafny.Map
						_ = _2271_v90
						_2271_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2270_v89).Dtor_cf82(), _2264_v86)
						var _2272_v91 _dafny.MultiSet
						_ = _2272_v91
						_2272_v91 = _dafny.MultiSetOf(_2264_v86)
						var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_2267_v88), 0))
						_ = _index358
						(_2267_v88).ArraySet1((func() _dafny.Sequence {
							if (_2271_v90).Contains(((_2272_v91).Update(_dafny.SeqOf(_2149_v0, _2149_v0, false, !(_2149_v0), Companion_Default___.Fm17(globalState)), Companion_Default___.Abs(_2263_v85))).Cardinality()) {
								return (_2271_v90).Get(((_2272_v91).Update(_dafny.SeqOf(_2149_v0, _2149_v0, false, !(_2149_v0), Companion_Default___.Fm17(globalState)), Companion_Default___.Abs(_2263_v85))).Cardinality()).(_dafny.Sequence)
							}
							return _2264_v86
						})(), (_index358).Int())
						_2263_v85 = p0
						var _2273_v92 _dafny.Array
						_ = _2273_v92
						var _len0_56 _dafny.Int = _dafny.IntOfInt64(2)
						_ = _len0_56
						var _nw351 _dafny.Array
						_ = _nw351
						if _len0_56.Cmp(_dafny.Zero) == 0 {
							_nw351 = _dafny.NewArray(_len0_56)
						} else {
							var _init56 func(_dafny.Int) _dafny.Int = (func(_2274_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_2275_i9 _dafny.Int) _dafny.Int {
									return Companion_Default___.SafeDivisionInt(_2275_i9, _2274_p0)
								}
							})(p0)
							_ = _init56
							var _element0_56 = _init56(_dafny.Zero)
							_ = _element0_56
							_nw351 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
							(_nw351).ArraySet1(_element0_56, 0)
							var _nativeLen0_56 = (_len0_56).Int()
							_ = _nativeLen0_56
							for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
								(_nw351).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
							}
						}
						_2273_v92 = _nw351
						var _2276_v93 T1
						_ = _2276_v93
						var _nw352 *C5 = New_C5_()
						_ = _nw352
						_nw352.Ctor__(_2273_v92, (_this).F0(), _this.F1())
						_2276_v93 = _nw352
						var _2277_v94 _dafny.Sequence
						_ = _2277_v94
						_2277_v94 = _dafny.SeqOf(_2276_v93, _2276_v93)
						var _2278_v95 _dafny.Array
						_ = _2278_v95
						var _nwElement0_78 T1 = _2276_v93
						_ = _nwElement0_78
						var _nw353 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_78, nil, _dafny.IntOfInt64(13))
						_ = _nw353
						(_nw353).ArraySet1(_nwElement0_78, 0)
						(_nw353).ArraySet1((func() T1 {
							if _2149_v0 {
								return _2276_v93
							}
							return _2276_v93
						})(), 1)
						(_nw353).ArraySet1(_2276_v93, 2)
						(_nw353).ArraySet1(_2276_v93, 3)
						(_nw353).ArraySet1(_2276_v93, 4)
						(_nw353).ArraySet1(_2276_v93, 5)
						(_nw353).ArraySet1((_2277_v94).Select((Companion_Default___.SafeIndex(_2263_v85, _dafny.IntOfUint32((_2277_v94).Cardinality()))).Uint32()).(T1), 6)
						(_nw353).ArraySet1(_2276_v93, 7)
						(_nw353).ArraySet1(_2276_v93, 8)
						(_nw353).ArraySet1(_2276_v93, 9)
						(_nw353).ArraySet1(_2276_v93, 10)
						(_nw353).ArraySet1(_2276_v93, 11)
						(_nw353).ArraySet1(_2276_v93, 12)
						_2278_v95 = _nw353
						var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_2278_v95), 0))
						_ = _index359
						(_2278_v95).ArraySet1(_2276_v93, (_index359).Int())
						var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_2278_v95), 0))
						_ = _index360
						(_2278_v95).ArraySet1(_2276_v93, (_index360).Int())
						var _2279_v96 _dafny.CodePoint
						_ = _2279_v96
						_2279_v96 = _dafny.CodePoint('r')
						var _2280_v97 _dafny.Map
						_ = _2280_v97
						_2280_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2279_v96, _2263_v85)
						var _2281_v98 _dafny.Sequence
						_ = _2281_v98
						_2281_v98 = _dafny.SeqOf(_2264_v86)
						var _2282_v99 _dafny.Set
						_ = _2282_v99
						_2282_v99 = _dafny.SetOf(Companion_Default___.Fm19(_2149_v0, _2149_v0, _2280_v97, !(_2149_v0), globalState), (_2281_v98).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F2()), _dafny.IntOfUint32((_2281_v98).Cardinality()))).Uint32()).(_dafny.Sequence), Companion_Default___.Fm19(_2149_v0, _2149_v0, _2280_v97, _2149_v0, globalState))
						_2263_v85 = (_2282_v99).Cardinality()
					} else {
						_2263_v85 = p0
						var _2283_v100 _dafny.Array
						_ = _2283_v100
						var _nw354 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
						_ = _nw354
						_2283_v100 = _nw354
						var _2284_v101 *C5
						_ = _2284_v101
						var _nw355 *C5 = New_C5_()
						_ = _nw355
						_nw355.Ctor__(_2283_v100, _dafny.Companion_Sequence_.Concatenate((_this).F0(), (_this).F0()), Companion_D0_.Create_DC1_(_2263_v85))
						_2284_v101 = _nw355
						var _2285_v102 _dafny.Sequence
						_ = _2285_v102
						_2285_v102 = _dafny.UnicodeSeqOfUtf8Bytes("v")
						_2285_v102 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm16(_2285_v102, (_dafny.Zero).Minus(_dafny.IntOfUint32((_2285_v102).Cardinality())), globalState), _2285_v102)
						var _2286_v103 _dafny.Map
						_ = _2286_v103
						_2286_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2149_v0, _2149_v0)
						var _2287_v104 _dafny.CodePoint
						_ = _2287_v104
						_2287_v104 = _dafny.CodePoint('y')
						var _2288_v105 _dafny.Map
						_ = _2288_v105
						_2288_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2284_v101).F7(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2149_v0, Companion_Default___.Fm65((_2286_v103).Cardinality(), _2287_v104, (_this).F2(), globalState)))
						var _2289_v106 _dafny.Map
						_ = _2289_v106
						_2289_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2149_v0, p0)
						var _2290_v107 _dafny.Map
						_ = _2290_v107
						_2290_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _2289_v106)
						_2288_v105 = (_2288_v105).Update((_2284_v101).F7(), _2290_v107)
						var _2291_v108 _dafny.Set
						_ = _2291_v108
						_2291_v108 = _dafny.SetOf((_this).F2())
						var _2292_v109 _dafny.Map
						_ = _2292_v109
						_2292_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2260_v84, _2263_v85)
						var _rhs272 _dafny.Set = (_2291_v108).Union(_2291_v108)
						_ = _rhs272
						var _rhs273 bool = Companion_Default___.Fm17(globalState)
						_ = _rhs273
						var _rhs274 _dafny.Int = _dafny.IntOfInt64(363)
						_ = _rhs274
						var _rhs275 _dafny.Map = (_2292_v109).Merge(_2292_v109)
						_ = _rhs275
						var _rhs276 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(((_this).F2()).Minus((_2284_v101).Fm6((_this).F2(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-447)), globalState))))
						_ = _rhs276
						_2291_v108 = _rhs272
						_2149_v0 = _rhs273
						_2263_v85 = _rhs274
						_2292_v109 = _rhs275
						_2263_v85 = _rhs276
					}
					_2263_v85 = p0
					_2149_v0 = !(!(_2149_v0))
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _2294_v110 _dafny.Array
		_ = _2294_v110
		var _nw356 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
		_ = _nw356
		_2294_v110 = _nw356
		var _2295_v111 _dafny.Map
		_ = _2295_v111
		_2295_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2149_v0, _dafny.MultiSetOf((_this).F2()))
		var _2296_v112 _dafny.Map
		_ = _2296_v112
		_2296_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2294_v110, _2295_v111)
		var _2297_v113 _dafny.MultiSet
		_ = _2297_v113
		_2297_v113 = _dafny.MultiSetOf(_2263_v85)
		r0 = ((func() _dafny.Map {
			if (_2296_v112).Contains(_2294_v110) {
				return (_2296_v112).Get(_2294_v110).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2149_v0, _2297_v113)
		})()).Merge(_2295_v111)
		return r0
	}
}
func (_this *C8) M3(p0 _dafny.Int, p1 D0, p2 bool, globalState *GlobalState) (_dafny.Int, _dafny.Map, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _2298_i0 _dafny.Int
		_ = _2298_i0
		_2298_i0 = _dafny.Zero
		{
			for !(p2) {
				{
					if (_2298_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_2298_i0 = (_2298_i0).Plus(_dafny.One)
					var _2299_v0 _dafny.MultiSet
					_ = _2299_v0
					_2299_v0 = _dafny.MultiSetOf(p2)
					var _2300_v1 _dafny.MultiSet
					_ = _2300_v1
					_2300_v1 = _dafny.MultiSetOf((_this).F2(), _dafny.IntOfInt64(949), (_this).F2(), _dafny.IntOfInt64(303), (_this).F2())
					var _2301_v2 _dafny.Map
					_ = _2301_v2
					_2301_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), (_this).F2())
					var _2302_v3 _dafny.Map
					_ = _2302_v3
					_2302_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_dafny.Zero).Minus((_2301_v2).Cardinality()))
					var _2303_v4 _dafny.CodePoint
					_ = _2303_v4
					_2303_v4 = _dafny.CodePoint('m')
					var _2304_v5 _dafny.Map
					_ = _2304_v5
					_2304_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-939), p0)
					var _2305_v6 _dafny.Map
					_ = _2305_v6
					_2305_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(981), p2)
					var _2306_v7 _dafny.Map
					_ = _2306_v7
					_2306_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.CodePoint('m'))
					var _2307_v8 _dafny.Set
					_ = _2307_v8
					_2307_v8 = _dafny.SetOf(_2306_v7)
					var _2308_v9 _dafny.Array
					_ = _2308_v9
					var _nwElement0_79 bool = false
					_ = _nwElement0_79
					var _nw357 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_79, nil, _dafny.IntOfInt64(2))
					_ = _nw357
					(_nw357).ArraySet1(_nwElement0_79, 0)
					(_nw357).ArraySet1((func() bool {
						if (_2305_v6).Contains(_dafny.IntOfInt64(532)) {
							return (_2305_v6).Get(_dafny.IntOfInt64(532)).(bool)
						}
						return p2
					})(), 1)
					_2308_v9 = _nw357
					var _2309_v10 _dafny.Map
					_ = _2309_v10
					_2309_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2308_v9, p0)
					var _2310_v11 _dafny.Array
					_ = _2310_v11
					var _nwElement0_80 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(810))).Uint32(), func(coer126 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg126 _dafny.Int) interface{} {
							return coer126(arg126)
						}
					}(func(_2311_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('p')
					}))).Cardinality())
					_ = _nwElement0_80
					var _nw358 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_80, nil, _dafny.IntOfInt64(22))
					_ = _nw358
					(_nw358).ArraySet1(_nwElement0_80, 0)
					(_nw358).ArraySet1((_2299_v0).Cardinality(), 1)
					(_nw358).ArraySet1((_this).F2(), 2)
					(_nw358).ArraySet1(((_dafny.Zero).Minus((_this).F2())).Times((_2300_v1).Cardinality()), 3)
					(_nw358).ArraySet1((func() _dafny.Int {
						if p2 {
							return _dafny.IntOfUint32((Companion_Default___.Fm31(_dafny.IntOfInt64(-301), p2, p0, globalState)).Cardinality())
						}
						return p0
					})(), 4)
					(_nw358).ArraySet1(p0, 5)
					(_nw358).ArraySet1((_2301_v2).Cardinality(), 6)
					(_nw358).ArraySet1(((_this).F2()).Plus((_this).F2()), 7)
					(_nw358).ArraySet1(_dafny.IntOfInt64(23), 8)
					(_nw358).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p2, (_this).Fm5((_this).F2(), p2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2303_v4, _this.F1()), p2, globalState))).Cardinality()), 9)
					(_nw358).ArraySet1((_this).F2(), 10)
					(_nw358).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F2(), (_dafny.Zero).Minus((func() _dafny.Int {
						if (_2304_v5).Contains((_dafny.Zero).Minus((_this).F2())) {
							return (_2304_v5).Get((_dafny.Zero).Minus((_this).F2())).(_dafny.Int)
						}
						return (_this).F2()
					})()), (_this).F2(), (_2305_v6).Cardinality())).Cardinality())), 11)
					(_nw358).ArraySet1((_this).F2(), 12)
					(_nw358).ArraySet1((p0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2307_v8).Cardinality(), (_this).F2())).Cardinality()), 13)
					(_nw358).ArraySet1(_dafny.IntOfInt64(502), 14)
					(_nw358).ArraySet1((_this).Fm6((_dafny.Zero).Minus((_this).F2()), (_this).F2(), globalState), 15)
					(_nw358).ArraySet1((_this).F2(), 16)
					(_nw358).ArraySet1(((_2309_v10).Cardinality()).Minus(Companion_Default___.Fm44(p2, globalState)), 17)
					(_nw358).ArraySet1(p0, 18)
					(_nw358).ArraySet1((func() _dafny.Int {
						if (_2301_v2).Contains(!(p2)) {
							return (_2301_v2).Get(!(p2)).(_dafny.Int)
						}
						return _dafny.IntOfInt64(800)
					})(), 19)
					(_nw358).ArraySet1(_dafny.IntOfInt64(4), 20)
					(_nw358).ArraySet1(p0, 21)
					_2310_v11 = _nw358
					var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_2310_v11), 0))
					_ = _index361
					(_2310_v11).ArraySet1(p0, (_index361).Int())
					var _index362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_2310_v11), 0))
					_ = _index362
					(_2310_v11).ArraySet1((p0).Plus(p0), (_index362).Int())
					var _2312_v12 _dafny.Sequence
					_ = _2312_v12
					_2312_v12 = _dafny.SeqOf(p2, p2)
					var _2313_v13 _dafny.Map
					_ = _2313_v13
					_2313_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D7_.Create_DC17_(_2312_v12)).Dtor_cf34(), p2)
					r0 = ((_this).F0()).Select((Companion_Default___.SafeIndex((_2313_v13).Cardinality(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int)
					var _2314_v14 bool
					_ = _2314_v14
					_2314_v14 = true
					var _2315_v15 _dafny.Sequence
					_ = _2315_v15
					_2315_v15 = _dafny.UnicodeSeqOfUtf8Bytes("hivvd")
					_2314_v14 = (_this).Fm5(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2315_v15, _2315_v15)).Cardinality()), (func() bool {
						if true {
							return _2314_v14
						}
						return p2
					})(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2303_v4, _this.F1()), (_dafny.IntOfUint32((_2315_v15).Cardinality())).Cmp((_2310_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_2310_v11), 0))).Int()).(_dafny.Int)) >= 0, globalState)
					var _2316_v16 *C6
					_ = _2316_v16
					var _nw359 *C6 = New_C6_()
					_ = _nw359
					_nw359.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((_this).F0()).Cardinality()), p0))
					_2316_v16 = _nw359
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _2317_i2 _dafny.Int
		_ = _2317_i2
		_2317_i2 = _dafny.Zero
		{
			for (_dafny.IntOfInt64(642)).Cmp(Companion_Default___.SafeModuloInt((_this).F2(), p0)) <= 0 {
				{
					if (_2317_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_2317_i2 = (_2317_i2).Plus(_dafny.One)
					if p2 {
						var _2318_v17 _dafny.Array
						_ = _2318_v17
						var _nw360 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
						_ = _nw360
						_2318_v17 = _nw360
						_2318_v17 = _2318_v17
						var _2319_v18 _dafny.Sequence
						_ = _2319_v18
						_2319_v18 = _dafny.UnicodeSeqOfUtf8Bytes("xkbca")
						var _2320_v19 _dafny.CodePoint
						_ = _2320_v19
						_2320_v19 = _dafny.CodePoint('h')
						var _2321_v20 _dafny.Map
						_ = _2321_v20
						_2321_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_2319_v18, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("snuobpymy"), (Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("snuobpymy")).Cardinality()))).Uint32(), _2320_v19)), (_this).F2())
						_2321_v20 = (_2321_v20).Update(_2319_v18, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(459), (_dafny.Zero).Minus(p0)))
						var _2322_v21 _dafny.Map
						_ = _2322_v21
						_2322_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F2(), (_this).F2())
						var _2323_v22 D13
						_ = _2323_v22
						_2323_v22 = Companion_D13_.Create_DC31_(true, p0, (_this).F2())
						var _2324_v23 _dafny.Map
						_ = _2324_v23
						_2324_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)
						var _2325_v24 _dafny.Array
						_ = _2325_v24
						var _nwElement0_81 _dafny.Int = (_this).F2()
						_ = _nwElement0_81
						var _nw361 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_81, nil, _dafny.IntOfInt64(15))
						_ = _nw361
						(_nw361).ArraySet1(_nwElement0_81, 0)
						(_nw361).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
							if p2 {
								return _dafny.Companion_Sequence_.Update(_2319_v18, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2319_v18).Cardinality()))).Uint32(), _2320_v19)
							}
							return _dafny.UnicodeSeqOfUtf8Bytes("b")
						})()).Cardinality()), 1)
						(_nw361).ArraySet1(_dafny.IntOfUint32((_2319_v18).Cardinality()), 2)
						(_nw361).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.SetOf(p0, p0)).Cardinality(), p0), 3)
						(_nw361).ArraySet1(Companion_Default___.SafeModuloInt((_this).F2(), (Companion_Default___.Fm66((_2322_v21).Cardinality(), p2, globalState)).Dtor_cf78()), 4)
						(_nw361).ArraySet1((func() _dafny.Int {
							if p2 {
								return p0
							}
							return p0
						})(), 5)
						(_nw361).ArraySet1((_2323_v22).Dtor_cf55(), 6)
						(_nw361).ArraySet1((_this).F2(), 7)
						(_nw361).ArraySet1(p0, 8)
						(_nw361).ArraySet1((_2324_v23).Cardinality(), 9)
						(_nw361).ArraySet1(p0, 10)
						(_nw361).ArraySet1(p0, 11)
						(_nw361).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gk")).Cardinality()), 12)
						(_nw361).ArraySet1(((_this).F2()).Minus(Companion_Default___.Fm32(p0, globalState)), 13)
						(_nw361).ArraySet1(p0, 14)
						_2325_v24 = _nw361
						var _2326_v25 _dafny.MultiSet
						_ = _2326_v25
						_2326_v25 = _dafny.MultiSetOf((_this).F2())
						var _index363 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_2325_v24), 0))
						_ = _index363
						(_2325_v24).ArraySet1((_2326_v25).Cardinality(), (_index363).Int())
						var _2327_v26 _dafny.Array
						_ = _2327_v26
						var _len0_57 _dafny.Int = _dafny.IntOfInt64(25)
						_ = _len0_57
						var _nw362 _dafny.Array
						_ = _nw362
						if _len0_57.Cmp(_dafny.Zero) == 0 {
							_nw362 = _dafny.NewArray(_len0_57)
						} else {
							var _init57 func(_dafny.Int) bool = (func(_2328_p2 bool) func(_dafny.Int) bool {
								return func(_2329_i3 _dafny.Int) bool {
									return _2328_p2
								}
							})(p2)
							_ = _init57
							var _element0_57 = _init57(_dafny.Zero)
							_ = _element0_57
							_nw362 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
							(_nw362).ArraySet1(_element0_57, 0)
							var _nativeLen0_57 = (_len0_57).Int()
							_ = _nativeLen0_57
							for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
								(_nw362).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
							}
						}
						_2327_v26 = _nw362
						var _2330_v27 _dafny.Set
						_ = _2330_v27
						_2330_v27 = _dafny.SetOf((func() _dafny.Array {
							if p2 {
								return _2327_v26
							}
							return _2327_v26
						})(), _2327_v26, _2327_v26)
						var _2331_v28 bool
						_ = _2331_v28
						_2331_v28 = false
						var _index364 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_2325_v24), 0))
						_ = _index364
						var _rhs277 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F2(), p0)
						_ = _rhs277
						var _rhs278 _dafny.Int = p0
						_ = _rhs278
						var _rhs279 _dafny.Set = _2330_v27
						_ = _rhs279
						var _rhs280 bool = _2331_v28
						_ = _rhs280
						var _lhs154 _dafny.Array = _2325_v24
						_ = _lhs154
						var _lhs155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_2325_v24), 0))
						_ = _lhs155
						r2 = _rhs277
						(_lhs154).ArraySet1(_rhs278, (_lhs155).Int())
						_2330_v27 = _rhs279
						_2331_v28 = _rhs280
						var _index365 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2327_v26), 0))
						_ = _index365
						(_2327_v26).ArraySet1(!(_2331_v28) || (p2), (_index365).Int())
						var _index366 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2327_v26), 0))
						_ = _index366
						var _index367 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_2325_v24), 0))
						_ = _index367
						var _rhs281 bool = !(_2331_v28)
						_ = _rhs281
						var _rhs282 bool = _2331_v28
						_ = _rhs282
						var _rhs283 _dafny.Int = (_dafny.Zero).Minus((_2325_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_2325_v24), 0))).Int()).(_dafny.Int))
						_ = _rhs283
						var _lhs156 _dafny.Array = _2327_v26
						_ = _lhs156
						var _lhs157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2327_v26), 0))
						_ = _lhs157
						var _lhs158 _dafny.Array = _2325_v24
						_ = _lhs158
						var _lhs159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_2325_v24), 0))
						_ = _lhs159
						_2331_v28 = _rhs281
						(_lhs156).ArraySet1(_rhs282, (_lhs157).Int())
						(_lhs158).ArraySet1(_rhs283, (_lhs159).Int())
						var _index368 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_2325_v24), 0))
						_ = _index368
						(_2325_v24).ArraySet1(((_2325_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_2325_v24), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(141))).Uint32(), func(coer127 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg127 _dafny.Int) interface{} {
								return coer127(arg127)
							}
						}((func(_2332_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_2333_i4 _dafny.Int) _dafny.CodePoint {
								return _2332_v19
							}
						})(_2320_v19)))).Cardinality())), (_index368).Int())
					} else {
						var _2334_v29 _dafny.MultiSet
						_ = _2334_v29
						_2334_v29 = _dafny.MultiSetOf(p2, p2, p2, p2)
						var _2335_v30 _dafny.MultiSet
						_ = _2335_v30
						_2335_v30 = _dafny.MultiSetOf(p0)
						var _2336_v31 _dafny.Set
						_ = _2336_v31
						_2336_v31 = _dafny.SetOf((_this).F2(), (func() _dafny.Int {
							if (_2334_v29).Contains(p2) {
								return (_2334_v29).Multiplicity(p2)
							}
							return (func() _dafny.Int {
								if (_2335_v30).Contains((_this).F2()) {
									return (_2335_v30).Multiplicity((_this).F2())
								}
								return (_this).F2()
							})()
						})())
						var _2337_v32 D9
						_ = _2337_v32
						_2337_v32 = Companion_D9_.Create_DC22_(p0)
						var _2338_v33 _dafny.Map
						_ = _2338_v33
						_2338_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2336_v31, _2337_v32)
						_2338_v33 = _2338_v33
						var _2339_v34 _dafny.Array
						_ = _2339_v34
						var _nw363 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
						_ = _nw363
						_2339_v34 = _nw363
						var _index369 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_2339_v34), 0))
						_ = _index369
						(_2339_v34).ArraySet1(p0, (_index369).Int())
						var _index370 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_2339_v34), 0))
						_ = _index370
						(_2339_v34).ArraySet1((_dafny.Zero).Minus((p0).Minus((_this).F2())), (_index370).Int())
						var _2340_v35 _dafny.Sequence
						_ = _2340_v35
						_2340_v35 = _dafny.UnicodeSeqOfUtf8Bytes("ovve")
						_2340_v35 = _2340_v35
						var _2341_v36 bool
						_ = _2341_v36
						_2341_v36 = false
						_2341_v36 = _2341_v36
						_2341_v36 = (Companion_Default___.SafeDivisionInt(p0, (_2339_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_2339_v34), 0))).Int()).(_dafny.Int))).Cmp(_dafny.IntOfUint32(((func() _dafny.Sequence {
							if _2341_v36 {
								return (_this).F0()
							}
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(388))).Uint32(), func(coer128 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg128 _dafny.Int) interface{} {
									return coer128(arg128)
								}
							}(func(_2342_i5 _dafny.Int) _dafny.Int {
								return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-225), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Cardinality()
							}))
						})()).Cardinality())) != 0
					}
					var _2343_v37 _dafny.CodePoint
					_ = _2343_v37
					_2343_v37 = _dafny.CodePoint('l')
					var _2344_v38 D11
					_ = _2344_v38
					_2344_v38 = Companion_D11_.Create_DC25_(_2343_v37, p0, (p0).Cmp(_dafny.IntOfInt64(800)) != 0)
					var _2345_v39 _dafny.Set
					_ = _2345_v39
					_2345_v39 = _dafny.SetOf((_this).F2())
					var _2346_v40 _dafny.Map
					_ = _2346_v40
					_2346_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F2(), p2)
					var _2347_v41 _dafny.Map
					_ = _2347_v41
					_2347_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2345_v39).Cardinality(), (func() bool {
						if (_2346_v40).Contains(p0) {
							return (_2346_v40).Get(p0).(bool)
						}
						return p2
					})())
					var _2348_v42 D16
					_ = _2348_v42
					_2348_v42 = Companion_D16_.Create_DC44_(p0)
					var _2349_v43 _dafny.Set
					_ = _2349_v43
					_2349_v43 = _dafny.SetOf(((_2346_v40).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(p2)))).Cardinality(), ((_this).F2()).Minus((_this).F2()), (_2348_v42).Dtor_cf82(), p0)
					var _2350_v44 _dafny.Sequence
					_ = _2350_v44
					_2350_v44 = _dafny.UnicodeSeqOfUtf8Bytes("ntyajnp")
					var _2351_v45 bool
					_ = _2351_v45
					_2351_v45 = true
					var _2352_v46 _dafny.Map
					_ = _2352_v46
					_2352_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_2351_v45), _2351_v45)
					var _2353_v47 _dafny.Sequence
					_ = _2353_v47
					_2353_v47 = _dafny.SeqOf(p2, p2, p2, p2, p2)
					var _2354_v48 _dafny.Set
					_ = _2354_v48
					_2354_v48 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _2351_v45), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2353_v47).Select((Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32((_2353_v47).Cardinality()))).Uint32()).(bool), p2))
					var _rhs284 D11 = Companion_Default___.Fm67((_dafny.SetOf(_2352_v46)).Intersection(_2354_v48), globalState)
					_ = _rhs284
					var _rhs285 _dafny.Set = (_dafny.SetOf(p0)).Union(_dafny.SetOf((_dafny.MultiSetOf((_this).F2())).Cardinality()))
					_ = _rhs285
					var _rhs286 _dafny.Int = (_this).F2()
					_ = _rhs286
					var _rhs287 _dafny.Sequence = _2350_v44
					_ = _rhs287
					var _rhs288 bool = p2
					_ = _rhs288
					_2344_v38 = _rhs284
					_2345_v39 = _rhs285
					r0 = _rhs286
					_2350_v44 = _rhs287
					_2351_v45 = _rhs288
					var _2355_v49 _dafny.Set
					_ = _2355_v49
					_2355_v49 = _dafny.SetOf(_2350_v44, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-745))).Uint32(), func(coer129 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg129 _dafny.Int) interface{} {
							return coer129(arg129)
						}
					}((func(_2356_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2357_i6 _dafny.Int) _dafny.CodePoint {
							return _2356_v37
						}
					})(_2343_v37))))
					var _2358_v50 _dafny.Map
					_ = _2358_v50
					_2358_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2353_v47, (_2355_v49).Difference(_2355_v49))
					_2358_v50 = (_2358_v50).Update(_2353_v47, _dafny.SetOf(_2350_v44))
					var _2359_v51 *C2
					_ = _2359_v51
					var _nw364 *C2 = New_C2_()
					_ = _nw364
					_nw364.Ctor__()
					_2359_v51 = _nw364
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _2360_v52 D16
		_ = _2360_v52
		_2360_v52 = Companion_D16_.Create_DC44_((_this).F2())
		var _pat_let_tv62 = p2
		_ = _pat_let_tv62
		var _pat_let_tv63 = p2
		_ = _pat_let_tv63
		var _pat_let_tv64 = p2
		_ = _pat_let_tv64
		var _pat_let_tv65 = p2
		_ = _pat_let_tv65
		var _pat_let_tv66 = p2
		_ = _pat_let_tv66
		var _pat_let_tv67 = p2
		_ = _pat_let_tv67
		var _pat_let_tv68 = p2
		_ = _pat_let_tv68
		var _pat_let_tv69 = p2
		_ = _pat_let_tv69
		var _pat_let_tv70 = p2
		_ = _pat_let_tv70
		r2 = func(_source25 D16) _dafny.Int {
			if _source25.Is_DC42() {
				var _2361___mcc_h0 _dafny.Int = _source25.Get_().(D16_DC42).Cf76
				_ = _2361___mcc_h0
				var _2362___mcc_h1 _dafny.Int = _source25.Get_().(D16_DC42).Cf77
				_ = _2362___mcc_h1
				var _2363___mcc_h2 _dafny.Int = _source25.Get_().(D16_DC42).Cf78
				_ = _2363___mcc_h2
				var _2364_cf78 _dafny.Int = _2363___mcc_h2
				_ = _2364_cf78
				var _2365_cf77 _dafny.Int = _2362___mcc_h1
				_ = _2365_cf77
				var _2366_cf76 _dafny.Int = _2361___mcc_h0
				_ = _2366_cf76
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_pat_let_tv62)).Cardinality()))
			} else if _source25.Is_DC43() {
				var _2367___mcc_h3 _dafny.Sequence = _source25.Get_().(D16_DC43).Cf79
				_ = _2367___mcc_h3
				var _2368___mcc_h4 _dafny.CodePoint = _source25.Get_().(D16_DC43).Cf80
				_ = _2368___mcc_h4
				var _2369___mcc_h5 _dafny.Int = _source25.Get_().(D16_DC43).Cf81
				_ = _2369___mcc_h5
				var _2370_cf81 _dafny.Int = _2369___mcc_h5
				_ = _2370_cf81
				var _2371_cf80 _dafny.CodePoint = _2368___mcc_h4
				_ = _2371_cf80
				var _2372_cf79 _dafny.Sequence = _2367___mcc_h3
				_ = _2372_cf79
				return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _pat_let_tv63), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_pat_let_tv64), _pat_let_tv65), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv66, _pat_let_tv67), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv68, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv69, _pat_let_tv70))).Cardinality()
			} else if _source25.Is_DC44() {
				var _2373___mcc_h6 _dafny.Int = _source25.Get_().(D16_DC44).Cf82
				_ = _2373___mcc_h6
				var _2374_cf82 _dafny.Int = _2373___mcc_h6
				_ = _2374_cf82
				return Companion_Default___.SafeDivisionInt((_this).F2(), (_this).F2())
			} else {
				var _2375___mcc_h7 _dafny.Array = _source25.Get_().(D16_DC41).Cf75
				_ = _2375___mcc_h7
				var _2376_cf75 _dafny.Array = _2375___mcc_h7
				_ = _2376_cf75
				return (_this).F2()
			}
		}(_2360_v52)
		var _2377_v53 _dafny.Array
		_ = _2377_v53
		var _nw365 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
		_ = _nw365
		_2377_v53 = _nw365
		var _index371 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))
		_ = _index371
		(_2377_v53).ArraySet1(p2, (_index371).Int())
		var _2378_v54 bool
		_ = _2378_v54
		_2378_v54 = true
		var _index372 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))
		_ = _index372
		var _rhs289 bool = _2378_v54
		_ = _rhs289
		var _rhs290 bool = !(p2)
		_ = _rhs290
		var _lhs160 _dafny.Array = _2377_v53
		_ = _lhs160
		var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))
		_ = _lhs161
		(_lhs160).ArraySet1(_rhs289, (_lhs161).Int())
		_2378_v54 = _rhs290
		var _hi16 _dafny.Int = (func() _dafny.Int {
			if (_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool) {
				return (_dafny.Zero).Minus(_dafny.IntOfInt64(-773))
			}
			return p0
		})()
		_ = _hi16
		for _2379_i7 := (_dafny.IntOfInt64(945)).Minus(p0); _2379_i7.Cmp(_hi16) < 0; _2379_i7 = _2379_i7.Plus(_dafny.One) {
			_2378_v54 = (func() bool {
				if p2 {
					return (p0).Cmp((_this).F2()) >= 0
				}
				return (_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool)
			})()
			var _2380_v55 D12
			_ = _2380_v55
			_2380_v55 = Companion_D12_.Create_DC29_(true, p0, p2, _2379_i7, (_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool))
			var _index373 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))
			_ = _index373
			(_2377_v53).ArraySet1((func() bool {
				if (func() bool {
					if p2 {
						return (_2380_v55).Dtor_cf50()
					}
					return (_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool)
				})() {
					return p2
				}
				return (_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool)
			})(), (_index373).Int())
			var _2381_v56 _dafny.Array
			_ = _2381_v56
			var _nw366 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
			_ = _nw366
			_2381_v56 = _nw366
			_2381_v56 = _2381_v56
			var _2382_v57 *C2
			_ = _2382_v57
			var _nw367 *C2 = New_C2_()
			_ = _nw367
			_nw367.Ctor__()
			_2382_v57 = _nw367
		}
		var _2383_v58 D12
		_ = _2383_v58
		_2383_v58 = Companion_D12_.Create_DC28_(true, (_this).F2())
		var _2384_v59 _dafny.Map
		_ = _2384_v59
		_2384_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool)), (_2383_v58).Dtor_cf46())
		var _2385_v61 _dafny.Sequence
		_ = _2385_v61
		_2385_v61 = _dafny.SeqOf(_2378_v54)
		var _2386_v62 _dafny.Map
		_ = _2386_v62
		_2386_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2385_v61, p2)
		var _2387_v63 _dafny.Map
		_ = _2387_v63
		_2387_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool), func() _dafny.Map {
			var _coll85 = _dafny.NewMapBuilder()
			_ = _coll85
			for _iter88 := _dafny.Iterate((_2386_v62).Keys().Elements()); ; {
				_compr_85, _ok88 := _iter88()
				if !_ok88 {
					break
				}
				var _2388_v60 _dafny.Sequence
				_2388_v60 = interface{}(_compr_85).(_dafny.Sequence)
				if (_2386_v62).Contains(_2388_v60) {
					_coll85.Add(_2388_v60, (_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool))
				}
			}
			return _coll85.ToMap()
		}())
		var _2389_v64 _dafny.Set
		_ = _2389_v64
		_2389_v64 = _dafny.SetOf(_2378_v54)
		var _2390_v65 _dafny.Map
		_ = _2390_v65
		_2390_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _2391_v66 _dafny.Array
		_ = _2391_v66
		var _len0_58 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_58
		var _nw368 _dafny.Array
		_ = _nw368
		if _len0_58.Cmp(_dafny.Zero) == 0 {
			_nw368 = _dafny.NewArray(_len0_58)
		} else {
			var _init58 func(_dafny.Int) _dafny.Int = (func(_2392_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2393_i8 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_2393_i8, _2392_p0)
				}
			})(p0)
			_ = _init58
			var _element0_58 = _init58(_dafny.Zero)
			_ = _element0_58
			_nw368 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
			(_nw368).ArraySet1(_element0_58, 0)
			var _nativeLen0_58 = (_len0_58).Int()
			_ = _nativeLen0_58
			for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
				(_nw368).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
			}
		}
		_2391_v66 = _nw368
		var _2394_v67 _dafny.CodePoint
		_ = _2394_v67
		_2394_v67 = _dafny.CodePoint('i')
		var _2395_v68 _dafny.Map
		_ = _2395_v68
		_2395_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2391_v66, Companion_Default___.Fm65((_dafny.Zero).Minus((_this).F2()), _2394_v67, p0, globalState))
		var _2396_v69 _dafny.Sequence
		_ = _2396_v69
		_2396_v69 = _dafny.UnicodeSeqOfUtf8Bytes("haw")
		var _2397_v70 *C0
		_ = _2397_v70
		var _nw369 *C0 = New_C0_()
		_ = _nw369
		_nw369.Ctor__(_2396_v69)
		_2397_v70 = _nw369
		var _2398_v71 _dafny.Sequence
		_ = _2398_v71
		_2398_v71 = _dafny.SeqOf(_2397_v70)
		var _2399_v72 _dafny.Array
		_ = _2399_v72
		var _nwElement0_82 _dafny.Int = ((_this).F2()).Minus(p0)
		_ = _nwElement0_82
		var _nw370 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_82, nil, _dafny.IntOfInt64(27))
		_ = _nw370
		(_nw370).ArraySet1(_nwElement0_82, 0)
		(_nw370).ArraySet1((_dafny.IntOfInt64(941)).Times((_this).F2()), 1)
		(_nw370).ArraySet1(p0, 2)
		(_nw370).ArraySet1((_this).F2(), 3)
		(_nw370).ArraySet1(_dafny.IntOfUint32(((_this).F0()).Cardinality()), 4)
		(_nw370).ArraySet1(((_this).F2()).Plus((_this).F2()), 5)
		(_nw370).ArraySet1(p0, 6)
		(_nw370).ArraySet1((_this).F2(), 7)
		(_nw370).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf((func() bool {
			if (_2384_v59).Contains((_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool)) {
				return (_2384_v59).Get((_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool)).(bool)
			}
			return p2
		})())).Cardinality())).Plus(_dafny.IntOfInt64(323)), 8)
		(_nw370).ArraySet1((_this).F2(), 9)
		(_nw370).ArraySet1((((func() _dafny.Map {
			if (_2387_v63).Contains(!(p2)) {
				return (_2387_v63).Get(!(p2)).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2385_v61, _2378_v54)
		})()).Cardinality()).Plus((_this).F2()), 10)
		(_nw370).ArraySet1(((_2389_v64).Difference(_2389_v64)).Cardinality(), 11)
		(_nw370).ArraySet1((_this).F2(), 12)
		(_nw370).ArraySet1(Companion_Default___.SafeDivisionInt((_2390_v65).Cardinality(), ((func() _dafny.Map {
			if (_2395_v68).Contains(_2391_v66) {
				return (_2395_v68).Get(_2391_v66).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F2())
		})()).Cardinality()), 13)
		(_nw370).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F2(), _dafny.IntOfInt64(355)), 14)
		(_nw370).ArraySet1((p0).Times(_dafny.IntOfUint32((_2396_v69).Cardinality())), 15)
		(_nw370).ArraySet1((p0).Minus((_dafny.SetOf((_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool), !((_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool)), p2, p2)).Cardinality()), 16)
		(_nw370).ArraySet1((_dafny.Zero).Minus((_this).F2()), 17)
		(_nw370).ArraySet1(p0, 18)
		(_nw370).ArraySet1((func() _dafny.Int {
			if (_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool) {
				return _dafny.IntOfUint32((_2398_v71).Cardinality())
			}
			return (_this).F2()
		})(), 19)
		(_nw370).ArraySet1((p0).Times(Companion_Default___.Fm32(_dafny.IntOfInt64(502), globalState)), 20)
		(_nw370).ArraySet1((_dafny.Zero).Minus(((_2390_v65).Cardinality()).Plus(p0)), 21)
		(_nw370).ArraySet1(p0, 22)
		(_nw370).ArraySet1((_this).F2(), 23)
		(_nw370).ArraySet1(_dafny.IntOfUint32(((_this).F0()).Cardinality()), 24)
		(_nw370).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf((_this).F2())).Cardinality())).Minus(_dafny.IntOfInt64(-469)), 25)
		(_nw370).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), (_this).F2()), 26)
		_2399_v72 = _nw370
		var _2400_v73 D12
		_ = _2400_v73
		_2400_v73 = Companion_D12_.Create_DC29_(_2378_v54, _dafny.IntOfUint32((_2385_v61).Cardinality()), (_2377_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2377_v53), 0))).Int()).(bool), _dafny.IntOfInt64(196), p2)
		var _index374 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_2399_v72), 0))
		_ = _index374
		(_2399_v72).ArraySet1((_2400_v73).Dtor_cf49(), (_index374).Int())
		var _index375 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_2399_v72), 0))
		_ = _index375
		(_2399_v72).ArraySet1((((_dafny.Zero).Minus(p0)).Plus(_dafny.IntOfInt64(447))).Plus(_dafny.IntOfInt64(383)), (_index375).Int())
		var _2401_v74 _dafny.Map
		_ = _2401_v74
		_2401_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_2399_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_2399_v72), 0))).Int()).(_dafny.Int)), (_this).F2()), _2377_v53)
		r0 = (_2401_v74).Cardinality()
		var _2402_v75 _dafny.Map
		_ = _2402_v75
		_2402_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F0()).Select((Companion_Default___.SafeIndex((_2399_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_2399_v72), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int), _2394_v67)
		r1 = (_2402_v75).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2396_v69).Cardinality()), _2394_v67))
		r2 = Companion_Default___.SafeModuloInt((_2399_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_2399_v72), 0))).Int()).(_dafny.Int), p0)
		return r0, r1, r2
	}
}
func (_this *C8) M4(p0 bool, p1 _dafny.Int, p2 _dafny.MultiSet, p3 _dafny.Int, globalState *GlobalState) {
	{
		var _2403_v0 _dafny.Array
		_ = _2403_v0
		var _len0_59 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_59
		var _nw371 _dafny.Array
		_ = _nw371
		if _len0_59.Cmp(_dafny.Zero) == 0 {
			_nw371 = _dafny.NewArray(_len0_59)
		} else {
			var _init59 func(_dafny.Int) bool = func(_2404_i0 _dafny.Int) bool {
				return false
			}
			_ = _init59
			var _element0_59 = _init59(_dafny.Zero)
			_ = _element0_59
			_nw371 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
			(_nw371).ArraySet1(_element0_59, 0)
			var _nativeLen0_59 = (_len0_59).Int()
			_ = _nativeLen0_59
			for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
				(_nw371).ArraySet1(_init59(_dafny.IntOf(_i0_59)), _i0_59)
			}
		}
		_2403_v0 = _nw371
		var _2405_v1 _dafny.Array
		_ = _2405_v1
		var _nwElement0_83 _dafny.Array = _2403_v0
		_ = _nwElement0_83
		var _nw372 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_83, nil, _dafny.IntOfInt64(7))
		_ = _nw372
		(_nw372).ArraySet1(_nwElement0_83, 0)
		(_nw372).ArraySet1(_2403_v0, 1)
		(_nw372).ArraySet1(_2403_v0, 2)
		(_nw372).ArraySet1(_2403_v0, 3)
		(_nw372).ArraySet1(_2403_v0, 4)
		(_nw372).ArraySet1(_2403_v0, 5)
		(_nw372).ArraySet1(_2403_v0, 6)
		_2405_v1 = _nw372
		var _2406_v2 _dafny.CodePoint
		_ = _2406_v2
		_2406_v2 = _dafny.CodePoint('a')
		var _2407_v3 _dafny.Map
		_ = _2407_v3
		_2407_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2406_v2, _this.F1())
		var _2408_v4 _dafny.Map
		_ = _2408_v4
		_2408_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm5(p3, !(p0), _2407_v3, p0, globalState), _2405_v1)
		_2405_v1 = (func() _dafny.Array {
			if (_2408_v4).Contains((_dafny.IntOfUint32(((_this).F0()).Cardinality())).Cmp(p3) == 0) {
				return (_2408_v4).Get((_dafny.IntOfUint32(((_this).F0()).Cardinality())).Cmp(p3) == 0).(_dafny.Array)
			}
			return _2405_v1
		})()
		var _2409_v5 _dafny.Sequence
		_ = _2409_v5
		_2409_v5 = _dafny.SeqOf(!(p0), p0, p0)
		var _2410_v6 _dafny.Sequence
		_ = _2410_v6
		_2410_v6 = _dafny.SeqOf(_2409_v5)
		if !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2409_v5), _2410_v6), _dafny.SeqOf(p0)) {
			var _2411_v7 _dafny.Set
			_ = _2411_v7
			_2411_v7 = _dafny.SetOf(_2403_v0)
			var _2412_v8 _dafny.Set
			_ = _2412_v8
			_2412_v8 = _2411_v7
			_2412_v8 = _2412_v8
			var _2413_v9 bool
			_ = _2413_v9
			_2413_v9 = false
			_2413_v9 = _2413_v9
			if p0 {
				var _2414_v10 _dafny.Int
				_ = _2414_v10
				_2414_v10 = _dafny.IntOfInt64(249)
				_2414_v10 = p1
				var _2415_v11 _dafny.Map
				_ = _2415_v11
				_2415_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(868), !(p0))
				var _2416_v12 _dafny.Map
				_ = _2416_v12
				_2416_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2403_v0, !(_2413_v9))
				var _2417_v13 _dafny.Array
				_ = _2417_v13
				var _nwElement0_84 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2403_v0, (func() bool {
					if (_2415_v11).Contains(_dafny.IntOfInt64(601)) {
						return (_2415_v11).Get(_dafny.IntOfInt64(601)).(bool)
					}
					return _2413_v9
				})())
				_ = _nwElement0_84
				var _nw373 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_84, nil, _dafny.IntOfInt64(26))
				_ = _nw373
				(_nw373).ArraySet1(_nwElement0_84, 0)
				(_nw373).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2403_v0, p0)).Merge(_2416_v12), 1)
				(_nw373).ArraySet1(_2416_v12, 2)
				(_nw373).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2403_v0, p0)).Merge(_2416_v12), 3)
				(_nw373).ArraySet1((func() _dafny.Map {
					if p0 {
						return _2416_v12
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2403_v0, p0)
				})(), 4)
				(_nw373).ArraySet1(_2416_v12, 5)
				(_nw373).ArraySet1((_2416_v12).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2403_v0, _2413_v9)), 6)
				(_nw373).ArraySet1((_2416_v12).Merge(_2416_v12), 7)
				(_nw373).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2403_v0, p0)).Merge(_2416_v12), 8)
				(_nw373).ArraySet1(_2416_v12, 9)
				(_nw373).ArraySet1((_2416_v12).Merge(_2416_v12), 10)
				(_nw373).ArraySet1(_2416_v12, 11)
				(_nw373).ArraySet1(_2416_v12, 12)
				(_nw373).ArraySet1(_2416_v12, 13)
				(_nw373).ArraySet1((_2416_v12).Merge(_2416_v12), 14)
				(_nw373).ArraySet1(_2416_v12, 15)
				(_nw373).ArraySet1(_2416_v12, 16)
				(_nw373).ArraySet1(_2416_v12, 17)
				(_nw373).ArraySet1(_2416_v12, 18)
				(_nw373).ArraySet1(_2416_v12, 19)
				(_nw373).ArraySet1(_2416_v12, 20)
				(_nw373).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2403_v0, p0)).Merge(_2416_v12), 21)
				(_nw373).ArraySet1(_2416_v12, 22)
				(_nw373).ArraySet1(_2416_v12, 23)
				(_nw373).ArraySet1(_2416_v12, 24)
				(_nw373).ArraySet1(_2416_v12, 25)
				_2417_v13 = _nw373
				var _index376 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_2417_v13), 0))
				_ = _index376
				(_2417_v13).ArraySet1(_2416_v12, (_index376).Int())
				var _2418_v14 _dafny.Map
				_ = _2418_v14
				_2418_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.Zero).Minus((_this).F2()))
				var _2419_v15 _dafny.Array
				_ = _2419_v15
				var _len0_60 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_60
				var _nw374 _dafny.Array
				_ = _nw374
				if _len0_60.Cmp(_dafny.Zero) == 0 {
					_nw374 = _dafny.NewArray(_len0_60)
				} else {
					var _init60 func(_dafny.Int) _dafny.Int = (func(_2420_v10 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2421_i1 _dafny.Int) _dafny.Int {
							return (_2421_i1).Times(_2420_v10)
						}
					})(_2414_v10)
					_ = _init60
					var _element0_60 = _init60(_dafny.Zero)
					_ = _element0_60
					_nw374 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
					(_nw374).ArraySet1(_element0_60, 0)
					var _nativeLen0_60 = (_len0_60).Int()
					_ = _nativeLen0_60
					for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
						(_nw374).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
					}
				}
				_2419_v15 = _nw374
				var _2422_v16 *C5
				_ = _2422_v16
				var _nw375 *C5 = New_C5_()
				_ = _nw375
				_nw375.Ctor__(_2419_v15, _dafny.SeqOf((_this).F2(), p1), _this.F1())
				_2422_v16 = _nw375
				var _2423_v17 _dafny.Array
				_ = _2423_v17
				var _nwElement0_85 *C5 = _2422_v16
				_ = _nwElement0_85
				var _nw376 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_85, nil, _dafny.IntOfInt64(21))
				_ = _nw376
				(_nw376).ArraySet1(_nwElement0_85, 0)
				(_nw376).ArraySet1(_2422_v16, 1)
				(_nw376).ArraySet1(_2422_v16, 2)
				(_nw376).ArraySet1(_2422_v16, 3)
				(_nw376).ArraySet1(_2422_v16, 4)
				(_nw376).ArraySet1(_2422_v16, 5)
				(_nw376).ArraySet1(_2422_v16, 6)
				(_nw376).ArraySet1(_2422_v16, 7)
				(_nw376).ArraySet1(_2422_v16, 8)
				(_nw376).ArraySet1(_2422_v16, 9)
				(_nw376).ArraySet1(_2422_v16, 10)
				(_nw376).ArraySet1(_2422_v16, 11)
				(_nw376).ArraySet1(_2422_v16, 12)
				(_nw376).ArraySet1(_2422_v16, 13)
				(_nw376).ArraySet1(_2422_v16, 14)
				(_nw376).ArraySet1(_2422_v16, 15)
				(_nw376).ArraySet1(_2422_v16, 16)
				(_nw376).ArraySet1(_2422_v16, 17)
				(_nw376).ArraySet1(_2422_v16, 18)
				(_nw376).ArraySet1(_2422_v16, 19)
				(_nw376).ArraySet1(_2422_v16, 20)
				_2423_v17 = _nw376
				var _2424_v18 _dafny.Map
				_ = _2424_v18
				_2424_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2423_v17)
				var _2425_v19 _dafny.Sequence
				_ = _2425_v19
				_2425_v19 = _dafny.SeqOf((func() _dafny.Array {
					if (_2424_v18).Contains(p1) {
						return (_2424_v18).Get(p1).(_dafny.Array)
					}
					return _2423_v17
				})())
				var _2426_v20 D12
				_ = _2426_v20
				_2426_v20 = Companion_D12_.Create_DC29_(_2413_v9, p3, !(p0), _dafny.IntOfUint32((_2425_v19).Cardinality()), Companion_Default___.Fm30(_2414_v10, false, p0, globalState))
				var _index377 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_2417_v13), 0))
				_ = _index377
				var _rhs291 _dafny.Int = (Companion_Default___.SafeModuloInt((_2418_v14).Cardinality(), _dafny.IntOfInt64(832))).Minus((_2426_v20).Dtor_cf49())
				_ = _rhs291
				var _rhs292 bool = p0
				_ = _rhs292
				var _rhs293 _dafny.CodePoint = _2406_v2
				_ = _rhs293
				var _rhs294 bool = p0
				_ = _rhs294
				var _rhs295 _dafny.Map = ((_2416_v12).Merge(_2416_v12)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2403_v0, p0))
				_ = _rhs295
				var _lhs162 _dafny.Array = _2417_v13
				_ = _lhs162
				var _lhs163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_2417_v13), 0))
				_ = _lhs163
				_2414_v10 = _rhs291
				_2413_v9 = _rhs292
				_2406_v2 = _rhs293
				_2413_v9 = _rhs294
				(_lhs162).ArraySet1(_rhs295, (_lhs163).Int())
				var _2427_v21 _dafny.Sequence
				_ = _2427_v21
				_2427_v21 = _dafny.SeqOf(((_this).F2()).Minus(p3), p1, p1)
				_2427_v21 = _dafny.Companion_Sequence_.Concatenate((_this).F0(), _2427_v21)
				_2413_v9 = _2413_v9
				_2414_v10 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2427_v21, _2427_v21), (_this).F0())).Cardinality())
			} else {
				var _2428_v22 _dafny.Sequence
				_ = _2428_v22
				_2428_v22 = _dafny.UnicodeSeqOfUtf8Bytes("tqhfxuca")
				var _2429_v23 _dafny.Array
				_ = _2429_v23
				var _nwElement0_86 _dafny.Int = p1
				_ = _nwElement0_86
				var _nw377 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_86, nil, _dafny.IntOfInt64(4))
				_ = _nw377
				(_nw377).ArraySet1(_nwElement0_86, 0)
				(_nw377).ArraySet1((_this).F2(), 1)
				(_nw377).ArraySet1((_this).F2(), 2)
				(_nw377).ArraySet1(_dafny.IntOfUint32((_2428_v22).Cardinality()), 3)
				_2429_v23 = _nw377
				var _2430_v24 *C5
				_ = _2430_v24
				var _nw378 *C5 = New_C5_()
				_ = _nw378
				_nw378.Ctor__(_2429_v23, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(910))).Uint32(), func(coer130 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg130 _dafny.Int) interface{} {
						return coer130(arg130)
					}
				}((func(_2431_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2432_i2 _dafny.Int) _dafny.Int {
						return _2431_p3
					}
				})(p3))), _this.F1())
				_2430_v24 = _nw378
				var _2433_v25 *C5
				_ = _2433_v25
				_2433_v25 = _2430_v24
				var _2434_v26 _dafny.Set
				_ = _2434_v26
				_2434_v26 = _dafny.SetOf((_2433_v25), _2430_v24)
				var _2435_v27 _dafny.Map
				_ = _2435_v27
				_2435_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_2409_v5).Select((Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32((_2409_v5).Cardinality()))).Uint32()).(bool))
				var _2436_v28 _dafny.Map
				_ = _2436_v28
				_2436_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2434_v26).Union(_2434_v26), ((_2435_v27).Merge(_2435_v27)).Cardinality())
				_2436_v28 = (_2436_v28).Update(_2434_v26, (_this).F2())
				var _rhs296 bool = (func() bool {
					if _2413_v9 {
						return p0
					}
					return _2413_v9
				})()
				_ = _rhs296
				var _rhs297 bool = p0
				_ = _rhs297
				_2413_v9 = _rhs296
				_2413_v9 = _rhs297
				var _2437_v29 _dafny.Map
				_ = _2437_v29
				_2437_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((_this).F0()).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.Companion_Sequence_.Update(_2428_v22, (Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32((_2428_v22).Cardinality()))).Uint32(), _2406_v2))
				var _2438_v30 D14
				_ = _2438_v30
				_2438_v30 = Companion_D14_.Create_DC34_(_2437_v29)
				var _2439_v32 _dafny.Map
				_ = _2439_v32
				_2439_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F2(), p0)
				var _2440_v33 _dafny.Map
				_ = _2440_v33
				_2440_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2406_v2)
				var _2441_v35 _dafny.Set
				_ = _2441_v35
				_2441_v35 = _dafny.SetOf(p3, p1)
				var _pat_let_tv71 = p1
				_ = _pat_let_tv71
				var _pat_let_tv72 = _2428_v22
				_ = _pat_let_tv72
				var _pat_let_tv73 = _2437_v29
				_ = _pat_let_tv73
				var _pat_let_tv74 = _2437_v29
				_ = _pat_let_tv74
				var _2442_v36 _dafny.Array
				_ = _2442_v36
				var _nwElement0_87 D14 = _2438_v30
				_ = _nwElement0_87
				var _nw379 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_87, nil, _dafny.IntOfInt64(23))
				_ = _nw379
				(_nw379).ArraySet1(_nwElement0_87, 0)
				(_nw379).ArraySet1(Companion_D14_.Create_DC34_(func() _dafny.Map {
					var _coll86 = _dafny.NewMapBuilder()
					_ = _coll86
					for _iter89 := _dafny.Iterate((_2439_v32).Keys().Elements()); ; {
						_compr_86, _ok89 := _iter89()
						if !_ok89 {
							break
						}
						var _2443_v31 _dafny.Int
						_2443_v31 = interface{}(_compr_86).(_dafny.Int)
						if (_2439_v32).Contains(_2443_v31) {
							_coll86.Add((_2443_v31).Times(p3), _2428_v22)
						}
					}
					return _coll86.ToMap()
				}()), 1)
				(_nw379).ArraySet1(func(_pat_let51_0 D14) D14 {
					return func(_2444_dt__update__tmp_h0 D14) D14 {
						return func(_pat_let52_0 _dafny.Map) D14 {
							return func(_2445_dt__update_hcf62_h0 _dafny.Map) D14 {
								return Companion_D14_.Create_DC34_(_2445_dt__update_hcf62_h0)
							}(_pat_let52_0)
						}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv71, _pat_let_tv72))
					}(_pat_let51_0)
				}(Companion_D14_.Create_DC34_(_2437_v29)), 2)
				(_nw379).ArraySet1(_2438_v30, 3)
				(_nw379).ArraySet1(_2438_v30, 4)
				(_nw379).ArraySet1(_2438_v30, 5)
				(_nw379).ArraySet1(Companion_Default___.Fm68(p0, _2428_v22, globalState), 6)
				(_nw379).ArraySet1(Companion_Default___.Fm68(!(p0), _2428_v22, globalState), 7)
				(_nw379).ArraySet1(_2438_v30, 8)
				(_nw379).ArraySet1(Companion_D14_.Create_DC34_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_2440_v33).Cardinality()), _2428_v22)), 9)
				(_nw379).ArraySet1(Companion_D14_.Create_DC34_(func() _dafny.Map {
					var _coll87 = _dafny.NewMapBuilder()
					_ = _coll87
					for _iter90 := _dafny.Iterate((_2441_v35).Elements()); ; {
						_compr_87, _ok90 := _iter90()
						if !_ok90 {
							break
						}
						var _2446_v34 _dafny.Int
						_2446_v34 = interface{}(_compr_87).(_dafny.Int)
						if (_2441_v35).Contains(_2446_v34) {
							_coll87.Add(Companion_Default___.SafeModuloInt(_2446_v34, p3), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-583))).Uint32(), func(coer131 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg131 _dafny.Int) interface{} {
									return coer131(arg131)
								}
							}((func(_2447_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_2448_i3 _dafny.Int) _dafny.CodePoint {
									return _2447_v2
								}
							})(_2406_v2))))
						}
					}
					return _coll87.ToMap()
				}()), 10)
				(_nw379).ArraySet1(_2438_v30, 11)
				(_nw379).ArraySet1(_2438_v30, 12)
				(_nw379).ArraySet1(Companion_Default___.Fm68(p0, _2428_v22, globalState), 13)
				(_nw379).ArraySet1(_2438_v30, 14)
				(_nw379).ArraySet1(_2438_v30, 15)
				(_nw379).ArraySet1(_2438_v30, 16)
				(_nw379).ArraySet1(func(_pat_let53_0 D14) D14 {
					return func(_2449_dt__update__tmp_h1 D14) D14 {
						return func(_pat_let54_0 _dafny.Map) D14 {
							return func(_2450_dt__update_hcf62_h1 _dafny.Map) D14 {
								return Companion_D14_.Create_DC34_(_2450_dt__update_hcf62_h1)
							}(_pat_let54_0)
						}(_pat_let_tv73)
					}(_pat_let53_0)
				}(_2438_v30), 17)
				(_nw379).ArraySet1(_2438_v30, 18)
				(_nw379).ArraySet1(_2438_v30, 19)
				(_nw379).ArraySet1(_2438_v30, 20)
				(_nw379).ArraySet1(_2438_v30, 21)
				(_nw379).ArraySet1(func(_pat_let55_0 D14) D14 {
					return func(_2451_dt__update__tmp_h2 D14) D14 {
						return func(_pat_let56_0 _dafny.Map) D14 {
							return func(_2452_dt__update_hcf62_h2 _dafny.Map) D14 {
								return Companion_D14_.Create_DC34_(_2452_dt__update_hcf62_h2)
							}(_pat_let56_0)
						}(_pat_let_tv74)
					}(_pat_let55_0)
				}(_2438_v30), 22)
				_2442_v36 = _nw379
				var _index378 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_2442_v36), 0))
				_ = _index378
				(_2442_v36).ArraySet1(_2438_v30, (_index378).Int())
				var _index379 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_2442_v36), 0))
				_ = _index379
				(_2442_v36).ArraySet1(_2438_v30, (_index379).Int())
				_2439_v32 = (Companion_Default___.Fm21(globalState)).Update(p1, _2413_v9)
				_2413_v9 = p0
			}
			var _2453_v37 *C3
			_ = _2453_v37
			var _nw380 *C3 = New_C3_()
			_ = _nw380
			_nw380.Ctor__()
			_2453_v37 = _nw380
			var _2454_v38 _dafny.Array
			_ = _2454_v38
			var _nw381 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(26))
			_ = _nw381
			_2454_v38 = _nw381
			var _2455_v39 _dafny.Sequence
			_ = _2455_v39
			_2455_v39 = _dafny.SeqOf(_2406_v2)
			var _2456_v40 _dafny.MultiSet
			_ = _2456_v40
			_2456_v40 = _dafny.MultiSetOf(_2455_v39)
			var _2457_v42 D2
			_ = _2457_v42
			_2457_v42 = Companion_D2_.Create_DC3_(func() _dafny.Set {
				var _coll88 = _dafny.NewBuilder()
				_ = _coll88
				for _iter91 := _dafny.Iterate((_2456_v40).Elements()); ; {
					_compr_88, _ok91 := _iter91()
					if !_ok91 {
						break
					}
					var _2458_v41 _dafny.Sequence
					_2458_v41 = interface{}(_compr_88).(_dafny.Sequence)
					if (_2456_v40).Contains(_2458_v41) {
						_coll88.Add(_2458_v41)
					}
				}
				return _coll88.ToSet()
			}())
			var _pat_let_tv75 = _2455_v39
			_ = _pat_let_tv75
			var _index380 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_2454_v38), 0))
			_ = _index380
			(_2454_v38).ArraySet1(func(_pat_let57_0 D2) D2 {
				return func(_2459_dt__update__tmp_h3 D2) D2 {
					return func(_pat_let58_0 _dafny.Set) D2 {
						return func(_2460_dt__update_hcf4_h0 _dafny.Set) D2 {
							return Companion_D2_.Create_DC3_(_2460_dt__update_hcf4_h0)
						}(_pat_let58_0)
					}(_dafny.SetOf(_pat_let_tv75))
				}(_pat_let57_0)
			}(_2457_v42), (_index380).Int())
			var _index381 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_2454_v38), 0))
			_ = _index381
			(_2454_v38).ArraySet1(_2457_v42, (_index381).Int())
		} else {
			var _index382 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))
			_ = _index382
			(_2403_v0).ArraySet1(p0, (_index382).Int())
			var _2461_v43 _dafny.Int
			_ = _2461_v43
			_2461_v43 = _dafny.IntOfInt64(950)
			var _2462_v44 _dafny.Sequence
			_ = _2462_v44
			_2462_v44 = _dafny.UnicodeSeqOfUtf8Bytes("wvwvfkb")
			var _2463_v45 _dafny.Array
			_ = _2463_v45
			var _nw382 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw382
			_2463_v45 = _nw382
			var _index383 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))
			_ = _index383
			(_2463_v45).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F0(), (_this).F2())).Cardinality(), _dafny.IntOfInt64(232)), (_index383).Int())
			var _2464_v46 _dafny.Set
			_ = _2464_v46
			_2464_v46 = _dafny.SetOf(false, p0)
			var _2465_v47 _dafny.Map
			_ = _2465_v47
			_2465_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2464_v46).Cardinality(), _2463_v45)
			var _index384 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))
			_ = _index384
			var _index385 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))
			_ = _index385
			var _rhs298 bool = p0
			_ = _rhs298
			var _rhs299 _dafny.Int = Companion_Default___.Fm20(!(p0), globalState)
			_ = _rhs299
			var _rhs300 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_2462_v44, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_2462_v44).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
				if p0 {
					return _dafny.CodePoint('u')
				}
				return _dafny.CodePoint('q')
			})())
			_ = _rhs300
			var _rhs301 _dafny.Int = (func() _dafny.Int {
				if (_2465_v47).Contains((_this).F2()) {
					return Companion_Default___.SafeDivisionInt(Companion_Default___.Fm44(p0, globalState), _2461_v43)
				}
				return p3
			})()
			_ = _rhs301
			var _lhs164 _dafny.Array = _2403_v0
			_ = _lhs164
			var _lhs165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))
			_ = _lhs165
			var _lhs166 _dafny.Array = _2463_v45
			_ = _lhs166
			var _lhs167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))
			_ = _lhs167
			(_lhs164).ArraySet1(_rhs298, (_lhs165).Int())
			_2461_v43 = _rhs299
			_2462_v44 = _rhs300
			(_lhs166).ArraySet1(_rhs301, (_lhs167).Int())
			if (_2403_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))).Int()).(bool) {
				var _index386 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))
				_ = _index386
				(_2463_v45).ArraySet1((_this).F2(), (_index386).Int())
				var _index387 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))
				_ = _index387
				(_2403_v0).ArraySet1(!((_2403_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))).Int()).(bool)), (_index387).Int())
				var _index388 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))
				_ = _index388
				(_2463_v45).ArraySet1((_dafny.Zero).Minus(((_this).F0()).Select((Companion_Default___.SafeIndex((_this).F2(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int)), (_index388).Int())
				var _2466_v48 T0
				_ = _2466_v48
				var _nw383 *C3 = New_C3_()
				_ = _nw383
				_nw383.Ctor__()
				_2466_v48 = _nw383
				_2466_v48 = _2466_v48
				_2462_v44 = _dafny.Companion_Sequence_.Concatenate(_2462_v44, _dafny.Companion_Sequence_.Concatenate(_2462_v44, _dafny.UnicodeSeqOfUtf8Bytes("oo")))
			} else {
				var _index389 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))
				_ = _index389
				(_2403_v0).ArraySet1(((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_2462_v44).Cardinality()), _2461_v43))).Cmp((func() _dafny.Int {
					if true {
						return (_this).F2()
					}
					return Companion_Default___.Fm20(Companion_Default___.Fm17(globalState), globalState)
				})()) == 0, (_index389).Int())
				var _2467_v49 _dafny.Map
				_ = _2467_v49
				_2467_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2461_v43, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-869))).Uint32(), func(coer132 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg132 _dafny.Int) interface{} {
						return coer132(arg132)
					}
				}(func(_2468_i4 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(538)
				})))
				var _index390 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))
				_ = _index390
				(_2463_v45).ArraySet1((func() _dafny.Int {
					if (_2403_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))).Int()).(bool) {
						return (_this).F2()
					}
					return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_2467_v49).Contains(p1) {
							return (_2467_v49).Get(p1).(_dafny.Sequence)
						}
						return (_this).F0()
					})()).Cardinality()), (_dafny.SetOf((_this).F2())).Cardinality())
				})(), (_index390).Int())
				var _2469_v50 _dafny.Map
				_ = _2469_v50
				_2469_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2409_v5).Cardinality()), _dafny.IntOfInt64(930))
				var _2470_v51 _dafny.MultiSet
				_ = _2470_v51
				_2470_v51 = _dafny.MultiSetOf((_2469_v50).Cardinality())
				var _2471_v52 _dafny.MultiSet
				_ = _2471_v52
				_2471_v52 = _dafny.MultiSetOf((_2463_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
					if (_2469_v50).Contains(_dafny.IntOfInt64(-461)) {
						return (_2469_v50).Get(_dafny.IntOfInt64(-461)).(_dafny.Int)
					}
					return (_this).F2()
				})(), p3, (_this).F2(), ((_2470_v51).Cardinality()).Plus(p3))
				var _index391 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))
				_ = _index391
				(_2463_v45).ArraySet1((func() _dafny.Int {
					if (_2470_v51).Contains((_dafny.IntOfInt64(936)).Times(p3)) {
						return (_2470_v51).Multiplicity((_dafny.IntOfInt64(936)).Times(p3))
					}
					return Companion_Default___.Fm44(p0, globalState)
				})(), (_index391).Int())
				var _index392 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))
				_ = _index392
				(_2463_v45).ArraySet1(((_this).F0()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int), (_index392).Int())
				var _2472_v53 *C4
				_ = _2472_v53
				var _nw384 *C4 = New_C4_()
				_ = _nw384
				_nw384.Ctor__()
				_2472_v53 = _nw384
				var _2473_v54 _dafny.Map
				_ = _2473_v54
				_2473_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2472_v53)
				_2473_v54 = (_2473_v54).Update((_2403_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))).Int()).(bool), _2472_v53)
			}
			var _2474_v55 _dafny.Map
			_ = _2474_v55
			_2474_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _2406_v2)
			_2406_v2 = (func() _dafny.CodePoint {
				if (_2474_v55).Contains(p1) {
					return (_2474_v55).Get(p1).(_dafny.CodePoint)
				}
				return _2406_v2
			})()
			var _2475_v56 _dafny.Map
			_ = _2475_v56
			_2475_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_2463_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))).Int()).(_dafny.Int))
			var _2476_v57 _dafny.Map
			_ = _2476_v57
			_2476_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(303), _2475_v56)
			if (((_2476_v57).Update(p3, _2475_v56)).Cardinality()).Cmp(p3) != 0 {
				var _2477_v58 *C2
				_ = _2477_v58
				var _nw385 *C2 = New_C2_()
				_ = _nw385
				_nw385.Ctor__()
				_2477_v58 = _nw385
				var _2478_v59 _dafny.Set
				_ = _2478_v59
				_2478_v59 = _dafny.SetOf(_2477_v58)
				var _2479_v60 _dafny.Array
				_ = _2479_v60
				var _nw386 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
				_ = _nw386
				_2479_v60 = _nw386
				var _2480_v61 _dafny.Map
				_ = _2480_v61
				_2480_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2478_v59).Union(_2478_v59), _2479_v60)
				_2480_v61 = (_2480_v61).Update(_dafny.SetOf(_2477_v58), _2479_v60)
				_2461_v43 = _2461_v43
				var _2481_v62 _dafny.Sequence
				_ = _2481_v62
				_2481_v62 = _dafny.SeqOf((_this).F2())
				_2481_v62 = (Companion_D2_.Create_DC5_(_2461_v43, _2481_v62, false)).Dtor_cf9()
				var _2482_v63 _dafny.Array
				_ = _2482_v63
				var _len0_61 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_61
				var _nw387 _dafny.Array
				_ = _nw387
				if _len0_61.Cmp(_dafny.Zero) == 0 {
					_nw387 = _dafny.NewArray(_len0_61)
				} else {
					var _init61 func(_dafny.Int) _dafny.CodePoint = (func(_2483_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2484_i5 _dafny.Int) _dafny.CodePoint {
							return _2483_v2
						}
					})(_2406_v2)
					_ = _init61
					var _element0_61 = _init61(_dafny.Zero)
					_ = _element0_61
					_nw387 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
					(_nw387).ArraySet1CodePoint(_element0_61, 0)
					var _nativeLen0_61 = (_len0_61).Int()
					_ = _nativeLen0_61
					for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
						(_nw387).ArraySet1CodePoint(_init61(_dafny.IntOf(_i0_61)), _i0_61)
					}
				}
				_2482_v63 = _nw387
				var _2485_v64 _dafny.Sequence
				_ = _2485_v64
				_2485_v64 = _dafny.SeqOf(_2482_v63)
				_2485_v64 = _2485_v64
			} else {
				var _index393 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))
				_ = _index393
				(_2463_v45).ArraySet1((func() _dafny.Int {
					if p0 {
						return _2461_v43
					}
					return (_2463_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))).Int()).(_dafny.Int)
				})(), (_index393).Int())
				_2462_v44 = _2462_v44
				var _2486_v65 _dafny.Sequence
				_ = _2486_v65
				_2486_v65 = _dafny.SeqOf((_2463_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))).Int()).(_dafny.Int))
				_2486_v65 = _2486_v65
				_2406_v2 = (_2462_v44).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(p3, (_dafny.MultiSetOf(_2462_v44)).Cardinality()), _dafny.IntOfUint32((_2462_v44).Cardinality()))).Uint32()).(_dafny.CodePoint)
				var _2487_v66 _dafny.Array
				_ = _2487_v66
				var _nw388 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
				_ = _nw388
				_2487_v66 = _nw388
				var _2488_v67 _dafny.Array
				_ = _2488_v67
				var _nw389 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
				_ = _nw389
				_2488_v67 = _nw389
				var _index394 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_2487_v66), 0))
				_ = _index394
				(_2487_v66).ArraySet1((func() _dafny.Array {
					if (_2409_v5).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_2409_v5).Cardinality()))).Uint32()).(bool) {
						return _2488_v67
					}
					return _2488_v67
				})(), (_index394).Int())
				var _2489_v68 _dafny.Set
				_ = _2489_v68
				_2489_v68 = _dafny.SetOf(_dafny.IntOfUint32((_2462_v44).Cardinality()))
				var _2490_v69 _dafny.Map
				_ = _2490_v69
				_2490_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_2463_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))).Int()).(_dafny.Int))).Cardinality()), (_2403_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))).Int()).(bool))
				var _2491_v70 _dafny.Map
				_ = _2491_v70
				_2491_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfUint32((_2462_v44).Cardinality()), (_2463_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))).Int()).(_dafny.Int), (_2463_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))).Int()).(_dafny.Int), (_2463_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))).Int()).(_dafny.Int), p1), _2490_v69)
				var _index395 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_2487_v66), 0))
				_ = _index395
				var _index396 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))
				_ = _index396
				var _rhs302 _dafny.Array = _2488_v67
				_ = _rhs302
				var _rhs303 _dafny.Int = (_this).F2()
				_ = _rhs303
				var _rhs304 bool = !((_2491_v70).Merge(_2491_v70)).Contains(_2489_v68)
				_ = _rhs304
				var _lhs168 _dafny.Array = _2487_v66
				_ = _lhs168
				var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_2487_v66), 0))
				_ = _lhs169
				var _lhs170 _dafny.Array = _2403_v0
				_ = _lhs170
				var _lhs171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))
				_ = _lhs171
				(_lhs168).ArraySet1(_rhs302, (_lhs169).Int())
				_2461_v43 = _rhs303
				(_lhs170).ArraySet1(_rhs304, (_lhs171).Int())
			}
			var _2492_v71 D13
			_ = _2492_v71
			_2492_v71 = Companion_D13_.Create_DC32_(_2462_v44, p0, _2461_v43, _2463_v45)
			var _2493_v72 _dafny.MultiSet
			_ = _2493_v72
			_2493_v72 = _dafny.MultiSetOf(_2462_v44, (_2492_v71).Dtor_cf57())
			var _index397 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_2403_v0), 0))
			_ = _index397
			(_2403_v0).ArraySet1((_this).Fm5(p3, p0, _2407_v3, ((_2463_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))).Int()).(_dafny.Int)).Cmp((func() _dafny.Int {
				if (_2493_v72).Contains(_dafny.UnicodeSeqOfUtf8Bytes("cu")) {
					return (_2493_v72).Multiplicity(_dafny.UnicodeSeqOfUtf8Bytes("cu"))
				}
				return (_2463_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_2463_v45), 0))).Int()).(_dafny.Int)
			})()) < 0, globalState), (_index397).Int())
		}
		var _2494_v73 _dafny.Array
		_ = _2494_v73
		var _nw390 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw390
		_2494_v73 = _nw390
		var _index398 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_2494_v73), 0))
		_ = _index398
		(_2494_v73).ArraySet1(p3, (_index398).Int())
		var _index399 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_2494_v73), 0))
		_ = _index399
		(_2494_v73).ArraySet1((p1).Times((_this).F2()), (_index399).Int())
		var _index400 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_2403_v0), 0))
		_ = _index400
		(_2403_v0).ArraySet1(p0, (_index400).Int())
		var _2495_v74 _dafny.Set
		_ = _2495_v74
		_2495_v74 = _dafny.SetOf((_this).F0(), _dafny.SeqOf(p3))
		var _2496_v75 D15
		_ = _2496_v75
		_2496_v75 = Companion_D15_.Create_DC37_(_2495_v74)
		var _2497_v76 _dafny.Map
		_ = _2497_v76
		_2497_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2494_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_2494_v73), 0))).Int()).(_dafny.Int), _2496_v75)
		var _2498_v77 _dafny.Array
		_ = _2498_v77
		var _nw391 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(12))
		_ = _nw391
		_2498_v77 = _nw391
		var _2499_v78 D20
		_ = _2499_v78
		_2499_v78 = Companion_D20_.Create_DC52_(_2498_v77)
		var _2500_v79 _dafny.Array
		_ = _2500_v79
		var _nwElement0_88 _dafny.Array = _2498_v77
		_ = _nwElement0_88
		var _nw392 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_88, nil, _dafny.IntOfInt64(22))
		_ = _nw392
		(_nw392).ArraySet1(_nwElement0_88, 0)
		(_nw392).ArraySet1(_2498_v77, 1)
		(_nw392).ArraySet1(_2498_v77, 2)
		(_nw392).ArraySet1(_2498_v77, 3)
		(_nw392).ArraySet1(_2498_v77, 4)
		(_nw392).ArraySet1(_2498_v77, 5)
		(_nw392).ArraySet1(_2498_v77, 6)
		(_nw392).ArraySet1(_2498_v77, 7)
		(_nw392).ArraySet1(_2498_v77, 8)
		(_nw392).ArraySet1(_2498_v77, 9)
		(_nw392).ArraySet1(_2498_v77, 10)
		(_nw392).ArraySet1((_2499_v78).Dtor_cf93(), 11)
		(_nw392).ArraySet1(_2498_v77, 12)
		(_nw392).ArraySet1(_2498_v77, 13)
		(_nw392).ArraySet1(_2498_v77, 14)
		(_nw392).ArraySet1(_2498_v77, 15)
		(_nw392).ArraySet1(_2498_v77, 16)
		(_nw392).ArraySet1((_2499_v78).Dtor_cf93(), 17)
		(_nw392).ArraySet1(_2498_v77, 18)
		(_nw392).ArraySet1(_2498_v77, 19)
		(_nw392).ArraySet1(_2498_v77, 20)
		(_nw392).ArraySet1(_2498_v77, 21)
		_2500_v79 = _nw392
		var _index401 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_2500_v79), 0))
		_ = _index401
		(_2500_v79).ArraySet1(_2498_v77, (_index401).Int())
		var _2501_v81 _dafny.Map
		_ = _2501_v81
		_2501_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_this).F2(), (_2494_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_2494_v73), 0))).Int()).(_dafny.Int)), Companion_Default___.Fm69(false, globalState))
		var _2502_v82 _dafny.Set
		_ = _2502_v82
		_2502_v82 = _dafny.SetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(p3)))
		var _2503_v83 _dafny.Map
		_ = _2503_v83
		_2503_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F2())
		var _index402 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_2403_v0), 0))
		_ = _index402
		var _index403 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_2500_v79), 0))
		_ = _index403
		var _index404 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_2494_v73), 0))
		_ = _index404
		var _rhs305 bool = p0
		_ = _rhs305
		var _rhs306 _dafny.Map = (func() _dafny.Map {
			var _coll89 = _dafny.NewMapBuilder()
			_ = _coll89
			for _iter92 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(311), _dafny.IntOfInt64(702))); ; {
				_compr_89, _ok92 := _iter92()
				if !_ok92 {
					break
				}
				var _2504_v80 _dafny.Int
				_2504_v80 = interface{}(_compr_89).(_dafny.Int)
				if ((_dafny.IntOfInt64(311)).Cmp(_2504_v80) <= 0) && ((_2504_v80).Cmp(_dafny.IntOfInt64(702)) < 0) {
					_coll89.Add((_2504_v80).Minus(_dafny.IntOfInt64(-843)), _2496_v75)
				}
			}
			return _coll89.ToMap()
		}()).Merge((func() _dafny.Map {
			if (_2501_v81).Contains(_2502_v82) {
				return (_2501_v81).Get(_2502_v82).(_dafny.Map)
			}
			return _2497_v76
		})())
		_ = _rhs306
		var _rhs307 _dafny.Array = _2498_v77
		_ = _rhs307
		var _rhs308 _dafny.Int = ((_this).F0()).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_2503_v83).Contains(false) {
				return (_2503_v83).Get(false).(_dafny.Int)
			}
			return (_this).F2()
		})(), _dafny.IntOfUint32(((_this).F0()).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs308
		var _lhs172 _dafny.Array = _2403_v0
		_ = _lhs172
		var _lhs173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_2403_v0), 0))
		_ = _lhs173
		var _lhs174 _dafny.Array = _2500_v79
		_ = _lhs174
		var _lhs175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_2500_v79), 0))
		_ = _lhs175
		var _lhs176 _dafny.Array = _2494_v73
		_ = _lhs176
		var _lhs177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_2494_v73), 0))
		_ = _lhs177
		(_lhs172).ArraySet1(_rhs305, (_lhs173).Int())
		_2497_v76 = _rhs306
		(_lhs174).ArraySet1(_rhs307, (_lhs175).Int())
		(_lhs176).ArraySet1(_rhs308, (_lhs177).Int())
		_2406_v2 = _dafny.CodePoint('x')
		var _2505_v84 *C4
		_ = _2505_v84
		var _nw393 *C4 = New_C4_()
		_ = _nw393
		_nw393.Ctor__()
		_2505_v84 = _nw393
	}
}
func (_this *C8) F2() _dafny.Int {
	{
		return _this._f2
	}
}
func (_this *C8) F3() D0 {
	{
		return _this._f3
	}
}

// End of class C8

// Definition of class C9
type C9 struct {
	dummy byte
}

func New_C9_() *C9 {
	_this := C9{}

	return &_this
}

type CompanionStruct_C9_ struct {
}

var Companion_C9_ = CompanionStruct_C9_{}

func (_this *C9) Equals(other *C9) bool {
	return _this == other
}

func (_this *C9) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C9)
	return ok && _this.Equals(other)
}

func (*C9) String() string {
	return "_module.C9"
}

func Type_C9_() _dafny.TypeDescriptor {
	return type_C9_{}
}

type type_C9_ struct {
}

func (_this type_C9_) Default() interface{} {
	return (*C9)(nil)
}

func (_this type_C9_) String() string {
	return "main.C9"
}
func (_this *C9) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C9{}

func (_this *C9) Ctor__() {
	{
	}
}
func (_this *C9) Fm0(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(63)
	}
}
func (_this *C9) M0(p0 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var _2506_v0 bool
		_ = _2506_v0
		_2506_v0 = false
		var _2507_v1 _dafny.MultiSet
		_ = _2507_v1
		_2507_v1 = _dafny.MultiSetOf(_2506_v0)
		r0 = (_this).Fm0(_2506_v0, (_2507_v1).Cardinality(), globalState)
		var _2508_v2 _dafny.CodePoint
		_ = _2508_v2
		_2508_v2 = _dafny.CodePoint('q')
		var _2509_v3 _dafny.Sequence
		_ = _2509_v3
		_2509_v3 = _dafny.SeqOf(_dafny.CodePoint('w'), _2508_v2)
		var _2510_v4 _dafny.Sequence
		_ = _2510_v4
		_2510_v4 = _dafny.SeqOf(_2509_v3, _2509_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-249))).Uint32(), func(coer133 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg133 _dafny.Int) interface{} {
				return coer133(arg133)
			}
		}(func(_2511_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(398))).Uint32(), func(coer134 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg134 _dafny.Int) interface{} {
				return coer134(arg134)
			}
		}((func(_2512_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2513_i1 _dafny.Int) _dafny.CodePoint {
				return _2512_v2
			}
		})(_2508_v2))), _2509_v3), _dafny.SeqOf(_2508_v2))
		_2510_v4 = _2510_v4
		var _2514_v5 _dafny.MultiSet
		_ = _2514_v5
		_2514_v5 = _dafny.MultiSetOf(_2508_v2)
		var _2515_v6 _dafny.Array
		_ = _2515_v6
		var _nwElement0_89 _dafny.MultiSet = _2514_v5
		_ = _nwElement0_89
		var _nw394 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_89, nil, _dafny.One)
		_ = _nw394
		(_nw394).ArraySet1(_nwElement0_89, 0)
		_2515_v6 = _nw394
		for _iter93 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2515_v6), 0))); ; {
			_guard_loop_3, _ok93 := _iter93()
			if !_ok93 {
				break
			}
			var _2516_i2 _dafny.Int
			_2516_i2 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_2516_i2).Sign() != -1) && ((_2516_i2).Cmp(_dafny.ArrayLen((_2515_v6), 0)) < 0)) {
				(_2515_v6).ArraySet1(_2514_v5, (_2516_i2).Int())
			}
		}
		var _2517_v7 _dafny.Set
		_ = _2517_v7
		_2517_v7 = _dafny.SetOf(_2506_v0)
		var _2518_i3 _dafny.Int
		_ = _2518_i3
		_2518_i3 = _dafny.Zero
		{
			for (_2517_v7).IsProperSubsetOf(_2517_v7) {
				{
					if (_2518_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_2518_i3 = (_2518_i3).Plus(_dafny.One)
					var _2519_v8 _dafny.Map
					_ = _2519_v8
					_2519_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2517_v7).IsDisjointFrom(_2517_v7), p0)
					var _2520_v9 _dafny.Sequence
					_ = _2520_v9
					_2520_v9 = _dafny.SeqOf(true)
					var _2521_v10 _dafny.Array
					_ = _2521_v10
					var _nwElement0_90 bool = _2506_v0
					_ = _nwElement0_90
					var _nw395 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_90, nil, _dafny.IntOfInt64(13))
					_ = _nw395
					(_nw395).ArraySet1(_nwElement0_90, 0)
					(_nw395).ArraySet1(_2506_v0, 1)
					(_nw395).ArraySet1(false, 2)
					(_nw395).ArraySet1(_2506_v0, 3)
					(_nw395).ArraySet1(_2506_v0, 4)
					(_nw395).ArraySet1(!(!(_2506_v0)), 5)
					(_nw395).ArraySet1(_2506_v0, 6)
					(_nw395).ArraySet1(_2506_v0, 7)
					(_nw395).ArraySet1(_2506_v0, 8)
					(_nw395).ArraySet1(_2506_v0, 9)
					(_nw395).ArraySet1((_2520_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2509_v3).Cardinality()), _dafny.IntOfUint32((_2520_v9).Cardinality()))).Uint32()).(bool), 10)
					(_nw395).ArraySet1(_2506_v0, 11)
					(_nw395).ArraySet1(_2506_v0, 12)
					_2521_v10 = _nw395
					var _2522_v11 _dafny.Set
					_ = _2522_v11
					_2522_v11 = _dafny.SetOf(_2521_v10)
					_2519_v8 = (_2519_v8).Update(_2506_v0, (_dafny.Zero).Minus(((_2522_v11).Difference(_2522_v11)).Cardinality()))
					r1 = (_dafny.Zero).Minus(p0)
					var _2523_v12 _dafny.Sequence
					_ = _2523_v12
					_2523_v12 = _dafny.SeqOf((p0).Plus(p0))
					_2523_v12 = _dafny.Companion_Sequence_.Update(_dafny.SeqOf(p0, _dafny.IntOfUint32((_2509_v3).Cardinality()), p0), (Companion_Default___.SafeIndex(((Companion_Default___.Fm1(p0, _2506_v0, p0, _2506_v0, globalState)).Cardinality()).Minus((_this).Fm0(false, p0, globalState)), _dafny.IntOfUint32((_dafny.SeqOf(p0, _dafny.IntOfUint32((_2509_v3).Cardinality()), p0)).Cardinality()))).Uint32(), p0)
					_2506_v0 = _2506_v0
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
		var _2524_v13 _dafny.Map
		_ = _2524_v13
		_2524_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2506_v0)
		_2524_v13 = _2524_v13
		var _2525_v14 _dafny.Array
		_ = _2525_v14
		var _len0_62 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_62
		var _nw396 _dafny.Array
		_ = _nw396
		if _len0_62.Cmp(_dafny.Zero) == 0 {
			_nw396 = _dafny.NewArray(_len0_62)
		} else {
			var _init62 func(_dafny.Int) _dafny.Set = (func(_2526_v7 _dafny.Set) func(_dafny.Int) _dafny.Set {
				return func(_2527_i4 _dafny.Int) _dafny.Set {
					return _2526_v7
				}
			})(_2517_v7)
			_ = _init62
			var _element0_62 = _init62(_dafny.Zero)
			_ = _element0_62
			_nw396 = _dafny.NewArrayFromExample(_element0_62, nil, _len0_62)
			(_nw396).ArraySet1(_element0_62, 0)
			var _nativeLen0_62 = (_len0_62).Int()
			_ = _nativeLen0_62
			for _i0_62 := 1; _i0_62 < _nativeLen0_62; _i0_62++ {
				(_nw396).ArraySet1(_init62(_dafny.IntOf(_i0_62)), _i0_62)
			}
		}
		_2525_v14 = _nw396
		var _2528_v15 D0
		_ = _2528_v15
		_2528_v15 = Companion_D0_.Create_DC0_((func() _dafny.Array {
			if _2506_v0 {
				return _2525_v14
			}
			return _2525_v14
		})(), p0)
		var _source26 D0 = _2528_v15
		_ = _source26
		if _source26.Is_DC0() {
			var _2529___mcc_h0 _dafny.Array = _source26.Get_().(D0_DC0).Cf0
			_ = _2529___mcc_h0
			var _2530___mcc_h1 _dafny.Int = _source26.Get_().(D0_DC0).Cf1
			_ = _2530___mcc_h1
			var _2531_cf1 _dafny.Int = _2530___mcc_h1
			_ = _2531_cf1
			var _2532_cf0 _dafny.Array = _2529___mcc_h0
			_ = _2532_cf0
			r0 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("moityagb")).Cardinality()), p0))
			var _2533_v16 _dafny.Sequence
			_ = _2533_v16
			_2533_v16 = _dafny.SeqOf(_2506_v0, _2506_v0, true, _2506_v0)
			var _2534_v17 _dafny.Array
			_ = _2534_v17
			var _nwElement0_91 bool = !(_2506_v0) || (!(!(_2506_v0)))
			_ = _nwElement0_91
			var _nw397 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_91, nil, _dafny.IntOfInt64(9))
			_ = _nw397
			(_nw397).ArraySet1(_nwElement0_91, 0)
			(_nw397).ArraySet1((_2533_v16).Select((Companion_Default___.SafeIndex((_this).Fm0(_2506_v0, p0, globalState), _dafny.IntOfUint32((_2533_v16).Cardinality()))).Uint32()).(bool), 1)
			(_nw397).ArraySet1((_2533_v16).Select((Companion_Default___.SafeIndex(_2531_cf1, _dafny.IntOfUint32((_2533_v16).Cardinality()))).Uint32()).(bool), 2)
			(_nw397).ArraySet1(!_dafny.Companion_Sequence_.Equal(_2533_v16, _2533_v16), 3)
			(_nw397).ArraySet1((func() bool {
				if _2506_v0 {
					return _2506_v0
				}
				return _2506_v0
			})(), 4)
			(_nw397).ArraySet1((p0).Cmp((_dafny.Zero).Minus(_2531_cf1)) >= 0, 5)
			(_nw397).ArraySet1((func() bool {
				if _2506_v0 {
					return !(_2506_v0)
				}
				return _2506_v0
			})(), 6)
			(_nw397).ArraySet1(_2506_v0, 7)
			(_nw397).ArraySet1(!(_2506_v0), 8)
			_2534_v17 = _nw397
			var _index405 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_2534_v17), 0))
			_ = _index405
			(_2534_v17).ArraySet1(_2506_v0, (_index405).Int())
			var _2535_v18 _dafny.Map
			_ = _2535_v18
			_2535_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2506_v0) && (_2506_v0), _2506_v0)
			var _index406 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_2534_v17), 0))
			_ = _index406
			(_2534_v17).ArraySet1(!((func() bool {
				if (_2535_v18).Contains(!(_2506_v0)) {
					return (_2535_v18).Get(!(_2506_v0)).(bool)
				}
				return _2506_v0
			})()), (_index406).Int())
			var _2536_v19 _dafny.Sequence
			_ = _2536_v19
			_2536_v19 = _dafny.SeqOf(_2517_v7, _dafny.SetOf(_2506_v0), _2517_v7, _2517_v7, _dafny.SetOf((_2534_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_2534_v17), 0))).Int()).(bool)))
			var _2537_v20 _dafny.Set
			_ = _2537_v20
			_2537_v20 = _dafny.SetOf(p0, _2531_cf1, _dafny.IntOfUint32((_2536_v19).Cardinality()))
			var _index407 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_2534_v17), 0))
			_ = _index407
			var _rhs309 _dafny.Int = (_dafny.IntOfInt64(568)).Plus(p0)
			_ = _rhs309
			var _rhs310 _dafny.Sequence = _2509_v3
			_ = _rhs310
			var _rhs311 bool = !(_2506_v0)
			_ = _rhs311
			var _rhs312 bool = (_2537_v20).IsProperSubsetOf(_2537_v20)
			_ = _rhs312
			var _lhs178 _dafny.Array = _2534_v17
			_ = _lhs178
			var _lhs179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_2534_v17), 0))
			_ = _lhs179
			r0 = _rhs309
			_2509_v3 = _rhs310
			(_lhs178).ArraySet1(_rhs311, (_lhs179).Int())
			_2506_v0 = _rhs312
			var _2538_v21 _dafny.Array
			_ = _2538_v21
			var _nw398 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
			_ = _nw398
			_2538_v21 = _nw398
			var _2539_v22 _dafny.Array
			_ = _2539_v22
			var _nw399 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
			_ = _nw399
			_2539_v22 = _nw399
			var _index408 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_2538_v21), 0))
			_ = _index408
			(_2538_v21).ArraySet1(_2539_v22, (_index408).Int())
			var _index409 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_2538_v21), 0))
			_ = _index409
			var _index410 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_2534_v17), 0))
			_ = _index410
			var _rhs313 _dafny.Array = _2539_v22
			_ = _rhs313
			var _rhs314 bool = (_2533_v16).Select((Companion_Default___.SafeIndex(_2531_cf1, _dafny.IntOfUint32((_2533_v16).Cardinality()))).Uint32()).(bool)
			_ = _rhs314
			var _rhs315 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2510_v4, _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(363))).Uint32(), func(coer135 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg135 _dafny.Int) interface{} {
					return coer135(arg135)
				}
			}((func(_2540_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2541_i5 _dafny.Int) _dafny.CodePoint {
					return _2540_v2
				}
			})(_2508_v2))), _2509_v3))
			_ = _rhs315
			var _lhs180 _dafny.Array = _2538_v21
			_ = _lhs180
			var _lhs181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_2538_v21), 0))
			_ = _lhs181
			var _lhs182 _dafny.Array = _2534_v17
			_ = _lhs182
			var _lhs183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_2534_v17), 0))
			_ = _lhs183
			(_lhs180).ArraySet1(_rhs313, (_lhs181).Int())
			(_lhs182).ArraySet1(_rhs314, (_lhs183).Int())
			_2510_v4 = _rhs315
		} else {
			var _2542___mcc_h2 _dafny.Int = _source26.Get_().(D0_DC1).Cf2
			_ = _2542___mcc_h2
			var _2543_cf2 _dafny.Int = _2542___mcc_h2
			_ = _2543_cf2
			if !(!((func() bool {
				if _2506_v0 {
					return Companion_Default___.Fm2((_2506_v0), _2543_cf2, globalState)
				}
				return _2506_v0
			})())) {
				_2543_cf2 = ((_dafny.Zero).Minus(_2543_cf2)).Plus(p0)
				_2506_v0 = _2506_v0
				var _2544_v23 _dafny.Sequence
				_ = _2544_v23
				_2544_v23 = _dafny.SeqOf(_2506_v0)
				_2506_v0 = Companion_Default___.Fm2((_2544_v23).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2544_v23).Cardinality()))).Uint32()).(bool), p0, globalState)
				_2506_v0 = Companion_Default___.Fm2(_2506_v0, Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(829)), globalState)
				var _2545_v24 _dafny.Map
				_ = _2545_v24
				_2545_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2544_v23, _2508_v2)
				_2545_v24 = Companion_Default___.Fm3(_2508_v2, (_2543_cf2).Times((_dafny.Zero).Minus(_2543_cf2)), _2506_v0, _2506_v0, globalState)
			} else {
				var _2546_v25 _dafny.Array
				_ = _2546_v25
				var _nw400 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw400
				_2546_v25 = _nw400
				var _2547_v26 _dafny.Map
				_ = _2547_v26
				_2547_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2506_v0, _2543_cf2)
				var _index411 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_2546_v25), 0))
				_ = _index411
				(_2546_v25).ArraySet1((func() _dafny.Int {
					if (_2547_v26).Contains(true) {
						return (_2547_v26).Get(true).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_2509_v3).Cardinality())
				})(), (_index411).Int())
				var _index412 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_2546_v25), 0))
				_ = _index412
				(_2546_v25).ArraySet1((_this).Fm0((_2506_v0) && (_2506_v0), ((Companion_Default___.Fm4(_2509_v3, _2508_v2, _2506_v0, globalState)).Cardinality()).Minus(p0), globalState), (_index412).Int())
				_2509_v3 = _dafny.Companion_Sequence_.Concatenate(_2509_v3, _dafny.UnicodeSeqOfUtf8Bytes("rbslwhc"))
				var _2548_v27 _dafny.Map
				_ = _2548_v27
				_2548_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2506_v0, _2508_v2)
				var _2549_v28 bool
				_ = _2549_v28
				_2549_v28 = false
				_2548_v27 = (_2548_v27).Update(_2549_v28, _dafny.CodePoint('o'))
				var _2550_v29 _dafny.Array
				_ = _2550_v29
				var _nwElement0_92 bool = !(_2506_v0) || (_2506_v0)
				_ = _nwElement0_92
				var _nw401 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_92, nil, _dafny.IntOfInt64(5))
				_ = _nw401
				(_nw401).ArraySet1(_nwElement0_92, 0)
				(_nw401).ArraySet1(true, 1)
				(_nw401).ArraySet1(_2506_v0, 2)
				(_nw401).ArraySet1(_2506_v0, 3)
				(_nw401).ArraySet1(true, 4)
				_2550_v29 = _nw401
				var _index413 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_2550_v29), 0))
				_ = _index413
				(_2550_v29).ArraySet1(false, (_index413).Int())
				var _2551_v30 D0
				_ = _2551_v30
				_2551_v30 = Companion_D0_.Create_DC1_((_dafny.Zero).Minus((_2546_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_2546_v25), 0))).Int()).(_dafny.Int)))
				var _2552_v31 _dafny.Map
				_ = _2552_v31
				_2552_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2551_v30, _2517_v7)
				var _index414 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_2550_v29), 0))
				_ = _index414
				(_2550_v29).ArraySet1(_2506_v0, (_index414).Int())
				var _2553_v32 _dafny.Map
				_ = _2553_v32
				_2553_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _2506_v0)
				var _index415 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_2550_v29), 0))
				_ = _index415
				var _index416 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_2550_v29), 0))
				_ = _index416
				var _rhs316 _dafny.Int = Companion_Default___.SafeModuloInt((_2546_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_2546_v25), 0))).Int()).(_dafny.Int), p0)
				_ = _rhs316
				var _rhs317 bool = !(_2506_v0)
				_ = _rhs317
				var _rhs318 _dafny.Int = Companion_Default___.SafeModuloInt(_2543_cf2, (_2546_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_2546_v25), 0))).Int()).(_dafny.Int))
				_ = _rhs318
				var _rhs319 _dafny.Map = (func() _dafny.Map {
					if (_2553_v32).Contains(_2506_v0) {
						return _2552_v31
					}
					return (_2552_v31).Update(_2551_v30, _2517_v7)
				})()
				_ = _rhs319
				var _rhs320 bool = true
				_ = _rhs320
				var _lhs184 _dafny.Array = _2550_v29
				_ = _lhs184
				var _lhs185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_2550_v29), 0))
				_ = _lhs185
				var _lhs186 _dafny.Array = _2550_v29
				_ = _lhs186
				var _lhs187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_2550_v29), 0))
				_ = _lhs187
				r0 = _rhs316
				(_lhs184).ArraySet1(_rhs317, (_lhs185).Int())
				r0 = _rhs318
				_2552_v31 = _rhs319
				(_lhs186).ArraySet1(_rhs320, (_lhs187).Int())
				var _index417 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_2550_v29), 0))
				_ = _index417
				(_2550_v29).ArraySet1(((_2546_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_2546_v25), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(470)) <= 0, (_index417).Int())
			}
			var _2554_v33 *C3
			_ = _2554_v33
			var _nw402 *C3 = New_C3_()
			_ = _nw402
			_nw402.Ctor__()
			_2554_v33 = _nw402
			_2543_cf2 = (p0).Times(p0)
			_2506_v0 = !((_2506_v0) == (!((_dafny.SetOf(true)).IsSubsetOf(_2517_v7))))
		}
		r0 = p0
		var _2555_v34 T0
		_ = _2555_v34
		var _nw403 *C3 = New_C3_()
		_ = _nw403
		_nw403.Ctor__()
		_2555_v34 = _nw403
		var _2556_v35 _dafny.Sequence
		_ = _2556_v35
		_2556_v35 = _dafny.SeqOf(_2506_v0, _2506_v0, _2506_v0, _2506_v0)
		var _2557_v36 _dafny.Map
		_ = _2557_v36
		_2557_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2555_v34, _2556_v35)
		r1 = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_2557_v36).Contains(_2555_v34) {
				return (_2557_v36).Get(_2555_v34).(_dafny.Sequence)
			}
			return _2556_v35
		})()).Cardinality())
		r2 = _dafny.Companion_Sequence_.Concatenate(_2556_v35, _2556_v35)
		return r0, r1, r2
	}
}

// End of class C9
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
