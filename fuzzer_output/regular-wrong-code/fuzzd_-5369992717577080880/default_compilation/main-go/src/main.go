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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, globalState *GlobalState) D0 {
	if (Companion_D4_.Create_DC15_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kuw")).Cardinality()), true)).Dtor_cf24() {
		return Companion_D0_.Create_DC0_(_dafny.IntOfInt64(-581))
	} else {
		return Companion_D0_.Create_DC0_((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(412), true)).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_(true, ((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(16))).Cardinality()))).Plus((_dafny.Zero).Minus((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(904), _dafny.IntOfInt64(148))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(904)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(148)) < 0) {
				_coll0.Add((_0_v0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(635), true))).Cardinality()), true)
			}
		}
		return _coll0.ToMap()
	}()).Cardinality())), (_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dh")).Cardinality())))).Union(_dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-400))))))
}
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) bool {
	return !(false) || (false)
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, (Companion_D0_.Create_DC2_(true, _dafny.IntOfInt64(163), _dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(713))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-701), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality())
	}))).Cardinality()))))).Dtor_cf2())).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gxdaqqlt")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(!(true), true)), _dafny.SeqOf((Companion_D5_.Create_DC18_(_dafny.IntOfInt64(-672), !(!(true)))).Dtor_cf28(), false))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	if (_dafny.MultiSetOf(true)).Equals(_dafny.MultiSetOf(false, true)) {
		return _dafny.UnicodeSeqOfUtf8Bytes("yo")
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jjeyx"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(600))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(false)).Intersection((_dafny.MultiSetOf(true)).Difference(_dafny.MultiSetOf(!((Companion_D0_.Create_DC1_(true)).Dtor_cf1()))))
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(987))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_3_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(398)
	})), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xrmaqmxx")).Cardinality())))), (func() _dafny.Sequence {
		if !(false) {
			return _dafny.SeqOf(_dafny.IntOfInt64(42), _dafny.IntOfInt64(627))
		}
		return _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(false)), func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(421))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_4_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			}))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _5_v0 _dafny.CodePoint
				_5_v0 = interface{}(_compr_1).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(421))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg4 _dafny.Int) interface{} {
						return coer4(arg4)
					}
				}(func(_4_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				})), _5_v0) {
					_coll1.Add(_5_v0, true)
				}
			}
			return _coll1.ToMap()
		}())).Cardinality(), _dafny.IntOfInt64(-641), _dafny.IntOfInt64(944))
	})())
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), !(true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (Companion_D0_.Create_DC1_((Companion_D2_.Create_DC8_(true, _dafny.IntOfInt64(488), false)).Dtor_cf7())).Dtor_cf1())), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))))
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(982))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_6_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("v")
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(641))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_7_i1 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("hchl")
	}))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-593))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_8_i2 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(278))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_9_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		}))
	})), _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(990))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_10_i4 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	})))))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mlvvv")).Cardinality()), false)).Merge(func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(480), _dafny.IntOfInt64(-305))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _11_v0 _dafny.Int
			_11_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(480)).Cmp(_11_v0) <= 0) && ((_11_v0).Cmp(_dafny.IntOfInt64(-305)) < 0) {
				_coll2.Add((_11_v0).Times(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), true)
			}
		}
		return _coll2.ToMap()
	}())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, !(true), true, false)).Cardinality()), !(true)))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(Companion_Default___.SafeModuloInt((_dafny.SetOf(false)).Cardinality(), _dafny.IntOfInt64(837)), (_dafny.Zero).Minus((_dafny.IntOfInt64(437)).Plus((func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-549), _dafny.IntOfInt64(598))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _12_v0 _dafny.Int
			_12_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(-549)).Cmp(_12_v0) <= 0) && ((_12_v0).Cmp(_dafny.IntOfInt64(598)) < 0) {
				_coll3.Add((_12_v0).Plus((_dafny.Zero).Minus(_dafny.IntOfInt64(-485))), false)
			}
		}
		return _coll3.ToMap()
	}()).Cardinality())), _dafny.IntOfInt64(368), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Cardinality(), (_dafny.Zero).Minus((_dafny.SetOf(!(false), true)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm18(globalState *GlobalState) _dafny.CodePoint {
	if ((func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((func() _dafny.Set {
			var _coll5 = _dafny.NewBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality()), _dafny.IntOfInt64(356))).Keys().Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _13_v1 _dafny.Int
				_13_v1 = interface{}(_compr_5).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality()), _dafny.IntOfInt64(356))).Contains(_13_v1) {
					_coll5.Add(Companion_Default___.SafeDivisionInt(_13_v1, _dafny.IntOfInt64(-91)))
				}
			}
			return _coll5.ToSet()
		}()).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _14_v0 _dafny.Int
			_14_v0 = interface{}(_compr_4).(_dafny.Int)
			if (func() _dafny.Set {
				var _coll6 = _dafny.NewBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality()), _dafny.IntOfInt64(356))).Keys().Elements()); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _15_v1 _dafny.Int
					_15_v1 = interface{}(_compr_6).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality()), _dafny.IntOfInt64(356))).Contains(_15_v1) {
						_coll6.Add(Companion_Default___.SafeDivisionInt(_15_v1, _dafny.IntOfInt64(-91)))
					}
				}
				return _coll6.ToSet()
			}()).Contains(_14_v0) {
				_coll4.Add((_14_v0).Times((_dafny.SetOf(false, !(true), false, true)).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(972))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}(func(_16_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				})))
			}
		}
		return _coll4.ToMap()
	}()).Cardinality()).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("delwixswv")).Cardinality()))) <= 0 {
		return _dafny.CodePoint('v')
	} else {
		return _dafny.CodePoint('m')
	}
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((Companion_D7_.Create_DC25_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-252))).Uint32(), func(coer11 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_17_i0 _dafny.Int) D3 {
		return Companion_D3_.Create_DC9_(_dafny.SeqOf(_dafny.CodePoint('a'), _dafny.CodePoint('f')))
	})))).Dtor_cf41(), _dafny.SeqOf(Companion_D3_.Create_DC9_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(356))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_18_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))), Companion_D3_.Create_DC9_(_dafny.SeqOf(_dafny.CodePoint('u')))))
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(599))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_19_i0 _dafny.Int) _dafny.Map {
		return func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dxqwlk")).Cardinality()), _dafny.CodePoint('w'))).Keys().Elements()); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _20_v0 _dafny.Int
				_20_v0 = interface{}(_compr_7).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dxqwlk")).Cardinality()), _dafny.CodePoint('w'))).Contains(_20_v0) {
					_coll7.Add((_20_v0).Times(_dafny.IntOfInt64(757)), true)
				}
			}
			return _coll7.ToMap()
		}()
	})))
}
func (_static *CompanionStruct_Default___) Fm21(p0 bool, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC20_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), !(false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 _dafny.Map, globalState *GlobalState) _dafny.Int {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var _hi0 _dafny.Int = p0
	_ = _hi0
	for _21_i0 := p0; _21_i0.Cmp(_hi0) < 0; _21_i0 = _21_i0.Plus(_dafny.One) {
		var _22_v0 _dafny.CodePoint
		_ = _22_v0
		_22_v0 = _dafny.CodePoint('x')
		var _23_v1 _dafny.Map
		_ = _23_v1
		_23_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_22_v0, _22_v0)
		var _24_v2 _dafny.Map
		_ = _24_v2
		_24_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_23_v1, p2)
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((p1), 0))
		_ = _index0
		(p1).ArraySet1(p0, (_index0).Int())
		var _25_v3 _dafny.Sequence
		_ = _25_v3
		_25_v3 = _dafny.UnicodeSeqOfUtf8Bytes("nodnggll")
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((p1), 0))
		_ = _index1
		var _rhs0 _dafny.Map = _24_v2
		_ = _rhs0
		var _rhs1 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(246))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_26_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_27_i1 _dafny.Int) _dafny.CodePoint {
				return _26_v0
			}
		})(_22_v0)))).Cardinality())
		_ = _rhs1
		var _rhs2 _dafny.Int = _dafny.IntOfUint32((_25_v3).Cardinality())
		_ = _rhs2
		var _lhs0 _dafny.Array = p1
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((p1), 0))
		_ = _lhs1
		_24_v2 = _rhs0
		r0 = _rhs1
		(_lhs0).ArraySet1(_rhs2, (_lhs1).Int())
		var _28_v4 D5
		_ = _28_v4
		_28_v4 = Companion_D5_.Create_DC18_(p0, p2)
		var _pat_let_tv0 = p2
		_ = _pat_let_tv0
		var _pat_let_tv1 = p0
		_ = _pat_let_tv1
		var _pat_let_tv2 = p2
		_ = _pat_let_tv2
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((p1), 0))
		_ = _index2
		var _rhs3 _dafny.Int = p0
		_ = _rhs3
		var _rhs4 _dafny.Int = Companion_Default___.SafeDivisionInt((_28_v4).Dtor_cf27(), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))
		_ = _rhs4
		var _rhs5 D5 = func(_pat_let0_0 D5) D5 {
			return func(_33_dt__update__tmp_h2 D5) D5 {
				return func(_pat_let5_0 bool) D5 {
					return func(_34_dt__update_hcf28_h1 bool) D5 {
						return Companion_D5_.Create_DC18_((_33_dt__update__tmp_h2).Dtor_cf27(), _34_dt__update_hcf28_h1)
					}(_pat_let5_0)
				}(_pat_let_tv2)
			}(_pat_let0_0)
		}(func(_pat_let1_0 D5) D5 {
			return func(_31_dt__update__tmp_h1 D5) D5 {
				return func(_pat_let4_0 _dafny.Int) D5 {
					return func(_32_dt__update_hcf27_h0 _dafny.Int) D5 {
						return Companion_D5_.Create_DC18_(_32_dt__update_hcf27_h0, (_31_dt__update__tmp_h1).Dtor_cf28())
					}(_pat_let4_0)
				}((_dafny.Zero).Minus(_pat_let_tv1))
			}(_pat_let1_0)
		}(func(_pat_let2_0 D5) D5 {
			return func(_29_dt__update__tmp_h0 D5) D5 {
				return func(_pat_let3_0 bool) D5 {
					return func(_30_dt__update_hcf28_h0 bool) D5 {
						return Companion_D5_.Create_DC18_((_29_dt__update__tmp_h0).Dtor_cf27(), _30_dt__update_hcf28_h0)
					}(_pat_let3_0)
				}(_pat_let_tv0)
			}(_pat_let2_0)
		}(_28_v4)))
		_ = _rhs5
		var _lhs2 _dafny.Array = p1
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((p1), 0))
		_ = _lhs3
		(_lhs2).ArraySet1(_rhs3, (_lhs3).Int())
		r0 = _rhs4
		_28_v4 = _rhs5
		var _35_v5 *C1
		_ = _35_v5
		var _nw0 *C1 = New_C1_()
		_ = _nw0
		_nw0.Ctor__(p2, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), p2)
		_35_v5 = _nw0
		var _36_v6 _dafny.Array
		_ = _36_v6
		var _nwElement0_0 *C1 = _35_v5
		_ = _nwElement0_0
		var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(5))
		_ = _nw1
		(_nw1).ArraySet1(_nwElement0_0, 0)
		(_nw1).ArraySet1(_35_v5, 1)
		(_nw1).ArraySet1(_35_v5, 2)
		(_nw1).ArraySet1(_35_v5, 3)
		(_nw1).ArraySet1(_35_v5, 4)
		_36_v6 = _nw1
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_36_v6), 0))
		_ = _index3
		(_36_v6).ArraySet1(_35_v5, (_index3).Int())
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_36_v6), 0))
		_ = _index4
		(_36_v6).ArraySet1(_35_v5, (_index4).Int())
		(globalState).F3 = !(_dafny.Companion_Sequence_.Contains(_25_v3, _22_v0))
	}
	var _hi1 _dafny.Int = p0
	_ = _hi1
	for _37_i2 := p0; _37_i2.Cmp(_hi1) < 0; _37_i2 = _37_i2.Plus(_dafny.One) {
		var _38_v7 _dafny.Array
		_ = _38_v7
		var _nwElement0_1 bool = p2
		_ = _nwElement0_1
		var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(14))
		_ = _nw2
		(_nw2).ArraySet1(_nwElement0_1, 0)
		(_nw2).ArraySet1(true, 1)
		(_nw2).ArraySet1(p2, 2)
		(_nw2).ArraySet1(p2, 3)
		(_nw2).ArraySet1(Companion_Default___.Fm2(globalState), 4)
		(_nw2).ArraySet1(p2, 5)
		(_nw2).ArraySet1(p2, 6)
		(_nw2).ArraySet1(p2, 7)
		(_nw2).ArraySet1(p2, 8)
		(_nw2).ArraySet1(p2, 9)
		(_nw2).ArraySet1(p2, 10)
		(_nw2).ArraySet1(true, 11)
		(_nw2).ArraySet1(p2, 12)
		(_nw2).ArraySet1(p2, 13)
		_38_v7 = _nw2
		var _39_v8 _dafny.Sequence
		_ = _39_v8
		_39_v8 = _dafny.SeqOf(_38_v7, _38_v7)
		var _40_v9 _dafny.Array
		_ = _40_v9
		var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(14))
		_ = _nw3
		_40_v9 = _nw3
		var _41_v10 _dafny.Map
		_ = _41_v10
		_41_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_39_v8, _40_v9)
		_41_v10 = (_41_v10).Update(_dafny.Companion_Sequence_.Concatenate(_39_v8, _39_v8), _40_v9)
		var _42_v11 _dafny.Sequence
		_ = _42_v11
		_42_v11 = _dafny.SeqOf(_37_i2)
		var _43_v12 _dafny.Set
		_ = _43_v12
		_43_v12 = _dafny.SetOf((_dafny.Zero).Minus(p0), p0, (_42_v11).Select((Companion_Default___.SafeIndex(_37_i2, _dafny.IntOfUint32((_42_v11).Cardinality()))).Uint32()).(_dafny.Int), p0, _37_i2)
		var _44_v13 D3
		_ = _44_v13
		_44_v13 = Companion_D3_.Create_DC10_((_43_v12).Cardinality(), p1, p2)
		var _source0 D3 = _44_v13
		_ = _source0
		if _source0.Is_DC10() {
			var _45___mcc_h0 _dafny.Int = _source0.Get_().(D3_DC10).Cf11
			_ = _45___mcc_h0
			var _46___mcc_h1 _dafny.Array = _source0.Get_().(D3_DC10).Cf12
			_ = _46___mcc_h1
			var _47___mcc_h2 bool = _source0.Get_().(D3_DC10).Cf13
			_ = _47___mcc_h2
			var _48_cf13 bool = _47___mcc_h2
			_ = _48_cf13
			var _49_cf12 _dafny.Array = _46___mcc_h1
			_ = _49_cf12
			var _50_cf11 _dafny.Int = _45___mcc_h0
			_ = _50_cf11
			var _51_v14 _dafny.Map
			_ = _51_v14
			_51_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_49_cf12), 0))
			_ = _index5
			(_49_cf12).ArraySet1((func() _dafny.Int {
				if p2 {
					return _50_cf11
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_51_v14).Cardinality())).Cardinality()
			})(), (_index5).Int())
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_49_cf12), 0))
			_ = _index6
			(_49_cf12).ArraySet1((func() _dafny.Int {
				if false {
					return (_37_i2).Times(p0)
				}
				return p0
			})(), (_index6).Int())
			var _52_v15 *C2
			_ = _52_v15
			var _nw4 *C2 = New_C2_()
			_ = _nw4
			_nw4.Ctor__(true, _48_cf13)
			_52_v15 = _nw4
			_52_v15 = _52_v15
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_38_v7), 0))
			_ = _index7
			(_38_v7).ArraySet1(_52_v15.F10, (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_38_v7), 0))
			_ = _index8
			(_38_v7).ArraySet1(p2, (_index8).Int())
			var _53_v16 _dafny.Array
			_ = _53_v16
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw5
			_53_v16 = _nw5
			var _54_v17 D0
			_ = _54_v17
			_54_v17 = Companion_D0_.Create_DC1_(_52_v15.F10)
			var _55_v18 _dafny.Set
			_ = _55_v18
			_55_v18 = _dafny.SetOf((_54_v17).Dtor_cf1(), _52_v15.F10)
			var _56_v19 bool
			_ = _56_v19
			var _57_v20 bool
			_ = _57_v20
			var _58_v21 _dafny.Map
			_ = _58_v21
			var _out0 bool
			_ = _out0
			var _out1 bool
			_ = _out1
			var _out2 _dafny.Map
			_ = _out2
			_out0, _out1, _out2 = (_52_v15).M3((p0).Minus((_49_cf12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_49_cf12), 0))).Int()).(_dafny.Int)), _53_v16, true, _55_v18, globalState)
			_56_v19 = _out0
			_57_v20 = _out1
			_58_v21 = _out2
		} else if _source0.Is_DC11() {
			var _59___mcc_h3 _dafny.Int = _source0.Get_().(D3_DC11).Cf14
			_ = _59___mcc_h3
			var _60___mcc_h4 _dafny.Sequence = _source0.Get_().(D3_DC11).Cf15
			_ = _60___mcc_h4
			var _61___mcc_h5 _dafny.Int = _source0.Get_().(D3_DC11).Cf16
			_ = _61___mcc_h5
			var _62___mcc_h6 *C0 = _source0.Get_().(D3_DC11).Cf17
			_ = _62___mcc_h6
			var _63_cf17 *C0 = _62___mcc_h6
			_ = _63_cf17
			var _64_cf16 _dafny.Int = _61___mcc_h5
			_ = _64_cf16
			var _65_cf15 _dafny.Sequence = _60___mcc_h4
			_ = _65_cf15
			var _66_cf14 _dafny.Int = _59___mcc_h3
			_ = _66_cf14
			var _67_v23 _dafny.MultiSet
			_ = _67_v23
			_67_v23 = _dafny.MultiSetOf(func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-756), _dafny.IntOfInt64(716))); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _68_v22 _dafny.Int
					_68_v22 = interface{}(_compr_8).(_dafny.Int)
					if ((_dafny.IntOfInt64(-756)).Cmp(_68_v22) <= 0) && ((_68_v22).Cmp(_dafny.IntOfInt64(716)) < 0) {
						_coll8.Add(Companion_Default___.SafeModuloInt(_68_v22, _dafny.IntOfUint32((_65_cf15).Cardinality())), !(p2))
					}
				}
				return _coll8.ToMap()
			}())
			var _69_v24 D2
			_ = _69_v24
			_69_v24 = Companion_D2_.Create_DC8_((_63_cf17).F13(), _64_cf16, (_63_cf17).F13())
			var _pat_let_tv3 = _63_cf17
			_ = _pat_let_tv3
			var _pat_let_tv4 = _64_cf16
			_ = _pat_let_tv4
			var _pat_let_tv5 = _63_cf17
			_ = _pat_let_tv5
			var _rhs6 bool = (func(_pat_let6_0 D2) D2 {
				return func(_73_dt__update__tmp_h4 D2) D2 {
					return func(_pat_let10_0 bool) D2 {
						return func(_74_dt__update_hcf7_h1 bool) D2 {
							return Companion_D2_.Create_DC8_(_74_dt__update_hcf7_h1, (_73_dt__update__tmp_h4).Dtor_cf8(), (_73_dt__update__tmp_h4).Dtor_cf9())
						}(_pat_let10_0)
					}((_pat_let_tv5).F13())
				}(_pat_let6_0)
			}(func(_pat_let7_0 D2) D2 {
				return func(_70_dt__update__tmp_h3 D2) D2 {
					return func(_pat_let8_0 bool) D2 {
						return func(_71_dt__update_hcf7_h0 bool) D2 {
							return func(_pat_let9_0 _dafny.Int) D2 {
								return func(_72_dt__update_hcf8_h0 _dafny.Int) D2 {
									return Companion_D2_.Create_DC8_(_71_dt__update_hcf7_h0, _72_dt__update_hcf8_h0, (_70_dt__update__tmp_h3).Dtor_cf9())
								}(_pat_let9_0)
							}(_pat_let_tv4)
						}(_pat_let8_0)
					}((_pat_let_tv3).F13())
				}(_pat_let7_0)
			}(_69_v24))).Dtor_cf7()
			_ = _rhs6
			var _rhs7 _dafny.MultiSet = Companion_Default___.Fm20((_63_cf17).F13(), _42_v11, (_64_cf16).Cmp(_63_cf17.F14) == 0, globalState)
			_ = _rhs7
			var _rhs8 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_65_cf15, _dafny.Companion_Sequence_.Concatenate(_65_cf15, _65_cf15))
			_ = _rhs8
			var _lhs4 *GlobalState = globalState
			_ = _lhs4
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			_lhs4.F3 = _rhs6
			_67_v23 = _rhs7
			_lhs5.F3 = _rhs8
			var _75_v25 _dafny.Array
			_ = _75_v25
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw6
			_75_v25 = _nw6
			_75_v25 = p1
			_64_cf16 = (func() _dafny.Int {
				if (func() bool {
					if (_63_cf17).F13() {
						return p2
					}
					return (_63_cf17).F13()
				})() {
					return _37_i2
				}
				return _66_cf14
			})()
			var _76_v26 _dafny.Map
			_ = _76_v26
			_76_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_63_cf17).F13())
			var _77_v27 _dafny.Map
			_ = _77_v27
			_77_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_76_v26).Cardinality()).Cmp(_66_cf14) <= 0, !(p2))
			var _78_v28 _dafny.MultiSet
			_ = _78_v28
			_78_v28 = _dafny.MultiSetOf(p2)
			_77_v27 = (_77_v27).Update(!(_78_v28).Contains(p2), (_44_v13).Dtor_cf13())
		} else if _source0.Is_DC9() {
			var _79___mcc_h7 _dafny.Sequence = _source0.Get_().(D3_DC9).Cf10
			_ = _79___mcc_h7
			var _80_cf10 _dafny.Sequence = _79___mcc_h7
			_ = _80_cf10
			var _81_v29 _dafny.Set
			_ = _81_v29
			_81_v29 = _dafny.SetOf(p2, p2)
			(globalState).F0 = !(((_42_v11).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_42_v11).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((p0).Times((_81_v29).Cardinality())) >= 0)
			var _82_v30 _dafny.Array
			_ = _82_v30
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_0
			var _nw7 _dafny.Array
			_ = _nw7
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw7 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.CodePoint = (func(_83_p2 bool) func(_dafny.Int) _dafny.CodePoint {
					return func(_84_i3 _dafny.Int) _dafny.CodePoint {
						return (func() _dafny.CodePoint {
							if _83_p2 {
								return _dafny.CodePoint('c')
							}
							return _dafny.CodePoint('x')
						})()
					}
				})(p2)
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
			_82_v30 = _nw7
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_82_v30), 0))
			_ = _index9
			(_82_v30).ArraySet1CodePoint(_dafny.CodePoint('q'), (_index9).Int())
			var _85_v31 _dafny.CodePoint
			_ = _85_v31
			_85_v31 = _dafny.CodePoint('u')
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_82_v30), 0))
			_ = _index10
			(_82_v30).ArraySet1CodePoint(_85_v31, (_index10).Int())
			_40_v9 = _40_v9
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((p1), 0))
			_ = _index11
			(p1).ArraySet1(p0, (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((p1), 0))
			_ = _index12
			(p1).ArraySet1(p0, (_index12).Int())
		} else {
			var _86___mcc_h8 D3 = _source0.Get_().(D3_DC12).Cf18
			_ = _86___mcc_h8
			var _87_cf18 D3 = _86___mcc_h8
			_ = _87_cf18
			_38_v7 = _38_v7
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_38_v7), 0))
			_ = _index13
			(_38_v7).ArraySet1(!(p2), (_index13).Int())
			var _88_v32 _dafny.MultiSet
			_ = _88_v32
			_88_v32 = _dafny.MultiSetOf(p2, p2)
			var _89_v33 _dafny.Sequence
			_ = _89_v33
			_89_v33 = _dafny.SeqOf(_88_v32)
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_38_v7), 0))
			_ = _index14
			(_38_v7).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_89_v33, _89_v33), _89_v33), (_index14).Int())
			r0 = p0
			var _90_v34 _dafny.Array
			_ = _90_v34
			var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
			_ = _nw8
			_90_v34 = _nw8
			var _91_v35 _dafny.Array
			_ = _91_v35
			var _nw9 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw9
			_91_v35 = _nw9
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_90_v34), 0))
			_ = _index15
			(_90_v34).ArraySet1(_91_v35, (_index15).Int())
			var _92_v36 _dafny.CodePoint
			_ = _92_v36
			_92_v36 = _dafny.CodePoint('l')
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_90_v34), 0))
			_ = _index16
			(_90_v34).ArraySet1((_39_v8).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm3(_92_v36, true, p0, globalState), _dafny.IntOfUint32((_39_v8).Cardinality()))).Uint32()).(_dafny.Array), (_index16).Int())
		}
		var _93_v37 D4
		_ = _93_v37
		_93_v37 = Companion_D4_.Create_DC15_((p0).Plus(_37_i2), p2)
		var _rhs9 bool = p2
		_ = _rhs9
		var _rhs10 D4 = _93_v37
		_ = _rhs10
		var _lhs6 *GlobalState = globalState
		_ = _lhs6
		_lhs6.F3 = _rhs9
		_93_v37 = _rhs10
		var _94_v38 _dafny.Sequence
		_ = _94_v38
		_94_v38 = _dafny.SeqOf(p2, p2)
		var _95_v39 *C2
		_ = _95_v39
		var _nw10 *C2 = New_C2_()
		_ = _nw10
		_nw10.Ctor__((_94_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_94_v38).Cardinality()), _dafny.IntOfUint32((_94_v38).Cardinality()))).Uint32()).(bool), p2)
		_95_v39 = _nw10
	}
	var _hi2 _dafny.Int = p0
	_ = _hi2
	for _96_i4 := (p0).Plus(p0); _96_i4.Cmp(_hi2) < 0; _96_i4 = _96_i4.Plus(_dafny.One) {
		(globalState).F0 = !(p2)
		r0 = (_96_i4).Minus(_dafny.IntOfUint32((Companion_Default___.Fm11(false, p2, globalState)).Cardinality()))
		(globalState).F0 = ((func() bool {
			if !(!(p2)) {
				return p2
			}
			return p2
		})()) && (!((_96_i4).Cmp(p0) != 0))
		(globalState).F3 = p2
	}
	var _hi3 _dafny.Int = p0
	_ = _hi3
	for _97_i5 := p0; _97_i5.Cmp(_hi3) < 0; _97_i5 = _97_i5.Plus(_dafny.One) {
		var _98_v40 _dafny.Sequence
		_ = _98_v40
		_98_v40 = _dafny.SeqOf(p2, p2, p2)
		var _99_v41 *C1
		_ = _99_v41
		var _nw11 *C1 = New_C1_()
		_ = _nw11
		_nw11.Ctor__(!(p2) || (Companion_Default___.Fm2(globalState)), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_98_v40).Cardinality()), _97_i5), (func() bool {
			if p2 {
				return p2
			}
			return p2
		})())
		_99_v41 = _nw11
		_99_v41 = _99_v41
		var _100_v42 _dafny.Array
		_ = _100_v42
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_1
		var _nw12 _dafny.Array
		_ = _nw12
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw12 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_101_v41 *C1) func(_dafny.Int) bool {
				return func(_102_i6 _dafny.Int) bool {
					return (_101_v41).F11()
				}
			})(_99_v41)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw12 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw12).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw12).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_100_v42 = _nw12
		var _103_v43 _dafny.Map
		_ = _103_v43
		_103_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v42, (_99_v41).F11())
		_103_v43 = (_103_v43).Update(_100_v42, p2)
		var _104_v44 _dafny.MultiSet
		_ = _104_v44
		_104_v44 = _dafny.MultiSetOf(_dafny.SetOf(p0))
		(globalState).F3 = ((_104_v44).IsSubsetOf(_104_v44)) == ((_99_v41).F11())
		var _105_v45 *C3
		_ = _105_v45
		var _nw13 *C3 = New_C3_()
		_ = _nw13
		_nw13.Ctor__((_99_v41).F11())
		_105_v45 = _nw13
		var _106_v46 _dafny.Map
		_ = _106_v46
		_106_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v42, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v45, (_99_v41).F11()))
		var _107_v47 _dafny.Map
		_ = _107_v47
		_107_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v45, (_99_v41).F11())
		_106_v46 = (_106_v46).Update(_100_v42, _107_v47)
	}
	var _108_i7 _dafny.Int
	_ = _108_i7
	_108_i7 = _dafny.Zero
	{
		for !(!(p2)) {
			{
				if (_108_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_108_i7 = (_108_i7).Plus(_dafny.One)
				(globalState).F0 = p2
				var _109_v48 _dafny.MultiSet
				_ = _109_v48
				_109_v48 = _dafny.MultiSetOf(p0)
				var _110_v49 _dafny.MultiSet
				_ = _110_v49
				_110_v49 = _dafny.MultiSetOf(_109_v48)
				var _111_v50 D0
				_ = _111_v50
				_111_v50 = Companion_D0_.Create_DC2_(p2, p0, _110_v49)
				var _112_v51 _dafny.Array
				_ = _112_v51
				var _nwElement0_2 bool = p2
				_ = _nwElement0_2
				var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(21))
				_ = _nw14
				(_nw14).ArraySet1(_nwElement0_2, 0)
				(_nw14).ArraySet1(p2, 1)
				(_nw14).ArraySet1(p2, 2)
				(_nw14).ArraySet1(p2, 3)
				(_nw14).ArraySet1(p2, 4)
				(_nw14).ArraySet1(p2, 5)
				(_nw14).ArraySet1(p2, 6)
				(_nw14).ArraySet1(p2, 7)
				(_nw14).ArraySet1(p2, 8)
				(_nw14).ArraySet1(p2, 9)
				(_nw14).ArraySet1(p2, 10)
				(_nw14).ArraySet1(p2, 11)
				(_nw14).ArraySet1(p2, 12)
				(_nw14).ArraySet1(p2, 13)
				(_nw14).ArraySet1(!(p2), 14)
				(_nw14).ArraySet1((_111_v50).Dtor_cf2(), 15)
				(_nw14).ArraySet1(p2, 16)
				(_nw14).ArraySet1(true, 17)
				(_nw14).ArraySet1(p2, 18)
				(_nw14).ArraySet1(p2, 19)
				(_nw14).ArraySet1(p2, 20)
				_112_v51 = _nw14
				var _113_v52 _dafny.Sequence
				_ = _113_v52
				_113_v52 = _dafny.SeqOf(_112_v51, _112_v51, _112_v51)
				var _114_v53 _dafny.Sequence
				_ = _114_v53
				_114_v53 = _dafny.UnicodeSeqOfUtf8Bytes("nw")
				var _115_v54 _dafny.Sequence
				_ = _115_v54
				_115_v54 = _dafny.SeqOf(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_113_v52, _dafny.Companion_Sequence_.Update(_113_v52, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_114_v53).Cardinality()), _dafny.IntOfUint32((_113_v52).Cardinality()))).Uint32(), _112_v51))).Cardinality()))
				var _116_v55 _dafny.Sequence
				_ = _116_v55
				_116_v55 = _dafny.SeqOf(_115_v54)
				_115_v54 = (_116_v55).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_116_v55).Cardinality()))).Uint32()).(_dafny.Sequence)
				var _117_v56 _dafny.Sequence
				_ = _117_v56
				_117_v56 = _dafny.SeqOf(p2, p2)
				if _dafny.Companion_Sequence_.Equal(Companion_Default___.Fm10(globalState), _117_v56) {
					(globalState).F0 = p2
					var _118_v57 _dafny.Map
					_ = _118_v57
					_118_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Cmp(p0) == 0, p2)
					var _119_v58 _dafny.CodePoint
					_ = _119_v58
					_119_v58 = _dafny.CodePoint('k')
					var _120_v59 _dafny.Map
					_ = _120_v59
					_120_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _119_v58)
					var _rhs11 _dafny.Array = _112_v51
					_ = _rhs11
					var _rhs12 _dafny.Sequence = Companion_Default___.Fm13(p2, (func() bool {
						if p2 {
							return p2
						}
						return p2
					})(), p2, p0, globalState)
					_ = _rhs12
					var _rhs13 _dafny.Map = (Companion_Default___.Fm21(p2, globalState)).Dtor_cf30()
					_ = _rhs13
					var _rhs14 bool = p2
					_ = _rhs14
					var _rhs15 _dafny.Map = _120_v59
					_ = _rhs15
					var _lhs7 *GlobalState = globalState
					_ = _lhs7
					_112_v51 = _rhs11
					_115_v54 = _rhs12
					_118_v57 = _rhs13
					_lhs7.F3 = _rhs14
					_120_v59 = _rhs15
					var _121_v60 _dafny.Array
					_ = _121_v60
					var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
					_ = _nw15
					_121_v60 = _nw15
					var _122_v61 _dafny.Array
					_ = _122_v61
					var _nw16 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
					_ = _nw16
					_122_v61 = _nw16
					var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_121_v60), 0))
					_ = _index17
					(_121_v60).ArraySet1(_122_v61, (_index17).Int())
					var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_121_v60), 0))
					_ = _index18
					(_121_v60).ArraySet1(_122_v61, (_index18).Int())
					var _123_v62 _dafny.MultiSet
					_ = _123_v62
					_123_v62 = _dafny.MultiSetOf(p2, p2, p2, p2, p2)
					var _124_v63 _dafny.Set
					_ = _124_v63
					_124_v63 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(p0, (_123_v62).Cardinality())).Cardinality()))
					var _125_v64 *C0
					_ = _125_v64
					var _nw17 *C0 = New_C0_()
					_ = _nw17
					_nw17.Ctor__(true, (_115_v54).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_115_v54).Cardinality()))).Uint32()).(_dafny.Int), (_124_v63).IsProperSubsetOf(_124_v63))
					_125_v64 = _nw17
					_125_v64 = _125_v64
					var _126_v65 _dafny.Set
					_ = _126_v65
					_126_v65 = _dafny.SetOf(_109_v48, _109_v48, (_dafny.MultiSetFromSeq(_115_v54)).Union(_109_v48))
					var _127_v66 _dafny.Array
					_ = _127_v66
					var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
					_ = _nw18
					_127_v66 = _nw18
					var _128_v67 _dafny.Map
					_ = _128_v67
					_128_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_125_v64).F13(), _117_v56)
					var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_127_v66), 0))
					_ = _index19
					(_127_v66).ArraySet1(_128_v67, (_index19).Int())
					var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_127_v66), 0))
					_ = _index20
					var _rhs16 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, (_dafny.Zero).Minus(p0))
					_ = _rhs16
					var _rhs17 _dafny.Set = _126_v65
					_ = _rhs17
					var _rhs18 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_114_v53, _114_v53)
					_ = _rhs18
					var _rhs19 _dafny.MultiSet = _dafny.MultiSetOf(_125_v64.F14, Companion_Default___.SafeModuloInt(_125_v64.F14, _125_v64.F14))
					_ = _rhs19
					var _rhs20 _dafny.Map = ((_128_v67).Merge(_128_v67)).Merge(_128_v67)
					_ = _rhs20
					var _lhs8 *C0 = _125_v64
					_ = _lhs8
					var _lhs9 _dafny.Array = _127_v66
					_ = _lhs9
					var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_127_v66), 0))
					_ = _lhs10
					_lhs8.F14 = _rhs16
					_126_v65 = _rhs17
					_114_v53 = _rhs18
					_109_v48 = _rhs19
					(_lhs9).ArraySet1(_rhs20, (_lhs10).Int())
				} else {
					var _129_v68 _dafny.MultiSet
					_ = _129_v68
					_129_v68 = _dafny.MultiSetOf(_115_v54, _dafny.SeqOf(p0))
					(globalState).F0 = !((p0).Cmp((_129_v68).Cardinality()) >= 0)
					_115_v54 = _115_v54
					r0 = (p0).Times(p0)
					var _130_v69 _dafny.CodePoint
					_ = _130_v69
					_130_v69 = _dafny.CodePoint('l')
					var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((p1), 0))
					_ = _index21
					(p1).ArraySet1((_dafny.IntOfUint32((_114_v53).Cardinality())).Plus(Companion_Default___.Fm3(_130_v69, true, _dafny.IntOfUint32((_114_v53).Cardinality()), globalState)), (_index21).Int())
					var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((p1), 0))
					_ = _index22
					(p1).ArraySet1(p0, (_index22).Int())
					var _131_v70 *C0
					_ = _131_v70
					var _nw19 *C0 = New_C0_()
					_ = _nw19
					_nw19.Ctor__(false, p0, p2)
					_131_v70 = _nw19
					var _132_v71 T0
					_ = _132_v71
					var _nw20 *C3 = New_C3_()
					_ = _nw20
					_nw20.Ctor__((_131_v70).F13())
					_132_v71 = _nw20
					var _133_v72 _dafny.Map
					_ = _133_v72
					_133_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v70, _132_v71)
					var _134_v73 _dafny.Map
					_ = _134_v73
					_134_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_133_v72).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v70, _132_v71)))
					_134_v73 = (_134_v73).Update(!((_131_v70).F13()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v70, _132_v71))
				}
				var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_112_v51), 0))
				_ = _index23
				(_112_v51).ArraySet1(p2, (_index23).Int())
				var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_112_v51), 0))
				_ = _index24
				(_112_v51).ArraySet1(((p0).Cmp(p0) >= 0) || (p2), (_index24).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _hi4 _dafny.Int = p0
	_ = _hi4
	for _135_i8 := _dafny.IntOfInt64(702); _135_i8.Cmp(_hi4) < 0; _135_i8 = _135_i8.Plus(_dafny.One) {
		(globalState).F0 = (_135_i8).Cmp(_135_i8) <= 0
		var _136_v74 _dafny.Array
		_ = _136_v74
		var _nw21 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
		_ = _nw21
		_136_v74 = _nw21
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_136_v74), 0))
		_ = _index25
		(_136_v74).ArraySet1(p2, (_index25).Int())
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_136_v74), 0))
		_ = _index26
		(_136_v74).ArraySet1(false, (_index26).Int())
		var _137_v75 _dafny.Array
		_ = _137_v75
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(21))
		_ = _nw22
		_137_v75 = _nw22
		var _138_v76 _dafny.Sequence
		_ = _138_v76
		_138_v76 = _dafny.SeqOf(p2)
		var _139_v77 _dafny.Set
		_ = _139_v77
		_139_v77 = _dafny.SetOf(p2, p2, p2)
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_137_v75), 0))
		_ = _index27
		(_137_v75).ArraySet1((_dafny.SetOf(p2, p2, (_138_v76).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_138_v76).Cardinality()))).Uint32()).(bool), p2)).Union(_139_v77), (_index27).Int())
		var _140_v78 T0
		_ = _140_v78
		var _nw23 *C1 = New_C1_()
		_ = _nw23
		_nw23.Ctor__(p2, Companion_Default___.SafeModuloInt(_135_i8, p0), p2)
		_140_v78 = _nw23
		var _141_v79 _dafny.Set
		_ = _141_v79
		_141_v79 = _dafny.SetOf(_135_i8, p0, p0)
		var _142_v80 _dafny.MultiSet
		_ = _142_v80
		_142_v80 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hdndrwawf")).Cardinality()))
		var _143_v81 _dafny.Map
		_ = _143_v81
		_143_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Int {
			if (_142_v80).Contains(_135_i8) {
				return (_142_v80).Multiplicity(_135_i8)
			}
			return _135_i8
		})()).Minus(_135_i8), (p0).Times(_dafny.IntOfInt64(-532)))
		var _144_v82 _dafny.Sequence
		_ = _144_v82
		_144_v82 = _dafny.UnicodeSeqOfUtf8Bytes("koccov")
		var _145_v83 _dafny.Map
		_ = _145_v83
		_145_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.SeqOf(true, true, p2, (_140_v78).F9()))
		var _146_v84 *C2
		_ = _146_v84
		var _nw24 *C2 = New_C2_()
		_ = _nw24
		_nw24.Ctor__((_140_v78).F9(), (_140_v78).F9())
		_146_v84 = _nw24
		var _147_v85 _dafny.MultiSet
		_ = _147_v85
		_147_v85 = _dafny.MultiSetOf(_146_v84)
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_136_v74), 0))
		_ = _index28
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_136_v74), 0))
		_ = _index29
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_137_v75), 0))
		_ = _index30
		var _rhs21 bool = (_141_v79).IsProperSubsetOf(_141_v79)
		_ = _rhs21
		var _rhs22 bool = p2
		_ = _rhs22
		var _rhs23 _dafny.Int = (func() _dafny.Int {
			if (_143_v81).Contains(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfInt64(-199)), _dafny.IntOfUint32((_144_v82).Cardinality()))) {
				return (_143_v81).Get(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfInt64(-199)), _dafny.IntOfUint32((_144_v82).Cardinality()))).(_dafny.Int)
			}
			return ((_140_v78).Fm4(_145_v83, p2, (_140_v78).F9(), p0, globalState)).Minus((_147_v85).Cardinality())
		})()
		_ = _rhs23
		var _rhs24 _dafny.Set = _139_v77
		_ = _rhs24
		var _rhs25 T0 = _140_v78
		_ = _rhs25
		var _lhs11 _dafny.Array = _136_v74
		_ = _lhs11
		var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_136_v74), 0))
		_ = _lhs12
		var _lhs13 _dafny.Array = _136_v74
		_ = _lhs13
		var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_136_v74), 0))
		_ = _lhs14
		var _lhs15 _dafny.Array = _137_v75
		_ = _lhs15
		var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_137_v75), 0))
		_ = _lhs16
		(_lhs11).ArraySet1(_rhs21, (_lhs12).Int())
		(_lhs13).ArraySet1(_rhs22, (_lhs14).Int())
		r0 = _rhs23
		(_lhs15).ArraySet1(_rhs24, (_lhs16).Int())
		_140_v78 = _rhs25
		var _148_v86 _dafny.Array
		_ = _148_v86
		var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
		_ = _nw25
		_148_v86 = _nw25
		var _149_v87 D4
		_ = _149_v87
		_149_v87 = Companion_D4_.Create_DC13_(_148_v86)
		var _150_v88 D4
		_ = _150_v88
		_150_v88 = Companion_D4_.Create_DC16_(_149_v87)
		_150_v88 = _150_v88
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_136_v74), 0))
		_ = _index31
		(_136_v74).ArraySet1((_135_i8).Cmp(_dafny.IntOfInt64(-200)) == 0, (_index31).Int())
	}
	r0 = Companion_Default___.SafeDivisionInt((p0).Times(p0), p0)
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _151_v0 _dafny.Int
	_ = _151_v0
	_151_v0 = _dafny.IntOfInt64(636)
	var _152_v1 _dafny.Sequence
	_ = _152_v1
	_152_v1 = _dafny.SeqOf(_151_v0, _151_v0)
	var _153_v2 bool
	_ = _153_v2
	_153_v2 = true
	var _154_v3 _dafny.Map
	_ = _154_v3
	_154_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v2, _152_v1)
	var _155_v4 _dafny.MultiSet
	_ = _155_v4
	_155_v4 = _dafny.MultiSetOf(_152_v1, _152_v1, (func() _dafny.Sequence {
		if (_154_v3).Contains(!(false)) {
			return (_154_v3).Get(!(false)).(_dafny.Sequence)
		}
		return _dafny.SeqOf(_151_v0)
	})(), _152_v1, _152_v1)
	var _156_v5 _dafny.Array
	_ = _156_v5
	var _nw26 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
	_ = _nw26
	_156_v5 = _nw26
	var _157_v6 _dafny.MultiSet
	_ = _157_v6
	_157_v6 = _dafny.MultiSetOf(_156_v5, _156_v5)
	var _158_v7 D0
	_ = _158_v7
	_158_v7 = Companion_D0_.Create_DC0_((_dafny.Zero).Minus(_151_v0))
	var _159_v8 _dafny.Sequence
	_ = _159_v8
	_159_v8 = _dafny.UnicodeSeqOfUtf8Bytes("iyuhsbbd")
	var _160_v9 _dafny.Map
	_ = _160_v9
	_160_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v2, _151_v0)
	var _161_v10 _dafny.Array
	_ = _161_v10
	var _nwElement0_3 _dafny.Int = _151_v0
	_ = _nwElement0_3
	var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(17))
	_ = _nw27
	(_nw27).ArraySet1(_nwElement0_3, 0)
	(_nw27).ArraySet1(_151_v0, 1)
	(_nw27).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yqsopcbd")).Cardinality()))).Cardinality(), 2)
	(_nw27).ArraySet1((_dafny.Zero).Minus(_151_v0), 3)
	(_nw27).ArraySet1(_151_v0, 4)
	(_nw27).ArraySet1(_151_v0, 5)
	(_nw27).ArraySet1((_157_v6).Cardinality(), 6)
	(_nw27).ArraySet1((_158_v7).Dtor_cf0(), 7)
	(_nw27).ArraySet1(_dafny.IntOfUint32((_152_v1).Cardinality()), 8)
	(_nw27).ArraySet1(_dafny.IntOfUint32((_159_v8).Cardinality()), 9)
	(_nw27).ArraySet1(_151_v0, 10)
	(_nw27).ArraySet1(_151_v0, 11)
	(_nw27).ArraySet1(_151_v0, 12)
	(_nw27).ArraySet1(_151_v0, 13)
	(_nw27).ArraySet1(_151_v0, 14)
	(_nw27).ArraySet1((_160_v9).Cardinality(), 15)
	(_nw27).ArraySet1(_151_v0, 16)
	_161_v10 = _nw27
	var _162_globalState *GlobalState
	_ = _162_globalState
	var _nw28 *GlobalState = New_GlobalState_()
	_ = _nw28
	_nw28.Ctor__(true, _dafny.IntOfInt64(463), true, false, false, _155_v4, _161_v10, _dafny.CodePoint('s'), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(798))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_163_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	})))
	_162_globalState = _nw28
	var _164_v11 _dafny.CodePoint
	_ = _164_v11
	_164_v11 = _dafny.CodePoint('p')
	var _pat_let_tv6 = _159_v8
	_ = _pat_let_tv6
	var _pat_let_tv7 = _159_v8
	_ = _pat_let_tv7
	var _pat_let_tv8 = _159_v8
	_ = _pat_let_tv8
	var _pat_let_tv9 = _159_v8
	_ = _pat_let_tv9
	var _pat_let_tv10 = _159_v8
	_ = _pat_let_tv10
	var _pat_let_tv11 = _159_v8
	_ = _pat_let_tv11
	var _pat_let_tv12 = _159_v8
	_ = _pat_let_tv12
	var _pat_let_tv13 = _159_v8
	_ = _pat_let_tv13
	var _pat_let_tv14 = _159_v8
	_ = _pat_let_tv14
	var _pat_let_tv15 = _159_v8
	_ = _pat_let_tv15
	var _pat_let_tv16 = _159_v8
	_ = _pat_let_tv16
	var _pat_let_tv17 = _159_v8
	_ = _pat_let_tv17
	_159_v8 = _dafny.Companion_Sequence_.Update(func(_source1 D0) _dafny.Sequence {
		if _source1.Is_DC1() {
			var _165___mcc_h0 bool = _source1.Get_().(D0_DC1).Cf1
			_ = _165___mcc_h0
			var _166_cf1 bool = _165___mcc_h0
			_ = _166_cf1
			return _pat_let_tv6
		} else if _source1.Is_DC2() {
			var _167___mcc_h1 bool = _source1.Get_().(D0_DC2).Cf2
			_ = _167___mcc_h1
			var _168___mcc_h2 _dafny.Int = _source1.Get_().(D0_DC2).Cf3
			_ = _168___mcc_h2
			var _169___mcc_h3 _dafny.MultiSet = _source1.Get_().(D0_DC2).Cf4
			_ = _169___mcc_h3
			var _170_cf4 _dafny.MultiSet = _169___mcc_h3
			_ = _170_cf4
			var _171_cf3 _dafny.Int = _168___mcc_h2
			_ = _171_cf3
			var _172_cf2 bool = _167___mcc_h1
			_ = _172_cf2
			return _dafny.Companion_Sequence_.Concatenate(_pat_let_tv7, _pat_let_tv8)
		} else if _source1.Is_DC3() {
			return _dafny.Companion_Sequence_.Concatenate(_pat_let_tv9, _pat_let_tv10)
		} else {
			var _173___mcc_h4 _dafny.Int = _source1.Get_().(D0_DC0).Cf0
			_ = _173___mcc_h4
			var _174_cf0 _dafny.Int = _173___mcc_h4
			_ = _174_cf0
			return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qbtq"), _pat_let_tv11)
		}
	}(Companion_Default___.Fm0(_153_v2, _162_globalState)), (Companion_Default___.SafeIndex(_151_v0, _dafny.IntOfUint32((func(_source2 D0) _dafny.Sequence {
		if _source2.Is_DC1() {
			var _175___mcc_h5 bool = _source2.Get_().(D0_DC1).Cf1
			_ = _175___mcc_h5
			var _176_cf1 bool = _175___mcc_h5
			_ = _176_cf1
			return _pat_let_tv12
		} else if _source2.Is_DC2() {
			var _177___mcc_h6 bool = _source2.Get_().(D0_DC2).Cf2
			_ = _177___mcc_h6
			var _178___mcc_h7 _dafny.Int = _source2.Get_().(D0_DC2).Cf3
			_ = _178___mcc_h7
			var _179___mcc_h8 _dafny.MultiSet = _source2.Get_().(D0_DC2).Cf4
			_ = _179___mcc_h8
			var _180_cf4 _dafny.MultiSet = _179___mcc_h8
			_ = _180_cf4
			var _181_cf3 _dafny.Int = _178___mcc_h7
			_ = _181_cf3
			var _182_cf2 bool = _177___mcc_h6
			_ = _182_cf2
			return _dafny.Companion_Sequence_.Concatenate(_pat_let_tv13, _pat_let_tv14)
		} else if _source2.Is_DC3() {
			return _dafny.Companion_Sequence_.Concatenate(_pat_let_tv15, _pat_let_tv16)
		} else {
			var _183___mcc_h9 _dafny.Int = _source2.Get_().(D0_DC0).Cf0
			_ = _183___mcc_h9
			var _184_cf0 _dafny.Int = _183___mcc_h9
			_ = _184_cf0
			return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qbtq"), _pat_let_tv17)
		}
	}(Companion_Default___.Fm0(_153_v2, _162_globalState))).Cardinality()))).Uint32(), _164_v11)
	var _hi5 _dafny.Int = _151_v0
	_ = _hi5
	for _185_i1 := Companion_Default___.SafeModuloInt(_151_v0, _151_v0); _185_i1.Cmp(_hi5) < 0; _185_i1 = _185_i1.Plus(_dafny.One) {
		var _source3 D0 = Companion_Default___.Fm1(_159_v8, true, _162_globalState)
		_ = _source3
		if _source3.Is_DC1() {
			var _186___mcc_h10 bool = _source3.Get_().(D0_DC1).Cf1
			_ = _186___mcc_h10
			var _187_cf1 bool = _186___mcc_h10
			_ = _187_cf1
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_156_v5), 0))
			_ = _index32
			(_156_v5).ArraySet1(_187_cf1, (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_156_v5), 0))
			_ = _index33
			(_156_v5).ArraySet1(Companion_Default___.Fm2(_162_globalState), (_index33).Int())
			var _188_v12 _dafny.Map
			_ = _188_v12
			_188_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v0, _159_v8)
			var _189_v13 D0
			_ = _189_v13
			_189_v13 = Companion_D0_.Create_DC1_(!(_188_v12).Equals(_188_v12))
			_189_v13 = _189_v13
			var _190_v14 _dafny.Map
			_ = _190_v14
			_190_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(413), _151_v0)
			var _191_v15 _dafny.Map
			_ = _191_v15
			_191_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_185_i1, _190_v14)
			var _192_v16 _dafny.Set
			_ = _192_v16
			_192_v16 = _dafny.SetOf(_dafny.IntOfUint32((_159_v8).Cardinality()))
			_151_v0 = Companion_Default___.SafeDivisionInt((_152_v1).Select((Companion_Default___.SafeIndex((((func() _dafny.Map {
				if (_191_v15).Contains(_151_v0) {
					return (_191_v15).Get(_151_v0).(_dafny.Map)
				}
				return _190_v14
			})()).Update((_192_v16).Cardinality(), _185_i1)).Cardinality(), _dafny.IntOfUint32((_152_v1).Cardinality()))).Uint32()).(_dafny.Int), _151_v0)
			var _193_v17 _dafny.MultiSet
			_ = _193_v17
			_193_v17 = _dafny.MultiSetOf((_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool), (_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool))
			var _194_v18 _dafny.Sequence
			_ = _194_v18
			_194_v18 = _dafny.SeqOf((_193_v17).IsSubsetOf(_193_v17))
			_151_v0 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_194_v18).Cardinality()))
		} else if _source3.Is_DC2() {
			var _195___mcc_h11 bool = _source3.Get_().(D0_DC2).Cf2
			_ = _195___mcc_h11
			var _196___mcc_h12 _dafny.Int = _source3.Get_().(D0_DC2).Cf3
			_ = _196___mcc_h12
			var _197___mcc_h13 _dafny.MultiSet = _source3.Get_().(D0_DC2).Cf4
			_ = _197___mcc_h13
			var _198_cf4 _dafny.MultiSet = _197___mcc_h13
			_ = _198_cf4
			var _199_cf3 _dafny.Int = _196___mcc_h12
			_ = _199_cf3
			var _200_cf2 bool = _195___mcc_h11
			_ = _200_cf2
			var _201_v19 _dafny.Map
			_ = _201_v19
			_201_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_200_cf2, _200_cf2)).Cardinality(), _200_cf2)
			var _202_v20 _dafny.Int
			_ = _202_v20
			var _out3 _dafny.Int
			_ = _out3
			_out3 = Companion_Default___.M0(_199_cf3, _161_v10, _153_v2, (_201_v19).Merge(_201_v19), _162_globalState)
			_202_v20 = _out3
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw29
			_156_v5 = _nw29
			(_162_globalState).F7 = _164_v11
			(_162_globalState).F0 = (_dafny.IntOfInt64(-274)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v2, (_152_v1).Select((Companion_Default___.SafeIndex(_202_v20, _dafny.IntOfUint32((_152_v1).Cardinality()))).Uint32()).(_dafny.Int))).Cardinality()) == 0
		} else if _source3.Is_DC3() {
			var _203_v21 _dafny.Map
			_ = _203_v21
			_203_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(363), _153_v2)
			var _204_v22 _dafny.Sequence
			_ = _204_v22
			_204_v22 = _dafny.SeqOf(_153_v2)
			var _205_v23 _dafny.Set
			_ = _205_v23
			_205_v23 = _dafny.SetOf(_151_v0, _185_i1, _dafny.IntOfUint32((_204_v22).Cardinality()))
			var _206_v24 _dafny.Int
			_ = _206_v24
			var _out4 _dafny.Int
			_ = _out4
			_out4 = Companion_Default___.M0((_185_i1).Times(_151_v0), _161_v10, _153_v2, (_203_v21).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(_164_v11, _153_v2, (_205_v23).Cardinality(), _162_globalState), _153_v2)), _162_globalState)
			_206_v24 = _out4
			_153_v2 = (_204_v22).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(_206_v24)), _dafny.IntOfUint32((_204_v22).Cardinality()))).Uint32()).(bool)
			(_162_globalState).F0 = _153_v2
			(_162_globalState).F3 = Companion_Default___.Fm2(_162_globalState)
		} else {
			var _207___mcc_h14 _dafny.Int = _source3.Get_().(D0_DC0).Cf0
			_ = _207___mcc_h14
			var _208_cf0 _dafny.Int = _207___mcc_h14
			_ = _208_cf0
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_161_v10), 0))
			_ = _index34
			(_161_v10).ArraySet1(_151_v0, (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_161_v10), 0))
			_ = _index35
			(_161_v10).ArraySet1(_dafny.IntOfUint32((_159_v8).Cardinality()), (_index35).Int())
			var _209_v25 *C3
			_ = _209_v25
			var _nw30 *C3 = New_C3_()
			_ = _nw30
			_nw30.Ctor__(_153_v2)
			_209_v25 = _nw30
			var _210_v26 _dafny.MultiSet
			_ = _210_v26
			_210_v26 = _dafny.MultiSetOf(_151_v0)
			var _211_v27 _dafny.Sequence
			_ = _211_v27
			_211_v27 = _dafny.SeqOf(_153_v2)
			var _212_v28 _dafny.Set
			_ = _212_v28
			_212_v28 = _dafny.SetOf((_211_v27).Select((Companion_Default___.SafeIndex(_151_v0, _dafny.IntOfUint32((_211_v27).Cardinality()))).Uint32()).(bool), _153_v2, _153_v2, _153_v2)
			var _213_v29 _dafny.Array
			_ = _213_v29
			var _nwElement0_4 _dafny.MultiSet = (Companion_Default___.Fm17(_152_v1, _153_v2, _162_globalState)).Difference(_210_v26)
			_ = _nwElement0_4
			var _nw31 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(23))
			_ = _nw31
			(_nw31).ArraySet1(_nwElement0_4, 0)
			(_nw31).ArraySet1(_210_v26, 1)
			(_nw31).ArraySet1(_210_v26, 2)
			(_nw31).ArraySet1(_210_v26, 3)
			(_nw31).ArraySet1(_210_v26, 4)
			(_nw31).ArraySet1((_210_v26).Intersection(_210_v26), 5)
			(_nw31).ArraySet1(_210_v26, 6)
			(_nw31).ArraySet1(_210_v26, 7)
			(_nw31).ArraySet1(_210_v26, 8)
			(_nw31).ArraySet1((_210_v26).Difference(_210_v26), 9)
			(_nw31).ArraySet1(_210_v26, 10)
			(_nw31).ArraySet1(_dafny.MultiSetOf(_151_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xpfylyv")).Cardinality())), 11)
			(_nw31).ArraySet1(_210_v26, 12)
			(_nw31).ArraySet1(_210_v26, 13)
			(_nw31).ArraySet1((_210_v26).Union(_210_v26), 14)
			(_nw31).ArraySet1(_dafny.MultiSetFromSeq(Companion_Default___.Fm13(_153_v2, _153_v2, _153_v2, (_212_v28).Cardinality(), _162_globalState)), 15)
			(_nw31).ArraySet1((_210_v26).Intersection(_dafny.MultiSetOf(_208_cf0)), 16)
			(_nw31).ArraySet1(_210_v26, 17)
			(_nw31).ArraySet1(Companion_Default___.Fm17(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(424))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_214_cf0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_215_i2 _dafny.Int) _dafny.Int {
					return _214_cf0
				}
			})(_208_cf0))), _153_v2, _162_globalState), 18)
			(_nw31).ArraySet1(_210_v26, 19)
			(_nw31).ArraySet1(_210_v26, 20)
			(_nw31).ArraySet1(_210_v26, 21)
			(_nw31).ArraySet1(_210_v26, 22)
			_213_v29 = _nw31
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_213_v29), 0))
			_ = _index36
			(_213_v29).ArraySet1((_210_v26).Intersection(_210_v26), (_index36).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_161_v10), 0))
			_ = _index37
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_213_v29), 0))
			_ = _index38
			var _rhs26 bool = _153_v2
			_ = _rhs26
			var _rhs27 _dafny.Int = _185_i1
			_ = _rhs27
			var _rhs28 _dafny.MultiSet = _210_v26
			_ = _rhs28
			var _rhs29 bool = Companion_Default___.Fm2(_162_globalState)
			_ = _rhs29
			var _lhs17 *GlobalState = _162_globalState
			_ = _lhs17
			var _lhs18 _dafny.Array = _161_v10
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_161_v10), 0))
			_ = _lhs19
			var _lhs20 _dafny.Array = _213_v29
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_213_v29), 0))
			_ = _lhs21
			var _lhs22 *GlobalState = _162_globalState
			_ = _lhs22
			_lhs17.F3 = _rhs26
			(_lhs18).ArraySet1(_rhs27, (_lhs19).Int())
			(_lhs20).ArraySet1(_rhs28, (_lhs21).Int())
			_lhs22.F0 = _rhs29
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_161_v10), 0))
			_ = _index39
			(_161_v10).ArraySet1(_208_cf0, (_index39).Int())
		}
		var _216_v30 _dafny.Map
		_ = _216_v30
		_216_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_153_v2)), _156_v5)
		_216_v30 = (_216_v30).Update(_153_v2, _156_v5)
		var _217_v31 _dafny.Set
		_ = _217_v31
		_217_v31 = _dafny.SetOf(_153_v2, _153_v2, false)
		_217_v31 = _dafny.SetOf(_153_v2, _153_v2)
		var _218_v32 *C3
		_ = _218_v32
		var _nw32 *C3 = New_C3_()
		_ = _nw32
		_nw32.Ctor__((_185_i1).Cmp(_151_v0) >= 0)
		_218_v32 = _nw32
	}
	(_162_globalState).F0 = _153_v2
	var _219_v33 _dafny.Set
	_ = _219_v33
	_219_v33 = _dafny.SetOf(true, false, _153_v2)
	var _hi6 _dafny.Int = _151_v0
	_ = _hi6
	for _220_i3 := (_219_v33).Cardinality(); _220_i3.Cmp(_hi6) < 0; _220_i3 = _220_i3.Plus(_dafny.One) {
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_156_v5), 0))
		_ = _index40
		(_156_v5).ArraySet1(Companion_Default___.Fm2(_162_globalState), (_index40).Int())
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_156_v5), 0))
		_ = _index41
		(_156_v5).ArraySet1(_153_v2, (_index41).Int())
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_161_v10), 0))
		_ = _index42
		(_161_v10).ArraySet1((_dafny.Zero).Minus(_151_v0), (_index42).Int())
		var _221_v34 _dafny.Set
		_ = _221_v34
		_221_v34 = _dafny.SetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_220_i3)), _151_v0, _220_i3, (_dafny.Zero).Minus(_220_i3))
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_161_v10), 0))
		_ = _index43
		(_161_v10).ArraySet1(Companion_Default___.Fm3(_164_v11, _153_v2, (_221_v34).Cardinality(), _162_globalState), (_index43).Int())
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_161_v10), 0))
		_ = _index44
		(_161_v10).ArraySet1(_220_i3, (_index44).Int())
		if (_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool) {
			(_162_globalState).F3 = (func() bool {
				if _153_v2 {
					return (_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool)
				}
				return _153_v2
			})()
			_151_v0 = _151_v0
			var _222_v35 *C0
			_ = _222_v35
			var _nw33 *C0 = New_C0_()
			_ = _nw33
			_nw33.Ctor__((_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool), (_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), _153_v2)
			_222_v35 = _nw33
			var _223_v36 D3
			_ = _223_v36
			_223_v36 = Companion_D3_.Create_DC11_(_151_v0, _dafny.UnicodeSeqOfUtf8Bytes("f"), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v2, _220_i3)).Cardinality(), _222_v35)
			var _224_v37 *C1
			_ = _224_v37
			var _nw34 *C1 = New_C1_()
			_ = _nw34
			_nw34.Ctor__(false, (_223_v36).Dtor_cf14(), (_222_v35).F13())
			_224_v37 = _nw34
			var _225_v38 _dafny.Map
			_ = _225_v38
			_225_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v37, _dafny.IntOfUint32((_159_v8).Cardinality()))
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_161_v10), 0))
			_ = _index45
			(_161_v10).ArraySet1((func() _dafny.Int {
				if _153_v2 {
					return (func() _dafny.Int {
						if (_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool) {
							return (_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)
						}
						return (_225_v38).Cardinality()
					})()
				}
				return _222_v35.F14
			})(), (_index45).Int())
			(_222_v35).F14 = _222_v35.F14
			var _226_v39 T0
			_ = _226_v39
			var _nw35 *C1 = New_C1_()
			_ = _nw35
			_nw35.Ctor__(_153_v2, (_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), !(!(_153_v2)))
			_226_v39 = _nw35
			var _227_v40 D4
			_ = _227_v40
			_227_v40 = Companion_D4_.Create_DC15_(_222_v35.F14, _153_v2)
			var _nw36 *C2 = New_C2_()
			_ = _nw36
			_nw36.Ctor__((_227_v40).Dtor_cf24(), (_226_v39).F9())
			_226_v39 = _nw36
		} else {
			_151_v0 = _151_v0
			var _228_v41 _dafny.Array
			_ = _228_v41
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_2
			var _nw37 _dafny.Array
			_ = _nw37
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw37 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Int = (func(_229_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_230_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_230_i4, _229_v0)
					}
				})(_151_v0)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw37 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw37).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw37).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_228_v41 = _nw37
			var _231_v42 D0
			_ = _231_v42
			_231_v42 = Companion_D0_.Create_DC1_((_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool))
			var _232_v43 _dafny.Map
			_ = _232_v43
			_232_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v0, (_231_v42).Dtor_cf1())
			var _233_v44 _dafny.Int
			_ = _233_v44
			var _out5 _dafny.Int
			_ = _out5
			_out5 = Companion_Default___.M0((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), _228_v41, !((_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool)), _232_v43, _162_globalState)
			_233_v44 = _out5
			var _234_v45 _dafny.Map
			_ = _234_v45
			_234_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm10(_162_globalState), _158_v7)
			var _235_v46 _dafny.Sequence
			_ = _235_v46
			_235_v46 = _dafny.SeqOf((_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool), _153_v2)
			_234_v45 = ((_234_v45).Update(_235_v46, Companion_D0_.Create_DC0_(_220_i3))).Update(Companion_Default___.Fm10(_162_globalState), _158_v7)
			_151_v0 = _dafny.IntOfInt64(-74)
			(_162_globalState).F7 = _164_v11
		}
	}
	var _236_v47 _dafny.Sequence
	_ = _236_v47
	_236_v47 = _dafny.SeqOf(_153_v2)
	(_162_globalState).F3 = _dafny.Companion_Sequence_.IsPrefixOf(_236_v47, _236_v47)
	var _237_v48 *C0
	_ = _237_v48
	var _nw38 *C0 = New_C0_()
	_ = _nw38
	_nw38.Ctor__(_153_v2, _151_v0, _153_v2)
	_237_v48 = _nw38
	if (func() bool {
		if _153_v2 {
			return _153_v2
		}
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_156_v5, _dafny.SetOf(_237_v48, _237_v48))).Cardinality()).Cmp(_151_v0) < 0
	})() {
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_161_v10), 0))
		_ = _index46
		(_161_v10).ArraySet1(_237_v48.F14, (_index46).Int())
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_161_v10), 0))
		_ = _index47
		(_161_v10).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_237_v48.F14), (_237_v48.F14).Plus(_dafny.IntOfInt64(-739))), (_index47).Int())
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_156_v5), 0))
		_ = _index48
		(_156_v5).ArraySet1((_237_v48).F13(), (_index48).Int())
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_156_v5), 0))
		_ = _index49
		(_156_v5).ArraySet1(true, (_index49).Int())
		var _238_v49 T0
		_ = _238_v49
		var _nw39 *C3 = New_C3_()
		_ = _nw39
		_nw39.Ctor__(((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)).Cmp((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)) < 0)
		_238_v49 = _nw39
		_238_v49 = _238_v49
		if (_dafny.IntOfInt64(26)).Cmp(Companion_Default___.SafeModuloInt(_237_v48.F14, _237_v48.F14)) != 0 {
			var _239_v50 _dafny.Array
			_ = _239_v50
			var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw40
			_239_v50 = _nw40
			var _240_v51 _dafny.Sequence
			_ = _240_v51
			_240_v51 = _dafny.SeqOf(_239_v50, _239_v50, _239_v50)
			var _241_v52 _dafny.Sequence
			_ = _241_v52
			_241_v52 = _dafny.SeqOf((_240_v51).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_159_v8).Cardinality()), _dafny.IntOfUint32((_240_v51).Cardinality()))).Uint32()).(_dafny.Array), _239_v50)
			var _242_v53 _dafny.Array
			_ = _242_v53
			var _nwElement0_5 _dafny.Array = _239_v50
			_ = _nwElement0_5
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(8))
			_ = _nw41
			(_nw41).ArraySet1(_nwElement0_5, 0)
			(_nw41).ArraySet1(_239_v50, 1)
			(_nw41).ArraySet1((func() _dafny.Array {
				if _153_v2 {
					return _239_v50
				}
				return _239_v50
			})(), 2)
			(_nw41).ArraySet1((_241_v52).Select((Companion_Default___.SafeIndex((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_241_v52).Cardinality()))).Uint32()).(_dafny.Array), 3)
			(_nw41).ArraySet1(_239_v50, 4)
			(_nw41).ArraySet1(_239_v50, 5)
			(_nw41).ArraySet1(_239_v50, 6)
			(_nw41).ArraySet1(_239_v50, 7)
			_242_v53 = _nw41
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_242_v53), 0))
			_ = _index50
			(_242_v53).ArraySet1(_239_v50, (_index50).Int())
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_242_v53), 0))
			_ = _index51
			var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw42
			(_242_v53).ArraySet1(_nw42, (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_161_v10), 0))
			_ = _index52
			(_161_v10).ArraySet1(_237_v48.F14, (_index52).Int())
			var _243_v54 _dafny.Map
			_ = _243_v54
			_243_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool), _239_v50)
			_243_v54 = (func() _dafny.Map {
				if !(Companion_Default___.Fm2(_162_globalState)) || (_153_v2) {
					return _243_v54
				}
				return (_243_v54).Update(true, _239_v50)
			})()
			var _244_v55 _dafny.MultiSet
			_ = _244_v55
			_244_v55 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), _151_v0))
			var _245_v56 _dafny.Map
			_ = _245_v56
			_245_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_237_v48.F14, (_160_v9).Cardinality())
			var _246_v57 _dafny.Sequence
			_ = _246_v57
			_246_v57 = _dafny.SeqOf(_236_v47)
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_161_v10), 0))
			_ = _index53
			(_161_v10).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if (_244_v55).Contains((func() _dafny.Int {
					if (_245_v56).Contains((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)) {
						return (_245_v56).Get((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)).(_dafny.Int)
					}
					return _237_v48.F14
				})()) {
					return (_244_v55).Multiplicity((func() _dafny.Int {
						if (_245_v56).Contains((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)) {
							return (_245_v56).Get((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int)).(_dafny.Int)
						}
						return _237_v48.F14
					})())
				}
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_236_v47, (_246_v57).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_151_v0), _dafny.IntOfUint32((_246_v57).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality())
			})())), (_index53).Int())
			var _247_v58 _dafny.Map
			_ = _247_v58
			_247_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_237_v48).F13(), (_237_v48).F13())
			var _248_v59 *C2
			_ = _248_v59
			var _nw43 *C2 = New_C2_()
			_ = _nw43
			_nw43.Ctor__(!(_247_v58).Equals(_247_v58), (_237_v48).F13())
			_248_v59 = _nw43
			var _249_v60 _dafny.Map
			_ = _249_v60
			_249_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_162_globalState), _248_v59)
			_248_v59 = (func() *C2 {
				if _153_v2 {
					return _248_v59
				}
				return (func() *C2 {
					if (_249_v60).Contains((_237_v48).F13()) {
						return (_249_v60).Get((_237_v48).F13()).(*C2)
					}
					return _248_v59
				})()
			})()
		} else {
			var _250_v61 _dafny.Map
			_ = _250_v61
			_250_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_237_v48).F13())
			var _251_v62 _dafny.Map
			_ = _251_v62
			_251_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_250_v61).Cardinality(), (_237_v48).F13())
			var _252_v63 _dafny.Map
			_ = _252_v63
			_252_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v9, !(_251_v62).Equals(_251_v62))
			_252_v63 = _252_v63
			_237_v48 = _237_v48
			var _253_v64 _dafny.Array
			_ = _253_v64
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
			_ = _nw44
			_253_v64 = _nw44
			var _254_v65 *C1
			_ = _254_v65
			var _nw45 *C1 = New_C1_()
			_ = _nw45
			_nw45.Ctor__(true, (_dafny.Zero).Minus(_237_v48.F14), (_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool))
			_254_v65 = _nw45
			var _255_v66 _dafny.Map
			_ = _255_v66
			_255_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v64, _254_v65)
			_255_v66 = (_255_v66).Update(_253_v64, _254_v65)
			(_237_v48).F14 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(778), _237_v48.F14)
			var _256_v67 _dafny.MultiSet
			_ = _256_v67
			_256_v67 = _dafny.MultiSetOf((_254_v65).F12())
			var _257_v68 *C0
			_ = _257_v68
			var _nw46 *C0 = New_C0_()
			_ = _nw46
			_nw46.Ctor__((_256_v67).Contains(_dafny.IntOfInt64(550)), _dafny.IntOfInt64(114), (_238_v49).F9())
			_257_v68 = _nw46
		}
		var _258_v69 *C2
		_ = _258_v69
		var _nw47 *C2 = New_C2_()
		_ = _nw47
		_nw47.Ctor__(_dafny.Companion_Sequence_.Contains(_159_v8, _164_v11), (_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool))
		_258_v69 = _nw47
		_258_v69 = _258_v69
	} else {
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_156_v5), 0))
		_ = _index54
		(_156_v5).ArraySet1((_237_v48).F13(), (_index54).Int())
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_156_v5), 0))
		_ = _index55
		(_156_v5).ArraySet1((_237_v48).F13(), (_index55).Int())
		(_162_globalState).F3 = (_237_v48).F13()
		var _259_v70 D0
		_ = _259_v70
		_259_v70 = Companion_D0_.Create_DC3_()
		var _source4 D0 = (func() D0 {
			if (_237_v48).F13() {
				return _259_v70
			}
			return _259_v70
		})()
		_ = _source4
		if _source4.Is_DC1() {
			var _260___mcc_h15 bool = _source4.Get_().(D0_DC1).Cf1
			_ = _260___mcc_h15
			var _261_cf1 bool = _260___mcc_h15
			_ = _261_cf1
			var _262_v71 *C1
			_ = _262_v71
			var _nw48 *C1 = New_C1_()
			_ = _nw48
			_nw48.Ctor__((_237_v48).F13(), _237_v48.F14, !((_237_v48).F13()))
			_262_v71 = _nw48
			var _263_v72 _dafny.Map
			_ = _263_v72
			_263_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(186), _262_v71)
			var _264_v73 _dafny.Map
			_ = _264_v73
			_264_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_263_v72).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(366), _262_v71))).Cardinality(), _236_v47)
			var _rhs30 _dafny.Map = (_264_v73).Merge(_264_v73)
			_ = _rhs30
			var _rhs31 _dafny.Int = (_dafny.Zero).Minus(_237_v48.F14)
			_ = _rhs31
			var _lhs23 *C0 = _237_v48
			_ = _lhs23
			_264_v73 = _rhs30
			_lhs23.F14 = _rhs31
			_159_v8 = _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("p"), (Companion_Default___.SafeIndex((_262_v71).F12(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()))).Uint32(), Companion_Default___.Fm18(_162_globalState))
			var _265_v74 _dafny.Array
			_ = _265_v74
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_3
			var _nw49 _dafny.Array
			_ = _nw49
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw49 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = (func(_266_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_267_i5 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(904))).Uint32(), func(coer17 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
							return func(arg17 _dafny.Int) interface{} {
								return coer17(arg17)
							}
						}((func(_268_v8 _dafny.Sequence) func(_dafny.Int) D3 {
							return func(_269_i6 _dafny.Int) D3 {
								return Companion_D3_.Create_DC9_(_268_v8)
							}
						})(_266_v8)))
					}
				})(_159_v8)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw49 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw49).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw49).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_265_v74 = _nw49
			var _270_v75 _dafny.Map
			_ = _270_v75
			_270_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_262_v71).F12(), _dafny.IntOfInt64(-277))
			var _271_v76 D3
			_ = _271_v76
			_271_v76 = Companion_D3_.Create_DC9_(_159_v8)
			var _272_v77 _dafny.Sequence
			_ = _272_v77
			_272_v77 = _dafny.SeqOf(_271_v76, _271_v76, _271_v76)
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_265_v74), 0))
			_ = _index56
			(_265_v74).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm19(_270_v75, _162_globalState), _272_v77), (_index56).Int())
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_265_v74), 0))
			_ = _index57
			(_265_v74).ArraySet1(_272_v77, (_index57).Int())
			var _273_v79 _dafny.Int
			_ = _273_v79
			var _out6 _dafny.Int
			_ = _out6
			_out6 = Companion_Default___.M0(_237_v48.F14, _161_v10, !_dafny.Companion_Sequence_.Equal(_159_v8, _dafny.UnicodeSeqOfUtf8Bytes("gretekrl")), func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(515), _dafny.IntOfInt64(497))); ; {
					_compr_9, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _274_v78 _dafny.Int
					_274_v78 = interface{}(_compr_9).(_dafny.Int)
					if ((_dafny.IntOfInt64(515)).Cmp(_274_v78) <= 0) && ((_274_v78).Cmp(_dafny.IntOfInt64(497)) < 0) {
						_coll9.Add((_274_v78).Minus(_237_v48.F14), false)
					}
				}
				return _coll9.ToMap()
			}(), _162_globalState)
			_273_v79 = _out6
		} else if _source4.Is_DC2() {
			var _275___mcc_h16 bool = _source4.Get_().(D0_DC2).Cf2
			_ = _275___mcc_h16
			var _276___mcc_h17 _dafny.Int = _source4.Get_().(D0_DC2).Cf3
			_ = _276___mcc_h17
			var _277___mcc_h18 _dafny.MultiSet = _source4.Get_().(D0_DC2).Cf4
			_ = _277___mcc_h18
			var _278_cf4 _dafny.MultiSet = _277___mcc_h18
			_ = _278_cf4
			var _279_cf3 _dafny.Int = _276___mcc_h17
			_ = _279_cf3
			var _280_cf2 bool = _275___mcc_h16
			_ = _280_cf2
			var _281_v80 *C1
			_ = _281_v80
			var _nw50 *C1 = New_C1_()
			_ = _nw50
			_nw50.Ctor__((_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool), (_279_cf3).Times(_279_cf3), (_237_v48).F13())
			_281_v80 = _nw50
			_160_v9 = (_160_v9).Update((_281_v80).F11(), (_160_v9).Cardinality())
			var _282_v81 *C2
			_ = _282_v81
			var _nw51 *C2 = New_C2_()
			_ = _nw51
			_nw51.Ctor__((_237_v48).F13(), _dafny.Companion_Sequence_.IsPrefixOf(_152_v1, _152_v1))
			_282_v81 = _nw51
			var _nw52 *C2 = New_C2_()
			_ = _nw52
			_nw52.Ctor__((_151_v0).Cmp(_237_v48.F14) == 0, ((_281_v80).F12()).Cmp(_279_cf3) >= 0)
			_282_v81 = _nw52
			var _rhs32 _dafny.CodePoint = Companion_Default___.Fm18(_162_globalState)
			_ = _rhs32
			var _rhs33 _dafny.Int = _dafny.IntOfInt64(712)
			_ = _rhs33
			var _lhs24 *GlobalState = _162_globalState
			_ = _lhs24
			_lhs24.F7 = _rhs32
			_279_cf3 = _rhs33
		} else if _source4.Is_DC3() {
			(_237_v48).F14 = _237_v48.F14
			var _283_v82 _dafny.Map
			_ = _283_v82
			_283_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v2, !_dafny.Companion_Sequence_.Equal(_159_v8, _159_v8))
			var _284_v83 _dafny.Array
			_ = _284_v83
			var _nw53 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw53
			_284_v83 = _nw53
			var _285_v84 _dafny.Map
			_ = _285_v84
			_285_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC7_(_284_v83), _159_v8)
			var _286_v85 _dafny.Map
			_ = _286_v85
			_286_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
				if (_285_v84).Contains(Companion_D2_.Create_DC7_(_284_v83)) {
					return (_285_v84).Get(Companion_D2_.Create_DC7_(_284_v83)).(_dafny.Sequence)
				}
				return _159_v8
			})(), _236_v47)
			var _rhs34 _dafny.Map = _283_v82
			_ = _rhs34
			var _rhs35 _dafny.Map = (_286_v85).Merge(_286_v85)
			_ = _rhs35
			var _rhs36 _dafny.Int = _151_v0
			_ = _rhs36
			var _lhs25 *C0 = _237_v48
			_ = _lhs25
			_283_v82 = _rhs34
			_286_v85 = _rhs35
			_lhs25.F14 = _rhs36
			var _287_v86 _dafny.Map
			_ = _287_v86
			_287_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_159_v8).Cardinality()), _151_v0)
			var _288_v87 _dafny.Map
			_ = _288_v87
			_288_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_287_v86).Update(_dafny.IntOfUint32((_159_v8).Cardinality()), _237_v48.F14)).Cardinality(), Companion_Default___.Fm2(_162_globalState))
			var _289_v88 _dafny.Int
			_ = _289_v88
			var _out7 _dafny.Int
			_ = _out7
			_out7 = Companion_Default___.M0((_dafny.Zero).Minus(_dafny.IntOfUint32((_159_v8).Cardinality())), _161_v10, _153_v2, _288_v87, _162_globalState)
			_289_v88 = _out7
			var _290_v89 _dafny.Int
			_ = _290_v89
			var _out8 _dafny.Int
			_ = _out8
			_out8 = Companion_Default___.M0((_dafny.Zero).Minus((_151_v0).Plus(_dafny.IntOfInt64(514))), _161_v10, (_237_v48).F13(), _288_v87, _162_globalState)
			_290_v89 = _out8
		} else {
			var _291___mcc_h19 _dafny.Int = _source4.Get_().(D0_DC0).Cf0
			_ = _291___mcc_h19
			var _292_cf0 _dafny.Int = _291___mcc_h19
			_ = _292_cf0
			(_162_globalState).F0 = (_237_v48).F13()
			var _293_v90 _dafny.Array
			_ = _293_v90
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_4
			var _nw54 _dafny.Array
			_ = _nw54
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw54 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Sequence = (func(_294_v11 _dafny.CodePoint, _295_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_296_i7 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-815))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg18 _dafny.Int) interface{} {
								return coer18(arg18)
							}
						}((func(_297_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_298_i8 _dafny.Int) _dafny.CodePoint {
								return _297_v11
							}
						})(_294_v11))), _295_v8)
					}
				})(_164_v11, _159_v8)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw54 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw54).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw54).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_293_v90 = _nw54
			_293_v90 = _293_v90
			var _299_v91 *C0
			_ = _299_v91
			var _nw55 *C0 = New_C0_()
			_ = _nw55
			_nw55.Ctor__(!((_237_v48).F13()), _237_v48.F14, _153_v2)
			_299_v91 = _nw55
			var _300_v92 *C1
			_ = _300_v92
			var _nw56 *C1 = New_C1_()
			_ = _nw56
			_nw56.Ctor__(true, _237_v48.F14, (_237_v48).F13())
			_300_v92 = _nw56
			var _301_v93 _dafny.Sequence
			_ = _301_v93
			_301_v93 = _dafny.SeqOf(_300_v92)
			_153_v2 = ((func() _dafny.Int {
				if (_237_v48).F13() {
					return _dafny.IntOfInt64(428)
				}
				return _dafny.IntOfUint32((_301_v93).Cardinality())
			})()).Cmp(_151_v0) < 0
		}
		var _302_v94 _dafny.Array
		_ = _302_v94
		var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(15))
		_ = _nw57
		_302_v94 = _nw57
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_302_v94), 0))
		_ = _index58
		(_302_v94).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_236_v47, (Companion_Default___.SafeIndex(_151_v0, _dafny.IntOfUint32((_236_v47).Cardinality()))).Uint32(), (_237_v48).F13()), _236_v47)), (_index58).Int())
		var _303_v95 _dafny.MultiSet
		_ = _303_v95
		_303_v95 = _dafny.MultiSetOf(true)
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_302_v94), 0))
		_ = _index59
		(_302_v94).ArraySet1(_303_v95, (_index59).Int())
		_151_v0 = _237_v48.F14
	}
	var _304_v96 _dafny.MultiSet
	_ = _304_v96
	_304_v96 = _dafny.MultiSetOf(_151_v0)
	_151_v0 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_237_v48).F13(), (_304_v96).Cardinality())).Cardinality()).Plus(_dafny.IntOfInt64(771))
	if (func() bool {
		if _153_v2 {
			return (Companion_Default___.Fm2(_162_globalState)) == ((_237_v48).F13())
		}
		return _153_v2
	})() {
		(_237_v48).F14 = _237_v48.F14
		var _305_v97 D0
		_ = _305_v97
		_305_v97 = Companion_D0_.Create_DC3_()
		var _306_v98 _dafny.Set
		_ = _306_v98
		_306_v98 = _dafny.SetOf(_305_v97, _305_v97)
		(_162_globalState).F3 = ((_304_v96).Cardinality()).Cmp((_306_v98).Cardinality()) == 0
		var _307_v99 *C2
		_ = _307_v99
		var _nw58 *C2 = New_C2_()
		_ = _nw58
		_nw58.Ctor__((_dafny.SetOf(_153_v2)).IsProperSubsetOf(_219_v33), (_237_v48).F13())
		_307_v99 = _nw58
		_307_v99 = _307_v99
		var _308_v100 *C2
		_ = _308_v100
		var _nw59 *C2 = New_C2_()
		_ = _nw59
		_nw59.Ctor__(false, _153_v2)
		_308_v100 = _nw59
		_237_v48 = _237_v48
	} else {
		(_162_globalState).F0 = (_237_v48).F13()
		(_162_globalState).F7 = _164_v11
		_152_v1 = _dafny.SeqOf((Companion_Default___.Fm3(_164_v11, (_237_v48).F13(), Companion_Default___.Fm3(_164_v11, (_237_v48).F13(), _237_v48.F14, _162_globalState), _162_globalState)).Minus(_237_v48.F14), _dafny.IntOfInt64(978), _151_v0)
		if (_237_v48).F13() {
			var _309_v101 *C3
			_ = _309_v101
			var _nw60 *C3 = New_C3_()
			_ = _nw60
			_nw60.Ctor__(false)
			_309_v101 = _nw60
			var _310_v102 _dafny.Sequence
			_ = _310_v102
			_310_v102 = _dafny.SeqOf(_161_v10, _161_v10, _161_v10)
			(_237_v48).F14 = (_dafny.IntOfUint32((_310_v102).Cardinality())).Plus((_151_v0).Plus(Companion_Default___.Fm3(_164_v11, (_237_v48).F13(), _237_v48.F14, _162_globalState)))
			var _311_v103 _dafny.Map
			_ = _311_v103
			_311_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_237_v48.F14, (_237_v48).F13())
			var _312_v104 _dafny.Int
			_ = _312_v104
			var _out9 _dafny.Int
			_ = _out9
			_out9 = Companion_Default___.M0(_237_v48.F14, _161_v10, (_237_v48).F13(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_237_v48.F14, (_237_v48).F13())).Merge(_311_v103), _162_globalState)
			_312_v104 = _out9
			(_237_v48).F14 = _237_v48.F14
			var _313_v105 *C3
			_ = _313_v105
			var _nw61 *C3 = New_C3_()
			_ = _nw61
			_nw61.Ctor__((_237_v48).F13())
			_313_v105 = _nw61
		} else {
			var _314_v106 *C3
			_ = _314_v106
			var _nw62 *C3 = New_C3_()
			_ = _nw62
			_nw62.Ctor__(_153_v2)
			_314_v106 = _nw62
			var _315_v107 _dafny.Sequence
			_ = _315_v107
			_315_v107 = _dafny.SeqOf(_237_v48)
			var _316_v108 D3
			_ = _316_v108
			_316_v108 = Companion_D3_.Create_DC11_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_314_v106, _314_v106), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_151_v0), _dafny.IntOfUint32((_dafny.SeqOf(_314_v106, _314_v106)).Cardinality()))).Uint32(), _314_v106)).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(310))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_317_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_318_i9 _dafny.Int) _dafny.CodePoint {
					return _317_v11
				}
			})(_164_v11))), _237_v48.F14, (_315_v107).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.IntOfUint32((_315_v107).Cardinality()))).Uint32()).(*C0))
			_316_v108 = _316_v108
			_314_v106 = (Companion_D5_.Create_DC17_(_314_v106)).Dtor_cf26()
			var _319_v109 T0
			_ = _319_v109
			var _nw63 *C3 = New_C3_()
			_ = _nw63
			_nw63.Ctor__((_237_v48).F13())
			_319_v109 = _nw63
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_156_v5), 0))
			_ = _index60
			(_156_v5).ArraySet1((_237_v48).F13(), (_index60).Int())
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_156_v5), 0))
			_ = _index61
			(_156_v5).ArraySet1((_237_v48).F13(), (_index61).Int())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_156_v5), 0))
			_ = _index62
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_156_v5), 0))
			_ = _index63
			var _rhs37 T0 = _319_v109
			_ = _rhs37
			var _rhs38 bool = !((func() bool {
				if _153_v2 {
					return (_237_v48).F13()
				}
				return (_237_v48).F13()
			})()) || ((_237_v48).F13())
			_ = _rhs38
			var _rhs39 bool = (_237_v48).F13()
			_ = _rhs39
			var _rhs40 bool = (_237_v48).F13()
			_ = _rhs40
			var _lhs26 _dafny.Array = _156_v5
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_156_v5), 0))
			_ = _lhs27
			var _lhs28 _dafny.Array = _156_v5
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_156_v5), 0))
			_ = _lhs29
			_319_v109 = _rhs37
			(_lhs26).ArraySet1(_rhs38, (_lhs27).Int())
			_153_v2 = _rhs39
			(_lhs28).ArraySet1(_rhs40, (_lhs29).Int())
			(_237_v48).F14 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_152_v1, _152_v1)).Cardinality())).Times(Companion_Default___.SafeDivisionInt((_152_v1).Select((Companion_Default___.SafeIndex(_237_v48.F14, _dafny.IntOfUint32((_152_v1).Cardinality()))).Uint32()).(_dafny.Int), _237_v48.F14))
			var _320_v110 D4
			_ = _320_v110
			_320_v110 = Companion_D4_.Create_DC15_(_dafny.IntOfInt64(718), true)
			var _321_v111 _dafny.Map
			_ = _321_v111
			_321_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_320_v110).Dtor_cf24() {
					return _151_v0
				}
				return (_dafny.SetOf(_dafny.SeqOf((_237_v48).F13(), (_319_v109).F9()))).Cardinality()
			})(), (_237_v48).F13())
			var _rhs41 _dafny.Array = _161_v10
			_ = _rhs41
			var _rhs42 _dafny.Map = _321_v111
			_ = _rhs42
			var _rhs43 _dafny.Int = (_219_v33).Cardinality()
			_ = _rhs43
			var _rhs44 _dafny.CodePoint = _164_v11
			_ = _rhs44
			var _lhs30 *GlobalState = _162_globalState
			_ = _lhs30
			_161_v10 = _rhs41
			_321_v111 = _rhs42
			_151_v0 = _rhs43
			_lhs30.F7 = _rhs44
		}
		_156_v5 = _156_v5
	}
	var _322_v112 *C3
	_ = _322_v112
	var _nw64 *C3 = New_C3_()
	_ = _nw64
	_nw64.Ctor__(_153_v2)
	_322_v112 = _nw64
	var _323_v113 _dafny.Set
	_ = _323_v113
	_323_v113 = _dafny.SetOf((_dafny.Zero).Minus((_219_v33).Cardinality()))
	var _324_v114 D5
	_ = _324_v114
	_324_v114 = Companion_D5_.Create_DC18_((_323_v113).Cardinality(), (_237_v48).F13())
	_151_v0 = (func() _dafny.Int {
		if !((_237_v48.F14).Cmp((_324_v114).Dtor_cf27()) >= 0) {
			return _151_v0
		}
		return _151_v0
	})()
	if false {
		var _325_v115 _dafny.Map
		_ = _325_v115
		_325_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_237_v48.F14, _237_v48.F14)
		var _326_v116 _dafny.Map
		_ = _326_v116
		_326_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v0, (_237_v48).F13())
		var _327_v117 _dafny.Int
		_ = _327_v117
		var _out10 _dafny.Int
		_ = _out10
		_out10 = Companion_Default___.M0((_dafny.MultiSetOf(_325_v115, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v0, _237_v48.F14))).Cardinality(), _161_v10, !(_323_v113).Equals(_323_v113), _326_v116, _162_globalState)
		_327_v117 = _out10
		var _328_v118 D0
		_ = _328_v118
		_328_v118 = Companion_D0_.Create_DC3_()
		_328_v118 = _328_v118
		var _329_v119 _dafny.Sequence
		_ = _329_v119
		_329_v119 = _dafny.SeqOf(_161_v10, _161_v10)
		var _330_v120 _dafny.Array
		_ = _330_v120
		var _nwElement0_6 _dafny.Array = _161_v10
		_ = _nwElement0_6
		var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(22))
		_ = _nw65
		(_nw65).ArraySet1(_nwElement0_6, 0)
		(_nw65).ArraySet1(_161_v10, 1)
		(_nw65).ArraySet1(_161_v10, 2)
		(_nw65).ArraySet1(_161_v10, 3)
		(_nw65).ArraySet1(_161_v10, 4)
		(_nw65).ArraySet1(_161_v10, 5)
		(_nw65).ArraySet1(_161_v10, 6)
		(_nw65).ArraySet1(_161_v10, 7)
		(_nw65).ArraySet1((_329_v119).Select((Companion_Default___.SafeIndex(_237_v48.F14, _dafny.IntOfUint32((_329_v119).Cardinality()))).Uint32()).(_dafny.Array), 8)
		(_nw65).ArraySet1(_161_v10, 9)
		(_nw65).ArraySet1(_161_v10, 10)
		(_nw65).ArraySet1(_161_v10, 11)
		(_nw65).ArraySet1(_161_v10, 12)
		(_nw65).ArraySet1(_161_v10, 13)
		(_nw65).ArraySet1(_161_v10, 14)
		(_nw65).ArraySet1(_161_v10, 15)
		(_nw65).ArraySet1(_161_v10, 16)
		(_nw65).ArraySet1(_161_v10, 17)
		(_nw65).ArraySet1(_161_v10, 18)
		(_nw65).ArraySet1(_161_v10, 19)
		(_nw65).ArraySet1(_161_v10, 20)
		(_nw65).ArraySet1((_329_v119).Select((Companion_Default___.SafeIndex(_151_v0, _dafny.IntOfUint32((_329_v119).Cardinality()))).Uint32()).(_dafny.Array), 21)
		_330_v120 = _nw65
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_330_v120), 0))
		_ = _index64
		(_330_v120).ArraySet1((_329_v119).Select((Companion_Default___.SafeIndex(_327_v117, _dafny.IntOfUint32((_329_v119).Cardinality()))).Uint32()).(_dafny.Array), (_index64).Int())
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_330_v120), 0))
		_ = _index65
		(_330_v120).ArraySet1(_161_v10, (_index65).Int())
		var _331_v121 _dafny.MultiSet
		_ = _331_v121
		_331_v121 = _dafny.MultiSetOf(_237_v48)
		var _332_v122 *C2
		_ = _332_v122
		var _nw66 *C2 = New_C2_()
		_ = _nw66
		_nw66.Ctor__(false, (_237_v48).F13())
		_332_v122 = _nw66
		var _333_v123 _dafny.Sequence
		_ = _333_v123
		_333_v123 = _dafny.SeqOf(_332_v122)
		var _334_v124 _dafny.Map
		_ = _334_v124
		_334_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_331_v121, (_333_v123).Select((Companion_Default___.SafeIndex(_327_v117, _dafny.IntOfUint32((_333_v123).Cardinality()))).Uint32()).(*C2))
		_334_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_237_v48), _332_v122)
		var _335_v125 T0
		_ = _335_v125
		var _nw67 *C3 = New_C3_()
		_ = _nw67
		_nw67.Ctor__((_237_v48).F13())
		_335_v125 = _nw67
		var _336_v126 *C0
		_ = _336_v126
		var _nw68 *C0 = New_C0_()
		_ = _nw68
		_nw68.Ctor__(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_335_v125, (_335_v125).F9())).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_236_v47, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_237_v48.F14, _327_v117)).Cardinality(), _dafny.IntOfUint32((_236_v47).Cardinality()))).Uint32(), (_237_v48).F13())).Cardinality())) < 0, _237_v48.F14, true)
		_336_v126 = _nw68
	} else {
		var _337_v127 _dafny.MultiSet
		_ = _337_v127
		_337_v127 = _dafny.MultiSetOf(!((_237_v48).F13()), (_237_v48).F13(), !(true))
		(_162_globalState).F3 = !(!((true) == ((_237_v48).F13())) || ((_337_v127).IsDisjointFrom(Companion_Default___.Fm12(!((_237_v48).F13()), (_237_v48).F13(), _237_v48.F14, (_237_v48).F13(), _162_globalState))))
		var _338_v128 _dafny.Map
		_ = _338_v128
		_338_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_237_v48).F13(), true)
		_338_v128 = (_338_v128).Update((_237_v48).F13(), (_237_v48).F13())
		var _339_v129 D3
		_ = _339_v129
		_339_v129 = Companion_D3_.Create_DC10_(_237_v48.F14, _161_v10, false)
		_161_v10 = (_339_v129).Dtor_cf12()
		var _340_v130 *C3
		_ = _340_v130
		var _nw69 *C3 = New_C3_()
		_ = _nw69
		_nw69.Ctor__((_237_v48.F14).Cmp(_237_v48.F14) != 0)
		_340_v130 = _nw69
		var _341_v131 _dafny.Map
		_ = _341_v131
		_341_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_237_v48.F14, _dafny.Companion_Sequence_.Equal(_159_v8, _dafny.Companion_Sequence_.Update(_159_v8, (Companion_Default___.SafeIndex(_237_v48.F14, _dafny.IntOfUint32((_159_v8).Cardinality()))).Uint32(), _164_v11)))
		_341_v131 = (_341_v131).Update((_237_v48.F14).Minus(_237_v48.F14), _153_v2)
	}
	var _hi7 _dafny.Int = (_219_v33).Cardinality()
	_ = _hi7
	for _342_i10 := _151_v0; _342_i10.Cmp(_hi7) < 0; _342_i10 = _342_i10.Plus(_dafny.One) {
		var _343_v132 D4
		_ = _343_v132
		_343_v132 = Companion_D4_.Create_DC15_(_dafny.IntOfInt64(631), _153_v2)
		var _344_v133 _dafny.Map
		_ = _344_v133
		_344_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_343_v132, _156_v5)
		var _345_v134 D2
		_ = _345_v134
		_345_v134 = Companion_D2_.Create_DC7_((func() _dafny.Array {
			if (_344_v133).Contains(_343_v132) {
				return (_344_v133).Get(_343_v132).(_dafny.Array)
			}
			return _156_v5
		})())
		var _source5 D2 = _345_v134
		_ = _source5
		if _source5.Is_DC8() {
			var _346___mcc_h20 bool = _source5.Get_().(D2_DC8).Cf7
			_ = _346___mcc_h20
			var _347___mcc_h21 _dafny.Int = _source5.Get_().(D2_DC8).Cf8
			_ = _347___mcc_h21
			var _348___mcc_h22 bool = _source5.Get_().(D2_DC8).Cf9
			_ = _348___mcc_h22
			var _349_cf9 bool = _348___mcc_h22
			_ = _349_cf9
			var _350_cf8 _dafny.Int = _347___mcc_h21
			_ = _350_cf8
			var _351_cf7 bool = _346___mcc_h20
			_ = _351_cf7
			var _352_v135 _dafny.Array
			_ = _352_v135
			var _nwElement0_7 _dafny.CodePoint = _164_v11
			_ = _nwElement0_7
			var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(5))
			_ = _nw70
			(_nw70).ArraySet1CodePoint(_nwElement0_7, 0)
			(_nw70).ArraySet1CodePoint(_164_v11, 1)
			(_nw70).ArraySet1CodePoint(_164_v11, 2)
			(_nw70).ArraySet1CodePoint(_164_v11, 3)
			(_nw70).ArraySet1CodePoint(_dafny.CodePoint('n'), 4)
			_352_v135 = _nw70
			var _353_v136 _dafny.Set
			_ = _353_v136
			_353_v136 = _dafny.SetOf(_352_v135)
			var _354_v137 _dafny.Int
			_ = _354_v137
			var _355_v138 D0
			_ = _355_v138
			var _356_v139 _dafny.Int
			_ = _356_v139
			var _out11 _dafny.Int
			_ = _out11
			var _out12 D0
			_ = _out12
			var _out13 _dafny.Int
			_ = _out13
			_out11, _out12, _out13 = (_322_v112).M1(_353_v136, _162_globalState)
			_354_v137 = _out11
			_355_v138 = _out12
			_356_v139 = _out13
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_161_v10), 0))
			_ = _index66
			(_161_v10).ArraySet1(_151_v0, (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_161_v10), 0))
			_ = _index67
			(_161_v10).ArraySet1((_dafny.IntOfInt64(353)).Times((func() _dafny.Int {
				if (_155_v4).Contains(_152_v1) {
					return (_155_v4).Multiplicity(_152_v1)
				}
				return _342_i10
			})()), (_index67).Int())
			_153_v2 = !(!((_237_v48).F13()) || ((_324_v114).Dtor_cf28())) || ((_237_v48).F13())
			(_237_v48).F14 = _237_v48.F14
		} else {
			var _357___mcc_h23 _dafny.Array = _source5.Get_().(D2_DC7).Cf6
			_ = _357___mcc_h23
			var _358_cf6 _dafny.Array = _357___mcc_h23
			_ = _358_cf6
			(_162_globalState).F0 = _153_v2
			(_162_globalState).F0 = _153_v2
			var _359_v140 _dafny.Sequence
			_ = _359_v140
			_359_v140 = _dafny.SeqOf(_161_v10, _161_v10, _161_v10, _161_v10)
			var _360_v141 _dafny.Map
			_ = _360_v141
			_360_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_164_v11, _159_v8)
			var _361_v142 _dafny.Int
			_ = _361_v142
			var _out14 _dafny.Int
			_ = _out14
			_out14 = Companion_Default___.M0((_dafny.IntOfInt64(209)).Minus(_237_v48.F14), (_359_v140).Select((Companion_Default___.SafeIndex(_237_v48.F14, _dafny.IntOfUint32((_359_v140).Cardinality()))).Uint32()).(_dafny.Array), !((_237_v48).F13()), Companion_Default___.Fm16(_360_v141, _342_i10, _162_globalState), _162_globalState)
			_361_v142 = _out14
			(_237_v48).F14 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_237_v48.F14).Minus(_151_v0), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(684), (_dafny.Zero).Minus(_237_v48.F14))))
		}
		var _362_v143 _dafny.Map
		_ = _362_v143
		_362_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_i10, _237_v48.F14)
		_362_v143 = _362_v143
		var _363_v144 _dafny.Array
		_ = _363_v144
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_5
		var _nw71 _dafny.Array
		_ = _nw71
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw71 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.CodePoint = (func(_364_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_365_i11 _dafny.Int) _dafny.CodePoint {
					return _364_v11
				}
			})(_164_v11)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw71 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw71).ArraySet1CodePoint(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw71).ArraySet1CodePoint(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_363_v144 = _nw71
		var _366_v145 _dafny.Set
		_ = _366_v145
		_366_v145 = _dafny.SetOf(_363_v144)
		var _367_v146 _dafny.Int
		_ = _367_v146
		var _368_v147 D0
		_ = _368_v147
		var _369_v148 _dafny.Int
		_ = _369_v148
		var _out15 _dafny.Int
		_ = _out15
		var _out16 D0
		_ = _out16
		var _out17 _dafny.Int
		_ = _out17
		_out15, _out16, _out17 = (_322_v112).M1(_366_v145, _162_globalState)
		_367_v146 = _out15
		_368_v147 = _out16
		_369_v148 = _out17
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_156_v5), 0))
		_ = _index68
		(_156_v5).ArraySet1((_237_v48).F13(), (_index68).Int())
		var _370_v149 _dafny.Sequence
		_ = _370_v149
		_370_v149 = _dafny.SeqOf(_237_v48)
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_156_v5), 0))
		_ = _index69
		(_156_v5).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_370_v149, _370_v149), _dafny.Companion_Sequence_.Concatenate(_370_v149, _370_v149)), (_index69).Int())
	}
	var _371_v150 D5
	_ = _371_v150
	_371_v150 = Companion_D5_.Create_DC19_(_219_v33)
	var _372_v151 _dafny.Map
	_ = _372_v151
	_372_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_371_v150, _156_v5)
	var _373_v152 _dafny.Array
	_ = _373_v152
	var _nwElement0_8 _dafny.Map = _372_v151
	_ = _nwElement0_8
	var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(27))
	_ = _nw72
	(_nw72).ArraySet1(_nwElement0_8, 0)
	(_nw72).ArraySet1(_372_v151, 1)
	(_nw72).ArraySet1(_372_v151, 2)
	(_nw72).ArraySet1((_372_v151).Merge(_372_v151), 3)
	(_nw72).ArraySet1(_372_v151, 4)
	(_nw72).ArraySet1((_372_v151).Update(_371_v150, _156_v5), 5)
	(_nw72).ArraySet1(_372_v151, 6)
	(_nw72).ArraySet1((_372_v151).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_371_v150, _156_v5)), 7)
	(_nw72).ArraySet1(_372_v151, 8)
	(_nw72).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_371_v150, _156_v5)).Merge(_372_v151), 9)
	(_nw72).ArraySet1(_372_v151, 10)
	(_nw72).ArraySet1(_372_v151, 11)
	(_nw72).ArraySet1(_372_v151, 12)
	(_nw72).ArraySet1((_372_v151).Merge(_372_v151), 13)
	(_nw72).ArraySet1(_372_v151, 14)
	(_nw72).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_371_v150, _156_v5), 15)
	(_nw72).ArraySet1((_372_v151).Merge(_372_v151), 16)
	(_nw72).ArraySet1(_372_v151, 17)
	(_nw72).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_371_v150, _156_v5)).Merge(_372_v151), 18)
	(_nw72).ArraySet1(_372_v151, 19)
	(_nw72).ArraySet1(_372_v151, 20)
	(_nw72).ArraySet1(((_372_v151).Update(_371_v150, _156_v5)).Merge(_372_v151), 21)
	(_nw72).ArraySet1(_372_v151, 22)
	(_nw72).ArraySet1((_372_v151).Merge(_372_v151), 23)
	(_nw72).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_371_v150, _156_v5), 24)
	(_nw72).ArraySet1(_372_v151, 25)
	(_nw72).ArraySet1(_372_v151, 26)
	_373_v152 = _nw72
	var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_373_v152), 0))
	_ = _index70
	(_373_v152).ArraySet1(_372_v151, (_index70).Int())
	var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_373_v152), 0))
	_ = _index71
	(_373_v152).ArraySet1(_372_v151, (_index71).Int())
	var _374_v153 _dafny.MultiSet
	_ = _374_v153
	_374_v153 = _dafny.MultiSetOf(_153_v2)
	var _375_v154 _dafny.Map
	_ = _375_v154
	_375_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((Companion_D0_.Create_DC1_(false)).Dtor_cf1()), ((_237_v48).F13()) == ((_237_v48).F13()))
	var _rhs45 bool = (_219_v33).IsSubsetOf(_219_v33)
	_ = _rhs45
	var _rhs46 bool = (func() bool {
		if (_375_v154).Contains(Companion_Default___.Fm2(_162_globalState)) {
			return (_375_v154).Get(Companion_Default___.Fm2(_162_globalState)).(bool)
		}
		return !((_237_v48).F13())
	})()
	_ = _rhs46
	var _rhs47 _dafny.MultiSet = (_dafny.MultiSetOf(_153_v2, (_237_v48).F13())).Difference(_374_v153)
	_ = _rhs47
	var _rhs48 _dafny.Array = _156_v5
	_ = _rhs48
	var _rhs49 _dafny.Int = _237_v48.F14
	_ = _rhs49
	var _lhs31 *GlobalState = _162_globalState
	_ = _lhs31
	var _lhs32 *C0 = _237_v48
	_ = _lhs32
	_153_v2 = _rhs45
	_lhs31.F3 = _rhs46
	_374_v153 = _rhs47
	_156_v5 = _rhs48
	_lhs32.F14 = _rhs49
	if _153_v2 {
		var _376_v155 _dafny.MultiSet
		_ = _376_v155
		_376_v155 = _dafny.MultiSetOf(_164_v11)
		var _377_v156 *C2
		_ = _377_v156
		var _nw73 *C2 = New_C2_()
		_ = _nw73
		_nw73.Ctor__((_237_v48.F14).Cmp((func() _dafny.Int {
			if (_376_v155).Contains(_164_v11) {
				return (_376_v155).Multiplicity(_164_v11)
			}
			return _151_v0
		})()) > 0, (_237_v48).F13())
		_377_v156 = _nw73
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_161_v10), 0))
		_ = _index72
		(_161_v10).ArraySet1(_237_v48.F14, (_index72).Int())
		var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_161_v10), 0))
		_ = _index73
		(_161_v10).ArraySet1(((_323_v113).Difference(_323_v113)).Cardinality(), (_index73).Int())
		var _378_v157 T0
		_ = _378_v157
		var _nw74 *C1 = New_C1_()
		_ = _nw74
		_nw74.Ctor__((_237_v48).F13(), _237_v48.F14, (_237_v48).F13())
		_378_v157 = _nw74
		var _379_v158 _dafny.MultiSet
		_ = _379_v158
		_379_v158 = _dafny.MultiSetOf(_378_v157, _378_v157)
		var _380_v159 *C0
		_ = _380_v159
		var _nw75 *C0 = New_C0_()
		_ = _nw75
		_nw75.Ctor__((_237_v48).F13(), (_dafny.Zero).Minus((_379_v158).Cardinality()), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), _156_v5)).Cardinality()).Cmp(_237_v48.F14) != 0)
		_380_v159 = _nw75
		var _381_v160 _dafny.Array
		_ = _381_v160
		var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
		_ = _nw76
		_381_v160 = _nw76
		var _382_v161 _dafny.Map
		_ = _382_v161
		_382_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_381_v160, (_378_v157).F9())
		var _383_v162 *C1
		_ = _383_v162
		var _nw77 *C1 = New_C1_()
		_ = _nw77
		_nw77.Ctor__((func() bool {
			if (_382_v161).Contains(_381_v160) {
				return (_382_v161).Get(_381_v160).(bool)
			}
			return (_380_v159).F13()
		})(), _dafny.IntOfUint32((Companion_Default___.Fm11(Companion_Default___.Fm2(_162_globalState), (_237_v48).F13(), _162_globalState)).Cardinality()), Companion_Default___.Fm2(_162_globalState))
		_383_v162 = _nw77
		var _384_v163 _dafny.Map
		_ = _384_v163
		_384_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_159_v8).Cardinality()), _237_v48.F14)
		var _385_v164 _dafny.Sequence
		_ = _385_v164
		_385_v164 = _dafny.SeqOf(_383_v162, _383_v162, _383_v162)
		var _rhs50 *C1 = (_385_v164).Select((Companion_Default___.SafeIndex((_161_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_161_v10), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_385_v164).Cardinality()))).Uint32()).(*C1)
		_ = _rhs50
		var _rhs51 bool = Companion_Default___.Fm2(_162_globalState)
		_ = _rhs51
		var _rhs52 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v0, _dafny.IntOfUint32((_159_v8).Cardinality()))
		_ = _rhs52
		var _lhs33 *C2 = _377_v156
		_ = _lhs33
		_383_v162 = _rhs50
		_lhs33.F10 = _rhs51
		_384_v163 = _rhs52
		var _386_v165 _dafny.Map
		_ = _386_v165
		_386_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_237_v48.F14, _237_v48.F14), _380_v159.F14)
		(_237_v48).F14 = (func() _dafny.Int {
			if (_386_v165).Contains(_384_v163) {
				return (_386_v165).Get(_384_v163).(_dafny.Int)
			}
			return _237_v48.F14
		})()
	} else {
		_160_v9 = (_160_v9).Update(_153_v2, _237_v48.F14)
		var _387_v166 _dafny.Set
		_ = _387_v166
		_387_v166 = _dafny.SetOf(_156_v5, _156_v5)
		_387_v166 = _387_v166
		(_162_globalState).F3 = !(_153_v2) || ((_237_v48).F13())
		(_237_v48).F14 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(522))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_388_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_389_i12 _dafny.Int) _dafny.CodePoint {
				return _388_v11
			}
		})(_164_v11)))).Cardinality())
		var _390_v167 *C1
		_ = _390_v167
		var _nw78 *C1 = New_C1_()
		_ = _nw78
		_nw78.Ctor__((_237_v48).F13(), _237_v48.F14, _153_v2)
		_390_v167 = _nw78
		var _391_v168 _dafny.Array
		_ = _391_v168
		var _nwElement0_9 *C1 = _390_v167
		_ = _nwElement0_9
		var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(3))
		_ = _nw79
		(_nw79).ArraySet1(_nwElement0_9, 0)
		(_nw79).ArraySet1(_390_v167, 1)
		(_nw79).ArraySet1(_390_v167, 2)
		_391_v168 = _nw79
		var _392_v169 _dafny.MultiSet
		_ = _392_v169
		_392_v169 = _dafny.MultiSetOf(_dafny.MultiSetOf((_237_v48).F13(), (_237_v48).F13(), (_237_v48).F13(), (_237_v48).F13()))
		var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_156_v5), 0))
		_ = _index74
		(_156_v5).ArraySet1(_153_v2, (_index74).Int())
		var _393_v170 T0
		_ = _393_v170
		var _nw80 *C3 = New_C3_()
		_ = _nw80
		_nw80.Ctor__((_237_v48).F13())
		_393_v170 = _nw80
		var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_156_v5), 0))
		_ = _index75
		var _rhs53 _dafny.Array = _391_v168
		_ = _rhs53
		var _rhs54 _dafny.MultiSet = (_392_v169).Intersection((_392_v169).Union(_dafny.MultiSetOf(_374_v153)))
		_ = _rhs54
		var _rhs55 bool = (func() bool {
			if (_237_v48).F13() {
				return (_237_v48).F13()
			}
			return (_390_v167).Fm8((_393_v170).F9(), _162_globalState)
		})()
		_ = _rhs55
		var _rhs56 T0 = _393_v170
		_ = _rhs56
		var _lhs34 _dafny.Array = _156_v5
		_ = _lhs34
		var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(820), _dafny.ArrayLen((_156_v5), 0))
		_ = _lhs35
		_391_v168 = _rhs53
		_392_v169 = _rhs54
		(_lhs34).ArraySet1(_rhs55, (_lhs35).Int())
		_393_v170 = _rhs56
	}
	var _hi8 _dafny.Int = _237_v48.F14
	_ = _hi8
	for _394_i13 := (_dafny.Zero).Minus(_237_v48.F14); _394_i13.Cmp(_hi8) < 0; _394_i13 = _394_i13.Plus(_dafny.One) {
		_153_v2 = !((_237_v48).F13()) || ((_237_v48).F13())
		var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_156_v5), 0))
		_ = _index76
		(_156_v5).ArraySet1((_237_v48).F13(), (_index76).Int())
		var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_156_v5), 0))
		_ = _index77
		(_156_v5).ArraySet1((func() bool {
			if (_237_v48).F13() {
				return false
			}
			return (_237_v48).F13()
		})(), (_index77).Int())
		var _395_v171 _dafny.Map
		_ = _395_v171
		_395_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _236_v47)
		var _396_v172 _dafny.Map
		_ = _396_v172
		_396_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_394_i13, (_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool))
		var _nwElement0_10 _dafny.Int = _dafny.IntOfInt64(-787)
		_ = _nwElement0_10
		var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(9))
		_ = _nw81
		(_nw81).ArraySet1(_nwElement0_10, 0)
		(_nw81).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_237_v48.F14)), 1)
		(_nw81).ArraySet1((_237_v48).Fm4(_395_v171, _153_v2, (_237_v48).F13(), (_322_v112).Fm4(_395_v171, (_237_v48).F13(), false, _237_v48.F14, _162_globalState), _162_globalState), 2)
		(_nw81).ArraySet1(_237_v48.F14, 3)
		(_nw81).ArraySet1((_237_v48.F14).Minus((_374_v153).Cardinality()), 4)
		(_nw81).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_237_v48.F14)), 5)
		(_nw81).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_375_v154).Contains((_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool)) {
				return (_375_v154).Get((_156_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_156_v5), 0))).Int()).(bool)).(bool)
			}
			return Companion_Default___.Fm2(_162_globalState)
		})(), _396_v172)).Cardinality(), 6)
		(_nw81).ArraySet1(_394_i13, 7)
		(_nw81).ArraySet1(_237_v48.F14, 8)
		_161_v10 = _nw81
		(_162_globalState).F7 = Companion_Default___.Fm18(_162_globalState)
	}
	_dafny.Print(_151_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_152_v1, _dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_153_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_154_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_155_v4).Equals(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636)), _dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636)), _dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636)), _dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636)), _dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_v6).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v7).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_v8.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(636))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v10).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_162_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_162_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F5()).Equals(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636)), _dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636)), _dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636)), _dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636)), _dafny.SeqOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(636)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_162_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_162_globalState).F8(), _dafny.SeqOf(_dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_164_v11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_219_v33).Equals(_dafny.SetOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_236_v47, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_237_v48).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_237_v48.F14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_304_v96).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(636))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_323_v113).Equals(_dafny.SetOf(_dafny.IntOfInt64(-2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_324_v114).Dtor_cf27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_324_v114).Dtor_cf28())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_371_v150).Dtor_cf29()).Equals(_dafny.SetOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_372_v151).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_373_v152).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Map)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_374_v153).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_375_v154).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
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

