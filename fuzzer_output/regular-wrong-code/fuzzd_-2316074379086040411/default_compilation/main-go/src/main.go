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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(573))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.Int {
		return (_dafny.MultiSetOf(false)).Cardinality()
	})), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), true)).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-708)), _dafny.SeqOf(_dafny.IntOfInt64(224), (_dafny.MultiSetOf(!(true))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((func() _dafny.CodePoint {
		if true {
			return _dafny.CodePoint('c')
		}
		return _dafny.CodePoint('h')
	})())
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) D0 {
	var _source0 D5 = Companion_D5_.Create_DC14_(Companion_D5_.Create_DC13_(_dafny.CodePoint('y'), false, _dafny.IntOfInt64(349), _dafny.CodePoint('u')))
	_ = _source0
	if _source0.Is_DC12() {
		var _1___mcc_h0 _dafny.Int = _source0.Get_().(D5_DC12).Cf17
		_ = _1___mcc_h0
		var _2_cf17 _dafny.Int = _1___mcc_h0
		_ = _2_cf17
		return Companion_D0_.Create_DC0_(_dafny.CodePoint('d'))
	} else if _source0.Is_DC13() {
		var _3___mcc_h1 _dafny.CodePoint = _source0.Get_().(D5_DC13).Cf18
		_ = _3___mcc_h1
		var _4___mcc_h2 bool = _source0.Get_().(D5_DC13).Cf19
		_ = _4___mcc_h2
		var _5___mcc_h3 _dafny.Int = _source0.Get_().(D5_DC13).Cf20
		_ = _5___mcc_h3
		var _6___mcc_h4 _dafny.CodePoint = _source0.Get_().(D5_DC13).Cf21
		_ = _6___mcc_h4
		var _7_cf21 _dafny.CodePoint = _6___mcc_h4
		_ = _7_cf21
		var _8_cf20 _dafny.Int = _5___mcc_h3
		_ = _8_cf20
		var _9_cf19 bool = _4___mcc_h2
		_ = _9_cf19
		var _10_cf18 _dafny.CodePoint = _3___mcc_h1
		_ = _10_cf18
		return Companion_D0_.Create_DC0_(_dafny.CodePoint('v'))
	} else if _source0.Is_DC11() {
		var _11___mcc_h5 _dafny.MultiSet = _source0.Get_().(D5_DC11).Cf16
		_ = _11___mcc_h5
		var _12_cf16 _dafny.MultiSet = _11___mcc_h5
		_ = _12_cf16
		return Companion_D0_.Create_DC0_(_dafny.CodePoint('o'))
	} else {
		var _13___mcc_h6 D5 = _source0.Get_().(D5_DC14).Cf22
		_ = _13___mcc_h6
		var _14_cf22 D5 = _13___mcc_h6
		_ = _14_cf22
		return Companion_D0_.Create_DC0_(_dafny.CodePoint('y'))
	}
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Map, globalState *GlobalState) _dafny.Int {
	var _source1 D5 = Companion_D5_.Create_DC11_(_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(588))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_15_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))))
	_ = _source1
	if _source1.Is_DC12() {
		var _16___mcc_h0 _dafny.Int = _source1.Get_().(D5_DC12).Cf17
		_ = _16___mcc_h0
		var _17_cf17 _dafny.Int = _16___mcc_h0
		_ = _17_cf17
		return _17_cf17
	} else if _source1.Is_DC13() {
		var _18___mcc_h1 _dafny.CodePoint = _source1.Get_().(D5_DC13).Cf18
		_ = _18___mcc_h1
		var _19___mcc_h2 bool = _source1.Get_().(D5_DC13).Cf19
		_ = _19___mcc_h2
		var _20___mcc_h3 _dafny.Int = _source1.Get_().(D5_DC13).Cf20
		_ = _20___mcc_h3
		var _21___mcc_h4 _dafny.CodePoint = _source1.Get_().(D5_DC13).Cf21
		_ = _21___mcc_h4
		var _22_cf21 _dafny.CodePoint = _21___mcc_h4
		_ = _22_cf21
		var _23_cf20 _dafny.Int = _20___mcc_h3
		_ = _23_cf20
		var _24_cf19 bool = _19___mcc_h2
		_ = _24_cf19
		var _25_cf18 _dafny.CodePoint = _18___mcc_h1
		_ = _25_cf18
		return _dafny.IntOfUint32((_dafny.SeqOf(_23_cf20)).Cardinality())
	} else if _source1.Is_DC11() {
		var _26___mcc_h5 _dafny.MultiSet = _source1.Get_().(D5_DC11).Cf16
		_ = _26___mcc_h5
		var _27_cf16 _dafny.MultiSet = _26___mcc_h5
		_ = _27_cf16
		return _dafny.IntOfInt64(890)
	} else {
		var _28___mcc_h6 D5 = _source1.Get_().(D5_DC14).Cf22
		_ = _28___mcc_h6
		var _29_cf22 D5 = _28___mcc_h6
		_ = _29_cf22
		return _dafny.IntOfInt64(-862)
	}
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("hab")
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(245), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(478), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(30))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_30_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	}))).Cardinality())))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfInt64(247))))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(true)), _dafny.IntOfInt64(994))).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _31_v0 _dafny.Sequence
			_31_v0 = interface{}(_compr_0).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(true)), _dafny.IntOfInt64(994))).Contains(_31_v0) {
				_coll0.Add(_31_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("af")).Cardinality()))
			}
		}
		return _coll0.ToMap()
	}()).Merge((func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(625))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_32_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf(false)
		})))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _33_v1 _dafny.Sequence
			_33_v1 = interface{}(_compr_1).(_dafny.Sequence)
			if (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(625))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_32_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(false)
			})))).Contains(_33_v1) {
				_coll1.Add(_33_v1, _dafny.IntOfInt64(467))
			}
		}
		return _coll1.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(false), false, false, false, true), _dafny.IntOfInt64(461))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Sequence, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('y')
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false), true, false, false, true), _dafny.SeqOf(false, false)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(!(false))))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(!(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(274))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_34_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	})), _dafny.UnicodeSeqOfUtf8Bytes("fuonoct"))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.MultiSet, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('u'))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _35_v0 _dafny.CodePoint
			_35_v0 = interface{}(_compr_2).(_dafny.CodePoint)
			if (_dafny.SetOf(_dafny.CodePoint('u'))).Contains(_35_v0) {
				_coll2.Add(_35_v0, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(149), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality()))
			}
		}
		return _coll2.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _dafny.IntOfInt64(46)))).Merge((func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), false)).Keys().Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _36_v1 _dafny.CodePoint
			_36_v1 = interface{}(_compr_3).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), false)).Contains(_36_v1) {
				_coll3.Add(_36_v1, (_dafny.MultiSetOf(_dafny.IntOfInt64(72))).Cardinality())
			}
		}
		return _coll3.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(518))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm12(globalState *GlobalState) bool {
	return !(true) || (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(841), _dafny.IntOfInt64(902))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _37_v0 _dafny.Int
			_37_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(841)).Cmp(_37_v0) <= 0) && ((_37_v0).Cmp(_dafny.IntOfInt64(902)) < 0) {
				_coll4.Add((_37_v0).Minus((_dafny.MultiSetOf(true)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-948))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg6 _dafny.Int) interface{} {
						return coer6(arg6)
					}
				}(func(_38_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				}))).Cardinality()))
			}
		}
		return _coll4.ToMap()
	}())).Cardinality()).Cmp(_dafny.IntOfInt64(-711)) >= 0)
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(true, true)).Union((func() _dafny.MultiSet {
		if true {
			return _dafny.MultiSetOf(false)
		}
		return _dafny.MultiSetOf(false, true)
	})())
}
func (_static *CompanionStruct_Default___) M1(p0 _dafny.CodePoint, p1 _dafny.Sequence, p2 bool, p3 _dafny.Sequence, globalState *GlobalState) (*C0, _dafny.Array, _dafny.Int, _dafny.Int) {
	var r0 *C0 = (*C0)(nil)
	_ = r0
	var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	var r3 _dafny.Int = _dafny.Zero
	_ = r3
	var _39_v0 _dafny.Set
	_ = _39_v0
	_39_v0 = _dafny.SetOf(true, Companion_Default___.Fm12(globalState))
	var _40_v1 _dafny.Int
	_ = _40_v1
	_40_v1 = _dafny.IntOfInt64(952)
	var _41_v4 _dafny.Map
	_ = _41_v4
	_41_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(327), _dafny.IntOfInt64(247))); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _42_v2 _dafny.Int
			_42_v2 = interface{}(_compr_5).(_dafny.Int)
			if ((_dafny.IntOfInt64(327)).Cmp(_42_v2) <= 0) && ((_42_v2).Cmp(_dafny.IntOfInt64(247)) < 0) {
				_coll5.Add((_42_v2).Times((func() _dafny.Set {
					var _coll6 = _dafny.NewBuilder()
					_ = _coll6
					for _iter6 := _dafny.Iterate((_dafny.MultiSetFromSeq(p3)).Elements()); ; {
						_compr_6, _ok6 := _iter6()
						if !_ok6 {
							break
						}
						var _43_v3 _dafny.Int
						_43_v3 = interface{}(_compr_6).(_dafny.Int)
						if (_dafny.MultiSetFromSeq(p3)).Contains(_43_v3) {
							_coll6.Add(Companion_Default___.SafeDivisionInt(_43_v3, _dafny.IntOfInt64(665)))
						}
					}
					return _coll6.ToSet()
				}()).Cardinality()))
			}
		}
		return _coll5.ToSet()
	}(), p2)
	var _44_v5 _dafny.Map
	_ = _44_v5
	_44_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v1, (_41_v4).Cardinality())
	var _45_i0 _dafny.Int
	_ = _45_i0
	_45_i0 = _dafny.Zero
	{
		for ((_39_v0).Cardinality()).Cmp((_dafny.Zero).Minus((func() _dafny.Int {
			if (_44_v5).Contains(_40_v1) {
				return (_44_v5).Get(_40_v1).(_dafny.Int)
			}
			return _40_v1
		})())) > 0 {
			{
				if (_45_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_45_i0 = (_45_i0).Plus(_dafny.One)
				(globalState).F2 = _40_v1
				_44_v5 = (_44_v5).Update(_40_v1, Companion_Default___.SafeModuloInt(_40_v1, _dafny.IntOfInt64(196)))
				var _46_v6 D4
				_ = _46_v6
				_46_v6 = Companion_D4_.Create_DC9_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(234))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg7 _dafny.Int) interface{} {
						return coer7(arg7)
					}
				}((func(_47_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_48_i1 _dafny.Int) _dafny.Int {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _47_v1)).Cardinality()
					}
				})(_40_v1)))).Cardinality()), p2)
				var _source2 D4 = (func() D4 {
					if p2 {
						return _46_v6
					}
					return _46_v6
				})()
				_ = _source2
				if _source2.Is_DC9() {
					var _49___mcc_h0 _dafny.Int = _source2.Get_().(D4_DC9).Cf14
					_ = _49___mcc_h0
					var _50___mcc_h1 bool = _source2.Get_().(D4_DC9).Cf15
					_ = _50___mcc_h1
					var _51_cf15 bool = _50___mcc_h1
					_ = _51_cf15
					var _52_cf14 _dafny.Int = _49___mcc_h0
					_ = _52_cf14
					var _53_v7 _dafny.Array
					_ = _53_v7
					var _nwElement0_0 _dafny.Int = _52_cf14
					_ = _nwElement0_0
					var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.One)
					_ = _nw0
					(_nw0).ArraySet1(_nwElement0_0, 0)
					_53_v7 = _nw0
					var _54_v8 _dafny.Sequence
					_ = _54_v8
					_54_v8 = _dafny.UnicodeSeqOfUtf8Bytes("e")
					var _55_v9 _dafny.MultiSet
					_ = _55_v9
					_55_v9 = _dafny.MultiSetOf(_54_v8)
					var _56_v10 D5
					_ = _56_v10
					_56_v10 = Companion_D5_.Create_DC11_(_55_v9)
					var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_53_v7), 0))
					_ = _index0
					(_53_v7).ArraySet1(((_55_v9).Union((_56_v10).Dtor_cf16())).Cardinality(), (_index0).Int())
					var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_53_v7), 0))
					_ = _index1
					(_53_v7).ArraySet1((func() _dafny.Int {
						if (_44_v5).Contains(_52_cf14) {
							return (_44_v5).Get(_52_cf14).(_dafny.Int)
						}
						return _40_v1
					})(), (_index1).Int())
					(globalState).F10 = p2
					var _57_v11 *C0
					_ = _57_v11
					var _nw1 *C0 = New_C0_()
					_ = _nw1
					_nw1.Ctor__()
					_57_v11 = _nw1
					r0 = _57_v11
					var _58_v12 _dafny.Array
					_ = _58_v12
					var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
					_ = _nw2
					_58_v12 = _nw2
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_58_v12), 0))
					_ = _index2
					(_58_v12).ArraySet1(_54_v8, (_index2).Int())
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_58_v12), 0))
					_ = _index3
					(_58_v12).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_54_v8, _54_v8), (_index3).Int())
				} else if _source2.Is_DC10() {
					var _59_v13 _dafny.Set
					_ = _59_v13
					_59_v13 = _dafny.SetOf(_40_v1, _40_v1, _40_v1, _40_v1, _40_v1)
					var _60_v14 _dafny.Map
					_ = _60_v14
					_60_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_59_v13).Cardinality(), p2)
					var _61_v15 D0
					_ = _61_v15
					_61_v15 = Companion_D0_.Create_DC1_(!((func() bool {
						if (_60_v14).Contains(_40_v1) {
							return (_60_v14).Get(_40_v1).(bool)
						}
						return p2
					})()), false, p2)
					var _62_v16 _dafny.Sequence
					_ = _62_v16
					_62_v16 = _dafny.UnicodeSeqOfUtf8Bytes("tgxywina")
					var _63_v17 _dafny.Map
					_ = _63_v17
					_63_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
					r3 = (func() _dafny.Int {
						if (_61_v15).Dtor_cf3() {
							return (func() _dafny.Int {
								if p2 {
									return _dafny.IntOfUint32((_62_v16).Cardinality())
								}
								return _40_v1
							})()
						}
						return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_46_v6).Dtor_cf15())).Merge(_63_v17)).Cardinality()
					})()
					(globalState).F9 = _dafny.IntOfInt64(-423)
					var _64_v18 _dafny.Sequence
					_ = _64_v18
					_64_v18 = _dafny.SeqOf(p2)
					var _65_v19 _dafny.Array
					_ = _65_v19
					var _nwElement0_1 _dafny.Sequence = _64_v18
					_ = _nwElement0_1
					var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(4))
					_ = _nw3
					(_nw3).ArraySet1(_nwElement0_1, 0)
					(_nw3).ArraySet1(_dafny.SeqOf((_64_v18).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_40_v1), _dafny.IntOfUint32((_64_v18).Cardinality()))).Uint32()).(bool)), 1)
					(_nw3).ArraySet1(_dafny.SeqOf(p2, !(p2)), 2)
					(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_64_v18, _64_v18), 3)
					_65_v19 = _nw3
					r1 = _65_v19
					var _66_v20 *C0
					_ = _66_v20
					var _nw4 *C0 = New_C0_()
					_ = _nw4
					_nw4.Ctor__()
					_66_v20 = _nw4
				} else {
					var _67___mcc_h2 _dafny.Map = _source2.Get_().(D4_DC8).Cf13
					_ = _67___mcc_h2
					var _68_cf13 _dafny.Map = _67___mcc_h2
					_ = _68_cf13
					var _69_v21 *C0
					_ = _69_v21
					var _nw5 *C0 = New_C0_()
					_ = _nw5
					_nw5.Ctor__()
					_69_v21 = _nw5
					r0 = (func() *C0 {
						if p2 {
							return _69_v21
						}
						return _69_v21
					})()
					(globalState).F2 = ((_40_v1).Times(_40_v1)).Times((_40_v1).Minus(_dafny.IntOfInt64(-153)))
					(globalState).F9 = _40_v1
					(globalState).F2 = (p3).Select((Companion_Default___.SafeIndex(_40_v1, _dafny.IntOfUint32((p3).Cardinality()))).Uint32()).(_dafny.Int)
				}
				var _70_v22 _dafny.Sequence
				_ = _70_v22
				_70_v22 = _dafny.UnicodeSeqOfUtf8Bytes("cnyebc")
				_40_v1 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_70_v22).Cardinality()), (_dafny.IntOfInt64(-808)).Plus(_dafny.IntOfInt64(133)))
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _71_v23 *C0
	_ = _71_v23
	var _nw6 *C0 = New_C0_()
	_ = _nw6
	_nw6.Ctor__()
	_71_v23 = _nw6
	var _72_v24 _dafny.Array
	_ = _72_v24
	var _nwElement0_2 *C0 = _71_v23
	_ = _nwElement0_2
	var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(8))
	_ = _nw7
	(_nw7).ArraySet1(_nwElement0_2, 0)
	(_nw7).ArraySet1(_71_v23, 1)
	(_nw7).ArraySet1(_71_v23, 2)
	(_nw7).ArraySet1(_71_v23, 3)
	(_nw7).ArraySet1(_71_v23, 4)
	(_nw7).ArraySet1(_71_v23, 5)
	(_nw7).ArraySet1(_71_v23, 6)
	(_nw7).ArraySet1(_71_v23, 7)
	_72_v24 = _nw7
	var _rhs0 _dafny.Int = _40_v1
	_ = _rhs0
	var _rhs1 _dafny.Array = _72_v24
	_ = _rhs1
	var _rhs2 bool = p2
	_ = _rhs2
	var _rhs3 _dafny.Int = _40_v1
	_ = _rhs3
	var _rhs4 *C0 = _71_v23
	_ = _rhs4
	var _lhs0 *GlobalState = globalState
	_ = _lhs0
	var _lhs1 *GlobalState = globalState
	_ = _lhs1
	r3 = _rhs0
	_72_v24 = _rhs1
	_lhs0.F13 = _rhs2
	_lhs1.F2 = _rhs3
	_71_v23 = _rhs4
	var _73_v25 _dafny.Sequence
	_ = _73_v25
	_73_v25 = _dafny.SeqOf(p2)
	var _74_v26 _dafny.Map
	_ = _74_v26
	_74_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v25, _40_v1)
	var _hi0 _dafny.Int = _40_v1
	_ = _hi0
	for _75_i2 := Companion_Default___.Fm4(_74_v26, globalState); _75_i2.Cmp(_hi0) < 0; _75_i2 = _75_i2.Plus(_dafny.One) {
		var _76_v27 *C0
		_ = _76_v27
		var _nw8 *C0 = New_C0_()
		_ = _nw8
		_nw8.Ctor__()
		_76_v27 = _nw8
		(globalState).F10 = p2
		var _77_v28 _dafny.Sequence
		_ = _77_v28
		_77_v28 = _dafny.UnicodeSeqOfUtf8Bytes("dw")
		var _78_v29 _dafny.Map
		_ = _78_v29
		_78_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v1, _77_v28)
		var _79_v30 _dafny.Array
		_ = _79_v30
		var _nwElement0_3 _dafny.Sequence = (func() _dafny.Sequence {
			if p2 {
				return _77_v28
			}
			return _77_v28
		})()
		_ = _nwElement0_3
		var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(6))
		_ = _nw9
		(_nw9).ArraySet1(_nwElement0_3, 0)
		(_nw9).ArraySet1(_77_v28, 1)
		(_nw9).ArraySet1((func() _dafny.Sequence {
			if (_78_v29).Contains(_dafny.IntOfInt64(619)) {
				return (_78_v29).Get(_dafny.IntOfInt64(619)).(_dafny.Sequence)
			}
			return _77_v28
		})(), 2)
		(_nw9).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_77_v28, _77_v28), 3)
		(_nw9).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hicn"), 4)
		(_nw9).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("etaeup"), _77_v28), 5)
		_79_v30 = _nw9
		_79_v30 = _79_v30
		var _80_v31 _dafny.Map
		_ = _80_v31
		_80_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		var _81_v32 _dafny.Map
		_ = _81_v32
		_81_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_80_v31).Contains(p2) {
				return (_80_v31).Get(p2).(bool)
			}
			return p2
		})(), p2)
		var _82_v33 _dafny.Map
		_ = _82_v33
		_82_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, Companion_Default___.Fm10(p2, p2, _40_v1, p0, globalState))
		(globalState).F10 = (func() bool {
			if (_80_v31).Contains(true) {
				return (_80_v31).Get(true).(bool)
			}
			return (_dafny.SetOf(p2, p2, p2, p2, p2)).IsSubsetOf((func() _dafny.Set {
				if (_82_v33).Contains(p2) {
					return (_82_v33).Get(p2).(_dafny.Set)
				}
				return _dafny.SetOf(p2, p2)
			})())
		})()
	}
	var _83_v34 _dafny.CodePoint
	_ = _83_v34
	_83_v34 = _dafny.CodePoint('s')
	_83_v34 = p0
	var _84_v35 _dafny.Array
	_ = _84_v35
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(23)
	_ = _len0_0
	var _nw10 _dafny.Array
	_ = _nw10
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw10 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Int = (func(_85_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_86_i4 _dafny.Int) _dafny.Int {
				return (_86_i4).Minus(_85_v1)
			}
		})(_40_v1)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw10 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw10).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw10).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_84_v35 = _nw10
	var _87_v36 D1
	_ = _87_v36
	_87_v36 = Companion_D1_.Create_DC3_(p3, p2)
	var _88_v37 _dafny.Map
	_ = _88_v37
	_88_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v35, (_87_v36).Dtor_cf5())
	var _89_v38 _dafny.Map
	_ = _89_v38
	_89_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfInt64(-464))
	var _90_v39 _dafny.Sequence
	_ = _90_v39
	_90_v39 = _dafny.SeqOf(_89_v38)
	var _91_v40 _dafny.Array
	_ = _91_v40
	var _nwElement0_4 bool = p2
	_ = _nwElement0_4
	var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(28))
	_ = _nw11
	(_nw11).ArraySet1(_nwElement0_4, 0)
	(_nw11).ArraySet1(p2, 1)
	(_nw11).ArraySet1(!(p2), 2)
	(_nw11).ArraySet1(p2, 3)
	(_nw11).ArraySet1(Companion_Default___.Fm12(globalState), 4)
	(_nw11).ArraySet1(false, 5)
	(_nw11).ArraySet1(false, 6)
	(_nw11).ArraySet1(p2, 7)
	(_nw11).ArraySet1(!(p2), 8)
	(_nw11).ArraySet1(true, 9)
	(_nw11).ArraySet1(Companion_Default___.Fm12(globalState), 10)
	(_nw11).ArraySet1((_88_v37).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v35, p3)), 11)
	(_nw11).ArraySet1(p2, 12)
	(_nw11).ArraySet1(p2, 13)
	(_nw11).ArraySet1((_71_v23).Fm3(!(p2), p2, Companion_Default___.Fm13(globalState), _73_v25, globalState), 14)
	(_nw11).ArraySet1(p2, 15)
	(_nw11).ArraySet1(p2, 16)
	(_nw11).ArraySet1(p2, 17)
	(_nw11).ArraySet1(p2, 18)
	(_nw11).ArraySet1(true, 19)
	(_nw11).ArraySet1(p2, 20)
	(_nw11).ArraySet1(p2, 21)
	(_nw11).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_73_v25, _dafny.SeqOf(p2, p2)), 22)
	(_nw11).ArraySet1(p2, 23)
	(_nw11).ArraySet1(p2, 24)
	(_nw11).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_90_v39, _90_v39), 25)
	(_nw11).ArraySet1(p2, 26)
	(_nw11).ArraySet1((_40_v1).Cmp(Companion_Default___.Fm4(_74_v26, globalState)) < 0, 27)
	_91_v40 = _nw11
	for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_91_v40), 0))); ; {
		_guard_loop_0, _ok7 := _iter7()
		if !_ok7 {
			break
		}
		var _92_i3 _dafny.Int
		_92_i3 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_92_i3).Sign() != -1) && ((_92_i3).Cmp(_dafny.ArrayLen((_91_v40), 0)) < 0)) {
			(_91_v40).ArraySet1((_dafny.MultiSetOf(_40_v1)).IsSubsetOf((func() _dafny.MultiSet {
				if true {
					return _dafny.MultiSetOf(_40_v1)
				}
				return _dafny.MultiSetOf(_40_v1, _dafny.IntOfUint32((_dafny.SeqOf(_40_v1)).Cardinality()), _40_v1, _40_v1, _40_v1)
			})()), (_92_i3).Int())
		}
	}
	var _93_v41 _dafny.Array
	_ = _93_v41
	var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
	_ = _nw12
	_93_v41 = _nw12
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_93_v41), 0))
	_ = _index4
	(_93_v41).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfInt64(842)), (_index4).Int())
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_93_v41), 0))
	_ = _index5
	(_93_v41).ArraySet1((_89_v38).Merge((_89_v38).Merge(_89_v38)), (_index5).Int())
	var _nw13 *C0 = New_C0_()
	_ = _nw13
	_nw13.Ctor__()
	r0 = _nw13
	var _94_v42 _dafny.Array
	_ = _94_v42
	var _nwElement0_5 _dafny.Sequence = _73_v25
	_ = _nwElement0_5
	var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(14))
	_ = _nw14
	(_nw14).ArraySet1(_nwElement0_5, 0)
	(_nw14).ArraySet1(_dafny.SeqOf(p2), 1)
	(_nw14).ArraySet1(_73_v25, 2)
	(_nw14).ArraySet1(_73_v25, 3)
	(_nw14).ArraySet1(_73_v25, 4)
	(_nw14).ArraySet1(Companion_Default___.Fm9(globalState), 5)
	(_nw14).ArraySet1(_73_v25, 6)
	(_nw14).ArraySet1(_73_v25, 7)
	(_nw14).ArraySet1(_73_v25, 8)
	(_nw14).ArraySet1(Companion_Default___.Fm9(globalState), 9)
	(_nw14).ArraySet1(_73_v25, 10)
	(_nw14).ArraySet1(_73_v25, 11)
	(_nw14).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_73_v25, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_40_v1), _dafny.IntOfUint32((_73_v25).Cardinality()))).Uint32(), Companion_Default___.Fm12(globalState)), (Companion_Default___.SafeIndex(_40_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_73_v25, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_40_v1), _dafny.IntOfUint32((_73_v25).Cardinality()))).Uint32(), Companion_Default___.Fm12(globalState))).Cardinality()))).Uint32(), !(p2)), 12)
	(_nw14).ArraySet1(_73_v25, 13)
	_94_v42 = _nw14
	var _95_v43 _dafny.Map
	_ = _95_v43
	_95_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_v23, _94_v42)
	r1 = (func() _dafny.Array {
		if (_95_v43).Contains(_71_v23) {
			return (_95_v43).Get(_71_v23).(_dafny.Array)
		}
		return _94_v42
	})()
	r2 = Companion_Default___.SafeModuloInt(_40_v1, _40_v1)
	r3 = _40_v1
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _96_v0 _dafny.Array
	_ = _96_v0
	var _nw15 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
	_ = _nw15
	_96_v0 = _nw15
	var _97_v1 _dafny.Sequence
	_ = _97_v1
	_97_v1 = _dafny.UnicodeSeqOfUtf8Bytes("uwjfy")
	var _98_globalState *GlobalState
	_ = _98_globalState
	var _nw16 *GlobalState = New_GlobalState_()
	_ = _nw16
	_nw16.Ctor__(_dafny.IntOfInt64(511), _dafny.CodePoint('o'), _dafny.IntOfInt64(-937), _dafny.IntOfInt64(29), _96_v0, false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(220))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_99_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(402))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_100_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	})), _dafny.CodePoint('y'), _dafny.IntOfInt64(949), false, _dafny.CodePoint('j'), _dafny.IntOfInt64(820), true, _dafny.Companion_Sequence_.Concatenate(_97_v1, _97_v1), false)
	_98_globalState = _nw16
	var _101_v2 bool
	_ = _101_v2
	_101_v2 = false
	var _102_v3 _dafny.Int
	_ = _102_v3
	_102_v3 = _dafny.IntOfInt64(397)
	if ((_101_v2) == (_101_v2)) == (!((_102_v3).Cmp(_102_v3) >= 0)) {
		var _103_v4 _dafny.Sequence
		_ = _103_v4
		_103_v4 = _dafny.SeqOf(_102_v3, _dafny.IntOfInt64(92), _102_v3, _102_v3, (_dafny.MultiSetOf(_101_v2)).Cardinality())
		(_98_globalState).F0 = (_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm0(_103_v4, Companion_Default___.SafeModuloInt(_102_v3, _102_v3), _98_globalState)).Cardinality()))
		var _104_v5 _dafny.MultiSet
		_ = _104_v5
		_104_v5 = _dafny.MultiSetOf(_dafny.IntOfInt64(-498), _102_v3)
		(_98_globalState).F10 = !(((_104_v5).Union(_104_v5)).IsSubsetOf(_104_v5))
		if _101_v2 {
			(_98_globalState).F7 = _97_v1
			var _105_v6 _dafny.Map
			_ = _105_v6
			_105_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if _101_v2 {
					return _dafny.IntOfInt64(241)
				}
				return (_104_v5).Cardinality()
			})(), Companion_Default___.Fm1(_102_v3, _98_globalState))
			var _106_v7 _dafny.CodePoint
			_ = _106_v7
			_106_v7 = _dafny.CodePoint('a')
			var _107_v8 _dafny.MultiSet
			_ = _107_v8
			_107_v8 = _dafny.MultiSetOf((Companion_Default___.Fm2(_101_v2, _101_v2, _102_v3, _98_globalState)).Dtor_cf0(), _106_v7, _dafny.CodePoint('t'), _106_v7)
			_105_v6 = (_105_v6).Update(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(96))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}(func(_108_i2 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(936)
			}))).Cardinality()), _107_v8)
			var _109_v9 *C0
			_ = _109_v9
			var _nw17 *C0 = New_C0_()
			_ = _nw17
			_nw17.Ctor__()
			_109_v9 = _nw17
			var _110_v10 _dafny.Array
			_ = _110_v10
			var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw18
			_110_v10 = _nw18
			var _111_v11 _dafny.Map
			_ = _111_v11
			_111_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _110_v10)
			var _112_v12 _dafny.Sequence
			_ = _112_v12
			_112_v12 = _dafny.SeqOf(_110_v10, (func() _dafny.Array {
				if (_111_v11).Contains(_101_v2) {
					return (_111_v11).Get(_101_v2).(_dafny.Array)
				}
				return _110_v10
			})(), _110_v10)
			(_98_globalState).F9 = _dafny.IntOfUint32((_112_v12).Cardinality())
			var _113_v13 *C0
			_ = _113_v13
			var _nw19 *C0 = New_C0_()
			_ = _nw19
			_nw19.Ctor__()
			_113_v13 = _nw19
		} else {
			var _114_v14 _dafny.CodePoint
			_ = _114_v14
			_114_v14 = _dafny.CodePoint('l')
			var _115_v15 D0
			_ = _115_v15
			_115_v15 = Companion_D0_.Create_DC0_(_114_v14)
			var _116_v16 _dafny.Sequence
			_ = _116_v16
			_116_v16 = _dafny.SeqOf(_115_v15)
			var _117_v17 *C0
			_ = _117_v17
			var _118_v18 _dafny.Array
			_ = _118_v18
			var _119_v19 _dafny.Int
			_ = _119_v19
			var _120_v20 _dafny.Int
			_ = _120_v20
			var _out0 *C0
			_ = _out0
			var _out1 _dafny.Array
			_ = _out1
			var _out2 _dafny.Int
			_ = _out2
			var _out3 _dafny.Int
			_ = _out3
			_out0, _out1, _out2, _out3 = Companion_Default___.M1(_dafny.CodePoint('j'), _116_v16, _101_v2, _dafny.Companion_Sequence_.Concatenate(_103_v4, _dafny.SeqOf(_102_v3)), _98_globalState)
			_117_v17 = _out0
			_118_v18 = _out1
			_119_v19 = _out2
			_120_v20 = _out3
			(_98_globalState).F10 = _101_v2
			var _121_v21 _dafny.Map
			_ = _121_v21
			_121_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), _101_v2)
			var _122_v22 _dafny.Set
			_ = _122_v22
			_122_v22 = _dafny.SetOf(_121_v21)
			var _123_v23 _dafny.Sequence
			_ = _123_v23
			_123_v23 = _dafny.SeqOf(_97_v1, _dafny.UnicodeSeqOfUtf8Bytes("ckdgwau"), _dafny.UnicodeSeqOfUtf8Bytes("vu"))
			var _124_v24 _dafny.Map
			_ = _124_v24
			_124_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _123_v23)
			var _125_v25 _dafny.Array
			_ = _125_v25
			var _nwElement0_6 _dafny.Int = (_122_v22).Cardinality()
			_ = _nwElement0_6
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(29))
			_ = _nw20
			(_nw20).ArraySet1(_nwElement0_6, 0)
			(_nw20).ArraySet1(_119_v19, 1)
			(_nw20).ArraySet1(_119_v19, 2)
			(_nw20).ArraySet1(_119_v19, 3)
			(_nw20).ArraySet1((_dafny.Zero).Minus(_120_v20), 4)
			(_nw20).ArraySet1(_102_v3, 5)
			(_nw20).ArraySet1(_dafny.IntOfUint32((_97_v1).Cardinality()), 6)
			(_nw20).ArraySet1(_119_v19, 7)
			(_nw20).ArraySet1(_119_v19, 8)
			(_nw20).ArraySet1(_102_v3, 9)
			(_nw20).ArraySet1(_120_v20, 10)
			(_nw20).ArraySet1(_119_v19, 11)
			(_nw20).ArraySet1(_102_v3, 12)
			(_nw20).ArraySet1(_120_v20, 13)
			(_nw20).ArraySet1(_dafny.IntOfInt64(597), 14)
			(_nw20).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_124_v24).Contains(_101_v2) {
					return (_124_v24).Get(_101_v2).(_dafny.Sequence)
				}
				return _123_v23
			})()).Cardinality()), 15)
			(_nw20).ArraySet1(_102_v3, 16)
			(_nw20).ArraySet1(_102_v3, 17)
			(_nw20).ArraySet1(_119_v19, 18)
			(_nw20).ArraySet1((_dafny.Zero).Minus(_120_v20), 19)
			(_nw20).ArraySet1(_102_v3, 20)
			(_nw20).ArraySet1(_119_v19, 21)
			(_nw20).ArraySet1(_120_v20, 22)
			(_nw20).ArraySet1(_102_v3, 23)
			(_nw20).ArraySet1(_119_v19, 24)
			(_nw20).ArraySet1((_dafny.MultiSetOf(_101_v2)).Cardinality(), 25)
			(_nw20).ArraySet1(_119_v19, 26)
			(_nw20).ArraySet1(_120_v20, 27)
			(_nw20).ArraySet1(_dafny.IntOfInt64(291), 28)
			_125_v25 = _nw20
			var _126_v26 _dafny.Array
			_ = _126_v26
			var _nwElement0_7 _dafny.Array = _125_v25
			_ = _nwElement0_7
			var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(26))
			_ = _nw21
			(_nw21).ArraySet1(_nwElement0_7, 0)
			(_nw21).ArraySet1(_125_v25, 1)
			(_nw21).ArraySet1(_125_v25, 2)
			(_nw21).ArraySet1(_125_v25, 3)
			(_nw21).ArraySet1(_125_v25, 4)
			(_nw21).ArraySet1(_125_v25, 5)
			(_nw21).ArraySet1(_125_v25, 6)
			(_nw21).ArraySet1(_125_v25, 7)
			(_nw21).ArraySet1(_125_v25, 8)
			(_nw21).ArraySet1(_125_v25, 9)
			(_nw21).ArraySet1(_125_v25, 10)
			(_nw21).ArraySet1(_125_v25, 11)
			(_nw21).ArraySet1(_125_v25, 12)
			(_nw21).ArraySet1(_125_v25, 13)
			(_nw21).ArraySet1(_125_v25, 14)
			(_nw21).ArraySet1(_125_v25, 15)
			(_nw21).ArraySet1(_125_v25, 16)
			(_nw21).ArraySet1(_125_v25, 17)
			(_nw21).ArraySet1(_125_v25, 18)
			(_nw21).ArraySet1(_125_v25, 19)
			(_nw21).ArraySet1(_125_v25, 20)
			(_nw21).ArraySet1(_125_v25, 21)
			(_nw21).ArraySet1(_125_v25, 22)
			(_nw21).ArraySet1(_125_v25, 23)
			(_nw21).ArraySet1(_125_v25, 24)
			(_nw21).ArraySet1(_125_v25, 25)
			_126_v26 = _nw21
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_126_v26), 0))
			_ = _index6
			(_126_v26).ArraySet1(_125_v25, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_126_v26), 0))
			_ = _index7
			(_126_v26).ArraySet1(_125_v25, (_index7).Int())
			var _127_v27 _dafny.Array
			_ = _127_v27
			var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
			_ = _nw22
			_127_v27 = _nw22
			var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw23
			_127_v27 = _nw23
			var _128_v28 _dafny.Sequence
			_ = _128_v28
			_128_v28 = _dafny.SeqOf(_101_v2)
			(_98_globalState).F0 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_97_v1).Cardinality()), _dafny.IntOfUint32((_128_v28).Cardinality()))))
		}
		var _129_v29 _dafny.CodePoint
		_ = _129_v29
		_129_v29 = _dafny.CodePoint('d')
		var _130_v30 D0
		_ = _130_v30
		_130_v30 = Companion_D0_.Create_DC0_(_129_v29)
		var _131_v31 _dafny.Sequence
		_ = _131_v31
		_131_v31 = _dafny.SeqOf(_130_v30, _130_v30, Companion_D0_.Create_DC0_(_129_v29))
		var _132_v32 _dafny.Map
		_ = _132_v32
		_132_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v3, _102_v3)
		var _133_v33 *C0
		_ = _133_v33
		var _134_v34 _dafny.Array
		_ = _134_v34
		var _135_v35 _dafny.Int
		_ = _135_v35
		var _136_v36 _dafny.Int
		_ = _136_v36
		var _out4 *C0
		_ = _out4
		var _out5 _dafny.Array
		_ = _out5
		var _out6 _dafny.Int
		_ = _out6
		var _out7 _dafny.Int
		_ = _out7
		_out4, _out5, _out6, _out7 = Companion_Default___.M1(_129_v29, _131_v31, _101_v2, _dafny.SeqOf((_132_v32).Cardinality(), _102_v3), _98_globalState)
		_133_v33 = _out4
		_134_v34 = _out5
		_135_v35 = _out6
		_136_v36 = _out7
		var _137_v37 _dafny.Array
		_ = _137_v37
		var _len0_1 _dafny.Int = _dafny.One
		_ = _len0_1
		var _nw24 _dafny.Array
		_ = _nw24
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw24 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Sequence = (func(_138_v29 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
				return func(_139_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(168))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg11 _dafny.Int) interface{} {
							return coer11(arg11)
						}
					}((func(_140_v29 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_141_i4 _dafny.Int) _dafny.CodePoint {
							return _140_v29
						}
					})(_138_v29)))
				}
			})(_129_v29)
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
		_137_v37 = _nw24
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_137_v37), 0))
		_ = _index8
		(_137_v37).ArraySet1(_97_v1, (_index8).Int())
		var _142_v38 _dafny.Sequence
		_ = _142_v38
		_142_v38 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("alnbjixe"), _97_v1)
		var _143_v39 _dafny.Map
		_ = _143_v39
		_143_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _101_v2)
		var _144_v40 _dafny.Map
		_ = _144_v40
		_144_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_143_v39).Cardinality(), _97_v1)
		var _145_v41 _dafny.Sequence
		_ = _145_v41
		_145_v41 = _dafny.SeqOf(_101_v2, _101_v2)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_137_v37), 0))
		_ = _index9
		var _rhs5 bool = _dafny.Companion_Sequence_.IsPrefixOf((_142_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.IntOfUint32((_142_v38).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5(_101_v2, _101_v2, _98_globalState), _dafny.Companion_Sequence_.Update(_97_v1, (Companion_Default___.SafeIndex((_104_v5).Cardinality(), _dafny.IntOfUint32((_97_v1).Cardinality()))).Uint32(), _129_v29)))
		_ = _rhs5
		var _rhs6 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-282))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_146_v29 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_147_i5 _dafny.Int) _dafny.CodePoint {
				return _146_v29
			}
		})(_129_v29))), (func() _dafny.Sequence {
			if (_144_v40).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_145_v41).Cardinality()))) {
				return (_144_v40).Get((_dafny.Zero).Minus(_dafny.IntOfUint32((_145_v41).Cardinality()))).(_dafny.Sequence)
			}
			return _97_v1
		})())
		_ = _rhs6
		var _rhs7 _dafny.Sequence = _97_v1
		_ = _rhs7
		var _lhs2 *GlobalState = _98_globalState
		_ = _lhs2
		var _lhs3 *GlobalState = _98_globalState
		_ = _lhs3
		var _lhs4 _dafny.Array = _137_v37
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_137_v37), 0))
		_ = _lhs5
		_lhs2.F15 = _rhs5
		_lhs3.F7 = _rhs6
		(_lhs4).ArraySet1(_rhs7, (_lhs5).Int())
	} else {
		var _148_v42 _dafny.Map
		_ = _148_v42
		_148_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _dafny.IntOfInt64(726))
		var _149_v43 _dafny.Sequence
		_ = _149_v43
		_149_v43 = _dafny.SeqOf((func() _dafny.Int {
			if (_148_v42).Contains(_101_v2) {
				return (_148_v42).Get(_101_v2).(_dafny.Int)
			}
			return _dafny.IntOfInt64(-708)
		})(), _dafny.IntOfUint32((_97_v1).Cardinality()), _dafny.IntOfInt64(-797), _102_v3)
		var _150_v44 _dafny.Sequence
		_ = _150_v44
		_150_v44 = _dafny.SeqOf(_149_v43, _149_v43, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(671))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}((func(_151_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_152_i6 _dafny.Int) _dafny.Int {
				return _151_v3
			}
		})(_102_v3))))
		_150_v44 = _150_v44
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_96_v0), 0))
		_ = _index10
		(_96_v0).ArraySet1(_101_v2, (_index10).Int())
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_96_v0), 0))
		_ = _index11
		(_96_v0).ArraySet1(true, (_index11).Int())
		var _153_v45 D1
		_ = _153_v45
		_153_v45 = Companion_D1_.Create_DC4_((_96_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_96_v0), 0))).Int()).(bool), _102_v3, _97_v1)
		var _154_v46 _dafny.Set
		_ = _154_v46
		_154_v46 = _dafny.SetOf(_153_v45, _153_v45, _153_v45)
		_154_v46 = _154_v46
		var _155_v47 _dafny.MultiSet
		_ = _155_v47
		_155_v47 = _dafny.MultiSetOf(_101_v2)
		(_98_globalState).F0 = (((_155_v47).Union(_155_v47)).Cardinality()).Times(_102_v3)
		var _156_v48 _dafny.Map
		_ = _156_v48
		_156_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_148_v42).Cardinality(), _102_v3)
		_156_v48 = (_156_v48).Update(_102_v3, _102_v3)
	}
	if _101_v2 {
		(_98_globalState).F13 = ((func() bool {
			if false {
				return _101_v2
			}
			return _101_v2
		})()) == (false)
		var _157_v49 *C0
		_ = _157_v49
		var _nw25 *C0 = New_C0_()
		_ = _nw25
		_nw25.Ctor__()
		_157_v49 = _nw25
		var _158_v50 _dafny.CodePoint
		_ = _158_v50
		_158_v50 = _dafny.CodePoint('f')
		_158_v50 = _158_v50
		var _159_v51 _dafny.Set
		_ = _159_v51
		_159_v51 = _dafny.SetOf(_101_v2)
		var _160_v52 _dafny.Sequence
		_ = _160_v52
		_160_v52 = _dafny.SeqOf(_102_v3, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _102_v3)).Cardinality(), (_159_v51).Cardinality(), _dafny.IntOfInt64(298))
		var _161_v53 D1
		_ = _161_v53
		_161_v53 = Companion_D1_.Create_DC3_(_160_v52, _101_v2)
		if !((_161_v53).Dtor_cf6()) {
			var _162_v54 _dafny.Sequence
			_ = _162_v54
			_162_v54 = _dafny.SeqOf(true)
			(_98_globalState).F10 = ((_102_v3).Times(_102_v3)).Cmp(Companion_Default___.Fm4(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v54, _dafny.IntOfInt64(904)), _98_globalState)) == 0
			(_98_globalState).F10 = _101_v2
			var _163_v55 _dafny.Array
			_ = _163_v55
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
			_ = _nw26
			_163_v55 = _nw26
			var _164_v56 _dafny.Sequence
			_ = _164_v56
			_164_v56 = _dafny.SeqOf(_157_v49)
			var _165_v57 _dafny.Map
			_ = _165_v57
			_165_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v3, _dafny.IntOfUint32((_164_v56).Cardinality()))
			var _166_v58 D1
			_ = _166_v58
			_166_v58 = Companion_D1_.Create_DC4_(_101_v2, _102_v3, _97_v1)
			var _167_v59 _dafny.Sequence
			_ = _167_v59
			_167_v59 = _dafny.SeqOf(_165_v57)
			var _168_v60 _dafny.Map
			_ = _168_v60
			_168_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _157_v49)
			var _169_v61 _dafny.Map
			_ = _169_v61
			_169_v61 = _165_v57
			var _170_v62 _dafny.Map
			_ = _170_v62
			_170_v62 = (_169_v61)
			var _nwElement0_8 _dafny.Map = _165_v57
			_ = _nwElement0_8
			var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(17))
			_ = _nw27
			(_nw27).ArraySet1(_nwElement0_8, 0)
			(_nw27).ArraySet1(_165_v57, 1)
			(_nw27).ArraySet1(_165_v57, 2)
			(_nw27).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v3, (_166_v58).Dtor_cf8()), 3)
			(_nw27).ArraySet1(((_165_v57).Update(_102_v3, _102_v3)).Merge(_165_v57), 4)
			(_nw27).ArraySet1(Companion_Default___.Fm6(false, _101_v2, _102_v3, _dafny.IntOfInt64(702), _98_globalState), 5)
			(_nw27).ArraySet1((_165_v57).Merge((_167_v59).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-488), _dafny.IntOfUint32((_167_v59).Cardinality()))).Uint32()).(_dafny.Map)), 6)
			(_nw27).ArraySet1(_165_v57, 7)
			(_nw27).ArraySet1((_165_v57).Update(_102_v3, (_168_v60).Cardinality()), 8)
			(_nw27).ArraySet1((_169_v61), 9)
			(_nw27).ArraySet1(_165_v57, 10)
			(_nw27).ArraySet1(func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-838), _dafny.IntOfInt64(-263))); ; {
					_compr_7, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _171_v63 _dafny.Int
					_171_v63 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(-838)).Cmp(_171_v63) <= 0) && ((_171_v63).Cmp(_dafny.IntOfInt64(-263)) < 0) {
						_coll7.Add((_171_v63).Minus(_102_v3), _102_v3)
					}
				}
				return _coll7.ToMap()
			}(), 11)
			(_nw27).ArraySet1(_165_v57, 12)
			(_nw27).ArraySet1((_165_v57).Update(_102_v3, _102_v3), 13)
			(_nw27).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(446), _102_v3), 14)
			(_nw27).ArraySet1(_165_v57, 15)
			(_nw27).ArraySet1(func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter9 := _dafny.Iterate((_160_v52).Elements()); ; {
					_compr_8, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _172_v64 _dafny.Int
					_172_v64 = interface{}(_compr_8).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_160_v52, _172_v64) {
						_coll8.Add(Companion_Default___.SafeModuloInt(_172_v64, _102_v3), _dafny.IntOfInt64(991))
					}
				}
				return _coll8.ToMap()
			}(), 16)
			_163_v55 = _nw27
			var _173_v65 _dafny.Map
			_ = _173_v65
			_173_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_160_v52, (Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_160_v52).Cardinality()))).Uint32(), _102_v3)).Cardinality()))
			var _174_v66 _dafny.MultiSet
			_ = _174_v66
			_174_v66 = _dafny.MultiSetOf(_101_v2)
			_173_v65 = (_173_v65).Update(_dafny.Companion_Sequence_.IsPrefixOf(_97_v1, _97_v1), (func() _dafny.Int {
				if !((_157_v49).Fm3(_101_v2, false, _174_v66, _dafny.SeqOf(_101_v2), _98_globalState)) {
					return _102_v3
				}
				return _102_v3
			})())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_96_v0), 0))
			_ = _index12
			(_96_v0).ArraySet1(_101_v2, (_index12).Int())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_96_v0), 0))
			_ = _index13
			(_96_v0).ArraySet1(false, (_index13).Int())
		} else {
			_96_v0 = _96_v0
			var _175_v67 *C0
			_ = _175_v67
			var _nw28 *C0 = New_C0_()
			_ = _nw28
			_nw28.Ctor__()
			_175_v67 = _nw28
			(_98_globalState).F0 = _102_v3
			(_98_globalState).F10 = _101_v2
			_102_v3 = (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if _101_v2 {
					return _102_v3
				}
				return _dafny.IntOfInt64(421)
			})()))
		}
		(_98_globalState).F15 = _101_v2
	} else {
		(_98_globalState).F2 = _102_v3
		var _176_v68 _dafny.Sequence
		_ = _176_v68
		_176_v68 = _dafny.SeqOf(true)
		var _177_v69 _dafny.Map
		_ = _177_v69
		_177_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_101_v2, _101_v2), _176_v68), _101_v2)
		_177_v69 = _177_v69
		var _178_v70 _dafny.MultiSet
		_ = _178_v70
		_178_v70 = _dafny.MultiSetOf(_101_v2, _101_v2)
		var _179_v71 _dafny.Array
		_ = _179_v71
		var _nwElement0_9 _dafny.MultiSet = _178_v70
		_ = _nwElement0_9
		var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(4))
		_ = _nw29
		(_nw29).ArraySet1(_nwElement0_9, 0)
		(_nw29).ArraySet1((_178_v70).Difference(_178_v70), 1)
		(_nw29).ArraySet1(_178_v70, 2)
		(_nw29).ArraySet1((_178_v70).Update(_101_v2, Companion_Default___.Abs(_102_v3)), 3)
		_179_v71 = _nw29
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_179_v71), 0))
		_ = _index14
		(_179_v71).ArraySet1((_178_v70).Intersection(_178_v70), (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_179_v71), 0))
		_ = _index15
		(_179_v71).ArraySet1((_dafny.MultiSetOf(_101_v2)).Union(_178_v70), (_index15).Int())
		var _180_v72 *C0
		_ = _180_v72
		var _nw30 *C0 = New_C0_()
		_ = _nw30
		_nw30.Ctor__()
		_180_v72 = _nw30
		var _181_v73 *C0
		_ = _181_v73
		_181_v73 = _180_v72
		_180_v72 = (func() *C0 {
			if _101_v2 {
				return _180_v72
			}
			return (func() *C0 {
				if _101_v2 {
					return (_181_v73)
				}
				return _180_v72
			})()
		})()
		if _101_v2 {
			(_180_v72).M0(_176_v68, _102_v3, _98_globalState)
			var _182_v74 _dafny.Array
			_ = _182_v74
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_2
			var _nw31 _dafny.Array
			_ = _nw31
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw31 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Int = func(_183_i7 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_183_i7, (_dafny.SetOf(_dafny.CodePoint('y'))).Cardinality())
				}
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw31 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw31).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw31).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_182_v74 = _nw31
			var _184_v75 _dafny.Map
			_ = _184_v75
			_184_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm5(_101_v2, true, _98_globalState), _182_v74)
			var _185_v76 _dafny.Sequence
			_ = _185_v76
			_185_v76 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v1, _182_v74), (_184_v75).Update(_97_v1, _182_v74), _184_v75)
			var _186_v77 _dafny.Sequence
			_ = _186_v77
			_186_v77 = _dafny.SeqOf(_102_v3, _dafny.IntOfInt64(510), _dafny.IntOfInt64(-366))
			var _187_v78 _dafny.Map
			_ = _187_v78
			_187_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_185_v76).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_97_v1).Cardinality()), _dafny.IntOfUint32((_185_v76).Cardinality()))).Uint32()).(_dafny.Map), Companion_Default___.SafeDivisionInt(_102_v3, (_186_v77).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.SetOf(_101_v2, _101_v2)).Cardinality()), _dafny.IntOfUint32((_186_v77).Cardinality()))).Uint32()).(_dafny.Int)))
			_187_v78 = (_187_v78).Update((_184_v75).Merge(_184_v75), _dafny.IntOfInt64(-798))
			(_98_globalState).F13 = _101_v2
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_182_v74), 0))
			_ = _index16
			(_182_v74).ArraySet1(_102_v3, (_index16).Int())
			var _188_v79 _dafny.MultiSet
			_ = _188_v79
			_188_v79 = _dafny.MultiSetOf(_102_v3, _102_v3)
			var _189_v80 _dafny.Map
			_ = _189_v80
			_189_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v72, _188_v79)
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_182_v74), 0))
			_ = _index17
			(_182_v74).ArraySet1((_dafny.IntOfInt64(618)).Plus(((_189_v80).Cardinality()).Times(_102_v3)), (_index17).Int())
			var _190_v81 _dafny.Set
			_ = _190_v81
			_190_v81 = _dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_101_v2, _101_v2))).Cardinality())
			var _191_v82 _dafny.Map
			_ = _191_v82
			_191_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, (_182_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_182_v74), 0))).Int()).(_dafny.Int))
			var _192_v83 _dafny.Map
			_ = _192_v83
			_192_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_190_v81).IsSubsetOf(_190_v81), (_97_v1).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm4(Companion_Default___.Fm7(_dafny.IntOfUint32((Companion_Default___.Fm5(true, false, _98_globalState)).Cardinality()), (_191_v82).Cardinality(), _98_globalState), _98_globalState), _dafny.IntOfUint32((_97_v1).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _193_v84 _dafny.Array
			_ = _193_v84
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_3
			var _nw32 _dafny.Array
			_ = _nw32
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw32 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.CodePoint = func(_194_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('q')
				}
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw32 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw32).ArraySet1CodePoint(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw32).ArraySet1CodePoint(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_193_v84 = _nw32
			var _195_v85 _dafny.MultiSet
			_ = _195_v85
			_195_v85 = _dafny.MultiSetOf(_193_v84)
			var _196_v86 _dafny.Sequence
			_ = _196_v86
			_196_v86 = _dafny.SeqOf(_195_v85, _195_v85)
			var _197_v87 _dafny.Sequence
			_ = _197_v87
			_197_v87 = _dafny.SeqOf(_193_v84)
			var _198_v88 _dafny.CodePoint
			_ = _198_v88
			_198_v88 = _dafny.CodePoint('s')
			_192_v83 = (_192_v83).Update(((_196_v86).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.IntOfUint32((_196_v86).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsDisjointFrom(_dafny.MultiSetOf((_197_v87).Select((Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_197_v87).Cardinality()))).Uint32()).(_dafny.Array), _193_v84, _193_v84)), _198_v88)
		} else {
			var _199_v89 _dafny.Sequence
			_ = _199_v89
			_199_v89 = _dafny.SeqOf(_102_v3)
			var _200_v90 _dafny.MultiSet
			_ = _200_v90
			_200_v90 = _dafny.MultiSetOf(_102_v3)
			var _201_v91 _dafny.Map
			_ = _201_v91
			_201_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_101_v2), _102_v3)
			var _202_v92 _dafny.CodePoint
			_ = _202_v92
			_202_v92 = _dafny.CodePoint('r')
			var _203_v93 _dafny.MultiSet
			_ = _203_v93
			_203_v93 = _dafny.MultiSetOf(Companion_Default___.Fm8(_199_v89, (_200_v90).Update(_dafny.IntOfInt64(-538), Companion_Default___.Abs((_dafny.MultiSetOf(_102_v3, _102_v3, _102_v3, _102_v3)).Cardinality())), _101_v2, Companion_Default___.Fm4(_201_v91, _98_globalState), _98_globalState), _202_v92, _202_v92, _202_v92)
			var _204_v94 _dafny.Sequence
			_ = _204_v94
			_204_v94 = _dafny.SeqOf(_102_v3, (_203_v93).Cardinality(), Companion_Default___.Fm4(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), _102_v3), _98_globalState), _102_v3)
			_102_v3 = (_dafny.Zero).Minus((_204_v94).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_102_v3), _dafny.IntOfUint32((_204_v94).Cardinality()))).Uint32()).(_dafny.Int))
			var _205_v95 _dafny.Set
			_ = _205_v95
			_205_v95 = _dafny.SetOf(_101_v2)
			var _206_v96 _dafny.Sequence
			_ = _206_v96
			_206_v96 = _dafny.SeqOf(_205_v95, _205_v95, _205_v95, _205_v95)
			var _207_v97 _dafny.Map
			_ = _207_v97
			_207_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(((_206_v96).Select((Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_206_v96).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()), _180_v72)
			var _208_v98 *C0
			_ = _208_v98
			var _209_v99 _dafny.Array
			_ = _209_v99
			var _210_v100 _dafny.Int
			_ = _210_v100
			var _211_v101 _dafny.Int
			_ = _211_v101
			var _out8 *C0
			_ = _out8
			var _out9 _dafny.Array
			_ = _out9
			var _out10 _dafny.Int
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out8, _out9, _out10, _out11 = Companion_Default___.M1(_202_v92, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(981))).Uint32(), func(coer14 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_212_v92 _dafny.CodePoint) func(_dafny.Int) D0 {
				return func(_213_i9 _dafny.Int) D0 {
					return Companion_D0_.Create_DC0_(_212_v92)
				}
			})(_202_v92))), _101_v2, _dafny.Companion_Sequence_.Update(_199_v89, (Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_199_v89).Cardinality()))).Uint32(), (_207_v97).Cardinality()), _98_globalState)
			_208_v98 = _out8
			_209_v99 = _out9
			_210_v100 = _out10
			_211_v101 = _out11
			_102_v3 = Companion_Default___.Fm4(Companion_Default___.Fm7(_210_v100, _211_v101, _98_globalState), _98_globalState)
			var _214_v102 _dafny.Sequence
			_ = _214_v102
			_214_v102 = _dafny.SeqOf(_176_v68, Companion_Default___.Fm9(_98_globalState))
			var _215_v103 _dafny.Map
			_ = _215_v103
			_215_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _101_v2)
			var _216_v104 _dafny.Sequence
			_ = _216_v104
			_216_v104 = _dafny.SeqOf(_215_v103)
			(_180_v72).M0((_214_v102).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_216_v104).Cardinality()), _dafny.IntOfUint32((_214_v102).Cardinality()))).Uint32()).(_dafny.Sequence), Companion_Default___.Fm4(_201_v91, _98_globalState), _98_globalState)
			var _217_v105 D0
			_ = _217_v105
			_217_v105 = Companion_D0_.Create_DC0_(_202_v92)
			var _218_v106 _dafny.Sequence
			_ = _218_v106
			_218_v106 = _dafny.SeqOf(_217_v105)
			var _219_v108 *C0
			_ = _219_v108
			var _220_v109 _dafny.Array
			_ = _220_v109
			var _221_v110 _dafny.Int
			_ = _221_v110
			var _222_v111 _dafny.Int
			_ = _222_v111
			var _out12 *C0
			_ = _out12
			var _out13 _dafny.Array
			_ = _out13
			var _out14 _dafny.Int
			_ = _out14
			var _out15 _dafny.Int
			_ = _out15
			_out12, _out13, _out14, _out15 = Companion_Default___.M1(_202_v92, _218_v106, _101_v2, _dafny.SeqOf((_dafny.Zero).Minus(_211_v101), Companion_Default___.Fm4(func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter10 := _dafny.Iterate((_dafny.SeqOf(_176_v68)).Elements()); ; {
					_compr_9, _ok10 := _iter10()
					if !_ok10 {
						break
					}
					var _223_v107 _dafny.Sequence
					_223_v107 = interface{}(_compr_9).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_176_v68), _223_v107) {
						_coll9.Add(_223_v107, (_205_v95).Cardinality())
					}
				}
				return _coll9.ToMap()
			}(), _98_globalState)), _98_globalState)
			_219_v108 = _out12
			_220_v109 = _out13
			_221_v110 = _out14
			_222_v111 = _out15
		}
	}
	var _224_v112 _dafny.Array
	_ = _224_v112
	var _nw33 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(17))
	_ = _nw33
	_224_v112 = _nw33
	var _225_v113 _dafny.Sequence
	_ = _225_v113
	_225_v113 = _dafny.SeqOf(_224_v112)
	(_98_globalState).F2 = _dafny.IntOfUint32((_225_v113).Cardinality())
	var _226_v114 *C0
	_ = _226_v114
	var _nw34 *C0 = New_C0_()
	_ = _nw34
	_nw34.Ctor__()
	_226_v114 = _nw34
	if _101_v2 {
		var _227_v115 _dafny.Array
		_ = _227_v115
		var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
		_ = _nw35
		_227_v115 = _nw35
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))
		_ = _index18
		(_227_v115).ArraySet1(_102_v3, (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))
		_ = _index19
		(_227_v115).ArraySet1(Companion_Default___.SafeDivisionInt(_102_v3, _102_v3), (_index19).Int())
		if !(_101_v2) || ((!(_101_v2)) && (_101_v2)) {
			var _228_v116 _dafny.Array
			_ = _228_v116
			var _nwElement0_10 _dafny.Array = _96_v0
			_ = _nwElement0_10
			var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(9))
			_ = _nw36
			(_nw36).ArraySet1(_nwElement0_10, 0)
			(_nw36).ArraySet1(_96_v0, 1)
			(_nw36).ArraySet1(_96_v0, 2)
			(_nw36).ArraySet1(_96_v0, 3)
			(_nw36).ArraySet1(_96_v0, 4)
			(_nw36).ArraySet1(_96_v0, 5)
			(_nw36).ArraySet1(_96_v0, 6)
			(_nw36).ArraySet1(_96_v0, 7)
			(_nw36).ArraySet1(_96_v0, 8)
			_228_v116 = _nw36
			var _229_v117 _dafny.Map
			_ = _229_v117
			_229_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_228_v116, false)
			var _230_v118 _dafny.MultiSet
			_ = _230_v118
			_230_v118 = _dafny.MultiSetOf(_101_v2)
			var _231_v119 _dafny.Sequence
			_ = _231_v119
			_231_v119 = _dafny.SeqOf(true, _101_v2, _101_v2)
			_229_v117 = (_229_v117).Update(_228_v116, !((_226_v114).Fm3(_101_v2, _101_v2, _230_v118, _231_v119, _98_globalState)))
			(_98_globalState).F10 = !((_101_v2) && (_101_v2))
			(_226_v114).M0(_231_v119, _102_v3, _98_globalState)
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))
			_ = _index20
			(_227_v115).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(142), _dafny.IntOfInt64(898)), (_227_v115).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))).Int()).(_dafny.Int))), (_index20).Int())
			(_98_globalState).F2 = (_dafny.Zero).Minus((func() _dafny.Int {
				if _101_v2 {
					return (_dafny.Zero).Minus((_227_v115).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))).Int()).(_dafny.Int))
				}
				return _102_v3
			})())
		} else {
			_97_v1 = _dafny.UnicodeSeqOfUtf8Bytes("jkgtl")
			_226_v114 = _226_v114
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))
			_ = _index21
			(_227_v115).ArraySet1((_227_v115).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))).Int()).(_dafny.Int), (_index21).Int())
			var _232_v120 _dafny.Map
			_ = _232_v120
			_232_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _101_v2)
			var _233_v121 _dafny.Map
			_ = _233_v121
			_233_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _232_v120)
			var _234_v122 _dafny.Map
			_ = _234_v122
			_234_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_232_v120, (func() _dafny.Map {
				if (_233_v121).Contains(true) {
					return (_233_v121).Get(true).(_dafny.Map)
				}
				return _232_v120
			})()), _101_v2)
			var _235_v123 _dafny.MultiSet
			_ = _235_v123
			_235_v123 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _101_v2), _232_v120, _232_v120)
			(_98_globalState).F13 = (func() bool {
				if (_234_v122).Contains(_235_v123) {
					return (_234_v122).Get(_235_v123).(bool)
				}
				return _101_v2
			})()
			var _236_v124 _dafny.CodePoint
			_ = _236_v124
			_236_v124 = _dafny.CodePoint('w')
			var _237_v125 _dafny.Map
			_ = _237_v125
			_237_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_236_v124, _96_v0)
			(_98_globalState).F13 = (_237_v125).Contains(_236_v124)
		}
		var _nw37 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
		_ = _nw37
		_96_v0 = _nw37
		if _101_v2 {
			_102_v3 = _dafny.IntOfInt64(285)
			var _238_v126 _dafny.Sequence
			_ = _238_v126
			_238_v126 = _dafny.SeqOf(_101_v2, _101_v2, _101_v2)
			(_226_v114).M0(_238_v126, (_227_v115).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))).Int()).(_dafny.Int), _98_globalState)
			var _239_v127 _dafny.Set
			_ = _239_v127
			_239_v127 = _dafny.SetOf(_101_v2, _101_v2)
			var _240_v128 _dafny.Map
			_ = _240_v128
			_240_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_227_v115).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))).Int()).(_dafny.Int), (_239_v127).Cardinality())
			var _241_v129 _dafny.MultiSet
			_ = _241_v129
			_241_v129 = _dafny.MultiSetOf(_101_v2)
			var _242_v130 _dafny.Map
			_ = _242_v130
			_242_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, (_226_v114).Fm3(false, _101_v2, _241_v129, _238_v126, _98_globalState))
			var _243_v131 _dafny.Map
			_ = _243_v131
			_243_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, (func() _dafny.Int {
				if (_240_v128).Contains(_dafny.IntOfInt64(577)) {
					return (_240_v128).Get(_dafny.IntOfInt64(577)).(_dafny.Int)
				}
				return (_242_v130).Cardinality()
			})())
			var _244_v132 _dafny.Map
			_ = _244_v132
			_244_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v3, _243_v131)
			var _245_v133 D1
			_ = _245_v133
			_245_v133 = Companion_D1_.Create_DC4_(false, (_227_v115).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))).Int()).(_dafny.Int), _97_v1)
			var _246_v134 _dafny.Sequence
			_ = _246_v134
			_246_v134 = _dafny.SeqOf(_245_v133)
			var _247_v135 D1
			_ = _247_v135
			_247_v135 = Companion_D1_.Create_DC4_(_101_v2, _dafny.IntOfUint32((_246_v134).Cardinality()), _97_v1)
			_244_v132 = (_244_v132).Update(((_247_v135).Dtor_cf8()).Minus(_102_v3), _243_v131)
			var _248_v136 _dafny.Set
			_ = _248_v136
			_248_v136 = _dafny.SetOf(_226_v114)
			var _249_v137 _dafny.Map
			_ = _249_v137
			_249_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v3, _248_v136)
			_249_v137 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_227_v115).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))).Int()).(_dafny.Int), _248_v136)).Merge(_249_v137)
			(_98_globalState).F10 = _101_v2
		} else {
			var _250_v139 _dafny.Sequence
			_ = _250_v139
			_250_v139 = _dafny.SeqOf((func() _dafny.Map {
				var _coll10 = _dafny.NewMapBuilder()
				_ = _coll10
				for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-420), _dafny.IntOfInt64(960))); ; {
					_compr_10, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _251_v138 _dafny.Int
					_251_v138 = interface{}(_compr_10).(_dafny.Int)
					if ((_dafny.IntOfInt64(-420)).Cmp(_251_v138) <= 0) && ((_251_v138).Cmp(_dafny.IntOfInt64(960)) < 0) {
						_coll10.Add((_251_v138).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_227_v115).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))).Int()).(_dafny.Int), _101_v2)).Cardinality()), _dafny.MultiSetOf(_101_v2, _101_v2, _101_v2, _101_v2, !(_101_v2)))
					}
				}
				return _coll10.ToMap()
			}()).Cardinality(), _dafny.IntOfUint32((_97_v1).Cardinality()))
			var _252_v140 D1
			_ = _252_v140
			_252_v140 = Companion_D1_.Create_DC3_(_250_v139, _101_v2)
			(_98_globalState).F10 = (_252_v140).Dtor_cf6()
			var _253_v141 *C0
			_ = _253_v141
			var _nw38 *C0 = New_C0_()
			_ = _nw38
			_nw38.Ctor__()
			_253_v141 = _nw38
			(_98_globalState).F15 = false
			var _254_v142 _dafny.Array
			_ = _254_v142
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_4
			var _nw39 _dafny.Array
			_ = _nw39
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw39 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Sequence = (func(_255_v1 _dafny.Sequence, _256_v115 _dafny.Array) func(_dafny.Int) _dafny.Sequence {
					return func(_257_i10 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Update(_255_v1, (Companion_Default___.SafeIndex((_256_v115).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_256_v115), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_255_v1).Cardinality()))).Uint32(), _dafny.CodePoint('m'))
					}
				})(_97_v1, _227_v115)
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
			_254_v142 = _nw39
			var _258_v143 _dafny.Map
			_ = _258_v143
			_258_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_101_v2), _97_v1)
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_254_v142), 0))
			_ = _index22
			(_254_v142).ArraySet1((func() _dafny.Sequence {
				if (_258_v143).Contains(_101_v2) {
					return (_258_v143).Get(_101_v2).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("pojsssw")
			})(), (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_254_v142), 0))
			_ = _index23
			(_254_v142).ArraySet1(_97_v1, (_index23).Int())
			var _259_v144 D1
			_ = _259_v144
			_259_v144 = Companion_D1_.Create_DC2_(_102_v3)
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_227_v115), 0))
			_ = _index24
			(_227_v115).ArraySet1((_259_v144).Dtor_cf4(), (_index24).Int())
		}
		var _260_v145 _dafny.MultiSet
		_ = _260_v145
		_260_v145 = _dafny.MultiSetOf(_101_v2)
		var _261_v146 _dafny.Sequence
		_ = _261_v146
		_261_v146 = _dafny.SeqOf(_101_v2)
		var _262_v147 _dafny.Map
		_ = _262_v147
		_262_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_226_v114).Fm3(_101_v2, _101_v2, _260_v145, _261_v146, _98_globalState), _102_v3)
		_262_v147 = _262_v147
	} else {
		var _263_v148 _dafny.Sequence
		_ = _263_v148
		_263_v148 = _dafny.SeqOf(!(_101_v2))
		(_226_v114).M0(_263_v148, Companion_Default___.SafeDivisionInt(_102_v3, _102_v3), _98_globalState)
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_224_v112), 0))
		_ = _index25
		(_224_v112).ArraySet1(_226_v114, (_index25).Int())
		var _264_v149 _dafny.Map
		_ = _264_v149
		_264_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v1, _226_v114)
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_224_v112), 0))
		_ = _index26
		var _rhs8 bool = true
		_ = _rhs8
		var _rhs9 *C0 = (func() *C0 {
			if (_264_v149).Contains(_dafny.UnicodeSeqOfUtf8Bytes("ucwu")) {
				return (_264_v149).Get(_dafny.UnicodeSeqOfUtf8Bytes("ucwu")).(*C0)
			}
			return _226_v114
		})()
		_ = _rhs9
		var _lhs6 *GlobalState = _98_globalState
		_ = _lhs6
		var _lhs7 _dafny.Array = _224_v112
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_224_v112), 0))
		_ = _lhs8
		_lhs6.F13 = _rhs8
		(_lhs7).ArraySet1(_rhs9, (_lhs8).Int())
		var _265_v150 _dafny.Map
		_ = _265_v150
		_265_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dioapv")).Cardinality()))
		_265_v150 = (_265_v150).Update(_102_v3, _102_v3)
		var _266_v151 _dafny.CodePoint
		_ = _266_v151
		_266_v151 = _dafny.CodePoint('v')
		var _rhs10 _dafny.CodePoint = _266_v151
		_ = _rhs10
		var _rhs11 _dafny.Array = _96_v0
		_ = _rhs11
		_266_v151 = _rhs10
		_96_v0 = _rhs11
		(_98_globalState).F10 = (func() bool {
			if (_102_v3).Cmp(_102_v3) < 0 {
				return !(_101_v2) || ((_263_v148).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_97_v1).Cardinality())), _dafny.IntOfUint32((_263_v148).Cardinality()))).Uint32()).(bool))
			}
			return _101_v2
		})()
	}
	var _267_i11 _dafny.Int
	_ = _267_i11
	_267_i11 = _dafny.Zero
	{
		for _101_v2 {
			{
				if (_267_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_267_i11 = (_267_i11).Plus(_dafny.One)
				var _268_v152 _dafny.Array
				_ = _268_v152
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_5
				var _nw40 _dafny.Array
				_ = _nw40
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw40 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) _dafny.Sequence = (func(_269_v2 bool) func(_dafny.Int) _dafny.Sequence {
						return func(_270_i12 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_269_v2, _269_v2), _dafny.SeqOf(false))
						}
					})(_101_v2)
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw40 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw40).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw40).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_268_v152 = _nw40
				var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_268_v152), 0))
				_ = _index27
				(_268_v152).ArraySet1(_dafny.SeqOf(_101_v2), (_index27).Int())
				var _271_v153 _dafny.Sequence
				_ = _271_v153
				_271_v153 = _dafny.SeqOf(_101_v2, !(_101_v2))
				var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_268_v152), 0))
				_ = _index28
				(_268_v152).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_271_v153, _271_v153), (_index28).Int())
				var _272_v154 _dafny.Map
				_ = _272_v154
				_272_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_271_v153, _dafny.IntOfUint32((_97_v1).Cardinality()))
				var _273_v155 _dafny.Set
				_ = _273_v155
				_273_v155 = _dafny.SetOf(false)
				var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_268_v152), 0))
				_ = _index29
				var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_268_v152), 0))
				_ = _index30
				var _rhs12 _dafny.Int = _102_v3
				_ = _rhs12
				var _rhs13 _dafny.Int = _102_v3
				_ = _rhs13
				var _rhs14 _dafny.Sequence = _271_v153
				_ = _rhs14
				var _rhs15 _dafny.Int = Companion_Default___.Fm4(_272_v154, _98_globalState)
				_ = _rhs15
				var _rhs16 _dafny.Sequence = (func() _dafny.Sequence {
					if (_dafny.SetOf(true, _101_v2, _101_v2, _101_v2)).IsProperSubsetOf(_273_v155) {
						return _271_v153
					}
					return _271_v153
				})()
				_ = _rhs16
				var _lhs9 *GlobalState = _98_globalState
				_ = _lhs9
				var _lhs10 _dafny.Array = _268_v152
				_ = _lhs10
				var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_268_v152), 0))
				_ = _lhs11
				var _lhs12 *GlobalState = _98_globalState
				_ = _lhs12
				var _lhs13 _dafny.Array = _268_v152
				_ = _lhs13
				var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(881), _dafny.ArrayLen((_268_v152), 0))
				_ = _lhs14
				_102_v3 = _rhs12
				_lhs9.F0 = _rhs13
				(_lhs10).ArraySet1(_rhs14, (_lhs11).Int())
				_lhs12.F9 = _rhs15
				(_lhs13).ArraySet1(_rhs16, (_lhs14).Int())
				(_98_globalState).F9 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if _101_v2 {
						return _102_v3
					}
					return _102_v3
				})(), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_102_v3, _102_v3)))
				(_98_globalState).F2 = _102_v3
				var _274_v156 D1
				_ = _274_v156
				_274_v156 = Companion_D1_.Create_DC2_(_102_v3)
				var _275_v157 D1
				_ = _275_v157
				_275_v157 = Companion_D1_.Create_DC2_((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
					if _101_v2 {
						return (_274_v156).Dtor_cf4()
					}
					return (_273_v155).Cardinality()
				})())))
				var _source3 D1 = _274_v156
				_ = _source3
				if _source3.Is_DC3() {
					var _276___mcc_h0 _dafny.Sequence = _source3.Get_().(D1_DC3).Cf5
					_ = _276___mcc_h0
					var _277___mcc_h1 bool = _source3.Get_().(D1_DC3).Cf6
					_ = _277___mcc_h1
					var _278_cf6 bool = _277___mcc_h1
					_ = _278_cf6
					var _279_cf5 _dafny.Sequence = _276___mcc_h0
					_ = _279_cf5
					(_98_globalState).F10 = !(_278_cf6) || (_dafny.Companion_Sequence_.Equal(_97_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(679))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg15 _dafny.Int) interface{} {
							return coer15(arg15)
						}
					}(func(_280_i13 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('m')
					}))))
					var _281_v158 _dafny.Sequence
					_ = _281_v158
					_281_v158 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_97_v1, _dafny.UnicodeSeqOfUtf8Bytes("ijprsydy")), _dafny.Companion_Sequence_.Concatenate(_97_v1, _dafny.UnicodeSeqOfUtf8Bytes("h")), _dafny.UnicodeSeqOfUtf8Bytes("xceaadkiw"))
					var _282_v159 _dafny.MultiSet
					_ = _282_v159
					_282_v159 = _dafny.MultiSetOf(_101_v2, _101_v2)
					var _283_v160 _dafny.Map
					_ = _283_v160
					_283_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v3, _dafny.CodePoint('b'))
					var _284_v161 _dafny.CodePoint
					_ = _284_v161
					_284_v161 = _dafny.CodePoint('n')
					var _rhs17 _dafny.Sequence = _281_v158
					_ = _rhs17
					var _rhs18 *C0 = _226_v114
					_ = _rhs18
					var _rhs19 _dafny.Int = (_102_v3).Minus((func() _dafny.Int {
						if (_282_v159).Contains(_278_cf6) {
							return (_282_v159).Multiplicity(_278_cf6)
						}
						return _102_v3
					})())
					_ = _rhs19
					var _rhs20 _dafny.Sequence = _dafny.SeqOf((func() _dafny.CodePoint {
						if (_283_v160).Contains(_dafny.IntOfInt64(-218)) {
							return (_283_v160).Get(_dafny.IntOfInt64(-218)).(_dafny.CodePoint)
						}
						return _284_v161
					})(), _284_v161, _284_v161, _284_v161, _284_v161)
					_ = _rhs20
					var _rhs21 bool = (_102_v3).Cmp(_102_v3) <= 0
					_ = _rhs21
					var _lhs15 *GlobalState = _98_globalState
					_ = _lhs15
					var _lhs16 *GlobalState = _98_globalState
					_ = _lhs16
					_281_v158 = _rhs17
					_226_v114 = _rhs18
					_lhs15.F0 = _rhs19
					_97_v1 = _rhs20
					_lhs16.F15 = _rhs21
					var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_96_v0), 0))
					_ = _index31
					(_96_v0).ArraySet1(_278_cf6, (_index31).Int())
					var _285_v162 _dafny.Array
					_ = _285_v162
					var _len0_6 _dafny.Int = _dafny.IntOfInt64(2)
					_ = _len0_6
					var _nw41 _dafny.Array
					_ = _nw41
					if _len0_6.Cmp(_dafny.Zero) == 0 {
						_nw41 = _dafny.NewArray(_len0_6)
					} else {
						var _init6 func(_dafny.Int) _dafny.Set = (func(_286_v155 _dafny.Set) func(_dafny.Int) _dafny.Set {
							return func(_287_i14 _dafny.Int) _dafny.Set {
								return _286_v155
							}
						})(_273_v155)
						_ = _init6
						var _element0_6 = _init6(_dafny.Zero)
						_ = _element0_6
						_nw41 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
						(_nw41).ArraySet1(_element0_6, 0)
						var _nativeLen0_6 = (_len0_6).Int()
						_ = _nativeLen0_6
						for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
							(_nw41).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
						}
					}
					_285_v162 = _nw41
					var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_285_v162), 0))
					_ = _index32
					(_285_v162).ArraySet1((_273_v155).Intersection(_dafny.SetOf(_278_cf6, _101_v2)), (_index32).Int())
					var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_96_v0), 0))
					_ = _index33
					var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_285_v162), 0))
					_ = _index34
					var _rhs22 bool = _101_v2
					_ = _rhs22
					var _rhs23 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rugsxcnc")).Cardinality()), _dafny.IntOfUint32((_279_cf5).Cardinality()))
					_ = _rhs23
					var _rhs24 _dafny.Set = (_273_v155).Difference(Companion_Default___.Fm10(_101_v2, _278_cf6, _102_v3, _284_v161, _98_globalState))
					_ = _rhs24
					var _rhs25 _dafny.Sequence = _279_cf5
					_ = _rhs25
					var _lhs17 _dafny.Array = _96_v0
					_ = _lhs17
					var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_96_v0), 0))
					_ = _lhs18
					var _lhs19 *GlobalState = _98_globalState
					_ = _lhs19
					var _lhs20 _dafny.Array = _285_v162
					_ = _lhs20
					var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_285_v162), 0))
					_ = _lhs21
					(_lhs17).ArraySet1(_rhs22, (_lhs18).Int())
					_lhs19.F2 = _rhs23
					(_lhs20).ArraySet1(_rhs24, (_lhs21).Int())
					_279_cf5 = _rhs25
					var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_96_v0), 0))
					_ = _index35
					var _rhs26 *C0 = _226_v114
					_ = _rhs26
					var _rhs27 bool = _278_cf6
					_ = _rhs27
					var _rhs28 bool = (_96_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_96_v0), 0))).Int()).(bool)
					_ = _rhs28
					var _lhs22 *GlobalState = _98_globalState
					_ = _lhs22
					var _lhs23 _dafny.Array = _96_v0
					_ = _lhs23
					var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_96_v0), 0))
					_ = _lhs24
					_226_v114 = _rhs26
					_lhs22.F15 = _rhs27
					(_lhs23).ArraySet1(_rhs28, (_lhs24).Int())
				} else if _source3.Is_DC4() {
					var _288___mcc_h2 bool = _source3.Get_().(D1_DC4).Cf7
					_ = _288___mcc_h2
					var _289___mcc_h3 _dafny.Int = _source3.Get_().(D1_DC4).Cf8
					_ = _289___mcc_h3
					var _290___mcc_h4 _dafny.Sequence = _source3.Get_().(D1_DC4).Cf9
					_ = _290___mcc_h4
					var _291_cf9 _dafny.Sequence = _290___mcc_h4
					_ = _291_cf9
					var _292_cf8 _dafny.Int = _289___mcc_h3
					_ = _292_cf8
					var _293_cf7 bool = _288___mcc_h2
					_ = _293_cf7
					var _294_v163 _dafny.MultiSet
					_ = _294_v163
					_294_v163 = _dafny.MultiSetOf(Companion_Default___.Fm4(_272_v154, _98_globalState))
					var _295_v164 _dafny.CodePoint
					_ = _295_v164
					_295_v164 = _dafny.CodePoint('n')
					var _296_v165 _dafny.Map
					_ = _296_v165
					_296_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_101_v2), _295_v164)
					var _297_v166 _dafny.MultiSet
					_ = _297_v166
					_297_v166 = _dafny.MultiSetOf(_101_v2)
					var _298_v167 _dafny.Sequence
					_ = _298_v167
					_298_v167 = _dafny.SeqOf((func() _dafny.Int {
						if (_294_v163).Contains(_102_v3) {
							return (_294_v163).Multiplicity(_102_v3)
						}
						return (_297_v166).Cardinality()
					})(), _292_cf8)
					var _299_v168 _dafny.Map
					_ = _299_v168
					_299_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _dafny.MultiSetFromSeq(_298_v167))
					var _300_v169 _dafny.Sequence
					_ = _300_v169
					_300_v169 = _dafny.SeqOf((_296_v165).Cardinality(), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mfd")).Cardinality())), _102_v3), (_292_cf8).Minus((_299_v168).Cardinality()))
					var _rhs29 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_295_v164, _295_v164), _dafny.Companion_Sequence_.Update(_291_cf9, (Companion_Default___.SafeIndex(_292_cf8, _dafny.IntOfUint32((_291_cf9).Cardinality()))).Uint32(), _295_v164)), (func() _dafny.Sequence {
						if !(_101_v2) {
							return _97_v1
						}
						return _97_v1
					})())
					_ = _rhs29
					var _rhs30 _dafny.Int = (_300_v169).Select((Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_300_v169).Cardinality()))).Uint32()).(_dafny.Int)
					_ = _rhs30
					var _rhs31 _dafny.Int = _292_cf8
					_ = _rhs31
					var _rhs32 _dafny.MultiSet = (_dafny.MultiSetOf(_292_cf8, _292_cf8, _102_v3, (_dafny.Zero).Minus(_292_cf8))).Intersection((_294_v163).Intersection(_294_v163))
					_ = _rhs32
					var _lhs25 *GlobalState = _98_globalState
					_ = _lhs25
					var _lhs26 *GlobalState = _98_globalState
					_ = _lhs26
					var _lhs27 *GlobalState = _98_globalState
					_ = _lhs27
					_lhs25.F7 = _rhs29
					_lhs26.F9 = _rhs30
					_lhs27.F2 = _rhs31
					_294_v163 = _rhs32
					var _301_v170 _dafny.Array
					_ = _301_v170
					var _nwElement0_11 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_97_v1, _97_v1)
					_ = _nwElement0_11
					var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(8))
					_ = _nw42
					(_nw42).ArraySet1(_nwElement0_11, 0)
					(_nw42).ArraySet1(_97_v1, 1)
					(_nw42).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ws"), 2)
					(_nw42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_97_v1, _291_cf9), 3)
					(_nw42).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(741))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg16 _dafny.Int) interface{} {
							return coer16(arg16)
						}
					}((func(_302_v164 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_303_i15 _dafny.Int) _dafny.CodePoint {
							return _302_v164
						}
					})(_295_v164))), _291_cf9), (Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(741))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg17 _dafny.Int) interface{} {
							return coer17(arg17)
						}
					}((func(_304_v164 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_305_i15 _dafny.Int) _dafny.CodePoint {
							return _304_v164
						}
					})(_295_v164))), _291_cf9)).Cardinality()))).Uint32(), _295_v164), 4)
					(_nw42).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-212))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg18 _dafny.Int) interface{} {
							return coer18(arg18)
						}
					}(func(_306_i16 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('f')
					})), 5)
					(_nw42).ArraySet1(_291_cf9, 6)
					(_nw42).ArraySet1(_97_v1, 7)
					_301_v170 = _nw42
					var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_301_v170), 0))
					_ = _index36
					(_301_v170).ArraySet1(_291_cf9, (_index36).Int())
					var _307_v171 _dafny.Sequence
					_ = _307_v171
					_307_v171 = _dafny.SeqOf(_297_v166)
					var _pat_let_tv0 = _101_v2
					_ = _pat_let_tv0
					var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_301_v170), 0))
					_ = _index37
					(_301_v170).ArraySet1((func() _dafny.Sequence {
						if (_226_v114).Fm3(_293_cf7, _101_v2, (_307_v171).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.IntOfUint32((_307_v171).Cardinality()))).Uint32()).(_dafny.MultiSet), _271_v153, _98_globalState) {
							return _dafny.Companion_Sequence_.Concatenate(_291_cf9, _291_cf9)
						}
						return Companion_Default___.Fm5((func(_pat_let0_0 D0) D0 {
							return func(_308_dt__update__tmp_h0 D0) D0 {
								return func(_pat_let1_0 bool) D0 {
									return func(_309_dt__update_hcf2_h0 bool) D0 {
										return Companion_D0_.Create_DC1_((_308_dt__update__tmp_h0).Dtor_cf1(), _309_dt__update_hcf2_h0, (_308_dt__update__tmp_h0).Dtor_cf3())
									}(_pat_let1_0)
								}(_pat_let_tv0)
							}(_pat_let0_0)
						}(Companion_D0_.Create_DC1_((_226_v114).Fm3(true, _101_v2, _297_v166, _271_v153, _98_globalState), _101_v2, _101_v2))).Dtor_cf2(), (_226_v114).Fm3(!(_293_cf7), _101_v2, _297_v166, (_268_v152).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_268_v152), 0))).Int()).(_dafny.Sequence), _98_globalState), _98_globalState)
					})(), (_index37).Int())
					var _310_v172 _dafny.Sequence
					_ = _310_v172
					_310_v172 = _dafny.SeqOf(_271_v153)
					_310_v172 = _dafny.Companion_Sequence_.Concatenate(_310_v172, _310_v172)
					_295_v164 = _295_v164
				} else if _source3.Is_DC2() {
					var _311___mcc_h5 _dafny.Int = _source3.Get_().(D1_DC2).Cf4
					_ = _311___mcc_h5
					var _312_cf4 _dafny.Int = _311___mcc_h5
					_ = _312_cf4
					var _313_v173 *C0
					_ = _313_v173
					var _nw43 *C0 = New_C0_()
					_ = _nw43
					_nw43.Ctor__()
					_313_v173 = _nw43
					var _314_v174 _dafny.Array
					_ = _314_v174
					var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
					_ = _nw44
					_314_v174 = _nw44
					var _315_v175 _dafny.Map
					_ = _315_v175
					_315_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(671), _312_cf4)
					var _316_v176 _dafny.Set
					_ = _316_v176
					_316_v176 = _dafny.SetOf(_315_v175, _315_v175, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_312_cf4, Companion_Default___.Fm4(_272_v154, _98_globalState)))
					var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_314_v174), 0))
					_ = _index38
					(_314_v174).ArraySet1(_316_v176, (_index38).Int())
					var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_314_v174), 0))
					_ = _index39
					(_314_v174).ArraySet1(_316_v176, (_index39).Int())
					var _317_v177 D1
					_ = _317_v177
					_317_v177 = Companion_D1_.Create_DC4_((_273_v155).IsDisjointFrom(_273_v155), _102_v3, _97_v1)
					_317_v177 = _317_v177
					(_98_globalState).F9 = Companion_Default___.Fm4(_272_v154, _98_globalState)
				} else {
					var _318___mcc_h6 D1 = _source3.Get_().(D1_DC5).Cf10
					_ = _318___mcc_h6
					var _319_cf10 D1 = _318___mcc_h6
					_ = _319_cf10
					(_226_v114).M0(_dafny.SeqOf(_101_v2), Companion_Default___.SafeModuloInt(_102_v3, _102_v3), _98_globalState)
					var _320_v178 _dafny.Sequence
					_ = _320_v178
					_320_v178 = _dafny.SeqOf(_226_v114, _226_v114)
					var _321_v179 _dafny.Map
					_ = _321_v179
					_321_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_v114, _320_v178)
					_321_v179 = (_321_v179).Update(_226_v114, (func() _dafny.Sequence {
						if _101_v2 {
							return _320_v178
						}
						return _320_v178
					})())
					var _322_v180 _dafny.CodePoint
					_ = _322_v180
					_322_v180 = _dafny.CodePoint('b')
					var _323_v181 D0
					_ = _323_v181
					_323_v181 = Companion_D0_.Create_DC0_(_322_v180)
					var _324_v182 _dafny.Sequence
					_ = _324_v182
					_324_v182 = _dafny.SeqOf(Companion_D0_.Create_DC0_(_322_v180), _323_v181, _323_v181)
					var _325_v183 _dafny.Map
					_ = _325_v183
					_325_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _dafny.SeqOf(_102_v3))
					var _326_v184 _dafny.Sequence
					_ = _326_v184
					_326_v184 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jgeiw")).Cardinality()))
					var _327_v185 *C0
					_ = _327_v185
					var _328_v186 _dafny.Array
					_ = _328_v186
					var _329_v187 _dafny.Int
					_ = _329_v187
					var _330_v188 _dafny.Int
					_ = _330_v188
					var _out16 *C0
					_ = _out16
					var _out17 _dafny.Array
					_ = _out17
					var _out18 _dafny.Int
					_ = _out18
					var _out19 _dafny.Int
					_ = _out19
					_out16, _out17, _out18, _out19 = Companion_Default___.M1(_322_v180, (func() _dafny.Sequence {
						if _101_v2 {
							return _dafny.SeqOf(Companion_D0_.Create_DC0_(_322_v180), func(_pat_let2_0 D0) D0 {
								return func(_331_dt__update__tmp_h1 D0) D0 {
									return func(_pat_let3_0 _dafny.CodePoint) D0 {
										return func(_332_dt__update_hcf0_h0 _dafny.CodePoint) D0 {
											return Companion_D0_.Create_DC0_(_332_dt__update_hcf0_h0)
										}(_pat_let3_0)
									}(_dafny.CodePoint('m'))
								}(_pat_let2_0)
							}(Companion_D0_.Create_DC0_(_322_v180)), Companion_D0_.Create_DC0_(_dafny.CodePoint('o')), Companion_D0_.Create_DC0_(_322_v180))
						}
						return _324_v182
					})(), _101_v2, (func() _dafny.Sequence {
						if (_325_v183).Contains(_101_v2) {
							return (_325_v183).Get(_101_v2).(_dafny.Sequence)
						}
						return _326_v184
					})(), _98_globalState)
					_327_v185 = _out16
					_328_v186 = _out17
					_329_v187 = _out18
					_330_v188 = _out19
					(_98_globalState).F2 = (_dafny.IntOfInt64(389)).Times(Companion_Default___.Fm4((_272_v154).Update(_271_v153, _102_v3), _98_globalState))
				}
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _333_v189 _dafny.CodePoint
	_ = _333_v189
	_333_v189 = _dafny.CodePoint('p')
	var _334_v190 _dafny.Map
	_ = _334_v190
	_334_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_333_v189, _102_v3)
	var _335_v191 _dafny.Map
	_ = _335_v191
	_335_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _101_v2)
	var _336_v192 _dafny.MultiSet
	_ = _336_v192
	_336_v192 = _dafny.MultiSetOf(_101_v2, _101_v2)
	var _337_v193 _dafny.Sequence
	_ = _337_v193
	_337_v193 = _dafny.SeqOf(_101_v2)
	var _338_v194 _dafny.Sequence
	_ = _338_v194
	_338_v194 = _dafny.SeqOf((_226_v114).Fm3((func() bool {
		if (_335_v191).Contains(_101_v2) {
			return (_335_v191).Get(_101_v2).(bool)
		}
		return _101_v2
	})(), _101_v2, _336_v192, _337_v193, _98_globalState), false)
	var _339_v195 _dafny.Sequence
	_ = _339_v195
	_339_v195 = _dafny.SeqOf(_334_v190, Companion_Default___.Fm11(_dafny.MultiSetFromSeq(_337_v193), _102_v3, _101_v2, _98_globalState))
	var _340_i17 _dafny.Int
	_ = _340_i17
	_340_i17 = _dafny.Zero
	{
		for (_226_v114).Fm3(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ipj"), _97_v1), _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_339_v195, (Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_339_v195).Cardinality()))).Uint32(), _334_v190), (Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_339_v195, (Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_339_v195).Cardinality()))).Uint32(), _334_v190)).Cardinality()))).Uint32(), _334_v190), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_334_v190), (Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_dafny.SeqOf(_334_v190)).Cardinality()))).Uint32(), Companion_Default___.Fm11(_dafny.MultiSetOf(_101_v2), _102_v3, _101_v2, _98_globalState))), (_336_v192).Intersection(_dafny.MultiSetOf(_101_v2, _101_v2)), _338_v194, _98_globalState) {
			{
				if (_340_i17).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_340_i17 = (_340_i17).Plus(_dafny.One)
				var _341_v196 _dafny.Array
				_ = _341_v196
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_7
				var _nw45 _dafny.Array
				_ = _nw45
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw45 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.Sequence = func(_342_i18 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("cixtikg")
					}
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw45 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw45).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw45).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_341_v196 = _nw45
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_341_v196), 0))
				_ = _index40
				(_341_v196).ArraySet1(_97_v1, (_index40).Int())
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_341_v196), 0))
				_ = _index41
				(_341_v196).ArraySet1(_97_v1, (_index41).Int())
				var _343_v197 _dafny.Array
				_ = _343_v197
				var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(14))
				_ = _nw46
				_343_v197 = _nw46
				var _344_v198 _dafny.Array
				_ = _344_v198
				var _nwElement0_12 _dafny.Array = _343_v197
				_ = _nwElement0_12
				var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(22))
				_ = _nw47
				(_nw47).ArraySet1(_nwElement0_12, 0)
				(_nw47).ArraySet1(_343_v197, 1)
				(_nw47).ArraySet1(_343_v197, 2)
				(_nw47).ArraySet1(_343_v197, 3)
				(_nw47).ArraySet1(_343_v197, 4)
				(_nw47).ArraySet1(_343_v197, 5)
				(_nw47).ArraySet1(_343_v197, 6)
				(_nw47).ArraySet1(_343_v197, 7)
				(_nw47).ArraySet1(_343_v197, 8)
				(_nw47).ArraySet1(_343_v197, 9)
				(_nw47).ArraySet1(_343_v197, 10)
				(_nw47).ArraySet1(_343_v197, 11)
				(_nw47).ArraySet1(_343_v197, 12)
				(_nw47).ArraySet1(_343_v197, 13)
				(_nw47).ArraySet1(_343_v197, 14)
				(_nw47).ArraySet1(_343_v197, 15)
				(_nw47).ArraySet1(_343_v197, 16)
				(_nw47).ArraySet1(_343_v197, 17)
				(_nw47).ArraySet1((func() _dafny.Array {
					if _101_v2 {
						return _343_v197
					}
					return _343_v197
				})(), 18)
				(_nw47).ArraySet1(_343_v197, 19)
				(_nw47).ArraySet1(_343_v197, 20)
				(_nw47).ArraySet1(_343_v197, 21)
				_344_v198 = _nw47
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_344_v198), 0))
				_ = _index42
				(_344_v198).ArraySet1(_343_v197, (_index42).Int())
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_344_v198), 0))
				_ = _index43
				(_344_v198).ArraySet1(_343_v197, (_index43).Int())
				_96_v0 = _96_v0
				var _345_v199 *C0
				_ = _345_v199
				var _nw48 *C0 = New_C0_()
				_ = _nw48
				_nw48.Ctor__()
				_345_v199 = _nw48
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _346_v200 _dafny.Sequence
	_ = _346_v200
	_346_v200 = _dafny.SeqOf(_102_v3)
	(_226_v114).M0(_338_v194, (_346_v200).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_97_v1).Cardinality()), _dafny.IntOfUint32((_346_v200).Cardinality()))).Uint32()).(_dafny.Int), _98_globalState)
	_338_v194 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_101_v2, _101_v2, _101_v2, _101_v2), (func() _dafny.Sequence {
		if _101_v2 {
			return _dafny.SeqOf(_101_v2, _101_v2, _101_v2)
		}
		return _338_v194
	})())
	var _347_i19 _dafny.Int
	_ = _347_i19
	_347_i19 = _dafny.Zero
	{
		for (!((_226_v114).Fm3(_101_v2, _101_v2, _dafny.MultiSetOf(_101_v2, _101_v2, _101_v2), _337_v193, _98_globalState))) && (_101_v2) {
			{
				if (_347_i19).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_347_i19 = (_347_i19).Plus(_dafny.One)
				(_98_globalState).F10 = _101_v2
				_101_v2 = _101_v2
				(_98_globalState).F15 = !(_101_v2)
				(_98_globalState).F15 = ((_336_v192).Union(_336_v192)).IsProperSubsetOf((_336_v192).Difference(_336_v192))
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	var _348_i20 _dafny.Int
	_ = _348_i20
	_348_i20 = _dafny.Zero
	{
		for (_101_v2) == (_101_v2) {
			{
				if (_348_i20).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L4
				}
				_348_i20 = (_348_i20).Plus(_dafny.One)
				var _349_v201 _dafny.Map
				_ = _349_v201
				_349_v201 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _336_v192)
				(_98_globalState).F15 = ((func() _dafny.MultiSet {
					if (_349_v201).Contains(_101_v2) {
						return (_349_v201).Get(_101_v2).(_dafny.MultiSet)
					}
					return _dafny.MultiSetOf(_101_v2, _101_v2)
				})()).IsSubsetOf(_336_v192)
				(_98_globalState).F15 = (_102_v3).Cmp(_102_v3) <= 0
				var _350_v202 _dafny.Array
				_ = _350_v202
				var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw49
				_350_v202 = _nw49
				var _351_v203 _dafny.Map
				_ = _351_v203
				_351_v203 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _102_v3)
				var _352_v204 _dafny.Set
				_ = _352_v204
				_352_v204 = _dafny.SetOf((_351_v203).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v2, _102_v3)), (_351_v203).Update(_101_v2, _102_v3), _351_v203)
				var _353_v205 _dafny.Map
				_ = _353_v205
				_353_v205 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v1, _97_v1)
				var _rhs33 *C0 = _226_v114
				_ = _rhs33
				var _rhs34 _dafny.Array = _350_v202
				_ = _rhs34
				var _rhs35 bool = !((_353_v205).Update(_97_v1, Companion_Default___.Fm5(_101_v2, _101_v2, _98_globalState))).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-480))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}((func(_354_v189 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_355_i21 _dafny.Int) _dafny.CodePoint {
						return _354_v189
					}
				})(_333_v189))))
				_ = _rhs35
				var _rhs36 _dafny.Set = _352_v204
				_ = _rhs36
				var _lhs28 *GlobalState = _98_globalState
				_ = _lhs28
				_226_v114 = _rhs33
				_350_v202 = _rhs34
				_lhs28.F10 = _rhs35
				_352_v204 = _rhs36
				if (_102_v3).Cmp(_102_v3) <= 0 {
					(_98_globalState).F9 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_97_v1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.IntOfUint32((_97_v1).Cardinality()))).Uint32(), _333_v189)).Cardinality())
					var _356_v206 _dafny.Array
					_ = _356_v206
					var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
					_ = _nw50
					_356_v206 = _nw50
					var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_356_v206), 0))
					_ = _index44
					(_356_v206).ArraySet1(_350_v202, (_index44).Int())
					var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_356_v206), 0))
					_ = _index45
					(_356_v206).ArraySet1(_350_v202, (_index45).Int())
					var _357_v207 _dafny.Map
					_ = _357_v207
					_357_v207 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(72), _102_v3)
					var _358_v208 _dafny.Sequence
					_ = _358_v208
					_358_v208 = _dafny.SeqOf(_357_v207)
					var _359_v209 _dafny.MultiSet
					_ = _359_v209
					_359_v209 = _dafny.MultiSetOf(_102_v3, (((_358_v208).Select((Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_358_v208).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()).Plus(_102_v3), (_102_v3).Plus(_102_v3), Companion_Default___.SafeModuloInt(_102_v3, _102_v3), _102_v3)
					var _360_v211 _dafny.Map
					_ = _360_v211
					_360_v211 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_359_v209).Cardinality(), (_226_v114).Fm3(_101_v2, _101_v2, _dafny.MultiSetOf(_101_v2, _101_v2), _337_v193, _98_globalState))
					var _361_v212 _dafny.Map
					_ = _361_v212
					_361_v212 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_360_v211, (Companion_D1_.Create_DC4_(_101_v2, _102_v3, _97_v1)).Dtor_cf7())
					var _362_v215 _dafny.Set
					_ = _362_v215
					_362_v215 = _dafny.SetOf(_360_v211)
					var _363_v216 _dafny.Sequence
					_ = _363_v216
					_363_v216 = _dafny.SeqOf(_362_v215)
					var _364_v217 _dafny.Sequence
					_ = _364_v217
					_364_v217 = _dafny.SeqOf(func() _dafny.Map {
						var _coll11 = _dafny.NewMapBuilder()
						_ = _coll11
						for _iter12 := _dafny.Iterate(((_363_v216).Select((Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_363_v216).Cardinality()))).Uint32()).(_dafny.Set)).Elements()); ; {
							_compr_11, _ok12 := _iter12()
							if !_ok12 {
								break
							}
							var _365_v214 _dafny.Map
							_365_v214 = interface{}(_compr_11).(_dafny.Map)
							if ((_363_v216).Select((Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_363_v216).Cardinality()))).Uint32()).(_dafny.Set)).Contains(_365_v214) {
								_coll11.Add(_365_v214, _101_v2)
							}
						}
						return _coll11.ToMap()
					}())
					var _366_v218 _dafny.Set
					_ = _366_v218
					_366_v218 = _dafny.SetOf(_101_v2)
					var _367_v219 D4
					_ = _367_v219
					_367_v219 = Companion_D4_.Create_DC8_(_361_v212)
					var _368_v220 _dafny.Array
					_ = _368_v220
					var _nwElement0_13 _dafny.Map = (func() _dafny.Map {
						var _coll12 = _dafny.NewMapBuilder()
						_ = _coll12
						for _iter13 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-660))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg20 _dafny.Int) interface{} {
								return coer20(arg20)
							}
						}((func(_369_v3 _dafny.Int, _370_v2 bool) func(_dafny.Int) _dafny.Map {
							return func(_371_i22 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_369_v3, _370_v2)
							}
						})(_102_v3, _101_v2)))).Elements()); ; {
							_compr_12, _ok13 := _iter13()
							if !_ok13 {
								break
							}
							var _372_v210 _dafny.Map
							_372_v210 = interface{}(_compr_12).(_dafny.Map)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-660))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
								return func(arg21 _dafny.Int) interface{} {
									return coer21(arg21)
								}
							}((func(_373_v3 _dafny.Int, _374_v2 bool) func(_dafny.Int) _dafny.Map {
								return func(_371_i22 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_373_v3, _374_v2)
								}
							})(_102_v3, _101_v2))), _372_v210) {
								_coll12.Add(_372_v210, _101_v2)
							}
						}
						return _coll12.ToMap()
					}()).Merge(_361_v212)
					_ = _nwElement0_13
					var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(20))
					_ = _nw51
					(_nw51).ArraySet1(_nwElement0_13, 0)
					(_nw51).ArraySet1(_361_v212, 1)
					(_nw51).ArraySet1(_361_v212, 2)
					(_nw51).ArraySet1(_361_v212, 3)
					(_nw51).ArraySet1(_361_v212, 4)
					(_nw51).ArraySet1(_361_v212, 5)
					(_nw51).ArraySet1((_361_v212).Update(func() _dafny.Map {
						var _coll13 = _dafny.NewMapBuilder()
						_ = _coll13
						for _iter14 := _dafny.Iterate((_360_v211).Keys().Elements()); ; {
							_compr_13, _ok14 := _iter14()
							if !_ok14 {
								break
							}
							var _375_v213 _dafny.Int
							_375_v213 = interface{}(_compr_13).(_dafny.Int)
							if (_360_v211).Contains(_375_v213) {
								_coll13.Add((_375_v213).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality())), true)
							}
						}
						return _coll13.ToMap()
					}(), _101_v2), 6)
					(_nw51).ArraySet1(_361_v212, 7)
					(_nw51).ArraySet1(_361_v212, 8)
					(_nw51).ArraySet1((_361_v212).Update(_360_v211, _101_v2), 9)
					(_nw51).ArraySet1((_361_v212).Update(_360_v211, !(_101_v2)), 10)
					(_nw51).ArraySet1(((_364_v217).Select((Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_364_v217).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v3, _101_v2)).Update(_102_v3, _101_v2), (_226_v114).Fm3(_101_v2, _101_v2, _dafny.MultiSetOf(_101_v2), _338_v194, _98_globalState))), 11)
					(_nw51).ArraySet1(_361_v212, 12)
					(_nw51).ArraySet1(_361_v212, 13)
					(_nw51).ArraySet1((_361_v212).Merge(_361_v212), 14)
					(_nw51).ArraySet1(_361_v212, 15)
					(_nw51).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_366_v218).Cardinality()), false), (func() bool {
						if (_360_v211).Contains((_335_v191).Cardinality()) {
							return (_360_v211).Get((_335_v191).Cardinality()).(bool)
						}
						return true
					})()), 16)
					(_nw51).ArraySet1(_361_v212, 17)
					(_nw51).ArraySet1((_367_v219).Dtor_cf13(), 18)
					(_nw51).ArraySet1(_361_v212, 19)
					_368_v220 = _nw51
					var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_368_v220), 0))
					_ = _index46
					(_368_v220).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_360_v211, _101_v2)).Merge((_361_v212).Update(func() _dafny.Map {
						var _coll14 = _dafny.NewMapBuilder()
						_ = _coll14
						for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-640), _dafny.IntOfInt64(23))); ; {
							_compr_14, _ok15 := _iter15()
							if !_ok15 {
								break
							}
							var _376_v221 _dafny.Int
							_376_v221 = interface{}(_compr_14).(_dafny.Int)
							if ((_dafny.IntOfInt64(-640)).Cmp(_376_v221) <= 0) && ((_376_v221).Cmp(_dafny.IntOfInt64(23)) < 0) {
								_coll14.Add((_376_v221).Minus((_351_v203).Cardinality()), _101_v2)
							}
						}
						return _coll14.ToMap()
					}(), _101_v2)), (_index46).Int())
					var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_368_v220), 0))
					_ = _index47
					var _rhs37 _dafny.MultiSet = _dafny.MultiSetOf((_102_v3).Minus((_dafny.Zero).Minus((func() _dafny.Int {
						if (_351_v203).Contains(_101_v2) {
							return (_351_v203).Get(_101_v2).(_dafny.Int)
						}
						return _102_v3
					})())))
					_ = _rhs37
					var _rhs38 *C0 = _226_v114
					_ = _rhs38
					var _rhs39 _dafny.Map = (func() _dafny.Map {
						if !(_101_v2) {
							return _361_v212
						}
						return ((_364_v217).Select((Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_364_v217).Cardinality()))).Uint32()).(_dafny.Map)).Merge(func() _dafny.Map {
							var _coll15 = _dafny.NewMapBuilder()
							_ = _coll15
							for _iter16 := _dafny.Iterate((_dafny.SetOf(_360_v211)).Elements()); ; {
								_compr_15, _ok16 := _iter16()
								if !_ok16 {
									break
								}
								var _377_v222 _dafny.Map
								_377_v222 = interface{}(_compr_15).(_dafny.Map)
								if (_dafny.SetOf(_360_v211)).Contains(_377_v222) {
									_coll15.Add(_377_v222, false)
								}
							}
							return _coll15.ToMap()
						}())
					})()
					_ = _rhs39
					var _rhs40 *C0 = _226_v114
					_ = _rhs40
					var _lhs29 _dafny.Array = _368_v220
					_ = _lhs29
					var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_368_v220), 0))
					_ = _lhs30
					_359_v209 = _rhs37
					_226_v114 = _rhs38
					(_lhs29).ArraySet1(_rhs39, (_lhs30).Int())
					_226_v114 = _rhs40
					var _378_v223 _dafny.Array
					_ = _378_v223
					var _nw52 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(3))
					_ = _nw52
					_378_v223 = _nw52
					var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_96_v0), 0))
					_ = _index48
					(_96_v0).ArraySet1(_101_v2, (_index48).Int())
					var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_96_v0), 0))
					_ = _index49
					var _rhs41 _dafny.Int = (_dafny.Zero).Minus((_102_v3).Times(_102_v3))
					_ = _rhs41
					var _rhs42 _dafny.CodePoint = _dafny.CodePoint('v')
					_ = _rhs42
					var _rhs43 _dafny.Array = _378_v223
					_ = _rhs43
					var _rhs44 bool = (_101_v2) && ((_337_v193).Select((Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_337_v193).Cardinality()))).Uint32()).(bool))
					_ = _rhs44
					var _lhs31 *GlobalState = _98_globalState
					_ = _lhs31
					var _lhs32 _dafny.Array = _96_v0
					_ = _lhs32
					var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_96_v0), 0))
					_ = _lhs33
					_lhs31.F2 = _rhs41
					_333_v189 = _rhs42
					_378_v223 = _rhs43
					(_lhs32).ArraySet1(_rhs44, (_lhs33).Int())
					var _379_v224 *C0
					_ = _379_v224
					var _380_v225 _dafny.Array
					_ = _380_v225
					var _381_v226 _dafny.Int
					_ = _381_v226
					var _382_v227 _dafny.Int
					_ = _382_v227
					var _out20 *C0
					_ = _out20
					var _out21 _dafny.Array
					_ = _out21
					var _out22 _dafny.Int
					_ = _out22
					var _out23 _dafny.Int
					_ = _out23
					_out20, _out21, _out22, _out23 = Companion_Default___.M1(_333_v189, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-33))).Uint32(), func(coer22 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
						return func(arg22 _dafny.Int) interface{} {
							return coer22(arg22)
						}
					}((func(_383_v189 _dafny.CodePoint) func(_dafny.Int) D0 {
						return func(_384_i23 _dafny.Int) D0 {
							return Companion_D0_.Create_DC0_(_383_v189)
						}
					})(_333_v189))), _101_v2, _346_v200, _98_globalState)
					_379_v224 = _out20
					_380_v225 = _out21
					_381_v226 = _out22
					_382_v227 = _out23
				} else {
					(_226_v114).M0(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_101_v2), _338_v194), _dafny.IntOfUint32((_346_v200).Cardinality()), _98_globalState)
					_226_v114 = _226_v114
					var _385_v228 *C0
					_ = _385_v228
					var _nw53 *C0 = New_C0_()
					_ = _nw53
					_nw53.Ctor__()
					_385_v228 = _nw53
					var _386_v229 _dafny.Map
					_ = _386_v229
					_386_v229 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v3, _dafny.UnicodeSeqOfUtf8Bytes("emrwvbq"))
					var _387_v230 _dafny.MultiSet
					_ = _387_v230
					_387_v230 = _dafny.MultiSetOf(_dafny.IntOfUint32((_346_v200).Cardinality()))
					var _388_v231 _dafny.MultiSet
					_ = _388_v231
					_388_v231 = _dafny.MultiSetOf(_387_v230, _387_v230)
					(_98_globalState).F7 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
						if (_386_v229).Contains((_388_v231).Cardinality()) {
							return (_386_v229).Get((_388_v231).Cardinality()).(_dafny.Sequence)
						}
						return _dafny.UnicodeSeqOfUtf8Bytes("skwwff")
					})(), Companion_Default___.Fm5(_101_v2, _101_v2, _98_globalState)), _97_v1)
					(_98_globalState).F2 = Companion_Default___.SafeModuloInt(_102_v3, (_102_v3).Times(_102_v3))
				}
				goto C4
			}
		C4:
		}
		goto L4
	}
