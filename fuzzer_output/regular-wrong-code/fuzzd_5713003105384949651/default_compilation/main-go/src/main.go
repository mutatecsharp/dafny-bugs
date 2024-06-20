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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-490), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("guwwkcoj"))).Cardinality()), _dafny.IntOfInt64(-612))))
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(309))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})), _dafny.UnicodeSeqOfUtf8Bytes("dwtkddnh")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kq"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(218))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))))
}
func (_static *CompanionStruct_Default___) Fm4(globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((func() _dafny.Int {
		if !((true) && (false)) {
			return (_dafny.IntOfInt64(512)).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('s'))).Cardinality())
		}
		return Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(268), true)).Cardinality(), _dafny.IntOfInt64(589), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ddwiwxcq")).Cardinality()))).Cardinality()), _dafny.IntOfInt64(-555))
	})())
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC3_(_dafny.CodePoint('d'))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('s')
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(795))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	})), _dafny.UnicodeSeqOfUtf8Bytes("l"))
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(144), (func() _dafny.Set {
			var _coll1 = _dafny.NewBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((func() _dafny.Set {
				var _coll2 = _dafny.NewBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(66), _dafny.IntOfInt64(817))); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _3_v0 _dafny.Int
					_3_v0 = interface{}(_compr_2).(_dafny.Int)
					if ((_dafny.IntOfInt64(66)).Cmp(_3_v0) <= 0) && ((_3_v0).Cmp(_dafny.IntOfInt64(817)) < 0) {
						_coll2.Add(Companion_Default___.SafeModuloInt(_3_v0, _dafny.IntOfInt64(401)))
					}
				}
				return _coll2.ToSet()
			}()).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _4_v1 _dafny.Int
				_4_v1 = interface{}(_compr_1).(_dafny.Int)
				if (func() _dafny.Set {
					var _coll3 = _dafny.NewBuilder()
					_ = _coll3
					for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(66), _dafny.IntOfInt64(817))); ; {
						_compr_3, _ok3 := _iter3()
						if !_ok3 {
							break
						}
						var _5_v0 _dafny.Int
						_5_v0 = interface{}(_compr_3).(_dafny.Int)
						if ((_dafny.IntOfInt64(66)).Cmp(_5_v0) <= 0) && ((_5_v0).Cmp(_dafny.IntOfInt64(817)) < 0) {
							_coll3.Add(Companion_Default___.SafeModuloInt(_5_v0, _dafny.IntOfInt64(401)))
						}
					}
					return _coll3.ToSet()
				}()).Contains(_4_v1) {
					_coll1.Add(Companion_Default___.SafeModuloInt(_4_v1, (_dafny.SetOf(true, true)).Cardinality()))
				}
			}
			return _coll1.ToSet()
		}()).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Cardinality(), true)).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _6_v2 _dafny.Int
			_6_v2 = interface{}(_compr_0).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(144), (func() _dafny.Set {
				var _coll4 = _dafny.NewBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate((func() _dafny.Set {
					var _coll5 = _dafny.NewBuilder()
					_ = _coll5
					for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(66), _dafny.IntOfInt64(817))); ; {
						_compr_5, _ok5 := _iter5()
						if !_ok5 {
							break
						}
						var _7_v0 _dafny.Int
						_7_v0 = interface{}(_compr_5).(_dafny.Int)
						if ((_dafny.IntOfInt64(66)).Cmp(_7_v0) <= 0) && ((_7_v0).Cmp(_dafny.IntOfInt64(817)) < 0) {
							_coll5.Add(Companion_Default___.SafeModuloInt(_7_v0, _dafny.IntOfInt64(401)))
						}
					}
					return _coll5.ToSet()
				}()).Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _8_v1 _dafny.Int
					_8_v1 = interface{}(_compr_4).(_dafny.Int)
					if (func() _dafny.Set {
						var _coll6 = _dafny.NewBuilder()
						_ = _coll6
						for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(66), _dafny.IntOfInt64(817))); ; {
							_compr_6, _ok6 := _iter6()
							if !_ok6 {
								break
							}
							var _9_v0 _dafny.Int
							_9_v0 = interface{}(_compr_6).(_dafny.Int)
							if ((_dafny.IntOfInt64(66)).Cmp(_9_v0) <= 0) && ((_9_v0).Cmp(_dafny.IntOfInt64(817)) < 0) {
								_coll6.Add(Companion_Default___.SafeModuloInt(_9_v0, _dafny.IntOfInt64(401)))
							}
						}
						return _coll6.ToSet()
					}()).Contains(_8_v1) {
						_coll4.Add(Companion_Default___.SafeModuloInt(_8_v1, (_dafny.SetOf(true, true)).Cardinality()))
					}
				}
				return _coll4.ToSet()
			}()).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Cardinality(), true)).Contains(_6_v2) {
				_coll0.Add(Companion_Default___.SafeDivisionInt(_6_v2, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vup")).Cardinality()))))
			}
		}
		return _coll0.ToSet()
	}()).Intersection((func() _dafny.Set {
		if !(true) {
			return func() _dafny.Set {
				var _coll7 = _dafny.NewBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(704), _dafny.IntOfInt64(-431))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _10_v3 _dafny.Int
					_10_v3 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(704)).Cmp(_10_v3) <= 0) && ((_10_v3).Cmp(_dafny.IntOfInt64(-431)) < 0) {
						_coll7.Add((_10_v3).Minus(_dafny.IntOfInt64(-676)))
					}
				}
				return _coll7.ToSet()
			}()
		}
		return _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qdem")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true, true, true, !((Companion_D0_.Create_DC0_(true)).Dtor_cf0()), false)).Cardinality()))
	})())
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false, false, !(!(!(false))), true)).Cardinality()))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(-769), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(798))).Cardinality())))).Difference((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-619)))).Union(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(752))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_11_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	}))).Cardinality()), _dafny.IntOfInt64(803), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D1_.Create_DC5_(false, false)).Dtor_cf8(), true)).Cardinality(), _dafny.IntOfInt64(648))))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	if (func() bool {
		if true {
			return true
		}
		return true
	})() {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-402)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mbafox")).Cardinality()), (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), !(!(false))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Cardinality()))
	} else {
		return _dafny.SeqOf(_dafny.IntOfInt64(-643))
	}
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.CodePoint, p1 _dafny.MultiSet, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC5_(true, (_dafny.SetOf(_dafny.SeqOf(false), _dafny.SeqOf(true))).IsSubsetOf(_dafny.SetOf(_dafny.SeqOf(false), _dafny.SeqOf(!(true), true, false, true, true))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(920), _dafny.IntOfInt64(101))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _12_v0 _dafny.Int
			_12_v0 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(920)).Cmp(_12_v0) <= 0) && ((_12_v0).Cmp(_dafny.IntOfInt64(101)) < 0) {
				_coll8.Add((_12_v0).Plus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(621))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg4 _dafny.Int) interface{} {
						return coer4(arg4)
					}
				}(func(_13_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yfuahe")).Cardinality())
				})))).Cardinality())
			}
		}
		return _coll8.ToMap()
	}())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(Companion_D11_.Create_DC25_(), Companion_D11_.Create_DC25_(), Companion_D11_.Create_DC25_())).Cardinality()), _dafny.IntOfInt64(-411))).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _14_v1 _dafny.Int
			_14_v1 = interface{}(_compr_9).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(Companion_D11_.Create_DC25_(), Companion_D11_.Create_DC25_(), Companion_D11_.Create_DC25_())).Cardinality()), _dafny.IntOfInt64(-411)), _14_v1) {
				_coll9.Add(Companion_Default___.SafeModuloInt(_14_v1, _dafny.IntOfInt64(280)))
			}
		}
		return _coll9.ToSet()
	}()).Cardinality(), _dafny.IntOfInt64(114)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(424), _dafny.IntOfInt64(154)))))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(true)).Difference(_dafny.SetOf(false))).Difference((_dafny.SetOf(false)).Intersection(_dafny.SetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(!(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.CodePoint('g')), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(258))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_15_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	})))), true, (false) || (false))
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 _dafny.Sequence, p2 bool, p3 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(216))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_16_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("nrkvxhch")
	})))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ugtro")))), _dafny.IntOfInt64(149))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("odvbie"))).Intersection(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("efm"), _dafny.UnicodeSeqOfUtf8Bytes("ctebj")))).Difference((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("wmdxukk"), _dafny.UnicodeSeqOfUtf8Bytes("m"), _dafny.UnicodeSeqOfUtf8Bytes("dlv"))).Difference(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("affp"), _dafny.UnicodeSeqOfUtf8Bytes("v"))))
}
func (_static *CompanionStruct_Default___) Fm19(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cjrbft")).Cardinality())), _dafny.IntOfInt64(587))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-296), _dafny.IntOfInt64(464))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(537), _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(932), _dafny.IntOfInt64(302))); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _17_v0 _dafny.Int
			_17_v0 = interface{}(_compr_10).(_dafny.Int)
			if ((_dafny.IntOfInt64(932)).Cmp(_17_v0) <= 0) && ((_17_v0).Cmp(_dafny.IntOfInt64(302)) < 0) {
				_coll10.Add(Companion_Default___.SafeDivisionInt(_17_v0, (_dafny.Zero).Minus(_dafny.IntOfInt64(-106))), false)
			}
		}
		return _coll10.ToMap()
	}()).Cardinality())).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(!(false), true)).Cardinality(), false)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf((_dafny.MultiSetOf(true)).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(795)), _dafny.SeqOf(_dafny.IntOfInt64(521)))).Cardinality()), !(false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(505), true)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(265))).Cardinality(), !(false))))
}
func (_static *CompanionStruct_Default___) Fm21(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	var _source0 D0 = Companion_D0_.Create_DC1_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(908))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_18_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	}))).Cardinality())), _dafny.IntOfInt64(-506), _dafny.IntOfInt64(562))
	_ = _source0
	if _source0.Is_DC1() {
		var _19___mcc_h0 _dafny.Int = _source0.Get_().(D0_DC1).Cf1
		_ = _19___mcc_h0
		var _20___mcc_h1 _dafny.Int = _source0.Get_().(D0_DC1).Cf2
		_ = _20___mcc_h1
		var _21___mcc_h2 _dafny.Int = _source0.Get_().(D0_DC1).Cf3
		_ = _21___mcc_h2
		var _22_cf3 _dafny.Int = _21___mcc_h2
		_ = _22_cf3
		var _23_cf2 _dafny.Int = _20___mcc_h1
		_ = _23_cf2
		var _24_cf1 _dafny.Int = _19___mcc_h0
		_ = _24_cf1
		return func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate((_dafny.SeqOf(Companion_D1_.Create_DC6_(Companion_D1_.Create_DC3_(_dafny.CodePoint('f'))), Companion_D1_.Create_DC6_(Companion_D1_.Create_DC5_(false, true)))).Elements()); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _25_v0 D1
				_25_v0 = interface{}(_compr_11).(D1)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D1_.Create_DC6_(Companion_D1_.Create_DC3_(_dafny.CodePoint('f'))), Companion_D1_.Create_DC6_(Companion_D1_.Create_DC5_(false, true))), _25_v0) {
					_coll11.Add(_25_v0, _dafny.CodePoint('y'))
				}
			}
			return _coll11.ToMap()
		}()
	} else if _source0.Is_DC0() {
		var _26___mcc_h3 bool = _source0.Get_().(D0_DC0).Cf0
		_ = _26___mcc_h3
		var _27_cf0 bool = _26___mcc_h3
		_ = _27_cf0
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC6_(Companion_D1_.Create_DC4_(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(888))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_28_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		})))), _dafny.IntOfInt64(720))), _dafny.CodePoint('k'))
	} else {
		var _29___mcc_h4 D0 = _source0.Get_().(D0_DC2).Cf4
		_ = _29___mcc_h4
		var _30_cf4 D0 = _29___mcc_h4
		_ = _30_cf4
		return (Companion_D15_.Create_DC31_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC6_(Companion_D1_.Create_DC6_(Companion_D1_.Create_DC3_(_dafny.CodePoint('g')))), _dafny.CodePoint('r')))).Dtor_cf43()
	}
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Array, p1 _dafny.Int, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) {
	var _31_v0 bool
	_ = _31_v0
	_31_v0 = false
	var _32_v1 _dafny.Sequence
	_ = _32_v1
	_32_v1 = _dafny.SeqOf(_31_v0, _31_v0, _31_v0)
	var _33_v2 _dafny.Map
	_ = _33_v2
	_33_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v0, _32_v1)
	_33_v2 = (_33_v2).Update(true, _32_v1)
	if (_32_v1).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_32_v1).Cardinality()))).Uint32()).(bool) {
		var _34_v3 D3
		_ = _34_v3
		_34_v3 = Companion_D3_.Create_DC10_(_31_v0, p3, _31_v0, p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !(_31_v0)))
		var _35_v4 D3
		_ = _35_v4
		_35_v4 = Companion_D3_.Create_DC11_(_34_v3)
		_35_v4 = _35_v4
		var _36_v5 D1
		_ = _36_v5
		_36_v5 = Companion_D1_.Create_DC5_(_31_v0, _31_v0)
		_31_v0 = (_36_v5).Dtor_cf8()
		var _37_v6 D1
		_ = _37_v6
		_37_v6 = Companion_D1_.Create_DC5_(_31_v0, _31_v0)
		var _38_v7 D1
		_ = _38_v7
		_38_v7 = Companion_D1_.Create_DC6_(_37_v6)
		var _39_v8 *C0
		_ = _39_v8
		var _nw0 *C0 = New_C0_()
		_ = _nw0
		_nw0.Ctor__()
		_39_v8 = _nw0
		var _40_v9 _dafny.MultiSet
		_ = _40_v9
		_40_v9 = _dafny.MultiSetOf(_31_v0, true)
		var _41_v10 _dafny.Sequence
		_ = _41_v10
		_41_v10 = _dafny.SeqOf(p1)
		var _42_v11 _dafny.Sequence
		_ = _42_v11
		_42_v11 = _dafny.SeqOf((_dafny.Zero).Minus((_40_v9).Cardinality()), p3, _dafny.IntOfUint32((_41_v10).Cardinality()))
		var _43_v12 _dafny.Sequence
		_ = _43_v12
		_43_v12 = _dafny.SeqOf(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_42_v11, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_42_v11).Cardinality()))).Uint32(), p1)).Cardinality()))
		var _44_v13 *C2
		_ = _44_v13
		var _nw1 *C2 = New_C2_()
		_ = _nw1
		_nw1.Ctor__(p1)
		_44_v13 = _nw1
		var _45_v14 _dafny.Map
		_ = _45_v14
		_45_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _44_v13)
		var _46_v15 _dafny.Map
		_ = _46_v15
		_46_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v0, (_45_v14).Cardinality())
		var _47_v16 *C1
		_ = _47_v16
		var _nw2 *C1 = New_C1_()
		_ = _nw2
		_nw2.Ctor__(false)
		_47_v16 = _nw2
		var _48_v17 _dafny.Map
		_ = _48_v17
		_48_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v0, _47_v16)
		var _49_v18 _dafny.Array
		_ = _49_v18
		var _nwElement0_0 _dafny.Int = p1
		_ = _nwElement0_0
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(24))
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_0, 0)
		(_nw3).ArraySet1(p3, 1)
		(_nw3).ArraySet1(p3, 2)
		(_nw3).ArraySet1((_43_v12).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_43_v12).Cardinality()))).Uint32()).(_dafny.Int), 3)
		(_nw3).ArraySet1(p1, 4)
		(_nw3).ArraySet1(p3, 5)
		(_nw3).ArraySet1(p1, 6)
		(_nw3).ArraySet1(p3, 7)
		(_nw3).ArraySet1(p1, 8)
		(_nw3).ArraySet1(p3, 9)
		(_nw3).ArraySet1(_dafny.IntOfUint32((_41_v10).Cardinality()), 10)
		(_nw3).ArraySet1(p3, 11)
		(_nw3).ArraySet1(p3, 12)
		(_nw3).ArraySet1((_dafny.Zero).Minus(p1), 13)
		(_nw3).ArraySet1(_dafny.One, 14)
		(_nw3).ArraySet1(p1, 15)
		(_nw3).ArraySet1((func() _dafny.Int {
			if (_46_v15).Contains(_31_v0) {
				return (_46_v15).Get(_31_v0).(_dafny.Int)
			}
			return p3
		})(), 16)
		(_nw3).ArraySet1(_dafny.IntOfInt64(597), 17)
		(_nw3).ArraySet1(p3, 18)
		(_nw3).ArraySet1(p1, 19)
		(_nw3).ArraySet1(_dafny.IntOfInt64(-469), 20)
		(_nw3).ArraySet1(_44_v13.F13, 21)
		(_nw3).ArraySet1(_dafny.IntOfInt64(896), 22)
		(_nw3).ArraySet1((_48_v17).Cardinality(), 23)
		_49_v18 = _nw3
		var _50_v19 _dafny.Sequence
		_ = _50_v19
		_50_v19 = _dafny.SeqOf(_49_v18, _49_v18)
		var _51_v20 _dafny.Array
		_ = _51_v20
		var _nwElement0_1 _dafny.Array = _49_v18
		_ = _nwElement0_1
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(22))
		_ = _nw4
		(_nw4).ArraySet1(_nwElement0_1, 0)
		(_nw4).ArraySet1(_49_v18, 1)
		(_nw4).ArraySet1(_49_v18, 2)
		(_nw4).ArraySet1(_49_v18, 3)
		(_nw4).ArraySet1(_49_v18, 4)
		(_nw4).ArraySet1(_49_v18, 5)
		(_nw4).ArraySet1(_49_v18, 6)
		(_nw4).ArraySet1(_49_v18, 7)
		(_nw4).ArraySet1((_50_v19).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_50_v19).Cardinality()))).Uint32()).(_dafny.Array), 8)
		(_nw4).ArraySet1(_49_v18, 9)
		(_nw4).ArraySet1(_49_v18, 10)
		(_nw4).ArraySet1(_49_v18, 11)
		(_nw4).ArraySet1(_49_v18, 12)
		(_nw4).ArraySet1(_49_v18, 13)
		(_nw4).ArraySet1(_49_v18, 14)
		(_nw4).ArraySet1(_49_v18, 15)
		(_nw4).ArraySet1(_49_v18, 16)
		(_nw4).ArraySet1(_49_v18, 17)
		(_nw4).ArraySet1(_49_v18, 18)
		(_nw4).ArraySet1(_49_v18, 19)
		(_nw4).ArraySet1(_49_v18, 20)
		(_nw4).ArraySet1(_49_v18, 21)
		_51_v20 = _nw4
		var _rhs0 D1 = _38_v7
		_ = _rhs0
		var _rhs1 _dafny.Array = _49_v18
		_ = _rhs1
		var _rhs2 bool = _31_v0
		_ = _rhs2
		var _rhs3 *C0 = _39_v8
		_ = _rhs3
		var _rhs4 _dafny.Array = _51_v20
		_ = _rhs4
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		_38_v7 = _rhs0
		_lhs0.F10 = _rhs1
		_31_v0 = _rhs2
		_39_v8 = _rhs3
		_51_v20 = _rhs4
		var _52_v21 _dafny.Sequence
		_ = _52_v21
		_52_v21 = _32_v1
		var _53_v22 _dafny.Array
		_ = _53_v22
		var _nwElement0_2 _dafny.Sequence = (func() _dafny.Sequence {
			if _47_v16.F14 {
				return _32_v1
			}
			return _32_v1
		})()
		_ = _nwElement0_2
		var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(13))
		_ = _nw5
		(_nw5).ArraySet1(_nwElement0_2, 0)
		(_nw5).ArraySet1(_32_v1, 1)
		(_nw5).ArraySet1(_32_v1, 2)
		(_nw5).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_47_v16.F14, _47_v16.F14, _47_v16.F14, _47_v16.F14), _32_v1), 3)
		(_nw5).ArraySet1(_dafny.SeqOf(_47_v16.F14), 4)
		(_nw5).ArraySet1(_32_v1, 5)
		(_nw5).ArraySet1(_32_v1, 6)
		(_nw5).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_32_v1, _32_v1), 7)
		(_nw5).ArraySet1(_dafny.SeqOf(_31_v0, _47_v16.F14), 8)
		(_nw5).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_32_v1, _32_v1), 9)
		(_nw5).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_32_v1, _32_v1), 10)
		(_nw5).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_52_v21), _32_v1), 11)
		(_nw5).ArraySet1(_32_v1, 12)
		_53_v22 = _nw5
		var _54_v23 _dafny.CodePoint
		_ = _54_v23
		_54_v23 = _dafny.CodePoint('l')
		var _rhs5 _dafny.Array = _53_v22
		_ = _rhs5
		var _rhs6 _dafny.CodePoint = _54_v23
		_ = _rhs6
		var _rhs7 bool = (p3).Cmp((_dafny.Zero).Minus(_44_v13.F13)) >= 0
		_ = _rhs7
		var _lhs1 *C1 = _47_v16
		_ = _lhs1
		_53_v22 = _rhs5
		_54_v23 = _rhs6
		_lhs1.F14 = _rhs7
		var _55_v24 _dafny.Set
		_ = _55_v24
		_55_v24 = _dafny.SetOf(_32_v1)
		_55_v24 = (_55_v24).Union(_55_v24)
	} else {
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((p0), 0))
		_ = _index0
		(p0).ArraySet1(_31_v0, (_index0).Int())
		var _56_v25 _dafny.Array
		_ = _56_v25
		var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
		_ = _nw6
		_56_v25 = _nw6
		var _57_v26 _dafny.Sequence
		_ = _57_v26
		_57_v26 = _dafny.SeqOf(p1, p3)
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_56_v25), 0))
		_ = _index1
		(_56_v25).ArraySet1(_57_v26, (_index1).Int())
		var _58_v27 _dafny.Sequence
		_ = _58_v27
		_58_v27 = _dafny.UnicodeSeqOfUtf8Bytes("wwcctfdwq")
		var _59_v28 _dafny.Map
		_ = _59_v28
		_59_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, !(_31_v0))
		var _60_v29 D7
		_ = _60_v29
		_60_v29 = Companion_D7_.Create_DC17_(_31_v0)
		var _61_v30 D11
		_ = _61_v30
		_61_v30 = Companion_D11_.Create_DC23_(_57_v26)
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((p0), 0))
		_ = _index2
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_56_v25), 0))
		_ = _index3
		var _rhs8 bool = (_dafny.Companion_Sequence_.Contains(_57_v26, _dafny.IntOfInt64(233))) && (_dafny.Companion_Sequence_.IsPrefixOf(_58_v27, _dafny.UnicodeSeqOfUtf8Bytes("g")))
		_ = _rhs8
		var _rhs9 _dafny.Int = ((_59_v28).Merge(Companion_Default___.Fm20(_31_v0, (_dafny.Zero).Minus(p3), Companion_Default___.Fm0(true, _31_v0, (_60_v29).Dtor_cf25(), _31_v0, globalState), globalState))).Cardinality()
		_ = _rhs9
		var _rhs10 bool = (func() bool {
			if (func() bool {
				if !(Companion_Default___.Fm0(_31_v0, _31_v0, _31_v0, _31_v0, globalState)) {
					return _31_v0
				}
				return false
			})() {
				return _31_v0
			}
			return (func() bool {
				if _31_v0 {
					return _31_v0
				}
				return _31_v0
			})()
		})()
		_ = _rhs10
		var _rhs11 _dafny.Sequence = _dafny.Companion_Sequence_.Update((_61_v30).Dtor_cf34(), (Companion_Default___.SafeIndex((p1).Times(p3), _dafny.IntOfUint32(((_61_v30).Dtor_cf34()).Cardinality()))).Uint32(), (_dafny.Zero).Minus(p1))
		_ = _rhs11
		var _lhs2 _dafny.Array = p0
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((p0), 0))
		_ = _lhs3
		var _lhs4 *GlobalState = globalState
		_ = _lhs4
		var _lhs5 _dafny.Array = _56_v25
		_ = _lhs5
		var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_56_v25), 0))
		_ = _lhs6
		(_lhs2).ArraySet1(_rhs8, (_lhs3).Int())
		_lhs4.F2 = _rhs9
		_31_v0 = _rhs10
		(_lhs5).ArraySet1(_rhs11, (_lhs6).Int())
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((p0), 0))
		_ = _index4
		(p0).ArraySet1(_31_v0, (_index4).Int())
		var _62_v31 *C1
		_ = _62_v31
		var _nw7 *C1 = New_C1_()
		_ = _nw7
		_nw7.Ctor__(_31_v0)
		_62_v31 = _nw7
		var _63_v32 _dafny.Map
		_ = _63_v32
		_63_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Cmp(p1) <= 0, _62_v31)
		var _64_v33 _dafny.Map
		_ = _64_v33
		_64_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_57_v26).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_57_v26).Cardinality()))).Uint32()).(_dafny.Int), _62_v31)
		_62_v31 = (func() *C1 {
			if (_63_v32).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((p0), 0))).Int()).(bool)) {
				return (_63_v32).Get((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((p0), 0))).Int()).(bool)).(*C1)
			}
			return (func() *C1 {
				if (_64_v33).Contains(p3) {
					return (_64_v33).Get(p3).(*C1)
				}
				return _62_v31
			})()
		})()
		(globalState).F2 = ((_56_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_56_v25), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_58_v27).Cardinality()), _dafny.IntOfUint32(((_56_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_56_v25), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Int)
		var _65_v34 _dafny.MultiSet
		_ = _65_v34
		_65_v34 = _dafny.MultiSetOf(_31_v0, _31_v0)
		var _66_v35 *C2
		_ = _66_v35
		var _nw8 *C2 = New_C2_()
		_ = _nw8
		_nw8.Ctor__((_65_v34).Cardinality())
		_66_v35 = _nw8
		var _67_v36 _dafny.CodePoint
		_ = _67_v36
		_67_v36 = _dafny.CodePoint('x')
		var _68_v37 _dafny.Map
		_ = _68_v37
		_68_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_66_v35, _67_v36)
		(globalState).F2 = (p1).Minus((_68_v37).Cardinality())
	}
	if _31_v0 {
		_31_v0 = _31_v0
		var _69_v38 *C0
		_ = _69_v38
		var _nw9 *C0 = New_C0_()
		_ = _nw9
		_nw9.Ctor__()
		_69_v38 = _nw9
		var _70_v39 _dafny.Sequence
		_ = _70_v39
		_70_v39 = _dafny.SeqOf(_69_v38, _69_v38)
		var _71_v40 _dafny.Sequence
		_ = _71_v40
		_71_v40 = _70_v39
		var _72_v41 _dafny.Map
		_ = _72_v41
		_72_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_71_v40), (_dafny.Zero).Minus(p1))
		var _73_v42 _dafny.Map
		_ = _73_v42
		_73_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(718), Companion_Default___.Fm4(globalState))
		var _74_v43 _dafny.Map
		_ = _74_v43
		_74_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v0, _31_v0)
		var _75_v44 _dafny.Sequence
		_ = _75_v44
		_75_v44 = _dafny.SeqOf(_74_v43)
		_72_v41 = (_72_v41).Update(_dafny.Companion_Sequence_.Update(_70_v39, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_70_v39).Cardinality()))).Uint32(), _69_v38), ((func() _dafny.Int {
			if (_73_v42).Contains(_dafny.IntOfUint32((_75_v44).Cardinality())) {
				return (_73_v42).Get(_dafny.IntOfUint32((_75_v44).Cardinality())).(_dafny.Int)
			}
			return p3
		})()).Minus(p1))
		var _76_v45 _dafny.Set
		_ = _76_v45
		_76_v45 = _dafny.SetOf(p3)
		var _77_v46 _dafny.Map
		_ = _77_v46
		_77_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_76_v45).IsProperSubsetOf(_dafny.SetOf(p1))), p3)
		_77_v46 = (_77_v46).Update(Companion_Default___.Fm0(_31_v0, _31_v0, !(_31_v0), _31_v0, globalState), p3)
		var _78_v47 D11
		_ = _78_v47
		_78_v47 = Companion_D11_.Create_DC24_(p3)
		_78_v47 = _78_v47
		if !(_31_v0) {
			var _79_v48 *C0
			_ = _79_v48
			var _nw10 *C0 = New_C0_()
			_ = _nw10
			_nw10.Ctor__()
			_79_v48 = _nw10
			_31_v0 = _31_v0
			var _80_v49 _dafny.Array
			_ = _80_v49
			var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw11
			_80_v49 = _nw11
			var _81_v50 *C2
			_ = _81_v50
			var _nw12 *C2 = New_C2_()
			_ = _nw12
			_nw12.Ctor__(p3)
			_81_v50 = _nw12
			var _82_v51 _dafny.Map
			_ = _82_v51
			_82_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_81_v50, _80_v49)
			var _83_v52 _dafny.Array
			_ = _83_v52
			var _nwElement0_3 _dafny.Array = _80_v49
			_ = _nwElement0_3
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(25))
			_ = _nw13
			(_nw13).ArraySet1(_nwElement0_3, 0)
			(_nw13).ArraySet1(_80_v49, 1)
			(_nw13).ArraySet1(_80_v49, 2)
			(_nw13).ArraySet1(_80_v49, 3)
			(_nw13).ArraySet1(_80_v49, 4)
			(_nw13).ArraySet1(_80_v49, 5)
			(_nw13).ArraySet1(_80_v49, 6)
			(_nw13).ArraySet1(_80_v49, 7)
			(_nw13).ArraySet1(_80_v49, 8)
			(_nw13).ArraySet1(_80_v49, 9)
			(_nw13).ArraySet1(_80_v49, 10)
			(_nw13).ArraySet1(_80_v49, 11)
			(_nw13).ArraySet1(_80_v49, 12)
			(_nw13).ArraySet1(_80_v49, 13)
			(_nw13).ArraySet1((func() _dafny.Array {
				if (_82_v51).Contains(_81_v50) {
					return (_82_v51).Get(_81_v50).(_dafny.Array)
				}
				return _80_v49
			})(), 14)
			(_nw13).ArraySet1(_80_v49, 15)
			(_nw13).ArraySet1(_80_v49, 16)
			(_nw13).ArraySet1(_80_v49, 17)
			(_nw13).ArraySet1(_80_v49, 18)
			(_nw13).ArraySet1(_80_v49, 19)
			(_nw13).ArraySet1(_80_v49, 20)
			(_nw13).ArraySet1(_80_v49, 21)
			(_nw13).ArraySet1(_80_v49, 22)
			(_nw13).ArraySet1(_80_v49, 23)
			(_nw13).ArraySet1(_80_v49, 24)
			_83_v52 = _nw13
			var _84_v53 _dafny.Map
			_ = _84_v53
			_84_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), _80_v49)
			var _85_v54 _dafny.CodePoint
			_ = _85_v54
			_85_v54 = _dafny.CodePoint('g')
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_83_v52), 0))
			_ = _index5
			(_83_v52).ArraySet1((func() _dafny.Array {
				if (_84_v53).Contains(_85_v54) {
					return (_84_v53).Get(_85_v54).(_dafny.Array)
				}
				return _80_v49
			})(), (_index5).Int())
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_83_v52), 0))
			_ = _index6
			(_83_v52).ArraySet1(_80_v49, (_index6).Int())
			var _86_v55 _dafny.Map
			_ = _86_v55
			_86_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v38, (_77_v46).Cardinality())
			_86_v55 = (_86_v55).Update(_69_v38, p3)
			(globalState).F2 = _81_v50.F13
		} else {
			var _87_v56 _dafny.Array
			_ = _87_v56
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_0
			var _nw14 _dafny.Array
			_ = _nw14
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw14 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Int = (func(_88_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_89_i0 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_89_i0, _88_p3)
					}
				})(p3)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw14 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw14).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw14).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_87_v56 = _nw14
			var _90_v57 _dafny.Map
			_ = _90_v57
			_90_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_v56, p1)
			var _91_v58 _dafny.MultiSet
			_ = _91_v58
			_91_v58 = _dafny.MultiSetOf(p3, p3)
			var _92_v59 _dafny.Set
			_ = _92_v59
			_92_v59 = _dafny.SetOf(_31_v0, _31_v0, _31_v0)
			_73_v42 = (_73_v42).Update((func() _dafny.Int {
				if (_90_v57).Contains(_87_v56) {
					return (_90_v57).Get(_87_v56).(_dafny.Int)
				}
				return (_69_v38).Fm8((_91_v58).Cardinality(), _92_v59, p1, globalState)
			})(), p1)
			var _93_v60 _dafny.Array
			_ = _93_v60
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_1
			var _nw15 _dafny.Array
			_ = _nw15
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw15 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) D1 = func(_94_i1 _dafny.Int) D1 {
					return Companion_D1_.Create_DC6_(Companion_D1_.Create_DC6_(Companion_D1_.Create_DC3_(_dafny.CodePoint('h'))))
				}
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw15 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw15).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw15).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_93_v60 = _nw15
			var _95_v61 D2
			_ = _95_v61
			_95_v61 = Companion_D2_.Create_DC7_(_93_v60)
			var _96_v62 _dafny.CodePoint
			_ = _96_v62
			_96_v62 = _dafny.CodePoint('m')
			var _rhs12 _dafny.Int = (p1).Minus(Companion_Default___.SafeDivisionInt(p3, (Companion_Default___.Fm11(_31_v0, _96_v62, globalState)).Cardinality()))
			_ = _rhs12
			var _rhs13 bool = ((_dafny.Zero).Minus(p3)).Cmp((_dafny.Zero).Minus(p1)) >= 0
			_ = _rhs13
			var _rhs14 D2 = _95_v61
			_ = _rhs14
			var _lhs7 *GlobalState = globalState
			_ = _lhs7
			_lhs7.F2 = _rhs12
			_31_v0 = _rhs13
			_95_v61 = _rhs14
			(globalState).F2 = (_dafny.Zero).Minus((_dafny.Zero).Minus(p1))
			var _97_v63 *C2
			_ = _97_v63
			var _nw16 *C2 = New_C2_()
			_ = _nw16
			_nw16.Ctor__(p1)
			_97_v63 = _nw16
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_87_v56), 0))
			_ = _index7
			(_87_v56).ArraySet1(Companion_Default___.Fm2(_31_v0, globalState), (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_87_v56), 0))
			_ = _index8
			(_87_v56).ArraySet1(_97_v63.F13, (_index8).Int())
		}
	} else {
		var _98_v64 D7
		_ = _98_v64
		_98_v64 = Companion_D7_.Create_DC18_(Companion_D7_.Create_DC17_(_31_v0))
		var _source1 D7 = _98_v64
		_ = _source1
		if _source1.Is_DC17() {
			var _99___mcc_h0 bool = _source1.Get_().(D7_DC17).Cf25
			_ = _99___mcc_h0
			var _100_cf25 bool = _99___mcc_h0
			_ = _100_cf25
			_100_cf25 = ((p1).Cmp(p3) != 0) && (!(_31_v0))
			var _101_v65 _dafny.Map
			_ = _101_v65
			_101_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _100_cf25)
			var _102_v66 _dafny.Sequence
			_ = _102_v66
			_102_v66 = _dafny.UnicodeSeqOfUtf8Bytes("dpwnr")
			var _103_v67 _dafny.Map
			_ = _103_v67
			_103_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((Companion_D3_.Create_DC10_(_31_v0, p3, _31_v0, p0, _101_v65)).Dtor_cf16()), _102_v66)
			var _104_v68 _dafny.Sequence
			_ = _104_v68
			_104_v68 = _dafny.SeqOf((_103_v67).Cardinality())
			var _rhs15 _dafny.Int = (_104_v68).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.IntOfUint32((_104_v68).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs15
			var _rhs16 _dafny.Int = p1
			_ = _rhs16
			var _lhs8 *GlobalState = globalState
			_ = _lhs8
			var _lhs9 *GlobalState = globalState
			_ = _lhs9
			_lhs8.F2 = _rhs15
			_lhs9.F2 = _rhs16
			var _105_v69 *C2
			_ = _105_v69
			var _nw17 *C2 = New_C2_()
			_ = _nw17
			_nw17.Ctor__((func() _dafny.Int {
				if _31_v0 {
					return p1
				}
				return p1
			})())
			_105_v69 = _nw17
			_105_v69 = _105_v69
			var _106_v70 *C3
			_ = _106_v70
			var _nw18 *C3 = New_C3_()
			_ = _nw18
			_nw18.Ctor__(_100_cf25, _dafny.IntOfInt64(7))
			_106_v70 = _nw18
		} else if _source1.Is_DC16() {
			var _107___mcc_h1 *C0 = _source1.Get_().(D7_DC16).Cf24
			_ = _107___mcc_h1
			var _108_cf24 *C0 = _107___mcc_h1
			_ = _108_cf24
			var _109_v71 _dafny.Array
			_ = _109_v71
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_2
			var _nw19 _dafny.Array
			_ = _nw19
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw19 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Sequence = func(_110_i2 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('a'), _dafny.CodePoint('y'), _dafny.CodePoint('f'))
				}
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw19 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw19).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw19).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_109_v71 = _nw19
			var _111_v72 _dafny.Array
			_ = _111_v72
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_3
			var _nw20 _dafny.Array
			_ = _nw20
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw20 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Int = (func(_112_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_113_i3 _dafny.Int) _dafny.Int {
						return (_113_i3).Minus((_dafny.Zero).Minus(_112_p3))
					}
				})(p3)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw20 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw20).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw20).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_111_v72 = _nw20
			var _114_v73 _dafny.Sequence
			_ = _114_v73
			_114_v73 = _dafny.UnicodeSeqOfUtf8Bytes("joaugkus")
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_111_v72), 0))
			_ = _index9
			(_111_v72).ArraySet1(_dafny.IntOfUint32((_114_v73).Cardinality()), (_index9).Int())
			var _115_v74 _dafny.Map
			_ = _115_v74
			_115_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Cmp(p3) != 0, _dafny.Companion_Sequence_.Concatenate(_114_v73, _114_v73))
			var _116_v75 _dafny.Array
			_ = _116_v75
			var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(16))
			_ = _nw21
			_116_v75 = _nw21
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_111_v72), 0))
			_ = _index10
			var _rhs17 _dafny.Array = (func() _dafny.Array {
				if _31_v0 {
					return _109_v71
				}
				return _109_v71
			})()
			_ = _rhs17
			var _rhs18 _dafny.Int = p1
			_ = _rhs18
			var _rhs19 _dafny.Array = _109_v71
			_ = _rhs19
			var _rhs20 _dafny.Map = (_115_v74).Merge(_115_v74)
			_ = _rhs20
			var _rhs21 _dafny.Array = _116_v75
			_ = _rhs21
			var _lhs10 _dafny.Array = _111_v72
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_111_v72), 0))
			_ = _lhs11
			_109_v71 = _rhs17
			(_lhs10).ArraySet1(_rhs18, (_lhs11).Int())
			_109_v71 = _rhs19
			_115_v74 = _rhs20
			_116_v75 = _rhs21
			var _117_v76 _dafny.Set
			_ = _117_v76
			_117_v76 = _dafny.SetOf(_31_v0)
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((p0), 0))
			_ = _index11
			(p0).ArraySet1((_117_v76).IsDisjointFrom(_117_v76), (_index11).Int())
			var _118_v77 _dafny.Array
			_ = _118_v77
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_4
			var _nw22 _dafny.Array
			_ = _nw22
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw22 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Map = (func(_119_p3 _dafny.Int, _120_p1 _dafny.Int, _121_v72 _dafny.Array) func(_dafny.Int) _dafny.Map {
					return func(_122_i4 _dafny.Int) _dafny.Map {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_p3, _120_p1)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_121_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_121_v72), 0))).Int()).(_dafny.Int), _119_p3))
					}
				})(p3, p1, _111_v72)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw22 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw22).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw22).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_118_v77 = _nw22
			var _123_v78 _dafny.Map
			_ = _123_v78
			_123_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nufq")).Cardinality()), p1)
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_118_v77), 0))
			_ = _index12
			(_118_v77).ArraySet1(_123_v78, (_index12).Int())
			var _124_v79 _dafny.MultiSet
			_ = _124_v79
			_124_v79 = _dafny.MultiSetOf(p1)
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((p0), 0))
			_ = _index13
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_118_v77), 0))
			_ = _index14
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_111_v72), 0))
			_ = _index15
			var _rhs22 _dafny.Int = (p1).Minus(p3)
			_ = _rhs22
			var _rhs23 bool = (_124_v79).IsSubsetOf(_124_v79)
			_ = _rhs23
			var _rhs24 _dafny.Map = _123_v78
			_ = _rhs24
			var _rhs25 _dafny.Int = p1
			_ = _rhs25
			var _lhs12 *GlobalState = globalState
			_ = _lhs12
			var _lhs13 _dafny.Array = p0
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((p0), 0))
			_ = _lhs14
			var _lhs15 _dafny.Array = _118_v77
			_ = _lhs15
			var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_118_v77), 0))
			_ = _lhs16
			var _lhs17 _dafny.Array = _111_v72
			_ = _lhs17
			var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_111_v72), 0))
			_ = _lhs18
			_lhs12.F2 = _rhs22
			(_lhs13).ArraySet1(_rhs23, (_lhs14).Int())
			(_lhs15).ArraySet1(_rhs24, (_lhs16).Int())
			(_lhs17).ArraySet1(_rhs25, (_lhs18).Int())
			var _125_v80 *C1
			_ = _125_v80
			var _nw23 *C1 = New_C1_()
			_ = _nw23
			_nw23.Ctor__((Companion_D7_.Create_DC17_(false)).Dtor_cf25())
			_125_v80 = _nw23
			var _126_v81 _dafny.Map
			_ = _126_v81
			_126_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((p0), 0))).Int()).(bool), _125_v80)
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((p0), 0))
			_ = _index16
			(p0).ArraySet1(!(!((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_114_v73, _114_v73)).Cardinality())).Cmp(((_126_v81).Merge(_126_v81)).Cardinality()) > 0)), (_index16).Int())
			(globalState).F2 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
					return (_dafny.Zero).Minus(((_111_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_111_v72), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfUint32((_114_v73).Cardinality())))
				}
				return (_dafny.Zero).Minus(((_117_v76).Union(_117_v76)).Cardinality())
			})())
		} else {
			var _127___mcc_h2 D7 = _source1.Get_().(D7_DC18).Cf26
			_ = _127___mcc_h2
			var _128_cf26 D7 = _127___mcc_h2
			_ = _128_cf26
			var _129_v82 _dafny.MultiSet
			_ = _129_v82
			_129_v82 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("gwvsosvy"))
			var _130_v83 D1
			_ = _130_v83
			_130_v83 = Companion_D1_.Create_DC4_(_129_v82, p1)
			var _131_v84 _dafny.Sequence
			_ = _131_v84
			_131_v84 = _dafny.UnicodeSeqOfUtf8Bytes("qcm")
			var _pat_let_tv0 = _129_v82
			_ = _pat_let_tv0
			var _pat_let_tv1 = _131_v84
			_ = _pat_let_tv1
			var _pat_let_tv2 = _131_v84
			_ = _pat_let_tv2
			_130_v83 = func(_pat_let0_0 D1) D1 {
				return func(_132_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let1_0 _dafny.MultiSet) D1 {
						return func(_133_dt__update_hcf6_h0 _dafny.MultiSet) D1 {
							return Companion_D1_.Create_DC4_(_133_dt__update_hcf6_h0, (_132_dt__update__tmp_h0).Dtor_cf7())
						}(_pat_let1_0)
					}((_pat_let_tv0).Intersection(_dafny.MultiSetOf(_pat_let_tv1, _pat_let_tv2)))
				}(_pat_let0_0)
			}(_130_v83)
			(globalState).F2 = Companion_Default___.SafeDivisionInt(p1, p1)
			var _134_v85 *C3
			_ = _134_v85
			var _nw24 *C3 = New_C3_()
			_ = _nw24
			_nw24.Ctor__(_dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm9(_31_v0, globalState), _131_v84), p1)
			_134_v85 = _nw24
			_31_v0 = _31_v0
		}
		var _135_v86 _dafny.MultiSet
		_ = _135_v86
		_135_v86 = _dafny.MultiSetOf(false)
		var _136_v87 _dafny.MultiSet
		_ = _136_v87
		_136_v87 = _dafny.MultiSetOf((_135_v86).Cardinality(), p3)
		var _137_v88 _dafny.Sequence
		_ = _137_v88
		_137_v88 = _dafny.UnicodeSeqOfUtf8Bytes("fpbqp")
		var _138_v89 _dafny.Map
		_ = _138_v89
		_138_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v88, Companion_Default___.Fm0(_31_v0, _31_v0, _31_v0, _31_v0, globalState))
		var _139_v90 _dafny.Set
		_ = _139_v90
		_139_v90 = _dafny.SetOf((_dafny.IntOfInt64(383)).Cmp((func() _dafny.Int {
			if (_136_v87).Contains(p3) {
				return (_136_v87).Multiplicity(p3)
			}
			return p3
		})()) > 0, _31_v0, (func() bool {
			if (_138_v89).Contains(_137_v88) {
				return (_138_v89).Get(_137_v88).(bool)
			}
			return _31_v0
		})())
		_139_v90 = (_dafny.SetOf(_31_v0, _31_v0, false)).Intersection((_139_v90).Intersection(_139_v90))
		(globalState).F2 = (func() _dafny.Int {
			if (_136_v87).Contains(_dafny.IntOfInt64(530)) {
				return (_136_v87).Multiplicity(_dafny.IntOfInt64(530))
			}
			return (_dafny.MultiSetOf(_31_v0, _31_v0)).Cardinality()
		})()
		var _140_v91 *C1
		_ = _140_v91
		var _nw25 *C1 = New_C1_()
		_ = _nw25
		_nw25.Ctor__(_31_v0)
		_140_v91 = _nw25
		_140_v91 = _140_v91
		var _141_v92 _dafny.MultiSet
		_ = _141_v92
		_141_v92 = _dafny.MultiSetOf(_137_v88, _137_v88)
		var _142_v93 _dafny.Set
		_ = _142_v93
		_142_v93 = _dafny.SetOf((func() _dafny.Int {
			if (_141_v92).Contains(_137_v88) {
				return (_141_v92).Multiplicity(_137_v88)
			}
			return p1
		})(), _dafny.IntOfInt64(311), p1, p1)
		_142_v93 = _142_v93
	}
	var _143_v94 D1
	_ = _143_v94
	_143_v94 = Companion_D1_.Create_DC5_(_31_v0, Companion_Default___.Fm0(_31_v0, _31_v0, false, _31_v0, globalState))
	_31_v0 = !((_143_v94).Dtor_cf9())
	var _144_v95 _dafny.Sequence
	_ = _144_v95
	_144_v95 = _dafny.UnicodeSeqOfUtf8Bytes("jfta")
	var _145_v96 _dafny.Map
	_ = _145_v96
	_145_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v0, _144_v95)
	var _146_v97 _dafny.MultiSet
	_ = _146_v97
	_146_v97 = _dafny.MultiSetOf(p3)
	if (_dafny.MultiSetOf((_145_v96).Cardinality())).IsProperSubsetOf(_146_v97) {
		var _147_v98 _dafny.Array
		_ = _147_v98
		var _nw26 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
		_ = _nw26
		_147_v98 = _nw26
		_147_v98 = _147_v98
		var _148_v99 *C1
		_ = _148_v99
		var _nw27 *C1 = New_C1_()
		_ = _nw27
		_nw27.Ctor__(_31_v0)
		_148_v99 = _nw27
		var _149_v100 *C3
		_ = _149_v100
		var _nw28 *C3 = New_C3_()
		_ = _nw28
		_nw28.Ctor__(_31_v0, p1)
		_149_v100 = _nw28
		var _150_v101 _dafny.Sequence
		_ = _150_v101
		_150_v101 = _dafny.SeqOf(_149_v100)
		var _rhs26 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_32_v1, _32_v1)
		_ = _rhs26
		var _rhs27 bool = _148_v99.F14
		_ = _rhs27
		var _rhs28 _dafny.Sequence = _32_v1
		_ = _rhs28
		var _rhs29 _dafny.Int = p3
		_ = _rhs29
		var _rhs30 _dafny.Int = (p3).Plus(_dafny.IntOfUint32((_150_v101).Cardinality()))
		_ = _rhs30
		var _lhs19 *GlobalState = globalState
		_ = _lhs19
		var _lhs20 *GlobalState = globalState
		_ = _lhs20
		_32_v1 = _rhs26
		_31_v0 = _rhs27
		_32_v1 = _rhs28
		_lhs19.F2 = _rhs29
		_lhs20.F2 = _rhs30
		var _151_v102 _dafny.CodePoint
		_ = _151_v102
		_151_v102 = _dafny.CodePoint('v')
		var _152_v103 _dafny.Array
		_ = _152_v103
		var _nwElement0_4 _dafny.CodePoint = _151_v102
		_ = _nwElement0_4
		var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(21))
		_ = _nw29
		(_nw29).ArraySet1CodePoint(_nwElement0_4, 0)
		(_nw29).ArraySet1CodePoint(_151_v102, 1)
		(_nw29).ArraySet1CodePoint(_151_v102, 2)
		(_nw29).ArraySet1CodePoint(_151_v102, 3)
		(_nw29).ArraySet1CodePoint((func() _dafny.CodePoint {
			if Companion_Default___.Fm0(_148_v99.F14, (_149_v100).F11(), _148_v99.F14, _148_v99.F14, globalState) {
				return _151_v102
			}
			return _151_v102
		})(), 4)
		(_nw29).ArraySet1CodePoint(_151_v102, 5)
		(_nw29).ArraySet1CodePoint(_151_v102, 6)
		(_nw29).ArraySet1CodePoint((func() _dafny.CodePoint {
			if _31_v0 {
				return _151_v102
			}
			return _151_v102
		})(), 7)
		(_nw29).ArraySet1CodePoint(_151_v102, 8)
		(_nw29).ArraySet1CodePoint(_151_v102, 9)
		(_nw29).ArraySet1CodePoint(_151_v102, 10)
		(_nw29).ArraySet1CodePoint(_151_v102, 11)
		(_nw29).ArraySet1CodePoint((func() _dafny.CodePoint {
			if _31_v0 {
				return _151_v102
			}
			return _dafny.CodePoint('x')
		})(), 12)
		(_nw29).ArraySet1CodePoint(_151_v102, 13)
		(_nw29).ArraySet1CodePoint(_151_v102, 14)
		(_nw29).ArraySet1CodePoint(_151_v102, 15)
		(_nw29).ArraySet1CodePoint(_151_v102, 16)
		(_nw29).ArraySet1CodePoint(_151_v102, 17)
		(_nw29).ArraySet1CodePoint(_151_v102, 18)
		(_nw29).ArraySet1CodePoint(_151_v102, 19)
		(_nw29).ArraySet1CodePoint(_dafny.CodePoint('y'), 20)
		_152_v103 = _nw29
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(8))
		_ = _nw30
		_152_v103 = _nw30
		(_148_v99).F14 = (p1).Cmp(_dafny.IntOfUint32((_32_v1).Cardinality())) == 0
	} else {
		var _153_v104 *C0
		_ = _153_v104
		var _nw31 *C0 = New_C0_()
		_ = _nw31
		_nw31.Ctor__()
		_153_v104 = _nw31
		(globalState).F2 = p3
		var _154_v105 _dafny.Array
		_ = _154_v105
		var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
		_ = _nw32
		_154_v105 = _nw32
		var _155_v106 _dafny.Sequence
		_ = _155_v106
		_155_v106 = _dafny.SeqOf(_154_v105, _154_v105)
		var _156_v107 _dafny.Set
		_ = _156_v107
		_156_v107 = _dafny.SetOf(_31_v0)
		var _157_v108 _dafny.Array
		_ = _157_v108
		var _nwElement0_5 _dafny.Array = _154_v105
		_ = _nwElement0_5
		var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(19))
		_ = _nw33
		(_nw33).ArraySet1(_nwElement0_5, 0)
		(_nw33).ArraySet1(_154_v105, 1)
		(_nw33).ArraySet1(_154_v105, 2)
		(_nw33).ArraySet1(_154_v105, 3)
		(_nw33).ArraySet1(_154_v105, 4)
		(_nw33).ArraySet1(_154_v105, 5)
		(_nw33).ArraySet1((_155_v106).Select((Companion_Default___.SafeIndex((_156_v107).Cardinality(), _dafny.IntOfUint32((_155_v106).Cardinality()))).Uint32()).(_dafny.Array), 6)
		(_nw33).ArraySet1(_154_v105, 7)
		(_nw33).ArraySet1(_154_v105, 8)
		(_nw33).ArraySet1(_154_v105, 9)
		(_nw33).ArraySet1(_154_v105, 10)
		(_nw33).ArraySet1(_154_v105, 11)
		(_nw33).ArraySet1(_154_v105, 12)
		(_nw33).ArraySet1(_154_v105, 13)
		(_nw33).ArraySet1(_154_v105, 14)
		(_nw33).ArraySet1(_154_v105, 15)
		(_nw33).ArraySet1(_154_v105, 16)
		(_nw33).ArraySet1(_154_v105, 17)
		(_nw33).ArraySet1(_154_v105, 18)
		_157_v108 = _nw33
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_157_v108), 0))
		_ = _index17
		(_157_v108).ArraySet1(_154_v105, (_index17).Int())
		var _158_v109 _dafny.MultiSet
		_ = _158_v109
		_158_v109 = _dafny.MultiSetOf(_144_v95)
		var _159_v110 D1
		_ = _159_v110
		_159_v110 = Companion_D1_.Create_DC6_(Companion_D1_.Create_DC4_(_158_v109, p1))
		var _160_v111 D1
		_ = _160_v111
		_160_v111 = Companion_D1_.Create_DC3_(_dafny.CodePoint('q'))
		var _161_v112 _dafny.CodePoint
		_ = _161_v112
		_161_v112 = _dafny.CodePoint('a')
		var _pat_let_tv3 = _160_v111
		_ = _pat_let_tv3
		var _162_v113 _dafny.Map
		_ = _162_v113
		_162_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let2_0 D1) D1 {
			return func(_163_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let3_0 D1) D1 {
					return func(_164_dt__update_hcf10_h0 D1) D1 {
						return Companion_D1_.Create_DC6_(_164_dt__update_hcf10_h0)
					}(_pat_let3_0)
				}(_pat_let_tv3)
			}(_pat_let2_0)
		}(_159_v110), _161_v112)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_157_v108), 0))
		_ = _index18
		var _rhs31 _dafny.Array = _154_v105
		_ = _rhs31
		var _rhs32 _dafny.Array = _154_v105
		_ = _rhs32
		var _rhs33 _dafny.Map = ((Companion_Default___.Fm21(_31_v0, _31_v0, _dafny.IntOfInt64(-627), _31_v0, globalState)).Merge(_162_v113)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_159_v110, _dafny.CodePoint('w')))
		_ = _rhs33
		var _lhs21 *GlobalState = globalState
		_ = _lhs21
		var _lhs22 _dafny.Array = _157_v108
		_ = _lhs22
		var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_157_v108), 0))
		_ = _lhs23
		_lhs21.F10 = _rhs31
		(_lhs22).ArraySet1(_rhs32, (_lhs23).Int())
		_162_v113 = _rhs33
		_31_v0 = !(_31_v0) || (_31_v0)
		_31_v0 = (func() bool {
			if _31_v0 {
				return _31_v0
			}
			return !(_31_v0)
		})()
	}
	if !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(789))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_165_i5 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	})), _144_v95), _144_v95) {
		var _166_v114 _dafny.Map
		_ = _166_v114
		_166_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		var _167_v115 _dafny.Map
		_ = _167_v115
		_167_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _31_v0)
		_31_v0 = Companion_Default___.Fm0(_31_v0, (p3).Cmp((_dafny.Zero).Minus(p3)) >= 0, (p1).Cmp((_166_v114).Cardinality()) != 0, !((func() bool {
			if (_167_v115).Contains(_dafny.IntOfInt64(61)) {
				return (_167_v115).Get(_dafny.IntOfInt64(61)).(bool)
			}
			return _31_v0
		})()), globalState)
		_31_v0 = _31_v0
		var _168_v116 *C2
		_ = _168_v116
		var _nw34 *C2 = New_C2_()
		_ = _nw34
		_nw34.Ctor__(p3)
		_168_v116 = _nw34
		var _169_v117 D8
		_ = _169_v117
		_169_v117 = Companion_D8_.Create_DC19_(_168_v116)
		_169_v117 = (func() D8 {
			if (p3).Cmp(_168_v116.F13) >= 0 {
				return _169_v117
			}
			return _169_v117
		})()
		var _170_v118 _dafny.Array
		_ = _170_v118
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_5
		var _nw35 _dafny.Array
		_ = _nw35
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw35 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Sequence = (func(_171_v95 _dafny.Sequence, _172_v116 *C2) func(_dafny.Int) _dafny.Sequence {
				return func(_173_i6 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_171_v95, (Companion_Default___.SafeIndex(_172_v116.F13, _dafny.IntOfUint32((_171_v95).Cardinality()))).Uint32(), _dafny.CodePoint('o')), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(429))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg10 _dafny.Int) interface{} {
							return coer10(arg10)
						}
					}(func(_174_i7 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('v')
					})))
				}
			})(_144_v95, _168_v116)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw35 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw35).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw35).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_170_v118 = _nw35
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_170_v118), 0))
		_ = _index19
		(_170_v118).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_144_v95, Companion_Default___.Fm9(_31_v0, globalState)), (_index19).Int())
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_170_v118), 0))
		_ = _index20
		(_170_v118).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-589))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_175_i8 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		})), (_index20).Int())
		var _176_v119 _dafny.Set
		_ = _176_v119
		_176_v119 = _dafny.SetOf(_168_v116.F13)
		var _177_v120 _dafny.Sequence
		_ = _177_v120
		_177_v120 = _dafny.SeqOf(_176_v119)
		_31_v0 = (_176_v119).IsProperSubsetOf((_177_v120).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_168_v116.F13), _dafny.IntOfUint32((_177_v120).Cardinality()))).Uint32()).(_dafny.Set))
	} else {
		var _178_v121 _dafny.Array
		_ = _178_v121
		var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
		_ = _nw36
		_178_v121 = _nw36
		var _179_v122 _dafny.Map
		_ = _179_v122
		_179_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v0, p1)
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_178_v121), 0))
		_ = _index21
		(_178_v121).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(107))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_180_i9 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}))).Cardinality()))).Merge(_179_v122), (_index21).Int())
		var _181_v123 _dafny.Sequence
		_ = _181_v123
		_181_v123 = _dafny.SeqOf(p1)
		var _182_v124 _dafny.Set
		_ = _182_v124
		_182_v124 = _dafny.SetOf(_31_v0)
		var _183_v125 _dafny.Sequence
		_ = _183_v125
		_183_v125 = _dafny.SeqOf(_182_v124, _182_v124)
		var _184_v127 _dafny.Array
		_ = _184_v127
		var _nwElement0_6 _dafny.Int = p3
		_ = _nwElement0_6
		var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(17))
		_ = _nw37
		(_nw37).ArraySet1(_nwElement0_6, 0)
		(_nw37).ArraySet1(p3, 1)
		(_nw37).ArraySet1(p3, 2)
		(_nw37).ArraySet1(p3, 3)
		(_nw37).ArraySet1(_dafny.IntOfUint32((_181_v123).Cardinality()), 4)
		(_nw37).ArraySet1((_182_v124).Cardinality(), 5)
		(_nw37).ArraySet1(p3, 6)
		(_nw37).ArraySet1((_146_v97).Cardinality(), 7)
		(_nw37).ArraySet1((func() _dafny.Map {
			var _coll12 = _dafny.NewMapBuilder()
			_ = _coll12
			for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(175), _dafny.IntOfInt64(637))); ; {
				_compr_12, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _185_v126 _dafny.Int
				_185_v126 = interface{}(_compr_12).(_dafny.Int)
				if ((_dafny.IntOfInt64(175)).Cmp(_185_v126) <= 0) && ((_185_v126).Cmp(_dafny.IntOfInt64(637)) < 0) {
					_coll12.Add((_185_v126).Minus(p1), _31_v0)
				}
			}
			return _coll12.ToMap()
		}()).Cardinality(), 8)
		(_nw37).ArraySet1(p3, 9)
		(_nw37).ArraySet1((_146_v97).Cardinality(), 10)
		(_nw37).ArraySet1(p1, 11)
		(_nw37).ArraySet1((_dafny.Zero).Minus(p1), 12)
		(_nw37).ArraySet1(p3, 13)
		(_nw37).ArraySet1(p3, 14)
		(_nw37).ArraySet1(p1, 15)
		(_nw37).ArraySet1(p3, 16)
		_184_v127 = _nw37
		var _186_v128 _dafny.Sequence
		_ = _186_v128
		_186_v128 = _dafny.SeqOf(_184_v127, _184_v127)
		var _187_v129 D11
		_ = _187_v129
		_187_v129 = Companion_D11_.Create_DC26_(p3, (_186_v128).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_186_v128).Cardinality()))).Uint32()).(_dafny.Array), _31_v0)
		var _188_v130 _dafny.Array
		_ = _188_v130
		var _nwElement0_7 _dafny.Set = Companion_Default___.Fm15(p3, _dafny.IntOfUint32((_181_v123).Cardinality()), Companion_Default___.Fm0(Companion_Default___.Fm0(_31_v0, _31_v0, _31_v0, !(_31_v0), globalState), _31_v0, _31_v0, !(_31_v0), globalState), globalState)
		_ = _nwElement0_7
		var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(15))
		_ = _nw38
		(_nw38).ArraySet1(_nwElement0_7, 0)
		(_nw38).ArraySet1(_182_v124, 1)
		(_nw38).ArraySet1((_183_v125).Select((Companion_Default___.SafeIndex((_181_v123).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_181_v123).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_183_v125).Cardinality()))).Uint32()).(_dafny.Set), 2)
		(_nw38).ArraySet1(_182_v124, 3)
		(_nw38).ArraySet1((_182_v124).Intersection(_182_v124), 4)
		(_nw38).ArraySet1(_182_v124, 5)
		(_nw38).ArraySet1((func() _dafny.Set {
			if _31_v0 {
				return _182_v124
			}
			return Companion_Default___.Fm15(p1, p1, _31_v0, globalState)
		})(), 6)
		(_nw38).ArraySet1(_182_v124, 7)
		(_nw38).ArraySet1((_182_v124).Union(_dafny.SetOf(!(_31_v0), (_187_v129).Dtor_cf38())), 8)
		(_nw38).ArraySet1(_182_v124, 9)
		(_nw38).ArraySet1(Companion_Default___.Fm15(p1, p1, _31_v0, globalState), 10)
		(_nw38).ArraySet1(_182_v124, 11)
		(_nw38).ArraySet1(_dafny.SetOf(false), 12)
		(_nw38).ArraySet1(_182_v124, 13)
		(_nw38).ArraySet1(_182_v124, 14)
		_188_v130 = _nw38
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_188_v130), 0))
		_ = _index22
		(_188_v130).ArraySet1(_182_v124, (_index22).Int())
		var _189_v131 _dafny.Set
		_ = _189_v131
		_189_v131 = _dafny.SetOf(p1, p3)
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_178_v121), 0))
		_ = _index23
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_188_v130), 0))
		_ = _index24
		var _rhs34 _dafny.Map = _179_v122
		_ = _rhs34
		var _rhs35 _dafny.Set = _dafny.SetOf((_189_v131).IsDisjointFrom(_189_v131))
		_ = _rhs35
		var _lhs24 _dafny.Array = _178_v121
		_ = _lhs24
		var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_178_v121), 0))
		_ = _lhs25
		var _lhs26 _dafny.Array = _188_v130
		_ = _lhs26
		var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_188_v130), 0))
		_ = _lhs27
		(_lhs24).ArraySet1(_rhs34, (_lhs25).Int())
		(_lhs26).ArraySet1(_rhs35, (_lhs27).Int())
		var _190_v132 D0
		_ = _190_v132
		_190_v132 = Companion_D0_.Create_DC0_(_31_v0)
		_31_v0 = Companion_Default___.Fm0(!((_190_v132).Dtor_cf0()) || (_31_v0), Companion_Default___.Fm0(_31_v0, _31_v0, _31_v0, _31_v0, globalState), ((_dafny.Zero).Minus(p3)).Cmp(p3) < 0, Companion_Default___.Fm0(_31_v0, Companion_Default___.Fm0(_31_v0, _31_v0, _31_v0, _31_v0, globalState), _31_v0, _31_v0, globalState), globalState)
		_31_v0 = !(true)
		var _191_v133 _dafny.Array
		_ = _191_v133
		var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
		_ = _nw39
		_191_v133 = _nw39
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_191_v133), 0))
		_ = _index25
		(_191_v133).ArraySet1(_181_v123, (_index25).Int())
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_191_v133), 0))
		_ = _index26
		(_191_v133).ArraySet1(_181_v123, (_index26).Int())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_184_v127), 0))
		_ = _index27
		(_184_v127).ArraySet1(p3, (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_184_v127), 0))
		_ = _index28
		(_184_v127).ArraySet1((_dafny.Zero).Minus(p1), (_index28).Int())
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _192_v0 bool
	_ = _192_v0
	_192_v0 = true
	var _193_v1 _dafny.Set
	_ = _193_v1
	_193_v1 = _dafny.SetOf(!(!(_192_v0)))
	var _194_v2 _dafny.Set
	_ = _194_v2
	_194_v2 = _dafny.SetOf(_dafny.SetOf(_192_v0, !(!(_192_v0))), _193_v1)
	var _195_v3 _dafny.Int
	_ = _195_v3
	_195_v3 = _dafny.IntOfInt64(848)
	var _196_v4 _dafny.Sequence
	_ = _196_v4
	_196_v4 = _dafny.UnicodeSeqOfUtf8Bytes("bk")
	var _197_v5 _dafny.Map
	_ = _197_v5
	_197_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v0, _196_v4)
	var _198_v6 _dafny.Sequence
	_ = _198_v6
	_198_v6 = _dafny.SeqOf((_197_v5).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_195_v3, _195_v3)).Cardinality(), _195_v3)
	var _199_v7 _dafny.Sequence
	_ = _199_v7
	_199_v7 = _dafny.SeqOf(_193_v1, _193_v1, _193_v1)
	var _200_v8 _dafny.Array
	_ = _200_v8
	var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
	_ = _nw40
	_200_v8 = _nw40
	var _201_v9 _dafny.Map
	_ = _201_v9
	_201_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_199_v7).Select((Companion_Default___.SafeIndex(_195_v3, _dafny.IntOfUint32((_199_v7).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), _200_v8)
	var _202_v10 _dafny.MultiSet
	_ = _202_v10
	_202_v10 = _dafny.MultiSetOf(!(_192_v0))
	var _203_globalState *GlobalState
	_ = _203_globalState
	var _nw41 *GlobalState = New_GlobalState_()
	_ = _nw41
	_nw41.Ctor__(_dafny.IntOfInt64(245), false, _dafny.IntOfInt64(862), false, (_dafny.SetOf(_193_v1, _193_v1)).Difference(_194_v2), (func() _dafny.Map {
		if _192_v0 {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_195_v3, _192_v0)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_195_v3, !(_192_v0))
	})(), _198_v6, _198_v6, false, true, (func() _dafny.Array {
		if (_201_v9).Contains((func() _dafny.Int {
			if (_202_v10).Contains(_192_v0) {
				return (_202_v10).Multiplicity(_192_v0)
			}
			return _195_v3
		})()) {
			return (_201_v9).Get((func() _dafny.Int {
				if (_202_v10).Contains(_192_v0) {
					return (_202_v10).Multiplicity(_192_v0)
				}
				return _195_v3
			})()).(_dafny.Array)
		}
		return _200_v8
	})())
	_203_globalState = _nw41
	(_203_globalState).F2 = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-661), (_dafny.Zero).Minus(_195_v3))).Times(_195_v3)
	var _204_v11 _dafny.Map
	_ = _204_v11
	_204_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v0, (_dafny.IntOfInt64(758)).Cmp(_195_v3) >= 0)
	var _205_v12 _dafny.Map
	_ = _205_v12
	_205_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v0, _195_v3)
	_204_v11 = (_204_v11).Update((_195_v3).Cmp(_195_v3) <= 0, (_195_v3).Cmp((_205_v12).Cardinality()) == 0)
	var _206_v13 _dafny.Array
	_ = _206_v13
	var _len0_6 _dafny.Int = _dafny.IntOfInt64(15)
	_ = _len0_6
	var _nw42 _dafny.Array
	_ = _nw42
	if _len0_6.Cmp(_dafny.Zero) == 0 {
		_nw42 = _dafny.NewArray(_len0_6)
	} else {
		var _init6 func(_dafny.Int) bool = (func(_207_v0 bool) func(_dafny.Int) bool {
			return func(_208_i0 _dafny.Int) bool {
				return !(_207_v0)
			}
		})(_192_v0)
		_ = _init6
		var _element0_6 = _init6(_dafny.Zero)
		_ = _element0_6
		_nw42 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
		(_nw42).ArraySet1(_element0_6, 0)
		var _nativeLen0_6 = (_len0_6).Int()
		_ = _nativeLen0_6
		for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
			(_nw42).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
		}
	}
	_206_v13 = _nw42
	var _209_v14 _dafny.Array
	_ = _209_v14
	var _len0_7 _dafny.Int = _dafny.IntOfInt64(14)
	_ = _len0_7
	var _nw43 _dafny.Array
	_ = _nw43
	if _len0_7.Cmp(_dafny.Zero) == 0 {
		_nw43 = _dafny.NewArray(_len0_7)
	} else {
		var _init7 func(_dafny.Int) _dafny.Map = (func(_210_v11 _dafny.Map) func(_dafny.Int) _dafny.Map {
			return func(_211_i1 _dafny.Int) _dafny.Map {
				return _210_v11
			}
		})(_204_v11)
		_ = _init7
		var _element0_7 = _init7(_dafny.Zero)
		_ = _element0_7
		_nw43 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
		(_nw43).ArraySet1(_element0_7, 0)
		var _nativeLen0_7 = (_len0_7).Int()
		_ = _nativeLen0_7
		for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
			(_nw43).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
		}
	}
	_209_v14 = _nw43
	Companion_Default___.M0(_206_v13, _195_v3, _209_v14, _195_v3, _203_globalState)
	if !(_192_v0) {
		var _212_v15 D0
		_ = _212_v15
		_212_v15 = Companion_D0_.Create_DC0_(true)
		var _213_v16 _dafny.Sequence
		_ = _213_v16
		_213_v16 = _dafny.SeqOf(Companion_Default___.Fm0(_192_v0, _192_v0, _192_v0, (_212_v15).Dtor_cf0(), _203_globalState))
		Companion_Default___.M0(_206_v13, (_dafny.IntOfUint32((_213_v16).Cardinality())).Times(_195_v3), _209_v14, (_193_v1).Cardinality(), _203_globalState)
		_192_v0 = _192_v0
		var _214_v17 *C2
		_ = _214_v17
		var _nw44 *C2 = New_C2_()
		_ = _nw44
		_nw44.Ctor__(_195_v3)
		_214_v17 = _nw44
		_192_v0 = _192_v0
		_192_v0 = _192_v0
	} else {
		(_203_globalState).F2 = _195_v3
		_192_v0 = _192_v0
		var _215_v18 _dafny.CodePoint
		_ = _215_v18
		_215_v18 = _dafny.CodePoint('t')
		var _216_v19 _dafny.Sequence
		_ = _216_v19
		_216_v19 = _dafny.SeqOf(_192_v0, _192_v0)
		var _source2 D1 = Companion_Default___.Fm13(_215_v18, _dafny.MultiSetFromSeq(_216_v19), _203_globalState)
		_ = _source2
		if _source2.Is_DC4() {
			var _217___mcc_h0 _dafny.MultiSet = _source2.Get_().(D1_DC4).Cf6
			_ = _217___mcc_h0
			var _218___mcc_h1 _dafny.Int = _source2.Get_().(D1_DC4).Cf7
			_ = _218___mcc_h1
			var _219_cf7 _dafny.Int = _218___mcc_h1
			_ = _219_cf7
			var _220_cf6 _dafny.MultiSet = _217___mcc_h0
			_ = _220_cf6
			var _rhs36 bool = _192_v0
			_ = _rhs36
			var _rhs37 _dafny.Int = Companion_Default___.SafeDivisionInt(_195_v3, _219_cf7)
			_ = _rhs37
			var _rhs38 _dafny.Array = _206_v13
			_ = _rhs38
			var _rhs39 _dafny.Int = (_dafny.Zero).Minus(_219_cf7)
			_ = _rhs39
			var _lhs28 *GlobalState = _203_globalState
			_ = _lhs28
			_192_v0 = _rhs36
			_lhs28.F2 = _rhs37
			_206_v13 = _rhs38
			_195_v3 = _rhs39
			(_203_globalState).F2 = _219_cf7
			var _221_v20 *C1
			_ = _221_v20
			var _nw45 *C1 = New_C1_()
			_ = _nw45
			_nw45.Ctor__(true)
			_221_v20 = _nw45
			_221_v20 = _221_v20
			var _222_v21 _dafny.Sequence
			_ = _222_v21
			var _223_v22 D1
			_ = _223_v22
			var _224_v23 _dafny.Int
			_ = _224_v23
			var _out0 _dafny.Sequence
			_ = _out0
			var _out1 D1
			_ = _out1
			var _out2 _dafny.Int
			_ = _out2
			_out0, _out1, _out2 = (_221_v20).M4(_205_v12, _203_globalState)
			_222_v21 = _out0
			_223_v22 = _out1
			_224_v23 = _out2
		} else if _source2.Is_DC5() {
			var _225___mcc_h2 bool = _source2.Get_().(D1_DC5).Cf8
			_ = _225___mcc_h2
			var _226___mcc_h3 bool = _source2.Get_().(D1_DC5).Cf9
			_ = _226___mcc_h3
			var _227_cf9 bool = _226___mcc_h3
			_ = _227_cf9
			var _228_cf8 bool = _225___mcc_h2
			_ = _228_cf8
			var _229_v24 D4
			_ = _229_v24
			_229_v24 = Companion_D4_.Create_DC13_(_215_v18)
			var _230_v25 _dafny.Map
			_ = _230_v25
			_230_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_cf9, _229_v24)
			var _231_v26 _dafny.Set
			_ = _231_v26
			_231_v26 = _dafny.SetOf(_195_v3, _195_v3)
			var _232_v27 D3
			_ = _232_v27
			_232_v27 = Companion_D3_.Create_DC10_(_192_v0, _195_v3, _227_cf9, _206_v13, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_231_v26).Cardinality(), _227_cf9))
			_230_v25 = (_230_v25).Update((_232_v27).Dtor_cf16(), _229_v24)
			_204_v11 = (_204_v11).Update(_192_v0, _192_v0)
			_227_cf9 = _227_cf9
			var _233_v28 *C3
			_ = _233_v28
			var _nw46 *C3 = New_C3_()
			_ = _nw46
			_nw46.Ctor__(_227_cf9, ((_198_v6).Select((Companion_Default___.SafeIndex(_195_v3, _dafny.IntOfUint32((_198_v6).Cardinality()))).Uint32()).(_dafny.Int)).Times(_dafny.IntOfUint32((_216_v19).Cardinality())))
			_233_v28 = _nw46
		} else if _source2.Is_DC3() {
			var _234___mcc_h4 _dafny.CodePoint = _source2.Get_().(D1_DC3).Cf5
			_ = _234___mcc_h4
			var _235_cf5 _dafny.CodePoint = _234___mcc_h4
			_ = _235_cf5
			(_203_globalState).F2 = _195_v3
			_192_v0 = _192_v0
			(_203_globalState).F2 = Companion_Default___.SafeDivisionInt(_195_v3, _195_v3)
			_195_v3 = Companion_Default___.Fm2(!(_192_v0), _203_globalState)
		} else {
			var _236___mcc_h5 D1 = _source2.Get_().(D1_DC6).Cf10
			_ = _236___mcc_h5
			var _237_cf10 D1 = _236___mcc_h5
			_ = _237_cf10
			var _238_v29 D1
			_ = _238_v29
			_238_v29 = Companion_D1_.Create_DC3_(_215_v18)
			var _239_v30 _dafny.Sequence
			_ = _239_v30
			_239_v30 = _dafny.SeqOf(Companion_Default___.Fm5(_dafny.IntOfInt64(454), _dafny.IntOfInt64(47), _dafny.IntOfInt64(921), _216_v19, _203_globalState), _238_v29)
			var _240_v31 _dafny.Map
			_ = _240_v31
			_240_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mawdmqm"), _196_v4), _dafny.IntOfUint32((_239_v30).Cardinality()))
			_240_v31 = ((func() _dafny.Map {
				if true {
					return _240_v31
				}
				return _240_v31
			})()).Merge(_240_v31)
			Companion_Default___.M0(_206_v13, ((_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_216_v19)).Cardinality())).Times(_195_v3), _209_v14, _195_v3, _203_globalState)
			var _241_v32 _dafny.Set
			_ = _241_v32
			_241_v32 = _dafny.SetOf(_dafny.IntOfUint32((_216_v19).Cardinality()), _195_v3)
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_200_v8), 0))
			_ = _index29
			(_200_v8).ArraySet1((_241_v32).Cardinality(), (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_200_v8), 0))
			_ = _index30
			(_200_v8).ArraySet1(_195_v3, (_index30).Int())
			(_203_globalState).F2 = _195_v3
		}
		(_203_globalState).F2 = _195_v3
		_202_v10 = _202_v10
	}
	var _hi0 _dafny.Int = _195_v3
	_ = _hi0
	for _242_i2 := (_195_v3).Times(_195_v3); _242_i2.Cmp(_hi0) < 0; _242_i2 = _242_i2.Plus(_dafny.One) {
		var _243_v33 *C2
		_ = _243_v33
		var _nw47 *C2 = New_C2_()
		_ = _nw47
		_nw47.Ctor__(_195_v3)
		_243_v33 = _nw47
		var _244_v34 *C1
		_ = _244_v34
		var _nw48 *C1 = New_C1_()
		_ = _nw48
		_nw48.Ctor__(_192_v0)
		_244_v34 = _nw48
		var _245_v35 _dafny.Map
		_ = _245_v35
		_245_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_195_v3, _244_v34)
		var _246_v36 *C0
		_ = _246_v36
		var _nw49 *C0 = New_C0_()
		_ = _nw49
		_nw49.Ctor__()
		_246_v36 = _nw49
		var _247_v37 _dafny.Map
		_ = _247_v37
		_247_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_v35, _246_v36)
		var _248_v38 _dafny.Map
		_ = _248_v38
		_248_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() *C0 {
			if (_247_v37).Contains(_245_v35) {
				return (_247_v37).Get(_245_v35).(*C0)
			}
			return _246_v36
		})(), _243_v33)
		var _249_v39 _dafny.Sequence
		_ = _249_v39
		_249_v39 = _dafny.SeqOf(!(_244_v34.F14))
		var _250_v40 _dafny.Array
		_ = _250_v40
		var _nwElement0_8 *C2 = _243_v33
		_ = _nwElement0_8
		var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(25))
		_ = _nw50
		(_nw50).ArraySet1(_nwElement0_8, 0)
		(_nw50).ArraySet1(_243_v33, 1)
		(_nw50).ArraySet1(_243_v33, 2)
		(_nw50).ArraySet1(_243_v33, 3)
		(_nw50).ArraySet1(_243_v33, 4)
		(_nw50).ArraySet1(_243_v33, 5)
		(_nw50).ArraySet1(_243_v33, 6)
		(_nw50).ArraySet1(_243_v33, 7)
		(_nw50).ArraySet1(_243_v33, 8)
		(_nw50).ArraySet1(_243_v33, 9)
		(_nw50).ArraySet1(_243_v33, 10)
		(_nw50).ArraySet1(_243_v33, 11)
		(_nw50).ArraySet1(_243_v33, 12)
		(_nw50).ArraySet1((func() *C2 {
			if _192_v0 {
				return _243_v33
			}
			return _243_v33
		})(), 13)
		(_nw50).ArraySet1(_243_v33, 14)
		(_nw50).ArraySet1(_243_v33, 15)
		(_nw50).ArraySet1(_243_v33, 16)
		(_nw50).ArraySet1(_243_v33, 17)
		(_nw50).ArraySet1(_243_v33, 18)
		(_nw50).ArraySet1(_243_v33, 19)
		(_nw50).ArraySet1((func() *C2 {
			if (_248_v38).Contains(_246_v36) {
				return (_248_v38).Get(_246_v36).(*C2)
			}
			return _243_v33
		})(), 20)
		(_nw50).ArraySet1((func() *C2 {
			if (_249_v39).Select((Companion_Default___.SafeIndex(_243_v33.F13, _dafny.IntOfUint32((_249_v39).Cardinality()))).Uint32()).(bool) {
				return _243_v33
			}
			return _243_v33
		})(), 21)
		(_nw50).ArraySet1(_243_v33, 22)
		(_nw50).ArraySet1(_243_v33, 23)
		(_nw50).ArraySet1(_243_v33, 24)
		_250_v40 = _nw50
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_250_v40), 0))
		_ = _index31
		(_250_v40).ArraySet1(_243_v33, (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_250_v40), 0))
		_ = _index32
		var _rhs40 bool = !(((_195_v3).Times(_dafny.IntOfInt64(-807))).Cmp(_243_v33.F13) <= 0)
		_ = _rhs40
		var _rhs41 *C2 = _243_v33
		_ = _rhs41
		var _lhs29 _dafny.Array = _250_v40
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_250_v40), 0))
		_ = _lhs30
		_192_v0 = _rhs40
		(_lhs29).ArraySet1(_rhs41, (_lhs30).Int())
		_196_v4 = _dafny.Companion_Sequence_.Concatenate(_196_v4, _196_v4)
		var _251_v42 _dafny.MultiSet
		_ = _251_v42
		_251_v42 = _dafny.MultiSetOf(_242_i2, _243_v33.F13)
		var _252_v43 _dafny.Set
		_ = _252_v43
		_252_v43 = _dafny.SetOf((_251_v42).Cardinality())
		(_244_v34).F14 = !((func() _dafny.Map {
			var _coll13 = _dafny.NewMapBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_252_v43).Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _253_v41 _dafny.Int
				_253_v41 = interface{}(_compr_13).(_dafny.Int)
				if (_252_v43).Contains(_253_v41) {
					_coll13.Add(Companion_Default___.SafeModuloInt(_253_v41, _195_v3), _243_v33.F13)
				}
			}
			return _coll13.ToMap()
		}()).Update((_dafny.Zero).Minus(_243_v33.F13), (_252_v43).Cardinality())).Contains(_195_v3)
		var _254_v44 _dafny.MultiSet
		_ = _254_v44
		_254_v44 = _dafny.MultiSetOf(_200_v8, _200_v8, _200_v8, _200_v8)
		if (_254_v44).Contains(_200_v8) {
			(_203_globalState).F2 = (_dafny.Zero).Minus(((_198_v6).Select((Companion_Default___.SafeIndex(_243_v33.F13, _dafny.IntOfUint32((_198_v6).Cardinality()))).Uint32()).(_dafny.Int)).Times(_242_i2))
			_196_v4 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_196_v4, _196_v4), _196_v4)
			var _255_v45 _dafny.MultiSet
			_ = _255_v45
			_255_v45 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_196_v4, _dafny.UnicodeSeqOfUtf8Bytes("ux")), _196_v4, _dafny.Companion_Sequence_.Concatenate(_196_v4, _196_v4), _196_v4, _196_v4)
			_255_v45 = _255_v45
			var _256_v46 _dafny.Array
			_ = _256_v46
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_8
			var _nw51 _dafny.Array
			_ = _nw51
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw51 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Sequence = func(_257_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("k")
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw51 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw51).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw51).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_256_v46 = _nw51
			var _258_v47 _dafny.CodePoint
			_ = _258_v47
			_258_v47 = _dafny.CodePoint('d')
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_256_v46), 0))
			_ = _index33
			(_256_v46).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dfldjy"), _dafny.UnicodeSeqOfUtf8Bytes("e")), (Companion_Default___.SafeIndex(_243_v33.F13, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dfldjy"), _dafny.UnicodeSeqOfUtf8Bytes("e"))).Cardinality()))).Uint32(), _258_v47), (_index33).Int())
			var _259_v48 D2
			_ = _259_v48
			_259_v48 = Companion_D2_.Create_DC8_(_dafny.IntOfUint32((_249_v39).Cardinality()))
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_256_v46), 0))
			_ = _index34
			var _rhs42 _dafny.Int = (_243_v33.F13).Times((_259_v48).Dtor_cf12())
			_ = _rhs42
			var _rhs43 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("cqdmq")
			_ = _rhs43
			var _rhs44 _dafny.Int = _243_v33.F13
			_ = _rhs44
			var _rhs45 _dafny.Int = (_243_v33.F13).Plus(_243_v33.F13)
			_ = _rhs45
			var _lhs31 *C2 = _243_v33
			_ = _lhs31
			var _lhs32 _dafny.Array = _256_v46
			_ = _lhs32
			var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_256_v46), 0))
			_ = _lhs33
			var _lhs34 *GlobalState = _203_globalState
			_ = _lhs34
			var _lhs35 *GlobalState = _203_globalState
			_ = _lhs35
			_lhs31.F13 = _rhs42
			(_lhs32).ArraySet1(_rhs43, (_lhs33).Int())
			_lhs34.F2 = _rhs44
			_lhs35.F2 = _rhs45
			var _260_v49 _dafny.Map
			_ = _260_v49
			_260_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_205_v12, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_195_v3, _244_v34))
			_260_v49 = _260_v49
		} else {
			_204_v11 = _204_v11
			var _261_v50 _dafny.CodePoint
			_ = _261_v50
			_261_v50 = _dafny.CodePoint('h')
			var _262_v51 D1
			_ = _262_v51
			_262_v51 = Companion_D1_.Create_DC3_(_261_v50)
			var _263_v52 *C3
			_ = _263_v52
			var _nw52 *C3 = New_C3_()
			_ = _nw52
			_nw52.Ctor__(!((_246_v36).Fm7(false, _262_v51, _196_v4, _244_v34.F14, _203_globalState)), _195_v3)
			_263_v52 = _nw52
			_263_v52 = _263_v52
			var _264_v53 _dafny.Array
			_ = _264_v53
			var _nw53 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.One)
			_ = _nw53
			_264_v53 = _nw53
			var _265_v54 D4
			_ = _265_v54
			_265_v54 = Companion_D4_.Create_DC12_(_193_v1)
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_264_v53), 0))
			_ = _index35
			(_264_v53).ArraySet1(_265_v54, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_264_v53), 0))
			_ = _index36
			(_264_v53).ArraySet1(_265_v54, (_index36).Int())
			var _266_v55 D1
			_ = _266_v55
			_266_v55 = Companion_D1_.Create_DC5_(_192_v0, _244_v34.F14)
			var _pat_let_tv4 = _244_v34
			_ = _pat_let_tv4
			_192_v0 = (func(_pat_let4_0 D1) D1 {
				return func(_267_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let5_0 bool) D1 {
						return func(_268_dt__update_hcf9_h0 bool) D1 {
							return Companion_D1_.Create_DC5_((_267_dt__update__tmp_h0).Dtor_cf8(), _268_dt__update_hcf9_h0)
						}(_pat_let5_0)
					}(_pat_let_tv4.F14)
				}(_pat_let4_0)
			}(_266_v55)).Dtor_cf9()
			(_244_v34).F14 = (_244_v34.F14) == (_244_v34.F14)
		}
	}
	var _269_v56 _dafny.Map
	_ = _269_v56
	_269_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("sharml"), _195_v3)
	var _270_v57 _dafny.CodePoint
	_ = _270_v57
	_270_v57 = _dafny.CodePoint('c')
	var _271_v58 *C3
	_ = _271_v58
	var _nw54 *C3 = New_C3_()
	_ = _nw54
	_nw54.Ctor__(_192_v0, _195_v3)
	_271_v58 = _nw54
	var _272_v59 _dafny.Map
	_ = _272_v59
	_272_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_270_v57, _271_v58)
	var _273_v60 _dafny.Map
	_ = _273_v60
	_273_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm4(_203_globalState), (func() *C3 {
		if (_272_v59).Contains(_270_v57) {
			return (_272_v59).Get(_270_v57).(*C3)
		}
		return _271_v58
	})())
	var _274_v61 _dafny.Map
	_ = _274_v61
	_274_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_273_v60, _200_v8)
	var _275_v62 _dafny.Sequence
	_ = _275_v62
	_275_v62 = _dafny.SeqOf(((_269_v56).Cardinality()).Cmp((_274_v61).Cardinality()) > 0)
	_275_v62 = _275_v62
	var _276_v63 bool
	_ = _276_v63
	var _277_v64 _dafny.Map
	_ = _277_v64
	var _278_v65 bool
	_ = _278_v65
	var _out3 bool
	_ = _out3
	var _out4 _dafny.Map
	_ = _out4
	var _out5 bool
	_ = _out5
	_out3, _out4, _out5 = (_271_v58).M2(_192_v0, (_271_v58).F11(), _203_globalState)
	_276_v63 = _out3
	_277_v64 = _out4
	_278_v65 = _out5
	var _279_v66 _dafny.Map
	_ = _279_v66
	_279_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_((_271_v58).F12(), (_271_v58).F12(), (_271_v58).F12()), (_271_v58).F11())
	var _280_v67 _dafny.Map
	_ = _280_v67
	_280_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_271_v58).F12(), _195_v3)
	var _281_v68 _dafny.Map
	_ = _281_v68
	_281_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v0, _280_v67)
	var _282_v69 _dafny.Map
	_ = _282_v69
	_282_v69 = _201_v9
	var _283_v70 _dafny.MultiSet
	_ = _283_v70
	_283_v70 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_282_v69, (_271_v58).F12()))
	var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
	_ = _index37
	(_200_v8).ArraySet1((_283_v70).Cardinality(), (_index37).Int())
	var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
	_ = _index38
	var _rhs46 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_271_v58).F12(), ((_204_v11).Merge(_204_v11)).Cardinality()))
	_ = _rhs46
	var _rhs47 _dafny.Map = _279_v66
	_ = _rhs47
	var _rhs48 _dafny.Map = Companion_Default___.Fm14(_dafny.Companion_Sequence_.Concatenate(_196_v4, _196_v4), _203_globalState)
	_ = _rhs48
	var _rhs49 _dafny.Int = (_dafny.IntOfUint32((_196_v4).Cardinality())).Plus((_271_v58).F12())
	_ = _rhs49
	var _lhs36 _dafny.Array = _200_v8
	_ = _lhs36
	var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
	_ = _lhs37
	_195_v3 = _rhs46
	_279_v66 = _rhs47
	_281_v68 = _rhs48
	(_lhs36).ArraySet1(_rhs49, (_lhs37).Int())
	var _284_i4 _dafny.Int
	_ = _284_i4
	_284_i4 = _dafny.Zero
	{
		for ((_271_v58).F12()).Cmp((_202_v10).Cardinality()) >= 0 {
			{
				if (_284_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_284_i4 = (_284_i4).Plus(_dafny.One)
				var _285_v71 _dafny.Map
				_ = _285_v71
				_285_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int), (_271_v58).F12()), _196_v4)
				var _286_v72 _dafny.Sequence
				_ = _286_v72
				_286_v72 = _dafny.SeqOf(_271_v58)
				var _287_v73 _dafny.Sequence
				_ = _287_v73
				_287_v73 = _286_v72
				_285_v71 = (_285_v71).Update(_dafny.IntOfUint32((_287_v73).Cardinality()), _dafny.SeqOf(Companion_Default___.Fm6((_271_v58).F12(), _192_v0, _278_v65, _dafny.CodePoint('n'), _203_globalState)))
				_192_v0 = ((_202_v10).Intersection(_dafny.MultiSetOf(true))).Equals(_202_v10)
				var _288_v74 D4
				_ = _288_v74
				_288_v74 = Companion_D4_.Create_DC13_((func() _dafny.CodePoint {
					if (_271_v58).F11() {
						return _270_v57
					}
					return _270_v57
				})())
				_288_v74 = _288_v74
				var _289_v75 D1
				_ = _289_v75
				_289_v75 = Companion_D1_.Create_DC5_((_271_v58).F11(), _192_v0)
				var _290_v76 _dafny.Sequence
				_ = _290_v76
				_290_v76 = _dafny.SeqOf(_289_v75)
				_290_v76 = _290_v76
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _291_v77 *C3
	_ = _291_v77
	var _nw55 *C3 = New_C3_()
	_ = _nw55
	_nw55.Ctor__(!((_271_v58).F11()), _dafny.IntOfInt64(651))
	_291_v77 = _nw55
	var _292_v78 _dafny.Map
	_ = _292_v78
	_292_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_291_v77, (_291_v77).F12())
	_292_v78 = (_292_v78).Update(_291_v77, (_291_v77).F12())
	_195_v3 = (_271_v58).F12()
	(_291_v77).M1(!(_276_v63), _203_globalState)
	var _293_v79 _dafny.Array
	_ = _293_v79
	var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
	_ = _nw56
	_293_v79 = _nw56
	var _294_v80 _dafny.MultiSet
	_ = _294_v80
	_294_v80 = _dafny.MultiSetOf(_293_v79, _293_v79, _293_v79)
	_294_v80 = _294_v80
	var _source3 D1 = Companion_Default___.Fm13(_270_v57, _dafny.MultiSetOf((_271_v58).F11(), _278_v65, !((_291_v77).F11())), _203_globalState)
	_ = _source3
	if _source3.Is_DC4() {
		var _295___mcc_h6 _dafny.MultiSet = _source3.Get_().(D1_DC4).Cf6
		_ = _295___mcc_h6
		var _296___mcc_h7 _dafny.Int = _source3.Get_().(D1_DC4).Cf7
		_ = _296___mcc_h7
		var _297_cf7 _dafny.Int = _296___mcc_h7
		_ = _297_cf7
		var _298_cf6 _dafny.MultiSet = _295___mcc_h6
		_ = _298_cf6
		var _299_v81 D1
		_ = _299_v81
		_299_v81 = Companion_D1_.Create_DC3_(_270_v57)
		var _source4 D1 = _299_v81
		_ = _source4
		if _source4.Is_DC4() {
			var _300___mcc_h12 _dafny.MultiSet = _source4.Get_().(D1_DC4).Cf6
			_ = _300___mcc_h12
			var _301___mcc_h13 _dafny.Int = _source4.Get_().(D1_DC4).Cf7
			_ = _301___mcc_h13
			var _302_cf7 _dafny.Int = _301___mcc_h13
			_ = _302_cf7
			var _303_cf6 _dafny.MultiSet = _300___mcc_h12
			_ = _303_cf6
			var _304_v82 _dafny.Set
			_ = _304_v82
			_304_v82 = _dafny.SetOf(_dafny.IntOfInt64(832), (_291_v77).F12(), (_291_v77).F12())
			var _305_v83 *C3
			_ = _305_v83
			var _nw57 *C3 = New_C3_()
			_ = _nw57
			_nw57.Ctor__(!(_192_v0), (_304_v82).Cardinality())
			_305_v83 = _nw57
			var _306_v84 _dafny.Map
			_ = _306_v84
			_306_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nuwybql")).Cardinality()), ((_305_v83).F11()) || (Companion_Default___.Fm0(false, (_271_v58).F11(), _192_v0, (_291_v77).F11(), _203_globalState)))
			var _307_v85 _dafny.MultiSet
			_ = _307_v85
			_307_v85 = _dafny.MultiSetOf(_204_v11)
			_306_v84 = (_306_v84).Update((_dafny.Zero).Minus(((_307_v85).Difference(_dafny.MultiSetOf(_204_v11, _204_v11, _204_v11))).Cardinality()), (_291_v77).F11())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _index39
			(_200_v8).ArraySet1((_271_v58).F12(), (_index39).Int())
			_276_v63 = _278_v65
		} else if _source4.Is_DC5() {
			var _308___mcc_h14 bool = _source4.Get_().(D1_DC5).Cf8
			_ = _308___mcc_h14
			var _309___mcc_h15 bool = _source4.Get_().(D1_DC5).Cf9
			_ = _309___mcc_h15
			var _310_cf9 bool = _309___mcc_h15
			_ = _310_cf9
			var _311_cf8 bool = _308___mcc_h14
			_ = _311_cf8
			var _312_v86 _dafny.MultiSet
			_ = _312_v86
			_312_v86 = _dafny.MultiSetOf((_271_v58).Fm1(_196_v4, (_271_v58).F12(), _203_globalState), (_291_v77).F12(), (_291_v77).F12(), _dafny.IntOfUint32((_196_v4).Cardinality()), (_197_v5).Cardinality())
			Companion_Default___.M0(_206_v13, ((func() _dafny.Int {
				if (_312_v86).Contains(_dafny.IntOfInt64(-408)) {
					return (_312_v86).Multiplicity(_dafny.IntOfInt64(-408))
				}
				return (_291_v77).F12()
			})()).Plus((_291_v77).F12()), _209_v14, _195_v3, _203_globalState)
			_196_v4 = _dafny.Companion_Sequence_.Concatenate(_196_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(889))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_313_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			})))
			var _314_v87 bool
			_ = _314_v87
			var _315_v88 _dafny.Map
			_ = _315_v88
			var _316_v89 bool
			_ = _316_v89
			var _out6 bool
			_ = _out6
			var _out7 _dafny.Map
			_ = _out7
			var _out8 bool
			_ = _out8
			_out6, _out7, _out8 = (_271_v58).M2((_195_v3).Cmp(_297_cf7) != 0, _276_v63, _203_globalState)
			_314_v87 = _out6
			_315_v88 = _out7
			_316_v89 = _out8
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _index40
			(_200_v8).ArraySet1((_dafny.Zero).Minus((_291_v77).F12()), (_index40).Int())
		} else if _source4.Is_DC3() {
			var _317___mcc_h16 _dafny.CodePoint = _source4.Get_().(D1_DC3).Cf5
			_ = _317___mcc_h16
			var _318_cf5 _dafny.CodePoint = _317___mcc_h16
			_ = _318_cf5
			_276_v63 = !(!(!((func() bool {
				if false {
					return (_271_v58).F11()
				}
				return (_271_v58).F11()
			})()) || (!((_271_v58).F11()))))
			_297_cf7 = (_297_cf7).Minus((_dafny.IntOfInt64(10)).Minus((func() _dafny.Set {
				var _coll14 = _dafny.NewBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(393), _dafny.IntOfInt64(611))); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _319_v90 _dafny.Int
					_319_v90 = interface{}(_compr_14).(_dafny.Int)
					if ((_dafny.IntOfInt64(393)).Cmp(_319_v90) <= 0) && ((_319_v90).Cmp(_dafny.IntOfInt64(611)) < 0) {
						_coll14.Add(Companion_Default___.SafeDivisionInt(_319_v90, (_271_v58).F12()))
					}
				}
				return _coll14.ToSet()
			}()).Cardinality()))
			_204_v11 = (_204_v11).Update((_271_v58).F11(), _276_v63)
			var _320_v91 *C0
			_ = _320_v91
			var _nw58 *C0 = New_C0_()
			_ = _nw58
			_nw58.Ctor__()
			_320_v91 = _nw58
		} else {
			var _321___mcc_h17 D1 = _source4.Get_().(D1_DC6).Cf10
			_ = _321___mcc_h17
			var _322_cf10 D1 = _321___mcc_h17
			_ = _322_cf10
			var _323_v92 *C1
			_ = _323_v92
			var _nw59 *C1 = New_C1_()
			_ = _nw59
			_nw59.Ctor__((_291_v77).F11())
			_323_v92 = _nw59
			var _324_v93 _dafny.Set
			_ = _324_v93
			_324_v93 = _dafny.SetOf(_323_v92, _323_v92, _323_v92)
			(_271_v58).M1((_324_v93).IsProperSubsetOf(_324_v93), _203_globalState)
			var _325_v94 _dafny.Array
			_ = _325_v94
			var _nw60 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw60
			_325_v94 = _nw60
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((_325_v94), 0))
			_ = _index41
			(_325_v94).ArraySet1(_196_v4, (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((_325_v94), 0))
			_ = _index42
			(_325_v94).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("adpwq"), (_index42).Int())
			var _326_v95 _dafny.MultiSet
			_ = _326_v95
			_326_v95 = _dafny.MultiSetOf(_271_v58)
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _index43
			var _rhs50 _dafny.Int = _195_v3
			_ = _rhs50
			var _rhs51 _dafny.MultiSet = _326_v95
			_ = _rhs51
			var _rhs52 bool = (Companion_Default___.SafeModuloInt((_291_v77).F12(), (_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int))).Cmp(((_291_v77).F12()).Times(_195_v3)) == 0
			_ = _rhs52
			var _lhs38 _dafny.Array = _200_v8
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _lhs39
			(_lhs38).ArraySet1(_rhs50, (_lhs39).Int())
			_326_v95 = _rhs51
			_276_v63 = _rhs52
			_195_v3 = (_271_v58).F12()
		}
		_206_v13 = _206_v13
		if _192_v0 {
			(_203_globalState).F2 = (func() _dafny.Int {
				if (_271_v58).F11() {
					return (_dafny.Zero).Minus((_291_v77).F12())
				}
				return _dafny.IntOfInt64(-201)
			})()
			var _327_v96 D7
			_ = _327_v96
			_327_v96 = Companion_D7_.Create_DC17_(_192_v0)
			var _328_v97 _dafny.Set
			_ = _328_v97
			_328_v97 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(!(_276_v63), (_327_v96).Dtor_cf25())).Cardinality()))
			var _rhs53 bool = (_dafny.SetOf(_297_cf7)).IsSubsetOf(_328_v97)
			_ = _rhs53
			var _rhs54 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_291_v77).F11() {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-994))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg14 _dafny.Int) interface{} {
							return coer14(arg14)
						}
					}((func(_329_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_330_i6 _dafny.Int) _dafny.CodePoint {
							return _329_v57
						}
					})(_270_v57)))
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-963))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}((func(_331_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_332_i7 _dafny.Int) _dafny.CodePoint {
						return _331_v57
					}
				})(_270_v57)))
			})()).Cardinality())
			_ = _rhs54
			var _lhs40 *GlobalState = _203_globalState
			_ = _lhs40
			_192_v0 = _rhs53
			_lhs40.F2 = _rhs54
			var _333_v98 D0
			_ = _333_v98
			_333_v98 = Companion_D0_.Create_DC1_((_271_v58).F12(), _dafny.IntOfUint32((_198_v6).Cardinality()), (_291_v77).F12())
			var _334_v99 *C1
			_ = _334_v99
			var _nw61 *C1 = New_C1_()
			_ = _nw61
			_nw61.Ctor__(((_333_v98).Dtor_cf1()).Cmp(_297_cf7) == 0)
			_334_v99 = _nw61
			var _335_v100 _dafny.Array
			_ = _335_v100
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_9
			var _nw62 _dafny.Array
			_ = _nw62
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw62 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Sequence = func(_336_i8 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("n")
				}
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw62 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw62).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw62).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_335_v100 = _nw62
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_335_v100), 0))
			_ = _index44
			(_335_v100).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_196_v4, _196_v4), (_index44).Int())
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_335_v100), 0))
			_ = _index45
			(_335_v100).ArraySet1(_196_v4, (_index45).Int())
			var _337_v101 *C1
			_ = _337_v101
			var _nw63 *C1 = New_C1_()
			_ = _nw63
			_nw63.Ctor__((_291_v77).F11())
			_337_v101 = _nw63
		} else {
			_275_v62 = _dafny.Companion_Sequence_.Concatenate(_275_v62, _dafny.Companion_Sequence_.Concatenate(_275_v62, _275_v62))
			Companion_Default___.M0(_206_v13, (_271_v58).F12(), _209_v14, (_291_v77).F12(), _203_globalState)
			var _338_v102 _dafny.Map
			_ = _338_v102
			_338_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_292_v78, (_291_v77).Fm1(_196_v4, _195_v3, _203_globalState))
			_297_cf7 = (func() _dafny.Int {
				if (!((_271_v58).F11())) == (Companion_Default___.Fm0(true, (_271_v58).F11(), (_271_v58).F11(), (_271_v58).F11(), _203_globalState)) {
					return (_271_v58).F12()
				}
				return (func() _dafny.Int {
					if (_338_v102).Contains(_292_v78) {
						return (_338_v102).Get(_292_v78).(_dafny.Int)
					}
					return (_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int)
				})()
			})()
			(_203_globalState).F2 = (_297_cf7).Minus((_291_v77).F12())
			_206_v13 = (func() _dafny.Array {
				if (_291_v77).F11() {
					return _206_v13
				}
				return _206_v13
			})()
		}
		var _339_v103 _dafny.MultiSet
		_ = _339_v103
		_339_v103 = _dafny.MultiSetOf(_200_v8, (func() _dafny.Array {
			if (_291_v77).F11() {
				return _200_v8
			}
			return _200_v8
		})())
		var _340_v104 _dafny.Sequence
		_ = _340_v104
		_340_v104 = _dafny.SeqOf(_291_v77, _271_v58)
		var _341_v105 _dafny.Sequence
		_ = _341_v105
		_341_v105 = _340_v104
		var _342_v106 _dafny.Array
		_ = _342_v106
		var _nwElement0_9 _dafny.Sequence = _340_v104
		_ = _nwElement0_9
		var _nw64 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(9))
		_ = _nw64
		(_nw64).ArraySet1(_nwElement0_9, 0)
		(_nw64).ArraySet1(_341_v105, 1)
		(_nw64).ArraySet1(_341_v105, 2)
		(_nw64).ArraySet1(_340_v104, 3)
		(_nw64).ArraySet1(_340_v104, 4)
		(_nw64).ArraySet1(_341_v105, 5)
		(_nw64).ArraySet1(_341_v105, 6)
		(_nw64).ArraySet1(_dafny.Companion_Sequence_.Update(_340_v104, (Companion_Default___.SafeIndex((_291_v77).F12(), _dafny.IntOfUint32((_340_v104).Cardinality()))).Uint32(), _271_v58), 7)
		(_nw64).ArraySet1(_341_v105, 8)
		_342_v106 = _nw64
		var _343_v107 _dafny.Map
		_ = _343_v107
		_343_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_v106, _dafny.MultiSetOf(_200_v8))
		_339_v103 = (func() _dafny.MultiSet {
			if (_343_v107).Contains(_342_v106) {
				return (_343_v107).Get(_342_v106).(_dafny.MultiSet)
			}
			return _339_v103
		})()
	} else if _source3.Is_DC5() {
		var _344___mcc_h8 bool = _source3.Get_().(D1_DC5).Cf8
		_ = _344___mcc_h8
		var _345___mcc_h9 bool = _source3.Get_().(D1_DC5).Cf9
		_ = _345___mcc_h9
		var _346_cf9 bool = _345___mcc_h9
		_ = _346_cf9
		var _347_cf8 bool = _344___mcc_h8
		_ = _347_cf8
		var _348_v108 _dafny.MultiSet
		_ = _348_v108
		_348_v108 = _dafny.MultiSetOf(Companion_Default___.Fm9((_291_v77).F11(), _203_globalState))
		var _349_v109 D1
		_ = _349_v109
		_349_v109 = Companion_D1_.Create_DC4_(_348_v108, (_291_v77).F12())
		_349_v109 = _349_v109
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_206_v13), 0))
		_ = _index46
		(_206_v13).ArraySet1((_271_v58).F11(), (_index46).Int())
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_206_v13), 0))
		_ = _index47
		(_206_v13).ArraySet1(false, (_index47).Int())
		var _350_v110 _dafny.Map
		_ = _350_v110
		_350_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_347_cf8, _275_v62)
		var _351_v111 _dafny.Map
		_ = _351_v111
		_351_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
			if (_350_v110).Contains(_346_cf9) {
				return (_350_v110).Get(_346_cf9).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Update(_275_v62, (Companion_Default___.SafeIndex(_195_v3, _dafny.IntOfUint32((_275_v62).Cardinality()))).Uint32(), _276_v63)
		})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_275_v62, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-519), _dafny.IntOfUint32((_275_v62).Cardinality()))).Uint32(), !(!(_347_cf8)))).Cardinality()))
		var _352_v112 _dafny.MultiSet
		_ = _352_v112
		_352_v112 = _dafny.MultiSetOf((_351_v111).Cardinality())
		_195_v3 = (_dafny.SetOf(((_271_v58).F12()).Minus((_352_v112).Cardinality()))).Cardinality()
		var _353_v113 D1
		_ = _353_v113
		_353_v113 = Companion_D1_.Create_DC5_(_346_cf9, _347_cf8)
		if (_353_v113).Dtor_cf9() {
			_196_v4 = _dafny.Companion_Sequence_.Concatenate(_196_v4, _196_v4)
			var _354_v114 *C3
			_ = _354_v114
			var _nw65 *C3 = New_C3_()
			_ = _nw65
			_nw65.Ctor__((_291_v77).F11(), Companion_Default___.SafeDivisionInt((Companion_Default___.Fm15((_291_v77).F12(), (_291_v77).F12(), _192_v0, _203_globalState)).Cardinality(), _195_v3))
			_354_v114 = _nw65
			var _355_v115 _dafny.Sequence
			_ = _355_v115
			_355_v115 = _dafny.SeqOf(_200_v8, _200_v8)
			var _356_v116 _dafny.Map
			_ = _356_v116
			_356_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_270_v57, _200_v8)
			var _357_v117 _dafny.Array
			_ = _357_v117
			var _nwElement0_10 _dafny.Array = (func() _dafny.Array {
				if (_201_v9).Contains(_dafny.IntOfInt64(83)) {
					return (_201_v9).Get(_dafny.IntOfInt64(83)).(_dafny.Array)
				}
				return (_355_v115).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-561), _dafny.IntOfUint32((_355_v115).Cardinality()))).Uint32()).(_dafny.Array)
			})()
			_ = _nwElement0_10
			var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(20))
			_ = _nw66
			(_nw66).ArraySet1(_nwElement0_10, 0)
			(_nw66).ArraySet1(_200_v8, 1)
			(_nw66).ArraySet1((func() _dafny.Array {
				if (_291_v77).F11() {
					return _200_v8
				}
				return _200_v8
			})(), 2)
			(_nw66).ArraySet1((func() _dafny.Array {
				if (_356_v116).Contains(_270_v57) {
					return (_356_v116).Get(_270_v57).(_dafny.Array)
				}
				return _200_v8
			})(), 3)
			(_nw66).ArraySet1(_200_v8, 4)
			(_nw66).ArraySet1(_200_v8, 5)
			(_nw66).ArraySet1(_200_v8, 6)
			(_nw66).ArraySet1(_200_v8, 7)
			(_nw66).ArraySet1(_200_v8, 8)
			(_nw66).ArraySet1(_200_v8, 9)
			(_nw66).ArraySet1(_200_v8, 10)
			(_nw66).ArraySet1(_200_v8, 11)
			(_nw66).ArraySet1(_200_v8, 12)
			(_nw66).ArraySet1(_200_v8, 13)
			(_nw66).ArraySet1(_200_v8, 14)
			(_nw66).ArraySet1(_200_v8, 15)
			(_nw66).ArraySet1(_200_v8, 16)
			(_nw66).ArraySet1(_200_v8, 17)
			(_nw66).ArraySet1(_200_v8, 18)
			(_nw66).ArraySet1(_200_v8, 19)
			_357_v117 = _nw66
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_357_v117), 0))
			_ = _index48
			(_357_v117).ArraySet1(_200_v8, (_index48).Int())
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_357_v117), 0))
			_ = _index49
			(_357_v117).ArraySet1(_200_v8, (_index49).Int())
			_354_v114 = _271_v58
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _index50
			(_200_v8).ArraySet1(((_dafny.Zero).Minus((_354_v114).F12())).Minus(_dafny.IntOfInt64(500)), (_index50).Int())
		} else {
			var _358_v118 _dafny.Array
			_ = _358_v118
			var _nw67 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw67
			_358_v118 = _nw67
			Companion_Default___.M0(_358_v118, Companion_Default___.SafeDivisionInt(_195_v3, (_271_v58).F12()), _209_v14, (_271_v58).F12(), _203_globalState)
			var _359_v119 bool
			_ = _359_v119
			var _360_v120 _dafny.Map
			_ = _360_v120
			var _361_v121 bool
			_ = _361_v121
			var _out9 bool
			_ = _out9
			var _out10 _dafny.Map
			_ = _out10
			var _out11 bool
			_ = _out11
			_out9, _out10, _out11 = (_271_v58).M2(_347_cf8, ((_271_v58).F12()).Cmp((_291_v77).F12()) <= 0, _203_globalState)
			_359_v119 = _out9
			_360_v120 = _out10
			_361_v121 = _out11
			var _rhs55 _dafny.Int = (_198_v6).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if _192_v0 {
					return (_291_v77).F12()
				}
				return _195_v3
			})(), _dafny.IntOfUint32((_198_v6).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs55
			var _rhs56 bool = (_275_v62).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.IntOfUint32((_275_v62).Cardinality()))).Uint32()).(bool)
			_ = _rhs56
			var _rhs57 bool = _dafny.Companion_Sequence_.Contains(_196_v4, _270_v57)
			_ = _rhs57
			var _lhs41 *GlobalState = _203_globalState
			_ = _lhs41
			_lhs41.F2 = _rhs55
			_359_v119 = _rhs56
			_276_v63 = _rhs57
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_206_v13), 0))
			_ = _index51
			var _rhs58 _dafny.Int = (func() _dafny.Int {
				if (_202_v10).Contains(_347_cf8) {
					return (_202_v10).Multiplicity(_347_cf8)
				}
				return (_291_v77).F12()
			})()
			_ = _rhs58
			var _rhs59 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_275_v62, Companion_Default___.Fm16(_dafny.IntOfInt64(345), (_353_v113).Dtor_cf9(), _203_globalState))
			_ = _rhs59
			var _rhs60 bool = true
			_ = _rhs60
			var _rhs61 _dafny.Array = (func() _dafny.Array {
				if !(_278_v65) {
					return _200_v8
				}
				return _200_v8
			})()
			_ = _rhs61
			var _rhs62 bool = !(_276_v63)
			_ = _rhs62
			var _lhs42 *GlobalState = _203_globalState
			_ = _lhs42
			var _lhs43 _dafny.Array = _206_v13
			_ = _lhs43
			var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_206_v13), 0))
			_ = _lhs44
			var _lhs45 *GlobalState = _203_globalState
			_ = _lhs45
			_lhs42.F2 = _rhs58
			_275_v62 = _rhs59
			(_lhs43).ArraySet1(_rhs60, (_lhs44).Int())
			_lhs45.F10 = _rhs61
			_346_cf9 = _rhs62
			var _362_v122 *C2
			_ = _362_v122
			var _nw68 *C2 = New_C2_()
			_ = _nw68
			_nw68.Ctor__(_195_v3)
			_362_v122 = _nw68
		}
	} else if _source3.Is_DC3() {
		var _363___mcc_h10 _dafny.CodePoint = _source3.Get_().(D1_DC3).Cf5
		_ = _363___mcc_h10
		var _364_cf5 _dafny.CodePoint = _363___mcc_h10
		_ = _364_cf5
		var _365_v123 *C0
		_ = _365_v123
		var _nw69 *C0 = New_C0_()
		_ = _nw69
		_nw69.Ctor__()
		_365_v123 = _nw69
		var _366_v124 D7
		_ = _366_v124
		_366_v124 = Companion_D7_.Create_DC16_(_365_v123)
		_366_v124 = _366_v124
		var _367_v125 *C0
		_ = _367_v125
		var _nw70 *C0 = New_C0_()
		_ = _nw70
		_nw70.Ctor__()
		_367_v125 = _nw70
		Companion_Default___.M0(_206_v13, (_291_v77).F12(), _209_v14, (_291_v77).F12(), _203_globalState)
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_206_v13), 0))
		_ = _index52
		(_206_v13).ArraySet1(_276_v63, (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_206_v13), 0))
		_ = _index53
		(_206_v13).ArraySet1(_276_v63, (_index53).Int())
	} else {
		var _368___mcc_h11 D1 = _source3.Get_().(D1_DC6).Cf10
		_ = _368___mcc_h11
		var _369_cf10 D1 = _368___mcc_h11
		_ = _369_cf10
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_206_v13), 0))
		_ = _index54
		(_206_v13).ArraySet1(true, (_index54).Int())
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_206_v13), 0))
		_ = _index55
		(_206_v13).ArraySet1(!(_276_v63), (_index55).Int())
		(_291_v77).M1((_291_v77).F11(), _203_globalState)
		var _370_v126 _dafny.Array
		_ = _370_v126
		var _nw71 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
		_ = _nw71
		_370_v126 = _nw71
		var _371_v127 _dafny.Map
		_ = _371_v127
		_371_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_291_v77).F12(), _192_v0)
		var _372_v128 D3
		_ = _372_v128
		_372_v128 = Companion_D3_.Create_DC10_((_206_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_206_v13), 0))).Int()).(bool), (_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int), true, _370_v126, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_271_v58).F12(), (_291_v77).F11()))
		var _pat_let_tv5 = _370_v126
		_ = _pat_let_tv5
		var _pat_let_tv6 = _271_v58
		_ = _pat_let_tv6
		var _pat_let_tv7 = _192_v0
		_ = _pat_let_tv7
		var _source5 D3 = func(_pat_let6_0 D3) D3 {
			return func(_373_dt__update__tmp_h4 D3) D3 {
				return func(_pat_let7_0 _dafny.Array) D3 {
					return func(_374_dt__update_hcf17_h0 _dafny.Array) D3 {
						return func(_pat_let8_0 bool) D3 {
							return func(_375_dt__update_hcf14_h0 bool) D3 {
								return func(_pat_let9_0 bool) D3 {
									return func(_376_dt__update_hcf16_h0 bool) D3 {
										return Companion_D3_.Create_DC10_(_375_dt__update_hcf14_h0, (_373_dt__update__tmp_h4).Dtor_cf15(), _376_dt__update_hcf16_h0, _374_dt__update_hcf17_h0, (_373_dt__update__tmp_h4).Dtor_cf18())
									}(_pat_let9_0)
								}(_pat_let_tv7)
							}(_pat_let8_0)
						}((_pat_let_tv6).F11())
					}(_pat_let7_0)
				}(_pat_let_tv5)
			}(_pat_let6_0)
		}((func() D3 {
			if (_291_v77).F11() {
				return Companion_D3_.Create_DC10_(_278_v65, (_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int), (_271_v58).F11(), _370_v126, _371_v127)
			}
			return _372_v128
		})())
		_ = _source5
		if _source5.Is_DC10() {
			var _377___mcc_h18 bool = _source5.Get_().(D3_DC10).Cf14
			_ = _377___mcc_h18
			var _378___mcc_h19 _dafny.Int = _source5.Get_().(D3_DC10).Cf15
			_ = _378___mcc_h19
			var _379___mcc_h20 bool = _source5.Get_().(D3_DC10).Cf16
			_ = _379___mcc_h20
			var _380___mcc_h21 _dafny.Array = _source5.Get_().(D3_DC10).Cf17
			_ = _380___mcc_h21
			var _381___mcc_h22 _dafny.Map = _source5.Get_().(D3_DC10).Cf18
			_ = _381___mcc_h22
			var _382_cf18 _dafny.Map = _381___mcc_h22
			_ = _382_cf18
			var _383_cf17 _dafny.Array = _380___mcc_h21
			_ = _383_cf17
			var _384_cf16 bool = _379___mcc_h20
			_ = _384_cf16
			var _385_cf15 _dafny.Int = _378___mcc_h19
			_ = _385_cf15
			var _386_cf14 bool = _377___mcc_h18
			_ = _386_cf14
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _index56
			(_200_v8).ArraySet1((_271_v58).Fm1(_196_v4, (_280_v67).Cardinality(), _203_globalState), (_index56).Int())
			_270_v57 = _dafny.CodePoint('w')
			_195_v3 = ((_dafny.Zero).Minus((_271_v58).F12())).Times((_271_v58).F12())
			_277_v64 = (_277_v64).Update(_275_v62, Companion_Default___.Fm0(_192_v0, (_271_v58).F11(), _276_v63, (_271_v58).F11(), _203_globalState))
		} else if _source5.Is_DC9() {
			var _387___mcc_h23 _dafny.Array = _source5.Get_().(D3_DC9).Cf13
			_ = _387___mcc_h23
			var _388_cf13 _dafny.Array = _387___mcc_h23
			_ = _388_cf13
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_293_v79), 0))
			_ = _index57
			(_293_v79).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_271_v58).F11() {
					return _270_v57
				}
				return _270_v57
			})(), (_index57).Int())
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_293_v79), 0))
			_ = _index58
			(_293_v79).ArraySet1CodePoint((func() _dafny.CodePoint {
				if !_dafny.Companion_Sequence_.Contains(_275_v62, _276_v63) {
					return _270_v57
				}
				return _270_v57
			})(), (_index58).Int())
			var _389_v129 _dafny.MultiSet
			_ = _389_v129
			_389_v129 = _dafny.MultiSetOf((_271_v58).F12())
			_389_v129 = _389_v129
			var _390_v130 D8
			_ = _390_v130
			_390_v130 = Companion_D8_.Create_DC20_((_291_v77).F11(), (_206_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_206_v13), 0))).Int()).(bool), (func() bool {
				if (_371_v127).Contains((_204_v11).Cardinality()) {
					return (_371_v127).Get((_204_v11).Cardinality()).(bool)
				}
				return (_291_v77).F11()
			})(), (_271_v58).F11())
			var _391_v131 _dafny.Set
			_ = _391_v131
			_391_v131 = _dafny.SetOf(_196_v4)
			var _392_v132 *C1
			_ = _392_v132
			var _nw72 *C1 = New_C1_()
			_ = _nw72
			_nw72.Ctor__((_291_v77).F11())
			_392_v132 = _nw72
			var _393_v133 _dafny.Map
			_ = _393_v133
			_393_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_392_v132, _278_v65)
			_278_v65 = Companion_Default___.Fm0((_390_v130).Dtor_cf28(), _276_v63, (_391_v131).IsDisjointFrom(_391_v131), (func() bool {
				if (_393_v133).Contains(_392_v132) {
					return (_393_v133).Get(_392_v132).(bool)
				}
				return false
			})(), _203_globalState)
			var _394_v134 _dafny.Array
			_ = _394_v134
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_10
			var _nw73 _dafny.Array
			_ = _nw73
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw73 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Sequence = (func(_395_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_396_i9 _dafny.Int) _dafny.Sequence {
						return _395_v6
					}
				})(_198_v6)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw73 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw73).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw73).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_394_v134 = _nw73
			var _397_v135 _dafny.Map
			_ = _397_v135
			_397_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_394_v134, true)
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_206_v13), 0))
			_ = _index59
			(_206_v13).ArraySet1((func() bool {
				if (_397_v135).Contains(_394_v134) {
					return (_397_v135).Get(_394_v134).(bool)
				}
				return (_291_v77).F11()
			})(), (_index59).Int())
		} else {
			var _398___mcc_h24 D3 = _source5.Get_().(D3_DC11).Cf19
			_ = _398___mcc_h24
			var _399_cf19 D3 = _398___mcc_h24
			_ = _399_cf19
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_206_v13), 0))
			_ = _index60
			var _rhs63 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.Fm2(_278_v65, _203_globalState), (_291_v77).F12())
			_ = _rhs63
			var _rhs64 bool = !(_192_v0) || (true)
			_ = _rhs64
			var _lhs46 *GlobalState = _203_globalState
			_ = _lhs46
			var _lhs47 _dafny.Array = _206_v13
			_ = _lhs47
			var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_206_v13), 0))
			_ = _lhs48
			_lhs46.F2 = _rhs63
			(_lhs47).ArraySet1(_rhs64, (_lhs48).Int())
			var _400_v136 _dafny.MultiSet
			_ = _400_v136
			_400_v136 = _dafny.MultiSetOf(_195_v3, (_291_v77).F12(), (_291_v77).F12())
			(_203_globalState).F2 = (func() _dafny.Int {
				if ((_271_v58).F12()).Cmp((_400_v136).Cardinality()) < 0 {
					return (func() _dafny.Int {
						if (_202_v10).Contains((_291_v77).F11()) {
							return (_202_v10).Multiplicity((_291_v77).F11())
						}
						return (_291_v77).F12()
					})()
				}
				return (_dafny.Zero).Minus((_271_v58).F12())
			})()
			_278_v65 = (_372_v128).Dtor_cf16()
			_270_v57 = (func() _dafny.CodePoint {
				if (_dafny.SetOf(!((_291_v77).F11()), _278_v65, (_291_v77).F11(), (_271_v58).F11())).IsSubsetOf(_dafny.SetOf((_206_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_206_v13), 0))).Int()).(bool), !(_276_v63), _276_v63)) {
					return _dafny.CodePoint('c')
				}
				return _270_v57
			})()
		}
		_280_v67 = (_280_v67).Update(_195_v3, (_dafny.Zero).Minus((_205_v12).Cardinality()))
	}
	var _source6 D0 = Companion_D0_.Create_DC0_(_192_v0)
	_ = _source6
	if _source6.Is_DC1() {
		var _401___mcc_h25 _dafny.Int = _source6.Get_().(D0_DC1).Cf1
		_ = _401___mcc_h25
		var _402___mcc_h26 _dafny.Int = _source6.Get_().(D0_DC1).Cf2
		_ = _402___mcc_h26
		var _403___mcc_h27 _dafny.Int = _source6.Get_().(D0_DC1).Cf3
		_ = _403___mcc_h27
		var _404_cf3 _dafny.Int = _403___mcc_h27
		_ = _404_cf3
		var _405_cf2 _dafny.Int = _402___mcc_h26
		_ = _405_cf2
		var _406_cf1 _dafny.Int = _401___mcc_h25
		_ = _406_cf1
		var _source7 D1 = Companion_Default___.Fm17((_291_v77).F11(), _dafny.Companion_Sequence_.Concatenate(_198_v6, _198_v6), (_275_v62).Select((Companion_Default___.SafeIndex(_405_cf2, _dafny.IntOfUint32((_275_v62).Cardinality()))).Uint32()).(bool), _192_v0, _203_globalState)
		_ = _source7
		if _source7.Is_DC4() {
			var _407___mcc_h30 _dafny.MultiSet = _source7.Get_().(D1_DC4).Cf6
			_ = _407___mcc_h30
			var _408___mcc_h31 _dafny.Int = _source7.Get_().(D1_DC4).Cf7
			_ = _408___mcc_h31
			var _409_cf7 _dafny.Int = _408___mcc_h31
			_ = _409_cf7
			var _410_cf6 _dafny.MultiSet = _407___mcc_h30
			_ = _410_cf6
			var _411_v137 D11
			_ = _411_v137
			_411_v137 = Companion_D11_.Create_DC23_(_198_v6)
			_198_v6 = _dafny.Companion_Sequence_.Concatenate((_411_v137).Dtor_cf34(), _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_271_v58).F11() {
					return _198_v6
				}
				return _198_v6
			})(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_196_v4).Cardinality()), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_271_v58).F11() {
					return _198_v6
				}
				return _198_v6
			})()).Cardinality()))).Uint32(), (_197_v5).Cardinality()))
			var _412_v138 bool
			_ = _412_v138
			var _413_v139 _dafny.Map
			_ = _413_v139
			var _414_v140 bool
			_ = _414_v140
			var _out12 bool
			_ = _out12
			var _out13 _dafny.Map
			_ = _out13
			var _out14 bool
			_ = _out14
			_out12, _out13, _out14 = (_291_v77).M2((_271_v58).F11(), ((_dafny.Zero).Minus(_406_cf1)).Cmp((_291_v77).F12()) <= 0, _203_globalState)
			_412_v138 = _out12
			_413_v139 = _out13
			_414_v140 = _out14
			var _415_v141 bool
			_ = _415_v141
			var _416_v142 _dafny.Map
			_ = _416_v142
			var _417_v143 bool
			_ = _417_v143
			var _out15 bool
			_ = _out15
			var _out16 _dafny.Map
			_ = _out16
			var _out17 bool
			_ = _out17
			_out15, _out16, _out17 = (_271_v58).M2(((_291_v77).F11()) || (_276_v63), ((_271_v58).F12()).Cmp((_291_v77).F12()) != 0, _203_globalState)
			_415_v141 = _out15
			_416_v142 = _out16
			_417_v143 = _out17
			var _418_v144 _dafny.Sequence
			_ = _418_v144
			_418_v144 = _dafny.SeqOf((func() _dafny.Sequence {
				if _278_v65 {
					return _dafny.SeqOf((_271_v58).F11(), (_291_v77).F11())
				}
				return _275_v62
			})(), _275_v62, Companion_Default___.Fm16((_291_v77).F12(), (_291_v77).F11(), _203_globalState))
			_418_v144 = _418_v144
		} else if _source7.Is_DC5() {
			var _419___mcc_h32 bool = _source7.Get_().(D1_DC5).Cf8
			_ = _419___mcc_h32
			var _420___mcc_h33 bool = _source7.Get_().(D1_DC5).Cf9
			_ = _420___mcc_h33
			var _421_cf9 bool = _420___mcc_h33
			_ = _421_cf9
			var _422_cf8 bool = _419___mcc_h32
			_ = _422_cf8
			var _423_v145 _dafny.Map
			_ = _423_v145
			_423_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_404_cf3, ((_271_v58).F12()).Cmp((_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int)) <= 0)
			_423_v145 = _423_v145
			(_203_globalState).F2 = (_291_v77).F12()
			var _424_v146 D0
			_ = _424_v146
			_424_v146 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(287), (_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int), (_271_v58).F12())
			_195_v3 = (_198_v6).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_424_v146).Dtor_cf2(), (_dafny.Zero).Minus((func() _dafny.Int {
				if (_202_v10).Contains((_271_v58).F11()) {
					return (_202_v10).Multiplicity((_271_v58).F11())
				}
				return (_271_v58).F12()
			})())), _dafny.IntOfUint32((_198_v6).Cardinality()))).Uint32()).(_dafny.Int)
			_196_v4 = _196_v4
		} else if _source7.Is_DC3() {
			var _425___mcc_h34 _dafny.CodePoint = _source7.Get_().(D1_DC3).Cf5
			_ = _425___mcc_h34
			var _426_cf5 _dafny.CodePoint = _425___mcc_h34
			_ = _426_cf5
			_275_v62 = _275_v62
			var _427_v147 _dafny.Map
			_ = _427_v147
			_427_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_196_v4).Cardinality()), _206_v13)
			_206_v13 = (func() _dafny.Array {
				if (_427_v147).Contains((_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int)) {
					return (_427_v147).Get((_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int)).(_dafny.Array)
				}
				return _206_v13
			})()
			_196_v4 = _dafny.Companion_Sequence_.Concatenate(_196_v4, _196_v4)
			_276_v63 = (Companion_Default___.SafeDivisionInt((_271_v58).F12(), (_dafny.Zero).Minus((_271_v58).F12()))).Cmp((_291_v77).F12()) <= 0
		} else {
			var _428___mcc_h35 D1 = _source7.Get_().(D1_DC6).Cf10
			_ = _428___mcc_h35
			var _429_cf10 D1 = _428___mcc_h35
			_ = _429_cf10
			var _430_v148 _dafny.Array
			_ = _430_v148
			var _nwElement0_11 _dafny.Array = _206_v13
			_ = _nwElement0_11
			var _nw74 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(15))
			_ = _nw74
			(_nw74).ArraySet1(_nwElement0_11, 0)
			(_nw74).ArraySet1(_206_v13, 1)
			(_nw74).ArraySet1(_206_v13, 2)
			(_nw74).ArraySet1(_206_v13, 3)
			(_nw74).ArraySet1(_206_v13, 4)
			(_nw74).ArraySet1(_206_v13, 5)
			(_nw74).ArraySet1(_206_v13, 6)
			(_nw74).ArraySet1(_206_v13, 7)
			(_nw74).ArraySet1(_206_v13, 8)
			(_nw74).ArraySet1(_206_v13, 9)
			(_nw74).ArraySet1(_206_v13, 10)
			(_nw74).ArraySet1(_206_v13, 11)
			(_nw74).ArraySet1(_206_v13, 12)
			(_nw74).ArraySet1(_206_v13, 13)
			(_nw74).ArraySet1(_206_v13, 14)
			_430_v148 = _nw74
			var _431_v149 _dafny.Map
			_ = _431_v149
			_431_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(673), _430_v148)
			var _432_v150 _dafny.Map
			_ = _432_v150
			_432_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_271_v58).F11(), (func() _dafny.Array {
				if (_431_v149).Contains((_271_v58).F12()) {
					return (_431_v149).Get((_271_v58).F12()).(_dafny.Array)
				}
				return _430_v148
			})())
			_432_v150 = (_432_v150).Update(true, _430_v148)
			_270_v57 = _270_v57
			(_203_globalState).F2 = (_dafny.Zero).Minus(_404_cf3)
			var _433_v151 _dafny.Array
			_ = _433_v151
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_11
			var _nw75 _dafny.Array
			_ = _nw75
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw75 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) D11 = (func(_434_v58 *C3) func(_dafny.Int) D11 {
					return func(_435_i10 _dafny.Int) D11 {
						return (func() D11 {
							if (_434_v58).F11() {
								return Companion_D11_.Create_DC25_()
							}
							return Companion_D11_.Create_DC25_()
						})()
					}
				})(_271_v58)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw75 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw75).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw75).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_433_v151 = _nw75
			var _436_v152 D11
			_ = _436_v152
			_436_v152 = Companion_D11_.Create_DC25_()
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_433_v151), 0))
			_ = _index61
			(_433_v151).ArraySet1(_436_v152, (_index61).Int())
			var _437_v153 _dafny.Map
			_ = _437_v153
			_437_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_404_cf3, (_271_v58).Fm1(_196_v4, (_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int), _203_globalState)), _dafny.SeqOf((_291_v77).F12()))
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _index62
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _index63
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_433_v151), 0))
			_ = _index64
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _index65
			var _rhs65 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_437_v153).Contains(_406_cf1) {
					return (_437_v153).Get(_406_cf1).(_dafny.Sequence)
				}
				return _198_v6
			})()).Cardinality())
			_ = _rhs65
			var _rhs66 bool = false
			_ = _rhs66
			var _rhs67 _dafny.Int = (func() _dafny.Int {
				if _dafny.Companion_Sequence_.Contains(_196_v4, _270_v57) {
					return _406_cf1
				}
				return (_291_v77).F12()
			})()
			_ = _rhs67
			var _rhs68 D11 = _436_v152
			_ = _rhs68
			var _rhs69 _dafny.Int = Companion_Default___.SafeModuloInt(_195_v3, _dafny.IntOfUint32((_196_v4).Cardinality()))
			_ = _rhs69
			var _lhs49 _dafny.Array = _200_v8
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _lhs50
			var _lhs51 _dafny.Array = _200_v8
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _lhs52
			var _lhs53 _dafny.Array = _433_v151
			_ = _lhs53
			var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_433_v151), 0))
			_ = _lhs54
			var _lhs55 _dafny.Array = _200_v8
			_ = _lhs55
			var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _lhs56
			(_lhs49).ArraySet1(_rhs65, (_lhs50).Int())
			_278_v65 = _rhs66
			(_lhs51).ArraySet1(_rhs67, (_lhs52).Int())
			(_lhs53).ArraySet1(_rhs68, (_lhs54).Int())
			(_lhs55).ArraySet1(_rhs69, (_lhs56).Int())
		}
		var _438_v154 *C3
		_ = _438_v154
		var _nw76 *C3 = New_C3_()
		_ = _nw76
		_nw76.Ctor__(_276_v63, Companion_Default___.SafeModuloInt((_280_v67).Cardinality(), _404_cf3))
		_438_v154 = _nw76
		var _source8 D1 = Companion_Default___.Fm17((_271_v58).F11(), _198_v6, !((_dafny.IntOfUint32((_dafny.SeqOf((_271_v58).F11())).Cardinality())).Cmp((_291_v77).F12()) <= 0), _dafny.Companion_Sequence_.IsPrefixOf(_196_v4, _196_v4), _203_globalState)
		_ = _source8
		if _source8.Is_DC4() {
			var _439___mcc_h36 _dafny.MultiSet = _source8.Get_().(D1_DC4).Cf6
			_ = _439___mcc_h36
			var _440___mcc_h37 _dafny.Int = _source8.Get_().(D1_DC4).Cf7
			_ = _440___mcc_h37
			var _441_cf7 _dafny.Int = _440___mcc_h37
			_ = _441_cf7
			var _442_cf6 _dafny.MultiSet = _439___mcc_h36
			_ = _442_cf6
			_278_v65 = _dafny.Companion_Sequence_.IsProperPrefixOf(_196_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(379))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_443_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_444_i11 _dafny.Int) _dafny.CodePoint {
					return _443_v57
				}
			})(_270_v57))))
			_278_v65 = Companion_Default___.Fm0(!(_205_v12).Equals(_205_v12), (_291_v77).F11(), _192_v0, !((_438_v154).F11()), _203_globalState)
			_276_v63 = (func() bool {
				if (_271_v58).F11() {
					return (_291_v77).F11()
				}
				return (_438_v154).F11()
			})()
			var _445_v155 *C1
			_ = _445_v155
			var _nw77 *C1 = New_C1_()
			_ = _nw77
			_nw77.Ctor__((!((_291_v77).F11())) == ((_271_v58).F11()))
			_445_v155 = _nw77
		} else if _source8.Is_DC5() {
			var _446___mcc_h38 bool = _source8.Get_().(D1_DC5).Cf8
			_ = _446___mcc_h38
			var _447___mcc_h39 bool = _source8.Get_().(D1_DC5).Cf9
			_ = _447___mcc_h39
			var _448_cf9 bool = _447___mcc_h39
			_ = _448_cf9
			var _449_cf8 bool = _446___mcc_h38
			_ = _449_cf8
			_196_v4 = _196_v4
			(_291_v77).M1(((_271_v58).F12()).Cmp((_438_v154).F12()) >= 0, _203_globalState)
			var _450_v156 *C1
			_ = _450_v156
			var _nw78 *C1 = New_C1_()
			_ = _nw78
			_nw78.Ctor__(_dafny.Companion_Sequence_.IsProperPrefixOf(_275_v62, _275_v62))
			_450_v156 = _nw78
			var _451_v157 *C3
			_ = _451_v157
			var _nw79 *C3 = New_C3_()
			_ = _nw79
			_nw79.Ctor__((_291_v77).F11(), Companion_Default___.SafeModuloInt(_404_cf3, (Companion_Default___.Fm18(_405_cf2, _203_globalState)).Cardinality()))
			_451_v157 = _nw79
		} else if _source8.Is_DC3() {
			var _452___mcc_h40 _dafny.CodePoint = _source8.Get_().(D1_DC3).Cf5
			_ = _452___mcc_h40
			var _453_cf5 _dafny.CodePoint = _452___mcc_h40
			_ = _453_cf5
			var _454_v158 D1
			_ = _454_v158
			_454_v158 = Companion_D1_.Create_DC3_(_453_cf5)
			_453_cf5 = (_454_v158).Dtor_cf5()
			var _455_v159 _dafny.Sequence
			_ = _455_v159
			_455_v159 = _dafny.SeqOf((_202_v10).Update((_271_v58).F11(), Companion_Default___.Abs((_291_v77).F12())))
			_270_v57 = Companion_Default___.Fm6((_438_v154).F12(), !_dafny.Companion_Sequence_.Equal(_455_v159, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(712))).Uint32(), func(coer17 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_456_v77 *C3) func(_dafny.Int) _dafny.MultiSet {
				return func(_457_i12 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf((_456_v77).F11())
				}
			})(_291_v77)))), (_291_v77).F11(), _453_cf5, _203_globalState)
			(_203_globalState).F2 = _dafny.IntOfUint32((_198_v6).Cardinality())
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_206_v13), 0))
			_ = _index66
			(_206_v13).ArraySet1(true, (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_206_v13), 0))
			_ = _index67
			(_206_v13).ArraySet1(_276_v63, (_index67).Int())
		} else {
			var _458___mcc_h41 D1 = _source8.Get_().(D1_DC6).Cf10
			_ = _458___mcc_h41
			var _459_cf10 D1 = _458___mcc_h41
			_ = _459_cf10
			var _460_v160 _dafny.Map
			_ = _460_v160
			_460_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_200_v8, (_271_v58).F11())
			_460_v160 = (_460_v160).Update(_200_v8, !_dafny.Companion_Sequence_.Contains(_196_v4, Companion_Default___.Fm6(_195_v3, (_291_v77).F11(), _192_v0, _270_v57, _203_globalState)))
			_404_cf3 = (Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_405_cf2), (_438_v154).F12())).Plus((_438_v154).Fm1(_dafny.UnicodeSeqOfUtf8Bytes("sn"), _406_cf1, _203_globalState))
			var _461_v161 *C0
			_ = _461_v161
			var _nw80 *C0 = New_C0_()
			_ = _nw80
			_nw80.Ctor__()
			_461_v161 = _nw80
			var _462_v162 D7
			_ = _462_v162
			_462_v162 = Companion_D7_.Create_DC16_(_461_v161)
			var _463_v163 _dafny.Map
			_ = _463_v163
			_463_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_462_v162).Dtor_cf24(), _dafny.IntOfUint32((_198_v6).Cardinality()))
			var _464_v164 _dafny.Map
			_ = _464_v164
			_464_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_463_v163, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(376))).Uint32(), func(coer18 func(_dafny.Int) D11) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}(func(_465_i13 _dafny.Int) D11 {
				return Companion_D11_.Create_DC25_()
			}))).Cardinality())))
			_464_v164 = _464_v164
			(_203_globalState).F2 = _195_v3
		}
		var _466_v165 D3
		_ = _466_v165
		_466_v165 = Companion_D3_.Create_DC9_((func() _dafny.Array {
			if true {
				return _206_v13
			}
			return _206_v13
		})())
		var _source9 D3 = _466_v165
		_ = _source9
		if _source9.Is_DC10() {
			var _467___mcc_h42 bool = _source9.Get_().(D3_DC10).Cf14
			_ = _467___mcc_h42
			var _468___mcc_h43 _dafny.Int = _source9.Get_().(D3_DC10).Cf15
			_ = _468___mcc_h43
			var _469___mcc_h44 bool = _source9.Get_().(D3_DC10).Cf16
			_ = _469___mcc_h44
			var _470___mcc_h45 _dafny.Array = _source9.Get_().(D3_DC10).Cf17
			_ = _470___mcc_h45
			var _471___mcc_h46 _dafny.Map = _source9.Get_().(D3_DC10).Cf18
			_ = _471___mcc_h46
			var _472_cf18 _dafny.Map = _471___mcc_h46
			_ = _472_cf18
			var _473_cf17 _dafny.Array = _470___mcc_h45
			_ = _473_cf17
			var _474_cf16 bool = _469___mcc_h44
			_ = _474_cf16
			var _475_cf15 _dafny.Int = _468___mcc_h43
			_ = _475_cf15
			var _476_cf14 bool = _467___mcc_h42
			_ = _476_cf14
			var _rhs70 bool = (_271_v58).F11()
			_ = _rhs70
			var _rhs71 _dafny.CodePoint = (_196_v4).Select((Companion_Default___.SafeIndex(_475_cf15, _dafny.IntOfUint32((_196_v4).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs71
			var _rhs72 bool = !(!(false))
			_ = _rhs72
			_474_cf16 = _rhs70
			_270_v57 = _rhs71
			_276_v63 = _rhs72
			var _477_v166 _dafny.Sequence
			_ = _477_v166
			_477_v166 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_275_v62, _275_v62))
			_477_v166 = _477_v166
			_475_cf15 = ((_280_v67).Merge((_280_v67).Merge(_280_v67))).Cardinality()
			_204_v11 = _204_v11
		} else if _source9.Is_DC9() {
			var _478___mcc_h47 _dafny.Array = _source9.Get_().(D3_DC9).Cf13
			_ = _478___mcc_h47
			var _479_cf13 _dafny.Array = _478___mcc_h47
			_ = _479_cf13
			_195_v3 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_196_v4, _dafny.UnicodeSeqOfUtf8Bytes("qgqwbmtnb")), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(58))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}(func(_480_i14 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(58))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}(func(_481_i14 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			}))).Cardinality()))).Uint32(), Companion_Default___.Fm6(_195_v3, (_291_v77).F11(), _278_v65, _dafny.CodePoint('b'), _203_globalState)), _196_v4))).Cardinality())
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _index68
			(_200_v8).ArraySet1((_271_v58).F12(), (_index68).Int())
			var _482_v167 D0
			_ = _482_v167
			_482_v167 = Companion_D0_.Create_DC1_((_271_v58).F12(), (_dafny.Zero).Minus((_271_v58).Fm1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(637))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_483_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_484_i15 _dafny.Int) _dafny.CodePoint {
					return _483_v57
				}
			})(_270_v57))), (_271_v58).F12(), _203_globalState)), (_438_v154).F12())
			_195_v3 = (_482_v167).Dtor_cf3()
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_479_cf13), 0))
			_ = _index69
			(_479_cf13).ArraySet1(!(true) || (_278_v65), (_index69).Int())
			var _485_v168 D11
			_ = _485_v168
			_485_v168 = Companion_D11_.Create_DC26_((_291_v77).F12(), _200_v8, (_438_v154).F11())
			var _486_v169 D0
			_ = _486_v169
			_486_v169 = Companion_D0_.Create_DC0_((_485_v168).Dtor_cf38())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_479_cf13), 0))
			_ = _index70
			(_479_cf13).ArraySet1((_486_v169).Dtor_cf0(), (_index70).Int())
		} else {
			var _487___mcc_h48 D3 = _source9.Get_().(D3_DC11).Cf19
			_ = _487___mcc_h48
			var _488_cf19 D3 = _487___mcc_h48
			_ = _488_cf19
			_405_cf2 = ((_193_v1).Cardinality()).Minus(_405_cf2)
			var _489_v170 _dafny.Array
			_ = _489_v170
			var _nw81 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
			_ = _nw81
			_489_v170 = _nw81
			var _490_v171 *C1
			_ = _490_v171
			var _nw82 *C1 = New_C1_()
			_ = _nw82
			_nw82.Ctor__((_291_v77).F11())
			_490_v171 = _nw82
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_489_v170), 0))
			_ = _index71
			(_489_v170).ArraySet1(_490_v171, (_index71).Int())
			var _491_v172 *C1
			_ = _491_v172
			_491_v172 = _490_v171
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_489_v170), 0))
			_ = _index72
			(_489_v170).ArraySet1((_491_v172), (_index72).Int())
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _index73
			(_200_v8).ArraySet1(((_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int)).Times((_271_v58).F12()), (_index73).Int())
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_206_v13), 0))
			_ = _index74
			(_206_v13).ArraySet1(_278_v65, (_index74).Int())
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_206_v13), 0))
			_ = _index75
			(_206_v13).ArraySet1((_438_v154).F11(), (_index75).Int())
		}
	} else if _source6.Is_DC0() {
		var _492___mcc_h28 bool = _source6.Get_().(D0_DC0).Cf0
		_ = _492___mcc_h28
		var _493_cf0 bool = _492___mcc_h28
		_ = _493_cf0
		var _494_v173 _dafny.Map
		_ = _494_v173
		_494_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_278_v65, _270_v57)
		var _495_v174 bool
		_ = _495_v174
		var _496_v175 _dafny.Map
		_ = _496_v175
		var _497_v176 bool
		_ = _497_v176
		var _out18 bool
		_ = _out18
		var _out19 _dafny.Map
		_ = _out19
		var _out20 bool
		_ = _out20
		_out18, _out19, _out20 = (_291_v77).M2(((_494_v173).Cardinality()).Cmp((_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int)) == 0, (_271_v58).F11(), _203_globalState)
		_495_v174 = _out18
		_496_v175 = _out19
		_497_v176 = _out20
		(_271_v58).M1(_278_v65, _203_globalState)
		_495_v174 = !_dafny.Companion_Sequence_.Equal(_198_v6, _dafny.Companion_Sequence_.Concatenate(_198_v6, _198_v6))
		var _498_v177 _dafny.Map
		_ = _498_v177
		_498_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_204_v11).Cardinality(), _495_v174)
		var _499_v178 D3
		_ = _499_v178
		_499_v178 = Companion_D3_.Create_DC10_(_495_v174, (_271_v58).F12(), _276_v63, _206_v13, _498_v177)
		_204_v11 = (_204_v11).Update(((_499_v178).Dtor_cf15()).Cmp((_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int)) >= 0, _276_v63)
	} else {
		var _500___mcc_h29 D0 = _source6.Get_().(D0_DC2).Cf4
		_ = _500___mcc_h29
		var _501_cf4 D0 = _500___mcc_h29
		_ = _501_cf4
		var _502_v179 _dafny.Sequence
		_ = _502_v179
		_502_v179 = _dafny.SeqOf(_196_v4, _196_v4, _196_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(570))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}((func(_503_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_504_i16 _dafny.Int) _dafny.CodePoint {
				return _503_v57
			}
		})(_270_v57))))
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_502_v179, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(24))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}((func(_505_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_506_i17 _dafny.Int) _dafny.Sequence {
				return _505_v4
			}
		})(_196_v4)))) {
			var _507_v180 _dafny.Map
			_ = _507_v180
			_507_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_271_v58).F12()), (_271_v58).F11())
			_507_v180 = (_507_v180).Update(_195_v3, true)
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_206_v13), 0))
			_ = _index76
			(_206_v13).ArraySet1((_291_v77).F11(), (_index76).Int())
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_206_v13), 0))
			_ = _index77
			(_206_v13).ArraySet1((Companion_D3_.Create_DC10_((_291_v77).F11(), (_291_v77).F12(), _278_v65, _206_v13, _507_v180)).Dtor_cf16(), (_index77).Int())
			(_291_v77).M1(((_271_v58).F12()).Cmp((_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int)) >= 0, _203_globalState)
			_196_v4 = _196_v4
			(_291_v77).M1((_291_v77).F11(), _203_globalState)
		} else {
			var _508_v181 _dafny.Array
			_ = _508_v181
			var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
			_ = _nw83
			_508_v181 = _nw83
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_508_v181), 0))
			_ = _index78
			(_508_v181).ArraySet1(_206_v13, (_index78).Int())
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_508_v181), 0))
			_ = _index79
			(_508_v181).ArraySet1(_206_v13, (_index79).Int())
			var _509_v182 *C3
			_ = _509_v182
			var _nw84 *C3 = New_C3_()
			_ = _nw84
			_nw84.Ctor__((_291_v77).F11(), (_291_v77).F12())
			_509_v182 = _nw84
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _index80
			(_200_v8).ArraySet1(Companion_Default___.Fm4(_203_globalState), (_index80).Int())
			var _510_v183 _dafny.Sequence
			_ = _510_v183
			_510_v183 = _196_v4
			(_203_globalState).F2 = _dafny.IntOfUint32((_510_v183).Cardinality())
			var _511_v184 D11
			_ = _511_v184
			_511_v184 = Companion_D11_.Create_DC24_((_509_v182).F12())
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))
			_ = _index81
			(_200_v8).ArraySet1((_511_v184).Dtor_cf35(), (_index81).Int())
		}
		var _512_v185 D0
		_ = _512_v185
		_512_v185 = Companion_D0_.Create_DC1_(_195_v3, (_291_v77).F12(), (_271_v58).F12())
		var _rhs73 _dafny.Int = ((_291_v77).F12()).Times((_198_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_275_v62).Cardinality()), _dafny.IntOfUint32((_198_v6).Cardinality()))).Uint32()).(_dafny.Int))
		_ = _rhs73
		var _rhs74 bool = (_dafny.IntOfInt64(694)).Cmp((_512_v185).Dtor_cf2()) > 0
		_ = _rhs74
		var _rhs75 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_291_v77).F12()), (_271_v58).F12())
		_ = _rhs75
		var _rhs76 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_271_v58).F12(), _dafny.IntOfInt64(697)), (_291_v77).F12())
		_ = _rhs76
		_195_v3 = _rhs73
		_278_v65 = _rhs74
		_280_v67 = _rhs75
		_195_v3 = _rhs76
		var _513_v186 _dafny.Map
		_ = _513_v186
		_513_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_196_v4, true)
		var _514_v187 D0
		_ = _514_v187
		_514_v187 = Companion_D0_.Create_DC0_((func() bool {
			if (_513_v186).Contains(_dafny.UnicodeSeqOfUtf8Bytes("ddnbjifo")) {
				return (_513_v186).Get(_dafny.UnicodeSeqOfUtf8Bytes("ddnbjifo")).(bool)
			}
			return _278_v65
		})())
		var _source10 D0 = _514_v187
		_ = _source10
		if _source10.Is_DC1() {
			var _515___mcc_h49 _dafny.Int = _source10.Get_().(D0_DC1).Cf1
			_ = _515___mcc_h49
			var _516___mcc_h50 _dafny.Int = _source10.Get_().(D0_DC1).Cf2
			_ = _516___mcc_h50
			var _517___mcc_h51 _dafny.Int = _source10.Get_().(D0_DC1).Cf3
			_ = _517___mcc_h51
			var _518_cf3 _dafny.Int = _517___mcc_h51
			_ = _518_cf3
			var _519_cf2 _dafny.Int = _516___mcc_h50
			_ = _519_cf2
			var _520_cf1 _dafny.Int = _515___mcc_h49
			_ = _520_cf1
			var _521_v188 _dafny.Map
			_ = _521_v188
			_521_v188 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm19(_203_globalState), _195_v3)
			_519_cf2 = (func() _dafny.Int {
				if (_271_v58).F11() {
					return Companion_Default___.SafeModuloInt((_291_v77).F12(), _519_cf2)
				}
				return (_521_v188).Cardinality()
			})()
			_278_v65 = (_291_v77).F11()
			_192_v0 = ((_291_v77).F12()).Cmp(_518_cf3) <= 0
			_192_v0 = (_291_v77).F11()
		} else if _source10.Is_DC0() {
			var _522___mcc_h52 bool = _source10.Get_().(D0_DC0).Cf0
			_ = _522___mcc_h52
			var _523_cf0 bool = _522___mcc_h52
			_ = _523_cf0
			var _524_v189 *C1
			_ = _524_v189
			var _nw85 *C1 = New_C1_()
			_ = _nw85
			_nw85.Ctor__(false)
			_524_v189 = _nw85
			_523_cf0 = _523_cf0
			_513_v186 = (_513_v186).Update(_196_v4, true)
			var _525_v190 bool
			_ = _525_v190
			var _526_v191 _dafny.Map
			_ = _526_v191
			var _527_v192 bool
			_ = _527_v192
			var _out21 bool
			_ = _out21
			var _out22 _dafny.Map
			_ = _out22
			var _out23 bool
			_ = _out23
			_out21, _out22, _out23 = (_291_v77).M2((func() bool {
				if _278_v65 {
					return (_291_v77).F11()
				}
				return (_271_v58).F11()
			})(), !(_276_v63), _203_globalState)
			_525_v190 = _out21
			_526_v191 = _out22
			_527_v192 = _out23
		} else {
			var _528___mcc_h53 D0 = _source10.Get_().(D0_DC2).Cf4
			_ = _528___mcc_h53
			var _529_cf4 D0 = _528___mcc_h53
			_ = _529_cf4
			var _530_v193 bool
			_ = _530_v193
			var _531_v194 _dafny.Map
			_ = _531_v194
			var _532_v195 bool
			_ = _532_v195
			var _out24 bool
			_ = _out24
			var _out25 _dafny.Map
			_ = _out25
			var _out26 bool
			_ = _out26
			_out24, _out25, _out26 = (_291_v77).M2(!((_271_v58).F11()), _278_v65, _203_globalState)
			_530_v193 = _out24
			_531_v194 = _out25
			_532_v195 = _out26
			var _533_v196 D7
			_ = _533_v196
			_533_v196 = Companion_D7_.Create_DC17_((_291_v77).F11())
			_278_v65 = Companion_Default___.Fm0((_532_v195) == ((_291_v77).F11()), (_533_v196).Dtor_cf25(), _530_v193, _276_v63, _203_globalState)
			_532_v195 = _192_v0
			var _534_v197 _dafny.Array
			_ = _534_v197
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_12
			var _nw86 _dafny.Array
			_ = _nw86
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw86 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Sequence = (func(_535_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_536_i18 _dafny.Int) _dafny.Sequence {
						return (_535_v4)
					}
				})(_196_v4)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw86 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw86).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw86).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_534_v197 = _nw86
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_534_v197), 0))
			_ = _index82
			(_534_v197).ArraySet1(_196_v4, (_index82).Int())
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_534_v197), 0))
			_ = _index83
			var _rhs77 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("c"), _196_v4)
			_ = _rhs77
			var _rhs78 bool = _532_v195
			_ = _rhs78
			var _lhs57 _dafny.Array = _534_v197
			_ = _lhs57
			var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_534_v197), 0))
			_ = _lhs58
			(_lhs57).ArraySet1(_rhs77, (_lhs58).Int())
			_530_v193 = _rhs78
		}
		var _537_v198 _dafny.Set
		_ = _537_v198
		_537_v198 = _dafny.SetOf((_291_v77).F12(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-9))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}((func(_538_v1 _dafny.Set) func(_dafny.Int) _dafny.Int {
			return func(_539_i19 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus((_538_v1).Cardinality())
			}
		})(_193_v1)))).Cardinality()), (_200_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_200_v8), 0))).Int()).(_dafny.Int))
		var _540_v199 _dafny.Map
		_ = _540_v199
		_540_v199 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_537_v198, _192_v0)
		var _541_v200 _dafny.Sequence
		_ = _541_v200
		_541_v200 = _dafny.SeqOf((_540_v199).Merge(_540_v199), _540_v199, _540_v199, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_537_v198, false))
		var _rhs79 _dafny.Int = (func() _dafny.Int {
			if (_271_v58).F11() {
				return (_dafny.Zero).Minus(((_271_v58).F12()).Minus(_195_v3))
			}
			return Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_271_v58).F12()), _dafny.IntOfUint32((_196_v4).Cardinality()))
		})()
		_ = _rhs79
		var _rhs80 _dafny.Map = (_541_v200).Select((Companion_Default___.SafeIndex(_195_v3, _dafny.IntOfUint32((_541_v200).Cardinality()))).Uint32()).(_dafny.Map)
		_ = _rhs80
		_195_v3 = _rhs79
		_540_v199 = _rhs80
	}
	_dafny.Print(_192_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_193_v1).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_194_v2).Equals(_dafny.SetOf(_dafny.SetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_195_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_196_v4.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_197_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("bk"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_198_v6, _dafny.SeqOf(_dafny.One, _dafny.One, _dafny.IntOfInt64(848))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_199_v7, _dafny.SeqOf(_dafny.SetOf(true), _dafny.SetOf(true), _dafny.SetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_200_v8).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_201_v9).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_202_v10).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_203_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_203_globalState).F4()).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_203_globalState).F5()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(848), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_203_globalState).F6(), _dafny.SeqOf(_dafny.One, _dafny.One, _dafny.IntOfInt64(848))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_203_globalState.F7, _dafny.SeqOf(_dafny.IntOfInt64(490))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_globalState.F10).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_204_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(848))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_206_v13).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_209_v14).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v56).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("sharml"), _dafny.IntOfInt64(848))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_270_v57)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_271_v58).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_271_v58).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v59).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_273_v60).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_v61).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_275_v62, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_276_v63)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v64).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false, true, true), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_278_v65)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_279_v66).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.IntOfInt64(848), _dafny.IntOfInt64(848), _dafny.IntOfInt64(848)), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_280_v67).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(848), _dafny.IntOfInt64(848))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v68).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(424), _dafny.IntOfInt64(154))).UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(2), _dafny.IntOfInt64(114)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_282_v69).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_283_v70).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_284_i4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_291_v77).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_291_v77).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_292_v78).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_294_v80).Cardinality())
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
	Cf2 _dafny.Int
	Cf3 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Int, Cf3 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
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
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf0() bool {
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
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
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
	Cf6 _dafny.MultiSet
	Cf7 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf6 _dafny.MultiSet, Cf7 _dafny.Int) D1 {
	return D1{D1_DC4{Cf6, Cf7}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf8 bool
	Cf9 bool
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf8 bool, Cf9 bool) D1 {
	return D1{D1_DC5{Cf8, Cf9}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC3 struct {
	Cf5 _dafny.CodePoint
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf5 _dafny.CodePoint) D1 {
	return D1{D1_DC3{Cf5}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC6 struct {
	Cf10 D1
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf10 D1) D1 {
	return D1{D1_DC6{Cf10}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.EmptyMultiSet, _dafny.Zero)
}

func (_this D1) Dtor_cf6() _dafny.MultiSet {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) Dtor_cf9() bool {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf5() _dafny.CodePoint {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf10() D1 {
	return _this.Get_().(D1_DC6).Cf10
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf6.Equals(data2.Cf6) && data1.Cf7.Cmp(data2.Cf7) == 0
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9 == data2.Cf9
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf5 == data2.Cf5
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
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

type D2_DC8 struct {
	Cf12 _dafny.Int
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf12 _dafny.Int) D2 {
	return D2{D2_DC8{Cf12}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC7 struct {
	Cf11 _dafny.Array
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf11 _dafny.Array) D2 {
	return D2{D2_DC7{Cf11}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(_dafny.Zero)
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf12
}

func (_this D2) Dtor_cf11() _dafny.Array {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf11 == data2.Cf11
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
	Cf14 bool
	Cf15 _dafny.Int
	Cf16 bool
	Cf17 _dafny.Array
	Cf18 _dafny.Map
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf14 bool, Cf15 _dafny.Int, Cf16 bool, Cf17 _dafny.Array, Cf18 _dafny.Map) D3 {
	return D3{D3_DC10{Cf14, Cf15, Cf16, Cf17, Cf18}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC9 struct {
	Cf13 _dafny.Array
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf13 _dafny.Array) D3 {
	return D3{D3_DC9{Cf13}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC11 struct {
	Cf19 D3
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf19 D3) D3 {
	return D3{D3_DC11{Cf19}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(false, _dafny.Zero, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptyMap)
}

func (_this D3) Dtor_cf14() bool {
	return _this.Get_().(D3_DC10).Cf14
}

func (_this D3) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf15
}

func (_this D3) Dtor_cf16() bool {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) Dtor_cf17() _dafny.Array {
	return _this.Get_().(D3_DC10).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Map {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) Dtor_cf13() _dafny.Array {
	return _this.Get_().(D3_DC9).Cf13
}

func (_this D3) Dtor_cf19() D3 {
	return _this.Get_().(D3_DC11).Cf19
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf19) + ")"
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
			return ok && data1.Cf14 == data2.Cf14 && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16 == data2.Cf16 && data1.Cf17 == data2.Cf17 && data1.Cf18.Equals(data2.Cf18)
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf13 == data2.Cf13
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf19.Equals(data2.Cf19)
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
	Cf21 _dafny.CodePoint
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf21 _dafny.CodePoint) D4 {
	return D4{D4_DC13{Cf21}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC12 struct {
	Cf20 _dafny.Set
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf20 _dafny.Set) D4 {
	return D4{D4_DC12{Cf20}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_(_dafny.CodePoint('D'))
}

func (_this D4) Dtor_cf21() _dafny.CodePoint {
	return _this.Get_().(D4_DC13).Cf21
}

func (_this D4) Dtor_cf20() _dafny.Set {
	return _this.Get_().(D4_DC12).Cf20
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf20) + ")"
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
			return ok && data1.Cf21 == data2.Cf21
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf20.Equals(data2.Cf20)
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

type D5_DC14 struct {
	Cf22 _dafny.Map
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf22 _dafny.Map) D5 {
	return D5{D5_DC14{Cf22}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D5) Dtor_cf22() _dafny.Map {
	return _this.Get_().(D5_DC14).Cf22
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	Cf23 _dafny.MultiSet
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf23 _dafny.MultiSet) D6 {
	return D6{D6_DC15{Cf23}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D6) Dtor_cf23() _dafny.MultiSet {
	return _this.Get_().(D6_DC15).Cf23
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf23) + ")"
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
			return ok && data1.Cf23.Equals(data2.Cf23)
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

type D7_DC17 struct {
	Cf25 bool
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf25 bool) D7 {
	return D7{D7_DC17{Cf25}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

type D7_DC16 struct {
	Cf24 *C0
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf24 *C0) D7 {
	return D7{D7_DC16{Cf24}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

type D7_DC18 struct {
	Cf26 D7
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf26 D7) D7 {
	return D7{D7_DC18{Cf26}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC17_(false)
}

func (_this D7) Dtor_cf25() bool {
	return _this.Get_().(D7_DC17).Cf25
}

func (_this D7) Dtor_cf24() *C0 {
	return _this.Get_().(D7_DC16).Cf24
}

func (_this D7) Dtor_cf26() D7 {
	return _this.Get_().(D7_DC18).Cf26
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D7_DC16:
		{
			return "D7.DC16" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf26) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf25 == data2.Cf25
		}
	case D7_DC16:
		{
			data2, ok := other.Get_().(D7_DC16)
			return ok && data1.Cf24 == data2.Cf24
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf26.Equals(data2.Cf26)
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
	Cf28 bool
	Cf29 bool
	Cf30 bool
	Cf31 bool
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf28 bool, Cf29 bool, Cf30 bool, Cf31 bool) D8 {
	return D8{D8_DC20{Cf28, Cf29, Cf30, Cf31}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

type D8_DC19 struct {
	Cf27 *C2
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf27 *C2) D8 {
	return D8{D8_DC19{Cf27}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC20_(false, false, false, false)
}

func (_this D8) Dtor_cf28() bool {
	return _this.Get_().(D8_DC20).Cf28
}

func (_this D8) Dtor_cf29() bool {
	return _this.Get_().(D8_DC20).Cf29
}

func (_this D8) Dtor_cf30() bool {
	return _this.Get_().(D8_DC20).Cf30
}

func (_this D8) Dtor_cf31() bool {
	return _this.Get_().(D8_DC20).Cf31
}

func (_this D8) Dtor_cf27() *C2 {
	return _this.Get_().(D8_DC19).Cf27
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf27) + ")"
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
			return ok && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29 && data1.Cf30 == data2.Cf30 && data1.Cf31 == data2.Cf31
		}
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
			return ok && data1.Cf27 == data2.Cf27
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

type D9_DC21 struct {
	Cf32 _dafny.Sequence
}

func (D9_DC21) isD9() {}

func (CompanionStruct_D9_) Create_DC21_(Cf32 _dafny.Sequence) D9 {
	return D9{D9_DC21{Cf32}}
}

func (_this D9) Is_DC21() bool {
	_, ok := _this.Get_().(D9_DC21)
	return ok
}

func (CompanionStruct_D9_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D9) Dtor_cf32() _dafny.Sequence {
	return _this.Get_().(D9_DC21).Cf32
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC21:
		{
			return "D9.DC21" + "(" + data.Cf32.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC21:
		{
			data2, ok := other.Get_().(D9_DC21)
			return ok && data1.Cf32.Equals(data2.Cf32)
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

type D10_DC22 struct {
	Cf33 _dafny.Sequence
}

func (D10_DC22) isD10() {}

func (CompanionStruct_D10_) Create_DC22_(Cf33 _dafny.Sequence) D10 {
	return D10{D10_DC22{Cf33}}
}

func (_this D10) Is_DC22() bool {
	_, ok := _this.Get_().(D10_DC22)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D10) Dtor_cf33() _dafny.Sequence {
	return _this.Get_().(D10_DC22).Cf33
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC22:
		{
			return "D10.DC22" + "(" + _dafny.String(data.Cf33) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC22:
		{
			data2, ok := other.Get_().(D10_DC22)
			return ok && data1.Cf33.Equals(data2.Cf33)
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

type D11_DC24 struct {
	Cf35 _dafny.Int
}

func (D11_DC24) isD11() {}

func (CompanionStruct_D11_) Create_DC24_(Cf35 _dafny.Int) D11 {
	return D11{D11_DC24{Cf35}}
}

func (_this D11) Is_DC24() bool {
	_, ok := _this.Get_().(D11_DC24)
	return ok
}

type D11_DC25 struct {
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_() D11 {
	return D11{D11_DC25{}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

type D11_DC26 struct {
	Cf36 _dafny.Int
	Cf37 _dafny.Array
	Cf38 bool
}

func (D11_DC26) isD11() {}

func (CompanionStruct_D11_) Create_DC26_(Cf36 _dafny.Int, Cf37 _dafny.Array, Cf38 bool) D11 {
	return D11{D11_DC26{Cf36, Cf37, Cf38}}
}

func (_this D11) Is_DC26() bool {
	_, ok := _this.Get_().(D11_DC26)
	return ok
}

type D11_DC23 struct {
	Cf34 _dafny.Sequence
}

func (D11_DC23) isD11() {}

func (CompanionStruct_D11_) Create_DC23_(Cf34 _dafny.Sequence) D11 {
	return D11{D11_DC23{Cf34}}
}

func (_this D11) Is_DC23() bool {
	_, ok := _this.Get_().(D11_DC23)
	return ok
}

type D11_DC27 struct {
	Cf39 D11
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf39 D11) D11 {
	return D11{D11_DC27{Cf39}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC24_(_dafny.Zero)
}

func (_this D11) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D11_DC24).Cf35
}

func (_this D11) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D11_DC26).Cf36
}

func (_this D11) Dtor_cf37() _dafny.Array {
	return _this.Get_().(D11_DC26).Cf37
}

func (_this D11) Dtor_cf38() bool {
	return _this.Get_().(D11_DC26).Cf38
}

func (_this D11) Dtor_cf34() _dafny.Sequence {
	return _this.Get_().(D11_DC23).Cf34
}

func (_this D11) Dtor_cf39() D11 {
	return _this.Get_().(D11_DC27).Cf39
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC24:
		{
			return "D11.DC24" + "(" + _dafny.String(data.Cf35) + ")"
		}
	case D11_DC25:
		{
			return "D11.DC25"
		}
	case D11_DC26:
		{
			return "D11.DC26" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D11_DC23:
		{
			return "D11.DC23" + "(" + _dafny.String(data.Cf34) + ")"
		}
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf39) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC24:
		{
			data2, ok := other.Get_().(D11_DC24)
			return ok && data1.Cf35.Cmp(data2.Cf35) == 0
		}
	case D11_DC25:
		{
			_, ok := other.Get_().(D11_DC25)
			return ok
		}
	case D11_DC26:
		{
			data2, ok := other.Get_().(D11_DC26)
			return ok && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37 == data2.Cf37 && data1.Cf38 == data2.Cf38
		}
	case D11_DC23:
		{
			data2, ok := other.Get_().(D11_DC23)
			return ok && data1.Cf34.Equals(data2.Cf34)
		}
	case D11_DC27:
		{
			data2, ok := other.Get_().(D11_DC27)
			return ok && data1.Cf39.Equals(data2.Cf39)
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
	Cf40 *C1
}

func (D12_DC28) isD12() {}

func (CompanionStruct_D12_) Create_DC28_(Cf40 *C1) D12 {
	return D12{D12_DC28{Cf40}}
}

func (_this D12) Is_DC28() bool {
	_, ok := _this.Get_().(D12_DC28)
	return ok
}

func (CompanionStruct_D12_) Default() *C1 {
	return (*C1)(nil)
}

func (_this D12) Dtor_cf40() *C1 {
	return _this.Get_().(D12_DC28).Cf40
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC28:
		{
			return "D12.DC28" + "(" + _dafny.String(data.Cf40) + ")"
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
			return ok && data1.Cf40 == data2.Cf40
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

type D13_DC29 struct {
	Cf41 _dafny.Sequence
}

func (D13_DC29) isD13() {}

func (CompanionStruct_D13_) Create_DC29_(Cf41 _dafny.Sequence) D13 {
	return D13{D13_DC29{Cf41}}
}

func (_this D13) Is_DC29() bool {
	_, ok := _this.Get_().(D13_DC29)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D13) Dtor_cf41() _dafny.Sequence {
	return _this.Get_().(D13_DC29).Cf41
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC29:
		{
			return "D13.DC29" + "(" + _dafny.String(data.Cf41) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC29:
		{
			data2, ok := other.Get_().(D13_DC29)
			return ok && data1.Cf41.Equals(data2.Cf41)
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

type D14_DC30 struct {
	Cf42 _dafny.Sequence
}

func (D14_DC30) isD14() {}

func (CompanionStruct_D14_) Create_DC30_(Cf42 _dafny.Sequence) D14 {
	return D14{D14_DC30{Cf42}}
}

func (_this D14) Is_DC30() bool {
	_, ok := _this.Get_().(D14_DC30)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D14) Dtor_cf42() _dafny.Sequence {
	return _this.Get_().(D14_DC30).Cf42
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC30:
		{
			return "D14.DC30" + "(" + _dafny.String(data.Cf42) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC30:
		{
			data2, ok := other.Get_().(D14_DC30)
			return ok && data1.Cf42.Equals(data2.Cf42)
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

type D15_DC32 struct {
	Cf44 bool
}

func (D15_DC32) isD15() {}

func (CompanionStruct_D15_) Create_DC32_(Cf44 bool) D15 {
	return D15{D15_DC32{Cf44}}
}

func (_this D15) Is_DC32() bool {
	_, ok := _this.Get_().(D15_DC32)
	return ok
}

type D15_DC31 struct {
	Cf43 _dafny.Map
}

func (D15_DC31) isD15() {}

func (CompanionStruct_D15_) Create_DC31_(Cf43 _dafny.Map) D15 {
	return D15{D15_DC31{Cf43}}
}

func (_this D15) Is_DC31() bool {
	_, ok := _this.Get_().(D15_DC31)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC32_(false)
}

func (_this D15) Dtor_cf44() bool {
	return _this.Get_().(D15_DC32).Cf44
}

func (_this D15) Dtor_cf43() _dafny.Map {
	return _this.Get_().(D15_DC31).Cf43
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC32:
		{
			return "D15.DC32" + "(" + _dafny.String(data.Cf44) + ")"
		}
	case D15_DC31:
		{
			return "D15.DC31" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC32:
		{
			data2, ok := other.Get_().(D15_DC32)
			return ok && data1.Cf44 == data2.Cf44
		}
	case D15_DC31:
		{
			data2, ok := other.Get_().(D15_DC31)
			return ok && data1.Cf43.Equals(data2.Cf43)
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

// Definition of class GlobalState
type GlobalState struct {
	F2  _dafny.Int
	F7  _dafny.Sequence
	F10 _dafny.Array
	_f0 _dafny.Int
	_f1 bool
	_f3 bool
	_f4 _dafny.Set
	_f5 _dafny.Map
	_f6 _dafny.Sequence
	_f8 bool
	_f9 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = _dafny.Zero
	_this.F7 = _dafny.EmptySeq
	_this.F10 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f0 = _dafny.Zero
	_this._f1 = false
	_this._f3 = false
	_this._f4 = _dafny.EmptySet
	_this._f5 = _dafny.EmptyMap
	_this._f6 = _dafny.EmptySeq
	_this._f8 = false
	_this._f9 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 bool, f2 _dafny.Int, f3 bool, f4 _dafny.Set, f5 _dafny.Map, f6 _dafny.Sequence, f7 _dafny.Sequence, f8 bool, f9 bool, f10 _dafny.Array) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Set {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Map {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Sequence {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F8() bool {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
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
func (_this *C0) Fm7(p0 bool, p1 D1, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) bool {
	{
		return ((_dafny.IntOfInt64(107)).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wt")).Cardinality()))).Cmp(_dafny.IntOfInt64(-225)) < 0
	}
}
func (_this *C0) Fm8(p0 _dafny.Int, p1 _dafny.Set, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.IntOfInt64(341)).Minus(_dafny.IntOfInt64(-256))).Plus(((_dafny.SetOf(_dafny.IntOfInt64(789), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lan")).Cardinality()), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(150))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ghqr")).Cardinality()))).Cardinality(), false)).Cardinality())).Intersection(_dafny.SetOf(_dafny.IntOfInt64(330)))).Cardinality())
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	F14 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this.F14 = false
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

func (_this *C1) Ctor__(f14 bool) {
	{
		(_this).F14 = f14
	}
}
func (_this *C1) M4(p0 _dafny.Map, globalState *GlobalState) (_dafny.Sequence, D1, _dafny.Int) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 D1 = Companion_D1_.Default()
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _542_v0 _dafny.Int
		_ = _542_v0
		_542_v0 = _dafny.IntOfInt64(-769)
		r2 = _542_v0
		_542_v0 = (_542_v0).Plus(_dafny.IntOfInt64(358))
		var _543_v1 _dafny.Sequence
		_ = _543_v1
		_543_v1 = _dafny.SeqOf(_dafny.IntOfInt64(341), _542_v0)
		var _544_v2 _dafny.CodePoint
		_ = _544_v2
		_544_v2 = _dafny.CodePoint('q')
		var _545_v3 _dafny.MultiSet
		_ = _545_v3
		_545_v3 = _dafny.MultiSetOf(_542_v0)
		var _hi1 _dafny.Int = ((_545_v3).Cardinality()).Times(_542_v0)
		_ = _hi1
		for _546_i0 := (_543_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("n"), (Companion_Default___.SafeIndex(_542_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("n")).Cardinality()))).Uint32(), _544_v2)).Cardinality()), _dafny.IntOfUint32((_543_v1).Cardinality()))).Uint32()).(_dafny.Int); _546_i0.Cmp(_hi1) < 0; _546_i0 = _546_i0.Plus(_dafny.One) {
			var _547_v4 _dafny.Array
			_ = _547_v4
			var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw87
			_547_v4 = _nw87
			var _548_v5 _dafny.Array
			_ = _548_v5
			var _nw88 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
			_ = _nw88
			_548_v5 = _nw88
			var _549_v6 _dafny.MultiSet
			_ = _549_v6
			_549_v6 = _dafny.MultiSetOf(_548_v5, _548_v5)
			var _550_v7 _dafny.Map
			_ = _550_v7
			_550_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F14)
			var _551_v8 _dafny.Sequence
			_ = _551_v8
			_551_v8 = _dafny.SeqOf(_550_v7, _550_v7, _550_v7, _550_v7)
			var _552_v9 _dafny.Map
			_ = _552_v9
			_552_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_551_v8).Cardinality()), _542_v0)
			var _553_v10 _dafny.Array
			_ = _553_v10
			var _nwElement0_12 _dafny.Int = _dafny.IntOfUint32((_543_v1).Cardinality())
			_ = _nwElement0_12
			var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(19))
			_ = _nw89
			(_nw89).ArraySet1(_nwElement0_12, 0)
			(_nw89).ArraySet1(_dafny.IntOfInt64(917), 1)
			(_nw89).ArraySet1(_546_i0, 2)
			(_nw89).ArraySet1(_542_v0, 3)
			(_nw89).ArraySet1((_dafny.Zero).Minus((_dafny.SetOf(_this.F14, _this.F14)).Cardinality()), 4)
			(_nw89).ArraySet1((_549_v6).Cardinality(), 5)
			(_nw89).ArraySet1(_542_v0, 6)
			(_nw89).ArraySet1(_546_i0, 7)
			(_nw89).ArraySet1(_dafny.IntOfInt64(691), 8)
			(_nw89).ArraySet1((_550_v7).Cardinality(), 9)
			(_nw89).ArraySet1(_546_i0, 10)
			(_nw89).ArraySet1(_542_v0, 11)
			(_nw89).ArraySet1(_546_i0, 12)
			(_nw89).ArraySet1(_546_i0, 13)
			(_nw89).ArraySet1(_dafny.IntOfInt64(731), 14)
			(_nw89).ArraySet1(_546_i0, 15)
			(_nw89).ArraySet1((func() _dafny.Int {
				if (_552_v9).Contains(_546_i0) {
					return (_552_v9).Get(_546_i0).(_dafny.Int)
				}
				return Companion_Default___.Fm4(globalState)
			})(), 16)
			(_nw89).ArraySet1(_542_v0, 17)
			(_nw89).ArraySet1(_546_i0, 18)
			_553_v10 = _nw89
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(770), _dafny.ArrayLen((_547_v4), 0))
			_ = _index84
			(_547_v4).ArraySet1(_553_v10, (_index84).Int())
			var _554_v11 _dafny.Map
			_ = _554_v11
			_554_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_546_i0, _this.F14)
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(770), _dafny.ArrayLen((_547_v4), 0))
			_ = _index85
			var _rhs81 _dafny.Array = _553_v10
			_ = _rhs81
			var _rhs82 bool = _this.F14
			_ = _rhs82
			var _rhs83 bool = (!(_this.F14)) == ((func() bool {
				if (_554_v11).Contains(_546_i0) {
					return (_554_v11).Get(_546_i0).(bool)
				}
				return _this.F14
			})())
			_ = _rhs83
			var _lhs59 _dafny.Array = _547_v4
			_ = _lhs59
			var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(770), _dafny.ArrayLen((_547_v4), 0))
			_ = _lhs60
			var _lhs61 *C1 = _this
			_ = _lhs61
			var _lhs62 *C1 = _this
			_ = _lhs62
			(_lhs59).ArraySet1(_rhs81, (_lhs60).Int())
			_lhs61.F14 = _rhs82
			_lhs62.F14 = _rhs83
			var _555_v12 _dafny.Set
			_ = _555_v12
			_555_v12 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Cardinality(), _this.F14)).Cardinality(), (p0).Cardinality())
			var _556_v13 _dafny.Map
			_ = _556_v13
			_556_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_555_v12, _542_v0)
			var _557_v14 _dafny.MultiSet
			_ = _557_v14
			_557_v14 = _dafny.MultiSetOf((_546_i0).Cmp((_556_v13).Cardinality()) >= 0)
			_557_v14 = _557_v14
			(globalState).F2 = _546_i0
			(_this).F14 = (func() bool {
				if _this.F14 {
					return _this.F14
				}
				return _this.F14
			})()
		}
		var _558_v15 _dafny.Map
		_ = _558_v15
		_558_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_542_v0, _this.F14)
		_558_v15 = (_558_v15).Update((_dafny.IntOfUint32((_543_v1).Cardinality())).Plus(_542_v0), _this.F14)
		var _559_v16 _dafny.Array
		_ = _559_v16
		var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
		_ = _nw90
		_559_v16 = _nw90
		var _560_v17 _dafny.Map
		_ = _560_v17
		_560_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14, !(_this.F14))
		var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_559_v16), 0))
		_ = _index86
		(_559_v16).ArraySet1(_560_v17, (_index86).Int())
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_559_v16), 0))
		_ = _index87
		(_559_v16).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14, _this.F14), (_index87).Int())
		var _561_v18 _dafny.Sequence
		_ = _561_v18
		_561_v18 = _dafny.SeqOf(!(_this.F14), _this.F14, _this.F14)
		var _562_v19 _dafny.Array
		_ = _562_v19
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_13
		var _nw91 _dafny.Array
		_ = _nw91
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw91 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) bool = func(_563_i1 _dafny.Int) bool {
				return _this.F14
			}
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw91 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw91).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw91).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_562_v19 = _nw91
		var _564_v20 D3
		_ = _564_v20
		_564_v20 = Companion_D3_.Create_DC10_(true, _542_v0, _this.F14, _562_v19, _558_v15)
		var _pat_let_tv8 = _558_v15
		_ = _pat_let_tv8
		var _pat_let_tv9 = _564_v20
		_ = _pat_let_tv9
		var _pat_let_tv10 = globalState
		_ = _pat_let_tv10
		var _source11 D1 = func(_pat_let10_0 D1) D1 {
			return func(_565_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let11_0 _dafny.CodePoint) D1 {
					return func(_566_dt__update_hcf5_h0 _dafny.CodePoint) D1 {
						return Companion_D1_.Create_DC3_(_566_dt__update_hcf5_h0)
					}(_pat_let11_0)
				}(Companion_Default___.Fm6((_pat_let_tv8).Cardinality(), (_pat_let_tv9).Dtor_cf14(), !(_this.F14), _dafny.CodePoint('a'), _pat_let_tv10))
			}(_pat_let10_0)
		}(Companion_Default___.Fm5(_542_v0, _542_v0, _542_v0, _561_v18, globalState))
		_ = _source11
		if _source11.Is_DC4() {
			var _567___mcc_h0 _dafny.MultiSet = _source11.Get_().(D1_DC4).Cf6
			_ = _567___mcc_h0
			var _568___mcc_h1 _dafny.Int = _source11.Get_().(D1_DC4).Cf7
			_ = _568___mcc_h1
			var _569_cf7 _dafny.Int = _568___mcc_h1
			_ = _569_cf7
			var _570_cf6 _dafny.MultiSet = _567___mcc_h0
			_ = _570_cf6
			var _571_v21 *C0
			_ = _571_v21
			var _nw92 *C0 = New_C0_()
			_ = _nw92
			_nw92.Ctor__()
			_571_v21 = _nw92
			var _572_v22 _dafny.Sequence
			_ = _572_v22
			_572_v22 = _dafny.SeqOf(_562_v19)
			var _573_v23 _dafny.Array
			_ = _573_v23
			var _nwElement0_13 _dafny.Array = _562_v19
			_ = _nwElement0_13
			var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(23))
			_ = _nw93
			(_nw93).ArraySet1(_nwElement0_13, 0)
			(_nw93).ArraySet1((func() _dafny.Array {
				if _this.F14 {
					return _562_v19
				}
				return _562_v19
			})(), 1)
			(_nw93).ArraySet1(_562_v19, 2)
			(_nw93).ArraySet1(_562_v19, 3)
			(_nw93).ArraySet1(_562_v19, 4)
			(_nw93).ArraySet1(_562_v19, 5)
			(_nw93).ArraySet1(_562_v19, 6)
			(_nw93).ArraySet1((_564_v20).Dtor_cf17(), 7)
			(_nw93).ArraySet1(_562_v19, 8)
			(_nw93).ArraySet1(_562_v19, 9)
			(_nw93).ArraySet1(_562_v19, 10)
			(_nw93).ArraySet1(_562_v19, 11)
			(_nw93).ArraySet1((_572_v22).Select((Companion_Default___.SafeIndex(_569_cf7, _dafny.IntOfUint32((_572_v22).Cardinality()))).Uint32()).(_dafny.Array), 12)
			(_nw93).ArraySet1(_562_v19, 13)
			(_nw93).ArraySet1(_562_v19, 14)
			(_nw93).ArraySet1(_562_v19, 15)
			(_nw93).ArraySet1(_562_v19, 16)
			(_nw93).ArraySet1(_562_v19, 17)
			(_nw93).ArraySet1(_562_v19, 18)
			(_nw93).ArraySet1(_562_v19, 19)
			(_nw93).ArraySet1(_562_v19, 20)
			(_nw93).ArraySet1(_562_v19, 21)
			(_nw93).ArraySet1(_562_v19, 22)
			_573_v23 = _nw93
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_573_v23), 0))
			_ = _index88
			(_573_v23).ArraySet1(_562_v19, (_index88).Int())
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_573_v23), 0))
			_ = _index89
			(_573_v23).ArraySet1(_562_v19, (_index89).Int())
			var _574_v24 _dafny.Map
			_ = _574_v24
			_574_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_this.F14, _this.F14, _this.F14, _this.F14, globalState), _562_v19)
			_562_v19 = (func() _dafny.Array {
				if (_574_v24).Contains(_this.F14) {
					return (_574_v24).Get(_this.F14).(_dafny.Array)
				}
				return _562_v19
			})()
			_545_v3 = _545_v3
		} else if _source11.Is_DC5() {
			var _575___mcc_h2 bool = _source11.Get_().(D1_DC5).Cf8
			_ = _575___mcc_h2
			var _576___mcc_h3 bool = _source11.Get_().(D1_DC5).Cf9
			_ = _576___mcc_h3
			var _577_cf9 bool = _576___mcc_h3
			_ = _577_cf9
			var _578_cf8 bool = _575___mcc_h2
			_ = _578_cf8
			var _579_v25 _dafny.Map
			_ = _579_v25
			_579_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_542_v0, _dafny.UnicodeSeqOfUtf8Bytes("vsixj"))
			var _580_v26 _dafny.Sequence
			_ = _580_v26
			_580_v26 = _dafny.UnicodeSeqOfUtf8Bytes("oukgvej")
			var _581_v27 _dafny.Map
			_ = _581_v27
			_581_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_542_v0, (func() _dafny.Int {
				if _578_cf8 {
					return _dafny.IntOfInt64(883)
				}
				return _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_579_v25).Contains(_542_v0) {
						return (_579_v25).Get(_542_v0).(_dafny.Sequence)
					}
					return _580_v26
				})()).Cardinality())
			})())
			_581_v27 = (_581_v27).Update(_542_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lysyusnd")).Cardinality()))
			var _582_v28 _dafny.Map
			_ = _582_v28
			_582_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if _578_cf8 {
					return false
				}
				return _this.F14
			})(), _542_v0)
			_582_v28 = (_582_v28).Update((_this.F14) || (_578_cf8), _dafny.IntOfInt64(177))
			var _583_v29 _dafny.Array
			_ = _583_v29
			var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(15))
			_ = _nw94
			_583_v29 = _nw94
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_583_v29), 0))
			_ = _index90
			(_583_v29).ArraySet1CodePoint(_544_v2, (_index90).Int())
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_583_v29), 0))
			_ = _index91
			(_583_v29).ArraySet1CodePoint(_544_v2, (_index91).Int())
			var _584_v30 *C0
			_ = _584_v30
			var _nw95 *C0 = New_C0_()
			_ = _nw95
			_nw95.Ctor__()
			_584_v30 = _nw95
		} else if _source11.Is_DC3() {
			var _585___mcc_h4 _dafny.CodePoint = _source11.Get_().(D1_DC3).Cf5
			_ = _585___mcc_h4
			var _586_cf5 _dafny.CodePoint = _585___mcc_h4
			_ = _586_cf5
			var _587_v31 *C0
			_ = _587_v31
			var _nw96 *C0 = New_C0_()
			_ = _nw96
			_nw96.Ctor__()
			_587_v31 = _nw96
			(_this).F14 = (_561_v18).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_545_v3).Contains(((_559_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_559_v16), 0))).Int()).(_dafny.Map)).Cardinality()) {
					return (_545_v3).Multiplicity(((_559_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_559_v16), 0))).Int()).(_dafny.Map)).Cardinality())
				}
				return _542_v0
			})(), _dafny.IntOfUint32((_561_v18).Cardinality()))).Uint32()).(bool)
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_562_v19), 0))
			_ = _index92
			(_562_v19).ArraySet1(false, (_index92).Int())
			var _588_v32 _dafny.Map
			_ = _588_v32
			_588_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_562_v19, _542_v0)
			var _589_v33 _dafny.Sequence
			_ = _589_v33
			_589_v33 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_562_v19, _542_v0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_562_v19, (_dafny.Zero).Minus(_542_v0)))
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_562_v19), 0))
			_ = _index93
			var _rhs84 bool = _this.F14
			_ = _rhs84
			var _rhs85 _dafny.Map = ((func() _dafny.Map {
				if false {
					return _588_v32
				}
				return _588_v32
			})()).Merge((_589_v33).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_542_v0), _dafny.IntOfUint32((_589_v33).Cardinality()))).Uint32()).(_dafny.Map))
			_ = _rhs85
			var _rhs86 bool = (_this.F14) == (true)
			_ = _rhs86
			var _rhs87 bool = _this.F14
			_ = _rhs87
			var _rhs88 _dafny.Sequence = Companion_Default___.Fm9(_this.F14, globalState)
			_ = _rhs88
			var _lhs63 _dafny.Array = _562_v19
			_ = _lhs63
			var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_562_v19), 0))
			_ = _lhs64
			var _lhs65 *C1 = _this
			_ = _lhs65
			var _lhs66 *C1 = _this
			_ = _lhs66
			(_lhs63).ArraySet1(_rhs84, (_lhs64).Int())
			_588_v32 = _rhs85
			_lhs65.F14 = _rhs86
			_lhs66.F14 = _rhs87
			r0 = _rhs88
			var _590_v34 _dafny.Sequence
			_ = _590_v34
			_590_v34 = _dafny.UnicodeSeqOfUtf8Bytes("ihdetle")
			var _591_v35 _dafny.Set
			_ = _591_v35
			_591_v35 = _dafny.SetOf((_dafny.Zero).Minus(_542_v0), _542_v0)
			var _592_v36 _dafny.MultiSet
			_ = _592_v36
			_592_v36 = _dafny.MultiSetOf(_590_v34)
			var _593_v37 D1
			_ = _593_v37
			_593_v37 = Companion_D1_.Create_DC4_(_592_v36, _542_v0)
			var _594_v38 _dafny.Map
			_ = _594_v38
			_594_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_562_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_562_v19), 0))).Int()).(bool), _561_v18)
			var _595_v39 _dafny.Array
			_ = _595_v39
			var _nwElement0_14 _dafny.Int = _542_v0
			_ = _nwElement0_14
			var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(17))
			_ = _nw97
			(_nw97).ArraySet1(_nwElement0_14, 0)
			(_nw97).ArraySet1(_542_v0, 1)
			(_nw97).ArraySet1(_542_v0, 2)
			(_nw97).ArraySet1(_542_v0, 3)
			(_nw97).ArraySet1(_dafny.IntOfInt64(137), 4)
			(_nw97).ArraySet1(Companion_Default___.SafeDivisionInt(_542_v0, _542_v0), 5)
			(_nw97).ArraySet1((_dafny.IntOfUint32((_590_v34).Cardinality())).Plus((_591_v35).Cardinality()), 6)
			(_nw97).ArraySet1(_542_v0, 7)
			(_nw97).ArraySet1((_593_v37).Dtor_cf7(), 8)
			(_nw97).ArraySet1((_542_v0).Plus((_594_v38).Cardinality()), 9)
			(_nw97).ArraySet1((func() _dafny.Int {
				if (_545_v3).Contains(_542_v0) {
					return (_545_v3).Multiplicity(_542_v0)
				}
				return _dafny.IntOfInt64(745)
			})(), 10)
			(_nw97).ArraySet1(_542_v0, 11)
			(_nw97).ArraySet1((_542_v0).Plus(_542_v0), 12)
			(_nw97).ArraySet1(Companion_Default___.Fm4(globalState), 13)
			(_nw97).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(386), _542_v0), 14)
			(_nw97).ArraySet1(_542_v0, 15)
			(_nw97).ArraySet1((_543_v1).Select((Companion_Default___.SafeIndex((_591_v35).Cardinality(), _dafny.IntOfUint32((_543_v1).Cardinality()))).Uint32()).(_dafny.Int), 16)
			_595_v39 = _nw97
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_595_v39), 0))
			_ = _index94
			(_595_v39).ArraySet1(_542_v0, (_index94).Int())
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_595_v39), 0))
			_ = _index95
			(_595_v39).ArraySet1(_542_v0, (_index95).Int())
		} else {
			var _596___mcc_h5 D1 = _source11.Get_().(D1_DC6).Cf10
			_ = _596___mcc_h5
			var _597_cf10 D1 = _596___mcc_h5
			_ = _597_cf10
			var _598_v40 _dafny.Array
			_ = _598_v40
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_14
			var _nw98 _dafny.Array
			_ = _nw98
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw98 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) D1 = (func(_599_v2 _dafny.CodePoint) func(_dafny.Int) D1 {
					return func(_600_i2 _dafny.Int) D1 {
						return Companion_D1_.Create_DC6_(Companion_D1_.Create_DC3_(_599_v2))
					}
				})(_544_v2)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw98 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw98).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw98).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_598_v40 = _nw98
			var _601_v41 D2
			_ = _601_v41
			_601_v41 = Companion_D2_.Create_DC7_(_598_v40)
			_601_v41 = _601_v41
			(_this).F14 = (func() bool {
				if _this.F14 {
					return _this.F14
				}
				return _this.F14
			})()
			var _602_v42 D2
			_ = _602_v42
			_602_v42 = Companion_D2_.Create_DC8_(_dafny.IntOfInt64(675))
			var _603_v43 _dafny.MultiSet
			_ = _603_v43
			_603_v43 = _dafny.MultiSetOf(_602_v42)
			(globalState).F2 = ((func() _dafny.MultiSet {
				if _this.F14 {
					return (_dafny.MultiSetOf(_602_v42)).Difference(_dafny.MultiSetOf(_602_v42, _602_v42))
				}
				return ((_dafny.MultiSetOf(_602_v42)).Update(_602_v42, Companion_Default___.Abs(_542_v0))).Union(_603_v43)
			})()).Cardinality()
			var _604_v44 *C0
			_ = _604_v44
			var _nw99 *C0 = New_C0_()
			_ = _nw99
			_nw99.Ctor__()
			_604_v44 = _nw99
		}
		var _605_v45 _dafny.Sequence
		_ = _605_v45
		_605_v45 = _dafny.UnicodeSeqOfUtf8Bytes("nwhkcx")
		var _606_v46 _dafny.Sequence
		_ = _606_v46
		_606_v46 = _dafny.SeqOf(_605_v45, _dafny.Companion_Sequence_.Update(_605_v45, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_542_v0), _dafny.IntOfUint32((_605_v45).Cardinality()))).Uint32(), _544_v2), _605_v45, _605_v45)
		r0 = (_606_v46).Select((Companion_Default___.SafeIndex((_dafny.IntOfUint32((_543_v1).Cardinality())).Times(_542_v0), _dafny.IntOfUint32((_606_v46).Cardinality()))).Uint32()).(_dafny.Sequence)
		var _607_v47 D1
		_ = _607_v47
		_607_v47 = Companion_D1_.Create_DC5_(_this.F14, _dafny.Companion_Sequence_.IsProperPrefixOf(_605_v45, _605_v45))
		r1 = _607_v47
		var _608_v48 _dafny.Sequence
		_ = _608_v48
		_608_v48 = _dafny.SeqOf(_562_v19, _562_v19, _562_v19, _562_v19)
		r2 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_608_v48).Cardinality()), (_542_v0).Times(_dafny.IntOfUint32((_605_v45).Cardinality())))
		return r0, r1, r2
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	F13 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this.F13 = _dafny.Zero
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) Ctor__(f13 _dafny.Int) {
	{
		(_this).F13 = f13
	}
}
func (_this *C2) M3(p0 _dafny.Array, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) {
	{
		var _hi2 _dafny.Int = _this.F13
		_ = _hi2
		for _609_i0 := _this.F13; _609_i0.Cmp(_hi2) < 0; _609_i0 = _609_i0.Plus(_dafny.One) {
			var _610_v0 bool
			_ = _610_v0
			_610_v0 = true
			var _611_v1 _dafny.CodePoint
			_ = _611_v1
			_611_v1 = _dafny.CodePoint('e')
			var _612_v2 _dafny.Map
			_ = _612_v2
			_612_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_611_v1, _610_v0)
			var _613_v3 _dafny.Map
			_ = _613_v3
			_613_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_609_i0, _610_v0)
			var _614_v4 _dafny.Map
			_ = _614_v4
			_614_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_613_v3).Cardinality(), _this.F13)
			(_this).F13 = ((func() _dafny.Int {
				if _610_v0 {
					return _this.F13
				}
				return (_612_v2).Cardinality()
			})()).Plus((_609_i0).Minus((_614_v4).Cardinality()))
			var _615_v5 _dafny.Array
			_ = _615_v5
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_15
			var _nw100 _dafny.Array
			_ = _nw100
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw100 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Sequence = (func(_616_v0 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_617_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_616_v0)
					}
				})(_610_v0)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw100 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw100).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw100).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_615_v5 = _nw100
			_615_v5 = _615_v5
			var _618_v6 _dafny.Sequence
			_ = _618_v6
			_618_v6 = _dafny.SeqOf(false)
			var _619_v7 _dafny.Map
			_ = _619_v7
			_619_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(true, !(_610_v0), _610_v0, (_618_v6).Select((Companion_Default___.SafeIndex(_this.F13, _dafny.IntOfUint32((_618_v6).Cardinality()))).Uint32()).(bool), globalState), _610_v0)
			var _620_v8 _dafny.Sequence
			_ = _620_v8
			_620_v8 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p1).Cardinality()), _609_i0))
			var _621_v9 _dafny.Array
			_ = _621_v9
			var _nwElement0_15 _dafny.Int = _dafny.IntOfUint32((p1).Cardinality())
			_ = _nwElement0_15
			var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(29))
			_ = _nw101
			(_nw101).ArraySet1(_nwElement0_15, 0)
			(_nw101).ArraySet1((p2).Select((Companion_Default___.SafeIndex(_609_i0, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(_dafny.Int), 1)
			(_nw101).ArraySet1(_dafny.IntOfInt64(908), 2)
			(_nw101).ArraySet1(_609_i0, 3)
			(_nw101).ArraySet1((_619_v7).Cardinality(), 4)
			(_nw101).ArraySet1((_this.F13).Minus(Companion_Default___.Fm2(_610_v0, globalState)), 5)
			(_nw101).ArraySet1(_609_i0, 6)
			(_nw101).ArraySet1(_this.F13, 7)
			(_nw101).ArraySet1((_dafny.Zero).Minus(_609_i0), 8)
			(_nw101).ArraySet1(_dafny.IntOfInt64(943), 9)
			(_nw101).ArraySet1(_this.F13, 10)
			(_nw101).ArraySet1(_this.F13, 11)
			(_nw101).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 12)
			(_nw101).ArraySet1(_dafny.IntOfInt64(917), 13)
			(_nw101).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_620_v8, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(681))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}((func(_622_i0 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_623_i2 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_622_i0, (_dafny.SetOf(_this.F13)).Cardinality())
				}
			})(_609_i0))))).Cardinality()), 14)
			(_nw101).ArraySet1((_dafny.IntOfInt64(-545)).Times(_609_i0), 15)
			(_nw101).ArraySet1(_609_i0, 16)
			(_nw101).ArraySet1(_609_i0, 17)
			(_nw101).ArraySet1(_this.F13, 18)
			(_nw101).ArraySet1(_this.F13, 19)
			(_nw101).ArraySet1(_609_i0, 20)
			(_nw101).ArraySet1(_this.F13, 21)
			(_nw101).ArraySet1((_this.F13).Minus(_this.F13), 22)
			(_nw101).ArraySet1(_this.F13, 23)
			(_nw101).ArraySet1(_this.F13, 24)
			(_nw101).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ngpktyvor"), p1)).Cardinality()), 25)
			(_nw101).ArraySet1(_609_i0, 26)
			(_nw101).ArraySet1(_this.F13, 27)
			(_nw101).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_610_v0)).Cardinality()), 28)
			_621_v9 = _nw101
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_621_v9), 0))
			_ = _index96
			(_621_v9).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-155))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_624_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_625_i3 _dafny.Int) _dafny.CodePoint {
					return _624_v1
				}
			})(_611_v1)))).Cardinality()), (_index96).Int())
			var _626_v10 D0
			_ = _626_v10
			_626_v10 = Companion_D0_.Create_DC0_(_610_v0)
			var _627_v11 _dafny.Set
			_ = _627_v11
			_627_v11 = _dafny.SetOf(_626_v10)
			var _628_v12 _dafny.Map
			_ = _628_v12
			_628_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_610_v0, _627_v11)
			var _629_v13 _dafny.Map
			_ = _629_v13
			_629_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_626_v10, _dafny.IntOfInt64(680))
			var _630_v15 _dafny.Map
			_ = _630_v15
			_630_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Set {
				if (_628_v12).Contains(false) {
					return (_628_v12).Get(false).(_dafny.Set)
				}
				return func() _dafny.Set {
					var _coll15 = _dafny.NewBuilder()
					_ = _coll15
					for _iter15 := _dafny.Iterate((_629_v13).Keys().Elements()); ; {
						_compr_15, _ok15 := _iter15()
						if !_ok15 {
							break
						}
						var _631_v14 D0
						_631_v14 = interface{}(_compr_15).(D0)
						if (_629_v13).Contains(_631_v14) {
							_coll15.Add(_631_v14)
						}
					}
					return _coll15.ToSet()
				}()
			})())
			var _632_v16 _dafny.Set
			_ = _632_v16
			_632_v16 = _dafny.SetOf(_610_v0, _610_v0, !(false), _610_v0, _610_v0)
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_621_v9), 0))
			_ = _index97
			var _rhs89 _dafny.Int = ((func() _dafny.Set {
				if (_628_v12).Contains(_610_v0) {
					return (_628_v12).Get(_610_v0).(_dafny.Set)
				}
				return (func() _dafny.Set {
					if _610_v0 {
						return _627_v11
					}
					return _627_v11
				})()
			})()).Cardinality()
			_ = _rhs89
			var _rhs90 bool = (_dafny.SetOf(_610_v0, _610_v0)).IsProperSubsetOf(_632_v16)
			_ = _rhs90
			var _lhs67 _dafny.Array = _621_v9
			_ = _lhs67
			var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_621_v9), 0))
			_ = _lhs68
			(_lhs67).ArraySet1(_rhs89, (_lhs68).Int())
			_610_v0 = _rhs90
			_610_v0 = _610_v0
		}
		var _633_v17 _dafny.Sequence
		_ = _633_v17
		_633_v17 = _dafny.SeqOf(p0)
		(_this).F13 = _dafny.IntOfUint32((_633_v17).Cardinality())
		var _634_v18 bool
		_ = _634_v18
		_634_v18 = true
		var _rhs91 bool = (_this.F13).Cmp(_this.F13) != 0
		_ = _rhs91
		var _rhs92 bool = (Companion_Default___.SafeDivisionInt(_this.F13, _this.F13)).Cmp(_this.F13) > 0
		_ = _rhs92
		_634_v18 = _rhs91
		_634_v18 = _rhs92
		_634_v18 = false
		var _hi3 _dafny.Int = Companion_Default___.SafeModuloInt(_this.F13, _this.F13)
		_ = _hi3
		for _635_i4 := _this.F13; _635_i4.Cmp(_hi3) < 0; _635_i4 = _635_i4.Plus(_dafny.One) {
			var _636_v19 _dafny.Set
			_ = _636_v19
			_636_v19 = _dafny.SetOf(_634_v18, false, true, Companion_Default___.Fm0(_634_v18, _634_v18, _634_v18, _634_v18, globalState))
			if !(!((_636_v19).IsSubsetOf(_636_v19))) {
				var _637_v20 _dafny.Array
				_ = _637_v20
				var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
				_ = _nw102
				_637_v20 = _nw102
				var _638_v21 _dafny.CodePoint
				_ = _638_v21
				_638_v21 = _dafny.CodePoint('l')
				var _639_v22 _dafny.Map
				_ = _639_v22
				_639_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_638_v21, p1)
				var _640_v23 D1
				_ = _640_v23
				_640_v23 = Companion_D1_.Create_DC3_(_638_v21)
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_637_v20), 0))
				_ = _index98
				(_637_v20).ArraySet1((func() _dafny.Sequence {
					if (_639_v22).Contains((_640_v23).Dtor_cf5()) {
						return (_639_v22).Get((_640_v23).Dtor_cf5()).(_dafny.Sequence)
					}
					return p1
				})(), (_index98).Int())
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_637_v20), 0))
				_ = _index99
				(_637_v20).ArraySet1(Companion_Default___.Fm3(_634_v18, p1, _this.F13, globalState), (_index99).Int())
				var _641_v24 _dafny.Array
				_ = _641_v24
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw103
				_641_v24 = _nw103
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_641_v24), 0))
				_ = _index100
				(_641_v24).ArraySet1((_dafny.Zero).Minus(_635_i4), (_index100).Int())
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_641_v24), 0))
				_ = _index101
				(_641_v24).ArraySet1(_635_i4, (_index101).Int())
				var _642_v25 _dafny.Sequence
				_ = _642_v25
				_642_v25 = _dafny.SeqOf(p2)
				_642_v25 = _642_v25
				var _643_v26 _dafny.Map
				_ = _643_v26
				_643_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_634_v18), _634_v18)
				var _644_v27 _dafny.Array
				_ = _644_v27
				var _nwElement0_16 bool = (func() bool {
					if _634_v18 {
						return Companion_Default___.Fm0(_634_v18, _634_v18, _634_v18, !(_634_v18), globalState)
					}
					return _634_v18
				})()
				_ = _nwElement0_16
				var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(21))
				_ = _nw104
				(_nw104).ArraySet1(_nwElement0_16, 0)
				(_nw104).ArraySet1((_636_v19).IsProperSubsetOf(_636_v19), 1)
				(_nw104).ArraySet1(!(_634_v18), 2)
				(_nw104).ArraySet1(_634_v18, 3)
				(_nw104).ArraySet1(_634_v18, 4)
				(_nw104).ArraySet1(_634_v18, 5)
				(_nw104).ArraySet1(Companion_Default___.Fm0(_634_v18, _634_v18, _634_v18, _634_v18, globalState), 6)
				(_nw104).ArraySet1((func() bool {
					if _634_v18 {
						return _634_v18
					}
					return _634_v18
				})(), 7)
				(_nw104).ArraySet1(Companion_Default___.Fm0(_634_v18, false, _634_v18, false, globalState), 8)
				(_nw104).ArraySet1(_634_v18, 9)
				(_nw104).ArraySet1(_634_v18, 10)
				(_nw104).ArraySet1((func() bool {
					if (_643_v26).Contains(_634_v18) {
						return (_643_v26).Get(_634_v18).(bool)
					}
					return _634_v18
				})(), 11)
				(_nw104).ArraySet1((func() bool {
					if (_643_v26).Contains(_634_v18) {
						return (_643_v26).Get(_634_v18).(bool)
					}
					return _634_v18
				})(), 12)
				(_nw104).ArraySet1(_634_v18, 13)
				(_nw104).ArraySet1(_634_v18, 14)
				(_nw104).ArraySet1(_634_v18, 15)
				(_nw104).ArraySet1(_634_v18, 16)
				(_nw104).ArraySet1(_634_v18, 17)
				(_nw104).ArraySet1(!(_634_v18), 18)
				(_nw104).ArraySet1(_634_v18, 19)
				(_nw104).ArraySet1(false, 20)
				_644_v27 = _nw104
				var _nw105 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
				_ = _nw105
				_644_v27 = _nw105
				var _645_v28 D1
				_ = _645_v28
				_645_v28 = Companion_D1_.Create_DC5_(_634_v18, _634_v18)
				var _646_v29 _dafny.Sequence
				_ = _646_v29
				_646_v29 = _dafny.SeqOf(_645_v28)
				var _647_v30 _dafny.Array
				_ = _647_v30
				var _nw106 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(5))
				_ = _nw106
				_647_v30 = _nw106
				var _648_v31 D2
				_ = _648_v31
				_648_v31 = Companion_D2_.Create_DC7_(_647_v30)
				var _649_v32 _dafny.MultiSet
				_ = _649_v32
				_649_v32 = _dafny.MultiSetOf((_637_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_637_v20), 0))).Int()).(_dafny.Sequence), p1)
				var _650_v33 D1
				_ = _650_v33
				_650_v33 = Companion_D1_.Create_DC4_(_649_v32, _635_i4)
				var _651_v34 _dafny.Map
				_ = _651_v34
				_651_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_650_v33, _647_v30)
				var _652_v35 _dafny.Array
				_ = _652_v35
				var _nwElement0_17 _dafny.Array = _647_v30
				_ = _nwElement0_17
				var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(29))
				_ = _nw107
				(_nw107).ArraySet1(_nwElement0_17, 0)
				(_nw107).ArraySet1(_647_v30, 1)
				(_nw107).ArraySet1(_647_v30, 2)
				(_nw107).ArraySet1(_647_v30, 3)
				(_nw107).ArraySet1(_647_v30, 4)
				(_nw107).ArraySet1(_647_v30, 5)
				(_nw107).ArraySet1(_647_v30, 6)
				(_nw107).ArraySet1(_647_v30, 7)
				(_nw107).ArraySet1(_647_v30, 8)
				(_nw107).ArraySet1(_647_v30, 9)
				(_nw107).ArraySet1(_647_v30, 10)
				(_nw107).ArraySet1(_647_v30, 11)
				(_nw107).ArraySet1(_647_v30, 12)
				(_nw107).ArraySet1(_647_v30, 13)
				(_nw107).ArraySet1((func() _dafny.Array {
					if !(_634_v18) {
						return _647_v30
					}
					return _647_v30
				})(), 14)
				(_nw107).ArraySet1((_648_v31).Dtor_cf11(), 15)
				(_nw107).ArraySet1(_647_v30, 16)
				(_nw107).ArraySet1((func() _dafny.Array {
					if (_651_v34).Contains(_650_v33) {
						return (_651_v34).Get(_650_v33).(_dafny.Array)
					}
					return _647_v30
				})(), 17)
				(_nw107).ArraySet1(_647_v30, 18)
				(_nw107).ArraySet1(_647_v30, 19)
				(_nw107).ArraySet1(_647_v30, 20)
				(_nw107).ArraySet1(_647_v30, 21)
				(_nw107).ArraySet1(_647_v30, 22)
				(_nw107).ArraySet1(_647_v30, 23)
				(_nw107).ArraySet1(_647_v30, 24)
				(_nw107).ArraySet1(_647_v30, 25)
				(_nw107).ArraySet1(_647_v30, 26)
				(_nw107).ArraySet1((func() _dafny.Array {
					if _634_v18 {
						return _647_v30
					}
					return _647_v30
				})(), 27)
				(_nw107).ArraySet1(_647_v30, 28)
				_652_v35 = _nw107
				var _rhs93 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqOf(Companion_D1_.Create_DC5_(_634_v18, _634_v18)), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(620), _this.F13), _dafny.IntOfUint32((_dafny.SeqOf(Companion_D1_.Create_DC5_(_634_v18, _634_v18))).Cardinality()))).Uint32(), _645_v28)
				_ = _rhs93
				var _rhs94 bool = !(_634_v18)
				_ = _rhs94
				var _rhs95 _dafny.Array = _652_v35
				_ = _rhs95
				_646_v29 = _rhs93
				_634_v18 = _rhs94
				_652_v35 = _rhs95
			} else {
				var _653_v36 _dafny.Map
				_ = _653_v36
				_653_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(223), _634_v18)
				var _654_v37 D3
				_ = _654_v37
				_654_v37 = Companion_D3_.Create_DC10_(_634_v18, _this.F13, _634_v18, p0, _653_v36)
				var _655_v38 _dafny.CodePoint
				_ = _655_v38
				_655_v38 = _dafny.CodePoint('i')
				var _656_v39 _dafny.Map
				_ = _656_v39
				_656_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_654_v37).Dtor_cf17(), _655_v38)
				_656_v39 = (_656_v39).Update(p0, (func() _dafny.CodePoint {
					if _634_v18 {
						return _655_v38
					}
					return _655_v38
				})())
				var _657_v40 _dafny.Array
				_ = _657_v40
				var _nw108 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw108
				_657_v40 = _nw108
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_657_v40), 0))
				_ = _index102
				(_657_v40).ArraySet1(_635_i4, (_index102).Int())
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_657_v40), 0))
				_ = _index103
				var _rhs96 _dafny.Int = _this.F13
				_ = _rhs96
				var _rhs97 _dafny.Array = _657_v40
				_ = _rhs97
				var _rhs98 _dafny.Int = _this.F13
				_ = _rhs98
				var _lhs69 *C2 = _this
				_ = _lhs69
				var _lhs70 *GlobalState = globalState
				_ = _lhs70
				var _lhs71 _dafny.Array = _657_v40
				_ = _lhs71
				var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_657_v40), 0))
				_ = _lhs72
				_lhs69.F13 = _rhs96
				_lhs70.F10 = _rhs97
				(_lhs71).ArraySet1(_rhs98, (_lhs72).Int())
				var _658_v41 _dafny.Array
				_ = _658_v41
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_16
				var _nw109 _dafny.Array
				_ = _nw109
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw109 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) bool = (func(_659_v18 bool) func(_dafny.Int) bool {
						return func(_660_i5 _dafny.Int) bool {
							return _659_v18
						}
					})(_634_v18)
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw109 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw109).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw109).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_658_v41 = _nw109
				_658_v41 = _658_v41
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_657_v40), 0))
				_ = _index104
				var _rhs99 bool = _634_v18
				_ = _rhs99
				var _rhs100 _dafny.Int = (_635_i4).Times(Companion_Default___.Fm2(_634_v18, globalState))
				_ = _rhs100
				var _rhs101 _dafny.Int = (_this.F13).Plus(((_657_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_657_v40), 0))).Int()).(_dafny.Int)).Times(_this.F13))
				_ = _rhs101
				var _rhs102 _dafny.Int = _635_i4
				_ = _rhs102
				var _rhs103 _dafny.Int = _this.F13
				_ = _rhs103
				var _lhs73 *C2 = _this
				_ = _lhs73
				var _lhs74 *GlobalState = globalState
				_ = _lhs74
				var _lhs75 *GlobalState = globalState
				_ = _lhs75
				var _lhs76 _dafny.Array = _657_v40
				_ = _lhs76
				var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_657_v40), 0))
				_ = _lhs77
				_634_v18 = _rhs99
				_lhs73.F13 = _rhs100
				_lhs74.F2 = _rhs101
				_lhs75.F2 = _rhs102
				(_lhs76).ArraySet1(_rhs103, (_lhs77).Int())
				var _661_v42 _dafny.Map
				_ = _661_v42
				_661_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(458), _this.F13), (_dafny.IntOfUint32((p2).Cardinality())).Plus(_635_i4))
				_661_v42 = _661_v42
			}
			var _662_v43 _dafny.Array
			_ = _662_v43
			var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
			_ = _nw110
			_662_v43 = _nw110
			_662_v43 = _662_v43
			var _663_v44 _dafny.CodePoint
			_ = _663_v44
			_663_v44 = _dafny.CodePoint('a')
			var _664_v45 _dafny.Set
			_ = _664_v45
			_664_v45 = _dafny.SetOf(_663_v44)
			var _665_v46 _dafny.Map
			_ = _665_v46
			_665_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p1).Cardinality()), _634_v18)
			var _666_v47 D3
			_ = _666_v47
			_666_v47 = Companion_D3_.Create_DC10_(_634_v18, (_664_v45).Cardinality(), _634_v18, (_633_v17).Select((Companion_Default___.SafeIndex(_this.F13, _dafny.IntOfUint32((_633_v17).Cardinality()))).Uint32()).(_dafny.Array), _665_v46)
			if (_666_v47).Dtor_cf16() {
				var _667_v48 _dafny.Map
				_ = _667_v48
				_667_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_634_v18, (_634_v18) || (_634_v18))
				_634_v18 = (func() bool {
					if (_667_v48).Contains(!(!(_634_v18) || (_634_v18))) {
						return (_667_v48).Get(!(!(_634_v18) || (_634_v18))).(bool)
					}
					return false
				})()
				(_this).F13 = _this.F13
				var _668_v49 _dafny.Array
				_ = _668_v49
				var _nw111 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
				_ = _nw111
				_668_v49 = _nw111
				var _nw112 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw112
				_668_v49 = _nw112
				var _669_v50 _dafny.MultiSet
				_ = _669_v50
				_669_v50 = _dafny.MultiSetOf(_635_i4)
				var _rhs104 _dafny.Int = _this.F13
				_ = _rhs104
				var _rhs105 _dafny.Int = _this.F13
				_ = _rhs105
				var _rhs106 bool = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_635_i4)), _634_v18)).Merge((_665_v46).Update(_this.F13, _634_v18))).Contains((func() _dafny.Int {
					if (_669_v50).Contains(Companion_Default___.Fm2(true, globalState)) {
						return (_669_v50).Multiplicity(Companion_Default___.Fm2(true, globalState))
					}
					return _635_i4
				})())
				_ = _rhs106
				var _rhs107 _dafny.CodePoint = _dafny.CodePoint('c')
				_ = _rhs107
				var _rhs108 bool = _634_v18
				_ = _rhs108
				var _lhs78 *C2 = _this
				_ = _lhs78
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				_lhs78.F13 = _rhs104
				_lhs79.F2 = _rhs105
				_634_v18 = _rhs106
				_663_v44 = _rhs107
				_634_v18 = _rhs108
				var _670_v51 _dafny.Array
				_ = _670_v51
				var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
				_ = _nw113
				_670_v51 = _nw113
				var _671_v52 _dafny.MultiSet
				_ = _671_v52
				_671_v52 = _dafny.MultiSetOf(_634_v18)
				Companion_Default___.M0(p0, _this.F13, _670_v51, (_671_v52).Cardinality(), globalState)
			} else {
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((p0), 0))
				_ = _index105
				(p0).ArraySet1(_634_v18, (_index105).Int())
				var _672_v53 _dafny.Array
				_ = _672_v53
				var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
				_ = _nw114
				_672_v53 = _nw114
				var _673_v54 _dafny.Map
				_ = _673_v54
				_673_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_672_v53), 0))
				_ = _index106
				(_672_v53).ArraySet1(_673_v54, (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((p0), 0))
				_ = _index107
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_672_v53), 0))
				_ = _index108
				var _rhs109 bool = _634_v18
				_ = _rhs109
				var _rhs110 _dafny.Map = _673_v54
				_ = _rhs110
				var _lhs80 _dafny.Array = p0
				_ = _lhs80
				var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((p0), 0))
				_ = _lhs81
				var _lhs82 _dafny.Array = _672_v53
				_ = _lhs82
				var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_672_v53), 0))
				_ = _lhs83
				(_lhs80).ArraySet1(_rhs109, (_lhs81).Int())
				(_lhs82).ArraySet1(_rhs110, (_lhs83).Int())
				var _674_v55 _dafny.Array
				_ = _674_v55
				var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
				_ = _nw115
				_674_v55 = _nw115
				var _675_v56 _dafny.Array
				_ = _675_v56
				var _nw116 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw116
				_675_v56 = _nw116
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_674_v55), 0))
				_ = _index109
				(_674_v55).ArraySet1(_675_v56, (_index109).Int())
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_674_v55), 0))
				_ = _index110
				(_674_v55).ArraySet1(_675_v56, (_index110).Int())
				var _676_v57 _dafny.MultiSet
				_ = _676_v57
				_676_v57 = _dafny.MultiSetOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((p0), 0))).Int()).(bool), true, !((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((p0), 0))).Int()).(bool)))
				_676_v57 = _676_v57
				var _677_v58 *C1
				_ = _677_v58
				var _nw117 *C1 = New_C1_()
				_ = _nw117
				_nw117.Ctor__(_634_v18)
				_677_v58 = _nw117
				var _678_v59 _dafny.Sequence
				_ = _678_v59
				_678_v59 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-930))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg27 _dafny.Int) interface{} {
						return coer27(arg27)
					}
				}((func(_679_v44 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_680_i6 _dafny.Int) _dafny.CodePoint {
						return _679_v44
					}
				})(_663_v44))), p1)
				var _681_v60 _dafny.Map
				_ = _681_v60
				_681_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((p0), 0))).Int()).(bool), true)
				var _682_v61 _dafny.Map
				_ = _682_v61
				_682_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_677_v58.F14, (_681_v60).Cardinality())
				var _683_v62 _dafny.Set
				_ = _683_v62
				_683_v62 = _dafny.SetOf(_dafny.IntOfUint32((p2).Cardinality()), (func() _dafny.Int {
					if (_682_v61).Contains(Companion_Default___.Fm0(_634_v18, _677_v58.F14, _677_v58.F14, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState)) {
						return (_682_v61).Get(Companion_Default___.Fm0(_634_v18, _677_v58.F14, _677_v58.F14, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState)).(_dafny.Int)
					}
					return _635_i4
				})(), _this.F13)
				var _684_v63 _dafny.Array
				_ = _684_v63
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_17
				var _nw118 _dafny.Array
				_ = _nw118
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw118 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.Int = (func(_685_i4 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_686_i7 _dafny.Int) _dafny.Int {
							return (_686_i7).Minus(_685_i4)
						}
					})(_635_i4)
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw118 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw118).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw118).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_684_v63 = _nw118
				var _687_v64 _dafny.Sequence
				_ = _687_v64
				_687_v64 = _dafny.SeqOf(_684_v63, _684_v63, _684_v63, _684_v63)
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((p0), 0))
				_ = _index111
				var _rhs111 bool = (_636_v19).IsDisjointFrom(_dafny.SetOf(_634_v18))
				_ = _rhs111
				var _rhs112 _dafny.CodePoint = (p1).Select((Companion_Default___.SafeIndex((_683_v62).Cardinality(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs112
				var _rhs113 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_687_v64, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_684_v63, _684_v63), _687_v64))).Cardinality())
				_ = _rhs113
				var _rhs114 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_678_v59, (Companion_Default___.SafeIndex(_635_i4, _dafny.IntOfUint32((_678_v59).Cardinality()))).Uint32(), p1), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F13), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_678_v59, (Companion_Default___.SafeIndex(_635_i4, _dafny.IntOfUint32((_678_v59).Cardinality()))).Uint32(), p1)).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Concatenate(p1, p1))
				_ = _rhs114
				var _rhs115 bool = false
				_ = _rhs115
				var _lhs84 *C2 = _this
				_ = _lhs84
				var _lhs85 _dafny.Array = p0
				_ = _lhs85
				var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((p0), 0))
				_ = _lhs86
				_634_v18 = _rhs111
				_663_v44 = _rhs112
				_lhs84.F13 = _rhs113
				_678_v59 = _rhs114
				(_lhs85).ArraySet1(_rhs115, (_lhs86).Int())
			}
			var _688_v65 _dafny.Map
			_ = _688_v65
			_688_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_634_v18, _636_v19)
			var _689_v66 D4
			_ = _689_v66
			_689_v66 = Companion_D4_.Create_DC12_(_636_v19)
			var _690_v67 _dafny.Map
			_ = _690_v67
			_690_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_635_i4, (_689_v66).Dtor_cf20())
			var _691_v68 _dafny.MultiSet
			_ = _691_v68
			_691_v68 = _dafny.MultiSetOf(_this.F13)
			var _692_v69 _dafny.Sequence
			_ = _692_v69
			_692_v69 = _dafny.SeqOf(_634_v18, false)
			var _693_v70 _dafny.Map
			_ = _693_v70
			_693_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_634_v18, (_692_v69).Select((Companion_Default___.SafeIndex(_635_i4, _dafny.IntOfUint32((_692_v69).Cardinality()))).Uint32()).(bool))
			var _694_v71 _dafny.Map
			_ = _694_v71
			_694_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_634_v18), _663_v44)
			var _695_v72 _dafny.Array
			_ = _695_v72
			var _nwElement0_18 _dafny.Int = _dafny.IntOfInt64(647)
			_ = _nwElement0_18
			var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(29))
			_ = _nw119
			(_nw119).ArraySet1(_nwElement0_18, 0)
			(_nw119).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if _634_v18 {
					return p1
				}
				return p1
			})(), (Companion_Default___.SafeIndex(_this.F13, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if _634_v18 {
					return p1
				}
				return p1
			})()).Cardinality()))).Uint32(), _663_v44)).Cardinality()), 1)
			(_nw119).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(466))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}(func(_696_i8 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			}))).Cardinality()), (p2).Select((Companion_Default___.SafeIndex(_635_i4, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(_dafny.Int)), 2)
			(_nw119).ArraySet1((_this.F13).Minus((_636_v19).Cardinality()), 3)
			(_nw119).ArraySet1(((func() _dafny.Set {
				if (_688_v65).Contains(_634_v18) {
					return (_688_v65).Get(_634_v18).(_dafny.Set)
				}
				return (func() _dafny.Set {
					if (_690_v67).Contains(_this.F13) {
						return (_690_v67).Get(_this.F13).(_dafny.Set)
					}
					return _636_v19
				})()
			})()).Cardinality(), 4)
			(_nw119).ArraySet1(_635_i4, 5)
			(_nw119).ArraySet1(Companion_Default___.SafeModuloInt(_this.F13, _dafny.IntOfInt64(182)), 6)
			(_nw119).ArraySet1(Companion_Default___.SafeModuloInt(_this.F13, _this.F13), 7)
			(_nw119).ArraySet1(_635_i4, 8)
			(_nw119).ArraySet1(_this.F13, 9)
			(_nw119).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 10)
			(_nw119).ArraySet1((func() _dafny.Int {
				if (_691_v68).Contains(_dafny.IntOfUint32((_692_v69).Cardinality())) {
					return (_691_v68).Multiplicity(_dafny.IntOfUint32((_692_v69).Cardinality()))
				}
				return _635_i4
			})(), 11)
			(_nw119).ArraySet1(_635_i4, 12)
			(_nw119).ArraySet1(_this.F13, 13)
			(_nw119).ArraySet1(_635_i4, 14)
			(_nw119).ArraySet1(_635_i4, 15)
			(_nw119).ArraySet1(_dafny.IntOfInt64(773), 16)
			(_nw119).ArraySet1((_693_v70).Cardinality(), 17)
			(_nw119).ArraySet1(Companion_Default___.Fm4(globalState), 18)
			(_nw119).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, p1)).Cardinality()), 19)
			(_nw119).ArraySet1(_this.F13, 20)
			(_nw119).ArraySet1(_635_i4, 21)
			(_nw119).ArraySet1(_635_i4, 22)
			(_nw119).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetOf(_663_v44, _663_v44, _663_v44, (func() _dafny.CodePoint {
				if (_694_v71).Contains(_634_v18) {
					return (_694_v71).Get(_634_v18).(_dafny.CodePoint)
				}
				return _663_v44
			})(), _663_v44)).Cardinality())), 23)
			(_nw119).ArraySet1(_this.F13, 24)
			(_nw119).ArraySet1(_this.F13, 25)
			(_nw119).ArraySet1(_635_i4, 26)
			(_nw119).ArraySet1(_635_i4, 27)
			(_nw119).ArraySet1(_this.F13, 28)
			_695_v72 = _nw119
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_695_v72), 0))
			_ = _index112
			(_695_v72).ArraySet1(_this.F13, (_index112).Int())
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_695_v72), 0))
			_ = _index113
			(_695_v72).ArraySet1(_this.F13, (_index113).Int())
		}
		_634_v18 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F13), _dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex(_this.F13, _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), _this.F13)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(802))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_697_i9 _dafny.Int) _dafny.Int {
			return _this.F13
		})))
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f11 bool
	_f12 _dafny.Int
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f11 = false
	_this._f12 = _dafny.Zero
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

