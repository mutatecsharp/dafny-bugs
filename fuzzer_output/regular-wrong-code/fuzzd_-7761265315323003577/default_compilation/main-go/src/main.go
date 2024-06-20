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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Set, p2 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source0 D1 = Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_dafny.CodePoint('s'), _dafny.CodePoint('l'), _dafny.CodePoint('q')))
	_ = _source0
	if _source0.Is_DC4() {
		var _0___mcc_h0 D0 = _source0.Get_().(D1_DC4).Cf7
		_ = _0___mcc_h0
		var _1___mcc_h1 _dafny.Int = _source0.Get_().(D1_DC4).Cf8
		_ = _1___mcc_h1
		var _2___mcc_h2 _dafny.Sequence = _source0.Get_().(D1_DC4).Cf9
		_ = _2___mcc_h2
		var _3___mcc_h3 bool = _source0.Get_().(D1_DC4).Cf10
		_ = _3___mcc_h3
		var _4___mcc_h4 _dafny.Int = _source0.Get_().(D1_DC4).Cf11
		_ = _4___mcc_h4
		var _5_cf11 _dafny.Int = _4___mcc_h4
		_ = _5_cf11
		var _6_cf10 bool = _3___mcc_h3
		_ = _6_cf10
		var _7_cf9 _dafny.Sequence = _2___mcc_h2
		_ = _7_cf9
		var _8_cf8 _dafny.Int = _1___mcc_h1
		_ = _8_cf8
		var _9_cf7 D0 = _0___mcc_h0
		_ = _9_cf7
		return (_7_cf9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_7_cf9).Cardinality()), _dafny.IntOfUint32((_7_cf9).Cardinality()))).Uint32()).(_dafny.CodePoint)
	} else {
		var _10___mcc_h5 _dafny.MultiSet = _source0.Get_().(D1_DC3).Cf6
		_ = _10___mcc_h5
		var _11_cf6 _dafny.MultiSet = _10___mcc_h5
		_ = _11_cf6
		return _dafny.CodePoint('n')
	}
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.Zero).Minus((func() _dafny.Int {
		if false {
			return (_dafny.MultiSetOf(false, !(!(false)), false)).Cardinality()
		}
		return _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf((func() _dafny.Map {
			var _coll0 = _dafny.NewMapBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(-912))).Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _12_v0 _dafny.Int
				_12_v0 = interface{}(_compr_0).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfInt64(-912))).Contains(_12_v0) {
					_coll0.Add(Companion_Default___.SafeModuloInt(_12_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()), (func() _dafny.Map {
						var _coll1 = _dafny.NewMapBuilder()
						_ = _coll1
						for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(676), _dafny.IntOfInt64(222))); ; {
							_compr_1, _ok1 := _iter1()
							if !_ok1 {
								break
							}
							var _13_v1 _dafny.Int
							_13_v1 = interface{}(_compr_1).(_dafny.Int)
							if ((_dafny.IntOfInt64(676)).Cmp(_13_v1) <= 0) && ((_13_v1).Cmp(_dafny.IntOfInt64(222)) < 0) {
								_coll1.Add((_13_v1).Plus(_dafny.IntOfInt64(-910)), true)
							}
						}
						return _coll1.ToMap()
					}()).Cardinality())
				}
			}
			return _coll0.ToMap()
		}()).Cardinality())).Cardinality())).Cardinality())
	})()), (_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), false)).Cardinality()).Times(_dafny.IntOfInt64(-907))), ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lk")).Cardinality()))).Plus(_dafny.IntOfInt64(802)))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return ((_dafny.IntOfInt64(-235)).Minus(_dafny.IntOfInt64(-236))).Times(_dafny.IntOfInt64(663))
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, false), _dafny.SeqOf(false, false)), _dafny.SeqOf(false))
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf((Companion_D11_.Create_DC33_(_dafny.IntOfInt64(-562), !(false), _dafny.IntOfInt64(369), false, _dafny.SetOf(true))).Dtor_cf65(), false))).Cardinality(), _dafny.IntOfInt64(412))).Cardinality(), _dafny.IntOfInt64(706))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_dafny.SetOf(true, false)).Cardinality())).Cardinality(), _dafny.IntOfInt64(906)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(697))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_14_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	}))).Cardinality()))).Merge(func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(294), _dafny.IntOfInt64(803))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _15_v0 _dafny.Int
			_15_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(294)).Cmp(_15_v0) <= 0) && ((_15_v0).Cmp(_dafny.IntOfInt64(803)) < 0) {
				_coll2.Add((_15_v0).Times((_dafny.MultiSetOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality()))).Cardinality()), _dafny.IntOfInt64(781))
			}
		}
		return _coll2.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC8_(false, !((_dafny.SetOf(_dafny.IntOfInt64(90))).IsProperSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(894), (_dafny.MultiSetOf(false, true, false)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(357))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_16_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	})), _dafny.UnicodeSeqOfUtf8Bytes("fqjluaqn")), _dafny.UnicodeSeqOfUtf8Bytes("pfpcdepv"))
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(902))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_17_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	})), _dafny.UnicodeSeqOfUtf8Bytes("go")), _dafny.UnicodeSeqOfUtf8Bytes("ywfd"))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(400), (_dafny.MultiSetOf(false, false, !((Companion_D0_.Create_DC1_(!(true), (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-171))).Uint32(), func(coer3 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_18_i0 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.SetOf(!(true))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qxcxo")).Cardinality())))
	})))).Cardinality(), _dafny.CodePoint('h'))).Dtor_cf2()), !(false))).Cardinality())).Cardinality(), (_dafny.MultiSetOf(false)).Cardinality(), (_dafny.MultiSetOf(false, true, false, true)).Cardinality())).Difference(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(538))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_19_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-397)
	}))))
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, p1 bool, p2 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_((_dafny.IntOfInt64(-119)).Cmp(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SeqOf(func() _dafny.Set {
			var _coll4 = _dafny.NewBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-648), _dafny.IntOfInt64(965))); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _20_v1 _dafny.Int
				_20_v1 = interface{}(_compr_4).(_dafny.Int)
				if ((_dafny.IntOfInt64(-648)).Cmp(_20_v1) <= 0) && ((_20_v1).Cmp(_dafny.IntOfInt64(965)) < 0) {
					_coll4.Add(Companion_Default___.SafeDivisionInt(_20_v1, _dafny.IntOfInt64(422)))
				}
			}
			return _coll4.ToSet()
		}())).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _21_v0 _dafny.Set
			_21_v0 = interface{}(_compr_3).(_dafny.Set)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(func() _dafny.Set {
				var _coll5 = _dafny.NewBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-648), _dafny.IntOfInt64(965))); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _22_v1 _dafny.Int
					_22_v1 = interface{}(_compr_5).(_dafny.Int)
					if ((_dafny.IntOfInt64(-648)).Cmp(_22_v1) <= 0) && ((_22_v1).Cmp(_dafny.IntOfInt64(965)) < 0) {
						_coll5.Add(Companion_Default___.SafeDivisionInt(_22_v1, _dafny.IntOfInt64(422)))
					}
				}
				return _coll5.ToSet()
			}()), _21_v0) {
				_coll3.Add(_21_v0, true)
			}
		}
		return _coll3.ToMap()
	}()).Cardinality()))).Cardinality())) >= 0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qnlh"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(886))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_23_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	})))).Cardinality()), _dafny.CodePoint('r'))
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(761))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_24_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-716), (_dafny.SetOf(true)).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(993))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_25_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)).Cardinality()))
	}))), (func() _dafny.Sequence {
		if true {
			return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(495))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_26_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			}))).Cardinality()))
		}
		return _dafny.SeqOf(_dafny.IntOfInt64(90), _dafny.IntOfInt64(979))
	})())
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("tqkyhgesq"), _dafny.IntOfInt64(-595))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("bk"), _dafny.IntOfInt64(753)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mrd"), _dafny.IntOfInt64(673)))
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Set, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("yhclj")), (Companion_D21_.Create_DC49_(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("jcnirnr"), _dafny.UnicodeSeqOfUtf8Bytes("ttxfurlkd")))).Dtor_cf88())
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("xty")
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (Companion_D9_.Create_DC25_(_dafny.IntOfInt64(-77), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tussinc")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tk")).Cardinality()))).Cardinality())), _dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()), _dafny.CodePoint('n'))).Dtor_cf47())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(10))).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-344)))
}
func (_static *CompanionStruct_Default___) Fm32(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) D0 {
	if !((false) || (false)) {
		return Companion_D0_.Create_DC0_(_dafny.SeqOf(!(false), false), _dafny.IntOfInt64(468))
	} else {
		return Companion_D0_.Create_DC0_(_dafny.SeqOf(true), _dafny.IntOfInt64(501))
	}
}
func (_static *CompanionStruct_Default___) Fm33(p0 D3, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return (((Companion_D22_.Create_DC51_(func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(593), _dafny.IntOfInt64(-876))).Keys().Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _27_v0 _dafny.Int
			_27_v0 = interface{}(_compr_6).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(593), _dafny.IntOfInt64(-876))).Contains(_27_v0) {
				_coll6.Add((_27_v0).Minus(_dafny.IntOfInt64(505)), true)
			}
		}
		return _coll6.ToMap()
	}())).Dtor_cf93()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(482), !(!(true))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(381))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_28_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(true)).Cardinality())
	}))).Cardinality())), true))
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true)).Intersection((_dafny.SetOf(!(false))).Union(_dafny.SetOf(false, !(false), false, false, true)))
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.IntOfInt64(778), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(584), _dafny.SeqOf(true))).Cardinality()).Minus(_dafny.IntOfInt64(-262)), _dafny.IntOfInt64(812), _dafny.IntOfInt64(881), (_dafny.IntOfInt64(222)).Times(_dafny.IntOfInt64(-684)))
}
func (_static *CompanionStruct_Default___) Fm36(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ekwnkte")).Cardinality()), _dafny.IntOfInt64(380), _dafny.IntOfInt64(628), _dafny.IntOfInt64(-978), _dafny.IntOfInt64(812)), _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-503))), _dafny.MultiSetOf((Companion_D0_.Create_DC1_(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xuckykvh")).Cardinality()), _dafny.CodePoint('f'))).Dtor_cf3(), _dafny.IntOfInt64(815)))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(377), _dafny.IntOfInt64(315)), _dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-684))).Cardinality())), _dafny.IntOfInt64(826)), _dafny.MultiSetOf(_dafny.IntOfInt64(150), _dafny.IntOfInt64(489)))))).Difference((_dafny.MultiSetOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(326), !(false))).Cardinality()))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(868))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_29_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("ijuq")
	}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.MultiSetOf(true))).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC19_(_dafny.IntOfInt64(543), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(691), _dafny.IntOfInt64(460)), (_dafny.IntOfInt64(95)).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-772)))).Cardinality())).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm38(globalState *GlobalState) D5 {
	if (func() bool {
		if !(true) {
			return true
		}
		return !(!(false))
	})() {
		return Companion_D5_.Create_DC16_(_dafny.SetOf(!(false), false, true))
	} else {
		return Companion_D5_.Create_DC16_(_dafny.SetOf(true))
	}
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(767))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_30_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	}))))
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.SetOf(func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(12), _dafny.IntOfInt64(391))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _31_v0 _dafny.Int
			_31_v0 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(12)).Cmp(_31_v0) <= 0) && ((_31_v0).Cmp(_dafny.IntOfInt64(391)) < 0) {
				_coll7.Add((_31_v0).Times((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality()), false)
			}
		}
		return _coll7.ToMap()
	}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(670), false))).Cardinality()), _dafny.IntOfInt64(-354)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), !(!(false)))).Cardinality())).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm41(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)))
}
func (_static *CompanionStruct_Default___) Fm42(p0 bool, p1 _dafny.MultiSet, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("emmw")).Cardinality()), (_dafny.IntOfInt64(749)).Plus(_dafny.IntOfInt64(-156)))
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-894))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_32_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(344))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_33_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-599)
		}))
	})), _dafny.SeqOf(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-910))), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, !(true))).Cardinality())))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(792))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_34_i2 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(28))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_35_i3 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf(_dafny.IntOfInt64(375), _dafny.IntOfInt64(729))
		}))).Cardinality()))
	})))
}
func (_static *CompanionStruct_Default___) Fm44(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false))).Cardinality(), false)).Cardinality()), (_dafny.SetOf(_dafny.IntOfInt64(499), _dafny.IntOfInt64(373), (_dafny.SetOf(true)).Cardinality())).Cardinality()), Companion_D9_.Create_DC25_(_dafny.IntOfInt64(788), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(311))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_36_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))).Cardinality()), !(false))).Cardinality())).Cardinality()), (_dafny.SetOf(false, false)).Cardinality(), _dafny.CodePoint('u')))
}
func (_static *CompanionStruct_Default___) Fm45(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D4 {
	var _source1 D4 = Companion_D4_.Create_DC15_(true, !(false), _dafny.UnicodeSeqOfUtf8Bytes("vypeoo"))
	_ = _source1
	if _source1.Is_DC13() {
		var _37___mcc_h0 _dafny.Sequence = _source1.Get_().(D4_DC13).Cf25
		_ = _37___mcc_h0
		var _38___mcc_h1 bool = _source1.Get_().(D4_DC13).Cf26
		_ = _38___mcc_h1
		var _39___mcc_h2 _dafny.Int = _source1.Get_().(D4_DC13).Cf27
		_ = _39___mcc_h2
		var _40_cf27 _dafny.Int = _39___mcc_h2
		_ = _40_cf27
		var _41_cf26 bool = _38___mcc_h1
		_ = _41_cf26
		var _42_cf25 _dafny.Sequence = _37___mcc_h0
		_ = _42_cf25
		return Companion_D4_.Create_DC13_(_dafny.SeqOf(_40_cf27), _41_cf26, _40_cf27)
	} else if _source1.Is_DC14() {
		var _43___mcc_h3 bool = _source1.Get_().(D4_DC14).Cf28
		_ = _43___mcc_h3
		var _44___mcc_h4 _dafny.Map = _source1.Get_().(D4_DC14).Cf29
		_ = _44___mcc_h4
		var _45___mcc_h5 _dafny.Int = _source1.Get_().(D4_DC14).Cf30
		_ = _45___mcc_h5
		var _46___mcc_h6 _dafny.Map = _source1.Get_().(D4_DC14).Cf31
		_ = _46___mcc_h6
		var _47_cf31 _dafny.Map = _46___mcc_h6
		_ = _47_cf31
		var _48_cf30 _dafny.Int = _45___mcc_h5
		_ = _48_cf30
		var _49_cf29 _dafny.Map = _44___mcc_h4
		_ = _49_cf29
		var _50_cf28 bool = _43___mcc_h3
		_ = _50_cf28
		return Companion_D4_.Create_DC13_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(43))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_51_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(652)
		})), _50_cf28, _48_cf30)
	} else if _source1.Is_DC15() {
		var _52___mcc_h7 bool = _source1.Get_().(D4_DC15).Cf32
		_ = _52___mcc_h7
		var _53___mcc_h8 bool = _source1.Get_().(D4_DC15).Cf33
		_ = _53___mcc_h8
		var _54___mcc_h9 _dafny.Sequence = _source1.Get_().(D4_DC15).Cf34
		_ = _54___mcc_h9
		var _55_cf34 _dafny.Sequence = _54___mcc_h9
		_ = _55_cf34
		var _56_cf33 bool = _53___mcc_h8
		_ = _56_cf33
		var _57_cf32 bool = _52___mcc_h7
		_ = _57_cf32
		return Companion_D4_.Create_DC13_(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_cf33, _dafny.IntOfInt64(34))).Cardinality(), (Companion_D1_.Create_DC4_(Companion_D0_.Create_DC0_(_dafny.SeqOf(_57_cf32, _57_cf32, _56_cf33, true), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_56_cf33))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_cf33, _dafny.IntOfInt64(-314))).Cardinality(), _55_cf34, _56_cf33, _dafny.IntOfInt64(222))).Dtor_cf11()), _57_cf32, (_dafny.Zero).Minus(_dafny.IntOfInt64(-758)))
	} else {
		var _58___mcc_h10 _dafny.Sequence = _source1.Get_().(D4_DC12).Cf24
		_ = _58___mcc_h10
		var _59_cf24 _dafny.Sequence = _58___mcc_h10
		_ = _59_cf24
		return Companion_D4_.Create_DC13_(_dafny.SeqOf(_dafny.IntOfInt64(368)), false, _dafny.IntOfInt64(-642))
	}
}
func (_static *CompanionStruct_Default___) Fm46(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false))), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))))
}
func (_static *CompanionStruct_Default___) Fm47(p0 _dafny.CodePoint, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(719))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_60_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	})))
}
func (_static *CompanionStruct_Default___) Fm48(globalState *GlobalState) _dafny.MultiSet {
	return ((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetOf(false)
		}
		return _dafny.MultiSetOf(false, true)
	})()).Union((_dafny.MultiSetOf(false)).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(true, false, false))))
}
func (_static *CompanionStruct_Default___) Fm49(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) D2 {
	if (Companion_D4_.Create_DC13_(_dafny.SeqOf(_dafny.IntOfInt64(-20), _dafny.IntOfInt64(680)), !(!(!(false))), _dafny.IntOfInt64(-244))).Dtor_cf26() {
		return Companion_D2_.Create_DC8_(!(true), true)
	} else {
		return Companion_D2_.Create_DC8_(true, true)
	}
}
func (_static *CompanionStruct_Default___) Fm50(p0 _dafny.Sequence, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC11_(Companion_D3_.Create_DC11_(Companion_D3_.Create_DC9_(_dafny.SeqOf(_dafny.IntOfInt64(-215)))))
}
func (_static *CompanionStruct_Default___) Fm51(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(577))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}(func(_61_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-201)
	})), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(491))).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(986)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dyeiakrtw")).Cardinality()), _dafny.IntOfInt64(-582), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(755), _dafny.IntOfInt64(-79), (_dafny.Zero).Minus(_dafny.IntOfInt64(-688)))).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(122))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg20 _dafny.Int) interface{} {
			return coer20(arg20)
		}
	}(func(_62_i1 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(23))).Cardinality()
	}))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(822))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}(func(_63_i2 _dafny.Int) _dafny.Int {
		return (Companion_D11_.Create_DC31_(((Companion_D4_.Create_DC14_(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.Zero, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true)))).Dtor_cf29()).Cardinality(), (_dafny.MultiSetOf(false)).Cardinality())).Dtor_cf59()
	})))
}
func (_static *CompanionStruct_Default___) Fm52(p0 bool, globalState *GlobalState) D14 {
	return Companion_D14_.Create_DC37_((_dafny.SetOf(false, true)).IsSubsetOf(_dafny.SetOf(false)), _dafny.UnicodeSeqOfUtf8Bytes("k"))
}
func (_static *CompanionStruct_Default___) Fm53(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_dafny.CodePoint('v'))))).Union(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-950))).Uint32(), func(coer22 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
		return func(arg22 _dafny.Int) interface{} {
			return coer22(arg22)
		}
	}(func(_64_i0 _dafny.Int) D1 {
		return Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_dafny.CodePoint('j')))
	})), _dafny.SeqOf(Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_dafny.CodePoint('i'), _dafny.CodePoint('k'))), Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_dafny.CodePoint('l')))))))
}
func (_static *CompanionStruct_Default___) Fm54(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.CodePoint('f'))
}
func (_static *CompanionStruct_Default___) Fm55(globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC15_((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iptf")).Cardinality())).Cmp(_dafny.IntOfInt64(-985)) <= 0, !((false) && (true)), (Companion_D4_.Create_DC15_(true, false, _dafny.UnicodeSeqOfUtf8Bytes("uj"))).Dtor_cf34())
}
func (_static *CompanionStruct_Default___) Fm56(p0 _dafny.Sequence, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) D11 {
	return Companion_D11_.Create_DC31_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(804), (_dafny.MultiSetFromSeq(_dafny.SeqOf(true, !(true)))).Cardinality()), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(316))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}(func(_65_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	}))).Cardinality()), _dafny.IntOfInt64(827)))).Cardinality()), (_dafny.IntOfInt64(518)).Plus(_dafny.IntOfInt64(449)))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _66_v0 _dafny.CodePoint
	_ = _66_v0
	_66_v0 = _dafny.CodePoint('s')
	var _67_v1 _dafny.Sequence
	_ = _67_v1
	_67_v1 = _dafny.SeqOf(_dafny.CodePoint('f'), _66_v0)
	var _68_v2 _dafny.Int
	_ = _68_v2
	_68_v2 = _dafny.IntOfInt64(291)
	var _69_v3 _dafny.Sequence
	_ = _69_v3
	_69_v3 = _dafny.SeqOf(_68_v2)
	var _70_v4 _dafny.Map
	_ = _70_v4
	_70_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_69_v3).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vyifsmwe")).Cardinality()))
	var _71_globalState *GlobalState
	_ = _71_globalState
	var _nw0 *GlobalState = New_GlobalState_()
	_ = _nw0
	_nw0.Ctor__(true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(215))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg24 _dafny.Int) interface{} {
			return coer24(arg24)
		}
	}(func(_72_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	})), _67_v1, true, _70_v4, _dafny.IntOfInt64(583))
	_71_globalState = _nw0
	var _73_v5 bool
	_ = _73_v5
	_73_v5 = false
	var _74_v6 D0
	_ = _74_v6
	_74_v6 = Companion_D0_.Create_DC1_(_73_v5, Companion_Default___.SafeModuloInt(_68_v2, _68_v2), Companion_Default___.Fm0(_68_v2, Companion_Default___.Fm1(_71_globalState), _73_v5, _71_globalState))
	var _source2 D0 = _74_v6
	_ = _source2
	if _source2.Is_DC0() {
		var _75___mcc_h0 _dafny.Sequence = _source2.Get_().(D0_DC0).Cf0
		_ = _75___mcc_h0
		var _76___mcc_h1 _dafny.Int = _source2.Get_().(D0_DC0).Cf1
		_ = _76___mcc_h1
		var _77_cf1 _dafny.Int = _76___mcc_h1
		_ = _77_cf1
		var _78_cf0 _dafny.Sequence = _75___mcc_h0
		_ = _78_cf0
		var _79_v7 _dafny.Map
		_ = _79_v7
		_79_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_68_v2, _71_globalState), _66_v0)
		_79_v7 = _79_v7
		var _80_v8 _dafny.Array
		_ = _80_v8
		var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
		_ = _nw1
		_80_v8 = _nw1
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))
		_ = _index0
		(_80_v8).ArraySet1(_dafny.IntOfInt64(-792), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))
		_ = _index1
		(_80_v8).ArraySet1((_77_cf1).Minus(Companion_Default___.SafeModuloInt(_77_cf1, _dafny.IntOfUint32((_78_cf0).Cardinality()))), (_index1).Int())
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))
		_ = _index2
		(_80_v8).ArraySet1((_80_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))).Int()).(_dafny.Int), (_index2).Int())
		var _81_v9 D0
		_ = _81_v9
		_81_v9 = Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(_73_v5, _68_v2, _66_v0))
		var _82_v10 _dafny.Sequence
		_ = _82_v10
		_82_v10 = _dafny.SeqOf(_67_v1, _67_v1, _67_v1, _67_v1, _67_v1)
		var _83_v11 D0
		_ = _83_v11
		_83_v11 = Companion_D0_.Create_DC0_(Companion_Default___.Fm4(_68_v2, _dafny.IntOfUint32((_82_v10).Cardinality()), (_80_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))).Int()).(_dafny.Int), _71_globalState), _77_cf1)
		var _84_v12 D0
		_ = _84_v12
		_84_v12 = Companion_D0_.Create_DC2_(_83_v11)
		var _85_v13 D0
		_ = _85_v13
		_85_v13 = Companion_D0_.Create_DC2_(_83_v11)
		var _pat_let_tv0 = _83_v11
		_ = _pat_let_tv0
		var _source3 D0 = (func() D0 {
			if Companion_Default___.Fm3(_71_globalState) {
				return _81_v9
			}
			return func(_pat_let0_0 D0) D0 {
				return func(_86_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let1_0 D0) D0 {
						return func(_87_dt__update_hcf5_h0 D0) D0 {
							return Companion_D0_.Create_DC2_(_87_dt__update_hcf5_h0)
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(_81_v9)
		})()
		_ = _source3
		if _source3.Is_DC0() {
			var _88___mcc_h6 _dafny.Sequence = _source3.Get_().(D0_DC0).Cf0
			_ = _88___mcc_h6
			var _89___mcc_h7 _dafny.Int = _source3.Get_().(D0_DC0).Cf1
			_ = _89___mcc_h7
			var _90_cf1 _dafny.Int = _89___mcc_h7
			_ = _90_cf1
			var _91_cf0 _dafny.Sequence = _88___mcc_h6
			_ = _91_cf0
			var _92_v14 *C1
			_ = _92_v14
			var _nw2 *C1 = New_C1_()
			_ = _nw2
			_nw2.Ctor__(_68_v2)
			_92_v14 = _nw2
			_73_v5 = (_91_cf0).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_90_cf1), _dafny.IntOfUint32((_91_cf0).Cardinality()))).Uint32()).(bool)
			_90_cf1 = (_dafny.Zero).Minus((_80_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))).Int()).(_dafny.Int))
			var _93_v15 T0
			_ = _93_v15
			var _nw3 *C5 = New_C5_()
			_ = _nw3
			_nw3.Ctor__(_90_cf1, _90_cf1)
			_93_v15 = _nw3
			var _94_v16 _dafny.MultiSet
			_ = _94_v16
			_94_v16 = _dafny.MultiSetOf(_93_v15)
			var _95_v17 _dafny.Map
			_ = _95_v17
			_95_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v2, _73_v5)
			var _96_v18 _dafny.Map
			_ = _96_v18
			_96_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_95_v17).Contains(_90_cf1) {
					return (_95_v17).Get(_90_cf1).(bool)
				}
				return !(_73_v5)
			})(), (_80_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))).Int()).(_dafny.Int))
			var _97_v19 _dafny.Array
			_ = _97_v19
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_0
			var _nw4 _dafny.Array
			_ = _nw4
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw4 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) bool = (func(_98_v5 bool) func(_dafny.Int) bool {
					return func(_99_i1 _dafny.Int) bool {
						return _98_v5
					}
				})(_73_v5)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw4 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw4).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw4).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_97_v19 = _nw4
			var _100_v20 _dafny.Map
			_ = _100_v20
			_100_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(547), _67_v1)
			var _101_v21 _dafny.MultiSet
			_ = _101_v21
			_101_v21 = _dafny.MultiSetOf(_68_v2)
			var _102_v22 _dafny.MultiSet
			_ = _102_v22
			_102_v22 = _dafny.MultiSetOf(_101_v21)
			var _103_v23 T1
			_ = _103_v23
			var _nw5 *C4 = New_C4_()
			_ = _nw5
			_nw5.Ctor__((func() _dafny.Int {
				if (_94_v16).Contains(_93_v15) {
					return (_94_v16).Multiplicity(_93_v15)
				}
				return _68_v2
			})(), ((_96_v18).Update(_73_v5, _dafny.IntOfUint32((_67_v1).Cardinality()))).Cardinality(), _97_v19, _100_v20, _68_v2, _102_v22)
			_103_v23 = _nw5
			var _104_v24 D6
			_ = _104_v24
			_104_v24 = Companion_D6_.Create_DC18_(_103_v23)
			var _105_v25 _dafny.Map
			_ = _105_v25
			_105_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v8, _69_v3)
			var _106_v26 _dafny.Map
			_ = _106_v26
			_106_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_104_v24, _105_v25)
			(_71_globalState).F5 = (((func() _dafny.Map {
				if (_106_v26).Contains(_104_v24) {
					return (_106_v26).Get(_104_v24).(_dafny.Map)
				}
				return _105_v25
			})()).Update(_80_v8, _69_v3)).Cardinality()
		} else if _source3.Is_DC1() {
			var _107___mcc_h8 bool = _source3.Get_().(D0_DC1).Cf2
			_ = _107___mcc_h8
			var _108___mcc_h9 _dafny.Int = _source3.Get_().(D0_DC1).Cf3
			_ = _108___mcc_h9
			var _109___mcc_h10 _dafny.CodePoint = _source3.Get_().(D0_DC1).Cf4
			_ = _109___mcc_h10
			var _110_cf4 _dafny.CodePoint = _109___mcc_h10
			_ = _110_cf4
			var _111_cf3 _dafny.Int = _108___mcc_h9
			_ = _111_cf3
			var _112_cf2 bool = _107___mcc_h8
			_ = _112_cf2
			var _113_v27 _dafny.Array
			_ = _113_v27
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
			_ = _nw6
			_113_v27 = _nw6
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
			_ = _nw7
			_113_v27 = _nw7
			var _114_v28 _dafny.Map
			_ = _114_v28
			_114_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_111_cf3), _112_cf2)
			_114_v28 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_80_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))).Int()).(_dafny.Int), !(_112_cf2))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_cf1, _112_cf2))
			var _115_v29 D6
			_ = _115_v29
			_115_v29 = Companion_D6_.Create_DC19_(_68_v2, (_80_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(121))
			var _pat_let_tv1 = _80_v8
			_ = _pat_let_tv1
			var _pat_let_tv2 = _80_v8
			_ = _pat_let_tv2
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))
			_ = _index3
			(_80_v8).ArraySet1(Companion_Default___.Fm2((func(_pat_let2_0 D6) D6 {
				return func(_116_dt__update__tmp_h1 D6) D6 {
					return func(_pat_let3_0 _dafny.Int) D6 {
						return func(_117_dt__update_hcf41_h0 _dafny.Int) D6 {
							return Companion_D6_.Create_DC19_((_116_dt__update__tmp_h1).Dtor_cf39(), (_116_dt__update__tmp_h1).Dtor_cf40(), _117_dt__update_hcf41_h0)
						}(_pat_let3_0)
					}((_pat_let_tv2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_pat_let_tv1), 0))).Int()).(_dafny.Int))
				}(_pat_let2_0)
			}(_115_v29)).Dtor_cf41(), _71_globalState), (_index3).Int())
			var _118_v30 *C6
			_ = _118_v30
			var _nw8 *C6 = New_C6_()
			_ = _nw8
			_nw8.Ctor__(_68_v2, _73_v5)
			_118_v30 = _nw8
			var _119_v31 _dafny.Set
			_ = _119_v31
			_119_v31 = _dafny.SetOf(_118_v30, _118_v30, _118_v30, _118_v30, _118_v30)
			var _120_v32 _dafny.Map
			_ = _120_v32
			_120_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_v31, _78_cf0)
			var _121_v33 _dafny.Sequence
			_ = _121_v33
			_121_v33 = _dafny.SeqOf(_119_v31, _dafny.SetOf(_118_v30, _118_v30, _118_v30))
			_120_v32 = (_120_v32).Update((_121_v33).Select((Companion_Default___.SafeIndex(_68_v2, _dafny.IntOfUint32((_121_v33).Cardinality()))).Uint32()).(_dafny.Set), _78_cf0)
		} else {
			var _122___mcc_h11 D0 = _source3.Get_().(D0_DC2).Cf5
			_ = _122___mcc_h11
			var _123_cf5 D0 = _122___mcc_h11
			_ = _123_cf5
			var _124_v34 _dafny.Array
			_ = _124_v34
			var _nw9 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw9
			_124_v34 = _nw9
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_124_v34), 0))
			_ = _index4
			(_124_v34).ArraySet1(_73_v5, (_index4).Int())
			var _125_v35 _dafny.Map
			_ = _125_v35
			_125_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v5, _73_v5)
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_124_v34), 0))
			_ = _index5
			(_124_v34).ArraySet1((func() bool {
				if (_125_v35).Contains(_73_v5) {
					return (_125_v35).Get(_73_v5).(bool)
				}
				return _73_v5
			})(), (_index5).Int())
			var _126_v36 *C9
			_ = _126_v36
			var _nw10 *C9 = New_C9_()
			_ = _nw10
			_nw10.Ctor__((_80_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))).Int()).(_dafny.Int))
			_126_v36 = _nw10
			var _127_v37 _dafny.MultiSet
			_ = _127_v37
			_127_v37 = _dafny.MultiSetOf((_124_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_124_v34), 0))).Int()).(bool))
			var _128_v38 _dafny.Map
			_ = _128_v38
			_128_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v5, (_dafny.Zero).Minus((_127_v37).Cardinality()))
			_68_v2 = (_68_v2).Times(Companion_Default___.SafeModuloInt(_68_v2, (func() _dafny.Int {
				if (_128_v38).Contains(true) {
					return (_128_v38).Get(true).(_dafny.Int)
				}
				return _dafny.IntOfInt64(858)
			})()))
			var _129_v39 _dafny.Array
			_ = _129_v39
			var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
			_ = _nw11
			_129_v39 = _nw11
			var _130_v40 _dafny.Array
			_ = _130_v40
			var _nwElement0_0 _dafny.Array = _129_v39
			_ = _nwElement0_0
			var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(6))
			_ = _nw12
			(_nw12).ArraySet1(_nwElement0_0, 0)
			(_nw12).ArraySet1(_129_v39, 1)
			(_nw12).ArraySet1(_129_v39, 2)
			(_nw12).ArraySet1(_129_v39, 3)
			(_nw12).ArraySet1((func() _dafny.Array {
				if _73_v5 {
					return _129_v39
				}
				return _129_v39
			})(), 4)
			(_nw12).ArraySet1(_129_v39, 5)
			_130_v40 = _nw12
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_130_v40), 0))
			_ = _index6
			(_130_v40).ArraySet1(_129_v39, (_index6).Int())
			var _131_v41 *C10
			_ = _131_v41
			var _nw13 *C10 = New_C10_()
			_ = _nw13
			_nw13.Ctor__(_77_cf1)
			_131_v41 = _nw13
			var _132_v42 _dafny.Map
			_ = _132_v42
			_132_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v41, _dafny.IntOfInt64(692))
			var _133_v43 _dafny.Sequence
			_ = _133_v43
			_133_v43 = _dafny.SeqOf(_125_v35)
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_130_v40), 0))
			_ = _index7
			var _rhs0 _dafny.Array = _129_v39
			_ = _rhs0
			var _rhs1 _dafny.MultiSet = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_78_cf0, (Companion_Default___.SafeIndex((_127_v37).Cardinality(), _dafny.IntOfUint32((_78_cf0).Cardinality()))).Uint32(), (_124_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_124_v34), 0))).Int()).(bool)), Companion_Default___.Fm4((_80_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))).Int()).(_dafny.Int), _68_v2, (_132_v42).Cardinality(), _71_globalState))))
			_ = _rhs1
			var _rhs2 bool = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_133_v43, Companion_Default___.Fm46(_71_globalState)), _133_v43)
			_ = _rhs2
			var _rhs3 _dafny.Int = ((_80_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_80_v8), 0))).Int()).(_dafny.Int)).Plus((_dafny.IntOfUint32((_67_v1).Cardinality())).Plus(_68_v2))
			_ = _rhs3
			var _lhs0 _dafny.Array = _130_v40
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_130_v40), 0))
			_ = _lhs1
			var _lhs2 *GlobalState = _71_globalState
			_ = _lhs2
			var _lhs3 *GlobalState = _71_globalState
			_ = _lhs3
			(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
			_127_v37 = _rhs1
			_lhs2.F0 = _rhs2
			_lhs3.F5 = _rhs3
		}
	} else if _source2.Is_DC1() {
		var _134___mcc_h2 bool = _source2.Get_().(D0_DC1).Cf2
		_ = _134___mcc_h2
		var _135___mcc_h3 _dafny.Int = _source2.Get_().(D0_DC1).Cf3
		_ = _135___mcc_h3
		var _136___mcc_h4 _dafny.CodePoint = _source2.Get_().(D0_DC1).Cf4
		_ = _136___mcc_h4
		var _137_cf4 _dafny.CodePoint = _136___mcc_h4
		_ = _137_cf4
		var _138_cf3 _dafny.Int = _135___mcc_h3
		_ = _138_cf3
		var _139_cf2 bool = _134___mcc_h2
		_ = _139_cf2
		_73_v5 = _73_v5
		var _140_v44 _dafny.Array
		_ = _140_v44
		var _nw14 _dafny.Array = _dafny.NewArray(_dafny.One)
		_ = _nw14
		_140_v44 = _nw14
		var _141_v45 D20
		_ = _141_v45
		_141_v45 = Companion_D20_.Create_DC46_(_140_v44)
		var _142_v46 _dafny.Map
		_ = _142_v46
		_142_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_141_v45).Dtor_cf84(), _138_cf3)
		_142_v46 = (_142_v46).Update(_140_v44, _68_v2)
		(_71_globalState).F0 = _139_cf2
		(_71_globalState).F5 = _138_cf3
	} else {
		var _143___mcc_h5 D0 = _source2.Get_().(D0_DC2).Cf5
		_ = _143___mcc_h5
		var _144_cf5 D0 = _143___mcc_h5
		_ = _144_cf5
		var _145_v47 _dafny.Array
		_ = _145_v47
		var _nwElement0_1 _dafny.Int = _dafny.IntOfInt64(563)
		_ = _nwElement0_1
		var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(14))
		_ = _nw15
		(_nw15).ArraySet1(_nwElement0_1, 0)
		(_nw15).ArraySet1(_68_v2, 1)
		(_nw15).ArraySet1(_68_v2, 2)
		(_nw15).ArraySet1(_68_v2, 3)
		(_nw15).ArraySet1(_68_v2, 4)
		(_nw15).ArraySet1((_70_v4).Cardinality(), 5)
		(_nw15).ArraySet1(_dafny.IntOfInt64(-185), 6)
		(_nw15).ArraySet1(_68_v2, 7)
		(_nw15).ArraySet1(_68_v2, 8)
		(_nw15).ArraySet1(_68_v2, 9)
		(_nw15).ArraySet1((_dafny.Zero).Minus(_68_v2), 10)
		(_nw15).ArraySet1(_dafny.IntOfInt64(-959), 11)
		(_nw15).ArraySet1(_dafny.IntOfUint32((_69_v3).Cardinality()), 12)
		(_nw15).ArraySet1(_dafny.IntOfUint32((_67_v1).Cardinality()), 13)
		_145_v47 = _nw15
		var _146_v48 _dafny.Array
		_ = _146_v48
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_1
		var _nw16 _dafny.Array
		_ = _nw16
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw16 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = func(_147_i2 _dafny.Int) bool {
				return true
			}
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw16 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw16).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw16).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_146_v48 = _nw16
		var _148_v49 _dafny.Sequence
		_ = _148_v49
		_148_v49 = _dafny.SeqOf(false)
		var _149_v50 _dafny.Map
		_ = _149_v50
		_149_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_148_v49).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("vuxo"))
		var _150_v51 *C7
		_ = _150_v51
		var _nw17 *C7 = New_C7_()
		_ = _nw17
		_nw17.Ctor__(_145_v47, _145_v47, Companion_Default___.Fm2(_68_v2, _71_globalState), _146_v48, _149_v50)
		_150_v51 = _nw17
		var _rhs4 _dafny.Int = _68_v2
		_ = _rhs4
		var _rhs5 _dafny.Int = (_68_v2).Minus(_68_v2)
		_ = _rhs5
		var _lhs4 *GlobalState = _71_globalState
		_ = _lhs4
		var _lhs5 *GlobalState = _71_globalState
		_ = _lhs5
		_lhs4.F5 = _rhs4
		_lhs5.F5 = _rhs5
		var _151_v52 *C0
		_ = _151_v52
		var _nw18 *C0 = New_C0_()
		_ = _nw18
		_nw18.Ctor__(_146_v48, ((_150_v51).Fm14(_71_globalState)).Cmp(_dafny.IntOfUint32((_69_v3).Cardinality())) >= 0)
		_151_v52 = _nw18
		_151_v52 = _151_v52
		var _152_v53 *C8
		_ = _152_v53
		var _nw19 *C8 = New_C8_()
		_ = _nw19
		_nw19.Ctor__()
		_152_v53 = _nw19
	}
	(_71_globalState).F5 = Companion_Default___.SafeDivisionInt(_68_v2, _68_v2)
	var _153_v54 _dafny.Map
	_ = _153_v54
	_153_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v5, Companion_Default___.Fm2(_68_v2, _71_globalState))
	var _154_v55 _dafny.Map
	_ = _154_v55
	_154_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v54, _dafny.MultiSetOf(!(_73_v5)))
	_154_v55 = (_154_v55).Merge(_154_v55)
	_73_v5 = (func() bool {
		if _73_v5 {
			return _73_v5
		}
		return _73_v5
	})()
	var _155_v56 _dafny.Sequence
	_ = _155_v56
	_155_v56 = _dafny.SeqOf(_73_v5, _73_v5)
	var _156_v57 _dafny.Set
	_ = _156_v57
	_156_v57 = _dafny.SetOf((_dafny.Zero).Minus(_68_v2), Companion_Default___.Fm2(_68_v2, _71_globalState))
	var _157_v58 _dafny.Set
	_ = _157_v58
	_157_v58 = _dafny.SetOf(_73_v5, _73_v5)
	var _158_v59 _dafny.Array
	_ = _158_v59
	var _nw20 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
	_ = _nw20
	_158_v59 = _nw20
	var _159_v60 _dafny.Map
	_ = _159_v60
	_159_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v2, _67_v1)
	var _160_v61 T1
	_ = _160_v61
	var _nw21 *C3 = New_C3_()
	_ = _nw21
	_nw21.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v2, (_dafny.Zero).Minus(_68_v2)), _158_v59, _159_v60, _68_v2)
	_160_v61 = _nw21
	var _161_v62 _dafny.Map
	_ = _161_v62
	_161_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v3, _160_v61)
	var _162_v63 _dafny.Set
	_ = _162_v63
	_162_v63 = _dafny.SetOf(_66_v0)
	var _163_v64 _dafny.MultiSet
	_ = _163_v64
	_163_v64 = _dafny.MultiSetOf(_157_v58)
	var _164_v65 _dafny.Array
	_ = _164_v65
	var _nwElement0_2 _dafny.Int = Companion_Default___.Fm2(_68_v2, _71_globalState)
	_ = _nwElement0_2
	var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(25))
	_ = _nw22
	(_nw22).ArraySet1(_nwElement0_2, 0)
	(_nw22).ArraySet1(_dafny.IntOfUint32((_155_v56).Cardinality()), 1)
	(_nw22).ArraySet1(_68_v2, 2)
	(_nw22).ArraySet1(_68_v2, 3)
	(_nw22).ArraySet1(_68_v2, 4)
	(_nw22).ArraySet1(_68_v2, 5)
	(_nw22).ArraySet1(_68_v2, 6)
	(_nw22).ArraySet1(_68_v2, 7)
	(_nw22).ArraySet1(_68_v2, 8)
	(_nw22).ArraySet1(_dafny.IntOfInt64(-22), 9)
	(_nw22).ArraySet1(_68_v2, 10)
	(_nw22).ArraySet1(_68_v2, 11)
	(_nw22).ArraySet1((_156_v57).Cardinality(), 12)
	(_nw22).ArraySet1((_157_v58).Cardinality(), 13)
	(_nw22).ArraySet1(_dafny.IntOfInt64(-993), 14)
	(_nw22).ArraySet1((_161_v62).Cardinality(), 15)
	(_nw22).ArraySet1((_70_v4).Cardinality(), 16)
	(_nw22).ArraySet1(_68_v2, 17)
	(_nw22).ArraySet1(_68_v2, 18)
	(_nw22).ArraySet1(_dafny.IntOfInt64(253), 19)
	(_nw22).ArraySet1((_160_v61).F6(), 20)
	(_nw22).ArraySet1((_162_v63).Cardinality(), 21)
	(_nw22).ArraySet1((_160_v61).F6(), 22)
	(_nw22).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ibypqjp")).Cardinality()), 23)
	(_nw22).ArraySet1((func() _dafny.Int {
		if (_163_v64).Contains(_157_v58) {
			return (_163_v64).Multiplicity(_157_v58)
		}
		return (_160_v61).F6()
	})(), 24)
	_164_v65 = _nw22
	var _165_v66 _dafny.Map
	_ = _165_v66
	_165_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_164_v65, !(_73_v5))
	var _166_v67 D10
	_ = _166_v67
	_166_v67 = Companion_D10_.Create_DC28_(_165_v66)
	_166_v67 = _166_v67
	var _167_v68 _dafny.Map
	_ = _167_v68
	_167_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v5, _160_v61.F9())
	var _168_v69 _dafny.Sequence
	_ = _168_v69
	var _169_v70 bool
	_ = _169_v70
	var _out0 _dafny.Sequence
	_ = _out0
	var _out1 bool
	_ = _out1
	_out0, _out1 = (_160_v61).M6(_73_v5, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_73_v5, _73_v5), Companion_Default___.Fm4((_160_v61).F6(), _68_v2, (_167_v68).Cardinality(), _71_globalState)), _68_v2, _71_globalState)
	_168_v69 = _out0
	_169_v70 = _out1
	(_71_globalState).F2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_67_v1, _67_v1), _67_v1)
	var _170_v71 _dafny.Sequence
	_ = _170_v71
	var _171_v72 bool
	_ = _171_v72
	var _out2 _dafny.Sequence
	_ = _out2
	var _out3 bool
	_ = _out3
	_out2, _out3 = (_160_v61).M6(_73_v5, Companion_Default___.Fm4((_160_v61).F6(), _68_v2, (_160_v61).F6(), _71_globalState), _dafny.IntOfUint32((_dafny.SeqOf(_73_v5, _169_v70, _73_v5, _73_v5)).Cardinality()), _71_globalState)
	_170_v71 = _out2
	_171_v72 = _out3
	var _172_i3 _dafny.Int
	_ = _172_i3
	_172_i3 = _dafny.Zero
	{
		for !(_171_v72) {
			{
				if (_172_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_172_i3 = (_172_i3).Plus(_dafny.One)
				var _173_v73 *C6
				_ = _173_v73
				var _nw23 *C6 = New_C6_()
				_ = _nw23
				_nw23.Ctor__((_160_v61).F6(), _169_v70)
				_173_v73 = _nw23
				_173_v73 = _173_v73
				var _174_v74 _dafny.Sequence
				_ = _174_v74
				var _175_v75 bool
				_ = _175_v75
				var _out4 _dafny.Sequence
				_ = _out4
				var _out5 bool
				_ = _out5
				_out4, _out5 = (_160_v61).M6(((_dafny.Zero).Minus((_160_v61).F6())).Cmp((_160_v61).F6()) <= 0, _dafny.Companion_Sequence_.Concatenate(_155_v56, _155_v56), Companion_Default___.SafeDivisionInt(_68_v2, (_dafny.Zero).Minus((_160_v61).F6())), _71_globalState)
				_174_v74 = _out4
				_175_v75 = _out5
				var _176_v76 _dafny.MultiSet
				_ = _176_v76
				_176_v76 = _dafny.MultiSetOf((_156_v57).IsSubsetOf(_156_v57), true, _73_v5, _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(311))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg25 _dafny.Int) interface{} {
						return coer25(arg25)
					}
				}(func(_177_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('r')
				})), _67_v1), _171_v72)
				_176_v76 = (_176_v76).Intersection(_176_v76)
				_69_v3 = Companion_Default___.Fm40(Companion_Default___.SafeDivisionInt((_dafny.MultiSetFromSeq(_155_v56)).Cardinality(), (_160_v61).F6()), _157_v58, _71_globalState)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	_73_v5 = Companion_Default___.Fm3(_71_globalState)
	var _hi0 _dafny.Int = Companion_Default___.SafeModuloInt((_160_v61).F6(), (_156_v57).Cardinality())
	_ = _hi0
	for _178_i5 := _68_v2; _178_i5.Cmp(_hi0) < 0; _178_i5 = _178_i5.Plus(_dafny.One) {
		var _179_v77 _dafny.Sequence
		_ = _179_v77
		var _180_v78 bool
		_ = _180_v78
		var _out6 _dafny.Sequence
		_ = _out6
		var _out7 bool
		_ = _out7
		_out6, _out7 = (_160_v61).M6(_171_v72, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_169_v70), _155_v56), _dafny.IntOfInt64(481), _71_globalState)
		_179_v77 = _out6
		_180_v78 = _out7
		if !(_169_v70) {
			var _181_v79 _dafny.Map
			_ = _181_v79
			_181_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _180_v78)
			var _182_v80 _dafny.Set
			_ = _182_v80
			_182_v80 = _dafny.SetOf(_181_v79)
			var _183_v81 _dafny.Array
			_ = _183_v81
			var _nwElement0_3 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(175))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_184_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_185_i6 _dafny.Int) _dafny.CodePoint {
					return _184_v0
				}
			})(_66_v0)))
			_ = _nwElement0_3
			var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(8))
			_ = _nw24
			(_nw24).ArraySet1(_nwElement0_3, 0)
			(_nw24).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("kosnitju"), 1)
			(_nw24).ArraySet1(Companion_Default___.Fm17((_160_v61).F6(), _dafny.IntOfUint32((_67_v1).Cardinality()), _71_globalState), 2)
			(_nw24).ArraySet1(_67_v1, 3)
			(_nw24).ArraySet1(_67_v1, 4)
			(_nw24).ArraySet1(_67_v1, 5)
			(_nw24).ArraySet1(_67_v1, 6)
			(_nw24).ArraySet1(_67_v1, 7)
			_183_v81 = _nw24
			var _186_v82 D10
			_ = _186_v82
			_186_v82 = Companion_D10_.Create_DC29_((_160_v61).F6(), _183_v81, _180_v78, _178_i5)
			var _rhs6 _dafny.Array = _164_v65
			_ = _rhs6
			var _rhs7 _dafny.Set = _182_v80
			_ = _rhs7
			var _rhs8 bool = (func() bool {
				if (_181_v79).Contains(_169_v70) {
					return (_181_v79).Get(_169_v70).(bool)
				}
				return (_186_v82).Dtor_cf55()
			})()
			_ = _rhs8
			var _lhs6 *GlobalState = _71_globalState
			_ = _lhs6
			_164_v65 = _rhs6
			_182_v80 = _rhs7
			_lhs6.F0 = _rhs8
			var _187_v83 _dafny.Array
			_ = _187_v83
			_187_v83 = _160_v61.F9()
			var _188_v84 *C7
			_ = _188_v84
			var _nw25 *C7 = New_C7_()
			_ = _nw25
			_nw25.Ctor__(_164_v65, _164_v65, Companion_Default___.SafeDivisionInt((_160_v61).F6(), _dafny.IntOfUint32(((_168_v69).Select((Companion_Default___.SafeIndex((_160_v61).F6(), _dafny.IntOfUint32((_168_v69).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())), (_187_v83), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v2, _67_v1))
			_188_v84 = _nw25
			var _189_v85 *C1
			_ = _189_v85
			var _nw26 *C1 = New_C1_()
			_ = _nw26
			_nw26.Ctor__((_dafny.IntOfInt64(684)).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_67_v1).Cardinality()))))
			_189_v85 = _nw26
			_189_v85 = _189_v85
			var _190_v86 D11
			_ = _190_v86
			_190_v86 = Companion_D11_.Create_DC31_((_160_v61).F6(), (_dafny.Zero).Minus((_160_v61).F6()))
			_190_v86 = Companion_Default___.Fm56(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-459))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_191_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_192_i7 _dafny.Int) _dafny.CodePoint {
					return _191_v0
				}
			})(_66_v0))), _73_v5, _67_v1, _71_globalState)
			var _193_v87 T0
			_ = _193_v87
			var _nw27 *C1 = New_C1_()
			_ = _nw27
			_nw27.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_69_v3, (Companion_Default___.SafeIndex((_160_v61).F6(), _dafny.IntOfUint32((_69_v3).Cardinality()))).Uint32(), _178_i5)).Cardinality()))
			_193_v87 = _nw27
			var _194_v88 _dafny.Sequence
			_ = _194_v88
			_194_v88 = _dafny.SeqOf(_193_v87, _193_v87, _193_v87, _193_v87)
			var _195_v89 _dafny.Array
			_ = _195_v89
			var _nwElement0_4 T0 = _193_v87
			_ = _nwElement0_4
			var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(22))
			_ = _nw28
			(_nw28).ArraySet1(_nwElement0_4, 0)
			(_nw28).ArraySet1(_193_v87, 1)
			(_nw28).ArraySet1(_193_v87, 2)
			(_nw28).ArraySet1(_193_v87, 3)
			(_nw28).ArraySet1(_193_v87, 4)
			(_nw28).ArraySet1(_193_v87, 5)
			(_nw28).ArraySet1(_193_v87, 6)
			(_nw28).ArraySet1(_193_v87, 7)
			(_nw28).ArraySet1(_193_v87, 8)
			(_nw28).ArraySet1(_193_v87, 9)
			(_nw28).ArraySet1(_193_v87, 10)
			(_nw28).ArraySet1(_193_v87, 11)
			(_nw28).ArraySet1(_193_v87, 12)
			(_nw28).ArraySet1(_193_v87, 13)
			(_nw28).ArraySet1((_194_v88).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm2((_193_v87).F6(), _71_globalState), _dafny.IntOfUint32((_194_v88).Cardinality()))).Uint32()).(T0), 14)
			(_nw28).ArraySet1(_193_v87, 15)
			(_nw28).ArraySet1(_193_v87, 16)
			(_nw28).ArraySet1(_193_v87, 17)
			(_nw28).ArraySet1(_193_v87, 18)
			(_nw28).ArraySet1(_193_v87, 19)
			(_nw28).ArraySet1(_193_v87, 20)
			(_nw28).ArraySet1(_193_v87, 21)
			_195_v89 = _nw28
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_195_v89), 0))
			_ = _index8
			(_195_v89).ArraySet1(_193_v87, (_index8).Int())
			var _196_v90 D2
			_ = _196_v90
			_196_v90 = Companion_D2_.Create_DC6_((_160_v61).F6(), _188_v84.F12)
			var _197_v91 _dafny.MultiSet
			_ = _197_v91
			_197_v91 = _dafny.MultiSetOf((_196_v90).Dtor_cf14(), _188_v84.F12)
			var _198_v92 _dafny.Map
			_ = _198_v92
			_198_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_70_v4).Update((_160_v61).F6(), (_160_v61).F6())).Cardinality(), !(_73_v5))
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_195_v89), 0))
			_ = _index9
			var _rhs9 T0 = _193_v87
			_ = _rhs9
			var _rhs10 bool = (((_198_v92).Cardinality()).Plus(Companion_Default___.Fm2((_160_v61).F6(), _71_globalState))).Cmp((_160_v61).F6()) <= 0
			_ = _rhs10
			var _rhs11 _dafny.Sequence = _155_v56
			_ = _rhs11
			var _rhs12 _dafny.CodePoint = (_67_v1).Select((Companion_Default___.SafeIndex(_68_v2, _dafny.IntOfUint32((_67_v1).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs12
			var _rhs13 _dafny.MultiSet = _197_v91
			_ = _rhs13
			var _lhs7 _dafny.Array = _195_v89
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_195_v89), 0))
			_ = _lhs8
			(_lhs7).ArraySet1(_rhs9, (_lhs8).Int())
			_171_v72 = _rhs10
			_155_v56 = _rhs11
			_66_v0 = _rhs12
			_197_v91 = _rhs13
		} else {
			(_160_v61).F9_set_(_160_v61.F9())
			var _199_v93 D4
			_ = _199_v93
			_199_v93 = Companion_D4_.Create_DC13_(_69_v3, !(_169_v70), (_160_v61).F6())
			var _200_v94 D14
			_ = _200_v94
			_200_v94 = Companion_D14_.Create_DC37_(_73_v5, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(743))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}(func(_201_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})))
			_199_v93 = Companion_D4_.Create_DC13_((func() _dafny.Sequence {
				if _169_v70 {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(569))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg29 _dafny.Int) interface{} {
							return coer29(arg29)
						}
					}((func(_202_v70 bool) func(_dafny.Int) _dafny.Int {
						return func(_203_i8 _dafny.Int) _dafny.Int {
							return (_dafny.MultiSetOf(_202_v70)).Cardinality()
						}
					})(_169_v70)))
				}
				return _69_v3
			})(), (_200_v94).Dtor_cf70(), ((_dafny.Zero).Minus((_160_v61).F6())).Times((_160_v61).F6()))
			var _204_v95 _dafny.Array
			_ = _204_v95
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(21))
			_ = _nw29
			_204_v95 = _nw29
			var _205_v96 _dafny.Map
			_ = _205_v96
			_205_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_178_i5, _180_v78)
			var _206_v97 _dafny.Map
			_ = _206_v97
			_206_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_205_v96, _169_v70)
			var _207_v98 D2
			_ = _207_v98
			_207_v98 = Companion_D2_.Create_DC6_((_206_v97).Cardinality(), _164_v65)
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_204_v95), 0))
			_ = _index10
			(_204_v95).ArraySet1(_207_v98, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_204_v95), 0))
			_ = _index11
			(_204_v95).ArraySet1(Companion_D2_.Create_DC6_(_178_i5, _164_v65), (_index11).Int())
			_155_v56 = _dafny.Companion_Sequence_.Concatenate(_155_v56, _155_v56)
			_66_v0 = _66_v0
		}
		var _208_v99 _dafny.Array
		_ = _208_v99
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_2
		var _nw30 _dafny.Array
		_ = _nw30
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw30 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.CodePoint = (func(_209_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_210_i10 _dafny.Int) _dafny.CodePoint {
					return _209_v0
				}
			})(_66_v0)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw30 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw30).ArraySet1CodePoint(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw30).ArraySet1CodePoint(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_208_v99 = _nw30
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_208_v99), 0))
		_ = _index12
		(_208_v99).ArraySet1CodePoint(_66_v0, (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_208_v99), 0))
		_ = _index13
		(_208_v99).ArraySet1CodePoint(_66_v0, (_index13).Int())
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_164_v65), 0))
		_ = _index14
		(_164_v65).ArraySet1((_160_v61).F6(), (_index14).Int())
		var _211_v100 _dafny.Array
		_ = _211_v100
		var _nw31 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(11))
		_ = _nw31
		_211_v100 = _nw31
		var _212_v101 D1
		_ = _212_v101
		_212_v101 = Companion_D1_.Create_DC4_(func(_pat_let4_0 D0) D0 {
			return func(_213_dt__update__tmp_h2 D0) D0 {
				return func(_pat_let5_0 _dafny.Int) D0 {
					return func(_214_dt__update_hcf1_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_((_213_dt__update__tmp_h2).Dtor_cf0(), _214_dt__update_hcf1_h0)
					}(_pat_let5_0)
				}(_dafny.IntOfInt64(-120))
			}(_pat_let4_0)
		}(Companion_D0_.Create_DC0_(_155_v56, _178_i5)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v72, _171_v72)).Cardinality(), _67_v1, _169_v70, _68_v2)
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_211_v100), 0))
		_ = _index15
		(_211_v100).ArraySet1(_212_v101, (_index15).Int())
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_164_v65), 0))
		_ = _index16
		(_164_v65).ArraySet1(_178_i5, (_index16).Int())
		var _215_v102 _dafny.Map
		_ = _215_v102
		_215_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_208_v99).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_208_v99), 0))).Int()), _73_v5)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_164_v65), 0))
		_ = _index17
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_211_v100), 0))
		_ = _index18
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_164_v65), 0))
		_ = _index19
		var _rhs14 _dafny.Int = Companion_Default___.Fm2(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(101))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}((func(_216_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_217_i11 _dafny.Int) _dafny.CodePoint {
				return _216_v0
			}
		})(_66_v0))), _dafny.UnicodeSeqOfUtf8Bytes("r"))).Cardinality()), _71_globalState)
		_ = _rhs14
		var _rhs15 _dafny.Int = Companion_Default___.Fm2(_178_i5, _71_globalState)
		_ = _rhs15
		var _rhs16 D1 = _212_v101
		_ = _rhs16
		var _rhs17 _dafny.Int = Companion_Default___.SafeDivisionInt((_215_v102).Cardinality(), _68_v2)
		_ = _rhs17
		var _rhs18 _dafny.Int = _178_i5
		_ = _rhs18
		var _lhs9 _dafny.Array = _164_v65
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_164_v65), 0))
		_ = _lhs10
		var _lhs11 _dafny.Array = _211_v100
		_ = _lhs11
		var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_211_v100), 0))
		_ = _lhs12
		var _lhs13 _dafny.Array = _164_v65
		_ = _lhs13
		var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_164_v65), 0))
		_ = _lhs14
		(_lhs9).ArraySet1(_rhs14, (_lhs10).Int())
		_68_v2 = _rhs15
		(_lhs11).ArraySet1(_rhs16, (_lhs12).Int())
		(_lhs13).ArraySet1(_rhs17, (_lhs14).Int())
		_68_v2 = _rhs18
	}
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_164_v65), 0))
	_ = _index20
	(_164_v65).ArraySet1(Companion_Default___.SafeDivisionInt(_68_v2, (_dafny.SetOf(_171_v72)).Cardinality()), (_index20).Int())
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_164_v65), 0))
	_ = _index21
	(_164_v65).ArraySet1((_160_v61).F6(), (_index21).Int())
	for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_164_v65), 0))); ; {
		_guard_loop_0, _ok8 := _iter8()
		if !_ok8 {
			break
		}
		var _218_i12 _dafny.Int
		_218_i12 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_218_i12).Sign() != -1) && ((_218_i12).Cmp(_dafny.ArrayLen((_164_v65), 0)) < 0)) {
			(_164_v65).ArraySet1(Companion_Default___.SafeModuloInt(_218_i12, (_153_v54).Cardinality()), (_218_i12).Int())
		}
	}
	var _arr0 _dafny.Array = _160_v61.F9()
	_ = _arr0
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_160_v61.F9()), 0))
	_ = _index22
	(_arr0).ArraySet1(!(_171_v72) || (_171_v72), (_index22).Int())
	var _arr1 _dafny.Array = _160_v61.F9()
	_ = _arr1
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_160_v61.F9()), 0))
	_ = _index23
	var _rhs19 bool = ((_164_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_164_v65), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(309)) > 0
	_ = _rhs19
	var _rhs20 bool = _73_v5
	_ = _rhs20
	var _lhs15 _dafny.Array = _160_v61.F9()
	_ = _lhs15
	var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_160_v61.F9()), 0))
	_ = _lhs16
	_171_v72 = _rhs19
	(_lhs15).ArraySet1(_rhs20, (_lhs16).Int())
	var _219_v103 _dafny.MultiSet
	_ = _219_v103
	_219_v103 = _dafny.MultiSetOf(_68_v2, (_160_v61).F6())
	var _220_v104 _dafny.CodePoint
	_ = _220_v104
	var _221_v105 _dafny.Int
	_ = _221_v105
	var _222_v106 _dafny.Array
	_ = _222_v106
	var _223_v107 _dafny.Int
	_ = _223_v107
	var _out8 _dafny.CodePoint
	_ = _out8
	var _out9 _dafny.Int
	_ = _out9
	var _out10 _dafny.Array
	_ = _out10
	var _out11 _dafny.Int
	_ = _out11
	_out8, _out9, _out10, _out11 = (_160_v61).M7((_219_v103).Union(Companion_Default___.Fm19(_dafny.IntOfInt64(534), _71_globalState)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_169_v70, _169_v70), _155_v56)).Cardinality()), _71_globalState)
	_220_v104 = _out8
	_221_v105 = _out9
	_222_v106 = _out10
	_223_v107 = _out11
	var _224_i13 _dafny.Int
	_ = _224_i13
	_224_i13 = _dafny.Zero
	{
		for Companion_Default___.Fm3(_71_globalState) {
			{
				if (_224_i13).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_224_i13 = (_224_i13).Plus(_dafny.One)
				var _225_v108 _dafny.Array
				_ = _225_v108
				var _nw32 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(17))
				_ = _nw32
				_225_v108 = _nw32
				var _226_v109 *C0
				_ = _226_v109
				var _nw33 *C0 = New_C0_()
				_ = _nw33
				_nw33.Ctor__(_160_v61.F9(), _73_v5)
				_226_v109 = _nw33
				var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_225_v108), 0))
				_ = _index24
				(_225_v108).ArraySet1(_226_v109, (_index24).Int())
				var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_225_v108), 0))
				_ = _index25
				(_225_v108).ArraySet1((func() *C0 {
					if _169_v70 {
						return _226_v109
					}
					return _226_v109
				})(), (_index25).Int())
				_68_v2 = (_dafny.Zero).Minus((_160_v61).F6())
				_169_v70 = _73_v5
				if _171_v72 {
					var _227_v110 _dafny.CodePoint
					_ = _227_v110
					var _228_v111 _dafny.Int
					_ = _228_v111
					var _229_v112 _dafny.Array
					_ = _229_v112
					var _230_v113 _dafny.Int
					_ = _230_v113
					var _out12 _dafny.CodePoint
					_ = _out12
					var _out13 _dafny.Int
					_ = _out13
					var _out14 _dafny.Array
					_ = _out14
					var _out15 _dafny.Int
					_ = _out15
					_out12, _out13, _out14, _out15 = (_160_v61).M7(_219_v103, (func() _dafny.Int {
						if _171_v72 {
							return (_164_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_164_v65), 0))).Int()).(_dafny.Int)
						}
						return (_164_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_164_v65), 0))).Int()).(_dafny.Int)
					})(), _71_globalState)
					_227_v110 = _out12
					_228_v111 = _out13
					_229_v112 = _out14
					_230_v113 = _out15
					var _231_v114 _dafny.Sequence
					_ = _231_v114
					_231_v114 = _dafny.SeqOf(_164_v65, _164_v65)
					_231_v114 = _231_v114
					var _232_v115 _dafny.MultiSet
					_ = _232_v115
					_232_v115 = _dafny.MultiSetOf(false, (_155_v56).Select((Companion_Default___.SafeIndex(_230_v113, _dafny.IntOfUint32((_155_v56).Cardinality()))).Uint32()).(bool))
					var _233_v116 _dafny.Map
					_ = _233_v116
					_233_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v5, _226_v109.F20)
					_223_v107 = (((_232_v115).Cardinality()).Times(_228_v111)).Minus((_dafny.IntOfUint32((_67_v1).Cardinality())).Plus((_233_v116).Cardinality()))
					var _234_v117 _dafny.Map
					_ = _234_v117
					_234_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_74_v6, true)
					_234_v117 = (_234_v117).Update(_74_v6, (_160_v61.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(541), _dafny.ArrayLen((_160_v61.F9()), 0))).Int()).(bool))
					_220_v104 = Companion_Default___.Fm0(Companion_Default___.Fm2((_160_v61).F6(), _71_globalState), (Companion_Default___.Fm1(_71_globalState)).Union(_156_v57), _dafny.Companion_Sequence_.Equal(_67_v1, _dafny.UnicodeSeqOfUtf8Bytes("eiyyqrevp")), _71_globalState)
				} else {
					_169_v70 = (_68_v2).Cmp(Companion_Default___.Fm2(_dafny.IntOfInt64(-764), _71_globalState)) != 0
					_223_v107 = (_dafny.Zero).Minus((_160_v61).F6())
					var _rhs21 bool = (_226_v109.F20) || (_171_v72)
					_ = _rhs21
					var _rhs22 T1 = _160_v61
					_ = _rhs22
					var _rhs23 bool = ((_164_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_164_v65), 0))).Int()).(_dafny.Int)).Cmp((_68_v2).Minus((_69_v3).Select((Companion_Default___.SafeIndex((_164_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_164_v65), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_69_v3).Cardinality()))).Uint32()).(_dafny.Int))) <= 0
					_ = _rhs23
					var _rhs24 _dafny.Int = _223_v107
					_ = _rhs24
					var _rhs25 _dafny.Set = _156_v57
					_ = _rhs25
					var _lhs17 *GlobalState = _71_globalState
					_ = _lhs17
					var _lhs18 *GlobalState = _71_globalState
					_ = _lhs18
					_73_v5 = _rhs21
					_160_v61 = _rhs22
					_lhs17.F0 = _rhs23
					_lhs18.F5 = _rhs24
					_156_v57 = _rhs25
					var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_164_v65), 0))
					_ = _index26
					(_164_v65).ArraySet1(_68_v2, (_index26).Int())
					_155_v56 = _155_v56
				}
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	_dafny.Print(_66_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_67_v1, _dafny.SeqOf(_dafny.CodePoint('f'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_68_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_69_v3, _dafny.SeqOf(_dafny.IntOfInt64(-2), _dafny.IntOfInt64(-354), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_70_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_71_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_71_globalState).F1(), _dafny.SeqOf(_dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'), _dafny.CodePoint('p'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_71_globalState.F2.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F4()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_71_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_73_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_74_v6).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_74_v6).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_74_v6).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v54).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(663))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_154_v55).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(663)), _dafny.MultiSetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_155_v56, _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v57).Equals(_dafny.SetOf(_dafny.IntOfInt64(-291), _dafny.IntOfInt64(663))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_v58).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v59).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v59).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v59).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v59).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v60).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(291), _dafny.SeqOf(_dafny.CodePoint('f'), _dafny.CodePoint('s')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v61.F9()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v61.F9()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v61.F9()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v61.F9()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_160_v61).F10()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(291), _dafny.SeqOf(_dafny.CodePoint('f'), _dafny.CodePoint('s')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v61).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v62).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_v63).Equals(_dafny.SetOf(_dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v64).Equals(_dafny.MultiSetOf(_dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v65).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v66).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_166_v67).Dtor_cf52()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_167_v68).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_168_v69, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("yhclj"), _dafny.UnicodeSeqOfUtf8Bytes("jcnirnr"), _dafny.UnicodeSeqOfUtf8Bytes("ttxfurlkd"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_169_v70)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_170_v71, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("yhclj"), _dafny.UnicodeSeqOfUtf8Bytes("jcnirnr"), _dafny.UnicodeSeqOfUtf8Bytes("ttxfurlkd"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_171_v72)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_172_i3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_219_v103).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(291), _dafny.IntOfInt64(291))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_220_v104)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_221_v105)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v106).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v106).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v106).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v106).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v106).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v106).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v106).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_223_v107)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_224_i13)
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
	Cf0 _dafny.Sequence
	Cf1 _dafny.Int
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Sequence, Cf1 _dafny.Int) D0 {
	return D0{D0_DC0{Cf0, Cf1}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
	Cf2 bool
	Cf3 _dafny.Int
	Cf4 _dafny.CodePoint
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf2 bool, Cf3 _dafny.Int, Cf4 _dafny.CodePoint) D0 {
	return D0{D0_DC1{Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf5 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf5 D0) D0 {
	return D0{D0_DC2{Cf5}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.EmptySeq, _dafny.Zero)
}

func (_this D0) Dtor_cf0() _dafny.Sequence {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.CodePoint {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() D0 {
	return _this.Get_().(D0_DC2).Cf5
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf0.Equals(data2.Cf0) && data1.Cf1.Cmp(data2.Cf1) == 0
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4 == data2.Cf4
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf5.Equals(data2.Cf5)
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
	Cf7  D0
	Cf8  _dafny.Int
	Cf9  _dafny.Sequence
	Cf10 bool
	Cf11 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf7 D0, Cf8 _dafny.Int, Cf9 _dafny.Sequence, Cf10 bool, Cf11 _dafny.Int) D1 {
	return D1{D1_DC4{Cf7, Cf8, Cf9, Cf10, Cf11}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf6 _dafny.MultiSet
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf6 _dafny.MultiSet) D1 {
	return D1{D1_DC3{Cf6}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(Companion_D0_.Default(), _dafny.Zero, _dafny.EmptySeq, false, _dafny.Zero)
}

func (_this D1) Dtor_cf7() D0 {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() bool {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) Dtor_cf6() _dafny.MultiSet {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + data.Cf9.VerbatimString(true) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf7.Equals(data2.Cf7) && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Equals(data2.Cf9) && data1.Cf10 == data2.Cf10 && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf6.Equals(data2.Cf6)
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
	Cf13 _dafny.Int
	Cf14 _dafny.Array
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf13 _dafny.Int, Cf14 _dafny.Array) D2 {
	return D2{D2_DC6{Cf13, Cf14}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf15 _dafny.Sequence
	Cf16 _dafny.Array
	Cf17 _dafny.Int
	Cf18 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf15 _dafny.Sequence, Cf16 _dafny.Array, Cf17 _dafny.Int, Cf18 _dafny.Int) D2 {
	return D2{D2_DC7{Cf15, Cf16, Cf17, Cf18}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC8 struct {
	Cf19 bool
	Cf20 bool
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf19 bool, Cf20 bool) D2 {
	return D2{D2_DC8{Cf19, Cf20}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC5 struct {
	Cf12 _dafny.Map
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf12 _dafny.Map) D2 {
	return D2{D2_DC5{Cf12}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D2) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Array {
	return _this.Get_().(D2_DC6).Cf14
}

func (_this D2) Dtor_cf15() _dafny.Sequence {
	return _this.Get_().(D2_DC7).Cf15
}

func (_this D2) Dtor_cf16() _dafny.Array {
	return _this.Get_().(D2_DC7).Cf16
}

func (_this D2) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf17
}

func (_this D2) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf18
}

func (_this D2) Dtor_cf19() bool {
	return _this.Get_().(D2_DC8).Cf19
}

func (_this D2) Dtor_cf20() bool {
	return _this.Get_().(D2_DC8).Cf20
}

func (_this D2) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D2_DC5).Cf12
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14 == data2.Cf14
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf15.Equals(data2.Cf15) && data1.Cf16 == data2.Cf16 && data1.Cf17.Cmp(data2.Cf17) == 0 && data1.Cf18.Cmp(data2.Cf18) == 0
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf19 == data2.Cf19 && data1.Cf20 == data2.Cf20
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
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
	Cf22 _dafny.Map
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf22 _dafny.Map) D3 {
	return D3{D3_DC10{Cf22}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC9 struct {
	Cf21 _dafny.Sequence
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf21 _dafny.Sequence) D3 {
	return D3{D3_DC9{Cf21}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC11 struct {
	Cf23 D3
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf23 D3) D3 {
	return D3{D3_DC11{Cf23}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(_dafny.EmptyMap)
}

func (_this D3) Dtor_cf22() _dafny.Map {
	return _this.Get_().(D3_DC10).Cf22
}

func (_this D3) Dtor_cf21() _dafny.Sequence {
	return _this.Get_().(D3_DC9).Cf21
}

func (_this D3) Dtor_cf23() D3 {
	return _this.Get_().(D3_DC11).Cf23
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf23) + ")"
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
			return ok && data1.Cf22.Equals(data2.Cf22)
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf21.Equals(data2.Cf21)
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf23.Equals(data2.Cf23)
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
	Cf25 _dafny.Sequence
	Cf26 bool
	Cf27 _dafny.Int
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf25 _dafny.Sequence, Cf26 bool, Cf27 _dafny.Int) D4 {
	return D4{D4_DC13{Cf25, Cf26, Cf27}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC14 struct {
	Cf28 bool
	Cf29 _dafny.Map
	Cf30 _dafny.Int
	Cf31 _dafny.Map
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf28 bool, Cf29 _dafny.Map, Cf30 _dafny.Int, Cf31 _dafny.Map) D4 {
	return D4{D4_DC14{Cf28, Cf29, Cf30, Cf31}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC15 struct {
	Cf32 bool
	Cf33 bool
	Cf34 _dafny.Sequence
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf32 bool, Cf33 bool, Cf34 _dafny.Sequence) D4 {
	return D4{D4_DC15{Cf32, Cf33, Cf34}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

type D4_DC12 struct {
	Cf24 _dafny.Sequence
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf24 _dafny.Sequence) D4 {
	return D4{D4_DC12{Cf24}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_(_dafny.EmptySeq, false, _dafny.Zero)
}

func (_this D4) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D4_DC13).Cf25
}

func (_this D4) Dtor_cf26() bool {
	return _this.Get_().(D4_DC13).Cf26
}

func (_this D4) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf27
}

func (_this D4) Dtor_cf28() bool {
	return _this.Get_().(D4_DC14).Cf28
}

func (_this D4) Dtor_cf29() _dafny.Map {
	return _this.Get_().(D4_DC14).Cf29
}

func (_this D4) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf30
}

func (_this D4) Dtor_cf31() _dafny.Map {
	return _this.Get_().(D4_DC14).Cf31
}

func (_this D4) Dtor_cf32() bool {
	return _this.Get_().(D4_DC15).Cf32
}

func (_this D4) Dtor_cf33() bool {
	return _this.Get_().(D4_DC15).Cf33
}

func (_this D4) Dtor_cf34() _dafny.Sequence {
	return _this.Get_().(D4_DC15).Cf34
}

func (_this D4) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D4_DC12).Cf24
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + data.Cf34.VerbatimString(true) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf24) + ")"
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
			return ok && data1.Cf25.Equals(data2.Cf25) && data1.Cf26 == data2.Cf26 && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf28 == data2.Cf28 && data1.Cf29.Equals(data2.Cf29) && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31.Equals(data2.Cf31)
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33 && data1.Cf34.Equals(data2.Cf34)
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf24.Equals(data2.Cf24)
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
	Cf36 _dafny.Sequence
	Cf37 bool
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf36 _dafny.Sequence, Cf37 bool) D5 {
	return D5{D5_DC17{Cf36, Cf37}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC16 struct {
	Cf35 _dafny.Set
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf35 _dafny.Set) D5 {
	return D5{D5_DC16{Cf35}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC17_(_dafny.EmptySeq, false)
}

func (_this D5) Dtor_cf36() _dafny.Sequence {
	return _this.Get_().(D5_DC17).Cf36
}

func (_this D5) Dtor_cf37() bool {
	return _this.Get_().(D5_DC17).Cf37
}

func (_this D5) Dtor_cf35() _dafny.Set {
	return _this.Get_().(D5_DC16).Cf35
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf35) + ")"
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
			return ok && data1.Cf36.Equals(data2.Cf36) && data1.Cf37 == data2.Cf37
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf35.Equals(data2.Cf35)
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
	Cf39 _dafny.Int
	Cf40 _dafny.Int
	Cf41 _dafny.Int
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf39 _dafny.Int, Cf40 _dafny.Int, Cf41 _dafny.Int) D6 {
	return D6{D6_DC19{Cf39, Cf40, Cf41}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

type D6_DC18 struct {
	Cf38 T1
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf38 T1) D6 {
	return D6{D6_DC18{Cf38}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC20 struct {
	Cf42 D6
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_(Cf42 D6) D6 {
	return D6{D6_DC20{Cf42}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC19_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D6) Dtor_cf39() _dafny.Int {
	return _this.Get_().(D6_DC19).Cf39
}

func (_this D6) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D6_DC19).Cf40
}

func (_this D6) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D6_DC19).Cf41
}

func (_this D6) Dtor_cf38() T1 {
	return _this.Get_().(D6_DC18).Cf38
}

func (_this D6) Dtor_cf42() D6 {
	return _this.Get_().(D6_DC20).Cf42
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf38) + ")"
		}
	case D6_DC20:
		{
			return "D6.DC20" + "(" + _dafny.String(data.Cf42) + ")"
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
			return ok && data1.Cf39.Cmp(data2.Cf39) == 0 && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41.Cmp(data2.Cf41) == 0
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && _dafny.AreEqual(data1.Cf38, data2.Cf38)
		}
	case D6_DC20:
		{
			data2, ok := other.Get_().(D6_DC20)
			return ok && data1.Cf42.Equals(data2.Cf42)
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

type D7_DC22 struct {
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_() D7 {
	return D7{D7_DC22{}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

type D7_DC21 struct {
	Cf43 _dafny.Array
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf43 _dafny.Array) D7 {
	return D7{D7_DC21{Cf43}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC22_()
}

func (_this D7) Dtor_cf43() _dafny.Array {
	return _this.Get_().(D7_DC21).Cf43
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC22:
		{
			return "D7.DC22"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC22:
		{
			_, ok := other.Get_().(D7_DC22)
			return ok
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf43 == data2.Cf43
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
	Cf44 _dafny.Array
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf44 _dafny.Array) D8 {
	return D8{D8_DC23{Cf44}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D8) Dtor_cf44() _dafny.Array {
	return _this.Get_().(D8_DC23).Cf44
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf44) + ")"
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
			return ok && data1.Cf44 == data2.Cf44
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
	Cf46 _dafny.Int
	Cf47 _dafny.Int
	Cf48 _dafny.Int
	Cf49 _dafny.CodePoint
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf46 _dafny.Int, Cf47 _dafny.Int, Cf48 _dafny.Int, Cf49 _dafny.CodePoint) D9 {
	return D9{D9_DC25{Cf46, Cf47, Cf48, Cf49}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC26 struct {
	Cf50 _dafny.Array
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf50 _dafny.Array) D9 {
	return D9{D9_DC26{Cf50}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC24 struct {
	Cf45 _dafny.MultiSet
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf45 _dafny.MultiSet) D9 {
	return D9{D9_DC24{Cf45}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC27 struct {
	Cf51 D9
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf51 D9) D9 {
	return D9{D9_DC27{Cf51}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC25_(_dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.CodePoint('D'))
}

func (_this D9) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf46
}

func (_this D9) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf47
}

func (_this D9) Dtor_cf48() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf48
}

func (_this D9) Dtor_cf49() _dafny.CodePoint {
	return _this.Get_().(D9_DC25).Cf49
}

func (_this D9) Dtor_cf50() _dafny.Array {
	return _this.Get_().(D9_DC26).Cf50
}

func (_this D9) Dtor_cf45() _dafny.MultiSet {
	return _this.Get_().(D9_DC24).Cf45
}

func (_this D9) Dtor_cf51() D9 {
	return _this.Get_().(D9_DC27).Cf51
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ")"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf50) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf45) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf51) + ")"
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
			return ok && data1.Cf46.Cmp(data2.Cf46) == 0 && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48.Cmp(data2.Cf48) == 0 && data1.Cf49 == data2.Cf49
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf50 == data2.Cf50
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf45.Equals(data2.Cf45)
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf51.Equals(data2.Cf51)
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
	Cf53 _dafny.Int
	Cf54 _dafny.Array
	Cf55 bool
	Cf56 _dafny.Int
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf53 _dafny.Int, Cf54 _dafny.Array, Cf55 bool, Cf56 _dafny.Int) D10 {
	return D10{D10_DC29{Cf53, Cf54, Cf55, Cf56}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

type D10_DC28 struct {
	Cf52 _dafny.Map
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf52 _dafny.Map) D10 {
	return D10{D10_DC28{Cf52}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC29_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.Zero)
}

func (_this D10) Dtor_cf53() _dafny.Int {
	return _this.Get_().(D10_DC29).Cf53
}

func (_this D10) Dtor_cf54() _dafny.Array {
	return _this.Get_().(D10_DC29).Cf54
}

func (_this D10) Dtor_cf55() bool {
	return _this.Get_().(D10_DC29).Cf55
}

func (_this D10) Dtor_cf56() _dafny.Int {
	return _this.Get_().(D10_DC29).Cf56
}

func (_this D10) Dtor_cf52() _dafny.Map {
	return _this.Get_().(D10_DC28).Cf52
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ")"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf52) + ")"
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
			return ok && data1.Cf53.Cmp(data2.Cf53) == 0 && data1.Cf54 == data2.Cf54 && data1.Cf55 == data2.Cf55 && data1.Cf56.Cmp(data2.Cf56) == 0
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf52.Equals(data2.Cf52)
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
	Cf58 _dafny.Int
	Cf59 _dafny.Int
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf58 _dafny.Int, Cf59 _dafny.Int) D11 {
	return D11{D11_DC31{Cf58, Cf59}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

type D11_DC32 struct {
	Cf60 _dafny.MultiSet
	Cf61 _dafny.Sequence
}

func (D11_DC32) isD11() {}

func (CompanionStruct_D11_) Create_DC32_(Cf60 _dafny.MultiSet, Cf61 _dafny.Sequence) D11 {
	return D11{D11_DC32{Cf60, Cf61}}
}

func (_this D11) Is_DC32() bool {
	_, ok := _this.Get_().(D11_DC32)
	return ok
}

type D11_DC33 struct {
	Cf62 _dafny.Int
	Cf63 bool
	Cf64 _dafny.Int
	Cf65 bool
	Cf66 _dafny.Set
}

func (D11_DC33) isD11() {}

func (CompanionStruct_D11_) Create_DC33_(Cf62 _dafny.Int, Cf63 bool, Cf64 _dafny.Int, Cf65 bool, Cf66 _dafny.Set) D11 {
	return D11{D11_DC33{Cf62, Cf63, Cf64, Cf65, Cf66}}
}

func (_this D11) Is_DC33() bool {
	_, ok := _this.Get_().(D11_DC33)
	return ok
}

type D11_DC30 struct {
	Cf57 _dafny.Set
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf57 _dafny.Set) D11 {
	return D11{D11_DC30{Cf57}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC31_(_dafny.Zero, _dafny.Zero)
}

func (_this D11) Dtor_cf58() _dafny.Int {
	return _this.Get_().(D11_DC31).Cf58
}

func (_this D11) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D11_DC31).Cf59
}

func (_this D11) Dtor_cf60() _dafny.MultiSet {
	return _this.Get_().(D11_DC32).Cf60
}

func (_this D11) Dtor_cf61() _dafny.Sequence {
	return _this.Get_().(D11_DC32).Cf61
}

func (_this D11) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D11_DC33).Cf62
}

func (_this D11) Dtor_cf63() bool {
	return _this.Get_().(D11_DC33).Cf63
}

func (_this D11) Dtor_cf64() _dafny.Int {
	return _this.Get_().(D11_DC33).Cf64
}

func (_this D11) Dtor_cf65() bool {
	return _this.Get_().(D11_DC33).Cf65
}

func (_this D11) Dtor_cf66() _dafny.Set {
	return _this.Get_().(D11_DC33).Cf66
}

func (_this D11) Dtor_cf57() _dafny.Set {
	return _this.Get_().(D11_DC30).Cf57
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D11_DC32:
		{
			return "D11.DC32" + "(" + _dafny.String(data.Cf60) + ", " + data.Cf61.VerbatimString(true) + ")"
		}
	case D11_DC33:
		{
			return "D11.DC33" + "(" + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ", " + _dafny.String(data.Cf66) + ")"
		}
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf57) + ")"
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
			return ok && data1.Cf58.Cmp(data2.Cf58) == 0 && data1.Cf59.Cmp(data2.Cf59) == 0
		}
	case D11_DC32:
		{
			data2, ok := other.Get_().(D11_DC32)
			return ok && data1.Cf60.Equals(data2.Cf60) && data1.Cf61.Equals(data2.Cf61)
		}
	case D11_DC33:
		{
			data2, ok := other.Get_().(D11_DC33)
			return ok && data1.Cf62.Cmp(data2.Cf62) == 0 && data1.Cf63 == data2.Cf63 && data1.Cf64.Cmp(data2.Cf64) == 0 && data1.Cf65 == data2.Cf65 && data1.Cf66.Equals(data2.Cf66)
		}
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf57.Equals(data2.Cf57)
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

type D12_DC34 struct {
	Cf67 _dafny.MultiSet
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf67 _dafny.MultiSet) D12 {
	return D12{D12_DC34{Cf67}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

func (CompanionStruct_D12_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D12) Dtor_cf67() _dafny.MultiSet {
	return _this.Get_().(D12_DC34).Cf67
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf67) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
			return ok && data1.Cf67.Equals(data2.Cf67)
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

type D13_DC35 struct {
	Cf68 _dafny.Map
}

func (D13_DC35) isD13() {}

func (CompanionStruct_D13_) Create_DC35_(Cf68 _dafny.Map) D13 {
	return D13{D13_DC35{Cf68}}
}

func (_this D13) Is_DC35() bool {
	_, ok := _this.Get_().(D13_DC35)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D13) Dtor_cf68() _dafny.Map {
	return _this.Get_().(D13_DC35).Cf68
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC35:
		{
			return "D13.DC35" + "(" + _dafny.String(data.Cf68) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC35:
		{
			data2, ok := other.Get_().(D13_DC35)
			return ok && data1.Cf68.Equals(data2.Cf68)
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

type D14_DC37 struct {
	Cf70 bool
	Cf71 _dafny.Sequence
}

func (D14_DC37) isD14() {}

func (CompanionStruct_D14_) Create_DC37_(Cf70 bool, Cf71 _dafny.Sequence) D14 {
	return D14{D14_DC37{Cf70, Cf71}}
}

func (_this D14) Is_DC37() bool {
	_, ok := _this.Get_().(D14_DC37)
	return ok
}

type D14_DC36 struct {
	Cf69 _dafny.Map
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf69 _dafny.Map) D14 {
	return D14{D14_DC36{Cf69}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC37_(false, _dafny.EmptySeq)
}

func (_this D14) Dtor_cf70() bool {
	return _this.Get_().(D14_DC37).Cf70
}

func (_this D14) Dtor_cf71() _dafny.Sequence {
	return _this.Get_().(D14_DC37).Cf71
}

func (_this D14) Dtor_cf69() _dafny.Map {
	return _this.Get_().(D14_DC36).Cf69
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC37:
		{
			return "D14.DC37" + "(" + _dafny.String(data.Cf70) + ", " + data.Cf71.VerbatimString(true) + ")"
		}
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf69) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC37:
		{
			data2, ok := other.Get_().(D14_DC37)
			return ok && data1.Cf70 == data2.Cf70 && data1.Cf71.Equals(data2.Cf71)
		}
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && data1.Cf69.Equals(data2.Cf69)
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
	Cf72 _dafny.Array
}

func (D15_DC38) isD15() {}

func (CompanionStruct_D15_) Create_DC38_(Cf72 _dafny.Array) D15 {
	return D15{D15_DC38{Cf72}}
}

func (_this D15) Is_DC38() bool {
	_, ok := _this.Get_().(D15_DC38)
	return ok
}

func (CompanionStruct_D15_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D15) Dtor_cf72() _dafny.Array {
	return _this.Get_().(D15_DC38).Cf72
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC38:
		{
			return "D15.DC38" + "(" + _dafny.String(data.Cf72) + ")"
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
			return ok && data1.Cf72 == data2.Cf72
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
	Cf73 _dafny.Map
}

func (D16_DC39) isD16() {}

func (CompanionStruct_D16_) Create_DC39_(Cf73 _dafny.Map) D16 {
	return D16{D16_DC39{Cf73}}
}

func (_this D16) Is_DC39() bool {
	_, ok := _this.Get_().(D16_DC39)
	return ok
}

func (CompanionStruct_D16_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D16) Dtor_cf73() _dafny.Map {
	return _this.Get_().(D16_DC39).Cf73
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC39:
		{
			return "D16.DC39" + "(" + _dafny.String(data.Cf73) + ")"
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
			return ok && data1.Cf73.Equals(data2.Cf73)
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

type D17_DC41 struct {
	Cf75 bool
	Cf76 _dafny.CodePoint
}

func (D17_DC41) isD17() {}

func (CompanionStruct_D17_) Create_DC41_(Cf75 bool, Cf76 _dafny.CodePoint) D17 {
	return D17{D17_DC41{Cf75, Cf76}}
}

func (_this D17) Is_DC41() bool {
	_, ok := _this.Get_().(D17_DC41)
	return ok
}

type D17_DC40 struct {
	Cf74 *C3
}

func (D17_DC40) isD17() {}

func (CompanionStruct_D17_) Create_DC40_(Cf74 *C3) D17 {
	return D17{D17_DC40{Cf74}}
}

func (_this D17) Is_DC40() bool {
	_, ok := _this.Get_().(D17_DC40)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC41_(false, _dafny.CodePoint('D'))
}

func (_this D17) Dtor_cf75() bool {
	return _this.Get_().(D17_DC41).Cf75
}

func (_this D17) Dtor_cf76() _dafny.CodePoint {
	return _this.Get_().(D17_DC41).Cf76
}

func (_this D17) Dtor_cf74() *C3 {
	return _this.Get_().(D17_DC40).Cf74
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC41:
		{
			return "D17.DC41" + "(" + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ")"
		}
	case D17_DC40:
		{
			return "D17.DC40" + "(" + _dafny.String(data.Cf74) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC41:
		{
			data2, ok := other.Get_().(D17_DC41)
			return ok && data1.Cf75 == data2.Cf75 && data1.Cf76 == data2.Cf76
		}
	case D17_DC40:
		{
			data2, ok := other.Get_().(D17_DC40)
			return ok && data1.Cf74 == data2.Cf74
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

type D18_DC43 struct {
	Cf78 _dafny.Int
}

func (D18_DC43) isD18() {}

func (CompanionStruct_D18_) Create_DC43_(Cf78 _dafny.Int) D18 {
	return D18{D18_DC43{Cf78}}
}

func (_this D18) Is_DC43() bool {
	_, ok := _this.Get_().(D18_DC43)
	return ok
}

type D18_DC44 struct {
	Cf79 bool
	Cf80 bool
	Cf81 _dafny.MultiSet
	Cf82 *C4
}

func (D18_DC44) isD18() {}

func (CompanionStruct_D18_) Create_DC44_(Cf79 bool, Cf80 bool, Cf81 _dafny.MultiSet, Cf82 *C4) D18 {
	return D18{D18_DC44{Cf79, Cf80, Cf81, Cf82}}
}

func (_this D18) Is_DC44() bool {
	_, ok := _this.Get_().(D18_DC44)
	return ok
}

type D18_DC42 struct {
	Cf77 _dafny.Map
}

func (D18_DC42) isD18() {}

func (CompanionStruct_D18_) Create_DC42_(Cf77 _dafny.Map) D18 {
	return D18{D18_DC42{Cf77}}
}

func (_this D18) Is_DC42() bool {
	_, ok := _this.Get_().(D18_DC42)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC43_(_dafny.Zero)
}

func (_this D18) Dtor_cf78() _dafny.Int {
	return _this.Get_().(D18_DC43).Cf78
}

func (_this D18) Dtor_cf79() bool {
	return _this.Get_().(D18_DC44).Cf79
}

func (_this D18) Dtor_cf80() bool {
	return _this.Get_().(D18_DC44).Cf80
}

func (_this D18) Dtor_cf81() _dafny.MultiSet {
	return _this.Get_().(D18_DC44).Cf81
}

func (_this D18) Dtor_cf82() *C4 {
	return _this.Get_().(D18_DC44).Cf82
}

func (_this D18) Dtor_cf77() _dafny.Map {
	return _this.Get_().(D18_DC42).Cf77
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC43:
		{
			return "D18.DC43" + "(" + _dafny.String(data.Cf78) + ")"
		}
	case D18_DC44:
		{
			return "D18.DC44" + "(" + _dafny.String(data.Cf79) + ", " + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ", " + _dafny.String(data.Cf82) + ")"
		}
	case D18_DC42:
		{
			return "D18.DC42" + "(" + _dafny.String(data.Cf77) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC43:
		{
			data2, ok := other.Get_().(D18_DC43)
			return ok && data1.Cf78.Cmp(data2.Cf78) == 0
		}
	case D18_DC44:
		{
			data2, ok := other.Get_().(D18_DC44)
			return ok && data1.Cf79 == data2.Cf79 && data1.Cf80 == data2.Cf80 && data1.Cf81.Equals(data2.Cf81) && data1.Cf82 == data2.Cf82
		}
	case D18_DC42:
		{
			data2, ok := other.Get_().(D18_DC42)
			return ok && data1.Cf77.Equals(data2.Cf77)
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

type D19_DC45 struct {
	Cf83 T0
}

func (D19_DC45) isD19() {}

func (CompanionStruct_D19_) Create_DC45_(Cf83 T0) D19 {
	return D19{D19_DC45{Cf83}}
}

func (_this D19) Is_DC45() bool {
	_, ok := _this.Get_().(D19_DC45)
	return ok
}

func (CompanionStruct_D19_) Default() T0 {
	return (T0)(nil)
}

func (_this D19) Dtor_cf83() T0 {
	return _this.Get_().(D19_DC45).Cf83
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC45:
		{
			return "D19.DC45" + "(" + _dafny.String(data.Cf83) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC45:
		{
			data2, ok := other.Get_().(D19_DC45)
			return ok && _dafny.AreEqual(data1.Cf83, data2.Cf83)
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

type D20_DC47 struct {
	Cf85 bool
	Cf86 _dafny.Int
}

func (D20_DC47) isD20() {}

func (CompanionStruct_D20_) Create_DC47_(Cf85 bool, Cf86 _dafny.Int) D20 {
	return D20{D20_DC47{Cf85, Cf86}}
}

func (_this D20) Is_DC47() bool {
	_, ok := _this.Get_().(D20_DC47)
	return ok
}

type D20_DC48 struct {
	Cf87 _dafny.Map
}

func (D20_DC48) isD20() {}

func (CompanionStruct_D20_) Create_DC48_(Cf87 _dafny.Map) D20 {
	return D20{D20_DC48{Cf87}}
}

func (_this D20) Is_DC48() bool {
	_, ok := _this.Get_().(D20_DC48)
	return ok
}

type D20_DC46 struct {
	Cf84 _dafny.Array
}

func (D20_DC46) isD20() {}

func (CompanionStruct_D20_) Create_DC46_(Cf84 _dafny.Array) D20 {
	return D20{D20_DC46{Cf84}}
}

func (_this D20) Is_DC46() bool {
	_, ok := _this.Get_().(D20_DC46)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC47_(false, _dafny.Zero)
}

func (_this D20) Dtor_cf85() bool {
	return _this.Get_().(D20_DC47).Cf85
}

func (_this D20) Dtor_cf86() _dafny.Int {
	return _this.Get_().(D20_DC47).Cf86
}

func (_this D20) Dtor_cf87() _dafny.Map {
	return _this.Get_().(D20_DC48).Cf87
}

func (_this D20) Dtor_cf84() _dafny.Array {
	return _this.Get_().(D20_DC46).Cf84
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC47:
		{
			return "D20.DC47" + "(" + _dafny.String(data.Cf85) + ", " + _dafny.String(data.Cf86) + ")"
		}
	case D20_DC48:
		{
			return "D20.DC48" + "(" + _dafny.String(data.Cf87) + ")"
		}
	case D20_DC46:
		{
			return "D20.DC46" + "(" + _dafny.String(data.Cf84) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC47:
		{
			data2, ok := other.Get_().(D20_DC47)
			return ok && data1.Cf85 == data2.Cf85 && data1.Cf86.Cmp(data2.Cf86) == 0
		}
	case D20_DC48:
		{
			data2, ok := other.Get_().(D20_DC48)
			return ok && data1.Cf87.Equals(data2.Cf87)
		}
	case D20_DC46:
		{
			data2, ok := other.Get_().(D20_DC46)
			return ok && data1.Cf84 == data2.Cf84
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

type D21_DC50 struct {
	Cf89 bool
	Cf90 _dafny.Int
	Cf91 bool
	Cf92 bool
}

func (D21_DC50) isD21() {}

func (CompanionStruct_D21_) Create_DC50_(Cf89 bool, Cf90 _dafny.Int, Cf91 bool, Cf92 bool) D21 {
	return D21{D21_DC50{Cf89, Cf90, Cf91, Cf92}}
}

func (_this D21) Is_DC50() bool {
	_, ok := _this.Get_().(D21_DC50)
	return ok
}

type D21_DC49 struct {
	Cf88 _dafny.Sequence
}

func (D21_DC49) isD21() {}

func (CompanionStruct_D21_) Create_DC49_(Cf88 _dafny.Sequence) D21 {
	return D21{D21_DC49{Cf88}}
}

func (_this D21) Is_DC49() bool {
	_, ok := _this.Get_().(D21_DC49)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC50_(false, _dafny.Zero, false, false)
}

func (_this D21) Dtor_cf89() bool {
	return _this.Get_().(D21_DC50).Cf89
}

func (_this D21) Dtor_cf90() _dafny.Int {
	return _this.Get_().(D21_DC50).Cf90
}

func (_this D21) Dtor_cf91() bool {
	return _this.Get_().(D21_DC50).Cf91
}

func (_this D21) Dtor_cf92() bool {
	return _this.Get_().(D21_DC50).Cf92
}

func (_this D21) Dtor_cf88() _dafny.Sequence {
	return _this.Get_().(D21_DC49).Cf88
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC50:
		{
			return "D21.DC50" + "(" + _dafny.String(data.Cf89) + ", " + _dafny.String(data.Cf90) + ", " + _dafny.String(data.Cf91) + ", " + _dafny.String(data.Cf92) + ")"
		}
	case D21_DC49:
		{
			return "D21.DC49" + "(" + _dafny.String(data.Cf88) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC50:
		{
			data2, ok := other.Get_().(D21_DC50)
			return ok && data1.Cf89 == data2.Cf89 && data1.Cf90.Cmp(data2.Cf90) == 0 && data1.Cf91 == data2.Cf91 && data1.Cf92 == data2.Cf92
		}
	case D21_DC49:
		{
			data2, ok := other.Get_().(D21_DC49)
			return ok && data1.Cf88.Equals(data2.Cf88)
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
	Cf94 _dafny.Map
}

func (D22_DC52) isD22() {}

func (CompanionStruct_D22_) Create_DC52_(Cf94 _dafny.Map) D22 {
	return D22{D22_DC52{Cf94}}
}

func (_this D22) Is_DC52() bool {
	_, ok := _this.Get_().(D22_DC52)
	return ok
}

type D22_DC53 struct {
	Cf95 bool
	Cf96 _dafny.CodePoint
	Cf97 bool
	Cf98 bool
}

func (D22_DC53) isD22() {}

func (CompanionStruct_D22_) Create_DC53_(Cf95 bool, Cf96 _dafny.CodePoint, Cf97 bool, Cf98 bool) D22 {
	return D22{D22_DC53{Cf95, Cf96, Cf97, Cf98}}
}

func (_this D22) Is_DC53() bool {
	_, ok := _this.Get_().(D22_DC53)
	return ok
}

type D22_DC51 struct {
	Cf93 _dafny.Map
}

func (D22_DC51) isD22() {}

func (CompanionStruct_D22_) Create_DC51_(Cf93 _dafny.Map) D22 {
	return D22{D22_DC51{Cf93}}
}

func (_this D22) Is_DC51() bool {
	_, ok := _this.Get_().(D22_DC51)
	return ok
}

type D22_DC54 struct {
	Cf99 D22
}

func (D22_DC54) isD22() {}

func (CompanionStruct_D22_) Create_DC54_(Cf99 D22) D22 {
	return D22{D22_DC54{Cf99}}
}

func (_this D22) Is_DC54() bool {
	_, ok := _this.Get_().(D22_DC54)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC52_(_dafny.EmptyMap)
}

func (_this D22) Dtor_cf94() _dafny.Map {
	return _this.Get_().(D22_DC52).Cf94
}

func (_this D22) Dtor_cf95() bool {
	return _this.Get_().(D22_DC53).Cf95
}

func (_this D22) Dtor_cf96() _dafny.CodePoint {
	return _this.Get_().(D22_DC53).Cf96
}

func (_this D22) Dtor_cf97() bool {
	return _this.Get_().(D22_DC53).Cf97
}

func (_this D22) Dtor_cf98() bool {
	return _this.Get_().(D22_DC53).Cf98
}

func (_this D22) Dtor_cf93() _dafny.Map {
	return _this.Get_().(D22_DC51).Cf93
}

func (_this D22) Dtor_cf99() D22 {
	return _this.Get_().(D22_DC54).Cf99
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC52:
		{
			return "D22.DC52" + "(" + _dafny.String(data.Cf94) + ")"
		}
	case D22_DC53:
		{
			return "D22.DC53" + "(" + _dafny.String(data.Cf95) + ", " + _dafny.String(data.Cf96) + ", " + _dafny.String(data.Cf97) + ", " + _dafny.String(data.Cf98) + ")"
		}
	case D22_DC51:
		{
			return "D22.DC51" + "(" + _dafny.String(data.Cf93) + ")"
		}
	case D22_DC54:
		{
			return "D22.DC54" + "(" + _dafny.String(data.Cf99) + ")"
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
			data2, ok := other.Get_().(D22_DC52)
			return ok && data1.Cf94.Equals(data2.Cf94)
		}
	case D22_DC53:
		{
			data2, ok := other.Get_().(D22_DC53)
			return ok && data1.Cf95 == data2.Cf95 && data1.Cf96 == data2.Cf96 && data1.Cf97 == data2.Cf97 && data1.Cf98 == data2.Cf98
		}
	case D22_DC51:
		{
			data2, ok := other.Get_().(D22_DC51)
			return ok && data1.Cf93.Equals(data2.Cf93)
		}
	case D22_DC54:
		{
			data2, ok := other.Get_().(D22_DC54)
			return ok && data1.Cf99.Equals(data2.Cf99)
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm5(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence
	Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map
	F6() _dafny.Int
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
	Fm5(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence
	Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map
	F6() _dafny.Int
	F9() _dafny.Array
	F9_set_(value _dafny.Array)
	Fm13(p0 _dafny.CodePoint, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map
	M6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool)
	M7(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) (_dafny.CodePoint, _dafny.Int, _dafny.Array, _dafny.Int)
	F10() _dafny.Map
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

// Definition of trait T2
type T2 interface {
	String() string
	Fm5(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence
	Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map
	F6() _dafny.Int
	F16() _dafny.MultiSet
	F16_set_(value _dafny.MultiSet)
	Fm21(p0 _dafny.Int, globalState *GlobalState) _dafny.Int
	Fm22(p0 _dafny.MultiSet, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool
}
type CompanionStruct_T2_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T2_ = CompanionStruct_T2_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T2_) CastTo_(x interface{}) T2 {
	var t T2
	t, _ = x.(T2)
	return t
}

// End of trait T2

// Definition of class GlobalState
type GlobalState struct {
	F0  bool
	F2  _dafny.Sequence
	F5  _dafny.Int
	_f1 _dafny.Sequence
	_f3 bool
	_f4 _dafny.Map
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this.F2 = _dafny.EmptySeq
	_this.F5 = _dafny.Zero
	_this._f1 = _dafny.EmptySeq
	_this._f3 = false
	_this._f4 = _dafny.EmptyMap
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Sequence, f2 _dafny.Sequence, f3 bool, f4 _dafny.Map, f5 _dafny.Int) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
	}
}
func (_this *GlobalState) F1() _dafny.Sequence {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Map {
	{
		return _this._f4
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F19 _dafny.Array
	F20 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F19 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F20 = false
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

func (_this *C0) Ctor__(f19 _dafny.Array, f20 bool) {
	{
		(_this).F19 = f19
		(_this).F20 = f20
	}
}
func (_this *C0) Fm23(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(638)
	}
}
func (_this *C0) Fm24(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqOf(!(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F20, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false)).Cardinality(), _this.F20)).Cardinality(), _this.F20)).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-605))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}(func(_235_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		})))).Cardinality())).Cardinality()).Cmp(_dafny.IntOfInt64(779)) < 0), _this.F20, _this.F20, _this.F20)
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f6 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f6 = _dafny.Zero
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

func (_this *C1) F6() _dafny.Int {
	return _this._f6
}
func (_this *C1) Ctor__(f6 _dafny.Int) {
	{
		(_this)._f6 = f6
	}
}
func (_this *C1) Fm5(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		var _source4 D3 = (func() D3 {
			if false {
				return Companion_D3_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), false), false))
			}
			return Companion_D3_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(214), false), false))
		})()
		_ = _source4
		if _source4.Is_DC10() {
			var _236___mcc_h0 _dafny.Map = _source4.Get_().(D3_DC10).Cf22
			_ = _236___mcc_h0
			var _237_cf22 _dafny.Map = _236___mcc_h0
			_ = _237_cf22
			return _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('h'), _dafny.CodePoint('f'), _dafny.CodePoint('r')))
		} else if _source4.Is_DC9() {
			var _238___mcc_h1 _dafny.Sequence = _source4.Get_().(D3_DC9).Cf21
			_ = _238___mcc_h1
			var _239_cf21 _dafny.Sequence = _238___mcc_h1
			_ = _239_cf21
			return (Companion_D4_.Create_DC12_(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('g')), _dafny.MultiSetOf(_dafny.CodePoint('m'), _dafny.CodePoint('k'), _dafny.CodePoint('s'), _dafny.CodePoint('t'))))).Dtor_cf24()
		} else {
			var _240___mcc_h2 D3 = _source4.Get_().(D3_DC11).Cf23
			_ = _240___mcc_h2
			var _241_cf23 D3 = _240___mcc_h2
			_ = _241_cf23
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('q'))), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('k'), _dafny.CodePoint('i'))))
		}
	}
}
func (_this *C1) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter9 := _dafny.Iterate(((_dafny.MultiSetOf((_this).F6(), (_this).F6(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pxn")).Cardinality()), (_this).F6())).Union(_dafny.MultiSetOf((_this).F6()))).Elements()); ; {
				_compr_8, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _242_v0 _dafny.Int
				_242_v0 = interface{}(_compr_8).(_dafny.Int)
				if ((_dafny.MultiSetOf((_this).F6(), (_this).F6(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pxn")).Cardinality()), (_this).F6())).Union(_dafny.MultiSetOf((_this).F6()))).Contains(_242_v0) {
					_coll8.Add((_242_v0).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(272), (_this).F6())).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(256))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg32 _dafny.Int) interface{} {
							return coer32(arg32)
						}
					}(func(_243_i0 _dafny.Int) _dafny.Int {
						return (_this).F6()
					})), _dafny.SeqOf((_this).F6())))
				}
			}
			return _coll8.ToMap()
		}()
	}
}
func (_this *C1) Fm28(globalState *GlobalState) bool {
	{
		return ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F6()), _dafny.SeqOf((_this).F6(), (_this).F6()))).Cardinality()))).Cmp(_dafny.IntOfInt64(618)) <= 0
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f16 _dafny.MultiSet
	_f9  _dafny.Array
	_f6  _dafny.Int
	_f10 _dafny.Map
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f16 = _dafny.EmptyMultiSet
	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.Zero
	_this._f10 = _dafny.EmptyMap
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C2{}
var _ T2 = &C2{}
var _ T1 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F16() _dafny.MultiSet {
	return _this._f16
}
func (_this *C2) F16_set_(value _dafny.MultiSet) {
	_this._f16 = value
}
func (_this *C2) F9() _dafny.Array {
	return _this._f9
}
func (_this *C2) F9_set_(value _dafny.Array) {
	_this._f9 = value
}
func (_this *C2) F6() _dafny.Int {
	return _this._f6
}
func (_this *C2) F10() _dafny.Map {
	return _this._f10
}
func (_this *C2) Ctor__(f6 _dafny.Int, f16 _dafny.MultiSet, f9 _dafny.Array, f10 _dafny.Map) {
	{
		(_this)._f6 = f6
		(_this)._f16 = f16
		(_this)._f9 = f9
		(_this)._f10 = f10
	}
}
func (_this *C2) Fm5(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		var _source5 D2 = Companion_D2_.Create_DC8_(!(!(false)), false)
		_ = _source5
		if _source5.Is_DC6() {
			var _244___mcc_h0 _dafny.Int = _source5.Get_().(D2_DC6).Cf13
			_ = _244___mcc_h0
			var _245___mcc_h1 _dafny.Array = _source5.Get_().(D2_DC6).Cf14
			_ = _245___mcc_h1
			var _246_cf14 _dafny.Array = _245___mcc_h1
			_ = _246_cf14
			var _247_cf13 _dafny.Int = _244___mcc_h0
			_ = _247_cf13
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('s'), _dafny.CodePoint('h')))), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('d'))))
		} else if _source5.Is_DC7() {
			var _248___mcc_h2 _dafny.Sequence = _source5.Get_().(D2_DC7).Cf15
			_ = _248___mcc_h2
			var _249___mcc_h3 _dafny.Array = _source5.Get_().(D2_DC7).Cf16
			_ = _249___mcc_h3
			var _250___mcc_h4 _dafny.Int = _source5.Get_().(D2_DC7).Cf17
			_ = _250___mcc_h4
			var _251___mcc_h5 _dafny.Int = _source5.Get_().(D2_DC7).Cf18
			_ = _251___mcc_h5
			var _252_cf18 _dafny.Int = _251___mcc_h5
			_ = _252_cf18
			var _253_cf17 _dafny.Int = _250___mcc_h4
			_ = _253_cf17
			var _254_cf16 _dafny.Array = _249___mcc_h3
			_ = _254_cf16
			var _255_cf15 _dafny.Sequence = _248___mcc_h2
			_ = _255_cf15
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('h')), _dafny.MultiSetOf(_dafny.CodePoint('r')), _dafny.MultiSetOf(_dafny.CodePoint('e'), _dafny.CodePoint('u'))), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('h')), _dafny.MultiSetOf(_dafny.CodePoint('y')), _dafny.MultiSetOf(_dafny.CodePoint('s'))))
		} else if _source5.Is_DC8() {
			var _256___mcc_h6 bool = _source5.Get_().(D2_DC8).Cf19
			_ = _256___mcc_h6
			var _257___mcc_h7 bool = _source5.Get_().(D2_DC8).Cf20
			_ = _257___mcc_h7
			var _258_cf20 bool = _257___mcc_h7
			_ = _258_cf20
			var _259_cf19 bool = _256___mcc_h6
			_ = _259_cf19
			return _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('m'), _dafny.CodePoint('d'), _dafny.CodePoint('i'), _dafny.CodePoint('k'), _dafny.CodePoint('o')), _dafny.MultiSetOf(_dafny.CodePoint('h')))
		} else {
			var _260___mcc_h8 _dafny.Map = _source5.Get_().(D2_DC5).Cf12
			_ = _260___mcc_h8
			var _261_cf12 _dafny.Map = _260___mcc_h8
			_ = _261_cf12
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('q'))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(523))).Uint32(), func(coer33 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}(func(_262_i0 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(_dafny.CodePoint('u'))
			})))
		}
	}
}
func (_this *C2) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.SeqOf((_this).F6()))
	}
}
func (_this *C2) Fm21(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F6()).Minus((_this).F6())
	}
}
func (_this *C2) Fm22(p0 _dafny.MultiSet, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('m'), _dafny.CodePoint('d')), _dafny.CodePoint('d')) {
			return true
		} else {
			return ((_dafny.Zero).Minus((_this).F6())).Cmp(_dafny.IntOfInt64(248)) <= 0
		}
	}
}
func (_this *C2) Fm13(p0 _dafny.CodePoint, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(!(!(!(true)))))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-918))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}(func(_263_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		})))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(2))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}(func(_264_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('i')
		})))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("rxhrrtyr"))))
	}
}
func (_this *C2) M6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _265_v0 _dafny.Map
		_ = _265_v0
		_265_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), !(p0))
		var _266_v1 _dafny.Set
		_ = _266_v1
		_266_v1 = _dafny.SetOf((_265_v0).Merge(_265_v0))
		var _267_v2 _dafny.Map
		_ = _267_v2
		_267_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(91), (_this).F6())
		var _268_v3 D2
		_ = _268_v3
		_268_v3 = Companion_D2_.Create_DC5_(_267_v2)
		var _269_v4 _dafny.Set
		_ = _269_v4
		_269_v4 = _dafny.SetOf(p0)
		var _270_v5 _dafny.MultiSet
		_ = _270_v5
		_270_v5 = _dafny.MultiSetOf(p0)
		var _271_v6 _dafny.MultiSet
		_ = _271_v6
		_271_v6 = _dafny.MultiSetOf((_this).F6(), (_this).F6())
		var _272_v7 _dafny.Map
		_ = _272_v7
		_272_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_271_v6).Update((_this).F6(), Companion_Default___.Abs(p2)))
		var _rhs26 bool = (((_269_v4).Cardinality()).Minus(_dafny.IntOfInt64(648))).Cmp((_this).F6()) >= 0
		_ = _rhs26
		var _rhs27 _dafny.Set = ((func() _dafny.Set {
			if (_this).Fm22((_270_v5).Update(p0, Companion_Default___.Abs((_this).F6())), p0, (_this).F6(), globalState) {
				return _266_v1
			}
			return _266_v1
		})()).Difference((func() _dafny.Set {
			if p0 {
				return _266_v1
			}
			return _266_v1
		})())
		_ = _rhs27
		var _rhs28 D2 = Companion_D2_.Create_DC5_((_267_v2).Merge(_267_v2))
		_ = _rhs28
		var _rhs29 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F6(), ((func() _dafny.MultiSet {
			if (_272_v7).Contains(p0) {
				return (_272_v7).Get(p0).(_dafny.MultiSet)
			}
			return _271_v6
		})()).Cardinality())
		_ = _rhs29
		var _rhs30 bool = p0
		_ = _rhs30
		var _lhs19 *GlobalState = globalState
		_ = _lhs19
		var _lhs20 *GlobalState = globalState
		_ = _lhs20
		_lhs19.F0 = _rhs26
		_266_v1 = _rhs27
		_268_v3 = _rhs28
		_lhs20.F5 = _rhs29
		r1 = _rhs30
		var _273_v8 _dafny.Array
		_ = _273_v8
		var _nwElement0_5 _dafny.Int = (_this).F6()
		_ = _nwElement0_5
		var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(3))
		_ = _nw34
		(_nw34).ArraySet1(_nwElement0_5, 0)
		(_nw34).ArraySet1((_this).F6(), 1)
		(_nw34).ArraySet1((_this).F6(), 2)
		_273_v8 = _nw34
		var _274_v9 _dafny.Set
		_ = _274_v9
		_274_v9 = _dafny.SetOf((_this).F6())
		var _275_v10 _dafny.Sequence
		_ = _275_v10
		_275_v10 = _dafny.SeqOf(_274_v9, _274_v9, _274_v9, _274_v9)
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_273_v8), 0))
		_ = _index27
		(_273_v8).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_275_v10, (Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_dafny.SeqOf(p2))).Cardinality(), _dafny.IntOfUint32((_275_v10).Cardinality()))).Uint32(), _274_v9), _275_v10)).Cardinality()), (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_273_v8), 0))
		_ = _index28
		var _rhs31 bool = ((func() _dafny.Int {
			if (_this).Fm22(_270_v5, p0, (_this).F6(), globalState) {
				return p2
			}
			return (_this).F6()
		})()).Cmp(Companion_Default___.SafeModuloInt((_this).F6(), (_this).F6())) >= 0
		_ = _rhs31
		var _rhs32 bool = true
		_ = _rhs32
		var _rhs33 _dafny.Int = (func() _dafny.Int {
			if (_271_v6).IsDisjointFrom(_271_v6) {
				return (_269_v4).Cardinality()
			}
			return ((_this).F6()).Plus((_this).F6())
		})()
		_ = _rhs33
		var _rhs34 bool = p0
		_ = _rhs34
		var _lhs21 *GlobalState = globalState
		_ = _lhs21
		var _lhs22 *GlobalState = globalState
		_ = _lhs22
		var _lhs23 _dafny.Array = _273_v8
		_ = _lhs23
		var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_273_v8), 0))
		_ = _lhs24
		var _lhs25 *GlobalState = globalState
		_ = _lhs25
		_lhs21.F0 = _rhs31
		_lhs22.F0 = _rhs32
		(_lhs23).ArraySet1(_rhs33, (_lhs24).Int())
		_lhs25.F0 = _rhs34
		var _276_v11 *C1
		_ = _276_v11
		var _nw35 *C1 = New_C1_()
		_ = _nw35
		_nw35.Ctor__(_dafny.IntOfInt64(-380))
		_276_v11 = _nw35
		(globalState).F0 = !((p2).Cmp((_273_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_273_v8), 0))).Int()).(_dafny.Int)) >= 0) || (p0)
		var _277_v12 _dafny.Array
		_ = _277_v12
		var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(24))
		_ = _nw36
		_277_v12 = _nw36
		var _278_v13 D2
		_ = _278_v13
		_278_v13 = Companion_D2_.Create_DC7_(p1, _277_v12, (_this).F6(), (_269_v4).Cardinality())
		var _279_v14 _dafny.Map
		_ = _279_v14
		_279_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_278_v13, _dafny.UnicodeSeqOfUtf8Bytes("gtvimpmrf"))
		var _280_v15 _dafny.Sequence
		_ = _280_v15
		_280_v15 = _dafny.UnicodeSeqOfUtf8Bytes("ohcxxhxh")
		var _281_v16 _dafny.CodePoint
		_ = _281_v16
		_281_v16 = _dafny.CodePoint('t')
		(globalState).F5 = ((_273_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_273_v8), 0))).Int()).(_dafny.Int)).Times((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_279_v14).Contains(_278_v13) {
				return (_279_v14).Get(_278_v13).(_dafny.Sequence)
			}
			return _280_v15
		})(), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_279_v14).Contains(_278_v13) {
				return (_279_v14).Get(_278_v13).(_dafny.Sequence)
			}
			return _280_v15
		})()).Cardinality()))).Uint32(), _281_v16)).Cardinality())).Plus(_dafny.IntOfUint32((_280_v15).Cardinality())))
		var _arr2 _dafny.Array = _this.F9()
		_ = _arr2
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_this.F9()), 0))
		_ = _index29
		(_arr2).ArraySet1(p0, (_index29).Int())
		var _arr3 _dafny.Array = _this.F9()
		_ = _arr3
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_this.F9()), 0))
		_ = _index30
		(_arr3).ArraySet1(!(p0), (_index30).Int())
		r0 = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm29(_274_v9, globalState), (Companion_Default___.SafeIndex((func() _dafny.Int {
			if p0 {
				return (_this).F6()
			}
			return (_273_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_273_v8), 0))).Int()).(_dafny.Int)
		})(), _dafny.IntOfUint32((Companion_Default___.Fm29(_274_v9, globalState)).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Concatenate(_280_v15, _280_v15))
		var _282_v17 _dafny.Sequence
		_ = _282_v17
		_282_v17 = _dafny.SeqOf((_267_v2).Cardinality())
		r1 = (p2).Cmp((_dafny.IntOfUint32((_282_v17).Cardinality())).Times((_273_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_273_v8), 0))).Int()).(_dafny.Int))) >= 0
		return r0, r1
	}
}
func (_this *C2) M7(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) (_dafny.CodePoint, _dafny.Int, _dafny.Array, _dafny.Int) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _283_v0 bool
		_ = _283_v0
		_283_v0 = true
		if !(_283_v0) || (!(!(_283_v0))) {
			var _284_v1 _dafny.Set
			_ = _284_v1
			_284_v1 = _dafny.SetOf(_283_v0)
			var _285_v2 D5
			_ = _285_v2
			_285_v2 = Companion_D5_.Create_DC16_(_284_v1)
			var _286_v3 _dafny.Sequence
			_ = _286_v3
			_286_v3 = _dafny.SeqOf((_285_v2).Dtor_cf35(), (_284_v1).Difference(_284_v1), (_284_v1).Difference(_284_v1))
			var _287_v4 _dafny.Array
			_ = _287_v4
			var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(12))
			_ = _nw37
			_287_v4 = _nw37
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_287_v4), 0))
			_ = _index31
			(_287_v4).ArraySet1((_dafny.SetOf(_283_v0)).Difference((_285_v2).Dtor_cf35()), (_index31).Int())
			var _288_v5 _dafny.CodePoint
			_ = _288_v5
			_288_v5 = _dafny.CodePoint('v')
			var _289_v6 _dafny.Map
			_ = _289_v6
			_289_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_283_v0), (_this).F6())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_287_v4), 0))
			_ = _index32
			var _rhs35 _dafny.CodePoint = _288_v5
			_ = _rhs35
			var _rhs36 _dafny.Sequence = (func() _dafny.Sequence {
				if _283_v0 {
					return _286_v3
				}
				return _286_v3
			})()
			_ = _rhs36
			var _rhs37 _dafny.Int = (_this).Fm21(((_289_v6).Merge(_289_v6)).Cardinality(), globalState)
			_ = _rhs37
			var _rhs38 _dafny.Int = p1
			_ = _rhs38
			var _rhs39 _dafny.Set = ((_284_v1).Intersection(_284_v1)).Intersection(_dafny.SetOf(_283_v0, true, _283_v0))
			_ = _rhs39
			var _lhs26 *GlobalState = globalState
			_ = _lhs26
			var _lhs27 _dafny.Array = _287_v4
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_287_v4), 0))
			_ = _lhs28
			r0 = _rhs35
			_286_v3 = _rhs36
			_lhs26.F5 = _rhs37
			r3 = _rhs38
			(_lhs27).ArraySet1(_rhs39, (_lhs28).Int())
			var _290_v7 _dafny.Map
			_ = _290_v7
			_290_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v0, !(_283_v0))
			var _291_v8 D4
			_ = _291_v8
			_291_v8 = Companion_D4_.Create_DC14_(_283_v0, _290_v7, (_this).F6(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v0, _283_v0))
			var _292_v9 _dafny.Sequence
			_ = _292_v9
			_292_v9 = _dafny.SeqOf(false, _283_v0)
			var _293_v10 _dafny.Map
			_ = _293_v10
			_293_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_291_v8, _292_v9)
			_293_v10 = (_293_v10).Update(_291_v8, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(true), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Uint32(), !((_292_v9).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_292_v9).Cardinality()))).Uint32()).(bool))))
			r3 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v0, _283_v0)).Merge((_290_v7).Merge(_290_v7))).Cardinality()
			var _294_v11 _dafny.Sequence
			_ = _294_v11
			_294_v11 = _dafny.UnicodeSeqOfUtf8Bytes("cidjjdvpa")
			(globalState).F2 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ik"), _dafny.Companion_Sequence_.Concatenate(_294_v11, _294_v11)), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_this).F6(), p1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ik"), _dafny.Companion_Sequence_.Concatenate(_294_v11, _294_v11))).Cardinality()))).Uint32(), _288_v5)
			(globalState).F0 = Companion_Default___.Fm3(globalState)
		} else {
			(globalState).F0 = _283_v0
			if !(_283_v0) {
				(globalState).F0 = _283_v0
				var _295_v12 _dafny.Array
				_ = _295_v12
				var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
				_ = _nw38
				_295_v12 = _nw38
				var _296_v13 _dafny.Array
				_ = _296_v13
				var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
				_ = _nw39
				_296_v13 = _nw39
				var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_295_v12), 0))
				_ = _index33
				(_295_v12).ArraySet1(_296_v13, (_index33).Int())
				var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_295_v12), 0))
				_ = _index34
				(_295_v12).ArraySet1(_296_v13, (_index34).Int())
				var _297_v14 _dafny.Sequence
				_ = _297_v14
				_297_v14 = _dafny.UnicodeSeqOfUtf8Bytes("rdc")
				_283_v0 = _dafny.Companion_Sequence_.Equal(_297_v14, _297_v14)
				var _298_v15 _dafny.Array
				_ = _298_v15
				var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
				_ = _nw40
				_298_v15 = _nw40
				var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_298_v15), 0))
				_ = _index35
				(_298_v15).ArraySet1(_297_v14, (_index35).Int())
				var _299_v16 _dafny.Array
				_ = _299_v16
				var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw41
				_299_v16 = _nw41
				var _300_v17 _dafny.Sequence
				_ = _300_v17
				_300_v17 = _dafny.SeqOf(p1)
				var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_299_v16), 0))
				_ = _index36
				(_299_v16).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if _283_v0 {
						return p1
					}
					return _dafny.IntOfUint32((_300_v17).Cardinality())
				})()), (_index36).Int())
				var _301_v18 _dafny.Sequence
				_ = _301_v18
				_301_v18 = _dafny.SeqOf(true)
				var _302_v19 D4
				_ = _302_v19
				_302_v19 = Companion_D4_.Create_DC13_(_dafny.SeqOf(_dafny.IntOfInt64(-406)), _283_v0, p1)
				var _303_v20 _dafny.Map
				_ = _303_v20
				_303_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_302_v19, _dafny.IntOfUint32((_300_v17).Cardinality()))
				var _304_v21 _dafny.Map
				_ = _304_v21
				_304_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_303_v20).Contains(_302_v19) {
						return (_303_v20).Get(_302_v19).(_dafny.Int)
					}
					return p1
				})(), (_this).F6())
				var _305_v22 _dafny.CodePoint
				_ = _305_v22
				_305_v22 = _dafny.CodePoint('c')
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_298_v15), 0))
				_ = _index37
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_299_v16), 0))
				_ = _index38
				var _rhs40 _dafny.Int = _dafny.IntOfUint32((_300_v17).Cardinality())
				_ = _rhs40
				var _rhs41 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F6(), Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfUint32((_301_v18).Cardinality())))
				_ = _rhs41
				var _rhs42 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_297_v14, (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_304_v21).Contains((_dafny.Zero).Minus((_304_v21).Cardinality())) {
						return (_304_v21).Get((_dafny.Zero).Minus((_304_v21).Cardinality())).(_dafny.Int)
					}
					return p1
				})(), _dafny.IntOfUint32((_297_v14).Cardinality()))).Uint32(), _305_v22)
				_ = _rhs42
				var _rhs43 _dafny.Int = (p1).Times(p1)
				_ = _rhs43
				var _rhs44 _dafny.Int = p1
				_ = _rhs44
				var _lhs29 *GlobalState = globalState
				_ = _lhs29
				var _lhs30 *GlobalState = globalState
				_ = _lhs30
				var _lhs31 _dafny.Array = _298_v15
				_ = _lhs31
				var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_298_v15), 0))
				_ = _lhs32
				var _lhs33 _dafny.Array = _299_v16
				_ = _lhs33
				var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_299_v16), 0))
				_ = _lhs34
				var _lhs35 *GlobalState = globalState
				_ = _lhs35
				_lhs29.F5 = _rhs40
				_lhs30.F5 = _rhs41
				(_lhs31).ArraySet1(_rhs42, (_lhs32).Int())
				(_lhs33).ArraySet1(_rhs43, (_lhs34).Int())
				_lhs35.F5 = _rhs44
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_299_v16), 0))
				_ = _index39
				(_299_v16).ArraySet1(p1, (_index39).Int())
			} else {
				var _306_v23 _dafny.CodePoint
				_ = _306_v23
				_306_v23 = _dafny.CodePoint('j')
				var _307_v24 _dafny.Sequence
				_ = _307_v24
				_307_v24 = _dafny.UnicodeSeqOfUtf8Bytes("tgruuabrh")
				var _308_v25 _dafny.Sequence
				_ = _308_v25
				_308_v25 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("yjxblq"), _dafny.UnicodeSeqOfUtf8Bytes("kfxtlw")), (func() _dafny.Sequence {
					if _283_v0 {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(293))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg36 _dafny.Int) interface{} {
								return coer36(arg36)
							}
						}(func(_309_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('u')
						}))
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("yiaoosnib")
				})(), Companion_Default___.Fm30(true, _306_v23, globalState), _307_v24)
				(globalState).F2 = (_308_v25).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_308_v25).Cardinality()))).Uint32()).(_dafny.Sequence)
				(globalState).F0 = !(!(_283_v0))
				_283_v0 = ((_this).F6()).Cmp((_this).F6()) != 0
				(globalState).F2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_307_v24, _307_v24), _307_v24)
				(globalState).F0 = _283_v0
			}
			var _310_v26 _dafny.Map
			_ = _310_v26
			_310_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v0, (_this).F6())
			var _311_v27 _dafny.Sequence
			_ = _311_v27
			_311_v27 = _dafny.SeqOf((_310_v26).Merge(_310_v26))
			var _rhs45 _dafny.Map = (_310_v26).Merge(Companion_Default___.Fm31(_283_v0, (_this).F6(), Companion_Default___.Fm2(p1, globalState), _dafny.SeqOf(_283_v0), globalState))
			_ = _rhs45
			var _rhs46 _dafny.Sequence = _dafny.SeqOf(_310_v26, (_310_v26).Merge(_310_v26), _310_v26, _310_v26)
			_ = _rhs46
			_310_v26 = _rhs45
			_311_v27 = _rhs46
			var _312_v28 _dafny.Sequence
			_ = _312_v28
			_312_v28 = _dafny.SeqOf(_283_v0)
			var _313_v29 _dafny.Map
			_ = _313_v29
			_313_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_312_v28, _312_v28), _283_v0)
			_313_v29 = _313_v29
			var _314_v30 _dafny.Map
			_ = _314_v30
			_314_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v0, _283_v0)
			var _315_v31 _dafny.CodePoint
			_ = _315_v31
			_315_v31 = _dafny.CodePoint('i')
			var _316_v32 _dafny.MultiSet
			_ = _316_v32
			_316_v32 = _dafny.MultiSetOf(_315_v31, _315_v31, _315_v31, _315_v31, _dafny.CodePoint('m'))
			var _317_v33 _dafny.MultiSet
			_ = _317_v33
			_317_v33 = _dafny.MultiSetOf(!(_283_v0), !(_283_v0))
			var _318_v34 D4
			_ = _318_v34
			_318_v34 = Companion_D4_.Create_DC14_(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v0, true), _dafny.IntOfInt64(-428), _314_v30)
			var _319_v35 _dafny.Set
			_ = _319_v35
			_319_v35 = _dafny.SetOf(true, !(_283_v0), (_318_v34).Dtor_cf28(), _283_v0, true)
			var _320_v36 _dafny.Array
			_ = _320_v36
			var _nwElement0_6 _dafny.Int = (_314_v30).Cardinality()
			_ = _nwElement0_6
			var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(14))
			_ = _nw42
			(_nw42).ArraySet1(_nwElement0_6, 0)
			(_nw42).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_315_v31, _283_v0)).Cardinality(), 1)
			(_nw42).ArraySet1(((_this).F6()).Minus(p1), 2)
			(_nw42).ArraySet1((func() _dafny.Int {
				if (_316_v32).Contains(_315_v31) {
					return (_316_v32).Multiplicity(_315_v31)
				}
				return p1
			})(), 3)
			(_nw42).ArraySet1(_dafny.IntOfUint32((_312_v28).Cardinality()), 4)
			(_nw42).ArraySet1(Companion_Default___.Fm2((_this).F6(), globalState), 5)
			(_nw42).ArraySet1((_dafny.Zero).Minus(p1), 6)
			(_nw42).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _283_v0)).Cardinality(), 7)
			(_nw42).ArraySet1((Companion_Default___.Fm2(p1, globalState)).Minus((_317_v33).Cardinality()), 8)
			(_nw42).ArraySet1(p1, 9)
			(_nw42).ArraySet1((_this).F6(), 10)
			(_nw42).ArraySet1((_this).Fm21(p1, globalState), 11)
			(_nw42).ArraySet1(p1, 12)
			(_nw42).ArraySet1((func() _dafny.Int {
				if _283_v0 {
					return (_319_v35).Cardinality()
				}
				return (_this).F6()
			})(), 13)
			_320_v36 = _nw42
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_320_v36), 0))
			_ = _index40
			(_320_v36).ArraySet1(_dafny.IntOfInt64(-446), (_index40).Int())
			var _321_v37 D0
			_ = _321_v37
			_321_v37 = Companion_D0_.Create_DC0_(_312_v28, _dafny.IntOfInt64(593))
			var _322_v38 D1
			_ = _322_v38
			_322_v38 = Companion_D1_.Create_DC4_(_321_v37, (_this).Fm21(p1, globalState), _dafny.UnicodeSeqOfUtf8Bytes("rdeikwgw"), _283_v0, p1)
			var _323_v39 _dafny.Sequence
			_ = _323_v39
			_323_v39 = _dafny.UnicodeSeqOfUtf8Bytes("s")
			var _324_v40 _dafny.Map
			_ = _324_v40
			_324_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v0, (_323_v39).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_323_v39).Cardinality()), _dafny.IntOfUint32((_323_v39).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _pat_let_tv3 = _324_v40
			_ = _pat_let_tv3
			var _pat_let_tv4 = _323_v39
			_ = _pat_let_tv4
			var _325_v42 D1
			_ = _325_v42
			_325_v42 = Companion_D1_.Create_DC4_(_321_v37, (_this).F6(), (func(_pat_let6_0 D1) D1 {
				return func(_326_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let7_0 _dafny.Int) D1 {
						return func(_327_dt__update_hcf8_h0 _dafny.Int) D1 {
							return func(_pat_let8_0 _dafny.Sequence) D1 {
								return func(_328_dt__update_hcf9_h0 _dafny.Sequence) D1 {
									return Companion_D1_.Create_DC4_((_326_dt__update__tmp_h0).Dtor_cf7(), _327_dt__update_hcf8_h0, _328_dt__update_hcf9_h0, (_326_dt__update__tmp_h0).Dtor_cf10(), (_326_dt__update__tmp_h0).Dtor_cf11())
								}(_pat_let8_0)
							}(_pat_let_tv4)
						}(_pat_let7_0)
					}((_pat_let_tv3).Cardinality())
				}(_pat_let6_0)
			}(_322_v38)).Dtor_cf9(), _283_v0, (func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(401), _dafny.IntOfInt64(516))); ; {
					_compr_9, _ok10 := _iter10()
					if !_ok10 {
						break
					}
					var _329_v41 _dafny.Int
					_329_v41 = interface{}(_compr_9).(_dafny.Int)
					if ((_dafny.IntOfInt64(401)).Cmp(_329_v41) <= 0) && ((_329_v41).Cmp(_dafny.IntOfInt64(516)) < 0) {
						_coll9.Add((_329_v41).Plus(p1), p1)
					}
				}
				return _coll9.ToMap()
			}()).Cardinality())
			var _330_v43 _dafny.Sequence
			_ = _330_v43
			_330_v43 = _dafny.SeqOf(_325_v42, _325_v42)
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_320_v36), 0))
			_ = _index41
			(_320_v36).ArraySet1(p1, (_index41).Int())
			var _331_v44 _dafny.Set
			_ = _331_v44
			_331_v44 = _dafny.SetOf(_320_v36)
			var _332_v45 _dafny.Map
			_ = _332_v45
			_332_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _283_v0), _283_v0)
			var _333_v46 D3
			_ = _333_v46
			_333_v46 = Companion_D3_.Create_DC10_(_332_v45)
			var _334_v47 _dafny.Set
			_ = _334_v47
			_334_v47 = _dafny.SetOf(_dafny.MultiSetOf(_333_v46, _333_v46, Companion_D3_.Create_DC10_(_332_v45)))
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_320_v36), 0))
			_ = _index42
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_320_v36), 0))
			_ = _index43
			var _rhs47 _dafny.Int = (_this).F6()
			_ = _rhs47
			var _rhs48 _dafny.Sequence = _dafny.SeqOf(Companion_D1_.Create_DC4_(Companion_D0_.Create_DC0_(_312_v28, _dafny.IntOfInt64(-852)), (_this).F6(), _323_v39, _283_v0, (_this).F6()))
			_ = _rhs48
			var _rhs49 _dafny.Int = (((_334_v47).Difference(_334_v47)).Union(_334_v47)).Cardinality()
			_ = _rhs49
			var _rhs50 _dafny.Set = _dafny.SetOf(_320_v36, _320_v36, _320_v36, _320_v36)
			_ = _rhs50
			var _lhs36 _dafny.Array = _320_v36
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_320_v36), 0))
			_ = _lhs37
			var _lhs38 _dafny.Array = _320_v36
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_320_v36), 0))
			_ = _lhs39
			(_lhs36).ArraySet1(_rhs47, (_lhs37).Int())
			_330_v43 = _rhs48
			(_lhs38).ArraySet1(_rhs49, (_lhs39).Int())
			_331_v44 = _rhs50
		}
		var _335_v48 _dafny.Array
		_ = _335_v48
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_3
		var _nw43 _dafny.Array
		_ = _nw43
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw43 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Sequence = func(_336_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(651))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}(func(_337_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				}))
			}
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw43 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw43).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw43).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_335_v48 = _nw43
		var _338_v49 _dafny.Sequence
		_ = _338_v49
		_338_v49 = _dafny.UnicodeSeqOfUtf8Bytes("aabonynb")
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_335_v48), 0))
		_ = _index44
		(_335_v48).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_338_v49, _338_v49), (_index44).Int())
		var _339_v50 _dafny.Map
		_ = _339_v50
		_339_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_this).F6(), (_this).F6()), _dafny.IntOfUint32((_338_v49).Cardinality()))
		var _340_v51 _dafny.Map
		_ = _340_v51
		_340_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v0, true)
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_335_v48), 0))
		_ = _index45
		var _rhs51 _dafny.Int = (func() _dafny.Int {
			if (_339_v50).Contains((_dafny.IntOfInt64(731)).Minus((_this).F6())) {
				return (_339_v50).Get((_dafny.IntOfInt64(731)).Minus((_this).F6())).(_dafny.Int)
			}
			return (_this).F6()
		})()
		_ = _rhs51
		var _rhs52 bool = _283_v0
		_ = _rhs52
		var _rhs53 bool = (p1).Cmp((_340_v51).Cardinality()) <= 0
		_ = _rhs53
		var _rhs54 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_338_v49, _338_v49)
		_ = _rhs54
		var _lhs40 _dafny.Array = _335_v48
		_ = _lhs40
		var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_335_v48), 0))
		_ = _lhs41
		r1 = _rhs51
		_283_v0 = _rhs52
		_283_v0 = _rhs53
		(_lhs40).ArraySet1(_rhs54, (_lhs41).Int())
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_335_v48), 0))
		_ = _index46
		(_335_v48).ArraySet1(_338_v49, (_index46).Int())
		var _hi1 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F6(), p1)
		_ = _hi1
		for _341_i3 := p1; _341_i3.Cmp(_hi1) < 0; _341_i3 = _341_i3.Plus(_dafny.One) {
			var _342_v52 *C0
			_ = _342_v52
			var _nw44 *C0 = New_C0_()
			_ = _nw44
			_nw44.Ctor__(_this.F9(), (func() bool {
				if _283_v0 {
					return _283_v0
				}
				return _283_v0
			})())
			_342_v52 = _nw44
			var _343_v53 _dafny.CodePoint
			_ = _343_v53
			_343_v53 = _dafny.CodePoint('t')
			var _344_v54 _dafny.Map
			_ = _344_v54
			_344_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_343_v53, _342_v52.F20)
			(_342_v52).F20 = !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_343_v53, _283_v0)), (_344_v54).Merge(_344_v54))
			r3 = _dafny.IntOfInt64(15)
			var _345_v55 _dafny.MultiSet
			_ = _345_v55
			_345_v55 = _dafny.MultiSetOf(_342_v52.F20)
			var _346_v56 D1
			_ = _346_v56
			_346_v56 = Companion_D1_.Create_DC4_(func(_pat_let9_0 D0) D0 {
				return func(_347_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let10_0 _dafny.Int) D0 {
						return func(_348_dt__update_hcf1_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC0_((_347_dt__update__tmp_h1).Dtor_cf0(), _348_dt__update_hcf1_h0)
						}(_pat_let10_0)
					}(_dafny.IntOfInt64(811))
				}(_pat_let9_0)
			}(Companion_Default___.Fm32(_283_v0, _342_v52.F20, _341_i3, _283_v0, globalState)), (func() _dafny.Int {
				if (_345_v55).Contains(_342_v52.F20) {
					return (_345_v55).Multiplicity(_342_v52.F20)
				}
				return (_this).F6()
			})(), _338_v49, _283_v0, (_dafny.Zero).Minus(p1))
			var _349_v57 _dafny.Set
			_ = _349_v57
			_349_v57 = _dafny.SetOf(!(_283_v0))
			var _350_v58 D5
			_ = _350_v58
			_350_v58 = Companion_D5_.Create_DC16_(_349_v57)
			var _pat_let_tv5 = _350_v58
			_ = _pat_let_tv5
			var _351_v59 _dafny.Map
			_ = _351_v59
			_351_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_v52.F20, func(_pat_let11_0 D1) D1 {
				return func(_352_dt__update__tmp_h2 D1) D1 {
					return func(_pat_let12_0 _dafny.Int) D1 {
						return func(_353_dt__update_hcf8_h1 _dafny.Int) D1 {
							return Companion_D1_.Create_DC4_((_352_dt__update__tmp_h2).Dtor_cf7(), _353_dt__update_hcf8_h1, (_352_dt__update__tmp_h2).Dtor_cf9(), (_352_dt__update__tmp_h2).Dtor_cf10(), (_352_dt__update__tmp_h2).Dtor_cf11())
						}(_pat_let12_0)
					}(((_pat_let_tv5).Dtor_cf35()).Cardinality())
				}(_pat_let11_0)
			}(_346_v56))
			_351_v59 = (_351_v59).Update(!(_283_v0), _346_v56)
		}
		var _nw45 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
		_ = _nw45
		(_this).F9_set_(_nw45)
		var _354_v60 _dafny.Sequence
		_ = _354_v60
		_354_v60 = _dafny.SeqOf((_this).F6())
		var _355_v61 D3
		_ = _355_v61
		_355_v61 = Companion_D3_.Create_DC9_(_dafny.SeqOf((_this).F6(), (_this).F6()))
		var _356_v62 _dafny.Sequence
		_ = _356_v62
		_356_v62 = _dafny.SeqOf(_this.F9(), (func() _dafny.Array {
			if _283_v0 {
				return _this.F9()
			}
			return _this.F9()
		})(), _this.F9(), _this.F9())
		var _357_v64 D3
		_ = _357_v64
		_357_v64 = Companion_D3_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(620), _dafny.IntOfInt64(163))); ; {
				_compr_10, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _358_v63 _dafny.Int
				_358_v63 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(620)).Cmp(_358_v63) <= 0) && ((_358_v63).Cmp(_dafny.IntOfInt64(163)) < 0) {
					_coll10.Add((_358_v63).Times(_dafny.IntOfUint32((_354_v60).Cardinality())), _283_v0)
				}
			}
			return _coll10.ToMap()
		}(), _283_v0))
		var _pat_let_tv6 = p1
		_ = _pat_let_tv6
		var _pat_let_tv7 = p1
		_ = _pat_let_tv7
		var _pat_let_tv8 = p1
		_ = _pat_let_tv8
		var _pat_let_tv9 = _283_v0
		_ = _pat_let_tv9
		var _pat_let_tv10 = _283_v0
		_ = _pat_let_tv10
		var _rhs55 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_338_v49, _338_v49)).Cardinality())).Minus(Companion_Default___.SafeDivisionInt((_this).F6(), _dafny.IntOfUint32((_354_v60).Cardinality())))
		_ = _rhs55
		var _rhs56 bool = func(_source6 D3) bool {
			if _source6.Is_DC10() {
				var _359___mcc_h0 _dafny.Map = _source6.Get_().(D3_DC10).Cf22
				_ = _359___mcc_h0
				var _360_cf22 _dafny.Map = _359___mcc_h0
				_ = _360_cf22
				return !(!(((_dafny.Zero).Minus(_pat_let_tv6)).Cmp(_pat_let_tv7) != 0))
			} else if _source6.Is_DC9() {
				var _361___mcc_h1 _dafny.Sequence = _source6.Get_().(D3_DC9).Cf21
				_ = _361___mcc_h1
				var _362_cf21 _dafny.Sequence = _361___mcc_h1
				_ = _362_cf21
				return ((_dafny.Zero).Minus(_pat_let_tv8)).Cmp((_this).F6()) != 0
			} else {
				var _363___mcc_h2 D3 = _source6.Get_().(D3_DC11).Cf23
				_ = _363___mcc_h2
				var _364_cf23 D3 = _363___mcc_h2
				_ = _364_cf23
				return !(_pat_let_tv9) || (_pat_let_tv10)
			}
		}(Companion_D3_.Create_DC11_(Companion_D3_.Create_DC11_(Companion_D3_.Create_DC11_(Companion_D3_.Create_DC11_(_355_v61)))))
		_ = _rhs56
		var _rhs57 _dafny.Array = (_356_v62).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(((_this).F6()).Times((_this).F6())), _dafny.IntOfUint32((_356_v62).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs57
		var _rhs58 bool = !(_283_v0)
		_ = _rhs58
		var _rhs59 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F6(), ((Companion_Default___.Fm33(_357_v64, _dafny.IntOfInt64(969), _283_v0, _283_v0, globalState)).Cardinality()).Plus((_this).F6()))
		_ = _rhs59
		var _lhs42 *C2 = _this
		_ = _lhs42
		var _lhs43 *GlobalState = globalState
		_ = _lhs43
		var _lhs44 *GlobalState = globalState
		_ = _lhs44
		r3 = _rhs55
		_283_v0 = _rhs56
		_lhs42.F9_set_(_rhs57)
		_lhs43.F0 = _rhs58
		_lhs44.F5 = _rhs59
		r0 = _dafny.CodePoint('j')
		r1 = Companion_Default___.Fm2(p1, globalState)
		r2 = _this.F9()
		r3 = (((_this).F6()).Plus(_dafny.IntOfUint32((_338_v49).Cardinality()))).Minus((_this).F6())
		return r0, r1, r2, r3
	}
}
func (_this *C2) M15(p0 bool, p1 bool, p2 _dafny.Map, globalState *GlobalState) {
	{
		(globalState).F0 = !(p0)
		var _hi2 _dafny.Int = (_dafny.IntOfInt64(442)).Minus(_dafny.IntOfInt64(553))
		_ = _hi2
		for _365_i0 := _dafny.IntOfInt64(832); _365_i0.Cmp(_hi2) < 0; _365_i0 = _365_i0.Plus(_dafny.One) {
			var _366_v0 *C1
			_ = _366_v0
			var _nw46 *C1 = New_C1_()
			_ = _nw46
			_nw46.Ctor__((_this).F6())
			_366_v0 = _nw46
			var _367_v1 _dafny.Array
			_ = _367_v1
			var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw47
			_367_v1 = _nw47
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))
			_ = _index47
			(_367_v1).ArraySet1((_365_i0).Plus(_365_i0), (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_367_v1), 0))
			_ = _index48
			(_367_v1).ArraySet1(_365_i0, (_index48).Int())
			var _368_v2 _dafny.Array
			_ = _368_v2
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_4
			var _nw48 _dafny.Array
			_ = _nw48
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw48 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Set = func(_369_i1 _dafny.Int) _dafny.Set {
					return (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F6())).Cardinality()))).Difference(_dafny.SetOf((_this).F6()))
				}
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw48 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw48).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw48).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_368_v2 = _nw48
			var _370_v3 _dafny.Set
			_ = _370_v3
			_370_v3 = _dafny.SetOf(_365_i0)
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_368_v2), 0))
			_ = _index49
			(_368_v2).ArraySet1((_370_v3).Intersection(_370_v3), (_index49).Int())
			var _371_v4 _dafny.Sequence
			_ = _371_v4
			_371_v4 = _dafny.SeqOf(p1)
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))
			_ = _index50
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_367_v1), 0))
			_ = _index51
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_368_v2), 0))
			_ = _index52
			var _rhs60 _dafny.Int = (_this).F6()
			_ = _rhs60
			var _rhs61 _dafny.Int = (func() _dafny.Int {
				if _dafny.Companion_Sequence_.Equal(_371_v4, _371_v4) {
					return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(682), _dafny.IntOfInt64(-15))
				}
				return ((_this).F6()).Minus((_this).F6())
			})()
			_ = _rhs61
			var _rhs62 _dafny.Set = _370_v3
			_ = _rhs62
			var _lhs45 _dafny.Array = _367_v1
			_ = _lhs45
			var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))
			_ = _lhs46
			var _lhs47 _dafny.Array = _367_v1
			_ = _lhs47
			var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_367_v1), 0))
			_ = _lhs48
			var _lhs49 _dafny.Array = _368_v2
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_368_v2), 0))
			_ = _lhs50
			(_lhs45).ArraySet1(_rhs60, (_lhs46).Int())
			(_lhs47).ArraySet1(_rhs61, (_lhs48).Int())
			(_lhs49).ArraySet1(_rhs62, (_lhs50).Int())
			var _372_v5 _dafny.Sequence
			_ = _372_v5
			_372_v5 = _dafny.UnicodeSeqOfUtf8Bytes("fwds")
			if !(true) || ((_371_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_372_v5).Cardinality()), _dafny.IntOfUint32((_371_v4).Cardinality()))).Uint32()).(bool)) {
				var _373_v6 _dafny.Array
				_ = _373_v6
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_5
				var _nw49 _dafny.Array
				_ = _nw49
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw49 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) _dafny.Sequence = (func(_374_v5 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_375_i2 _dafny.Int) _dafny.Sequence {
							return _374_v5
						}
					})(_372_v5)
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw49 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw49).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw49).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_373_v6 = _nw49
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_373_v6), 0))
				_ = _index53
				(_373_v6).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_372_v5, _372_v5), (_index53).Int())
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_373_v6), 0))
				_ = _index54
				(_373_v6).ArraySet1(_372_v5, (_index54).Int())
				var _376_v7 _dafny.Map
				_ = _376_v7
				_376_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
				var _377_v8 _dafny.Sequence
				_ = _377_v8
				_377_v8 = _dafny.SeqOf((_376_v7).Cardinality(), _365_i0)
				var _378_v9 D4
				_ = _378_v9
				_378_v9 = Companion_D4_.Create_DC13_(_377_v8, p1, (_this).F6())
				var _379_v10 _dafny.Map
				_ = _379_v10
				_379_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm21(_365_i0, globalState), (_378_v9).Dtor_cf26())
				var _380_v11 _dafny.Set
				_ = _380_v11
				_380_v11 = _dafny.SetOf(p0, false, p1)
				var _381_v12 _dafny.Set
				_ = _381_v12
				_381_v12 = _dafny.SetOf(_380_v11, _380_v11)
				_379_v10 = (_379_v10).Update((_367_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))).Int()).(_dafny.Int), !((_381_v12).Contains(_380_v11)))
				(globalState).F5 = _365_i0
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))
				_ = _index55
				(_367_v1).ArraySet1((_367_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))).Int()).(_dafny.Int), (_index55).Int())
				var _382_v13 _dafny.Map
				_ = _382_v13
				_382_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_367_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))).Int()).(_dafny.Int), (_367_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))).Int()).(_dafny.Int))
				_382_v13 = (_382_v13).Update(Companion_Default___.SafeModuloInt((_this).F6(), _dafny.IntOfInt64(-29)), (_367_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))).Int()).(_dafny.Int))
			} else {
				var _383_v14 _dafny.CodePoint
				_ = _383_v14
				_383_v14 = _dafny.CodePoint('p')
				var _384_v15 _dafny.Map
				_ = _384_v15
				_384_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_367_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))).Int()).(_dafny.Int), p1)
				var _385_v16 _dafny.Sequence
				_ = _385_v16
				_385_v16 = _dafny.SeqOf((_367_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))).Int()).(_dafny.Int), (_384_v15).Cardinality())
				var _386_v17 _dafny.Map
				_ = _386_v17
				_386_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D4_.Create_DC13_(_385_v16, p1, _365_i0)).Dtor_cf27(), p0)
				var _387_v18 _dafny.Map
				_ = _387_v18
				_387_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v14, ((_386_v17).Update((_this).F6(), p0)).Cardinality())
				_387_v18 = (_387_v18).Update(_383_v14, _dafny.IntOfInt64(43))
				(globalState).F0 = _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(32))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}(func(_388_i3 _dafny.Int) _dafny.Int {
					return (_this).F6()
				})), (_367_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))).Int()).(_dafny.Int))
				(globalState).F5 = (_dafny.Zero).Minus((_dafny.IntOfUint32((_371_v4).Cardinality())).Times((_dafny.Zero).Minus((_367_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))).Int()).(_dafny.Int))))
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))
				_ = _index56
				(_367_v1).ArraySet1((_367_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))).Int()).(_dafny.Int), (_index56).Int())
				(globalState).F0 = p0
			}
			var _389_v19 _dafny.Map
			_ = _389_v19
			_389_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_365_i0, p0)
			var _390_v20 _dafny.Map
			_ = _390_v20
			_390_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (func() _dafny.Int {
				if !((func() bool {
					if (_389_v19).Contains(_dafny.IntOfInt64(947)) {
						return (_389_v19).Get(_dafny.IntOfInt64(947)).(bool)
					}
					return p1
				})()) {
					return (_367_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_367_v1), 0))).Int()).(_dafny.Int)
				}
				return _dafny.IntOfInt64(229)
			})())
			_390_v20 = _390_v20
		}
		var _391_v21 _dafny.Set
		_ = _391_v21
		_391_v21 = _dafny.SetOf(_this.F9())
		var _392_v22 _dafny.Map
		_ = _392_v22
		_392_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_391_v21).Cardinality())
		var _393_v23 _dafny.Sequence
		_ = _393_v23
		_393_v23 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
		var _hi3 _dafny.Int = (((_392_v22).Update(p0, (_this).F6())).Cardinality()).Minus((_393_v23).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F6()), _dafny.IntOfUint32((_393_v23).Cardinality()))).Uint32()).(_dafny.Int))
		_ = _hi3
		for _394_i4 := (_this).F6(); _394_i4.Cmp(_hi3) < 0; _394_i4 = _394_i4.Plus(_dafny.One) {
			var _395_v24 *C0
			_ = _395_v24
			var _nw50 *C0 = New_C0_()
			_ = _nw50
			_nw50.Ctor__(_this.F9(), p1)
			_395_v24 = _nw50
			var _396_v25 *C1
			_ = _396_v25
			var _nw51 *C1 = New_C1_()
			_ = _nw51
			_nw51.Ctor__((_this).F6())
			_396_v25 = _nw51
			var _397_v26 _dafny.Map
			_ = _397_v26
			_397_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _395_v24.F19)
			var _398_v27 _dafny.CodePoint
			_ = _398_v27
			_398_v27 = _dafny.CodePoint('c')
			var _399_v28 _dafny.Map
			_ = _399_v28
			_399_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_398_v27, (_this).F6())
			var _400_v29 _dafny.Map
			_ = _400_v29
			_400_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SetOf((_dafny.SetOf(_394_i4)).Cardinality()))
			(_395_v24).F19 = (func() _dafny.Array {
				if (_397_v26).Contains(Companion_Default___.SafeModuloInt((_399_v28).Cardinality(), (_400_v29).Cardinality())) {
					return (_397_v26).Get(Companion_Default___.SafeModuloInt((_399_v28).Cardinality(), (_400_v29).Cardinality())).(_dafny.Array)
				}
				return _this.F9()
			})()
			var _401_v30 *C1
			_ = _401_v30
			var _nw52 *C1 = New_C1_()
			_ = _nw52
			_nw52.Ctor__((_this).F6())
			_401_v30 = _nw52
		}
		var _402_v31 _dafny.MultiSet
		_ = _402_v31
		_402_v31 = _dafny.MultiSetOf(_393_v23)
		var _403_v32 _dafny.Sequence
		_ = _403_v32
		_403_v32 = _dafny.UnicodeSeqOfUtf8Bytes("waawxeny")
		var _404_v33 _dafny.Map
		_ = _404_v33
		_404_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v32, _dafny.SeqOf(Companion_Default___.Fm2(_dafny.IntOfUint32((_403_v32).Cardinality()), globalState)))
		var _405_v34 _dafny.Map
		_ = _405_v34
		_405_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
		var _406_v35 D4
		_ = _406_v35
		_406_v35 = Companion_D4_.Create_DC15_(p1, p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(53))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}(func(_407_i6 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		})))
		var _408_v36 _dafny.MultiSet
		_ = _408_v36
		_408_v36 = _dafny.MultiSetOf(Companion_D4_.Create_DC15_(p1, p0, _403_v32), _406_v35)
		var _409_v37 _dafny.Sequence
		_ = _409_v37
		_409_v37 = _dafny.SeqOf(_dafny.SeqOf((_this).F6()))
		var _410_v38 _dafny.Array
		_ = _410_v38
		var _nwElement0_7 _dafny.Int = (func() _dafny.Int {
			if (_402_v31).Contains((func() _dafny.Sequence {
				if (_404_v33).Contains(_403_v32) {
					return (_404_v33).Get(_403_v32).(_dafny.Sequence)
				}
				return _393_v23
			})()) {
				return (_402_v31).Multiplicity((func() _dafny.Sequence {
					if (_404_v33).Contains(_403_v32) {
						return (_404_v33).Get(_403_v32).(_dafny.Sequence)
					}
					return _393_v23
				})())
			}
			return (_405_v34).Cardinality()
		})()
		_ = _nwElement0_7
		var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(11))
		_ = _nw53
		(_nw53).ArraySet1(_nwElement0_7, 0)
		(_nw53).ArraySet1((func() _dafny.Int {
			if (_392_v22).Contains(p0) {
				return (_392_v22).Get(p0).(_dafny.Int)
			}
			return (_this).F6()
		})(), 1)
		(_nw53).ArraySet1(((_408_v36).Cardinality()).Times(_dafny.IntOfInt64(383)), 2)
		(_nw53).ArraySet1((_this).F6(), 3)
		(_nw53).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F6(), _dafny.IntOfUint32((_403_v32).Cardinality())), 4)
		(_nw53).ArraySet1(_dafny.IntOfUint32((_393_v23).Cardinality()), 5)
		(_nw53).ArraySet1((func() _dafny.Int {
			if p1 {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())).Cardinality()
			}
			return (_this).F6()
		})(), 6)
		(_nw53).ArraySet1((_this).F6(), 7)
		(_nw53).ArraySet1((_this).F6(), 8)
		(_nw53).ArraySet1(((_this).F6()).Times((Companion_Default___.Fm34(_dafny.IntOfUint32(((_409_v37).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_409_v37).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), globalState)).Cardinality()), 9)
		(_nw53).ArraySet1((_393_v23).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.IntOfUint32((_393_v23).Cardinality()))).Uint32()).(_dafny.Int), 10)
		_410_v38 = _nw53
		for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_410_v38), 0))); ; {
			_guard_loop_1, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _411_i5 _dafny.Int
			_411_i5 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_411_i5).Sign() != -1) && ((_411_i5).Cmp(_dafny.ArrayLen((_410_v38), 0)) < 0)) {
				(_410_v38).ArraySet1((_411_i5).Times(_dafny.IntOfInt64(404)), (_411_i5).Int())
			}
		}
		var _hi4 _dafny.Int = (_this).F6()
		_ = _hi4
		for _412_i7 := (_dafny.Zero).Minus((_this).F6()); _412_i7.Cmp(_hi4) < 0; _412_i7 = _412_i7.Plus(_dafny.One) {
			(globalState).F0 = p0
			var _413_v39 _dafny.Sequence
			_ = _413_v39
			_413_v39 = _dafny.SeqOf(p1)
			_413_v39 = _413_v39
			var _414_v40 _dafny.Array
			_ = _414_v40
			var _nwElement0_8 _dafny.Sequence = _dafny.SeqOf(p1)
			_ = _nwElement0_8
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(9))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_8, 0)
			(_nw54).ArraySet1(_413_v39, 1)
			(_nw54).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_413_v39, _413_v39), 2)
			(_nw54).ArraySet1(_413_v39, 3)
			(_nw54).ArraySet1(_413_v39, 4)
			(_nw54).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_413_v39, _413_v39), 5)
			(_nw54).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm3(globalState), p1, p0, p0, p1), _dafny.SeqOf(p0, p1)), 6)
			(_nw54).ArraySet1(_413_v39, 7)
			(_nw54).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_413_v39, _dafny.SeqOf(true, p1)), 8)
			_414_v40 = _nw54
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_414_v40), 0))
			_ = _index57
			(_414_v40).ArraySet1(_413_v39, (_index57).Int())
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_414_v40), 0))
			_ = _index58
			(_414_v40).ArraySet1(_413_v39, (_index58).Int())
			var _415_v41 _dafny.Array
			_ = _415_v41
			var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
			_ = _nw55
			_415_v41 = _nw55
			var _416_v42 _dafny.Array
			_ = _416_v42
			var _nwElement0_9 _dafny.Array = _this.F9()
			_ = _nwElement0_9
			var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(4))
			_ = _nw56
			(_nw56).ArraySet1(_nwElement0_9, 0)
			(_nw56).ArraySet1(_this.F9(), 1)
			(_nw56).ArraySet1(_this.F9(), 2)
			(_nw56).ArraySet1(_this.F9(), 3)
			_416_v42 = _nw56
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_415_v41), 0))
			_ = _index59
			(_415_v41).ArraySet1(_416_v42, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_415_v41), 0))
			_ = _index60
			(_415_v41).ArraySet1(_416_v42, (_index60).Int())
		}
		(globalState).F0 = p0
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f9  _dafny.Array
	_f6  _dafny.Int
	_f10 _dafny.Map
	_f21 _dafny.Map
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.Zero
	_this._f10 = _dafny.EmptyMap
	_this._f21 = _dafny.EmptyMap
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F9() _dafny.Array {
	return _this._f9
}
func (_this *C3) F9_set_(value _dafny.Array) {
	_this._f9 = value
}
func (_this *C3) F6() _dafny.Int {
	return _this._f6
}
func (_this *C3) F10() _dafny.Map {
	return _this._f10
}
func (_this *C3) Ctor__(f21 _dafny.Map, f9 _dafny.Array, f10 _dafny.Map, f6 _dafny.Int) {
	{
		(_this)._f21 = f21
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f6 = f6
	}
}
func (_this *C3) Fm13(p0 _dafny.CodePoint, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("o"), _dafny.UnicodeSeqOfUtf8Bytes("qhkakwoox")))
	}
}
func (_this *C3) Fm5(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('b')), _dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('b'), _dafny.CodePoint('n'), _dafny.CodePoint('b')))), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('m'), _dafny.CodePoint('g'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('e')))).Intersection(_dafny.MultiSetOf(_dafny.CodePoint('n'))), _dafny.MultiSetOf(_dafny.CodePoint('s')), _dafny.MultiSetOf(_dafny.CodePoint('p'), _dafny.CodePoint('k'), _dafny.CodePoint('x'), _dafny.CodePoint('g'), _dafny.CodePoint('y')), _dafny.MultiSetOf(_dafny.CodePoint('i')))
	}
}
func (_this *C3) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.SeqOf(_dafny.IntOfInt64(-371), (_this).F6(), (_this).F6(), (_this).F6(), (_dafny.Zero).Minus((_this).F6())))
	}
}
func (_this *C3) M6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _417_v0 _dafny.Sequence
		_ = _417_v0
		_417_v0 = _dafny.SeqOf((_this).F6(), _dafny.IntOfInt64(-821))
		var _418_v1 D3
		_ = _418_v1
		_418_v1 = Companion_D3_.Create_DC9_(_417_v0)
		var _pat_let_tv11 = _417_v0
		_ = _pat_let_tv11
		(globalState).F5 = (_dafny.Zero).Minus(_dafny.IntOfUint32(((func(_pat_let13_0 D3) D3 {
			return func(_419_dt__update__tmp_h0 D3) D3 {
				return func(_pat_let14_0 _dafny.Sequence) D3 {
					return func(_420_dt__update_hcf21_h0 _dafny.Sequence) D3 {
						return Companion_D3_.Create_DC9_(_420_dt__update_hcf21_h0)
					}(_pat_let14_0)
				}(_pat_let_tv11)
			}(_pat_let13_0)
		}(_418_v1)).Dtor_cf21()).Cardinality()))
		var _421_v2 _dafny.Array
		_ = _421_v2
		var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
		_ = _nw57
		_421_v2 = _nw57
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))
		_ = _index61
		(_421_v2).ArraySet1(p2, (_index61).Int())
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))
		_ = _index62
		(_421_v2).ArraySet1((_this).F6(), (_index62).Int())
		(globalState).F2 = _dafny.UnicodeSeqOfUtf8Bytes("djjjt")
		_417_v0 = _417_v0
		var _422_v3 _dafny.Array
		_ = _422_v3
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_6
		var _nw58 _dafny.Array
		_ = _nw58
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw58 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.MultiSet = (func(_423_v0 _dafny.Sequence) func(_dafny.Int) _dafny.MultiSet {
				return func(_424_i0 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetFromSeq(_423_v0)
				}
			})(_417_v0)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw58 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw58).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw58).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_422_v3 = _nw58
		var _425_v4 D2
		_ = _425_v4
		_425_v4 = Companion_D2_.Create_DC7_(p1, _422_v3, ((_this).F6()).Minus((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int)), ((_this).F6()).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("olpetvk")).Cardinality())))
		var _source7 D2 = _425_v4
		_ = _source7
		if _source7.Is_DC6() {
			var _426___mcc_h0 _dafny.Int = _source7.Get_().(D2_DC6).Cf13
			_ = _426___mcc_h0
			var _427___mcc_h1 _dafny.Array = _source7.Get_().(D2_DC6).Cf14
			_ = _427___mcc_h1
			var _428_cf14 _dafny.Array = _427___mcc_h1
			_ = _428_cf14
			var _429_cf13 _dafny.Int = _426___mcc_h0
			_ = _429_cf13
			var _arr4 _dafny.Array = _this.F9()
			_ = _arr4
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index63
			(_arr4).ArraySet1(p0, (_index63).Int())
			var _430_v5 _dafny.Map
			_ = _430_v5
			_430_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9(), p0)
			var _431_v6 D0
			_ = _431_v6
			_431_v6 = Companion_D0_.Create_DC0_(p1, _dafny.IntOfUint32((p1).Cardinality()))
			var _432_v7 _dafny.Set
			_ = _432_v7
			_432_v7 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("hjgjnji"))
			var _pat_let_tv12 = _421_v2
			_ = _pat_let_tv12
			var _pat_let_tv13 = _421_v2
			_ = _pat_let_tv13
			var _433_v8 D1
			_ = _433_v8
			_433_v8 = Companion_D1_.Create_DC4_(func(_pat_let15_0 D0) D0 {
				return func(_434_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let16_0 _dafny.Int) D0 {
						return func(_435_dt__update_hcf1_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC0_((_434_dt__update__tmp_h1).Dtor_cf0(), _435_dt__update_hcf1_h0)
						}(_pat_let16_0)
					}((_pat_let_tv13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_pat_let_tv12), 0))).Int()).(_dafny.Int))
				}(_pat_let15_0)
			}(_431_v6), _429_cf13, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(921))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg40 _dafny.Int) interface{} {
					return coer40(arg40)
				}
			}(func(_436_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			})), true, (_432_v7).Cardinality())
			var _arr5 _dafny.Array = _this.F9()
			_ = _arr5
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index64
			var _rhs63 _dafny.Int = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9(), p0)).Merge(_430_v5)).Merge(_430_v5)).Cardinality()
			_ = _rhs63
			var _rhs64 bool = (_433_v8).Dtor_cf10()
			_ = _rhs64
			var _rhs65 bool = p0
			_ = _rhs65
			var _lhs51 *GlobalState = globalState
			_ = _lhs51
			var _lhs52 *GlobalState = globalState
			_ = _lhs52
			var _lhs53 _dafny.Array = _this.F9()
			_ = _lhs53
			var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_this.F9()), 0))
			_ = _lhs54
			_lhs51.F5 = _rhs63
			_lhs52.F0 = _rhs64
			(_lhs53).ArraySet1(_rhs65, (_lhs54).Int())
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))
			_ = _index65
			(_421_v2).ArraySet1(_429_cf13, (_index65).Int())
			var _437_v9 _dafny.Map
			_ = _437_v9
			_437_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_429_cf13, globalState), p0)
			var _438_v10 _dafny.Map
			_ = _438_v10
			_438_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (func() bool {
				if (_437_v9).Contains((_this).F6()) {
					return (_437_v9).Get((_this).F6()).(bool)
				}
				return (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)
			})())
			_437_v9 = (_437_v9).Update(_429_cf13, p0)
			var _arr6 _dafny.Array = _this.F9()
			_ = _arr6
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index66
			(_arr6).ArraySet1(p0, (_index66).Int())
		} else if _source7.Is_DC7() {
			var _439___mcc_h2 _dafny.Sequence = _source7.Get_().(D2_DC7).Cf15
			_ = _439___mcc_h2
			var _440___mcc_h3 _dafny.Array = _source7.Get_().(D2_DC7).Cf16
			_ = _440___mcc_h3
			var _441___mcc_h4 _dafny.Int = _source7.Get_().(D2_DC7).Cf17
			_ = _441___mcc_h4
			var _442___mcc_h5 _dafny.Int = _source7.Get_().(D2_DC7).Cf18
			_ = _442___mcc_h5
			var _443_cf18 _dafny.Int = _442___mcc_h5
			_ = _443_cf18
			var _444_cf17 _dafny.Int = _441___mcc_h4
			_ = _444_cf17
			var _445_cf16 _dafny.Array = _440___mcc_h3
			_ = _445_cf16
			var _446_cf15 _dafny.Sequence = _439___mcc_h2
			_ = _446_cf15
			var _447_v11 _dafny.Map
			_ = _447_v11
			_447_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9(), p0)
			_447_v11 = (_447_v11).Update(_this.F9(), p0)
			var _448_v12 _dafny.MultiSet
			_ = _448_v12
			_448_v12 = _dafny.MultiSetOf(p0, p0, p0, !(true))
			var _449_v13 _dafny.Sequence
			_ = _449_v13
			_449_v13 = _dafny.UnicodeSeqOfUtf8Bytes("o")
			var _rhs66 _dafny.Int = Companion_Default___.SafeDivisionInt((p2).Plus(((_448_v12).Update(p0, Companion_Default___.Abs(_443_cf18))).Cardinality()), p2)
			_ = _rhs66
			var _rhs67 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lcenstrw"), _449_v13)
			_ = _rhs67
			var _rhs68 _dafny.Int = (_this).F6()
			_ = _rhs68
			var _rhs69 bool = p0
			_ = _rhs69
			var _lhs55 *GlobalState = globalState
			_ = _lhs55
			var _lhs56 *GlobalState = globalState
			_ = _lhs56
			var _lhs57 *GlobalState = globalState
			_ = _lhs57
			var _lhs58 *GlobalState = globalState
			_ = _lhs58
			_lhs55.F5 = _rhs66
			_lhs56.F2 = _rhs67
			_lhs57.F5 = _rhs68
			_lhs58.F0 = _rhs69
			var _450_v14 _dafny.Set
			_ = _450_v14
			_450_v14 = _dafny.SetOf(false, p0)
			var _451_v15 _dafny.Map
			_ = _451_v15
			_451_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _450_v14)
			_451_v15 = (_451_v15).Update(p0, _450_v14)
			var _452_v16 T0
			_ = _452_v16
			var _nw59 *C1 = New_C1_()
			_ = _nw59
			_nw59.Ctor__(Companion_Default___.SafeDivisionInt(_444_cf17, _dafny.IntOfInt64(67)))
			_452_v16 = _nw59
			var _453_v17 _dafny.CodePoint
			_ = _453_v17
			_453_v17 = _dafny.CodePoint('o')
			var _454_v18 _dafny.Map
			_ = _454_v18
			_454_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC1_(p0, (_452_v16).F6(), _453_v17)).Dtor_cf2(), (func() _dafny.MultiSet {
				if p0 {
					return _dafny.MultiSetOf(p0, p0)
				}
				return _dafny.MultiSetFromSeq(_446_cf15)
			})())
			var _rhs70 T0 = _452_v16
			_ = _rhs70
			var _rhs71 _dafny.Map = (func() _dafny.Map {
				if p0 {
					return _454_v18
				}
				return _454_v18
			})()
			_ = _rhs71
			_452_v16 = _rhs70
			_454_v18 = _rhs71
		} else if _source7.Is_DC8() {
			var _455___mcc_h6 bool = _source7.Get_().(D2_DC8).Cf19
			_ = _455___mcc_h6
			var _456___mcc_h7 bool = _source7.Get_().(D2_DC8).Cf20
			_ = _456___mcc_h7
			var _457_cf20 bool = _456___mcc_h7
			_ = _457_cf20
			var _458_cf19 bool = _455___mcc_h6
			_ = _458_cf19
			if ((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int)).Cmp((func() _dafny.Set {
				var _coll11 = _dafny.NewBuilder()
				_ = _coll11
				for _iter13 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(619))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}((func(_459_p1 _dafny.Sequence, _460_cf20 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_461_i2 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Update(_459_p1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.IntOfUint32((_459_p1).Cardinality()))).Uint32(), _460_cf20)
					}
				})(p1, _457_cf20)))).Elements()); ; {
					_compr_11, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _462_v19 _dafny.Sequence
					_462_v19 = interface{}(_compr_11).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(619))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg42 _dafny.Int) interface{} {
							return coer42(arg42)
						}
					}((func(_463_p1 _dafny.Sequence, _464_cf20 bool) func(_dafny.Int) _dafny.Sequence {
						return func(_461_i2 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Update(_463_p1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.IntOfUint32((_463_p1).Cardinality()))).Uint32(), _464_cf20)
						}
					})(p1, _457_cf20))), _462_v19) {
						_coll11.Add(_462_v19)
					}
				}
				return _coll11.ToSet()
			}()).Cardinality()) == 0 {
				var _465_v20 _dafny.Sequence
				_ = _465_v20
				_465_v20 = _dafny.UnicodeSeqOfUtf8Bytes("hgq")
				(globalState).F2 = _465_v20
				var _466_v21 _dafny.Int
				_ = _466_v21
				var _467_v22 _dafny.CodePoint
				_ = _467_v22
				var _468_v23 bool
				_ = _468_v23
				var _out16 _dafny.Int
				_ = _out16
				var _out17 _dafny.CodePoint
				_ = _out17
				var _out18 bool
				_ = _out18
				_out16, _out17, _out18 = (_this).M14(Companion_Default___.Fm3(globalState), true, globalState)
				_466_v21 = _out16
				_467_v22 = _out17
				_468_v23 = _out18
				(globalState).F5 = (p2).Minus((_this).F6())
				_458_cf19 = !((_468_v23) == (p0))
				var _469_v24 *C0
				_ = _469_v24
				var _nw60 *C0 = New_C0_()
				_ = _nw60
				_nw60.Ctor__(_this.F9(), p0)
				_469_v24 = _nw60
			} else {
				var _470_v26 _dafny.CodePoint
				_ = _470_v26
				_470_v26 = _dafny.CodePoint('r')
				var _471_v27 _dafny.Map
				_ = _471_v27
				_471_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_470_v26, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-134))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}((func(_472_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_473_i3 _dafny.Int) _dafny.CodePoint {
						return _472_v26
					}
				})(_470_v26)))).Cardinality()))
				var _474_v28 _dafny.Map
				_ = _474_v28
				_474_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _457_cf20)
				var _475_v29 _dafny.MultiSet
				_ = _475_v29
				_475_v29 = _dafny.MultiSetOf((_474_v28).Cardinality(), p2, (_this).F6())
				var _476_v30 _dafny.Map
				_ = _476_v30
				_476_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int), p0)
				(globalState).F0 = !((_475_v29).Update((_476_v30).Cardinality(), Companion_Default___.Abs((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int)))).Contains((func() _dafny.Map {
					var _coll12 = _dafny.NewMapBuilder()
					_ = _coll12
					for _iter14 := _dafny.Iterate(((_471_v27).Update(_470_v26, (_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int))).Keys().Elements()); ; {
						_compr_12, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _477_v25 _dafny.CodePoint
						_477_v25 = interface{}(_compr_12).(_dafny.CodePoint)
						if ((_471_v27).Update(_470_v26, (_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int))).Contains(_477_v25) {
							_coll12.Add(_477_v25, p0)
						}
					}
					return _coll12.ToMap()
				}()).Cardinality())
				var _478_v31 _dafny.Array
				_ = _478_v31
				var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
				_ = _nw61
				_478_v31 = _nw61
				_478_v31 = _478_v31
				var _479_v32 _dafny.Map
				_ = _479_v32
				_479_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_458_cf19, _this.F9())
				var _480_v33 _dafny.Set
				_ = _480_v33
				_480_v33 = _dafny.SetOf(_dafny.IntOfInt64(790))
				var _481_v34 _dafny.MultiSet
				_ = _481_v34
				_481_v34 = _dafny.MultiSetOf(_458_cf19, _458_cf19)
				var _482_v35 _dafny.MultiSet
				_ = _482_v35
				_482_v35 = _dafny.MultiSetOf(Companion_Default___.Fm35(p2, (_480_v33).Cardinality(), _458_cf19, _dafny.IntOfUint32((p1).Cardinality()), globalState), _dafny.MultiSetOf(((_481_v34).Update(false, Companion_Default___.Abs((_this).F6()))).Cardinality()), _475_v29)
				var _483_v36 T1
				_ = _483_v36
				var _nw62 *C2 = New_C2_()
				_ = _nw62
				_nw62.Ctor__(p2, _482_v35, _this.F9(), (_this).F10())
				_483_v36 = _nw62
				var _484_v37 _dafny.Array
				_ = _484_v37
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_7
				var _nw63 _dafny.Array
				_ = _nw63
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw63 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.Sequence = (func(_485_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_486_i4 _dafny.Int) _dafny.Sequence {
							return _485_p1
						}
					})(p1)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw63 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw63).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw63).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_484_v37 = _nw63
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_484_v37), 0))
				_ = _index67
				(_484_v37).ArraySet1(p1, (_index67).Int())
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_484_v37), 0))
				_ = _index68
				var _rhs72 _dafny.Map = _479_v32
				_ = _rhs72
				var _rhs73 T1 = (Companion_D6_.Create_DC18_(_483_v36)).Dtor_cf38()
				_ = _rhs73
				var _rhs74 _dafny.Sequence = p1
				_ = _rhs74
				var _lhs59 _dafny.Array = _484_v37
				_ = _lhs59
				var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_484_v37), 0))
				_ = _lhs60
				_479_v32 = _rhs72
				_483_v36 = _rhs73
				(_lhs59).ArraySet1(_rhs74, (_lhs60).Int())
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))
				_ = _index69
				(_421_v2).ArraySet1((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int), (_index69).Int())
				var _487_v38 *C2
				_ = _487_v38
				var _nw64 *C2 = New_C2_()
				_ = _nw64
				_nw64.Ctor__((_483_v36).F6(), (_482_v35).Union(_482_v35), _this.F9(), (_this).F10())
				_487_v38 = _nw64
			}
			_457_cf20 = (func() bool {
				if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(p2, (_this).F6()), _417_v0) {
					return _458_cf19
				}
				return _457_cf20
			})()
			var _488_v39 _dafny.Map
			_ = _488_v39
			_488_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0)
			var _489_v40 D4
			_ = _489_v40
			_489_v40 = Companion_D4_.Create_DC13_(_417_v0, false, (_488_v39).Cardinality())
			_489_v40 = _489_v40
			var _490_v41 D4
			_ = _490_v41
			_490_v41 = Companion_D4_.Create_DC14_(_457_cf20, _488_v39, (_this).F6(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0))
			var _491_v42 _dafny.Map
			_ = _491_v42
			_491_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_490_v41, ((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_457_cf20, _457_cf20)).Cardinality()) < 0)
			_491_v42 = (_491_v42).Update(_490_v41, _457_cf20)
		} else {
			var _492___mcc_h8 _dafny.Map = _source7.Get_().(D2_DC5).Cf12
			_ = _492___mcc_h8
			var _493_cf12 _dafny.Map = _492___mcc_h8
			_ = _493_cf12
			var _494_v43 D2
			_ = _494_v43
			_494_v43 = Companion_D2_.Create_DC8_(p0, p0)
			var _495_v44 _dafny.Map
			_ = _495_v44
			_495_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_417_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-880))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}((func(_496_v2 _dafny.Array) func(_dafny.Int) _dafny.Int {
				return func(_497_i5 _dafny.Int) _dafny.Int {
					return (_496_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_496_v2), 0))).Int()).(_dafny.Int)
				}
			})(_421_v2))))).Cardinality()), (_494_v43).Dtor_cf20())
			var _498_v45 _dafny.Map
			_ = _498_v45
			_498_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool), p0)
			_495_v44 = (_495_v44).Update((_498_v45).Cardinality(), true)
			var _499_v46 _dafny.Array
			_ = _499_v46
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_8
			var _nw65 _dafny.Array
			_ = _nw65
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw65 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Set = func(_500_i6 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-810))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg45 _dafny.Int) interface{} {
							return coer45(arg45)
						}
					}(func(_501_i7 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('y')
					})), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-810))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg46 _dafny.Int) interface{} {
							return coer46(arg46)
						}
					}(func(_502_i7 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('y')
					}))).Cardinality()))).Uint32(), _dafny.CodePoint('l')))
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw65 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw65).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw65).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_499_v46 = _nw65
			var _503_v47 _dafny.CodePoint
			_ = _503_v47
			_503_v47 = _dafny.CodePoint('b')
			var _504_v48 _dafny.Sequence
			_ = _504_v48
			_504_v48 = _dafny.SeqOf(_503_v47, _503_v47)
			var _505_v49 _dafny.Set
			_ = _505_v49
			_505_v49 = _dafny.SetOf(_504_v48)
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_499_v46), 0))
			_ = _index70
			(_499_v46).ArraySet1(_505_v49, (_index70).Int())
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_499_v46), 0))
			_ = _index71
			(_499_v46).ArraySet1(_505_v49, (_index71).Int())
			if false {
				_503_v47 = _503_v47
				var _506_v50 _dafny.MultiSet
				_ = _506_v50
				_506_v50 = _dafny.MultiSetOf(_dafny.CodePoint('l'), _503_v47)
				var _507_v51 _dafny.MultiSet
				_ = _507_v51
				_507_v51 = _dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(619), (_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int)))
				var _508_v52 *C2
				_ = _508_v52
				var _nw66 *C2 = New_C2_()
				_ = _nw66
				_nw66.Ctor__((_dafny.Zero).Minus((_506_v50).Cardinality()), (Companion_Default___.Fm36(globalState)).Difference(_507_v51), _this.F9(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_504_v48).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(670))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_509_v47 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_510_i8 _dafny.Int) _dafny.CodePoint {
						return _509_v47
					}
				})(_503_v47)))))
				_508_v52 = _nw66
				var _511_v53 _dafny.Map
				_ = _511_v53
				_511_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(p2, (_this).F6()), _421_v2)
				_511_v53 = _511_v53
				var _arr7 _dafny.Array = _this.F9()
				_ = _arr7
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index72
				(_arr7).ArraySet1(p0, (_index72).Int())
				var _512_v54 _dafny.Map
				_ = _512_v54
				_512_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _504_v48)
				var _513_v55 _dafny.MultiSet
				_ = _513_v55
				_513_v55 = _dafny.MultiSetOf(p0)
				var _514_v56 _dafny.MultiSet
				_ = _514_v56
				_514_v56 = _dafny.MultiSetOf((_this).F6())
				var _515_v57 _dafny.Set
				_ = _515_v57
				_515_v57 = _dafny.SetOf(_dafny.IntOfInt64(739), (_dafny.Zero).Minus((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int)), (_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int))
				var _pat_let_tv14 = p1
				_ = _pat_let_tv14
				var _516_v58 _dafny.Array
				_ = _516_v58
				var _nwElement0_10 _dafny.Map = _512_v54
				_ = _nwElement0_10
				var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(16))
				_ = _nw67
				(_nw67).ArraySet1(_nwElement0_10, 0)
				(_nw67).ArraySet1((_512_v54).Update(p0, _504_v48), 1)
				(_nw67).ArraySet1(_512_v54, 2)
				(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _504_v48), 3)
				(_nw67).ArraySet1(_512_v54, 4)
				(_nw67).ArraySet1(_512_v54, 5)
				(_nw67).ArraySet1((_this).Fm13(_dafny.CodePoint('s'), _513_v55, p0, (_514_v56).Cardinality(), globalState), 6)
				(_nw67).ArraySet1(_512_v54, 7)
				(_nw67).ArraySet1(_512_v54, 8)
				(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _504_v48), 9)
				(_nw67).ArraySet1(_512_v54, 10)
				(_nw67).ArraySet1(_512_v54, 11)
				(_nw67).ArraySet1(_512_v54, 12)
				(_nw67).ArraySet1(_512_v54, 13)
				(_nw67).ArraySet1(_512_v54, 14)
				(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (Companion_D1_.Create_DC4_(func(_pat_let17_0 D0) D0 {
					return func(_517_dt__update__tmp_h2 D0) D0 {
						return func(_pat_let18_0 _dafny.Sequence) D0 {
							return func(_518_dt__update_hcf0_h0 _dafny.Sequence) D0 {
								return Companion_D0_.Create_DC0_(_518_dt__update_hcf0_h0, (_517_dt__update__tmp_h2).Dtor_cf1())
							}(_pat_let18_0)
						}(_pat_let_tv14)
					}(_pat_let17_0)
				}(Companion_D0_.Create_DC0_(p1, (_this).F6())), (_this).F6(), _504_v48, p0, (_515_v57).Cardinality())).Dtor_cf9()), 15)
				_516_v58 = _nw67
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_516_v58), 0))
				_ = _index73
				(_516_v58).ArraySet1(_512_v54, (_index73).Int())
				var _519_v59 D0
				_ = _519_v59
				_519_v59 = Companion_D0_.Create_DC1_(p0, (_this).F6(), _503_v47)
				var _pat_let_tv15 = p0
				_ = _pat_let_tv15
				var _arr8 _dafny.Array = _this.F9()
				_ = _arr8
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index74
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_516_v58), 0))
				_ = _index75
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))
				_ = _index76
				var _rhs75 bool = p0
				_ = _rhs75
				var _rhs76 _dafny.Map = _512_v54
				_ = _rhs76
				var _rhs77 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_504_v48, (Companion_Default___.SafeIndex((p2).Times(p2), _dafny.IntOfUint32((_504_v48).Cardinality()))).Uint32(), _503_v47)).Cardinality())
				_ = _rhs77
				var _rhs78 D0 = func(_pat_let19_0 D0) D0 {
					return func(_520_dt__update__tmp_h3 D0) D0 {
						return func(_pat_let20_0 bool) D0 {
							return func(_521_dt__update_hcf2_h0 bool) D0 {
								return Companion_D0_.Create_DC1_(_521_dt__update_hcf2_h0, (_520_dt__update__tmp_h3).Dtor_cf3(), (_520_dt__update__tmp_h3).Dtor_cf4())
							}(_pat_let20_0)
						}(_pat_let_tv15)
					}(_pat_let19_0)
				}(_519_v59)
				_ = _rhs78
				var _rhs79 bool = _dafny.Companion_Sequence_.Contains(p1, p0)
				_ = _rhs79
				var _lhs61 _dafny.Array = _this.F9()
				_ = _lhs61
				var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_this.F9()), 0))
				_ = _lhs62
				var _lhs63 _dafny.Array = _516_v58
				_ = _lhs63
				var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_516_v58), 0))
				_ = _lhs64
				var _lhs65 _dafny.Array = _421_v2
				_ = _lhs65
				var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))
				_ = _lhs66
				(_lhs61).ArraySet1(_rhs75, (_lhs62).Int())
				(_lhs63).ArraySet1(_rhs76, (_lhs64).Int())
				(_lhs65).ArraySet1(_rhs77, (_lhs66).Int())
				_519_v59 = _rhs78
				r1 = _rhs79
				var _522_v60 _dafny.Set
				_ = _522_v60
				_522_v60 = _dafny.SetOf(p0, (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
				var _523_v61 D0
				_ = _523_v61
				_523_v61 = Companion_D0_.Create_DC0_(_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex((_522_v60).Cardinality(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), p0), _dafny.IntOfInt64(-92))
				var _524_v62 D1
				_ = _524_v62
				_524_v62 = Companion_D1_.Create_DC4_(_523_v61, (_this).F6(), _504_v48, p0, p2)
				var _525_v63 _dafny.Set
				_ = _525_v63
				_525_v63 = _dafny.SetOf((func() bool {
					if (_495_v44).Contains((_this).F6()) {
						return (_495_v44).Get((_this).F6()).(bool)
					}
					return (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)
				})(), (_524_v62).Dtor_cf10())
				var _526_v64 _dafny.Set
				_ = _526_v64
				_526_v64 = _dafny.SetOf(_522_v60, _525_v63, _dafny.SetOf((func() bool {
					if (_498_v45).Contains(false) {
						return (_498_v45).Get(false).(bool)
					}
					return (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)
				})()))
				var _527_v65 _dafny.Sequence
				_ = _527_v65
				_527_v65 = _dafny.SeqOf(_dafny.MultiSetOf(_525_v63, _522_v60))
				var _528_v67 _dafny.Map
				_ = _528_v67
				_528_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_525_v63, (_this).F6())
				var _529_v69 _dafny.Map
				_ = _529_v69
				_529_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(p0, (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)), _522_v60)
				var _530_v71 _dafny.Sequence
				_ = _530_v71
				_530_v71 = _dafny.SeqOf(_526_v64)
				var _531_v72 _dafny.Map
				_ = _531_v72
				_531_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_525_v63, (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
				var _532_v74 _dafny.Array
				_ = _532_v74
				var _nwElement0_11 _dafny.Set = _526_v64
				_ = _nwElement0_11
				var _nw68 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(27))
				_ = _nw68
				(_nw68).ArraySet1(_nwElement0_11, 0)
				(_nw68).ArraySet1(_526_v64, 1)
				(_nw68).ArraySet1(_526_v64, 2)
				(_nw68).ArraySet1(_526_v64, 3)
				(_nw68).ArraySet1(_526_v64, 4)
				(_nw68).ArraySet1((_526_v64).Intersection(_526_v64), 5)
				(_nw68).ArraySet1(_dafny.SetOf(_525_v63, _525_v63, _522_v60), 6)
				(_nw68).ArraySet1(_526_v64, 7)
				(_nw68).ArraySet1((func() _dafny.Set {
					if p0 {
						return _526_v64
					}
					return _526_v64
				})(), 8)
				(_nw68).ArraySet1(func() _dafny.Set {
					var _coll13 = _dafny.NewBuilder()
					_ = _coll13
					for _iter15 := _dafny.Iterate(((_527_v65).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.IntOfUint32((_527_v65).Cardinality()))).Uint32()).(_dafny.MultiSet)).Elements()); ; {
						_compr_13, _ok15 := _iter15()
						if !_ok15 {
							break
						}
						var _533_v66 _dafny.Set
						_533_v66 = interface{}(_compr_13).(_dafny.Set)
						if ((_527_v65).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.IntOfUint32((_527_v65).Cardinality()))).Uint32()).(_dafny.MultiSet)).Contains(_533_v66) {
							_coll13.Add(_533_v66)
						}
					}
					return _coll13.ToSet()
				}(), 9)
				(_nw68).ArraySet1((_526_v64).Union(_526_v64), 10)
				(_nw68).ArraySet1(func() _dafny.Set {
					var _coll14 = _dafny.NewBuilder()
					_ = _coll14
					for _iter16 := _dafny.Iterate((_528_v67).Keys().Elements()); ; {
						_compr_14, _ok16 := _iter16()
						if !_ok16 {
							break
						}
						var _534_v68 _dafny.Set
						_534_v68 = interface{}(_compr_14).(_dafny.Set)
						if (_528_v67).Contains(_534_v68) {
							_coll14.Add(_534_v68)
						}
					}
					return _coll14.ToSet()
				}(), 11)
				(_nw68).ArraySet1(_526_v64, 12)
				(_nw68).ArraySet1(_dafny.SetOf(_dafny.SetOf((p1).Select((Companion_Default___.SafeIndex((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)), _522_v60), 13)
				(_nw68).ArraySet1(_526_v64, 14)
				(_nw68).ArraySet1((_526_v64).Difference(func() _dafny.Set {
					var _coll15 = _dafny.NewBuilder()
					_ = _coll15
					for _iter17 := _dafny.Iterate((_529_v69).Keys().Elements()); ; {
						_compr_15, _ok17 := _iter17()
						if !_ok17 {
							break
						}
						var _535_v70 _dafny.Set
						_535_v70 = interface{}(_compr_15).(_dafny.Set)
						if (_529_v69).Contains(_535_v70) {
							_coll15.Add(_535_v70)
						}
					}
					return _coll15.ToSet()
				}()), 15)
				(_nw68).ArraySet1((_530_v71).Select((Companion_Default___.SafeIndex((_513_v55).Cardinality(), _dafny.IntOfUint32((_530_v71).Cardinality()))).Uint32()).(_dafny.Set), 16)
				(_nw68).ArraySet1(func() _dafny.Set {
					var _coll16 = _dafny.NewBuilder()
					_ = _coll16
					for _iter18 := _dafny.Iterate((_531_v72).Keys().Elements()); ; {
						_compr_16, _ok18 := _iter18()
						if !_ok18 {
							break
						}
						var _536_v73 _dafny.Set
						_536_v73 = interface{}(_compr_16).(_dafny.Set)
						if (_531_v72).Contains(_536_v73) {
							_coll16.Add(_536_v73)
						}
					}
					return _coll16.ToSet()
				}(), 17)
				(_nw68).ArraySet1((_530_v71).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_530_v71).Cardinality()))).Uint32()).(_dafny.Set), 18)
				(_nw68).ArraySet1((_526_v64).Intersection(_dafny.SetOf(_522_v60, _522_v60)), 19)
				(_nw68).ArraySet1(_dafny.SetOf(_525_v63), 20)
				(_nw68).ArraySet1((_dafny.SetOf(_525_v63, _dafny.SetOf(p0), _525_v63, _522_v60, _522_v60)).Intersection(_526_v64), 21)
				(_nw68).ArraySet1(_526_v64, 22)
				(_nw68).ArraySet1(_526_v64, 23)
				(_nw68).ArraySet1(_526_v64, 24)
				(_nw68).ArraySet1((func() _dafny.Set {
					if p0 {
						return _526_v64
					}
					return _526_v64
				})(), 25)
				(_nw68).ArraySet1(_526_v64, 26)
				_532_v74 = _nw68
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_532_v74), 0))
				_ = _index77
				(_532_v74).ArraySet1(_dafny.SetOf(_525_v63), (_index77).Int())
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_532_v74), 0))
				_ = _index78
				var _rhs80 bool = (p0) && ((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
				_ = _rhs80
				var _rhs81 bool = (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)
				_ = _rhs81
				var _rhs82 bool = (p0) && ((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
				_ = _rhs82
				var _rhs83 _dafny.Int = Companion_Default___.Fm2(((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int)).Minus(p2), globalState)
				_ = _rhs83
				var _rhs84 _dafny.Set = _526_v64
				_ = _rhs84
				var _lhs67 *GlobalState = globalState
				_ = _lhs67
				var _lhs68 *GlobalState = globalState
				_ = _lhs68
				var _lhs69 *GlobalState = globalState
				_ = _lhs69
				var _lhs70 *GlobalState = globalState
				_ = _lhs70
				var _lhs71 _dafny.Array = _532_v74
				_ = _lhs71
				var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_532_v74), 0))
				_ = _lhs72
				_lhs67.F0 = _rhs80
				_lhs68.F0 = _rhs81
				_lhs69.F0 = _rhs82
				_lhs70.F5 = _rhs83
				(_lhs71).ArraySet1(_rhs84, (_lhs72).Int())
			} else {
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))
				_ = _index79
				(_421_v2).ArraySet1((_dafny.Zero).Minus(((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int)).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F6())))), (_index79).Int())
				var _537_v75 _dafny.MultiSet
				_ = _537_v75
				_537_v75 = _dafny.MultiSetOf((_this).F6())
				var _538_v76 _dafny.MultiSet
				_ = _538_v76
				_538_v76 = _dafny.MultiSetOf(_537_v75)
				var _539_v77 *C2
				_ = _539_v77
				var _nw69 *C2 = New_C2_()
				_ = _nw69
				_nw69.Ctor__((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int), (_538_v76).Difference(_538_v76), _this.F9(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _504_v48))
				_539_v77 = _nw69
				var _540_v78 _dafny.Set
				_ = _540_v78
				_540_v78 = _dafny.SetOf(p0)
				var _541_v79 D5
				_ = _541_v79
				_541_v79 = Companion_D5_.Create_DC16_(_540_v78)
				_541_v79 = _541_v79
				(globalState).F5 = (_539_v77).Fm21(_dafny.IntOfUint32((p1).Cardinality()), globalState)
				_421_v2 = _421_v2
			}
			(globalState).F5 = (_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int)
		}
		var _542_v80 D5
		_ = _542_v80
		_542_v80 = Companion_D5_.Create_DC17_(_417_v0, p0)
		var _source8 D5 = _542_v80
		_ = _source8
		if _source8.Is_DC17() {
			var _543___mcc_h9 _dafny.Sequence = _source8.Get_().(D5_DC17).Cf36
			_ = _543___mcc_h9
			var _544___mcc_h10 bool = _source8.Get_().(D5_DC17).Cf37
			_ = _544___mcc_h10
			var _545_cf37 bool = _544___mcc_h10
			_ = _545_cf37
			var _546_cf36 _dafny.Sequence = _543___mcc_h9
			_ = _546_cf36
			var _547_v81 _dafny.Map
			_ = _547_v81
			_547_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int))
			var _548_v82 _dafny.Map
			_ = _548_v82
			_548_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_545_cf37, _545_cf37)
			var _549_v83 _dafny.Map
			_ = _549_v83
			_549_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm37(_545_cf37, globalState)).Dtor_cf41(), (func() bool {
				if (_548_v82).Contains(p0) {
					return (_548_v82).Get(p0).(bool)
				}
				return _545_cf37
			})())
			_547_v81 = (_547_v81).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, p1)).Cardinality()), (_549_v83).Cardinality())
			var _550_v84 *C1
			_ = _550_v84
			var _nw70 *C1 = New_C1_()
			_ = _nw70
			_nw70.Ctor__(p2)
			_550_v84 = _nw70
			var _551_v85 _dafny.Sequence
			_ = _551_v85
			_551_v85 = _dafny.SeqOf(_550_v84, _550_v84)
			var _552_v86 _dafny.Map
			_ = _552_v86
			_552_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_551_v85).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p1).Cardinality()), _dafny.IntOfUint32((_551_v85).Cardinality()))).Uint32()).(*C1), _545_cf37)
			var _arr9 _dafny.Array = _this.F9()
			_ = _arr9
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index80
			(_arr9).ArraySet1((func() bool {
				if (_552_v86).Contains(_550_v84) {
					return (_552_v86).Get(_550_v84).(bool)
				}
				return _545_cf37
			})(), (_index80).Int())
			var _553_v87 _dafny.CodePoint
			_ = _553_v87
			_553_v87 = _dafny.CodePoint('r')
			var _554_v88 _dafny.MultiSet
			_ = _554_v88
			_554_v88 = _dafny.MultiSetOf(_553_v87)
			var _555_v89 _dafny.Map
			_ = _555_v89
			_555_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_545_cf37, _554_v88)
			var _556_v90 _dafny.Sequence
			_ = _556_v90
			_556_v90 = _dafny.SeqOf(_554_v88, _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('v'), _553_v87, _553_v87)), _554_v88, (func() _dafny.MultiSet {
				if (_555_v89).Contains(_545_cf37) {
					return (_555_v89).Get(_545_cf37).(_dafny.MultiSet)
				}
				return (_554_v88).Update(_dafny.CodePoint('c'), Companion_Default___.Abs(p2))
			})())
			var _557_v91 D4
			_ = _557_v91
			_557_v91 = Companion_D4_.Create_DC12_(_556_v90)
			var _arr10 _dafny.Array = _this.F9()
			_ = _arr10
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index81
			var _rhs85 bool = false
			_ = _rhs85
			var _rhs86 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p1), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()))).Uint32(), _dafny.SeqOf(p0, !(_545_cf37), true, p0)), _dafny.SeqOf(p1))).Cardinality())).Cmp((_dafny.IntOfInt64(182)).Minus(Companion_Default___.Fm2(_dafny.IntOfInt64(565), globalState))) > 0
			_ = _rhs86
			var _rhs87 D4 = _557_v91
			_ = _rhs87
			var _lhs73 _dafny.Array = _this.F9()
			_ = _lhs73
			var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_this.F9()), 0))
			_ = _lhs74
			(_lhs73).ArraySet1(_rhs85, (_lhs74).Int())
			_545_cf37 = _rhs86
			_557_v91 = _rhs87
			var _558_v92 _dafny.MultiSet
			_ = _558_v92
			_558_v92 = _dafny.MultiSetOf(_dafny.IntOfInt64(-930))
			var _559_v93 _dafny.Array
			_ = _559_v93
			var _nw71 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(10))
			_ = _nw71
			_559_v93 = _nw71
			var _560_v94 D0
			_ = _560_v94
			_560_v94 = Companion_D0_.Create_DC1_(_545_cf37, (_this).F6(), _553_v87)
			var _561_v95 _dafny.Sequence
			_ = _561_v95
			_561_v95 = _dafny.UnicodeSeqOfUtf8Bytes("vst")
			var _562_v96 _dafny.Array
			_ = _562_v96
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_9
			var _nw72 _dafny.Array
			_ = _nw72
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw72 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) D0 = (func(_563_p1 _dafny.Sequence) func(_dafny.Int) D0 {
					return func(_564_i11 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_(_563_p1, (_this).F6())
					}
				})(p1)
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
			_562_v96 = _nw72
			var _565_v97 D7
			_ = _565_v97
			_565_v97 = Companion_D7_.Create_DC21_(_562_v96)
			var _566_v98 _dafny.Set
			_ = _566_v98
			_566_v98 = _dafny.SetOf((_565_v97).Dtor_cf43())
			var _rhs88 bool = (func() bool {
				if _545_cf37 {
					return (_560_v94).Dtor_cf2()
				}
				return ((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)) == (_545_cf37)
			})()
			_ = _rhs88
			var _rhs89 _dafny.MultiSet = Companion_Default___.Fm35((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int), ((_this).F6()).Times((_dafny.Zero).Minus((_this).F6())), false, p2, globalState)
			_ = _rhs89
			var _rhs90 _dafny.Array = _421_v2
			_ = _rhs90
			var _rhs91 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(928))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg48 _dafny.Int) interface{} {
					return coer48(arg48)
				}
			}((func(_567_v87 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
				return func(_568_i9 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(921))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg49 _dafny.Int) interface{} {
							return coer49(arg49)
						}
					}(func(_569_i10 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('a')
					})), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(921))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg50 _dafny.Int) interface{} {
							return coer50(arg50)
						}
					}(func(_570_i10 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('a')
					}))).Cardinality()))).Uint32(), _567_v87)
				}
			})(_553_v87))), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-96), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(928))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}((func(_571_v87 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
				return func(_572_i9 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(921))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg52 _dafny.Int) interface{} {
							return coer52(arg52)
						}
					}(func(_573_i10 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('a')
					})), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(921))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg53 _dafny.Int) interface{} {
							return coer53(arg53)
						}
					}(func(_574_i10 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('a')
					}))).Cardinality()))).Uint32(), _571_v87)
				}
			})(_553_v87)))).Cardinality()))).Uint32(), _561_v95)).Cardinality())
			_ = _rhs91
			var _rhs92 _dafny.Array = (func() _dafny.Array {
				if !(_566_v98).Equals(_566_v98) {
					return _559_v93
				}
				return _559_v93
			})()
			_ = _rhs92
			var _lhs75 *GlobalState = globalState
			_ = _lhs75
			_545_cf37 = _rhs88
			_558_v92 = _rhs89
			_421_v2 = _rhs90
			_lhs75.F5 = _rhs91
			_559_v93 = _rhs92
			var _arr11 _dafny.Array = _this.F9()
			_ = _arr11
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index82
			(_arr11).ArraySet1(p0, (_index82).Int())
		} else {
			var _575___mcc_h11 _dafny.Set = _source8.Get_().(D5_DC16).Cf35
			_ = _575___mcc_h11
			var _576_cf35 _dafny.Set = _575___mcc_h11
			_ = _576_cf35
			var _577_v99 _dafny.Map
			_ = _577_v99
			_577_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(p1)).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(416))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg54 _dafny.Int) interface{} {
					return coer54(arg54)
				}
			}(func(_578_i14 _dafny.Int) _dafny.Int {
				return (_this).F6()
			})))
			var _579_v100 _dafny.MultiSet
			_ = _579_v100
			_579_v100 = _dafny.MultiSetOf(p0, true)
			var _580_v101 _dafny.CodePoint
			_ = _580_v101
			_580_v101 = _dafny.CodePoint('h')
			var _581_v102 _dafny.MultiSet
			_ = _581_v102
			_581_v102 = _dafny.MultiSetOf(_dafny.CodePoint('u'), _580_v101)
			var _582_v103 _dafny.Map
			_ = _582_v103
			_582_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _dafny.Companion_Sequence_.Update(_417_v0, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.IntOfUint32((_417_v0).Cardinality()))).Uint32(), (func() _dafny.Int {
				if (_579_v100).Contains(p0) {
					return (_579_v100).Multiplicity(p0)
				}
				return (_581_v102).Cardinality()
			})()))
			var _583_v104 _dafny.MultiSet
			_ = _583_v104
			_583_v104 = _dafny.MultiSetOf((_this).F6(), (_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int))
			var _584_v105 _dafny.Sequence
			_ = _584_v105
			_584_v105 = _dafny.SeqOf(_417_v0)
			var _585_v106 _dafny.Sequence
			_ = _585_v106
			_585_v106 = _dafny.UnicodeSeqOfUtf8Bytes("nhw")
			var _586_v107 _dafny.Sequence
			_ = _586_v107
			_586_v107 = _dafny.SeqOf(_417_v0, _417_v0, (_584_v105).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm2(_dafny.IntOfUint32((_585_v106).Cardinality()), globalState), _dafny.IntOfUint32((_584_v105).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _587_v108 _dafny.Array
			_ = _587_v108
			var _nwElement0_12 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(49))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg55 _dafny.Int) interface{} {
					return coer55(arg55)
				}
			}((func(_588_v2 _dafny.Array) func(_dafny.Int) _dafny.Int {
				return func(_589_i12 _dafny.Int) _dafny.Int {
					return (_588_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_588_v2), 0))).Int()).(_dafny.Int)
				}
			})(_421_v2))), _417_v0)
			_ = _nwElement0_12
			var _nw73 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(14))
			_ = _nw73
			(_nw73).ArraySet1(_nwElement0_12, 0)
			(_nw73).ArraySet1(_417_v0, 1)
			(_nw73).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_417_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(332))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg56 _dafny.Int) interface{} {
					return coer56(arg56)
				}
			}((func(_590_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_591_i13 _dafny.Int) _dafny.Int {
					return _590_p2
				}
			})(p2)))), 2)
			(_nw73).ArraySet1((func() _dafny.Sequence {
				if (_577_v99).Contains(_dafny.IntOfUint32((_417_v0).Cardinality())) {
					return (_577_v99).Get(_dafny.IntOfUint32((_417_v0).Cardinality())).(_dafny.Sequence)
				}
				return _dafny.SeqOf(p2, (_this).F6())
			})(), 3)
			(_nw73).ArraySet1(_dafny.Companion_Sequence_.Update(_417_v0, (Companion_Default___.SafeIndex((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_417_v0).Cardinality()))).Uint32(), Companion_Default___.Fm2((_this).F6(), globalState)), 4)
			(_nw73).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_417_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(640))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg57 _dafny.Int) interface{} {
					return coer57(arg57)
				}
			}(func(_592_i15 _dafny.Int) _dafny.Int {
				return (_this).F6()
			}))), 5)
			(_nw73).ArraySet1(_417_v0, 6)
			(_nw73).ArraySet1(_417_v0, 7)
			(_nw73).ArraySet1((func() _dafny.Sequence {
				if (_582_v103).Contains(p0) {
					return (_582_v103).Get(p0).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Update(_417_v0, (Companion_Default___.SafeIndex((_583_v104).Cardinality(), _dafny.IntOfUint32((_417_v0).Cardinality()))).Uint32(), (_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int))
			})(), 8)
			(_nw73).ArraySet1(_417_v0, 9)
			(_nw73).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_417_v0, _417_v0), 10)
			(_nw73).ArraySet1((_584_v105).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_584_v105).Cardinality()))).Uint32()).(_dafny.Sequence), 11)
			(_nw73).ArraySet1((func() _dafny.Sequence {
				if !(!(!(p0))) {
					return _417_v0
				}
				return _417_v0
			})(), 12)
			(_nw73).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int), (_this).F6(), (_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(953), p2), _dafny.SeqOf((_421_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_421_v2), 0))).Int()).(_dafny.Int))), 13)
			_587_v108 = _nw73
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_10
			var _nw74 _dafny.Array
			_ = _nw74
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw74 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Sequence = (func(_593_v0 _dafny.Sequence, _594_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_595_i16 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_593_v0, _dafny.Companion_Sequence_.Update(_593_v0, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_593_v0).Cardinality()))).Uint32(), _dafny.IntOfUint32((_594_p1).Cardinality())))
					}
				})(_417_v0, p1)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw74 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw74).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw74).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_587_v108 = _nw74
			var _596_v109 D5
			_ = _596_v109
			_596_v109 = Companion_D5_.Create_DC16_(_576_cf35)
			var _597_v110 D5
			_ = _597_v110
			_597_v110 = Companion_D5_.Create_DC16_((_596_v109).Dtor_cf35())
			var _pat_let_tv16 = globalState
			_ = _pat_let_tv16
			var _pat_let_tv17 = _576_cf35
			_ = _pat_let_tv17
			var _rhs93 _dafny.Int = p2
			_ = _rhs93
			var _rhs94 D5 = func(_pat_let21_0 D5) D5 {
				return func(_598_dt__update__tmp_h4 D5) D5 {
					return func(_pat_let22_0 _dafny.Set) D5 {
						return func(_599_dt__update_hcf35_h0 _dafny.Set) D5 {
							return Companion_D5_.Create_DC16_(_599_dt__update_hcf35_h0)
						}(_pat_let22_0)
					}(((Companion_Default___.Fm38(_pat_let_tv16)).Dtor_cf35()).Union(_pat_let_tv17))
				}(_pat_let21_0)
			}(_597_v110)
			_ = _rhs94
			var _lhs76 *GlobalState = globalState
			_ = _lhs76
			_lhs76.F5 = _rhs93
			_596_v109 = _rhs94
			(globalState).F5 = (func() _dafny.Int {
				if p0 {
					return (_this).F6()
				}
				return p2
			})()
			var _600_v111 _dafny.Sequence
			_ = _600_v111
			_600_v111 = _dafny.SeqOf(_585_v106, _585_v106)
			(globalState).F5 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_600_v111, _dafny.SeqOf(_585_v106, _dafny.UnicodeSeqOfUtf8Bytes("f")))).Cardinality()), _dafny.IntOfUint32((_585_v106).Cardinality()))
		}
		var _601_v112 _dafny.Set
		_ = _601_v112
		_601_v112 = _dafny.SetOf(_dafny.IntOfInt64(-87))
		r0 = Companion_Default___.Fm29(_601_v112, globalState)
		r1 = p0
		return r0, r1
	}
}
func (_this *C3) M7(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) (_dafny.CodePoint, _dafny.Int, _dafny.Array, _dafny.Int) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _602_v0 _dafny.CodePoint
		_ = _602_v0
		_602_v0 = _dafny.CodePoint('j')
		r0 = _602_v0
		(_this).F9_set_((_this.F9()))
		var _603_v1 _dafny.Array
		_ = _603_v1
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_11
		var _nw75 _dafny.Array
		_ = _nw75
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw75 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Int = func(_604_i0 _dafny.Int) _dafny.Int {
				return (_604_i0).Plus((_this).F6())
			}
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
		_603_v1 = _nw75
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_603_v1), 0))
		_ = _index83
		(_603_v1).ArraySet1((_this).F6(), (_index83).Int())
		var _605_v2 bool
		_ = _605_v2
		_605_v2 = false
		var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_603_v1), 0))
		_ = _index84
		var _rhs95 _dafny.Int = _dafny.IntOfInt64(-180)
		_ = _rhs95
		var _rhs96 _dafny.Int = _dafny.IntOfInt64(84)
		_ = _rhs96
		var _rhs97 bool = _605_v2
		_ = _rhs97
		var _lhs77 _dafny.Array = _603_v1
		_ = _lhs77
		var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_603_v1), 0))
		_ = _lhs78
		var _lhs79 *GlobalState = globalState
		_ = _lhs79
		r3 = _rhs95
		(_lhs77).ArraySet1(_rhs96, (_lhs78).Int())
		_lhs79.F0 = _rhs97
		(globalState).F0 = _605_v2
		var _606_v3 _dafny.Map
		_ = _606_v3
		_606_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v2, p1)
		var _607_v5 _dafny.Sequence
		_ = _607_v5
		_607_v5 = _dafny.SeqOf((_603_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_603_v1), 0))).Int()).(_dafny.Int), p1, (_603_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_603_v1), 0))).Int()).(_dafny.Int))
		if ((Companion_Default___.Fm1(globalState)).Cardinality()).Cmp(((_dafny.SetOf(_606_v3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_605_v2), (func() _dafny.Map {
			var _coll17 = _dafny.NewMapBuilder()
			_ = _coll17
			for _iter19 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_607_v5, _605_v2)).Keys().Elements()); ; {
				_compr_17, _ok19 := _iter19()
				if !_ok19 {
					break
				}
				var _608_v4 _dafny.Sequence
				_608_v4 = interface{}(_compr_17).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_607_v5, _605_v2)).Contains(_608_v4) {
					_coll17.Add(_608_v4, _605_v2)
				}
			}
			return _coll17.ToMap()
		}()).Cardinality()))).Union(_dafny.SetOf(_606_v3, _606_v3))).Cardinality()) <= 0 {
			var _609_v6 *C0
			_ = _609_v6
			var _nw76 *C0 = New_C0_()
			_ = _nw76
			_nw76.Ctor__(_this.F9(), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("gxtawrg"), _dafny.UnicodeSeqOfUtf8Bytes("qi")))
			_609_v6 = _nw76
			(globalState).F0 = Companion_Default___.Fm3(globalState)
			var _610_v7 _dafny.Sequence
			_ = _610_v7
			_610_v7 = _dafny.SeqOf(_609_v6.F20)
			var _611_v8 D0
			_ = _611_v8
			_611_v8 = Companion_D0_.Create_DC0_(_610_v7, (_603_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_603_v1), 0))).Int()).(_dafny.Int))
			var _612_v9 _dafny.Sequence
			_ = _612_v9
			_612_v9 = _dafny.UnicodeSeqOfUtf8Bytes("fsmbn")
			var _613_v10 D1
			_ = _613_v10
			_613_v10 = Companion_D1_.Create_DC4_(_611_v8, p1, _612_v9, _605_v2, (_this).F6())
			var _614_v11 _dafny.Int
			_ = _614_v11
			var _615_v12 _dafny.CodePoint
			_ = _615_v12
			var _616_v13 bool
			_ = _616_v13
			var _out19 _dafny.Int
			_ = _out19
			var _out20 _dafny.CodePoint
			_ = _out20
			var _out21 bool
			_ = _out21
			_out19, _out20, _out21 = (_this).M14((_613_v10).Dtor_cf10(), _609_v6.F20, globalState)
			_614_v11 = _out19
			_615_v12 = _out20
			_616_v13 = _out21
			var _617_v14 _dafny.Array
			_ = _617_v14
			var _nw77 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
			_ = _nw77
			_617_v14 = _nw77
			var _618_v15 _dafny.Map
			_ = _618_v15
			_618_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_617_v14, _605_v2)
			_618_v15 = (_618_v15).Update(_617_v14, Companion_Default___.Fm3(globalState))
			var _619_v16 _dafny.MultiSet
			_ = _619_v16
			_619_v16 = _dafny.MultiSetOf(_603_v1)
			(_609_v6).F20 = (_619_v16).IsProperSubsetOf((_dafny.MultiSetOf(_603_v1)).Intersection(_619_v16))
		} else {
			var _620_v17 _dafny.Array
			_ = _620_v17
			var _nw78 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(13))
			_ = _nw78
			_620_v17 = _nw78
			var _nw79 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(25))
			_ = _nw79
			_620_v17 = _nw79
			var _621_v18 _dafny.MultiSet
			_ = _621_v18
			_621_v18 = _dafny.MultiSetOf((p0).Update((_603_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_603_v1), 0))).Int()).(_dafny.Int), Companion_Default___.Abs((_this).F6())))
			var _622_v19 *C2
			_ = _622_v19
			var _nw80 *C2 = New_C2_()
			_ = _nw80
			_nw80.Ctor__((_603_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_603_v1), 0))).Int()).(_dafny.Int), (_621_v18).Update(p0, Companion_Default___.Abs((_603_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_603_v1), 0))).Int()).(_dafny.Int))), _this.F9(), Companion_Default___.Fm39(_602_v0, globalState))
			_622_v19 = _nw80
			var _623_v20 _dafny.Set
			_ = _623_v20
			_623_v20 = _dafny.SetOf((_603_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_603_v1), 0))).Int()).(_dafny.Int), (_this).F6(), (p1).Plus(Companion_Default___.Fm2((_this).F6(), globalState)))
			_623_v20 = _623_v20
			var _624_v21 _dafny.Array
			_ = _624_v21
			var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
			_ = _nw81
			_624_v21 = _nw81
			var _625_v22 _dafny.Array
			_ = _625_v22
			var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(2))
			_ = _nw82
			_625_v22 = _nw82
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_624_v21), 0))
			_ = _index85
			(_624_v21).ArraySet1(_625_v22, (_index85).Int())
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_624_v21), 0))
			_ = _index86
			(_624_v21).ArraySet1(_625_v22, (_index86).Int())
			if !(_605_v2) || (_605_v2) {
				var _626_v23 _dafny.Int
				_ = _626_v23
				var _627_v24 _dafny.CodePoint
				_ = _627_v24
				var _628_v25 bool
				_ = _628_v25
				var _out22 _dafny.Int
				_ = _out22
				var _out23 _dafny.CodePoint
				_ = _out23
				var _out24 bool
				_ = _out24
				_out22, _out23, _out24 = (_this).M14(true, _605_v2, globalState)
				_626_v23 = _out22
				_627_v24 = _out23
				_628_v25 = _out24
				var _629_v26 _dafny.Array
				_ = _629_v26
				_629_v26 = _this.F9()
				_629_v26 = _629_v26
				var _630_v27 _dafny.Map
				_ = _630_v27
				_630_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
				_605_v2 = (_626_v23).Cmp((func() _dafny.Int {
					if (_630_v27).Contains(p0) {
						return (_630_v27).Get(p0).(_dafny.Int)
					}
					return (_this).F6()
				})()) != 0
				var _631_v28 _dafny.Sequence
				_ = _631_v28
				_631_v28 = _dafny.SeqOf(_628_v25)
				var _632_v29 _dafny.Map
				_ = _632_v29
				_632_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_628_v25, _628_v25)
				var _633_v30 _dafny.Map
				_ = _633_v30
				_633_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_631_v28).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.IntOfUint32((_631_v28).Cardinality()))).Uint32()).(bool)), (func() bool {
					if (_632_v29).Contains(_628_v25) {
						return (_632_v29).Get(_628_v25).(bool)
					}
					return true
				})())
				r1 = ((_dafny.Zero).Minus((func() _dafny.Int {
					if (_606_v3).Contains(_605_v2) {
						return (_606_v3).Get(_605_v2).(_dafny.Int)
					}
					return (_633_v30).Cardinality()
				})())).Plus(_dafny.IntOfInt64(116))
				_623_v20 = _623_v20
			} else {
				var _634_v31 D9
				_ = _634_v31
				_634_v31 = Companion_D9_.Create_DC24_(_621_v18)
				var _635_v32 *C2
				_ = _635_v32
				var _nw83 *C2 = New_C2_()
				_ = _nw83
				_nw83.Ctor__(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ytxvws")).Cardinality()), (_621_v18).Union((_634_v31).Dtor_cf45()), _this.F9(), (_this).F10())
				_635_v32 = _nw83
				var _636_v33 _dafny.Array
				_ = _636_v33
				var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
				_ = _nw84
				_636_v33 = _nw84
				_636_v33 = _636_v33
				_603_v1 = _603_v1
				var _637_v34 _dafny.Set
				_ = _637_v34
				_637_v34 = _dafny.SetOf(_605_v2, _605_v2, true)
				(globalState).F5 = (_637_v34).Cardinality()
				var _638_v35 *C0
				_ = _638_v35
				var _nw85 *C0 = New_C0_()
				_ = _nw85
				_nw85.Ctor__(_this.F9(), !(_605_v2))
				_638_v35 = _nw85
			}
		}
		var _639_v36 _dafny.Array
		_ = _639_v36
		var _nwElement0_13 _dafny.Map = (_this).F10()
		_ = _nwElement0_13
		var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(6))
		_ = _nw86
		(_nw86).ArraySet1(_nwElement0_13, 0)
		(_nw86).ArraySet1((_this).F10(), 1)
		(_nw86).ArraySet1(((_this).F10()).Merge((_this).F10()), 2)
		(_nw86).ArraySet1((_this).F10(), 3)
		(_nw86).ArraySet1((_this).F10(), 4)
		(_nw86).ArraySet1((_this).F10(), 5)
		_639_v36 = _nw86
		var _640_v37 _dafny.Sequence
		_ = _640_v37
		_640_v37 = _dafny.UnicodeSeqOfUtf8Bytes("rtg")
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_639_v36), 0))
		_ = _index87
		(_639_v36).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_603_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_603_v1), 0))).Int()).(_dafny.Int), _640_v37), (_index87).Int())
		var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_639_v36), 0))
		_ = _index88
		(_639_v36).ArraySet1((_this).F10(), (_index88).Int())
		r0 = _602_v0
		var _641_v38 _dafny.Sequence
		_ = _641_v38
		_641_v38 = _dafny.SeqOf(_605_v2, Companion_Default___.Fm3(globalState))
		r1 = ((_603_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_603_v1), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfUint32((_641_v38).Cardinality()))
		var _nwElement0_14 bool = _605_v2
		_ = _nwElement0_14
		var _nw87 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(7))
		_ = _nw87
		(_nw87).ArraySet1(_nwElement0_14, 0)
		(_nw87).ArraySet1(false, 1)
		(_nw87).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_641_v38, _dafny.Companion_Sequence_.Update(_641_v38, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_641_v38).Cardinality()))).Uint32(), Companion_Default___.Fm3(globalState))), 2)
		(_nw87).ArraySet1((_dafny.IntOfInt64(-757)).Cmp(_dafny.IntOfUint32((_607_v5).Cardinality())) == 0, 3)
		(_nw87).ArraySet1((func() bool {
			if _605_v2 {
				return _605_v2
			}
			return _605_v2
		})(), 4)
		(_nw87).ArraySet1(_605_v2, 5)
		(_nw87).ArraySet1(!(_605_v2) || (true), 6)
		r2 = _nw87
		r3 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_640_v37).Cardinality()), (_this).F6()))
		return r0, r1, r2, r3
	}
}
func (_this *C3) M14(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.CodePoint, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r1
		var r2 bool = false
		_ = r2
		if p1 {
			r0 = (_this).F6()
			(globalState).F5 = (_this).F6()
			var _642_v0 *C0
			_ = _642_v0
			var _nw88 *C0 = New_C0_()
			_ = _nw88
			_nw88.Ctor__(_this.F9(), p0)
			_642_v0 = _nw88
			var _643_v1 _dafny.Sequence
			_ = _643_v1
			_643_v1 = _dafny.SeqOf(((_this).F6()).Minus((_this).F6()), (_this).F6())
			_643_v1 = _dafny.Companion_Sequence_.Concatenate(_643_v1, _643_v1)
			var _644_v2 _dafny.Map
			_ = _644_v2
			_644_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p0)
			_644_v2 = (_644_v2).Update(_dafny.IntOfInt64(-864), _642_v0.F20)
		} else {
			var _rhs98 bool = p0
			_ = _rhs98
			var _rhs99 _dafny.Int = (_this).F6()
			_ = _rhs99
			var _lhs80 *GlobalState = globalState
			_ = _lhs80
			var _lhs81 *GlobalState = globalState
			_ = _lhs81
			_lhs80.F0 = _rhs98
			_lhs81.F5 = _rhs99
			var _arr12 _dafny.Array = _this.F9()
			_ = _arr12
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index89
			(_arr12).ArraySet1(p0, (_index89).Int())
			var _arr13 _dafny.Array = _this.F9()
			_ = _arr13
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index90
			(_arr13).ArraySet1(p0, (_index90).Int())
			(globalState).F5 = (func() _dafny.Int {
				if p1 {
					return (_this).F6()
				}
				return (_this).F6()
			})()
			if !(p0) {
				var _645_v3 _dafny.Sequence
				_ = _645_v3
				_645_v3 = _dafny.SeqOf(p1)
				var _646_v4 D0
				_ = _646_v4
				_646_v4 = Companion_D0_.Create_DC0_(_645_v3, (_this).F6())
				var _647_v5 _dafny.CodePoint
				_ = _647_v5
				_647_v5 = _dafny.CodePoint('v')
				var _648_v6 D1
				_ = _648_v6
				_648_v6 = Companion_D1_.Create_DC4_(_646_v4, (_this).F6(), Companion_Default___.Fm30(p0, _647_v5, globalState), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), (_this).F6())
				(globalState).F2 = (_648_v6).Dtor_cf9()
				var _649_v7 _dafny.Set
				_ = _649_v7
				_649_v7 = _dafny.SetOf(_647_v5, _647_v5, _647_v5)
				var _650_v8 _dafny.MultiSet
				_ = _650_v8
				_650_v8 = _dafny.MultiSetOf(_649_v7, _649_v7, _649_v7)
				var _651_v9 _dafny.Set
				_ = _651_v9
				_651_v9 = _dafny.SetOf((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), p0)
				var _652_v10 _dafny.Map
				_ = _652_v10
				_652_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_651_v9).Cardinality(), p0)
				var _653_v11 _dafny.Map
				_ = _653_v11
				_653_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_652_v10, p1)
				var _654_v12 D3
				_ = _654_v12
				_654_v12 = Companion_D3_.Create_DC10_(_653_v11)
				var _655_v13 _dafny.Sequence
				_ = _655_v13
				_655_v13 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (Companion_Default___.Fm33(_654_v12, (_this).F6(), false, p1, globalState)).Cardinality())).Cardinality())
				var _656_v14 _dafny.Sequence
				_ = _656_v14
				_656_v14 = _dafny.UnicodeSeqOfUtf8Bytes("b")
				var _657_v15 _dafny.Array
				_ = _657_v15
				var _nwElement0_15 _dafny.Int = (_this).F6()
				_ = _nwElement0_15
				var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(22))
				_ = _nw89
				(_nw89).ArraySet1(_nwElement0_15, 0)
				(_nw89).ArraySet1((_this).F6(), 1)
				(_nw89).ArraySet1((_dafny.Zero).Minus((_this).F6()), 2)
				(_nw89).ArraySet1((_this).F6(), 3)
				(_nw89).ArraySet1((_this).F6(), 4)
				(_nw89).ArraySet1((_this).F6(), 5)
				(_nw89).ArraySet1((_this).F6(), 6)
				(_nw89).ArraySet1(((_this).F6()).Minus((_this).F6()), 7)
				(_nw89).ArraySet1(_dafny.IntOfInt64(247), 8)
				(_nw89).ArraySet1((_this).F6(), 9)
				(_nw89).ArraySet1(((_this).F6()).Times(_dafny.IntOfInt64(558)), 10)
				(_nw89).ArraySet1(((_this).F6()).Minus(Companion_Default___.Fm2((_dafny.Zero).Minus((_this).F6()), globalState)), 11)
				(_nw89).ArraySet1((_this).F6(), 12)
				(_nw89).ArraySet1((_this).F6(), 13)
				(_nw89).ArraySet1((_this).F6(), 14)
				(_nw89).ArraySet1((_this).F6(), 15)
				(_nw89).ArraySet1((_650_v8).Cardinality(), 16)
				(_nw89).ArraySet1(((_this).F6()).Plus(_dafny.IntOfUint32((_655_v13).Cardinality())), 17)
				(_nw89).ArraySet1((_this).F6(), 18)
				(_nw89).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_656_v14, _656_v14)).Cardinality()), 19)
				(_nw89).ArraySet1((_this).F6(), 20)
				(_nw89).ArraySet1((_this).F6(), 21)
				_657_v15 = _nw89
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_657_v15), 0))
				_ = _index91
				(_657_v15).ArraySet1((_this).F6(), (_index91).Int())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_657_v15), 0))
				_ = _index92
				(_657_v15).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_647_v5)).Cardinality()), (_index92).Int())
				var _658_v16 _dafny.Array
				_ = _658_v16
				var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
				_ = _nw90
				_658_v16 = _nw90
				var _659_v17 _dafny.Sequence
				_ = _659_v17
				_659_v17 = _dafny.SeqOf(_658_v16, _658_v16)
				var _660_v18 D9
				_ = _660_v18
				_660_v18 = Companion_D9_.Create_DC26_((_659_v17).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_659_v17).Cardinality()))).Uint32()).(_dafny.Array))
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_657_v15), 0))
				_ = _index93
				var _arr14 _dafny.Array = _this.F9()
				_ = _arr14
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index94
				var _rhs100 D9 = _660_v18
				_ = _rhs100
				var _rhs101 _dafny.Int = Companion_Default___.SafeModuloInt((_657_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_657_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_655_v13, _655_v13)).Cardinality()))
				_ = _rhs101
				var _rhs102 bool = p0
				_ = _rhs102
				var _lhs82 _dafny.Array = _657_v15
				_ = _lhs82
				var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_657_v15), 0))
				_ = _lhs83
				var _lhs84 _dafny.Array = _this.F9()
				_ = _lhs84
				var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_this.F9()), 0))
				_ = _lhs85
				_660_v18 = _rhs100
				(_lhs82).ArraySet1(_rhs101, (_lhs83).Int())
				(_lhs84).ArraySet1(_rhs102, (_lhs85).Int())
				var _661_v19 _dafny.Set
				_ = _661_v19
				_661_v19 = _dafny.SetOf(Companion_Default___.SafeDivisionInt((_657_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_657_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_656_v14).Cardinality())))
				_661_v19 = func() _dafny.Set {
					var _coll18 = _dafny.NewBuilder()
					_ = _coll18
					for _iter20 := _dafny.Iterate((_655_v13).Elements()); ; {
						_compr_18, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _662_v20 _dafny.Int
						_662_v20 = interface{}(_compr_18).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_655_v13, _662_v20) {
							_coll18.Add(Companion_Default___.SafeDivisionInt(_662_v20, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfInt64(-986)), Companion_D4_.Create_DC12_(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('n'), _dafny.CodePoint('h')), _dafny.MultiSetOf(_dafny.CodePoint('c'), _dafny.CodePoint('m')))))).Cardinality()).Times((_dafny.MultiSetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.SetOf(!(!(!(false))))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ehssnh")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-954))).Cardinality()))).Cardinality())).Cardinality())))
						}
					}
					return _coll18.ToSet()
				}()
				var _663_v21 _dafny.Array
				_ = _663_v21
				var _nwElement0_16 _dafny.CodePoint = _647_v5
				_ = _nwElement0_16
				var _nw91 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(28))
				_ = _nw91
				(_nw91).ArraySet1CodePoint(_nwElement0_16, 0)
				(_nw91).ArraySet1CodePoint(_647_v5, 1)
				(_nw91).ArraySet1CodePoint(_647_v5, 2)
				(_nw91).ArraySet1CodePoint(_647_v5, 3)
				(_nw91).ArraySet1CodePoint(_dafny.CodePoint('j'), 4)
				(_nw91).ArraySet1CodePoint(_647_v5, 5)
				(_nw91).ArraySet1CodePoint(_647_v5, 6)
				(_nw91).ArraySet1CodePoint(_647_v5, 7)
				(_nw91).ArraySet1CodePoint(_647_v5, 8)
				(_nw91).ArraySet1CodePoint(_647_v5, 9)
				(_nw91).ArraySet1CodePoint(_dafny.CodePoint('p'), 10)
				(_nw91).ArraySet1CodePoint(_647_v5, 11)
				(_nw91).ArraySet1CodePoint(_647_v5, 12)
				(_nw91).ArraySet1CodePoint((Companion_D0_.Create_DC1_(p1, (_this).F6(), _647_v5)).Dtor_cf4(), 13)
				(_nw91).ArraySet1CodePoint(_647_v5, 14)
				(_nw91).ArraySet1CodePoint(_647_v5, 15)
				(_nw91).ArraySet1CodePoint(_647_v5, 16)
				(_nw91).ArraySet1CodePoint(_647_v5, 17)
				(_nw91).ArraySet1CodePoint(_647_v5, 18)
				(_nw91).ArraySet1CodePoint(_647_v5, 19)
				(_nw91).ArraySet1CodePoint(_647_v5, 20)
				(_nw91).ArraySet1CodePoint(_647_v5, 21)
				(_nw91).ArraySet1CodePoint(_647_v5, 22)
				(_nw91).ArraySet1CodePoint(_647_v5, 23)
				(_nw91).ArraySet1CodePoint(_647_v5, 24)
				(_nw91).ArraySet1CodePoint(_647_v5, 25)
				(_nw91).ArraySet1CodePoint(_dafny.CodePoint('r'), 26)
				(_nw91).ArraySet1CodePoint(_647_v5, 27)
				_663_v21 = _nw91
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_663_v21), 0))
				_ = _index95
				(_663_v21).ArraySet1CodePoint(_647_v5, (_index95).Int())
				var _664_v22 _dafny.Map
				_ = _664_v22
				_664_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)), p1)
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_663_v21), 0))
				_ = _index96
				(_663_v21).ArraySet1CodePoint(Companion_Default___.Fm0(_dafny.IntOfUint32((_645_v3).Cardinality()), _dafny.SetOf((_this).F6(), _dafny.IntOfUint32((_645_v3).Cardinality())), (func() bool {
					if (_664_v22).Contains(p0) {
						return (_664_v22).Get(p0).(bool)
					}
					return (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)
				})(), globalState), (_index96).Int())
			} else {
				var _665_v23 _dafny.MultiSet
				_ = _665_v23
				_665_v23 = _dafny.MultiSetOf((_this).F6(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("idv")).Cardinality()))
				(globalState).F5 = (((_this).F6()).Minus((_665_v23).Cardinality())).Plus(Companion_Default___.SafeDivisionInt((_this).F6(), (_this).F6()))
				var _666_v24 _dafny.Map
				_ = _666_v24
				_666_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F6())
				_666_v24 = (_666_v24).Update(p0, (_this).F6())
				var _667_v25 _dafny.Array
				_ = _667_v25
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_12
				var _nw92 _dafny.Array
				_ = _nw92
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw92 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) _dafny.Sequence = func(_668_i0 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf((_this).F6(), (_this).F6())
					}
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw92 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw92).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw92).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_667_v25 = _nw92
				_667_v25 = _667_v25
				var _669_v26 _dafny.Array
				_ = _669_v26
				var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
				_ = _nw93
				_669_v26 = _nw93
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_669_v26), 0))
				_ = _index97
				(_669_v26).ArraySet1((_this).F6(), (_index97).Int())
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_669_v26), 0))
				_ = _index98
				var _rhs103 bool = p1
				_ = _rhs103
				var _rhs104 _dafny.Int = (_this).F6()
				_ = _rhs104
				var _lhs86 *GlobalState = globalState
				_ = _lhs86
				var _lhs87 _dafny.Array = _669_v26
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_669_v26), 0))
				_ = _lhs88
				_lhs86.F0 = _rhs103
				(_lhs87).ArraySet1(_rhs104, (_lhs88).Int())
				var _670_v27 _dafny.Sequence
				_ = _670_v27
				_670_v27 = _dafny.SeqOf((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), Companion_Default___.Fm3(globalState), p1)
				var _671_v28 _dafny.Array
				_ = _671_v28
				var _nw94 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
				_ = _nw94
				_671_v28 = _nw94
				var _672_v29 _dafny.Map
				_ = _672_v29
				_672_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _671_v28)
				var _673_v30 _dafny.Set
				_ = _673_v30
				_673_v30 = _dafny.SetOf(Companion_Default___.SafeDivisionInt((_669_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_669_v26), 0))).Int()).(_dafny.Int), (_669_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_669_v26), 0))).Int()).(_dafny.Int)), (_dafny.IntOfInt64(263)).Times((_672_v29).Cardinality()))
				var _674_v31 _dafny.Sequence
				_ = _674_v31
				_674_v31 = _dafny.SeqOf((_this).F6())
				var _675_v32 D2
				_ = _675_v32
				_675_v32 = Companion_D2_.Create_DC6_((_669_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_669_v26), 0))).Int()).(_dafny.Int), _669_v26)
				var _rhs105 _dafny.Array = _669_v26
				_ = _rhs105
				var _rhs106 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_674_v31, _dafny.Companion_Sequence_.Concatenate(_674_v31, _674_v31))).Cardinality())
				_ = _rhs106
				var _rhs107 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_670_v27, _dafny.Companion_Sequence_.Update(_670_v27, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_670_v27).Cardinality()))).Uint32(), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)))
				_ = _rhs107
				var _rhs108 _dafny.Set = (_673_v30).Intersection(_dafny.SetOf(((_this).F21()).Cardinality()))
				_ = _rhs108
				var _rhs109 _dafny.Int = (_675_v32).Dtor_cf13()
				_ = _rhs109
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				_669_v26 = _rhs105
				r0 = _rhs106
				_670_v27 = _rhs107
				_673_v30 = _rhs108
				_lhs89.F5 = _rhs109
			}
			var _676_v33 _dafny.Map
			_ = _676_v33
			_676_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), true)
			var _677_v34 _dafny.Sequence
			_ = _677_v34
			_677_v34 = _dafny.SeqOf((_this).F6(), (_this).F6(), (_676_v33).Cardinality())
			var _678_v35 _dafny.Set
			_ = _678_v35
			_678_v35 = _dafny.SetOf((_dafny.Zero).Minus((_this).F6()), (_this).F6(), (_dafny.Zero).Minus((_this).F6()))
			var _679_v36 _dafny.Sequence
			_ = _679_v36
			_679_v36 = _dafny.UnicodeSeqOfUtf8Bytes("jfb")
			var _680_v37 _dafny.Map
			_ = _680_v37
			_680_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(globalState), p0)
			var _681_v38 D4
			_ = _681_v38
			_681_v38 = Companion_D4_.Create_DC14_(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)), Companion_Default___.Fm2((_this).F6(), globalState), _680_v37)
			var _682_v39 _dafny.CodePoint
			_ = _682_v39
			_682_v39 = _dafny.CodePoint('i')
			var _683_v40 _dafny.MultiSet
			_ = _683_v40
			_683_v40 = _dafny.MultiSetOf((_this).F6())
			var _684_v41 _dafny.MultiSet
			_ = _684_v41
			_684_v41 = _dafny.MultiSetOf((_683_v40).Cardinality(), _dafny.IntOfInt64(877))
			var _685_v42 D0
			_ = _685_v42
			_685_v42 = Companion_D0_.Create_DC1_((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), (_this).F6(), _682_v39)
			var _686_v43 _dafny.Sequence
			_ = _686_v43
			_686_v43 = _dafny.SeqOf(_679_v36)
			var _687_v44 _dafny.Array
			_ = _687_v44
			var _nwElement0_17 _dafny.Int = _dafny.IntOfInt64(-382)
			_ = _nwElement0_17
			var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(24))
			_ = _nw95
			(_nw95).ArraySet1(_nwElement0_17, 0)
			(_nw95).ArraySet1((_this).F6(), 1)
			(_nw95).ArraySet1((_this).F6(), 2)
			(_nw95).ArraySet1((_this).F6(), 3)
			(_nw95).ArraySet1(_dafny.IntOfUint32((_679_v36).Cardinality()), 4)
			(_nw95).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetOf(Companion_Default___.Fm0((_this).F6(), _678_v35, p0, globalState), _682_v39)).Cardinality()))), 5)
			(_nw95).ArraySet1(_dafny.IntOfUint32((_679_v36).Cardinality()), 6)
			(_nw95).ArraySet1((_this).F6(), 7)
			(_nw95).ArraySet1((_this).F6(), 8)
			(_nw95).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_682_v39), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_dafny.SeqOf(_682_v39)).Cardinality()))).Uint32(), _dafny.CodePoint('m'))).Cardinality()), 9)
			(_nw95).ArraySet1(_dafny.IntOfUint32((_677_v34).Cardinality()), 10)
			(_nw95).ArraySet1(_dafny.IntOfInt64(850), 11)
			(_nw95).ArraySet1(_dafny.IntOfInt64(216), 12)
			(_nw95).ArraySet1((_683_v40).Cardinality(), 13)
			(_nw95).ArraySet1(_dafny.IntOfInt64(-913), 14)
			(_nw95).ArraySet1((_dafny.Zero).Minus((_this).F6()), 15)
			(_nw95).ArraySet1((_this).F6(), 16)
			(_nw95).ArraySet1((_this).F6(), 17)
			(_nw95).ArraySet1((_685_v42).Dtor_cf3(), 18)
			(_nw95).ArraySet1((_this).F6(), 19)
			(_nw95).ArraySet1(_dafny.IntOfInt64(227), 20)
			(_nw95).ArraySet1((_dafny.Zero).Minus((_this).F6()), 21)
			(_nw95).ArraySet1((func() _dafny.Int {
				if (_684_v41).Contains((_this).F6()) {
					return (_684_v41).Multiplicity((_this).F6())
				}
				return _dafny.IntOfUint32(((_686_v43).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_686_v43).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
			})(), 22)
			(_nw95).ArraySet1((_this).F6(), 23)
			_687_v44 = _nw95
			var _688_v45 D2
			_ = _688_v45
			_688_v45 = Companion_D2_.Create_DC6_(_dafny.IntOfInt64(181), _687_v44)
			var _689_v46 _dafny.Array
			_ = _689_v46
			var _nwElement0_18 _dafny.Int = ((_this).F6()).Plus((_this).F6())
			_ = _nwElement0_18
			var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(27))
			_ = _nw96
			(_nw96).ArraySet1(_nwElement0_18, 0)
			(_nw96).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-882)), 1)
			(_nw96).ArraySet1(_dafny.IntOfInt64(-131), 2)
			(_nw96).ArraySet1((_this).F6(), 3)
			(_nw96).ArraySet1((_this).F6(), 4)
			(_nw96).ArraySet1(((_this).F6()).Plus(_dafny.IntOfInt64(-829)), 5)
			(_nw96).ArraySet1(_dafny.IntOfInt64(21), 6)
			(_nw96).ArraySet1((_this).F6(), 7)
			(_nw96).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_677_v34).Cardinality()), (_678_v35).Cardinality()), 8)
			(_nw96).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_679_v36).Cardinality())), 9)
			(_nw96).ArraySet1((_this).F6(), 10)
			(_nw96).ArraySet1((_this).F6(), 11)
			(_nw96).ArraySet1((_this).F6(), 12)
			(_nw96).ArraySet1((_dafny.Zero).Minus((_681_v38).Dtor_cf30()), 13)
			(_nw96).ArraySet1((_this).F6(), 14)
			(_nw96).ArraySet1(_dafny.IntOfUint32((_679_v36).Cardinality()), 15)
			(_nw96).ArraySet1((_this).F6(), 16)
			(_nw96).ArraySet1((_this).F6(), 17)
			(_nw96).ArraySet1((_this).F6(), 18)
			(_nw96).ArraySet1((_this).F6(), 19)
			(_nw96).ArraySet1((_this).F6(), 20)
			(_nw96).ArraySet1((_this).F6(), 21)
			(_nw96).ArraySet1((_this).F6(), 22)
			(_nw96).ArraySet1((_this).F6(), 23)
			(_nw96).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm30(p1, _682_v39, globalState), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(543))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg58 _dafny.Int) interface{} {
					return coer58(arg58)
				}
			}((func(_690_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_691_i1 _dafny.Int) _dafny.CodePoint {
					return _690_v39
				}
			})(_682_v39)))).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm30(p1, _682_v39, globalState)).Cardinality()))).Uint32(), _682_v39)).Cardinality()), 24)
			(_nw96).ArraySet1((_this).F6(), 25)
			(_nw96).ArraySet1((_688_v45).Dtor_cf13(), 26)
			_689_v46 = _nw96
			var _692_v47 _dafny.Map
			_ = _692_v47
			_692_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.SetOf(Companion_Default___.Fm2(_dafny.IntOfUint32((_679_v36).Cardinality()), globalState), (_this).F6(), (_this).F6()))
			var _693_v48 _dafny.Array
			_ = _693_v48
			var _nwElement0_19 D2 = Companion_D2_.Create_DC6_((_this).F6(), _687_v44)
			_ = _nwElement0_19
			var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(23))
			_ = _nw97
			(_nw97).ArraySet1(_nwElement0_19, 0)
			(_nw97).ArraySet1(_688_v45, 1)
			(_nw97).ArraySet1(_688_v45, 2)
			(_nw97).ArraySet1(_688_v45, 3)
			(_nw97).ArraySet1(_688_v45, 4)
			(_nw97).ArraySet1(_688_v45, 5)
			(_nw97).ArraySet1(_688_v45, 6)
			(_nw97).ArraySet1(_688_v45, 7)
			(_nw97).ArraySet1(Companion_D2_.Create_DC6_((_this).F6(), _687_v44), 8)
			(_nw97).ArraySet1(_688_v45, 9)
			(_nw97).ArraySet1(_688_v45, 10)
			(_nw97).ArraySet1(_688_v45, 11)
			(_nw97).ArraySet1(_688_v45, 12)
			(_nw97).ArraySet1(_688_v45, 13)
			(_nw97).ArraySet1(_688_v45, 14)
			(_nw97).ArraySet1(_688_v45, 15)
			(_nw97).ArraySet1(Companion_D2_.Create_DC6_(((func() _dafny.Set {
				if (_692_v47).Contains((_this).F6()) {
					return (_692_v47).Get((_this).F6()).(_dafny.Set)
				}
				return _678_v35
			})()).Cardinality(), _689_v46), 16)
			(_nw97).ArraySet1(_688_v45, 17)
			(_nw97).ArraySet1(_688_v45, 18)
			(_nw97).ArraySet1(_688_v45, 19)
			(_nw97).ArraySet1(_688_v45, 20)
			(_nw97).ArraySet1(_688_v45, 21)
			(_nw97).ArraySet1(_688_v45, 22)
			_693_v48 = _nw97
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_693_v48), 0))
			_ = _index99
			(_693_v48).ArraySet1(_688_v45, (_index99).Int())
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_693_v48), 0))
			_ = _index100
			var _rhs110 _dafny.Int = (func() _dafny.Int {
				if p0 {
					return (_dafny.Zero).Minus((_this).F6())
				}
				return (_this).F6()
			})()
			_ = _rhs110
			var _rhs111 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("i")
			_ = _rhs111
			var _rhs112 _dafny.Int = (_this).F6()
			_ = _rhs112
			var _rhs113 _dafny.Array = _687_v44
			_ = _rhs113
			var _rhs114 D2 = _688_v45
			_ = _rhs114
			var _lhs90 *GlobalState = globalState
			_ = _lhs90
			var _lhs91 *GlobalState = globalState
			_ = _lhs91
			var _lhs92 *GlobalState = globalState
			_ = _lhs92
			var _lhs93 _dafny.Array = _693_v48
			_ = _lhs93
			var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_693_v48), 0))
			_ = _lhs94
			_lhs90.F5 = _rhs110
			_lhs91.F2 = _rhs111
			_lhs92.F5 = _rhs112
			_689_v46 = _rhs113
			(_lhs93).ArraySet1(_rhs114, (_lhs94).Int())
		}
		var _694_v49 _dafny.Array
		_ = _694_v49
		var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
		_ = _nw98
		_694_v49 = _nw98
		var _695_v50 _dafny.Map
		_ = _695_v50
		_695_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_694_v49, p1)
		var _696_v51 D10
		_ = _696_v51
		_696_v51 = Companion_D10_.Create_DC28_(_695_v50)
		_695_v50 = (_696_v51).Dtor_cf52()
		var _697_v52 _dafny.Set
		_ = _697_v52
		_697_v52 = _dafny.SetOf((_this).F6(), (_this).F6())
		var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_694_v49), 0))
		_ = _index101
		(_694_v49).ArraySet1((_697_v52).Cardinality(), (_index101).Int())
		var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_694_v49), 0))
		_ = _index102
		(_694_v49).ArraySet1((_this).F6(), (_index102).Int())
		var _698_v53 _dafny.Sequence
		_ = _698_v53
		_698_v53 = _dafny.SeqOf(_dafny.IntOfInt64(409))
		var _699_v54 _dafny.MultiSet
		_ = _699_v54
		_699_v54 = _dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(665), _dafny.IntOfInt64(647))).Cardinality(), (_this).F6())
		var _700_v55 _dafny.MultiSet
		_ = _700_v55
		_700_v55 = _dafny.MultiSetOf(_699_v54)
		var _701_v56 T0
		_ = _701_v56
		var _nw99 *C2 = New_C2_()
		_ = _nw99
		_nw99.Ctor__((_dafny.Zero).Minus((_698_v53).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_698_v53).Cardinality()))).Uint32()).(_dafny.Int)), _700_v55, _this.F9(), (_this).F10())
		_701_v56 = _nw99
		var _702_v57 _dafny.Map
		_ = _702_v57
		_702_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_701_v56, p0)
		var _703_v58 _dafny.Map
		_ = _703_v58
		_703_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_702_v57, _698_v53)
		var _704_v59 _dafny.Map
		_ = _704_v59
		_704_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _703_v58)
		_704_v59 = (_704_v59).Update((_694_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_694_v49), 0))).Int()).(_dafny.Int), _703_v58)
		var _705_i2 _dafny.Int
		_ = _705_i2
		_705_i2 = _dafny.Zero
		{
			for p0 {
				{
					if (_705_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_705_i2 = (_705_i2).Plus(_dafny.One)
					r2 = p1
					r0 = _dafny.IntOfInt64(839)
					var _arr15 _dafny.Array = _this.F9()
					_ = _arr15
					var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_this.F9()), 0))
					_ = _index103
					(_arr15).ArraySet1(p1, (_index103).Int())
					var _arr16 _dafny.Array = _this.F9()
					_ = _arr16
					var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_this.F9()), 0))
					_ = _index104
					(_arr16).ArraySet1((func() bool {
						if (p0) || (p0) {
							return p1
						}
						return true
					})(), (_index104).Int())
					var _706_v60 _dafny.Array
					_ = _706_v60
					var _nwElement0_20 _dafny.Array = _694_v49
					_ = _nwElement0_20
					var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(6))
					_ = _nw100
					(_nw100).ArraySet1(_nwElement0_20, 0)
					(_nw100).ArraySet1(_694_v49, 1)
					(_nw100).ArraySet1(_694_v49, 2)
					(_nw100).ArraySet1(_694_v49, 3)
					(_nw100).ArraySet1(_694_v49, 4)
					(_nw100).ArraySet1(_694_v49, 5)
					_706_v60 = _nw100
					var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_706_v60), 0))
					_ = _index105
					(_706_v60).ArraySet1(_694_v49, (_index105).Int())
					var _707_v61 _dafny.Sequence
					_ = _707_v61
					_707_v61 = _dafny.SeqOf(_698_v53, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(640))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg59 _dafny.Int) interface{} {
							return coer59(arg59)
						}
					}((func(_708_v56 T0) func(_dafny.Int) _dafny.Int {
						return func(_709_i3 _dafny.Int) _dafny.Int {
							return (_708_v56).F6()
						}
					})(_701_v56))), _698_v53)
					var _710_v62 _dafny.Array
					_ = _710_v62
					var _len0_13 _dafny.Int = _dafny.IntOfInt64(9)
					_ = _len0_13
					var _nw101 _dafny.Array
					_ = _nw101
					if _len0_13.Cmp(_dafny.Zero) == 0 {
						_nw101 = _dafny.NewArray(_len0_13)
					} else {
						var _init13 func(_dafny.Int) bool = (func(_711_p0 bool) func(_dafny.Int) bool {
							return func(_712_i4 _dafny.Int) bool {
								return _711_p0
							}
						})(p0)
						_ = _init13
						var _element0_13 = _init13(_dafny.Zero)
						_ = _element0_13
						_nw101 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
						(_nw101).ArraySet1(_element0_13, 0)
						var _nativeLen0_13 = (_len0_13).Int()
						_ = _nativeLen0_13
						for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
							(_nw101).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
						}
					}
					_710_v62 = _nw101
					var _713_v63 _dafny.Map
					_ = _713_v63
					_713_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_707_v61).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_707_v61).Cardinality()))).Uint32()).(_dafny.Sequence), _710_v62)
					var _714_v64 _dafny.Set
					_ = _714_v64
					_714_v64 = _dafny.SetOf((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), true)
					var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_694_v49), 0))
					_ = _index106
					var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_706_v60), 0))
					_ = _index107
					var _rhs115 _dafny.Int = (Companion_Default___.SafeModuloInt((_this).F6(), _dafny.IntOfUint32((Companion_Default___.Fm40((_this).F6(), _714_v64, globalState)).Cardinality()))).Times((_dafny.Zero).Minus((_701_v56).F6()))
					_ = _rhs115
					var _rhs116 _dafny.Int = (_694_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_694_v49), 0))).Int()).(_dafny.Int)
					_ = _rhs116
					var _rhs117 _dafny.Array = _694_v49
					_ = _rhs117
					var _rhs118 _dafny.Map = ((_713_v63).Merge(_713_v63)).Merge(_713_v63)
					_ = _rhs118
					var _rhs119 bool = p1
					_ = _rhs119
					var _lhs95 _dafny.Array = _694_v49
					_ = _lhs95
					var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_694_v49), 0))
					_ = _lhs96
					var _lhs97 _dafny.Array = _706_v60
					_ = _lhs97
					var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_706_v60), 0))
					_ = _lhs98
					var _lhs99 *GlobalState = globalState
					_ = _lhs99
					r0 = _rhs115
					(_lhs95).ArraySet1(_rhs116, (_lhs96).Int())
					(_lhs97).ArraySet1(_rhs117, (_lhs98).Int())
					_713_v63 = _rhs118
					_lhs99.F0 = _rhs119
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _rhs120 bool = p0
		_ = _rhs120
		var _rhs121 _dafny.Array = (func() _dafny.Array {
			if p0 {
				return _694_v49
			}
			return _694_v49
		})()
		_ = _rhs121
		r2 = _rhs120
		_694_v49 = _rhs121
		r0 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wwrq")).Cardinality()), (_694_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_694_v49), 0))).Int()).(_dafny.Int))
		var _715_v65 _dafny.CodePoint
		_ = _715_v65
		_715_v65 = _dafny.CodePoint('y')
		r1 = _715_v65
		var _716_v66 _dafny.Set
		_ = _716_v66
		_716_v66 = _dafny.SetOf(p0, p1, true)
		r2 = (_716_v66).IsProperSubsetOf(_dafny.SetOf(p0))
		return r0, r1, r2
	}
}
func (_this *C3) F21() _dafny.Map {
	{
		return _this._f21
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f9  _dafny.Array
	_f16 _dafny.MultiSet
	_f6  _dafny.Int
	_f10 _dafny.Map
	_f17 _dafny.Int
	_f18 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f16 = _dafny.EmptyMultiSet
	_this._f6 = _dafny.Zero
	_this._f10 = _dafny.EmptyMap
	_this._f17 = _dafny.Zero
	_this._f18 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_, Companion_T2_.TraitID_}
}

var _ T1 = &C4{}
var _ T0 = &C4{}
var _ T2 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F9() _dafny.Array {
	return _this._f9
}
func (_this *C4) F9_set_(value _dafny.Array) {
	_this._f9 = value
}
func (_this *C4) F16() _dafny.MultiSet {
	return _this._f16
}
func (_this *C4) F16_set_(value _dafny.MultiSet) {
	_this._f16 = value
}
func (_this *C4) F6() _dafny.Int {
	return _this._f6
}
func (_this *C4) F10() _dafny.Map {
	return _this._f10
}
func (_this *C4) Ctor__(f17 _dafny.Int, f18 _dafny.Int, f9 _dafny.Array, f10 _dafny.Map, f6 _dafny.Int, f16 _dafny.MultiSet) {
	{
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f6 = f6
		(_this)._f16 = f16
	}
}
func (_this *C4) Fm13(p0 _dafny.CodePoint, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var _source9 D1 = Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_dafny.CodePoint('n'), _dafny.CodePoint('o'), _dafny.CodePoint('p'), _dafny.CodePoint('p')))
		_ = _source9
		if _source9.Is_DC4() {
			var _717___mcc_h0 D0 = _source9.Get_().(D1_DC4).Cf7
			_ = _717___mcc_h0
			var _718___mcc_h1 _dafny.Int = _source9.Get_().(D1_DC4).Cf8
			_ = _718___mcc_h1
			var _719___mcc_h2 _dafny.Sequence = _source9.Get_().(D1_DC4).Cf9
			_ = _719___mcc_h2
			var _720___mcc_h3 bool = _source9.Get_().(D1_DC4).Cf10
			_ = _720___mcc_h3
			var _721___mcc_h4 _dafny.Int = _source9.Get_().(D1_DC4).Cf11
			_ = _721___mcc_h4
			var _722_cf11 _dafny.Int = _721___mcc_h4
			_ = _722_cf11
			var _723_cf10 bool = _720___mcc_h3
			_ = _723_cf10
			var _724_cf9 _dafny.Sequence = _719___mcc_h2
			_ = _724_cf9
			var _725_cf8 _dafny.Int = _718___mcc_h1
			_ = _725_cf8
			var _726_cf7 D0 = _717___mcc_h0
			_ = _726_cf7
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _724_cf9)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_723_cf10, _724_cf9))
		} else {
			var _727___mcc_h5 _dafny.MultiSet = _source9.Get_().(D1_DC3).Cf6
			_ = _727___mcc_h5
			var _728_cf6 _dafny.MultiSet = _727___mcc_h5
			_ = _728_cf6
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("kg"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("xmpjtp")))
		}
	}
}
func (_this *C4) Fm5(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		var _source10 D0 = Companion_D0_.Create_DC0_(_dafny.SeqOf(false, true), (_this).F6())
		_ = _source10
		if _source10.Is_DC0() {
			var _729___mcc_h0 _dafny.Sequence = _source10.Get_().(D0_DC0).Cf0
			_ = _729___mcc_h0
			var _730___mcc_h1 _dafny.Int = _source10.Get_().(D0_DC0).Cf1
			_ = _730___mcc_h1
			var _731_cf1 _dafny.Int = _730___mcc_h1
			_ = _731_cf1
			var _732_cf0 _dafny.Sequence = _729___mcc_h0
			_ = _732_cf0
			return _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('a'))), _dafny.MultiSetOf(_dafny.CodePoint('w'), _dafny.CodePoint('f')), _dafny.MultiSetOf(_dafny.CodePoint('d')), _dafny.MultiSetOf(_dafny.CodePoint('u')), _dafny.MultiSetOf(_dafny.CodePoint('h'), _dafny.CodePoint('d')))
		} else if _source10.Is_DC1() {
			var _733___mcc_h2 bool = _source10.Get_().(D0_DC1).Cf2
			_ = _733___mcc_h2
			var _734___mcc_h3 _dafny.Int = _source10.Get_().(D0_DC1).Cf3
			_ = _734___mcc_h3
			var _735___mcc_h4 _dafny.CodePoint = _source10.Get_().(D0_DC1).Cf4
			_ = _735___mcc_h4
			var _736_cf4 _dafny.CodePoint = _735___mcc_h4
			_ = _736_cf4
			var _737_cf3 _dafny.Int = _734___mcc_h3
			_ = _737_cf3
			var _738_cf2 bool = _733___mcc_h2
			_ = _738_cf2
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_736_cf4), _dafny.MultiSetOf(_736_cf4)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(448))).Uint32(), func(coer60 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg60 _dafny.Int) interface{} {
					return coer60(arg60)
				}
			}((func(_739_cf4 _dafny.CodePoint) func(_dafny.Int) _dafny.MultiSet {
				return func(_740_i0 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_739_cf4)
				}
			})(_736_cf4))))
		} else {
			var _741___mcc_h5 D0 = _source10.Get_().(D0_DC2).Cf5
			_ = _741___mcc_h5
			var _742_cf5 D0 = _741___mcc_h5
			_ = _742_cf5
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(409))).Uint32(), func(coer61 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg61 _dafny.Int) interface{} {
					return coer61(arg61)
				}
			}(func(_743_i1 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(_dafny.CodePoint('e'), _dafny.CodePoint('g'))
			})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(613))).Uint32(), func(coer62 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg62 _dafny.Int) interface{} {
					return coer62(arg62)
				}
			}(func(_744_i2 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(237))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}(func(_745_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				})))
			})))
		}
	}
}
func (_this *C4) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _dafny.SeqOf((_this).F18()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(669))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg64 _dafny.Int) interface{} {
				return coer64(arg64)
			}
		}(func(_746_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		}))).Cardinality()), _dafny.SeqOf((_this).F18()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()), _dafny.SeqOf((_this).F18()))))
	}
}
func (_this *C4) Fm21(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt((_this).F18(), ((_this).F17()).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false))).Cardinality()))
	}
}
func (_this *C4) Fm22(p0 _dafny.MultiSet, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		var _source11 D0 = (func() D0 {
			if !(false) {
				return Companion_D0_.Create_DC1_(false, (_this).F6(), _dafny.CodePoint('y'))
			}
			return Companion_D0_.Create_DC1_(true, (_this).F18(), _dafny.CodePoint('n'))
		})()
		_ = _source11
		if _source11.Is_DC0() {
			var _747___mcc_h0 _dafny.Sequence = _source11.Get_().(D0_DC0).Cf0
			_ = _747___mcc_h0
			var _748___mcc_h1 _dafny.Int = _source11.Get_().(D0_DC0).Cf1
			_ = _748___mcc_h1
			var _749_cf1 _dafny.Int = _748___mcc_h1
			_ = _749_cf1
			var _750_cf0 _dafny.Sequence = _747___mcc_h0
			_ = _750_cf0
			return !(true)
		} else if _source11.Is_DC1() {
			var _751___mcc_h2 bool = _source11.Get_().(D0_DC1).Cf2
			_ = _751___mcc_h2
			var _752___mcc_h3 _dafny.Int = _source11.Get_().(D0_DC1).Cf3
			_ = _752___mcc_h3
			var _753___mcc_h4 _dafny.CodePoint = _source11.Get_().(D0_DC1).Cf4
			_ = _753___mcc_h4
			var _754_cf4 _dafny.CodePoint = _753___mcc_h4
			_ = _754_cf4
			var _755_cf3 _dafny.Int = _752___mcc_h3
			_ = _755_cf3
			var _756_cf2 bool = _751___mcc_h2
			_ = _756_cf2
			return _756_cf2
		} else {
			var _757___mcc_h5 D0 = _source11.Get_().(D0_DC2).Cf5
			_ = _757___mcc_h5
			var _758_cf5 D0 = _757___mcc_h5
			_ = _758_cf5
			return (_dafny.SetOf(false)).IsProperSubsetOf(_dafny.SetOf(true))
		}
	}
}
func (_this *C4) M6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _759_v0 _dafny.Map
		_ = _759_v0
		_759_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(18), p0)
		var _hi5 _dafny.Int = (p2).Minus(p2)
		_ = _hi5
		for _760_i0 := ((_759_v0).Merge(_759_v0)).Cardinality(); _760_i0.Cmp(_hi5) < 0; _760_i0 = _760_i0.Plus(_dafny.One) {
			(globalState).F5 = p2
			var _761_v1 D0
			_ = _761_v1
			_761_v1 = Companion_D0_.Create_DC0_(p1, (_this).F18())
			var _762_v2 _dafny.Sequence
			_ = _762_v2
			_762_v2 = _dafny.UnicodeSeqOfUtf8Bytes("jnbsg")
			var _763_v3 D1
			_ = _763_v3
			_763_v3 = Companion_D1_.Create_DC4_(_761_v1, (_this).F6(), (func() _dafny.Sequence {
				if p0 {
					return _762_v2
				}
				return _762_v2
			})(), p0, (_this).F18())
			_763_v3 = Companion_D1_.Create_DC4_(_761_v1, _dafny.IntOfInt64(403), _762_v2, p0, (_this).F18())
			_762_v2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_762_v2, _762_v2), _dafny.Companion_Sequence_.Concatenate(_762_v2, _762_v2))
			var _764_v4 _dafny.Array
			_ = _764_v4
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_14
			var _nw102 _dafny.Array
			_ = _nw102
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw102 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) D1 = func(_765_i1 _dafny.Int) D1 {
					return Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_dafny.CodePoint('g')))
				}
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw102 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw102).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw102).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_764_v4 = _nw102
			var _766_v5 _dafny.Array
			_ = _766_v5
			var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(26))
			_ = _nw103
			_766_v5 = _nw103
			var _767_v6 _dafny.Map
			_ = _767_v6
			_767_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_764_v4, _766_v5)
			var _768_v7 _dafny.Map
			_ = _768_v7
			_768_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (p1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-465), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool))
			var _769_v8 _dafny.MultiSet
			_ = _769_v8
			_769_v8 = _dafny.MultiSetOf(p0, (func() bool {
				if (_768_v7).Contains(!(p0)) {
					return (_768_v7).Get(!(p0)).(bool)
				}
				return p0
			})())
			var _770_v9 D2
			_ = _770_v9
			_770_v9 = Companion_D2_.Create_DC7_(_dafny.Companion_Sequence_.Concatenate(p1, Companion_Default___.Fm4(_760_i0, (_dafny.Zero).Minus(_760_i0), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality()), globalState)), (func() _dafny.Array {
				if (_767_v6).Contains(_764_v4) {
					return (_767_v6).Get(_764_v4).(_dafny.Array)
				}
				return _766_v5
			})(), Companion_Default___.Fm2((_this).F17(), globalState), ((_this).F18()).Plus((_769_v8).Cardinality()))
			_770_v9 = _770_v9
		}
		(globalState).F5 = (_this).F18()
		var _771_i2 _dafny.Int
		_ = _771_i2
		_771_i2 = _dafny.Zero
		{
			for p0 {
				{
					if (_771_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_771_i2 = (_771_i2).Plus(_dafny.One)
					var _772_v10 *C0
					_ = _772_v10
					var _nw104 *C0 = New_C0_()
					_ = _nw104
					_nw104.Ctor__(_this.F9(), Companion_Default___.Fm3(globalState))
					_772_v10 = _nw104
					(globalState).F5 = (_this).F18()
					var _773_v11 D0
					_ = _773_v11
					_773_v11 = Companion_D0_.Create_DC0_(_dafny.SeqOf(p0), (_this).F6())
					(globalState).F5 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((Companion_Default___.Fm25(p2, _772_v10.F20, globalState)).Cardinality()), (_773_v11).Dtor_cf1())
					var _774_v12 _dafny.Array
					_ = _774_v12
					var _len0_15 _dafny.Int = _dafny.IntOfInt64(9)
					_ = _len0_15
					var _nw105 _dafny.Array
					_ = _nw105
					if _len0_15.Cmp(_dafny.Zero) == 0 {
						_nw105 = _dafny.NewArray(_len0_15)
					} else {
						var _init15 func(_dafny.Int) _dafny.Int = func(_775_i3 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_775_i3, (_this).F18())
						}
						_ = _init15
						var _element0_15 = _init15(_dafny.Zero)
						_ = _element0_15
						_nw105 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
						(_nw105).ArraySet1(_element0_15, 0)
						var _nativeLen0_15 = (_len0_15).Int()
						_ = _nativeLen0_15
						for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
							(_nw105).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
						}
					}
					_774_v12 = _nw105
					var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_774_v12), 0))
					_ = _index108
					(_774_v12).ArraySet1((_this).F6(), (_index108).Int())
					var _776_v13 _dafny.Sequence
					_ = _776_v13
					_776_v13 = _dafny.SeqOf(_dafny.IntOfInt64(8), (_this).F17())
					var _777_v14 D2
					_ = _777_v14
					_777_v14 = Companion_D2_.Create_DC8_(_772_v10.F20, p0)
					var _778_v15 _dafny.MultiSet
					_ = _778_v15
					_778_v15 = _dafny.MultiSetOf((_this).F6())
					var _779_v16 _dafny.CodePoint
					_ = _779_v16
					_779_v16 = _dafny.CodePoint('l')
					var _780_v17 _dafny.Map
					_ = _780_v17
					_780_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_779_v16, (_this).F6())
					var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_774_v12), 0))
					_ = _index109
					var _rhs122 bool = _772_v10.F20
					_ = _rhs122
					var _rhs123 _dafny.Int = (_this).F6()
					_ = _rhs123
					var _rhs124 bool = _dafny.Companion_Sequence_.IsPrefixOf(_776_v13, _776_v13)
					_ = _rhs124
					var _rhs125 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_772_v10.F20, _777_v14)).Cardinality()
					_ = _rhs125
					var _rhs126 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((Companion_Default___.Fm4((_dafny.Zero).Minus(((_778_v15).Update((_this).F6(), Companion_Default___.Abs((_dafny.Zero).Minus((_this).F18())))).Cardinality()), _dafny.IntOfInt64(621), (_this).F18(), globalState)).Cardinality()), Companion_Default___.SafeModuloInt((func() _dafny.Int {
						if (_780_v17).Contains(_dafny.CodePoint('g')) {
							return (_780_v17).Get(_dafny.CodePoint('g')).(_dafny.Int)
						}
						return (_this).F6()
					})(), (_this).F6()))
					_ = _rhs126
					var _lhs100 _dafny.Array = _774_v12
					_ = _lhs100
					var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_774_v12), 0))
					_ = _lhs101
					var _lhs102 *GlobalState = globalState
					_ = _lhs102
					var _lhs103 *GlobalState = globalState
					_ = _lhs103
					var _lhs104 *GlobalState = globalState
					_ = _lhs104
					r1 = _rhs122
					(_lhs100).ArraySet1(_rhs123, (_lhs101).Int())
					_lhs102.F0 = _rhs124
					_lhs103.F5 = _rhs125
					_lhs104.F5 = _rhs126
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		if p0 {
			var _781_v18 D0
			_ = _781_v18
			_781_v18 = Companion_D0_.Create_DC0_(p1, p2)
			var _782_v19 D0
			_ = _782_v19
			_782_v19 = Companion_D0_.Create_DC2_(_781_v18)
			var _783_v20 _dafny.Map
			_ = _783_v20
			_783_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _782_v19)
			var _784_v21 _dafny.Set
			_ = _784_v21
			_784_v21 = _dafny.SetOf(_783_v20)
			(globalState).F0 = !(((func() _dafny.Set {
				var _coll19 = _dafny.NewBuilder()
				_ = _coll19
				for _iter21 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-885))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}((func(_785_v20 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_786_i4 _dafny.Int) _dafny.Map {
						return _785_v20
					}
				})(_783_v20)))).Elements()); ; {
					_compr_19, _ok21 := _iter21()
					if !_ok21 {
						break
					}
					var _787_v22 _dafny.Map
					_787_v22 = interface{}(_compr_19).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-885))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg66 _dafny.Int) interface{} {
							return coer66(arg66)
						}
					}((func(_788_v20 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_786_i4 _dafny.Int) _dafny.Map {
							return _788_v20
						}
					})(_783_v20))), _787_v22) {
						_coll19.Add(_787_v22)
					}
				}
				return _coll19.ToSet()
			}()).Union(_784_v21)).IsProperSubsetOf(_784_v21))
			var _789_v23 _dafny.CodePoint
			_ = _789_v23
			_789_v23 = _dafny.CodePoint('f')
			_789_v23 = _789_v23
			var _790_v24 *C0
			_ = _790_v24
			var _nw106 *C0 = New_C0_()
			_ = _nw106
			_nw106.Ctor__(_this.F9(), p0)
			_790_v24 = _nw106
			_790_v24 = _790_v24
			var _791_v25 _dafny.Array
			_ = _791_v25
			var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(2))
			_ = _nw107
			_791_v25 = _nw107
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_791_v25), 0))
			_ = _index110
			(_791_v25).ArraySet1CodePoint(_789_v23, (_index110).Int())
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_791_v25), 0))
			_ = _index111
			(_791_v25).ArraySet1CodePoint(_789_v23, (_index111).Int())
			var _792_v26 _dafny.Sequence
			_ = _792_v26
			_792_v26 = _dafny.UnicodeSeqOfUtf8Bytes("wjkiaxif")
			var _793_v27 _dafny.MultiSet
			_ = _793_v27
			_793_v27 = _dafny.MultiSetOf(_dafny.IntOfUint32((_792_v26).Cardinality()))
			(globalState).F0 = ((_this).Fm21((_dafny.Zero).Minus((_759_v0).Cardinality()), globalState)).Cmp(Companion_Default___.SafeModuloInt((_this).F17(), (_793_v27).Cardinality())) < 0
		} else {
			if ((_this).F18()).Cmp(p2) != 0 {
				var _794_v28 _dafny.Array
				_ = _794_v28
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_16
				var _nw108 _dafny.Array
				_ = _nw108
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw108 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.Int = func(_795_i5 _dafny.Int) _dafny.Int {
						return (_795_i5).Minus((_this).F6())
					}
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
				_794_v28 = _nw108
				var _796_v29 _dafny.Map
				_ = _796_v29
				_796_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_794_v28, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, p1)).Cardinality()))
				_796_v29 = _796_v29
				var _797_v30 *C0
				_ = _797_v30
				var _nw109 *C0 = New_C0_()
				_ = _nw109
				_nw109.Ctor__(_this.F9(), p0)
				_797_v30 = _nw109
				var _798_v31 *C0
				_ = _798_v31
				var _nw110 *C0 = New_C0_()
				_ = _nw110
				_nw110.Ctor__(_this.F9(), p0)
				_798_v31 = _nw110
				var _799_v32 _dafny.Sequence
				_ = _799_v32
				_799_v32 = _dafny.SeqOf(_798_v31.F20)
				var _rhs127 _dafny.Sequence = p1
				_ = _rhs127
				var _rhs128 bool = true
				_ = _rhs128
				var _lhs105 *C0 = _798_v31
				_ = _lhs105
				_799_v32 = _rhs127
				_lhs105.F20 = _rhs128
				var _800_v33 *C0
				_ = _800_v33
				var _nw111 *C0 = New_C0_()
				_ = _nw111
				_nw111.Ctor__(_this.F9(), p0)
				_800_v33 = _nw111
			} else {
				(globalState).F0 = (_dafny.SetOf(p0)).IsSubsetOf((_dafny.SetOf(p0, p0, false)).Difference(_dafny.SetOf(p0, p0)))
				(globalState).F5 = (_this).F17()
				(globalState).F0 = p0
				var _801_v34 *C0
				_ = _801_v34
				var _nw112 *C0 = New_C0_()
				_ = _nw112
				_nw112.Ctor__((func() _dafny.Array {
					if p0 {
						return _this.F9()
					}
					return _this.F9()
				})(), p0)
				_801_v34 = _nw112
				var _802_v35 _dafny.Array
				_ = _802_v35
				var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw113
				_802_v35 = _nw113
				var _803_v36 _dafny.Sequence
				_ = _803_v36
				_803_v36 = _dafny.UnicodeSeqOfUtf8Bytes("koabsqd")
				var _804_v37 _dafny.CodePoint
				_ = _804_v37
				_804_v37 = _dafny.CodePoint('v')
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_802_v35), 0))
				_ = _index112
				(_802_v35).ArraySet1((_dafny.IntOfUint32((_803_v36).Cardinality())).Times(_dafny.IntOfUint32((Companion_Default___.Fm26((_this).F6(), _804_v37, _801_v34.F20, _804_v37, globalState)).Cardinality())), (_index112).Int())
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_802_v35), 0))
				_ = _index113
				(_802_v35).ArraySet1((_this).F6(), (_index113).Int())
			}
			var _805_v38 _dafny.Sequence
			_ = _805_v38
			_805_v38 = _dafny.SeqOf((_this).F17())
			var _806_v39 *C0
			_ = _806_v39
			var _nw114 *C0 = New_C0_()
			_ = _nw114
			_nw114.Ctor__(_this.F9(), ((_this).F17()).Cmp((_805_v38).Select((Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_805_v38).Cardinality()))).Uint32()).(_dafny.Int)) > 0)
			_806_v39 = _nw114
			var _807_v40 _dafny.Sequence
			_ = _807_v40
			_807_v40 = _dafny.UnicodeSeqOfUtf8Bytes("snhg")
			var _808_v41 _dafny.Set
			_ = _808_v41
			_808_v41 = _dafny.SetOf(_806_v39.F20)
			var _809_v42 _dafny.CodePoint
			_ = _809_v42
			_809_v42 = _dafny.CodePoint('b')
			var _810_v43 _dafny.Map
			_ = _810_v43
			_810_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_808_v41).Cardinality())).Cardinality(), _809_v42)
			var _811_v45 _dafny.MultiSet
			_ = _811_v45
			_811_v45 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(18))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg67 _dafny.Int) interface{} {
					return coer67(arg67)
				}
			}((func(_812_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_813_i6 _dafny.Int) _dafny.CodePoint {
					return _812_v42
				}
			})(_809_v42))))
			var _814_v46 _dafny.Map
			_ = _814_v46
			_814_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_807_v40, _dafny.IntOfInt64(-128))
			var _815_v48 _dafny.Array
			_ = _815_v48
			var _nwElement0_21 _dafny.Map = (func() _dafny.Map {
				if p0 {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_807_v40, _dafny.IntOfUint32((_dafny.SeqOf((_807_v40).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_807_v40).Cardinality()))).Uint32()).(_dafny.CodePoint), (func() _dafny.CodePoint {
						if (_810_v43).Contains((_this).F17()) {
							return (_810_v43).Get((_this).F17()).(_dafny.CodePoint)
						}
						return (_807_v40).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_807_v40).Cardinality()))).Uint32()).(_dafny.CodePoint)
					})())).Cardinality()))
				}
				return func() _dafny.Map {
					var _coll20 = _dafny.NewMapBuilder()
					_ = _coll20
					for _iter22 := _dafny.Iterate((_811_v45).Elements()); ; {
						_compr_20, _ok22 := _iter22()
						if !_ok22 {
							break
						}
						var _816_v44 _dafny.Sequence
						_816_v44 = interface{}(_compr_20).(_dafny.Sequence)
						if (_811_v45).Contains(_816_v44) {
							_coll20.Add(_816_v44, _dafny.IntOfInt64(313))
						}
					}
					return _coll20.ToMap()
				}()
			})()
			_ = _nwElement0_21
			var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(13))
			_ = _nw115
			(_nw115).ArraySet1(_nwElement0_21, 0)
			(_nw115).ArraySet1(_814_v46, 1)
			(_nw115).ArraySet1(_814_v46, 2)
			(_nw115).ArraySet1(_814_v46, 3)
			(_nw115).ArraySet1(_814_v46, 4)
			(_nw115).ArraySet1(Companion_Default___.Fm27(globalState), 5)
			(_nw115).ArraySet1(_814_v46, 6)
			(_nw115).ArraySet1(_814_v46, 7)
			(_nw115).ArraySet1(_814_v46, 8)
			(_nw115).ArraySet1(_814_v46, 9)
			(_nw115).ArraySet1((_814_v46).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_807_v40, (_dafny.Zero).Minus(_dafny.IntOfInt64(-528)))), 10)
			(_nw115).ArraySet1(_814_v46, 11)
			(_nw115).ArraySet1((func() _dafny.Map {
				var _coll21 = _dafny.NewMapBuilder()
				_ = _coll21
				for _iter23 := _dafny.Iterate(((_dafny.MultiSetOf(_807_v40)).Update(_807_v40, Companion_Default___.Abs((_this).F17()))).Elements()); ; {
					_compr_21, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _817_v47 _dafny.Sequence
					_817_v47 = interface{}(_compr_21).(_dafny.Sequence)
					if ((_dafny.MultiSetOf(_807_v40)).Update(_807_v40, Companion_Default___.Abs((_this).F17()))).Contains(_817_v47) {
						_coll21.Add(_817_v47, (_this).F17())
					}
				}
				return _coll21.ToMap()
			}()).Merge(_814_v46), 12)
			_815_v48 = _nw115
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_815_v48), 0))
			_ = _index114
			(_815_v48).ArraySet1(_814_v46, (_index114).Int())
			var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_815_v48), 0))
			_ = _index115
			(_815_v48).ArraySet1(_814_v46, (_index115).Int())
			var _818_v49 _dafny.Set
			_ = _818_v49
			_818_v49 = _dafny.SetOf((_this).F17(), (_dafny.Zero).Minus(p2), (_dafny.Zero).Minus((_this).F6()))
			var _819_v50 _dafny.MultiSet
			_ = _819_v50
			_819_v50 = _dafny.MultiSetOf((_this).F18(), p2, (_818_v49).Cardinality())
			var _rhs129 _dafny.MultiSet = (func() _dafny.MultiSet {
				if !((p0) == (p0)) {
					return _819_v50
				}
				return _819_v50
			})()
			_ = _rhs129
			var _rhs130 bool = _806_v39.F20
			_ = _rhs130
			var _lhs106 *C0 = _806_v39
			_ = _lhs106
			_819_v50 = _rhs129
			_lhs106.F20 = _rhs130
			var _arr17 _dafny.Array = _806_v39.F19
			_ = _arr17
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_806_v39.F19), 0))
			_ = _index116
			(_arr17).ArraySet1(p0, (_index116).Int())
			var _arr18 _dafny.Array = _806_v39.F19
			_ = _arr18
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_806_v39.F19), 0))
			_ = _index117
			(_arr18).ArraySet1((_819_v50).IsSubsetOf((_819_v50).Union(_819_v50)), (_index117).Int())
		}
		var _820_i7 _dafny.Int
		_ = _820_i7
		_820_i7 = _dafny.Zero
		{
			for p0 {
				{
					if (_820_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_820_i7 = (_820_i7).Plus(_dafny.One)
					var _821_v51 _dafny.MultiSet
					_ = _821_v51
					_821_v51 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9(), (_this).F18())).Cardinality())
					_759_v0 = (_759_v0).Update(Companion_Default___.SafeModuloInt(p2, (_821_v51).Cardinality()), true)
					var _arr19 _dafny.Array = _this.F9()
					_ = _arr19
					var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))
					_ = _index118
					(_arr19).ArraySet1(p0, (_index118).Int())
					var _822_v52 _dafny.Array
					_ = _822_v52
					var _len0_17 _dafny.Int = _dafny.IntOfInt64(8)
					_ = _len0_17
					var _nw116 _dafny.Array
					_ = _nw116
					if _len0_17.Cmp(_dafny.Zero) == 0 {
						_nw116 = _dafny.NewArray(_len0_17)
					} else {
						var _init17 func(_dafny.Int) _dafny.Int = func(_823_i8 _dafny.Int) _dafny.Int {
							return (_823_i8).Minus((_this).F17())
						}
						_ = _init17
						var _element0_17 = _init17(_dafny.Zero)
						_ = _element0_17
						_nw116 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
						(_nw116).ArraySet1(_element0_17, 0)
						var _nativeLen0_17 = (_len0_17).Int()
						_ = _nativeLen0_17
						for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
							(_nw116).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
						}
					}
					_822_v52 = _nw116
					var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_822_v52), 0))
					_ = _index119
					(_822_v52).ArraySet1(p2, (_index119).Int())
					var _824_v53 _dafny.Sequence
					_ = _824_v53
					_824_v53 = _dafny.UnicodeSeqOfUtf8Bytes("wmae")
					var _arr20 _dafny.Array = _this.F9()
					_ = _arr20
					var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))
					_ = _index120
					var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_822_v52), 0))
					_ = _index121
					var _rhs131 bool = p0
					_ = _rhs131
					var _rhs132 _dafny.Int = (_this).F6()
					_ = _rhs132
					var _rhs133 bool = p0
					_ = _rhs133
					var _rhs134 _dafny.Int = (_dafny.IntOfUint32((_824_v53).Cardinality())).Plus((_this).F17())
					_ = _rhs134
					var _lhs107 *GlobalState = globalState
					_ = _lhs107
					var _lhs108 *GlobalState = globalState
					_ = _lhs108
					var _lhs109 _dafny.Array = _this.F9()
					_ = _lhs109
					var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))
					_ = _lhs110
					var _lhs111 _dafny.Array = _822_v52
					_ = _lhs111
					var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_822_v52), 0))
					_ = _lhs112
					_lhs107.F0 = _rhs131
					_lhs108.F5 = _rhs132
					(_lhs109).ArraySet1(_rhs133, (_lhs110).Int())
					(_lhs111).ArraySet1(_rhs134, (_lhs112).Int())
					(globalState).F5 = p2
					var _825_v54 _dafny.MultiSet
					_ = _825_v54
					_825_v54 = _dafny.MultiSetOf(_822_v52)
					var _rhs135 _dafny.Array = (func() _dafny.Array {
						if (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool) {
							return (func() _dafny.Array {
								if p0 {
									return _822_v52
								}
								return _822_v52
							})()
						}
						return _822_v52
					})()
					_ = _rhs135
					var _rhs136 _dafny.MultiSet = _825_v54
					_ = _rhs136
					var _rhs137 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(751))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg68 _dafny.Int) interface{} {
							return coer68(arg68)
						}
					}((func(_826_v53 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
						return func(_827_i9 _dafny.Int) _dafny.CodePoint {
							return (_826_v53).Select((Companion_Default___.SafeIndex((_this).F17(), _dafny.IntOfUint32((_826_v53).Cardinality()))).Uint32()).(_dafny.CodePoint)
						}
					})(_824_v53)))
					_ = _rhs137
					var _rhs138 bool = Companion_Default___.Fm3(globalState)
					_ = _rhs138
					var _lhs113 *GlobalState = globalState
					_ = _lhs113
					var _lhs114 *GlobalState = globalState
					_ = _lhs114
					_822_v52 = _rhs135
					_825_v54 = _rhs136
					_lhs113.F2 = _rhs137
					_lhs114.F0 = _rhs138
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _828_v55 _dafny.Array
		_ = _828_v55
		var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(21))
		_ = _nw117
		_828_v55 = _nw117
		var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_828_v55), 0))
		_ = _index122
		(_828_v55).ArraySet1(_dafny.SetOf((_this).F6(), (_this).F17()), (_index122).Int())
		var _829_v56 _dafny.Map
		_ = _829_v56
		_829_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p2), (_this).F17())
		var _830_v57 D2
		_ = _830_v57
		_830_v57 = Companion_D2_.Create_DC5_(_829_v56)
		var _831_v58 T1
		_ = _831_v58
		var _nw118 *C3 = New_C3_()
		_ = _nw118
		_nw118.Ctor__((_830_v57).Dtor_cf12(), _this.F9(), (_this).F10(), _dafny.IntOfInt64(13))
		_831_v58 = _nw118
		var _832_v59 _dafny.Sequence
		_ = _832_v59
		_832_v59 = _dafny.SeqOf(_831_v58)
		var _833_v60 _dafny.Map
		_ = _833_v60
		_833_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _834_v61 _dafny.Sequence
		_ = _834_v61
		_834_v61 = _dafny.SeqOf(_833_v60, _833_v60)
		var _835_v62 _dafny.Set
		_ = _835_v62
		_835_v62 = _dafny.SetOf(p2, (_dafny.MultiSetFromSeq(_832_v59)).Cardinality(), _dafny.IntOfUint32((_834_v61).Cardinality()))
		var _836_v63 D11
		_ = _836_v63
		_836_v63 = Companion_D11_.Create_DC30_(_835_v62)
		var _837_v64 _dafny.Sequence
		_ = _837_v64
		_837_v64 = _dafny.SeqOf((_835_v62).Difference((_836_v63).Dtor_cf57()))
		var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_828_v55), 0))
		_ = _index123
		var _rhs139 _dafny.Int = ((_this).F18()).Minus((_this).F18())
		_ = _rhs139
		var _rhs140 _dafny.Array = _this.F9()
		_ = _rhs140
		var _rhs141 _dafny.Set = (_837_v64).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)).Cardinality(), _dafny.IntOfUint32((_837_v64).Cardinality()))).Uint32()).(_dafny.Set)
		_ = _rhs141
		var _lhs115 *GlobalState = globalState
		_ = _lhs115
		var _lhs116 *C4 = _this
		_ = _lhs116
		var _lhs117 _dafny.Array = _828_v55
		_ = _lhs117
		var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_828_v55), 0))
		_ = _lhs118
		_lhs115.F5 = _rhs139
		_lhs116.F9_set_(_rhs140)
		(_lhs117).ArraySet1(_rhs141, (_lhs118).Int())
		r0 = Companion_Default___.Fm29((func() _dafny.Set {
			if p0 {
				return (_828_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_828_v55), 0))).Int()).(_dafny.Set)
			}
			return (_828_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_828_v55), 0))).Int()).(_dafny.Set)
		})(), globalState)
		r1 = p0
		return r0, r1
	}
}
func (_this *C4) M7(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) (_dafny.CodePoint, _dafny.Int, _dafny.Array, _dafny.Int) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _838_v0 _dafny.CodePoint
		_ = _838_v0
		_838_v0 = _dafny.CodePoint('v')
		var _839_v1 _dafny.MultiSet
		_ = _839_v1
		_839_v1 = _dafny.MultiSetOf(false)
		var _840_v2 _dafny.Sequence
		_ = _840_v2
		_840_v2 = _dafny.SeqOf((_this).F6())
		var _841_v3 _dafny.MultiSet
		_ = _841_v3
		_841_v3 = _dafny.MultiSetOf(_dafny.SeqOf((_this).F17(), (_839_v1).Cardinality()), _840_v2)
		var _842_v5 _dafny.Map
		_ = _842_v5
		_842_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_838_v0, (func() _dafny.Set {
			var _coll22 = _dafny.NewBuilder()
			_ = _coll22
			for _iter24 := _dafny.Iterate((_841_v3).Elements()); ; {
				_compr_22, _ok24 := _iter24()
				if !_ok24 {
					break
				}
				var _843_v4 _dafny.Sequence
				_843_v4 = interface{}(_compr_22).(_dafny.Sequence)
				if (_841_v3).Contains(_843_v4) {
					_coll22.Add(_843_v4)
				}
			}
			return _coll22.ToSet()
		}()).Cardinality())
		var _hi6 _dafny.Int = (func() _dafny.Int {
			if (_842_v5).Contains(_838_v0) {
				return (_842_v5).Get(_838_v0).(_dafny.Int)
			}
			return (func() _dafny.Map {
				var _coll23 = _dafny.NewMapBuilder()
				_ = _coll23
				for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(35), _dafny.IntOfInt64(958))); ; {
					_compr_23, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _844_v6 _dafny.Int
					_844_v6 = interface{}(_compr_23).(_dafny.Int)
					if ((_dafny.IntOfInt64(35)).Cmp(_844_v6) <= 0) && ((_844_v6).Cmp(_dafny.IntOfInt64(958)) < 0) {
						_coll23.Add((_844_v6).Minus((_this).F18()), true)
					}
				}
				return _coll23.ToMap()
			}()).Cardinality()
		})()
		_ = _hi6
		for _845_i0 := (_this).F6(); _845_i0.Cmp(_hi6) < 0; _845_i0 = _845_i0.Plus(_dafny.One) {
			var _846_v7 _dafny.Sequence
			_ = _846_v7
			_846_v7 = _dafny.UnicodeSeqOfUtf8Bytes("rrvwuy")
			var _847_v8 bool
			_ = _847_v8
			_847_v8 = false
			var _848_v9 _dafny.Map
			_ = _848_v9
			_848_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_847_v8, p1)
			var _849_v10 *C3
			_ = _849_v10
			var _nw119 *C3 = New_C3_()
			_ = _nw119
			_nw119.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_845_i0, (_this).F6()), _this.F9(), (_this).F10(), _845_i0)
			_849_v10 = _nw119
			var _850_v11 _dafny.Map
			_ = _850_v11
			_850_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_849_v10, (_this).F17())
			var _851_v12 _dafny.Array
			_ = _851_v12
			var _nwElement0_22 _dafny.Int = p1
			_ = _nwElement0_22
			var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(10))
			_ = _nw120
			(_nw120).ArraySet1(_nwElement0_22, 0)
			(_nw120).ArraySet1(_dafny.IntOfInt64(-173), 1)
			(_nw120).ArraySet1(_dafny.IntOfUint32((_846_v7).Cardinality()), 2)
			(_nw120).ArraySet1((_848_v9).Cardinality(), 3)
			(_nw120).ArraySet1((_850_v11).Cardinality(), 4)
			(_nw120).ArraySet1(_dafny.IntOfInt64(-505), 5)
			(_nw120).ArraySet1((_this).F6(), 6)
			(_nw120).ArraySet1(p1, 7)
			(_nw120).ArraySet1((_this).F18(), 8)
			(_nw120).ArraySet1(_dafny.IntOfInt64(-558), 9)
			_851_v12 = _nw120
			var _852_v13 D5
			_ = _852_v13
			_852_v13 = Companion_D5_.Create_DC17_(_840_v2, !(_847_v8))
			var _853_v14 _dafny.Map
			_ = _853_v14
			_853_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_851_v12, (func() bool {
				if _847_v8 {
					return _847_v8
				}
				return (_852_v13).Dtor_cf37()
			})())
			_853_v14 = (_853_v14).Update(_851_v12, _847_v8)
			if _847_v8 {
				var _854_v15 _dafny.Map
				_ = _854_v15
				_854_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(199))).Uint32(), func(coer69 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg69 _dafny.Int) interface{} {
						return coer69(arg69)
					}
				}(func(_855_i1 _dafny.Int) _dafny.Int {
					return (_this).F18()
				})), !(_847_v8) || (_847_v8))
				_854_v15 = _854_v15
				(globalState).F0 = _847_v8
				var _856_v16 D4
				_ = _856_v16
				_856_v16 = Companion_D4_.Create_DC14_(_847_v8, Companion_Default___.Fm41(globalState), _845_i0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_847_v8, _847_v8))
				var _857_v17 _dafny.Map
				_ = _857_v17
				_857_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_856_v16).Dtor_cf28())
				_857_v17 = (_857_v17).Update(((_857_v17).Merge(_857_v17)).Cardinality(), (_dafny.IntOfInt64(-185)).Cmp((_this).F18()) > 0)
				var _858_v18 _dafny.Array
				_ = _858_v18
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_18
				var _nw121 _dafny.Array
				_ = _nw121
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw121 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Map = (func(_859_v8 bool) func(_dafny.Int) _dafny.Map {
						return func(_860_i2 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_859_v8, _859_v8)
						}
					})(_847_v8)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw121 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw121).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw121).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_858_v18 = _nw121
				var _861_v19 _dafny.Set
				_ = _861_v19
				_861_v19 = _dafny.SetOf(Companion_D9_.Create_DC26_(_858_v18))
				var _862_v20 _dafny.Map
				_ = _862_v20
				_862_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_847_v8, _861_v19)
				var _863_v21 _dafny.Array
				_ = _863_v21
				var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(13))
				_ = _nw122
				_863_v21 = _nw122
				var _864_v22 D10
				_ = _864_v22
				_864_v22 = Companion_D10_.Create_DC28_(_853_v14)
				var _rhs142 _dafny.Map = ((_862_v20).Merge(_862_v20)).Merge(_862_v20)
				_ = _rhs142
				var _rhs143 _dafny.Array = _863_v21
				_ = _rhs143
				var _rhs144 D10 = _864_v22
				_ = _rhs144
				var _rhs145 bool = _847_v8
				_ = _rhs145
				var _lhs119 *GlobalState = globalState
				_ = _lhs119
				_862_v20 = _rhs142
				_863_v21 = _rhs143
				_864_v22 = _rhs144
				_lhs119.F0 = _rhs145
				_838_v0 = _838_v0
			} else {
				var _arr21 _dafny.Array = _this.F9()
				_ = _arr21
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index124
				(_arr21).ArraySet1(_847_v8, (_index124).Int())
				var _865_v23 _dafny.Map
				_ = _865_v23
				_865_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_847_v8), _847_v8)
				var _866_v24 D4
				_ = _866_v24
				_866_v24 = Companion_D4_.Create_DC14_(_847_v8, _865_v23, p1, _865_v23)
				var _arr22 _dafny.Array = _this.F9()
				_ = _arr22
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index125
				(_arr22).ArraySet1((!(_847_v8)) && ((_866_v24).Dtor_cf28()), (_index125).Int())
				var _867_v25 _dafny.Array
				_ = _867_v25
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_19
				var _nw123 _dafny.Array
				_ = _nw123
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw123 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) _dafny.Map = (func(_868_v10 *C3) func(_dafny.Int) _dafny.Map {
						return func(_869_i3 _dafny.Int) _dafny.Map {
							return ((_868_v10).F21()).Merge((_868_v10).F21())
						}
					})(_849_v10)
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw123 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw123).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw123).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_867_v25 = _nw123
				_867_v25 = _867_v25
				var _870_v26 _dafny.Sequence
				_ = _870_v26
				_870_v26 = _dafny.SeqOf(_dafny.MultiSetOf(_838_v0))
				var _871_v27 D4
				_ = _871_v27
				_871_v27 = Companion_D4_.Create_DC12_(_870_v26)
				var _872_v28 _dafny.Array
				_ = _872_v28
				var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(21))
				_ = _nw124
				_872_v28 = _nw124
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_872_v28), 0))
				_ = _index126
				(_872_v28).ArraySet1((_dafny.MultiSetOf(_840_v2)).Update(_dafny.SeqOf(_dafny.IntOfInt64(766)), Companion_Default___.Abs(p1)), (_index126).Int())
				var _873_v29 _dafny.Array
				_ = _873_v29
				var _nw125 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
				_ = _nw125
				_873_v29 = _nw125
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_872_v28), 0))
				_ = _index127
				var _rhs146 D4 = _871_v27
				_ = _rhs146
				var _rhs147 _dafny.MultiSet = _841_v3
				_ = _rhs147
				var _rhs148 _dafny.Array = _873_v29
				_ = _rhs148
				var _lhs120 _dafny.Array = _872_v28
				_ = _lhs120
				var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_872_v28), 0))
				_ = _lhs121
				_871_v27 = _rhs146
				(_lhs120).ArraySet1(_rhs147, (_lhs121).Int())
				_873_v29 = _rhs148
				r3 = (((func() _dafny.Map {
					if _847_v8 {
						return _865_v23
					}
					return _865_v23
				})()).Merge(_865_v23)).Cardinality()
				var _arr23 _dafny.Array = _this.F9()
				_ = _arr23
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index128
				var _arr24 _dafny.Array = _this.F9()
				_ = _arr24
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index129
				var _arr25 _dafny.Array = _this.F9()
				_ = _arr25
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index130
				var _rhs149 bool = _847_v8
				_ = _rhs149
				var _rhs150 bool = !((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
				_ = _rhs150
				var _rhs151 _dafny.CodePoint = (Companion_D0_.Create_DC1_(_847_v8, p1, _838_v0)).Dtor_cf4()
				_ = _rhs151
				var _rhs152 bool = _847_v8
				_ = _rhs152
				var _lhs122 _dafny.Array = _this.F9()
				_ = _lhs122
				var _lhs123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))
				_ = _lhs123
				var _lhs124 _dafny.Array = _this.F9()
				_ = _lhs124
				var _lhs125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))
				_ = _lhs125
				var _lhs126 _dafny.Array = _this.F9()
				_ = _lhs126
				var _lhs127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_this.F9()), 0))
				_ = _lhs127
				(_lhs122).ArraySet1(_rhs149, (_lhs123).Int())
				(_lhs124).ArraySet1(_rhs150, (_lhs125).Int())
				_838_v0 = _rhs151
				(_lhs126).ArraySet1(_rhs152, (_lhs127).Int())
			}
			var _874_v30 _dafny.Array
			_ = _874_v30
			var _nw126 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
			_ = _nw126
			_874_v30 = _nw126
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_874_v30), 0))
			_ = _index131
			(_874_v30).ArraySet1(_840_v2, (_index131).Int())
			var _875_v31 _dafny.Map
			_ = _875_v31
			_875_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _847_v8)
			var _876_v32 D4
			_ = _876_v32
			_876_v32 = Companion_D4_.Create_DC14_(_847_v8, _875_v31, (_this).F17(), _875_v31)
			var _877_v33 _dafny.Sequence
			_ = _877_v33
			_877_v33 = _dafny.SeqOf((_this).Fm22((_dafny.MultiSetOf(_847_v8)).Update((_876_v32).Dtor_cf28(), Companion_Default___.Abs((_this).F6())), false, (_this).F6(), globalState))
			var _878_v34 _dafny.Set
			_ = _878_v34
			_878_v34 = _dafny.SetOf(p1, _dafny.IntOfUint32((_877_v33).Cardinality()), (_839_v1).Cardinality(), (_this).F6())
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_874_v30), 0))
			_ = _index132
			(_874_v30).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_840_v2, _840_v2), _dafny.Companion_Sequence_.Concatenate(_840_v2, _dafny.SeqOf((_878_v34).Cardinality(), (_this).F17(), (_this).F18()))), (_index132).Int())
			r3 = ((p1).Times((_this).F18())).Plus((_this).F6())
		}
		var _879_v37 _dafny.Array
		_ = _879_v37
		var _len0_20 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_20
		var _nw127 _dafny.Array
		_ = _nw127
		if _len0_20.Cmp(_dafny.Zero) == 0 {
			_nw127 = _dafny.NewArray(_len0_20)
		} else {
			var _init20 func(_dafny.Int) _dafny.Map = (func(_880_v1 _dafny.MultiSet) func(_dafny.Int) _dafny.Map {
				return func(_881_i4 _dafny.Int) _dafny.Map {
					return func() _dafny.Map {
						var _coll24 = _dafny.NewMapBuilder()
						_ = _coll24
						for _iter26 := _dafny.Iterate((func() _dafny.Set {
							var _coll25 = _dafny.NewBuilder()
							_ = _coll25
							for _iter27 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("d"), true)).Keys().Elements()); ; {
								_compr_25, _ok27 := _iter27()
								if !_ok27 {
									break
								}
								var _882_v36 _dafny.Sequence
								_882_v36 = interface{}(_compr_25).(_dafny.Sequence)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("d"), true)).Contains(_882_v36) {
									_coll25.Add(_882_v36)
								}
							}
							return _coll25.ToSet()
						}()).Elements()); ; {
							_compr_24, _ok26 := _iter26()
							if !_ok26 {
								break
							}
							var _883_v35 _dafny.Sequence
							_883_v35 = interface{}(_compr_24).(_dafny.Sequence)
							if (func() _dafny.Set {
								var _coll26 = _dafny.NewBuilder()
								_ = _coll26
								for _iter28 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("d"), true)).Keys().Elements()); ; {
									_compr_26, _ok28 := _iter28()
									if !_ok28 {
										break
									}
									var _884_v36 _dafny.Sequence
									_884_v36 = interface{}(_compr_26).(_dafny.Sequence)
									if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("d"), true)).Contains(_884_v36) {
										_coll26.Add(_884_v36)
									}
								}
								return _coll26.ToSet()
							}()).Contains(_883_v35) {
								_coll24.Add(_883_v35, (_880_v1).Cardinality())
							}
						}
						return _coll24.ToMap()
					}()
				}
			})(_839_v1)
			_ = _init20
			var _element0_20 = _init20(_dafny.Zero)
			_ = _element0_20
			_nw127 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
			(_nw127).ArraySet1(_element0_20, 0)
			var _nativeLen0_20 = (_len0_20).Int()
			_ = _nativeLen0_20
			for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
				(_nw127).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
			}
		}
		_879_v37 = _nw127
		var _885_v38 _dafny.Sequence
		_ = _885_v38
		_885_v38 = _dafny.UnicodeSeqOfUtf8Bytes("jgfwo")
		var _886_v39 _dafny.Map
		_ = _886_v39
		_886_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_885_v38, (_this).F17())
		var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_879_v37), 0))
		_ = _index133
		(_879_v37).ArraySet1(_886_v39, (_index133).Int())
		var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_879_v37), 0))
		_ = _index134
		(_879_v37).ArraySet1(_886_v39, (_index134).Int())
		_839_v1 = _839_v1
		var _887_v40 bool
		_ = _887_v40
		_887_v40 = true
		var _888_v41 _dafny.Set
		_ = _888_v41
		_888_v41 = _dafny.SetOf(_887_v40)
		_888_v41 = ((_888_v41).Difference(_dafny.SetOf(false))).Difference((_888_v41).Difference(_dafny.SetOf(_887_v40)))
		var _hi7 _dafny.Int = _dafny.IntOfUint32((_885_v38).Cardinality())
		_ = _hi7
		for _889_i5 := (_this).F17(); _889_i5.Cmp(_hi7) < 0; _889_i5 = _889_i5.Plus(_dafny.One) {
			var _890_v42 _dafny.Map
			_ = _890_v42
			_890_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(115))
			var _891_v43 _dafny.Sequence
			_ = _891_v43
			_891_v43 = _dafny.SeqOf(_this.F9(), _this.F9())
			var _892_v44 _dafny.Sequence
			_ = _892_v44
			_892_v44 = _dafny.SeqOf(_this.F9(), (_891_v43).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(797), _dafny.IntOfUint32((_891_v43).Cardinality()))).Uint32()).(_dafny.Array), _this.F9(), _this.F9())
			var _893_v45 D4
			_ = _893_v45
			_893_v45 = Companion_D4_.Create_DC13_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-484))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg70 _dafny.Int) interface{} {
					return coer70(arg70)
				}
			}(func(_894_i6 _dafny.Int) _dafny.Int {
				return (_this).F6()
			})), _887_v40, (_this).F18())
			var _895_v46 _dafny.Map
			_ = _895_v46
			_895_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_887_v40, (_893_v45).Dtor_cf26())
			var _896_v47 *C3
			_ = _896_v47
			var _nw128 *C3 = New_C3_()
			_ = _nw128
			_nw128.Ctor__(_890_v42, (_892_v44).Select((Companion_Default___.SafeIndex((_895_v46).Cardinality(), _dafny.IntOfUint32((_892_v44).Cardinality()))).Uint32()).(_dafny.Array), (_this).F10(), _dafny.IntOfUint32((_885_v38).Cardinality()))
			_896_v47 = _nw128
			r1 = (_this).F6()
			var _897_v48 bool
			_ = _897_v48
			var _898_v49 bool
			_ = _898_v49
			var _899_v50 _dafny.Array
			_ = _899_v50
			var _out25 bool
			_ = _out25
			var _out26 bool
			_ = _out26
			var _out27 _dafny.Array
			_ = _out27
			_out25, _out26, _out27 = (_this).M12(_887_v40, (_dafny.SetOf(true)).IsSubsetOf(_888_v41), globalState)
			_897_v48 = _out25
			_898_v49 = _out26
			_899_v50 = _out27
			var _900_v51 _dafny.Set
			_ = _900_v51
			_900_v51 = _dafny.SetOf(_dafny.IntOfInt64(8), (_dafny.Zero).Minus((_dafny.Zero).Minus(p1)))
			if !(_900_v51).Equals(_900_v51) {
				var _901_v52 _dafny.Array
				_ = _901_v52
				_901_v52 = _this.F9()
				var _902_v53 *C0
				_ = _902_v53
				var _nw129 *C0 = New_C0_()
				_ = _nw129
				_nw129.Ctor__((_901_v52), _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm30(_887_v40, _838_v0, globalState), _885_v38))
				_902_v53 = _nw129
				var _rhs153 _dafny.Int = p1
				_ = _rhs153
				var _rhs154 _dafny.CodePoint = _dafny.CodePoint('o')
				_ = _rhs154
				r3 = _rhs153
				_838_v0 = _rhs154
				var _903_v55 *C3
				_ = _903_v55
				var _nw130 *C3 = New_C3_()
				_ = _nw130
				_nw130.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_this).F17()), _902_v53.F19, ((_this).F10()).Merge(func() _dafny.Map {
					var _coll27 = _dafny.NewMapBuilder()
					_ = _coll27
					for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(978), _dafny.IntOfInt64(814))); ; {
						_compr_27, _ok29 := _iter29()
						if !_ok29 {
							break
						}
						var _904_v54 _dafny.Int
						_904_v54 = interface{}(_compr_27).(_dafny.Int)
						if ((_dafny.IntOfInt64(978)).Cmp(_904_v54) <= 0) && ((_904_v54).Cmp(_dafny.IntOfInt64(814)) < 0) {
							_coll27.Add((_904_v54).Times(_dafny.IntOfInt64(456)), _885_v38)
						}
					}
					return _coll27.ToMap()
				}()), (_dafny.Zero).Minus((_this).F6()))
				_903_v55 = _nw130
				var _905_v56 _dafny.Map
				_ = _905_v56
				_905_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_903_v55, _dafny.IntOfUint32((_885_v38).Cardinality()))
				_890_v42 = (_890_v42).Update((_905_v56).Cardinality(), (_this).F6())
				var _906_v57 _dafny.Map
				_ = _906_v57
				_906_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_897_v48), _899_v50)
				var _907_v58 _dafny.Map
				_ = _907_v58
				_907_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_906_v57, (_this).F18())
				var _908_v59 _dafny.Map
				_ = _908_v59
				_908_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _906_v57)
				_907_v58 = (_907_v58).Update((func() _dafny.Map {
					if (_908_v59).Contains((_this).F6()) {
						return (_908_v59).Get((_this).F6()).(_dafny.Map)
					}
					return _906_v57
				})(), ((_this).F6()).Minus((_840_v2).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_840_v2).Cardinality()))).Uint32()).(_dafny.Int)))
			} else {
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_899_v50), 0))
				_ = _index135
				(_899_v50).ArraySet1((_839_v1).Cardinality(), (_index135).Int())
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_899_v50), 0))
				_ = _index136
				(_899_v50).ArraySet1((_this).F17(), (_index136).Int())
				var _909_v60 _dafny.Set
				_ = _909_v60
				_909_v60 = _dafny.SetOf(_838_v0)
				var _910_v61 _dafny.Map
				_ = _910_v61
				_910_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_dafny.SetOf(_838_v0, _838_v0, _838_v0)).Union(_909_v60))
				_910_v61 = func() _dafny.Map {
					var _coll28 = _dafny.NewMapBuilder()
					_ = _coll28
					for _iter30 := _dafny.Iterate((((_896_v47).F21()).Merge(Companion_Default___.Fm42(_898_v49, p0, (_dafny.Zero).Minus((_this).F6()), _838_v0, globalState))).Keys().Elements()); ; {
						_compr_28, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _911_v62 _dafny.Int
						_911_v62 = interface{}(_compr_28).(_dafny.Int)
						if (((_896_v47).F21()).Merge(Companion_Default___.Fm42(_898_v49, p0, (_dafny.Zero).Minus((_this).F6()), _838_v0, globalState))).Contains(_911_v62) {
							_coll28.Add((_911_v62).Times((func() _dafny.Map {
								var _coll29 = _dafny.NewMapBuilder()
								_ = _coll29
								for _iter31 := _dafny.Iterate((_dafny.MultiSetOf(Companion_D0_.Create_DC1_(_898_v49, _dafny.IntOfInt64(265), _838_v0), Companion_D0_.Create_DC1_(_897_v48, _dafny.IntOfUint32((_885_v38).Cardinality()), _838_v0), Companion_D0_.Create_DC1_(_887_v40, (_this).F18(), _dafny.CodePoint('g')), Companion_D0_.Create_DC1_(_897_v48, _dafny.IntOfInt64(376), _838_v0), Companion_D0_.Create_DC1_(_898_v49, (_this).F6(), _838_v0))).Elements()); ; {
									_compr_29, _ok31 := _iter31()
									if !_ok31 {
										break
									}
									var _912_v63 D0
									_912_v63 = interface{}(_compr_29).(D0)
									if (_dafny.MultiSetOf(Companion_D0_.Create_DC1_(_898_v49, _dafny.IntOfInt64(265), _838_v0), Companion_D0_.Create_DC1_(_897_v48, _dafny.IntOfUint32((_885_v38).Cardinality()), _838_v0), Companion_D0_.Create_DC1_(_887_v40, (_this).F18(), _dafny.CodePoint('g')), Companion_D0_.Create_DC1_(_897_v48, _dafny.IntOfInt64(376), _838_v0), Companion_D0_.Create_DC1_(_898_v49, (_this).F6(), _838_v0))).Contains(_912_v63) {
										_coll29.Add(_912_v63, (_dafny.Zero).Minus((_899_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_899_v50), 0))).Int()).(_dafny.Int)))
									}
								}
								return _coll29.ToMap()
							}()).Cardinality()), _dafny.SetOf(_838_v0, _838_v0))
						}
					}
					return _coll28.ToMap()
				}()
				var _913_v64 _dafny.Array
				_ = _913_v64
				var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
				_ = _nw131
				_913_v64 = _nw131
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_913_v64), 0))
				_ = _index137
				(_913_v64).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(948))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}((func(_914_v38 _dafny.Sequence, _915_v50 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_916_i7 _dafny.Int) _dafny.CodePoint {
						return (_914_v38).Select((Companion_Default___.SafeIndex((_915_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_915_v50), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_914_v38).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_885_v38, _899_v50))), (_index137).Int())
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_913_v64), 0))
				_ = _index138
				(_913_v64).ArraySet1(_885_v38, (_index138).Int())
				var _917_v65 *C0
				_ = _917_v65
				var _nw132 *C0 = New_C0_()
				_ = _nw132
				_nw132.Ctor__(_this.F9(), _887_v40)
				_917_v65 = _nw132
				var _918_v66 _dafny.Array
				_ = _918_v66
				var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
				_ = _nw133
				_918_v66 = _nw133
				var _919_v67 _dafny.Array
				_ = _919_v67
				var _nwElement0_23 _dafny.Int = (_this).F18()
				_ = _nwElement0_23
				var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(2))
				_ = _nw134
				(_nw134).ArraySet1(_nwElement0_23, 0)
				(_nw134).ArraySet1((_this).F6(), 1)
				_919_v67 = _nw134
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_918_v66), 0))
				_ = _index139
				(_918_v66).ArraySet1(_919_v67, (_index139).Int())
				var _920_v68 _dafny.Array
				_ = _920_v68
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_21
				var _nw135 _dafny.Array
				_ = _nw135
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw135 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Map = (func(_921_v65 *C0, _922_v49 bool) func(_dafny.Int) _dafny.Map {
						return func(_923_i8 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_921_v65.F20, _922_v49)
						}
					})(_917_v65, _898_v49)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw135 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw135).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw135).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_920_v68 = _nw135
				var _924_v69 D9
				_ = _924_v69
				_924_v69 = Companion_D9_.Create_DC26_(_920_v68)
				var _925_v70 _dafny.Sequence
				_ = _925_v70
				_925_v70 = _dafny.SeqOf(_917_v65.F20, _898_v49)
				var _926_v71 _dafny.Map
				_ = _926_v71
				_926_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_924_v69, _dafny.Companion_Sequence_.Update(_925_v70, (Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_925_v70).Cardinality()))).Uint32(), _887_v40))
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_918_v66), 0))
				_ = _index140
				var _nwElement0_24 _dafny.Int = p1
				_ = _nwElement0_24
				var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(12))
				_ = _nw136
				(_nw136).ArraySet1(_nwElement0_24, 0)
				(_nw136).ArraySet1((_926_v71).Cardinality(), 1)
				(_nw136).ArraySet1((_this).F18(), 2)
				(_nw136).ArraySet1(_dafny.IntOfInt64(566), 3)
				(_nw136).ArraySet1((func() _dafny.Int {
					if _897_v48 {
						return (_this).F6()
					}
					return (_dafny.Zero).Minus((_this).F17())
				})(), 4)
				(_nw136).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(114), _889_i5), 5)
				(_nw136).ArraySet1(_dafny.IntOfInt64(-915), 6)
				(_nw136).ArraySet1((_this).F17(), 7)
				(_nw136).ArraySet1((_this).F6(), 8)
				(_nw136).ArraySet1(p1, 9)
				(_nw136).ArraySet1((_899_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_899_v50), 0))).Int()).(_dafny.Int), 10)
				(_nw136).ArraySet1((_this).F6(), 11)
				(_918_v66).ArraySet1(_nw136, (_index140).Int())
			}
		}
		var _927_v72 _dafny.Map
		_ = _927_v72
		_927_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9(), _885_v38)
		var _928_i9 _dafny.Int
		_ = _928_i9
		_928_i9 = _dafny.Zero
		{
			for !(_927_v72).Equals(_927_v72) {
				{
					if (_928_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_928_i9 = (_928_i9).Plus(_dafny.One)
					var _929_v73 _dafny.Sequence
					_ = _929_v73
					_929_v73 = _dafny.SeqOf(_887_v40)
					var _930_v74 _dafny.Map
					_ = _930_v74
					_930_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_885_v38).Cardinality()), _929_v73)
					var _931_v75 _dafny.Map
					_ = _931_v75
					_931_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _887_v40)
					var _932_v76 D5
					_ = _932_v76
					_932_v76 = Companion_D5_.Create_DC17_(_840_v2, _887_v40)
					var _933_v77 _dafny.Sequence
					_ = _933_v77
					_933_v77 = _dafny.SeqOf(Companion_D5_.Create_DC17_(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(868))).Uint32(), func(coer72 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg72 _dafny.Int) interface{} {
							return coer72(arg72)
						}
					}((func(_934_v40 bool) func(_dafny.Int) _dafny.Int {
						return func(_935_i10 _dafny.Int) _dafny.Int {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_934_v40, (_this).F18())).Cardinality()
						}
					})(_887_v40))), (Companion_Default___.SafeIndex((_930_v74).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(868))).Uint32(), func(coer73 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg73 _dafny.Int) interface{} {
							return coer73(arg73)
						}
					}((func(_936_v40 bool) func(_dafny.Int) _dafny.Int {
						return func(_937_i10 _dafny.Int) _dafny.Int {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_936_v40, (_this).F18())).Cardinality()
						}
					})(_887_v40)))).Cardinality()))).Uint32(), (_931_v75).Cardinality()), !(_887_v40)), Companion_D5_.Create_DC17_(_840_v2, _887_v40), _932_v76)
					var _938_v78 _dafny.Sequence
					_ = _938_v78
					_938_v78 = _dafny.SeqOf(Companion_Default___.Fm35((_this).F17(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), _this.F9())).Cardinality(), _887_v40, (_888_v41).Cardinality(), globalState))
					var _939_v79 _dafny.Map
					_ = _939_v79
					_939_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_933_v77, (func() _dafny.Int {
						if _887_v40 {
							return (_this).F17()
						}
						return (((_938_v78).Select((Companion_Default___.SafeIndex((_this).F17(), _dafny.IntOfUint32((_938_v78).Cardinality()))).Uint32()).(_dafny.MultiSet)).Update((_this).F17(), Companion_Default___.Abs((_this).F6()))).Cardinality()
					})())
					_939_v79 = _939_v79
					var _940_v80 _dafny.Set
					_ = _940_v80
					_940_v80 = _dafny.SetOf((_this).F6())
					var _941_v81 _dafny.Sequence
					_ = _941_v81
					_941_v81 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_885_v38, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_885_v38).Cardinality()))).Uint32(), Companion_Default___.Fm0((_this).F6(), _940_v80, (_929_v73).Select((Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_929_v73).Cardinality()))).Uint32()).(bool), globalState)))
					var _942_v82 _dafny.Array
					_ = _942_v82
					var _nwElement0_25 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_885_v38, _dafny.UnicodeSeqOfUtf8Bytes("dpxwvvn"))
					_ = _nwElement0_25
					var _nw137 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(17))
					_ = _nw137
					(_nw137).ArraySet1(_nwElement0_25, 0)
					(_nw137).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("oplylgwjk"), 1)
					(_nw137).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(439))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg74 _dafny.Int) interface{} {
							return coer74(arg74)
						}
					}((func(_943_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_944_i11 _dafny.Int) _dafny.CodePoint {
							return _943_v0
						}
					})(_838_v0))), 2)
					(_nw137).ArraySet1(_885_v38, 3)
					(_nw137).ArraySet1(_885_v38, 4)
					(_nw137).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("huf"), 5)
					(_nw137).ArraySet1(Companion_Default___.Fm30(_887_v40, _dafny.CodePoint('n'), globalState), 6)
					(_nw137).ArraySet1(_885_v38, 7)
					(_nw137).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("a"), _885_v38), 8)
					(_nw137).ArraySet1(_885_v38, 9)
					(_nw137).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("xkaku"), 10)
					(_nw137).ArraySet1(_885_v38, 11)
					(_nw137).ArraySet1(_dafny.Companion_Sequence_.Update((_941_v81).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_929_v73).Cardinality()), _dafny.IntOfUint32((_941_v81).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32(((_941_v81).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_929_v73).Cardinality()), _dafny.IntOfUint32((_941_v81).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _838_v0), 12)
					(_nw137).ArraySet1(Companion_Default___.Fm25((_this).F17(), _887_v40, globalState), 13)
					(_nw137).ArraySet1(Companion_Default___.Fm30(_887_v40, _838_v0, globalState), 14)
					(_nw137).ArraySet1(_885_v38, 15)
					(_nw137).ArraySet1(_885_v38, 16)
					_942_v82 = _nw137
					var _945_v83 _dafny.Sequence
					_ = _945_v83
					_945_v83 = _dafny.SeqOf(_942_v82, _942_v82, _942_v82, _942_v82)
					_942_v82 = (_945_v83).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.IntOfUint32((_945_v83).Cardinality()))).Uint32()).(_dafny.Array)
					_887_v40 = (_dafny.IntOfInt64(-609)).Cmp((_this).F18()) <= 0
					if !(_887_v40) {
						var _946_v84 _dafny.Sequence
						_ = _946_v84
						_946_v84 = _dafny.SeqOf(Companion_D5_.Create_DC16_(_888_v41))
						var _947_v85 _dafny.Set
						_ = _947_v85
						_947_v85 = _dafny.SetOf(_946_v84)
						var _948_v86 _dafny.Map
						_ = _948_v86
						_948_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_946_v84, (_this).F6())
						_947_v85 = func() _dafny.Set {
							var _coll30 = _dafny.NewBuilder()
							_ = _coll30
							for _iter32 := _dafny.Iterate((_948_v86).Keys().Elements()); ; {
								_compr_30, _ok32 := _iter32()
								if !_ok32 {
									break
								}
								var _949_v87 _dafny.Sequence
								_949_v87 = interface{}(_compr_30).(_dafny.Sequence)
								if (_948_v86).Contains(_949_v87) {
									_coll30.Add(_949_v87)
								}
							}
							return _coll30.ToSet()
						}()
						var _950_v88 D9
						_ = _950_v88
						_950_v88 = Companion_D9_.Create_DC25_((func() _dafny.Int {
							if _887_v40 {
								return p1
							}
							return (_this).F6()
						})(), (func() _dafny.Int {
							if (_839_v1).Contains(_887_v40) {
								return (_839_v1).Multiplicity(_887_v40)
							}
							return _dafny.IntOfInt64(139)
						})(), (func() _dafny.Int {
							if _887_v40 {
								return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-255))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
									return func(arg75 _dafny.Int) interface{} {
										return coer75(arg75)
									}
								}(func(_951_i12 _dafny.Int) _dafny.Int {
									return (_this).F18()
								}))).Cardinality())
							}
							return (_dafny.Zero).Minus((_this).F18())
						})(), _838_v0)
						_950_v88 = _950_v88
						var _952_v89 D0
						_ = _952_v89
						_952_v89 = Companion_D0_.Create_DC0_(_929_v73, (_this).F18())
						var _rhs155 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_885_v38, _885_v38)
						_ = _rhs155
						var _rhs156 _dafny.Int = ((func() _dafny.Int {
							if _887_v40 {
								return (_952_v89).Dtor_cf1()
							}
							return (_this).F18()
						})()).Plus(((_this).F17()).Times((_this).F17()))
						_ = _rhs156
						var _lhs128 *GlobalState = globalState
						_ = _lhs128
						_lhs128.F2 = _rhs155
						r3 = _rhs156
						var _953_v90 _dafny.Array
						_ = _953_v90
						var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
						_ = _nw138
						_953_v90 = _nw138
						var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_953_v90), 0))
						_ = _index141
						(_953_v90).ArraySet1((_this).F18(), (_index141).Int())
						var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_953_v90), 0))
						_ = _index142
						var _rhs157 bool = _887_v40
						_ = _rhs157
						var _rhs158 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-824), p1)
						_ = _rhs158
						var _rhs159 _dafny.Set = (_dafny.SetOf(_887_v40)).Difference(_888_v41)
						_ = _rhs159
						var _rhs160 _dafny.Int = (Companion_D4_.Create_DC13_(_840_v2, _887_v40, (_dafny.Zero).Minus((_this).F17()))).Dtor_cf27()
						_ = _rhs160
						var _lhs129 *GlobalState = globalState
						_ = _lhs129
						var _lhs130 _dafny.Array = _953_v90
						_ = _lhs130
						var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_953_v90), 0))
						_ = _lhs131
						_lhs129.F0 = _rhs157
						r3 = _rhs158
						_888_v41 = _rhs159
						(_lhs130).ArraySet1(_rhs160, (_lhs131).Int())
						r3 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_885_v38, _885_v38)).Cardinality())
					} else {
						(globalState).F0 = _887_v40
						var _arr26 _dafny.Array = _this.F9()
						_ = _arr26
						var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_this.F9()), 0))
						_ = _index143
						(_arr26).ArraySet1(_887_v40, (_index143).Int())
						var _954_v91 _dafny.Sequence
						_ = _954_v91
						_954_v91 = _dafny.SeqOf(_840_v2, _840_v2)
						var _arr27 _dafny.Array = _this.F9()
						_ = _arr27
						var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_this.F9()), 0))
						_ = _index144
						var _rhs161 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("afv"), _885_v38)).Cardinality())).Cmp(Companion_Default___.SafeDivisionInt((_this).F18(), (_dafny.Zero).Minus(p1))) > 0
						_ = _rhs161
						var _rhs162 _dafny.Sequence = _954_v91
						_ = _rhs162
						var _rhs163 _dafny.Int = (_this).F6()
						_ = _rhs163
						var _lhs132 _dafny.Array = _this.F9()
						_ = _lhs132
						var _lhs133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_this.F9()), 0))
						_ = _lhs133
						var _lhs134 *GlobalState = globalState
						_ = _lhs134
						(_lhs132).ArraySet1(_rhs161, (_lhs133).Int())
						_954_v91 = _rhs162
						_lhs134.F5 = _rhs163
						var _arr28 _dafny.Array = _this.F9()
						_ = _arr28
						var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_this.F9()), 0))
						_ = _index145
						(_arr28).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), (_index145).Int())
						var _955_v92 _dafny.MultiSet
						_ = _955_v92
						_955_v92 = _839_v1
						var _956_v93 _dafny.Map
						_ = _956_v93
						_956_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_955_v92))
						_839_v1 = (_839_v1).Union(((_839_v1).Update(false, Companion_Default___.Abs((_this).F17()))).Difference((func() _dafny.MultiSet {
							if (_956_v93).Contains((_this).Fm21(_dafny.IntOfUint32((_929_v73).Cardinality()), globalState)) {
								return (_956_v93).Get((_this).Fm21(_dafny.IntOfUint32((_929_v73).Cardinality()), globalState)).(_dafny.MultiSet)
							}
							return _dafny.MultiSetOf(_887_v40, (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
						})()))
						(globalState).F5 = (_dafny.Zero).Minus((_this).F6())
					}
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _957_v94 D7
		_ = _957_v94
		_957_v94 = Companion_D7_.Create_DC22_()
		var _pat_let_tv18 = _838_v0
		_ = _pat_let_tv18
		r0 = func(_source12 D7) _dafny.CodePoint {
			if _source12.Is_DC22() {
				return _pat_let_tv18
			} else {
				var _958___mcc_h0 _dafny.Array = _source12.Get_().(D7_DC21).Cf43
				_ = _958___mcc_h0
				var _959_cf43 _dafny.Array = _958___mcc_h0
				_ = _959_cf43
				return _dafny.CodePoint('r')
			}
		}((func() D7 {
			if _887_v40 {
				return _957_v94
			}
			return _957_v94
		})())
		var _960_v95 _dafny.MultiSet
		_ = _960_v95
		_960_v95 = _dafny.MultiSetOf(_this.F9())
		r1 = (func() _dafny.Int {
			if (_960_v95).Contains(_this.F9()) {
				return (_960_v95).Multiplicity(_this.F9())
			}
			return (_this).F18()
		})()
		r2 = _this.F9()
		r3 = (_this).F17()
		return r0, r1, r2, r3
	}
}
func (_this *C4) M12(p0 bool, p1 bool, globalState *GlobalState) (bool, bool, _dafny.Array) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var _961_v0 _dafny.Sequence
		_ = _961_v0
		_961_v0 = _dafny.SeqOf((_this).F6())
		var _962_v1 _dafny.Sequence
		_ = _962_v1
		_962_v1 = _dafny.SeqOf(((_this).F18()).Minus((_this).F17()), (_961_v0).Select((Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_961_v0).Cardinality()))).Uint32()).(_dafny.Int), (_this).F18(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(880), (_this).F6()), ((_this).F6()).Plus((_this).F18()))
		var _963_v2 _dafny.Sequence
		_ = _963_v2
		_963_v2 = _dafny.UnicodeSeqOfUtf8Bytes("xxnkqk")
		var _rhs164 _dafny.Sequence = _963_v2
		_ = _rhs164
		var _rhs165 _dafny.Sequence = _961_v0
		_ = _rhs165
		var _rhs166 _dafny.Int = (_dafny.IntOfInt64(899)).Minus((_this).F17())
		_ = _rhs166
		var _lhs135 *GlobalState = globalState
		_ = _lhs135
		var _lhs136 *GlobalState = globalState
		_ = _lhs136
		_lhs135.F2 = _rhs164
		_961_v0 = _rhs165
		_lhs136.F5 = _rhs166
		var _arr29 _dafny.Array = _this.F9()
		_ = _arr29
		var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_this.F9()), 0))
		_ = _index146
		(_arr29).ArraySet1(p0, (_index146).Int())
		var _964_v3 _dafny.Map
		_ = _964_v3
		_964_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(240), p0)
		var _965_v4 _dafny.Map
		_ = _965_v4
		_965_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm22((_dafny.MultiSetOf(!(true))), (func() bool {
			if (_964_v3).Contains((_this).F17()) {
				return (_964_v3).Get((_this).F17()).(bool)
			}
			return p1
		})(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F17(), (_this).F17(), (_this).F17())).Cardinality()), globalState), !(p1) || (p1))
		var _966_v5 _dafny.Map
		_ = _966_v5
		_966_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F6())
		var _967_v6 _dafny.MultiSet
		_ = _967_v6
		_967_v6 = _dafny.MultiSetOf((_this).F18())
		var _968_v8 _dafny.Set
		_ = _968_v8
		_968_v8 = _dafny.SetOf((_this).F6())
		var _969_v9 *C2
		_ = _969_v9
		var _nw139 *C2 = New_C2_()
		_ = _nw139
		_nw139.Ctor__(((_966_v5).Merge(_966_v5)).Cardinality(), _dafny.MultiSetOf(_967_v6), _this.F9(), (func() _dafny.Map {
			var _coll31 = _dafny.NewMapBuilder()
			_ = _coll31
			for _iter33 := _dafny.Iterate((_968_v8).Elements()); ; {
				_compr_31, _ok33 := _iter33()
				if !_ok33 {
					break
				}
				var _970_v7 _dafny.Int
				_970_v7 = interface{}(_compr_31).(_dafny.Int)
				if (_968_v8).Contains(_970_v7) {
					_coll31.Add(Companion_Default___.SafeModuloInt(_970_v7, _dafny.IntOfInt64(-615)), _dafny.UnicodeSeqOfUtf8Bytes("m"))
				}
			}
			return _coll31.ToMap()
		}()).Update((_this).F17(), _963_v2))
		_969_v9 = _nw139
		var _971_v10 _dafny.Sequence
		_ = _971_v10
		_971_v10 = _dafny.SeqOf(_964_v3)
		var _972_v11 _dafny.Sequence
		_ = _972_v11
		_972_v11 = _dafny.SeqOf(true)
		var _973_v12 _dafny.Map
		_ = _973_v12
		_973_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9(), _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_972_v11, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_963_v2).Cardinality()), _dafny.IntOfUint32((_972_v11).Cardinality()))).Uint32(), p1)))
		var _974_v13 _dafny.Sequence
		_ = _974_v13
		_974_v13 = _dafny.SeqOf(((_971_v10).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_971_v10).Cardinality()))).Uint32()).(_dafny.Map)).Update((_973_v12).Cardinality(), p1), _964_v3)
		var _975_v14 _dafny.Set
		_ = _975_v14
		_975_v14 = _dafny.SetOf(!(false), ((_this).F6()).Cmp((_this).F17()) > 0, _dafny.Companion_Sequence_.Equal(_971_v10, _974_v13))
		var _976_v15 _dafny.CodePoint
		_ = _976_v15
		_976_v15 = _dafny.CodePoint('y')
		var _977_v16 _dafny.Map
		_ = _977_v16
		_977_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_976_v15, !((_972_v11).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_972_v11).Cardinality()))).Uint32()).(bool)))
		var _arr30 _dafny.Array = _this.F9()
		_ = _arr30
		var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_this.F9()), 0))
		_ = _index147
		var _rhs167 bool = (func() bool {
			if (_977_v16).Contains(_976_v15) {
				return (_977_v16).Get(_976_v15).(bool)
			}
			return Companion_Default___.Fm3(globalState)
		})()
		_ = _rhs167
		var _rhs168 _dafny.Map = _965_v4
		_ = _rhs168
		var _rhs169 *C2 = _969_v9
		_ = _rhs169
		var _rhs170 _dafny.Set = _975_v14
		_ = _rhs170
		var _rhs171 bool = p1
		_ = _rhs171
		var _lhs137 _dafny.Array = _this.F9()
		_ = _lhs137
		var _lhs138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_this.F9()), 0))
		_ = _lhs138
		(_lhs137).ArraySet1(_rhs167, (_lhs138).Int())
		_965_v4 = _rhs168
		_969_v9 = _rhs169
		_975_v14 = _rhs170
		r1 = _rhs171
		var _hi8 _dafny.Int = (func() _dafny.Int {
			if !(p1) {
				return _dafny.IntOfUint32((_972_v11).Cardinality())
			}
			return (_this).F6()
		})()
		_ = _hi8
		for _978_i0 := (_this).F18(); _978_i0.Cmp(_hi8) < 0; _978_i0 = _978_i0.Plus(_dafny.One) {
			if _dafny.Companion_Sequence_.Contains(_972_v11, ((_dafny.Zero).Minus((_969_v9).Fm21((_967_v6).Cardinality(), globalState))).Cmp((_dafny.Zero).Minus((_this).F17())) != 0) {
				(globalState).F5 = (func() _dafny.Int {
					if p0 {
						return _978_i0
					}
					return _978_i0
				})()
				(globalState).F0 = p0
				var _979_v17 _dafny.Array
				_ = _979_v17
				var _nw140 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
				_ = _nw140
				_979_v17 = _nw140
				var _980_v18 _dafny.Map
				_ = _980_v18
				_980_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_979_v17, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F17())).Cardinality(), (_this).F17()))
				_980_v18 = _980_v18
				(globalState).F0 = !(!(p1))
				var _981_v19 _dafny.Sequence
				_ = _981_v19
				_981_v19 = _dafny.SeqOf(_961_v0, _961_v0, _961_v0, _dafny.SeqOf((_this).F6()))
				var _982_v20 _dafny.Map
				_ = _982_v20
				_982_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_981_v19, Companion_Default___.Fm43(_976_v15, (_this).F17(), globalState)), (_this).F6())
				_982_v20 = (_982_v20).Update(_981_v19, (_this).F17())
			} else {
				var _983_v21 _dafny.Map
				_ = _983_v21
				_983_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), (_this).F18())
				(globalState).F5 = (Companion_D0_.Create_DC1_(p1, (_dafny.SetOf((_this).F6(), (_983_v21).Cardinality())).Cardinality(), _976_v15)).Dtor_cf3()
				var _984_v22 _dafny.Array
				_ = _984_v22
				var _nw141 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw141
				_984_v22 = _nw141
				var _985_v23 *C2
				_ = _985_v23
				var _nw142 *C2 = New_C2_()
				_ = _nw142
				_nw142.Ctor__((_dafny.Zero).Minus((_this).F18()), (_this.F16()).Intersection(_this.F16()), _984_v22, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _dafny.Companion_Sequence_.Update(_963_v2, (Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_963_v2).Cardinality()))).Uint32(), _976_v15)))
				_985_v23 = _nw142
				var _986_v24 D3
				_ = _986_v24
				_986_v24 = Companion_D3_.Create_DC9_(_dafny.Companion_Sequence_.Update(_962_v1, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_962_v1).Cardinality()))).Uint32(), Companion_Default___.Fm2(_978_i0, globalState)))
				var _987_v25 *C2
				_ = _987_v25
				var _nw143 *C2 = New_C2_()
				_ = _nw143
				_nw143.Ctor__((_this).F18(), (_this.F16()).Update(_dafny.MultiSetFromSeq((_986_v24).Dtor_cf21()), Companion_Default___.Abs(_dafny.IntOfUint32((_963_v2).Cardinality()))), (_984_v22), (_this).F10())
				_987_v25 = _nw143
				(globalState).F0 = p1
				_976_v15 = (func() _dafny.CodePoint {
					if (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool) {
						return _976_v15
					}
					return _976_v15
				})()
			}
			(globalState).F5 = _978_i0
			var _988_v26 _dafny.Map
			_ = _988_v26
			_988_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _978_i0)
			_988_v26 = (_988_v26).Update(((_this).F18()).Cmp((_this).F18()) > 0, _978_i0)
			var _989_v27 D2
			_ = _989_v27
			var _990_v28 _dafny.Sequence
			_ = _990_v28
			var _991_v29 _dafny.Sequence
			_ = _991_v29
			var _out28 D2
			_ = _out28
			var _out29 _dafny.Sequence
			_ = _out29
			var _out30 _dafny.Sequence
			_ = _out30
			_out28, _out29, _out30 = (_this).M13((_this).F18(), globalState)
			_989_v27 = _out28
			_990_v28 = _out29
			_991_v29 = _out30
		}
		var _992_v30 _dafny.Array
		_ = _992_v30
		var _nw144 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
		_ = _nw144
		_992_v30 = _nw144
		var _993_v31 *C0
		_ = _993_v31
		var _nw145 *C0 = New_C0_()
		_ = _nw145
		_nw145.Ctor__(_992_v30, (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
		_993_v31 = _nw145
		var _994_v32 _dafny.Map
		_ = _994_v32
		_994_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F17())
		var _995_v33 _dafny.MultiSet
		_ = _995_v33
		_995_v33 = _dafny.MultiSetOf(p1)
		_967_v6 = Companion_Default___.Fm35((_994_v32).Cardinality(), (_this).F17(), ((_this).F18()).Cmp((_967_v6).Cardinality()) == 0, (func() _dafny.Int {
			if (_995_v33).Contains(p1) {
				return (_995_v33).Multiplicity(p1)
			}
			return (_this).F17()
		})(), globalState)
		var _arr31 _dafny.Array = _this.F9()
		_ = _arr31
		var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_this.F9()), 0))
		_ = _index148
		var _rhs172 bool = false
		_ = _rhs172
		var _rhs173 _dafny.CodePoint = _976_v15
		_ = _rhs173
		var _lhs139 _dafny.Array = _this.F9()
		_ = _lhs139
		var _lhs140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_this.F9()), 0))
		_ = _lhs140
		(_lhs139).ArraySet1(_rhs172, (_lhs140).Int())
		_976_v15 = _rhs173
		r0 = p1
		r1 = p0
		var _996_v34 _dafny.Array
		_ = _996_v34
		var _len0_22 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_22
		var _nw146 _dafny.Array
		_ = _nw146
		if _len0_22.Cmp(_dafny.Zero) == 0 {
			_nw146 = _dafny.NewArray(_len0_22)
		} else {
			var _init22 func(_dafny.Int) _dafny.Int = (func(_997_v33 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
				return func(_998_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_998_i1, (_997_v33).Cardinality())
				}
			})(_995_v33)
			_ = _init22
			var _element0_22 = _init22(_dafny.Zero)
			_ = _element0_22
			_nw146 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
			(_nw146).ArraySet1(_element0_22, 0)
			var _nativeLen0_22 = (_len0_22).Int()
			_ = _nativeLen0_22
			for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
				(_nw146).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
			}
		}
		_996_v34 = _nw146
		r2 = _996_v34
		return r0, r1, r2
	}
}
func (_this *C4) M13(p0 _dafny.Int, globalState *GlobalState) (D2, _dafny.Sequence, _dafny.Sequence) {
	{
		var r0 D2 = Companion_D2_.Default()
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		if ((_this).F17()).Cmp(Companion_Default___.SafeDivisionInt(p0, p0)) == 0 {
			var _999_v0 bool
			_ = _999_v0
			_999_v0 = false
			var _1000_v1 _dafny.CodePoint
			_ = _1000_v1
			_1000_v1 = _dafny.CodePoint('j')
			var _1001_v2 _dafny.Sequence
			_ = _1001_v2
			_1001_v2 = _dafny.SeqOf((_this).F18())
			var _1002_v3 D4
			_ = _1002_v3
			_1002_v3 = Companion_D4_.Create_DC13_(_1001_v2, _999_v0, (_this).F18())
			var _1003_v4 _dafny.MultiSet
			_ = _1003_v4
			_1003_v4 = _dafny.MultiSetOf((_1002_v3).Dtor_cf26())
			var _1004_v5 _dafny.Set
			_ = _1004_v5
			_1004_v5 = _dafny.SetOf(_999_v0, true, _999_v0)
			var _1005_v6 _dafny.Map
			_ = _1005_v6
			_1005_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
				if _999_v0 {
					return _1000_v1
				}
				return _dafny.CodePoint('a')
			})(), Companion_Default___.Fm44(p0, (_1003_v4).Cardinality(), _999_v0, (_1004_v5).Cardinality(), globalState))
			var _1006_v7 D9
			_ = _1006_v7
			_1006_v7 = Companion_D9_.Create_DC25_((_dafny.Zero).Minus((_this).F17()), _dafny.IntOfUint32((_1001_v2).Cardinality()), (_this).F17(), _1000_v1)
			_1005_v6 = (_1005_v6).Update(_1000_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1006_v7))
			if false {
				var _1007_v8 *C2
				_ = _1007_v8
				var _nw147 *C2 = New_C2_()
				_ = _nw147
				_nw147.Ctor__(_dafny.IntOfInt64(-796), _this.F16(), _this.F9(), (_this).F10())
				_1007_v8 = _nw147
				var _1008_v9 _dafny.Map
				_ = _1008_v9
				_1008_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1007_v8, _999_v0)
				var _1009_v10 _dafny.Sequence
				_ = _1009_v10
				_1009_v10 = _dafny.SeqOf(!(_999_v0))
				var _1010_v11 _dafny.Map
				_ = _1010_v11
				_1010_v11 = _1008_v9
				var _1011_v12 _dafny.Array
				_ = _1011_v12
				var _nwElement0_26 _dafny.Map = _1008_v9
				_ = _nwElement0_26
				var _nw148 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(28))
				_ = _nw148
				(_nw148).ArraySet1(_nwElement0_26, 0)
				(_nw148).ArraySet1(_1008_v9, 1)
				(_nw148).ArraySet1(_1008_v9, 2)
				(_nw148).ArraySet1(_1008_v9, 3)
				(_nw148).ArraySet1((func() _dafny.Map {
					if _999_v0 {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1007_v8, _999_v0)
					}
					return _1008_v9
				})(), 4)
				(_nw148).ArraySet1(_1008_v9, 5)
				(_nw148).ArraySet1(_1008_v9, 6)
				(_nw148).ArraySet1((_1008_v9).Merge(_1008_v9), 7)
				(_nw148).ArraySet1(_1008_v9, 8)
				(_nw148).ArraySet1(_1008_v9, 9)
				(_nw148).ArraySet1(_1008_v9, 10)
				(_nw148).ArraySet1((func() _dafny.Map {
					if true {
						return _1008_v9
					}
					return _1008_v9
				})(), 11)
				(_nw148).ArraySet1((_1008_v9).Update(_1007_v8, _999_v0), 12)
				(_nw148).ArraySet1((_1008_v9).Update(_1007_v8, true), 13)
				(_nw148).ArraySet1(_1008_v9, 14)
				(_nw148).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1007_v8, _999_v0)).Merge(_1008_v9), 15)
				(_nw148).ArraySet1((_1008_v9).Merge((_1008_v9).Update(_1007_v8, (_1009_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tk")).Cardinality()), _dafny.IntOfUint32((_1009_v10).Cardinality()))).Uint32()).(bool))), 16)
				(_nw148).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1007_v8, _999_v0)).Merge(_1008_v9), 17)
				(_nw148).ArraySet1((_1010_v11), 18)
				(_nw148).ArraySet1((_1008_v9).Update(_1007_v8, _999_v0), 19)
				(_nw148).ArraySet1(_1008_v9, 20)
				(_nw148).ArraySet1(_1008_v9, 21)
				(_nw148).ArraySet1((_1008_v9).Merge(_1008_v9), 22)
				(_nw148).ArraySet1(_1008_v9, 23)
				(_nw148).ArraySet1(_1008_v9, 24)
				(_nw148).ArraySet1(_1008_v9, 25)
				(_nw148).ArraySet1(_1008_v9, 26)
				(_nw148).ArraySet1(_1008_v9, 27)
				_1011_v12 = _nw148
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1011_v12), 0))
				_ = _index149
				(_1011_v12).ArraySet1((_1008_v9).Merge(_1008_v9), (_index149).Int())
				var _1012_v13 _dafny.Map
				_ = _1012_v13
				_1012_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_999_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_999_v0, true))
				var _1013_v14 _dafny.Map
				_ = _1013_v14
				_1013_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_999_v0, false)
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1011_v12), 0))
				_ = _index150
				var _rhs174 _dafny.Map = _1008_v9
				_ = _rhs174
				var _rhs175 D4 = Companion_Default___.Fm45((_dafny.IntOfUint32((_1001_v2).Cardinality())).Minus((_this).F17()), (_this).F6(), globalState)
				_ = _rhs175
				var _rhs176 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1009_v10, _dafny.Companion_Sequence_.Concatenate(_1009_v10, Companion_Default___.Fm4((_this).F18(), p0, ((func() _dafny.Map {
					if (_1012_v13).Contains(_999_v0) {
						return (_1012_v13).Get(_999_v0).(_dafny.Map)
					}
					return _1013_v14
				})()).Cardinality(), globalState)))
				_ = _rhs176
				var _rhs177 bool = _999_v0
				_ = _rhs177
				var _rhs178 bool = _999_v0
				_ = _rhs178
				var _lhs141 _dafny.Array = _1011_v12
				_ = _lhs141
				var _lhs142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1011_v12), 0))
				_ = _lhs142
				var _lhs143 *GlobalState = globalState
				_ = _lhs143
				var _lhs144 *GlobalState = globalState
				_ = _lhs144
				(_lhs141).ArraySet1(_rhs174, (_lhs142).Int())
				_1002_v3 = _rhs175
				_1009_v10 = _rhs176
				_lhs143.F0 = _rhs177
				_lhs144.F0 = _rhs178
				var _1014_v15 *C1
				_ = _1014_v15
				var _nw149 *C1 = New_C1_()
				_ = _nw149
				_nw149.Ctor__(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uemcv")).Cardinality()))
				_1014_v15 = _nw149
				var _1015_v16 _dafny.Array
				_ = _1015_v16
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_23
				var _nw150 _dafny.Array
				_ = _nw150
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw150 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Int = func(_1016_i0 _dafny.Int) _dafny.Int {
						return (_1016_i0).Times((_dafny.Zero).Minus((_this).F17()))
					}
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw150 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw150).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw150).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_1015_v16 = _nw150
				var _1017_v17 D2
				_ = _1017_v17
				_1017_v17 = Companion_D2_.Create_DC6_((_this).F17(), _1015_v16)
				_1015_v16 = (_1017_v17).Dtor_cf14()
				var _1018_v18 *C1
				_ = _1018_v18
				var _nw151 *C1 = New_C1_()
				_ = _nw151
				_nw151.Ctor__((_this).F18())
				_1018_v18 = _nw151
				(_this).F9_set_(_this.F9())
			} else {
				var _1019_v19 T1
				_ = _1019_v19
				var _nw152 *C2 = New_C2_()
				_ = _nw152
				_nw152.Ctor__((_this).F17(), _this.F16(), _this.F9(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _dafny.UnicodeSeqOfUtf8Bytes("nbgkv")))
				_1019_v19 = _nw152
				var _1020_v20 D6
				_ = _1020_v20
				_1020_v20 = Companion_D6_.Create_DC18_(_1019_v19)
				var _1021_v21 D6
				_ = _1021_v21
				_1021_v21 = Companion_D6_.Create_DC20_(_1020_v20)
				_1021_v21 = Companion_D6_.Create_DC20_(_1020_v20)
				var _1022_v22 _dafny.Map
				_ = _1022_v22
				_1022_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_999_v0, _999_v0)
				var _1023_v23 _dafny.Sequence
				_ = _1023_v23
				_1023_v23 = _dafny.SeqOf(_999_v0, (func() bool {
					if (_1022_v22).Contains(false) {
						return (_1022_v22).Get(false).(bool)
					}
					return true
				})())
				var _1024_v24 D6
				_ = _1024_v24
				_1024_v24 = Companion_D6_.Create_DC19_(_dafny.IntOfUint32((_1023_v23).Cardinality()), _dafny.IntOfUint32((_1001_v2).Cardinality()), (_1019_v19).F6())
				var _1025_v25 _dafny.MultiSet
				_ = _1025_v25
				_1025_v25 = _dafny.MultiSetOf(_1024_v24, _1024_v24, _1024_v24, _1024_v24)
				(globalState).F5 = (func() _dafny.Int {
					if (_1025_v25).Contains(_1024_v24) {
						return (_1025_v25).Multiplicity(_1024_v24)
					}
					return (_this).F6()
				})()
				var _1026_v26 _dafny.Array
				_ = _1026_v26
				var _nw153 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
				_ = _nw153
				_1026_v26 = _nw153
				var _1027_v27 _dafny.Map
				_ = _1027_v27
				_1027_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_999_v0) || (_999_v0), _1026_v26)
				var _1028_v28 _dafny.Sequence
				_ = _1028_v28
				_1028_v28 = _dafny.SeqOf(_1026_v26, _1026_v26)
				_1027_v27 = (_1027_v27).Update(_999_v0, (func() _dafny.Array {
					if (_1027_v27).Contains(_999_v0) {
						return (_1027_v27).Get(_999_v0).(_dafny.Array)
					}
					return (_1028_v28).Select((Companion_Default___.SafeIndex((_this).F17(), _dafny.IntOfUint32((_1028_v28).Cardinality()))).Uint32()).(_dafny.Array)
				})())
				var _1029_v29 _dafny.Array
				_ = _1029_v29
				var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw154
				_1029_v29 = _nw154
				var _1030_v30 _dafny.Map
				_ = _1030_v30
				_1030_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_999_v0, p0)
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_1029_v29), 0))
				_ = _index151
				(_1029_v29).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(264), (_1030_v30).Cardinality()), (_index151).Int())
				var _1031_v31 _dafny.Map
				_ = _1031_v31
				_1031_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(914), _999_v0)
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_1029_v29), 0))
				_ = _index152
				(_1029_v29).ArraySet1(Companion_Default___.Fm2((func() _dafny.Int {
					if (_1030_v30).Contains(_999_v0) {
						return (_1030_v30).Get(_999_v0).(_dafny.Int)
					}
					return (_1031_v31).Cardinality()
				})(), globalState), (_index152).Int())
				var _1032_v32 _dafny.Sequence
				_ = _1032_v32
				_1032_v32 = _dafny.UnicodeSeqOfUtf8Bytes("lnnj")
				var _1033_v33 _dafny.Sequence
				_ = _1033_v33
				_1033_v33 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hieeu"), _1032_v32), _1032_v32, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(130))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg76 _dafny.Int) interface{} {
						return coer76(arg76)
					}
				}((func(_1034_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1035_i1 _dafny.Int) _dafny.CodePoint {
						return _1034_v1
					}
				})(_1000_v1))))
				var _rhs179 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(704))).Uint32(), func(coer77 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg77 _dafny.Int) interface{} {
						return coer77(arg77)
					}
				}(func(_1036_i2 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("xrfic")
				}))
				_ = _rhs179
				var _rhs180 _dafny.Int = (_this).F18()
				_ = _rhs180
				var _lhs145 *GlobalState = globalState
				_ = _lhs145
				_1033_v33 = _rhs179
				_lhs145.F5 = _rhs180
			}
			var _1037_v34 _dafny.Map
			_ = _1037_v34
			_1037_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm22(_dafny.MultiSetOf(_999_v0, !(_999_v0)), _999_v0, _dafny.IntOfInt64(66), globalState), (_this).F6())
			var _1038_v35 _dafny.Array
			_ = _1038_v35
			var _len0_24 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_24
			var _nw155 _dafny.Array
			_ = _nw155
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw155 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) _dafny.Int = func(_1039_i3 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1039_i3, _dafny.IntOfInt64(-228))
				}
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw155 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw155).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw155).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_1038_v35 = _nw155
			var _1040_v36 _dafny.MultiSet
			_ = _1040_v36
			_1040_v36 = _dafny.MultiSetOf((_this).F17())
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_1038_v35), 0))
			_ = _index153
			(_1038_v35).ArraySet1((func() _dafny.Int {
				if _999_v0 {
					return p0
				}
				return (_1040_v36).Cardinality()
			})(), (_index153).Int())
			var _1041_v37 _dafny.Sequence
			_ = _1041_v37
			_1041_v37 = _dafny.UnicodeSeqOfUtf8Bytes("w")
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_1038_v35), 0))
			_ = _index154
			var _rhs181 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1041_v37, _1041_v37)).Cardinality())
			_ = _rhs181
			var _rhs182 _dafny.Map = _1037_v34
			_ = _rhs182
			var _rhs183 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfInt64(-547))
			_ = _rhs183
			var _rhs184 _dafny.Int = (_this).F18()
			_ = _rhs184
			var _rhs185 _dafny.Array = (func() _dafny.Array {
				if _999_v0 {
					return _1038_v35
				}
				return (func() _dafny.Array {
					if _999_v0 {
						return _1038_v35
					}
					return _1038_v35
				})()
			})()
			_ = _rhs185
			var _lhs146 *GlobalState = globalState
			_ = _lhs146
			var _lhs147 _dafny.Array = _1038_v35
			_ = _lhs147
			var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_1038_v35), 0))
			_ = _lhs148
			var _lhs149 *GlobalState = globalState
			_ = _lhs149
			_lhs146.F5 = _rhs181
			_1037_v34 = _rhs182
			(_lhs147).ArraySet1(_rhs183, (_lhs148).Int())
			_lhs149.F5 = _rhs184
			_1038_v35 = _rhs185
			if _999_v0 {
				(globalState).F5 = Companion_Default___.SafeDivisionInt((_this).F6(), p0)
				var _1042_v38 *C2
				_ = _1042_v38
				var _nw156 *C2 = New_C2_()
				_ = _nw156
				_nw156.Ctor__((p0).Plus(_dafny.IntOfUint32((_1041_v37).Cardinality())), _this.F16(), _this.F9(), ((_this).F10()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _dafny.Companion_Sequence_.Update(_1041_v37, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(939))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg78 _dafny.Int) interface{} {
						return coer78(arg78)
					}
				}((func(_1043_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1044_i4 _dafny.Int) _dafny.CodePoint {
						return _1043_v1
					}
				})(_1000_v1)))).Cardinality()), _dafny.IntOfUint32((_1041_v37).Cardinality()))).Uint32(), _1000_v1))))
				_1042_v38 = _nw156
				var _1045_v39 *C1
				_ = _1045_v39
				var _nw157 *C1 = New_C1_()
				_ = _nw157
				_nw157.Ctor__(Companion_Default___.SafeDivisionInt((_this).F17(), (_this).F17()))
				_1045_v39 = _nw157
				(globalState).F0 = false
				_1037_v34 = (_1037_v34).Update(((_this).F18()).Cmp(_dafny.IntOfInt64(-227)) <= 0, _dafny.IntOfInt64(68))
			} else {
				var _1046_v40 _dafny.Map
				_ = _1046_v40
				_1046_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_999_v0, _999_v0)
				_1046_v40 = (_1046_v40).Update(_999_v0, _999_v0)
				var _1047_v41 _dafny.Array
				_ = _1047_v41
				var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
				_ = _nw158
				_1047_v41 = _nw158
				var _1048_v42 _dafny.Array
				_ = _1048_v42
				var _nwElement0_27 _dafny.Array = _1047_v41
				_ = _nwElement0_27
				var _nw159 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(6))
				_ = _nw159
				(_nw159).ArraySet1(_nwElement0_27, 0)
				(_nw159).ArraySet1(_1047_v41, 1)
				(_nw159).ArraySet1(_1047_v41, 2)
				(_nw159).ArraySet1(_1047_v41, 3)
				(_nw159).ArraySet1(_1047_v41, 4)
				(_nw159).ArraySet1(_1047_v41, 5)
				_1048_v42 = _nw159
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_1048_v42), 0))
				_ = _index155
				(_1048_v42).ArraySet1(_1047_v41, (_index155).Int())
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_1048_v42), 0))
				_ = _index156
				var _rhs186 _dafny.Array = _1047_v41
				_ = _rhs186
				var _rhs187 _dafny.Int = (_this).F17()
				_ = _rhs187
				var _rhs188 bool = !(_999_v0)
				_ = _rhs188
				var _lhs150 _dafny.Array = _1048_v42
				_ = _lhs150
				var _lhs151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_1048_v42), 0))
				_ = _lhs151
				var _lhs152 *GlobalState = globalState
				_ = _lhs152
				var _lhs153 *GlobalState = globalState
				_ = _lhs153
				(_lhs150).ArraySet1(_rhs186, (_lhs151).Int())
				_lhs152.F5 = _rhs187
				_lhs153.F0 = _rhs188
				var _1049_v43 _dafny.Sequence
				_ = _1049_v43
				_1049_v43 = _dafny.SeqOf(_999_v0)
				r1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_999_v0), _1049_v43), _1049_v43)
				(globalState).F5 = (_this).F17()
				var _arr32 _dafny.Array = _this.F9()
				_ = _arr32
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index157
				(_arr32).ArraySet1(_999_v0, (_index157).Int())
				var _arr33 _dafny.Array = _this.F9()
				_ = _arr33
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index158
				(_arr33).ArraySet1((!(_999_v0) || (_999_v0)) == (_999_v0), (_index158).Int())
			}
			var _1050_v44 _dafny.Array
			_ = _1050_v44
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_25
			var _nw160 _dafny.Array
			_ = _nw160
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw160 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) D5 = (func(_1051_v0 bool, _1052_v5 _dafny.Set) func(_dafny.Int) D5 {
					return func(_1053_i5 _dafny.Int) D5 {
						return (func() D5 {
							if _1051_v0 {
								return Companion_D5_.Create_DC16_(_dafny.SetOf(_1051_v0))
							}
							return Companion_D5_.Create_DC16_(_1052_v5)
						})()
					}
				})(_999_v0, _1004_v5)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw160 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw160).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw160).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_1050_v44 = _nw160
			var _1054_v45 D5
			_ = _1054_v45
			_1054_v45 = Companion_D5_.Create_DC16_(_1004_v5)
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_1050_v44), 0))
			_ = _index159
			(_1050_v44).ArraySet1(_1054_v45, (_index159).Int())
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_1050_v44), 0))
			_ = _index160
			(_1050_v44).ArraySet1(_1054_v45, (_index160).Int())
		} else {
			var _1055_v46 bool
			_ = _1055_v46
			_1055_v46 = false
			var _1056_v47 _dafny.Sequence
			_ = _1056_v47
			_1056_v47 = _dafny.SeqOf(p0)
			var _1057_v48 _dafny.Map
			_ = _1057_v48
			_1057_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1055_v46, _1056_v47)
			var _1058_v49 _dafny.CodePoint
			_ = _1058_v49
			_1058_v49 = _dafny.CodePoint('j')
			var _1059_v50 _dafny.Map
			_ = _1059_v50
			_1059_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _1055_v46)
			var _1060_v51 _dafny.MultiSet
			_ = _1060_v51
			_1060_v51 = _dafny.MultiSetOf(_1059_v50)
			_1057_v48 = (_1057_v48).Update((p0).Cmp(p0) > 0, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm26((_this).F6(), _1058_v49, !(_1055_v46), _1058_v49, globalState), _dafny.SeqOf(Companion_Default___.Fm2((_1060_v51).Cardinality(), globalState), _dafny.IntOfInt64(904))))
			var _1061_v52 D6
			_ = _1061_v52
			_1061_v52 = Companion_D6_.Create_DC19_(_dafny.IntOfInt64(313), p0, (_this).F17())
			var _1062_v53 D6
			_ = _1062_v53
			_1062_v53 = Companion_D6_.Create_DC20_(_1061_v52)
			var _source13 D6 = _1062_v53
			_ = _source13
			if _source13.Is_DC19() {
				var _1063___mcc_h0 _dafny.Int = _source13.Get_().(D6_DC19).Cf39
				_ = _1063___mcc_h0
				var _1064___mcc_h1 _dafny.Int = _source13.Get_().(D6_DC19).Cf40
				_ = _1064___mcc_h1
				var _1065___mcc_h2 _dafny.Int = _source13.Get_().(D6_DC19).Cf41
				_ = _1065___mcc_h2
				var _1066_cf41 _dafny.Int = _1065___mcc_h2
				_ = _1066_cf41
				var _1067_cf40 _dafny.Int = _1064___mcc_h1
				_ = _1067_cf40
				var _1068_cf39 _dafny.Int = _1063___mcc_h0
				_ = _1068_cf39
				var _arr34 _dafny.Array = _this.F9()
				_ = _arr34
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index161
				(_arr34).ArraySet1(_1055_v46, (_index161).Int())
				var _arr35 _dafny.Array = _this.F9()
				_ = _arr35
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index162
				(_arr35).ArraySet1(_1055_v46, (_index162).Int())
				var _1069_v54 _dafny.Array
				_ = _1069_v54
				var _nw161 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw161
				_1069_v54 = _nw161
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_1069_v54), 0))
				_ = _index163
				(_1069_v54).ArraySet1((_this).F6(), (_index163).Int())
				var _1070_v55 _dafny.Sequence
				_ = _1070_v55
				_1070_v55 = _dafny.SeqOf(_1069_v54, _1069_v54, _1069_v54, _1069_v54, _1069_v54)
				var _1071_v56 _dafny.Map
				_ = _1071_v56
				_1071_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1058_v49, (_this).F10())
				var _1072_v57 D3
				_ = _1072_v57
				_1072_v57 = Companion_D3_.Create_DC9_(_1056_v47)
				var _1073_v58 _dafny.MultiSet
				_ = _1073_v58
				_1073_v58 = _dafny.MultiSetOf((_1072_v57).Dtor_cf21())
				var _1074_v59 D0
				_ = _1074_v59
				_1074_v59 = Companion_D0_.Create_DC0_(Companion_Default___.Fm4(p0, _1067_cf40, (_1073_v58).Cardinality(), globalState), _dafny.IntOfInt64(644))
				var _arr36 _dafny.Array = _this.F9()
				_ = _arr36
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index164
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_1069_v54), 0))
				_ = _index165
				var _rhs189 bool = ((func() _dafny.Map {
					if (_1071_v56).Contains(_1058_v49) {
						return (_1071_v56).Get(_1058_v49).(_dafny.Map)
					}
					return (_this).F10()
				})()).Contains((_dafny.Zero).Minus(p0))
				_ = _rhs189
				var _rhs190 _dafny.Int = (_1074_v59).Dtor_cf1()
				_ = _rhs190
				var _rhs191 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_1068_cf39), _dafny.IntOfInt64(481))
				_ = _rhs191
				var _rhs192 _dafny.Sequence = _1070_v55
				_ = _rhs192
				var _lhs154 _dafny.Array = _this.F9()
				_ = _lhs154
				var _lhs155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_this.F9()), 0))
				_ = _lhs155
				var _lhs156 _dafny.Array = _1069_v54
				_ = _lhs156
				var _lhs157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(560), _dafny.ArrayLen((_1069_v54), 0))
				_ = _lhs157
				var _lhs158 *GlobalState = globalState
				_ = _lhs158
				(_lhs154).ArraySet1(_rhs189, (_lhs155).Int())
				(_lhs156).ArraySet1(_rhs190, (_lhs157).Int())
				_lhs158.F5 = _rhs191
				_1070_v55 = _rhs192
				var _1075_v60 _dafny.Array
				_ = _1075_v60
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_26
				var _nw162 _dafny.Array
				_ = _nw162
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw162 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.Map = (func(_1076_v46 bool) func(_dafny.Int) _dafny.Map {
						return func(_1077_i6 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1076_v46)
						}
					})(_1055_v46)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw162 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw162).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw162).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_1075_v60 = _nw162
				var _1078_v61 D9
				_ = _1078_v61
				_1078_v61 = Companion_D9_.Create_DC26_(_1075_v60)
				_1078_v61 = _1078_v61
				(globalState).F0 = (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)
			} else if _source13.Is_DC18() {
				var _1079___mcc_h3 T1 = _source13.Get_().(D6_DC18).Cf38
				_ = _1079___mcc_h3
				var _1080_cf38 T1 = _1079___mcc_h3
				_ = _1080_cf38
				var _1081_v62 _dafny.Sequence
				_ = _1081_v62
				_1081_v62 = _dafny.SeqOf(_1056_v47, _1056_v47, _1056_v47)
				_1081_v62 = _1081_v62
				(globalState).F5 = (_this).F17()
				var _1082_v63 _dafny.Map
				_ = _1082_v63
				_1082_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(59))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg79 _dafny.Int) interface{} {
						return coer79(arg79)
					}
				}(func(_1083_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				})))
				var _1084_v64 _dafny.Sequence
				_ = _1084_v64
				_1084_v64 = _dafny.UnicodeSeqOfUtf8Bytes("oaqqtxvm")
				var _1085_v65 _dafny.Map
				_ = _1085_v65
				_1085_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1055_v46, _1084_v64)
				_1082_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() _dafny.Sequence {
					if (_1085_v65).Contains(_1055_v46) {
						return (_1085_v65).Get(_1055_v46).(_dafny.Sequence)
					}
					return _1084_v64
				})())
				var _1086_v66 _dafny.Array
				_ = _1086_v66
				var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw163
				_1086_v66 = _nw163
				var _1087_v67 _dafny.Array
				_ = _1087_v67
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_27
				var _nw164 _dafny.Array
				_ = _nw164
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw164 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) _dafny.Int = func(_1088_i8 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1088_i8, (_this).F6())
					}
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw164 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw164).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw164).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_1087_v67 = _nw164
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_1086_v66), 0))
				_ = _index166
				(_1086_v66).ArraySet1((func() _dafny.Array {
					if _1055_v46 {
						return _1087_v67
					}
					return _1087_v67
				})(), (_index166).Int())
				var _1089_v68 _dafny.Sequence
				_ = _1089_v68
				_1089_v68 = _dafny.SeqOf(_1055_v46, _1055_v46, (p0).Cmp(p0) <= 0)
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_1086_v66), 0))
				_ = _index167
				var _rhs193 _dafny.Int = (_1056_v47).Select((Companion_Default___.SafeIndex((_this).F17(), _dafny.IntOfUint32((_1056_v47).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs193
				var _rhs194 _dafny.Int = (_this).F18()
				_ = _rhs194
				var _rhs195 bool = (_1089_v68).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1089_v68).Cardinality()))).Uint32()).(bool)
				_ = _rhs195
				var _rhs196 _dafny.Array = _1087_v67
				_ = _rhs196
				var _lhs159 *GlobalState = globalState
				_ = _lhs159
				var _lhs160 *GlobalState = globalState
				_ = _lhs160
				var _lhs161 _dafny.Array = _1086_v66
				_ = _lhs161
				var _lhs162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_1086_v66), 0))
				_ = _lhs162
				_lhs159.F5 = _rhs193
				_lhs160.F5 = _rhs194
				_1055_v46 = _rhs195
				(_lhs161).ArraySet1(_rhs196, (_lhs162).Int())
			} else {
				var _1090___mcc_h4 D6 = _source13.Get_().(D6_DC20).Cf42
				_ = _1090___mcc_h4
				var _1091_cf42 D6 = _1090___mcc_h4
				_ = _1091_cf42
				var _1092_v69 _dafny.Sequence
				_ = _1092_v69
				_1092_v69 = _dafny.SeqOf(_1055_v46)
				var _1093_v70 _dafny.Map
				_ = _1093_v70
				_1093_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1092_v69).Select((Companion_Default___.SafeIndex((_this).F17(), _dafny.IntOfUint32((_1092_v69).Cardinality()))).Uint32()).(bool), !(_1055_v46))
				var _1094_v71 _dafny.Sequence
				_ = _1094_v71
				_1094_v71 = _dafny.SeqOf(_1093_v70)
				_1094_v71 = (func() _dafny.Sequence {
					if _1055_v46 {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1093_v70, _1093_v70), Companion_Default___.Fm46(globalState))
					}
					return _dafny.Companion_Sequence_.Concatenate(_1094_v71, _1094_v71)
				})()
				(globalState).F0 = false
				var _1095_v72 _dafny.MultiSet
				_ = _1095_v72
				_1095_v72 = _dafny.MultiSetOf((_this).F17())
				var _1096_v74 _dafny.Array
				_ = _1096_v74
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_28
				var _nw165 _dafny.Array
				_ = _nw165
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw165 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) _dafny.Int = (func(_1097_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1098_i9 _dafny.Int) _dafny.Int {
							return (_1098_i9).Plus(_1097_p0)
						}
					})(p0)
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw165 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw165).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw165).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_1096_v74 = _nw165
				var _1099_v75 _dafny.Map
				_ = _1099_v75
				_1099_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1055_v46, _1096_v74)
				var _pat_let_tv19 = _1095_v72
				_ = _pat_let_tv19
				var _pat_let_tv20 = _1095_v72
				_ = _pat_let_tv20
				var _1100_v76 _dafny.Map
				_ = _1100_v76
				_1100_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let23_0 D11) D11 {
					return func(_1101_dt__update__tmp_h0 D11) D11 {
						return func(_pat_let24_0 _dafny.Set) D11 {
							return func(_1103_dt__update_hcf57_h0 _dafny.Set) D11 {
								return Companion_D11_.Create_DC30_(_1103_dt__update_hcf57_h0)
							}(_pat_let24_0)
						}(func() _dafny.Set {
							var _coll32 = _dafny.NewBuilder()
							_ = _coll32
							for _iter34 := _dafny.Iterate((_pat_let_tv19).Elements()); ; {
								_compr_32, _ok34 := _iter34()
								if !_ok34 {
									break
								}
								var _1102_v73 _dafny.Int
								_1102_v73 = interface{}(_compr_32).(_dafny.Int)
								if (_pat_let_tv20).Contains(_1102_v73) {
									_coll32.Add(Companion_Default___.SafeDivisionInt(_1102_v73, (_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality()))
								}
							}
							return _coll32.ToSet()
						}())
					}(_pat_let23_0)
				}(Companion_D11_.Create_DC30_(_dafny.SetOf((_this).F17()))), _1099_v75)
				var _1104_v77 _dafny.Map
				_ = _1104_v77
				_1104_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _dafny.IntOfInt64(931))
				var _1105_v78 *C3
				_ = _1105_v78
				var _nw166 *C3 = New_C3_()
				_ = _nw166
				_nw166.Ctor__(_1104_v77, _this.F9(), (_this).F10(), p0)
				_1105_v78 = _nw166
				var _1106_v79 _dafny.Sequence
				_ = _1106_v79
				_1106_v79 = _dafny.SeqOf(_1105_v78, _1105_v78)
				var _1107_v80 _dafny.Map
				_ = _1107_v80
				_1107_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1106_v79).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1106_v79).Cardinality()))).Uint32()).(*C3), false)
				var _1108_v81 _dafny.Set
				_ = _1108_v81
				_1108_v81 = _dafny.SetOf((_1107_v80).Cardinality())
				var _1109_v82 D11
				_ = _1109_v82
				_1109_v82 = Companion_D11_.Create_DC30_(_1108_v81)
				_1100_v76 = (_1100_v76).Update(_1109_v82, _1099_v75)
				var _rhs197 _dafny.Map = (_1093_v70).Update(_1055_v46, (p0).Cmp(p0) <= 0)
				_ = _rhs197
				var _rhs198 bool = _1055_v46
				_ = _rhs198
				var _rhs199 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_this).F18(), (_this).F6()), (_this).F6())
				_ = _rhs199
				var _lhs163 *GlobalState = globalState
				_ = _lhs163
				var _lhs164 *GlobalState = globalState
				_ = _lhs164
				_1093_v70 = _rhs197
				_lhs163.F0 = _rhs198
				_lhs164.F5 = _rhs199
			}
			var _arr37 _dafny.Array = _this.F9()
			_ = _arr37
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index168
			(_arr37).ArraySet1(_1055_v46, (_index168).Int())
			var _1110_v83 _dafny.Set
			_ = _1110_v83
			_1110_v83 = _dafny.SetOf(_1055_v46)
			var _arr38 _dafny.Array = _this.F9()
			_ = _arr38
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index169
			(_arr38).ArraySet1((_1110_v83).IsProperSubsetOf(_1110_v83), (_index169).Int())
			var _arr39 _dafny.Array = _this.F9()
			_ = _arr39
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index170
			(_arr39).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), (_index170).Int())
			var _arr40 _dafny.Array = _this.F9()
			_ = _arr40
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index171
			(_arr40).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), (_index171).Int())
		}
		var _1111_v84 _dafny.Array
		_ = _1111_v84
		var _len0_29 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_29
		var _nw167 _dafny.Array
		_ = _nw167
		if _len0_29.Cmp(_dafny.Zero) == 0 {
			_nw167 = _dafny.NewArray(_len0_29)
		} else {
			var _init29 func(_dafny.Int) _dafny.Sequence = func(_1112_i10 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("rkntbsman")
			}
			_ = _init29
			var _element0_29 = _init29(_dafny.Zero)
			_ = _element0_29
			_nw167 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
			(_nw167).ArraySet1(_element0_29, 0)
			var _nativeLen0_29 = (_len0_29).Int()
			_ = _nativeLen0_29
			for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
				(_nw167).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
			}
		}
		_1111_v84 = _nw167
		var _1113_v85 bool
		_ = _1113_v85
		_1113_v85 = true
		var _1114_v86 D10
		_ = _1114_v86
		_1114_v86 = Companion_D10_.Create_DC29_(p0, _1111_v84, _1113_v85, (_this).F18())
		var _source14 D10 = _1114_v86
		_ = _source14
		if _source14.Is_DC29() {
			var _1115___mcc_h5 _dafny.Int = _source14.Get_().(D10_DC29).Cf53
			_ = _1115___mcc_h5
			var _1116___mcc_h6 _dafny.Array = _source14.Get_().(D10_DC29).Cf54
			_ = _1116___mcc_h6
			var _1117___mcc_h7 bool = _source14.Get_().(D10_DC29).Cf55
			_ = _1117___mcc_h7
			var _1118___mcc_h8 _dafny.Int = _source14.Get_().(D10_DC29).Cf56
			_ = _1118___mcc_h8
			var _1119_cf56 _dafny.Int = _1118___mcc_h8
			_ = _1119_cf56
			var _1120_cf55 bool = _1117___mcc_h7
			_ = _1120_cf55
			var _1121_cf54 _dafny.Array = _1116___mcc_h6
			_ = _1121_cf54
			var _1122_cf53 _dafny.Int = _1115___mcc_h5
			_ = _1122_cf53
			(globalState).F2 = _dafny.UnicodeSeqOfUtf8Bytes("na")
			var _arr41 _dafny.Array = _this.F9()
			_ = _arr41
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index172
			(_arr41).ArraySet1(_1113_v85, (_index172).Int())
			var _1123_v87 _dafny.CodePoint
			_ = _1123_v87
			_1123_v87 = _dafny.CodePoint('d')
			var _1124_v88 _dafny.Set
			_ = _1124_v88
			_1124_v88 = _dafny.SetOf(_dafny.IntOfInt64(754), (_this).F18())
			var _1125_v89 _dafny.Array
			_ = _1125_v89
			var _nwElement0_28 _dafny.CodePoint = _1123_v87
			_ = _nwElement0_28
			var _nw168 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(21))
			_ = _nw168
			(_nw168).ArraySet1CodePoint(_nwElement0_28, 0)
			(_nw168).ArraySet1CodePoint(_1123_v87, 1)
			(_nw168).ArraySet1CodePoint(_1123_v87, 2)
			(_nw168).ArraySet1CodePoint(_1123_v87, 3)
			(_nw168).ArraySet1CodePoint(_1123_v87, 4)
			(_nw168).ArraySet1CodePoint(Companion_Default___.Fm0((_this).F18(), _1124_v88, _1113_v85, globalState), 5)
			(_nw168).ArraySet1CodePoint(_1123_v87, 6)
			(_nw168).ArraySet1CodePoint(_1123_v87, 7)
			(_nw168).ArraySet1CodePoint(_1123_v87, 8)
			(_nw168).ArraySet1CodePoint(_1123_v87, 9)
			(_nw168).ArraySet1CodePoint(_1123_v87, 10)
			(_nw168).ArraySet1CodePoint(_1123_v87, 11)
			(_nw168).ArraySet1CodePoint(_1123_v87, 12)
			(_nw168).ArraySet1CodePoint(_1123_v87, 13)
			(_nw168).ArraySet1CodePoint(_1123_v87, 14)
			(_nw168).ArraySet1CodePoint(_1123_v87, 15)
			(_nw168).ArraySet1CodePoint(_1123_v87, 16)
			(_nw168).ArraySet1CodePoint(_1123_v87, 17)
			(_nw168).ArraySet1CodePoint(_1123_v87, 18)
			(_nw168).ArraySet1CodePoint(_1123_v87, 19)
			(_nw168).ArraySet1CodePoint(_1123_v87, 20)
			_1125_v89 = _nw168
			var _1126_v90 _dafny.Sequence
			_ = _1126_v90
			_1126_v90 = _dafny.SeqOf(_1125_v89, _1125_v89, _1125_v89)
			var _arr42 _dafny.Array = _this.F9()
			_ = _arr42
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index173
			(_arr42).ArraySet1(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1125_v89, _1123_v87)).Contains((_1126_v90).Select((Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_1126_v90).Cardinality()))).Uint32()).(_dafny.Array)), (_index173).Int())
			var _1127_v91 _dafny.Sequence
			_ = _1127_v91
			_1127_v91 = _dafny.SeqOf((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
			var _1128_v92 _dafny.Sequence
			_ = _1128_v92
			_1128_v92 = _dafny.SeqOf(_1127_v91, _1127_v91)
			var _1129_v93 _dafny.Map
			_ = _1129_v93
			_1129_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1128_v92, _1128_v92)).Cardinality())), (_this).F17())
			_1129_v93 = (_1129_v93).Update((_1124_v88).Cardinality(), p0)
			var _1130_v94 D0
			_ = _1130_v94
			_1130_v94 = Companion_D0_.Create_DC0_(_1127_v91, _1122_cf53)
			var _1131_v95 D1
			_ = _1131_v95
			_1131_v95 = Companion_D1_.Create_DC4_(_1130_v94, Companion_Default___.SafeModuloInt((_this).F6(), (_this).F6()), _dafny.UnicodeSeqOfUtf8Bytes("pceere"), _1120_cf55, p0)
			var _source15 D1 = _1131_v95
			_ = _source15
			if _source15.Is_DC4() {
				var _1132___mcc_h10 D0 = _source15.Get_().(D1_DC4).Cf7
				_ = _1132___mcc_h10
				var _1133___mcc_h11 _dafny.Int = _source15.Get_().(D1_DC4).Cf8
				_ = _1133___mcc_h11
				var _1134___mcc_h12 _dafny.Sequence = _source15.Get_().(D1_DC4).Cf9
				_ = _1134___mcc_h12
				var _1135___mcc_h13 bool = _source15.Get_().(D1_DC4).Cf10
				_ = _1135___mcc_h13
				var _1136___mcc_h14 _dafny.Int = _source15.Get_().(D1_DC4).Cf11
				_ = _1136___mcc_h14
				var _1137_cf11 _dafny.Int = _1136___mcc_h14
				_ = _1137_cf11
				var _1138_cf10 bool = _1135___mcc_h13
				_ = _1138_cf10
				var _1139_cf9 _dafny.Sequence = _1134___mcc_h12
				_ = _1139_cf9
				var _1140_cf8 _dafny.Int = _1133___mcc_h11
				_ = _1140_cf8
				var _1141_cf7 D0 = _1132___mcc_h10
				_ = _1141_cf7
				var _1142_v96 _dafny.Map
				_ = _1142_v96
				_1142_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1111_v84)
				var _1143_v97 _dafny.Array
				_ = _1143_v97
				var _nwElement0_29 _dafny.Array = _1121_cf54
				_ = _nwElement0_29
				var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(21))
				_ = _nw169
				(_nw169).ArraySet1(_nwElement0_29, 0)
				(_nw169).ArraySet1(_1111_v84, 1)
				(_nw169).ArraySet1((func() _dafny.Array {
					if (_1142_v96).Contains((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)) {
						return (_1142_v96).Get((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)).(_dafny.Array)
					}
					return _1111_v84
				})(), 2)
				(_nw169).ArraySet1(_1111_v84, 3)
				(_nw169).ArraySet1(_1121_cf54, 4)
				(_nw169).ArraySet1(_1111_v84, 5)
				(_nw169).ArraySet1(_1111_v84, 6)
				(_nw169).ArraySet1(_1111_v84, 7)
				(_nw169).ArraySet1(_1121_cf54, 8)
				(_nw169).ArraySet1(_1121_cf54, 9)
				(_nw169).ArraySet1(_1121_cf54, 10)
				(_nw169).ArraySet1(_1121_cf54, 11)
				(_nw169).ArraySet1((Companion_D10_.Create_DC29_((_this).F6(), _1111_v84, _1113_v85, Companion_Default___.Fm2(_dafny.IntOfInt64(534), globalState))).Dtor_cf54(), 12)
				(_nw169).ArraySet1(_1111_v84, 13)
				(_nw169).ArraySet1(_1111_v84, 14)
				(_nw169).ArraySet1(_1121_cf54, 15)
				(_nw169).ArraySet1(_1111_v84, 16)
				(_nw169).ArraySet1(_1111_v84, 17)
				(_nw169).ArraySet1(_1111_v84, 18)
				(_nw169).ArraySet1(_1111_v84, 19)
				(_nw169).ArraySet1(_1111_v84, 20)
				_1143_v97 = _nw169
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_1143_v97), 0))
				_ = _index174
				(_1143_v97).ArraySet1(_1121_cf54, (_index174).Int())
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_1143_v97), 0))
				_ = _index175
				(_1143_v97).ArraySet1(_1121_cf54, (_index175).Int())
				var _arr43 _dafny.Array = _this.F9()
				_ = _arr43
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index176
				(_arr43).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_1127_v91, _dafny.SeqOf(_1138_cf10, _1120_cf55)), _1128_v92), (_index176).Int())
				_1122_cf53 = Companion_Default___.Fm2(_dafny.IntOfInt64(615), globalState)
				var _1144_v98 _dafny.Array
				_ = _1144_v98
				var _nw170 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(20))
				_ = _nw170
				_1144_v98 = _nw170
				var _1145_v99 _dafny.Array
				_ = _1145_v99
				var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
				_ = _nw171
				_1145_v99 = _nw171
				var _rhs200 _dafny.Array = _1144_v98
				_ = _rhs200
				var _rhs201 _dafny.Array = _1145_v99
				_ = _rhs201
				var _rhs202 _dafny.Sequence = Companion_Default___.Fm25(_1137_cf11, !_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-665))).Uint32(), func(coer80 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg80 _dafny.Int) interface{} {
						return coer80(arg80)
					}
				}((func(_1146_cf11 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1147_i11 _dafny.Int) _dafny.Int {
						return _1146_cf11
					}
				})(_1137_cf11))), (_this).F6()), globalState)
				_ = _rhs202
				_1144_v98 = _rhs200
				_1145_v99 = _rhs201
				r2 = _rhs202
			} else {
				var _1148___mcc_h15 _dafny.MultiSet = _source15.Get_().(D1_DC3).Cf6
				_ = _1148___mcc_h15
				var _1149_cf6 _dafny.MultiSet = _1148___mcc_h15
				_ = _1149_cf6
				var _1150_v100 _dafny.Array
				_ = _1150_v100
				var _nw172 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw172
				_1150_v100 = _nw172
				var _1151_v101 *C0
				_ = _1151_v101
				var _nw173 *C0 = New_C0_()
				_ = _nw173
				_nw173.Ctor__(_1150_v100, !(!((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))))
				_1151_v101 = _nw173
				var _1152_v102 *C1
				_ = _1152_v102
				var _nw174 *C1 = New_C1_()
				_ = _nw174
				_nw174.Ctor__((_this).F6())
				_1152_v102 = _nw174
				var _1153_v103 _dafny.Map
				_ = _1153_v103
				_1153_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(187), _1152_v102)
				var _1154_v104 _dafny.MultiSet
				_ = _1154_v104
				_1154_v104 = _dafny.MultiSetOf(_1119_cf56, (_1153_v103).Cardinality())
				_1154_v104 = (_1154_v104).Update((_this).F18(), Companion_Default___.Abs((_this).F18()))
				var _1155_v105 _dafny.Sequence
				_ = _1155_v105
				_1155_v105 = _dafny.UnicodeSeqOfUtf8Bytes("uyva")
				var _1156_v106 _dafny.MultiSet
				_ = _1156_v106
				_1156_v106 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("jxfruork"), _1155_v105, _1155_v105)
				var _1157_v108 D0
				_ = _1157_v108
				_1157_v108 = Companion_D0_.Create_DC1_(_1151_v101.F20, _1119_cf56, _1123_v87)
				var _1158_v109 _dafny.Set
				_ = _1158_v109
				_1158_v109 = _dafny.SetOf(_1113_v85, _1120_cf55)
				_1156_v106 = ((Companion_Default___.Fm47(_1123_v87, func() _dafny.Map {
					var _coll33 = _dafny.NewMapBuilder()
					_ = _coll33
					for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(355), _dafny.IntOfInt64(690))); ; {
						_compr_33, _ok35 := _iter35()
						if !_ok35 {
							break
						}
						var _1159_v107 _dafny.Int
						_1159_v107 = interface{}(_compr_33).(_dafny.Int)
						if ((_dafny.IntOfInt64(355)).Cmp(_1159_v107) <= 0) && ((_1159_v107).Cmp(_dafny.IntOfInt64(690)) < 0) {
							_coll33.Add(Companion_Default___.SafeModuloInt(_1159_v107, (_this).F6()), _1113_v85)
						}
					}
					return _coll33.ToMap()
				}(), _1120_cf55, globalState)).Difference(_dafny.MultiSetOf(_1155_v105))).Difference((_1156_v106).Update(Companion_Default___.Fm30(_1113_v85, (_1157_v108).Dtor_cf4(), globalState), Companion_Default___.Abs((_1158_v109).Cardinality())))
				_1123_v87 = _1123_v87
			}
		} else {
			var _1160___mcc_h9 _dafny.Map = _source14.Get_().(D10_DC28).Cf52
			_ = _1160___mcc_h9
			var _1161_cf52 _dafny.Map = _1160___mcc_h9
			_ = _1161_cf52
			var _1162_v110 *C0
			_ = _1162_v110
			var _nw175 *C0 = New_C0_()
			_ = _nw175
			_nw175.Ctor__(_this.F9(), _1113_v85)
			_1162_v110 = _nw175
			(globalState).F2 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-170))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg81 _dafny.Int) interface{} {
					return coer81(arg81)
				}
			}(func(_1163_i12 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			}))
			var _1164_v111 _dafny.Sequence
			_ = _1164_v111
			_1164_v111 = _dafny.SeqOf(_1162_v110.F20)
			var _1165_v112 _dafny.MultiSet
			_ = _1165_v112
			_1165_v112 = _dafny.MultiSetOf((_this).F18(), Companion_Default___.Fm2(_dafny.IntOfUint32((_1164_v111).Cardinality()), globalState), (_this).F17(), (_this).F18())
			var _1166_v113 _dafny.Map
			_ = _1166_v113
			_1166_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1113_v85, (_1165_v112).Intersection(_1165_v112))
			var _1167_v114 _dafny.Sequence
			_ = _1167_v114
			_1167_v114 = _dafny.SeqOf((_this).F6(), p0, (_this).F17(), (_this).F17())
			_1166_v113 = (_1166_v113).Update(_1162_v110.F20, _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_1167_v114, (Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_1167_v114).Cardinality()))).Uint32(), (_this).F17())))
			var _arr44 _dafny.Array = _1162_v110.F19
			_ = _arr44
			var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1162_v110.F19), 0))
			_ = _index177
			(_arr44).ArraySet1((func() bool {
				if _1162_v110.F20 {
					return _1113_v85
				}
				return _1113_v85
			})(), (_index177).Int())
			var _1168_v115 _dafny.Array
			_ = _1168_v115
			var _len0_30 _dafny.Int = _dafny.One
			_ = _len0_30
			var _nw176 _dafny.Array
			_ = _nw176
			if _len0_30.Cmp(_dafny.Zero) == 0 {
				_nw176 = _dafny.NewArray(_len0_30)
			} else {
				var _init30 func(_dafny.Int) _dafny.Int = func(_1169_i13 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_1169_i13, (_this).F17())
				}
				_ = _init30
				var _element0_30 = _init30(_dafny.Zero)
				_ = _element0_30
				_nw176 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
				(_nw176).ArraySet1(_element0_30, 0)
				var _nativeLen0_30 = (_len0_30).Int()
				_ = _nativeLen0_30
				for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
					(_nw176).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
				}
			}
			_1168_v115 = _nw176
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_1168_v115), 0))
			_ = _index178
			(_1168_v115).ArraySet1(_dafny.IntOfInt64(134), (_index178).Int())
			var _1170_v116 _dafny.Sequence
			_ = _1170_v116
			_1170_v116 = _dafny.UnicodeSeqOfUtf8Bytes("tl")
			var _1171_v117 _dafny.Sequence
			_ = _1171_v117
			_1171_v117 = _dafny.SeqOf(_1170_v116)
			var _arr45 _dafny.Array = _1162_v110.F19
			_ = _arr45
			var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1162_v110.F19), 0))
			_ = _index179
			var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_1168_v115), 0))
			_ = _index180
			var _rhs203 bool = !(_1113_v85)
			_ = _rhs203
			var _rhs204 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfInt64(367))
			_ = _rhs204
			var _rhs205 _dafny.Int = p0
			_ = _rhs205
			var _rhs206 bool = (Companion_D4_.Create_DC13_(_1167_v114, false, p0)).Dtor_cf26()
			_ = _rhs206
			var _rhs207 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F17(), _dafny.IntOfUint32(((_1171_v117).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.IntOfUint32((_1171_v117).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), (_this).F17(), (_this).F17()), _1167_v114)).Cardinality())
			_ = _rhs207
			var _lhs165 _dafny.Array = _1162_v110.F19
			_ = _lhs165
			var _lhs166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1162_v110.F19), 0))
			_ = _lhs166
			var _lhs167 *GlobalState = globalState
			_ = _lhs167
			var _lhs168 _dafny.Array = _1168_v115
			_ = _lhs168
			var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_1168_v115), 0))
			_ = _lhs169
			var _lhs170 *GlobalState = globalState
			_ = _lhs170
			var _lhs171 *GlobalState = globalState
			_ = _lhs171
			(_lhs165).ArraySet1(_rhs203, (_lhs166).Int())
			_lhs167.F5 = _rhs204
			(_lhs168).ArraySet1(_rhs205, (_lhs169).Int())
			_lhs170.F0 = _rhs206
			_lhs171.F5 = _rhs207
		}
		var _1172_v118 _dafny.CodePoint
		_ = _1172_v118
		_1172_v118 = _dafny.CodePoint('i')
		var _1173_v119 _dafny.Sequence
		_ = _1173_v119
		_1173_v119 = _dafny.SeqOf(_1172_v118)
		var _1174_i14 _dafny.Int
		_ = _1174_i14
		_1174_i14 = _dafny.Zero
		{
			for !(((_this).F6()).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(981), Companion_Default___.Fm2(_dafny.IntOfUint32((_1173_v119).Cardinality()), globalState))) < 0) {
				{
					if (_1174_i14).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1174_i14 = (_1174_i14).Plus(_dafny.One)
					(globalState).F5 = Companion_Default___.SafeDivisionInt((_this).F17(), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1173_v119).Cardinality()), (_this).F18()))
					var _1175_v121 _dafny.Sequence
					_ = _1175_v121
					_1175_v121 = _dafny.SeqOf(p0)
					var _1176_v122 _dafny.Set
					_ = _1176_v122
					_1176_v122 = _dafny.SetOf(_dafny.IntOfUint32((_1175_v121).Cardinality()), (_this).F6())
					var _1177_v123 *C2
					_ = _1177_v123
					var _nw177 *C2 = New_C2_()
					_ = _nw177
					_nw177.Ctor__(p0, _this.F16(), _this.F9(), func() _dafny.Map {
						var _coll34 = _dafny.NewMapBuilder()
						_ = _coll34
						for _iter36 := _dafny.Iterate((_1176_v122).Elements()); ; {
							_compr_34, _ok36 := _iter36()
							if !_ok36 {
								break
							}
							var _1178_v120 _dafny.Int
							_1178_v120 = interface{}(_compr_34).(_dafny.Int)
							if (_1176_v122).Contains(_1178_v120) {
								_coll34.Add(Companion_Default___.SafeModuloInt(_1178_v120, _dafny.IntOfInt64(138)), _1173_v119)
							}
						}
						return _coll34.ToMap()
					}())
					_1177_v123 = _nw177
					_1176_v122 = (func() _dafny.Set {
						if !(_1113_v85) {
							return (_1176_v122).Difference(_1176_v122)
						}
						return (_1176_v122).Union(_1176_v122)
					})()
					var _1179_v124 _dafny.Map
					_ = _1179_v124
					_1179_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1113_v85, _1113_v85)
					var _1180_v125 _dafny.MultiSet
					_ = _1180_v125
					_1180_v125 = _dafny.MultiSetOf(_1113_v85)
					var _1181_v126 _dafny.Map
					_ = _1181_v126
					_1181_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_this).F18())
					var _rhs208 bool = ((p0).Minus(p0)).Cmp((_this).F6()) > 0
					_ = _rhs208
					var _rhs209 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1113_v85, (_1177_v123).Fm22(_1180_v125, _1113_v85, (_1181_v126).Cardinality(), globalState))
					_ = _rhs209
					var _lhs172 *GlobalState = globalState
					_ = _lhs172
					_lhs172.F0 = _rhs208
					_1179_v124 = _rhs209
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1182_v127 _dafny.MultiSet
		_ = _1182_v127
		_1182_v127 = _dafny.MultiSetOf(_1172_v118)
		var _1183_v128 _dafny.Map
		_ = _1183_v128
		_1183_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1182_v127)
		var _1184_v129 _dafny.Array
		_ = _1184_v129
		var _nw178 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
		_ = _nw178
		_1184_v129 = _nw178
		var _1185_v130 _dafny.MultiSet
		_ = _1185_v130
		_1185_v130 = _dafny.MultiSetOf(_1184_v129, _1184_v129, _1184_v129, _1184_v129)
		var _1186_v131 _dafny.Sequence
		_ = _1186_v131
		_1186_v131 = _dafny.SeqOf((_dafny.MultiSetOf(p0, (_1185_v130).Cardinality())).Cardinality())
		_1183_v128 = (_1183_v128).Update(_dafny.IntOfUint32((_1186_v131).Cardinality()), _1182_v127)
		var _1187_v132 _dafny.Array
		_ = _1187_v132
		var _nw179 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
		_ = _nw179
		_1187_v132 = _nw179
		var _1188_v133 _dafny.Sequence
		_ = _1188_v133
		_1188_v133 = _dafny.SeqOf(_1113_v85)
		var _1189_v134 _dafny.Sequence
		_ = _1189_v134
		_1189_v134 = _dafny.SeqOf(_1188_v133, _1188_v133, _dafny.SeqOf(_1113_v85))
		var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1187_v132), 0))
		_ = _index181
		(_1187_v132).ArraySet1(_1189_v134, (_index181).Int())
		var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1187_v132), 0))
		_ = _index182
		(_1187_v132).ArraySet1(_1189_v134, (_index182).Int())
		var _1190_v135 D0
		_ = _1190_v135
		_1190_v135 = Companion_D0_.Create_DC1_(_1113_v85, p0, (_1173_v119).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1173_v119).Cardinality()))).Uint32()).(_dafny.CodePoint))
		var _1191_v136 _dafny.Map
		_ = _1191_v136
		_1191_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9(), _1113_v85)
		var _1192_v137 _dafny.Map
		_ = _1192_v137
		_1192_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1113_v85, (func() _dafny.Map {
			if (_1190_v135).Dtor_cf2() {
				return _1191_v136
			}
			return _1191_v136
		})())
		_1192_v137 = (_1192_v137).Update(!(_1113_v85) || (_1113_v85), _1191_v136)
		r0 = Companion_D2_.Create_DC8_(true, !(((_this).F6()).Cmp(Companion_Default___.Fm2(_dafny.IntOfUint32((_1173_v119).Cardinality()), globalState)) < 0))
		r1 = _dafny.SeqOf(!(Companion_Default___.Fm3(globalState)) || (false))
		r2 = _dafny.Companion_Sequence_.Concatenate(_1173_v119, _1173_v119)
		return r0, r1, r2
	}
}
func (_this *C4) F17() _dafny.Int {
	{
		return _this._f17
	}
}
func (_this *C4) F18() _dafny.Int {
	{
		return _this._f18
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f6  _dafny.Int
	_f15 _dafny.Int
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f6 = _dafny.Zero
	_this._f15 = _dafny.Zero
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

func (_this *C5) F6() _dafny.Int {
	return _this._f6
}
func (_this *C5) Ctor__(f15 _dafny.Int, f6 _dafny.Int) {
	{
		(_this)._f15 = f15
		(_this)._f6 = f6
	}
}
func (_this *C5) Fm5(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('h')), _dafny.MultiSetOf(_dafny.CodePoint('j'), _dafny.CodePoint('g'))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('r'), _dafny.CodePoint('q'), _dafny.CodePoint('d')))), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('c')))))
	}
}
func (_this *C5) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _dafny.SeqOf((_this).F6()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC0_(_dafny.SeqOf(false, (Companion_D2_.Create_DC8_(!(false), false)).Dtor_cf20()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ig")).Cardinality()))).Dtor_cf1(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(174))).Uint32(), func(coer82 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg82 _dafny.Int) interface{} {
				return coer82(arg82)
			}
		}(func(_1193_i0 _dafny.Int) _dafny.Int {
			return (_this).F15()
		}))))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _dafny.SeqOf((Companion_D1_.Create_DC4_(Companion_D0_.Create_DC0_(_dafny.SeqOf(!(true), true), (_this).F6()), (_this).F6(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(108))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg83 _dafny.Int) interface{} {
				return coer83(arg83)
			}
		}(func(_1194_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		})), false, (_this).F15())).Dtor_cf11(), (_this).F6()))).Merge(func() _dafny.Map {
			var _coll35 = _dafny.NewMapBuilder()
			_ = _coll35
			for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(94), _dafny.IntOfInt64(657))); ; {
				_compr_35, _ok37 := _iter37()
				if !_ok37 {
					break
				}
				var _1195_v0 _dafny.Int
				_1195_v0 = interface{}(_compr_35).(_dafny.Int)
				if ((_dafny.IntOfInt64(94)).Cmp(_1195_v0) <= 0) && ((_1195_v0).Cmp(_dafny.IntOfInt64(657)) < 0) {
					_coll35.Add((_1195_v0).Plus((_this).F6()), _dafny.SeqOf((_dafny.MultiSetOf(true)).Cardinality()))
				}
			}
			return _coll35.ToMap()
		}()))
	}
}
func (_this *C5) M10(p0 bool, p1 _dafny.Int, globalState *GlobalState) (_dafny.Map, bool) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var _1196_v0 _dafny.Array
		_ = _1196_v0
		var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
		_ = _nw180
		_1196_v0 = _nw180
		var _1197_v1 D2
		_ = _1197_v1
		_1197_v1 = Companion_D2_.Create_DC6_(p1, _1196_v0)
		var _source16 D2 = (func() D2 {
			if p0 {
				return Companion_D2_.Create_DC6_((_1197_v1).Dtor_cf13(), _1196_v0)
			}
			return _1197_v1
		})()
		_ = _source16
		if _source16.Is_DC6() {
			var _1198___mcc_h0 _dafny.Int = _source16.Get_().(D2_DC6).Cf13
			_ = _1198___mcc_h0
			var _1199___mcc_h1 _dafny.Array = _source16.Get_().(D2_DC6).Cf14
			_ = _1199___mcc_h1
			var _1200_cf14 _dafny.Array = _1199___mcc_h1
			_ = _1200_cf14
			var _1201_cf13 _dafny.Int = _1198___mcc_h0
			_ = _1201_cf13
			var _1202_v2 _dafny.Sequence
			_ = _1202_v2
			_1202_v2 = _dafny.SeqOf((p1).Cmp(p1) != 0, false, p0)
			_1202_v2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0), _1202_v2), _1202_v2)
			(globalState).F0 = p0
			r1 = p0
			var _1203_v3 _dafny.CodePoint
			_ = _1203_v3
			_1203_v3 = _dafny.CodePoint('m')
			var _1204_v4 _dafny.MultiSet
			_ = _1204_v4
			_1204_v4 = _dafny.MultiSetOf(_1203_v3, _1203_v3)
			var _1205_v5 D1
			_ = _1205_v5
			_1205_v5 = Companion_D1_.Create_DC3_(_1204_v4)
			var _1206_v6 _dafny.Set
			_ = _1206_v6
			_1206_v6 = _dafny.SetOf(_1205_v5)
			var _1207_v7 _dafny.Set
			_ = _1207_v7
			_1207_v7 = _dafny.SetOf(_1206_v6, _1206_v6, _1206_v6, _1206_v6, _1206_v6)
			var _1208_v8 _dafny.Map
			_ = _1208_v8
			_1208_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1207_v7)
			_1208_v8 = (_1208_v8).Update(p0, _1207_v7)
		} else if _source16.Is_DC7() {
			var _1209___mcc_h2 _dafny.Sequence = _source16.Get_().(D2_DC7).Cf15
			_ = _1209___mcc_h2
			var _1210___mcc_h3 _dafny.Array = _source16.Get_().(D2_DC7).Cf16
			_ = _1210___mcc_h3
			var _1211___mcc_h4 _dafny.Int = _source16.Get_().(D2_DC7).Cf17
			_ = _1211___mcc_h4
			var _1212___mcc_h5 _dafny.Int = _source16.Get_().(D2_DC7).Cf18
			_ = _1212___mcc_h5
			var _1213_cf18 _dafny.Int = _1212___mcc_h5
			_ = _1213_cf18
			var _1214_cf17 _dafny.Int = _1211___mcc_h4
			_ = _1214_cf17
			var _1215_cf16 _dafny.Array = _1210___mcc_h3
			_ = _1215_cf16
			var _1216_cf15 _dafny.Sequence = _1209___mcc_h2
			_ = _1216_cf15
			var _1217_v9 _dafny.Sequence
			_ = _1217_v9
			_1217_v9 = _dafny.UnicodeSeqOfUtf8Bytes("y")
			var _1218_v10 _dafny.MultiSet
			_ = _1218_v10
			_1218_v10 = _dafny.MultiSetOf(!_dafny.Companion_Sequence_.Equal(_1217_v9, _1217_v9))
			_1218_v10 = (_1218_v10).Union(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0), _1216_cf15)))
			_1217_v9 = _dafny.Companion_Sequence_.Concatenate(_1217_v9, _1217_v9)
			var _1219_v11 _dafny.Map
			_ = _1219_v11
			_1219_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1217_v9)
			var _1220_v12 _dafny.CodePoint
			_ = _1220_v12
			_1220_v12 = _dafny.CodePoint('i')
			var _1221_v13 _dafny.Map
			_ = _1221_v13
			_1221_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1217_v9)
			_1219_v11 = (_1219_v11).Update(_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm18(p0, globalState), _1220_v12), (func() _dafny.Sequence {
				if (_1221_v13).Contains((_this).F6()) {
					return (_1221_v13).Get((_this).F6()).(_dafny.Sequence)
				}
				return _1217_v9
			})())
			if p0 {
				var _1222_v14 _dafny.MultiSet
				_ = _1222_v14
				_1222_v14 = _dafny.MultiSetOf(_1213_cf18, (_this).F15())
				var _1223_v15 _dafny.Map
				_ = _1223_v15
				_1223_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1222_v14).IsProperSubsetOf(Companion_Default___.Fm19(p1, globalState)), (_this).F15())
				_1223_v15 = (_1223_v15).Update(_dafny.Companion_Sequence_.IsPrefixOf(_1217_v9, _dafny.Companion_Sequence_.Update(_1217_v9, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1217_v9).Cardinality()), _dafny.IntOfUint32((_1217_v9).Cardinality()))).Uint32(), _dafny.CodePoint('x'))), (_dafny.Zero).Minus((_this).F15()))
				var _1224_v16 _dafny.MultiSet
				_ = _1224_v16
				_1224_v16 = _dafny.MultiSetOf(_1220_v12)
				var _1225_v17 _dafny.Map
				_ = _1225_v17
				_1225_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1224_v16).Contains(_1220_v12) {
						return (_1224_v16).Multiplicity(_1220_v12)
					}
					return (_this).F6()
				})(), _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("dgs"), _1217_v9))
				var _1226_v18 _dafny.Set
				_ = _1226_v18
				_1226_v18 = _dafny.SetOf(p0, p0)
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_1196_v0), 0))
				_ = _index183
				(_1196_v0).ArraySet1(((_1226_v18).Union(_1226_v18)).Cardinality(), (_index183).Int())
				var _1227_v19 _dafny.Map
				_ = _1227_v19
				_1227_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1220_v12)
				var _1228_v20 _dafny.Set
				_ = _1228_v20
				_1228_v20 = _dafny.SetOf(_dafny.IntOfInt64(58), (_this).F6())
				var _1229_v21 D0
				_ = _1229_v21
				_1229_v21 = Companion_D0_.Create_DC1_(!(!_dafny.Companion_Sequence_.Contains(_1217_v9, (func() _dafny.CodePoint {
					if (_1227_v19).Contains(p0) {
						return (_1227_v19).Get(p0).(_dafny.CodePoint)
					}
					return Companion_Default___.Fm0((_this).F15(), _1228_v20, !(p0), globalState)
				})())), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_1213_cf18))), _1220_v12)
				var _1230_v22 _dafny.Array
				_ = _1230_v22
				var _nw181 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw181
				_1230_v22 = _nw181
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_1196_v0), 0))
				_ = _index184
				var _rhs210 _dafny.Map = _1225_v17
				_ = _rhs210
				var _rhs211 _dafny.Int = p1
				_ = _rhs211
				var _rhs212 D0 = Companion_Default___.Fm20(p0, (p0) || (p0), p0, globalState)
				_ = _rhs212
				var _rhs213 _dafny.Array = _1230_v22
				_ = _rhs213
				var _lhs173 _dafny.Array = _1196_v0
				_ = _lhs173
				var _lhs174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_1196_v0), 0))
				_ = _lhs174
				_1225_v17 = _rhs210
				(_lhs173).ArraySet1(_rhs211, (_lhs174).Int())
				_1229_v21 = _rhs212
				_1230_v22 = _rhs213
				r1 = Companion_Default___.Fm3(globalState)
				_1213_cf18 = (((_1196_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_1196_v0), 0))).Int()).(_dafny.Int)).Times((_this).F6())).Times(_1213_cf18)
				var _1231_v23 _dafny.Array
				_ = _1231_v23
				var _nwElement0_30 _dafny.Sequence = _1216_cf15
				_ = _nwElement0_30
				var _nw182 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(11))
				_ = _nw182
				(_nw182).ArraySet1(_nwElement0_30, 0)
				(_nw182).ArraySet1(_1216_cf15, 1)
				(_nw182).ArraySet1(_1216_cf15, 2)
				(_nw182).ArraySet1(_dafny.SeqOf(p0), 3)
				(_nw182).ArraySet1(Companion_Default___.Fm4((_this).F6(), p1, _1213_cf18, globalState), 4)
				(_nw182).ArraySet1(_1216_cf15, 5)
				(_nw182).ArraySet1(_1216_cf15, 6)
				(_nw182).ArraySet1(_dafny.SeqOf(p0), 7)
				(_nw182).ArraySet1(_1216_cf15, 8)
				(_nw182).ArraySet1(_dafny.SeqOf(p0, p0), 9)
				(_nw182).ArraySet1(_1216_cf15, 10)
				_1231_v23 = _nw182
				var _1232_v24 _dafny.Map
				_ = _1232_v24
				_1232_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_1228_v20).Cardinality(), _dafny.IntOfUint32((_1217_v9).Cardinality())), _1231_v23)
				_1232_v24 = (_1232_v24).Update((_dafny.Zero).Minus((_this).F6()), _1231_v23)
			} else {
				var _1233_v25 *C1
				_ = _1233_v25
				var _nw183 *C1 = New_C1_()
				_ = _nw183
				_nw183.Ctor__(_dafny.IntOfInt64(-736))
				_1233_v25 = _nw183
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_1196_v0), 0))
				_ = _index185
				(_1196_v0).ArraySet1(_1214_cf17, (_index185).Int())
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_1196_v0), 0))
				_ = _index186
				(_1196_v0).ArraySet1(_1214_cf17, (_index186).Int())
				_1220_v12 = _1220_v12
				(globalState).F0 = p0
				r1 = false
			}
		} else if _source16.Is_DC8() {
			var _1234___mcc_h6 bool = _source16.Get_().(D2_DC8).Cf19
			_ = _1234___mcc_h6
			var _1235___mcc_h7 bool = _source16.Get_().(D2_DC8).Cf20
			_ = _1235___mcc_h7
			var _1236_cf20 bool = _1235___mcc_h7
			_ = _1236_cf20
			var _1237_cf19 bool = _1234___mcc_h6
			_ = _1237_cf19
			var _1238_v26 _dafny.Array
			_ = _1238_v26
			var _nwElement0_31 bool = _1237_cf19
			_ = _nwElement0_31
			var _nw184 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(5))
			_ = _nw184
			(_nw184).ArraySet1(_nwElement0_31, 0)
			(_nw184).ArraySet1(false, 1)
			(_nw184).ArraySet1(_1237_cf19, 2)
			(_nw184).ArraySet1(false, 3)
			(_nw184).ArraySet1(p0, 4)
			_1238_v26 = _nw184
			var _1239_v27 _dafny.Sequence
			_ = _1239_v27
			_1239_v27 = _dafny.UnicodeSeqOfUtf8Bytes("dusfea")
			var _1240_v28 _dafny.Sequence
			_ = _1240_v28
			_1240_v28 = _dafny.SeqOf(_1239_v27, _1239_v27)
			var _1241_v29 _dafny.Map
			_ = _1241_v29
			_1241_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(622), (_1240_v28).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1240_v28).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _1242_v30 *C3
			_ = _1242_v30
			var _nw185 *C3 = New_C3_()
			_ = _nw185
			_nw185.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wbwsr")).Cardinality()), p1), (func() _dafny.Array {
				if _1236_cf20 {
					return _1238_v26
				}
				return _1238_v26
			})(), _1241_v29, (_this).F6())
			_1242_v30 = _nw185
			var _1243_v31 *C1
			_ = _1243_v31
			var _nw186 *C1 = New_C1_()
			_ = _nw186
			_nw186.Ctor__(p1)
			_1243_v31 = _nw186
			var _1244_v32 _dafny.Map
			_ = _1244_v32
			_1244_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(427))
			var _1245_v33 _dafny.CodePoint
			_ = _1245_v33
			_1245_v33 = _dafny.CodePoint('j')
			var _1246_v34 _dafny.Map
			_ = _1246_v34
			_1246_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1245_v33, p0)
			r0 = (_1244_v32).Update(p0, (_1246_v34).Cardinality())
			var _1247_v35 _dafny.Sequence
			_ = _1247_v35
			_1247_v35 = _dafny.SeqOf(_1237_cf19, _1237_cf19)
			var _1248_v36 _dafny.Map
			_ = _1248_v36
			_1248_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1247_v35)
			var _1249_v37 _dafny.Map
			_ = _1249_v37
			_1249_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _1236_cf20)
			var _1250_v38 _dafny.Sequence
			_ = _1250_v38
			_1250_v38 = _dafny.SeqOf(_1249_v37)
			var _1251_v39 _dafny.Map
			_ = _1251_v39
			_1251_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1248_v36).Contains(!(true)) {
					return (_1248_v36).Get(!(true)).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Update(_1247_v35, (Companion_Default___.SafeIndex((_this).F15(), _dafny.IntOfUint32((_1247_v35).Cardinality()))).Uint32(), p0)
			})()).Cardinality())).Cmp(_dafny.IntOfInt64(87)) == 0, !_dafny.Companion_Sequence_.Equal(_1250_v38, _1250_v38))
			var _rhs214 _dafny.Map = _1251_v39
			_ = _rhs214
			var _rhs215 _dafny.Array = _1196_v0
			_ = _rhs215
			var _rhs216 _dafny.Int = p1
			_ = _rhs216
			var _lhs175 *GlobalState = globalState
			_ = _lhs175
			_1251_v39 = _rhs214
			_1196_v0 = _rhs215
			_lhs175.F5 = _rhs216
		} else {
			var _1252___mcc_h8 _dafny.Map = _source16.Get_().(D2_DC5).Cf12
			_ = _1252___mcc_h8
			var _1253_cf12 _dafny.Map = _1252___mcc_h8
			_ = _1253_cf12
			var _1254_v40 _dafny.Set
			_ = _1254_v40
			_1254_v40 = _dafny.SetOf(p0)
			_1254_v40 = _1254_v40
			var _1255_v41 _dafny.MultiSet
			_ = _1255_v41
			_1255_v41 = _dafny.MultiSetOf(p1)
			var _1256_v42 _dafny.CodePoint
			_ = _1256_v42
			_1256_v42 = _dafny.CodePoint('d')
			var _1257_v43 _dafny.Set
			_ = _1257_v43
			_1257_v43 = _dafny.SetOf(Companion_Default___.Fm42(p0, _1255_v41, p1, _1256_v42, globalState))
			var _1258_v44 _dafny.Sequence
			_ = _1258_v44
			_1258_v44 = _dafny.UnicodeSeqOfUtf8Bytes("s")
			var _1259_v45 _dafny.MultiSet
			_ = _1259_v45
			_1259_v45 = _dafny.MultiSetOf(_dafny.SeqOf((_this).F15(), _dafny.IntOfUint32((_1258_v44).Cardinality())))
			var _1260_v46 _dafny.Map
			_ = _1260_v46
			_1260_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1257_v43, (func() _dafny.Int {
				if (_1259_v45).Contains(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(Companion_Default___.Fm2(_dafny.IntOfInt64(-138), globalState)), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqOf(Companion_Default___.Fm2(_dafny.IntOfInt64(-138), globalState))).Cardinality()))).Uint32(), (_this).F6())) {
					return (_1259_v45).Multiplicity(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(Companion_Default___.Fm2(_dafny.IntOfInt64(-138), globalState)), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqOf(Companion_Default___.Fm2(_dafny.IntOfInt64(-138), globalState))).Cardinality()))).Uint32(), (_this).F6()))
				}
				return _dafny.IntOfUint32((_1258_v44).Cardinality())
			})())
			_1260_v46 = (_1260_v46).Update((_1257_v43).Intersection(func() _dafny.Set {
				var _coll36 = _dafny.NewBuilder()
				_ = _coll36
				for _iter38 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1253_cf12, p1)).Keys().Elements()); ; {
					_compr_36, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _1261_v47 _dafny.Map
					_1261_v47 = interface{}(_compr_36).(_dafny.Map)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1253_cf12, p1)).Contains(_1261_v47) {
						_coll36.Add(_1261_v47)
					}
				}
				return _coll36.ToSet()
			}()), (_this).F6())
			var _1262_v48 _dafny.Sequence
			_ = _1262_v48
			_1262_v48 = _dafny.SeqOf(p0)
			var _1263_v49 _dafny.Map
			_ = _1263_v49
			_1263_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1262_v48, p0)
			var _1264_v50 _dafny.Set
			_ = _1264_v50
			_1264_v50 = _dafny.SetOf(_1196_v0)
			var _1265_v51 _dafny.Map
			_ = _1265_v51
			_1265_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1263_v49, _1264_v50)
			_1265_v51 = (_1265_v51).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p0), Companion_Default___.Fm3(globalState)), _1264_v50)
			var _1266_v52 _dafny.Map
			_ = _1266_v52
			_1266_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, ((_this).F15()).Cmp(p1) == 0)
			_1266_v52 = (_1266_v52).Update(p0, p0)
		}
		var _1267_v53 _dafny.Array
		_ = _1267_v53
		var _nw187 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
		_ = _nw187
		_1267_v53 = _nw187
		var _1268_v54 _dafny.Array
		_ = _1268_v54
		_1268_v54 = _1267_v53
		var _1269_v55 _dafny.Map
		_ = _1269_v55
		_1269_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1268_v54, p0)
		var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1196_v0), 0))
		_ = _index187
		(_1196_v0).ArraySet1((_1269_v55).Cardinality(), (_index187).Int())
		var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1196_v0), 0))
		_ = _index188
		(_1196_v0).ArraySet1(p1, (_index188).Int())
		var _1270_v56 _dafny.Map
		_ = _1270_v56
		_1270_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1196_v0, ((_1196_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1196_v0), 0))).Int()).(_dafny.Int)).Times((_1196_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1196_v0), 0))).Int()).(_dafny.Int)))
		var _1271_v57 _dafny.Map
		_ = _1271_v57
		_1271_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1196_v0)
		_1270_v56 = (_1270_v56).Update((func() _dafny.Array {
			if (_1271_v57).Contains(false) {
				return (_1271_v57).Get(false).(_dafny.Array)
			}
			return _1196_v0
		})(), (_this).F15())
		var _1272_v58 _dafny.CodePoint
		_ = _1272_v58
		_1272_v58 = _dafny.CodePoint('n')
		var _1273_v59 _dafny.Map
		_ = _1273_v59
		_1273_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1272_v58, _1267_v53)
		_1273_v59 = (_1273_v59).Update(_1272_v58, _1267_v53)
		var _1274_v60 _dafny.Map
		_ = _1274_v60
		_1274_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F15())
		var _1275_v61 _dafny.Sequence
		_ = _1275_v61
		_1275_v61 = _dafny.SeqOf((_1274_v60).Cardinality(), (_this).F15())
		var _1276_v62 _dafny.Sequence
		_ = _1276_v62
		_1276_v62 = _dafny.SeqOf(p0)
		var _1277_v63 _dafny.Map
		_ = _1277_v63
		_1277_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_1276_v62).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm1(globalState)).Cardinality(), _dafny.IntOfUint32((_1276_v62).Cardinality()))).Uint32()).(bool))
		var _1278_v64 _dafny.Map
		_ = _1278_v64
		_1278_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1277_v63, p0)
		r1 = (func() bool {
			if p0 {
				return (_dafny.IntOfUint32((_1275_v61).Cardinality())).Cmp((Companion_Default___.Fm33(Companion_D3_.Create_DC10_(_1278_v64), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(863))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}(func(_1279_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				}))).Cardinality()), !(p0), (func() bool {
					if (_1277_v63).Contains((_1196_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1196_v0), 0))).Int()).(_dafny.Int)) {
						return (_1277_v63).Get((_1196_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1196_v0), 0))).Int()).(_dafny.Int)).(bool)
					}
					return false
				})(), globalState)).Cardinality()) < 0
			}
			return p0
		})()
		var _1280_v65 _dafny.Set
		_ = _1280_v65
		_1280_v65 = _dafny.SetOf(false, p0)
		var _1281_v66 D11
		_ = _1281_v66
		_1281_v66 = Companion_D11_.Create_DC33_((_this).F15(), p0, (_this).F15(), p0, _1280_v65)
		var _1282_v67 _dafny.MultiSet
		_ = _1282_v67
		_1282_v67 = _dafny.MultiSetOf((_1281_v66).Dtor_cf64(), (_this).F15(), _dafny.IntOfInt64(112), (_this).F6(), _dafny.IntOfInt64(477))
		var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1196_v0), 0))
		_ = _index189
		(_1196_v0).ArraySet1(((p1).Minus((_dafny.MultiSetOf((_1196_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1196_v0), 0))).Int()).(_dafny.Int))).Cardinality())).Times((Companion_Default___.Fm42(p0, _1282_v67, (_1282_v67).Cardinality(), _1272_v58, globalState)).Cardinality()), (_index189).Int())
		r0 = (Companion_D14_.Create_DC36_(_1274_v60)).Dtor_cf69()
		r1 = p0
		return r0, r1
	}
}
func (_this *C5) M11(p0 bool, p1 bool, p2 _dafny.Map, globalState *GlobalState) (_dafny.Int, _dafny.Sequence, bool, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _1283_v0 _dafny.Sequence
		_ = _1283_v0
		_1283_v0 = _dafny.UnicodeSeqOfUtf8Bytes("sc")
		var _1284_v1 _dafny.Sequence
		_ = _1284_v1
		_1284_v1 = _dafny.SeqOf((_this).F15())
		var _1285_v2 _dafny.CodePoint
		_ = _1285_v2
		_1285_v2 = _dafny.CodePoint('o')
		var _1286_v3 D9
		_ = _1286_v3
		_1286_v3 = Companion_D9_.Create_DC25_((_this).F6(), (_this).F6(), (_this).F15(), _1285_v2)
		var _1287_v4 D9
		_ = _1287_v4
		_1287_v4 = Companion_D9_.Create_DC27_(_1286_v3)
		var _1288_v5 D9
		_ = _1288_v5
		_1288_v5 = Companion_D9_.Create_DC27_(_1287_v4)
		var _1289_v6 _dafny.Array
		_ = _1289_v6
		var _nwElement0_32 D9 = Companion_D9_.Create_DC27_(Companion_D9_.Create_DC25_(_dafny.IntOfUint32((_1283_v0).Cardinality()), _dafny.IntOfUint32((_1284_v1).Cardinality()), (_this).F6(), _1285_v2))
		_ = _nwElement0_32
		var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(6))
		_ = _nw188
		(_nw188).ArraySet1(_nwElement0_32, 0)
		(_nw188).ArraySet1(_1288_v5, 1)
		(_nw188).ArraySet1(_1288_v5, 2)
		(_nw188).ArraySet1(_1288_v5, 3)
		(_nw188).ArraySet1(_1288_v5, 4)
		(_nw188).ArraySet1(_1288_v5, 5)
		_1289_v6 = _nw188
		var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_1289_v6), 0))
		_ = _index190
		(_1289_v6).ArraySet1(_1288_v5, (_index190).Int())
		var _1290_v7 _dafny.MultiSet
		_ = _1290_v7
		_1290_v7 = _dafny.MultiSetOf(true)
		var _1291_v8 _dafny.Map
		_ = _1291_v8
		_1291_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1290_v7)
		var _1292_v9 _dafny.Sequence
		_ = _1292_v9
		_1292_v9 = _dafny.SeqOf((func() _dafny.MultiSet {
			if p1 {
				return _1290_v7
			}
			return (func() _dafny.MultiSet {
				if (_1291_v8).Contains(p1) {
					return (_1291_v8).Get(p1).(_dafny.MultiSet)
				}
				return _1290_v7
			})()
		})())
		var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_1289_v6), 0))
		_ = _index191
		var _rhs217 _dafny.Int = ((_this).F6()).Plus(_dafny.IntOfInt64(10))
		_ = _rhs217
		var _rhs218 D9 = _1288_v5
		_ = _rhs218
		var _rhs219 _dafny.MultiSet = (_1292_v9).Select((Companion_Default___.SafeIndex(((_this).F15()).Plus((_1284_v1).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1284_v1).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.IntOfUint32((_1292_v9).Cardinality()))).Uint32()).(_dafny.MultiSet)
		_ = _rhs219
		var _lhs176 _dafny.Array = _1289_v6
		_ = _lhs176
		var _lhs177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_1289_v6), 0))
		_ = _lhs177
		r3 = _rhs217
		(_lhs176).ArraySet1(_rhs218, (_lhs177).Int())
		_1290_v7 = _rhs219
		var _hi9 _dafny.Int = (_1284_v1).Select((Companion_Default___.SafeIndex((_this).F15(), _dafny.IntOfUint32((_1284_v1).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _hi9
		for _1293_i0 := _dafny.IntOfInt64(417); _1293_i0.Cmp(_hi9) < 0; _1293_i0 = _1293_i0.Plus(_dafny.One) {
			r0 = ((_dafny.Zero).Minus((_this).F15())).Minus((func() _dafny.Int {
				if (_1290_v7).Contains(p0) {
					return (_1290_v7).Multiplicity(p0)
				}
				return (_this).F6()
			})())
			var _1294_v10 _dafny.Map
			_ = _1294_v10
			_1294_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(863)).Cmp((_dafny.Zero).Minus((_this).F15())) == 0, p1)
			_1294_v10 = (_1294_v10).Update(p1, Companion_Default___.Fm3(globalState))
			var _1295_v11 _dafny.MultiSet
			_ = _1295_v11
			_1295_v11 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1283_v0).Cardinality()))
			var _1296_v12 _dafny.MultiSet
			_ = _1296_v12
			_1296_v12 = _dafny.MultiSetOf(_1295_v11, _1295_v11)
			var _1297_v13 D4
			_ = _1297_v13
			_1297_v13 = Companion_D4_.Create_DC13_(_1284_v1, p0, (_1284_v1).Select((Companion_Default___.SafeIndex((_this).F15(), _dafny.IntOfUint32((_1284_v1).Cardinality()))).Uint32()).(_dafny.Int))
			var _1298_v14 _dafny.Array
			_ = _1298_v14
			var _nwElement0_33 bool = p1
			_ = _nwElement0_33
			var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(18))
			_ = _nw189
			(_nw189).ArraySet1(_nwElement0_33, 0)
			(_nw189).ArraySet1(p1, 1)
			(_nw189).ArraySet1(p0, 2)
			(_nw189).ArraySet1(p0, 3)
			(_nw189).ArraySet1(p0, 4)
			(_nw189).ArraySet1(p0, 5)
			(_nw189).ArraySet1(p0, 6)
			(_nw189).ArraySet1((_1297_v13).Dtor_cf26(), 7)
			(_nw189).ArraySet1(p1, 8)
			(_nw189).ArraySet1(false, 9)
			(_nw189).ArraySet1(p1, 10)
			(_nw189).ArraySet1(p0, 11)
			(_nw189).ArraySet1(true, 12)
			(_nw189).ArraySet1(p1, 13)
			(_nw189).ArraySet1(p1, 14)
			(_nw189).ArraySet1(p1, 15)
			(_nw189).ArraySet1(p1, 16)
			(_nw189).ArraySet1(!(p0), 17)
			_1298_v14 = _nw189
			var _1299_v15 _dafny.Map
			_ = _1299_v15
			_1299_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1285_v2, _1298_v14)
			var _1300_v16 *C2
			_ = _1300_v16
			var _nw190 *C2 = New_C2_()
			_ = _nw190
			_nw190.Ctor__((_this).F6(), _1296_v12, (func() _dafny.Array {
				if (_1299_v15).Contains(_1285_v2) {
					return (_1299_v15).Get(_1285_v2).(_dafny.Array)
				}
				return _1298_v14
			})(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), Companion_Default___.Fm18(p1, globalState)))
			_1300_v16 = _nw190
			(globalState).F0 = false
		}
		var _source17 D5 = Companion_D5_.Create_DC17_(_1284_v1, p0)
		_ = _source17
		if _source17.Is_DC17() {
			var _1301___mcc_h0 _dafny.Sequence = _source17.Get_().(D5_DC17).Cf36
			_ = _1301___mcc_h0
			var _1302___mcc_h1 bool = _source17.Get_().(D5_DC17).Cf37
			_ = _1302___mcc_h1
			var _1303_cf37 bool = _1302___mcc_h1
			_ = _1303_cf37
			var _1304_cf36 _dafny.Sequence = _1301___mcc_h0
			_ = _1304_cf36
			var _1305_v17 _dafny.Array
			_ = _1305_v17
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_31
			var _nw191 _dafny.Array
			_ = _nw191
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw191 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) bool = (func(_1306_p1 bool) func(_dafny.Int) bool {
					return func(_1307_i1 _dafny.Int) bool {
						return (func() bool {
							if _1306_p1 {
								return true
							}
							return _1306_p1
						})()
					}
				})(p1)
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
			_1305_v17 = _nw191
			var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1305_v17), 0))
			_ = _index192
			(_1305_v17).ArraySet1(!(p0), (_index192).Int())
			var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1305_v17), 0))
			_ = _index193
			(_1305_v17).ArraySet1(p0, (_index193).Int())
			var _1308_v18 _dafny.Sequence
			_ = _1308_v18
			_1308_v18 = _dafny.SeqOf(_1303_cf37)
			var _1309_v19 _dafny.Map
			_ = _1309_v19
			_1309_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), (_1305_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1305_v17), 0))).Int()).(bool))
			var _1310_v20 _dafny.Set
			_ = _1310_v20
			_1310_v20 = _dafny.SetOf((_1309_v19).Cardinality(), _dafny.IntOfInt64(667), (_this).F6())
			var _1311_v21 _dafny.Sequence
			_ = _1311_v21
			_1311_v21 = _dafny.SeqOf(_1310_v20, _1310_v20)
			var _1312_v22 _dafny.Array
			_ = _1312_v22
			var _nwElement0_34 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1303_cf37, true)).Cardinality()
			_ = _nwElement0_34
			var _nw192 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(22))
			_ = _nw192
			(_nw192).ArraySet1(_nwElement0_34, 0)
			(_nw192).ArraySet1((_this).F15(), 1)
			(_nw192).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1308_v18, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1304_cf36).Cardinality()), _dafny.IntOfUint32((_1308_v18).Cardinality()))).Uint32(), p0)).Cardinality()), 2)
			(_nw192).ArraySet1((_this).F15(), 3)
			(_nw192).ArraySet1(((_1311_v21).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1311_v21).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), 4)
			(_nw192).ArraySet1(_dafny.IntOfUint32((_1283_v0).Cardinality()), 5)
			(_nw192).ArraySet1((_this).F6(), 6)
			(_nw192).ArraySet1(Companion_Default___.Fm2((_this).F6(), globalState), 7)
			(_nw192).ArraySet1((_this).F15(), 8)
			(_nw192).ArraySet1((_this).F6(), 9)
			(_nw192).ArraySet1((_this).F6(), 10)
			(_nw192).ArraySet1((_this).F15(), 11)
			(_nw192).ArraySet1((_dafny.Zero).Minus((_this).F15()), 12)
			(_nw192).ArraySet1((_this).F6(), 13)
			(_nw192).ArraySet1((_dafny.Zero).Minus((_this).F15()), 14)
			(_nw192).ArraySet1((_1284_v1).Select((Companion_Default___.SafeIndex((_this).F15(), _dafny.IntOfUint32((_1284_v1).Cardinality()))).Uint32()).(_dafny.Int), 15)
			(_nw192).ArraySet1(Companion_Default___.Fm2((_this).F6(), globalState), 16)
			(_nw192).ArraySet1(_dafny.IntOfInt64(697), 17)
			(_nw192).ArraySet1((_this).F15(), 18)
			(_nw192).ArraySet1((_this).F15(), 19)
			(_nw192).ArraySet1(_dafny.IntOfInt64(576), 20)
			(_nw192).ArraySet1((_this).F15(), 21)
			_1312_v22 = _nw192
			var _1313_v23 _dafny.Map
			_ = _1313_v23
			_1313_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1312_v22, (_1305_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1305_v17), 0))).Int()).(bool))
			_1303_cf37 = !((func() bool {
				if (_1313_v23).Contains(_1312_v22) {
					return (_1313_v23).Get(_1312_v22).(bool)
				}
				return p0
			})())
			r2 = _dafny.Companion_Sequence_.IsPrefixOf(_1283_v0, _1283_v0)
			_1309_v19 = (_1309_v19).Update((_this).F6(), (_1305_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1305_v17), 0))).Int()).(bool))
		} else {
			var _1314___mcc_h2 _dafny.Set = _source17.Get_().(D5_DC16).Cf35
			_ = _1314___mcc_h2
			var _1315_cf35 _dafny.Set = _1314___mcc_h2
			_ = _1315_cf35
			var _1316_v24 _dafny.Array
			_ = _1316_v24
			var _nw193 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw193
			_1316_v24 = _nw193
			_1316_v24 = _1316_v24
			(globalState).F5 = (_this).F15()
			var _1317_v25 _dafny.Sequence
			_ = _1317_v25
			_1317_v25 = _dafny.SeqOf(p1)
			var _1318_v26 _dafny.Map
			_ = _1318_v26
			_1318_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1317_v25, (_this).F15())
			_1318_v26 = (_1318_v26).Update(_1317_v25, _dafny.IntOfInt64(-172))
			r2 = !(!(p1)) || (((_this).F15()).Cmp((_this).F15()) == 0)
		}
		var _1319_i2 _dafny.Int
		_ = _1319_i2
		_1319_i2 = _dafny.Zero
		{
			for p0 {
				{
					if (_1319_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1319_i2 = (_1319_i2).Plus(_dafny.One)
					var _1320_v27 _dafny.Sequence
					_ = _1320_v27
					_1320_v27 = _dafny.SeqOf((func() bool {
						if p0 {
							return p1
						}
						return p1
					})())
					var _1321_v28 _dafny.Map
					_ = _1321_v28
					_1321_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _1320_v27)
					_1320_v27 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm3(globalState)), (func() _dafny.Sequence {
						if (_1321_v28).Contains((_this).F6()) {
							return (_1321_v28).Get((_this).F6()).(_dafny.Sequence)
						}
						return _1320_v27
					})())
					var _1322_v29 _dafny.Array
					_ = _1322_v29
					var _len0_32 _dafny.Int = _dafny.IntOfInt64(14)
					_ = _len0_32
					var _nw194 _dafny.Array
					_ = _nw194
					if _len0_32.Cmp(_dafny.Zero) == 0 {
						_nw194 = _dafny.NewArray(_len0_32)
					} else {
						var _init32 func(_dafny.Int) _dafny.Int = func(_1323_i3 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1323_i3, (_this).F15())
						}
						_ = _init32
						var _element0_32 = _init32(_dafny.Zero)
						_ = _element0_32
						_nw194 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
						(_nw194).ArraySet1(_element0_32, 0)
						var _nativeLen0_32 = (_len0_32).Int()
						_ = _nativeLen0_32
						for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
							(_nw194).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
						}
					}
					_1322_v29 = _nw194
					var _1324_v30 D2
					_ = _1324_v30
					_1324_v30 = Companion_D2_.Create_DC6_((_this).F6(), _1322_v29)
					var _1325_v31 _dafny.Sequence
					_ = _1325_v31
					_1325_v31 = _dafny.SeqOf(_1322_v29, _1322_v29)
					var _pat_let_tv21 = _1322_v29
					_ = _pat_let_tv21
					var _pat_let_tv22 = _1322_v29
					_ = _pat_let_tv22
					var _1326_v32 _dafny.Array
					_ = _1326_v32
					var _nwElement0_35 _dafny.Array = _1322_v29
					_ = _nwElement0_35
					var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(11))
					_ = _nw195
					(_nw195).ArraySet1(_nwElement0_35, 0)
					(_nw195).ArraySet1(_1322_v29, 1)
					(_nw195).ArraySet1(_1322_v29, 2)
					(_nw195).ArraySet1(_1322_v29, 3)
					(_nw195).ArraySet1(_1322_v29, 4)
					(_nw195).ArraySet1((func(_pat_let25_0 D2) D2 {
						return func(_1329_dt__update__tmp_h1 D2) D2 {
							return func(_pat_let28_0 _dafny.Array) D2 {
								return func(_1330_dt__update_hcf14_h1 _dafny.Array) D2 {
									return Companion_D2_.Create_DC6_((_1329_dt__update__tmp_h1).Dtor_cf13(), _1330_dt__update_hcf14_h1)
								}(_pat_let28_0)
							}(_pat_let_tv22)
						}(_pat_let25_0)
					}(func(_pat_let26_0 D2) D2 {
						return func(_1327_dt__update__tmp_h0 D2) D2 {
							return func(_pat_let27_0 _dafny.Array) D2 {
								return func(_1328_dt__update_hcf14_h0 _dafny.Array) D2 {
									return Companion_D2_.Create_DC6_((_1327_dt__update__tmp_h0).Dtor_cf13(), _1328_dt__update_hcf14_h0)
								}(_pat_let27_0)
							}(_pat_let_tv21)
						}(_pat_let26_0)
					}(_1324_v30))).Dtor_cf14(), 5)
					(_nw195).ArraySet1((_1325_v31).Select((Companion_Default___.SafeIndex(((_1290_v7).Update(p0, Companion_Default___.Abs((_this).F15()))).Cardinality(), _dafny.IntOfUint32((_1325_v31).Cardinality()))).Uint32()).(_dafny.Array), 6)
					(_nw195).ArraySet1(_1322_v29, 7)
					(_nw195).ArraySet1(_1322_v29, 8)
					(_nw195).ArraySet1(_1322_v29, 9)
					(_nw195).ArraySet1(_1322_v29, 10)
					_1326_v32 = _nw195
					var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1326_v32), 0))
					_ = _index194
					(_1326_v32).ArraySet1(_1322_v29, (_index194).Int())
					var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1326_v32), 0))
					_ = _index195
					(_1326_v32).ArraySet1(_1322_v29, (_index195).Int())
					(globalState).F5 = (_this).F15()
					var _1331_v33 *C1
					_ = _1331_v33
					var _nw196 *C1 = New_C1_()
					_ = _nw196
					_nw196.Ctor__((_this).F15())
					_1331_v33 = _nw196
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1332_v34 _dafny.Sequence
		_ = _1332_v34
		_1332_v34 = _dafny.SeqOf(p0)
		var _1333_v35 _dafny.Sequence
		_ = _1333_v35
		_1333_v35 = _dafny.SeqOf((_1332_v34).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.SetOf(p1)).Cardinality()), _dafny.IntOfUint32((_1332_v34).Cardinality()))).Uint32()).(bool), false)
		var _1334_v37 _dafny.Map
		_ = _1334_v37
		_1334_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			var _coll37 = _dafny.NewMapBuilder()
			_ = _coll37
			for _iter39 := _dafny.Iterate((_1284_v1).Elements()); ; {
				_compr_37, _ok39 := _iter39()
				if !_ok39 {
					break
				}
				var _1335_v36 _dafny.Int
				_1335_v36 = interface{}(_compr_37).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_1284_v1, _1335_v36) {
					_coll37.Add((_1335_v36).Times((_dafny.Zero).Minus((_this).F15())), (_this).F6())
				}
			}
			return _coll37.ToMap()
		}()).Cardinality(), p0)
		var _1336_v38 _dafny.Map
		_ = _1336_v38
		_1336_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1334_v37, p0)
		var _1337_v39 D3
		_ = _1337_v39
		_1337_v39 = Companion_D3_.Create_DC10_(_1336_v38)
		var _pat_let_tv23 = _1336_v38
		_ = _pat_let_tv23
		var _1338_v40 _dafny.Array
		_ = _1338_v40
		var _nwElement0_36 D3 = (func() D3 {
			if (_1332_v34).Select((Companion_Default___.SafeIndex((_this).F15(), _dafny.IntOfUint32((_1332_v34).Cardinality()))).Uint32()).(bool) {
				return _1337_v39
			}
			return func(_pat_let29_0 D3) D3 {
				return func(_1339_dt__update__tmp_h2 D3) D3 {
					return func(_pat_let30_0 _dafny.Map) D3 {
						return func(_1340_dt__update_hcf22_h0 _dafny.Map) D3 {
							return Companion_D3_.Create_DC10_(_1340_dt__update_hcf22_h0)
						}(_pat_let30_0)
					}(_pat_let_tv23)
				}(_pat_let29_0)
			}(_1337_v39)
		})()
		_ = _nwElement0_36
		var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(3))
		_ = _nw197
		(_nw197).ArraySet1(_nwElement0_36, 0)
		(_nw197).ArraySet1(Companion_D3_.Create_DC10_(_1336_v38), 1)
		(_nw197).ArraySet1(_1337_v39, 2)
		_1338_v40 = _nw197
		var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_1338_v40), 0))
		_ = _index196
		(_1338_v40).ArraySet1(_1337_v39, (_index196).Int())
		var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_1338_v40), 0))
		_ = _index197
		(_1338_v40).ArraySet1((func() D3 {
			if (p1) == (p1) {
				return _1337_v39
			}
			return _1337_v39
		})(), (_index197).Int())
		if p1 {
			if true {
				(globalState).F2 = Companion_Default___.Fm30(p1, _1285_v2, globalState)
				var _1341_v41 _dafny.Array
				_ = _1341_v41
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_33
				var _nw198 _dafny.Array
				_ = _nw198
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw198 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) bool = (func(_1342_p1 bool) func(_dafny.Int) bool {
						return func(_1343_i4 _dafny.Int) bool {
							return !(_1342_p1)
						}
					})(p1)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw198 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw198).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw198).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1341_v41 = _nw198
				var _1344_v42 _dafny.Map
				_ = _1344_v42
				_1344_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _dafny.UnicodeSeqOfUtf8Bytes("amxb"))
				var _1345_v43 _dafny.MultiSet
				_ = _1345_v43
				_1345_v43 = _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_1284_v1, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1284_v1).Cardinality()))).Uint32(), _dafny.IntOfInt64(680))))
				var _1346_v44 *C4
				_ = _1346_v44
				var _nw199 *C4 = New_C4_()
				_ = _nw199
				_nw199.Ctor__(Companion_Default___.Fm2(Companion_Default___.Fm2((_this).F6(), globalState), globalState), (_this).F15(), (func() _dafny.Array {
					if p0 {
						return _1341_v41
					}
					return _1341_v41
				})(), (_1344_v42).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(141))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg85 _dafny.Int) interface{} {
						return coer85(arg85)
					}
				}((func(_1347_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1348_i5 _dafny.Int) _dafny.CodePoint {
						return _1347_v2
					}
				})(_1285_v2)))).Cardinality()), _1283_v0)).Update(Companion_Default___.Fm2((_1290_v7).Cardinality(), globalState), _dafny.UnicodeSeqOfUtf8Bytes("coc"))), (_this).F6(), _1345_v43)
				_1346_v44 = _nw199
				(globalState).F5 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1284_v1).Cardinality()), (_this).F6())
				var _1349_v45 D0
				_ = _1349_v45
				_1349_v45 = Companion_D0_.Create_DC1_(p1, (_1346_v44).F18(), _1285_v2)
				var _1350_v46 _dafny.Map
				_ = _1350_v46
				_1350_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
				var _1351_v47 D4
				_ = _1351_v47
				_1351_v47 = Companion_D4_.Create_DC14_(p1, _1350_v46, (_this).F15(), _1350_v46)
				var _1352_v48 _dafny.Sequence
				_ = _1352_v48
				_1352_v48 = _dafny.SeqOf(_1351_v47, Companion_D4_.Create_DC14_(p1, _1350_v46, (_this).F6(), _1350_v46))
				var _1353_v49 _dafny.Sequence
				_ = _1353_v49
				_1353_v49 = _dafny.SeqOf(_1352_v48, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(710))).Uint32(), func(coer86 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
					return func(arg86 _dafny.Int) interface{} {
						return coer86(arg86)
					}
				}((func(_1354_v47 D4) func(_dafny.Int) D4 {
					return func(_1355_i6 _dafny.Int) D4 {
						return _1354_v47
					}
				})(_1351_v47))), _1352_v48, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1352_v48, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1352_v48).Cardinality()))).Uint32(), _1351_v47), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1352_v48, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1352_v48).Cardinality()))).Uint32(), _1351_v47)).Cardinality()))).Uint32(), Companion_D4_.Create_DC14_(p1, _1350_v46, (_dafny.Zero).Minus((_1346_v44).F17()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0))), _1352_v48)
				var _rhs220 bool = (_1349_v45).Dtor_cf2()
				_ = _rhs220
				var _rhs221 _dafny.Int = (_dafny.IntOfUint32((_1353_v49).Cardinality())).Times((_this).F6())
				_ = _rhs221
				var _rhs222 bool = (!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("kx"), _1283_v0)) || (p0)
				_ = _rhs222
				var _rhs223 _dafny.CodePoint = _dafny.CodePoint('t')
				_ = _rhs223
				var _lhs178 *GlobalState = globalState
				_ = _lhs178
				var _lhs179 *GlobalState = globalState
				_ = _lhs179
				var _lhs180 *GlobalState = globalState
				_ = _lhs180
				_lhs178.F0 = _rhs220
				_lhs179.F5 = _rhs221
				_lhs180.F0 = _rhs222
				_1285_v2 = _rhs223
				var _1356_v50 _dafny.Set
				_ = _1356_v50
				_1356_v50 = _dafny.SetOf(p0)
				var _1357_v51 *C1
				_ = _1357_v51
				var _nw200 *C1 = New_C1_()
				_ = _nw200
				_nw200.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(548), (_1346_v44).Fm21((_1356_v50).Cardinality(), globalState)))
				_1357_v51 = _nw200
			} else {
				(globalState).F0 = p0
				var _1358_v52 _dafny.MultiSet
				_ = _1358_v52
				_1358_v52 = _dafny.MultiSetOf(_dafny.IntOfInt64(28), (_this).F15())
				var _1359_v53 _dafny.MultiSet
				_ = _1359_v53
				_1359_v53 = _dafny.MultiSetOf(_1358_v52)
				var _1360_v54 _dafny.Array
				_ = _1360_v54
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_34
				var _nw201 _dafny.Array
				_ = _nw201
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw201 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) bool = func(_1361_i7 _dafny.Int) bool {
						return true
					}
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw201 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw201).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw201).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1360_v54 = _nw201
				var _1362_v55 _dafny.Map
				_ = _1362_v55
				_1362_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F6()), _1283_v0)
				var _1363_v57 *C2
				_ = _1363_v57
				var _nw202 *C2 = New_C2_()
				_ = _nw202
				_nw202.Ctor__(_dafny.IntOfInt64(-676), (_1359_v53).Union(_dafny.MultiSetOf(_1358_v52, _1358_v52)), _1360_v54, (_1362_v55).Merge(func() _dafny.Map {
					var _coll38 = _dafny.NewMapBuilder()
					_ = _coll38
					for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-58), _dafny.IntOfInt64(555))); ; {
						_compr_38, _ok40 := _iter40()
						if !_ok40 {
							break
						}
						var _1364_v56 _dafny.Int
						_1364_v56 = interface{}(_compr_38).(_dafny.Int)
						if ((_dafny.IntOfInt64(-58)).Cmp(_1364_v56) <= 0) && ((_1364_v56).Cmp(_dafny.IntOfInt64(555)) < 0) {
							_coll38.Add(Companion_Default___.SafeModuloInt(_1364_v56, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1358_v52).Cardinality())).Cardinality()), _1283_v0)
						}
					}
					return _coll38.ToMap()
				}()))
				_1363_v57 = _nw202
				var _1365_v59 _dafny.Array
				_ = _1365_v59
				var _nwElement0_37 _dafny.Int = (_1290_v7).Cardinality()
				_ = _nwElement0_37
				var _nw203 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(26))
				_ = _nw203
				(_nw203).ArraySet1(_nwElement0_37, 0)
				(_nw203).ArraySet1((_this).F6(), 1)
				(_nw203).ArraySet1((_this).F15(), 2)
				(_nw203).ArraySet1((_this).F6(), 3)
				(_nw203).ArraySet1((_this).F15(), 4)
				(_nw203).ArraySet1((_this).F15(), 5)
				(_nw203).ArraySet1((_this).F15(), 6)
				(_nw203).ArraySet1((_this).F6(), 7)
				(_nw203).ArraySet1(_dafny.IntOfUint32((_1283_v0).Cardinality()), 8)
				(_nw203).ArraySet1((_this).F6(), 9)
				(_nw203).ArraySet1((_this).F15(), 10)
				(_nw203).ArraySet1((_this).F6(), 11)
				(_nw203).ArraySet1((_1358_v52).Cardinality(), 12)
				(_nw203).ArraySet1(_dafny.IntOfInt64(99), 13)
				(_nw203).ArraySet1((_this).F15(), 14)
				(_nw203).ArraySet1(_dafny.IntOfInt64(-832), 15)
				(_nw203).ArraySet1((_this).F6(), 16)
				(_nw203).ArraySet1((_this).F6(), 17)
				(_nw203).ArraySet1((_1363_v57).Fm21((_this).F15(), globalState), 18)
				(_nw203).ArraySet1((p2).Cardinality(), 19)
				(_nw203).ArraySet1((_this).F6(), 20)
				(_nw203).ArraySet1((_this).F15(), 21)
				(_nw203).ArraySet1((_this).F15(), 22)
				(_nw203).ArraySet1((_this).F15(), 23)
				(_nw203).ArraySet1((_this).F6(), 24)
				(_nw203).ArraySet1((_this).F15(), 25)
				_1365_v59 = _nw203
				var _1366_v60 D2
				_ = _1366_v60
				_1366_v60 = Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1283_v0, (Companion_Default___.SafeIndex((_this).F15(), _dafny.IntOfUint32((_1283_v0).Cardinality()))).Uint32(), _1285_v2)).Cardinality()), _1365_v59)
				var _1367_v61 *C2
				_ = _1367_v61
				var _nw204 *C2 = New_C2_()
				_ = _nw204
				_nw204.Ctor__((_this).F15(), _1359_v53, _1360_v54, func() _dafny.Map {
					var _coll39 = _dafny.NewMapBuilder()
					_ = _coll39
					for _iter41 := _dafny.Iterate((Companion_Default___.Fm35((_1366_v60).Dtor_cf13(), (_this).F6(), p1, (_this).F6(), globalState)).Elements()); ; {
						_compr_39, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _1368_v58 _dafny.Int
						_1368_v58 = interface{}(_compr_39).(_dafny.Int)
						if (Companion_Default___.Fm35((_1366_v60).Dtor_cf13(), (_this).F6(), p1, (_this).F6(), globalState)).Contains(_1368_v58) {
							_coll39.Add((_1368_v58).Plus(_dafny.IntOfUint32((_1332_v34).Cardinality())), _1283_v0)
						}
					}
					return _coll39.ToMap()
				}())
				_1367_v61 = _nw204
				var _1369_v62 _dafny.Map
				_ = _1369_v62
				_1369_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(651), _1365_v59)
				_1369_v62 = (_1369_v62).Update((_this).F6(), _1365_v59)
				var _1370_v63 _dafny.Map
				_ = _1370_v63
				_1370_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1285_v2, p0)
				var _1371_v64 _dafny.Set
				_ = _1371_v64
				_1371_v64 = _dafny.SetOf(p0, p1)
				_1370_v63 = (_1370_v63).Update(_1285_v2, !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_1367_v61).Fm21((_this).F6(), globalState), (Companion_Default___.Fm35((_this).F6(), (_this).F6(), p1, (_this).F6(), globalState)).Cardinality(), (_this).F6(), _dafny.IntOfInt64(991), (_this).F15()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm40((_this).F6(), _1371_v64, globalState)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_1367_v61).Fm21((_this).F6(), globalState), (Companion_Default___.Fm35((_this).F6(), (_this).F6(), p1, (_this).F6(), globalState)).Cardinality(), (_this).F6(), _dafny.IntOfInt64(991), (_this).F15())).Cardinality()))).Uint32(), _dafny.IntOfInt64(409)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(761))).Uint32(), func(coer87 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg87 _dafny.Int) interface{} {
						return coer87(arg87)
					}
				}(func(_1372_i8 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tvotwjin")).Cardinality())
				}))))
			}
			var _1373_v65 _dafny.Sequence
			_ = _1373_v65
			_1373_v65 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-371))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg88 _dafny.Int) interface{} {
					return coer88(arg88)
				}
			}((func(_1374_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1375_i9 _dafny.Int) _dafny.CodePoint {
					return _1374_v2
				}
			})(_1285_v2))), _1283_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(714))).Uint32(), func(coer89 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg89 _dafny.Int) interface{} {
					return coer89(arg89)
				}
			}((func(_1376_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1377_i10 _dafny.Int) _dafny.CodePoint {
					return _1376_v2
				}
			})(_1285_v2))))
			_1373_v65 = (func() _dafny.Sequence {
				if (p1) || (p1) {
					return _1373_v65
				}
				return _1373_v65
			})()
			var _1378_v66 _dafny.Array
			_ = _1378_v66
			var _nw205 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw205
			_1378_v66 = _nw205
			var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_1378_v66), 0))
			_ = _index198
			(_1378_v66).ArraySet1((_1290_v7).IsSubsetOf(Companion_Default___.Fm48(globalState)), (_index198).Int())
			var _1379_v67 _dafny.Array
			_ = _1379_v67
			var _nw206 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(11))
			_ = _nw206
			_1379_v67 = _nw206
			var _1380_v68 _dafny.Map
			_ = _1380_v68
			_1380_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1379_v67, p0)
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_1378_v66), 0))
			_ = _index199
			var _rhs224 bool = p1
			_ = _rhs224
			var _rhs225 bool = ((_this).F15()).Cmp((_this).F15()) > 0
			_ = _rhs225
			var _rhs226 _dafny.Map = _1380_v68
			_ = _rhs226
			var _rhs227 _dafny.Int = (_this).F15()
			_ = _rhs227
			var _rhs228 bool = (_1333_v35).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1333_v35).Cardinality()))).Uint32()).(bool)
			_ = _rhs228
			var _lhs181 _dafny.Array = _1378_v66
			_ = _lhs181
			var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_1378_v66), 0))
			_ = _lhs182
			var _lhs183 *GlobalState = globalState
			_ = _lhs183
			(_lhs181).ArraySet1(_rhs224, (_lhs182).Int())
			r2 = _rhs225
			_1380_v68 = _rhs226
			r3 = _rhs227
			_lhs183.F0 = _rhs228
			var _1381_v69 _dafny.Map
			_ = _1381_v69
			_1381_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F15())
			var _1382_v70 _dafny.Map
			_ = _1382_v70
			_1382_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1378_v66, ((func() _dafny.Int {
				if (_1381_v69).Contains(true) {
					return (_1381_v69).Get(true).(_dafny.Int)
				}
				return (_this).F6()
			})()).Cmp(Companion_Default___.Fm2((_this).F15(), globalState)) > 0)
			_1382_v70 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1378_v66, (_1378_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_1378_v66), 0))).Int()).(bool))).Merge(_1382_v70)
			var _1383_v71 _dafny.MultiSet
			_ = _1383_v71
			_1383_v71 = _dafny.MultiSetOf((_this).F15(), (_this).F15())
			var _1384_v72 _dafny.MultiSet
			_ = _1384_v72
			_1384_v72 = _dafny.MultiSetOf(_1383_v71, _dafny.MultiSetFromSeq(_1284_v1))
			var _1385_v73 _dafny.Map
			_ = _1385_v73
			_1385_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sigwl")).Cardinality())), _1283_v0)
			var _1386_v74 _dafny.Sequence
			_ = _1386_v74
			_1386_v74 = _dafny.SeqOf(_1385_v73)
			var _1387_v75 *C2
			_ = _1387_v75
			var _nw207 *C2 = New_C2_()
			_ = _nw207
			_nw207.Ctor__((_this).F15(), _1384_v72, _1378_v66, (_1386_v74).Select((Companion_Default___.SafeIndex((_this).F15(), _dafny.IntOfUint32((_1386_v74).Cardinality()))).Uint32()).(_dafny.Map))
			_1387_v75 = _nw207
		} else {
			var _1388_v76 _dafny.Set
			_ = _1388_v76
			_1388_v76 = _dafny.SetOf(_1285_v2)
			var _1389_v78 _dafny.Sequence
			_ = _1389_v78
			_1389_v78 = _dafny.SeqOf(_1388_v76, func() _dafny.Set {
				var _coll40 = _dafny.NewBuilder()
				_ = _coll40
				for _iter42 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1285_v2, p0)).Keys().Elements()); ; {
					_compr_40, _ok42 := _iter42()
					if !_ok42 {
						break
					}
					var _1390_v77 _dafny.CodePoint
					_1390_v77 = interface{}(_compr_40).(_dafny.CodePoint)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1285_v2, p0)).Contains(_1390_v77) {
						_coll40.Add(_1390_v77)
					}
				}
				return _coll40.ToSet()
			}())
			(globalState).F5 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1389_v78).Cardinality()), (func() _dafny.Int {
				if p0 {
					return _dafny.IntOfInt64(988)
				}
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(699))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg90 _dafny.Int) interface{} {
						return coer90(arg90)
					}
				}(func(_1391_i11 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('q')
				}))).Cardinality())
			})())
			var _1392_v79 _dafny.Map
			_ = _1392_v79
			_1392_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(554))).Uint32(), func(coer91 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg91 _dafny.Int) interface{} {
					return coer91(arg91)
				}
			}(func(_1393_i12 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-145)
			})), _1284_v1), p0)
			_1392_v79 = (_1392_v79).Update(p0, p0)
			(globalState).F0 = p1
			_1290_v7 = (_1290_v7).Union(_dafny.MultiSetOf(p0))
			var _1394_v80 _dafny.Set
			_ = _1394_v80
			_1394_v80 = _dafny.SetOf((_this).F15())
			var _1395_v81 _dafny.Array
			_ = _1395_v81
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_35
			var _nw208 _dafny.Array
			_ = _nw208
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw208 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) _dafny.Int = func(_1396_i13 _dafny.Int) _dafny.Int {
					return (_1396_i13).Minus((_dafny.Zero).Minus((_this).F15()))
				}
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw208 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw208).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw208).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1395_v81 = _nw208
			var _1397_v82 _dafny.Map
			_ = _1397_v82
			_1397_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _1395_v81)
			var _1398_v83 _dafny.MultiSet
			_ = _1398_v83
			_1398_v83 = _dafny.MultiSetOf((_dafny.Zero).Minus((_this).F6()))
			var _1399_v84 _dafny.MultiSet
			_ = _1399_v84
			_1399_v84 = _dafny.MultiSetOf((_1284_v1).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm2((_1398_v83).Cardinality(), globalState), _dafny.IntOfUint32((_1284_v1).Cardinality()))).Uint32()).(_dafny.Int))
			var _1400_v85 _dafny.Map
			_ = _1400_v85
			_1400_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.CodePoint('a'))
			var _1401_v86 _dafny.Array
			_ = _1401_v86
			var _nwElement0_38 _dafny.Int = (_this).F15()
			_ = _nwElement0_38
			var _nw209 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(16))
			_ = _nw209
			(_nw209).ArraySet1(_nwElement0_38, 0)
			(_nw209).ArraySet1((_1394_v80).Cardinality(), 1)
			(_nw209).ArraySet1((_this).F15(), 2)
			(_nw209).ArraySet1((_this).F6(), 3)
			(_nw209).ArraySet1((_1397_v82).Cardinality(), 4)
			(_nw209).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("loqmtebnn")).Cardinality()), 5)
			(_nw209).ArraySet1((_this).F15(), 6)
			(_nw209).ArraySet1((_this).F6(), 7)
			(_nw209).ArraySet1((_this).F15(), 8)
			(_nw209).ArraySet1((_this).F6(), 9)
			(_nw209).ArraySet1(_dafny.IntOfUint32((_1283_v0).Cardinality()), 10)
			(_nw209).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)).Cardinality(), 11)
			(_nw209).ArraySet1((_this).F6(), 12)
			(_nw209).ArraySet1((_this).F15(), 13)
			(_nw209).ArraySet1((func() _dafny.Int {
				if (_1399_v84).Contains((_1400_v85).Cardinality()) {
					return (_1399_v84).Multiplicity((_1400_v85).Cardinality())
				}
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))
			})(), 14)
			(_nw209).ArraySet1((_this).F15(), 15)
			_1401_v86 = _nw209
			var _1402_v87 D2
			_ = _1402_v87
			_1402_v87 = Companion_D2_.Create_DC6_((_dafny.IntOfUint32((_dafny.SeqOf(_1333_v35)).Cardinality())).Plus((_this).F15()), _1401_v86)
			var _source18 D2 = _1402_v87
			_ = _source18
			if _source18.Is_DC6() {
				var _1403___mcc_h3 _dafny.Int = _source18.Get_().(D2_DC6).Cf13
				_ = _1403___mcc_h3
				var _1404___mcc_h4 _dafny.Array = _source18.Get_().(D2_DC6).Cf14
				_ = _1404___mcc_h4
				var _1405_cf14 _dafny.Array = _1404___mcc_h4
				_ = _1405_cf14
				var _1406_cf13 _dafny.Int = _1403___mcc_h3
				_ = _1406_cf13
				(globalState).F5 = (Companion_Default___.SafeModuloInt((_this).F6(), _dafny.IntOfUint32((_1283_v0).Cardinality()))).Times((_this).F6())
				r3 = ((_this).F6()).Plus((_dafny.Zero).Minus(((_this).F15()).Plus(_dafny.IntOfInt64(337))))
				(globalState).F2 = _1283_v0
				(globalState).F0 = (func() bool {
					if p0 {
						return p0
					}
					return ((_this).F15()).Cmp((_this).F15()) >= 0
				})()
			} else if _source18.Is_DC7() {
				var _1407___mcc_h5 _dafny.Sequence = _source18.Get_().(D2_DC7).Cf15
				_ = _1407___mcc_h5
				var _1408___mcc_h6 _dafny.Array = _source18.Get_().(D2_DC7).Cf16
				_ = _1408___mcc_h6
				var _1409___mcc_h7 _dafny.Int = _source18.Get_().(D2_DC7).Cf17
				_ = _1409___mcc_h7
				var _1410___mcc_h8 _dafny.Int = _source18.Get_().(D2_DC7).Cf18
				_ = _1410___mcc_h8
				var _1411_cf18 _dafny.Int = _1410___mcc_h8
				_ = _1411_cf18
				var _1412_cf17 _dafny.Int = _1409___mcc_h7
				_ = _1412_cf17
				var _1413_cf16 _dafny.Array = _1408___mcc_h6
				_ = _1413_cf16
				var _1414_cf15 _dafny.Sequence = _1407___mcc_h5
				_ = _1414_cf15
				(globalState).F0 = (func() bool {
					if (func() bool {
						if p0 {
							return p0
						}
						return p0
					})() {
						return p1
					}
					return (_1414_cf15).Select((Companion_Default___.SafeIndex((_this).F15(), _dafny.IntOfUint32((_1414_cf15).Cardinality()))).Uint32()).(bool)
				})()
				var _1415_v88 D0
				_ = _1415_v88
				_1415_v88 = Companion_D0_.Create_DC1_(!(p0), _dafny.IntOfInt64(88), _dafny.CodePoint('f'))
				var _1416_v89 _dafny.Map
				_ = _1416_v89
				_1416_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1290_v7).Cardinality())
				var _pat_let_tv24 = _1285_v2
				_ = _pat_let_tv24
				var _1417_v90 _dafny.Array
				_ = _1417_v90
				var _nwElement0_39 D0 = _1415_v88
				_ = _nwElement0_39
				var _nw210 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(26))
				_ = _nw210
				(_nw210).ArraySet1(_nwElement0_39, 0)
				(_nw210).ArraySet1(_1415_v88, 1)
				(_nw210).ArraySet1(_1415_v88, 2)
				(_nw210).ArraySet1(_1415_v88, 3)
				(_nw210).ArraySet1((func() D0 {
					if p1 {
						return _1415_v88
					}
					return Companion_D0_.Create_DC1_(p1, _1411_cf18, _1285_v2)
				})(), 4)
				(_nw210).ArraySet1(_1415_v88, 5)
				(_nw210).ArraySet1(_1415_v88, 6)
				(_nw210).ArraySet1(Companion_Default___.Fm20(p1, !(p0), true, globalState), 7)
				(_nw210).ArraySet1(func(_pat_let31_0 D0) D0 {
					return func(_1418_dt__update__tmp_h3 D0) D0 {
						return func(_pat_let32_0 _dafny.CodePoint) D0 {
							return func(_1419_dt__update_hcf4_h0 _dafny.CodePoint) D0 {
								return Companion_D0_.Create_DC1_((_1418_dt__update__tmp_h3).Dtor_cf2(), (_1418_dt__update__tmp_h3).Dtor_cf3(), _1419_dt__update_hcf4_h0)
							}(_pat_let32_0)
						}(_pat_let_tv24)
					}(_pat_let31_0)
				}(_1415_v88), 8)
				(_nw210).ArraySet1(_1415_v88, 9)
				(_nw210).ArraySet1(_1415_v88, 10)
				(_nw210).ArraySet1(Companion_D0_.Create_DC1_((_1333_v35).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1416_v89).Contains(p1) {
						return (_1416_v89).Get(p1).(_dafny.Int)
					}
					return (_this).F6()
				})(), _dafny.IntOfUint32((_1333_v35).Cardinality()))).Uint32()).(bool), (_this).F15(), _1285_v2), 11)
				(_nw210).ArraySet1(_1415_v88, 12)
				(_nw210).ArraySet1((func() D0 {
					if p1 {
						return _1415_v88
					}
					return _1415_v88
				})(), 13)
				(_nw210).ArraySet1(Companion_Default___.Fm20(p0, p1, p1, globalState), 14)
				(_nw210).ArraySet1(_1415_v88, 15)
				(_nw210).ArraySet1(_1415_v88, 16)
				(_nw210).ArraySet1(_1415_v88, 17)
				(_nw210).ArraySet1(Companion_D0_.Create_DC1_(p1, _1412_cf17, _1285_v2), 18)
				(_nw210).ArraySet1(_1415_v88, 19)
				(_nw210).ArraySet1(_1415_v88, 20)
				(_nw210).ArraySet1(_1415_v88, 21)
				(_nw210).ArraySet1(_1415_v88, 22)
				(_nw210).ArraySet1(_1415_v88, 23)
				(_nw210).ArraySet1(_1415_v88, 24)
				(_nw210).ArraySet1(_1415_v88, 25)
				_1417_v90 = _nw210
				_1417_v90 = _1417_v90
				var _1420_v91 _dafny.Array
				_ = _1420_v91
				var _nwElement0_40 bool = true
				_ = _nwElement0_40
				var _nw211 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(8))
				_ = _nw211
				(_nw211).ArraySet1(_nwElement0_40, 0)
				(_nw211).ArraySet1(!(p1), 1)
				(_nw211).ArraySet1(p0, 2)
				(_nw211).ArraySet1(p0, 3)
				(_nw211).ArraySet1(!(false), 4)
				(_nw211).ArraySet1(p1, 5)
				(_nw211).ArraySet1(p1, 6)
				(_nw211).ArraySet1(p1, 7)
				_1420_v91 = _nw211
				var _1421_v92 *C0
				_ = _1421_v92
				var _nw212 *C0 = New_C0_()
				_ = _nw212
				_nw212.Ctor__(_1420_v91, true)
				_1421_v92 = _nw212
				(globalState).F2 = _dafny.Companion_Sequence_.Concatenate(_1283_v0, _dafny.Companion_Sequence_.Concatenate(_1283_v0, _1283_v0))
			} else if _source18.Is_DC8() {
				var _1422___mcc_h9 bool = _source18.Get_().(D2_DC8).Cf19
				_ = _1422___mcc_h9
				var _1423___mcc_h10 bool = _source18.Get_().(D2_DC8).Cf20
				_ = _1423___mcc_h10
				var _1424_cf20 bool = _1423___mcc_h10
				_ = _1424_cf20
				var _1425_cf19 bool = _1422___mcc_h9
				_ = _1425_cf19
				_1285_v2 = _dafny.CodePoint('l')
				(globalState).F5 = (_this).F15()
				var _1426_v93 _dafny.Sequence
				_ = _1426_v93
				_1426_v93 = _dafny.SeqOf(_dafny.SetOf((_this).F6(), _dafny.IntOfUint32((_1333_v35).Cardinality())))
				var _1427_v94 _dafny.Map
				_ = _1427_v94
				_1427_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1283_v0, _1286_v3)
				var _pat_let_tv25 = _1427_v94
				_ = _pat_let_tv25
				var _pat_let_tv26 = _1283_v0
				_ = _pat_let_tv26
				var _pat_let_tv27 = _1427_v94
				_ = _pat_let_tv27
				var _pat_let_tv28 = _1283_v0
				_ = _pat_let_tv28
				var _pat_let_tv29 = _1286_v3
				_ = _pat_let_tv29
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_1289_v6), 0))
				_ = _index200
				var _rhs229 _dafny.Set = (_1426_v93).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1426_v93).Cardinality()))).Uint32()).(_dafny.Set)
				_ = _rhs229
				var _rhs230 D9 = func(_pat_let33_0 D9) D9 {
					return func(_1428_dt__update__tmp_h4 D9) D9 {
						return func(_pat_let34_0 D9) D9 {
							return func(_1429_dt__update_hcf51_h0 D9) D9 {
								return Companion_D9_.Create_DC27_(_1429_dt__update_hcf51_h0)
							}(_pat_let34_0)
						}((func() D9 {
							if (_pat_let_tv25).Contains(_pat_let_tv26) {
								return (_pat_let_tv27).Get(_pat_let_tv28).(D9)
							}
							return Companion_D9_.Create_DC27_(_pat_let_tv29)
						})())
					}(_pat_let33_0)
				}(_1288_v5)
				_ = _rhs230
				var _rhs231 bool = p0
				_ = _rhs231
				var _lhs184 _dafny.Array = _1289_v6
				_ = _lhs184
				var _lhs185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_1289_v6), 0))
				_ = _lhs185
				var _lhs186 *GlobalState = globalState
				_ = _lhs186
				_1394_v80 = _rhs229
				(_lhs184).ArraySet1(_rhs230, (_lhs185).Int())
				_lhs186.F0 = _rhs231
				var _1430_v95 D9
				_ = _1430_v95
				_1430_v95 = Companion_D9_.Create_DC25_((_this).F15(), (_this).F15(), (_this).F6(), _1285_v2)
				var _1431_v96 _dafny.Map
				_ = _1431_v96
				_1431_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1402_v87).Dtor_cf14(), _1430_v95)
				_1431_v96 = (_1431_v96).Update(_1401_v86, _1430_v95)
			} else {
				var _1432___mcc_h11 _dafny.Map = _source18.Get_().(D2_DC5).Cf12
				_ = _1432___mcc_h11
				var _1433_cf12 _dafny.Map = _1432___mcc_h11
				_ = _1433_cf12
				var _1434_v97 _dafny.Set
				_ = _1434_v97
				_1434_v97 = _dafny.SetOf(_1433_cf12)
				_1434_v97 = _1434_v97
				var _1435_v98 *C1
				_ = _1435_v98
				var _nw213 *C1 = New_C1_()
				_ = _nw213
				_nw213.Ctor__(((_this).F6()).Minus((_this).F15()))
				_1435_v98 = _nw213
				var _rhs232 bool = p0
				_ = _rhs232
				var _rhs233 bool = true
				_ = _rhs233
				var _lhs187 *GlobalState = globalState
				_ = _lhs187
				r2 = _rhs232
				_lhs187.F0 = _rhs233
				_1395_v81 = _1401_v86
			}
		}
		r0 = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1283_v0).Cardinality()), _dafny.IntOfInt64(352))).Times((_this).F15())
		r1 = _1284_v1
		r2 = !(p0) || (((_this).F6()).Cmp(_dafny.IntOfUint32((_dafny.SeqOf((_this).F6(), _dafny.IntOfInt64(895), (_this).F6(), (_this).F6(), (_this).F6())).Cardinality())) > 0)
		var _1436_v99 _dafny.Array
		_ = _1436_v99
		var _nw214 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(19))
		_ = _nw214
		_1436_v99 = _nw214
		var _1437_v100 D2
		_ = _1437_v100
		_1437_v100 = Companion_D2_.Create_DC7_(_1332_v34, _1436_v99, _dafny.IntOfUint32((_1283_v0).Cardinality()), (_dafny.Zero).Minus((_this).F6()))
		r3 = (_1437_v100).Dtor_cf18()
		return r0, r1, r2, r3
	}
}
func (_this *C5) F15() _dafny.Int {
	{
		return _this._f15
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f13 _dafny.Int
	_f14 bool
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f13 = _dafny.Zero
	_this._f14 = false
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

func (_this *C6) Ctor__(f13 _dafny.Int, f14 bool) {
	{
		(_this)._f13 = f13
		(_this)._f14 = f14
	}
}
func (_this *C6) Fm15(p0 _dafny.Int, p1 D2, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(903)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-992))).Uint32(), func(coer92 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg92 _dafny.Int) interface{} {
				return coer92(arg92)
			}
		}(func(_1438_i0 _dafny.Int) _dafny.Int {
			return (_this).F13()
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer93 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg93 _dafny.Int) interface{} {
				return coer93(arg93)
			}
		}(func(_1439_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(741)
		}))))
	}
}
func (_this *C6) Fm16(p0 D2, p1 D0, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F13()).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F14()))).Cardinality()))
	}
}
func (_this *C6) M8(globalState *GlobalState) (_dafny.Int, D2, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 D2 = Companion_D2_.Default()
		_ = r1
		var r2 bool = false
		_ = r2
		var _1440_v0 _dafny.Map
		_ = _1440_v0
		_1440_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13())
		var _1441_v1 _dafny.Map
		_ = _1441_v1
		_1441_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_1440_v0).Contains((_this).F13()) {
				return (_1440_v0).Get((_this).F13()).(_dafny.Int)
			}
			return _dafny.IntOfInt64(956)
		})(), (_this).F14())
		_1441_v1 = (_1441_v1).Update((_this).F13(), (_this).F14())
		(globalState).F0 = Companion_Default___.Fm3(globalState)
		var _1442_i0 _dafny.Int
		_ = _1442_i0
		_1442_i0 = _dafny.Zero
		{
			for (_this).F14() {
				{
					if (_1442_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1442_i0 = (_1442_i0).Plus(_dafny.One)
					var _1443_v2 _dafny.Array
					_ = _1443_v2
					var _nw215 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
					_ = _nw215
					_1443_v2 = _nw215
					var _1444_v3 _dafny.MultiSet
					_ = _1444_v3
					_1444_v3 = _dafny.MultiSetOf((_this).F13())
					var _1445_v4 _dafny.Sequence
					_ = _1445_v4
					_1445_v4 = _dafny.SeqOf((_1444_v3).Cardinality(), (_this).F13())
					var _1446_v5 _dafny.Map
					_ = _1446_v5
					_1446_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1443_v2, _1445_v4)
					_1446_v5 = (_1446_v5).Update(_1443_v2, _1445_v4)
					var _1447_v6 _dafny.Map
					_ = _1447_v6
					_1447_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-212)))
					_1447_v6 = (_1447_v6).Update((_this).F14(), (_this).F13())
					r2 = (_this).F14()
					if (_this).F14() {
						(globalState).F0 = !(false)
						(globalState).F5 = (_this).F13()
						var _1448_v7 _dafny.Sequence
						_ = _1448_v7
						_1448_v7 = _dafny.UnicodeSeqOfUtf8Bytes("cfytvh")
						var _1449_v8 _dafny.Sequence
						_ = _1449_v8
						_1449_v8 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("cevep"), (func() _dafny.Sequence {
							if (_this).F14() {
								return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(447))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg94 _dafny.Int) interface{} {
										return coer94(arg94)
									}
								}(func(_1450_i1 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('h')
								}))
							}
							return _dafny.UnicodeSeqOfUtf8Bytes("o")
						})(), _dafny.Companion_Sequence_.Concatenate(_1448_v7, _1448_v7))
						(globalState).F2 = (_1449_v8).Select((Companion_Default___.SafeIndex(((_dafny.Zero).Minus((_this).F13())).Plus(_dafny.IntOfInt64(97)), _dafny.IntOfUint32((_1449_v8).Cardinality()))).Uint32()).(_dafny.Sequence)
						var _1451_v9 _dafny.CodePoint
						_ = _1451_v9
						_1451_v9 = _dafny.CodePoint('e')
						_1451_v9 = _1451_v9
						(globalState).F5 = (_this).F13()
					} else {
						var _1452_v10 _dafny.Set
						_ = _1452_v10
						_1452_v10 = _dafny.SetOf(_1443_v2, _1443_v2, _1443_v2, _1443_v2, _1443_v2)
						_1452_v10 = _1452_v10
						var _1453_v11 _dafny.Sequence
						_ = _1453_v11
						_1453_v11 = _dafny.SeqOf((_this).F14(), (_this).F14())
						var _1454_v12 _dafny.Sequence
						_ = _1454_v12
						_1454_v12 = _dafny.UnicodeSeqOfUtf8Bytes("snuxhfgl")
						var _1455_v13 _dafny.Set
						_ = _1455_v13
						_1455_v13 = _dafny.SetOf(_1454_v12)
						var _1456_v14 D0
						_ = _1456_v14
						_1456_v14 = Companion_D0_.Create_DC0_(_1453_v11, (_1455_v13).Cardinality())
						var _1457_v15 D1
						_ = _1457_v15
						_1457_v15 = Companion_D1_.Create_DC4_(_1456_v14, _dafny.IntOfInt64(923), Companion_Default___.Fm17((_this).F13(), (_this).F13(), globalState), (_this).F14(), (_this).F13())
						(globalState).F2 = (_1457_v15).Dtor_cf9()
						var _1458_v16 _dafny.Array
						_ = _1458_v16
						var _nw216 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
						_ = _nw216
						_1458_v16 = _nw216
						var _1459_v17 _dafny.MultiSet
						_ = _1459_v17
						_1459_v17 = _dafny.MultiSetOf(_1458_v16, _1458_v16, _1458_v16, _1458_v16)
						var _1460_v18 _dafny.Sequence
						_ = _1460_v18
						_1460_v18 = _dafny.SeqOf(_1445_v4)
						var _1461_v19 _dafny.CodePoint
						_ = _1461_v19
						_1461_v19 = _dafny.CodePoint('f')
						var _1462_v20 _dafny.Map
						_ = _1462_v20
						_1462_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1461_v19, _1458_v16)
						var _1463_v21 _dafny.Sequence
						_ = _1463_v21
						_1463_v21 = _dafny.SeqOf((_dafny.MultiSetOf(_1458_v16)).Update((func() _dafny.Array {
							if (_1462_v20).Contains(_dafny.CodePoint('e')) {
								return (_1462_v20).Get(_dafny.CodePoint('e')).(_dafny.Array)
							}
							return _1458_v16
						})(), Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.SeqOf(func(_pat_let35_0 D0) D0 {
							return func(_1464_dt__update__tmp_h0 D0) D0 {
								return func(_pat_let36_0 _dafny.Int) D0 {
									return func(_1465_dt__update_hcf1_h0 _dafny.Int) D0 {
										return Companion_D0_.Create_DC0_((_1464_dt__update__tmp_h0).Dtor_cf0(), _1465_dt__update_hcf1_h0)
									}(_pat_let36_0)
								}((_this).F13())
							}(_pat_let35_0)
						}(_1456_v14), _1456_v14, _1456_v14, _1456_v14)).Cardinality()))), _1459_v17)
						_1459_v17 = ((((_1459_v17).Update(_1458_v16, Companion_Default___.Abs(_dafny.IntOfInt64(774)))).Update(_1458_v16, Companion_Default___.Abs(_dafny.IntOfUint32(((_1460_v18).Select((Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_1460_v18).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())))).Union(_1459_v17)).Union((_1459_v17).Difference((_1463_v21).Select((Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_1463_v21).Cardinality()))).Uint32()).(_dafny.MultiSet)))
						var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen((_1458_v16), 0))
						_ = _index201
						(_1458_v16).ArraySet1((_1444_v3).Cardinality(), (_index201).Int())
						var _1466_v22 _dafny.Set
						_ = _1466_v22
						_1466_v22 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("vp"), (Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vp")).Cardinality()))).Uint32(), _dafny.CodePoint('q'))).Cardinality()))
						var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen((_1458_v16), 0))
						_ = _index202
						var _rhs234 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F13(), (_1466_v22).Cardinality())
						_ = _rhs234
						var _rhs235 _dafny.Int = (_dafny.Zero).Minus(((_this).F13()).Times((func() _dafny.Int {
							if (_1444_v3).Contains((_this).F13()) {
								return (_1444_v3).Multiplicity((_this).F13())
							}
							return (_this).F13()
						})()))
						_ = _rhs235
						var _lhs188 _dafny.Array = _1458_v16
						_ = _lhs188
						var _lhs189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen((_1458_v16), 0))
						_ = _lhs189
						var _lhs190 *GlobalState = globalState
						_ = _lhs190
						(_lhs188).ArraySet1(_rhs234, (_lhs189).Int())
						_lhs190.F5 = _rhs235
						var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen((_1458_v16), 0))
						_ = _index203
						(_1458_v16).ArraySet1((_this).F13(), (_index203).Int())
					}
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _hi10 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F14())).Cardinality()
		_ = _hi10
		for _1467_i2 := (_this).F13(); _1467_i2.Cmp(_hi10) < 0; _1467_i2 = _1467_i2.Plus(_dafny.One) {
			var _1468_v23 _dafny.Map
			_ = _1468_v23
			_1468_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F14()), (_this).F13())
			var _1469_v24 D2
			_ = _1469_v24
			_1469_v24 = Companion_D2_.Create_DC5_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_1468_v23).Contains((_this).F14()) {
					return (_1468_v23).Get((_this).F14()).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(_1467_i2)
			})(), (_dafny.Zero).Minus(_1467_i2)))
			var _source19 D2 = _1469_v24
			_ = _source19
			if _source19.Is_DC6() {
				var _1470___mcc_h0 _dafny.Int = _source19.Get_().(D2_DC6).Cf13
				_ = _1470___mcc_h0
				var _1471___mcc_h1 _dafny.Array = _source19.Get_().(D2_DC6).Cf14
				_ = _1471___mcc_h1
				var _1472_cf14 _dafny.Array = _1471___mcc_h1
				_ = _1472_cf14
				var _1473_cf13 _dafny.Int = _1470___mcc_h0
				_ = _1473_cf13
				var _1474_v25 _dafny.Map
				_ = _1474_v25
				_1474_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F14())
				_1474_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), Companion_Default___.Fm3(globalState))
				var _1475_v26 _dafny.Sequence
				_ = _1475_v26
				_1475_v26 = _dafny.SeqOf(_1468_v23, _1468_v23, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F14()), _1473_cf13))
				var _1476_v27 _dafny.Sequence
				_ = _1476_v27
				_1476_v27 = _dafny.UnicodeSeqOfUtf8Bytes("qk")
				var _1477_v28 _dafny.Sequence
				_ = _1477_v28
				_1477_v28 = _dafny.SeqOf((_this).F14())
				r2 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_1475_v26, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1476_v27).Cardinality()), _dafny.IntOfUint32((_1475_v26).Cardinality()))).Uint32(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _dafny.IntOfUint32((_1477_v28).Cardinality()))).Update((_this).F14(), _dafny.IntOfInt64(83))).Update((_this).F14(), _1473_cf13)), _1475_v26)
				_1474_v25 = (_1474_v25).Update(true, (func() bool {
					if (_1474_v25).Contains((_this).F14()) {
						return (_1474_v25).Get((_this).F14()).(bool)
					}
					return !((_this).F14())
				})())
				(globalState).F0 = (_this).F14()
			} else if _source19.Is_DC7() {
				var _1478___mcc_h2 _dafny.Sequence = _source19.Get_().(D2_DC7).Cf15
				_ = _1478___mcc_h2
				var _1479___mcc_h3 _dafny.Array = _source19.Get_().(D2_DC7).Cf16
				_ = _1479___mcc_h3
				var _1480___mcc_h4 _dafny.Int = _source19.Get_().(D2_DC7).Cf17
				_ = _1480___mcc_h4
				var _1481___mcc_h5 _dafny.Int = _source19.Get_().(D2_DC7).Cf18
				_ = _1481___mcc_h5
				var _1482_cf18 _dafny.Int = _1481___mcc_h5
				_ = _1482_cf18
				var _1483_cf17 _dafny.Int = _1480___mcc_h4
				_ = _1483_cf17
				var _1484_cf16 _dafny.Array = _1479___mcc_h3
				_ = _1484_cf16
				var _1485_cf15 _dafny.Sequence = _1478___mcc_h2
				_ = _1485_cf15
				var _1486_v29 *C1
				_ = _1486_v29
				var _nw217 *C1 = New_C1_()
				_ = _nw217
				_nw217.Ctor__((_this).F13())
				_1486_v29 = _nw217
				var _1487_v30 _dafny.Sequence
				_ = _1487_v30
				_1487_v30 = _dafny.SeqOf(_1485_cf15, _1485_cf15, _1485_cf15, _1485_cf15, _1485_cf15)
				_1487_v30 = _1487_v30
				(globalState).F5 = (_this).F13()
				(globalState).F5 = (_this).F13()
			} else if _source19.Is_DC8() {
				var _1488___mcc_h6 bool = _source19.Get_().(D2_DC8).Cf19
				_ = _1488___mcc_h6
				var _1489___mcc_h7 bool = _source19.Get_().(D2_DC8).Cf20
				_ = _1489___mcc_h7
				var _1490_cf20 bool = _1489___mcc_h7
				_ = _1490_cf20
				var _1491_cf19 bool = _1488___mcc_h6
				_ = _1491_cf19
				(globalState).F5 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_this).F13(), (_this).F13()), (_this).F13())
				var _1492_v31 _dafny.Array
				_ = _1492_v31
				var _nw218 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
				_ = _nw218
				_1492_v31 = _nw218
				var _1493_v32 _dafny.Array
				_ = _1493_v32
				var _nw219 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw219
				_1493_v32 = _nw219
				var _1494_v33 _dafny.MultiSet
				_ = _1494_v33
				_1494_v33 = _dafny.MultiSetOf(_1493_v32, _1493_v32)
				var _1495_v34 _dafny.CodePoint
				_ = _1495_v34
				_1495_v34 = _dafny.CodePoint('g')
				var _1496_v35 _dafny.MultiSet
				_ = _1496_v35
				_1496_v35 = _dafny.MultiSetOf(_1495_v34)
				var _1497_v36 _dafny.Sequence
				_ = _1497_v36
				_1497_v36 = _dafny.SeqOf((_this).F14(), Companion_Default___.Fm3(globalState))
				var _1498_v37 _dafny.Array
				_ = _1498_v37
				var _nwElement0_41 _dafny.Map = _1468_v23
				_ = _nwElement0_41
				var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(9))
				_ = _nw220
				(_nw220).ArraySet1(_nwElement0_41, 0)
				(_nw220).ArraySet1(_1468_v23, 1)
				(_nw220).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1491_cf19, _1467_i2), 2)
				(_nw220).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1490_cf20, (_1494_v33).Cardinality()), 3)
				(_nw220).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_1496_v35).Cardinality()), 4)
				(_nw220).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1491_cf19, (_this).F13())).Update(_1491_cf19, (_dafny.Zero).Minus(_1467_i2)), 5)
				(_nw220).ArraySet1(_1468_v23, 6)
				(_nw220).ArraySet1(Companion_Default___.Fm31(_1491_cf19, (_this).F13(), _1467_i2, _1497_v36, globalState), 7)
				(_nw220).ArraySet1(_1468_v23, 8)
				_1498_v37 = _nw220
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1492_v31), 0))
				_ = _index204
				(_1492_v31).ArraySet1(_1498_v37, (_index204).Int())
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_1492_v31), 0))
				_ = _index205
				(_1492_v31).ArraySet1(_1498_v37, (_index205).Int())
				(globalState).F5 = (_dafny.Zero).Minus(((_this).F13()).Minus(((_this).F13()).Minus((_this).F13())))
				var _1499_v38 _dafny.Array
				_ = _1499_v38
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_36
				var _nw221 _dafny.Array
				_ = _nw221
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw221 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) bool = func(_1500_i3 _dafny.Int) bool {
						return _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("tsfvbtav"), _dafny.CodePoint('c'))
					}
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw221 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw221).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw221).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1499_v38 = _nw221
				var _1501_v39 _dafny.Sequence
				_ = _1501_v39
				_1501_v39 = _dafny.SeqOf(_1499_v38)
				_1499_v38 = (_1501_v39).Select((Companion_Default___.SafeIndex((_dafny.IntOfInt64(-871)).Minus(_dafny.IntOfUint32((Companion_Default___.Fm4(_1467_i2, (_this).F13(), (_this).F13(), globalState)).Cardinality())), _dafny.IntOfUint32((_1501_v39).Cardinality()))).Uint32()).(_dafny.Array)
			} else {
				var _1502___mcc_h8 _dafny.Map = _source19.Get_().(D2_DC5).Cf12
				_ = _1502___mcc_h8
				var _1503_cf12 _dafny.Map = _1502___mcc_h8
				_ = _1503_cf12
				var _1504_v40 _dafny.Array
				_ = _1504_v40
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_37
				var _nw222 _dafny.Array
				_ = _nw222
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw222 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) _dafny.Int = (func(_1505_i2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1506_i4 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1506_i4, (_dafny.Zero).Minus(_1505_i2))
						}
					})(_1467_i2)
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw222 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw222).ArraySet1(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw222).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1504_v40 = _nw222
				var _1507_v41 _dafny.MultiSet
				_ = _1507_v41
				_1507_v41 = _dafny.MultiSetOf(_1467_i2)
				var _1508_v42 _dafny.MultiSet
				_ = _1508_v42
				_1508_v42 = _dafny.MultiSetOf(_1507_v41)
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1504_v40), 0))
				_ = _index206
				(_1504_v40).ArraySet1(Companion_Default___.SafeModuloInt((_1508_v42).Cardinality(), _1467_i2), (_index206).Int())
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1504_v40), 0))
				_ = _index207
				(_1504_v40).ArraySet1((_this).F13(), (_index207).Int())
				var _1509_v43 _dafny.Sequence
				_ = _1509_v43
				_1509_v43 = _dafny.UnicodeSeqOfUtf8Bytes("ecuurfubj")
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1504_v40), 0))
				_ = _index208
				(_1504_v40).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gyxgc"), _1509_v43)).Cardinality())), (_index208).Int())
				var _1510_v44 _dafny.CodePoint
				_ = _1510_v44
				_1510_v44 = _dafny.CodePoint('t')
				var _1511_v45 _dafny.Set
				_ = _1511_v45
				_1511_v45 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_1509_v43, (Companion_Default___.SafeIndex((_1504_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1504_v40), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1509_v43).Cardinality()))).Uint32(), _1510_v44), _1509_v43, _1509_v43)
				var _1512_v46 _dafny.Sequence
				_ = _1512_v46
				_1512_v46 = _dafny.SeqOf((_1511_v45).Intersection(_1511_v45))
				_1511_v45 = (_1512_v46).Select((Companion_Default___.SafeIndex(_1467_i2, _dafny.IntOfUint32((_1512_v46).Cardinality()))).Uint32()).(_dafny.Set)
				var _1513_v47 _dafny.Array
				_ = _1513_v47
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_38
				var _nw223 _dafny.Array
				_ = _nw223
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw223 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) bool = func(_1514_i5 _dafny.Int) bool {
						return (_this).F14()
					}
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw223 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw223).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw223).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1513_v47 = _nw223
				var _1515_v48 _dafny.Map
				_ = _1515_v48
				_1515_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _1509_v43)
				var _1516_v49 *C3
				_ = _1516_v49
				var _nw224 *C3 = New_C3_()
				_ = _nw224
				_nw224.Ctor__(_1440_v0, _1513_v47, _1515_v48, (_this).F13())
				_1516_v49 = _nw224
				var _1517_v50 _dafny.Sequence
				_ = _1517_v50
				_1517_v50 = _dafny.SeqOf(_1516_v49)
				var _1518_v51 _dafny.Set
				_ = _1518_v51
				_1518_v51 = _dafny.SetOf((_1517_v50).Select((Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_1517_v50).Cardinality()))).Uint32()).(*C3))
				var _rhs236 bool = (_this).F14()
				_ = _rhs236
				var _rhs237 _dafny.Set = (_1518_v51).Union((_dafny.SetOf(_1516_v49, _1516_v49)).Difference(_1518_v51))
				_ = _rhs237
				var _lhs191 *GlobalState = globalState
				_ = _lhs191
				_lhs191.F0 = _rhs236
				_1518_v51 = _rhs237
			}
			var _1519_v52 _dafny.Sequence
			_ = _1519_v52
			_1519_v52 = _dafny.SeqOf(_1467_i2, _1467_i2)
			var _1520_v53 _dafny.Set
			_ = _1520_v53
			_1520_v53 = _dafny.SetOf(_dafny.Companion_Sequence_.IsProperPrefixOf(_1519_v52, _1519_v52), ((_this).F13()).Cmp((_dafny.Zero).Minus((_this).F13())) >= 0, (_this).F14(), (_this).F14(), (_this).F14())
			_1520_v53 = (_1520_v53).Intersection(_1520_v53)
			var _1521_v54 _dafny.Map
			_ = _1521_v54
			_1521_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), !((_this).F14()))
			var _1522_v55 D4
			_ = _1522_v55
			_1522_v55 = Companion_D4_.Create_DC14_((_this).F14(), _1521_v54, (_this).F13(), _1521_v54)
			var _1523_v56 _dafny.Sequence
			_ = _1523_v56
			_1523_v56 = _dafny.SeqOf((_1522_v55).Dtor_cf28(), (_this).F14())
			var _1524_v57 _dafny.Map
			_ = _1524_v57
			_1524_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1467_i2, _1523_v56)
			var _1525_v58 _dafny.Sequence
			_ = _1525_v58
			_1525_v58 = _dafny.UnicodeSeqOfUtf8Bytes("wdgqi")
			var _1526_v59 _dafny.Map
			_ = _1526_v59
			_1526_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_this).F14(), (_this).F14()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1467_i2, _1467_i2)).Update(_dafny.IntOfUint32((_1525_v58).Cardinality()), _1467_i2))
			_1524_v57 = (_1524_v57).Update((_1526_v59).Cardinality(), _1523_v56)
			(globalState).F5 = _1467_i2
		}
		var _1527_v60 _dafny.Array
		_ = _1527_v60
		var _nw225 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw225
		_1527_v60 = _nw225
		var _1528_v61 D2
		_ = _1528_v61
		_1528_v61 = Companion_D2_.Create_DC6_((_this).F13(), _1527_v60)
		_1528_v61 = _1528_v61
		var _1529_v62 D2
		_ = _1529_v62
		_1529_v62 = Companion_D2_.Create_DC8_(false, (_this).F14())
		var _1530_v63 _dafny.CodePoint
		_ = _1530_v63
		_1530_v63 = _dafny.CodePoint('h')
		var _1531_v64 D0
		_ = _1531_v64
		_1531_v64 = Companion_D0_.Create_DC1_((_this).F14(), (_this).F13(), _1530_v63)
		var _hi11 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F13(), (_this).Fm16(Companion_Default___.Fm49((_this).F13(), (_this).Fm16(_1529_v62, _1531_v64, globalState), (_this).F14(), (_this).F13(), globalState), _1531_v64, globalState))
		_ = _hi11
		for _1532_i6 := (_this).F13(); _1532_i6.Cmp(_hi11) < 0; _1532_i6 = _1532_i6.Plus(_dafny.One) {
			(globalState).F0 = (_this).F14()
			_1440_v0 = (_1440_v0).Update((_this).F13(), (_dafny.IntOfInt64(-670)).Plus((_this).F13()))
			var _1533_v65 _dafny.MultiSet
			_ = _1533_v65
			_1533_v65 = _dafny.MultiSetOf((_this).F14(), false)
			r0 = ((_1533_v65).Cardinality()).Times(Companion_Default___.Fm2((_dafny.SetOf(_1532_i6)).Cardinality(), globalState))
			var _1534_v66 _dafny.Sequence
			_ = _1534_v66
			_1534_v66 = _dafny.SeqOf((_this).F13())
			(globalState).F0 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1534_v66, _1534_v66)).Cardinality())).Cmp(_dafny.IntOfInt64(-581)) <= 0
		}
		r0 = (_this).F13()
		var _1535_v67 D2
		_ = _1535_v67
		_1535_v67 = Companion_D2_.Create_DC5_(_1440_v0)
		r1 = _1535_v67
		r2 = (_this).F14()
		return r0, r1, r2
	}
}
func (_this *C6) M9(p0 _dafny.Sequence, p1 _dafny.Sequence, globalState *GlobalState) (_dafny.Set, _dafny.Int, _dafny.MultiSet) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r2
		var _1536_v0 _dafny.Map
		_ = _1536_v0
		_1536_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F13())
		var _1537_v1 _dafny.MultiSet
		_ = _1537_v1
		_1537_v1 = _dafny.MultiSetOf((_this).F13())
		var _1538_v2 _dafny.MultiSet
		_ = _1538_v2
		_1538_v2 = _dafny.MultiSetOf(_1537_v1)
		var _1539_v3 _dafny.Array
		_ = _1539_v3
		var _len0_39 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_39
		var _nw226 _dafny.Array
		_ = _nw226
		if _len0_39.Cmp(_dafny.Zero) == 0 {
			_nw226 = _dafny.NewArray(_len0_39)
		} else {
			var _init39 func(_dafny.Int) bool = func(_1540_i0 _dafny.Int) bool {
				return (_this).F14()
			}
			_ = _init39
			var _element0_39 = _init39(_dafny.Zero)
			_ = _element0_39
			_nw226 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
			(_nw226).ArraySet1(_element0_39, 0)
			var _nativeLen0_39 = (_len0_39).Int()
			_ = _nativeLen0_39
			for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
				(_nw226).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
			}
		}
		_1539_v3 = _nw226
		var _1541_v4 _dafny.Map
		_ = _1541_v4
		_1541_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2((_this).F13(), globalState), p0)
		var _1542_v5 *C2
		_ = _1542_v5
		var _nw227 *C2 = New_C2_()
		_ = _nw227
		_nw227.Ctor__((_this).F13(), _1538_v2, _1539_v3, _1541_v4)
		_1542_v5 = _nw227
		var _1543_v6 _dafny.Sequence
		_ = _1543_v6
		_1543_v6 = _dafny.SeqOf((_this).F14(), (_this).F14())
		var _1544_v7 _dafny.Sequence
		_ = _1544_v7
		_1544_v7 = _dafny.SeqOf(_dafny.IntOfUint32((_1543_v6).Cardinality()))
		var _1545_v8 _dafny.Map
		_ = _1545_v8
		_1545_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1542_v5, _1544_v7)
		_1536_v0 = (_1536_v0).Update(_dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
			if (_1545_v8).Contains(_1542_v5) {
				return (_1545_v8).Get(_1542_v5).(_dafny.Sequence)
			}
			return _1544_v7
		})(), (_this).F13()), (_this).F13())
		var _1546_i1 _dafny.Int
		_ = _1546_i1
		_1546_i1 = _dafny.Zero
		{
			for (_this).F14() {
				{
					if (_1546_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1546_i1 = (_1546_i1).Plus(_dafny.One)
					r1 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F13()), (_this).F13())
					var _1547_v9 _dafny.Set
					_ = _1547_v9
					_1547_v9 = _dafny.SetOf(_1539_v3, _1539_v3, _1539_v3, _1539_v3)
					_1547_v9 = _1547_v9
					(globalState).F5 = (_this).F13()
					var _1548_v11 _dafny.Map
					_ = _1548_v11
					_1548_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F14())
					var _1549_v12 *C2
					_ = _1549_v12
					var _nw228 *C2 = New_C2_()
					_ = _nw228
					_nw228.Ctor__((_this).F13(), _1538_v2, _1539_v3, (func() _dafny.Map {
						var _coll41 = _dafny.NewMapBuilder()
						_ = _coll41
						for _iter43 := _dafny.Iterate((_1548_v11).Keys().Elements()); ; {
							_compr_41, _ok43 := _iter43()
							if !_ok43 {
								break
							}
							var _1550_v10 _dafny.Int
							_1550_v10 = interface{}(_compr_41).(_dafny.Int)
							if (_1548_v11).Contains(_1550_v10) {
								_coll41.Add((_1550_v10).Plus((_this).F13()), p1)
							}
						}
						return _coll41.ToMap()
					}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-554), p1)))
					_1549_v12 = _nw228
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _1551_v13 D2
		_ = _1551_v13
		_1551_v13 = Companion_D2_.Create_DC8_((_this).F14(), (_this).F14())
		var _1552_v14 _dafny.CodePoint
		_ = _1552_v14
		_1552_v14 = _dafny.CodePoint('d')
		var _1553_v15 _dafny.Map
		_ = _1553_v15
		_1553_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _dafny.SetOf((_this).F13()))
		var _1554_v16 _dafny.Set
		_ = _1554_v16
		_1554_v16 = _dafny.SetOf((_1536_v0).Cardinality(), (_this).F13(), (_this).F13())
		var _rhs238 _dafny.Int = ((func() _dafny.Int {
			if false {
				return _dafny.IntOfInt64(449)
			}
			return (_this).F13()
		})()).Plus(((_this).F13()).Plus((_this).Fm16(func(_pat_let37_0 D2) D2 {
			return func(_1555_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let38_0 bool) D2 {
					return func(_1556_dt__update_hcf19_h0 bool) D2 {
						return Companion_D2_.Create_DC8_(_1556_dt__update_hcf19_h0, (_1555_dt__update__tmp_h0).Dtor_cf20())
					}(_pat_let38_0)
				}((_this).F14())
			}(_pat_let37_0)
		}(_1551_v13), Companion_D0_.Create_DC1_(true, (_this).F13(), _1552_v14), globalState)))
		_ = _rhs238
		var _rhs239 _dafny.Array = _1539_v3
		_ = _rhs239
		var _rhs240 _dafny.Sequence = _1543_v6
		_ = _rhs240
		var _rhs241 bool = ((_this).F13()).Cmp((_dafny.Zero).Minus((_this).F13())) < 0
		_ = _rhs241
		var _rhs242 bool = (_1554_v16).IsSubsetOf((func() _dafny.Set {
			if (_1553_v15).Contains((_this).F13()) {
				return (_1553_v15).Get((_this).F13()).(_dafny.Set)
			}
			return _1554_v16
		})())
		_ = _rhs242
		var _lhs192 *GlobalState = globalState
		_ = _lhs192
		var _lhs193 *GlobalState = globalState
		_ = _lhs193
		r1 = _rhs238
		_1539_v3 = _rhs239
		_1543_v6 = _rhs240
		_lhs192.F0 = _rhs241
		_lhs193.F0 = _rhs242
		if (func() bool {
			if ((_this).F13()).Cmp((_this).F13()) == 0 {
				return (_this).F14()
			}
			return _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("q"), _1552_v14)
		})() {
			var _1557_v17 _dafny.Set
			_ = _1557_v17
			_1557_v17 = _dafny.SetOf((_this).F14(), (_this).F14(), (_this).F14())
			(globalState).F0 = (_1557_v17).IsDisjointFrom(_1557_v17)
			var _1558_v18 _dafny.Map
			_ = _1558_v18
			_1558_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13())
			var _1559_v19 D2
			_ = _1559_v19
			_1559_v19 = Companion_D2_.Create_DC5_(_1558_v18)
			var _rhs243 D2 = Companion_D2_.Create_DC5_(_1558_v18)
			_ = _rhs243
			var _rhs244 bool = (_1543_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-56), _dafny.IntOfUint32((_1543_v6).Cardinality()))).Uint32()).(bool)
			_ = _rhs244
			var _rhs245 _dafny.Int = ((_this).F13()).Minus((_1544_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.IntOfUint32((_1544_v7).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs245
			var _lhs194 *GlobalState = globalState
			_ = _lhs194
			var _lhs195 *GlobalState = globalState
			_ = _lhs195
			_1559_v19 = _rhs243
			_lhs194.F0 = _rhs244
			_lhs195.F5 = _rhs245
			r1 = ((_this).F13()).Plus((_this).F13())
			var _1560_v20 _dafny.Map
			_ = _1560_v20
			_1560_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1543_v6).Select((Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_1543_v6).Cardinality()))).Uint32()).(bool), false)
			(globalState).F0 = (func() bool {
				if (_1560_v20).Contains(((_this).F14()) || (false)) {
					return (_1560_v20).Get(((_this).F14()) || (false)).(bool)
				}
				return !((_1554_v16).IsDisjointFrom(_1554_v16))
			})()
			var _1561_v21 _dafny.Array
			_ = _1561_v21
			var _len0_40 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_40
			var _nw229 _dafny.Array
			_ = _nw229
			if _len0_40.Cmp(_dafny.Zero) == 0 {
				_nw229 = _dafny.NewArray(_len0_40)
			} else {
				var _init40 func(_dafny.Int) _dafny.Sequence = func(_1562_i2 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("jfhhroih")
				}
				_ = _init40
				var _element0_40 = _init40(_dafny.Zero)
				_ = _element0_40
				_nw229 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
				(_nw229).ArraySet1(_element0_40, 0)
				var _nativeLen0_40 = (_len0_40).Int()
				_ = _nativeLen0_40
				for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
					(_nw229).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
				}
			}
			_1561_v21 = _nw229
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_1561_v21), 0))
			_ = _index209
			(_1561_v21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pli"), p0), (_index209).Int())
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_1561_v21), 0))
			_ = _index210
			(_1561_v21).ArraySet1(p0, (_index210).Int())
		} else {
			_1554_v16 = _1554_v16
			var _1563_v22 _dafny.Array
			_ = _1563_v22
			var _nwElement0_42 _dafny.Array = _1539_v3
			_ = _nwElement0_42
			var _nw230 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(11))
			_ = _nw230
			(_nw230).ArraySet1(_nwElement0_42, 0)
			(_nw230).ArraySet1((func() _dafny.Array {
				if (_this).F14() {
					return _1539_v3
				}
				return _1539_v3
			})(), 1)
			(_nw230).ArraySet1(_1539_v3, 2)
			(_nw230).ArraySet1(_1539_v3, 3)
			(_nw230).ArraySet1(_1539_v3, 4)
			(_nw230).ArraySet1(_1539_v3, 5)
			(_nw230).ArraySet1(_1539_v3, 6)
			(_nw230).ArraySet1(_1539_v3, 7)
			(_nw230).ArraySet1(_1539_v3, 8)
			(_nw230).ArraySet1(_1539_v3, 9)
			(_nw230).ArraySet1(_1539_v3, 10)
			_1563_v22 = _nw230
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1563_v22), 0))
			_ = _index211
			(_1563_v22).ArraySet1(_1539_v3, (_index211).Int())
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1563_v22), 0))
			_ = _index212
			(_1563_v22).ArraySet1(_1539_v3, (_index212).Int())
			var _1564_v23 _dafny.Map
			_ = _1564_v23
			_1564_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1544_v7).Cardinality()), (_this).F14())
			(globalState).F5 = ((_1564_v23).Cardinality()).Plus(((_this).F13()).Times((_this).F13()))
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_1539_v3), 0))
			_ = _index213
			(_1539_v3).ArraySet1((_this).F14(), (_index213).Int())
			var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_1539_v3), 0))
			_ = _index214
			(_1539_v3).ArraySet1(false, (_index214).Int())
			_1544_v7 = _1544_v7
		}
		var _1565_v24 _dafny.Set
		_ = _1565_v24
		_1565_v24 = _dafny.SetOf(true)
		(globalState).F5 = (_1565_v24).Cardinality()
		var _1566_v25 D1
		_ = _1566_v25
		_1566_v25 = Companion_D1_.Create_DC4_(Companion_D0_.Create_DC0_(_dafny.SeqOf((_this).F14(), (_this).F14()), (_this).F13()), (_dafny.Zero).Minus((_this).F13()), p1, (_this).F14(), (_this).F13())
		var _1567_i3 _dafny.Int
		_ = _1567_i3
		_1567_i3 = _dafny.Zero
		{
			for !(((_this).F13()).Cmp(_dafny.IntOfInt64(623)) != 0) || (((_this).F13()).Cmp((_1566_v25).Dtor_cf8()) >= 0) {
				{
					if (_1567_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1567_i3 = (_1567_i3).Plus(_dafny.One)
					var _1568_v26 _dafny.Array
					_ = _1568_v26
					var _nw231 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
					_ = _nw231
					_1568_v26 = _nw231
					var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_1568_v26), 0))
					_ = _index215
					(_1568_v26).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1544_v7, _dafny.Companion_Sequence_.Update(_1544_v7, (Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_1544_v7).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.SeqOf(!((_this).F14()))).Cardinality()))), (_index215).Int())
					var _1569_v27 _dafny.Map
					_ = _1569_v27
					_1569_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13())
					var _1570_v28 D0
					_ = _1570_v28
					_1570_v28 = Companion_D0_.Create_DC0_(Companion_Default___.Fm4((_this).F13(), (_this).F13(), (_this).F13(), globalState), (_this).F13())
					var _1571_v29 _dafny.Set
					_ = _1571_v29
					_1571_v29 = _dafny.SetOf(_1539_v3)
					var _pat_let_tv30 = _1570_v28
					_ = _pat_let_tv30
					var _pat_let_tv31 = globalState
					_ = _pat_let_tv31
					var _1572_v30 _dafny.Array
					_ = _1572_v30
					var _nwElement0_43 _dafny.Int = (_dafny.IntOfUint32((_1543_v6).Cardinality())).Times((_this).F13())
					_ = _nwElement0_43
					var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(5))
					_ = _nw232
					(_nw232).ArraySet1(_nwElement0_43, 0)
					(_nw232).ArraySet1((_this).F13(), 1)
					(_nw232).ArraySet1(((_1569_v27).Merge(_1569_v27)).Cardinality(), 2)
					(_nw232).ArraySet1((_dafny.Zero).Minus((func(_pat_let39_0 D1) D1 {
						return func(_1573_dt__update__tmp_h1 D1) D1 {
							return func(_pat_let40_0 D0) D1 {
								return func(_1574_dt__update_hcf7_h0 D0) D1 {
									return func(_pat_let41_0 _dafny.Sequence) D1 {
										return func(_1575_dt__update_hcf9_h0 _dafny.Sequence) D1 {
											return func(_pat_let42_0 bool) D1 {
												return func(_1576_dt__update_hcf10_h0 bool) D1 {
													return Companion_D1_.Create_DC4_(_1574_dt__update_hcf7_h0, (_1573_dt__update__tmp_h1).Dtor_cf8(), _1575_dt__update_hcf9_h0, _1576_dt__update_hcf10_h0, (_1573_dt__update__tmp_h1).Dtor_cf11())
												}(_pat_let42_0)
											}((_this).F14())
										}(_pat_let41_0)
									}(Companion_Default___.Fm18((_this).F14(), _pat_let_tv31))
								}(_pat_let40_0)
							}(_pat_let_tv30)
						}(_pat_let39_0)
					}(_1566_v25)).Dtor_cf11()), 3)
					(_nw232).ArraySet1((_1571_v29).Cardinality(), 4)
					_1572_v30 = _nw232
					var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))
					_ = _index216
					(_1572_v30).ArraySet1(Companion_Default___.Fm2((_this).F13(), globalState), (_index216).Int())
					var _1577_v31 _dafny.Map
					_ = _1577_v31
					_1577_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), p0)
					var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_1568_v26), 0))
					_ = _index217
					var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))
					_ = _index218
					var _rhs246 _dafny.Sequence = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm26(Companion_Default___.SafeModuloInt(((_1577_v31).Update((_this).F14(), p1)).Cardinality(), (_1569_v27).Cardinality()), _1552_v14, (_this).F14(), _1552_v14, globalState), (Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((Companion_Default___.Fm26(Companion_Default___.SafeModuloInt(((_1577_v31).Update((_this).F14(), p1)).Cardinality(), (_1569_v27).Cardinality()), _1552_v14, (_this).F14(), _1552_v14, globalState)).Cardinality()))).Uint32(), (_this).F13())
					_ = _rhs246
					var _rhs247 bool = ((Companion_Default___.Fm34(_dafny.IntOfUint32((p0).Cardinality()), globalState)).Intersection(_1565_v24)).IsProperSubsetOf(_1565_v24)
					_ = _rhs247
					var _rhs248 _dafny.Sequence = _1544_v7
					_ = _rhs248
					var _rhs249 _dafny.Int = (_this).F13()
					_ = _rhs249
					var _rhs250 bool = Companion_Default___.Fm3(globalState)
					_ = _rhs250
					var _lhs196 *GlobalState = globalState
					_ = _lhs196
					var _lhs197 _dafny.Array = _1568_v26
					_ = _lhs197
					var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_1568_v26), 0))
					_ = _lhs198
					var _lhs199 _dafny.Array = _1572_v30
					_ = _lhs199
					var _lhs200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))
					_ = _lhs200
					var _lhs201 *GlobalState = globalState
					_ = _lhs201
					_1544_v7 = _rhs246
					_lhs196.F0 = _rhs247
					(_lhs197).ArraySet1(_rhs248, (_lhs198).Int())
					(_lhs199).ArraySet1(_rhs249, (_lhs200).Int())
					_lhs201.F0 = _rhs250
					r1 = (_this).F13()
					if (_this).F14() {
						var _1578_v32 _dafny.Array
						_ = _1578_v32
						var _nw233 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
						_ = _nw233
						_1578_v32 = _nw233
						_1578_v32 = _1578_v32
						var _1579_v33 _dafny.Map
						_ = _1579_v33
						_1579_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1568_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_1568_v26), 0))).Int()).(_dafny.Sequence), (_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int))
						var _1580_v34 _dafny.Map
						_ = _1580_v34
						_1580_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int), _1569_v27)
						var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))
						_ = _index219
						(_1572_v30).ArraySet1((func() _dafny.Int {
							if (_1579_v33).Contains(_dafny.Companion_Sequence_.Concatenate((_1568_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_1568_v26), 0))).Int()).(_dafny.Sequence), _1544_v7)) {
								return (_1579_v33).Get(_dafny.Companion_Sequence_.Concatenate((_1568_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_1568_v26), 0))).Int()).(_dafny.Sequence), _1544_v7)).(_dafny.Int)
							}
							return (_1580_v34).Cardinality()
						})(), (_index219).Int())
						var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_1568_v26), 0))
						_ = _index220
						(_1568_v26).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm40((_this).F13(), _dafny.SetOf((_this).F14(), (_this).F14()), globalState), (_1568_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_1568_v26), 0))).Int()).(_dafny.Sequence)), (_index220).Int())
						var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))
						_ = _index221
						var _rhs251 _dafny.Int = (_dafny.Zero).Minus((_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int))
						_ = _rhs251
						var _rhs252 _dafny.Int = Companion_Default___.Fm2(_dafny.IntOfUint32((_dafny.SeqOf(_1565_v24, _1565_v24)).Cardinality()), globalState)
						_ = _rhs252
						var _lhs202 _dafny.Array = _1572_v30
						_ = _lhs202
						var _lhs203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))
						_ = _lhs203
						var _lhs204 *GlobalState = globalState
						_ = _lhs204
						(_lhs202).ArraySet1(_rhs251, (_lhs203).Int())
						_lhs204.F5 = _rhs252
						var _1581_v35 _dafny.MultiSet
						_ = _1581_v35
						_1581_v35 = _dafny.MultiSetOf(Companion_Default___.Fm3(globalState))
						(globalState).F0 = (_1542_v5).Fm22(_1581_v35, ((_this).F13()).Cmp((_this).F13()) >= 0, (_this).F13(), globalState)
					} else {
						var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1539_v3), 0))
						_ = _index222
						(_1539_v3).ArraySet1((_this).F14(), (_index222).Int())
						var _1582_v36 _dafny.Map
						_ = _1582_v36
						_1582_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int), _1543_v6)
						var _1583_v38 _dafny.Map
						_ = _1583_v38
						_1583_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int), (_this).F14())
						var _1584_v40 D11
						_ = _1584_v40
						_1584_v40 = Companion_D11_.Create_DC30_(func() _dafny.Set {
							var _coll42 = _dafny.NewBuilder()
							_ = _coll42
							for _iter44 := _dafny.Iterate((_1583_v38).Keys().Elements()); ; {
								_compr_42, _ok44 := _iter44()
								if !_ok44 {
									break
								}
								var _1585_v39 _dafny.Int
								_1585_v39 = interface{}(_compr_42).(_dafny.Int)
								if (_1583_v38).Contains(_1585_v39) {
									_coll42.Add((_1585_v39).Times(_dafny.IntOfInt64(342)))
								}
							}
							return _coll42.ToSet()
						}())
						var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1539_v3), 0))
						_ = _index223
						(_1539_v3).ArraySet1((((_1584_v40).Dtor_cf57()).Intersection(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-74))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg95 _dafny.Int) interface{} {
								return coer95(arg95)
							}
						}((func(_1586_v14 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1587_i4 _dafny.Int) _dafny.CodePoint {
								return _1586_v14
							}
						})(_1552_v14)))).Cardinality()), (_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int)))).IsProperSubsetOf((_1554_v16).Intersection(func() _dafny.Set {
							var _coll43 = _dafny.NewBuilder()
							_ = _coll43
							for _iter45 := _dafny.Iterate((_1582_v36).Keys().Elements()); ; {
								_compr_43, _ok45 := _iter45()
								if !_ok45 {
									break
								}
								var _1588_v37 _dafny.Int
								_1588_v37 = interface{}(_compr_43).(_dafny.Int)
								if (_1582_v36).Contains(_1588_v37) {
									_coll43.Add((_1588_v37).Plus(_dafny.IntOfInt64(744)))
								}
							}
							return _coll43.ToSet()
						}())), (_index223).Int())
						var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_1568_v26), 0))
						_ = _index224
						(_1568_v26).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(401))).Uint32(), func(coer96 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg96 _dafny.Int) interface{} {
								return coer96(arg96)
							}
						}((func(_1589_v16 _dafny.Set) func(_dafny.Int) _dafny.Int {
							return func(_1590_i5 _dafny.Int) _dafny.Int {
								return (_dafny.Zero).Minus((_1589_v16).Cardinality())
							}
						})(_1554_v16))), _1544_v7), (Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(401))).Uint32(), func(coer97 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg97 _dafny.Int) interface{} {
								return coer97(arg97)
							}
						}((func(_1591_v16 _dafny.Set) func(_dafny.Int) _dafny.Int {
							return func(_1592_i5 _dafny.Int) _dafny.Int {
								return (_dafny.Zero).Minus((_1591_v16).Cardinality())
							}
						})(_1554_v16))), _1544_v7)).Cardinality()))).Uint32(), (_dafny.IntOfInt64(771)).Times((_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int))), (_index224).Int())
						(globalState).F5 = (_this).F13()
						var _1593_v41 D11
						_ = _1593_v41
						_1593_v41 = Companion_D11_.Create_DC33_(((_this).F13()).Minus((_1569_v27).Cardinality()), ((_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int)).Cmp((_this).F13()) >= 0, (_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int), false, _1565_v24)
						var _1594_v43 D0
						_ = _1594_v43
						_1594_v43 = Companion_D0_.Create_DC1_((_this).F14(), (func() _dafny.Map {
							var _coll44 = _dafny.NewMapBuilder()
							_ = _coll44
							for _iter46 := _dafny.Iterate((p1).Elements()); ; {
								_compr_44, _ok46 := _iter46()
								if !_ok46 {
									break
								}
								var _1595_v42 _dafny.CodePoint
								_1595_v42 = interface{}(_compr_44).(_dafny.CodePoint)
								if _dafny.Companion_Sequence_.Contains(p1, _1595_v42) {
									_coll44.Add(_1595_v42, _1543_v6)
								}
							}
							return _coll44.ToMap()
						}()).Cardinality(), _1552_v14)
						var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))
						_ = _index225
						var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))
						_ = _index226
						var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1539_v3), 0))
						_ = _index227
						var _rhs253 D11 = _1593_v41
						_ = _rhs253
						var _rhs254 _dafny.Int = (_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int)
						_ = _rhs254
						var _rhs255 _dafny.Int = (_this).Fm16(_1551_v13, _1594_v43, globalState)
						_ = _rhs255
						var _rhs256 bool = ((_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int)).Cmp((_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int)) >= 0
						_ = _rhs256
						var _rhs257 _dafny.Set = _1571_v29
						_ = _rhs257
						var _lhs205 _dafny.Array = _1572_v30
						_ = _lhs205
						var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))
						_ = _lhs206
						var _lhs207 _dafny.Array = _1572_v30
						_ = _lhs207
						var _lhs208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))
						_ = _lhs208
						var _lhs209 _dafny.Array = _1539_v3
						_ = _lhs209
						var _lhs210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1539_v3), 0))
						_ = _lhs210
						_1593_v41 = _rhs253
						(_lhs205).ArraySet1(_rhs254, (_lhs206).Int())
						(_lhs207).ArraySet1(_rhs255, (_lhs208).Int())
						(_lhs209).ArraySet1(_rhs256, (_lhs210).Int())
						_1571_v29 = _rhs257
						var _1596_v44 _dafny.MultiSet
						_ = _1596_v44
						_1596_v44 = _dafny.MultiSetOf((_1539_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1539_v3), 0))).Int()).(bool), (_this).F14(), (_1539_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1539_v3), 0))).Int()).(bool))
						var _1597_v45 D5
						_ = _1597_v45
						_1597_v45 = Companion_D5_.Create_DC17_(_1544_v7, !((_this).F14()))
						var _rhs258 bool = (_1542_v5).Fm22(_1596_v44, (_1577_v31).Contains((_this).F14()), ((_this).F13()).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uuiikup")).Cardinality())), globalState)
						_ = _rhs258
						var _rhs259 bool = ((_1554_v16).Intersection(_dafny.SetOf(_dafny.IntOfUint32((_1543_v6).Cardinality()), (_this).F13(), (_this).F13(), (_this).F13()))).IsProperSubsetOf(_1554_v16)
						_ = _rhs259
						var _rhs260 bool = (func() bool {
							if (_1583_v38).Contains(_dafny.IntOfUint32((p1).Cardinality())) {
								return (_1583_v38).Get(_dafny.IntOfUint32((p1).Cardinality())).(bool)
							}
							return (_this).F14()
						})()
						_ = _rhs260
						var _rhs261 _dafny.Int = (func() _dafny.Int {
							if (_this).F14() {
								return (_dafny.Zero).Minus((_1572_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1572_v30), 0))).Int()).(_dafny.Int))
							}
							return Companion_Default___.SafeDivisionInt((_this).F13(), (_this).F13())
						})()
						_ = _rhs261
						var _rhs262 bool = (_1597_v45).Dtor_cf37()
						_ = _rhs262
						var _lhs211 *GlobalState = globalState
						_ = _lhs211
						var _lhs212 *GlobalState = globalState
						_ = _lhs212
						var _lhs213 *GlobalState = globalState
						_ = _lhs213
						var _lhs214 *GlobalState = globalState
						_ = _lhs214
						var _lhs215 *GlobalState = globalState
						_ = _lhs215
						_lhs211.F0 = _rhs258
						_lhs212.F0 = _rhs259
						_lhs213.F0 = _rhs260
						_lhs214.F5 = _rhs261
						_lhs215.F0 = _rhs262
					}
					(globalState).F2 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xpbyxlev"), Companion_Default___.Fm18((_this).F14(), globalState))
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1598_v46 _dafny.Array
		_ = _1598_v46
		var _len0_41 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_41
		var _nw234 _dafny.Array
		_ = _nw234
		if _len0_41.Cmp(_dafny.Zero) == 0 {
			_nw234 = _dafny.NewArray(_len0_41)
		} else {
			var _init41 func(_dafny.Int) _dafny.Int = func(_1599_i6 _dafny.Int) _dafny.Int {
				return (_1599_i6).Times((_this).F13())
			}
			_ = _init41
			var _element0_41 = _init41(_dafny.Zero)
			_ = _element0_41
			_nw234 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
			(_nw234).ArraySet1(_element0_41, 0)
			var _nativeLen0_41 = (_len0_41).Int()
			_ = _nativeLen0_41
			for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
				(_nw234).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
			}
		}
		_1598_v46 = _nw234
		var _1600_v47 _dafny.Set
		_ = _1600_v47
		_1600_v47 = _dafny.SetOf(_1598_v46, _1598_v46, _1598_v46, _1598_v46)
		r0 = _1600_v47
		r1 = ((_this).F13()).Plus(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F13()), (_this).F13()))
		r2 = (_1537_v1).Intersection((_1537_v1).Difference(_1537_v1))
		return r0, r1, r2
	}
}
func (_this *C6) F13() _dafny.Int {
	{
		return _this._f13
	}
}
func (_this *C6) F14() bool {
	{
		return _this._f14
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f9  _dafny.Array
	_f6  _dafny.Int
	_f10 _dafny.Map
	F11  _dafny.Array
	F12  _dafny.Array
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.Zero
	_this._f10 = _dafny.EmptyMap
	_this.F11 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F12 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C7{}
var _ T1 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) F9() _dafny.Array {
	return _this._f9
}
func (_this *C7) F9_set_(value _dafny.Array) {
	_this._f9 = value
}
func (_this *C7) F6() _dafny.Int {
	return _this._f6
}
func (_this *C7) F10() _dafny.Map {
	return _this._f10
}
func (_this *C7) Ctor__(f11 _dafny.Array, f12 _dafny.Array, f6 _dafny.Int, f9 _dafny.Array, f10 _dafny.Map) {
	{
		(_this).F11 = f11
		(_this).F12 = f12
		(_this)._f6 = f6
		(_this)._f9 = f9
		(_this)._f10 = f10
	}
}
func (_this *C7) Fm5(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-609))).Uint32(), func(coer98 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg98 _dafny.Int) interface{} {
				return coer98(arg98)
			}
		}(func(_1601_i0 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetOf(_dafny.CodePoint('u'))
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(596))).Uint32(), func(coer99 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg99 _dafny.Int) interface{} {
				return coer99(arg99)
			}
		}(func(_1602_i1 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetOf(_dafny.CodePoint('u'), _dafny.CodePoint('l'), (Companion_D0_.Create_DC1_(false, _dafny.IntOfInt64(711), _dafny.CodePoint('g'))).Dtor_cf4(), _dafny.CodePoint('k'))
		}))), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('l'), _dafny.CodePoint('b'))))
	}
}
func (_this *C7) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F6()), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F6())).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F6(), (_this).F6())).Cardinality()), (_this).F6(), (_this).F6(), (_this).F6())))
	}
}
func (_this *C7) Fm13(p0 _dafny.CodePoint, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return ((func() _dafny.Map {
			if false {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("j"))
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("ggjslaqas"))
		})()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("kdhe"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(298))).Uint32(), func(coer100 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg100 _dafny.Int) interface{} {
				return coer100(arg100)
			}
		}(func(_1603_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})))))
	}
}
func (_this *C7) Fm14(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(((Companion_D0_.Create_DC0_(_dafny.SeqOf(false, false), (_this).F6())).Dtor_cf1()).Minus(((_this).F6()).Plus((_this).F6())))
	}
}
func (_this *C7) M6(p0 bool, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _hi12 _dafny.Int = (_this).Fm14(globalState)
		_ = _hi12
		for _1604_i0 := _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0, p0), p1)).Cardinality()); _1604_i0.Cmp(_hi12) < 0; _1604_i0 = _1604_i0.Plus(_dafny.One) {
			(globalState).F5 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if p0 {
					return p2
				}
				return _1604_i0
			})(), _1604_i0))
			(globalState).F2 = _dafny.UnicodeSeqOfUtf8Bytes("bksq")
			var _1605_v0 _dafny.Sequence
			_ = _1605_v0
			_1605_v0 = _dafny.UnicodeSeqOfUtf8Bytes("rwrd")
			var _1606_v1 D0
			_ = _1606_v1
			_1606_v1 = Companion_D0_.Create_DC0_(Companion_Default___.Fm4(_1604_i0, _1604_i0, (_this).F6(), globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1605_v0, _1605_v0)).Cardinality()))
			_1606_v1 = _1606_v1
			var _1607_v2 _dafny.MultiSet
			_ = _1607_v2
			_1607_v2 = _dafny.MultiSetOf(_this.F9())
			if (_1607_v2).IsProperSubsetOf(_1607_v2) {
				var _1608_v3 *C0
				_ = _1608_v3
				var _nw235 *C0 = New_C0_()
				_ = _nw235
				_nw235.Ctor__(_this.F9(), p0)
				_1608_v3 = _nw235
				(globalState).F5 = (_this).F6()
				(globalState).F5 = ((p2).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vsvuef")).Cardinality()))).Plus(p2)
				var _1609_v4 _dafny.Array
				_ = _1609_v4
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_42
				var _nw236 _dafny.Array
				_ = _nw236
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw236 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) _dafny.Sequence = (func(_1610_v3 *C0) func(_dafny.Int) _dafny.Sequence {
						return func(_1611_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_1610_v3.F20)
						}
					})(_1608_v3)
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw236 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw236).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw236).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1609_v4 = _nw236
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_1609_v4), 0))
				_ = _index228
				(_1609_v4).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1608_v3.F20, _1608_v3.F20), p1), (_index228).Int())
				var _1612_v5 _dafny.Sequence
				_ = _1612_v5
				_1612_v5 = _dafny.SeqOf(p0)
				var _1613_v6 _dafny.Map
				_ = _1613_v6
				_1613_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(551))).Uint32(), func(coer101 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg101 _dafny.Int) interface{} {
						return coer101(arg101)
					}
				}(func(_1614_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				}))).Cardinality()), _dafny.SeqOf(!(_1608_v3.F20), _1608_v3.F20))
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_1609_v4), 0))
				_ = _index229
				var _rhs263 bool = Companion_Default___.Fm3(globalState)
				_ = _rhs263
				var _rhs264 _dafny.Sequence = (func() _dafny.Sequence {
					if (_1613_v6).Contains((_this).F6()) {
						return (_1613_v6).Get((_this).F6()).(_dafny.Sequence)
					}
					return _1612_v5
				})()
				_ = _rhs264
				var _rhs265 _dafny.Sequence = p1
				_ = _rhs265
				var _rhs266 _dafny.Array = _this.F12
				_ = _rhs266
				var _rhs267 _dafny.Int = _1604_i0
				_ = _rhs267
				var _lhs216 *GlobalState = globalState
				_ = _lhs216
				var _lhs217 _dafny.Array = _1609_v4
				_ = _lhs217
				var _lhs218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_1609_v4), 0))
				_ = _lhs218
				var _lhs219 *C7 = _this
				_ = _lhs219
				var _lhs220 *GlobalState = globalState
				_ = _lhs220
				_lhs216.F0 = _rhs263
				(_lhs217).ArraySet1(_rhs264, (_lhs218).Int())
				_1612_v5 = _rhs265
				_lhs219.F11 = _rhs266
				_lhs220.F5 = _rhs267
				var _1615_v8 _dafny.Set
				_ = _1615_v8
				_1615_v8 = _dafny.SetOf(false, _1608_v3.F20, _1608_v3.F20)
				var _1616_v9 _dafny.Map
				_ = _1616_v9
				_1616_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(930), (Companion_Default___.Fm2(_1604_i0, globalState)).Cmp((func() _dafny.Map {
					var _coll45 = _dafny.NewMapBuilder()
					_ = _coll45
					for _iter47 := _dafny.Iterate((_dafny.SeqOf(_1615_v8, _1615_v8)).Elements()); ; {
						_compr_45, _ok47 := _iter47()
						if !_ok47 {
							break
						}
						var _1617_v7 _dafny.Set
						_1617_v7 = interface{}(_compr_45).(_dafny.Set)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_1615_v8, _1615_v8), _1617_v7) {
							_coll45.Add(_1617_v7, _1608_v3.F20)
						}
					}
					return _coll45.ToMap()
				}()).Cardinality()) == 0)
				(_1608_v3).F20 = (func() bool {
					if (_1616_v9).Contains(Companion_Default___.SafeModuloInt(p2, _dafny.IntOfUint32((_1605_v0).Cardinality()))) {
						return (_1616_v9).Get(Companion_Default___.SafeModuloInt(p2, _dafny.IntOfUint32((_1605_v0).Cardinality()))).(bool)
					}
					return p0
				})()
			} else {
				var _1618_v10 _dafny.MultiSet
				_ = _1618_v10
				_1618_v10 = _dafny.MultiSetOf(p2)
				var _1619_v11 _dafny.Map
				_ = _1619_v11
				_1619_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1618_v10, p2)
				_1619_v11 = _1619_v11
				var _1620_v12 D2
				_ = _1620_v12
				_1620_v12 = Companion_D2_.Create_DC6_(_dafny.IntOfInt64(835), _this.F11)
				var _1621_v13 _dafny.Array
				_ = _1621_v13
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_43
				var _nw237 _dafny.Array
				_ = _nw237
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw237 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) _dafny.Sequence = func(_1622_i3 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(957))).Uint32(), func(coer102 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg102 _dafny.Int) interface{} {
								return coer102(arg102)
							}
						}(func(_1623_i4 _dafny.Int) _dafny.Int {
							return (_this).F6()
						}))
					}
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw237 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw237).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw237).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_1621_v13 = _nw237
				var _1624_v14 _dafny.Sequence
				_ = _1624_v14
				_1624_v14 = _dafny.SeqOf((_1618_v10).Cardinality())
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_1621_v13), 0))
				_ = _index230
				(_1621_v13).ArraySet1(_1624_v14, (_index230).Int())
				var _1625_v15 _dafny.Array
				_ = _1625_v15
				var _nw238 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(20))
				_ = _nw238
				_1625_v15 = _nw238
				var _pat_let_tv32 = p1
				_ = _pat_let_tv32
				var _1626_v16 _dafny.Array
				_ = _1626_v16
				var _nwElement0_44 D0 = _1606_v1
				_ = _nwElement0_44
				var _nw239 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(16))
				_ = _nw239
				(_nw239).ArraySet1(_nwElement0_44, 0)
				(_nw239).ArraySet1(Companion_D0_.Create_DC0_(p1, _dafny.IntOfUint32((_1605_v0).Cardinality())), 1)
				(_nw239).ArraySet1(_1606_v1, 2)
				(_nw239).ArraySet1(_1606_v1, 3)
				(_nw239).ArraySet1(_1606_v1, 4)
				(_nw239).ArraySet1(Companion_Default___.Fm32(p0, p0, p2, p0, globalState), 5)
				(_nw239).ArraySet1(_1606_v1, 6)
				(_nw239).ArraySet1(_1606_v1, 7)
				(_nw239).ArraySet1(_1606_v1, 8)
				(_nw239).ArraySet1(Companion_Default___.Fm32(p0, p0, _1604_i0, p0, globalState), 9)
				(_nw239).ArraySet1(_1606_v1, 10)
				(_nw239).ArraySet1(Companion_D0_.Create_DC0_(_dafny.SeqOf(p0), _1604_i0), 11)
				(_nw239).ArraySet1(_1606_v1, 12)
				(_nw239).ArraySet1(func(_pat_let43_0 D0) D0 {
					return func(_1627_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let44_0 _dafny.Sequence) D0 {
							return func(_1628_dt__update_hcf0_h0 _dafny.Sequence) D0 {
								return Companion_D0_.Create_DC0_(_1628_dt__update_hcf0_h0, (_1627_dt__update__tmp_h0).Dtor_cf1())
							}(_pat_let44_0)
						}(_pat_let_tv32)
					}(_pat_let43_0)
				}(_1606_v1), 13)
				(_nw239).ArraySet1(_1606_v1, 14)
				(_nw239).ArraySet1(_1606_v1, 15)
				_1626_v16 = _nw239
				var _1629_v17 D7
				_ = _1629_v17
				_1629_v17 = Companion_D7_.Create_DC21_(_1626_v16)
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1625_v15), 0))
				_ = _index231
				(_1625_v15).ArraySet1(_1629_v17, (_index231).Int())
				var _1630_v18 _dafny.Sequence
				_ = _1630_v18
				_1630_v18 = _dafny.SeqOf(_this.F12, _this.F11)
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_1621_v13), 0))
				_ = _index232
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1625_v15), 0))
				_ = _index233
				var _rhs268 D2 = _1620_v12
				_ = _rhs268
				var _rhs269 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)).Cardinality()), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)).Cardinality())).Cardinality()))).Uint32(), _dafny.IntOfInt64(848)), _dafny.Companion_Sequence_.Concatenate(_1624_v14, _1624_v14))
				_ = _rhs269
				var _rhs270 D7 = _1629_v17
				_ = _rhs270
				var _rhs271 bool = p0
				_ = _rhs271
				var _rhs272 _dafny.Array = (_1630_v18).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1630_v18).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs272
				var _lhs221 _dafny.Array = _1621_v13
				_ = _lhs221
				var _lhs222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_1621_v13), 0))
				_ = _lhs222
				var _lhs223 _dafny.Array = _1625_v15
				_ = _lhs223
				var _lhs224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1625_v15), 0))
				_ = _lhs224
				var _lhs225 *GlobalState = globalState
				_ = _lhs225
				var _lhs226 *C7 = _this
				_ = _lhs226
				_1620_v12 = _rhs268
				(_lhs221).ArraySet1(_rhs269, (_lhs222).Int())
				(_lhs223).ArraySet1(_rhs270, (_lhs224).Int())
				_lhs225.F0 = _rhs271
				_lhs226.F11 = _rhs272
				var _1631_v19 _dafny.MultiSet
				_ = _1631_v19
				_1631_v19 = _dafny.MultiSetOf(_1618_v10)
				var _1632_v20 *C2
				_ = _1632_v20
				var _nw240 *C2 = New_C2_()
				_ = _nw240
				_nw240.Ctor__(((_this).F6()).Plus(_1604_i0), (func() _dafny.MultiSet {
					if !(p0) {
						return _1631_v19
					}
					return _1631_v19
				})(), _this.F9(), (_this).F10())
				_1632_v20 = _nw240
				var _1633_v21 _dafny.Map
				_ = _1633_v21
				_1633_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1605_v0)
				_1633_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.UnicodeSeqOfUtf8Bytes("dkcllx"))
				(globalState).F0 = (func() bool {
					if p0 {
						return p0
					}
					return p0
				})()
			}
		}
		var _arr46 _dafny.Array = _this.F9()
		_ = _arr46
		var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))
		_ = _index234
		(_arr46).ArraySet1(p0, (_index234).Int())
		var _1634_v22 _dafny.Map
		_ = _1634_v22
		_1634_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F6())
		var _arr47 _dafny.Array = _this.F9()
		_ = _arr47
		var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))
		_ = _index235
		(_arr47).ArraySet1((func() bool {
			if p0 {
				return !((_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0))).Cardinality())).Cmp((_1634_v22).Cardinality()) == 0)
			}
			return (func() bool {
				if p0 {
					return p0
				}
				return p0
			})()
		})(), (_index235).Int())
		if p0 {
			var _1635_v23 _dafny.Map
			_ = _1635_v23
			_1635_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), !(p0))
			var _1636_v24 _dafny.Map
			_ = _1636_v24
			_1636_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1635_v23, (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
			var _1637_v25 D3
			_ = _1637_v25
			_1637_v25 = Companion_D3_.Create_DC10_(_1636_v24)
			var _1638_v26 D3
			_ = _1638_v26
			_1638_v26 = Companion_D3_.Create_DC11_(_1637_v25)
			var _1639_v27 _dafny.Map
			_ = _1639_v27
			_1639_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
			_1638_v26 = Companion_Default___.Fm50((func() _dafny.Sequence {
				if (_1639_v27).Contains((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)) {
					return (_1639_v27).Get((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)).(_dafny.Sequence)
				}
				return p1
			})(), globalState)
			(globalState).F5 = (_dafny.Zero).Minus((_this).F6())
			var _1640_v28 _dafny.Sequence
			_ = _1640_v28
			_1640_v28 = _dafny.UnicodeSeqOfUtf8Bytes("bvftkwuvg")
			(globalState).F5 = Companion_Default___.SafeModuloInt((_dafny.IntOfUint32((_1640_v28).Cardinality())).Plus((_this).F6()), p2)
			if (p2).Cmp(((_this).F6()).Times((_this).F6())) > 0 {
				var _1641_v29 _dafny.Sequence
				_ = _1641_v29
				_1641_v29 = _dafny.SeqOf(_this.F11)
				var _1642_v30 _dafny.Map
				_ = _1642_v30
				_1642_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)), p0)
				var _1643_v31 D4
				_ = _1643_v31
				_1643_v31 = Companion_D4_.Create_DC14_((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), _1642_v30, (_this).F6(), _1642_v30)
				var _1644_v32 _dafny.MultiSet
				_ = _1644_v32
				_1644_v32 = _dafny.MultiSetOf((_1643_v31).Dtor_cf28())
				var _1645_v33 _dafny.Sequence
				_ = _1645_v33
				_1645_v33 = _dafny.SeqOf((func() _dafny.Int {
					if (_1644_v32).Contains((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)) {
						return (_1644_v32).Multiplicity((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
					}
					return (_this).F6()
				})(), (_this).F6())
				var _rhs273 _dafny.Sequence = _1641_v29
				_ = _rhs273
				var _rhs274 bool = ((_1645_v33).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1645_v33).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(Companion_Default___.Fm2(p2, globalState)) < 0
				_ = _rhs274
				var _lhs227 *GlobalState = globalState
				_ = _lhs227
				_1641_v29 = _rhs273
				_lhs227.F0 = _rhs274
				var _1646_v34 _dafny.Array
				_ = _1646_v34
				var _nw241 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
				_ = _nw241
				_1646_v34 = _nw241
				var _1647_v35 _dafny.Map
				_ = _1647_v35
				_1647_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1646_v34, (_this).F6())
				(globalState).F5 = (func() _dafny.Int {
					if (_1647_v35).Contains(_1646_v34) {
						return (_1647_v35).Get(_1646_v34).(_dafny.Int)
					}
					return p2
				})()
				_1645_v33 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1645_v33, _dafny.SeqOf(Companion_Default___.Fm2(p2, globalState))), (Companion_Default___.SafeIndex(Companion_Default___.Fm2(_dafny.IntOfUint32((p1).Cardinality()), globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1645_v33, _dafny.SeqOf(Companion_Default___.Fm2(p2, globalState)))).Cardinality()))).Uint32(), (_dafny.IntOfInt64(745)).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(94))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg103 _dafny.Int) interface{} {
						return coer103(arg103)
					}
				}(func(_1648_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}))).Cardinality())))
				var _arr48 _dafny.Array = _this.F9()
				_ = _arr48
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))
				_ = _index236
				(_arr48).ArraySet1(p0, (_index236).Int())
				var _1649_v36 _dafny.MultiSet
				_ = _1649_v36
				_1649_v36 = _dafny.MultiSetOf((_dafny.Zero).Minus((_this).F6()))
				var _1650_v37 _dafny.MultiSet
				_ = _1650_v37
				_1650_v37 = _dafny.MultiSetOf(_1649_v36, _1649_v36, _1649_v36)
				var _1651_v38 *C2
				_ = _1651_v38
				var _nw242 *C2 = New_C2_()
				_ = _nw242
				_nw242.Ctor__((_dafny.Zero).Minus((_1635_v23).Cardinality()), (_1650_v37).Difference(_1650_v37), _1646_v34, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (func() _dafny.Sequence {
					if ((_this).F10()).Contains(_dafny.IntOfUint32((_1640_v28).Cardinality())) {
						return ((_this).F10()).Get(_dafny.IntOfUint32((_1640_v28).Cardinality())).(_dafny.Sequence)
					}
					return _1640_v28
				})())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F6())).Cardinality(), _1640_v28)))
				_1651_v38 = _nw242
			} else {
				var _1652_v39 *C6
				_ = _1652_v39
				var _nw243 *C6 = New_C6_()
				_ = _nw243
				_nw243.Ctor__(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p2), (_this).F6()), (func() bool {
					if (_1635_v23).Contains(p2) {
						return (_1635_v23).Get(p2).(bool)
					}
					return (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)
				})())
				_1652_v39 = _nw243
				_1652_v39 = _1652_v39
				var _1653_v40 _dafny.Map
				_ = _1653_v40
				_1653_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfInt64(-125))
				var _1654_v41 _dafny.Sequence
				_ = _1654_v41
				_1654_v41 = _dafny.SeqOf((_1652_v39).F13())
				var _1655_v42 _dafny.Array
				_ = _1655_v42
				var _nwElement0_45 bool = (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)
				_ = _nwElement0_45
				var _nw244 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(16))
				_ = _nw244
				(_nw244).ArraySet1(_nwElement0_45, 0)
				(_nw244).ArraySet1(p0, 1)
				(_nw244).ArraySet1(p0, 2)
				(_nw244).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 3)
				(_nw244).ArraySet1((p1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1654_v41).Cardinality()), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool), 4)
				(_nw244).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 5)
				(_nw244).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 6)
				(_nw244).ArraySet1(p0, 7)
				(_nw244).ArraySet1(p0, 8)
				(_nw244).ArraySet1(p0, 9)
				(_nw244).ArraySet1(p0, 10)
				(_nw244).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 11)
				(_nw244).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 12)
				(_nw244).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 13)
				(_nw244).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 14)
				(_nw244).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 15)
				_1655_v42 = _nw244
				var _1656_v43 *C3
				_ = _1656_v43
				var _nw245 *C3 = New_C3_()
				_ = _nw245
				_nw245.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_1653_v40).Cardinality()), _1655_v42, (_this).F10(), (_1654_v41).Select((Companion_Default___.SafeIndex((_1652_v39).F13(), _dafny.IntOfUint32((_1654_v41).Cardinality()))).Uint32()).(_dafny.Int))
				_1656_v43 = _nw245
				var _1657_v44 _dafny.Set
				_ = _1657_v44
				_1657_v44 = _dafny.SetOf((_1652_v39).F13(), (_this).F6())
				var _1658_v45 D11
				_ = _1658_v45
				_1658_v45 = Companion_D11_.Create_DC30_(_1657_v44)
				_1658_v45 = _1658_v45
				var _1659_v46 *C5
				_ = _1659_v46
				var _nw246 *C5 = New_C5_()
				_ = _nw246
				_nw246.Ctor__(((_this).F6()).Times((_1652_v39).F13()), _dafny.IntOfInt64(690))
				_1659_v46 = _nw246
				var _1660_v47 _dafny.Array
				_ = _1660_v47
				var _nw247 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(23))
				_ = _nw247
				_1660_v47 = _nw247
				var _pat_let_tv33 = _1637_v25
				_ = _pat_let_tv33
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_1660_v47), 0))
				_ = _index237
				(_1660_v47).ArraySet1(func(_pat_let45_0 D3) D3 {
					return func(_1661_dt__update__tmp_h1 D3) D3 {
						return func(_pat_let46_0 D3) D3 {
							return func(_1662_dt__update_hcf23_h0 D3) D3 {
								return Companion_D3_.Create_DC11_(_1662_dt__update_hcf23_h0)
							}(_pat_let46_0)
						}(_pat_let_tv33)
					}(_pat_let45_0)
				}(_1638_v26), (_index237).Int())
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_1660_v47), 0))
				_ = _index238
				(_1660_v47).ArraySet1(_1638_v26, (_index238).Int())
			}
			(globalState).F5 = (Companion_D1_.Create_DC4_(Companion_Default___.Fm32(p0, p0, (_this).F6(), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), globalState), p2, _1640_v28, (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), p2)).Dtor_cf8()
		} else {
			var _1663_v48 _dafny.CodePoint
			_ = _1663_v48
			_1663_v48 = _dafny.CodePoint('u')
			var _1664_v49 _dafny.Map
			_ = _1664_v49
			_1664_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1663_v48, !(p0))
			_1664_v49 = (_1664_v49).Update(_1663_v48, p0)
			var _1665_v50 _dafny.MultiSet
			_ = _1665_v50
			_1665_v50 = _dafny.MultiSetOf(p0)
			var _1666_v51 _dafny.Set
			_ = _1666_v51
			_1666_v51 = _dafny.SetOf((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
			(globalState).F5 = (_dafny.Zero).Minus(((func() _dafny.Int {
				if (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool) {
					return p2
				}
				return (func() _dafny.Int {
					if (_1665_v50).Contains((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)) {
						return (_1665_v50).Multiplicity((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
					}
					return (_this).F6()
				})()
			})()).Minus((func() _dafny.Int {
				if true {
					return (_this).F6()
				}
				return (_1666_v51).Cardinality()
			})()))
			var _1667_v52 D0
			_ = _1667_v52
			_1667_v52 = Companion_D0_.Create_DC1_((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), p2, _1663_v48)
			var _rhs275 _dafny.CodePoint = (_1667_v52).Dtor_cf4()
			_ = _rhs275
			var _rhs276 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(p2)))
			_ = _rhs276
			var _lhs228 *GlobalState = globalState
			_ = _lhs228
			_1663_v48 = _rhs275
			_lhs228.F5 = _rhs276
			var _rhs277 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(p2, (_this).Fm14(globalState)), (_this).F6())
			_ = _rhs277
			var _rhs278 _dafny.Int = ((_this).F6()).Minus(((_this).F6()).Times(p2))
			_ = _rhs278
			var _lhs229 *GlobalState = globalState
			_ = _lhs229
			var _lhs230 *GlobalState = globalState
			_ = _lhs230
			_lhs229.F5 = _rhs277
			_lhs230.F5 = _rhs278
			var _1668_v53 _dafny.MultiSet
			_ = _1668_v53
			_1668_v53 = _dafny.MultiSetOf(_1663_v48)
			var _1669_v54 _dafny.Sequence
			_ = _1669_v54
			_1669_v54 = _dafny.SeqOf(_1668_v53, _1668_v53, _1668_v53)
			var _1670_v55 D4
			_ = _1670_v55
			_1670_v55 = Companion_D4_.Create_DC12_(_1669_v54)
			var _1671_v56 _dafny.Map
			_ = _1671_v56
			_1671_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1670_v55, p2)
			var _1672_v57 _dafny.Sequence
			_ = _1672_v57
			_1672_v57 = _dafny.UnicodeSeqOfUtf8Bytes("ogh")
			var _arr49 _dafny.Array = _this.F9()
			_ = _arr49
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index239
			var _rhs279 bool = (func() bool {
				if (Companion_Default___.Fm2((_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F6(), (_this).F6()))).Cardinality(), globalState)).Cmp((_1671_v56).Cardinality()) >= 0 {
					return (!(!(p0))) == (Companion_Default___.Fm3(globalState))
				}
				return !(!(!((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))))
			})()
			_ = _rhs279
			var _rhs280 _dafny.Sequence = _1672_v57
			_ = _rhs280
			var _lhs231 _dafny.Array = _this.F9()
			_ = _lhs231
			var _lhs232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))
			_ = _lhs232
			var _lhs233 *GlobalState = globalState
			_ = _lhs233
			(_lhs231).ArraySet1(_rhs279, (_lhs232).Int())
			_lhs233.F2 = _rhs280
		}
		var _1673_v58 _dafny.MultiSet
		_ = _1673_v58
		_1673_v58 = _dafny.MultiSetOf((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
		var _1674_v59 D6
		_ = _1674_v59
		_1674_v59 = Companion_D6_.Create_DC19_((_this).F6(), (_this).F6(), _dafny.IntOfInt64(428))
		var _1675_v60 _dafny.Map
		_ = _1675_v60
		_1675_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_1673_v58).Contains((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)) {
				return (_1673_v58).Multiplicity((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
			}
			return (_dafny.Zero).Minus(p2)
		})(), (_1674_v59).Dtor_cf41())
		_1675_v60 = (func() _dafny.Map {
			var _coll46 = _dafny.NewMapBuilder()
			_ = _coll46
			for _iter48 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F6()), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqOf((_this).F6())).Cardinality()))).Uint32(), (_1675_v60).Cardinality())).Elements()); ; {
				_compr_46, _ok48 := _iter48()
				if !_ok48 {
					break
				}
				var _1676_v61 _dafny.Int
				_1676_v61 = interface{}(_compr_46).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F6()), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqOf((_this).F6())).Cardinality()))).Uint32(), (_1675_v60).Cardinality()), _1676_v61) {
					_coll46.Add(Companion_Default___.SafeDivisionInt(_1676_v61, (_this).F6()), p2)
				}
			}
			return _coll46.ToMap()
		}()).Update(p2, (_this).F6())
		var _1677_v62 _dafny.CodePoint
		_ = _1677_v62
		_1677_v62 = _dafny.CodePoint('y')
		(globalState).F5 = (_dafny.IntOfUint32((_dafny.SeqOf(_1677_v62)).Cardinality())).Plus((_this).F6())
		var _1678_v63 _dafny.Array
		_ = _1678_v63
		var _len0_44 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_44
		var _nw248 _dafny.Array
		_ = _nw248
		if _len0_44.Cmp(_dafny.Zero) == 0 {
			_nw248 = _dafny.NewArray(_len0_44)
		} else {
			var _init44 func(_dafny.Int) _dafny.Map = (func(_1679_p0 bool) func(_dafny.Int) _dafny.Map {
				return func(_1680_i6 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1679_p0, (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
				}
			})(p0)
			_ = _init44
			var _element0_44 = _init44(_dafny.Zero)
			_ = _element0_44
			_nw248 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
			(_nw248).ArraySet1(_element0_44, 0)
			var _nativeLen0_44 = (_len0_44).Int()
			_ = _nativeLen0_44
			for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
				(_nw248).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
			}
		}
		_1678_v63 = _nw248
		var _1681_v64 D9
		_ = _1681_v64
		_1681_v64 = Companion_D9_.Create_DC26_(_1678_v63)
		var _source20 D9 = _1681_v64
		_ = _source20
		if _source20.Is_DC25() {
			var _1682___mcc_h0 _dafny.Int = _source20.Get_().(D9_DC25).Cf46
			_ = _1682___mcc_h0
			var _1683___mcc_h1 _dafny.Int = _source20.Get_().(D9_DC25).Cf47
			_ = _1683___mcc_h1
			var _1684___mcc_h2 _dafny.Int = _source20.Get_().(D9_DC25).Cf48
			_ = _1684___mcc_h2
			var _1685___mcc_h3 _dafny.CodePoint = _source20.Get_().(D9_DC25).Cf49
			_ = _1685___mcc_h3
			var _1686_cf49 _dafny.CodePoint = _1685___mcc_h3
			_ = _1686_cf49
			var _1687_cf48 _dafny.Int = _1684___mcc_h2
			_ = _1687_cf48
			var _1688_cf47 _dafny.Int = _1683___mcc_h1
			_ = _1688_cf47
			var _1689_cf46 _dafny.Int = _1682___mcc_h0
			_ = _1689_cf46
			var _1690_v65 _dafny.Sequence
			_ = _1690_v65
			_1690_v65 = _dafny.SeqOf(_1689_cf46)
			var _1691_v66 _dafny.MultiSet
			_ = _1691_v66
			_1691_v66 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1690_v65).Cardinality()), _1688_cf47)
			var _1692_v67 _dafny.Sequence
			_ = _1692_v67
			_1692_v67 = _dafny.SeqOf(_1686_cf49)
			var _1693_v68 _dafny.Map
			_ = _1693_v68
			_1693_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1691_v66).Cardinality()).Plus(_1688_cf47), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(50))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg104 _dafny.Int) interface{} {
					return coer104(arg104)
				}
			}((func(_1694_cf49 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1695_i7 _dafny.Int) _dafny.CodePoint {
					return _1694_cf49
				}
			})(_1686_cf49))), _1692_v67))
			_1693_v68 = (_1693_v68).Update((_1690_v65).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfInt64(-709)), _dafny.IntOfUint32((_1690_v65).Cardinality()))).Uint32()).(_dafny.Int), _1692_v67)
			if p0 {
				var _1696_v69 D2
				_ = _1696_v69
				_1696_v69 = Companion_D2_.Create_DC6_(_1687_cf48, _this.F11)
				var _1697_v70 _dafny.Sequence
				_ = _1697_v70
				_1697_v70 = _dafny.SeqOf((_1696_v69).Dtor_cf14(), _this.F11)
				_1697_v70 = _dafny.SeqOf(_this.F12, _this.F12, _this.F11)
				(globalState).F0 = p0
				var _1698_v71 _dafny.Sequence
				_ = _1698_v71
				_1698_v71 = _dafny.SeqOf(_1690_v65)
				var _1699_v72 _dafny.Set
				_ = _1699_v72
				_1699_v72 = _dafny.SetOf((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
				var _1700_v73 _dafny.Array
				_ = _1700_v73
				var _nwElement0_46 _dafny.Sequence = _1690_v65
				_ = _nwElement0_46
				var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(28))
				_ = _nw249
				(_nw249).ArraySet1(_nwElement0_46, 0)
				(_nw249).ArraySet1(_1690_v65, 1)
				(_nw249).ArraySet1(_1690_v65, 2)
				(_nw249).ArraySet1(_1690_v65, 3)
				(_nw249).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-787))).Uint32(), func(coer105 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg105 _dafny.Int) interface{} {
						return coer105(arg105)
					}
				}(func(_1701_i8 _dafny.Int) _dafny.Int {
					return (_this).F6()
				})), 4)
				(_nw249).ArraySet1(_1690_v65, 5)
				(_nw249).ArraySet1(_1690_v65, 6)
				(_nw249).ArraySet1(_dafny.Companion_Sequence_.Update(_1690_v65, (Companion_Default___.SafeIndex(Companion_Default___.Fm2(p2, globalState), _dafny.IntOfUint32((_1690_v65).Cardinality()))).Uint32(), (_this).F6()), 7)
				(_nw249).ArraySet1(_1690_v65, 8)
				(_nw249).ArraySet1(_1690_v65, 9)
				(_nw249).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-670))).Uint32(), func(coer106 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg106 _dafny.Int) interface{} {
						return coer106(arg106)
					}
				}(func(_1702_i9 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(-939)
				})), 10)
				(_nw249).ArraySet1(_1690_v65, 11)
				(_nw249).ArraySet1(_1690_v65, 12)
				(_nw249).ArraySet1(_1690_v65, 13)
				(_nw249).ArraySet1(_1690_v65, 14)
				(_nw249).ArraySet1(_1690_v65, 15)
				(_nw249).ArraySet1(_1690_v65, 16)
				(_nw249).ArraySet1(_1690_v65, 17)
				(_nw249).ArraySet1(_dafny.SeqOf(_dafny.IntOfInt64(-584)), 18)
				(_nw249).ArraySet1(_dafny.SeqOf(_1688_cf47, p2), 19)
				(_nw249).ArraySet1(_1690_v65, 20)
				(_nw249).ArraySet1(_1690_v65, 21)
				(_nw249).ArraySet1(_1690_v65, 22)
				(_nw249).ArraySet1(_dafny.Companion_Sequence_.Update((_1698_v71).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1698_v71).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_1688_cf47, _dafny.IntOfUint32(((_1698_v71).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1698_v71).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _1688_cf47), 23)
				(_nw249).ArraySet1(_1690_v65, 24)
				(_nw249).ArraySet1(_1690_v65, 25)
				(_nw249).ArraySet1(_dafny.Companion_Sequence_.Update(_1690_v65, (Companion_Default___.SafeIndex(_1689_cf46, _dafny.IntOfUint32((_1690_v65).Cardinality()))).Uint32(), (_1699_v72).Cardinality()), 26)
				(_nw249).ArraySet1(_dafny.SeqOf((_this).F6(), p2), 27)
				_1700_v73 = _nw249
				var _1703_v74 _dafny.Map
				_ = _1703_v74
				_1703_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1700_v73, _1692_v67)
				_1703_v74 = (_1703_v74).Update(_1700_v73, _1692_v67)
				var _1704_v75 _dafny.Map
				_ = _1704_v75
				_1704_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1677_v62, _dafny.IntOfInt64(-779))
				_1704_v75 = (_1704_v75).Merge(_1704_v75)
				(globalState).F0 = p0
			} else {
				var _1705_v76 _dafny.Map
				_ = _1705_v76
				_1705_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
				var _1706_v77 D4
				_ = _1706_v77
				_1706_v77 = Companion_D4_.Create_DC15_((func() bool {
					if (_1705_v76).Contains((_dafny.Zero).Minus(_1687_cf48)) {
						return (_1705_v76).Get((_dafny.Zero).Minus(_1687_cf48)).(bool)
					}
					return p0
				})(), p0, _1692_v67)
				r1 = !((_1706_v77).Dtor_cf33())
				var _1707_v78 _dafny.Array
				_ = _1707_v78
				var _nwElement0_47 bool = true
				_ = _nwElement0_47
				var _nw250 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(15))
				_ = _nw250
				(_nw250).ArraySet1(_nwElement0_47, 0)
				(_nw250).ArraySet1(false, 1)
				(_nw250).ArraySet1(p0, 2)
				(_nw250).ArraySet1(!(p0), 3)
				(_nw250).ArraySet1(!((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)), 4)
				(_nw250).ArraySet1(Companion_Default___.Fm3(globalState), 5)
				(_nw250).ArraySet1(p0, 6)
				(_nw250).ArraySet1((_1706_v77).Dtor_cf33(), 7)
				(_nw250).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 8)
				(_nw250).ArraySet1(p0, 9)
				(_nw250).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 10)
				(_nw250).ArraySet1(p0, 11)
				(_nw250).ArraySet1(p0, 12)
				(_nw250).ArraySet1(p0, 13)
				(_nw250).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 14)
				_1707_v78 = _nw250
				var _1708_v79 *C3
				_ = _1708_v79
				var _nw251 *C3 = New_C3_()
				_ = _nw251
				_nw251.Ctor__((_1675_v60).Merge(_1675_v60), _1707_v78, (_1693_v68).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("thytt"), (Companion_Default___.SafeIndex(_1689_cf46, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("thytt")).Cardinality()))).Uint32(), _1686_cf49))), _1687_cf48)
				_1708_v79 = _nw251
				(globalState).F5 = _1688_cf47
				var _1709_v80 _dafny.Map
				_ = _1709_v80
				_1709_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), _1686_cf49)
				var _1710_v81 _dafny.Set
				_ = _1710_v81
				_1710_v81 = _dafny.SetOf(_1688_cf47, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(752))).Uint32(), func(coer107 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg107 _dafny.Int) interface{} {
						return coer107(arg107)
					}
				}((func(_1711_cf46 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1712_i10 _dafny.Int) _dafny.Int {
						return _1711_cf46
					}
				})(_1689_cf46)))).Cardinality())))
				_1686_cf49 = Companion_Default___.Fm0((_1709_v80).Cardinality(), _1710_v81, (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), globalState)
				var _1713_v82 _dafny.MultiSet
				_ = _1713_v82
				_1713_v82 = _dafny.MultiSetOf((_1691_v66).Update(_dafny.IntOfInt64(772), Companion_Default___.Abs((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(734))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg108 _dafny.Int) interface{} {
						return coer108(arg108)
					}
				}((func(_1714_cf49 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1715_i11 _dafny.Int) _dafny.CodePoint {
						return _1714_cf49
					}
				})(_1686_cf49)))).Cardinality())))), _1691_v66)
				var _1716_v83 T0
				_ = _1716_v83
				var _nw252 *C2 = New_C2_()
				_ = _nw252
				_nw252.Ctor__(_1689_cf46, _1713_v82, _1707_v78, _1693_v68)
				_1716_v83 = _nw252
				var _1717_v84 _dafny.Sequence
				_ = _1717_v84
				_1717_v84 = _dafny.SeqOf(_1716_v83, _1716_v83, _1716_v83)
				var _1718_v85 T1
				_ = _1718_v85
				var _nw253 *C3 = New_C3_()
				_ = _nw253
				_nw253.Ctor__((_1675_v60).Merge(_1675_v60), (func() _dafny.Array {
					if p0 {
						return _1707_v78
					}
					return _1707_v78
				})(), (func() _dafny.Map {
					if false {
						return (_this).F10()
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1692_v67).Cardinality()), _dafny.Companion_Sequence_.Update(_1692_v67, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1717_v84).Cardinality()), _dafny.IntOfUint32((_1692_v67).Cardinality()))).Uint32(), _dafny.CodePoint('j')))
				})(), (_dafny.Zero).Minus((_1673_v58).Cardinality()))
				_1718_v85 = _nw253
				var _nw254 *C2 = New_C2_()
				_ = _nw254
				_nw254.Ctor__((_1716_v83).F6(), _1713_v82, _1718_v85.F9(), ((_this).F10()).Merge((_1718_v85).F10()))
				_1718_v85 = _nw254
			}
			var _arr50 _dafny.Array = _this.F9()
			_ = _arr50
			var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index240
			(_arr50).ArraySet1(((func() _dafny.Int {
				if Companion_Default___.Fm3(globalState) {
					return _dafny.IntOfInt64(93)
				}
				return p2
			})()).Cmp(_1688_cf47) != 0, (_index240).Int())
			var _arr51 _dafny.Array = _this.F9()
			_ = _arr51
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index241
			(_arr51).ArraySet1(!(!(((_dafny.IntOfUint32((p1).Cardinality())).Minus(_1687_cf48)).Cmp((_1673_v58).Cardinality()) == 0)), (_index241).Int())
		} else if _source20.Is_DC26() {
			var _1719___mcc_h4 _dafny.Array = _source20.Get_().(D9_DC26).Cf50
			_ = _1719___mcc_h4
			var _1720_cf50 _dafny.Array = _1719___mcc_h4
			_ = _1720_cf50
			var _1721_v86 _dafny.Sequence
			_ = _1721_v86
			_1721_v86 = _dafny.SeqOf(_1677_v62)
			var _1722_v87 _dafny.Set
			_ = _1722_v87
			_1722_v87 = _dafny.SetOf(p0, true)
			var _1723_v88 D11
			_ = _1723_v88
			_1723_v88 = Companion_D11_.Create_DC33_(p2, p0, _dafny.IntOfUint32((_1721_v86).Cardinality()), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), _1722_v87)
			if (_1723_v88).Dtor_cf65() {
				(globalState).F5 = (_this).F6()
				var _1724_v89 _dafny.Map
				_ = _1724_v89
				_1724_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1721_v86).Cardinality()), _1677_v62)
				var _1725_v90 _dafny.Map
				_ = _1725_v90
				_1725_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1724_v89).Cardinality(), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
				var _1726_v91 _dafny.Map
				_ = _1726_v91
				_1726_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
				_1725_v90 = (_1725_v90).Update(p2, (_1726_v91).Contains((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)))
				(globalState).F0 = (p0) && (!_dafny.Companion_Sequence_.Contains(p1, (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)))
				var _rhs281 _dafny.Int = p2
				_ = _rhs281
				var _rhs282 _dafny.Int = ((Companion_D11_.Create_DC31_(_dafny.IntOfInt64(-330), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))).Cardinality())).Dtor_cf59()).Times((_dafny.Zero).Minus((_1722_v87).Cardinality()))
				_ = _rhs282
				var _lhs234 *GlobalState = globalState
				_ = _lhs234
				var _lhs235 *GlobalState = globalState
				_ = _lhs235
				_lhs234.F5 = _rhs281
				_lhs235.F5 = _rhs282
				(globalState).F5 = (_this).F6()
			} else {
				var _1727_v92 _dafny.Array
				_ = _1727_v92
				var _nw255 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(26))
				_ = _nw255
				_1727_v92 = _nw255
				var _1728_v93 D2
				_ = _1728_v93
				_1728_v93 = Companion_D2_.Create_DC8_((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
				var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_1727_v92), 0))
				_ = _index242
				(_1727_v92).ArraySet1(_1728_v93, (_index242).Int())
				var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_1727_v92), 0))
				_ = _index243
				(_1727_v92).ArraySet1((func() D2 {
					if p0 {
						return _1728_v93
					}
					return _1728_v93
				})(), (_index243).Int())
				(globalState).F5 = (p2).Plus((_this).F6())
				var _1729_v94 *C1
				_ = _1729_v94
				var _nw256 *C1 = New_C1_()
				_ = _nw256
				_nw256.Ctor__(_dafny.IntOfUint32((_1721_v86).Cardinality()))
				_1729_v94 = _nw256
				var _1730_v95 _dafny.Array
				_ = _1730_v95
				var _nw257 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
				_ = _nw257
				_1730_v95 = _nw257
				var _arr52 _dafny.Array = _this.F11
				_ = _arr52
				var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_this.F11), 0))
				_ = _index244
				(_arr52).ArraySet1((_dafny.Zero).Minus(p2), (_index244).Int())
				var _1731_v96 _dafny.MultiSet
				_ = _1731_v96
				_1731_v96 = _dafny.MultiSetOf(p1, _dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1721_v86, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1721_v86).Cardinality()))).Uint32(), _1677_v62)).Cardinality()), _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)), p1, p1)
				var _arr53 _dafny.Array = _this.F11
				_ = _arr53
				var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_this.F11), 0))
				_ = _index245
				var _rhs283 _dafny.Array = _1730_v95
				_ = _rhs283
				var _rhs284 _dafny.Int = ((func() _dafny.MultiSet {
					if Companion_Default___.Fm3(globalState) {
						return (_1731_v96).Update(p1, Companion_Default___.Abs(_dafny.IntOfInt64(569)))
					}
					return _1731_v96
				})()).Cardinality()
				_ = _rhs284
				var _lhs236 _dafny.Array = _this.F11
				_ = _lhs236
				var _lhs237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_this.F11), 0))
				_ = _lhs237
				_1730_v95 = _rhs283
				(_lhs236).ArraySet1(_rhs284, (_lhs237).Int())
				var _1732_v97 _dafny.Sequence
				_ = _1732_v97
				_1732_v97 = _dafny.SeqOf(_this.F12)
				var _1733_v98 _dafny.Sequence
				_ = _1733_v98
				_1733_v98 = _dafny.SeqOf((_this).F6(), p2)
				var _1734_v99 _dafny.Array
				_ = _1734_v99
				var _nwElement0_48 _dafny.Array = _this.F12
				_ = _nwElement0_48
				var _nw258 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(9))
				_ = _nw258
				(_nw258).ArraySet1(_nwElement0_48, 0)
				(_nw258).ArraySet1(_this.F12, 1)
				(_nw258).ArraySet1(_this.F12, 2)
				(_nw258).ArraySet1(_this.F12, 3)
				(_nw258).ArraySet1(_this.F12, 4)
				(_nw258).ArraySet1(_this.F12, 5)
				(_nw258).ArraySet1(_this.F12, 6)
				(_nw258).ArraySet1(_this.F12, 7)
				(_nw258).ArraySet1((_1732_v97).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_1733_v98).Cardinality())))), _dafny.IntOfUint32((_1732_v97).Cardinality()))).Uint32()).(_dafny.Array), 8)
				_1734_v99 = _nw258
				var _1735_v101 _dafny.Set
				_ = _1735_v101
				_1735_v101 = _dafny.SetOf((_this.F11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_this.F11), 0))).Int()).(_dafny.Int))
				var _1736_v102 _dafny.Sequence
				_ = _1736_v102
				_1736_v102 = _dafny.SeqOf((func() _dafny.Set {
					var _coll47 = _dafny.NewBuilder()
					_ = _coll47
					for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-932), _dafny.IntOfInt64(-282))); ; {
						_compr_47, _ok49 := _iter49()
						if !_ok49 {
							break
						}
						var _1737_v100 _dafny.Int
						_1737_v100 = interface{}(_compr_47).(_dafny.Int)
						if ((_dafny.IntOfInt64(-932)).Cmp(_1737_v100) <= 0) && ((_1737_v100).Cmp(_dafny.IntOfInt64(-282)) < 0) {
							_coll47.Add(Companion_Default___.SafeModuloInt(_1737_v100, p2))
						}
					}
					return _coll47.ToSet()
				}()).IsProperSubsetOf(_1735_v101), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), ((_this).F6()).Cmp(_dafny.IntOfInt64(-21)) != 0)
				var _1738_v103 _dafny.Array
				_ = _1738_v103
				var _nw259 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
				_ = _nw259
				_1738_v103 = _nw259
				var _arr54 _dafny.Array = _this.F11
				_ = _arr54
				var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_this.F11), 0))
				_ = _index246
				var _rhs285 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("psadyloiw")).Cardinality())
				_ = _rhs285
				var _rhs286 _dafny.Array = _1734_v99
				_ = _rhs286
				var _rhs287 _dafny.Sequence = Companion_Default___.Fm4(p2, (_dafny.IntOfInt64(-846)).Times(_dafny.IntOfInt64(17)), (func() _dafny.Int {
					if p0 {
						return (_this.F11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_this.F11), 0))).Int()).(_dafny.Int)
					}
					return (_this).F6()
				})(), globalState)
				_ = _rhs287
				var _rhs288 _dafny.Array = _1738_v103
				_ = _rhs288
				var _lhs238 _dafny.Array = _this.F11
				_ = _lhs238
				var _lhs239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_this.F11), 0))
				_ = _lhs239
				(_lhs238).ArraySet1(_rhs285, (_lhs239).Int())
				_1734_v99 = _rhs286
				_1736_v102 = _rhs287
				_1738_v103 = _rhs288
			}
			var _1739_v104 _dafny.MultiSet
			_ = _1739_v104
			_1739_v104 = _dafny.MultiSetOf(_dafny.IntOfInt64(105))
			var _1740_v105 _dafny.MultiSet
			_ = _1740_v105
			_1740_v105 = _dafny.MultiSetOf(_1739_v104, _1739_v104)
			var _1741_v106 _dafny.Array
			_ = _1741_v106
			var _nwElement0_49 bool = p0
			_ = _nwElement0_49
			var _nw260 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(8))
			_ = _nw260
			(_nw260).ArraySet1(_nwElement0_49, 0)
			(_nw260).ArraySet1(p0, 1)
			(_nw260).ArraySet1(p0, 2)
			(_nw260).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 3)
			(_nw260).ArraySet1(false, 4)
			(_nw260).ArraySet1(p0, 5)
			(_nw260).ArraySet1(p0, 6)
			(_nw260).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 7)
			_1741_v106 = _nw260
			var _1742_v107 T0
			_ = _1742_v107
			var _nw261 *C2 = New_C2_()
			_ = _nw261
			_nw261.Ctor__((_this).F6(), _1740_v105, _1741_v106, (_this).F10())
			_1742_v107 = _nw261
			var _1743_v108 _dafny.Set
			_ = _1743_v108
			_1743_v108 = _dafny.SetOf(_1742_v107, _1742_v107, _1742_v107)
			var _1744_v109 _dafny.Map
			_ = _1744_v109
			_1744_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1743_v108, _dafny.CodePoint('g'))
			_1744_v109 = (func() _dafny.Map {
				if Companion_Default___.Fm3(globalState) {
					return _1744_v109
				}
				return _1744_v109
			})()
			(globalState).F5 = p2
			var _arr55 _dafny.Array = _this.F9()
			_ = _arr55
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index247
			(_arr55).ArraySet1(!(p0), (_index247).Int())
		} else if _source20.Is_DC24() {
			var _1745___mcc_h5 _dafny.MultiSet = _source20.Get_().(D9_DC24).Cf45
			_ = _1745___mcc_h5
			var _1746_cf45 _dafny.MultiSet = _1745___mcc_h5
			_ = _1746_cf45
			var _arr56 _dafny.Array = _this.F9()
			_ = _arr56
			var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))
			_ = _index248
			(_arr56).ArraySet1(p0, (_index248).Int())
			var _1747_v110 _dafny.Set
			_ = _1747_v110
			_1747_v110 = _dafny.SetOf(p0, p0, p0)
			var _1748_v111 _dafny.Array
			_ = _1748_v111
			var _nwElement0_50 bool = false
			_ = _nwElement0_50
			var _nw262 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(6))
			_ = _nw262
			(_nw262).ArraySet1(_nwElement0_50, 0)
			(_nw262).ArraySet1((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), 1)
			(_nw262).ArraySet1(p0, 2)
			(_nw262).ArraySet1(!((_1747_v110).IsProperSubsetOf(_1747_v110)), 3)
			(_nw262).ArraySet1(false, 4)
			(_nw262).ArraySet1((_dafny.IntOfUint32((p1).Cardinality())).Cmp((_this).F6()) == 0, 5)
			_1748_v111 = _nw262
			var _1749_v112 _dafny.Sequence
			_ = _1749_v112
			_1749_v112 = _dafny.SeqOf((func() _dafny.Array {
				if p0 {
					return _1748_v111
				}
				return _1748_v111
			})())
			_1748_v111 = (_1749_v112).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1749_v112).Cardinality()))).Uint32()).(_dafny.Array)
			(globalState).F5 = (_this).F6()
			if ((_this).F6()).Cmp((_this).F6()) <= 0 {
				var _1750_v113 *C0
				_ = _1750_v113
				var _nw263 *C0 = New_C0_()
				_ = _nw263
				_nw263.Ctor__(_1748_v111, (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
				_1750_v113 = _nw263
				var _1751_v114 _dafny.Sequence
				_ = _1751_v114
				_1751_v114 = _dafny.UnicodeSeqOfUtf8Bytes("hqae")
				var _1752_v115 D9
				_ = _1752_v115
				_1752_v115 = Companion_D9_.Create_DC25_(p2, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1751_v114).Cardinality()), p2), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(125))).Uint32(), func(coer109 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
					return func(arg109 _dafny.Int) interface{} {
						return coer109(arg109)
					}
				}((func(_1753_p2 _dafny.Int) func(_dafny.Int) D3 {
					return func(_1754_i12 _dafny.Int) D3 {
						return Companion_D3_.Create_DC9_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(266))).Uint32(), func(coer110 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg110 _dafny.Int) interface{} {
								return coer110(arg110)
							}
						}((func(_1755_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1756_i13 _dafny.Int) _dafny.Int {
								return _1755_p2
							}
						})(_1753_p2))))
					}
				})(p2)))).Cardinality())), _1677_v62)
				_1752_v115 = _1752_v115
				var _1757_v116 _dafny.Sequence
				_ = _1757_v116
				_1757_v116 = _dafny.SeqOf(_dafny.SetOf(p0))
				var _1758_v117 _dafny.Map
				_ = _1758_v117
				_1758_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1757_v116).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1751_v114).Cardinality()), _dafny.IntOfUint32((_1757_v116).Cardinality()))).Uint32()).(_dafny.Set)).Union(_1747_v110), !((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)))
				_1758_v117 = (_1758_v117).Update(_1747_v110, _1750_v113.F20)
				var _1759_v118 _dafny.Map
				_ = _1759_v118
				_1759_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
				_1759_v118 = (_1759_v118).Merge((_1759_v118).Merge(func() _dafny.Map {
					var _coll48 = _dafny.NewMapBuilder()
					_ = _coll48
					for _iter50 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1675_v60).Cardinality())).Keys().Elements()); ; {
						_compr_48, _ok50 := _iter50()
						if !_ok50 {
							break
						}
						var _1760_v119 _dafny.Int
						_1760_v119 = interface{}(_compr_48).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1675_v60).Cardinality())).Contains(_1760_v119) {
							_coll48.Add(Companion_Default___.SafeDivisionInt(_1760_v119, p2), p0)
						}
					}
					return _coll48.ToMap()
				}()))
				var _1761_v120 _dafny.Sequence
				_ = _1761_v120
				_1761_v120 = _dafny.SeqOf(p2)
				var _1762_v121 _dafny.Map
				_ = _1762_v121
				_1762_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v113.F20, _1750_v113.F20)
				(globalState).F5 = (((_1761_v120).Select((Companion_Default___.SafeIndex((_1762_v121).Cardinality(), _dafny.IntOfUint32((_1761_v120).Cardinality()))).Uint32()).(_dafny.Int)).Times((_this).F6())).Times(p2)
			} else {
				var _1763_v122 _dafny.MultiSet
				_ = _1763_v122
				_1763_v122 = _dafny.MultiSetOf(p2, p2)
				var _arr57 _dafny.Array = _this.F12
				_ = _arr57
				var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_this.F12), 0))
				_ = _index249
				(_arr57).ArraySet1(((func() _dafny.Int {
					if (_1763_v122).Contains(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(998))).Uint32(), func(coer111 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg111 _dafny.Int) interface{} {
							return coer111(arg111)
						}
					}(func(_1764_i14 _dafny.Int) _dafny.Int {
						return (_this).F6()
					}))).Cardinality())) {
						return (_1763_v122).Multiplicity(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(998))).Uint32(), func(coer112 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg112 _dafny.Int) interface{} {
								return coer112(arg112)
							}
						}(func(_1765_i14 _dafny.Int) _dafny.Int {
							return (_this).F6()
						}))).Cardinality()))
					}
					return p2
				})()).Minus(p2), (_index249).Int())
				var _arr58 _dafny.Array = _this.F12
				_ = _arr58
				var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_this.F12), 0))
				_ = _index250
				(_arr58).ArraySet1((_this).F6(), (_index250).Int())
				var _1766_v123 *C6
				_ = _1766_v123
				var _nw264 *C6 = New_C6_()
				_ = _nw264
				_nw264.Ctor__((_this).F6(), (func() bool {
					if (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool) {
						return (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)
					}
					return (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool)
				})())
				_1766_v123 = _nw264
				var _1767_v124 _dafny.Array
				_ = _1767_v124
				var _nw265 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(26))
				_ = _nw265
				_1767_v124 = _nw265
				var _1768_v125 _dafny.Array
				_ = _1768_v125
				_1768_v125 = _1767_v124
				_1767_v124 = (_1768_v125)
				var _1769_v126 _dafny.Map
				_ = _1769_v126
				_1769_v126 = (_this).F10()
				var _1770_v127 *C3
				_ = _1770_v127
				var _nw266 *C3 = New_C3_()
				_ = _nw266
				_nw266.Ctor__(_1675_v60, _1748_v111, (_1769_v126).Merge((_this).F10()), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if ((_this).F10()).Contains((_1766_v123).F13()) {
						return ((_this).F10()).Get((_1766_v123).F13()).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(646))).Uint32(), func(coer113 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg113 _dafny.Int) interface{} {
							return coer113(arg113)
						}
					}((func(_1771_v62 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1772_i15 _dafny.Int) _dafny.CodePoint {
							return _1771_v62
						}
					})(_1677_v62)))
				})()).Cardinality()))
				_1770_v127 = _nw266
				var _1773_v128 *C5
				_ = _1773_v128
				var _nw267 *C5 = New_C5_()
				_ = _nw267
				_nw267.Ctor__((_this.F12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_this.F12), 0))).Int()).(_dafny.Int), (_1766_v123).F13())
				_1773_v128 = _nw267
			}
		} else {
			var _1774___mcc_h6 D9 = _source20.Get_().(D9_DC27).Cf51
			_ = _1774___mcc_h6
			var _1775_cf51 D9 = _1774___mcc_h6
			_ = _1775_cf51
			var _1776_v129 _dafny.Array
			_ = _1776_v129
			var _len0_45 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_45
			var _nw268 _dafny.Array
			_ = _nw268
			if _len0_45.Cmp(_dafny.Zero) == 0 {
				_nw268 = _dafny.NewArray(_len0_45)
			} else {
				var _init45 func(_dafny.Int) D0 = (func(_1777_p1 _dafny.Sequence, _1778_p2 _dafny.Int) func(_dafny.Int) D0 {
					return func(_1779_i16 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_(_1777_p1, _1778_p2)
					}
				})(p1, p2)
				_ = _init45
				var _element0_45 = _init45(_dafny.Zero)
				_ = _element0_45
				_nw268 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
				(_nw268).ArraySet1(_element0_45, 0)
				var _nativeLen0_45 = (_len0_45).Int()
				_ = _nativeLen0_45
				for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
					(_nw268).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
				}
			}
			_1776_v129 = _nw268
			var _1780_v130 D7
			_ = _1780_v130
			_1780_v130 = Companion_D7_.Create_DC21_(_1776_v129)
			_1780_v130 = _1780_v130
			var _1781_v131 _dafny.Sequence
			_ = _1781_v131
			_1781_v131 = _dafny.UnicodeSeqOfUtf8Bytes("dcbeliug")
			var _1782_v132 *C5
			_ = _1782_v132
			var _nw269 *C5 = New_C5_()
			_ = _nw269
			_nw269.Ctor__((_dafny.Zero).Minus(_dafny.IntOfUint32((_1781_v131).Cardinality())), _dafny.IntOfInt64(510))
			_1782_v132 = _nw269
			var _1783_v133 _dafny.Map
			_ = _1783_v133
			_1783_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1782_v132, _this.F12)
			var _rhs289 _dafny.Int = (p2).Times((_this).F6())
			_ = _rhs289
			var _rhs290 _dafny.Int = (_this).F6()
			_ = _rhs290
			var _rhs291 _dafny.Map = (_1783_v133).Merge(_1783_v133)
			_ = _rhs291
			var _lhs240 *GlobalState = globalState
			_ = _lhs240
			var _lhs241 *GlobalState = globalState
			_ = _lhs241
			_lhs240.F5 = _rhs289
			_lhs241.F5 = _rhs290
			_1783_v133 = _rhs291
			var _1784_v134 _dafny.Map
			_ = _1784_v134
			_1784_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (func() bool {
				if (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool) {
					return !((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
				}
				return p0
			})())
			_1784_v134 = (_1784_v134).Update((_1782_v132).F15(), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
			var _arr59 _dafny.Array = _this.F11
			_ = _arr59
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_this.F11), 0))
			_ = _index251
			(_arr59).ArraySet1((_1782_v132).F15(), (_index251).Int())
			var _1785_v135 _dafny.Sequence
			_ = _1785_v135
			_1785_v135 = _dafny.SeqOf(_this.F12, _this.F11)
			var _1786_v136 _dafny.Array
			_ = _1786_v136
			var _nw270 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw270
			_1786_v136 = _nw270
			var _1787_v137 _dafny.Sequence
			_ = _1787_v137
			_1787_v137 = _dafny.SeqOf((_1782_v132).F15(), (_this).F6(), (_dafny.Zero).Minus((_1782_v132).F15()))
			var _1788_v138 _dafny.MultiSet
			_ = _1788_v138
			_1788_v138 = _dafny.MultiSetOf(_1787_v137, _1787_v137)
			var _1789_v139 _dafny.Sequence
			_ = _1789_v139
			_1789_v139 = _dafny.SeqOf(_1787_v137)
			var _arr60 _dafny.Array = _this.F11
			_ = _arr60
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_this.F11), 0))
			_ = _index252
			var _rhs292 _dafny.Int = (_dafny.Zero).Minus((_1782_v132).F15())
			_ = _rhs292
			var _rhs293 _dafny.Sequence = (func() _dafny.Sequence {
				if p0 {
					return _1781_v131
				}
				return _dafny.Companion_Sequence_.Update(_1781_v131, (Companion_Default___.SafeIndex((_this).Fm14(globalState), _dafny.IntOfUint32((_1781_v131).Cardinality()))).Uint32(), _1677_v62)
			})()
			_ = _rhs293
			var _rhs294 _dafny.Sequence = (func() _dafny.Sequence {
				if (_dafny.MultiSetFromSeq(_1789_v139)).IsSubsetOf(_1788_v138) {
					return _1785_v135
				}
				return _1785_v135
			})()
			_ = _rhs294
			var _rhs295 _dafny.Int = (_1782_v132).F15()
			_ = _rhs295
			var _rhs296 _dafny.Array = _1786_v136
			_ = _rhs296
			var _lhs242 _dafny.Array = _this.F11
			_ = _lhs242
			var _lhs243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_this.F11), 0))
			_ = _lhs243
			var _lhs244 *GlobalState = globalState
			_ = _lhs244
			var _lhs245 *GlobalState = globalState
			_ = _lhs245
			(_lhs242).ArraySet1(_rhs292, (_lhs243).Int())
			_lhs244.F2 = _rhs293
			_1785_v135 = _rhs294
			_lhs245.F5 = _rhs295
			_1786_v136 = _rhs296
		}
		r0 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(780))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg114 _dafny.Int) interface{} {
				return coer114(arg114)
			}
		}(func(_1790_i17 _dafny.Int) _dafny.Sequence {
			return _dafny.UnicodeSeqOfUtf8Bytes("qvp")
		}))
		var _1791_v140 _dafny.Set
		_ = _1791_v140
		_1791_v140 = _dafny.SetOf((_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool))
		var _1792_v141 _dafny.MultiSet
		_ = _1792_v141
		_1792_v141 = _dafny.MultiSetOf(_dafny.IntOfInt64(976), (_1791_v140).Cardinality())
		var _1793_v142 D11
		_ = _1793_v142
		_1793_v142 = Companion_D11_.Create_DC31_((_this).F6(), (_this).F6())
		var _1794_v143 _dafny.Map
		_ = _1794_v143
		_1794_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1792_v141, (p1).Select((Companion_Default___.SafeIndex((_1793_v142).Dtor_cf58(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool))
		var _1795_v144 D5
		_ = _1795_v144
		_1795_v144 = Companion_D5_.Create_DC16_(_1791_v140)
		var _1796_v145 D11
		_ = _1796_v145
		_1796_v145 = Companion_D11_.Create_DC33_((_1794_v143).Cardinality(), p0, (_this).F6(), (_this.F9()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_this.F9()), 0))).Int()).(bool), (_1795_v144).Dtor_cf35())
		r1 = (_1796_v145).Dtor_cf63()
		return r0, r1
	}
}
func (_this *C7) M7(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) (_dafny.CodePoint, _dafny.Int, _dafny.Array, _dafny.Int) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		(globalState).F5 = (_this).F6()
		var _1797_v0 bool
		_ = _1797_v0
		_1797_v0 = false
		var _1798_v1 _dafny.Sequence
		_ = _1798_v1
		_1798_v1 = _dafny.UnicodeSeqOfUtf8Bytes("ldvfwsx")
		(globalState).F0 = (func() bool {
			if _1797_v0 {
				return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("gsjgptuvl"), _1798_v1)
			}
			return _1797_v0
		})()
		(globalState).F0 = Companion_Default___.Fm3(globalState)
		_1797_v0 = !(_1797_v0)
		var _1799_v2 D7
		_ = _1799_v2
		_1799_v2 = Companion_D7_.Create_DC22_()
		var _1800_v3 _dafny.Sequence
		_ = _1800_v3
		_1800_v3 = _dafny.SeqOf(_1797_v0)
		var _1801_v4 _dafny.Map
		_ = _1801_v4
		_1801_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1797_v0), _1797_v0)
		var _1802_v5 D1
		_ = _1802_v5
		_1802_v5 = Companion_D1_.Create_DC4_(Companion_D0_.Create_DC0_(_1800_v3, (_1801_v4).Cardinality()), Companion_Default___.Fm2((_this).F6(), globalState), Companion_Default___.Fm18(_1797_v0, globalState), _1797_v0, (_dafny.Zero).Minus((_this).F6()))
		var _rhs297 bool = (_1802_v5).Dtor_cf10()
		_ = _rhs297
		var _rhs298 D7 = _1799_v2
		_ = _rhs298
		var _rhs299 bool = _1797_v0
		_ = _rhs299
		var _rhs300 bool = Companion_Default___.Fm3(globalState)
		_ = _rhs300
		var _lhs246 *GlobalState = globalState
		_ = _lhs246
		var _lhs247 *GlobalState = globalState
		_ = _lhs247
		var _lhs248 *GlobalState = globalState
		_ = _lhs248
		_lhs246.F0 = _rhs297
		_1799_v2 = _rhs298
		_lhs247.F0 = _rhs299
		_lhs248.F0 = _rhs300
		var _1803_v6 _dafny.CodePoint
		_ = _1803_v6
		_1803_v6 = _dafny.CodePoint('l')
		var _1804_v7 _dafny.Map
		_ = _1804_v7
		_1804_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(287), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ljwkjfnhn"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ljwkjfnhn")).Cardinality()))).Uint32(), _1803_v6)).Cardinality())))
		var _1805_v9 _dafny.Map
		_ = _1805_v9
		_1805_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.SetOf(_1803_v6))
		var _1806_v10 *C3
		_ = _1806_v10
		var _nw271 *C3 = New_C3_()
		_ = _nw271
		_nw271.Ctor__(_1804_v7, _this.F9(), func() _dafny.Map {
			var _coll49 = _dafny.NewMapBuilder()
			_ = _coll49
			for _iter51 := _dafny.Iterate((_1805_v9).Keys().Elements()); ; {
				_compr_49, _ok51 := _iter51()
				if !_ok51 {
					break
				}
				var _1807_v8 _dafny.Int
				_1807_v8 = interface{}(_compr_49).(_dafny.Int)
				if (_1805_v9).Contains(_1807_v8) {
					_coll49.Add((_1807_v8).Plus(p1), _dafny.UnicodeSeqOfUtf8Bytes("f"))
				}
			}
			return _coll49.ToMap()
		}(), p1)
		_1806_v10 = _nw271
		r0 = _1803_v6
		var _1808_v11 _dafny.MultiSet
		_ = _1808_v11
		_1808_v11 = _dafny.MultiSetOf(_1797_v0)
		r1 = (func() _dafny.Int {
			if _1797_v0 {
				return Companion_Default___.SafeModuloInt((_this).F6(), (_this).F6())
			}
			return ((_this).F6()).Plus((_dafny.Zero).Minus((func() _dafny.Int {
				if (_1808_v11).Contains(false) {
					return (_1808_v11).Multiplicity(false)
				}
				return _dafny.IntOfInt64(788)
			})()))
		})()
		var _nw272 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
		_ = _nw272
		r2 = _nw272
		r3 = p1
		return r0, r1, r2, r3
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	dummy byte
}

func New_C8_() *C8 {
	_this := C8{}

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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) Ctor__() {
	{
	}
}
func (_this *C8) Fm10(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetOf(_dafny.IntOfInt64(697), _dafny.IntOfInt64(243), _dafny.IntOfInt64(884), _dafny.IntOfInt64(-12))).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(835)))
	}
}
func (_this *C8) Fm11(globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt((_dafny.IntOfUint32((_dafny.SeqOf(Companion_D1_.Create_DC4_(Companion_D0_.Create_DC0_(_dafny.SeqOf(!(true), false), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("skr")).Cardinality())), _dafny.IntOfInt64(887), _dafny.UnicodeSeqOfUtf8Bytes("lqerk"), false, _dafny.IntOfInt64(739)))).Cardinality())).Minus(_dafny.IntOfInt64(-290)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ctlfngak"), _dafny.UnicodeSeqOfUtf8Bytes("snnrnqplc"))).Cardinality()))
	}
}
func (_this *C8) M4(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) (_dafny.Map, _dafny.Map, _dafny.Int) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1809_v0 _dafny.Map
		_ = _1809_v0
		_1809_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)
		var _1810_v1 _dafny.Array
		_ = _1810_v1
		var _nwElement0_51 bool = p1
		_ = _nwElement0_51
		var _nw273 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(5))
		_ = _nw273
		(_nw273).ArraySet1(_nwElement0_51, 0)
		(_nw273).ArraySet1(Companion_Default___.Fm3(globalState), 1)
		(_nw273).ArraySet1((p0) && (p3), 2)
		(_nw273).ArraySet1(p0, 3)
		(_nw273).ArraySet1((func() bool {
			if !((func() bool {
				if (_1809_v0).Contains(p3) {
					return (_1809_v0).Get(p3).(bool)
				}
				return Companion_Default___.Fm3(globalState)
			})()) {
				return (func() bool {
					if (_1809_v0).Contains(p0) {
						return (_1809_v0).Get(p0).(bool)
					}
					return p0
				})()
			}
			return p3
		})(), 4)
		_1810_v1 = _nw273
		for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1810_v1), 0))); ; {
			_guard_loop_2, _ok52 := _iter52()
			if !_ok52 {
				break
			}
			var _1811_i0 _dafny.Int
			_1811_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1811_i0).Sign() != -1) && ((_1811_i0).Cmp(_dafny.ArrayLen((_1810_v1), 0)) < 0)) {
				(_1810_v1).ArraySet1((func() bool {
					if p3 {
						return true
					}
					return (p0) == (true)
				})(), (_1811_i0).Int())
			}
		}
		var _1812_v2 _dafny.Sequence
		_ = _1812_v2
		_1812_v2 = _dafny.SeqOf(p2)
		var _1813_v3 _dafny.MultiSet
		_ = _1813_v3
		_1813_v3 = _dafny.MultiSetOf(_1812_v2, _1812_v2)
		if (_1813_v3).IsDisjointFrom(_1813_v3) {
			var _1814_v4 _dafny.CodePoint
			_ = _1814_v4
			_1814_v4 = _dafny.CodePoint('x')
			var _1815_v5 D0
			_ = _1815_v5
			_1815_v5 = Companion_D0_.Create_DC1_(!(!(p0)), _dafny.IntOfInt64(172), _1814_v4)
			var _1816_v6 _dafny.Set
			_ = _1816_v6
			_1816_v6 = _dafny.SetOf((_this).Fm11(globalState), p2, p2)
			var _1817_v8 _dafny.Sequence
			_ = _1817_v8
			_1817_v8 = _dafny.UnicodeSeqOfUtf8Bytes("mjcroura")
			var _1818_v9 _dafny.Map
			_ = _1818_v9
			_1818_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1814_v4, _dafny.IntOfUint32((_1817_v8).Cardinality()))
			var _1819_v10 _dafny.Sequence
			_ = _1819_v10
			_1819_v10 = _dafny.SeqOf(_1818_v9)
			var _rhs301 bool = (_1816_v6).IsDisjointFrom(func() _dafny.Set {
				var _coll50 = _dafny.NewBuilder()
				_ = _coll50
				for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-492), _dafny.IntOfInt64(679))); ; {
					_compr_50, _ok53 := _iter53()
					if !_ok53 {
						break
					}
					var _1820_v7 _dafny.Int
					_1820_v7 = interface{}(_compr_50).(_dafny.Int)
					if ((_dafny.IntOfInt64(-492)).Cmp(_1820_v7) <= 0) && ((_1820_v7).Cmp(_dafny.IntOfInt64(679)) < 0) {
						_coll50.Add((_1820_v7).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_1815_v5).Dtor_cf2(), p1)).Cardinality())))
					}
				}
				return _coll50.ToSet()
			}())
			_ = _rhs301
			var _rhs302 D0 = Companion_D0_.Create_DC1_(p3, p2, (_1815_v5).Dtor_cf4())
			_ = _rhs302
			var _rhs303 bool = ((_1819_v10).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1819_v10).Cardinality()))).Uint32()).(_dafny.Map)).Contains(_dafny.CodePoint('k'))
			_ = _rhs303
			var _rhs304 bool = p0
			_ = _rhs304
			var _lhs249 *GlobalState = globalState
			_ = _lhs249
			var _lhs250 *GlobalState = globalState
			_ = _lhs250
			var _lhs251 *GlobalState = globalState
			_ = _lhs251
			_lhs249.F0 = _rhs301
			_1815_v5 = _rhs302
			_lhs250.F0 = _rhs303
			_lhs251.F0 = _rhs304
			var _1821_v11 _dafny.Array
			_ = _1821_v11
			var _nw274 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw274
			_1821_v11 = _nw274
			var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_1821_v11), 0))
			_ = _index253
			(_1821_v11).ArraySet1(p2, (_index253).Int())
			var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_1821_v11), 0))
			_ = _index254
			(_1821_v11).ArraySet1((p2).Times(_dafny.IntOfInt64(240)), (_index254).Int())
			var _pat_let_tv34 = p3
			_ = _pat_let_tv34
			_1814_v4 = (func(_pat_let47_0 D0) D0 {
				return func(_1822_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let48_0 bool) D0 {
						return func(_1823_dt__update_hcf2_h0 bool) D0 {
							return Companion_D0_.Create_DC1_(_1823_dt__update_hcf2_h0, (_1822_dt__update__tmp_h0).Dtor_cf3(), (_1822_dt__update__tmp_h0).Dtor_cf4())
						}(_pat_let48_0)
					}(_pat_let_tv34)
				}(_pat_let47_0)
			}(Companion_D0_.Create_DC1_(p1, (_1821_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_1821_v11), 0))).Int()).(_dafny.Int), _1814_v4))).Dtor_cf4()
			(globalState).F5 = (_1821_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_1821_v11), 0))).Int()).(_dafny.Int)
			var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_1821_v11), 0))
			_ = _index255
			(_1821_v11).ArraySet1((_this).Fm11(globalState), (_index255).Int())
		} else {
			var _1824_v12 _dafny.Array
			_ = _1824_v12
			var _len0_46 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_46
			var _nw275 _dafny.Array
			_ = _nw275
			if _len0_46.Cmp(_dafny.Zero) == 0 {
				_nw275 = _dafny.NewArray(_len0_46)
			} else {
				var _init46 func(_dafny.Int) _dafny.Sequence = (func(_1825_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1826_i1 _dafny.Int) _dafny.Sequence {
						return _1825_v2
					}
				})(_1812_v2)
				_ = _init46
				var _element0_46 = _init46(_dafny.Zero)
				_ = _element0_46
				_nw275 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
				(_nw275).ArraySet1(_element0_46, 0)
				var _nativeLen0_46 = (_len0_46).Int()
				_ = _nativeLen0_46
				for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
					(_nw275).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
				}
			}
			_1824_v12 = _nw275
			_1824_v12 = (func() _dafny.Array {
				if !(p0) {
					return _1824_v12
				}
				return _1824_v12
			})()
			var _1827_v13 _dafny.Array
			_ = _1827_v13
			var _nw276 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw276
			_1827_v13 = _nw276
			_1827_v13 = _1827_v13
			var _1828_v14 D2
			_ = _1828_v14
			_1828_v14 = Companion_D2_.Create_DC6_(_dafny.IntOfInt64(11), _1827_v13)
			_1828_v14 = _1828_v14
			_1812_v2 = _1812_v2
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_1810_v1), 0))
			_ = _index256
			(_1810_v1).ArraySet1(p1, (_index256).Int())
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_1810_v1), 0))
			_ = _index257
			(_1810_v1).ArraySet1(p0, (_index257).Int())
		}
		(globalState).F2 = _dafny.UnicodeSeqOfUtf8Bytes("hrnahy")
		var _1829_v15 D2
		_ = _1829_v15
		_1829_v15 = Companion_D2_.Create_DC8_(p0, !(false))
		var _pat_let_tv35 = _1809_v0
		_ = _pat_let_tv35
		var _pat_let_tv36 = _1809_v0
		_ = _pat_let_tv36
		var _pat_let_tv37 = p1
		_ = _pat_let_tv37
		var _pat_let_tv38 = _1809_v0
		_ = _pat_let_tv38
		_1809_v0 = func(_source21 D2) _dafny.Map {
			if _source21.Is_DC6() {
				var _1830___mcc_h0 _dafny.Int = _source21.Get_().(D2_DC6).Cf13
				_ = _1830___mcc_h0
				var _1831___mcc_h1 _dafny.Array = _source21.Get_().(D2_DC6).Cf14
				_ = _1831___mcc_h1
				var _1832_cf14 _dafny.Array = _1831___mcc_h1
				_ = _1832_cf14
				var _1833_cf13 _dafny.Int = _1830___mcc_h0
				_ = _1833_cf13
				return _pat_let_tv35
			} else if _source21.Is_DC7() {
				var _1834___mcc_h2 _dafny.Sequence = _source21.Get_().(D2_DC7).Cf15
				_ = _1834___mcc_h2
				var _1835___mcc_h3 _dafny.Array = _source21.Get_().(D2_DC7).Cf16
				_ = _1835___mcc_h3
				var _1836___mcc_h4 _dafny.Int = _source21.Get_().(D2_DC7).Cf17
				_ = _1836___mcc_h4
				var _1837___mcc_h5 _dafny.Int = _source21.Get_().(D2_DC7).Cf18
				_ = _1837___mcc_h5
				var _1838_cf18 _dafny.Int = _1837___mcc_h5
				_ = _1838_cf18
				var _1839_cf17 _dafny.Int = _1836___mcc_h4
				_ = _1839_cf17
				var _1840_cf16 _dafny.Array = _1835___mcc_h3
				_ = _1840_cf16
				var _1841_cf15 _dafny.Sequence = _1834___mcc_h2
				_ = _1841_cf15
				return _pat_let_tv36
			} else if _source21.Is_DC8() {
				var _1842___mcc_h6 bool = _source21.Get_().(D2_DC8).Cf19
				_ = _1842___mcc_h6
				var _1843___mcc_h7 bool = _source21.Get_().(D2_DC8).Cf20
				_ = _1843___mcc_h7
				var _1844_cf20 bool = _1843___mcc_h7
				_ = _1844_cf20
				var _1845_cf19 bool = _1842___mcc_h6
				_ = _1845_cf19
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv37, _1844_cf20)
			} else {
				var _1846___mcc_h8 _dafny.Map = _source21.Get_().(D2_DC5).Cf12
				_ = _1846___mcc_h8
				var _1847_cf12 _dafny.Map = _1846___mcc_h8
				_ = _1847_cf12
				return _pat_let_tv38
			}
		}(_1829_v15)
		var _1848_v16 _dafny.CodePoint
		_ = _1848_v16
		_1848_v16 = _dafny.CodePoint('i')
		var _1849_v17 _dafny.Sequence
		_ = _1849_v17
		_1849_v17 = _dafny.UnicodeSeqOfUtf8Bytes("m")
		var _source22 D2 = Companion_Default___.Fm12(p2, !(_dafny.Companion_Sequence_.Contains(_1849_v17, _1848_v16)), !(!(!(p0))), globalState)
		_ = _source22
		if _source22.Is_DC6() {
			var _1850___mcc_h9 _dafny.Int = _source22.Get_().(D2_DC6).Cf13
			_ = _1850___mcc_h9
			var _1851___mcc_h10 _dafny.Array = _source22.Get_().(D2_DC6).Cf14
			_ = _1851___mcc_h10
			var _1852_cf14 _dafny.Array = _1851___mcc_h10
			_ = _1852_cf14
			var _1853_cf13 _dafny.Int = _1850___mcc_h9
			_ = _1853_cf13
			var _1854_v18 _dafny.Map
			_ = _1854_v18
			_1854_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
			var _1855_v19 _dafny.Map
			_ = _1855_v19
			_1855_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1848_v16, false)
			var _1856_v21 _dafny.MultiSet
			_ = _1856_v21
			_1856_v21 = _dafny.MultiSetOf(_1848_v16)
			var _1857_v22 _dafny.Sequence
			_ = _1857_v22
			_1857_v22 = _dafny.SeqOf(_1855_v19, func() _dafny.Map {
				var _coll51 = _dafny.NewMapBuilder()
				_ = _coll51
				for _iter54 := _dafny.Iterate((_1856_v21).Elements()); ; {
					_compr_51, _ok54 := _iter54()
					if !_ok54 {
						break
					}
					var _1858_v20 _dafny.CodePoint
					_1858_v20 = interface{}(_compr_51).(_dafny.CodePoint)
					if (_1856_v21).Contains(_1858_v20) {
						_coll51.Add(_1858_v20, !(p1))
					}
				}
				return _coll51.ToMap()
			}())
			_1854_v18 = (_1854_v18).Update(p0, _dafny.IntOfUint32((_1857_v22).Cardinality()))
			_1853_cf13 = (Companion_Default___.SafeDivisionInt(p2, _dafny.IntOfInt64(650))).Times(_dafny.IntOfInt64(91))
			var _1859_v23 _dafny.Array
			_ = _1859_v23
			var _nw277 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(16))
			_ = _nw277
			_1859_v23 = _nw277
			var _1860_v24 _dafny.Map
			_ = _1860_v24
			_1860_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1849_v17).Cardinality()), p1)
			var _1861_v25 _dafny.Sequence
			_ = _1861_v25
			_1861_v25 = _dafny.SeqOf(_1860_v24)
			var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(479), _dafny.ArrayLen((_1859_v23), 0))
			_ = _index258
			(_1859_v23).ArraySet1((_1861_v25).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1861_v25).Cardinality()))).Uint32()).(_dafny.Map), (_index258).Int())
			var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(479), _dafny.ArrayLen((_1859_v23), 0))
			_ = _index259
			(_1859_v23).ArraySet1(_1860_v24, (_index259).Int())
			var _1862_v26 _dafny.Set
			_ = _1862_v26
			_1862_v26 = _dafny.SetOf(true, p1)
			if (_dafny.SetOf(p1, !(p3))).IsSubsetOf(_1862_v26) {
				var _1863_v27 _dafny.Map
				_ = _1863_v27
				_1863_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1853_cf13, _1848_v16)
				_1863_v27 = (_1863_v27).Update(_1853_cf13, _1848_v16)
				var _1864_v28 _dafny.Array
				_ = _1864_v28
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_47
				var _nw278 _dafny.Array
				_ = _nw278
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw278 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) _dafny.Sequence = (func(_1865_v2 _dafny.Sequence, _1866_p2 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_1867_i2 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_1865_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-622))).Uint32(), func(coer115 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg115 _dafny.Int) interface{} {
									return coer115(arg115)
								}
							}((func(_1868_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_1869_i3 _dafny.Int) _dafny.Int {
									return _1868_p2
								}
							})(_1866_p2))))
						}
					})(_1812_v2, p2)
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw278 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw278).ArraySet1(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw278).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_1864_v28 = _nw278
				var _rhs305 bool = ((_1853_cf13).Times(_1853_cf13)).Cmp(p2) < 0
				_ = _rhs305
				var _rhs306 _dafny.Int = (_1853_cf13).Minus(_dafny.IntOfUint32((_1849_v17).Cardinality()))
				_ = _rhs306
				var _rhs307 _dafny.Array = _1864_v28
				_ = _rhs307
				var _lhs252 *GlobalState = globalState
				_ = _lhs252
				_lhs252.F0 = _rhs305
				r2 = _rhs306
				_1864_v28 = _rhs307
				(globalState).F5 = (_1853_cf13).Times((_dafny.Zero).Minus(_1853_cf13))
				(globalState).F0 = p3
				var _1870_v29 *C0
				_ = _1870_v29
				var _nw279 *C0 = New_C0_()
				_ = _nw279
				_nw279.Ctor__(_1810_v1, p1)
				_1870_v29 = _nw279
			} else {
				(globalState).F2 = _1849_v17
				(globalState).F0 = p1
				(globalState).F0 = p3
				(globalState).F0 = (func() bool {
					if p3 {
						return p1
					}
					return p0
				})()
				var _1871_v30 *C6
				_ = _1871_v30
				var _nw280 *C6 = New_C6_()
				_ = _nw280
				_nw280.Ctor__((func() _dafny.Int {
					if false {
						return p2
					}
					return p2
				})(), p0)
				_1871_v30 = _nw280
			}
		} else if _source22.Is_DC7() {
			var _1872___mcc_h11 _dafny.Sequence = _source22.Get_().(D2_DC7).Cf15
			_ = _1872___mcc_h11
			var _1873___mcc_h12 _dafny.Array = _source22.Get_().(D2_DC7).Cf16
			_ = _1873___mcc_h12
			var _1874___mcc_h13 _dafny.Int = _source22.Get_().(D2_DC7).Cf17
			_ = _1874___mcc_h13
			var _1875___mcc_h14 _dafny.Int = _source22.Get_().(D2_DC7).Cf18
			_ = _1875___mcc_h14
			var _1876_cf18 _dafny.Int = _1875___mcc_h14
			_ = _1876_cf18
			var _1877_cf17 _dafny.Int = _1874___mcc_h13
			_ = _1877_cf17
			var _1878_cf16 _dafny.Array = _1873___mcc_h12
			_ = _1878_cf16
			var _1879_cf15 _dafny.Sequence = _1872___mcc_h11
			_ = _1879_cf15
			if p0 {
				var _1880_v31 _dafny.Map
				_ = _1880_v31
				_1880_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(537), p2)
				var _1881_v32 *C3
				_ = _1881_v32
				var _nw281 *C3 = New_C3_()
				_ = _nw281
				_nw281.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1849_v17).Cardinality()), p2), _1810_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(426), _1849_v17), _dafny.IntOfInt64(-162))
				_1881_v32 = _nw281
				var _1882_v33 _dafny.Sequence
				_ = _1882_v33
				_1882_v33 = _dafny.SeqOf(_1881_v32, _1881_v32)
				var _1883_v34 _dafny.Map
				_ = _1883_v34
				_1883_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1880_v31).Update(_dafny.IntOfUint32((_1882_v33).Cardinality()), _dafny.IntOfInt64(-92))).Cardinality(), p3)
				var _1884_v35 _dafny.Map
				_ = _1884_v35
				_1884_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("dppy"), _1876_cf18)
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(479), _dafny.ArrayLen((_1810_v1), 0))
				_ = _index260
				(_1810_v1).ArraySet1((func() bool {
					if (_1883_v34).Contains((_1884_v35).Cardinality()) {
						return (_1883_v34).Get((_1884_v35).Cardinality()).(bool)
					}
					return !(p1)
				})(), (_index260).Int())
				var _1885_v36 _dafny.Sequence
				_ = _1885_v36
				_1885_v36 = _dafny.SeqOf(Companion_Default___.Fm51(globalState))
				var _1886_v37 _dafny.Set
				_ = _1886_v37
				_1886_v37 = _dafny.SetOf((_1809_v0).Cardinality())
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(479), _dafny.ArrayLen((_1810_v1), 0))
				_ = _index261
				var _rhs308 _dafny.MultiSet = (_1885_v36).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1885_v36).Cardinality()))).Uint32()).(_dafny.MultiSet)
				_ = _rhs308
				var _rhs309 _dafny.Int = _1876_cf18
				_ = _rhs309
				var _rhs310 bool = (_1886_v37).IsSubsetOf((_1886_v37).Difference(_dafny.SetOf(_1877_cf17)))
				_ = _rhs310
				var _lhs253 *GlobalState = globalState
				_ = _lhs253
				var _lhs254 _dafny.Array = _1810_v1
				_ = _lhs254
				var _lhs255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(479), _dafny.ArrayLen((_1810_v1), 0))
				_ = _lhs255
				_1813_v3 = _rhs308
				_lhs253.F5 = _rhs309
				(_lhs254).ArraySet1(_rhs310, (_lhs255).Int())
				var _1887_v38 *C5
				_ = _1887_v38
				var _nw282 *C5 = New_C5_()
				_ = _nw282
				_nw282.Ctor__(_dafny.IntOfUint32((_1879_cf15).Cardinality()), _1876_cf18)
				_1887_v38 = _nw282
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(479), _dafny.ArrayLen((_1810_v1), 0))
				_ = _index262
				(_1810_v1).ArraySet1(p0, (_index262).Int())
				var _1888_v39 _dafny.Map
				_ = _1888_v39
				_1888_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if p0 {
						return _1812_v2
					}
					return _dafny.SeqOf(_1876_cf18)
				})()).Cardinality()), Companion_Default___.Fm0((_1887_v38).F15(), _dafny.SetOf(_1877_cf17), (_1810_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(479), _dafny.ArrayLen((_1810_v1), 0))).Int()).(bool), globalState))
				_1888_v39 = (_1888_v39).Update((func() _dafny.Int {
					if p1 {
						return (_1887_v38).F15()
					}
					return (_1887_v38).F15()
				})(), _1848_v16)
				var _1889_v40 *C5
				_ = _1889_v40
				var _nw283 *C5 = New_C5_()
				_ = _nw283
				_nw283.Ctor__(_1876_cf18, (_dafny.Zero).Minus((_this).Fm11(globalState)))
				_1889_v40 = _nw283
			} else {
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))
				_ = _index263
				(_1810_v1).ArraySet1(p1, (_index263).Int())
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_1810_v1), 0))
				_ = _index264
				(_1810_v1).ArraySet1(p0, (_index264).Int())
				var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))
				_ = _index265
				var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_1810_v1), 0))
				_ = _index266
				var _rhs311 bool = (p0) && (p1)
				_ = _rhs311
				var _rhs312 bool = p1
				_ = _rhs312
				var _rhs313 bool = !(p1)
				_ = _rhs313
				var _rhs314 bool = p1
				_ = _rhs314
				var _lhs256 *GlobalState = globalState
				_ = _lhs256
				var _lhs257 _dafny.Array = _1810_v1
				_ = _lhs257
				var _lhs258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))
				_ = _lhs258
				var _lhs259 *GlobalState = globalState
				_ = _lhs259
				var _lhs260 _dafny.Array = _1810_v1
				_ = _lhs260
				var _lhs261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_1810_v1), 0))
				_ = _lhs261
				_lhs256.F0 = _rhs311
				(_lhs257).ArraySet1(_rhs312, (_lhs258).Int())
				_lhs259.F0 = _rhs313
				(_lhs260).ArraySet1(_rhs314, (_lhs261).Int())
				var _1890_v41 _dafny.MultiSet
				_ = _1890_v41
				_1890_v41 = _dafny.MultiSetOf(p3, (_1810_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))).Int()).(bool), p3)
				var _1891_v42 *C5
				_ = _1891_v42
				var _nw284 *C5 = New_C5_()
				_ = _nw284
				_nw284.Ctor__((_this).Fm11(globalState), (_dafny.Zero).Minus((_dafny.Zero).Minus(((_1890_v41).Cardinality()).Plus(_dafny.IntOfInt64(-241)))))
				_1891_v42 = _nw284
				var _1892_v43 _dafny.Array
				_ = _1892_v43
				var _nwElement0_52 bool = p0
				_ = _nwElement0_52
				var _nw285 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(28))
				_ = _nw285
				(_nw285).ArraySet1(_nwElement0_52, 0)
				(_nw285).ArraySet1(!(!(p0)), 1)
				(_nw285).ArraySet1(p1, 2)
				(_nw285).ArraySet1((_1810_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))).Int()).(bool), 3)
				(_nw285).ArraySet1(Companion_Default___.Fm3(globalState), 4)
				(_nw285).ArraySet1((_1810_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))).Int()).(bool), 5)
				(_nw285).ArraySet1((_dafny.IntOfInt64(-985)).Cmp(_1877_cf17) <= 0, 6)
				(_nw285).ArraySet1(p3, 7)
				(_nw285).ArraySet1(!(true), 8)
				(_nw285).ArraySet1(p0, 9)
				(_nw285).ArraySet1(p0, 10)
				(_nw285).ArraySet1(p3, 11)
				(_nw285).ArraySet1(true, 12)
				(_nw285).ArraySet1(p3, 13)
				(_nw285).ArraySet1(p0, 14)
				(_nw285).ArraySet1(!(p0), 15)
				(_nw285).ArraySet1(true, 16)
				(_nw285).ArraySet1(!(p3) || (p3), 17)
				(_nw285).ArraySet1((_1810_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))).Int()).(bool), 18)
				(_nw285).ArraySet1((_1810_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))).Int()).(bool), 19)
				(_nw285).ArraySet1((_1810_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))).Int()).(bool), 20)
				(_nw285).ArraySet1(p3, 21)
				(_nw285).ArraySet1(p3, 22)
				(_nw285).ArraySet1(!(p0) || (p1), 23)
				(_nw285).ArraySet1((_1810_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))).Int()).(bool), 24)
				(_nw285).ArraySet1(true, 25)
				(_nw285).ArraySet1(!_dafny.Companion_Sequence_.Equal(_1879_cf15, _1879_cf15), 26)
				(_nw285).ArraySet1(Companion_Default___.Fm3(globalState), 27)
				_1892_v43 = _nw285
				var _1893_v44 _dafny.Array
				_ = _1893_v44
				var _len0_48 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_48
				var _nw286 _dafny.Array
				_ = _nw286
				if _len0_48.Cmp(_dafny.Zero) == 0 {
					_nw286 = _dafny.NewArray(_len0_48)
				} else {
					var _init48 func(_dafny.Int) _dafny.Int = (func(_1894_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1895_i4 _dafny.Int) _dafny.Int {
							return (_1895_i4).Times(_1894_p2)
						}
					})(p2)
					_ = _init48
					var _element0_48 = _init48(_dafny.Zero)
					_ = _element0_48
					_nw286 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
					(_nw286).ArraySet1(_element0_48, 0)
					var _nativeLen0_48 = (_len0_48).Int()
					_ = _nativeLen0_48
					for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
						(_nw286).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
					}
				}
				_1893_v44 = _nw286
				var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))
				_ = _index267
				var _rhs315 bool = !(p3)
				_ = _rhs315
				var _rhs316 bool = (_1810_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))).Int()).(bool)
				_ = _rhs316
				var _rhs317 _dafny.Array = _1892_v43
				_ = _rhs317
				var _rhs318 _dafny.Array = _1893_v44
				_ = _rhs318
				var _lhs262 _dafny.Array = _1810_v1
				_ = _lhs262
				var _lhs263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))
				_ = _lhs263
				var _lhs264 *GlobalState = globalState
				_ = _lhs264
				(_lhs262).ArraySet1(_rhs315, (_lhs263).Int())
				_lhs264.F0 = _rhs316
				_1892_v43 = _rhs317
				_1893_v44 = _rhs318
				var _1896_v45 _dafny.Array
				_ = _1896_v45
				var _len0_49 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_49
				var _nw287 _dafny.Array
				_ = _nw287
				if _len0_49.Cmp(_dafny.Zero) == 0 {
					_nw287 = _dafny.NewArray(_len0_49)
				} else {
					var _init49 func(_dafny.Int) _dafny.Sequence = (func(_1897_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1898_i5 _dafny.Int) _dafny.Sequence {
							return _1897_v2
						}
					})(_1812_v2)
					_ = _init49
					var _element0_49 = _init49(_dafny.Zero)
					_ = _element0_49
					_nw287 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
					(_nw287).ArraySet1(_element0_49, 0)
					var _nativeLen0_49 = (_len0_49).Int()
					_ = _nativeLen0_49
					for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
						(_nw287).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
					}
				}
				_1896_v45 = _nw287
				var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_1896_v45), 0))
				_ = _index268
				(_1896_v45).ArraySet1(_1812_v2, (_index268).Int())
				var _1899_v46 _dafny.Map
				_ = _1899_v46
				_1899_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1812_v2).Cardinality()), _1877_cf17)
				var _1900_v47 _dafny.Map
				_ = _1900_v47
				_1900_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1899_v46, _dafny.UnicodeSeqOfUtf8Bytes("dgimpf"))
				var _1901_v48 D14
				_ = _1901_v48
				_1901_v48 = Companion_D14_.Create_DC37_(Companion_Default___.Fm3(globalState), (func() _dafny.Sequence {
					if (_1900_v47).Contains(_1899_v46) {
						return (_1900_v47).Get(_1899_v46).(_dafny.Sequence)
					}
					return _1849_v17
				})())
				var _1902_v49 D5
				_ = _1902_v49
				_1902_v49 = Companion_D5_.Create_DC17_(_1812_v2, p3)
				var _pat_let_tv39 = p0
				_ = _pat_let_tv39
				var _pat_let_tv40 = _1810_v1
				_ = _pat_let_tv40
				var _pat_let_tv41 = _1810_v1
				_ = _pat_let_tv41
				var _pat_let_tv42 = _1849_v17
				_ = _pat_let_tv42
				var _1903_v50 _dafny.Array
				_ = _1903_v50
				var _nwElement0_53 D14 = _1901_v48
				_ = _nwElement0_53
				var _nw288 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(29))
				_ = _nw288
				(_nw288).ArraySet1(_nwElement0_53, 0)
				(_nw288).ArraySet1(Companion_D14_.Create_DC37_((func() bool {
					if (_1809_v0).Contains(p1) {
						return (_1809_v0).Get(p1).(bool)
					}
					return (_1810_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))).Int()).(bool)
				})(), _1849_v17), 1)
				(_nw288).ArraySet1(Companion_D14_.Create_DC37_(Companion_Default___.Fm3(globalState), _dafny.UnicodeSeqOfUtf8Bytes("im")), 2)
				(_nw288).ArraySet1(_1901_v48, 3)
				(_nw288).ArraySet1(_1901_v48, 4)
				(_nw288).ArraySet1(_1901_v48, 5)
				(_nw288).ArraySet1(_1901_v48, 6)
				(_nw288).ArraySet1(_1901_v48, 7)
				(_nw288).ArraySet1(_1901_v48, 8)
				(_nw288).ArraySet1(_1901_v48, 9)
				(_nw288).ArraySet1(_1901_v48, 10)
				(_nw288).ArraySet1(_1901_v48, 11)
				(_nw288).ArraySet1((func() D14 {
					if true {
						return Companion_D14_.Create_DC37_(p3, _1849_v17)
					}
					return _1901_v48
				})(), 12)
				(_nw288).ArraySet1(func(_pat_let49_0 D14) D14 {
					return func(_1904_dt__update__tmp_h1 D14) D14 {
						return func(_pat_let50_0 bool) D14 {
							return func(_1905_dt__update_hcf70_h0 bool) D14 {
								return Companion_D14_.Create_DC37_(_1905_dt__update_hcf70_h0, (_1904_dt__update__tmp_h1).Dtor_cf71())
							}(_pat_let50_0)
						}(_pat_let_tv39)
					}(_pat_let49_0)
				}(_1901_v48), 13)
				(_nw288).ArraySet1(Companion_D14_.Create_DC37_((_1810_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1810_v1), 0))).Int()).(bool), _1849_v17), 14)
				(_nw288).ArraySet1(_1901_v48, 15)
				(_nw288).ArraySet1(Companion_D14_.Create_DC37_(p3, _1849_v17), 16)
				(_nw288).ArraySet1(_1901_v48, 17)
				(_nw288).ArraySet1(_1901_v48, 18)
				(_nw288).ArraySet1(_1901_v48, 19)
				(_nw288).ArraySet1(_1901_v48, 20)
				(_nw288).ArraySet1(_1901_v48, 21)
				(_nw288).ArraySet1(func(_pat_let51_0 D14) D14 {
					return func(_1906_dt__update__tmp_h2 D14) D14 {
						return func(_pat_let52_0 bool) D14 {
							return func(_1907_dt__update_hcf70_h1 bool) D14 {
								return Companion_D14_.Create_DC37_(_1907_dt__update_hcf70_h1, (_1906_dt__update__tmp_h2).Dtor_cf71())
							}(_pat_let52_0)
						}((_pat_let_tv41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_pat_let_tv40), 0))).Int()).(bool))
					}(_pat_let51_0)
				}(_1901_v48), 22)
				(_nw288).ArraySet1(_1901_v48, 23)
				(_nw288).ArraySet1(_1901_v48, 24)
				(_nw288).ArraySet1(_1901_v48, 25)
				(_nw288).ArraySet1((func() D14 {
					if (_1902_v49).Dtor_cf37() {
						return _1901_v48
					}
					return _1901_v48
				})(), 26)
				(_nw288).ArraySet1(_1901_v48, 27)
				(_nw288).ArraySet1(func(_pat_let53_0 D14) D14 {
					return func(_1908_dt__update__tmp_h3 D14) D14 {
						return func(_pat_let54_0 _dafny.Sequence) D14 {
							return func(_1909_dt__update_hcf71_h0 _dafny.Sequence) D14 {
								return Companion_D14_.Create_DC37_((_1908_dt__update__tmp_h3).Dtor_cf70(), _1909_dt__update_hcf71_h0)
							}(_pat_let54_0)
						}(_pat_let_tv42)
					}(_pat_let53_0)
				}(Companion_Default___.Fm52(p1, globalState)), 28)
				_1903_v50 = _nw288
				var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_1903_v50), 0))
				_ = _index269
				(_1903_v50).ArraySet1(Companion_D14_.Create_DC37_(p0, _1849_v17), (_index269).Int())
				var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_1896_v45), 0))
				_ = _index270
				var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_1903_v50), 0))
				_ = _index271
				var _rhs319 _dafny.CodePoint = _1848_v16
				_ = _rhs319
				var _rhs320 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1812_v2, (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_this).Fm11(globalState), _1876_cf18), _dafny.IntOfUint32((_1812_v2).Cardinality()))).Uint32(), _1876_cf18)
				_ = _rhs320
				var _rhs321 _dafny.Array = _1878_cf16
				_ = _rhs321
				var _rhs322 bool = p0
				_ = _rhs322
				var _rhs323 D14 = _1901_v48
				_ = _rhs323
				var _lhs265 _dafny.Array = _1896_v45
				_ = _lhs265
				var _lhs266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_1896_v45), 0))
				_ = _lhs266
				var _lhs267 *GlobalState = globalState
				_ = _lhs267
				var _lhs268 _dafny.Array = _1903_v50
				_ = _lhs268
				var _lhs269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_1903_v50), 0))
				_ = _lhs269
				_1848_v16 = _rhs319
				(_lhs265).ArraySet1(_rhs320, (_lhs266).Int())
				_1878_cf16 = _rhs321
				_lhs267.F0 = _rhs322
				(_lhs268).ArraySet1(_rhs323, (_lhs269).Int())
				var _1910_v51 _dafny.Map
				_ = _1910_v51
				_1910_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1877_cf17, p1)
				_1910_v51 = (_1910_v51).Update(p2, p1)
			}
			var _source23 D2 = _1829_v15
			_ = _source23
			if _source23.Is_DC6() {
				var _1911___mcc_h18 _dafny.Int = _source23.Get_().(D2_DC6).Cf13
				_ = _1911___mcc_h18
				var _1912___mcc_h19 _dafny.Array = _source23.Get_().(D2_DC6).Cf14
				_ = _1912___mcc_h19
				var _1913_cf14 _dafny.Array = _1912___mcc_h19
				_ = _1913_cf14
				var _1914_cf13 _dafny.Int = _1911___mcc_h18
				_ = _1914_cf13
				var _1915_v52 _dafny.MultiSet
				_ = _1915_v52
				_1915_v52 = _dafny.MultiSetOf(_dafny.CodePoint('h'), _1848_v16)
				var _1916_v53 D1
				_ = _1916_v53
				_1916_v53 = Companion_D1_.Create_DC3_(_1915_v52)
				var _1917_v54 _dafny.MultiSet
				_ = _1917_v54
				_1917_v54 = _dafny.MultiSetOf(_1916_v53, _1916_v53, _1916_v53, _1916_v53, _1916_v53)
				var _rhs324 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(61))).Uint32(), func(coer116 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg116 _dafny.Int) interface{} {
						return coer116(arg116)
					}
				}((func(_1918_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1919_i6 _dafny.Int) _dafny.CodePoint {
						return _1918_v16
					}
				})(_1848_v16)))
				_ = _rhs324
				var _rhs325 bool = (Companion_Default___.Fm53(_1848_v16, globalState)).IsDisjointFrom(_1917_v54)
				_ = _rhs325
				var _lhs270 *GlobalState = globalState
				_ = _lhs270
				var _lhs271 *GlobalState = globalState
				_ = _lhs271
				_lhs270.F2 = _rhs324
				_lhs271.F0 = _rhs325
				(globalState).F5 = _1877_cf17
				(globalState).F0 = p0
				var _1920_v55 _dafny.Array
				_ = _1920_v55
				var _nw289 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
				_ = _nw289
				_1920_v55 = _nw289
				var _1921_v56 D10
				_ = _1921_v56
				_1921_v56 = Companion_D10_.Create_DC29_(_dafny.IntOfInt64(534), _1920_v55, true, p2)
				var _1922_v57 _dafny.Map
				_ = _1922_v57
				_1922_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(525))
				var _pat_let_tv43 = _1920_v55
				_ = _pat_let_tv43
				var _pat_let_tv44 = _1914_cf13
				_ = _pat_let_tv44
				var _pat_let_tv45 = p3
				_ = _pat_let_tv45
				var _1923_v58 _dafny.Array
				_ = _1923_v58
				var _nwElement0_54 D10 = _1921_v56
				_ = _nwElement0_54
				var _nw290 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(23))
				_ = _nw290
				(_nw290).ArraySet1(_nwElement0_54, 0)
				(_nw290).ArraySet1(_1921_v56, 1)
				(_nw290).ArraySet1(_1921_v56, 2)
				(_nw290).ArraySet1(_1921_v56, 3)
				(_nw290).ArraySet1(_1921_v56, 4)
				(_nw290).ArraySet1((func() D10 {
					if p3 {
						return func(_pat_let55_0 D10) D10 {
							return func(_1924_dt__update__tmp_h4 D10) D10 {
								return func(_pat_let56_0 _dafny.Array) D10 {
									return func(_1925_dt__update_hcf54_h0 _dafny.Array) D10 {
										return func(_pat_let57_0 _dafny.Int) D10 {
											return func(_1926_dt__update_hcf56_h0 _dafny.Int) D10 {
												return func(_pat_let58_0 bool) D10 {
													return func(_1927_dt__update_hcf55_h0 bool) D10 {
														return Companion_D10_.Create_DC29_((_1924_dt__update__tmp_h4).Dtor_cf53(), _1925_dt__update_hcf54_h0, _1927_dt__update_hcf55_h0, _1926_dt__update_hcf56_h0)
													}(_pat_let58_0)
												}(_pat_let_tv45)
											}(_pat_let57_0)
										}(_pat_let_tv44)
									}(_pat_let56_0)
								}(_pat_let_tv43)
							}(_pat_let55_0)
						}(Companion_D10_.Create_DC29_(_dafny.IntOfUint32((_1849_v17).Cardinality()), _1920_v55, p1, _dafny.IntOfInt64(-766)))
					}
					return _1921_v56
				})(), 5)
				(_nw290).ArraySet1(_1921_v56, 6)
				(_nw290).ArraySet1(_1921_v56, 7)
				(_nw290).ArraySet1(_1921_v56, 8)
				(_nw290).ArraySet1(_1921_v56, 9)
				(_nw290).ArraySet1(_1921_v56, 10)
				(_nw290).ArraySet1(Companion_D10_.Create_DC29_(_1876_cf18, _1920_v55, p0, _1877_cf17), 11)
				(_nw290).ArraySet1(_1921_v56, 12)
				(_nw290).ArraySet1(Companion_D10_.Create_DC29_(_1877_cf17, _1920_v55, p3, _1914_cf13), 13)
				(_nw290).ArraySet1(_1921_v56, 14)
				(_nw290).ArraySet1(_1921_v56, 15)
				(_nw290).ArraySet1(_1921_v56, 16)
				(_nw290).ArraySet1(_1921_v56, 17)
				(_nw290).ArraySet1(_1921_v56, 18)
				(_nw290).ArraySet1(_1921_v56, 19)
				(_nw290).ArraySet1(Companion_D10_.Create_DC29_(((_1922_v57).Update(p0, _1876_cf18)).Cardinality(), _1920_v55, false, p2), 20)
				(_nw290).ArraySet1(_1921_v56, 21)
				(_nw290).ArraySet1(_1921_v56, 22)
				_1923_v58 = _nw290
				_1923_v58 = _1923_v58
			} else if _source23.Is_DC7() {
				var _1928___mcc_h20 _dafny.Sequence = _source23.Get_().(D2_DC7).Cf15
				_ = _1928___mcc_h20
				var _1929___mcc_h21 _dafny.Array = _source23.Get_().(D2_DC7).Cf16
				_ = _1929___mcc_h21
				var _1930___mcc_h22 _dafny.Int = _source23.Get_().(D2_DC7).Cf17
				_ = _1930___mcc_h22
				var _1931___mcc_h23 _dafny.Int = _source23.Get_().(D2_DC7).Cf18
				_ = _1931___mcc_h23
				var _1932_cf18 _dafny.Int = _1931___mcc_h23
				_ = _1932_cf18
				var _1933_cf17 _dafny.Int = _1930___mcc_h22
				_ = _1933_cf17
				var _1934_cf16 _dafny.Array = _1929___mcc_h21
				_ = _1934_cf16
				var _1935_cf15 _dafny.Sequence = _1928___mcc_h20
				_ = _1935_cf15
				var _1936_v59 _dafny.MultiSet
				_ = _1936_v59
				_1936_v59 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1812_v2).Cardinality()), p2, p2, _1933_cf17, _dafny.IntOfInt64(534))
				var _1937_v60 _dafny.Sequence
				_ = _1937_v60
				_1937_v60 = _dafny.SeqOf(_1812_v2, _1812_v2)
				var _1938_v61 _dafny.Set
				_ = _1938_v61
				_1938_v61 = _dafny.SetOf(_dafny.IntOfUint32(((_1937_v60).Select((Companion_Default___.SafeIndex(_1933_cf17, _dafny.IntOfUint32((_1937_v60).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
				_1848_v16 = Companion_Default___.Fm0((_1936_v59).Cardinality(), _1938_v61, p0, globalState)
				var _1939_v62 _dafny.Map
				_ = _1939_v62
				_1939_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1934_cf16, _1933_cf17)
				_1939_v62 = (_1939_v62).Update(_1934_cf16, _1932_cf18)
				var _rhs326 bool = !(p3) || ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1932_cf18)).Contains(p1))
				_ = _rhs326
				var _rhs327 _dafny.Int = _1932_cf18
				_ = _rhs327
				var _lhs272 *GlobalState = globalState
				_ = _lhs272
				_lhs272.F0 = _rhs326
				r2 = _rhs327
				var _1940_v63 _dafny.Map
				_ = _1940_v63
				_1940_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1933_cf17, Companion_Default___.Fm0(_1933_cf17, _1938_v61, p0, globalState))
				var _1941_v64 _dafny.MultiSet
				_ = _1941_v64
				_1941_v64 = _dafny.MultiSetOf(_1940_v63, _1940_v63)
				var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_1810_v1), 0))
				_ = _index272
				(_1810_v1).ArraySet1((_1941_v64).IsDisjointFrom(_1941_v64), (_index272).Int())
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_1810_v1), 0))
				_ = _index273
				(_1810_v1).ArraySet1(p0, (_index273).Int())
			} else if _source23.Is_DC8() {
				var _1942___mcc_h24 bool = _source23.Get_().(D2_DC8).Cf19
				_ = _1942___mcc_h24
				var _1943___mcc_h25 bool = _source23.Get_().(D2_DC8).Cf20
				_ = _1943___mcc_h25
				var _1944_cf20 bool = _1943___mcc_h25
				_ = _1944_cf20
				var _1945_cf19 bool = _1942___mcc_h24
				_ = _1945_cf19
				var _1946_v65 *C1
				_ = _1946_v65
				var _nw291 *C1 = New_C1_()
				_ = _nw291
				_nw291.Ctor__(_1877_cf17)
				_1946_v65 = _nw291
				var _1947_v66 _dafny.Map
				_ = _1947_v66
				_1947_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
				var _1948_v67 _dafny.Set
				_ = _1948_v67
				_1948_v67 = _dafny.SetOf(_1944_cf20)
				var _1949_v68 _dafny.Array
				_ = _1949_v68
				var _nwElement0_55 _dafny.Int = p2
				_ = _nwElement0_55
				var _nw292 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(9))
				_ = _nw292
				(_nw292).ArraySet1(_nwElement0_55, 0)
				(_nw292).ArraySet1((_1947_v66).Cardinality(), 1)
				(_nw292).ArraySet1(_1877_cf17, 2)
				(_nw292).ArraySet1(p2, 3)
				(_nw292).ArraySet1((_1948_v67).Cardinality(), 4)
				(_nw292).ArraySet1(_1876_cf18, 5)
				(_nw292).ArraySet1(p2, 6)
				(_nw292).ArraySet1(_1876_cf18, 7)
				(_nw292).ArraySet1(_dafny.IntOfUint32((_1849_v17).Cardinality()), 8)
				_1949_v68 = _nw292
				var _1950_v69 _dafny.Map
				_ = _1950_v69
				_1950_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1949_v68, _1944_cf20)
				var _1951_v70 _dafny.Set
				_ = _1951_v70
				_1951_v70 = _dafny.SetOf((_1950_v69).Cardinality())
				var _1952_v71 _dafny.Map
				_ = _1952_v71
				_1952_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1951_v70).Cardinality(), p1)
				var _1953_v72 _dafny.Map
				_ = _1953_v72
				_1953_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1877_cf17, Companion_Default___.Fm4((_1947_v66).Cardinality(), p2, _1877_cf17, globalState))
				var _1954_v73 _dafny.Array
				_ = _1954_v73
				var _nwElement0_56 _dafny.Int = p2
				_ = _nwElement0_56
				var _nw293 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(11))
				_ = _nw293
				(_nw293).ArraySet1(_nwElement0_56, 0)
				(_nw293).ArraySet1(_1876_cf18, 1)
				(_nw293).ArraySet1(_1877_cf17, 2)
				(_nw293).ArraySet1(_1876_cf18, 3)
				(_nw293).ArraySet1(_1877_cf17, 4)
				(_nw293).ArraySet1(((_1953_v72).Cardinality()).Plus(p2), 5)
				(_nw293).ArraySet1(_1877_cf17, 6)
				(_nw293).ArraySet1(_1877_cf17, 7)
				(_nw293).ArraySet1(_dafny.IntOfInt64(902), 8)
				(_nw293).ArraySet1(_1877_cf17, 9)
				(_nw293).ArraySet1(_dafny.IntOfInt64(905), 10)
				_1954_v73 = _nw293
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1954_v73), 0))
				_ = _index274
				(_1954_v73).ArraySet1(_1876_cf18, (_index274).Int())
				var _1955_v74 _dafny.MultiSet
				_ = _1955_v74
				_1955_v74 = _dafny.MultiSetOf(p0)
				var _1956_v75 _dafny.MultiSet
				_ = _1956_v75
				_1956_v75 = _1955_v74
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1954_v73), 0))
				_ = _index275
				var _rhs328 _dafny.Int = (_dafny.Zero).Minus((_1876_cf18).Minus(_1877_cf17))
				_ = _rhs328
				var _rhs329 _dafny.Int = p2
				_ = _rhs329
				var _rhs330 _dafny.Int = p2
				_ = _rhs330
				var _rhs331 _dafny.MultiSet = ((_1955_v74).Intersection((_1956_v75))).Intersection((_1955_v74).Difference(_1955_v74))
				_ = _rhs331
				var _lhs273 _dafny.Array = _1954_v73
				_ = _lhs273
				var _lhs274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1954_v73), 0))
				_ = _lhs274
				var _lhs275 *GlobalState = globalState
				_ = _lhs275
				_1876_cf18 = _rhs328
				(_lhs273).ArraySet1(_rhs329, (_lhs274).Int())
				_lhs275.F5 = _rhs330
				_1955_v74 = _rhs331
				var _1957_v76 _dafny.Set
				_ = _1957_v76
				_1957_v76 = _dafny.SetOf(_1951_v70, _1951_v70)
				var _1958_v77 _dafny.Map
				_ = _1958_v77
				_1958_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if _1945_cf19 {
						return _1877_cf17
					}
					return _1876_cf18
				})(), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1957_v76).Cardinality(), (_dafny.Zero).Minus((_1954_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1954_v73), 0))).Int()).(_dafny.Int)))))
				var _1959_v78 _dafny.Map
				_ = _1959_v78
				_1959_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1810_v1, p1)
				var _1960_v79 D14
				_ = _1960_v79
				_1960_v79 = Companion_D14_.Create_DC37_(_1944_cf20, _1849_v17)
				var _1961_v80 D0
				_ = _1961_v80
				_1961_v80 = Companion_D0_.Create_DC1_(true, _1876_cf18, Companion_Default___.Fm0((_1954_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1954_v73), 0))).Int()).(_dafny.Int), _1951_v70, _1945_cf19, globalState))
				var _pat_let_tv46 = p3
				_ = _pat_let_tv46
				var _pat_let_tv47 = _1848_v16
				_ = _pat_let_tv47
				var _rhs332 _dafny.Map = (_1958_v77).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1879_cf15).Cardinality()), (func(_pat_let59_0 D0) D0 {
					return func(_1962_dt__update__tmp_h5 D0) D0 {
						return func(_pat_let60_0 bool) D0 {
							return func(_1963_dt__update_hcf2_h1 bool) D0 {
								return func(_pat_let61_0 _dafny.CodePoint) D0 {
									return func(_1964_dt__update_hcf4_h0 _dafny.CodePoint) D0 {
										return Companion_D0_.Create_DC1_(_1963_dt__update_hcf2_h1, (_1962_dt__update__tmp_h5).Dtor_cf3(), _1964_dt__update_hcf4_h0)
									}(_pat_let61_0)
								}(_pat_let_tv47)
							}(_pat_let60_0)
						}(_pat_let_tv46)
					}(_pat_let59_0)
				}(_1961_v80)).Dtor_cf3()))
				_ = _rhs332
				var _rhs333 _dafny.Map = (func() _dafny.Map {
					if _1944_cf20 {
						return _1959_v78
					}
					return _1959_v78
				})()
				_ = _rhs333
				var _rhs334 D14 = _1960_v79
				_ = _rhs334
				_1958_v77 = _rhs332
				_1959_v78 = _rhs333
				_1960_v79 = _rhs334
				var _1965_v82 *C3
				_ = _1965_v82
				var _nw294 *C3 = New_C3_()
				_ = _nw294
				_nw294.Ctor__(_1958_v77, _1810_v1, func() _dafny.Map {
					var _coll52 = _dafny.NewMapBuilder()
					_ = _coll52
					for _iter55 := _dafny.Iterate((_1947_v66).Keys().Elements()); ; {
						_compr_52, _ok55 := _iter55()
						if !_ok55 {
							break
						}
						var _1966_v81 _dafny.Int
						_1966_v81 = interface{}(_compr_52).(_dafny.Int)
						if (_1947_v66).Contains(_1966_v81) {
							_coll52.Add((_1966_v81).Times(_1876_cf18), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(517))).Uint32(), func(coer117 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg117 _dafny.Int) interface{} {
									return coer117(arg117)
								}
							}((func(_1967_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1968_i7 _dafny.Int) _dafny.CodePoint {
									return _1967_v16
								}
							})(_1848_v16))))
						}
					}
					return _coll52.ToMap()
				}(), _1876_cf18)
				_1965_v82 = _nw294
				var _1969_v83 _dafny.Sequence
				_ = _1969_v83
				_1969_v83 = _dafny.SeqOf(_1965_v82)
				var _1970_v84 D17
				_ = _1970_v84
				_1970_v84 = Companion_D17_.Create_DC40_(_1965_v82)
				var _1971_v85 _dafny.Array
				_ = _1971_v85
				var _nwElement0_57 *C3 = _1965_v82
				_ = _nwElement0_57
				var _nw295 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(20))
				_ = _nw295
				(_nw295).ArraySet1(_nwElement0_57, 0)
				(_nw295).ArraySet1(_1965_v82, 1)
				(_nw295).ArraySet1(_1965_v82, 2)
				(_nw295).ArraySet1((func() *C3 {
					if p0 {
						return _1965_v82
					}
					return _1965_v82
				})(), 3)
				(_nw295).ArraySet1(_1965_v82, 4)
				(_nw295).ArraySet1(_1965_v82, 5)
				(_nw295).ArraySet1(_1965_v82, 6)
				(_nw295).ArraySet1(_1965_v82, 7)
				(_nw295).ArraySet1((_1969_v83).Select((Companion_Default___.SafeIndex(_1876_cf18, _dafny.IntOfUint32((_1969_v83).Cardinality()))).Uint32()).(*C3), 8)
				(_nw295).ArraySet1(_1965_v82, 9)
				(_nw295).ArraySet1(_1965_v82, 10)
				(_nw295).ArraySet1(_1965_v82, 11)
				(_nw295).ArraySet1(_1965_v82, 12)
				(_nw295).ArraySet1(_1965_v82, 13)
				(_nw295).ArraySet1(_1965_v82, 14)
				(_nw295).ArraySet1(_1965_v82, 15)
				(_nw295).ArraySet1(_1965_v82, 16)
				(_nw295).ArraySet1((_1970_v84).Dtor_cf74(), 17)
				(_nw295).ArraySet1(_1965_v82, 18)
				(_nw295).ArraySet1(_1965_v82, 19)
				_1971_v85 = _nw295
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_1971_v85), 0))
				_ = _index276
				(_1971_v85).ArraySet1(_1965_v82, (_index276).Int())
				var _1972_v86 _dafny.Map
				_ = _1972_v86
				_1972_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1877_cf17, _1965_v82)
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_1971_v85), 0))
				_ = _index277
				(_1971_v85).ArraySet1((func() *C3 {
					if (_1972_v86).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus(((_1954_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1954_v73), 0))).Int()).(_dafny.Int)).Plus((_1954_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1954_v73), 0))).Int()).(_dafny.Int))))) {
						return (_1972_v86).Get((_dafny.Zero).Minus((_dafny.Zero).Minus(((_1954_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1954_v73), 0))).Int()).(_dafny.Int)).Plus((_1954_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1954_v73), 0))).Int()).(_dafny.Int))))).(*C3)
					}
					return _1965_v82
				})(), (_index277).Int())
			} else {
				var _1973___mcc_h26 _dafny.Map = _source23.Get_().(D2_DC5).Cf12
				_ = _1973___mcc_h26
				var _1974_cf12 _dafny.Map = _1973___mcc_h26
				_ = _1974_cf12
				var _1975_v87 _dafny.Array
				_ = _1975_v87
				var _nw296 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
				_ = _nw296
				_1975_v87 = _nw296
				var _1976_v88 _dafny.MultiSet
				_ = _1976_v88
				_1976_v88 = _dafny.MultiSetOf(_1975_v87)
				var _1977_v89 _dafny.Map
				_ = _1977_v89
				_1977_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(403))).Uint32(), func(coer118 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg118 _dafny.Int) interface{} {
						return coer118(arg118)
					}
				}((func(_1978_cf18 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1979_i8 _dafny.Int) _dafny.Int {
						return _1978_cf18
					}
				})(_1876_cf18))))
				var _1980_v90 _dafny.MultiSet
				_ = _1980_v90
				_1980_v90 = _dafny.MultiSetOf(p1)
				var _1981_v91 _dafny.Map
				_ = _1981_v91
				_1981_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(284), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1849_v17, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.IntOfUint32((_1849_v17).Cardinality()))).Uint32(), _1848_v16), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(396))).Uint32(), func(coer119 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg119 _dafny.Int) interface{} {
						return coer119(arg119)
					}
				}((func(_1982_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1983_i9 _dafny.Int) _dafny.CodePoint {
						return _1982_v16
					}
				})(_1848_v16)))))
				var _rhs335 bool = (_1976_v88).IsDisjointFrom(_1976_v88)
				_ = _rhs335
				var _rhs336 bool = (_1877_cf17).Cmp(_1877_cf17) != 0
				_ = _rhs336
				var _rhs337 bool = (_1876_cf18).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1977_v89).Contains(p0) {
						return (_1977_v89).Get(p0).(_dafny.Sequence)
					}
					return _1812_v2
				})()).Cardinality()), (_1980_v90).Cardinality())) <= 0
				_ = _rhs337
				var _rhs338 _dafny.Sequence = (func() _dafny.Sequence {
					if (_1981_v91).Contains(_1876_cf18) {
						return (_1981_v91).Get(_1876_cf18).(_dafny.Sequence)
					}
					return _1849_v17
				})()
				_ = _rhs338
				var _lhs276 *GlobalState = globalState
				_ = _lhs276
				var _lhs277 *GlobalState = globalState
				_ = _lhs277
				var _lhs278 *GlobalState = globalState
				_ = _lhs278
				_lhs276.F0 = _rhs335
				_lhs277.F0 = _rhs336
				_lhs278.F0 = _rhs337
				_1849_v17 = _rhs338
				_1810_v1 = _1810_v1
				_1848_v16 = _1848_v16
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_1975_v87), 0))
				_ = _index278
				(_1975_v87).ArraySet1(_1877_cf17, (_index278).Int())
				var _1984_v92 _dafny.Sequence
				_ = _1984_v92
				_1984_v92 = _dafny.SeqOf(_1975_v87)
				var _1985_v93 D4
				_ = _1985_v93
				_1985_v93 = Companion_D4_.Create_DC15_(p1, p0, _dafny.Companion_Sequence_.Update(_1849_v17, (Companion_Default___.SafeIndex(_1877_cf17, _dafny.IntOfUint32((_1849_v17).Cardinality()))).Uint32(), _1848_v16))
				var _1986_v94 _dafny.Map
				_ = _1986_v94
				_1986_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1985_v93, _1975_v87)
				var _1987_v95 _dafny.Map
				_ = _1987_v95
				_1987_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (func() _dafny.Array {
					if (_1986_v94).Contains(_1985_v93) {
						return (_1986_v94).Get(_1985_v93).(_dafny.Array)
					}
					return _1975_v87
				})())
				var _1988_v96 D2
				_ = _1988_v96
				_1988_v96 = Companion_D2_.Create_DC6_(p2, _1975_v87)
				var _1989_v97 _dafny.Array
				_ = _1989_v97
				var _nwElement0_58 _dafny.Array = _1975_v87
				_ = _nwElement0_58
				var _nw297 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(28))
				_ = _nw297
				(_nw297).ArraySet1(_nwElement0_58, 0)
				(_nw297).ArraySet1((_1984_v92).Select((Companion_Default___.SafeIndex((_1987_v95).Cardinality(), _dafny.IntOfUint32((_1984_v92).Cardinality()))).Uint32()).(_dafny.Array), 1)
				(_nw297).ArraySet1(_1975_v87, 2)
				(_nw297).ArraySet1(_1975_v87, 3)
				(_nw297).ArraySet1(_1975_v87, 4)
				(_nw297).ArraySet1(_1975_v87, 5)
				(_nw297).ArraySet1(_1975_v87, 6)
				(_nw297).ArraySet1(_1975_v87, 7)
				(_nw297).ArraySet1((func() _dafny.Array {
					if p1 {
						return _1975_v87
					}
					return _1975_v87
				})(), 8)
				(_nw297).ArraySet1(_1975_v87, 9)
				(_nw297).ArraySet1((func() _dafny.Array {
					if p0 {
						return _1975_v87
					}
					return (_1984_v92).Select((Companion_Default___.SafeIndex(_1877_cf17, _dafny.IntOfUint32((_1984_v92).Cardinality()))).Uint32()).(_dafny.Array)
				})(), 10)
				(_nw297).ArraySet1(_1975_v87, 11)
				(_nw297).ArraySet1(_1975_v87, 12)
				(_nw297).ArraySet1(_1975_v87, 13)
				(_nw297).ArraySet1(_1975_v87, 14)
				(_nw297).ArraySet1((_1988_v96).Dtor_cf14(), 15)
				(_nw297).ArraySet1(_1975_v87, 16)
				(_nw297).ArraySet1(_1975_v87, 17)
				(_nw297).ArraySet1(_1975_v87, 18)
				(_nw297).ArraySet1(_1975_v87, 19)
				(_nw297).ArraySet1((func() _dafny.Array {
					if p3 {
						return _1975_v87
					}
					return _1975_v87
				})(), 20)
				(_nw297).ArraySet1(_1975_v87, 21)
				(_nw297).ArraySet1(_1975_v87, 22)
				(_nw297).ArraySet1(_1975_v87, 23)
				(_nw297).ArraySet1(_1975_v87, 24)
				(_nw297).ArraySet1(_1975_v87, 25)
				(_nw297).ArraySet1(_1975_v87, 26)
				(_nw297).ArraySet1(_1975_v87, 27)
				_1989_v97 = _nw297
				var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_1989_v97), 0))
				_ = _index279
				(_1989_v97).ArraySet1(_1975_v87, (_index279).Int())
				var _1990_v98 _dafny.Set
				_ = _1990_v98
				_1990_v98 = _dafny.SetOf((_1849_v17).Select((Companion_Default___.SafeIndex(_1876_cf18, _dafny.IntOfUint32((_1849_v17).Cardinality()))).Uint32()).(_dafny.CodePoint), _1848_v16, _1848_v16)
				var _1991_v99 _dafny.Map
				_ = _1991_v99
				_1991_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_1876_cf18, globalState), _1812_v2)
				var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_1975_v87), 0))
				_ = _index280
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_1989_v97), 0))
				_ = _index281
				var _rhs339 _dafny.Int = _1877_cf17
				_ = _rhs339
				var _rhs340 _dafny.Array = _1975_v87
				_ = _rhs340
				var _rhs341 _dafny.Set = (_1990_v98).Intersection(_dafny.SetOf(_1848_v16, _dafny.CodePoint('y')))
				_ = _rhs341
				var _rhs342 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1991_v99).Contains(_dafny.IntOfUint32((_1879_cf15).Cardinality())) {
						return (_1991_v99).Get(_dafny.IntOfUint32((_1879_cf15).Cardinality())).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(836))).Uint32(), func(coer120 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg120 _dafny.Int) interface{} {
							return coer120(arg120)
						}
					}((func(_1992_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1993_i10 _dafny.Int) _dafny.Int {
							return _1992_p2
						}
					})(p2)))
				})()).Cardinality()))
				_ = _rhs342
				var _lhs279 _dafny.Array = _1975_v87
				_ = _lhs279
				var _lhs280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_1975_v87), 0))
				_ = _lhs280
				var _lhs281 _dafny.Array = _1989_v97
				_ = _lhs281
				var _lhs282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_1989_v97), 0))
				_ = _lhs282
				var _lhs283 *GlobalState = globalState
				_ = _lhs283
				(_lhs279).ArraySet1(_rhs339, (_lhs280).Int())
				(_lhs281).ArraySet1(_rhs340, (_lhs282).Int())
				_1990_v98 = _rhs341
				_lhs283.F5 = _rhs342
			}
			var _1994_v100 _dafny.Array
			_ = _1994_v100
			var _nw298 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
			_ = _nw298
			_1994_v100 = _nw298
			var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_1994_v100), 0))
			_ = _index282
			(_1994_v100).ArraySet1(_1812_v2, (_index282).Int())
			var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_1994_v100), 0))
			_ = _index283
			(_1994_v100).ArraySet1(_1812_v2, (_index283).Int())
			(globalState).F5 = (_1877_cf17).Plus(_1877_cf17)
		} else if _source22.Is_DC8() {
			var _1995___mcc_h15 bool = _source22.Get_().(D2_DC8).Cf19
			_ = _1995___mcc_h15
			var _1996___mcc_h16 bool = _source22.Get_().(D2_DC8).Cf20
			_ = _1996___mcc_h16
			var _1997_cf20 bool = _1996___mcc_h16
			_ = _1997_cf20
			var _1998_cf19 bool = _1995___mcc_h15
			_ = _1998_cf19
			var _1999_v101 _dafny.Array
			_ = _1999_v101
			var _nw299 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw299
			_1999_v101 = _nw299
			var _2000_v102 _dafny.Map
			_ = _2000_v102
			_2000_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1999_v101, _1997_cf20)
			var _2001_v103 D10
			_ = _2001_v103
			_2001_v103 = Companion_D10_.Create_DC28_(_2000_v102)
			var _source24 D10 = _2001_v103
			_ = _source24
			if _source24.Is_DC29() {
				var _2002___mcc_h27 _dafny.Int = _source24.Get_().(D10_DC29).Cf53
				_ = _2002___mcc_h27
				var _2003___mcc_h28 _dafny.Array = _source24.Get_().(D10_DC29).Cf54
				_ = _2003___mcc_h28
				var _2004___mcc_h29 bool = _source24.Get_().(D10_DC29).Cf55
				_ = _2004___mcc_h29
				var _2005___mcc_h30 _dafny.Int = _source24.Get_().(D10_DC29).Cf56
				_ = _2005___mcc_h30
				var _2006_cf56 _dafny.Int = _2005___mcc_h30
				_ = _2006_cf56
				var _2007_cf55 bool = _2004___mcc_h29
				_ = _2007_cf55
				var _2008_cf54 _dafny.Array = _2003___mcc_h28
				_ = _2008_cf54
				var _2009_cf53 _dafny.Int = _2002___mcc_h27
				_ = _2009_cf53
				var _2010_v104 _dafny.Map
				_ = _2010_v104
				_2010_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2009_cf53, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(958))).Uint32(), func(coer121 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg121 _dafny.Int) interface{} {
						return coer121(arg121)
					}
				}((func(_2011_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2012_i11 _dafny.Int) _dafny.CodePoint {
						return _2011_v16
					}
				})(_1848_v16))))
				var _2013_v105 *C7
				_ = _2013_v105
				var _nw300 *C7 = New_C7_()
				_ = _nw300
				_nw300.Ctor__(_1999_v101, _1999_v101, _2009_cf53, _1810_v1, _2010_v104)
				_2013_v105 = _nw300
				var _2014_v106 _dafny.Set
				_ = _2014_v106
				_2014_v106 = _dafny.SetOf(p3)
				var _2015_v107 D11
				_ = _2015_v107
				_2015_v107 = Companion_D11_.Create_DC33_(_2009_cf53, _2007_cf55, _dafny.IntOfInt64(655), false, _2014_v106)
				var _2016_v108 _dafny.Map
				_ = _2016_v108
				_2016_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_2015_v107).Dtor_cf65(), _1998_cf19, _1998_cf19, p1, _1998_cf19), _2010_v104)
				var _2017_v109 _dafny.Sequence
				_ = _2017_v109
				_2017_v109 = _dafny.SeqOf(p0, _2007_cf55)
				var _nw301 *C7 = New_C7_()
				_ = _nw301
				_nw301.Ctor__(_2013_v105.F11, _1999_v101, _2009_cf53, _1810_v1, (func() _dafny.Map {
					if (_2016_v108).Contains(_2017_v109) {
						return (_2016_v108).Get(_2017_v109).(_dafny.Map)
					}
					return _2010_v104
				})())
				_2013_v105 = _nw301
				_1849_v17 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1849_v17, _1849_v17), _1849_v17)
				var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_2008_cf54), 0))
				_ = _index284
				(_2008_cf54).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("btcbtiuf"), _1849_v17), (_index284).Int())
				var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_2008_cf54), 0))
				_ = _index285
				var _rhs343 _dafny.Int = p2
				_ = _rhs343
				var _rhs344 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1849_v17, _1849_v17)
				_ = _rhs344
				var _lhs284 _dafny.Array = _2008_cf54
				_ = _lhs284
				var _lhs285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_2008_cf54), 0))
				_ = _lhs285
				_2009_cf53 = _rhs343
				(_lhs284).ArraySet1(_rhs344, (_lhs285).Int())
				(globalState).F5 = (p2).Plus(_2009_cf53)
			} else {
				var _2018___mcc_h31 _dafny.Map = _source24.Get_().(D10_DC28).Cf52
				_ = _2018___mcc_h31
				var _2019_cf52 _dafny.Map = _2018___mcc_h31
				_ = _2019_cf52
				var _2020_v110 _dafny.MultiSet
				_ = _2020_v110
				_2020_v110 = _dafny.MultiSetOf(Companion_Default___.Fm0(_dafny.IntOfInt64(192), _dafny.SetOf(_dafny.IntOfInt64(-474)), _1998_cf19, globalState), _1848_v16)
				var _2021_v111 D1
				_ = _2021_v111
				_2021_v111 = Companion_D1_.Create_DC3_(_2020_v110)
				var _pat_let_tv48 = _2020_v110
				_ = _pat_let_tv48
				var _2022_v112 _dafny.Map
				_ = _2022_v112
				_2022_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, func(_pat_let62_0 D1) D1 {
					return func(_2023_dt__update__tmp_h6 D1) D1 {
						return func(_pat_let63_0 _dafny.MultiSet) D1 {
							return func(_2024_dt__update_hcf6_h0 _dafny.MultiSet) D1 {
								return Companion_D1_.Create_DC3_(_2024_dt__update_hcf6_h0)
							}(_pat_let63_0)
						}(_pat_let_tv48)
					}(_pat_let62_0)
				}(_2021_v111))
				_2022_v112 = (_2022_v112).Update(p2, _2021_v111)
				var _2025_v113 _dafny.MultiSet
				_ = _2025_v113
				_2025_v113 = _dafny.MultiSetOf(p2, _dafny.IntOfUint32((_1849_v17).Cardinality()))
				var _2026_v114 _dafny.Sequence
				_ = _2026_v114
				_2026_v114 = _dafny.SeqOf(_2025_v113, _2025_v113, _2025_v113)
				var _2027_v115 D11
				_ = _2027_v115
				_2027_v115 = Companion_D11_.Create_DC31_(p2, p2)
				var _pat_let_tv49 = p2
				_ = _pat_let_tv49
				var _pat_let_tv50 = p2
				_ = _pat_let_tv50
				var _pat_let_tv51 = _1849_v17
				_ = _pat_let_tv51
				var _2028_v116 _dafny.Map
				_ = _2028_v116
				_2028_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2026_v114).Cardinality()), (func(_pat_let64_0 D11) D11 {
					return func(_2033_dt__update__tmp_h9 D11) D11 {
						return func(_pat_let69_0 _dafny.Int) D11 {
							return func(_2034_dt__update_hcf58_h0 _dafny.Int) D11 {
								return Companion_D11_.Create_DC31_(_2034_dt__update_hcf58_h0, (_2033_dt__update__tmp_h9).Dtor_cf59())
							}(_pat_let69_0)
						}(_dafny.IntOfUint32((_pat_let_tv51).Cardinality()))
					}(_pat_let64_0)
				}(func(_pat_let65_0 D11) D11 {
					return func(_2031_dt__update__tmp_h8 D11) D11 {
						return func(_pat_let68_0 _dafny.Int) D11 {
							return func(_2032_dt__update_hcf59_h1 _dafny.Int) D11 {
								return Companion_D11_.Create_DC31_((_2031_dt__update__tmp_h8).Dtor_cf58(), _2032_dt__update_hcf59_h1)
							}(_pat_let68_0)
						}(_pat_let_tv50)
					}(_pat_let65_0)
				}(func(_pat_let66_0 D11) D11 {
					return func(_2029_dt__update__tmp_h7 D11) D11 {
						return func(_pat_let67_0 _dafny.Int) D11 {
							return func(_2030_dt__update_hcf59_h0 _dafny.Int) D11 {
								return Companion_D11_.Create_DC31_((_2029_dt__update__tmp_h7).Dtor_cf58(), _2030_dt__update_hcf59_h0)
							}(_pat_let67_0)
						}(_pat_let_tv49)
					}(_pat_let66_0)
				}(_2027_v115)))).Dtor_cf59())
				_2028_v116 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)).Merge(_2028_v116)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2))
				_1998_cf19 = _1997_cf20
				(globalState).F5 = (_this).Fm11(globalState)
			}
			var _2035_v117 _dafny.Array
			_ = _2035_v117
			var _nw302 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
			_ = _nw302
			_2035_v117 = _nw302
			var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_2035_v117), 0))
			_ = _index286
			(_2035_v117).ArraySet1(_1849_v17, (_index286).Int())
			var _2036_v118 _dafny.Sequence
			_ = _2036_v118
			_2036_v118 = _dafny.SeqOf(_1812_v2)
			var _2037_v119 _dafny.Sequence
			_ = _2037_v119
			_2037_v119 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_1849_v17, _1849_v17))
			var _2038_v120 _dafny.Set
			_ = _2038_v120
			_2038_v120 = _dafny.SetOf(_1999_v101, _1999_v101)
			var _2039_v121 _dafny.MultiSet
			_ = _2039_v121
			_2039_v121 = _dafny.MultiSetOf((_2038_v120).Cardinality())
			var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_2035_v117), 0))
			_ = _index287
			var _rhs345 _dafny.Sequence = (_2036_v118).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(291))).Uint32(), func(coer122 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg122 _dafny.Int) interface{} {
					return coer122(arg122)
				}
			}((func(_2040_v17 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_2041_i12 _dafny.Int) _dafny.Sequence {
					return _2040_v17
				}
			})(_1849_v17)))).Cardinality()), _dafny.IntOfUint32((_2036_v118).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs345
			var _rhs346 _dafny.Sequence = (_2037_v119).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_2039_v121).Contains(_dafny.IntOfInt64(974)) {
					return (_2039_v121).Multiplicity(_dafny.IntOfInt64(974))
				}
				return p2
			})(), _dafny.IntOfUint32((_2037_v119).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs346
			var _lhs286 _dafny.Array = _2035_v117
			_ = _lhs286
			var _lhs287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_2035_v117), 0))
			_ = _lhs287
			_1812_v2 = _rhs345
			(_lhs286).ArraySet1(_rhs346, (_lhs287).Int())
			(globalState).F5 = Companion_Default___.SafeModuloInt(p2, _dafny.IntOfInt64(-255))
			var _2042_v122 _dafny.MultiSet
			_ = _2042_v122
			_2042_v122 = _dafny.MultiSetOf(_1997_cf20)
			var _2043_v123 _dafny.Map
			_ = _2043_v123
			_2043_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2042_v122, p1)
			_2043_v123 = (_2043_v123).Update(_dafny.MultiSetOf(_1997_cf20, _1997_cf20, (func() bool {
				if (_1809_v0).Contains(p3) {
					return (_1809_v0).Get(p3).(bool)
				}
				return p0
			})()), p3)
		} else {
			var _2044___mcc_h17 _dafny.Map = _source22.Get_().(D2_DC5).Cf12
			_ = _2044___mcc_h17
			var _2045_cf12 _dafny.Map = _2044___mcc_h17
			_ = _2045_cf12
			var _2046_v124 _dafny.Array
			_ = _2046_v124
			var _len0_50 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_50
			var _nw303 _dafny.Array
			_ = _nw303
			if _len0_50.Cmp(_dafny.Zero) == 0 {
				_nw303 = _dafny.NewArray(_len0_50)
			} else {
				var _init50 func(_dafny.Int) _dafny.Int = (func(_2047_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2048_i13 _dafny.Int) _dafny.Int {
						return (_2048_i13).Minus((_dafny.Zero).Minus((_dafny.SetOf(_2047_p2, _2047_p2, _2047_p2)).Cardinality()))
					}
				})(p2)
				_ = _init50
				var _element0_50 = _init50(_dafny.Zero)
				_ = _element0_50
				_nw303 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
				(_nw303).ArraySet1(_element0_50, 0)
				var _nativeLen0_50 = (_len0_50).Int()
				_ = _nativeLen0_50
				for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
					(_nw303).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
				}
			}
			_2046_v124 = _nw303
			var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_2046_v124), 0))
			_ = _index288
			(_2046_v124).ArraySet1(p2, (_index288).Int())
			var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_2046_v124), 0))
			_ = _index289
			(_2046_v124).ArraySet1(p2, (_index289).Int())
			var _2049_v125 _dafny.Map
			_ = _2049_v125
			_2049_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, Companion_Default___.Fm2(p2, globalState))
			var _2050_v126 _dafny.Map
			_ = _2050_v126
			_2050_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2049_v125, Companion_Default___.SafeModuloInt((_2046_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_2046_v124), 0))).Int()).(_dafny.Int), p2))
			var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_2046_v124), 0))
			_ = _index290
			(_2046_v124).ArraySet1((func() _dafny.Int {
				if (_2050_v126).Contains(_2049_v125) {
					return (_2050_v126).Get(_2049_v125).(_dafny.Int)
				}
				return Companion_Default___.SafeDivisionInt((_2046_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_2046_v124), 0))).Int()).(_dafny.Int), (_2046_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_2046_v124), 0))).Int()).(_dafny.Int))
			})(), (_index290).Int())
			var _2051_v127 D4
			_ = _2051_v127
			_2051_v127 = Companion_D4_.Create_DC13_(_dafny.SeqOf(_dafny.IntOfUint32((_1849_v17).Cardinality()), (_2046_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_2046_v124), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm17((_dafny.Zero).Minus(p2), p2, globalState)).Cardinality()), (_dafny.Zero).Minus((_2046_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_2046_v124), 0))).Int()).(_dafny.Int)), p2), false, _dafny.IntOfUint32((_1812_v2).Cardinality()))
			_2051_v127 = _2051_v127
			var _2052_v128 _dafny.Array
			_ = _2052_v128
			var _nwElement0_59 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1849_v17, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_2046_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_2046_v124), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_1849_v17).Cardinality()))).Uint32(), _1848_v16)
			_ = _nwElement0_59
			var _nw304 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(2))
			_ = _nw304
			(_nw304).ArraySet1(_nwElement0_59, 0)
			(_nw304).ArraySet1(_1849_v17, 1)
			_2052_v128 = _nw304
			var _2053_v129 D10
			_ = _2053_v129
			_2053_v129 = Companion_D10_.Create_DC29_(p2, _2052_v128, p3, p2)
			var _2054_v130 _dafny.Map
			_ = _2054_v130
			_2054_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), _2053_v129)
			var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_1810_v1), 0))
			_ = _index291
			(_1810_v1).ArraySet1((p2).Cmp((_2046_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_2046_v124), 0))).Int()).(_dafny.Int)) > 0, (_index291).Int())
			var _2055_v131 _dafny.Map
			_ = _2055_v131
			_2055_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1848_v16, (func() bool {
				if p1 {
					return p1
				}
				return p0
			})())
			var _2056_v132 _dafny.Sequence
			_ = _2056_v132
			_2056_v132 = _dafny.SeqOf(_2054_v130)
			var _2057_v133 _dafny.Map
			_ = _2057_v133
			_2057_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2046_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_2046_v124), 0))).Int()).(_dafny.Int), _2054_v130)
			var _2058_v134 _dafny.Map
			_ = _2058_v134
			_2058_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, true)
			var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_1810_v1), 0))
			_ = _index292
			var _rhs347 _dafny.Int = _dafny.IntOfInt64(326)
			_ = _rhs347
			var _rhs348 _dafny.Array = _2046_v124
			_ = _rhs348
			var _rhs349 _dafny.Map = ((func() _dafny.Map {
				if p3 {
					return ((_2056_v132).Select((Companion_Default___.SafeIndex((_2046_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_2046_v124), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2056_v132).Cardinality()))).Uint32()).(_dafny.Map)).Update(p0, _2053_v129)
				}
				return _2054_v130
			})()).Merge((func() _dafny.Map {
				if (_2057_v133).Contains(p2) {
					return (_2057_v133).Get(p2).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2053_v129)
			})())
			_ = _rhs349
			var _rhs350 bool = (Companion_D2_.Create_DC8_((func() bool {
				if (_2058_v134).Contains(_dafny.IntOfInt64(-284)) {
					return (_2058_v134).Get(_dafny.IntOfInt64(-284)).(bool)
				}
				return p1
			})(), p3)).Dtor_cf19()
			_ = _rhs350
			var _rhs351 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1848_v16, p0)
			_ = _rhs351
			var _lhs288 _dafny.Array = _1810_v1
			_ = _lhs288
			var _lhs289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_1810_v1), 0))
			_ = _lhs289
			r2 = _rhs347
			_2046_v124 = _rhs348
			_2054_v130 = _rhs349
			(_lhs288).ArraySet1(_rhs350, (_lhs289).Int())
			_2055_v131 = _rhs351
		}
		var _2059_v135 _dafny.Set
		_ = _2059_v135
		_2059_v135 = _dafny.SetOf(!(p0), p3)
		var _2060_v136 D6
		_ = _2060_v136
		_2060_v136 = Companion_D6_.Create_DC19_(p2, (_this).Fm11(globalState), ((_2059_v135).Union(_2059_v135)).Cardinality())
		var _2061_v137 _dafny.Array
		_ = _2061_v137
		var _nw305 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
		_ = _nw305
		_2061_v137 = _nw305
		var _2062_v138 _dafny.MultiSet
		_ = _2062_v138
		_2062_v138 = _dafny.MultiSetOf(_2061_v137, _2061_v137, _2061_v137)
		var _rhs352 D6 = _2060_v136
		_ = _rhs352
		var _rhs353 _dafny.Int = ((_2062_v138).Union(_2062_v138)).Cardinality()
		_ = _rhs353
		var _lhs290 *GlobalState = globalState
		_ = _lhs290
		_2060_v136 = _rhs352
		_lhs290.F5 = _rhs353
		var _2063_v139 _dafny.Array
		_ = _2063_v139
		var _nw306 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw306
		_2063_v139 = _nw306
		var _2064_v140 D2
		_ = _2064_v140
		_2064_v140 = Companion_D2_.Create_DC6_(p2, _2063_v139)
		var _2065_v141 _dafny.Map
		_ = _2065_v141
		_2065_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2064_v140, p0)
		var _2066_v142 _dafny.Sequence
		_ = _2066_v142
		_2066_v142 = _dafny.SeqOf(_2065_v141)
		r0 = ((_2065_v141).Merge(_2065_v141)).Merge((func() _dafny.Map {
			if false {
				return _2065_v141
			}
			return (_2066_v142).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2066_v142).Cardinality()))).Uint32()).(_dafny.Map)
		})())
		r1 = _1809_v0
		r2 = (_1812_v2).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1812_v2).Cardinality()))).Uint32()).(_dafny.Int)
		return r0, r1, r2
	}
}
func (_this *C8) M5(p0 bool, p1 bool, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _2067_v0 _dafny.Int
		_ = _2067_v0
		_2067_v0 = _dafny.IntOfInt64(372)
		var _2068_v1 _dafny.Map
		_ = _2068_v1
		_2068_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2067_v0, _dafny.IntOfInt64(100))
		var _2069_v2 _dafny.Map
		_ = _2069_v2
		_2069_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2067_v0, _2068_v1)
		r2 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2067_v0, _2068_v1)).Merge(_2069_v2)).Cardinality()
		_2067_v0 = _2067_v0
		var _2070_v3 _dafny.MultiSet
		_ = _2070_v3
		_2070_v3 = _dafny.MultiSetOf(_dafny.CodePoint('h'))
		var _2071_v4 D1
		_ = _2071_v4
		_2071_v4 = Companion_D1_.Create_DC3_(_2070_v3)
		var _2072_v5 _dafny.CodePoint
		_ = _2072_v5
		_2072_v5 = _dafny.CodePoint('o')
		var _2073_v6 _dafny.Array
		_ = _2073_v6
		var _nwElement0_60 D1 = _2071_v4
		_ = _nwElement0_60
		var _nw307 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(23))
		_ = _nw307
		(_nw307).ArraySet1(_nwElement0_60, 0)
		(_nw307).ArraySet1(_2071_v4, 1)
		(_nw307).ArraySet1(_2071_v4, 2)
		(_nw307).ArraySet1(_2071_v4, 3)
		(_nw307).ArraySet1(_2071_v4, 4)
		(_nw307).ArraySet1(_2071_v4, 5)
		(_nw307).ArraySet1(_2071_v4, 6)
		(_nw307).ArraySet1(_2071_v4, 7)
		(_nw307).ArraySet1(Companion_D1_.Create_DC3_(_2070_v3), 8)
		(_nw307).ArraySet1(_2071_v4, 9)
		(_nw307).ArraySet1(Companion_D1_.Create_DC3_(_2070_v3), 10)
		(_nw307).ArraySet1(_2071_v4, 11)
		(_nw307).ArraySet1((func() D1 {
			if p0 {
				return _2071_v4
			}
			return _2071_v4
		})(), 12)
		(_nw307).ArraySet1(_2071_v4, 13)
		(_nw307).ArraySet1(_2071_v4, 14)
		(_nw307).ArraySet1(Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_2072_v5)), 15)
		(_nw307).ArraySet1(_2071_v4, 16)
		(_nw307).ArraySet1(_2071_v4, 17)
		(_nw307).ArraySet1(_2071_v4, 18)
		(_nw307).ArraySet1(_2071_v4, 19)
		(_nw307).ArraySet1((func() D1 {
			if p0 {
				return _2071_v4
			}
			return _2071_v4
		})(), 20)
		(_nw307).ArraySet1(_2071_v4, 21)
		(_nw307).ArraySet1(_2071_v4, 22)
		_2073_v6 = _nw307
		for _iter56 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2073_v6), 0))); ; {
			_guard_loop_3, _ok56 := _iter56()
			if !_ok56 {
				break
			}
			var _2074_i0 _dafny.Int
			_2074_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_2074_i0).Sign() != -1) && ((_2074_i0).Cmp(_dafny.ArrayLen((_2073_v6), 0)) < 0)) {
				(_2073_v6).ArraySet1((func() D1 {
					if p0 {
						return _2071_v4
					}
					return _2071_v4
				})(), (_2074_i0).Int())
			}
		}
		var _2075_v7 _dafny.Array
		_ = _2075_v7
		var _len0_51 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_51
		var _nw308 _dafny.Array
		_ = _nw308
		if _len0_51.Cmp(_dafny.Zero) == 0 {
			_nw308 = _dafny.NewArray(_len0_51)
		} else {
			var _init51 func(_dafny.Int) _dafny.Sequence = func(_2076_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("sggnxrdx")
			}
			_ = _init51
			var _element0_51 = _init51(_dafny.Zero)
			_ = _element0_51
			_nw308 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
			(_nw308).ArraySet1(_element0_51, 0)
			var _nativeLen0_51 = (_len0_51).Int()
			_ = _nativeLen0_51
			for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
				(_nw308).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
			}
		}
		_2075_v7 = _nw308
		var _2077_v8 D10
		_ = _2077_v8
		_2077_v8 = Companion_D10_.Create_DC29_(_2067_v0, _2075_v7, p0, _2067_v0)
		var _2078_v9 _dafny.Array
		_ = _2078_v9
		var _nwElement0_61 _dafny.Int = (_2077_v8).Dtor_cf53()
		_ = _nwElement0_61
		var _nw309 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.One)
		_ = _nw309
		(_nw309).ArraySet1(_nwElement0_61, 0)
		_2078_v9 = _nw309
		var _2079_v10 _dafny.Map
		_ = _2079_v10
		_2079_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2078_v9, p0)
		_2079_v10 = (_2079_v10).Update(_2078_v9, p0)
		var _2080_v11 _dafny.Set
		_ = _2080_v11
		_2080_v11 = _dafny.SetOf(p1, p1, (p0) || (!(false)))
		_2080_v11 = (Companion_D11_.Create_DC33_(_dafny.IntOfInt64(-74), p0, (_this).Fm11(globalState), p0, _2080_v11)).Dtor_cf66()
		var _hi13 _dafny.Int = _2067_v0
		_ = _hi13
		for _2081_i2 := (_dafny.Zero).Minus(_2067_v0); _2081_i2.Cmp(_hi13) < 0; _2081_i2 = _2081_i2.Plus(_dafny.One) {
			var _2082_v12 _dafny.Sequence
			_ = _2082_v12
			_2082_v12 = _dafny.SeqOf(_dafny.IntOfInt64(428), _2067_v0)
			var _2083_v13 D11
			_ = _2083_v13
			_2083_v13 = Companion_D11_.Create_DC31_(_dafny.IntOfInt64(67), (_2082_v12).Select((Companion_Default___.SafeIndex(_2081_i2, _dafny.IntOfUint32((_2082_v12).Cardinality()))).Uint32()).(_dafny.Int))
			var _pat_let_tv52 = _2068_v1
			_ = _pat_let_tv52
			var _rhs354 D11 = func(_pat_let70_0 D11) D11 {
				return func(_2084_dt__update__tmp_h0 D11) D11 {
					return func(_pat_let71_0 _dafny.Int) D11 {
						return func(_2085_dt__update_hcf59_h0 _dafny.Int) D11 {
							return Companion_D11_.Create_DC31_((_2084_dt__update__tmp_h0).Dtor_cf58(), _2085_dt__update_hcf59_h0)
						}(_pat_let71_0)
					}((_2081_i2).Plus((_pat_let_tv52).Cardinality()))
				}(_pat_let70_0)
			}(_2083_v13)
			_ = _rhs354
			var _rhs355 _dafny.Int = (_2067_v0).Times(_2067_v0)
			_ = _rhs355
			_2083_v13 = _rhs354
			r2 = _rhs355
			var _2086_v14 _dafny.Sequence
			_ = _2086_v14
			_2086_v14 = _dafny.SeqOf(_2070_v3, _2070_v3)
			var _2087_v15 _dafny.Sequence
			_ = _2087_v15
			_2087_v15 = _dafny.UnicodeSeqOfUtf8Bytes("iguyj")
			var _2088_v16 D4
			_ = _2088_v16
			_2088_v16 = Companion_D4_.Create_DC12_(_dafny.Companion_Sequence_.Update(_2086_v14, (Companion_Default___.SafeIndex(_2081_i2, _dafny.IntOfUint32((_2086_v14).Cardinality()))).Uint32(), (_2070_v3).Update(_2072_v5, Companion_Default___.Abs(_dafny.IntOfUint32((_2087_v15).Cardinality())))))
			_2088_v16 = _2088_v16
			var _2089_v17 _dafny.Array
			_ = _2089_v17
			var _nw310 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
			_ = _nw310
			_2089_v17 = _nw310
			var _2090_v18 _dafny.Array
			_ = _2090_v18
			_2090_v18 = _2089_v17
			var _2091_v19 _dafny.Map
			_ = _2091_v19
			_2091_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2067_v0, _dafny.UnicodeSeqOfUtf8Bytes("by"))
			var _2092_v20 *C3
			_ = _2092_v20
			var _nw311 *C3 = New_C3_()
			_ = _nw311
			_nw311.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(87), _2067_v0), (_2090_v18), _2091_v19, _2081_i2)
			_2092_v20 = _nw311
			_2067_v0 = _2067_v0
		}
		r0 = p1
		var _2093_v21 _dafny.Map
		_ = _2093_v21
		_2093_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2067_v0, p1)
		var _2094_v22 _dafny.Map
		_ = _2094_v22
		_2094_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2093_v21, p1)
		var _2095_v23 D3
		_ = _2095_v23
		_2095_v23 = Companion_D3_.Create_DC10_(_2094_v22)
		var _pat_let_tv53 = p0
		_ = _pat_let_tv53
		var _pat_let_tv54 = p1
		_ = _pat_let_tv54
		var _pat_let_tv55 = p0
		_ = _pat_let_tv55
		r1 = !(func(_source25 D3) bool {
			if _source25.Is_DC10() {
				var _2096___mcc_h0 _dafny.Map = _source25.Get_().(D3_DC10).Cf22
				_ = _2096___mcc_h0
				var _2097_cf22 _dafny.Map = _2096___mcc_h0
				_ = _2097_cf22
				return _pat_let_tv53
			} else if _source25.Is_DC9() {
				var _2098___mcc_h1 _dafny.Sequence = _source25.Get_().(D3_DC9).Cf21
				_ = _2098___mcc_h1
				var _2099_cf21 _dafny.Sequence = _2098___mcc_h1
				_ = _2099_cf21
				return _pat_let_tv54
			} else {
				var _2100___mcc_h2 D3 = _source25.Get_().(D3_DC11).Cf23
				_ = _2100___mcc_h2
				var _2101_cf23 D3 = _2100___mcc_h2
				_ = _2101_cf23
				return _pat_let_tv55
			}
		}(_2095_v23))
		r2 = (_dafny.Zero).Minus(_2067_v0)
		return r0, r1, r2
	}
}

// End of class C8

// Definition of class C9
type C9 struct {
	_f6 _dafny.Int
}

func New_C9_() *C9 {
	_this := C9{}

	_this._f6 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C9{}
var _ _dafny.TraitOffspring = &C9{}

func (_this *C9) F6() _dafny.Int {
	return _this._f6
}
func (_this *C9) Ctor__(f6 _dafny.Int) {
	{
		(_this)._f6 = f6
	}
}
func (_this *C9) Fm5(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('c'), _dafny.CodePoint('v'), _dafny.CodePoint('w'))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('q'), _dafny.CodePoint('e'))), _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(26))).Uint32(), func(coer123 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg123 _dafny.Int) interface{} {
				return coer123(arg123)
			}
		}(func(_2102_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		}))), _dafny.MultiSetOf(_dafny.CodePoint('u'), _dafny.CodePoint('a')), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('r'), _dafny.CodePoint('a'))), (Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_dafny.CodePoint('v'), _dafny.CodePoint('w')))).Dtor_cf6(), _dafny.MultiSetOf(_dafny.CodePoint('p')))))
	}
}
func (_this *C9) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.SeqOf((_this).F6()))
	}
}
func (_this *C9) M3(p0 bool, p1 _dafny.Array, p2 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _2103_v0 _dafny.Array
		_ = _2103_v0
		var _nw312 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
		_ = _nw312
		_2103_v0 = _nw312
		var _2104_v1 _dafny.Sequence
		_ = _2104_v1
		_2104_v1 = _dafny.UnicodeSeqOfUtf8Bytes("hbqce")
		var _2105_v2 _dafny.Set
		_ = _2105_v2
		_2105_v2 = _dafny.SetOf((_this).F6(), _dafny.IntOfUint32((_2104_v1).Cardinality()))
		var _2106_v3 _dafny.Set
		_ = _2106_v3
		_2106_v3 = _dafny.SetOf(false)
		var _2107_v4 _dafny.Sequence
		_ = _2107_v4
		_2107_v4 = _dafny.SeqOf((_2106_v3).Cardinality())
		var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_2103_v0), 0))
		_ = _index293
		(_2103_v0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_2105_v2).Cardinality()), _dafny.Companion_Sequence_.Update(_2107_v4, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2107_v4).Cardinality()))).Uint32(), (_this).F6())), (_index293).Int())
		var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_2103_v0), 0))
		_ = _index294
		(_2103_v0).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2107_v4, _2107_v4), _2107_v4), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2107_v4, _2107_v4), _2107_v4)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_2104_v1).Cardinality())), (_index294).Int())
		var _2108_v5 _dafny.Sequence
		_ = _2108_v5
		_2108_v5 = _dafny.SeqOf(p2, p2)
		(globalState).F0 = !(((_this).F6()).Cmp((func() _dafny.Int {
			if p0 {
				return _dafny.IntOfUint32((_2108_v5).Cardinality())
			}
			return (_this).F6()
		})()) >= 0)
		var _2109_v6 D0
		_ = _2109_v6
		_2109_v6 = Companion_D0_.Create_DC0_(_2108_v5, (_this).F6())
		var _source26 D0 = _2109_v6
		_ = _source26
		if _source26.Is_DC0() {
			var _2110___mcc_h0 _dafny.Sequence = _source26.Get_().(D0_DC0).Cf0
			_ = _2110___mcc_h0
			var _2111___mcc_h1 _dafny.Int = _source26.Get_().(D0_DC0).Cf1
			_ = _2111___mcc_h1
			var _2112_cf1 _dafny.Int = _2111___mcc_h1
			_ = _2112_cf1
			var _2113_cf0 _dafny.Sequence = _2110___mcc_h0
			_ = _2113_cf0
			var _2114_v7 _dafny.CodePoint
			_ = _2114_v7
			_2114_v7 = _dafny.CodePoint('e')
			var _2115_v8 _dafny.Map
			_ = _2115_v8
			_2115_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if p0 {
					return _2112_cf1
				}
				return (_this).F6()
			})(), (_this).F6())
			var _2116_v9 _dafny.Map
			_ = _2116_v9
			_2116_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.CodePoint('v'))
			var _rhs356 bool = true
			_ = _rhs356
			var _rhs357 bool = p0
			_ = _rhs357
			var _rhs358 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_2116_v9).Contains(_dafny.Companion_Sequence_.IsPrefixOf(_2104_v1, _2104_v1)) {
					return (_2116_v9).Get(_dafny.Companion_Sequence_.IsPrefixOf(_2104_v1, _2104_v1)).(_dafny.CodePoint)
				}
				return _2114_v7
			})()
			_ = _rhs358
			var _rhs359 _dafny.Map = (_2115_v8).Update(Companion_Default___.Fm2(_2112_cf1, globalState), (_this).F6())
			_ = _rhs359
			var _lhs291 *GlobalState = globalState
			_ = _lhs291
			var _lhs292 *GlobalState = globalState
			_ = _lhs292
			_lhs291.F0 = _rhs356
			_lhs292.F0 = _rhs357
			_2114_v7 = _rhs358
			_2115_v8 = _rhs359
			(globalState).F0 = false
			if p0 {
				var _2117_v10 _dafny.Sequence
				_ = _2117_v10
				_2117_v10 = _dafny.SeqOf(_2104_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-641))).Uint32(), func(coer124 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg124 _dafny.Int) interface{} {
						return coer124(arg124)
					}
				}((func(_2118_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2119_i1 _dafny.Int) _dafny.CodePoint {
						return _2118_v7
					}
				})(_2114_v7))), _dafny.Companion_Sequence_.Update(_2104_v1, (Companion_Default___.SafeIndex(_2112_cf1, _dafny.IntOfUint32((_2104_v1).Cardinality()))).Uint32(), _2114_v7))
				_2104_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(885))).Uint32(), func(coer125 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg125 _dafny.Int) interface{} {
						return coer125(arg125)
					}
				}((func(_2120_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2121_i0 _dafny.Int) _dafny.CodePoint {
						return _2120_v7
					}
				})(_2114_v7))), (_2117_v10).Select((Companion_Default___.SafeIndex((func() _dafny.Map {
					var _coll53 = _dafny.NewMapBuilder()
					_ = _coll53
					for _iter57 := _dafny.Iterate((func() _dafny.Map {
						var _coll54 = _dafny.NewMapBuilder()
						_ = _coll54
						for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(381), _dafny.IntOfInt64(-819))); ; {
							_compr_54, _ok58 := _iter58()
							if !_ok58 {
								break
							}
							var _2122_v12 _dafny.Int
							_2122_v12 = interface{}(_compr_54).(_dafny.Int)
							if ((_dafny.IntOfInt64(381)).Cmp(_2122_v12) <= 0) && ((_2122_v12).Cmp(_dafny.IntOfInt64(-819)) < 0) {
								_coll54.Add((_2122_v12).Minus((_this).F6()), (_dafny.MultiSetOf(!(p0), p2)).Cardinality())
							}
						}
						return _coll54.ToMap()
					}()).Keys().Elements()); ; {
						_compr_53, _ok57 := _iter57()
						if !_ok57 {
							break
						}
						var _2123_v11 _dafny.Int
						_2123_v11 = interface{}(_compr_53).(_dafny.Int)
						if (func() _dafny.Map {
							var _coll55 = _dafny.NewMapBuilder()
							_ = _coll55
							for _iter59 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(381), _dafny.IntOfInt64(-819))); ; {
								_compr_55, _ok59 := _iter59()
								if !_ok59 {
									break
								}
								var _2122_v12 _dafny.Int
								_2122_v12 = interface{}(_compr_55).(_dafny.Int)
								if ((_dafny.IntOfInt64(381)).Cmp(_2122_v12) <= 0) && ((_2122_v12).Cmp(_dafny.IntOfInt64(-819)) < 0) {
									_coll55.Add((_2122_v12).Minus((_this).F6()), (_dafny.MultiSetOf(!(p0), p2)).Cardinality())
								}
							}
							return _coll55.ToMap()
						}()).Contains(_2123_v11) {
							_coll53.Add((_2123_v11).Minus(_2112_cf1), _2112_cf1)
						}
					}
					return _coll53.ToMap()
				}()).Cardinality(), _dafny.IntOfUint32((_2117_v10).Cardinality()))).Uint32()).(_dafny.Sequence))
				(globalState).F0 = p2
				var _2124_v13 _dafny.MultiSet
				_ = _2124_v13
				_2124_v13 = _dafny.MultiSetOf(p1, p1)
				var _2125_v14 _dafny.MultiSet
				_ = _2125_v14
				_2125_v14 = _dafny.MultiSetOf(_2104_v1)
				var _2126_v15 _dafny.Map
				_ = _2126_v15
				_2126_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _2125_v14)
				var _2127_v16 _dafny.Map
				_ = _2127_v16
				_2127_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2124_v13).IsProperSubsetOf(_2124_v13), (func() _dafny.MultiSet {
					if (_2126_v15).Contains(p0) {
						return (_2126_v15).Get(p0).(_dafny.MultiSet)
					}
					return _dafny.MultiSetOf(_2104_v1, _2104_v1, _2104_v1)
				})())
				_2127_v16 = (_2127_v16).Update((Companion_Default___.Fm3(globalState)) == (p0), _2125_v14)
				(globalState).F0 = (p0) || (false)
				var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((p1), 0))
				_ = _index295
				(p1).ArraySet1((_this).F6(), (_index295).Int())
				var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((p1), 0))
				_ = _index296
				(p1).ArraySet1(Companion_Default___.Fm2(Companion_Default___.Fm2(_2112_cf1, globalState), globalState), (_index296).Int())
			} else {
				var _2128_v17 _dafny.MultiSet
				_ = _2128_v17
				_2128_v17 = _dafny.MultiSetOf((_dafny.SetOf(_2112_cf1, (_this).F6())).Cardinality(), (_this).F6(), (_this).F6(), _2112_cf1)
				(globalState).F0 = (_dafny.MultiSetOf((_this).F6(), (_this).F6())).IsDisjointFrom(_2128_v17)
				_2109_v6 = _2109_v6
				_2115_v8 = (_2115_v8).Update((_this).F6(), (_this).F6())
				var _2129_v18 _dafny.Array
				_ = _2129_v18
				var _len0_52 _dafny.Int = _dafny.One
				_ = _len0_52
				var _nw313 _dafny.Array
				_ = _nw313
				if _len0_52.Cmp(_dafny.Zero) == 0 {
					_nw313 = _dafny.NewArray(_len0_52)
				} else {
					var _init52 func(_dafny.Int) bool = func(_2130_i2 _dafny.Int) bool {
						return false
					}
					_ = _init52
					var _element0_52 = _init52(_dafny.Zero)
					_ = _element0_52
					_nw313 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
					(_nw313).ArraySet1(_element0_52, 0)
					var _nativeLen0_52 = (_len0_52).Int()
					_ = _nativeLen0_52
					for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
						(_nw313).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
					}
				}
				_2129_v18 = _nw313
				var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_2129_v18), 0))
				_ = _index297
				(_2129_v18).ArraySet1(!(p0), (_index297).Int())
				var _2131_v19 _dafny.Map
				_ = _2131_v19
				_2131_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2112_cf1, false)
				var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_2129_v18), 0))
				_ = _index298
				(_2129_v18).ArraySet1(((_2131_v19).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2112_cf1, p0))).Equals((_2131_v19).Merge(_2131_v19)), (_index298).Int())
				var _2132_v20 _dafny.MultiSet
				_ = _2132_v20
				_2132_v20 = _dafny.MultiSetOf(p2, (func() bool {
					if (_2131_v19).Contains(_dafny.IntOfInt64(158)) {
						return (_2131_v19).Get(_dafny.IntOfInt64(158)).(bool)
					}
					return p0
				})(), p2)
				var _2133_v21 D1
				_ = _2133_v21
				_2133_v21 = Companion_D1_.Create_DC4_(_2109_v6, (_2132_v20).Cardinality(), _2104_v1, (_2108_v5).Select((Companion_Default___.SafeIndex(_2112_cf1, _dafny.IntOfUint32((_2108_v5).Cardinality()))).Uint32()).(bool), _2112_cf1)
				var _2134_v22 _dafny.Map
				_ = _2134_v22
				_2134_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.Fm2(_dafny.IntOfInt64(971), globalState)), _2133_v21)
				_2134_v22 = _2134_v22
			}
			var _2135_v23 _dafny.Array
			_ = _2135_v23
			var _nw314 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
			_ = _nw314
			_2135_v23 = _nw314
			var _2136_v24 _dafny.Array
			_ = _2136_v24
			var _nw315 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(6))
			_ = _nw315
			_2136_v24 = _nw315
			var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_2135_v23), 0))
			_ = _index299
			(_2135_v23).ArraySet1(_2136_v24, (_index299).Int())
			var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_2135_v23), 0))
			_ = _index300
			var _len0_53 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_53
			var _nw316 _dafny.Array
			_ = _nw316
			if _len0_53.Cmp(_dafny.Zero) == 0 {
				_nw316 = _dafny.NewArray(_len0_53)
			} else {
				var _init53 func(_dafny.Int) D0 = (func(_2137_v6 D0) func(_dafny.Int) D0 {
					return func(_2138_i3 _dafny.Int) D0 {
						return _2137_v6
					}
				})(_2109_v6)
				_ = _init53
				var _element0_53 = _init53(_dafny.Zero)
				_ = _element0_53
				_nw316 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
				(_nw316).ArraySet1(_element0_53, 0)
				var _nativeLen0_53 = (_len0_53).Int()
				_ = _nativeLen0_53
				for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
					(_nw316).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
				}
			}
			(_2135_v23).ArraySet1(_nw316, (_index300).Int())
		} else if _source26.Is_DC1() {
			var _2139___mcc_h2 bool = _source26.Get_().(D0_DC1).Cf2
			_ = _2139___mcc_h2
			var _2140___mcc_h3 _dafny.Int = _source26.Get_().(D0_DC1).Cf3
			_ = _2140___mcc_h3
			var _2141___mcc_h4 _dafny.CodePoint = _source26.Get_().(D0_DC1).Cf4
			_ = _2141___mcc_h4
			var _2142_cf4 _dafny.CodePoint = _2141___mcc_h4
			_ = _2142_cf4
			var _2143_cf3 _dafny.Int = _2140___mcc_h3
			_ = _2143_cf3
			var _2144_cf2 bool = _2139___mcc_h2
			_ = _2144_cf2
			_2143_cf3 = (_this).F6()
			r0 = !(p0) || (p0)
			var _2145_v25 _dafny.Array
			_ = _2145_v25
			var _len0_54 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_54
			var _nw317 _dafny.Array
			_ = _nw317
			if _len0_54.Cmp(_dafny.Zero) == 0 {
				_nw317 = _dafny.NewArray(_len0_54)
			} else {
				var _init54 func(_dafny.Int) _dafny.Sequence = (func(_2146_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_2147_i4 _dafny.Int) _dafny.Sequence {
						return _2146_v1
					}
				})(_2104_v1)
				_ = _init54
				var _element0_54 = _init54(_dafny.Zero)
				_ = _element0_54
				_nw317 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
				(_nw317).ArraySet1(_element0_54, 0)
				var _nativeLen0_54 = (_len0_54).Int()
				_ = _nativeLen0_54
				for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
					(_nw317).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
				}
			}
			_2145_v25 = _nw317
			var _2148_v26 _dafny.Map
			_ = _2148_v26
			_2148_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2145_v25, true)
			_2148_v26 = (_2148_v26).Update(_2145_v25, p0)
			_2143_cf3 = (func() _dafny.Int {
				if p0 {
					return Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(p2, p2)).Cardinality()), _2143_cf3)
				}
				return _2143_cf3
			})()
		} else {
			var _2149___mcc_h5 D0 = _source26.Get_().(D0_DC2).Cf5
			_ = _2149___mcc_h5
			var _2150_cf5 D0 = _2149___mcc_h5
			_ = _2150_cf5
			(globalState).F5 = (_this).F6()
			(globalState).F5 = (_this).F6()
			var _2151_v27 _dafny.Map
			_ = _2151_v27
			_2151_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
			var _2152_v31 _dafny.Array
			_ = _2152_v31
			var _nwElement0_62 _dafny.Map = _2151_v27
			_ = _nwElement0_62
			var _nw318 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(16))
			_ = _nw318
			(_nw318).ArraySet1(_nwElement0_62, 0)
			(_nw318).ArraySet1((_2151_v27).Update((_this).F6(), _dafny.IntOfInt64(-510)), 1)
			(_nw318).ArraySet1(_2151_v27, 2)
			(_nw318).ArraySet1(_2151_v27, 3)
			(_nw318).ArraySet1((_2151_v27).Merge(_2151_v27), 4)
			(_nw318).ArraySet1(_2151_v27, 5)
			(_nw318).ArraySet1((func() _dafny.Map {
				if p2 {
					return _2151_v27
				}
				return _2151_v27
			})(), 6)
			(_nw318).ArraySet1(func() _dafny.Map {
				var _coll56 = _dafny.NewMapBuilder()
				_ = _coll56
				for _iter60 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-1), _dafny.IntOfInt64(458))); ; {
					_compr_56, _ok60 := _iter60()
					if !_ok60 {
						break
					}
					var _2153_v28 _dafny.Int
					_2153_v28 = interface{}(_compr_56).(_dafny.Int)
					if ((_dafny.IntOfInt64(-1)).Cmp(_2153_v28) <= 0) && ((_2153_v28).Cmp(_dafny.IntOfInt64(458)) < 0) {
						_coll56.Add(Companion_Default___.SafeModuloInt(_2153_v28, _dafny.IntOfInt64(76)), (_this).F6())
					}
				}
				return _coll56.ToMap()
			}(), 7)
			(_nw318).ArraySet1(func() _dafny.Map {
				var _coll57 = _dafny.NewMapBuilder()
				_ = _coll57
				for _iter61 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(52), _dafny.IntOfInt64(438))); ; {
					_compr_57, _ok61 := _iter61()
					if !_ok61 {
						break
					}
					var _2154_v29 _dafny.Int
					_2154_v29 = interface{}(_compr_57).(_dafny.Int)
					if ((_dafny.IntOfInt64(52)).Cmp(_2154_v29) <= 0) && ((_2154_v29).Cmp(_dafny.IntOfInt64(438)) < 0) {
						_coll57.Add((_2154_v29).Plus((_this).F6()), (_this).F6())
					}
				}
				return _coll57.ToMap()
			}(), 8)
			(_nw318).ArraySet1(Companion_Default___.Fm9(globalState), 9)
			(_nw318).ArraySet1((_2151_v27).Merge((_2151_v27).Update(_dafny.IntOfUint32((_2108_v5).Cardinality()), (_this).F6())), 10)
			(_nw318).ArraySet1(func() _dafny.Map {
				var _coll58 = _dafny.NewMapBuilder()
				_ = _coll58
				for _iter62 := _dafny.Iterate((_2107_v4).Elements()); ; {
					_compr_58, _ok62 := _iter62()
					if !_ok62 {
						break
					}
					var _2155_v30 _dafny.Int
					_2155_v30 = interface{}(_compr_58).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_2107_v4, _2155_v30) {
						_coll58.Add((_2155_v30).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((_2103_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_2103_v0), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32(((_2103_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_2103_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.SeqOf(p2, p2)).Cardinality())), (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p0)).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_2103_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_2103_v0), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32(((_2103_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_2103_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.SeqOf(p2, p2)).Cardinality()))).Cardinality()))).Uint32(), (_this).F6())).Cardinality())), (_this).F6())
					}
				}
				return _coll58.ToMap()
			}(), 11)
			(_nw318).ArraySet1(_2151_v27, 12)
			(_nw318).ArraySet1((Companion_D2_.Create_DC5_(_2151_v27)).Dtor_cf12(), 13)
			(_nw318).ArraySet1(_2151_v27, 14)
			(_nw318).ArraySet1((_2151_v27).Merge(_2151_v27), 15)
			_2152_v31 = _nw318
			_2152_v31 = (func() _dafny.Array {
				if (p2) || (!(p2)) {
					return _2152_v31
				}
				return _2152_v31
			})()
			var _rhs360 _dafny.Int = (_2107_v4).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F6()), _dafny.IntOfUint32((_2107_v4).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs360
			var _rhs361 bool = p2
			_ = _rhs361
			var _rhs362 _dafny.Sequence = _2104_v1
			_ = _rhs362
			var _rhs363 _dafny.Int = (_this).F6()
			_ = _rhs363
			var _lhs293 *GlobalState = globalState
			_ = _lhs293
			var _lhs294 *GlobalState = globalState
			_ = _lhs294
			_lhs293.F5 = _rhs360
			r0 = _rhs361
			_2104_v1 = _rhs362
			_lhs294.F5 = _rhs363
		}
		var _2156_v32 _dafny.CodePoint
		_ = _2156_v32
		_2156_v32 = _dafny.CodePoint('f')
		var _2157_v33 D0
		_ = _2157_v33
		_2157_v33 = Companion_D0_.Create_DC1_(p0, (_this).F6(), _2156_v32)
		var _source27 D0 = (func() D0 {
			if p0 {
				return _2157_v33
			}
			return _2157_v33
		})()
		_ = _source27
		if _source27.Is_DC0() {
			var _2158___mcc_h6 _dafny.Sequence = _source27.Get_().(D0_DC0).Cf0
			_ = _2158___mcc_h6
			var _2159___mcc_h7 _dafny.Int = _source27.Get_().(D0_DC0).Cf1
			_ = _2159___mcc_h7
			var _2160_cf1 _dafny.Int = _2159___mcc_h7
			_ = _2160_cf1
			var _2161_cf0 _dafny.Sequence = _2158___mcc_h6
			_ = _2161_cf0
			_2160_cf1 = (_this).F6()
			var _2162_v34 _dafny.MultiSet
			_ = _2162_v34
			_2162_v34 = _dafny.MultiSetOf(_2156_v32)
			var _2163_v35 _dafny.Array
			_ = _2163_v35
			var _nwElement0_63 bool = p2
			_ = _nwElement0_63
			var _nw319 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(28))
			_ = _nw319
			(_nw319).ArraySet1(_nwElement0_63, 0)
			(_nw319).ArraySet1(!(p2), 1)
			(_nw319).ArraySet1(p2, 2)
			(_nw319).ArraySet1(p0, 3)
			(_nw319).ArraySet1(p2, 4)
			(_nw319).ArraySet1(p0, 5)
			(_nw319).ArraySet1(p2, 6)
			(_nw319).ArraySet1(p2, 7)
			(_nw319).ArraySet1(false, 8)
			(_nw319).ArraySet1(p2, 9)
			(_nw319).ArraySet1(Companion_Default___.Fm3(globalState), 10)
			(_nw319).ArraySet1(Companion_Default___.Fm3(globalState), 11)
			(_nw319).ArraySet1(p2, 12)
			(_nw319).ArraySet1(false, 13)
			(_nw319).ArraySet1(p0, 14)
			(_nw319).ArraySet1(p2, 15)
			(_nw319).ArraySet1(p0, 16)
			(_nw319).ArraySet1(p2, 17)
			(_nw319).ArraySet1(p2, 18)
			(_nw319).ArraySet1(p0, 19)
			(_nw319).ArraySet1(p0, 20)
			(_nw319).ArraySet1(p0, 21)
			(_nw319).ArraySet1(p2, 22)
			(_nw319).ArraySet1(p0, 23)
			(_nw319).ArraySet1(p2, 24)
			(_nw319).ArraySet1(p2, 25)
			(_nw319).ArraySet1(p0, 26)
			(_nw319).ArraySet1(p2, 27)
			_2163_v35 = _nw319
			var _2164_v36 _dafny.Map
			_ = _2164_v36
			_2164_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _2104_v1)
			var _2165_v37 _dafny.MultiSet
			_ = _2165_v37
			_2165_v37 = _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_2107_v4))
			var _2166_v38 *C4
			_ = _2166_v38
			var _nw320 *C4 = New_C4_()
			_ = _nw320
			_nw320.Ctor__(_2160_cf1, (func() _dafny.Int {
				if (_2162_v34).Contains(_2156_v32) {
					return (_2162_v34).Multiplicity(_2156_v32)
				}
				return _dafny.IntOfUint32((_2104_v1).Cardinality())
			})(), (func() _dafny.Array {
				if p0 {
					return _2163_v35
				}
				return _2163_v35
			})(), (_2164_v36).Merge(_2164_v36), _2160_cf1, _2165_v37)
			_2166_v38 = _nw320
			var _2167_v39 _dafny.Map
			_ = _2167_v39
			_2167_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p2)
			var _2168_v40 D1
			_ = _2168_v40
			_2168_v40 = Companion_D1_.Create_DC4_(_2109_v6, Companion_Default___.Fm2((_2167_v39).Cardinality(), globalState), _2104_v1, false, (_2166_v38).F18())
			var _2169_v41 _dafny.Map
			_ = _2169_v41
			_2169_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2168_v40, _2163_v35)
			(globalState).F5 = (_dafny.Zero).Minus((_2169_v41).Cardinality())
			(globalState).F0 = (p2) == (p2)
		} else if _source27.Is_DC1() {
			var _2170___mcc_h8 bool = _source27.Get_().(D0_DC1).Cf2
			_ = _2170___mcc_h8
			var _2171___mcc_h9 _dafny.Int = _source27.Get_().(D0_DC1).Cf3
			_ = _2171___mcc_h9
			var _2172___mcc_h10 _dafny.CodePoint = _source27.Get_().(D0_DC1).Cf4
			_ = _2172___mcc_h10
			var _2173_cf4 _dafny.CodePoint = _2172___mcc_h10
			_ = _2173_cf4
			var _2174_cf3 _dafny.Int = _2171___mcc_h9
			_ = _2174_cf3
			var _2175_cf2 bool = _2170___mcc_h8
			_ = _2175_cf2
			var _2176_v42 _dafny.Map
			_ = _2176_v42
			_2176_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2157_v33, !(_2175_cf2))
			(globalState).F5 = ((_dafny.Zero).Minus((_this).F6())).Times((_2176_v42).Cardinality())
			_2174_cf3 = (func() _dafny.Int {
				if (_2108_v5).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2108_v5).Cardinality()))).Uint32()).(bool) {
					return ((_this).F6()).Minus(_dafny.IntOfInt64(169))
				}
				return (_this).F6()
			})()
			var _2177_v43 _dafny.Array
			_ = _2177_v43
			var _nw321 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw321
			_2177_v43 = _nw321
			_2177_v43 = p1
			var _2178_v44 _dafny.Array
			_ = _2178_v44
			var _len0_55 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_55
			var _nw322 _dafny.Array
			_ = _nw322
			if _len0_55.Cmp(_dafny.Zero) == 0 {
				_nw322 = _dafny.NewArray(_len0_55)
			} else {
				var _init55 func(_dafny.Int) _dafny.Map = (func(_2179_p2 bool, _2180_p0 bool) func(_dafny.Int) _dafny.Map {
					return func(_2181_i5 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_2179_p2)), _2180_p0)
					}
				})(p2, p0)
				_ = _init55
				var _element0_55 = _init55(_dafny.Zero)
				_ = _element0_55
				_nw322 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
				(_nw322).ArraySet1(_element0_55, 0)
				var _nativeLen0_55 = (_len0_55).Int()
				_ = _nativeLen0_55
				for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
					(_nw322).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
				}
			}
			_2178_v44 = _nw322
			var _2182_v45 _dafny.Map
			_ = _2182_v45
			_2182_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(globalState), p2)
			var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_2178_v44), 0))
			_ = _index301
			(_2178_v44).ArraySet1(_2182_v45, (_index301).Int())
			var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_2178_v44), 0))
			_ = _index302
			(_2178_v44).ArraySet1(_2182_v45, (_index302).Int())
		} else {
			var _2183___mcc_h11 D0 = _source27.Get_().(D0_DC2).Cf5
			_ = _2183___mcc_h11
			var _2184_cf5 D0 = _2183___mcc_h11
			_ = _2184_cf5
			(globalState).F5 = ((_this).F6()).Plus((_this).F6())
			var _2185_v46 _dafny.Map
			_ = _2185_v46
			_2185_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2104_v1, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2104_v1).Cardinality()))).Uint32(), _2156_v32)).Cardinality()))
			_2185_v46 = (_2185_v46).Update((_this).F6(), (_this).F6())
			var _2186_v47 _dafny.Map
			_ = _2186_v47
			_2186_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(789), !(p0))
			var _rhs364 _dafny.Set = (func() _dafny.Set {
				if (func() bool {
					if (_2186_v47).Contains((_this).F6()) {
						return (_2186_v47).Get((_this).F6()).(bool)
					}
					return p2
				})() {
					return _dafny.SetOf((_this).F6(), (_this).F6())
				}
				return _dafny.SetOf((_this).F6())
			})()
			_ = _rhs364
			var _rhs365 bool = ((_this).F6()).Cmp((_this).F6()) >= 0
			_ = _rhs365
			_2105_v2 = _rhs364
			r0 = _rhs365
			(globalState).F0 = p2
		}
		var _2187_v48 _dafny.Map
		_ = _2187_v48
		_2187_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("icjgt")).Cardinality()))
		var _2188_v49 D14
		_ = _2188_v49
		_2188_v49 = Companion_D14_.Create_DC36_(_2187_v48)
		var _2189_v50 _dafny.Array
		_ = _2189_v50
		var _nwElement0_64 D14 = (func() D14 {
			if !(p2) {
				return _2188_v49
			}
			return _2188_v49
		})()
		_ = _nwElement0_64
		var _nw323 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(2))
		_ = _nw323
		(_nw323).ArraySet1(_nwElement0_64, 0)
		(_nw323).ArraySet1(_2188_v49, 1)
		_2189_v50 = _nw323
		var _2190_v51 _dafny.Sequence
		_ = _2190_v51
		_2190_v51 = _dafny.SeqOf(Companion_Default___.Fm41(globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0))
		var _2191_v52 _dafny.Map
		_ = _2191_v52
		_2191_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
		var _2192_v53 _dafny.MultiSet
		_ = _2192_v53
		_2192_v53 = _dafny.MultiSetOf(_2191_v52, _2191_v52)
		var _rhs366 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _rhs366
		var _rhs367 bool = (_2192_v53).IsSubsetOf((_dafny.MultiSetFromSeq(_2190_v51)).Difference(_2192_v53))
		_ = _rhs367
		var _rhs368 _dafny.Array = _2189_v50
		_ = _rhs368
		var _lhs295 *GlobalState = globalState
		_ = _lhs295
		var _lhs296 *GlobalState = globalState
		_ = _lhs296
		_lhs295.F5 = _rhs366
		_lhs296.F0 = _rhs367
		_2189_v50 = _rhs368
		var _2193_v54 _dafny.MultiSet
		_ = _2193_v54
		_2193_v54 = _dafny.MultiSetOf(p0, Companion_Default___.Fm3(globalState))
		var _2194_v55 _dafny.Map
		_ = _2194_v55
		_2194_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2193_v54, _dafny.CodePoint('a'))
		var _2195_v56 _dafny.Sequence
		_ = _2195_v56
		_2195_v56 = _dafny.SeqOf(_2194_v55)
		if (Companion_Default___.SafeModuloInt(((_2195_v56).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2195_v56).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), (_this).F6())).Cmp((_this).F6()) == 0 {
			var _2196_v57 *C6
			_ = _2196_v57
			var _nw324 *C6 = New_C6_()
			_ = _nw324
			_nw324.Ctor__((_this).F6(), true)
			_2196_v57 = _nw324
			var _2197_v58 D6
			_ = _2197_v58
			_2197_v58 = Companion_D6_.Create_DC19_((_this).F6(), (_this).F6(), (_2105_v2).Cardinality())
			var _pat_let_tv56 = _2196_v57
			_ = _pat_let_tv56
			_2197_v58 = func(_pat_let72_0 D6) D6 {
				return func(_2198_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let73_0 _dafny.Int) D6 {
						return func(_2199_dt__update_hcf41_h0 _dafny.Int) D6 {
							return Companion_D6_.Create_DC19_((_2198_dt__update__tmp_h0).Dtor_cf39(), (_2198_dt__update__tmp_h0).Dtor_cf40(), _2199_dt__update_hcf41_h0)
						}(_pat_let73_0)
					}((_pat_let_tv56).F13())
				}(_pat_let72_0)
			}(Companion_D6_.Create_DC19_(_dafny.IntOfUint32((_2104_v1).Cardinality()), (_2196_v57).F13(), (_2196_v57).F13()))
			var _2200_v59 _dafny.Map
			_ = _2200_v59
			_2200_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_2196_v57).F14()) && (p2), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(p2)), _dafny.IntOfInt64(-942))).Update(p2, (_this).F6()))
			_2200_v59 = (_2200_v59).Update(!((_2196_v57).F14()), ((_2187_v48).Update(true, (_2196_v57).F13())).Update(false, (_this).F6()))
			var _2201_v60 _dafny.Array
			_ = _2201_v60
			var _nw325 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
			_ = _nw325
			_2201_v60 = _nw325
			_2201_v60 = _2201_v60
			var _2202_v61 _dafny.Array
			_ = _2202_v61
			var _len0_56 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_56
			var _nw326 _dafny.Array
			_ = _nw326
			if _len0_56.Cmp(_dafny.Zero) == 0 {
				_nw326 = _dafny.NewArray(_len0_56)
			} else {
				var _init56 func(_dafny.Int) D0 = (func(_2203_v6 D0) func(_dafny.Int) D0 {
					return func(_2204_i6 _dafny.Int) D0 {
						return _2203_v6
					}
				})(_2109_v6)
				_ = _init56
				var _element0_56 = _init56(_dafny.Zero)
				_ = _element0_56
				_nw326 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
				(_nw326).ArraySet1(_element0_56, 0)
				var _nativeLen0_56 = (_len0_56).Int()
				_ = _nativeLen0_56
				for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
					(_nw326).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
				}
			}
			_2202_v61 = _nw326
			var _2205_v62 D7
			_ = _2205_v62
			_2205_v62 = Companion_D7_.Create_DC21_(_2202_v61)
			var _source28 D7 = _2205_v62
			_ = _source28
			if _source28.Is_DC22() {
				var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((p1), 0))
				_ = _index303
				(p1).ArraySet1(((_2196_v57).F13()).Minus((_this).F6()), (_index303).Int())
				var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((p1), 0))
				_ = _index304
				(p1).ArraySet1((_this).F6(), (_index304).Int())
				var _rhs369 bool = p2
				_ = _rhs369
				var _rhs370 _dafny.Int = (_dafny.Zero).Minus((_2107_v4).Select((Companion_Default___.SafeIndex((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2107_v4).Cardinality()))).Uint32()).(_dafny.Int))
				_ = _rhs370
				var _rhs371 _dafny.Sequence = _2104_v1
				_ = _rhs371
				var _rhs372 _dafny.Int = (_this).F6()
				_ = _rhs372
				var _lhs297 *GlobalState = globalState
				_ = _lhs297
				var _lhs298 *GlobalState = globalState
				_ = _lhs298
				var _lhs299 *GlobalState = globalState
				_ = _lhs299
				var _lhs300 *GlobalState = globalState
				_ = _lhs300
				_lhs297.F0 = _rhs369
				_lhs298.F5 = _rhs370
				_lhs299.F2 = _rhs371
				_lhs300.F5 = _rhs372
				var _2206_v63 *C1
				_ = _2206_v63
				var _nw327 *C1 = New_C1_()
				_ = _nw327
				_nw327.Ctor__((_this).F6())
				_2206_v63 = _nw327
				(globalState).F0 = p0
			} else {
				var _2207___mcc_h12 _dafny.Array = _source28.Get_().(D7_DC21).Cf43
				_ = _2207___mcc_h12
				var _2208_cf43 _dafny.Array = _2207___mcc_h12
				_ = _2208_cf43
				(globalState).F0 = (_2196_v57).F14()
				var _2209_v64 *C1
				_ = _2209_v64
				var _nw328 *C1 = New_C1_()
				_ = _nw328
				_nw328.Ctor__(_dafny.IntOfInt64(469))
				_2209_v64 = _nw328
				var _2210_v65 _dafny.Map
				_ = _2210_v65
				_2210_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2187_v48, _2196_v57)
				var _2211_v66 _dafny.Map
				_ = _2211_v66
				_2211_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _2196_v57)
				var _2212_v67 _dafny.Array
				_ = _2212_v67
				var _nwElement0_65 *C6 = (func() *C6 {
					if (_2210_v65).Contains(_2187_v48) {
						return (_2210_v65).Get(_2187_v48).(*C6)
					}
					return _2196_v57
				})()
				_ = _nwElement0_65
				var _nw329 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(14))
				_ = _nw329
				(_nw329).ArraySet1(_nwElement0_65, 0)
				(_nw329).ArraySet1(_2196_v57, 1)
				(_nw329).ArraySet1(_2196_v57, 2)
				(_nw329).ArraySet1(_2196_v57, 3)
				(_nw329).ArraySet1(_2196_v57, 4)
				(_nw329).ArraySet1(_2196_v57, 5)
				(_nw329).ArraySet1(_2196_v57, 6)
				(_nw329).ArraySet1(_2196_v57, 7)
				(_nw329).ArraySet1(_2196_v57, 8)
				(_nw329).ArraySet1(_2196_v57, 9)
				(_nw329).ArraySet1(_2196_v57, 10)
				(_nw329).ArraySet1(_2196_v57, 11)
				(_nw329).ArraySet1((func() *C6 {
					if (_2211_v66).Contains((_2209_v64).Fm28(globalState)) {
						return (_2211_v66).Get((_2209_v64).Fm28(globalState)).(*C6)
					}
					return _2196_v57
				})(), 12)
				(_nw329).ArraySet1(_2196_v57, 13)
				_2212_v67 = _nw329
				_2212_v67 = _2212_v67
				_2191_v52 = (_2191_v52).Update((_2196_v57).F14(), ((_2196_v57).F13()).Cmp((_2196_v57).F13()) >= 0)
			}
		} else {
			var _2213_v68 _dafny.Map
			_ = _2213_v68
			_2213_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-195), _2191_v52)
			var _2214_v69 _dafny.MultiSet
			_ = _2214_v69
			_2214_v69 = _dafny.MultiSetOf(p1, p1)
			var _2215_v71 _dafny.Sequence
			_ = _2215_v71
			_2215_v71 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _2191_v52), _2213_v68)
			var _2216_v72 _dafny.Array
			_ = _2216_v72
			var _nwElement0_66 _dafny.Map = _2213_v68
			_ = _nwElement0_66
			var _nw330 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(29))
			_ = _nw330
			(_nw330).ArraySet1(_nwElement0_66, 0)
			(_nw330).ArraySet1(_2213_v68, 1)
			(_nw330).ArraySet1((_2213_v68).Merge(_2213_v68), 2)
			(_nw330).ArraySet1(_2213_v68, 3)
			(_nw330).ArraySet1(_2213_v68, 4)
			(_nw330).ArraySet1(_2213_v68, 5)
			(_nw330).ArraySet1((_2213_v68).Update((_2214_v69).Cardinality(), _2191_v52), 6)
			(_nw330).ArraySet1(_2213_v68, 7)
			(_nw330).ArraySet1(_2213_v68, 8)
			(_nw330).ArraySet1((_2213_v68).Merge(_2213_v68), 9)
			(_nw330).ArraySet1(_2213_v68, 10)
			(_nw330).ArraySet1(_2213_v68, 11)
			(_nw330).ArraySet1(func() _dafny.Map {
				var _coll59 = _dafny.NewMapBuilder()
				_ = _coll59
				for _iter63 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(604), _dafny.IntOfInt64(-336))); ; {
					_compr_59, _ok63 := _iter63()
					if !_ok63 {
						break
					}
					var _2217_v70 _dafny.Int
					_2217_v70 = interface{}(_compr_59).(_dafny.Int)
					if ((_dafny.IntOfInt64(604)).Cmp(_2217_v70) <= 0) && ((_2217_v70).Cmp(_dafny.IntOfInt64(-336)) < 0) {
						_coll59.Add((_2217_v70).Plus((_this).F6()), _2191_v52)
					}
				}
				return _coll59.ToMap()
			}(), 12)
			(_nw330).ArraySet1(_2213_v68, 13)
			(_nw330).ArraySet1((_2213_v68).Update(_dafny.IntOfUint32((_2108_v5).Cardinality()), _2191_v52), 14)
			(_nw330).ArraySet1(_2213_v68, 15)
			(_nw330).ArraySet1(_2213_v68, 16)
			(_nw330).ArraySet1((_2213_v68).Merge(_2213_v68), 17)
			(_nw330).ArraySet1((_2213_v68).Merge((_2215_v71).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_2103_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_2103_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_2215_v71).Cardinality()))).Uint32()).(_dafny.Map)), 18)
			(_nw330).ArraySet1((_2213_v68).Merge(_2213_v68), 19)
			(_nw330).ArraySet1(_2213_v68, 20)
			(_nw330).ArraySet1(_2213_v68, 21)
			(_nw330).ArraySet1(_2213_v68, 22)
			(_nw330).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(854), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)), 23)
			(_nw330).ArraySet1(_2213_v68, 24)
			(_nw330).ArraySet1(_2213_v68, 25)
			(_nw330).ArraySet1(_2213_v68, 26)
			(_nw330).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_2190_v51).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2104_v1).Cardinality()), _dafny.IntOfUint32((_2190_v51).Cardinality()))).Uint32()).(_dafny.Map)), 27)
			(_nw330).ArraySet1(_2213_v68, 28)
			_2216_v72 = _nw330
			var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2216_v72), 0))
			_ = _index305
			(_2216_v72).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)), (_index305).Int())
			var _2218_v73 D4
			_ = _2218_v73
			_2218_v73 = Companion_D4_.Create_DC14_(false, _2191_v52, (_this).F6(), _2191_v52)
			var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2216_v72), 0))
			_ = _index306
			(_2216_v72).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_2218_v73).Dtor_cf29())).Merge(_2213_v68), (_index306).Int())
			(globalState).F5 = (_this).F6()
			if (p0) || (p0) {
				var _2219_v74 _dafny.Sequence
				_ = _2219_v74
				_2219_v74 = _dafny.SeqOf(_2108_v5, _2108_v5, _dafny.SeqOf(p2), _2108_v5)
				_2219_v74 = _2219_v74
				(globalState).F0 = p0
				var _2220_v75 *C5
				_ = _2220_v75
				var _nw331 *C5 = New_C5_()
				_ = _nw331
				_nw331.Ctor__((_this).F6(), (_this).F6())
				_2220_v75 = _nw331
				var _2221_v76 _dafny.Map
				_ = _2221_v76
				_2221_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm26((_this).F6(), _2156_v32, (_2108_v5).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2108_v5).Cardinality()))).Uint32()).(bool), _2156_v32, globalState), (_this).F6()), _2104_v1)
				_2221_v76 = (_2221_v76).Update(p0, _2104_v1)
				var _2222_v77 _dafny.Sequence
				_ = _2222_v77
				_2222_v77 = _dafny.SeqOf(_2193_v54, _2193_v54, _2193_v54, _dafny.MultiSetFromSeq(_2108_v5), (_2193_v54).Update(p2, Companion_Default___.Abs((_dafny.Zero).Minus((_this).F6()))))
				var _2223_v78 _dafny.Set
				_ = _2223_v78
				_2223_v78 = _dafny.SetOf(_2107_v4, _2107_v4)
				var _2224_v79 _dafny.Map
				_ = _2224_v79
				_2224_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2223_v78, _2222_v77)
				_2222_v77 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.MultiSetOf(p0)).Update(p0, Companion_Default___.Abs(_dafny.IntOfInt64(683)))), _2222_v77), (func() _dafny.Sequence {
					if (_2224_v79).Contains(_2223_v78) {
						return (_2224_v79).Get(_2223_v78).(_dafny.Sequence)
					}
					return _2222_v77
				})())
			} else {
				var _2225_v80 _dafny.Array
				_ = _2225_v80
				var _nw332 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
				_ = _nw332
				_2225_v80 = _nw332
				var _2226_v81 _dafny.Map
				_ = _2226_v81
				_2226_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2156_v32, _2225_v80)
				_2226_v81 = (_2226_v81).Update(Companion_Default___.Fm0((_this).F6(), _2105_v2, p2, globalState), _2225_v80)
				var _2227_v82 _dafny.Map
				_ = _2227_v82
				_2227_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p2)
				var _2228_v83 _dafny.MultiSet
				_ = _2228_v83
				_2228_v83 = _dafny.MultiSetOf((_dafny.MultiSetOf(p2, true, (func() bool {
					if (_2227_v82).Contains((_this).F6()) {
						return (_2227_v82).Get((_this).F6()).(bool)
					}
					return !(p0)
				})())).Cardinality(), (_dafny.IntOfInt64(843)).Minus(_dafny.IntOfUint32((_2104_v1).Cardinality())), ((_this).F6()).Times((_this).F6()), (_this).F6(), ((_this).F6()).Plus(_dafny.IntOfInt64(758)))
				var _2229_v84 _dafny.Sequence
				_ = _2229_v84
				_2229_v84 = _dafny.SeqOf(_2104_v1, _2104_v1)
				var _rhs373 bool = p2
				_ = _rhs373
				var _rhs374 bool = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-848), (_this).F6())).Cmp((_this).F6()) >= 0
				_ = _rhs374
				var _rhs375 _dafny.Int = ((_this).F6()).Plus((_this).F6())
				_ = _rhs375
				var _rhs376 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_2229_v84, _2229_v84), _2229_v84)
				_ = _rhs376
				var _rhs377 _dafny.MultiSet = _2228_v83
				_ = _rhs377
				var _lhs301 *GlobalState = globalState
				_ = _lhs301
				var _lhs302 *GlobalState = globalState
				_ = _lhs302
				var _lhs303 *GlobalState = globalState
				_ = _lhs303
				var _lhs304 *GlobalState = globalState
				_ = _lhs304
				_lhs301.F0 = _rhs373
				_lhs302.F0 = _rhs374
				_lhs303.F5 = _rhs375
				_lhs304.F0 = _rhs376
				_2228_v83 = _rhs377
				_2227_v82 = (_2227_v82).Update((_this).F6(), ((_this).F6()).Cmp((_this).F6()) < 0)
				(globalState).F0 = p2
				var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((p1), 0))
				_ = _index307
				(p1).ArraySet1((_this).F6(), (_index307).Int())
				var _2230_v85 _dafny.Sequence
				_ = _2230_v85
				_2230_v85 = _dafny.SeqOf(_2157_v33)
				var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((p1), 0))
				_ = _index308
				(p1).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2230_v85, _2230_v85)).Cardinality()), (_index308).Int())
			}
			var _2231_v86 _dafny.Sequence
			_ = _2231_v86
			_2231_v86 = _dafny.SeqOf(_dafny.SeqOf(p2))
			var _2232_v87 _dafny.Sequence
			_ = _2232_v87
			_2232_v87 = _dafny.SeqOf(_2231_v86, _2231_v86, _2231_v86, _2231_v86)
			var _2233_v88 _dafny.Map
			_ = _2233_v88
			_2233_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2232_v87).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2232_v87).Cardinality()))).Uint32()).(_dafny.Sequence), p0)
			_2233_v88 = (_2233_v88).Update(_dafny.SeqOf(_2108_v5), Companion_Default___.Fm3(globalState))
			var _2234_v89 _dafny.Array
			_ = _2234_v89
			var _nw333 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
			_ = _nw333
			_2234_v89 = _nw333
			var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_2234_v89), 0))
			_ = _index309
			(_2234_v89).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2104_v1, _2104_v1), (_index309).Int())
			var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_2234_v89), 0))
			_ = _index310
			(_2234_v89).ArraySet1(_2104_v1, (_index310).Int())
		}
		r0 = p0
		return r0
	}
}

// End of class C9

// Definition of class C10
type C10 struct {
	_f6 _dafny.Int
}

func New_C10_() *C10 {
	_this := C10{}

	_this._f6 = _dafny.Zero
	return &_this
}

type CompanionStruct_C10_ struct {
}

var Companion_C10_ = CompanionStruct_C10_{}

func (_this *C10) Equals(other *C10) bool {
	return _this == other
}

func (_this *C10) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C10)
	return ok && _this.Equals(other)
}

func (*C10) String() string {
	return "_module.C10"
}

func Type_C10_() _dafny.TypeDescriptor {
	return type_C10_{}
}

type type_C10_ struct {
}

func (_this type_C10_) Default() interface{} {
	return (*C10)(nil)
}

func (_this type_C10_) String() string {
	return "main.C10"
}
func (_this *C10) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C10{}
var _ _dafny.TraitOffspring = &C10{}

func (_this *C10) F6() _dafny.Int {
	return _this._f6
}
func (_this *C10) Ctor__(f6 _dafny.Int) {
	{
		(_this)._f6 = f6
	}
}
func (_this *C10) Fm5(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('o')))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(771))).Uint32(), func(coer126 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg126 _dafny.Int) interface{} {
				return coer126(arg126)
			}
		}(func(_2235_i0 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetFromSeq((Companion_D1_.Create_DC4_(Companion_D0_.Create_DC0_(_dafny.SeqOf(true, !(true)), _dafny.IntOfInt64(796)), (_this).F6(), _dafny.UnicodeSeqOfUtf8Bytes("lxfutcy"), false, (_this).F6())).Dtor_cf9())
		}))), (func() _dafny.Sequence {
			if false {
				return _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('y'), _dafny.CodePoint('l')))
			}
			return _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('c')), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('p'))))
		})())
	}
}
func (_this *C10) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.SeqOf(_dafny.IntOfInt64(471), _dafny.IntOfUint32((_dafny.SeqOf(false, false, false, false, false)).Cardinality())))).Merge(func() _dafny.Map {
			var _coll60 = _dafny.NewMapBuilder()
			_ = _coll60
			for _iter64 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(858), _dafny.IntOfInt64(166))); ; {
				_compr_60, _ok64 := _iter64()
				if !_ok64 {
					break
				}
				var _2236_v0 _dafny.Int
				_2236_v0 = interface{}(_compr_60).(_dafny.Int)
				if ((_dafny.IntOfInt64(858)).Cmp(_2236_v0) <= 0) && ((_2236_v0).Cmp(_dafny.IntOfInt64(166)) < 0) {
					_coll60.Add(Companion_Default___.SafeModuloInt(_2236_v0, (_this).F6()), _dafny.SeqOf((_this).F6(), (_this).F6()))
				}
			}
			return _coll60.ToMap()
		}())
	}
}
func (_this *C10) M1(p0 _dafny.Sequence, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _2237_v0 _dafny.Sequence
		_ = _2237_v0
		_2237_v0 = _dafny.SeqOf(_dafny.IntOfUint32((p0).Cardinality()), (_this).F6())
		var _2238_v1 bool
		_ = _2238_v1
		_2238_v1 = false
		var _2239_v2 _dafny.Sequence
		_ = _2239_v2
		_2239_v2 = _dafny.SeqOf(_2238_v1, _2238_v1)
		var _2240_v3 D0
		_ = _2240_v3
		_2240_v3 = Companion_D0_.Create_DC0_(_2239_v2, (_this).F6())
		var _2241_v4 _dafny.Sequence
		_ = _2241_v4
		_2241_v4 = _dafny.SeqOf(((_this).F6()).Times((_2237_v0).Select((Companion_Default___.SafeIndex((_2240_v3).Dtor_cf1(), _dafny.IntOfUint32((_2237_v0).Cardinality()))).Uint32()).(_dafny.Int)))
		(globalState).F5 = (_2237_v0).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2237_v0).Cardinality()))).Uint32()).(_dafny.Int)
		var _2242_v5 _dafny.Array
		_ = _2242_v5
		var _len0_57 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_57
		var _nw334 _dafny.Array
		_ = _nw334
		if _len0_57.Cmp(_dafny.Zero) == 0 {
			_nw334 = _dafny.NewArray(_len0_57)
		} else {
			var _init57 func(_dafny.Int) bool = (func(_2243_v1 bool) func(_dafny.Int) bool {
				return func(_2244_i0 _dafny.Int) bool {
					return _2243_v1
				}
			})(_2238_v1)
			_ = _init57
			var _element0_57 = _init57(_dafny.Zero)
			_ = _element0_57
			_nw334 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
			(_nw334).ArraySet1(_element0_57, 0)
			var _nativeLen0_57 = (_len0_57).Int()
			_ = _nativeLen0_57
			for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
				(_nw334).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
			}
		}
		_2242_v5 = _nw334
		_2242_v5 = _2242_v5
		var _2245_i1 _dafny.Int
		_ = _2245_i1
		_2245_i1 = _dafny.Zero
		{
			for _2238_v1 {
				{
					if (_2245_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_2245_i1 = (_2245_i1).Plus(_dafny.One)
					var _2246_v6 _dafny.Map
					_ = _2246_v6
					_2246_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(7), _2238_v1)
					var _2247_v7 *C6
					_ = _2247_v7
					var _nw335 *C6 = New_C6_()
					_ = _nw335
					_nw335.Ctor__((_2246_v6).Cardinality(), _2238_v1)
					_2247_v7 = _nw335
					var _2248_v8 _dafny.Map
					_ = _2248_v8
					_2248_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _2247_v7)
					var _2249_v9 D18
					_ = _2249_v9
					_2249_v9 = Companion_D18_.Create_DC42_(_2248_v8)
					var _2250_v10 _dafny.MultiSet
					_ = _2250_v10
					_2250_v10 = _dafny.MultiSetOf(((_2249_v9).Dtor_cf77()).Cardinality())
					var _2251_v11 _dafny.Map
					_ = _2251_v11
					_2251_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_dafny.Zero).Minus((func() _dafny.Int {
						if (_2250_v10).Contains((_this).F6()) {
							return (_2250_v10).Multiplicity((_this).F6())
						}
						return _dafny.IntOfInt64(-417)
					})()))
					var _2252_v12 _dafny.Map
					_ = _2252_v12
					_2252_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p0)
					var _2253_v13 *C3
					_ = _2253_v13
					var _nw336 *C3 = New_C3_()
					_ = _nw336
					_nw336.Ctor__(_2251_v11, _2242_v5, (_2252_v12).Merge(_2252_v12), (_2247_v7).F13())
					_2253_v13 = _nw336
					var _2254_v14 _dafny.Map
					_ = _2254_v14
					_2254_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2247_v7).F14(), (_this).F6())
					var _2255_v15 _dafny.Sequence
					_ = _2255_v15
					_2255_v15 = _dafny.SeqOf(_2241_v4)
					var _2256_v16 _dafny.Set
					_ = _2256_v16
					_2256_v16 = _dafny.SetOf((_2247_v7).F14())
					var _2257_v17 _dafny.Array
					_ = _2257_v17
					var _nwElement0_67 _dafny.Int = _dafny.IntOfInt64(-254)
					_ = _nwElement0_67
					var _nw337 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(28))
					_ = _nw337
					(_nw337).ArraySet1(_nwElement0_67, 0)
					(_nw337).ArraySet1((_this).F6(), 1)
					(_nw337).ArraySet1((func() _dafny.Int {
						if (_2250_v10).Contains((_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F6(), _dafny.IntOfInt64(953), _dafny.IntOfInt64(72)))).Cardinality()) {
							return (_2250_v10).Multiplicity((_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F6(), _dafny.IntOfInt64(953), _dafny.IntOfInt64(72)))).Cardinality())
						}
						return (_this).F6()
					})(), 2)
					(_nw337).ArraySet1((_this).F6(), 3)
					(_nw337).ArraySet1((_2247_v7).F13(), 4)
					(_nw337).ArraySet1((_2247_v7).F13(), 5)
					(_nw337).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(((func() _dafny.Int {
						if (_2250_v10).Contains((_2241_v4).Select((Companion_Default___.SafeIndex((_2247_v7).F13(), _dafny.IntOfUint32((_2241_v4).Cardinality()))).Uint32()).(_dafny.Int)) {
							return (_2250_v10).Multiplicity((_2241_v4).Select((Companion_Default___.SafeIndex((_2247_v7).F13(), _dafny.IntOfUint32((_2241_v4).Cardinality()))).Uint32()).(_dafny.Int))
						}
						return (_dafny.Zero).Minus(_dafny.IntOfUint32((_2239_v2).Cardinality()))
					})()).Plus((_2254_v14).Cardinality()))), 6)
					(_nw337).ArraySet1((_this).F6(), 7)
					(_nw337).ArraySet1((_this).F6(), 8)
					(_nw337).ArraySet1((func() _dafny.Int {
						if (_2247_v7).F14() {
							return (_this).F6()
						}
						return (_dafny.Zero).Minus((_this).F6())
					})(), 9)
					(_nw337).ArraySet1((_this).F6(), 10)
					(_nw337).ArraySet1((_2247_v7).F13(), 11)
					(_nw337).ArraySet1((_2247_v7).F13(), 12)
					(_nw337).ArraySet1((_2247_v7).F13(), 13)
					(_nw337).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2237_v0, (_2255_v15).Select((Companion_Default___.SafeIndex((_2256_v16).Cardinality(), _dafny.IntOfUint32((_2255_v15).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()), 14)
					(_nw337).ArraySet1(_dafny.IntOfInt64(175), 15)
					(_nw337).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(509), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(352))).Uint32(), func(coer127 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg127 _dafny.Int) interface{} {
							return coer127(arg127)
						}
					}(func(_2258_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('w')
					}))).Cardinality())), 16)
					(_nw337).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F6(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yl")).Cardinality())), 17)
					(_nw337).ArraySet1((_this).F6(), 18)
					(_nw337).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F6())), 19)
					(_nw337).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 20)
					(_nw337).ArraySet1(((_2247_v7).F13()).Plus((Companion_Default___.Fm19((_2247_v7).F13(), globalState)).Cardinality()), 21)
					(_nw337).ArraySet1((_2247_v7).F13(), 22)
					(_nw337).ArraySet1(Companion_Default___.SafeDivisionInt((_2247_v7).F13(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_2247_v7).F13())).Cardinality()))), 23)
					(_nw337).ArraySet1((_2247_v7).F13(), 24)
					(_nw337).ArraySet1(_dafny.IntOfInt64(118), 25)
					(_nw337).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((p0).Cardinality())), _dafny.IntOfUint32(((Companion_D5_.Create_DC17_(_2241_v4, _2238_v1)).Dtor_cf36()).Cardinality()))), 26)
					(_nw337).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F6()), (_2247_v7).F13()), 27)
					_2257_v17 = _nw337
					var _2259_v18 D2
					_ = _2259_v18
					_2259_v18 = Companion_D2_.Create_DC8_(_2238_v1, false)
					var _2260_v19 _dafny.CodePoint
					_ = _2260_v19
					_2260_v19 = _dafny.CodePoint('f')
					var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_2257_v17), 0))
					_ = _index311
					(_2257_v17).ArraySet1((_2247_v7).Fm16(_2259_v18, Companion_D0_.Create_DC1_(_2238_v1, (_this).F6(), _2260_v19), globalState), (_index311).Int())
					var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_2257_v17), 0))
					_ = _index312
					(_2257_v17).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_this).F6(), _dafny.IntOfInt64(-30)), (_2247_v7).F13()), (_index312).Int())
					var _2261_v20 _dafny.Map
					_ = _2261_v20
					_2261_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(!(_2238_v1)))
					var _2262_v21 _dafny.Sequence
					_ = _2262_v21
					_2262_v21 = _dafny.SeqOf(p0)
					var _2263_v22 *C8
					_ = _2263_v22
					var _nw338 *C8 = New_C8_()
					_ = _nw338
					_nw338.Ctor__()
					_2263_v22 = _nw338
					var _2264_v23 _dafny.Sequence
					_ = _2264_v23
					_2264_v23 = _dafny.SeqOf(_2263_v22)
					_2261_v20 = (_2261_v20).Update(_dafny.Companion_Sequence_.Concatenate((_2262_v21).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2264_v23).Cardinality()), _dafny.IntOfUint32((_2262_v21).Cardinality()))).Uint32()).(_dafny.Sequence), p0), (_2247_v7).F14())
					var _2265_v24 D9
					_ = _2265_v24
					_2265_v24 = Companion_D9_.Create_DC25_(_dafny.IntOfInt64(765), (func() _dafny.Int {
						if (_2247_v7).F14() {
							return (_2247_v7).F13()
						}
						return (_this).F6()
					})(), (_2257_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_2257_v17), 0))).Int()).(_dafny.Int), _2260_v19)
					var _2266_v25 _dafny.MultiSet
					_ = _2266_v25
					_2266_v25 = _dafny.MultiSetOf(_2250_v10)
					var _2267_v26 _dafny.Sequence
					_ = _2267_v26
					_2267_v26 = _dafny.SeqOf(Companion_Default___.Fm4((_this).F6(), (_2257_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_2257_v17), 0))).Int()).(_dafny.Int), (_2247_v7).F13(), globalState))
					var _2268_v27 _dafny.Sequence
					_ = _2268_v27
					_2268_v27 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_2237_v0, (Companion_Default___.SafeIndex((_2257_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_2257_v17), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2237_v0).Cardinality()))).Uint32(), (_2257_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_2257_v17), 0))).Int()).(_dafny.Int))), _2250_v10)
					var _rhs378 bool = ((_2239_v2).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2239_v2).Cardinality()))).Uint32()).(bool)) && ((_2239_v2).Select((Companion_Default___.SafeIndex((_2265_v24).Dtor_cf48(), _dafny.IntOfUint32((_2239_v2).Cardinality()))).Uint32()).(bool))
					_ = _rhs378
					var _rhs379 _dafny.Array = _2242_v5
					_ = _rhs379
					var _rhs380 D9 = _2265_v24
					_ = _rhs380
					var _rhs381 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_2239_v2), _2267_v26)
					_ = _rhs381
					var _rhs382 _dafny.MultiSet = ((_2266_v25).Difference(_dafny.MultiSetFromSeq(_2268_v27))).Union(_2266_v25)
					_ = _rhs382
					var _lhs305 *GlobalState = globalState
					_ = _lhs305
					var _lhs306 *GlobalState = globalState
					_ = _lhs306
					_lhs305.F0 = _rhs378
					_2242_v5 = _rhs379
					_2265_v24 = _rhs380
					_lhs306.F0 = _rhs381
					_2266_v25 = _rhs382
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _2269_v28 _dafny.CodePoint
		_ = _2269_v28
		_2269_v28 = _dafny.CodePoint('q')
		_2269_v28 = _2269_v28
		(globalState).F2 = _dafny.UnicodeSeqOfUtf8Bytes("henupki")
		var _2270_v29 _dafny.Map
		_ = _2270_v29
		_2270_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
		var _2271_v30 D6
		_ = _2271_v30
		_2271_v30 = Companion_D6_.Create_DC19_((_2270_v29).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F6(), (_this).F6(), (_this).F6())).Cardinality()), (_this).F6())
		var _2272_v31 D6
		_ = _2272_v31
		_2272_v31 = Companion_D6_.Create_DC20_(_2271_v30)
		var _pat_let_tv57 = _2271_v30
		_ = _pat_let_tv57
		var _2273_v32 _dafny.Array
		_ = _2273_v32
		var _nwElement0_68 D6 = Companion_D6_.Create_DC20_(_2271_v30)
		_ = _nwElement0_68
		var _nw339 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(29))
		_ = _nw339
		(_nw339).ArraySet1(_nwElement0_68, 0)
		(_nw339).ArraySet1(Companion_D6_.Create_DC20_(_2271_v30), 1)
		(_nw339).ArraySet1(_2272_v31, 2)
		(_nw339).ArraySet1(_2272_v31, 3)
		(_nw339).ArraySet1(func(_pat_let74_0 D6) D6 {
			return func(_2274_dt__update__tmp_h0 D6) D6 {
				return func(_pat_let75_0 D6) D6 {
					return func(_2275_dt__update_hcf42_h0 D6) D6 {
						return Companion_D6_.Create_DC20_(_2275_dt__update_hcf42_h0)
					}(_pat_let75_0)
				}(_pat_let_tv57)
			}(_pat_let74_0)
		}(_2272_v31), 4)
		(_nw339).ArraySet1(_2272_v31, 5)
		(_nw339).ArraySet1(_2272_v31, 6)
		(_nw339).ArraySet1(Companion_D6_.Create_DC20_(Companion_D6_.Create_DC20_(_2271_v30)), 7)
		(_nw339).ArraySet1(Companion_D6_.Create_DC20_(_2271_v30), 8)
		(_nw339).ArraySet1(_2272_v31, 9)
		(_nw339).ArraySet1(_2272_v31, 10)
		(_nw339).ArraySet1(_2272_v31, 11)
		(_nw339).ArraySet1(_2272_v31, 12)
		(_nw339).ArraySet1(_2272_v31, 13)
		(_nw339).ArraySet1(_2272_v31, 14)
		(_nw339).ArraySet1(_2272_v31, 15)
		(_nw339).ArraySet1(_2272_v31, 16)
		(_nw339).ArraySet1(_2272_v31, 17)
		(_nw339).ArraySet1(_2272_v31, 18)
		(_nw339).ArraySet1(_2272_v31, 19)
		(_nw339).ArraySet1(_2272_v31, 20)
		(_nw339).ArraySet1(_2272_v31, 21)
		(_nw339).ArraySet1(_2272_v31, 22)
		(_nw339).ArraySet1(_2272_v31, 23)
		(_nw339).ArraySet1(_2272_v31, 24)
		(_nw339).ArraySet1(_2272_v31, 25)
		(_nw339).ArraySet1(_2272_v31, 26)
		(_nw339).ArraySet1(_2272_v31, 27)
		(_nw339).ArraySet1(_2272_v31, 28)
		_2273_v32 = _nw339
		var _2276_v33 _dafny.Array
		_ = _2276_v33
		_2276_v33 = _2273_v32
		var _source29 _dafny.Array = _2276_v33
		_ = _source29
		var _2277___mcc_h0 _dafny.Array = _source29
		_ = _2277___mcc_h0
		var _2278_cf72 _dafny.Array = _2277___mcc_h0
		_ = _2278_cf72
		var _2279_v34 D1
		_ = _2279_v34
		_2279_v34 = Companion_D1_.Create_DC4_(_2240_v3, (_this).F6(), p0, _2238_v1, Companion_Default___.Fm2((_this).F6(), globalState))
		var _source30 D1 = _2279_v34
		_ = _source30
		if _source30.Is_DC4() {
			var _2280___mcc_h1 D0 = _source30.Get_().(D1_DC4).Cf7
			_ = _2280___mcc_h1
			var _2281___mcc_h2 _dafny.Int = _source30.Get_().(D1_DC4).Cf8
			_ = _2281___mcc_h2
			var _2282___mcc_h3 _dafny.Sequence = _source30.Get_().(D1_DC4).Cf9
			_ = _2282___mcc_h3
			var _2283___mcc_h4 bool = _source30.Get_().(D1_DC4).Cf10
			_ = _2283___mcc_h4
			var _2284___mcc_h5 _dafny.Int = _source30.Get_().(D1_DC4).Cf11
			_ = _2284___mcc_h5
			var _2285_cf11 _dafny.Int = _2284___mcc_h5
			_ = _2285_cf11
			var _2286_cf10 bool = _2283___mcc_h4
			_ = _2286_cf10
			var _2287_cf9 _dafny.Sequence = _2282___mcc_h3
			_ = _2287_cf9
			var _2288_cf8 _dafny.Int = _2281___mcc_h2
			_ = _2288_cf8
			var _2289_cf7 D0 = _2280___mcc_h1
			_ = _2289_cf7
			(globalState).F5 = _2288_cf8
			var _2290_v35 _dafny.MultiSet
			_ = _2290_v35
			_2290_v35 = _dafny.MultiSetOf(_2285_cf11)
			var _2291_v36 _dafny.Sequence
			_ = _2291_v36
			_2291_v36 = _dafny.SeqOf(_2290_v35)
			var _2292_v37 _dafny.MultiSet
			_ = _2292_v37
			_2292_v37 = _dafny.MultiSetOf(_2238_v1)
			var _2293_v38 _dafny.Sequence
			_ = _2293_v38
			_2293_v38 = _dafny.SeqOf(_2290_v35, _2290_v35, _2290_v35, _2290_v35, (_2291_v36).Select((Companion_Default___.SafeIndex((_2292_v37).Cardinality(), _dafny.IntOfUint32((_2291_v36).Cardinality()))).Uint32()).(_dafny.MultiSet))
			var _2294_v39 _dafny.Array
			_ = _2294_v39
			var _nwElement0_69 _dafny.Int = Companion_Default___.Fm2(_2285_cf11, globalState)
			_ = _nwElement0_69
			var _nw340 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(15))
			_ = _nw340
			(_nw340).ArraySet1(_nwElement0_69, 0)
			(_nw340).ArraySet1(_dafny.IntOfInt64(-393), 1)
			(_nw340).ArraySet1(_2285_cf11, 2)
			(_nw340).ArraySet1(_dafny.IntOfInt64(-989), 3)
			(_nw340).ArraySet1(((_this).F6()).Plus(_dafny.IntOfInt64(999)), 4)
			(_nw340).ArraySet1(Companion_Default___.SafeModuloInt(_2288_cf8, _2285_cf11), 5)
			(_nw340).ArraySet1((_dafny.Zero).Minus(((func() _dafny.Int {
				if (_2270_v29).Contains(_dafny.IntOfUint32((_2287_cf9).Cardinality())) {
					return (_2270_v29).Get(_dafny.IntOfUint32((_2287_cf9).Cardinality())).(_dafny.Int)
				}
				return (_2290_v35).Cardinality()
			})()).Times(_dafny.IntOfUint32((_2291_v36).Cardinality()))), 6)
			(_nw340).ArraySet1((_dafny.IntOfUint32((p0).Cardinality())).Times(_2285_cf11), 7)
			(_nw340).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2288_cf8, (_2239_v2).Select((Companion_Default___.SafeIndex(_2285_cf11, _dafny.IntOfUint32((_2239_v2).Cardinality()))).Uint32()).(bool))).Cardinality(), 8)
			(_nw340).ArraySet1((_this).F6(), 9)
			(_nw340).ArraySet1((func() _dafny.Int {
				if (_2290_v35).Contains((_this).F6()) {
					return (_2290_v35).Multiplicity((_this).F6())
				}
				return _2285_cf11
			})(), 10)
			(_nw340).ArraySet1((_2288_cf8).Minus((_dafny.Zero).Minus(_2285_cf11)), 11)
			(_nw340).ArraySet1(_dafny.IntOfInt64(664), 12)
			(_nw340).ArraySet1((_2292_v37).Cardinality(), 13)
			(_nw340).ArraySet1(_2288_cf8, 14)
			_2294_v39 = _nw340
			var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_2294_v39), 0))
			_ = _index313
			(_2294_v39).ArraySet1(_2288_cf8, (_index313).Int())
			var _2295_v40 D2
			_ = _2295_v40
			_2295_v40 = Companion_D2_.Create_DC6_(_2285_cf11, _2294_v39)
			var _2296_v41 _dafny.MultiSet
			_ = _2296_v41
			_2296_v41 = _dafny.MultiSetOf((_2295_v40).Dtor_cf14())
			var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_2294_v39), 0))
			_ = _index314
			(_2294_v39).ArraySet1(((_2296_v41).Union(_dafny.MultiSetOf(_2294_v39, _2294_v39))).Cardinality(), (_index314).Int())
			(globalState).F0 = _2286_cf10
			var _2297_v42 _dafny.Map
			_ = _2297_v42
			_2297_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2285_cf11, p0)
			var _2298_v43 *C7
			_ = _2298_v43
			var _nw341 *C7 = New_C7_()
			_ = _nw341
			_nw341.Ctor__(_2294_v39, _2294_v39, Companion_Default___.Fm2((_this).F6(), globalState), _2242_v5, _2297_v42)
			_2298_v43 = _nw341
			var _rhs383 *C7 = _2298_v43
			_ = _rhs383
			var _rhs384 _dafny.Int = _2285_cf11
			_ = _rhs384
			_2298_v43 = _rhs383
			_2288_cf8 = _rhs384
		} else {
			var _2299___mcc_h6 _dafny.MultiSet = _source30.Get_().(D1_DC3).Cf6
			_ = _2299___mcc_h6
			var _2300_cf6 _dafny.MultiSet = _2299___mcc_h6
			_ = _2300_cf6
			(globalState).F0 = !(!_dafny.Companion_Sequence_.Contains(p0, _2269_v28))
			_2240_v3 = _2240_v3
			(globalState).F0 = ((_this).F6()).Cmp(_dafny.IntOfInt64(322)) > 0
			var _2301_v44 *C1
			_ = _2301_v44
			var _nw342 *C1 = New_C1_()
			_ = _nw342
			_nw342.Ctor__((_this).F6())
			_2301_v44 = _nw342
		}
		var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_2242_v5), 0))
		_ = _index315
		(_2242_v5).ArraySet1(_2238_v1, (_index315).Int())
		var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_2242_v5), 0))
		_ = _index316
		(_2242_v5).ArraySet1((_dafny.IntOfUint32((p0).Cardinality())).Cmp((_this).F6()) < 0, (_index316).Int())
		var _2302_v45 _dafny.Set
		_ = _2302_v45
		_2302_v45 = _dafny.SetOf((_this).F6())
		var _2303_v46 _dafny.MultiSet
		_ = _2303_v46
		_2303_v46 = _dafny.MultiSetOf((_2302_v45).Cardinality(), _dafny.IntOfInt64(951))
		var _2304_v47 _dafny.MultiSet
		_ = _2304_v47
		_2304_v47 = _dafny.MultiSetOf(_2303_v46)
		var _2305_v48 _dafny.Array
		_ = _2305_v48
		var _nw343 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
		_ = _nw343
		_2305_v48 = _nw343
		var _2306_v50 T2
		_ = _2306_v50
		var _nw344 *C2 = New_C2_()
		_ = _nw344
		_nw344.Ctor__((_this).F6(), _2304_v47, _2305_v48, func() _dafny.Map {
			var _coll61 = _dafny.NewMapBuilder()
			_ = _coll61
			for _iter65 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(993), _dafny.IntOfInt64(-151))); ; {
				_compr_61, _ok65 := _iter65()
				if !_ok65 {
					break
				}
				var _2307_v49 _dafny.Int
				_2307_v49 = interface{}(_compr_61).(_dafny.Int)
				if ((_dafny.IntOfInt64(993)).Cmp(_2307_v49) <= 0) && ((_2307_v49).Cmp(_dafny.IntOfInt64(-151)) < 0) {
					_coll61.Add((_2307_v49).Minus((_this).F6()), p0)
				}
			}
			return _coll61.ToMap()
		}())
		_2306_v50 = _nw344
		var _2308_v51 _dafny.MultiSet
		_ = _2308_v51
		_2308_v51 = _dafny.MultiSetOf(_2306_v50)
		var _2309_v52 _dafny.Map
		_ = _2309_v52
		_2309_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2306_v50).F6(), p0)
		var _2310_v53 _dafny.Map
		_ = _2310_v53
		_2310_v53 = _2309_v52
		var _2311_v54 _dafny.MultiSet
		_ = _2311_v54
		_2311_v54 = _dafny.MultiSetOf(_2310_v53)
		var _2312_v55 _dafny.MultiSet
		_ = _2312_v55
		_2312_v55 = _dafny.MultiSetOf(!(_2238_v1))
		var _2313_v56 _dafny.Array
		_ = _2313_v56
		var _nwElement0_70 _dafny.Int = (_this).F6()
		_ = _nwElement0_70
		var _nw345 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(17))
		_ = _nw345
		(_nw345).ArraySet1(_nwElement0_70, 0)
		(_nw345).ArraySet1(((_this).F6()).Plus((_dafny.Zero).Minus((_this).F6())), 1)
		(_nw345).ArraySet1((_this).F6(), 2)
		(_nw345).ArraySet1((_this).F6(), 3)
		(_nw345).ArraySet1(((_this).F6()).Plus((_this).F6()), 4)
		(_nw345).ArraySet1((_this).F6(), 5)
		(_nw345).ArraySet1((_this).F6(), 6)
		(_nw345).ArraySet1((_2308_v51).Cardinality(), 7)
		(_nw345).ArraySet1((_2306_v50).F6(), 8)
		(_nw345).ArraySet1(((_2306_v50).F6()).Minus(Companion_Default___.Fm2((_this).F6(), globalState)), 9)
		(_nw345).ArraySet1((_this).F6(), 10)
		(_nw345).ArraySet1(((_2306_v50).F6()).Minus((_2306_v50).F6()), 11)
		(_nw345).ArraySet1(((_2306_v50).F6()).Minus((_this).F6()), 12)
		(_nw345).ArraySet1((_this).F6(), 13)
		(_nw345).ArraySet1((_this).F6(), 14)
		(_nw345).ArraySet1((_2306_v50).F6(), 15)
		(_nw345).ArraySet1(((((_2311_v54).Update(_2310_v53, Companion_Default___.Abs((_2312_v55).Cardinality()))).Update(_2310_v53, Companion_Default___.Abs((_this).F6()))).Intersection(_2311_v54)).Cardinality(), 16)
		_2313_v56 = _nw345
		var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2313_v56), 0))
		_ = _index317
		(_2313_v56).ArraySet1((_2306_v50).F6(), (_index317).Int())
		var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2313_v56), 0))
		_ = _index318
		(_2313_v56).ArraySet1(Companion_Default___.Fm2((_2306_v50).F6(), globalState), (_index318).Int())
		if Companion_Default___.Fm3(globalState) {
			var _2314_v57 _dafny.Sequence
			_ = _2314_v57
			_2314_v57 = _dafny.SeqOf(_2306_v50)
			var _2315_v58 _dafny.Map
			_ = _2315_v58
			_2315_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_2239_v2, (Companion_Default___.SafeIndex((_2312_v55).Cardinality(), _dafny.IntOfUint32((_2239_v2).Cardinality()))).Uint32(), false), _dafny.IntOfUint32((_2314_v57).Cardinality()))
			_2315_v58 = (_2315_v58).Update(_2239_v2, ((_2313_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2313_v56), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfInt64(804)))
			var _2316_v59 bool
			_ = _2316_v59
			var _out31 bool
			_ = _out31
			_out31 = (_this).M2(globalState)
			_2316_v59 = _out31
			var _2317_v60 _dafny.MultiSet
			_ = _2317_v60
			_2317_v60 = _dafny.MultiSetOf(_dafny.SeqOf((_this).F6()), _2241_v4)
			var _2318_v61 _dafny.Map
			_ = _2318_v61
			_2318_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2313_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2313_v56), 0))).Int()).(_dafny.Int), _2317_v60)
			_2237_v0 = _dafny.SeqOf(((_2306_v50).F6()).Plus((func() _dafny.Set {
				var _coll62 = _dafny.NewBuilder()
				_ = _coll62
				for _iter66 := _dafny.Iterate(((func() _dafny.MultiSet {
					if (_2318_v61).Contains((_2306_v50).F6()) {
						return (_2318_v61).Get((_2306_v50).F6()).(_dafny.MultiSet)
					}
					return _2317_v60
				})()).Elements()); ; {
					_compr_62, _ok66 := _iter66()
					if !_ok66 {
						break
					}
					var _2319_v62 _dafny.Sequence
					_2319_v62 = interface{}(_compr_62).(_dafny.Sequence)
					if ((func() _dafny.MultiSet {
						if (_2318_v61).Contains((_2306_v50).F6()) {
							return (_2318_v61).Get((_2306_v50).F6()).(_dafny.MultiSet)
						}
						return _2317_v60
					})()).Contains(_2319_v62) {
						_coll62.Add(_2319_v62)
					}
				}
				return _coll62.ToSet()
			}()).Cardinality()))
			var _2320_v63 _dafny.Array
			_ = _2320_v63
			var _nw346 _dafny.Array = _dafny.NewArrayWithValue(Companion_D11_.Default(), _dafny.IntOfInt64(18))
			_ = _nw346
			_2320_v63 = _nw346
			var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_2242_v5), 0))
			_ = _index319
			var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2313_v56), 0))
			_ = _index320
			var _rhs385 _dafny.Array = _2320_v63
			_ = _rhs385
			var _rhs386 bool = !(Companion_Default___.Fm3(globalState))
			_ = _rhs386
			var _rhs387 bool = _2238_v1
			_ = _rhs387
			var _rhs388 _dafny.Int = Companion_Default___.SafeDivisionInt((_2306_v50).F6(), _dafny.IntOfUint32((p0).Cardinality()))
			_ = _rhs388
			var _rhs389 _dafny.Array = _2313_v56
			_ = _rhs389
			var _lhs307 _dafny.Array = _2242_v5
			_ = _lhs307
			var _lhs308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_2242_v5), 0))
			_ = _lhs308
			var _lhs309 _dafny.Array = _2313_v56
			_ = _lhs309
			var _lhs310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2313_v56), 0))
			_ = _lhs310
			_2320_v63 = _rhs385
			_2316_v59 = _rhs386
			(_lhs307).ArraySet1(_rhs387, (_lhs308).Int())
			(_lhs309).ArraySet1(_rhs388, (_lhs310).Int())
			_2313_v56 = _rhs389
			(globalState).F5 = (_dafny.Zero).Minus((((_2306_v50).F6()).Times((_2306_v50).F6())).Times((_2237_v0).Select((Companion_Default___.SafeIndex((_2313_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2313_v56), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2237_v0).Cardinality()))).Uint32()).(_dafny.Int)))
		} else {
			var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2313_v56), 0))
			_ = _index321
			(_2313_v56).ArraySet1((_this).F6(), (_index321).Int())
			var _2321_v64 *C9
			_ = _2321_v64
			var _nw347 *C9 = New_C9_()
			_ = _nw347
			_nw347.Ctor__(((_2306_v50).F6()).Plus((_2313_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2313_v56), 0))).Int()).(_dafny.Int)))
			_2321_v64 = _nw347
			var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2313_v56), 0))
			_ = _index322
			var _rhs390 bool = (_2239_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_2306_v50).F6(), (_dafny.Zero).Minus((_2306_v50).F6()))).Cardinality()), _dafny.IntOfUint32((_2239_v2).Cardinality()))).Uint32()).(bool)
			_ = _rhs390
			var _rhs391 *C9 = _2321_v64
			_ = _rhs391
			var _rhs392 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_2306_v50).F6(), (_2306_v50).F6()))
			_ = _rhs392
			var _rhs393 _dafny.Int = (_dafny.Zero).Minus((_this).F6())
			_ = _rhs393
			var _lhs311 _dafny.Array = _2313_v56
			_ = _lhs311
			var _lhs312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_2313_v56), 0))
			_ = _lhs312
			var _lhs313 *GlobalState = globalState
			_ = _lhs313
			_2238_v1 = _rhs390
			_2321_v64 = _rhs391
			(_lhs311).ArraySet1(_rhs392, (_lhs312).Int())
			_lhs313.F5 = _rhs393
			var _2322_v65 _dafny.Map
			_ = _2322_v65
			_2322_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2306_v50).F6(), (_2242_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_2242_v5), 0))).Int()).(bool))
			(globalState).F2 = _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_2322_v65).Cardinality(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _2269_v28)
			(globalState).F5 = (_2306_v50).F6()
			var _2323_v66 bool
			_ = _2323_v66
			var _out32 bool
			_ = _out32
			_out32 = (_this).M2(globalState)
			_2323_v66 = _out32
		}
		r0 = ((_this).F6()).Cmp((_this).F6()) == 0
		return r0
	}
}
func (_this *C10) M2(globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _2324_v0 bool
		_ = _2324_v0
		_2324_v0 = false
		(globalState).F2 = Companion_Default___.Fm18(_2324_v0, globalState)
		var _2325_v1 _dafny.Array
		_ = _2325_v1
		var _nwElement0_71 bool = (false) || (_2324_v0)
		_ = _nwElement0_71
		var _nw348 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(5))
		_ = _nw348
		(_nw348).ArraySet1(_nwElement0_71, 0)
		(_nw348).ArraySet1(_2324_v0, 1)
		(_nw348).ArraySet1(_2324_v0, 2)
		(_nw348).ArraySet1((_2324_v0) || (_2324_v0), 3)
		(_nw348).ArraySet1(_2324_v0, 4)
		_2325_v1 = _nw348
		var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))
		_ = _index323
		(_2325_v1).ArraySet1(_2324_v0, (_index323).Int())
		var _2326_v2 _dafny.Sequence
		_ = _2326_v2
		_2326_v2 = _dafny.SeqOf(_2324_v0)
		var _2327_v3 _dafny.MultiSet
		_ = _2327_v3
		_2327_v3 = _dafny.MultiSetOf(_2326_v2, _2326_v2)
		var _2328_v4 _dafny.Sequence
		_ = _2328_v4
		_2328_v4 = _dafny.UnicodeSeqOfUtf8Bytes("rpavxdbs")
		var _2329_v5 _dafny.Map
		_ = _2329_v5
		_2329_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _2328_v4)
		var _2330_v6 _dafny.Map
		_ = _2330_v6
		_2330_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
		var _2331_v7 _dafny.MultiSet
		_ = _2331_v7
		_2331_v7 = _dafny.MultiSetOf(_dafny.IntOfInt64(264))
		var _2332_v8 _dafny.MultiSet
		_ = _2332_v8
		_2332_v8 = _dafny.MultiSetOf((_2331_v7).Cardinality())
		var _2333_v9 _dafny.MultiSet
		_ = _2333_v9
		_2333_v9 = _dafny.MultiSetOf(_dafny.MultiSetOf((_this).F6()), _2332_v8, _2332_v8, _2332_v8, _2331_v7)
		var _2334_v10 *C4
		_ = _2334_v10
		var _nw349 *C4 = New_C4_()
		_ = _nw349
		_nw349.Ctor__(_dafny.IntOfUint32((_2328_v4).Cardinality()), (_this).F6(), _2325_v1, _2329_v5, (func() _dafny.Int {
			if (_2330_v6).Contains((_dafny.SetOf((_this).F6())).Cardinality()) {
				return (_2330_v6).Get((_dafny.SetOf((_this).F6())).Cardinality()).(_dafny.Int)
			}
			return (_this).F6()
		})(), _2333_v9)
		_2334_v10 = _nw349
		var _2335_v11 D18
		_ = _2335_v11
		_2335_v11 = Companion_D18_.Create_DC44_(_2324_v0, _2324_v0, _2327_v3, _2334_v10)
		var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))
		_ = _index324
		var _rhs394 bool = (_2335_v11).Dtor_cf80()
		_ = _rhs394
		var _rhs395 _dafny.Int = (_2334_v10).F17()
		_ = _rhs395
		var _rhs396 bool = _2324_v0
		_ = _rhs396
		var _rhs397 bool = _2324_v0
		_ = _rhs397
		var _rhs398 _dafny.Array = _2325_v1
		_ = _rhs398
		var _lhs314 *GlobalState = globalState
		_ = _lhs314
		var _lhs315 _dafny.Array = _2325_v1
		_ = _lhs315
		var _lhs316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))
		_ = _lhs316
		var _lhs317 *GlobalState = globalState
		_ = _lhs317
		r0 = _rhs394
		_lhs314.F5 = _rhs395
		(_lhs315).ArraySet1(_rhs396, (_lhs316).Int())
		_lhs317.F0 = _rhs397
		_2325_v1 = _rhs398
		if false {
			(globalState).F5 = (_2334_v10).Fm21((_2334_v10).F17(), globalState)
			var _2336_v12 _dafny.CodePoint
			_ = _2336_v12
			_2336_v12 = _dafny.CodePoint('f')
			var _2337_v13 _dafny.Set
			_ = _2337_v13
			_2337_v13 = _dafny.SetOf((_2334_v10).F17(), (_dafny.Zero).Minus((_2334_v10).F18()))
			var _2338_v14 D18
			_ = _2338_v14
			_2338_v14 = Companion_D18_.Create_DC43_(_dafny.IntOfInt64(266))
			var _2339_v15 _dafny.Sequence
			_ = _2339_v15
			_2339_v15 = _dafny.SeqOf((_2334_v10).F17(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xsvnce")).Cardinality()))
			var _2340_v17 _dafny.Set
			_ = _2340_v17
			_2340_v17 = _dafny.SetOf((func() _dafny.CodePoint {
				if false {
					return _2336_v12
				}
				return (_2328_v4).Select((Companion_Default___.SafeIndex((_2337_v13).Cardinality(), _dafny.IntOfUint32((_2328_v4).Cardinality()))).Uint32()).(_dafny.CodePoint)
			})(), _2336_v12, Companion_Default___.Fm0((func() _dafny.Int {
				if (_2330_v6).Contains((_2338_v14).Dtor_cf78()) {
					return (_2330_v6).Get((_2338_v14).Dtor_cf78()).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_2339_v15).Cardinality())
			})(), func() _dafny.Set {
				var _coll63 = _dafny.NewBuilder()
				_ = _coll63
				for _iter67 := _dafny.Iterate((_2339_v15).Elements()); ; {
					_compr_63, _ok67 := _iter67()
					if !_ok67 {
						break
					}
					var _2341_v16 _dafny.Int
					_2341_v16 = interface{}(_compr_63).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_2339_v15, _2341_v16) {
						_coll63.Add((_2341_v16).Times(_dafny.IntOfInt64(2)))
					}
				}
				return _coll63.ToSet()
			}(), (_2325_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))).Int()).(bool), globalState), _2336_v12)
			var _2342_v18 _dafny.Array
			_ = _2342_v18
			var _nw350 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
			_ = _nw350
			_2342_v18 = _nw350
			var _2343_v19 _dafny.Map
			_ = _2343_v19
			_2343_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2342_v18, _2340_v17)
			var _2344_v20 _dafny.Sequence
			_ = _2344_v20
			_2344_v20 = _dafny.SeqOf(_2326_v2, _2326_v2, _2326_v2, _2326_v2)
			var _2345_v21 D9
			_ = _2345_v21
			_2345_v21 = Companion_D9_.Create_DC25_(_dafny.IntOfUint32(((_2344_v20).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("q")).Cardinality()), _dafny.IntOfUint32((_2344_v20).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), (_2334_v10).F18(), (_2334_v10).F18(), (_2328_v4).Select((Companion_Default___.SafeIndex((_2334_v10).F18(), _dafny.IntOfUint32((_2328_v4).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _2346_v22 _dafny.Sequence
			_ = _2346_v22
			_2346_v22 = _dafny.SeqOf(Companion_D9_.Create_DC25_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qqakq")).Cardinality()), (_2334_v10).F18(), (Companion_Default___.Fm9(globalState)).Cardinality(), _2336_v12), _2345_v21)
			_2340_v17 = (func() _dafny.Set {
				if (_2343_v19).Contains(_2342_v18) {
					return (_2343_v19).Get(_2342_v18).(_dafny.Set)
				}
				return Companion_Default___.Fm54((_2325_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))).Int()).(bool), Companion_Default___.Fm0((_2334_v10).F17(), _2337_v13, (_2325_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))).Int()).(bool), globalState), _2346_v22, globalState)
			})()
			_2336_v12 = _2336_v12
			var _2347_v23 _dafny.Array
			_ = _2347_v23
			var _nw351 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw351
			_2347_v23 = _nw351
			var _2348_v24 _dafny.Array
			_ = _2348_v24
			var _nw352 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(22))
			_ = _nw352
			_2348_v24 = _nw352
			var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_2347_v23), 0))
			_ = _index325
			(_2347_v23).ArraySet1(_2348_v24, (_index325).Int())
			var _2349_v25 T0
			_ = _2349_v25
			var _nw353 *C9 = New_C9_()
			_ = _nw353
			_nw353.Ctor__((_2334_v10).F18())
			_2349_v25 = _nw353
			var _2350_v26 _dafny.Map
			_ = _2350_v26
			_2350_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_2349_v25).F6()), _2349_v25)
			var _2351_v27 T0
			_ = _2351_v27
			_2351_v27 = _2349_v25
			var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_2347_v23), 0))
			_ = _index326
			var _nwElement0_72 T0 = _2349_v25
			_ = _nwElement0_72
			var _nw354 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(20))
			_ = _nw354
			(_nw354).ArraySet1(_nwElement0_72, 0)
			(_nw354).ArraySet1(_2349_v25, 1)
			(_nw354).ArraySet1(_2349_v25, 2)
			(_nw354).ArraySet1(_2349_v25, 3)
			(_nw354).ArraySet1(_2349_v25, 4)
			(_nw354).ArraySet1(_2349_v25, 5)
			(_nw354).ArraySet1(_2349_v25, 6)
			(_nw354).ArraySet1(_2349_v25, 7)
			(_nw354).ArraySet1(_2349_v25, 8)
			(_nw354).ArraySet1(_2349_v25, 9)
			(_nw354).ArraySet1((func() T0 {
				if (_2350_v26).Contains((_2334_v10).F17()) {
					return (_2350_v26).Get((_2334_v10).F17()).(T0)
				}
				return _2349_v25
			})(), 10)
			(_nw354).ArraySet1(_2349_v25, 11)
			(_nw354).ArraySet1(_2349_v25, 12)
			(_nw354).ArraySet1(_2349_v25, 13)
			(_nw354).ArraySet1(_2349_v25, 14)
			(_nw354).ArraySet1((_2351_v27), 15)
			(_nw354).ArraySet1(_2349_v25, 16)
			(_nw354).ArraySet1(_2349_v25, 17)
			(_nw354).ArraySet1(_2349_v25, 18)
			(_nw354).ArraySet1((_2349_v25), 19)
			(_2347_v23).ArraySet1(_nw354, (_index326).Int())
			var _2352_v28 _dafny.Map
			_ = _2352_v28
			_2352_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2324_v0, (_2334_v10).F17())
			var _2353_v29 _dafny.Map
			_ = _2353_v29
			_2353_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_2325_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))).Int()).(bool)), ((_2352_v28).Merge(_2352_v28)).Cardinality())
			_2352_v28 = (_2352_v28).Update(_dafny.Companion_Sequence_.IsProperPrefixOf(_2328_v4, _dafny.UnicodeSeqOfUtf8Bytes("ywxdsc")), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2339_v15, _2339_v15)).Cardinality()))
		} else {
			var _2354_v30 D0
			_ = _2354_v30
			_2354_v30 = Companion_D0_.Create_DC0_(_dafny.SeqOf(_2324_v0, false, (_2325_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))).Int()).(bool), (_2325_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))).Int()).(bool)), (_2334_v10).F18())
			var _2355_v31 _dafny.Array
			_ = _2355_v31
			var _len0_58 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_58
			var _nw355 _dafny.Array
			_ = _nw355
			if _len0_58.Cmp(_dafny.Zero) == 0 {
				_nw355 = _dafny.NewArray(_len0_58)
			} else {
				var _init58 func(_dafny.Int) _dafny.Int = func(_2356_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_2356_i0, (_this).F6())
				}
				_ = _init58
				var _element0_58 = _init58(_dafny.Zero)
				_ = _element0_58
				_nw355 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
				(_nw355).ArraySet1(_element0_58, 0)
				var _nativeLen0_58 = (_len0_58).Int()
				_ = _nativeLen0_58
				for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
					(_nw355).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
				}
			}
			_2355_v31 = _nw355
			var _2357_v32 _dafny.MultiSet
			_ = _2357_v32
			_2357_v32 = _dafny.MultiSetOf(_2355_v31, _2355_v31)
			var _2358_v33 _dafny.Map
			_ = _2358_v33
			_2358_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_2325_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))).Int()).(bool)), _2355_v31)
			var _2359_v34 _dafny.Map
			_ = _2359_v34
			_2359_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2354_v30, (_2357_v32).Update((func() _dafny.Array {
				if (_2358_v33).Contains((_2325_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))).Int()).(bool)) {
					return (_2358_v33).Get((_2325_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))).Int()).(bool)).(_dafny.Array)
				}
				return _2355_v31
			})(), Companion_Default___.Abs((_this).F6())))
			_2359_v34 = (_2359_v34).Update(_2354_v30, (_2357_v32).Difference(_2357_v32))
			var _2360_v35 _dafny.Sequence
			_ = _2360_v35
			_2360_v35 = _dafny.SeqOf((_2334_v10).F17(), (_2334_v10).F18(), _dafny.IntOfInt64(-387), (_2334_v10).F18())
			var _2361_v36 _dafny.Map
			_ = _2361_v36
			_2361_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_2360_v35, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-334), _dafny.IntOfUint32((_2360_v35).Cardinality()))).Uint32(), (_2334_v10).F18()), _dafny.SeqOf(Companion_Default___.Fm55(globalState)))
			_2361_v36 = _2361_v36
			(globalState).F0 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_2360_v35, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(249))).Uint32(), func(coer128 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg128 _dafny.Int) interface{} {
					return coer128(arg128)
				}
			}((func(_2362_v10 *C4) func(_dafny.Int) _dafny.Int {
				return func(_2363_i1 _dafny.Int) _dafny.Int {
					return (_2362_v10).F17()
				}
			})(_2334_v10)))), _2360_v35)
			var _2364_v37 _dafny.Map
			_ = _2364_v37
			_2364_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2324_v0, _dafny.IntOfInt64(788))
			var _2365_v38 _dafny.Sequence
			_ = _2365_v38
			_2365_v38 = _dafny.SeqOf(_2364_v37)
			(globalState).F5 = ((((_2365_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2360_v35).Cardinality()), _dafny.IntOfUint32((_2365_v38).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_2364_v37)).Cardinality()).Times(((_dafny.Zero).Minus((_2334_v10).F17())).Plus((_2334_v10).F18()))
			var _2366_v39 T2
			_ = _2366_v39
			var _nw356 *C4 = New_C4_()
			_ = _nw356
			_nw356.Ctor__((_2334_v10).F18(), _dafny.IntOfInt64(82), _2325_v1, (_2329_v5).Update((_2334_v10).F18(), _2328_v4), (_this).F6(), _2333_v9)
			_2366_v39 = _nw356
			var _2367_v40 _dafny.Map
			_ = _2367_v40
			_2367_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2334_v10).F17(), _2366_v39)
			var _2368_v41 _dafny.Map
			_ = _2368_v41
			_2368_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2367_v40, (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality())).Cmp((_2334_v10).F18()) > 0)
			_2368_v41 = (_2368_v41).Update(_2367_v40, _2324_v0)
		}
		var _2369_v42 _dafny.Array
		_ = _2369_v42
		var _nw357 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(28))
		_ = _nw357
		_2369_v42 = _nw357
		var _2370_v43 _dafny.CodePoint
		_ = _2370_v43
		_2370_v43 = _dafny.CodePoint('t')
		var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2369_v42), 0))
		_ = _index327
		(_2369_v42).ArraySet1CodePoint(_2370_v43, (_index327).Int())
		var _2371_v44 D14
		_ = _2371_v44
		_2371_v44 = Companion_D14_.Create_DC37_((_2325_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))).Int()).(bool), _2328_v4)
		var _pat_let_tv58 = _2334_v10
		_ = _pat_let_tv58
		var _pat_let_tv59 = _2324_v0
		_ = _pat_let_tv59
		var _pat_let_tv60 = _2334_v10
		_ = _pat_let_tv60
		var _pat_let_tv61 = _2334_v10
		_ = _pat_let_tv61
		var _pat_let_tv62 = _2370_v43
		_ = _pat_let_tv62
		var _pat_let_tv63 = _2325_v1
		_ = _pat_let_tv63
		var _pat_let_tv64 = _2325_v1
		_ = _pat_let_tv64
		var _pat_let_tv65 = _2325_v1
		_ = _pat_let_tv65
		var _pat_let_tv66 = _2325_v1
		_ = _pat_let_tv66
		var _pat_let_tv67 = _2370_v43
		_ = _pat_let_tv67
		var _pat_let_tv68 = _2370_v43
		_ = _pat_let_tv68
		var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2369_v42), 0))
		_ = _index328
		(_2369_v42).ArraySet1CodePoint(func(_source31 D14) _dafny.CodePoint {
			if _source31.Is_DC37() {
				var _2372___mcc_h0 bool = _source31.Get_().(D14_DC37).Cf70
				_ = _2372___mcc_h0
				var _2373___mcc_h1 _dafny.Sequence = _source31.Get_().(D14_DC37).Cf71
				_ = _2373___mcc_h1
				var _2374_cf71 _dafny.Sequence = _2373___mcc_h1
				_ = _2374_cf71
				var _2375_cf70 bool = _2372___mcc_h0
				_ = _2375_cf70
				return (Companion_D9_.Create_DC25_((_pat_let_tv58).F17(), (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv59, _dafny.IntOfUint32((_dafny.SeqOf((_pat_let_tv60).F17())).Cardinality()))).Cardinality(), (_pat_let_tv61).F18())).Cardinality(), _dafny.IntOfInt64(-684), _pat_let_tv62)).Dtor_cf49()
			} else {
				var _2376___mcc_h2 _dafny.Map = _source31.Get_().(D14_DC36).Cf69
				_ = _2376___mcc_h2
				var _2377_cf69 _dafny.Map = _2376___mcc_h2
				_ = _2377_cf69
				if (_pat_let_tv64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_pat_let_tv63), 0))).Int()).(bool) {
					return (Companion_D0_.Create_DC1_((_pat_let_tv66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_pat_let_tv65), 0))).Int()).(bool), (_this).F6(), _pat_let_tv67)).Dtor_cf4()
				} else {
					return _pat_let_tv68
				}
			}
		}(_2371_v44), (_index328).Int())
		var _2378_v45 _dafny.Sequence
		_ = _2378_v45
		_2378_v45 = _dafny.SeqOf(_dafny.IntOfInt64(701))
		var _rhs399 _dafny.Int = (_2378_v45).Select((Companion_Default___.SafeIndex((_2334_v10).F18(), _dafny.IntOfUint32((_2378_v45).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs399
		var _rhs400 _dafny.Int = (_2378_v45).Select((Companion_Default___.SafeIndex(((Companion_Default___.Fm56(_2328_v4, _2324_v0, _2328_v4, globalState)).Dtor_cf59()).Minus((_this).F6()), _dafny.IntOfUint32((_2378_v45).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs400
		var _rhs401 _dafny.Int = (_2334_v10).F17()
		_ = _rhs401
		var _lhs318 *GlobalState = globalState
		_ = _lhs318
		var _lhs319 *GlobalState = globalState
		_ = _lhs319
		var _lhs320 *GlobalState = globalState
		_ = _lhs320
		_lhs318.F5 = _rhs399
		_lhs319.F5 = _rhs400
		_lhs320.F5 = _rhs401
		var _2379_v46 _dafny.Map
		_ = _2379_v46
		_2379_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _2324_v0)
		var _2380_v47 _dafny.Set
		_ = _2380_v47
		_2380_v47 = _dafny.SetOf(_2379_v46, (_2379_v46).Update((_2325_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))).Int()).(bool), true), _2379_v46)
		(globalState).F5 = (_2380_v47).Cardinality()
		r0 = (_2325_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2325_v1), 0))).Int()).(bool)
		return r0
	}
}

// End of class C10

// Definition of class C11
type C11 struct {
	_f6 _dafny.Int
	F8  bool
	_f7 _dafny.Array
}

func New_C11_() *C11 {
	_this := C11{}

	_this._f6 = _dafny.Zero
	_this.F8 = false
	_this._f7 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	return &_this
}

type CompanionStruct_C11_ struct {
}

var Companion_C11_ = CompanionStruct_C11_{}

func (_this *C11) Equals(other *C11) bool {
	return _this == other
}

func (_this *C11) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C11)
	return ok && _this.Equals(other)
}

func (*C11) String() string {
	return "_module.C11"
}

func Type_C11_() _dafny.TypeDescriptor {
	return type_C11_{}
}

type type_C11_ struct {
}

func (_this type_C11_) Default() interface{} {
	return (*C11)(nil)
}

func (_this type_C11_) String() string {
	return "main.C11"
}
func (_this *C11) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C11{}
var _ _dafny.TraitOffspring = &C11{}

func (_this *C11) F6() _dafny.Int {
	return _this._f6
}
func (_this *C11) Ctor__(f7 _dafny.Array, f8 bool, f6 _dafny.Int) {
	{
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this)._f6 = f6
	}
}
func (_this *C11) Fm5(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		var _source32 D0 = (func() D0 {
			if _this.F8 {
				return Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(_this.F8, (_this).F6(), _dafny.CodePoint('l')))
			}
			return Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_(_dafny.SeqOf(_this.F8), (_this).F6()))
		})()
		_ = _source32
		if _source32.Is_DC0() {
			var _2381___mcc_h0 _dafny.Sequence = _source32.Get_().(D0_DC0).Cf0
			_ = _2381___mcc_h0
			var _2382___mcc_h1 _dafny.Int = _source32.Get_().(D0_DC0).Cf1
			_ = _2382___mcc_h1
			var _2383_cf1 _dafny.Int = _2382___mcc_h1
			_ = _2383_cf1
			var _2384_cf0 _dafny.Sequence = _2381___mcc_h0
			_ = _2384_cf0
			return _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('y'))))
		} else if _source32.Is_DC1() {
			var _2385___mcc_h2 bool = _source32.Get_().(D0_DC1).Cf2
			_ = _2385___mcc_h2
			var _2386___mcc_h3 _dafny.Int = _source32.Get_().(D0_DC1).Cf3
			_ = _2386___mcc_h3
			var _2387___mcc_h4 _dafny.CodePoint = _source32.Get_().(D0_DC1).Cf4
			_ = _2387___mcc_h4
			var _2388_cf4 _dafny.CodePoint = _2387___mcc_h4
			_ = _2388_cf4
			var _2389_cf3 _dafny.Int = _2386___mcc_h3
			_ = _2389_cf3
			var _2390_cf2 bool = _2385___mcc_h2
			_ = _2390_cf2
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_2388_cf4)), _dafny.SeqOf(_dafny.MultiSetOf(_2388_cf4, _2388_cf4)))
		} else {
			var _2391___mcc_h5 D0 = _source32.Get_().(D0_DC2).Cf5
			_ = _2391___mcc_h5
			var _2392_cf5 D0 = _2391___mcc_h5
			_ = _2392_cf5
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(538))).Uint32(), func(coer129 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg129 _dafny.Int) interface{} {
					return coer129(arg129)
				}
			}(func(_2393_i0 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('f')))
			})), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('g')), (Companion_D1_.Create_DC3_(_dafny.MultiSetOf(_dafny.CodePoint('j')))).Dtor_cf6(), _dafny.MultiSetOf(_dafny.CodePoint('k'))))
		}
	}
}
func (_this *C11) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.SeqOf((_this).F6()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.SeqOf((_dafny.SetOf(_this.F8, _this.F8)).Cardinality(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_this.F8)).Cardinality()))).Cardinality(), (_this).F6(), (_this).F6())))
	}
}
func (_this *C11) Fm7(p0 bool, p1 _dafny.Sequence, p2 _dafny.Set, p3 bool, globalState *GlobalState) bool {
	{
		return _this.F8
	}
}
func (_this *C11) Fm8(globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (func() D0 {
			if _this.F8 {
				return Companion_D0_.Create_DC0_(_dafny.SeqOf(false), (_this).F6())
			}
			return Companion_D0_.Create_DC0_(_dafny.SeqOf(_this.F8), (_dafny.Zero).Minus((_this).F6()))
		})())
	}
}
func (_this *C11) M0(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Array, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var _2394_v0 _dafny.Sequence
		_ = _2394_v0
		_2394_v0 = _dafny.SeqOf((_this).F6())
		var _2395_v1 _dafny.Set
		_ = _2395_v1
		_2395_v1 = _dafny.SetOf(_this.F8, _this.F8)
		var _2396_v2 _dafny.Sequence
		_ = _2396_v2
		_2396_v2 = _dafny.SeqOf(p1, _this.F8, false, _this.F8)
		var _2397_v3 _dafny.Sequence
		_ = _2397_v3
		_2397_v3 = _dafny.UnicodeSeqOfUtf8Bytes("tbon")
		var _2398_v4 D0
		_ = _2398_v4
		_2398_v4 = Companion_D0_.Create_DC1_(_this.F8, (_this).F6(), _dafny.CodePoint('p'))
		var _2399_v5 _dafny.MultiSet
		_ = _2399_v5
		_2399_v5 = _dafny.MultiSetOf((_this).F6())
		var _2400_v6 _dafny.MultiSet
		_ = _2400_v6
		_2400_v6 = _dafny.MultiSetOf(_2399_v5)
		var _2401_v7 _dafny.Map
		_ = _2401_v7
		_2401_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.UnicodeSeqOfUtf8Bytes("ogtymm"))
		var _2402_v8 T0
		_ = _2402_v8
		var _nw358 *C2 = New_C2_()
		_ = _nw358
		_nw358.Ctor__(_dafny.IntOfUint32((_2397_v3).Cardinality()), _2400_v6, p2, _2401_v7)
		_2402_v8 = _nw358
		var _2403_v9 _dafny.Sequence
		_ = _2403_v9
		_2403_v9 = _dafny.SeqOf(_2402_v8)
		var _2404_v10 D9
		_ = _2404_v10
		_2404_v10 = Companion_D9_.Create_DC24_(_2400_v6)
		var _2405_v11 _dafny.Sequence
		_ = _2405_v11
		_2405_v11 = _dafny.SeqOf(_2404_v10, _2404_v10, _2404_v10)
		var _2406_v12 _dafny.Array
		_ = _2406_v12
		var _nwElement0_73 bool = !(p1) || (_this.F8)
		_ = _nwElement0_73
		var _nw359 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(20))
		_ = _nw359
		(_nw359).ArraySet1(_nwElement0_73, 0)
		(_nw359).ArraySet1(_this.F8, 1)
		(_nw359).ArraySet1((_dafny.IntOfInt64(365)).Cmp((_this).F6()) >= 0, 2)
		(_nw359).ArraySet1(_this.F8, 3)
		(_nw359).ArraySet1(true, 4)
		(_nw359).ArraySet1(_dafny.Companion_Sequence_.Equal(_2394_v0, _2394_v0), 5)
		(_nw359).ArraySet1(((_dafny.Zero).Minus((_2395_v1).Cardinality())).Cmp((_dafny.Zero).Minus((_this).F6())) == 0, 6)
		(_nw359).ArraySet1(!_dafny.Companion_Sequence_.Contains(_2396_v2, p1), 7)
		(_nw359).ArraySet1(p1, 8)
		(_nw359).ArraySet1(false, 9)
		(_nw359).ArraySet1(!((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2397_v3, (Companion_Default___.SafeIndex((_2398_v4).Dtor_cf3(), _dafny.IntOfUint32((_2397_v3).Cardinality()))).Uint32(), p0)).Cardinality())).Cmp((_this).F6()) <= 0), 10)
		(_nw359).ArraySet1(p1, 11)
		(_nw359).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_2403_v9, _2403_v9), 12)
		(_nw359).ArraySet1(p1, 13)
		(_nw359).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_2405_v11, _2405_v11), 14)
		(_nw359).ArraySet1((_2396_v2).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2396_v2).Cardinality()))).Uint32()).(bool), 15)
		(_nw359).ArraySet1(((_2402_v8).F6()).Cmp(_dafny.IntOfInt64(746)) < 0, 16)
		(_nw359).ArraySet1(!(_this.F8), 17)
		(_nw359).ArraySet1(p1, 18)
		(_nw359).ArraySet1((_dafny.IntOfUint32((_2397_v3).Cardinality())).Cmp((_2402_v8).F6()) >= 0, 19)
		_2406_v12 = _nw359
		_2406_v12 = _2406_v12
		(globalState).F5 = (_2402_v8).F6()
		(globalState).F5 = (_2402_v8).F6()
		var _2407_v13 *C4
		_ = _2407_v13
		var _nw360 *C4 = New_C4_()
		_ = _nw360
		_nw360.Ctor__((_2402_v8).F6(), (_this).F6(), p2, _2401_v7, (_this).F6(), _2400_v6)
		_2407_v13 = _nw360
		var _2408_v14 _dafny.MultiSet
		_ = _2408_v14
		_2408_v14 = _dafny.MultiSetOf(_2407_v13, _2407_v13)
		var _hi14 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
			if p1 {
				return (_2395_v1).Cardinality()
			}
			return (_2402_v8).F6()
		})())
		_ = _hi14
		for _2409_i0 := (func() _dafny.Int {
			if (_2408_v14).Contains(_2407_v13) {
				return (_2408_v14).Multiplicity(_2407_v13)
			}
			return (_2407_v13).F17()
		})(); _2409_i0.Cmp(_hi14) < 0; _2409_i0 = _2409_i0.Plus(_dafny.One) {
			var _2410_v15 _dafny.Map
			_ = _2410_v15
			_2410_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_2407_v13).F18()).Plus((_2407_v13).F17()), (_dafny.Zero).Minus((_2407_v13).F18()))
			_2410_v15 = (func() _dafny.Map {
				var _coll64 = _dafny.NewMapBuilder()
				_ = _coll64
				for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-162), _dafny.IntOfInt64(475))); ; {
					_compr_64, _ok68 := _iter68()
					if !_ok68 {
						break
					}
					var _2411_v16 _dafny.Int
					_2411_v16 = interface{}(_compr_64).(_dafny.Int)
					if ((_dafny.IntOfInt64(-162)).Cmp(_2411_v16) <= 0) && ((_2411_v16).Cmp(_dafny.IntOfInt64(475)) < 0) {
						_coll64.Add((_2411_v16).Plus((_dafny.MultiSetOf((_2407_v13).F18(), (_2402_v8).F6())).Cardinality()), (_this).F6())
					}
				}
				return _coll64.ToMap()
			}()).Merge(_2410_v15)
			if _this.F8 {
				var _2412_v17 _dafny.Array
				_ = _2412_v17
				var _len0_59 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_59
				var _nw361 _dafny.Array
				_ = _nw361
				if _len0_59.Cmp(_dafny.Zero) == 0 {
					_nw361 = _dafny.NewArray(_len0_59)
				} else {
					var _init59 func(_dafny.Int) _dafny.Int = (func(_2413_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_2414_i1 _dafny.Int) _dafny.Int {
							return (_2414_i1).Plus(_dafny.IntOfUint32((_2413_v3).Cardinality()))
						}
					})(_2397_v3)
					_ = _init59
					var _element0_59 = _init59(_dafny.Zero)
					_ = _element0_59
					_nw361 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
					(_nw361).ArraySet1(_element0_59, 0)
					var _nativeLen0_59 = (_len0_59).Int()
					_ = _nativeLen0_59
					for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
						(_nw361).ArraySet1(_init59(_dafny.IntOf(_i0_59)), _i0_59)
					}
				}
				_2412_v17 = _nw361
				var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_2412_v17), 0))
				_ = _index329
				(_2412_v17).ArraySet1((_dafny.IntOfInt64(291)).Times((_2407_v13).F17()), (_index329).Int())
				var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_2412_v17), 0))
				_ = _index330
				var _rhs402 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2397_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(750))).Uint32(), func(coer130 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg130 _dafny.Int) interface{} {
						return coer130(arg130)
					}
				}(func(_2415_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				}))), (Companion_Default___.SafeIndex((_2402_v8).F6(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2397_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(750))).Uint32(), func(coer131 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg131 _dafny.Int) interface{} {
						return coer131(arg131)
					}
				}(func(_2416_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				})))).Cardinality()))).Uint32(), p0)).Cardinality())
				_ = _rhs402
				var _rhs403 bool = _this.F8
				_ = _rhs403
				var _rhs404 _dafny.Sequence = _2397_v3
				_ = _rhs404
				var _rhs405 bool = p1
				_ = _rhs405
				var _lhs321 _dafny.Array = _2412_v17
				_ = _lhs321
				var _lhs322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_2412_v17), 0))
				_ = _lhs322
				var _lhs323 *GlobalState = globalState
				_ = _lhs323
				var _lhs324 *GlobalState = globalState
				_ = _lhs324
				(_lhs321).ArraySet1(_rhs402, (_lhs322).Int())
				_lhs323.F0 = _rhs403
				_2397_v3 = _rhs404
				_lhs324.F0 = _rhs405
				var _2417_v18 D14
				_ = _2417_v18
				_2417_v18 = Companion_D14_.Create_DC37_(_this.F8, _2397_v3)
				(globalState).F0 = (_2417_v18).Dtor_cf70()
				(globalState).F5 = (_2394_v0).Select((Companion_Default___.SafeIndex((_2407_v13).F18(), _dafny.IntOfUint32((_2394_v0).Cardinality()))).Uint32()).(_dafny.Int)
				var _2418_v19 _dafny.Map
				_ = _2418_v19
				_2418_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2402_v8, _2406_v12)
				_2406_v12 = (func() _dafny.Array {
					if (_2418_v19).Contains(_2402_v8) {
						return (_2418_v19).Get(_2402_v8).(_dafny.Array)
					}
					return p2
				})()
				var _2419_v20 _dafny.CodePoint
				_ = _2419_v20
				var _2420_v21 _dafny.Int
				_ = _2420_v21
				var _2421_v22 _dafny.Array
				_ = _2421_v22
				var _2422_v23 _dafny.Int
				_ = _2422_v23
				var _out33 _dafny.CodePoint
				_ = _out33
				var _out34 _dafny.Int
				_ = _out34
				var _out35 _dafny.Array
				_ = _out35
				var _out36 _dafny.Int
				_ = _out36
				_out33, _out34, _out35, _out36 = (_2407_v13).M7((_2399_v5).Union(_2399_v5), (_2407_v13).F18(), globalState)
				_2419_v20 = _out33
				_2420_v21 = _out34
				_2421_v22 = _out35
				_2422_v23 = _out36
			} else {
				var _2423_v24 _dafny.Array
				_ = _2423_v24
				var _nw362 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw362
				_2423_v24 = _nw362
				var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_2423_v24), 0))
				_ = _index331
				(_2423_v24).ArraySet1((_dafny.IntOfUint32((_2396_v2).Cardinality())).Times((_2407_v13).F17()), (_index331).Int())
				var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_2423_v24), 0))
				_ = _index332
				var _rhs406 _dafny.Int = (_2407_v13).F18()
				_ = _rhs406
				var _rhs407 _dafny.Int = (_2407_v13).F17()
				_ = _rhs407
				var _lhs325 *GlobalState = globalState
				_ = _lhs325
				var _lhs326 _dafny.Array = _2423_v24
				_ = _lhs326
				var _lhs327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_2423_v24), 0))
				_ = _lhs327
				_lhs325.F5 = _rhs406
				(_lhs326).ArraySet1(_rhs407, (_lhs327).Int())
				var _rhs408 _dafny.MultiSet = ((_2399_v5).Update((_2402_v8).F6(), Companion_Default___.Abs(Companion_Default___.Fm2((_2402_v8).F6(), globalState)))).Difference(_2399_v5)
				_ = _rhs408
				var _rhs409 _dafny.Map = (_2410_v15).Merge(_2410_v15)
				_ = _rhs409
				var _rhs410 bool = !(((_2402_v8).F6()).Cmp(Companion_Default___.SafeModuloInt((_2407_v13).F18(), _dafny.IntOfInt64(747))) >= 0)
				_ = _rhs410
				var _rhs411 _dafny.Int = (_2402_v8).F6()
				_ = _rhs411
				var _rhs412 _dafny.Array = _2406_v12
				_ = _rhs412
				var _lhs328 *GlobalState = globalState
				_ = _lhs328
				var _lhs329 *GlobalState = globalState
				_ = _lhs329
				_2399_v5 = _rhs408
				_2410_v15 = _rhs409
				_lhs328.F0 = _rhs410
				_lhs329.F5 = _rhs411
				_2406_v12 = _rhs412
				var _2424_v25 _dafny.Array
				_ = _2424_v25
				var _nwElement0_74 T0 = _2402_v8
				_ = _nwElement0_74
				var _nw363 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(23))
				_ = _nw363
				(_nw363).ArraySet1(_nwElement0_74, 0)
				(_nw363).ArraySet1(_2402_v8, 1)
				(_nw363).ArraySet1(_2402_v8, 2)
				(_nw363).ArraySet1(_2402_v8, 3)
				(_nw363).ArraySet1(_2402_v8, 4)
				(_nw363).ArraySet1(_2402_v8, 5)
				(_nw363).ArraySet1(_2402_v8, 6)
				(_nw363).ArraySet1(_2402_v8, 7)
				(_nw363).ArraySet1(_2402_v8, 8)
				(_nw363).ArraySet1(_2402_v8, 9)
				(_nw363).ArraySet1(_2402_v8, 10)
				(_nw363).ArraySet1(_2402_v8, 11)
				(_nw363).ArraySet1(_2402_v8, 12)
				(_nw363).ArraySet1(_2402_v8, 13)
				(_nw363).ArraySet1(_2402_v8, 14)
				(_nw363).ArraySet1(_2402_v8, 15)
				(_nw363).ArraySet1(_2402_v8, 16)
				(_nw363).ArraySet1(_2402_v8, 17)
				(_nw363).ArraySet1(_2402_v8, 18)
				(_nw363).ArraySet1(_2402_v8, 19)
				(_nw363).ArraySet1(_2402_v8, 20)
				(_nw363).ArraySet1(_2402_v8, 21)
				(_nw363).ArraySet1(_2402_v8, 22)
				_2424_v25 = _nw363
				var _2425_v26 _dafny.Map
				_ = _2425_v26
				_2425_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2424_v25, _2423_v24)
				_2425_v26 = (_2425_v26).Update(_2424_v25, _2423_v24)
				var _2426_v27 D2
				_ = _2426_v27
				var _2427_v28 _dafny.Sequence
				_ = _2427_v28
				var _2428_v29 _dafny.Sequence
				_ = _2428_v29
				var _out37 D2
				_ = _out37
				var _out38 _dafny.Sequence
				_ = _out38
				var _out39 _dafny.Sequence
				_ = _out39
				_out37, _out38, _out39 = (_2407_v13).M13((_2407_v13).F18(), globalState)
				_2426_v27 = _out37
				_2427_v28 = _out38
				_2428_v29 = _out39
				var _2429_v30 _dafny.MultiSet
				_ = _2429_v30
				_2429_v30 = _dafny.MultiSetOf(_2396_v2)
				var _2430_v31 D18
				_ = _2430_v31
				_2430_v31 = Companion_D18_.Create_DC44_(true, p1, (_2429_v30).Intersection(_2429_v30), _2407_v13)
				_2430_v31 = _2430_v31
			}
			var _2431_v32 *C5
			_ = _2431_v32
			var _nw364 *C5 = New_C5_()
			_ = _nw364
			_nw364.Ctor__((_2407_v13).F18(), Companion_Default___.SafeModuloInt((_this).F6(), (_2402_v8).F6()))
			_2431_v32 = _nw364
			var _2432_v33 _dafny.Map
			_ = _2432_v33
			_2432_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.CodePoint('n'))
			var _rhs413 _dafny.Int = (_2407_v13).F17()
			_ = _rhs413
			var _rhs414 _dafny.Int = ((func() _dafny.Map {
				if _this.F8 {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_2397_v3, (Companion_Default___.SafeIndex((_2431_v32).F15(), _dafny.IntOfUint32((_2397_v3).Cardinality()))).Uint32(), p0), (Companion_Default___.SafeIndex((_2402_v8).F6(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2397_v3, (Companion_Default___.SafeIndex((_2431_v32).F15(), _dafny.IntOfUint32((_2397_v3).Cardinality()))).Uint32(), p0)).Cardinality()))).Uint32(), p0)).Cardinality()), p0)
				}
				return _2432_v33
			})()).Cardinality()
			_ = _rhs414
			var _rhs415 _dafny.Int = (_2407_v13).F18()
			_ = _rhs415
			var _rhs416 *C5 = _2431_v32
			_ = _rhs416
			var _lhs330 *GlobalState = globalState
			_ = _lhs330
			var _lhs331 *GlobalState = globalState
			_ = _lhs331
			var _lhs332 *GlobalState = globalState
			_ = _lhs332
			_lhs330.F5 = _rhs413
			_lhs331.F5 = _rhs414
			_lhs332.F5 = _rhs415
			_2431_v32 = _rhs416
			(globalState).F0 = p1
		}
		var _2433_v34 _dafny.Array
		_ = _2433_v34
		var _len0_60 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_60
		var _nw365 _dafny.Array
		_ = _nw365
		if _len0_60.Cmp(_dafny.Zero) == 0 {
			_nw365 = _dafny.NewArray(_len0_60)
		} else {
			var _init60 func(_dafny.Int) _dafny.Int = func(_2434_i3 _dafny.Int) _dafny.Int {
				return (_2434_i3).Times((_this).F6())
			}
			_ = _init60
			var _element0_60 = _init60(_dafny.Zero)
			_ = _element0_60
			_nw365 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
			(_nw365).ArraySet1(_element0_60, 0)
			var _nativeLen0_60 = (_len0_60).Int()
			_ = _nativeLen0_60
			for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
				(_nw365).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
			}
		}
		_2433_v34 = _nw365
		var _2435_v35 *C7
		_ = _2435_v35
		var _nw366 *C7 = New_C7_()
		_ = _nw366
		_nw366.Ctor__(_2433_v34, _2433_v34, (_2402_v8).F6(), _2406_v12, _2401_v7)
		_2435_v35 = _nw366
		var _2436_v36 *C2
		_ = _2436_v36
		var _nw367 *C2 = New_C2_()
		_ = _nw367
		_nw367.Ctor__(_dafny.IntOfUint32((_2396_v2).Cardinality()), _dafny.MultiSetOf(_dafny.MultiSetOf((_2407_v13).F17(), (_2402_v8).F6()), _2399_v5), p2, _2401_v7)
		_2436_v36 = _nw367
		var _2437_v37 _dafny.Map
		_ = _2437_v37
		_2437_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2436_v36, true)
		var _source33 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2436_v36, Companion_Default___.Fm3(globalState))
		_ = _source33
		var _2438___mcc_h0 _dafny.Map = _source33
		_ = _2438___mcc_h0
		var _2439_cf68 _dafny.Map = _2438___mcc_h0
		_ = _2439_cf68
		var _2440_v38 _dafny.CodePoint
		_ = _2440_v38
		_2440_v38 = _dafny.CodePoint('a')
		var _rhs417 _dafny.Sequence = _2396_v2
		_ = _rhs417
		var _rhs418 _dafny.CodePoint = p0
		_ = _rhs418
		_2396_v2 = _rhs417
		_2440_v38 = _rhs418
		var _2441_v39 _dafny.Array
		_ = _2441_v39
		var _nw368 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
		_ = _nw368
		_2441_v39 = _nw368
		var _2442_v40 _dafny.Array
		_ = _2442_v40
		var _nwElement0_75 _dafny.Array = _2441_v39
		_ = _nwElement0_75
		var _nw369 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(27))
		_ = _nw369
		(_nw369).ArraySet1(_nwElement0_75, 0)
		(_nw369).ArraySet1(_2441_v39, 1)
		(_nw369).ArraySet1(_2441_v39, 2)
		(_nw369).ArraySet1(_2441_v39, 3)
		(_nw369).ArraySet1(_2441_v39, 4)
		(_nw369).ArraySet1(_2441_v39, 5)
		(_nw369).ArraySet1(_2441_v39, 6)
		(_nw369).ArraySet1(_2441_v39, 7)
		(_nw369).ArraySet1(_2441_v39, 8)
		(_nw369).ArraySet1(_2441_v39, 9)
		(_nw369).ArraySet1(_2441_v39, 10)
		(_nw369).ArraySet1(_2441_v39, 11)
		(_nw369).ArraySet1(_2441_v39, 12)
		(_nw369).ArraySet1(_2441_v39, 13)
		(_nw369).ArraySet1(_2441_v39, 14)
		(_nw369).ArraySet1(_2441_v39, 15)
		(_nw369).ArraySet1(_2441_v39, 16)
		(_nw369).ArraySet1(_2441_v39, 17)
		(_nw369).ArraySet1(_2441_v39, 18)
		(_nw369).ArraySet1(_2441_v39, 19)
		(_nw369).ArraySet1(_2441_v39, 20)
		(_nw369).ArraySet1(_2441_v39, 21)
		(_nw369).ArraySet1(_2441_v39, 22)
		(_nw369).ArraySet1(_2441_v39, 23)
		(_nw369).ArraySet1(_2441_v39, 24)
		(_nw369).ArraySet1(_2441_v39, 25)
		(_nw369).ArraySet1(_2441_v39, 26)
		_2442_v40 = _nw369
		var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_2442_v40), 0))
		_ = _index333
		(_2442_v40).ArraySet1(_2441_v39, (_index333).Int())
		var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_2442_v40), 0))
		_ = _index334
		var _rhs419 _dafny.Array = _2441_v39
		_ = _rhs419
		var _rhs420 bool = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(977), (_2402_v8).F6())).Cmp((_2407_v13).F18()) < 0
		_ = _rhs420
		var _lhs333 _dafny.Array = _2442_v40
		_ = _lhs333
		var _lhs334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_2442_v40), 0))
		_ = _lhs334
		var _lhs335 *C11 = _this
		_ = _lhs335
		(_lhs333).ArraySet1(_rhs419, (_lhs334).Int())
		_lhs335.F8 = _rhs420
		(_this).F8 = true
		var _2443_v41 _dafny.Array
		_ = _2443_v41
		var _nw370 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
		_ = _nw370
		_2443_v41 = _nw370
		var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_2443_v41), 0))
		_ = _index335
		(_2443_v41).ArraySet1(_2397_v3, (_index335).Int())
		var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_2443_v41), 0))
		_ = _index336
		(_2443_v41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2397_v3, _dafny.Companion_Sequence_.Concatenate(_2397_v3, _2397_v3)), (_index336).Int())
		r0 = _2397_v3
		return r0
	}
}
func (_this *C11) F7() _dafny.Array {
	{
		return _this._f7
	}
}

// End of class C11
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
