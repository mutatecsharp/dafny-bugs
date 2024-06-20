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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	return !((_dafny.MultiSetOf(true)).IsDisjointFrom(_dafny.MultiSetOf(false, false, !(false), false))) || ((false) || (true))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	if true {
		return _dafny.UnicodeSeqOfUtf8Bytes("hgwprn")
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(792))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})), _dafny.UnicodeSeqOfUtf8Bytes("tsxqnnw"))
	}
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((_dafny.IntOfUint32(((func() _dafny.Sequence {
		if !(true) {
			return _dafny.UnicodeSeqOfUtf8Bytes("fv")
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("uosyjyo")
	})()).Cardinality())).Times((Companion_D4_.Create_DC15_(!(false), (_dafny.MultiSetOf(true)).Cardinality())).Dtor_cf24()))
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(!(false), false, false)).Union(_dafny.SetOf(true))).Difference((_dafny.SetOf(true, true)).Intersection(_dafny.SetOf(false, !(true))))
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC7_(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()).Minus(_dafny.IntOfInt64(242)))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.MultiSet, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D1_.Create_DC3_(false, false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(627))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	})), false, _dafny.IntOfInt64(45))).Dtor_cf5(), _dafny.IntOfInt64(7))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(177)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-120))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.MultiSetOf(false)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 D3, p2 _dafny.Set, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(59), _dafny.IntOfInt64(-991))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _2_v0 _dafny.Int
			_2_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(59)).Cmp(_2_v0) <= 0) && ((_2_v0).Cmp(_dafny.IntOfInt64(-991)) < 0) {
				_coll0.Add(Companion_Default___.SafeDivisionInt(_2_v0, _dafny.IntOfInt64(958)), (true) || (true))
			}
		}
		return _coll0.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(500), _dafny.IntOfInt64(122))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _3_v0 _dafny.Int
			_3_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(500)).Cmp(_3_v0) <= 0) && ((_3_v0).Cmp(_dafny.IntOfInt64(122)) < 0) {
				_coll1.Add((_3_v0).Times(_dafny.IntOfInt64(-706)))
			}
		}
		return _coll1.ToSet()
	}()).Difference(func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-473), _dafny.IntOfInt64(223))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _4_v1 _dafny.Int
			_4_v1 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(-473)).Cmp(_4_v1) <= 0) && ((_4_v1).Cmp(_dafny.IntOfInt64(223)) < 0) {
				_coll2.Add(Companion_Default___.SafeModuloInt(_4_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(461), true)).Cardinality()))
			}
		}
		return _coll2.ToSet()
	}())).Union(_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-290))))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	if false {
		if false {
			return _dafny.SeqOf(false, !(!(false)))
		} else {
			return _dafny.SeqOf(false)
		}
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false, (Companion_D6_.Create_DC20_(true)).Dtor_cf34()), _dafny.SeqOf(true))
	}
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D3_.Create_DC12_(Companion_D3_.Create_DC10_(false, _dafny.IntOfInt64(524), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(380))).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D3_.Create_DC12_(Companion_D3_.Create_DC12_(Companion_D3_.Create_DC11_()))))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D3_.Create_DC12_(Companion_D3_.Create_DC10_(false, _dafny.IntOfInt64(682), (func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(337), _dafny.IntOfInt64(918))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _5_v0 _dafny.Int
			_5_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(337)).Cmp(_5_v0) <= 0) && ((_5_v0).Cmp(_dafny.IntOfInt64(918)) < 0) {
				_coll3.Add(Companion_Default___.SafeDivisionInt(_5_v0, _dafny.IntOfInt64(99)))
			}
		}
		return _coll3.ToSet()
	}()).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D3_.Create_DC12_(Companion_D3_.Create_DC9_(_dafny.SeqOf((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())).Cardinality()))))))
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) D3 {
	var _source0 D1 = (func() D1 {
		if true {
			return Companion_D1_.Create_DC2_(_dafny.IntOfInt64(376))
		}
		return Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("arg")).Cardinality()))
	})()
	_ = _source0
	if _source0.Is_DC3() {
		var _6___mcc_h0 bool = _source0.Get_().(D1_DC3).Cf5
		_ = _6___mcc_h0
		var _7___mcc_h1 bool = _source0.Get_().(D1_DC3).Cf6
		_ = _7___mcc_h1
		var _8___mcc_h2 _dafny.Sequence = _source0.Get_().(D1_DC3).Cf7
		_ = _8___mcc_h2
		var _9___mcc_h3 bool = _source0.Get_().(D1_DC3).Cf8
		_ = _9___mcc_h3
		var _10___mcc_h4 _dafny.Int = _source0.Get_().(D1_DC3).Cf9
		_ = _10___mcc_h4
		var _11_cf9 _dafny.Int = _10___mcc_h4
		_ = _11_cf9
		var _12_cf8 bool = _9___mcc_h3
		_ = _12_cf8
		var _13_cf7 _dafny.Sequence = _8___mcc_h2
		_ = _13_cf7
		var _14_cf6 bool = _7___mcc_h1
		_ = _14_cf6
		var _15_cf5 bool = _6___mcc_h0
		_ = _15_cf5
		return Companion_D3_.Create_DC12_(Companion_D3_.Create_DC10_(_12_cf8, _11_cf9, _11_cf9))
	} else if _source0.Is_DC2() {
		var _16___mcc_h5 _dafny.Int = _source0.Get_().(D1_DC2).Cf4
		_ = _16___mcc_h5
		var _17_cf4 _dafny.Int = _16___mcc_h5
		_ = _17_cf4
		return Companion_D3_.Create_DC12_(Companion_D3_.Create_DC11_())
	} else {
		var _18___mcc_h6 D1 = _source0.Get_().(D1_DC4).Cf10
		_ = _18___mcc_h6
		var _19_cf10 D1 = _18___mcc_h6
		_ = _19_cf10
		return Companion_D3_.Create_DC12_(Companion_D3_.Create_DC12_(Companion_D3_.Create_DC12_(Companion_D3_.Create_DC9_(_dafny.SeqOf(_dafny.IntOfInt64(27))))))
	}
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(254), _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nk")).Cardinality()), (_dafny.SetOf(false, true)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(240), func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(937), _dafny.IntOfInt64(842))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _20_v0 _dafny.Int
			_20_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(937)).Cmp(_20_v0) <= 0) && ((_20_v0).Cmp(_dafny.IntOfInt64(842)) < 0) {
				_coll4.Add(Companion_Default___.SafeModuloInt(_20_v0, (_dafny.Zero).Minus((_dafny.MultiSetOf(true, true)).Cardinality())))
			}
		}
		return _coll4.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm14(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_(Companion_D1_.Create_DC4_(Companion_D1_.Create_DC3_(true, false, _dafny.UnicodeSeqOfUtf8Bytes("eewew"), true, _dafny.IntOfInt64(243))))
}
func (_static *CompanionStruct_Default___) Fm15(p0 D3, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetFromSeq(_dafny.SeqOf(true, false, false, false, false))).Intersection((_dafny.MultiSetOf(false)).Union(_dafny.MultiSetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) D4 {
	var _source1 D6 = Companion_D6_.Create_DC20_(true)
	_ = _source1
	if _source1.Is_DC19() {
		var _21___mcc_h0 _dafny.Int = _source1.Get_().(D6_DC19).Cf31
		_ = _21___mcc_h0
		var _22___mcc_h1 bool = _source1.Get_().(D6_DC19).Cf32
		_ = _22___mcc_h1
		var _23___mcc_h2 _dafny.CodePoint = _source1.Get_().(D6_DC19).Cf33
		_ = _23___mcc_h2
		var _24_cf33 _dafny.CodePoint = _23___mcc_h2
		_ = _24_cf33
		var _25_cf32 bool = _22___mcc_h1
		_ = _25_cf32
		var _26_cf31 _dafny.Int = _21___mcc_h0
		_ = _26_cf31
		return Companion_D4_.Create_DC15_(_25_cf32, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (Companion_D0_.Create_DC0_(_25_cf32)).Dtor_cf0())).Cardinality())
	} else if _source1.Is_DC20() {
		var _27___mcc_h3 bool = _source1.Get_().(D6_DC20).Cf34
		_ = _27___mcc_h3
		var _28_cf34 bool = _27___mcc_h3
		_ = _28_cf34
		return Companion_D4_.Create_DC15_(_28_cf34, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_28_cf34)).Cardinality()), (_dafny.MultiSetOf(_dafny.CodePoint('y'), _dafny.CodePoint('m'), _dafny.CodePoint('t'), _dafny.CodePoint('q'), _dafny.CodePoint('a'))).Cardinality())).Cardinality()), _dafny.MultiSetOf(_dafny.IntOfInt64(673), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_cf34, true)).Cardinality()))).Cardinality()))
	} else if _source1.Is_DC21() {
		return Companion_D4_.Create_DC15_(true, (_dafny.SetOf((_dafny.Zero).Minus((_dafny.SetOf((_dafny.SetOf(true)).Cardinality(), (func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((func() _dafny.Set {
				var _coll6 = _dafny.NewBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(572), _dafny.IntOfInt64(240))); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _29_v1 _dafny.Int
					_29_v1 = interface{}(_compr_6).(_dafny.Int)
					if ((_dafny.IntOfInt64(572)).Cmp(_29_v1) <= 0) && ((_29_v1).Cmp(_dafny.IntOfInt64(240)) < 0) {
						_coll6.Add((_29_v1).Times(_dafny.IntOfInt64(414)))
					}
				}
				return _coll6.ToSet()
			}()).Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _30_v0 _dafny.Int
				_30_v0 = interface{}(_compr_5).(_dafny.Int)
				if (func() _dafny.Set {
					var _coll7 = _dafny.NewBuilder()
					_ = _coll7
					for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(572), _dafny.IntOfInt64(240))); ; {
						_compr_7, _ok7 := _iter7()
						if !_ok7 {
							break
						}
						var _31_v1 _dafny.Int
						_31_v1 = interface{}(_compr_7).(_dafny.Int)
						if ((_dafny.IntOfInt64(572)).Cmp(_31_v1) <= 0) && ((_31_v1).Cmp(_dafny.IntOfInt64(240)) < 0) {
							_coll7.Add((_31_v1).Times(_dafny.IntOfInt64(414)))
						}
					}
					return _coll7.ToSet()
				}()).Contains(_30_v0) {
					_coll5.Add(Companion_Default___.SafeModuloInt(_30_v0, (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(691))).Cardinality())), true)
				}
			}
			return _coll5.ToMap()
		}()).Cardinality())).Cardinality()))).Cardinality())
	} else {
		var _32___mcc_h4 _dafny.CodePoint = _source1.Get_().(D6_DC18).Cf30
		_ = _32___mcc_h4
		var _33_cf30 _dafny.CodePoint = _32___mcc_h4
		_ = _33_cf30
		if true {
			return Companion_D4_.Create_DC15_(!(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(592))).Cardinality())
		} else {
			return Companion_D4_.Create_DC15_(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ngymo")).Cardinality()))
		}
	}
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(true, false)).Cardinality())).Cardinality()))).Union(_dafny.SetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(943))).Cardinality())), _dafny.SetOf(_dafny.IntOfInt64(-890), (func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(445), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(873), _dafny.IntOfInt64(28))).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jfd")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cngsgffno")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("giiveerwc")).Cardinality()))).Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _34_v0 _dafny.Int
			_34_v0 = interface{}(_compr_8).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(445), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(873), _dafny.IntOfInt64(28))).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jfd")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cngsgffno")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("giiveerwc")).Cardinality())), _34_v0) {
				_coll8.Add((_34_v0).Plus((Companion_D4_.Create_DC15_(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(112))).Uint32(), func(coer2 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg2 _dafny.Int) interface{} {
						return coer2(arg2)
					}
				}(func(_35_i0 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf((_dafny.MultiSetOf(true)).Cardinality())
				}))).Cardinality()))).Dtor_cf24()), _dafny.IntOfInt64(-492))
			}
		}
		return _coll8.ToMap()
	}()).Cardinality()), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(868))).Cardinality())), _dafny.SetOf(_dafny.IntOfInt64(472)), _dafny.SetOf((_dafny.SetOf((Companion_D4_.Create_DC15_(false, _dafny.IntOfInt64(655))).Dtor_cf24())).Cardinality(), _dafny.IntOfInt64(-685), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf((func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(233), _dafny.IntOfInt64(864))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _36_v1 _dafny.Int
			_36_v1 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(233)).Cmp(_36_v1) <= 0) && ((_36_v1).Cmp(_dafny.IntOfInt64(864)) < 0) {
				_coll9.Add(Companion_Default___.SafeModuloInt(_36_v1, _dafny.IntOfInt64(830)))
			}
		}
		return _coll9.ToSet()
	}()).Cardinality(), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()))).Cardinality()), true)).Cardinality(), (_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('h')
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(300))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_37_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-673))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_38_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		}))
	}))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, globalState *GlobalState) bool {
	var r0 bool = false
	_ = r0
	r0 = p0
	var _39_v0 D0
	_ = _39_v0
	_39_v0 = Companion_D0_.Create_DC0_(p0)
	var _40_i0 _dafny.Int
	_ = _40_i0
	_40_i0 = _dafny.Zero
	{
		var _pat_let_tv3 = p0
		_ = _pat_let_tv3
		var _pat_let_tv4 = p0
		_ = _pat_let_tv4
		for func(_source3 D0) bool {
			if _source3.Is_DC0() {
				var _99___mcc_h6 bool = _source3.Get_().(D0_DC0).Cf0
				_ = _99___mcc_h6
				var _100_cf0 bool = _99___mcc_h6
				_ = _100_cf0
				return _pat_let_tv3
			} else {
				var _101___mcc_h7 _dafny.MultiSet = _source3.Get_().(D0_DC1).Cf1
				_ = _101___mcc_h7
				var _102___mcc_h8 _dafny.Array = _source3.Get_().(D0_DC1).Cf2
				_ = _102___mcc_h8
				var _103___mcc_h9 _dafny.Array = _source3.Get_().(D0_DC1).Cf3
				_ = _103___mcc_h9
				var _104_cf3 _dafny.Array = _103___mcc_h9
				_ = _104_cf3
				var _105_cf2 _dafny.Array = _102___mcc_h8
				_ = _105_cf2
				var _106_cf1 _dafny.MultiSet = _101___mcc_h7
				_ = _106_cf1
				return ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-575))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg7 _dafny.Int) interface{} {
						return coer7(arg7)
					}
				}(func(_107_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				}))).Cardinality()))).Cmp((Companion_D1_.Create_DC3_(_pat_let_tv4, true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(431))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}(func(_108_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('r')
				})), false, (_dafny.Zero).Minus(_dafny.IntOfInt64(-332)))).Dtor_cf9()) == 0
			}
		}(_39_v0) {
			{
				if (_40_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_40_i0 = (_40_i0).Plus(_dafny.One)
				r0 = p0
				var _41_v1 _dafny.Int
				_ = _41_v1
				_41_v1 = _dafny.IntOfInt64(660)
				var _42_v2 _dafny.Sequence
				_ = _42_v2
				_42_v2 = _dafny.SeqOf(_41_v1, _41_v1)
				var _43_v3 _dafny.Array
				_ = _43_v3
				var _nwElement0_0 _dafny.Int = _41_v1
				_ = _nwElement0_0
				var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(3))
				_ = _nw0
				(_nw0).ArraySet1(_nwElement0_0, 0)
				(_nw0).ArraySet1(_41_v1, 1)
				(_nw0).ArraySet1(_dafny.IntOfUint32((_42_v2).Cardinality()), 2)
				_43_v3 = _nw0
				var _44_v4 D2
				_ = _44_v4
				_44_v4 = Companion_D2_.Create_DC6_(_41_v1, true, _43_v3)
				var _pat_let_tv0 = p0
				_ = _pat_let_tv0
				var _pat_let_tv1 = p0
				_ = _pat_let_tv1
				var _pat_let_tv2 = _43_v3
				_ = _pat_let_tv2
				var _source2 D2 = func(_pat_let0_0 D2) D2 {
					return func(_45_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let1_0 bool) D2 {
							return func(_46_dt__update_hcf13_h0 bool) D2 {
								return func(_pat_let2_0 _dafny.Array) D2 {
									return func(_47_dt__update_hcf14_h0 _dafny.Array) D2 {
										return Companion_D2_.Create_DC6_((_45_dt__update__tmp_h0).Dtor_cf12(), _46_dt__update_hcf13_h0, _47_dt__update_hcf14_h0)
									}(_pat_let2_0)
								}(_pat_let_tv2)
							}(_pat_let1_0)
						}((_pat_let_tv0) == (_pat_let_tv1))
					}(_pat_let0_0)
				}(_44_v4)
				_ = _source2
				if _source2.Is_DC6() {
					var _48___mcc_h0 _dafny.Int = _source2.Get_().(D2_DC6).Cf12
					_ = _48___mcc_h0
					var _49___mcc_h1 bool = _source2.Get_().(D2_DC6).Cf13
					_ = _49___mcc_h1
					var _50___mcc_h2 _dafny.Array = _source2.Get_().(D2_DC6).Cf14
					_ = _50___mcc_h2
					var _51_cf14 _dafny.Array = _50___mcc_h2
					_ = _51_cf14
					var _52_cf13 bool = _49___mcc_h1
					_ = _52_cf13
					var _53_cf12 _dafny.Int = _48___mcc_h0
					_ = _53_cf12
					var _54_v5 _dafny.Set
					_ = _54_v5
					_54_v5 = _dafny.SetOf(false)
					var _55_v6 *C0
					_ = _55_v6
					var _nw1 *C0 = New_C0_()
					_ = _nw1
					_nw1.Ctor__(_53_cf12, (func() _dafny.Set {
						if p0 {
							return _dafny.SetOf(!(p0), Companion_Default___.Fm0(!(p0), p0, _53_cf12, globalState), _52_cf13)
						}
						return _54_v5
					})())
					_55_v6 = _nw1
					var _56_v7 _dafny.Sequence
					_ = _56_v7
					_56_v7 = _dafny.SeqOf(p0, p0, (_dafny.SetOf(_52_cf13)).Contains(p0))
					r0 = (_56_v7).Select((Companion_Default___.SafeIndex(_53_cf12, _dafny.IntOfUint32((_56_v7).Cardinality()))).Uint32()).(bool)
					var _57_v8 *C0
					_ = _57_v8
					var _nw2 *C0 = New_C0_()
					_ = _nw2
					_nw2.Ctor__((_55_v6).F6(), _54_v5)
					_57_v8 = _nw2
					var _58_v9 _dafny.Set
					_ = _58_v9
					_58_v9 = _dafny.SetOf((_55_v6).F6(), _41_v1)
					var _59_v10 _dafny.Sequence
					_ = _59_v10
					_59_v10 = _dafny.SeqOf(_58_v9)
					var _60_v11 _dafny.Map
					_ = _60_v11
					_60_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_57_v8).F6(), _dafny.UnicodeSeqOfUtf8Bytes("upqhkypa"))
					var _61_v12 _dafny.Map
					_ = _61_v12
					_61_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_60_v11).Cardinality(), _58_v9)
					var _62_v15 _dafny.Map
					_ = _62_v15
					_62_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_57_v8).F6())
					var _63_v16 _dafny.Array
					_ = _63_v16
					var _nwElement0_1 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_59_v10, _dafny.SeqOf((func() _dafny.Set {
						if (_61_v12).Contains((_55_v6).F6()) {
							return (_61_v12).Get((_55_v6).F6()).(_dafny.Set)
						}
						return (_59_v10).Select((Companion_Default___.SafeIndex(_53_cf12, _dafny.IntOfUint32((_59_v10).Cardinality()))).Uint32()).(_dafny.Set)
					})()))
					_ = _nwElement0_1
					var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(15))
					_ = _nw3
					(_nw3).ArraySet1(_nwElement0_1, 0)
					(_nw3).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(169))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
						return func(arg5 _dafny.Int) interface{} {
							return coer5(arg5)
						}
					}((func(_64_cf12 _dafny.Int) func(_dafny.Int) _dafny.Set {
						return func(_65_i3 _dafny.Int) _dafny.Set {
							return func() _dafny.Set {
								var _coll10 = _dafny.NewBuilder()
								_ = _coll10
								for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(239), _dafny.IntOfInt64(328))); ; {
									_compr_10, _ok10 := _iter10()
									if !_ok10 {
										break
									}
									var _66_v13 _dafny.Int
									_66_v13 = interface{}(_compr_10).(_dafny.Int)
									if ((_dafny.IntOfInt64(239)).Cmp(_66_v13) <= 0) && ((_66_v13).Cmp(_dafny.IntOfInt64(328)) < 0) {
										_coll10.Add(Companion_Default___.SafeModuloInt(_66_v13, _64_cf12))
									}
								}
								return _coll10.ToSet()
							}()
						}
					})(_53_cf12))), 1)
					(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_59_v10, (Companion_Default___.SafeIndex((_57_v8).F6(), _dafny.IntOfUint32((_59_v10).Cardinality()))).Uint32(), _58_v9), _59_v10), 2)
					(_nw3).ArraySet1(_59_v10, 3)
					(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_59_v10, _59_v10), 4)
					(_nw3).ArraySet1(_59_v10, 5)
					(_nw3).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(845))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
						return func(arg6 _dafny.Int) interface{} {
							return coer6(arg6)
						}
					}((func(_67_v9 _dafny.Set) func(_dafny.Int) _dafny.Set {
						return func(_68_i4 _dafny.Int) _dafny.Set {
							return _67_v9
						}
					})(_58_v9))), 6)
					(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_58_v9), _dafny.SeqOf(_58_v9, _58_v9)), 7)
					(_nw3).ArraySet1(_dafny.Companion_Sequence_.Update(_59_v10, (Companion_Default___.SafeIndex((_55_v6).F6(), _dafny.IntOfUint32((_59_v10).Cardinality()))).Uint32(), func() _dafny.Set {
						var _coll11 = _dafny.NewBuilder()
						_ = _coll11
						for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-118), _dafny.IntOfInt64(257))); ; {
							_compr_11, _ok11 := _iter11()
							if !_ok11 {
								break
							}
							var _69_v14 _dafny.Int
							_69_v14 = interface{}(_compr_11).(_dafny.Int)
							if ((_dafny.IntOfInt64(-118)).Cmp(_69_v14) <= 0) && ((_69_v14).Cmp(_dafny.IntOfInt64(257)) < 0) {
								_coll11.Add(Companion_Default___.SafeDivisionInt(_69_v14, _dafny.IntOfUint32((_56_v7).Cardinality())))
							}
						}
						return _coll11.ToSet()
					}()), 8)
					(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_59_v10, _59_v10), 9)
					(_nw3).ArraySet1(_dafny.Companion_Sequence_.Update(_59_v10, (Companion_Default___.SafeIndex((_62_v15).Cardinality(), _dafny.IntOfUint32((_59_v10).Cardinality()))).Uint32(), _58_v9), 10)
					(_nw3).ArraySet1(_59_v10, 11)
					(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_59_v10, _59_v10), 12)
					(_nw3).ArraySet1(_59_v10, 13)
					(_nw3).ArraySet1(_59_v10, 14)
					_63_v16 = _nw3
					var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_63_v16), 0))
					_ = _index0
					(_63_v16).ArraySet1(_59_v10, (_index0).Int())
					var _70_v17 _dafny.MultiSet
					_ = _70_v17
					_70_v17 = _dafny.MultiSetOf(_53_cf12)
					var _71_v18 *C0
					_ = _71_v18
					var _nw4 *C0 = New_C0_()
					_ = _nw4
					_nw4.Ctor__(Companion_Default___.SafeDivisionInt((_57_v8).F6(), (func() _dafny.Int {
						if (_62_v15).Contains(p0) {
							return (_62_v15).Get(p0).(_dafny.Int)
						}
						return (_70_v17).Cardinality()
					})()), (_57_v8).F7())
					_71_v18 = _nw4
					var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_63_v16), 0))
					_ = _index1
					var _rhs0 _dafny.Sequence = _59_v10
					_ = _rhs0
					var _rhs1 *C0 = _71_v18
					_ = _rhs1
					var _rhs2 bool = p0
					_ = _rhs2
					var _lhs0 _dafny.Array = _63_v16
					_ = _lhs0
					var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_63_v16), 0))
					_ = _lhs1
					(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
					_71_v18 = _rhs1
					r0 = _rhs2
				} else if _source2.Is_DC7() {
					var _72___mcc_h3 _dafny.Int = _source2.Get_().(D2_DC7).Cf15
					_ = _72___mcc_h3
					var _73_cf15 _dafny.Int = _72___mcc_h3
					_ = _73_cf15
					var _74_v19 _dafny.Set
					_ = _74_v19
					_74_v19 = _dafny.SetOf(_dafny.IntOfInt64(602))
					r0 = (func() bool {
						if p0 {
							return !(p0)
						}
						return (_74_v19).IsProperSubsetOf(_74_v19)
					})()
					r0 = Companion_Default___.Fm0(true, p0, Companion_Default___.SafeDivisionInt(_73_cf15, _73_cf15), globalState)
					var _75_v20 _dafny.Set
					_ = _75_v20
					_75_v20 = _dafny.SetOf(_74_v19)
					var _76_v21 _dafny.Map
					_ = _76_v21
					_76_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
						if p0 {
							return (_75_v20).Cardinality()
						}
						return (_dafny.Zero).Minus(_73_cf15)
					})(), _41_v1)
					_76_v21 = _76_v21
					var _77_v22 _dafny.Set
					_ = _77_v22
					_77_v22 = _dafny.SetOf(false, p0, p0)
					var _78_v23 *C0
					_ = _78_v23
					var _nw5 *C0 = New_C0_()
					_ = _nw5
					_nw5.Ctor__((_dafny.Zero).Minus(_73_cf15), _77_v22)
					_78_v23 = _nw5
				} else if _source2.Is_DC5() {
					var _79___mcc_h4 _dafny.Array = _source2.Get_().(D2_DC5).Cf11
					_ = _79___mcc_h4
					var _80_cf11 _dafny.Array = _79___mcc_h4
					_ = _80_cf11
					var _81_v24 _dafny.CodePoint
					_ = _81_v24
					_81_v24 = _dafny.CodePoint('w')
					var _82_v25 _dafny.Map
					_ = _82_v25
					_82_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _81_v24)
					_82_v25 = (_82_v25).Update(p0, (func() _dafny.CodePoint {
						if false {
							return _81_v24
						}
						return _81_v24
					})())
					var _83_v26 *C0
					_ = _83_v26
					var _nw6 *C0 = New_C0_()
					_ = _nw6
					_nw6.Ctor__(Companion_Default___.SafeModuloInt(_41_v1, _41_v1), _dafny.SetOf(p0, p0))
					_83_v26 = _nw6
					_83_v26 = _83_v26
					var _84_v27 _dafny.Set
					_ = _84_v27
					_84_v27 = _dafny.SetOf((_83_v26).F6(), _41_v1, _dafny.IntOfInt64(-816), _41_v1)
					var _85_v28 *C0
					_ = _85_v28
					var _nw7 *C0 = New_C0_()
					_ = _nw7
					_nw7.Ctor__((_84_v27).Cardinality(), (_83_v26).F7())
					_85_v28 = _nw7
					(globalState).F0 = (_41_v1).Minus((_83_v26).F6())
				} else {
					var _86___mcc_h5 D2 = _source2.Get_().(D2_DC8).Cf16
					_ = _86___mcc_h5
					var _87_cf16 D2 = _86___mcc_h5
					_ = _87_cf16
					_41_v1 = (func() _dafny.Map {
						var _coll12 = _dafny.NewMapBuilder()
						_ = _coll12
						for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(933), _dafny.IntOfInt64(887))); ; {
							_compr_12, _ok12 := _iter12()
							if !_ok12 {
								break
							}
							var _88_v29 _dafny.Int
							_88_v29 = interface{}(_compr_12).(_dafny.Int)
							if ((_dafny.IntOfInt64(933)).Cmp(_88_v29) <= 0) && ((_88_v29).Cmp(_dafny.IntOfInt64(887)) < 0) {
								_coll12.Add(Companion_Default___.SafeDivisionInt(_88_v29, _dafny.IntOfUint32((_dafny.SeqOf(!(p0))).Cardinality())), p0)
							}
						}
						return _coll12.ToMap()
					}()).Cardinality()
					var _89_v31 _dafny.Map
					_ = _89_v31
					_89_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_41_v1, _41_v1)
					var _90_v32 _dafny.Map
					_ = _90_v32
					_90_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v31, _41_v1)
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(142), _dafny.ArrayLen((_43_v3), 0))
					_ = _index2
					(_43_v3).ArraySet1((func() _dafny.Int {
						if p0 {
							return _41_v1
						}
						return ((func() _dafny.Map {
							var _coll13 = _dafny.NewMapBuilder()
							_ = _coll13
							for _iter13 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.MultiSetOf(p0, p0)).Cardinality())).Elements()); ; {
								_compr_13, _ok13 := _iter13()
								if !_ok13 {
									break
								}
								var _91_v30 _dafny.Int
								_91_v30 = interface{}(_compr_13).(_dafny.Int)
								if (_dafny.MultiSetOf((_dafny.MultiSetOf(p0, p0)).Cardinality())).Contains(_91_v30) {
									_coll13.Add((_91_v30).Times(_dafny.IntOfInt64(877)), p0)
								}
							}
							return _coll13.ToMap()
						}()).Update((_90_v32).Cardinality(), p0)).Cardinality()
					})(), (_index2).Int())
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(142), _dafny.ArrayLen((_43_v3), 0))
					_ = _index3
					(_43_v3).ArraySet1(_41_v1, (_index3).Int())
					var _92_v33 _dafny.Array
					_ = _92_v33
					var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
					_ = _nw8
					_92_v33 = _nw8
					_92_v33 = _92_v33
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(142), _dafny.ArrayLen((_43_v3), 0))
					_ = _index4
					(_43_v3).ArraySet1(Companion_Default___.SafeModuloInt((_43_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(142), _dafny.ArrayLen((_43_v3), 0))).Int()).(_dafny.Int), (_43_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(142), _dafny.ArrayLen((_43_v3), 0))).Int()).(_dafny.Int)), (_index4).Int())
				}
				var _93_v34 _dafny.Array
				_ = _93_v34
				var _nw9 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
				_ = _nw9
				_93_v34 = _nw9
				var _94_v35 _dafny.MultiSet
				_ = _94_v35
				_94_v35 = _dafny.MultiSetOf((_41_v1).Cmp(_41_v1) >= 0)
				var _95_v36 _dafny.Set
				_ = _95_v36
				_95_v36 = _dafny.SetOf(_41_v1)
				var _96_v37 D3
				_ = _96_v37
				_96_v37 = Companion_D3_.Create_DC11_()
				var _rhs3 _dafny.Array = _93_v34
				_ = _rhs3
				var _rhs4 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_41_v1, _41_v1), ((func() _dafny.Set {
					if p0 {
						return _95_v36
					}
					return _95_v36
				})()).Cardinality())
				_ = _rhs4
				var _rhs5 _dafny.MultiSet = Companion_Default___.Fm15(_96_v37, p0, globalState)
				_ = _rhs5
				var _lhs2 *GlobalState = globalState
				_ = _lhs2
				_93_v34 = _rhs3
				_lhs2.F0 = _rhs4
				_94_v35 = _rhs5
				var _97_v38 _dafny.Set
				_ = _97_v38
				_97_v38 = _dafny.SetOf(true)
				var _98_v39 *C0
				_ = _98_v39
				var _nw10 *C0 = New_C0_()
				_ = _nw10
				_nw10.Ctor__(_41_v1, _97_v38)
				_98_v39 = _nw10
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _109_v40 _dafny.Int
	_ = _109_v40
	_109_v40 = _dafny.IntOfInt64(-284)
	var _source4 D4 = Companion_D4_.Create_DC15_((_109_v40).Cmp(_dafny.IntOfInt64(305)) >= 0, _109_v40)
	_ = _source4
	if _source4.Is_DC14() {
		r0 = p0
		r0 = p0
		r0 = !(p0)
		var _110_v41 _dafny.Array
		_ = _110_v41
		var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
		_ = _nw11
		_110_v41 = _nw11
		var _111_v42 _dafny.Array
		_ = _111_v42
		var _nwElement0_2 _dafny.Array = _110_v41
		_ = _nwElement0_2
		var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(17))
		_ = _nw12
		(_nw12).ArraySet1(_nwElement0_2, 0)
		(_nw12).ArraySet1(_110_v41, 1)
		(_nw12).ArraySet1(_110_v41, 2)
		(_nw12).ArraySet1(_110_v41, 3)
		(_nw12).ArraySet1(_110_v41, 4)
		(_nw12).ArraySet1(_110_v41, 5)
		(_nw12).ArraySet1(_110_v41, 6)
		(_nw12).ArraySet1(_110_v41, 7)
		(_nw12).ArraySet1(_110_v41, 8)
		(_nw12).ArraySet1(_110_v41, 9)
		(_nw12).ArraySet1(_110_v41, 10)
		(_nw12).ArraySet1(_110_v41, 11)
		(_nw12).ArraySet1(_110_v41, 12)
		(_nw12).ArraySet1(_110_v41, 13)
		(_nw12).ArraySet1(_110_v41, 14)
		(_nw12).ArraySet1(_110_v41, 15)
		(_nw12).ArraySet1(_110_v41, 16)
		_111_v42 = _nw12
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_111_v42), 0))
		_ = _index5
		(_111_v42).ArraySet1(_110_v41, (_index5).Int())
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_111_v42), 0))
		_ = _index6
		(_111_v42).ArraySet1(_110_v41, (_index6).Int())
	} else if _source4.Is_DC15() {
		var _112___mcc_h10 bool = _source4.Get_().(D4_DC15).Cf23
		_ = _112___mcc_h10
		var _113___mcc_h11 _dafny.Int = _source4.Get_().(D4_DC15).Cf24
		_ = _113___mcc_h11
		var _114_cf24 _dafny.Int = _113___mcc_h11
		_ = _114_cf24
		var _115_cf23 bool = _112___mcc_h10
		_ = _115_cf23
		var _116_v43 _dafny.Sequence
		_ = _116_v43
		_116_v43 = _dafny.UnicodeSeqOfUtf8Bytes("sficpirr")
		var _117_v44 _dafny.CodePoint
		_ = _117_v44
		_117_v44 = _dafny.CodePoint('d')
		var _118_v45 D1
		_ = _118_v45
		_118_v45 = Companion_D1_.Create_DC3_(_115_cf23, Companion_Default___.Fm0(p0, _115_cf23, _114_cf24, globalState), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_116_v43, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(141))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_119_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		}))), (Companion_Default___.SafeIndex(_109_v40, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_116_v43, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(141))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_120_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		})))).Cardinality()))).Uint32(), _117_v44), (_dafny.IntOfUint32((_116_v43).Cardinality())).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _109_v40)).Cardinality()) <= 0, (_dafny.Zero).Minus(_109_v40))
		var _rhs6 D1 = _118_v45
		_ = _rhs6
		var _rhs7 _dafny.Int = Companion_Default___.Fm4(_109_v40, p0, globalState)
		_ = _rhs7
		_118_v45 = _rhs6
		_114_cf24 = _rhs7
		_116_v43 = _dafny.UnicodeSeqOfUtf8Bytes("m")
		if false {
			(globalState).F4 = _dafny.IntOfInt64(206)
			var _121_v46 _dafny.Set
			_ = _121_v46
			_121_v46 = _dafny.SetOf(false, _115_cf23)
			var _122_v47 *C0
			_ = _122_v47
			var _nw13 *C0 = New_C0_()
			_ = _nw13
			_nw13.Ctor__(_114_cf24, _121_v46)
			_122_v47 = _nw13
			var _123_v48 _dafny.Map
			_ = _123_v48
			_123_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v47, _115_cf23)
			var _124_v49 _dafny.Map
			_ = _124_v49
			_124_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(838), _122_v47)
			_115_cf23 = Companion_Default___.Fm0(p0, !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_123_v48).Update((func() *C0 {
				if (_124_v49).Contains(_114_cf24) {
					return (_124_v49).Get(_114_cf24).(*C0)
				}
				return _122_v47
			})(), p0)).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_115_cf23, (_116_v43).Select((Companion_Default___.SafeIndex((_122_v47).F6(), _dafny.IntOfUint32((_116_v43).Cardinality()))).Uint32()).(_dafny.CodePoint)))).Contains(_dafny.IntOfUint32((_116_v43).Cardinality())), (_dafny.Zero).Minus((_122_v47).F6()), globalState)
			var _125_v50 *C0
			_ = _125_v50
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__((_122_v47).F6(), _121_v46)
			_125_v50 = _nw14
			var _126_v51 _dafny.Set
			_ = _126_v51
			_126_v51 = _dafny.SetOf(_124_v49)
			_115_cf23 = (_126_v51).IsSubsetOf(_126_v51)
			var _127_v52 _dafny.Map
			_ = _127_v52
			_127_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v47).F6(), p0)
			_127_v52 = (_127_v52).Update((_dafny.Zero).Minus((_125_v50).F6()), p0)
		} else {
			r0 = !(p0)
			var _128_v53 _dafny.Set
			_ = _128_v53
			_128_v53 = _dafny.SetOf(_115_cf23)
			var _129_v54 *C0
			_ = _129_v54
			var _nw15 *C0 = New_C0_()
			_ = _nw15
			_nw15.Ctor__(_109_v40, _128_v53)
			_129_v54 = _nw15
			var _130_v55 D4
			_ = _130_v55
			_130_v55 = Companion_D4_.Create_DC13_(_129_v54)
			var _131_v56 _dafny.Map
			_ = _131_v56
			_131_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_130_v55, _115_cf23)
			var _132_v57 *C0
			_ = _132_v57
			var _nw16 *C0 = New_C0_()
			_ = _nw16
			_nw16.Ctor__((_131_v56).Cardinality(), _128_v53)
			_132_v57 = _nw16
			var _133_v58 _dafny.Sequence
			_ = _133_v58
			_133_v58 = _dafny.SeqOf(p0)
			var _134_v59 _dafny.Map
			_ = _134_v59
			_134_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsPrefixOf(_116_v43, _116_v43), !_dafny.Companion_Sequence_.Contains(_133_v58, p0))
			_134_v59 = _134_v59
			r0 = _115_cf23
			var _135_v60 *C0
			_ = _135_v60
			var _nw17 *C0 = New_C0_()
			_ = _nw17
			_nw17.Ctor__((_dafny.Zero).Minus((_132_v57).F6()), (_128_v53).Union(_dafny.SetOf(_115_cf23)))
			_135_v60 = _nw17
		}
		var _136_v61 _dafny.Set
		_ = _136_v61
		_136_v61 = _dafny.SetOf(_109_v40)
		var _137_v62 _dafny.Map
		_ = _137_v62
		_137_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v44, _136_v61)
		var _138_v63 _dafny.Array
		_ = _138_v63
		var _nwElement0_3 _dafny.Int = (_dafny.IntOfUint32((_116_v43).Cardinality())).Minus(_114_cf24)
		_ = _nwElement0_3
		var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(10))
		_ = _nw18
		(_nw18).ArraySet1(_nwElement0_3, 0)
		(_nw18).ArraySet1(_114_cf24, 1)
		(_nw18).ArraySet1(_109_v40, 2)
		(_nw18).ArraySet1(_109_v40, 3)
		(_nw18).ArraySet1(_dafny.IntOfInt64(446), 4)
		(_nw18).ArraySet1(Companion_Default___.Fm4(_114_cf24, p0, globalState), 5)
		(_nw18).ArraySet1(Companion_Default___.Fm4(_109_v40, _115_cf23, globalState), 6)
		(_nw18).ArraySet1(_109_v40, 7)
		(_nw18).ArraySet1(_109_v40, 8)
		(_nw18).ArraySet1((_dafny.Zero).Minus((_109_v40).Minus((_137_v62).Cardinality())), 9)
		_138_v63 = _nw18
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_138_v63), 0))
		_ = _index7
		(_138_v63).ArraySet1(_109_v40, (_index7).Int())
		var _139_v64 _dafny.Array
		_ = _139_v64
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_0
		var _nw19 _dafny.Array
		_ = _nw19
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw19 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Sequence = (func(_140_v43 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_141_i6 _dafny.Int) _dafny.Sequence {
					return _140_v43
				}
			})(_116_v43)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw19 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw19).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw19).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_139_v64 = _nw19
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_139_v64), 0))
		_ = _index8
		(_139_v64).ArraySet1(_116_v43, (_index8).Int())
		var _142_v65 _dafny.Array
		_ = _142_v65
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_1
		var _nw20 _dafny.Array
		_ = _nw20
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw20 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_143_v43 _dafny.Sequence) func(_dafny.Int) bool {
				return func(_144_i7 _dafny.Int) bool {
					return !(!(_dafny.Companion_Sequence_.IsPrefixOf(_143_v43, _143_v43)))
				}
			})(_116_v43)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw20 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw20).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw20).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_142_v65 = _nw20
		var _145_v66 _dafny.Sequence
		_ = _145_v66
		_145_v66 = _dafny.SeqOf(p0)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_142_v65), 0))
		_ = _index9
		(_142_v65).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(p0, _115_cf23, _115_cf23), _145_v66), (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_138_v63), 0))
		_ = _index10
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_139_v64), 0))
		_ = _index11
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_142_v65), 0))
		_ = _index12
		var _rhs8 _dafny.Int = _109_v40
		_ = _rhs8
		var _rhs9 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("ihfjckhnn")
		_ = _rhs9
		var _rhs10 _dafny.Int = _109_v40
		_ = _rhs10
		var _rhs11 _dafny.Int = ((_114_cf24).Times(_dafny.IntOfUint32((_116_v43).Cardinality()))).Minus(_dafny.IntOfInt64(205))
		_ = _rhs11
		var _rhs12 bool = p0
		_ = _rhs12
		var _lhs3 _dafny.Array = _138_v63
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_138_v63), 0))
		_ = _lhs4
		var _lhs5 _dafny.Array = _139_v64
		_ = _lhs5
		var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_139_v64), 0))
		_ = _lhs6
		var _lhs7 *GlobalState = globalState
		_ = _lhs7
		var _lhs8 *GlobalState = globalState
		_ = _lhs8
		var _lhs9 _dafny.Array = _142_v65
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_142_v65), 0))
		_ = _lhs10
		(_lhs3).ArraySet1(_rhs8, (_lhs4).Int())
		(_lhs5).ArraySet1(_rhs9, (_lhs6).Int())
		_lhs7.F0 = _rhs10
		_lhs8.F0 = _rhs11
		(_lhs9).ArraySet1(_rhs12, (_lhs10).Int())
	} else if _source4.Is_DC16() {
		var _146___mcc_h12 _dafny.Set = _source4.Get_().(D4_DC16).Cf25
		_ = _146___mcc_h12
		var _147___mcc_h13 _dafny.CodePoint = _source4.Get_().(D4_DC16).Cf26
		_ = _147___mcc_h13
		var _148___mcc_h14 _dafny.Int = _source4.Get_().(D4_DC16).Cf27
		_ = _148___mcc_h14
		var _149___mcc_h15 _dafny.CodePoint = _source4.Get_().(D4_DC16).Cf28
		_ = _149___mcc_h15
		var _150_cf28 _dafny.CodePoint = _149___mcc_h15
		_ = _150_cf28
		var _151_cf27 _dafny.Int = _148___mcc_h14
		_ = _151_cf27
		var _152_cf26 _dafny.CodePoint = _147___mcc_h13
		_ = _152_cf26
		var _153_cf25 _dafny.Set = _146___mcc_h12
		_ = _153_cf25
		var _154_v67 _dafny.Sequence
		_ = _154_v67
		_154_v67 = _dafny.UnicodeSeqOfUtf8Bytes("adyhsx")
		var _155_v68 _dafny.Map
		_ = _155_v68
		_155_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_154_v67, p0)
		var _156_v69 _dafny.MultiSet
		_ = _156_v69
		_156_v69 = _dafny.MultiSetOf(_109_v40, (_155_v68).Cardinality())
		var _157_v70 _dafny.Sequence
		_ = _157_v70
		_157_v70 = _dafny.SeqOf(_dafny.IntOfUint32((_154_v67).Cardinality()), _151_cf27)
		var _158_v71 D2
		_ = _158_v71
		_158_v71 = Companion_D2_.Create_DC7_((_157_v70).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.IntOfUint32((_157_v70).Cardinality()))).Uint32()).(_dafny.Int))
		var _159_v72 _dafny.Map
		_ = _159_v72
		_159_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-151), _158_v71)
		var _160_v73 _dafny.MultiSet
		_ = _160_v73
		_160_v73 = _dafny.MultiSetOf((_151_cf27).Cmp((func() _dafny.Int {
			if (_156_v69).Contains(_109_v40) {
				return (_156_v69).Multiplicity(_109_v40)
			}
			return (_159_v72).Cardinality()
		})()) >= 0)
		var _161_v74 _dafny.Array
		_ = _161_v74
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_2
		var _nw21 _dafny.Array
		_ = _nw21
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw21 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_162_v70 _dafny.Sequence, _163_cf26 _dafny.CodePoint) func(_dafny.Int) _dafny.Int {
				return func(_164_i8 _dafny.Int) _dafny.Int {
					return (_164_i8).Plus((_162_v70).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(3))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg11 _dafny.Int) interface{} {
							return coer11(arg11)
						}
					}((func(_165_cf26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_166_i9 _dafny.Int) _dafny.CodePoint {
							return _165_cf26
						}
					})(_163_cf26)))).Cardinality()), _dafny.IntOfUint32((_162_v70).Cardinality()))).Uint32()).(_dafny.Int))
				}
			})(_157_v70, _152_cf26)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw21 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw21).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw21).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_161_v74 = _nw21
		var _167_v75 _dafny.Map
		_ = _167_v75
		_167_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v74, p0)
		var _168_v76 _dafny.Map
		_ = _168_v76
		_168_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(365), _dafny.UnicodeSeqOfUtf8Bytes("csu"))
		(globalState).F0 = (func() _dafny.Int {
			if (_160_v73).Contains(Companion_Default___.Fm0(p0, (func() bool {
				if (_167_v75).Contains(_161_v74) {
					return (_167_v75).Get(_161_v74).(bool)
				}
				return p0
			})(), (_168_v76).Cardinality(), globalState)) {
				return (_160_v73).Multiplicity(Companion_Default___.Fm0(p0, (func() bool {
					if (_167_v75).Contains(_161_v74) {
						return (_167_v75).Get(_161_v74).(bool)
					}
					return p0
				})(), (_168_v76).Cardinality(), globalState))
			}
			return _109_v40
		})()
		r0 = p0
		var _169_v77 _dafny.Array
		_ = _169_v77
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
		_ = _nw22
		_169_v77 = _nw22
		_169_v77 = _169_v77
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_161_v74), 0))
		_ = _index13
		(_161_v74).ArraySet1((_151_cf27).Plus(_109_v40), (_index13).Int())
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_161_v74), 0))
		_ = _index14
		(_161_v74).ArraySet1(_151_cf27, (_index14).Int())
	} else {
		var _170___mcc_h16 *C0 = _source4.Get_().(D4_DC13).Cf22
		_ = _170___mcc_h16
		var _171_cf22 *C0 = _170___mcc_h16
		_ = _171_cf22
		var _172_v78 D3
		_ = _172_v78
		_172_v78 = Companion_D3_.Create_DC10_(p0, _109_v40, _109_v40)
		var _source5 D3 = _172_v78
		_ = _source5
		if _source5.Is_DC10() {
			var _173___mcc_h17 bool = _source5.Get_().(D3_DC10).Cf18
			_ = _173___mcc_h17
			var _174___mcc_h18 _dafny.Int = _source5.Get_().(D3_DC10).Cf19
			_ = _174___mcc_h18
			var _175___mcc_h19 _dafny.Int = _source5.Get_().(D3_DC10).Cf20
			_ = _175___mcc_h19
			var _176_cf20 _dafny.Int = _175___mcc_h19
			_ = _176_cf20
			var _177_cf19 _dafny.Int = _174___mcc_h18
			_ = _177_cf19
			var _178_cf18 bool = _173___mcc_h17
			_ = _178_cf18
			var _179_v79 _dafny.Sequence
			_ = _179_v79
			_179_v79 = _dafny.SeqOf(_109_v40)
			_178_cf18 = !_dafny.Companion_Sequence_.Equal(_179_v79, _179_v79)
			_178_cf18 = _178_cf18
			var _180_v80 _dafny.Array
			_ = _180_v80
			var _nwElement0_4 _dafny.Int = _dafny.IntOfInt64(-215)
			_ = _nwElement0_4
			var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(9))
			_ = _nw23
			(_nw23).ArraySet1(_nwElement0_4, 0)
			(_nw23).ArraySet1(_dafny.IntOfUint32((_179_v79).Cardinality()), 1)
			(_nw23).ArraySet1((_171_cf22).F6(), 2)
			(_nw23).ArraySet1((_171_cf22).F6(), 3)
			(_nw23).ArraySet1((_171_cf22).F6(), 4)
			(_nw23).ArraySet1(_177_cf19, 5)
			(_nw23).ArraySet1(_176_cf20, 6)
			(_nw23).ArraySet1(_109_v40, 7)
			(_nw23).ArraySet1(_109_v40, 8)
			_180_v80 = _nw23
			var _181_v81 D2
			_ = _181_v81
			_181_v81 = Companion_D2_.Create_DC6_(_176_cf20, p0, _180_v80)
			var _182_v82 _dafny.MultiSet
			_ = _182_v82
			_182_v82 = _dafny.MultiSetOf((_181_v81).Dtor_cf12(), (_171_cf22).F6())
			var _183_v83 _dafny.Map
			_ = _183_v83
			_183_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_171_cf22).F6(), _182_v82)
			(globalState).F4 = (_171_cf22).Fm1(((func() _dafny.MultiSet {
				if (_183_v83).Contains(_109_v40) {
					return (_183_v83).Get(_109_v40).(_dafny.MultiSet)
				}
				return _182_v82
			})()).Update((_dafny.Zero).Minus((_171_cf22).F6()), Companion_Default___.Abs((_171_cf22).F6())), _39_v0, globalState)
			_180_v80 = _180_v80
		} else if _source5.Is_DC11() {
			var _184_v84 _dafny.Map
			_ = _184_v84
			_184_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
			var _185_v85 _dafny.Sequence
			_ = _185_v85
			_185_v85 = _dafny.SeqOf(_109_v40, (_171_cf22).F6())
			var _186_v86 _dafny.Map
			_ = _186_v86
			_186_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_171_cf22).F6(), _185_v85)
			var _187_v87 _dafny.Map
			_ = _187_v87
			_187_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_184_v84).Cardinality(), _dafny.Companion_Sequence_.Concatenate(_185_v85, (func() _dafny.Sequence {
				if (_186_v86).Contains(_dafny.IntOfInt64(603)) {
					return (_186_v86).Get(_dafny.IntOfInt64(603)).(_dafny.Sequence)
				}
				return _185_v85
			})()))
			var _188_v88 _dafny.CodePoint
			_ = _188_v88
			_188_v88 = _dafny.CodePoint('j')
			var _189_v89 _dafny.MultiSet
			_ = _189_v89
			_189_v89 = _dafny.MultiSetOf(p0, p0, !(p0))
			var _190_v90 _dafny.Sequence
			_ = _190_v90
			_190_v90 = _dafny.UnicodeSeqOfUtf8Bytes("pcyod")
			var _191_v91 _dafny.Sequence
			_ = _191_v91
			_191_v91 = _dafny.SeqOf((Companion_Default___.Fm7(_189_v89, _190_v90, globalState)).Update(p0, (_171_cf22).F6()))
			_186_v86 = (_186_v86).Update((_171_cf22).Fm2((Companion_Default___.Fm16(_109_v40, _188_v88, globalState)).Dtor_cf23(), _185_v85, _185_v85, (_191_v91).Select((Companion_Default___.SafeIndex(_109_v40, _dafny.IntOfUint32((_191_v91).Cardinality()))).Uint32()).(_dafny.Map), globalState), _185_v85)
			r0 = !(!(p0))
			var _192_v92 _dafny.Map
			_ = _192_v92
			_192_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v40, Companion_Default___.Fm0(p0, p0, (_171_cf22).F6(), globalState))
			var _193_v93 _dafny.Set
			_ = _193_v93
			_193_v93 = _dafny.SetOf(_109_v40, _109_v40, (_171_cf22).F6(), (_171_cf22).F6(), _109_v40)
			_192_v92 = (_192_v92).Update((_171_cf22).F6(), (func() _dafny.Set {
				var _coll14 = _dafny.NewBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(312), _dafny.IntOfInt64(235))); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _194_v94 _dafny.Int
					_194_v94 = interface{}(_compr_14).(_dafny.Int)
					if ((_dafny.IntOfInt64(312)).Cmp(_194_v94) <= 0) && ((_194_v94).Cmp(_dafny.IntOfInt64(235)) < 0) {
						_coll14.Add(Companion_Default___.SafeModuloInt(_194_v94, (_171_cf22).F6()))
					}
				}
				return _coll14.ToSet()
			}()).IsProperSubsetOf(_193_v93))
			(globalState).F4 = (_dafny.Zero).Minus((_171_cf22).F6())
		} else if _source5.Is_DC9() {
			var _195___mcc_h20 _dafny.Sequence = _source5.Get_().(D3_DC9).Cf17
			_ = _195___mcc_h20
			var _196_cf17 _dafny.Sequence = _195___mcc_h20
			_ = _196_cf17
			var _197_v95 _dafny.Array
			_ = _197_v95
			var _nw24 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw24
			_197_v95 = _nw24
			_197_v95 = _197_v95
			var _198_v96 _dafny.Array
			_ = _198_v96
			var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw25
			_198_v96 = _nw25
			var _rhs13 bool = p0
			_ = _rhs13
			var _rhs14 _dafny.Array = _198_v96
			_ = _rhs14
			var _rhs15 _dafny.Int = _109_v40
			_ = _rhs15
			var _rhs16 _dafny.Int = (_171_cf22).F6()
			_ = _rhs16
			var _lhs11 *GlobalState = globalState
			_ = _lhs11
			var _lhs12 *GlobalState = globalState
			_ = _lhs12
			r0 = _rhs13
			_198_v96 = _rhs14
			_lhs11.F0 = _rhs15
			_lhs12.F4 = _rhs16
			var _199_v97 *C0
			_ = _199_v97
			var _nw26 *C0 = New_C0_()
			_ = _nw26
			_nw26.Ctor__((_171_cf22).F6(), (_171_cf22).F7())
			_199_v97 = _nw26
			var _200_v98 _dafny.CodePoint
			_ = _200_v98
			_200_v98 = _dafny.CodePoint('c')
			_200_v98 = _200_v98
		} else {
			var _201___mcc_h21 D3 = _source5.Get_().(D3_DC12).Cf21
			_ = _201___mcc_h21
			var _202_cf21 D3 = _201___mcc_h21
			_ = _202_cf21
			var _nw27 *C0 = New_C0_()
			_ = _nw27
			_nw27.Ctor__((func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(320), _dafny.IntOfInt64(-184))); ; {
					_compr_15, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _203_v99 _dafny.Int
					_203_v99 = interface{}(_compr_15).(_dafny.Int)
					if ((_dafny.IntOfInt64(320)).Cmp(_203_v99) <= 0) && ((_203_v99).Cmp(_dafny.IntOfInt64(-184)) < 0) {
						_coll15.Add((_203_v99).Plus((_171_cf22).F6()), p0)
					}
				}
				return _coll15.ToMap()
			}()).Cardinality(), (_171_cf22).F7())
			_171_cf22 = _nw27
			var _204_v100 _dafny.Array
			_ = _204_v100
			var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw28
			_204_v100 = _nw28
			_204_v100 = _204_v100
			var _205_v101 _dafny.Sequence
			_ = _205_v101
			_205_v101 = _dafny.SeqOf(p0)
			var _206_v102 _dafny.Array
			_ = _206_v102
			var _nwElement0_5 bool = p0
			_ = _nwElement0_5
			var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(15))
			_ = _nw29
			(_nw29).ArraySet1(_nwElement0_5, 0)
			(_nw29).ArraySet1(p0, 1)
			(_nw29).ArraySet1(!(!(p0)) || (p0), 2)
			(_nw29).ArraySet1(p0, 3)
			(_nw29).ArraySet1(p0, 4)
			(_nw29).ArraySet1((_205_v101).Select((Companion_Default___.SafeIndex((_171_cf22).F6(), _dafny.IntOfUint32((_205_v101).Cardinality()))).Uint32()).(bool), 5)
			(_nw29).ArraySet1(p0, 6)
			(_nw29).ArraySet1((_172_v78).Dtor_cf18(), 7)
			(_nw29).ArraySet1((p0) == (p0), 8)
			(_nw29).ArraySet1(false, 9)
			(_nw29).ArraySet1(p0, 10)
			(_nw29).ArraySet1(p0, 11)
			(_nw29).ArraySet1(p0, 12)
			(_nw29).ArraySet1(p0, 13)
			(_nw29).ArraySet1(Companion_Default___.Fm0(!(p0), false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_205_v101).Cardinality())), globalState), 14)
			_206_v102 = _nw29
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_206_v102), 0))
			_ = _index15
			(_206_v102).ArraySet1(p0, (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_206_v102), 0))
			_ = _index16
			(_206_v102).ArraySet1(!(p0), (_index16).Int())
			var _207_v103 _dafny.Sequence
			_ = _207_v103
			_207_v103 = _dafny.UnicodeSeqOfUtf8Bytes("pibona")
			var _208_v104 _dafny.Sequence
			_ = _208_v104
			_208_v104 = _dafny.SeqOf(_109_v40, _dafny.IntOfUint32((_207_v103).Cardinality()))
			var _209_v105 _dafny.Map
			_ = _209_v105
			_209_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(p0, (_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), _dafny.IntOfUint32((_208_v104).Cardinality()), globalState), _171_cf22)
			var _210_v106 _dafny.Map
			_ = _210_v106
			_210_v106 = _209_v105
			var _211_v107 _dafny.Array
			_ = _211_v107
			var _nwElement0_6 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(p0, (_206_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_206_v102), 0))).Int()).(bool), (_171_cf22).F6(), globalState), _171_cf22)
			_ = _nwElement0_6
			var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(3))
			_ = _nw30
			(_nw30).ArraySet1(_nwElement0_6, 0)
			(_nw30).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _171_cf22), 1)
			(_nw30).ArraySet1((_210_v106).Merge(_209_v105), 2)
			_211_v107 = _nw30
			_211_v107 = _211_v107
		}
		if !(p0) {
			var _212_v108 _dafny.CodePoint
			_ = _212_v108
			_212_v108 = _dafny.CodePoint('w')
			var _213_v109 _dafny.Map
			_ = _213_v109
			_213_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v108, _171_cf22)
			_213_v109 = (_213_v109).Update((func() _dafny.CodePoint {
				if !(p0) {
					return _212_v108
				}
				return _212_v108
			})(), _171_cf22)
			var _214_v110 *C0
			_ = _214_v110
			var _nw31 *C0 = New_C0_()
			_ = _nw31
			_nw31.Ctor__(((_171_cf22).F6()).Times((_171_cf22).F6()), ((_171_cf22).F7()).Intersection(_dafny.SetOf(p0)))
			_214_v110 = _nw31
			(globalState).F4 = (_171_cf22).F6()
			var _215_v111 _dafny.MultiSet
			_ = _215_v111
			_215_v111 = _dafny.MultiSetOf((_171_cf22).F6(), (_171_cf22).F6())
			var _216_v112 _dafny.Sequence
			_ = _216_v112
			_216_v112 = _dafny.SeqOf((_215_v111).Cardinality(), _dafny.IntOfInt64(-745))
			var _217_v114 _dafny.Sequence
			_ = _217_v114
			_217_v114 = _dafny.SeqOf((func() _dafny.Set {
				var _coll16 = _dafny.NewBuilder()
				_ = _coll16
				for _iter16 := _dafny.Iterate((_216_v112).Elements()); ; {
					_compr_16, _ok16 := _iter16()
					if !_ok16 {
						break
					}
					var _218_v113 _dafny.Int
					_218_v113 = interface{}(_compr_16).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_216_v112, _218_v113) {
						_coll16.Add(Companion_Default___.SafeDivisionInt(_218_v113, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yr")).Cardinality())))
					}
				}
				return _coll16.ToSet()
			}()).Cardinality())
			_217_v114 = _217_v114
			r0 = (((_171_cf22).F6()).Times((_171_cf22).F6())).Cmp(_109_v40) < 0
		} else {
			var _219_v115 _dafny.Set
			_ = _219_v115
			_219_v115 = _dafny.SetOf(_dafny.IntOfInt64(682))
			_219_v115 = _dafny.SetOf(_109_v40, _109_v40, _109_v40)
			var _220_v116 _dafny.Map
			_ = _220_v116
			_220_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_cf22, false)
			var _221_v117 _dafny.Sequence
			_ = _221_v117
			_221_v117 = _dafny.SeqOf(p0, false, !(_220_v116).Equals(_220_v116))
			_221_v117 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if false {
					return _221_v117
				}
				return _221_v117
			})(), _221_v117)
			var _222_v118 _dafny.Array
			_ = _222_v118
			var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
			_ = _nw32
			_222_v118 = _nw32
			_222_v118 = _222_v118
			var _223_v119 _dafny.Array
			_ = _223_v119
			var _nw33 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw33
			_223_v119 = _nw33
			var _224_v120 _dafny.CodePoint
			_ = _224_v120
			_224_v120 = _dafny.CodePoint('y')
			var _225_v122 _dafny.Map
			_ = _225_v122
			_225_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v120, _dafny.IntOfInt64(-9))
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_223_v119), 0))
			_ = _index17
			(_223_v119).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v120, p0)).Equals(func() _dafny.Map {
				var _coll17 = _dafny.NewMapBuilder()
				_ = _coll17
				for _iter17 := _dafny.Iterate((_225_v122).Keys().Elements()); ; {
					_compr_17, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _226_v121 _dafny.CodePoint
					_226_v121 = interface{}(_compr_17).(_dafny.CodePoint)
					if (_225_v122).Contains(_226_v121) {
						_coll17.Add(_226_v121, p0)
					}
				}
				return _coll17.ToMap()
			}()), (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_223_v119), 0))
			_ = _index18
			(_223_v119).ArraySet1(p0, (_index18).Int())
			var _227_v123 _dafny.Array
			_ = _227_v123
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_3
			var _nw34 _dafny.Array
			_ = _nw34
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw34 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Int = (func(_228_cf22 *C0) func(_dafny.Int) _dafny.Int {
					return func(_229_i10 _dafny.Int) _dafny.Int {
						return (_229_i10).Plus((_228_cf22).F6())
					}
				})(_171_cf22)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw34 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw34).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw34).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_227_v123 = _nw34
			var _230_v124 _dafny.Map
			_ = _230_v124
			_230_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v119, _227_v123)
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_223_v119), 0))
			_ = _index19
			(_223_v119).ArraySet1(!(_230_v124).Contains(_223_v119), (_index19).Int())
		}
		var _231_v125 _dafny.Set
		_ = _231_v125
		_231_v125 = _dafny.SetOf((_171_cf22).F6())
		var _232_v126 _dafny.Map
		_ = _232_v126
		_232_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_231_v125, _109_v40)
		var _233_v128 _dafny.Map
		_ = _233_v128
		_233_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v40, func() _dafny.Set {
			var _coll18 = _dafny.NewBuilder()
			_ = _coll18
			for _iter18 := _dafny.Iterate((_232_v126).Keys().Elements()); ; {
				_compr_18, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _234_v127 _dafny.Set
				_234_v127 = interface{}(_compr_18).(_dafny.Set)
				if (_232_v126).Contains(_234_v127) {
					_coll18.Add(_234_v127)
				}
			}
			return _coll18.ToSet()
		}())
		var _235_v129 _dafny.Set
		_ = _235_v129
		_235_v129 = _dafny.SetOf(_231_v125, _231_v125)
		_109_v40 = ((((func() _dafny.Set {
			if (_233_v128).Contains(_dafny.IntOfInt64(398)) {
				return (_233_v128).Get(_dafny.IntOfInt64(398)).(_dafny.Set)
			}
			return _235_v129
		})()).Difference(Companion_Default___.Fm17(_109_v40, p0, (_dafny.Zero).Minus((_171_cf22).F6()), globalState))).Union(_235_v129)).Cardinality()
		var _236_v130 _dafny.MultiSet
		_ = _236_v130
		_236_v130 = _dafny.MultiSetOf((_171_cf22).F6())
		(globalState).F0 = ((_171_cf22).Fm1(_236_v130, _39_v0, globalState)).Times((_171_cf22).F6())
	}
	var _237_v131 _dafny.Sequence
	_ = _237_v131
	_237_v131 = _dafny.UnicodeSeqOfUtf8Bytes("xukko")
	if (func() bool {
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_237_v131, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(910))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_238_i11 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		}))) {
			return ((_dafny.Zero).Minus(_109_v40)).Cmp(_109_v40) == 0
		}
		return ((_dafny.Zero).Minus(_109_v40)).Cmp(_109_v40) < 0
	})() {
		r0 = p0
		if ((_109_v40).Minus(_109_v40)).Cmp(_109_v40) >= 0 {
			var _239_v132 D4
			_ = _239_v132
			_239_v132 = Companion_D4_.Create_DC15_(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rq")).Cardinality()))
			var _240_v133 _dafny.Map
			_ = _240_v133
			_240_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(p0, p0, _dafny.IntOfUint32((_237_v131).Cardinality()), globalState), (_239_v132).Dtor_cf23())
			var _241_v134 _dafny.Set
			_ = _241_v134
			_241_v134 = _dafny.SetOf(_109_v40, _109_v40)
			var _242_v135 _dafny.Set
			_ = _242_v135
			_242_v135 = _dafny.SetOf(p0, p0)
			var _243_v136 *C0
			_ = _243_v136
			var _nw35 *C0 = New_C0_()
			_ = _nw35
			_nw35.Ctor__(_109_v40, _242_v135)
			_243_v136 = _nw35
			var _244_v137 _dafny.Map
			_ = _244_v137
			_244_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_v136, _109_v40)
			var _245_v138 _dafny.MultiSet
			_ = _245_v138
			_245_v138 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (func() _dafny.Int {
				if (_244_v137).Contains(_243_v136) {
					return (_244_v137).Get(_243_v136).(_dafny.Int)
				}
				return (_243_v136).F6()
			})())
			_240_v133 = (_240_v133).Update((_241_v134).IsDisjointFrom(_dafny.SetOf(_109_v40, (_dafny.SetOf(p0)).Cardinality(), _dafny.IntOfUint32((_237_v131).Cardinality()), (func() _dafny.Int {
				if (_245_v138).Contains(_109_v40) {
					return (_245_v138).Multiplicity(_109_v40)
				}
				return _109_v40
			})())), p0)
			var _246_v139 _dafny.Map
			_ = _246_v139
			_246_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm18(_109_v40, globalState), _109_v40)
			var _247_v140 _dafny.CodePoint
			_ = _247_v140
			_247_v140 = _dafny.CodePoint('l')
			var _248_v141 _dafny.Map
			_ = _248_v141
			_248_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v40, p0)
			(globalState).F0 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_246_v139).Contains(_247_v140) {
					return (_246_v139).Get(_247_v140).(_dafny.Int)
				}
				return ((_248_v141).Update(_109_v40, p0)).Cardinality()
			})())
			var _249_v142 *C0
			_ = _249_v142
			var _nw36 *C0 = New_C0_()
			_ = _nw36
			_nw36.Ctor__(((_243_v136).F6()).Times(_109_v40), (_243_v136).F7())
			_249_v142 = _nw36
			_247_v140 = _247_v140
			var _250_v143 _dafny.MultiSet
			_ = _250_v143
			_250_v143 = _dafny.MultiSetOf(p0)
			var _251_v144 _dafny.Map
			_ = _251_v144
			_251_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_250_v143, _247_v140)
			var _252_v145 _dafny.Sequence
			_ = _252_v145
			_252_v145 = _dafny.SeqOf(p0, p0)
			_251_v144 = (_251_v144).Update(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_252_v145, _252_v145)), _247_v140)
		} else {
			r0 = p0
			var _253_v146 _dafny.Array
			_ = _253_v146
			var _nw37 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(12))
			_ = _nw37
			_253_v146 = _nw37
			var _254_v147 _dafny.Set
			_ = _254_v147
			_254_v147 = _dafny.SetOf(p0)
			var _255_v148 *C0
			_ = _255_v148
			var _nw38 *C0 = New_C0_()
			_ = _nw38
			_nw38.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _109_v40)).Cardinality(), _254_v147)
			_255_v148 = _nw38
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_253_v146), 0))
			_ = _index20
			(_253_v146).ArraySet1(_255_v148, (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_253_v146), 0))
			_ = _index21
			var _rhs17 *C0 = _255_v148
			_ = _rhs17
			var _rhs18 _dafny.Int = _109_v40
			_ = _rhs18
			var _rhs19 _dafny.Int = (_255_v148).F6()
			_ = _rhs19
			var _rhs20 *C0 = _255_v148
			_ = _rhs20
			var _lhs13 _dafny.Array = _253_v146
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_253_v146), 0))
			_ = _lhs14
			var _lhs15 *GlobalState = globalState
			_ = _lhs15
			var _lhs16 *GlobalState = globalState
			_ = _lhs16
			(_lhs13).ArraySet1(_rhs17, (_lhs14).Int())
			_lhs15.F4 = _rhs18
			_lhs16.F4 = _rhs19
			_255_v148 = _rhs20
			var _256_v149 _dafny.Array
			_ = _256_v149
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_4
			var _nw39 _dafny.Array
			_ = _nw39
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw39 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.MultiSet = (func(_257_v148 *C0, _258_v40 _dafny.Int, _259_p0 bool, _260_v131 _dafny.Sequence) func(_dafny.Int) _dafny.MultiSet {
					return func(_261_i12 _dafny.Int) _dafny.MultiSet {
						return (_dafny.MultiSetOf((_257_v148).F6(), (_257_v148).F6())).Intersection(_dafny.MultiSetOf((_257_v148).F6(), _258_v40, (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_259_p0)).Cardinality()), _258_v40, _dafny.IntOfUint32((_260_v131).Cardinality()), _258_v40)).Cardinality()))
					}
				})(_255_v148, _109_v40, p0, _237_v131)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw39 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw39).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw39).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_256_v149 = _nw39
			var _262_v150 _dafny.MultiSet
			_ = _262_v150
			_262_v150 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_109_v40, (_255_v148).F6())).Cardinality()), (_dafny.Zero).Minus(_109_v40))
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_256_v149), 0))
			_ = _index22
			(_256_v149).ArraySet1(_262_v150, (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_256_v149), 0))
			_ = _index23
			(_256_v149).ArraySet1(((func() _dafny.MultiSet {
				if p0 {
					return _262_v150
				}
				return _262_v150
			})()).Union(_262_v150), (_index23).Int())
			var _263_v151 _dafny.MultiSet
			_ = _263_v151
			_263_v151 = _dafny.MultiSetOf(!(p0))
			(globalState).F0 = ((_255_v148).F6()).Plus((_263_v151).Cardinality())
			r0 = _dafny.Companion_Sequence_.IsPrefixOf(_237_v131, _237_v131)
		}
		(globalState).F0 = (func() _dafny.Int {
			if p0 {
				return (_dafny.Zero).Minus(_109_v40)
			}
			return _109_v40
		})()
		var _264_v152 *C0
		_ = _264_v152
		var _nw40 *C0 = New_C0_()
		_ = _nw40
		_nw40.Ctor__(_109_v40, _dafny.SetOf(p0))
		_264_v152 = _nw40
		var _265_v153 _dafny.Map
		_ = _265_v153
		_265_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_264_v152).F6())
		var _266_v154 _dafny.Sequence
		_ = _266_v154
		_266_v154 = _dafny.SeqOf(_109_v40, (_264_v152).Fm2(p0, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_109_v40), (Companion_Default___.SafeIndex(_109_v40, _dafny.IntOfUint32((_dafny.SeqOf(_109_v40)).Cardinality()))).Uint32(), _109_v40), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(629))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}((func(_267_v152 *C0) func(_dafny.Int) _dafny.Int {
			return func(_268_i14 _dafny.Int) _dafny.Int {
				return (_267_v152).F6()
			}
		})(_264_v152))), _265_v153, globalState))
		_109_v40 = Companion_Default___.SafeModuloInt(_109_v40, (_264_v152).Fm2(true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(27))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_269_v152 *C0) func(_dafny.Int) _dafny.Int {
			return func(_270_i13 _dafny.Int) _dafny.Int {
				return (_269_v152).F6()
			}
		})(_264_v152))), _266_v154, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((Companion_Default___.Fm3(_dafny.SeqOf(p0), globalState)).Cardinality())), globalState))
	} else {
		(globalState).F0 = _109_v40
		var _271_v155 _dafny.Array
		_ = _271_v155
		var _nw41 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw41
		_271_v155 = _nw41
		var _272_v156 _dafny.Map
		_ = _272_v156
		_272_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_271_v155), _dafny.UnicodeSeqOfUtf8Bytes("x"))
		var _273_v157 _dafny.Sequence
		_ = _273_v157
		_273_v157 = _dafny.SeqOf(_271_v155, _271_v155)
		var _274_v158 _dafny.CodePoint
		_ = _274_v158
		_274_v158 = _dafny.CodePoint('t')
		var _275_v159 _dafny.Sequence
		_ = _275_v159
		_275_v159 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("jxrvpul"), _dafny.Companion_Sequence_.Update(_237_v131, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_272_v156).Contains(_dafny.Companion_Sequence_.Update(_273_v157, (Companion_Default___.SafeIndex(_109_v40, _dafny.IntOfUint32((_273_v157).Cardinality()))).Uint32(), _271_v155)) {
				return (_272_v156).Get(_dafny.Companion_Sequence_.Update(_273_v157, (Companion_Default___.SafeIndex(_109_v40, _dafny.IntOfUint32((_273_v157).Cardinality()))).Uint32(), _271_v155)).(_dafny.Sequence)
			}
			return _237_v131
		})()).Cardinality()), _dafny.IntOfUint32((_237_v131).Cardinality()))).Uint32(), _274_v158))
		var _276_v160 _dafny.Sequence
		_ = _276_v160
		_276_v160 = _dafny.SeqOf(p0)
		var _277_v161 _dafny.Array
		_ = _277_v161
		var _nwElement0_7 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_275_v159, _275_v159)
		_ = _nwElement0_7
		var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(29))
		_ = _nw42
		(_nw42).ArraySet1(_nwElement0_7, 0)
		(_nw42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_275_v159, _275_v159), 1)
		(_nw42).ArraySet1(_dafny.SeqOf(_237_v131, _dafny.Companion_Sequence_.Update(_237_v131, (Companion_Default___.SafeIndex(_109_v40, _dafny.IntOfUint32((_237_v131).Cardinality()))).Uint32(), _274_v158)), 2)
		(_nw42).ArraySet1(_275_v159, 3)
		(_nw42).ArraySet1(_275_v159, 4)
		(_nw42).ArraySet1(_dafny.SeqOf(_237_v131, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(328))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}((func(_278_v131 _dafny.Sequence, _279_p0 bool, _280_v158 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_281_i15 _dafny.Int) _dafny.CodePoint {
				return (Companion_D6_.Create_DC19_(_dafny.IntOfUint32((_278_v131).Cardinality()), _279_p0, _280_v158)).Dtor_cf33()
			}
		})(_237_v131, p0, _274_v158)))), 5)
		(_nw42).ArraySet1(_275_v159, 6)
		(_nw42).ArraySet1(Companion_Default___.Fm19(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-351))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_282_v40 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_283_i16 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_282_v40), true)).Cardinality()
			}
		})(_109_v40)))).Cardinality()), globalState), 7)
		(_nw42).ArraySet1(_275_v159, 8)
		(_nw42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_275_v159, _275_v159), 9)
		(_nw42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_275_v159, _275_v159), 10)
		(_nw42).ArraySet1(_275_v159, 11)
		(_nw42).ArraySet1(_275_v159, 12)
		(_nw42).ArraySet1(_275_v159, 13)
		(_nw42).ArraySet1(_275_v159, 14)
		(_nw42).ArraySet1(_dafny.Companion_Sequence_.Update(_275_v159, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_237_v131).Cardinality()), _dafny.IntOfUint32((_275_v159).Cardinality()))).Uint32(), _237_v131), 15)
		(_nw42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("qrr")), _275_v159), 16)
		(_nw42).ArraySet1(_275_v159, 17)
		(_nw42).ArraySet1(_275_v159, 18)
		(_nw42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_275_v159, _dafny.Companion_Sequence_.Update(_275_v159, (Companion_Default___.SafeIndex(_109_v40, _dafny.IntOfUint32((_275_v159).Cardinality()))).Uint32(), _dafny.UnicodeSeqOfUtf8Bytes("igyetaiv"))), 19)
		(_nw42).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(905))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_284_v131 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_285_i17 _dafny.Int) _dafny.Sequence {
				return _284_v131
			}
		})(_237_v131))), 20)
		(_nw42).ArraySet1(_dafny.SeqOf(Companion_Default___.Fm3(_276_v160, globalState)), 21)
		(_nw42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_275_v159, _275_v159), 22)
		(_nw42).ArraySet1(_dafny.SeqOf(_237_v131, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(353))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_286_v158 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_287_i18 _dafny.Int) _dafny.CodePoint {
				return _286_v158
			}
		})(_274_v158)))), 23)
		(_nw42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_275_v159, _275_v159), 24)
		(_nw42).ArraySet1(_275_v159, 25)
		(_nw42).ArraySet1(_275_v159, 26)
		(_nw42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm19(p0, _109_v40, globalState), _275_v159), 27)
		(_nw42).ArraySet1(_dafny.SeqOf(_237_v131, _237_v131), 28)
		_277_v161 = _nw42
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_277_v161), 0))
		_ = _index24
		(_277_v161).ArraySet1(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(480))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}(func(_288_i19 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		}))), (_index24).Int())
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_277_v161), 0))
		_ = _index25
		(_277_v161).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_275_v159, _275_v159), (_index25).Int())
		var _289_v162 _dafny.Array
		_ = _289_v162
		var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
		_ = _nw43
		_289_v162 = _nw43
		var _290_v163 _dafny.Array
		_ = _290_v163
		var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
		_ = _nw44
		_290_v163 = _nw44
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_289_v162), 0))
		_ = _index26
		(_289_v162).ArraySet1(_290_v163, (_index26).Int())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_289_v162), 0))
		_ = _index27
		var _nw45 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
		_ = _nw45
		(_289_v162).ArraySet1(_nw45, (_index27).Int())
		var _291_v165 _dafny.Array
		_ = _291_v165
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_5
		var _nw46 _dafny.Array
		_ = _nw46
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw46 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Sequence = (func(_292_p0 bool, _293_v131 _dafny.Sequence, _294_v158 _dafny.CodePoint, _295_v40 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_296_i20 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(Companion_D1_.Create_DC3_(_292_p0, _292_p0, _293_v131, true, _dafny.IntOfUint32((_293_v131).Cardinality())), Companion_D1_.Create_DC3_(_292_p0, true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(235))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg20 _dafny.Int) interface{} {
							return coer20(arg20)
						}
					}((func(_297_v158 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_298_i21 _dafny.Int) _dafny.CodePoint {
							return _297_v158
						}
					})(_294_v158))), _292_p0, _295_v40), Companion_D1_.Create_DC3_(_292_p0, _292_p0, _293_v131, !(false), (func() _dafny.Set {
						var _coll19 = _dafny.NewBuilder()
						_ = _coll19
						for _iter19 := _dafny.Iterate((_dafny.SeqOf(_294_v158, _294_v158)).Elements()); ; {
							_compr_19, _ok19 := _iter19()
							if !_ok19 {
								break
							}
							var _299_v164 _dafny.CodePoint
							_299_v164 = interface{}(_compr_19).(_dafny.CodePoint)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_294_v158, _294_v158), _299_v164) {
								_coll19.Add(_299_v164)
							}
						}
						return _coll19.ToSet()
					}()).Cardinality()))
				}
			})(p0, _237_v131, _274_v158, _109_v40)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw46 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw46).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw46).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_291_v165 = _nw46
		var _300_v166 _dafny.Map
		_ = _300_v166
		_300_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.MultiSetFromSeq(_276_v160))
		var _rhs21 _dafny.Int = ((_300_v166).Cardinality()).Times(_109_v40)
		_ = _rhs21
		var _rhs22 _dafny.Int = _109_v40
		_ = _rhs22
		var _rhs23 bool = p0
		_ = _rhs23
		var _rhs24 _dafny.Array = (func() _dafny.Array {
			if p0 {
				return _291_v165
			}
			return _291_v165
		})()
		_ = _rhs24
		var _lhs17 *GlobalState = globalState
		_ = _lhs17
		_109_v40 = _rhs21
		_lhs17.F0 = _rhs22
		r0 = _rhs23
		_291_v165 = _rhs24
		var _301_v167 D2
		_ = _301_v167
		_301_v167 = Companion_D2_.Create_DC7_(_dafny.IntOfUint32((_237_v131).Cardinality()))
		var _302_v168 _dafny.Set
		_ = _302_v168
		_302_v168 = _dafny.SetOf(p0)
		var _303_v169 *C0
		_ = _303_v169
		var _nw47 *C0 = New_C0_()
		_ = _nw47
		_nw47.Ctor__((_109_v40).Plus((_301_v167).Dtor_cf15()), _302_v168)
		_303_v169 = _nw47
	}
	var _304_v170 _dafny.CodePoint
	_ = _304_v170
	_304_v170 = _dafny.CodePoint('t')
	var _305_v171 D6
	_ = _305_v171
	_305_v171 = Companion_D6_.Create_DC19_(_109_v40, Companion_Default___.Fm0(!(!(p0)), p0, _109_v40, globalState), _304_v170)
	var _pat_let_tv5 = _109_v40
	_ = _pat_let_tv5
	var _pat_let_tv6 = _39_v0
	_ = _pat_let_tv6
	var _pat_let_tv7 = _39_v0
	_ = _pat_let_tv7
	var _pat_let_tv8 = p0
	_ = _pat_let_tv8
	var _pat_let_tv9 = _109_v40
	_ = _pat_let_tv9
	var _pat_let_tv10 = _109_v40
	_ = _pat_let_tv10
	var _pat_let_tv11 = _109_v40
	_ = _pat_let_tv11
	var _pat_let_tv12 = _109_v40
	_ = _pat_let_tv12
	var _pat_let_tv13 = _109_v40
	_ = _pat_let_tv13
	r0 = func(_source6 D6) bool {
		if _source6.Is_DC19() {
			var _306___mcc_h22 _dafny.Int = _source6.Get_().(D6_DC19).Cf31
			_ = _306___mcc_h22
			var _307___mcc_h23 bool = _source6.Get_().(D6_DC19).Cf32
			_ = _307___mcc_h23
			var _308___mcc_h24 _dafny.CodePoint = _source6.Get_().(D6_DC19).Cf33
			_ = _308___mcc_h24
			var _309_cf33 _dafny.CodePoint = _308___mcc_h24
			_ = _309_cf33
			var _310_cf32 bool = _307___mcc_h23
			_ = _310_cf32
			var _311_cf31 _dafny.Int = _306___mcc_h22
			_ = _311_cf31
			return _310_cf32
		} else if _source6.Is_DC20() {
			var _312___mcc_h25 bool = _source6.Get_().(D6_DC20).Cf34
			_ = _312___mcc_h25
			var _313_cf34 bool = _312___mcc_h25
			_ = _313_cf34
			return (_pat_let_tv5).Cmp((func() _dafny.Map {
				var _coll20 = _dafny.NewMapBuilder()
				_ = _coll20
				for _iter20 := _dafny.Iterate((_dafny.SeqOf(_pat_let_tv6)).Elements()); ; {
					_compr_20, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _314_v172 D0
					_314_v172 = interface{}(_compr_20).(D0)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_pat_let_tv7), _314_v172) {
						_coll20.Add(_314_v172, _pat_let_tv8)
					}
				}
				return _coll20.ToMap()
			}()).Cardinality()) > 0
		} else if _source6.Is_DC21() {
			return (_dafny.MultiSetOf(_dafny.IntOfInt64(14), _pat_let_tv9)).IsSubsetOf(_dafny.MultiSetOf(_pat_let_tv10, _pat_let_tv11))
		} else {
			var _315___mcc_h26 _dafny.CodePoint = _source6.Get_().(D6_DC18).Cf30
			_ = _315___mcc_h26
			var _316_cf30 _dafny.CodePoint = _315___mcc_h26
			_ = _316_cf30
			return (Companion_D3_.Create_DC10_(true, _pat_let_tv12, _pat_let_tv13)).Dtor_cf18()
		}
	}(_305_v171)
	var _317_i22 _dafny.Int
	_ = _317_i22
	_317_i22 = _dafny.Zero
	{
		for !(p0) {
			{
				if (_317_i22).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_317_i22 = (_317_i22).Plus(_dafny.One)
				r0 = p0
				r0 = p0
				(globalState).F0 = _109_v40
				var _318_v173 _dafny.Set
				_ = _318_v173
				_318_v173 = _dafny.SetOf(p0, p0)
				var _319_v174 *C0
				_ = _319_v174
				var _nw48 *C0 = New_C0_()
				_ = _nw48
				_nw48.Ctor__((_dafny.Zero).Minus(Companion_Default___.Fm4(_109_v40, p0, globalState)), _318_v173)
				_319_v174 = _nw48
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	r0 = p0
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _320_v0 _dafny.Int
	_ = _320_v0
	_320_v0 = _dafny.IntOfInt64(-45)
	var _321_v1 bool
	_ = _321_v1
	_321_v1 = true
	var _322_v2 _dafny.Sequence
	_ = _322_v2
	_322_v2 = _dafny.SeqOf(_321_v1, _321_v1)
	var _323_v3 _dafny.Map
	_ = _323_v3
	_323_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v0, _322_v2)
	var _324_globalState *GlobalState
	_ = _324_globalState
	var _nw49 *GlobalState = New_GlobalState_()
	_ = _nw49
	_nw49.Ctor__(_dafny.IntOfInt64(977), _dafny.IntOfInt64(424), true, _dafny.IntOfInt64(358), _dafny.IntOfInt64(-429), _323_v3)
	_324_globalState = _nw49
	var _325_i0 _dafny.Int
	_ = _325_i0
	_325_i0 = _dafny.Zero
	{
		for (_dafny.IntOfInt64(671)).Cmp((func() _dafny.Int {
			if _321_v1 {
				return _dafny.IntOfInt64(548)
			}
			return _320_v0
		})()) >= 0 {
			{
				if (_325_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_325_i0 = (_325_i0).Plus(_dafny.One)
				var _326_v4 _dafny.CodePoint
				_ = _326_v4
				_326_v4 = _dafny.CodePoint('d')
				var _327_v5 _dafny.Sequence
				_ = _327_v5
				_327_v5 = _dafny.UnicodeSeqOfUtf8Bytes("ja")
				_326_v4 = (_327_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_321_v1), _dafny.SeqOf(_321_v1))).Cardinality()), _dafny.IntOfUint32((_327_v5).Cardinality()))).Uint32()).(_dafny.CodePoint)
				var _328_v6 _dafny.Array
				_ = _328_v6
				var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
				_ = _nw50
				_328_v6 = _nw50
				var _329_v7 _dafny.MultiSet
				_ = _329_v7
				_329_v7 = _dafny.MultiSetOf(_328_v6, _328_v6, _328_v6)
				var _330_v8 D0
				_ = _330_v8
				_330_v8 = Companion_D0_.Create_DC0_(_321_v1)
				var _331_v9 _dafny.Array
				_ = _331_v9
				var _nwElement0_8 bool = !(_321_v1)
				_ = _nwElement0_8
				var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(29))
				_ = _nw51
				(_nw51).ArraySet1(_nwElement0_8, 0)
				(_nw51).ArraySet1(_321_v1, 1)
				(_nw51).ArraySet1(_321_v1, 2)
				(_nw51).ArraySet1(_321_v1, 3)
				(_nw51).ArraySet1(_321_v1, 4)
				(_nw51).ArraySet1(_321_v1, 5)
				(_nw51).ArraySet1(_321_v1, 6)
				(_nw51).ArraySet1(_321_v1, 7)
				(_nw51).ArraySet1(true, 8)
				(_nw51).ArraySet1(_321_v1, 9)
				(_nw51).ArraySet1(_321_v1, 10)
				(_nw51).ArraySet1(true, 11)
				(_nw51).ArraySet1(_321_v1, 12)
				(_nw51).ArraySet1(_321_v1, 13)
				(_nw51).ArraySet1(_321_v1, 14)
				(_nw51).ArraySet1(Companion_Default___.Fm0(_321_v1, _321_v1, _320_v0, _324_globalState), 15)
				(_nw51).ArraySet1(_321_v1, 16)
				(_nw51).ArraySet1(_321_v1, 17)
				(_nw51).ArraySet1(_321_v1, 18)
				(_nw51).ArraySet1(!(_321_v1), 19)
				(_nw51).ArraySet1(false, 20)
				(_nw51).ArraySet1(_321_v1, 21)
				(_nw51).ArraySet1(_321_v1, 22)
				(_nw51).ArraySet1(_321_v1, 23)
				(_nw51).ArraySet1(true, 24)
				(_nw51).ArraySet1(_321_v1, 25)
				(_nw51).ArraySet1(_321_v1, 26)
				(_nw51).ArraySet1((_330_v8).Dtor_cf0(), 27)
				(_nw51).ArraySet1(_321_v1, 28)
				_331_v9 = _nw51
				var _332_v10 _dafny.Array
				_ = _332_v10
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_6
				var _nw52 _dafny.Array
				_ = _nw52
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw52 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) _dafny.MultiSet = (func(_333_v1 bool) func(_dafny.Int) _dafny.MultiSet {
						return func(_334_i1 _dafny.Int) _dafny.MultiSet {
							return _dafny.MultiSetOf((_dafny.MultiSetOf(_333_v1)).Cardinality())
						}
					})(_321_v1)
					_ = _init6
					var _element0_6 = _init6(_dafny.Zero)
					_ = _element0_6
					_nw52 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
					(_nw52).ArraySet1(_element0_6, 0)
					var _nativeLen0_6 = (_len0_6).Int()
					_ = _nativeLen0_6
					for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
						(_nw52).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
					}
				}
				_332_v10 = _nw52
				var _335_v11 D0
				_ = _335_v11
				_335_v11 = Companion_D0_.Create_DC1_(_329_v7, _331_v9, _332_v10)
				var _source7 D0 = _335_v11
				_ = _source7
				if _source7.Is_DC0() {
					var _336___mcc_h0 bool = _source7.Get_().(D0_DC0).Cf0
					_ = _336___mcc_h0
					var _337_cf0 bool = _336___mcc_h0
					_ = _337_cf0
					var _338_v12 _dafny.Set
					_ = _338_v12
					_338_v12 = _dafny.SetOf(_320_v0, _320_v0)
					var _339_v13 _dafny.Sequence
					_ = _339_v13
					_339_v13 = _dafny.SeqOf(_338_v12, _338_v12)
					_339_v13 = _339_v13
					var _340_v14 _dafny.Set
					_ = _340_v14
					_340_v14 = _dafny.SetOf(true, _321_v1, _321_v1)
					var _341_v15 *C0
					_ = _341_v15
					var _nw53 *C0 = New_C0_()
					_ = _nw53
					_nw53.Ctor__(Companion_Default___.SafeDivisionInt(_320_v0, _320_v0), (_340_v14).Difference(_340_v14))
					_341_v15 = _nw53
					var _342_v16 _dafny.Map
					_ = _342_v16
					_342_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_321_v1, _341_v15)
					_342_v16 = (_342_v16).Update(_337_cf0, (func() *C0 {
						if (_342_v16).Contains(_321_v1) {
							return (_342_v16).Get(_321_v1).(*C0)
						}
						return _341_v15
					})())
					_328_v6 = (func() _dafny.Array {
						if true {
							return _328_v6
						}
						return _328_v6
					})()
				} else {
					var _343___mcc_h1 _dafny.MultiSet = _source7.Get_().(D0_DC1).Cf1
					_ = _343___mcc_h1
					var _344___mcc_h2 _dafny.Array = _source7.Get_().(D0_DC1).Cf2
					_ = _344___mcc_h2
					var _345___mcc_h3 _dafny.Array = _source7.Get_().(D0_DC1).Cf3
					_ = _345___mcc_h3
					var _346_cf3 _dafny.Array = _345___mcc_h3
					_ = _346_cf3
					var _347_cf2 _dafny.Array = _344___mcc_h2
					_ = _347_cf2
					var _348_cf1 _dafny.MultiSet = _343___mcc_h1
					_ = _348_cf1
					_328_v6 = _328_v6
					var _349_v17 _dafny.MultiSet
					_ = _349_v17
					_349_v17 = _dafny.MultiSetOf(_321_v1)
					_349_v17 = (_349_v17).Union(_349_v17)
					var _350_v18 bool
					_ = _350_v18
					var _out0 bool
					_ = _out0
					_out0 = Companion_Default___.M0(!((func() bool {
						if _321_v1 {
							return _321_v1
						}
						return _321_v1
					})()), _324_globalState)
					_350_v18 = _out0
					_320_v0 = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nroqns")).Cardinality())
				}
				(_324_globalState).F4 = _320_v0
				var _351_v19 _dafny.Set
				_ = _351_v19
				_351_v19 = _dafny.SetOf(_320_v0, (_320_v0).Minus(_dafny.IntOfUint32((_327_v5).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_327_v5, _327_v5)).Cardinality()), _320_v0, Companion_Default___.SafeDivisionInt(_320_v0, _320_v0))
				var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_328_v6), 0))
				_ = _index28
				(_328_v6).ArraySet1(_320_v0, (_index28).Int())
				var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_328_v6), 0))
				_ = _index29
				var _rhs25 bool = _321_v1
				_ = _rhs25
				var _rhs26 _dafny.Set = _351_v19
				_ = _rhs26
				var _rhs27 _dafny.Int = _320_v0
				_ = _rhs27
				var _lhs18 _dafny.Array = _328_v6
				_ = _lhs18
				var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_328_v6), 0))
				_ = _lhs19
				_321_v1 = _rhs25
				_351_v19 = _rhs26
				(_lhs18).ArraySet1(_rhs27, (_lhs19).Int())
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _352_v20 _dafny.Sequence
	_ = _352_v20
	_352_v20 = _dafny.UnicodeSeqOfUtf8Bytes("ss")
	var _353_v21 _dafny.MultiSet
	_ = _353_v21
	_353_v21 = _dafny.MultiSetOf(_320_v0, _320_v0)
	var _source8 D0 = Companion_D0_.Create_DC0_((_dafny.IntOfUint32((_352_v20).Cardinality())).Cmp((_353_v21).Cardinality()) > 0)
	_ = _source8
	if _source8.Is_DC0() {
		var _354___mcc_h4 bool = _source8.Get_().(D0_DC0).Cf0
		_ = _354___mcc_h4
		var _355_cf0 bool = _354___mcc_h4
		_ = _355_cf0
		var _356_v22 _dafny.Array
		_ = _356_v22
		var _nw54 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
		_ = _nw54
		_356_v22 = _nw54
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_356_v22), 0))
		_ = _index30
		(_356_v22).ArraySet1((_355_cf0) || (true), (_index30).Int())
		var _357_v23 _dafny.CodePoint
		_ = _357_v23
		_357_v23 = _dafny.CodePoint('s')
		var _358_v24 _dafny.MultiSet
		_ = _358_v24
		_358_v24 = _dafny.MultiSetOf(Companion_Default___.Fm3(_322_v2, _324_globalState))
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_356_v22), 0))
		_ = _index31
		var _rhs28 bool = _321_v1
		_ = _rhs28
		var _rhs29 _dafny.CodePoint = _357_v23
		_ = _rhs29
		var _rhs30 _dafny.Int = _320_v0
		_ = _rhs30
		var _rhs31 _dafny.Int = (_358_v24).Cardinality()
		_ = _rhs31
		var _lhs20 _dafny.Array = _356_v22
		_ = _lhs20
		var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_356_v22), 0))
		_ = _lhs21
		var _lhs22 *GlobalState = _324_globalState
		_ = _lhs22
		var _lhs23 *GlobalState = _324_globalState
		_ = _lhs23
		(_lhs20).ArraySet1(_rhs28, (_lhs21).Int())
		_357_v23 = _rhs29
		_lhs22.F4 = _rhs30
		_lhs23.F0 = _rhs31
		var _359_v25 _dafny.Map
		_ = _359_v25
		_359_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_355_cf0, Companion_Default___.SafeDivisionInt(_320_v0, _320_v0))
		var _360_v26 _dafny.MultiSet
		_ = _360_v26
		_360_v26 = _dafny.MultiSetOf((_356_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_356_v22), 0))).Int()).(bool))
		_359_v25 = (_359_v25).Update((_dafny.MultiSetFromSeq(_322_v2)).IsDisjointFrom((_360_v26).Update(!(!(true)), Companion_Default___.Abs(_320_v0))), (_dafny.Zero).Minus((_dafny.Zero).Minus(_320_v0)))
		_357_v23 = _357_v23
		var _361_v27 _dafny.Map
		_ = _361_v27
		_361_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_356_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_356_v22), 0))).Int()).(bool), _dafny.SetOf(_321_v1, (_356_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_356_v22), 0))).Int()).(bool), _321_v1, (_356_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_356_v22), 0))).Int()).(bool), _355_cf0))
		var _362_v28 _dafny.Set
		_ = _362_v28
		_362_v28 = _dafny.SetOf(false, false, _355_cf0)
		var _363_v29 _dafny.Array
		_ = _363_v29
		var _nwElement0_9 _dafny.Int = _320_v0
		_ = _nwElement0_9
		var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(22))
		_ = _nw55
		(_nw55).ArraySet1(_nwElement0_9, 0)
		(_nw55).ArraySet1(_320_v0, 1)
		(_nw55).ArraySet1(_320_v0, 2)
		(_nw55).ArraySet1(_320_v0, 3)
		(_nw55).ArraySet1((_320_v0).Minus(Companion_Default___.Fm4(_320_v0, _355_cf0, _324_globalState)), 4)
		(_nw55).ArraySet1(_320_v0, 5)
		(_nw55).ArraySet1(_320_v0, 6)
		(_nw55).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_352_v20, _352_v20)).Cardinality()), 7)
		(_nw55).ArraySet1(Companion_Default___.SafeDivisionInt(_320_v0, _320_v0), 8)
		(_nw55).ArraySet1(_320_v0, 9)
		(_nw55).ArraySet1(_320_v0, 10)
		(_nw55).ArraySet1(_dafny.IntOfInt64(233), 11)
		(_nw55).ArraySet1(Companion_Default___.SafeDivisionInt(_320_v0, _320_v0), 12)
		(_nw55).ArraySet1(_320_v0, 13)
		(_nw55).ArraySet1((_320_v0).Minus(_320_v0), 14)
		(_nw55).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_320_v0)), 15)
		(_nw55).ArraySet1(_320_v0, 16)
		(_nw55).ArraySet1((((_361_v27).Update(_321_v1, _362_v28)).Cardinality()).Times(_320_v0), 17)
		(_nw55).ArraySet1(Companion_Default___.SafeModuloInt(_320_v0, _320_v0), 18)
		(_nw55).ArraySet1(_320_v0, 19)
		(_nw55).ArraySet1(_320_v0, 20)
		(_nw55).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("oywkaa"), _dafny.UnicodeSeqOfUtf8Bytes("k"))).Cardinality()), 21)
		_363_v29 = _nw55
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_363_v29), 0))
		_ = _index32
		(_363_v29).ArraySet1(_320_v0, (_index32).Int())
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_363_v29), 0))
		_ = _index33
		(_363_v29).ArraySet1(Companion_Default___.SafeModuloInt(_320_v0, _320_v0), (_index33).Int())
	} else {
		var _364___mcc_h5 _dafny.MultiSet = _source8.Get_().(D0_DC1).Cf1
		_ = _364___mcc_h5
		var _365___mcc_h6 _dafny.Array = _source8.Get_().(D0_DC1).Cf2
		_ = _365___mcc_h6
		var _366___mcc_h7 _dafny.Array = _source8.Get_().(D0_DC1).Cf3
		_ = _366___mcc_h7
		var _367_cf3 _dafny.Array = _366___mcc_h7
		_ = _367_cf3
		var _368_cf2 _dafny.Array = _365___mcc_h6
		_ = _368_cf2
		var _369_cf1 _dafny.MultiSet = _364___mcc_h5
		_ = _369_cf1
		(_324_globalState).F4 = _320_v0
		var _370_v30 _dafny.Set
		_ = _370_v30
		_370_v30 = _dafny.SetOf(_321_v1)
		var _371_v31 *C0
		_ = _371_v31
		var _nw56 *C0 = New_C0_()
		_ = _nw56
		_nw56.Ctor__(_320_v0, _370_v30)
		_371_v31 = _nw56
		var _372_v32 _dafny.Array
		_ = _372_v32
		var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw57
		_372_v32 = _nw57
		var _373_v33 _dafny.Map
		_ = _373_v33
		_373_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_371_v31, _372_v32)
		_373_v33 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_371_v31, _372_v32)).Merge(_373_v33)
		var _374_v34 _dafny.CodePoint
		_ = _374_v34
		_374_v34 = _dafny.CodePoint('t')
		var _375_v35 _dafny.Array
		_ = _375_v35
		var _nwElement0_10 _dafny.Array = _372_v32
		_ = _nwElement0_10
		var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(16))
		_ = _nw58
		(_nw58).ArraySet1(_nwElement0_10, 0)
		(_nw58).ArraySet1(_372_v32, 1)
		(_nw58).ArraySet1(_372_v32, 2)
		(_nw58).ArraySet1(_372_v32, 3)
		(_nw58).ArraySet1(_372_v32, 4)
		(_nw58).ArraySet1(_372_v32, 5)
		(_nw58).ArraySet1(_372_v32, 6)
		(_nw58).ArraySet1(_372_v32, 7)
		(_nw58).ArraySet1(_372_v32, 8)
		(_nw58).ArraySet1(_372_v32, 9)
		(_nw58).ArraySet1(_372_v32, 10)
		(_nw58).ArraySet1(_372_v32, 11)
		(_nw58).ArraySet1((func() _dafny.Array {
			if _321_v1 {
				return _372_v32
			}
			return _372_v32
		})(), 12)
		(_nw58).ArraySet1(_372_v32, 13)
		(_nw58).ArraySet1(_372_v32, 14)
		(_nw58).ArraySet1(_372_v32, 15)
		_375_v35 = _nw58
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_375_v35), 0))
		_ = _index34
		(_375_v35).ArraySet1(_372_v32, (_index34).Int())
		var _376_v36 _dafny.Map
		_ = _376_v36
		_376_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_372_v32, _367_cf3)
		var _377_v37 D0
		_ = _377_v37
		_377_v37 = Companion_D0_.Create_DC1_(_369_cf1, _368_cf2, (func() _dafny.Array {
			if (_376_v36).Contains(_372_v32) {
				return (_376_v36).Get(_372_v32).(_dafny.Array)
			}
			return _367_cf3
		})())
		var _378_v38 _dafny.Sequence
		_ = _378_v38
		_378_v38 = _dafny.SeqOf(_368_cf2)
		var _379_v39 _dafny.Array
		_ = _379_v39
		var _nwElement0_11 _dafny.Array = _368_cf2
		_ = _nwElement0_11
		var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(28))
		_ = _nw59
		(_nw59).ArraySet1(_nwElement0_11, 0)
		(_nw59).ArraySet1(_368_cf2, 1)
		(_nw59).ArraySet1(_368_cf2, 2)
		(_nw59).ArraySet1(_368_cf2, 3)
		(_nw59).ArraySet1(_368_cf2, 4)
		(_nw59).ArraySet1(_368_cf2, 5)
		(_nw59).ArraySet1(_368_cf2, 6)
		(_nw59).ArraySet1(_368_cf2, 7)
		(_nw59).ArraySet1(_368_cf2, 8)
		(_nw59).ArraySet1(_368_cf2, 9)
		(_nw59).ArraySet1((func() _dafny.Array {
			if _321_v1 {
				return _368_cf2
			}
			return _368_cf2
		})(), 10)
		(_nw59).ArraySet1(_368_cf2, 11)
		(_nw59).ArraySet1((_377_v37).Dtor_cf2(), 12)
		(_nw59).ArraySet1(_368_cf2, 13)
		(_nw59).ArraySet1(_368_cf2, 14)
		(_nw59).ArraySet1(_368_cf2, 15)
		(_nw59).ArraySet1(_368_cf2, 16)
		(_nw59).ArraySet1(_368_cf2, 17)
		(_nw59).ArraySet1(_368_cf2, 18)
		(_nw59).ArraySet1(_368_cf2, 19)
		(_nw59).ArraySet1(_368_cf2, 20)
		(_nw59).ArraySet1(_368_cf2, 21)
		(_nw59).ArraySet1(_368_cf2, 22)
		(_nw59).ArraySet1(_368_cf2, 23)
		(_nw59).ArraySet1(_368_cf2, 24)
		(_nw59).ArraySet1(_368_cf2, 25)
		(_nw59).ArraySet1(_368_cf2, 26)
		(_nw59).ArraySet1((_378_v38).Select((Companion_Default___.SafeIndex(_320_v0, _dafny.IntOfUint32((_378_v38).Cardinality()))).Uint32()).(_dafny.Array), 27)
		_379_v39 = _nw59
		var _380_v40 _dafny.MultiSet
		_ = _380_v40
		_380_v40 = _dafny.MultiSetOf(false, _321_v1)
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_375_v35), 0))
		_ = _index35
		var _rhs32 _dafny.CodePoint = _374_v34
		_ = _rhs32
		var _rhs33 _dafny.Array = _372_v32
		_ = _rhs33
		var _rhs34 bool = _321_v1
		_ = _rhs34
		var _rhs35 _dafny.Array = _379_v39
		_ = _rhs35
		var _rhs36 _dafny.Int = ((func() _dafny.MultiSet {
			if !(_321_v1) {
				return _380_v40
			}
			return (func() _dafny.MultiSet {
				if _321_v1 {
					return (_380_v40).Update(true, Companion_Default___.Abs(_dafny.IntOfInt64(-722)))
				}
				return _380_v40
			})()
		})()).Cardinality()
		_ = _rhs36
		var _lhs24 _dafny.Array = _375_v35
		_ = _lhs24
		var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_375_v35), 0))
		_ = _lhs25
		_374_v34 = _rhs32
		(_lhs24).ArraySet1(_rhs33, (_lhs25).Int())
		_321_v1 = _rhs34
		_379_v39 = _rhs35
		_320_v0 = _rhs36
		var _381_v42 _dafny.Set
		_ = _381_v42
		_381_v42 = _dafny.SetOf((_dafny.Zero).Minus((_371_v31).F6()))
		_321_v1 = ((func() _dafny.Set {
			var _coll21 = _dafny.NewBuilder()
			_ = _coll21
			for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-485), _dafny.IntOfInt64(640))); ; {
				_compr_21, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _382_v41 _dafny.Int
				_382_v41 = interface{}(_compr_21).(_dafny.Int)
				if ((_dafny.IntOfInt64(-485)).Cmp(_382_v41) <= 0) && ((_382_v41).Cmp(_dafny.IntOfInt64(640)) < 0) {
					_coll21.Add((_382_v41).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_374_v34, _321_v1)).Cardinality()))
				}
			}
			return _coll21.ToSet()
		}()).Intersection(_381_v42)).IsProperSubsetOf(_381_v42)
	}
	var _383_v43 D1
	_ = _383_v43
	_383_v43 = Companion_D1_.Create_DC3_(_321_v1, _321_v1, _352_v20, _321_v1, _320_v0)
	var _pat_let_tv14 = _321_v1
	_ = _pat_let_tv14
	var _pat_let_tv15 = _321_v1
	_ = _pat_let_tv15
	var _pat_let_tv16 = _321_v1
	_ = _pat_let_tv16
	var _pat_let_tv17 = _321_v1
	_ = _pat_let_tv17
	if func(_source9 D1) bool {
		if _source9.Is_DC3() {
			var _384___mcc_h8 bool = _source9.Get_().(D1_DC3).Cf5
			_ = _384___mcc_h8
			var _385___mcc_h9 bool = _source9.Get_().(D1_DC3).Cf6
			_ = _385___mcc_h9
			var _386___mcc_h10 _dafny.Sequence = _source9.Get_().(D1_DC3).Cf7
			_ = _386___mcc_h10
			var _387___mcc_h11 bool = _source9.Get_().(D1_DC3).Cf8
			_ = _387___mcc_h11
			var _388___mcc_h12 _dafny.Int = _source9.Get_().(D1_DC3).Cf9
			_ = _388___mcc_h12
			var _389_cf9 _dafny.Int = _388___mcc_h12
			_ = _389_cf9
			var _390_cf8 bool = _387___mcc_h11
			_ = _390_cf8
			var _391_cf7 _dafny.Sequence = _386___mcc_h10
			_ = _391_cf7
			var _392_cf6 bool = _385___mcc_h9
			_ = _392_cf6
			var _393_cf5 bool = _384___mcc_h8
			_ = _393_cf5
			return _pat_let_tv14
		} else if _source9.Is_DC2() {
			var _394___mcc_h13 _dafny.Int = _source9.Get_().(D1_DC2).Cf4
			_ = _394___mcc_h13
			var _395_cf4 _dafny.Int = _394___mcc_h13
			_ = _395_cf4
			return _pat_let_tv15
		} else {
			var _396___mcc_h14 D1 = _source9.Get_().(D1_DC4).Cf10
			_ = _396___mcc_h14
			var _397_cf10 D1 = _396___mcc_h14
			_ = _397_cf10
			if _pat_let_tv16 {
				return true
			} else {
				return _pat_let_tv17
			}
		}
	}(_383_v43) {
		(_324_globalState).F4 = _320_v0
		var _398_v44 _dafny.MultiSet
		_ = _398_v44
		_398_v44 = _dafny.MultiSetOf(_321_v1, _321_v1, _321_v1)
		var _399_v45 _dafny.Map
		_ = _399_v45
		_399_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_352_v20, _321_v1)
		var _400_v46 *C0
		_ = _400_v46
		var _nw60 *C0 = New_C0_()
		_ = _nw60
		_nw60.Ctor__((_398_v44).Cardinality(), _dafny.SetOf(!((func() bool {
			if (_399_v45).Contains(_352_v20) {
				return (_399_v45).Get(_352_v20).(bool)
			}
			return _321_v1
		})()), _321_v1))
		_400_v46 = _nw60
		var _401_v47 _dafny.Array
		_ = _401_v47
		var _nw61 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
		_ = _nw61
		_401_v47 = _nw61
		_401_v47 = _401_v47
		(_324_globalState).F4 = _320_v0
		var _402_v48 _dafny.CodePoint
		_ = _402_v48
		_402_v48 = _dafny.CodePoint('b')
		_399_v45 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_352_v20, (Companion_Default___.SafeIndex(_320_v0, _dafny.IntOfUint32((_352_v20).Cardinality()))).Uint32(), _402_v48), false)).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(896))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}((func(_403_v48 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_404_i2 _dafny.Int) _dafny.CodePoint {
				return _403_v48
			}
		})(_402_v48))), _321_v1)
	} else {
		var _405_v49 _dafny.MultiSet
		_ = _405_v49
		_405_v49 = _dafny.MultiSetOf(_321_v1, _321_v1)
		if (((func() _dafny.MultiSet {
			if _321_v1 {
				return _405_v49
			}
			return _405_v49
		})()).Cardinality()).Cmp(_320_v0) <= 0 {
			_321_v1 = !(_321_v1)
			_321_v1 = _321_v1
			_321_v1 = false
			var _406_v50 _dafny.Array
			_ = _406_v50
			var _nwElement0_12 _dafny.Int = _320_v0
			_ = _nwElement0_12
			var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(3))
			_ = _nw62
			(_nw62).ArraySet1(_nwElement0_12, 0)
			(_nw62).ArraySet1((_dafny.IntOfInt64(274)).Minus((_383_v43).Dtor_cf9()), 1)
			(_nw62).ArraySet1((_320_v0).Minus(_320_v0), 2)
			_406_v50 = _nw62
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_406_v50), 0))
			_ = _index36
			(_406_v50).ArraySet1(_320_v0, (_index36).Int())
			var _407_v51 _dafny.Map
			_ = _407_v51
			_407_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm4(_320_v0, _321_v1, _324_globalState), _321_v1)
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_406_v50), 0))
			_ = _index37
			(_406_v50).ArraySet1((_407_v51).Cardinality(), (_index37).Int())
			var _408_v52 _dafny.Array
			_ = _408_v52
			var _nw63 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw63
			_408_v52 = _nw63
			var _409_v53 _dafny.Map
			_ = _409_v53
			_409_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_405_v49, _408_v52)
			(_324_globalState).F0 = (_409_v53).Cardinality()
		} else {
			_321_v1 = false
			var _410_v54 _dafny.Array
			_ = _410_v54
			var _nwElement0_13 bool = (_320_v0).Cmp(_320_v0) >= 0
			_ = _nwElement0_13
			var _nw64 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(7))
			_ = _nw64
			(_nw64).ArraySet1(_nwElement0_13, 0)
			(_nw64).ArraySet1(_321_v1, 1)
			(_nw64).ArraySet1(_321_v1, 2)
			(_nw64).ArraySet1(_321_v1, 3)
			(_nw64).ArraySet1(_321_v1, 4)
			(_nw64).ArraySet1(_321_v1, 5)
			(_nw64).ArraySet1(_321_v1, 6)
			_410_v54 = _nw64
			_410_v54 = (func() _dafny.Array {
				if false {
					return _410_v54
				}
				return _410_v54
			})()
			var _411_v55 _dafny.Sequence
			_ = _411_v55
			_411_v55 = _dafny.SeqOf(_320_v0)
			var _412_v56 _dafny.CodePoint
			_ = _412_v56
			_412_v56 = _dafny.CodePoint('f')
			var _413_v57 _dafny.Set
			_ = _413_v57
			_413_v57 = _dafny.SetOf(_412_v56, _412_v56, _412_v56)
			var _414_v58 _dafny.Sequence
			_ = _414_v58
			_414_v58 = _dafny.SeqOf(_dafny.IntOfUint32((_411_v55).Cardinality()), _320_v0, _320_v0, _320_v0, (_413_v57).Cardinality())
			var _415_v59 *C0
			_ = _415_v59
			var _nw65 *C0 = New_C0_()
			_ = _nw65
			_nw65.Ctor__((_411_v55).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm4(_320_v0, _321_v1, _324_globalState), _dafny.IntOfUint32((_411_v55).Cardinality()))).Uint32()).(_dafny.Int), _dafny.SetOf(!(_321_v1)))
			_415_v59 = _nw65
			(_324_globalState).F0 = _320_v0
			(_324_globalState).F4 = (_320_v0).Times((_383_v43).Dtor_cf9())
		}
		var _416_v60 _dafny.Map
		_ = _416_v60
		_416_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v0, (_320_v0).Plus(_320_v0))
		var _417_v61 _dafny.Set
		_ = _417_v61
		_417_v61 = _dafny.SetOf((_dafny.Zero).Minus(_320_v0), _320_v0)
		_416_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_417_v61).Cardinality(), _320_v0)
		var _418_v62 _dafny.Set
		_ = _418_v62
		_418_v62 = _dafny.SetOf(!(_321_v1))
		var _419_v63 *C0
		_ = _419_v63
		var _nw66 *C0 = New_C0_()
		_ = _nw66
		_nw66.Ctor__(Companion_Default___.SafeModuloInt(_320_v0, _320_v0), (func() _dafny.Set {
			if _321_v1 {
				return _418_v62
			}
			return _418_v62
		})())
		_419_v63 = _nw66
		var _420_v64 _dafny.CodePoint
		_ = _420_v64
		_420_v64 = _dafny.CodePoint('u')
		(_324_globalState).F0 = (func() _dafny.Int {
			if !_dafny.Companion_Sequence_.Contains(_352_v20, _420_v64) {
				return _dafny.IntOfInt64(177)
			}
			return ((_419_v63).F6()).Minus((_419_v63).F6())
		})()
		if true {
			var _421_v65 bool
			_ = _421_v65
			var _out1 bool
			_ = _out1
			_out1 = Companion_Default___.M0(_321_v1, _324_globalState)
			_421_v65 = _out1
			_321_v1 = _421_v65
			var _422_v66 _dafny.Array
			_ = _422_v66
			var _nw67 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
			_ = _nw67
			_422_v66 = _nw67
			var _423_v67 _dafny.Sequence
			_ = _423_v67
			_423_v67 = _dafny.SeqOf(_422_v66, _422_v66, _422_v66)
			var _pat_let_tv18 = _422_v66
			_ = _pat_let_tv18
			var _424_v68 _dafny.Array
			_ = _424_v68
			var _nwElement0_14 _dafny.Array = _422_v66
			_ = _nwElement0_14
			var _nw68 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(22))
			_ = _nw68
			(_nw68).ArraySet1(_nwElement0_14, 0)
			(_nw68).ArraySet1(_422_v66, 1)
			(_nw68).ArraySet1((_423_v67).Select((Companion_Default___.SafeIndex((_419_v63).F6(), _dafny.IntOfUint32((_423_v67).Cardinality()))).Uint32()).(_dafny.Array), 2)
			(_nw68).ArraySet1(_422_v66, 3)
			(_nw68).ArraySet1(_422_v66, 4)
			(_nw68).ArraySet1(_422_v66, 5)
			(_nw68).ArraySet1((_423_v67).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm4((_419_v63).F6(), _421_v65, _324_globalState), _dafny.IntOfUint32((_423_v67).Cardinality()))).Uint32()).(_dafny.Array), 6)
			(_nw68).ArraySet1(_422_v66, 7)
			(_nw68).ArraySet1(_422_v66, 8)
			(_nw68).ArraySet1(_422_v66, 9)
			(_nw68).ArraySet1((_423_v67).Select((Companion_Default___.SafeIndex((_419_v63).F6(), _dafny.IntOfUint32((_423_v67).Cardinality()))).Uint32()).(_dafny.Array), 10)
			(_nw68).ArraySet1(_422_v66, 11)
			(_nw68).ArraySet1(_422_v66, 12)
			(_nw68).ArraySet1(_422_v66, 13)
			(_nw68).ArraySet1(_422_v66, 14)
			(_nw68).ArraySet1(_422_v66, 15)
			(_nw68).ArraySet1(_422_v66, 16)
			(_nw68).ArraySet1((func(_pat_let3_0 D2) D2 {
				return func(_425_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let4_0 _dafny.Array) D2 {
						return func(_426_dt__update_hcf11_h0 _dafny.Array) D2 {
							return Companion_D2_.Create_DC5_(_426_dt__update_hcf11_h0)
						}(_pat_let4_0)
					}(_pat_let_tv18)
				}(_pat_let3_0)
			}(Companion_D2_.Create_DC5_(_422_v66))).Dtor_cf11(), 17)
			(_nw68).ArraySet1(_422_v66, 18)
			(_nw68).ArraySet1(_422_v66, 19)
			(_nw68).ArraySet1(_422_v66, 20)
			(_nw68).ArraySet1(_422_v66, 21)
			_424_v68 = _nw68
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_424_v68), 0))
			_ = _index38
			(_424_v68).ArraySet1(_422_v66, (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_424_v68), 0))
			_ = _index39
			(_424_v68).ArraySet1(_422_v66, (_index39).Int())
			var _427_v69 _dafny.Array
			_ = _427_v69
			var _nwElement0_15 bool = _421_v65
			_ = _nwElement0_15
			var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(28))
			_ = _nw69
			(_nw69).ArraySet1(_nwElement0_15, 0)
			(_nw69).ArraySet1(Companion_Default___.Fm0(_321_v1, (Companion_D1_.Create_DC3_((_383_v43).Dtor_cf6(), _421_v65, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(997))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}((func(_428_v64 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_429_i3 _dafny.Int) _dafny.CodePoint {
					return _428_v64
				}
			})(_420_v64))), _421_v65, _dafny.IntOfInt64(639))).Dtor_cf5(), (_419_v63).F6(), _324_globalState), 1)
			(_nw69).ArraySet1(_321_v1, 2)
			(_nw69).ArraySet1(_321_v1, 3)
			(_nw69).ArraySet1(_321_v1, 4)
			(_nw69).ArraySet1(_321_v1, 5)
			(_nw69).ArraySet1(false, 6)
			(_nw69).ArraySet1((_383_v43).Dtor_cf5(), 7)
			(_nw69).ArraySet1(_321_v1, 8)
			(_nw69).ArraySet1(_321_v1, 9)
			(_nw69).ArraySet1(_321_v1, 10)
			(_nw69).ArraySet1(!(_421_v65), 11)
			(_nw69).ArraySet1(_421_v65, 12)
			(_nw69).ArraySet1(_421_v65, 13)
			(_nw69).ArraySet1(_321_v1, 14)
			(_nw69).ArraySet1(_421_v65, 15)
			(_nw69).ArraySet1(_321_v1, 16)
			(_nw69).ArraySet1(_421_v65, 17)
			(_nw69).ArraySet1(true, 18)
			(_nw69).ArraySet1(_421_v65, 19)
			(_nw69).ArraySet1(_321_v1, 20)
			(_nw69).ArraySet1(_421_v65, 21)
			(_nw69).ArraySet1(_421_v65, 22)
			(_nw69).ArraySet1(true, 23)
			(_nw69).ArraySet1(false, 24)
			(_nw69).ArraySet1(_421_v65, 25)
			(_nw69).ArraySet1(_421_v65, 26)
			(_nw69).ArraySet1(_321_v1, 27)
			_427_v69 = _nw69
			var _430_v70 _dafny.Map
			_ = _430_v70
			_430_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_427_v69, _421_v65)
			var _431_v71 *C0
			_ = _431_v71
			var _nw70 *C0 = New_C0_()
			_ = _nw70
			_nw70.Ctor__(((func() _dafny.Int {
				if (_405_v49).Contains(_421_v65) {
					return (_405_v49).Multiplicity(_421_v65)
				}
				return (_419_v63).F6()
			})()).Times((_430_v70).Cardinality()), _418_v62)
			_431_v71 = _nw70
			var _432_v72 _dafny.Map
			_ = _432_v72
			_432_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_353_v21).Cardinality())
			(_324_globalState).F4 = (_432_v72).Cardinality()
		} else {
			var _433_v73 *C0
			_ = _433_v73
			var _nw71 *C0 = New_C0_()
			_ = _nw71
			_nw71.Ctor__(_dafny.IntOfInt64(483), (_419_v63).F7())
			_433_v73 = _nw71
			var _434_v74 _dafny.Array
			_ = _434_v74
			var _nw72 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
			_ = _nw72
			_434_v74 = _nw72
			var _435_v75 _dafny.Map
			_ = _435_v75
			_435_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_419_v63, _434_v74)
			var _436_v76 _dafny.Map
			_ = _436_v76
			_436_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v0, _434_v74)
			_435_v75 = (_435_v75).Update(_433_v73, (func() _dafny.Array {
				if (_436_v76).Contains((_419_v63).F6()) {
					return (_436_v76).Get((_419_v63).F6()).(_dafny.Array)
				}
				return (func() _dafny.Array {
					if (_436_v76).Contains(_320_v0) {
						return (_436_v76).Get(_320_v0).(_dafny.Array)
					}
					return _434_v74
				})()
			})())
			var _437_v77 _dafny.Array
			_ = _437_v77
			var _len0_7 _dafny.Int = _dafny.One
			_ = _len0_7
			var _nw73 _dafny.Array
			_ = _nw73
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw73 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Int = (func(_438_v63 *C0) func(_dafny.Int) _dafny.Int {
					return func(_439_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_439_i4, (_438_v63).F6())
					}
				})(_419_v63)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw73 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw73).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw73).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_437_v77 = _nw73
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_437_v77), 0))
			_ = _index40
			(_437_v77).ArraySet1((_419_v63).F6(), (_index40).Int())
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_437_v77), 0))
			_ = _index41
			(_437_v77).ArraySet1(_320_v0, (_index41).Int())
			var _440_v78 _dafny.Map
			_ = _440_v78
			_440_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_417_v61).Union(_dafny.SetOf(_dafny.IntOfInt64(858))), _dafny.Companion_Sequence_.Update(_352_v20, (Companion_Default___.SafeIndex((_433_v73).F6(), _dafny.IntOfUint32((_352_v20).Cardinality()))).Uint32(), _420_v64))
			var _pat_let_tv19 = _321_v1
			_ = _pat_let_tv19
			var _pat_let_tv20 = _321_v1
			_ = _pat_let_tv20
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_437_v77), 0))
			_ = _index42
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_437_v77), 0))
			_ = _index43
			var _rhs37 _dafny.Sequence = (func() _dafny.Sequence {
				if (_440_v78).Contains((_417_v61).Union(_417_v61)) {
					return (_440_v78).Get((_417_v61).Union(_417_v61)).(_dafny.Sequence)
				}
				return (func(_pat_let5_0 D1) D1 {
					return func(_441_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let6_0 bool) D1 {
							return func(_442_dt__update_hcf6_h0 bool) D1 {
								return func(_pat_let7_0 bool) D1 {
									return func(_443_dt__update_hcf8_h0 bool) D1 {
										return Companion_D1_.Create_DC3_((_441_dt__update__tmp_h1).Dtor_cf5(), _442_dt__update_hcf6_h0, (_441_dt__update__tmp_h1).Dtor_cf7(), _443_dt__update_hcf8_h0, (_441_dt__update__tmp_h1).Dtor_cf9())
									}(_pat_let7_0)
								}(_pat_let_tv20)
							}(_pat_let6_0)
						}(_pat_let_tv19)
					}(_pat_let5_0)
				}(_383_v43)).Dtor_cf7()
			})()
			_ = _rhs37
			var _rhs38 _dafny.Int = (_419_v63).F6()
			_ = _rhs38
			var _rhs39 _dafny.Int = Companion_Default___.Fm4(_320_v0, _dafny.Companion_Sequence_.Equal(_352_v20, _352_v20), _324_globalState)
			_ = _rhs39
			var _rhs40 _dafny.Int = ((_419_v63).F6()).Plus(_320_v0)
			_ = _rhs40
			var _rhs41 _dafny.Int = (_419_v63).F6()
			_ = _rhs41
			var _lhs26 _dafny.Array = _437_v77
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_437_v77), 0))
			_ = _lhs27
			var _lhs28 *GlobalState = _324_globalState
			_ = _lhs28
			var _lhs29 _dafny.Array = _437_v77
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_437_v77), 0))
			_ = _lhs30
			var _lhs31 *GlobalState = _324_globalState
			_ = _lhs31
			_352_v20 = _rhs37
			(_lhs26).ArraySet1(_rhs38, (_lhs27).Int())
			_lhs28.F4 = _rhs39
			(_lhs29).ArraySet1(_rhs40, (_lhs30).Int())
			_lhs31.F4 = _rhs41
			var _444_v79 _dafny.Set
			_ = _444_v79
			_444_v79 = _dafny.SetOf(_420_v64)
			(_324_globalState).F4 = ((_437_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_437_v77), 0))).Int()).(_dafny.Int)).Times(((_444_v79).Intersection(_444_v79)).Cardinality())
			var _445_v80 _dafny.Sequence
			_ = _445_v80
			_445_v80 = _dafny.SeqOf((_419_v63).F7(), _dafny.SetOf(_321_v1, _321_v1, _321_v1))
			var _446_v81 _dafny.Map
			_ = _446_v81
			_446_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_321_v1, _320_v0)
			var _447_v82 _dafny.Sequence
			_ = _447_v82
			_447_v82 = _dafny.SeqOf(_dafny.IntOfInt64(-799))
			var _448_v83 *C0
			_ = _448_v83
			var _nw74 *C0 = New_C0_()
			_ = _nw74
			_nw74.Ctor__(_dafny.IntOfInt64(-604), (_445_v80).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((func() _dafny.Int {
				if (_446_v81).Contains(_321_v1) {
					return (_446_v81).Get(_321_v1).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_447_v82).Cardinality())
			})()), _dafny.IntOfUint32((_445_v80).Cardinality()))).Uint32()).(_dafny.Set))
			_448_v83 = _nw74
		}
	}
	var _449_v84 _dafny.Map
	_ = _449_v84
	_449_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_321_v1, _dafny.IntOfUint32((_352_v20).Cardinality()))
	var _450_v85 _dafny.Map
	_ = _450_v85
	_450_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(492), _449_v84)
	var _451_v86 _dafny.Map
	_ = _451_v86
	_451_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_450_v85).Cardinality(), _449_v84)
	var _452_v87 _dafny.MultiSet
	_ = _452_v87
	_452_v87 = _dafny.MultiSetOf(_321_v1, _321_v1, false, ((_450_v85).Cardinality()).Cmp((_dafny.Zero).Minus(_320_v0)) != 0, !(_321_v1) || (_321_v1))
	_452_v87 = (_452_v87).Difference(_452_v87)
	_321_v1 = _321_v1
	var _453_v88 _dafny.Array
	_ = _453_v88
	var _nw75 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
	_ = _nw75
	_453_v88 = _nw75
	var _454_v89 _dafny.Set
	_ = _454_v89
	_454_v89 = _dafny.SetOf(_321_v1)
	var _rhs42 _dafny.Int = Companion_Default___.SafeModuloInt((_454_v89).Cardinality(), (_320_v0).Times(_dafny.IntOfInt64(578)))
	_ = _rhs42
	var _rhs43 bool = false
	_ = _rhs43
	var _rhs44 _dafny.Array = _453_v88
	_ = _rhs44
	var _rhs45 bool = !(!(_321_v1))
	_ = _rhs45
	var _rhs46 bool = _321_v1
	_ = _rhs46
	var _lhs32 *GlobalState = _324_globalState
	_ = _lhs32
	_lhs32.F0 = _rhs42
	_321_v1 = _rhs43
	_453_v88 = _rhs44
	_321_v1 = _rhs45
	_321_v1 = _rhs46
	var _455_v90 bool
	_ = _455_v90
	var _out2 bool
	_ = _out2
	_out2 = Companion_Default___.M0(_321_v1, _324_globalState)
	_455_v90 = _out2
	var _456_v91 _dafny.Map
	_ = _456_v91
	_456_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v0, true)
	var _457_v92 _dafny.Map
	_ = _457_v92
	_457_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _320_v0)
	var _458_v93 D1
	_ = _458_v93
	_458_v93 = Companion_D1_.Create_DC2_(_320_v0)
	var _459_v94 _dafny.Array
	_ = _459_v94
	var _nwElement0_16 _dafny.MultiSet = _452_v87
	_ = _nwElement0_16
	var _nw76 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(8))
	_ = _nw76
	(_nw76).ArraySet1(_nwElement0_16, 0)
	(_nw76).ArraySet1((_452_v87).Difference(_452_v87), 1)
	(_nw76).ArraySet1(_452_v87, 2)
	(_nw76).ArraySet1(_dafny.MultiSetOf(Companion_Default___.Fm0(!(_321_v1), (func() bool {
		if (_456_v91).Contains((_383_v43).Dtor_cf9()) {
			return (_456_v91).Get((_383_v43).Dtor_cf9()).(bool)
		}
		return _455_v90
	})(), (_457_v92).Cardinality(), _324_globalState), _321_v1, _321_v1, _321_v1, _321_v1), 3)
	(_nw76).ArraySet1((_452_v87).Difference(_452_v87), 4)
	(_nw76).ArraySet1((_452_v87).Update(_321_v1, Companion_Default___.Abs((_458_v93).Dtor_cf4())), 5)
	(_nw76).ArraySet1(_dafny.MultiSetOf(_455_v90, _455_v90), 6)
	(_nw76).ArraySet1(_452_v87, 7)
	_459_v94 = _nw76
	for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_459_v94), 0))); ; {
		_guard_loop_0, _ok22 := _iter22()
		if !_ok22 {
			break
		}
		var _460_i5 _dafny.Int
		_460_i5 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_460_i5).Sign() != -1) && ((_460_i5).Cmp(_dafny.ArrayLen((_459_v94), 0)) < 0)) {
			(_459_v94).ArraySet1(_452_v87, (_460_i5).Int())
		}
	}
	_352_v20 = _dafny.Companion_Sequence_.Concatenate(_352_v20, _352_v20)
	var _source10 D1 = _458_v93
	_ = _source10
	if _source10.Is_DC3() {
		var _461___mcc_h15 bool = _source10.Get_().(D1_DC3).Cf5
		_ = _461___mcc_h15
		var _462___mcc_h16 bool = _source10.Get_().(D1_DC3).Cf6
		_ = _462___mcc_h16
		var _463___mcc_h17 _dafny.Sequence = _source10.Get_().(D1_DC3).Cf7
		_ = _463___mcc_h17
		var _464___mcc_h18 bool = _source10.Get_().(D1_DC3).Cf8
		_ = _464___mcc_h18
		var _465___mcc_h19 _dafny.Int = _source10.Get_().(D1_DC3).Cf9
		_ = _465___mcc_h19
		var _466_cf9 _dafny.Int = _465___mcc_h19
		_ = _466_cf9
		var _467_cf8 bool = _464___mcc_h18
		_ = _467_cf8
		var _468_cf7 _dafny.Sequence = _463___mcc_h17
		_ = _468_cf7
		var _469_cf6 bool = _462___mcc_h16
		_ = _469_cf6
		var _470_cf5 bool = _461___mcc_h15
		_ = _470_cf5
		var _471_v95 _dafny.Map
		_ = _471_v95
		_471_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_469_cf6, false)
		_471_v95 = (_471_v95).Update(!(func() _dafny.Map {
			var _coll22 = _dafny.NewMapBuilder()
			_ = _coll22
			for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(954), _dafny.IntOfInt64(-549))); ; {
				_compr_22, _ok23 := _iter23()
				if !_ok23 {
					break
				}
				var _472_v96 _dafny.Int
				_472_v96 = interface{}(_compr_22).(_dafny.Int)
				if ((_dafny.IntOfInt64(954)).Cmp(_472_v96) <= 0) && ((_472_v96).Cmp(_dafny.IntOfInt64(-549)) < 0) {
					_coll22.Add(Companion_Default___.SafeModuloInt(_472_v96, (_471_v95).Cardinality()), _470_cf5)
				}
			}
			return _coll22.ToMap()
		}()).Equals(_456_v91), _470_cf5)
		var _473_v97 _dafny.CodePoint
		_ = _473_v97
		_473_v97 = _dafny.CodePoint('r')
		_473_v97 = _473_v97
		(_324_globalState).F4 = (_320_v0).Plus(_dafny.IntOfUint32((_468_cf7).Cardinality()))
		_469_cf6 = (func() bool {
			if (_471_v95).Contains((_467_cf8) == (_467_cf8)) {
				return (_471_v95).Get((_467_cf8) == (_467_cf8)).(bool)
			}
			return true
		})()
	} else if _source10.Is_DC2() {
		var _474___mcc_h20 _dafny.Int = _source10.Get_().(D1_DC2).Cf4
		_ = _474___mcc_h20
		var _475_cf4 _dafny.Int = _474___mcc_h20
		_ = _475_cf4
		_454_v89 = Companion_Default___.Fm5(_324_globalState)
		var _476_v98 _dafny.Sequence
		_ = _476_v98
		_476_v98 = _dafny.SeqOf(_dafny.IntOfUint32((_322_v2).Cardinality()))
		if !(((_476_v98).Select((Companion_Default___.SafeIndex(_320_v0, _dafny.IntOfUint32((_476_v98).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_475_cf4) != 0) {
			var _477_v99 _dafny.Map
			_ = _477_v99
			_477_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-19))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}(func(_478_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})), _322_v2)
			var _479_v101 _dafny.CodePoint
			_ = _479_v101
			_479_v101 = _dafny.CodePoint('d')
			var _480_v102 _dafny.Map
			_ = _480_v102
			_480_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v0, _479_v101)
			var _481_v103 _dafny.Sequence
			_ = _481_v103
			_481_v103 = _dafny.SeqOf(_477_v99, _477_v99, _477_v99, func() _dafny.Map {
				var _coll23 = _dafny.NewMapBuilder()
				_ = _coll23
				for _iter24 := _dafny.Iterate((_dafny.MultiSetOf(_352_v20, _352_v20)).Elements()); ; {
					_compr_23, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _482_v100 _dafny.Sequence
					_482_v100 = interface{}(_compr_23).(_dafny.Sequence)
					if (_dafny.MultiSetOf(_352_v20, _352_v20)).Contains(_482_v100) {
						_coll23.Add(_482_v100, _dafny.SeqOf(true))
					}
				}
				return _coll23.ToMap()
			}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_352_v20, _dafny.Companion_Sequence_.Update(_322_v2, (Companion_Default___.SafeIndex((_480_v102).Cardinality(), _dafny.IntOfUint32((_322_v2).Cardinality()))).Uint32(), _321_v1)))
			var _483_v105 _dafny.Map
			_ = _483_v105
			_483_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v0, (func() _dafny.Map {
				if _455_v90 {
					return ((_481_v103).Select((Companion_Default___.SafeIndex(_320_v0, _dafny.IntOfUint32((_481_v103).Cardinality()))).Uint32()).(_dafny.Map)).Update(_352_v20, _322_v2)
				}
				return func() _dafny.Map {
					var _coll24 = _dafny.NewMapBuilder()
					_ = _coll24
					for _iter25 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_352_v20, _321_v1)).Keys().Elements()); ; {
						_compr_24, _ok25 := _iter25()
						if !_ok25 {
							break
						}
						var _484_v104 _dafny.Sequence
						_484_v104 = interface{}(_compr_24).(_dafny.Sequence)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_352_v20, _321_v1)).Contains(_484_v104) {
							_coll24.Add(_484_v104, _322_v2)
						}
					}
					return _coll24.ToMap()
				}()
			})())
			_483_v105 = (_483_v105).Update(_475_cf4, ((_477_v99).Update(_dafny.UnicodeSeqOfUtf8Bytes("kb"), _322_v2)).Merge(_477_v99))
			_321_v1 = !(_321_v1) || (_321_v1)
			var _485_v106 _dafny.Sequence
			_ = _485_v106
			_485_v106 = _dafny.SeqOf(_449_v84)
			var _486_v107 _dafny.Sequence
			_ = _486_v107
			_486_v107 = _dafny.SeqOf((_485_v106).Select((Companion_Default___.SafeIndex(_320_v0, _dafny.IntOfUint32((_485_v106).Cardinality()))).Uint32()).(_dafny.Map), _449_v84, _449_v84, _449_v84)
			var _487_v108 *C0
			_ = _487_v108
			var _nw77 *C0 = New_C0_()
			_ = _nw77
			_nw77.Ctor__((((_486_v107).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_352_v20).Cardinality()), _dafny.IntOfUint32((_486_v107).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_449_v84)).Cardinality(), _454_v89)
			_487_v108 = _nw77
			(_324_globalState).F4 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_352_v20).Cardinality()), _320_v0)
			var _488_v109 D2
			_ = _488_v109
			_488_v109 = Companion_D2_.Create_DC7_((_475_cf4).Times((_487_v108).F6()))
			var _pat_let_tv21 = _487_v108
			_ = _pat_let_tv21
			_488_v109 = func(_pat_let8_0 D2) D2 {
				return func(_489_dt__update__tmp_h2 D2) D2 {
					return func(_pat_let9_0 _dafny.Int) D2 {
						return func(_490_dt__update_hcf15_h0 _dafny.Int) D2 {
							return Companion_D2_.Create_DC7_(_490_dt__update_hcf15_h0)
						}(_pat_let9_0)
					}((_pat_let_tv21).F6())
				}(_pat_let8_0)
			}(_488_v109)
		} else {
			var _491_v110 _dafny.Set
			_ = _491_v110
			_491_v110 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-884))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}(func(_492_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			}))).Cardinality()), _320_v0)
			var _493_v111 _dafny.Map
			_ = _493_v111
			_493_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v0, (_491_v110).Union(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_455_v90)).Cardinality()), _320_v0)))
			_493_v111 = (_493_v111).Update(_320_v0, _491_v110)
			var _494_v112 _dafny.Array
			_ = _494_v112
			var _nwElement0_17 _dafny.Int = _320_v0
			_ = _nwElement0_17
			var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(2))
			_ = _nw78
			(_nw78).ArraySet1(_nwElement0_17, 0)
			(_nw78).ArraySet1(_475_cf4, 1)
			_494_v112 = _nw78
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_494_v112), 0))
			_ = _index44
			(_494_v112).ArraySet1((_dafny.IntOfInt64(301)).Plus(_dafny.IntOfUint32((_352_v20).Cardinality())), (_index44).Int())
			var _495_v113 _dafny.Sequence
			_ = _495_v113
			_495_v113 = _dafny.SeqOf(Companion_D2_.Create_DC5_(_494_v112))
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_494_v112), 0))
			_ = _index45
			var _rhs47 _dafny.Int = (_dafny.Zero).Minus((Companion_D2_.Create_DC6_(_320_v0, _321_v1, _494_v112)).Dtor_cf12())
			_ = _rhs47
			var _rhs48 _dafny.Sequence = _dafny.SeqOf(Companion_D2_.Create_DC5_(_494_v112))
			_ = _rhs48
			var _rhs49 _dafny.Int = Companion_Default___.SafeDivisionInt((_454_v89).Cardinality(), _320_v0)
			_ = _rhs49
			var _lhs33 _dafny.Array = _494_v112
			_ = _lhs33
			var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_494_v112), 0))
			_ = _lhs34
			(_lhs33).ArraySet1(_rhs47, (_lhs34).Int())
			_495_v113 = _rhs48
			_320_v0 = _rhs49
			var _496_v114 *C0
			_ = _496_v114
			var _nw79 *C0 = New_C0_()
			_ = _nw79
			_nw79.Ctor__((_320_v0).Minus(_475_cf4), Companion_Default___.Fm5(_324_globalState))
			_496_v114 = _nw79
			var _497_v115 bool
			_ = _497_v115
			var _out3 bool
			_ = _out3
			_out3 = Companion_Default___.M0(_455_v90, _324_globalState)
			_497_v115 = _out3
			var _498_v116 _dafny.Map
			_ = _498_v116
			_498_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_496_v114, _321_v1)
			_497_v115 = !((func() bool {
				if (_498_v116).Contains((func() *C0 {
					if false {
						return _496_v114
					}
					return _496_v114
				})()) {
					return (_498_v116).Get((func() *C0 {
						if false {
							return _496_v114
						}
						return _496_v114
					})()).(bool)
				}
				return _497_v115
			})())
		}
		var _499_v117 _dafny.Array
		_ = _499_v117
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_8
		var _nw80 _dafny.Array
		_ = _nw80
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw80 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) _dafny.Int = (func(_500_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_501_i8 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_501_i8, _500_v0)
				}
			})(_320_v0)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw80 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw80).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw80).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_499_v117 = _nw80
		var _502_v118 D2
		_ = _502_v118
		_502_v118 = Companion_D2_.Create_DC5_(_499_v117)
		var _503_v119 _dafny.Map
		_ = _503_v119
		_503_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_499_v117, (func() D2 {
			if _455_v90 {
				return _502_v118
			}
			return _502_v118
		})())
		_503_v119 = (_503_v119).Update(_499_v117, _502_v118)
		(_324_globalState).F0 = Companion_Default___.SafeModuloInt(_475_cf4, (_dafny.Zero).Minus(_320_v0))
	} else {
		var _504___mcc_h21 D1 = _source10.Get_().(D1_DC4).Cf10
		_ = _504___mcc_h21
		var _505_cf10 D1 = _504___mcc_h21
		_ = _505_cf10
		var _506_v120 _dafny.CodePoint
		_ = _506_v120
		_506_v120 = _dafny.CodePoint('f')
		var _507_v121 _dafny.MultiSet
		_ = _507_v121
		_507_v121 = _dafny.MultiSetOf(_506_v120)
		var _508_v122 _dafny.Array
		_ = _508_v122
		var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
		_ = _nw81
		_508_v122 = _nw81
		var _509_v123 _dafny.Map
		_ = _509_v123
		_509_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_507_v121, _508_v122)
		_509_v123 = (_509_v123).Update(_507_v121, _508_v122)
		(_324_globalState).F0 = (_dafny.Zero).Minus(_320_v0)
		_321_v1 = (_320_v0).Cmp(Companion_Default___.SafeModuloInt(_320_v0, _320_v0)) <= 0
		var _510_v124 _dafny.Set
		_ = _510_v124
		_510_v124 = _dafny.SetOf(_320_v0, _320_v0, _320_v0)
		(_324_globalState).F4 = (_dafny.Zero).Minus(((_510_v124).Intersection(_510_v124)).Cardinality())
	}
	var _511_v125 _dafny.Sequence
	_ = _511_v125
	_511_v125 = _dafny.SeqOf(_320_v0, _320_v0)
	var _512_v126 _dafny.Array
	_ = _512_v126
	var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
	_ = _nw82
	_512_v126 = _nw82
	var _513_v127 D2
	_ = _513_v127
	_513_v127 = Companion_D2_.Create_DC6_(_dafny.IntOfUint32((Companion_Default___.Fm3(_322_v2, _324_globalState)).Cardinality()), _dafny.Companion_Sequence_.Equal(_511_v125, _511_v125), _512_v126)
	var _source11 D2 = _513_v127
	_ = _source11
	if _source11.Is_DC6() {
		var _514___mcc_h22 _dafny.Int = _source11.Get_().(D2_DC6).Cf12
		_ = _514___mcc_h22
		var _515___mcc_h23 bool = _source11.Get_().(D2_DC6).Cf13
		_ = _515___mcc_h23
		var _516___mcc_h24 _dafny.Array = _source11.Get_().(D2_DC6).Cf14
		_ = _516___mcc_h24
		var _517_cf14 _dafny.Array = _516___mcc_h24
		_ = _517_cf14
		var _518_cf13 bool = _515___mcc_h23
		_ = _518_cf13
		var _519_cf12 _dafny.Int = _514___mcc_h22
		_ = _519_cf12
		_320_v0 = ((_dafny.IntOfInt64(761)).Times(_519_cf12)).Minus(_519_cf12)
		_321_v1 = _455_v90
		var _520_v128 _dafny.Map
		_ = _520_v128
		_520_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm4(_519_cf12, _518_cf13, _324_globalState), Companion_Default___.SafeModuloInt(_320_v0, _dafny.IntOfUint32((_511_v125).Cardinality())))
		var _rhs50 bool = _455_v90
		_ = _rhs50
		var _rhs51 bool = false
		_ = _rhs51
		var _rhs52 _dafny.Map = (_520_v128).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_519_cf12, _519_cf12))
		_ = _rhs52
		_455_v90 = _rhs50
		_455_v90 = _rhs51
		_520_v128 = _rhs52
		_455_v90 = _321_v1
	} else if _source11.Is_DC7() {
		var _521___mcc_h25 _dafny.Int = _source11.Get_().(D2_DC7).Cf15
		_ = _521___mcc_h25
		var _522_cf15 _dafny.Int = _521___mcc_h25
		_ = _522_cf15
		var _523_v129 bool
		_ = _523_v129
		var _out4 bool
		_ = _out4
		_out4 = Companion_Default___.M0((func() bool {
			if _321_v1 {
				return (func() bool {
					if (_456_v91).Contains(_522_cf15) {
						return (_456_v91).Get(_522_cf15).(bool)
					}
					return !(_455_v90)
				})()
			}
			return _321_v1
		})(), _324_globalState)
		_523_v129 = _out4
		var _524_v130 *C0
		_ = _524_v130
		var _nw83 *C0 = New_C0_()
		_ = _nw83
		_nw83.Ctor__(_320_v0, _454_v89)
		_524_v130 = _nw83
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_453_v88), 0))
		_ = _index46
		(_453_v88).ArraySet1(_524_v130, (_index46).Int())
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_453_v88), 0))
		_ = _index47
		var _rhs53 *C0 = _524_v130
		_ = _rhs53
		var _rhs54 _dafny.Int = _320_v0
		_ = _rhs54
		var _rhs55 _dafny.Sequence = _511_v125
		_ = _rhs55
		var _rhs56 *C0 = _524_v130
		_ = _rhs56
		var _lhs35 _dafny.Array = _453_v88
		_ = _lhs35
		var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_453_v88), 0))
		_ = _lhs36
		var _lhs37 *GlobalState = _324_globalState
		_ = _lhs37
		(_lhs35).ArraySet1(_rhs53, (_lhs36).Int())
		_lhs37.F0 = _rhs54
		_511_v125 = _rhs55
		_524_v130 = _rhs56
		var _525_v131 _dafny.CodePoint
		_ = _525_v131
		_525_v131 = _dafny.CodePoint('a')
		_525_v131 = _525_v131
		var _526_v132 D2
		_ = _526_v132
		_526_v132 = Companion_D2_.Create_DC7_(_522_cf15)
		var _527_v133 _dafny.Set
		_ = _527_v133
		_527_v133 = _dafny.SetOf(_522_cf15)
		var _528_v134 _dafny.Array
		_ = _528_v134
		var _nwElement0_18 D2 = _526_v132
		_ = _nwElement0_18
		var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(28))
		_ = _nw84
		(_nw84).ArraySet1(_nwElement0_18, 0)
		(_nw84).ArraySet1(_526_v132, 1)
		(_nw84).ArraySet1(_526_v132, 2)
		(_nw84).ArraySet1((func() D2 {
			if Companion_Default___.Fm0(_455_v90, _321_v1, (_527_v133).Cardinality(), _324_globalState) {
				return _526_v132
			}
			return Companion_D2_.Create_DC7_(_dafny.IntOfUint32((_352_v20).Cardinality()))
		})(), 3)
		(_nw84).ArraySet1(_526_v132, 4)
		(_nw84).ArraySet1(_526_v132, 5)
		(_nw84).ArraySet1(Companion_D2_.Create_DC7_(_dafny.IntOfInt64(68)), 6)
		(_nw84).ArraySet1(_526_v132, 7)
		(_nw84).ArraySet1(_526_v132, 8)
		(_nw84).ArraySet1(_526_v132, 9)
		(_nw84).ArraySet1(Companion_D2_.Create_DC7_(_320_v0), 10)
		(_nw84).ArraySet1(_526_v132, 11)
		(_nw84).ArraySet1(_526_v132, 12)
		(_nw84).ArraySet1(_526_v132, 13)
		(_nw84).ArraySet1((func() D2 {
			if false {
				return _526_v132
			}
			return _526_v132
		})(), 14)
		(_nw84).ArraySet1(_526_v132, 15)
		(_nw84).ArraySet1(_526_v132, 16)
		(_nw84).ArraySet1(_526_v132, 17)
		(_nw84).ArraySet1(_526_v132, 18)
		(_nw84).ArraySet1(Companion_D2_.Create_DC7_(_dafny.IntOfUint32((_322_v2).Cardinality())), 19)
		(_nw84).ArraySet1(_526_v132, 20)
		(_nw84).ArraySet1(_526_v132, 21)
		(_nw84).ArraySet1(_526_v132, 22)
		(_nw84).ArraySet1(_526_v132, 23)
		(_nw84).ArraySet1((func() D2 {
			if _455_v90 {
				return _526_v132
			}
			return _526_v132
		})(), 24)
		(_nw84).ArraySet1(_526_v132, 25)
		(_nw84).ArraySet1(_526_v132, 26)
		(_nw84).ArraySet1((func() D2 {
			if _321_v1 {
				return _526_v132
			}
			return _526_v132
		})(), 27)
		_528_v134 = _nw84
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_528_v134), 0))
		_ = _index48
		(_528_v134).ArraySet1(_526_v132, (_index48).Int())
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_528_v134), 0))
		_ = _index49
		(_528_v134).ArraySet1(Companion_Default___.Fm6(_324_globalState), (_index49).Int())
	} else if _source11.Is_DC5() {
		var _529___mcc_h26 _dafny.Array = _source11.Get_().(D2_DC5).Cf11
		_ = _529___mcc_h26
		var _530_cf11 _dafny.Array = _529___mcc_h26
		_ = _530_cf11
		if (func() bool {
			if _321_v1 {
				return _455_v90
			}
			return _321_v1
		})() {
			_455_v90 = ((_dafny.Zero).Minus(_320_v0)).Cmp((_dafny.Zero).Minus(Companion_Default___.Fm4(_dafny.IntOfInt64(-120), _321_v1, _324_globalState))) >= 0
			_452_v87 = (_dafny.MultiSetOf(_321_v1, _455_v90, _321_v1, _321_v1, _455_v90)).Update(!(!(_455_v90)), Companion_Default___.Abs(_320_v0))
			(_324_globalState).F4 = (_511_v125).Select((Companion_Default___.SafeIndex(_320_v0, _dafny.IntOfUint32((_511_v125).Cardinality()))).Uint32()).(_dafny.Int)
			var _531_v135 _dafny.Array
			_ = _531_v135
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_9
			var _nw85 _dafny.Array
			_ = _nw85
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw85 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Sequence = (func(_532_v20 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_533_i9 _dafny.Int) _dafny.Sequence {
						return _532_v20
					}
				})(_352_v20)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw85 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw85).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw85).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_531_v135 = _nw85
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_531_v135), 0))
			_ = _index50
			(_531_v135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_352_v20, _dafny.UnicodeSeqOfUtf8Bytes("lkypk")), (_index50).Int())
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_531_v135), 0))
			_ = _index51
			(_531_v135).ArraySet1(_352_v20, (_index51).Int())
			_449_v84 = (_449_v84).Update(true, _320_v0)
		} else {
			var _534_v136 bool
			_ = _534_v136
			var _out5 bool
			_ = _out5
			_out5 = Companion_Default___.M0(((_dafny.Zero).Minus(_320_v0)).Cmp(_dafny.IntOfUint32((_511_v125).Cardinality())) != 0, _324_globalState)
			_534_v136 = _out5
			var _535_v137 bool
			_ = _535_v137
			var _out6 bool
			_ = _out6
			_out6 = Companion_Default___.M0(_321_v1, _324_globalState)
			_535_v137 = _out6
			_534_v136 = !_dafny.Companion_Sequence_.Equal(_352_v20, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qwjyx"), _352_v20))
			(_324_globalState).F0 = _320_v0
			var _536_v138 *C0
			_ = _536_v138
			var _nw86 *C0 = New_C0_()
			_ = _nw86
			_nw86.Ctor__(_320_v0, _454_v89)
			_536_v138 = _nw86
		}
		(_324_globalState).F4 = _dafny.IntOfInt64(-260)
		var _537_v139 _dafny.Array
		_ = _537_v139
		var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
		_ = _nw87
		_537_v139 = _nw87
		_537_v139 = _537_v139
		var _538_v140 _dafny.Map
		_ = _538_v140
		_538_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v0, _320_v0)
		(_324_globalState).F0 = (func() _dafny.Int {
			if (_538_v140).Contains((_dafny.Zero).Minus(_320_v0)) {
				return (_538_v140).Get((_dafny.Zero).Minus(_320_v0)).(_dafny.Int)
			}
			return _320_v0
		})()
	} else {
		var _539___mcc_h27 D2 = _source11.Get_().(D2_DC8).Cf16
		_ = _539___mcc_h27
		var _540_cf16 D2 = _539___mcc_h27
		_ = _540_cf16
		var _541_v141 _dafny.Map
		_ = _541_v141
		_541_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_455_v90, _321_v1)
		var _542_v142 *C0
		_ = _542_v142
		var _nw88 *C0 = New_C0_()
		_ = _nw88
		_nw88.Ctor__(_320_v0, _454_v89)
		_542_v142 = _nw88
		var _543_v143 _dafny.Map
		_ = _543_v143
		_543_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_542_v142, (_542_v142).F6())
		if ((_dafny.Zero).Minus(((func() _dafny.Int {
			if (_353_v21).Contains(_dafny.IntOfUint32((_322_v2).Cardinality())) {
				return (_353_v21).Multiplicity(_dafny.IntOfUint32((_322_v2).Cardinality()))
			}
			return (_541_v141).Cardinality()
		})()).Times(_320_v0))).Cmp(((_543_v143).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_542_v142, (_542_v142).F6()))).Cardinality()) > 0 {
			var _544_v144 _dafny.MultiSet
			_ = _544_v144
			_544_v144 = _dafny.MultiSetOf(_512_v126)
			var _545_v145 _dafny.Array
			_ = _545_v145
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_10
			var _nw89 _dafny.Array
			_ = _nw89
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw89 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) bool = func(_546_i10 _dafny.Int) bool {
					return true
				}
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw89 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw89).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw89).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_545_v145 = _nw89
			var _547_v146 _dafny.Array
			_ = _547_v146
			var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(16))
			_ = _nw90
			_547_v146 = _nw90
			var _548_v147 D0
			_ = _548_v147
			_548_v147 = Companion_D0_.Create_DC1_(_544_v144, _545_v145, (func() _dafny.Array {
				if _321_v1 {
					return _547_v146
				}
				return _547_v146
			})())
			_548_v147 = _548_v147
			(_324_globalState).F4 = (_542_v142).F6()
			_449_v84 = Companion_Default___.Fm7(_dafny.MultiSetOf((func() bool {
				if (_456_v91).Contains(_320_v0) {
					return (_456_v91).Get(_320_v0).(bool)
				}
				return _455_v90
			})(), false), _352_v20, _324_globalState)
			_321_v1 = _321_v1
			_449_v84 = (_449_v84).Update((_383_v43).Dtor_cf6(), _320_v0)
		} else {
			var _549_v148 D0
			_ = _549_v148
			_549_v148 = Companion_D0_.Create_DC0_(false)
			var _550_v149 *C0
			_ = _550_v149
			var _nw91 *C0 = New_C0_()
			_ = _nw91
			_nw91.Ctor__((_542_v142).Fm1(_dafny.MultiSetOf((_542_v142).F6()), _549_v148, _324_globalState), _dafny.SetOf(false, _321_v1))
			_550_v149 = _nw91
			var _551_v150 *C0
			_ = _551_v150
			var _nw92 *C0 = New_C0_()
			_ = _nw92
			_nw92.Ctor__((_542_v142).F6(), (_454_v89).Intersection((_550_v149).F7()))
			_551_v150 = _nw92
			var _552_v151 D3
			_ = _552_v151
			_552_v151 = Companion_D3_.Create_DC9_(_511_v125)
			var _553_v152 _dafny.MultiSet
			_ = _553_v152
			_553_v152 = _dafny.MultiSetOf(_dafny.SeqOf(_320_v0), _dafny.Companion_Sequence_.Update((_552_v151).Dtor_cf17(), (Companion_Default___.SafeIndex((_449_v84).Cardinality(), _dafny.IntOfUint32(((_552_v151).Dtor_cf17()).Cardinality()))).Uint32(), (_542_v142).F6()), _511_v125)
			var _554_v153 *C0
			_ = _554_v153
			var _nw93 *C0 = New_C0_()
			_ = _nw93
			_nw93.Ctor__((_dafny.Zero).Minus((func() _dafny.Int {
				if (_553_v152).Contains(_dafny.Companion_Sequence_.Update(_511_v125, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.IntOfUint32((_511_v125).Cardinality()))).Uint32(), _dafny.IntOfUint32((_352_v20).Cardinality()))) {
					return (_553_v152).Multiplicity(_dafny.Companion_Sequence_.Update(_511_v125, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.IntOfUint32((_511_v125).Cardinality()))).Uint32(), _dafny.IntOfUint32((_352_v20).Cardinality())))
				}
				return _dafny.IntOfInt64(624)
			})()), _dafny.SetOf(_455_v90))
			_554_v153 = _nw93
			_455_v90 = _455_v90
			var _555_v154 bool
			_ = _555_v154
			var _out7 bool
			_ = _out7
			_out7 = Companion_Default___.M0(_321_v1, _324_globalState)
			_555_v154 = _out7
		}
		_455_v90 = !(false) || (!((_452_v87).Contains(Companion_Default___.Fm0(false, _321_v1, _320_v0, _324_globalState))))
		var _556_v155 D3
		_ = _556_v155
		_556_v155 = Companion_D3_.Create_DC10_(_321_v1, (Companion_Default___.Fm9(_dafny.IntOfInt64(664), (_353_v21).Cardinality(), _324_globalState)).Cardinality(), _320_v0)
		var _557_v156 D3
		_ = _557_v156
		_557_v156 = Companion_D3_.Create_DC12_(_556_v155)
		_456_v91 = Companion_Default___.Fm8(_321_v1, _557_v156, _454_v89, _320_v0, _324_globalState)
		if (Companion_Default___.SafeDivisionInt((_542_v142).F6(), (_542_v142).F6())).Cmp(((_542_v142).F6()).Plus(_dafny.IntOfUint32((_352_v20).Cardinality()))) != 0 {
			(_324_globalState).F0 = (_542_v142).F6()
			var _558_v157 bool
			_ = _558_v157
			var _out8 bool
			_ = _out8
			_out8 = Companion_Default___.M0(!((_dafny.MultiSetOf(_321_v1, !(true))).IsSubsetOf(_452_v87)), _324_globalState)
			_558_v157 = _out8
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_453_v88), 0))
			_ = _index52
			(_453_v88).ArraySet1(_542_v142, (_index52).Int())
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_453_v88), 0))
			_ = _index53
			var _rhs57 _dafny.MultiSet = (_353_v21).Difference(_353_v21)
			_ = _rhs57
			var _rhs58 _dafny.Int = (func() _dafny.Int {
				if !(_558_v157) || (_558_v157) {
					return (_320_v0).Times((_542_v142).F6())
				}
				return (_541_v141).Cardinality()
			})()
			_ = _rhs58
			var _rhs59 bool = !(_455_v90)
			_ = _rhs59
			var _rhs60 *C0 = _542_v142
			_ = _rhs60
			var _lhs38 *GlobalState = _324_globalState
			_ = _lhs38
			var _lhs39 _dafny.Array = _453_v88
			_ = _lhs39
			var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_453_v88), 0))
			_ = _lhs40
			_353_v21 = _rhs57
			_lhs38.F0 = _rhs58
			_321_v1 = _rhs59
			(_lhs39).ArraySet1(_rhs60, (_lhs40).Int())
			var _559_v158 _dafny.Map
			_ = _559_v158
			_559_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm10((_456_v91).Cardinality(), (_542_v142).F6(), _324_globalState), _dafny.SeqOf(_455_v90)), Companion_D3_.Create_DC12_(_556_v155))
			_559_v158 = Companion_Default___.Fm11(_324_globalState)
			var _560_v159 _dafny.Array
			_ = _560_v159
			var _nw94 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw94
			_560_v159 = _nw94
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_560_v159), 0))
			_ = _index54
			(_560_v159).ArraySet1(_455_v90, (_index54).Int())
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_560_v159), 0))
			_ = _index55
			var _rhs61 bool = _455_v90
			_ = _rhs61
			var _rhs62 _dafny.Sequence = _352_v20
			_ = _rhs62
			var _lhs41 _dafny.Array = _560_v159
			_ = _lhs41
			var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_560_v159), 0))
			_ = _lhs42
			(_lhs41).ArraySet1(_rhs61, (_lhs42).Int())
			_352_v20 = _rhs62
		} else {
			_455_v90 = _455_v90
			_455_v90 = true
			var _561_v160 _dafny.Map
			_ = _561_v160
			_561_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_322_v2, (_542_v142).F6())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_512_v126), 0))
			_ = _index56
			(_512_v126).ArraySet1((_561_v160).Cardinality(), (_index56).Int())
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_512_v126), 0))
			_ = _index57
			(_512_v126).ArraySet1((_542_v142).F6(), (_index57).Int())
			var _562_v161 _dafny.Map
			_ = _562_v161
			_562_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_542_v142, (func() bool {
				if (_456_v91).Contains((_512_v126).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_512_v126), 0))).Int()).(_dafny.Int)) {
					return (_456_v91).Get((_512_v126).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_512_v126), 0))).Int()).(_dafny.Int)).(bool)
				}
				return _321_v1
			})())
			_562_v161 = (_562_v161).Update(_542_v142, _455_v90)
			var _nw95 *C0 = New_C0_()
			_ = _nw95
			_nw95.Ctor__(Companion_Default___.SafeModuloInt((_542_v142).F6(), _320_v0), _454_v89)
			_542_v142 = _nw95
		}
	}
	_321_v1 = (_320_v0).Cmp(_dafny.IntOfInt64(183)) < 0
	for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_512_v126), 0))); ; {
		_guard_loop_1, _ok26 := _iter26()
		if !_ok26 {
			break
		}
		var _563_i11 _dafny.Int
		_563_i11 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_563_i11).Sign() != -1) && ((_563_i11).Cmp(_dafny.ArrayLen((_512_v126), 0)) < 0)) {
			(_512_v126).ArraySet1((_563_i11).Minus((func() _dafny.Int {
				if _455_v90 {
					return _320_v0
				}
				return _320_v0
			})()), (_563_i11).Int())
		}
	}
	var _564_v162 D3
	_ = _564_v162
	_564_v162 = Companion_D3_.Create_DC9_(_511_v125)
	var _source12 D3 = (func() D3 {
		if (func() bool {
			if _455_v90 {
				return _455_v90
			}
			return !(_455_v90)
		})() {
			return Companion_D3_.Create_DC12_(Companion_D3_.Create_DC12_(_564_v162))
		}
		return Companion_Default___.Fm12(_321_v1, (Companion_Default___.Fm13(_455_v90, _455_v90, _324_globalState)).Cardinality(), _455_v90, _dafny.IntOfInt64(49), _324_globalState)
	})()
	_ = _source12
	if _source12.Is_DC10() {
		var _565___mcc_h28 bool = _source12.Get_().(D3_DC10).Cf18
		_ = _565___mcc_h28
		var _566___mcc_h29 _dafny.Int = _source12.Get_().(D3_DC10).Cf19
		_ = _566___mcc_h29
		var _567___mcc_h30 _dafny.Int = _source12.Get_().(D3_DC10).Cf20
		_ = _567___mcc_h30
		var _568_cf20 _dafny.Int = _567___mcc_h30
		_ = _568_cf20
		var _569_cf19 _dafny.Int = _566___mcc_h29
		_ = _569_cf19
		var _570_cf18 bool = _565___mcc_h28
		_ = _570_cf18
		var _571_v163 _dafny.Sequence
		_ = _571_v163
		_571_v163 = _dafny.SeqOf(_449_v84, _449_v84)
		var _572_v164 bool
		_ = _572_v164
		var _out9 bool
		_ = _out9
		_out9 = Companion_Default___.M0(_dafny.Companion_Sequence_.Equal(_571_v163, _571_v163), _324_globalState)
		_572_v164 = _out9
		_352_v20 = _dafny.Companion_Sequence_.Concatenate(_352_v20, _352_v20)
		var _573_v165 _dafny.Sequence
		_ = _573_v165
		_573_v165 = _dafny.SeqOf(_322_v2)
		var _574_v166 *C0
		_ = _574_v166
		var _nw96 *C0 = New_C0_()
		_ = _nw96
		_nw96.Ctor__((_449_v84).Cardinality(), _454_v89)
		_574_v166 = _nw96
		var _575_v167 _dafny.Sequence
		_ = _575_v167
		_575_v167 = _dafny.SeqOf(_574_v166, _574_v166)
		var _576_v168 _dafny.Map
		_ = _576_v168
		_576_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_575_v167).Select((Companion_Default___.SafeIndex(_569_cf19, _dafny.IntOfUint32((_575_v167).Cardinality()))).Uint32()).(*C0), _455_v90)
		var _577_v169 _dafny.Map
		_ = _577_v169
		_577_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_574_v166, !(_572_v164)), _321_v1)
		var _578_v170 _dafny.Array
		_ = _578_v170
		var _nwElement0_19 bool = _570_cf18
		_ = _nwElement0_19
		var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(17))
		_ = _nw97
		(_nw97).ArraySet1(_nwElement0_19, 0)
		(_nw97).ArraySet1(_321_v1, 1)
		(_nw97).ArraySet1(_321_v1, 2)
		(_nw97).ArraySet1(_570_cf18, 3)
		(_nw97).ArraySet1(_dafny.Companion_Sequence_.Equal(_352_v20, _352_v20), 4)
		(_nw97).ArraySet1(!(_dafny.Companion_Sequence_.Contains((_573_v165).Select((Companion_Default___.SafeIndex(_569_cf19, _dafny.IntOfUint32((_573_v165).Cardinality()))).Uint32()).(_dafny.Sequence), _321_v1)), 5)
		(_nw97).ArraySet1(_321_v1, 6)
		(_nw97).ArraySet1((_dafny.IntOfUint32((_511_v125).Cardinality())).Cmp((_449_v84).Cardinality()) != 0, 7)
		(_nw97).ArraySet1(_455_v90, 8)
		(_nw97).ArraySet1((_568_cf20).Cmp(_dafny.IntOfInt64(-329)) != 0, 9)
		(_nw97).ArraySet1(_572_v164, 10)
		(_nw97).ArraySet1(true, 11)
		(_nw97).ArraySet1(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_576_v168, _570_cf18)).Equals(_577_v169), 12)
		(_nw97).ArraySet1((func() bool {
			if _321_v1 {
				return _570_cf18
			}
			return _455_v90
		})(), 13)
		(_nw97).ArraySet1(_572_v164, 14)
		(_nw97).ArraySet1(_570_cf18, 15)
		(_nw97).ArraySet1(_455_v90, 16)
		_578_v170 = _nw97
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_578_v170), 0))
		_ = _index58
		(_578_v170).ArraySet1(_570_cf18, (_index58).Int())
		var _579_v172 _dafny.Set
		_ = _579_v172
		_579_v172 = _dafny.SetOf(_322_v2)
		var _580_v173 _dafny.Set
		_ = _580_v173
		_580_v173 = _dafny.SetOf((_579_v172).Cardinality())
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_578_v170), 0))
		_ = _index59
		(_578_v170).ArraySet1((_dafny.SetOf((_511_v125).Select((Companion_Default___.SafeIndex((_353_v21).Cardinality(), _dafny.IntOfUint32((_511_v125).Cardinality()))).Uint32()).(_dafny.Int), (_574_v166).Fm2(_570_cf18, _511_v125, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(611))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}((func(_581_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_582_i12 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					var _coll25 = _dafny.NewMapBuilder()
					_ = _coll25
					for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(908), _dafny.IntOfInt64(-712))); ; {
						_compr_25, _ok27 := _iter27()
						if !_ok27 {
							break
						}
						var _583_v171 _dafny.Int
						_583_v171 = interface{}(_compr_25).(_dafny.Int)
						if ((_dafny.IntOfInt64(908)).Cmp(_583_v171) <= 0) && ((_583_v171).Cmp(_dafny.IntOfInt64(-712)) < 0) {
							_coll25.Add((_583_v171).Times(_581_v0), false)
						}
					}
					return _coll25.ToMap()
				}()).Cardinality(), true)).Cardinality()
			}
		})(_320_v0))), _449_v84, _324_globalState))).IsDisjointFrom(_580_v173), (_index59).Int())
		var _584_v174 _dafny.CodePoint
		_ = _584_v174
		_584_v174 = _dafny.CodePoint('a')
		_569_cf19 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_352_v20, (Companion_Default___.SafeIndex(((_574_v166).F6()).Times(_320_v0), _dafny.IntOfUint32((_352_v20).Cardinality()))).Uint32(), _584_v174)).Cardinality())
	} else if _source12.Is_DC11() {
		var _585_v175 D3
		_ = _585_v175
		_585_v175 = Companion_D3_.Create_DC9_(_dafny.SeqOf(_320_v0))
		var _pat_let_tv22 = _511_v125
		_ = _pat_let_tv22
		var _586_v176 _dafny.MultiSet
		_ = _586_v176
		_586_v176 = _dafny.MultiSetOf(_585_v175, _585_v175, _585_v175, func(_pat_let10_0 D3) D3 {
			return func(_587_dt__update__tmp_h3 D3) D3 {
				return func(_pat_let11_0 _dafny.Sequence) D3 {
					return func(_588_dt__update_hcf17_h0 _dafny.Sequence) D3 {
						return Companion_D3_.Create_DC9_(_588_dt__update_hcf17_h0)
					}(_pat_let11_0)
				}(_pat_let_tv22)
			}(_pat_let10_0)
		}(_585_v175))
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_512_v126), 0))
		_ = _index60
		(_512_v126).ArraySet1((_dafny.Zero).Minus((_586_v176).Cardinality()), (_index60).Int())
		var _589_v177 _dafny.Map
		_ = _589_v177
		_589_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_321_v1, true)
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_512_v126), 0))
		_ = _index61
		(_512_v126).ArraySet1(Companion_Default___.Fm4(_320_v0, (func() bool {
			if (_589_v177).Contains(_321_v1) {
				return (_589_v177).Get(_321_v1).(bool)
			}
			return _455_v90
		})(), _324_globalState), (_index61).Int())
		var _590_v178 _dafny.Set
		_ = _590_v178
		_590_v178 = _dafny.SetOf(_449_v84, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC0_(false)).Dtor_cf0(), _dafny.IntOfUint32((_322_v2).Cardinality())), _449_v84)
		var _591_v179 _dafny.Set
		_ = _591_v179
		_591_v179 = _dafny.SetOf(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-660), _320_v0), (_590_v178).Cardinality(), _dafny.IntOfInt64(253))
		_591_v179 = _591_v179
		var _592_v180 *C0
		_ = _592_v180
		var _nw98 *C0 = New_C0_()
		_ = _nw98
		_nw98.Ctor__(Companion_Default___.Fm4((_591_v179).Cardinality(), _455_v90, _324_globalState), _454_v89)
		_592_v180 = _nw98
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_512_v126), 0))
		_ = _index62
		(_512_v126).ArraySet1(_320_v0, (_index62).Int())
	} else if _source12.Is_DC9() {
		var _593___mcc_h31 _dafny.Sequence = _source12.Get_().(D3_DC9).Cf17
		_ = _593___mcc_h31
		var _594_cf17 _dafny.Sequence = _593___mcc_h31
		_ = _594_cf17
		_449_v84 = ((_449_v84).Merge(_449_v84)).Merge(_449_v84)
		if _321_v1 {
			var _595_v181 _dafny.Map
			_ = _595_v181
			_595_v181 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_320_v0), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if _455_v90 {
					return _352_v20
				}
				return _352_v20
			})()).Cardinality()))
			_595_v181 = (_595_v181).Merge(_595_v181)
			var _596_v182 _dafny.CodePoint
			_ = _596_v182
			_596_v182 = _dafny.CodePoint('g')
			var _597_v183 _dafny.Map
			_ = _597_v183
			_597_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _596_v182)
			var _598_v184 _dafny.Map
			_ = _598_v184
			_598_v184 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_352_v20, (func() _dafny.CodePoint {
				if (_597_v183).Contains(_321_v1) {
					return (_597_v183).Get(_321_v1).(_dafny.CodePoint)
				}
				return _596_v182
			})())
			_598_v184 = (_598_v184).Update(_352_v20, _596_v182)
			_320_v0 = (_320_v0).Plus(_320_v0)
			var _599_v185 D2
			_ = _599_v185
			_599_v185 = Companion_D2_.Create_DC5_(_512_v126)
			var _pat_let_tv23 = _512_v126
			_ = _pat_let_tv23
			var _pat_let_tv24 = _512_v126
			_ = _pat_let_tv24
			_599_v185 = func(_pat_let12_0 D2) D2 {
				return func(_600_dt__update__tmp_h4 D2) D2 {
					return func(_pat_let13_0 _dafny.Array) D2 {
						return func(_601_dt__update_hcf11_h1 _dafny.Array) D2 {
							return Companion_D2_.Create_DC5_(_601_dt__update_hcf11_h1)
						}(_pat_let13_0)
					}((func() _dafny.Array {
						if true {
							return _pat_let_tv23
						}
						return _pat_let_tv24
					})())
				}(_pat_let12_0)
			}(_599_v185)
			var _602_v186 *C0
			_ = _602_v186
			var _nw99 *C0 = New_C0_()
			_ = _nw99
			_nw99.Ctor__(_320_v0, _454_v89)
			_602_v186 = _nw99
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_453_v88), 0))
			_ = _index63
			(_453_v88).ArraySet1(_602_v186, (_index63).Int())
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_453_v88), 0))
			_ = _index64
			(_453_v88).ArraySet1((func() *C0 {
				if _321_v1 {
					return _602_v186
				}
				return _602_v186
			})(), (_index64).Int())
		} else {
			var _603_v187 *C0
			_ = _603_v187
			var _nw100 *C0 = New_C0_()
			_ = _nw100
			_nw100.Ctor__(_320_v0, _dafny.SetOf(false, _321_v1))
			_603_v187 = _nw100
			_603_v187 = _603_v187
			_352_v20 = _352_v20
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_512_v126), 0))
			_ = _index65
			(_512_v126).ArraySet1(_320_v0, (_index65).Int())
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_512_v126), 0))
			_ = _index66
			(_512_v126).ArraySet1(_320_v0, (_index66).Int())
			_603_v187 = _603_v187
			var _604_v188 _dafny.Sequence
			_ = _604_v188
			_604_v188 = _dafny.SeqOf(_352_v20)
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_512_v126), 0))
			_ = _index67
			(_512_v126).ArraySet1(((_dafny.IntOfUint32((_604_v188).Cardinality())).Times(_dafny.IntOfInt64(-818))).Times((_452_v87).Cardinality()), (_index67).Int())
		}
		var _605_v189 _dafny.Array
		_ = _605_v189
		var _nwElement0_20 bool = (_322_v2).Select((Companion_Default___.SafeIndex(_320_v0, _dafny.IntOfUint32((_322_v2).Cardinality()))).Uint32()).(bool)
		_ = _nwElement0_20
		var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(26))
		_ = _nw101
		(_nw101).ArraySet1(_nwElement0_20, 0)
		(_nw101).ArraySet1(_321_v1, 1)
		(_nw101).ArraySet1(_321_v1, 2)
		(_nw101).ArraySet1(_455_v90, 3)
		(_nw101).ArraySet1(false, 4)
		(_nw101).ArraySet1(true, 5)
		(_nw101).ArraySet1(!(false), 6)
		(_nw101).ArraySet1(_321_v1, 7)
		(_nw101).ArraySet1(_321_v1, 8)
		(_nw101).ArraySet1(_455_v90, 9)
		(_nw101).ArraySet1(_321_v1, 10)
		(_nw101).ArraySet1(!(Companion_Default___.Fm0(_321_v1, _321_v1, Companion_Default___.Fm4(_dafny.IntOfUint32((_322_v2).Cardinality()), _321_v1, _324_globalState), _324_globalState)), 11)
		(_nw101).ArraySet1(_321_v1, 12)
		(_nw101).ArraySet1(_321_v1, 13)
		(_nw101).ArraySet1(_321_v1, 14)
		(_nw101).ArraySet1(false, 15)
		(_nw101).ArraySet1(_455_v90, 16)
		(_nw101).ArraySet1(_321_v1, 17)
		(_nw101).ArraySet1(_455_v90, 18)
		(_nw101).ArraySet1(true, 19)
		(_nw101).ArraySet1(_455_v90, 20)
		(_nw101).ArraySet1(Companion_Default___.Fm0(_321_v1, (_383_v43).Dtor_cf5(), _320_v0, _324_globalState), 21)
		(_nw101).ArraySet1(!(_321_v1), 22)
		(_nw101).ArraySet1(_321_v1, 23)
		(_nw101).ArraySet1(_321_v1, 24)
		(_nw101).ArraySet1(_321_v1, 25)
		_605_v189 = _nw101
		var _606_v190 _dafny.Map
		_ = _606_v190
		_606_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_455_v90, _605_v189)
		var _607_v191 _dafny.Array
		_ = _607_v191
		var _nwElement0_21 _dafny.Array = (func() _dafny.Array {
			if (_606_v190).Contains(_321_v1) {
				return (_606_v190).Get(_321_v1).(_dafny.Array)
			}
			return _605_v189
		})()
		_ = _nwElement0_21
		var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(26))
		_ = _nw102
		(_nw102).ArraySet1(_nwElement0_21, 0)
		(_nw102).ArraySet1(_605_v189, 1)
		(_nw102).ArraySet1(_605_v189, 2)
		(_nw102).ArraySet1(_605_v189, 3)
		(_nw102).ArraySet1(_605_v189, 4)
		(_nw102).ArraySet1(_605_v189, 5)
		(_nw102).ArraySet1((func() _dafny.Array {
			if _455_v90 {
				return _605_v189
			}
			return _605_v189
		})(), 6)
		(_nw102).ArraySet1(_605_v189, 7)
		(_nw102).ArraySet1(_605_v189, 8)
		(_nw102).ArraySet1(_605_v189, 9)
		(_nw102).ArraySet1(_605_v189, 10)
		(_nw102).ArraySet1(_605_v189, 11)
		(_nw102).ArraySet1(_605_v189, 12)
		(_nw102).ArraySet1(_605_v189, 13)
		(_nw102).ArraySet1(_605_v189, 14)
		(_nw102).ArraySet1(_605_v189, 15)
		(_nw102).ArraySet1(_605_v189, 16)
		(_nw102).ArraySet1(_605_v189, 17)
		(_nw102).ArraySet1(_605_v189, 18)
		(_nw102).ArraySet1((func() _dafny.Array {
			if _321_v1 {
				return _605_v189
			}
			return _605_v189
		})(), 19)
		(_nw102).ArraySet1(_605_v189, 20)
		(_nw102).ArraySet1(_605_v189, 21)
		(_nw102).ArraySet1(_605_v189, 22)
		(_nw102).ArraySet1((func() _dafny.Array {
			if _321_v1 {
				return _605_v189
			}
			return _605_v189
		})(), 23)
		(_nw102).ArraySet1(_605_v189, 24)
		(_nw102).ArraySet1(_605_v189, 25)
		_607_v191 = _nw102
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(215), _dafny.ArrayLen((_607_v191), 0))
		_ = _index68
		(_607_v191).ArraySet1(_605_v189, (_index68).Int())
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(215), _dafny.ArrayLen((_607_v191), 0))
		_ = _index69
		(_607_v191).ArraySet1(_605_v189, (_index69).Int())
		var _source13 D1 = _458_v93
		_ = _source13
		if _source13.Is_DC3() {
			var _608___mcc_h33 bool = _source13.Get_().(D1_DC3).Cf5
			_ = _608___mcc_h33
			var _609___mcc_h34 bool = _source13.Get_().(D1_DC3).Cf6
			_ = _609___mcc_h34
			var _610___mcc_h35 _dafny.Sequence = _source13.Get_().(D1_DC3).Cf7
			_ = _610___mcc_h35
			var _611___mcc_h36 bool = _source13.Get_().(D1_DC3).Cf8
			_ = _611___mcc_h36
			var _612___mcc_h37 _dafny.Int = _source13.Get_().(D1_DC3).Cf9
			_ = _612___mcc_h37
			var _613_cf9 _dafny.Int = _612___mcc_h37
			_ = _613_cf9
			var _614_cf8 bool = _611___mcc_h36
			_ = _614_cf8
			var _615_cf7 _dafny.Sequence = _610___mcc_h35
			_ = _615_cf7
			var _616_cf6 bool = _609___mcc_h34
			_ = _616_cf6
			var _617_cf5 bool = _608___mcc_h33
			_ = _617_cf5
			var _618_v192 _dafny.CodePoint
			_ = _618_v192
			_618_v192 = _dafny.CodePoint('v')
			var _619_v193 _dafny.Map
			_ = _619_v193
			_619_v193 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_614_cf8, false)
			var _620_v194 _dafny.Array
			_ = _620_v194
			var _nwElement0_22 _dafny.Set = _dafny.SetOf(_614_cf8, (func() bool {
				if (_619_v193).Contains(_455_v90) {
					return (_619_v193).Get(_455_v90).(bool)
				}
				return _616_cf6
			})())
			_ = _nwElement0_22
			var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(2))
			_ = _nw103
			(_nw103).ArraySet1(_nwElement0_22, 0)
			(_nw103).ArraySet1(_dafny.SetOf(_617_cf5, (_322_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_511_v125).Cardinality()), _dafny.IntOfUint32((_322_v2).Cardinality()))).Uint32()).(bool), !(_455_v90)), 1)
			_620_v194 = _nw103
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_620_v194), 0))
			_ = _index70
			(_620_v194).ArraySet1(_454_v89, (_index70).Int())
			var _621_v195 _dafny.Map
			_ = _621_v195
			_621_v195 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v0, Companion_D3_.Create_DC11_())
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_620_v194), 0))
			_ = _index71
			var _rhs63 _dafny.Int = _613_cf9
			_ = _rhs63
			var _rhs64 bool = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_594_cf17, _511_v125))).IsProperSubsetOf(_dafny.MultiSetOf(_320_v0, _320_v0, _320_v0))
			_ = _rhs64
			var _rhs65 _dafny.CodePoint = _618_v192
			_ = _rhs65
			var _rhs66 _dafny.Set = (func() _dafny.Set {
				if !(((_621_v195).Cardinality()).Cmp((_457_v92).Cardinality()) >= 0) {
					return _454_v89
				}
				return _454_v89
			})()
			_ = _rhs66
			var _lhs43 *GlobalState = _324_globalState
			_ = _lhs43
			var _lhs44 _dafny.Array = _620_v194
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_620_v194), 0))
			_ = _lhs45
			_lhs43.F0 = _rhs63
			_614_cf8 = _rhs64
			_618_v192 = _rhs65
			(_lhs44).ArraySet1(_rhs66, (_lhs45).Int())
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_605_v189), 0))
			_ = _index72
			(_605_v189).ArraySet1(true, (_index72).Int())
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_605_v189), 0))
			_ = _index73
			(_605_v189).ArraySet1(Companion_Default___.Fm0(_321_v1, _455_v90, _613_cf9, _324_globalState), (_index73).Int())
			var _622_v196 *C0
			_ = _622_v196
			var _nw104 *C0 = New_C0_()
			_ = _nw104
			_nw104.Ctor__(Companion_Default___.SafeDivisionInt(_613_cf9, (_dafny.Zero).Minus((_606_v190).Cardinality())), Companion_Default___.Fm5(_324_globalState))
			_622_v196 = _nw104
			_616_cf6 = true
		} else if _source13.Is_DC2() {
			var _623___mcc_h38 _dafny.Int = _source13.Get_().(D1_DC2).Cf4
			_ = _623___mcc_h38
			var _624_cf4 _dafny.Int = _623___mcc_h38
			_ = _624_cf4
			var _625_v197 *C0
			_ = _625_v197
			var _nw105 *C0 = New_C0_()
			_ = _nw105
			_nw105.Ctor__(Companion_Default___.SafeModuloInt(Companion_Default___.Fm4(_624_cf4, _321_v1, _324_globalState), (_dafny.Zero).Minus(_320_v0)), (func() _dafny.Set {
				if _455_v90 {
					return _dafny.SetOf(_455_v90)
				}
				return _454_v89
			})())
			_625_v197 = _nw105
			var _626_v198 _dafny.Map
			_ = _626_v198
			_626_v198 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_455_v90, (_322_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.IntOfUint32((_322_v2).Cardinality()))).Uint32()).(bool))
			_625_v197 = (func() *C0 {
				if (func() bool {
					if (_626_v198).Contains(_455_v90) {
						return (_626_v198).Get(_455_v90).(bool)
					}
					return _455_v90
				})() {
					return _625_v197
				}
				return _625_v197
			})()
			var _627_v199 D4
			_ = _627_v199
			_627_v199 = Companion_D4_.Create_DC13_(_625_v197)
			_625_v197 = (_627_v199).Dtor_cf22()
			(_324_globalState).F4 = ((_625_v197).F6()).Plus(_624_cf4)
			(_324_globalState).F4 = (((_449_v84).Merge(_449_v84)).Merge((_449_v84).Merge(_449_v84))).Cardinality()
		} else {
			var _628___mcc_h39 D1 = _source13.Get_().(D1_DC4).Cf10
			_ = _628___mcc_h39
			var _629_cf10 D1 = _628___mcc_h39
			_ = _629_cf10
			var _630_v200 _dafny.Sequence
			_ = _630_v200
			_630_v200 = _dafny.SeqOf(_512_v126)
			var _631_v201 _dafny.Map
			_ = _631_v201
			_631_v201 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_455_v90, _321_v1, _320_v0, _324_globalState), _630_v200)
			_631_v201 = (_631_v201).Update(_455_v90, (func() _dafny.Sequence {
				if _321_v1 {
					return _630_v200
				}
				return _630_v200
			})())
			(_324_globalState).F4 = (((_449_v84).Merge((_449_v84).Update(_321_v1, _320_v0))).Cardinality()).Times(Companion_Default___.Fm4((_dafny.Zero).Minus(_dafny.IntOfUint32((_352_v20).Cardinality())), Companion_Default___.Fm0(_321_v1, _321_v1, _320_v0, _324_globalState), _324_globalState))
			var _632_v202 *C0
			_ = _632_v202
			var _nw106 *C0 = New_C0_()
			_ = _nw106
			_nw106.Ctor__(_dafny.IntOfInt64(382), _454_v89)
			_632_v202 = _nw106
			var _633_v203 _dafny.Set
			_ = _633_v203
			_633_v203 = _dafny.SetOf(_632_v202, _632_v202)
			var _634_v204 _dafny.CodePoint
			_ = _634_v204
			_634_v204 = _dafny.CodePoint('r')
			var _635_v205 D4
			_ = _635_v205
			_635_v205 = Companion_D4_.Create_DC16_(_633_v203, _634_v204, (_632_v202).F6(), _634_v204)
			var _636_v206 _dafny.Map
			_ = _636_v206
			_636_v206 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v0, (_dafny.Zero).Minus(_320_v0))
			var _rhs67 bool = (Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_632_v202).F6(), _632_v202)).Cardinality(), (_632_v202).F6())).Cmp((_dafny.Zero).Minus(_320_v0)) != 0
			_ = _rhs67
			var _rhs68 D4 = _635_v205
			_ = _rhs68
			var _rhs69 _dafny.Map = _636_v206
			_ = _rhs69
			var _rhs70 _dafny.Array = _512_v126
			_ = _rhs70
			_455_v90 = _rhs67
			_635_v205 = _rhs68
			_636_v206 = _rhs69
			_512_v126 = _rhs70
		}
	} else {
		var _637___mcc_h32 D3 = _source12.Get_().(D3_DC12).Cf21
		_ = _637___mcc_h32
		var _638_cf21 D3 = _637___mcc_h32
		_ = _638_cf21
		var _639_v207 bool
		_ = _639_v207
		var _out10 bool
		_ = _out10
		_out10 = Companion_Default___.M0(!(_455_v90), _324_globalState)
		_639_v207 = _out10
		(_324_globalState).F0 = _320_v0
		var _640_v208 D3
		_ = _640_v208
		_640_v208 = Companion_D3_.Create_DC10_(!(_455_v90), _320_v0, _dafny.IntOfInt64(-144))
		var _pat_let_tv25 = _320_v0
		_ = _pat_let_tv25
		var _641_v209 _dafny.Array
		_ = _641_v209
		var _nwElement0_23 bool = _455_v90
		_ = _nwElement0_23
		var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(23))
		_ = _nw107
		(_nw107).ArraySet1(_nwElement0_23, 0)
		(_nw107).ArraySet1(_455_v90, 1)
		(_nw107).ArraySet1(_455_v90, 2)
		(_nw107).ArraySet1(!(_321_v1), 3)
		(_nw107).ArraySet1(_455_v90, 4)
		(_nw107).ArraySet1(false, 5)
		(_nw107).ArraySet1((_639_v207) && (!(_639_v207)), 6)
		(_nw107).ArraySet1(_321_v1, 7)
		(_nw107).ArraySet1(_321_v1, 8)
		(_nw107).ArraySet1(!(true), 9)
		(_nw107).ArraySet1(_455_v90, 10)
		(_nw107).ArraySet1(_321_v1, 11)
		(_nw107).ArraySet1(_639_v207, 12)
		(_nw107).ArraySet1(((_dafny.MultiSetOf(_321_v1)).Cardinality()).Cmp(_dafny.IntOfInt64(-508)) < 0, 13)
		(_nw107).ArraySet1(!(_455_v90) || (false), 14)
		(_nw107).ArraySet1(_455_v90, 15)
		(_nw107).ArraySet1(_321_v1, 16)
		(_nw107).ArraySet1(_639_v207, 17)
		(_nw107).ArraySet1((func(_pat_let14_0 D3) D3 {
			return func(_642_dt__update__tmp_h5 D3) D3 {
				return func(_pat_let15_0 _dafny.Int) D3 {
					return func(_643_dt__update_hcf20_h0 _dafny.Int) D3 {
						return Companion_D3_.Create_DC10_((_642_dt__update__tmp_h5).Dtor_cf18(), (_642_dt__update__tmp_h5).Dtor_cf19(), _643_dt__update_hcf20_h0)
					}(_pat_let15_0)
				}(_pat_let_tv25)
			}(_pat_let14_0)
		}(_640_v208)).Dtor_cf18(), 18)
		(_nw107).ArraySet1(_455_v90, 19)
		(_nw107).ArraySet1(_455_v90, 20)
		(_nw107).ArraySet1(_639_v207, 21)
		(_nw107).ArraySet1((_322_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-229), _dafny.IntOfUint32((_322_v2).Cardinality()))).Uint32()).(bool), 22)
		_641_v209 = _nw107
		var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_641_v209), 0))
		_ = _index74
		(_641_v209).ArraySet1(_455_v90, (_index74).Int())
		var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_641_v209), 0))
		_ = _index75
		(_641_v209).ArraySet1(_321_v1, (_index75).Int())
		var _644_v210 _dafny.Map
		_ = _644_v210
		_644_v210 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_456_v91).Cardinality(), _320_v0)
		var _645_v212 *C0
		_ = _645_v212
		var _nw108 *C0 = New_C0_()
		_ = _nw108
		_nw108.Ctor__((func() _dafny.Set {
			var _coll26 = _dafny.NewBuilder()
			_ = _coll26
			for _iter28 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_644_v210, _322_v2)).Keys().Elements()); ; {
				_compr_26, _ok28 := _iter28()
				if !_ok28 {
					break
				}
				var _646_v211 _dafny.Map
				_646_v211 = interface{}(_compr_26).(_dafny.Map)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_644_v210, _322_v2)).Contains(_646_v211) {
					_coll26.Add(_646_v211)
				}
			}
			return _coll26.ToSet()
		}()).Cardinality(), _dafny.SetOf(_321_v1))
		_645_v212 = _nw108
	}
	if !((_452_v87).Union(_452_v87)).Equals(_452_v87) {
		var _647_v213 _dafny.CodePoint
		_ = _647_v213
		_647_v213 = _dafny.CodePoint('m')
		_647_v213 = _647_v213
		_456_v91 = (_456_v91).Merge(_456_v91)
		_352_v20 = _dafny.UnicodeSeqOfUtf8Bytes("qyqejkm")
		var _648_v214 D2
		_ = _648_v214
		_648_v214 = Companion_D2_.Create_DC7_(Companion_Default___.Fm4((_dafny.Zero).Minus(_320_v0), _321_v1, _324_globalState))
		var _649_v215 D2
		_ = _649_v215
		_649_v215 = Companion_D2_.Create_DC8_(_648_v214)
		var _source14 D2 = _649_v215
		_ = _source14
		if _source14.Is_DC6() {
			var _650___mcc_h40 _dafny.Int = _source14.Get_().(D2_DC6).Cf12
			_ = _650___mcc_h40
			var _651___mcc_h41 bool = _source14.Get_().(D2_DC6).Cf13
			_ = _651___mcc_h41
			var _652___mcc_h42 _dafny.Array = _source14.Get_().(D2_DC6).Cf14
			_ = _652___mcc_h42
			var _653_cf14 _dafny.Array = _652___mcc_h42
			_ = _653_cf14
			var _654_cf13 bool = _651___mcc_h41
			_ = _654_cf13
			var _655_cf12 _dafny.Int = _650___mcc_h40
			_ = _655_cf12
			var _656_v216 _dafny.Array
			_ = _656_v216
			var _nwElement0_24 bool = (_320_v0).Cmp((_511_v125).Select((Companion_Default___.SafeIndex(_655_cf12, _dafny.IntOfUint32((_511_v125).Cardinality()))).Uint32()).(_dafny.Int)) == 0
			_ = _nwElement0_24
			var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(6))
			_ = _nw109
			(_nw109).ArraySet1(_nwElement0_24, 0)
			(_nw109).ArraySet1(true, 1)
			(_nw109).ArraySet1(!((func() bool {
				if (_456_v91).Contains(_655_cf12) {
					return (_456_v91).Get(_655_cf12).(bool)
				}
				return _654_cf13
			})()), 2)
			(_nw109).ArraySet1(_455_v90, 3)
			(_nw109).ArraySet1(_321_v1, 4)
			(_nw109).ArraySet1(_321_v1, 5)
			_656_v216 = _nw109
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_656_v216), 0))
			_ = _index76
			(_656_v216).ArraySet1(_321_v1, (_index76).Int())
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_656_v216), 0))
			_ = _index77
			(_656_v216).ArraySet1(_321_v1, (_index77).Int())
			var _657_v217 _dafny.Array
			_ = _657_v217
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_11
			var _nw110 _dafny.Array
			_ = _nw110
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw110 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.Sequence = (func(_658_v20 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_659_i13 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_658_v20, _658_v20)
					}
				})(_352_v20)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw110 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw110).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw110).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_657_v217 = _nw110
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_657_v217), 0))
			_ = _index78
			(_657_v217).ArraySet1(_352_v20, (_index78).Int())
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_657_v217), 0))
			_ = _index79
			(_657_v217).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_352_v20, _352_v20), (_index79).Int())
			var _660_v220 _dafny.Array
			_ = _660_v220
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_12
			var _nw111 _dafny.Array
			_ = _nw111
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw111 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Set = (func(_661_cf12 _dafny.Int) func(_dafny.Int) _dafny.Set {
					return func(_662_i14 _dafny.Int) _dafny.Set {
						return (func() _dafny.Set {
							var _coll27 = _dafny.NewBuilder()
							_ = _coll27
							for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(873), _dafny.IntOfInt64(129))); ; {
								_compr_27, _ok29 := _iter29()
								if !_ok29 {
									break
								}
								var _663_v218 _dafny.Int
								_663_v218 = interface{}(_compr_27).(_dafny.Int)
								if ((_dafny.IntOfInt64(873)).Cmp(_663_v218) <= 0) && ((_663_v218).Cmp(_dafny.IntOfInt64(129)) < 0) {
									_coll27.Add((_663_v218).Minus(_661_cf12))
								}
							}
							return _coll27.ToSet()
						}()).Difference(func() _dafny.Set {
							var _coll28 = _dafny.NewBuilder()
							_ = _coll28
							for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(15), _dafny.IntOfInt64(772))); ; {
								_compr_28, _ok30 := _iter30()
								if !_ok30 {
									break
								}
								var _664_v219 _dafny.Int
								_664_v219 = interface{}(_compr_28).(_dafny.Int)
								if ((_dafny.IntOfInt64(15)).Cmp(_664_v219) <= 0) && ((_664_v219).Cmp(_dafny.IntOfInt64(772)) < 0) {
									_coll28.Add((_664_v219).Plus(_661_cf12))
								}
							}
							return _coll28.ToSet()
						}())
					}
				})(_655_cf12)
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
			_660_v220 = _nw111
			var _665_v221 _dafny.Set
			_ = _665_v221
			_665_v221 = _dafny.SetOf((_dafny.SetOf(_655_cf12, _dafny.IntOfUint32((_352_v20).Cardinality()))).Cardinality(), _320_v0)
			var _666_v222 _dafny.Sequence
			_ = _666_v222
			_666_v222 = _dafny.SeqOf(_454_v89, _454_v89)
			var _667_v223 *C0
			_ = _667_v223
			var _nw112 *C0 = New_C0_()
			_ = _nw112
			_nw112.Ctor__(_655_cf12, (_666_v222).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_657_v217).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_657_v217), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_666_v222).Cardinality()))).Uint32()).(_dafny.Set))
			_667_v223 = _nw112
			var _668_v224 _dafny.Set
			_ = _668_v224
			_668_v224 = _dafny.SetOf(_667_v223)
			var _669_v225 D4
			_ = _669_v225
			_669_v225 = Companion_D4_.Create_DC16_(_668_v224, _647_v213, _320_v0, _647_v213)
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_660_v220), 0))
			_ = _index80
			(_660_v220).ArraySet1((func() _dafny.Set {
				if true {
					return _665_v221
				}
				return _dafny.SetOf((_665_v221).Cardinality(), _655_cf12, _655_cf12, (_669_v225).Dtor_cf27(), _655_cf12)
			})(), (_index80).Int())
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_660_v220), 0))
			_ = _index81
			(_660_v220).ArraySet1(_665_v221, (_index81).Int())
			var _670_v226 D3
			_ = _670_v226
			_670_v226 = Companion_D3_.Create_DC12_(_564_v162)
			var _671_v227 _dafny.Map
			_ = _671_v227
			_671_v227 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_670_v226, _667_v223)
			_671_v227 = (_671_v227).Update(Companion_D3_.Create_DC12_(_564_v162), _667_v223)
		} else if _source14.Is_DC7() {
			var _672___mcc_h43 _dafny.Int = _source14.Get_().(D2_DC7).Cf15
			_ = _672___mcc_h43
			var _673_cf15 _dafny.Int = _672___mcc_h43
			_ = _673_cf15
			var _674_v228 *C0
			_ = _674_v228
			var _nw113 *C0 = New_C0_()
			_ = _nw113
			_nw113.Ctor__(_dafny.IntOfUint32((_511_v125).Cardinality()), _454_v89)
			_674_v228 = _nw113
			var _675_v229 D4
			_ = _675_v229
			_675_v229 = Companion_D4_.Create_DC13_(_674_v228)
			var _676_v230 _dafny.Map
			_ = _676_v230
			_676_v230 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_647_v213, Companion_Default___.Fm0(_455_v90, _321_v1, _320_v0, _324_globalState))
			var _677_v231 D1
			_ = _677_v231
			_677_v231 = Companion_D1_.Create_DC3_(_321_v1, Companion_Default___.Fm0(Companion_Default___.Fm0(_455_v90, _321_v1, _dafny.IntOfInt64(844), _324_globalState), (func() bool {
				if (_676_v230).Contains(_647_v213) {
					return (_676_v230).Get(_647_v213).(bool)
				}
				return true
			})(), _dafny.IntOfUint32((_322_v2).Cardinality()), _324_globalState), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("tim"), (Companion_Default___.SafeIndex((_452_v87).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tim")).Cardinality()))).Uint32(), _dafny.CodePoint('u')), _321_v1, _320_v0)
			var _678_v232 D1
			_ = _678_v232
			_678_v232 = Companion_D1_.Create_DC4_(_677_v231)
			var _rhs71 D4 = Companion_D4_.Create_DC13_((func() *C0 {
				if _455_v90 {
					return _674_v228
				}
				return _674_v228
			})())
			_ = _rhs71
			var _rhs72 D1 = Companion_D1_.Create_DC4_(Companion_Default___.Fm14(_324_globalState))
			_ = _rhs72
			var _rhs73 _dafny.Int = _673_cf15
			_ = _rhs73
			var _lhs46 *GlobalState = _324_globalState
			_ = _lhs46
			_675_v229 = _rhs71
			_678_v232 = _rhs72
			_lhs46.F0 = _rhs73
			var _679_v233 bool
			_ = _679_v233
			var _out11 bool
			_ = _out11
			_out11 = Companion_Default___.M0(!(_321_v1), _324_globalState)
			_679_v233 = _out11
			_679_v233 = !((func() bool {
				if Companion_Default___.Fm0(!(_455_v90), _321_v1, _dafny.IntOfUint32((_511_v125).Cardinality()), _324_globalState) {
					return _679_v233
				}
				return _321_v1
			})())
			_320_v0 = _dafny.IntOfInt64(230)
		} else if _source14.Is_DC5() {
			var _680___mcc_h44 _dafny.Array = _source14.Get_().(D2_DC5).Cf11
			_ = _680___mcc_h44
			var _681_cf11 _dafny.Array = _680___mcc_h44
			_ = _681_cf11
			(_324_globalState).F0 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_320_v0, (func() _dafny.Int {
				if (_322_v2).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm4(_320_v0, _321_v1, _324_globalState), _dafny.IntOfUint32((_322_v2).Cardinality()))).Uint32()).(bool) {
					return _320_v0
				}
				return _320_v0
			})()))
			var _682_v234 _dafny.Array
			_ = _682_v234
			var _nw114 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw114
			_682_v234 = _nw114
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_682_v234), 0))
			_ = _index82
			(_682_v234).ArraySet1(_321_v1, (_index82).Int())
			var _683_v235 D3
			_ = _683_v235
			_683_v235 = Companion_D3_.Create_DC12_(Companion_D3_.Create_DC11_())
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_682_v234), 0))
			_ = _index83
			var _rhs74 bool = _455_v90
			_ = _rhs74
			var _rhs75 D3 = _683_v235
			_ = _rhs75
			var _rhs76 _dafny.Array = _682_v234
			_ = _rhs76
			var _rhs77 bool = (Companion_Default___.Fm4(_320_v0, _321_v1, _324_globalState)).Cmp(((Companion_D2_.Create_DC6_(_320_v0, _455_v90, _512_v126)).Dtor_cf12()).Times(_320_v0)) <= 0
			_ = _rhs77
			var _rhs78 bool = Companion_Default___.Fm0(_455_v90, _321_v1, _320_v0, _324_globalState)
			_ = _rhs78
			var _lhs47 _dafny.Array = _682_v234
			_ = _lhs47
			var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_682_v234), 0))
			_ = _lhs48
			(_lhs47).ArraySet1(_rhs74, (_lhs48).Int())
			_683_v235 = _rhs75
			_682_v234 = _rhs76
			_455_v90 = _rhs77
			_455_v90 = _rhs78
			var _684_v236 D3
			_ = _684_v236
			_684_v236 = Companion_D3_.Create_DC10_(_455_v90, _320_v0, _320_v0)
			(_324_globalState).F0 = (_dafny.IntOfInt64(722)).Plus(((_dafny.Zero).Minus((_452_v87).Cardinality())).Times((_684_v236).Dtor_cf19()))
			(_324_globalState).F0 = ((_320_v0).Times(_320_v0)).Minus(_320_v0)
		} else {
			var _685___mcc_h45 D2 = _source14.Get_().(D2_DC8).Cf16
			_ = _685___mcc_h45
			var _686_cf16 D2 = _685___mcc_h45
			_ = _686_cf16
			(_324_globalState).F0 = (_320_v0).Times(Companion_Default___.SafeDivisionInt(_320_v0, (_dafny.Zero).Minus(((_dafny.MultiSetFromSeq(_322_v2)).Update(_321_v1, Companion_Default___.Abs(_320_v0))).Cardinality())))
			_455_v90 = !(true)
			var _687_v237 *C0
			_ = _687_v237
			var _nw115 *C0 = New_C0_()
			_ = _nw115
			_nw115.Ctor__(_320_v0, _454_v89)
			_687_v237 = _nw115
			var _688_v238 D3
			_ = _688_v238
			_688_v238 = Companion_D3_.Create_DC10_(_455_v90, _dafny.IntOfInt64(354), (_687_v237).F6())
			_321_v1 = (_688_v238).Dtor_cf18()
		}
		var _pat_let_tv26 = _455_v90
		_ = _pat_let_tv26
		var _689_v239 _dafny.Sequence
		_ = _689_v239
		_689_v239 = _dafny.SeqOf(func(_pat_let16_0 D4) D4 {
			return func(_690_dt__update__tmp_h6 D4) D4 {
				return func(_pat_let17_0 bool) D4 {
					return func(_691_dt__update_hcf23_h0 bool) D4 {
						return Companion_D4_.Create_DC15_(_691_dt__update_hcf23_h0, (_690_dt__update__tmp_h6).Dtor_cf24())
					}(_pat_let17_0)
				}(_pat_let_tv26)
			}(_pat_let16_0)
		}(Companion_D4_.Create_DC15_(_321_v1, _320_v0)), Companion_D4_.Create_DC15_(!(_321_v1), _dafny.IntOfInt64(412)))
		(_324_globalState).F0 = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_352_v20).Cardinality()), _320_v0)).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_689_v239, _689_v239)).Cardinality()))
	} else {
		var _692_v240 *C0
		_ = _692_v240
		var _nw116 *C0 = New_C0_()
		_ = _nw116
		_nw116.Ctor__(_dafny.IntOfUint32((_352_v20).Cardinality()), _454_v89)
		_692_v240 = _nw116
		(_324_globalState).F0 = (_692_v240).F6()
		var _693_v241 _dafny.Map
		_ = _693_v241
		_693_v241 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_352_v20, _352_v20), (_322_v2).Select((Companion_Default___.SafeIndex((_692_v240).F6(), _dafny.IntOfUint32((_322_v2).Cardinality()))).Uint32()).(bool))
		_693_v241 = (_693_v241).Update(_dafny.Companion_Sequence_.Concatenate(_352_v20, _352_v20), _455_v90)
		var _694_v242 *C0
		_ = _694_v242
		var _nw117 *C0 = New_C0_()
		_ = _nw117
		_nw117.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_457_v92).Cardinality()), _511_v125)).Cardinality()), (_692_v240).F7())
		_694_v242 = _nw117
		var _695_v243 bool
		_ = _695_v243
		var _out12 bool
		_ = _out12
		_out12 = Companion_Default___.M0(true, _324_globalState)
		_695_v243 = _out12
	}
	var _696_v244 _dafny.Array
	_ = _696_v244
	var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
	_ = _nw118
	_696_v244 = _nw118
	var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.ArrayLen((_696_v244), 0))
	_ = _index84
	(_696_v244).ArraySet1(_dafny.SeqOf(_512_v126, _512_v126), (_index84).Int())
	var _697_v245 _dafny.Sequence
	_ = _697_v245
	_697_v245 = _dafny.SeqOf(_512_v126, _512_v126)
	var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.ArrayLen((_696_v244), 0))
	_ = _index85
	(_696_v244).ArraySet1(_697_v245, (_index85).Int())
	_dafny.Print(_320_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_321_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_322_v2, _dafny.SeqOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_323_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-45), _dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_324_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_324_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_324_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_324_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_324_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_324_globalState).F5()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-45), _dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_325_i0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_352_v20.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_353_v21).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(6), _dafny.IntOfInt64(6))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_383_v43).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_383_v43).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_383_v43).Dtor_cf7().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_383_v43).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_383_v43).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_449_v84).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_450_v85).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(492), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_451_v86).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_452_v87).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_454_v89).Equals(_dafny.SetOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_455_v90)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_456_v91).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(6), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_457_v92).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _dafny.IntOfInt64(6))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_458_v93).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_459_v94).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_459_v94).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_459_v94).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_459_v94).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_459_v94).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_459_v94).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_459_v94).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_459_v94).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_511_v125, _dafny.SeqOf(_dafny.Zero, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_512_v126).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_513_v127).Dtor_cf12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_513_v127).Dtor_cf13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_513_v127).Dtor_cf14()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_564_v162).Dtor_cf17(), _dafny.SeqOf(_dafny.Zero, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32(((_696_v244).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence)).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_697_v245).Cardinality()))
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

