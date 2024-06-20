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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bur"), _dafny.UnicodeSeqOfUtf8Bytes("noxdr"))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeModuloInt((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iu")).Cardinality())).Times(_dafny.IntOfInt64(694)), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('q'))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(90))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-787))))
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if true {
			return _dafny.SeqOf(true)
		}
		return _dafny.SeqOf(false)
	})(), _dafny.SeqOf(false, true))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return (func() _dafny.Sequence {
		if true {
			return _dafny.UnicodeSeqOfUtf8Bytes("kgdqdqqcq")
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("qcmmbgxd")
	})()
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rchljvs"), true)).Merge(func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("qxsrnkgpo"))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Sequence
			_0_v0 = interface{}(_compr_0).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("qxsrnkgpo")), _0_v0) {
				_coll0.Add(_0_v0, false)
			}
		}
		return _coll0.ToMap()
	}())).Merge(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(283))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		})), _dafny.UnicodeSeqOfUtf8Bytes("ukufam"))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v1 _dafny.Sequence
			_2_v1 = interface{}(_compr_1).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(283))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('w')
			})), _dafny.UnicodeSeqOfUtf8Bytes("ukufam")), _2_v1) {
				_coll1.Add(_2_v1, !(!(false)))
			}
		}
		return _coll1.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Map, p1 _dafny.Int, p2 D2, p3 bool, globalState *GlobalState) _dafny.Sequence {
	var _source0 D2 = Companion_D2_.Create_DC3_(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("djvdjhw"), _dafny.UnicodeSeqOfUtf8Bytes("vqfwcrmm")))
	_ = _source0
	if _source0.Is_DC4() {
		var _3___mcc_h0 _dafny.Int = _source0.Get_().(D2_DC4).Cf6
		_ = _3___mcc_h0
		var _4___mcc_h1 _dafny.Int = _source0.Get_().(D2_DC4).Cf7
		_ = _4___mcc_h1
		var _5___mcc_h2 _dafny.Int = _source0.Get_().(D2_DC4).Cf8
		_ = _5___mcc_h2
		var _6_cf8 _dafny.Int = _5___mcc_h2
		_ = _6_cf8
		var _7_cf7 _dafny.Int = _4___mcc_h1
		_ = _7_cf7
		var _8_cf6 _dafny.Int = _3___mcc_h0
		_ = _8_cf6
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_6_cf8)).Cardinality())), _dafny.SeqOf(_6_cf8, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(true)).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cfcbcpfd")).Cardinality())), false)).Cardinality())))
	} else if _source0.Is_DC5() {
		return _dafny.SeqOf((func() _dafny.Map {
			var _coll2 = _dafny.NewMapBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(119), (Companion_D8_.Create_DC16_(true, (Companion_D16_.Create_DC39_(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()), true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.CodePoint('s'))).Dtor_cf62())).Dtor_cf24())).Keys().Elements()); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _9_v0 _dafny.Int
				_9_v0 = interface{}(_compr_2).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(119), (Companion_D8_.Create_DC16_(true, (Companion_D16_.Create_DC39_(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()), true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.CodePoint('s'))).Dtor_cf62())).Dtor_cf24())).Contains(_9_v0) {
					_coll2.Add(Companion_Default___.SafeDivisionInt(_9_v0, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(333))).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfInt64(807)))
				}
			}
			return _coll2.ToMap()
		}()).Cardinality())
	} else {
		var _10___mcc_h3 _dafny.MultiSet = _source0.Get_().(D2_DC3).Cf5
		_ = _10___mcc_h3
		var _11_cf5 _dafny.MultiSet = _10___mcc_h3
		_ = _11_cf5
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(732))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_12_i0 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(952))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_13_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			}))).Cardinality()))
		}))
	}
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfInt64(950))).Intersection(_dafny.SetOf((func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(177))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_14_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality()))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _15_v0 _dafny.Int
			_15_v0 = interface{}(_compr_3).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(177))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_14_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			}))).Cardinality())), _15_v0) {
				_coll3.Add((_15_v0).Plus(_dafny.IntOfInt64(443)))
			}
		}
		return _coll3.ToSet()
	}()).Cardinality()))).Union((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('k'))).Cardinality()))).Union(_dafny.SetOf((_dafny.SetOf(!(false))).Cardinality(), _dafny.IntOfInt64(778))))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(false)
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source1 D22 = Companion_D22_.Create_DC51_(_dafny.SetOf(_dafny.IntOfInt64(103)))
	_ = _source1
	if _source1.Is_DC52() {
		return _dafny.CodePoint('q')
	} else if _source1.Is_DC53() {
		var _16___mcc_h0 _dafny.Int = _source1.Get_().(D22_DC53).Cf87
		_ = _16___mcc_h0
		var _17___mcc_h1 bool = _source1.Get_().(D22_DC53).Cf88
		_ = _17___mcc_h1
		var _18___mcc_h2 _dafny.Sequence = _source1.Get_().(D22_DC53).Cf89
		_ = _18___mcc_h2
		var _19___mcc_h3 _dafny.Int = _source1.Get_().(D22_DC53).Cf90
		_ = _19___mcc_h3
		var _20_cf90 _dafny.Int = _19___mcc_h3
		_ = _20_cf90
		var _21_cf89 _dafny.Sequence = _18___mcc_h2
		_ = _21_cf89
		var _22_cf88 bool = _17___mcc_h1
		_ = _22_cf88
		var _23_cf87 _dafny.Int = _16___mcc_h0
		_ = _23_cf87
		return _dafny.CodePoint('h')
	} else {
		var _24___mcc_h4 _dafny.Set = _source1.Get_().(D22_DC51).Cf86
		_ = _24___mcc_h4
		var _25_cf86 _dafny.Set = _24___mcc_h4
		_ = _25_cf86
		return _dafny.CodePoint('j')
	}
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((Companion_D6_.Create_DC9_(func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(40), _dafny.IntOfInt64(-988))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _26_v0 _dafny.Int
			_26_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(40)).Cmp(_26_v0) <= 0) && ((_26_v0).Cmp(_dafny.IntOfInt64(-988)) < 0) {
				_coll4.Add((_26_v0).Times(_dafny.IntOfInt64(144)), true)
			}
		}
		return _coll4.ToMap()
	}())).Dtor_cf12()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(58), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), false)).Cardinality(), false)))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SeqOf(true, false), _dafny.SeqOf(!(false), true), _dafny.SeqOf(!(!(true)), false, true))
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	if false {
		return _dafny.MultiSetOf(false, false)
	} else {
		return _dafny.MultiSetOf(true)
	}
}
func (_static *CompanionStruct_Default___) Fm26(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.IntOfInt64(373))
}
func (_static *CompanionStruct_Default___) Fm27(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(12))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_27_i0 _dafny.Int) _dafny.Int {
		return Companion_Default___.SafeDivisionInt((_dafny.SetOf(Companion_D13_.Create_DC28_(_dafny.UnicodeSeqOfUtf8Bytes("livdrujf"), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D6_.Create_DC9_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((func() _dafny.Set {
			var _coll5 = _dafny.NewBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(951), _dafny.IntOfInt64(703))); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _28_v0 _dafny.Int
				_28_v0 = interface{}(_compr_5).(_dafny.Int)
				if ((_dafny.IntOfInt64(951)).Cmp(_28_v0) <= 0) && ((_28_v0).Cmp(_dafny.IntOfInt64(703)) < 0) {
					_coll5.Add(Companion_Default___.SafeModuloInt(_28_v0, _dafny.IntOfInt64(631)))
				}
			}
			return _coll5.ToSet()
		}()).Cardinality()), false)))).Cardinality(), true))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-719)))
	}))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((func() _dafny.Set {
		if false {
			return _dafny.SetOf(_dafny.CodePoint('e'), _dafny.CodePoint('o'), _dafny.CodePoint('f'), _dafny.CodePoint('q'))
		}
		return _dafny.SetOf(_dafny.CodePoint('b'), _dafny.CodePoint('p'))
	})(), (func() _dafny.Set {
		var _coll6 = _dafny.NewBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('b'))).Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _29_v0 _dafny.CodePoint
			_29_v0 = interface{}(_compr_6).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('b')), _29_v0) {
				_coll6.Add(_29_v0)
			}
		}
		return _coll6.ToSet()
	}()).Difference(_dafny.SetOf(_dafny.CodePoint('w'), _dafny.CodePoint('f'), _dafny.CodePoint('a'), _dafny.CodePoint('t'), _dafny.CodePoint('o'))), (_dafny.SetOf(_dafny.CodePoint('a'), _dafny.CodePoint('w'), _dafny.CodePoint('f'), _dafny.CodePoint('t'), _dafny.CodePoint('k'))).Difference(_dafny.SetOf(_dafny.CodePoint('e'), _dafny.CodePoint('x'), _dafny.CodePoint('n'))), func() _dafny.Set {
		var _coll7 = _dafny.NewBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), false)).Keys().Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _30_v1 _dafny.CodePoint
			_30_v1 = interface{}(_compr_7).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), false)).Contains(_30_v1) {
				_coll7.Add(_30_v1)
			}
		}
		return _coll7.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(96))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_31_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-482)
		}))).Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _32_v0 _dafny.Int
			_32_v0 = interface{}(_compr_8).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(96))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_31_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-482)
			})), _32_v0) {
				_coll8.Add(Companion_Default___.SafeDivisionInt(_32_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wuy")).Cardinality())), true)
			}
		}
		return _coll8.ToMap()
	}()).Cardinality(), false)).Cardinality(), true)).Cardinality())))).Union(_dafny.SetOf(Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-716)))))).Difference((_dafny.SetOf(Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hucnnesk")).Cardinality()))), Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(759))), Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))), Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(473), _dafny.IntOfInt64(328)), _dafny.SetOf(_dafny.IntOfInt64(464)), _dafny.SetOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("v")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true, false, true, false, true)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tloyglr")).Cardinality()))).Cardinality()))).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _33_v1 _dafny.Set
			_33_v1 = interface{}(_compr_9).(_dafny.Set)
			if (_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(473), _dafny.IntOfInt64(328)), _dafny.SetOf(_dafny.IntOfInt64(464)), _dafny.SetOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("v")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true, false, true, false, true)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tloyglr")).Cardinality()))).Cardinality()))).Contains(_33_v1) {
				_coll9.Add(_33_v1)
			}
		}
		return _coll9.ToSet()
	}()).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfInt64(772))).Cardinality()))), Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(832))))).Difference(_dafny.SetOf(Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality()))))))
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-526), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(144), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(354), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(368)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-154))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_34_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-107))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_35_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(741)
		}))).Cardinality())
	})), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("pbacfmeic"), _dafny.IntOfInt64(500))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-351)), _dafny.IntOfInt64(576), _dafny.IntOfInt64(622)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(961))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_36_i2 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(901)
	})))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(313), _dafny.IntOfInt64(538))).Cardinality()))).Cardinality())))).Merge(func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality())))))).Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _37_v0 _dafny.Int
			_37_v0 = interface{}(_compr_10).(_dafny.Int)
			if (_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality())))))).Contains(_37_v0) {
				_coll10.Add((_37_v0).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()), (Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(841)))).Dtor_cf26())
			}
		}
		return _coll10.ToMap()
	}())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(703), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(669))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, globalState *GlobalState) D6 {
	if !(false) {
		return Companion_D6_.Create_DC10_((Companion_D11_.Create_DC24_(_dafny.IntOfInt64(413), _dafny.SeqOf(true), _dafny.CodePoint('f'))).Dtor_cf35())
	} else {
		return Companion_D6_.Create_DC10_(_dafny.SeqOf(true))
	}
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, p1 D6, globalState *GlobalState) _dafny.Sequence {
	var _source2 D22 = Companion_D22_.Create_DC52_()
	_ = _source2
	if _source2.Is_DC52() {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("re")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), true)).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(674))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_38_i0 _dafny.Int) _dafny.Int {
			return (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dppgedr")).Cardinality()))).Cardinality()
		})))
	} else if _source2.Is_DC53() {
		var _39___mcc_h0 _dafny.Int = _source2.Get_().(D22_DC53).Cf87
		_ = _39___mcc_h0
		var _40___mcc_h1 bool = _source2.Get_().(D22_DC53).Cf88
		_ = _40___mcc_h1
		var _41___mcc_h2 _dafny.Sequence = _source2.Get_().(D22_DC53).Cf89
		_ = _41___mcc_h2
		var _42___mcc_h3 _dafny.Int = _source2.Get_().(D22_DC53).Cf90
		_ = _42___mcc_h3
		var _43_cf90 _dafny.Int = _42___mcc_h3
		_ = _43_cf90
		var _44_cf89 _dafny.Sequence = _41___mcc_h2
		_ = _44_cf89
		var _45_cf88 bool = _40___mcc_h1
		_ = _45_cf88
		var _46_cf87 _dafny.Int = _39___mcc_h0
		_ = _46_cf87
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_46_cf87), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(140))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_47_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(921)
		})))
	} else {
		var _48___mcc_h4 _dafny.Set = _source2.Get_().(D22_DC51).Cf86
		_ = _48___mcc_h4
		var _49_cf86 _dafny.Set = _48___mcc_h4
		_ = _49_cf86
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-42))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_50_i2 _dafny.Int) _dafny.Int {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-538), true)).Cardinality()
		}))
	}
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, p1 _dafny.Int, p2 D7, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(false)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(true)))
}
func (_static *CompanionStruct_Default___) Fm36(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) D7 {
	var _source3 D6 = Companion_D6_.Create_DC9_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(19), !(true)))
	_ = _source3
	if _source3.Is_DC10() {
		var _51___mcc_h0 _dafny.Sequence = _source3.Get_().(D6_DC10).Cf13
		_ = _51___mcc_h0
		var _52_cf13 _dafny.Sequence = _51___mcc_h0
		_ = _52_cf13
		return Companion_D7_.Create_DC13_(_dafny.SetOf(false), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xcxj")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(628))).Cardinality(), _dafny.IntOfInt64(-166))).Cardinality())
	} else if _source3.Is_DC11() {
		var _53___mcc_h1 _dafny.Int = _source3.Get_().(D6_DC11).Cf14
		_ = _53___mcc_h1
		var _54___mcc_h2 _dafny.Int = _source3.Get_().(D6_DC11).Cf15
		_ = _54___mcc_h2
		var _55___mcc_h3 _dafny.Int = _source3.Get_().(D6_DC11).Cf16
		_ = _55___mcc_h3
		var _56_cf16 _dafny.Int = _55___mcc_h3
		_ = _56_cf16
		var _57_cf15 _dafny.Int = _54___mcc_h2
		_ = _57_cf15
		var _58_cf14 _dafny.Int = _53___mcc_h1
		_ = _58_cf14
		return Companion_D7_.Create_DC13_(_dafny.SetOf(true, true), _57_cf15, _56_cf16)
	} else {
		var _59___mcc_h4 _dafny.Map = _source3.Get_().(D6_DC9).Cf12
		_ = _59___mcc_h4
		var _60_cf12 _dafny.Map = _59___mcc_h4
		_ = _60_cf12
		return Companion_D7_.Create_DC13_(_dafny.SetOf(false), _dafny.IntOfInt64(461), _dafny.IntOfInt64(-56))
	}
}
func (_static *CompanionStruct_Default___) Fm37(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(((_dafny.SetOf(_dafny.IntOfInt64(863))).Cardinality()).Cmp((_dafny.SetOf(false, true)).Cardinality()) == 0, _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("dexnes"), _dafny.UnicodeSeqOfUtf8Bytes("sijt")), !((_dafny.IntOfInt64(369)).Cmp((_dafny.SetOf(false)).Cardinality()) != 0), false)
}
func (_static *CompanionStruct_Default___) Fm38(p0 _dafny.Map, p1 bool, p2 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(Companion_Default___.SafeDivisionInt((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(278))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_61_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jrpag")).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(true)).Cardinality(), _dafny.IntOfInt64(566), _dafny.IntOfUint32((_dafny.SeqOf(true, true, true, true)).Cardinality()))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('a')
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Sequence, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(Companion_D11_.Create_DC23_(false, _dafny.IntOfInt64(575), _dafny.IntOfInt64(896)), Companion_D11_.Create_DC23_(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(169))).Cardinality(), _dafny.IntOfInt64(427)))
}
func (_static *CompanionStruct_Default___) Fm41(globalState *GlobalState) _dafny.Set {
	var _source4 D15 = Companion_D15_.Create_DC36_(_dafny.IntOfInt64(4), false, true, true, _dafny.IntOfInt64(388))
	_ = _source4
	if _source4.Is_DC34() {
		return _dafny.SetOf(false, false, !((Companion_D20_.Create_DC45_(true, true)).Dtor_cf76()), !(!(false)), true)
	} else if _source4.Is_DC35() {
		var _62___mcc_h0 bool = _source4.Get_().(D15_DC35).Cf50
		_ = _62___mcc_h0
		var _63___mcc_h1 T1 = _source4.Get_().(D15_DC35).Cf51
		_ = _63___mcc_h1
		var _64___mcc_h2 bool = _source4.Get_().(D15_DC35).Cf52
		_ = _64___mcc_h2
		var _65___mcc_h3 _dafny.Sequence = _source4.Get_().(D15_DC35).Cf53
		_ = _65___mcc_h3
		var _66___mcc_h4 _dafny.Int = _source4.Get_().(D15_DC35).Cf54
		_ = _66___mcc_h4
		var _67_cf54 _dafny.Int = _66___mcc_h4
		_ = _67_cf54
		var _68_cf53 _dafny.Sequence = _65___mcc_h3
		_ = _68_cf53
		var _69_cf52 bool = _64___mcc_h2
		_ = _69_cf52
		var _70_cf51 T1 = _63___mcc_h1
		_ = _70_cf51
		var _71_cf50 bool = _62___mcc_h0
		_ = _71_cf50
		return ((_70_cf51).F12()).Union((_70_cf51).F12())
	} else if _source4.Is_DC36() {
		var _72___mcc_h5 _dafny.Int = _source4.Get_().(D15_DC36).Cf55
		_ = _72___mcc_h5
		var _73___mcc_h6 bool = _source4.Get_().(D15_DC36).Cf56
		_ = _73___mcc_h6
		var _74___mcc_h7 bool = _source4.Get_().(D15_DC36).Cf57
		_ = _74___mcc_h7
		var _75___mcc_h8 bool = _source4.Get_().(D15_DC36).Cf58
		_ = _75___mcc_h8
		var _76___mcc_h9 _dafny.Int = _source4.Get_().(D15_DC36).Cf59
		_ = _76___mcc_h9
		var _77_cf59 _dafny.Int = _76___mcc_h9
		_ = _77_cf59
		var _78_cf58 bool = _75___mcc_h8
		_ = _78_cf58
		var _79_cf57 bool = _74___mcc_h7
		_ = _79_cf57
		var _80_cf56 bool = _73___mcc_h6
		_ = _80_cf56
		var _81_cf55 _dafny.Int = _72___mcc_h5
		_ = _81_cf55
		return (_dafny.SetOf(_78_cf58, _79_cf57)).Difference(_dafny.SetOf(_78_cf58))
	} else if _source4.Is_DC33() {
		var _82___mcc_h10 _dafny.Map = _source4.Get_().(D15_DC33).Cf49
		_ = _82___mcc_h10
		var _83_cf49 _dafny.Map = _82___mcc_h10
		_ = _83_cf49
		return _dafny.SetOf(true)
	} else {
		var _84___mcc_h11 D15 = _source4.Get_().(D15_DC37).Cf60
		_ = _84___mcc_h11
		var _85_cf60 D15 = _84___mcc_h11
		_ = _85_cf60
		return _dafny.SetOf(!(!(true)), !(true))
	}
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.Set, globalState *GlobalState) D9 {
	return Companion_D9_.Create_DC19_()
}
func (_static *CompanionStruct_Default___) Fm43(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("bha"), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(703))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("hpxjw"), _dafny.MultiSetOf(_dafny.IntOfInt64(-382))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(529))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_86_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(697))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_87_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	}))).Cardinality()), (func() _dafny.Set {
		var _coll11 = _dafny.NewBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(347), (func() _dafny.Map {
			var _coll12 = _dafny.NewMapBuilder()
			_ = _coll12
			for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(633), _dafny.IntOfInt64(-258))); ; {
				_compr_12, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _88_v0 _dafny.Int
				_88_v0 = interface{}(_compr_12).(_dafny.Int)
				if ((_dafny.IntOfInt64(633)).Cmp(_88_v0) <= 0) && ((_88_v0).Cmp(_dafny.IntOfInt64(-258)) < 0) {
					_coll12.Add(Companion_Default___.SafeModuloInt(_88_v0, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())
				}
			}
			return _coll12.ToMap()
		}()).Cardinality())).Elements()); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _89_v1 _dafny.Int
			_89_v1 = interface{}(_compr_11).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(347), (func() _dafny.Map {
				var _coll13 = _dafny.NewMapBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(633), _dafny.IntOfInt64(-258))); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _88_v0 _dafny.Int
					_88_v0 = interface{}(_compr_13).(_dafny.Int)
					if ((_dafny.IntOfInt64(633)).Cmp(_88_v0) <= 0) && ((_88_v0).Cmp(_dafny.IntOfInt64(-258)) < 0) {
						_coll13.Add(Companion_Default___.SafeModuloInt(_88_v0, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())
					}
				}
				return _coll13.ToMap()
			}()).Cardinality()), _89_v1) {
				_coll11.Add((_89_v1).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(634), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(993))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg18 _dafny.Int) interface{} {
						return coer18(arg18)
					}
				}(func(_90_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				}))).Cardinality()))).Cardinality(), (func() _dafny.Map {
					var _coll14 = _dafny.NewMapBuilder()
					_ = _coll14
					for _iter14 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("wra"), _dafny.IntOfInt64(574))).Keys().Elements()); ; {
						_compr_14, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _91_v2 _dafny.Sequence
						_91_v2 = interface{}(_compr_14).(_dafny.Sequence)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("wra"), _dafny.IntOfInt64(574))).Contains(_91_v2) {
							_coll14.Add(_91_v2, false)
						}
					}
					return _coll14.ToMap()
				}()).Cardinality())).Cardinality()))
			}
		}
		return _coll11.ToSet()
	}()).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm44(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((func() _dafny.Map {
		var _coll15 = _dafny.NewMapBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate((_dafny.MultiSetOf(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC16_(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('p'))).Cardinality())))).Elements()); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _92_v0 D8
			_92_v0 = interface{}(_compr_15).(D8)
			if (_dafny.MultiSetOf(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC16_(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('p'))).Cardinality())))).Contains(_92_v0) {
				_coll15.Add(_92_v0, (_dafny.Zero).Minus((_dafny.SetOf(!((Companion_D21_.Create_DC48_(false)).Dtor_cf78()), !(false))).Cardinality()))
			}
		}
		return _coll15.ToMap()
	}()).Merge(func() _dafny.Map {
		var _coll16 = _dafny.NewMapBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(727))).Uint32(), func(coer19 func(_dafny.Int) D8) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}(func(_93_i0 _dafny.Int) D8 {
			return Companion_D8_.Create_DC17_(Companion_D8_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
		}))).Elements()); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _94_v1 D8
			_94_v1 = interface{}(_compr_16).(D8)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(727))).Uint32(), func(coer20 func(_dafny.Int) D8) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}(func(_93_i0 _dafny.Int) D8 {
				return Companion_D8_.Create_DC17_(Companion_D8_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
			})), _94_v1) {
				_coll16.Add(_94_v1, (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(true)).Cardinality())))).Cardinality())
			}
		}
		return _coll16.ToMap()
	}()), func() _dafny.Map {
		var _coll17 = _dafny.NewMapBuilder()
		_ = _coll17
		for _iter17 := _dafny.Iterate((_dafny.SeqOf(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC16_(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(456))).Cardinality())))))).Elements()); ; {
			_compr_17, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _95_v2 D8
			_95_v2 = interface{}(_compr_17).(D8)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC16_(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(456))).Cardinality()))))), _95_v2) {
				_coll17.Add(_95_v2, _dafny.IntOfInt64(988))
			}
		}
		return _coll17.ToMap()
	}(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(819))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}(func(_96_i1 _dafny.Int) _dafny.Int {
		return (_dafny.SetOf(!(true), true, true, false)).Cardinality()
	}))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true)))), _dafny.IntOfInt64(-665))))
}
func (_static *CompanionStruct_Default___) Fm45(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.IntOfInt64(-462))).Union((_dafny.MultiSetOf(_dafny.IntOfInt64(-710))).Union(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mxu")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm46(globalState *GlobalState) _dafny.Map {
	return (((Companion_D16_.Create_DC38_(func() _dafny.Map {
		var _coll18 = _dafny.NewMapBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate((func() _dafny.Map {
			var _coll19 = _dafny.NewMapBuilder()
			_ = _coll19
			for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(945), _dafny.IntOfInt64(197))); ; {
				_compr_19, _ok19 := _iter19()
				if !_ok19 {
					break
				}
				var _97_v1 _dafny.Int
				_97_v1 = interface{}(_compr_19).(_dafny.Int)
				if ((_dafny.IntOfInt64(945)).Cmp(_97_v1) <= 0) && ((_97_v1).Cmp(_dafny.IntOfInt64(197)) < 0) {
					_coll19.Add((_97_v1).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())
				}
			}
			return _coll19.ToMap()
		}()).Keys().Elements()); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _98_v0 _dafny.Int
			_98_v0 = interface{}(_compr_18).(_dafny.Int)
			if (func() _dafny.Map {
				var _coll20 = _dafny.NewMapBuilder()
				_ = _coll20
				for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(945), _dafny.IntOfInt64(197))); ; {
					_compr_20, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _97_v1 _dafny.Int
					_97_v1 = interface{}(_compr_20).(_dafny.Int)
					if ((_dafny.IntOfInt64(945)).Cmp(_97_v1) <= 0) && ((_97_v1).Cmp(_dafny.IntOfInt64(197)) < 0) {
						_coll20.Add((_97_v1).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())
					}
				}
				return _coll20.ToMap()
			}()).Contains(_98_v0) {
				_coll18.Add((_98_v0).Minus(_dafny.IntOfInt64(-664)), (_dafny.MultiSetOf(false)).Cardinality())
			}
		}
		return _coll18.ToMap()
	}())).Dtor_cf61()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(173), _dafny.IntOfInt64(261))).Cardinality()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(666))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vvgkqkf")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("w")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm47(p0 _dafny.Set, p1 bool, p2 _dafny.Sequence, p3 D15, globalState *GlobalState) D13 {
	return Companion_D13_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC16_(false, (func() _dafny.Map {
		var _coll21 = _dafny.NewMapBuilder()
		_ = _coll21
		for _iter21 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Keys().Elements()); ; {
			_compr_21, _ok21 := _iter21()
			if !_ok21 {
				break
			}
			var _99_v0 _dafny.Map
			_99_v0 = interface{}(_compr_21).(_dafny.Map)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Contains(_99_v0) {
				_coll21.Add(_99_v0, (_dafny.SetOf(true)).Cardinality())
			}
		}
		return _coll21.ToMap()
	}()).Cardinality())), _dafny.IntOfInt64(777)))
}
func (_static *CompanionStruct_Default___) Fm48(p0 _dafny.Int, globalState *GlobalState) D8 {
	if false {
		return Companion_D8_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true)))
	} else {
		return Companion_D8_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false))
	}
}
func (_static *CompanionStruct_Default___) Fm49(p0 bool, globalState *GlobalState) D0 {
	if (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(564))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg22 _dafny.Int) interface{} {
			return coer22(arg22)
		}
	}(func(_100_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))).Cardinality())).Cmp((_dafny.SetOf(false, false, false, !(true), true)).Cardinality()) > 0 {
		return Companion_D0_.Create_DC1_(true, _dafny.IntOfInt64(817), true)
	} else {
		return Companion_D0_.Create_DC1_(!(true), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), false)
	}
}
func (_static *CompanionStruct_Default___) Fm50(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) D11 {
	return Companion_D11_.Create_DC23_((func() bool {
		if false {
			return true
		}
		return !(false)
	})(), (func() _dafny.Int {
		if true {
			return _dafny.IntOfInt64(-90)
		}
		return _dafny.IntOfInt64(147)
	})(), ((_dafny.SetOf(true)).Difference(_dafny.SetOf(false))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm51(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D14_.Create_DC32_(Companion_D14_.Create_DC31_(true, _dafny.IntOfInt64(177), _dafny.IntOfInt64(233)))), _dafny.SeqOf(Companion_D14_.Create_DC32_(Companion_D14_.Create_DC32_(Companion_D14_.Create_DC32_(Companion_D14_.Create_DC31_(false, _dafny.IntOfInt64(71), _dafny.IntOfInt64(802))))), Companion_D14_.Create_DC32_(Companion_D14_.Create_DC32_(Companion_D14_.Create_DC32_(Companion_D14_.Create_DC31_(false, _dafny.IntOfInt64(-926), _dafny.IntOfInt64(123))))), Companion_D14_.Create_DC32_(Companion_D14_.Create_DC32_(Companion_D14_.Create_DC31_(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality(), (_dafny.Zero).Minus((func() _dafny.Set {
		var _coll22 = _dafny.NewBuilder()
		_ = _coll22
		for _iter22 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('w'), _dafny.CodePoint('f'))).Elements()); ; {
			_compr_22, _ok22 := _iter22()
			if !_ok22 {
				break
			}
			var _101_v0 _dafny.CodePoint
			_101_v0 = interface{}(_compr_22).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('w'), _dafny.CodePoint('f')), _101_v0) {
				_coll22.Add(_101_v0)
			}
		}
		return _coll22.ToSet()
	}()).Cardinality())))))), _dafny.SeqOf(Companion_D14_.Create_DC32_(Companion_D14_.Create_DC30_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.UnicodeSeqOfUtf8Bytes("immtempc")))), Companion_D14_.Create_DC32_(Companion_D14_.Create_DC31_(true, _dafny.IntOfInt64(903), (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-189))).Cardinality())).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm52(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(!(false), _dafny.IntOfInt64(949), false), false)).Merge((func() _dafny.Map {
		var _coll23 = _dafny.NewMapBuilder()
		_ = _coll23
		for _iter23 := _dafny.Iterate((func() _dafny.Map {
			var _coll24 = _dafny.NewMapBuilder()
			_ = _coll24
			for _iter24 := _dafny.Iterate((_dafny.SeqOf(Companion_D0_.Create_DC1_(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nkuwoejgm")).Cardinality()), false), Companion_D0_.Create_DC1_(false, _dafny.IntOfInt64(45), !(false)), Companion_D0_.Create_DC1_(!(false), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(568))).Cardinality()), !(!(false))))).Elements()); ; {
				_compr_24, _ok24 := _iter24()
				if !_ok24 {
					break
				}
				var _102_v1 D0
				_102_v1 = interface{}(_compr_24).(D0)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D0_.Create_DC1_(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nkuwoejgm")).Cardinality()), false), Companion_D0_.Create_DC1_(false, _dafny.IntOfInt64(45), !(false)), Companion_D0_.Create_DC1_(!(false), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(568))).Cardinality()), !(!(false)))), _102_v1) {
					_coll24.Add(_102_v1, _dafny.SetOf((_dafny.MultiSetOf(true)).Cardinality()))
				}
			}
			return _coll24.ToMap()
		}()).Keys().Elements()); ; {
			_compr_23, _ok23 := _iter23()
			if !_ok23 {
				break
			}
			var _103_v0 D0
			_103_v0 = interface{}(_compr_23).(D0)
			if (func() _dafny.Map {
				var _coll25 = _dafny.NewMapBuilder()
				_ = _coll25
				for _iter25 := _dafny.Iterate((_dafny.SeqOf(Companion_D0_.Create_DC1_(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nkuwoejgm")).Cardinality()), false), Companion_D0_.Create_DC1_(false, _dafny.IntOfInt64(45), !(false)), Companion_D0_.Create_DC1_(!(false), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(568))).Cardinality()), !(!(false))))).Elements()); ; {
					_compr_25, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _102_v1 D0
					_102_v1 = interface{}(_compr_25).(D0)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D0_.Create_DC1_(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nkuwoejgm")).Cardinality()), false), Companion_D0_.Create_DC1_(false, _dafny.IntOfInt64(45), !(false)), Companion_D0_.Create_DC1_(!(false), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(568))).Cardinality()), !(!(false)))), _102_v1) {
						_coll25.Add(_102_v1, _dafny.SetOf((_dafny.MultiSetOf(true)).Cardinality()))
					}
				}
				return _coll25.ToMap()
			}()).Contains(_103_v0) {
				_coll23.Add(_103_v0, true)
			}
		}
		return _coll23.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(true, (_dafny.Zero).Minus(_dafny.IntOfInt64(-656)), true), false)))
}
func (_static *CompanionStruct_Default___) Fm53(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.SetOf(false), _dafny.SetOf(!(!(true)), !(!(false))))
}
func (_static *CompanionStruct_Default___) Fm54(p0 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('l'), (func() _dafny.Map {
		var _coll26 = _dafny.NewMapBuilder()
		_ = _coll26
		for _iter26 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), _dafny.CodePoint('h')), Companion_D6_.Create_DC9_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}(func(_104_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality()), true)))).Keys().Elements()); ; {
			_compr_26, _ok26 := _iter26()
			if !_ok26 {
				break
			}
			var _105_v0 _dafny.Map
			_105_v0 = interface{}(_compr_26).(_dafny.Map)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), _dafny.CodePoint('h')), Companion_D6_.Create_DC9_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}(func(_104_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			}))).Cardinality()), true)))).Contains(_105_v0) {
				_coll26.Add(_105_v0, false)
			}
		}
		return _coll26.ToMap()
	}()).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm55(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(!(true), false)).Cardinality(), _dafny.IntOfInt64(294))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))).Merge(func() _dafny.Map {
		var _coll27 = _dafny.NewMapBuilder()
		_ = _coll27
		for _iter27 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(464))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_106_i0 _dafny.Int) _dafny.Set {
			return _dafny.SetOf((_dafny.SetOf(false)).Cardinality(), _dafny.IntOfInt64(-889))
		}))).Elements()); ; {
			_compr_27, _ok27 := _iter27()
			if !_ok27 {
				break
			}
			var _107_v0 _dafny.Set
			_107_v0 = interface{}(_compr_27).(_dafny.Set)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(464))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}(func(_106_i0 _dafny.Int) _dafny.Set {
				return _dafny.SetOf((_dafny.SetOf(false)).Cardinality(), _dafny.IntOfInt64(-889))
			})), _107_v0) {
				_coll27.Add(_107_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))
			}
		}
		return _coll27.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm56(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(501))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg27 _dafny.Int) interface{} {
			return coer27(arg27)
		}
	}(func(_108_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(_dafny.IntOfInt64(125), (_dafny.MultiSetOf(_dafny.CodePoint('h'), _dafny.CodePoint('p'))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('s'))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (Companion_D15_.Create_DC36_((_dafny.SetOf(_dafny.IntOfInt64(878))).Cardinality(), true, true, true, _dafny.IntOfInt64(714))).Dtor_cf57()))).Cardinality())), _dafny.IntOfInt64(750))
	})), _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(-861)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dgbtjar")).Cardinality()), _dafny.IntOfInt64(24)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(97))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg28 _dafny.Int) interface{} {
			return coer28(arg28)
		}
	}(func(_109_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(63)
	}))))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _110_v0 bool
	_ = _110_v0
	_110_v0 = false
	var _111_v1 _dafny.Int
	_ = _111_v1
	_111_v1 = _dafny.IntOfInt64(754)
	var _112_v2 _dafny.Sequence
	_ = _112_v2
	_112_v2 = _dafny.SeqOf(true, _110_v0, _110_v0)
	var _113_v3 _dafny.Map
	_ = _113_v3
	_113_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v1, _dafny.IntOfUint32((_112_v2).Cardinality()))
	var _114_v4 _dafny.Map
	_ = _114_v4
	_114_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v0, (func() _dafny.Int {
		if (_113_v3).Contains(_111_v1) {
			return (_113_v3).Get(_111_v1).(_dafny.Int)
		}
		return _111_v1
	})())
	var _115_v5 _dafny.Array
	_ = _115_v5
	var _nwElement0_0 _dafny.Map = _114_v4
	_ = _nwElement0_0
	var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(8))
	_ = _nw0
	(_nw0).ArraySet1(_nwElement0_0, 0)
	(_nw0).ArraySet1(_114_v4, 1)
	(_nw0).ArraySet1(_114_v4, 2)
	(_nw0).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v0, _111_v1), 3)
	(_nw0).ArraySet1(_114_v4, 4)
	(_nw0).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v0, _111_v1), 5)
	(_nw0).ArraySet1(_114_v4, 6)
	(_nw0).ArraySet1(_114_v4, 7)
	_115_v5 = _nw0
	var _116_globalState *GlobalState
	_ = _116_globalState
	var _nw1 *GlobalState = New_GlobalState_()
	_ = _nw1
	_nw1.Ctor__(false, _dafny.IntOfInt64(92), _dafny.IntOfInt64(431), _dafny.IntOfInt64(-645), false, _115_v5, true, true, false, _dafny.IntOfInt64(-357), true)
	_116_globalState = _nw1
	var _117_i0 _dafny.Int
	_ = _117_i0
	_117_i0 = _dafny.Zero
	{
		for _110_v0 {
			{
				if (_117_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_117_i0 = (_117_i0).Plus(_dafny.One)
				(_116_globalState).F6 = _110_v0
				if !(_110_v0) {
					var _118_v6 D0
					_ = _118_v6
					_118_v6 = Companion_D0_.Create_DC1_(_110_v0, _111_v1, _110_v0)
					var _119_v7 _dafny.Sequence
					_ = _119_v7
					_119_v7 = _dafny.SeqOf((_dafny.SetOf(_110_v0, _110_v0)).Cardinality())
					var _120_v8 _dafny.MultiSet
					_ = _120_v8
					_120_v8 = _dafny.MultiSetOf(_111_v1, _111_v1, _111_v1, (_dafny.Zero).Minus(_111_v1), _111_v1)
					var _121_v9 _dafny.Set
					_ = _121_v9
					_121_v9 = _dafny.SetOf((func() _dafny.Int {
						if (_120_v8).Contains(_111_v1) {
							return (_120_v8).Multiplicity(_111_v1)
						}
						return _dafny.IntOfInt64(664)
					})())
					var _122_v10 _dafny.Sequence
					_ = _122_v10
					_122_v10 = _dafny.SeqOf(_111_v1, (_118_v6).Dtor_cf2(), _111_v1, _dafny.IntOfUint32((_119_v7).Cardinality()), (_121_v9).Cardinality())
					var _rhs0 bool = !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_111_v1, _111_v1), _119_v7))
					_ = _rhs0
					var _rhs1 bool = _110_v0
					_ = _rhs1
					var _rhs2 bool = !(false) || (_110_v0)
					_ = _rhs2
					var _lhs0 *GlobalState = _116_globalState
					_ = _lhs0
					var _lhs1 *GlobalState = _116_globalState
					_ = _lhs1
					_lhs0.F7 = _rhs0
					_lhs1.F7 = _rhs1
					_110_v0 = _rhs2
					var _123_v11 _dafny.Array
					_ = _123_v11
					var _nw2 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
					_ = _nw2
					_123_v11 = _nw2
					_123_v11 = _123_v11
					var _124_v12 _dafny.Sequence
					_ = _124_v12
					_124_v12 = _dafny.SeqOf(_123_v11)
					var _125_v13 _dafny.Set
					_ = _125_v13
					_125_v13 = _dafny.SetOf(Companion_Default___.Fm0(_116_globalState))
					var _126_v14 _dafny.Map
					_ = _126_v14
					_126_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v1, _125_v13)
					var _127_v15 _dafny.Array
					_ = _127_v15
					var _nwElement0_1 _dafny.Int = (_dafny.Zero).Minus((_dafny.IntOfUint32((_119_v7).Cardinality())).Times((_113_v3).Cardinality()))
					_ = _nwElement0_1
					var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(17))
					_ = _nw3
					(_nw3).ArraySet1(_nwElement0_1, 0)
					(_nw3).ArraySet1(_111_v1, 1)
					(_nw3).ArraySet1(_111_v1, 2)
					(_nw3).ArraySet1(_111_v1, 3)
					(_nw3).ArraySet1(Companion_Default___.SafeDivisionInt(_111_v1, _111_v1), 4)
					(_nw3).ArraySet1(_dafny.IntOfUint32((_124_v12).Cardinality()), 5)
					(_nw3).ArraySet1((_dafny.Zero).Minus((_126_v14).Cardinality()), 6)
					(_nw3).ArraySet1(_111_v1, 7)
					(_nw3).ArraySet1(_111_v1, 8)
					(_nw3).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(753))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg29 _dafny.Int) interface{} {
							return coer29(arg29)
						}
					}(func(_128_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('s')
					}))).Cardinality()), 9)
					(_nw3).ArraySet1(_111_v1, 10)
					(_nw3).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_112_v2, _112_v2)).Cardinality()), 11)
					(_nw3).ArraySet1(((_dafny.Zero).Minus(_111_v1)).Times((_113_v3).Cardinality()), 12)
					(_nw3).ArraySet1(_111_v1, 13)
					(_nw3).ArraySet1(_111_v1, 14)
					(_nw3).ArraySet1((func() _dafny.Int {
						if false {
							return _111_v1
						}
						return _111_v1
					})(), 15)
					(_nw3).ArraySet1(_111_v1, 16)
					_127_v15 = _nw3
					var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_127_v15), 0))
					_ = _index0
					(_127_v15).ArraySet1(_111_v1, (_index0).Int())
					var _129_v16 D0
					_ = _129_v16
					_129_v16 = Companion_D0_.Create_DC0_(_111_v1)
					var _pat_let_tv0 = _111_v1
					_ = _pat_let_tv0
					var _130_v17 _dafny.Array
					_ = _130_v17
					var _nwElement0_2 D0 = Companion_D0_.Create_DC0_(_111_v1)
					_ = _nwElement0_2
					var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(12))
					_ = _nw4
					(_nw4).ArraySet1(_nwElement0_2, 0)
					(_nw4).ArraySet1(_129_v16, 1)
					(_nw4).ArraySet1(_129_v16, 2)
					(_nw4).ArraySet1(_129_v16, 3)
					(_nw4).ArraySet1(_129_v16, 4)
					(_nw4).ArraySet1(_129_v16, 5)
					(_nw4).ArraySet1(Companion_D0_.Create_DC0_(_111_v1), 6)
					(_nw4).ArraySet1(_129_v16, 7)
					(_nw4).ArraySet1(_129_v16, 8)
					(_nw4).ArraySet1(_129_v16, 9)
					(_nw4).ArraySet1(_129_v16, 10)
					(_nw4).ArraySet1(func(_pat_let0_0 D0) D0 {
						return func(_131_dt__update__tmp_h0 D0) D0 {
							return func(_pat_let1_0 _dafny.Int) D0 {
								return func(_132_dt__update_hcf0_h0 _dafny.Int) D0 {
									return Companion_D0_.Create_DC0_(_132_dt__update_hcf0_h0)
								}(_pat_let1_0)
							}(_pat_let_tv0)
						}(_pat_let0_0)
					}(_129_v16), 11)
					_130_v17 = _nw4
					var _133_v18 _dafny.MultiSet
					_ = _133_v18
					_133_v18 = _dafny.MultiSetOf(_130_v17, _130_v17)
					var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_127_v15), 0))
					_ = _index1
					var _rhs3 bool = _110_v0
					_ = _rhs3
					var _rhs4 bool = (_133_v18).IsProperSubsetOf(((_133_v18).Update(_130_v17, Companion_Default___.Abs(_111_v1))).Difference(_133_v18))
					_ = _rhs4
					var _rhs5 _dafny.Int = (_dafny.Zero).Minus((_111_v1).Minus(_111_v1))
					_ = _rhs5
					var _lhs2 *GlobalState = _116_globalState
					_ = _lhs2
					var _lhs3 *GlobalState = _116_globalState
					_ = _lhs3
					var _lhs4 _dafny.Array = _127_v15
					_ = _lhs4
					var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_127_v15), 0))
					_ = _lhs5
					_lhs2.F6 = _rhs3
					_lhs3.F6 = _rhs4
					(_lhs4).ArraySet1(_rhs5, (_lhs5).Int())
					var _rhs6 _dafny.Array = _127_v15
					_ = _rhs6
					var _rhs7 bool = (_dafny.IntOfUint32((_119_v7).Cardinality())).Cmp(_111_v1) > 0
					_ = _rhs7
					var _rhs8 _dafny.Int = ((_127_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_127_v15), 0))).Int()).(_dafny.Int)).Times(_111_v1)
					_ = _rhs8
					var _lhs6 *GlobalState = _116_globalState
					_ = _lhs6
					_127_v15 = _rhs6
					_lhs6.F6 = _rhs7
					_111_v1 = _rhs8
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_127_v15), 0))
					_ = _index2
					(_127_v15).ArraySet1(_111_v1, (_index2).Int())
				} else {
					var _134_v19 _dafny.Sequence
					_ = _134_v19
					_134_v19 = _dafny.UnicodeSeqOfUtf8Bytes("prcfca")
					var _135_v20 _dafny.Map
					_ = _135_v20
					_135_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v0, _114_v4)
					var _136_v21 _dafny.Map
					_ = _136_v21
					_136_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(_116_globalState), _134_v19), Companion_Default___.Fm2(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xngayfcy")).Cardinality()), (_135_v20).Cardinality(), _116_globalState))
					var _137_v22 _dafny.Sequence
					_ = _137_v22
					_137_v22 = _dafny.SeqOf(_111_v1, _111_v1)
					_136_v21 = (_136_v21).Update((func() _dafny.Sequence {
						if _110_v0 {
							return _134_v19
						}
						return _134_v19
					})(), (_137_v22).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_111_v1), _dafny.IntOfUint32((_137_v22).Cardinality()))).Uint32()).(_dafny.Int))
					var _138_v23 _dafny.MultiSet
					_ = _138_v23
					_138_v23 = _dafny.MultiSetOf(Companion_Default___.Fm2(_111_v1, _111_v1, _116_globalState))
					var _139_v24 _dafny.Map
					_ = _139_v24
					_139_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v23, _110_v0)
					_139_v24 = (_139_v24).Update(_dafny.MultiSetFromSeq(_137_v22), _110_v0)
					var _140_v25 _dafny.Map
					_ = _140_v25
					_140_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _110_v0)
					var _141_v26 _dafny.Set
					_ = _141_v26
					_141_v26 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v0, _110_v0))
					var _142_v27 _dafny.Map
					_ = _142_v27
					_142_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v1, !(_110_v0))
					var _143_v28 D0
					_ = _143_v28
					_143_v28 = Companion_D0_.Create_DC1_(!(_110_v0), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v1, _110_v0)).Cardinality(), _110_v0)
					var _144_v29 _dafny.Array
					_ = _144_v29
					var _nwElement0_3 bool = _110_v0
					_ = _nwElement0_3
					var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(24))
					_ = _nw5
					(_nw5).ArraySet1(_nwElement0_3, 0)
					(_nw5).ArraySet1((_dafny.SetOf(_140_v25, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v0, _110_v0))).IsSubsetOf(_141_v26), 1)
					(_nw5).ArraySet1(_110_v0, 2)
					(_nw5).ArraySet1(_110_v0, 3)
					(_nw5).ArraySet1(!(_110_v0), 4)
					(_nw5).ArraySet1(_110_v0, 5)
					(_nw5).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_137_v22, _137_v22), 6)
					(_nw5).ArraySet1(_110_v0, 7)
					(_nw5).ArraySet1(!(!(true)), 8)
					(_nw5).ArraySet1((_138_v23).IsSubsetOf(_138_v23), 9)
					(_nw5).ArraySet1((true) || (_110_v0), 10)
					(_nw5).ArraySet1(_110_v0, 11)
					(_nw5).ArraySet1(_110_v0, 12)
					(_nw5).ArraySet1((_111_v1).Cmp(_dafny.IntOfInt64(135)) < 0, 13)
					(_nw5).ArraySet1(true, 14)
					(_nw5).ArraySet1(_110_v0, 15)
					(_nw5).ArraySet1(Companion_Default___.Fm0(_116_globalState), 16)
					(_nw5).ArraySet1(_110_v0, 17)
					(_nw5).ArraySet1((_dafny.IntOfInt64(930)).Cmp(_111_v1) == 0, 18)
					(_nw5).ArraySet1(_110_v0, 19)
					(_nw5).ArraySet1((func() bool {
						if (_142_v27).Contains(_111_v1) {
							return (_142_v27).Get(_111_v1).(bool)
						}
						return _110_v0
					})(), 20)
					(_nw5).ArraySet1((_143_v28).Dtor_cf1(), 21)
					(_nw5).ArraySet1(_110_v0, 22)
					(_nw5).ArraySet1(!(_142_v27).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v1, (_112_v2).Select((Companion_Default___.SafeIndex(_111_v1, _dafny.IntOfUint32((_112_v2).Cardinality()))).Uint32()).(bool))).Cardinality(), !(Companion_Default___.Fm0(_116_globalState)))), 23)
					_144_v29 = _nw5
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_144_v29), 0))
					_ = _index3
					(_144_v29).ArraySet1(_110_v0, (_index3).Int())
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_144_v29), 0))
					_ = _index4
					(_144_v29).ArraySet1(_110_v0, (_index4).Int())
					var _145_v30 _dafny.CodePoint
					_ = _145_v30
					_145_v30 = _dafny.CodePoint('v')
					var _146_v31 *C2
					_ = _146_v31
					var _nw6 *C2 = New_C2_()
					_ = _nw6
					_nw6.Ctor__(_145_v30, _111_v1)
					_146_v31 = _nw6
					var _147_v32 bool
					_ = _147_v32
					var _148_v33 bool
					_ = _148_v33
					var _149_v34 bool
					_ = _149_v34
					var _150_v35 _dafny.Int
					_ = _150_v35
					var _out0 bool
					_ = _out0
					var _out1 bool
					_ = _out1
					var _out2 bool
					_ = _out2
					var _out3 _dafny.Int
					_ = _out3
					_out0, _out1, _out2, _out3 = (_146_v31).M8(_116_globalState)
					_147_v32 = _out0
					_148_v33 = _out1
					_149_v34 = _out2
					_150_v35 = _out3
				}
				if _110_v0 {
					var _151_v36 _dafny.Set
					_ = _151_v36
					_151_v36 = _dafny.SetOf(_110_v0)
					var _152_v37 _dafny.Set
					_ = _152_v37
					_152_v37 = _dafny.SetOf(_110_v0, _110_v0, _110_v0, (_112_v2).Select((Companion_Default___.SafeIndex((_151_v36).Cardinality(), _dafny.IntOfUint32((_112_v2).Cardinality()))).Uint32()).(bool), _110_v0)
					var _153_v38 _dafny.Sequence
					_ = _153_v38
					_153_v38 = _dafny.SeqOf(_151_v36)
					var _154_v39 _dafny.Sequence
					_ = _154_v39
					_154_v39 = _dafny.UnicodeSeqOfUtf8Bytes("eua")
					var _155_v40 D7
					_ = _155_v40
					_155_v40 = Companion_D7_.Create_DC13_((_153_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_154_v39).Cardinality()), _dafny.IntOfUint32((_153_v38).Cardinality()))).Uint32()).(_dafny.Set), _111_v1, _111_v1)
					var _pat_let_tv1 = _111_v1
					_ = _pat_let_tv1
					var _pat_let_tv2 = _111_v1
					_ = _pat_let_tv2
					var _pat_let_tv3 = _116_globalState
					_ = _pat_let_tv3
					_111_v1 = (_111_v1).Times((func(_pat_let2_0 D7) D7 {
						return func(_156_dt__update__tmp_h1 D7) D7 {
							return func(_pat_let3_0 _dafny.Int) D7 {
								return func(_157_dt__update_hcf19_h0 _dafny.Int) D7 {
									return Companion_D7_.Create_DC13_((_156_dt__update__tmp_h1).Dtor_cf18(), _157_dt__update_hcf19_h0, (_156_dt__update__tmp_h1).Dtor_cf20())
								}(_pat_let3_0)
							}(Companion_Default___.Fm2(_pat_let_tv1, _pat_let_tv2, _pat_let_tv3))
						}(_pat_let2_0)
					}(_155_v40)).Dtor_cf19())
					_154_v39 = Companion_Default___.Fm1(_116_globalState)
					var _158_v41 _dafny.CodePoint
					_ = _158_v41
					_158_v41 = _dafny.CodePoint('n')
					var _159_v42 _dafny.Array
					_ = _159_v42
					var _len0_0 _dafny.Int = _dafny.IntOfInt64(8)
					_ = _len0_0
					var _nw7 _dafny.Array
					_ = _nw7
					if _len0_0.Cmp(_dafny.Zero) == 0 {
						_nw7 = _dafny.NewArray(_len0_0)
					} else {
						var _init0 func(_dafny.Int) _dafny.CodePoint = (func(_160_v41 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_161_i2 _dafny.Int) _dafny.CodePoint {
								return _160_v41
							}
						})(_158_v41)
						_ = _init0
						var _element0_0 = _init0(_dafny.Zero)
						_ = _element0_0
						_nw7 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
						(_nw7).ArraySet1CodePoint(_element0_0, 0)
						var _nativeLen0_0 = (_len0_0).Int()
						_ = _nativeLen0_0
						for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
							(_nw7).ArraySet1CodePoint(_init0(_dafny.IntOf(_i0_0)), _i0_0)
						}
					}
					_159_v42 = _nw7
					var _162_v43 D21
					_ = _162_v43
					_162_v43 = Companion_D21_.Create_DC50_(_159_v42, (_dafny.Zero).Minus(_111_v1), _110_v0)
					var _163_v44 _dafny.Sequence
					_ = _163_v44
					_163_v44 = _dafny.SeqOf(_154_v39)
					var _164_v45 _dafny.Sequence
					_ = _164_v45
					_164_v45 = _dafny.SeqOf(_111_v1, _111_v1, _111_v1, _111_v1, _dafny.IntOfUint32((_163_v44).Cardinality()))
					var _165_v46 _dafny.Array
					_ = _165_v46
					var _nwElement0_4 bool = true
					_ = _nwElement0_4
					var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(24))
					_ = _nw8
					(_nw8).ArraySet1(_nwElement0_4, 0)
					(_nw8).ArraySet1(!_dafny.Companion_Sequence_.Contains(_154_v39, _158_v41), 1)
					(_nw8).ArraySet1(Companion_Default___.Fm0(_116_globalState), 2)
					(_nw8).ArraySet1(_110_v0, 3)
					(_nw8).ArraySet1(Companion_Default___.Fm0(_116_globalState), 4)
					(_nw8).ArraySet1(_110_v0, 5)
					(_nw8).ArraySet1((_162_v43).Dtor_cf85(), 6)
					(_nw8).ArraySet1((_112_v2).Select((Companion_Default___.SafeIndex(_111_v1, _dafny.IntOfUint32((_112_v2).Cardinality()))).Uint32()).(bool), 7)
					(_nw8).ArraySet1(_110_v0, 8)
					(_nw8).ArraySet1(!(_110_v0), 9)
					(_nw8).ArraySet1((_111_v1).Cmp((_164_v45).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(_110_v0)).Cardinality(), _dafny.IntOfUint32((_164_v45).Cardinality()))).Uint32()).(_dafny.Int)) > 0, 10)
					(_nw8).ArraySet1(_110_v0, 11)
					(_nw8).ArraySet1(_110_v0, 12)
					(_nw8).ArraySet1(Companion_Default___.Fm0(_116_globalState), 13)
					(_nw8).ArraySet1(_110_v0, 14)
					(_nw8).ArraySet1((_111_v1).Cmp(_dafny.IntOfInt64(305)) == 0, 15)
					(_nw8).ArraySet1(!(_110_v0), 16)
					(_nw8).ArraySet1(_110_v0, 17)
					(_nw8).ArraySet1(true, 18)
					(_nw8).ArraySet1((!(_110_v0)) && (_110_v0), 19)
					(_nw8).ArraySet1(true, 20)
					(_nw8).ArraySet1(Companion_Default___.Fm0(_116_globalState), 21)
					(_nw8).ArraySet1(_110_v0, 22)
					(_nw8).ArraySet1(_110_v0, 23)
					_165_v46 = _nw8
					_165_v46 = _165_v46
					var _166_v47 _dafny.MultiSet
					_ = _166_v47
					_166_v47 = _dafny.MultiSetOf((_111_v1).Cmp(_111_v1) <= 0)
					var _167_v48 _dafny.Map
					_ = _167_v48
					_167_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v0, _110_v0)
					var _168_v49 _dafny.Map
					_ = _168_v49
					_168_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v0, _167_v48)
					var _169_v50 D16
					_ = _169_v50
					_169_v50 = Companion_D16_.Create_DC39_(_111_v1, _110_v0, (func() _dafny.Map {
						if (_168_v49).Contains(_110_v0) {
							return (_168_v49).Get(_110_v0).(_dafny.Map)
						}
						return _167_v48
					})(), _158_v41)
					var _170_v51 D6
					_ = _170_v51
					_170_v51 = Companion_D6_.Create_DC10_(_112_v2)
					var _171_v52 D19
					_ = _171_v52
					_171_v52 = Companion_D19_.Create_DC43_(_110_v0, (_169_v50).Dtor_cf65(), _110_v0, _170_v51, _154_v39)
					var _rhs9 _dafny.CodePoint = (_171_v52).Dtor_cf70()
					_ = _rhs9
					var _rhs10 bool = _110_v0
					_ = _rhs10
					var _rhs11 _dafny.MultiSet = _166_v47
					_ = _rhs11
					_158_v41 = _rhs9
					_110_v0 = _rhs10
					_166_v47 = _rhs11
					var _172_v53 _dafny.Array
					_ = _172_v53
					var _len0_1 _dafny.Int = _dafny.IntOfInt64(24)
					_ = _len0_1
					var _nw9 _dafny.Array
					_ = _nw9
					if _len0_1.Cmp(_dafny.Zero) == 0 {
						_nw9 = _dafny.NewArray(_len0_1)
					} else {
						var _init1 func(_dafny.Int) _dafny.Int = (func(_173_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_174_i3 _dafny.Int) _dafny.Int {
								return (_174_i3).Times(_173_v1)
							}
						})(_111_v1)
						_ = _init1
						var _element0_1 = _init1(_dafny.Zero)
						_ = _element0_1
						_nw9 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
						(_nw9).ArraySet1(_element0_1, 0)
						var _nativeLen0_1 = (_len0_1).Int()
						_ = _nativeLen0_1
						for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
							(_nw9).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
						}
					}
					_172_v53 = _nw9
					var _175_v54 _dafny.Array
					_ = _175_v54
					var _nwElement0_5 _dafny.Array = _172_v53
					_ = _nwElement0_5
					var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(29))
					_ = _nw10
					(_nw10).ArraySet1(_nwElement0_5, 0)
					(_nw10).ArraySet1(_172_v53, 1)
					(_nw10).ArraySet1(_172_v53, 2)
					(_nw10).ArraySet1(_172_v53, 3)
					(_nw10).ArraySet1(_172_v53, 4)
					(_nw10).ArraySet1(_172_v53, 5)
					(_nw10).ArraySet1(_172_v53, 6)
					(_nw10).ArraySet1(_172_v53, 7)
					(_nw10).ArraySet1(_172_v53, 8)
					(_nw10).ArraySet1(_172_v53, 9)
					(_nw10).ArraySet1(_172_v53, 10)
					(_nw10).ArraySet1(_172_v53, 11)
					(_nw10).ArraySet1(_172_v53, 12)
					(_nw10).ArraySet1(_172_v53, 13)
					(_nw10).ArraySet1(_172_v53, 14)
					(_nw10).ArraySet1(_172_v53, 15)
					(_nw10).ArraySet1(_172_v53, 16)
					(_nw10).ArraySet1(_172_v53, 17)
					(_nw10).ArraySet1(_172_v53, 18)
					(_nw10).ArraySet1(_172_v53, 19)
					(_nw10).ArraySet1(_172_v53, 20)
					(_nw10).ArraySet1(_172_v53, 21)
					(_nw10).ArraySet1(_172_v53, 22)
					(_nw10).ArraySet1(_172_v53, 23)
					(_nw10).ArraySet1(_172_v53, 24)
					(_nw10).ArraySet1(_172_v53, 25)
					(_nw10).ArraySet1(_172_v53, 26)
					(_nw10).ArraySet1(_172_v53, 27)
					(_nw10).ArraySet1(_172_v53, 28)
					_175_v54 = _nw10
					var _176_v55 _dafny.Array
					_ = _176_v55
					_176_v55 = _175_v54
					var _177_v56 _dafny.Array
					_ = _177_v56
					var _nwElement0_6 _dafny.Array = _175_v54
					_ = _nwElement0_6
					var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(23))
					_ = _nw11
					(_nw11).ArraySet1(_nwElement0_6, 0)
					(_nw11).ArraySet1(_176_v55, 1)
					(_nw11).ArraySet1(_176_v55, 2)
					(_nw11).ArraySet1(_176_v55, 3)
					(_nw11).ArraySet1(_175_v54, 4)
					(_nw11).ArraySet1(_176_v55, 5)
					(_nw11).ArraySet1(_176_v55, 6)
					(_nw11).ArraySet1(_176_v55, 7)
					(_nw11).ArraySet1(_176_v55, 8)
					(_nw11).ArraySet1(_176_v55, 9)
					(_nw11).ArraySet1(_176_v55, 10)
					(_nw11).ArraySet1(_176_v55, 11)
					(_nw11).ArraySet1(_176_v55, 12)
					(_nw11).ArraySet1(_175_v54, 13)
					(_nw11).ArraySet1(_175_v54, 14)
					(_nw11).ArraySet1(_176_v55, 15)
					(_nw11).ArraySet1(_176_v55, 16)
					(_nw11).ArraySet1(_176_v55, 17)
					(_nw11).ArraySet1(_176_v55, 18)
					(_nw11).ArraySet1(_175_v54, 19)
					(_nw11).ArraySet1(_175_v54, 20)
					(_nw11).ArraySet1(_176_v55, 21)
					(_nw11).ArraySet1(_176_v55, 22)
					_177_v56 = _nw11
					var _178_v57 _dafny.Map
					_ = _178_v57
					_178_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v56, _172_v53)
					var _179_v58 D11
					_ = _179_v58
					_179_v58 = Companion_D11_.Create_DC22_(_178_v57)
					_179_v58 = _179_v58
				} else {
					var _180_v59 _dafny.CodePoint
					_ = _180_v59
					_180_v59 = _dafny.CodePoint('k')
					var _181_v60 _dafny.Sequence
					_ = _181_v60
					_181_v60 = _dafny.SeqOf(_180_v59, _180_v59, _180_v59, _180_v59, _180_v59)
					(_116_globalState).F1 = (_dafny.MultiSetFromSeq(_181_v60)).Cardinality()
					_110_v0 = _110_v0
					var _182_v61 _dafny.MultiSet
					_ = _182_v61
					_182_v61 = _dafny.MultiSetOf(_112_v2)
					var _183_v62 _dafny.Array
					_ = _183_v62
					var _nwElement0_7 _dafny.MultiSet = (_182_v61).Union(_182_v61)
					_ = _nwElement0_7
					var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(16))
					_ = _nw12
					(_nw12).ArraySet1(_nwElement0_7, 0)
					(_nw12).ArraySet1((_dafny.MultiSetOf(_112_v2, _112_v2, _112_v2)).Difference(_dafny.MultiSetOf(_112_v2, _112_v2, _112_v2, _112_v2, _dafny.SeqOf(!(_110_v0)))), 1)
					(_nw12).ArraySet1(_182_v61, 2)
					(_nw12).ArraySet1((_182_v61).Update(_112_v2, Companion_Default___.Abs(Companion_Default___.Fm2(_111_v1, _111_v1, _116_globalState))), 3)
					(_nw12).ArraySet1(_182_v61, 4)
					(_nw12).ArraySet1(_dafny.MultiSetOf(_112_v2), 5)
					(_nw12).ArraySet1((_182_v61).Difference(_182_v61), 6)
					(_nw12).ArraySet1(_182_v61, 7)
					(_nw12).ArraySet1(_182_v61, 8)
					(_nw12).ArraySet1(_182_v61, 9)
					(_nw12).ArraySet1((_182_v61).Intersection(_dafny.MultiSetOf(_112_v2, _dafny.Companion_Sequence_.Update(_112_v2, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_110_v0)).Cardinality()), _dafny.IntOfUint32((_112_v2).Cardinality()))).Uint32(), true))), 10)
					(_nw12).ArraySet1(_182_v61, 11)
					(_nw12).ArraySet1(_182_v61, 12)
					(_nw12).ArraySet1(_182_v61, 13)
					(_nw12).ArraySet1(_182_v61, 14)
					(_nw12).ArraySet1((_dafny.MultiSetOf(_dafny.SeqOf(_110_v0))).Union(_182_v61), 15)
					_183_v62 = _nw12
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_183_v62), 0))
					_ = _index5
					(_183_v62).ArraySet1((_dafny.MultiSetOf(_112_v2, _112_v2)).Intersection(_182_v61), (_index5).Int())
					var _184_v63 _dafny.Sequence
					_ = _184_v63
					_184_v63 = _dafny.SeqOf(_dafny.SeqOf(_110_v0, _110_v0, true, _110_v0, _110_v0))
					var _185_v64 _dafny.MultiSet
					_ = _185_v64
					_185_v64 = _dafny.MultiSetOf(false)
					var _186_v65 _dafny.MultiSet
					_ = _186_v65
					_186_v65 = _dafny.MultiSetOf(_111_v1, _111_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ymwi"), (Companion_Default___.SafeIndex(_111_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ymwi")).Cardinality()))).Uint32(), _180_v59)).Cardinality()))
					var _187_v66 _dafny.Map
					_ = _187_v66
					_187_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_186_v65).Cardinality(), _112_v2)
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_183_v62), 0))
					_ = _index6
					(_183_v62).ArraySet1((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_184_v63, _dafny.SeqOf(_112_v2)), (Companion_Default___.SafeIndex((_185_v64).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_184_v63, _dafny.SeqOf(_112_v2))).Cardinality()))).Uint32(), (func() _dafny.Sequence {
						if (_187_v66).Contains(_111_v1) {
							return (_187_v66).Get(_111_v1).(_dafny.Sequence)
						}
						return _112_v2
					})()))).Update(_112_v2, Companion_Default___.Abs(Companion_Default___.Fm2(_111_v1, _111_v1, _116_globalState))), (_index6).Int())
					var _188_v67 _dafny.Array
					_ = _188_v67
					var _nwElement0_8 bool = _110_v0
					_ = _nwElement0_8
					var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(7))
					_ = _nw13
					(_nw13).ArraySet1(_nwElement0_8, 0)
					(_nw13).ArraySet1(!_dafny.Companion_Sequence_.Contains(_112_v2, _110_v0), 1)
					(_nw13).ArraySet1(!(_110_v0), 2)
					(_nw13).ArraySet1((func() bool {
						if _110_v0 {
							return _110_v0
						}
						return _110_v0
					})(), 3)
					(_nw13).ArraySet1(_110_v0, 4)
					(_nw13).ArraySet1(_110_v0, 5)
					(_nw13).ArraySet1(_110_v0, 6)
					_188_v67 = _nw13
					var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_188_v67), 0))
					_ = _index7
					(_188_v67).ArraySet1(_110_v0, (_index7).Int())
					var _189_v68 _dafny.Set
					_ = _189_v68
					_189_v68 = _dafny.SetOf(_110_v0, _110_v0)
					var _190_v69 T1
					_ = _190_v69
					var _nw14 *C8 = New_C8_()
					_ = _nw14
					_nw14.Ctor__(_110_v0, _189_v68, (_186_v65).Cardinality())
					_190_v69 = _nw14
					var _191_v70 _dafny.Set
					_ = _191_v70
					_191_v70 = _dafny.SetOf(_190_v69)
					var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_188_v67), 0))
					_ = _index8
					var _rhs12 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_112_v2).Cardinality()), (_dafny.Zero).Minus(((_191_v70).Cardinality()).Minus(_111_v1)))
					_ = _rhs12
					var _rhs13 _dafny.Int = (func() _dafny.Int {
						if (_113_v3).Contains(((_dafny.MultiSetOf((_190_v69).F11())).Cardinality()).Minus((_190_v69).F11())) {
							return (_113_v3).Get(((_dafny.MultiSetOf((_190_v69).F11())).Cardinality()).Minus((_190_v69).F11())).(_dafny.Int)
						}
						return (_dafny.Zero).Minus((_dafny.SetOf(_110_v0)).Cardinality())
					})()
					_ = _rhs13
					var _rhs14 bool = (Companion_Default___.SafeDivisionInt(_111_v1, _111_v1)).Cmp(_111_v1) > 0
					_ = _rhs14
					var _lhs7 *GlobalState = _116_globalState
					_ = _lhs7
					var _lhs8 *GlobalState = _116_globalState
					_ = _lhs8
					var _lhs9 _dafny.Array = _188_v67
					_ = _lhs9
					var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_188_v67), 0))
					_ = _lhs10
					_lhs7.F1 = _rhs12
					_lhs8.F1 = _rhs13
					(_lhs9).ArraySet1(_rhs14, (_lhs10).Int())
					var _192_v71 D0
					_ = _192_v71
					_192_v71 = Companion_D0_.Create_DC1_(false, (_190_v69).F11(), _110_v0)
					var _193_v72 _dafny.Array
					_ = _193_v72
					var _nwElement0_9 _dafny.Int = Companion_Default___.SafeDivisionInt(_111_v1, (_192_v71).Dtor_cf2())
					_ = _nwElement0_9
					var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.One)
					_ = _nw15
					(_nw15).ArraySet1(_nwElement0_9, 0)
					_193_v72 = _nw15
					var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_193_v72), 0))
					_ = _index9
					(_193_v72).ArraySet1(Companion_Default___.SafeModuloInt((_190_v69).F11(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("day")).Cardinality())), (_index9).Int())
					var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_193_v72), 0))
					_ = _index10
					var _rhs15 _dafny.Int = Companion_Default___.SafeModuloInt((_190_v69).F11(), _111_v1)
					_ = _rhs15
					var _rhs16 _dafny.CodePoint = _180_v59
					_ = _rhs16
					var _rhs17 _dafny.Int = (_190_v69).F11()
					_ = _rhs17
					var _lhs11 _dafny.Array = _193_v72
					_ = _lhs11
					var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_193_v72), 0))
					_ = _lhs12
					(_lhs11).ArraySet1(_rhs15, (_lhs12).Int())
					_180_v59 = _rhs16
					_111_v1 = _rhs17
				}
				var _194_v73 _dafny.Sequence
				_ = _194_v73
				_194_v73 = _dafny.SeqOf(_111_v1)
				var _195_v74 _dafny.MultiSet
				_ = _195_v74
				_195_v74 = _dafny.MultiSetOf(_111_v1)
				var _196_v75 _dafny.Array
				_ = _196_v75
				var _nwElement0_10 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_194_v73)).Update(_111_v1, Companion_Default___.Abs(_111_v1))
				_ = _nwElement0_10
				var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(4))
				_ = _nw16
				(_nw16).ArraySet1(_nwElement0_10, 0)
				(_nw16).ArraySet1(_dafny.MultiSetOf(_111_v1, _111_v1), 1)
				(_nw16).ArraySet1(_195_v74, 2)
				(_nw16).ArraySet1(_195_v74, 3)
				_196_v75 = _nw16
				var _197_v76 _dafny.Array
				_ = _197_v76
				_197_v76 = _196_v75
				_197_v76 = _197_v76
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _198_v77 _dafny.Sequence
	_ = _198_v77
	_198_v77 = _dafny.UnicodeSeqOfUtf8Bytes("lyaxksq")
	var _199_v78 D13
	_ = _199_v78
	_199_v78 = Companion_D13_.Create_DC28_(_dafny.Companion_Sequence_.Concatenate(_198_v77, _198_v77), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_111_v1, _dafny.IntOfInt64(668))), _110_v0)
	var _pat_let_tv4 = _110_v0
	_ = _pat_let_tv4
	var _pat_let_tv5 = _198_v77
	_ = _pat_let_tv5
	_199_v78 = func(_pat_let4_0 D13) D13 {
		return func(_200_dt__update__tmp_h3 D13) D13 {
			return func(_pat_let5_0 bool) D13 {
				return func(_201_dt__update_hcf42_h0 bool) D13 {
					return func(_pat_let6_0 _dafny.Sequence) D13 {
						return func(_202_dt__update_hcf40_h0 _dafny.Sequence) D13 {
							return Companion_D13_.Create_DC28_(_202_dt__update_hcf40_h0, (_200_dt__update__tmp_h3).Dtor_cf41(), _201_dt__update_hcf42_h0)
						}(_pat_let6_0)
					}(_pat_let_tv5)
				}(_pat_let5_0)
			}(_pat_let_tv4)
		}(_pat_let4_0)
	}(_199_v78)
	if _110_v0 {
		var _203_v79 _dafny.Array
		_ = _203_v79
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_2
		var _nw17 _dafny.Array
		_ = _nw17
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw17 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_204_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_205_i4 _dafny.Int) _dafny.Int {
					return (_205_i4).Plus(_204_v1)
				}
			})(_111_v1)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw17 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw17).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw17).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_203_v79 = _nw17
		var _206_v80 D19
		_ = _206_v80
		_206_v80 = Companion_D19_.Create_DC42_(_203_v79)
		_203_v79 = (_206_v80).Dtor_cf68()
		var _207_v81 _dafny.Array
		_ = _207_v81
		var _nw18 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(17))
		_ = _nw18
		_207_v81 = _nw18
		var _208_v82 _dafny.MultiSet
		_ = _208_v82
		_208_v82 = _dafny.MultiSetOf(_111_v1, _111_v1)
		var _209_v83 _dafny.Set
		_ = _209_v83
		_209_v83 = _dafny.SetOf(_110_v0, _110_v0)
		var _210_v84 *C4
		_ = _210_v84
		var _nw19 *C4 = New_C4_()
		_ = _nw19
		_nw19.Ctor__(_208_v82, _110_v0, _209_v83, (_113_v3).Cardinality())
		_210_v84 = _nw19
		var _211_v85 _dafny.Map
		_ = _211_v85
		_211_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_203_v79, _210_v84)
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_207_v81), 0))
		_ = _index11
		(_207_v81).ArraySet1((func() *C4 {
			if (_211_v85).Contains(_203_v79) {
				return (_211_v85).Get(_203_v79).(*C4)
			}
			return _210_v84
		})(), (_index11).Int())
		var _212_v86 _dafny.Sequence
		_ = _212_v86
		_212_v86 = _dafny.SeqOf(_dafny.IntOfUint32((_112_v2).Cardinality()))
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_207_v81), 0))
		_ = _index12
		var _rhs18 *C4 = _210_v84
		_ = _rhs18
		var _rhs19 bool = (_210_v84).F17()
		_ = _rhs19
		var _rhs20 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_212_v86, _212_v86)
		_ = _rhs20
		var _lhs13 _dafny.Array = _207_v81
		_ = _lhs13
		var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_207_v81), 0))
		_ = _lhs14
		var _lhs15 *GlobalState = _116_globalState
		_ = _lhs15
		(_lhs13).ArraySet1(_rhs18, (_lhs14).Int())
		_lhs15.F7 = _rhs19
		_212_v86 = _rhs20
		var _213_v87 *C1
		_ = _213_v87
		var _nw20 *C1 = New_C1_()
		_ = _nw20
		_nw20.Ctor__(_110_v0, (_210_v84).F17())
		_213_v87 = _nw20
		var _214_v88 _dafny.Set
		_ = _214_v88
		_214_v88 = _dafny.SetOf(_213_v87, _213_v87, _213_v87)
		if (Companion_Default___.SafeModuloInt((_214_v88).Cardinality(), _111_v1)).Cmp(_111_v1) == 0 {
			(_116_globalState).F6 = (_210_v84).F17()
			(_116_globalState).F1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_210_v84).F17()), _112_v2)).Cardinality())
			var _215_v89 _dafny.Array
			_ = _215_v89
			var _nwElement0_11 _dafny.Sequence = _198_v77
			_ = _nwElement0_11
			var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(11))
			_ = _nw21
			(_nw21).ArraySet1(_nwElement0_11, 0)
			(_nw21).ArraySet1(Companion_Default___.Fm1(_116_globalState), 1)
			(_nw21).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(722))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}(func(_216_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('v')
			})), 2)
			(_nw21).ArraySet1(_198_v77, 3)
			(_nw21).ArraySet1(_198_v77, 4)
			(_nw21).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(181))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}(func(_217_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			})), 5)
			(_nw21).ArraySet1(_198_v77, 6)
			(_nw21).ArraySet1(Companion_Default___.Fm1(_116_globalState), 7)
			(_nw21).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ltloi"), 8)
			(_nw21).ArraySet1(_198_v77, 9)
			(_nw21).ArraySet1(_198_v77, 10)
			_215_v89 = _nw21
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_215_v89), 0))
			_ = _index13
			(_215_v89).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("brktfp"), (_index13).Int())
			var _218_v90 _dafny.MultiSet
			_ = _218_v90
			_218_v90 = _dafny.MultiSetOf(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("tk"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(121))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}(func(_219_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			}))))
			var _220_v91 D24
			_ = _220_v91
			_220_v91 = Companion_D24_.Create_DC58_((_210_v84).F16())
			var _221_v92 _dafny.Map
			_ = _221_v92
			_221_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_220_v91).Dtor_cf96(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_212_v86).Cardinality()), Companion_Default___.Fm2(_111_v1, _dafny.IntOfUint32((_112_v2).Cardinality()), _116_globalState)))
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_215_v89), 0))
			_ = _index14
			var _rhs21 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_198_v77, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-425))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}(func(_222_i8 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			})), _198_v77))
			_ = _rhs21
			var _rhs22 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_198_v77, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(792))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}(func(_223_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			})))
			_ = _rhs22
			var _rhs23 bool = (_210_v84).Fm4((_111_v1).Times(_111_v1), _221_v92, _111_v1, _110_v0, _116_globalState)
			_ = _rhs23
			var _rhs24 _dafny.Int = _111_v1
			_ = _rhs24
			var _rhs25 _dafny.MultiSet = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_213_v87).F18()), _112_v2))
			_ = _rhs25
			var _lhs16 _dafny.Array = _215_v89
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_215_v89), 0))
			_ = _lhs17
			var _lhs18 *GlobalState = _116_globalState
			_ = _lhs18
			var _lhs19 *GlobalState = _116_globalState
			_ = _lhs19
			_198_v77 = _rhs21
			(_lhs16).ArraySet1(_rhs22, (_lhs17).Int())
			_lhs18.F7 = _rhs23
			_lhs19.F1 = _rhs24
			_218_v90 = _rhs25
			var _224_v93 _dafny.CodePoint
			_ = _224_v93
			_224_v93 = _dafny.CodePoint('a')
			var _225_v94 _dafny.Map
			_ = _225_v94
			_225_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v1, (_210_v84).F16())
			var _226_v95 _dafny.Map
			_ = _226_v95
			_226_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_225_v94).Cardinality(), (_213_v87).F19())
			var _227_v96 _dafny.Map
			_ = _227_v96
			_227_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v93, _226_v95)
			var _228_v97 _dafny.Array
			_ = _228_v97
			var _229_v98 _dafny.Map
			_ = _229_v98
			var _out4 _dafny.Array
			_ = _out4
			var _out5 _dafny.Map
			_ = _out5
			_out4, _out5 = (_210_v84).M1(((func() _dafny.Map {
				if (_227_v96).Contains(_224_v93) {
					return (_227_v96).Get(_224_v93).(_dafny.Map)
				}
				return _226_v95
			})()).Cardinality(), _116_globalState)
			_228_v97 = _out4
			_229_v98 = _out5
			var _rhs26 _dafny.Int = _dafny.IntOfUint32((_198_v77).Cardinality())
			_ = _rhs26
			var _rhs27 _dafny.Int = (_111_v1).Plus(_111_v1)
			_ = _rhs27
			var _rhs28 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_111_v1, (_dafny.Zero).Minus(_111_v1)))
			_ = _rhs28
			var _rhs29 _dafny.Int = _dafny.IntOfInt64(271)
			_ = _rhs29
			var _rhs30 bool = (_213_v87).F18()
			_ = _rhs30
			var _lhs20 *GlobalState = _116_globalState
			_ = _lhs20
			var _lhs21 *GlobalState = _116_globalState
			_ = _lhs21
			var _lhs22 *GlobalState = _116_globalState
			_ = _lhs22
			var _lhs23 *GlobalState = _116_globalState
			_ = _lhs23
			var _lhs24 *GlobalState = _116_globalState
			_ = _lhs24
			_lhs20.F1 = _rhs26
			_lhs21.F1 = _rhs27
			_lhs22.F1 = _rhs28
			_lhs23.F1 = _rhs29
			_lhs24.F7 = _rhs30
		} else {
			var _230_v99 _dafny.Map
			_ = _230_v99
			_230_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_213_v87).F18(), (_213_v87).F18())
			_230_v99 = (_230_v99).Update(true, (_213_v87).F18())
			var _231_v100 _dafny.Array
			_ = _231_v100
			var _232_v101 _dafny.Map
			_ = _232_v101
			var _out6 _dafny.Array
			_ = _out6
			var _out7 _dafny.Map
			_ = _out7
			_out6, _out7 = (_210_v84).M1(_111_v1, _116_globalState)
			_231_v100 = _out6
			_232_v101 = _out7
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_203_v79), 0))
			_ = _index15
			(_203_v79).ArraySet1(Companion_Default___.SafeDivisionInt(_111_v1, _111_v1), (_index15).Int())
			var _233_v102 _dafny.Set
			_ = _233_v102
			_233_v102 = _dafny.SetOf(_111_v1)
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_203_v79), 0))
			_ = _index16
			(_203_v79).ArraySet1((_210_v84).Fm3(_111_v1, ((_dafny.Zero).Minus((_233_v102).Cardinality())).Cmp(Companion_Default___.Fm2(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("a")).Cardinality()), _111_v1, _116_globalState)) > 0, (_213_v87).F19(), _116_globalState), (_index16).Int())
			var _234_v103 _dafny.CodePoint
			_ = _234_v103
			_234_v103 = _dafny.CodePoint('p')
			var _235_v104 T0
			_ = _235_v104
			var _nw22 *C2 = New_C2_()
			_ = _nw22
			_nw22.Ctor__(_234_v103, _111_v1)
			_235_v104 = _nw22
			var _236_v105 _dafny.Array
			_ = _236_v105
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_3
			var _nw23 _dafny.Array
			_ = _nw23
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw23 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) bool = (func(_237_v87 *C1) func(_dafny.Int) bool {
					return func(_238_i10 _dafny.Int) bool {
						return (_237_v87).F19()
					}
				})(_213_v87)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw23 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw23).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw23).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_236_v105 = _nw23
			var _239_v106 _dafny.Map
			_ = _239_v106
			_239_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_236_v105, _dafny.IntOfUint32((_212_v86).Cardinality()))
			var _nw24 *C2 = New_C2_()
			_ = _nw24
			_nw24.Ctor__(_234_v103, (func() _dafny.Int {
				if (_239_v106).Contains(_236_v105) {
					return (_239_v106).Get(_236_v105).(_dafny.Int)
				}
				return _111_v1
			})())
			_235_v104 = _nw24
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_203_v79), 0))
			_ = _index17
			(_203_v79).ArraySet1(_dafny.IntOfInt64(146), (_index17).Int())
		}
		var _240_v107 _dafny.Array
		_ = _240_v107
		var _241_v108 _dafny.Map
		_ = _241_v108
		var _out8 _dafny.Array
		_ = _out8
		var _out9 _dafny.Map
		_ = _out9
		_out8, _out9 = (_210_v84).M1(_111_v1, _116_globalState)
		_240_v107 = _out8
		_241_v108 = _out9
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_203_v79), 0))
		_ = _index18
		(_203_v79).ArraySet1((_210_v84).Fm6(_110_v0, _111_v1, _116_globalState), (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_203_v79), 0))
		_ = _index19
		var _rhs31 _dafny.Int = (_111_v1).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_198_v77).Cardinality()), _111_v1))
		_ = _rhs31
		var _rhs32 _dafny.Int = _111_v1
		_ = _rhs32
		var _rhs33 _dafny.Int = (_111_v1).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("s")).Cardinality()))
		_ = _rhs33
		var _lhs25 _dafny.Array = _203_v79
		_ = _lhs25
		var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_203_v79), 0))
		_ = _lhs26
		var _lhs27 *GlobalState = _116_globalState
		_ = _lhs27
		(_lhs25).ArraySet1(_rhs31, (_lhs26).Int())
		_111_v1 = _rhs32
		_lhs27.F1 = _rhs33
	} else {
		var _242_v109 _dafny.MultiSet
		_ = _242_v109
		_242_v109 = _dafny.MultiSetOf(_111_v1)
		var _243_v110 _dafny.Set
		_ = _243_v110
		_243_v110 = _dafny.SetOf(_110_v0)
		var _244_v111 *C3
		_ = _244_v111
		var _nw25 *C3 = New_C3_()
		_ = _nw25
		_nw25.Ctor__((func() _dafny.Int {
			if _110_v0 {
				return _111_v1
			}
			return (_114_v4).Cardinality()
		})(), _111_v1, Companion_Default___.SafeModuloInt((_242_v109).Cardinality(), _111_v1), _243_v110)
		_244_v111 = _nw25
		_244_v111 = _244_v111
		var _245_v112 _dafny.Array
		_ = _245_v112
		var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
		_ = _nw26
		_245_v112 = _nw26
		var _246_v113 _dafny.CodePoint
		_ = _246_v113
		_246_v113 = _dafny.CodePoint('r')
		var _247_v114 _dafny.Map
		_ = _247_v114
		_247_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_v112, _246_v113)
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))
		_ = _index20
		(_245_v112).ArraySet1((_244_v111.F23).Minus((func() _dafny.Int {
			if (_113_v3).Contains((_247_v114).Cardinality()) {
				return (_113_v3).Get((_247_v114).Cardinality()).(_dafny.Int)
			}
			return (_244_v111).F22()
		})()), (_index20).Int())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))
		_ = _index21
		(_245_v112).ArraySet1((_244_v111).F22(), (_index21).Int())
		_111_v1 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-630), (func() _dafny.Int {
			if _110_v0 {
				return (_244_v111).F22()
			}
			return (_dafny.Zero).Minus((_245_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))).Int()).(_dafny.Int))
		})())
		var _248_v115 _dafny.Array
		_ = _248_v115
		var _249_v116 _dafny.Map
		_ = _249_v116
		var _out10 _dafny.Array
		_ = _out10
		var _out11 _dafny.Map
		_ = _out11
		_out10, _out11 = (_244_v111).M1((_244_v111).F22(), _116_globalState)
		_248_v115 = _out10
		_249_v116 = _out11
		var _250_v117 _dafny.Sequence
		_ = _250_v117
		_250_v117 = _dafny.SeqOf(_111_v1)
		var _source5 D7 = Companion_D7_.Create_DC13_((_243_v110).Union(_243_v110), (_244_v111).F22(), _dafny.IntOfUint32((_250_v117).Cardinality()))
		_ = _source5
		if _source5.Is_DC13() {
			var _251___mcc_h0 _dafny.Set = _source5.Get_().(D7_DC13).Cf18
			_ = _251___mcc_h0
			var _252___mcc_h1 _dafny.Int = _source5.Get_().(D7_DC13).Cf19
			_ = _252___mcc_h1
			var _253___mcc_h2 _dafny.Int = _source5.Get_().(D7_DC13).Cf20
			_ = _253___mcc_h2
			var _254_cf20 _dafny.Int = _253___mcc_h2
			_ = _254_cf20
			var _255_cf19 _dafny.Int = _252___mcc_h1
			_ = _255_cf19
			var _256_cf18 _dafny.Set = _251___mcc_h0
			_ = _256_cf18
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))
			_ = _index22
			(_245_v112).ArraySet1((_245_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))).Int()).(_dafny.Int), (_index22).Int())
			var _257_v118 D22
			_ = _257_v118
			_257_v118 = Companion_D22_.Create_DC52_()
			_257_v118 = _257_v118
			var _258_v119 _dafny.Map
			_ = _258_v119
			_258_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v0, true)
			var _259_v120 _dafny.Set
			_ = _259_v120
			_259_v120 = _dafny.SetOf((func() _dafny.Int {
				if (_242_v109).Contains(Companion_Default___.Fm2((_245_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))).Int()).(_dafny.Int), (_245_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))).Int()).(_dafny.Int), _116_globalState)) {
					return (_242_v109).Multiplicity(Companion_Default___.Fm2((_245_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))).Int()).(_dafny.Int), (_245_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))).Int()).(_dafny.Int), _116_globalState))
				}
				return _254_cf20
			})(), _244_v111.F23, (_dafny.IntOfInt64(866)).Plus(_dafny.IntOfInt64(72)), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(807))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}((func(_260_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_261_i11 _dafny.Int) _dafny.Int {
					return _260_v1
				}
			})(_111_v1))), (Companion_Default___.SafeIndex((_258_v119).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(807))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}((func(_262_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_263_i11 _dafny.Int) _dafny.Int {
					return _262_v1
				}
			})(_111_v1)))).Cardinality()))).Uint32(), (_dafny.Zero).Minus(_244_v111.F23))).Cardinality()), _111_v1), (_258_v119).Cardinality())
			_259_v120 = _dafny.SetOf(_111_v1)
			var _264_v122 _dafny.Sequence
			_ = _264_v122
			_264_v122 = _dafny.SeqOf(_dafny.SetOf((func() _dafny.Int {
				if (_242_v109).Contains((_245_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))).Int()).(_dafny.Int)) {
					return (_242_v109).Multiplicity((_245_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))).Int()).(_dafny.Int))
				}
				return (_244_v111).F22()
			})()), _dafny.SetOf(_dafny.IntOfInt64(261)), func() _dafny.Set {
				var _coll28 = _dafny.NewBuilder()
				_ = _coll28
				for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(113), _dafny.IntOfInt64(-958))); ; {
					_compr_28, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _265_v121 _dafny.Int
					_265_v121 = interface{}(_compr_28).(_dafny.Int)
					if ((_dafny.IntOfInt64(113)).Cmp(_265_v121) <= 0) && ((_265_v121).Cmp(_dafny.IntOfInt64(-958)) < 0) {
						_coll28.Add((_265_v121).Times(_111_v1))
					}
				}
				return _coll28.ToSet()
			}())
			_264_v122 = _264_v122
		} else if _source5.Is_DC12() {
			var _266___mcc_h3 _dafny.Sequence = _source5.Get_().(D7_DC12).Cf17
			_ = _266___mcc_h3
			var _267_cf17 _dafny.Sequence = _266___mcc_h3
			_ = _267_cf17
			var _268_v123 _dafny.Map
			_ = _268_v123
			_268_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v0, _dafny.Companion_Sequence_.Concatenate(_198_v77, _198_v77))
			_268_v123 = (_268_v123).Update(false, _dafny.Companion_Sequence_.Concatenate(_198_v77, _198_v77))
			(_116_globalState).F6 = !((Companion_Default___.Fm2(_dafny.IntOfUint32((_198_v77).Cardinality()), _244_v111.F23, _116_globalState)).Cmp((_245_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))).Int()).(_dafny.Int)) < 0)
			var _269_v124 _dafny.Array
			_ = _269_v124
			var _nw27 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw27
			_269_v124 = _nw27
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_269_v124), 0))
			_ = _index23
			(_269_v124).ArraySet1(_110_v0, (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_269_v124), 0))
			_ = _index24
			(_269_v124).ArraySet1(Companion_Default___.Fm0(_116_globalState), (_index24).Int())
			var _270_v125 *C4
			_ = _270_v125
			var _nw28 *C4 = New_C4_()
			_ = _nw28
			_nw28.Ctor__(_242_v109, false, _243_v110, (_245_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))).Int()).(_dafny.Int))
			_270_v125 = _nw28
			var _271_v126 D25
			_ = _271_v126
			_271_v126 = Companion_D25_.Create_DC61_(_270_v125)
			_270_v125 = (_271_v126).Dtor_cf98()
		} else {
			var _272___mcc_h4 D7 = _source5.Get_().(D7_DC14).Cf21
			_ = _272___mcc_h4
			var _273_cf21 D7 = _272___mcc_h4
			_ = _273_cf21
			var _274_v127 D6
			_ = _274_v127
			_274_v127 = Companion_D6_.Create_DC10_(_dafny.SeqOf(false, _110_v0))
			var _275_v128 D19
			_ = _275_v128
			_275_v128 = Companion_D19_.Create_DC43_(_110_v0, _246_v113, true, _274_v127, _198_v77)
			var _276_v129 _dafny.Array
			_ = _276_v129
			var _nwElement0_12 bool = true
			_ = _nwElement0_12
			var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(10))
			_ = _nw29
			(_nw29).ArraySet1(_nwElement0_12, 0)
			(_nw29).ArraySet1(_110_v0, 1)
			(_nw29).ArraySet1(!((_244_v111).Fm5(Companion_D0_.Create_DC0_(_244_v111.F23), _112_v2, _111_v1, (_275_v128).Dtor_cf70(), _116_globalState)), 2)
			(_nw29).ArraySet1(_110_v0, 3)
			(_nw29).ArraySet1(_110_v0, 4)
			(_nw29).ArraySet1(true, 5)
			(_nw29).ArraySet1(_110_v0, 6)
			(_nw29).ArraySet1(_110_v0, 7)
			(_nw29).ArraySet1(false, 8)
			(_nw29).ArraySet1(false, 9)
			_276_v129 = _nw29
			var _277_v130 _dafny.Set
			_ = _277_v130
			_277_v130 = _dafny.SetOf(_276_v129, _276_v129, _276_v129)
			_277_v130 = (_dafny.SetOf(_276_v129, _276_v129, _276_v129)).Union(_277_v130)
			_111_v1 = _111_v1
			(_116_globalState).F6 = (func() bool {
				if _110_v0 {
					return _110_v0
				}
				return !(func() _dafny.Set {
					var _coll29 = _dafny.NewBuilder()
					_ = _coll29
					for _iter29 := _dafny.Iterate((Companion_Default___.Fm27(_110_v0, _110_v0, _116_globalState)).Elements()); ; {
						_compr_29, _ok29 := _iter29()
						if !_ok29 {
							break
						}
						var _278_v131 _dafny.Int
						_278_v131 = interface{}(_compr_29).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm27(_110_v0, _110_v0, _116_globalState), _278_v131) {
							_coll29.Add(Companion_Default___.SafeDivisionInt(_278_v131, _dafny.IntOfInt64(-307)))
						}
					}
					return _coll29.ToSet()
				}()).Contains(_244_v111.F23)
			})()
			var _279_v132 *C7
			_ = _279_v132
			var _nw30 *C7 = New_C7_()
			_ = _nw30
			_nw30.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(966))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}((func(_280_v111 *C3, _281_v1 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_282_i12 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_280_v111.F23, _281_v1)
				}
			})(_244_v111, _111_v1))), _243_v110, (_245_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_245_v112), 0))).Int()).(_dafny.Int))
			_279_v132 = _nw30
		}
	}
	var _283_v133 _dafny.CodePoint
	_ = _283_v133
	_283_v133 = _dafny.CodePoint('o')
	var _284_v134 D6
	_ = _284_v134
	_284_v134 = Companion_D6_.Create_DC10_(_dafny.SeqOf(_110_v0, _110_v0))
	(_116_globalState).F7 = (Companion_D19_.Create_DC43_(false, _283_v133, _110_v0, _284_v134, _198_v77)).Dtor_cf71()
	var _285_i13 _dafny.Int
	_ = _285_i13
	_285_i13 = _dafny.Zero
	{
		for _110_v0 {
			{
				if (_285_i13).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_285_i13 = (_285_i13).Plus(_dafny.One)
				var _286_v135 *C5
				_ = _286_v135
				var _nw31 *C5 = New_C5_()
				_ = _nw31
				_nw31.Ctor__()
				_286_v135 = _nw31
				var _rhs34 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_110_v0, true, _110_v0), _dafny.SeqOf(_110_v0)), _112_v2)
				_ = _rhs34
				var _rhs35 _dafny.Int = _111_v1
				_ = _rhs35
				var _rhs36 _dafny.Int = _111_v1
				_ = _rhs36
				var _rhs37 *C5 = _286_v135
				_ = _rhs37
				var _rhs38 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if true {
						return _111_v1
					}
					return _111_v1
				})(), _111_v1)
				_ = _rhs38
				var _lhs28 *GlobalState = _116_globalState
				_ = _lhs28
				var _lhs29 *GlobalState = _116_globalState
				_ = _lhs29
				var _lhs30 *GlobalState = _116_globalState
				_ = _lhs30
				var _lhs31 *GlobalState = _116_globalState
				_ = _lhs31
				_lhs28.F6 = _rhs34
				_lhs29.F1 = _rhs35
				_lhs30.F1 = _rhs36
				_286_v135 = _rhs37
				_lhs31.F1 = _rhs38
				var _287_v136 _dafny.Set
				_ = _287_v136
				_287_v136 = _dafny.SetOf(_110_v0, _110_v0, _110_v0, _110_v0)
				var _288_v137 *C7
				_ = _288_v137
				var _nw32 *C7 = New_C7_()
				_ = _nw32
				_nw32.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(269))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}((func(_289_v0 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_290_i14 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_289_v0)).Cardinality()))
					}
				})(_110_v0))), Companion_Default___.Fm56(_111_v1, _111_v1, _116_globalState)), _287_v136, (_111_v1).Plus(_111_v1))
				_288_v137 = _nw32
				_288_v137 = _288_v137
				var _291_v138 _dafny.Array
				_ = _291_v138
				var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw33
				_291_v138 = _nw33
				var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_291_v138), 0))
				_ = _index25
				(_291_v138).ArraySet1((_dafny.Zero).Minus((_dafny.SetOf(_111_v1)).Cardinality()), (_index25).Int())
				var _292_v139 *C0
				_ = _292_v139
				var _nw34 *C0 = New_C0_()
				_ = _nw34
				_nw34.Ctor__(_112_v2, _111_v1)
				_292_v139 = _nw34
				var _293_v140 _dafny.MultiSet
				_ = _293_v140
				_293_v140 = _dafny.MultiSetOf(_292_v139, _292_v139)
				var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_291_v138), 0))
				_ = _index26
				(_291_v138).ArraySet1(Companion_Default___.SafeDivisionInt(_111_v1, ((_293_v140).Intersection(_293_v140)).Cardinality()), (_index26).Int())
				(_116_globalState).F7 = _110_v0
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	(_116_globalState).F7 = true
	(_116_globalState).F1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(812))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg39 _dafny.Int) interface{} {
			return coer39(arg39)
		}
	}((func(_294_v133 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_295_i15 _dafny.Int) _dafny.CodePoint {
			return _294_v133
		}
	})(_283_v133))), _dafny.Companion_Sequence_.Update(_198_v77, (Companion_Default___.SafeIndex(_111_v1, _dafny.IntOfUint32((_198_v77).Cardinality()))).Uint32(), _283_v133))).Cardinality())
	var _296_v141 _dafny.Array
	_ = _296_v141
	var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
	_ = _nw35
	_296_v141 = _nw35
	var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))
	_ = _index27
	(_296_v141).ArraySet1(_dafny.IntOfUint32((_112_v2).Cardinality()), (_index27).Int())
	var _297_v142 _dafny.Sequence
	_ = _297_v142
	_297_v142 = _dafny.SeqOf(_111_v1)
	var _298_v143 _dafny.Sequence
	_ = _298_v143
	_298_v143 = _dafny.SeqOf(_111_v1, (_dafny.Zero).Minus(_111_v1), _dafny.IntOfUint32((_297_v142).Cardinality()), _111_v1)
	var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))
	_ = _index28
	var _rhs39 _dafny.Int = _111_v1
	_ = _rhs39
	var _rhs40 bool = !(!(!_dafny.Companion_Sequence_.Equal(_297_v142, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-502))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg40 _dafny.Int) interface{} {
			return coer40(arg40)
		}
	}((func(_299_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_300_i16 _dafny.Int) _dafny.Int {
			return _299_v1
		}
	})(_111_v1))))))
	_ = _rhs40
	var _lhs32 _dafny.Array = _296_v141
	_ = _lhs32
	var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))
	_ = _lhs33
	var _lhs34 *GlobalState = _116_globalState
	_ = _lhs34
	(_lhs32).ArraySet1(_rhs39, (_lhs33).Int())
	_lhs34.F7 = _rhs40
	var _301_v145 D8
	_ = _301_v145
	_301_v145 = Companion_D8_.Create_DC16_(true, _111_v1)
	var _302_v146 D8
	_ = _302_v146
	_302_v146 = Companion_D8_.Create_DC17_(_301_v145)
	var _303_v147 _dafny.Sequence
	_ = _303_v147
	_303_v147 = _dafny.SeqOf(_302_v146, _302_v146)
	var _304_v148 _dafny.Map
	_ = _304_v148
	_304_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v133, func() _dafny.Map {
		var _coll30 = _dafny.NewMapBuilder()
		_ = _coll30
		for _iter30 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_303_v147, (Companion_Default___.SafeIndex((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_303_v147).Cardinality()))).Uint32(), _302_v146))).Elements()); ; {
			_compr_30, _ok30 := _iter30()
			if !_ok30 {
				break
			}
			var _305_v144 D8
			_305_v144 = interface{}(_compr_30).(D8)
			if (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_303_v147, (Companion_Default___.SafeIndex((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_303_v147).Cardinality()))).Uint32(), _302_v146))).Contains(_305_v144) {
				_coll30.Add(_305_v144, (_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int))
			}
		}
		return _coll30.ToMap()
	}())
	var _306_v149 _dafny.Map
	_ = _306_v149
	_306_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC17_(_301_v145), Companion_Default___.Fm2((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_298_v143).Cardinality()), _116_globalState))
	var _307_v150 D13
	_ = _307_v150
	_307_v150 = Companion_D13_.Create_DC27_((func() _dafny.Map {
		if (_304_v148).Contains(_283_v133) {
			return (_304_v148).Get(_283_v133).(_dafny.Map)
		}
		return _306_v149
	})())
	var _308_v151 _dafny.Map
	_ = _308_v151
	_308_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int), _307_v150)
	_308_v151 = (_308_v151).Update((_dafny.Zero).Minus((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int)), _307_v150)
	var _309_v152 _dafny.MultiSet
	_ = _309_v152
	_309_v152 = _dafny.MultiSetOf(_111_v1)
	var _310_v153 *C4
	_ = _310_v153
	var _nw36 *C4 = New_C4_()
	_ = _nw36
	_nw36.Ctor__(_309_v152, !(_110_v0), _dafny.SetOf(_110_v0), (func() _dafny.Int {
		if _110_v0 {
			return _111_v1
		}
		return (_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int)
	})())
	_310_v153 = _nw36
	var _311_v154 _dafny.Array
	_ = _311_v154
	var _nw37 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
	_ = _nw37
	_311_v154 = _nw37
	var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))
	_ = _index29
	(_311_v154).ArraySet1(_110_v0, (_index29).Int())
	var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))
	_ = _index30
	var _rhs41 bool = _110_v0
	_ = _rhs41
	var _rhs42 *C4 = (func() *C4 {
		if (_310_v153).F17() {
			return _310_v153
		}
		return _310_v153
	})()
	_ = _rhs42
	var _rhs43 bool = (_310_v153).F17()
	_ = _rhs43
	var _rhs44 bool = (_310_v153).F17()
	_ = _rhs44
	var _lhs35 *GlobalState = _116_globalState
	_ = _lhs35
	var _lhs36 *GlobalState = _116_globalState
	_ = _lhs36
	var _lhs37 _dafny.Array = _311_v154
	_ = _lhs37
	var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))
	_ = _lhs38
	_lhs35.F7 = _rhs41
	_310_v153 = _rhs42
	_lhs36.F7 = _rhs43
	(_lhs37).ArraySet1(_rhs44, (_lhs38).Int())
	(_116_globalState).F1 = Companion_Default___.SafeModuloInt(_111_v1, _dafny.IntOfInt64(223))
	var _312_v155 *C5
	_ = _312_v155
	var _nw38 *C5 = New_C5_()
	_ = _nw38
	_nw38.Ctor__()
	_312_v155 = _nw38
	var _313_v156 _dafny.Sequence
	_ = _313_v156
	_313_v156 = _dafny.SeqOf(_312_v155)
	var _314_v157 D20
	_ = _314_v157
	_314_v157 = Companion_D20_.Create_DC44_((_313_v156).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_198_v77).Cardinality()), _dafny.IntOfUint32((_313_v156).Cardinality()))).Uint32()).(*C5))
	_314_v157 = _314_v157
	var _315_v159 _dafny.Map
	_ = _315_v159
	_315_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
	var _316_v160 _dafny.Set
	_ = _316_v160
	_316_v160 = _dafny.SetOf(_dafny.MultiSetOf((_315_v159).Cardinality(), _111_v1))
	var _317_v161 _dafny.MultiSet
	_ = _317_v161
	_317_v161 = _dafny.MultiSetOf((_310_v153).Fm4(_111_v1, func() _dafny.Map {
		var _coll31 = _dafny.NewMapBuilder()
		_ = _coll31
		for _iter31 := _dafny.Iterate((_316_v160).Elements()); ; {
			_compr_31, _ok31 := _iter31()
			if !_ok31 {
				break
			}
			var _318_v158 _dafny.MultiSet
			_318_v158 = interface{}(_compr_31).(_dafny.MultiSet)
			if (_316_v160).Contains(_318_v158) {
				_coll31.Add(_318_v158, _113_v3)
			}
		}
		return _coll31.ToMap()
	}(), _111_v1, (_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool), _116_globalState), (_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool))
	var _319_v163 D22
	_ = _319_v163
	_319_v163 = Companion_D22_.Create_DC51_(func() _dafny.Set {
		var _coll32 = _dafny.NewBuilder()
		_ = _coll32
		for _iter32 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int), (_310_v153).F17())).Keys().Elements()); ; {
			_compr_32, _ok32 := _iter32()
			if !_ok32 {
				break
			}
			var _320_v162 _dafny.Int
			_320_v162 = interface{}(_compr_32).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int), (_310_v153).F17())).Contains(_320_v162) {
				_coll32.Add(Companion_Default___.SafeModuloInt(_320_v162, _dafny.IntOfInt64(-545)))
			}
		}
		return _coll32.ToSet()
	}())
	var _321_v164 _dafny.Sequence
	_ = _321_v164
	_321_v164 = _dafny.SeqOf(_319_v163)
	(_116_globalState).F6 = (_dafny.IntOfUint32((_297_v142).Cardinality())).Cmp(Companion_Default___.SafeModuloInt((func() _dafny.Int {
		if (_317_v161).Contains((_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool)) {
			return (_317_v161).Multiplicity((_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool))
		}
		return (_dafny.Zero).Minus(_111_v1)
	})(), (_dafny.MultiSetFromSeq(_321_v164)).Cardinality())) > 0
	for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_296_v141), 0))); ; {
		_guard_loop_0, _ok33 := _iter33()
		if !_ok33 {
			break
		}
		var _322_i17 _dafny.Int
		_322_i17 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_322_i17).Sign() != -1) && ((_322_i17).Cmp(_dafny.ArrayLen((_296_v141), 0)) < 0)) {
			(_296_v141).ArraySet1((_322_i17).Minus(Companion_Default___.SafeDivisionInt((_dafny.SetOf((func() _dafny.Set {
				var _coll33 = _dafny.NewBuilder()
				_ = _coll33
				for _iter34 := _dafny.Iterate((_dafny.MultiSetOf(_283_v133)).Elements()); ; {
					_compr_33, _ok34 := _iter34()
					if !_ok34 {
						break
					}
					var _323_v165 _dafny.CodePoint
					_323_v165 = interface{}(_compr_33).(_dafny.CodePoint)
					if (_dafny.MultiSetOf(_283_v133)).Contains(_323_v165) {
						_coll33.Add(_323_v165)
					}
				}
				return _coll33.ToSet()
			}()).Cardinality(), _111_v1)).Cardinality(), _dafny.IntOfUint32((_112_v2).Cardinality()))), (_322_i17).Int())
		}
	}
	var _324_v166 D24
	_ = _324_v166
	_324_v166 = Companion_D24_.Create_DC58_((_310_v153).F16())
	var _325_i18 _dafny.Int
	_ = _325_i18
	_325_i18 = _dafny.Zero
	{
		var _pat_let_tv6 = _298_v143
		_ = _pat_let_tv6
		var _pat_let_tv7 = _298_v143
		_ = _pat_let_tv7
		var _pat_let_tv8 = _111_v1
		_ = _pat_let_tv8
		var _pat_let_tv9 = _311_v154
		_ = _pat_let_tv9
		var _pat_let_tv10 = _311_v154
		_ = _pat_let_tv10
		var _pat_let_tv11 = _298_v143
		_ = _pat_let_tv11
		var _pat_let_tv12 = _111_v1
		_ = _pat_let_tv12
		var _pat_let_tv13 = _298_v143
		_ = _pat_let_tv13
		var _pat_let_tv14 = _296_v141
		_ = _pat_let_tv14
		var _pat_let_tv15 = _296_v141
		_ = _pat_let_tv15
		var _pat_let_tv16 = _296_v141
		_ = _pat_let_tv16
		var _pat_let_tv17 = _296_v141
		_ = _pat_let_tv17
		var _pat_let_tv18 = _111_v1
		_ = _pat_let_tv18
		var _pat_let_tv19 = _296_v141
		_ = _pat_let_tv19
		var _pat_let_tv20 = _296_v141
		_ = _pat_let_tv20
		for func(_source6 D24) bool {
			if _source6.Is_DC59() {
				var _351___mcc_h5 _dafny.CodePoint = _source6.Get_().(D24_DC59).Cf97
				_ = _351___mcc_h5
				var _352_cf97 _dafny.CodePoint = _351___mcc_h5
				_ = _352_cf97
				return !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_pat_let_tv6, _pat_let_tv7), _dafny.SeqOf(_dafny.SeqOf(_pat_let_tv8)))
			} else if _source6.Is_DC60() {
				return (_pat_let_tv10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_pat_let_tv9), 0))).Int()).(bool)
			} else {
				var _353___mcc_h6 _dafny.MultiSet = _source6.Get_().(D24_DC58).Cf96
				_ = _353___mcc_h6
				var _354_cf96 _dafny.MultiSet = _353___mcc_h6
				_ = _354_cf96
				return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_pat_let_tv11, (Companion_Default___.SafeIndex(_pat_let_tv12, _dafny.IntOfUint32((_pat_let_tv13).Cardinality()))).Uint32(), (_pat_let_tv15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_pat_let_tv14), 0))).Int()).(_dafny.Int)), _dafny.SeqOf((_pat_let_tv17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_pat_let_tv16), 0))).Int()).(_dafny.Int), _pat_let_tv18, (_pat_let_tv20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_pat_let_tv19), 0))).Int()).(_dafny.Int)))
			}
		}((func() D24 {
			if (_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool) {
				return Companion_D24_.Create_DC58_((_310_v153).F16())
			}
			return _324_v166
		})()) {
			{
				if (_325_i18).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_325_i18 = (_325_i18).Plus(_dafny.One)
				var _326_v167 *C0
				_ = _326_v167
				var _nw39 *C0 = New_C0_()
				_ = _nw39
				_nw39.Ctor__(_dafny.Companion_Sequence_.Concatenate(_112_v2, _112_v2), Companion_Default___.SafeModuloInt((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int), ((_310_v153).F16()).Cardinality()))
				_326_v167 = _nw39
				var _327_v168 _dafny.Array
				_ = _327_v168
				var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
				_ = _nw40
				_327_v168 = _nw40
				var _328_v169 _dafny.Array
				_ = _328_v169
				var _nwElement0_13 _dafny.Sequence = _198_v77
				_ = _nwElement0_13
				var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(27))
				_ = _nw41
				(_nw41).ArraySet1(_nwElement0_13, 0)
				(_nw41).ArraySet1(_198_v77, 1)
				(_nw41).ArraySet1(Companion_Default___.Fm1(_116_globalState), 2)
				(_nw41).ArraySet1(_198_v77, 3)
				(_nw41).ArraySet1(_198_v77, 4)
				(_nw41).ArraySet1(_198_v77, 5)
				(_nw41).ArraySet1(_198_v77, 6)
				(_nw41).ArraySet1(_198_v77, 7)
				(_nw41).ArraySet1(_198_v77, 8)
				(_nw41).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("smkfbwtdu"), 9)
				(_nw41).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(667))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}((func(_329_v133 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_330_i19 _dafny.Int) _dafny.CodePoint {
						return _329_v133
					}
				})(_283_v133))), 10)
				(_nw41).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mfcvhqiw"), 11)
				(_nw41).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(849))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_331_v133 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_332_i20 _dafny.Int) _dafny.CodePoint {
						return _331_v133
					}
				})(_283_v133))), 12)
				(_nw41).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("kjwo"), 13)
				(_nw41).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(996))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}(func(_333_i21 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				})), 14)
				(_nw41).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gnhdyskyq"), 15)
				(_nw41).ArraySet1(_198_v77, 16)
				(_nw41).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-3))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}((func(_334_v133 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_335_i22 _dafny.Int) _dafny.CodePoint {
						return _334_v133
					}
				})(_283_v133))), 17)
				(_nw41).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ykantl"), 18)
				(_nw41).ArraySet1(_198_v77, 19)
				(_nw41).ArraySet1(_198_v77, 20)
				(_nw41).ArraySet1(_dafny.Companion_Sequence_.Update(_198_v77, (Companion_Default___.SafeIndex((_326_v167).F21(), _dafny.IntOfUint32((_198_v77).Cardinality()))).Uint32(), _283_v133), 21)
				(_nw41).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(780))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}((func(_336_v133 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_337_i23 _dafny.Int) _dafny.CodePoint {
						return _336_v133
					}
				})(_283_v133))), 22)
				(_nw41).ArraySet1(_198_v77, 23)
				(_nw41).ArraySet1(_198_v77, 24)
				(_nw41).ArraySet1(_198_v77, 25)
				(_nw41).ArraySet1(_198_v77, 26)
				_328_v169 = _nw41
				var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_327_v168), 0))
				_ = _index31
				(_327_v168).ArraySet1(_328_v169, (_index31).Int())
				var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_327_v168), 0))
				_ = _index32
				(_327_v168).ArraySet1(_328_v169, (_index32).Int())
				_198_v77 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_198_v77, _198_v77), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bpbqxjl"), _198_v77))
				if (_310_v153).F17() {
					var _338_v170 _dafny.Sequence
					_ = _338_v170
					_338_v170 = _dafny.SeqOf(_317_v161, _317_v161, _317_v161, _317_v161)
					var _339_v171 _dafny.Set
					_ = _339_v171
					_339_v171 = _dafny.SetOf((_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool))
					var _340_v172 _dafny.Array
					_ = _340_v172
					var _nwElement0_14 _dafny.MultiSet = _317_v161
					_ = _nwElement0_14
					var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(25))
					_ = _nw42
					(_nw42).ArraySet1(_nwElement0_14, 0)
					(_nw42).ArraySet1(_dafny.MultiSetOf((_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool), (_310_v153).F17()), 1)
					(_nw42).ArraySet1((_dafny.MultiSetOf(!((_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool)))).Difference(_317_v161), 2)
					(_nw42).ArraySet1(_317_v161, 3)
					(_nw42).ArraySet1(_317_v161, 4)
					(_nw42).ArraySet1(_dafny.MultiSetFromSeq(_112_v2), 5)
					(_nw42).ArraySet1((Companion_Default___.Fm37(_116_globalState)).Union(_317_v161), 6)
					(_nw42).ArraySet1(_317_v161, 7)
					(_nw42).ArraySet1(_317_v161, 8)
					(_nw42).ArraySet1(_317_v161, 9)
					(_nw42).ArraySet1(_dafny.MultiSetOf((_326_v167).Fm21(true, _111_v1, _116_globalState), (_310_v153).F17(), (_310_v153).F17()), 10)
					(_nw42).ArraySet1(_dafny.MultiSetOf((_310_v153).F17()), 11)
					(_nw42).ArraySet1(_317_v161, 12)
					(_nw42).ArraySet1((_317_v161).Difference(_317_v161), 13)
					(_nw42).ArraySet1(_317_v161, 14)
					(_nw42).ArraySet1(_dafny.MultiSetOf((_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool), (_310_v153).F17(), !(_110_v0)), 15)
					(_nw42).ArraySet1(_317_v161, 16)
					(_nw42).ArraySet1((func() _dafny.MultiSet {
						if !((_310_v153).F17()) {
							return _317_v161
						}
						return _dafny.MultiSetFromSeq(_112_v2)
					})(), 17)
					(_nw42).ArraySet1((_338_v170).Select((Companion_Default___.SafeIndex(_111_v1, _dafny.IntOfUint32((_338_v170).Cardinality()))).Uint32()).(_dafny.MultiSet), 18)
					(_nw42).ArraySet1(_317_v161, 19)
					(_nw42).ArraySet1((_317_v161).Union(_317_v161), 20)
					(_nw42).ArraySet1((_317_v161).Difference(Companion_Default___.Fm37(_116_globalState)), 21)
					(_nw42).ArraySet1((_dafny.MultiSetOf(_110_v0, !((_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool)))).Difference(_317_v161), 22)
					(_nw42).ArraySet1(_317_v161, 23)
					(_nw42).ArraySet1((_317_v161).Update((_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool), Companion_Default___.Abs((_339_v171).Cardinality())), 24)
					_340_v172 = _nw42
					var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_340_v172), 0))
					_ = _index33
					(_340_v172).ArraySet1(_317_v161, (_index33).Int())
					var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_340_v172), 0))
					_ = _index34
					(_340_v172).ArraySet1((_317_v161).Union(_dafny.MultiSetOf((_310_v153).F17())), (_index34).Int())
					var _341_v173 _dafny.Map
					_ = _341_v173
					_341_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v4, !((_310_v153).F17()) || ((_310_v153).F17()))
					var _rhs45 _dafny.Int = (Companion_Default___.SafeModuloInt((_298_v143).Select((Companion_Default___.SafeIndex((_326_v167).F21(), _dafny.IntOfUint32((_298_v143).Cardinality()))).Uint32()).(_dafny.Int), (_326_v167).F21())).Plus(Companion_Default___.SafeModuloInt((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int), (_326_v167).F21()))
					_ = _rhs45
					var _rhs46 bool = (_310_v153).F17()
					_ = _rhs46
					var _rhs47 bool = (func() bool {
						if (_310_v153).F17() {
							return (_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool)
						}
						return !((_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool))
					})()
					_ = _rhs47
					var _rhs48 _dafny.Int = ((_326_v167).F21()).Minus(_dafny.IntOfUint32((_198_v77).Cardinality()))
					_ = _rhs48
					var _rhs49 _dafny.Map = _341_v173
					_ = _rhs49
					var _lhs39 *GlobalState = _116_globalState
					_ = _lhs39
					var _lhs40 *GlobalState = _116_globalState
					_ = _lhs40
					var _lhs41 *GlobalState = _116_globalState
					_ = _lhs41
					_lhs39.F1 = _rhs45
					_110_v0 = _rhs46
					_lhs40.F6 = _rhs47
					_lhs41.F1 = _rhs48
					_341_v173 = _rhs49
					var _342_v174 _dafny.Array
					_ = _342_v174
					var _nw43 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(29))
					_ = _nw43
					_342_v174 = _nw43
					var _343_v175 *C8
					_ = _343_v175
					var _nw44 *C8 = New_C8_()
					_ = _nw44
					_nw44.Ctor__(!((_310_v153).F17()), _339_v171, _dafny.IntOfUint32((_198_v77).Cardinality()))
					_343_v175 = _nw44
					var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_342_v174), 0))
					_ = _index35
					(_342_v174).ArraySet1(_343_v175, (_index35).Int())
					var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_342_v174), 0))
					_ = _index36
					(_342_v174).ArraySet1(_343_v175, (_index36).Int())
					var _344_v176 *C3
					_ = _344_v176
					var _nw45 *C3 = New_C3_()
					_ = _nw45
					_nw45.Ctor__(_111_v1, _dafny.IntOfInt64(282), (_dafny.IntOfUint32((_198_v77).Cardinality())).Times((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int)), _339_v171)
					_344_v176 = _nw45
					var _345_v177 _dafny.Array
					_ = _345_v177
					var _346_v178 _dafny.Map
					_ = _346_v178
					var _out12 _dafny.Array
					_ = _out12
					var _out13 _dafny.Map
					_ = _out13
					_out12, _out13 = (_343_v175).M1((_344_v176.F23).Minus((_343_v175).Fm3((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int), (_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool), (_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool), _116_globalState)), _116_globalState)
					_345_v177 = _out12
					_346_v178 = _out13
				} else {
					(_116_globalState).F7 = !(_110_v0) || ((_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool))
					var _347_v179 _dafny.Array
					_ = _347_v179
					var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
					_ = _nw46
					_347_v179 = _nw46
					var _348_v180 _dafny.Array
					_ = _348_v180
					var _nw47 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(6))
					_ = _nw47
					_348_v180 = _nw47
					var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_347_v179), 0))
					_ = _index37
					(_347_v179).ArraySet1(_348_v180, (_index37).Int())
					var _349_v181 _dafny.Sequence
					_ = _349_v181
					_349_v181 = _dafny.SeqOf(_348_v180)
					var _350_v182 _dafny.Map
					_ = _350_v182
					_350_v182 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_310_v153).F17(), _297_v142)
					var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_347_v179), 0))
					_ = _index38
					var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))
					_ = _index39
					var _rhs50 _dafny.Array = (_349_v181).Select((Companion_Default___.SafeIndex((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_349_v181).Cardinality()))).Uint32()).(_dafny.Array)
					_ = _rhs50
					var _rhs51 _dafny.Int = (_dafny.Zero).Minus(((_350_v182).Cardinality()).Minus(((_326_v167).F21()).Plus((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int))))
					_ = _rhs51
					var _rhs52 _dafny.Int = _111_v1
					_ = _rhs52
					var _lhs42 _dafny.Array = _347_v179
					_ = _lhs42
					var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_347_v179), 0))
					_ = _lhs43
					var _lhs44 *GlobalState = _116_globalState
					_ = _lhs44
					var _lhs45 _dafny.Array = _296_v141
					_ = _lhs45
					var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))
					_ = _lhs46
					(_lhs42).ArraySet1(_rhs50, (_lhs43).Int())
					_lhs44.F1 = _rhs51
					(_lhs45).ArraySet1(_rhs52, (_lhs46).Int())
					var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))
					_ = _index40
					var _rhs53 _dafny.Int = (_312_v155).Fm9((_326_v167).F21(), Companion_Default___.Fm2(_111_v1, _dafny.IntOfInt64(325), _116_globalState), (_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool), _116_globalState)
					_ = _rhs53
					var _rhs54 _dafny.Array = _296_v141
					_ = _rhs54
					var _rhs55 _dafny.Int = (_326_v167).F21()
					_ = _rhs55
					var _rhs56 bool = (_310_v153).Fm4(Companion_Default___.SafeDivisionInt(_111_v1, _111_v1), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_310_v153).F16(), _113_v3), _111_v1, !((_310_v153).F17()) || ((_310_v153).F17()), _116_globalState)
					_ = _rhs56
					var _lhs47 *GlobalState = _116_globalState
					_ = _lhs47
					var _lhs48 _dafny.Array = _296_v141
					_ = _lhs48
					var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))
					_ = _lhs49
					var _lhs50 *GlobalState = _116_globalState
					_ = _lhs50
					_lhs47.F1 = _rhs53
					_296_v141 = _rhs54
					(_lhs48).ArraySet1(_rhs55, (_lhs49).Int())
					_lhs50.F6 = _rhs56
					(_116_globalState).F1 = (_dafny.Zero).Minus((func() _dafny.Int {
						if (_310_v153).F17() {
							return (func() _dafny.Int {
								if (_310_v153).F17() {
									return (_326_v167).F21()
								}
								return (_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int)
							})()
						}
						return _111_v1
					})())
					var _rhs57 _dafny.Int = (_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int)
					_ = _rhs57
					var _rhs58 _dafny.Int = (func() _dafny.Int {
						if (_114_v4).Contains(_110_v0) {
							return (_114_v4).Get(_110_v0).(_dafny.Int)
						}
						return _111_v1
					})()
					_ = _rhs58
					var _rhs59 bool = (_310_v153).F17()
					_ = _rhs59
					var _rhs60 bool = ((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int)).Cmp((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int)) != 0
					_ = _rhs60
					var _lhs51 *GlobalState = _116_globalState
					_ = _lhs51
					var _lhs52 *GlobalState = _116_globalState
					_ = _lhs52
					var _lhs53 *GlobalState = _116_globalState
					_ = _lhs53
					_lhs51.F1 = _rhs57
					_lhs52.F1 = _rhs58
					_lhs53.F7 = _rhs59
					_110_v0 = _rhs60
				}
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _355_i24 _dafny.Int
	_ = _355_i24
	_355_i24 = _dafny.Zero
	{
		for (_311_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_311_v154), 0))).Int()).(bool) {
			{
				if (_355_i24).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_355_i24 = (_355_i24).Plus(_dafny.One)
				var _356_v183 _dafny.MultiSet
				_ = _356_v183
				var _out14 _dafny.MultiSet
				_ = _out14
				_out14 = (_310_v153).M0(_283_v133, _116_globalState)
				_356_v183 = _out14
				_312_v155 = _312_v155
				_296_v141 = _296_v141
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))
				_ = _index41
				(_296_v141).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_296_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_296_v141), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
					if true {
						return _111_v1
					}
					return _111_v1
				})())))), (_index41).Int())
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	_dafny.Print(_110_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_111_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_112_v2, _dafny.SeqOf(true, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_113_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(754), _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_115_v5).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_115_v5).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_115_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_115_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(754))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_115_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_115_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(754))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_115_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_115_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_116_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_116_globalState.F5).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_116_globalState.F5).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_116_globalState.F5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_116_globalState.F5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(754))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_116_globalState.F5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_116_globalState.F5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(754))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_116_globalState.F5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_116_globalState.F5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_116_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_116_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_117_i0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_198_v77.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_199_v78).Dtor_cf40().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_199_v78).Dtor_cf41())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_199_v78).Dtor_cf42())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_283_v133)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_284_v134).Dtor_cf13(), _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_285_i13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v141).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_297_v142, _dafny.SeqOf(_dafny.IntOfInt64(630))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_298_v143, _dafny.SeqOf(_dafny.IntOfInt64(630), _dafny.IntOfInt64(-630), _dafny.One, _dafny.IntOfInt64(630))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_301_v145).Dtor_cf23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_301_v145).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_302_v146).Dtor_cf25()).Dtor_cf23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_302_v146).Dtor_cf25()).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_303_v147, _dafny.SeqOf(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC16_(true, _dafny.IntOfInt64(630))), Companion_D8_.Create_DC17_(Companion_D8_.Create_DC16_(true, _dafny.IntOfInt64(630))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_304_v148).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC16_(true, _dafny.IntOfInt64(630))), _dafny.IntOfInt64(630)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v149).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC16_(true, _dafny.IntOfInt64(630))), _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_307_v150).Dtor_cf39()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC16_(true, _dafny.IntOfInt64(630))), _dafny.IntOfInt64(630))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v151).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(630), Companion_D13_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC16_(true, _dafny.IntOfInt64(630))), _dafny.IntOfInt64(630)))).UpdateUnsafe(_dafny.IntOfInt64(-630), Companion_D13_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC17_(Companion_D8_.Create_DC16_(true, _dafny.IntOfInt64(630))), _dafny.IntOfInt64(630))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_v152).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(630))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_310_v153).F16()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(630))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v153).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_311_v154).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_313_v156).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_315_v159).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_316_v160).Equals(_dafny.SetOf(_dafny.MultiSetOf(_dafny.One, _dafny.IntOfInt64(630)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_317_v161).Equals(_dafny.MultiSetOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_319_v163).Dtor_cf86()).Equals(_dafny.SetOf(_dafny.IntOfInt64(85))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_321_v164, _dafny.SeqOf(Companion_D22_.Create_DC51_(_dafny.SetOf(_dafny.IntOfInt64(85))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_324_v166).Dtor_cf96()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(630))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_325_i18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_355_i24)
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
	Cf2 _dafny.Int
	Cf3 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 _dafny.Int, Cf3 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Int
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Int) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, _dafny.Zero, false)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
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
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3 == data2.Cf3
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0
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
	Cf4 _dafny.Sequence
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf4 _dafny.Sequence) D1 {
	return D1{D1_DC2{Cf4}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D1) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + data.Cf4.VerbatimString(true) + ")"
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
			return ok && data1.Cf4.Equals(data2.Cf4)
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
	Cf6 _dafny.Int
	Cf7 _dafny.Int
	Cf8 _dafny.Int
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf6 _dafny.Int, Cf7 _dafny.Int, Cf8 _dafny.Int) D2 {
	return D2{D2_DC4{Cf6, Cf7, Cf8}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC5 struct {
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_() D2 {
	return D2{D2_DC5{}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC3 struct {
	Cf5 _dafny.MultiSet
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf5 _dafny.MultiSet) D2 {
	return D2{D2_DC3{Cf5}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC4_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf6
}

func (_this D2) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf7
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf8
}

func (_this D2) Dtor_cf5() _dafny.MultiSet {
	return _this.Get_().(D2_DC3).Cf5
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5"
		}
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8.Cmp(data2.Cf8) == 0
		}
	case D2_DC5:
		{
			_, ok := other.Get_().(D2_DC5)
			return ok
		}
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf5.Equals(data2.Cf5)
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

type D3_DC6 struct {
	Cf9 _dafny.Array
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf9 _dafny.Array) D3 {
	return D3{D3_DC6{Cf9}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D3) Dtor_cf9() _dafny.Array {
	return _this.Get_().(D3_DC6).Cf9
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf9) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC6:
		{
			data2, ok := other.Get_().(D3_DC6)
			return ok && data1.Cf9 == data2.Cf9
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

type D4_DC7 struct {
	Cf10 _dafny.Sequence
}

func (D4_DC7) isD4() {}

func (CompanionStruct_D4_) Create_DC7_(Cf10 _dafny.Sequence) D4 {
	return D4{D4_DC7{Cf10}}
}

func (_this D4) Is_DC7() bool {
	_, ok := _this.Get_().(D4_DC7)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D4) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D4_DC7).Cf10
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC7:
		{
			return "D4.DC7" + "(" + _dafny.String(data.Cf10) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC7:
		{
			data2, ok := other.Get_().(D4_DC7)
			return ok && data1.Cf10.Equals(data2.Cf10)
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

type D5_DC8 struct {
	Cf11 _dafny.MultiSet
}

func (D5_DC8) isD5() {}

func (CompanionStruct_D5_) Create_DC8_(Cf11 _dafny.MultiSet) D5 {
	return D5{D5_DC8{Cf11}}
}

func (_this D5) Is_DC8() bool {
	_, ok := _this.Get_().(D5_DC8)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D5) Dtor_cf11() _dafny.MultiSet {
	return _this.Get_().(D5_DC8).Cf11
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC8:
		{
			return "D5.DC8" + "(" + _dafny.String(data.Cf11) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC8:
		{
			data2, ok := other.Get_().(D5_DC8)
			return ok && data1.Cf11.Equals(data2.Cf11)
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

type D6_DC10 struct {
	Cf13 _dafny.Sequence
}

func (D6_DC10) isD6() {}

func (CompanionStruct_D6_) Create_DC10_(Cf13 _dafny.Sequence) D6 {
	return D6{D6_DC10{Cf13}}
}

func (_this D6) Is_DC10() bool {
	_, ok := _this.Get_().(D6_DC10)
	return ok
}

type D6_DC11 struct {
	Cf14 _dafny.Int
	Cf15 _dafny.Int
	Cf16 _dafny.Int
}

func (D6_DC11) isD6() {}

func (CompanionStruct_D6_) Create_DC11_(Cf14 _dafny.Int, Cf15 _dafny.Int, Cf16 _dafny.Int) D6 {
	return D6{D6_DC11{Cf14, Cf15, Cf16}}
}

func (_this D6) Is_DC11() bool {
	_, ok := _this.Get_().(D6_DC11)
	return ok
}

type D6_DC9 struct {
	Cf12 _dafny.Map
}

func (D6_DC9) isD6() {}

func (CompanionStruct_D6_) Create_DC9_(Cf12 _dafny.Map) D6 {
	return D6{D6_DC9{Cf12}}
}

func (_this D6) Is_DC9() bool {
	_, ok := _this.Get_().(D6_DC9)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC10_(_dafny.EmptySeq)
}

func (_this D6) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D6_DC10).Cf13
}

func (_this D6) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D6_DC11).Cf14
}

func (_this D6) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D6_DC11).Cf15
}

func (_this D6) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D6_DC11).Cf16
}

func (_this D6) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D6_DC9).Cf12
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC10:
		{
			return "D6.DC10" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D6_DC11:
		{
			return "D6.DC11" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D6_DC9:
		{
			return "D6.DC9" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC10:
		{
			data2, ok := other.Get_().(D6_DC10)
			return ok && data1.Cf13.Equals(data2.Cf13)
		}
	case D6_DC11:
		{
			data2, ok := other.Get_().(D6_DC11)
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Cmp(data2.Cf16) == 0
		}
	case D6_DC9:
		{
			data2, ok := other.Get_().(D6_DC9)
			return ok && data1.Cf12.Equals(data2.Cf12)
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

type D7_DC13 struct {
	Cf18 _dafny.Set
	Cf19 _dafny.Int
	Cf20 _dafny.Int
}

func (D7_DC13) isD7() {}

func (CompanionStruct_D7_) Create_DC13_(Cf18 _dafny.Set, Cf19 _dafny.Int, Cf20 _dafny.Int) D7 {
	return D7{D7_DC13{Cf18, Cf19, Cf20}}
}

func (_this D7) Is_DC13() bool {
	_, ok := _this.Get_().(D7_DC13)
	return ok
}

type D7_DC12 struct {
	Cf17 _dafny.Sequence
}

func (D7_DC12) isD7() {}

func (CompanionStruct_D7_) Create_DC12_(Cf17 _dafny.Sequence) D7 {
	return D7{D7_DC12{Cf17}}
}

func (_this D7) Is_DC12() bool {
	_, ok := _this.Get_().(D7_DC12)
	return ok
}

type D7_DC14 struct {
	Cf21 D7
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf21 D7) D7 {
	return D7{D7_DC14{Cf21}}
}

func (_this D7) Is_DC14() bool {
	_, ok := _this.Get_().(D7_DC14)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC13_(_dafny.EmptySet, _dafny.Zero, _dafny.Zero)
}

func (_this D7) Dtor_cf18() _dafny.Set {
	return _this.Get_().(D7_DC13).Cf18
}

func (_this D7) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D7_DC13).Cf19
}

func (_this D7) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D7_DC13).Cf20
}

func (_this D7) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D7_DC12).Cf17
}

func (_this D7) Dtor_cf21() D7 {
	return _this.Get_().(D7_DC14).Cf21
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC13:
		{
			return "D7.DC13" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D7_DC12:
		{
			return "D7.DC12" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf21) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC13:
		{
			data2, ok := other.Get_().(D7_DC13)
			return ok && data1.Cf18.Equals(data2.Cf18) && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D7_DC12:
		{
			data2, ok := other.Get_().(D7_DC12)
			return ok && data1.Cf17.Equals(data2.Cf17)
		}
	case D7_DC14:
		{
			data2, ok := other.Get_().(D7_DC14)
			return ok && data1.Cf21.Equals(data2.Cf21)
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

type D8_DC16 struct {
	Cf23 bool
	Cf24 _dafny.Int
}

func (D8_DC16) isD8() {}

func (CompanionStruct_D8_) Create_DC16_(Cf23 bool, Cf24 _dafny.Int) D8 {
	return D8{D8_DC16{Cf23, Cf24}}
}

func (_this D8) Is_DC16() bool {
	_, ok := _this.Get_().(D8_DC16)
	return ok
}

type D8_DC15 struct {
	Cf22 _dafny.Map
}

func (D8_DC15) isD8() {}

func (CompanionStruct_D8_) Create_DC15_(Cf22 _dafny.Map) D8 {
	return D8{D8_DC15{Cf22}}
}

func (_this D8) Is_DC15() bool {
	_, ok := _this.Get_().(D8_DC15)
	return ok
}

type D8_DC17 struct {
	Cf25 D8
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf25 D8) D8 {
	return D8{D8_DC17{Cf25}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC16_(false, _dafny.Zero)
}

func (_this D8) Dtor_cf23() bool {
	return _this.Get_().(D8_DC16).Cf23
}

func (_this D8) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D8_DC16).Cf24
}

func (_this D8) Dtor_cf22() _dafny.Map {
	return _this.Get_().(D8_DC15).Cf22
}

func (_this D8) Dtor_cf25() D8 {
	return _this.Get_().(D8_DC17).Cf25
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC16:
		{
			return "D8.DC16" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D8_DC15:
		{
			return "D8.DC15" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D8_DC17:
		{
			return "D8.DC17" + "(" + _dafny.String(data.Cf25) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC16:
		{
			data2, ok := other.Get_().(D8_DC16)
			return ok && data1.Cf23 == data2.Cf23 && data1.Cf24.Cmp(data2.Cf24) == 0
		}
	case D8_DC15:
		{
			data2, ok := other.Get_().(D8_DC15)
			return ok && data1.Cf22.Equals(data2.Cf22)
		}
	case D8_DC17:
		{
			data2, ok := other.Get_().(D8_DC17)
			return ok && data1.Cf25.Equals(data2.Cf25)
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

type D9_DC19 struct {
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_() D9 {
	return D9{D9_DC19{}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

type D9_DC18 struct {
	Cf26 _dafny.Map
}

func (D9_DC18) isD9() {}

func (CompanionStruct_D9_) Create_DC18_(Cf26 _dafny.Map) D9 {
	return D9{D9_DC18{Cf26}}
}

func (_this D9) Is_DC18() bool {
	_, ok := _this.Get_().(D9_DC18)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC19_()
}

func (_this D9) Dtor_cf26() _dafny.Map {
	return _this.Get_().(D9_DC18).Cf26
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC19:
		{
			return "D9.DC19"
		}
	case D9_DC18:
		{
			return "D9.DC18" + "(" + _dafny.String(data.Cf26) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC19:
		{
			_, ok := other.Get_().(D9_DC19)
			return ok
		}
	case D9_DC18:
		{
			data2, ok := other.Get_().(D9_DC18)
			return ok && data1.Cf26.Equals(data2.Cf26)
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

type D10_DC21 struct {
	Cf28 bool
	Cf29 bool
}

func (D10_DC21) isD10() {}

func (CompanionStruct_D10_) Create_DC21_(Cf28 bool, Cf29 bool) D10 {
	return D10{D10_DC21{Cf28, Cf29}}
}

func (_this D10) Is_DC21() bool {
	_, ok := _this.Get_().(D10_DC21)
	return ok
}

type D10_DC20 struct {
	Cf27 _dafny.Map
}

func (D10_DC20) isD10() {}

func (CompanionStruct_D10_) Create_DC20_(Cf27 _dafny.Map) D10 {
	return D10{D10_DC20{Cf27}}
}

func (_this D10) Is_DC20() bool {
	_, ok := _this.Get_().(D10_DC20)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC21_(false, false)
}

func (_this D10) Dtor_cf28() bool {
	return _this.Get_().(D10_DC21).Cf28
}

func (_this D10) Dtor_cf29() bool {
	return _this.Get_().(D10_DC21).Cf29
}

func (_this D10) Dtor_cf27() _dafny.Map {
	return _this.Get_().(D10_DC20).Cf27
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC21:
		{
			return "D10.DC21" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D10_DC20:
		{
			return "D10.DC20" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC21:
		{
			data2, ok := other.Get_().(D10_DC21)
			return ok && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29
		}
	case D10_DC20:
		{
			data2, ok := other.Get_().(D10_DC20)
			return ok && data1.Cf27.Equals(data2.Cf27)
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

type D11_DC23 struct {
	Cf31 bool
	Cf32 _dafny.Int
	Cf33 _dafny.Int
}

func (D11_DC23) isD11() {}

func (CompanionStruct_D11_) Create_DC23_(Cf31 bool, Cf32 _dafny.Int, Cf33 _dafny.Int) D11 {
	return D11{D11_DC23{Cf31, Cf32, Cf33}}
}

func (_this D11) Is_DC23() bool {
	_, ok := _this.Get_().(D11_DC23)
	return ok
}

type D11_DC24 struct {
	Cf34 _dafny.Int
	Cf35 _dafny.Sequence
	Cf36 _dafny.CodePoint
}

func (D11_DC24) isD11() {}

func (CompanionStruct_D11_) Create_DC24_(Cf34 _dafny.Int, Cf35 _dafny.Sequence, Cf36 _dafny.CodePoint) D11 {
	return D11{D11_DC24{Cf34, Cf35, Cf36}}
}

func (_this D11) Is_DC24() bool {
	_, ok := _this.Get_().(D11_DC24)
	return ok
}

type D11_DC22 struct {
	Cf30 _dafny.Map
}

func (D11_DC22) isD11() {}

func (CompanionStruct_D11_) Create_DC22_(Cf30 _dafny.Map) D11 {
	return D11{D11_DC22{Cf30}}
}

func (_this D11) Is_DC22() bool {
	_, ok := _this.Get_().(D11_DC22)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC23_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D11) Dtor_cf31() bool {
	return _this.Get_().(D11_DC23).Cf31
}

func (_this D11) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D11_DC23).Cf32
}

func (_this D11) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D11_DC23).Cf33
}

func (_this D11) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D11_DC24).Cf34
}

func (_this D11) Dtor_cf35() _dafny.Sequence {
	return _this.Get_().(D11_DC24).Cf35
}

func (_this D11) Dtor_cf36() _dafny.CodePoint {
	return _this.Get_().(D11_DC24).Cf36
}

func (_this D11) Dtor_cf30() _dafny.Map {
	return _this.Get_().(D11_DC22).Cf30
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC23:
		{
			return "D11.DC23" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D11_DC24:
		{
			return "D11.DC24" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D11_DC22:
		{
			return "D11.DC22" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC23:
		{
			data2, ok := other.Get_().(D11_DC23)
			return ok && data1.Cf31 == data2.Cf31 && data1.Cf32.Cmp(data2.Cf32) == 0 && data1.Cf33.Cmp(data2.Cf33) == 0
		}
	case D11_DC24:
		{
			data2, ok := other.Get_().(D11_DC24)
			return ok && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35.Equals(data2.Cf35) && data1.Cf36 == data2.Cf36
		}
	case D11_DC22:
		{
			data2, ok := other.Get_().(D11_DC22)
			return ok && data1.Cf30.Equals(data2.Cf30)
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

type D12_DC26 struct {
	Cf38 _dafny.Map
}

func (D12_DC26) isD12() {}

func (CompanionStruct_D12_) Create_DC26_(Cf38 _dafny.Map) D12 {
	return D12{D12_DC26{Cf38}}
}

func (_this D12) Is_DC26() bool {
	_, ok := _this.Get_().(D12_DC26)
	return ok
}

type D12_DC25 struct {
	Cf37 _dafny.Array
}

func (D12_DC25) isD12() {}

func (CompanionStruct_D12_) Create_DC25_(Cf37 _dafny.Array) D12 {
	return D12{D12_DC25{Cf37}}
}

func (_this D12) Is_DC25() bool {
	_, ok := _this.Get_().(D12_DC25)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC26_(_dafny.EmptyMap)
}

func (_this D12) Dtor_cf38() _dafny.Map {
	return _this.Get_().(D12_DC26).Cf38
}

func (_this D12) Dtor_cf37() _dafny.Array {
	return _this.Get_().(D12_DC25).Cf37
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC26:
		{
			return "D12.DC26" + "(" + _dafny.String(data.Cf38) + ")"
		}
	case D12_DC25:
		{
			return "D12.DC25" + "(" + _dafny.String(data.Cf37) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC26:
		{
			data2, ok := other.Get_().(D12_DC26)
			return ok && data1.Cf38.Equals(data2.Cf38)
		}
	case D12_DC25:
		{
			data2, ok := other.Get_().(D12_DC25)
			return ok && data1.Cf37 == data2.Cf37
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

type D13_DC28 struct {
	Cf40 _dafny.Sequence
	Cf41 _dafny.Int
	Cf42 bool
}

func (D13_DC28) isD13() {}

func (CompanionStruct_D13_) Create_DC28_(Cf40 _dafny.Sequence, Cf41 _dafny.Int, Cf42 bool) D13 {
	return D13{D13_DC28{Cf40, Cf41, Cf42}}
}

func (_this D13) Is_DC28() bool {
	_, ok := _this.Get_().(D13_DC28)
	return ok
}

type D13_DC27 struct {
	Cf39 _dafny.Map
}

func (D13_DC27) isD13() {}

func (CompanionStruct_D13_) Create_DC27_(Cf39 _dafny.Map) D13 {
	return D13{D13_DC27{Cf39}}
}

func (_this D13) Is_DC27() bool {
	_, ok := _this.Get_().(D13_DC27)
	return ok
}

type D13_DC29 struct {
	Cf43 D13
}

func (D13_DC29) isD13() {}

func (CompanionStruct_D13_) Create_DC29_(Cf43 D13) D13 {
	return D13{D13_DC29{Cf43}}
}

func (_this D13) Is_DC29() bool {
	_, ok := _this.Get_().(D13_DC29)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC28_(_dafny.EmptySeq, _dafny.Zero, false)
}

func (_this D13) Dtor_cf40() _dafny.Sequence {
	return _this.Get_().(D13_DC28).Cf40
}

func (_this D13) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D13_DC28).Cf41
}

func (_this D13) Dtor_cf42() bool {
	return _this.Get_().(D13_DC28).Cf42
}

func (_this D13) Dtor_cf39() _dafny.Map {
	return _this.Get_().(D13_DC27).Cf39
}

func (_this D13) Dtor_cf43() D13 {
	return _this.Get_().(D13_DC29).Cf43
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC28:
		{
			return "D13.DC28" + "(" + data.Cf40.VerbatimString(true) + ", " + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ")"
		}
	case D13_DC27:
		{
			return "D13.DC27" + "(" + _dafny.String(data.Cf39) + ")"
		}
	case D13_DC29:
		{
			return "D13.DC29" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC28:
		{
			data2, ok := other.Get_().(D13_DC28)
			return ok && data1.Cf40.Equals(data2.Cf40) && data1.Cf41.Cmp(data2.Cf41) == 0 && data1.Cf42 == data2.Cf42
		}
	case D13_DC27:
		{
			data2, ok := other.Get_().(D13_DC27)
			return ok && data1.Cf39.Equals(data2.Cf39)
		}
	case D13_DC29:
		{
			data2, ok := other.Get_().(D13_DC29)
			return ok && data1.Cf43.Equals(data2.Cf43)
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

type D14_DC31 struct {
	Cf45 bool
	Cf46 _dafny.Int
	Cf47 _dafny.Int
}

func (D14_DC31) isD14() {}

func (CompanionStruct_D14_) Create_DC31_(Cf45 bool, Cf46 _dafny.Int, Cf47 _dafny.Int) D14 {
	return D14{D14_DC31{Cf45, Cf46, Cf47}}
}

func (_this D14) Is_DC31() bool {
	_, ok := _this.Get_().(D14_DC31)
	return ok
}

type D14_DC30 struct {
	Cf44 _dafny.Map
}

func (D14_DC30) isD14() {}

func (CompanionStruct_D14_) Create_DC30_(Cf44 _dafny.Map) D14 {
	return D14{D14_DC30{Cf44}}
}

func (_this D14) Is_DC30() bool {
	_, ok := _this.Get_().(D14_DC30)
	return ok
}

type D14_DC32 struct {
	Cf48 D14
}

func (D14_DC32) isD14() {}

func (CompanionStruct_D14_) Create_DC32_(Cf48 D14) D14 {
	return D14{D14_DC32{Cf48}}
}

func (_this D14) Is_DC32() bool {
	_, ok := _this.Get_().(D14_DC32)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC31_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D14) Dtor_cf45() bool {
	return _this.Get_().(D14_DC31).Cf45
}

func (_this D14) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D14_DC31).Cf46
}

func (_this D14) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D14_DC31).Cf47
}

func (_this D14) Dtor_cf44() _dafny.Map {
	return _this.Get_().(D14_DC30).Cf44
}

func (_this D14) Dtor_cf48() D14 {
	return _this.Get_().(D14_DC32).Cf48
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC31:
		{
			return "D14.DC31" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ")"
		}
	case D14_DC30:
		{
			return "D14.DC30" + "(" + _dafny.String(data.Cf44) + ")"
		}
	case D14_DC32:
		{
			return "D14.DC32" + "(" + _dafny.String(data.Cf48) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC31:
		{
			data2, ok := other.Get_().(D14_DC31)
			return ok && data1.Cf45 == data2.Cf45 && data1.Cf46.Cmp(data2.Cf46) == 0 && data1.Cf47.Cmp(data2.Cf47) == 0
		}
	case D14_DC30:
		{
			data2, ok := other.Get_().(D14_DC30)
			return ok && data1.Cf44.Equals(data2.Cf44)
		}
	case D14_DC32:
		{
			data2, ok := other.Get_().(D14_DC32)
			return ok && data1.Cf48.Equals(data2.Cf48)
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

type D15_DC34 struct {
}

func (D15_DC34) isD15() {}

func (CompanionStruct_D15_) Create_DC34_() D15 {
	return D15{D15_DC34{}}
}

func (_this D15) Is_DC34() bool {
	_, ok := _this.Get_().(D15_DC34)
	return ok
}

type D15_DC35 struct {
	Cf50 bool
	Cf51 T1
	Cf52 bool
	Cf53 _dafny.Sequence
	Cf54 _dafny.Int
}

func (D15_DC35) isD15() {}

func (CompanionStruct_D15_) Create_DC35_(Cf50 bool, Cf51 T1, Cf52 bool, Cf53 _dafny.Sequence, Cf54 _dafny.Int) D15 {
	return D15{D15_DC35{Cf50, Cf51, Cf52, Cf53, Cf54}}
}

func (_this D15) Is_DC35() bool {
	_, ok := _this.Get_().(D15_DC35)
	return ok
}

type D15_DC36 struct {
	Cf55 _dafny.Int
	Cf56 bool
	Cf57 bool
	Cf58 bool
	Cf59 _dafny.Int
}

func (D15_DC36) isD15() {}

func (CompanionStruct_D15_) Create_DC36_(Cf55 _dafny.Int, Cf56 bool, Cf57 bool, Cf58 bool, Cf59 _dafny.Int) D15 {
	return D15{D15_DC36{Cf55, Cf56, Cf57, Cf58, Cf59}}
}

func (_this D15) Is_DC36() bool {
	_, ok := _this.Get_().(D15_DC36)
	return ok
}

type D15_DC33 struct {
	Cf49 _dafny.Map
}

func (D15_DC33) isD15() {}

func (CompanionStruct_D15_) Create_DC33_(Cf49 _dafny.Map) D15 {
	return D15{D15_DC33{Cf49}}
}

func (_this D15) Is_DC33() bool {
	_, ok := _this.Get_().(D15_DC33)
	return ok
}

type D15_DC37 struct {
	Cf60 D15
}

func (D15_DC37) isD15() {}

func (CompanionStruct_D15_) Create_DC37_(Cf60 D15) D15 {
	return D15{D15_DC37{Cf60}}
}

func (_this D15) Is_DC37() bool {
	_, ok := _this.Get_().(D15_DC37)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC34_()
}

func (_this D15) Dtor_cf50() bool {
	return _this.Get_().(D15_DC35).Cf50
}

func (_this D15) Dtor_cf51() T1 {
	return _this.Get_().(D15_DC35).Cf51
}

func (_this D15) Dtor_cf52() bool {
	return _this.Get_().(D15_DC35).Cf52
}

func (_this D15) Dtor_cf53() _dafny.Sequence {
	return _this.Get_().(D15_DC35).Cf53
}

func (_this D15) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D15_DC35).Cf54
}

func (_this D15) Dtor_cf55() _dafny.Int {
	return _this.Get_().(D15_DC36).Cf55
}

func (_this D15) Dtor_cf56() bool {
	return _this.Get_().(D15_DC36).Cf56
}

func (_this D15) Dtor_cf57() bool {
	return _this.Get_().(D15_DC36).Cf57
}

func (_this D15) Dtor_cf58() bool {
	return _this.Get_().(D15_DC36).Cf58
}

func (_this D15) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D15_DC36).Cf59
}

func (_this D15) Dtor_cf49() _dafny.Map {
	return _this.Get_().(D15_DC33).Cf49
}

func (_this D15) Dtor_cf60() D15 {
	return _this.Get_().(D15_DC37).Cf60
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC34:
		{
			return "D15.DC34"
		}
	case D15_DC35:
		{
			return "D15.DC35" + "(" + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D15_DC36:
		{
			return "D15.DC36" + "(" + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D15_DC33:
		{
			return "D15.DC33" + "(" + _dafny.String(data.Cf49) + ")"
		}
	case D15_DC37:
		{
			return "D15.DC37" + "(" + _dafny.String(data.Cf60) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC34:
		{
			_, ok := other.Get_().(D15_DC34)
			return ok
		}
	case D15_DC35:
		{
			data2, ok := other.Get_().(D15_DC35)
			return ok && data1.Cf50 == data2.Cf50 && _dafny.AreEqual(data1.Cf51, data2.Cf51) && data1.Cf52 == data2.Cf52 && data1.Cf53.Equals(data2.Cf53) && data1.Cf54.Cmp(data2.Cf54) == 0
		}
	case D15_DC36:
		{
			data2, ok := other.Get_().(D15_DC36)
			return ok && data1.Cf55.Cmp(data2.Cf55) == 0 && data1.Cf56 == data2.Cf56 && data1.Cf57 == data2.Cf57 && data1.Cf58 == data2.Cf58 && data1.Cf59.Cmp(data2.Cf59) == 0
		}
	case D15_DC33:
		{
			data2, ok := other.Get_().(D15_DC33)
			return ok && data1.Cf49.Equals(data2.Cf49)
		}
	case D15_DC37:
		{
			data2, ok := other.Get_().(D15_DC37)
			return ok && data1.Cf60.Equals(data2.Cf60)
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

type D16_DC39 struct {
	Cf62 _dafny.Int
	Cf63 bool
	Cf64 _dafny.Map
	Cf65 _dafny.CodePoint
}

func (D16_DC39) isD16() {}

func (CompanionStruct_D16_) Create_DC39_(Cf62 _dafny.Int, Cf63 bool, Cf64 _dafny.Map, Cf65 _dafny.CodePoint) D16 {
	return D16{D16_DC39{Cf62, Cf63, Cf64, Cf65}}
}

func (_this D16) Is_DC39() bool {
	_, ok := _this.Get_().(D16_DC39)
	return ok
}

type D16_DC38 struct {
	Cf61 _dafny.Map
}

func (D16_DC38) isD16() {}

func (CompanionStruct_D16_) Create_DC38_(Cf61 _dafny.Map) D16 {
	return D16{D16_DC38{Cf61}}
}

func (_this D16) Is_DC38() bool {
	_, ok := _this.Get_().(D16_DC38)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC39_(_dafny.Zero, false, _dafny.EmptyMap, _dafny.CodePoint('D'))
}

func (_this D16) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D16_DC39).Cf62
}

func (_this D16) Dtor_cf63() bool {
	return _this.Get_().(D16_DC39).Cf63
}

func (_this D16) Dtor_cf64() _dafny.Map {
	return _this.Get_().(D16_DC39).Cf64
}

func (_this D16) Dtor_cf65() _dafny.CodePoint {
	return _this.Get_().(D16_DC39).Cf65
}

func (_this D16) Dtor_cf61() _dafny.Map {
	return _this.Get_().(D16_DC38).Cf61
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC39:
		{
			return "D16.DC39" + "(" + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ")"
		}
	case D16_DC38:
		{
			return "D16.DC38" + "(" + _dafny.String(data.Cf61) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC39:
		{
			data2, ok := other.Get_().(D16_DC39)
			return ok && data1.Cf62.Cmp(data2.Cf62) == 0 && data1.Cf63 == data2.Cf63 && data1.Cf64.Equals(data2.Cf64) && data1.Cf65 == data2.Cf65
		}
	case D16_DC38:
		{
			data2, ok := other.Get_().(D16_DC38)
			return ok && data1.Cf61.Equals(data2.Cf61)
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

type D17_DC40 struct {
	Cf66 _dafny.Array
}

func (D17_DC40) isD17() {}

func (CompanionStruct_D17_) Create_DC40_(Cf66 _dafny.Array) D17 {
	return D17{D17_DC40{Cf66}}
}

func (_this D17) Is_DC40() bool {
	_, ok := _this.Get_().(D17_DC40)
	return ok
}

func (CompanionStruct_D17_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D17) Dtor_cf66() _dafny.Array {
	return _this.Get_().(D17_DC40).Cf66
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC40:
		{
			return "D17.DC40" + "(" + _dafny.String(data.Cf66) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC40:
		{
			data2, ok := other.Get_().(D17_DC40)
			return ok && data1.Cf66 == data2.Cf66
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

type D18_DC41 struct {
	Cf67 _dafny.Map
}

func (D18_DC41) isD18() {}

func (CompanionStruct_D18_) Create_DC41_(Cf67 _dafny.Map) D18 {
	return D18{D18_DC41{Cf67}}
}

func (_this D18) Is_DC41() bool {
	_, ok := _this.Get_().(D18_DC41)
	return ok
}

func (CompanionStruct_D18_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D18) Dtor_cf67() _dafny.Map {
	return _this.Get_().(D18_DC41).Cf67
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC41:
		{
			return "D18.DC41" + "(" + _dafny.String(data.Cf67) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC41:
		{
			data2, ok := other.Get_().(D18_DC41)
			return ok && data1.Cf67.Equals(data2.Cf67)
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

type D19_DC43 struct {
	Cf69 bool
	Cf70 _dafny.CodePoint
	Cf71 bool
	Cf72 D6
	Cf73 _dafny.Sequence
}

func (D19_DC43) isD19() {}

func (CompanionStruct_D19_) Create_DC43_(Cf69 bool, Cf70 _dafny.CodePoint, Cf71 bool, Cf72 D6, Cf73 _dafny.Sequence) D19 {
	return D19{D19_DC43{Cf69, Cf70, Cf71, Cf72, Cf73}}
}

func (_this D19) Is_DC43() bool {
	_, ok := _this.Get_().(D19_DC43)
	return ok
}

type D19_DC42 struct {
	Cf68 _dafny.Array
}

func (D19_DC42) isD19() {}

func (CompanionStruct_D19_) Create_DC42_(Cf68 _dafny.Array) D19 {
	return D19{D19_DC42{Cf68}}
}

func (_this D19) Is_DC42() bool {
	_, ok := _this.Get_().(D19_DC42)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC43_(false, _dafny.CodePoint('D'), false, Companion_D6_.Default(), _dafny.EmptySeq)
}

func (_this D19) Dtor_cf69() bool {
	return _this.Get_().(D19_DC43).Cf69
}

func (_this D19) Dtor_cf70() _dafny.CodePoint {
	return _this.Get_().(D19_DC43).Cf70
}

func (_this D19) Dtor_cf71() bool {
	return _this.Get_().(D19_DC43).Cf71
}

func (_this D19) Dtor_cf72() D6 {
	return _this.Get_().(D19_DC43).Cf72
}

func (_this D19) Dtor_cf73() _dafny.Sequence {
	return _this.Get_().(D19_DC43).Cf73
}

func (_this D19) Dtor_cf68() _dafny.Array {
	return _this.Get_().(D19_DC42).Cf68
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC43:
		{
			return "D19.DC43" + "(" + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ", " + data.Cf73.VerbatimString(true) + ")"
		}
	case D19_DC42:
		{
			return "D19.DC42" + "(" + _dafny.String(data.Cf68) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC43:
		{
			data2, ok := other.Get_().(D19_DC43)
			return ok && data1.Cf69 == data2.Cf69 && data1.Cf70 == data2.Cf70 && data1.Cf71 == data2.Cf71 && data1.Cf72.Equals(data2.Cf72) && data1.Cf73.Equals(data2.Cf73)
		}
	case D19_DC42:
		{
			data2, ok := other.Get_().(D19_DC42)
			return ok && data1.Cf68 == data2.Cf68
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

type D20_DC45 struct {
	Cf75 bool
	Cf76 bool
}

func (D20_DC45) isD20() {}

func (CompanionStruct_D20_) Create_DC45_(Cf75 bool, Cf76 bool) D20 {
	return D20{D20_DC45{Cf75, Cf76}}
}

func (_this D20) Is_DC45() bool {
	_, ok := _this.Get_().(D20_DC45)
	return ok
}

type D20_DC46 struct {
}

func (D20_DC46) isD20() {}

func (CompanionStruct_D20_) Create_DC46_() D20 {
	return D20{D20_DC46{}}
}

func (_this D20) Is_DC46() bool {
	_, ok := _this.Get_().(D20_DC46)
	return ok
}

type D20_DC44 struct {
	Cf74 *C5
}

func (D20_DC44) isD20() {}

func (CompanionStruct_D20_) Create_DC44_(Cf74 *C5) D20 {
	return D20{D20_DC44{Cf74}}
}

func (_this D20) Is_DC44() bool {
	_, ok := _this.Get_().(D20_DC44)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC45_(false, false)
}

func (_this D20) Dtor_cf75() bool {
	return _this.Get_().(D20_DC45).Cf75
}

func (_this D20) Dtor_cf76() bool {
	return _this.Get_().(D20_DC45).Cf76
}

func (_this D20) Dtor_cf74() *C5 {
	return _this.Get_().(D20_DC44).Cf74
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC45:
		{
			return "D20.DC45" + "(" + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ")"
		}
	case D20_DC46:
		{
			return "D20.DC46"
		}
	case D20_DC44:
		{
			return "D20.DC44" + "(" + _dafny.String(data.Cf74) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC45:
		{
			data2, ok := other.Get_().(D20_DC45)
			return ok && data1.Cf75 == data2.Cf75 && data1.Cf76 == data2.Cf76
		}
	case D20_DC46:
		{
			_, ok := other.Get_().(D20_DC46)
			return ok
		}
	case D20_DC44:
		{
			data2, ok := other.Get_().(D20_DC44)
			return ok && data1.Cf74 == data2.Cf74
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

type D21_DC48 struct {
	Cf78 bool
}

func (D21_DC48) isD21() {}

func (CompanionStruct_D21_) Create_DC48_(Cf78 bool) D21 {
	return D21{D21_DC48{Cf78}}
}

func (_this D21) Is_DC48() bool {
	_, ok := _this.Get_().(D21_DC48)
	return ok
}

type D21_DC49 struct {
	Cf79 bool
	Cf80 _dafny.Array
	Cf81 _dafny.Int
	Cf82 _dafny.Int
}

func (D21_DC49) isD21() {}

func (CompanionStruct_D21_) Create_DC49_(Cf79 bool, Cf80 _dafny.Array, Cf81 _dafny.Int, Cf82 _dafny.Int) D21 {
	return D21{D21_DC49{Cf79, Cf80, Cf81, Cf82}}
}

func (_this D21) Is_DC49() bool {
	_, ok := _this.Get_().(D21_DC49)
	return ok
}

type D21_DC50 struct {
	Cf83 _dafny.Array
	Cf84 _dafny.Int
	Cf85 bool
}

func (D21_DC50) isD21() {}

func (CompanionStruct_D21_) Create_DC50_(Cf83 _dafny.Array, Cf84 _dafny.Int, Cf85 bool) D21 {
	return D21{D21_DC50{Cf83, Cf84, Cf85}}
}

func (_this D21) Is_DC50() bool {
	_, ok := _this.Get_().(D21_DC50)
	return ok
}

type D21_DC47 struct {
	Cf77 _dafny.Set
}

func (D21_DC47) isD21() {}

func (CompanionStruct_D21_) Create_DC47_(Cf77 _dafny.Set) D21 {
	return D21{D21_DC47{Cf77}}
}

func (_this D21) Is_DC47() bool {
	_, ok := _this.Get_().(D21_DC47)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC48_(false)
}

func (_this D21) Dtor_cf78() bool {
	return _this.Get_().(D21_DC48).Cf78
}

func (_this D21) Dtor_cf79() bool {
	return _this.Get_().(D21_DC49).Cf79
}

func (_this D21) Dtor_cf80() _dafny.Array {
	return _this.Get_().(D21_DC49).Cf80
}

func (_this D21) Dtor_cf81() _dafny.Int {
	return _this.Get_().(D21_DC49).Cf81
}

func (_this D21) Dtor_cf82() _dafny.Int {
	return _this.Get_().(D21_DC49).Cf82
}

func (_this D21) Dtor_cf83() _dafny.Array {
	return _this.Get_().(D21_DC50).Cf83
}

func (_this D21) Dtor_cf84() _dafny.Int {
	return _this.Get_().(D21_DC50).Cf84
}

func (_this D21) Dtor_cf85() bool {
	return _this.Get_().(D21_DC50).Cf85
}

func (_this D21) Dtor_cf77() _dafny.Set {
	return _this.Get_().(D21_DC47).Cf77
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC48:
		{
			return "D21.DC48" + "(" + _dafny.String(data.Cf78) + ")"
		}
	case D21_DC49:
		{
			return "D21.DC49" + "(" + _dafny.String(data.Cf79) + ", " + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ", " + _dafny.String(data.Cf82) + ")"
		}
	case D21_DC50:
		{
			return "D21.DC50" + "(" + _dafny.String(data.Cf83) + ", " + _dafny.String(data.Cf84) + ", " + _dafny.String(data.Cf85) + ")"
		}
	case D21_DC47:
		{
			return "D21.DC47" + "(" + _dafny.String(data.Cf77) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC48:
		{
			data2, ok := other.Get_().(D21_DC48)
			return ok && data1.Cf78 == data2.Cf78
		}
	case D21_DC49:
		{
			data2, ok := other.Get_().(D21_DC49)
			return ok && data1.Cf79 == data2.Cf79 && data1.Cf80 == data2.Cf80 && data1.Cf81.Cmp(data2.Cf81) == 0 && data1.Cf82.Cmp(data2.Cf82) == 0
		}
	case D21_DC50:
		{
			data2, ok := other.Get_().(D21_DC50)
			return ok && data1.Cf83 == data2.Cf83 && data1.Cf84.Cmp(data2.Cf84) == 0 && data1.Cf85 == data2.Cf85
		}
	case D21_DC47:
		{
			data2, ok := other.Get_().(D21_DC47)
			return ok && data1.Cf77.Equals(data2.Cf77)
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

type D22_DC52 struct {
}

func (D22_DC52) isD22() {}

func (CompanionStruct_D22_) Create_DC52_() D22 {
	return D22{D22_DC52{}}
}

func (_this D22) Is_DC52() bool {
	_, ok := _this.Get_().(D22_DC52)
	return ok
}

type D22_DC53 struct {
	Cf87 _dafny.Int
	Cf88 bool
	Cf89 _dafny.Sequence
	Cf90 _dafny.Int
}

func (D22_DC53) isD22() {}

func (CompanionStruct_D22_) Create_DC53_(Cf87 _dafny.Int, Cf88 bool, Cf89 _dafny.Sequence, Cf90 _dafny.Int) D22 {
	return D22{D22_DC53{Cf87, Cf88, Cf89, Cf90}}
}

func (_this D22) Is_DC53() bool {
	_, ok := _this.Get_().(D22_DC53)
	return ok
}

type D22_DC51 struct {
	Cf86 _dafny.Set
}

func (D22_DC51) isD22() {}

func (CompanionStruct_D22_) Create_DC51_(Cf86 _dafny.Set) D22 {
	return D22{D22_DC51{Cf86}}
}

func (_this D22) Is_DC51() bool {
	_, ok := _this.Get_().(D22_DC51)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC52_()
}

func (_this D22) Dtor_cf87() _dafny.Int {
	return _this.Get_().(D22_DC53).Cf87
}

func (_this D22) Dtor_cf88() bool {
	return _this.Get_().(D22_DC53).Cf88
}

func (_this D22) Dtor_cf89() _dafny.Sequence {
	return _this.Get_().(D22_DC53).Cf89
}

func (_this D22) Dtor_cf90() _dafny.Int {
	return _this.Get_().(D22_DC53).Cf90
}

func (_this D22) Dtor_cf86() _dafny.Set {
	return _this.Get_().(D22_DC51).Cf86
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC52:
		{
			return "D22.DC52"
		}
	case D22_DC53:
		{
			return "D22.DC53" + "(" + _dafny.String(data.Cf87) + ", " + _dafny.String(data.Cf88) + ", " + data.Cf89.VerbatimString(true) + ", " + _dafny.String(data.Cf90) + ")"
		}
	case D22_DC51:
		{
			return "D22.DC51" + "(" + _dafny.String(data.Cf86) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC52:
		{
			_, ok := other.Get_().(D22_DC52)
			return ok
		}
	case D22_DC53:
		{
			data2, ok := other.Get_().(D22_DC53)
			return ok && data1.Cf87.Cmp(data2.Cf87) == 0 && data1.Cf88 == data2.Cf88 && data1.Cf89.Equals(data2.Cf89) && data1.Cf90.Cmp(data2.Cf90) == 0
		}
	case D22_DC51:
		{
			data2, ok := other.Get_().(D22_DC51)
			return ok && data1.Cf86.Equals(data2.Cf86)
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

type D23_DC55 struct {
	Cf92 _dafny.Sequence
	Cf93 _dafny.Array
	Cf94 _dafny.Int
}

func (D23_DC55) isD23() {}

func (CompanionStruct_D23_) Create_DC55_(Cf92 _dafny.Sequence, Cf93 _dafny.Array, Cf94 _dafny.Int) D23 {
	return D23{D23_DC55{Cf92, Cf93, Cf94}}
}

func (_this D23) Is_DC55() bool {
	_, ok := _this.Get_().(D23_DC55)
	return ok
}

type D23_DC56 struct {
}

func (D23_DC56) isD23() {}

func (CompanionStruct_D23_) Create_DC56_() D23 {
	return D23{D23_DC56{}}
}

func (_this D23) Is_DC56() bool {
	_, ok := _this.Get_().(D23_DC56)
	return ok
}

type D23_DC54 struct {
	Cf91 *C0
}

func (D23_DC54) isD23() {}

func (CompanionStruct_D23_) Create_DC54_(Cf91 *C0) D23 {
	return D23{D23_DC54{Cf91}}
}

func (_this D23) Is_DC54() bool {
	_, ok := _this.Get_().(D23_DC54)
	return ok
}

type D23_DC57 struct {
	Cf95 D23
}

func (D23_DC57) isD23() {}

func (CompanionStruct_D23_) Create_DC57_(Cf95 D23) D23 {
	return D23{D23_DC57{Cf95}}
}

func (_this D23) Is_DC57() bool {
	_, ok := _this.Get_().(D23_DC57)
	return ok
}

func (CompanionStruct_D23_) Default() D23 {
	return Companion_D23_.Create_DC55_(_dafny.EmptySeq, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D23) Dtor_cf92() _dafny.Sequence {
	return _this.Get_().(D23_DC55).Cf92
}

func (_this D23) Dtor_cf93() _dafny.Array {
	return _this.Get_().(D23_DC55).Cf93
}

func (_this D23) Dtor_cf94() _dafny.Int {
	return _this.Get_().(D23_DC55).Cf94
}

func (_this D23) Dtor_cf91() *C0 {
	return _this.Get_().(D23_DC54).Cf91
}

func (_this D23) Dtor_cf95() D23 {
	return _this.Get_().(D23_DC57).Cf95
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC55:
		{
			return "D23.DC55" + "(" + _dafny.String(data.Cf92) + ", " + _dafny.String(data.Cf93) + ", " + _dafny.String(data.Cf94) + ")"
		}
	case D23_DC56:
		{
			return "D23.DC56"
		}
	case D23_DC54:
		{
			return "D23.DC54" + "(" + _dafny.String(data.Cf91) + ")"
		}
	case D23_DC57:
		{
			return "D23.DC57" + "(" + _dafny.String(data.Cf95) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC55:
		{
			data2, ok := other.Get_().(D23_DC55)
			return ok && data1.Cf92.Equals(data2.Cf92) && data1.Cf93 == data2.Cf93 && data1.Cf94.Cmp(data2.Cf94) == 0
		}
	case D23_DC56:
		{
			_, ok := other.Get_().(D23_DC56)
			return ok
		}
	case D23_DC54:
		{
			data2, ok := other.Get_().(D23_DC54)
			return ok && data1.Cf91 == data2.Cf91
		}
	case D23_DC57:
		{
			data2, ok := other.Get_().(D23_DC57)
			return ok && data1.Cf95.Equals(data2.Cf95)
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

type D24_DC59 struct {
	Cf97 _dafny.CodePoint
}

func (D24_DC59) isD24() {}

func (CompanionStruct_D24_) Create_DC59_(Cf97 _dafny.CodePoint) D24 {
	return D24{D24_DC59{Cf97}}
}

func (_this D24) Is_DC59() bool {
	_, ok := _this.Get_().(D24_DC59)
	return ok
}

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

type D24_DC58 struct {
	Cf96 _dafny.MultiSet
}

func (D24_DC58) isD24() {}

func (CompanionStruct_D24_) Create_DC58_(Cf96 _dafny.MultiSet) D24 {
	return D24{D24_DC58{Cf96}}
}

func (_this D24) Is_DC58() bool {
	_, ok := _this.Get_().(D24_DC58)
	return ok
}

func (CompanionStruct_D24_) Default() D24 {
	return Companion_D24_.Create_DC59_(_dafny.CodePoint('D'))
}

func (_this D24) Dtor_cf97() _dafny.CodePoint {
	return _this.Get_().(D24_DC59).Cf97
}

func (_this D24) Dtor_cf96() _dafny.MultiSet {
	return _this.Get_().(D24_DC58).Cf96
}

func (_this D24) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D24_DC59:
		{
			return "D24.DC59" + "(" + _dafny.String(data.Cf97) + ")"
		}
	case D24_DC60:
		{
			return "D24.DC60"
		}
	case D24_DC58:
		{
			return "D24.DC58" + "(" + _dafny.String(data.Cf96) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D24) Equals(other D24) bool {
	switch data1 := _this.Get_().(type) {
	case D24_DC59:
		{
			data2, ok := other.Get_().(D24_DC59)
			return ok && data1.Cf97 == data2.Cf97
		}
	case D24_DC60:
		{
			_, ok := other.Get_().(D24_DC60)
			return ok
		}
	case D24_DC58:
		{
			data2, ok := other.Get_().(D24_DC58)
			return ok && data1.Cf96.Equals(data2.Cf96)
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

type D25_DC62 struct {
	Cf99 bool
}

func (D25_DC62) isD25() {}

func (CompanionStruct_D25_) Create_DC62_(Cf99 bool) D25 {
	return D25{D25_DC62{Cf99}}
}

func (_this D25) Is_DC62() bool {
	_, ok := _this.Get_().(D25_DC62)
	return ok
}

type D25_DC61 struct {
	Cf98 *C4
}

func (D25_DC61) isD25() {}

func (CompanionStruct_D25_) Create_DC61_(Cf98 *C4) D25 {
	return D25{D25_DC61{Cf98}}
}

func (_this D25) Is_DC61() bool {
	_, ok := _this.Get_().(D25_DC61)
	return ok
}

func (CompanionStruct_D25_) Default() D25 {
	return Companion_D25_.Create_DC62_(false)
}

func (_this D25) Dtor_cf99() bool {
	return _this.Get_().(D25_DC62).Cf99
}

func (_this D25) Dtor_cf98() *C4 {
	return _this.Get_().(D25_DC61).Cf98
}

func (_this D25) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D25_DC62:
		{
			return "D25.DC62" + "(" + _dafny.String(data.Cf99) + ")"
		}
	case D25_DC61:
		{
			return "D25.DC61" + "(" + _dafny.String(data.Cf98) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D25) Equals(other D25) bool {
	switch data1 := _this.Get_().(type) {
	case D25_DC62:
		{
			data2, ok := other.Get_().(D25_DC62)
			return ok && data1.Cf99 == data2.Cf99
		}
	case D25_DC61:
		{
			data2, ok := other.Get_().(D25_DC61)
			return ok && data1.Cf98 == data2.Cf98
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

type D26_DC63 struct {
	Cf100 _dafny.Map
}

func (D26_DC63) isD26() {}

func (CompanionStruct_D26_) Create_DC63_(Cf100 _dafny.Map) D26 {
	return D26{D26_DC63{Cf100}}
}

func (_this D26) Is_DC63() bool {
	_, ok := _this.Get_().(D26_DC63)
	return ok
}

func (CompanionStruct_D26_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D26) Dtor_cf100() _dafny.Map {
	return _this.Get_().(D26_DC63).Cf100
}

func (_this D26) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D26_DC63:
		{
			return "D26.DC63" + "(" + _dafny.String(data.Cf100) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D26) Equals(other D26) bool {
	switch data1 := _this.Get_().(type) {
	case D26_DC63:
		{
			data2, ok := other.Get_().(D26_DC63)
			return ok && data1.Cf100.Equals(data2.Cf100)
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm3(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int
	Fm4(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool
	F11() _dafny.Int
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
	Fm3(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int
	Fm4(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool
	F11() _dafny.Int
	Fm5(p0 D0, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) bool
	Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int
	M0(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet
	M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Map)
	F12() _dafny.Set
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
	F1   _dafny.Int
	F5   _dafny.Array
	F6   bool
	F7   bool
	_f0  bool
	_f2  _dafny.Int
	_f3  _dafny.Int
	_f4  bool
	_f8  bool
	_f9  _dafny.Int
	_f10 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.Zero
	_this.F5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F6 = false
	_this.F7 = false
	_this._f0 = false
	_this._f2 = _dafny.Zero
	_this._f3 = _dafny.Zero
	_this._f4 = false
	_this._f8 = false
	_this._f9 = _dafny.Zero
	_this._f10 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 _dafny.Int, f3 _dafny.Int, f4 bool, f5 _dafny.Array, f6 bool, f7 bool, f8 bool, f9 _dafny.Int, f10 bool) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this).F6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F2() _dafny.Int {
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
func (_this *GlobalState) F8() bool {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() bool {
	{
		return _this._f10
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f20 _dafny.Sequence
	_f21 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f20 = _dafny.EmptySeq
	_this._f21 = _dafny.Zero
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

func (_this *C0) Ctor__(f20 _dafny.Sequence, f21 _dafny.Int) {
	{
		(_this)._f20 = f20
		(_this)._f21 = f21
	}
}
func (_this *C0) Fm21(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		if !(_dafny.MultiSetOf(false, true)).Equals(_dafny.MultiSetOf(true)) {
			return (!(!(true))) && (false)
		} else {
			return false
		}
	}
}
func (_this *C0) F20() _dafny.Sequence {
	{
		return _this._f20
	}
}
func (_this *C0) F21() _dafny.Int {
	{
		return _this._f21
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f18 bool
	_f19 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f18 = false
	_this._f19 = false
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__(f18 bool, f19 bool) {
	{
		(_this)._f18 = f18
		(_this)._f19 = f19
	}
}
func (_this *C1) Fm13(p0 D2, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_this).F18()
	}
}
func (_this *C1) Fm14(p0 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F18()
	}
}
func (_this *C1) M6(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) (_dafny.Int, _dafny.Array) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _357_v0 _dafny.MultiSet
		_ = _357_v0
		_357_v0 = _dafny.MultiSetOf((_this).F18(), (_this).F18(), (_this).F18())
		var _358_v1 _dafny.Set
		_ = _358_v1
		_358_v1 = _dafny.SetOf(Companion_Default___.Fm15(p2, globalState))
		var _359_v3 _dafny.Set
		_ = _359_v3
		_359_v3 = _dafny.SetOf(_358_v1, _358_v1, func() _dafny.Set {
			var _coll34 = _dafny.NewBuilder()
			_ = _coll34
			for _iter35 := _dafny.Iterate((Companion_Default___.Fm16(p2, (p3).Cardinality(), globalState)).Keys().Elements()); ; {
				_compr_34, _ok35 := _iter35()
				if !_ok35 {
					break
				}
				var _360_v2 _dafny.Sequence
				_360_v2 = interface{}(_compr_34).(_dafny.Sequence)
				if (Companion_Default___.Fm16(p2, (p3).Cardinality(), globalState)).Contains(_360_v2) {
					_coll34.Add(_360_v2)
				}
			}
			return _coll34.ToSet()
		}())
		var _361_v4 _dafny.Sequence
		_ = _361_v4
		_361_v4 = _dafny.SeqOf(p1)
		var _362_v5 _dafny.Map
		_ = _362_v5
		_362_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), p1)
		var _363_v6 _dafny.Map
		_ = _363_v6
		_363_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_362_v5, (_this).F19())
		var _364_v8 D2
		_ = _364_v8
		_364_v8 = Companion_D2_.Create_DC5_()
		var _365_v9 _dafny.Map
		_ = _365_v9
		_365_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_dafny.Zero).Minus(p1))
		var _366_v10 _dafny.Sequence
		_ = _366_v10
		_366_v10 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm17(_363_v6, p1, Companion_D2_.Create_DC5_(), (_this).F18(), globalState), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm17(_363_v6, p1, Companion_D2_.Create_DC5_(), (_this).F18(), globalState)).Cardinality()))).Uint32(), ((_365_v9).Update(_dafny.IntOfInt64(-171), p1)).Cardinality()), _361_v4, Companion_Default___.Fm17(_363_v6, _dafny.IntOfInt64(415), _364_v8, (_this).F19(), globalState))
		var _367_v11 _dafny.Sequence
		_ = _367_v11
		_367_v11 = _dafny.SeqOf(_361_v4, Companion_Default___.Fm17(_363_v6, (func() _dafny.Map {
			var _coll35 = _dafny.NewMapBuilder()
			_ = _coll35
			for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-770), _dafny.IntOfInt64(-776))); ; {
				_compr_35, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _368_v7 _dafny.Int
				_368_v7 = interface{}(_compr_35).(_dafny.Int)
				if ((_dafny.IntOfInt64(-770)).Cmp(_368_v7) <= 0) && ((_368_v7).Cmp(_dafny.IntOfInt64(-776)) < 0) {
					_coll35.Add((_368_v7).Times(p2), (_this).F18())
				}
			}
			return _coll35.ToMap()
		}()).Cardinality(), _364_v8, (_this).F18(), globalState), _361_v4, (_366_v10).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_366_v10).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _369_v12 D2
		_ = _369_v12
		var _out15 D2
		_ = _out15
		_out15 = (_this).M7(_357_v0, _359_v3, (_367_v11).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm2(p2, _dafny.IntOfInt64(450), globalState), _dafny.IntOfUint32((_367_v11).Cardinality()))).Uint32()).(_dafny.Sequence), globalState)
		_369_v12 = _out15
		var _370_v13 _dafny.Sequence
		_ = _370_v13
		_370_v13 = _dafny.SeqOf((_this).F19(), (_this).F18(), (p2).Cmp(p2) < 0, (_this).F18())
		if (_370_v13).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_370_v13).Cardinality()))).Uint32()).(bool) {
			var _371_v14 _dafny.Array
			_ = _371_v14
			var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
			_ = _nw48
			_371_v14 = _nw48
			_371_v14 = _371_v14
			var _372_v15 _dafny.Array
			_ = _372_v15
			var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
			_ = _nw49
			_372_v15 = _nw49
			var _373_v16 _dafny.CodePoint
			_ = _373_v16
			_373_v16 = _dafny.CodePoint('k')
			var _374_v17 _dafny.Array
			_ = _374_v17
			var _nw50 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw50
			_374_v17 = _nw50
			var _375_v18 _dafny.Set
			_ = _375_v18
			_375_v18 = _dafny.SetOf(_374_v17)
			var _376_v19 _dafny.Array
			_ = _376_v19
			var _nwElement0_15 _dafny.Sequence = Companion_Default___.Fm1(globalState)
			_ = _nwElement0_15
			var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(21))
			_ = _nw51
			(_nw51).ArraySet1(_nwElement0_15, 0)
			(_nw51).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(149))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg46 _dafny.Int) interface{} {
					return coer46(arg46)
				}
			}(func(_377_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			})), 1)
			(_nw51).ArraySet1(p0, 2)
			(_nw51).ArraySet1(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _373_v16), 3)
			(_nw51).ArraySet1(p0, 4)
			(_nw51).ArraySet1(p0, 5)
			(_nw51).ArraySet1(p0, 6)
			(_nw51).ArraySet1(p0, 7)
			(_nw51).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vpilsnhk"), 8)
			(_nw51).ArraySet1(p0, 9)
			(_nw51).ArraySet1(p0, 10)
			(_nw51).ArraySet1(p0, 11)
			(_nw51).ArraySet1(p0, 12)
			(_nw51).ArraySet1(p0, 13)
			(_nw51).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gn"), 14)
			(_nw51).ArraySet1(p0, 15)
			(_nw51).ArraySet1(p0, 16)
			(_nw51).ArraySet1(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_361_v4).Select((Companion_Default___.SafeIndex((_375_v18).Cardinality(), _dafny.IntOfUint32((_361_v4).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _dafny.CodePoint('a')), 17)
			(_nw51).ArraySet1(p0, 18)
			(_nw51).ArraySet1(p0, 19)
			(_nw51).ArraySet1(p0, 20)
			_376_v19 = _nw51
			var _378_v20 _dafny.Map
			_ = _378_v20
			_378_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_372_v15, _376_v19)
			_378_v20 = (_378_v20).Update(_372_v15, (func() _dafny.Array {
				if (_370_v13).Select((Companion_Default___.SafeIndex((_365_v9).Cardinality(), _dafny.IntOfUint32((_370_v13).Cardinality()))).Uint32()).(bool) {
					return _376_v19
				}
				return _376_v19
			})())
			var _379_v21 _dafny.Array
			_ = _379_v21
			var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
			_ = _nw52
			_379_v21 = _nw52
			var _380_v22 _dafny.Set
			_ = _380_v22
			_380_v22 = _dafny.SetOf(p1)
			var _381_v23 _dafny.Array
			_ = _381_v23
			var _nwElement0_16 _dafny.Set = _380_v22
			_ = _nwElement0_16
			var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(3))
			_ = _nw53
			(_nw53).ArraySet1(_nwElement0_16, 0)
			(_nw53).ArraySet1(_380_v22, 1)
			(_nw53).ArraySet1(Companion_Default___.Fm18(_dafny.IntOfUint32((Companion_Default___.Fm19(true, (_this).F18(), globalState)).Cardinality()), globalState), 2)
			_381_v23 = _nw53
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_379_v21), 0))
			_ = _index42
			(_379_v21).ArraySet1(_381_v23, (_index42).Int())
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_379_v21), 0))
			_ = _index43
			var _nwElement0_17 _dafny.Set = _380_v22
			_ = _nwElement0_17
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(3))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_17, 0)
			(_nw54).ArraySet1(_380_v22, 1)
			(_nw54).ArraySet1(_380_v22, 2)
			(_379_v21).ArraySet1(_nw54, (_index43).Int())
			var _382_v24 _dafny.Set
			_ = _382_v24
			_382_v24 = _dafny.SetOf((_this).F19())
			var _383_v25 _dafny.Map
			_ = _383_v25
			_383_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_382_v24, (_this).F19())
			(globalState).F1 = (_383_v25).Cardinality()
			(globalState).F7 = (_this).F18()
		} else {
			var _384_v26 _dafny.Array
			_ = _384_v26
			var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
			_ = _nw55
			_384_v26 = _nw55
			var _385_v27 _dafny.Array
			_ = _385_v27
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_4
			var _nw56 _dafny.Array
			_ = _nw56
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw56 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) bool = func(_386_i1 _dafny.Int) bool {
					return false
				}
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw56 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw56).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw56).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_385_v27 = _nw56
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_384_v26), 0))
			_ = _index44
			(_384_v26).ArraySet1(_385_v27, (_index44).Int())
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_384_v26), 0))
			_ = _index45
			(_384_v26).ArraySet1(_385_v27, (_index45).Int())
			(globalState).F6 = Companion_Default___.Fm0(globalState)
			var _387_v28 _dafny.CodePoint
			_ = _387_v28
			_387_v28 = _dafny.CodePoint('f')
			var _388_v29 _dafny.Set
			_ = _388_v29
			_388_v29 = _dafny.SetOf(p2)
			var _389_v30 _dafny.Array
			_ = _389_v30
			var _nwElement0_18 _dafny.CodePoint = _387_v28
			_ = _nwElement0_18
			var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(5))
			_ = _nw57
			(_nw57).ArraySet1CodePoint(_nwElement0_18, 0)
			(_nw57).ArraySet1CodePoint(_387_v28, 1)
			(_nw57).ArraySet1CodePoint(_387_v28, 2)
			(_nw57).ArraySet1CodePoint(_387_v28, 3)
			(_nw57).ArraySet1CodePoint(_387_v28, 4)
			_389_v30 = _nw57
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_389_v30), 0))
			_ = _index46
			(_389_v30).ArraySet1CodePoint(_387_v28, (_index46).Int())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v30), 0))
			_ = _index47
			(_389_v30).ArraySet1CodePoint(_387_v28, (_index47).Int())
			var _390_v31 _dafny.MultiSet
			_ = _390_v31
			_390_v31 = _dafny.MultiSetOf(Companion_D2_.Create_DC4_(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((p0).Cardinality()), p2))
			var _391_v32 _dafny.Map
			_ = _391_v32
			_391_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_390_v31, _388_v29)
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_389_v30), 0))
			_ = _index48
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v30), 0))
			_ = _index49
			var _rhs61 _dafny.CodePoint = _387_v28
			_ = _rhs61
			var _rhs62 _dafny.Set = (func() _dafny.Set {
				if (_391_v32).Contains(_dafny.MultiSetOf(_369_v12, _369_v12, _369_v12)) {
					return (_391_v32).Get(_dafny.MultiSetOf(_369_v12, _369_v12, _369_v12)).(_dafny.Set)
				}
				return _388_v29
			})()
			_ = _rhs62
			var _rhs63 _dafny.CodePoint = Companion_Default___.Fm20(_370_v13, (p1).Minus(_dafny.IntOfInt64(936)), (_this).F18(), globalState)
			_ = _rhs63
			var _rhs64 _dafny.CodePoint = _387_v28
			_ = _rhs64
			var _lhs54 _dafny.Array = _389_v30
			_ = _lhs54
			var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_389_v30), 0))
			_ = _lhs55
			var _lhs56 _dafny.Array = _389_v30
			_ = _lhs56
			var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_389_v30), 0))
			_ = _lhs57
			_387_v28 = _rhs61
			_388_v29 = _rhs62
			(_lhs54).ArraySet1CodePoint(_rhs63, (_lhs55).Int())
			(_lhs56).ArraySet1CodePoint(_rhs64, (_lhs57).Int())
			_385_v27 = _385_v27
			var _392_v33 _dafny.Sequence
			_ = _392_v33
			_392_v33 = _370_v13
			var _393_v34 *C0
			_ = _393_v34
			var _nw58 *C0 = New_C0_()
			_ = _nw58
			_nw58.Ctor__((_392_v33), (Companion_Default___.Fm2(p1, p2, globalState)).Times(_dafny.IntOfInt64(353)))
			_393_v34 = _nw58
		}
		(globalState).F6 = (_this).F19()
		var _394_v35 D0
		_ = _394_v35
		_394_v35 = Companion_D0_.Create_DC0_(p1)
		var _source7 D0 = _394_v35
		_ = _source7
		if _source7.Is_DC1() {
			var _395___mcc_h0 bool = _source7.Get_().(D0_DC1).Cf1
			_ = _395___mcc_h0
			var _396___mcc_h1 _dafny.Int = _source7.Get_().(D0_DC1).Cf2
			_ = _396___mcc_h1
			var _397___mcc_h2 bool = _source7.Get_().(D0_DC1).Cf3
			_ = _397___mcc_h2
			var _398_cf3 bool = _397___mcc_h2
			_ = _398_cf3
			var _399_cf2 _dafny.Int = _396___mcc_h1
			_ = _399_cf2
			var _400_cf1 bool = _395___mcc_h0
			_ = _400_cf1
			(globalState).F7 = (_370_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-413), _dafny.IntOfUint32((_370_v13).Cardinality()))).Uint32()).(bool)
			var _401_v36 _dafny.MultiSet
			_ = _401_v36
			_401_v36 = _dafny.MultiSetOf(_399_cf2, (_dafny.Zero).Minus(p1))
			_398_cf3 = ((func() _dafny.Int {
				if (_401_v36).Contains((_dafny.Zero).Minus((p3).Cardinality())) {
					return (_401_v36).Multiplicity((_dafny.Zero).Minus((p3).Cardinality()))
				}
				return p2
			})()).Cmp(_399_cf2) < 0
			var _402_v37 *C0
			_ = _402_v37
			var _nw59 *C0 = New_C0_()
			_ = _nw59
			_nw59.Ctor__(_dafny.Companion_Sequence_.Concatenate(_370_v13, _370_v13), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(906), _399_cf2))
			_402_v37 = _nw59
			var _403_v38 _dafny.Map
			_ = _403_v38
			_403_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_364_v8, ((Companion_Default___.Fm22(p0, true, globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F19()))).Cardinality())
			var _404_v39 _dafny.Sequence
			_ = _404_v39
			_404_v39 = _dafny.UnicodeSeqOfUtf8Bytes("jr")
			var _405_v40 _dafny.Array
			_ = _405_v40
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_5
			var _nw60 _dafny.Array
			_ = _nw60
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw60 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Int = (func(_406_v37 *C0) func(_dafny.Int) _dafny.Int {
					return func(_407_i2 _dafny.Int) _dafny.Int {
						return (_407_i2).Minus((_406_v37).F21())
					}
				})(_402_v37)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw60 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw60).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw60).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_405_v40 = _nw60
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_405_v40), 0))
			_ = _index50
			(_405_v40).ArraySet1((p2).Plus(_399_cf2), (_index50).Int())
			var _408_v41 _dafny.CodePoint
			_ = _408_v41
			_408_v41 = _dafny.CodePoint('v')
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_405_v40), 0))
			_ = _index51
			var _rhs65 _dafny.Map = _403_v38
			_ = _rhs65
			var _rhs66 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(globalState), _404_v39), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((func() _dafny.Int {
				if (_this).F18() {
					return _399_cf2
				}
				return (_dafny.Zero).Minus((_369_v12).Dtor_cf8())
			})()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(globalState), _404_v39)).Cardinality()))).Uint32(), _408_v41)
			_ = _rhs66
			var _rhs67 D2 = _369_v12
			_ = _rhs67
			var _rhs68 _dafny.Int = p1
			_ = _rhs68
			var _rhs69 _dafny.Int = p1
			_ = _rhs69
			var _lhs58 _dafny.Array = _405_v40
			_ = _lhs58
			var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_405_v40), 0))
			_ = _lhs59
			_403_v38 = _rhs65
			_404_v39 = _rhs66
			_369_v12 = _rhs67
			(_lhs58).ArraySet1(_rhs68, (_lhs59).Int())
			r0 = _rhs69
		} else {
			var _409___mcc_h3 _dafny.Int = _source7.Get_().(D0_DC0).Cf0
			_ = _409___mcc_h3
			var _410_cf0 _dafny.Int = _409___mcc_h3
			_ = _410_cf0
			var _411_v42 _dafny.Array
			_ = _411_v42
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_6
			var _nw61 _dafny.Array
			_ = _nw61
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw61 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Int = (func(_412_v0 _dafny.MultiSet, _413_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_414_i3 _dafny.Int) _dafny.Int {
						return (_414_i3).Minus((func() _dafny.Int {
							if (_412_v0).Contains((_this).F19()) {
								return (_412_v0).Multiplicity((_this).F19())
							}
							return _413_p1
						})())
					}
				})(_357_v0, p1)
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
			_411_v42 = _nw61
			var _415_v43 _dafny.MultiSet
			_ = _415_v43
			_415_v43 = _dafny.MultiSetOf(_411_v42)
			var _416_v44 _dafny.Sequence
			_ = _416_v44
			_416_v44 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(973))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}(func(_417_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			})))
			var _418_v45 _dafny.Map
			_ = _418_v45
			_418_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_415_v43).Intersection(_415_v43), _416_v44)
			_418_v45 = (_418_v45).Update(_415_v43, _416_v44)
			(globalState).F7 = (((_361_v4).Select((Companion_Default___.SafeIndex((_362_v5).Cardinality(), _dafny.IntOfUint32((_361_v4).Cardinality()))).Uint32()).(_dafny.Int)).Plus(_410_cf0)).Cmp(_410_cf0) <= 0
			var _419_v46 _dafny.MultiSet
			_ = _419_v46
			_419_v46 = _dafny.MultiSetOf(_361_v4)
			var _420_v47 _dafny.Array
			_ = _420_v47
			var _nw62 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw62
			_420_v47 = _nw62
			var _421_v48 _dafny.MultiSet
			_ = _421_v48
			_421_v48 = _dafny.MultiSetOf(_420_v47, _420_v47, _420_v47)
			var _422_v49 _dafny.Map
			_ = _422_v49
			_422_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), (_this).F18())
			var _423_v50 _dafny.Array
			_ = _423_v50
			var _nwElement0_19 bool = (_this).F19()
			_ = _nwElement0_19
			var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(24))
			_ = _nw63
			(_nw63).ArraySet1(_nwElement0_19, 0)
			(_nw63).ArraySet1((_419_v46).IsDisjointFrom(_419_v46), 1)
			(_nw63).ArraySet1((_this).F18(), 2)
			(_nw63).ArraySet1(!((_this).F18()) || ((_this).F19()), 3)
			(_nw63).ArraySet1(!((_this).F18()), 4)
			(_nw63).ArraySet1((_this).F18(), 5)
			(_nw63).ArraySet1((_this).F19(), 6)
			(_nw63).ArraySet1((_this).F18(), 7)
			(_nw63).ArraySet1((_this).F18(), 8)
			(_nw63).ArraySet1((p3).Equals(p3), 9)
			(_nw63).ArraySet1((_this).F18(), 10)
			(_nw63).ArraySet1((_this).F19(), 11)
			(_nw63).ArraySet1(false, 12)
			(_nw63).ArraySet1((_this).F19(), 13)
			(_nw63).ArraySet1((_421_v48).IsProperSubsetOf(_dafny.MultiSetOf(_420_v47)), 14)
			(_nw63).ArraySet1((_this).F18(), 15)
			(_nw63).ArraySet1((_this).F19(), 16)
			(_nw63).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm1(globalState), _dafny.UnicodeSeqOfUtf8Bytes("kppg")), 17)
			(_nw63).ArraySet1((_this).F18(), 18)
			(_nw63).ArraySet1((_this).F19(), 19)
			(_nw63).ArraySet1((_this).F18(), 20)
			(_nw63).ArraySet1((_this).F18(), 21)
			(_nw63).ArraySet1(!(((_422_v49).Cardinality()).Cmp(p1) >= 0), 22)
			(_nw63).ArraySet1((_this).F18(), 23)
			_423_v50 = _nw63
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_420_v47), 0))
			_ = _index52
			(_420_v47).ArraySet1((_this).F19(), (_index52).Int())
			var _424_v51 _dafny.Map
			_ = _424_v51
			_424_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _357_v0)
			var _425_v52 _dafny.Map
			_ = _425_v52
			_425_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm18(p2, globalState)).Cardinality(), _dafny.MultiSetFromSeq(_370_v13))
			var _426_v53 _dafny.MultiSet
			_ = _426_v53
			_426_v53 = (func() _dafny.MultiSet {
				if (_424_v51).Contains((_this).F19()) {
					return (_424_v51).Get((_this).F19()).(_dafny.MultiSet)
				}
				return _357_v0
			})()
			var _427_v54 _dafny.Array
			_ = _427_v54
			var _nwElement0_20 _dafny.MultiSet = _357_v0
			_ = _nwElement0_20
			var _nw64 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(28))
			_ = _nw64
			(_nw64).ArraySet1(_nwElement0_20, 0)
			(_nw64).ArraySet1(_357_v0, 1)
			(_nw64).ArraySet1((func() _dafny.MultiSet {
				if (_this).F19() {
					return _357_v0
				}
				return _dafny.MultiSetFromSeq(_370_v13)
			})(), 2)
			(_nw64).ArraySet1((_357_v0).Difference(_dafny.MultiSetOf((_this).F18())), 3)
			(_nw64).ArraySet1(_357_v0, 4)
			(_nw64).ArraySet1((_357_v0).Union(_357_v0), 5)
			(_nw64).ArraySet1(_357_v0, 6)
			(_nw64).ArraySet1((func() _dafny.MultiSet {
				if (_424_v51).Contains(!((_this).F18())) {
					return (_424_v51).Get(!((_this).F18())).(_dafny.MultiSet)
				}
				return _357_v0
			})(), 7)
			(_nw64).ArraySet1(_dafny.MultiSetOf((_this).F19()), 8)
			(_nw64).ArraySet1(_357_v0, 9)
			(_nw64).ArraySet1(_dafny.MultiSetOf((_this).F19(), !((_this).F19()), (_this).F18(), (_this).F19()), 10)
			(_nw64).ArraySet1(_357_v0, 11)
			(_nw64).ArraySet1((func() _dafny.MultiSet {
				if false {
					return _dafny.MultiSetOf((_this).F19(), false)
				}
				return _357_v0
			})(), 12)
			(_nw64).ArraySet1(_357_v0, 13)
			(_nw64).ArraySet1((func() _dafny.MultiSet {
				if (_425_v52).Contains(_410_cf0) {
					return (_425_v52).Get(_410_cf0).(_dafny.MultiSet)
				}
				return _dafny.MultiSetFromSeq(_370_v13)
			})(), 14)
			(_nw64).ArraySet1(_357_v0, 15)
			(_nw64).ArraySet1(_357_v0, 16)
			(_nw64).ArraySet1((_357_v0).Difference((_426_v53)), 17)
			(_nw64).ArraySet1(_357_v0, 18)
			(_nw64).ArraySet1((_357_v0).Union((_dafny.MultiSetOf((_this).F18())).Update((_this).F19(), Companion_Default___.Abs(_dafny.IntOfInt64(727)))), 19)
			(_nw64).ArraySet1(_357_v0, 20)
			(_nw64).ArraySet1(_357_v0, 21)
			(_nw64).ArraySet1((_dafny.MultiSetOf((_this).F19(), (_this).F19())).Intersection((_357_v0).Update(!((_this).F18()), Companion_Default___.Abs(p2))), 22)
			(_nw64).ArraySet1(_357_v0, 23)
			(_nw64).ArraySet1(_357_v0, 24)
			(_nw64).ArraySet1(_357_v0, 25)
			(_nw64).ArraySet1(_357_v0, 26)
			(_nw64).ArraySet1(_357_v0, 27)
			_427_v54 = _nw64
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_427_v54), 0))
			_ = _index53
			(_427_v54).ArraySet1((_357_v0).Update((_this).F18(), Companion_Default___.Abs(p2)), (_index53).Int())
			var _428_v55 D0
			_ = _428_v55
			_428_v55 = Companion_D0_.Create_DC1_(!((_this).F18()), p2, (_this).F19())
			var _429_v56 _dafny.CodePoint
			_ = _429_v56
			_429_v56 = _dafny.CodePoint('t')
			var _430_v57 _dafny.Array
			_ = _430_v57
			var _nwElement0_21 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_428_v55).Dtor_cf1() {
					return _429_v56
				}
				return Companion_Default___.Fm20(_370_v13, _410_cf0, (_this).F19(), globalState)
			})()
			_ = _nwElement0_21
			var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(11))
			_ = _nw65
			(_nw65).ArraySet1CodePoint(_nwElement0_21, 0)
			(_nw65).ArraySet1CodePoint(_dafny.CodePoint('t'), 1)
			(_nw65).ArraySet1CodePoint(_dafny.CodePoint('b'), 2)
			(_nw65).ArraySet1CodePoint(_dafny.CodePoint('x'), 3)
			(_nw65).ArraySet1CodePoint(_dafny.CodePoint('r'), 4)
			(_nw65).ArraySet1CodePoint(_429_v56, 5)
			(_nw65).ArraySet1CodePoint(_dafny.CodePoint('q'), 6)
			(_nw65).ArraySet1CodePoint(_dafny.CodePoint('e'), 7)
			(_nw65).ArraySet1CodePoint(_429_v56, 8)
			(_nw65).ArraySet1CodePoint(_429_v56, 9)
			(_nw65).ArraySet1CodePoint(_429_v56, 10)
			_430_v57 = _nw65
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_430_v57), 0))
			_ = _index54
			(_430_v57).ArraySet1CodePoint(_429_v56, (_index54).Int())
			var _431_v58 _dafny.Map
			_ = _431_v58
			_431_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), p0)
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_420_v47), 0))
			_ = _index55
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_427_v54), 0))
			_ = _index56
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_430_v57), 0))
			_ = _index57
			var _rhs70 _dafny.Int = _410_cf0
			_ = _rhs70
			var _rhs71 bool = !((_dafny.IntOfInt64(705)).Cmp((_410_cf0).Plus(_410_cf0)) >= 0)
			_ = _rhs71
			var _rhs72 bool = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(p0, p0), (func() _dafny.Sequence {
				if (_431_v58).Contains((_this).F18()) {
					return (_431_v58).Get((_this).F18()).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(91))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_432_v56 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_433_i5 _dafny.Int) _dafny.CodePoint {
						return _432_v56
					}
				})(_429_v56)))
			})())
			_ = _rhs72
			var _rhs73 _dafny.MultiSet = _357_v0
			_ = _rhs73
			var _rhs74 _dafny.CodePoint = _429_v56
			_ = _rhs74
			var _lhs60 *GlobalState = globalState
			_ = _lhs60
			var _lhs61 _dafny.Array = _420_v47
			_ = _lhs61
			var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_420_v47), 0))
			_ = _lhs62
			var _lhs63 *GlobalState = globalState
			_ = _lhs63
			var _lhs64 _dafny.Array = _427_v54
			_ = _lhs64
			var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_427_v54), 0))
			_ = _lhs65
			var _lhs66 _dafny.Array = _430_v57
			_ = _lhs66
			var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_430_v57), 0))
			_ = _lhs67
			_lhs60.F1 = _rhs70
			(_lhs61).ArraySet1(_rhs71, (_lhs62).Int())
			_lhs63.F6 = _rhs72
			(_lhs64).ArraySet1(_rhs73, (_lhs65).Int())
			(_lhs66).ArraySet1CodePoint(_rhs74, (_lhs67).Int())
			(globalState).F6 = (_420_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_420_v47), 0))).Int()).(bool)
		}
		(globalState).F1 = p2
		var _434_v59 D0
		_ = _434_v59
		_434_v59 = Companion_D0_.Create_DC1_(((_this).F18()) == ((_this).F19()), p1, !((_this).F18()) || ((_this).F18()))
		var _pat_let_tv21 = p1
		_ = _pat_let_tv21
		_434_v59 = func(_pat_let7_0 D0) D0 {
			return func(_435_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let8_0 bool) D0 {
					return func(_436_dt__update_hcf1_h0 bool) D0 {
						return func(_pat_let9_0 _dafny.Int) D0 {
							return func(_437_dt__update_hcf2_h0 _dafny.Int) D0 {
								return Companion_D0_.Create_DC1_(_436_dt__update_hcf1_h0, _437_dt__update_hcf2_h0, (_435_dt__update__tmp_h0).Dtor_cf3())
							}(_pat_let9_0)
						}(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(473), _pat_let_tv21))
					}(_pat_let8_0)
				}((_this).F18())
			}(_pat_let7_0)
		}(_434_v59)
		r0 = p1
		var _438_v60 D6
		_ = _438_v60
		_438_v60 = Companion_D6_.Create_DC9_(p3)
		var _pat_let_tv22 = p3
		_ = _pat_let_tv22
		var _439_v61 _dafny.Array
		_ = _439_v61
		var _nwElement0_22 _dafny.Map = p3
		_ = _nwElement0_22
		var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(10))
		_ = _nw66
		(_nw66).ArraySet1(_nwElement0_22, 0)
		(_nw66).ArraySet1((_438_v60).Dtor_cf12(), 1)
		(_nw66).ArraySet1((p3).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F19())), 2)
		(_nw66).ArraySet1(p3, 3)
		(_nw66).ArraySet1(p3, 4)
		(_nw66).ArraySet1((p3).Merge(p3), 5)
		(_nw66).ArraySet1(p3, 6)
		(_nw66).ArraySet1(((func(_pat_let10_0 D6) D6 {
			return func(_440_dt__update__tmp_h1 D6) D6 {
				return func(_pat_let11_0 _dafny.Map) D6 {
					return func(_441_dt__update_hcf12_h0 _dafny.Map) D6 {
						return Companion_D6_.Create_DC9_(_441_dt__update_hcf12_h0)
					}(_pat_let11_0)
				}(_pat_let_tv22)
			}(_pat_let10_0)
		}(_438_v60)).Dtor_cf12()).Merge((p3).Update(p2, (_this).F19())), 7)
		(_nw66).ArraySet1((p3).Merge(p3), 8)
		(_nw66).ArraySet1(p3, 9)
		_439_v61 = _nw66
		r1 = _439_v61
		return r0, r1
	}
}
func (_this *C1) M7(p0 _dafny.MultiSet, p1 _dafny.Set, p2 _dafny.Sequence, globalState *GlobalState) D2 {
	{
		var r0 D2 = Companion_D2_.Default()
		_ = r0
		var _442_v0 _dafny.Sequence
		_ = _442_v0
		_442_v0 = _dafny.UnicodeSeqOfUtf8Bytes("rpix")
		var _443_i0 _dafny.Int
		_ = _443_i0
		_443_i0 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm1(globalState), _dafny.Companion_Sequence_.Concatenate(_442_v0, _442_v0)) {
				{
					if (_443_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_443_i0 = (_443_i0).Plus(_dafny.One)
					var _444_v1 _dafny.Int
					_ = _444_v1
					_444_v1 = _dafny.IntOfInt64(239)
					(globalState).F6 = (func() bool {
						if ((_this).F19()) == ((_this).F19()) {
							return (_this).F18()
						}
						return (_444_v1).Cmp(_444_v1) < 0
					})()
					var _445_v2 _dafny.CodePoint
					_ = _445_v2
					_445_v2 = _dafny.CodePoint('s')
					_445_v2 = _445_v2
					var _446_v3 _dafny.Map
					_ = _446_v3
					_446_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _445_v2)
					_445_v2 = (func() _dafny.CodePoint {
						if (_446_v3).Contains((_this).F19()) {
							return (_446_v3).Get((_this).F19()).(_dafny.CodePoint)
						}
						return _445_v2
					})()
					var _447_v4 D0
					_ = _447_v4
					_447_v4 = Companion_D0_.Create_DC1_((_this).F19(), _444_v1, (p0).Contains((_this).F18()))
					var _448_v5 _dafny.Array
					_ = _448_v5
					var _len0_7 _dafny.Int = _dafny.IntOfInt64(24)
					_ = _len0_7
					var _nw67 _dafny.Array
					_ = _nw67
					if _len0_7.Cmp(_dafny.Zero) == 0 {
						_nw67 = _dafny.NewArray(_len0_7)
					} else {
						var _init7 func(_dafny.Int) _dafny.Int = (func(_449_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_450_i1 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeDivisionInt(_450_i1, _449_v1)
							}
						})(_444_v1)
						_ = _init7
						var _element0_7 = _init7(_dafny.Zero)
						_ = _element0_7
						_nw67 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
						(_nw67).ArraySet1(_element0_7, 0)
						var _nativeLen0_7 = (_len0_7).Int()
						_ = _nativeLen0_7
						for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
							(_nw67).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
						}
					}
					_448_v5 = _nw67
					var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_448_v5), 0))
					_ = _index58
					(_448_v5).ArraySet1(_444_v1, (_index58).Int())
					var _451_v6 _dafny.MultiSet
					_ = _451_v6
					_451_v6 = _dafny.MultiSetOf(_dafny.IntOfInt64(794))
					var _452_v7 _dafny.Map
					_ = _452_v7
					_452_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_451_v6, _444_v1)
					var _453_v8 _dafny.Sequence
					_ = _453_v8
					_453_v8 = _dafny.SeqOf((_this).F19())
					var _454_v9 _dafny.Sequence
					_ = _454_v9
					_454_v9 = _dafny.SeqOf(_442_v0, _442_v0, _442_v0, _dafny.UnicodeSeqOfUtf8Bytes("efknn"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(735))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg49 _dafny.Int) interface{} {
							return coer49(arg49)
						}
					}((func(_455_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_456_i2 _dafny.Int) _dafny.CodePoint {
							return _455_v2
						}
					})(_445_v2))))
					var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_448_v5), 0))
					_ = _index59
					var _rhs75 D0 = Companion_D0_.Create_DC1_((_this).F19(), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((Companion_Default___.Fm23(_444_v1, _444_v1, globalState)).Cardinality()), _444_v1), (_this).F19())
					_ = _rhs75
					var _rhs76 _dafny.Int = (_dafny.Zero).Minus(_444_v1)
					_ = _rhs76
					var _rhs77 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
						if (_452_v7).Contains(_451_v6) {
							return (_452_v7).Get(_451_v6).(_dafny.Int)
						}
						return _dafny.IntOfUint32((_453_v8).Cardinality())
					})())
					_ = _rhs77
					var _rhs78 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_this).F19() {
							return _442_v0
						}
						return (_454_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.IntOfUint32((_454_v9).Cardinality()))).Uint32()).(_dafny.Sequence)
					})()).Cardinality())
					_ = _rhs78
					var _lhs68 *GlobalState = globalState
					_ = _lhs68
					var _lhs69 _dafny.Array = _448_v5
					_ = _lhs69
					var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_448_v5), 0))
					_ = _lhs70
					_447_v4 = _rhs75
					_lhs68.F1 = _rhs76
					(_lhs69).ArraySet1(_rhs77, (_lhs70).Int())
					_444_v1 = _rhs78
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _457_v10 _dafny.Int
		_ = _457_v10
		_457_v10 = _dafny.IntOfInt64(907)
		var _458_v12 _dafny.Array
		_ = _458_v12
		var _nwElement0_23 _dafny.Int = Companion_Default___.Fm2(_457_v10, _457_v10, globalState)
		_ = _nwElement0_23
		var _nw68 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(5))
		_ = _nw68
		(_nw68).ArraySet1(_nwElement0_23, 0)
		(_nw68).ArraySet1(_457_v10, 1)
		(_nw68).ArraySet1(_dafny.IntOfUint32((_442_v0).Cardinality()), 2)
		(_nw68).ArraySet1((func() _dafny.Map {
			var _coll36 = _dafny.NewMapBuilder()
			_ = _coll36
			for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(180), _dafny.IntOfInt64(711))); ; {
				_compr_36, _ok37 := _iter37()
				if !_ok37 {
					break
				}
				var _459_v11 _dafny.Int
				_459_v11 = interface{}(_compr_36).(_dafny.Int)
				if ((_dafny.IntOfInt64(180)).Cmp(_459_v11) <= 0) && ((_459_v11).Cmp(_dafny.IntOfInt64(711)) < 0) {
					_coll36.Add((_459_v11).Plus(_dafny.IntOfUint32((_442_v0).Cardinality())), (_dafny.SetOf(_457_v10)).Cardinality())
				}
			}
			return _coll36.ToMap()
		}()).Cardinality(), 3)
		(_nw68).ArraySet1(_457_v10, 4)
		_458_v12 = _nw68
		var _460_v13 _dafny.Set
		_ = _460_v13
		_460_v13 = _dafny.SetOf(_458_v12)
		_460_v13 = _460_v13
		var _461_v14 _dafny.Map
		_ = _461_v14
		_461_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _457_v10)
		var _462_v15 _dafny.Set
		_ = _462_v15
		_462_v15 = _dafny.SetOf(!((_this).F19()), (_this).F18())
		_457_v10 = Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_461_v14).Contains((_this).F18()) {
				return (_461_v14).Get((_this).F18()).(_dafny.Int)
			}
			return _457_v10
		})(), Companion_Default___.SafeModuloInt((_462_v15).Cardinality(), _457_v10))
		(globalState).F6 = (_this).F19()
		var _463_i3 _dafny.Int
		_ = _463_i3
		_463_i3 = _dafny.Zero
		{
			for (_this).F18() {
				{
					if (_463_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_463_i3 = (_463_i3).Plus(_dafny.One)
					var _464_v16 _dafny.Set
					_ = _464_v16
					_464_v16 = _dafny.SetOf(_dafny.IntOfInt64(789))
					var _465_v17 D7
					_ = _465_v17
					_465_v17 = Companion_D7_.Create_DC12_(p2)
					_457_v10 = Companion_Default___.Fm2(((_dafny.SetOf(_457_v10, _457_v10, _457_v10, _457_v10, _457_v10)).Union(_464_v16)).Cardinality(), (_dafny.IntOfUint32(((_465_v17).Dtor_cf17()).Cardinality())).Minus(_457_v10), globalState)
					var _466_v18 D2
					_ = _466_v18
					_466_v18 = Companion_D2_.Create_DC4_(_457_v10, _457_v10, _457_v10)
					(globalState).F1 = Companion_Default___.Fm2(Companion_Default___.Fm2(_457_v10, _457_v10, globalState), (_457_v10).Times((_466_v18).Dtor_cf6()), globalState)
					var _467_v19 _dafny.Array
					_ = _467_v19
					var _len0_8 _dafny.Int = _dafny.IntOfInt64(3)
					_ = _len0_8
					var _nw69 _dafny.Array
					_ = _nw69
					if _len0_8.Cmp(_dafny.Zero) == 0 {
						_nw69 = _dafny.NewArray(_len0_8)
					} else {
						var _init8 func(_dafny.Int) _dafny.Set = (func(_468_v16 _dafny.Set) func(_dafny.Int) _dafny.Set {
							return func(_469_i4 _dafny.Int) _dafny.Set {
								return _468_v16
							}
						})(_464_v16)
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
					_467_v19 = _nw69
					var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_467_v19), 0))
					_ = _index60
					(_467_v19).ArraySet1(_464_v16, (_index60).Int())
					var _470_v20 _dafny.Map
					_ = _470_v20
					_470_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), Companion_Default___.Fm18(_dafny.IntOfInt64(519), globalState))
					var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_467_v19), 0))
					_ = _index61
					(_467_v19).ArraySet1(((func() _dafny.Set {
						if (_470_v20).Contains((_this).F18()) {
							return (_470_v20).Get((_this).F18()).(_dafny.Set)
						}
						return _464_v16
					})()).Union(_464_v16), (_index61).Int())
					(globalState).F6 = (_this).F19()
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _471_v21 _dafny.Array
		_ = _471_v21
		var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.One)
		_ = _nw70
		_471_v21 = _nw70
		_471_v21 = _471_v21
		var _472_v22 D0
		_ = _472_v22
		_472_v22 = Companion_D0_.Create_DC1_((_this).F19(), _457_v10, false)
		var _pat_let_tv23 = _457_v10
		_ = _pat_let_tv23
		var _pat_let_tv24 = _457_v10
		_ = _pat_let_tv24
		var _pat_let_tv25 = _457_v10
		_ = _pat_let_tv25
		r0 = func(_source8 _dafny.MultiSet) D2 {
			var _473___mcc_h0 _dafny.MultiSet = _source8
			_ = _473___mcc_h0
			var _474_cf11 _dafny.MultiSet = _473___mcc_h0
			_ = _474_cf11
			return Companion_D2_.Create_DC4_(_pat_let_tv23, _pat_let_tv24, _pat_let_tv25)
		}(Companion_Default___.Fm24((_472_v22).Dtor_cf3(), globalState))
		return r0
	}
}
func (_this *C1) F18() bool {
	{
		return _this._f18
	}
}
func (_this *C1) F19() bool {
	{
		return _this._f19
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f11 _dafny.Int
	_f24 _dafny.CodePoint
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f11 = _dafny.Zero
	_this._f24 = _dafny.CodePoint('D')
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

func (_this *C2) F11() _dafny.Int {
	return _this._f11
}
func (_this *C2) Ctor__(f24 _dafny.CodePoint, f11 _dafny.Int) {
	{
		(_this)._f24 = f24
		(_this)._f11 = f11
	}
}
func (_this *C2) Fm3(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F11()).Plus((_this).F11())
	}
}
func (_this *C2) Fm4(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(337))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg50 _dafny.Int) interface{} {
				return coer50(arg50)
			}
		}(func(_475_i0 _dafny.Int) _dafny.CodePoint {
			return (_this).F24()
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(428))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg51 _dafny.Int) interface{} {
				return coer51(arg51)
			}
		}(func(_476_i1 _dafny.Int) _dafny.CodePoint {
			return (_this).F24()
		}))), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pykmo"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-42))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg52 _dafny.Int) interface{} {
				return coer52(arg52)
			}
		}(func(_477_i2 _dafny.Int) _dafny.CodePoint {
			return (_this).F24()
		}))))
	}
}
func (_this *C2) M8(globalState *GlobalState) (bool, bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _478_v0 bool
		_ = _478_v0
		_478_v0 = false
		var _479_v1 _dafny.Map
		_ = _479_v1
		_479_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-775), (_this).Fm3((_this).F11(), !(_478_v0), _478_v0, globalState))
		var _480_v2 D8
		_ = _480_v2
		_480_v2 = Companion_D8_.Create_DC16_(true, (func() _dafny.Int {
			if (_479_v1).Contains((_this).F11()) {
				return (_479_v1).Get((_this).F11()).(_dafny.Int)
			}
			return (_this).F11()
		})())
		(globalState).F1 = ((_480_v2).Dtor_cf24()).Minus(((_this).F11()).Plus((_this).F11()))
		var _481_v3 _dafny.Sequence
		_ = _481_v3
		_481_v3 = _dafny.SeqOf(_478_v0, _478_v0)
		var _482_v4 *C0
		_ = _482_v4
		var _nw71 *C0 = New_C0_()
		_ = _nw71
		_nw71.Ctor__(_dafny.Companion_Sequence_.Concatenate(_481_v3, _481_v3), ((_this).F11()).Times(_dafny.IntOfInt64(687)))
		_482_v4 = _nw71
		var _483_v5 _dafny.Map
		_ = _483_v5
		_483_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_482_v4).F21()), !(_478_v0))
		var _484_v6 D6
		_ = _484_v6
		_484_v6 = Companion_D6_.Create_DC9_((_483_v5).Merge((_483_v5).Update((_482_v4).F21(), _478_v0)))
		var _source9 D6 = _484_v6
		_ = _source9
		if _source9.Is_DC10() {
			var _485___mcc_h0 _dafny.Sequence = _source9.Get_().(D6_DC10).Cf13
			_ = _485___mcc_h0
			var _486_cf13 _dafny.Sequence = _485___mcc_h0
			_ = _486_cf13
			var _487_v7 _dafny.Sequence
			_ = _487_v7
			_487_v7 = _dafny.UnicodeSeqOfUtf8Bytes("qpp")
			var _488_v8 _dafny.Sequence
			_ = _488_v8
			_488_v8 = _dafny.SeqOf(_dafny.IntOfUint32((_487_v7).Cardinality()))
			var _489_v9 D7
			_ = _489_v9
			_489_v9 = Companion_D7_.Create_DC12_(_488_v8)
			var _490_v10 _dafny.Map
			_ = _490_v10
			_490_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _478_v0)
			var _491_v11 _dafny.Map
			_ = _491_v11
			_491_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_490_v10, _478_v0)
			var _492_v12 _dafny.Map
			_ = _492_v12
			_492_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_491_v11).Cardinality(), _488_v8)
			var _493_v13 _dafny.Array
			_ = _493_v13
			var _nwElement0_24 _dafny.Sequence = _488_v8
			_ = _nwElement0_24
			var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(7))
			_ = _nw72
			(_nw72).ArraySet1(_nwElement0_24, 0)
			(_nw72).ArraySet1(_488_v8, 1)
			(_nw72).ArraySet1(_488_v8, 2)
			(_nw72).ArraySet1((_489_v9).Dtor_cf17(), 3)
			(_nw72).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(364))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}((func(_494_v4 *C0) func(_dafny.Int) _dafny.Int {
				return func(_495_i0 _dafny.Int) _dafny.Int {
					return (_494_v4).F21()
				}
			})(_482_v4))), 4)
			(_nw72).ArraySet1(_488_v8, 5)
			(_nw72).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_488_v8, (func() _dafny.Sequence {
				if (_492_v12).Contains(_dafny.IntOfUint32(((_482_v4).F20()).Cardinality())) {
					return (_492_v12).Get(_dafny.IntOfUint32(((_482_v4).F20()).Cardinality())).(_dafny.Sequence)
				}
				return _488_v8
			})()), 6)
			_493_v13 = _nw72
			var _rhs79 bool = ((_this).F11()).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus((_482_v4).F21()))) < 0
			_ = _rhs79
			var _rhs80 _dafny.Array = (func() _dafny.Array {
				if _478_v0 {
					return _493_v13
				}
				return _493_v13
			})()
			_ = _rhs80
			var _rhs81 bool = _478_v0
			_ = _rhs81
			var _lhs71 *GlobalState = globalState
			_ = _lhs71
			var _lhs72 *GlobalState = globalState
			_ = _lhs72
			_lhs71.F6 = _rhs79
			_493_v13 = _rhs80
			_lhs72.F6 = _rhs81
			var _496_v14 _dafny.MultiSet
			_ = _496_v14
			_496_v14 = _dafny.MultiSetOf(_478_v0)
			if !(_496_v14).Contains(_478_v0) {
				var _497_v15 _dafny.Map
				_ = _497_v15
				_497_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_478_v0, (_482_v4).F21())
				r3 = Companion_Default___.SafeDivisionInt((_this).Fm3(_dafny.IntOfUint32(((_482_v4).F20()).Cardinality()), _478_v0, _478_v0, globalState), (func() _dafny.Int {
					if (_497_v15).Contains(_478_v0) {
						return (_497_v15).Get(_478_v0).(_dafny.Int)
					}
					return (_this).F11()
				})())
				var _498_v16 _dafny.Set
				_ = _498_v16
				_498_v16 = _dafny.SetOf((_this).F24())
				var _499_v17 _dafny.Set
				_ = _499_v17
				_499_v17 = _dafny.SetOf(_498_v16)
				var _500_v18 _dafny.Map
				_ = _500_v18
				_500_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_487_v7, _499_v17)
				_500_v18 = (_500_v18).Update(_487_v7, Companion_Default___.Fm28(_478_v0, _dafny.IntOfInt64(-68), globalState))
				var _501_v19 _dafny.Array
				_ = _501_v19
				var _nw73 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw73
				_501_v19 = _nw73
				var _502_v20 _dafny.Array
				_ = _502_v20
				var _nwElement0_25 _dafny.Array = _501_v19
				_ = _nwElement0_25
				var _nw74 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(13))
				_ = _nw74
				(_nw74).ArraySet1(_nwElement0_25, 0)
				(_nw74).ArraySet1(_501_v19, 1)
				(_nw74).ArraySet1(_501_v19, 2)
				(_nw74).ArraySet1(_501_v19, 3)
				(_nw74).ArraySet1(_501_v19, 4)
				(_nw74).ArraySet1(_501_v19, 5)
				(_nw74).ArraySet1(_501_v19, 6)
				(_nw74).ArraySet1(_501_v19, 7)
				(_nw74).ArraySet1(_501_v19, 8)
				(_nw74).ArraySet1(_501_v19, 9)
				(_nw74).ArraySet1(_501_v19, 10)
				(_nw74).ArraySet1(_501_v19, 11)
				(_nw74).ArraySet1((func() _dafny.Array {
					if false {
						return _501_v19
					}
					return _501_v19
				})(), 12)
				_502_v20 = _nw74
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_502_v20), 0))
				_ = _index62
				(_502_v20).ArraySet1(_501_v19, (_index62).Int())
				var _503_v21 _dafny.MultiSet
				_ = _503_v21
				_503_v21 = _dafny.MultiSetOf((_this).F11(), (_482_v4).F21())
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_502_v20), 0))
				_ = _index63
				var _rhs82 _dafny.Int = (_dafny.Zero).Minus((_482_v4).F21())
				_ = _rhs82
				var _rhs83 _dafny.Int = ((_this).F11()).Plus(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_503_v21).Contains((_482_v4).F21()) {
						return (_503_v21).Multiplicity((_482_v4).F21())
					}
					return (func() _dafny.Int {
						if (_479_v1).Contains((_482_v4).F21()) {
							return (_479_v1).Get((_482_v4).F21()).(_dafny.Int)
						}
						return (_482_v4).F21()
					})()
				})(), (_this).F11()))
				_ = _rhs83
				var _rhs84 _dafny.Array = _501_v19
				_ = _rhs84
				var _lhs73 *GlobalState = globalState
				_ = _lhs73
				var _lhs74 *GlobalState = globalState
				_ = _lhs74
				var _lhs75 _dafny.Array = _502_v20
				_ = _lhs75
				var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_502_v20), 0))
				_ = _lhs76
				_lhs73.F1 = _rhs82
				_lhs74.F1 = _rhs83
				(_lhs75).ArraySet1(_rhs84, (_lhs76).Int())
				(globalState).F7 = _478_v0
				var _504_v22 _dafny.Set
				_ = _504_v22
				_504_v22 = _dafny.SetOf((func() bool {
					if _478_v0 {
						return _478_v0
					}
					return false
				})())
				var _505_v23 _dafny.CodePoint
				_ = _505_v23
				_505_v23 = _dafny.CodePoint('v')
				var _rhs85 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(249))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}((func(_506_v7 _dafny.Sequence, _507_v0 bool, _508_v1 _dafny.Map, _509_v4 *C0) func(_dafny.Int) _dafny.Int {
					return func(_510_i1 _dafny.Int) _dafny.Int {
						return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_506_v7).Cardinality()), _507_v0)).Cardinality()).Minus((_dafny.Zero).Minus((func() _dafny.Int {
							if (_508_v1).Contains(_dafny.IntOfInt64(968)) {
								return (_508_v1).Get(_dafny.IntOfInt64(968)).(_dafny.Int)
							}
							return (_509_v4).F21()
						})()))
					}
				})(_487_v7, _478_v0, _479_v1, _482_v4)))
				_ = _rhs85
				var _rhs86 _dafny.Set = _504_v22
				_ = _rhs86
				var _rhs87 _dafny.CodePoint = (func() _dafny.CodePoint {
					if _478_v0 {
						return (_this).F24()
					}
					return _505_v23
				})()
				_ = _rhs87
				var _rhs88 _dafny.Int = (_482_v4).F21()
				_ = _rhs88
				var _lhs77 *GlobalState = globalState
				_ = _lhs77
				_488_v8 = _rhs85
				_504_v22 = _rhs86
				_505_v23 = _rhs87
				_lhs77.F1 = _rhs88
			} else {
				_483_v5 = (_483_v5).Update((func() _dafny.Int {
					if _478_v0 {
						return Companion_Default___.Fm2((_dafny.Zero).Minus((_482_v4).F21()), (_dafny.Zero).Minus((_this).F11()), globalState)
					}
					return _dafny.IntOfInt64(-92)
				})(), _478_v0)
				var _511_v24 _dafny.Set
				_ = _511_v24
				_511_v24 = _dafny.SetOf((func() _dafny.Sequence {
					if _478_v0 {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(260))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg55 _dafny.Int) interface{} {
								return coer55(arg55)
							}
						}((func(_512_v4 *C0) func(_dafny.Int) _dafny.Int {
							return func(_513_i2 _dafny.Int) _dafny.Int {
								return (_512_v4).F21()
							}
						})(_482_v4)))
					}
					return _488_v8
				})())
				_511_v24 = _511_v24
				var _514_v25 _dafny.CodePoint
				_ = _514_v25
				_514_v25 = _dafny.CodePoint('i')
				var _515_v26 _dafny.Set
				_ = _515_v26
				_515_v26 = _dafny.SetOf(Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_478_v0, (_482_v4).F21())))
				var _516_v27 _dafny.Array
				_ = _516_v27
				var _nwElement0_26 _dafny.Set = _515_v26
				_ = _nwElement0_26
				var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(9))
				_ = _nw75
				(_nw75).ArraySet1(_nwElement0_26, 0)
				(_nw75).ArraySet1(_515_v26, 1)
				(_nw75).ArraySet1(_515_v26, 2)
				(_nw75).ArraySet1(_515_v26, 3)
				(_nw75).ArraySet1((_515_v26).Difference(_515_v26), 4)
				(_nw75).ArraySet1(_515_v26, 5)
				(_nw75).ArraySet1(Companion_Default___.Fm29(globalState), 6)
				(_nw75).ArraySet1(_515_v26, 7)
				(_nw75).ArraySet1(_515_v26, 8)
				_516_v27 = _nw75
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_516_v27), 0))
				_ = _index64
				(_516_v27).ArraySet1(_515_v26, (_index64).Int())
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_516_v27), 0))
				_ = _index65
				var _rhs89 _dafny.CodePoint = (_this).F24()
				_ = _rhs89
				var _rhs90 _dafny.Set = _515_v26
				_ = _rhs90
				var _lhs78 _dafny.Array = _516_v27
				_ = _lhs78
				var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_516_v27), 0))
				_ = _lhs79
				_514_v25 = _rhs89
				(_lhs78).ArraySet1(_rhs90, (_lhs79).Int())
				var _517_v28 D8
				_ = _517_v28
				_517_v28 = Companion_D8_.Create_DC15_(_490_v10)
				var _518_v29 D8
				_ = _518_v29
				_518_v29 = Companion_D8_.Create_DC17_(_517_v28)
				_518_v29 = _518_v29
				var _519_v30 *C0
				_ = _519_v30
				var _nw76 *C0 = New_C0_()
				_ = _nw76
				_nw76.Ctor__(_dafny.SeqOf(_478_v0), Companion_Default___.SafeDivisionInt((_482_v4).F21(), (_482_v4).F21()))
				_519_v30 = _nw76
			}
			var _520_v31 _dafny.Array
			_ = _520_v31
			var _nw77 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw77
			_520_v31 = _nw77
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_520_v31), 0))
			_ = _index66
			(_520_v31).ArraySet1((_482_v4).F21(), (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_520_v31), 0))
			_ = _index67
			(_520_v31).ArraySet1(Companion_Default___.SafeModuloInt((_482_v4).F21(), (_480_v2).Dtor_cf24()), (_index67).Int())
			r3 = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(359), (_482_v4).F21())).Times((_this).F11())
		} else if _source9.Is_DC11() {
			var _521___mcc_h1 _dafny.Int = _source9.Get_().(D6_DC11).Cf14
			_ = _521___mcc_h1
			var _522___mcc_h2 _dafny.Int = _source9.Get_().(D6_DC11).Cf15
			_ = _522___mcc_h2
			var _523___mcc_h3 _dafny.Int = _source9.Get_().(D6_DC11).Cf16
			_ = _523___mcc_h3
			var _524_cf16 _dafny.Int = _523___mcc_h3
			_ = _524_cf16
			var _525_cf15 _dafny.Int = _522___mcc_h2
			_ = _525_cf15
			var _526_cf14 _dafny.Int = _521___mcc_h1
			_ = _526_cf14
			var _527_v32 _dafny.Sequence
			_ = _527_v32
			_527_v32 = _dafny.SeqOf((_this).F11(), _526_cf14)
			var _528_v33 _dafny.Array
			_ = _528_v33
			var _nwElement0_27 _dafny.Int = (_this).F11()
			_ = _nwElement0_27
			var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(8))
			_ = _nw78
			(_nw78).ArraySet1(_nwElement0_27, 0)
			(_nw78).ArraySet1((_482_v4).F21(), 1)
			(_nw78).ArraySet1(_dafny.IntOfInt64(552), 2)
			(_nw78).ArraySet1(_dafny.IntOfUint32((_527_v32).Cardinality()), 3)
			(_nw78).ArraySet1(_524_cf16, 4)
			(_nw78).ArraySet1(_dafny.IntOfInt64(683), 5)
			(_nw78).ArraySet1((_dafny.Zero).Minus((_this).F11()), 6)
			(_nw78).ArraySet1(_526_cf14, 7)
			_528_v33 = _nw78
			var _529_v34 _dafny.Map
			_ = _529_v34
			_529_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if _478_v0 {
					return _528_v33
				}
				return _528_v33
			})(), _478_v0)
			_529_v34 = (_529_v34).Update(_528_v33, (_482_v4).Fm21(_478_v0, _dafny.IntOfInt64(777), globalState))
			var _530_v35 _dafny.Map
			_ = _530_v35
			_530_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_479_v1).Cardinality(), _528_v33)
			_530_v35 = _530_v35
			var _531_v36 _dafny.Map
			_ = _531_v36
			_531_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_527_v32).Cardinality())).Times((_482_v4).F21()), (_this).F24())
			var _532_v37 D7
			_ = _532_v37
			_532_v37 = Companion_D7_.Create_DC12_(_527_v32)
			var _533_v39 _dafny.MultiSet
			_ = _533_v39
			_533_v39 = _dafny.MultiSetOf(_dafny.IntOfInt64(557), _524_cf16, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iwljmxlt")).Cardinality()))
			var _534_v40 _dafny.Map
			_ = _534_v40
			_534_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_478_v0, _482_v4)
			var _rhs91 bool = (func() _dafny.Set {
				var _coll37 = _dafny.NewBuilder()
				_ = _coll37
				for _iter38 := _dafny.Iterate((_dafny.SeqOf((_532_v37).Dtor_cf17(), _dafny.SeqOf(_dafny.IntOfInt64(144)))).Elements()); ; {
					_compr_37, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _535_v38 _dafny.Sequence
					_535_v38 = interface{}(_compr_37).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_532_v37).Dtor_cf17(), _dafny.SeqOf(_dafny.IntOfInt64(144))), _535_v38) {
						_coll37.Add(_535_v38)
					}
				}
				return _coll37.ToSet()
			}()).Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_482_v4).F21()), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_482_v4).F21()), _dafny.IntOfUint32((_dafny.SeqOf((_482_v4).F21())).Cardinality()))).Uint32(), (_this).F11()), _dafny.SeqOf((_482_v4).F21())))
			_ = _rhs91
			var _rhs92 _dafny.Int = Companion_Default___.SafeModuloInt((_482_v4).F21(), ((_533_v39).Cardinality()).Plus((_dafny.Zero).Minus((_482_v4).F21())))
			_ = _rhs92
			var _rhs93 _dafny.Map = _531_v36
			_ = _rhs93
			var _rhs94 bool = !((_534_v40).Merge(_534_v40)).Contains(_478_v0)
			_ = _rhs94
			var _rhs95 bool = ((_482_v4).F20()).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfInt64(-724)), _dafny.IntOfUint32(((_482_v4).F20()).Cardinality()))).Uint32()).(bool)
			_ = _rhs95
			var _lhs80 *GlobalState = globalState
			_ = _lhs80
			var _lhs81 *GlobalState = globalState
			_ = _lhs81
			var _lhs82 *GlobalState = globalState
			_ = _lhs82
			_lhs80.F7 = _rhs91
			_lhs81.F1 = _rhs92
			_531_v36 = _rhs93
			_lhs82.F6 = _rhs94
			r0 = _rhs95
			_526_cf14 = (_482_v4).F21()
		} else {
			var _536___mcc_h4 _dafny.Map = _source9.Get_().(D6_DC9).Cf12
			_ = _536___mcc_h4
			var _537_cf12 _dafny.Map = _536___mcc_h4
			_ = _537_cf12
			if (_478_v0) || (((_this).F11()).Cmp((_482_v4).F21()) <= 0) {
				(globalState).F1 = (_482_v4).F21()
				(globalState).F7 = false
				var _538_v41 _dafny.Array
				_ = _538_v41
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_9
				var _nw79 _dafny.Array
				_ = _nw79
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw79 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.Int = func(_539_i3 _dafny.Int) _dafny.Int {
						return (_539_i3).Times((_this).F11())
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
				_538_v41 = _nw79
				var _540_v42 _dafny.Sequence
				_ = _540_v42
				_540_v42 = _dafny.SeqOf(_538_v41, _538_v41, _538_v41)
				var _541_v43 _dafny.Sequence
				_ = _541_v43
				_541_v43 = _dafny.SeqOf((_540_v42).Select((Companion_Default___.SafeIndex((_482_v4).F21(), _dafny.IntOfUint32((_540_v42).Cardinality()))).Uint32()).(_dafny.Array), _538_v41, _538_v41)
				var _542_v44 _dafny.Sequence
				_ = _542_v44
				_542_v44 = _dafny.UnicodeSeqOfUtf8Bytes("miujifbj")
				var _543_v45 _dafny.Sequence
				_ = _543_v45
				_543_v45 = _dafny.SeqOf((_482_v4).F21())
				var _544_v46 _dafny.Sequence
				_ = _544_v46
				_544_v46 = _dafny.SeqOf(_543_v45)
				var _545_v47 D9
				_ = _545_v47
				_545_v47 = Companion_D9_.Create_DC19_()
				var _546_v48 _dafny.Map
				_ = _546_v48
				_546_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_482_v4).F21()), _479_v1)
				var _547_v50 _dafny.Map
				_ = _547_v50
				_547_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qfr")).Cardinality()), _dafny.SetOf((_this).F11()))
				var _548_v52 _dafny.Set
				_ = _548_v52
				_548_v52 = _dafny.SetOf((_482_v4).F21(), (_this).F11(), (_482_v4).F21())
				var _549_v53 _dafny.Array
				_ = _549_v53
				var _nwElement0_28 bool = (func() bool {
					if !(_478_v0) {
						return _478_v0
					}
					return _478_v0
				})()
				_ = _nwElement0_28
				var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(20))
				_ = _nw80
				(_nw80).ArraySet1(_nwElement0_28, 0)
				(_nw80).ArraySet1(_478_v0, 1)
				(_nw80).ArraySet1(_dafny.Companion_Sequence_.Equal(_544_v46, _544_v46), 2)
				(_nw80).ArraySet1(_478_v0, 3)
				(_nw80).ArraySet1((_482_v4).Fm21(_478_v0, (_482_v4).F21(), globalState), 4)
				(_nw80).ArraySet1(_478_v0, 5)
				(_nw80).ArraySet1((Companion_Default___.Fm30((_482_v4).F21(), false, (_482_v4).F21(), _dafny.IntOfInt64(930), globalState)).Contains((_482_v4).F21()), 6)
				(_nw80).ArraySet1(true, 7)
				(_nw80).ArraySet1(_478_v0, 8)
				(_nw80).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_545_v47), _545_v47), 9)
				(_nw80).ArraySet1(_478_v0, 10)
				(_nw80).ArraySet1((_this).Fm4((_this).F11(), _546_v48, (func() _dafny.Map {
					var _coll38 = _dafny.NewMapBuilder()
					_ = _coll38
					for _iter39 := _dafny.Iterate((Companion_Default___.Fm31((_this).F24(), globalState)).Elements()); ; {
						_compr_38, _ok39 := _iter39()
						if !_ok39 {
							break
						}
						var _550_v49 _dafny.Sequence
						_550_v49 = interface{}(_compr_38).(_dafny.Sequence)
						if (Companion_Default___.Fm31((_this).F24(), globalState)).Contains(_550_v49) {
							_coll38.Add(_550_v49, _478_v0)
						}
					}
					return _coll38.ToMap()
				}()).Cardinality(), false, globalState), 11)
				(_nw80).ArraySet1((func() bool {
					if _478_v0 {
						return (_481_v3).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_481_v3).Cardinality()))).Uint32()).(bool)
					}
					return _478_v0
				})(), 12)
				(_nw80).ArraySet1(false, 13)
				(_nw80).ArraySet1(!((_548_v52).IsSubsetOf((func() _dafny.Set {
					if (_547_v50).Contains((_482_v4).F21()) {
						return (_547_v50).Get((_482_v4).F21()).(_dafny.Set)
					}
					return func() _dafny.Set {
						var _coll39 = _dafny.NewBuilder()
						_ = _coll39
						for _iter40 := _dafny.Iterate((_dafny.SeqOf((_482_v4).F21())).Elements()); ; {
							_compr_39, _ok40 := _iter40()
							if !_ok40 {
								break
							}
							var _551_v51 _dafny.Int
							_551_v51 = interface{}(_compr_39).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_482_v4).F21()), _551_v51) {
								_coll39.Add((_551_v51).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("knsv")).Cardinality())))
							}
						}
						return _coll39.ToSet()
					}()
				})())), 14)
				(_nw80).ArraySet1(_478_v0, 15)
				(_nw80).ArraySet1(false, 16)
				(_nw80).ArraySet1(_478_v0, 17)
				(_nw80).ArraySet1(((_482_v4).F20()).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F11()), _dafny.IntOfUint32(((_482_v4).F20()).Cardinality()))).Uint32()).(bool), 18)
				(_nw80).ArraySet1(_478_v0, 19)
				_549_v53 = _nw80
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))
				_ = _index68
				(_549_v53).ArraySet1(!(_478_v0), (_index68).Int())
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))
				_ = _index69
				var _rhs96 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_540_v42, _541_v43), _541_v43)
				_ = _rhs96
				var _rhs97 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(globalState), (func() _dafny.Sequence {
					if !(true) {
						return _542_v44
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("vtq")
				})())
				_ = _rhs97
				var _rhs98 bool = true
				_ = _rhs98
				var _rhs99 _dafny.Int = (_482_v4).F21()
				_ = _rhs99
				var _lhs83 _dafny.Array = _549_v53
				_ = _lhs83
				var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))
				_ = _lhs84
				var _lhs85 *GlobalState = globalState
				_ = _lhs85
				_541_v43 = _rhs96
				_542_v44 = _rhs97
				(_lhs83).ArraySet1(_rhs98, (_lhs84).Int())
				_lhs85.F1 = _rhs99
				var _552_v54 _dafny.Map
				_ = _552_v54
				_552_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_549_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))).Int()).(bool), (_479_v1).Cardinality())
				var _553_v55 _dafny.Map
				_ = _553_v55
				_553_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_482_v4).F21()).Times(_dafny.IntOfUint32((_542_v44).Cardinality())), _552_v54)
				var _554_v57 _dafny.MultiSet
				_ = _554_v57
				_554_v57 = _dafny.MultiSetOf((_482_v4).F21())
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))
				_ = _index70
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))
				_ = _index71
				var _rhs100 bool = (func() bool {
					if (_this).Fm4((_482_v4).F21(), _546_v48, (_this).F11(), (_this).Fm4(_dafny.IntOfInt64(635), func() _dafny.Map {
						var _coll40 = _dafny.NewMapBuilder()
						_ = _coll40
						for _iter41 := _dafny.Iterate((_dafny.SetOf(_554_v57)).Elements()); ; {
							_compr_40, _ok41 := _iter41()
							if !_ok41 {
								break
							}
							var _555_v56 _dafny.MultiSet
							_555_v56 = interface{}(_compr_40).(_dafny.MultiSet)
							if (_dafny.SetOf(_554_v57)).Contains(_555_v56) {
								_coll40.Add(_555_v56, _479_v1)
							}
						}
						return _coll40.ToMap()
					}(), (_482_v4).F21(), (_549_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))).Int()).(bool), globalState), globalState) {
						return (_549_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))).Int()).(bool)
					}
					return (_549_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))).Int()).(bool)
				})()
				_ = _rhs100
				var _rhs101 bool = !(!(_483_v5).Equals(_537_cf12)) || ((func() bool {
					if _478_v0 {
						return _478_v0
					}
					return (_549_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))).Int()).(bool)
				})())
				_ = _rhs101
				var _rhs102 bool = _478_v0
				_ = _rhs102
				var _rhs103 _dafny.Map = (Companion_Default___.Fm32((_482_v4).F21(), (_482_v4).F21(), (_482_v4).F20(), (_482_v4).F21(), globalState)).Merge(_553_v55)
				_ = _rhs103
				var _rhs104 bool = (_549_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))).Int()).(bool)
				_ = _rhs104
				var _lhs86 _dafny.Array = _549_v53
				_ = _lhs86
				var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))
				_ = _lhs87
				var _lhs88 _dafny.Array = _549_v53
				_ = _lhs88
				var _lhs89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_549_v53), 0))
				_ = _lhs89
				r2 = _rhs100
				r0 = _rhs101
				(_lhs86).ArraySet1(_rhs102, (_lhs87).Int())
				_553_v55 = _rhs103
				(_lhs88).ArraySet1(_rhs104, (_lhs89).Int())
				var _556_v58 _dafny.Array
				_ = _556_v58
				var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(6))
				_ = _nw81
				_556_v58 = _nw81
				var _557_v59 _dafny.Map
				_ = _557_v59
				_557_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_556_v58, _549_v53)
				_557_v59 = (_557_v59).Update(_556_v58, _549_v53)
			} else {
				(globalState).F6 = (!(_478_v0)) && ((_dafny.MultiSetOf((_482_v4).F21())).Contains((_482_v4).F21()))
				var _558_v60 _dafny.Map
				_ = _558_v60
				_558_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_478_v0, _dafny.SeqOf((_482_v4).F21(), (_482_v4).F21()))
				var _559_v61 _dafny.Sequence
				_ = _559_v61
				_559_v61 = _dafny.SeqOf(_dafny.CodePoint('h'))
				var _560_v62 _dafny.Sequence
				_ = _560_v62
				_560_v62 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(442))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}(func(_561_i4 _dafny.Int) _dafny.CodePoint {
					return (_this).F24()
				})))
				var _562_v63 D7
				_ = _562_v63
				_562_v63 = Companion_D7_.Create_DC12_(_dafny.SeqOf(_dafny.IntOfUint32((_559_v61).Cardinality()), _dafny.IntOfUint32(((_560_v62).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_560_v62).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())))
				var _563_v64 _dafny.Sequence
				_ = _563_v64
				_563_v64 = _dafny.SeqOf((_this).F11(), _dafny.IntOfUint32((_559_v61).Cardinality()), (_482_v4).F21())
				var _pat_let_tv26 = _563_v64
				_ = _pat_let_tv26
				_558_v60 = (_558_v60).Update(Companion_Default___.Fm0(globalState), _dafny.Companion_Sequence_.Concatenate((func(_pat_let12_0 D7) D7 {
					return func(_564_dt__update__tmp_h0 D7) D7 {
						return func(_pat_let13_0 _dafny.Sequence) D7 {
							return func(_565_dt__update_hcf17_h0 _dafny.Sequence) D7 {
								return Companion_D7_.Create_DC12_(_565_dt__update_hcf17_h0)
							}(_pat_let13_0)
						}(_pat_let_tv26)
					}(_pat_let12_0)
				}(_562_v63)).Dtor_cf17(), _563_v64))
				var _566_v66 _dafny.MultiSet
				_ = _566_v66
				_566_v66 = _dafny.MultiSetOf(func() _dafny.Set {
					var _coll41 = _dafny.NewBuilder()
					_ = _coll41
					for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(891), _dafny.IntOfInt64(318))); ; {
						_compr_41, _ok42 := _iter42()
						if !_ok42 {
							break
						}
						var _567_v65 _dafny.Int
						_567_v65 = interface{}(_compr_41).(_dafny.Int)
						if ((_dafny.IntOfInt64(891)).Cmp(_567_v65) <= 0) && ((_567_v65).Cmp(_dafny.IntOfInt64(318)) < 0) {
							_coll41.Add((_567_v65).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(338))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg57 _dafny.Int) interface{} {
									return coer57(arg57)
								}
							}((func(_568_v4 *C0) func(_dafny.Int) _dafny.Int {
								return func(_569_i5 _dafny.Int) _dafny.Int {
									return (_568_v4).F21()
								}
							})(_482_v4)))).Cardinality())))
						}
					}
					return _coll41.ToSet()
				}())
				_566_v66 = (_566_v66).Difference(_566_v66)
				var _570_v67 _dafny.MultiSet
				_ = _570_v67
				_570_v67 = _dafny.MultiSetOf((_482_v4).F21(), (_482_v4).F21(), (_482_v4).F21())
				_570_v67 = _570_v67
				(globalState).F1 = ((_482_v4).F21()).Minus((_482_v4).F21())
			}
			var _571_v68 _dafny.Map
			_ = _571_v68
			_571_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_478_v0)), (_this).F11())
			var _572_v69 _dafny.Set
			_ = _572_v69
			_572_v69 = _dafny.SetOf((_this).Fm3(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("sbmbrcsg"), (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sbmbrcsg")).Cardinality()))).Uint32(), (_this).F24())).Cardinality()), _478_v0, _478_v0, globalState), (_482_v4).F21(), Companion_Default___.Fm2((_482_v4).F21(), (_this).F11(), globalState), (_571_v68).Cardinality())
			_572_v69 = _572_v69
			var _573_v70 _dafny.CodePoint
			_ = _573_v70
			_573_v70 = _dafny.CodePoint('d')
			_573_v70 = (func() _dafny.CodePoint {
				if _478_v0 {
					return (_this).F24()
				}
				return _573_v70
			})()
			var _574_v71 D9
			_ = _574_v71
			_574_v71 = Companion_D9_.Create_DC19_()
			_574_v71 = _574_v71
		}
		var _575_v72 _dafny.Sequence
		_ = _575_v72
		_575_v72 = _dafny.UnicodeSeqOfUtf8Bytes("e")
		var _576_v73 _dafny.Map
		_ = _576_v73
		_576_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_575_v72, _478_v0)
		r2 = (_576_v73).Equals(_576_v73)
		var _577_v74 _dafny.MultiSet
		_ = _577_v74
		_577_v74 = _dafny.MultiSetOf((_this).F11(), (_482_v4).F21())
		var _578_v76 _dafny.Map
		_ = _578_v76
		_578_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_577_v74, func() _dafny.Map {
			var _coll42 = _dafny.NewMapBuilder()
			_ = _coll42
			for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(679), _dafny.IntOfInt64(599))); ; {
				_compr_42, _ok43 := _iter43()
				if !_ok43 {
					break
				}
				var _579_v75 _dafny.Int
				_579_v75 = interface{}(_compr_42).(_dafny.Int)
				if ((_dafny.IntOfInt64(679)).Cmp(_579_v75) <= 0) && ((_579_v75).Cmp(_dafny.IntOfInt64(599)) < 0) {
					_coll42.Add(Companion_Default___.SafeDivisionInt(_579_v75, (Companion_D7_.Create_DC13_(_dafny.SetOf(_478_v0), (_this).F11(), (_482_v4).F21())).Dtor_cf20()), _dafny.IntOfUint32((_575_v72).Cardinality()))
				}
			}
			return _coll42.ToMap()
		}())
		var _580_v77 D0
		_ = _580_v77
		_580_v77 = Companion_D0_.Create_DC1_(_478_v0, (_this).F11(), _478_v0)
		r2 = !((_this).Fm4(_dafny.IntOfInt64(-452), _578_v76, (_482_v4).F21(), (_580_v77).Dtor_cf1(), globalState))
		var _581_v78 _dafny.Set
		_ = _581_v78
		_581_v78 = _dafny.SetOf(false, _478_v0)
		var _582_v79 _dafny.Map
		_ = _582_v79
		_582_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _581_v78)
		var _583_v80 D7
		_ = _583_v80
		_583_v80 = Companion_D7_.Create_DC13_((func() _dafny.Set {
			if (_582_v79).Contains((_this).F24()) {
				return (_582_v79).Get((_this).F24()).(_dafny.Set)
			}
			return _581_v78
		})(), (_482_v4).F21(), (_this).F11())
		var _584_v81 _dafny.Map
		_ = _584_v81
		_584_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _478_v0)
		_583_v80 = Companion_D7_.Create_DC13_((_dafny.SetOf(_478_v0, (func() bool {
			if (_584_v81).Contains((_this).F24()) {
				return (_584_v81).Get((_this).F24()).(bool)
			}
			return _478_v0
		})())).Union(_581_v78), (_482_v4).F21(), _dafny.IntOfInt64(71))
		var _585_v82 _dafny.Sequence
		_ = _585_v82
		_585_v82 = _dafny.SeqOf((_this).F11())
		var _586_v83 _dafny.Sequence
		_ = _586_v83
		_586_v83 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_585_v82, _478_v0)).Cardinality())
		var _587_v84 _dafny.Array
		_ = _587_v84
		var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
		_ = _nw82
		_587_v84 = _nw82
		var _588_v85 _dafny.Map
		_ = _588_v85
		_588_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _587_v84)
		r0 = (_482_v4).Fm21((_dafny.MultiSetOf((_this).F11())).IsProperSubsetOf(_577_v74), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_585_v82).Cardinality()), (_588_v85).Cardinality()), globalState)
		r1 = (func() bool {
			if (_576_v73).Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xbqkmva"), _dafny.UnicodeSeqOfUtf8Bytes("bnu"))) {
				return (_576_v73).Get(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xbqkmva"), _dafny.UnicodeSeqOfUtf8Bytes("bnu"))).(bool)
			}
			return _478_v0
		})()
		var _589_v86 _dafny.Sequence
		_ = _589_v86
		_589_v86 = _dafny.SeqOf(_481_v3)
		var _590_v87 _dafny.Sequence
		_ = _590_v87
		_590_v87 = _dafny.SeqOf(_589_v86, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(308))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg58 _dafny.Int) interface{} {
				return coer58(arg58)
			}
		}((func(_591_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_592_i6 _dafny.Int) _dafny.Sequence {
				return _591_v3
			}
		})(_481_v3))), _589_v86, _589_v86)
		r2 = (((_this).F11()).Times((_482_v4).F21())).Cmp(_dafny.IntOfUint32(((_590_v87).Select((Companion_Default___.SafeIndex((_482_v4).F21(), _dafny.IntOfUint32((_590_v87).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())) < 0
		r3 = ((_482_v4).F21()).Minus((_this).F11())
		return r0, r1, r2, r3
	}
}
func (_this *C2) M9(p0 bool, p1 _dafny.Array, globalState *GlobalState) (D7, _dafny.Array) {
	{
		var r0 D7 = Companion_D7_.Default()
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _593_i0 _dafny.Int
		_ = _593_i0
		_593_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_593_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_593_i0 = (_593_i0).Plus(_dafny.One)
					var _594_v0 D0
					_ = _594_v0
					_594_v0 = Companion_D0_.Create_DC1_(p0, (_this).F11(), Companion_Default___.Fm0(globalState))
					var _pat_let_tv27 = p0
					_ = _pat_let_tv27
					var _pat_let_tv28 = p0
					_ = _pat_let_tv28
					var _pat_let_tv29 = p0
					_ = _pat_let_tv29
					var _source10 D0 = func(_pat_let14_0 D0) D0 {
						return func(_595_dt__update__tmp_h0 D0) D0 {
							return func(_pat_let15_0 bool) D0 {
								return func(_596_dt__update_hcf1_h0 bool) D0 {
									return Companion_D0_.Create_DC1_(_596_dt__update_hcf1_h0, (_595_dt__update__tmp_h0).Dtor_cf2(), (_595_dt__update__tmp_h0).Dtor_cf3())
								}(_pat_let15_0)
							}((func() bool {
								if _pat_let_tv27 {
									return _pat_let_tv28
								}
								return _pat_let_tv29
							})())
						}(_pat_let14_0)
					}(_594_v0)
					_ = _source10
					if _source10.Is_DC1() {
						var _597___mcc_h0 bool = _source10.Get_().(D0_DC1).Cf1
						_ = _597___mcc_h0
						var _598___mcc_h1 _dafny.Int = _source10.Get_().(D0_DC1).Cf2
						_ = _598___mcc_h1
						var _599___mcc_h2 bool = _source10.Get_().(D0_DC1).Cf3
						_ = _599___mcc_h2
						var _600_cf3 bool = _599___mcc_h2
						_ = _600_cf3
						var _601_cf2 _dafny.Int = _598___mcc_h1
						_ = _601_cf2
						var _602_cf1 bool = _597___mcc_h0
						_ = _602_cf1
						(globalState).F6 = false
						var _603_v1 bool
						_ = _603_v1
						var _604_v2 bool
						_ = _604_v2
						var _605_v3 bool
						_ = _605_v3
						var _606_v4 _dafny.Int
						_ = _606_v4
						var _out16 bool
						_ = _out16
						var _out17 bool
						_ = _out17
						var _out18 bool
						_ = _out18
						var _out19 _dafny.Int
						_ = _out19
						_out16, _out17, _out18, _out19 = (_this).M8(globalState)
						_603_v1 = _out16
						_604_v2 = _out17
						_605_v3 = _out18
						_606_v4 = _out19
						var _607_v5 _dafny.Array
						_ = _607_v5
						var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
						_ = _nw83
						_607_v5 = _nw83
						_607_v5 = _607_v5
						var _608_v6 D9
						_ = _608_v6
						_608_v6 = Companion_D9_.Create_DC19_()
						_608_v6 = _608_v6
					} else {
						var _609___mcc_h3 _dafny.Int = _source10.Get_().(D0_DC0).Cf0
						_ = _609___mcc_h3
						var _610_cf0 _dafny.Int = _609___mcc_h3
						_ = _610_cf0
						var _611_v7 _dafny.Map
						_ = _611_v7
						_611_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
						(globalState).F6 = (((_611_v7).Merge(_611_v7)).Cardinality()).Cmp(_610_cf0) > 0
						var _612_v8 _dafny.Map
						_ = _612_v8
						_612_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_610_cf0, _610_cf0)
						var _613_v9 _dafny.Sequence
						_ = _613_v9
						_613_v9 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-762))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg59 _dafny.Int) interface{} {
								return coer59(arg59)
							}
						}((func(_614_v8 _dafny.Map) func(_dafny.Int) _dafny.Map {
							return func(_615_i1 _dafny.Int) _dafny.Map {
								return _614_v8
							}
						})(_612_v8))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-221))).Uint32(), func(coer60 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg60 _dafny.Int) interface{} {
								return coer60(arg60)
							}
						}((func(_616_v8 _dafny.Map) func(_dafny.Int) _dafny.Map {
							return func(_617_i2 _dafny.Int) _dafny.Map {
								return _616_v8
							}
						})(_612_v8))))
						var _618_v10 _dafny.Array
						_ = _618_v10
						var _nwElement0_29 _dafny.Int = _610_cf0
						_ = _nwElement0_29
						var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(11))
						_ = _nw84
						(_nw84).ArraySet1(_nwElement0_29, 0)
						(_nw84).ArraySet1((_this).F11(), 1)
						(_nw84).ArraySet1((_this).F11(), 2)
						(_nw84).ArraySet1(((_612_v8).Cardinality()).Plus(_dafny.IntOfUint32(((_613_v9).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_613_v9).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())), 3)
						(_nw84).ArraySet1(_610_cf0, 4)
						(_nw84).ArraySet1((_this).F11(), 5)
						(_nw84).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F11(), (_this).F11()), 6)
						(_nw84).ArraySet1((_dafny.Zero).Minus(_610_cf0), 7)
						(_nw84).ArraySet1((_dafny.SetOf(!(p0))).Cardinality(), 8)
						(_nw84).ArraySet1(_610_cf0, 9)
						(_nw84).ArraySet1((_this).F11(), 10)
						_618_v10 = _nw84
						var _619_v11 _dafny.Set
						_ = _619_v11
						_619_v11 = _dafny.SetOf(_610_cf0)
						var _620_v12 _dafny.Map
						_ = _620_v12
						_620_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_610_cf0, _619_v11)
						var _rhs105 _dafny.Array = _618_v10
						_ = _rhs105
						var _rhs106 _dafny.Int = ((_dafny.Zero).Minus(_610_cf0)).Minus((_this).F11())
						_ = _rhs106
						var _rhs107 bool = p0
						_ = _rhs107
						var _rhs108 bool = (((func() _dafny.Set {
							if (_620_v12).Contains(_610_cf0) {
								return (_620_v12).Get(_610_cf0).(_dafny.Set)
							}
							return _dafny.SetOf(_610_cf0, (_this).F11())
						})()).Cardinality()).Cmp((_this).F11()) != 0
						_ = _rhs108
						var _lhs90 *GlobalState = globalState
						_ = _lhs90
						var _lhs91 *GlobalState = globalState
						_ = _lhs91
						var _lhs92 *GlobalState = globalState
						_ = _lhs92
						_618_v10 = _rhs105
						_lhs90.F1 = _rhs106
						_lhs91.F7 = _rhs107
						_lhs92.F7 = _rhs108
						var _621_v13 _dafny.MultiSet
						_ = _621_v13
						_621_v13 = _dafny.MultiSetOf(_610_cf0)
						var _622_v14 _dafny.Map
						_ = _622_v14
						_622_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_this).F11())).Union(_621_v13), p0)
						(globalState).F7 = (func() bool {
							if (_622_v14).Contains((_621_v13).Difference(_dafny.MultiSetOf((_this).F11()))) {
								return (_622_v14).Get((_621_v13).Difference(_dafny.MultiSetOf((_this).F11()))).(bool)
							}
							return p0
						})()
						var _623_v15 _dafny.Sequence
						_ = _623_v15
						_623_v15 = _dafny.SeqOf(p0, p0)
						var _624_v16 _dafny.Map
						_ = _624_v16
						_624_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, false)
						var _625_v17 *C0
						_ = _625_v17
						var _nw85 *C0 = New_C0_()
						_ = _nw85
						_nw85.Ctor__(_623_v15, Companion_Default___.Fm2(_610_cf0, (_624_v16).Cardinality(), globalState))
						_625_v17 = _nw85
						var _626_v18 _dafny.Map
						_ = _626_v18
						_626_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _625_v17)
						var _627_v19 _dafny.Map
						_ = _627_v19
						_627_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() *C0 {
							if (_626_v18).Contains(p0) {
								return (_626_v18).Get(p0).(*C0)
							}
							return _625_v17
						})(), _610_cf0)
						(globalState).F1 = (func() _dafny.Int {
							if (_627_v19).Contains(_625_v17) {
								return (_627_v19).Get(_625_v17).(_dafny.Int)
							}
							return _dafny.IntOfInt64(893)
						})()
					}
					var _628_v20 _dafny.Array
					_ = _628_v20
					var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
					_ = _nw86
					_628_v20 = _nw86
					var _629_v21 _dafny.Array
					_ = _629_v21
					var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
					_ = _nw87
					_629_v21 = _nw87
					var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_628_v20), 0))
					_ = _index72
					(_628_v20).ArraySet1(_629_v21, (_index72).Int())
					var _630_v22 _dafny.Sequence
					_ = _630_v22
					_630_v22 = _dafny.SeqOf(_629_v21, _629_v21)
					var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_628_v20), 0))
					_ = _index73
					(_628_v20).ArraySet1((_630_v22).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_630_v22).Cardinality()))).Uint32()).(_dafny.Array), (_index73).Int())
					var _631_v23 _dafny.Sequence
					_ = _631_v23
					_631_v23 = _dafny.UnicodeSeqOfUtf8Bytes("vb")
					var _632_v24 _dafny.Sequence
					_ = _632_v24
					_632_v24 = _dafny.SeqOf(p0)
					var _633_v25 _dafny.MultiSet
					_ = _633_v25
					_633_v25 = _dafny.MultiSetOf((func() _dafny.Sequence {
						if p0 {
							return _632_v24
						}
						return _632_v24
					})(), _632_v24, _dafny.Companion_Sequence_.Concatenate(_632_v24, _632_v24), _632_v24)
					var _634_v26 _dafny.MultiSet
					_ = _634_v26
					_634_v26 = _dafny.MultiSetOf(p0)
					var _635_v27 _dafny.Map
					_ = _635_v27
					_635_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _634_v26)
					var _636_v29 _dafny.Map
					_ = _636_v29
					_636_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F11())
					var _637_v30 _dafny.Map
					_ = _637_v30
					_637_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((func() _dafny.Set {
						var _coll43 = _dafny.NewBuilder()
						_ = _coll43
						for _iter44 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_631_v23, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_631_v23).Cardinality()))).Uint32(), (_this).F24())).Elements()); ; {
							_compr_43, _ok44 := _iter44()
							if !_ok44 {
								break
							}
							var _638_v28 _dafny.CodePoint
							_638_v28 = interface{}(_compr_43).(_dafny.CodePoint)
							if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_631_v23, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_631_v23).Cardinality()))).Uint32(), (_this).F24()), _638_v28) {
								_coll43.Add(_638_v28)
							}
						}
						return _coll43.ToSet()
					}()).Cardinality()), _636_v29)
					var _rhs109 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_631_v23, (Companion_Default___.SafeIndex(((_635_v27).Update((_this).F11(), _dafny.MultiSetOf(p0, p0))).Cardinality(), _dafny.IntOfUint32((_631_v23).Cardinality()))).Uint32(), (_this).F24()), _dafny.UnicodeSeqOfUtf8Bytes("vjdgtkgo")), _dafny.Companion_Sequence_.Concatenate(_631_v23, _631_v23))
					_ = _rhs109
					var _rhs110 bool = !(true) || (Companion_Default___.Fm0(globalState))
					_ = _rhs110
					var _rhs111 _dafny.MultiSet = (func() _dafny.MultiSet {
						if (p0) && (p0) {
							return _633_v25
						}
						return (_633_v25).Difference(_dafny.MultiSetOf(_632_v24, _632_v24, _632_v24, _632_v24, _632_v24))
					})()
					_ = _rhs111
					var _rhs112 _dafny.Int = (func() _dafny.Int {
						if (_634_v26).Contains((_this).Fm4((_this).F11(), _637_v30, (_this).F11(), p0, globalState)) {
							return (_634_v26).Multiplicity((_this).Fm4((_this).F11(), _637_v30, (_this).F11(), p0, globalState))
						}
						return (_this).F11()
					})()
					_ = _rhs112
					var _lhs93 *GlobalState = globalState
					_ = _lhs93
					var _lhs94 *GlobalState = globalState
					_ = _lhs94
					_631_v23 = _rhs109
					_lhs93.F7 = _rhs110
					_633_v25 = _rhs111
					_lhs94.F1 = _rhs112
					(globalState).F7 = p0
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _639_v31 _dafny.MultiSet
		_ = _639_v31
		_639_v31 = _dafny.MultiSetOf((_this).F11())
		var _640_v32 _dafny.Sequence
		_ = _640_v32
		_640_v32 = _dafny.SeqOf((_this).F11(), (_this).F11())
		var _hi0 _dafny.Int = _dafny.IntOfInt64(770)
		_ = _hi0
		for _641_i3 := ((_639_v31).Difference(_dafny.MultiSetFromSeq(_640_v32))).Cardinality(); _641_i3.Cmp(_hi0) < 0; _641_i3 = _641_i3.Plus(_dafny.One) {
			(globalState).F6 = p0
			var _642_v33 D8
			_ = _642_v33
			_642_v33 = Companion_D8_.Create_DC16_(p0, _641_i3)
			var _source11 D8 = _642_v33
			_ = _source11
			if _source11.Is_DC16() {
				var _643___mcc_h4 bool = _source11.Get_().(D8_DC16).Cf23
				_ = _643___mcc_h4
				var _644___mcc_h5 _dafny.Int = _source11.Get_().(D8_DC16).Cf24
				_ = _644___mcc_h5
				var _645_cf24 _dafny.Int = _644___mcc_h5
				_ = _645_cf24
				var _646_cf23 bool = _643___mcc_h4
				_ = _646_cf23
				var _647_v34 _dafny.Array
				_ = _647_v34
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_10
				var _nw88 _dafny.Array
				_ = _nw88
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw88 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Int = func(_648_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_648_i4, (_this).F11())
					}
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw88 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw88).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw88).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_647_v34 = _nw88
				var _649_v35 _dafny.Map
				_ = _649_v35
				_649_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_647_v34, p0)
				var _650_v36 _dafny.Sequence
				_ = _650_v36
				_650_v36 = _dafny.SeqOf(p0, _646_cf23)
				var _651_v37 _dafny.MultiSet
				_ = _651_v37
				_651_v37 = _dafny.MultiSetOf((_650_v36).Select((Companion_Default___.SafeIndex(_645_cf24, _dafny.IntOfUint32((_650_v36).Cardinality()))).Uint32()).(bool), !(p0))
				var _652_v38 _dafny.MultiSet
				_ = _652_v38
				_652_v38 = _651_v37
				var _653_v39 _dafny.Set
				_ = _653_v39
				_653_v39 = _dafny.SetOf(p0, p0)
				var _654_v40 _dafny.Map
				_ = _654_v40
				_654_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(_646_cf23, Companion_Default___.Fm0(globalState), _646_cf23, p0, p0)).IsProperSubsetOf(_653_v39))
				var _655_v41 D10
				_ = _655_v41
				_655_v41 = Companion_D10_.Create_DC20_(_649_v35)
				var _rhs113 bool = (func() bool {
					if (_654_v40).Contains(_646_cf23) {
						return (_654_v40).Get(_646_cf23).(bool)
					}
					return false
				})()
				_ = _rhs113
				var _rhs114 _dafny.Map = (_649_v35).Merge((_655_v41).Dtor_cf27())
				_ = _rhs114
				var _rhs115 _dafny.Int = (_this).F11()
				_ = _rhs115
				var _rhs116 _dafny.MultiSet = _652_v38
				_ = _rhs116
				var _lhs95 *GlobalState = globalState
				_ = _lhs95
				var _lhs96 *GlobalState = globalState
				_ = _lhs96
				_lhs95.F6 = _rhs113
				_649_v35 = _rhs114
				_lhs96.F1 = _rhs115
				_652_v38 = _rhs116
				var _656_v42 _dafny.Array
				_ = _656_v42
				var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
				_ = _nw89
				_656_v42 = _nw89
				var _657_v43 _dafny.Array
				_ = _657_v43
				_657_v43 = _656_v42
				var _658_v44 _dafny.Array
				_ = _658_v44
				var _nwElement0_30 _dafny.Array = _656_v42
				_ = _nwElement0_30
				var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(12))
				_ = _nw90
				(_nw90).ArraySet1(_nwElement0_30, 0)
				(_nw90).ArraySet1(_657_v43, 1)
				(_nw90).ArraySet1(_657_v43, 2)
				(_nw90).ArraySet1(_657_v43, 3)
				(_nw90).ArraySet1(_657_v43, 4)
				(_nw90).ArraySet1(_657_v43, 5)
				(_nw90).ArraySet1(_657_v43, 6)
				(_nw90).ArraySet1(_657_v43, 7)
				(_nw90).ArraySet1(_657_v43, 8)
				(_nw90).ArraySet1(_656_v42, 9)
				(_nw90).ArraySet1(_657_v43, 10)
				(_nw90).ArraySet1(_657_v43, 11)
				_658_v44 = _nw90
				var _659_v45 _dafny.Map
				_ = _659_v45
				_659_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_658_v44, _647_v34)
				var _660_v46 _dafny.Map
				_ = _660_v46
				_660_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_659_v45, p0)
				var _661_v47 D11
				_ = _661_v47
				_661_v47 = Companion_D11_.Create_DC22_((_659_v45).Update(_658_v44, _647_v34))
				_660_v46 = (_660_v46).Update((_659_v45).Merge((_661_v47).Dtor_cf30()), p0)
				var _662_v48 _dafny.Map
				_ = _662_v48
				_662_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_641_i3).Minus(_645_cf24))
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_647_v34), 0))
				_ = _index74
				(_647_v34).ArraySet1(_dafny.IntOfInt64(-915), (_index74).Int())
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_647_v34), 0))
				_ = _index75
				var _rhs117 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfInt64(196))
				_ = _rhs117
				var _rhs118 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(414))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}(func(_663_i5 _dafny.Int) _dafny.CodePoint {
					return (_this).F24()
				}))).Cardinality()))
				_ = _rhs118
				var _rhs119 _dafny.Int = (_dafny.Zero).Minus(((_this).F11()).Times(_dafny.IntOfInt64(630)))
				_ = _rhs119
				var _lhs97 _dafny.Array = _647_v34
				_ = _lhs97
				var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_647_v34), 0))
				_ = _lhs98
				_662_v48 = _rhs117
				(_lhs97).ArraySet1(_rhs118, (_lhs98).Int())
				_645_cf24 = _rhs119
				var _664_v49 _dafny.Map
				_ = _664_v49
				_664_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_656_v42, _645_cf24)
				_664_v49 = (_664_v49).Update((func() _dafny.Array {
					if _646_cf23 {
						return _657_v43
					}
					return _657_v43
				})(), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_645_cf24, (_this).Fm3((_this).F11(), _646_cf23, p0, globalState))))
			} else if _source11.Is_DC15() {
				var _665___mcc_h6 _dafny.Map = _source11.Get_().(D8_DC15).Cf22
				_ = _665___mcc_h6
				var _666_cf22 _dafny.Map = _665___mcc_h6
				_ = _666_cf22
				_640_v32 = _640_v32
				var _667_v50 _dafny.Sequence
				_ = _667_v50
				_667_v50 = _dafny.SeqOf(p0, p0, p0, p0)
				var _668_v51 D6
				_ = _668_v51
				_668_v51 = Companion_D6_.Create_DC10_(_dafny.Companion_Sequence_.Concatenate(_667_v50, _667_v50))
				_668_v51 = _668_v51
				var _669_v52 D11
				_ = _669_v52
				_669_v52 = Companion_D11_.Create_DC23_(false, _641_i3, _dafny.IntOfInt64(662))
				var _670_v53 _dafny.Sequence
				_ = _670_v53
				_670_v53 = _dafny.UnicodeSeqOfUtf8Bytes("o")
				var _671_v54 _dafny.MultiSet
				_ = _671_v54
				_671_v54 = _dafny.MultiSetOf(p0, !(p0))
				var _672_v55 _dafny.Array
				_ = _672_v55
				var _nwElement0_31 _dafny.MultiSet = (_639_v31).Union(_639_v31)
				_ = _nwElement0_31
				var _nw91 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(27))
				_ = _nw91
				(_nw91).ArraySet1(_nwElement0_31, 0)
				(_nw91).ArraySet1(_dafny.MultiSetOf((_this).F11()), 1)
				(_nw91).ArraySet1(_639_v31, 2)
				(_nw91).ArraySet1(_639_v31, 3)
				(_nw91).ArraySet1((_639_v31).Intersection(_639_v31), 4)
				(_nw91).ArraySet1(_dafny.MultiSetOf(_641_i3, _dafny.IntOfUint32((_dafny.SeqOf(_641_i3)).Cardinality()), _641_i3, (_this).F11(), (_this).F11()), 5)
				(_nw91).ArraySet1(_639_v31, 6)
				(_nw91).ArraySet1((_dafny.MultiSetOf((_669_v52).Dtor_cf33())).Difference(_639_v31), 7)
				(_nw91).ArraySet1(_639_v31, 8)
				(_nw91).ArraySet1((func() _dafny.MultiSet {
					if p0 {
						return _639_v31
					}
					return _639_v31
				})(), 9)
				(_nw91).ArraySet1(_639_v31, 10)
				(_nw91).ArraySet1(_639_v31, 11)
				(_nw91).ArraySet1(_639_v31, 12)
				(_nw91).ArraySet1(_639_v31, 13)
				(_nw91).ArraySet1((_639_v31).Intersection(_639_v31), 14)
				(_nw91).ArraySet1(_639_v31, 15)
				(_nw91).ArraySet1((_639_v31).Update(Companion_Default___.Fm2(_dafny.IntOfInt64(888), _dafny.IntOfUint32((_670_v53).Cardinality()), globalState), Companion_Default___.Abs((_this).F11())), 16)
				(_nw91).ArraySet1(((_dafny.MultiSetOf(_dafny.IntOfInt64(58), (_this).F11())).Update(_641_i3, Companion_Default___.Abs(_641_i3))).Intersection(_639_v31), 17)
				(_nw91).ArraySet1(_639_v31, 18)
				(_nw91).ArraySet1(_639_v31, 19)
				(_nw91).ArraySet1(_639_v31, 20)
				(_nw91).ArraySet1(_639_v31, 21)
				(_nw91).ArraySet1(_639_v31, 22)
				(_nw91).ArraySet1(((_639_v31).Update((_this).F11(), Companion_Default___.Abs((_671_v54).Cardinality()))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(436))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_673_i3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_674_i6 _dafny.Int) _dafny.Int {
						return _673_i3
					}
				})(_641_i3))))), 23)
				(_nw91).ArraySet1(_639_v31, 24)
				(_nw91).ArraySet1(_dafny.MultiSetOf(_641_i3, _641_i3), 25)
				(_nw91).ArraySet1(_dafny.MultiSetOf((_this).F11()), 26)
				_672_v55 = _nw91
				_672_v55 = _672_v55
				var _675_v56 *C0
				_ = _675_v56
				var _nw92 *C0 = New_C0_()
				_ = _nw92
				_nw92.Ctor__(_667_v50, ((_this).F11()).Minus(_dafny.IntOfInt64(-174)))
				_675_v56 = _nw92
			} else {
				var _676___mcc_h7 D8 = _source11.Get_().(D8_DC17).Cf25
				_ = _676___mcc_h7
				var _677_cf25 D8 = _676___mcc_h7
				_ = _677_cf25
				var _678_v57 _dafny.Array
				_ = _678_v57
				var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
				_ = _nw93
				_678_v57 = _nw93
				var _679_v58 _dafny.MultiSet
				_ = _679_v58
				_679_v58 = _dafny.MultiSetOf((_this).F24())
				var _680_v59 _dafny.Array
				_ = _680_v59
				var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw94
				_680_v59 = _nw94
				var _681_v60 _dafny.Sequence
				_ = _681_v60
				_681_v60 = _dafny.SeqOf(_680_v59)
				var _682_v61 _dafny.Map
				_ = _682_v61
				_682_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_679_v58, (_681_v60).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_681_v60).Cardinality()))).Uint32()).(_dafny.Array))
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_678_v57), 0))
				_ = _index76
				(_678_v57).ArraySet1(_682_v61, (_index76).Int())
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_678_v57), 0))
				_ = _index77
				var _rhs120 _dafny.Map = ((_682_v61).Merge(_682_v61)).Merge((_682_v61).Update(_679_v58, _680_v59))
				_ = _rhs120
				var _rhs121 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_681_v60, _681_v60), _681_v60)).Cardinality())
				_ = _rhs121
				var _rhs122 bool = p0
				_ = _rhs122
				var _rhs123 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_640_v32, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F11()), _640_v32))
				_ = _rhs123
				var _lhs99 _dafny.Array = _678_v57
				_ = _lhs99
				var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_678_v57), 0))
				_ = _lhs100
				var _lhs101 *GlobalState = globalState
				_ = _lhs101
				var _lhs102 *GlobalState = globalState
				_ = _lhs102
				(_lhs99).ArraySet1(_rhs120, (_lhs100).Int())
				_lhs101.F1 = _rhs121
				_lhs102.F7 = _rhs122
				_640_v32 = _rhs123
				var _683_v62 *C1
				_ = _683_v62
				var _nw95 *C1 = New_C1_()
				_ = _nw95
				_nw95.Ctor__(p0, !((Companion_Default___.Fm2((_this).F11(), _641_i3, globalState)).Cmp(_641_i3) > 0))
				_683_v62 = _nw95
				var _684_v63 _dafny.Sequence
				_ = _684_v63
				_684_v63 = _dafny.UnicodeSeqOfUtf8Bytes("fgsyppx")
				var _685_v64 _dafny.MultiSet
				_ = _685_v64
				_685_v64 = _dafny.MultiSetOf(_684_v63, _dafny.UnicodeSeqOfUtf8Bytes("gg"))
				var _686_v65 D2
				_ = _686_v65
				_686_v65 = Companion_D2_.Create_DC3_(_685_v64)
				var _687_v66 _dafny.Sequence
				_ = _687_v66
				_687_v66 = _dafny.SeqOf((_683_v62).F19(), (_683_v62).Fm13(_686_v65, Companion_Default___.Fm0(globalState), (_683_v62).F18(), (_this).F11(), globalState))
				var _688_v67 D6
				_ = _688_v67
				_688_v67 = Companion_D6_.Create_DC10_(_687_v66)
				var _pat_let_tv30 = _687_v66
				_ = _pat_let_tv30
				var _pat_let_tv31 = _683_v62
				_ = _pat_let_tv31
				var _pat_let_tv32 = _683_v62
				_ = _pat_let_tv32
				var _689_v68 _dafny.Array
				_ = _689_v68
				var _nwElement0_32 D6 = _688_v67
				_ = _nwElement0_32
				var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(25))
				_ = _nw96
				(_nw96).ArraySet1(_nwElement0_32, 0)
				(_nw96).ArraySet1(Companion_D6_.Create_DC10_(_687_v66), 1)
				(_nw96).ArraySet1(Companion_D6_.Create_DC10_((_688_v67).Dtor_cf13()), 2)
				(_nw96).ArraySet1(_688_v67, 3)
				(_nw96).ArraySet1(_688_v67, 4)
				(_nw96).ArraySet1(Companion_D6_.Create_DC10_(_687_v66), 5)
				(_nw96).ArraySet1(Companion_Default___.Fm33(p0, globalState), 6)
				(_nw96).ArraySet1(_688_v67, 7)
				(_nw96).ArraySet1(_688_v67, 8)
				(_nw96).ArraySet1(_688_v67, 9)
				(_nw96).ArraySet1(_688_v67, 10)
				(_nw96).ArraySet1(_688_v67, 11)
				(_nw96).ArraySet1(_688_v67, 12)
				(_nw96).ArraySet1(_688_v67, 13)
				(_nw96).ArraySet1(_688_v67, 14)
				(_nw96).ArraySet1(func(_pat_let16_0 D6) D6 {
					return func(_690_dt__update__tmp_h1 D6) D6 {
						return func(_pat_let17_0 _dafny.Sequence) D6 {
							return func(_691_dt__update_hcf13_h0 _dafny.Sequence) D6 {
								return Companion_D6_.Create_DC10_(_691_dt__update_hcf13_h0)
							}(_pat_let17_0)
						}(_pat_let_tv30)
					}(_pat_let16_0)
				}(_688_v67), 15)
				(_nw96).ArraySet1(Companion_D6_.Create_DC10_(_687_v66), 16)
				(_nw96).ArraySet1(_688_v67, 17)
				(_nw96).ArraySet1(Companion_D6_.Create_DC10_(_687_v66), 18)
				(_nw96).ArraySet1(_688_v67, 19)
				(_nw96).ArraySet1(Companion_D6_.Create_DC10_(_dafny.SeqOf((_683_v62).F19(), (_683_v62).F19())), 20)
				(_nw96).ArraySet1(func(_pat_let18_0 D6) D6 {
					return func(_692_dt__update__tmp_h2 D6) D6 {
						return func(_pat_let19_0 _dafny.Sequence) D6 {
							return func(_693_dt__update_hcf13_h1 _dafny.Sequence) D6 {
								return Companion_D6_.Create_DC10_(_693_dt__update_hcf13_h1)
							}(_pat_let19_0)
						}(_dafny.SeqOf((_pat_let_tv31).F19(), !((_pat_let_tv32).F18())))
					}(_pat_let18_0)
				}(_688_v67), 21)
				(_nw96).ArraySet1(_688_v67, 22)
				(_nw96).ArraySet1(_688_v67, 23)
				(_nw96).ArraySet1(_688_v67, 24)
				_689_v68 = _nw96
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_689_v68), 0))
				_ = _index78
				(_689_v68).ArraySet1(Companion_D6_.Create_DC10_(_687_v66), (_index78).Int())
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_689_v68), 0))
				_ = _index79
				(_689_v68).ArraySet1(_688_v67, (_index79).Int())
				var _694_v69 _dafny.Map
				_ = _694_v69
				_694_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _680_v59)
				var _695_v70 _dafny.Map
				_ = _695_v70
				_695_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_694_v69).Contains(!(!((_683_v62).F18()))), (p0) && ((_683_v62).F19()))
				var _696_v71 _dafny.Set
				_ = _696_v71
				_696_v71 = _dafny.SetOf((_683_v62).F19())
				_695_v70 = (_695_v70).Update(((_this).F11()).Cmp(_dafny.IntOfInt64(429)) > 0, (_696_v71).IsProperSubsetOf(_696_v71))
			}
			var _697_v72 _dafny.Map
			_ = _697_v72
			_697_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(globalState), _dafny.Companion_Sequence_.Concatenate(_640_v32, _640_v32))
			(globalState).F1 = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_697_v72).Contains(p0) {
					return (_697_v72).Get(p0).(_dafny.Sequence)
				}
				return _640_v32
			})()).Cardinality())
			var _698_v73 *C1
			_ = _698_v73
			var _nw97 *C1 = New_C1_()
			_ = _nw97
			_nw97.Ctor__(p0, p0)
			_698_v73 = _nw97
		}
		var _699_v74 _dafny.Sequence
		_ = _699_v74
		_699_v74 = _dafny.SeqOf(p0)
		var _700_v75 *C0
		_ = _700_v75
		var _nw98 *C0 = New_C0_()
		_ = _nw98
		_nw98.Ctor__(_dafny.Companion_Sequence_.Concatenate(_699_v74, _699_v74), ((_this).F11()).Minus((_this).F11()))
		_700_v75 = _nw98
		var _701_v76 _dafny.Map
		_ = _701_v76
		_701_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), p0)
		var _702_v77 D8
		_ = _702_v77
		_702_v77 = Companion_D8_.Create_DC16_((func() bool {
			if (_701_v76).Contains((_700_v75).F21()) {
				return (_701_v76).Get((_700_v75).F21()).(bool)
			}
			return p0
		})(), (_700_v75).F21())
		var _pat_let_tv33 = p0
		_ = _pat_let_tv33
		var _pat_let_tv34 = p0
		_ = _pat_let_tv34
		if func(_source12 D8) bool {
			if _source12.Is_DC16() {
				var _703___mcc_h17 bool = _source12.Get_().(D8_DC16).Cf23
				_ = _703___mcc_h17
				var _704___mcc_h18 _dafny.Int = _source12.Get_().(D8_DC16).Cf24
				_ = _704___mcc_h18
				var _705_cf24 _dafny.Int = _704___mcc_h18
				_ = _705_cf24
				var _706_cf23 bool = _703___mcc_h17
				_ = _706_cf23
				return true
			} else if _source12.Is_DC15() {
				var _707___mcc_h19 _dafny.Map = _source12.Get_().(D8_DC15).Cf22
				_ = _707___mcc_h19
				var _708_cf22 _dafny.Map = _707___mcc_h19
				_ = _708_cf22
				return true
			} else {
				var _709___mcc_h20 D8 = _source12.Get_().(D8_DC17).Cf25
				_ = _709___mcc_h20
				var _710_cf25 D8 = _709___mcc_h20
				_ = _710_cf25
				return (_pat_let_tv33) == (false)
			}
		}(func(_pat_let20_0 D8) D8 {
			return func(_711_dt__update__tmp_h3 D8) D8 {
				return func(_pat_let21_0 bool) D8 {
					return func(_712_dt__update_hcf23_h0 bool) D8 {
						return Companion_D8_.Create_DC16_(_712_dt__update_hcf23_h0, (_711_dt__update__tmp_h3).Dtor_cf24())
					}(_pat_let21_0)
				}(_pat_let_tv34)
			}(_pat_let20_0)
		}(_702_v77)) {
			var _713_v78 D2
			_ = _713_v78
			_713_v78 = Companion_D2_.Create_DC5_()
			var _source13 D2 = _713_v78
			_ = _source13
			if _source13.Is_DC4() {
				var _714___mcc_h8 _dafny.Int = _source13.Get_().(D2_DC4).Cf6
				_ = _714___mcc_h8
				var _715___mcc_h9 _dafny.Int = _source13.Get_().(D2_DC4).Cf7
				_ = _715___mcc_h9
				var _716___mcc_h10 _dafny.Int = _source13.Get_().(D2_DC4).Cf8
				_ = _716___mcc_h10
				var _717_cf8 _dafny.Int = _716___mcc_h10
				_ = _717_cf8
				var _718_cf7 _dafny.Int = _715___mcc_h9
				_ = _718_cf7
				var _719_cf6 _dafny.Int = _714___mcc_h8
				_ = _719_cf6
				var _720_v79 _dafny.Array
				_ = _720_v79
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_11
				var _nw99 _dafny.Array
				_ = _nw99
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw99 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) D8 = (func(_721_v77 D8) func(_dafny.Int) D8 {
						return func(_722_i7 _dafny.Int) D8 {
							return _721_v77
						}
					})(_702_v77)
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw99 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw99).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw99).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_720_v79 = _nw99
				var _723_v80 D12
				_ = _723_v80
				_723_v80 = Companion_D12_.Create_DC25_(_720_v79)
				var _pat_let_tv35 = _720_v79
				_ = _pat_let_tv35
				var _rhs124 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_700_v75).F21(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)).Cardinality()), (_700_v75).F21())
				_ = _rhs124
				var _rhs125 _dafny.Array = (func(_pat_let22_0 D12) D12 {
					return func(_724_dt__update__tmp_h4 D12) D12 {
						return func(_pat_let23_0 _dafny.Array) D12 {
							return func(_725_dt__update_hcf37_h0 _dafny.Array) D12 {
								return Companion_D12_.Create_DC25_(_725_dt__update_hcf37_h0)
							}(_pat_let23_0)
						}(_pat_let_tv35)
					}(_pat_let22_0)
				}(_723_v80)).Dtor_cf37()
				_ = _rhs125
				var _lhs103 *GlobalState = globalState
				_ = _lhs103
				_lhs103.F1 = _rhs124
				_720_v79 = _rhs125
				var _726_v81 _dafny.Array
				_ = _726_v81
				var _nw100 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
				_ = _nw100
				_726_v81 = _nw100
				var _727_v82 _dafny.MultiSet
				_ = _727_v82
				_727_v82 = _dafny.MultiSetOf(p0, p0)
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_726_v81), 0))
				_ = _index80
				(_726_v81).ArraySet1((_727_v82).Cardinality(), (_index80).Int())
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_726_v81), 0))
				_ = _index81
				(_726_v81).ArraySet1((_this).F11(), (_index81).Int())
				r1 = p1
				var _728_v83 _dafny.Array
				_ = _728_v83
				var _nwElement0_33 _dafny.Sequence = _dafny.SeqOf((_700_v75).F21(), (_700_v75).F21(), _dafny.IntOfInt64(737))
				_ = _nwElement0_33
				var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(10))
				_ = _nw101
				(_nw101).ArraySet1(_nwElement0_33, 0)
				(_nw101).ArraySet1(_640_v32, 1)
				(_nw101).ArraySet1(_dafny.SeqOf(_719_cf6), 2)
				(_nw101).ArraySet1(_640_v32, 3)
				(_nw101).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_700_v75).F21()), _dafny.SeqOf(_718_cf7)), 4)
				(_nw101).ArraySet1(_640_v32, 5)
				(_nw101).ArraySet1(_640_v32, 6)
				(_nw101).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(296))).Uint32(), func(coer63 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}(func(_729_i8 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf((_this).F24())).Cardinality())
				})), 7)
				(_nw101).ArraySet1(_640_v32, 8)
				(_nw101).ArraySet1(_640_v32, 9)
				_728_v83 = _nw101
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_728_v83), 0))
				_ = _index82
				(_728_v83).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(47))).Uint32(), func(coer64 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg64 _dafny.Int) interface{} {
						return coer64(arg64)
					}
				}((func(_730_cf6 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_731_i9 _dafny.Int) _dafny.Int {
						return _730_cf6
					}
				})(_719_cf6))), Companion_Default___.Fm34(p0, Companion_D6_.Create_DC11_(_dafny.IntOfInt64(100), _718_cf7, (_700_v75).F21()), globalState)), (_index82).Int())
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_728_v83), 0))
				_ = _index83
				(_728_v83).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus((_726_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_726_v81), 0))).Int()).(_dafny.Int))), _640_v32), _640_v32), (_index83).Int())
			} else if _source13.Is_DC5() {
				var _732_v84 _dafny.Set
				_ = _732_v84
				_732_v84 = _dafny.SetOf(p0, false)
				var _733_v85 _dafny.Array
				_ = _733_v85
				var _nwElement0_34 _dafny.Set = _732_v84
				_ = _nwElement0_34
				var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(3))
				_ = _nw102
				(_nw102).ArraySet1(_nwElement0_34, 0)
				(_nw102).ArraySet1((_732_v84).Difference(_dafny.SetOf(p0)), 1)
				(_nw102).ArraySet1((_732_v84).Union(_732_v84), 2)
				_733_v85 = _nw102
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_733_v85), 0))
				_ = _index84
				(_733_v85).ArraySet1(_732_v84, (_index84).Int())
				var _734_v86 _dafny.Sequence
				_ = _734_v86
				_734_v86 = _dafny.UnicodeSeqOfUtf8Bytes("cym")
				var _735_v87 D6
				_ = _735_v87
				_735_v87 = Companion_D6_.Create_DC11_((_this).F11(), (_700_v75).F21(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_734_v86).Cardinality())))
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_733_v85), 0))
				_ = _index85
				var _rhs126 _dafny.Sequence = Companion_Default___.Fm34((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("c")).Cardinality())).Cmp((_700_v75).F21()) < 0, _735_v87, globalState)
				_ = _rhs126
				var _rhs127 _dafny.Set = _732_v84
				_ = _rhs127
				var _lhs104 _dafny.Array = _733_v85
				_ = _lhs104
				var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_733_v85), 0))
				_ = _lhs105
				_640_v32 = _rhs126
				(_lhs104).ArraySet1(_rhs127, (_lhs105).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))
				_ = _index86
				(p1).ArraySet1(p0, (_index86).Int())
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))
				_ = _index87
				(p1).ArraySet1(!(p0), (_index87).Int())
				(globalState).F1 = Companion_Default___.SafeDivisionInt((_dafny.IntOfUint32((_640_v32).Cardinality())).Times((_dafny.Zero).Minus((_this).F11())), (_dafny.IntOfInt64(48)).Times((_700_v75).F21()))
				var _736_v88 _dafny.Map
				_ = _736_v88
				_736_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool), p0)
				var _737_v89 _dafny.Array
				_ = _737_v89
				var _nwElement0_35 bool = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool)
				_ = _nwElement0_35
				var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(23))
				_ = _nw103
				(_nw103).ArraySet1(_nwElement0_35, 0)
				(_nw103).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool), 1)
				(_nw103).ArraySet1(false, 2)
				(_nw103).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool), 3)
				(_nw103).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool), 4)
				(_nw103).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool), 5)
				(_nw103).ArraySet1(p0, 6)
				(_nw103).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool), 7)
				(_nw103).ArraySet1(p0, 8)
				(_nw103).ArraySet1(!(p0) || (p0), 9)
				(_nw103).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool), 10)
				(_nw103).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool), 11)
				(_nw103).ArraySet1((func() bool {
					if (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool) {
						return p0
					}
					return !(p0)
				})(), 12)
				(_nw103).ArraySet1((_700_v75).Fm21((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool), (_700_v75).F21(), globalState), 13)
				(_nw103).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool), 14)
				(_nw103).ArraySet1(false, 15)
				(_nw103).ArraySet1(false, 16)
				(_nw103).ArraySet1((p0) && ((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool)), 17)
				(_nw103).ArraySet1((func() bool {
					if p0 {
						return true
					}
					return (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool)
				})(), 18)
				(_nw103).ArraySet1(p0, 19)
				(_nw103).ArraySet1((_700_v75).Fm21(p0, (_735_v87).Dtor_cf14(), globalState), 20)
				(_nw103).ArraySet1((func() bool {
					if (_736_v88).Contains(p0) {
						return (_736_v88).Get(p0).(bool)
					}
					return !((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((p1), 0))).Int()).(bool))
				})(), 21)
				(_nw103).ArraySet1(p0, 22)
				_737_v89 = _nw103
				r1 = _737_v89
			} else {
				var _738___mcc_h11 _dafny.MultiSet = _source13.Get_().(D2_DC3).Cf5
				_ = _738___mcc_h11
				var _739_cf5 _dafny.MultiSet = _738___mcc_h11
				_ = _739_cf5
				var _740_v90 _dafny.Sequence
				_ = _740_v90
				_740_v90 = _dafny.SeqOf(_dafny.SeqOf(p0, true), _dafny.SeqOf(p0))
				var _741_v91 _dafny.Map
				_ = _741_v91
				_741_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F24())
				var _742_v92 *C0
				_ = _742_v92
				var _nw104 *C0 = New_C0_()
				_ = _nw104
				_nw104.Ctor__(_dafny.Companion_Sequence_.Update((_740_v90).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_740_v90).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_741_v91).Cardinality(), _dafny.IntOfUint32(((_740_v90).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_740_v90).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), p0), (_this).F11())
				_742_v92 = _nw104
				var _743_v93 _dafny.Map
				_ = _743_v93
				_743_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm1(globalState))
				var _744_v94 _dafny.Sequence
				_ = _744_v94
				_744_v94 = _dafny.UnicodeSeqOfUtf8Bytes("qxgefovt")
				var _745_v95 _dafny.Map
				_ = _745_v95
				_745_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
					if (_743_v93).Contains(p0) {
						return (_743_v93).Get(p0).(_dafny.Sequence)
					}
					return _744_v94
				})(), (_700_v75).F21())
				var _746_v96 _dafny.Map
				_ = _746_v96
				_746_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_700_v75).F21())
				_745_v95 = (_745_v95).Update(_dafny.Companion_Sequence_.Concatenate(_744_v94, _744_v94), ((_746_v96).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_700_v75).F21()))).Cardinality())
				var _747_v97 _dafny.MultiSet
				_ = _747_v97
				_747_v97 = _dafny.MultiSetOf(p0, p0)
				(globalState).F1 = (Companion_Default___.SafeDivisionInt((_700_v75).F21(), (_this).F11())).Times(((func() _dafny.Int {
					if (_747_v97).Contains(p0) {
						return (_747_v97).Multiplicity(p0)
					}
					return (_700_v75).F21()
				})()).Times((_742_v92).F21()))
				var _748_v98 *C0
				_ = _748_v98
				var _nw105 *C0 = New_C0_()
				_ = _nw105
				_nw105.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0), _dafny.Companion_Sequence_.Update((_700_v75).F20(), (Companion_Default___.SafeIndex((_700_v75).F21(), _dafny.IntOfUint32(((_700_v75).F20()).Cardinality()))).Uint32(), p0)), ((_746_v96).Merge(_746_v96)).Cardinality())
				_748_v98 = _nw105
			}
			(globalState).F1 = (_dafny.Zero).Minus((_700_v75).F21())
			var _749_v99 _dafny.Sequence
			_ = _749_v99
			_749_v99 = _dafny.UnicodeSeqOfUtf8Bytes("uesixv")
			var _rhs128 _dafny.Int = (_700_v75).F21()
			_ = _rhs128
			var _rhs129 _dafny.Sequence = _749_v99
			_ = _rhs129
			var _lhs106 *GlobalState = globalState
			_ = _lhs106
			_lhs106.F1 = _rhs128
			_749_v99 = _rhs129
			(globalState).F1 = ((_this).F11()).Plus((_this).F11())
			var _rhs130 bool = ((_700_v75).F21()).Cmp(((_700_v75).F21()).Times((_700_v75).F21())) <= 0
			_ = _rhs130
			var _rhs131 bool = !(p0)
			_ = _rhs131
			var _lhs107 *GlobalState = globalState
			_ = _lhs107
			var _lhs108 *GlobalState = globalState
			_ = _lhs108
			_lhs107.F7 = _rhs130
			_lhs108.F6 = _rhs131
		} else {
			(globalState).F7 = p0
			var _750_v100 _dafny.Array
			_ = _750_v100
			var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
			_ = _nw106
			_750_v100 = _nw106
			var _751_v101 _dafny.Array
			_ = _751_v101
			var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
			_ = _nw107
			_751_v101 = _nw107
			var _752_v102 _dafny.Array
			_ = _752_v102
			_752_v102 = _751_v101
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_750_v100), 0))
			_ = _index88
			(_750_v100).ArraySet1(_752_v102, (_index88).Int())
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_750_v100), 0))
			_ = _index89
			(_750_v100).ArraySet1(_752_v102, (_index89).Int())
			var _753_v103 *C0
			_ = _753_v103
			var _nw108 *C0 = New_C0_()
			_ = _nw108
			_nw108.Ctor__(_dafny.Companion_Sequence_.Concatenate(_699_v74, _699_v74), (_700_v75).F21())
			_753_v103 = _nw108
			var _754_v104 D7
			_ = _754_v104
			_754_v104 = Companion_D7_.Create_DC12_(_dafny.Companion_Sequence_.Update(_640_v32, (Companion_Default___.SafeIndex((_753_v103).F21(), _dafny.IntOfUint32((_640_v32).Cardinality()))).Uint32(), (_753_v103).F21()))
			var _source14 D7 = _754_v104
			_ = _source14
			if _source14.Is_DC13() {
				var _755___mcc_h12 _dafny.Set = _source14.Get_().(D7_DC13).Cf18
				_ = _755___mcc_h12
				var _756___mcc_h13 _dafny.Int = _source14.Get_().(D7_DC13).Cf19
				_ = _756___mcc_h13
				var _757___mcc_h14 _dafny.Int = _source14.Get_().(D7_DC13).Cf20
				_ = _757___mcc_h14
				var _758_cf20 _dafny.Int = _757___mcc_h14
				_ = _758_cf20
				var _759_cf19 _dafny.Int = _756___mcc_h13
				_ = _759_cf19
				var _760_cf18 _dafny.Set = _755___mcc_h12
				_ = _760_cf18
				var _761_v105 bool
				_ = _761_v105
				var _762_v106 bool
				_ = _762_v106
				var _763_v107 bool
				_ = _763_v107
				var _764_v108 _dafny.Int
				_ = _764_v108
				var _out20 bool
				_ = _out20
				var _out21 bool
				_ = _out21
				var _out22 bool
				_ = _out22
				var _out23 _dafny.Int
				_ = _out23
				_out20, _out21, _out22, _out23 = (_this).M8(globalState)
				_761_v105 = _out20
				_762_v106 = _out21
				_763_v107 = _out22
				_764_v108 = _out23
				var _765_v109 _dafny.Sequence
				_ = _765_v109
				_765_v109 = _dafny.UnicodeSeqOfUtf8Bytes("mbcbg")
				var _rhs132 _dafny.Sequence = _dafny.SeqOf(_dafny.Companion_Sequence_.IsPrefixOf(_765_v109, _765_v109))
				_ = _rhs132
				var _rhs133 _dafny.Array = p1
				_ = _rhs133
				_699_v74 = _rhs132
				r1 = _rhs133
				(globalState).F6 = p0
				_763_v107 = !(_761_v105)
			} else if _source14.Is_DC12() {
				var _766___mcc_h15 _dafny.Sequence = _source14.Get_().(D7_DC12).Cf17
				_ = _766___mcc_h15
				var _767_cf17 _dafny.Sequence = _766___mcc_h15
				_ = _767_cf17
				(globalState).F7 = !(p0)
				var _768_v110 *C0
				_ = _768_v110
				var _nw109 *C0 = New_C0_()
				_ = _nw109
				_nw109.Ctor__((_753_v103).F20(), (_700_v75).F21())
				_768_v110 = _nw109
				var _769_v111 _dafny.Sequence
				_ = _769_v111
				_769_v111 = _dafny.UnicodeSeqOfUtf8Bytes("ckt")
				_769_v111 = _769_v111
				var _770_v112 _dafny.Sequence
				_ = _770_v112
				_770_v112 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F24()))
				var _771_v113 _dafny.Sequence
				_ = _771_v113
				_771_v113 = _769_v111
				var _772_v114 _dafny.Map
				_ = _772_v114
				_772_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_770_v112).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm2((_753_v103).F21(), _dafny.IntOfUint32((_769_v111).Cardinality()), globalState), _dafny.IntOfUint32((_770_v112).Cardinality()))).Uint32()).(_dafny.Map), _771_v113)
				var _773_v115 _dafny.Map
				_ = _773_v115
				_773_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F24())
				_772_v114 = (_772_v114).Update(_773_v115, _771_v113)
			} else {
				var _774___mcc_h16 D7 = _source14.Get_().(D7_DC14).Cf21
				_ = _774___mcc_h16
				var _775_cf21 D7 = _774___mcc_h16
				_ = _775_cf21
				var _776_v116 _dafny.Map
				_ = _776_v116
				_776_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				(globalState).F7 = (_753_v103).Fm21((func() bool {
					if (_776_v116).Contains(p0) {
						return (_776_v116).Get(p0).(bool)
					}
					return p0
				})(), (_700_v75).F21(), globalState)
				(globalState).F7 = p0
				var _777_v117 _dafny.Set
				_ = _777_v117
				_777_v117 = _dafny.SetOf(_dafny.IntOfInt64(324), (_700_v75).F21(), ((_700_v75).F21()).Plus((_753_v103).F21()))
				var _rhs134 _dafny.Int = (_777_v117).Cardinality()
				_ = _rhs134
				var _rhs135 bool = _dafny.Companion_Sequence_.Equal(_699_v74, Companion_Default___.Fm35((_753_v103).F21(), (_700_v75).F21(), Companion_Default___.Fm36(p0, p0, (_this).F11(), globalState), _dafny.IntOfInt64(60), globalState))
				_ = _rhs135
				var _lhs109 *GlobalState = globalState
				_ = _lhs109
				var _lhs110 *GlobalState = globalState
				_ = _lhs110
				_lhs109.F1 = _rhs134
				_lhs110.F7 = _rhs135
				_640_v32 = _640_v32
			}
			if (func() bool {
				if p0 {
					return false
				}
				return p0
			})() {
				_701_v76 = (_701_v76).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(535))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}((func(_778_v103 *C0) func(_dafny.Int) _dafny.Int {
					return func(_779_i10 _dafny.Int) _dafny.Int {
						return (_778_v103).F21()
					}
				})(_753_v103))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(85))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg66 _dafny.Int) interface{} {
						return coer66(arg66)
					}
				}((func(_780_v75 *C0) func(_dafny.Int) _dafny.Int {
					return func(_781_i11 _dafny.Int) _dafny.Int {
						return (func() _dafny.Map {
							var _coll44 = _dafny.NewMapBuilder()
							_ = _coll44
							for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-190), _dafny.IntOfInt64(507))); ; {
								_compr_44, _ok45 := _iter45()
								if !_ok45 {
									break
								}
								var _782_v118 _dafny.Int
								_782_v118 = interface{}(_compr_44).(_dafny.Int)
								if ((_dafny.IntOfInt64(-190)).Cmp(_782_v118) <= 0) && ((_782_v118).Cmp(_dafny.IntOfInt64(507)) < 0) {
									_coll44.Add((_782_v118).Minus((_780_v75).F21()), (_dafny.Zero).Minus((_780_v75).F21()))
								}
							}
							return _coll44.ToMap()
						}()).Cardinality()
					}
				})(_700_v75))))).Cardinality()), p0)
				(globalState).F1 = (_753_v103).F21()
				var _783_v119 _dafny.Sequence
				_ = _783_v119
				_783_v119 = _dafny.UnicodeSeqOfUtf8Bytes("v")
				var _784_v120 _dafny.Sequence
				_ = _784_v120
				_784_v120 = _dafny.SeqOf(_783_v119)
				(globalState).F7 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_783_v119, _dafny.UnicodeSeqOfUtf8Bytes("lviqvc")), _dafny.Companion_Sequence_.Concatenate(_783_v119, (_784_v120).Select((Companion_Default___.SafeIndex((_700_v75).F21(), _dafny.IntOfUint32((_784_v120).Cardinality()))).Uint32()).(_dafny.Sequence)))
				(globalState).F1 = _dafny.IntOfUint32((Companion_Default___.Fm1(globalState)).Cardinality())
				var _785_v121 _dafny.MultiSet
				_ = _785_v121
				_785_v121 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.IsProperPrefixOf(_783_v119, _783_v119), !_dafny.Companion_Sequence_.Contains(_783_v119, (_this).F24()))
				var _rhs136 _dafny.Int = ((_this).Fm3(_dafny.IntOfInt64(983), true, p0, globalState)).Minus((_700_v75).F21())
				_ = _rhs136
				var _rhs137 _dafny.MultiSet = ((func() _dafny.MultiSet {
					if p0 {
						return Companion_Default___.Fm37(globalState)
					}
					return _dafny.MultiSetFromSeq((_700_v75).F20())
				})()).Update((func() bool {
					if false {
						return p0
					}
					return p0
				})(), Companion_Default___.Abs((_753_v103).F21()))
				_ = _rhs137
				var _rhs138 bool = !(p0)
				_ = _rhs138
				var _lhs111 *GlobalState = globalState
				_ = _lhs111
				var _lhs112 *GlobalState = globalState
				_ = _lhs112
				_lhs111.F1 = _rhs136
				_785_v121 = _rhs137
				_lhs112.F7 = _rhs138
			} else {
				(globalState).F6 = p0
				(globalState).F1 = (func() _dafny.Int {
					if p0 {
						return (_dafny.Zero).Minus((_this).F11())
					}
					return (_753_v103).F21()
				})()
				var _786_v122 _dafny.Array
				_ = _786_v122
				var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw110
				_786_v122 = _nw110
				_786_v122 = _786_v122
				var _787_v123 _dafny.Array
				_ = _787_v123
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_12
				var _nw111 _dafny.Array
				_ = _nw111
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw111 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) _dafny.Sequence = func(_788_i12 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("pswwmeq")
					}
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw111 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw111).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw111).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_787_v123 = _nw111
				var _rhs139 _dafny.Array = _787_v123
				_ = _rhs139
				var _rhs140 _dafny.Int = (_753_v103).F21()
				_ = _rhs140
				var _lhs113 *GlobalState = globalState
				_ = _lhs113
				_787_v123 = _rhs139
				_lhs113.F1 = _rhs140
				var _789_v124 _dafny.MultiSet
				_ = _789_v124
				_789_v124 = _dafny.MultiSetOf(_699_v74)
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_786_v122), 0))
				_ = _index90
				(_786_v122).ArraySet1(((func() _dafny.MultiSet {
					if p0 {
						return _789_v124
					}
					return _789_v124
				})()).Cardinality(), (_index90).Int())
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_786_v122), 0))
				_ = _index91
				(_786_v122).ArraySet1(Companion_Default___.Fm2((_dafny.Zero).Minus((_this).Fm3((_700_v75).F21(), !(p0), p0, globalState)), (_700_v75).F21(), globalState), (_index91).Int())
			}
		}
		var _790_v125 D6
		_ = _790_v125
		_790_v125 = Companion_D6_.Create_DC11_(_dafny.IntOfInt64(85), (_700_v75).F21(), (_700_v75).F21())
		if _dafny.Companion_Sequence_.Equal(Companion_Default___.Fm34(p0, _790_v125, globalState), _dafny.SeqOf((_700_v75).F21(), (_this).F11())) {
			(globalState).F1 = (_700_v75).F21()
			if false {
				var _791_v126 _dafny.Sequence
				_ = _791_v126
				_791_v126 = _dafny.UnicodeSeqOfUtf8Bytes("jkdpcyeue")
				var _792_v127 _dafny.Map
				_ = _792_v127
				_792_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _791_v126)
				(globalState).F1 = (Companion_Default___.Fm38(_792_v127, p0, false, globalState)).Dtor_cf0()
				var _793_v128 _dafny.Array
				_ = _793_v128
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_13
				var _nw112 _dafny.Array
				_ = _nw112
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw112 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Int = (func(_794_v126 _dafny.Sequence, _795_p0 bool) func(_dafny.Int) _dafny.Int {
						return func(_796_i13 _dafny.Int) _dafny.Int {
							return (_796_i13).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_794_v126).Cardinality()), _795_p0)).Cardinality())
						}
					})(_791_v126, p0)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw112 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw112).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw112).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_793_v128 = _nw112
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_793_v128), 0))
				_ = _index92
				(_793_v128).ArraySet1((_700_v75).F21(), (_index92).Int())
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_793_v128), 0))
				_ = _index93
				(_793_v128).ArraySet1((_700_v75).F21(), (_index93).Int())
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_793_v128), 0))
				_ = _index94
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_793_v128), 0))
				_ = _index95
				var _rhs141 _dafny.Int = (_700_v75).F21()
				_ = _rhs141
				var _rhs142 _dafny.Int = ((_700_v75).F21()).Times((_700_v75).F21())
				_ = _rhs142
				var _lhs114 _dafny.Array = _793_v128
				_ = _lhs114
				var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_793_v128), 0))
				_ = _lhs115
				var _lhs116 _dafny.Array = _793_v128
				_ = _lhs116
				var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_793_v128), 0))
				_ = _lhs117
				(_lhs114).ArraySet1(_rhs141, (_lhs115).Int())
				(_lhs116).ArraySet1(_rhs142, (_lhs117).Int())
				var _rhs143 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_793_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_793_v128), 0))).Int()).(_dafny.Int)))
				_ = _rhs143
				var _rhs144 bool = p0
				_ = _rhs144
				var _rhs145 bool = _dafny.Companion_Sequence_.IsPrefixOf(_640_v32, _640_v32)
				_ = _rhs145
				var _rhs146 bool = (func() bool {
					if !(false) {
						return p0
					}
					return true
				})()
				_ = _rhs146
				var _lhs118 *GlobalState = globalState
				_ = _lhs118
				var _lhs119 *GlobalState = globalState
				_ = _lhs119
				var _lhs120 *GlobalState = globalState
				_ = _lhs120
				var _lhs121 *GlobalState = globalState
				_ = _lhs121
				_lhs118.F1 = _rhs143
				_lhs119.F7 = _rhs144
				_lhs120.F7 = _rhs145
				_lhs121.F6 = _rhs146
				(globalState).F6 = p0
				(globalState).F1 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_700_v75).F21(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(35))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}(func(_797_i14 _dafny.Int) _dafny.CodePoint {
					return (_this).F24()
				})), _dafny.UnicodeSeqOfUtf8Bytes("w"))).Cardinality())))
			} else {
				var _798_v129 bool
				_ = _798_v129
				var _799_v130 bool
				_ = _799_v130
				var _800_v131 bool
				_ = _800_v131
				var _801_v132 _dafny.Int
				_ = _801_v132
				var _out24 bool
				_ = _out24
				var _out25 bool
				_ = _out25
				var _out26 bool
				_ = _out26
				var _out27 _dafny.Int
				_ = _out27
				_out24, _out25, _out26, _out27 = (_this).M8(globalState)
				_798_v129 = _out24
				_799_v130 = _out25
				_800_v131 = _out26
				_801_v132 = _out27
				(globalState).F6 = !((_700_v75).Fm21(false, Companion_Default___.SafeDivisionInt(_801_v132, (_700_v75).F21()), globalState))
				r1 = p1
				var _802_v133 D11
				_ = _802_v133
				_802_v133 = Companion_D11_.Create_DC23_(_799_v130, (_this).F11(), (_700_v75).F21())
				(globalState).F7 = (_802_v133).Dtor_cf31()
				(globalState).F1 = (_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt((_700_v75).F21(), (_dafny.Zero).Minus((_700_v75).F21()))).Plus((_this).F11()))
			}
			var _803_v134 _dafny.MultiSet
			_ = _803_v134
			_803_v134 = _dafny.MultiSetOf((_this).F24(), (_this).F24())
			var _804_v135 _dafny.Sequence
			_ = _804_v135
			_804_v135 = _dafny.UnicodeSeqOfUtf8Bytes("vds")
			_701_v76 = (_701_v76).Update((func() _dafny.Int {
				if (_803_v134).Contains((_this).F24()) {
					return (_803_v134).Multiplicity((_this).F24())
				}
				return (_this).Fm3((_700_v75).F21(), p0, p0, globalState)
			})(), !_dafny.Companion_Sequence_.Equal(_804_v135, _804_v135))
			var _805_v136 *C0
			_ = _805_v136
			var _nw113 *C0 = New_C0_()
			_ = _nw113
			_nw113.Ctor__(_699_v74, (_700_v75).F21())
			_805_v136 = _nw113
			var _806_v137 *C0
			_ = _806_v137
			var _nw114 *C0 = New_C0_()
			_ = _nw114
			_nw114.Ctor__((func() _dafny.Sequence {
				if p0 {
					return (_700_v75).F20()
				}
				return _699_v74
			})(), (_805_v136).F21())
			_806_v137 = _nw114
		} else {
			var _807_v138 *C0
			_ = _807_v138
			var _nw115 *C0 = New_C0_()
			_ = _nw115
			_nw115.Ctor__((_700_v75).F20(), Companion_Default___.SafeDivisionInt(Companion_Default___.Fm2((_this).F11(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(451))).Uint32(), func(coer68 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg68 _dafny.Int) interface{} {
					return coer68(arg68)
				}
			}((func(_808_v75 *C0) func(_dafny.Int) _dafny.Int {
				return func(_809_i15 _dafny.Int) _dafny.Int {
					return (_808_v75).F21()
				}
			})(_700_v75)))).Cardinality()), globalState), (func() _dafny.Int {
				if (_639_v31).Contains((_this).F11()) {
					return (_639_v31).Multiplicity((_this).F11())
				}
				return (_700_v75).F21()
			})()))
			_807_v138 = _nw115
			(globalState).F6 = (func() bool {
				if p0 {
					return (_700_v75).Fm21(p0, (_700_v75).F21(), globalState)
				}
				return p0
			})()
			var _810_v139 *C0
			_ = _810_v139
			var _nw116 *C0 = New_C0_()
			_ = _nw116
			_nw116.Ctor__((_700_v75).F20(), (_700_v75).F21())
			_810_v139 = _nw116
			(globalState).F1 = (_807_v138).F21()
			var _811_v140 _dafny.Map
			_ = _811_v140
			_811_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_807_v138).F20()).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).Fm3((_dafny.Zero).Minus((_810_v139).F21()), p0, p0, globalState)), _dafny.IntOfUint32(((_807_v138).F20()).Cardinality()))).Uint32()).(bool), p0)
			var _812_v141 _dafny.MultiSet
			_ = _812_v141
			_812_v141 = _dafny.MultiSetOf(p0)
			var _813_v142 _dafny.Sequence
			_ = _813_v142
			_813_v142 = _dafny.UnicodeSeqOfUtf8Bytes("pag")
			var _814_v143 _dafny.Array
			_ = _814_v143
			var _nwElement0_36 _dafny.Int = (_807_v138).F21()
			_ = _nwElement0_36
			var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(23))
			_ = _nw117
			(_nw117).ArraySet1(_nwElement0_36, 0)
			(_nw117).ArraySet1((_810_v139).F21(), 1)
			(_nw117).ArraySet1(_dafny.IntOfUint32(((_807_v138).F20()).Cardinality()), 2)
			(_nw117).ArraySet1((_700_v75).F21(), 3)
			(_nw117).ArraySet1(_dafny.IntOfInt64(-913), 4)
			(_nw117).ArraySet1((_700_v75).F21(), 5)
			(_nw117).ArraySet1((_810_v139).F21(), 6)
			(_nw117).ArraySet1((_700_v75).F21(), 7)
			(_nw117).ArraySet1((_807_v138).F21(), 8)
			(_nw117).ArraySet1((_807_v138).F21(), 9)
			(_nw117).ArraySet1((_700_v75).F21(), 10)
			(_nw117).ArraySet1((_639_v31).Cardinality(), 11)
			(_nw117).ArraySet1((_807_v138).F21(), 12)
			(_nw117).ArraySet1((_this).F11(), 13)
			(_nw117).ArraySet1((_807_v138).F21(), 14)
			(_nw117).ArraySet1((_700_v75).F21(), 15)
			(_nw117).ArraySet1(_dafny.IntOfUint32((_813_v142).Cardinality()), 16)
			(_nw117).ArraySet1(_dafny.IntOfUint32((_640_v32).Cardinality()), 17)
			(_nw117).ArraySet1((_701_v76).Cardinality(), 18)
			(_nw117).ArraySet1(_dafny.IntOfUint32((_813_v142).Cardinality()), 19)
			(_nw117).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sowrbha")).Cardinality()), 20)
			(_nw117).ArraySet1((_700_v75).F21(), 21)
			(_nw117).ArraySet1((_this).F11(), 22)
			_814_v143 = _nw117
			var _815_v144 _dafny.Sequence
			_ = _815_v144
			_815_v144 = _dafny.SeqOf(_814_v143)
			var _816_v145 _dafny.Map
			_ = _816_v145
			_816_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_640_v32, p0)
			var _817_v146 _dafny.Map
			_ = _817_v146
			_817_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_639_v31).Contains(_dafny.IntOfInt64(26)) {
					return (_639_v31).Multiplicity(_dafny.IntOfInt64(26))
				}
				return (_700_v75).F21()
			})(), Companion_Default___.Fm39(_dafny.IntOfUint32((_813_v142).Cardinality()), p0, _dafny.IntOfUint32((_640_v32).Cardinality()), globalState))
			var _818_v147 _dafny.Array
			_ = _818_v147
			var _nwElement0_37 _dafny.Int = _dafny.IntOfInt64(660)
			_ = _nwElement0_37
			var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(27))
			_ = _nw118
			(_nw118).ArraySet1(_nwElement0_37, 0)
			(_nw118).ArraySet1((func() _dafny.Int {
				if true {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _811_v140)).Cardinality()
				}
				return (_810_v139).F21()
			})(), 1)
			(_nw118).ArraySet1(_dafny.IntOfUint32((_699_v74).Cardinality()), 2)
			(_nw118).ArraySet1((_700_v75).F21(), 3)
			(_nw118).ArraySet1((_807_v138).F21(), 4)
			(_nw118).ArraySet1((_810_v139).F21(), 5)
			(_nw118).ArraySet1(((_807_v138).F21()).Times((_700_v75).F21()), 6)
			(_nw118).ArraySet1((_700_v75).F21(), 7)
			(_nw118).ArraySet1(Companion_Default___.Fm2((_this).F11(), (_700_v75).F21(), globalState), 8)
			(_nw118).ArraySet1(((_807_v138).F21()).Minus((_812_v141).Cardinality()), 9)
			(_nw118).ArraySet1(_dafny.IntOfUint32((_815_v144).Cardinality()), 10)
			(_nw118).ArraySet1((_this).F11(), 11)
			(_nw118).ArraySet1((_dafny.Zero).Minus((_700_v75).F21()), 12)
			(_nw118).ArraySet1((_this).F11(), 13)
			(_nw118).ArraySet1(_dafny.IntOfInt64(98), 14)
			(_nw118).ArraySet1((_700_v75).F21(), 15)
			(_nw118).ArraySet1((_816_v145).Cardinality(), 16)
			(_nw118).ArraySet1(((_700_v75).F21()).Plus((_700_v75).F21()), 17)
			(_nw118).ArraySet1(Companion_Default___.SafeModuloInt((_807_v138).F21(), ((_dafny.MultiSetOf(p0)).Update(p0, Companion_Default___.Abs((_700_v75).F21()))).Cardinality()), 18)
			(_nw118).ArraySet1((_this).F11(), 19)
			(_nw118).ArraySet1((_810_v139).F21(), 20)
			(_nw118).ArraySet1((_817_v146).Cardinality(), 21)
			(_nw118).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(983), _dafny.IntOfUint32((_813_v142).Cardinality())), 22)
			(_nw118).ArraySet1((_807_v138).F21(), 23)
			(_nw118).ArraySet1((_810_v139).F21(), 24)
			(_nw118).ArraySet1((_this).F11(), 25)
			(_nw118).ArraySet1(((_807_v138).F21()).Times((_807_v138).F21()), 26)
			_818_v147 = _nw118
			var _819_v148 _dafny.Map
			_ = _819_v148
			_819_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _814_v143)
			var _820_v149 _dafny.Map
			_ = _820_v149
			_820_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm40((_807_v138).F20(), (_this).F24(), (_this).F11(), globalState)).Cardinality(), _818_v147)
			_814_v143 = (func() _dafny.Array {
				if (_819_v148).Contains(p1) {
					return (_819_v148).Get(p1).(_dafny.Array)
				}
				return (func() _dafny.Array {
					if (_820_v149).Contains((_810_v139).F21()) {
						return (_820_v149).Get((_810_v139).F21()).(_dafny.Array)
					}
					return _814_v143
				})()
			})()
		}
		var _821_v150 _dafny.Sequence
		_ = _821_v150
		_821_v150 = _dafny.UnicodeSeqOfUtf8Bytes("tbi")
		var _822_v151 *C1
		_ = _822_v151
		var _nw119 *C1 = New_C1_()
		_ = _nw119
		_nw119.Ctor__(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-580))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg69 _dafny.Int) interface{} {
				return coer69(arg69)
			}
		}(func(_823_i16 _dafny.Int) _dafny.CodePoint {
			return (_this).F24()
		})), _821_v150), (p0) && (p0))
		_822_v151 = _nw119
		var _824_v152 _dafny.Set
		_ = _824_v152
		_824_v152 = _dafny.SetOf(p0)
		var _825_v153 D7
		_ = _825_v153
		_825_v153 = Companion_D7_.Create_DC13_((_824_v152).Union(_dafny.SetOf(p0)), (_639_v31).Cardinality(), _dafny.IntOfInt64(263))
		r0 = _825_v153
		r1 = p1
		return r0, r1
	}
}
func (_this *C2) F24() _dafny.CodePoint {
	{
		return _this._f24
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f12 _dafny.Set
	_f11 _dafny.Int
	F23  _dafny.Int
	_f22 _dafny.Int
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f12 = _dafny.EmptySet
	_this._f11 = _dafny.Zero
	_this.F23 = _dafny.Zero
	_this._f22 = _dafny.Zero
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

func (_this *C3) F12() _dafny.Set {
	return _this._f12
}
func (_this *C3) F11() _dafny.Int {
	return _this._f11
}
func (_this *C3) Ctor__(f22 _dafny.Int, f23 _dafny.Int, f11 _dafny.Int, f12 _dafny.Set) {
	{
		(_this)._f22 = f22
		(_this).F23 = f23
		(_this)._f11 = f11
		(_this)._f12 = f12
	}
}
func (_this *C3) Fm3(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F11()).Plus((_this).F11())
	}
}
func (_this *C3) Fm4(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return !(!(!(!((_this.F23).Cmp((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true, false), _dafny.SeqOf(!(!(true)), true, false, true)))).Cardinality()) <= 0))))
	}
}
func (_this *C3) Fm5(p0 D0, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		var _source15 D7 = Companion_D7_.Create_DC14_(Companion_D7_.Create_DC14_(Companion_D7_.Create_DC12_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-233))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg70 _dafny.Int) interface{} {
				return coer70(arg70)
			}
		}(func(_826_i0 _dafny.Int) _dafny.Int {
			return (_this).F11()
		})))))
		_ = _source15
		if _source15.Is_DC13() {
			var _827___mcc_h0 _dafny.Set = _source15.Get_().(D7_DC13).Cf18
			_ = _827___mcc_h0
			var _828___mcc_h1 _dafny.Int = _source15.Get_().(D7_DC13).Cf19
			_ = _828___mcc_h1
			var _829___mcc_h2 _dafny.Int = _source15.Get_().(D7_DC13).Cf20
			_ = _829___mcc_h2
			var _830_cf20 _dafny.Int = _829___mcc_h2
			_ = _830_cf20
			var _831_cf19 _dafny.Int = _828___mcc_h1
			_ = _831_cf19
			var _832_cf18 _dafny.Set = _827___mcc_h0
			_ = _832_cf18
			return !(true)
		} else if _source15.Is_DC12() {
			var _833___mcc_h3 _dafny.Sequence = _source15.Get_().(D7_DC12).Cf17
			_ = _833___mcc_h3
			var _834_cf17 _dafny.Sequence = _833___mcc_h3
			_ = _834_cf17
			return ((_this).F11()).Cmp((_this).F11()) != 0
		} else {
			var _835___mcc_h4 D7 = _source15.Get_().(D7_DC14).Cf21
			_ = _835___mcc_h4
			var _836_cf21 D7 = _835___mcc_h4
			_ = _836_cf21
			return (true) && (false)
		}
	}
}
func (_this *C3) Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(694)
	}
}
func (_this *C3) Fm25(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((Companion_D0_.Create_DC1_(false, _dafny.IntOfInt64(444), true)).Dtor_cf1()), !(false))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge((Companion_D8_.Create_DC15_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Dtor_cf22()))
	}
}
func (_this *C3) M0(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var _837_v0 _dafny.Sequence
		_ = _837_v0
		_837_v0 = _dafny.UnicodeSeqOfUtf8Bytes("njxqrukhg")
		(globalState).F6 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xihlij"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(598))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg71 _dafny.Int) interface{} {
				return coer71(arg71)
			}
		}(func(_838_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))), _837_v0)
		var _839_v1 bool
		_ = _839_v1
		_839_v1 = false
		if _839_v1 {
			var _840_v2 _dafny.Sequence
			_ = _840_v2
			_840_v2 = _dafny.SeqOf((_this).F11())
			var _rhs147 bool = !(_839_v1)
			_ = _rhs147
			var _rhs148 _dafny.Int = (func() _dafny.Int {
				if false {
					return (_840_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_837_v0).Cardinality()), _dafny.IntOfUint32((_840_v2).Cardinality()))).Uint32()).(_dafny.Int)
				}
				return (_this).F11()
			})()
			_ = _rhs148
			var _rhs149 bool = ((_this).F11()).Cmp(((_this).F22()).Minus(_dafny.IntOfUint32((_837_v0).Cardinality()))) <= 0
			_ = _rhs149
			var _lhs122 *GlobalState = globalState
			_ = _lhs122
			var _lhs123 *C3 = _this
			_ = _lhs123
			_lhs122.F7 = _rhs147
			_lhs123.F23 = _rhs148
			_839_v1 = _rhs149
			var _841_v3 _dafny.MultiSet
			_ = _841_v3
			_841_v3 = _dafny.MultiSetOf(_839_v1, _839_v1)
			var _842_v4 _dafny.Sequence
			_ = _842_v4
			_842_v4 = _dafny.SeqOf(true, false, _839_v1)
			var _843_v5 D0
			_ = _843_v5
			_843_v5 = Companion_D0_.Create_DC0_((_this).F11())
			var _844_v6 _dafny.Sequence
			_ = _844_v6
			_844_v6 = _dafny.SeqOf(_840_v2)
			var _845_v7 _dafny.Array
			_ = _845_v7
			var _nw120 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw120
			_845_v7 = _nw120
			var _846_v8 _dafny.Map
			_ = _846_v8
			_846_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_845_v7, _839_v1)
			var _847_v9 _dafny.Array
			_ = _847_v9
			var _nwElement0_38 bool = (_841_v3).IsSubsetOf(_841_v3)
			_ = _nwElement0_38
			var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(21))
			_ = _nw121
			(_nw121).ArraySet1(_nwElement0_38, 0)
			(_nw121).ArraySet1(_839_v1, 1)
			(_nw121).ArraySet1(_839_v1, 2)
			(_nw121).ArraySet1((_this).Fm5(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(-65)), _842_v4, _dafny.IntOfInt64(533), p0, globalState), 3)
			(_nw121).ArraySet1(_839_v1, 4)
			(_nw121).ArraySet1(((_this).Fm6(_839_v1, _this.F23, globalState)).Cmp((_this).F22()) != 0, 5)
			(_nw121).ArraySet1((func() bool {
				if _839_v1 {
					return false
				}
				return !(_839_v1)
			})(), 6)
			(_nw121).ArraySet1(_839_v1, 7)
			(_nw121).ArraySet1(_839_v1, 8)
			(_nw121).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("qkcgskmp"), _dafny.UnicodeSeqOfUtf8Bytes("ju")), 9)
			(_nw121).ArraySet1((_this).Fm5(_843_v5, _842_v4, (_this).F11(), _dafny.CodePoint('o'), globalState), 10)
			(_nw121).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_844_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(76))).Uint32(), func(coer72 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg72 _dafny.Int) interface{} {
					return coer72(arg72)
				}
			}((func(_848_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_849_i1 _dafny.Int) _dafny.Sequence {
					return _848_v2
				}
			})(_840_v2)))), 11)
			(_nw121).ArraySet1(((_this).F11()).Cmp(_this.F23) >= 0, 12)
			(_nw121).ArraySet1((_839_v1) && (_839_v1), 13)
			(_nw121).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_840_v2, _dafny.Companion_Sequence_.Update(_840_v2, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_840_v2).Cardinality()))).Uint32(), _dafny.IntOfUint32(((_844_v6).Select((Companion_Default___.SafeIndex(_this.F23, _dafny.IntOfUint32((_844_v6).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))), 14)
			(_nw121).ArraySet1(!((func() bool {
				if (_846_v8).Contains(_845_v7) {
					return (_846_v8).Get(_845_v7).(bool)
				}
				return _839_v1
			})()) || (_839_v1), 15)
			(_nw121).ArraySet1((_842_v4).Select((Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_842_v4).Cardinality()))).Uint32()).(bool), 16)
			(_nw121).ArraySet1(!(false) || (true), 17)
			(_nw121).ArraySet1(((_this).F22()).Cmp(_this.F23) >= 0, 18)
			(_nw121).ArraySet1(_839_v1, 19)
			(_nw121).ArraySet1(_839_v1, 20)
			_847_v9 = _nw121
			var _850_v10 _dafny.MultiSet
			_ = _850_v10
			_850_v10 = _dafny.MultiSetOf(_dafny.IntOfInt64(145), ((_this).F12()).Cardinality())
			var _851_v11 _dafny.Map
			_ = _851_v11
			_851_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_839_v1, (_this).F11())
			var _852_v12 _dafny.Map
			_ = _852_v12
			_852_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_851_v11).Cardinality(), _dafny.IntOfUint32((_837_v0).Cardinality()))
			var _853_v13 _dafny.Map
			_ = _853_v13
			_853_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_850_v10, _852_v12)
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_847_v9), 0))
			_ = _index96
			(_847_v9).ArraySet1((_this).Fm4(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oxsuipcwe")).Cardinality()), _853_v13, _this.F23, _839_v1, globalState), (_index96).Int())
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_847_v9), 0))
			_ = _index97
			(_847_v9).ArraySet1(_839_v1, (_index97).Int())
			var _854_v14 *C0
			_ = _854_v14
			var _nw122 *C0 = New_C0_()
			_ = _nw122
			_nw122.Ctor__(_842_v4, (func() _dafny.Int {
				if (_851_v11).Contains((_847_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_847_v9), 0))).Int()).(bool)) {
					return (_851_v11).Get((_847_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_847_v9), 0))).Int()).(bool)).(_dafny.Int)
				}
				return (_dafny.Zero).Minus((_this).F22())
			})())
			_854_v14 = _nw122
			_854_v14 = _854_v14
			var _855_v15 _dafny.Sequence
			_ = _855_v15
			_855_v15 = _837_v0
			var _rhs150 bool = _839_v1
			_ = _rhs150
			var _rhs151 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("nb")
			_ = _rhs151
			var _lhs124 *GlobalState = globalState
			_ = _lhs124
			_lhs124.F7 = _rhs150
			_855_v15 = _rhs151
			var _856_v16 _dafny.Map
			_ = _856_v16
			_856_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_841_v3, Companion_Default___.Fm26(globalState))
			_851_v11 = (_851_v11).Update(((_this).F11()).Cmp((_851_v11).Cardinality()) > 0, (_856_v16).Cardinality())
		} else {
			var _857_v17 _dafny.Sequence
			_ = _857_v17
			_857_v17 = _dafny.SeqOf(_this.F23)
			var _858_v18 D8
			_ = _858_v18
			_858_v18 = Companion_D8_.Create_DC16_(_839_v1, (_dafny.Zero).Minus(_this.F23))
			var _859_v19 D8
			_ = _859_v19
			_859_v19 = Companion_D8_.Create_DC17_(_858_v18)
			var _860_v20 _dafny.Map
			_ = _860_v20
			_860_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_857_v17, _858_v18)
			var _861_v21 D8
			_ = _861_v21
			_861_v21 = Companion_D8_.Create_DC17_((func() D8 {
				if (_860_v20).Contains(_dafny.SeqOf(_this.F23)) {
					return (_860_v20).Get(_dafny.SeqOf(_this.F23)).(D8)
				}
				return _858_v18
			})())
			_861_v21 = _861_v21
			var _862_v22 D0
			_ = _862_v22
			_862_v22 = Companion_D0_.Create_DC0_(_this.F23)
			var _863_v23 _dafny.Sequence
			_ = _863_v23
			_863_v23 = _dafny.SeqOf(_839_v1)
			var _864_v24 _dafny.MultiSet
			_ = _864_v24
			_864_v24 = _dafny.MultiSetOf((_this).F11())
			var _865_v25 _dafny.MultiSet
			_ = _865_v25
			_865_v25 = _dafny.MultiSetOf(_864_v24)
			var _866_v26 _dafny.Sequence
			_ = _866_v26
			_866_v26 = _dafny.SeqOf(_863_v23)
			var _867_v27 _dafny.Array
			_ = _867_v27
			var _nwElement0_39 bool = !(_839_v1) || (_839_v1)
			_ = _nwElement0_39
			var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(18))
			_ = _nw123
			(_nw123).ArraySet1(_nwElement0_39, 0)
			(_nw123).ArraySet1(!(_839_v1), 1)
			(_nw123).ArraySet1(_839_v1, 2)
			(_nw123).ArraySet1(true, 3)
			(_nw123).ArraySet1(_839_v1, 4)
			(_nw123).ArraySet1((_this).Fm5(_862_v22, _863_v23, _this.F23, p0, globalState), 5)
			(_nw123).ArraySet1((_this.F23).Cmp(_this.F23) > 0, 6)
			(_nw123).ArraySet1(!(!(_839_v1)), 7)
			(_nw123).ArraySet1(_839_v1, 8)
			(_nw123).ArraySet1(_839_v1, 9)
			(_nw123).ArraySet1(!((_dafny.SetOf((_864_v24).Cardinality())).IsProperSubsetOf(_dafny.SetOf((_this).F11()))), 10)
			(_nw123).ArraySet1((_865_v25).IsSubsetOf(_865_v25), 11)
			(_nw123).ArraySet1(_839_v1, 12)
			(_nw123).ArraySet1(!(_839_v1), 13)
			(_nw123).ArraySet1(_839_v1, 14)
			(_nw123).ArraySet1(_839_v1, 15)
			(_nw123).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32(((_866_v26).Select((Companion_Default___.SafeIndex(_this.F23, _dafny.IntOfUint32((_866_v26).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Cmp(_this.F23) == 0, 16)
			(_nw123).ArraySet1(_839_v1, 17)
			_867_v27 = _nw123
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_867_v27), 0))
			_ = _index98
			(_867_v27).ArraySet1(_839_v1, (_index98).Int())
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_867_v27), 0))
			_ = _index99
			(_867_v27).ArraySet1(_839_v1, (_index99).Int())
			var _868_v28 _dafny.Map
			_ = _868_v28
			_868_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_867_v27, (_dafny.Zero).Minus((_this.F23).Plus(_dafny.IntOfInt64(732))))
			_868_v28 = (_868_v28).Update(_867_v27, (_dafny.IntOfUint32((_837_v0).Cardinality())).Minus((_this).F22()))
			_839_v1 = true
			var _869_v29 _dafny.Array
			_ = _869_v29
			var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw124
			_869_v29 = _nw124
			var _870_v30 _dafny.Set
			_ = _870_v30
			_870_v30 = _dafny.SetOf((_this).F22(), _this.F23, (_dafny.SetOf((_this).F11())).Cardinality(), _this.F23)
			var _871_v31 _dafny.Map
			_ = _871_v31
			_871_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(491), (_870_v30).Cardinality())
			var _872_v32 _dafny.Map
			_ = _872_v32
			_872_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_869_v29, _871_v31)
			_872_v32 = (_872_v32).Update((func() _dafny.Array {
				if (_867_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_867_v27), 0))).Int()).(bool) {
					return _869_v29
				}
				return _869_v29
			})(), _871_v31)
		}
		var _873_i2 _dafny.Int
		_ = _873_i2
		_873_i2 = _dafny.Zero
		{
			for false {
				{
					if (_873_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_873_i2 = (_873_i2).Plus(_dafny.One)
					var _874_v33 *C1
					_ = _874_v33
					var _nw125 *C1 = New_C1_()
					_ = _nw125
					_nw125.Ctor__(_839_v1, !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("kdveoar"), _dafny.UnicodeSeqOfUtf8Bytes("csqsqpon")))
					_874_v33 = _nw125
					(globalState).F6 = false
					(globalState).F1 = (_this).F22()
					(_this).F23 = (_dafny.Zero).Minus((_this).F11())
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _875_v34 _dafny.Sequence
		_ = _875_v34
		_875_v34 = _dafny.SeqOf(true, _839_v1)
		var _876_v35 *C0
		_ = _876_v35
		var _nw126 *C0 = New_C0_()
		_ = _nw126
		_nw126.Ctor__(_dafny.Companion_Sequence_.Concatenate(_875_v34, _875_v34), (_dafny.IntOfInt64(461)).Minus((_this).F11()))
		_876_v35 = _nw126
		var _877_v36 _dafny.Set
		_ = _877_v36
		_877_v36 = _dafny.SetOf(_dafny.IntOfInt64(342))
		var _878_v37 _dafny.Array
		_ = _878_v37
		var _nwElement0_40 bool = !((_dafny.IntOfUint32((_837_v0).Cardinality())).Cmp(_this.F23) == 0)
		_ = _nwElement0_40
		var _nw127 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(16))
		_ = _nw127
		(_nw127).ArraySet1(_nwElement0_40, 0)
		(_nw127).ArraySet1(_839_v1, 1)
		(_nw127).ArraySet1((_dafny.SetOf((_this).F22())).IsProperSubsetOf(_877_v36), 2)
		(_nw127).ArraySet1(false, 3)
		(_nw127).ArraySet1(!(_839_v1), 4)
		(_nw127).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf(_839_v1)).Cardinality())).Cmp(_dafny.IntOfInt64(-62)) > 0, 5)
		(_nw127).ArraySet1(_839_v1, 6)
		(_nw127).ArraySet1((func() bool {
			if _839_v1 {
				return _839_v1
			}
			return _839_v1
		})(), 7)
		(_nw127).ArraySet1(false, 8)
		(_nw127).ArraySet1(_839_v1, 9)
		(_nw127).ArraySet1(_839_v1, 10)
		(_nw127).ArraySet1((_839_v1) && (_839_v1), 11)
		(_nw127).ArraySet1(_839_v1, 12)
		(_nw127).ArraySet1(_839_v1, 13)
		(_nw127).ArraySet1(_839_v1, 14)
		(_nw127).ArraySet1(_839_v1, 15)
		_878_v37 = _nw127
		var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))
		_ = _index100
		(_878_v37).ArraySet1(_839_v1, (_index100).Int())
		var _879_v38 _dafny.Map
		_ = _879_v38
		_879_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_839_v1, (_876_v35).F21())
		var _880_v39 _dafny.MultiSet
		_ = _880_v39
		_880_v39 = _dafny.MultiSetOf(_879_v38)
		var _881_v40 _dafny.Map
		_ = _881_v40
		_881_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_876_v35).F21(), _880_v39)
		var _882_v41 D9
		_ = _882_v41
		_882_v41 = Companion_D9_.Create_DC18_(_879_v38)
		var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))
		_ = _index101
		(_878_v37).ArraySet1(((_880_v39).Difference(_880_v39)).IsSubsetOf((func() _dafny.MultiSet {
			if (_881_v40).Contains(_this.F23) {
				return (_881_v40).Get(_this.F23).(_dafny.MultiSet)
			}
			return _dafny.MultiSetOf(_879_v38, _879_v38, (_882_v41).Dtor_cf26())
		})()), (_index101).Int())
		var _883_v42 _dafny.Map
		_ = _883_v42
		_883_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool), (_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool))
		var _884_i3 _dafny.Int
		_ = _884_i3
		_884_i3 = _dafny.Zero
		{
			for (_883_v42).Contains((_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool)) {
				{
					if (_884_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_884_i3 = (_884_i3).Plus(_dafny.One)
					if !(_839_v1) {
						var _885_v43 _dafny.MultiSet
						_ = _885_v43
						_885_v43 = _dafny.MultiSetOf((_this).F22(), (_876_v35).F21(), (_this).Fm3((_dafny.Zero).Minus((_this).F11()), _839_v1, _839_v1, globalState))
						var _886_v44 _dafny.Map
						_ = _886_v44
						_886_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_876_v35).F21(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-99))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg73 _dafny.Int) interface{} {
								return coer73(arg73)
							}
						}((func(_887_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_888_i4 _dafny.Int) _dafny.CodePoint {
								return _887_p0
							}
						})(p0))))
						(_this).F23 = (((_this).F22()).Minus((_876_v35).F21())).Minus(Companion_Default___.SafeModuloInt((_885_v43).Cardinality(), (_this).Fm6((_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool), (_886_v44).Cardinality(), globalState)))
						(globalState).F7 = _839_v1
						_839_v1 = (_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool)
						(globalState).F1 = Companion_Default___.SafeModuloInt((_876_v35).F21(), _this.F23)
						var _889_v45 _dafny.MultiSet
						_ = _889_v45
						_889_v45 = _dafny.MultiSetOf((_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool), _839_v1)
						var _890_v46 _dafny.Map
						_ = _890_v46
						_890_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_889_v45, _878_v37)
						_890_v46 = _890_v46
					} else {
						r0 = _dafny.MultiSetOf(_839_v1, (_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool), (func() bool {
							if (_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool) {
								return _839_v1
							}
							return _839_v1
						})())
						var _891_v47 _dafny.Map
						_ = _891_v47
						_891_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm26(globalState)).Cardinality(), _839_v1)
						var _892_v48 _dafny.CodePoint
						_ = _892_v48
						_892_v48 = _dafny.CodePoint('d')
						var _893_v49 _dafny.Map
						_ = _893_v49
						_893_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(Companion_Default___.Fm27(_839_v1, _839_v1, globalState)), false)
						var _rhs152 bool = (_893_v49).Equals((_893_v49).Merge(_893_v49))
						_ = _rhs152
						var _rhs153 _dafny.Int = ((_this).F11()).Minus(_this.F23)
						_ = _rhs153
						var _rhs154 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_837_v0, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qsx"), _dafny.UnicodeSeqOfUtf8Bytes("mhrhniii")))
						_ = _rhs154
						var _rhs155 _dafny.Map = (_891_v47).Merge(_891_v47)
						_ = _rhs155
						var _rhs156 _dafny.CodePoint = p0
						_ = _rhs156
						var _lhs125 *GlobalState = globalState
						_ = _lhs125
						var _lhs126 *GlobalState = globalState
						_ = _lhs126
						_lhs125.F6 = _rhs152
						_lhs126.F1 = _rhs153
						_837_v0 = _rhs154
						_891_v47 = _rhs155
						_892_v48 = _rhs156
						var _894_v50 _dafny.MultiSet
						_ = _894_v50
						_894_v50 = _dafny.MultiSetOf(false)
						var _895_v51 _dafny.Sequence
						_ = _895_v51
						_895_v51 = _dafny.SeqOf((_dafny.Zero).Minus((func() _dafny.Int {
							if (_894_v50).Contains(_839_v1) {
								return (_894_v50).Multiplicity(_839_v1)
							}
							return _this.F23
						})()))
						var _896_v52 _dafny.MultiSet
						_ = _896_v52
						_896_v52 = _dafny.MultiSetOf((_895_v51).Select((Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_895_v51).Cardinality()))).Uint32()).(_dafny.Int), _this.F23)
						_896_v52 = _896_v52
						var _897_v53 _dafny.Map
						_ = _897_v53
						_897_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_839_v1, p0)
						var _898_v54 _dafny.Array
						_ = _898_v54
						var _nwElement0_41 _dafny.Int = (_this).F11()
						_ = _nwElement0_41
						var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(19))
						_ = _nw128
						(_nw128).ArraySet1(_nwElement0_41, 0)
						(_nw128).ArraySet1((_897_v53).Cardinality(), 1)
						(_nw128).ArraySet1(_this.F23, 2)
						(_nw128).ArraySet1(_dafny.IntOfInt64(522), 3)
						(_nw128).ArraySet1(_dafny.IntOfUint32((_895_v51).Cardinality()), 4)
						(_nw128).ArraySet1(_dafny.IntOfInt64(845), 5)
						(_nw128).ArraySet1((_this).F11(), 6)
						(_nw128).ArraySet1((_876_v35).F21(), 7)
						(_nw128).ArraySet1((_876_v35).F21(), 8)
						(_nw128).ArraySet1((_this).F22(), 9)
						(_nw128).ArraySet1((_876_v35).F21(), 10)
						(_nw128).ArraySet1((_883_v42).Cardinality(), 11)
						(_nw128).ArraySet1((_dafny.MultiSetOf((_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool), _839_v1)).Cardinality(), 12)
						(_nw128).ArraySet1(_this.F23, 13)
						(_nw128).ArraySet1(_dafny.IntOfUint32((_837_v0).Cardinality()), 14)
						(_nw128).ArraySet1((_this).F22(), 15)
						(_nw128).ArraySet1((_dafny.Zero).Minus((_this).F22()), 16)
						(_nw128).ArraySet1(_dafny.IntOfInt64(138), 17)
						(_nw128).ArraySet1((_876_v35).F21(), 18)
						_898_v54 = _nw128
						var _899_v55 _dafny.Set
						_ = _899_v55
						_899_v55 = _dafny.SetOf(_898_v54)
						_899_v55 = (_899_v55).Intersection((_899_v55).Union(_899_v55))
						var _900_v56 _dafny.Sequence
						_ = _900_v56
						_900_v56 = _837_v0
						var _901_v57 _dafny.Sequence
						_ = _901_v57
						_901_v57 = _dafny.SeqOf(_837_v0)
						(globalState).F7 = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_901_v57, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.IntOfUint32((_901_v57).Cardinality()))).Uint32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(537))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg74 _dafny.Int) interface{} {
								return coer74(arg74)
							}
						}((func(_902_v48 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_903_i5 _dafny.Int) _dafny.CodePoint {
								return _902_v48
							}
						})(_892_v48)))), (_900_v56))
					}
					(globalState).F6 = (_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool)
					var _904_v58 _dafny.CodePoint
					_ = _904_v58
					_904_v58 = _dafny.CodePoint('w')
					_904_v58 = p0
					var _905_v59 _dafny.Sequence
					_ = _905_v59
					_905_v59 = _dafny.SeqOf((_876_v35).F21())
					var _906_v60 T0
					_ = _906_v60
					var _nw129 *C2 = New_C2_()
					_ = _nw129
					_nw129.Ctor__(_904_v58, _dafny.IntOfUint32((_905_v59).Cardinality()))
					_906_v60 = _nw129
					var _907_v61 _dafny.Map
					_ = _907_v61
					_907_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_906_v60, _878_v37)
					var _908_v62 D0
					_ = _908_v62
					_908_v62 = Companion_D0_.Create_DC0_((_907_v61).Cardinality())
					(globalState).F6 = (_this).Fm5(_908_v62, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool)), (_876_v35).F20()), ((_this).Fm6((_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool), (_this).F11(), globalState)).Times(_this.F23), _dafny.CodePoint('e'), globalState)
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _909_v63 _dafny.MultiSet
		_ = _909_v63
		_909_v63 = _dafny.MultiSetOf(_839_v1)
		r0 = ((_909_v63).Intersection(_dafny.MultiSetOf((_878_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_878_v37), 0))).Int()).(bool), _839_v1, _839_v1, _839_v1))).Update((_839_v1) && (_839_v1), Companion_Default___.Abs((_this).F11()))
		return r0
	}
}
func (_this *C3) M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Map) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _910_v0 bool
		_ = _910_v0
		_910_v0 = true
		if _910_v0 {
			var _911_v1 _dafny.CodePoint
			_ = _911_v1
			_911_v1 = _dafny.CodePoint('v')
			var _912_v2 _dafny.Array
			_ = _912_v2
			var _nwElement0_42 _dafny.Int = (_this).F22()
			_ = _nwElement0_42
			var _nw130 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.One)
			_ = _nw130
			(_nw130).ArraySet1(_nwElement0_42, 0)
			_912_v2 = _nw130
			var _913_v3 *C2
			_ = _913_v3
			var _nw131 *C2 = New_C2_()
			_ = _nw131
			_nw131.Ctor__(_911_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_912_v2, (_this).F11())).Cardinality())
			_913_v3 = _nw131
			(globalState).F6 = _910_v0
			var _914_v4 _dafny.Set
			_ = _914_v4
			_914_v4 = _dafny.SetOf(_dafny.CodePoint('f'), _dafny.CodePoint('i'))
			(globalState).F1 = (_dafny.Zero).Minus((_this).Fm3((_914_v4).Cardinality(), !(_910_v0), _910_v0, globalState))
			var _915_v5 _dafny.Set
			_ = _915_v5
			_915_v5 = _dafny.SetOf((_this).F22(), (_dafny.IntOfInt64(283)).Times(p0), ((_this).F22()).Minus(_dafny.IntOfInt64(565)), p0)
			(globalState).F1 = (_915_v5).Cardinality()
			var _916_v6 _dafny.Set
			_ = _916_v6
			_916_v6 = _dafny.SetOf(_910_v0)
			_916_v6 = Companion_Default___.Fm41(globalState)
		} else {
			var _917_v7 _dafny.Sequence
			_ = _917_v7
			_917_v7 = _dafny.SeqOf(_910_v0)
			var _918_v8 _dafny.Sequence
			_ = _918_v8
			_918_v8 = _dafny.SeqOf(p0, p0)
			var _919_v9 _dafny.Set
			_ = _919_v9
			_919_v9 = _dafny.SetOf(_918_v8)
			var _920_v10 _dafny.Map
			_ = _920_v10
			_920_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_910_v0, _919_v9)
			var _921_v11 _dafny.Map
			_ = _921_v11
			_921_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_910_v0, ((func() _dafny.Set {
				if (_920_v10).Contains(_910_v0) {
					return (_920_v10).Get(_910_v0).(_dafny.Set)
				}
				return _919_v9
			})()).Cardinality())
			var _922_v12 _dafny.Map
			_ = _922_v12
			_922_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_921_v11).Contains(_910_v0) {
					return (_921_v11).Get(_910_v0).(_dafny.Int)
				}
				return (_this).F11()
			})(), _917_v7)
			var _923_v13 *C1
			_ = _923_v13
			var _nw132 *C1 = New_C1_()
			_ = _nw132
			_nw132.Ctor__(!_dafny.Companion_Sequence_.Equal(_917_v7, (func() _dafny.Sequence {
				if (_922_v12).Contains((_this).F22()) {
					return (_922_v12).Get((_this).F22()).(_dafny.Sequence)
				}
				return _917_v7
			})()), true)
			_923_v13 = _nw132
			(globalState).F7 = _910_v0
			var _924_v14 _dafny.Sequence
			_ = _924_v14
			_924_v14 = _dafny.UnicodeSeqOfUtf8Bytes("vea")
			var _925_v15 D7
			_ = _925_v15
			_925_v15 = Companion_D7_.Create_DC13_((_this).F12(), (_this).F11(), p0)
			var _926_v16 *C0
			_ = _926_v16
			var _nw133 *C0 = New_C0_()
			_ = _nw133
			_nw133.Ctor__(_dafny.Companion_Sequence_.Concatenate(_917_v7, Companion_Default___.Fm35(_dafny.IntOfUint32((_924_v14).Cardinality()), p0, _925_v15, (_this).F11(), globalState)), p0)
			_926_v16 = _nw133
			var _927_v17 _dafny.Array
			_ = _927_v17
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_14
			var _nw134 _dafny.Array
			_ = _nw134
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw134 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) bool = (func(_928_v0 bool) func(_dafny.Int) bool {
					return func(_929_i0 _dafny.Int) bool {
						return _928_v0
					}
				})(_910_v0)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw134 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw134).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw134).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_927_v17 = _nw134
			var _rhs157 *C0 = _926_v16
			_ = _rhs157
			var _rhs158 _dafny.Array = _927_v17
			_ = _rhs158
			var _rhs159 bool = _910_v0
			_ = _rhs159
			var _rhs160 bool = (func() bool {
				if !(_910_v0) {
					return _910_v0
				}
				return (func() bool {
					if (_923_v13).F19() {
						return (_923_v13).F18()
					}
					return (_923_v13).F18()
				})()
			})()
			_ = _rhs160
			var _rhs161 bool = !(Companion_Default___.Fm0(globalState)) || ((_923_v13).F19())
			_ = _rhs161
			var _lhs127 *GlobalState = globalState
			_ = _lhs127
			var _lhs128 *GlobalState = globalState
			_ = _lhs128
			_926_v16 = _rhs157
			_927_v17 = _rhs158
			_910_v0 = _rhs159
			_lhs127.F7 = _rhs160
			_lhs128.F7 = _rhs161
			(globalState).F1 = _this.F23
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_927_v17), 0))
			_ = _index102
			(_927_v17).ArraySet1(_910_v0, (_index102).Int())
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_927_v17), 0))
			_ = _index103
			(_927_v17).ArraySet1(_910_v0, (_index103).Int())
		}
		var _930_v18 _dafny.Sequence
		_ = _930_v18
		_930_v18 = _dafny.SeqOf(_910_v0)
		var _931_v19 _dafny.Sequence
		_ = _931_v19
		_931_v19 = _dafny.SeqOf(_910_v0, (_930_v18).Select((Companion_Default___.SafeIndex(_this.F23, _dafny.IntOfUint32((_930_v18).Cardinality()))).Uint32()).(bool), _910_v0, !(_910_v0), _910_v0)
		var _932_i1 _dafny.Int
		_ = _932_i1
		_932_i1 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.Equal(_931_v19, _930_v18) {
				{
					if (_932_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_932_i1 = (_932_i1).Plus(_dafny.One)
					var _933_v20 _dafny.CodePoint
					_ = _933_v20
					_933_v20 = _dafny.CodePoint('u')
					_933_v20 = _933_v20
					var _934_v21 D0
					_ = _934_v21
					_934_v21 = Companion_D0_.Create_DC0_((_this).F11())
					var _935_v22 _dafny.MultiSet
					_ = _935_v22
					_935_v22 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("bu"))
					var _936_v23 _dafny.MultiSet
					_ = _936_v23
					_936_v23 = _dafny.MultiSetOf((_this).Fm5(_934_v21, _931_v19, (_935_v22).Cardinality(), _933_v20, globalState), Companion_Default___.Fm0(globalState), _910_v0, _910_v0)
					var _937_v24 _dafny.MultiSet
					_ = _937_v24
					_937_v24 = _936_v23
					var _source16 _dafny.MultiSet = _937_v24
					_ = _source16
					var _938___mcc_h0 _dafny.MultiSet = _source16
					_ = _938___mcc_h0
					var _939_cf11 _dafny.MultiSet = _938___mcc_h0
					_ = _939_cf11
					var _rhs162 bool = _910_v0
					_ = _rhs162
					var _rhs163 bool = !((func() bool {
						if _910_v0 {
							return _910_v0
						}
						return !(_910_v0) || (false)
					})())
					_ = _rhs163
					var _lhs129 *GlobalState = globalState
					_ = _lhs129
					var _lhs130 *GlobalState = globalState
					_ = _lhs130
					_lhs129.F7 = _rhs162
					_lhs130.F6 = _rhs163
					(globalState).F6 = _910_v0
					var _940_v25 _dafny.Array
					_ = _940_v25
					var _nwElement0_43 _dafny.MultiSet = _937_v24
					_ = _nwElement0_43
					var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(8))
					_ = _nw135
					(_nw135).ArraySet1(_nwElement0_43, 0)
					(_nw135).ArraySet1(_937_v24, 1)
					(_nw135).ArraySet1(_937_v24, 2)
					(_nw135).ArraySet1(_936_v23, 3)
					(_nw135).ArraySet1(_937_v24, 4)
					(_nw135).ArraySet1(_937_v24, 5)
					(_nw135).ArraySet1(_937_v24, 6)
					(_nw135).ArraySet1(_939_cf11, 7)
					_940_v25 = _nw135
					var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_940_v25), 0))
					_ = _index104
					(_940_v25).ArraySet1(_937_v24, (_index104).Int())
					var _941_v26 _dafny.Sequence
					_ = _941_v26
					_941_v26 = _dafny.SeqOf((_this).F22())
					var _942_v27 _dafny.Map
					_ = _942_v27
					_942_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_941_v26).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.IntOfUint32((_941_v26).Cardinality()))).Uint32()).(_dafny.Int), false)
					var _943_v28 _dafny.Sequence
					_ = _943_v28
					_943_v28 = _dafny.SeqOf(_931_v19, _930_v18)
					var _944_v29 _dafny.Set
					_ = _944_v29
					_944_v29 = _dafny.SetOf(_dafny.SeqOf(_910_v0), (_943_v28).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_943_v28).Cardinality()))).Uint32()).(_dafny.Sequence))
					var _945_v31 _dafny.MultiSet
					_ = _945_v31
					_945_v31 = _dafny.MultiSetOf(p0, _dafny.IntOfInt64(496))
					var _946_v33 _dafny.Map
					_ = _946_v33
					_946_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_945_v31, func() _dafny.Map {
						var _coll45 = _dafny.NewMapBuilder()
						_ = _coll45
						for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(381), _dafny.IntOfInt64(912))); ; {
							_compr_45, _ok46 := _iter46()
							if !_ok46 {
								break
							}
							var _947_v32 _dafny.Int
							_947_v32 = interface{}(_compr_45).(_dafny.Int)
							if ((_dafny.IntOfInt64(381)).Cmp(_947_v32) <= 0) && ((_947_v32).Cmp(_dafny.IntOfInt64(912)) < 0) {
								_coll45.Add((_947_v32).Minus((_this).F11()), (_this).F11())
							}
						}
						return _coll45.ToMap()
					}())
					var _948_v34 _dafny.Array
					_ = _948_v34
					var _nwElement0_44 bool = _910_v0
					_ = _nwElement0_44
					var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(10))
					_ = _nw136
					(_nw136).ArraySet1(_nwElement0_44, 0)
					(_nw136).ArraySet1(_910_v0, 1)
					(_nw136).ArraySet1(_910_v0, 2)
					(_nw136).ArraySet1(_910_v0, 3)
					(_nw136).ArraySet1(_910_v0, 4)
					(_nw136).ArraySet1(_910_v0, 5)
					(_nw136).ArraySet1(_910_v0, 6)
					(_nw136).ArraySet1(_910_v0, 7)
					(_nw136).ArraySet1((func() bool {
						if (_942_v27).Contains((func() _dafny.Set {
							var _coll46 = _dafny.NewBuilder()
							_ = _coll46
							for _iter47 := _dafny.Iterate((_944_v29).Elements()); ; {
								_compr_46, _ok47 := _iter47()
								if !_ok47 {
									break
								}
								var _949_v30 _dafny.Sequence
								_949_v30 = interface{}(_compr_46).(_dafny.Sequence)
								if (_944_v29).Contains(_949_v30) {
									_coll46.Add(_949_v30)
								}
							}
							return _coll46.ToSet()
						}()).Cardinality()) {
							return (_942_v27).Get((func() _dafny.Set {
								var _coll47 = _dafny.NewBuilder()
								_ = _coll47
								for _iter48 := _dafny.Iterate((_944_v29).Elements()); ; {
									_compr_47, _ok48 := _iter48()
									if !_ok48 {
										break
									}
									var _950_v30 _dafny.Sequence
									_950_v30 = interface{}(_compr_47).(_dafny.Sequence)
									if (_944_v29).Contains(_950_v30) {
										_coll47.Add(_950_v30)
									}
								}
								return _coll47.ToSet()
							}()).Cardinality()).(bool)
						}
						return (_this).Fm4(_dafny.IntOfInt64(970), _946_v33, (Companion_Default___.Fm26(globalState)).Cardinality(), _910_v0, globalState)
					})(), 8)
					(_nw136).ArraySet1((_945_v31).IsProperSubsetOf(_945_v31), 9)
					_948_v34 = _nw136
					var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_948_v34), 0))
					_ = _index105
					(_948_v34).ArraySet1(_910_v0, (_index105).Int())
					var _951_v35 *C0
					_ = _951_v35
					var _nw137 *C0 = New_C0_()
					_ = _nw137
					_nw137.Ctor__(_931_v19, (_this).F22())
					_951_v35 = _nw137
					var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_940_v25), 0))
					_ = _index106
					var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_948_v34), 0))
					_ = _index107
					var _rhs164 _dafny.MultiSet = (func() _dafny.MultiSet {
						if true {
							return _937_v24
						}
						return _939_cf11
					})()
					_ = _rhs164
					var _rhs165 bool = _910_v0
					_ = _rhs165
					var _rhs166 *C0 = _951_v35
					_ = _rhs166
					var _lhs131 _dafny.Array = _940_v25
					_ = _lhs131
					var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_940_v25), 0))
					_ = _lhs132
					var _lhs133 _dafny.Array = _948_v34
					_ = _lhs133
					var _lhs134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_948_v34), 0))
					_ = _lhs134
					(_lhs131).ArraySet1(_rhs164, (_lhs132).Int())
					(_lhs133).ArraySet1(_rhs165, (_lhs134).Int())
					_951_v35 = _rhs166
					var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_948_v34), 0))
					_ = _index108
					(_948_v34).ArraySet1(true, (_index108).Int())
					var _952_v36 _dafny.Map
					_ = _952_v36
					_952_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm26(globalState))
					var _rhs167 bool = true
					_ = _rhs167
					var _rhs168 _dafny.Int = (_dafny.Zero).Minus(((_this).F11()).Minus(p0))
					_ = _rhs168
					var _rhs169 bool = !(_952_v36).Equals(_952_v36)
					_ = _rhs169
					var _lhs135 *C3 = _this
					_ = _lhs135
					var _lhs136 *GlobalState = globalState
					_ = _lhs136
					_910_v0 = _rhs167
					_lhs135.F23 = _rhs168
					_lhs136.F6 = _rhs169
					var _953_v38 _dafny.MultiSet
					_ = _953_v38
					_953_v38 = _dafny.MultiSetOf(func() _dafny.Map {
						var _coll48 = _dafny.NewMapBuilder()
						_ = _coll48
						for _iter49 := _dafny.Iterate((_dafny.SetOf(p0)).Elements()); ; {
							_compr_48, _ok49 := _iter49()
							if !_ok49 {
								break
							}
							var _954_v37 _dafny.Int
							_954_v37 = interface{}(_compr_48).(_dafny.Int)
							if (_dafny.SetOf(p0)).Contains(_954_v37) {
								_coll48.Add((_954_v37).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_910_v0, _dafny.IntOfUint32((_dafny.SeqOf(_this.F23, (_dafny.Zero).Minus((_this).F11()))).Cardinality()))).Cardinality()), _933_v20)
							}
						}
						return _coll48.ToMap()
					}())
					var _955_v39 _dafny.Map
					_ = _955_v39
					_955_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _933_v20)
					var _956_v40 _dafny.Sequence
					_ = _956_v40
					_956_v40 = _dafny.SeqOf(_dafny.MultiSetOf(_955_v39, _955_v39), _953_v38)
					var _957_v41 _dafny.Set
					_ = _957_v41
					_957_v41 = _dafny.SetOf(_930_v18)
					if !(((func() _dafny.MultiSet {
						if _910_v0 {
							return (_956_v40).Select((Companion_Default___.SafeIndex((_957_v41).Cardinality(), _dafny.IntOfUint32((_956_v40).Cardinality()))).Uint32()).(_dafny.MultiSet)
						}
						return _953_v38
					})()).IsSubsetOf(_953_v38)) {
						var _958_v42 _dafny.Array
						_ = _958_v42
						var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
						_ = _nw138
						_958_v42 = _nw138
						var _959_v43 _dafny.Map
						_ = _959_v43
						_959_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_958_v42, _910_v0)
						_959_v43 = _959_v43
						var _960_v44 _dafny.Array
						_ = _960_v44
						var _len0_15 _dafny.Int = _dafny.IntOfInt64(5)
						_ = _len0_15
						var _nw139 _dafny.Array
						_ = _nw139
						if _len0_15.Cmp(_dafny.Zero) == 0 {
							_nw139 = _dafny.NewArray(_len0_15)
						} else {
							var _init15 func(_dafny.Int) bool = (func(_961_v0 bool) func(_dafny.Int) bool {
								return func(_962_i2 _dafny.Int) bool {
									return _961_v0
								}
							})(_910_v0)
							_ = _init15
							var _element0_15 = _init15(_dafny.Zero)
							_ = _element0_15
							_nw139 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
							(_nw139).ArraySet1(_element0_15, 0)
							var _nativeLen0_15 = (_len0_15).Int()
							_ = _nativeLen0_15
							for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
								(_nw139).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
							}
						}
						_960_v44 = _nw139
						_960_v44 = (func() _dafny.Array {
							if _910_v0 {
								return _960_v44
							}
							return _960_v44
						})()
						(globalState).F7 = ((_dafny.Zero).Minus((_this).F22())).Cmp(_this.F23) != 0
						var _963_v45 _dafny.Map
						_ = _963_v45
						_963_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _910_v0)
						var _964_v46 _dafny.Map
						_ = _964_v46
						_964_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F22()), _963_v45)
						var _965_v47 D7
						_ = _965_v47
						_965_v47 = Companion_D7_.Create_DC13_((_this).F12(), (_936_v23).Cardinality(), _this.F23)
						var _966_v48 _dafny.Sequence
						_ = _966_v48
						_966_v48 = _dafny.UnicodeSeqOfUtf8Bytes("vxka")
						var _967_v49 _dafny.Map
						_ = _967_v49
						_967_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, (_this).F11())
						var _968_v50 *C1
						_ = _968_v50
						var _nw140 *C1 = New_C1_()
						_ = _nw140
						_nw140.Ctor__(_910_v0, _910_v0)
						_968_v50 = _nw140
						var _969_v51 _dafny.Map
						_ = _969_v51
						_969_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _968_v50)
						var _rhs170 _dafny.Map = (_964_v46).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _963_v45))
						_ = _rhs170
						var _rhs171 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm35((_this).F22(), (func() _dafny.Int {
							if (_967_v49).Contains(p0) {
								return (_967_v49).Get(p0).(_dafny.Int)
							}
							return (_dafny.Zero).Minus(_this.F23)
						})(), Companion_D7_.Create_DC13_((_this).F12(), p0, p0), _this.F23, globalState), (Companion_Default___.SafeIndex((_969_v51).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm35((_this).F22(), (func() _dafny.Int {
							if (_967_v49).Contains(p0) {
								return (_967_v49).Get(p0).(_dafny.Int)
							}
							return (_dafny.Zero).Minus(_this.F23)
						})(), Companion_D7_.Create_DC13_((_this).F12(), p0, p0), _this.F23, globalState)).Cardinality()))).Uint32(), (_968_v50).F18()), _930_v18), (Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm35((_this).F22(), (func() _dafny.Int {
							if (_967_v49).Contains(p0) {
								return (_967_v49).Get(p0).(_dafny.Int)
							}
							return (_dafny.Zero).Minus(_this.F23)
						})(), Companion_D7_.Create_DC13_((_this).F12(), p0, p0), _this.F23, globalState), (Companion_Default___.SafeIndex((_969_v51).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm35((_this).F22(), (func() _dafny.Int {
							if (_967_v49).Contains(p0) {
								return (_967_v49).Get(p0).(_dafny.Int)
							}
							return (_dafny.Zero).Minus(_this.F23)
						})(), Companion_D7_.Create_DC13_((_this).F12(), p0, p0), _this.F23, globalState)).Cardinality()))).Uint32(), (_968_v50).F18()), _930_v18)).Cardinality()))).Uint32(), _910_v0), (Companion_Default___.SafeIndex(_this.F23, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm35((_this).F22(), (func() _dafny.Int {
							if (_967_v49).Contains(p0) {
								return (_967_v49).Get(p0).(_dafny.Int)
							}
							return (_dafny.Zero).Minus(_this.F23)
						})(), Companion_D7_.Create_DC13_((_this).F12(), p0, p0), _this.F23, globalState), (Companion_Default___.SafeIndex((_969_v51).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm35((_this).F22(), (func() _dafny.Int {
							if (_967_v49).Contains(p0) {
								return (_967_v49).Get(p0).(_dafny.Int)
							}
							return (_dafny.Zero).Minus(_this.F23)
						})(), Companion_D7_.Create_DC13_((_this).F12(), p0, p0), _this.F23, globalState)).Cardinality()))).Uint32(), (_968_v50).F18()), _930_v18), (Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm35((_this).F22(), (func() _dafny.Int {
							if (_967_v49).Contains(p0) {
								return (_967_v49).Get(p0).(_dafny.Int)
							}
							return (_dafny.Zero).Minus(_this.F23)
						})(), Companion_D7_.Create_DC13_((_this).F12(), p0, p0), _this.F23, globalState), (Companion_Default___.SafeIndex((_969_v51).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm35((_this).F22(), (func() _dafny.Int {
							if (_967_v49).Contains(p0) {
								return (_967_v49).Get(p0).(_dafny.Int)
							}
							return (_dafny.Zero).Minus(_this.F23)
						})(), Companion_D7_.Create_DC13_((_this).F12(), p0, p0), _this.F23, globalState)).Cardinality()))).Uint32(), (_968_v50).F18()), _930_v18)).Cardinality()))).Uint32(), _910_v0)).Cardinality()))).Uint32(), (_968_v50).F18())
						_ = _rhs171
						var _rhs172 D7 = (func() D7 {
							if (_968_v50).F19() {
								return _965_v47
							}
							return _965_v47
						})()
						_ = _rhs172
						var _rhs173 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_966_v48, _dafny.UnicodeSeqOfUtf8Bytes("eppqimqu"))
						_ = _rhs173
						_964_v46 = _rhs170
						_930_v18 = _rhs171
						_965_v47 = _rhs172
						_966_v48 = _rhs173
						var _970_v52 _dafny.Sequence
						_ = _970_v52
						_970_v52 = _dafny.SeqOf(_dafny.IntOfInt64(-590), _this.F23)
						var _971_v53 _dafny.MultiSet
						_ = _971_v53
						_971_v53 = _dafny.MultiSetOf((_this).F22())
						var _972_v54 _dafny.Sequence
						_ = _972_v54
						_972_v54 = _dafny.SeqOf(_970_v52, _970_v52, _970_v52, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0, (_971_v53).Cardinality(), _dafny.IntOfInt64(894)), _970_v52), _dafny.SeqOf((_this).F11()))
						_970_v52 = (_972_v54).Select((Companion_Default___.SafeIndex((_dafny.IntOfUint32((_970_v52).Cardinality())).Times((_this).F11()), _dafny.IntOfUint32((_972_v54).Cardinality()))).Uint32()).(_dafny.Sequence)
					} else {
						var _973_v55 _dafny.Array
						_ = _973_v55
						var _nw141 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
						_ = _nw141
						_973_v55 = _nw141
						var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_973_v55), 0))
						_ = _index109
						(_973_v55).ArraySet1(_910_v0, (_index109).Int())
						var _974_v56 _dafny.Sequence
						_ = _974_v56
						_974_v56 = _dafny.UnicodeSeqOfUtf8Bytes("iq")
						var _975_v57 D2
						_ = _975_v57
						_975_v57 = Companion_D2_.Create_DC3_(_935_v22)
						var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_973_v55), 0))
						_ = _index110
						(_973_v55).ArraySet1(((_975_v57).Dtor_cf5()).IsProperSubsetOf((_dafny.MultiSetOf(_974_v56)).Union(_935_v22)), (_index110).Int())
						_936_v23 = (_936_v23).Union((_dafny.MultiSetOf((_973_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_973_v55), 0))).Int()).(bool), (_973_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_973_v55), 0))).Int()).(bool), (_973_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_973_v55), 0))).Int()).(bool), (_973_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_973_v55), 0))).Int()).(bool), _910_v0)).Intersection(_936_v23))
						var _976_v58 *C2
						_ = _976_v58
						var _nw142 *C2 = New_C2_()
						_ = _nw142
						_nw142.Ctor__(Companion_Default___.Fm39(_dafny.IntOfInt64(175), (_973_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_973_v55), 0))).Int()).(bool), (_this).F11(), globalState), p0)
						_976_v58 = _nw142
						var _977_v59 _dafny.Map
						_ = _977_v59
						_977_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_910_v0, (_dafny.Zero).Minus((_this.F23).Times((_this).F22())))
						var _978_v60 _dafny.MultiSet
						_ = _978_v60
						_978_v60 = _dafny.MultiSetOf(_dafny.IntOfInt64(204))
						var _979_v61 _dafny.Map
						_ = _979_v61
						_979_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F11())
						var _980_v62 _dafny.Map
						_ = _980_v62
						_980_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_978_v60, _979_v61)
						_977_v59 = (_977_v59).Update(!((_this).Fm4(_this.F23, _980_v62, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_974_v56, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_974_v56).Cardinality()))).Uint32(), _933_v20), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_974_v56, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_974_v56).Cardinality()))).Uint32(), _933_v20)).Cardinality()))).Uint32(), _933_v20)).Cardinality()), true, globalState)), (_dafny.Zero).Minus((_this).F22()))
						var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_973_v55), 0))
						_ = _index111
						(_973_v55).ArraySet1(!(!(_910_v0)), (_index111).Int())
					}
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _981_v63 _dafny.Set
		_ = _981_v63
		_981_v63 = _dafny.SetOf((_this).F11())
		var _rhs174 bool = _910_v0
		_ = _rhs174
		var _rhs175 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F22(), (_this).F11())
		_ = _rhs175
		var _rhs176 _dafny.Int = p0
		_ = _rhs176
		var _rhs177 bool = !(_981_v63).Contains((_this).F22())
		_ = _rhs177
		var _lhs137 *C3 = _this
		_ = _lhs137
		var _lhs138 *C3 = _this
		_ = _lhs138
		var _lhs139 *GlobalState = globalState
		_ = _lhs139
		_910_v0 = _rhs174
		_lhs137.F23 = _rhs175
		_lhs138.F23 = _rhs176
		_lhs139.F6 = _rhs177
		var _982_v64 _dafny.Sequence
		_ = _982_v64
		_982_v64 = _dafny.UnicodeSeqOfUtf8Bytes("qnfp")
		var _983_v65 _dafny.Sequence
		_ = _983_v65
		_983_v65 = _dafny.SeqOf(_dafny.IntOfUint32((_982_v64).Cardinality()))
		var _984_v66 _dafny.MultiSet
		_ = _984_v66
		_984_v66 = _dafny.MultiSetOf((_983_v65).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_983_v65).Cardinality()))).Uint32()).(_dafny.Int), (_this).F11(), (_this).F11())
		var _985_v67 _dafny.Map
		_ = _985_v67
		_985_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_982_v64, _984_v66)
		r1 = _985_v67
		var _986_v68 _dafny.Sequence
		_ = _986_v68
		_986_v68 = _dafny.SeqOf(_930_v18, _930_v18, _930_v18)
		var _987_i3 _dafny.Int
		_ = _987_i3
		_987_i3 = _dafny.Zero
		{
			for !(true) || (_dafny.Companion_Sequence_.Contains(_986_v68, _931_v19)) {
				{
					if (_987_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_987_i3 = (_987_i3).Plus(_dafny.One)
					if _910_v0 {
						var _988_v69 _dafny.Map
						_ = _988_v69
						_988_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(785), _982_v64)
						var _rhs178 bool = (_910_v0) == ((_dafny.MultiSetOf((func() _dafny.Sequence {
							if (_988_v69).Contains((_this).F22()) {
								return (_988_v69).Get((_this).F22()).(_dafny.Sequence)
							}
							return _982_v64
						})())).IsProperSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(761))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg75 _dafny.Int) interface{} {
								return coer75(arg75)
							}
						}((func(_989_v64 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_990_i4 _dafny.Int) _dafny.Sequence {
								return _989_v64
							}
						})(_982_v64))))))
						_ = _rhs178
						var _rhs179 _dafny.Int = p0
						_ = _rhs179
						var _rhs180 _dafny.Int = _dafny.IntOfInt64(598)
						_ = _rhs180
						var _lhs140 *GlobalState = globalState
						_ = _lhs140
						var _lhs141 *GlobalState = globalState
						_ = _lhs141
						var _lhs142 *GlobalState = globalState
						_ = _lhs142
						_lhs140.F6 = _rhs178
						_lhs141.F1 = _rhs179
						_lhs142.F1 = _rhs180
						(globalState).F1 = (_this).F22()
						var _991_v70 _dafny.CodePoint
						_ = _991_v70
						_991_v70 = _dafny.CodePoint('g')
						var _992_v72 *C0
						_ = _992_v72
						var _nw143 *C0 = New_C0_()
						_ = _nw143
						_nw143.Ctor__(_931_v19, (_this).F22())
						_992_v72 = _nw143
						var _993_v73 _dafny.Sequence
						_ = _993_v73
						_993_v73 = _dafny.SeqOf(_992_v72)
						var _994_v74 D0
						_ = _994_v74
						_994_v74 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_993_v73).Cardinality()))
						var _995_v75 _dafny.MultiSet
						_ = _995_v75
						_995_v75 = _dafny.MultiSetOf(_994_v74, _994_v74, _994_v74, _994_v74, _994_v74)
						_991_v70 = Companion_Default___.Fm39(_dafny.IntOfInt64(599), Companion_Default___.Fm0(globalState), (func() _dafny.Map {
							var _coll49 = _dafny.NewMapBuilder()
							_ = _coll49
							for _iter50 := _dafny.Iterate((_995_v75).Elements()); ; {
								_compr_49, _ok50 := _iter50()
								if !_ok50 {
									break
								}
								var _996_v71 D0
								_996_v71 = interface{}(_compr_49).(D0)
								if (_995_v75).Contains(_996_v71) {
									_coll49.Add(_996_v71, (_this).F11())
								}
							}
							return _coll49.ToMap()
						}()).Cardinality(), globalState)
						(globalState).F7 = _dafny.Companion_Sequence_.Equal(_930_v18, _dafny.Companion_Sequence_.Concatenate((_992_v72).F20(), _930_v18))
						var _997_v76 _dafny.Array
						_ = _997_v76
						var _len0_16 _dafny.Int = _dafny.One
						_ = _len0_16
						var _nw144 _dafny.Array
						_ = _nw144
						if _len0_16.Cmp(_dafny.Zero) == 0 {
							_nw144 = _dafny.NewArray(_len0_16)
						} else {
							var _init16 func(_dafny.Int) _dafny.MultiSet = (func(_998_v0 bool, _999_v66 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
								return func(_1000_i5 _dafny.Int) _dafny.MultiSet {
									return (func() _dafny.MultiSet {
										if _998_v0 {
											return _999_v66
										}
										return _999_v66
									})()
								}
							})(_910_v0, _984_v66)
							_ = _init16
							var _element0_16 = _init16(_dafny.Zero)
							_ = _element0_16
							_nw144 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
							(_nw144).ArraySet1(_element0_16, 0)
							var _nativeLen0_16 = (_len0_16).Int()
							_ = _nativeLen0_16
							for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
								(_nw144).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
							}
						}
						_997_v76 = _nw144
						var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_997_v76), 0))
						_ = _index112
						(_997_v76).ArraySet1(_984_v66, (_index112).Int())
						var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_997_v76), 0))
						_ = _index113
						var _rhs181 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, (_this).F22())
						_ = _rhs181
						var _rhs182 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_982_v64, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_982_v64, _982_v64), (Companion_Default___.SafeIndex((_983_v65).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_983_v65).Cardinality()), _dafny.IntOfUint32((_983_v65).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_982_v64, _982_v64)).Cardinality()))).Uint32(), _991_v70))
						_ = _rhs182
						var _rhs183 _dafny.Int = _this.F23
						_ = _rhs183
						var _rhs184 _dafny.MultiSet = _984_v66
						_ = _rhs184
						var _lhs143 *GlobalState = globalState
						_ = _lhs143
						var _lhs144 *C3 = _this
						_ = _lhs144
						var _lhs145 _dafny.Array = _997_v76
						_ = _lhs145
						var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_997_v76), 0))
						_ = _lhs146
						_lhs143.F1 = _rhs181
						_982_v64 = _rhs182
						_lhs144.F23 = _rhs183
						(_lhs145).ArraySet1(_rhs184, (_lhs146).Int())
					} else {
						(globalState).F1 = (_this).F22()
						var _1001_v77 _dafny.Sequence
						_ = _1001_v77
						_1001_v77 = _dafny.SeqOf(_983_v65)
						var _1002_v78 _dafny.Map
						_ = _1002_v78
						_1002_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Contains(_1001_v77, _dafny.SeqOf(_this.F23, _dafny.IntOfInt64(-896))), _this.F23)
						_1002_v78 = _1002_v78
						var _1003_v79 _dafny.Map
						_ = _1003_v79
						_1003_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-274))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg76 _dafny.Int) interface{} {
								return coer76(arg76)
							}
						}(func(_1004_i6 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('q')
						})))
						var _1005_v80 _dafny.CodePoint
						_ = _1005_v80
						_1005_v80 = _dafny.CodePoint('n')
						var _rhs185 _dafny.Int = (((_981_v63).Cardinality()).Minus((_this).Fm3(p0, false, _910_v0, globalState))).Times((func() _dafny.Int {
							if _910_v0 {
								return (_this).F11()
							}
							return (_this).F11()
						})())
						_ = _rhs185
						var _rhs186 bool = !(!((_931_v19).Select((Companion_Default___.SafeIndex(_this.F23, _dafny.IntOfUint32((_931_v19).Cardinality()))).Uint32()).(bool)))
						_ = _rhs186
						var _rhs187 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(globalState), _dafny.SeqOf(_1005_v80)), _982_v64)
						_ = _rhs187
						var _rhs188 _dafny.Map = _1003_v79
						_ = _rhs188
						var _rhs189 bool = true
						_ = _rhs189
						var _lhs147 *GlobalState = globalState
						_ = _lhs147
						var _lhs148 *GlobalState = globalState
						_ = _lhs148
						_lhs147.F1 = _rhs185
						_lhs148.F6 = _rhs186
						_982_v64 = _rhs187
						_1003_v79 = _rhs188
						_910_v0 = _rhs189
						(globalState).F1 = (_this).F11()
						var _1006_v81 *C1
						_ = _1006_v81
						var _nw145 *C1 = New_C1_()
						_ = _nw145
						_nw145.Ctor__(_910_v0, (_984_v66).IsDisjointFrom(_984_v66))
						_1006_v81 = _nw145
					}
					var _1007_v82 _dafny.Map
					_ = _1007_v82
					_1007_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, _930_v18)
					var _1008_v83 *C0
					_ = _1008_v83
					var _nw146 *C0 = New_C0_()
					_ = _nw146
					_nw146.Ctor__((func() _dafny.Sequence {
						if (_1007_v82).Contains((_983_v65).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_983_v65).Cardinality()), _dafny.IntOfUint32((_983_v65).Cardinality()))).Uint32()).(_dafny.Int)) {
							return (_1007_v82).Get((_983_v65).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_983_v65).Cardinality()), _dafny.IntOfUint32((_983_v65).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Sequence)
						}
						return _930_v18
					})(), _this.F23)
					_1008_v83 = _nw146
					var _1009_v84 _dafny.Sequence
					_ = _1009_v84
					_1009_v84 = _dafny.SeqOf((_this).F12(), (_this).F12(), (_this).F12(), (_this).F12())
					_1009_v84 = _1009_v84
					_983_v65 = _dafny.SeqOf((_this).F22(), p0, (_1008_v83).F21(), _this.F23, (_this).F22())
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1010_v85 _dafny.Set
		_ = _1010_v85
		_1010_v85 = _dafny.SetOf(_983_v65)
		var _1011_v86 D9
		_ = _1011_v86
		_1011_v86 = Companion_D9_.Create_DC19_()
		var _1012_v87 _dafny.Array
		_ = _1012_v87
		var _nwElement0_45 D9 = Companion_Default___.Fm42(_1010_v85, globalState)
		_ = _nwElement0_45
		var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(18))
		_ = _nw147
		(_nw147).ArraySet1(_nwElement0_45, 0)
		(_nw147).ArraySet1(_1011_v86, 1)
		(_nw147).ArraySet1((func() D9 {
			if _910_v0 {
				return _1011_v86
			}
			return _1011_v86
		})(), 2)
		(_nw147).ArraySet1(_1011_v86, 3)
		(_nw147).ArraySet1(_1011_v86, 4)
		(_nw147).ArraySet1(_1011_v86, 5)
		(_nw147).ArraySet1(_1011_v86, 6)
		(_nw147).ArraySet1(_1011_v86, 7)
		(_nw147).ArraySet1(_1011_v86, 8)
		(_nw147).ArraySet1(_1011_v86, 9)
		(_nw147).ArraySet1(_1011_v86, 10)
		(_nw147).ArraySet1(_1011_v86, 11)
		(_nw147).ArraySet1(_1011_v86, 12)
		(_nw147).ArraySet1(_1011_v86, 13)
		(_nw147).ArraySet1(_1011_v86, 14)
		(_nw147).ArraySet1(_1011_v86, 15)
		(_nw147).ArraySet1(Companion_Default___.Fm42(_1010_v85, globalState), 16)
		(_nw147).ArraySet1(_1011_v86, 17)
		_1012_v87 = _nw147
		for _iter51 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1012_v87), 0))); ; {
			_guard_loop_1, _ok51 := _iter51()
			if !_ok51 {
				break
			}
			var _1013_i7 _dafny.Int
			_1013_i7 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_1013_i7).Sign() != -1) && ((_1013_i7).Cmp(_dafny.ArrayLen((_1012_v87), 0)) < 0)) {
				(_1012_v87).ArraySet1((func() D9 {
					if (_this.F23).Cmp(_this.F23) < 0 {
						return (func() D9 {
							if _910_v0 {
								return _1011_v86
							}
							return _1011_v86
						})()
					}
					return _1011_v86
				})(), (_1013_i7).Int())
			}
		}
		var _1014_v88 _dafny.Array
		_ = _1014_v88
		var _len0_17 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_17
		var _nw148 _dafny.Array
		_ = _nw148
		if _len0_17.Cmp(_dafny.Zero) == 0 {
			_nw148 = _dafny.NewArray(_len0_17)
		} else {
			var _init17 func(_dafny.Int) _dafny.MultiSet = (func(_1015_v0 bool) func(_dafny.Int) _dafny.MultiSet {
				return func(_1016_i8 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_1015_v0)
				}
			})(_910_v0)
			_ = _init17
			var _element0_17 = _init17(_dafny.Zero)
			_ = _element0_17
			_nw148 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
			(_nw148).ArraySet1(_element0_17, 0)
			var _nativeLen0_17 = (_len0_17).Int()
			_ = _nativeLen0_17
			for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
				(_nw148).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
			}
		}
		_1014_v88 = _nw148
		var _1017_v89 _dafny.Map
		_ = _1017_v89
		_1017_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _1014_v88)
		r0 = (func() _dafny.Array {
			if (_1017_v89).Contains((_dafny.Zero).Minus(_dafny.IntOfInt64(-6))) {
				return (_1017_v89).Get((_dafny.Zero).Minus(_dafny.IntOfInt64(-6))).(_dafny.Array)
			}
			return _1014_v88
		})()
		r1 = (_985_v67).Merge(Companion_Default___.Fm43(_910_v0, globalState))
		return r0, r1
	}
}
func (_this *C3) F22() _dafny.Int {
	{
		return _this._f22
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f12 _dafny.Set
	_f11 _dafny.Int
	_f16 _dafny.MultiSet
	_f17 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f12 = _dafny.EmptySet
	_this._f11 = _dafny.Zero
	_this._f16 = _dafny.EmptyMultiSet
	_this._f17 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C4{}
var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F12() _dafny.Set {
	return _this._f12
}
func (_this *C4) F11() _dafny.Int {
	return _this._f11
}
func (_this *C4) Ctor__(f16 _dafny.MultiSet, f17 bool, f12 _dafny.Set, f11 _dafny.Int) {
	{
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f12 = f12
		(_this)._f11 = f11
	}
}
func (_this *C4) Fm5(p0 D0, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		return (_this).F17()
	}
}
func (_this *C4) Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (((_this).F16()).Intersection((_this).F16())).Cardinality()
	}
}
func (_this *C4) Fm3(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("njlnvedvb"), _dafny.UnicodeSeqOfUtf8Bytes("jjlvfx")), _dafny.UnicodeSeqOfUtf8Bytes("uocufamr"))).Cardinality())
	}
}
func (_this *C4) Fm4(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return (_this).F17()
	}
}
func (_this *C4) M0(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		(globalState).F7 = (_this).F17()
		var _1018_i0 _dafny.Int
		_ = _1018_i0
		_1018_i0 = _dafny.Zero
		{
			for (_this).F17() {
				{
					if (_1018_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_1018_i0 = (_1018_i0).Plus(_dafny.One)
					var _1019_v0 _dafny.MultiSet
					_ = _1019_v0
					_1019_v0 = _dafny.MultiSetOf((_this).F17(), (_this).F17(), (_this).F17())
					r0 = ((func() _dafny.MultiSet {
						if (_this).F17() {
							return _1019_v0
						}
						return _dafny.MultiSetOf((_this).F17())
					})()).Difference(_1019_v0)
					var _1020_v1 _dafny.Array
					_ = _1020_v1
					var _nw149 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
					_ = _nw149
					_1020_v1 = _nw149
					var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1020_v1), 0))
					_ = _index114
					(_1020_v1).ArraySet1((_this).F17(), (_index114).Int())
					var _1021_v2 _dafny.Array
					_ = _1021_v2
					var _nw150 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(4))
					_ = _nw150
					_1021_v2 = _nw150
					var _1022_v3 _dafny.MultiSet
					_ = _1022_v3
					_1022_v3 = _dafny.MultiSetOf(_1021_v2)
					var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1020_v1), 0))
					_ = _index115
					(_1020_v1).ArraySet1(((_dafny.MultiSetOf(_1021_v2, _1021_v2)).Difference(_1022_v3)).IsDisjointFrom(_1022_v3), (_index115).Int())
					(globalState).F7 = false
					var _1023_v4 _dafny.Array
					_ = _1023_v4
					var _nw151 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(4))
					_ = _nw151
					_1023_v4 = _nw151
					var _1024_v5 _dafny.Sequence
					_ = _1024_v5
					_1024_v5 = _dafny.SeqOf((_this).F17(), (_this).F17())
					var _1025_v6 _dafny.Map
					_ = _1025_v6
					_1025_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F17()), (_this).F11())
					var _1026_v7 _dafny.Set
					_ = _1026_v7
					_1026_v7 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1020_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1020_v1), 0))).Int()).(bool), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1024_v5, (_this).Fm3((_this).F11(), (_1020_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1020_v1), 0))).Int()).(bool), (_1020_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1020_v1), 0))).Int()).(bool), globalState))).Cardinality()), _1025_v6)
					var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1023_v4), 0))
					_ = _index116
					(_1023_v4).ArraySet1((_1026_v7).Difference(_1026_v7), (_index116).Int())
					var _1027_v8 _dafny.Array
					_ = _1027_v8
					var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
					_ = _nw152
					_1027_v8 = _nw152
					var _1028_v9 _dafny.MultiSet
					_ = _1028_v9
					_1028_v9 = _dafny.MultiSetOf(_1027_v8)
					var _1029_v10 D2
					_ = _1029_v10
					_1029_v10 = Companion_D2_.Create_DC4_((_dafny.Zero).Minus((_this).F11()), (func() _dafny.Int {
						if (_1028_v9).Contains(_1027_v8) {
							return (_1028_v9).Multiplicity(_1027_v8)
						}
						return (_this).F11()
					})(), (_this).F11())
					var _1030_v11 _dafny.Sequence
					_ = _1030_v11
					_1030_v11 = _dafny.SeqOf((_this).F11(), (_this).F11())
					var _1031_v12 _dafny.Sequence
					_ = _1031_v12
					_1031_v12 = _dafny.UnicodeSeqOfUtf8Bytes("j")
					var _1032_v13 _dafny.Sequence
					_ = _1032_v13
					_1032_v13 = _dafny.SeqOf(_dafny.IntOfUint32((_1030_v11).Cardinality()), _dafny.IntOfUint32((_1024_v5).Cardinality()), _dafny.IntOfUint32((_1031_v12).Cardinality()))
					var _pat_let_tv36 = _1030_v11
					_ = _pat_let_tv36
					var _1033_v14 _dafny.Map
					_ = _1033_v14
					_1033_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), func(_pat_let24_0 D2) D2 {
						return func(_1034_dt__update__tmp_h0 D2) D2 {
							return func(_pat_let25_0 _dafny.Int) D2 {
								return func(_1035_dt__update_hcf6_h0 _dafny.Int) D2 {
									return func(_pat_let26_0 _dafny.Int) D2 {
										return func(_1036_dt__update_hcf8_h0 _dafny.Int) D2 {
											return Companion_D2_.Create_DC4_(_1035_dt__update_hcf6_h0, (_1034_dt__update__tmp_h0).Dtor_cf7(), _1036_dt__update_hcf8_h0)
										}(_pat_let26_0)
									}(_dafny.IntOfUint32((_pat_let_tv36).Cardinality()))
								}(_pat_let25_0)
							}((_this).F11())
						}(_pat_let24_0)
					}(_1029_v10))
					var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_1023_v4), 0))
					_ = _index117
					(_1023_v4).ArraySet1(Companion_Default___.Fm10((_this).F11(), _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
						if !((_1020_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1020_v1), 0))).Int()).(bool)) {
							return _dafny.UnicodeSeqOfUtf8Bytes("smwrdj")
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("tctiechk")
					})(), (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32(((func() _dafny.Sequence {
						if !((_1020_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1020_v1), 0))).Int()).(bool)) {
							return _dafny.UnicodeSeqOfUtf8Bytes("smwrdj")
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("tctiechk")
					})()).Cardinality()))).Uint32(), _dafny.CodePoint('q')), (_1033_v14).Cardinality(), (_1020_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1020_v1), 0))).Int()).(bool), globalState), (_index117).Int())
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _1037_v15 _dafny.Array
		_ = _1037_v15
		var _nw153 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
		_ = _nw153
		_1037_v15 = _nw153
		var _1038_v16 _dafny.Sequence
		_ = _1038_v16
		_1038_v16 = _dafny.UnicodeSeqOfUtf8Bytes("oc")
		var _rhs190 _dafny.Int = (_this).Fm6((_this).F17(), Companion_Default___.SafeModuloInt((_this).F11(), (_dafny.Zero).Minus((_this).F11())), globalState)
		_ = _rhs190
		var _rhs191 _dafny.Int = (_this).F11()
		_ = _rhs191
		var _rhs192 _dafny.Array = _1037_v15
		_ = _rhs192
		var _rhs193 _dafny.Array = _1037_v15
		_ = _rhs193
		var _rhs194 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bjar"), _1038_v16)).Cardinality())).Plus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F17())).Merge(Companion_Default___.Fm11(globalState))).Cardinality())
		_ = _rhs194
		var _lhs149 *GlobalState = globalState
		_ = _lhs149
		var _lhs150 *GlobalState = globalState
		_ = _lhs150
		var _lhs151 *GlobalState = globalState
		_ = _lhs151
		_lhs149.F1 = _rhs190
		_lhs150.F1 = _rhs191
		_1037_v15 = _rhs192
		_1037_v15 = _rhs193
		_lhs151.F1 = _rhs194
		var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))
		_ = _index118
		(_1037_v15).ArraySet1((_this).F17(), (_index118).Int())
		var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))
		_ = _index119
		(_1037_v15).ArraySet1((_this).F17(), (_index119).Int())
		(globalState).F6 = (_this).F17()
		var _1039_v17 _dafny.Map
		_ = _1039_v17
		_1039_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool)), (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool))
		if (_1039_v17).Contains(Companion_Default___.Fm12((_dafny.Zero).Minus((_this).F11()), _1038_v16, globalState)) {
			var _1040_v18 _dafny.Map
			_ = _1040_v18
			_1040_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F11())
			var _rhs195 _dafny.Int = (_1040_v18).Cardinality()
			_ = _rhs195
			var _rhs196 bool = (_this).F17()
			_ = _rhs196
			var _lhs152 *GlobalState = globalState
			_ = _lhs152
			var _lhs153 *GlobalState = globalState
			_ = _lhs153
			_lhs152.F1 = _rhs195
			_lhs153.F7 = _rhs196
			var _1041_v19 _dafny.Map
			_ = _1041_v19
			_1041_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool), (_this).F17())
			var _1042_v20 _dafny.Map
			_ = _1042_v20
			_1042_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1038_v16).Cardinality()), (_this).F17())
			var _1043_v21 _dafny.Map
			_ = _1043_v21
			_1043_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() bool {
				if (_1042_v20).Contains((_this).F11()) {
					return (_1042_v20).Get((_this).F11()).(bool)
				}
				return false
			})())
			var _1044_v22 _dafny.Sequence
			_ = _1044_v22
			_1044_v22 = _dafny.SeqOf((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F11(), (_this).F11())), (func() _dafny.Int {
				if (func() bool {
					if (_1041_v19).Contains((_this).F17()) {
						return (_1041_v19).Get((_this).F17()).(bool)
					}
					return (_this).F17()
				})() {
					return (_this).F11()
				}
				return (_this).F11()
			})(), (_this).F11(), (_dafny.Zero).Minus(((((_1043_v21).Update(p0, (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool))).Update(_dafny.CodePoint('w'), (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool))).Merge(_1043_v21)).Cardinality()), Companion_Default___.SafeDivisionInt((_this).F11(), _dafny.IntOfInt64(923)))
			var _1045_v23 _dafny.Array
			_ = _1045_v23
			var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw154
			_1045_v23 = _nw154
			var _1046_v24 _dafny.Array
			_ = _1046_v24
			var _nw155 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
			_ = _nw155
			_1046_v24 = _nw155
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_1045_v23), 0))
			_ = _index120
			(_1045_v23).ArraySet1(_1046_v24, (_index120).Int())
			var _1047_v26 _dafny.Map
			_ = _1047_v26
			_1047_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool))
			var _1048_v28 _dafny.Array
			_ = _1048_v28
			_1048_v28 = _1046_v24
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_1045_v23), 0))
			_ = _index121
			var _rhs197 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1044_v22, _1044_v22)
			_ = _rhs197
			var _rhs198 _dafny.Int = (_1044_v22).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_1044_v22).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs198
			var _rhs199 _dafny.Array = (func() _dafny.Array {
				if (func() bool {
					if (_this).F17() {
						return (_this).Fm4((_this).F11(), func() _dafny.Map {
							var _coll50 = _dafny.NewMapBuilder()
							_ = _coll50
							for _iter52 := _dafny.Iterate((_1047_v26).Keys().Elements()); ; {
								_compr_50, _ok52 := _iter52()
								if !_ok52 {
									break
								}
								var _1049_v25 _dafny.MultiSet
								_1049_v25 = interface{}(_compr_50).(_dafny.MultiSet)
								if (_1047_v26).Contains(_1049_v25) {
									_coll50.Add(_1049_v25, func() _dafny.Map {
										var _coll51 = _dafny.NewMapBuilder()
										_ = _coll51
										for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(257), _dafny.IntOfInt64(278))); ; {
											_compr_51, _ok53 := _iter53()
											if !_ok53 {
												break
											}
											var _1050_v27 _dafny.Int
											_1050_v27 = interface{}(_compr_51).(_dafny.Int)
											if ((_dafny.IntOfInt64(257)).Cmp(_1050_v27) <= 0) && ((_1050_v27).Cmp(_dafny.IntOfInt64(278)) < 0) {
												_coll51.Add((_1050_v27).Times(((_this).F12()).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_this).F17(), (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool))).Cardinality()))
											}
										}
										return _coll51.ToMap()
									}())
								}
							}
							return _coll50.ToMap()
						}(), (_this).F11(), (_this).F17(), globalState)
					}
					return (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool)
				})() {
					return _1046_v24
				}
				return (_1048_v28)
			})()
			_ = _rhs199
			var _rhs200 _dafny.Array = (func() _dafny.Array {
				if (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool) {
					return _1037_v15
				}
				return _1037_v15
			})()
			_ = _rhs200
			var _lhs154 *GlobalState = globalState
			_ = _lhs154
			var _lhs155 _dafny.Array = _1045_v23
			_ = _lhs155
			var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_1045_v23), 0))
			_ = _lhs156
			_1044_v22 = _rhs197
			_lhs154.F1 = _rhs198
			(_lhs155).ArraySet1(_rhs199, (_lhs156).Int())
			_1037_v15 = _rhs200
			if !((_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool)) {
				var _1051_v29 _dafny.Int
				_ = _1051_v29
				var _1052_v30 bool
				_ = _1052_v30
				var _1053_v31 _dafny.Int
				_ = _1053_v31
				var _out28 _dafny.Int
				_ = _out28
				var _out29 bool
				_ = _out29
				var _out30 _dafny.Int
				_ = _out30
				_out28, _out29, _out30 = (_this).M5((_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool), !((_this).F17()) || (true), (_this).F11(), globalState)
				_1051_v29 = _out28
				_1052_v30 = _out29
				_1053_v31 = _out30
				_1038_v16 = Companion_Default___.Fm1(globalState)
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))
				_ = _index122
				(_1037_v15).ArraySet1((_this).F17(), (_index122).Int())
				var _1054_v32 _dafny.Sequence
				_ = _1054_v32
				_1054_v32 = _dafny.SeqOf((_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool), (_this).F17(), (_this).F17())
				var _1055_v33 *C0
				_ = _1055_v33
				var _nw156 *C0 = New_C0_()
				_ = _nw156
				_nw156.Ctor__(_1054_v32, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(648))).Uint32(), func(coer77 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg77 _dafny.Int) interface{} {
						return coer77(arg77)
					}
				}((func(_1056_v20 _dafny.Map) func(_dafny.Int) _dafny.Int {
					return func(_1057_i1 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus((_1056_v20).Cardinality())
					}
				})(_1042_v20)))).Cardinality()))
				_1055_v33 = _nw156
				var _1058_v34 T1
				_ = _1058_v34
				var _nw157 *C3 = New_C3_()
				_ = _nw157
				_nw157.Ctor__(_dafny.IntOfUint32((_1038_v16).Cardinality()), _dafny.IntOfInt64(729), _1053_v31, (_this).F12())
				_1058_v34 = _nw157
				var _1059_v35 _dafny.Sequence
				_ = _1059_v35
				_1059_v35 = _dafny.SeqOf(_1058_v34, _1058_v34, _1058_v34, _1058_v34)
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))
				_ = _index123
				var _rhs201 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F11(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_1059_v35).Cardinality()))).Cardinality())
				_ = _rhs201
				var _rhs202 _dafny.Int = (func() _dafny.Int {
					if ((_this).F16()).Contains((_this).F11()) {
						return ((_this).F16()).Multiplicity((_this).F11())
					}
					return (_1053_v31).Times(_1051_v29)
				})()
				_ = _rhs202
				var _rhs203 bool = !(!((func() bool {
					if (_1042_v20).Contains((_1058_v34).F11()) {
						return (_1042_v20).Get((_1058_v34).F11()).(bool)
					}
					return _1052_v30
				})()))
				_ = _rhs203
				var _rhs204 bool = _1052_v30
				_ = _rhs204
				var _lhs157 *GlobalState = globalState
				_ = _lhs157
				var _lhs158 *GlobalState = globalState
				_ = _lhs158
				var _lhs159 _dafny.Array = _1037_v15
				_ = _lhs159
				var _lhs160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))
				_ = _lhs160
				_1053_v31 = _rhs201
				_lhs157.F1 = _rhs202
				_lhs158.F6 = _rhs203
				(_lhs159).ArraySet1(_rhs204, (_lhs160).Int())
			} else {
				var _1060_v36 _dafny.Array
				_ = _1060_v36
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_18
				var _nw158 _dafny.Array
				_ = _nw158
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw158 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Sequence = (func(_1061_v16 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1062_i2 _dafny.Int) _dafny.Sequence {
							return (_1061_v16)
						}
					})(_1038_v16)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw158 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw158).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw158).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_1060_v36 = _nw158
				_1060_v36 = _1060_v36
				var _1063_v37 _dafny.Map
				_ = _1063_v37
				_1063_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1044_v22, (_this).F11())
				var _1064_v38 _dafny.Map
				_ = _1064_v38
				_1064_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1063_v37, (_this).F11())
				(globalState).F1 = ((_dafny.Zero).Minus((_this).F11())).Times((func() _dafny.Int {
					if (_1064_v38).Contains(_1063_v37) {
						return (_1064_v38).Get(_1063_v37).(_dafny.Int)
					}
					return (_this).F11()
				})())
				(globalState).F6 = (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool)
				var _1065_v39 *C2
				_ = _1065_v39
				var _nw159 *C2 = New_C2_()
				_ = _nw159
				_nw159.Ctor__(_dafny.CodePoint('d'), Companion_Default___.SafeDivisionInt((_this).F11(), (_this).F11()))
				_1065_v39 = _nw159
				(globalState).F6 = ((_this).F16()).Contains((_this).F11())
			}
			var _1066_v41 _dafny.Set
			_ = _1066_v41
			_1066_v41 = _dafny.SetOf(_dafny.IntOfUint32((_1044_v22).Cardinality()), (func() _dafny.Map {
				var _coll52 = _dafny.NewMapBuilder()
				_ = _coll52
				for _iter54 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-494))).Uint32(), func(coer78 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg78 _dafny.Int) interface{} {
						return coer78(arg78)
					}
				}(func(_1067_i3 _dafny.Int) _dafny.Int {
					return (_this).F11()
				}))).Elements()); ; {
					_compr_52, _ok54 := _iter54()
					if !_ok54 {
						break
					}
					var _1068_v40 _dafny.Int
					_1068_v40 = interface{}(_compr_52).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-494))).Uint32(), func(coer79 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg79 _dafny.Int) interface{} {
							return coer79(arg79)
						}
					}(func(_1067_i3 _dafny.Int) _dafny.Int {
						return (_this).F11()
					})), _1068_v40) {
						_coll52.Add((_1068_v40).Minus((_dafny.MultiSetOf((_this).F11())).Cardinality()), true)
					}
				}
				return _coll52.ToMap()
			}()).Cardinality())
			(globalState).F1 = ((_1066_v41).Union(_1066_v41)).Cardinality()
			var _1069_v42 *C3
			_ = _1069_v42
			var _nw160 *C3 = New_C3_()
			_ = _nw160
			_nw160.Ctor__(_dafny.IntOfUint32((_1038_v16).Cardinality()), (_this).F11(), (_this).F11(), (_this).F12())
			_1069_v42 = _nw160
			var _1070_v43 _dafny.Sequence
			_ = _1070_v43
			_1070_v43 = _dafny.SeqOf(_1069_v42, _1069_v42)
			(globalState).F1 = _dafny.IntOfUint32((_1070_v43).Cardinality())
		} else {
			var _1071_v44 _dafny.Array
			_ = _1071_v44
			var _nw161 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
			_ = _nw161
			_1071_v44 = _nw161
			var _1072_v45 _dafny.Sequence
			_ = _1072_v45
			_1072_v45 = _dafny.SeqOf((_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool))
			var _1073_v46 D11
			_ = _1073_v46
			_1073_v46 = Companion_D11_.Create_DC24_((_this).F11(), _1072_v45, p0)
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))
			_ = _index124
			var _rhs205 _dafny.Int = (_1073_v46).Dtor_cf34()
			_ = _rhs205
			var _rhs206 _dafny.Int = (_this).F11()
			_ = _rhs206
			var _rhs207 bool = (_this).F17()
			_ = _rhs207
			var _rhs208 _dafny.Int = (_this).F11()
			_ = _rhs208
			var _rhs209 _dafny.Array = _1071_v44
			_ = _rhs209
			var _lhs161 *GlobalState = globalState
			_ = _lhs161
			var _lhs162 *GlobalState = globalState
			_ = _lhs162
			var _lhs163 _dafny.Array = _1037_v15
			_ = _lhs163
			var _lhs164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))
			_ = _lhs164
			var _lhs165 *GlobalState = globalState
			_ = _lhs165
			_lhs161.F1 = _rhs205
			_lhs162.F1 = _rhs206
			(_lhs163).ArraySet1(_rhs207, (_lhs164).Int())
			_lhs165.F1 = _rhs208
			_1071_v44 = _rhs209
			var _1074_v47 _dafny.Map
			_ = _1074_v47
			_1074_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F11()).Plus((_this).F11()), ((_this).F11()).Plus((_this).F11()))
			_1074_v47 = (_1074_v47).Update((_this).F11(), (_this).F11())
			var _1075_v48 _dafny.Array
			_ = _1075_v48
			var _nw162 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw162
			_1075_v48 = _nw162
			var _1076_v49 D6
			_ = _1076_v49
			_1076_v49 = Companion_D6_.Create_DC11_((_this).F11(), (_this).F11(), (_dafny.Zero).Minus((_this).F11()))
			var _1077_v50 _dafny.Map
			_ = _1077_v50
			_1077_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1076_v49, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("go"), (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("go")).Cardinality()))).Uint32(), p0))
			var _1078_v51 _dafny.MultiSet
			_ = _1078_v51
			_1078_v51 = _dafny.MultiSetOf((func() _dafny.Sequence {
				if (_1077_v50).Contains(_1076_v49) {
					return (_1077_v50).Get(_1076_v49).(_dafny.Sequence)
				}
				return _1038_v16
			})())
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_1075_v48), 0))
			_ = _index125
			(_1075_v48).ArraySet1((_1078_v51).Cardinality(), (_index125).Int())
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_1075_v48), 0))
			_ = _index126
			(_1075_v48).ArraySet1((_dafny.IntOfInt64(-940)).Times((_dafny.Zero).Minus((_this).F11())), (_index126).Int())
			var _1079_v52 _dafny.Set
			_ = _1079_v52
			_1079_v52 = _dafny.SetOf((_this).F11(), (_this).F11())
			if (((_this).F11()).Minus((_1079_v52).Cardinality())).Cmp(_dafny.IntOfUint32((_1038_v16).Cardinality())) > 0 {
				(globalState).F1 = _dafny.IntOfInt64(-492)
				var _1080_v53 _dafny.MultiSet
				_ = _1080_v53
				_1080_v53 = _dafny.MultiSetOf((_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool))
				var _1081_v54 *C1
				_ = _1081_v54
				var _nw163 *C1 = New_C1_()
				_ = _nw163
				_nw163.Ctor__((_dafny.MultiSetFromSeq(_1072_v45)).IsSubsetOf(_1080_v53), (_this).F17())
				_1081_v54 = _nw163
				var _1082_v55 _dafny.Sequence
				_ = _1082_v55
				_1082_v55 = _dafny.SeqOf((_1075_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_1075_v48), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-48), (_1075_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_1075_v48), 0))).Int()).(_dafny.Int))
				(globalState).F7 = (Companion_Default___.Fm0(globalState)) && (((_1082_v55).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_1082_v55).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((_this).F11()) > 0)
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))
				_ = _index127
				(_1037_v15).ArraySet1((_this).F17(), (_index127).Int())
				(globalState).F7 = (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool)
			} else {
				var _1083_v56 _dafny.Map
				_ = _1083_v56
				_1083_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F17())
				var _1084_v57 D8
				_ = _1084_v57
				_1084_v57 = Companion_D8_.Create_DC15_(_1083_v56)
				var _1085_v58 D8
				_ = _1085_v58
				_1085_v58 = Companion_D8_.Create_DC17_(_1084_v57)
				var _1086_v59 _dafny.Map
				_ = _1086_v59
				_1086_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1085_v58, (_1075_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_1075_v48), 0))).Int()).(_dafny.Int))
				var _1087_v60 _dafny.Sequence
				_ = _1087_v60
				_1087_v60 = _dafny.SeqOf(_1086_v59, _1086_v59, _1086_v59, _1086_v59)
				var _1088_v61 _dafny.MultiSet
				_ = _1088_v61
				_1088_v61 = _dafny.MultiSetOf(Companion_Default___.Fm0(globalState))
				var _1089_v62 D13
				_ = _1089_v62
				_1089_v62 = Companion_D13_.Create_DC27_(_1086_v59)
				var _1090_v63 _dafny.Map
				_ = _1090_v63
				_1090_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _1087_v60)
				var _1091_v64 _dafny.Map
				_ = _1091_v64
				_1091_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(11), (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool))
				var _1092_v65 _dafny.Array
				_ = _1092_v65
				var _nwElement0_46 _dafny.Sequence = _1087_v60
				_ = _nwElement0_46
				var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(15))
				_ = _nw164
				(_nw164).ArraySet1(_nwElement0_46, 0)
				(_nw164).ArraySet1(_1087_v60, 1)
				(_nw164).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1087_v60, (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1088_v61).Contains(true) {
						return (_1088_v61).Multiplicity(true)
					}
					return (_this).F11()
				})(), _dafny.IntOfUint32((_1087_v60).Cardinality()))).Uint32(), _1086_v59), _1087_v60), 2)
				(_nw164).ArraySet1(_1087_v60, 3)
				(_nw164).ArraySet1(_1087_v60, 4)
				(_nw164).ArraySet1(_dafny.SeqOf(_1086_v59, _1086_v59, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1085_v58, (_this).F11())), 5)
				(_nw164).ArraySet1(_1087_v60, 6)
				(_nw164).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1089_v62).Dtor_cf39(), _1086_v59, (_1089_v62).Dtor_cf39()), _dafny.SeqOf(_1086_v59)), (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1089_v62).Dtor_cf39(), _1086_v59, (_1089_v62).Dtor_cf39()), _dafny.SeqOf(_1086_v59))).Cardinality()))).Uint32(), _1086_v59), 7)
				(_nw164).ArraySet1((func() _dafny.Sequence {
					if (_1090_v63).Contains((func() bool {
						if (_1091_v64).Contains((_this).F11()) {
							return (_1091_v64).Get((_this).F11()).(bool)
						}
						return (_this).F17()
					})()) {
						return (_1090_v63).Get((func() bool {
							if (_1091_v64).Contains((_this).F11()) {
								return (_1091_v64).Get((_this).F11()).(bool)
							}
							return (_this).F17()
						})()).(_dafny.Sequence)
					}
					return Companion_Default___.Fm44(!((_this).F17()), (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool), globalState)
				})(), 8)
				(_nw164).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm44((_this).F17(), (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool), globalState), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.IntOfUint32((Companion_Default___.Fm44((_this).F17(), (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool), globalState)).Cardinality()))).Uint32(), _1086_v59), _1087_v60), 9)
				(_nw164).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1087_v60, _1087_v60), 10)
				(_nw164).ArraySet1(_dafny.SeqOf(_1086_v59, _1086_v59), 11)
				(_nw164).ArraySet1(Companion_Default___.Fm44(!((_this).F17()), (_this).F17(), globalState), 12)
				(_nw164).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1087_v60, _1087_v60), 13)
				(_nw164).ArraySet1(_dafny.SeqOf((_1086_v59).Update(Companion_D8_.Create_DC17_(_1084_v57), (_dafny.Zero).Minus(_dafny.IntOfInt64(-8)))), 14)
				_1092_v65 = _nw164
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_1092_v65), 0))
				_ = _index128
				(_1092_v65).ArraySet1((func() _dafny.Sequence {
					if (_1090_v63).Contains(!((_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool))) {
						return (_1090_v63).Get(!((_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool))).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-255))).Uint32(), func(coer80 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg80 _dafny.Int) interface{} {
							return coer80(arg80)
						}
					}((func(_1093_v59 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_1094_i4 _dafny.Int) _dafny.Map {
							return _1093_v59
						}
					})(_1086_v59)))
				})(), (_index128).Int())
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_1092_v65), 0))
				_ = _index129
				(_1092_v65).ArraySet1(_1087_v60, (_index129).Int())
				var _1095_v66 _dafny.Array
				_ = _1095_v66
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_19
				var _nw165 _dafny.Array
				_ = _nw165
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw165 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) D8 = (func(_1096_v58 D8) func(_dafny.Int) D8 {
						return func(_1097_i5 _dafny.Int) D8 {
							return _1096_v58
						}
					})(_1085_v58)
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw165 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw165).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw165).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_1095_v66 = _nw165
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_1095_v66), 0))
				_ = _index130
				(_1095_v66).ArraySet1(_1085_v58, (_index130).Int())
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_1095_v66), 0))
				_ = _index131
				(_1095_v66).ArraySet1(_1085_v58, (_index131).Int())
				_1037_v15 = _1037_v15
				(globalState).F6 = !(((_this).F11()).Cmp((_1075_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_1075_v48), 0))).Int()).(_dafny.Int)) >= 0)
				var _1098_v67 _dafny.Array
				_ = _1098_v67
				var _nw166 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
				_ = _nw166
				_1098_v67 = _nw166
				var _1099_v68 _dafny.Map
				_ = _1099_v68
				_1099_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _1098_v67)
				(globalState).F1 = ((_1099_v68).Update(((_this).F17()) || ((_this).F17()), _1098_v67)).Cardinality()
			}
			(globalState).F7 = (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool)
		}
		var _1100_v70 _dafny.Map
		_ = _1100_v70
		_1100_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F16()).Update((_this).F11(), Companion_Default___.Abs((_this).F11())), (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool))
		var _1101_v72 _dafny.MultiSet
		_ = _1101_v72
		_1101_v72 = _dafny.MultiSetOf((_this).Fm4((_this).F11(), func() _dafny.Map {
			var _coll53 = _dafny.NewMapBuilder()
			_ = _coll53
			for _iter55 := _dafny.Iterate((_1100_v70).Keys().Elements()); ; {
				_compr_53, _ok55 := _iter55()
				if !_ok55 {
					break
				}
				var _1102_v69 _dafny.MultiSet
				_1102_v69 = interface{}(_compr_53).(_dafny.MultiSet)
				if (_1100_v70).Contains(_1102_v69) {
					_coll53.Add(_1102_v69, func() _dafny.Map {
						var _coll54 = _dafny.NewMapBuilder()
						_ = _coll54
						for _iter56 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(958), _dafny.IntOfInt64(898))); ; {
							_compr_54, _ok56 := _iter56()
							if !_ok56 {
								break
							}
							var _1103_v71 _dafny.Int
							_1103_v71 = interface{}(_compr_54).(_dafny.Int)
							if ((_dafny.IntOfInt64(958)).Cmp(_1103_v71) <= 0) && ((_1103_v71).Cmp(_dafny.IntOfInt64(898)) < 0) {
								_coll54.Add(Companion_Default___.SafeModuloInt(_1103_v71, (_this).F11()), _dafny.IntOfInt64(-113))
							}
						}
						return _coll54.ToMap()
					}())
				}
			}
			return _coll53.ToMap()
		}(), _dafny.IntOfInt64(893), (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool), globalState))
		var _1104_v73 D0
		_ = _1104_v73
		_1104_v73 = Companion_D0_.Create_DC0_((_this).F11())
		var _1105_v74 _dafny.Sequence
		_ = _1105_v74
		_1105_v74 = _dafny.SeqOf(true, (_1037_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1037_v15), 0))).Int()).(bool), (_this).F17(), (_this).F17())
		r0 = (_1101_v72).Update((_this).Fm5(_1104_v73, _1105_v74, (_this).F11(), _dafny.CodePoint('x'), globalState), Companion_Default___.Abs(((_this).F11()).Times(((_this).F12()).Cardinality())))
		return r0
	}
}
func (_this *C4) M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Map) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _1106_i0 _dafny.Int
		_ = _1106_i0
		_1106_i0 = _dafny.Zero
		{
			for (_this).F17() {
				{
					if (_1106_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_1106_i0 = (_1106_i0).Plus(_dafny.One)
					var _1107_v0 _dafny.CodePoint
					_ = _1107_v0
					_1107_v0 = _dafny.CodePoint('c')
					var _1108_v1 *C2
					_ = _1108_v1
					var _nw167 *C2 = New_C2_()
					_ = _nw167
					_nw167.Ctor__(_1107_v0, _dafny.IntOfInt64(813))
					_1108_v1 = _nw167
					(globalState).F1 = (_this).F11()
					var _1109_v2 _dafny.Sequence
					_ = _1109_v2
					_1109_v2 = _dafny.UnicodeSeqOfUtf8Bytes("ulwlw")
					var _1110_v3 _dafny.Sequence
					_ = _1110_v3
					_1110_v3 = _dafny.SeqOf(true)
					var _1111_v4 _dafny.Array
					_ = _1111_v4
					var _nwElement0_47 _dafny.Sequence = Companion_Default___.Fm12((_this).F11(), _1109_v2, globalState)
					_ = _nwElement0_47
					var _nw168 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(9))
					_ = _nw168
					(_nw168).ArraySet1(_nwElement0_47, 0)
					(_nw168).ArraySet1(_dafny.Companion_Sequence_.Update(_1110_v3, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_1110_v3).Cardinality()))).Uint32(), (_this).F17()), 1)
					(_nw168).ArraySet1(_1110_v3, 2)
					(_nw168).ArraySet1(_1110_v3, 3)
					(_nw168).ArraySet1(_dafny.Companion_Sequence_.Update(_1110_v3, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_1110_v3).Cardinality()))).Uint32(), (_this).F17()), 4)
					(_nw168).ArraySet1(_1110_v3, 5)
					(_nw168).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1110_v3, _1110_v3), 6)
					(_nw168).ArraySet1(_1110_v3, 7)
					(_nw168).ArraySet1(_1110_v3, 8)
					_1111_v4 = _nw168
					var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_1111_v4), 0))
					_ = _index132
					(_1111_v4).ArraySet1(_1110_v3, (_index132).Int())
					var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_1111_v4), 0))
					_ = _index133
					(_1111_v4).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1110_v3, _1110_v3), (_index133).Int())
					var _1112_v5 _dafny.Map
					_ = _1112_v5
					_1112_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F17())
					var _1113_v6 _dafny.Map
					_ = _1113_v6
					_1113_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F17())
					var _1114_v7 _dafny.Array
					_ = _1114_v7
					var _nwElement0_48 _dafny.Int = (_1112_v5).Cardinality()
					_ = _nwElement0_48
					var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(9))
					_ = _nw169
					(_nw169).ArraySet1(_nwElement0_48, 0)
					(_nw169).ArraySet1(p0, 1)
					(_nw169).ArraySet1((_this).F11(), 2)
					(_nw169).ArraySet1(p0, 3)
					(_nw169).ArraySet1((_1113_v6).Cardinality(), 4)
					(_nw169).ArraySet1(_dafny.IntOfInt64(815), 5)
					(_nw169).ArraySet1((_this).F11(), 6)
					(_nw169).ArraySet1(((_this).F12()).Cardinality(), 7)
					(_nw169).ArraySet1(_dafny.IntOfInt64(745), 8)
					_1114_v7 = _nw169
					var _1115_v8 _dafny.Map
					_ = _1115_v8
					_1115_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1114_v7, !(((_this).F11()).Cmp(_dafny.IntOfUint32(((_1111_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_1111_v4), 0))).Int()).(_dafny.Sequence)).Cardinality())) >= 0))
					_1115_v8 = (_1115_v8).Update(_1114_v7, (_this).F17())
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
		var _1116_v9 _dafny.Sequence
		_ = _1116_v9
		_1116_v9 = _dafny.UnicodeSeqOfUtf8Bytes("g")
		var _1117_v10 *C3
		_ = _1117_v10
		var _nw170 *C3 = New_C3_()
		_ = _nw170
		_nw170.Ctor__((Companion_D13_.Create_DC28_(_1116_v9, (_this).F11(), (_this).F17())).Dtor_cf41(), (p0).Plus(_dafny.IntOfUint32((_1116_v9).Cardinality())), Companion_Default___.Fm2(p0, p0, globalState), (_this).F12())
		_1117_v10 = _nw170
		if (false) || (((_this).F17()) || ((_this).F17())) {
			var _1118_v11 _dafny.Sequence
			_ = _1118_v11
			_1118_v11 = _dafny.SeqOf((_this).F16(), (_this).F16())
			if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(185))).Uint32(), func(coer81 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg81 _dafny.Int) interface{} {
					return coer81(arg81)
				}
			}(func(_1119_i1 _dafny.Int) _dafny.MultiSet {
				return (_this).F16()
			})), _1118_v11) {
				var _1120_v13 _dafny.Sequence
				_ = _1120_v13
				_1120_v13 = _dafny.SeqOf((_dafny.Zero).Minus((_1117_v10).F22()))
				var _1121_v14 _dafny.Set
				_ = _1121_v14
				_1121_v14 = _dafny.SetOf((_this).F11(), (_1117_v10).F22())
				var _1122_v15 _dafny.Sequence
				_ = _1122_v15
				_1122_v15 = _dafny.SeqOf((_this).F17())
				var _1123_v16 D0
				_ = _1123_v16
				_1123_v16 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(-405))
				var _1124_v17 _dafny.Map
				_ = _1124_v17
				_1124_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(811), (_this).F11())
				var _1125_v18 _dafny.CodePoint
				_ = _1125_v18
				_1125_v18 = _dafny.CodePoint('u')
				var _1126_v19 _dafny.Array
				_ = _1126_v19
				var _nwElement0_49 bool = (_this).F17()
				_ = _nwElement0_49
				var _nw171 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(28))
				_ = _nw171
				(_nw171).ArraySet1(_nwElement0_49, 0)
				(_nw171).ArraySet1((_this).F17(), 1)
				(_nw171).ArraySet1(!((_this).F17()), 2)
				(_nw171).ArraySet1((_1117_v10.F23).Cmp((_dafny.Zero).Minus(_1117_v10.F23)) >= 0, 3)
				(_nw171).ArraySet1(((_this).F11()).Cmp((func() _dafny.Set {
					var _coll55 = _dafny.NewBuilder()
					_ = _coll55
					for _iter57 := _dafny.Iterate((_1118_v11).Elements()); ; {
						_compr_55, _ok57 := _iter57()
						if !_ok57 {
							break
						}
						var _1127_v12 _dafny.MultiSet
						_1127_v12 = interface{}(_compr_55).(_dafny.MultiSet)
						if _dafny.Companion_Sequence_.Contains(_1118_v11, _1127_v12) {
							_coll55.Add(_1127_v12)
						}
					}
					return _coll55.ToSet()
				}()).Cardinality()) != 0, 4)
				(_nw171).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1116_v9, _1116_v9), 5)
				(_nw171).ArraySet1(!((_this).F17()) || (!(!(!((_this).F17())))), 6)
				(_nw171).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1120_v13, _1120_v13), 7)
				(_nw171).ArraySet1((_this).F17(), 8)
				(_nw171).ArraySet1((_this).F17(), 9)
				(_nw171).ArraySet1((_dafny.SetOf(p0)).IsDisjointFrom(_1121_v14), 10)
				(_nw171).ArraySet1((_this).F17(), 11)
				(_nw171).ArraySet1((_this).F17(), 12)
				(_nw171).ArraySet1(true, 13)
				(_nw171).ArraySet1(!_dafny.Companion_Sequence_.Contains(_1122_v15, (_this).F17()), 14)
				(_nw171).ArraySet1((_this).F17(), 15)
				(_nw171).ArraySet1((_this).F17(), 16)
				(_nw171).ArraySet1((_this).F17(), 17)
				(_nw171).ArraySet1((_this).F17(), 18)
				(_nw171).ArraySet1((_this).F17(), 19)
				(_nw171).ArraySet1((_this).F17(), 20)
				(_nw171).ArraySet1(!(true), 21)
				(_nw171).ArraySet1((_this).F17(), 22)
				(_nw171).ArraySet1((_1117_v10).Fm5(_1123_v16, _1122_v15, (_1120_v13).Select((Companion_Default___.SafeIndex((_1124_v17).Cardinality(), _dafny.IntOfUint32((_1120_v13).Cardinality()))).Uint32()).(_dafny.Int), _1125_v18, globalState), 23)
				(_nw171).ArraySet1(!((_this).F17()) || (true), 24)
				(_nw171).ArraySet1((_this).F17(), 25)
				(_nw171).ArraySet1((_this).F17(), 26)
				(_nw171).ArraySet1((_dafny.IntOfUint32((_1116_v9).Cardinality())).Cmp(_1117_v10.F23) > 0, 27)
				_1126_v19 = _nw171
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1126_v19), 0))
				_ = _index134
				(_1126_v19).ArraySet1((_this).Fm4((_1117_v10).Fm6((_this).F17(), _1117_v10.F23, globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), _1124_v17), (_1117_v10).F22(), true, globalState), (_index134).Int())
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1126_v19), 0))
				_ = _index135
				(_1126_v19).ArraySet1((Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1116_v9).Cardinality()), _1117_v10.F23)).Cmp(_1117_v10.F23) != 0, (_index135).Int())
				_1125_v18 = _1125_v18
				var _1128_v20 _dafny.Array
				_ = _1128_v20
				var _nw172 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
				_ = _nw172
				_1128_v20 = _nw172
				_1128_v20 = _1128_v20
				var _1129_v21 _dafny.Sequence
				_ = _1129_v21
				_1129_v21 = _1116_v9
				_1129_v21 = _1129_v21
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1126_v19), 0))
				_ = _index136
				(_1126_v19).ArraySet1((_1117_v10).Fm5(_1123_v16, _1122_v15, p0, Companion_Default___.Fm39(p0, false, p0, globalState), globalState), (_index136).Int())
			} else {
				var _1130_v22 _dafny.Map
				_ = _1130_v22
				_1130_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1117_v10.F23, (_1117_v10).F22())
				var _1131_v23 _dafny.Sequence
				_ = _1131_v23
				_1131_v23 = _dafny.SeqOf((_1130_v22).Cardinality())
				var _1132_v24 _dafny.Sequence
				_ = _1132_v24
				_1132_v24 = _dafny.SeqOf(((_1117_v10).F22()).Plus(_dafny.IntOfUint32((_1131_v23).Cardinality())))
				var _1133_v25 D10
				_ = _1133_v25
				_1133_v25 = Companion_D10_.Create_DC21_((_this).F17(), (_this).F17())
				var _1134_v28 _dafny.Map
				_ = _1134_v28
				_1134_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F17())
				var _rhs210 _dafny.Int = (_1117_v10).F22()
				_ = _rhs210
				var _rhs211 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(78))).Uint32(), func(coer82 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg82 _dafny.Int) interface{} {
						return coer82(arg82)
					}
				}((func(_1135_v10 *C3) func(_dafny.Int) _dafny.Int {
					return func(_1136_i2 _dafny.Int) _dafny.Int {
						return (func() _dafny.Map {
							var _coll56 = _dafny.NewMapBuilder()
							_ = _coll56
							for _iter58 := _dafny.Iterate((func() _dafny.Set {
								var _coll57 = _dafny.NewBuilder()
								_ = _coll57
								for _iter59 := _dafny.Iterate((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F17()))).Elements()); ; {
									_compr_57, _ok59 := _iter59()
									if !_ok59 {
										break
									}
									var _1137_v27 _dafny.Map
									_1137_v27 = interface{}(_compr_57).(_dafny.Map)
									if (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F17()))).Contains(_1137_v27) {
										_coll57.Add(_1137_v27)
									}
								}
								return _coll57.ToSet()
							}()).Elements()); ; {
								_compr_56, _ok58 := _iter58()
								if !_ok58 {
									break
								}
								var _1138_v26 _dafny.Map
								_1138_v26 = interface{}(_compr_56).(_dafny.Map)
								if (func() _dafny.Set {
									var _coll58 = _dafny.NewBuilder()
									_ = _coll58
									for _iter60 := _dafny.Iterate((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F17()))).Elements()); ; {
										_compr_58, _ok60 := _iter60()
										if !_ok60 {
											break
										}
										var _1139_v27 _dafny.Map
										_1139_v27 = interface{}(_compr_58).(_dafny.Map)
										if (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F17()))).Contains(_1139_v27) {
											_coll58.Add(_1139_v27)
										}
									}
									return _coll58.ToSet()
								}()).Contains(_1138_v26) {
									_coll56.Add(_1138_v26, _1135_v10.F23)
								}
							}
							return _coll56.ToMap()
						}()).Cardinality()
					}
				})(_1117_v10))), _1132_v24), _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_1117_v10).F22(), (_1117_v10).F22(), _dafny.IntOfUint32((_dafny.SeqOf(_1117_v10.F23, _1117_v10.F23, _1117_v10.F23)).Cardinality()), _dafny.IntOfInt64(-313)), (Companion_Default___.SafeIndex((_1134_v28).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_1117_v10).F22(), (_1117_v10).F22(), _dafny.IntOfUint32((_dafny.SeqOf(_1117_v10.F23, _1117_v10.F23, _1117_v10.F23)).Cardinality()), _dafny.IntOfInt64(-313))).Cardinality()))).Uint32(), (_1117_v10).F22()))
				_ = _rhs211
				var _rhs212 D10 = _1133_v25
				_ = _rhs212
				var _lhs166 *C3 = _1117_v10
				_ = _lhs166
				_lhs166.F23 = _rhs210
				_1132_v24 = _rhs211
				_1133_v25 = _rhs212
				var _1140_v29 *C3
				_ = _1140_v29
				var _nw173 *C3 = New_C3_()
				_ = _nw173
				_nw173.Ctor__((_dafny.Zero).Minus((_this).Fm3((_this).F11(), (_this).F17(), true, globalState)), (_this).F11(), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1116_v9).Cardinality()), (_dafny.Zero).Minus((_1134_v28).Cardinality())), (_this).F12())
				_1140_v29 = _nw173
				var _1141_v30 _dafny.Array
				_ = _1141_v30
				var _nw174 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(25))
				_ = _nw174
				_1141_v30 = _nw174
				var _1142_v31 _dafny.MultiSet
				_ = _1142_v31
				_1142_v31 = _dafny.MultiSetOf((_this).F17())
				var _1143_v32 T1
				_ = _1143_v32
				var _nw175 *C3 = New_C3_()
				_ = _nw175
				_nw175.Ctor__((_1142_v31).Cardinality(), (_1140_v29).F22(), _dafny.IntOfUint32((_1116_v9).Cardinality()), (_this).F12())
				_1143_v32 = _nw175
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_1141_v30), 0))
				_ = _index137
				(_1141_v30).ArraySet1(_1143_v32, (_index137).Int())
				var _1144_v33 _dafny.Sequence
				_ = _1144_v33
				_1144_v33 = _dafny.SeqOf(_1116_v9)
				var _1145_v34 _dafny.Sequence
				_ = _1145_v34
				_1145_v34 = _dafny.SeqOf(_1131_v23, _1131_v23)
				var _1146_v35 _dafny.Array
				_ = _1146_v35
				var _nwElement0_50 _dafny.Int = _1140_v29.F23
				_ = _nwElement0_50
				var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(7))
				_ = _nw176
				(_nw176).ArraySet1(_nwElement0_50, 0)
				(_nw176).ArraySet1((_dafny.Zero).Minus(_1117_v10.F23), 1)
				(_nw176).ArraySet1((p0).Plus((_1117_v10).F22()), 2)
				(_nw176).ArraySet1((_1117_v10).F22(), 3)
				(_nw176).ArraySet1(_dafny.IntOfUint32((_1144_v33).Cardinality()), 4)
				(_nw176).ArraySet1((func() _dafny.Int {
					if (_this).F17() {
						return (_1117_v10).F22()
					}
					return _1117_v10.F23
				})(), 5)
				(_nw176).ArraySet1(Companion_Default___.SafeModuloInt((_1143_v32).F11(), _dafny.IntOfUint32((_1145_v34).Cardinality())), 6)
				_1146_v35 = _nw176
				var _1147_v36 _dafny.CodePoint
				_ = _1147_v36
				_1147_v36 = _dafny.CodePoint('o')
				var _1148_v37 _dafny.Map
				_ = _1148_v37
				_1148_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1147_v36, (_this).F11())
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1146_v35), 0))
				_ = _index138
				(_1146_v35).ArraySet1(((_1148_v37).Merge(_1148_v37)).Cardinality(), (_index138).Int())
				var _1149_v39 _dafny.Set
				_ = _1149_v39
				_1149_v39 = _dafny.SetOf((_1116_v9).Select((Companion_Default___.SafeIndex((_1117_v10).F22(), _dafny.IntOfUint32((_1116_v9).Cardinality()))).Uint32()).(_dafny.CodePoint), _1147_v36, _1147_v36)
				var _1150_v40 _dafny.Map
				_ = _1150_v40
				_1150_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll59 = _dafny.NewMapBuilder()
					_ = _coll59
					for _iter61 := _dafny.Iterate((_1149_v39).Elements()); ; {
						_compr_59, _ok61 := _iter61()
						if !_ok61 {
							break
						}
						var _1151_v38 _dafny.CodePoint
						_1151_v38 = interface{}(_compr_59).(_dafny.CodePoint)
						if (_1149_v39).Contains(_1151_v38) {
							_coll59.Add(_1151_v38, (_this).F17())
						}
					}
					return _coll59.ToMap()
				}(), (_this).F17())
				var _1152_v41 _dafny.Map
				_ = _1152_v41
				_1152_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1147_v36, !((_this).F17()))
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_1141_v30), 0))
				_ = _index139
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1146_v35), 0))
				_ = _index140
				var _rhs213 T1 = _1143_v32
				_ = _rhs213
				var _rhs214 _dafny.Int = (_dafny.Zero).Minus((_1143_v32).F11())
				_ = _rhs214
				var _rhs215 _dafny.Int = (_1130_v22).Cardinality()
				_ = _rhs215
				var _rhs216 _dafny.Map = (_1150_v40).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1152_v41, (_this).F17())).Update((_1152_v41).Update(_1147_v36, (_this).F17()), (_this).F17()))
				_ = _rhs216
				var _lhs167 _dafny.Array = _1141_v30
				_ = _lhs167
				var _lhs168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_1141_v30), 0))
				_ = _lhs168
				var _lhs169 *C3 = _1117_v10
				_ = _lhs169
				var _lhs170 _dafny.Array = _1146_v35
				_ = _lhs170
				var _lhs171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1146_v35), 0))
				_ = _lhs171
				(_lhs167).ArraySet1(_rhs213, (_lhs168).Int())
				_lhs169.F23 = _rhs214
				(_lhs170).ArraySet1(_rhs215, (_lhs171).Int())
				_1150_v40 = _rhs216
				var _1153_v42 _dafny.Array
				_ = _1153_v42
				var _nw177 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
				_ = _nw177
				_1153_v42 = _nw177
				var _1154_v43 D0
				_ = _1154_v43
				_1154_v43 = Companion_D0_.Create_DC0_((_1146_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1146_v35), 0))).Int()).(_dafny.Int))
				var _1155_v44 _dafny.Set
				_ = _1155_v44
				_1155_v44 = _dafny.SetOf((_1146_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1146_v35), 0))).Int()).(_dafny.Int), (_1154_v43).Dtor_cf0())
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1153_v42), 0))
				_ = _index141
				(_1153_v42).ArraySet1((_1117_v10.F23).Cmp((_1155_v44).Cardinality()) == 0, (_index141).Int())
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1153_v42), 0))
				_ = _index142
				(_1153_v42).ArraySet1((_this).F17(), (_index142).Int())
				var _1156_v45 _dafny.Map
				_ = _1156_v45
				_1156_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F11()), (func() bool {
					if (_1153_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1153_v42), 0))).Int()).(bool) {
						return (_this).F17()
					}
					return (_this).F17()
				})())
				_1156_v45 = (_1156_v45).Update(_dafny.IntOfUint32((_1116_v9).Cardinality()), Companion_Default___.Fm0(globalState))
			}
			var _1157_v46 *C2
			_ = _1157_v46
			var _nw178 *C2 = New_C2_()
			_ = _nw178
			_nw178.Ctor__(_dafny.CodePoint('g'), (_this).F11())
			_1157_v46 = _nw178
			var _1158_v47 _dafny.Sequence
			_ = _1158_v47
			_1158_v47 = _dafny.SeqOf(_1117_v10.F23, _dafny.IntOfUint32((_1116_v9).Cardinality()), _1117_v10.F23)
			(globalState).F6 = (func() bool {
				if (Companion_Default___.Fm45(false, (_this).F17(), (_1117_v10).F22(), _dafny.IntOfUint32((_1116_v9).Cardinality()), globalState)).IsProperSubsetOf(_dafny.MultiSetFromSeq(_1158_v47)) {
					return (_this).F17()
				}
				return (_this).F17()
			})()
			var _1159_v48 D13
			_ = _1159_v48
			_1159_v48 = Companion_D13_.Create_DC28_(_1116_v9, _dafny.IntOfInt64(-728), (_this).F17())
			var _pat_let_tv37 = globalState
			_ = _pat_let_tv37
			var _pat_let_tv38 = _1116_v9
			_ = _pat_let_tv38
			_1159_v48 = func(_pat_let27_0 D13) D13 {
				return func(_1160_dt__update__tmp_h0 D13) D13 {
					return func(_pat_let28_0 _dafny.Sequence) D13 {
						return func(_1161_dt__update_hcf40_h0 _dafny.Sequence) D13 {
							return Companion_D13_.Create_DC28_(_1161_dt__update_hcf40_h0, (_1160_dt__update__tmp_h0).Dtor_cf41(), (_1160_dt__update__tmp_h0).Dtor_cf42())
						}(_pat_let28_0)
					}(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(_pat_let_tv37), _pat_let_tv38))
				}(_pat_let27_0)
			}(Companion_D13_.Create_DC28_(_dafny.UnicodeSeqOfUtf8Bytes("bsb"), (_this).F11(), (_this).F17()))
			(globalState).F6 = !(!(((_this).F17()) || (true)))
		} else {
			(globalState).F1 = Companion_Default___.SafeDivisionInt((_1117_v10).F22(), p0)
			var _1162_v49 _dafny.CodePoint
			_ = _1162_v49
			_1162_v49 = _dafny.CodePoint('u')
			var _1163_v50 _dafny.Sequence
			_ = _1163_v50
			_1163_v50 = _dafny.SeqOf((_this).F17(), (_this).F17())
			var _1164_v51 _dafny.Sequence
			_ = _1164_v51
			_1164_v51 = _dafny.SeqOf(_1116_v9, _1116_v9, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(730))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg83 _dafny.Int) interface{} {
					return coer83(arg83)
				}
			}((func(_1165_v49 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1166_i3 _dafny.Int) _dafny.CodePoint {
					return _1165_v49
				}
			})(_1162_v49))))
			var _1167_v52 _dafny.Map
			_ = _1167_v52
			_1167_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _1117_v10.F23)
			var _1168_v53 _dafny.Sequence
			_ = _1168_v53
			_1168_v53 = _dafny.SeqOf(_1117_v10.F23, (_1117_v10).F22())
			_1162_v49 = Companion_Default___.Fm20(_dafny.Companion_Sequence_.Concatenate(_1163_v50, _1163_v50), Companion_Default___.SafeModuloInt((_1117_v10).F22(), _1117_v10.F23), _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.IntOfUint32(((_1164_v51).Select((Companion_Default___.SafeIndex((_1167_v52).Cardinality(), _dafny.IntOfUint32((_1164_v51).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())), _1168_v53), globalState)
			var _1169_v54 _dafny.Array
			_ = _1169_v54
			var _len0_20 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_20
			var _nw179 _dafny.Array
			_ = _nw179
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw179 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) _dafny.MultiSet = func(_1170_i4 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F11()))
				}
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw179 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw179).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw179).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_1169_v54 = _nw179
			_1169_v54 = _1169_v54
			var _1171_v55 _dafny.Map
			_ = _1171_v55
			_1171_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F17())
			var _1172_v56 _dafny.Sequence
			_ = _1172_v56
			_1172_v56 = _dafny.SeqOf(_1171_v55)
			var _1173_v57 *C3
			_ = _1173_v57
			var _nw180 *C3 = New_C3_()
			_ = _nw180
			_nw180.Ctor__(((_1172_v56).Select((Companion_Default___.SafeIndex((_1117_v10).F22(), _dafny.IntOfUint32((_1172_v56).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), (_this).F11(), _1117_v10.F23, (_this).F12())
			_1173_v57 = _nw180
			var _rhs217 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _rhs217
			var _rhs218 bool = !(((_this).F17()) && (!((_this).F17())))
			_ = _rhs218
			var _lhs172 *C3 = _1117_v10
			_ = _lhs172
			var _lhs173 *GlobalState = globalState
			_ = _lhs173
			_lhs172.F23 = _rhs217
			_lhs173.F7 = _rhs218
		}
		var _1174_v58 _dafny.Sequence
		_ = _1174_v58
		_1174_v58 = _dafny.SeqOf((_this).F17(), (_this).F17())
		var _1175_v59 _dafny.Sequence
		_ = _1175_v59
		_1175_v59 = _dafny.SeqOf(_dafny.IntOfUint32((_1174_v58).Cardinality()), _1117_v10.F23, (_this).F11())
		var _1176_v60 _dafny.MultiSet
		_ = _1176_v60
		_1176_v60 = _dafny.MultiSetOf((_this).F17())
		var _1177_v61 _dafny.MultiSet
		_ = _1177_v61
		_1177_v61 = _1176_v60
		var _1178_v62 _dafny.Array
		_ = _1178_v62
		var _nwElement0_51 _dafny.Int = (_this).F11()
		_ = _nwElement0_51
		var _nw181 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(16))
		_ = _nw181
		(_nw181).ArraySet1(_nwElement0_51, 0)
		(_nw181).ArraySet1((_1175_v59).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), _dafny.IntOfUint32((_1175_v59).Cardinality()))).Uint32()).(_dafny.Int), 1)
		(_nw181).ArraySet1(p0, 2)
		(_nw181).ArraySet1(((_1117_v10).F22()).Minus((_this).F11()), 3)
		(_nw181).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_1174_v58).Cardinality())), 4)
		(_nw181).ArraySet1((_this).F11(), 5)
		(_nw181).ArraySet1((_this).F11(), 6)
		(_nw181).ArraySet1(((_this).F11()).Plus(p0), 7)
		(_nw181).ArraySet1((_1177_v61).Cardinality(), 8)
		(_nw181).ArraySet1((_this).F11(), 9)
		(_nw181).ArraySet1(_1117_v10.F23, 10)
		(_nw181).ArraySet1((Companion_Default___.Fm11(globalState)).Cardinality(), 11)
		(_nw181).ArraySet1((_1117_v10).F22(), 12)
		(_nw181).ArraySet1((_1117_v10).Fm3(p0, (_this).F17(), (_this).F17(), globalState), 13)
		(_nw181).ArraySet1(p0, 14)
		(_nw181).ArraySet1((_this).F11(), 15)
		_1178_v62 = _nw181
		var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_1178_v62), 0))
		_ = _index143
		(_1178_v62).ArraySet1(_dafny.IntOfInt64(255), (_index143).Int())
		var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_1178_v62), 0))
		_ = _index144
		var _rhs219 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
			if (_1176_v60).Contains((_this).F17()) {
				return (_1176_v60).Multiplicity((_this).F17())
			}
			return _1117_v10.F23
		})())
		_ = _rhs219
		var _rhs220 _dafny.Int = (_1117_v10).F22()
		_ = _rhs220
		var _rhs221 _dafny.Int = _dafny.IntOfInt64(297)
		_ = _rhs221
		var _lhs174 _dafny.Array = _1178_v62
		_ = _lhs174
		var _lhs175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_1178_v62), 0))
		_ = _lhs175
		var _lhs176 *C3 = _1117_v10
		_ = _lhs176
		var _lhs177 *C3 = _1117_v10
		_ = _lhs177
		(_lhs174).ArraySet1(_rhs219, (_lhs175).Int())
		_lhs176.F23 = _rhs220
		_lhs177.F23 = _rhs221
		var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_1178_v62), 0))
		_ = _index145
		(_1178_v62).ArraySet1(_dafny.IntOfInt64(600), (_index145).Int())
		(_1117_v10).F23 = (_1117_v10).F22()
		var _1179_v63 _dafny.Array
		_ = _1179_v63
		var _nw182 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(18))
		_ = _nw182
		_1179_v63 = _nw182
		r0 = _1179_v63
		var _1180_v64 _dafny.Map
		_ = _1180_v64
		_1180_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(825))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg84 _dafny.Int) interface{} {
				return coer84(arg84)
			}
		}(func(_1181_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})), (_this).F16())
		r1 = _1180_v64
		return r0, r1
	}
}
func (_this *C4) M5(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) (_dafny.Int, bool, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		if p1 {
			var _1182_v0 _dafny.Array
			_ = _1182_v0
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_21
			var _nw183 _dafny.Array
			_ = _nw183
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw183 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) bool = (func(_1183_p1 bool) func(_dafny.Int) bool {
					return func(_1184_i0 _dafny.Int) bool {
						return _1183_p1
					}
				})(p1)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw183 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw183).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw183).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_1182_v0 = _nw183
			var _1185_v1 _dafny.Set
			_ = _1185_v1
			_1185_v1 = _dafny.SetOf(Companion_Default___.Fm46(globalState))
			var _1186_v2 _dafny.Map
			_ = _1186_v2
			_1186_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_1182_v0), _1185_v1)
			var _1187_v3 _dafny.MultiSet
			_ = _1187_v3
			_1187_v3 = _dafny.MultiSetOf(_1182_v0, _1182_v0, _1182_v0)
			_1186_v2 = (_1186_v2).Update(_1187_v3, (_1185_v1).Union(_1185_v1))
			var _1188_v5 _dafny.Map
			_ = _1188_v5
			_1188_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _1189_v6 D8
			_ = _1189_v6
			_1189_v6 = Companion_D8_.Create_DC15_(_1188_v5)
			var _1190_v7 D8
			_ = _1190_v7
			_1190_v7 = Companion_D8_.Create_DC17_(_1189_v6)
			var _1191_v8 D13
			_ = _1191_v8
			_1191_v8 = Companion_D13_.Create_DC27_(func() _dafny.Map {
				var _coll60 = _dafny.NewMapBuilder()
				_ = _coll60
				for _iter62 := _dafny.Iterate((_dafny.SeqOf(_1190_v7)).Elements()); ; {
					_compr_60, _ok62 := _iter62()
					if !_ok62 {
						break
					}
					var _1192_v4 D8
					_1192_v4 = interface{}(_compr_60).(D8)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_1190_v7), _1192_v4) {
						_coll60.Add(_1192_v4, (_this).F11())
					}
				}
				return _coll60.ToMap()
			}())
			var _1193_v9 D13
			_ = _1193_v9
			_1193_v9 = Companion_D13_.Create_DC29_(_1191_v8)
			var _1194_v10 D13
			_ = _1194_v10
			_1194_v10 = Companion_D13_.Create_DC29_(_1191_v8)
			var _1195_v11 _dafny.Map
			_ = _1195_v11
			_1195_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1190_v7, _dafny.IntOfInt64(992))
			_1194_v10 = Companion_D13_.Create_DC29_(Companion_D13_.Create_DC27_(_1195_v11))
			(globalState).F7 = p1
			var _1196_v12 _dafny.Sequence
			_ = _1196_v12
			_1196_v12 = _dafny.SeqOf(p1, p0, p1)
			var _1197_v13 _dafny.Sequence
			_ = _1197_v13
			_1197_v13 = _dafny.SeqOf(!(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(p1), _1196_v12)), p0, (func() bool {
				if p0 {
					return !(p0)
				}
				return p0
			})(), !(((_this).F16()).IsSubsetOf((_this).F16())), p0)
			var _1198_v14 _dafny.Map
			_ = _1198_v14
			_1198_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1188_v5)
			var _1199_v15 _dafny.Sequence
			_ = _1199_v15
			_1199_v15 = _dafny.UnicodeSeqOfUtf8Bytes("aw")
			_1196_v12 = Companion_Default___.Fm12(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F11()), (_1198_v14).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-390))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg85 _dafny.Int) interface{} {
					return coer85(arg85)
				}
			}(func(_1200_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			})), _1199_v15), globalState)
			var _1201_v16 *C3
			_ = _1201_v16
			var _nw184 *C3 = New_C3_()
			_ = _nw184
			_nw184.Ctor__((_this).F11(), (_dafny.Zero).Minus(((_this).F11()).Minus(_dafny.IntOfInt64(267))), p2, (_this).F12())
			_1201_v16 = _nw184
		} else {
			var _1202_v17 *C3
			_ = _1202_v17
			var _nw185 *C3 = New_C3_()
			_ = _nw185
			_nw185.Ctor__(p2, _dafny.IntOfInt64(-116), (func() _dafny.Int {
				if ((_this).F16()).Contains(p2) {
					return ((_this).F16()).Multiplicity(p2)
				}
				return p2
			})(), (_this).F12())
			_1202_v17 = _nw185
			var _1203_v18 _dafny.Array
			_ = _1203_v18
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_22
			var _nw186 _dafny.Array
			_ = _nw186
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw186 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Map = (func(_1204_p1 bool, _1205_p0 bool) func(_dafny.Int) _dafny.Map {
					return func(_1206_i2 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D11_.Create_DC24_(_dafny.IntOfInt64(2), _dafny.SeqOf(_1204_p1, _1205_p0), _dafny.CodePoint('n')), _1204_p1)
					}
				})(p1, p0)
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw186 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw186).ArraySet1(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw186).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_1203_v18 = _nw186
			var _1207_v19 _dafny.CodePoint
			_ = _1207_v19
			_1207_v19 = _dafny.CodePoint('f')
			var _1208_v20 D11
			_ = _1208_v20
			_1208_v20 = Companion_D11_.Create_DC24_(p2, _dafny.SeqOf(p1, p1), _1207_v19)
			var _1209_v21 _dafny.Map
			_ = _1209_v21
			_1209_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1208_v20, p1)
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1203_v18), 0))
			_ = _index146
			(_1203_v18).ArraySet1(_1209_v21, (_index146).Int())
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1203_v18), 0))
			_ = _index147
			(_1203_v18).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1208_v20, p1)).Merge((_1209_v21).Merge(_1209_v21)), (_index147).Int())
			var _1210_v22 D0
			_ = _1210_v22
			_1210_v22 = Companion_D0_.Create_DC1_((_this).F17(), _1202_v17.F23, !(p1))
			var _1211_v23 _dafny.Sequence
			_ = _1211_v23
			_1211_v23 = _dafny.SeqOf((_1202_v17).F22())
			var _1212_v24 D7
			_ = _1212_v24
			_1212_v24 = Companion_D7_.Create_DC12_(_1211_v23)
			var _1213_v25 _dafny.Map
			_ = _1213_v25
			_1213_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm39((_1202_v17).F22(), !(p1), (_1210_v22).Dtor_cf2(), globalState), _1212_v24)
			_1213_v25 = (_1213_v25).Update(_1207_v19, _1212_v24)
			(_1202_v17).F23 = (_dafny.Zero).Minus((_1202_v17).F22())
			var _1214_v26 *C2
			_ = _1214_v26
			var _nw187 *C2 = New_C2_()
			_ = _nw187
			_nw187.Ctor__(_1207_v19, _dafny.IntOfInt64(-968))
			_1214_v26 = _nw187
		}
		var _1215_v27 _dafny.Sequence
		_ = _1215_v27
		_1215_v27 = _dafny.SeqOf(p1)
		var _1216_v28 D6
		_ = _1216_v28
		_1216_v28 = Companion_D6_.Create_DC10_(_1215_v27)
		var _pat_let_tv39 = p2
		_ = _pat_let_tv39
		var _pat_let_tv40 = p2
		_ = _pat_let_tv40
		var _pat_let_tv41 = p2
		_ = _pat_let_tv41
		r0 = (_dafny.Zero).Minus(func(_source17 D6) _dafny.Int {
			if _source17.Is_DC10() {
				var _1217___mcc_h0 _dafny.Sequence = _source17.Get_().(D6_DC10).Cf13
				_ = _1217___mcc_h0
				var _1218_cf13 _dafny.Sequence = _1217___mcc_h0
				_ = _1218_cf13
				return _pat_let_tv39
			} else if _source17.Is_DC11() {
				var _1219___mcc_h1 _dafny.Int = _source17.Get_().(D6_DC11).Cf14
				_ = _1219___mcc_h1
				var _1220___mcc_h2 _dafny.Int = _source17.Get_().(D6_DC11).Cf15
				_ = _1220___mcc_h2
				var _1221___mcc_h3 _dafny.Int = _source17.Get_().(D6_DC11).Cf16
				_ = _1221___mcc_h3
				var _1222_cf16 _dafny.Int = _1221___mcc_h3
				_ = _1222_cf16
				var _1223_cf15 _dafny.Int = _1220___mcc_h2
				_ = _1223_cf15
				var _1224_cf14 _dafny.Int = _1219___mcc_h1
				_ = _1224_cf14
				return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(437), (_this).F11())
			} else {
				var _1225___mcc_h4 _dafny.Map = _source17.Get_().(D6_DC9).Cf12
				_ = _1225___mcc_h4
				var _1226_cf12 _dafny.Map = _1225___mcc_h4
				_ = _1226_cf12
				return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _pat_let_tv40)).Cardinality()).Times(_pat_let_tv41)
			}
		}(_1216_v28))
		r0 = p2
		var _1227_v29 _dafny.Array
		_ = _1227_v29
		var _len0_23 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_23
		var _nw188 _dafny.Array
		_ = _nw188
		if _len0_23.Cmp(_dafny.Zero) == 0 {
			_nw188 = _dafny.NewArray(_len0_23)
		} else {
			var _init23 func(_dafny.Int) bool = func(_1228_i3 _dafny.Int) bool {
				return (_this).F17()
			}
			_ = _init23
			var _element0_23 = _init23(_dafny.Zero)
			_ = _element0_23
			_nw188 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
			(_nw188).ArraySet1(_element0_23, 0)
			var _nativeLen0_23 = (_len0_23).Int()
			_ = _nativeLen0_23
			for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
				(_nw188).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
			}
		}
		_1227_v29 = _nw188
		var _nw189 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
		_ = _nw189
		_1227_v29 = _nw189
		var _1229_v30 _dafny.Array
		_ = _1229_v30
		var _nw190 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
		_ = _nw190
		_1229_v30 = _nw190
		var _1230_v31 _dafny.Array
		_ = _1230_v31
		var _len0_24 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_24
		var _nw191 _dafny.Array
		_ = _nw191
		if _len0_24.Cmp(_dafny.Zero) == 0 {
			_nw191 = _dafny.NewArray(_len0_24)
		} else {
			var _init24 func(_dafny.Int) D2 = func(_1231_i4 _dafny.Int) D2 {
				return Companion_D2_.Create_DC5_()
			}
			_ = _init24
			var _element0_24 = _init24(_dafny.Zero)
			_ = _element0_24
			_nw191 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
			(_nw191).ArraySet1(_element0_24, 0)
			var _nativeLen0_24 = (_len0_24).Int()
			_ = _nativeLen0_24
			for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
				(_nw191).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
			}
		}
		_1230_v31 = _nw191
		var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_1229_v30), 0))
		_ = _index148
		(_1229_v30).ArraySet1(_1230_v31, (_index148).Int())
		var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_1229_v30), 0))
		_ = _index149
		var _rhs222 _dafny.Array = _1230_v31
		_ = _rhs222
		var _rhs223 bool = p1
		_ = _rhs223
		var _rhs224 bool = (_this).F17()
		_ = _rhs224
		var _rhs225 _dafny.Int = p2
		_ = _rhs225
		var _rhs226 _dafny.Int = _dafny.IntOfInt64(93)
		_ = _rhs226
		var _lhs178 _dafny.Array = _1229_v30
		_ = _lhs178
		var _lhs179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_1229_v30), 0))
		_ = _lhs179
		var _lhs180 *GlobalState = globalState
		_ = _lhs180
		var _lhs181 *GlobalState = globalState
		_ = _lhs181
		(_lhs178).ArraySet1(_rhs222, (_lhs179).Int())
		r1 = _rhs223
		_lhs180.F6 = _rhs224
		r2 = _rhs225
		_lhs181.F1 = _rhs226
		var _1232_i5 _dafny.Int
		_ = _1232_i5
		_1232_i5 = _dafny.Zero
		{
			for (_this).F17() {
				{
					if (_1232_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L13
					}
					_1232_i5 = (_1232_i5).Plus(_dafny.One)
					(globalState).F1 = (_this).F11()
					var _1233_v32 *C0
					_ = _1233_v32
					var _nw192 *C0 = New_C0_()
					_ = _nw192
					_nw192.Ctor__((_1216_v28).Dtor_cf13(), (_this).F11())
					_1233_v32 = _nw192
					var _1234_v33 _dafny.Map
					_ = _1234_v33
					_1234_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1233_v32)
					var _1235_v34 _dafny.Array
					_ = _1235_v34
					var _nw193 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
					_ = _nw193
					_1235_v34 = _nw193
					var _1236_v35 _dafny.Map
					_ = _1236_v35
					_1236_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1234_v33).Cardinality(), _1235_v34)
					_1236_v35 = (_1236_v35).Update(p2, _1235_v34)
					var _1237_v36 _dafny.CodePoint
					_ = _1237_v36
					_1237_v36 = _dafny.CodePoint('b')
					_1237_v36 = _1237_v36
					(globalState).F1 = ((_dafny.Zero).Minus(((_dafny.SetOf(_dafny.IntOfInt64(593), p2)).Cardinality()).Minus(p2))).Plus(_dafny.IntOfUint32(((_1233_v32).F20()).Cardinality()))
					goto C13
				}
			C13:
			}
			goto L13
		}
	L13:
		var _1238_v37 _dafny.MultiSet
		_ = _1238_v37
		_1238_v37 = _dafny.MultiSetOf(_dafny.SeqOf((_this).F17()), _dafny.SeqOf(p1, (_this).F17()), _1215_v27, _1215_v27)
		r0 = (_dafny.Zero).Minus(((_1238_v37).Update(_dafny.SeqOf((_this).F17(), p0, true), Companion_Default___.Abs((_dafny.Zero).Minus((_this).F11())))).Cardinality())
		r1 = !(((_this).F12()).Difference(Companion_Default___.Fm41(globalState))).Contains(p0)
		var _1239_v38 _dafny.MultiSet
		_ = _1239_v38
		_1239_v38 = _dafny.MultiSetOf(p0, true, false, p1)
		var _1240_v39 _dafny.MultiSet
		_ = _1240_v39
		_1240_v39 = _dafny.MultiSetOf(_1239_v38)
		r2 = (func() _dafny.Int {
			if p1 {
				return (func() _dafny.Int {
					if (_1240_v39).Contains(((_1239_v38).Update(true, Companion_Default___.Abs(p2))).Update(false, Companion_Default___.Abs(p2))) {
						return (_1240_v39).Multiplicity(((_1239_v38).Update(true, Companion_Default___.Abs(p2))).Update(false, Companion_Default___.Abs(p2)))
					}
					return p2
				})()
			}
			return _dafny.IntOfInt64(14)
		})()
		return r0, r1, r2
	}
}
func (_this *C4) F16() _dafny.MultiSet {
	{
		return _this._f16
	}
}
func (_this *C4) F17() bool {
	{
		return _this._f17
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	dummy byte
}

func New_C5_() *C5 {
	_this := C5{}

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

func (_this *C5) Ctor__() {
	{
	}
}
func (_this *C5) Fm9(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(((_dafny.IntOfInt64(539)).Plus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("usdxyfb")).Cardinality()))))).Minus(_dafny.IntOfInt64(933)))
	}
}
func (_this *C5) M4(p0 _dafny.Sequence, p1 _dafny.Array, globalState *GlobalState) {
	{
		var _1241_v0 _dafny.Int
		_ = _1241_v0
		_1241_v0 = _dafny.IntOfInt64(826)
		(globalState).F1 = _1241_v0
		var _1242_v1 bool
		_ = _1242_v1
		_1242_v1 = true
		var _1243_v2 _dafny.CodePoint
		_ = _1243_v2
		_1243_v2 = _dafny.CodePoint('k')
		var _1244_v3 _dafny.Map
		_ = _1244_v3
		_1244_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v1, (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1241_v0, _1243_v2)).Cardinality(), _1241_v0)).IsProperSubsetOf(_dafny.MultiSetOf((_dafny.Zero).Minus(_1241_v0))))
		if (func() bool {
			if (_1244_v3).Contains(_1242_v1) {
				return (_1244_v3).Get(_1242_v1).(bool)
			}
			return _1242_v1
		})() {
			var _1245_v4 _dafny.Map
			_ = _1245_v4
			_1245_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1241_v0, _dafny.IntOfInt64(821))
			_1245_v4 = (_1245_v4).Update(_1241_v0, (_dafny.Zero).Minus(_1241_v0))
			var _1246_v5 _dafny.Sequence
			_ = _1246_v5
			_1246_v5 = _dafny.SeqOf(_1242_v1)
			(globalState).F6 = (_1242_v1) && ((_1246_v5).Select((Companion_Default___.SafeIndex(_1241_v0, _dafny.IntOfUint32((_1246_v5).Cardinality()))).Uint32()).(bool))
			var _1247_v6 _dafny.Map
			_ = _1247_v6
			_1247_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1243_v2, _1241_v0)
			var _1248_v7 *C0
			_ = _1248_v7
			var _nw194 *C0 = New_C0_()
			_ = _nw194
			_nw194.Ctor__(_1246_v5, (func() _dafny.Int {
				if (_1247_v6).Contains(_1243_v2) {
					return (_1247_v6).Get(_1243_v2).(_dafny.Int)
				}
				return _1241_v0
			})())
			_1248_v7 = _nw194
			var _1249_v8 _dafny.Array
			_ = _1249_v8
			var _nw195 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
			_ = _nw195
			_1249_v8 = _nw195
			var _1250_v9 _dafny.Set
			_ = _1250_v9
			_1250_v9 = _dafny.SetOf(false)
			var _1251_v10 T0
			_ = _1251_v10
			var _nw196 *C3 = New_C3_()
			_ = _nw196
			_nw196.Ctor__(_dafny.IntOfInt64(611), (_dafny.Zero).Minus(_1241_v0), _1241_v0, _1250_v9)
			_1251_v10 = _nw196
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1249_v8), 0))
			_ = _index150
			(_1249_v8).ArraySet1(_1251_v10, (_index150).Int())
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1249_v8), 0))
			_ = _index151
			(_1249_v8).ArraySet1(_1251_v10, (_index151).Int())
			(globalState).F1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(p0, p0), p0)).Cardinality())
		} else {
			(globalState).F1 = _1241_v0
			var _1252_v11 _dafny.Sequence
			_ = _1252_v11
			_1252_v11 = _dafny.SeqOf(Companion_Default___.SafeModuloInt(_1241_v0, _1241_v0), _1241_v0)
			_1252_v11 = _dafny.SeqOf(_1241_v0)
			var _1253_v12 *C3
			_ = _1253_v12
			var _nw197 *C3 = New_C3_()
			_ = _nw197
			_nw197.Ctor__(_1241_v0, _1241_v0, _1241_v0, _dafny.SetOf(_1242_v1))
			_1253_v12 = _nw197
			var _1254_v13 _dafny.Array
			_ = _1254_v13
			var _nwElement0_52 *C3 = _1253_v12
			_ = _nwElement0_52
			var _nw198 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(17))
			_ = _nw198
			(_nw198).ArraySet1(_nwElement0_52, 0)
			(_nw198).ArraySet1(_1253_v12, 1)
			(_nw198).ArraySet1(_1253_v12, 2)
			(_nw198).ArraySet1((func() *C3 {
				if _1242_v1 {
					return _1253_v12
				}
				return _1253_v12
			})(), 3)
			(_nw198).ArraySet1(_1253_v12, 4)
			(_nw198).ArraySet1(_1253_v12, 5)
			(_nw198).ArraySet1(_1253_v12, 6)
			(_nw198).ArraySet1(_1253_v12, 7)
			(_nw198).ArraySet1(_1253_v12, 8)
			(_nw198).ArraySet1(_1253_v12, 9)
			(_nw198).ArraySet1(_1253_v12, 10)
			(_nw198).ArraySet1(_1253_v12, 11)
			(_nw198).ArraySet1((func() *C3 {
				if !(!(!(true))) {
					return _1253_v12
				}
				return _1253_v12
			})(), 12)
			(_nw198).ArraySet1(_1253_v12, 13)
			(_nw198).ArraySet1((func() *C3 {
				if _1242_v1 {
					return _1253_v12
				}
				return _1253_v12
			})(), 14)
			(_nw198).ArraySet1(_1253_v12, 15)
			(_nw198).ArraySet1(_1253_v12, 16)
			_1254_v13 = _nw198
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1254_v13), 0))
			_ = _index152
			(_1254_v13).ArraySet1(_1253_v12, (_index152).Int())
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1254_v13), 0))
			_ = _index153
			(_1254_v13).ArraySet1((func() *C3 {
				if Companion_Default___.Fm0(globalState) {
					return _1253_v12
				}
				return _1253_v12
			})(), (_index153).Int())
			var _1255_v14 D0
			_ = _1255_v14
			_1255_v14 = Companion_D0_.Create_DC0_(_1253_v12.F23)
			var _pat_let_tv42 = _1241_v0
			_ = _pat_let_tv42
			_1255_v14 = func(_pat_let29_0 D0) D0 {
				return func(_1256_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let30_0 _dafny.Int) D0 {
						return func(_1257_dt__update_hcf0_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC0_(_1257_dt__update_hcf0_h0)
						}(_pat_let30_0)
					}(_pat_let_tv42)
				}(_pat_let29_0)
			}(_1255_v14)
			var _1258_v15 _dafny.Sequence
			_ = _1258_v15
			_1258_v15 = _dafny.SeqOf(_1242_v1, _1242_v1)
			var _1259_v16 _dafny.Sequence
			_ = _1259_v16
			_1259_v16 = _dafny.SeqOf(_1258_v15)
			(globalState).F1 = (_dafny.IntOfUint32((_1259_v16).Cardinality())).Times((_dafny.Zero).Minus(_1241_v0))
		}
		var _1260_v17 D13
		_ = _1260_v17
		_1260_v17 = Companion_D13_.Create_DC28_(p0, _1241_v0, _1242_v1)
		if (_1260_v17).Dtor_cf42() {
			var _1261_v18 _dafny.Array
			_ = _1261_v18
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_25
			var _nw199 _dafny.Array
			_ = _nw199
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw199 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Map = (func(_1262_v3 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_1263_i0 _dafny.Int) _dafny.Map {
						return (_1262_v3).Merge(_1262_v3)
					}
				})(_1244_v3)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw199 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw199).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw199).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_1261_v18 = _nw199
			_1261_v18 = _1261_v18
			(globalState).F1 = _dafny.IntOfInt64(498)
			var _1264_v19 _dafny.MultiSet
			_ = _1264_v19
			_1264_v19 = _dafny.MultiSetOf(Companion_Default___.Fm0(globalState), _1242_v1)
			(globalState).F7 = !((_1264_v19).IsDisjointFrom(_1264_v19)) || (_1242_v1)
			(globalState).F6 = _1242_v1
			var _1265_v20 _dafny.Array
			_ = _1265_v20
			var _nw200 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw200
			_1265_v20 = _nw200
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1265_v20), 0))
			_ = _index154
			(_1265_v20).ArraySet1(_1241_v0, (_index154).Int())
			var _1266_v21 _dafny.Sequence
			_ = _1266_v21
			_1266_v21 = _dafny.SeqOf(_1242_v1)
			var _1267_v22 _dafny.MultiSet
			_ = _1267_v22
			_1267_v22 = _dafny.MultiSetOf(_1266_v21, _1266_v21, _1266_v21, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1266_v21, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_1266_v21).Cardinality()))).Uint32(), _1242_v1), _dafny.Companion_Sequence_.Update(_1266_v21, (Companion_Default___.SafeIndex(_1241_v0, _dafny.IntOfUint32((_1266_v21).Cardinality()))).Uint32(), _1242_v1)))
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1265_v20), 0))
			_ = _index155
			(_1265_v20).ArraySet1((func() _dafny.Int {
				if (_1267_v22).Contains((func() _dafny.Sequence {
					if !(_1242_v1) {
						return _1266_v21
					}
					return _dafny.SeqOf(_1242_v1)
				})()) {
					return (_1267_v22).Multiplicity((func() _dafny.Sequence {
						if !(_1242_v1) {
							return _1266_v21
						}
						return _dafny.SeqOf(_1242_v1)
					})())
				}
				return _1241_v0
			})(), (_index155).Int())
		} else {
			var _1268_v23 _dafny.Array
			_ = _1268_v23
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_26
			var _nw201 _dafny.Array
			_ = _nw201
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw201 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) _dafny.Int = (func(_1269_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1270_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1270_i1, _1269_v0)
					}
				})(_1241_v0)
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw201 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw201).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw201).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_1268_v23 = _nw201
			var _1271_v24 _dafny.Map
			_ = _1271_v24
			_1271_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1268_v23, true)
			var _1272_v25 D10
			_ = _1272_v25
			_1272_v25 = Companion_D10_.Create_DC20_(_1271_v24)
			var _source18 D10 = _1272_v25
			_ = _source18
			if _source18.Is_DC21() {
				var _1273___mcc_h0 bool = _source18.Get_().(D10_DC21).Cf28
				_ = _1273___mcc_h0
				var _1274___mcc_h1 bool = _source18.Get_().(D10_DC21).Cf29
				_ = _1274___mcc_h1
				var _1275_cf29 bool = _1274___mcc_h1
				_ = _1275_cf29
				var _1276_cf28 bool = _1273___mcc_h0
				_ = _1276_cf28
				var _1277_v26 _dafny.MultiSet
				_ = _1277_v26
				_1277_v26 = _dafny.MultiSetOf(Companion_Default___.Fm2(_1241_v0, _dafny.IntOfUint32((p0).Cardinality()), globalState))
				_1277_v26 = _1277_v26
				var _1278_v27 _dafny.Sequence
				_ = _1278_v27
				_1278_v27 = _dafny.SeqOf(_1242_v1)
				var _1279_v28 D6
				_ = _1279_v28
				_1279_v28 = Companion_D6_.Create_DC10_(_1278_v27)
				_1279_v28 = (func() D6 {
					if (func() bool {
						if (_1244_v3).Contains(_1275_cf29) {
							return (_1244_v3).Get(_1275_cf29).(bool)
						}
						return Companion_Default___.Fm0(globalState)
					})() {
						return _1279_v28
					}
					return _1279_v28
				})()
				var _1280_v29 _dafny.Array
				_ = _1280_v29
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_27
				var _nw202 _dafny.Array
				_ = _nw202
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw202 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) _dafny.Map = (func(_1281_cf28 bool, _1282_p0 _dafny.Sequence, _1283_v1 bool) func(_dafny.Int) _dafny.Map {
						return func(_1284_i2 _dafny.Int) _dafny.Map {
							return ((Companion_D14_.Create_DC30_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1281_cf28, _1282_p0))).Dtor_cf44()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1283_v1, _1282_p0))
						}
					})(_1276_cf28, p0, _1242_v1)
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw202 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw202).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw202).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_1280_v29 = _nw202
				var _1285_v30 _dafny.Map
				_ = _1285_v30
				_1285_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v1, p0)
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_1280_v29), 0))
				_ = _index156
				(_1280_v29).ArraySet1(_1285_v30, (_index156).Int())
				var _1286_v31 _dafny.Sequence
				_ = _1286_v31
				_1286_v31 = _dafny.SeqOf(_1285_v30)
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_1280_v29), 0))
				_ = _index157
				(_1280_v29).ArraySet1((_1286_v31).Select((Companion_Default___.SafeIndex(_1241_v0, _dafny.IntOfUint32((_1286_v31).Cardinality()))).Uint32()).(_dafny.Map), (_index157).Int())
				var _1287_v32 _dafny.Set
				_ = _1287_v32
				_1287_v32 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, p0)).Cardinality()))
				var _1288_v33 _dafny.Map
				_ = _1288_v33
				_1288_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1241_v0, _1241_v0)
				var _1289_v34 _dafny.Map
				_ = _1289_v34
				_1289_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1243_v2, Companion_Default___.Fm0(globalState))
				var _1290_v36 _dafny.Map
				_ = _1290_v36
				_1290_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1243_v2, _1241_v0)
				var _1291_v37 _dafny.Array
				_ = _1291_v37
				var _nwElement0_53 bool = !(_1276_cf28) || (_1276_cf28)
				_ = _nwElement0_53
				var _nw203 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(29))
				_ = _nw203
				(_nw203).ArraySet1(_nwElement0_53, 0)
				(_nw203).ArraySet1(_1242_v1, 1)
				(_nw203).ArraySet1(((_this).Fm9((_dafny.Zero).Minus((_1288_v33).Cardinality()), _dafny.IntOfInt64(-717), _1242_v1, globalState)).Cmp(_1241_v0) < 0, 2)
				(_nw203).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(func() _dafny.Map {
					var _coll61 = _dafny.NewMapBuilder()
					_ = _coll61
					for _iter63 := _dafny.Iterate((_1290_v36).Keys().Elements()); ; {
						_compr_61, _ok63 := _iter63()
						if !_ok63 {
							break
						}
						var _1292_v35 _dafny.CodePoint
						_1292_v35 = interface{}(_compr_61).(_dafny.CodePoint)
						if (_1290_v36).Contains(_1292_v35) {
							_coll61.Add(_1292_v35, _1275_cf29)
						}
					}
					return _coll61.ToMap()
				}(), _1289_v34), _1289_v34), 3)
				(_nw203).ArraySet1(!((_1278_v27).Select((Companion_Default___.SafeIndex(_1241_v0, _dafny.IntOfUint32((_1278_v27).Cardinality()))).Uint32()).(bool)), 4)
				(_nw203).ArraySet1(_1276_cf28, 5)
				(_nw203).ArraySet1((func() bool {
					if !(_1275_cf29) {
						return _1275_cf29
					}
					return (_1278_v27).Select((Companion_Default___.SafeIndex(_1241_v0, _dafny.IntOfUint32((_1278_v27).Cardinality()))).Uint32()).(bool)
				})(), 6)
				(_nw203).ArraySet1(true, 7)
				(_nw203).ArraySet1(false, 8)
				(_nw203).ArraySet1((func() bool {
					if _1242_v1 {
						return _1276_cf28
					}
					return _1242_v1
				})(), 9)
				(_nw203).ArraySet1(!(_1276_cf28), 10)
				(_nw203).ArraySet1(!(_1275_cf29) || (Companion_Default___.Fm0(globalState)), 11)
				(_nw203).ArraySet1(!(_dafny.Companion_Sequence_.IsProperPrefixOf(p0, p0)), 12)
				(_nw203).ArraySet1(_1242_v1, 13)
				(_nw203).ArraySet1(false, 14)
				(_nw203).ArraySet1(_1242_v1, 15)
				(_nw203).ArraySet1((_1242_v1) && (_1276_cf28), 16)
				(_nw203).ArraySet1(_1242_v1, 17)
				(_nw203).ArraySet1(_1276_cf28, 18)
				(_nw203).ArraySet1((func() bool {
					if _1242_v1 {
						return _1275_cf29
					}
					return _1242_v1
				})(), 19)
				(_nw203).ArraySet1(true, 20)
				(_nw203).ArraySet1((func() bool {
					if (_1244_v3).Contains(_1242_v1) {
						return (_1244_v3).Get(_1242_v1).(bool)
					}
					return _1242_v1
				})(), 21)
				(_nw203).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(p0, _dafny.UnicodeSeqOfUtf8Bytes("ley")), 22)
				(_nw203).ArraySet1(_1276_cf28, 23)
				(_nw203).ArraySet1(_1276_cf28, 24)
				(_nw203).ArraySet1(!(_1242_v1), 25)
				(_nw203).ArraySet1((_1242_v1) && (_1276_cf28), 26)
				(_nw203).ArraySet1((func() bool {
					if _1276_cf28 {
						return _1276_cf28
					}
					return true
				})(), 27)
				(_nw203).ArraySet1(_1275_cf29, 28)
				_1291_v37 = _nw203
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1291_v37), 0))
				_ = _index158
				(_1291_v37).ArraySet1(_1242_v1, (_index158).Int())
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1291_v37), 0))
				_ = _index159
				var _rhs227 _dafny.Set = (_1287_v32).Difference(_1287_v32)
				_ = _rhs227
				var _rhs228 bool = Companion_Default___.Fm0(globalState)
				_ = _rhs228
				var _lhs182 _dafny.Array = _1291_v37
				_ = _lhs182
				var _lhs183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1291_v37), 0))
				_ = _lhs183
				_1287_v32 = _rhs227
				(_lhs182).ArraySet1(_rhs228, (_lhs183).Int())
			} else {
				var _1293___mcc_h2 _dafny.Map = _source18.Get_().(D10_DC20).Cf27
				_ = _1293___mcc_h2
				var _1294_cf27 _dafny.Map = _1293___mcc_h2
				_ = _1294_cf27
				var _1295_v38 _dafny.Sequence
				_ = _1295_v38
				_1295_v38 = _dafny.SeqOf(_1242_v1)
				var _1296_v39 *C1
				_ = _1296_v39
				var _nw204 *C1 = New_C1_()
				_ = _nw204
				_nw204.Ctor__(_1242_v1, (_1295_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_1295_v38).Cardinality()))).Uint32()).(bool))
				_1296_v39 = _nw204
				(globalState).F7 = (_1296_v39).F19()
				var _1297_v40 D8
				_ = _1297_v40
				_1297_v40 = Companion_D8_.Create_DC16_(_1242_v1, (_this).Fm9(_1241_v0, _1241_v0, _1242_v1, globalState))
				var _1298_v41 D8
				_ = _1298_v41
				_1298_v41 = Companion_D8_.Create_DC17_(_1297_v40)
				var _1299_v42 D8
				_ = _1299_v42
				_1299_v42 = Companion_D8_.Create_DC17_(_1297_v40)
				var _1300_v43 _dafny.Map
				_ = _1300_v43
				_1300_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1299_v42, _1241_v0)
				var _1301_v44 D13
				_ = _1301_v44
				_1301_v44 = Companion_D13_.Create_DC27_(_1300_v43)
				_1301_v44 = _1301_v44
				(globalState).F7 = true
			}
			(globalState).F1 = _1241_v0
			var _1302_v45 _dafny.Sequence
			_ = _1302_v45
			_1302_v45 = _dafny.SeqOf((_dafny.Zero).Minus(_1241_v0))
			_1242_v1 = _dafny.Companion_Sequence_.IsProperPrefixOf(_1302_v45, _dafny.SeqOf(_1241_v0))
			var _1303_v46 _dafny.Sequence
			_ = _1303_v46
			_1303_v46 = _dafny.SeqOf(_1242_v1, !(_1242_v1), _1242_v1, _1242_v1)
			var _1304_v47 _dafny.Sequence
			_ = _1304_v47
			_1304_v47 = _1303_v46
			var _source19 _dafny.Sequence = _1303_v46
			_ = _source19
			var _1305___mcc_h3 _dafny.Sequence = _source19
			_ = _1305___mcc_h3
			var _1306_cf10 _dafny.Sequence = _1305___mcc_h3
			_ = _1306_cf10
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_1268_v23), 0))
			_ = _index160
			(_1268_v23).ArraySet1(_1241_v0, (_index160).Int())
			var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_1268_v23), 0))
			_ = _index161
			var _rhs229 _dafny.Int = _1241_v0
			_ = _rhs229
			var _rhs230 _dafny.Int = _1241_v0
			_ = _rhs230
			var _rhs231 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_1302_v45).Cardinality()))
			_ = _rhs231
			var _lhs184 *GlobalState = globalState
			_ = _lhs184
			var _lhs185 *GlobalState = globalState
			_ = _lhs185
			var _lhs186 _dafny.Array = _1268_v23
			_ = _lhs186
			var _lhs187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_1268_v23), 0))
			_ = _lhs187
			_lhs184.F1 = _rhs229
			_lhs185.F1 = _rhs230
			(_lhs186).ArraySet1(_rhs231, (_lhs187).Int())
			var _1307_v48 D8
			_ = _1307_v48
			_1307_v48 = Companion_D8_.Create_DC16_(_1242_v1, _1241_v0)
			_1307_v48 = _1307_v48
			var _1308_v49 *C1
			_ = _1308_v49
			var _nw205 *C1 = New_C1_()
			_ = _nw205
			_nw205.Ctor__(_1242_v1, _1242_v1)
			_1308_v49 = _nw205
			var _1309_v50 _dafny.Map
			_ = _1309_v50
			_1309_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1268_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_1268_v23), 0))).Int()).(_dafny.Int), _1241_v0)
			var _1310_v51 *C1
			_ = _1310_v51
			var _nw206 *C1 = New_C1_()
			_ = _nw206
			_nw206.Ctor__(_dafny.Companion_Sequence_.IsPrefixOf(p0, p0), !(_1309_v50).Equals(_1309_v50))
			_1310_v51 = _nw206
			var _1311_v52 _dafny.Map
			_ = _1311_v52
			_1311_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1241_v0, _1268_v23)
			_1311_v52 = (_1311_v52).Update(_1241_v0, _1268_v23)
		}
		var _1312_i3 _dafny.Int
		_ = _1312_i3
		_1312_i3 = _dafny.Zero
		{
			for !((_1241_v0).Cmp(_1241_v0) != 0) {
				{
					if (_1312_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L14
					}
					_1312_i3 = (_1312_i3).Plus(_dafny.One)
					var _1313_v53 _dafny.Array
					_ = _1313_v53
					var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
					_ = _nw207
					_1313_v53 = _nw207
					var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1313_v53), 0))
					_ = _index162
					(_1313_v53).ArraySet1(_1244_v3, (_index162).Int())
					var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1313_v53), 0))
					_ = _index163
					(_1313_v53).ArraySet1(Companion_Default___.Fm11(globalState), (_index163).Int())
					if !(_1242_v1) || ((_1242_v1) == (_1242_v1)) {
						var _1314_v54 *C2
						_ = _1314_v54
						var _nw208 *C2 = New_C2_()
						_ = _nw208
						_nw208.Ctor__(_1243_v2, _1241_v0)
						_1314_v54 = _nw208
						var _1315_v55 _dafny.Map
						_ = _1315_v55
						_1315_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v1, _1314_v54)
						var _1316_v56 _dafny.Array
						_ = _1316_v56
						var _len0_28 _dafny.Int = _dafny.IntOfInt64(14)
						_ = _len0_28
						var _nw209 _dafny.Array
						_ = _nw209
						if _len0_28.Cmp(_dafny.Zero) == 0 {
							_nw209 = _dafny.NewArray(_len0_28)
						} else {
							var _init28 func(_dafny.Int) bool = (func(_1317_v1 bool) func(_dafny.Int) bool {
								return func(_1318_i4 _dafny.Int) bool {
									return _1317_v1
								}
							})(_1242_v1)
							_ = _init28
							var _element0_28 = _init28(_dafny.Zero)
							_ = _element0_28
							_nw209 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
							(_nw209).ArraySet1(_element0_28, 0)
							var _nativeLen0_28 = (_len0_28).Int()
							_ = _nativeLen0_28
							for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
								(_nw209).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
							}
						}
						_1316_v56 = _nw209
						var _1319_v57 _dafny.Array
						_ = _1319_v57
						var _nwElement0_54 _dafny.Array = _1316_v56
						_ = _nwElement0_54
						var _nw210 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(13))
						_ = _nw210
						(_nw210).ArraySet1(_nwElement0_54, 0)
						(_nw210).ArraySet1(_1316_v56, 1)
						(_nw210).ArraySet1(_1316_v56, 2)
						(_nw210).ArraySet1(_1316_v56, 3)
						(_nw210).ArraySet1(_1316_v56, 4)
						(_nw210).ArraySet1(_1316_v56, 5)
						(_nw210).ArraySet1(_1316_v56, 6)
						(_nw210).ArraySet1(_1316_v56, 7)
						(_nw210).ArraySet1(_1316_v56, 8)
						(_nw210).ArraySet1(_1316_v56, 9)
						(_nw210).ArraySet1(_1316_v56, 10)
						(_nw210).ArraySet1(_1316_v56, 11)
						(_nw210).ArraySet1(_1316_v56, 12)
						_1319_v57 = _nw210
						var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_1316_v56), 0))
						_ = _index164
						(_1316_v56).ArraySet1(_1242_v1, (_index164).Int())
						var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_1316_v56), 0))
						_ = _index165
						var _rhs232 _dafny.Map = _1315_v55
						_ = _rhs232
						var _rhs233 _dafny.Array = _1319_v57
						_ = _rhs233
						var _rhs234 _dafny.Int = _1241_v0
						_ = _rhs234
						var _rhs235 bool = _1242_v1
						_ = _rhs235
						var _lhs188 _dafny.Array = _1316_v56
						_ = _lhs188
						var _lhs189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_1316_v56), 0))
						_ = _lhs189
						_1315_v55 = _rhs232
						_1319_v57 = _rhs233
						_1241_v0 = _rhs234
						(_lhs188).ArraySet1(_rhs235, (_lhs189).Int())
						var _1320_v58 _dafny.Sequence
						_ = _1320_v58
						_1320_v58 = _dafny.UnicodeSeqOfUtf8Bytes("hnmuw")
						_1320_v58 = p0
						(globalState).F7 = (_1316_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_1316_v56), 0))).Int()).(bool)
						var _1321_v59 _dafny.Sequence
						_ = _1321_v59
						_1321_v59 = _dafny.SeqOf(_1320_v58, p0, p0, p0)
						(globalState).F1 = (_1314_v54).Fm3(_dafny.IntOfUint32((_1321_v59).Cardinality()), (_1241_v0).Cmp(_1241_v0) > 0, false, globalState)
						var _1322_v60 _dafny.MultiSet
						_ = _1322_v60
						_1322_v60 = _dafny.MultiSetOf(_1241_v0)
						var _1323_v61 _dafny.Set
						_ = _1323_v61
						_1323_v61 = _dafny.SetOf(_1242_v1)
						var _1324_v62 D7
						_ = _1324_v62
						_1324_v62 = Companion_D7_.Create_DC13_(_1323_v61, _1241_v0, _1241_v0)
						var _1325_v63 *C4
						_ = _1325_v63
						var _nw211 *C4 = New_C4_()
						_ = _nw211
						_nw211.Ctor__(_1322_v60, false, (_1324_v62).Dtor_cf18(), (_1241_v0).Minus(_1241_v0))
						_1325_v63 = _nw211
					} else {
						var _1326_v64 _dafny.MultiSet
						_ = _1326_v64
						_1326_v64 = _dafny.MultiSetOf(Companion_Default___.Fm2(_1241_v0, _1241_v0, globalState))
						_1241_v0 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_1241_v0, Companion_Default___.SafeModuloInt(_1241_v0, (func() _dafny.Int {
							if (_1326_v64).Contains(_1241_v0) {
								return (_1326_v64).Multiplicity(_1241_v0)
							}
							return _dafny.IntOfInt64(-984)
						})())))
						var _1327_v65 D9
						_ = _1327_v65
						_1327_v65 = Companion_D9_.Create_DC19_()
						_1327_v65 = _1327_v65
						_1242_v1 = _1242_v1
						var _1328_v66 _dafny.Sequence
						_ = _1328_v66
						_1328_v66 = _dafny.UnicodeSeqOfUtf8Bytes("rxlji")
						_1328_v66 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1328_v66, p0), (Companion_Default___.SafeIndex(_1241_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1328_v66, p0)).Cardinality()))).Uint32(), _1243_v2)
						var _1329_v67 _dafny.Array
						_ = _1329_v67
						var _nw212 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
						_ = _nw212
						_1329_v67 = _nw212
						var _1330_v68 _dafny.Set
						_ = _1330_v68
						_1330_v68 = _dafny.SetOf(true, _1242_v1, _1242_v1)
						var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1329_v67), 0))
						_ = _index166
						(_1329_v67).ArraySet1((_1330_v68).IsDisjointFrom(_dafny.SetOf(_1242_v1)), (_index166).Int())
						var _1331_v69 _dafny.Sequence
						_ = _1331_v69
						_1331_v69 = _dafny.SeqOf(_1241_v0)
						var _1332_v70 _dafny.Set
						_ = _1332_v70
						_1332_v70 = _dafny.SetOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_1331_v69)).Cardinality())), _1241_v0, _1241_v0, _1241_v0)
						var _1333_v71 _dafny.Array
						_ = _1333_v71
						var _nwElement0_55 _dafny.Int = (_dafny.Zero).Minus(_1241_v0)
						_ = _nwElement0_55
						var _nw213 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(5))
						_ = _nw213
						(_nw213).ArraySet1(_nwElement0_55, 0)
						(_nw213).ArraySet1(_1241_v0, 1)
						(_nw213).ArraySet1((_1241_v0).Times(_1241_v0), 2)
						(_nw213).ArraySet1(_1241_v0, 3)
						(_nw213).ArraySet1((_1241_v0).Plus((_1332_v70).Cardinality()), 4)
						_1333_v71 = _nw213
						var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_1333_v71), 0))
						_ = _index167
						(_1333_v71).ArraySet1(_1241_v0, (_index167).Int())
						var _1334_v72 _dafny.Sequence
						_ = _1334_v72
						_1334_v72 = _dafny.SeqOf(!(_1242_v1), _1242_v1)
						var _1335_v73 _dafny.Map
						_ = _1335_v73
						_1335_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1241_v0, _1334_v72)
						var _1336_v74 _dafny.Map
						_ = _1336_v74
						_1336_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).Fm9((_1335_v73).Cardinality(), (_dafny.Zero).Minus((_dafny.Zero).Minus((_1330_v68).Cardinality())), _1242_v1, globalState)), (_dafny.Zero).Minus(_1241_v0))
						var _1337_v75 _dafny.Sequence
						_ = _1337_v75
						_1337_v75 = _dafny.SeqOf(_1336_v74)
						var _1338_v76 _dafny.MultiSet
						_ = _1338_v76
						_1338_v76 = _dafny.MultiSetOf(Companion_Default___.Fm39(_1241_v0, _1242_v1, _1241_v0, globalState))
						var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1329_v67), 0))
						_ = _index168
						var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_1333_v71), 0))
						_ = _index169
						var _rhs236 bool = (_1338_v76).IsDisjointFrom(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1328_v66, Companion_Default___.Fm1(globalState))))
						_ = _rhs236
						var _rhs237 bool = !(_1242_v1) || (_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_1334_v72, (Companion_Default___.SafeIndex(_1241_v0, _dafny.IntOfUint32((_1334_v72).Cardinality()))).Uint32(), _1242_v1), (func() _dafny.Sequence {
							if (_1335_v73).Contains(_1241_v0) {
								return (_1335_v73).Get(_1241_v0).(_dafny.Sequence)
							}
							return _1334_v72
						})()))
						_ = _rhs237
						var _rhs238 _dafny.Int = (_dafny.Zero).Minus(_1241_v0)
						_ = _rhs238
						var _rhs239 _dafny.Int = (Companion_Default___.SafeModuloInt(_1241_v0, _dafny.IntOfInt64(-527))).Times((_dafny.IntOfInt64(338)).Minus(_dafny.IntOfUint32((p0).Cardinality())))
						_ = _rhs239
						var _rhs240 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1337_v75, (Companion_Default___.SafeIndex(_1241_v0, _dafny.IntOfUint32((_1337_v75).Cardinality()))).Uint32(), (_1336_v74).Merge(_1336_v74))
						_ = _rhs240
						var _lhs190 _dafny.Array = _1329_v67
						_ = _lhs190
						var _lhs191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1329_v67), 0))
						_ = _lhs191
						var _lhs192 _dafny.Array = _1333_v71
						_ = _lhs192
						var _lhs193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_1333_v71), 0))
						_ = _lhs193
						var _lhs194 *GlobalState = globalState
						_ = _lhs194
						_1242_v1 = _rhs236
						(_lhs190).ArraySet1(_rhs237, (_lhs191).Int())
						(_lhs192).ArraySet1(_rhs238, (_lhs193).Int())
						_lhs194.F1 = _rhs239
						_1337_v75 = _rhs240
					}
					(globalState).F7 = !(_1242_v1)
					if _1242_v1 {
						var _1339_v77 _dafny.Sequence
						_ = _1339_v77
						_1339_v77 = _dafny.SeqOf(_1242_v1, _1242_v1)
						var _1340_v78 _dafny.Map
						_ = _1340_v78
						_1340_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1241_v0, _1241_v0)
						var _1341_v79 _dafny.Array
						_ = _1341_v79
						_1341_v79 = p1
						var _1342_v80 _dafny.Set
						_ = _1342_v80
						_1342_v80 = _dafny.SetOf(_1243_v2)
						var _1343_v81 _dafny.Map
						_ = _1343_v81
						_1343_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1341_v79, _1342_v80)
						var _1344_v82 _dafny.Array
						_ = _1344_v82
						var _nwElement0_56 _dafny.Int = _1241_v0
						_ = _nwElement0_56
						var _nw214 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(10))
						_ = _nw214
						(_nw214).ArraySet1(_nwElement0_56, 0)
						(_nw214).ArraySet1(_dafny.IntOfInt64(-279), 1)
						(_nw214).ArraySet1(_1241_v0, 2)
						(_nw214).ArraySet1(_1241_v0, 3)
						(_nw214).ArraySet1(_1241_v0, 4)
						(_nw214).ArraySet1((_dafny.Zero).Minus(_1241_v0), 5)
						(_nw214).ArraySet1(_1241_v0, 6)
						(_nw214).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1339_v77, (Companion_Default___.SafeIndex(Companion_Default___.Fm2(_1241_v0, (_1340_v78).Cardinality(), globalState), _dafny.IntOfUint32((_1339_v77).Cardinality()))).Uint32(), !(_1242_v1))).Cardinality()), 7)
						(_nw214).ArraySet1(_1241_v0, 8)
						(_nw214).ArraySet1(((func() _dafny.Set {
							if (_1343_v81).Contains(_1341_v79) {
								return (_1343_v81).Get(_1341_v79).(_dafny.Set)
							}
							return _dafny.SetOf(_1243_v2)
						})()).Cardinality(), 9)
						_1344_v82 = _nw214
						var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((p1), 0))
						_ = _index170
						(p1).ArraySet1(_1344_v82, (_index170).Int())
						var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((p1), 0))
						_ = _index171
						(p1).ArraySet1(_1344_v82, (_index171).Int())
						var _1345_v83 _dafny.Array
						_ = _1345_v83
						var _nw215 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
						_ = _nw215
						_1345_v83 = _nw215
						var _1346_v84 _dafny.Map
						_ = _1346_v84
						_1346_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v1, _1345_v83)
						var _1347_v86 D15
						_ = _1347_v86
						_1347_v86 = Companion_D15_.Create_DC33_(_1346_v84)
						var _1348_v87 _dafny.Map
						_ = _1348_v87
						_1348_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
							var _coll62 = _dafny.NewMapBuilder()
							_ = _coll62
							for _iter64 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-81), _dafny.IntOfInt64(373))); ; {
								_compr_62, _ok64 := _iter64()
								if !_ok64 {
									break
								}
								var _1349_v85 _dafny.Int
								_1349_v85 = interface{}(_compr_62).(_dafny.Int)
								if ((_dafny.IntOfInt64(-81)).Cmp(_1349_v85) <= 0) && ((_1349_v85).Cmp(_dafny.IntOfInt64(373)) < 0) {
									_coll62.Add(Companion_Default___.SafeModuloInt(_1349_v85, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v1, _1241_v0)).Cardinality()), (_dafny.MultiSetOf(_1241_v0)).Cardinality())
								}
							}
							return _coll62.ToMap()
						}(), (_1347_v86).Dtor_cf49())
						var _1350_v88 D16
						_ = _1350_v88
						_1350_v88 = Companion_D16_.Create_DC38_(_1340_v78)
						_1346_v84 = (func() _dafny.Map {
							if (_1348_v87).Contains((_1350_v88).Dtor_cf61()) {
								return (_1348_v87).Get((_1350_v88).Dtor_cf61()).(_dafny.Map)
							}
							return _1346_v84
						})()
						var _1351_v89 _dafny.Sequence
						_ = _1351_v89
						_1351_v89 = _dafny.SeqOf((_dafny.Zero).Minus(_1241_v0), _dafny.IntOfInt64(-243))
						(globalState).F7 = (_dafny.IntOfUint32((_1351_v89).Cardinality())).Cmp(_1241_v0) < 0
						var _arr0 _dafny.Array = _dafny.ArrayCastTo((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((p1), 0))).Int()))
						_ = _arr0
						var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(724), _dafny.ArrayLen((_dafny.ArrayCastTo((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((p1), 0))).Int()))), 0))
						_ = _index172
						(_arr0).ArraySet1(_dafny.IntOfInt64(161), (_index172).Int())
						var _arr1 _dafny.Array = _dafny.ArrayCastTo((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((p1), 0))).Int()))
						_ = _arr1
						var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(724), _dafny.ArrayLen((_dafny.ArrayCastTo((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((p1), 0))).Int()))), 0))
						_ = _index173
						var _rhs241 _dafny.Int = (_1241_v0).Times(_1241_v0)
						_ = _rhs241
						var _rhs242 _dafny.CodePoint = _1243_v2
						_ = _rhs242
						var _lhs195 _dafny.Array = _dafny.ArrayCastTo((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((p1), 0))).Int()))
						_ = _lhs195
						var _lhs196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(724), _dafny.ArrayLen((_dafny.ArrayCastTo((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((p1), 0))).Int()))), 0))
						_ = _lhs196
						(_lhs195).ArraySet1(_rhs241, (_lhs196).Int())
						_1243_v2 = _rhs242
						var _1352_v90 _dafny.Map
						_ = _1352_v90
						_1352_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1244_v3).Merge((_1313_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1313_v53), 0))).Int()).(_dafny.Map))).Cardinality(), _1351_v89)
						_1352_v90 = _1352_v90
					} else {
						(globalState).F1 = _1241_v0
						var _1353_v91 _dafny.Array
						_ = _1353_v91
						var _nw216 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
						_ = _nw216
						_1353_v91 = _nw216
						var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1353_v91), 0))
						_ = _index174
						(_1353_v91).ArraySet1(_1242_v1, (_index174).Int())
						var _1354_v92 _dafny.Map
						_ = _1354_v92
						_1354_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1241_v0, !(false))
						var _1355_v93 _dafny.Map
						_ = _1355_v93
						_1355_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v1, _1241_v0)
						var _1356_v94 _dafny.Sequence
						_ = _1356_v94
						_1356_v94 = _dafny.SeqOf(_1242_v1)
						var _1357_v95 _dafny.Sequence
						_ = _1357_v95
						_1357_v95 = _dafny.SeqOf(_1241_v0)
						var _1358_v96 _dafny.Array
						_ = _1358_v96
						var _nwElement0_57 _dafny.Int = (_dafny.Zero).Minus((_1241_v0).Plus(_dafny.IntOfInt64(755)))
						_ = _nwElement0_57
						var _nw217 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(17))
						_ = _nw217
						(_nw217).ArraySet1(_nwElement0_57, 0)
						(_nw217).ArraySet1(Companion_Default___.SafeModuloInt((_1354_v92).Cardinality(), _1241_v0), 1)
						(_nw217).ArraySet1((_dafny.Zero).Minus(_1241_v0), 2)
						(_nw217).ArraySet1(_1241_v0, 3)
						(_nw217).ArraySet1((_dafny.IntOfUint32((p0).Cardinality())).Times(_1241_v0), 4)
						(_nw217).ArraySet1((func() _dafny.Int {
							if (_1355_v93).Contains(_1242_v1) {
								return (_1355_v93).Get(_1242_v1).(_dafny.Int)
							}
							return (_dafny.Zero).Minus(_1241_v0)
						})(), 5)
						(_nw217).ArraySet1(_1241_v0, 6)
						(_nw217).ArraySet1(_dafny.IntOfInt64(784), 7)
						(_nw217).ArraySet1(_dafny.IntOfUint32((_1356_v94).Cardinality()), 8)
						(_nw217).ArraySet1(_1241_v0, 9)
						(_nw217).ArraySet1(_1241_v0, 10)
						(_nw217).ArraySet1(_1241_v0, 11)
						(_nw217).ArraySet1((_1241_v0).Plus(_dafny.IntOfUint32((_1356_v94).Cardinality())), 12)
						(_nw217).ArraySet1((_1357_v95).Select((Companion_Default___.SafeIndex(_1241_v0, _dafny.IntOfUint32((_1357_v95).Cardinality()))).Uint32()).(_dafny.Int), 13)
						(_nw217).ArraySet1(_dafny.IntOfUint32((_1356_v94).Cardinality()), 14)
						(_nw217).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 15)
						(_nw217).ArraySet1((_dafny.Zero).Minus(_1241_v0), 16)
						_1358_v96 = _nw217
						var _1359_v97 _dafny.Array
						_ = _1359_v97
						var _nw218 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
						_ = _nw218
						_1359_v97 = _nw218
						var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1359_v97), 0))
						_ = _index175
						(_1359_v97).ArraySet1(_1353_v91, (_index175).Int())
						var _1360_v98 _dafny.Sequence
						_ = _1360_v98
						_1360_v98 = _dafny.UnicodeSeqOfUtf8Bytes("isyb")
						var _1361_v99 _dafny.Set
						_ = _1361_v99
						_1361_v99 = _dafny.SetOf(_1241_v0, _1241_v0)
						var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1353_v91), 0))
						_ = _index176
						var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1359_v97), 0))
						_ = _index177
						var _rhs243 bool = false
						_ = _rhs243
						var _rhs244 _dafny.Array = (func() _dafny.Array {
							if !(_1242_v1) {
								return _1358_v96
							}
							return _1358_v96
						})()
						_ = _rhs244
						var _rhs245 _dafny.Array = _1353_v91
						_ = _rhs245
						var _rhs246 _dafny.Sequence = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm1(globalState), (Companion_Default___.SafeIndex(((_1361_v99).Difference(_1361_v99)).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm1(globalState)).Cardinality()))).Uint32(), _1243_v2)
						_ = _rhs246
						var _lhs197 _dafny.Array = _1353_v91
						_ = _lhs197
						var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1353_v91), 0))
						_ = _lhs198
						var _lhs199 _dafny.Array = _1359_v97
						_ = _lhs199
						var _lhs200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1359_v97), 0))
						_ = _lhs200
						(_lhs197).ArraySet1(_rhs243, (_lhs198).Int())
						_1358_v96 = _rhs244
						(_lhs199).ArraySet1(_rhs245, (_lhs200).Int())
						_1360_v98 = _rhs246
						var _1362_v100 _dafny.Set
						_ = _1362_v100
						_1362_v100 = _dafny.SetOf(_dafny.ArrayCastTo((_1359_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1359_v97), 0))).Int())), _dafny.ArrayCastTo((_1359_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1359_v97), 0))).Int())))
						var _1363_v101 _dafny.Map
						_ = _1363_v101
						_1363_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1362_v100).Union(_1362_v100), _1241_v0)
						_1363_v101 = (_1363_v101).Update(_1362_v100, (_1355_v93).Cardinality())
						_1358_v96 = _1358_v96
						var _1364_v102 *C0
						_ = _1364_v102
						var _nw219 *C0 = New_C0_()
						_ = _nw219
						_nw219.Ctor__(_1356_v94, (_dafny.Zero).Minus((func() _dafny.Int {
							if !(_1242_v1) {
								return _1241_v0
							}
							return _1241_v0
						})()))
						_1364_v102 = _nw219
					}
					goto C14
				}
			C14:
			}
			goto L14
		}
	L14:
		var _1365_v103 _dafny.Map
		_ = _1365_v103
		_1365_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1241_v0)
		var _1366_v104 _dafny.Set
		_ = _1366_v104
		_1366_v104 = _dafny.SetOf((_dafny.Zero).Minus((_1365_v103).Cardinality()), _1241_v0, _1241_v0)
		var _1367_v105 _dafny.MultiSet
		_ = _1367_v105
		_1367_v105 = _dafny.MultiSetOf(_1241_v0)
		(globalState).F1 = ((func() _dafny.Int {
			if _1242_v1 {
				return _1241_v0
			}
			return (_1366_v104).Cardinality()
		})()).Plus((func() _dafny.Int {
			if (_1367_v105).Contains(_1241_v0) {
				return (_1367_v105).Multiplicity(_1241_v0)
			}
			return _1241_v0
		})())
		var _1368_v106 _dafny.Sequence
		_ = _1368_v106
		_1368_v106 = _dafny.SeqOf(_1242_v1, _1242_v1)
		var _1369_v107 *C2
		_ = _1369_v107
		var _nw220 *C2 = New_C2_()
		_ = _nw220
		_nw220.Ctor__(_1243_v2, _1241_v0)
		_1369_v107 = _nw220
		var _1370_v108 _dafny.MultiSet
		_ = _1370_v108
		_1370_v108 = _dafny.MultiSetOf(_1369_v107)
		var _1371_v109 _dafny.Array
		_ = _1371_v109
		var _nwElement0_58 _dafny.Int = _1241_v0
		_ = _nwElement0_58
		var _nw221 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(14))
		_ = _nw221
		(_nw221).ArraySet1(_nwElement0_58, 0)
		(_nw221).ArraySet1((_1365_v103).Cardinality(), 1)
		(_nw221).ArraySet1(_dafny.IntOfUint32((_1368_v106).Cardinality()), 2)
		(_nw221).ArraySet1((_1365_v103).Cardinality(), 3)
		(_nw221).ArraySet1(_dafny.IntOfInt64(83), 4)
		(_nw221).ArraySet1(_dafny.IntOfInt64(347), 5)
		(_nw221).ArraySet1(_1241_v0, 6)
		(_nw221).ArraySet1(_1241_v0, 7)
		(_nw221).ArraySet1(_1241_v0, 8)
		(_nw221).ArraySet1(_1241_v0, 9)
		(_nw221).ArraySet1(_dafny.IntOfInt64(956), 10)
		(_nw221).ArraySet1(_1241_v0, 11)
		(_nw221).ArraySet1((_dafny.Zero).Minus(_1241_v0), 12)
		(_nw221).ArraySet1((_1370_v108).Cardinality(), 13)
		_1371_v109 = _nw221
		var _1372_v110 _dafny.Array
		_ = _1372_v110
		var _nw222 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
		_ = _nw222
		_1372_v110 = _nw222
		var _1373_v111 _dafny.Map
		_ = _1373_v111
		_1373_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1371_v109, _1372_v110)
		(globalState).F7 = (!(_1242_v1)) && ((_1373_v111).Contains(_1371_v109))
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f12 _dafny.Set
	_f11 _dafny.Int
	_f15 D2
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f12 = _dafny.EmptySet
	_this._f11 = _dafny.Zero
	_this._f15 = Companion_D2_.Default()
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C6{}
var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F12() _dafny.Set {
	return _this._f12
}
func (_this *C6) F11() _dafny.Int {
	return _this._f11
}
func (_this *C6) Ctor__(f15 D2, f12 _dafny.Set, f11 _dafny.Int) {
	{
		(_this)._f15 = f15
		(_this)._f12 = f12
		(_this)._f11 = f11
	}
}
func (_this *C6) Fm5(p0 D0, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		if false {
			return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uoywkyi")).Cardinality())).Cmp((_this).F11()) <= 0
		} else {
			return !(!(true))
		}
	}
}
func (_this *C6) Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((_this).F11(), ((func() _dafny.Map {
			if !(!(true)) {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)
		})()).Cardinality())
	}
}
func (_this *C6) Fm3(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((func() _dafny.Map {
			var _coll63 = _dafny.NewMapBuilder()
			_ = _coll63
			for _iter65 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('p'), _dafny.CodePoint('y'))).Elements()); ; {
				_compr_63, _ok65 := _iter65()
				if !_ok65 {
					break
				}
				var _1374_v0 _dafny.CodePoint
				_1374_v0 = interface{}(_compr_63).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('p'), _dafny.CodePoint('y')), _1374_v0) {
					_coll63.Add(_1374_v0, _dafny.UnicodeSeqOfUtf8Bytes("mhdueatrb"))
				}
			}
			return _coll63.ToMap()
		}()).Cardinality()).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-288))).Uint32(), func(coer86 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg86 _dafny.Int) interface{} {
				return coer86(arg86)
			}
		}(func(_1375_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(941))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg87 _dafny.Int) interface{} {
				return coer87(arg87)
			}
		}(func(_1376_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		})))).Cardinality())))
	}
}
func (_this *C6) Fm4(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		var _source20 D2 = (_this).F15()
		_ = _source20
		if _source20.Is_DC4() {
			var _1377___mcc_h0 _dafny.Int = _source20.Get_().(D2_DC4).Cf6
			_ = _1377___mcc_h0
			var _1378___mcc_h1 _dafny.Int = _source20.Get_().(D2_DC4).Cf7
			_ = _1378___mcc_h1
			var _1379___mcc_h2 _dafny.Int = _source20.Get_().(D2_DC4).Cf8
			_ = _1379___mcc_h2
			var _1380_cf8 _dafny.Int = _1379___mcc_h2
			_ = _1380_cf8
			var _1381_cf7 _dafny.Int = _1378___mcc_h1
			_ = _1381_cf7
			var _1382_cf6 _dafny.Int = _1377___mcc_h0
			_ = _1382_cf6
			return true
		} else if _source20.Is_DC5() {
			return (Companion_D0_.Create_DC1_(true, _dafny.IntOfInt64(252), true)).Dtor_cf1()
		} else {
			var _1383___mcc_h3 _dafny.MultiSet = _source20.Get_().(D2_DC3).Cf5
			_ = _1383___mcc_h3
			var _1384_cf5 _dafny.MultiSet = _1383___mcc_h3
			_ = _1384_cf5
			return ((_this).F12()).IsDisjointFrom(_dafny.SetOf((Companion_D0_.Create_DC1_(false, (_this).F11(), true)).Dtor_cf3()))
		}
	}
}
func (_this *C6) M0(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var _1385_v0 _dafny.Array
		_ = _1385_v0
		var _nw223 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(8))
		_ = _nw223
		_1385_v0 = _nw223
		var _1386_v1 _dafny.Array
		_ = _1386_v1
		var _nw224 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw224
		_1386_v1 = _nw224
		var _1387_v2 _dafny.Set
		_ = _1387_v2
		_1387_v2 = _dafny.SetOf(_1386_v1)
		var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_1385_v0), 0))
		_ = _index178
		(_1385_v0).ArraySet1(_1387_v2, (_index178).Int())
		var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_1385_v0), 0))
		_ = _index179
		(_1385_v0).ArraySet1(_1387_v2, (_index179).Int())
		var _1388_v3 bool
		_ = _1388_v3
		_1388_v3 = true
		var _1389_i0 _dafny.Int
		_ = _1389_i0
		_1389_i0 = _dafny.Zero
		{
			for _1388_v3 {
				{
					if (_1389_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L15
					}
					_1389_i0 = (_1389_i0).Plus(_dafny.One)
					var _1390_v4 _dafny.Map
					_ = _1390_v4
					_1390_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1388_v3, _1388_v3)
					_1390_v4 = _1390_v4
					var _1391_v5 _dafny.MultiSet
					_ = _1391_v5
					_1391_v5 = _dafny.MultiSetOf((_this).F11())
					var _1392_v6 _dafny.Map
					_ = _1392_v6
					_1392_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1391_v5, !(_1388_v3) || (_1388_v3))
					(globalState).F1 = (_1392_v6).Cardinality()
					_1388_v3 = Companion_Default___.Fm0(globalState)
					(globalState).F6 = _1388_v3
					goto C15
				}
			C15:
			}
			goto L15
		}
	L15:
		(globalState).F1 = (_this).F11()
		if (func() bool {
			if _1388_v3 {
				return _1388_v3
			}
			return _1388_v3
		})() {
			var _1393_v7 _dafny.Map
			_ = _1393_v7
			_1393_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _dafny.IntOfInt64(194))
			var _1394_v8 _dafny.Array
			_ = _1394_v8
			var _nwElement0_59 _dafny.Int = _dafny.IntOfInt64(-194)
			_ = _nwElement0_59
			var _nw225 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(6))
			_ = _nw225
			(_nw225).ArraySet1(_nwElement0_59, 0)
			(_nw225).ArraySet1((_this).F11(), 1)
			(_nw225).ArraySet1((_this).Fm3((_this).F11(), _1388_v3, _1388_v3, globalState), 2)
			(_nw225).ArraySet1((_this).F11(), 3)
			(_nw225).ArraySet1((_this).F11(), 4)
			(_nw225).ArraySet1(Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_1393_v7).Contains((_this).F11()) {
					return (_1393_v7).Get((_this).F11()).(_dafny.Int)
				}
				return (_this).F11()
			})(), (_this).F11()), 5)
			_1394_v8 = _nw225
			var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_1394_v8), 0))
			_ = _index180
			(_1394_v8).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F11(), (_this).F11())), (_index180).Int())
			var _1395_v9 _dafny.Sequence
			_ = _1395_v9
			_1395_v9 = _dafny.SeqOf(!(_1388_v3), _1388_v3)
			var _1396_v10 _dafny.Sequence
			_ = _1396_v10
			_1396_v10 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1388_v3))
			var _1397_v11 _dafny.Sequence
			_ = _1397_v11
			_1397_v11 = _dafny.SeqOf((_this).F11(), (_this).F11(), (_this).F11())
			var _1398_v12 _dafny.Map
			_ = _1398_v12
			_1398_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F12()).Cardinality(), false)
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_1394_v8), 0))
			_ = _index181
			var _rhs247 _dafny.Int = (_this).F11()
			_ = _rhs247
			var _rhs248 bool = (_1395_v9).Select((Companion_Default___.SafeIndex((((_1396_v10).Select((Companion_Default___.SafeIndex((_this).Fm3(_dafny.IntOfUint32((_1397_v11).Cardinality()), _1388_v3, _1388_v3, globalState), _dafny.IntOfUint32((_1396_v10).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_1398_v12)).Cardinality(), _dafny.IntOfUint32((_1395_v9).Cardinality()))).Uint32()).(bool)
			_ = _rhs248
			var _rhs249 bool = _1388_v3
			_ = _rhs249
			var _lhs201 _dafny.Array = _1394_v8
			_ = _lhs201
			var _lhs202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_1394_v8), 0))
			_ = _lhs202
			var _lhs203 *GlobalState = globalState
			_ = _lhs203
			var _lhs204 *GlobalState = globalState
			_ = _lhs204
			(_lhs201).ArraySet1(_rhs247, (_lhs202).Int())
			_lhs203.F7 = _rhs248
			_lhs204.F6 = _rhs249
			var _1399_v13 _dafny.Array
			_ = _1399_v13
			var _nw226 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
			_ = _nw226
			_1399_v13 = _nw226
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_1399_v13), 0))
			_ = _index182
			(_1399_v13).ArraySet1(_1395_v9, (_index182).Int())
			var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_1399_v13), 0))
			_ = _index183
			(_1399_v13).ArraySet1(_1395_v9, (_index183).Int())
			var _1400_v14 _dafny.Map
			_ = _1400_v14
			_1400_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F11())
			(globalState).F1 = (Companion_Default___.Fm2(_dafny.IntOfInt64(384), (_1394_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_1394_v8), 0))).Int()).(_dafny.Int), globalState)).Plus((_dafny.Zero).Minus((func() _dafny.Int {
				if (_1393_v7).Contains((_1394_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_1394_v8), 0))).Int()).(_dafny.Int)) {
					return (_1393_v7).Get((_1394_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_1394_v8), 0))).Int()).(_dafny.Int)).(_dafny.Int)
				}
				return (_1400_v14).Cardinality()
			})()))
			var _1401_v15 _dafny.Sequence
			_ = _1401_v15
			_1401_v15 = _dafny.UnicodeSeqOfUtf8Bytes("irx")
			_1401_v15 = _dafny.Companion_Sequence_.Update(_1401_v15, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.IntOfUint32((_1401_v15).Cardinality()))).Uint32(), _dafny.CodePoint('m'))
			var _1402_v16 *C2
			_ = _1402_v16
			var _nw227 *C2 = New_C2_()
			_ = _nw227
			_nw227.Ctor__(p0, ((_this).F11()).Minus((_this).F11()))
			_1402_v16 = _nw227
		} else {
			var _1403_v17 _dafny.Map
			_ = _1403_v17
			_1403_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1388_v3, _1388_v3)
			var _1404_v18 D8
			_ = _1404_v18
			_1404_v18 = Companion_D8_.Create_DC15_(_1403_v17)
			var _1405_v19 D8
			_ = _1405_v19
			_1405_v19 = Companion_D8_.Create_DC17_(Companion_D8_.Create_DC17_(_1404_v18))
			var _1406_v20 _dafny.Map
			_ = _1406_v20
			_1406_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1405_v19, (_this).F11())
			var _1407_v21 D13
			_ = _1407_v21
			_1407_v21 = Companion_D13_.Create_DC27_(_1406_v20)
			var _1408_v22 D13
			_ = _1408_v22
			_1408_v22 = Companion_D13_.Create_DC29_(_1407_v21)
			var _source21 D13 = _1408_v22
			_ = _source21
			if _source21.Is_DC28() {
				var _1409___mcc_h0 _dafny.Sequence = _source21.Get_().(D13_DC28).Cf40
				_ = _1409___mcc_h0
				var _1410___mcc_h1 _dafny.Int = _source21.Get_().(D13_DC28).Cf41
				_ = _1410___mcc_h1
				var _1411___mcc_h2 bool = _source21.Get_().(D13_DC28).Cf42
				_ = _1411___mcc_h2
				var _1412_cf42 bool = _1411___mcc_h2
				_ = _1412_cf42
				var _1413_cf41 _dafny.Int = _1410___mcc_h1
				_ = _1413_cf41
				var _1414_cf40 _dafny.Sequence = _1409___mcc_h0
				_ = _1414_cf40
				var _1415_v23 _dafny.MultiSet
				_ = _1415_v23
				_1415_v23 = _dafny.MultiSetOf(_1412_cf42)
				var _1416_v24 _dafny.Map
				_ = _1416_v24
				_1416_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1415_v23).Cardinality(), _1412_cf42)
				var _1417_v25 _dafny.Map
				_ = _1417_v25
				_1417_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1388_v3, Companion_D12_.Create_DC26_(_1416_v24))
				var _1418_v26 _dafny.MultiSet
				_ = _1418_v26
				_1418_v26 = _dafny.MultiSetOf(_1417_v25, _1417_v25)
				var _1419_v27 _dafny.Sequence
				_ = _1419_v27
				_1419_v27 = _dafny.SeqOf(_1413_cf41, _dafny.IntOfInt64(262))
				var _1420_v28 _dafny.Map
				_ = _1420_v28
				_1420_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1412_cf42), _1413_cf41)
				var _1421_v29 _dafny.Array
				_ = _1421_v29
				var _nwElement0_60 _dafny.Int = (_this).F11()
				_ = _nwElement0_60
				var _nw228 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(20))
				_ = _nw228
				(_nw228).ArraySet1(_nwElement0_60, 0)
				(_nw228).ArraySet1(Companion_Default___.SafeDivisionInt(_1413_cf41, _dafny.IntOfInt64(328)), 1)
				(_nw228).ArraySet1((_this).F11(), 2)
				(_nw228).ArraySet1((_this).F11(), 3)
				(_nw228).ArraySet1(_1413_cf41, 4)
				(_nw228).ArraySet1(_dafny.IntOfInt64(6), 5)
				(_nw228).ArraySet1(_1413_cf41, 6)
				(_nw228).ArraySet1((_this).F11(), 7)
				(_nw228).ArraySet1(((_dafny.Zero).Minus((func() _dafny.Int {
					if (_1418_v26).Contains(_1417_v25) {
						return (_1418_v26).Multiplicity(_1417_v25)
					}
					return (_this).F11()
				})())).Plus((_this).F11()), 8)
				(_nw228).ArraySet1((_this).F11(), 9)
				(_nw228).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1419_v27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(371))).Uint32(), func(coer88 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg88 _dafny.Int) interface{} {
						return coer88(arg88)
					}
				}(func(_1422_i1 _dafny.Int) _dafny.Int {
					return (_this).F11()
				})))).Cardinality()), 10)
				(_nw228).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F11(), (func() _dafny.Int {
					if (_1420_v28).Contains(_1388_v3) {
						return (_1420_v28).Get(_1388_v3).(_dafny.Int)
					}
					return (_this).F11()
				})()), 11)
				(_nw228).ArraySet1(_1413_cf41, 12)
				(_nw228).ArraySet1(_1413_cf41, 13)
				(_nw228).ArraySet1(((_1415_v23).Cardinality()).Minus((_this).Fm6(true, (_this).F11(), globalState)), 14)
				(_nw228).ArraySet1(_1413_cf41, 15)
				(_nw228).ArraySet1((_this).F11(), 16)
				(_nw228).ArraySet1(_1413_cf41, 17)
				(_nw228).ArraySet1(_1413_cf41, 18)
				(_nw228).ArraySet1(_dafny.IntOfInt64(798), 19)
				_1421_v29 = _nw228
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_1421_v29), 0))
				_ = _index184
				(_1421_v29).ArraySet1((_this).F11(), (_index184).Int())
				var _1423_v30 _dafny.Map
				_ = _1423_v30
				_1423_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1414_cf40, (_1415_v23).Cardinality())
				var _1424_v31 _dafny.Map
				_ = _1424_v31
				_1424_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1412_cf42), _1414_cf40)
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_1421_v29), 0))
				_ = _index185
				(_1421_v29).ArraySet1((func() _dafny.Int {
					if (_1423_v30).Contains((func() _dafny.Sequence {
						if (_1424_v31).Contains(_1388_v3) {
							return (_1424_v31).Get(_1388_v3).(_dafny.Sequence)
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-548))).Uint32(), func(coer89 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg89 _dafny.Int) interface{} {
								return coer89(arg89)
							}
						}(func(_1425_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('f')
						}))
					})()) {
						return (_1423_v30).Get((func() _dafny.Sequence {
							if (_1424_v31).Contains(_1388_v3) {
								return (_1424_v31).Get(_1388_v3).(_dafny.Sequence)
							}
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-548))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg90 _dafny.Int) interface{} {
									return coer90(arg90)
								}
							}(func(_1426_i2 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('f')
							}))
						})()).(_dafny.Int)
					}
					return (_this).Fm3((_this).F11(), _1412_cf42, _1388_v3, globalState)
				})(), (_index185).Int())
				var _1427_v32 _dafny.Map
				_ = _1427_v32
				_1427_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1421_v29, _1388_v3)
				var _1428_v33 D10
				_ = _1428_v33
				_1428_v33 = Companion_D10_.Create_DC20_(_1427_v32)
				var _1429_v34 _dafny.MultiSet
				_ = _1429_v34
				_1429_v34 = _dafny.MultiSetOf(_1428_v33, _1428_v33)
				_1429_v34 = _1429_v34
				var _1430_v35 *C1
				_ = _1430_v35
				var _nw229 *C1 = New_C1_()
				_ = _nw229
				_nw229.Ctor__(false, _1412_cf42)
				_1430_v35 = _nw229
				var _1431_v36 _dafny.Sequence
				_ = _1431_v36
				_1431_v36 = _dafny.SeqOf(_1430_v35)
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_1421_v29), 0))
				_ = _index186
				var _rhs250 _dafny.Int = ((_1416_v24).Cardinality()).Times(_1413_cf41)
				_ = _rhs250
				var _rhs251 _dafny.Int = (_dafny.Zero).Minus(((_dafny.Zero).Minus((_dafny.Zero).Minus((_1413_cf41).Times((_1421_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_1421_v29), 0))).Int()).(_dafny.Int))))).Minus(_1413_cf41))
				_ = _rhs251
				var _rhs252 bool = (_1430_v35).F19()
				_ = _rhs252
				var _rhs253 _dafny.Sequence = _1431_v36
				_ = _rhs253
				var _lhs205 _dafny.Array = _1421_v29
				_ = _lhs205
				var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_1421_v29), 0))
				_ = _lhs206
				var _lhs207 *GlobalState = globalState
				_ = _lhs207
				var _lhs208 *GlobalState = globalState
				_ = _lhs208
				(_lhs205).ArraySet1(_rhs250, (_lhs206).Int())
				_lhs207.F1 = _rhs251
				_lhs208.F6 = _rhs252
				_1431_v36 = _rhs253
				(globalState).F1 = (_this).F11()
			} else if _source21.Is_DC27() {
				var _1432___mcc_h3 _dafny.Map = _source21.Get_().(D13_DC27).Cf39
				_ = _1432___mcc_h3
				var _1433_cf39 _dafny.Map = _1432___mcc_h3
				_ = _1433_cf39
				var _1434_v37 _dafny.Sequence
				_ = _1434_v37
				_1434_v37 = _dafny.SeqOf(_1388_v3, _1388_v3, _1388_v3)
				var _1435_v38 _dafny.Map
				_ = _1435_v38
				_1435_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F11()), _dafny.SeqOf((_this).F11(), (_dafny.Zero).Minus((_this).F11()), _dafny.IntOfUint32((_1434_v37).Cardinality()))), (_this).F11())
				_1435_v38 = _1435_v38
				var _1436_v39 _dafny.Sequence
				_ = _1436_v39
				_1436_v39 = _dafny.SeqOf((_this).F11())
				var _1437_v40 _dafny.Sequence
				_ = _1437_v40
				_1437_v40 = _dafny.UnicodeSeqOfUtf8Bytes("bdlkv")
				var _1438_v41 _dafny.Map
				_ = _1438_v41
				_1438_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_dafny.IntOfInt64(78), (_this).F11(), globalState), (_this).F11())
				var _1439_v42 _dafny.MultiSet
				_ = _1439_v42
				_1439_v42 = _dafny.MultiSetOf((_this).F11())
				var _1440_v43 _dafny.Map
				_ = _1440_v43
				_1440_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1439_v42, _1438_v41)
				var _1441_v44 _dafny.Array
				_ = _1441_v44
				var _nwElement0_61 _dafny.Int = (_this).F11()
				_ = _nwElement0_61
				var _nw230 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(16))
				_ = _nw230
				(_nw230).ArraySet1(_nwElement0_61, 0)
				(_nw230).ArraySet1(((_this).F11()).Times((_this).F11()), 1)
				(_nw230).ArraySet1((_this).F11(), 2)
				(_nw230).ArraySet1(_dafny.IntOfInt64(-574), 3)
				(_nw230).ArraySet1((_this).F11(), 4)
				(_nw230).ArraySet1((_this).F11(), 5)
				(_nw230).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfUint32((_1436_v39).Cardinality())).Plus((_this).F11())), 6)
				(_nw230).ArraySet1((_this).F11(), 7)
				(_nw230).ArraySet1((_dafny.IntOfUint32((_1437_v40).Cardinality())).Times((_this).F11()), 8)
				(_nw230).ArraySet1((_1403_v17).Cardinality(), 9)
				(_nw230).ArraySet1((_this).F11(), 10)
				(_nw230).ArraySet1(Companion_Default___.SafeModuloInt((_1438_v41).Cardinality(), (_this).F11()), 11)
				(_nw230).ArraySet1((_this).F11(), 12)
				(_nw230).ArraySet1((_this).F11(), 13)
				(_nw230).ArraySet1((_dafny.Zero).Minus((_1436_v39).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_1436_v39).Cardinality()))).Uint32()).(_dafny.Int)), 14)
				(_nw230).ArraySet1(((_this).Fm3((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F11())), (_this).Fm4((_this).F11(), _1440_v43, (_this).F11(), _1388_v3, globalState), _1388_v3, globalState)).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1388_v3)).Cardinality()), 15)
				_1441_v44 = _nw230
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_1441_v44), 0))
				_ = _index187
				(_1441_v44).ArraySet1(((_1439_v42).Union(_1439_v42)).Cardinality(), (_index187).Int())
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_1441_v44), 0))
				_ = _index188
				(_1441_v44).ArraySet1((_this).F11(), (_index188).Int())
				var _1442_v45 _dafny.Map
				_ = _1442_v45
				_1442_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1386_v1, (_this).F11())
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_1441_v44), 0))
				_ = _index189
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_1441_v44), 0))
				_ = _index190
				var _rhs254 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1434_v37).Cardinality()), (_this).F11())
				_ = _rhs254
				var _rhs255 _dafny.Int = ((func() _dafny.Int {
					if (_1442_v45).Contains(_1386_v1) {
						return (_1442_v45).Get(_1386_v1).(_dafny.Int)
					}
					return (_this).F11()
				})()).Minus((_this).F11())
				_ = _rhs255
				var _lhs209 _dafny.Array = _1441_v44
				_ = _lhs209
				var _lhs210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_1441_v44), 0))
				_ = _lhs210
				var _lhs211 _dafny.Array = _1441_v44
				_ = _lhs211
				var _lhs212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_1441_v44), 0))
				_ = _lhs212
				(_lhs209).ArraySet1(_rhs254, (_lhs210).Int())
				(_lhs211).ArraySet1(_rhs255, (_lhs212).Int())
				var _1443_v46 D10
				_ = _1443_v46
				_1443_v46 = Companion_D10_.Create_DC21_((func() bool {
					if (_1403_v17).Contains(_1388_v3) {
						return (_1403_v17).Get(_1388_v3).(bool)
					}
					return _1388_v3
				})(), _1388_v3)
				var _1444_v47 _dafny.Map
				_ = _1444_v47
				_1444_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1443_v46, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1441_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_1441_v44), 0))).Int()).(_dafny.Int), (_this).F11()))
				_1444_v47 = (_1444_v47).Update(_1443_v46, _1438_v41)
				var _1445_v48 _dafny.Array
				_ = _1445_v48
				var _nw231 _dafny.Array = _dafny.NewArrayWithValue(Companion_D12_.Default(), _dafny.IntOfInt64(25))
				_ = _nw231
				_1445_v48 = _nw231
				var _1446_v49 _dafny.Array
				_ = _1446_v49
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_29
				var _nw232 _dafny.Array
				_ = _nw232
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw232 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) D8 = (func(_1447_v3 bool) func(_dafny.Int) D8 {
						return func(_1448_i3 _dafny.Int) D8 {
							return Companion_D8_.Create_DC16_(_1447_v3, _dafny.IntOfInt64(-553))
						}
					})(_1388_v3)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw232 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw232).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw232).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1446_v49 = _nw232
				var _1449_v50 D12
				_ = _1449_v50
				_1449_v50 = Companion_D12_.Create_DC25_(_1446_v49)
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_1445_v48), 0))
				_ = _index191
				(_1445_v48).ArraySet1(_1449_v50, (_index191).Int())
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_1445_v48), 0))
				_ = _index192
				(_1445_v48).ArraySet1(Companion_D12_.Create_DC25_(_1446_v49), (_index192).Int())
			} else {
				var _1450___mcc_h4 D13 = _source21.Get_().(D13_DC29).Cf43
				_ = _1450___mcc_h4
				var _1451_cf43 D13 = _1450___mcc_h4
				_ = _1451_cf43
				var _1452_v51 _dafny.Sequence
				_ = _1452_v51
				_1452_v51 = _dafny.UnicodeSeqOfUtf8Bytes("x")
				_1452_v51 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("djqq"), _1452_v51)
				var _1453_v52 _dafny.CodePoint
				_ = _1453_v52
				_1453_v52 = _dafny.CodePoint('k')
				_1453_v52 = p0
				var _1454_v53 D15
				_ = _1454_v53
				_1454_v53 = Companion_D15_.Create_DC36_((_this).F11(), _1388_v3, _1388_v3, _1388_v3, (_dafny.Zero).Minus((_this).F11()))
				(globalState).F1 = (_this).Fm6((_1454_v53).Dtor_cf58(), (_this).F11(), globalState)
				_1453_v52 = p0
			}
			var _1455_v54 D15
			_ = _1455_v54
			_1455_v54 = Companion_D15_.Create_DC34_()
			var _1456_v55 _dafny.Set
			_ = _1456_v55
			_1456_v55 = _dafny.SetOf(_1455_v54, _1455_v54, _1455_v54)
			var _source22 D13 = Companion_Default___.Fm47(_1456_v55, _1388_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(407))).Uint32(), func(coer91 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg91 _dafny.Int) interface{} {
					return coer91(arg91)
				}
			}((func(_1457_v3 bool) func(_dafny.Int) _dafny.Sequence {
				return func(_1458_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(!(_1457_v3))
				}
			})(_1388_v3))), _1455_v54, globalState)
			_ = _source22
			if _source22.Is_DC28() {
				var _1459___mcc_h5 _dafny.Sequence = _source22.Get_().(D13_DC28).Cf40
				_ = _1459___mcc_h5
				var _1460___mcc_h6 _dafny.Int = _source22.Get_().(D13_DC28).Cf41
				_ = _1460___mcc_h6
				var _1461___mcc_h7 bool = _source22.Get_().(D13_DC28).Cf42
				_ = _1461___mcc_h7
				var _1462_cf42 bool = _1461___mcc_h7
				_ = _1462_cf42
				var _1463_cf41 _dafny.Int = _1460___mcc_h6
				_ = _1463_cf41
				var _1464_cf40 _dafny.Sequence = _1459___mcc_h5
				_ = _1464_cf40
				var _1465_v56 D8
				_ = _1465_v56
				_1465_v56 = Companion_D8_.Create_DC15_(_1403_v17)
				var _1466_v57 _dafny.MultiSet
				_ = _1466_v57
				_1466_v57 = _dafny.MultiSetOf(_1388_v3, _1462_cf42)
				var _1467_v58 _dafny.Map
				_ = _1467_v58
				_1467_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1464_cf40, (func() _dafny.Int {
					if (_1466_v57).Contains(_1388_v3) {
						return (_1466_v57).Multiplicity(_1388_v3)
					}
					return _1463_cf41
				})())
				var _1468_v60 _dafny.Sequence
				_ = _1468_v60
				_1468_v60 = _dafny.SeqOf(_1464_cf40, _1464_cf40, _1464_cf40, _dafny.UnicodeSeqOfUtf8Bytes("nhsif"), _1464_cf40)
				var _1469_v61 T1
				_ = _1469_v61
				var _nw233 *C4 = New_C4_()
				_ = _nw233
				_nw233.Ctor__(_dafny.MultiSetOf((_this).F11()), _1388_v3, _dafny.SetOf(_1388_v3), (_this).F11())
				_1469_v61 = _nw233
				var _1470_v62 _dafny.Array
				_ = _1470_v62
				var _nwElement0_62 T1 = _1469_v61
				_ = _nwElement0_62
				var _nw234 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(17))
				_ = _nw234
				(_nw234).ArraySet1(_nwElement0_62, 0)
				(_nw234).ArraySet1(_1469_v61, 1)
				(_nw234).ArraySet1(_1469_v61, 2)
				(_nw234).ArraySet1(_1469_v61, 3)
				(_nw234).ArraySet1(_1469_v61, 4)
				(_nw234).ArraySet1(_1469_v61, 5)
				(_nw234).ArraySet1(_1469_v61, 6)
				(_nw234).ArraySet1(_1469_v61, 7)
				(_nw234).ArraySet1(_1469_v61, 8)
				(_nw234).ArraySet1(_1469_v61, 9)
				(_nw234).ArraySet1(_1469_v61, 10)
				(_nw234).ArraySet1(_1469_v61, 11)
				(_nw234).ArraySet1(_1469_v61, 12)
				(_nw234).ArraySet1(_1469_v61, 13)
				(_nw234).ArraySet1(_1469_v61, 14)
				(_nw234).ArraySet1(_1469_v61, 15)
				(_nw234).ArraySet1(_1469_v61, 16)
				_1470_v62 = _nw234
				var _1471_v63 _dafny.MultiSet
				_ = _1471_v63
				_1471_v63 = _dafny.MultiSetOf(_1470_v62, _1470_v62, _1470_v62)
				var _1472_v64 _dafny.Sequence
				_ = _1472_v64
				_1472_v64 = _dafny.SeqOf((_this).F11(), _1463_cf41)
				var _1473_v65 _dafny.Map
				_ = _1473_v65
				_1473_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1466_v57)
				var _1474_v66 _dafny.Array
				_ = _1474_v66
				var _nwElement0_63 _dafny.Int = _1463_cf41
				_ = _nwElement0_63
				var _nw235 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(29))
				_ = _nw235
				(_nw235).ArraySet1(_nwElement0_63, 0)
				(_nw235).ArraySet1((_this).Fm3((_this).F11(), _1388_v3, _1388_v3, globalState), 1)
				(_nw235).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_1463_cf41), (_this).F11()), 2)
				(_nw235).ArraySet1((_this).F11(), 3)
				(_nw235).ArraySet1((_1463_cf41).Times((_this).F11()), 4)
				(_nw235).ArraySet1((_this).Fm6(_1462_cf42, _1463_cf41, globalState), 5)
				(_nw235).ArraySet1(_1463_cf41, 6)
				(_nw235).ArraySet1(_1463_cf41, 7)
				(_nw235).ArraySet1(((_1467_v58).Merge(func() _dafny.Map {
					var _coll64 = _dafny.NewMapBuilder()
					_ = _coll64
					for _iter66 := _dafny.Iterate((_1468_v60).Elements()); ; {
						_compr_64, _ok66 := _iter66()
						if !_ok66 {
							break
						}
						var _1475_v59 _dafny.Sequence
						_1475_v59 = interface{}(_compr_64).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_1468_v60, _1475_v59) {
							_coll64.Add(_1475_v59, _dafny.IntOfInt64(510))
						}
					}
					return _coll64.ToMap()
				}())).Cardinality(), 8)
				(_nw235).ArraySet1((_this).F11(), 9)
				(_nw235).ArraySet1(_1463_cf41, 10)
				(_nw235).ArraySet1((_this).F11(), 11)
				(_nw235).ArraySet1((_this).F11(), 12)
				(_nw235).ArraySet1((_this).F11(), 13)
				(_nw235).ArraySet1((_dafny.Zero).Minus(((_1471_v63).Cardinality()).Minus((_this).F11())), 14)
				(_nw235).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-181), (_1469_v61).Fm3(_dafny.IntOfUint32((_1464_cf40).Cardinality()), _1462_cf42, _1462_cf42, globalState)), 15)
				(_nw235).ArraySet1((_1469_v61).F11(), 16)
				(_nw235).ArraySet1(_dafny.IntOfInt64(438), 17)
				(_nw235).ArraySet1((_this).Fm3((_dafny.Zero).Minus((_1469_v61).F11()), _1388_v3, !(_1388_v3), globalState), 18)
				(_nw235).ArraySet1((_1463_cf41).Times((_this).F11()), 19)
				(_nw235).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1388_v3, _1462_cf42))).Update(_1388_v3, _1403_v17)).Cardinality(), 20)
				(_nw235).ArraySet1((_1403_v17).Cardinality(), 21)
				(_nw235).ArraySet1(_dafny.IntOfUint32((_1472_v64).Cardinality()), 22)
				(_nw235).ArraySet1(_1463_cf41, 23)
				(_nw235).ArraySet1((_this).F11(), 24)
				(_nw235).ArraySet1(((_this).F12()).Cardinality(), 25)
				(_nw235).ArraySet1((_1463_cf41).Plus((_this).F11()), 26)
				(_nw235).ArraySet1((_1473_v65).Cardinality(), 27)
				(_nw235).ArraySet1(_1463_cf41, 28)
				_1474_v66 = _nw235
				var _1476_v67 _dafny.MultiSet
				_ = _1476_v67
				_1476_v67 = _dafny.MultiSetOf((_1463_cf41).Plus((_this).F11()))
				var _1477_v68 _dafny.Map
				_ = _1477_v68
				_1477_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1463_cf41, _1463_cf41)
				var _1478_v69 _dafny.Sequence
				_ = _1478_v69
				_1478_v69 = _dafny.SeqOf(!(_1388_v3), _1388_v3)
				var _rhs256 _dafny.Int = (func() _dafny.Int {
					if (_1476_v67).Contains((_dafny.Zero).Minus((func() _dafny.Int {
						if (_1477_v68).Contains(_dafny.IntOfInt64(-405)) {
							return (_1477_v68).Get(_dafny.IntOfInt64(-405)).(_dafny.Int)
						}
						return (_this).F11()
					})())) {
						return (_1476_v67).Multiplicity((_dafny.Zero).Minus((func() _dafny.Int {
							if (_1477_v68).Contains(_dafny.IntOfInt64(-405)) {
								return (_1477_v68).Get(_dafny.IntOfInt64(-405)).(_dafny.Int)
							}
							return (_this).F11()
						})()))
					}
					return (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(762), _dafny.IntOfUint32((_1478_v69).Cardinality())))
				})()
				_ = _rhs256
				var _rhs257 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1464_cf40, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(443))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg92 _dafny.Int) interface{} {
						return coer92(arg92)
					}
				}((func(_1479_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1480_i5 _dafny.Int) _dafny.CodePoint {
						return _1479_p0
					}
				})(p0))))).Cardinality())
				_ = _rhs257
				var _rhs258 D8 = Companion_Default___.Fm48(_dafny.IntOfInt64(458), globalState)
				_ = _rhs258
				var _rhs259 bool = (_1462_cf42) == (_1462_cf42)
				_ = _rhs259
				var _rhs260 _dafny.Array = _1474_v66
				_ = _rhs260
				var _lhs213 *GlobalState = globalState
				_ = _lhs213
				var _lhs214 *GlobalState = globalState
				_ = _lhs214
				var _lhs215 *GlobalState = globalState
				_ = _lhs215
				_lhs213.F1 = _rhs256
				_lhs214.F1 = _rhs257
				_1465_v56 = _rhs258
				_lhs215.F7 = _rhs259
				_1474_v66 = _rhs260
				_1388_v3 = true
				var _1481_v70 _dafny.Array
				_ = _1481_v70
				var _nw236 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(29))
				_ = _nw236
				_1481_v70 = _nw236
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_1481_v70), 0))
				_ = _index193
				(_1481_v70).ArraySet1CodePoint(p0, (_index193).Int())
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_1481_v70), 0))
				_ = _index194
				(_1481_v70).ArraySet1CodePoint(p0, (_index194).Int())
				var _1482_v71 D6
				_ = _1482_v71
				_1482_v71 = Companion_D6_.Create_DC10_(_1478_v69)
				var _rhs261 bool = _1462_cf42
				_ = _rhs261
				var _rhs262 D6 = _1482_v71
				_ = _rhs262
				var _lhs216 *GlobalState = globalState
				_ = _lhs216
				_lhs216.F6 = _rhs261
				_1482_v71 = _rhs262
			} else if _source22.Is_DC27() {
				var _1483___mcc_h8 _dafny.Map = _source22.Get_().(D13_DC27).Cf39
				_ = _1483___mcc_h8
				var _1484_cf39 _dafny.Map = _1483___mcc_h8
				_ = _1484_cf39
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_1386_v1), 0))
				_ = _index195
				(_1386_v1).ArraySet1((_1388_v3) == (_1388_v3), (_index195).Int())
				var _1485_v72 _dafny.Sequence
				_ = _1485_v72
				_1485_v72 = _dafny.UnicodeSeqOfUtf8Bytes("ctfp")
				var _1486_v73 D8
				_ = _1486_v73
				_1486_v73 = Companion_D8_.Create_DC16_(_1388_v3, _dafny.IntOfUint32((_1485_v72).Cardinality()))
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_1386_v1), 0))
				_ = _index196
				(_1386_v1).ArraySet1((_1486_v73).Dtor_cf23(), (_index196).Int())
				var _1487_v74 _dafny.Sequence
				_ = _1487_v74
				_1487_v74 = _dafny.SeqOf((_this).F11())
				_1487_v74 = _dafny.SeqOf((_this).F11(), (_this).F11(), (_this).F11())
				var _1488_v75 _dafny.Map
				_ = _1488_v75
				_1488_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1388_v3, _dafny.IntOfInt64(879))
				var _1489_v76 _dafny.Sequence
				_ = _1489_v76
				_1489_v76 = _dafny.SeqOf(_1488_v75)
				var _1490_v77 _dafny.Map
				_ = _1490_v77
				_1490_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2((_dafny.Zero).Minus(_dafny.IntOfUint32((_1489_v76).Cardinality())), (_this).F11(), globalState), _1487_v74)
				_1490_v77 = (_1490_v77).Update((func() _dafny.Int {
					if (_1386_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_1386_v1), 0))).Int()).(bool) {
						return (_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm1(globalState)).Cardinality()))
					}
					return _dafny.IntOfInt64(649)
				})(), _1487_v74)
				var _1491_v78 _dafny.Array
				_ = _1491_v78
				var _nw237 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
				_ = _nw237
				_1491_v78 = _nw237
				var _1492_v79 _dafny.Map
				_ = _1492_v79
				_1492_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1491_v78)
				_1492_v79 = (_1492_v79).Update((_this).F11(), _1491_v78)
			} else {
				var _1493___mcc_h9 D13 = _source22.Get_().(D13_DC29).Cf43
				_ = _1493___mcc_h9
				var _1494_cf43 D13 = _1493___mcc_h9
				_ = _1494_cf43
				var _1495_v80 _dafny.Sequence
				_ = _1495_v80
				_1495_v80 = _dafny.SeqOf(_1388_v3)
				var _1496_v81 _dafny.Set
				_ = _1496_v81
				_1496_v81 = _dafny.SetOf(p0, p0, p0, Companion_Default___.Fm20(_1495_v80, (_this).F11(), _1388_v3, globalState))
				_1496_v81 = _1496_v81
				var _1497_v82 _dafny.Sequence
				_ = _1497_v82
				_1497_v82 = _dafny.UnicodeSeqOfUtf8Bytes("l")
				var _1498_v83 _dafny.Sequence
				_ = _1498_v83
				_1498_v83 = _dafny.SeqOf((_this).F11())
				var _1499_v84 _dafny.Array
				_ = _1499_v84
				var _nwElement0_64 _dafny.Int = (_this).F11()
				_ = _nwElement0_64
				var _nw238 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(3))
				_ = _nw238
				(_nw238).ArraySet1(_nwElement0_64, 0)
				(_nw238).ArraySet1((_this).F11(), 1)
				(_nw238).ArraySet1((_this).F11(), 2)
				_1499_v84 = _nw238
				var _1500_v85 _dafny.MultiSet
				_ = _1500_v85
				_1500_v85 = _dafny.MultiSetOf(_1388_v3, _1388_v3)
				var _1501_v86 _dafny.Map
				_ = _1501_v86
				_1501_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_1499_v84, _1499_v84)).Cardinality(), (_1500_v85).Cardinality())
				var _1502_v87 _dafny.Array
				_ = _1502_v87
				var _nwElement0_65 _dafny.Int = ((_this).F11()).Times(_dafny.IntOfUint32((_1497_v82).Cardinality()))
				_ = _nwElement0_65
				var _nw239 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(15))
				_ = _nw239
				(_nw239).ArraySet1(_nwElement0_65, 0)
				(_nw239).ArraySet1((_this).F11(), 1)
				(_nw239).ArraySet1((_1498_v83).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1501_v86).Contains((Companion_D11_.Create_DC23_(Companion_Default___.Fm0(globalState), _dafny.IntOfInt64(19), (_this).F11())).Dtor_cf33()) {
						return (_1501_v86).Get((Companion_D11_.Create_DC23_(Companion_Default___.Fm0(globalState), _dafny.IntOfInt64(19), (_this).F11())).Dtor_cf33()).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hwt")).Cardinality())
				})(), _dafny.IntOfUint32((_1498_v83).Cardinality()))).Uint32()).(_dafny.Int), 2)
				(_nw239).ArraySet1((_this).F11(), 3)
				(_nw239).ArraySet1(Companion_Default___.SafeModuloInt((_this).F11(), (_this).F11()), 4)
				(_nw239).ArraySet1((_this).F11(), 5)
				(_nw239).ArraySet1((_this).F11(), 6)
				(_nw239).ArraySet1((_this).Fm3((_this).F11(), !(_1388_v3), _1388_v3, globalState), 7)
				(_nw239).ArraySet1((_this).F11(), 8)
				(_nw239).ArraySet1((_this).F11(), 9)
				(_nw239).ArraySet1((_this).F11(), 10)
				(_nw239).ArraySet1(((_this).F11()).Plus(_dafny.IntOfUint32((_1495_v80).Cardinality())), 11)
				(_nw239).ArraySet1((_this).F11(), 12)
				(_nw239).ArraySet1((_this).F11(), 13)
				(_nw239).ArraySet1((_this).F11(), 14)
				_1502_v87 = _nw239
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_1499_v84), 0))
				_ = _index197
				(_1499_v84).ArraySet1(_dafny.IntOfInt64(835), (_index197).Int())
				var _1503_v88 D14
				_ = _1503_v88
				_1503_v88 = Companion_D14_.Create_DC31_(_1388_v3, (_this).F11(), (_this).F11())
				var _1504_v89 _dafny.MultiSet
				_ = _1504_v89
				_1504_v89 = _dafny.MultiSetOf(_1497_v82)
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_1499_v84), 0))
				_ = _index198
				(_1499_v84).ArraySet1(Companion_Default___.SafeModuloInt((_1503_v88).Dtor_cf46(), Companion_Default___.SafeModuloInt((_this).F11(), (_1504_v89).Cardinality())), (_index198).Int())
				var _1505_v90 _dafny.Array
				_ = _1505_v90
				var _nw240 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(18))
				_ = _nw240
				_1505_v90 = _nw240
				var _1506_v91 _dafny.Array
				_ = _1506_v91
				_1506_v91 = _1505_v90
				_1505_v90 = (_1506_v91)
				(globalState).F1 = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1495_v80, _dafny.Companion_Sequence_.Concatenate(_1495_v80, _1495_v80)), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-764), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1495_v80, _dafny.Companion_Sequence_.Concatenate(_1495_v80, _1495_v80))).Cardinality()))).Uint32(), _1388_v3), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_1499_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_1499_v84), 0))).Int()).(_dafny.Int), (_1499_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_1499_v84), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1495_v80, _dafny.Companion_Sequence_.Concatenate(_1495_v80, _1495_v80)), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-764), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1495_v80, _dafny.Companion_Sequence_.Concatenate(_1495_v80, _1495_v80))).Cardinality()))).Uint32(), _1388_v3)).Cardinality()))).Uint32(), ((_this).F11()).Cmp((_this).F11()) == 0), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-761), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1495_v80, _dafny.Companion_Sequence_.Concatenate(_1495_v80, _1495_v80)), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-764), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1495_v80, _dafny.Companion_Sequence_.Concatenate(_1495_v80, _1495_v80))).Cardinality()))).Uint32(), _1388_v3), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_1499_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_1499_v84), 0))).Int()).(_dafny.Int), (_1499_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_1499_v84), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1495_v80, _dafny.Companion_Sequence_.Concatenate(_1495_v80, _1495_v80)), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-764), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1495_v80, _dafny.Companion_Sequence_.Concatenate(_1495_v80, _1495_v80))).Cardinality()))).Uint32(), _1388_v3)).Cardinality()))).Uint32(), ((_this).F11()).Cmp((_this).F11()) == 0)).Cardinality()))).Uint32(), _1388_v3))).Cardinality()
			}
			var _1507_v92 _dafny.Array
			_ = _1507_v92
			var _nw241 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw241
			_1507_v92 = _nw241
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1507_v92), 0))
			_ = _index199
			(_1507_v92).ArraySet1((_this).F11(), (_index199).Int())
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1507_v92), 0))
			_ = _index200
			(_1507_v92).ArraySet1((_this).Fm3(((_this).F11()).Minus((_this).F11()), _1388_v3, !(true), globalState), (_index200).Int())
			(globalState).F7 = ((_1507_v92).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1507_v92), 0))).Int()).(_dafny.Int)).Cmp((_this).F11()) == 0
			var _1508_v93 _dafny.Sequence
			_ = _1508_v93
			_1508_v93 = _dafny.SeqOf(_1507_v92)
			var _1509_v94 _dafny.MultiSet
			_ = _1509_v94
			_1509_v94 = _dafny.MultiSetOf(_1388_v3)
			var _1510_v95 _dafny.Map
			_ = _1510_v95
			_1510_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(51), _1388_v3)
			var _1511_v96 _dafny.Array
			_ = _1511_v96
			var _nwElement0_66 _dafny.Array = _1507_v92
			_ = _nwElement0_66
			var _nw242 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(13))
			_ = _nw242
			(_nw242).ArraySet1(_nwElement0_66, 0)
			(_nw242).ArraySet1(_1507_v92, 1)
			(_nw242).ArraySet1(_1507_v92, 2)
			(_nw242).ArraySet1(_1507_v92, 3)
			(_nw242).ArraySet1(_1507_v92, 4)
			(_nw242).ArraySet1(_1507_v92, 5)
			(_nw242).ArraySet1((_1508_v93).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1509_v94).Contains(_1388_v3) {
					return (_1509_v94).Multiplicity(_1388_v3)
				}
				return (_1510_v95).Cardinality()
			})(), _dafny.IntOfUint32((_1508_v93).Cardinality()))).Uint32()).(_dafny.Array), 6)
			(_nw242).ArraySet1(_1507_v92, 7)
			(_nw242).ArraySet1(_1507_v92, 8)
			(_nw242).ArraySet1(_1507_v92, 9)
			(_nw242).ArraySet1(_1507_v92, 10)
			(_nw242).ArraySet1(_1507_v92, 11)
			(_nw242).ArraySet1(_1507_v92, 12)
			_1511_v96 = _nw242
			var _1512_v97 _dafny.Array
			_ = _1512_v97
			var _nwElement0_67 _dafny.Array = _1511_v96
			_ = _nwElement0_67
			var _nw243 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(8))
			_ = _nw243
			(_nw243).ArraySet1(_nwElement0_67, 0)
			(_nw243).ArraySet1(_1511_v96, 1)
			(_nw243).ArraySet1(_1511_v96, 2)
			(_nw243).ArraySet1(_1511_v96, 3)
			(_nw243).ArraySet1((_1511_v96), 4)
			(_nw243).ArraySet1(_1511_v96, 5)
			(_nw243).ArraySet1(_1511_v96, 6)
			(_nw243).ArraySet1(_1511_v96, 7)
			_1512_v97 = _nw243
			var _1513_v98 _dafny.Map
			_ = _1513_v98
			_1513_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1511_v96)
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_1512_v97), 0))
			_ = _index201
			(_1512_v97).ArraySet1((func() _dafny.Array {
				if (_1513_v98).Contains((_1507_v92).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1507_v92), 0))).Int()).(_dafny.Int)) {
					return (_1513_v98).Get((_1507_v92).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1507_v92), 0))).Int()).(_dafny.Int)).(_dafny.Array)
				}
				return _1511_v96
			})(), (_index201).Int())
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_1512_v97), 0))
			_ = _index202
			var _nw244 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
			_ = _nw244
			(_1512_v97).ArraySet1(_nw244, (_index202).Int())
		}
		var _1514_v99 D15
		_ = _1514_v99
		_1514_v99 = Companion_D15_.Create_DC36_((_this).F11(), _1388_v3, _1388_v3, _1388_v3, (_this).F11())
		(globalState).F1 = (_dafny.Zero).Minus((_1514_v99).Dtor_cf55())
		var _1515_v100 _dafny.Sequence
		_ = _1515_v100
		_1515_v100 = _dafny.UnicodeSeqOfUtf8Bytes("dohvyph")
		var _hi1 _dafny.Int = _dafny.IntOfUint32((_1515_v100).Cardinality())
		_ = _hi1
		for _1516_i6 := ((_this).F11()).Times((_this).F11()); _1516_i6.Cmp(_hi1) < 0; _1516_i6 = _1516_i6.Plus(_dafny.One) {
			var _1517_v101 _dafny.Array
			_ = _1517_v101
			var _nw245 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
			_ = _nw245
			_1517_v101 = _nw245
			var _1518_v102 _dafny.Array
			_ = _1518_v102
			_1518_v102 = _1517_v101
			var _1519_v103 _dafny.Array
			_ = _1519_v103
			var _nwElement0_68 _dafny.Array = _1518_v102
			_ = _nwElement0_68
			var _nw246 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(8))
			_ = _nw246
			(_nw246).ArraySet1(_nwElement0_68, 0)
			(_nw246).ArraySet1(_1518_v102, 1)
			(_nw246).ArraySet1(_1517_v101, 2)
			(_nw246).ArraySet1(_1518_v102, 3)
			(_nw246).ArraySet1(_1518_v102, 4)
			(_nw246).ArraySet1(_1517_v101, 5)
			(_nw246).ArraySet1(_1518_v102, 6)
			(_nw246).ArraySet1(_1518_v102, 7)
			_1519_v103 = _nw246
			var _1520_v104 _dafny.Sequence
			_ = _1520_v104
			_1520_v104 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F11()), _1516_i6)
			var _1521_v105 D0
			_ = _1521_v105
			_1521_v105 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_1520_v104).Cardinality()))
			var _1522_v106 _dafny.Sequence
			_ = _1522_v106
			_1522_v106 = _dafny.SeqOf(!(_1388_v3))
			var _1523_v107 _dafny.Sequence
			_ = _1523_v107
			_1523_v107 = _dafny.SeqOf((_this).Fm5(func(_pat_let31_0 D0) D0 {
				return func(_1524_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let32_0 _dafny.Int) D0 {
						return func(_1525_dt__update_hcf0_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC0_(_1525_dt__update_hcf0_h0)
						}(_pat_let32_0)
					}((_this).F11())
				}(_pat_let31_0)
			}(_1521_v105), _1522_v106, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(481))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg93 _dafny.Int) interface{} {
					return coer93(arg93)
				}
			}((func(_1526_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1527_i7 _dafny.Int) _dafny.CodePoint {
					return _1526_p0
				}
			})(p0)))).Cardinality()), p0, globalState))
			var _1528_v108 _dafny.MultiSet
			_ = _1528_v108
			_1528_v108 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1522_v106).Cardinality()))
			var _1529_v109 _dafny.Set
			_ = _1529_v109
			_1529_v109 = _dafny.SetOf((_this).F11(), _1516_i6)
			var _1530_v110 _dafny.Map
			_ = _1530_v110
			_1530_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1388_v3, _dafny.IntOfInt64(716))
			var _1531_v111 _dafny.Sequence
			_ = _1531_v111
			_1531_v111 = _dafny.SeqOf(_1520_v104, _1520_v104)
			var _1532_v112 _dafny.Array
			_ = _1532_v112
			var _nwElement0_69 _dafny.Int = _dafny.IntOfInt64(840)
			_ = _nwElement0_69
			var _nw247 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(21))
			_ = _nw247
			(_nw247).ArraySet1(_nwElement0_69, 0)
			(_nw247).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1516_i6)).Cardinality(), 1)
			(_nw247).ArraySet1(_1516_i6, 2)
			(_nw247).ArraySet1((_this).F11(), 3)
			(_nw247).ArraySet1((_1528_v108).Cardinality(), 4)
			(_nw247).ArraySet1((_1529_v109).Cardinality(), 5)
			(_nw247).ArraySet1((_this).F11(), 6)
			(_nw247).ArraySet1((_this).F11(), 7)
			(_nw247).ArraySet1(_1516_i6, 8)
			(_nw247).ArraySet1((func() _dafny.Int {
				if (_1530_v110).Contains(_1388_v3) {
					return (_1530_v110).Get(_1388_v3).(_dafny.Int)
				}
				return _1516_i6
			})(), 9)
			(_nw247).ArraySet1(_1516_i6, 10)
			(_nw247).ArraySet1((_this).F11(), 11)
			(_nw247).ArraySet1(Companion_Default___.Fm2((_this).F11(), _1516_i6, globalState), 12)
			(_nw247).ArraySet1(_dafny.IntOfUint32((_1523_v107).Cardinality()), 13)
			(_nw247).ArraySet1(_dafny.IntOfUint32((_1531_v111).Cardinality()), 14)
			(_nw247).ArraySet1(_1516_i6, 15)
			(_nw247).ArraySet1(_1516_i6, 16)
			(_nw247).ArraySet1(_1516_i6, 17)
			(_nw247).ArraySet1(_1516_i6, 18)
			(_nw247).ArraySet1(_1516_i6, 19)
			(_nw247).ArraySet1(_1516_i6, 20)
			_1532_v112 = _nw247
			var _1533_v113 _dafny.Map
			_ = _1533_v113
			_1533_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1519_v103, _1532_v112)
			var _1534_v114 D11
			_ = _1534_v114
			_1534_v114 = Companion_D11_.Create_DC22_((_1533_v113).Merge(_1533_v113))
			var _source23 D11 = _1534_v114
			_ = _source23
			if _source23.Is_DC23() {
				var _1535___mcc_h10 bool = _source23.Get_().(D11_DC23).Cf31
				_ = _1535___mcc_h10
				var _1536___mcc_h11 _dafny.Int = _source23.Get_().(D11_DC23).Cf32
				_ = _1536___mcc_h11
				var _1537___mcc_h12 _dafny.Int = _source23.Get_().(D11_DC23).Cf33
				_ = _1537___mcc_h12
				var _1538_cf33 _dafny.Int = _1537___mcc_h12
				_ = _1538_cf33
				var _1539_cf32 _dafny.Int = _1536___mcc_h11
				_ = _1539_cf32
				var _1540_cf31 bool = _1535___mcc_h10
				_ = _1540_cf31
				var _1541_v115 _dafny.MultiSet
				_ = _1541_v115
				_1541_v115 = _dafny.MultiSetOf(_1386_v1, _1386_v1)
				var _1542_v116 _dafny.Map
				_ = _1542_v116
				_1542_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1539_cf32, p0)
				var _rhs263 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1515_v100, _1515_v100), (Companion_Default___.SafeIndex(_1538_cf33, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1515_v100, _1515_v100)).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
					if (_1542_v116).Contains(_1538_cf33) {
						return (_1542_v116).Get(_1538_cf33).(_dafny.CodePoint)
					}
					return p0
				})())).Cardinality())).Minus(_1539_cf32)
				_ = _rhs263
				var _rhs264 _dafny.MultiSet = _1541_v115
				_ = _rhs264
				var _lhs217 *GlobalState = globalState
				_ = _lhs217
				_lhs217.F1 = _rhs263
				_1541_v115 = _rhs264
				_1521_v105 = _1521_v105
				_1528_v108 = (_1528_v108).Union(_1528_v108)
				var _1543_v117 *C2
				_ = _1543_v117
				var _nw248 *C2 = New_C2_()
				_ = _nw248
				_nw248.Ctor__((func() _dafny.CodePoint {
					if _1540_cf31 {
						return p0
					}
					return _dafny.CodePoint('n')
				})(), (_1539_cf32).Times(_1539_cf32))
				_1543_v117 = _nw248
			} else if _source23.Is_DC24() {
				var _1544___mcc_h13 _dafny.Int = _source23.Get_().(D11_DC24).Cf34
				_ = _1544___mcc_h13
				var _1545___mcc_h14 _dafny.Sequence = _source23.Get_().(D11_DC24).Cf35
				_ = _1545___mcc_h14
				var _1546___mcc_h15 _dafny.CodePoint = _source23.Get_().(D11_DC24).Cf36
				_ = _1546___mcc_h15
				var _1547_cf36 _dafny.CodePoint = _1546___mcc_h15
				_ = _1547_cf36
				var _1548_cf35 _dafny.Sequence = _1545___mcc_h14
				_ = _1548_cf35
				var _1549_cf34 _dafny.Int = _1544___mcc_h13
				_ = _1549_cf34
				var _1550_v118 _dafny.Set
				_ = _1550_v118
				_1550_v118 = _dafny.SetOf(p0)
				_1550_v118 = _1550_v118
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_30
				var _nw249 _dafny.Array
				_ = _nw249
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw249 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) _dafny.Int = func(_1551_i8 _dafny.Int) _dafny.Int {
						return (_1551_i8).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("v")).Cardinality()))
					}
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw249 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw249).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw249).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1532_v112 = _nw249
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_1532_v112), 0))
				_ = _index203
				(_1532_v112).ArraySet1(_1516_i6, (_index203).Int())
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.ArrayLen((_1532_v112), 0))
				_ = _index204
				(_1532_v112).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_1516_i6)), (_index204).Int())
				_1549_cf34 = _dafny.IntOfInt64(700)
			} else {
				var _1552___mcc_h16 _dafny.Map = _source23.Get_().(D11_DC22).Cf30
				_ = _1552___mcc_h16
				var _1553_cf30 _dafny.Map = _1552___mcc_h16
				_ = _1553_cf30
				(globalState).F1 = (Companion_Default___.Fm2((_this).F11(), _1516_i6, globalState)).Minus(_1516_i6)
				var _1554_v119 _dafny.Map
				_ = _1554_v119
				_1554_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1516_i6, _1516_i6)
				var _rhs265 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1520_v104, _dafny.SeqOf((_this).F11(), _1516_i6, _dafny.IntOfUint32((_1515_v100).Cardinality())))
				_ = _rhs265
				var _rhs266 bool = (Companion_Default___.SafeDivisionInt(_1516_i6, _1516_i6)).Cmp(_1516_i6) >= 0
				_ = _rhs266
				var _rhs267 _dafny.Array = _1532_v112
				_ = _rhs267
				var _rhs268 _dafny.Int = Companion_Default___.SafeDivisionInt((_1554_v119).Cardinality(), (_this).F11())
				_ = _rhs268
				var _rhs269 _dafny.Array = _1386_v1
				_ = _rhs269
				var _lhs218 *GlobalState = globalState
				_ = _lhs218
				var _lhs219 *GlobalState = globalState
				_ = _lhs219
				_1520_v104 = _rhs265
				_lhs218.F6 = _rhs266
				_1532_v112 = _rhs267
				_lhs219.F1 = _rhs268
				_1386_v1 = _rhs269
				var _1555_v120 D11
				_ = _1555_v120
				_1555_v120 = Companion_D11_.Create_DC24_((_this).F11(), _1523_v107, p0)
				r0 = (_dafny.MultiSetFromSeq((_1555_v120).Dtor_cf35())).Union(_dafny.MultiSetOf(false, true, _1388_v3, _1388_v3))
				var _1556_v122 _dafny.Sequence
				_ = _1556_v122
				_1556_v122 = _dafny.SeqOf(func() _dafny.Map {
					var _coll65 = _dafny.NewMapBuilder()
					_ = _coll65
					for _iter67 := _dafny.Iterate((_1520_v104).Elements()); ; {
						_compr_65, _ok67 := _iter67()
						if !_ok67 {
							break
						}
						var _1557_v121 _dafny.Int
						_1557_v121 = interface{}(_compr_65).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1520_v104, _1557_v121) {
							_coll65.Add((_1557_v121).Minus(_1516_i6), _1388_v3)
						}
					}
					return _coll65.ToMap()
				}())
				var _1558_v123 D0
				_ = _1558_v123
				_1558_v123 = Companion_D0_.Create_DC1_(_1388_v3, ((_1556_v122).Select((Companion_Default___.SafeIndex(_1516_i6, _dafny.IntOfUint32((_1556_v122).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _1388_v3)
				var _1559_v124 _dafny.Array
				_ = _1559_v124
				var _nwElement0_70 D0 = _1558_v123
				_ = _nwElement0_70
				var _nw250 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(10))
				_ = _nw250
				(_nw250).ArraySet1(_nwElement0_70, 0)
				(_nw250).ArraySet1(Companion_Default___.Fm49(_1388_v3, globalState), 1)
				(_nw250).ArraySet1(_1558_v123, 2)
				(_nw250).ArraySet1(_1558_v123, 3)
				(_nw250).ArraySet1(_1558_v123, 4)
				(_nw250).ArraySet1(_1558_v123, 5)
				(_nw250).ArraySet1(_1558_v123, 6)
				(_nw250).ArraySet1(_1558_v123, 7)
				(_nw250).ArraySet1(_1558_v123, 8)
				(_nw250).ArraySet1(_1558_v123, 9)
				_1559_v124 = _nw250
				var _1560_v125 _dafny.MultiSet
				_ = _1560_v125
				_1560_v125 = _dafny.MultiSetOf(_1559_v124)
				var _1561_v126 _dafny.Sequence
				_ = _1561_v126
				_1561_v126 = _dafny.SeqOf(_1559_v124, _1559_v124)
				var _1562_v127 _dafny.Sequence
				_ = _1562_v127
				_1562_v127 = _dafny.SeqOf(_1559_v124, _1559_v124, (_1561_v126).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-797))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg94 _dafny.Int) interface{} {
						return coer94(arg94)
					}
				}((func(_1563_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1564_i9 _dafny.Int) _dafny.CodePoint {
						return _1563_p0
					}
				})(p0)))).Cardinality()), _dafny.IntOfUint32((_1561_v126).Cardinality()))).Uint32()).(_dafny.Array), _1559_v124)
				_1560_v125 = (_dafny.MultiSetOf(_1559_v124, _1559_v124, _1559_v124, (_1562_v127).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_1562_v127).Cardinality()))).Uint32()).(_dafny.Array), _1559_v124)).Difference(_1560_v125)
			}
			(globalState).F1 = _1516_i6
			var _1565_v128 _dafny.Map
			_ = _1565_v128
			_1565_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1388_v3, p0)
			var _1566_v129 D11
			_ = _1566_v129
			_1566_v129 = Companion_D11_.Create_DC23_(_1388_v3, _1516_i6, (_this).F11())
			var _1567_v130 _dafny.Array
			_ = _1567_v130
			var _nwElement0_71 _dafny.CodePoint = p0
			_ = _nwElement0_71
			var _nw251 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(17))
			_ = _nw251
			(_nw251).ArraySet1CodePoint(_nwElement0_71, 0)
			(_nw251).ArraySet1CodePoint((_1515_v100).Select((Companion_Default___.SafeIndex(_1516_i6, _dafny.IntOfUint32((_1515_v100).Cardinality()))).Uint32()).(_dafny.CodePoint), 1)
			(_nw251).ArraySet1CodePoint(_dafny.CodePoint('l'), 2)
			(_nw251).ArraySet1CodePoint(p0, 3)
			(_nw251).ArraySet1CodePoint(_dafny.CodePoint('q'), 4)
			(_nw251).ArraySet1CodePoint(_dafny.CodePoint('e'), 5)
			(_nw251).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_1565_v128).Contains((_1566_v129).Dtor_cf31()) {
					return (_1565_v128).Get((_1566_v129).Dtor_cf31()).(_dafny.CodePoint)
				}
				return p0
			})(), 6)
			(_nw251).ArraySet1CodePoint(p0, 7)
			(_nw251).ArraySet1CodePoint(p0, 8)
			(_nw251).ArraySet1CodePoint(p0, 9)
			(_nw251).ArraySet1CodePoint(p0, 10)
			(_nw251).ArraySet1CodePoint(p0, 11)
			(_nw251).ArraySet1CodePoint(_dafny.CodePoint('v'), 12)
			(_nw251).ArraySet1CodePoint(p0, 13)
			(_nw251).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _1388_v3 {
					return p0
				}
				return p0
			})(), 14)
			(_nw251).ArraySet1CodePoint(_dafny.CodePoint('l'), 15)
			(_nw251).ArraySet1CodePoint(p0, 16)
			_1567_v130 = _nw251
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_1567_v130), 0))
			_ = _index205
			(_1567_v130).ArraySet1CodePoint(p0, (_index205).Int())
			var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(502), _dafny.ArrayLen((_1567_v130), 0))
			_ = _index206
			(_1567_v130).ArraySet1CodePoint(p0, (_index206).Int())
			_1515_v100 = _1515_v100
		}
		var _1568_v131 _dafny.MultiSet
		_ = _1568_v131
		_1568_v131 = _dafny.MultiSetOf(_1388_v3, _1388_v3, true)
		r0 = _1568_v131
		return r0
	}
}
func (_this *C6) M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Map) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		(globalState).F1 = Companion_Default___.SafeDivisionInt((_this).F11(), (_dafny.Zero).Minus(p0))
		var _1569_v0 _dafny.Array
		_ = _1569_v0
		var _nw252 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
		_ = _nw252
		_1569_v0 = _nw252
		var _1570_v1 _dafny.Sequence
		_ = _1570_v1
		_1570_v1 = _dafny.UnicodeSeqOfUtf8Bytes("sva")
		var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))
		_ = _index207
		(_1569_v0).ArraySet1(_1570_v1, (_index207).Int())
		var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))
		_ = _index208
		(_1569_v0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(globalState), _dafny.Companion_Sequence_.Concatenate(_1570_v1, _1570_v1)), (_index208).Int())
		var _1571_v2 bool
		_ = _1571_v2
		_1571_v2 = true
		if !(_1571_v2) {
			var _1572_v3 _dafny.Sequence
			_ = _1572_v3
			_1572_v3 = _dafny.SeqOf(_1571_v2)
			var _1573_v4 _dafny.MultiSet
			_ = _1573_v4
			_1573_v4 = _dafny.MultiSetOf(_1571_v2)
			var _source24 D11 = Companion_Default___.Fm50(_dafny.IntOfUint32((_1572_v3).Cardinality()), _1571_v2, _1571_v2, (_1572_v3).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1573_v4).Contains(_1571_v2) {
					return (_1573_v4).Multiplicity(_1571_v2)
				}
				return p0
			})(), _dafny.IntOfUint32((_1572_v3).Cardinality()))).Uint32()).(bool), globalState)
			_ = _source24
			if _source24.Is_DC23() {
				var _1574___mcc_h0 bool = _source24.Get_().(D11_DC23).Cf31
				_ = _1574___mcc_h0
				var _1575___mcc_h1 _dafny.Int = _source24.Get_().(D11_DC23).Cf32
				_ = _1575___mcc_h1
				var _1576___mcc_h2 _dafny.Int = _source24.Get_().(D11_DC23).Cf33
				_ = _1576___mcc_h2
				var _1577_cf33 _dafny.Int = _1576___mcc_h2
				_ = _1577_cf33
				var _1578_cf32 _dafny.Int = _1575___mcc_h1
				_ = _1578_cf32
				var _1579_cf31 bool = _1574___mcc_h0
				_ = _1579_cf31
				var _1580_v5 _dafny.Array
				_ = _1580_v5
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_31
				var _nw253 _dafny.Array
				_ = _nw253
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw253 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) _dafny.Map = (func(_1581_v2 bool, _1582_cf31 bool) func(_dafny.Int) _dafny.Map {
						return func(_1583_i0 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1581_v2, _1581_v2)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1582_cf31, false))
						}
					})(_1571_v2, _1579_cf31)
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw253 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw253).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw253).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1580_v5 = _nw253
				_1580_v5 = _1580_v5
				(globalState).F7 = _1579_cf31
				var _1584_v6 _dafny.Map
				_ = _1584_v6
				_1584_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v2, _1579_cf31)
				var _1585_v7 _dafny.Sequence
				_ = _1585_v7
				_1585_v7 = _dafny.SeqOf(_1584_v6)
				var _1586_v8 D8
				_ = _1586_v8
				_1586_v8 = Companion_D8_.Create_DC15_((_1585_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.IntOfUint32((_1585_v7).Cardinality()))).Uint32()).(_dafny.Map))
				_1586_v8 = _1586_v8
				var _1587_v9 _dafny.Array
				_ = _1587_v9
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_32
				var _nw254 _dafny.Array
				_ = _nw254
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw254 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) _dafny.Sequence = (func(_1588_v3 _dafny.Sequence, _1589_v1 _dafny.Sequence, _1590_cf31 bool) func(_dafny.Int) _dafny.Sequence {
						return func(_1591_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Update(_1588_v3, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1589_v1).Cardinality()), _dafny.IntOfUint32((_1588_v3).Cardinality()))).Uint32(), !(_1590_cf31))
						}
					})(_1572_v3, _1570_v1, _1579_cf31)
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw254 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw254).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw254).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1587_v9 = _nw254
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_1587_v9), 0))
				_ = _index209
				(_1587_v9).ArraySet1(_1572_v3, (_index209).Int())
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_1587_v9), 0))
				_ = _index210
				var _rhs270 _dafny.Sequence = Companion_Default___.Fm19(_1579_cf31, !(_1579_cf31), globalState)
				_ = _rhs270
				var _rhs271 bool = _1571_v2
				_ = _rhs271
				var _lhs220 _dafny.Array = _1587_v9
				_ = _lhs220
				var _lhs221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_1587_v9), 0))
				_ = _lhs221
				(_lhs220).ArraySet1(_rhs270, (_lhs221).Int())
				_1571_v2 = _rhs271
			} else if _source24.Is_DC24() {
				var _1592___mcc_h3 _dafny.Int = _source24.Get_().(D11_DC24).Cf34
				_ = _1592___mcc_h3
				var _1593___mcc_h4 _dafny.Sequence = _source24.Get_().(D11_DC24).Cf35
				_ = _1593___mcc_h4
				var _1594___mcc_h5 _dafny.CodePoint = _source24.Get_().(D11_DC24).Cf36
				_ = _1594___mcc_h5
				var _1595_cf36 _dafny.CodePoint = _1594___mcc_h5
				_ = _1595_cf36
				var _1596_cf35 _dafny.Sequence = _1593___mcc_h4
				_ = _1596_cf35
				var _1597_cf34 _dafny.Int = _1592___mcc_h3
				_ = _1597_cf34
				var _1598_v10 _dafny.Array
				_ = _1598_v10
				var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
				_ = _nw255
				_1598_v10 = _nw255
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_1598_v10), 0))
				_ = _index211
				(_1598_v10).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F11()), p0), (_index211).Int())
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_1598_v10), 0))
				_ = _index212
				(_1598_v10).ArraySet1((_this).F11(), (_index212).Int())
				_1598_v10 = _1598_v10
				var _1599_v11 _dafny.Set
				_ = _1599_v11
				_1599_v11 = _dafny.SetOf((_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence), _1570_v1)
				_1599_v11 = _1599_v11
				var _1600_v12 _dafny.Map
				_ = _1600_v12
				_1600_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _dafny.Companion_Sequence_.Concatenate(_1570_v1, _dafny.Companion_Sequence_.Update(_1570_v1, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1570_v1).Cardinality()))).Uint32(), _1595_cf36)))
				_1600_v12 = (_1600_v12).Update((_this).F11(), (_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence))
			} else {
				var _1601___mcc_h6 _dafny.Map = _source24.Get_().(D11_DC22).Cf30
				_ = _1601___mcc_h6
				var _1602_cf30 _dafny.Map = _1601___mcc_h6
				_ = _1602_cf30
				var _1603_v13 _dafny.Array
				_ = _1603_v13
				var _nwElement0_72 _dafny.Int = p0
				_ = _nwElement0_72
				var _nw256 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(6))
				_ = _nw256
				(_nw256).ArraySet1(_nwElement0_72, 0)
				(_nw256).ArraySet1((p0).Times(p0), 1)
				(_nw256).ArraySet1(p0, 2)
				(_nw256).ArraySet1(Companion_Default___.Fm2((_this).F11(), (_dafny.MultiSetFromSeq(_1572_v3)).Cardinality(), globalState), 3)
				(_nw256).ArraySet1(((_this).F12()).Cardinality(), 4)
				(_nw256).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F11(), (_dafny.Zero).Minus((_this).F11())), 5)
				_1603_v13 = _nw256
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1603_v13), 0))
				_ = _index213
				(_1603_v13).ArraySet1((_this).F11(), (_index213).Int())
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1603_v13), 0))
				_ = _index214
				(_1603_v13).ArraySet1((_this).F11(), (_index214).Int())
				var _1604_v14 _dafny.Array
				_ = _1604_v14
				var _nw257 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
				_ = _nw257
				_1604_v14 = _nw257
				var _1605_v15 _dafny.Map
				_ = _1605_v15
				_1605_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1604_v14, _dafny.IntOfUint32(((_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()))
				var _1606_v16 _dafny.MultiSet
				_ = _1606_v16
				_1606_v16 = _dafny.MultiSetOf((_this).F11(), (_this).F11(), (_this).F11(), (_dafny.Zero).Minus(p0))
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1603_v13), 0))
				_ = _index215
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1603_v13), 0))
				_ = _index216
				var _rhs272 _dafny.Int = (func() _dafny.Int {
					if (_1605_v15).Contains(_1604_v14) {
						return (_1605_v15).Get(_1604_v14).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1572_v3, _dafny.SeqOf(_1571_v2))).Cardinality())
				})()
				_ = _rhs272
				var _rhs273 bool = false
				_ = _rhs273
				var _rhs274 _dafny.Int = (_dafny.Zero).Minus(p0)
				_ = _rhs274
				var _rhs275 bool = (_1606_v16).IsSubsetOf((_1606_v16).Update((_this).F11(), Companion_Default___.Abs((_this).F11())))
				_ = _rhs275
				var _lhs222 _dafny.Array = _1603_v13
				_ = _lhs222
				var _lhs223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1603_v13), 0))
				_ = _lhs223
				var _lhs224 *GlobalState = globalState
				_ = _lhs224
				var _lhs225 _dafny.Array = _1603_v13
				_ = _lhs225
				var _lhs226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1603_v13), 0))
				_ = _lhs226
				var _lhs227 *GlobalState = globalState
				_ = _lhs227
				(_lhs222).ArraySet1(_rhs272, (_lhs223).Int())
				_lhs224.F7 = _rhs273
				(_lhs225).ArraySet1(_rhs274, (_lhs226).Int())
				_lhs227.F7 = _rhs275
				var _1607_v17 _dafny.Map
				_ = _1607_v17
				_1607_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v2, p0)
				var _1608_v18 _dafny.Set
				_ = _1608_v18
				_1608_v18 = _dafny.SetOf(_1607_v17)
				_1608_v18 = _dafny.SetOf(_1607_v17, _1607_v17)
				var _1609_v19 _dafny.Map
				_ = _1609_v19
				_1609_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1572_v3, _1572_v3)).Cardinality()), _1572_v3)
				var _1610_v20 *C1
				_ = _1610_v20
				var _nw258 *C1 = New_C1_()
				_ = _nw258
				_nw258.Ctor__(_1571_v2, _1571_v2)
				_1610_v20 = _nw258
				var _1611_v21 _dafny.Map
				_ = _1611_v21
				_1611_v21 = _1609_v19
				var _1612_v22 D13
				_ = _1612_v22
				_1612_v22 = Companion_D13_.Create_DC28_(_dafny.UnicodeSeqOfUtf8Bytes("bqjjbmfo"), (_1603_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1603_v13), 0))).Int()).(_dafny.Int), (_1610_v20).F18())
				var _1613_v23 _dafny.Map
				_ = _1613_v23
				_1613_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_1610_v20).F18() {
						return _1571_v2
					}
					return (_1610_v20).F18()
				})(), _1610_v20)
				var _rhs276 _dafny.Map = (func() _dafny.Map {
					if _1571_v2 {
						return (_1611_v21)
					}
					return (_1609_v19).Update((_1612_v22).Dtor_cf41(), _1572_v3)
				})()
				_ = _rhs276
				var _rhs277 *C1 = (func() *C1 {
					if (_1613_v23).Contains((_1610_v20).F19()) {
						return (_1613_v23).Get((_1610_v20).F19()).(*C1)
					}
					return _1610_v20
				})()
				_ = _rhs277
				_1609_v19 = _rhs276
				_1610_v20 = _rhs277
				var _1614_v24 _dafny.Sequence
				_ = _1614_v24
				_1614_v24 = _dafny.SeqOf((_1603_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1603_v13), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-974))
				var _1615_v25 _dafny.Map
				_ = _1615_v25
				_1615_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1603_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1603_v13), 0))).Int()).(_dafny.Int), (_dafny.IntOfUint32((_1614_v24).Cardinality())).Times(p0))
				_1615_v25 = (_1615_v25).Update(_dafny.IntOfInt64(799), (_1603_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1603_v13), 0))).Int()).(_dafny.Int))
			}
			var _1616_v26 _dafny.Map
			_ = _1616_v26
			_1616_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32(((_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Cardinality(), p0)
			var _1617_v27 D16
			_ = _1617_v27
			_1617_v27 = Companion_D16_.Create_DC38_(_1616_v26)
			var _1618_v28 _dafny.Map
			_ = _1618_v28
			_1618_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1617_v27, _1571_v2)
			var _pat_let_tv43 = _1616_v26
			_ = _pat_let_tv43
			var _1619_v29 _dafny.Map
			_ = _1619_v29
			_1619_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), (_1618_v28).Update(func(_pat_let33_0 D16) D16 {
				return func(_1620_dt__update__tmp_h0 D16) D16 {
					return func(_pat_let34_0 _dafny.Map) D16 {
						return func(_1621_dt__update_hcf61_h0 _dafny.Map) D16 {
							return Companion_D16_.Create_DC38_(_1621_dt__update_hcf61_h0)
						}(_pat_let34_0)
					}(_pat_let_tv43)
				}(_pat_let33_0)
			}(Companion_D16_.Create_DC38_(Companion_Default___.Fm46(globalState))), false)))
			var _1622_v30 _dafny.Map
			_ = _1622_v30
			_1622_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _1618_v28)
			_1619_v29 = (_1619_v29).Update(_1571_v2, _1622_v30)
			var _1623_v31 _dafny.Array
			_ = _1623_v31
			var _nw259 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
			_ = _nw259
			_1623_v31 = _nw259
			var _1624_v32 _dafny.Map
			_ = _1624_v32
			_1624_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1623_v31, (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F11())))
			var _1625_v33 _dafny.Array
			_ = _1625_v33
			var _nw260 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
			_ = _nw260
			_1625_v33 = _nw260
			var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_1625_v33), 0))
			_ = _index217
			(_1625_v33).ArraySet1((_this).F11(), (_index217).Int())
			var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_1625_v33), 0))
			_ = _index218
			var _rhs278 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1623_v31, (p0).Plus(p0))
			_ = _rhs278
			var _rhs279 bool = !(_1571_v2)
			_ = _rhs279
			var _rhs280 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F11(), (_this).F11())
			_ = _rhs280
			var _rhs281 _dafny.Int = (func() _dafny.Int {
				if _1571_v2 {
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((_1572_v3).Cardinality()))
				}
				return _dafny.IntOfInt64(183)
			})()
			_ = _rhs281
			var _lhs228 *GlobalState = globalState
			_ = _lhs228
			var _lhs229 *GlobalState = globalState
			_ = _lhs229
			var _lhs230 _dafny.Array = _1625_v33
			_ = _lhs230
			var _lhs231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_1625_v33), 0))
			_ = _lhs231
			_1624_v32 = _rhs278
			_lhs228.F7 = _rhs279
			_lhs229.F1 = _rhs280
			(_lhs230).ArraySet1(_rhs281, (_lhs231).Int())
			var _1626_v34 _dafny.Sequence
			_ = _1626_v34
			_1626_v34 = _dafny.SeqOf((_this).F11())
			(globalState).F1 = (_this).Fm3(_dafny.IntOfUint32((_1626_v34).Cardinality()), _1571_v2, true, globalState)
			(globalState).F7 = _1571_v2
		} else {
			var _1627_v35 _dafny.Array
			_ = _1627_v35
			var _nw261 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw261
			_1627_v35 = _nw261
			var _1628_v36 _dafny.Sequence
			_ = _1628_v36
			_1628_v36 = _dafny.SeqOf((_this).F11(), (_this).F11())
			var _1629_v37 D14
			_ = _1629_v37
			_1629_v37 = Companion_D14_.Create_DC31_(_1571_v2, (_this).F11(), (_this).F11())
			var _1630_v38 D14
			_ = _1630_v38
			_1630_v38 = Companion_D14_.Create_DC32_(_1629_v37)
			var _1631_v39 _dafny.Sequence
			_ = _1631_v39
			_1631_v39 = _dafny.SeqOf(_1630_v38)
			var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1627_v35), 0))
			_ = _index219
			(_1627_v35).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm51(_1628_v36, p0, p0, p0, globalState), _1631_v39)).Cardinality()), (_index219).Int())
			var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1627_v35), 0))
			_ = _index220
			(_1627_v35).ArraySet1((_dafny.Zero).Minus((_this).F11()), (_index220).Int())
			if _1571_v2 {
				var _1632_v40 _dafny.Array
				_ = _1632_v40
				var _nw262 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
				_ = _nw262
				_1632_v40 = _nw262
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_1632_v40), 0))
				_ = _index221
				(_1632_v40).ArraySet1(_1571_v2, (_index221).Int())
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_1632_v40), 0))
				_ = _index222
				(_1632_v40).ArraySet1(_1571_v2, (_index222).Int())
				_1628_v36 = _1628_v36
				(globalState).F1 = (_dafny.Zero).Minus((_this).F11())
				var _1633_v41 _dafny.Array
				_ = _1633_v41
				var _nw263 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw263
				_1633_v41 = _nw263
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_1633_v41), 0))
				_ = _index223
				(_1633_v41).ArraySet1(_1632_v40, (_index223).Int())
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_1633_v41), 0))
				_ = _index224
				(_1633_v41).ArraySet1(_1632_v40, (_index224).Int())
				var _1634_v42 _dafny.Sequence
				_ = _1634_v42
				_1634_v42 = _dafny.SeqOf((_1632_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_1632_v40), 0))).Int()).(bool))
				var _1635_v43 *C0
				_ = _1635_v43
				var _nw264 *C0 = New_C0_()
				_ = _nw264
				_nw264.Ctor__(_1634_v42, _dafny.IntOfInt64(225))
				_1635_v43 = _nw264
			} else {
				_1627_v35 = _1627_v35
				(globalState).F1 = (_this).F11()
				var _1636_v44 D19
				_ = _1636_v44
				_1636_v44 = Companion_D19_.Create_DC42_(_1627_v35)
				var _1637_v45 _dafny.Array
				_ = _1637_v45
				var _nwElement0_73 _dafny.Array = _1627_v35
				_ = _nwElement0_73
				var _nw265 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(20))
				_ = _nw265
				(_nw265).ArraySet1(_nwElement0_73, 0)
				(_nw265).ArraySet1(_1627_v35, 1)
				(_nw265).ArraySet1(_1627_v35, 2)
				(_nw265).ArraySet1(_1627_v35, 3)
				(_nw265).ArraySet1(_1627_v35, 4)
				(_nw265).ArraySet1(_1627_v35, 5)
				(_nw265).ArraySet1(_1627_v35, 6)
				(_nw265).ArraySet1(_1627_v35, 7)
				(_nw265).ArraySet1(_1627_v35, 8)
				(_nw265).ArraySet1(_1627_v35, 9)
				(_nw265).ArraySet1(_1627_v35, 10)
				(_nw265).ArraySet1(_1627_v35, 11)
				(_nw265).ArraySet1(_1627_v35, 12)
				(_nw265).ArraySet1(_1627_v35, 13)
				(_nw265).ArraySet1(_1627_v35, 14)
				(_nw265).ArraySet1(_1627_v35, 15)
				(_nw265).ArraySet1(_1627_v35, 16)
				(_nw265).ArraySet1(_1627_v35, 17)
				(_nw265).ArraySet1(_1627_v35, 18)
				(_nw265).ArraySet1((_1636_v44).Dtor_cf68(), 19)
				_1637_v45 = _nw265
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_1637_v45), 0))
				_ = _index225
				(_1637_v45).ArraySet1(_1627_v35, (_index225).Int())
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_1637_v45), 0))
				_ = _index226
				(_1637_v45).ArraySet1(_1627_v35, (_index226).Int())
				var _1638_v46 _dafny.CodePoint
				_ = _1638_v46
				_1638_v46 = _dafny.CodePoint('l')
				var _1639_v47 _dafny.Sequence
				_ = _1639_v47
				_1639_v47 = _dafny.SeqOf(_1571_v2, _1571_v2, _1571_v2)
				var _rhs282 bool = (func() bool {
					if _1571_v2 {
						return _1571_v2
					}
					return (_1639_v47).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_1639_v47).Cardinality()))).Uint32()).(bool)
				})()
				_ = _rhs282
				var _rhs283 _dafny.CodePoint = _1638_v46
				_ = _rhs283
				_1571_v2 = _rhs282
				_1638_v46 = _rhs283
				(globalState).F1 = p0
			}
			(globalState).F1 = (_this).F11()
			(globalState).F7 = (_1571_v2) || (_1571_v2)
			(globalState).F7 = Companion_Default___.Fm0(globalState)
		}
		if (func() bool {
			if _1571_v2 {
				return !(_1571_v2)
			}
			return (p0).Cmp(_dafny.IntOfInt64(522)) != 0
		})() {
			var _1640_v48 D15
			_ = _1640_v48
			_1640_v48 = Companion_D15_.Create_DC34_()
			_1640_v48 = (func() D15 {
				if _1571_v2 {
					return _1640_v48
				}
				return (func() D15 {
					if _1571_v2 {
						return _1640_v48
					}
					return _1640_v48
				})()
			})()
			if _1571_v2 {
				(globalState).F6 = (((_this).F12()).Difference((_this).F12())).IsSubsetOf((_this).F12())
				_1570_v1 = _dafny.UnicodeSeqOfUtf8Bytes("xcrg")
				var _1641_v49 _dafny.CodePoint
				_ = _1641_v49
				_1641_v49 = _dafny.CodePoint('h')
				_1641_v49 = (_1570_v1).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1570_v1).Cardinality()))).Uint32()).(_dafny.CodePoint)
				var _1642_v50 D13
				_ = _1642_v50
				_1642_v50 = Companion_D13_.Create_DC28_(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("qaqirwuan"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qaqirwuan")).Cardinality()))).Uint32(), _1641_v49), (_this).F11(), Companion_Default___.Fm0(globalState))
				(globalState).F6 = (_1642_v50).Dtor_cf42()
				var _1643_v51 _dafny.Array
				_ = _1643_v51
				var _nwElement0_74 _dafny.Int = (_this).F11()
				_ = _nwElement0_74
				var _nw266 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(3))
				_ = _nw266
				(_nw266).ArraySet1(_nwElement0_74, 0)
				(_nw266).ArraySet1(p0, 1)
				(_nw266).ArraySet1(p0, 2)
				_1643_v51 = _nw266
				var _1644_v52 _dafny.Sequence
				_ = _1644_v52
				_1644_v52 = _dafny.SeqOf(_1643_v51, _1643_v51)
				var _1645_v53 _dafny.MultiSet
				_ = _1645_v53
				_1645_v53 = _dafny.MultiSetOf(_dafny.IntOfInt64(591))
				var _1646_v54 _dafny.Map
				_ = _1646_v54
				_1646_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v2, _1571_v2)
				_1644_v52 = _dafny.Companion_Sequence_.Update(_1644_v52, (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1645_v53).Contains((_1646_v54).Cardinality()) {
						return (_1645_v53).Multiplicity((_1646_v54).Cardinality())
					}
					return (_this).F11()
				})(), _dafny.IntOfUint32((_1644_v52).Cardinality()))).Uint32(), _1643_v51)
			} else {
				var _1647_v55 _dafny.Array
				_ = _1647_v55
				var _nw267 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
				_ = _nw267
				_1647_v55 = _nw267
				var _1648_v56 _dafny.Sequence
				_ = _1648_v56
				_1648_v56 = _dafny.SeqOf(Companion_Default___.Fm52((_dafny.Zero).Minus((_this).F11()), (_this).F11(), globalState))
				var _1649_v57 _dafny.Map
				_ = _1649_v57
				_1649_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1647_v55, _dafny.Companion_Sequence_.Concatenate(_1648_v56, _1648_v56))
				_1649_v57 = (_1649_v57).Update(_1647_v55, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-320))).Uint32(), func(coer95 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg95 _dafny.Int) interface{} {
						return coer95(arg95)
					}
				}((func(_1650_v2 bool) func(_dafny.Int) _dafny.Map {
					return func(_1651_i2 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_1650_v2, _dafny.IntOfUint32((_dafny.SeqOf(_1650_v2, _1650_v2)).Cardinality()), !(_1650_v2)), _1650_v2)
					}
				})(_1571_v2))), _1648_v56))
				(globalState).F1 = ((_this).F11()).Minus(p0)
				var _1652_v58 _dafny.Sequence
				_ = _1652_v58
				_1652_v58 = _dafny.SeqOf(p0)
				(globalState).F6 = !((_dafny.IntOfUint32((_1652_v58).Cardinality())).Cmp(((_this).F11()).Times(p0)) >= 0)
				(globalState).F1 = (p0).Minus((_this).F11())
				(globalState).F1 = Companion_Default___.SafeModuloInt(p0, (_this).F11())
			}
			var _1653_v59 _dafny.Sequence
			_ = _1653_v59
			_1653_v59 = _dafny.SeqOf(p0)
			var _1654_v60 _dafny.Set
			_ = _1654_v60
			_1654_v60 = _dafny.SetOf(_1653_v59, _dafny.Companion_Sequence_.Concatenate(_1653_v59, _1653_v59))
			_1654_v60 = (_1654_v60).Difference(_1654_v60)
			var _1655_v61 D0
			_ = _1655_v61
			_1655_v61 = Companion_D0_.Create_DC0_((_this).F11())
			var _1656_v62 _dafny.Sequence
			_ = _1656_v62
			_1656_v62 = _dafny.SeqOf(false)
			var _1657_v63 _dafny.MultiSet
			_ = _1657_v63
			_1657_v63 = _dafny.MultiSetOf(p0)
			var _1658_v64 _dafny.CodePoint
			_ = _1658_v64
			_1658_v64 = _dafny.CodePoint('r')
			if (_this).Fm5(_1655_v61, _1656_v62, (_1657_v63).Cardinality(), _1658_v64, globalState) {
				(globalState).F7 = _1571_v2
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))
				_ = _index227
				(_1569_v0).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence), _dafny.Companion_Sequence_.Update((_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1653_v59).Cardinality()), _dafny.IntOfUint32(((_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _1658_v64)), (_index227).Int())
				var _1659_v65 _dafny.Array
				_ = _1659_v65
				var _nw268 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(29))
				_ = _nw268
				_1659_v65 = _nw268
				var _1660_v66 *C4
				_ = _1660_v66
				var _nw269 *C4 = New_C4_()
				_ = _nw269
				_nw269.Ctor__(_1657_v63, _1571_v2, (_this).F12(), _dafny.IntOfInt64(-299))
				_1660_v66 = _nw269
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1659_v65), 0))
				_ = _index228
				(_1659_v65).ArraySet1((func() *C4 {
					if _1571_v2 {
						return _1660_v66
					}
					return _1660_v66
				})(), (_index228).Int())
				var _1661_v67 _dafny.MultiSet
				_ = _1661_v67
				_1661_v67 = _dafny.MultiSetOf((_1660_v66).F17(), true, _1571_v2, (_1660_v66).F17(), true)
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1659_v65), 0))
				_ = _index229
				var _nw270 *C4 = New_C4_()
				_ = _nw270
				_nw270.Ctor__((func() _dafny.MultiSet {
					if false {
						return _1657_v63
					}
					return _dafny.MultiSetOf((_1653_v59).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_this).F11())).Cardinality()), _dafny.IntOfUint32((_1653_v59).Cardinality()))).Uint32()).(_dafny.Int))
				})(), _1571_v2, (_this).F12(), (_dafny.Zero).Minus((p0).Times((func() _dafny.Int {
					if (_1661_v67).Contains(true) {
						return (_1661_v67).Multiplicity(true)
					}
					return p0
				})())))
				(_1659_v65).ArraySet1(_nw270, (_index229).Int())
				var _1662_v68 _dafny.Array
				_ = _1662_v68
				var _nw271 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw271
				_1662_v68 = _nw271
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_1662_v68), 0))
				_ = _index230
				(_1662_v68).ArraySet1((_this).F11(), (_index230).Int())
				var _1663_v69 _dafny.Array
				_ = _1663_v69
				var _len0_33 _dafny.Int = _dafny.One
				_ = _len0_33
				var _nw272 _dafny.Array
				_ = _nw272
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw272 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) bool = (func(_1664_v2 bool) func(_dafny.Int) bool {
						return func(_1665_i3 _dafny.Int) bool {
							return _1664_v2
						}
					})(_1571_v2)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw272 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw272).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw272).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1663_v69 = _nw272
				var _1666_v70 _dafny.Array
				_ = _1666_v70
				var _nwElement0_75 _dafny.Array = _1663_v69
				_ = _nwElement0_75
				var _nw273 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(3))
				_ = _nw273
				(_nw273).ArraySet1(_nwElement0_75, 0)
				(_nw273).ArraySet1(_1663_v69, 1)
				(_nw273).ArraySet1(_1663_v69, 2)
				_1666_v70 = _nw273
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_1666_v70), 0))
				_ = _index231
				(_1666_v70).ArraySet1(_1663_v69, (_index231).Int())
				var _1667_v71 _dafny.Map
				_ = _1667_v71
				_1667_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1656_v62, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1656_v62).Cardinality()))).Uint32(), (_1660_v66).F17())).Cardinality()), p0)
				var _1668_v72 _dafny.Sequence
				_ = _1668_v72
				_1668_v72 = _dafny.SeqOf(_1667_v71, _1667_v71)
				var _1669_v73 _dafny.Map
				_ = _1669_v73
				_1669_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v2, ((_1668_v72).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1668_v72).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())
				var _1670_v74 _dafny.Map
				_ = _1670_v74
				_1670_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1660_v66).F17())
				var _1671_v75 _dafny.Map
				_ = _1671_v75
				_1671_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1571_v2), _1571_v2)
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_1662_v68), 0))
				_ = _index232
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_1666_v70), 0))
				_ = _index233
				var _rhs284 _dafny.Int = (func() _dafny.Int {
					if (_1669_v73).Contains(!((_1660_v66).F17()) || ((func() bool {
						if (_1670_v74).Contains((_1671_v75).Cardinality()) {
							return (_1670_v74).Get((_1671_v75).Cardinality()).(bool)
						}
						return _1571_v2
					})())) {
						return (_1669_v73).Get(!((_1660_v66).F17()) || ((func() bool {
							if (_1670_v74).Contains((_1671_v75).Cardinality()) {
								return (_1670_v74).Get((_1671_v75).Cardinality()).(bool)
							}
							return _1571_v2
						})())).(_dafny.Int)
					}
					return (_this).F11()
				})()
				_ = _rhs284
				var _rhs285 _dafny.Array = _1663_v69
				_ = _rhs285
				var _rhs286 bool = ((_1660_v66).F17()) || (Companion_Default___.Fm0(globalState))
				_ = _rhs286
				var _rhs287 bool = (_1656_v62).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1656_v62).Cardinality()))).Uint32()).(bool)
				_ = _rhs287
				var _rhs288 bool = (_1660_v66).F17()
				_ = _rhs288
				var _lhs232 _dafny.Array = _1662_v68
				_ = _lhs232
				var _lhs233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_1662_v68), 0))
				_ = _lhs233
				var _lhs234 _dafny.Array = _1666_v70
				_ = _lhs234
				var _lhs235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_1666_v70), 0))
				_ = _lhs235
				var _lhs236 *GlobalState = globalState
				_ = _lhs236
				var _lhs237 *GlobalState = globalState
				_ = _lhs237
				var _lhs238 *GlobalState = globalState
				_ = _lhs238
				(_lhs232).ArraySet1(_rhs284, (_lhs233).Int())
				(_lhs234).ArraySet1(_rhs285, (_lhs235).Int())
				_lhs236.F6 = _rhs286
				_lhs237.F7 = _rhs287
				_lhs238.F7 = _rhs288
				var _1672_v76 _dafny.Array
				_ = _1672_v76
				var _nwElement0_76 _dafny.Array = _1662_v68
				_ = _nwElement0_76
				var _nw274 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(18))
				_ = _nw274
				(_nw274).ArraySet1(_nwElement0_76, 0)
				(_nw274).ArraySet1(_1662_v68, 1)
				(_nw274).ArraySet1(_1662_v68, 2)
				(_nw274).ArraySet1(_1662_v68, 3)
				(_nw274).ArraySet1(_1662_v68, 4)
				(_nw274).ArraySet1(_1662_v68, 5)
				(_nw274).ArraySet1(_1662_v68, 6)
				(_nw274).ArraySet1(_1662_v68, 7)
				(_nw274).ArraySet1(_1662_v68, 8)
				(_nw274).ArraySet1(_1662_v68, 9)
				(_nw274).ArraySet1(_1662_v68, 10)
				(_nw274).ArraySet1(_1662_v68, 11)
				(_nw274).ArraySet1(_1662_v68, 12)
				(_nw274).ArraySet1(_1662_v68, 13)
				(_nw274).ArraySet1(_1662_v68, 14)
				(_nw274).ArraySet1(_1662_v68, 15)
				(_nw274).ArraySet1(_1662_v68, 16)
				(_nw274).ArraySet1(_1662_v68, 17)
				_1672_v76 = _nw274
				var _1673_v77 _dafny.Array
				_ = _1673_v77
				_1673_v77 = _1672_v76
				var _1674_v78 _dafny.Array
				_ = _1674_v78
				var _nwElement0_77 _dafny.Array = _1673_v77
				_ = _nwElement0_77
				var _nw275 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(16))
				_ = _nw275
				(_nw275).ArraySet1(_nwElement0_77, 0)
				(_nw275).ArraySet1(_1673_v77, 1)
				(_nw275).ArraySet1(_1673_v77, 2)
				(_nw275).ArraySet1(_1673_v77, 3)
				(_nw275).ArraySet1(_1673_v77, 4)
				(_nw275).ArraySet1(_1673_v77, 5)
				(_nw275).ArraySet1(_1672_v76, 6)
				(_nw275).ArraySet1(_1673_v77, 7)
				(_nw275).ArraySet1(_1673_v77, 8)
				(_nw275).ArraySet1(_1673_v77, 9)
				(_nw275).ArraySet1(_1673_v77, 10)
				(_nw275).ArraySet1(_1673_v77, 11)
				(_nw275).ArraySet1(_1673_v77, 12)
				(_nw275).ArraySet1(_1673_v77, 13)
				(_nw275).ArraySet1(_1673_v77, 14)
				(_nw275).ArraySet1(_1673_v77, 15)
				_1674_v78 = _nw275
				var _1675_v79 _dafny.Map
				_ = _1675_v79
				_1675_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1674_v78, _1662_v68)
				var _1676_v80 D11
				_ = _1676_v80
				_1676_v80 = Companion_D11_.Create_DC22_(_1675_v79)
				var _pat_let_tv44 = _1675_v79
				_ = _pat_let_tv44
				var _1677_v81 D11
				_ = _1677_v81
				_1677_v81 = Companion_D11_.Create_DC22_((func(_pat_let35_0 D11) D11 {
					return func(_1678_dt__update__tmp_h1 D11) D11 {
						return func(_pat_let36_0 _dafny.Map) D11 {
							return func(_1679_dt__update_hcf30_h0 _dafny.Map) D11 {
								return Companion_D11_.Create_DC22_(_1679_dt__update_hcf30_h0)
							}(_pat_let36_0)
						}(_pat_let_tv44)
					}(_pat_let35_0)
				}(_1676_v80)).Dtor_cf30())
				var _arr2 _dafny.Array = _dafny.ArrayCastTo((_1666_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_1666_v70), 0))).Int()))
				_ = _arr2
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_dafny.ArrayCastTo((_1666_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_1666_v70), 0))).Int()))), 0))
				_ = _index234
				(_arr2).ArraySet1(!((_1660_v66).F17()) || ((_1660_v66).F17()), (_index234).Int())
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_1662_v68), 0))
				_ = _index235
				var _arr3 _dafny.Array = _dafny.ArrayCastTo((_1666_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_1666_v70), 0))).Int()))
				_ = _arr3
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_dafny.ArrayCastTo((_1666_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_1666_v70), 0))).Int()))), 0))
				_ = _index236
				var _rhs289 _dafny.Int = (Companion_Default___.SafeModuloInt((_1662_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_1662_v68), 0))).Int()).(_dafny.Int), (_1660_v66).Fm6((_1660_v66).F17(), (_1667_v71).Cardinality(), globalState))).Times((_1662_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_1662_v68), 0))).Int()).(_dafny.Int))
				_ = _rhs289
				var _rhs290 _dafny.Int = _dafny.IntOfUint32((_1570_v1).Cardinality())
				_ = _rhs290
				var _rhs291 D11 = _1676_v80
				_ = _rhs291
				var _rhs292 _dafny.Int = (_this).F11()
				_ = _rhs292
				var _rhs293 bool = (_1660_v66).F17()
				_ = _rhs293
				var _lhs239 _dafny.Array = _1662_v68
				_ = _lhs239
				var _lhs240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_1662_v68), 0))
				_ = _lhs240
				var _lhs241 *GlobalState = globalState
				_ = _lhs241
				var _lhs242 *GlobalState = globalState
				_ = _lhs242
				var _lhs243 _dafny.Array = _dafny.ArrayCastTo((_1666_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_1666_v70), 0))).Int()))
				_ = _lhs243
				var _lhs244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_dafny.ArrayCastTo((_1666_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_1666_v70), 0))).Int()))), 0))
				_ = _lhs244
				(_lhs239).ArraySet1(_rhs289, (_lhs240).Int())
				_lhs241.F1 = _rhs290
				_1676_v80 = _rhs291
				_lhs242.F1 = _rhs292
				(_lhs243).ArraySet1(_rhs293, (_lhs244).Int())
			} else {
				_1656_v62 = _dafny.Companion_Sequence_.Concatenate(_1656_v62, _1656_v62)
				var _1680_v82 _dafny.MultiSet
				_ = _1680_v82
				_1680_v82 = _dafny.MultiSetOf(_1571_v2, _1571_v2, _1571_v2, _1571_v2, _1571_v2)
				var _1681_v83 _dafny.Map
				_ = _1681_v83
				_1681_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1680_v82)
				var _1682_v84 _dafny.Sequence
				_ = _1682_v84
				_1682_v84 = _dafny.SeqOf(_dafny.MultiSetOf(_1571_v2), (func() _dafny.MultiSet {
					if (_1681_v83).Contains((_this).F11()) {
						return (_1681_v83).Get((_this).F11()).(_dafny.MultiSet)
					}
					return _1680_v82
				})())
				_1682_v84 = _dafny.SeqOf(Companion_Default___.Fm37(globalState))
				(globalState).F1 = Companion_Default___.SafeModuloInt((_this).Fm6(true, (_this).F11(), globalState), p0)
				_1571_v2 = _1571_v2
				var _rhs294 _dafny.Int = (_this).Fm6(_1571_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1653_v59, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(95))).Uint32(), func(coer96 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg96 _dafny.Int) interface{} {
						return coer96(arg96)
					}
				}((func(_1683_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1684_i4 _dafny.Int) _dafny.Int {
						return _1683_p0
					}
				})(p0))))).Cardinality()), globalState)
				_ = _rhs294
				var _rhs295 _dafny.Int = (_1653_v59).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(p0, (_this).F11()), _dafny.IntOfUint32((_1653_v59).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs295
				var _rhs296 bool = _1571_v2
				_ = _rhs296
				var _rhs297 _dafny.Int = _dafny.IntOfInt64(474)
				_ = _rhs297
				var _lhs245 *GlobalState = globalState
				_ = _lhs245
				var _lhs246 *GlobalState = globalState
				_ = _lhs246
				var _lhs247 *GlobalState = globalState
				_ = _lhs247
				var _lhs248 *GlobalState = globalState
				_ = _lhs248
				_lhs245.F1 = _rhs294
				_lhs246.F1 = _rhs295
				_lhs247.F7 = _rhs296
				_lhs248.F1 = _rhs297
			}
			var _1685_v85 _dafny.Map
			_ = _1685_v85
			_1685_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F11())
			_1685_v85 = _1685_v85
		} else {
			var _1686_v86 _dafny.Map
			_ = _1686_v86
			_1686_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(225))).Uint32(), func(coer97 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg97 _dafny.Int) interface{} {
					return coer97(arg97)
				}
			}(func(_1687_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})))
			var _1688_v87 D14
			_ = _1688_v87
			_1688_v87 = Companion_D14_.Create_DC30_(_1686_v86)
			var _1689_v88 D14
			_ = _1689_v88
			_1689_v88 = Companion_D14_.Create_DC32_(_1688_v87)
			var _1690_v89 D7
			_ = _1690_v89
			_1690_v89 = Companion_D7_.Create_DC13_(((_this).F12()).Difference((_this).F12()), (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F11())), _dafny.IntOfInt64(-673))
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))
			_ = _index237
			var _rhs298 _dafny.Sequence = _1570_v1
			_ = _rhs298
			var _rhs299 D14 = _1689_v88
			_ = _rhs299
			var _rhs300 _dafny.Int = Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(p0))).Cardinality()))
			_ = _rhs300
			var _rhs301 D7 = _1690_v89
			_ = _rhs301
			var _rhs302 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F11(), p0)
			_ = _rhs302
			var _lhs249 _dafny.Array = _1569_v0
			_ = _lhs249
			var _lhs250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))
			_ = _lhs250
			var _lhs251 *GlobalState = globalState
			_ = _lhs251
			var _lhs252 *GlobalState = globalState
			_ = _lhs252
			(_lhs249).ArraySet1(_rhs298, (_lhs250).Int())
			_1689_v88 = _rhs299
			_lhs251.F1 = _rhs300
			_1690_v89 = _rhs301
			_lhs252.F1 = _rhs302
			(globalState).F1 = (_this).F11()
			var _1691_v90 D15
			_ = _1691_v90
			_1691_v90 = Companion_D15_.Create_DC36_(p0, _1571_v2, _1571_v2, _1571_v2, (_this).F11())
			if (_1691_v90).Dtor_cf56() {
				(globalState).F6 = (_1571_v2) || (_1571_v2)
				(globalState).F1 = _dafny.IntOfInt64(841)
				var _1692_v91 _dafny.MultiSet
				_ = _1692_v91
				_1692_v91 = _dafny.MultiSetOf(p0)
				var _1693_v92 T1
				_ = _1693_v92
				var _nw276 *C4 = New_C4_()
				_ = _nw276
				_nw276.Ctor__(_1692_v91, _1571_v2, (_this).F12(), (_this).F11())
				_1693_v92 = _nw276
				var _1694_v93 _dafny.Map
				_ = _1694_v93
				_1694_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1571_v2), _1693_v92)
				var _1695_v94 _dafny.Sequence
				_ = _1695_v94
				_1695_v94 = _dafny.SeqOf(_1693_v92)
				var _1696_v95 _dafny.Array
				_ = _1696_v95
				var _nwElement0_78 T1 = _1693_v92
				_ = _nwElement0_78
				var _nw277 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_78, nil, _dafny.IntOfInt64(26))
				_ = _nw277
				(_nw277).ArraySet1(_nwElement0_78, 0)
				(_nw277).ArraySet1(_1693_v92, 1)
				(_nw277).ArraySet1(_1693_v92, 2)
				(_nw277).ArraySet1(_1693_v92, 3)
				(_nw277).ArraySet1(_1693_v92, 4)
				(_nw277).ArraySet1(_1693_v92, 5)
				(_nw277).ArraySet1(_1693_v92, 6)
				(_nw277).ArraySet1(_1693_v92, 7)
				(_nw277).ArraySet1(_1693_v92, 8)
				(_nw277).ArraySet1(_1693_v92, 9)
				(_nw277).ArraySet1(_1693_v92, 10)
				(_nw277).ArraySet1(_1693_v92, 11)
				(_nw277).ArraySet1(_1693_v92, 12)
				(_nw277).ArraySet1(_1693_v92, 13)
				(_nw277).ArraySet1(_1693_v92, 14)
				(_nw277).ArraySet1(_1693_v92, 15)
				(_nw277).ArraySet1(_1693_v92, 16)
				(_nw277).ArraySet1(_1693_v92, 17)
				(_nw277).ArraySet1(_1693_v92, 18)
				(_nw277).ArraySet1(_1693_v92, 19)
				(_nw277).ArraySet1((func() T1 {
					if (_1694_v93).Contains(_1571_v2) {
						return (_1694_v93).Get(_1571_v2).(T1)
					}
					return _1693_v92
				})(), 20)
				(_nw277).ArraySet1(_1693_v92, 21)
				(_nw277).ArraySet1(_1693_v92, 22)
				(_nw277).ArraySet1(_1693_v92, 23)
				(_nw277).ArraySet1((_1695_v94).Select((Companion_Default___.SafeIndex((Companion_D0_.Create_DC1_(_1571_v2, (_1693_v92).F11(), _1571_v2)).Dtor_cf2(), _dafny.IntOfUint32((_1695_v94).Cardinality()))).Uint32()).(T1), 24)
				(_nw277).ArraySet1(_1693_v92, 25)
				_1696_v95 = _nw277
				_1696_v95 = _1696_v95
				var _1697_v96 _dafny.Sequence
				_ = _1697_v96
				_1697_v96 = _dafny.SeqOf(p0, (_1693_v92).F11())
				var _1698_v97 D6
				_ = _1698_v97
				_1698_v97 = Companion_D6_.Create_DC11_(_dafny.IntOfUint32((_1697_v96).Cardinality()), (_this).F11(), p0)
				_1697_v96 = _dafny.Companion_Sequence_.Concatenate(_1697_v96, Companion_Default___.Fm34(_1571_v2, _1698_v97, globalState))
				(globalState).F1 = p0
			} else {
				(globalState).F6 = (_1571_v2) && (Companion_Default___.Fm0(globalState))
				var _1699_v98 D9
				_ = _1699_v98
				_1699_v98 = Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v2, p0))
				var _1700_v99 _dafny.Sequence
				_ = _1700_v99
				_1700_v99 = _dafny.SeqOf(_1571_v2)
				var _1701_v100 _dafny.Map
				_ = _1701_v100
				_1701_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v2, _dafny.IntOfUint32((_1700_v99).Cardinality()))
				var _pat_let_tv45 = _1701_v100
				_ = _pat_let_tv45
				(globalState).F1 = Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1571_v2)).Cardinality(), ((func(_pat_let37_0 D9) D9 {
					return func(_1702_dt__update__tmp_h2 D9) D9 {
						return func(_pat_let38_0 _dafny.Map) D9 {
							return func(_1703_dt__update_hcf26_h0 _dafny.Map) D9 {
								return Companion_D9_.Create_DC18_(_1703_dt__update_hcf26_h0)
							}(_pat_let38_0)
						}(_pat_let_tv45)
					}(_pat_let37_0)
				}(_1699_v98)).Dtor_cf26()).Cardinality())
				_1701_v100 = _1701_v100
				var _1704_v101 _dafny.CodePoint
				_ = _1704_v101
				_1704_v101 = _dafny.CodePoint('r')
				_1704_v101 = _1704_v101
				(globalState).F1 = (_this).F11()
			}
			var _1705_v102 *C1
			_ = _1705_v102
			var _nw278 *C1 = New_C1_()
			_ = _nw278
			_nw278.Ctor__(_1571_v2, !(_1571_v2) || (_1571_v2))
			_1705_v102 = _nw278
			_1705_v102 = _1705_v102
			var _1706_v103 _dafny.CodePoint
			_ = _1706_v103
			_1706_v103 = _dafny.CodePoint('s')
			var _1707_v104 _dafny.Set
			_ = _1707_v104
			_1707_v104 = _dafny.SetOf(_1706_v103, _1706_v103, _1706_v103, _1706_v103, _1706_v103)
			var _1708_v105 _dafny.Map
			_ = _1708_v105
			_1708_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1707_v104).Cardinality(), (_1705_v102).F18())
			var _1709_v106 _dafny.Map
			_ = _1709_v106
			_1709_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()), (func() bool {
				if (_1708_v105).Contains((_this).F11()) {
					return (_1708_v105).Get((_this).F11()).(bool)
				}
				return _1571_v2
			})())
			var _1710_v107 _dafny.Map
			_ = _1710_v107
			_1710_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v2, (_dafny.MultiSetOf(_1571_v2, _1571_v2)).Cardinality())
			var _1711_v108 _dafny.MultiSet
			_ = _1711_v108
			_1711_v108 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1570_v1).Cardinality()))
			var _1712_v109 _dafny.Array
			_ = _1712_v109
			var _nwElement0_79 bool = false
			_ = _nwElement0_79
			var _nw279 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_79, nil, _dafny.IntOfInt64(16))
			_ = _nw279
			(_nw279).ArraySet1(_nwElement0_79, 0)
			(_nw279).ArraySet1(!(_1571_v2) || ((_1705_v102).F18()), 1)
			(_nw279).ArraySet1(_1571_v2, 2)
			(_nw279).ArraySet1(((_1705_v102).F18()) && ((_1705_v102).F18()), 3)
			(_nw279).ArraySet1((func() bool {
				if (_1708_v105).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_1706_v103, _1706_v103)).Cardinality())) {
					return (_1708_v105).Get(_dafny.IntOfUint32((_dafny.SeqOf(_1706_v103, _1706_v103)).Cardinality())).(bool)
				}
				return (_1705_v102).F19()
			})(), 4)
			(_nw279).ArraySet1(!(!((_1705_v102).F19())), 5)
			(_nw279).ArraySet1(((_dafny.MultiSetOf((_1705_v102).F18(), false)).Cardinality()).Cmp((_this).F11()) != 0, 6)
			(_nw279).ArraySet1(!((_1705_v102).F19()) || ((_1705_v102).F19()), 7)
			(_nw279).ArraySet1((_1705_v102).F19(), 8)
			(_nw279).ArraySet1((p0).Cmp((_this).F11()) >= 0, 9)
			(_nw279).ArraySet1((_1705_v102).F19(), 10)
			(_nw279).ArraySet1((_1705_v102).F19(), 11)
			(_nw279).ArraySet1((p0).Cmp(_dafny.IntOfUint32(((_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence)).Cardinality())) == 0, 12)
			(_nw279).ArraySet1(((_1710_v107).Cardinality()).Cmp(_dafny.IntOfInt64(-68)) <= 0, 13)
			(_nw279).ArraySet1((_1711_v108).IsProperSubsetOf(_1711_v108), 14)
			(_nw279).ArraySet1(!((_1705_v102).F19()) || (_1571_v2), 15)
			_1712_v109 = _nw279
			var _1713_v110 D14
			_ = _1713_v110
			_1713_v110 = Companion_D14_.Create_DC31_((_1705_v102).F18(), (_this).F11(), _dafny.IntOfInt64(341))
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1712_v109), 0))
			_ = _index238
			(_1712_v109).ArraySet1(((_this).F11()).Cmp((_1713_v110).Dtor_cf46()) < 0, (_index238).Int())
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1712_v109), 0))
			_ = _index239
			(_1712_v109).ArraySet1((_1705_v102).F19(), (_index239).Int())
		}
		var _1714_v111 _dafny.Array
		_ = _1714_v111
		var _len0_34 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_34
		var _nw280 _dafny.Array
		_ = _nw280
		if _len0_34.Cmp(_dafny.Zero) == 0 {
			_nw280 = _dafny.NewArray(_len0_34)
		} else {
			var _init34 func(_dafny.Int) _dafny.Int = func(_1715_i6 _dafny.Int) _dafny.Int {
				return (_1715_i6).Times((_this).F11())
			}
			_ = _init34
			var _element0_34 = _init34(_dafny.Zero)
			_ = _element0_34
			_nw280 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
			(_nw280).ArraySet1(_element0_34, 0)
			var _nativeLen0_34 = (_len0_34).Int()
			_ = _nativeLen0_34
			for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
				(_nw280).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
			}
		}
		_1714_v111 = _nw280
		var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))
		_ = _index240
		(_1714_v111).ArraySet1(p0, (_index240).Int())
		var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))
		_ = _index241
		(_1714_v111).ArraySet1((_this).F11(), (_index241).Int())
		if _1571_v2 {
			var _1716_v112 _dafny.CodePoint
			_ = _1716_v112
			_1716_v112 = _dafny.CodePoint('k')
			var _1717_v113 _dafny.Sequence
			_ = _1717_v113
			_1717_v113 = _dafny.SeqOf(_1571_v2)
			var _1718_v114 _dafny.Set
			_ = _1718_v114
			_1718_v114 = _dafny.SetOf((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int))
			var _1719_v115 _dafny.Map
			_ = _1719_v115
			_1719_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v2, _1571_v2)
			var _1720_v116 _dafny.Map
			_ = _1720_v116
			_1720_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1718_v114, _1719_v115)
			var _1721_v117 _dafny.MultiSet
			_ = _1721_v117
			_1721_v117 = _dafny.MultiSetOf(_1716_v112, Companion_Default___.Fm20(_1717_v113, (_dafny.Zero).Minus((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int)), _1571_v2, globalState), Companion_Default___.Fm20(_1717_v113, (_1720_v116).Cardinality(), _1571_v2, globalState))
			(globalState).F1 = ((func() _dafny.Int {
				if (_1721_v117).Contains(_1716_v112) {
					return (_1721_v117).Multiplicity(_1716_v112)
				}
				return (_this).F11()
			})()).Times(_dafny.IntOfUint32((_dafny.SeqOf((_1717_v113).Select((Companion_Default___.SafeIndex((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1717_v113).Cardinality()))).Uint32()).(bool), _1571_v2)).Cardinality()))
			var _1722_v118 _dafny.Map
			_ = _1722_v118
			_1722_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v2, (((_this).F15()).Dtor_cf8()).Times(_dafny.IntOfUint32((_1570_v1).Cardinality())))
			_1722_v118 = (_1722_v118).Update(((_this).F11()).Cmp(_dafny.IntOfInt64(102)) < 0, (_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int))
			_1570_v1 = (_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence)
			(globalState).F1 = (_this).F11()
			var _1723_v119 D0
			_ = _1723_v119
			_1723_v119 = Companion_D0_.Create_DC0_((p0).Times((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int)))
			var _source25 D0 = _1723_v119
			_ = _source25
			if _source25.Is_DC1() {
				var _1724___mcc_h7 bool = _source25.Get_().(D0_DC1).Cf1
				_ = _1724___mcc_h7
				var _1725___mcc_h8 _dafny.Int = _source25.Get_().(D0_DC1).Cf2
				_ = _1725___mcc_h8
				var _1726___mcc_h9 bool = _source25.Get_().(D0_DC1).Cf3
				_ = _1726___mcc_h9
				var _1727_cf3 bool = _1726___mcc_h9
				_ = _1727_cf3
				var _1728_cf2 _dafny.Int = _1725___mcc_h8
				_ = _1728_cf2
				var _1729_cf1 bool = _1724___mcc_h7
				_ = _1729_cf1
				var _1730_v120 _dafny.Map
				_ = _1730_v120
				_1730_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int), true)
				(globalState).F1 = (((_1730_v120).Merge(_1730_v120)).Cardinality()).Minus((p0).Times((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int)))
				var _1731_v121 _dafny.MultiSet
				_ = _1731_v121
				_1731_v121 = _dafny.MultiSetOf(_1727_cf3)
				var _1732_v122 _dafny.Map
				_ = _1732_v122
				_1732_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int), (_this).F11())
				var _1733_v123 _dafny.Map
				_ = _1733_v123
				_1733_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_1728_cf2), _1732_v122)
				_1571_v2 = ((_dafny.MultiSetOf((_this).Fm4((_1731_v121).Cardinality(), _1733_v123, p0, _1729_cf1, globalState), _1729_cf1)).Difference(_dafny.MultiSetOf(_1727_cf3))).IsProperSubsetOf(_1731_v121)
				_1717_v113 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1717_v113, _1717_v113), _1717_v113), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_1728_cf2, (_dafny.MultiSetOf(_dafny.IntOfInt64(559))).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1717_v113, _1717_v113), _1717_v113)).Cardinality()))).Uint32(), ((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(518)) >= 0)
				var _1734_v124 _dafny.Map
				_ = _1734_v124
				_1734_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1716_v112, _1571_v2)
				var _1735_v125 _dafny.Array
				_ = _1735_v125
				var _nwElement0_80 bool = ((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int)).Cmp(p0) != 0
				_ = _nwElement0_80
				var _nw281 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_80, nil, _dafny.IntOfInt64(26))
				_ = _nw281
				(_nw281).ArraySet1(_nwElement0_80, 0)
				(_nw281).ArraySet1(_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm1(globalState), _dafny.UnicodeSeqOfUtf8Bytes("vlp")), 1)
				(_nw281).ArraySet1((_1717_v113).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1717_v113).Cardinality()))).Uint32()).(bool), 2)
				(_nw281).ArraySet1(!(!(_1727_cf3) || (_1729_cf1)), 3)
				(_nw281).ArraySet1((func() bool {
					if (_1734_v124).Contains(_1716_v112) {
						return (_1734_v124).Get(_1716_v112).(bool)
					}
					return _1729_cf1
				})(), 4)
				(_nw281).ArraySet1(!(true), 5)
				(_nw281).ArraySet1((_1717_v113).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.IntOfUint32((_1717_v113).Cardinality()))).Uint32()).(bool), 6)
				(_nw281).ArraySet1(_1571_v2, 7)
				(_nw281).ArraySet1(_1727_cf3, 8)
				(_nw281).ArraySet1(_1727_cf3, 9)
				(_nw281).ArraySet1(_1727_cf3, 10)
				(_nw281).ArraySet1(_1571_v2, 11)
				(_nw281).ArraySet1(_1729_cf1, 12)
				(_nw281).ArraySet1(_1727_cf3, 13)
				(_nw281).ArraySet1(_1727_cf3, 14)
				(_nw281).ArraySet1(_1571_v2, 15)
				(_nw281).ArraySet1(!((_1731_v121).IsProperSubsetOf(_dafny.MultiSetOf(!(_1571_v2)))), 16)
				(_nw281).ArraySet1(_1571_v2, 17)
				(_nw281).ArraySet1(_1729_cf1, 18)
				(_nw281).ArraySet1(_1727_cf3, 19)
				(_nw281).ArraySet1(!(_1729_cf1), 20)
				(_nw281).ArraySet1(_1571_v2, 21)
				(_nw281).ArraySet1((!(_1729_cf1)) || (_1727_cf3), 22)
				(_nw281).ArraySet1(_1727_cf3, 23)
				(_nw281).ArraySet1((_1729_cf1) || (_1571_v2), 24)
				(_nw281).ArraySet1(_1571_v2, 25)
				_1735_v125 = _nw281
				var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1735_v125), 0))
				_ = _index242
				(_1735_v125).ArraySet1(false, (_index242).Int())
				var _1736_v127 _dafny.Sequence
				_ = _1736_v127
				_1736_v127 = _dafny.SeqOf(_1718_v114)
				var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1735_v125), 0))
				_ = _index243
				var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))
				_ = _index244
				var _rhs303 bool = (func() _dafny.Set {
					var _coll66 = _dafny.NewBuilder()
					_ = _coll66
					for _iter68 := _dafny.Iterate((_1718_v114).Elements()); ; {
						_compr_66, _ok68 := _iter68()
						if !_ok68 {
							break
						}
						var _1737_v126 _dafny.Int
						_1737_v126 = interface{}(_compr_66).(_dafny.Int)
						if (_1718_v114).Contains(_1737_v126) {
							_coll66.Add(Companion_Default___.SafeModuloInt(_1737_v126, _dafny.IntOfInt64(575)))
						}
					}
					return _coll66.ToSet()
				}()).IsDisjointFrom((_1736_v127).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1736_v127).Cardinality()))).Uint32()).(_dafny.Set))
				_ = _rhs303
				var _rhs304 _dafny.Sequence = Companion_Default___.Fm1(globalState)
				_ = _rhs304
				var _rhs305 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm2(p0, p0, globalState))
				_ = _rhs305
				var _lhs253 _dafny.Array = _1735_v125
				_ = _lhs253
				var _lhs254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1735_v125), 0))
				_ = _lhs254
				var _lhs255 _dafny.Array = _1569_v0
				_ = _lhs255
				var _lhs256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))
				_ = _lhs256
				var _lhs257 *GlobalState = globalState
				_ = _lhs257
				(_lhs253).ArraySet1(_rhs303, (_lhs254).Int())
				(_lhs255).ArraySet1(_rhs304, (_lhs256).Int())
				_lhs257.F1 = _rhs305
			} else {
				var _1738___mcc_h10 _dafny.Int = _source25.Get_().(D0_DC0).Cf0
				_ = _1738___mcc_h10
				var _1739_cf0 _dafny.Int = _1738___mcc_h10
				_ = _1739_cf0
				var _1740_v128 _dafny.Array
				_ = _1740_v128
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_35
				var _nw282 _dafny.Array
				_ = _nw282
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw282 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) _dafny.MultiSet = (func(_1741_v2 bool) func(_dafny.Int) _dafny.MultiSet {
						return func(_1742_i7 _dafny.Int) _dafny.MultiSet {
							return _dafny.MultiSetOf(!(!(_1741_v2)), (Companion_D10_.Create_DC21_(_1741_v2, _1741_v2)).Dtor_cf28(), !(_1741_v2))
						}
					})(_1571_v2)
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw282 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw282).ArraySet1(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw282).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1740_v128 = _nw282
				var _rhs306 _dafny.Array = _1740_v128
				_ = _rhs306
				var _rhs307 _dafny.CodePoint = _dafny.CodePoint('j')
				_ = _rhs307
				var _rhs308 _dafny.Int = Companion_Default___.SafeDivisionInt((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
					if _1571_v2 {
						return _dafny.IntOfInt64(412)
					}
					return p0
				})())
				_ = _rhs308
				var _rhs309 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_1570_v1, (func() _dafny.Sequence {
					if _1571_v2 {
						return _dafny.UnicodeSeqOfUtf8Bytes("qqo")
					}
					return _1570_v1
				})())
				_ = _rhs309
				var _lhs258 *GlobalState = globalState
				_ = _lhs258
				var _lhs259 *GlobalState = globalState
				_ = _lhs259
				_1740_v128 = _rhs306
				_1716_v112 = _rhs307
				_lhs258.F1 = _rhs308
				_lhs259.F7 = _rhs309
				(globalState).F6 = _1571_v2
				var _1743_v129 _dafny.Array
				_ = _1743_v129
				var _nw283 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(21))
				_ = _nw283
				_1743_v129 = _nw283
				var _rhs310 bool = !(_1571_v2)
				_ = _rhs310
				var _rhs311 _dafny.Array = _1743_v129
				_ = _rhs311
				var _lhs260 *GlobalState = globalState
				_ = _lhs260
				_lhs260.F7 = _rhs310
				_1743_v129 = _rhs311
				var _1744_v130 _dafny.Array
				_ = _1744_v130
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_36
				var _nw284 _dafny.Array
				_ = _nw284
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw284 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) bool = (func(_1745_v2 bool) func(_dafny.Int) bool {
						return func(_1746_i8 _dafny.Int) bool {
							return _1745_v2
						}
					})(_1571_v2)
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw284 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw284).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw284).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1744_v130 = _nw284
				var _1747_v131 _dafny.Map
				_ = _1747_v131
				_1747_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _dafny.IntOfInt64(812))).Cardinality(), (_this).F11())
				var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_1744_v130), 0))
				_ = _index245
				(_1744_v130).ArraySet1((_1747_v131).Equals(_1747_v131), (_index245).Int())
				var _1748_v132 _dafny.Map
				_ = _1748_v132
				_1748_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1714_v111, _1571_v2)
				var _1749_v133 D10
				_ = _1749_v133
				_1749_v133 = Companion_D10_.Create_DC20_(_1748_v132)
				var _1750_v134 _dafny.Map
				_ = _1750_v134
				_1750_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1749_v133, (p0).Minus((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int)))
				var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_1744_v130), 0))
				_ = _index246
				var _rhs312 bool = _1571_v2
				_ = _rhs312
				var _rhs313 _dafny.Map = _1750_v134
				_ = _rhs313
				var _rhs314 _dafny.Array = _1714_v111
				_ = _rhs314
				var _lhs261 _dafny.Array = _1744_v130
				_ = _lhs261
				var _lhs262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_1744_v130), 0))
				_ = _lhs262
				(_lhs261).ArraySet1(_rhs312, (_lhs262).Int())
				_1750_v134 = _rhs313
				_1714_v111 = _rhs314
			}
		} else {
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_37
			var _nw285 _dafny.Array
			_ = _nw285
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw285 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) _dafny.Int = (func(_1751_v111 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_1752_i9 _dafny.Int) _dafny.Int {
						return (_1752_i9).Times((_1751_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1751_v111), 0))).Int()).(_dafny.Int))
					}
				})(_1714_v111)
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw285 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw285).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw285).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1714_v111 = _nw285
			var _1753_v135 *C5
			_ = _1753_v135
			var _nw286 *C5 = New_C5_()
			_ = _nw286
			_nw286.Ctor__()
			_1753_v135 = _nw286
			var _1754_v136 _dafny.Array
			_ = _1754_v136
			var _nwElement0_81 *C5 = _1753_v135
			_ = _nwElement0_81
			var _nw287 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_81, nil, _dafny.IntOfInt64(14))
			_ = _nw287
			(_nw287).ArraySet1(_nwElement0_81, 0)
			(_nw287).ArraySet1(_1753_v135, 1)
			(_nw287).ArraySet1(_1753_v135, 2)
			(_nw287).ArraySet1((Companion_D20_.Create_DC44_(_1753_v135)).Dtor_cf74(), 3)
			(_nw287).ArraySet1(_1753_v135, 4)
			(_nw287).ArraySet1(_1753_v135, 5)
			(_nw287).ArraySet1(_1753_v135, 6)
			(_nw287).ArraySet1(_1753_v135, 7)
			(_nw287).ArraySet1(_1753_v135, 8)
			(_nw287).ArraySet1(_1753_v135, 9)
			(_nw287).ArraySet1(_1753_v135, 10)
			(_nw287).ArraySet1(_1753_v135, 11)
			(_nw287).ArraySet1(_1753_v135, 12)
			(_nw287).ArraySet1(_1753_v135, 13)
			_1754_v136 = _nw287
			var _1755_v137 _dafny.Set
			_ = _1755_v137
			_1755_v137 = _dafny.SetOf(_1754_v136, _1754_v136, (func() _dafny.Array {
				if _1571_v2 {
					return _1754_v136
				}
				return _1754_v136
			})())
			_1755_v137 = ((_1755_v137).Union(_1755_v137)).Intersection(_dafny.SetOf(_1754_v136, _1754_v136))
			_1570_v1 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(globalState), _1570_v1)
			var _1756_v138 _dafny.Sequence
			_ = _1756_v138
			_1756_v138 = _dafny.SeqOf(_1571_v2)
			(globalState).F7 = (_1756_v138).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.IntOfUint32((_1756_v138).Cardinality()))).Uint32()).(bool)
			var _1757_v139 _dafny.CodePoint
			_ = _1757_v139
			_1757_v139 = _dafny.CodePoint('f')
			var _1758_v140 _dafny.Map
			_ = _1758_v140
			_1758_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1571_v2)
			var _1759_v141 _dafny.MultiSet
			_ = _1759_v141
			_1759_v141 = _dafny.MultiSetOf(_1758_v140, _1758_v140, _1758_v140)
			var _1760_v142 _dafny.Array
			_ = _1760_v142
			var _nwElement0_82 bool = (p0).Cmp(_dafny.IntOfInt64(883)) != 0
			_ = _nwElement0_82
			var _nw288 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_82, nil, _dafny.IntOfInt64(8))
			_ = _nw288
			(_nw288).ArraySet1(_nwElement0_82, 0)
			(_nw288).ArraySet1(_1571_v2, 1)
			(_nw288).ArraySet1(Companion_Default___.Fm0(globalState), 2)
			(_nw288).ArraySet1(_1571_v2, 3)
			(_nw288).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("c"), (_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence)), 4)
			(_nw288).ArraySet1(_1571_v2, 5)
			(_nw288).ArraySet1(!_dafny.Companion_Sequence_.Contains((_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence), _1757_v139), 6)
			(_nw288).ArraySet1((_1759_v141).IsDisjointFrom(_1759_v141), 7)
			_1760_v142 = _nw288
			var _1761_v144 _dafny.MultiSet
			_ = _1761_v144
			_1761_v144 = _dafny.MultiSetOf((_this).F11(), (_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int), p0)
			var _1762_v145 _dafny.Sequence
			_ = _1762_v145
			_1762_v145 = _dafny.SeqOf(_1761_v144)
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_1760_v142), 0))
			_ = _index247
			(_1760_v142).ArraySet1((_this).Fm4((_this).F11(), func() _dafny.Map {
				var _coll67 = _dafny.NewMapBuilder()
				_ = _coll67
				for _iter69 := _dafny.Iterate((_1762_v145).Elements()); ; {
					_compr_67, _ok69 := _iter69()
					if !_ok69 {
						break
					}
					var _1763_v143 _dafny.MultiSet
					_1763_v143 = interface{}(_compr_67).(_dafny.MultiSet)
					if _dafny.Companion_Sequence_.Contains(_1762_v145, _1763_v143) {
						_coll67.Add(_1763_v143, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus(p0)))
					}
				}
				return _coll67.ToMap()
			}(), (_this).F11(), (_1756_v138).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1756_v138).Cardinality()))).Uint32()).(bool), globalState), (_index247).Int())
			var _1764_v146 _dafny.Set
			_ = _1764_v146
			_1764_v146 = _dafny.SetOf(_1757_v139)
			var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_1760_v142), 0))
			_ = _index248
			(_1760_v142).ArraySet1((_1764_v146).IsSubsetOf(_1764_v146), (_index248).Int())
		}
		var _1765_v147 _dafny.MultiSet
		_ = _1765_v147
		_1765_v147 = _dafny.MultiSetOf(_1571_v2, _1571_v2, _1571_v2)
		var _1766_v148 _dafny.Set
		_ = _1766_v148
		_1766_v148 = _dafny.SetOf(_dafny.IntOfUint32((_1570_v1).Cardinality()))
		var _1767_v149 _dafny.Map
		_ = _1767_v149
		_1767_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1766_v148).Cardinality(), _1571_v2)
		var _1768_v150 D12
		_ = _1768_v150
		_1768_v150 = Companion_D12_.Create_DC26_(_1767_v149)
		var _1769_v151 _dafny.Map
		_ = _1769_v151
		_1769_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1768_v150, _1765_v147)
		var _1770_v152 _dafny.Array
		_ = _1770_v152
		var _nwElement0_83 _dafny.MultiSet = _1765_v147
		_ = _nwElement0_83
		var _nw289 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_83, nil, _dafny.IntOfInt64(7))
		_ = _nw289
		(_nw289).ArraySet1(_nwElement0_83, 0)
		(_nw289).ArraySet1(_1765_v147, 1)
		(_nw289).ArraySet1(_1765_v147, 2)
		(_nw289).ArraySet1(_dafny.MultiSetOf(_1571_v2, _1571_v2), 3)
		(_nw289).ArraySet1(_1765_v147, 4)
		(_nw289).ArraySet1(_1765_v147, 5)
		(_nw289).ArraySet1((func() _dafny.MultiSet {
			if (_1769_v151).Contains(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int)), _1571_v2))) {
				return (_1769_v151).Get(Companion_D12_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1714_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_1714_v111), 0))).Int()).(_dafny.Int)), _1571_v2))).(_dafny.MultiSet)
			}
			return _1765_v147
		})(), 6)
		_1770_v152 = _nw289
		r0 = _1770_v152
		var _1771_v153 _dafny.Sequence
		_ = _1771_v153
		_1771_v153 = _dafny.SeqOf(true, false)
		var _1772_v154 _dafny.MultiSet
		_ = _1772_v154
		_1772_v154 = _dafny.MultiSetOf(_dafny.IntOfUint32(((_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()))
		var _1773_v155 _dafny.Map
		_ = _1773_v155
		_1773_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update((_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32(((_1569_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1569_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), Companion_Default___.Fm20(_1771_v153, _dafny.IntOfUint32((_1570_v1).Cardinality()), _1571_v2, globalState)), _1772_v154)
		r1 = (_1773_v155).Merge((_1773_v155).Merge(_1773_v155))
		return r0, r1
	}
}
func (_this *C6) F15() D2 {
	{
		return _this._f15
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f12 _dafny.Set
	_f11 _dafny.Int
	_f14 _dafny.Sequence
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f12 = _dafny.EmptySet
	_this._f11 = _dafny.Zero
	_this._f14 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C7{}
var _ T0 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) F12() _dafny.Set {
	return _this._f12
}
func (_this *C7) F11() _dafny.Int {
	return _this._f11
}
func (_this *C7) Ctor__(f14 _dafny.Sequence, f12 _dafny.Set, f11 _dafny.Int) {
	{
		(_this)._f14 = f14
		(_this)._f12 = f12
		(_this)._f11 = f11
	}
}
func (_this *C7) Fm5(p0 D0, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		return !((Companion_D0_.Create_DC1_(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("s")).Cardinality()), true)).Dtor_cf1())
	}
}
func (_this *C7) Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var _source26 D0 = Companion_D0_.Create_DC0_((_this).F11())
		_ = _source26
		if _source26.Is_DC1() {
			var _1774___mcc_h0 bool = _source26.Get_().(D0_DC1).Cf1
			_ = _1774___mcc_h0
			var _1775___mcc_h1 _dafny.Int = _source26.Get_().(D0_DC1).Cf2
			_ = _1775___mcc_h1
			var _1776___mcc_h2 bool = _source26.Get_().(D0_DC1).Cf3
			_ = _1776___mcc_h2
			var _1777_cf3 bool = _1776___mcc_h2
			_ = _1777_cf3
			var _1778_cf2 _dafny.Int = _1775___mcc_h1
			_ = _1778_cf2
			var _1779_cf1 bool = _1774___mcc_h0
			_ = _1779_cf1
			return _1778_cf2
		} else {
			var _1780___mcc_h3 _dafny.Int = _source26.Get_().(D0_DC0).Cf0
			_ = _1780___mcc_h3
			var _1781_cf0 _dafny.Int = _1780___mcc_h3
			_ = _1781_cf0
			return (_this).F11()
		}
	}
}
func (_this *C7) Fm3(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(func(_source27 D0) _dafny.Int {
			if _source27.Is_DC1() {
				var _1782___mcc_h0 bool = _source27.Get_().(D0_DC1).Cf1
				_ = _1782___mcc_h0
				var _1783___mcc_h1 _dafny.Int = _source27.Get_().(D0_DC1).Cf2
				_ = _1783___mcc_h1
				var _1784___mcc_h2 bool = _source27.Get_().(D0_DC1).Cf3
				_ = _1784___mcc_h2
				var _1785_cf3 bool = _1784___mcc_h2
				_ = _1785_cf3
				var _1786_cf2 _dafny.Int = _1783___mcc_h1
				_ = _1786_cf2
				var _1787_cf1 bool = _1782___mcc_h0
				_ = _1787_cf1
				return ((_this).F11()).Plus((_this).F11())
			} else {
				var _1788___mcc_h3 _dafny.Int = _source27.Get_().(D0_DC0).Cf0
				_ = _1788___mcc_h3
				var _1789_cf0 _dafny.Int = _1788___mcc_h3
				_ = _1789_cf0
				return (_this).F11()
			}
		}((func() D0 {
			if true {
				return Companion_D0_.Create_DC1_(!(false), (_this).F11(), true)
			}
			return Companion_D0_.Create_DC1_(true, (_this).F11(), false)
		})()))
	}
}
func (_this *C7) Fm4(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.IsProperPrefixOf((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(97))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg98 _dafny.Int) interface{} {
				return coer98(arg98)
			}
		}(func(_1790_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		}))), _dafny.UnicodeSeqOfUtf8Bytes("kdsxn"))
	}
}
func (_this *C7) M0(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var _1791_v0 bool
		_ = _1791_v0
		_1791_v0 = true
		var _1792_v1 _dafny.MultiSet
		_ = _1792_v1
		_1792_v1 = _dafny.MultiSetOf(_1791_v0, _1791_v0, _1791_v0, _1791_v0)
		var _1793_v2 _dafny.Map
		_ = _1793_v2
		_1793_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1791_v0, _dafny.UnicodeSeqOfUtf8Bytes("lupwwwla"))
		var _hi2 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1793_v2).Contains(_1791_v0) {
				return (_1793_v2).Get(_1791_v0).(_dafny.Sequence)
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-92))).Uint32(), func(coer99 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg99 _dafny.Int) interface{} {
					return coer99(arg99)
				}
			}((func(_1794_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1795_i1 _dafny.Int) _dafny.CodePoint {
					return _1794_p0
				}
			})(p0)))
		})()).Cardinality())
		_ = _hi2
		for _1796_i0 := (_1792_v1).Cardinality(); _1796_i0.Cmp(_hi2) < 0; _1796_i0 = _1796_i0.Plus(_dafny.One) {
			var _1797_v3 _dafny.Sequence
			_ = _1797_v3
			_1797_v3 = _dafny.UnicodeSeqOfUtf8Bytes("yyslj")
			var _1798_v4 _dafny.MultiSet
			_ = _1798_v4
			_1798_v4 = _dafny.MultiSetOf(_1797_v3)
			var _1799_v5 D2
			_ = _1799_v5
			_1799_v5 = Companion_D2_.Create_DC3_(_1798_v4)
			var _1800_v6 _dafny.Sequence
			_ = _1800_v6
			_1800_v6 = _dafny.SeqOf(_1791_v0, _1791_v0, !(((_1799_v5).Dtor_cf5()).IsSubsetOf(_1798_v4)), (_1791_v0) == (false))
			_1800_v6 = _dafny.Companion_Sequence_.Update(_1800_v6, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_1800_v6).Cardinality()))).Uint32(), _1791_v0)
			(globalState).F1 = _1796_i0
			(globalState).F6 = (_dafny.SetOf(_1791_v0)).IsSubsetOf((_this).F12())
			var _1801_v7 _dafny.Map
			_ = _1801_v7
			_1801_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1791_v0)
			var _1802_v8 _dafny.Map
			_ = _1802_v8
			_1802_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1793_v2).Contains(_1791_v0) {
					return (_1793_v2).Get(_1791_v0).(_dafny.Sequence)
				}
				return _1797_v3
			})()).Cardinality()), (_1801_v7).Cardinality())
			var _1803_v9 *C2
			_ = _1803_v9
			var _nw290 *C2 = New_C2_()
			_ = _nw290
			_nw290.Ctor__(Companion_Default___.Fm20(_1800_v6, (_this).F11(), _1791_v0, globalState), ((_1802_v8).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-721), (_this).F11()))).Cardinality())
			_1803_v9 = _nw290
		}
		var _1804_v10 _dafny.MultiSet
		_ = _1804_v10
		_1804_v10 = _dafny.MultiSetOf((_this).F11())
		var _1805_v11 *C4
		_ = _1805_v11
		var _nw291 *C4 = New_C4_()
		_ = _nw291
		_nw291.Ctor__(_1804_v10, _1791_v0, ((_this).F12()).Intersection((_this).F12()), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-754), (_1792_v1).Cardinality()))
		_1805_v11 = _nw291
		var _hi3 _dafny.Int = (_this).F11()
		_ = _hi3
		for _1806_i2 := (_this).F11(); _1806_i2.Cmp(_hi3) < 0; _1806_i2 = _1806_i2.Plus(_dafny.One) {
			var _1807_v12 D9
			_ = _1807_v12
			_1807_v12 = Companion_D9_.Create_DC19_()
			var _1808_v13 _dafny.Sequence
			_ = _1808_v13
			_1808_v13 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F11()), (_this).F11())
			var _1809_v14 _dafny.Array
			_ = _1809_v14
			var _nwElement0_84 D9 = _1807_v12
			_ = _nwElement0_84
			var _nw292 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_84, nil, _dafny.IntOfInt64(15))
			_ = _nw292
			(_nw292).ArraySet1(_nwElement0_84, 0)
			(_nw292).ArraySet1(_1807_v12, 1)
			(_nw292).ArraySet1((func() D9 {
				if (_1805_v11).F17() {
					return _1807_v12
				}
				return _1807_v12
			})(), 2)
			(_nw292).ArraySet1(_1807_v12, 3)
			(_nw292).ArraySet1(Companion_D9_.Create_DC19_(), 4)
			(_nw292).ArraySet1(_1807_v12, 5)
			(_nw292).ArraySet1(Companion_D9_.Create_DC19_(), 6)
			(_nw292).ArraySet1(_1807_v12, 7)
			(_nw292).ArraySet1(_1807_v12, 8)
			(_nw292).ArraySet1(_1807_v12, 9)
			(_nw292).ArraySet1(_1807_v12, 10)
			(_nw292).ArraySet1(Companion_Default___.Fm42(_dafny.SetOf(_1808_v13), globalState), 11)
			(_nw292).ArraySet1(Companion_D9_.Create_DC19_(), 12)
			(_nw292).ArraySet1((func() D9 {
				if _1791_v0 {
					return _1807_v12
				}
				return _1807_v12
			})(), 13)
			(_nw292).ArraySet1(_1807_v12, 14)
			_1809_v14 = _nw292
			var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_1809_v14), 0))
			_ = _index249
			(_1809_v14).ArraySet1(_1807_v12, (_index249).Int())
			var _1810_v15 D8
			_ = _1810_v15
			_1810_v15 = Companion_D8_.Create_DC16_((_1805_v11).F17(), _1806_i2)
			var _1811_v16 D14
			_ = _1811_v16
			_1811_v16 = Companion_D14_.Create_DC31_(_1791_v0, (_dafny.Zero).Minus(_1806_i2), (_this).F11())
			var _1812_v17 _dafny.Sequence
			_ = _1812_v17
			_1812_v17 = _dafny.SeqOf((_1805_v11).F17(), _1791_v0, _1791_v0)
			var _1813_v18 _dafny.Array
			_ = _1813_v18
			var _nwElement0_85 bool = (_dafny.IntOfInt64(537)).Cmp(_1806_i2) <= 0
			_ = _nwElement0_85
			var _nw293 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_85, nil, _dafny.IntOfInt64(26))
			_ = _nw293
			(_nw293).ArraySet1(_nwElement0_85, 0)
			(_nw293).ArraySet1((func() bool {
				if (_1805_v11).F17() {
					return false
				}
				return (_1805_v11).F17()
			})(), 1)
			(_nw293).ArraySet1(true, 2)
			(_nw293).ArraySet1((_1805_v11).F17(), 3)
			(_nw293).ArraySet1((_1810_v15).Dtor_cf23(), 4)
			(_nw293).ArraySet1(((_this).F11()).Cmp((_this).F11()) < 0, 5)
			(_nw293).ArraySet1((_1805_v11).F17(), 6)
			(_nw293).ArraySet1((_1805_v11).F17(), 7)
			(_nw293).ArraySet1((_1805_v11).F17(), 8)
			(_nw293).ArraySet1(Companion_Default___.Fm0(globalState), 9)
			(_nw293).ArraySet1((_1806_i2).Cmp(_dafny.IntOfInt64(861)) > 0, 10)
			(_nw293).ArraySet1(((_this).F11()).Cmp((_this).F11()) < 0, 11)
			(_nw293).ArraySet1(false, 12)
			(_nw293).ArraySet1((_1806_i2).Cmp(_1806_i2) >= 0, 13)
			(_nw293).ArraySet1((_1805_v11).F17(), 14)
			(_nw293).ArraySet1((_1811_v16).Dtor_cf45(), 15)
			(_nw293).ArraySet1(((_dafny.SetOf((_1805_v11).F17(), (_1805_v11).F17())).Cardinality()).Cmp(_1806_i2) != 0, 16)
			(_nw293).ArraySet1((_1805_v11).F17(), 17)
			(_nw293).ArraySet1(_1791_v0, 18)
			(_nw293).ArraySet1(_1791_v0, 19)
			(_nw293).ArraySet1((_1805_v11).F17(), 20)
			(_nw293).ArraySet1((_1805_v11).F17(), 21)
			(_nw293).ArraySet1((_1812_v17).Select((Companion_Default___.SafeIndex((_1808_v13).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1806_i2), _dafny.IntOfUint32((_1808_v13).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1812_v17).Cardinality()))).Uint32()).(bool), 22)
			(_nw293).ArraySet1((_1805_v11).F17(), 23)
			(_nw293).ArraySet1((_1805_v11).F17(), 24)
			(_nw293).ArraySet1((_1805_v11).F17(), 25)
			_1813_v18 = _nw293
			var _1814_v19 _dafny.Sequence
			_ = _1814_v19
			_1814_v19 = _dafny.UnicodeSeqOfUtf8Bytes("kdckwdveg")
			var _1815_v20 _dafny.Sequence
			_ = _1815_v20
			_1815_v20 = _1814_v19
			var _1816_v21 _dafny.Map
			_ = _1816_v21
			_1816_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1815_v20, (_1805_v11).F17())
			var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1813_v18), 0))
			_ = _index250
			(_1813_v18).ArraySet1((func() bool {
				if (func() bool {
					if (_1816_v21).Contains(_1815_v20) {
						return (_1816_v21).Get(_1815_v20).(bool)
					}
					return (_1805_v11).F17()
				})() {
					return (_1805_v11).F17()
				}
				return false
			})(), (_index250).Int())
			var _1817_v22 _dafny.Map
			_ = _1817_v22
			_1817_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1805_v11).F17(), (_1805_v11).F17())
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_1809_v14), 0))
			_ = _index251
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1813_v18), 0))
			_ = _index252
			var _rhs315 D9 = _1807_v12
			_ = _rhs315
			var _rhs316 bool = (_1806_i2).Cmp((_1805_v11).Fm6((_1805_v11).F17(), Companion_Default___.Fm2(_1806_i2, _1806_i2, globalState), globalState)) != 0
			_ = _rhs316
			var _rhs317 bool = (_1812_v17).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_1817_v22).Cardinality(), _1806_i2), _dafny.IntOfUint32((_1812_v17).Cardinality()))).Uint32()).(bool)
			_ = _rhs317
			var _rhs318 bool = (_1805_v11).F17()
			_ = _rhs318
			var _lhs263 _dafny.Array = _1809_v14
			_ = _lhs263
			var _lhs264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_1809_v14), 0))
			_ = _lhs264
			var _lhs265 *GlobalState = globalState
			_ = _lhs265
			var _lhs266 _dafny.Array = _1813_v18
			_ = _lhs266
			var _lhs267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1813_v18), 0))
			_ = _lhs267
			var _lhs268 *GlobalState = globalState
			_ = _lhs268
			(_lhs263).ArraySet1(_rhs315, (_lhs264).Int())
			_lhs265.F6 = _rhs316
			(_lhs266).ArraySet1(_rhs317, (_lhs267).Int())
			_lhs268.F6 = _rhs318
			var _1818_v23 D16
			_ = _1818_v23
			_1818_v23 = Companion_D16_.Create_DC39_((_this).F11(), (_1813_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1813_v18), 0))).Int()).(bool), _1817_v22, p0)
			var _1819_v24 _dafny.Map
			_ = _1819_v24
			_1819_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1805_v11, _1812_v17)
			var _1820_v25 T0
			_ = _1820_v25
			var _nw294 *C3 = New_C3_()
			_ = _nw294
			_nw294.Ctor__(Companion_Default___.SafeDivisionInt((_1818_v23).Dtor_cf62(), (_this).F11()), (_dafny.Zero).Minus((_this).F11()), (_1819_v24).Cardinality(), ((_this).F12()).Difference((_this).F12()))
			_1820_v25 = _nw294
			var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1813_v18), 0))
			_ = _index253
			var _rhs319 bool = !((_1805_v11).F17())
			_ = _rhs319
			var _rhs320 bool = (_1805_v11).F17()
			_ = _rhs320
			var _rhs321 T0 = _1820_v25
			_ = _rhs321
			var _rhs322 _dafny.Int = (_1820_v25).F11()
			_ = _rhs322
			var _lhs269 _dafny.Array = _1813_v18
			_ = _lhs269
			var _lhs270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1813_v18), 0))
			_ = _lhs270
			var _lhs271 *GlobalState = globalState
			_ = _lhs271
			_1791_v0 = _rhs319
			(_lhs269).ArraySet1(_rhs320, (_lhs270).Int())
			_1820_v25 = _rhs321
			_lhs271.F1 = _rhs322
			(globalState).F6 = !_dafny.Companion_Sequence_.Equal(_1814_v19, (func() _dafny.Sequence {
				if !((_1813_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1813_v18), 0))).Int()).(bool)) {
					return _dafny.Companion_Sequence_.Update(_1814_v19, (Companion_Default___.SafeIndex(_1806_i2, _dafny.IntOfUint32((_1814_v19).Cardinality()))).Uint32(), p0)
				}
				return _1814_v19
			})())
			var _1821_v26 D7
			_ = _1821_v26
			_1821_v26 = Companion_D7_.Create_DC13_((_this).F12(), _1806_i2, (_dafny.IntOfUint32((_1814_v19).Cardinality())).Minus((_this).F11()))
			var _1822_v27 _dafny.Set
			_ = _1822_v27
			_1822_v27 = _dafny.SetOf((_1820_v25).F11())
			var _1823_v28 _dafny.Map
			_ = _1823_v28
			_1823_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1814_v19, (Companion_Default___.SafeIndex(_1806_i2, _dafny.IntOfUint32((_1814_v19).Cardinality()))).Uint32(), p0)).Cardinality()))
			var _pat_let_tv46 = _1823_v28
			_ = _pat_let_tv46
			_1821_v26 = (func() D7 {
				if (_1822_v27).IsSubsetOf(_1822_v27) {
					return Companion_Default___.Fm36((_1813_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1813_v18), 0))).Int()).(bool), Companion_Default___.Fm0(globalState), (_1820_v25).F11(), globalState)
				}
				return func(_pat_let39_0 D7) D7 {
					return func(_1824_dt__update__tmp_h0 D7) D7 {
						return func(_pat_let40_0 _dafny.Set) D7 {
							return func(_1825_dt__update_hcf18_h0 _dafny.Set) D7 {
								return func(_pat_let41_0 _dafny.Int) D7 {
									return func(_1826_dt__update_hcf20_h0 _dafny.Int) D7 {
										return Companion_D7_.Create_DC13_(_1825_dt__update_hcf18_h0, (_1824_dt__update__tmp_h0).Dtor_cf19(), _1826_dt__update_hcf20_h0)
									}(_pat_let41_0)
								}((_pat_let_tv46).Cardinality())
							}(_pat_let40_0)
						}((_this).F12())
					}(_pat_let39_0)
				}(_1821_v26)
			})()
		}
		(globalState).F1 = (_dafny.Zero).Minus((_this).F11())
		var _1827_v29 _dafny.Sequence
		_ = _1827_v29
		_1827_v29 = _dafny.UnicodeSeqOfUtf8Bytes("sctxpfjb")
		var _hi4 _dafny.Int = (_dafny.IntOfInt64(-812)).Minus(_dafny.IntOfUint32((_1827_v29).Cardinality()))
		_ = _hi4
		for _1828_i3 := (_dafny.Zero).Minus((_this).F11()); _1828_i3.Cmp(_hi4) < 0; _1828_i3 = _1828_i3.Plus(_dafny.One) {
			var _1829_v30 _dafny.Array
			_ = _1829_v30
			var _1830_v31 _dafny.Map
			_ = _1830_v31
			var _out31 _dafny.Array
			_ = _out31
			var _out32 _dafny.Map
			_ = _out32
			_out31, _out32 = (_1805_v11).M1(_1828_i3, globalState)
			_1829_v30 = _out31
			_1830_v31 = _out32
			var _1831_v32 _dafny.Sequence
			_ = _1831_v32
			_1831_v32 = _dafny.SeqOf((_1805_v11).F17(), (_1805_v11).F17())
			var _1832_v33 *C1
			_ = _1832_v33
			var _nw295 *C1 = New_C1_()
			_ = _nw295
			_nw295.Ctor__((_1805_v11).F17(), (_this).Fm5(Companion_D0_.Create_DC0_((_this).F11()), _1831_v32, _1828_i3, p0, globalState))
			_1832_v33 = _nw295
			if (_1832_v33).F18() {
				var _1833_v34 _dafny.Map
				_ = _1833_v34
				_1833_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1828_i3, p0)
				var _1834_v35 _dafny.Sequence
				_ = _1834_v35
				_1834_v35 = _dafny.SeqOf(_1828_i3)
				var _1835_v36 _dafny.Sequence
				_ = _1835_v36
				_1835_v36 = _dafny.SeqOf((_1833_v34).Cardinality(), (_1828_i3).Minus((_this).F11()), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1834_v35).Cardinality()), _1828_i3))
				_1835_v36 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(534))).Uint32(), func(coer100 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}((func(_1836_i3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1837_i4 _dafny.Int) _dafny.Int {
						return _1836_i3
					}
				})(_1828_i3)))
				var _1838_v37 _dafny.Array
				_ = _1838_v37
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_38
				var _nw296 _dafny.Array
				_ = _nw296
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw296 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) _dafny.Int = func(_1839_i5 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_1839_i5, (_this).F11())
					}
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw296 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw296).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw296).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1838_v37 = _nw296
				var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1838_v37), 0))
				_ = _index254
				(_1838_v37).ArraySet1(((_this).F11()).Times((_this).F11()), (_index254).Int())
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1838_v37), 0))
				_ = _index255
				(_1838_v37).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F11(), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1828_i3, _dafny.Companion_Sequence_.Update(_1827_v29, (Companion_Default___.SafeIndex(_1828_i3, _dafny.IntOfUint32((_1827_v29).Cardinality()))).Uint32(), p0))).Cardinality())), (_index255).Int())
				var _1840_v38 _dafny.Set
				_ = _1840_v38
				_1840_v38 = _dafny.SetOf((_this).Fm6((_1832_v33).F19(), (func() _dafny.Int {
					if ((_1805_v11).F16()).Contains(_dafny.IntOfUint32((_1831_v32).Cardinality())) {
						return ((_1805_v11).F16()).Multiplicity(_dafny.IntOfUint32((_1831_v32).Cardinality()))
					}
					return (_1838_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1838_v37), 0))).Int()).(_dafny.Int)
				})(), globalState), _dafny.IntOfInt64(532), _1828_i3)
				var _1841_v39 _dafny.Map
				_ = _1841_v39
				_1841_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_1840_v38).Union(_1840_v38))
				_1840_v38 = (func() _dafny.Set {
					if (_1841_v39).Contains((_dafny.Zero).Minus((_1838_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1838_v37), 0))).Int()).(_dafny.Int))) {
						return (_1841_v39).Get((_dafny.Zero).Minus((_1838_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1838_v37), 0))).Int()).(_dafny.Int))).(_dafny.Set)
					}
					return _dafny.SetOf(_dafny.IntOfUint32((Companion_Default___.Fm1(globalState)).Cardinality()), (_this).F11())
				})()
				var _pat_let_tv47 = _1832_v33
				_ = _pat_let_tv47
				var _pat_let_tv48 = _1832_v33
				_ = _pat_let_tv48
				var _pat_let_tv49 = _1805_v11
				_ = _pat_let_tv49
				(globalState).F7 = !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf((_1805_v11).F17(), (Companion_D19_.Create_DC43_((_1832_v33).F18(), p0, (_1805_v11).F17(), func(_pat_let42_0 D6) D6 {
					return func(_1842_dt__update__tmp_h1 D6) D6 {
						return func(_pat_let43_0 _dafny.Sequence) D6 {
							return func(_1843_dt__update_hcf13_h0 _dafny.Sequence) D6 {
								return Companion_D6_.Create_DC10_(_1843_dt__update_hcf13_h0)
							}(_pat_let43_0)
						}(_dafny.SeqOf((_pat_let_tv47).F18(), (_pat_let_tv48).F18(), (_pat_let_tv49).F17()))
					}(_pat_let42_0)
				}(Companion_D6_.Create_DC10_(_dafny.SeqOf((_1805_v11).F17()))), _1827_v29)).Dtor_cf69(), _1791_v0, (_1831_v32).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1792_v1).Contains((_1805_v11).F17()) {
						return (_1792_v1).Multiplicity((_1805_v11).F17())
					}
					return (_this).F11()
				})(), _dafny.IntOfUint32((_1831_v32).Cardinality()))).Uint32()).(bool)), _1831_v32)
				var _1844_v41 _dafny.Map
				_ = _1844_v41
				_1844_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1828_i3, (func() _dafny.Set {
					var _coll68 = _dafny.NewBuilder()
					_ = _coll68
					for _iter70 := _dafny.Iterate((_dafny.SeqOf(p0, p0)).Elements()); ; {
						_compr_68, _ok70 := _iter70()
						if !_ok70 {
							break
						}
						var _1845_v40 _dafny.CodePoint
						_1845_v40 = interface{}(_compr_68).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(p0, p0), _1845_v40) {
							_coll68.Add(_1845_v40)
						}
					}
					return _coll68.ToSet()
				}()).Cardinality())
				var _1846_v42 *C3
				_ = _1846_v42
				var _nw297 *C3 = New_C3_()
				_ = _nw297
				_nw297.Ctor__((_this).F11(), _1828_i3, _1828_i3, _dafny.SetOf((_1832_v33).F19()))
				_1846_v42 = _nw297
				var _1847_v43 _dafny.Map
				_ = _1847_v43
				_1847_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1844_v41, _1846_v42)
				var _1848_v44 _dafny.Set
				_ = _1848_v44
				_1848_v44 = _dafny.SetOf((func() *C3 {
					if (_1847_v43).Contains(_1844_v41) {
						return (_1847_v43).Get(_1844_v41).(*C3)
					}
					return _1846_v42
				})(), _1846_v42, _1846_v42)
				(globalState).F1 = (_1828_i3).Minus((_dafny.Zero).Minus((_1848_v44).Cardinality()))
			} else {
				var _1849_v45 _dafny.Set
				_ = _1849_v45
				_1849_v45 = _dafny.SetOf((_1805_v11).F17(), !((_1832_v33).Fm14((_this).F14(), globalState)))
				var _1850_v46 _dafny.Map
				_ = _1850_v46
				_1850_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_1828_i3), _1849_v45)
				var _1851_v47 D0
				_ = _1851_v47
				_1851_v47 = Companion_D0_.Create_DC0_((_this).F11())
				var _rhs323 _dafny.Int = ((_this).F11()).Minus((_this).F11())
				_ = _rhs323
				var _rhs324 bool = _1791_v0
				_ = _rhs324
				var _rhs325 _dafny.Int = _1828_i3
				_ = _rhs325
				var _rhs326 _dafny.Set = (func() _dafny.Set {
					if (_1832_v33).F19() {
						return (func() _dafny.Set {
							if (_1850_v46).Contains(_1851_v47) {
								return (_1850_v46).Get(_1851_v47).(_dafny.Set)
							}
							return Companion_Default___.Fm41(globalState)
						})()
					}
					return _1849_v45
				})()
				_ = _rhs326
				var _rhs327 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F11(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1827_v29, Companion_Default___.Fm1(globalState))).Cardinality()))
				_ = _rhs327
				var _lhs272 *GlobalState = globalState
				_ = _lhs272
				var _lhs273 *GlobalState = globalState
				_ = _lhs273
				var _lhs274 *GlobalState = globalState
				_ = _lhs274
				var _lhs275 *GlobalState = globalState
				_ = _lhs275
				_lhs272.F1 = _rhs323
				_lhs273.F7 = _rhs324
				_lhs274.F1 = _rhs325
				_1849_v45 = _rhs326
				_lhs275.F1 = _rhs327
				(globalState).F1 = (_this).F11()
				var _1852_v48 _dafny.Set
				_ = _1852_v48
				_1852_v48 = _dafny.SetOf(p0, Companion_Default___.Fm20(_1831_v32, (_1805_v11).Fm3(_1828_i3, (_1805_v11).F17(), (_1832_v33).F19(), globalState), (_1805_v11).F17(), globalState), p0, p0)
				(globalState).F1 = ((_this).F11()).Plus(((_this).Fm6((_1832_v33).F18(), (_this).F11(), globalState)).Times((_1852_v48).Cardinality()))
				(globalState).F7 = ((_dafny.Zero).Minus((_this).F11())).Cmp((_this).F11()) > 0
				var _rhs328 _dafny.Int = Companion_Default___.SafeDivisionInt(((_this).F11()).Minus((_dafny.Zero).Minus((_this).F11())), _dafny.IntOfInt64(-154))
				_ = _rhs328
				var _rhs329 bool = Companion_Default___.Fm0(globalState)
				_ = _rhs329
				var _rhs330 bool = (_1832_v33).Fm14((_this).F14(), globalState)
				_ = _rhs330
				var _rhs331 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1827_v29, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("stifwiti"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(265))).Uint32(), func(coer101 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg101 _dafny.Int) interface{} {
						return coer101(arg101)
					}
				}((func(_1853_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1854_i6 _dafny.Int) _dafny.CodePoint {
						return _1853_p0
					}
				})(p0)))), (Companion_Default___.SafeIndex(_1828_i3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("stifwiti"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(265))).Uint32(), func(coer102 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg102 _dafny.Int) interface{} {
						return coer102(arg102)
					}
				}((func(_1855_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1856_i6 _dafny.Int) _dafny.CodePoint {
						return _1855_p0
					}
				})(p0))))).Cardinality()))).Uint32(), p0))
				_ = _rhs331
				var _lhs276 *GlobalState = globalState
				_ = _lhs276
				var _lhs277 *GlobalState = globalState
				_ = _lhs277
				var _lhs278 *GlobalState = globalState
				_ = _lhs278
				_lhs276.F1 = _rhs328
				_lhs277.F6 = _rhs329
				_lhs278.F6 = _rhs330
				_1827_v29 = _rhs331
			}
			var _1857_v49 _dafny.Array
			_ = _1857_v49
			var _nw298 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
			_ = _nw298
			_1857_v49 = _nw298
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_1857_v49), 0))
			_ = _index256
			(_1857_v49).ArraySet1(_1829_v30, (_index256).Int())
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_1857_v49), 0))
			_ = _index257
			(_1857_v49).ArraySet1(_1829_v30, (_index257).Int())
		}
		var _1858_v50 _dafny.Array
		_ = _1858_v50
		var _nw299 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
		_ = _nw299
		_1858_v50 = _nw299
		_1858_v50 = _1858_v50
		r0 = Companion_Default___.Fm37(globalState)
		return r0
	}
}
func (_this *C7) M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Map) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _1859_v0 bool
		_ = _1859_v0
		_1859_v0 = false
		var _1860_v1 _dafny.CodePoint
		_ = _1860_v1
		_1860_v1 = _dafny.CodePoint('c')
		var _1861_v2 _dafny.Set
		_ = _1861_v2
		_1861_v2 = _dafny.SetOf(_1860_v1, _1860_v1, _1860_v1)
		var _1862_v3 _dafny.MultiSet
		_ = _1862_v3
		_1862_v3 = _dafny.MultiSetOf((_this).F12(), (_this).F12())
		var _1863_i0 _dafny.Int
		_ = _1863_i0
		_1863_i0 = _dafny.Zero
		{
			for (Companion_Default___.Fm53((_this).F11(), _1859_v0, (_dafny.Zero).Minus((_1861_v2).Cardinality()), globalState)).Equals(_1862_v3) {
				{
					if (_1863_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L16
					}
					_1863_i0 = (_1863_i0).Plus(_dafny.One)
					var _1864_v4 _dafny.Array
					_ = _1864_v4
					var _nw300 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
					_ = _nw300
					_1864_v4 = _nw300
					var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1864_v4), 0))
					_ = _index258
					(_1864_v4).ArraySet1(_1859_v0, (_index258).Int())
					var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1864_v4), 0))
					_ = _index259
					(_1864_v4).ArraySet1(_1859_v0, (_index259).Int())
					var _1865_v5 _dafny.Map
					_ = _1865_v5
					_1865_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(619), _1859_v0)
					var _1866_v6 _dafny.Map
					_ = _1866_v6
					_1866_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F11())
					(globalState).F7 = !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_1865_v5).Cardinality())).Equals(((_1866_v6).Update(p0, _dafny.IntOfInt64(-967))).Update((_this).F11(), p0))
					var _1867_v7 D0
					_ = _1867_v7
					_1867_v7 = Companion_D0_.Create_DC0_(p0)
					var _1868_v8 _dafny.Map
					_ = _1868_v8
					_1868_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1859_v0, (_1864_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1864_v4), 0))).Int()).(bool))
					var _1869_v9 D16
					_ = _1869_v9
					_1869_v9 = Companion_D16_.Create_DC39_(p0, _1859_v0, _1868_v8, _dafny.CodePoint('m'))
					var _1870_v10 _dafny.Sequence
					_ = _1870_v10
					_1870_v10 = _dafny.SeqOf((_1869_v9).Dtor_cf63())
					var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1864_v4), 0))
					_ = _index260
					(_1864_v4).ArraySet1((_this).Fm5(_1867_v7, _dafny.Companion_Sequence_.Concatenate(_1870_v10, _1870_v10), p0, _1860_v1, globalState), (_index260).Int())
					if (_1864_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1864_v4), 0))).Int()).(bool) {
						(globalState).F1 = p0
						var _1871_v11 _dafny.Array
						_ = _1871_v11
						var _nw301 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
						_ = _nw301
						_1871_v11 = _nw301
						var _nw302 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
						_ = _nw302
						_1871_v11 = _nw302
						(globalState).F1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("gh"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gh")).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
							if _1859_v0 {
								return _1860_v1
							}
							return _1860_v1
						})())).Cardinality())
						_1865_v5 = (_1865_v5).Update((_this).F11(), true)
						var _1872_v12 _dafny.MultiSet
						_ = _1872_v12
						_1872_v12 = _dafny.MultiSetOf(_1859_v0, (_1864_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1864_v4), 0))).Int()).(bool))
						(globalState).F7 = ((_1872_v12).Cardinality()).Cmp(_dafny.IntOfInt64(-682)) < 0
					} else {
						(globalState).F1 = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("psoaf")).Cardinality())
						var _1873_v13 _dafny.MultiSet
						_ = _1873_v13
						_1873_v13 = _dafny.MultiSetOf(p0, _dafny.IntOfInt64(331))
						var _1874_v14 _dafny.Sequence
						_ = _1874_v14
						_1874_v14 = _dafny.UnicodeSeqOfUtf8Bytes("acp")
						var _1875_v15 *C4
						_ = _1875_v15
						var _nw303 *C4 = New_C4_()
						_ = _nw303
						_nw303.Ctor__(_1873_v13, !_dafny.Companion_Sequence_.Equal(_1874_v14, _1874_v14), (_this).F12(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1860_v1)).Cardinality())
						_1875_v15 = _nw303
						(globalState).F6 = ((_dafny.IntOfInt64(-207)).Plus(p0)).Cmp(p0) <= 0
						(globalState).F1 = (_dafny.Zero).Minus(p0)
						var _1876_v16 _dafny.Array
						_ = _1876_v16
						var _nw304 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
						_ = _nw304
						_1876_v16 = _nw304
						var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.ArrayLen((_1876_v16), 0))
						_ = _index261
						(_1876_v16).ArraySet1((_dafny.Zero).Minus(p0), (_index261).Int())
						var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.ArrayLen((_1876_v16), 0))
						_ = _index262
						(_1876_v16).ArraySet1((p0).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F11())).Cardinality())), (_index262).Int())
					}
					goto C16
				}
			C16:
			}
			goto L16
		}
	L16:
		var _1877_v17 _dafny.MultiSet
		_ = _1877_v17
		_1877_v17 = _dafny.MultiSetOf(p0)
		var _1878_v18 T0
		_ = _1878_v18
		var _nw305 *C4 = New_C4_()
		_ = _nw305
		_nw305.Ctor__(_1877_v17, _1859_v0, _dafny.SetOf(_1859_v0), _dafny.IntOfUint32((_dafny.SeqOf(_1859_v0, false, _1859_v0, _1859_v0, _1859_v0)).Cardinality()))
		_1878_v18 = _nw305
		var _1879_v19 _dafny.Set
		_ = _1879_v19
		_1879_v19 = _dafny.SetOf(p0, (_dafny.SetOf(_1878_v18)).Cardinality())
		var _1880_v20 _dafny.Map
		_ = _1880_v20
		_1880_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1879_v19, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eo")).Cardinality()))
		var _1881_v21 _dafny.Map
		_ = _1881_v21
		_1881_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1880_v20).Cardinality(), Companion_Default___.Fm1(globalState))
		var _1882_v22 _dafny.Sequence
		_ = _1882_v22
		_1882_v22 = _dafny.SeqOf(_1860_v1)
		var _1883_v23 D16
		_ = _1883_v23
		_1883_v23 = Companion_D16_.Create_DC39_((_this).F11(), _1859_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1859_v0, false), _1860_v1)
		var _1884_v24 _dafny.Map
		_ = _1884_v24
		_1884_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1859_v0, _1859_v0)
		var _1885_v25 _dafny.Sequence
		_ = _1885_v25
		_1885_v25 = _dafny.SeqOf(_1884_v24, _1884_v24, _1884_v24, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1859_v0))
		var _1886_v26 _dafny.Sequence
		_ = _1886_v26
		_1886_v26 = _dafny.SeqOf(_1859_v0)
		var _1887_v27 _dafny.Map
		_ = _1887_v27
		_1887_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1859_v0, _dafny.IntOfInt64(555))
		var _1888_v28 _dafny.Array
		_ = _1888_v28
		var _nwElement0_86 bool = !(_1859_v0)
		_ = _nwElement0_86
		var _nw306 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_86, nil, _dafny.IntOfInt64(19))
		_ = _nw306
		(_nw306).ArraySet1(_nwElement0_86, 0)
		(_nw306).ArraySet1(_1859_v0, 1)
		(_nw306).ArraySet1(_1859_v0, 2)
		(_nw306).ArraySet1(((_this).F11()).Cmp((_this).F11()) >= 0, 3)
		(_nw306).ArraySet1(!(_1881_v21).Contains((_this).F11()), 4)
		(_nw306).ArraySet1(_1859_v0, 5)
		(_nw306).ArraySet1(_1859_v0, 6)
		(_nw306).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1882_v22, _dafny.SeqOf((_1883_v23).Dtor_cf65())), 7)
		(_nw306).ArraySet1(_1859_v0, 8)
		(_nw306).ArraySet1(_1859_v0, 9)
		(_nw306).ArraySet1(((_1878_v18).F11()).Cmp(_dafny.IntOfUint32((_1885_v25).Cardinality())) < 0, 10)
		(_nw306).ArraySet1(_1859_v0, 11)
		(_nw306).ArraySet1((func() bool {
			if (_1884_v24).Contains(_1859_v0) {
				return (_1884_v24).Get(_1859_v0).(bool)
			}
			return _1859_v0
		})(), 12)
		(_nw306).ArraySet1(_1859_v0, 13)
		(_nw306).ArraySet1(_1859_v0, 14)
		(_nw306).ArraySet1(_1859_v0, 15)
		(_nw306).ArraySet1(_1859_v0, 16)
		(_nw306).ArraySet1(_1859_v0, 17)
		(_nw306).ArraySet1((_1886_v26).Select((Companion_Default___.SafeIndex((Companion_D11_.Create_DC23_(true, (_1878_v18).F11(), (_1887_v27).Cardinality())).Dtor_cf32(), _dafny.IntOfUint32((_1886_v26).Cardinality()))).Uint32()).(bool), 18)
		_1888_v28 = _nw306
		var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1888_v28), 0))
		_ = _index263
		(_1888_v28).ArraySet1(_1859_v0, (_index263).Int())
		var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1888_v28), 0))
		_ = _index264
		(_1888_v28).ArraySet1(_1859_v0, (_index264).Int())
		var _1889_v29 _dafny.Array
		_ = _1889_v29
		var _nw307 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw307
		_1889_v29 = _nw307
		for _iter71 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1889_v29), 0))); ; {
			_guard_loop_2, _ok71 := _iter71()
			if !_ok71 {
				break
			}
			var _1890_i1 _dafny.Int
			_1890_i1 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1890_i1).Sign() != -1) && ((_1890_i1).Cmp(_dafny.ArrayLen((_1889_v29), 0)) < 0)) {
				(_1889_v29).ArraySet1(Companion_Default___.SafeDivisionInt(_1890_i1, (_dafny.Zero).Minus(p0)), (_1890_i1).Int())
			}
		}
		var _1891_v30 _dafny.Map
		_ = _1891_v30
		_1891_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(p0, p0), (_1888_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1888_v28), 0))).Int()).(bool))
		var _1892_v31 _dafny.Sequence
		_ = _1892_v31
		_1892_v31 = _dafny.SeqOf((_this).F11())
		_1891_v30 = (_1891_v30).Update(((_1878_v18).F11()).Plus(_dafny.IntOfUint32((_1892_v31).Cardinality())), _1859_v0)
		var _1893_v32 _dafny.MultiSet
		_ = _1893_v32
		_1893_v32 = _dafny.MultiSetOf(_1859_v0, ((_1878_v18).F11()).Cmp((_1878_v18).F11()) < 0)
		_1893_v32 = _1893_v32
		(globalState).F1 = (_this).F11()
		var _1894_v33 _dafny.Array
		_ = _1894_v33
		var _nw308 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(26))
		_ = _nw308
		_1894_v33 = _nw308
		r0 = _1894_v33
		var _1895_v34 _dafny.Map
		_ = _1895_v34
		_1895_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1882_v22, _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-102))).Uint32(), func(coer103 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg103 _dafny.Int) interface{} {
				return coer103(arg103)
			}
		}((func(_1896_v18 T0) func(_dafny.Int) _dafny.Int {
			return func(_1897_i2 _dafny.Int) _dafny.Int {
				return (_1896_v18).F11()
			}
		})(_1878_v18)))))
		r1 = (_1895_v34).Merge(_1895_v34)
		return r0, r1
	}
}
func (_this *C7) F14() _dafny.Sequence {
	{
		return _this._f14
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	_f12 _dafny.Set
	_f11 _dafny.Int
	_f13 bool
}

func New_C8_() *C8 {
	_this := C8{}

	_this._f12 = _dafny.EmptySet
	_this._f11 = _dafny.Zero
	_this._f13 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C8{}
var _ T0 = &C8{}
var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) F12() _dafny.Set {
	return _this._f12
}
func (_this *C8) F11() _dafny.Int {
	return _this._f11
}
func (_this *C8) Ctor__(f13 bool, f12 _dafny.Set, f11 _dafny.Int) {
	{
		(_this)._f13 = f13
		(_this)._f12 = f12
		(_this)._f11 = f11
	}
}
func (_this *C8) Fm5(p0 D0, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		if (_this).F13() {
			return (_this).F13()
		} else {
			return true
		}
	}
}
func (_this *C8) Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F11()
	}
}
func (_this *C8) Fm3(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.MultiSetOf((_this).F11(), (_this).F11())).Difference((_dafny.MultiSetOf((_this).F11())).Union(_dafny.MultiSetOf((_this).F11())))).Cardinality()
	}
}
func (_this *C8) Fm4(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ypsi"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(3))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg104 _dafny.Int) interface{} {
				return coer104(arg104)
			}
		}(func(_1898_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		})))
	}
}
func (_this *C8) Fm7(p0 _dafny.Int, p1 bool, p2 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F11()
	}
}
func (_this *C8) Fm8(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(-6)
	}
}
func (_this *C8) M0(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		if ((_this).F13()) && ((_this).F13()) {
			var _1899_v0 D0
			_ = _1899_v0
			_1899_v0 = Companion_D0_.Create_DC0_((_this).F11())
			var _1900_v1 _dafny.Array
			_ = _1900_v1
			var _nwElement0_87 bool = (_this).F13()
			_ = _nwElement0_87
			var _nw309 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_87, nil, _dafny.IntOfInt64(17))
			_ = _nw309
			(_nw309).ArraySet1(_nwElement0_87, 0)
			(_nw309).ArraySet1((_this).F13(), 1)
			(_nw309).ArraySet1((_this).F13(), 2)
			(_nw309).ArraySet1((_this).F13(), 3)
			(_nw309).ArraySet1((_this).F13(), 4)
			(_nw309).ArraySet1((_this).Fm5(_1899_v0, _dafny.SeqOf((_this).F13(), (_this).F13()), (_this).F11(), p0, globalState), 5)
			(_nw309).ArraySet1((_this).F13(), 6)
			(_nw309).ArraySet1((_this).F13(), 7)
			(_nw309).ArraySet1(false, 8)
			(_nw309).ArraySet1((_this).F13(), 9)
			(_nw309).ArraySet1((_this).F13(), 10)
			(_nw309).ArraySet1((_this).F13(), 11)
			(_nw309).ArraySet1((_this).F13(), 12)
			(_nw309).ArraySet1((_this).F13(), 13)
			(_nw309).ArraySet1((_this).F13(), 14)
			(_nw309).ArraySet1((_this).F13(), 15)
			(_nw309).ArraySet1((_this).F13(), 16)
			_1900_v1 = _nw309
			var _1901_v2 _dafny.Sequence
			_ = _1901_v2
			_1901_v2 = _dafny.SeqOf(_1900_v1, _1900_v1)
			var _1902_v3 _dafny.Map
			_ = _1902_v3
			_1902_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _1901_v2)
			var _1903_v4 _dafny.Sequence
			_ = _1903_v4
			_1903_v4 = _dafny.SeqOf(_1901_v2)
			var _1904_v5 _dafny.Sequence
			_ = _1904_v5
			_1904_v5 = _dafny.SeqOf((_this).F13(), !((_this).F13()), (_this).F13())
			var _1905_v6 _dafny.Array
			_ = _1905_v6
			var _nw310 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
			_ = _nw310
			_1905_v6 = _nw310
			var _1906_v7 _dafny.Map
			_ = _1906_v7
			_1906_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1905_v6, _1901_v2)
			var _1907_v8 _dafny.Array
			_ = _1907_v8
			var _nwElement0_88 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1901_v2, _1901_v2)
			_ = _nwElement0_88
			var _nw311 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_88, nil, _dafny.IntOfInt64(27))
			_ = _nw311
			(_nw311).ArraySet1(_nwElement0_88, 0)
			(_nw311).ArraySet1(_1901_v2, 1)
			(_nw311).ArraySet1(_1901_v2, 2)
			(_nw311).ArraySet1(_1901_v2, 3)
			(_nw311).ArraySet1(_1901_v2, 4)
			(_nw311).ArraySet1(_1901_v2, 5)
			(_nw311).ArraySet1(_1901_v2, 6)
			(_nw311).ArraySet1(_1901_v2, 7)
			(_nw311).ArraySet1(_1901_v2, 8)
			(_nw311).ArraySet1(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_1902_v3).Contains((_this).F13()) {
					return (_1902_v3).Get((_this).F13()).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Update((_1903_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1904_v5).Cardinality()), _dafny.IntOfUint32((_1903_v4).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_dafny.SetOf((_this).F13())).Cardinality(), _dafny.IntOfUint32(((_1903_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1904_v5).Cardinality()), _dafny.IntOfUint32((_1903_v4).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _1900_v1)
			})(), (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1902_v3).Contains((_this).F13()) {
					return (_1902_v3).Get((_this).F13()).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Update((_1903_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1904_v5).Cardinality()), _dafny.IntOfUint32((_1903_v4).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_dafny.SetOf((_this).F13())).Cardinality(), _dafny.IntOfUint32(((_1903_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1904_v5).Cardinality()), _dafny.IntOfUint32((_1903_v4).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _1900_v1)
			})()).Cardinality()))).Uint32(), (_1901_v2).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_1901_v2).Cardinality()))).Uint32()).(_dafny.Array)), 9)
			(_nw311).ArraySet1(_1901_v2, 10)
			(_nw311).ArraySet1(_1901_v2, 11)
			(_nw311).ArraySet1(_dafny.Companion_Sequence_.Update(_1901_v2, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_1901_v2).Cardinality()))).Uint32(), _1900_v1), 12)
			(_nw311).ArraySet1(_1901_v2, 13)
			(_nw311).ArraySet1(_1901_v2, 14)
			(_nw311).ArraySet1(_1901_v2, 15)
			(_nw311).ArraySet1((func() _dafny.Sequence {
				if (_1906_v7).Contains(_1905_v6) {
					return (_1906_v7).Get(_1905_v6).(_dafny.Sequence)
				}
				return _1901_v2
			})(), 16)
			(_nw311).ArraySet1(_1901_v2, 17)
			(_nw311).ArraySet1(_1901_v2, 18)
			(_nw311).ArraySet1(_1901_v2, 19)
			(_nw311).ArraySet1(_1901_v2, 20)
			(_nw311).ArraySet1(_1901_v2, 21)
			(_nw311).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1901_v2, _1901_v2), 22)
			(_nw311).ArraySet1(_1901_v2, 23)
			(_nw311).ArraySet1(_1901_v2, 24)
			(_nw311).ArraySet1(_dafny.SeqOf(_1900_v1), 25)
			(_nw311).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1901_v2, _1901_v2), 26)
			_1907_v8 = _nw311
			var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1907_v8), 0))
			_ = _index265
			(_1907_v8).ArraySet1(_1901_v2, (_index265).Int())
			var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1907_v8), 0))
			_ = _index266
			(_1907_v8).ArraySet1(_1901_v2, (_index266).Int())
			var _1908_v9 _dafny.Sequence
			_ = _1908_v9
			_1908_v9 = _dafny.UnicodeSeqOfUtf8Bytes("doukbfk")
			_1908_v9 = _1908_v9
			var _1909_v10 _dafny.CodePoint
			_ = _1909_v10
			_1909_v10 = _dafny.CodePoint('m')
			_1909_v10 = _1909_v10
			(globalState).F1 = (_this).F11()
			var _1910_v11 _dafny.Map
			_ = _1910_v11
			_1910_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F11())
			var _1911_v12 _dafny.Map
			_ = _1911_v12
			_1911_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1910_v11, _dafny.SetOf(Companion_Default___.Fm0(globalState)))
			var _1912_v13 *C2
			_ = _1912_v13
			var _nw312 *C2 = New_C2_()
			_ = _nw312
			_nw312.Ctor__(_1909_v10, ((_1911_v12).Merge(_1911_v12)).Cardinality())
			_1912_v13 = _nw312
		} else {
			var _1913_v14 _dafny.MultiSet
			_ = _1913_v14
			_1913_v14 = _dafny.MultiSetOf((_dafny.Zero).Minus((_this).F11()))
			var _1914_v15 T0
			_ = _1914_v15
			var _nw313 *C4 = New_C4_()
			_ = _nw313
			_nw313.Ctor__(_1913_v14, (_this).F13(), (_this).F12(), (_this).F11())
			_1914_v15 = _nw313
			var _1915_v16 _dafny.Map
			_ = _1915_v16
			_1915_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _1914_v15)
			var _1916_v17 _dafny.Map
			_ = _1916_v17
			_1916_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1913_v14, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_1915_v16).Cardinality()))
			var _1917_v18 _dafny.MultiSet
			_ = _1917_v18
			_1917_v18 = _dafny.MultiSetOf((_this).F13(), false)
			var _1918_v19 _dafny.Map
			_ = _1918_v19
			_1918_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm4((_this).F11(), _1916_v17, (func() _dafny.Int {
				if (_1917_v18).Contains((_this).F13()) {
					return (_1917_v18).Multiplicity((_this).F13())
				}
				return (_this).F11()
			})(), (_this).F13(), globalState), (_1914_v15).F11())
			var _1919_v21 _dafny.Sequence
			_ = _1919_v21
			_1919_v21 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iswspx")).Cardinality()))
			var _1920_v22 _dafny.Array
			_ = _1920_v22
			var _nwElement0_89 _dafny.Int = (_this).F11()
			_ = _nwElement0_89
			var _nw314 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_89, nil, _dafny.IntOfInt64(26))
			_ = _nw314
			(_nw314).ArraySet1(_nwElement0_89, 0)
			(_nw314).ArraySet1((_this).F11(), 1)
			(_nw314).ArraySet1((_1914_v15).F11(), 2)
			(_nw314).ArraySet1((func() _dafny.Int {
				if !((_this).F13()) {
					return _dafny.IntOfInt64(912)
				}
				return (_this).F11()
			})(), 3)
			(_nw314).ArraySet1(_dafny.IntOfInt64(331), 4)
			(_nw314).ArraySet1((_this).F11(), 5)
			(_nw314).ArraySet1(((_this).F12()).Cardinality(), 6)
			(_nw314).ArraySet1(Companion_Default___.SafeModuloInt((_1914_v15).F11(), _dafny.IntOfInt64(382)), 7)
			(_nw314).ArraySet1((_1914_v15).F11(), 8)
			(_nw314).ArraySet1((_1914_v15).F11(), 9)
			(_nw314).ArraySet1((_1914_v15).F11(), 10)
			(_nw314).ArraySet1((_1914_v15).F11(), 11)
			(_nw314).ArraySet1((_1914_v15).F11(), 12)
			(_nw314).ArraySet1((_1914_v15).F11(), 13)
			(_nw314).ArraySet1((_1914_v15).F11(), 14)
			(_nw314).ArraySet1((_dafny.IntOfInt64(934)).Minus((_this).F11()), 15)
			(_nw314).ArraySet1((_1914_v15).F11(), 16)
			(_nw314).ArraySet1((func() _dafny.Map {
				var _coll69 = _dafny.NewMapBuilder()
				_ = _coll69
				for _iter72 := _dafny.Iterate((_1913_v14).Elements()); ; {
					_compr_69, _ok72 := _iter72()
					if !_ok72 {
						break
					}
					var _1921_v20 _dafny.Int
					_1921_v20 = interface{}(_compr_69).(_dafny.Int)
					if (_1913_v14).Contains(_1921_v20) {
						_coll69.Add((_1921_v20).Plus((_1914_v15).F11()), _dafny.IntOfInt64(244))
					}
				}
				return _coll69.ToMap()
			}()).Cardinality(), 17)
			(_nw314).ArraySet1((_this).F11(), 18)
			(_nw314).ArraySet1((_1914_v15).F11(), 19)
			(_nw314).ArraySet1((_1914_v15).F11(), 20)
			(_nw314).ArraySet1((_1914_v15).F11(), 21)
			(_nw314).ArraySet1((_this).F11(), 22)
			(_nw314).ArraySet1(((_1914_v15).F11()).Minus((_this).F11()), 23)
			(_nw314).ArraySet1((_this).F11(), 24)
			(_nw314).ArraySet1((func() _dafny.Int {
				if true {
					return (_1919_v21).Select((Companion_Default___.SafeIndex((_1914_v15).F11(), _dafny.IntOfUint32((_1919_v21).Cardinality()))).Uint32()).(_dafny.Int)
				}
				return _dafny.IntOfInt64(723)
			})(), 25)
			_1920_v22 = _nw314
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_1920_v22), 0))
			_ = _index267
			(_1920_v22).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(51))).Uint32(), func(coer105 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}((func(_1922_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1923_i0 _dafny.Int) _dafny.CodePoint {
					return _1922_p0
				}
			})(p0)))).Cardinality()), (_index267).Int())
			var _1924_v23 D2
			_ = _1924_v23
			_1924_v23 = Companion_D2_.Create_DC4_(_dafny.IntOfInt64(894), (_1914_v15).F11(), (_this).F11())
			var _1925_v24 _dafny.Sequence
			_ = _1925_v24
			_1925_v24 = _dafny.UnicodeSeqOfUtf8Bytes("hlasr")
			var _1926_v25 *C6
			_ = _1926_v25
			var _nw315 *C6 = New_C6_()
			_ = _nw315
			_nw315.Ctor__(_1924_v23, _dafny.SetOf(!(false)), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1925_v24, _dafny.Companion_Sequence_.Update(Companion_Default___.Fm1(globalState), (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((Companion_Default___.Fm1(globalState)).Cardinality()))).Uint32(), p0))).Cardinality())))
			_1926_v25 = _nw315
			var _1927_v26 _dafny.Sequence
			_ = _1927_v26
			_1927_v26 = _dafny.SeqOf(!((_this).F13()), (_this).F13())
			var _1928_v27 _dafny.Sequence
			_ = _1928_v27
			_1928_v27 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("twrkpvv"))
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_1920_v22), 0))
			_ = _index268
			var _rhs332 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1927_v26).Cardinality()), _dafny.IntOfUint32((_1919_v21).Cardinality()))
			_ = _rhs332
			var _rhs333 _dafny.Map = (_1918_v19).Merge(_1918_v19)
			_ = _rhs333
			var _rhs334 _dafny.Int = ((_dafny.MultiSetOf(_1925_v24)).Intersection(_dafny.MultiSetFromSeq(_1928_v27))).Cardinality()
			_ = _rhs334
			var _rhs335 *C6 = _1926_v25
			_ = _rhs335
			var _rhs336 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_this).Fm4((_1914_v15).F11(), _1916_v17, (func() _dafny.Int {
					if (_1913_v14).Contains((_this).F11()) {
						return (_1913_v14).Multiplicity((_this).F11())
					}
					return _dafny.IntOfUint32((_1927_v26).Cardinality())
				})(), (_this).F13(), globalState) {
					return (_this).F11()
				}
				return (_1914_v15).F11()
			})(), (_this).F11())
			_ = _rhs336
			var _lhs279 *GlobalState = globalState
			_ = _lhs279
			var _lhs280 _dafny.Array = _1920_v22
			_ = _lhs280
			var _lhs281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_1920_v22), 0))
			_ = _lhs281
			var _lhs282 *GlobalState = globalState
			_ = _lhs282
			_lhs279.F1 = _rhs332
			_1918_v19 = _rhs333
			(_lhs280).ArraySet1(_rhs334, (_lhs281).Int())
			_1926_v25 = _rhs335
			_lhs282.F1 = _rhs336
			var _1929_v28 _dafny.Array
			_ = _1929_v28
			var _1930_v29 _dafny.Map
			_ = _1930_v29
			var _out33 _dafny.Array
			_ = _out33
			var _out34 _dafny.Map
			_ = _out34
			_out33, _out34 = (_1926_v25).M1((_1919_v21).Select((Companion_Default___.SafeIndex((_1920_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_1920_v22), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1919_v21).Cardinality()))).Uint32()).(_dafny.Int), globalState)
			_1929_v28 = _out33
			_1930_v29 = _out34
			(globalState).F1 = (_this).F11()
			var _1931_v30 _dafny.Sequence
			_ = _1931_v30
			_1931_v30 = _dafny.SeqOf(_1919_v21, _1919_v21)
			(globalState).F1 = ((_this).F11()).Plus(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_this).F13() {
					return _1919_v21
				}
				return (_1931_v30).Select((Companion_Default___.SafeIndex((_1914_v15).F11(), _dafny.IntOfUint32((_1931_v30).Cardinality()))).Uint32()).(_dafny.Sequence)
			})()).Cardinality()))
			var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_1920_v22), 0))
			_ = _index269
			(_1920_v22).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_1920_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_1920_v22), 0))).Int()).(_dafny.Int))), (_index269).Int())
		}
		var _1932_v31 _dafny.Sequence
		_ = _1932_v31
		_1932_v31 = _dafny.SeqOf(_dafny.IntOfInt64(-537))
		var _1933_v32 D0
		_ = _1933_v32
		_1933_v32 = Companion_D0_.Create_DC1_((_this).F13(), (_this).F11(), (_this).F13())
		var _1934_v33 _dafny.Sequence
		_ = _1934_v33
		_1934_v33 = _dafny.UnicodeSeqOfUtf8Bytes("liexbvcq")
		var _1935_v34 _dafny.Array
		_ = _1935_v34
		var _nw316 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
		_ = _nw316
		_1935_v34 = _nw316
		var _1936_v35 _dafny.Set
		_ = _1936_v35
		_1936_v35 = _dafny.SetOf(_1935_v34, _1935_v34)
		var _1937_v36 D21
		_ = _1937_v36
		_1937_v36 = Companion_D21_.Create_DC47_(_1936_v35)
		var _1938_v37 _dafny.Array
		_ = _1938_v37
		var _nwElement0_90 bool = !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf((_this).F11()), _1932_v31))
		_ = _nwElement0_90
		var _nw317 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_90, nil, _dafny.IntOfInt64(21))
		_ = _nw317
		(_nw317).ArraySet1(_nwElement0_90, 0)
		(_nw317).ArraySet1((_this).F13(), 1)
		(_nw317).ArraySet1((_1933_v32).Dtor_cf3(), 2)
		(_nw317).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_1934_v33, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_1934_v33).Cardinality()))).Uint32(), _dafny.CodePoint('q')), p0), 3)
		(_nw317).ArraySet1(!((_1937_v36).Dtor_cf77()).Equals(_dafny.SetOf(_1935_v34)), 4)
		(_nw317).ArraySet1((_this).F13(), 5)
		(_nw317).ArraySet1((_this).F13(), 6)
		(_nw317).ArraySet1(!((func() bool {
			if (_this).F13() {
				return (_this).F13()
			}
			return (_this).F13()
		})()), 7)
		(_nw317).ArraySet1(((_this).F11()).Cmp((_this).F11()) != 0, 8)
		(_nw317).ArraySet1((_this).F13(), 9)
		(_nw317).ArraySet1((_this).F13(), 10)
		(_nw317).ArraySet1((_this).F13(), 11)
		(_nw317).ArraySet1((_this).F13(), 12)
		(_nw317).ArraySet1((_this).F13(), 13)
		(_nw317).ArraySet1(!((_this).F13()) || ((_this).F13()), 14)
		(_nw317).ArraySet1((_this).F13(), 15)
		(_nw317).ArraySet1((_this).F13(), 16)
		(_nw317).ArraySet1((_this).F13(), 17)
		(_nw317).ArraySet1((_this).F13(), 18)
		(_nw317).ArraySet1((_this).F13(), 19)
		(_nw317).ArraySet1((_this).F13(), 20)
		_1938_v37 = _nw317
		var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
		_ = _index270
		(_1938_v37).ArraySet1(true, (_index270).Int())
		var _1939_v38 D7
		_ = _1939_v38
		_1939_v38 = Companion_D7_.Create_DC14_(Companion_D7_.Create_DC12_(_1932_v31))
		var _1940_v39 _dafny.Map
		_ = _1940_v39
		_1940_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_dafny.IntOfUint32((_1932_v31).Cardinality()), (_this).F11(), globalState), !((_this).F13()))
		var _1941_v41 _dafny.MultiSet
		_ = _1941_v41
		_1941_v41 = _dafny.MultiSetOf((_this).F11(), (_this).F11(), (_this).F11())
		var _1942_v42 _dafny.Set
		_ = _1942_v42
		_1942_v42 = _dafny.SetOf(_1940_v39, (func() _dafny.Map {
			if (_this).F13() {
				return _1940_v39
			}
			return _1940_v39
		})(), _1940_v39, (func() _dafny.Map {
			var _coll70 = _dafny.NewMapBuilder()
			_ = _coll70
			for _iter73 := _dafny.Iterate(((_1941_v41).Update((_this).F11(), Companion_Default___.Abs((_dafny.Zero).Minus((_this).F11())))).Elements()); ; {
				_compr_70, _ok73 := _iter73()
				if !_ok73 {
					break
				}
				var _1943_v40 _dafny.Int
				_1943_v40 = interface{}(_compr_70).(_dafny.Int)
				if ((_1941_v41).Update((_this).F11(), Companion_Default___.Abs((_dafny.Zero).Minus((_this).F11())))).Contains(_1943_v40) {
					_coll70.Add((_1943_v40).Times((_this).F11()), !((_this).F13()))
				}
			}
			return _coll70.ToMap()
		}()).Merge(_1940_v39), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1934_v33).Cardinality()), (_this).F13()))
		var _1944_v44 _dafny.Map
		_ = _1944_v44
		_1944_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1941_v41, func() _dafny.Map {
			var _coll71 = _dafny.NewMapBuilder()
			_ = _coll71
			for _iter74 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(50), _dafny.IntOfInt64(-202))); ; {
				_compr_71, _ok74 := _iter74()
				if !_ok74 {
					break
				}
				var _1945_v43 _dafny.Int
				_1945_v43 = interface{}(_compr_71).(_dafny.Int)
				if ((_dafny.IntOfInt64(50)).Cmp(_1945_v43) <= 0) && ((_1945_v43).Cmp(_dafny.IntOfInt64(-202)) < 0) {
					_coll71.Add(Companion_Default___.SafeDivisionInt(_1945_v43, (_this).F11()), (_dafny.Zero).Minus((_this).F11()))
				}
			}
			return _coll71.ToMap()
		}())
		var _1946_v45 _dafny.Sequence
		_ = _1946_v45
		_1946_v45 = _dafny.SeqOf(_1932_v31)
		var _1947_v46 *C7
		_ = _1947_v46
		var _nw318 *C7 = New_C7_()
		_ = _nw318
		_nw318.Ctor__(_1946_v45, (_this).F12(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F11()), true)).Cardinality())
		_1947_v46 = _nw318
		var _1948_v47 _dafny.Set
		_ = _1948_v47
		_1948_v47 = _dafny.SetOf(_1947_v46, _1947_v46)
		var _1949_v48 _dafny.Map
		_ = _1949_v48
		_1949_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm4((_this).F11(), _1944_v44, (_1948_v47).Cardinality(), (_this).F13(), globalState), (_this).F11())
		var _1950_v49 _dafny.Map
		_ = _1950_v49
		_1950_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).Fm3((_1949_v48).Cardinality(), (_this).F13(), (_this).F13(), globalState)), _dafny.IntOfInt64(452))
		var _1951_v50 _dafny.Map
		_ = _1951_v50
		_1951_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfUint32((_1932_v31).Cardinality())), _1950_v49)
		var _1952_v52 _dafny.Map
		_ = _1952_v52
		_1952_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll72 = _dafny.NewMapBuilder()
			_ = _coll72
			for _iter75 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F13())).Keys().Elements()); ; {
				_compr_72, _ok75 := _iter75()
				if !_ok75 {
					break
				}
				var _1953_v51 _dafny.Int
				_1953_v51 = interface{}(_compr_72).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F13())).Contains(_1953_v51) {
					_coll72.Add(Companion_Default___.SafeDivisionInt(_1953_v51, (_this).F11()), (_this).F13())
				}
			}
			return _coll72.ToMap()
		}(), (_this).F13())
		var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
		_ = _index271
		var _rhs337 bool = (_this).Fm4(((_this).F11()).Times((_this).F11()), _1944_v44, (_this).F11(), true, globalState)
		_ = _rhs337
		var _rhs338 D7 = _1939_v38
		_ = _rhs338
		var _rhs339 _dafny.Set = func() _dafny.Set {
			var _coll73 = _dafny.NewBuilder()
			_ = _coll73
			for _iter76 := _dafny.Iterate((_1952_v52).Keys().Elements()); ; {
				_compr_73, _ok76 := _iter76()
				if !_ok76 {
					break
				}
				var _1954_v53 _dafny.Map
				_1954_v53 = interface{}(_compr_73).(_dafny.Map)
				if (_1952_v52).Contains(_1954_v53) {
					_coll73.Add(_1954_v53)
				}
			}
			return _coll73.ToSet()
		}()
		_ = _rhs339
		var _rhs340 _dafny.Int = (_this).F11()
		_ = _rhs340
		var _lhs283 _dafny.Array = _1938_v37
		_ = _lhs283
		var _lhs284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
		_ = _lhs284
		var _lhs285 *GlobalState = globalState
		_ = _lhs285
		(_lhs283).ArraySet1(_rhs337, (_lhs284).Int())
		_1939_v38 = _rhs338
		_1942_v42 = _rhs339
		_lhs285.F1 = _rhs340
		var _1955_v54 _dafny.Map
		_ = _1955_v54
		_1955_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _dafny.CodePoint('h'))
		var _1956_v55 _dafny.Sequence
		_ = _1956_v55
		_1956_v55 = _dafny.SeqOf((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), (_this).F13())
		var _1957_v56 _dafny.Array
		_ = _1957_v56
		var _nwElement0_91 _dafny.Sequence = Companion_Default___.Fm1(globalState)
		_ = _nwElement0_91
		var _nw319 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_91, nil, _dafny.IntOfInt64(21))
		_ = _nw319
		(_nw319).ArraySet1(_nwElement0_91, 0)
		(_nw319).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("yur"), 1)
		(_nw319).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(677))).Uint32(), func(coer106 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg106 _dafny.Int) interface{} {
				return coer106(arg106)
			}
		}((func(_1958_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1959_i1 _dafny.Int) _dafny.CodePoint {
				return _1958_p0
			}
		})(p0))), 2)
		(_nw319).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(527))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg107 _dafny.Int) interface{} {
				return coer107(arg107)
			}
		}((func(_1960_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1961_i2 _dafny.Int) _dafny.CodePoint {
				return _1960_p0
			}
		})(p0))), 3)
		(_nw319).ArraySet1(_1934_v33, 4)
		(_nw319).ArraySet1(_1934_v33, 5)
		(_nw319).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ivoblt"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ivoblt")).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
			if (_1955_v54).Contains(_dafny.IntOfUint32((_1956_v55).Cardinality())) {
				return (_1955_v54).Get(_dafny.IntOfUint32((_1956_v55).Cardinality())).(_dafny.CodePoint)
			}
			return p0
		})()), 6)
		(_nw319).ArraySet1(_1934_v33, 7)
		(_nw319).ArraySet1(_1934_v33, 8)
		(_nw319).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1934_v33, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(716))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg108 _dafny.Int) interface{} {
				return coer108(arg108)
			}
		}((func(_1962_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1963_i3 _dafny.Int) _dafny.CodePoint {
				return _1962_p0
			}
		})(p0)))), 9)
		(_nw319).ArraySet1(_1934_v33, 10)
		(_nw319).ArraySet1(_1934_v33, 11)
		(_nw319).ArraySet1(_1934_v33, 12)
		(_nw319).ArraySet1(_1934_v33, 13)
		(_nw319).ArraySet1(_1934_v33, 14)
		(_nw319).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1934_v33, _1934_v33), 15)
		(_nw319).ArraySet1(_1934_v33, 16)
		(_nw319).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ugocbljk"), 17)
		(_nw319).ArraySet1(_1934_v33, 18)
		(_nw319).ArraySet1(_1934_v33, 19)
		(_nw319).ArraySet1(_1934_v33, 20)
		_1957_v56 = _nw319
		_1957_v56 = _1957_v56
		if (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool) {
			(globalState).F7 = !((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool)) || ((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool))
			(globalState).F1 = (_this).F11()
			(globalState).F6 = (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool)
			(globalState).F1 = _dafny.IntOfUint32((_1956_v55).Cardinality())
			var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
			_ = _index272
			(_1938_v37).ArraySet1((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), (_index272).Int())
		} else {
			var _1964_v57 _dafny.CodePoint
			_ = _1964_v57
			_1964_v57 = _dafny.CodePoint('k')
			var _1965_v58 _dafny.Array
			_ = _1965_v58
			var _nw320 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
			_ = _nw320
			_1965_v58 = _nw320
			var _1966_v59 _dafny.Set
			_ = _1966_v59
			_1966_v59 = _dafny.SetOf(_1965_v58)
			var _rhs341 _dafny.Int = (_dafny.IntOfInt64(163)).Times((_1966_v59).Cardinality())
			_ = _rhs341
			var _rhs342 _dafny.CodePoint = (_1934_v33).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1956_v55, _1956_v55)).Cardinality()), _dafny.IntOfUint32((_1934_v33).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs342
			var _rhs343 _dafny.Int = (_this).F11()
			_ = _rhs343
			var _lhs286 *GlobalState = globalState
			_ = _lhs286
			var _lhs287 *GlobalState = globalState
			_ = _lhs287
			_lhs286.F1 = _rhs341
			_1964_v57 = _rhs342
			_lhs287.F1 = _rhs343
			var _1967_v60 D7
			_ = _1967_v60
			_1967_v60 = Companion_D7_.Create_DC12_(_dafny.Companion_Sequence_.Concatenate(_1932_v31, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(335))).Uint32(), func(coer109 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg109 _dafny.Int) interface{} {
					return coer109(arg109)
				}
			}(func(_1968_i4 _dafny.Int) _dafny.Int {
				return (_this).F11()
			}))))
			var _source28 D7 = _1967_v60
			_ = _source28
			if _source28.Is_DC13() {
				var _1969___mcc_h0 _dafny.Set = _source28.Get_().(D7_DC13).Cf18
				_ = _1969___mcc_h0
				var _1970___mcc_h1 _dafny.Int = _source28.Get_().(D7_DC13).Cf19
				_ = _1970___mcc_h1
				var _1971___mcc_h2 _dafny.Int = _source28.Get_().(D7_DC13).Cf20
				_ = _1971___mcc_h2
				var _1972_cf20 _dafny.Int = _1971___mcc_h2
				_ = _1972_cf20
				var _1973_cf19 _dafny.Int = _1970___mcc_h1
				_ = _1973_cf19
				var _1974_cf18 _dafny.Set = _1969___mcc_h0
				_ = _1974_cf18
				var _1975_v61 D8
				_ = _1975_v61
				_1975_v61 = Companion_D8_.Create_DC16_((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), (_this).F11())
				var _1976_v62 *C3
				_ = _1976_v62
				var _nw321 *C3 = New_C3_()
				_ = _nw321
				_nw321.Ctor__(_1972_cf20, _dafny.IntOfInt64(36), _1972_cf20, (_this).F12())
				_1976_v62 = _nw321
				var _1977_v63 _dafny.Map
				_ = _1977_v63
				_1977_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1976_v62, _1976_v62.F23)
				var _1978_v64 _dafny.Set
				_ = _1978_v64
				_1978_v64 = _dafny.SetOf(_1977_v63)
				var _1979_v65 D2
				_ = _1979_v65
				_1979_v65 = Companion_D2_.Create_DC4_((_dafny.Zero).Minus((_1975_v61).Dtor_cf24()), (_1978_v64).Cardinality(), _1976_v62.F23)
				var _1980_v66 *C6
				_ = _1980_v66
				var _nw322 *C6 = New_C6_()
				_ = _nw322
				_nw322.Ctor__(_1979_v65, ((_this).F12()).Intersection(_1974_cf18), (_this).F11())
				_1980_v66 = _nw322
				var _1981_v67 _dafny.Sequence
				_ = _1981_v67
				_1981_v67 = _dafny.SeqOf(_1965_v58, _1965_v58, _1965_v58, _1965_v58)
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_1965_v58), 0))
				_ = _index273
				(_1965_v58).ArraySet1(_dafny.IntOfUint32((_1981_v67).Cardinality()), (_index273).Int())
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_1965_v58), 0))
				_ = _index274
				(_1965_v58).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_this).F11(), _1972_cf20), (_dafny.Zero).Minus((_1976_v62).F22())), (_index274).Int())
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_1965_v58), 0))
				_ = _index275
				(_1965_v58).ArraySet1((_dafny.Zero).Minus((((_1975_v61).Dtor_cf24()).Times(_dafny.IntOfInt64(-219))).Times((_this).F11())), (_index275).Int())
				var _1982_v68 _dafny.Set
				_ = _1982_v68
				_1982_v68 = _dafny.SetOf((_1976_v62).F22())
				var _1983_v70 _dafny.Sequence
				_ = _1983_v70
				_1983_v70 = _dafny.SeqOf(_1982_v68, _1982_v68, _1982_v68)
				var _1984_v71 _dafny.Map
				_ = _1984_v71
				_1984_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1964_v57, (_1965_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_1965_v58), 0))).Int()).(_dafny.Int))
				var _1985_v72 _dafny.Map
				_ = _1985_v72
				_1985_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(Companion_Default___.Fm54(_1949_v48, globalState), _1984_v71), _1982_v68)
				var _1986_v75 _dafny.Map
				_ = _1986_v75
				_1986_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1934_v33, Companion_Default___.Fm18(((_this).F12()).Cardinality(), globalState))
				var _1987_v76 _dafny.Array
				_ = _1987_v76
				var _nwElement0_92 _dafny.Set = _1982_v68
				_ = _nwElement0_92
				var _nw323 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_92, nil, _dafny.IntOfInt64(24))
				_ = _nw323
				(_nw323).ArraySet1(_nwElement0_92, 0)
				(_nw323).ArraySet1((func() _dafny.Set {
					var _coll74 = _dafny.NewBuilder()
					_ = _coll74
					for _iter77 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(212), _dafny.IntOfInt64(905))); ; {
						_compr_74, _ok77 := _iter77()
						if !_ok77 {
							break
						}
						var _1988_v69 _dafny.Int
						_1988_v69 = interface{}(_compr_74).(_dafny.Int)
						if ((_dafny.IntOfInt64(212)).Cmp(_1988_v69) <= 0) && ((_1988_v69).Cmp(_dafny.IntOfInt64(905)) < 0) {
							_coll74.Add(Companion_Default___.SafeModuloInt(_1988_v69, (_this).F11()))
						}
					}
					return _coll74.ToSet()
				}()).Union(_1982_v68), 1)
				(_nw323).ArraySet1((_1982_v68).Union(_dafny.SetOf(_1972_cf20)), 2)
				(_nw323).ArraySet1(_1982_v68, 3)
				(_nw323).ArraySet1((_1983_v70).Select((Companion_Default___.SafeIndex(_1973_cf19, _dafny.IntOfUint32((_1983_v70).Cardinality()))).Uint32()).(_dafny.Set), 4)
				(_nw323).ArraySet1((Companion_D22_.Create_DC51_(_1982_v68)).Dtor_cf86(), 5)
				(_nw323).ArraySet1(_dafny.SetOf(_dafny.IntOfInt64(49), _1972_cf20, (_1976_v62).F22()), 6)
				(_nw323).ArraySet1(_1982_v68, 7)
				(_nw323).ArraySet1(_1982_v68, 8)
				(_nw323).ArraySet1(_1982_v68, 9)
				(_nw323).ArraySet1((_1982_v68).Union(_1982_v68), 10)
				(_nw323).ArraySet1((_1982_v68).Intersection(_1982_v68), 11)
				(_nw323).ArraySet1(_1982_v68, 12)
				(_nw323).ArraySet1((func() _dafny.Set {
					if (_1985_v72).Contains(_dafny.MultiSetOf(func() _dafny.Map {
						var _coll75 = _dafny.NewMapBuilder()
						_ = _coll75
						for _iter78 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer110 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg110 _dafny.Int) interface{} {
								return coer110(arg110)
							}
						}((func(_1989_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1990_i5 _dafny.Int) _dafny.CodePoint {
								return _1989_p0
							}
						})(p0)))).Elements()); ; {
							_compr_75, _ok78 := _iter78()
							if !_ok78 {
								break
							}
							var _1991_v73 _dafny.CodePoint
							_1991_v73 = interface{}(_compr_75).(_dafny.CodePoint)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg111 _dafny.Int) interface{} {
									return coer111(arg111)
								}
							}((func(_1992_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1990_i5 _dafny.Int) _dafny.CodePoint {
									return _1992_p0
								}
							})(p0))), _1991_v73) {
								_coll75.Add(_1991_v73, _1976_v62.F23)
							}
						}
						return _coll75.ToMap()
					}())) {
						return (_1985_v72).Get(_dafny.MultiSetOf(func() _dafny.Map {
							var _coll76 = _dafny.NewMapBuilder()
							_ = _coll76
							for _iter79 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer112 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg112 _dafny.Int) interface{} {
									return coer112(arg112)
								}
							}((func(_1993_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1994_i5 _dafny.Int) _dafny.CodePoint {
									return _1993_p0
								}
							})(p0)))).Elements()); ; {
								_compr_76, _ok79 := _iter79()
								if !_ok79 {
									break
								}
								var _1995_v73 _dafny.CodePoint
								_1995_v73 = interface{}(_compr_76).(_dafny.CodePoint)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer113 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg113 _dafny.Int) interface{} {
										return coer113(arg113)
									}
								}((func(_1996_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_1994_i5 _dafny.Int) _dafny.CodePoint {
										return _1996_p0
									}
								})(p0))), _1995_v73) {
									_coll76.Add(_1995_v73, _1976_v62.F23)
								}
							}
							return _coll76.ToMap()
						}())).(_dafny.Set)
					}
					return _1982_v68
				})(), 13)
				(_nw323).ArraySet1(_dafny.SetOf(_dafny.IntOfUint32((_1934_v33).Cardinality()), (_dafny.MultiSetOf((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool))).Cardinality()), 14)
				(_nw323).ArraySet1(Companion_Default___.Fm26(globalState), 15)
				(_nw323).ArraySet1(_1982_v68, 16)
				(_nw323).ArraySet1((func() _dafny.Set {
					var _coll77 = _dafny.NewBuilder()
					_ = _coll77
					for _iter80 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-328), _dafny.IntOfInt64(346))); ; {
						_compr_77, _ok80 := _iter80()
						if !_ok80 {
							break
						}
						var _1997_v74 _dafny.Int
						_1997_v74 = interface{}(_compr_77).(_dafny.Int)
						if ((_dafny.IntOfInt64(-328)).Cmp(_1997_v74) <= 0) && ((_1997_v74).Cmp(_dafny.IntOfInt64(346)) < 0) {
							_coll77.Add(Companion_Default___.SafeDivisionInt(_1997_v74, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mak")).Cardinality())))
						}
					}
					return _coll77.ToSet()
				}()).Intersection(_dafny.SetOf(_1976_v62.F23, (_this).F11(), (_1932_v31).Select((Companion_Default___.SafeIndex((_1965_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_1965_v58), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1932_v31).Cardinality()))).Uint32()).(_dafny.Int))), 17)
				(_nw323).ArraySet1(_1982_v68, 18)
				(_nw323).ArraySet1((func() _dafny.Set {
					if (_1986_v75).Contains(_dafny.UnicodeSeqOfUtf8Bytes("hcfn")) {
						return (_1986_v75).Get(_dafny.UnicodeSeqOfUtf8Bytes("hcfn")).(_dafny.Set)
					}
					return _1982_v68
				})(), 19)
				(_nw323).ArraySet1(_1982_v68, 20)
				(_nw323).ArraySet1(_1982_v68, 21)
				(_nw323).ArraySet1(Companion_Default___.Fm26(globalState), 22)
				(_nw323).ArraySet1(_1982_v68, 23)
				_1987_v76 = _nw323
				_1987_v76 = _1987_v76
			} else if _source28.Is_DC12() {
				var _1998___mcc_h3 _dafny.Sequence = _source28.Get_().(D7_DC12).Cf17
				_ = _1998___mcc_h3
				var _1999_cf17 _dafny.Sequence = _1998___mcc_h3
				_ = _1999_cf17
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_1965_v58), 0))
				_ = _index276
				(_1965_v58).ArraySet1(((_this).F11()).Plus((_this).F11()), (_index276).Int())
				var _2000_v77 _dafny.Sequence
				_ = _2000_v77
				_2000_v77 = _dafny.SeqOf(_1940_v39, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F13()))
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_1965_v58), 0))
				_ = _index277
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
				_ = _index278
				var _rhs344 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1956_v55, _1956_v55)
				_ = _rhs344
				var _rhs345 _dafny.Int = (_this).Fm3(_dafny.IntOfInt64(-429), ((_2000_v77).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_2000_v77).Cardinality()))).Uint32()).(_dafny.Map)).Contains((_this).F11()), (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), globalState)
				_ = _rhs345
				var _rhs346 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F11(), (_this).F11())
				_ = _rhs346
				var _rhs347 bool = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_1934_v33, _1934_v33), _1934_v33)
				_ = _rhs347
				var _rhs348 _dafny.Int = (_this).F11()
				_ = _rhs348
				var _lhs288 _dafny.Array = _1965_v58
				_ = _lhs288
				var _lhs289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_1965_v58), 0))
				_ = _lhs289
				var _lhs290 *GlobalState = globalState
				_ = _lhs290
				var _lhs291 _dafny.Array = _1938_v37
				_ = _lhs291
				var _lhs292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
				_ = _lhs292
				var _lhs293 *GlobalState = globalState
				_ = _lhs293
				_1956_v55 = _rhs344
				(_lhs288).ArraySet1(_rhs345, (_lhs289).Int())
				_lhs290.F1 = _rhs346
				(_lhs291).ArraySet1(_rhs347, (_lhs292).Int())
				_lhs293.F1 = _rhs348
				_1934_v33 = _dafny.Companion_Sequence_.Concatenate(_1934_v33, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("okrvwip"), _1934_v33))
				(globalState).F1 = (_dafny.Zero).Minus((_this).F11())
				var _2001_v78 _dafny.Map
				_ = _2001_v78
				_2001_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1956_v55, _1956_v55), (_1965_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_1965_v58), 0))).Int()).(_dafny.Int))
				var _2002_v79 D7
				_ = _2002_v79
				_2002_v79 = Companion_D7_.Create_DC13_((_this).F12(), (_1955_v54).Cardinality(), _dafny.IntOfInt64(485))
				_2001_v78 = (_2001_v78).Update(Companion_Default___.Fm35(_dafny.IntOfUint32((_1934_v33).Cardinality()), (_1965_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_1965_v58), 0))).Int()).(_dafny.Int), _2002_v79, (_1941_v41).Cardinality(), globalState), (_this).F11())
			} else {
				var _2003___mcc_h4 D7 = _source28.Get_().(D7_DC14).Cf21
				_ = _2003___mcc_h4
				var _2004_cf21 D7 = _2003___mcc_h4
				_ = _2004_cf21
				_1934_v33 = _1934_v33
				(globalState).F1 = (_this).F11()
				var _2005_v80 D8
				_ = _2005_v80
				_2005_v80 = Companion_D8_.Create_DC16_((_this).F13(), (_dafny.SetOf((_this).F11())).Cardinality())
				var _2006_v81 D8
				_ = _2006_v81
				_2006_v81 = Companion_D8_.Create_DC17_(_2005_v80)
				var _2007_v82 D8
				_ = _2007_v82
				_2007_v82 = Companion_D8_.Create_DC17_(_2005_v80)
				var _2008_v83 D8
				_ = _2008_v83
				_2008_v83 = Companion_D8_.Create_DC17_(_2006_v81)
				var _2009_v84 _dafny.Map
				_ = _2009_v84
				_2009_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D8 {
					if (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool) {
						return _2008_v83
					}
					return _2008_v83
				})(), _dafny.CodePoint('p'))
				_2009_v84 = (_2009_v84).Update((func() D8 {
					if (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool) {
						return _2008_v83
					}
					return _2008_v83
				})(), p0)
				var _2010_v85 _dafny.Set
				_ = _2010_v85
				_2010_v85 = _dafny.SetOf(_dafny.MultiSetOf((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool)))
				var _2011_v86 _dafny.Map
				_ = _2011_v86
				_2011_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm12((_this).F11(), _dafny.UnicodeSeqOfUtf8Bytes("pwgdxagxs"), globalState), _1956_v55), _2010_v85)
				_2011_v86 = (_2011_v86).Update(_dafny.Companion_Sequence_.Concatenate(_1956_v55, _dafny.SeqOf(true, (_this).F13())), _2010_v85)
			}
			var _2012_v87 _dafny.Map
			_ = _2012_v87
			_2012_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1934_v33, (_this).F11())
			var _2013_v88 T0
			_ = _2013_v88
			var _nw324 *C4 = New_C4_()
			_ = _nw324
			_nw324.Ctor__(_dafny.MultiSetOf((_this).F11()), (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), (_this).F12(), (_this).F11())
			_2013_v88 = _nw324
			var _2014_v89 _dafny.Map
			_ = _2014_v89
			_2014_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2012_v87).Cardinality(), _2013_v88)
			var _2015_v90 _dafny.Map
			_ = _2015_v90
			_2015_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1950_v49).Cardinality(), _1955_v54)
			var _2016_v91 _dafny.Map
			_ = _2016_v91
			_2016_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), p0)
			var _2017_v92 D11
			_ = _2017_v92
			_2017_v92 = Companion_D11_.Create_DC23_(true, (_2013_v88).F11(), (_2016_v91).Cardinality())
			var _2018_v94 _dafny.Set
			_ = _2018_v94
			_2018_v94 = _dafny.SetOf(_1955_v54, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm46(globalState)).Cardinality(), _dafny.CodePoint('k')), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2013_v88).F11(), _dafny.CodePoint('p'))).Update((_2017_v92).Dtor_cf33(), _1964_v57), func() _dafny.Map {
				var _coll78 = _dafny.NewMapBuilder()
				_ = _coll78
				for _iter81 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(301), _dafny.IntOfInt64(662))); ; {
					_compr_78, _ok81 := _iter81()
					if !_ok81 {
						break
					}
					var _2019_v93 _dafny.Int
					_2019_v93 = interface{}(_compr_78).(_dafny.Int)
					if ((_dafny.IntOfInt64(301)).Cmp(_2019_v93) <= 0) && ((_2019_v93).Cmp(_dafny.IntOfInt64(662)) < 0) {
						_coll78.Add(Companion_Default___.SafeModuloInt(_2019_v93, (_this).F11()), p0)
					}
				}
				return _coll78.ToMap()
			}())
			if !(_2018_v94).Contains(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(315), _1964_v57)).Update((_2014_v89).Cardinality(), p0)).Merge((func() _dafny.Map {
				if (_2015_v90).Contains((_2013_v88).F11()) {
					return (_2015_v90).Get((_2013_v88).F11()).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1956_v55).Cardinality()), p0)
			})())) {
				var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
				_ = _index279
				(_1938_v37).ArraySet1(false, (_index279).Int())
				var _2020_v95 D0
				_ = _2020_v95
				_2020_v95 = Companion_D0_.Create_DC0_((_1932_v31).Select((Companion_Default___.SafeIndex((_1949_v48).Cardinality(), _dafny.IntOfUint32((_1932_v31).Cardinality()))).Uint32()).(_dafny.Int))
				var _pat_let_tv50 = _1967_v60
				_ = _pat_let_tv50
				var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
				_ = _index280
				(_1938_v37).ArraySet1((_1947_v46).Fm5(func(_pat_let44_0 D0) D0 {
					return func(_2023_dt__update__tmp_h1 D0) D0 {
						return func(_pat_let47_0 _dafny.Int) D0 {
							return func(_2024_dt__update_hcf0_h1 _dafny.Int) D0 {
								return Companion_D0_.Create_DC0_(_2024_dt__update_hcf0_h1)
							}(_pat_let47_0)
						}(_dafny.IntOfUint32(((_pat_let_tv50).Dtor_cf17()).Cardinality()))
					}(_pat_let44_0)
				}(func(_pat_let45_0 D0) D0 {
					return func(_2021_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let46_0 _dafny.Int) D0 {
							return func(_2022_dt__update_hcf0_h0 _dafny.Int) D0 {
								return Companion_D0_.Create_DC0_(_2022_dt__update_hcf0_h0)
							}(_pat_let46_0)
						}(_dafny.IntOfInt64(128))
					}(_pat_let45_0)
				}(_2020_v95)), _1956_v55, Companion_Default___.SafeDivisionInt((_this).F11(), _dafny.IntOfInt64(830)), _1964_v57, globalState), (_index280).Int())
				(globalState).F1 = (Companion_Default___.SafeDivisionInt((_this).F11(), (_this).F11())).Plus(_dafny.IntOfInt64(684))
				var _2025_v96 _dafny.Map
				_ = _2025_v96
				_2025_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13())
				var _2026_v97 D16
				_ = _2026_v97
				_2026_v97 = Companion_D16_.Create_DC39_((_1950_v49).Cardinality(), (_this).F13(), _2025_v96, _dafny.CodePoint('n'))
				var _2027_v98 *C2
				_ = _2027_v98
				var _nw325 *C2 = New_C2_()
				_ = _nw325
				_nw325.Ctor__((_2026_v97).Dtor_cf65(), (_2013_v88).F11())
				_2027_v98 = _nw325
				var _2028_v99 _dafny.Map
				_ = _2028_v99
				_2028_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2027_v98, (_2013_v88).F11())
				_2028_v99 = (_2028_v99).Update(_2027_v98, ((_2013_v88).F11()).Minus((_2013_v88).F11()))
				var _2029_v100 _dafny.Array
				_ = _2029_v100
				var _nw326 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
				_ = _nw326
				_2029_v100 = _nw326
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_2029_v100), 0))
				_ = _index281
				(_2029_v100).ArraySet1(_1935_v34, (_index281).Int())
				var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_2029_v100), 0))
				_ = _index282
				var _rhs349 _dafny.Array = _1935_v34
				_ = _rhs349
				var _rhs350 _dafny.Int = (_2013_v88).F11()
				_ = _rhs350
				var _lhs294 _dafny.Array = _2029_v100
				_ = _lhs294
				var _lhs295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_2029_v100), 0))
				_ = _lhs295
				var _lhs296 *GlobalState = globalState
				_ = _lhs296
				(_lhs294).ArraySet1(_rhs349, (_lhs295).Int())
				_lhs296.F1 = _rhs350
			} else {
				(globalState).F6 = (_2017_v92).Dtor_cf31()
				var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
				_ = _index283
				(_1938_v37).ArraySet1((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), (_index283).Int())
				(globalState).F1 = (_2013_v88).F11()
				var _2030_v101 *C2
				_ = _2030_v101
				var _nw327 *C2 = New_C2_()
				_ = _nw327
				_nw327.Ctor__(_dafny.CodePoint('f'), (_2013_v88).F11())
				_2030_v101 = _nw327
				var _pat_let_tv51 = _1934_v33
				_ = _pat_let_tv51
				(globalState).F1 = ((_2013_v88).F11()).Times((func(_pat_let48_0 D7) D7 {
					return func(_2031_dt__update__tmp_h2 D7) D7 {
						return func(_pat_let49_0 _dafny.Int) D7 {
							return func(_2032_dt__update_hcf19_h0 _dafny.Int) D7 {
								return Companion_D7_.Create_DC13_((_2031_dt__update__tmp_h2).Dtor_cf18(), _2032_dt__update_hcf19_h0, (_2031_dt__update__tmp_h2).Dtor_cf20())
							}(_pat_let49_0)
						}(_dafny.IntOfUint32((_pat_let_tv51).Cardinality()))
					}(_pat_let48_0)
				}(Companion_D7_.Create_DC13_((_this).F12(), (_2013_v88).F11(), (_2013_v88).F11()))).Dtor_cf20())
			}
			var _2033_v102 D6
			_ = _2033_v102
			_2033_v102 = Companion_D6_.Create_DC9_(_1940_v39)
			var _pat_let_tv52 = _1940_v39
			_ = _pat_let_tv52
			_2033_v102 = func(_pat_let50_0 D6) D6 {
				return func(_2034_dt__update__tmp_h3 D6) D6 {
					return func(_pat_let51_0 _dafny.Map) D6 {
						return func(_2035_dt__update_hcf12_h0 _dafny.Map) D6 {
							return Companion_D6_.Create_DC9_(_2035_dt__update_hcf12_h0)
						}(_pat_let51_0)
					}(_pat_let_tv52)
				}(_pat_let50_0)
			}(_2033_v102)
			(globalState).F6 = (!((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool))) || ((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool))
		}
		if ((func() _dafny.Int {
			if (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool) {
				return (_this).F11()
			}
			return (_this).F11()
		})()).Cmp((_this).F11()) <= 0 {
			var _rhs351 bool = (_this).F13()
			_ = _rhs351
			var _rhs352 bool = (false) && ((_this).F13())
			_ = _rhs352
			var _rhs353 _dafny.Int = (func() _dafny.Int {
				if ((_this).F13()) || (!((_this).F13())) {
					return (_this).F11()
				}
				return (_this).F11()
			})()
			_ = _rhs353
			var _lhs297 *GlobalState = globalState
			_ = _lhs297
			var _lhs298 *GlobalState = globalState
			_ = _lhs298
			var _lhs299 *GlobalState = globalState
			_ = _lhs299
			_lhs297.F7 = _rhs351
			_lhs298.F6 = _rhs352
			_lhs299.F1 = _rhs353
			_1955_v54 = (_1955_v54).Update(_dafny.IntOfInt64(569), p0)
			(globalState).F7 = (((_this).F11()).Minus((_this).F11())).Cmp((_this).F11()) > 0
			var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
			_ = _index284
			(_1938_v37).ArraySet1((_this).F13(), (_index284).Int())
			var _pat_let_tv53 = _1936_v35
			_ = _pat_let_tv53
			var _2036_v103 _dafny.MultiSet
			_ = _2036_v103
			_2036_v103 = _dafny.MultiSetOf(_1937_v36, _1937_v36, _1937_v36, _1937_v36, func(_pat_let52_0 D21) D21 {
				return func(_2037_dt__update__tmp_h4 D21) D21 {
					return func(_pat_let53_0 _dafny.Set) D21 {
						return func(_2038_dt__update_hcf77_h0 _dafny.Set) D21 {
							return Companion_D21_.Create_DC47_(_2038_dt__update_hcf77_h0)
						}(_pat_let53_0)
					}(_pat_let_tv53)
				}(_pat_let52_0)
			}(_1937_v36))
			_2036_v103 = (_2036_v103).Update(_1937_v36, Companion_Default___.Abs((_this).F11()))
		} else {
			(globalState).F1 = (_this).F11()
			(globalState).F1 = _dafny.IntOfInt64(112)
			var _2039_v104 D2
			_ = _2039_v104
			_2039_v104 = Companion_D2_.Create_DC4_(_dafny.IntOfInt64(121), (_this).F11(), _dafny.IntOfInt64(728))
			var _2040_v105 _dafny.Sequence
			_ = _2040_v105
			_2040_v105 = _dafny.SeqOf(_2039_v104)
			var _2041_v106 _dafny.MultiSet
			_ = _2041_v106
			_2041_v106 = _dafny.MultiSetOf((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), (_this).F13())
			(globalState).F1 = ((_this).F11()).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2040_v105, _2040_v105), (Companion_Default___.SafeIndex((_2041_v106).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2040_v105, _2040_v105)).Cardinality()))).Uint32(), Companion_D2_.Create_DC4_(_dafny.IntOfUint32((_1932_v31).Cardinality()), (_this).F11(), (_this).F11()))).Cardinality()))
			var _2042_v107 _dafny.Sequence
			_ = _2042_v107
			_2042_v107 = _dafny.SeqOf(_1956_v55)
			var _2043_v108 *C0
			_ = _2043_v108
			var _nw328 *C0 = New_C0_()
			_ = _nw328
			_nw328.Ctor__(_dafny.Companion_Sequence_.Concatenate((_2042_v107).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_2042_v107).Cardinality()))).Uint32()).(_dafny.Sequence), _1956_v55), (_this).Fm7((_this).F11(), (_this).F13(), _dafny.MultiSetFromSeq(_1956_v55), globalState))
			_2043_v108 = _nw328
			_2043_v108 = _2043_v108
			(globalState).F1 = _dafny.IntOfInt64(901)
		}
		var _2044_v109 D16
		_ = _2044_v109
		_2044_v109 = Companion_D16_.Create_DC38_(_1950_v49)
		var _source29 D15 = func(_source30 D16) D15 {
			if _source30.Is_DC39() {
				var _2045___mcc_h17 _dafny.Int = _source30.Get_().(D16_DC39).Cf62
				_ = _2045___mcc_h17
				var _2046___mcc_h18 bool = _source30.Get_().(D16_DC39).Cf63
				_ = _2046___mcc_h18
				var _2047___mcc_h19 _dafny.Map = _source30.Get_().(D16_DC39).Cf64
				_ = _2047___mcc_h19
				var _2048___mcc_h20 _dafny.CodePoint = _source30.Get_().(D16_DC39).Cf65
				_ = _2048___mcc_h20
				var _2049_cf65 _dafny.CodePoint = _2048___mcc_h20
				_ = _2049_cf65
				var _2050_cf64 _dafny.Map = _2047___mcc_h19
				_ = _2050_cf64
				var _2051_cf63 bool = _2046___mcc_h18
				_ = _2051_cf63
				var _2052_cf62 _dafny.Int = _2045___mcc_h17
				_ = _2052_cf62
				return Companion_D15_.Create_DC34_()
			} else {
				var _2053___mcc_h21 _dafny.Map = _source30.Get_().(D16_DC38).Cf61
				_ = _2053___mcc_h21
				var _2054_cf61 _dafny.Map = _2053___mcc_h21
				_ = _2054_cf61
				return Companion_D15_.Create_DC34_()
			}
		}(_2044_v109)
		_ = _source29
		if _source29.Is_DC34() {
			(globalState).F1 = (func() _dafny.Int {
				if false {
					return (_this).F11()
				}
				return (func() _dafny.Int {
					if Companion_Default___.Fm0(globalState) {
						return (_dafny.Zero).Minus((_this).F11())
					}
					return (_this).F11()
				})()
			})()
			var _2055_v110 T0
			_ = _2055_v110
			var _nw329 *C2 = New_C2_()
			_ = _nw329
			_nw329.Ctor__(Companion_Default___.Fm20(_dafny.SeqOf((_this).F13(), !((_this).F13()), (_this).F13(), (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool)), _dafny.IntOfInt64(311), (_this).F13(), globalState), _dafny.IntOfInt64(-613))
			_2055_v110 = _nw329
			var _2056_v111 _dafny.Map
			_ = _2056_v111
			_2056_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2055_v110, (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool))
			var _2057_v112 *C2
			_ = _2057_v112
			var _nw330 *C2 = New_C2_()
			_ = _nw330
			_nw330.Ctor__(p0, (_this).F11())
			_2057_v112 = _nw330
			var _2058_v113 _dafny.Set
			_ = _2058_v113
			_2058_v113 = _dafny.SetOf(_2057_v112, _2057_v112)
			var _rhs354 _dafny.Map = (_2056_v111).Merge(_2056_v111)
			_ = _rhs354
			var _rhs355 bool = !((_1956_v55).Select((Companion_Default___.SafeIndex(((_2058_v113).Union(_2058_v113)).Cardinality(), _dafny.IntOfUint32((_1956_v55).Cardinality()))).Uint32()).(bool))
			_ = _rhs355
			var _rhs356 _dafny.Int = (_this).F11()
			_ = _rhs356
			var _rhs357 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_this).F13() {
					return (_dafny.Zero).Minus((_this).F11())
				}
				return _dafny.IntOfInt64(25)
			})(), (_this).F11())
			_ = _rhs357
			var _lhs300 *GlobalState = globalState
			_ = _lhs300
			var _lhs301 *GlobalState = globalState
			_ = _lhs301
			var _lhs302 *GlobalState = globalState
			_ = _lhs302
			_2056_v111 = _rhs354
			_lhs300.F6 = _rhs355
			_lhs301.F1 = _rhs356
			_lhs302.F1 = _rhs357
			var _2059_v114 _dafny.Map
			_ = _2059_v114
			_2059_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2044_v109, (_2055_v110).F11())
			_2059_v114 = (_2059_v114).Update(_2044_v109, _dafny.IntOfInt64(-321))
			var _2060_v115 _dafny.Array
			_ = _2060_v115
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_39
			var _nw331 _dafny.Array
			_ = _nw331
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw331 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.Map = (func(_2061_v110 T0, _2062_v37 _dafny.Array) func(_dafny.Int) _dafny.Map {
					return func(_2063_i6 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_2061_v110).F11()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2062_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_2062_v37), 0))).Int()).(bool), (_2062_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_2062_v37), 0))).Int()).(bool)))
					}
				})(_2055_v110, _1938_v37)
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw331 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw331).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw331).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_2060_v115 = _nw331
			var _2064_v116 _dafny.Set
			_ = _2064_v116
			_2064_v116 = _dafny.SetOf((_2055_v110).F11(), (_2055_v110).F11())
			var _2065_v117 _dafny.Map
			_ = _2065_v117
			_2065_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
			var _2066_v118 _dafny.Map
			_ = _2066_v118
			_2066_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2064_v116, _2065_v117)
			var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_2060_v115), 0))
			_ = _index285
			(_2060_v115).ArraySet1(_2066_v118, (_index285).Int())
			var _2067_v119 _dafny.Sequence
			_ = _2067_v119
			_2067_v119 = _dafny.SeqOf(_1956_v55, _1956_v55, _1956_v55)
			var _2068_v120 D14
			_ = _2068_v120
			_2068_v120 = Companion_D14_.Create_DC31_(false, (_this).F11(), _dafny.IntOfInt64(397))
			var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_2060_v115), 0))
			_ = _index286
			var _rhs358 _dafny.Map = ((Companion_Default___.Fm55((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), _dafny.IntOfUint32((_1934_v33).Cardinality()), (_this).F13(), globalState)).Merge(_2066_v118)).Update(_2064_v116, _2065_v117)
			_ = _rhs358
			var _rhs359 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate((_2067_v119).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_2067_v119).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqOf((_2068_v120).Dtor_cf45(), (_this).F13())), _1956_v55)
			_ = _rhs359
			var _lhs303 _dafny.Array = _2060_v115
			_ = _lhs303
			var _lhs304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_2060_v115), 0))
			_ = _lhs304
			var _lhs305 *GlobalState = globalState
			_ = _lhs305
			(_lhs303).ArraySet1(_rhs358, (_lhs304).Int())
			_lhs305.F6 = _rhs359
		} else if _source29.Is_DC35() {
			var _2069___mcc_h5 bool = _source29.Get_().(D15_DC35).Cf50
			_ = _2069___mcc_h5
			var _2070___mcc_h6 T1 = _source29.Get_().(D15_DC35).Cf51
			_ = _2070___mcc_h6
			var _2071___mcc_h7 bool = _source29.Get_().(D15_DC35).Cf52
			_ = _2071___mcc_h7
			var _2072___mcc_h8 _dafny.Sequence = _source29.Get_().(D15_DC35).Cf53
			_ = _2072___mcc_h8
			var _2073___mcc_h9 _dafny.Int = _source29.Get_().(D15_DC35).Cf54
			_ = _2073___mcc_h9
			var _2074_cf54 _dafny.Int = _2073___mcc_h9
			_ = _2074_cf54
			var _2075_cf53 _dafny.Sequence = _2072___mcc_h8
			_ = _2075_cf53
			var _2076_cf52 bool = _2071___mcc_h7
			_ = _2076_cf52
			var _2077_cf51 T1 = _2070___mcc_h6
			_ = _2077_cf51
			var _2078_cf50 bool = _2069___mcc_h5
			_ = _2078_cf50
			var _2079_v121 *C7
			_ = _2079_v121
			var _nw332 *C7 = New_C7_()
			_ = _nw332
			_nw332.Ctor__((_1947_v46).F14(), ((_2077_cf51).F12()).Difference(_dafny.SetOf(_2078_cf50, (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), false)), (_2077_cf51).F11())
			_2079_v121 = _nw332
			var _2080_v122 D2
			_ = _2080_v122
			_2080_v122 = Companion_D2_.Create_DC4_((_this).F11(), (_this).F11(), (_this).F11())
			var _source31 D2 = _2080_v122
			_ = _source31
			if _source31.Is_DC4() {
				var _2081___mcc_h22 _dafny.Int = _source31.Get_().(D2_DC4).Cf6
				_ = _2081___mcc_h22
				var _2082___mcc_h23 _dafny.Int = _source31.Get_().(D2_DC4).Cf7
				_ = _2082___mcc_h23
				var _2083___mcc_h24 _dafny.Int = _source31.Get_().(D2_DC4).Cf8
				_ = _2083___mcc_h24
				var _2084_cf8 _dafny.Int = _2083___mcc_h24
				_ = _2084_cf8
				var _2085_cf7 _dafny.Int = _2082___mcc_h23
				_ = _2085_cf7
				var _2086_cf6 _dafny.Int = _2081___mcc_h22
				_ = _2086_cf6
				var _2087_v124 _dafny.Map
				_ = _2087_v124
				_2087_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm27((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), globalState), _1932_v31)).Cardinality()), func() _dafny.Map {
					var _coll79 = _dafny.NewMapBuilder()
					_ = _coll79
					for _iter82 := _dafny.Iterate((_1940_v39).Keys().Elements()); ; {
						_compr_79, _ok82 := _iter82()
						if !_ok82 {
							break
						}
						var _2088_v123 _dafny.Int
						_2088_v123 = interface{}(_compr_79).(_dafny.Int)
						if (_1940_v39).Contains(_2088_v123) {
							_coll79.Add((_2088_v123).Plus(_2084_cf8), (_this).F11())
						}
					}
					return _coll79.ToMap()
				}())
				_2087_v124 = (_2087_v124).Update(_2085_cf7, _1950_v49)
				var _2089_v125 *C0
				_ = _2089_v125
				var _nw333 *C0 = New_C0_()
				_ = _nw333
				_nw333.Ctor__(_1956_v55, _2074_cf54)
				_2089_v125 = _nw333
				(globalState).F6 = (_2078_cf50) && (_2078_cf50)
				var _2090_v126 _dafny.Map
				_ = _2090_v126
				_2090_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2076_cf52)
				var _2091_v128 _dafny.Set
				_ = _2091_v128
				_2091_v128 = _dafny.SetOf(_2090_v126, func() _dafny.Map {
					var _coll80 = _dafny.NewMapBuilder()
					_ = _coll80
					for _iter83 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(203))).Uint32(), func(coer114 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg114 _dafny.Int) interface{} {
							return coer114(arg114)
						}
					}(func(_2092_i7 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('w')
					}))).Elements()); ; {
						_compr_80, _ok83 := _iter83()
						if !_ok83 {
							break
						}
						var _2093_v127 _dafny.CodePoint
						_2093_v127 = interface{}(_compr_80).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(203))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg115 _dafny.Int) interface{} {
								return coer115(arg115)
							}
						}(func(_2092_i7 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('w')
						})), _2093_v127) {
							_coll80.Add(_2093_v127, (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool))
						}
					}
					return _coll80.ToMap()
				}(), _2090_v126)
				var _2094_v129 *C6
				_ = _2094_v129
				var _nw334 *C6 = New_C6_()
				_ = _nw334
				_nw334.Ctor__(_2080_v122, (_2077_cf51).F12(), Companion_Default___.SafeModuloInt(_2086_cf6, (_2091_v128).Cardinality()))
				_2094_v129 = _nw334
			} else if _source31.Is_DC5() {
				var _2095_v130 _dafny.Array
				_ = _2095_v130
				var _nw335 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw335
				_2095_v130 = _nw335
				var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_2095_v130), 0))
				_ = _index287
				(_2095_v130).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1934_v33).Cardinality()), _dafny.IntOfInt64(483)), (_index287).Int())
				var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_2095_v130), 0))
				_ = _index288
				(_2095_v130).ArraySet1((_dafny.Zero).Minus((_2077_cf51).F11()), (_index288).Int())
				var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_2095_v130), 0))
				_ = _index289
				var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_2095_v130), 0))
				_ = _index290
				var _rhs360 _dafny.Int = (_1932_v31).Select((Companion_Default___.SafeIndex(_2074_cf54, _dafny.IntOfUint32((_1932_v31).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs360
				var _rhs361 _dafny.Int = Companion_Default___.SafeDivisionInt(((_this).F11()).Plus((_this).F11()), ((_2077_cf51).F11()).Times(_dafny.IntOfUint32((_1934_v33).Cardinality())))
				_ = _rhs361
				var _rhs362 _dafny.Int = _dafny.IntOfInt64(518)
				_ = _rhs362
				var _lhs306 _dafny.Array = _2095_v130
				_ = _lhs306
				var _lhs307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_2095_v130), 0))
				_ = _lhs307
				var _lhs308 _dafny.Array = _2095_v130
				_ = _lhs308
				var _lhs309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_2095_v130), 0))
				_ = _lhs309
				(_lhs306).ArraySet1(_rhs360, (_lhs307).Int())
				(_lhs308).ArraySet1(_rhs361, (_lhs309).Int())
				_2074_cf54 = _rhs362
				(globalState).F1 = (_2095_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_2095_v130), 0))).Int()).(_dafny.Int)
				var _2096_v131 _dafny.Sequence
				_ = _2096_v131
				_2096_v131 = _dafny.SeqOf(_2095_v130, _2095_v130)
				var _2097_v132 _dafny.Array
				_ = _2097_v132
				var _nwElement0_93 _dafny.Array = _2095_v130
				_ = _nwElement0_93
				var _nw336 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_93, nil, _dafny.IntOfInt64(23))
				_ = _nw336
				(_nw336).ArraySet1(_nwElement0_93, 0)
				(_nw336).ArraySet1(_2095_v130, 1)
				(_nw336).ArraySet1(_2095_v130, 2)
				(_nw336).ArraySet1(_2095_v130, 3)
				(_nw336).ArraySet1(_2095_v130, 4)
				(_nw336).ArraySet1(_2095_v130, 5)
				(_nw336).ArraySet1(_2095_v130, 6)
				(_nw336).ArraySet1(_2095_v130, 7)
				(_nw336).ArraySet1(_2095_v130, 8)
				(_nw336).ArraySet1(_2095_v130, 9)
				(_nw336).ArraySet1((_2096_v131).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_2096_v131).Cardinality()))).Uint32()).(_dafny.Array), 10)
				(_nw336).ArraySet1(_2095_v130, 11)
				(_nw336).ArraySet1(_2095_v130, 12)
				(_nw336).ArraySet1(_2095_v130, 13)
				(_nw336).ArraySet1(_2095_v130, 14)
				(_nw336).ArraySet1(_2095_v130, 15)
				(_nw336).ArraySet1(_2095_v130, 16)
				(_nw336).ArraySet1(_2095_v130, 17)
				(_nw336).ArraySet1(_2095_v130, 18)
				(_nw336).ArraySet1(_2095_v130, 19)
				(_nw336).ArraySet1(_2095_v130, 20)
				(_nw336).ArraySet1(_2095_v130, 21)
				(_nw336).ArraySet1(_2095_v130, 22)
				_2097_v132 = _nw336
				var _2098_v133 _dafny.Array
				_ = _2098_v133
				_2098_v133 = _2097_v132
				_2098_v133 = _2098_v133
				var _2099_v134 *C6
				_ = _2099_v134
				var _nw337 *C6 = New_C6_()
				_ = _nw337
				_nw337.Ctor__(Companion_D2_.Create_DC4_((_this).Fm8(Companion_Default___.Fm39((_2095_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_2095_v130), 0))).Int()).(_dafny.Int), (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), (_2077_cf51).F11(), globalState), globalState), (_2095_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_2095_v130), 0))).Int()).(_dafny.Int), (_this).F11()), (_this).F12(), _dafny.IntOfUint32((_1934_v33).Cardinality()))
				_2099_v134 = _nw337
			} else {
				var _2100___mcc_h25 _dafny.MultiSet = _source31.Get_().(D2_DC3).Cf5
				_ = _2100___mcc_h25
				var _2101_cf5 _dafny.MultiSet = _2100___mcc_h25
				_ = _2101_cf5
				var _2102_v135 _dafny.Array
				_ = _2102_v135
				var _nwElement0_94 _dafny.Int = ((_2077_cf51).F11()).Plus((_this).F11())
				_ = _nwElement0_94
				var _nw338 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_94, nil, _dafny.IntOfInt64(5))
				_ = _nw338
				(_nw338).ArraySet1(_nwElement0_94, 0)
				(_nw338).ArraySet1(_dafny.IntOfInt64(289), 1)
				(_nw338).ArraySet1((_2077_cf51).F11(), 2)
				(_nw338).ArraySet1(_dafny.IntOfInt64(-902), 3)
				(_nw338).ArraySet1((_this).F11(), 4)
				_2102_v135 = _nw338
				var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_2102_v135), 0))
				_ = _index291
				(_2102_v135).ArraySet1((_this).F11(), (_index291).Int())
				var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_2102_v135), 0))
				_ = _index292
				var _rhs363 _dafny.Array = _1938_v37
				_ = _rhs363
				var _rhs364 _dafny.Int = _dafny.IntOfInt64(976)
				_ = _rhs364
				var _rhs365 _dafny.Int = (_2077_cf51).F11()
				_ = _rhs365
				var _lhs310 _dafny.Array = _2102_v135
				_ = _lhs310
				var _lhs311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_2102_v135), 0))
				_ = _lhs311
				var _lhs312 *GlobalState = globalState
				_ = _lhs312
				_1938_v37 = _rhs363
				(_lhs310).ArraySet1(_rhs364, (_lhs311).Int())
				_lhs312.F1 = _rhs365
				var _2103_v136 _dafny.MultiSet
				_ = _2103_v136
				_2103_v136 = _dafny.MultiSetOf((_this).F13(), _2078_cf50)
				(globalState).F1 = (func() _dafny.Int {
					if !(_2103_v136).Contains(_2078_cf50) {
						return Companion_Default___.SafeModuloInt((_this).F11(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vxxic")).Cardinality()))
					}
					return (_this).F11()
				})()
				_2076_cf52 = ((_this).F11()).Cmp(_2074_cf54) >= 0
				var _2104_v137 _dafny.Map
				_ = _2104_v137
				_2104_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).Fm4((_2077_cf51).F11(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1941_v41, _1950_v49), (_dafny.Zero).Minus(Companion_Default___.Fm2(_dafny.IntOfInt64(-392), (_this).F11(), globalState)), (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), globalState)), (func() _dafny.Array {
					if _2078_cf50 {
						return _2102_v135
					}
					return _2102_v135
				})())
				_2104_v137 = (_2104_v137).Update((_this).Fm4((_2101_cf5).Cardinality(), func() _dafny.Map {
					var _coll81 = _dafny.NewMapBuilder()
					_ = _coll81
					for _iter84 := _dafny.Iterate((_dafny.SetOf(_1941_v41)).Elements()); ; {
						_compr_81, _ok84 := _iter84()
						if !_ok84 {
							break
						}
						var _2105_v138 _dafny.MultiSet
						_2105_v138 = interface{}(_compr_81).(_dafny.MultiSet)
						if (_dafny.SetOf(_1941_v41)).Contains(_2105_v138) {
							_coll81.Add(_2105_v138, _1950_v49)
						}
					}
					return _coll81.ToMap()
				}(), _dafny.IntOfInt64(977), (_this).F13(), globalState), _2102_v135)
			}
			(globalState).F7 = (_this).Fm4((_this).F11(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1941_v41).Update((_2077_cf51).F11(), Companion_Default___.Abs((_2077_cf51).F11())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2077_cf51).F11(), (_2077_cf51).F11()))).Merge(_1951_v50), (_2080_v122).Dtor_cf7(), false, globalState)
			var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
			_ = _index293
			(_1938_v37).ArraySet1((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), (_index293).Int())
		} else if _source29.Is_DC36() {
			var _2106___mcc_h10 _dafny.Int = _source29.Get_().(D15_DC36).Cf55
			_ = _2106___mcc_h10
			var _2107___mcc_h11 bool = _source29.Get_().(D15_DC36).Cf56
			_ = _2107___mcc_h11
			var _2108___mcc_h12 bool = _source29.Get_().(D15_DC36).Cf57
			_ = _2108___mcc_h12
			var _2109___mcc_h13 bool = _source29.Get_().(D15_DC36).Cf58
			_ = _2109___mcc_h13
			var _2110___mcc_h14 _dafny.Int = _source29.Get_().(D15_DC36).Cf59
			_ = _2110___mcc_h14
			var _2111_cf59 _dafny.Int = _2110___mcc_h14
			_ = _2111_cf59
			var _2112_cf58 bool = _2109___mcc_h13
			_ = _2112_cf58
			var _2113_cf57 bool = _2108___mcc_h12
			_ = _2113_cf57
			var _2114_cf56 bool = _2107___mcc_h11
			_ = _2114_cf56
			var _2115_cf55 _dafny.Int = _2106___mcc_h10
			_ = _2115_cf55
			if _2112_cf58 {
				(globalState).F7 = Companion_Default___.Fm0(globalState)
				_1934_v33 = _dafny.Companion_Sequence_.Concatenate(_1934_v33, _1934_v33)
				_2112_cf58 = _2114_cf56
				var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
				_ = _index294
				(_1938_v37).ArraySet1((func() bool {
					if _2114_cf56 {
						return _2112_cf58
					}
					return (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool)
				})(), (_index294).Int())
				var _2116_v139 D8
				_ = _2116_v139
				_2116_v139 = Companion_D8_.Create_DC16_(_2114_cf56, _2115_cf55)
				var _2117_v140 D8
				_ = _2117_v140
				_2117_v140 = Companion_D8_.Create_DC17_(_2116_v139)
				_2117_v140 = _2117_v140
			} else {
				var _2118_v141 _dafny.Array
				_ = _2118_v141
				var _nw339 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw339
				_2118_v141 = _nw339
				var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_2118_v141), 0))
				_ = _index295
				(_2118_v141).ArraySet1((_this).F11(), (_index295).Int())
				var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_2118_v141), 0))
				_ = _index296
				(_2118_v141).ArraySet1(_dafny.IntOfInt64(-571), (_index296).Int())
				(globalState).F1 = _2111_cf59
				var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
				_ = _index297
				(_1938_v37).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm19((_this).F13(), false, globalState), _dafny.Companion_Sequence_.Concatenate(_1956_v55, _1956_v55)), (_index297).Int())
				var _2119_v142 *C2
				_ = _2119_v142
				var _nw340 *C2 = New_C2_()
				_ = _nw340
				_nw340.Ctor__(p0, (_this).F11())
				_2119_v142 = _nw340
				_1941_v41 = (_dafny.MultiSetOf((_this).F11(), _2111_cf59, _2111_cf59)).Union(_1941_v41)
			}
			var _2120_v143 _dafny.Map
			_ = _2120_v143
			_2120_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _1934_v33)
			var _2121_v144 _dafny.Sequence
			_ = _2121_v144
			_2121_v144 = _dafny.SeqOf(Companion_Default___.Fm22((func() _dafny.Sequence {
				if (_2120_v143).Contains(true) {
					return (_2120_v143).Get(true).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("gau")
			})(), _2112_cf58, globalState), _1940_v39)
			var _2122_v145 D8
			_ = _2122_v145
			_2122_v145 = Companion_D8_.Create_DC16_(_2113_cf57, (func() _dafny.Int {
				if (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool) {
					return _dafny.IntOfUint32((_2121_v144).Cardinality())
				}
				return (_this).F11()
			})())
			var _2123_v146 _dafny.Map
			_ = _2123_v146
			_2123_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1949_v48).Contains((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool)), _2112_cf58)
			var _2124_v147 T1
			_ = _2124_v147
			var _nw341 *C3 = New_C3_()
			_ = _nw341
			_nw341.Ctor__(_2115_cf55, _2115_cf55, (_this).F11(), (_this).F12())
			_2124_v147 = _nw341
			var _2125_v148 _dafny.Set
			_ = _2125_v148
			_2125_v148 = _dafny.SetOf(_2124_v147)
			var _2126_v149 _dafny.Sequence
			_ = _2126_v149
			_2126_v149 = _dafny.SeqOf(_2125_v148, _2125_v148, _2125_v148, _2125_v148)
			var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
			_ = _index298
			var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
			_ = _index299
			var _rhs366 bool = (_1956_v55).Select((Companion_Default___.SafeIndex((_2111_cf59).Times(_2115_cf55), _dafny.IntOfUint32((_1956_v55).Cardinality()))).Uint32()).(bool)
			_ = _rhs366
			var _rhs367 _dafny.Int = (_1947_v46).Fm3(((_this).F11()).Plus((_this).F11()), _2112_cf58, !((func() bool {
				if (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool) {
					return true
				}
				return (_this).F13()
			})()), globalState)
			_ = _rhs367
			var _rhs368 bool = (func() bool {
				if _2112_cf58 {
					return (_this).F13()
				}
				return (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool)
			})()
			_ = _rhs368
			var _rhs369 D8 = Companion_D8_.Create_DC16_((_this).F13(), (_this).F11())
			_ = _rhs369
			var _rhs370 bool = (func() bool {
				if (_2123_v146).Contains(_2113_cf57) {
					return (_2123_v146).Get(_2113_cf57).(bool)
				}
				return _dafny.Companion_Sequence_.IsProperPrefixOf(_2126_v149, _2126_v149)
			})()
			_ = _rhs370
			var _lhs313 _dafny.Array = _1938_v37
			_ = _lhs313
			var _lhs314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
			_ = _lhs314
			var _lhs315 *GlobalState = globalState
			_ = _lhs315
			var _lhs316 _dafny.Array = _1938_v37
			_ = _lhs316
			var _lhs317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))
			_ = _lhs317
			(_lhs313).ArraySet1(_rhs366, (_lhs314).Int())
			_lhs315.F1 = _rhs367
			(_lhs316).ArraySet1(_rhs368, (_lhs317).Int())
			_2122_v145 = _rhs369
			_2113_cf57 = _rhs370
			var _2127_v150 _dafny.Sequence
			_ = _2127_v150
			_2127_v150 = _dafny.SeqOf(_dafny.MultiSetOf(_2123_v146))
			var _2128_v151 _dafny.Map
			_ = _2128_v151
			_2128_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1934_v33, (_2127_v150).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_2127_v150).Cardinality()))).Uint32()).(_dafny.MultiSet))
			var _2129_v152 _dafny.MultiSet
			_ = _2129_v152
			_2129_v152 = _dafny.MultiSetOf(_2123_v146, _2123_v146)
			_2128_v151 = (_2128_v151).Update(_dafny.UnicodeSeqOfUtf8Bytes("fq"), (func() _dafny.MultiSet {
				if _2114_cf56 {
					return _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2113_cf57, (_this).F13()))
				}
				return _2129_v152
			})())
			_2115_cf55 = (_2124_v147).Fm3(((_2124_v147).F12()).Cardinality(), _2114_cf56, (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), globalState)
		} else if _source29.Is_DC33() {
			var _2130___mcc_h15 _dafny.Map = _source29.Get_().(D15_DC33).Cf49
			_ = _2130___mcc_h15
			var _2131_cf49 _dafny.Map = _2130___mcc_h15
			_ = _2131_cf49
			_1949_v48 = (_1949_v48).Update(false, (_this).F11())
			if ((_this).F11()).Cmp((_this).F11()) >= 0 {
				var _2132_v153 D21
				_ = _2132_v153
				_2132_v153 = Companion_D21_.Create_DC48_((_this).F13())
				var _2133_v154 _dafny.Map
				_ = _2133_v154
				_2133_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let54_0 D21) D21 {
					return func(_2134_dt__update__tmp_h5 D21) D21 {
						return func(_pat_let55_0 bool) D21 {
							return func(_2135_dt__update_hcf78_h0 bool) D21 {
								return Companion_D21_.Create_DC48_(_2135_dt__update_hcf78_h0)
							}(_pat_let55_0)
						}((_this).F13())
					}(_pat_let54_0)
				}(_2132_v153), (_this).F13())
				(globalState).F1 = Companion_Default___.SafeModuloInt(((_2133_v154).Cardinality()).Times(_dafny.IntOfInt64(-913)), ((_this).F11()).Plus(_dafny.IntOfInt64(391)))
				_1950_v49 = (_1950_v49).Merge((_1950_v49).Merge(_1950_v49))
				_1957_v56 = _1957_v56
				(globalState).F1 = (_this).F11()
				var _2136_v155 _dafny.Int
				_ = _2136_v155
				var _2137_v156 bool
				_ = _2137_v156
				var _out35 _dafny.Int
				_ = _out35
				var _out36 bool
				_ = _out36
				_out35, _out36 = (_this).M2(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(222), (_this).F11()), globalState)
				_2136_v155 = _out35
				_2137_v156 = _out36
			} else {
				(globalState).F1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uoj"), _1934_v33)).Cardinality())
				(globalState).F1 = _dafny.IntOfInt64(328)
				(globalState).F1 = _dafny.IntOfInt64(632)
				var _2138_v157 _dafny.Array
				_ = _2138_v157
				var _nw342 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(5))
				_ = _nw342
				_2138_v157 = _nw342
				var _2139_v158 _dafny.MultiSet
				_ = _2139_v158
				_2139_v158 = _dafny.MultiSetOf((_this).F13())
				var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_2138_v157), 0))
				_ = _index300
				(_2138_v157).ArraySet1(_2139_v158, (_index300).Int())
				var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_2138_v157), 0))
				_ = _index301
				(_2138_v157).ArraySet1(_2139_v158, (_index301).Int())
				var _2140_v159 _dafny.Array
				_ = _2140_v159
				var _nw343 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
				_ = _nw343
				_2140_v159 = _nw343
				var _2141_v160 _dafny.Array
				_ = _2141_v160
				var _nwElement0_95 _dafny.Int = (_this).F11()
				_ = _nwElement0_95
				var _nw344 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_95, nil, _dafny.IntOfInt64(9))
				_ = _nw344
				(_nw344).ArraySet1(_nwElement0_95, 0)
				(_nw344).ArraySet1(_dafny.IntOfInt64(283), 1)
				(_nw344).ArraySet1((_this).F11(), 2)
				(_nw344).ArraySet1((_this).F11(), 3)
				(_nw344).ArraySet1((_this).F11(), 4)
				(_nw344).ArraySet1((_this).F11(), 5)
				(_nw344).ArraySet1((_this).F11(), 6)
				(_nw344).ArraySet1((_this).F11(), 7)
				(_nw344).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dlq")).Cardinality()), 8)
				_2141_v160 = _nw344
				var _2142_v161 _dafny.Map
				_ = _2142_v161
				_2142_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2140_v159, _2141_v160)
				var _2143_v162 D11
				_ = _2143_v162
				_2143_v162 = Companion_D11_.Create_DC22_(_2142_v161)
				var _2144_v163 _dafny.MultiSet
				_ = _2144_v163
				_2144_v163 = _dafny.MultiSetOf(_2143_v162)
				var _2145_v164 D7
				_ = _2145_v164
				_2145_v164 = Companion_D7_.Create_DC13_((_this).F12(), (_2144_v163).Cardinality(), (_this).F11())
				var _2146_v165 *C7
				_ = _2146_v165
				var _nw345 *C7 = New_C7_()
				_ = _nw345
				_nw345.Ctor__((_1947_v46).F14(), (_2145_v164).Dtor_cf18(), (_this).F11())
				_2146_v165 = _nw345
			}
			var _2147_v166 _dafny.Map
			_ = _2147_v166
			_2147_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(319), _1940_v39)
			var _2148_v168 _dafny.Sequence
			_ = _2148_v168
			_2148_v168 = _dafny.SeqOf(_2147_v166, func() _dafny.Map {
				var _coll82 = _dafny.NewMapBuilder()
				_ = _coll82
				for _iter85 := _dafny.Iterate((_1932_v31).Elements()); ; {
					_compr_82, _ok85 := _iter85()
					if !_ok85 {
						break
					}
					var _2149_v167 _dafny.Int
					_2149_v167 = interface{}(_compr_82).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_1932_v31, _2149_v167) {
						_coll82.Add((_2149_v167).Plus((_this).F11()), _1940_v39)
					}
				}
				return _coll82.ToMap()
			}())
			_2147_v166 = (_2147_v166).Merge((_2148_v168).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_2148_v168).Cardinality()))).Uint32()).(_dafny.Map))
			var _2150_v169 _dafny.Sequence
			_ = _2150_v169
			_2150_v169 = _dafny.SeqOf(_1938_v37, _1938_v37, _1935_v34)
			var _rhs371 _dafny.Sequence = _2150_v169
			_ = _rhs371
			var _rhs372 bool = (_this).F13()
			_ = _rhs372
			var _lhs318 *GlobalState = globalState
			_ = _lhs318
			_2150_v169 = _rhs371
			_lhs318.F7 = _rhs372
		} else {
			var _2151___mcc_h16 D15 = _source29.Get_().(D15_DC37).Cf60
			_ = _2151___mcc_h16
			var _2152_cf60 D15 = _2151___mcc_h16
			_ = _2152_cf60
			(globalState).F1 = (func() _dafny.Int {
				if (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool) {
					return (_this).F11()
				}
				return Companion_Default___.SafeDivisionInt((_this).F11(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1934_v33, (Companion_Default___.SafeIndex(Companion_Default___.Fm2((_this).Fm6((_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool), _dafny.IntOfInt64(240), globalState), (_this).F11(), globalState), _dafny.IntOfUint32((_1934_v33).Cardinality()))).Uint32(), p0)).Cardinality()))
			})()
			(globalState).F1 = _dafny.IntOfUint32((_1934_v33).Cardinality())
			var _2153_v170 _dafny.Int
			_ = _2153_v170
			var _2154_v171 bool
			_ = _2154_v171
			var _out37 _dafny.Int
			_ = _out37
			var _out38 bool
			_ = _out38
			_out37, _out38 = (_this).M2(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1934_v33, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_1934_v33).Cardinality()))).Uint32(), p0)).Cardinality()), globalState)
			_2153_v170 = _out37
			_2154_v171 = _out38
			var _2155_v172 _dafny.Map
			_ = _2155_v172
			_2155_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1938_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_1938_v37), 0))).Int()).(bool))
			_2155_v172 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2154_v171)).Merge(_2155_v172)
		}
		r0 = Companion_Default___.Fm37(globalState)
		return r0
	}
}
func (_this *C8) M1(p0 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Map) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _2156_v0 _dafny.Array
		_ = _2156_v0
		var _len0_40 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_40
		var _nw346 _dafny.Array
		_ = _nw346
		if _len0_40.Cmp(_dafny.Zero) == 0 {
			_nw346 = _dafny.NewArray(_len0_40)
		} else {
			var _init40 func(_dafny.Int) _dafny.Int = func(_2157_i0 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_2157_i0, (_this).F11())
			}
			_ = _init40
			var _element0_40 = _init40(_dafny.Zero)
			_ = _element0_40
			_nw346 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
			(_nw346).ArraySet1(_element0_40, 0)
			var _nativeLen0_40 = (_len0_40).Int()
			_ = _nativeLen0_40
			for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
				(_nw346).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
			}
		}
		_2156_v0 = _nw346
		var _2158_v1 _dafny.CodePoint
		_ = _2158_v1
		_2158_v1 = _dafny.CodePoint('j')
		var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))
		_ = _index302
		(_2156_v0).ArraySet1((_dafny.Zero).Minus((_this).Fm8(_2158_v1, globalState)), (_index302).Int())
		var _2159_v2 _dafny.Sequence
		_ = _2159_v2
		_2159_v2 = _dafny.UnicodeSeqOfUtf8Bytes("h")
		var _2160_v3 _dafny.Sequence
		_ = _2160_v3
		_2160_v3 = _dafny.SeqOf((_this).F13())
		var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))
		_ = _index303
		(_2156_v0).ArraySet1((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2159_v2, _2159_v2)).Cardinality())).Plus(_dafny.IntOfUint32((_2160_v3).Cardinality())), (_index303).Int())
		_2158_v1 = _2158_v1
		var _2161_v4 _dafny.MultiSet
		_ = _2161_v4
		_2161_v4 = _dafny.MultiSetOf((_2160_v3).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_2160_v3).Cardinality()))).Uint32()).(bool))
		var _2162_v5 _dafny.Map
		_ = _2162_v5
		_2162_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), (_this).Fm7((_dafny.Zero).Minus((_2161_v4).Cardinality()), (_this).F13(), _2161_v4, globalState))
		var _hi5 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-346))).Uint32(), func(coer116 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg116 _dafny.Int) interface{} {
				return coer116(arg116)
			}
		}((func(_2163_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2164_i2 _dafny.Int) _dafny.CodePoint {
				return _2163_v1
			}
		})(_2158_v1)))).Cardinality())
		_ = _hi5
		for _2165_i1 := (func() _dafny.Int {
			if (_2162_v5).Contains((_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int)) {
				return (_2162_v5).Get((_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int)).(_dafny.Int)
			}
			return p0
		})(); _2165_i1.Cmp(_hi5) < 0; _2165_i1 = _2165_i1.Plus(_dafny.One) {
			var _2166_v6 _dafny.MultiSet
			_ = _2166_v6
			_2166_v6 = _dafny.MultiSetOf(_2165_i1, (_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int))
			(globalState).F1 = ((_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int)).Plus((_dafny.Zero).Minus(((_this).F11()).Times((_this).Fm3(_dafny.IntOfInt64(-774), (_this).Fm4(_2165_i1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2166_v6, func() _dafny.Map {
				var _coll83 = _dafny.NewMapBuilder()
				_ = _coll83
				for _iter86 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(163), _dafny.IntOfInt64(568))); ; {
					_compr_83, _ok86 := _iter86()
					if !_ok86 {
						break
					}
					var _2167_v7 _dafny.Int
					_2167_v7 = interface{}(_compr_83).(_dafny.Int)
					if ((_dafny.IntOfInt64(163)).Cmp(_2167_v7) <= 0) && ((_2167_v7).Cmp(_dafny.IntOfInt64(568)) < 0) {
						_coll83.Add((_2167_v7).Times(_dafny.IntOfInt64(252)), (_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int))
					}
				}
				return _coll83.ToMap()
			}()), (_this).F11(), (_this).F13(), globalState), (_this).F13(), globalState))))
			(globalState).F1 = _dafny.IntOfInt64(-211)
			var _2168_v8 _dafny.Map
			_ = _2168_v8
			_2168_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _2158_v1)
			if !((_2168_v8).Equals(_2168_v8)) || (((_this).F12()).IsSubsetOf((_this).F12())) {
				var _2169_v9 *C1
				_ = _2169_v9
				var _nw347 *C1 = New_C1_()
				_ = _nw347
				_nw347.Ctor__((_this).F13(), (_this).F13())
				_2169_v9 = _nw347
				var _2170_v10 _dafny.Map
				_ = _2170_v10
				_2170_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2169_v9, (_2169_v9).F18())
				_2170_v10 = (_2170_v10).Update(_2169_v9, (_this).F13())
				var _2171_v11 _dafny.Array
				_ = _2171_v11
				var _nw348 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(20))
				_ = _nw348
				_2171_v11 = _nw348
				_2171_v11 = _2171_v11
				var _2172_v12 *C2
				_ = _2172_v12
				var _nw349 *C2 = New_C2_()
				_ = _nw349
				_nw349.Ctor__(_2158_v1, ((_dafny.Zero).Minus(p0)).Minus((_dafny.Zero).Minus(p0)))
				_2172_v12 = _nw349
				var _2173_v13 _dafny.Map
				_ = _2173_v13
				_2173_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2169_v9).F18(), _2166_v6)
				var _2174_v14 _dafny.Map
				_ = _2174_v14
				_2174_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2173_v13).Cardinality(), !((_2169_v9).F19()))
				var _rhs373 bool = false
				_ = _rhs373
				var _rhs374 _dafny.Map = (_2174_v14).Merge(func() _dafny.Map {
					var _coll84 = _dafny.NewMapBuilder()
					_ = _coll84
					for _iter87 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(923), _dafny.IntOfInt64(-68))); ; {
						_compr_84, _ok87 := _iter87()
						if !_ok87 {
							break
						}
						var _2175_v15 _dafny.Int
						_2175_v15 = interface{}(_compr_84).(_dafny.Int)
						if ((_dafny.IntOfInt64(923)).Cmp(_2175_v15) <= 0) && ((_2175_v15).Cmp(_dafny.IntOfInt64(-68)) < 0) {
							_coll84.Add(Companion_Default___.SafeModuloInt(_2175_v15, p0), false)
						}
					}
					return _coll84.ToMap()
				}())
				_ = _rhs374
				var _lhs319 *GlobalState = globalState
				_ = _lhs319
				_lhs319.F7 = _rhs373
				_2174_v14 = _rhs374
				var _2176_v16 _dafny.Sequence
				_ = _2176_v16
				_2176_v16 = _dafny.SeqOf(_2162_v5, _2162_v5)
				var _2177_v17 _dafny.Set
				_ = _2177_v17
				_2177_v17 = _dafny.SetOf(p0, (_this).F11())
				var _2178_v18 _dafny.Map
				_ = _2178_v18
				_2178_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2176_v16).Select((Companion_Default___.SafeIndex((_2177_v17).Cardinality(), _dafny.IntOfUint32((_2176_v16).Cardinality()))).Uint32()).(_dafny.Map), (_2172_v12).F24())
				_2178_v18 = _2178_v18
			} else {
				(globalState).F6 = !_dafny.Companion_Sequence_.Equal(_2159_v2, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_2159_v2, (Companion_Default___.SafeIndex(_2165_i1, _dafny.IntOfUint32((_2159_v2).Cardinality()))).Uint32(), _2158_v1), _2159_v2))
				var _2179_v19 _dafny.Array
				_ = _2179_v19
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_41
				var _nw350 _dafny.Array
				_ = _nw350
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw350 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) bool = func(_2180_i3 _dafny.Int) bool {
						return ((_this).F13()) && (!((_this).F13()))
					}
					_ = _init41
					var _element0_41 = _init41(_dafny.Zero)
					_ = _element0_41
					_nw350 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
					(_nw350).ArraySet1(_element0_41, 0)
					var _nativeLen0_41 = (_len0_41).Int()
					_ = _nativeLen0_41
					for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
						(_nw350).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
					}
				}
				_2179_v19 = _nw350
				_2179_v19 = _2179_v19
				(globalState).F6 = ((_this).F12()).IsProperSubsetOf(((_this).F12()).Intersection(_dafny.SetOf((_this).F13(), (_this).F13())))
				_2159_v2 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(175))).Uint32(), func(coer117 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg117 _dafny.Int) interface{} {
						return coer117(arg117)
					}
				}((func(_2181_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2182_i4 _dafny.Int) _dafny.CodePoint {
						return _2181_v1
					}
				})(_2158_v1)))
				(globalState).F6 = !((_this).F13()) || ((_this).F13())
			}
			var _2183_v20 _dafny.Map
			_ = _2183_v20
			_2183_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2158_v1, true)
			var _2184_v21 _dafny.Sequence
			_ = _2184_v21
			_2184_v21 = _dafny.SeqOf(_2183_v20)
			var _rhs375 bool = (_this).F13()
			_ = _rhs375
			var _rhs376 bool = !(!((_this).F13()) || ((_this).F13()))
			_ = _rhs376
			var _rhs377 _dafny.Sequence = _2184_v21
			_ = _rhs377
			var _rhs378 bool = ((_this).F13()) == (!((func() bool {
				if (_this).F13() {
					return (_this).F13()
				}
				return (_this).F13()
			})()))
			_ = _rhs378
			var _rhs379 bool = (((_this).F11()).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if (_2166_v6).Contains(_dafny.IntOfUint32((_2159_v2).Cardinality())) {
					return (_2166_v6).Multiplicity(_dafny.IntOfUint32((_2159_v2).Cardinality()))
				}
				return (_this).F11()
			})()))).Cmp((_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int)) >= 0
			_ = _rhs379
			var _lhs320 *GlobalState = globalState
			_ = _lhs320
			var _lhs321 *GlobalState = globalState
			_ = _lhs321
			var _lhs322 *GlobalState = globalState
			_ = _lhs322
			var _lhs323 *GlobalState = globalState
			_ = _lhs323
			_lhs320.F7 = _rhs375
			_lhs321.F7 = _rhs376
			_2184_v21 = _rhs377
			_lhs322.F6 = _rhs378
			_lhs323.F6 = _rhs379
		}
		var _hi6 _dafny.Int = (_this).F11()
		_ = _hi6
		for _2185_i5 := (_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int); _2185_i5.Cmp(_hi6) < 0; _2185_i5 = _2185_i5.Plus(_dafny.One) {
			var _2186_v22 _dafny.Sequence
			_ = _2186_v22
			_2186_v22 = _dafny.SeqOf(_dafny.IntOfUint32((_2160_v3).Cardinality()))
			var _rhs380 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(291))).Uint32(), func(coer118 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg118 _dafny.Int) interface{} {
					return coer118(arg118)
				}
			}((func(_2187_i5 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2188_i6 _dafny.Int) _dafny.Int {
					return (_2187_i5).Minus((_this).F11())
				}
			})(_2185_i5)))
			_ = _rhs380
			var _rhs381 _dafny.Int = (_2185_i5).Minus(_2185_i5)
			_ = _rhs381
			var _lhs324 *GlobalState = globalState
			_ = _lhs324
			_2186_v22 = _rhs380
			_lhs324.F1 = _rhs381
			var _2189_v23 _dafny.Array
			_ = _2189_v23
			var _nw351 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw351
			_2189_v23 = _nw351
			var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_2189_v23), 0))
			_ = _index304
			(_2189_v23).ArraySet1(_2156_v0, (_index304).Int())
			var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_2189_v23), 0))
			_ = _index305
			(_2189_v23).ArraySet1(_2156_v0, (_index305).Int())
			var _2190_v24 _dafny.MultiSet
			_ = _2190_v24
			_2190_v24 = _dafny.MultiSetOf((_this).F11())
			var _2191_v25 _dafny.Map
			_ = _2191_v25
			_2191_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2190_v24, (_dafny.Zero).Minus((_this).F11()))
			_2191_v25 = (_2191_v25).Update(_2190_v24, Companion_Default___.Fm2((_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int), (_this).F11(), globalState))
			var _2192_v26 _dafny.Map
			_ = _2192_v26
			_2192_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_2185_i5, _dafny.IntOfUint32((Companion_Default___.Fm19((_this).F13(), (_this).F13(), globalState)).Cardinality()))), _2159_v2)
			_2159_v2 = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_2192_v26).Contains((_dafny.Zero).Minus((_this).F11())) {
					return (_2192_v26).Get((_dafny.Zero).Minus((_this).F11())).(_dafny.Sequence)
				}
				return _2159_v2
			})(), (Companion_Default___.SafeIndex(_2185_i5, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_2192_v26).Contains((_dafny.Zero).Minus((_this).F11())) {
					return (_2192_v26).Get((_dafny.Zero).Minus((_this).F11())).(_dafny.Sequence)
				}
				return _2159_v2
			})()).Cardinality()))).Uint32(), _2158_v1)
		}
		var _hi7 _dafny.Int = p0
		_ = _hi7
		for _2193_i7 := (_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int); _2193_i7.Cmp(_hi7) < 0; _2193_i7 = _2193_i7.Plus(_dafny.One) {
			var _2194_v27 _dafny.Map
			_ = _2194_v27
			_2194_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F13()), (_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int))
			if ((_2160_v3).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2160_v3).Cardinality()))).Uint32()).(bool)) == ((_2194_v27).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_dafny.Zero).Minus(_2193_i7)))) {
				var _2195_v28 *C5
				_ = _2195_v28
				var _nw352 *C5 = New_C5_()
				_ = _nw352
				_nw352.Ctor__()
				_2195_v28 = _nw352
				(globalState).F7 = (_this).F13()
				var _2196_v29 _dafny.Array
				_ = _2196_v29
				var _nw353 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
				_ = _nw353
				_2196_v29 = _nw353
				var _2197_v30 _dafny.Sequence
				_ = _2197_v30
				_2197_v30 = _dafny.SeqOf((_this).F11())
				var _2198_v31 _dafny.Map
				_ = _2198_v31
				_2198_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F13())
				var _2199_v32 D6
				_ = _2199_v32
				_2199_v32 = Companion_D6_.Create_DC10_(_2160_v3)
				var _2200_v33 D19
				_ = _2200_v33
				_2200_v33 = Companion_D19_.Create_DC43_(false, _dafny.CodePoint('r'), (_this).F13(), _2199_v32, _2159_v2)
				var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))
				_ = _index306
				var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))
				_ = _index307
				var _rhs382 _dafny.Array = _2196_v29
				_ = _rhs382
				var _rhs383 _dafny.Int = (_2197_v30).Select((Companion_Default___.SafeIndex((_2198_v31).Cardinality(), _dafny.IntOfUint32((_2197_v30).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs383
				var _rhs384 _dafny.Int = (func() _dafny.Int {
					if (_2200_v33).Dtor_cf69() {
						return p0
					}
					return (_this).F11()
				})()
				_ = _rhs384
				var _rhs385 bool = ((_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int)).Cmp((_this).F11()) != 0
				_ = _rhs385
				var _lhs325 _dafny.Array = _2156_v0
				_ = _lhs325
				var _lhs326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))
				_ = _lhs326
				var _lhs327 _dafny.Array = _2156_v0
				_ = _lhs327
				var _lhs328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))
				_ = _lhs328
				var _lhs329 *GlobalState = globalState
				_ = _lhs329
				_2196_v29 = _rhs382
				(_lhs325).ArraySet1(_rhs383, (_lhs326).Int())
				(_lhs327).ArraySet1(_rhs384, (_lhs328).Int())
				_lhs329.F7 = _rhs385
				var _2201_v34 _dafny.Set
				_ = _2201_v34
				_2201_v34 = _dafny.SetOf((_this).F11(), _2193_i7)
				var _2202_v35 _dafny.MultiSet
				_ = _2202_v35
				_2202_v35 = _dafny.MultiSetOf(_2201_v34, _2201_v34)
				var _2203_v36 _dafny.Map
				_ = _2203_v36
				_2203_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_2202_v35).Contains(_2201_v34) {
						return (_2202_v35).Multiplicity(_2201_v34)
					}
					return p0
				})(), (_this).F13())
				var _2204_v37 *C3
				_ = _2204_v37
				var _nw354 *C3 = New_C3_()
				_ = _nw354
				_nw354.Ctor__(((_dafny.Zero).Minus(p0)).Minus((_2203_v36).Cardinality()), Companion_Default___.SafeModuloInt((_this).F11(), (_2197_v30).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13()))).Cardinality(), _dafny.IntOfUint32((_2197_v30).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.IntOfUint32((_2197_v30).Cardinality()), (_this).F12())
				_2204_v37 = _nw354
				var _2205_v38 _dafny.Sequence
				_ = _2205_v38
				_2205_v38 = _dafny.SeqOf(_2197_v30, _2197_v30, _2197_v30, _2197_v30)
				var _2206_v39 _dafny.MultiSet
				_ = _2206_v39
				_2206_v39 = _dafny.MultiSetOf(_dafny.IntOfUint32((_2159_v2).Cardinality()))
				var _2207_v40 _dafny.Map
				_ = _2207_v40
				_2207_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2206_v39, _2162_v5)
				var _2208_v41 *C7
				_ = _2208_v41
				var _nw355 *C7 = New_C7_()
				_ = _nw355
				_nw355.Ctor__(_2205_v38, (_this).F12(), (_dafny.Zero).Minus((func() _dafny.Int {
					if !((_this).Fm4(p0, _2207_v40, (_2204_v37).F22(), (_this).F13(), globalState)) {
						return _dafny.IntOfUint32((_2160_v3).Cardinality())
					}
					return (_dafny.Zero).Minus((_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int))
				})()))
				_2208_v41 = _nw355
				var _rhs386 *C7 = _2208_v41
				_ = _rhs386
				var _rhs387 _dafny.Sequence = _2160_v3
				_ = _rhs387
				_2208_v41 = _rhs386
				_2160_v3 = _rhs387
			} else {
				var _2209_v42 _dafny.Set
				_ = _2209_v42
				_2209_v42 = _dafny.SetOf(_2159_v2, _2159_v2)
				_2209_v42 = (func() _dafny.Set {
					if (_this).F13() {
						return _2209_v42
					}
					return _2209_v42
				})()
				var _2210_v43 *C5
				_ = _2210_v43
				var _nw356 *C5 = New_C5_()
				_ = _nw356
				_nw356.Ctor__()
				_2210_v43 = _nw356
				var _2211_v44 _dafny.Array
				_ = _2211_v44
				var _nw357 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
				_ = _nw357
				_2211_v44 = _nw357
				var _2212_v45 _dafny.Map
				_ = _2212_v45
				_2212_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2158_v1, (_this).F13())
				var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_2211_v44), 0))
				_ = _index308
				(_2211_v44).ArraySet1((func() bool {
					if (_2212_v45).Contains(_2158_v1) {
						return (_2212_v45).Get(_2158_v1).(bool)
					}
					return (_this).F13()
				})(), (_index308).Int())
				var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_2211_v44), 0))
				_ = _index309
				(_2211_v44).ArraySet1((_this).F13(), (_index309).Int())
				var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_2211_v44), 0))
				_ = _index310
				(_2211_v44).ArraySet1(((_this).Fm6((_this).F13(), (_this).F11(), globalState)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tgekcnj"), _2159_v2)).Cardinality())) > 0, (_index310).Int())
				(globalState).F7 = !(((_2211_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_2211_v44), 0))).Int()).(bool)) || (false))
			}
			var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))
			_ = _index311
			(_2156_v0).ArraySet1(((_this).F11()).Times((_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int)), (_index311).Int())
			var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))
			_ = _index312
			(_2156_v0).ArraySet1((_this).F11(), (_index312).Int())
			var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))
			_ = _index313
			(_2156_v0).ArraySet1(p0, (_index313).Int())
		}
		var _2213_v46 _dafny.MultiSet
		_ = _2213_v46
		_2213_v46 = _dafny.MultiSetOf(_dafny.IntOfInt64(629))
		var _2214_v48 _dafny.Map
		_ = _2214_v48
		_2214_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2213_v46, func() _dafny.Map {
			var _coll85 = _dafny.NewMapBuilder()
			_ = _coll85
			for _iter88 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(737), _dafny.IntOfInt64(-125))); ; {
				_compr_85, _ok88 := _iter88()
				if !_ok88 {
					break
				}
				var _2215_v47 _dafny.Int
				_2215_v47 = interface{}(_compr_85).(_dafny.Int)
				if ((_dafny.IntOfInt64(737)).Cmp(_2215_v47) <= 0) && ((_2215_v47).Cmp(_dafny.IntOfInt64(-125)) < 0) {
					_coll85.Add(Companion_Default___.SafeModuloInt(_2215_v47, (_dafny.MultiSetOf(_2161_v4)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int), true)).Cardinality())
				}
			}
			return _coll85.ToMap()
		}())
		var _2216_v49 _dafny.Map
		_ = _2216_v49
		_2216_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(globalState), !((_this).F13()))
		var _2217_v50 _dafny.Array
		_ = _2217_v50
		var _nwElement0_96 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).Fm4((_this).F11(), _2214_v48, _dafny.IntOfInt64(128), true, globalState))
		_ = _nwElement0_96
		var _nw358 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_96, nil, _dafny.IntOfInt64(12))
		_ = _nw358
		(_nw358).ArraySet1(_nwElement0_96, 0)
		(_nw358).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13()), 1)
		(_nw358).ArraySet1(_2216_v49, 2)
		(_nw358).ArraySet1(_2216_v49, 3)
		(_nw358).ArraySet1(_2216_v49, 4)
		(_nw358).ArraySet1((_2216_v49).Merge(_2216_v49), 5)
		(_nw358).ArraySet1((_2216_v49).Merge(Companion_Default___.Fm11(globalState)), 6)
		(_nw358).ArraySet1(Companion_Default___.Fm11(globalState), 7)
		(_nw358).ArraySet1(_2216_v49, 8)
		(_nw358).ArraySet1(_2216_v49, 9)
		(_nw358).ArraySet1(_2216_v49, 10)
		(_nw358).ArraySet1(Companion_Default___.Fm11(globalState), 11)
		_2217_v50 = _nw358
		for _iter89 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2217_v50), 0))); ; {
			_guard_loop_3, _ok89 := _iter89()
			if !_ok89 {
				break
			}
			var _2218_i8 _dafny.Int
			_2218_i8 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_2218_i8).Sign() != -1) && ((_2218_i8).Cmp(_dafny.ArrayLen((_2217_v50), 0)) < 0)) {
				(_2217_v50).ArraySet1(_2216_v49, (_2218_i8).Int())
			}
		}
		var _2219_v51 _dafny.Array
		_ = _2219_v51
		var _len0_42 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_42
		var _nw359 _dafny.Array
		_ = _nw359
		if _len0_42.Cmp(_dafny.Zero) == 0 {
			_nw359 = _dafny.NewArray(_len0_42)
		} else {
			var _init42 func(_dafny.Int) _dafny.MultiSet = (func(_2220_v4 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
				return func(_2221_i9 _dafny.Int) _dafny.MultiSet {
					return _2220_v4
				}
			})(_2161_v4)
			_ = _init42
			var _element0_42 = _init42(_dafny.Zero)
			_ = _element0_42
			_nw359 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
			(_nw359).ArraySet1(_element0_42, 0)
			var _nativeLen0_42 = (_len0_42).Int()
			_ = _nativeLen0_42
			for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
				(_nw359).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
			}
		}
		_2219_v51 = _nw359
		var _2222_v52 _dafny.Sequence
		_ = _2222_v52
		_2222_v52 = _dafny.SeqOf(_2219_v51, _2219_v51, _2219_v51)
		r0 = (_2222_v52).Select((Companion_Default___.SafeIndex((_2156_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2156_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2222_v52).Cardinality()))).Uint32()).(_dafny.Array)
		var _2223_v53 _dafny.Map
		_ = _2223_v53
		_2223_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2159_v2, _2213_v46)
		r1 = (func() _dafny.Map {
			if (_2160_v3).Select((Companion_Default___.SafeIndex((_this).Fm8(_2158_v1, globalState), _dafny.IntOfUint32((_2160_v3).Cardinality()))).Uint32()).(bool) {
				return _2223_v53
			}
			return (_2223_v53).Merge(_2223_v53)
		})()
		return r0, r1
	}
}
func (_this *C8) M2(p0 _dafny.Int, globalState *GlobalState) (_dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var _2224_v0 _dafny.CodePoint
		_ = _2224_v0
		_2224_v0 = _dafny.CodePoint('f')
		var _2225_v1 _dafny.Sequence
		_ = _2225_v1
		_2225_v1 = _dafny.SeqOf(_2224_v0, _2224_v0)
		r1 = _dafny.Companion_Sequence_.Contains(_2225_v1, _2224_v0)
		r1 = (_this).F13()
		r0 = (_this).F11()
		if true {
			(globalState).F7 = ((_this).F13()) == ((_this).F13())
			var _2226_v2 _dafny.Sequence
			_ = _2226_v2
			_2226_v2 = _dafny.SeqOf((_this).F13(), (_this).F13(), false, (_this).F13(), (_this).F13())
			var _2227_v3 D6
			_ = _2227_v3
			_2227_v3 = Companion_D6_.Create_DC10_(_2226_v2)
			var _2228_v4 *C0
			_ = _2228_v4
			var _nw360 *C0 = New_C0_()
			_ = _nw360
			_nw360.Ctor__((_2227_v3).Dtor_cf13(), (_this).F11())
			_2228_v4 = _nw360
			var _2229_v5 D23
			_ = _2229_v5
			_2229_v5 = Companion_D23_.Create_DC54_(_2228_v4)
			var _2230_v6 _dafny.Array
			_ = _2230_v6
			var _nwElement0_97 *C0 = _2228_v4
			_ = _nwElement0_97
			var _nw361 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_97, nil, _dafny.IntOfInt64(13))
			_ = _nw361
			(_nw361).ArraySet1(_nwElement0_97, 0)
			(_nw361).ArraySet1(_2228_v4, 1)
			(_nw361).ArraySet1(_2228_v4, 2)
			(_nw361).ArraySet1(_2228_v4, 3)
			(_nw361).ArraySet1(_2228_v4, 4)
			(_nw361).ArraySet1(_2228_v4, 5)
			(_nw361).ArraySet1(_2228_v4, 6)
			(_nw361).ArraySet1(_2228_v4, 7)
			(_nw361).ArraySet1(_2228_v4, 8)
			(_nw361).ArraySet1(_2228_v4, 9)
			(_nw361).ArraySet1((_2229_v5).Dtor_cf91(), 10)
			(_nw361).ArraySet1(_2228_v4, 11)
			(_nw361).ArraySet1(_2228_v4, 12)
			_2230_v6 = _nw361
			var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_2230_v6), 0))
			_ = _index314
			(_2230_v6).ArraySet1(_2228_v4, (_index314).Int())
			var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_2230_v6), 0))
			_ = _index315
			(_2230_v6).ArraySet1(_2228_v4, (_index315).Int())
			var _2231_v7 D15
			_ = _2231_v7
			_2231_v7 = Companion_D15_.Create_DC34_()
			var _source32 D15 = _2231_v7
			_ = _source32
			if _source32.Is_DC34() {
				var _2232_v8 _dafny.Map
				_ = _2232_v8
				_2232_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2224_v0, (_this).F13())
				_2232_v8 = (func() _dafny.Map {
					if (_this).F13() {
						return (_2232_v8).Merge(_2232_v8)
					}
					return (_2232_v8).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2224_v0, (_this).F13())).Update(_2224_v0, (_this).F13()))
				})()
				var _2233_v9 _dafny.Set
				_ = _2233_v9
				_2233_v9 = _dafny.SetOf(_2224_v0, _2224_v0)
				var _2234_v10 _dafny.Sequence
				_ = _2234_v10
				_2234_v10 = _dafny.SeqOf(_2233_v9)
				var _2235_v11 _dafny.MultiSet
				_ = _2235_v11
				_2235_v11 = _dafny.MultiSetOf((_2234_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.IntOfUint32((_2234_v10).Cardinality()))).Uint32()).(_dafny.Set))
				r1 = (_2235_v11).IsSubsetOf(_dafny.MultiSetFromSeq(_2234_v10))
				_2224_v0 = _2224_v0
				var _rhs388 bool = (_this).F13()
				_ = _rhs388
				var _rhs389 _dafny.Int = (_dafny.Zero).Minus((_2228_v4).F21())
				_ = _rhs389
				var _rhs390 _dafny.CodePoint = _2224_v0
				_ = _rhs390
				var _rhs391 _dafny.Int = Companion_Default___.Fm2((_2228_v4).F21(), (_dafny.Zero).Minus((_2228_v4).F21()), globalState)
				_ = _rhs391
				var _lhs330 *GlobalState = globalState
				_ = _lhs330
				var _lhs331 *GlobalState = globalState
				_ = _lhs331
				var _lhs332 *GlobalState = globalState
				_ = _lhs332
				_lhs330.F6 = _rhs388
				_lhs331.F1 = _rhs389
				_2224_v0 = _rhs390
				_lhs332.F1 = _rhs391
			} else if _source32.Is_DC35() {
				var _2236___mcc_h0 bool = _source32.Get_().(D15_DC35).Cf50
				_ = _2236___mcc_h0
				var _2237___mcc_h1 T1 = _source32.Get_().(D15_DC35).Cf51
				_ = _2237___mcc_h1
				var _2238___mcc_h2 bool = _source32.Get_().(D15_DC35).Cf52
				_ = _2238___mcc_h2
				var _2239___mcc_h3 _dafny.Sequence = _source32.Get_().(D15_DC35).Cf53
				_ = _2239___mcc_h3
				var _2240___mcc_h4 _dafny.Int = _source32.Get_().(D15_DC35).Cf54
				_ = _2240___mcc_h4
				var _2241_cf54 _dafny.Int = _2240___mcc_h4
				_ = _2241_cf54
				var _2242_cf53 _dafny.Sequence = _2239___mcc_h3
				_ = _2242_cf53
				var _2243_cf52 bool = _2238___mcc_h2
				_ = _2243_cf52
				var _2244_cf51 T1 = _2237___mcc_h1
				_ = _2244_cf51
				var _2245_cf50 bool = _2236___mcc_h0
				_ = _2245_cf50
				var _2246_v12 _dafny.Array
				_ = _2246_v12
				var _nw362 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw362
				_2246_v12 = _nw362
				var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_2246_v12), 0))
				_ = _index316
				(_2246_v12).ArraySet1((_this).F13(), (_index316).Int())
				var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_2246_v12), 0))
				_ = _index317
				(_2246_v12).ArraySet1((_this).F13(), (_index317).Int())
				var _rhs392 bool = (_this).F13()
				_ = _rhs392
				var _rhs393 _dafny.Int = _dafny.IntOfInt64(845)
				_ = _rhs393
				var _rhs394 bool = _2243_cf52
				_ = _rhs394
				var _lhs333 *GlobalState = globalState
				_ = _lhs333
				var _lhs334 *GlobalState = globalState
				_ = _lhs334
				var _lhs335 *GlobalState = globalState
				_ = _lhs335
				_lhs333.F7 = _rhs392
				_lhs334.F1 = _rhs393
				_lhs335.F7 = _rhs394
				var _2247_v13 _dafny.Array
				_ = _2247_v13
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_43
				var _nw363 _dafny.Array
				_ = _nw363
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw363 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) _dafny.Int = (func(_2248_v4 *C0) func(_dafny.Int) _dafny.Int {
						return func(_2249_i0 _dafny.Int) _dafny.Int {
							return (_2249_i0).Minus((_2248_v4).F21())
						}
					})(_2228_v4)
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw363 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw363).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw363).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_2247_v13 = _nw363
				var _2250_v14 _dafny.Set
				_ = _2250_v14
				_2250_v14 = _dafny.SetOf(_2247_v13, _2247_v13, _2247_v13, _2247_v13)
				var _2251_v15 D2
				_ = _2251_v15
				_2251_v15 = Companion_D2_.Create_DC4_(p0, _2241_cf54, _dafny.IntOfUint32((_2242_cf53).Cardinality()))
				var _2252_v16 *C6
				_ = _2252_v16
				var _nw364 *C6 = New_C6_()
				_ = _nw364
				_nw364.Ctor__(_2251_v15, _dafny.SetOf(_2245_cf50), (_2228_v4).F21())
				_2252_v16 = _nw364
				var _2253_v17 _dafny.Map
				_ = _2253_v17
				_2253_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2252_v16, _2245_cf50)
				var _2254_v18 _dafny.Sequence
				_ = _2254_v18
				_2254_v18 = _dafny.SeqOf((_2228_v4).F21(), ((_2253_v17).Update(_2252_v16, (_this).F13())).Cardinality())
				var _2255_v19 _dafny.Sequence
				_ = _2255_v19
				_2255_v19 = _dafny.SeqOf(_2254_v18, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(p0, _2241_cf54, (_this).F11()), (Companion_Default___.SafeIndex((_2228_v4).F21(), _dafny.IntOfUint32((_dafny.SeqOf(p0, _2241_cf54, (_this).F11())).Cardinality()))).Uint32(), (_this).F11()), _dafny.SeqOf((_2244_cf51).F11()), _2254_v18, _2254_v18)
				var _2256_v20 *C7
				_ = _2256_v20
				var _nw365 *C7 = New_C7_()
				_ = _nw365
				_nw365.Ctor__(_2255_v19, (_this).F12(), p0)
				_2256_v20 = _nw365
				var _2257_v21 _dafny.Array
				_ = _2257_v21
				var _len0_44 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_44
				var _nw366 _dafny.Array
				_ = _nw366
				if _len0_44.Cmp(_dafny.Zero) == 0 {
					_nw366 = _dafny.NewArray(_len0_44)
				} else {
					var _init44 func(_dafny.Int) _dafny.CodePoint = (func(_2258_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2259_i1 _dafny.Int) _dafny.CodePoint {
							return _2258_v0
						}
					})(_2224_v0)
					_ = _init44
					var _element0_44 = _init44(_dafny.Zero)
					_ = _element0_44
					_nw366 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
					(_nw366).ArraySet1CodePoint(_element0_44, 0)
					var _nativeLen0_44 = (_len0_44).Int()
					_ = _nativeLen0_44
					for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
						(_nw366).ArraySet1CodePoint(_init44(_dafny.IntOf(_i0_44)), _i0_44)
					}
				}
				_2257_v21 = _nw366
				var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_2257_v21), 0))
				_ = _index318
				(_2257_v21).ArraySet1CodePoint(_2224_v0, (_index318).Int())
				var _2260_v22 _dafny.Sequence
				_ = _2260_v22
				_2260_v22 = _dafny.SeqOf(_2256_v20)
				var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_2257_v21), 0))
				_ = _index319
				var _rhs395 bool = _2243_cf52
				_ = _rhs395
				var _rhs396 _dafny.Set = (_2250_v14).Intersection((_dafny.SetOf(_2247_v13)).Union(_2250_v14))
				_ = _rhs396
				var _rhs397 *C7 = (_2260_v22).Select((Companion_Default___.SafeIndex((_2228_v4).F21(), _dafny.IntOfUint32((_2260_v22).Cardinality()))).Uint32()).(*C7)
				_ = _rhs397
				var _rhs398 _dafny.Int = (_this).F11()
				_ = _rhs398
				var _rhs399 _dafny.CodePoint = _2224_v0
				_ = _rhs399
				var _lhs336 *GlobalState = globalState
				_ = _lhs336
				var _lhs337 *GlobalState = globalState
				_ = _lhs337
				var _lhs338 _dafny.Array = _2257_v21
				_ = _lhs338
				var _lhs339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_2257_v21), 0))
				_ = _lhs339
				_lhs336.F6 = _rhs395
				_2250_v14 = _rhs396
				_2256_v20 = _rhs397
				_lhs337.F1 = _rhs398
				(_lhs338).ArraySet1CodePoint(_rhs399, (_lhs339).Int())
				var _2261_v23 _dafny.Set
				_ = _2261_v23
				_2261_v23 = _dafny.SetOf(_2246_v12)
				var _2262_v24 _dafny.Array
				_ = _2262_v24
				var _nwElement0_98 _dafny.Set = _2261_v23
				_ = _nwElement0_98
				var _nw367 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_98, nil, _dafny.IntOfInt64(10))
				_ = _nw367
				(_nw367).ArraySet1(_nwElement0_98, 0)
				(_nw367).ArraySet1((_2261_v23).Intersection(_2261_v23), 1)
				(_nw367).ArraySet1((_2261_v23).Union(_2261_v23), 2)
				(_nw367).ArraySet1(_2261_v23, 3)
				(_nw367).ArraySet1(_2261_v23, 4)
				(_nw367).ArraySet1((Companion_D21_.Create_DC47_(_2261_v23)).Dtor_cf77(), 5)
				(_nw367).ArraySet1((_2261_v23).Difference(_2261_v23), 6)
				(_nw367).ArraySet1((_2261_v23).Difference(_2261_v23), 7)
				(_nw367).ArraySet1(_2261_v23, 8)
				(_nw367).ArraySet1(_2261_v23, 9)
				_2262_v24 = _nw367
				var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_2262_v24), 0))
				_ = _index320
				(_2262_v24).ArraySet1(_2261_v23, (_index320).Int())
				var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_2262_v24), 0))
				_ = _index321
				(_2262_v24).ArraySet1(((_2261_v23).Union(_2261_v23)).Union(_2261_v23), (_index321).Int())
			} else if _source32.Is_DC36() {
				var _2263___mcc_h5 _dafny.Int = _source32.Get_().(D15_DC36).Cf55
				_ = _2263___mcc_h5
				var _2264___mcc_h6 bool = _source32.Get_().(D15_DC36).Cf56
				_ = _2264___mcc_h6
				var _2265___mcc_h7 bool = _source32.Get_().(D15_DC36).Cf57
				_ = _2265___mcc_h7
				var _2266___mcc_h8 bool = _source32.Get_().(D15_DC36).Cf58
				_ = _2266___mcc_h8
				var _2267___mcc_h9 _dafny.Int = _source32.Get_().(D15_DC36).Cf59
				_ = _2267___mcc_h9
				var _2268_cf59 _dafny.Int = _2267___mcc_h9
				_ = _2268_cf59
				var _2269_cf58 bool = _2266___mcc_h8
				_ = _2269_cf58
				var _2270_cf57 bool = _2265___mcc_h7
				_ = _2270_cf57
				var _2271_cf56 bool = _2264___mcc_h6
				_ = _2271_cf56
				var _2272_cf55 _dafny.Int = _2263___mcc_h5
				_ = _2272_cf55
				var _2273_v25 _dafny.MultiSet
				_ = _2273_v25
				_2273_v25 = _dafny.MultiSetOf(Companion_Default___.Fm2((_2228_v4).F21(), (_2228_v4).F21(), globalState))
				var _2274_v26 _dafny.Map
				_ = _2274_v26
				_2274_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2270_cf57, _2268_cf59)
				var _2275_v27 _dafny.Set
				_ = _2275_v27
				_2275_v27 = _dafny.SetOf((_2273_v25).Contains(_2268_cf59), ((func() _dafny.Int {
					if (_2274_v26).Contains(_2269_cf58) {
						return (_2274_v26).Get(_2269_cf58).(_dafny.Int)
					}
					return _2268_cf59
				})()).Cmp(_dafny.IntOfInt64(939)) >= 0, (_dafny.IntOfInt64(819)).Cmp(_dafny.IntOfInt64(369)) != 0, _2270_cf57, _2271_cf56)
				var _2276_v28 D7
				_ = _2276_v28
				_2276_v28 = Companion_D7_.Create_DC13_(Companion_Default___.Fm41(globalState), (_2228_v4).F21(), (_2228_v4).F21())
				var _pat_let_tv54 = _2275_v27
				_ = _pat_let_tv54
				_2275_v27 = (func(_pat_let56_0 D7) D7 {
					return func(_2277_dt__update__tmp_h0 D7) D7 {
						return func(_pat_let57_0 _dafny.Set) D7 {
							return func(_2278_dt__update_hcf18_h0 _dafny.Set) D7 {
								return Companion_D7_.Create_DC13_(_2278_dt__update_hcf18_h0, (_2277_dt__update__tmp_h0).Dtor_cf19(), (_2277_dt__update__tmp_h0).Dtor_cf20())
							}(_pat_let57_0)
						}(_pat_let_tv54)
					}(_pat_let56_0)
				}(_2276_v28)).Dtor_cf18()
				_2224_v0 = _2224_v0
				var _2279_v29 _dafny.Array
				_ = _2279_v29
				var _len0_45 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_45
				var _nw368 _dafny.Array
				_ = _nw368
				if _len0_45.Cmp(_dafny.Zero) == 0 {
					_nw368 = _dafny.NewArray(_len0_45)
				} else {
					var _init45 func(_dafny.Int) bool = func(_2280_i2 _dafny.Int) bool {
						return (_this).F13()
					}
					_ = _init45
					var _element0_45 = _init45(_dafny.Zero)
					_ = _element0_45
					_nw368 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
					(_nw368).ArraySet1(_element0_45, 0)
					var _nativeLen0_45 = (_len0_45).Int()
					_ = _nativeLen0_45
					for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
						(_nw368).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
					}
				}
				_2279_v29 = _nw368
				var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_2279_v29), 0))
				_ = _index322
				(_2279_v29).ArraySet1(!(_2270_cf57), (_index322).Int())
				var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_2279_v29), 0))
				_ = _index323
				(_2279_v29).ArraySet1(_2270_cf57, (_index323).Int())
				_2225_v1 = _dafny.Companion_Sequence_.Concatenate(_2225_v1, _2225_v1)
			} else if _source32.Is_DC33() {
				var _2281___mcc_h10 _dafny.Map = _source32.Get_().(D15_DC33).Cf49
				_ = _2281___mcc_h10
				var _2282_cf49 _dafny.Map = _2281___mcc_h10
				_ = _2282_cf49
				var _2283_v30 _dafny.Map
				_ = _2283_v30
				_2283_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_2228_v4).F21()), (_2228_v4).F21())
				_2283_v30 = (_2283_v30).Update(_dafny.IntOfInt64(574), (_this).F11())
				(globalState).F7 = (_this).F13()
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_2230_v6), 0))
				_ = _index324
				(_2230_v6).ArraySet1(_2228_v4, (_index324).Int())
				var _2284_v31 _dafny.Array
				_ = _2284_v31
				var _nw369 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw369
				_2284_v31 = _nw369
				var _2285_v32 D2
				_ = _2285_v32
				_2285_v32 = Companion_D2_.Create_DC4_((_2228_v4).F21(), _dafny.IntOfInt64(-724), (_2228_v4).F21())
				var _2286_v33 *C6
				_ = _2286_v33
				var _nw370 *C6 = New_C6_()
				_ = _nw370
				_nw370.Ctor__(_2285_v32, (_this).F12(), (_2228_v4).F21())
				_2286_v33 = _nw370
				var _2287_v34 _dafny.MultiSet
				_ = _2287_v34
				_2287_v34 = _dafny.MultiSetOf(_2286_v33, _2286_v33, _2286_v33)
				var _2288_v35 _dafny.Set
				_ = _2288_v35
				_2288_v35 = _dafny.SetOf(_2287_v34)
				var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_2284_v31), 0))
				_ = _index325
				(_2284_v31).ArraySet1((_2288_v35).IsProperSubsetOf(_2288_v35), (_index325).Int())
				var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_2284_v31), 0))
				_ = _index326
				(_2284_v31).ArraySet1(false, (_index326).Int())
			} else {
				var _2289___mcc_h11 D15 = _source32.Get_().(D15_DC37).Cf60
				_ = _2289___mcc_h11
				var _2290_cf60 D15 = _2289___mcc_h11
				_ = _2290_cf60
				r1 = (_this).F13()
				(globalState).F7 = true
				(globalState).F1 = (_2228_v4).F21()
				var _2291_v36 D11
				_ = _2291_v36
				_2291_v36 = Companion_D11_.Create_DC24_(_dafny.IntOfInt64(994), (_2228_v4).F20(), _2224_v0)
				var _2292_v37 _dafny.MultiSet
				_ = _2292_v37
				_2292_v37 = _dafny.MultiSetOf((_this).F13(), (_this).F13(), false)
				var _2293_v38 _dafny.Sequence
				_ = _2293_v38
				_2293_v38 = _dafny.SeqOf((_this).Fm7(p0, (_this).F13(), _2292_v37, globalState))
				var _2294_v39 D15
				_ = _2294_v39
				_2294_v39 = Companion_D15_.Create_DC36_((_this).F11(), (_this).F13(), (_this).F13(), (_this).F13(), _dafny.IntOfInt64(369))
				var _2295_v40 _dafny.Map
				_ = _2295_v40
				_2295_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2224_v0, (_this).F13())
				var _2296_v41 _dafny.MultiSet
				_ = _2296_v41
				_2296_v41 = _dafny.MultiSetOf((_this).F11())
				var _2297_v42 _dafny.Map
				_ = _2297_v42
				_2297_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2296_v41).Cardinality(), (_this).Fm6(false, (_2228_v4).F21(), globalState))
				var _2298_v43 D7
				_ = _2298_v43
				_2298_v43 = Companion_D7_.Create_DC13_(Companion_Default___.Fm41(globalState), (_this).F11(), (_2228_v4).F21())
				var _2299_v44 _dafny.Array
				_ = _2299_v44
				var _nwElement0_99 _dafny.Int = (_2291_v36).Dtor_cf34()
				_ = _nwElement0_99
				var _nw371 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_99, nil, _dafny.IntOfInt64(29))
				_ = _nw371
				(_nw371).ArraySet1(_nwElement0_99, 0)
				(_nw371).ArraySet1((_dafny.SetOf((_this).F13())).Cardinality(), 1)
				(_nw371).ArraySet1((_this).Fm8(_2224_v0, globalState), 2)
				(_nw371).ArraySet1((_2228_v4).F21(), 3)
				(_nw371).ArraySet1((_this).F11(), 4)
				(_nw371).ArraySet1(_dafny.IntOfInt64(822), 5)
				(_nw371).ArraySet1(_dafny.IntOfUint32((_2293_v38).Cardinality()), 6)
				(_nw371).ArraySet1((_this).F11(), 7)
				(_nw371).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2225_v1, _2225_v1)).Cardinality()), 8)
				(_nw371).ArraySet1(((_dafny.Zero).Minus((_2228_v4).F21())).Minus(p0), 9)
				(_nw371).ArraySet1(((_2294_v39).Dtor_cf55()).Times((_this).F11()), 10)
				(_nw371).ArraySet1((_dafny.Zero).Minus(((_2295_v40).Cardinality()).Plus(p0)), 11)
				(_nw371).ArraySet1(((_2297_v42).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2228_v4).F21(), (_this).F11()))).Cardinality(), 12)
				(_nw371).ArraySet1(p0, 13)
				(_nw371).ArraySet1(p0, 14)
				(_nw371).ArraySet1((_2292_v37).Cardinality(), 15)
				(_nw371).ArraySet1((func() _dafny.Int {
					if (_this).F13() {
						return (_this).F11()
					}
					return (_this).F11()
				})(), 16)
				(_nw371).ArraySet1(p0, 17)
				(_nw371).ArraySet1(p0, 18)
				(_nw371).ArraySet1((func() _dafny.Int {
					if (_2228_v4).Fm21((_this).F13(), _dafny.IntOfUint32((_2293_v38).Cardinality()), globalState) {
						return (_this).F11()
					}
					return _dafny.IntOfInt64(8)
				})(), 19)
				(_nw371).ArraySet1(_dafny.IntOfInt64(-315), 20)
				(_nw371).ArraySet1((_2292_v37).Cardinality(), 21)
				(_nw371).ArraySet1((_this).F11(), 22)
				(_nw371).ArraySet1((_2298_v43).Dtor_cf20(), 23)
				(_nw371).ArraySet1((_this).F11(), 24)
				(_nw371).ArraySet1((_2228_v4).F21(), 25)
				(_nw371).ArraySet1((_2228_v4).F21(), 26)
				(_nw371).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_2228_v4).F21()), (_2228_v4).F21()), 27)
				(_nw371).ArraySet1((_this).F11(), 28)
				_2299_v44 = _nw371
				var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_2299_v44), 0))
				_ = _index327
				(_2299_v44).ArraySet1(p0, (_index327).Int())
				var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_2299_v44), 0))
				_ = _index328
				(_2299_v44).ArraySet1((_this).Fm6((_this).F13(), (_2228_v4).F21(), globalState), (_index328).Int())
			}
			var _2300_v45 _dafny.Array
			_ = _2300_v45
			var _len0_46 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_46
			var _nw372 _dafny.Array
			_ = _nw372
			if _len0_46.Cmp(_dafny.Zero) == 0 {
				_nw372 = _dafny.NewArray(_len0_46)
			} else {
				var _init46 func(_dafny.Int) _dafny.Sequence = (func(_2301_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_2302_i3 _dafny.Int) _dafny.Sequence {
						return _2301_v1
					}
				})(_2225_v1)
				_ = _init46
				var _element0_46 = _init46(_dafny.Zero)
				_ = _element0_46
				_nw372 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
				(_nw372).ArraySet1(_element0_46, 0)
				var _nativeLen0_46 = (_len0_46).Int()
				_ = _nativeLen0_46
				for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
					(_nw372).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
				}
			}
			_2300_v45 = _nw372
			var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.ArrayLen((_2300_v45), 0))
			_ = _index329
			(_2300_v45).ArraySet1(_2225_v1, (_index329).Int())
			var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.ArrayLen((_2300_v45), 0))
			_ = _index330
			(_2300_v45).ArraySet1(_dafny.Companion_Sequence_.Update(_2225_v1, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2225_v1).Cardinality()))).Uint32(), _2224_v0), (_index330).Int())
			var _2303_v46 _dafny.Sequence
			_ = _2303_v46
			_2303_v46 = _dafny.SeqOf((_2228_v4).F21())
			var _2304_v47 _dafny.Map
			_ = _2304_v47
			_2304_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2224_v0, _2303_v46)
			var _2305_v48 _dafny.Sequence
			_ = _2305_v48
			_2305_v48 = _dafny.SeqOf(_2303_v46, (func() _dafny.Sequence {
				if (_2304_v47).Contains(_dafny.CodePoint('f')) {
					return (_2304_v47).Get(_dafny.CodePoint('f')).(_dafny.Sequence)
				}
				return _2303_v46
			})())
			var _2306_v49 *C7
			_ = _2306_v49
			var _nw373 *C7 = New_C7_()
			_ = _nw373
			_nw373.Ctor__(_2305_v48, (_this).F12(), (_2228_v4).F21())
			_2306_v49 = _nw373
		} else {
			var _2307_v50 _dafny.Sequence
			_ = _2307_v50
			_2307_v50 = _dafny.SeqOf((_this).F13(), (_this).F13(), !((_this).F13()))
			var _2308_v51 _dafny.MultiSet
			_ = _2308_v51
			_2308_v51 = _dafny.MultiSetOf(_2307_v50, _2307_v50, _2307_v50, _2307_v50, _2307_v50)
			(globalState).F1 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_2225_v1).Cardinality()), (_this).F11()), (_2308_v51).Cardinality())
			var _2309_v52 D2
			_ = _2309_v52
			_2309_v52 = Companion_D2_.Create_DC4_(p0, p0, p0)
			var _2310_v53 *C6
			_ = _2310_v53
			var _nw374 *C6 = New_C6_()
			_ = _nw374
			_nw374.Ctor__(func(_pat_let58_0 D2) D2 {
				return func(_2311_dt__update__tmp_h1 D2) D2 {
					return func(_pat_let59_0 _dafny.Int) D2 {
						return func(_2312_dt__update_hcf8_h0 _dafny.Int) D2 {
							return Companion_D2_.Create_DC4_((_2311_dt__update__tmp_h1).Dtor_cf6(), (_2311_dt__update__tmp_h1).Dtor_cf7(), _2312_dt__update_hcf8_h0)
						}(_pat_let59_0)
					}((_this).F11())
				}(_pat_let58_0)
			}(_2309_v52), Companion_Default___.Fm41(globalState), p0)
			_2310_v53 = _nw374
			var _2313_v54 D13
			_ = _2313_v54
			_2313_v54 = Companion_D13_.Create_DC28_(_2225_v1, p0, (_this).F13())
			r0 = ((_2313_v54).Dtor_cf41()).Minus((_this).F11())
			var _2314_v55 _dafny.Array
			_ = _2314_v55
			var _nw375 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
			_ = _nw375
			_2314_v55 = _nw375
			var _2315_v56 _dafny.Array
			_ = _2315_v56
			var _nw376 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
			_ = _nw376
			_2315_v56 = _nw376
			var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_2314_v55), 0))
			_ = _index331
			(_2314_v55).ArraySet1(_2315_v56, (_index331).Int())
			var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_2314_v55), 0))
			_ = _index332
			var _nw377 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw377
			(_2314_v55).ArraySet1(_nw377, (_index332).Int())
			if (_this).F13() {
				var _2316_v57 _dafny.Array
				_ = _2316_v57
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_47
				var _nw378 _dafny.Array
				_ = _nw378
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw378 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) _dafny.CodePoint = (func(_2317_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2318_i4 _dafny.Int) _dafny.CodePoint {
							return _2317_v0
						}
					})(_2224_v0)
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw378 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw378).ArraySet1CodePoint(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw378).ArraySet1CodePoint(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_2316_v57 = _nw378
				var _2319_v58 _dafny.Sequence
				_ = _2319_v58
				_2319_v58 = _dafny.SeqOf(_2316_v57, _2316_v57, _2316_v57)
				var _2320_v59 _dafny.Map
				_ = _2320_v59
				_2320_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _2316_v57)
				var _2321_v60 _dafny.Array
				_ = _2321_v60
				var _nwElement0_100 _dafny.Array = _2316_v57
				_ = _nwElement0_100
				var _nw379 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_100, nil, _dafny.IntOfInt64(20))
				_ = _nw379
				(_nw379).ArraySet1(_nwElement0_100, 0)
				(_nw379).ArraySet1(_2316_v57, 1)
				(_nw379).ArraySet1((_2319_v58).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2319_v58).Cardinality()))).Uint32()).(_dafny.Array), 2)
				(_nw379).ArraySet1(_2316_v57, 3)
				(_nw379).ArraySet1(_2316_v57, 4)
				(_nw379).ArraySet1(_2316_v57, 5)
				(_nw379).ArraySet1(_2316_v57, 6)
				(_nw379).ArraySet1(_2316_v57, 7)
				(_nw379).ArraySet1(_2316_v57, 8)
				(_nw379).ArraySet1(_2316_v57, 9)
				(_nw379).ArraySet1((func() _dafny.Array {
					if (_2320_v59).Contains((_this).F13()) {
						return (_2320_v59).Get((_this).F13()).(_dafny.Array)
					}
					return (_2319_v58).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2319_v58).Cardinality()))).Uint32()).(_dafny.Array)
				})(), 10)
				(_nw379).ArraySet1(_2316_v57, 11)
				(_nw379).ArraySet1(_2316_v57, 12)
				(_nw379).ArraySet1(_2316_v57, 13)
				(_nw379).ArraySet1((func() _dafny.Array {
					if (_this).F13() {
						return _2316_v57
					}
					return _2316_v57
				})(), 14)
				(_nw379).ArraySet1(_2316_v57, 15)
				(_nw379).ArraySet1(_2316_v57, 16)
				(_nw379).ArraySet1(_2316_v57, 17)
				(_nw379).ArraySet1(_2316_v57, 18)
				(_nw379).ArraySet1(_2316_v57, 19)
				_2321_v60 = _nw379
				var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2321_v60), 0))
				_ = _index333
				(_2321_v60).ArraySet1(_2316_v57, (_index333).Int())
				var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2321_v60), 0))
				_ = _index334
				(_2321_v60).ArraySet1(_2316_v57, (_index334).Int())
				var _2322_v61 *C2
				_ = _2322_v61
				var _nw380 *C2 = New_C2_()
				_ = _nw380
				_nw380.Ctor__(_dafny.CodePoint('s'), p0)
				_2322_v61 = _nw380
				var _2323_v62 _dafny.Sequence
				_ = _2323_v62
				_2323_v62 = _dafny.SeqOf(p0)
				(globalState).F1 = ((p0).Minus((_2323_v62).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_2323_v62).Cardinality()))).Uint32()).(_dafny.Int))).Plus(Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfUint32((_2225_v1).Cardinality())))
				var _2324_v63 D7
				_ = _2324_v63
				_2324_v63 = Companion_D7_.Create_DC12_(_2323_v62)
				_2324_v63 = _2324_v63
				var _2325_v64 _dafny.Array
				_ = _2325_v64
				var _nw381 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw381
				_2325_v64 = _nw381
				var _2326_v65 _dafny.Map
				_ = _2326_v65
				_2326_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F11()).Times(_dafny.IntOfUint32((_2225_v1).Cardinality())), _2325_v64)
				var _2327_v66 _dafny.Set
				_ = _2327_v66
				_2327_v66 = _dafny.SetOf((_this).F11(), (_this).F11())
				var _2328_v67 _dafny.MultiSet
				_ = _2328_v67
				_2328_v67 = _dafny.MultiSetOf((_2327_v66).Cardinality())
				var _rhs400 _dafny.Int = (_2326_v65).Cardinality()
				_ = _rhs400
				var _rhs401 bool = !(!((_this).F13()))
				_ = _rhs401
				var _rhs402 bool = (_2328_v67).IsSubsetOf((_2328_v67).Intersection(_dafny.MultiSetFromSeq(_2323_v62)))
				_ = _rhs402
				var _rhs403 _dafny.Int = Companion_Default___.SafeModuloInt(p0, (_this).F11())
				_ = _rhs403
				var _rhs404 bool = (_this).F13()
				_ = _rhs404
				var _lhs340 *GlobalState = globalState
				_ = _lhs340
				var _lhs341 *GlobalState = globalState
				_ = _lhs341
				var _lhs342 *GlobalState = globalState
				_ = _lhs342
				var _lhs343 *GlobalState = globalState
				_ = _lhs343
				r0 = _rhs400
				_lhs340.F6 = _rhs401
				_lhs341.F7 = _rhs402
				_lhs342.F1 = _rhs403
				_lhs343.F6 = _rhs404
			} else {
				var _2329_v68 _dafny.Array
				_ = _2329_v68
				var _nw382 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
				_ = _nw382
				_2329_v68 = _nw382
				var _2330_v69 _dafny.Set
				_ = _2330_v69
				_2330_v69 = _dafny.SetOf(_2329_v68)
				var _2331_v70 T1
				_ = _2331_v70
				var _nw383 *C6 = New_C6_()
				_ = _nw383
				_nw383.Ctor__((_2310_v53).F15(), (_this).F12(), (_2330_v69).Cardinality())
				_2331_v70 = _nw383
				var _2332_v71 _dafny.Map
				_ = _2332_v71
				_2332_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2331_v70, p0)
				var _2333_v72 _dafny.Sequence
				_ = _2333_v72
				_2333_v72 = _dafny.SeqOf((_2332_v71).Cardinality(), _dafny.IntOfUint32((_2225_v1).Cardinality()))
				var _2334_v73 _dafny.Sequence
				_ = _2334_v73
				_2334_v73 = _dafny.SeqOf(_2333_v72)
				var _2335_v74 _dafny.Map
				_ = _2335_v74
				_2335_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2331_v70).F11(), (_dafny.Zero).Minus((_this).F11()))
				var _2336_v75 *C7
				_ = _2336_v75
				var _nw384 *C7 = New_C7_()
				_ = _nw384
				_nw384.Ctor__(_2334_v73, (_dafny.SetOf((_this).F13())).Intersection((_this).F12()), ((func() _dafny.Int {
					if (_2335_v74).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_2225_v1).Cardinality()))) {
						return (_2335_v74).Get((_dafny.Zero).Minus(_dafny.IntOfUint32((_2225_v1).Cardinality()))).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_2225_v1).Cardinality())
				})()).Minus((_2331_v70).F11()))
				_2336_v75 = _nw384
				var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_2329_v68), 0))
				_ = _index335
				(_2329_v68).ArraySet1((func() bool {
					if (_this).F13() {
						return (_this).F13()
					}
					return (_this).F13()
				})(), (_index335).Int())
				var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_2329_v68), 0))
				_ = _index336
				(_2329_v68).ArraySet1(((_2331_v70).F12()).IsProperSubsetOf(((_this).F12()).Union((_this).F12())), (_index336).Int())
				var _2337_v76 _dafny.MultiSet
				_ = _2337_v76
				_2337_v76 = _dafny.MultiSetOf((_this).F13(), (_2329_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_2329_v68), 0))).Int()).(bool))
				var _2338_v77 _dafny.MultiSet
				_ = _2338_v77
				_2338_v77 = _2337_v76
				var _2339_v78 _dafny.Map
				_ = _2339_v78
				_2339_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_2329_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_2329_v68), 0))).Int()).(bool)), _2338_v77)
				var _2340_v79 D0
				_ = _2340_v79
				_2340_v79 = Companion_D0_.Create_DC1_((_this).F13(), (_2331_v70).F11(), (_this).F13())
				var _2341_v80 _dafny.Sequence
				_ = _2341_v80
				_2341_v80 = _dafny.SeqOf(_2340_v79)
				var _2342_v81 D8
				_ = _2342_v81
				_2342_v81 = Companion_D8_.Create_DC16_((_this).F13(), _dafny.IntOfUint32((_2341_v80).Cardinality()))
				var _2343_v82 _dafny.Map
				_ = _2343_v82
				_2343_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2224_v0)
				var _2344_v83 _dafny.Map
				_ = _2344_v83
				_2344_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_2329_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_2329_v68), 0))).Int()).(bool))
				var _2345_v84 D7
				_ = _2345_v84
				_2345_v84 = Companion_D7_.Create_DC13_((_2331_v70).F12(), (_this).F11(), (_this).F11())
				var _2346_v85 _dafny.Map
				_ = _2346_v85
				_2346_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _2345_v84)
				var _2347_v86 D7
				_ = _2347_v86
				_2347_v86 = Companion_D7_.Create_DC14_((func() D7 {
					if (_2346_v85).Contains((_this).F11()) {
						return (_2346_v85).Get((_this).F11()).(D7)
					}
					return _2345_v84
				})())
				var _2348_v87 _dafny.Map
				_ = _2348_v87
				_2348_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F11())
				var _2349_v88 _dafny.Array
				_ = _2349_v88
				var _nwElement0_101 _dafny.Int = (_this).F11()
				_ = _nwElement0_101
				var _nw385 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_101, nil, _dafny.IntOfInt64(29))
				_ = _nw385
				(_nw385).ArraySet1(_nwElement0_101, 0)
				(_nw385).ArraySet1(p0, 1)
				(_nw385).ArraySet1(p0, 2)
				(_nw385).ArraySet1((_this).F11(), 3)
				(_nw385).ArraySet1((_2339_v78).Cardinality(), 4)
				(_nw385).ArraySet1((_2331_v70).F11(), 5)
				(_nw385).ArraySet1(Companion_Default___.SafeDivisionInt(p0, p0), 6)
				(_nw385).ArraySet1(((_dafny.Zero).Minus((_2331_v70).F11())).Minus((_2342_v81).Dtor_cf24()), 7)
				(_nw385).ArraySet1((_2343_v82).Cardinality(), 8)
				(_nw385).ArraySet1((_dafny.IntOfInt64(787)).Times(p0), 9)
				(_nw385).ArraySet1(p0, 10)
				(_nw385).ArraySet1((_this).F11(), 11)
				(_nw385).ArraySet1((_dafny.Zero).Minus((_this).F11()), 12)
				(_nw385).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if (_this).F13() {
						return (_dafny.Zero).Minus((_2344_v83).Cardinality())
					}
					return p0
				})()), 13)
				(_nw385).ArraySet1((func() _dafny.Int {
					if Companion_Default___.Fm0(globalState) {
						return (_2331_v70).F11()
					}
					return (_2336_v75).Fm6((_this).F13(), (_2331_v70).F11(), globalState)
				})(), 14)
				(_nw385).ArraySet1((_this).F11(), 15)
				(_nw385).ArraySet1((_dafny.SetOf(_2347_v86)).Cardinality(), 16)
				(_nw385).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F11(), (_2337_v76).Cardinality()), 17)
				(_nw385).ArraySet1(p0, 18)
				(_nw385).ArraySet1(Companion_Default___.SafeModuloInt(p0, p0), 19)
				(_nw385).ArraySet1(_dafny.IntOfInt64(-888), 20)
				(_nw385).ArraySet1((_this).F11(), 21)
				(_nw385).ArraySet1(p0, 22)
				(_nw385).ArraySet1((_2331_v70).F11(), 23)
				(_nw385).ArraySet1((_dafny.Zero).Minus(((_this).F11()).Plus((_2331_v70).F11())), 24)
				(_nw385).ArraySet1(((_this).F11()).Plus((_2331_v70).F11()), 25)
				(_nw385).ArraySet1((func() _dafny.Int {
					if (_2329_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_2329_v68), 0))).Int()).(bool) {
						return (_2344_v83).Cardinality()
					}
					return (_2348_v87).Cardinality()
				})(), 26)
				(_nw385).ArraySet1(p0, 27)
				(_nw385).ArraySet1(p0, 28)
				_2349_v88 = _nw385
				var _len0_48 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_48
				var _nw386 _dafny.Array
				_ = _nw386
				if _len0_48.Cmp(_dafny.Zero) == 0 {
					_nw386 = _dafny.NewArray(_len0_48)
				} else {
					var _init48 func(_dafny.Int) _dafny.Int = (func(_2350_v70 T1) func(_dafny.Int) _dafny.Int {
						return func(_2351_i5 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_2351_i5, (_2350_v70).F11())
						}
					})(_2331_v70)
					_ = _init48
					var _element0_48 = _init48(_dafny.Zero)
					_ = _element0_48
					_nw386 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
					(_nw386).ArraySet1(_element0_48, 0)
					var _nativeLen0_48 = (_len0_48).Int()
					_ = _nativeLen0_48
					for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
						(_nw386).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
					}
				}
				_2349_v88 = _nw386
				r1 = (_2329_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_2329_v68), 0))).Int()).(bool)
				(globalState).F1 = ((p0).Times(_dafny.IntOfInt64(714))).Times((func() _dafny.Set {
					var _coll86 = _dafny.NewBuilder()
					_ = _coll86
					for _iter90 := _dafny.Iterate((_2344_v83).Keys().Elements()); ; {
						_compr_86, _ok90 := _iter90()
						if !_ok90 {
							break
						}
						var _2352_v89 _dafny.Int
						_2352_v89 = interface{}(_compr_86).(_dafny.Int)
						if (_2344_v83).Contains(_2352_v89) {
							_coll86.Add((_2352_v89).Minus(_dafny.IntOfInt64(611)))
						}
					}
					return _coll86.ToSet()
				}()).Cardinality())
			}
		}
		var _2353_v90 _dafny.Sequence
		_ = _2353_v90
		_2353_v90 = _dafny.SeqOf((_this).F11())
		var _2354_v91 _dafny.MultiSet
		_ = _2354_v91
		_2354_v91 = _dafny.MultiSetOf(p0, p0, (_this).F11())
		var _2355_v92 _dafny.Sequence
		_ = _2355_v92
		var _out39 _dafny.Sequence
		_ = _out39
		_out39 = (_this).M3((_this).F13(), _dafny.IntOfInt64(-814), (_this).F11(), (_2354_v91).IsProperSubsetOf(_dafny.MultiSetFromSeq(_2353_v90)), globalState)
		_2355_v92 = _out39
		var _2356_v93 _dafny.Sequence
		_ = _2356_v93
		_2356_v93 = _dafny.SeqOf(Companion_Default___.Fm0(globalState))
		var _2357_v94 _dafny.Sequence
		_ = _2357_v94
		_2357_v94 = _dafny.SeqOf((_2356_v93).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2356_v93).Cardinality()))).Uint32()).(bool), (_this).F13())
		var _2358_v95 D0
		_ = _2358_v95
		_2358_v95 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2356_v93, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2356_v93).Cardinality()))).Uint32(), !((_this).F13()))).Cardinality()))
		var _2359_v96 _dafny.MultiSet
		_ = _2359_v96
		_2359_v96 = _dafny.MultiSetOf(_2224_v0, _dafny.CodePoint('u'))
		var _2360_v97 _dafny.Sequence
		_ = _2360_v97
		var _out40 _dafny.Sequence
		_ = _out40
		_out40 = (_this).M3((_this).Fm5(_2358_v95, _2357_v94, (_2359_v96).Cardinality(), _2224_v0, globalState), (_this).F11(), (_this).F11(), (_this).F13(), globalState)
		_2360_v97 = _out40
		r0 = Companion_Default___.SafeDivisionInt(p0, (_this).F11())
		r1 = (_this).F13()
		return r0, r1
	}
}
func (_this *C8) M3(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var _2361_v0 _dafny.Sequence
		_ = _2361_v0
		_2361_v0 = _dafny.SeqOf(true)
		_2361_v0 = _dafny.Companion_Sequence_.Concatenate(_2361_v0, _2361_v0)
		var _2362_v1 _dafny.CodePoint
		_ = _2362_v1
		_2362_v1 = _dafny.CodePoint('v')
		var _2363_v2 _dafny.Map
		_ = _2363_v2
		_2363_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2362_v1, p0)
		_2363_v2 = _2363_v2
		var _2364_v3 _dafny.Sequence
		_ = _2364_v3
		_2364_v3 = _dafny.UnicodeSeqOfUtf8Bytes("ifvvwmi")
		if (func() bool {
			if !((_this).F13()) {
				return false
			}
			return !((_dafny.IntOfUint32((_2364_v3).Cardinality())).Cmp((_this).F11()) <= 0)
		})() {
			var _2365_v4 _dafny.Array
			_ = _2365_v4
			var _len0_49 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_49
			var _nw387 _dafny.Array
			_ = _nw387
			if _len0_49.Cmp(_dafny.Zero) == 0 {
				_nw387 = _dafny.NewArray(_len0_49)
			} else {
				var _init49 func(_dafny.Int) _dafny.Int = (func(_2366_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_2367_i0 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_2367_i0, _dafny.IntOfUint32((_2366_v0).Cardinality()))
					}
				})(_2361_v0)
				_ = _init49
				var _element0_49 = _init49(_dafny.Zero)
				_ = _element0_49
				_nw387 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
				(_nw387).ArraySet1(_element0_49, 0)
				var _nativeLen0_49 = (_len0_49).Int()
				_ = _nativeLen0_49
				for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
					(_nw387).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
				}
			}
			_2365_v4 = _nw387
			var _2368_v5 _dafny.Map
			_ = _2368_v5
			_2368_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2365_v4, p0)
			var _2369_v6 _dafny.Sequence
			_ = _2369_v6
			_2369_v6 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2365_v4, p0)).Update(_2365_v4, (_this).F13()))
			var _2370_v7 D19
			_ = _2370_v7
			_2370_v7 = Companion_D19_.Create_DC42_(_2365_v4)
			var _2371_v8 _dafny.Array
			_ = _2371_v8
			var _nwElement0_102 _dafny.Map = _2368_v5
			_ = _nwElement0_102
			var _nw388 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_102, nil, _dafny.IntOfInt64(15))
			_ = _nw388
			(_nw388).ArraySet1(_nwElement0_102, 0)
			(_nw388).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2365_v4, Companion_Default___.Fm0(globalState)), 1)
			(_nw388).ArraySet1((func() _dafny.Map {
				if true {
					return _2368_v5
				}
				return _2368_v5
			})(), 2)
			(_nw388).ArraySet1((_2368_v5).Merge(_2368_v5), 3)
			(_nw388).ArraySet1((_2368_v5).Merge(_2368_v5), 4)
			(_nw388).ArraySet1((func() _dafny.Map {
				if (_this).F13() {
					return (_2369_v6).Select((Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_2369_v6).Cardinality()))).Uint32()).(_dafny.Map)
				}
				return _2368_v5
			})(), 5)
			(_nw388).ArraySet1(_2368_v5, 6)
			(_nw388).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2370_v7).Dtor_cf68(), p3)).Merge(_2368_v5), 7)
			(_nw388).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2365_v4, p0)).Merge(_2368_v5), 8)
			(_nw388).ArraySet1(_2368_v5, 9)
			(_nw388).ArraySet1(_2368_v5, 10)
			(_nw388).ArraySet1(_2368_v5, 11)
			(_nw388).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2365_v4, (_this).F13()), 12)
			(_nw388).ArraySet1(_2368_v5, 13)
			(_nw388).ArraySet1((_2368_v5).Merge(_2368_v5), 14)
			_2371_v8 = _nw388
			var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_2371_v8), 0))
			_ = _index337
			(_2371_v8).ArraySet1(_2368_v5, (_index337).Int())
			var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_2371_v8), 0))
			_ = _index338
			(_2371_v8).ArraySet1(_2368_v5, (_index338).Int())
			var _2372_v9 _dafny.Set
			_ = _2372_v9
			_2372_v9 = _dafny.SetOf((_this).F11())
			var _2373_v10 _dafny.Map
			_ = _2373_v10
			_2373_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2372_v9).Cardinality(), p0)
			var _2374_v11 _dafny.Sequence
			_ = _2374_v11
			_2374_v11 = _dafny.SeqOf(_2373_v10)
			var _2375_v12 _dafny.Map
			_ = _2375_v12
			_2375_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_2374_v11).Cardinality()))
			var _2376_v13 _dafny.Array
			_ = _2376_v13
			var _nwElement0_103 _dafny.CodePoint = _2362_v1
			_ = _nwElement0_103
			var _nw389 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_103, nil, _dafny.IntOfInt64(22))
			_ = _nw389
			(_nw389).ArraySet1CodePoint(_nwElement0_103, 0)
			(_nw389).ArraySet1CodePoint(_2362_v1, 1)
			(_nw389).ArraySet1CodePoint(_2362_v1, 2)
			(_nw389).ArraySet1CodePoint(_2362_v1, 3)
			(_nw389).ArraySet1CodePoint(_2362_v1, 4)
			(_nw389).ArraySet1CodePoint(_2362_v1, 5)
			(_nw389).ArraySet1CodePoint(_2362_v1, 6)
			(_nw389).ArraySet1CodePoint(_dafny.CodePoint('y'), 7)
			(_nw389).ArraySet1CodePoint(_2362_v1, 8)
			(_nw389).ArraySet1CodePoint(_2362_v1, 9)
			(_nw389).ArraySet1CodePoint(_2362_v1, 10)
			(_nw389).ArraySet1CodePoint(_dafny.CodePoint('k'), 11)
			(_nw389).ArraySet1CodePoint(Companion_Default___.Fm20(_2361_v0, _dafny.IntOfInt64(663), p0, globalState), 12)
			(_nw389).ArraySet1CodePoint(_2362_v1, 13)
			(_nw389).ArraySet1CodePoint(_2362_v1, 14)
			(_nw389).ArraySet1CodePoint(_2362_v1, 15)
			(_nw389).ArraySet1CodePoint(_2362_v1, 16)
			(_nw389).ArraySet1CodePoint(_2362_v1, 17)
			(_nw389).ArraySet1CodePoint(_2362_v1, 18)
			(_nw389).ArraySet1CodePoint(_2362_v1, 19)
			(_nw389).ArraySet1CodePoint(_2362_v1, 20)
			(_nw389).ArraySet1CodePoint(_2362_v1, 21)
			_2376_v13 = _nw389
			var _2377_v14 _dafny.Map
			_ = _2377_v14
			_2377_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2376_v13, _2362_v1)
			var _2378_v15 _dafny.MultiSet
			_ = _2378_v15
			_2378_v15 = _dafny.MultiSetOf((func() _dafny.Int {
				if (_2375_v12).Contains((_this).F13()) {
					return (_2375_v12).Get((_this).F13()).(_dafny.Int)
				}
				return (_2377_v14).Cardinality()
			})(), (_dafny.Zero).Minus((_this).F11()))
			var _2379_v16 _dafny.Map
			_ = _2379_v16
			_2379_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2378_v15, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(673)))
			var _2380_v17 _dafny.Map
			_ = _2380_v17
			_2380_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm1(globalState)).Cardinality()), p2)
			(globalState).F7 = !((_this).Fm4((_dafny.Zero).Minus(p1), (_2379_v16).Update(_2378_v15, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_dafny.MultiSetOf(p2, p2)).Cardinality())), (func() _dafny.Int {
				if (_2380_v17).Contains((_this).F11()) {
					return (_2380_v17).Get((_this).F11()).(_dafny.Int)
				}
				return _dafny.IntOfInt64(-482)
			})(), p3, globalState))
			_2365_v4 = (func() _dafny.Array {
				if ((_this).F11()).Cmp((_this).F11()) > 0 {
					return (func() _dafny.Array {
						if p3 {
							return _2365_v4
						}
						return _2365_v4
					})()
				}
				return _2365_v4
			})()
			_2376_v13 = _2376_v13
			var _2381_v18 _dafny.Array
			_ = _2381_v18
			var _nw390 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
			_ = _nw390
			_2381_v18 = _nw390
			var _rhs405 _dafny.Int = p1
			_ = _rhs405
			var _rhs406 _dafny.Array = _2381_v18
			_ = _rhs406
			var _rhs407 bool = p0
			_ = _rhs407
			var _lhs344 *GlobalState = globalState
			_ = _lhs344
			var _lhs345 *GlobalState = globalState
			_ = _lhs345
			_lhs344.F1 = _rhs405
			_2381_v18 = _rhs406
			_lhs345.F7 = _rhs407
		} else {
			var _2382_v19 _dafny.MultiSet
			_ = _2382_v19
			_2382_v19 = _dafny.MultiSetOf((_this).F11())
			var _2383_v21 _dafny.Map
			_ = _2383_v21
			_2383_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2382_v19, func() _dafny.Map {
				var _coll87 = _dafny.NewMapBuilder()
				_ = _coll87
				for _iter91 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-821), _dafny.IntOfInt64(806))); ; {
					_compr_87, _ok91 := _iter91()
					if !_ok91 {
						break
					}
					var _2384_v20 _dafny.Int
					_2384_v20 = interface{}(_compr_87).(_dafny.Int)
					if ((_dafny.IntOfInt64(-821)).Cmp(_2384_v20) <= 0) && ((_2384_v20).Cmp(_dafny.IntOfInt64(806)) < 0) {
						_coll87.Add((_2384_v20).Plus((_this).F11()), p1)
					}
				}
				return _coll87.ToMap()
			}())
			var _2385_v22 _dafny.Array
			_ = _2385_v22
			var _nw391 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
			_ = _nw391
			_2385_v22 = _nw391
			var _2386_v23 _dafny.Sequence
			_ = _2386_v23
			_2386_v23 = _dafny.SeqOf(_2385_v22)
			(globalState).F7 = (_this).Fm4(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-837))).Uint32(), func(coer119 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg119 _dafny.Int) interface{} {
					return coer119(arg119)
				}
			}(func(_2387_i1 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(((_this).F12()).Cardinality())
			}))).Cardinality()), _2383_v21, Companion_Default___.SafeModuloInt((_this).F11(), (_this).F11()), _dafny.Companion_Sequence_.IsProperPrefixOf(_2386_v23, _2386_v23), globalState)
			(globalState).F1 = (_this).Fm8(_2362_v1, globalState)
			var _2388_v24 _dafny.Sequence
			_ = _2388_v24
			_2388_v24 = _dafny.SeqOf(_dafny.IntOfInt64(25))
			_2388_v24 = _2388_v24
			var _2389_v25 D6
			_ = _2389_v25
			_2389_v25 = Companion_D6_.Create_DC11_(p2, _dafny.IntOfInt64(647), (_this).F11())
			var _2390_v26 _dafny.Map
			_ = _2390_v26
			_2390_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm34(p3, _2389_v25, globalState)).Cardinality()), _dafny.IntOfInt64(60))
			var _2391_v27 _dafny.Sequence
			_ = _2391_v27
			_2391_v27 = _dafny.SeqOf(_2390_v26)
			var _2392_v28 _dafny.Set
			_ = _2392_v28
			_2392_v28 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(284))).Uint32(), func(coer120 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg120 _dafny.Int) interface{} {
					return coer120(arg120)
				}
			}((func(_2393_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2394_i2 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf(_2393_p1)).Cardinality())
				}
			})(p1))), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(340), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(284))).Uint32(), func(coer121 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg121 _dafny.Int) interface{} {
					return coer121(arg121)
				}
			}((func(_2395_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2396_i2 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf(_2395_p1)).Cardinality())
				}
			})(p1)))).Cardinality()))).Uint32(), p2), _2388_v24)
			var _2397_v31 _dafny.Map
			_ = _2397_v31
			_2397_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2362_v1)
			var _2398_v32 _dafny.Map
			_ = _2398_v32
			_2398_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2364_v3).Cardinality()), (_2397_v31).Cardinality())).Update(p2, (_this).F11())).Update(p1, (_this).F11()))
			var _2399_v33 _dafny.Array
			_ = _2399_v33
			var _nwElement0_104 _dafny.Map = (_2391_v27).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2361_v0).Cardinality()), _dafny.IntOfUint32((_2391_v27).Cardinality()))).Uint32()).(_dafny.Map)
			_ = _nwElement0_104
			var _nw392 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_104, nil, _dafny.IntOfInt64(19))
			_ = _nw392
			(_nw392).ArraySet1(_nwElement0_104, 0)
			(_nw392).ArraySet1(_2390_v26, 1)
			(_nw392).ArraySet1((func() _dafny.Map {
				if p3 {
					return (_2390_v26).Update(_dafny.IntOfInt64(-527), (_2392_v28).Cardinality())
				}
				return _2390_v26
			})(), 2)
			(_nw392).ArraySet1((_2390_v26).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)), 3)
			(_nw392).ArraySet1(func() _dafny.Map {
				var _coll88 = _dafny.NewMapBuilder()
				_ = _coll88
				for _iter92 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(760), _dafny.IntOfInt64(727))); ; {
					_compr_88, _ok92 := _iter92()
					if !_ok92 {
						break
					}
					var _2400_v29 _dafny.Int
					_2400_v29 = interface{}(_compr_88).(_dafny.Int)
					if ((_dafny.IntOfInt64(760)).Cmp(_2400_v29) <= 0) && ((_2400_v29).Cmp(_dafny.IntOfInt64(727)) < 0) {
						_coll88.Add((_2400_v29).Plus((_this).F11()), _dafny.IntOfInt64(86))
					}
				}
				return _coll88.ToMap()
			}(), 4)
			(_nw392).ArraySet1((_2390_v26).Update((func() _dafny.Int {
				if (_2382_v19).Contains(p2) {
					return (_2382_v19).Multiplicity(p2)
				}
				return _dafny.IntOfInt64(103)
			})(), p2), 5)
			(_nw392).ArraySet1(func() _dafny.Map {
				var _coll89 = _dafny.NewMapBuilder()
				_ = _coll89
				for _iter93 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(2), _dafny.IntOfInt64(656))); ; {
					_compr_89, _ok93 := _iter93()
					if !_ok93 {
						break
					}
					var _2401_v30 _dafny.Int
					_2401_v30 = interface{}(_compr_89).(_dafny.Int)
					if ((_dafny.IntOfInt64(2)).Cmp(_2401_v30) <= 0) && ((_2401_v30).Cmp(_dafny.IntOfInt64(656)) < 0) {
						_coll89.Add(Companion_Default___.SafeModuloInt(_2401_v30, (_this).F11()), _dafny.IntOfInt64(533))
					}
				}
				return _coll89.ToMap()
			}(), 6)
			(_nw392).ArraySet1(Companion_Default___.Fm46(globalState), 7)
			(_nw392).ArraySet1((func() _dafny.Map {
				if !(false) {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(254))).Uint32(), func(coer122 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg122 _dafny.Int) interface{} {
							return coer122(arg122)
						}
					}(func(_2402_i3 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('w')
					}))).Cardinality())))
				}
				return _2390_v26
			})(), 8)
			(_nw392).ArraySet1(Companion_Default___.Fm46(globalState), 9)
			(_nw392).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), p2), 10)
			(_nw392).ArraySet1((_2390_v26).Merge(_2390_v26), 11)
			(_nw392).ArraySet1(_2390_v26, 12)
			(_nw392).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2398_v32).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(704))).Uint32(), func(coer123 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg123 _dafny.Int) interface{} {
					return coer123(arg123)
				}
			}((func(_2403_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2404_i4 _dafny.Int) _dafny.CodePoint {
					return _2403_v1
				}
			})(_2362_v1)))).Cardinality())), 13)
			(_nw392).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(p2, _dafny.IntOfInt64(493), globalState), _dafny.IntOfInt64(256))).Merge(_2390_v26), 14)
			(_nw392).ArraySet1((_2390_v26).Merge(_2390_v26), 15)
			(_nw392).ArraySet1(_2390_v26, 16)
			(_nw392).ArraySet1(_2390_v26, 17)
			(_nw392).ArraySet1(_2390_v26, 18)
			_2399_v33 = _nw392
			var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_2399_v33), 0))
			_ = _index339
			(_2399_v33).ArraySet1(_2390_v26, (_index339).Int())
			var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_2399_v33), 0))
			_ = _index340
			var _rhs408 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfInt64(157))).Merge((Companion_Default___.Fm46(globalState)).Merge(_2390_v26))
			_ = _rhs408
			var _rhs409 _dafny.Int = (_dafny.IntOfInt64(138)).Minus(Companion_Default___.SafeModuloInt(p2, (_dafny.Zero).Minus((_this).F11())))
			_ = _rhs409
			var _lhs346 _dafny.Array = _2399_v33
			_ = _lhs346
			var _lhs347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_2399_v33), 0))
			_ = _lhs347
			var _lhs348 *GlobalState = globalState
			_ = _lhs348
			(_lhs346).ArraySet1(_rhs408, (_lhs347).Int())
			_lhs348.F1 = _rhs409
			var _2405_v34 D6
			_ = _2405_v34
			_2405_v34 = Companion_D6_.Create_DC10_(_2361_v0)
			var _2406_v35 D19
			_ = _2406_v35
			_2406_v35 = Companion_D19_.Create_DC43_(p3, _2362_v1, p3, _2405_v34, _2364_v3)
			var _2407_v36 _dafny.Map
			_ = _2407_v36
			_2407_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _2364_v3)
			var _2408_v37 _dafny.Array
			_ = _2408_v37
			var _nwElement0_105 _dafny.Sequence = (func() _dafny.Sequence {
				if p0 {
					return _dafny.UnicodeSeqOfUtf8Bytes("nowe")
				}
				return _2364_v3
			})()
			_ = _nwElement0_105
			var _nw393 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_105, nil, _dafny.IntOfInt64(16))
			_ = _nw393
			(_nw393).ArraySet1(_nwElement0_105, 0)
			(_nw393).ArraySet1(_2364_v3, 1)
			(_nw393).ArraySet1(_2364_v3, 2)
			(_nw393).ArraySet1(_2364_v3, 3)
			(_nw393).ArraySet1(_2364_v3, 4)
			(_nw393).ArraySet1(_2364_v3, 5)
			(_nw393).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2364_v3, _2364_v3), 6)
			(_nw393).ArraySet1(_2364_v3, 7)
			(_nw393).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(115))).Uint32(), func(coer124 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg124 _dafny.Int) interface{} {
					return coer124(arg124)
				}
			}((func(_2409_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2410_i5 _dafny.Int) _dafny.CodePoint {
					return _2409_v1
				}
			})(_2362_v1))), (_2406_v35).Dtor_cf73()), 8)
			(_nw393).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2364_v3, _2364_v3), 9)
			(_nw393).ArraySet1((func() _dafny.Sequence {
				if (_2407_v36).Contains(p2) {
					return (_2407_v36).Get(p2).(_dafny.Sequence)
				}
				return _2364_v3
			})(), 10)
			(_nw393).ArraySet1(_2364_v3, 11)
			(_nw393).ArraySet1(_2364_v3, 12)
			(_nw393).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("lcgs"), 13)
			(_nw393).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(910))).Uint32(), func(coer125 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg125 _dafny.Int) interface{} {
					return coer125(arg125)
				}
			}((func(_2411_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2412_i6 _dafny.Int) _dafny.CodePoint {
					return _2411_v1
				}
			})(_2362_v1))), _2364_v3), 14)
			(_nw393).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2364_v3, _2364_v3), 15)
			_2408_v37 = _nw393
			var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_2408_v37), 0))
			_ = _index341
			(_2408_v37).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2364_v3, _dafny.UnicodeSeqOfUtf8Bytes("dhbyei")), (_index341).Int())
			var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_2408_v37), 0))
			_ = _index342
			(_2408_v37).ArraySet1(_2364_v3, (_index342).Int())
		}
		if p3 {
			var _2413_v38 _dafny.MultiSet
			_ = _2413_v38
			_2413_v38 = _dafny.MultiSetOf(p0, false)
			var _2414_v39 T1
			_ = _2414_v39
			var _nw394 *C3 = New_C3_()
			_ = _nw394
			_nw394.Ctor__((_this).F11(), (func() _dafny.Int {
				if (_2413_v38).Contains((_this).F13()) {
					return (_2413_v38).Multiplicity((_this).F13())
				}
				return p2
			})(), (_this).F11(), _dafny.SetOf(p3))
			_2414_v39 = _nw394
			var _2415_v40 _dafny.Map
			_ = _2415_v40
			_2415_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2414_v39, (_this).F11())
			var _2416_v41 _dafny.MultiSet
			_ = _2416_v41
			_2416_v41 = _dafny.MultiSetOf(_2415_v40)
			var _2417_v42 _dafny.MultiSet
			_ = _2417_v42
			_2417_v42 = _dafny.MultiSetOf((_2414_v39).F11())
			var _2418_v43 _dafny.Array
			_ = _2418_v43
			var _nwElement0_106 _dafny.Int = (func() _dafny.Int {
				if (_2416_v41).Contains(_2415_v40) {
					return (_2416_v41).Multiplicity(_2415_v40)
				}
				return (_2414_v39).F11()
			})()
			_ = _nwElement0_106
			var _nw395 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_106, nil, _dafny.IntOfInt64(22))
			_ = _nw395
			(_nw395).ArraySet1(_nwElement0_106, 0)
			(_nw395).ArraySet1((_2414_v39).F11(), 1)
			(_nw395).ArraySet1(p2, 2)
			(_nw395).ArraySet1(Companion_Default___.SafeDivisionInt((_2414_v39).F11(), (_this).F11()), 3)
			(_nw395).ArraySet1((_this).F11(), 4)
			(_nw395).ArraySet1((_this).F11(), 5)
			(_nw395).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_2364_v3, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2364_v3).Cardinality()))).Uint32(), _2362_v1), _dafny.UnicodeSeqOfUtf8Bytes("lysmecy"))).Cardinality()), 6)
			(_nw395).ArraySet1(p1, 7)
			(_nw395).ArraySet1(((_this).Fm7((_2417_v42).Cardinality(), (_this).F13(), _2413_v38, globalState)).Minus((_this).F11()), 8)
			(_nw395).ArraySet1(p2, 9)
			(_nw395).ArraySet1((func() _dafny.Int {
				if (_this).F13() {
					return (_dafny.Zero).Minus(p1)
				}
				return _dafny.IntOfInt64(710)
			})(), 10)
			(_nw395).ArraySet1((_2414_v39).F11(), 11)
			(_nw395).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-666), (_this).F11()), 12)
			(_nw395).ArraySet1(p2, 13)
			(_nw395).ArraySet1(_dafny.IntOfInt64(93), 14)
			(_nw395).ArraySet1(Companion_Default___.SafeModuloInt(p1, _dafny.IntOfUint32((_2364_v3).Cardinality())), 15)
			(_nw395).ArraySet1(_dafny.IntOfInt64(-212), 16)
			(_nw395).ArraySet1((_2414_v39).F11(), 17)
			(_nw395).ArraySet1((_this).F11(), 18)
			(_nw395).ArraySet1(((_2414_v39).F11()).Times(_dafny.IntOfInt64(-92)), 19)
			(_nw395).ArraySet1((_dafny.IntOfInt64(-196)).Plus(p1), 20)
			(_nw395).ArraySet1((_this).F11(), 21)
			_2418_v43 = _nw395
			_2418_v43 = _2418_v43
			var _2419_v44 _dafny.Array
			_ = _2419_v44
			var _2420_v45 _dafny.Map
			_ = _2420_v45
			var _out41 _dafny.Array
			_ = _out41
			var _out42 _dafny.Map
			_ = _out42
			_out41, _out42 = (_2414_v39).M1(p1, globalState)
			_2419_v44 = _out41
			_2420_v45 = _out42
			var _rhs410 _dafny.Int = _dafny.IntOfInt64(-437)
			_ = _rhs410
			var _rhs411 _dafny.CodePoint = _2362_v1
			_ = _rhs411
			var _lhs349 *GlobalState = globalState
			_ = _lhs349
			_lhs349.F1 = _rhs410
			_2362_v1 = _rhs411
			(globalState).F1 = (_dafny.IntOfInt64(939)).Times(p2)
			(globalState).F1 = _dafny.IntOfUint32((_dafny.SeqOf(p3)).Cardinality())
		} else {
			var _2421_v47 _dafny.MultiSet
			_ = _2421_v47
			_2421_v47 = _dafny.MultiSetOf(_2364_v3)
			var _2422_v48 _dafny.Map
			_ = _2422_v48
			_2422_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (func() _dafny.Map {
				var _coll90 = _dafny.NewMapBuilder()
				_ = _coll90
				for _iter94 := _dafny.Iterate((_2421_v47).Elements()); ; {
					_compr_90, _ok94 := _iter94()
					if !_ok94 {
						break
					}
					var _2423_v46 _dafny.Sequence
					_2423_v46 = interface{}(_compr_90).(_dafny.Sequence)
					if (_2421_v47).Contains(_2423_v46) {
						_coll90.Add(_2423_v46, false)
					}
				}
				return _coll90.ToMap()
			}()).Cardinality())
			var _2424_v49 _dafny.Sequence
			_ = _2424_v49
			_2424_v49 = _dafny.SeqOf((_2422_v48).Cardinality())
			(globalState).F1 = ((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).Fm8(Companion_Default___.Fm20(_2361_v0, _dafny.IntOfUint32((_2424_v49).Cardinality()), p0, globalState), globalState), p1)))).Times(p1)
			(globalState).F1 = (_this).F11()
			var _2425_v50 D2
			_ = _2425_v50
			_2425_v50 = Companion_D2_.Create_DC4_(Companion_Default___.Fm2((_this).F11(), (_dafny.Zero).Minus(p1), globalState), p1, p2)
			var _2426_v51 *C6
			_ = _2426_v51
			var _nw396 *C6 = New_C6_()
			_ = _nw396
			_nw396.Ctor__(_2425_v50, (Companion_Default___.Fm41(globalState)).Union((_this).F12()), p1)
			_2426_v51 = _nw396
			var _2427_v52 _dafny.Map
			_ = _2427_v52
			_2427_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2422_v48)
			var _2428_v53 _dafny.Map
			_ = _2428_v53
			_2428_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_2422_v48).Update(p0, p1))).Merge(_2427_v52)).Cardinality(), _2426_v51)
			_2426_v51 = (func() *C6 {
				if (_2428_v53).Contains((_this).F11()) {
					return (_2428_v53).Get((_this).F11()).(*C6)
				}
				return _2426_v51
			})()
			var _2429_v54 _dafny.Array
			_ = _2429_v54
			var _nw397 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw397
			_2429_v54 = _nw397
			var _2430_v55 _dafny.Map
			_ = _2430_v55
			_2430_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2429_v54, p0)
			var _2431_v56 D10
			_ = _2431_v56
			_2431_v56 = Companion_D10_.Create_DC20_(_2430_v55)
			var _pat_let_tv55 = _2430_v55
			_ = _pat_let_tv55
			var _pat_let_tv56 = _2430_v55
			_ = _pat_let_tv56
			var _source33 D10 = func(_pat_let60_0 D10) D10 {
				return func(_2434_dt__update__tmp_h1 D10) D10 {
					return func(_pat_let63_0 _dafny.Map) D10 {
						return func(_2435_dt__update_hcf27_h1 _dafny.Map) D10 {
							return Companion_D10_.Create_DC20_(_2435_dt__update_hcf27_h1)
						}(_pat_let63_0)
					}(_pat_let_tv56)
				}(_pat_let60_0)
			}(func(_pat_let61_0 D10) D10 {
				return func(_2432_dt__update__tmp_h0 D10) D10 {
					return func(_pat_let62_0 _dafny.Map) D10 {
						return func(_2433_dt__update_hcf27_h0 _dafny.Map) D10 {
							return Companion_D10_.Create_DC20_(_2433_dt__update_hcf27_h0)
						}(_pat_let62_0)
					}(_pat_let_tv55)
				}(_pat_let61_0)
			}(_2431_v56))
			_ = _source33
			if _source33.Is_DC21() {
				var _2436___mcc_h0 bool = _source33.Get_().(D10_DC21).Cf28
				_ = _2436___mcc_h0
				var _2437___mcc_h1 bool = _source33.Get_().(D10_DC21).Cf29
				_ = _2437___mcc_h1
				var _2438_cf29 bool = _2437___mcc_h1
				_ = _2438_cf29
				var _2439_cf28 bool = _2436___mcc_h0
				_ = _2439_cf28
				(globalState).F1 = _dafny.IntOfInt64(178)
				var _2440_v57 _dafny.Array
				_ = _2440_v57
				var _nw398 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw398
				_2440_v57 = _nw398
				var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_2440_v57), 0))
				_ = _index343
				(_2440_v57).ArraySet1((_dafny.IntOfInt64(532)).Cmp((_this).F11()) < 0, (_index343).Int())
				var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_2440_v57), 0))
				_ = _index344
				(_2440_v57).ArraySet1(!_dafny.Companion_Sequence_.Equal(_2364_v3, _dafny.Companion_Sequence_.Concatenate(_2364_v3, _2364_v3)), (_index344).Int())
				(globalState).F6 = _2439_cf28
				var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_2429_v54), 0))
				_ = _index345
				(_2429_v54).ArraySet1((_this).F11(), (_index345).Int())
				var _2441_v58 _dafny.Array
				_ = _2441_v58
				var _nwElement0_107 _dafny.Array = _2440_v57
				_ = _nwElement0_107
				var _nw399 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_107, nil, _dafny.IntOfInt64(27))
				_ = _nw399
				(_nw399).ArraySet1(_nwElement0_107, 0)
				(_nw399).ArraySet1(_2440_v57, 1)
				(_nw399).ArraySet1(_2440_v57, 2)
				(_nw399).ArraySet1(_2440_v57, 3)
				(_nw399).ArraySet1(_2440_v57, 4)
				(_nw399).ArraySet1(_2440_v57, 5)
				(_nw399).ArraySet1(_2440_v57, 6)
				(_nw399).ArraySet1(_2440_v57, 7)
				(_nw399).ArraySet1(_2440_v57, 8)
				(_nw399).ArraySet1(_2440_v57, 9)
				(_nw399).ArraySet1(_2440_v57, 10)
				(_nw399).ArraySet1(_2440_v57, 11)
				(_nw399).ArraySet1(_2440_v57, 12)
				(_nw399).ArraySet1(_2440_v57, 13)
				(_nw399).ArraySet1(_2440_v57, 14)
				(_nw399).ArraySet1(_2440_v57, 15)
				(_nw399).ArraySet1(_2440_v57, 16)
				(_nw399).ArraySet1(_2440_v57, 17)
				(_nw399).ArraySet1(_2440_v57, 18)
				(_nw399).ArraySet1(_2440_v57, 19)
				(_nw399).ArraySet1(_2440_v57, 20)
				(_nw399).ArraySet1(_2440_v57, 21)
				(_nw399).ArraySet1(_2440_v57, 22)
				(_nw399).ArraySet1(_2440_v57, 23)
				(_nw399).ArraySet1(_2440_v57, 24)
				(_nw399).ArraySet1(_2440_v57, 25)
				(_nw399).ArraySet1(_2440_v57, 26)
				_2441_v58 = _nw399
				var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_2429_v54), 0))
				_ = _index346
				var _rhs412 _dafny.Int = p2
				_ = _rhs412
				var _rhs413 _dafny.Array = _2441_v58
				_ = _rhs413
				var _rhs414 _dafny.Array = _2429_v54
				_ = _rhs414
				var _lhs350 _dafny.Array = _2429_v54
				_ = _lhs350
				var _lhs351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_2429_v54), 0))
				_ = _lhs351
				(_lhs350).ArraySet1(_rhs412, (_lhs351).Int())
				_2441_v58 = _rhs413
				_2429_v54 = _rhs414
			} else {
				var _2442___mcc_h2 _dafny.Map = _source33.Get_().(D10_DC20).Cf27
				_ = _2442___mcc_h2
				var _2443_cf27 _dafny.Map = _2442___mcc_h2
				_ = _2443_cf27
				var _2444_v59 _dafny.MultiSet
				_ = _2444_v59
				_2444_v59 = _dafny.MultiSetOf((_this).F11(), (_this).F11())
				var _2445_v60 *C4
				_ = _2445_v60
				var _nw400 *C4 = New_C4_()
				_ = _nw400
				_nw400.Ctor__(_2444_v59, (_this).F13(), (_this).F12(), _dafny.IntOfInt64(587))
				_2445_v60 = _nw400
				var _2446_v61 _dafny.Sequence
				_ = _2446_v61
				_2446_v61 = _dafny.SeqOf(_2445_v60, _2445_v60)
				var _2447_v62 _dafny.Map
				_ = _2447_v62
				_2447_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_2446_v61).Cardinality())).Times(_dafny.IntOfUint32((_2364_v3).Cardinality())), ((_2424_v49).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2424_v49).Cardinality()))).Uint32()).(_dafny.Int)).Plus(p1))
				_2447_v62 = (_2447_v62).Update(p2, p2)
				var _2448_v63 _dafny.Map
				_ = _2448_v63
				_2448_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2445_v60).F16(), _2447_v62)
				var _2449_v64 _dafny.Array
				_ = _2449_v64
				var _nwElement0_108 bool = (_2445_v60).F17()
				_ = _nwElement0_108
				var _nw401 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_108, nil, _dafny.IntOfInt64(29))
				_ = _nw401
				(_nw401).ArraySet1(_nwElement0_108, 0)
				(_nw401).ArraySet1(p3, 1)
				(_nw401).ArraySet1(false, 2)
				(_nw401).ArraySet1((p2).Cmp((_2422_v48).Cardinality()) <= 0, 3)
				(_nw401).ArraySet1(true, 4)
				(_nw401).ArraySet1(p3, 5)
				(_nw401).ArraySet1((_2445_v60).F17(), 6)
				(_nw401).ArraySet1((_2445_v60).F17(), 7)
				(_nw401).ArraySet1(!((p2).Cmp((_this).F11()) <= 0), 8)
				(_nw401).ArraySet1((p1).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality())))) == 0, 9)
				(_nw401).ArraySet1(!((_2445_v60).F17()), 10)
				(_nw401).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_2364_v3, _2364_v3), 11)
				(_nw401).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_2364_v3, _dafny.Companion_Sequence_.Update(_2364_v3, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2364_v3).Cardinality()))).Uint32(), _2362_v1)), 12)
				(_nw401).ArraySet1(!(p3) || (p0), 13)
				(_nw401).ArraySet1(p0, 14)
				(_nw401).ArraySet1(p3, 15)
				(_nw401).ArraySet1(p0, 16)
				(_nw401).ArraySet1((_2445_v60).Fm4(p2, _2448_v63, p2, p3, globalState), 17)
				(_nw401).ArraySet1((p1).Cmp(_dafny.IntOfInt64(674)) != 0, 18)
				(_nw401).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("xwylrsvr"), _dafny.UnicodeSeqOfUtf8Bytes("rk")), 19)
				(_nw401).ArraySet1(false, 20)
				(_nw401).ArraySet1(p3, 21)
				(_nw401).ArraySet1((func() bool {
					if p3 {
						return p3
					}
					return (_this).F13()
				})(), 22)
				(_nw401).ArraySet1((_2445_v60).F17(), 23)
				(_nw401).ArraySet1(((_this).F11()).Cmp((_dafny.Zero).Minus((_this).F11())) != 0, 24)
				(_nw401).ArraySet1((_this).F13(), 25)
				(_nw401).ArraySet1(p3, 26)
				(_nw401).ArraySet1(((_this).F13()) == ((_2445_v60).F17()), 27)
				(_nw401).ArraySet1((_this).F13(), 28)
				_2449_v64 = _nw401
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_2449_v64), 0))
				_ = _index347
				(_2449_v64).ArraySet1(p0, (_index347).Int())
				var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_2449_v64), 0))
				_ = _index348
				(_2449_v64).ArraySet1((p3) || ((_2445_v60).F17()), (_index348).Int())
				(globalState).F1 = p2
				var _2450_v65 T0
				_ = _2450_v65
				var _nw402 *C3 = New_C3_()
				_ = _nw402
				_nw402.Ctor__(p2, (_dafny.Zero).Minus(p1), p2, (_this).F12())
				_2450_v65 = _nw402
				_2450_v65 = _2450_v65
			}
			var _2451_v66 _dafny.Array
			_ = _2451_v66
			var _nw403 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(13))
			_ = _nw403
			_2451_v66 = _nw403
			var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_2451_v66), 0))
			_ = _index349
			(_2451_v66).ArraySet1(_dafny.SetOf(p0, true, p0), (_index349).Int())
			var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_2451_v66), 0))
			_ = _index350
			(_2451_v66).ArraySet1((_this).F12(), (_index350).Int())
		}
		_2362_v1 = _2362_v1
		(globalState).F6 = !(p0)
		r0 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(133))).Uint32(), func(coer126 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg126 _dafny.Int) interface{} {
				return coer126(arg126)
			}
		}((func(_2452_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2453_i7 _dafny.Int) _dafny.CodePoint {
				return _2452_v1
			}
		})(_2362_v1))), _2364_v3)
		return r0
	}
}
func (_this *C8) F13() bool {
	{
		return _this._f13
	}
}

// End of class C8
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