type D0_DC2 struct {
	Cf2 bool
	Cf3 _dafny.Int
	Cf4 _dafny.MultiSet
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 bool, Cf3 _dafny.Int, Cf4 _dafny.MultiSet) D0 {
	return D0{D0_DC2{Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_() D0 {
	return D0{D0_DC3{}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
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
	return Companion_D0_.Create_DC1_(false)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) Dtor_cf4() _dafny.MultiSet {
	return _this.Get_().(D0_DC2).Cf4
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3"
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
			return ok && data1.Cf1 == data2.Cf1
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4.Equals(data2.Cf4)
		}
	case D0_DC3:
		{
			_, ok := other.Get_().(D0_DC3)
			return ok
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

type D1_DC5 struct {
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_() D1 {
	return D1{D1_DC5{}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_() D1 {
	return D1{D1_DC6{}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC4 struct {
	Cf5 _dafny.Sequence
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf5 _dafny.Sequence) D1 {
	return D1{D1_DC4{Cf5}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_()
}

func (_this D1) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5"
		}
	case D1_DC6:
		{
			return "D1.DC6"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf5) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC5:
		{
			_, ok := other.Get_().(D1_DC5)
			return ok
		}
	case D1_DC6:
		{
			_, ok := other.Get_().(D1_DC6)
			return ok
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
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
	Cf7 bool
	Cf8 _dafny.Int
	Cf9 bool
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf7 bool, Cf8 _dafny.Int, Cf9 bool) D2 {
	return D2{D2_DC8{Cf7, Cf8, Cf9}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC7 struct {
	Cf6 _dafny.Array
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf6 _dafny.Array) D2 {
	return D2{D2_DC7{Cf6}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(false, _dafny.Zero, false)
}

func (_this D2) Dtor_cf7() bool {
	return _this.Get_().(D2_DC8).Cf7
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf8
}

func (_this D2) Dtor_cf9() bool {
	return _this.Get_().(D2_DC8).Cf9
}

func (_this D2) Dtor_cf6() _dafny.Array {
	return _this.Get_().(D2_DC7).Cf6
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf7 == data2.Cf7 && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9 == data2.Cf9
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf6 == data2.Cf6
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
	Cf11 _dafny.Int
	Cf12 _dafny.Array
	Cf13 bool
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf11 _dafny.Int, Cf12 _dafny.Array, Cf13 bool) D3 {
	return D3{D3_DC10{Cf11, Cf12, Cf13}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC11 struct {
	Cf14 _dafny.Int
	Cf15 _dafny.Sequence
	Cf16 _dafny.Int
	Cf17 *C0
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf14 _dafny.Int, Cf15 _dafny.Sequence, Cf16 _dafny.Int, Cf17 *C0) D3 {
	return D3{D3_DC11{Cf14, Cf15, Cf16, Cf17}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC9 struct {
	Cf10 _dafny.Sequence
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf10 _dafny.Sequence) D3 {
	return D3{D3_DC9{Cf10}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC12 struct {
	Cf18 D3
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf18 D3) D3 {
	return D3{D3_DC12{Cf18}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D3) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf11
}

func (_this D3) Dtor_cf12() _dafny.Array {
	return _this.Get_().(D3_DC10).Cf12
}

func (_this D3) Dtor_cf13() bool {
	return _this.Get_().(D3_DC10).Cf13
}

func (_this D3) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf14
}

func (_this D3) Dtor_cf15() _dafny.Sequence {
	return _this.Get_().(D3_DC11).Cf15
}

func (_this D3) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf16
}

func (_this D3) Dtor_cf17() *C0 {
	return _this.Get_().(D3_DC11).Cf17
}

func (_this D3) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D3_DC9).Cf10
}

func (_this D3) Dtor_cf18() D3 {
	return _this.Get_().(D3_DC12).Cf18
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf14) + ", " + data.Cf15.VerbatimString(true) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + data.Cf10.VerbatimString(true) + ")"
		}
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf18) + ")"
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
			return ok && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15.Equals(data2.Cf15) && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17 == data2.Cf17
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf10.Equals(data2.Cf10)
		}
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf18.Equals(data2.Cf18)
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
	Cf20 bool
	Cf21 _dafny.Int
	Cf22 _dafny.Int
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf20 bool, Cf21 _dafny.Int, Cf22 _dafny.Int) D4 {
	return D4{D4_DC14{Cf20, Cf21, Cf22}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC15 struct {
	Cf23 _dafny.Int
	Cf24 bool
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf23 _dafny.Int, Cf24 bool) D4 {
	return D4{D4_DC15{Cf23, Cf24}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

type D4_DC13 struct {
	Cf19 _dafny.Array
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf19 _dafny.Array) D4 {
	return D4{D4_DC13{Cf19}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC16 struct {
	Cf25 D4
}

func (D4_DC16) isD4() {}

func (CompanionStruct_D4_) Create_DC16_(Cf25 D4) D4 {
	return D4{D4_DC16{Cf25}}
}

func (_this D4) Is_DC16() bool {
	_, ok := _this.Get_().(D4_DC16)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC14_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D4) Dtor_cf20() bool {
	return _this.Get_().(D4_DC14).Cf20
}

func (_this D4) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf21
}

func (_this D4) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf22
}

func (_this D4) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D4_DC15).Cf23
}

func (_this D4) Dtor_cf24() bool {
	return _this.Get_().(D4_DC15).Cf24
}

func (_this D4) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D4_DC13).Cf19
}

func (_this D4) Dtor_cf25() D4 {
	return _this.Get_().(D4_DC16).Cf25
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D4_DC16:
		{
			return "D4.DC16" + "(" + _dafny.String(data.Cf25) + ")"
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
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22.Cmp(data2.Cf22) == 0
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24 == data2.Cf24
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf19 == data2.Cf19
		}
	case D4_DC16:
		{
			data2, ok := other.Get_().(D4_DC16)
			return ok && data1.Cf25.Equals(data2.Cf25)
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

type D5_DC18 struct {
	Cf27 _dafny.Int
	Cf28 bool
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf27 _dafny.Int, Cf28 bool) D5 {
	return D5{D5_DC18{Cf27, Cf28}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

type D5_DC19 struct {
	Cf29 _dafny.Set
}

func (D5_DC19) isD5() {}

func (CompanionStruct_D5_) Create_DC19_(Cf29 _dafny.Set) D5 {
	return D5{D5_DC19{Cf29}}
}

func (_this D5) Is_DC19() bool {
	_, ok := _this.Get_().(D5_DC19)
	return ok
}

type D5_DC17 struct {
	Cf26 *C3
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf26 *C3) D5 {
	return D5{D5_DC17{Cf26}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC18_(_dafny.Zero, false)
}

func (_this D5) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D5_DC18).Cf27
}

func (_this D5) Dtor_cf28() bool {
	return _this.Get_().(D5_DC18).Cf28
}

func (_this D5) Dtor_cf29() _dafny.Set {
	return _this.Get_().(D5_DC19).Cf29
}

func (_this D5) Dtor_cf26() *C3 {
	return _this.Get_().(D5_DC17).Cf26
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D5_DC19:
		{
			return "D5.DC19" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf26) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
			return ok && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28 == data2.Cf28
		}
	case D5_DC19:
		{
			data2, ok := other.Get_().(D5_DC19)
			return ok && data1.Cf29.Equals(data2.Cf29)
		}
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf26 == data2.Cf26
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

type D6_DC21 struct {
	Cf31 _dafny.Sequence
	Cf32 bool
	Cf33 *C0
	Cf34 _dafny.Map
	Cf35 _dafny.CodePoint
}

func (D6_DC21) isD6() {}

func (CompanionStruct_D6_) Create_DC21_(Cf31 _dafny.Sequence, Cf32 bool, Cf33 *C0, Cf34 _dafny.Map, Cf35 _dafny.CodePoint) D6 {
	return D6{D6_DC21{Cf31, Cf32, Cf33, Cf34, Cf35}}
}

func (_this D6) Is_DC21() bool {
	_, ok := _this.Get_().(D6_DC21)
	return ok
}

type D6_DC22 struct {
	Cf36 _dafny.Int
}

func (D6_DC22) isD6() {}

func (CompanionStruct_D6_) Create_DC22_(Cf36 _dafny.Int) D6 {
	return D6{D6_DC22{Cf36}}
}

func (_this D6) Is_DC22() bool {
	_, ok := _this.Get_().(D6_DC22)
	return ok
}

type D6_DC23 struct {
	Cf37 bool
	Cf38 bool
	Cf39 _dafny.MultiSet
}

func (D6_DC23) isD6() {}

func (CompanionStruct_D6_) Create_DC23_(Cf37 bool, Cf38 bool, Cf39 _dafny.MultiSet) D6 {
	return D6{D6_DC23{Cf37, Cf38, Cf39}}
}

func (_this D6) Is_DC23() bool {
	_, ok := _this.Get_().(D6_DC23)
	return ok
}

type D6_DC20 struct {
	Cf30 _dafny.Map
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_(Cf30 _dafny.Map) D6 {
	return D6{D6_DC20{Cf30}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

type D6_DC24 struct {
	Cf40 D6
}

func (D6_DC24) isD6() {}

func (CompanionStruct_D6_) Create_DC24_(Cf40 D6) D6 {
	return D6{D6_DC24{Cf40}}
}

func (_this D6) Is_DC24() bool {
	_, ok := _this.Get_().(D6_DC24)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC21_(_dafny.EmptySeq, false, (*C0)(nil), _dafny.EmptyMap, _dafny.CodePoint('D'))
}

func (_this D6) Dtor_cf31() _dafny.Sequence {
	return _this.Get_().(D6_DC21).Cf31
}

func (_this D6) Dtor_cf32() bool {
	return _this.Get_().(D6_DC21).Cf32
}

func (_this D6) Dtor_cf33() *C0 {
	return _this.Get_().(D6_DC21).Cf33
}

func (_this D6) Dtor_cf34() _dafny.Map {
	return _this.Get_().(D6_DC21).Cf34
}

func (_this D6) Dtor_cf35() _dafny.CodePoint {
	return _this.Get_().(D6_DC21).Cf35
}

func (_this D6) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D6_DC22).Cf36
}

func (_this D6) Dtor_cf37() bool {
	return _this.Get_().(D6_DC23).Cf37
}

func (_this D6) Dtor_cf38() bool {
	return _this.Get_().(D6_DC23).Cf38
}

func (_this D6) Dtor_cf39() _dafny.MultiSet {
	return _this.Get_().(D6_DC23).Cf39
}

func (_this D6) Dtor_cf30() _dafny.Map {
	return _this.Get_().(D6_DC20).Cf30
}

func (_this D6) Dtor_cf40() D6 {
	return _this.Get_().(D6_DC24).Cf40
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC21:
		{
			return "D6.DC21" + "(" + data.Cf31.VerbatimString(true) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D6_DC22:
		{
			return "D6.DC22" + "(" + _dafny.String(data.Cf36) + ")"
		}
	case D6_DC23:
		{
			return "D6.DC23" + "(" + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D6_DC20:
		{
			return "D6.DC20" + "(" + _dafny.String(data.Cf30) + ")"
		}
	case D6_DC24:
		{
			return "D6.DC24" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC21:
		{
			data2, ok := other.Get_().(D6_DC21)
			return ok && data1.Cf31.Equals(data2.Cf31) && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33 && data1.Cf34.Equals(data2.Cf34) && data1.Cf35 == data2.Cf35
		}
	case D6_DC22:
		{
			data2, ok := other.Get_().(D6_DC22)
			return ok && data1.Cf36.Cmp(data2.Cf36) == 0
		}
	case D6_DC23:
		{
			data2, ok := other.Get_().(D6_DC23)
			return ok && data1.Cf37 == data2.Cf37 && data1.Cf38 == data2.Cf38 && data1.Cf39.Equals(data2.Cf39)
		}
	case D6_DC20:
		{
			data2, ok := other.Get_().(D6_DC20)
			return ok && data1.Cf30.Equals(data2.Cf30)
		}
	case D6_DC24:
		{
			data2, ok := other.Get_().(D6_DC24)
			return ok && data1.Cf40.Equals(data2.Cf40)
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

type D7_DC26 struct {
	Cf42 _dafny.Sequence
}

func (D7_DC26) isD7() {}

func (CompanionStruct_D7_) Create_DC26_(Cf42 _dafny.Sequence) D7 {
	return D7{D7_DC26{Cf42}}
}

func (_this D7) Is_DC26() bool {
	_, ok := _this.Get_().(D7_DC26)
	return ok
}

type D7_DC25 struct {
	Cf41 _dafny.Sequence
}

func (D7_DC25) isD7() {}

func (CompanionStruct_D7_) Create_DC25_(Cf41 _dafny.Sequence) D7 {
	return D7{D7_DC25{Cf41}}
}

func (_this D7) Is_DC25() bool {
	_, ok := _this.Get_().(D7_DC25)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC26_(_dafny.EmptySeq)
}

func (_this D7) Dtor_cf42() _dafny.Sequence {
	return _this.Get_().(D7_DC26).Cf42
}

func (_this D7) Dtor_cf41() _dafny.Sequence {
	return _this.Get_().(D7_DC25).Cf41
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC26:
		{
			return "D7.DC26" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D7_DC25:
		{
			return "D7.DC25" + "(" + _dafny.String(data.Cf41) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC26:
		{
			data2, ok := other.Get_().(D7_DC26)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	case D7_DC25:
		{
			data2, ok := other.Get_().(D7_DC25)
			return ok && data1.Cf41.Equals(data2.Cf41)
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm4(p0 _dafny.Map, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int
	F9() bool
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

// Definition of class GlobalState
type GlobalState struct {
	F0  bool
	F3  bool
	F7  _dafny.CodePoint
	_f1 _dafny.Int
	_f2 bool
	_f4 bool
	_f5 _dafny.MultiSet
	_f6 _dafny.Array
	_f8 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this.F3 = false
	_this.F7 = _dafny.CodePoint('D')
	_this._f1 = _dafny.Zero
	_this._f2 = false
	_this._f4 = false
	_this._f5 = _dafny.EmptyMultiSet
	_this._f6 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f8 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 bool, f3 bool, f4 bool, f5 _dafny.MultiSet, f6 _dafny.Array, f7 _dafny.CodePoint, f8 _dafny.Sequence) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
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
func (_this *GlobalState) F6() _dafny.Array {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F8() _dafny.Sequence {
	{
		return _this._f8
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f9  bool
	F14  _dafny.Int
	_f13 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f9 = false
	_this.F14 = _dafny.Zero
	_this._f13 = false
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

func (_this *C0) F9() bool {
	return _this._f9
}
func (_this *C0) Ctor__(f13 bool, f14 _dafny.Int, f9 bool) {
	{
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this)._f9 = f9
	}
}
func (_this *C0) Fm4(p0 _dafny.Map, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt(_this.F14, (_dafny.Zero).Minus(_dafny.IntOfInt64(-993)))
	}
}
func (_this *C0) F13() bool {
	{
		return _this._f13
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f9  bool
	_f11 bool
	_f12 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f9 = false
	_this._f11 = false
	_this._f12 = _dafny.Zero
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

func (_this *C1) F9() bool {
	return _this._f9
}
func (_this *C1) Ctor__(f11 bool, f12 _dafny.Int, f9 bool) {
	{
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f9 = f9
	}
}
func (_this *C1) Fm4(p0 _dafny.Map, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F12()).Times((_this).F12())
	}
}
func (_this *C1) Fm8(p0 bool, globalState *GlobalState) bool {
	{
		return (Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), !(false))).Cardinality(), (_this).F12())).Cmp((_this).F12()) >= 0
	}
}
func (_this *C1) Fm9(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_this).F11()
	}
}
func (_this *C1) M4(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) _dafny.CodePoint {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var _397_v0 _dafny.Array
		_ = _397_v0
		var _nw82 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(15))
		_ = _nw82
		_397_v0 = _nw82
		var _398_v1 D0
		_ = _398_v1
		_398_v1 = Companion_D0_.Create_DC3_()
		var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_397_v0), 0))
		_ = _index78
		(_397_v0).ArraySet1(_398_v1, (_index78).Int())
		var _399_v2 _dafny.Array
		_ = _399_v2
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_6
		var _nw83 _dafny.Array
		_ = _nw83
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw83 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) bool = func(_400_i0 _dafny.Int) bool {
				return (_this).F11()
			}
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw83 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw83).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw83).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_399_v2 = _nw83
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
		_ = _index79
		(_399_v2).ArraySet1(((_this).F12()).Cmp((_this).F12()) > 0, (_index79).Int())
		var _401_v3 _dafny.Int
		_ = _401_v3
		_401_v3 = _dafny.IntOfInt64(229)
		var _402_v4 _dafny.MultiSet
		_ = _402_v4
		_402_v4 = _dafny.MultiSetOf((_this).F9(), (_this).F9(), (_this).F9())
		var _403_v5 _dafny.Sequence
		_ = _403_v5
		_403_v5 = _dafny.SeqOf(_402_v4)
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_397_v0), 0))
		_ = _index80
		var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
		_ = _index81
		var _rhs57 bool = (_this).Fm8((_this).F9(), globalState)
		_ = _rhs57
		var _rhs58 D0 = _398_v1
		_ = _rhs58
		var _rhs59 bool = (_this).F9()
		_ = _rhs59
		var _rhs60 _dafny.Int = ((_403_v5).Select((Companion_Default___.SafeIndex(((_this).F12()).Minus(_dafny.IntOfInt64(-958)), _dafny.IntOfUint32((_403_v5).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()
		_ = _rhs60
		var _lhs36 *GlobalState = globalState
		_ = _lhs36
		var _lhs37 _dafny.Array = _397_v0
		_ = _lhs37
		var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_397_v0), 0))
		_ = _lhs38
		var _lhs39 _dafny.Array = _399_v2
		_ = _lhs39
		var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
		_ = _lhs40
		_lhs36.F0 = _rhs57
		(_lhs37).ArraySet1(_rhs58, (_lhs38).Int())
		(_lhs39).ArraySet1(_rhs59, (_lhs40).Int())
		_401_v3 = _rhs60
		var _source6 D0 = (func() D0 {
			if (_this).F11() {
				return (_397_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_397_v0), 0))).Int()).(D0)
			}
			return _398_v1
		})()
		_ = _source6
		if _source6.Is_DC1() {
			var _404___mcc_h0 bool = _source6.Get_().(D0_DC1).Cf1
			_ = _404___mcc_h0
			var _405_cf1 bool = _404___mcc_h0
			_ = _405_cf1
			var _406_v6 *C0
			_ = _406_v6
			var _nw84 *C0 = New_C0_()
			_ = _nw84
			_nw84.Ctor__((_this).F11(), p1, (_this).F9())
			_406_v6 = _nw84
			var _407_v7 _dafny.Array
			_ = _407_v7
			var _nw85 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw85
			_407_v7 = _nw85
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_407_v7), 0))
			_ = _index82
			(_407_v7).ArraySet1((_401_v3).Minus(_dafny.IntOfInt64(316)), (_index82).Int())
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_407_v7), 0))
			_ = _index83
			(_407_v7).ArraySet1(_406_v6.F14, (_index83).Int())
			var _408_v8 _dafny.Map
			_ = _408_v8
			_408_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
			var _409_v9 _dafny.Int
			_ = _409_v9
			var _out18 _dafny.Int
			_ = _out18
			_out18 = Companion_Default___.M0(_406_v6.F14, _407_v7, (_this).Fm9((_407_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_407_v7), 0))).Int()).(_dafny.Int), (_406_v6).F13(), (_this).F12(), globalState), _408_v8, globalState)
			_409_v9 = _out18
			(globalState).F0 = (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool)
		} else if _source6.Is_DC2() {
			var _410___mcc_h1 bool = _source6.Get_().(D0_DC2).Cf2
			_ = _410___mcc_h1
			var _411___mcc_h2 _dafny.Int = _source6.Get_().(D0_DC2).Cf3
			_ = _411___mcc_h2
			var _412___mcc_h3 _dafny.MultiSet = _source6.Get_().(D0_DC2).Cf4
			_ = _412___mcc_h3
			var _413_cf4 _dafny.MultiSet = _412___mcc_h3
			_ = _413_cf4
			var _414_cf3 _dafny.Int = _411___mcc_h2
			_ = _414_cf3
			var _415_cf2 bool = _410___mcc_h1
			_ = _415_cf2
			_415_cf2 = (_this).F11()
			_399_v2 = _399_v2
			var _416_v10 _dafny.Map
			_ = _416_v10
			_416_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_402_v4), (_this).F11())
			var _417_v11 _dafny.Sequence
			_ = _417_v11
			_417_v11 = _dafny.SeqOf(true)
			var _418_v12 _dafny.Set
			_ = _418_v12
			_418_v12 = _dafny.SetOf(_dafny.MultiSetFromSeq(_417_v11), _402_v4)
			_416_v10 = (_416_v10).Update(_418_v12, !((_this).F9()))
			var _419_v13 D0
			_ = _419_v13
			_419_v13 = Companion_D0_.Create_DC1_(_dafny.Companion_Sequence_.IsPrefixOf(_417_v11, _417_v11))
			var _source7 D0 = _419_v13
			_ = _source7
			if _source7.Is_DC1() {
				var _420___mcc_h5 bool = _source7.Get_().(D0_DC1).Cf1
				_ = _420___mcc_h5
				var _421_cf1 bool = _420___mcc_h5
				_ = _421_cf1
				_401_v3 = _414_cf3
				(globalState).F0 = (_417_v11).Select((Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_417_v11).Cardinality()))).Uint32()).(bool)
				var _422_v14 _dafny.Map
				_ = _422_v14
				_422_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _415_cf2)
				var _423_v15 _dafny.Sequence
				_ = _423_v15
				_423_v15 = _dafny.SeqOf(_417_v11, _417_v11, _dafny.SeqOf((func() bool {
					if (_422_v14).Contains(_dafny.IntOfInt64(874)) {
						return (_422_v14).Get(_dafny.IntOfInt64(874)).(bool)
					}
					return (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool)
				})()))
				var _424_v16 _dafny.Array
				_ = _424_v16
				var _nwElement0_11 _dafny.Sequence = _417_v11
				_ = _nwElement0_11
				var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(4))
				_ = _nw86
				(_nw86).ArraySet1(_nwElement0_11, 0)
				(_nw86).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_417_v11, _417_v11), 1)
				(_nw86).ArraySet1(_417_v11, 2)
				(_nw86).ArraySet1((_423_v15).Select((Companion_Default___.SafeIndex(_414_cf3, _dafny.IntOfUint32((_423_v15).Cardinality()))).Uint32()).(_dafny.Sequence), 3)
				_424_v16 = _nw86
				_424_v16 = _424_v16
				var _rhs61 _dafny.Map = _422_v14
				_ = _rhs61
				var _rhs62 _dafny.Int = ((func() _dafny.Int {
					if (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool) {
						return (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.SetOf(_415_cf2)).Cardinality()))
					}
					return (_this).F12()
				})()).Times((_this).F12())
				_ = _rhs62
				var _rhs63 _dafny.Array = (func() _dafny.Array {
					if (_this).Fm8((_this).F11(), globalState) {
						return _399_v2
					}
					return _399_v2
				})()
				_ = _rhs63
				_422_v14 = _rhs61
				_414_cf3 = _rhs62
				_399_v2 = _rhs63
			} else if _source7.Is_DC2() {
				var _425___mcc_h6 bool = _source7.Get_().(D0_DC2).Cf2
				_ = _425___mcc_h6
				var _426___mcc_h7 _dafny.Int = _source7.Get_().(D0_DC2).Cf3
				_ = _426___mcc_h7
				var _427___mcc_h8 _dafny.MultiSet = _source7.Get_().(D0_DC2).Cf4
				_ = _427___mcc_h8
				var _428_cf4 _dafny.MultiSet = _427___mcc_h8
				_ = _428_cf4
				var _429_cf3 _dafny.Int = _426___mcc_h7
				_ = _429_cf3
				var _430_cf2 bool = _425___mcc_h6
				_ = _430_cf2
				var _431_v17 _dafny.Map
				_ = _431_v17
				_431_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F9()), (_this).F11())
				var _432_v18 _dafny.Map
				_ = _432_v18
				_432_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _401_v3)
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
				_ = _index84
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
				_ = _index85
				var _rhs64 bool = ((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool)) == ((_this).F9())
				_ = _rhs64
				var _rhs65 bool = (_417_v11).Select((Companion_Default___.SafeIndex((_431_v17).Cardinality(), _dafny.IntOfUint32((_417_v11).Cardinality()))).Uint32()).(bool)
				_ = _rhs65
				var _rhs66 bool = (_this).F11()
				_ = _rhs66
				var _rhs67 _dafny.Int = _401_v3
				_ = _rhs67
				var _rhs68 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(false), (Companion_Default___.SafeIndex(_401_v3, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Uint32(), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool)), _417_v11), (func() _dafny.Sequence {
					if (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool) {
						return _dafny.SeqOf(_430_cf2, (_this).F9(), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))
					}
					return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm10(globalState), (Companion_Default___.SafeIndex((_432_v18).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm10(globalState)).Cardinality()))).Uint32(), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool)), (Companion_Default___.SafeIndex(_401_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm10(globalState), (Companion_Default___.SafeIndex((_432_v18).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm10(globalState)).Cardinality()))).Uint32(), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))).Cardinality()))).Uint32(), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))
				})())
				_ = _rhs68
				var _lhs41 *GlobalState = globalState
				_ = _lhs41
				var _lhs42 _dafny.Array = _399_v2
				_ = _lhs42
				var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
				_ = _lhs43
				var _lhs44 _dafny.Array = _399_v2
				_ = _lhs44
				var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
				_ = _lhs45
				_lhs41.F3 = _rhs64
				(_lhs42).ArraySet1(_rhs65, (_lhs43).Int())
				(_lhs44).ArraySet1(_rhs66, (_lhs45).Int())
				_429_cf3 = _rhs67
				_417_v11 = _rhs68
				(globalState).F0 = _430_cf2
				_401_v3 = p1
				var _433_v19 _dafny.Map
				_ = _433_v19
				_433_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_430_cf2, p2)
				var _434_v21 _dafny.Set
				_ = _434_v21
				_434_v21 = _dafny.SetOf(_401_v3, (_this).F12(), (func() _dafny.Map {
					var _coll10 = _dafny.NewMapBuilder()
					_ = _coll10
					for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-922), _dafny.IntOfInt64(289))); ; {
						_compr_10, _ok10 := _iter10()
						if !_ok10 {
							break
						}
						var _435_v20 _dafny.Int
						_435_v20 = interface{}(_compr_10).(_dafny.Int)
						if ((_dafny.IntOfInt64(-922)).Cmp(_435_v20) <= 0) && ((_435_v20).Cmp(_dafny.IntOfInt64(289)) < 0) {
							_coll10.Add((_435_v20).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("aee"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(640))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg21 _dafny.Int) interface{} {
									return coer21(arg21)
								}
							}(func(_436_i1 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('b')
							})))).Cardinality())), (_this).F9())
						}
					}
					return _coll10.ToMap()
				}()).Cardinality(), _429_cf3, (_this).F12())
				var _437_v22 _dafny.Sequence
				_ = _437_v22
				_437_v22 = _dafny.SeqOf(_434_v21)
				_433_v19 = (_433_v19).Update((_434_v21).IsProperSubsetOf((_437_v22).Select((Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_437_v22).Cardinality()))).Uint32()).(_dafny.Set)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(913))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}((func(_438_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_439_i2 _dafny.Int) _dafny.Int {
						return _438_v3
					}
				})(_401_v3))))
			} else if _source7.Is_DC3() {
				var _440_v23 _dafny.CodePoint
				_ = _440_v23
				_440_v23 = _dafny.CodePoint('q')
				var _441_v24 _dafny.Map
				_ = _441_v24
				_441_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _417_v11)
				var _442_v25 _dafny.Map
				_ = _442_v25
				_442_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_440_v23, (_this).Fm4(_441_v24, (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), (_this).Fm8(_415_cf2, globalState), (_this).F12(), globalState))
				_442_v25 = (_442_v25).Update(_440_v23, (_this).F12())
				var _443_v26 _dafny.Map
				_ = _443_v26
				_443_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), !(_415_cf2))
				_443_v26 = (_443_v26).Update((_this).Fm8((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), globalState), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))
				var _444_v27 _dafny.Array
				_ = _444_v27
				var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(16))
				_ = _nw87
				_444_v27 = _nw87
				_444_v27 = _444_v27
				var _445_v28 _dafny.Sequence
				_ = _445_v28
				_445_v28 = _dafny.SeqOf(_401_v3, (_this).F12(), Companion_Default___.SafeModuloInt(_401_v3, (_this).F12()))
				_445_v28 = p2
			} else {
				var _446___mcc_h9 _dafny.Int = _source7.Get_().(D0_DC0).Cf0
				_ = _446___mcc_h9
				var _447_cf0 _dafny.Int = _446___mcc_h9
				_ = _447_cf0
				var _448_v29 _dafny.Map
				_ = _448_v29
				_448_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), _414_cf3)
				var _449_v30 _dafny.Map
				_ = _449_v30
				_449_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _448_v29)
				var _450_v31 D0
				_ = _450_v31
				_450_v31 = Companion_D0_.Create_DC2_((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), (_449_v30).Cardinality(), _413_cf4)
				var _451_v32 _dafny.MultiSet
				_ = _451_v32
				_451_v32 = _dafny.MultiSetOf(_450_v31, _450_v31)
				var _452_v33 _dafny.Map
				_ = _452_v33
				_452_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_414_cf3, _415_cf2)
				var _453_v34 _dafny.Set
				_ = _453_v34
				_453_v34 = _dafny.SetOf(_414_cf3, (_452_v33).Cardinality())
				var _454_v35 _dafny.CodePoint
				_ = _454_v35
				_454_v35 = _dafny.CodePoint('k')
				var _455_v36 D1
				_ = _455_v36
				_455_v36 = Companion_D1_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(106))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg23 _dafny.Int) interface{} {
						return coer23(arg23)
					}
				}((func(_456_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_457_i3 _dafny.Int) _dafny.Int {
						return _456_p1
					}
				})(p1))))
				var _rhs69 bool = (_451_v32).IsProperSubsetOf((func() _dafny.MultiSet {
					if (_this).F9() {
						return _451_v32
					}
					return _451_v32
				})())
				_ = _rhs69
				var _rhs70 bool = ((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool)) == ((_453_v34).Equals(_453_v34))
				_ = _rhs70
				var _rhs71 _dafny.CodePoint = _454_v35
				_ = _rhs71
				var _rhs72 bool = (_this).F9()
				_ = _rhs72
				var _rhs73 _dafny.Int = _dafny.IntOfUint32(((_455_v36).Dtor_cf5()).Cardinality())
				_ = _rhs73
				var _lhs46 *GlobalState = globalState
				_ = _lhs46
				var _lhs47 *GlobalState = globalState
				_ = _lhs47
				var _lhs48 *GlobalState = globalState
				_ = _lhs48
				_415_cf2 = _rhs69
				_lhs46.F3 = _rhs70
				_lhs47.F7 = _rhs71
				_lhs48.F3 = _rhs72
				_414_cf3 = _rhs73
				var _458_v37 _dafny.Set
				_ = _458_v37
				_458_v37 = _dafny.SetOf(_399_v2, _399_v2, _399_v2)
				var _459_v38 _dafny.Sequence
				_ = _459_v38
				_459_v38 = _dafny.SeqOf(_458_v37, _458_v37, _458_v37, (_458_v37).Union(_458_v37))
				var _460_v39 _dafny.Sequence
				_ = _460_v39
				_460_v39 = _dafny.UnicodeSeqOfUtf8Bytes("rrmm")
				_458_v37 = (_459_v38).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.Fm3((_460_v39).Select((Companion_Default___.SafeIndex(_414_cf3, _dafny.IntOfUint32((_460_v39).Cardinality()))).Uint32()).(_dafny.CodePoint), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), _dafny.IntOfInt64(-349), globalState)), _dafny.IntOfUint32((_459_v38).Cardinality()))).Uint32()).(_dafny.Set)
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
				_ = _index86
				(_399_v2).ArraySet1((_450_v31).Dtor_cf2(), (_index86).Int())
				var _461_v40 _dafny.Map
				_ = _461_v40
				_461_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm11((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), globalState), _460_v39))
				_447_cf0 = (_461_v40).Cardinality()
			}
		} else if _source6.Is_DC3() {
			var _462_v41 _dafny.MultiSet
			_ = _462_v41
			_462_v41 = _dafny.MultiSetOf(_401_v3, p1, _401_v3)
			var _463_v42 D0
			_ = _463_v42
			_463_v42 = Companion_D0_.Create_DC1_((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))
			var _464_v43 *C0
			_ = _464_v43
			var _nw88 *C0 = New_C0_()
			_ = _nw88
			_nw88.Ctor__((_this).Fm9(p1, (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), p1, globalState), (func() _dafny.Int {
				if (_this).F11() {
					return (func() _dafny.Int {
						if (_462_v41).Contains(_dafny.IntOfInt64(973)) {
							return (_462_v41).Multiplicity(_dafny.IntOfInt64(973))
						}
						return _401_v3
					})()
				}
				return _401_v3
			})(), (_463_v42).Dtor_cf1())
			_464_v43 = _nw88
			_464_v43 = _464_v43
			var _465_v44 _dafny.Sequence
			_ = _465_v44
			_465_v44 = _dafny.UnicodeSeqOfUtf8Bytes("h")
			var _466_v45 _dafny.Map
			_ = _466_v45
			_466_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_465_v44, _464_v43.F14)
			var _467_v46 _dafny.Map
			_ = _467_v46
			_467_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_466_v45).Merge(_466_v45), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))
			_467_v46 = (_467_v46).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("lh"), _401_v3), (_this).F11())
			if (func() bool {
				if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("hgfq"), _465_v44) {
					return (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool)
				}
				return (_this).F11()
			})() {
				var _468_v47 _dafny.Sequence
				_ = _468_v47
				_468_v47 = _dafny.SeqOf(true)
				var _469_v48 _dafny.Map
				_ = _469_v48
				_469_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_468_v47).Select((Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_468_v47).Cardinality()))).Uint32()).(bool))
				var _470_v49 _dafny.Array
				_ = _470_v49
				var _nwElement0_12 _dafny.Int = ((_462_v41).Difference(_462_v41)).Cardinality()
				_ = _nwElement0_12
				var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(16))
				_ = _nw89
				(_nw89).ArraySet1(_nwElement0_12, 0)
				(_nw89).ArraySet1((func() _dafny.Int {
					if (_464_v43).F13() {
						return _401_v3
					}
					return _dafny.IntOfUint32((_465_v44).Cardinality())
				})(), 1)
				(_nw89).ArraySet1(_401_v3, 2)
				(_nw89).ArraySet1(_dafny.IntOfUint32((_465_v44).Cardinality()), 3)
				(_nw89).ArraySet1((_dafny.Zero).Minus(_401_v3), 4)
				(_nw89).ArraySet1(_dafny.IntOfInt64(-221), 5)
				(_nw89).ArraySet1((func() _dafny.Int {
					if (_464_v43).F13() {
						return _464_v43.F14
					}
					return _464_v43.F14
				})(), 6)
				(_nw89).ArraySet1(_dafny.IntOfUint32((_465_v44).Cardinality()), 7)
				(_nw89).ArraySet1(_464_v43.F14, 8)
				(_nw89).ArraySet1(_464_v43.F14, 9)
				(_nw89).ArraySet1(_464_v43.F14, 10)
				(_nw89).ArraySet1(_464_v43.F14, 11)
				(_nw89).ArraySet1((func() _dafny.Int {
					if (_this).F11() {
						return (_469_v48).Cardinality()
					}
					return (_dafny.SetOf(_464_v43, _464_v43)).Cardinality()
				})(), 12)
				(_nw89).ArraySet1((_dafny.Zero).Minus(_464_v43.F14), 13)
				(_nw89).ArraySet1(p1, 14)
				(_nw89).ArraySet1(_dafny.IntOfInt64(728), 15)
				_470_v49 = _nw89
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_470_v49), 0))
				_ = _index87
				(_470_v49).ArraySet1(((_this).F12()).Minus(p1), (_index87).Int())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_470_v49), 0))
				_ = _index88
				(_470_v49).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_464_v43.F14, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_465_v44, _465_v44)).Cardinality()))), (_index88).Int())
				var _471_v50 D2
				_ = _471_v50
				_471_v50 = Companion_D2_.Create_DC7_(_399_v2)
				var _472_v51 _dafny.Sequence
				_ = _472_v51
				_472_v51 = _dafny.SeqOf(_399_v2, _399_v2, (_471_v50).Dtor_cf6())
				_472_v51 = (func() _dafny.Sequence {
					if (_this).F11() {
						return (func() _dafny.Sequence {
							if (_this).F9() {
								return _472_v51
							}
							return _472_v51
						})()
					}
					return _472_v51
				})()
				var _473_v52 _dafny.MultiSet
				_ = _473_v52
				_473_v52 = _dafny.MultiSetOf(_462_v41, _462_v41, _462_v41)
				var _474_v53 D0
				_ = _474_v53
				_474_v53 = Companion_D0_.Create_DC2_((_464_v43).F13(), p1, _473_v52)
				var _475_v54 D2
				_ = _475_v54
				_475_v54 = Companion_D2_.Create_DC8_((_474_v53).Dtor_cf2(), (_this).F12(), (_464_v43).F13())
				var _476_v55 *C0
				_ = _476_v55
				var _nw90 *C0 = New_C0_()
				_ = _nw90
				_nw90.Ctor__((_this).F11(), (_475_v54).Dtor_cf8(), (_464_v43).F13())
				_476_v55 = _nw90
				var _477_v56 _dafny.CodePoint
				_ = _477_v56
				_477_v56 = _dafny.CodePoint('i')
				(globalState).F7 = _477_v56
				(_464_v43).F14 = _464_v43.F14
			} else {
				var _478_v57 _dafny.Map
				_ = _478_v57
				_478_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_464_v43.F14, Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(186), _464_v43.F14))
				var _479_v58 _dafny.CodePoint
				_ = _479_v58
				_479_v58 = _dafny.CodePoint('r')
				var _480_v59 _dafny.Map
				_ = _480_v59
				_480_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), Companion_Default___.Fm12(true, (_464_v43).F13(), _464_v43.F14, true, globalState))
				var _481_v60 _dafny.Sequence
				_ = _481_v60
				_481_v60 = _dafny.SeqOf((_this).F11(), (_this).F11(), (_464_v43).F13(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_479_v58, _402_v4)).Equals(_480_v59))
				var _482_v61 _dafny.Map
				_ = _482_v61
				_482_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_464_v43).F13(), _479_v58)
				var _483_v62 _dafny.Set
				_ = _483_v62
				_483_v62 = _dafny.SetOf((_this).Fm8((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), globalState))
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
				_ = _index89
				var _rhs74 bool = !((_481_v60).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_464_v43.F14, p1), _dafny.IntOfUint32((_481_v60).Cardinality()))).Uint32()).(bool))
				_ = _rhs74
				var _rhs75 _dafny.Int = (((_482_v61).Merge(_482_v61)).Cardinality()).Plus((_483_v62).Cardinality())
				_ = _rhs75
				var _rhs76 bool = (_464_v43).F13()
				_ = _rhs76
				var _rhs77 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.IntOfUint32((_465_v44).Cardinality())).Times(_464_v43.F14))
				_ = _rhs77
				var _lhs49 *GlobalState = globalState
				_ = _lhs49
				var _lhs50 *C0 = _464_v43
				_ = _lhs50
				var _lhs51 _dafny.Array = _399_v2
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
				_ = _lhs52
				_lhs49.F0 = _rhs74
				_lhs50.F14 = _rhs75
				(_lhs51).ArraySet1(_rhs76, (_lhs52).Int())
				_478_v57 = _rhs77
				(_464_v43).F14 = _464_v43.F14
				(_464_v43).F14 = (_this).F12()
				_465_v44 = _465_v44
				var _484_v63 _dafny.Map
				_ = _484_v63
				_484_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_464_v43.F14, _dafny.UnicodeSeqOfUtf8Bytes("xncydgy"))
				var _485_v64 D0
				_ = _485_v64
				_485_v64 = Companion_D0_.Create_DC0_(_464_v43.F14)
				_484_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_485_v64).Dtor_cf0(), Companion_Default___.Fm11((_464_v43).F13(), (_this).F11(), globalState))
			}
			_465_v44 = _dafny.Companion_Sequence_.Concatenate(_465_v44, _465_v44)
		} else {
			var _486___mcc_h4 _dafny.Int = _source6.Get_().(D0_DC0).Cf0
			_ = _486___mcc_h4
			var _487_cf0 _dafny.Int = _486___mcc_h4
			_ = _487_cf0
			var _488_v65 _dafny.CodePoint
			_ = _488_v65
			_488_v65 = _dafny.CodePoint('j')
			var _489_v66 _dafny.Sequence
			_ = _489_v66
			_489_v66 = _dafny.UnicodeSeqOfUtf8Bytes("p")
			var _490_v67 _dafny.MultiSet
			_ = _490_v67
			_490_v67 = _dafny.MultiSetOf(Companion_Default___.Fm3(_488_v65, (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), (_this).F12(), globalState), _dafny.IntOfUint32((_489_v66).Cardinality()))
			var _491_v68 D0
			_ = _491_v68
			_491_v68 = Companion_D0_.Create_DC0_(_487_cf0)
			_401_v3 = ((func() D0 {
				if (_this).F9() {
					return Companion_D0_.Create_DC0_((func() _dafny.Int {
						if (_490_v67).Contains(_dafny.IntOfInt64(278)) {
							return (_490_v67).Multiplicity(_dafny.IntOfInt64(278))
						}
						return (_this).F12()
					})())
				}
				return _491_v68
			})()).Dtor_cf0()
			var _492_v69 *C0
			_ = _492_v69
			var _nw91 *C0 = New_C0_()
			_ = _nw91
			_nw91.Ctor__(false, (func() _dafny.Int {
				if (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool) {
					return (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_487_cf0)))
				}
				return _401_v3
			})(), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))
			_492_v69 = _nw91
			var _493_v70 _dafny.Set
			_ = _493_v70
			_493_v70 = _dafny.SetOf(true)
			_493_v70 = (_493_v70).Union(_493_v70)
			var _494_v71 _dafny.Map
			_ = _494_v71
			_494_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_492_v69).F13(), _492_v69.F14)
			var _495_v72 D3
			_ = _495_v72
			_495_v72 = Companion_D3_.Create_DC11_(_487_cf0, _489_v66, (_494_v71).Cardinality(), _492_v69)
			var _496_v73 _dafny.Sequence
			_ = _496_v73
			_496_v73 = _dafny.SeqOf((_495_v72).Dtor_cf15(), _489_v66, _489_v66)
			var _497_v74 _dafny.Map
			_ = _497_v74
			_497_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(692))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_498_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_499_i4 _dafny.Int) _dafny.Int {
					return _498_v3
				}
			})(_401_v3))), p2), _496_v73)
			_496_v73 = (func() _dafny.Sequence {
				if (_497_v74).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-175))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg25 _dafny.Int) interface{} {
						return coer25(arg25)
					}
				}((func(_500_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_501_i5 _dafny.Int) _dafny.Int {
						return _500_p1
					}
				})(p1)))) {
					return (_497_v74).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-175))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg26 _dafny.Int) interface{} {
							return coer26(arg26)
						}
					}((func(_502_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_503_i5 _dafny.Int) _dafny.Int {
							return _502_p1
						}
					})(p1)))).(_dafny.Sequence)
				}
				return _496_v73
			})()
		}
		var _504_v75 _dafny.Array
		_ = _504_v75
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_7
		var _nw92 _dafny.Array
		_ = _nw92
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw92 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.CodePoint = func(_505_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			}
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw92 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw92).ArraySet1CodePoint(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw92).ArraySet1CodePoint(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_504_v75 = _nw92
		var _506_v76 _dafny.CodePoint
		_ = _506_v76
		_506_v76 = _dafny.CodePoint('e')
		var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_504_v75), 0))
		_ = _index90
		(_504_v75).ArraySet1CodePoint(_506_v76, (_index90).Int())
		var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_504_v75), 0))
		_ = _index91
		(_504_v75).ArraySet1CodePoint(_506_v76, (_index91).Int())
		var _507_i7 _dafny.Int
		_ = _507_i7
		_507_i7 = _dafny.Zero
		{
			for (_this).F9() {
				{
					if (_507_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_507_i7 = (_507_i7).Plus(_dafny.One)
					var _508_v77 _dafny.MultiSet
					_ = _508_v77
					_508_v77 = _dafny.MultiSetOf((_dafny.SetOf(_401_v3)).Cardinality())
					var _509_v78 _dafny.MultiSet
					_ = _509_v78
					_509_v78 = _dafny.MultiSetOf(_508_v77)
					var _510_v79 D0
					_ = _510_v79
					_510_v79 = Companion_D0_.Create_DC2_((_this).F11(), p1, _509_v78)
					var _511_v80 *C0
					_ = _511_v80
					var _nw93 *C0 = New_C0_()
					_ = _nw93
					_nw93.Ctor__(((_this).F9()) && ((_this).F9()), (_dafny.IntOfInt64(-293)).Plus((_510_v79).Dtor_cf3()), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))
					_511_v80 = _nw93
					var _512_v81 _dafny.Array
					_ = _512_v81
					var _nwElement0_13 _dafny.Sequence = p2
					_ = _nwElement0_13
					var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.One)
					_ = _nw94
					(_nw94).ArraySet1(_nwElement0_13, 0)
					_512_v81 = _nw94
					var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_512_v81), 0))
					_ = _index92
					(_512_v81).ArraySet1(Companion_Default___.Fm13((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), true, (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), _511_v80.F14, globalState), (_index92).Int())
					var _513_v82 _dafny.Sequence
					_ = _513_v82
					_513_v82 = _dafny.SeqOf((_this).F9())
					var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_512_v81), 0))
					_ = _index93
					var _rhs78 _dafny.Int = Companion_Default___.Fm3((_504_v75).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_504_v75), 0))).Int()), !((_511_v80).F13()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_513_v82, _513_v82)).Cardinality()), globalState)
					_ = _rhs78
					var _rhs79 _dafny.Int = _401_v3
					_ = _rhs79
					var _rhs80 _dafny.Sequence = p2
					_ = _rhs80
					var _lhs53 *C0 = _511_v80
					_ = _lhs53
					var _lhs54 *C0 = _511_v80
					_ = _lhs54
					var _lhs55 _dafny.Array = _512_v81
					_ = _lhs55
					var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_512_v81), 0))
					_ = _lhs56
					_lhs53.F14 = _rhs78
					_lhs54.F14 = _rhs79
					(_lhs55).ArraySet1(_rhs80, (_lhs56).Int())
					var _514_v83 D1
					_ = _514_v83
					_514_v83 = Companion_D1_.Create_DC6_()
					var _source8 D1 = (func() D1 {
						if (_511_v80).F13() {
							return _514_v83
						}
						return _514_v83
					})()
					_ = _source8
					if _source8.Is_DC5() {
						var _515_v84 _dafny.Map
						_ = _515_v84
						_515_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_399_v2, _511_v80.F14)
						var _516_v85 _dafny.Map
						_ = _516_v85
						_516_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_513_v82).Cardinality()), _511_v80.F14)
						_515_v84 = (_515_v84).Update(_399_v2, ((_516_v85).Merge(_516_v85)).Cardinality())
						var _517_v86 _dafny.Array
						_ = _517_v86
						var _nw95 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
						_ = _nw95
						_517_v86 = _nw95
						var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(480), _dafny.ArrayLen((_517_v86), 0))
						_ = _index94
						(_517_v86).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(950), _511_v80.F14), (_index94).Int())
						var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(480), _dafny.ArrayLen((_517_v86), 0))
						_ = _index95
						(_517_v86).ArraySet1(Companion_Default___.SafeDivisionInt(_511_v80.F14, (func() _dafny.Int {
							if true {
								return _511_v80.F14
							}
							return p1
						})()), (_index95).Int())
						var _518_v87 _dafny.Map
						_ = _518_v87
						_518_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _506_v76)
						var _519_v88 D0
						_ = _519_v88
						_519_v88 = Companion_D0_.Create_DC0_(p1)
						var _520_v89 _dafny.Map
						_ = _520_v89
						_520_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
							if (_518_v87).Contains((_this).F11()) {
								return (_518_v87).Get((_this).F11()).(_dafny.CodePoint)
							}
							return _506_v76
						})(), _519_v88)
						_520_v89 = (_520_v89).Update((_504_v75).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_504_v75), 0))).Int()), _519_v88)
						var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(480), _dafny.ArrayLen((_517_v86), 0))
						_ = _index96
						(_517_v86).ArraySet1(((_this).F12()).Times(p1), (_index96).Int())
					} else if _source8.Is_DC6() {
						var _521_v90 _dafny.Sequence
						_ = _521_v90
						_521_v90 = _dafny.UnicodeSeqOfUtf8Bytes("kpamgbonn")
						var _522_v91 _dafny.Set
						_ = _522_v91
						_522_v91 = _dafny.SetOf((_dafny.MultiSetOf((_511_v80).F13())).IsDisjointFrom(_402_v4), (_511_v80).F13(), (_511_v80).F13())
						var _523_v92 _dafny.Set
						_ = _523_v92
						_523_v92 = _dafny.SetOf((_dafny.Zero).Minus(_511_v80.F14))
						var _524_v93 _dafny.Sequence
						_ = _524_v93
						_524_v93 = _dafny.SeqOf(_523_v92, _523_v92, _523_v92, _523_v92)
						var _525_v94 _dafny.Map
						_ = _525_v94
						_525_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_511_v80).F13(), _522_v91)
						var _rhs81 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_521_v90, _521_v90), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_521_v90, _521_v90)).Cardinality()))).Uint32(), _506_v76), _521_v90)
						_ = _rhs81
						var _rhs82 bool = (_dafny.MultiSetFromSeq((func() _dafny.Sequence {
							if (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool) {
								return _524_v93
							}
							return _524_v93
						})())).IsSubsetOf(_dafny.MultiSetOf(_523_v92, _dafny.SetOf(p1)))
						_ = _rhs82
						var _rhs83 _dafny.Set = ((func() _dafny.Set {
							if (_525_v94).Contains(true) {
								return (_525_v94).Get(true).(_dafny.Set)
							}
							return _dafny.SetOf((_this).F11(), (_this).F9())
						})()).Difference(_522_v91)
						_ = _rhs83
						var _lhs57 *GlobalState = globalState
						_ = _lhs57
						_521_v90 = _rhs81
						_lhs57.F3 = _rhs82
						_522_v91 = _rhs83
						var _526_v95 _dafny.Map
						_ = _526_v95
						_526_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), (_511_v80).F13())
						(globalState).F0 = ((_511_v80).F13()) || ((func() bool {
							if (_526_v95).Contains((_dafny.Zero).Minus(p1)) {
								return (_526_v95).Get((_dafny.Zero).Minus(p1)).(bool)
							}
							return (_this).F9()
						})())
						(globalState).F7 = _506_v76
						var _527_v96 _dafny.Sequence
						_ = _527_v96
						_527_v96 = _dafny.SeqOf(_521_v90, _dafny.UnicodeSeqOfUtf8Bytes("wx"), _521_v90)
						var _528_v97 T0
						_ = _528_v97
						var _nw96 *C0 = New_C0_()
						_ = _nw96
						_nw96.Ctor__(_dafny.Companion_Sequence_.IsProperPrefixOf(_527_v96, _527_v96), p1, (_511_v80).F13())
						_528_v97 = _nw96
						var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_397_v0), 0))
						_ = _index97
						var _rhs84 _dafny.Array = _399_v2
						_ = _rhs84
						var _rhs85 T0 = _528_v97
						_ = _rhs85
						var _rhs86 bool = !(func() _dafny.Map {
							var _coll11 = _dafny.NewMapBuilder()
							_ = _coll11
							for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(736), _dafny.IntOfInt64(-955))); ; {
								_compr_11, _ok11 := _iter11()
								if !_ok11 {
									break
								}
								var _529_v98 _dafny.Int
								_529_v98 = interface{}(_compr_11).(_dafny.Int)
								if ((_dafny.IntOfInt64(736)).Cmp(_529_v98) <= 0) && ((_529_v98).Cmp(_dafny.IntOfInt64(-955)) < 0) {
									_coll11.Add(Companion_Default___.SafeDivisionInt(_529_v98, _511_v80.F14), func() _dafny.Map {
										var _coll12 = _dafny.NewMapBuilder()
										_ = _coll12
										for _iter12 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_((_this).F12()), (func() _dafny.Map {
											var _coll13 = _dafny.NewMapBuilder()
											_ = _coll13
											for _iter13 := _dafny.Iterate(((_512_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_512_v81), 0))).Int()).(_dafny.Sequence)).Elements()); ; {
												_compr_13, _ok13 := _iter13()
												if !_ok13 {
													break
												}
												var _530_v100 _dafny.Int
												_530_v100 = interface{}(_compr_13).(_dafny.Int)
												if _dafny.Companion_Sequence_.Contains((_512_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_512_v81), 0))).Int()).(_dafny.Sequence), _530_v100) {
													_coll13.Add((_530_v100).Plus((_dafny.Zero).Minus(p1)), _511_v80.F14)
												}
											}
											return _coll13.ToMap()
										}()).Cardinality())).Keys().Elements()); ; {
											_compr_12, _ok12 := _iter12()
											if !_ok12 {
												break
											}
											var _531_v99 D0
											_531_v99 = interface{}(_compr_12).(D0)
											if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_((_this).F12()), (func() _dafny.Map {
												var _coll14 = _dafny.NewMapBuilder()
												_ = _coll14
												for _iter14 := _dafny.Iterate(((_512_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_512_v81), 0))).Int()).(_dafny.Sequence)).Elements()); ; {
													_compr_14, _ok14 := _iter14()
													if !_ok14 {
														break
													}
													var _530_v100 _dafny.Int
													_530_v100 = interface{}(_compr_14).(_dafny.Int)
													if _dafny.Companion_Sequence_.Contains((_512_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_512_v81), 0))).Int()).(_dafny.Sequence), _530_v100) {
														_coll14.Add((_530_v100).Plus((_dafny.Zero).Minus(p1)), _511_v80.F14)
													}
												}
												return _coll14.ToMap()
											}()).Cardinality())).Contains(_531_v99) {
												_coll12.Add(_531_v99, _401_v3)
											}
										}
										return _coll12.ToMap()
									}())
								}
							}
							return _coll11.ToMap()
						}()).Contains((_522_v91).Cardinality())
						_ = _rhs86
						var _rhs87 D0 = (_397_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_397_v0), 0))).Int()).(D0)
						_ = _rhs87
						var _rhs88 bool = ((_this).F11()) && ((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))
						_ = _rhs88
						var _lhs58 *GlobalState = globalState
						_ = _lhs58
						var _lhs59 _dafny.Array = _397_v0
						_ = _lhs59
						var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_397_v0), 0))
						_ = _lhs60
						var _lhs61 *GlobalState = globalState
						_ = _lhs61
						_399_v2 = _rhs84
						_528_v97 = _rhs85
						_lhs58.F0 = _rhs86
						(_lhs59).ArraySet1(_rhs87, (_lhs60).Int())
						_lhs61.F0 = _rhs88
					} else {
						var _532___mcc_h10 _dafny.Sequence = _source8.Get_().(D1_DC4).Cf5
						_ = _532___mcc_h10
						var _533_cf5 _dafny.Sequence = _532___mcc_h10
						_ = _533_cf5
						(globalState).F0 = (_this).F9()
						var _534_v101 _dafny.Array
						_ = _534_v101
						var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
						_ = _nw97
						_534_v101 = _nw97
						var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_534_v101), 0))
						_ = _index98
						(_534_v101).ArraySet1(Companion_Default___.SafeModuloInt(p1, _511_v80.F14), (_index98).Int())
						var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_534_v101), 0))
						_ = _index99
						(_534_v101).ArraySet1(_401_v3, (_index99).Int())
						_401_v3 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(148), _401_v3)
						var _535_v102 D0
						_ = _535_v102
						_535_v102 = Companion_D0_.Create_DC1_((_this).Fm8(true, globalState))
						var _536_v103 _dafny.Sequence
						_ = _536_v103
						_536_v103 = _dafny.SeqOf((_512_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_512_v81), 0))).Int()).(_dafny.Sequence))
						var _537_v104 _dafny.Map
						_ = _537_v104
						_537_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_513_v82, _511_v80.F14)
						var _538_v105 _dafny.Set
						_ = _538_v105
						_538_v105 = _dafny.SetOf((_dafny.Zero).Minus((_511_v80.F14).Times(_dafny.IntOfUint32((_536_v103).Cardinality()))), (_537_v104).Cardinality())
						var _pat_let_tv18 = _511_v80
						_ = _pat_let_tv18
						var _rhs89 D0 = func(_pat_let11_0 D0) D0 {
							return func(_539_dt__update__tmp_h0 D0) D0 {
								return func(_pat_let12_0 bool) D0 {
									return func(_540_dt__update_hcf1_h0 bool) D0 {
										return Companion_D0_.Create_DC1_(_540_dt__update_hcf1_h0)
									}(_pat_let12_0)
								}((_pat_let_tv18).F13())
							}(_pat_let11_0)
						}(_535_v102)
						_ = _rhs89
						var _rhs90 _dafny.Set = _538_v105
						_ = _rhs90
						_535_v102 = _rhs89
						_538_v105 = _rhs90
					}
					var _541_v106 _dafny.Sequence
					_ = _541_v106
					_541_v106 = _dafny.SeqOf(_506_v76)
					var _542_v107 D3
					_ = _542_v107
					_542_v107 = Companion_D3_.Create_DC9_(_541_v106)
					var _543_v108 _dafny.Map
					_ = _543_v108
					_543_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _541_v106)
					var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
					_ = _index100
					var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
					_ = _index101
					var _rhs91 bool = (_511_v80.F14).Cmp((_511_v80.F14).Plus(_dafny.IntOfUint32(((_542_v107).Dtor_cf10()).Cardinality()))) > 0
					_ = _rhs91
					var _rhs92 bool = !(_543_v108).Equals(_543_v108)
					_ = _rhs92
					var _rhs93 _dafny.Int = p1
					_ = _rhs93
					var _rhs94 _dafny.Int = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mm")).Cardinality())).Minus(_dafny.IntOfUint32(((_512_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_512_v81), 0))).Int()).(_dafny.Sequence)).Cardinality()))
					_ = _rhs94
					var _lhs62 _dafny.Array = _399_v2
					_ = _lhs62
					var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
					_ = _lhs63
					var _lhs64 _dafny.Array = _399_v2
					_ = _lhs64
					var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))
					_ = _lhs65
					(_lhs62).ArraySet1(_rhs91, (_lhs63).Int())
					(_lhs64).ArraySet1(_rhs92, (_lhs65).Int())
					_401_v3 = _rhs93
					_401_v3 = _rhs94
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _544_v109 _dafny.Map
		_ = _544_v109
		_544_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), p1)
		var _545_v110 _dafny.MultiSet
		_ = _545_v110
		_545_v110 = _dafny.MultiSetOf(_dafny.IntOfInt64(336), (_dafny.Zero).Minus((_this).F12()))
		var _546_v111 _dafny.Sequence
		_ = _546_v111
		_546_v111 = _dafny.SeqOf(_545_v110)
		var _547_v112 _dafny.Array
		_ = _547_v112
		var _nwElement0_14 _dafny.Int = (_544_v109).Cardinality()
		_ = _nwElement0_14
		var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(21))
		_ = _nw98
		(_nw98).ArraySet1(_nwElement0_14, 0)
		(_nw98).ArraySet1(_401_v3, 1)
		(_nw98).ArraySet1(p1, 2)
		(_nw98).ArraySet1((_dafny.Zero).Minus((_this).F12()), 3)
		(_nw98).ArraySet1((_this).F12(), 4)
		(_nw98).ArraySet1((func() _dafny.Int {
			if (_402_v4).Contains((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool)) {
				return (_402_v4).Multiplicity((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))
			}
			return _401_v3
		})(), 5)
		(_nw98).ArraySet1((_this).F12(), 6)
		(_nw98).ArraySet1(_dafny.IntOfInt64(759), 7)
		(_nw98).ArraySet1(_401_v3, 8)
		(_nw98).ArraySet1((_dafny.Zero).Minus(p1), 9)
		(_nw98).ArraySet1(p1, 10)
		(_nw98).ArraySet1(p1, 11)
		(_nw98).ArraySet1(_dafny.IntOfInt64(-591), 12)
		(_nw98).ArraySet1(((_546_v111).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_546_v111).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), 13)
		(_nw98).ArraySet1(_401_v3, 14)
		(_nw98).ArraySet1(p1, 15)
		(_nw98).ArraySet1(_401_v3, 16)
		(_nw98).ArraySet1((_this).F12(), 17)
		(_nw98).ArraySet1(_dafny.IntOfUint32((p2).Cardinality()), 18)
		(_nw98).ArraySet1((_this).F12(), 19)
		(_nw98).ArraySet1((_dafny.Zero).Minus((_this).F12()), 20)
		_547_v112 = _nw98
		var _548_v113 _dafny.Map
		_ = _548_v113
		_548_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_547_v112, (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))
		_548_v113 = (_548_v113).Update(_547_v112, !(!(Companion_Default___.Fm2(globalState)) || ((_this).F11())))
		var _549_v114 D2
		_ = _549_v114
		_549_v114 = Companion_D2_.Create_DC7_(_399_v2)
		var _source9 D2 = _549_v114
		_ = _source9
		if _source9.Is_DC8() {
			var _550___mcc_h11 bool = _source9.Get_().(D2_DC8).Cf7
			_ = _550___mcc_h11
			var _551___mcc_h12 _dafny.Int = _source9.Get_().(D2_DC8).Cf8
			_ = _551___mcc_h12
			var _552___mcc_h13 bool = _source9.Get_().(D2_DC8).Cf9
			_ = _552___mcc_h13
			var _553_cf9 bool = _552___mcc_h13
			_ = _553_cf9
			var _554_cf8 _dafny.Int = _551___mcc_h12
			_ = _554_cf8
			var _555_cf7 bool = _550___mcc_h11
			_ = _555_cf7
			_554_cf8 = (_this).F12()
			var _556_v115 _dafny.Sequence
			_ = _556_v115
			_556_v115 = _dafny.SeqOf(false, (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))
			var _557_v116 _dafny.Map
			_ = _557_v116
			_557_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _556_v115)
			var _558_v117 *C0
			_ = _558_v117
			var _nw99 *C0 = New_C0_()
			_ = _nw99
			_nw99.Ctor__(_553_cf9, Companion_Default___.SafeModuloInt((_this).Fm4(_557_v116, (_this).F9(), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), (_this).F12(), globalState), _554_cf8), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool))
			_558_v117 = _nw99
			_555_cf7 = Companion_Default___.Fm2(globalState)
			var _559_v118 _dafny.Set
			_ = _559_v118
			_559_v118 = _dafny.SetOf((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), (_dafny.IntOfInt64(336)).Cmp(p1) <= 0, (_this).F9())
			_559_v118 = (_559_v118).Union(_dafny.SetOf(!(_553_cf9), (_this).F11(), (_558_v117).F13(), !((_558_v117).F13()), (_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool)))
		} else {
			var _560___mcc_h14 _dafny.Array = _source9.Get_().(D2_DC7).Cf6
			_ = _560___mcc_h14
			var _561_cf6 _dafny.Array = _560___mcc_h14
			_ = _561_cf6
			var _562_v119 *C0
			_ = _562_v119
			var _nw100 *C0 = New_C0_()
			_ = _nw100
			_nw100.Ctor__(false, _dafny.IntOfUint32((Companion_Default___.Fm11((_399_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_399_v2), 0))).Int()).(bool), (_this).F11(), globalState)).Cardinality()), (_this).F9())
			_562_v119 = _nw100
			var _563_v120 _dafny.Sequence
			_ = _563_v120
			_563_v120 = _dafny.UnicodeSeqOfUtf8Bytes("ji")
			var _564_v121 D3
			_ = _564_v121
			_564_v121 = Companion_D3_.Create_DC11_(_562_v119.F14, _563_v120, _562_v119.F14, _562_v119)
			var _565_v122 D3
			_ = _565_v122
			_565_v122 = Companion_D3_.Create_DC11_((_dafny.Zero).Minus(_562_v119.F14), _563_v120, _562_v119.F14, (_564_v121).Dtor_cf17())
			_562_v119 = (_564_v121).Dtor_cf17()
			var _566_v123 _dafny.Array
			_ = _566_v123
			var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
			_ = _nw101
			_566_v123 = _nw101
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_566_v123), 0))
			_ = _index102
			(_566_v123).ArraySet1((Companion_D1_.Create_DC4_(p2)).Dtor_cf5(), (_index102).Int())
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_566_v123), 0))
			_ = _index103
			(_566_v123).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_563_v120).Cardinality())), (_index103).Int())
			_563_v120 = _563_v120
			var _567_v124 _dafny.Sequence
			_ = _567_v124
			_567_v124 = _dafny.SeqOf((_401_v3).Cmp(_dafny.IntOfInt64(282)) != 0, (_this).F11(), (_562_v119).F13())
			(globalState).F3 = (_567_v124).Select((Companion_Default___.SafeIndex((p1).Plus(p1), _dafny.IntOfUint32((_567_v124).Cardinality()))).Uint32()).(bool)
		}
		r0 = (_504_v75).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_504_v75), 0))).Int())
		return r0
	}
}
func (_this *C1) F11() bool {
	{
		return _this._f11
	}
}
func (_this *C1) F12() _dafny.Int {
	{
		return _this._f12
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f9 bool
	F10 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f9 = false
	_this.F10 = false
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

func (_this *C2) F9() bool {
	return _this._f9
}
func (_this *C2) Ctor__(f10 bool, f9 bool) {
	{
		(_this).F10 = f10
		(_this)._f9 = f9
	}
}
func (_this *C2) Fm4(p0 _dafny.Map, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		if (_this).F9() {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.SetOf((_this).F9(), !((_this).F9()))).Cardinality(), _dafny.IntOfInt64(550), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ofopjn")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(983)), (_this).F9())).Cardinality()
		} else {
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ibh"), _dafny.UnicodeSeqOfUtf8Bytes("fwhmbs"))).Cardinality())
		}
	}
}
func (_this *C2) Fm6(p0 bool, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(266)
	}
}
func (_this *C2) Fm7(p0 _dafny.Map, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	{
		var _source10 D0 = Companion_D0_.Create_DC2_((_this).F9(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ekckaebw")).Cardinality()), _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(690))).Uint32(), func(coer27 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}(func(_568_i0 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(268), _dafny.IntOfInt64(682)))
		}))))
		_ = _source10
		if _source10.Is_DC1() {
			var _569___mcc_h0 bool = _source10.Get_().(D0_DC1).Cf1
			_ = _569___mcc_h0
			var _570_cf1 bool = _569___mcc_h0
			_ = _570_cf1
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(558), _this.F10)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-997), (_this).F9()))
		} else if _source10.Is_DC2() {
			var _571___mcc_h1 bool = _source10.Get_().(D0_DC2).Cf2
			_ = _571___mcc_h1
			var _572___mcc_h2 _dafny.Int = _source10.Get_().(D0_DC2).Cf3
			_ = _572___mcc_h2
			var _573___mcc_h3 _dafny.MultiSet = _source10.Get_().(D0_DC2).Cf4
			_ = _573___mcc_h3
			var _574_cf4 _dafny.MultiSet = _573___mcc_h3
			_ = _574_cf4
			var _575_cf3 _dafny.Int = _572___mcc_h2
			_ = _575_cf3
			var _576_cf2 bool = _571___mcc_h1
			_ = _576_cf2
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_575_cf3), _this.F10)
		} else if _source10.Is_DC3() {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_this.F10, (_this).F9(), true)).Cardinality(), _this.F10)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-643)), _this.F10))
		} else {
			var _577___mcc_h4 _dafny.Int = _source10.Get_().(D0_DC0).Cf0
			_ = _577___mcc_h4
			var _578_cf0 _dafny.Int = _577___mcc_h4
			_ = _578_cf0
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_578_cf0, (_this).F9())
		}
	}
}
func (_this *C2) M2(p0 bool, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) (_dafny.Set, bool) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 bool = false
		_ = r1
		var _579_v0 _dafny.Sequence
		_ = _579_v0
		_579_v0 = _dafny.UnicodeSeqOfUtf8Bytes("wqqru")
		_579_v0 = _dafny.UnicodeSeqOfUtf8Bytes("mysvhxlt")
		var _580_v1 _dafny.Sequence
		_ = _580_v1
		_580_v1 = _dafny.SeqOf(Companion_Default___.Fm2(globalState))
		var _581_v2 _dafny.Int
		_ = _581_v2
		_581_v2 = _dafny.IntOfInt64(579)
		var _582_v3 _dafny.Map
		_ = _582_v3
		_582_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_581_v2, _dafny.IntOfInt64(21))
		var _583_v4 _dafny.MultiSet
		_ = _583_v4
		_583_v4 = _dafny.MultiSetOf(_this.F10, (_580_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_580_v1, (Companion_Default___.SafeIndex((_582_v3).Cardinality(), _dafny.IntOfUint32((_580_v1).Cardinality()))).Uint32(), (_this).F9())).Cardinality()), _dafny.IntOfUint32((_580_v1).Cardinality()))).Uint32()).(bool), p0)
		var _584_v5 _dafny.MultiSet
		_ = _584_v5
		_584_v5 = _dafny.MultiSetOf(_581_v2)
		var _585_v6 T0
		_ = _585_v6
		var _nw102 *C0 = New_C0_()
		_ = _nw102
		_nw102.Ctor__(_this.F10, ((func() _dafny.Int {
			if (_583_v4).Contains(p0) {
				return (_583_v4).Multiplicity(p0)
			}
			return _581_v2
		})()).Plus(_581_v2), (_584_v5).IsProperSubsetOf(_584_v5))
		_585_v6 = _nw102
		_585_v6 = (func() T0 {
			if !((_585_v6).F9()) {
				return _585_v6
			}
			return _585_v6
		})()
		(_this).F10 = _this.F10
		var _586_v7 _dafny.Set
		_ = _586_v7
		_586_v7 = _dafny.SetOf((_585_v6).F9())
		_581_v2 = (_585_v6).Fm4(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _580_v1), (_581_v2).Cmp(_581_v2) <= 0, p0, (_586_v7).Cardinality(), globalState)
		var _587_v8 _dafny.Sequence
		_ = _587_v8
		_587_v8 = _dafny.SeqOf(Companion_Default___.Fm11(true, (_585_v6).F9(), globalState))
		(globalState).F0 = (_dafny.IntOfUint32((_587_v8).Cardinality())).Cmp(_dafny.IntOfInt64(748)) < 0
		var _hi9 _dafny.Int = _581_v2
		_ = _hi9
		for _588_i0 := _581_v2; _588_i0.Cmp(_hi9) < 0; _588_i0 = _588_i0.Plus(_dafny.One) {
			(globalState).F0 = !(_this.F10)
			_581_v2 = ((Companion_Default___.Fm3((p2).Select((Companion_Default___.SafeIndex(_581_v2, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(_dafny.CodePoint), p0, _581_v2, globalState)).Plus(_588_i0)).Times(_dafny.IntOfUint32((p2).Cardinality()))
			var _589_v9 _dafny.Array
			_ = _589_v9
			var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw103
			_589_v9 = _nw103
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_589_v9), 0))
			_ = _index104
			(_589_v9).ArraySet1(_588_i0, (_index104).Int())
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_589_v9), 0))
			_ = _index105
			(_589_v9).ArraySet1(_588_i0, (_index105).Int())
			var _590_v10 _dafny.Map
			_ = _590_v10
			_590_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_585_v6, true)
			var _591_v11 _dafny.Array
			_ = _591_v11
			var _nwElement0_15 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_585_v6, p1)
			_ = _nwElement0_15
			var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(6))
			_ = _nw104
			(_nw104).ArraySet1(_nwElement0_15, 0)
			(_nw104).ArraySet1(_590_v10, 1)
			(_nw104).ArraySet1((func() _dafny.Map {
				if true {
					return _590_v10
				}
				return _590_v10
			})(), 2)
			(_nw104).ArraySet1(_590_v10, 3)
			(_nw104).ArraySet1(_590_v10, 4)
			(_nw104).ArraySet1(_590_v10, 5)
			_591_v11 = _nw104
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_591_v11), 0))
			_ = _index106
			(_591_v11).ArraySet1(_590_v10, (_index106).Int())
			var _592_v12 _dafny.Array
			_ = _592_v12
			var _nwElement0_16 bool = !_dafny.Companion_Sequence_.Contains(_580_v1, _this.F10)
			_ = _nwElement0_16
			var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(2))
			_ = _nw105
			(_nw105).ArraySet1(_nwElement0_16, 0)
			(_nw105).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(p2, _579_v0), 1)
			_592_v12 = _nw105
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_592_v12), 0))
			_ = _index107
			(_592_v12).ArraySet1(Companion_Default___.Fm2(globalState), (_index107).Int())
			var _593_v13 _dafny.Sequence
			_ = _593_v13
			_593_v13 = _dafny.SeqOf(_dafny.IntOfInt64(171), _581_v2)
			var _594_v14 _dafny.Map
			_ = _594_v14
			_594_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_593_v13, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(390))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}((func(_595_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_596_i1 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_595_v0).Cardinality())
				}
			})(_579_v0))))
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_591_v11), 0))
			_ = _index108
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_592_v12), 0))
			_ = _index109
			var _rhs95 _dafny.Map = _590_v10
			_ = _rhs95
			var _rhs96 bool = (_585_v6).F9()
			_ = _rhs96
			var _rhs97 _dafny.Sequence = _580_v1
			_ = _rhs97
			var _rhs98 bool = (_this.F10) || ((_585_v6).F9())
			_ = _rhs98
			var _rhs99 _dafny.Int = ((_594_v14).Cardinality()).Minus(_588_i0)
			_ = _rhs99
			var _lhs66 _dafny.Array = _591_v11
			_ = _lhs66
			var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_591_v11), 0))
			_ = _lhs67
			var _lhs68 _dafny.Array = _592_v12
			_ = _lhs68
			var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_592_v12), 0))
			_ = _lhs69
			(_lhs66).ArraySet1(_rhs95, (_lhs67).Int())
			r1 = _rhs96
			_580_v1 = _rhs97
			(_lhs68).ArraySet1(_rhs98, (_lhs69).Int())
			_581_v2 = _rhs99
		}
		r0 = _586_v7
		r1 = (func() bool {
			if _this.F10 {
				return _this.F10
			}
			return (_this).F9()
		})()
		return r0, r1
	}
}
func (_this *C2) M3(p0 _dafny.Int, p1 _dafny.Array, p2 bool, p3 _dafny.Set, globalState *GlobalState) (bool, bool, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var _597_v1 _dafny.Array
		_ = _597_v1
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_8
		var _nw106 _dafny.Array
		_ = _nw106
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw106 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) bool = (func(_598_p0 _dafny.Int, _599_p2 bool) func(_dafny.Int) bool {
				return func(_600_i0 _dafny.Int) bool {
					return (func() _dafny.Set {
						var _coll15 = _dafny.NewBuilder()
						_ = _coll15
						for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(427), _dafny.IntOfInt64(24))); ; {
							_compr_15, _ok15 := _iter15()
							if !_ok15 {
								break
							}
							var _601_v0 _dafny.Int
							_601_v0 = interface{}(_compr_15).(_dafny.Int)
							if ((_dafny.IntOfInt64(427)).Cmp(_601_v0) <= 0) && ((_601_v0).Cmp(_dafny.IntOfInt64(24)) < 0) {
								_coll15.Add((_601_v0).Plus(_dafny.IntOfInt64(786)))
							}
						}
						return _coll15.ToSet()
					}()).IsProperSubsetOf(_dafny.SetOf(_598_p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(407))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg29 _dafny.Int) interface{} {
							return coer29(arg29)
						}
					}(func(_602_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('u')
					}))).Cardinality()), _dafny.IntOfInt64(514), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_599_p2, _this.F10)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-939))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg30 _dafny.Int) interface{} {
							return coer30(arg30)
						}
					}(func(_603_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('p')
					}))).Cardinality())))
				}
			})(p0, p2)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw106 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw106).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw106).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_597_v1 = _nw106
		var _604_v2 _dafny.Map
		_ = _604_v2
		_604_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F9())
		var _605_v3 _dafny.MultiSet
		_ = _605_v3
		_605_v3 = _dafny.MultiSetOf((_604_v2).Cardinality())
		var _606_v4 _dafny.Sequence
		_ = _606_v4
		_606_v4 = _dafny.SeqOf(p0, (_605_v3).Cardinality(), _dafny.IntOfInt64(333), p0)
		var _607_v5 _dafny.Map
		_ = _607_v5
		_607_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_605_v3).Contains((_606_v4).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_606_v4).Cardinality()))).Uint32()).(_dafny.Int)) {
				return (_605_v3).Multiplicity((_606_v4).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_606_v4).Cardinality()))).Uint32()).(_dafny.Int))
			}
			return p0
		})(), true)
		var _608_v6 _dafny.Sequence
		_ = _608_v6
		_608_v6 = _dafny.UnicodeSeqOfUtf8Bytes("isc")
		var _609_v7 _dafny.Map
		_ = _609_v7
		_609_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_607_v5).Cardinality(), _608_v6)
		var _610_v8 _dafny.MultiSet
		_ = _610_v8
		_610_v8 = _dafny.MultiSetOf((func() _dafny.Sequence {
			if _this.F10 {
				return _608_v6
			}
			return _608_v6
		})(), _dafny.UnicodeSeqOfUtf8Bytes("djiib"), _608_v6)
		var _611_v9 _dafny.Map
		_ = _611_v9
		_611_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10, _609_v7)
		var _612_v10 _dafny.Sequence
		_ = _612_v10
		_612_v10 = _dafny.SeqOf(_this.F10, _this.F10)
		var _613_v11 D2
		_ = _613_v11
		_613_v11 = Companion_D2_.Create_DC8_(p2, _dafny.IntOfInt64(618), (_this).F9())
		var _pat_let_tv19 = _610_v8
		_ = _pat_let_tv19
		var _pat_let_tv20 = _610_v8
		_ = _pat_let_tv20
		var _rhs100 _dafny.Array = p1
		_ = _rhs100
		var _rhs101 bool = (_this).F9()
		_ = _rhs101
		var _rhs102 bool = p2
		_ = _rhs102
		var _rhs103 _dafny.Map = (_609_v7).Merge((func() _dafny.Map {
			if (_611_v9).Contains((_612_v10).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_612_v10).Cardinality()))).Uint32()).(bool)) {
				return (_611_v9).Get((_612_v10).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_612_v10).Cardinality()))).Uint32()).(bool)).(_dafny.Map)
			}
			return _609_v7
		})())
		_ = _rhs103
		var _rhs104 _dafny.MultiSet = func(_source11 D2) _dafny.MultiSet {
			if _source11.Is_DC8() {
				var _614___mcc_h0 bool = _source11.Get_().(D2_DC8).Cf7
				_ = _614___mcc_h0
				var _615___mcc_h1 _dafny.Int = _source11.Get_().(D2_DC8).Cf8
				_ = _615___mcc_h1
				var _616___mcc_h2 bool = _source11.Get_().(D2_DC8).Cf9
				_ = _616___mcc_h2
				var _617_cf9 bool = _616___mcc_h2
				_ = _617_cf9
				var _618_cf8 _dafny.Int = _615___mcc_h1
				_ = _618_cf8
				var _619_cf7 bool = _614___mcc_h0
				_ = _619_cf7
				return _pat_let_tv19
			} else {
				var _620___mcc_h3 _dafny.Array = _source11.Get_().(D2_DC7).Cf6
				_ = _620___mcc_h3
				var _621_cf6 _dafny.Array = _620___mcc_h3
				_ = _621_cf6
				return _pat_let_tv20
			}
		}(_613_v11)
		_ = _rhs104
		_597_v1 = _rhs100
		r0 = _rhs101
		r1 = _rhs102
		_609_v7 = _rhs103
		_610_v8 = _rhs104
		var _622_v12 _dafny.Array
		_ = _622_v12
		var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
		_ = _nw107
		_622_v12 = _nw107
		_622_v12 = (Companion_D3_.Create_DC10_(p0, _622_v12, (_612_v10).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_612_v10).Cardinality()))).Uint32()).(bool))).Dtor_cf12()
		var _623_v13 _dafny.Sequence
		_ = _623_v13
		_623_v13 = _dafny.SeqOf(_597_v1, p1, p1, _597_v1, _597_v1)
		var _624_v14 *C0
		_ = _624_v14
		var _nw108 *C0 = New_C0_()
		_ = _nw108
		_nw108.Ctor__(true, _dafny.IntOfUint32((_623_v13).Cardinality()), (_this.F10) || (p2))
		_624_v14 = _nw108
		var _625_v15 _dafny.CodePoint
		_ = _625_v15
		_625_v15 = _dafny.CodePoint('k')
		var _pat_let_tv21 = _624_v14
		_ = _pat_let_tv21
		var _pat_let_tv22 = _624_v14
		_ = _pat_let_tv22
		var _pat_let_tv23 = _624_v14
		_ = _pat_let_tv23
		var _pat_let_tv24 = _624_v14
		_ = _pat_let_tv24
		(_this).F10 = func(_source12 D3) bool {
			if _source12.Is_DC10() {
				var _626___mcc_h4 _dafny.Int = _source12.Get_().(D3_DC10).Cf11
				_ = _626___mcc_h4
				var _627___mcc_h5 _dafny.Array = _source12.Get_().(D3_DC10).Cf12
				_ = _627___mcc_h5
				var _628___mcc_h6 bool = _source12.Get_().(D3_DC10).Cf13
				_ = _628___mcc_h6
				var _629_cf13 bool = _628___mcc_h6
				_ = _629_cf13
				var _630_cf12 _dafny.Array = _627___mcc_h5
				_ = _630_cf12
				var _631_cf11 _dafny.Int = _626___mcc_h4
				_ = _631_cf11
				return true
			} else if _source12.Is_DC11() {
				var _632___mcc_h7 _dafny.Int = _source12.Get_().(D3_DC11).Cf14
				_ = _632___mcc_h7
				var _633___mcc_h8 _dafny.Sequence = _source12.Get_().(D3_DC11).Cf15
				_ = _633___mcc_h8
				var _634___mcc_h9 _dafny.Int = _source12.Get_().(D3_DC11).Cf16
				_ = _634___mcc_h9
				var _635___mcc_h10 *C0 = _source12.Get_().(D3_DC11).Cf17
				_ = _635___mcc_h10
				var _636_cf17 *C0 = _635___mcc_h10
				_ = _636_cf17
				var _637_cf16 _dafny.Int = _634___mcc_h9
				_ = _637_cf16
				var _638_cf15 _dafny.Sequence = _633___mcc_h8
				_ = _638_cf15
				var _639_cf14 _dafny.Int = _632___mcc_h7
				_ = _639_cf14
				return (_pat_let_tv21).F13()
			} else if _source12.Is_DC9() {
				var _640___mcc_h11 _dafny.Sequence = _source12.Get_().(D3_DC9).Cf10
				_ = _640___mcc_h11
				var _641_cf10 _dafny.Sequence = _640___mcc_h11
				_ = _641_cf10
				return (_this).F9()
			} else {
				var _642___mcc_h12 D3 = _source12.Get_().(D3_DC12).Cf18
				_ = _642___mcc_h12
				var _643_cf18 D3 = _642___mcc_h12
				_ = _643_cf18
				return !((_dafny.SetOf(_pat_let_tv22.F14)).IsProperSubsetOf(_dafny.SetOf(_pat_let_tv23.F14, (_dafny.Zero).Minus(_pat_let_tv24.F14))))
			}
		}(Companion_D3_.Create_DC9_(_dafny.SeqOf(_625_v15, _625_v15)))
		var _644_v16 _dafny.Array
		_ = _644_v16
		var _nwElement0_17 _dafny.Array = p1
		_ = _nwElement0_17
		var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(15))
		_ = _nw109
		(_nw109).ArraySet1(_nwElement0_17, 0)
		(_nw109).ArraySet1(p1, 1)
		(_nw109).ArraySet1(_597_v1, 2)
		(_nw109).ArraySet1(p1, 3)
		(_nw109).ArraySet1(p1, 4)
		(_nw109).ArraySet1(p1, 5)
		(_nw109).ArraySet1(p1, 6)
		(_nw109).ArraySet1(p1, 7)
		(_nw109).ArraySet1(p1, 8)
		(_nw109).ArraySet1(p1, 9)
		(_nw109).ArraySet1(p1, 10)
		(_nw109).ArraySet1(p1, 11)
		(_nw109).ArraySet1(p1, 12)
		(_nw109).ArraySet1(p1, 13)
		(_nw109).ArraySet1(_597_v1, 14)
		_644_v16 = _nw109
		var _645_v17 D4
		_ = _645_v17
		_645_v17 = Companion_D4_.Create_DC13_(_644_v16)
		var _646_v18 _dafny.Array
		_ = _646_v18
		var _nwElement0_18 _dafny.Array = _644_v16
		_ = _nwElement0_18
		var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(28))
		_ = _nw110
		(_nw110).ArraySet1(_nwElement0_18, 0)
		(_nw110).ArraySet1(_644_v16, 1)
		(_nw110).ArraySet1(_644_v16, 2)
		(_nw110).ArraySet1(_644_v16, 3)
		(_nw110).ArraySet1(_644_v16, 4)
		(_nw110).ArraySet1(_644_v16, 5)
		(_nw110).ArraySet1(_644_v16, 6)
		(_nw110).ArraySet1(_644_v16, 7)
		(_nw110).ArraySet1(_644_v16, 8)
		(_nw110).ArraySet1(_644_v16, 9)
		(_nw110).ArraySet1((_645_v17).Dtor_cf19(), 10)
		(_nw110).ArraySet1(_644_v16, 11)
		(_nw110).ArraySet1(_644_v16, 12)
		(_nw110).ArraySet1(_644_v16, 13)
		(_nw110).ArraySet1(_644_v16, 14)
		(_nw110).ArraySet1(_644_v16, 15)
		(_nw110).ArraySet1(_644_v16, 16)
		(_nw110).ArraySet1(_644_v16, 17)
		(_nw110).ArraySet1(_644_v16, 18)
		(_nw110).ArraySet1(_644_v16, 19)
		(_nw110).ArraySet1(_644_v16, 20)
		(_nw110).ArraySet1(_644_v16, 21)
		(_nw110).ArraySet1(_644_v16, 22)
		(_nw110).ArraySet1(_644_v16, 23)
		(_nw110).ArraySet1(_644_v16, 24)
		(_nw110).ArraySet1(_644_v16, 25)
		(_nw110).ArraySet1((_645_v17).Dtor_cf19(), 26)
		(_nw110).ArraySet1(_644_v16, 27)
		_646_v18 = _nw110
		var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_646_v18), 0))
		_ = _index110
		(_646_v18).ArraySet1(_644_v16, (_index110).Int())
		var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_646_v18), 0))
		_ = _index111
		(_646_v18).ArraySet1(_644_v16, (_index111).Int())
		if (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_608_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(919))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}((func(_647_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_648_i3 _dafny.Int) _dafny.CodePoint {
				return _647_v15
			}
		})(_625_v15))))).Cardinality())).Cmp(_dafny.IntOfUint32((Companion_Default___.Fm10(globalState)).Cardinality())) < 0 {
			(globalState).F7 = _625_v15
			var _649_v19 _dafny.MultiSet
			_ = _649_v19
			_649_v19 = _dafny.MultiSetOf(p2)
			var _650_v20 _dafny.Map
			_ = _650_v20
			_650_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_624_v14, _649_v19)
			var _651_v21 _dafny.Map
			_ = _651_v21
			_651_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_650_v20, Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(66), _dafny.IntOfInt64(853)))
			_651_v21 = (_651_v21).Update(_650_v20, Companion_Default___.SafeDivisionInt(_624_v14.F14, _624_v14.F14))
			var _652_v22 _dafny.Array
			_ = _652_v22
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_9
			var _nw111 _dafny.Array
			_ = _nw111
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw111 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Map = (func(_653_v2 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_654_i4 _dafny.Int) _dafny.Map {
						return _653_v2
					}
				})(_604_v2)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw111 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw111).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw111).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_652_v22 = _nw111
			_652_v22 = _652_v22
			var _655_v23 _dafny.Array
			_ = _655_v23
			var _nw112 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(9))
			_ = _nw112
			_655_v23 = _nw112
			var _656_v24 _dafny.MultiSet
			_ = _656_v24
			_656_v24 = _dafny.MultiSetOf(_597_v1)
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_655_v23), 0))
			_ = _index112
			(_655_v23).ArraySet1((_656_v24).Difference(_dafny.MultiSetOf(p1)), (_index112).Int())
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_655_v23), 0))
			_ = _index113
			(_655_v23).ArraySet1(_656_v24, (_index113).Int())
			var _657_v25 D1
			_ = _657_v25
			_657_v25 = Companion_D1_.Create_DC4_(_606_v4)
			var _source13 D1 = _657_v25
			_ = _source13
			if _source13.Is_DC5() {
				(_624_v14).F14 = _dafny.IntOfInt64(421)
				var _658_v26 _dafny.Sequence
				_ = _658_v26
				_658_v26 = _dafny.SeqOf(_604_v2)
				var _659_v27 _dafny.Map
				_ = _659_v27
				_659_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_622_v12, Companion_Default___.Fm3(_625_v15, true, p0, globalState))
				var _rhs105 _dafny.Int = _624_v14.F14
				_ = _rhs105
				var _rhs106 _dafny.Sequence = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm14(_dafny.IntOfInt64(-443), globalState), (Companion_Default___.SafeIndex((_dafny.IntOfUint32((_dafny.SeqOf((_624_v14).F13())).Cardinality())).Minus(_624_v14.F14), _dafny.IntOfUint32((Companion_Default___.Fm14(_dafny.IntOfInt64(-443), globalState)).Cardinality()))).Uint32(), _604_v2)
				_ = _rhs106
				var _rhs107 bool = !(_659_v27).Equals((_659_v27).Merge(_659_v27))
				_ = _rhs107
				var _rhs108 _dafny.CodePoint = _625_v15
				_ = _rhs108
				var _rhs109 _dafny.Array = p1
				_ = _rhs109
				var _lhs70 *C0 = _624_v14
				_ = _lhs70
				var _lhs71 *GlobalState = globalState
				_ = _lhs71
				_lhs70.F14 = _rhs105
				_658_v26 = _rhs106
				r0 = _rhs107
				_lhs71.F7 = _rhs108
				_597_v1 = _rhs109
				var _660_v28 *C1
				_ = _660_v28
				var _nw113 *C1 = New_C1_()
				_ = _nw113
				_nw113.Ctor__((_624_v14).F13(), (_this).Fm6((_624_v14).F13(), (_this).F9(), globalState), p2)
				_660_v28 = _nw113
				_660_v28 = _660_v28
				var _661_v29 _dafny.Array
				_ = _661_v29
				var _nw114 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(29))
				_ = _nw114
				_661_v29 = _nw114
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_661_v29), 0))
				_ = _index114
				(_661_v29).ArraySet1(_613_v11, (_index114).Int())
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_661_v29), 0))
				_ = _index115
				(_661_v29).ArraySet1(Companion_D2_.Create_DC8_(false, (_660_v28).F12(), (_624_v14).F13()), (_index115).Int())
			} else if _source13.Is_DC6() {
				(_this).F10 = (func() bool {
					if _this.F10 {
						return !(p2)
					}
					return (_624_v14).F13()
				})()
				var _662_v30 _dafny.Sequence
				_ = _662_v30
				_662_v30 = _dafny.SeqOf(_608_v6, _608_v6)
				var _663_v31 _dafny.Sequence
				_ = _663_v31
				_663_v31 = _dafny.SeqOf(_610_v8)
				var _664_v32 _dafny.Array
				_ = _664_v32
				var _nwElement0_19 _dafny.MultiSet = _610_v8
				_ = _nwElement0_19
				var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(9))
				_ = _nw115
				(_nw115).ArraySet1(_nwElement0_19, 0)
				(_nw115).ArraySet1((_610_v8).Update(_608_v6, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gnmnf")).Cardinality()))), 1)
				(_nw115).ArraySet1(_610_v8, 2)
				(_nw115).ArraySet1((_dafny.MultiSetFromSeq(_662_v30)).Intersection(_610_v8), 3)
				(_nw115).ArraySet1((_610_v8).Union(_610_v8), 4)
				(_nw115).ArraySet1((_610_v8).Difference(_610_v8), 5)
				(_nw115).ArraySet1(_dafny.MultiSetFromSeq(Companion_Default___.Fm15((_this).F9(), globalState)), 6)
				(_nw115).ArraySet1((_dafny.MultiSetOf(_608_v6)).Intersection((_663_v31).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_663_v31).Cardinality()))).Uint32()).(_dafny.MultiSet)), 7)
				(_nw115).ArraySet1((_610_v8).Union(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("aonb"))), 8)
				_664_v32 = _nw115
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_664_v32), 0))
				_ = _index116
				(_664_v32).ArraySet1(_610_v8, (_index116).Int())
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_597_v1), 0))
				_ = _index117
				(_597_v1).ArraySet1(((_dafny.Zero).Minus(_624_v14.F14)).Cmp((_607_v5).Cardinality()) > 0, (_index117).Int())
				var _665_v33 _dafny.Set
				_ = _665_v33
				_665_v33 = _dafny.SetOf(_624_v14.F14, _dafny.IntOfInt64(141))
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_664_v32), 0))
				_ = _index118
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_597_v1), 0))
				_ = _index119
				var _rhs110 _dafny.MultiSet = _dafny.MultiSetOf((_665_v33).IsDisjointFrom(_665_v33))
				_ = _rhs110
				var _rhs111 bool = !(p2)
				_ = _rhs111
				var _rhs112 _dafny.Int = p0
				_ = _rhs112
				var _rhs113 _dafny.MultiSet = _610_v8
				_ = _rhs113
				var _rhs114 bool = (_624_v14).F13()
				_ = _rhs114
				var _lhs72 *GlobalState = globalState
				_ = _lhs72
				var _lhs73 *C0 = _624_v14
				_ = _lhs73
				var _lhs74 _dafny.Array = _664_v32
				_ = _lhs74
				var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_664_v32), 0))
				_ = _lhs75
				var _lhs76 _dafny.Array = _597_v1
				_ = _lhs76
				var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_597_v1), 0))
				_ = _lhs77
				_649_v19 = _rhs110
				_lhs72.F0 = _rhs111
				_lhs73.F14 = _rhs112
				(_lhs74).ArraySet1(_rhs113, (_lhs75).Int())
				(_lhs76).ArraySet1(_rhs114, (_lhs77).Int())
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_622_v12), 0))
				_ = _index120
				(_622_v12).ArraySet1(p0, (_index120).Int())
				var _666_v34 _dafny.MultiSet
				_ = _666_v34
				_666_v34 = _dafny.MultiSetOf(_605_v3)
				var _667_v35 T0
				_ = _667_v35
				var _nw116 *C1 = New_C1_()
				_ = _nw116
				_nw116.Ctor__(_this.F10, Companion_Default___.Fm3(_625_v15, (_624_v14).F13(), _dafny.IntOfInt64(-346), globalState), (_666_v34).IsProperSubsetOf((_666_v34).Update(_605_v3, Companion_Default___.Abs(_dafny.IntOfInt64(-782)))))
				_667_v35 = _nw116
				var _668_v36 _dafny.Map
				_ = _668_v36
				_668_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_649_v19, _606_v4)
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_622_v12), 0))
				_ = _index121
				var _rhs115 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_668_v36).Contains(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
						if _this.F10 {
							return _612_v10
						}
						return _612_v10
					})())) {
						return (_668_v36).Get(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
							if _this.F10 {
								return _612_v10
							}
							return _612_v10
						})())).(_dafny.Sequence)
					}
					return _dafny.Companion_Sequence_.Concatenate((_657_v25).Dtor_cf5(), _606_v4)
				})()).Cardinality())
				_ = _rhs115
				var _rhs116 bool = ((_624_v14).F13()) && ((_597_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_597_v1), 0))).Int()).(bool))
				_ = _rhs116
				var _rhs117 _dafny.Int = _dafny.IntOfInt64(243)
				_ = _rhs117
				var _rhs118 T0 = _667_v35
				_ = _rhs118
				var _rhs119 bool = !((_624_v14.F14).Cmp(Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfInt64(858))) < 0)
				_ = _rhs119
				var _lhs78 *C0 = _624_v14
				_ = _lhs78
				var _lhs79 _dafny.Array = _622_v12
				_ = _lhs79
				var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_622_v12), 0))
				_ = _lhs80
				var _lhs81 *GlobalState = globalState
				_ = _lhs81
				_lhs78.F14 = _rhs115
				r0 = _rhs116
				(_lhs79).ArraySet1(_rhs117, (_lhs80).Int())
				_667_v35 = _rhs118
				_lhs81.F0 = _rhs119
				var _669_v37 _dafny.Sequence
				_ = _669_v37
				_669_v37 = _dafny.SeqOf(_606_v4, _606_v4)
				var _670_v38 _dafny.Map
				_ = _670_v38
				_670_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F10) || (false), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_669_v37, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_669_v37).Cardinality()))).Uint32(), _606_v4), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_669_v37, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()), _dafny.IntOfUint32((_669_v37).Cardinality()))).Uint32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(184))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}(func(_671_i5 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yqssgr")).Cardinality())
				}))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_606_v4).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_669_v37, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()), _dafny.IntOfUint32((_669_v37).Cardinality()))).Uint32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(184))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}(func(_672_i5 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yqssgr")).Cardinality())
				})))).Cardinality()))).Uint32(), _606_v4)))
				_670_v38 = (_670_v38).Update(!(((_597_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_597_v1), 0))).Int()).(bool)) && ((_667_v35).F9())), _669_v37)
			} else {
				var _673___mcc_h13 _dafny.Sequence = _source13.Get_().(D1_DC4).Cf5
				_ = _673___mcc_h13
				var _674_cf5 _dafny.Sequence = _673___mcc_h13
				_ = _674_cf5
				var _675_v39 _dafny.Map
				_ = _675_v39
				_675_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.SeqOf((_624_v14).F13()))
				(_624_v14).F14 = Companion_Default___.SafeModuloInt((_624_v14.F14).Minus(p0), (_this).Fm4(_675_v39, (_624_v14).F13(), _this.F10, _624_v14.F14, globalState))
				_597_v1 = _597_v1
				var _676_v40 _dafny.Set
				_ = _676_v40
				_676_v40 = _dafny.SetOf(p1, _597_v1)
				_676_v40 = _676_v40
				(_624_v14).F14 = (_dafny.Zero).Minus(p0)
			}
		} else {
			_607_v5 = (_607_v5).Update(p0, (_624_v14).F13())
			var _677_v41 _dafny.Sequence
			_ = _677_v41
			_677_v41 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(globalState), (_624_v14).F13()))
			(globalState).F0 = (func() bool {
				if !(false) {
					return true
				}
				return _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_624_v14).F13()), (_612_v10).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_612_v10).Cardinality()))).Uint32()).(bool)), _604_v2, _604_v2), _677_v41)
			})()
			var _678_v42 *C0
			_ = _678_v42
			var _nw117 *C0 = New_C0_()
			_ = _nw117
			_nw117.Ctor__((_624_v14).F13(), p0, true)
			_678_v42 = _nw117
			var _679_v43 _dafny.Map
			_ = _679_v43
			_679_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_624_v14.F14, _606_v4)
			var _680_v44 _dafny.MultiSet
			_ = _680_v44
			_680_v44 = _dafny.MultiSetOf(_this.F10, (_624_v14).F13(), _this.F10, !(_this.F10))
			_679_v43 = (_679_v43).Update(((func() _dafny.Int {
				if (_680_v44).Contains((_624_v14).F13()) {
					return (_680_v44).Multiplicity((_624_v14).F13())
				}
				return (_dafny.Zero).Minus(_624_v14.F14)
			})()).Times(Companion_Default___.Fm3(_dafny.CodePoint('e'), false, _624_v14.F14, globalState)), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm13((_this).F9(), (_678_v42).F13(), (_624_v14).F13(), _dafny.IntOfUint32((_608_v6).Cardinality()), globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-594))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}((func(_681_v42 *C0) func(_dafny.Int) _dafny.Int {
				return func(_682_i6 _dafny.Int) _dafny.Int {
					return _681_v42.F14
				}
			})(_678_v42)))))
			(globalState).F7 = _625_v15
		}
		r0 = (_624_v14).F13()
		var _683_v45 D0
		_ = _683_v45
		_683_v45 = Companion_D0_.Create_DC0_((_dafny.Zero).Minus(p0))
		var _pat_let_tv25 = _612_v10
		_ = _pat_let_tv25
		var _pat_let_tv26 = _624_v14
		_ = _pat_let_tv26
		var _pat_let_tv27 = _612_v10
		_ = _pat_let_tv27
		var _pat_let_tv28 = _624_v14
		_ = _pat_let_tv28
		var _pat_let_tv29 = p2
		_ = _pat_let_tv29
		r1 = func(_source14 D0) bool {
			if _source14.Is_DC1() {
				var _684___mcc_h14 bool = _source14.Get_().(D0_DC1).Cf1
				_ = _684___mcc_h14
				var _685_cf1 bool = _684___mcc_h14
				_ = _685_cf1
				return (_pat_let_tv25).Select((Companion_Default___.SafeIndex(_pat_let_tv26.F14, _dafny.IntOfUint32((_pat_let_tv27).Cardinality()))).Uint32()).(bool)
			} else if _source14.Is_DC2() {
				var _686___mcc_h15 bool = _source14.Get_().(D0_DC2).Cf2
				_ = _686___mcc_h15
				var _687___mcc_h16 _dafny.Int = _source14.Get_().(D0_DC2).Cf3
				_ = _687___mcc_h16
				var _688___mcc_h17 _dafny.MultiSet = _source14.Get_().(D0_DC2).Cf4
				_ = _688___mcc_h17
				var _689_cf4 _dafny.MultiSet = _688___mcc_h17
				_ = _689_cf4
				var _690_cf3 _dafny.Int = _687___mcc_h16
				_ = _690_cf3
				var _691_cf2 bool = _686___mcc_h15
				_ = _691_cf2
				return (_pat_let_tv28.F14).Cmp(_690_cf3) >= 0
			} else if _source14.Is_DC3() {
				return _this.F10
			} else {
				var _692___mcc_h18 _dafny.Int = _source14.Get_().(D0_DC0).Cf0
				_ = _692___mcc_h18
				var _693_cf0 _dafny.Int = _692___mcc_h18
				_ = _693_cf0
				return _pat_let_tv29
			}
		}(_683_v45)
		var _694_v46 _dafny.Map
		_ = _694_v46
		_694_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _622_v12)
		r2 = (_694_v46).Update((_dafny.Zero).Minus(_624_v14.F14), _622_v12)
		return r0, r1, r2
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f9 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f9 = false
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