type D0_DC1 struct {
	Cf1 _dafny.MultiSet
	Cf2 _dafny.Array
	Cf3 _dafny.Array
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.MultiSet, Cf2 _dafny.Array, Cf3 _dafny.Array) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(false)
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.MultiSet {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf1.Equals(data2.Cf1) && data1.Cf2 == data2.Cf2 && data1.Cf3 == data2.Cf3
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

type D1_DC3 struct {
	Cf5 bool
	Cf6 bool
	Cf7 _dafny.Sequence
	Cf8 bool
	Cf9 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf5 bool, Cf6 bool, Cf7 _dafny.Sequence, Cf8 bool, Cf9 _dafny.Int) D1 {
	return D1{D1_DC3{Cf5, Cf6, Cf7, Cf8, Cf9}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC2 struct {
	Cf4 _dafny.Int
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf4 _dafny.Int) D1 {
	return D1{D1_DC2{Cf4}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC4 struct {
	Cf10 D1
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf10 D1) D1 {
	return D1{D1_DC4{Cf10}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(false, false, _dafny.EmptySeq, false, _dafny.Zero)
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf6() bool {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC3).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf9
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf10() D1 {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + data.Cf7.VerbatimString(true) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf10) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf5 == data2.Cf5 && data1.Cf6 == data2.Cf6 && data1.Cf7.Equals(data2.Cf7) && data1.Cf8 == data2.Cf8 && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf10.Equals(data2.Cf10)
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

type D2_DC6 struct {
	Cf12 _dafny.Int
	Cf13 bool
	Cf14 _dafny.Array
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf12 _dafny.Int, Cf13 bool, Cf14 _dafny.Array) D2 {
	return D2{D2_DC6{Cf12, Cf13, Cf14}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf15 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf15 _dafny.Int) D2 {
	return D2{D2_DC7{Cf15}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC5 struct {
	Cf11 _dafny.Array
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf11 _dafny.Array) D2 {
	return D2{D2_DC5{Cf11}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC8 struct {
	Cf16 D2
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf16 D2) D2 {
	return D2{D2_DC8{Cf16}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(_dafny.Zero, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf12
}

func (_this D2) Dtor_cf13() bool {
	return _this.Get_().(D2_DC6).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Array {
	return _this.Get_().(D2_DC6).Cf14
}

func (_this D2) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf15
}

func (_this D2) Dtor_cf11() _dafny.Array {
	return _this.Get_().(D2_DC5).Cf11
}

func (_this D2) Dtor_cf16() D2 {
	return _this.Get_().(D2_DC8).Cf16
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf16) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf12.Cmp(data2.Cf12) == 0 && data1.Cf13 == data2.Cf13 && data1.Cf14 == data2.Cf14
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf11 == data2.Cf11
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf16.Equals(data2.Cf16)
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
	Cf18 bool
	Cf19 _dafny.Int
	Cf20 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf18 bool, Cf19 _dafny.Int, Cf20 _dafny.Int) D3 {
	return D3{D3_DC10{Cf18, Cf19, Cf20}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC11 struct {
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_() D3 {
	return D3{D3_DC11{}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC9 struct {
	Cf17 _dafny.Sequence
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf17 _dafny.Sequence) D3 {
	return D3{D3_DC9{Cf17}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC12 struct {
	Cf21 D3
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf21 D3) D3 {
	return D3{D3_DC12{Cf21}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D3) Dtor_cf18() bool {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf19
}

func (_this D3) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf20
}

func (_this D3) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D3_DC9).Cf17
}

func (_this D3) Dtor_cf21() D3 {
	return _this.Get_().(D3_DC12).Cf21
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf18 == data2.Cf18 && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D3_DC11:
		{
			_, ok := other.Get_().(D3_DC11)
			return ok
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf17.Equals(data2.Cf17)
		}
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf21.Equals(data2.Cf21)
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

type D4_DC14 struct {
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_() D4 {
	return D4{D4_DC14{}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC15 struct {
	Cf23 bool
	Cf24 _dafny.Int
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf23 bool, Cf24 _dafny.Int) D4 {
	return D4{D4_DC15{Cf23, Cf24}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

type D4_DC16 struct {
	Cf25 _dafny.Set
	Cf26 _dafny.CodePoint
	Cf27 _dafny.Int
	Cf28 _dafny.CodePoint
}

func (D4_DC16) isD4() {}

func (CompanionStruct_D4_) Create_DC16_(Cf25 _dafny.Set, Cf26 _dafny.CodePoint, Cf27 _dafny.Int, Cf28 _dafny.CodePoint) D4 {
	return D4{D4_DC16{Cf25, Cf26, Cf27, Cf28}}
}

func (_this D4) Is_DC16() bool {
	_, ok := _this.Get_().(D4_DC16)
	return ok
}

type D4_DC13 struct {
	Cf22 *C0
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf22 *C0) D4 {
	return D4{D4_DC13{Cf22}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC14_()
}

func (_this D4) Dtor_cf23() bool {
	return _this.Get_().(D4_DC15).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D4_DC15).Cf24
}

func (_this D4) Dtor_cf25() _dafny.Set {
	return _this.Get_().(D4_DC16).Cf25
}

func (_this D4) Dtor_cf26() _dafny.CodePoint {
	return _this.Get_().(D4_DC16).Cf26
}

func (_this D4) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D4_DC16).Cf27
}

func (_this D4) Dtor_cf28() _dafny.CodePoint {
	return _this.Get_().(D4_DC16).Cf28
}

func (_this D4) Dtor_cf22() *C0 {
	return _this.Get_().(D4_DC13).Cf22
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC14:
		{
			return "D4.DC14"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D4_DC16:
		{
			return "D4.DC16" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf22) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC14:
		{
			_, ok := other.Get_().(D4_DC14)
			return ok
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf23 == data2.Cf23 && data1.Cf24.Cmp(data2.Cf24) == 0
		}
	case D4_DC16:
		{
			data2, ok := other.Get_().(D4_DC16)
			return ok && data1.Cf25.Equals(data2.Cf25) && data1.Cf26 == data2.Cf26 && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28 == data2.Cf28
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf22 == data2.Cf22
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
	Cf29 _dafny.Map
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf29 _dafny.Map) D5 {
	return D5{D5_DC17{Cf29}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D5) Dtor_cf29() _dafny.Map {
	return _this.Get_().(D5_DC17).Cf29
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf29) + ")"
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
			return ok && data1.Cf29.Equals(data2.Cf29)
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

type D6_DC19 struct {
	Cf31 _dafny.Int
	Cf32 bool
	Cf33 _dafny.CodePoint
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf31 _dafny.Int, Cf32 bool, Cf33 _dafny.CodePoint) D6 {
	return D6{D6_DC19{Cf31, Cf32, Cf33}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

type D6_DC20 struct {
	Cf34 bool
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_(Cf34 bool) D6 {
	return D6{D6_DC20{Cf34}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

type D6_DC21 struct {
}

func (D6_DC21) isD6() {}

func (CompanionStruct_D6_) Create_DC21_() D6 {
	return D6{D6_DC21{}}
}

func (_this D6) Is_DC21() bool {
	_, ok := _this.Get_().(D6_DC21)
	return ok
}

type D6_DC18 struct {
	Cf30 _dafny.CodePoint
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf30 _dafny.CodePoint) D6 {
	return D6{D6_DC18{Cf30}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC19_(_dafny.Zero, false, _dafny.CodePoint('D'))
}

func (_this D6) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D6_DC19).Cf31
}

func (_this D6) Dtor_cf32() bool {
	return _this.Get_().(D6_DC19).Cf32
}

func (_this D6) Dtor_cf33() _dafny.CodePoint {
	return _this.Get_().(D6_DC19).Cf33
}

func (_this D6) Dtor_cf34() bool {
	return _this.Get_().(D6_DC20).Cf34
}

func (_this D6) Dtor_cf30() _dafny.CodePoint {
	return _this.Get_().(D6_DC18).Cf30
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D6_DC20:
		{
			return "D6.DC20" + "(" + _dafny.String(data.Cf34) + ")"
		}
	case D6_DC21:
		{
			return "D6.DC21"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33
		}
	case D6_DC20:
		{
			data2, ok := other.Get_().(D6_DC20)
			return ok && data1.Cf34 == data2.Cf34
		}
	case D6_DC21:
		{
			_, ok := other.Get_().(D6_DC21)
			return ok
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf30 == data2.Cf30
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

// Definition of class GlobalState
type GlobalState struct {
	F0  _dafny.Int
	F4  _dafny.Int
	_f1 _dafny.Int
	_f2 bool
	_f3 _dafny.Int
	_f5 _dafny.Map
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F4 = _dafny.Zero
	_this._f1 = _dafny.Zero
	_this._f2 = false
	_this._f3 = _dafny.Zero
	_this._f5 = _dafny.EmptyMap
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Int, f2 bool, f3 _dafny.Int, f4 _dafny.Int, f5 _dafny.Map) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F5() _dafny.Map {
	{
		return _this._f5
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f6 _dafny.Int
	_f7 _dafny.Set
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f6 = _dafny.Zero
	_this._f7 = _dafny.EmptySet
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

func (_this *C0) Ctor__(f6 _dafny.Int, f7 _dafny.Set) {
	{
		(_this)._f6 = f6
		(_this)._f7 = f7
	}
}
func (_this *C0) Fm1(p0 _dafny.MultiSet, p1 D0, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt((_this).F6(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(287))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}(func(_698_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))).Cardinality()))
	}
}
func (_this *C0) Fm2(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 _dafny.Map, globalState *GlobalState) _dafny.Int {
	{
		return (Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality()))).Dtor_cf4()
	}
}
func (_this *C0) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *C0) F7() _dafny.Set {
	{
		return _this._f7
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