L4:
	_102_v3 = _102_v3
	var _389_v232 _dafny.Array
	_ = _389_v232
	var _len0_8 _dafny.Int = _dafny.IntOfInt64(20)
	_ = _len0_8
	var _nw54 _dafny.Array
	_ = _nw54
	if _len0_8.Cmp(_dafny.Zero) == 0 {
		_nw54 = _dafny.NewArray(_len0_8)
	} else {
		var _init8 func(_dafny.Int) _dafny.Int = (func(_390_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_391_i24 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_391_i24, (_dafny.Zero).Minus(_390_v3))
			}
		})(_102_v3)
		_ = _init8
		var _element0_8 = _init8(_dafny.Zero)
		_ = _element0_8
		_nw54 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
		(_nw54).ArraySet1(_element0_8, 0)
		var _nativeLen0_8 = (_len0_8).Int()
		_ = _nativeLen0_8
		for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
			(_nw54).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
		}
	}
	_389_v232 = _nw54
	_389_v232 = _389_v232
	var _rhs45 bool = _101_v2
	_ = _rhs45
	var _rhs46 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_97_v1, _dafny.Companion_Sequence_.Update(_97_v1, (Companion_Default___.SafeIndex(_102_v3, _dafny.IntOfUint32((_97_v1).Cardinality()))).Uint32(), _333_v189))
	_ = _rhs46
	var _lhs34 *GlobalState = _98_globalState
	_ = _lhs34
	_101_v2 = _rhs45
	_lhs34.F7 = _rhs46
	_101_v2 = (_336_v192).IsSubsetOf(_dafny.MultiSetOf(_101_v2))
	_389_v232 = _389_v232
	_dafny.Print((_96_v0).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_96_v0).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_97_v1.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_98_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_98_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_98_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_98_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_98_globalState).F6(), _dafny.SeqOf(_dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'), _dafny.CodePoint('u'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_98_globalState.F7.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_98_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_98_globalState.F10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_98_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_98_globalState).F14().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_98_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_101_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_102_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_225_v113).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_267_i11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_333_v189)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_334_v190).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.IntOfInt64(890))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_335_v191).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_336_v192).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_337_v193, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_338_v194, _dafny.SeqOf(false, false, false, false, true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_339_v195, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.IntOfInt64(890)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), _dafny.IntOfInt64(2)).UpdateUnsafe(_dafny.CodePoint('v'), _dafny.IntOfInt64(46)).UpdateUnsafe(_dafny.CodePoint('k'), _dafny.One).UpdateUnsafe(_dafny.CodePoint('n'), _dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_340_i17)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_346_v200, _dafny.SeqOf(_dafny.IntOfInt64(890))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_347_i19)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_348_i20)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v232).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
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
	Cf2 bool
	Cf3 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 bool, Cf3 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.CodePoint
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.CodePoint) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, false, false)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf0() _dafny.CodePoint {
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2 == data2.Cf2 && data1.Cf3 == data2.Cf3
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
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
	Cf5 _dafny.Sequence
	Cf6 bool
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf5 _dafny.Sequence, Cf6 bool) D1 {
	return D1{D1_DC3{Cf5, Cf6}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf7 bool
	Cf8 _dafny.Int
	Cf9 _dafny.Sequence
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf7 bool, Cf8 _dafny.Int, Cf9 _dafny.Sequence) D1 {
	return D1{D1_DC4{Cf7, Cf8, Cf9}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
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

type D1_DC5 struct {
	Cf10 D1
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf10 D1) D1 {
	return D1{D1_DC5{Cf10}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.EmptySeq, false)
}

func (_this D1) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf6() bool {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf10() D1 {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + data.Cf9.VerbatimString(true) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf5.Equals(data2.Cf5) && data1.Cf6 == data2.Cf6
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf7 == data2.Cf7 && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Equals(data2.Cf9)
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
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
	Cf11 _dafny.Map
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf11 _dafny.Map) D2 {
	return D2{D2_DC6{Cf11}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D2) Dtor_cf11() _dafny.Map {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	Cf12 *C0
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf12 *C0) D3 {
	return D3{D3_DC7{Cf12}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() *C0 {
	return (*C0)(nil)
}

func (_this D3) Dtor_cf12() *C0 {
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
			return ok && data1.Cf12 == data2.Cf12
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
	Cf14 _dafny.Int
	Cf15 bool
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf14 _dafny.Int, Cf15 bool) D4 {
	return D4{D4_DC9{Cf14, Cf15}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC10 struct {
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_() D4 {
	return D4{D4_DC10{}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC8 struct {
	Cf13 _dafny.Map
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf13 _dafny.Map) D4 {
	return D4{D4_DC8{Cf13}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC9_(_dafny.Zero, false)
}

func (_this D4) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf14
}

func (_this D4) Dtor_cf15() bool {
	return _this.Get_().(D4_DC9).Cf15
}

func (_this D4) Dtor_cf13() _dafny.Map {
	return _this.Get_().(D4_DC8).Cf13
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10"
		}
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15 == data2.Cf15
		}
	case D4_DC10:
		{
			_, ok := other.Get_().(D4_DC10)
			return ok
		}
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
			return ok && data1.Cf13.Equals(data2.Cf13)
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

type D5_DC12 struct {
	Cf17 _dafny.Int
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf17 _dafny.Int) D5 {
	return D5{D5_DC12{Cf17}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC13 struct {
	Cf18 _dafny.CodePoint
	Cf19 bool
	Cf20 _dafny.Int
	Cf21 _dafny.CodePoint
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf18 _dafny.CodePoint, Cf19 bool, Cf20 _dafny.Int, Cf21 _dafny.CodePoint) D5 {
	return D5{D5_DC13{Cf18, Cf19, Cf20, Cf21}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC11 struct {
	Cf16 _dafny.MultiSet
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf16 _dafny.MultiSet) D5 {
	return D5{D5_DC11{Cf16}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

type D5_DC14 struct {
	Cf22 D5
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf22 D5) D5 {
	return D5{D5_DC14{Cf22}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC12_(_dafny.Zero)
}

func (_this D5) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D5_DC12).Cf17
}

func (_this D5) Dtor_cf18() _dafny.CodePoint {
	return _this.Get_().(D5_DC13).Cf18
}

func (_this D5) Dtor_cf19() bool {
	return _this.Get_().(D5_DC13).Cf19
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf20
}

func (_this D5) Dtor_cf21() _dafny.CodePoint {
	return _this.Get_().(D5_DC13).Cf21
}

func (_this D5) Dtor_cf16() _dafny.MultiSet {
	return _this.Get_().(D5_DC11).Cf16
}

func (_this D5) Dtor_cf22() D5 {
	return _this.Get_().(D5_DC14).Cf22
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf16) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf22) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19 && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21 == data2.Cf21
		}
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
			return ok && data1.Cf16.Equals(data2.Cf16)
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf22.Equals(data2.Cf22)
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

// Definition of class GlobalState
type GlobalState struct {
	F0   _dafny.Int
	F2   _dafny.Int
	F7   _dafny.Sequence
	F9   _dafny.Int
	F10  bool
	F13  bool
	F15  bool
	_f1  _dafny.CodePoint
	_f3  _dafny.Int
	_f4  _dafny.Array
	_f5  bool
	_f6  _dafny.Sequence
	_f8  _dafny.CodePoint
	_f11 _dafny.CodePoint
	_f12 _dafny.Int
	_f14 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F2 = _dafny.Zero
	_this.F7 = _dafny.EmptySeq
	_this.F9 = _dafny.Zero
	_this.F10 = false
	_this.F13 = false
	_this.F15 = false
	_this._f1 = _dafny.CodePoint('D')
	_this._f3 = _dafny.Zero
	_this._f4 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f5 = false
	_this._f6 = _dafny.EmptySeq
	_this._f8 = _dafny.CodePoint('D')
	_this._f11 = _dafny.CodePoint('D')
	_this._f12 = _dafny.Zero
	_this._f14 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.CodePoint, f2 _dafny.Int, f3 _dafny.Int, f4 _dafny.Array, f5 bool, f6 _dafny.Sequence, f7 _dafny.Sequence, f8 _dafny.CodePoint, f9 _dafny.Int, f10 bool, f11 _dafny.CodePoint, f12 _dafny.Int, f13 bool, f14 _dafny.Sequence, f15 bool) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this).F9 = f9
		(_this).F10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
	}
}
func (_this *GlobalState) F1() _dafny.CodePoint {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Array {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Sequence {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F8() _dafny.CodePoint {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F11() _dafny.CodePoint {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() _dafny.Sequence {
	{
		return _this._f14
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
func (_this *C0) Fm3(p0 bool, p1 bool, p2 _dafny.MultiSet, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C0) M0(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) {
	{
		var _392_v0 D1
		_ = _392_v0
		_392_v0 = Companion_D1_.Create_DC2_(p1)
		var _393_v1 bool
		_ = _393_v1
		_393_v1 = true
		var _394_v2 _dafny.Map
		_ = _394_v2
		_394_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_393_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("snasecdf")).Cardinality()))
		var _395_v3 _dafny.Map
		_ = _395_v3
		_395_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(385), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1))
		var _396_v4 _dafny.Sequence
		_ = _396_v4
		_396_v4 = _dafny.SeqOf(_394_v2, (func() _dafny.Map {
			if (_395_v3).Contains(_dafny.IntOfUint32((p0).Cardinality())) {
				return (_395_v3).Get(_dafny.IntOfUint32((p0).Cardinality())).(_dafny.Map)
			}
			return _394_v2
		})())
		var _397_v5 _dafny.Array
		_ = _397_v5
		var _nwElement0_14 _dafny.Int = p1
		_ = _nwElement0_14
		var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(14))
		_ = _nw55
		(_nw55).ArraySet1(_nwElement0_14, 0)
		(_nw55).ArraySet1(p1, 1)
		(_nw55).ArraySet1(p1, 2)
		(_nw55).ArraySet1((_dafny.Zero).Minus((_392_v0).Dtor_cf4()), 3)
		(_nw55).ArraySet1(_dafny.IntOfUint32((_396_v4).Cardinality()), 4)
		(_nw55).ArraySet1(p1, 5)
		(_nw55).ArraySet1((_392_v0).Dtor_cf4(), 6)
		(_nw55).ArraySet1(p1, 7)
		(_nw55).ArraySet1(p1, 8)
		(_nw55).ArraySet1(p1, 9)
		(_nw55).ArraySet1(p1, 10)
		(_nw55).ArraySet1((_dafny.IntOfUint32((p0).Cardinality())).Times(p1), 11)
		(_nw55).ArraySet1(p1, 12)
		(_nw55).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p1), p1), 13)
		_397_v5 = _nw55
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))
		_ = _index50
		(_397_v5).ArraySet1((p1).Times(p1), (_index50).Int())
		var _398_v6 _dafny.Map
		_ = _398_v6
		_398_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))
		_ = _index51
		(_397_v5).ArraySet1(Companion_Default___.Fm4(_398_v6, globalState), (_index51).Int())
		var _399_v7 _dafny.Map
		_ = _399_v7
		_399_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_397_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))).Int()).(_dafny.Int)), _393_v1)
		var _400_v8 _dafny.Set
		_ = _400_v8
		_400_v8 = _dafny.SetOf(_399_v7)
		var _401_v9 _dafny.Map
		_ = _401_v9
		_401_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_400_v8).Cardinality(), p1)
		var _402_v10 _dafny.Set
		_ = _402_v10
		_402_v10 = _dafny.SetOf(_393_v1)
		var _403_v11 _dafny.Sequence
		_ = _403_v11
		_403_v11 = _dafny.SeqOf(_402_v10, _402_v10)
		_401_v9 = (_401_v9).Update(Companion_Default___.SafeModuloInt((_397_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_403_v11).Cardinality())))
		var _404_v12 _dafny.Map
		_ = _404_v12
		_404_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_393_v1, _397_v5)
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))
		_ = _index52
		(_397_v5).ArraySet1((((_404_v12).Merge(_404_v12)).Cardinality()).Minus(_dafny.IntOfInt64(800)), (_index52).Int())
		var _405_v13 _dafny.Array
		_ = _405_v13
		var _nw56 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw56
		_405_v13 = _nw56
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_405_v13), 0))
		_ = _index53
		(_405_v13).ArraySet1(true, (_index53).Int())
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_405_v13), 0))
		_ = _index54
		(_405_v13).ArraySet1(_393_v1, (_index54).Int())
		var _hi1 _dafny.Int = (_397_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))).Int()).(_dafny.Int)
		_ = _hi1
		for _406_i0 := p1; _406_i0.Cmp(_hi1) < 0; _406_i0 = _406_i0.Plus(_dafny.One) {
			(globalState).F0 = Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if _393_v1 {
					return p1
				}
				return (_397_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))).Int()).(_dafny.Int)
			})(), (_dafny.Zero).Minus((_397_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))).Int()).(_dafny.Int)))
			var _407_v14 _dafny.Sequence
			_ = _407_v14
			_407_v14 = _dafny.UnicodeSeqOfUtf8Bytes("cshphu")
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_405_v13), 0))
			_ = _index55
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))
			_ = _index56
			var _rhs47 _dafny.Int = (func() _dafny.Int {
				if (_401_v9).Contains(((_402_v10).Cardinality()).Times((_397_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))).Int()).(_dafny.Int))) {
					return (_401_v9).Get(((_402_v10).Cardinality()).Times((_397_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))).Int()).(_dafny.Int))).(_dafny.Int)
				}
				return Companion_Default___.SafeModuloInt(_406_i0, (_397_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))).Int()).(_dafny.Int))
			})()
			_ = _rhs47
			var _rhs48 bool = _dafny.Companion_Sequence_.Equal(_407_v14, _407_v14)
			_ = _rhs48
			var _rhs49 _dafny.Int = _dafny.IntOfInt64(686)
			_ = _rhs49
			var _lhs35 *GlobalState = globalState
			_ = _lhs35
			var _lhs36 _dafny.Array = _405_v13
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_405_v13), 0))
			_ = _lhs37
			var _lhs38 _dafny.Array = _397_v5
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))
			_ = _lhs39
			_lhs35.F9 = _rhs47
			(_lhs36).ArraySet1(_rhs48, (_lhs37).Int())
			(_lhs38).ArraySet1(_rhs49, (_lhs39).Int())
			(globalState).F13 = false
			(globalState).F0 = _dafny.IntOfInt64(-981)
		}
		var _ingredients0 = _dafny.NewBuilder()
		_ = _ingredients0
		for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_397_v5), 0))); ; {
			_guard_loop_1, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _408_i1 _dafny.Int
			_408_i1 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_408_i1).Sign() != -1) && ((_408_i1).Cmp(_dafny.ArrayLen((_397_v5), 0)) < 0)) {
				_ingredients0.Add(_dafny.TupleOf(_397_v5, (_408_i1).Int(), (_408_i1).Times(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(381))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg23 _dafny.Int) interface{} {
						return coer23(arg23)
					}
				}(func(_409_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				}))).Cardinality()), (_397_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_397_v5), 0))).Int()).(_dafny.Int)))))
			}
		}
		for _iter18 := _dafny.Iterate(_ingredients0); ; {
			_tup0, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
		}
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