func (_this *C3) F9() bool {
	return _this._f9
}
func (_this *C3) Ctor__(f9 bool) {
	{
		(_this)._f9 = f9
	}
}
func (_this *C3) Fm4(p0 _dafny.Map, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ouhfhsh")).Cardinality()), _dafny.IntOfInt64(-999))).Cardinality())).Minus((_dafny.IntOfInt64(636)).Times(_dafny.IntOfInt64(-559)))
	}
}
func (_this *C3) Fm5(globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('a')
	}
}
func (_this *C3) M1(p0 _dafny.Set, globalState *GlobalState) (_dafny.Int, D0, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 D0 = Companion_D0_.Default()
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _695_v0 _dafny.Int
		_ = _695_v0
		_695_v0 = _dafny.IntOfInt64(-2)
		var _696_v1 _dafny.Map
		_ = _696_v1
		_696_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_695_v0).Minus(_dafny.IntOfInt64(757)))
		var _697_v2 _dafny.Array
		_ = _697_v2
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_10
		var _nw118 _dafny.Array
		_ = _nw118
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw118 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.Int = (func(_698_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_699_i0 _dafny.Int) _dafny.Int {
					return (_699_i0).Minus(_698_v0)
				}
			})(_695_v0)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw118 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw118).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw118).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_697_v2 = _nw118
		var _700_v3 _dafny.Set
		_ = _700_v3
		_700_v3 = _dafny.SetOf(_697_v2, _697_v2, _697_v2, _697_v2)
		_696_v1 = (_696_v1).Update((_700_v3).IsProperSubsetOf(_dafny.SetOf(_697_v2, _697_v2, _697_v2)), _695_v0)
		var _701_v4 _dafny.Sequence
		_ = _701_v4
		_701_v4 = _dafny.SeqOf(_695_v0, _695_v0, _695_v0, _695_v0, _695_v0)
		var _702_v5 _dafny.Sequence
		_ = _702_v5
		_702_v5 = _dafny.UnicodeSeqOfUtf8Bytes("ps")
		var _703_v6 _dafny.MultiSet
		_ = _703_v6
		_703_v6 = _dafny.MultiSetOf((_dafny.Zero).Minus(_695_v0), _695_v0)
		var _704_v7 _dafny.MultiSet
		_ = _704_v7
		_704_v7 = _dafny.MultiSetOf(_dafny.MultiSetOf(_695_v0), _dafny.MultiSetOf(_dafny.IntOfUint32((_702_v5).Cardinality()), _695_v0), _703_v6)
		var _705_v8 _dafny.Map
		_ = _705_v8
		_705_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_701_v4, (Companion_D0_.Create_DC2_((_this).F9(), _695_v0, _704_v7)).Dtor_cf3())
		var _706_v9 _dafny.MultiSet
		_ = _706_v9
		_706_v9 = _dafny.MultiSetOf((_705_v8).Cardinality())
		var _707_v10 _dafny.Array
		_ = _707_v10
		var _nwElement0_20 _dafny.MultiSet = _703_v6
		_ = _nwElement0_20
		var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(6))
		_ = _nw119
		(_nw119).ArraySet1(_nwElement0_20, 0)
		(_nw119).ArraySet1(_706_v9, 1)
		(_nw119).ArraySet1(_dafny.MultiSetOf(_695_v0), 2)
		(_nw119).ArraySet1(_703_v6, 3)
		(_nw119).ArraySet1(_dafny.MultiSetFromSeq(_701_v4), 4)
		(_nw119).ArraySet1((_706_v9).Difference(_703_v6), 5)
		_707_v10 = _nw119
		var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_707_v10), 0))
		_ = _index122
		(_707_v10).ArraySet1((func() _dafny.MultiSet {
			if (_this).F9() {
				return _dafny.MultiSetOf((_dafny.Zero).Minus(_695_v0), _695_v0)
			}
			return _703_v6
		})(), (_index122).Int())
		var _708_v11 _dafny.Map
		_ = _708_v11
		_708_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_695_v0, _706_v9)
		var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_707_v10), 0))
		_ = _index123
		(_707_v10).ArraySet1((_706_v9).Intersection((func() _dafny.MultiSet {
			if (_708_v11).Contains(_695_v0) {
				return (_708_v11).Get(_695_v0).(_dafny.MultiSet)
			}
			return _706_v9
		})()), (_index123).Int())
		var _709_i1 _dafny.Int
		_ = _709_i1
		_709_i1 = _dafny.Zero
		{
			for (_this).F9() {
				{
					if (_709_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_709_i1 = (_709_i1).Plus(_dafny.One)
					(globalState).F3 = Companion_Default___.Fm2(globalState)
					var _710_v12 *C1
					_ = _710_v12
					var _nw120 *C1 = New_C1_()
					_ = _nw120
					_nw120.Ctor__((_this).F9(), _695_v0, (_this).F9())
					_710_v12 = _nw120
					var _711_v13 _dafny.CodePoint
					_ = _711_v13
					_711_v13 = _dafny.CodePoint('g')
					var _712_v14 T0
					_ = _712_v14
					var _nw121 *C0 = New_C0_()
					_ = _nw121
					_nw121.Ctor__((_this).F9(), _695_v0, (_this).F9())
					_712_v14 = _nw121
					var _713_v15 _dafny.Map
					_ = _713_v15
					_713_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_712_v14, (_710_v12).F11())
					var _714_v16 *C1
					_ = _714_v16
					var _nw122 *C1 = New_C1_()
					_ = _nw122
					_nw122.Ctor__((_dafny.MultiSetOf(_695_v0)).Contains(Companion_Default___.Fm3(_711_v13, (_this).F9(), (_710_v12).F12(), globalState)), (_dafny.MultiSetOf(_695_v0, (_710_v12).F12())).Cardinality(), (func() bool {
						if (_713_v15).Contains(_712_v14) {
							return (_713_v15).Get(_712_v14).(bool)
						}
						return (_this).F9()
					})())
					_714_v16 = _nw122
					var _715_v17 _dafny.Map
					_ = _715_v17
					_715_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_710_v12).F12(), _710_v12)
					var _716_v18 _dafny.Sequence
					_ = _716_v18
					_716_v18 = _dafny.SeqOf(_710_v12, _714_v16, _710_v12)
					_714_v16 = (func() *C1 {
						if (_715_v17).Contains((_714_v16).F12()) {
							return (_715_v17).Get((_714_v16).F12()).(*C1)
						}
						return (_716_v18).Select((Companion_Default___.SafeIndex((_710_v12).F12(), _dafny.IntOfUint32((_716_v18).Cardinality()))).Uint32()).(*C1)
					})()
					var _717_v19 _dafny.MultiSet
					_ = _717_v19
					_717_v19 = _dafny.MultiSetOf(!((_710_v12).F11()), (_this).F9(), (_this).F9(), (_712_v14).F9())
					var _718_v20 *C2
					_ = _718_v20
					var _nw123 *C2 = New_C2_()
					_ = _nw123
					_nw123.Ctor__((_717_v19).IsDisjointFrom(_717_v19), !_dafny.Companion_Sequence_.Contains(_702_v5, _711_v13))
					_718_v20 = _nw123
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _hi10 _dafny.Int = _695_v0
		_ = _hi10
		for _719_i2 := _695_v0; _719_i2.Cmp(_hi10) < 0; _719_i2 = _719_i2.Plus(_dafny.One) {
			r0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_702_v5, _702_v5)).Cardinality())
			var _720_v21 _dafny.Map
			_ = _720_v21
			_720_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _702_v5)
			_720_v21 = (_720_v21).Update(!(Companion_Default___.Fm2(globalState)) || ((_this).F9()), _dafny.Companion_Sequence_.Concatenate(_702_v5, _702_v5))
			var _721_v22 _dafny.Map
			_ = _721_v22
			_721_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_702_v5, !(false))
			var _722_v23 _dafny.Set
			_ = _722_v23
			_722_v23 = _dafny.SetOf((_this).F9(), (_this).F9(), (_this).F9(), (_this).F9(), (_this).F9())
			var _723_v24 _dafny.Map
			_ = _723_v24
			_723_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_721_v22, (_722_v23).Intersection(_dafny.SetOf((_this).F9(), (_this).F9(), (_this).F9(), (_this).F9(), !((_this).F9()))))
			_723_v24 = (_723_v24).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_702_v5, (_this).F9()), (_722_v23).Union(_722_v23))
			(globalState).F3 = (_this).F9()
		}
		var _724_v25 _dafny.CodePoint
		_ = _724_v25
		_724_v25 = _dafny.CodePoint('t')
		var _725_v26 _dafny.Map
		_ = _725_v26
		_725_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _724_v25)
		var _726_v27 _dafny.Map
		_ = _726_v27
		_726_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_695_v0, (_725_v26).Cardinality())
		_726_v27 = (_726_v27).Update(_695_v0, Companion_Default___.SafeDivisionInt((_703_v6).Cardinality(), _695_v0))
		if (_this).F9() {
			var _727_v28 _dafny.Array
			_ = _727_v28
			var _nwElement0_21 _dafny.Array = _697_v2
			_ = _nwElement0_21
			var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(3))
			_ = _nw124
			(_nw124).ArraySet1(_nwElement0_21, 0)
			(_nw124).ArraySet1(_697_v2, 1)
			(_nw124).ArraySet1(_697_v2, 2)
			_727_v28 = _nw124
			_727_v28 = _727_v28
			var _728_v29 _dafny.Sequence
			_ = _728_v29
			_728_v29 = _dafny.SeqOf(_696_v1)
			var _729_v30 D2
			_ = _729_v30
			_729_v30 = Companion_D2_.Create_DC8_((_this).F9(), (func() _dafny.Int {
				if (_this).F9() {
					return _695_v0
				}
				return (_726_v27).Cardinality()
			})(), _dafny.Companion_Sequence_.IsProperPrefixOf(_728_v29, _728_v29))
			var _source15 D2 = _729_v30
			_ = _source15
			if _source15.Is_DC8() {
				var _730___mcc_h0 bool = _source15.Get_().(D2_DC8).Cf7
				_ = _730___mcc_h0
				var _731___mcc_h1 _dafny.Int = _source15.Get_().(D2_DC8).Cf8
				_ = _731___mcc_h1
				var _732___mcc_h2 bool = _source15.Get_().(D2_DC8).Cf9
				_ = _732___mcc_h2
				var _733_cf9 bool = _732___mcc_h2
				_ = _733_cf9
				var _734_cf8 _dafny.Int = _731___mcc_h1
				_ = _734_cf8
				var _735_cf7 bool = _730___mcc_h0
				_ = _735_cf7
				(globalState).F3 = (_this).F9()
				var _736_v31 _dafny.MultiSet
				_ = _736_v31
				_736_v31 = _dafny.MultiSetOf(!(true), _733_cf9)
				(globalState).F3 = (Companion_Default___.SafeDivisionInt(_695_v0, (_736_v31).Cardinality())).Cmp((_dafny.Zero).Minus((_734_cf8).Times(_695_v0))) <= 0
				(globalState).F0 = (_this).F9()
				var _737_v32 _dafny.Map
				_ = _737_v32
				_737_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_724_v25, _dafny.UnicodeSeqOfUtf8Bytes("uathenfi"))
				var _738_v33 _dafny.Int
				_ = _738_v33
				var _out19 _dafny.Int
				_ = _out19
				_out19 = Companion_Default___.M0((func() _dafny.Int {
					if false {
						return _734_cf8
					}
					return _dafny.IntOfUint32((_701_v4).Cardinality())
				})(), (func() _dafny.Array {
					if _733_cf9 {
						return _697_v2
					}
					return _697_v2
				})(), _735_cf7, Companion_Default___.Fm16(_737_v32, _734_cf8, globalState), globalState)
				_738_v33 = _out19
			} else {
				var _739___mcc_h3 _dafny.Array = _source15.Get_().(D2_DC7).Cf6
				_ = _739___mcc_h3
				var _740_cf6 _dafny.Array = _739___mcc_h3
				_ = _740_cf6
				var _741_v34 _dafny.Map
				_ = _741_v34
				_741_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_695_v0), (_this).F9())
				var _742_v35 _dafny.Int
				_ = _742_v35
				var _out20 _dafny.Int
				_ = _out20
				_out20 = Companion_Default___.M0(Companion_Default___.Fm3(_724_v25, (_this).F9(), _695_v0, globalState), _697_v2, false, _741_v34, globalState)
				_742_v35 = _out20
				var _743_v36 T0
				_ = _743_v36
				var _nw125 *C0 = New_C0_()
				_ = _nw125
				_nw125.Ctor__(!_dafny.Companion_Sequence_.Contains(_702_v5, _724_v25), _695_v0, ((_this).F9()) || (!(false)))
				_743_v36 = _nw125
				var _744_v37 _dafny.Map
				_ = _744_v37
				_744_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_702_v5, _695_v0)
				var _745_v38 _dafny.Map
				_ = _745_v38
				_745_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(970), _697_v2)).Cardinality(), _702_v5)
				var _rhs120 T0 = _743_v36
				_ = _rhs120
				var _rhs121 bool = !(_745_v38).Equals(_745_v38)
				_ = _rhs121
				var _rhs122 _dafny.Int = _dafny.IntOfUint32((_701_v4).Cardinality())
				_ = _rhs122
				var _rhs123 _dafny.Map = _744_v37
				_ = _rhs123
				var _lhs82 *GlobalState = globalState
				_ = _lhs82
				_743_v36 = _rhs120
				_lhs82.F3 = _rhs121
				r2 = _rhs122
				_744_v37 = _rhs123
				r2 = Companion_Default___.SafeModuloInt(_695_v0, (_dafny.IntOfUint32((_701_v4).Cardinality())).Plus(_695_v0))
				_702_v5 = _dafny.UnicodeSeqOfUtf8Bytes("kmguhhp")
			}
			var _746_v39 *C0
			_ = _746_v39
			var _nw126 *C0 = New_C0_()
			_ = _nw126
			_nw126.Ctor__((_695_v0).Cmp(_dafny.IntOfInt64(-788)) < 0, _695_v0, (_this).F9())
			_746_v39 = _nw126
			_746_v39 = _746_v39
			var _747_v40 _dafny.Map
			_ = _747_v40
			_747_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F9())
			(globalState).F3 = (func() bool {
				if (_747_v40).Contains(!((_746_v39).F13())) {
					return (_747_v40).Get(!((_746_v39).F13())).(bool)
				}
				return false
			})()
			var _748_v41 _dafny.Map
			_ = _748_v41
			_748_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_707_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_707_v10), 0))).Int()).(_dafny.MultiSet)).Cardinality(), (func() _dafny.CodePoint {
				if (_746_v39).F13() {
					return (_this).Fm5(globalState)
				}
				return _724_v25
			})())
			_748_v41 = (_748_v41).Update(_695_v0, _724_v25)
		} else {
			var _749_v42 _dafny.Map
			_ = _749_v42
			_749_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_724_v25, Companion_Default___.SafeModuloInt(_695_v0, _695_v0))
			_749_v42 = _749_v42
			_702_v5 = _702_v5
			_695_v0 = (_dafny.IntOfInt64(150)).Times(_695_v0)
			r2 = _695_v0
			var _750_v43 _dafny.Array
			_ = _750_v43
			var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
			_ = _nw127
			_750_v43 = _nw127
			var _751_v44 _dafny.Array
			_ = _751_v44
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_11
			var _nw128 _dafny.Array
			_ = _nw128
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw128 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.CodePoint = (func(_752_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_753_i3 _dafny.Int) _dafny.CodePoint {
						return _752_v25
					}
				})(_724_v25)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw128 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw128).ArraySet1CodePoint(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw128).ArraySet1CodePoint(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_751_v44 = _nw128
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_750_v43), 0))
			_ = _index124
			(_750_v43).ArraySet1(_751_v44, (_index124).Int())
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_750_v43), 0))
			_ = _index125
			(_750_v43).ArraySet1(_751_v44, (_index125).Int())
		}
		r0 = ((_dafny.Zero).Minus(_695_v0)).Plus(Companion_Default___.SafeDivisionInt(_695_v0, _695_v0))
		var _754_v45 D1
		_ = _754_v45
		_754_v45 = Companion_D1_.Create_DC5_()
		var _pat_let_tv30 = _695_v0
		_ = _pat_let_tv30
		var _pat_let_tv31 = _695_v0
		_ = _pat_let_tv31
		var _pat_let_tv32 = _726_v27
		_ = _pat_let_tv32
		var _pat_let_tv33 = _695_v0
		_ = _pat_let_tv33
		var _pat_let_tv34 = _726_v27
		_ = _pat_let_tv34
		var _pat_let_tv35 = _695_v0
		_ = _pat_let_tv35
		var _pat_let_tv36 = _695_v0
		_ = _pat_let_tv36
		r1 = func(_source16 D1) D0 {
			if _source16.Is_DC5() {
				return Companion_D0_.Create_DC0_(_pat_let_tv30)
			} else if _source16.Is_DC6() {
				return Companion_D0_.Create_DC0_(_pat_let_tv31)
			} else {
				var _755___mcc_h4 _dafny.Sequence = _source16.Get_().(D1_DC4).Cf5
				_ = _755___mcc_h4
				var _756_cf5 _dafny.Sequence = _755___mcc_h4
				_ = _756_cf5
				return Companion_D0_.Create_DC0_((_dafny.SetOf((func() _dafny.Int {
					if (_pat_let_tv32).Contains(_pat_let_tv33) {
						return (_pat_let_tv34).Get(_pat_let_tv35).(_dafny.Int)
					}
					return _pat_let_tv36
				})())).Cardinality())
			}
		}(_754_v45)
		var _757_v46 _dafny.Set
		_ = _757_v46
		_757_v46 = _dafny.SetOf((_this).F9(), true, true)
		r2 = ((_757_v46).Intersection(_757_v46)).Cardinality()
		return r0, r1, r2
	}
}

// End of class C3
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