func (_this *C3) Ctor__(f11 bool, f12 _dafny.Int) {
	{
		(_this)._f11 = f11
		(_this)._f12 = f12
	}
}
func (_this *C3) Fm1(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ngc")).Cardinality()), (_dafny.MultiSetOf((_this).F11())).Cardinality()), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ttwonxung")).Cardinality()), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(173))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}(func(_698_i0 _dafny.Int) _dafny.Int {
			return (_this).F12()
		}))))).Cardinality()))
	}
}
func (_this *C3) M1(p0 bool, globalState *GlobalState) {
	{
		var _699_v0 D0
		_ = _699_v0
		_699_v0 = Companion_D0_.Create_DC0_(p0)
		var _700_v1 D0
		_ = _700_v1
		_700_v1 = Companion_D0_.Create_DC2_(_699_v0)
		var _701_v2 _dafny.Array
		_ = _701_v2
		var _nw120 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(14))
		_ = _nw120
		_701_v2 = _nw120
		var _702_v3 _dafny.Map
		_ = _702_v3
		_702_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_700_v1, _701_v2)
		if (_702_v3).Contains(Companion_D0_.Create_DC2_(_699_v0)) {
			var _703_v4 *C2
			_ = _703_v4
			var _nw121 *C2 = New_C2_()
			_ = _nw121
			_nw121.Ctor__((_dafny.Zero).Minus((_this).F12()))
			_703_v4 = _nw121
			var _704_v5 _dafny.MultiSet
			_ = _704_v5
			_704_v5 = _dafny.MultiSetOf((_703_v4.F13).Minus((_this).F12()), (_this).F12())
			_704_v5 = _704_v5
			var _705_v6 _dafny.Array
			_ = _705_v6
			var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
			_ = _nw122
			_705_v6 = _nw122
			var _706_v7 _dafny.Sequence
			_ = _706_v7
			_706_v7 = _dafny.SeqOf(p0)
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_705_v6), 0))
			_ = _index114
			(_705_v6).ArraySet1(_dafny.IntOfUint32((_706_v7).Cardinality()), (_index114).Int())
			var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_705_v6), 0))
			_ = _index115
			(_705_v6).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-878), (Companion_Default___.Fm10(globalState)).Cardinality()), (_index115).Int())
			var _707_v8 bool
			_ = _707_v8
			_707_v8 = false
			var _708_v9 _dafny.Set
			_ = _708_v9
			_708_v9 = _dafny.SetOf((_705_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_705_v6), 0))).Int()).(_dafny.Int))
			_707_v8 = !(!(_708_v9).Equals(_708_v9))
			var _709_v10 _dafny.CodePoint
			_ = _709_v10
			_709_v10 = _dafny.CodePoint('e')
			var _710_v11 _dafny.Sequence
			_ = _710_v11
			_710_v11 = _dafny.UnicodeSeqOfUtf8Bytes("wdfkv")
			if Companion_Default___.Fm0(_dafny.Companion_Sequence_.Contains(_710_v11, _709_v10), (func() bool {
				if (_this).F11() {
					return _707_v8
				}
				return (_this).F11()
			})(), (_this).F11(), (_this).F11(), globalState) {
				(globalState).F2 = ((_705_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_705_v6), 0))).Int()).(_dafny.Int)).Times(_703_v4.F13)
				_707_v8 = (_dafny.IntOfUint32((_706_v7).Cardinality())).Cmp((_this).F12()) < 0
				var _711_v12 _dafny.Array
				_ = _711_v12
				var _nw123 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw123
				_711_v12 = _nw123
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_711_v12), 0))
				_ = _index116
				(_711_v12).ArraySet1(p0, (_index116).Int())
				var _712_v13 _dafny.MultiSet
				_ = _712_v13
				_712_v13 = _dafny.MultiSetOf(_706_v7)
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_711_v12), 0))
				_ = _index117
				(_711_v12).ArraySet1((_712_v13).IsSubsetOf((_712_v13).Union(_712_v13)), (_index117).Int())
				(_703_v4).F13 = (_705_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_705_v6), 0))).Int()).(_dafny.Int)
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_705_v6), 0))
				_ = _index118
				(_705_v6).ArraySet1(((_this).F12()).Times(_703_v4.F13), (_index118).Int())
			} else {
				_700_v1 = _700_v1
				var _713_v14 *C1
				_ = _713_v14
				var _nw124 *C1 = New_C1_()
				_ = _nw124
				_nw124.Ctor__((_704_v5).IsProperSubsetOf(_704_v5))
				_713_v14 = _nw124
				var _714_v15 _dafny.Map
				_ = _714_v15
				_714_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_705_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_705_v6), 0))).Int()).(_dafny.Int), _709_v10)
				var _715_v16 _dafny.Map
				_ = _715_v16
				_715_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
					if (_714_v15).Contains((_this).F12()) {
						return (_714_v15).Get((_this).F12()).(_dafny.CodePoint)
					}
					return _709_v10
				})(), _dafny.Companion_Sequence_.Contains(_710_v11, _709_v10))
				_707_v8 = (func() bool {
					if (_715_v16).Contains(_dafny.CodePoint('o')) {
						return (_715_v16).Get(_dafny.CodePoint('o')).(bool)
					}
					return (_this).F11()
				})()
				var _716_v17 _dafny.Array
				_ = _716_v17
				var _nwElement0_19 bool = (_707_v8) == ((_706_v7).Select((Companion_Default___.SafeIndex(_703_v4.F13, _dafny.IntOfUint32((_706_v7).Cardinality()))).Uint32()).(bool))
				_ = _nwElement0_19
				var _nw125 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(9))
				_ = _nw125
				(_nw125).ArraySet1(_nwElement0_19, 0)
				(_nw125).ArraySet1((_this).F11(), 1)
				(_nw125).ArraySet1(!(_713_v14.F14), 2)
				(_nw125).ArraySet1((_this).F11(), 3)
				(_nw125).ArraySet1((_706_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-640), _dafny.IntOfUint32((_706_v7).Cardinality()))).Uint32()).(bool), 4)
				(_nw125).ArraySet1(_713_v14.F14, 5)
				(_nw125).ArraySet1(p0, 6)
				(_nw125).ArraySet1(_713_v14.F14, 7)
				(_nw125).ArraySet1((_this).F11(), 8)
				_716_v17 = _nw125
				_716_v17 = _716_v17
				var _717_v18 _dafny.Array
				_ = _717_v18
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_18
				var _nw126 _dafny.Array
				_ = _nw126
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw126 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Sequence = (func(_718_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_719_i0 _dafny.Int) _dafny.Sequence {
							return _718_v7
						}
					})(_706_v7)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw126 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw126).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw126).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_717_v18 = _nw126
				_717_v18 = _717_v18
			}
		} else {
			var _720_v19 _dafny.Sequence
			_ = _720_v19
			_720_v19 = _dafny.SeqOf((_this).F12(), (_this).F12())
			var _721_v20 _dafny.Sequence
			_ = _721_v20
			_721_v20 = _dafny.SeqOf((_this).F12(), _dafny.IntOfUint32((_720_v19).Cardinality()), _dafny.IntOfInt64(515))
			var _722_v21 _dafny.MultiSet
			_ = _722_v21
			_722_v21 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.IsProperPrefixOf(_720_v19, _721_v20))
			_722_v21 = (func() _dafny.MultiSet {
				if p0 {
					return _722_v21
				}
				return _722_v21
			})()
			var _723_v22 *C2
			_ = _723_v22
			var _nw127 *C2 = New_C2_()
			_ = _nw127
			_nw127.Ctor__(Companion_Default___.Fm4(globalState))
			_723_v22 = _nw127
			_723_v22 = _723_v22
			var _724_v23 bool
			_ = _724_v23
			_724_v23 = true
			_724_v23 = _724_v23
			var _725_v24 _dafny.Map
			_ = _725_v24
			_725_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_723_v22.F13, ((_this).F12()).Minus((_this).F12()))
			var _rhs116 bool = _724_v23
			_ = _rhs116
			var _rhs117 bool = p0
			_ = _rhs117
			var _rhs118 bool = ((_this).F12()).Cmp(_dafny.IntOfInt64(80)) > 0
			_ = _rhs118
			var _rhs119 bool = p0
			_ = _rhs119
			var _rhs120 _dafny.Map = (_725_v24).Merge(_725_v24)
			_ = _rhs120
			_724_v23 = _rhs116
			_724_v23 = _rhs117
			_724_v23 = _rhs118
			_724_v23 = _rhs119
			_725_v24 = _rhs120
			var _726_v25 *C1
			_ = _726_v25
			var _nw128 *C1 = New_C1_()
			_ = _nw128
			_nw128.Ctor__((_this).F11())
			_726_v25 = _nw128
		}
		var _727_v26 bool
		_ = _727_v26
		_727_v26 = false
		_727_v26 = !(p0)
		var _728_i1 _dafny.Int
		_ = _728_i1
		_728_i1 = _dafny.Zero
		{
			for p0 {
				{
					if (_728_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_728_i1 = (_728_i1).Plus(_dafny.One)
					var _729_v27 _dafny.MultiSet
					_ = _729_v27
					_729_v27 = _dafny.MultiSetOf((_this).F11(), (_this).F11())
					var _730_v28 _dafny.Sequence
					_ = _730_v28
					_730_v28 = _dafny.UnicodeSeqOfUtf8Bytes("dcu")
					var _731_v29 _dafny.CodePoint
					_ = _731_v29
					_731_v29 = _dafny.CodePoint('h')
					var _732_v30 _dafny.MultiSet
					_ = _732_v30
					_732_v30 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_730_v28, (Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_730_v28).Cardinality()))).Uint32(), _731_v29), _dafny.UnicodeSeqOfUtf8Bytes("rbbmc"))
					var _733_v31 bool
					_ = _733_v31
					var _734_v32 _dafny.Map
					_ = _734_v32
					var _735_v33 bool
					_ = _735_v33
					var _out27 bool
					_ = _out27
					var _out28 _dafny.Map
					_ = _out28
					var _out29 bool
					_ = _out29
					_out27, _out28, _out29 = (_this).M2(((_this).F12()).Cmp((_729_v27).Cardinality()) >= 0, (_732_v30).IsProperSubsetOf(_732_v30), globalState)
					_733_v31 = _out27
					_734_v32 = _out28
					_735_v33 = _out29
					var _736_v34 _dafny.Map
					_ = _736_v34
					_736_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_727_v26), _dafny.SetOf((_this).F11(), _733_v31, _733_v31))
					var _737_v35 _dafny.Set
					_ = _737_v35
					_737_v35 = _dafny.SetOf((_this).F11())
					var _738_v36 D4
					_ = _738_v36
					_738_v36 = Companion_D4_.Create_DC12_(_737_v35)
					_736_v34 = (_736_v34).Update(p0, (_738_v36).Dtor_cf20())
					_730_v28 = _dafny.Companion_Sequence_.Concatenate(_730_v28, _dafny.UnicodeSeqOfUtf8Bytes("ycmdjta"))
					var _739_v37 _dafny.Map
					_ = _739_v37
					_739_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(552), _727_v26)
					(globalState).F2 = ((_this).F12()).Plus(Companion_Default___.SafeModuloInt((_this).F12(), (_739_v37).Cardinality()))
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _hi4 _dafny.Int = (_dafny.Zero).Minus((_this).F12())
		_ = _hi4
		for _740_i2 := Companion_Default___.SafeModuloInt((_this).F12(), (_this).F12()); _740_i2.Cmp(_hi4) < 0; _740_i2 = _740_i2.Plus(_dafny.One) {
			var _741_v38 _dafny.Sequence
			_ = _741_v38
			_741_v38 = _dafny.SeqOf(true)
			_741_v38 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_741_v38, (Companion_Default___.SafeIndex(_740_i2, _dafny.IntOfUint32((_741_v38).Cardinality()))).Uint32(), p0), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_741_v38, (Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_741_v38).Cardinality()))).Uint32(), (_this).F11()), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_740_i2), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_741_v38, (Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_741_v38).Cardinality()))).Uint32(), (_this).F11())).Cardinality()))).Uint32(), !((_this).F11())))
			var _742_v39 _dafny.Array
			_ = _742_v39
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_19
			var _nw129 _dafny.Array
			_ = _nw129
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw129 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) bool = (func(_743_v26 bool) func(_dafny.Int) bool {
					return func(_744_i3 _dafny.Int) bool {
						return _743_v26
					}
				})(_727_v26)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw129 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw129).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw129).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_742_v39 = _nw129
			var _745_v40 _dafny.Map
			_ = _745_v40
			_745_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_742_v39, _742_v39)
			var _746_v41 _dafny.Sequence
			_ = _746_v41
			_746_v41 = _dafny.SeqOf((_this).F12())
			var _747_v42 _dafny.MultiSet
			_ = _747_v42
			_747_v42 = _dafny.MultiSetOf((_this).F12(), _dafny.IntOfUint32((_746_v41).Cardinality()))
			var _748_v43 _dafny.Map
			_ = _748_v43
			_748_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_727_v26, _747_v42)
			var _749_v44 _dafny.Map
			_ = _749_v44
			_749_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F11())
			var _750_v45 _dafny.Map
			_ = _750_v45
			_750_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), p0)
			var _751_v46 _dafny.Map
			_ = _751_v46
			_751_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_749_v44).Cardinality(), (_750_v45).Cardinality())
			var _752_v47 _dafny.Array
			_ = _752_v47
			var _nwElement0_20 _dafny.Int = (_this).F12()
			_ = _nwElement0_20
			var _nw130 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(9))
			_ = _nw130
			(_nw130).ArraySet1(_nwElement0_20, 0)
			(_nw130).ArraySet1(_740_i2, 1)
			(_nw130).ArraySet1((_751_v46).Cardinality(), 2)
			(_nw130).ArraySet1(_740_i2, 3)
			(_nw130).ArraySet1((_this).F12(), 4)
			(_nw130).ArraySet1(_740_i2, 5)
			(_nw130).ArraySet1((_this).F12(), 6)
			(_nw130).ArraySet1(_740_i2, 7)
			(_nw130).ArraySet1(_740_i2, 8)
			_752_v47 = _nw130
			var _753_v48 _dafny.CodePoint
			_ = _753_v48
			_753_v48 = _dafny.CodePoint('e')
			var _754_v49 _dafny.Sequence
			_ = _754_v49
			_754_v49 = _dafny.SeqOf(_753_v48)
			var _755_v50 _dafny.Sequence
			_ = _755_v50
			_755_v50 = _dafny.SeqOf(_747_v42)
			var _756_v51 D4
			_ = _756_v51
			_756_v51 = Companion_D4_.Create_DC13_(_753_v48)
			var _757_v52 _dafny.MultiSet
			_ = _757_v52
			_757_v52 = _dafny.MultiSetOf(_756_v51)
			var _758_v53 _dafny.MultiSet
			_ = _758_v53
			_758_v53 = _757_v52
			var _759_v54 _dafny.Set
			_ = _759_v54
			_759_v54 = _dafny.SetOf(((func() _dafny.MultiSet {
				if (_748_v43).Contains(_727_v26) {
					return (_748_v43).Get(_727_v26).(_dafny.MultiSet)
				}
				return _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_740_i2, _752_v47)).Cardinality())
			})()).Update(_dafny.IntOfUint32((_754_v49).Cardinality()), Companion_Default___.Abs(_dafny.IntOfInt64(-13))), (_747_v42).Union(_dafny.MultiSetOf((_this).F12(), (_this).F12())), _747_v42, (_755_v50).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_758_v53), (_this).F11())).Cardinality(), _740_i2, (_this).F12())).Cardinality(), _dafny.IntOfUint32((_755_v50).Cardinality()))).Uint32()).(_dafny.MultiSet), (_747_v42).Difference(_747_v42))
			var _760_v55 _dafny.Set
			_ = _760_v55
			_760_v55 = _dafny.SetOf((_this).F11())
			var _rhs121 bool = ((_760_v55).Intersection(_760_v55)).IsDisjointFrom(_760_v55)
			_ = _rhs121
			var _rhs122 _dafny.Map = _745_v40
			_ = _rhs122
			var _rhs123 _dafny.Set = _759_v54
			_ = _rhs123
			_727_v26 = _rhs121
			_745_v40 = _rhs122
			_759_v54 = _rhs123
			_750_v45 = (_750_v45).Update(((_this).F12()).Minus((_this).F12()), _727_v26)
			_727_v26 = true
		}
		var _761_v56 _dafny.CodePoint
		_ = _761_v56
		_761_v56 = _dafny.CodePoint('g')
		if Companion_Default___.Fm0(!_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-689))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}((func(_762_v56 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_763_i4 _dafny.Int) _dafny.CodePoint {
				return _762_v56
			}
		})(_761_v56))), _761_v56), _727_v26, (_727_v26) && ((_this).F11()), (_this).F11(), globalState) {
			var _764_v57 *C2
			_ = _764_v57
			var _nw131 *C2 = New_C2_()
			_ = _nw131
			_nw131.Ctor__(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("my")).Cardinality()))
			_764_v57 = _nw131
			if (_this).F11() {
				_727_v26 = !(_727_v26) || (!(!(p0)))
				var _765_v58 _dafny.Sequence
				_ = _765_v58
				_765_v58 = _dafny.UnicodeSeqOfUtf8Bytes("xoku")
				var _766_v59 _dafny.Set
				_ = _766_v59
				_766_v59 = _dafny.SetOf(p0, !(!((_this).F11())), _727_v26, _727_v26, _727_v26)
				var _rhs124 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("sedfiwcj")
				_ = _rhs124
				var _rhs125 bool = !(_766_v59).Equals(_766_v59)
				_ = _rhs125
				_765_v58 = _rhs124
				_727_v26 = _rhs125
				var _767_v60 _dafny.Set
				_ = _767_v60
				_767_v60 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(!(_727_v26))).Cardinality()))
				_727_v26 = !(((_767_v60).Union(_dafny.SetOf(_dafny.IntOfInt64(50), _764_v57.F13))).IsSubsetOf(_767_v60))
				_727_v26 = _727_v26
				var _768_v61 D0
				_ = _768_v61
				_768_v61 = Companion_D0_.Create_DC0_(true)
				var _769_v62 _dafny.Sequence
				_ = _769_v62
				_769_v62 = _dafny.SeqOf(!(p0) || ((_768_v61).Dtor_cf0()), !((_this).F11()))
				_769_v62 = _769_v62
			} else {
				var _770_v63 *C2
				_ = _770_v63
				var _nw132 *C2 = New_C2_()
				_ = _nw132
				_nw132.Ctor__(_764_v57.F13)
				_770_v63 = _nw132
				var _771_v64 bool
				_ = _771_v64
				var _772_v65 _dafny.Map
				_ = _772_v65
				var _773_v66 bool
				_ = _773_v66
				var _out30 bool
				_ = _out30
				var _out31 _dafny.Map
				_ = _out31
				var _out32 bool
				_ = _out32
				_out30, _out31, _out32 = (_this).M2(_727_v26, _727_v26, globalState)
				_771_v64 = _out30
				_772_v65 = _out31
				_773_v66 = _out32
				var _774_v67 _dafny.Map
				_ = _774_v67
				_774_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_773_v66, p0)
				var _775_v68 _dafny.Sequence
				_ = _775_v68
				_775_v68 = _dafny.UnicodeSeqOfUtf8Bytes("yfahqywud")
				var _776_v70 _dafny.Sequence
				_ = _776_v70
				_776_v70 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(80))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}(func(_777_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				})))
				var _778_v71 _dafny.MultiSet
				_ = _778_v71
				_778_v71 = _dafny.MultiSetOf(_764_v57.F13, (_774_v67).Cardinality(), _dafny.IntOfUint32((_775_v68).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
					var _coll16 = _dafny.NewMapBuilder()
					_ = _coll16
					for _iter16 := _dafny.Iterate((_776_v70).Elements()); ; {
						_compr_16, _ok16 := _iter16()
						if !_ok16 {
							break
						}
						var _779_v69 _dafny.Sequence
						_779_v69 = interface{}(_compr_16).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_776_v70, _779_v69) {
							_coll16.Add(_779_v69, (_this).F11())
						}
					}
					return _coll16.ToMap()
				}()).Cardinality())).Cardinality()))
				var _780_v72 *C1
				_ = _780_v72
				var _nw133 *C1 = New_C1_()
				_ = _nw133
				_nw133.Ctor__((_778_v71).IsSubsetOf(Companion_Default___.Fm11(Companion_Default___.Fm0(_727_v26, (_this).F11(), p0, (_this).F11(), globalState), _761_v56, globalState)))
				_780_v72 = _nw133
				_780_v72 = _780_v72
				var _781_v73 *C1
				_ = _781_v73
				var _nw134 *C1 = New_C1_()
				_ = _nw134
				_nw134.Ctor__(p0)
				_781_v73 = _nw134
				(_780_v72).F14 = _771_v64
			}
			(globalState).F2 = Companion_Default___.SafeModuloInt(_764_v57.F13, (_this).F12())
			_727_v26 = _727_v26
			var _782_v74 *C1
			_ = _782_v74
			var _nw135 *C1 = New_C1_()
			_ = _nw135
			_nw135.Ctor__((_764_v57.F13).Cmp(_dafny.IntOfInt64(184)) == 0)
			_782_v74 = _nw135
		} else {
			var _783_v75 _dafny.MultiSet
			_ = _783_v75
			_783_v75 = _dafny.MultiSetOf(_761_v56, _761_v56)
			_727_v26 = (_783_v75).IsDisjointFrom((_783_v75).Union(_783_v75))
			var _784_v76 *C2
			_ = _784_v76
			var _nw136 *C2 = New_C2_()
			_ = _nw136
			_nw136.Ctor__((_this).F12())
			_784_v76 = _nw136
			var _785_v77 _dafny.Sequence
			_ = _785_v77
			_785_v77 = _dafny.SeqOf((_this).F11(), (_this).F11())
			var _786_v78 _dafny.Map
			_ = _786_v78
			_786_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_785_v77).Cardinality()), (_this).F11())
			var _787_v79 _dafny.MultiSet
			_ = _787_v79
			_787_v79 = _dafny.MultiSetOf((_this).F11(), _727_v26, (_this).F11())
			var _788_v80 _dafny.Sequence
			_ = _788_v80
			_788_v80 = _dafny.UnicodeSeqOfUtf8Bytes("onxfcxpf")
			var _789_v81 _dafny.Array
			_ = _789_v81
			var _nw137 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw137
			_789_v81 = _nw137
			var _790_v82 _dafny.Array
			_ = _790_v82
			var _nwElement0_21 _dafny.Int = (_this).F12()
			_ = _nwElement0_21
			var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(15))
			_ = _nw138
			(_nw138).ArraySet1(_nwElement0_21, 0)
			(_nw138).ArraySet1((_this).F12(), 1)
			(_nw138).ArraySet1((_787_v79).Cardinality(), 2)
			(_nw138).ArraySet1((_dafny.IntOfInt64(584)).Minus(_784_v76.F13), 3)
			(_nw138).ArraySet1((_dafny.Zero).Minus((_784_v76.F13).Plus((_this).F12())), 4)
			(_nw138).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_788_v80).Cardinality()), _784_v76.F13), 5)
			(_nw138).ArraySet1((_this).F12(), 6)
			(_nw138).ArraySet1(((_this).F12()).Times(_784_v76.F13), 7)
			(_nw138).ArraySet1((_this).F12(), 8)
			(_nw138).ArraySet1((func() _dafny.Int {
				if (_this).F11() {
					return (_this).F12()
				}
				return _784_v76.F13
			})(), 9)
			(_nw138).ArraySet1((_this).Fm1(_dafny.UnicodeSeqOfUtf8Bytes("lnlmx"), Companion_Default___.Fm2((_this).F11(), globalState), globalState), 10)
			(_nw138).ArraySet1((_this).F12(), 11)
			(_nw138).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_789_v81, true)).Cardinality(), 12)
			(_nw138).ArraySet1((_this).F12(), 13)
			(_nw138).ArraySet1((_this).F12(), 14)
			_790_v82 = _nw138
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_790_v82), 0))
			_ = _index119
			(_790_v82).ArraySet1(_784_v76.F13, (_index119).Int())
			var _791_v83 *C0
			_ = _791_v83
			var _nw139 *C0 = New_C0_()
			_ = _nw139
			_nw139.Ctor__()
			_791_v83 = _nw139
			var _792_v84 D7
			_ = _792_v84
			_792_v84 = Companion_D7_.Create_DC16_(_791_v83)
			var _793_v85 _dafny.Array
			_ = _793_v85
			var _nwElement0_22 *C0 = _791_v83
			_ = _nwElement0_22
			var _nw140 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(27))
			_ = _nw140
			(_nw140).ArraySet1(_nwElement0_22, 0)
			(_nw140).ArraySet1(_791_v83, 1)
			(_nw140).ArraySet1(_791_v83, 2)
			(_nw140).ArraySet1(_791_v83, 3)
			(_nw140).ArraySet1(_791_v83, 4)
			(_nw140).ArraySet1(_791_v83, 5)
			(_nw140).ArraySet1(_791_v83, 6)
			(_nw140).ArraySet1(_791_v83, 7)
			(_nw140).ArraySet1(_791_v83, 8)
			(_nw140).ArraySet1(_791_v83, 9)
			(_nw140).ArraySet1(_791_v83, 10)
			(_nw140).ArraySet1(_791_v83, 11)
			(_nw140).ArraySet1(_791_v83, 12)
			(_nw140).ArraySet1(_791_v83, 13)
			(_nw140).ArraySet1(_791_v83, 14)
			(_nw140).ArraySet1(_791_v83, 15)
			(_nw140).ArraySet1(_791_v83, 16)
			(_nw140).ArraySet1(_791_v83, 17)
			(_nw140).ArraySet1(_791_v83, 18)
			(_nw140).ArraySet1(_791_v83, 19)
			(_nw140).ArraySet1(_791_v83, 20)
			(_nw140).ArraySet1(_791_v83, 21)
			(_nw140).ArraySet1(_791_v83, 22)
			(_nw140).ArraySet1(_791_v83, 23)
			(_nw140).ArraySet1((_792_v84).Dtor_cf24(), 24)
			(_nw140).ArraySet1(_791_v83, 25)
			(_nw140).ArraySet1(_791_v83, 26)
			_793_v85 = _nw140
			var _794_v86 _dafny.Sequence
			_ = _794_v86
			_794_v86 = _dafny.SeqOf(_793_v85, _793_v85, _793_v85, _793_v85)
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_790_v82), 0))
			_ = _index120
			var _rhs126 *C2 = (Companion_D8_.Create_DC19_(_784_v76)).Dtor_cf27()
			_ = _rhs126
			var _rhs127 bool = (_this).F11()
			_ = _rhs127
			var _rhs128 _dafny.Map = (_786_v78).Merge(((Companion_D3_.Create_DC10_(_727_v26, (_this).F12(), p0, _789_v81, _786_v78)).Dtor_cf18()).Merge(_786_v78))
			_ = _rhs128
			var _rhs129 _dafny.Int = (_this).F12()
			_ = _rhs129
			var _rhs130 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_794_v86, _794_v86)
			_ = _rhs130
			var _lhs87 _dafny.Array = _790_v82
			_ = _lhs87
			var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_790_v82), 0))
			_ = _lhs88
			_784_v76 = _rhs126
			_727_v26 = _rhs127
			_786_v78 = _rhs128
			(_lhs87).ArraySet1(_rhs129, (_lhs88).Int())
			_794_v86 = _rhs130
			var _795_v87 D4
			_ = _795_v87
			_795_v87 = Companion_D4_.Create_DC13_(_dafny.CodePoint('f'))
			var _796_v88 _dafny.Map
			_ = _796_v88
			_796_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _795_v87)
			var _797_v89 _dafny.Sequence
			_ = _797_v89
			_797_v89 = _dafny.SeqOf(_784_v76.F13, (_796_v88).Cardinality())
			var _798_v90 _dafny.Sequence
			_ = _798_v90
			_798_v90 = _dafny.SeqOf(_dafny.IntOfUint32((_797_v89).Cardinality()), _784_v76.F13)
			var _799_v91 _dafny.Sequence
			_ = _799_v91
			_799_v91 = _dafny.SeqOf(_dafny.IntOfUint32((_798_v90).Cardinality()))
			(_784_v76).M3(_789_v81, _788_v80, _dafny.Companion_Sequence_.Concatenate(_797_v89, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(911))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}(func(_800_i6 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-298)
			}))), globalState)
			if _727_v26 {
				(_784_v76).F13 = _784_v76.F13
				var _rhs131 _dafny.Int = (_784_v76.F13).Times(((_this).F12()).Plus(_dafny.IntOfUint32((_799_v91).Cardinality())))
				_ = _rhs131
				var _rhs132 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ebyufrei"), _788_v80)
				_ = _rhs132
				var _lhs89 *C2 = _784_v76
				_ = _lhs89
				_lhs89.F13 = _rhs131
				_788_v80 = _rhs132
				(globalState).F2 = (_790_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_790_v82), 0))).Int()).(_dafny.Int)
				_788_v80 = _dafny.Companion_Sequence_.Update(_788_v80, (Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_788_v80).Cardinality()))).Uint32(), _761_v56)
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_789_v81), 0))
				_ = _index121
				(_789_v81).ArraySet1(p0, (_index121).Int())
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_789_v81), 0))
				_ = _index122
				(_789_v81).ArraySet1(true, (_index122).Int())
			} else {
				var _801_v92 _dafny.MultiSet
				_ = _801_v92
				_801_v92 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("wtrbr"))
				var _802_v93 D1
				_ = _802_v93
				_802_v93 = Companion_D1_.Create_DC4_((_801_v92).Difference(_801_v92), (_dafny.Zero).Minus((_790_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_790_v82), 0))).Int()).(_dafny.Int)))
				_802_v93 = _802_v93
				(globalState).F2 = ((_790_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_790_v82), 0))).Int()).(_dafny.Int)).Plus((_this).F12())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_789_v81), 0))
				_ = _index123
				(_789_v81).ArraySet1(p0, (_index123).Int())
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_789_v81), 0))
				_ = _index124
				(_789_v81).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("brsji"), _788_v80), (_index124).Int())
				(globalState).F2 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_784_v76.F13, (func() _dafny.Map {
					var _coll17 = _dafny.NewMapBuilder()
					_ = _coll17
					for _iter17 := _dafny.Iterate((_786_v78).Keys().Elements()); ; {
						_compr_17, _ok17 := _iter17()
						if !_ok17 {
							break
						}
						var _803_v94 _dafny.Int
						_803_v94 = interface{}(_compr_17).(_dafny.Int)
						if (_786_v78).Contains(_803_v94) {
							_coll17.Add(Companion_Default___.SafeDivisionInt(_803_v94, (_this).F12()), _761_v56)
						}
					}
					return _coll17.ToMap()
				}()).Cardinality())))
				var _804_v95 _dafny.Map
				_ = _804_v95
				_804_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if _727_v26 {
						return _784_v76.F13
					}
					return (_790_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_790_v82), 0))).Int()).(_dafny.Int)
				})(), _788_v80)
				_804_v95 = _804_v95
			}
			var _805_v96 _dafny.Set
			_ = _805_v96
			_805_v96 = _dafny.SetOf(!(p0), (_this).F11())
			_727_v26 = (_805_v96).Contains(!((_dafny.SetOf((_797_v89).Select((Companion_Default___.SafeIndex(_784_v76.F13, _dafny.IntOfUint32((_797_v89).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.MultiSetOf(true)).Cardinality(), (_this).F12())).IsDisjointFrom(_dafny.SetOf((_790_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_790_v82), 0))).Int()).(_dafny.Int)))))
		}
		var _806_v97 _dafny.Sequence
		_ = _806_v97
		_806_v97 = _dafny.UnicodeSeqOfUtf8Bytes("ilyuyef")
		var _807_v98 _dafny.MultiSet
		_ = _807_v98
		_807_v98 = _dafny.MultiSetOf(_727_v26)
		var _808_v99 _dafny.Map
		_ = _808_v99
		_808_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_806_v97, (_807_v98).Cardinality())
		var _809_v100 _dafny.Array
		_ = _809_v100
		var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
		_ = _nw141
		_809_v100 = _nw141
		var _810_v101 _dafny.Map
		_ = _810_v101
		_810_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_808_v99, _809_v100)
		var _811_v102 _dafny.MultiSet
		_ = _811_v102
		_811_v102 = _dafny.MultiSetOf(_dafny.IntOfInt64(192))
		var _812_v103 _dafny.MultiSet
		_ = _812_v103
		_812_v103 = _dafny.MultiSetOf(_811_v102, _811_v102, _dafny.MultiSetOf((_this).F12()), (_dafny.MultiSetOf((_this).F12())).Update((_this).F12(), Companion_Default___.Abs((_this).F12())), _811_v102)
		_810_v101 = (_810_v101).Update(((_808_v99).Update(_dafny.UnicodeSeqOfUtf8Bytes("ufaahfrp"), (_this).F12())).Update(_dafny.UnicodeSeqOfUtf8Bytes("yaws"), (func() _dafny.Int {
			if (_812_v103).Contains(_811_v102) {
				return (_812_v103).Multiplicity(_811_v102)
			}
			return (_807_v98).Cardinality()
		})()), _809_v100)
	}
}
func (_this *C3) M2(p0 bool, p1 bool, globalState *GlobalState) (bool, _dafny.Map, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var r2 bool = false
		_ = r2
		var _813_v0 _dafny.Sequence
		_ = _813_v0
		_813_v0 = _dafny.UnicodeSeqOfUtf8Bytes("iteghrvta")
		var _814_v1 _dafny.Map
		_ = _814_v1
		_814_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_813_v0, (_this).F12())
		_814_v1 = (_814_v1).Update(_813_v0, (_this).F12())
		var _815_v2 _dafny.Array
		_ = _815_v2
		var _nw142 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
		_ = _nw142
		_815_v2 = _nw142
		var _816_v3 _dafny.Array
		_ = _816_v3
		var _nw143 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
		_ = _nw143
		_816_v3 = _nw143
		var _817_v4 _dafny.CodePoint
		_ = _817_v4
		_817_v4 = _dafny.CodePoint('w')
		var _818_v5 _dafny.MultiSet
		_ = _818_v5
		_818_v5 = _dafny.MultiSetOf(_817_v4)
		var _819_v6 _dafny.Sequence
		_ = _819_v6
		_819_v6 = _dafny.SeqOf(Companion_Default___.Fm2(p0, globalState))
		Companion_Default___.M0(_815_v2, (_this).F12(), _816_v3, (func() _dafny.Int {
			if (_818_v5).Contains(_817_v4) {
				return (_818_v5).Multiplicity(_817_v4)
			}
			return _dafny.IntOfUint32((_819_v6).Cardinality())
		})(), globalState)
		var _820_v7 *C0
		_ = _820_v7
		var _nw144 *C0 = New_C0_()
		_ = _nw144
		_nw144.Ctor__()
		_820_v7 = _nw144
		var _821_v8 _dafny.Array
		_ = _821_v8
		var _len0_20 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_20
		var _nw145 _dafny.Array
		_ = _nw145
		if _len0_20.Cmp(_dafny.Zero) == 0 {
			_nw145 = _dafny.NewArray(_len0_20)
		} else {
			var _init20 func(_dafny.Int) _dafny.Int = func(_822_i0 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_822_i0, (_this).F12())
			}
			_ = _init20
			var _element0_20 = _init20(_dafny.Zero)
			_ = _element0_20
			_nw145 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
			(_nw145).ArraySet1(_element0_20, 0)
			var _nativeLen0_20 = (_len0_20).Int()
			_ = _nativeLen0_20
			for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
				(_nw145).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
			}
		}
		_821_v8 = _nw145
		(globalState).F10 = _821_v8
		if p0 {
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_815_v2), 0))
			_ = _index125
			(_815_v2).ArraySet1((p0) == ((_this).F11()), (_index125).Int())
			var _823_v9 _dafny.Map
			_ = _823_v9
			_823_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _824_v10 _dafny.Set
			_ = _824_v10
			_824_v10 = _dafny.SetOf((func() bool {
				if (_823_v9).Contains(p1) {
					return (_823_v9).Get(p1).(bool)
				}
				return p1
			})(), (_this).F11())
			var _825_v11 _dafny.Map
			_ = _825_v11
			_825_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_824_v10, p0)
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_815_v2), 0))
			_ = _index126
			(_815_v2).ArraySet1((func() bool {
				if (_825_v11).Contains(_824_v10) {
					return (_825_v11).Get(_824_v10).(bool)
				}
				return (_this).F11()
			})(), (_index126).Int())
			if (_this).F11() {
				var _826_v12 _dafny.MultiSet
				_ = _826_v12
				_826_v12 = _dafny.MultiSetOf(p1)
				r2 = !((_826_v12).IsProperSubsetOf(_826_v12)) || (p0)
				var _827_v13 _dafny.MultiSet
				_ = _827_v13
				_827_v13 = _dafny.MultiSetOf(_824_v10)
				var _828_v14 _dafny.Sequence
				_ = _828_v14
				_828_v14 = _dafny.SeqOf(!(!((_827_v13).IsDisjointFrom(_827_v13))))
				_828_v14 = _828_v14
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_815_v2), 0))
				_ = _index127
				var _rhs133 _dafny.CodePoint = Companion_Default___.Fm6((_this).F12(), p0, Companion_Default___.Fm0((_this).F11(), true, p1, (_this).F11(), globalState), _817_v4, globalState)
				_ = _rhs133
				var _rhs134 bool = !((_815_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_815_v2), 0))).Int()).(bool))
				_ = _rhs134
				var _lhs90 _dafny.Array = _815_v2
				_ = _lhs90
				var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_815_v2), 0))
				_ = _lhs91
				_817_v4 = _rhs133
				(_lhs90).ArraySet1(_rhs134, (_lhs91).Int())
				var _829_v15 D1
				_ = _829_v15
				_829_v15 = Companion_D1_.Create_DC5_(p0, p0)
				r2 = (_820_v7).Fm7((func() bool {
					if p1 {
						return (_829_v15).Dtor_cf9()
					}
					return (_this).F11()
				})(), Companion_D1_.Create_DC3_(_817_v4), _813_v0, (_815_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_815_v2), 0))).Int()).(bool), globalState)
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_821_v8), 0))
				_ = _index128
				(_821_v8).ArraySet1((_this).F12(), (_index128).Int())
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_821_v8), 0))
				_ = _index129
				(_821_v8).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_this).F12(), (_this).F12()), (_this).F12()), (_index129).Int())
			} else {
				var _830_v16 _dafny.Array
				_ = _830_v16
				var _nw146 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(10))
				_ = _nw146
				_830_v16 = _nw146
				var _831_v17 *C2
				_ = _831_v17
				var _nw147 *C2 = New_C2_()
				_ = _nw147
				_nw147.Ctor__((_this).F12())
				_831_v17 = _nw147
				var _832_v18 D8
				_ = _832_v18
				_832_v18 = Companion_D8_.Create_DC19_(_831_v17)
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_830_v16), 0))
				_ = _index130
				(_830_v16).ArraySet1((func() D8 {
					if p1 {
						return _832_v18
					}
					return Companion_D8_.Create_DC19_(_831_v17)
				})(), (_index130).Int())
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_830_v16), 0))
				_ = _index131
				(_830_v16).ArraySet1(_832_v18, (_index131).Int())
				var _833_v19 _dafny.Map
				_ = _833_v19
				_833_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_831_v17.F13, _dafny.IntOfUint32((_813_v0).Cardinality()))
				var _834_v20 _dafny.Map
				_ = _834_v20
				_834_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_815_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_815_v2), 0))).Int()).(bool), _833_v19)
				_834_v20 = (_834_v20).Update(false, _833_v19)
				var _835_v21 D7
				_ = _835_v21
				_835_v21 = Companion_D7_.Create_DC17_((_815_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_815_v2), 0))).Int()).(bool))
				var _836_v22 D7
				_ = _836_v22
				_836_v22 = Companion_D7_.Create_DC18_(_835_v21)
				var _837_v23 D7
				_ = _837_v23
				_837_v23 = Companion_D7_.Create_DC18_(_835_v21)
				var _838_v24 _dafny.Map
				_ = _838_v24
				_838_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_837_v23, (_815_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_815_v2), 0))).Int()).(bool))
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_821_v8), 0))
				_ = _index132
				(_821_v8).ArraySet1((_819_v6).Select((Companion_Default___.SafeIndex((_838_v24).Cardinality(), _dafny.IntOfUint32((_819_v6).Cardinality()))).Uint32()).(_dafny.Int), (_index132).Int())
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_821_v8), 0))
				_ = _index133
				(_821_v8).ArraySet1(_831_v17.F13, (_index133).Int())
				var _839_v25 _dafny.Array
				_ = _839_v25
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_21
				var _nw148 _dafny.Array
				_ = _nw148
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw148 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Int = (func(_840_v8 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_841_i1 _dafny.Int) _dafny.Int {
							return (_841_i1).Times((_840_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_840_v8), 0))).Int()).(_dafny.Int))
						}
					})(_821_v8)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw148 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw148).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw148).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_839_v25 = _nw148
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_815_v2), 0))
				_ = _index134
				var _rhs135 _dafny.Array = _839_v25
				_ = _rhs135
				var _rhs136 _dafny.Int = _831_v17.F13
				_ = _rhs136
				var _rhs137 bool = p0
				_ = _rhs137
				var _rhs138 _dafny.Int = (_821_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_821_v8), 0))).Int()).(_dafny.Int)
				_ = _rhs138
				var _rhs139 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm12(_831_v17.F13, globalState), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_819_v6, _819_v6), (Companion_Default___.SafeIndex((_this).F12(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_819_v6, _819_v6)).Cardinality()))).Uint32(), _dafny.IntOfInt64(406)))
				_ = _rhs139
				var _lhs92 *GlobalState = globalState
				_ = _lhs92
				var _lhs93 *GlobalState = globalState
				_ = _lhs93
				var _lhs94 _dafny.Array = _815_v2
				_ = _lhs94
				var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_815_v2), 0))
				_ = _lhs95
				var _lhs96 *GlobalState = globalState
				_ = _lhs96
				var _lhs97 *GlobalState = globalState
				_ = _lhs97
				_lhs92.F10 = _rhs135
				_lhs93.F2 = _rhs136
				(_lhs94).ArraySet1(_rhs137, (_lhs95).Int())
				_lhs96.F2 = _rhs138
				_lhs97.F7 = _rhs139
				var _842_v26 _dafny.Array
				_ = _842_v26
				var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
				_ = _nw149
				_842_v26 = _nw149
				var _843_v27 _dafny.Array
				_ = _843_v27
				var _nw150 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw150
				_843_v27 = _nw150
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_842_v26), 0))
				_ = _index135
				(_842_v26).ArraySet1(_843_v27, (_index135).Int())
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_842_v26), 0))
				_ = _index136
				(_842_v26).ArraySet1(_843_v27, (_index136).Int())
			}
			(globalState).F10 = _821_v8
			_823_v9 = (_823_v9).Update((_this).F11(), (_815_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_815_v2), 0))).Int()).(bool))
			var _844_v28 *C2
			_ = _844_v28
			var _nw151 *C2 = New_C2_()
			_ = _nw151
			_nw151.Ctor__(((_this).F12()).Times((_this).F12()))
			_844_v28 = _nw151
		} else {
			var _845_v29 _dafny.Map
			_ = _845_v29
			_845_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _821_v8)
			_845_v29 = _845_v29
			(globalState).F7 = _819_v6
			var _846_v30 *C2
			_ = _846_v30
			var _nw152 *C2 = New_C2_()
			_ = _nw152
			_nw152.Ctor__(_dafny.IntOfUint32((_813_v0).Cardinality()))
			_846_v30 = _nw152
			var _847_v31 _dafny.Set
			_ = _847_v31
			_847_v31 = _dafny.SetOf(_846_v30)
			(globalState).F2 = ((_dafny.SetOf(_846_v30)).Difference(_847_v31)).Cardinality()
			var _848_v32 _dafny.Sequence
			_ = _848_v32
			_848_v32 = _813_v0
			var _849_v33 _dafny.Array
			_ = _849_v33
			var _nwElement0_23 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_848_v32), _813_v0)
			_ = _nwElement0_23
			var _nw153 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(13))
			_ = _nw153
			(_nw153).ArraySet1(_nwElement0_23, 0)
			(_nw153).ArraySet1(_813_v0, 1)
			(_nw153).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_813_v0, _813_v0), 2)
			(_nw153).ArraySet1(_813_v0, 3)
			(_nw153).ArraySet1(_813_v0, 4)
			(_nw153).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kctfg"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(908))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}((func(_850_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_851_i2 _dafny.Int) _dafny.CodePoint {
					return _850_v4
				}
			})(_817_v4)))), 5)
			(_nw153).ArraySet1(_813_v0, 6)
			(_nw153).ArraySet1(_813_v0, 7)
			(_nw153).ArraySet1(_813_v0, 8)
			(_nw153).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_813_v0, _813_v0), 9)
			(_nw153).ArraySet1(_813_v0, 10)
			(_nw153).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_813_v0, Companion_Default___.Fm9(p1, globalState)), 11)
			(_nw153).ArraySet1(_813_v0, 12)
			_849_v33 = _nw153
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_849_v33), 0))
			_ = _index137
			(_849_v33).ArraySet1(_813_v0, (_index137).Int())
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_849_v33), 0))
			_ = _index138
			(_849_v33).ArraySet1(_813_v0, (_index138).Int())
			var _852_v34 _dafny.Set
			_ = _852_v34
			_852_v34 = _dafny.SetOf((_this).F12())
			var _853_v35 _dafny.MultiSet
			_ = _853_v35
			_853_v35 = _dafny.MultiSetOf((_this).F11())
			var _854_v36 _dafny.Map
			_ = _854_v36
			_854_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), (_853_v35).Cardinality())
			_852_v34 = (func() _dafny.Set {
				var _coll18 = _dafny.NewBuilder()
				_ = _coll18
				for _iter18 := _dafny.Iterate((_854_v36).Keys().Elements()); ; {
					_compr_18, _ok18 := _iter18()
					if !_ok18 {
						break
					}
					var _855_v37 _dafny.Int
					_855_v37 = interface{}(_compr_18).(_dafny.Int)
					if (_854_v36).Contains(_855_v37) {
						_coll18.Add(Companion_Default___.SafeModuloInt(_855_v37, (_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(365)))).Cardinality())))
					}
				}
				return _coll18.ToSet()
			}()).Difference(_852_v34)
		}
		(globalState).F2 = (_this).F12()
		r0 = (_this).F11()
		var _856_v38 _dafny.Sequence
		_ = _856_v38
		_856_v38 = _dafny.SeqOf(!((_this).F11()), p0, (_this).F11())
		var _857_v39 _dafny.Map
		_ = _857_v39
		_857_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_856_v38, p0)
		r1 = _857_v39
		r2 = _dafny.Companion_Sequence_.IsPrefixOf(_856_v38, _856_v38)
		return r0, r1, r2
	}
}
func (_this *C3) F11() bool {
	{
		return _this._f11
	}
}
func (_this *C3) F12() _dafny.Int {
	{
		return _this._f12
	}
}

// End of class C3
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
