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
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.CodePoint {
	if (_dafny.IntOfInt64(608)).Cmp(_dafny.IntOfInt64(898)) != 0 {
		return _dafny.CodePoint('v')
	} else {
		return _dafny.CodePoint('l')
	}
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(100))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return (Companion_D1_.Create_DC1_(_dafny.CodePoint('d'))).Dtor_cf1()
	}))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), true))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source0 D11 = Companion_D11_.Create_DC30_()
	_ = _source0
	if _source0.Is_DC29() {
		var _1___mcc_h0 _dafny.Map = _source0.Get_().(D11_DC29).Cf43
		_ = _1___mcc_h0
		var _2_cf43 _dafny.Map = _1___mcc_h0
		_ = _2_cf43
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-222))).Uint32(), func(coer1 func(_dafny.Int) bool) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_3_i0 _dafny.Int) bool {
			return true
		})), _dafny.SeqOf(false, false, false))
	} else if _source0.Is_DC30() {
		return _dafny.SeqOf(!(false), !(!(!(false))), !(true), true)
	} else if _source0.Is_DC28() {
		var _4___mcc_h1 T1 = _source0.Get_().(D11_DC28).Cf42
		_ = _4___mcc_h1
		var _5_cf42 T1 = _4___mcc_h1
		_ = _5_cf42
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true, !(!(false)), true), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(593))).Uint32(), func(coer2 func(_dafny.Int) bool) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_6_i1 _dafny.Int) bool {
			return false
		})))
	} else {
		var _7___mcc_h2 D11 = _source0.Get_().(D11_DC31).Cf44
		_ = _7___mcc_h2
		var _8_cf44 D11 = _7___mcc_h2
		_ = _8_cf44
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(true))
	}
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((_dafny.Zero).Minus((((_dafny.SetOf(true, false)).Difference(_dafny.SetOf(false))).Intersection(_dafny.SetOf(!(!(false))))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-771))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	}))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	return ((_dafny.Zero).Minus(_dafny.IntOfInt64(-461))).Times(((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-343))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_10_i0 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('e'), true)).Cardinality()
	})))).Cardinality()).Plus(_dafny.IntOfInt64(518)))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("yvab"), true)).Cardinality()), _dafny.SeqOf((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(927), _dafny.IntOfInt64(999))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _11_v0 _dafny.Int
			_11_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(927)).Cmp(_11_v0) <= 0) && ((_11_v0).Cmp(_dafny.IntOfInt64(999)) < 0) {
				_coll0.Add((_11_v0).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ykwfkuyl")).Cardinality())), _dafny.IntOfInt64(298))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality(), _dafny.IntOfInt64(312), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality()), _dafny.IntOfInt64(792))), _dafny.SeqOf(_dafny.IntOfInt64(921)))
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_()))).Union(_dafny.MultiSetOf(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC1_(_dafny.CodePoint('y'))), Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_()), Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_())))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, globalState *GlobalState) D1 {
	var _source1 bool = false
	_ = _source1
	var _12___mcc_h0 bool = _source1
	_ = _12___mcc_h0
	var _13_cf0 bool = _12___mcc_h0
	_ = _13_cf0
	return Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_())
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) bool {
	return !((_dafny.IntOfUint32((_dafny.SeqOf(!(true), !(false))).Cardinality())).Cmp(_dafny.IntOfInt64(571)) != 0)
}
func (_static *CompanionStruct_Default___) Fm25(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.MultiSet, p1 bool, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC6_((_dafny.SetOf(_dafny.IntOfInt64(644))).Union(func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-907), _dafny.IntOfInt64(-218))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _14_v0 _dafny.Int
			_14_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(-907)).Cmp(_14_v0) <= 0) && ((_14_v0).Cmp(_dafny.IntOfInt64(-218)) < 0) {
				_coll1.Add((_14_v0).Times(_dafny.IntOfInt64(-978)))
			}
		}
		return _coll1.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.MultiSetOf(true, true, false))).Intersection(_dafny.SetOf(_dafny.MultiSetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return (Companion_D2_.Create_DC4_(_dafny.MultiSetFromSeq(_dafny.SeqOf(!(true))))).Dtor_cf3()
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 _dafny.MultiSet, globalState *GlobalState) _dafny.Set {
	var _source2 D10 = Companion_D10_.Create_DC25_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(529), ((Companion_D2_.Create_DC4_(_dafny.MultiSetOf(false))).Dtor_cf3()).Cardinality()))
	_ = _source2
	if _source2.Is_DC26() {
		var _15___mcc_h0 _dafny.Int = _source2.Get_().(D10_DC26).Cf40
		_ = _15___mcc_h0
		var _16_cf40 _dafny.Int = _15___mcc_h0
		_ = _16_cf40
		return (Companion_D3_.Create_DC6_(_dafny.SetOf((func() _dafny.Map {
			var _coll2 = _dafny.NewMapBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-914), _dafny.IntOfInt64(258))); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _17_v0 _dafny.Int
				_17_v0 = interface{}(_compr_2).(_dafny.Int)
				if ((_dafny.IntOfInt64(-914)).Cmp(_17_v0) <= 0) && ((_17_v0).Cmp(_dafny.IntOfInt64(258)) < 0) {
					_coll2.Add((_17_v0).Minus(_16_cf40), true)
				}
			}
			return _coll2.ToMap()
		}()).Cardinality()))).Dtor_cf9()
	} else if _source2.Is_DC25() {
		var _18___mcc_h1 _dafny.Map = _source2.Get_().(D10_DC25).Cf39
		_ = _18___mcc_h1
		var _19_cf39 _dafny.Map = _18___mcc_h1
		_ = _19_cf39
		return func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(574), _dafny.IntOfInt64(360))); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _20_v1 _dafny.Int
				_20_v1 = interface{}(_compr_3).(_dafny.Int)
				if ((_dafny.IntOfInt64(574)).Cmp(_20_v1) <= 0) && ((_20_v1).Cmp(_dafny.IntOfInt64(360)) < 0) {
					_coll3.Add(Companion_Default___.SafeDivisionInt(_20_v1, _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
						var _coll4 = _dafny.NewBuilder()
						_ = _coll4
						for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(772), _dafny.IntOfInt64(533))); ; {
							_compr_4, _ok4 := _iter4()
							if !_ok4 {
								break
							}
							var _21_v2 _dafny.Int
							_21_v2 = interface{}(_compr_4).(_dafny.Int)
							if ((_dafny.IntOfInt64(772)).Cmp(_21_v2) <= 0) && ((_21_v2).Cmp(_dafny.IntOfInt64(533)) < 0) {
								_coll4.Add((_21_v2).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(true, false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-308))).Cardinality()), _dafny.IntOfInt64(-914), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfInt64(50))).Cardinality())).Cardinality())))
							}
						}
						return _coll4.ToSet()
					}()).Cardinality())).Cardinality())))
				}
			}
			return _coll3.ToSet()
		}()
	} else {
		var _22___mcc_h2 D10 = _source2.Get_().(D10_DC27).Cf41
		_ = _22___mcc_h2
		var _23_cf41 D10 = _22___mcc_h2
		_ = _23_cf41
		return _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xlnmqckl")).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, globalState *GlobalState) _dafny.Sequence {
	if (_dafny.IntOfInt64(720)).Cmp((func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC6_(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rsyxbepi")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
			var _coll6 = _dafny.NewBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(853), _dafny.IntOfInt64(631))); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _24_v1 _dafny.Int
				_24_v1 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(853)).Cmp(_24_v1) <= 0) && ((_24_v1).Cmp(_dafny.IntOfInt64(631)) < 0) {
					_coll6.Add((_24_v1).Plus((_dafny.SetOf(false, false)).Cardinality()))
				}
			}
			return _coll6.ToSet()
		}()).Cardinality(), (_dafny.SetOf(true)).Cardinality())).Cardinality(), (_dafny.SetOf(true, false, !(false))).Cardinality())).Cardinality()))), _dafny.UnicodeSeqOfUtf8Bytes("veyog"))).Keys().Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _25_v0 D3
			_25_v0 = interface{}(_compr_5).(D3)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC6_(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rsyxbepi")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
				var _coll7 = _dafny.NewBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(853), _dafny.IntOfInt64(631))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _26_v1 _dafny.Int
					_26_v1 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(853)).Cmp(_26_v1) <= 0) && ((_26_v1).Cmp(_dafny.IntOfInt64(631)) < 0) {
						_coll7.Add((_26_v1).Plus((_dafny.SetOf(false, false)).Cardinality()))
					}
				}
				return _coll7.ToSet()
			}()).Cardinality(), (_dafny.SetOf(true)).Cardinality())).Cardinality(), (_dafny.SetOf(true, false, !(false))).Cardinality())).Cardinality()))), _dafny.UnicodeSeqOfUtf8Bytes("veyog"))).Contains(_25_v0) {
				_coll5.Add(_25_v0, _dafny.IntOfInt64(-649))
			}
		}
		return _coll5.ToMap()
	}()).Cardinality()) < 0 {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true)), _dafny.SeqOf(false))
	} else if !(!(true)) {
		return _dafny.SeqOf(true, true, true, false, true)
	} else {
		return _dafny.SeqOf(true, !(true))
	}
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, p1 _dafny.MultiSet, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!(!(false)), false)).Cardinality())), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(617))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(55), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-936))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_27_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	}))).Cardinality()))), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(424), (Companion_D8_.Create_DC22_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-954), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qejrjkfg")).Cardinality()))).Cardinality()))).Dtor_cf35())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(140))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_28_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(360)
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dhvkev")).Cardinality()))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(713), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(108))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_29_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm32(globalState *GlobalState) D5 {
	var _source3 D5 = Companion_D5_.Create_DC13_(false, _dafny.CodePoint('y'))
	_ = _source3
	if _source3.Is_DC12() {
		var _30___mcc_h0 bool = _source3.Get_().(D5_DC12).Cf19
		_ = _30___mcc_h0
		var _31_cf19 bool = _30___mcc_h0
		_ = _31_cf19
		return Companion_D5_.Create_DC13_(_31_cf19, _dafny.CodePoint('r'))
	} else if _source3.Is_DC13() {
		var _32___mcc_h1 bool = _source3.Get_().(D5_DC13).Cf20
		_ = _32___mcc_h1
		var _33___mcc_h2 _dafny.CodePoint = _source3.Get_().(D5_DC13).Cf21
		_ = _33___mcc_h2
		var _34_cf21 _dafny.CodePoint = _33___mcc_h2
		_ = _34_cf21
		var _35_cf20 bool = _32___mcc_h1
		_ = _35_cf20
		return Companion_D5_.Create_DC13_(_35_cf20, _34_cf21)
	} else if _source3.Is_DC11() {
		var _36___mcc_h3 _dafny.Map = _source3.Get_().(D5_DC11).Cf18
		_ = _36___mcc_h3
		var _37_cf18 _dafny.Map = _36___mcc_h3
		_ = _37_cf18
		return Companion_D5_.Create_DC13_(true, _dafny.CodePoint('f'))
	} else {
		var _38___mcc_h4 D5 = _source3.Get_().(D5_DC14).Cf22
		_ = _38___mcc_h4
		var _39_cf22 D5 = _38___mcc_h4
		_ = _39_cf22
		return Companion_D5_.Create_DC13_(false, _dafny.CodePoint('f'))
	}
}
func (_static *CompanionStruct_Default___) Fm33(globalState *GlobalState) bool {
	return ((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qqnloowd")).Cardinality())))).Cmp(_dafny.IntOfInt64(-548)) == 0
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, p1 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(Companion_D6_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(155))).Cardinality()), _dafny.SeqOf(true)), Companion_D6_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality())), _dafny.SeqOf(false)))).Intersection(func() _dafny.Set {
		var _coll8 = _dafny.NewBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D6_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), _dafny.IntOfInt64(699)), _dafny.SeqOf(false, true))))).Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _40_v0 D6
			_40_v0 = interface{}(_compr_8).(D6)
			if (_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D6_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), _dafny.IntOfInt64(699)), _dafny.SeqOf(false, true))))).Contains(_40_v0) {
				_coll8.Add(_40_v0)
			}
		}
		return _coll8.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC16_((func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.MultiSetOf(true), _dafny.MultiSetOf(true))).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _41_v0 _dafny.MultiSet
			_41_v0 = interface{}(_compr_9).(_dafny.MultiSet)
			if (_dafny.MultiSetOf(_dafny.MultiSetOf(true), _dafny.MultiSetOf(true))).Contains(_41_v0) {
				_coll9.Add(_41_v0, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
			}
		}
		return _coll9.ToMap()
	}()).Merge(func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(83))).Uint32(), func(coer8 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_42_i0 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(false, true)
			}))).Elements()); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _43_v2 _dafny.MultiSet
				_43_v2 = interface{}(_compr_11).(_dafny.MultiSet)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(83))).Uint32(), func(coer9 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg9 _dafny.Int) interface{} {
						return coer9(arg9)
					}
				}(func(_42_i0 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(false, true)
				})), _43_v2) {
					_coll11.Add(_43_v2, _dafny.IntOfInt64(775))
				}
			}
			return _coll11.ToMap()
		}()).Keys().Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _44_v1 _dafny.MultiSet
			_44_v1 = interface{}(_compr_10).(_dafny.MultiSet)
			if (func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(83))).Uint32(), func(coer10 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}(func(_42_i0 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(false, true)
				}))).Elements()); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _43_v2 _dafny.MultiSet
					_43_v2 = interface{}(_compr_12).(_dafny.MultiSet)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(83))).Uint32(), func(coer11 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
						return func(arg11 _dafny.Int) interface{} {
							return coer11(arg11)
						}
					}(func(_42_i0 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf(false, true)
					})), _43_v2) {
						_coll12.Add(_43_v2, _dafny.IntOfInt64(775))
					}
				}
				return _coll12.ToMap()
			}()).Contains(_44_v1) {
				_coll10.Add(_44_v1, _dafny.IntOfUint32((_dafny.SeqOf(true, true, false, false, !(false))).Cardinality()))
			}
		}
		return _coll10.ToMap()
	}()), _dafny.SeqOf(false))
}
func (_static *CompanionStruct_Default___) Fm36(p0 bool, p1 bool, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll13 = _dafny.NewBuilder()
		_ = _coll13
		for _iter13 := _dafny.Iterate((_dafny.SetOf(_dafny.MultiSetOf(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_()), Companion_D1_.Create_DC3_(Companion_D1_.Create_DC1_(_dafny.CodePoint('f')))))).Elements()); ; {
			_compr_13, _ok13 := _iter13()
			if !_ok13 {
				break
			}
			var _45_v0 _dafny.MultiSet
			_45_v0 = interface{}(_compr_13).(_dafny.MultiSet)
			if (_dafny.SetOf(_dafny.MultiSetOf(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_()), Companion_D1_.Create_DC3_(Companion_D1_.Create_DC1_(_dafny.CodePoint('f')))))).Contains(_45_v0) {
				_coll13.Add(_45_v0)
			}
		}
		return _coll13.ToSet()
	}()).Union((_dafny.SetOf(_dafny.MultiSetOf(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_())), Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_()), Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_()), Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_())), _dafny.MultiSetOf(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_()))))).Intersection(_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_()), Companion_D1_.Create_DC3_(Companion_D1_.Create_DC1_(_dafny.CodePoint('m'))), Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_()))))))
}
func (_static *CompanionStruct_Default___) Fm37(globalState *GlobalState) _dafny.Map {
	if !((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("npm")).Cardinality())).Cmp(_dafny.IntOfInt64(-818)) < 0) {
		return func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(93), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(694), _dafny.IntOfInt64(980))).Elements()); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _46_v0 _dafny.Int
				_46_v0 = interface{}(_compr_14).(_dafny.Int)
				if (_dafny.SetOf(_dafny.IntOfInt64(93), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(694), _dafny.IntOfInt64(980))).Contains(_46_v0) {
					_coll14.Add((_46_v0).Plus(_dafny.IntOfInt64(302)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(184))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg12 _dafny.Int) interface{} {
							return coer12(arg12)
						}
					}(func(_47_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('f')
					})))
				}
			}
			return _coll14.ToMap()
		}()
	} else {
		return (func() _dafny.Map {
			var _coll15 = _dafny.NewMapBuilder()
			_ = _coll15
			for _iter15 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(732), _dafny.IntOfInt64(894))).Elements()); ; {
				_compr_15, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _48_v1 _dafny.Int
				_48_v1 = interface{}(_compr_15).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(732), _dafny.IntOfInt64(894)), _48_v1) {
					_coll15.Add((_48_v1).Minus(_dafny.IntOfInt64(-626)), _dafny.UnicodeSeqOfUtf8Bytes("nf"))
				}
			}
			return _coll15.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-552))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_49_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		}))))
	}
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	if ((func() _dafny.Map {
		var _coll16 = _dafny.NewMapBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate((_dafny.SeqOf(_dafny.MultiSetOf(false), _dafny.MultiSetFromSeq(_dafny.SeqOf(false, false)))).Elements()); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _50_v0 _dafny.MultiSet
			_50_v0 = interface{}(_compr_16).(_dafny.MultiSet)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.MultiSetOf(false), _dafny.MultiSetFromSeq(_dafny.SeqOf(false, false))), _50_v0) {
				_coll16.Add(_50_v0, _dafny.IntOfInt64(-237))
			}
		}
		return _coll16.ToMap()
	}()).Cardinality()).Cmp(_dafny.IntOfInt64(47)) < 0 {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-275)), _dafny.IntOfInt64(295))).Merge(func() _dafny.Map {
			var _coll17 = _dafny.NewMapBuilder()
			_ = _coll17
			for _iter17 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(!(false))).Cardinality(), false)).Keys().Elements()); ; {
				_compr_17, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _51_v1 _dafny.Int
				_51_v1 = interface{}(_compr_17).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(!(false))).Cardinality(), false)).Contains(_51_v1) {
					_coll17.Add((_51_v1).Times((_dafny.Zero).Minus(_dafny.IntOfInt64(-744))), _dafny.IntOfInt64(-645))
				}
			}
			return _coll17.ToMap()
		}())
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			var _coll18 = _dafny.NewMapBuilder()
			_ = _coll18
			for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(887), _dafny.IntOfInt64(-826))); ; {
				_compr_18, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _52_v2 _dafny.Int
				_52_v2 = interface{}(_compr_18).(_dafny.Int)
				if ((_dafny.IntOfInt64(887)).Cmp(_52_v2) <= 0) && ((_52_v2).Cmp(_dafny.IntOfInt64(-826)) < 0) {
					_coll18.Add(Companion_Default___.SafeModuloInt(_52_v2, _dafny.IntOfInt64(-828)), true)
				}
			}
			return _coll18.ToMap()
		}()).Cardinality(), _dafny.IntOfInt64(28))
	}
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, p1 D3, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.MultiSetFromSeq(_dafny.SeqOf((func() _dafny.Map {
		var _coll19 = _dafny.NewMapBuilder()
		_ = _coll19
		for _iter19 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(500), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.MultiSetOf(!(true))).Cardinality()), true)).Cardinality())).Keys().Elements()); ; {
			_compr_19, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _53_v0 _dafny.Int
			_53_v0 = interface{}(_compr_19).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(500), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.MultiSetOf(!(true))).Cardinality()), true)).Cardinality())).Contains(_53_v0) {
				_coll19.Add(Companion_Default___.SafeModuloInt(_53_v0, _dafny.IntOfInt64(-296)), _dafny.SetOf(true))
			}
		}
		return _coll19.ToMap()
	}()).Cardinality()))).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(true, true)).Cardinality()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
		var _coll20 = _dafny.NewMapBuilder()
		_ = _coll20
		for _iter20 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(667))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_54_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(405)
		}))).Elements()); ; {
			_compr_20, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _55_v1 _dafny.Int
			_55_v1 = interface{}(_compr_20).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(667))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}(func(_54_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(405)
			})), _55_v1) {
				_coll20.Add((_55_v1).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qqnmhka")).Cardinality())), _dafny.IntOfInt64(965))
			}
		}
		return _coll20.ToMap()
	}()).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(true))), _dafny.IntOfInt64(224))))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) (_dafny.Array, bool) {
	var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r0
	var r1 bool = false
	_ = r1
	var _56_v0 _dafny.Sequence
	_ = _56_v0
	_56_v0 = _dafny.SeqOf(_dafny.IntOfInt64(-104))
	var _57_v1 _dafny.Map
	_ = _57_v1
	_57_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v0, p2)
	var _58_v2 D7
	_ = _58_v2
	_58_v2 = Companion_D7_.Create_DC18_(p2, (func() bool {
		if (_57_v1).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(748))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_59_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_60_i0 _dafny.Int) _dafny.Int {
				return _59_p0
			}
		})(p0)))) {
			return (_57_v1).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(748))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_61_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_62_i0 _dafny.Int) _dafny.Int {
					return _61_p0
				}
			})(p0)))).(bool)
		}
		return p1
	})())
	var _source4 D7 = _58_v2
	_ = _source4
	if _source4.Is_DC18() {
		var _63___mcc_h0 bool = _source4.Get_().(D7_DC18).Cf27
		_ = _63___mcc_h0
		var _64___mcc_h1 bool = _source4.Get_().(D7_DC18).Cf28
		_ = _64___mcc_h1
		var _65_cf28 bool = _64___mcc_h1
		_ = _65_cf28
		var _66_cf27 bool = _63___mcc_h0
		_ = _66_cf27
		var _67_v3 _dafny.Sequence
		_ = _67_v3
		_67_v3 = _dafny.SeqOf(!(true))
		var _68_v4 _dafny.Map
		_ = _68_v4
		_68_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D7_.Create_DC17_(_dafny.UnicodeSeqOfUtf8Bytes("mlycpc")), _67_v3)
		var _69_v5 _dafny.Map
		_ = _69_v5
		_69_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v4, (func() bool {
			if p2 {
				return true
			}
			return p1
		})())
		_69_v5 = (_69_v5).Update(_68_v4, _65_cf28)
		var _70_v6 _dafny.Map
		_ = _70_v6
		_70_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _66_cf27)
		var _71_v7 _dafny.Map
		_ = _71_v7
		_71_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v6, p0)
		var _72_v8 _dafny.Sequence
		_ = _72_v8
		_72_v8 = _dafny.UnicodeSeqOfUtf8Bytes("n")
		var _73_v9 *C5
		_ = _73_v9
		var _nw0 *C5 = New_C5_()
		_ = _nw0
		_nw0.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(911))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}(func(_74_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		})), (Companion_Default___.SafeIndex((_71_v7).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(911))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}(func(_75_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		}))).Cardinality()))).Uint32(), _dafny.CodePoint('n')), _72_v8), (_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), p0, p0)
		_73_v9 = _nw0
		var _76_v11 *C4
		_ = _76_v11
		var _nw1 *C4 = New_C4_()
		_ = _nw1
		_nw1.Ctor__(p1, (func() _dafny.Map {
			var _coll21 = _dafny.NewMapBuilder()
			_ = _coll21
			for _iter21 := _dafny.Iterate((_70_v6).Keys().Elements()); ; {
				_compr_21, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _77_v10 _dafny.Int
				_77_v10 = interface{}(_compr_21).(_dafny.Int)
				if (_70_v6).Contains(_77_v10) {
					_coll21.Add(Companion_Default___.SafeModuloInt(_77_v10, (_73_v9).F10()), (_73_v9).F10())
				}
			}
			return _coll21.ToMap()
		}()).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm21(p0, (_dafny.Zero).Minus(p0), p0, _dafny.UnicodeSeqOfUtf8Bytes("etffgh"), globalState)).Cardinality()))
		_76_v11 = _nw1
		var _78_v12 _dafny.MultiSet
		_ = _78_v12
		_78_v12 = _dafny.MultiSetOf(_76_v11)
		var _79_v13 _dafny.Sequence
		_ = _79_v13
		_79_v13 = _dafny.SeqOf(_76_v11)
		var _80_v14 _dafny.MultiSet
		_ = _80_v14
		_80_v14 = _dafny.MultiSetOf(_78_v12, _78_v12, _dafny.MultiSetOf(_76_v11, _76_v11), _dafny.MultiSetFromSeq(_79_v13))
		(globalState).F3 = (_dafny.MultiSetOf(_78_v12, _dafny.MultiSetOf(_76_v11))).IsSubsetOf(_80_v14)
		var _81_v15 _dafny.Int
		_ = _81_v15
		_81_v15 = _dafny.IntOfInt64(-986)
		_81_v15 = (_dafny.SetOf((_73_v9).F10(), _dafny.IntOfInt64(877), p0, _81_v15)).Cardinality()
	} else if _source4.Is_DC19() {
		var _82___mcc_h2 _dafny.Int = _source4.Get_().(D7_DC19).Cf29
		_ = _82___mcc_h2
		var _83___mcc_h3 bool = _source4.Get_().(D7_DC19).Cf30
		_ = _83___mcc_h3
		var _84___mcc_h4 _dafny.Int = _source4.Get_().(D7_DC19).Cf31
		_ = _84___mcc_h4
		var _85_cf31 _dafny.Int = _84___mcc_h4
		_ = _85_cf31
		var _86_cf30 bool = _83___mcc_h3
		_ = _86_cf30
		var _87_cf29 _dafny.Int = _82___mcc_h2
		_ = _87_cf29
		var _88_v16 *C3
		_ = _88_v16
		var _nw2 *C3 = New_C3_()
		_ = _nw2
		_nw2.Ctor__(p2, Companion_Default___.SafeModuloInt(Companion_Default___.Fm20(p2, p1, !(p1), globalState), _85_cf31), _dafny.IntOfInt64(-709))
		_88_v16 = _nw2
		var _89_v17 _dafny.Sequence
		_ = _89_v17
		_89_v17 = _dafny.UnicodeSeqOfUtf8Bytes("rtmdkqwc")
		var _90_v18 _dafny.Set
		_ = _90_v18
		_90_v18 = _dafny.SetOf((_88_v16).Fm13(_89_v17, globalState), p0)
		var _91_v19 _dafny.Map
		_ = _91_v19
		_91_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_90_v18).Intersection(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(p1, !(!(!(p1))))).Cardinality()), p0)), (func() bool {
			if (_88_v16).F12() {
				return true
			}
			return _86_cf30
		})())
		_91_v19 = (_91_v19).Update(_90_v18, !(_86_cf30) || ((_88_v16).F12()))
		_85_cf31 = _85_cf31
		_86_cf30 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ruevlxjng"), _dafny.Companion_Sequence_.Concatenate(_89_v17, _89_v17))
	} else if _source4.Is_DC20() {
		var _92___mcc_h5 _dafny.CodePoint = _source4.Get_().(D7_DC20).Cf32
		_ = _92___mcc_h5
		var _93___mcc_h6 _dafny.Int = _source4.Get_().(D7_DC20).Cf33
		_ = _93___mcc_h6
		var _94_cf33 _dafny.Int = _93___mcc_h6
		_ = _94_cf33
		var _95_cf32 _dafny.CodePoint = _92___mcc_h5
		_ = _95_cf32
		(globalState).F4 = p1
		var _96_v20 _dafny.Array
		_ = _96_v20
		var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
		_ = _nw3
		_96_v20 = _nw3
		var _97_v21 _dafny.Array
		_ = _97_v21
		var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
		_ = _nw4
		_97_v21 = _nw4
		var _98_v22 _dafny.Sequence
		_ = _98_v22
		_98_v22 = _dafny.UnicodeSeqOfUtf8Bytes("xjka")
		var _99_v23 _dafny.Map
		_ = _99_v23
		_99_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_cf33, _98_v22)
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_97_v21), 0))
		_ = _index0
		(_97_v21).ArraySet1(_99_v23, (_index0).Int())
		var _100_v24 D9
		_ = _100_v24
		_100_v24 = Companion_D9_.Create_DC23_(_96_v20)
		var _101_v25 _dafny.Map
		_ = _101_v25
		_101_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		var _102_v26 _dafny.Sequence
		_ = _102_v26
		_102_v26 = _dafny.SeqOf(p1, !(p1) || ((func() bool {
			if (_101_v25).Contains(false) {
				return (_101_v25).Get(false).(bool)
			}
			return p1
		})()))
		var _103_v27 _dafny.MultiSet
		_ = _103_v27
		_103_v27 = _dafny.MultiSetOf(_dafny.IntOfInt64(-877))
		var _104_v28 _dafny.Set
		_ = _104_v28
		_104_v28 = _dafny.SetOf(p0)
		var _105_v29 *C5
		_ = _105_v29
		var _nw5 *C5 = New_C5_()
		_ = _nw5
		_nw5.Ctor__(_98_v22, _94_cf33, (_103_v27).Cardinality(), (_104_v28).Cardinality())
		_105_v29 = _nw5
		var _106_v30 _dafny.Map
		_ = _106_v30
		_106_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v29, p2)
		var _107_v31 _dafny.Map
		_ = _107_v31
		_107_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_105_v29).F10(), p0), (Companion_Default___.Fm37(globalState)).Merge(Companion_Default___.Fm37(globalState)))
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_97_v21), 0))
		_ = _index1
		var _rhs0 bool = p1
		_ = _rhs0
		var _rhs1 _dafny.Array = (_100_v24).Dtor_cf36()
		_ = _rhs1
		var _rhs2 bool = (_102_v26).Select((Companion_Default___.SafeIndex(((_106_v30).Cardinality()).Times((_105_v29).F10()), _dafny.IntOfUint32((_102_v26).Cardinality()))).Uint32()).(bool)
		_ = _rhs2
		var _rhs3 _dafny.Map = (func() _dafny.Map {
			if (_107_v31).Contains((_dafny.Zero).Minus((_105_v29).F10())) {
				return (_107_v31).Get((_dafny.Zero).Minus((_105_v29).F10())).(_dafny.Map)
			}
			return (_99_v23).Merge(_99_v23)
		})()
		_ = _rhs3
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		var _lhs2 _dafny.Array = _97_v21
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_97_v21), 0))
		_ = _lhs3
		_lhs0.F4 = _rhs0
		_96_v20 = _rhs1
		_lhs1.F4 = _rhs2
		(_lhs2).ArraySet1(_rhs3, (_lhs3).Int())
		var _108_v32 _dafny.Array
		_ = _108_v32
		var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
		_ = _nw6
		_108_v32 = _nw6
		var _109_v33 _dafny.Map
		_ = _109_v33
		_109_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v32, _dafny.IntOfInt64(-969))
		_109_v33 = (_109_v33).Update(_108_v32, _94_cf33)
		_94_cf33 = Companion_Default___.SafeModuloInt((_94_cf33).Minus(_94_cf33), p0)
	} else {
		var _110___mcc_h7 _dafny.Sequence = _source4.Get_().(D7_DC17).Cf26
		_ = _110___mcc_h7
		var _111_cf26 _dafny.Sequence = _110___mcc_h7
		_ = _111_cf26
		var _112_v34 _dafny.Sequence
		_ = _112_v34
		_112_v34 = _dafny.SeqOf(!(true))
		var _113_v35 _dafny.Array
		_ = _113_v35
		var _nwElement0_0 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm30(false, globalState)).Cardinality())
		_ = _nwElement0_0
		var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(18))
		_ = _nw7
		(_nw7).ArraySet1(_nwElement0_0, 0)
		(_nw7).ArraySet1(_dafny.IntOfInt64(122), 1)
		(_nw7).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), 2)
		(_nw7).ArraySet1(Companion_Default___.Fm20(p1, p2, p1, globalState), 3)
		(_nw7).ArraySet1(p0, 4)
		(_nw7).ArraySet1(p0, 5)
		(_nw7).ArraySet1(p0, 6)
		(_nw7).ArraySet1(p0, 7)
		(_nw7).ArraySet1(p0, 8)
		(_nw7).ArraySet1(p0, 9)
		(_nw7).ArraySet1(_dafny.IntOfInt64(674), 10)
		(_nw7).ArraySet1(p0, 11)
		(_nw7).ArraySet1(p0, 12)
		(_nw7).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_112_v34).Cardinality())), 13)
		(_nw7).ArraySet1((_dafny.Zero).Minus(p0), 14)
		(_nw7).ArraySet1(p0, 15)
		(_nw7).ArraySet1(p0, 16)
		(_nw7).ArraySet1(_dafny.IntOfInt64(885), 17)
		_113_v35 = _nw7
		var _114_v36 _dafny.Sequence
		_ = _114_v36
		_114_v36 = _dafny.SeqOf(_113_v35, _113_v35)
		var _115_v37 _dafny.Map
		_ = _115_v37
		_115_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_114_v36).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_114_v36).Cardinality()))).Uint32()).(_dafny.Array))
		var _116_v38 _dafny.Map
		_ = _116_v38
		_116_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), p1)
		_115_v37 = (_115_v37).Update((func() bool {
			if (_116_v38).Contains(p0) {
				return (_116_v38).Get(p0).(bool)
			}
			return p1
		})(), _113_v35)
		(globalState).F4 = (func() bool {
			if p2 {
				return (func() bool {
					if p2 {
						return false
					}
					return p1
				})()
			}
			return true
		})()
		if (p0).Cmp(_dafny.IntOfInt64(240)) > 0 {
			var _117_v39 *C4
			_ = _117_v39
			var _nw8 *C4 = New_C4_()
			_ = _nw8
			_nw8.Ctor__(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm30(!(!(false)), globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm30(!(!(false)), globalState)).Cardinality()))).Uint32(), p1)).Cardinality()), (_dafny.Zero).Minus(p0))
			_117_v39 = _nw8
			var _118_v40 T0
			_ = _118_v40
			var _nw9 *C1 = New_C1_()
			_ = _nw9
			_nw9.Ctor__(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mhnx")).Cardinality()))
			_118_v40 = _nw9
			var _119_v41 _dafny.Sequence
			_ = _119_v41
			_119_v41 = _dafny.SeqOf(_118_v40, _118_v40)
			var _rhs4 *C4 = _117_v39
			_ = _rhs4
			var _rhs5 _dafny.Sequence = _119_v41
			_ = _rhs5
			_117_v39 = _rhs4
			_119_v41 = _rhs5
			var _120_v42 _dafny.Int
			_ = _120_v42
			_120_v42 = _dafny.Zero
			_120_v42 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_118_v40).F8(), (_118_v40).F7()), (_dafny.Zero).Minus((_dafny.Zero).Minus(_120_v42)))
			_56_v0 = (func() _dafny.Sequence {
				if p1 {
					return _56_v0
				}
				return _56_v0
			})()
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_113_v35), 0))
			_ = _index2
			(_113_v35).ArraySet1((_dafny.Zero).Minus(p0), (_index2).Int())
			var _121_v43 _dafny.Map
			_ = _121_v43
			_121_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_113_v35), 0))
			_ = _index3
			(_113_v35).ArraySet1((_120_v42).Times((func() _dafny.Int {
				if (_121_v43).Contains(Companion_Default___.Fm24((_117_v39).F11(), (_118_v40).F7(), _111_cf26, globalState)) {
					return (_121_v43).Get(Companion_Default___.Fm24((_117_v39).F11(), (_118_v40).F7(), _111_cf26, globalState)).(_dafny.Int)
				}
				return _120_v42
			})()), (_index3).Int())
			var _122_v44 _dafny.MultiSet
			_ = _122_v44
			_122_v44 = _dafny.MultiSetOf((Companion_Default___.Fm28(p2, globalState)).Cardinality(), _120_v42, _dafny.IntOfUint32((_111_cf26).Cardinality()), _120_v42)
			var _123_v45 _dafny.Array
			_ = _123_v45
			var _nwElement0_1 bool = (_117_v39).F11()
			_ = _nwElement0_1
			var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(9))
			_ = _nw10
			(_nw10).ArraySet1(_nwElement0_1, 0)
			(_nw10).ArraySet1(p2, 1)
			(_nw10).ArraySet1((_120_v42).Cmp((_118_v40).F7()) <= 0, 2)
			(_nw10).ArraySet1(p2, 3)
			(_nw10).ArraySet1(!(p1) || (p1), 4)
			(_nw10).ArraySet1(p2, 5)
			(_nw10).ArraySet1(true, 6)
			(_nw10).ArraySet1((_122_v44).Equals(_122_v44), 7)
			(_nw10).ArraySet1((_117_v39).F11(), 8)
			_123_v45 = _nw10
			var _124_v46 D10
			_ = _124_v46
			_124_v46 = Companion_D10_.Create_DC25_(Companion_Default___.Fm38((_117_v39).Fm1(_dafny.SetOf(p2), p0, (_dafny.Zero).Minus(_120_v42), _dafny.IntOfInt64(554), globalState), true, Companion_Default___.Fm24(p2, _dafny.IntOfInt64(816), _111_cf26, globalState), (_118_v40).F8(), globalState))
			var _125_v47 _dafny.Set
			_ = _125_v47
			_125_v47 = _dafny.SetOf(((_124_v46).Dtor_cf39()).Cardinality())
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_123_v45), 0))
			_ = _index4
			(_123_v45).ArraySet1((_125_v47).IsSubsetOf(func() _dafny.Set {
				var _coll22 = _dafny.NewBuilder()
				_ = _coll22
				for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(107), _dafny.IntOfInt64(642))); ; {
					_compr_22, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _126_v48 _dafny.Int
					_126_v48 = interface{}(_compr_22).(_dafny.Int)
					if ((_dafny.IntOfInt64(107)).Cmp(_126_v48) <= 0) && ((_126_v48).Cmp(_dafny.IntOfInt64(642)) < 0) {
						_coll22.Add((_126_v48).Times((_118_v40).F8()))
					}
				}
				return _coll22.ToSet()
			}()), (_index4).Int())
			var _127_v49 _dafny.Array
			_ = _127_v49
			var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(29))
			_ = _nw11
			_127_v49 = _nw11
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_127_v49), 0))
			_ = _index5
			(_127_v49).ArraySet1CodePoint(_dafny.CodePoint('e'), (_index5).Int())
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_123_v45), 0))
			_ = _index6
			(_123_v45).ArraySet1(p1, (_index6).Int())
			var _128_v50 _dafny.Array
			_ = _128_v50
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_0
			var _nw12 _dafny.Array
			_ = _nw12
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw12 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Map = (func(_129_p1 bool, _130_v40 T0) func(_dafny.Int) _dafny.Map {
					return func(_131_i2 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_p1, (_130_v40).F8())
					}
				})(p1, _118_v40)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw12 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw12).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw12).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_128_v50 = _nw12
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_128_v50), 0))
			_ = _index7
			(_128_v50).ArraySet1((_121_v43).Merge(_121_v43), (_index7).Int())
			var _132_v51 _dafny.CodePoint
			_ = _132_v51
			_132_v51 = _dafny.CodePoint('v')
			var _133_v52 D3
			_ = _133_v52
			_133_v52 = Companion_D3_.Create_DC6_(_125_v47)
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_123_v45), 0))
			_ = _index8
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_127_v49), 0))
			_ = _index9
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_123_v45), 0))
			_ = _index10
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_128_v50), 0))
			_ = _index11
			var _rhs6 bool = (_112_v34).Select((Companion_Default___.SafeIndex((_118_v40).F8(), _dafny.IntOfUint32((_112_v34).Cardinality()))).Uint32()).(bool)
			_ = _rhs6
			var _rhs7 _dafny.Int = _120_v42
			_ = _rhs7
			var _rhs8 _dafny.CodePoint = _132_v51
			_ = _rhs8
			var _rhs9 bool = p2
			_ = _rhs9
			var _rhs10 _dafny.Map = Companion_Default___.Fm39((_118_v40).F8(), _133_v52, p2, (_117_v39).F11(), globalState)
			_ = _rhs10
			var _lhs4 _dafny.Array = _123_v45
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_123_v45), 0))
			_ = _lhs5
			var _lhs6 _dafny.Array = _127_v49
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_127_v49), 0))
			_ = _lhs7
			var _lhs8 _dafny.Array = _123_v45
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_123_v45), 0))
			_ = _lhs9
			var _lhs10 _dafny.Array = _128_v50
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_128_v50), 0))
			_ = _lhs11
			(_lhs4).ArraySet1(_rhs6, (_lhs5).Int())
			_120_v42 = _rhs7
			(_lhs6).ArraySet1CodePoint(_rhs8, (_lhs7).Int())
			(_lhs8).ArraySet1(_rhs9, (_lhs9).Int())
			(_lhs10).ArraySet1(_rhs10, (_lhs11).Int())
		} else {
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_113_v35), 0))
			_ = _index12
			(_113_v35).ArraySet1(p0, (_index12).Int())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_113_v35), 0))
			_ = _index13
			(_113_v35).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_56_v0).Cardinality()), p0), (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_113_v35), 0))
			_ = _index14
			(_113_v35).ArraySet1((_113_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_113_v35), 0))).Int()).(_dafny.Int), (_index14).Int())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_113_v35), 0))
			_ = _index15
			(_113_v35).ArraySet1(p0, (_index15).Int())
			var _134_v53 _dafny.Map
			_ = _134_v53
			_134_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			var _135_v54 _dafny.CodePoint
			_ = _135_v54
			_135_v54 = _dafny.CodePoint('c')
			var _136_v55 _dafny.Array
			_ = _136_v55
			var _nwElement0_2 bool = p2
			_ = _nwElement0_2
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(23))
			_ = _nw13
			(_nw13).ArraySet1(_nwElement0_2, 0)
			(_nw13).ArraySet1(p2, 1)
			(_nw13).ArraySet1(p2, 2)
			(_nw13).ArraySet1(!(Companion_Default___.Fm24(p2, (_134_v53).Cardinality(), _dafny.SeqOf(_135_v54), globalState)), 3)
			(_nw13).ArraySet1(!(p1), 4)
			(_nw13).ArraySet1(p2, 5)
			(_nw13).ArraySet1(p1, 6)
			(_nw13).ArraySet1(p2, 7)
			(_nw13).ArraySet1(p2, 8)
			(_nw13).ArraySet1(p2, 9)
			(_nw13).ArraySet1(p1, 10)
			(_nw13).ArraySet1(p2, 11)
			(_nw13).ArraySet1(p2, 12)
			(_nw13).ArraySet1(p2, 13)
			(_nw13).ArraySet1(p1, 14)
			(_nw13).ArraySet1(p1, 15)
			(_nw13).ArraySet1(true, 16)
			(_nw13).ArraySet1(p1, 17)
			(_nw13).ArraySet1(p1, 18)
			(_nw13).ArraySet1(p1, 19)
			(_nw13).ArraySet1(p2, 20)
			(_nw13).ArraySet1(p1, 21)
			(_nw13).ArraySet1(true, 22)
			_136_v55 = _nw13
			var _137_v56 _dafny.Sequence
			_ = _137_v56
			_137_v56 = _dafny.SeqOf(_136_v55)
			var _138_v57 _dafny.MultiSet
			_ = _138_v57
			_138_v57 = _dafny.MultiSetOf((_113_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_113_v35), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_111_cf26).Cardinality()))
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_113_v35), 0))
			_ = _index16
			(_113_v35).ArraySet1(((Companion_D4_.Create_DC9_(p0, _137_v56, (_113_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_113_v35), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
				if (_138_v57).Contains((func() _dafny.Int {
					if (_138_v57).Contains(_dafny.IntOfInt64(516)) {
						return (_138_v57).Multiplicity(_dafny.IntOfInt64(516))
					}
					return (_113_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_113_v35), 0))).Int()).(_dafny.Int)
				})()) {
					return (_138_v57).Multiplicity((func() _dafny.Int {
						if (_138_v57).Contains(_dafny.IntOfInt64(516)) {
							return (_138_v57).Multiplicity(_dafny.IntOfInt64(516))
						}
						return (_113_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_113_v35), 0))).Int()).(_dafny.Int)
					})())
				}
				return _dafny.IntOfInt64(584)
			})())).Dtor_cf15()).Plus((_dafny.Zero).Minus((_113_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_113_v35), 0))).Int()).(_dafny.Int))), (_index16).Int())
			var _139_v58 _dafny.Array
			_ = _139_v58
			var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
			_ = _nw14
			_139_v58 = _nw14
			var _140_v59 _dafny.Array
			_ = _140_v59
			var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw15
			_140_v59 = _nw15
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_139_v58), 0))
			_ = _index17
			(_139_v58).ArraySet1(_140_v59, (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_139_v58), 0))
			_ = _index18
			(_139_v58).ArraySet1(_140_v59, (_index18).Int())
		}
		var _141_v60 *C3
		_ = _141_v60
		var _nw16 *C3 = New_C3_()
		_ = _nw16
		_nw16.Ctor__(p1, p0, p0)
		_141_v60 = _nw16
		var _142_v61 D8
		_ = _142_v61
		_142_v61 = Companion_D8_.Create_DC21_(_141_v60)
		var _143_v62 _dafny.Map
		_ = _143_v62
		_143_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v61, p2)
		_143_v62 = (_143_v62).Update(Companion_D8_.Create_DC21_(_141_v60), false)
	}
	var _144_v63 _dafny.CodePoint
	_ = _144_v63
	_144_v63 = _dafny.CodePoint('s')
	var _145_v64 _dafny.Array
	_ = _145_v64
	var _nwElement0_3 bool = !(!(p2))
	_ = _nwElement0_3
	var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(9))
	_ = _nw17
	(_nw17).ArraySet1(_nwElement0_3, 0)
	(_nw17).ArraySet1(p2, 1)
	(_nw17).ArraySet1(p1, 2)
	(_nw17).ArraySet1(p2, 3)
	(_nw17).ArraySet1(p1, 4)
	(_nw17).ArraySet1(true, 5)
	(_nw17).ArraySet1(p1, 6)
	(_nw17).ArraySet1(false, 7)
	(_nw17).ArraySet1(p1, 8)
	_145_v64 = _nw17
	var _146_v65 D6
	_ = _146_v65
	_146_v65 = Companion_D6_.Create_DC15_(_145_v64)
	var _147_v66 _dafny.Sequence
	_ = _147_v66
	_147_v66 = _dafny.UnicodeSeqOfUtf8Bytes("nbthqcwcx")
	var _148_v67 _dafny.Map
	_ = _148_v67
	_148_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_v65, _147_v66)
	var _149_i3 _dafny.Int
	_ = _149_i3
	_149_i3 = _dafny.Zero
	{
		for _dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
			if (_148_v67).Contains(_146_v65) {
				return (_148_v67).Get(_146_v65).(_dafny.Sequence)
			}
			return _147_v66
		})(), _144_v63) {
			{
				if (_149_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_149_i3 = (_149_i3).Plus(_dafny.One)
				var _150_v68 T0
				_ = _150_v68
				var _nw18 *C5 = New_C5_()
				_ = _nw18
				_nw18.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("e"), p0, p0, p0)
				_150_v68 = _nw18
				var _151_v69 T1
				_ = _151_v69
				var _nw19 *C5 = New_C5_()
				_ = _nw19
				_nw19.Ctor__(_147_v66, (_150_v68).F8(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ks")).Cardinality()), (_150_v68).F7())
				_151_v69 = _nw19
				var _152_v70 D11
				_ = _152_v70
				_152_v70 = Companion_D11_.Create_DC28_(_151_v69)
				var _153_v71 _dafny.Array
				_ = _153_v71
				var _nwElement0_4 T1 = _151_v69
				_ = _nwElement0_4
				var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(15))
				_ = _nw20
				(_nw20).ArraySet1(_nwElement0_4, 0)
				(_nw20).ArraySet1(_151_v69, 1)
				(_nw20).ArraySet1(_151_v69, 2)
				(_nw20).ArraySet1((_152_v70).Dtor_cf42(), 3)
				(_nw20).ArraySet1(_151_v69, 4)
				(_nw20).ArraySet1(_151_v69, 5)
				(_nw20).ArraySet1(_151_v69, 6)
				(_nw20).ArraySet1(_151_v69, 7)
				(_nw20).ArraySet1(_151_v69, 8)
				(_nw20).ArraySet1(_151_v69, 9)
				(_nw20).ArraySet1(_151_v69, 10)
				(_nw20).ArraySet1(_151_v69, 11)
				(_nw20).ArraySet1(_151_v69, 12)
				(_nw20).ArraySet1(_151_v69, 13)
				(_nw20).ArraySet1(_151_v69, 14)
				_153_v71 = _nw20
				var _154_v72 _dafny.Map
				_ = _154_v72
				_154_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_v68, _153_v71)
				var _155_v73 D12
				_ = _155_v73
				_155_v73 = Companion_D12_.Create_DC32_(_154_v72)
				var _156_v74 _dafny.Sequence
				_ = _156_v74
				_156_v74 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_v68, _153_v71))
				var _157_v75 _dafny.Map
				_ = _157_v75
				_157_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_150_v68).F8(), _154_v72)
				var _158_v76 _dafny.Array
				_ = _158_v76
				var _nwElement0_5 _dafny.Map = _154_v72
				_ = _nwElement0_5
				var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(21))
				_ = _nw21
				(_nw21).ArraySet1(_nwElement0_5, 0)
				(_nw21).ArraySet1(((_155_v73).Dtor_cf45()).Merge(_154_v72), 1)
				(_nw21).ArraySet1((_156_v74).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.IntOfUint32((_156_v74).Cardinality()))).Uint32()).(_dafny.Map), 2)
				(_nw21).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_v68, _153_v71), 3)
				(_nw21).ArraySet1(_154_v72, 4)
				(_nw21).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_v68, _153_v71), 5)
				(_nw21).ArraySet1(_154_v72, 6)
				(_nw21).ArraySet1(_154_v72, 7)
				(_nw21).ArraySet1(_154_v72, 8)
				(_nw21).ArraySet1((_154_v72).Update(_150_v68, _153_v71), 9)
				(_nw21).ArraySet1(_154_v72, 10)
				(_nw21).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_v68, _153_v71), 11)
				(_nw21).ArraySet1(_154_v72, 12)
				(_nw21).ArraySet1(_154_v72, 13)
				(_nw21).ArraySet1((_154_v72).Merge(_154_v72), 14)
				(_nw21).ArraySet1(_154_v72, 15)
				(_nw21).ArraySet1(_154_v72, 16)
				(_nw21).ArraySet1(_154_v72, 17)
				(_nw21).ArraySet1((_154_v72).Update(_150_v68, _153_v71), 18)
				(_nw21).ArraySet1(((_154_v72).Update(_150_v68, _153_v71)).Update(_150_v68, _153_v71), 19)
				(_nw21).ArraySet1((func() _dafny.Map {
					if (_157_v75).Contains(Companion_Default___.Fm16(p2, globalState)) {
						return (_157_v75).Get(Companion_Default___.Fm16(p2, globalState)).(_dafny.Map)
					}
					return _154_v72
				})(), 20)
				_158_v76 = _nw21
				var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_158_v76), 0))
				_ = _index19
				(_158_v76).ArraySet1(_154_v72, (_index19).Int())
				var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_158_v76), 0))
				_ = _index20
				(_158_v76).ArraySet1((_154_v72).Merge((_154_v72).Merge(_154_v72)), (_index20).Int())
				var _159_v77 _dafny.Map
				_ = _159_v77
				_159_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _151_v69)
				_159_v77 = (_159_v77).Update(Companion_Default___.Fm24(p2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_147_v66, (Companion_Default___.SafeIndex((_150_v68).F7(), _dafny.IntOfUint32((_147_v66).Cardinality()))).Uint32(), _144_v63)).Cardinality()), _147_v66, globalState), _151_v69)
				var _nw22 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw22
				_145_v64 = _nw22
				var _160_v78 _dafny.Map
				_ = _160_v78
				_160_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_v68, _151_v69)
				_160_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_150_v68, _151_v69)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))
	_ = _index21
	(_145_v64).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_147_v66, _147_v66), (_index21).Int())
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))
	_ = _index22
	(_145_v64).ArraySet1((p0).Cmp(p0) >= 0, (_index22).Int())
	var _161_v79 _dafny.Array
	_ = _161_v79
	var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
	_ = _nw23
	_161_v79 = _nw23
	var _162_v80 _dafny.Array
	_ = _162_v80
	var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
	_ = _nw24
	_162_v80 = _nw24
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))
	_ = _index23
	(_161_v79).ArraySet1(_162_v80, (_index23).Int())
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))
	_ = _index24
	(_161_v79).ArraySet1(_162_v80, (_index24).Int())
	if (_145_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))).Int()).(bool) {
		var _163_v81 D7
		_ = _163_v81
		_163_v81 = Companion_D7_.Create_DC19_(p0, (_145_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))).Int()).(bool), p0)
		var _164_v82 _dafny.Sequence
		_ = _164_v82
		_164_v82 = _dafny.SeqOf(_163_v81)
		var _165_v83 _dafny.MultiSet
		_ = _165_v83
		_165_v83 = _dafny.MultiSetOf(p0, p0, _dafny.IntOfUint32((_164_v82).Cardinality()), p0, _dafny.IntOfUint32((_147_v66).Cardinality()))
		_165_v83 = _165_v83
		var _arr0 _dafny.Array = _dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))
		_ = _arr0
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))
		_ = _index25
		(_arr0).ArraySet1(_dafny.IntOfInt64(237), (_index25).Int())
		var _166_v84 _dafny.MultiSet
		_ = _166_v84
		_166_v84 = _dafny.MultiSetOf(p2, (_145_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))).Int()).(bool))
		var _167_v85 *C1
		_ = _167_v85
		var _nw25 *C1 = New_C1_()
		_ = _nw25
		_nw25.Ctor__(p0, p0)
		_167_v85 = _nw25
		var _168_v86 _dafny.Map
		_ = _168_v86
		_168_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v85, p1)
		var _169_v87 *C1
		_ = _169_v87
		_169_v87 = _167_v85
		var _170_v88 D9
		_ = _170_v88
		_170_v88 = Companion_D9_.Create_DC23_(_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int())))
		var _171_v89 _dafny.Map
		_ = _171_v89
		_171_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_170_v88).Dtor_cf36())
		var _172_v90 _dafny.Array
		_ = _172_v90
		var _nw26 _dafny.Array = _dafny.NewArrayWithValue(Companion_D9_.Default(), _dafny.IntOfInt64(18))
		_ = _nw26
		_172_v90 = _nw26
		var _173_v91 _dafny.Map
		_ = _173_v91
		_173_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _172_v90)
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))
		_ = _index26
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))
		_ = _index27
		var _arr1 _dafny.Array = _dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))
		_ = _arr1
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))
		_ = _index28
		var _rhs11 bool = (_166_v84).IsDisjointFrom(_166_v84)
		_ = _rhs11
		var _rhs12 bool = (func() bool {
			if (_168_v86).Contains((_169_v87)) {
				return (_168_v86).Get((_169_v87)).(bool)
			}
			return true
		})()
		_ = _rhs12
		var _rhs13 _dafny.Array = (func() _dafny.Array {
			if (_171_v89).Contains(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _172_v90)).Merge(_173_v91)).Cardinality()) {
				return (_171_v89).Get(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _172_v90)).Merge(_173_v91)).Cardinality()).(_dafny.Array)
			}
			return (func() _dafny.Array {
				if p1 {
					return _dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))
				}
				return _162_v80
			})()
		})()
		_ = _rhs13
		var _rhs14 bool = false
		_ = _rhs14
		var _rhs15 _dafny.Int = (p0).Minus(_dafny.IntOfInt64(-541))
		_ = _rhs15
		var _lhs12 _dafny.Array = _145_v64
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))
		_ = _lhs13
		var _lhs14 *GlobalState = globalState
		_ = _lhs14
		var _lhs15 _dafny.Array = _161_v79
		_ = _lhs15
		var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))
		_ = _lhs16
		var _lhs17 *GlobalState = globalState
		_ = _lhs17
		var _lhs18 _dafny.Array = _dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))
		_ = _lhs18
		var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))
		_ = _lhs19
		(_lhs12).ArraySet1(_rhs11, (_lhs13).Int())
		_lhs14.F3 = _rhs12
		(_lhs15).ArraySet1(_rhs13, (_lhs16).Int())
		_lhs17.F4 = _rhs14
		(_lhs18).ArraySet1(_rhs15, (_lhs19).Int())
		(globalState).F4 = true
		var _174_v92 _dafny.Map
		_ = _174_v92
		_174_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_56_v0, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-735), _dafny.IntOfUint32((_56_v0).Cardinality()))).Uint32(), p0), _dafny.SeqOf((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))).Int()).(_dafny.Int))), p0)
		_174_v92 = ((_174_v92).Merge(_174_v92)).Merge(_174_v92)
		var _175_v93 _dafny.Array
		_ = _175_v93
		var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
		_ = _nw27
		_175_v93 = _nw27
		_175_v93 = (func() _dafny.Array {
			if true {
				return _175_v93
			}
			return _175_v93
		})()
	} else {
		var _176_v94 _dafny.Int
		_ = _176_v94
		_176_v94 = _dafny.IntOfInt64(-106)
		var _177_v95 _dafny.MultiSet
		_ = _177_v95
		_177_v95 = _dafny.MultiSetOf(p2)
		var _rhs16 bool = false
		_ = _rhs16
		var _rhs17 _dafny.Int = _176_v94
		_ = _rhs17
		var _rhs18 bool = ((_dafny.Zero).Minus(_176_v94)).Cmp(p0) == 0
		_ = _rhs18
		var _rhs19 bool = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(265), (func() _dafny.Int {
			if (_177_v95).Contains(p2) {
				return (_177_v95).Multiplicity(p2)
			}
			return (_dafny.Zero).Minus(_176_v94)
		})())).Cmp(p0) >= 0
		_ = _rhs19
		var _lhs20 *GlobalState = globalState
		_ = _lhs20
		var _lhs21 *GlobalState = globalState
		_ = _lhs21
		var _lhs22 *GlobalState = globalState
		_ = _lhs22
		_lhs20.F4 = _rhs16
		_176_v94 = _rhs17
		_lhs21.F4 = _rhs18
		_lhs22.F3 = _rhs19
		(globalState).F4 = p2
		var _178_v96 _dafny.Map
		_ = _178_v96
		_178_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_147_v66).Cardinality())).Cmp(p0) != 0, _dafny.MultiSetFromSeq(_56_v0))
		_178_v96 = (_178_v96).Update(p2, _dafny.MultiSetFromSeq(_56_v0))
		var _arr2 _dafny.Array = _dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))
		_ = _arr2
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))
		_ = _index29
		(_arr2).ArraySet1(p0, (_index29).Int())
		var _arr3 _dafny.Array = _dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))
		_ = _arr3
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))
		_ = _index30
		(_arr3).ArraySet1((Companion_Default___.Fm16(p1, globalState)).Minus(p0), (_index30).Int())
		var _179_v97 D1
		_ = _179_v97
		_179_v97 = Companion_D1_.Create_DC1_(_dafny.CodePoint('d'))
		var _source5 D1 = Companion_D1_.Create_DC3_(_179_v97)
		_ = _source5
		if _source5.Is_DC2() {
			var _nwElement0_6 _dafny.Sequence = _147_v66
			_ = _nwElement0_6
			var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(7))
			_ = _nw28
			(_nw28).ArraySet1(_nwElement0_6, 0)
			(_nw28).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mfn"), 1)
			(_nw28).ArraySet1((func() _dafny.Sequence {
				if !(p1) {
					return _147_v66
				}
				return _147_v66
			})(), 2)
			(_nw28).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ilvyw"), 3)
			(_nw28).ArraySet1(_147_v66, 4)
			(_nw28).ArraySet1(_147_v66, 5)
			(_nw28).ArraySet1(_147_v66, 6)
			r0 = _nw28
			_176_v94 = _176_v94
			var _180_v98 D3
			_ = _180_v98
			_180_v98 = Companion_D3_.Create_DC7_(p1)
			_180_v98 = _180_v98
			var _181_v99 _dafny.Array
			_ = _181_v99
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw29
			_181_v99 = _nw29
			var _arr4 _dafny.Array = _dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))
			_ = _arr4
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))
			_ = _index31
			var _rhs20 _dafny.Array = _181_v99
			_ = _rhs20
			var _rhs21 _dafny.Int = (_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))).Int()).(_dafny.Int)
			_ = _rhs21
			var _lhs23 _dafny.Array = _dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))
			_ = _lhs24
			_181_v99 = _rhs20
			(_lhs23).ArraySet1(_rhs21, (_lhs24).Int())
		} else if _source5.Is_DC1() {
			var _182___mcc_h8 _dafny.CodePoint = _source5.Get_().(D1_DC1).Cf1
			_ = _182___mcc_h8
			var _183_cf1 _dafny.CodePoint = _182___mcc_h8
			_ = _183_cf1
			(globalState).F3 = !_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("lq"), _183_cf1)
			var _184_v100 T1
			_ = _184_v100
			var _nw30 *C5 = New_C5_()
			_ = _nw30
			_nw30.Ctor__(_147_v66, (_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))).Int()).(_dafny.Int), p0, _dafny.IntOfInt64(780))
			_184_v100 = _nw30
			var _185_v101 _dafny.Map
			_ = _185_v101
			_185_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_184_v100, p0)
			var _186_v102 _dafny.Map
			_ = _186_v102
			_186_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_185_v101).Merge(_185_v101), p1)
			_186_v102 = (_186_v102).Update(_185_v101, p2)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))
			_ = _index32
			var _rhs22 bool = (_145_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))).Int()).(bool)
			_ = _rhs22
			var _rhs23 _dafny.Sequence = _56_v0
			_ = _rhs23
			var _lhs25 _dafny.Array = _145_v64
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))
			_ = _lhs26
			(_lhs25).ArraySet1(_rhs22, (_lhs26).Int())
			_56_v0 = _rhs23
			var _187_v103 _dafny.Array
			_ = _187_v103
			var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
			_ = _nw31
			_187_v103 = _nw31
			var _188_v104 _dafny.Sequence
			_ = _188_v104
			_188_v104 = _dafny.SeqOf(_145_v64, _145_v64, _145_v64)
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_187_v103), 0))
			_ = _index33
			(_187_v103).ArraySet1(_188_v104, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_187_v103), 0))
			_ = _index34
			(_187_v103).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_188_v104, _188_v104), _188_v104), (_index34).Int())
		} else {
			var _189___mcc_h9 D1 = _source5.Get_().(D1_DC3).Cf2
			_ = _189___mcc_h9
			var _190_cf2 D1 = _189___mcc_h9
			_ = _190_cf2
			var _191_v105 *C6
			_ = _191_v105
			var _nw32 *C6 = New_C6_()
			_ = _nw32
			_nw32.Ctor__(_dafny.IntOfUint32((_56_v0).Cardinality()), p0)
			_191_v105 = _nw32
			var _192_v106 _dafny.Sequence
			_ = _192_v106
			_192_v106 = _dafny.SeqOf(_191_v105, _191_v105)
			var _193_v107 _dafny.Map
			_ = _193_v107
			_193_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_192_v106).Cardinality()), _176_v94)
			var _194_v108 *C4
			_ = _194_v108
			var _nw33 *C4 = New_C4_()
			_ = _nw33
			_nw33.Ctor__(p2, _176_v94, _176_v94)
			_194_v108 = _nw33
			var _195_v109 _dafny.Map
			_ = _195_v109
			_195_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v108, (_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))).Int()).(_dafny.Int))
			var _196_v110 *C4
			_ = _196_v110
			var _nw34 *C4 = New_C4_()
			_ = _nw34
			_nw34.Ctor__(p1, p0, (func() _dafny.Int {
				if (_195_v109).Contains(_194_v108) {
					return (_195_v109).Get(_194_v108).(_dafny.Int)
				}
				return _dafny.IntOfInt64(325)
			})())
			_196_v110 = _nw34
			_176_v94 = (Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_193_v107).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_145_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))).Int()).(bool), _196_v110)).Cardinality()) {
					return (_193_v107).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_145_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))).Int()).(bool), _196_v110)).Cardinality()).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_56_v0).Cardinality())
			})(), _176_v94)).Plus((func() _dafny.Int {
				if (_193_v107).Contains(_dafny.IntOfInt64(64)) {
					return (_193_v107).Get(_dafny.IntOfInt64(64)).(_dafny.Int)
				}
				return _dafny.IntOfInt64(-306)
			})())
			var _197_v111 *C2
			_ = _197_v111
			var _nw35 *C2 = New_C2_()
			_ = _nw35
			_nw35.Ctor__((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))).Int()).(_dafny.Int), (_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))).Int()).(_dafny.Int))
			_197_v111 = _nw35
			var _arr5 _dafny.Array = _dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))
			_ = _arr5
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))
			_ = _index35
			var _rhs24 _dafny.Array = _145_v64
			_ = _rhs24
			var _rhs25 bool = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_147_v66, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(616))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_198_v63 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_199_i4 _dafny.Int) _dafny.CodePoint {
					return _198_v63
				}
			})(_144_v63)))), _147_v66)
			_ = _rhs25
			var _rhs26 _dafny.Int = _176_v94
			_ = _rhs26
			var _lhs27 *GlobalState = globalState
			_ = _lhs27
			var _lhs28 _dafny.Array = _dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_dafny.ArrayCastTo((_161_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_161_v79), 0))).Int()))), 0))
			_ = _lhs29
			_145_v64 = _rhs24
			_lhs27.F3 = _rhs25
			(_lhs28).ArraySet1(_rhs26, (_lhs29).Int())
			var _200_v112 _dafny.Array
			_ = _200_v112
			var _nw36 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
			_ = _nw36
			_200_v112 = _nw36
			var _201_v113 *C3
			_ = _201_v113
			var _nw37 *C3 = New_C3_()
			_ = _nw37
			_nw37.Ctor__(p1, p0, p0)
			_201_v113 = _nw37
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_200_v112), 0))
			_ = _index36
			(_200_v112).ArraySet1(_201_v113, (_index36).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_200_v112), 0))
			_ = _index37
			(_200_v112).ArraySet1(_201_v113, (_index37).Int())
		}
	}
	var _202_v114 _dafny.Map
	_ = _202_v114
	_202_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_147_v66), !((_145_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_145_v64), 0))).Int()).(bool)))
	var _203_v115 _dafny.MultiSet
	_ = _203_v115
	_203_v115 = _dafny.MultiSetOf(_147_v66)
	var _204_v116 *C0
	_ = _204_v116
	var _nw38 *C0 = New_C0_()
	_ = _nw38
	_nw38.Ctor__(_162_v80, p0, p0)
	_204_v116 = _nw38
	var _205_v117 _dafny.Map
	_ = _205_v117
	_205_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _204_v116)
	var _206_v118 _dafny.Map
	_ = _206_v118
	_206_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _205_v117)
	_202_v114 = (_202_v114).Update(_203_v115, (_206_v118).Contains(p0))
	var _207_v119 _dafny.Array
	_ = _207_v119
	var _nwElement0_7 _dafny.Sequence = _147_v66
	_ = _nwElement0_7
	var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(9))
	_ = _nw39
	(_nw39).ArraySet1(_nwElement0_7, 0)
	(_nw39).ArraySet1(_147_v66, 1)
	(_nw39).ArraySet1(_147_v66, 2)
	(_nw39).ArraySet1(_147_v66, 3)
	(_nw39).ArraySet1(_147_v66, 4)
	(_nw39).ArraySet1(_147_v66, 5)
	(_nw39).ArraySet1(_147_v66, 6)
	(_nw39).ArraySet1(_147_v66, 7)
	(_nw39).ArraySet1(_147_v66, 8)
	_207_v119 = _nw39
	var _208_v120 D14
	_ = _208_v120
	_208_v120 = Companion_D14_.Create_DC35_(_207_v119)
	r0 = (_208_v120).Dtor_cf48()
	var _209_v121 D5
	_ = _209_v121
	_209_v121 = Companion_D5_.Create_DC13_(p2, _144_v63)
	var _pat_let_tv0 = _147_v66
	_ = _pat_let_tv0
	var _pat_let_tv1 = p2
	_ = _pat_let_tv1
	r1 = func(_source6 D5) bool {
		if _source6.Is_DC12() {
			var _210___mcc_h10 bool = _source6.Get_().(D5_DC12).Cf19
			_ = _210___mcc_h10
			var _211_cf19 bool = _210___mcc_h10
			_ = _211_cf19
			return _dafny.Companion_Sequence_.IsPrefixOf(_pat_let_tv0, (Companion_D7_.Create_DC17_(_dafny.UnicodeSeqOfUtf8Bytes("v"))).Dtor_cf26())
		} else if _source6.Is_DC13() {
			var _212___mcc_h11 bool = _source6.Get_().(D5_DC13).Cf20
			_ = _212___mcc_h11
			var _213___mcc_h12 _dafny.CodePoint = _source6.Get_().(D5_DC13).Cf21
			_ = _213___mcc_h12
			var _214_cf21 _dafny.CodePoint = _213___mcc_h12
			_ = _214_cf21
			var _215_cf20 bool = _212___mcc_h11
			_ = _215_cf20
			return _215_cf20
		} else if _source6.Is_DC11() {
			var _216___mcc_h13 _dafny.Map = _source6.Get_().(D5_DC11).Cf18
			_ = _216___mcc_h13
			var _217_cf18 _dafny.Map = _216___mcc_h13
			_ = _217_cf18
			return _pat_let_tv1
		} else {
			var _218___mcc_h14 D5 = _source6.Get_().(D5_DC14).Cf22
			_ = _218___mcc_h14
			var _219_cf22 D5 = _218___mcc_h14
			_ = _219_cf22
			return true
		}
	}(_209_v121)
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _220_v0 _dafny.Array
	_ = _220_v0
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(18)
	_ = _len0_1
	var _nw40 _dafny.Array
	_ = _nw40
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw40 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) bool = func(_221_i1 _dafny.Int) bool {
			return false
		}
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw40 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw40).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw40).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_220_v0 = _nw40
	var _222_v1 _dafny.Map
	_ = _222_v1
	_222_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(230))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}(func(_223_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})), _220_v0)
	var _224_v2 _dafny.Sequence
	_ = _224_v2
	_224_v2 = _dafny.UnicodeSeqOfUtf8Bytes("w")
	var _225_globalState *GlobalState
	_ = _225_globalState
	var _nw41 *GlobalState = New_GlobalState_()
	_ = _nw41
	_nw41.Ctor__(_dafny.IntOfInt64(124), true, (_222_v1).Merge(_222_v1), false, true, _dafny.IntOfInt64(611), _224_v2)
	_225_globalState = _nw41
	var _226_v3 bool
	_ = _226_v3
	_226_v3 = false
	(_225_globalState).F3 = _226_v3
	var _227_v4 _dafny.Int
	_ = _227_v4
	_227_v4 = _dafny.IntOfInt64(756)
	var _228_i2 _dafny.Int
	_ = _228_i2
	_228_i2 = _dafny.Zero
	{
		for (Companion_Default___.SafeModuloInt(_227_v4, _227_v4)).Cmp(_227_v4) <= 0 {
			{
				if (_228_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_228_i2 = (_228_i2).Plus(_dafny.One)
				var _229_v5 _dafny.Array
				_ = _229_v5
				var _230_v6 bool
				_ = _230_v6
				var _out0 _dafny.Array
				_ = _out0
				var _out1 bool
				_ = _out1
				_out0, _out1 = Companion_Default___.M0(_227_v4, _226_v3, _226_v3, _225_globalState)
				_229_v5 = _out0
				_230_v6 = _out1
				var _231_v7 _dafny.Map
				_ = _231_v7
				_231_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _230_v6)
				var _232_v8 _dafny.Sequence
				_ = _232_v8
				_232_v8 = _dafny.SeqOf(_226_v3)
				_231_v7 = (_231_v7).Update(_dafny.Companion_Sequence_.Contains(_232_v8, _226_v3), _226_v3)
				var _233_v9 _dafny.Set
				_ = _233_v9
				_233_v9 = _dafny.SetOf(_224_v2)
				var _234_v10 *C2
				_ = _234_v10
				var _nw42 *C2 = New_C2_()
				_ = _nw42
				_nw42.Ctor__(((_233_v9).Cardinality()).Minus(_227_v4), (_dafny.Zero).Minus(_227_v4))
				_234_v10 = _nw42
				_231_v7 = (_231_v7).Update(_226_v3, _226_v3)
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _235_v11 _dafny.Sequence
	_ = _235_v11
	_235_v11 = _dafny.SeqOf(_dafny.IntOfInt64(88))
	(_225_globalState).F4 = ((_235_v11).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_227_v4), _dafny.IntOfUint32((_235_v11).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_224_v2).Cardinality())) == 0
	var _236_v12 _dafny.Map
	_ = _236_v12
	_236_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(823), false)
	var _237_v13 *C4
	_ = _237_v13
	var _nw43 *C4 = New_C4_()
	_ = _nw43
	_nw43.Ctor__(_226_v3, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_v3, _226_v3)).Cardinality(), _227_v4)
	_237_v13 = _nw43
	var _238_v14 _dafny.Sequence
	_ = _238_v14
	_238_v14 = _dafny.SeqOf(_237_v13, _237_v13)
	var _239_v15 _dafny.Array
	_ = _239_v15
	var _nwElement0_8 _dafny.Int = _dafny.IntOfInt64(218)
	_ = _nwElement0_8
	var _nw44 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(23))
	_ = _nw44
	(_nw44).ArraySet1(_nwElement0_8, 0)
	(_nw44).ArraySet1(_dafny.IntOfUint32((_224_v2).Cardinality()), 1)
	(_nw44).ArraySet1(_227_v4, 2)
	(_nw44).ArraySet1((_236_v12).Cardinality(), 3)
	(_nw44).ArraySet1(_227_v4, 4)
	(_nw44).ArraySet1(_dafny.IntOfInt64(709), 5)
	(_nw44).ArraySet1(_227_v4, 6)
	(_nw44).ArraySet1(_dafny.IntOfUint32((_238_v14).Cardinality()), 7)
	(_nw44).ArraySet1(_227_v4, 8)
	(_nw44).ArraySet1(_227_v4, 9)
	(_nw44).ArraySet1(_227_v4, 10)
	(_nw44).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_237_v13).F11(), (_237_v13).F11())).Cardinality()), 11)
	(_nw44).ArraySet1(_227_v4, 12)
	(_nw44).ArraySet1((_dafny.Zero).Minus(_227_v4), 13)
	(_nw44).ArraySet1(_dafny.Zero, 14)
	(_nw44).ArraySet1(_227_v4, 15)
	(_nw44).ArraySet1(_227_v4, 16)
	(_nw44).ArraySet1(_227_v4, 17)
	(_nw44).ArraySet1(_227_v4, 18)
	(_nw44).ArraySet1(_227_v4, 19)
	(_nw44).ArraySet1(_227_v4, 20)
	(_nw44).ArraySet1(_227_v4, 21)
	(_nw44).ArraySet1(_227_v4, 22)
	_239_v15 = _nw44
	var _240_v16 _dafny.Sequence
	_ = _240_v16
	_240_v16 = _dafny.SeqOf(_239_v15)
	var _241_v17 *C0
	_ = _241_v17
	var _nw45 *C0 = New_C0_()
	_ = _nw45
	_nw45.Ctor__((_240_v16).Select((Companion_Default___.SafeIndex(_227_v4, _dafny.IntOfUint32((_240_v16).Cardinality()))).Uint32()).(_dafny.Array), _dafny.IntOfInt64(525), _227_v4)
	_241_v17 = _nw45
	var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_220_v0), 0))
	_ = _index38
	(_220_v0).ArraySet1(!(!(!((_237_v13).F11()) || ((_237_v13).F11()))), (_index38).Int())
	var _242_v18 _dafny.Array
	_ = _242_v18
	var _nwElement0_9 _dafny.Array = _239_v15
	_ = _nwElement0_9
	var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.One)
	_ = _nw46
	(_nw46).ArraySet1(_nwElement0_9, 0)
	_242_v18 = _nw46
	var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_242_v18), 0))
	_ = _index39
	(_242_v18).ArraySet1(_241_v17.F13, (_index39).Int())
	var _243_v19 _dafny.Sequence
	_ = _243_v19
	_243_v19 = _dafny.SeqOf((_237_v13).F11(), _dafny.Companion_Sequence_.Equal(Companion_Default___.Fm8(true, _235_v11, _226_v3, _225_globalState), _224_v2), (func() bool {
		if !((_237_v13).F11()) {
			return _226_v3
		}
		return (_237_v13).F11()
	})(), !(_236_v12).Equals(_236_v12), ((_237_v13).F11()) == (_226_v3))
	var _244_v20 _dafny.CodePoint
	_ = _244_v20
	_244_v20 = _dafny.CodePoint('b')
	var _245_v21 D1
	_ = _245_v21
	_245_v21 = Companion_D1_.Create_DC1_(_244_v20)
	var _246_v22 D1
	_ = _246_v22
	_246_v22 = Companion_D1_.Create_DC3_(Companion_D1_.Create_DC1_(_244_v20))
	var _247_v23 _dafny.MultiSet
	_ = _247_v23
	_247_v23 = _dafny.MultiSetOf(Companion_D1_.Create_DC3_(_245_v21), _246_v22)
	var _248_v24 _dafny.Set
	_ = _248_v24
	_248_v24 = _dafny.SetOf(_247_v23, _247_v23, _dafny.MultiSetOf(_246_v22))
	var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_220_v0), 0))
	_ = _index40
	var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_242_v18), 0))
	_ = _index41
	var _rhs27 bool = !((func() bool {
		if (_227_v4).Cmp(_dafny.IntOfUint32((_224_v2).Cardinality())) >= 0 {
			return _226_v3
		}
		return _226_v3
	})())
	_ = _rhs27
	var _rhs28 _dafny.Array = _239_v15
	_ = _rhs28
	var _rhs29 _dafny.Sequence = _243_v19
	_ = _rhs29
	var _rhs30 _dafny.Int = Companion_Default___.Fm20(_226_v3, (_248_v24).IsProperSubsetOf(Companion_Default___.Fm36(_226_v3, (_237_v13).F11(), _225_globalState)), (_237_v13).F11(), _225_globalState)
	_ = _rhs30
	var _rhs31 *C4 = _237_v13
	_ = _rhs31
	var _lhs30 _dafny.Array = _220_v0
	_ = _lhs30
	var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_220_v0), 0))
	_ = _lhs31
	var _lhs32 _dafny.Array = _242_v18
	_ = _lhs32
	var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_242_v18), 0))
	_ = _lhs33
	(_lhs30).ArraySet1(_rhs27, (_lhs31).Int())
	(_lhs32).ArraySet1(_rhs28, (_lhs33).Int())
	_243_v19 = _rhs29
	_227_v4 = _rhs30
	_237_v13 = _rhs31
	var _249_v25 _dafny.MultiSet
	_ = _249_v25
	_249_v25 = _dafny.MultiSetOf(_241_v17.F13)
	var _250_v26 _dafny.Set
	_ = _250_v26
	_250_v26 = _dafny.SetOf((_237_v13).F11())
	var _251_v27 _dafny.Map
	_ = _251_v27
	_251_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_237_v13).F11(), (_241_v17).Fm1(_250_v26, _227_v4, _227_v4, _227_v4, _225_globalState))
	var _252_v28 _dafny.Map
	_ = _252_v28
	_252_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_227_v4), (_249_v25).Update(_241_v17.F13, Companion_Default___.Abs(_227_v4)))
	var _rhs32 _dafny.Array = _220_v0
	_ = _rhs32
	var _rhs33 _dafny.Int = (func() _dafny.Int {
		if (func() bool {
			if (_251_v27).Contains((_237_v13).F11()) {
				return (_251_v27).Get((_237_v13).F11()).(bool)
			}
			return (_237_v13).F11()
		})() {
			return ((_250_v26).Difference(_250_v26)).Cardinality()
		}
		return _227_v4
	})()
	_ = _rhs33
	var _rhs34 _dafny.MultiSet = (func() _dafny.MultiSet {
		if (_252_v28).Contains((_dafny.Zero).Minus(_227_v4)) {
			return (_252_v28).Get((_dafny.Zero).Minus(_227_v4)).(_dafny.MultiSet)
		}
		return _249_v25
	})()
	_ = _rhs34
	_220_v0 = _rhs32
	_227_v4 = _rhs33
	_249_v25 = _rhs34
	var _253_v29 T2
	_ = _253_v29
	var _nw47 *C2 = New_C2_()
	_ = _nw47
	_nw47.Ctor__(_227_v4, _227_v4)
	_253_v29 = _nw47
	var _254_v30 _dafny.Map
	_ = _254_v30
	_254_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v29, _236_v12)
	var _255_v31 *C2
	_ = _255_v31
	var _nw48 *C2 = New_C2_()
	_ = _nw48
	_nw48.Ctor__(_227_v4, (_dafny.IntOfInt64(356)).Plus(((func() _dafny.Map {
		if (_254_v30).Contains(_253_v29) {
			return (_254_v30).Get(_253_v29).(_dafny.Map)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_253_v29).F8(), _241_v17)).Cardinality(), (_220_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_220_v0), 0))).Int()).(bool))
	})()).Cardinality()))
	_255_v31 = _nw48
	var _256_i3 _dafny.Int
	_ = _256_i3
	_256_i3 = _dafny.Zero
	{
		for _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_224_v2, _224_v2), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(683))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}((func(_265_v20 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_266_i4 _dafny.Int) _dafny.CodePoint {
				return _265_v20
			}
		})(_244_v20)))) {
			{
				if (_256_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_256_i3 = (_256_i3).Plus(_dafny.One)
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_220_v0), 0))
				_ = _index42
				(_220_v0).ArraySet1(!((_237_v13).F11()), (_index42).Int())
				var _257_v32 _dafny.Array
				_ = _257_v32
				var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
				_ = _nw49
				_257_v32 = _nw49
				var _258_v33 _dafny.Map
				_ = _258_v33
				_258_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_220_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_220_v0), 0))).Int()).(bool)), (_253_v29).F7())
				var _259_v34 D6
				_ = _259_v34
				_259_v34 = Companion_D6_.Create_DC16_(_258_v33, _243_v19)
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_257_v32), 0))
				_ = _index43
				(_257_v32).ArraySet1((_259_v34).Dtor_cf25(), (_index43).Int())
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_257_v32), 0))
				_ = _index44
				(_257_v32).ArraySet1(_243_v19, (_index44).Int())
				var _260_v35 D7
				_ = _260_v35
				_260_v35 = Companion_D7_.Create_DC20_(_244_v20, _227_v4)
				var _261_v36 _dafny.Map
				_ = _261_v36
				_261_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v35, _244_v20)
				var _262_v37 *C4
				_ = _262_v37
				var _nw50 *C4 = New_C4_()
				_ = _nw50
				_nw50.Ctor__(_226_v3, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_261_v36, _261_v36)).Cardinality())), (_253_v29).F7())
				_262_v37 = _nw50
				_224_v2 = (func() _dafny.Sequence {
					if (_220_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_220_v0), 0))).Int()).(bool) {
						return _dafny.Companion_Sequence_.Concatenate((_255_v31).Fm14(_225_globalState), _224_v2)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(220))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg22 _dafny.Int) interface{} {
							return coer22(arg22)
						}
					}((func(_263_v20 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_264_i5 _dafny.Int) _dafny.CodePoint {
							return _263_v20
						}
					})(_244_v20)))
				})()
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _267_v38 _dafny.Array
	_ = _267_v38
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(5)
	_ = _len0_2
	var _nw51 _dafny.Array
	_ = _nw51
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw51 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Sequence = (func(_268_v0 _dafny.Array, _269_v13 *C4) func(_dafny.Int) _dafny.Sequence {
			return func(_270_i6 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(_dafny.SetOf((_268_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_268_v0), 0))).Int()).(bool), (_269_v13).F11()))
			}
		})(_220_v0, _237_v13)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw51 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw51).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw51).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_267_v38 = _nw51
	var _271_v39 _dafny.Sequence
	_ = _271_v39
	_271_v39 = _dafny.SeqOf(_250_v26)
	var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_267_v38), 0))
	_ = _index45
	(_267_v38).ArraySet1(_271_v39, (_index45).Int())
	var _272_v40 D4
	_ = _272_v40
	_272_v40 = Companion_D4_.Create_DC10_(_227_v4, (_237_v13).F11())
	var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_220_v0), 0))
	_ = _index46
	var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_267_v38), 0))
	_ = _index47
	var _rhs35 bool = (_272_v40).Dtor_cf17()
	_ = _rhs35
	var _rhs36 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_271_v39, _271_v39)
	_ = _rhs36
	var _lhs34 _dafny.Array = _220_v0
	_ = _lhs34
	var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_220_v0), 0))
	_ = _lhs35
	var _lhs36 _dafny.Array = _267_v38
	_ = _lhs36
	var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_267_v38), 0))
	_ = _lhs37
	(_lhs34).ArraySet1(_rhs35, (_lhs35).Int())
	(_lhs36).ArraySet1(_rhs36, (_lhs37).Int())
	var _273_v41 *C5
	_ = _273_v41
	var _nw52 *C5 = New_C5_()
	_ = _nw52
	_nw52.Ctor__(_224_v2, _227_v4, (_253_v29).F8(), (_253_v29).F7())
	_273_v41 = _nw52
	var _274_v42 _dafny.Map
	_ = _274_v42
	_274_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_273_v41, (_273_v41).F10())
	var _275_v43 _dafny.MultiSet
	_ = _275_v43
	_275_v43 = _dafny.MultiSetOf(_220_v0, _220_v0)
	(_225_globalState).F4 = (_237_v13).Fm1(_250_v26, (_253_v29).F8(), (func() _dafny.Int {
		if (_274_v42).Contains(_273_v41) {
			return (_274_v42).Get(_273_v41).(_dafny.Int)
		}
		return _227_v4
	})(), (func() _dafny.Int {
		if (_275_v43).Contains(_220_v0) {
			return (_275_v43).Multiplicity(_220_v0)
		}
		return _dafny.IntOfUint32((_235_v11).Cardinality())
	})(), _225_globalState)
	(_225_globalState).F3 = !(((Companion_Default___.Fm9(_226_v3, _225_globalState)).Cardinality()).Cmp(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_273_v41).F9(), (Companion_Default___.SafeIndex((_250_v26).Cardinality(), _dafny.IntOfUint32(((_273_v41).F9()).Cardinality()))).Uint32(), _244_v20)).Cardinality()))).Times((_253_v29).F7())) > 0)
	var _hi0 _dafny.Int = _227_v4
	_ = _hi0
	for _276_i7 := ((_253_v29).F7()).Minus(_dafny.IntOfUint32((_235_v11).Cardinality())); _276_i7.Cmp(_hi0) < 0; _276_i7 = _276_i7.Plus(_dafny.One) {
		_227_v4 = _227_v4
		var _277_v44 _dafny.Sequence
		_ = _277_v44
		_277_v44 = _dafny.SeqOf(_253_v29, _253_v29)
		var _278_v45 *C2
		_ = _278_v45
		var _nw53 *C2 = New_C2_()
		_ = _nw53
		_nw53.Ctor__((_253_v29).F7(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_277_v44, _277_v44)).Cardinality()))
		_278_v45 = _nw53
		_278_v45 = _278_v45
		_278_v45 = _278_v45
		(_278_v45).M1(_224_v2, _225_globalState)
	}
	var _279_v46 _dafny.Array
	_ = _279_v46
	var _nw54 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(9))
	_ = _nw54
	_279_v46 = _nw54
	for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_279_v46), 0))); ; {
		_guard_loop_0, _ok23 := _iter23()
		if !_ok23 {
			break
		}
		var _280_i8 _dafny.Int
		_280_i8 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_280_i8).Sign() != -1) && ((_280_i8).Cmp(_dafny.ArrayLen((_279_v46), 0)) < 0)) {
			(_279_v46).ArraySet1(Companion_D8_.Create_DC22_((_dafny.Zero).Minus((_273_v41).F10())), (_280_i8).Int())
		}
	}
	for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_241_v17.F13), 0))); ; {
		_guard_loop_1, _ok24 := _iter24()
		if !_ok24 {
			break
		}
		var _281_i9 _dafny.Int
		_281_i9 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_281_i9).Sign() != -1) && ((_281_i9).Cmp(_dafny.ArrayLen((_241_v17.F13), 0)) < 0)) {
			var _arr6 _dafny.Array = _241_v17.F13
			_ = _arr6
			(_arr6).ArraySet1(Companion_Default___.SafeDivisionInt(_281_i9, (_253_v29).F7()), (_281_i9).Int())
		}
	}
	(_225_globalState).F3 = !(_226_v3)
	_224_v2 = _224_v2
	_dafny.Print((_220_v0).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v0).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v1).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_224_v2.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_225_globalState).F2()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_225_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_225_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_globalState).F6().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_226_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_227_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_228_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_235_v11, _dafny.SeqOf(_dafny.IntOfInt64(88))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_236_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(823), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_237_v13).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_238_v14).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v15).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_240_v16).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v17.F13).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_242_v18).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_243_v19, _dafny.SeqOf(false, false, false, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_244_v20)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_245_v21).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_v22).Dtor_cf2()).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_247_v23).Equals(_dafny.MultiSetOf(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC1_(_dafny.CodePoint('b'))), Companion_D1_.Create_DC3_(Companion_D1_.Create_DC1_(_dafny.CodePoint('b'))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_248_v24).Equals(_dafny.SetOf(_dafny.MultiSetOf(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC1_(_dafny.CodePoint('b'))), Companion_D1_.Create_DC3_(Companion_D1_.Create_DC1_(_dafny.CodePoint('b')))), _dafny.MultiSetOf(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC1_(_dafny.CodePoint('b')))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_249_v25).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v26).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_251_v27).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v28).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v29).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v29).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v30).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_256_i3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_267_v38).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.SetOf(true, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_267_v38).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.SetOf(true, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_267_v38).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.SetOf(true, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_267_v38).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.SetOf(false), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_267_v38).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.SetOf(true, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_271_v39, _dafny.SeqOf(_dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v40).Dtor_cf16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v40).Dtor_cf17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_273_v41).F9().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_273_v41).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_v42).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v43).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_279_v46).ArrayGet1((_dafny.Zero).Int()).(D8)).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_279_v46).ArrayGet1((_dafny.One).Int()).(D8)).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_279_v46).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D8)).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_279_v46).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D8)).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_279_v46).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D8)).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_279_v46).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D8)).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_279_v46).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D8)).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_279_v46).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D8)).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_279_v46).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D8)).Dtor_cf35())
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

func (CompanionStruct_D0_) Default() bool {
	return false
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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

type D1_DC2 struct {
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_() D1 {
	return D1{D1_DC2{}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC1 struct {
	Cf1 _dafny.CodePoint
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.CodePoint) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

type D1_DC3 struct {
	Cf2 D1
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf2 D1) D1 {
	return D1{D1_DC3{Cf2}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_()
}

func (_this D1) Dtor_cf1() _dafny.CodePoint {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) Dtor_cf2() D1 {
	return _this.Get_().(D1_DC3).Cf2
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2"
		}
	case D1_DC1:
		{
			return "D1.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf2) + ")"
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
			_, ok := other.Get_().(D1_DC2)
			return ok
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1 == data2.Cf1
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf2.Equals(data2.Cf2)
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

type D2_DC5 struct {
	Cf4 bool
	Cf5 _dafny.Int
	Cf6 _dafny.Int
	Cf7 _dafny.Int
	Cf8 bool
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf4 bool, Cf5 _dafny.Int, Cf6 _dafny.Int, Cf7 _dafny.Int, Cf8 bool) D2 {
	return D2{D2_DC5{Cf4, Cf5, Cf6, Cf7, Cf8}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC4 struct {
	Cf3 _dafny.MultiSet
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf3 _dafny.MultiSet) D2 {
	return D2{D2_DC4{Cf3}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(false, _dafny.Zero, _dafny.Zero, _dafny.Zero, false)
}

func (_this D2) Dtor_cf4() bool {
	return _this.Get_().(D2_DC5).Cf4
}

func (_this D2) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf5
}

func (_this D2) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf6
}

func (_this D2) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf7
}

func (_this D2) Dtor_cf8() bool {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf3() _dafny.MultiSet {
	return _this.Get_().(D2_DC4).Cf3
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf3) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf4 == data2.Cf4 && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8 == data2.Cf8
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf3.Equals(data2.Cf3)
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
	Cf10 bool
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf10 bool) D3 {
	return D3{D3_DC7{Cf10}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC6 struct {
	Cf9 _dafny.Set
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf9 _dafny.Set) D3 {
	return D3{D3_DC6{Cf9}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC7_(false)
}

func (_this D3) Dtor_cf10() bool {
	return _this.Get_().(D3_DC7).Cf10
}

func (_this D3) Dtor_cf9() _dafny.Set {
	return _this.Get_().(D3_DC6).Cf9
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf10) + ")"
		}
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
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf10 == data2.Cf10
		}
	case D3_DC6:
		{
			data2, ok := other.Get_().(D3_DC6)
			return ok && data1.Cf9.Equals(data2.Cf9)
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
	Cf12 _dafny.Int
	Cf13 _dafny.Sequence
	Cf14 _dafny.Int
	Cf15 _dafny.Int
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf12 _dafny.Int, Cf13 _dafny.Sequence, Cf14 _dafny.Int, Cf15 _dafny.Int) D4 {
	return D4{D4_DC9{Cf12, Cf13, Cf14, Cf15}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC10 struct {
	Cf16 _dafny.Int
	Cf17 bool
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf16 _dafny.Int, Cf17 bool) D4 {
	return D4{D4_DC10{Cf16, Cf17}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC8 struct {
	Cf11 _dafny.Array
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf11 _dafny.Array) D4 {
	return D4{D4_DC8{Cf11}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC9_(_dafny.Zero, _dafny.EmptySeq, _dafny.Zero, _dafny.Zero)
}

func (_this D4) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf12
}

func (_this D4) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D4_DC9).Cf13
}

func (_this D4) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf14
}

func (_this D4) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf16
}

func (_this D4) Dtor_cf17() bool {
	return _this.Get_().(D4_DC10).Cf17
}

func (_this D4) Dtor_cf11() _dafny.Array {
	return _this.Get_().(D4_DC8).Cf11
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf12.Cmp(data2.Cf12) == 0 && data1.Cf13.Equals(data2.Cf13) && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15.Cmp(data2.Cf15) == 0
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17 == data2.Cf17
		}
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
			return ok && data1.Cf11 == data2.Cf11
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
	Cf19 bool
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf19 bool) D5 {
	return D5{D5_DC12{Cf19}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC13 struct {
	Cf20 bool
	Cf21 _dafny.CodePoint
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf20 bool, Cf21 _dafny.CodePoint) D5 {
	return D5{D5_DC13{Cf20, Cf21}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC11 struct {
	Cf18 _dafny.Map
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf18 _dafny.Map) D5 {
	return D5{D5_DC11{Cf18}}
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
	return Companion_D5_.Create_DC12_(false)
}

func (_this D5) Dtor_cf19() bool {
	return _this.Get_().(D5_DC12).Cf19
}

func (_this D5) Dtor_cf20() bool {
	return _this.Get_().(D5_DC13).Cf20
}

func (_this D5) Dtor_cf21() _dafny.CodePoint {
	return _this.Get_().(D5_DC13).Cf21
}

func (_this D5) Dtor_cf18() _dafny.Map {
	return _this.Get_().(D5_DC11).Cf18
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
			return "D5.DC12" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf18) + ")"
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
			return ok && data1.Cf19 == data2.Cf19
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21 == data2.Cf21
		}
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
			return ok && data1.Cf18.Equals(data2.Cf18)
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

type D6_DC16 struct {
	Cf24 _dafny.Map
	Cf25 _dafny.Sequence
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf24 _dafny.Map, Cf25 _dafny.Sequence) D6 {
	return D6{D6_DC16{Cf24, Cf25}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC15 struct {
	Cf23 _dafny.Array
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf23 _dafny.Array) D6 {
	return D6{D6_DC15{Cf23}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC16_(_dafny.EmptyMap, _dafny.EmptySeq)
}

func (_this D6) Dtor_cf24() _dafny.Map {
	return _this.Get_().(D6_DC16).Cf24
}

func (_this D6) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D6_DC16).Cf25
}

func (_this D6) Dtor_cf23() _dafny.Array {
	return _this.Get_().(D6_DC15).Cf23
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
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
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf24.Equals(data2.Cf24) && data1.Cf25.Equals(data2.Cf25)
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf23 == data2.Cf23
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
	Cf27 bool
	Cf28 bool
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf27 bool, Cf28 bool) D7 {
	return D7{D7_DC18{Cf27, Cf28}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC19 struct {
	Cf29 _dafny.Int
	Cf30 bool
	Cf31 _dafny.Int
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf29 _dafny.Int, Cf30 bool, Cf31 _dafny.Int) D7 {
	return D7{D7_DC19{Cf29, Cf30, Cf31}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC20 struct {
	Cf32 _dafny.CodePoint
	Cf33 _dafny.Int
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf32 _dafny.CodePoint, Cf33 _dafny.Int) D7 {
	return D7{D7_DC20{Cf32, Cf33}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC17 struct {
	Cf26 _dafny.Sequence
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf26 _dafny.Sequence) D7 {
	return D7{D7_DC17{Cf26}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC18_(false, false)
}

func (_this D7) Dtor_cf27() bool {
	return _this.Get_().(D7_DC18).Cf27
}

func (_this D7) Dtor_cf28() bool {
	return _this.Get_().(D7_DC18).Cf28
}

func (_this D7) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D7_DC19).Cf29
}

func (_this D7) Dtor_cf30() bool {
	return _this.Get_().(D7_DC19).Cf30
}

func (_this D7) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D7_DC19).Cf31
}

func (_this D7) Dtor_cf32() _dafny.CodePoint {
	return _this.Get_().(D7_DC20).Cf32
}

func (_this D7) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D7_DC20).Cf33
}

func (_this D7) Dtor_cf26() _dafny.Sequence {
	return _this.Get_().(D7_DC17).Cf26
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D7_DC17:
		{
			return "D7.DC17" + "(" + data.Cf26.VerbatimString(true) + ")"
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
			return ok && data1.Cf27 == data2.Cf27 && data1.Cf28 == data2.Cf28
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf29.Cmp(data2.Cf29) == 0 && data1.Cf30 == data2.Cf30 && data1.Cf31.Cmp(data2.Cf31) == 0
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf32 == data2.Cf32 && data1.Cf33.Cmp(data2.Cf33) == 0
		}
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
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

type D8_DC22 struct {
	Cf35 _dafny.Int
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf35 _dafny.Int) D8 {
	return D8{D8_DC22{Cf35}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC21 struct {
	Cf34 *C3
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf34 *C3) D8 {
	return D8{D8_DC21{Cf34}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC22_(_dafny.Zero)
}

func (_this D8) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D8_DC22).Cf35
}

func (_this D8) Dtor_cf34() *C3 {
	return _this.Get_().(D8_DC21).Cf34
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf35) + ")"
		}
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf34) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf35.Cmp(data2.Cf35) == 0
		}
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf34 == data2.Cf34
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

type D9_DC24 struct {
	Cf37 _dafny.Int
	Cf38 _dafny.Int
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf37 _dafny.Int, Cf38 _dafny.Int) D9 {
	return D9{D9_DC24{Cf37, Cf38}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC23 struct {
	Cf36 _dafny.Array
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf36 _dafny.Array) D9 {
	return D9{D9_DC23{Cf36}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC24_(_dafny.Zero, _dafny.Zero)
}

func (_this D9) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D9_DC24).Cf37
}

func (_this D9) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D9_DC24).Cf38
}

func (_this D9) Dtor_cf36() _dafny.Array {
	return _this.Get_().(D9_DC23).Cf36
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf36) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf36 == data2.Cf36
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

type D10_DC26 struct {
	Cf40 _dafny.Int
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf40 _dafny.Int) D10 {
	return D10{D10_DC26{Cf40}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

type D10_DC25 struct {
	Cf39 _dafny.Map
}

func (D10_DC25) isD10() {}

func (CompanionStruct_D10_) Create_DC25_(Cf39 _dafny.Map) D10 {
	return D10{D10_DC25{Cf39}}
}

func (_this D10) Is_DC25() bool {
	_, ok := _this.Get_().(D10_DC25)
	return ok
}

type D10_DC27 struct {
	Cf41 D10
}

func (D10_DC27) isD10() {}

func (CompanionStruct_D10_) Create_DC27_(Cf41 D10) D10 {
	return D10{D10_DC27{Cf41}}
}

func (_this D10) Is_DC27() bool {
	_, ok := _this.Get_().(D10_DC27)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC26_(_dafny.Zero)
}

func (_this D10) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D10_DC26).Cf40
}

func (_this D10) Dtor_cf39() _dafny.Map {
	return _this.Get_().(D10_DC25).Cf39
}

func (_this D10) Dtor_cf41() D10 {
	return _this.Get_().(D10_DC27).Cf41
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D10_DC25:
		{
			return "D10.DC25" + "(" + _dafny.String(data.Cf39) + ")"
		}
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf41) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC26:
		{
			data2, ok := other.Get_().(D10_DC26)
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0
		}
	case D10_DC25:
		{
			data2, ok := other.Get_().(D10_DC25)
			return ok && data1.Cf39.Equals(data2.Cf39)
		}
	case D10_DC27:
		{
			data2, ok := other.Get_().(D10_DC27)
			return ok && data1.Cf41.Equals(data2.Cf41)
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

type D11_DC29 struct {
	Cf43 _dafny.Map
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf43 _dafny.Map) D11 {
	return D11{D11_DC29{Cf43}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

type D11_DC30 struct {
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_() D11 {
	return D11{D11_DC30{}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC28 struct {
	Cf42 T1
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_(Cf42 T1) D11 {
	return D11{D11_DC28{Cf42}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

type D11_DC31 struct {
	Cf44 D11
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf44 D11) D11 {
	return D11{D11_DC31{Cf44}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC29_(_dafny.EmptyMap)
}

func (_this D11) Dtor_cf43() _dafny.Map {
	return _this.Get_().(D11_DC29).Cf43
}

func (_this D11) Dtor_cf42() T1 {
	return _this.Get_().(D11_DC28).Cf42
}

func (_this D11) Dtor_cf44() D11 {
	return _this.Get_().(D11_DC31).Cf44
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D11_DC30:
		{
			return "D11.DC30"
		}
	case D11_DC28:
		{
			return "D11.DC28" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
			return ok && data1.Cf43.Equals(data2.Cf43)
		}
	case D11_DC30:
		{
			_, ok := other.Get_().(D11_DC30)
			return ok
		}
	case D11_DC28:
		{
			data2, ok := other.Get_().(D11_DC28)
			return ok && _dafny.AreEqual(data1.Cf42, data2.Cf42)
		}
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
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

type D12_DC33 struct {
	Cf46 _dafny.Sequence
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf46 _dafny.Sequence) D12 {
	return D12{D12_DC33{Cf46}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

type D12_DC32 struct {
	Cf45 _dafny.Map
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf45 _dafny.Map) D12 {
	return D12{D12_DC32{Cf45}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC33_(_dafny.EmptySeq)
}

func (_this D12) Dtor_cf46() _dafny.Sequence {
	return _this.Get_().(D12_DC33).Cf46
}

func (_this D12) Dtor_cf45() _dafny.Map {
	return _this.Get_().(D12_DC32).Cf45
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC33:
		{
			return "D12.DC33" + "(" + data.Cf46.VerbatimString(true) + ")"
		}
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf45) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC33:
		{
			data2, ok := other.Get_().(D12_DC33)
			return ok && data1.Cf46.Equals(data2.Cf46)
		}
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
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

type D13_DC34 struct {
	Cf47 *C1
}

func (D13_DC34) isD13() {}

func (CompanionStruct_D13_) Create_DC34_(Cf47 *C1) D13 {
	return D13{D13_DC34{Cf47}}
}

func (_this D13) Is_DC34() bool {
	_, ok := _this.Get_().(D13_DC34)
	return ok
}

func (CompanionStruct_D13_) Default() *C1 {
	return (*C1)(nil)
}

func (_this D13) Dtor_cf47() *C1 {
	return _this.Get_().(D13_DC34).Cf47
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC34:
		{
			return "D13.DC34" + "(" + _dafny.String(data.Cf47) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC34:
		{
			data2, ok := other.Get_().(D13_DC34)
			return ok && data1.Cf47 == data2.Cf47
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

type D14_DC36 struct {
	Cf49 bool
	Cf50 _dafny.Sequence
	Cf51 _dafny.Sequence
	Cf52 _dafny.Sequence
	Cf53 _dafny.CodePoint
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf49 bool, Cf50 _dafny.Sequence, Cf51 _dafny.Sequence, Cf52 _dafny.Sequence, Cf53 _dafny.CodePoint) D14 {
	return D14{D14_DC36{Cf49, Cf50, Cf51, Cf52, Cf53}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

type D14_DC35 struct {
	Cf48 _dafny.Array
}

func (D14_DC35) isD14() {}

func (CompanionStruct_D14_) Create_DC35_(Cf48 _dafny.Array) D14 {
	return D14{D14_DC35{Cf48}}
}

func (_this D14) Is_DC35() bool {
	_, ok := _this.Get_().(D14_DC35)
	return ok
}

type D14_DC37 struct {
	Cf54 D14
}

func (D14_DC37) isD14() {}

func (CompanionStruct_D14_) Create_DC37_(Cf54 D14) D14 {
	return D14{D14_DC37{Cf54}}
}

func (_this D14) Is_DC37() bool {
	_, ok := _this.Get_().(D14_DC37)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC36_(false, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.CodePoint('D'))
}

func (_this D14) Dtor_cf49() bool {
	return _this.Get_().(D14_DC36).Cf49
}

func (_this D14) Dtor_cf50() _dafny.Sequence {
	return _this.Get_().(D14_DC36).Cf50
}

func (_this D14) Dtor_cf51() _dafny.Sequence {
	return _this.Get_().(D14_DC36).Cf51
}

func (_this D14) Dtor_cf52() _dafny.Sequence {
	return _this.Get_().(D14_DC36).Cf52
}

func (_this D14) Dtor_cf53() _dafny.CodePoint {
	return _this.Get_().(D14_DC36).Cf53
}

func (_this D14) Dtor_cf48() _dafny.Array {
	return _this.Get_().(D14_DC35).Cf48
}

func (_this D14) Dtor_cf54() D14 {
	return _this.Get_().(D14_DC37).Cf54
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ")"
		}
	case D14_DC35:
		{
			return "D14.DC35" + "(" + _dafny.String(data.Cf48) + ")"
		}
	case D14_DC37:
		{
			return "D14.DC37" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && data1.Cf49 == data2.Cf49 && data1.Cf50.Equals(data2.Cf50) && data1.Cf51.Equals(data2.Cf51) && data1.Cf52.Equals(data2.Cf52) && data1.Cf53 == data2.Cf53
		}
	case D14_DC35:
		{
			data2, ok := other.Get_().(D14_DC35)
			return ok && data1.Cf48 == data2.Cf48
		}
	case D14_DC37:
		{
			data2, ok := other.Get_().(D14_DC37)
			return ok && data1.Cf54.Equals(data2.Cf54)
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

// Definition of trait T0
type T0 interface {
	String() string
	F7() _dafny.Int
	F8() _dafny.Int
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
	F7() _dafny.Int
	F8() _dafny.Int
	Fm0(p0 bool, globalState *GlobalState) _dafny.Set
	Fm1(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool
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
	Fm0(p0 bool, globalState *GlobalState) _dafny.Set
	Fm1(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool
	F7() _dafny.Int
	F8() _dafny.Int
	Fm2(p0 bool, globalState *GlobalState) _dafny.Sequence
	M1(p0 _dafny.Sequence, globalState *GlobalState)
	M2(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, bool)
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
	F3  bool
	F4  bool
	_f0 _dafny.Int
	_f1 bool
	_f2 _dafny.Map
	_f5 _dafny.Int
	_f6 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F3 = false
	_this.F4 = false
	_this._f0 = _dafny.Zero
	_this._f1 = false
	_this._f2 = _dafny.EmptyMap
	_this._f5 = _dafny.Zero
	_this._f6 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 bool, f2 _dafny.Map, f3 bool, f4 bool, f5 _dafny.Int, f6 _dafny.Sequence) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
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
func (_this *GlobalState) F2() _dafny.Map {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Sequence {
	{
		return _this._f6
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f7 _dafny.Int
	_f8 _dafny.Int
	F13 _dafny.Array
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.Zero
	_this.F13 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C0{}
var _ T0 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) F7() _dafny.Int {
	return _this._f7
}
func (_this *C0) F8() _dafny.Int {
	return _this._f8
}
func (_this *C0) Ctor__(f13 _dafny.Array, f7 _dafny.Int, f8 _dafny.Int) {
	{
		(_this).F13 = f13
		(_this)._f7 = f7
		(_this)._f8 = f8
	}
}
func (_this *C0) Fm0(p0 bool, globalState *GlobalState) _dafny.Set {
	{
		return _dafny.SetOf((_dafny.SetOf(false)).Union(_dafny.SetOf(false)), (_dafny.SetOf(false)).Difference(_dafny.SetOf(true, !(true))), _dafny.SetOf(true, false, false))
	}
}
func (_this *C0) Fm1(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("nup"))).Union(func() _dafny.Set {
			var _coll23 = _dafny.NewBuilder()
			_ = _coll23
			for _iter25 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(620))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}(func(_282_i2 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(539))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg25 _dafny.Int) interface{} {
						return coer25(arg25)
					}
				}(func(_283_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				}))
			}))).Elements()); ; {
				_compr_23, _ok25 := _iter25()
				if !_ok25 {
					break
				}
				var _284_v0 _dafny.Sequence
				_284_v0 = interface{}(_compr_23).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(620))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg26 _dafny.Int) interface{} {
						return coer26(arg26)
					}
				}(func(_282_i2 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(539))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg27 _dafny.Int) interface{} {
							return coer27(arg27)
						}
					}(func(_283_i3 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('h')
					}))
				})), _284_v0) {
					_coll23.Add(_284_v0)
				}
			}
			return _coll23.ToSet()
		}())).IsSubsetOf((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("jg"))).Union(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("q"), _dafny.UnicodeSeqOfUtf8Bytes("xbglft"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(739))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}(func(_285_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-819))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_286_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		})), _dafny.UnicodeSeqOfUtf8Bytes("gkesmfsv"))))
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f7 _dafny.Int
	_f8 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.Zero
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

func (_this *C1) F7() _dafny.Int {
	return _this._f7
}
func (_this *C1) F8() _dafny.Int {
	return _this._f8
}
func (_this *C1) Ctor__(f7 _dafny.Int, f8 _dafny.Int) {
	{
		(_this)._f7 = f7
		(_this)._f8 = f8
	}
}
func (_this *C1) Fm18(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jof"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hno"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}(func(_287_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		}))))
	}
}
func (_this *C1) Fm19(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	{
		return (Companion_D1_.Create_DC1_(_dafny.CodePoint('m'))).Dtor_cf1()
	}
}
func (_this *C1) M6(globalState *GlobalState) (_dafny.Set, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _288_v0 _dafny.Array
		_ = _288_v0
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_3
		var _nw55 _dafny.Array
		_ = _nw55
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw55 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = func(_289_i0 _dafny.Int) bool {
				return ((_this).F7()).Cmp((_this).F7()) > 0
			}
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw55 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw55).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw55).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_288_v0 = _nw55
		var _290_v1 bool
		_ = _290_v1
		_290_v1 = false
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
		_ = _index48
		(_288_v0).ArraySet1(_290_v1, (_index48).Int())
		var _291_v2 bool
		_ = _291_v2
		_291_v2 = _290_v1
		var _292_v3 _dafny.Sequence
		_ = _292_v3
		_292_v3 = _dafny.SeqOf(_290_v1, _290_v1, (func() bool {
			if _290_v1 {
				return _290_v1
			}
			return _290_v1
		})())
		var _293_v4 _dafny.Set
		_ = _293_v4
		_293_v4 = _dafny.SetOf(_dafny.IntOfInt64(309))
		var _294_v5 D2
		_ = _294_v5
		_294_v5 = Companion_D2_.Create_DC5_(_290_v1, (_this).F7(), (_293_v4).Cardinality(), (_this).F7(), _290_v1)
		var _pat_let_tv2 = _291_v2
		_ = _pat_let_tv2
		var _pat_let_tv3 = _291_v2
		_ = _pat_let_tv3
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
		_ = _index49
		var _rhs37 bool = _290_v1
		_ = _rhs37
		var _rhs38 bool = (_292_v3).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_292_v3).Cardinality()))).Uint32()).(bool)
		_ = _rhs38
		var _rhs39 bool = func(_source7 D2) bool {
			if _source7.Is_DC5() {
				var _295___mcc_h0 bool = _source7.Get_().(D2_DC5).Cf4
				_ = _295___mcc_h0
				var _296___mcc_h1 _dafny.Int = _source7.Get_().(D2_DC5).Cf5
				_ = _296___mcc_h1
				var _297___mcc_h2 _dafny.Int = _source7.Get_().(D2_DC5).Cf6
				_ = _297___mcc_h2
				var _298___mcc_h3 _dafny.Int = _source7.Get_().(D2_DC5).Cf7
				_ = _298___mcc_h3
				var _299___mcc_h4 bool = _source7.Get_().(D2_DC5).Cf8
				_ = _299___mcc_h4
				var _300_cf8 bool = _299___mcc_h4
				_ = _300_cf8
				var _301_cf7 _dafny.Int = _298___mcc_h3
				_ = _301_cf7
				var _302_cf6 _dafny.Int = _297___mcc_h2
				_ = _302_cf6
				var _303_cf5 _dafny.Int = _296___mcc_h1
				_ = _303_cf5
				var _304_cf4 bool = _295___mcc_h0
				_ = _304_cf4
				return _pat_let_tv2
			} else {
				var _305___mcc_h5 _dafny.MultiSet = _source7.Get_().(D2_DC4).Cf3
				_ = _305___mcc_h5
				var _306_cf3 _dafny.MultiSet = _305___mcc_h5
				_ = _306_cf3
				return _pat_let_tv3
			}
		}(_294_v5)
		_ = _rhs39
		var _rhs40 bool = (_291_v2)
		_ = _rhs40
		var _lhs38 *GlobalState = globalState
		_ = _lhs38
		var _lhs39 _dafny.Array = _288_v0
		_ = _lhs39
		var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
		_ = _lhs40
		var _lhs41 *GlobalState = globalState
		_ = _lhs41
		_lhs38.F4 = _rhs37
		(_lhs39).ArraySet1(_rhs38, (_lhs40).Int())
		_291_v2 = _rhs39
		_lhs41.F4 = _rhs40
		var _307_v6 _dafny.Map
		_ = _307_v6
		_307_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F7())
		var _308_v7 _dafny.Sequence
		_ = _308_v7
		_308_v7 = _dafny.UnicodeSeqOfUtf8Bytes("tuyj")
		var _hi1 _dafny.Int = _dafny.IntOfInt64(422)
		_ = _hi1
		for _309_i1 := (func() _dafny.Int {
			if (_307_v6).Contains(Companion_Default___.Fm20(_290_v1, (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), true, globalState)) {
				return (_307_v6).Get(Companion_Default___.Fm20(_290_v1, (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), true, globalState)).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_308_v7).Cardinality())
		})(); _309_i1.Cmp(_hi1) < 0; _309_i1 = _309_i1.Plus(_dafny.One) {
			var _310_v8 _dafny.Array
			_ = _310_v8
			var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw56
			_310_v8 = _nw56
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))
			_ = _index50
			(_310_v8).ArraySet1((_this).F7(), (_index50).Int())
			var _311_v9 _dafny.Array
			_ = _311_v9
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_4
			var _nw57 _dafny.Array
			_ = _nw57
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw57 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Sequence = (func(_312_v6 _dafny.Map) func(_dafny.Int) _dafny.Sequence {
					return func(_313_i2 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf((_dafny.Zero).Minus((_this).F7()), (_this).F7(), (func() _dafny.Int {
							if (_312_v6).Contains((_this).F8()) {
								return (_312_v6).Get((_this).F8()).(_dafny.Int)
							}
							return (_this).F8()
						})())
					}
				})(_307_v6)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw57 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw57).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw57).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_311_v9 = _nw57
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))
			_ = _index51
			var _rhs41 _dafny.Int = (_this).F8()
			_ = _rhs41
			var _rhs42 bool = (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)
			_ = _rhs42
			var _rhs43 _dafny.Array = _311_v9
			_ = _rhs43
			var _rhs44 _dafny.Int = (_294_v5).Dtor_cf7()
			_ = _rhs44
			var _lhs42 _dafny.Array = _310_v8
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))
			_ = _lhs43
			(_lhs42).ArraySet1(_rhs41, (_lhs43).Int())
			_290_v1 = _rhs42
			_311_v9 = _rhs43
			r2 = _rhs44
			var _314_v10 *C0
			_ = _314_v10
			var _nw58 *C0 = New_C0_()
			_ = _nw58
			_nw58.Ctor__(_310_v8, _309_i1, _309_i1)
			_314_v10 = _nw58
			_314_v10 = _314_v10
			var _315_v11 _dafny.CodePoint
			_ = _315_v11
			_315_v11 = _dafny.CodePoint('y')
			_315_v11 = _315_v11
			var _316_v12 _dafny.Map
			_ = _316_v12
			_316_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), _290_v1)
			var _317_v13 _dafny.Map
			_ = _317_v13
			_317_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_316_v12, (_this).F8())
			if (_317_v13).Contains(_316_v12) {
				(globalState).F3 = (_292_v3).Select((Companion_Default___.SafeIndex((_294_v5).Dtor_cf5(), _dafny.IntOfUint32((_292_v3).Cardinality()))).Uint32()).(bool)
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))
				_ = _index52
				var _rhs45 bool = _290_v1
				_ = _rhs45
				var _rhs46 _dafny.Int = (_310_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))).Int()).(_dafny.Int)
				_ = _rhs46
				var _lhs44 *GlobalState = globalState
				_ = _lhs44
				var _lhs45 _dafny.Array = _310_v8
				_ = _lhs45
				var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))
				_ = _lhs46
				_lhs44.F4 = _rhs45
				(_lhs45).ArraySet1(_rhs46, (_lhs46).Int())
				var _318_v14 _dafny.Map
				_ = _318_v14
				_318_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_291_v2, (_310_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))).Int()).(_dafny.Int))
				var _319_v15 _dafny.Set
				_ = _319_v15
				_319_v15 = _dafny.SetOf(_318_v14)
				var _320_v16 _dafny.Sequence
				_ = _320_v16
				_320_v16 = _dafny.SeqOf(_293_v4)
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))
				_ = _index53
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
				_ = _index54
				var _rhs47 _dafny.Int = (_310_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))).Int()).(_dafny.Int)
				_ = _rhs47
				var _rhs48 _dafny.Sequence = _308_v7
				_ = _rhs48
				var _rhs49 _dafny.Set = _319_v15
				_ = _rhs49
				var _rhs50 bool = _290_v1
				_ = _rhs50
				var _rhs51 _dafny.Set = (_320_v16).Select((Companion_Default___.SafeIndex(((_this).F8()).Plus(_dafny.IntOfInt64(310)), _dafny.IntOfUint32((_320_v16).Cardinality()))).Uint32()).(_dafny.Set)
				_ = _rhs51
				var _lhs47 _dafny.Array = _310_v8
				_ = _lhs47
				var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))
				_ = _lhs48
				var _lhs49 _dafny.Array = _288_v0
				_ = _lhs49
				var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
				_ = _lhs50
				(_lhs47).ArraySet1(_rhs47, (_lhs48).Int())
				_308_v7 = _rhs48
				_319_v15 = _rhs49
				(_lhs49).ArraySet1(_rhs50, (_lhs50).Int())
				_293_v4 = _rhs51
				(globalState).F3 = (_309_i1).Cmp((_this).F8()) <= 0
				var _321_v17 _dafny.Map
				_ = _321_v17
				_321_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_309_i1).Cmp((_this).F8()) >= 0, (_310_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))).Int()).(_dafny.Int))
				_321_v17 = _321_v17
			} else {
				_308_v7 = _dafny.Companion_Sequence_.Concatenate(_308_v7, _308_v7)
				var _322_v18 D1
				_ = _322_v18
				_322_v18 = Companion_D1_.Create_DC1_(_315_v11)
				var _323_v19 _dafny.Sequence
				_ = _323_v19
				_323_v19 = _dafny.SeqOf(_322_v18, _322_v18, Companion_D1_.Create_DC1_(_dafny.CodePoint('c')))
				var _324_v20 *C0
				_ = _324_v20
				var _nw59 *C0 = New_C0_()
				_ = _nw59
				_nw59.Ctor__(_310_v8, _dafny.IntOfUint32((Companion_Default___.Fm21(_dafny.IntOfInt64(715), _dafny.IntOfUint32((_323_v19).Cardinality()), (_this).F7(), _308_v7, globalState)).Cardinality()), (_this).F8())
				_324_v20 = _nw59
				var _325_v21 _dafny.Map
				_ = _325_v21
				_325_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_294_v5), _290_v1)
				var _326_v22 _dafny.Sequence
				_ = _326_v22
				_326_v22 = _dafny.SeqOf(_294_v5)
				_325_v21 = (_325_v21).Update(_326_v22, false)
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))
				_ = _index55
				(_310_v8).ArraySet1(Companion_Default___.SafeDivisionInt((_310_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_310_v8), 0))).Int()).(_dafny.Int), Companion_Default___.SafeDivisionInt(Companion_Default___.Fm20((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), (_290_v1), (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), globalState), _dafny.IntOfInt64(225))), (_index55).Int())
				var _327_v23 D1
				_ = _327_v23
				_327_v23 = Companion_D1_.Create_DC1_(_315_v11)
				var _328_v24 D1
				_ = _328_v24
				_328_v24 = Companion_D1_.Create_DC3_(_327_v23)
				var _329_v25 _dafny.MultiSet
				_ = _329_v25
				_329_v25 = _dafny.MultiSetOf(_328_v24, Companion_D1_.Create_DC3_(_327_v23), _328_v24, _328_v24)
				_329_v25 = Companion_Default___.Fm22(globalState)
			}
		}
		var _source8 D2 = Companion_D2_.Create_DC5_((_292_v3).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_292_v3).Cardinality()))).Uint32()).(bool), (_this).F7(), (_293_v4).Cardinality(), (_this).F7(), (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))
		_ = _source8
		if _source8.Is_DC5() {
			var _330___mcc_h6 bool = _source8.Get_().(D2_DC5).Cf4
			_ = _330___mcc_h6
			var _331___mcc_h7 _dafny.Int = _source8.Get_().(D2_DC5).Cf5
			_ = _331___mcc_h7
			var _332___mcc_h8 _dafny.Int = _source8.Get_().(D2_DC5).Cf6
			_ = _332___mcc_h8
			var _333___mcc_h9 _dafny.Int = _source8.Get_().(D2_DC5).Cf7
			_ = _333___mcc_h9
			var _334___mcc_h10 bool = _source8.Get_().(D2_DC5).Cf8
			_ = _334___mcc_h10
			var _335_cf8 bool = _334___mcc_h10
			_ = _335_cf8
			var _336_cf7 _dafny.Int = _333___mcc_h9
			_ = _336_cf7
			var _337_cf6 _dafny.Int = _332___mcc_h8
			_ = _337_cf6
			var _338_cf5 _dafny.Int = _331___mcc_h7
			_ = _338_cf5
			var _339_cf4 bool = _330___mcc_h6
			_ = _339_cf4
			var _340_v26 _dafny.CodePoint
			_ = _340_v26
			_340_v26 = _dafny.CodePoint('q')
			var _341_v27 _dafny.Map
			_ = _341_v27
			_341_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("xbuccmlho"), (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xbuccmlho")).Cardinality()))).Uint32(), _340_v26), (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))
			var _342_v28 _dafny.Map
			_ = _342_v28
			_342_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), !(_341_v27).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(787))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}((func(_343_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_344_i3 _dafny.Int) _dafny.CodePoint {
					return _343_v26
				}
			})(_340_v26))), _339_cf4)))
			_342_v28 = (_342_v28).Update(_339_cf4, _335_cf8)
			var _345_v29 D1
			_ = _345_v29
			_345_v29 = Companion_D1_.Create_DC3_(Companion_Default___.Fm23(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qf")).Cardinality()), globalState))
			var _346_v30 D1
			_ = _346_v30
			_346_v30 = Companion_D1_.Create_DC3_(_345_v29)
			var _347_v31 _dafny.Sequence
			_ = _347_v31
			_347_v31 = _dafny.SeqOf(_338_cf5)
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
			_ = _index56
			var _rhs52 D1 = _346_v30
			_ = _rhs52
			var _rhs53 bool = _290_v1
			_ = _rhs53
			var _rhs54 bool = _339_cf4
			_ = _rhs54
			var _rhs55 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-405))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}((func(_348_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_349_i4 _dafny.Int) _dafny.CodePoint {
					return _348_v26
				}
			})(_340_v26))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(602))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}((func(_350_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_351_i5 _dafny.Int) _dafny.CodePoint {
					return _350_v26
				}
			})(_340_v26))), _308_v7)), (Companion_Default___.SafeIndex((_336_cf7).Minus(_dafny.IntOfUint32((_347_v31).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-405))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}((func(_352_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_353_i4 _dafny.Int) _dafny.CodePoint {
					return _352_v26
				}
			})(_340_v26))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(602))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}((func(_354_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_355_i5 _dafny.Int) _dafny.CodePoint {
					return _354_v26
				}
			})(_340_v26))), _308_v7))).Cardinality()))).Uint32(), _340_v26)
			_ = _rhs55
			var _lhs51 _dafny.Array = _288_v0
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
			_ = _lhs52
			_346_v30 = _rhs52
			_339_cf4 = _rhs53
			(_lhs51).ArraySet1(_rhs54, (_lhs52).Int())
			_308_v7 = _rhs55
			_339_cf4 = Companion_Default___.Fm24((_336_cf7).Cmp((_294_v5).Dtor_cf5()) < 0, _338_cf5, _dafny.SeqOf(_dafny.CodePoint('c')), globalState)
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
			_ = _index57
			var _rhs56 _dafny.Int = (_this).F8()
			_ = _rhs56
			var _rhs57 bool = !(_339_cf4) || (_339_cf4)
			_ = _rhs57
			var _rhs58 bool = !(_290_v1)
			_ = _rhs58
			var _lhs53 _dafny.Array = _288_v0
			_ = _lhs53
			var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
			_ = _lhs54
			_338_cf5 = _rhs56
			(_lhs53).ArraySet1(_rhs57, (_lhs54).Int())
			_335_cf8 = _rhs58
		} else {
			var _356___mcc_h11 _dafny.MultiSet = _source8.Get_().(D2_DC4).Cf3
			_ = _356___mcc_h11
			var _357_cf3 _dafny.MultiSet = _356___mcc_h11
			_ = _357_cf3
			var _358_v32 _dafny.Array
			_ = _358_v32
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_5
			var _nw60 _dafny.Array
			_ = _nw60
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw60 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Int = func(_359_i6 _dafny.Int) _dafny.Int {
					return (_359_i6).Times((_this).F7())
				}
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
			_358_v32 = _nw60
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))
			_ = _index58
			(_358_v32).ArraySet1((_this).F8(), (_index58).Int())
			var _360_v33 _dafny.Set
			_ = _360_v33
			_360_v33 = _dafny.SetOf((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))
			var _361_v34 *C0
			_ = _361_v34
			var _nw61 *C0 = New_C0_()
			_ = _nw61
			_nw61.Ctor__(_358_v32, _dafny.IntOfUint32((_308_v7).Cardinality()), (_this).F7())
			_361_v34 = _nw61
			var _362_v35 _dafny.Map
			_ = _362_v35
			_362_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v1, (func() *C0 {
				if false {
					return _361_v34
				}
				return _361_v34
			})())
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
			_ = _index59
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))
			_ = _index60
			var _rhs59 bool = !(!((!(((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)) || (!((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))))) && (_290_v1)))
			_ = _rhs59
			var _rhs60 _dafny.Set = (func() _dafny.Set {
				if _290_v1 {
					return (_360_v33).Intersection(_360_v33)
				}
				return _360_v33
			})()
			_ = _rhs60
			var _rhs61 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_362_v35).Cardinality()))
			_ = _rhs61
			var _rhs62 _dafny.Int = ((_this).F7()).Plus((_this).F7())
			_ = _rhs62
			var _lhs55 _dafny.Array = _288_v0
			_ = _lhs55
			var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
			_ = _lhs56
			var _lhs57 _dafny.Array = _358_v32
			_ = _lhs57
			var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))
			_ = _lhs58
			(_lhs55).ArraySet1(_rhs59, (_lhs56).Int())
			r0 = _rhs60
			(_lhs57).ArraySet1(_rhs61, (_lhs58).Int())
			r2 = _rhs62
			if ((_dafny.Zero).Minus((_this).F7())).Cmp(((_this).F7()).Times((_dafny.Zero).Minus((_this).F8()))) != 0 {
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))
				_ = _index61
				(_358_v32).ArraySet1(_dafny.IntOfInt64(117), (_index61).Int())
				var _363_v36 _dafny.Sequence
				_ = _363_v36
				_363_v36 = _dafny.SeqOf(_357_cf3)
				var _364_v37 _dafny.Array
				_ = _364_v37
				var _nwElement0_10 _dafny.MultiSet = ((_357_cf3).Update((_292_v3).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_292_v3).Cardinality()))).Uint32()).(bool), Companion_Default___.Abs((_358_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))).Int()).(_dafny.Int)))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(!((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)), _290_v1)))
				_ = _nwElement0_10
				var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(11))
				_ = _nw62
				(_nw62).ArraySet1(_nwElement0_10, 0)
				(_nw62).ArraySet1(_357_cf3, 1)
				(_nw62).ArraySet1(_357_cf3, 2)
				(_nw62).ArraySet1(_dafny.MultiSetFromSeq(_292_v3), 3)
				(_nw62).ArraySet1(_357_cf3, 4)
				(_nw62).ArraySet1(_357_cf3, 5)
				(_nw62).ArraySet1(_357_cf3, 6)
				(_nw62).ArraySet1((_357_cf3).Union(_357_cf3), 7)
				(_nw62).ArraySet1(((_363_v36).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F8()), _dafny.IntOfUint32((_363_v36).Cardinality()))).Uint32()).(_dafny.MultiSet)).Union(_357_cf3), 8)
				(_nw62).ArraySet1(_357_cf3, 9)
				(_nw62).ArraySet1(_357_cf3, 10)
				_364_v37 = _nw62
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_364_v37), 0))
				_ = _index62
				(_364_v37).ArraySet1(_357_cf3, (_index62).Int())
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_364_v37), 0))
				_ = _index63
				(_364_v37).ArraySet1(((_357_cf3).Intersection(_dafny.MultiSetOf(_290_v1))).Difference(_357_cf3), (_index63).Int())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))
				_ = _index64
				(_358_v32).ArraySet1(((_this).F8()).Times((func() _dafny.Int {
					if _290_v1 {
						return (_dafny.Zero).Minus((_358_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))).Int()).(_dafny.Int))
					}
					return (_358_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))).Int()).(_dafny.Int)
				})()), (_index64).Int())
				var _365_v38 _dafny.CodePoint
				_ = _365_v38
				_365_v38 = _dafny.CodePoint('u')
				var _366_v39 _dafny.Map
				_ = _366_v39
				_366_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_365_v38, (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))
				_366_v39 = (_366_v39).Update(_365_v38, (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))
				_ = _index65
				(_358_v32).ArraySet1((_358_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))).Int()).(_dafny.Int), (_index65).Int())
			} else {
				r1 = (func() _dafny.Int {
					if (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool) {
						return (func() _dafny.Int {
							if _290_v1 {
								return (_dafny.Zero).Minus((_this).F8())
							}
							return (_this).F8()
						})()
					}
					return (_358_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))).Int()).(_dafny.Int)
				})()
				var _367_v40 _dafny.Sequence
				_ = _367_v40
				_367_v40 = _dafny.SeqOf(Companion_D1_.Create_DC2_())
				_367_v40 = _367_v40
				var _368_v41 D2
				_ = _368_v41
				_368_v41 = Companion_D2_.Create_DC4_(_357_cf3)
				_368_v41 = _368_v41
				var _369_v42 *C0
				_ = _369_v42
				var _nw63 *C0 = New_C0_()
				_ = _nw63
				_nw63.Ctor__(_358_v32, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-810), (_358_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))).Int()).(_dafny.Int)), ((_358_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfUint32((_308_v7).Cardinality())))
				_369_v42 = _nw63
				var _370_v43 _dafny.Array
				_ = _370_v43
				var _nw64 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(14))
				_ = _nw64
				_370_v43 = _nw64
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_370_v43), 0))
				_ = _index66
				(_370_v43).ArraySet1(_361_v34, (_index66).Int())
				var _371_v44 _dafny.Map
				_ = _371_v44
				_371_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _290_v1)
				var _372_v45 _dafny.MultiSet
				_ = _372_v45
				_372_v45 = _dafny.MultiSetOf(Companion_Default___.Fm20((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), _290_v1, (func() bool {
					if (_371_v44).Contains((_this).F7()) {
						return (_371_v44).Get((_this).F7()).(bool)
					}
					return (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)
				})(), globalState), (_360_v33).Cardinality(), (_this).F8())
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_370_v43), 0))
				_ = _index67
				var _rhs63 *C0 = _369_v42
				_ = _rhs63
				var _rhs64 _dafny.Int = (_372_v45).Cardinality()
				_ = _rhs64
				var _lhs59 _dafny.Array = _370_v43
				_ = _lhs59
				var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_370_v43), 0))
				_ = _lhs60
				(_lhs59).ArraySet1(_rhs63, (_lhs60).Int())
				r1 = _rhs64
			}
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
			_ = _index68
			(_288_v0).ArraySet1((_292_v3).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if _290_v1 {
					return (_358_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))).Int()).(_dafny.Int)
				}
				return (_358_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))).Int()).(_dafny.Int)
			})(), _dafny.IntOfUint32((_292_v3).Cardinality()))).Uint32()).(bool), (_index68).Int())
			var _373_v46 _dafny.Set
			_ = _373_v46
			_373_v46 = _dafny.SetOf(_291_v2)
			var _374_v47 _dafny.Map
			_ = _374_v47
			_374_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_373_v46, Companion_D2_.Create_DC5_((_361_v34).Fm1(_360_v33, (_358_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))).Int()).(_dafny.Int), (_358_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))).Int()).(_dafny.Int), (_this).F7(), globalState), Companion_Default___.Fm20(_290_v1, (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), globalState), (_this).F8(), _dafny.IntOfInt64(-418), _290_v1))
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_358_v32), 0))
			_ = _index69
			(_358_v32).ArraySet1((_374_v47).Cardinality(), (_index69).Int())
		}
		var _375_v48 _dafny.CodePoint
		_ = _375_v48
		_375_v48 = _dafny.CodePoint('x')
		var _376_v49 D1
		_ = _376_v49
		_376_v49 = Companion_D1_.Create_DC1_(_375_v48)
		var _377_v50 D1
		_ = _377_v50
		_377_v50 = Companion_D1_.Create_DC3_(_376_v49)
		var _source9 D1 = _377_v50
		_ = _source9
		if _source9.Is_DC2() {
			var _378_v51 _dafny.Set
			_ = _378_v51
			_378_v51 = _dafny.SetOf(_288_v0)
			if ((_378_v51).IsSubsetOf(_378_v51)) && (_290_v1) {
				(globalState).F4 = _290_v1
				var _379_v52 D1
				_ = _379_v52
				_379_v52 = Companion_D1_.Create_DC1_((func() _dafny.CodePoint {
					if true {
						return _375_v48
					}
					return (_this).Fm19(_290_v1, (_294_v5).Dtor_cf8(), (_this).F8(), (_this).F7(), globalState)
				})())
				_379_v52 = _379_v52
				var _380_v53 _dafny.Array
				_ = _380_v53
				var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
				_ = _nw65
				_380_v53 = _nw65
				var _381_v54 _dafny.Map
				_ = _381_v54
				_381_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _288_v0)
				var _382_v55 *C0
				_ = _382_v55
				var _nw66 *C0 = New_C0_()
				_ = _nw66
				_nw66.Ctor__(_380_v53, (_dafny.Zero).Minus((_this).F7()), (_381_v54).Cardinality())
				_382_v55 = _nw66
				var _383_v56 _dafny.Map
				_ = _383_v56
				_383_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), (_dafny.IntOfInt64(299)).Cmp(Companion_Default___.Fm20(_290_v1, (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), globalState)) <= 0)
				_383_v56 = Companion_Default___.Fm25(globalState)
				r2 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if _290_v1 {
						return (_this).F8()
					}
					return (_this).F7()
				})(), (_this).F8())
			} else {
				var _384_v57 _dafny.Map
				_ = _384_v57
				_384_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F8()).Times(Companion_Default___.Fm20(_290_v1, _290_v1, (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), globalState)), (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))
				_384_v57 = (_384_v57).Update((_this).F7(), (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))
				var _385_v58 _dafny.Array
				_ = _385_v58
				var _nw67 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw67
				_385_v58 = _nw67
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_385_v58), 0))
				_ = _index70
				(_385_v58).ArraySet1((_this).F8(), (_index70).Int())
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_385_v58), 0))
				_ = _index71
				(_385_v58).ArraySet1((_this).F7(), (_index71).Int())
				var _386_v59 _dafny.Sequence
				_ = _386_v59
				_386_v59 = _dafny.SeqOf((_385_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_385_v58), 0))).Int()).(_dafny.Int))
				var _387_v60 _dafny.MultiSet
				_ = _387_v60
				_387_v60 = _dafny.MultiSetOf((_385_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_385_v58), 0))).Int()).(_dafny.Int))
				var _388_v61 _dafny.Map
				_ = _388_v61
				_388_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_386_v59, (Companion_Default___.SafeIndex(Companion_Default___.Fm20(Companion_Default___.Fm24(_290_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("trxrrblv")).Cardinality()), _dafny.SeqOf(_375_v48, _375_v48), globalState), (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), _290_v1, globalState), _dafny.IntOfUint32((_386_v59).Cardinality()))).Uint32(), (_this).F8()), (_293_v4).IsProperSubsetOf((Companion_Default___.Fm26(_387_v60, (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), globalState)).Dtor_cf9()))
				_388_v61 = _388_v61
				r2 = (_this).F8()
				_290_v1 = (_292_v3).Select((Companion_Default___.SafeIndex((_385_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_385_v58), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_292_v3).Cardinality()))).Uint32()).(bool)
			}
			r1 = (_294_v5).Dtor_cf5()
			(globalState).F4 = !(_290_v1)
			var _389_v62 _dafny.Sequence
			_ = _389_v62
			_389_v62 = _dafny.SeqOf((_this).F7(), (_this).F8())
			r2 = (_dafny.IntOfInt64(-476)).Minus((_389_v62).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm20((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), _290_v1, false, globalState), _dafny.IntOfUint32((_389_v62).Cardinality()))).Uint32()).(_dafny.Int))
		} else if _source9.Is_DC1() {
			var _390___mcc_h12 _dafny.CodePoint = _source9.Get_().(D1_DC1).Cf1
			_ = _390___mcc_h12
			var _391_cf1 _dafny.CodePoint = _390___mcc_h12
			_ = _391_cf1
			var _392_v63 _dafny.Array
			_ = _392_v63
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_6
			var _nw68 _dafny.Array
			_ = _nw68
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw68 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Int = func(_393_i7 _dafny.Int) _dafny.Int {
					return (_393_i7).Minus((_this).F7())
				}
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw68 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw68).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw68).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_392_v63 = _nw68
			var _394_v64 _dafny.Map
			_ = _394_v64
			_394_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v1, (_dafny.Zero).Minus((_this).F7()))
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_392_v63), 0))
			_ = _index72
			(_392_v63).ArraySet1(((_394_v64).Cardinality()).Times((_294_v5).Dtor_cf7()), (_index72).Int())
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_392_v63), 0))
			_ = _index73
			(_392_v63).ArraySet1(_dafny.IntOfInt64(-885), (_index73).Int())
			var _395_v65 _dafny.MultiSet
			_ = _395_v65
			_395_v65 = _dafny.MultiSetOf((_392_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_392_v63), 0))).Int()).(_dafny.Int), (_this).F7())
			(globalState).F4 = (_395_v65).Equals(_395_v65)
			var _396_v66 D3
			_ = _396_v66
			_396_v66 = Companion_D3_.Create_DC7_((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))
			var _397_v67 _dafny.Map
			_ = _397_v67
			_397_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v1, _396_v66)
			_397_v67 = (_397_v67).Update((func() bool {
				if _290_v1 {
					return (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)
				}
				return (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)
			})(), _396_v66)
			if (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool) {
				(globalState).F4 = (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)
				_391_cf1 = _391_cf1
				(globalState).F4 = ((_dafny.IntOfInt64(957)).Minus(_dafny.IntOfInt64(350))).Cmp((_this).F8()) == 0
				var _398_v68 _dafny.Array
				_ = _398_v68
				var _nwElement0_11 _dafny.Array = _392_v63
				_ = _nwElement0_11
				var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(17))
				_ = _nw69
				(_nw69).ArraySet1(_nwElement0_11, 0)
				(_nw69).ArraySet1(_392_v63, 1)
				(_nw69).ArraySet1(_392_v63, 2)
				(_nw69).ArraySet1(_392_v63, 3)
				(_nw69).ArraySet1(_392_v63, 4)
				(_nw69).ArraySet1(_392_v63, 5)
				(_nw69).ArraySet1(_392_v63, 6)
				(_nw69).ArraySet1(_392_v63, 7)
				(_nw69).ArraySet1(_392_v63, 8)
				(_nw69).ArraySet1(_392_v63, 9)
				(_nw69).ArraySet1(_392_v63, 10)
				(_nw69).ArraySet1(_392_v63, 11)
				(_nw69).ArraySet1(_392_v63, 12)
				(_nw69).ArraySet1(_392_v63, 13)
				(_nw69).ArraySet1(_392_v63, 14)
				(_nw69).ArraySet1(_392_v63, 15)
				(_nw69).ArraySet1(_392_v63, 16)
				_398_v68 = _nw69
				var _399_v69 _dafny.Sequence
				_ = _399_v69
				_399_v69 = _dafny.SeqOf(_398_v68, _398_v68)
				var _400_v70 _dafny.Map
				_ = _400_v70
				_400_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), _290_v1)
				var _401_v71 _dafny.Array
				_ = _401_v71
				var _nwElement0_12 _dafny.Array = _398_v68
				_ = _nwElement0_12
				var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(4))
				_ = _nw70
				(_nw70).ArraySet1(_nwElement0_12, 0)
				(_nw70).ArraySet1((_399_v69).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_399_v69).Cardinality()))).Uint32()).(_dafny.Array), 1)
				(_nw70).ArraySet1(_398_v68, 2)
				(_nw70).ArraySet1((_399_v69).Select((Companion_Default___.SafeIndex((_400_v70).Cardinality(), _dafny.IntOfUint32((_399_v69).Cardinality()))).Uint32()).(_dafny.Array), 3)
				_401_v71 = _nw70
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_401_v71), 0))
				_ = _index74
				(_401_v71).ArraySet1(_398_v68, (_index74).Int())
				var _402_v72 D4
				_ = _402_v72
				_402_v72 = Companion_D4_.Create_DC8_(_398_v68)
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_401_v71), 0))
				_ = _index75
				(_401_v71).ArraySet1((_402_v72).Dtor_cf11(), (_index75).Int())
				var _403_v73 _dafny.Map
				_ = _403_v73
				_403_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_392_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_392_v63), 0))).Int()).(_dafny.Int), true)
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_392_v63), 0))
				_ = _index76
				(_392_v63).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.IntOfUint32((_dafny.SeqOf(_290_v1)).Cardinality())).Minus(Companion_Default___.Fm20(_290_v1, _290_v1, (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), globalState)), Companion_Default___.SafeDivisionInt((_this).F8(), (_403_v73).Cardinality())), (_index76).Int())
			} else {
				var _404_v74 _dafny.Array
				_ = _404_v74
				var _nwElement0_13 _dafny.Sequence = _308_v7
				_ = _nwElement0_13
				var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(7))
				_ = _nw71
				(_nw71).ArraySet1(_nwElement0_13, 0)
				(_nw71).ArraySet1(_308_v7, 1)
				(_nw71).ArraySet1((_this).Fm18(globalState), 2)
				(_nw71).ArraySet1(_308_v7, 3)
				(_nw71).ArraySet1(_308_v7, 4)
				(_nw71).ArraySet1((_this).Fm18(globalState), 5)
				(_nw71).ArraySet1(_308_v7, 6)
				_404_v74 = _nw71
				var _405_v75 _dafny.Map
				_ = _405_v75
				_405_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v1, _404_v74)
				_405_v75 = (_405_v75).Update(_290_v1, _404_v74)
				var _406_v76 _dafny.Sequence
				_ = _406_v76
				_406_v76 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_292_v3))
				var _407_v78 _dafny.Map
				_ = _407_v78
				_407_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v1, (_294_v5).Dtor_cf8())
				var _408_v79 _dafny.MultiSet
				_ = _408_v79
				_408_v79 = _dafny.MultiSetOf((func() bool {
					if (_407_v78).Contains((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)) {
						return (_407_v78).Get((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)).(bool)
					}
					return (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)
				})())
				var _409_v80 _dafny.Set
				_ = _409_v80
				_409_v80 = _dafny.SetOf(_408_v79)
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
				_ = _index77
				(_288_v0).ArraySet1(!((func() _dafny.Set {
					if Companion_Default___.Fm24(_290_v1, (_this).F8(), _308_v7, globalState) {
						return func() _dafny.Set {
							var _coll24 = _dafny.NewBuilder()
							_ = _coll24
							for _iter26 := _dafny.Iterate((_406_v76).Elements()); ; {
								_compr_24, _ok26 := _iter26()
								if !_ok26 {
									break
								}
								var _410_v77 _dafny.MultiSet
								_410_v77 = interface{}(_compr_24).(_dafny.MultiSet)
								if _dafny.Companion_Sequence_.Contains(_406_v76, _410_v77) {
									_coll24.Add(_410_v77)
								}
							}
							return _coll24.ToSet()
						}()
					}
					return Companion_Default___.Fm27(globalState)
				})()).Equals(_409_v80), (_index77).Int())
				(globalState).F3 = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_292_v3, _292_v3), _292_v3)
				var _rhs65 bool = (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)
				_ = _rhs65
				var _rhs66 _dafny.Int = (_this).F8()
				_ = _rhs66
				var _lhs61 *GlobalState = globalState
				_ = _lhs61
				_lhs61.F4 = _rhs65
				r1 = _rhs66
				_307_v6 = ((_307_v6).Update((_392_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_392_v63), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus((_this).F7()))).Merge(_307_v6)
			}
		} else {
			var _411___mcc_h13 D1 = _source9.Get_().(D1_DC3).Cf2
			_ = _411___mcc_h13
			var _412_cf2 D1 = _411___mcc_h13
			_ = _412_cf2
			var _413_v81 _dafny.MultiSet
			_ = _413_v81
			_413_v81 = _dafny.MultiSetOf((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))
			var _414_v82 _dafny.Sequence
			_ = _414_v82
			_414_v82 = _dafny.SeqOf(_413_v81, _413_v81, _dafny.MultiSetOf(!(false), false))
			var _415_v83 _dafny.MultiSet
			_ = _415_v83
			_415_v83 = _dafny.MultiSetOf((_this).F7())
			var _416_v84 _dafny.Array
			_ = _416_v84
			var _nwElement0_14 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(970))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}(func(_417_i8 _dafny.Int) _dafny.Int {
				return (_this).F7()
			}))).Cardinality())
			_ = _nwElement0_14
			var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(24))
			_ = _nw72
			(_nw72).ArraySet1(_nwElement0_14, 0)
			(_nw72).ArraySet1((_this).F8(), 1)
			(_nw72).ArraySet1((_this).F7(), 2)
			(_nw72).ArraySet1((_this).F7(), 3)
			(_nw72).ArraySet1((_this).F8(), 4)
			(_nw72).ArraySet1(_dafny.IntOfInt64(487), 5)
			(_nw72).ArraySet1((_this).F8(), 6)
			(_nw72).ArraySet1((_this).F8(), 7)
			(_nw72).ArraySet1(_dafny.IntOfUint32((_308_v7).Cardinality()), 8)
			(_nw72).ArraySet1((_this).F7(), 9)
			(_nw72).ArraySet1((_this).F7(), 10)
			(_nw72).ArraySet1(_dafny.IntOfInt64(912), 11)
			(_nw72).ArraySet1((_dafny.SetOf((_this).F8(), (_294_v5).Dtor_cf5(), Companion_Default___.Fm20(_290_v1, _290_v1, (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), globalState))).Cardinality(), 12)
			(_nw72).ArraySet1((_this).F7(), 13)
			(_nw72).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_414_v82, (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_414_v82).Cardinality()))).Uint32(), _413_v81)).Cardinality()), 14)
			(_nw72).ArraySet1((_this).F7(), 15)
			(_nw72).ArraySet1((_307_v6).Cardinality(), 16)
			(_nw72).ArraySet1((_this).F7(), 17)
			(_nw72).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_308_v7, (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_308_v7).Cardinality()))).Uint32(), _375_v48)).Cardinality()), 18)
			(_nw72).ArraySet1((_415_v83).Cardinality(), 19)
			(_nw72).ArraySet1(_dafny.IntOfInt64(609), 20)
			(_nw72).ArraySet1((_this).F8(), 21)
			(_nw72).ArraySet1((_this).F8(), 22)
			(_nw72).ArraySet1((_this).F8(), 23)
			_416_v84 = _nw72
			var _418_v85 _dafny.Map
			_ = _418_v85
			_418_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v1, _dafny.CodePoint('l'))
			var _419_v86 *C0
			_ = _419_v86
			var _nw73 *C0 = New_C0_()
			_ = _nw73
			_nw73.Ctor__(_416_v84, (func() _dafny.Int {
				if (_415_v83).Contains((_this).F7()) {
					return (_415_v83).Multiplicity((_this).F7())
				}
				return (_418_v85).Cardinality()
			})(), (_this).F8())
			_419_v86 = _nw73
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
			_ = _index78
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
			_ = _index79
			var _rhs67 bool = Companion_Default___.Fm24((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), (_this).F8(), _308_v7, globalState)
			_ = _rhs67
			var _rhs68 bool = _290_v1
			_ = _rhs68
			var _rhs69 bool = (Companion_Default___.Fm20((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), true, globalState)).Cmp((_this).F7()) == 0
			_ = _rhs69
			var _rhs70 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(90))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}((func(_420_v48 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_421_i9 _dafny.Int) _dafny.CodePoint {
					return _420_v48
				}
			})(_375_v48)))
			_ = _rhs70
			var _lhs62 _dafny.Array = _288_v0
			_ = _lhs62
			var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
			_ = _lhs63
			var _lhs64 _dafny.Array = _288_v0
			_ = _lhs64
			var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
			_ = _lhs65
			var _lhs66 *GlobalState = globalState
			_ = _lhs66
			(_lhs62).ArraySet1(_rhs67, (_lhs63).Int())
			(_lhs64).ArraySet1(_rhs68, (_lhs65).Int())
			_lhs66.F3 = _rhs69
			_308_v7 = _rhs70
			var _422_v87 _dafny.Array
			_ = _422_v87
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_7
			var _nw74 _dafny.Array
			_ = _nw74
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw74 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Sequence = (func(_423_v1 bool, _424_v0 _dafny.Array, _425_v83 _dafny.MultiSet) func(_dafny.Int) _dafny.Sequence {
					return func(_426_i10 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf((_dafny.SetOf(Companion_D1_.Create_DC2_(), Companion_D1_.Create_DC2_(), Companion_D1_.Create_DC2_(), Companion_D1_.Create_DC2_(), Companion_D1_.Create_DC2_())).Cardinality(), (_this).F8(), _dafny.IntOfUint32((_dafny.SeqOf(_423_v1, (_424_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_424_v0), 0))).Int()).(bool), _423_v1)).Cardinality()), (_this).F8(), (_425_v83).Cardinality())
					}
				})(_290_v1, _288_v0, _415_v83)
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
			_422_v87 = _nw74
			var _427_v88 _dafny.Sequence
			_ = _427_v88
			_427_v88 = _dafny.SeqOf(_288_v0)
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_422_v87), 0))
			_ = _index80
			(_422_v87).ArraySet1(Companion_Default___.Fm21(Companion_Default___.Fm20((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), true, globalState), (_this).F7(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_427_v88, (Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_427_v88).Cardinality()))).Uint32(), _288_v0)).Cardinality()), _308_v7, globalState), (_index80).Int())
			var _428_v89 _dafny.Map
			_ = _428_v89
			_428_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), (_this).F8())
			var _429_v90 _dafny.Sequence
			_ = _429_v90
			_429_v90 = _dafny.SeqOf((_dafny.Zero).Minus((_428_v89).Cardinality()))
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_422_v87), 0))
			_ = _index81
			(_422_v87).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_292_v3).Cardinality()), (_this).F7()), _429_v90), (_index81).Int())
			if _290_v1 {
				var _430_v91 D2
				_ = _430_v91
				_430_v91 = Companion_D2_.Create_DC4_((Companion_Default___.Fm28((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), globalState)).Difference(_dafny.MultiSetOf((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))))
				_430_v91 = _430_v91
				var _arr7 _dafny.Array = _419_v86.F13
				_ = _arr7
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_419_v86.F13), 0))
				_ = _index82
				(_arr7).ArraySet1((_dafny.Zero).Minus((_this).F8()), (_index82).Int())
				var _arr8 _dafny.Array = _419_v86.F13
				_ = _arr8
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_419_v86.F13), 0))
				_ = _index83
				(_arr8).ArraySet1((_dafny.SetOf(_290_v1, (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))).Cardinality(), (_index83).Int())
				var _431_v92 _dafny.MultiSet
				_ = _431_v92
				_431_v92 = _dafny.MultiSetOf(_307_v6, _307_v6, _307_v6)
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
				_ = _index84
				(_288_v0).ArraySet1((_431_v92).IsProperSubsetOf(_431_v92), (_index84).Int())
				_413_v81 = _413_v81
				var _432_v93 _dafny.Map
				_ = _432_v93
				_432_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_419_v86.F13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_419_v86.F13), 0))).Int()).(_dafny.Int), _290_v1)
				var _433_v94 _dafny.Array
				_ = _433_v94
				var _nw75 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
				_ = _nw75
				_433_v94 = _nw75
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_433_v94), 0))
				_ = _index85
				(_433_v94).ArraySet1(_419_v86.F13, (_index85).Int())
				var _arr9 _dafny.Array = _419_v86.F13
				_ = _arr9
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_419_v86.F13), 0))
				_ = _index86
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_433_v94), 0))
				_ = _index87
				var _rhs71 _dafny.Int = (func() _dafny.Int {
					if (_428_v89).Contains(((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)) && (_290_v1)) {
						return (_428_v89).Get(((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool)) && (_290_v1)).(_dafny.Int)
					}
					return Companion_Default___.SafeModuloInt((_this).F7(), (_419_v86.F13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_419_v86.F13), 0))).Int()).(_dafny.Int))
				})()
				_ = _rhs71
				var _rhs72 _dafny.Map = _432_v93
				_ = _rhs72
				var _rhs73 _dafny.Int = (_this).F7()
				_ = _rhs73
				var _rhs74 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F8(), (_this).F8())
				_ = _rhs74
				var _rhs75 _dafny.Array = _419_v86.F13
				_ = _rhs75
				var _lhs67 _dafny.Array = _419_v86.F13
				_ = _lhs67
				var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_419_v86.F13), 0))
				_ = _lhs68
				var _lhs69 _dafny.Array = _433_v94
				_ = _lhs69
				var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_433_v94), 0))
				_ = _lhs70
				(_lhs67).ArraySet1(_rhs71, (_lhs68).Int())
				_432_v93 = _rhs72
				r1 = _rhs73
				r1 = _rhs74
				(_lhs69).ArraySet1(_rhs75, (_lhs70).Int())
			} else {
				r2 = (Companion_Default___.SafeModuloInt((_this).F8(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_419_v86.F13, _290_v1)).Cardinality())).Plus((_this).F8())
				var _434_v95 _dafny.Map
				_ = _434_v95
				_434_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool), ((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update((_422_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_422_v87), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(((Companion_D5_.Create_DC11_(_428_v89)).Dtor_cf18()).Cardinality(), _dafny.IntOfUint32(((_422_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_422_v87), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), (_this).F7()))).Cardinality()).Cmp((_dafny.SetOf((_this).F7())).Cardinality()) <= 0)
				var _435_v96 _dafny.Sequence
				_ = _435_v96
				_435_v96 = _dafny.SeqOf(_434_v95, _434_v95, _434_v95, (_434_v95).Merge(_434_v95))
				_434_v95 = (_435_v96).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_435_v96).Cardinality()))).Uint32()).(_dafny.Map)
				var _436_v97 _dafny.Array
				_ = _436_v97
				var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
				_ = _nw76
				_436_v97 = _nw76
				var _437_v98 _dafny.MultiSet
				_ = _437_v98
				_437_v98 = _dafny.MultiSetOf(_375_v48)
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
				_ = _index88
				var _rhs76 bool = (_437_v98).IsProperSubsetOf(_437_v98)
				_ = _rhs76
				var _rhs77 _dafny.Array = _436_v97
				_ = _rhs77
				var _rhs78 bool = _290_v1
				_ = _rhs78
				var _lhs71 *GlobalState = globalState
				_ = _lhs71
				var _lhs72 _dafny.Array = _288_v0
				_ = _lhs72
				var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))
				_ = _lhs73
				_lhs71.F4 = _rhs76
				_436_v97 = _rhs77
				(_lhs72).ArraySet1(_rhs78, (_lhs73).Int())
				r2 = (_this).F8()
				r1 = (_this).F8()
			}
		}
		var _438_v99 D4
		_ = _438_v99
		_438_v99 = Companion_D4_.Create_DC10_((_this).F7(), (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool))
		var _pat_let_tv4 = _308_v7
		_ = _pat_let_tv4
		r2 = _dafny.IntOfUint32((func(_source10 D4) _dafny.Sequence {
			if _source10.Is_DC9() {
				var _439___mcc_h14 _dafny.Int = _source10.Get_().(D4_DC9).Cf12
				_ = _439___mcc_h14
				var _440___mcc_h15 _dafny.Sequence = _source10.Get_().(D4_DC9).Cf13
				_ = _440___mcc_h15
				var _441___mcc_h16 _dafny.Int = _source10.Get_().(D4_DC9).Cf14
				_ = _441___mcc_h16
				var _442___mcc_h17 _dafny.Int = _source10.Get_().(D4_DC9).Cf15
				_ = _442___mcc_h17
				var _443_cf15 _dafny.Int = _442___mcc_h17
				_ = _443_cf15
				var _444_cf14 _dafny.Int = _441___mcc_h16
				_ = _444_cf14
				var _445_cf13 _dafny.Sequence = _440___mcc_h15
				_ = _445_cf13
				var _446_cf12 _dafny.Int = _439___mcc_h14
				_ = _446_cf12
				return _dafny.UnicodeSeqOfUtf8Bytes("cradsw")
			} else if _source10.Is_DC10() {
				var _447___mcc_h18 _dafny.Int = _source10.Get_().(D4_DC10).Cf16
				_ = _447___mcc_h18
				var _448___mcc_h19 bool = _source10.Get_().(D4_DC10).Cf17
				_ = _448___mcc_h19
				var _449_cf17 bool = _448___mcc_h19
				_ = _449_cf17
				var _450_cf16 _dafny.Int = _447___mcc_h18
				_ = _450_cf16
				return _dafny.UnicodeSeqOfUtf8Bytes("bqqe")
			} else {
				var _451___mcc_h20 _dafny.Array = _source10.Get_().(D4_DC8).Cf11
				_ = _451___mcc_h20
				var _452_cf11 _dafny.Array = _451___mcc_h20
				_ = _452_cf11
				return _pat_let_tv4
			}
		}(_438_v99)).Cardinality())
		(globalState).F4 = ((func() bool {
			if (_288_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_288_v0), 0))).Int()).(bool) {
				return _290_v1
			}
			return (_292_v3).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_292_v3).Cardinality()))).Uint32()).(bool)
		})()) && (_290_v1)
		var _453_v100 _dafny.Set
		_ = _453_v100
		_453_v100 = _dafny.SetOf(_290_v1)
		r0 = (_453_v100).Intersection((_453_v100).Intersection(_453_v100))
		r1 = (_this).F8()
		var _454_v101 _dafny.MultiSet
		_ = _454_v101
		_454_v101 = _dafny.MultiSetOf(_288_v0)
		r2 = ((_this).F8()).Plus((func() _dafny.Int {
			if false {
				return (_454_v101).Cardinality()
			}
			return (_this).F7()
		})())
		return r0, r1, r2
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f7 _dafny.Int
	_f8 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.Zero
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

func (_this *C2) F7() _dafny.Int {
	return _this._f7
}
func (_this *C2) F8() _dafny.Int {
	return _this._f8
}
func (_this *C2) Ctor__(f7 _dafny.Int, f8 _dafny.Int) {
	{
		(_this)._f7 = f7
		(_this)._f8 = f8
	}
}
func (_this *C2) Fm2(p0 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("ikasy")
	}
}
func (_this *C2) Fm0(p0 bool, globalState *GlobalState) _dafny.Set {
	{
		return _dafny.SetOf((_dafny.SetOf(false, false, true, false)).Difference(_dafny.SetOf(false)), (_dafny.SetOf(false, true)).Difference(_dafny.SetOf(true, false, true, true)), (_dafny.SetOf(!(false), true)).Intersection(_dafny.SetOf((false))))
	}
}
func (_this *C2) Fm1(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((func() bool {
			if false {
				return false
			}
			return true
		})())
	}
}
func (_this *C2) Fm14(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qx"), _dafny.UnicodeSeqOfUtf8Bytes("cv"))
	}
}
func (_this *C2) Fm15(globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C2) M1(p0 _dafny.Sequence, globalState *GlobalState) {
	{
		var _455_v0 _dafny.Array
		_ = _455_v0
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_8
		var _nw77 _dafny.Array
		_ = _nw77
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw77 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) bool = func(_456_i0 _dafny.Int) bool {
				return ((_this).F8()).Cmp((_this).F8()) < 0
			}
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw77 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw77).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw77).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_455_v0 = _nw77
		_455_v0 = _455_v0
		var _457_v1 bool
		_ = _457_v1
		_457_v1 = false
		var _458_v2 D2
		_ = _458_v2
		_458_v2 = Companion_D2_.Create_DC5_(_457_v1, (_this).F8(), (_this).F7(), (_this).F7(), _457_v1)
		var _hi2 _dafny.Int = (_458_v2).Dtor_cf5()
		_ = _hi2
		for _459_i1 := (_this).F7(); _459_i1.Cmp(_hi2) < 0; _459_i1 = _459_i1.Plus(_dafny.One) {
			var _460_v3 _dafny.Sequence
			_ = _460_v3
			_460_v3 = _dafny.UnicodeSeqOfUtf8Bytes("rsl")
			_460_v3 = _dafny.Companion_Sequence_.Concatenate(p0, p0)
			var _461_v4 _dafny.Array
			_ = _461_v4
			var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
			_ = _nw78
			_461_v4 = _nw78
			var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
			_ = _nw79
			_461_v4 = _nw79
			var _462_v5 _dafny.Array
			_ = _462_v5
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_9
			var _nw80 _dafny.Array
			_ = _nw80
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw80 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Int = (func(_463_i1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_464_i2 _dafny.Int) _dafny.Int {
						return (_464_i2).Plus(_463_i1)
					}
				})(_459_i1)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw80 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw80).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw80).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_462_v5 = _nw80
			var _465_v6 _dafny.Map
			_ = _465_v6
			_465_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(375), (_this).F7())
			var _466_v7 *C0
			_ = _466_v7
			var _nw81 *C0 = New_C0_()
			_ = _nw81
			_nw81.Ctor__(_462_v5, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_460_v3, p0)).Cardinality()), (_465_v6).Cardinality())
			_466_v7 = _nw81
			if (_dafny.IntOfInt64(860)).Cmp(_dafny.IntOfUint32(((_this).Fm2(_457_v1, globalState)).Cardinality())) < 0 {
				var _467_v8 _dafny.Int
				_ = _467_v8
				_467_v8 = _dafny.IntOfInt64(-417)
				_467_v8 = (_dafny.IntOfInt64(-602)).Minus((_this).F8())
				_467_v8 = _467_v8
				(globalState).F4 = _457_v1
				(globalState).F3 = (func() bool {
					if _457_v1 {
						return _457_v1
					}
					return (_459_i1).Cmp(_dafny.IntOfInt64(-965)) >= 0
				})()
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_455_v0), 0))
				_ = _index89
				(_455_v0).ArraySet1(_457_v1, (_index89).Int())
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_455_v0), 0))
				_ = _index90
				(_455_v0).ArraySet1(!(_457_v1), (_index90).Int())
			} else {
				var _468_v9 _dafny.MultiSet
				_ = _468_v9
				_468_v9 = _dafny.MultiSetOf((_this).F8())
				var _469_v10 _dafny.Sequence
				_ = _469_v10
				_469_v10 = _dafny.SeqOf(_468_v9, _468_v9, _468_v9)
				var _470_v11 _dafny.Map
				_ = _470_v11
				_470_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_455_v0, (_this).F7())
				var _471_v12 _dafny.Map
				_ = _471_v12
				_471_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_470_v11).Cardinality(), _465_v6)
				var _472_v13 _dafny.Map
				_ = _472_v13
				_472_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_469_v10).Cardinality()), (_471_v12).Merge(_471_v12))
				_472_v13 = (_472_v13).Update(Companion_Default___.Fm16(_457_v1, globalState), (_471_v12).Update(_459_i1, _465_v6))
				var _473_v14 _dafny.Sequence
				_ = _473_v14
				_473_v14 = _dafny.SeqOf((_this).F8())
				var _474_v15 _dafny.Sequence
				_ = _474_v15
				_474_v15 = _dafny.SeqOf(_473_v14)
				var _475_v16 _dafny.MultiSet
				_ = _475_v16
				_475_v16 = _dafny.MultiSetOf(_457_v1, _457_v1, !(_457_v1), _dafny.Companion_Sequence_.Equal((_474_v15).Select((Companion_Default___.SafeIndex((_458_v2).Dtor_cf5(), _dafny.IntOfUint32((_474_v15).Cardinality()))).Uint32()).(_dafny.Sequence), _473_v14))
				var _476_v17 _dafny.CodePoint
				_ = _476_v17
				_476_v17 = _dafny.CodePoint('a')
				var _477_v18 _dafny.Map
				_ = _477_v18
				_477_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_457_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_460_v3, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_459_i1), _dafny.IntOfUint32((_460_v3).Cardinality()))).Uint32(), _476_v17)).Cardinality()))
				var _478_v19 _dafny.Sequence
				_ = _478_v19
				_478_v19 = _dafny.SeqOf(true, !(!(_457_v1)), _457_v1)
				var _rhs79 bool = false
				_ = _rhs79
				var _rhs80 _dafny.MultiSet = ((_475_v16).Union(_475_v16)).Union(_475_v16)
				_ = _rhs80
				var _rhs81 _dafny.Map = (_477_v18).Merge(_477_v18)
				_ = _rhs81
				var _rhs82 _dafny.Sequence = _478_v19
				_ = _rhs82
				var _lhs74 *GlobalState = globalState
				_ = _lhs74
				_lhs74.F3 = _rhs79
				_475_v16 = _rhs80
				_477_v18 = _rhs81
				_478_v19 = _rhs82
				var _479_v20 *C0
				_ = _479_v20
				var _nw82 *C0 = New_C0_()
				_ = _nw82
				_nw82.Ctor__(_466_v7.F13, (_this).F8(), (_this).F7())
				_479_v20 = _nw82
				var _480_v21 _dafny.Array
				_ = _480_v21
				var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(2))
				_ = _nw83
				_480_v21 = _nw83
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_480_v21), 0))
				_ = _index91
				(_480_v21).ArraySet1CodePoint((p0).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm16(false, globalState), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index91).Int())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_480_v21), 0))
				_ = _index92
				(_480_v21).ArraySet1CodePoint(_476_v17, (_index92).Int())
				var _481_v22 *C0
				_ = _481_v22
				var _nw84 *C0 = New_C0_()
				_ = _nw84
				_nw84.Ctor__((func() _dafny.Array {
					if _457_v1 {
						return _466_v7.F13
					}
					return _462_v5
				})(), Companion_Default___.SafeModuloInt(_459_i1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(286))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}((func(_482_v21 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_483_i3 _dafny.Int) _dafny.CodePoint {
						return (_482_v21).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_482_v21), 0))).Int())
					}
				})(_480_v21)))).Cardinality())), Companion_Default___.Fm16(_457_v1, globalState))
				_481_v22 = _nw84
			}
		}
		var _484_v24 _dafny.MultiSet
		_ = _484_v24
		_484_v24 = _dafny.MultiSetOf(_457_v1, _457_v1)
		var _485_v25 _dafny.Set
		_ = _485_v25
		_485_v25 = _dafny.SetOf((_484_v24).Cardinality(), (_this).F8())
		var _486_v26 _dafny.Array
		_ = _486_v26
		var _nwElement0_15 _dafny.Set = func() _dafny.Set {
			var _coll25 = _dafny.NewBuilder()
			_ = _coll25
			for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(695), _dafny.IntOfInt64(-840))); ; {
				_compr_25, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _487_v23 _dafny.Int
				_487_v23 = interface{}(_compr_25).(_dafny.Int)
				if ((_dafny.IntOfInt64(695)).Cmp(_487_v23) <= 0) && ((_487_v23).Cmp(_dafny.IntOfInt64(-840)) < 0) {
					_coll25.Add((_487_v23).Times((_dafny.Zero).Minus((_this).F7())))
				}
			}
			return _coll25.ToSet()
		}()
		_ = _nwElement0_15
		var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(2))
		_ = _nw85
		(_nw85).ArraySet1(_nwElement0_15, 0)
		(_nw85).ArraySet1(_485_v25, 1)
		_486_v26 = _nw85
		var _488_v27 _dafny.Map
		_ = _488_v27
		_488_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F8()))), _486_v26)
		var _489_v28 _dafny.Array
		_ = _489_v28
		var _nwElement0_16 _dafny.Array = (func() _dafny.Array {
			if _457_v1 {
				return _486_v26
			}
			return _486_v26
		})()
		_ = _nwElement0_16
		var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(11))
		_ = _nw86
		(_nw86).ArraySet1(_nwElement0_16, 0)
		(_nw86).ArraySet1((func() _dafny.Array {
			if (_488_v27).Contains((_this).F8()) {
				return (_488_v27).Get((_this).F8()).(_dafny.Array)
			}
			return _486_v26
		})(), 1)
		(_nw86).ArraySet1(_486_v26, 2)
		(_nw86).ArraySet1(_486_v26, 3)
		(_nw86).ArraySet1(_486_v26, 4)
		(_nw86).ArraySet1(_486_v26, 5)
		(_nw86).ArraySet1(_486_v26, 6)
		(_nw86).ArraySet1(_486_v26, 7)
		(_nw86).ArraySet1(_486_v26, 8)
		(_nw86).ArraySet1(_486_v26, 9)
		(_nw86).ArraySet1(_486_v26, 10)
		_489_v28 = _nw86
		var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_489_v28), 0))
		_ = _index93
		(_489_v28).ArraySet1(_486_v26, (_index93).Int())
		var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_489_v28), 0))
		_ = _index94
		(_489_v28).ArraySet1(_486_v26, (_index94).Int())
		var _490_v29 _dafny.MultiSet
		_ = _490_v29
		_490_v29 = _dafny.MultiSetOf(_dafny.IntOfInt64(715), (_dafny.Zero).Minus((_this).F7()))
		var _491_v30 _dafny.Map
		_ = _491_v30
		_491_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_457_v1, _457_v1)
		var _492_v31 _dafny.Array
		_ = _492_v31
		var _nwElement0_17 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_490_v29).Cardinality(), _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_484_v24).Cardinality(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _dafny.CodePoint('y')))).Cardinality()
		_ = _nwElement0_17
		var _nw87 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(25))
		_ = _nw87
		(_nw87).ArraySet1(_nwElement0_17, 0)
		(_nw87).ArraySet1((_this).F8(), 1)
		(_nw87).ArraySet1((_485_v25).Cardinality(), 2)
		(_nw87).ArraySet1(_dafny.IntOfInt64(285), 3)
		(_nw87).ArraySet1((_this).F7(), 4)
		(_nw87).ArraySet1((_this).F8(), 5)
		(_nw87).ArraySet1((_this).F8(), 6)
		(_nw87).ArraySet1((_this).F8(), 7)
		(_nw87).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 8)
		(_nw87).ArraySet1((_this).F7(), 9)
		(_nw87).ArraySet1((_this).F8(), 10)
		(_nw87).ArraySet1((_this).F8(), 11)
		(_nw87).ArraySet1(_dafny.IntOfInt64(891), 12)
		(_nw87).ArraySet1(Companion_Default___.Fm16(_457_v1, globalState), 13)
		(_nw87).ArraySet1((_dafny.Zero).Minus((_this).F7()), 14)
		(_nw87).ArraySet1((_this).F8(), 15)
		(_nw87).ArraySet1((_this).F7(), 16)
		(_nw87).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 17)
		(_nw87).ArraySet1((_dafny.Zero).Minus((_491_v30).Cardinality()), 18)
		(_nw87).ArraySet1((_this).F8(), 19)
		(_nw87).ArraySet1((_this).F7(), 20)
		(_nw87).ArraySet1((_this).F7(), 21)
		(_nw87).ArraySet1((_this).F7(), 22)
		(_nw87).ArraySet1((_this).F7(), 23)
		(_nw87).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 24)
		_492_v31 = _nw87
		var _493_v32 *C0
		_ = _493_v32
		var _nw88 *C0 = New_C0_()
		_ = _nw88
		_nw88.Ctor__(_492_v31, (_this).F7(), (_this).F7())
		_493_v32 = _nw88
		(globalState).F4 = _457_v1
		var _494_v33 *C0
		_ = _494_v33
		var _nw89 *C0 = New_C0_()
		_ = _nw89
		_nw89.Ctor__(_493_v32.F13, (_dafny.Zero).Minus((_this).F8()), ((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("trnq"), p0, p0, p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(53))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}(func(_495_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})))).Update(p0, Companion_Default___.Abs((_this).F7()))).Cardinality())
		_494_v33 = _nw89
	}
}
func (_this *C2) M2(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _496_v0 _dafny.Array
		_ = _496_v0
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_10
		var _nw90 _dafny.Array
		_ = _nw90
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw90 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.Int = func(_497_i0 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_497_i0, (_this).F8())
			}
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw90 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw90).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw90).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_496_v0 = _nw90
		var _498_v1 *C0
		_ = _498_v1
		var _nw91 *C0 = New_C0_()
		_ = _nw91
		_nw91.Ctor__(_496_v0, Companion_Default___.SafeModuloInt(p0, (_this).F8()), _dafny.IntOfInt64(765))
		_498_v1 = _nw91
		(globalState).F4 = p1
		var _499_v2 _dafny.Sequence
		_ = _499_v2
		_499_v2 = _dafny.SeqOf(true, p3)
		var _500_v3 D2
		_ = _500_v3
		_500_v3 = Companion_D2_.Create_DC4_(_dafny.MultiSetFromSeq(_499_v2))
		var _source11 D2 = _500_v3
		_ = _source11
		if _source11.Is_DC5() {
			var _501___mcc_h0 bool = _source11.Get_().(D2_DC5).Cf4
			_ = _501___mcc_h0
			var _502___mcc_h1 _dafny.Int = _source11.Get_().(D2_DC5).Cf5
			_ = _502___mcc_h1
			var _503___mcc_h2 _dafny.Int = _source11.Get_().(D2_DC5).Cf6
			_ = _503___mcc_h2
			var _504___mcc_h3 _dafny.Int = _source11.Get_().(D2_DC5).Cf7
			_ = _504___mcc_h3
			var _505___mcc_h4 bool = _source11.Get_().(D2_DC5).Cf8
			_ = _505___mcc_h4
			var _506_cf8 bool = _505___mcc_h4
			_ = _506_cf8
			var _507_cf7 _dafny.Int = _504___mcc_h3
			_ = _507_cf7
			var _508_cf6 _dafny.Int = _503___mcc_h2
			_ = _508_cf6
			var _509_cf5 _dafny.Int = _502___mcc_h1
			_ = _509_cf5
			var _510_cf4 bool = _501___mcc_h0
			_ = _510_cf4
			if !(((_this).F8()).Cmp(_509_cf5) < 0) {
				(globalState).F4 = _510_cf4
				r2 = (func() bool {
					if p3 {
						return _506_cf8
					}
					return (func() bool {
						if _506_cf8 {
							return _506_cf8
						}
						return false
					})()
				})()
				(globalState).F3 = p1
				var _511_v4 _dafny.CodePoint
				_ = _511_v4
				_511_v4 = _dafny.CodePoint('c')
				var _512_v5 _dafny.Map
				_ = _512_v5
				_512_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if true {
						return p3
					}
					return false
				})(), _511_v4)
				_512_v5 = (_512_v5).Update(p1, _dafny.CodePoint('e'))
				var _513_v6 _dafny.Sequence
				_ = _513_v6
				_513_v6 = _dafny.UnicodeSeqOfUtf8Bytes("ksccjtf")
				_513_v6 = _513_v6
			} else {
				_508_cf6 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((_this).Fm14(globalState)).Cardinality()), _507_cf7)
				var _514_v7 _dafny.CodePoint
				_ = _514_v7
				_514_v7 = _dafny.CodePoint('g')
				var _rhs83 _dafny.CodePoint = _514_v7
				_ = _rhs83
				var _rhs84 _dafny.Sequence = _499_v2
				_ = _rhs84
				_514_v7 = _rhs83
				r0 = _rhs84
				var _515_v8 _dafny.Array
				_ = _515_v8
				var _nw92 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
				_ = _nw92
				_515_v8 = _nw92
				var _516_v9 _dafny.Sequence
				_ = _516_v9
				_516_v9 = _dafny.UnicodeSeqOfUtf8Bytes("ucud")
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_515_v8), 0))
				_ = _index95
				(_515_v8).ArraySet1(_dafny.Companion_Sequence_.Update(_516_v9, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.IntOfUint32((_516_v9).Cardinality()))).Uint32(), _514_v7), (_index95).Int())
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_515_v8), 0))
				_ = _index96
				(_515_v8).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_516_v9, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-60))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}((func(_517_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_518_i1 _dafny.Int) _dafny.CodePoint {
						return _517_v7
					}
				})(_514_v7)))), (_index96).Int())
				r0 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p3), _dafny.Companion_Sequence_.Concatenate(_499_v2, _dafny.SeqOf(p3)))
				var _519_v10 _dafny.Array
				_ = _519_v10
				var _nw93 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
				_ = _nw93
				_519_v10 = _nw93
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_519_v10), 0))
				_ = _index97
				(_519_v10).ArraySet1(_506_cf8, (_index97).Int())
				var _520_v11 _dafny.Sequence
				_ = _520_v11
				_520_v11 = _dafny.SeqOf(_508_cf6)
				var _521_v12 _dafny.Map
				_ = _521_v12
				_521_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_510_cf4, _520_v11)
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_519_v10), 0))
				_ = _index98
				var _rhs85 bool = p3
				_ = _rhs85
				var _rhs86 bool = p1
				_ = _rhs86
				var _rhs87 _dafny.Int = (func() _dafny.Int {
					if p1 {
						return _508_cf6
					}
					return _508_cf6
				})()
				_ = _rhs87
				var _rhs88 _dafny.Sequence = _520_v11
				_ = _rhs88
				var _rhs89 _dafny.Map = Companion_Default___.Fm17(_514_v7, _510_cf4, globalState)
				_ = _rhs89
				var _lhs75 _dafny.Array = _519_v10
				_ = _lhs75
				var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_519_v10), 0))
				_ = _lhs76
				(_lhs75).ArraySet1(_rhs85, (_lhs76).Int())
				r2 = _rhs86
				_508_cf6 = _rhs87
				_520_v11 = _rhs88
				_521_v12 = _rhs89
			}
			if p1 {
				var _522_v13 _dafny.Sequence
				_ = _522_v13
				_522_v13 = _dafny.UnicodeSeqOfUtf8Bytes("v")
				_522_v13 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-528))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}(func(_523_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				})), _522_v13)
				var _524_v14 _dafny.Map
				_ = _524_v14
				_524_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _dafny.UnicodeSeqOfUtf8Bytes("lb"))
				_524_v14 = (_524_v14).Update(_508_cf6, _522_v13)
				var _525_v15 T0
				_ = _525_v15
				var _nw94 *C1 = New_C1_()
				_ = _nw94
				_nw94.Ctor__((_this).F7(), (_this).F8())
				_525_v15 = _nw94
				_525_v15 = (func() T0 {
					if p1 {
						return _525_v15
					}
					return _525_v15
				})()
				var _526_v16 _dafny.Set
				_ = _526_v16
				_526_v16 = _dafny.SetOf(_506_cf8)
				var _527_v17 _dafny.Map
				_ = _527_v17
				_527_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_525_v15).F7(), _dafny.IntOfInt64(833))
				(globalState).F3 = (_498_v1).Fm1(_526_v16, (_this).F8(), (_dafny.Zero).Minus((_527_v17).Cardinality()), (_525_v15).F8(), globalState)
				var _528_v18 _dafny.Map
				_ = _528_v18
				_528_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_525_v15).F8())
				_528_v18 = (_528_v18).Update(p1, ((func() _dafny.Map {
					var _coll26 = _dafny.NewMapBuilder()
					_ = _coll26
					for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(437), _dafny.IntOfInt64(-935))); ; {
						_compr_26, _ok28 := _iter28()
						if !_ok28 {
							break
						}
						var _529_v19 _dafny.Int
						_529_v19 = interface{}(_compr_26).(_dafny.Int)
						if ((_dafny.IntOfInt64(437)).Cmp(_529_v19) <= 0) && ((_529_v19).Cmp(_dafny.IntOfInt64(-935)) < 0) {
							_coll26.Add(Companion_Default___.SafeModuloInt(_529_v19, _509_cf5), p3)
						}
					}
					return _coll26.ToMap()
				}()).Cardinality()).Minus(p0))
			} else {
				var _530_v20 _dafny.Map
				_ = _530_v20
				_530_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), (_509_cf5).Cmp(p2) < 0)
				_530_v20 = (_530_v20).Update(_506_cf8, _506_cf8)
				(globalState).F4 = (_508_cf6).Cmp(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(915), _507_cf7)) > 0
				var _531_v21 _dafny.CodePoint
				_ = _531_v21
				_531_v21 = _dafny.CodePoint('v')
				var _532_v22 _dafny.Sequence
				_ = _532_v22
				_532_v22 = _dafny.SeqOf(_531_v21)
				var _arr10 _dafny.Array = _498_v1.F13
				_ = _arr10
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_498_v1.F13), 0))
				_ = _index99
				(_arr10).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_532_v22, _dafny.SeqOf(_dafny.CodePoint('j')))).Cardinality()), (_index99).Int())
				var _533_v23 _dafny.Map
				_ = _533_v23
				_533_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_506_cf8), p0)
				var _534_v24 _dafny.MultiSet
				_ = _534_v24
				_534_v24 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfUint32((_532_v22).Cardinality())))
				var _arr11 _dafny.Array = _498_v1.F13
				_ = _arr11
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_498_v1.F13), 0))
				_ = _index100
				(_arr11).ArraySet1((func() _dafny.Int {
					if _506_cf8 {
						return (func() _dafny.Int {
							if (_533_v23).Contains(p3) {
								return (_533_v23).Get(p3).(_dafny.Int)
							}
							return (_534_v24).Cardinality()
						})()
					}
					return ((_this).F7()).Minus(p2)
				})(), (_index100).Int())
				_530_v20 = (_530_v20).Update(false, (_dafny.IntOfUint32((_532_v22).Cardinality())).Cmp(_507_cf7) == 0)
				var _535_v25 _dafny.Sequence
				_ = _535_v25
				_535_v25 = _dafny.SeqOf(_509_cf5, _507_cf7)
				var _536_v26 _dafny.MultiSet
				_ = _536_v26
				_536_v26 = _dafny.MultiSetOf(_535_v25)
				_536_v26 = _dafny.MultiSetOf(_535_v25, _535_v25)
			}
			_509_cf5 = _507_cf7
			(globalState).F4 = _510_cf4
		} else {
			var _537___mcc_h5 _dafny.MultiSet = _source11.Get_().(D2_DC4).Cf3
			_ = _537___mcc_h5
			var _538_cf3 _dafny.MultiSet = _537___mcc_h5
			_ = _538_cf3
			r1 = (_this).F8()
			var _539_v27 _dafny.CodePoint
			_ = _539_v27
			_539_v27 = _dafny.CodePoint('y')
			var _540_v28 _dafny.Set
			_ = _540_v28
			_540_v28 = _dafny.SetOf(_539_v27, _539_v27, _dafny.CodePoint('d'))
			var _541_v29 _dafny.Sequence
			_ = _541_v29
			_541_v29 = _dafny.SeqOf(_540_v28)
			(globalState).F4 = (p2).Cmp(_dafny.IntOfUint32((_541_v29).Cardinality())) >= 0
			var _542_v30 _dafny.Array
			_ = _542_v30
			var _nw95 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw95
			_542_v30 = _nw95
			var _543_v31 _dafny.Sequence
			_ = _543_v31
			_543_v31 = _dafny.SeqOf(_542_v30)
			var _544_v32 D4
			_ = _544_v32
			_544_v32 = Companion_D4_.Create_DC9_(p0, _543_v31, (_this).F7(), p0)
			var _545_v33 _dafny.Array
			_ = _545_v33
			var _nw96 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
			_ = _nw96
			_545_v33 = _nw96
			var _546_v34 D4
			_ = _546_v34
			_546_v34 = Companion_D4_.Create_DC8_(_545_v33)
			var _547_v35 _dafny.Map
			_ = _547_v35
			_547_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _546_v34)
			var _pat_let_tv5 = p0
			_ = _pat_let_tv5
			var _pat_let_tv6 = _547_v35
			_ = _pat_let_tv6
			var _source12 D2 = func(_pat_let0_0 D2) D2 {
				return func(_551_dt__update__tmp_h1 D2) D2 {
					return func(_pat_let4_0 _dafny.Int) D2 {
						return func(_552_dt__update_hcf6_h0 _dafny.Int) D2 {
							return Companion_D2_.Create_DC5_((_551_dt__update__tmp_h1).Dtor_cf4(), (_551_dt__update__tmp_h1).Dtor_cf5(), _552_dt__update_hcf6_h0, (_551_dt__update__tmp_h1).Dtor_cf7(), (_551_dt__update__tmp_h1).Dtor_cf8())
						}(_pat_let4_0)
					}((_dafny.Zero).Minus((_dafny.Zero).Minus((_pat_let_tv6).Cardinality())))
				}(_pat_let0_0)
			}(func(_pat_let1_0 D2) D2 {
				return func(_548_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let2_0 _dafny.Int) D2 {
						return func(_549_dt__update_hcf5_h0 _dafny.Int) D2 {
							return func(_pat_let3_0 bool) D2 {
								return func(_550_dt__update_hcf8_h0 bool) D2 {
									return Companion_D2_.Create_DC5_((_548_dt__update__tmp_h0).Dtor_cf4(), _549_dt__update_hcf5_h0, (_548_dt__update__tmp_h0).Dtor_cf6(), (_548_dt__update__tmp_h0).Dtor_cf7(), _550_dt__update_hcf8_h0)
								}(_pat_let3_0)
							}(false)
						}(_pat_let2_0)
					}(_pat_let_tv5)
				}(_pat_let1_0)
			}(Companion_D2_.Create_DC5_(true, (_this).F8(), _dafny.IntOfInt64(475), (_544_v32).Dtor_cf14(), p3)))
			_ = _source12
			if _source12.Is_DC5() {
				var _553___mcc_h6 bool = _source12.Get_().(D2_DC5).Cf4
				_ = _553___mcc_h6
				var _554___mcc_h7 _dafny.Int = _source12.Get_().(D2_DC5).Cf5
				_ = _554___mcc_h7
				var _555___mcc_h8 _dafny.Int = _source12.Get_().(D2_DC5).Cf6
				_ = _555___mcc_h8
				var _556___mcc_h9 _dafny.Int = _source12.Get_().(D2_DC5).Cf7
				_ = _556___mcc_h9
				var _557___mcc_h10 bool = _source12.Get_().(D2_DC5).Cf8
				_ = _557___mcc_h10
				var _558_cf8 bool = _557___mcc_h10
				_ = _558_cf8
				var _559_cf7 _dafny.Int = _556___mcc_h9
				_ = _559_cf7
				var _560_cf6 _dafny.Int = _555___mcc_h8
				_ = _560_cf6
				var _561_cf5 _dafny.Int = _554___mcc_h7
				_ = _561_cf5
				var _562_cf4 bool = _553___mcc_h6
				_ = _562_cf4
				var _563_v36 _dafny.Map
				_ = _563_v36
				_563_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_558_cf8), _498_v1.F13)
				var _564_v37 _dafny.Set
				_ = _564_v37
				_564_v37 = _dafny.SetOf((_563_v36).Cardinality())
				var _565_v38 D3
				_ = _565_v38
				_565_v38 = Companion_D3_.Create_DC6_((_564_v37).Union(_564_v37))
				_565_v38 = _565_v38
				(globalState).F4 = p3
				var _566_v39 _dafny.Set
				_ = _566_v39
				_566_v39 = _dafny.SetOf(p1)
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_542_v30), 0))
				_ = _index101
				(_542_v30).ArraySet1(true, (_index101).Int())
				var _567_v40 _dafny.Sequence
				_ = _567_v40
				_567_v40 = _dafny.SeqOf(_566_v39, _566_v39, _dafny.SetOf((_499_v2).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_499_v2).Cardinality()))).Uint32()).(bool)))
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_542_v30), 0))
				_ = _index102
				var _rhs90 _dafny.Set = (_567_v40).Select((Companion_Default___.SafeIndex((_dafny.IntOfUint32((_499_v2).Cardinality())).Plus(_559_cf7), _dafny.IntOfUint32((_567_v40).Cardinality()))).Uint32()).(_dafny.Set)
				_ = _rhs90
				var _rhs91 bool = (func() bool {
					if p3 {
						return ((_this).F8()).Cmp(p0) == 0
					}
					return p1
				})()
				_ = _rhs91
				var _lhs77 _dafny.Array = _542_v30
				_ = _lhs77
				var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_542_v30), 0))
				_ = _lhs78
				_566_v39 = _rhs90
				(_lhs77).ArraySet1(_rhs91, (_lhs78).Int())
				_560_cf6 = (_this).F8()
			} else {
				var _568___mcc_h11 _dafny.MultiSet = _source12.Get_().(D2_DC4).Cf3
				_ = _568___mcc_h11
				var _569_cf3 _dafny.MultiSet = _568___mcc_h11
				_ = _569_cf3
				(globalState).F4 = ((_this).F7()).Cmp(p2) > 0
				r1 = (_dafny.Zero).Minus(_dafny.IntOfInt64(-380))
				var _570_v41 _dafny.Array
				_ = _570_v41
				var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
				_ = _nw97
				_570_v41 = _nw97
				var _571_v44 _dafny.Set
				_ = _571_v44
				_571_v44 = _dafny.SetOf(p0)
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_570_v41), 0))
				_ = _index103
				(_570_v41).ArraySet1((func() _dafny.Set {
					var _coll27 = _dafny.NewBuilder()
					_ = _coll27
					for _iter29 := _dafny.Iterate((func() _dafny.Set {
						var _coll28 = _dafny.NewBuilder()
						_ = _coll28
						for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-205), _dafny.IntOfInt64(46))); ; {
							_compr_28, _ok30 := _iter30()
							if !_ok30 {
								break
							}
							var _572_v42 _dafny.Int
							_572_v42 = interface{}(_compr_28).(_dafny.Int)
							if ((_dafny.IntOfInt64(-205)).Cmp(_572_v42) <= 0) && ((_572_v42).Cmp(_dafny.IntOfInt64(46)) < 0) {
								_coll28.Add(Companion_Default___.SafeDivisionInt(_572_v42, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vudgnl")).Cardinality())))
							}
						}
						return _coll28.ToSet()
					}()).Elements()); ; {
						_compr_27, _ok29 := _iter29()
						if !_ok29 {
							break
						}
						var _573_v43 _dafny.Int
						_573_v43 = interface{}(_compr_27).(_dafny.Int)
						if (func() _dafny.Set {
							var _coll29 = _dafny.NewBuilder()
							_ = _coll29
							for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-205), _dafny.IntOfInt64(46))); ; {
								_compr_29, _ok31 := _iter31()
								if !_ok31 {
									break
								}
								var _574_v42 _dafny.Int
								_574_v42 = interface{}(_compr_29).(_dafny.Int)
								if ((_dafny.IntOfInt64(-205)).Cmp(_574_v42) <= 0) && ((_574_v42).Cmp(_dafny.IntOfInt64(46)) < 0) {
									_coll29.Add(Companion_Default___.SafeDivisionInt(_574_v42, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vudgnl")).Cardinality())))
								}
							}
							return _coll29.ToSet()
						}()).Contains(_573_v43) {
							_coll27.Add(Companion_Default___.SafeDivisionInt(_573_v43, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(432))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg42 _dafny.Int) interface{} {
									return coer42(arg42)
								}
							}(func(_575_i3 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('o')
							}))).Cardinality())))
						}
					}
					return _coll27.ToSet()
				}()).Difference(_571_v44), (_index103).Int())
				var _576_v46 D3
				_ = _576_v46
				_576_v46 = Companion_D3_.Create_DC6_(_dafny.SetOf(p2))
				var _pat_let_tv7 = p2
				_ = _pat_let_tv7
				var _pat_let_tv8 = _543_v31
				_ = _pat_let_tv8
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_570_v41), 0))
				_ = _index104
				var _rhs92 bool = (_540_v28).IsSubsetOf((_540_v28).Intersection(func() _dafny.Set {
					var _coll30 = _dafny.NewBuilder()
					_ = _coll30
					for _iter32 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(986))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg43 _dafny.Int) interface{} {
							return coer43(arg43)
						}
					}((func(_577_v27 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_578_i4 _dafny.Int) _dafny.CodePoint {
							return _577_v27
						}
					})(_539_v27)))).Elements()); ; {
						_compr_30, _ok32 := _iter32()
						if !_ok32 {
							break
						}
						var _579_v45 _dafny.CodePoint
						_579_v45 = interface{}(_compr_30).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(986))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg44 _dafny.Int) interface{} {
								return coer44(arg44)
							}
						}((func(_580_v27 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_578_i4 _dafny.Int) _dafny.CodePoint {
								return _580_v27
							}
						})(_539_v27))), _579_v45) {
							_coll30.Add(_579_v45)
						}
					}
					return _coll30.ToSet()
				}()))
				_ = _rhs92
				var _rhs93 _dafny.Sequence = _499_v2
				_ = _rhs93
				var _rhs94 _dafny.Int = (_this).F7()
				_ = _rhs94
				var _rhs95 D4 = func(_pat_let5_0 D4) D4 {
					return func(_581_dt__update__tmp_h2 D4) D4 {
						return func(_pat_let6_0 _dafny.Int) D4 {
							return func(_582_dt__update_hcf14_h0 _dafny.Int) D4 {
								return Companion_D4_.Create_DC9_((_581_dt__update__tmp_h2).Dtor_cf12(), (_581_dt__update__tmp_h2).Dtor_cf13(), _582_dt__update_hcf14_h0, (_581_dt__update__tmp_h2).Dtor_cf15())
							}(_pat_let6_0)
						}((_pat_let_tv7).Minus(_dafny.IntOfUint32((_pat_let_tv8).Cardinality())))
					}(_pat_let5_0)
				}(_544_v32)
				_ = _rhs95
				var _rhs96 _dafny.Set = (_576_v46).Dtor_cf9()
				_ = _rhs96
				var _lhs79 _dafny.Array = _570_v41
				_ = _lhs79
				var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_570_v41), 0))
				_ = _lhs80
				r2 = _rhs92
				r0 = _rhs93
				r1 = _rhs94
				_544_v32 = _rhs95
				(_lhs79).ArraySet1(_rhs96, (_lhs80).Int())
				_542_v30 = _542_v30
			}
			var _583_v47 _dafny.MultiSet
			_ = _583_v47
			_583_v47 = _dafny.MultiSetOf(p0, p2)
			var _584_v48 *C1
			_ = _584_v48
			var _nw98 *C1 = New_C1_()
			_ = _nw98
			_nw98.Ctor__((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm16(true, globalState))), (_dafny.Zero).Minus((Companion_Default___.Fm29(p1, _583_v47, globalState)).Cardinality()))
			_584_v48 = _nw98
		}
		var _585_v49 _dafny.MultiSet
		_ = _585_v49
		_585_v49 = _dafny.MultiSetOf(p1)
		var _586_v50 _dafny.Sequence
		_ = _586_v50
		_586_v50 = _dafny.SeqOf(Companion_Default___.Fm20(p3, !(false), p3, globalState), (func() _dafny.Int {
			if (_585_v49).Contains(p3) {
				return (_585_v49).Multiplicity(p3)
			}
			return (_dafny.Zero).Minus((_this).F7())
		})())
		var _hi3 _dafny.Int = p0
		_ = _hi3
		for _587_i5 := _dafny.IntOfUint32(((func() _dafny.Sequence {
			if p3 {
				return _586_v50
			}
			return _586_v50
		})()).Cardinality()); _587_i5.Cmp(_hi3) < 0; _587_i5 = _587_i5.Plus(_dafny.One) {
			(globalState).F3 = p1
			var _588_v51 _dafny.Sequence
			_ = _588_v51
			_588_v51 = _dafny.UnicodeSeqOfUtf8Bytes("nvldbus")
			var _589_v52 _dafny.Map
			_ = _589_v52
			_589_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfUint32((_588_v51).Cardinality()))
			var _590_v53 _dafny.MultiSet
			_ = _590_v53
			_590_v53 = _dafny.MultiSetOf(p0)
			_589_v52 = (_589_v52).Update(p3, (func() _dafny.Int {
				if (_590_v53).Contains(p2) {
					return (_590_v53).Multiplicity(p2)
				}
				return p0
			})())
			var _591_v54 _dafny.Array
			_ = _591_v54
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_11
			var _nw99 _dafny.Array
			_ = _nw99
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw99 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) bool = (func(_592_p3 bool) func(_dafny.Int) bool {
					return func(_593_i6 _dafny.Int) bool {
						return _592_p3
					}
				})(p3)
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
			_591_v54 = _nw99
			var _594_v55 D4
			_ = _594_v55
			_594_v55 = Companion_D4_.Create_DC9_(p2, _dafny.SeqOf(_591_v54, _591_v54, _591_v54, _591_v54, _591_v54), _dafny.IntOfInt64(688), (_this).F8())
			var _arr12 _dafny.Array = _498_v1.F13
			_ = _arr12
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_498_v1.F13), 0))
			_ = _index105
			(_arr12).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_589_v52).Contains(p1) {
					return (_589_v52).Get(p1).(_dafny.Int)
				}
				return p2
			})(), (_594_v55).Dtor_cf14()), (_index105).Int())
			var _arr13 _dafny.Array = _498_v1.F13
			_ = _arr13
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_498_v1.F13), 0))
			_ = _index106
			(_arr13).ArraySet1((p0).Times((_this).F7()), (_index106).Int())
			var _595_v56 _dafny.Map
			_ = _595_v56
			_595_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_588_v51).Cardinality()), _496_v0)
			var _596_v57 _dafny.Map
			_ = _596_v57
			_596_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_595_v56, _590_v53)
			_596_v57 = (_596_v57).Update(_595_v56, _dafny.MultiSetFromSeq((func() _dafny.Sequence {
				if p1 {
					return _586_v50
				}
				return _dafny.SeqOf((_590_v53).Cardinality(), (_this).F7(), _dafny.IntOfInt64(-717), _dafny.IntOfInt64(620), _dafny.IntOfInt64(742))
			})()))
		}
		r1 = Companion_Default___.Fm16((_this).Fm15(globalState), globalState)
		for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_496_v0), 0))); ; {
			_guard_loop_2, _ok33 := _iter33()
			if !_ok33 {
				break
			}
			var _597_i7 _dafny.Int
			_597_i7 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_597_i7).Sign() != -1) && ((_597_i7).Cmp(_dafny.ArrayLen((_496_v0), 0)) < 0)) {
				(_496_v0).ArraySet1(Companion_Default___.SafeModuloInt(_597_i7, (_this).F8()), (_597_i7).Int())
			}
		}
		var _598_v58 _dafny.Set
		_ = _598_v58
		_598_v58 = _dafny.SetOf(p3)
		r0 = _dafny.Companion_Sequence_.Update(_499_v2, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F7()), _dafny.IntOfUint32((_499_v2).Cardinality()))).Uint32(), ((_this).F7()).Cmp((_598_v58).Cardinality()) != 0)
		var _599_v59 _dafny.Set
		_ = _599_v59
		_599_v59 = _dafny.SetOf(p0)
		var _600_v60 D3
		_ = _600_v60
		_600_v60 = Companion_D3_.Create_DC6_(_599_v59)
		var _pat_let_tv9 = p0
		_ = _pat_let_tv9
		var _pat_let_tv10 = p2
		_ = _pat_let_tv10
		r1 = func(_source13 D3) _dafny.Int {
			if _source13.Is_DC7() {
				var _601___mcc_h12 bool = _source13.Get_().(D3_DC7).Cf10
				_ = _601___mcc_h12
				var _602_cf10 bool = _601___mcc_h12
				_ = _602_cf10
				if _602_cf10 {
					return _pat_let_tv9
				} else {
					return (_dafny.MultiSetOf(_pat_let_tv10, _dafny.IntOfInt64(-808), (_this).F8(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fjemwxaw")).Cardinality()), (_this).F8())).Cardinality()
				}
			} else {
				var _603___mcc_h13 _dafny.Set = _source13.Get_().(D3_DC6).Cf9
				_ = _603___mcc_h13
				var _604_cf9 _dafny.Set = _603___mcc_h13
				_ = _604_cf9
				return _dafny.IntOfInt64(332)
			}
		}(_600_v60)
		var _605_v61 _dafny.Sequence
		_ = _605_v61
		_605_v61 = _dafny.UnicodeSeqOfUtf8Bytes("rtyih")
		r2 = (Companion_Default___.Fm20(false, p3, p1, globalState)).Cmp(_dafny.IntOfUint32((_605_v61).Cardinality())) != 0
		return r0, r1, r2
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f7  _dafny.Int
	_f8  _dafny.Int
	_f12 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.Zero
	_this._f12 = false
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T2 = &C3{}
var _ T0 = &C3{}
var _ T1 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F7() _dafny.Int {
	return _this._f7
}
func (_this *C3) F8() _dafny.Int {
	return _this._f8
}
func (_this *C3) Ctor__(f12 bool, f7 _dafny.Int, f8 _dafny.Int) {
	{
		(_this)._f12 = f12
		(_this)._f7 = f7
		(_this)._f8 = f8
	}
}
func (_this *C3) Fm2(p0 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("skp")
	}
}
func (_this *C3) Fm0(p0 bool, globalState *GlobalState) _dafny.Set {
	{
		var _source14 bool = (_this).F12()
		_ = _source14
		var _606___mcc_h0 bool = _source14
		_ = _606___mcc_h0
		var _607_cf0 bool = _606___mcc_h0
		_ = _607_cf0
		if (_this).F12() {
			return _dafny.SetOf(_dafny.SetOf(false, _607_cf0))
		} else {
			return _dafny.SetOf(_dafny.SetOf(_607_cf0), _dafny.SetOf((_this).F12()))
		}
	}
}
func (_this *C3) Fm1(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return (Companion_Default___.SafeModuloInt((_dafny.SetOf((_this).F12(), (_this).F12(), (_this).F12(), (_this).F12(), !((_this).F12()))).Cardinality(), (func() _dafny.Map {
			var _coll31 = _dafny.NewMapBuilder()
			_ = _coll31
			for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(234), _dafny.IntOfInt64(342))); ; {
				_compr_31, _ok34 := _iter34()
				if !_ok34 {
					break
				}
				var _608_v0 _dafny.Int
				_608_v0 = interface{}(_compr_31).(_dafny.Int)
				if ((_dafny.IntOfInt64(234)).Cmp(_608_v0) <= 0) && ((_608_v0).Cmp(_dafny.IntOfInt64(342)) < 0) {
					_coll31.Add((_608_v0).Minus(_dafny.IntOfInt64(906)), (_this).F12())
				}
			}
			return _coll31.ToMap()
		}()).Cardinality())).Cmp(((_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F8(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xutwsx")).Cardinality()), (_this).F8()))).Cardinality()).Minus((_this).F8())) >= 0
	}
}
func (_this *C3) Fm12(globalState *GlobalState) bool {
	{
		return (_this).F12()
	}
}
func (_this *C3) Fm13(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F8()
	}
}
func (_this *C3) M1(p0 _dafny.Sequence, globalState *GlobalState) {
	{
		var _609_v0 _dafny.Sequence
		_ = _609_v0
		_609_v0 = _dafny.SeqOf((_this).F12(), !((_this).F12()), (_this).F12())
		var _hi4 _dafny.Int = (_dafny.Zero).Minus((_this).F7())
		_ = _hi4
		for _610_i0 := (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_609_v0, _609_v0)).Cardinality())); _610_i0.Cmp(_hi4) < 0; _610_i0 = _610_i0.Plus(_dafny.One) {
			var _611_v1 bool
			_ = _611_v1
			_611_v1 = (_this).F12()
			if !((_610_i0).Cmp((_dafny.MultiSetOf(_611_v1)).Cardinality()) > 0) {
				var _612_v2 _dafny.Array
				_ = _612_v2
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_12
				var _nw100 _dafny.Array
				_ = _nw100
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw100 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) _dafny.Int = (func(_613_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_614_i1 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_614_i1, _613_i0)
						}
					})(_610_i0)
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw100 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw100).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw100).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_612_v2 = _nw100
				var _615_v3 _dafny.MultiSet
				_ = _615_v3
				_615_v3 = _dafny.MultiSetOf((_this).F7())
				var _616_v4 _dafny.Map
				_ = _616_v4
				_616_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_612_v2, _615_v3)
				var _617_v5 _dafny.Array
				_ = _617_v5
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_13
				var _nw101 _dafny.Array
				_ = _nw101
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw101 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) bool = func(_618_i2 _dafny.Int) bool {
						return !(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(735))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg45 _dafny.Int) interface{} {
								return coer45(arg45)
							}
						}(func(_619_i3 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('b')
						})))).Equals(_dafny.MultiSetOf(_dafny.CodePoint('e')))
					}
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
				_617_v5 = _nw101
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_617_v5), 0))
				_ = _index107
				(_617_v5).ArraySet1((_this).F12(), (_index107).Int())
				var _620_v6 _dafny.Set
				_ = _620_v6
				_620_v6 = _dafny.SetOf((_this).F7(), (_this).F8())
				var _621_v7 _dafny.Sequence
				_ = _621_v7
				_621_v7 = _dafny.SeqOf((_620_v6).Cardinality())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_617_v5), 0))
				_ = _index108
				var _rhs97 _dafny.Map = (_616_v4).Merge(_616_v4)
				_ = _rhs97
				var _rhs98 bool = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_621_v7, _dafny.SeqOf((_this).F7())), _dafny.Companion_Sequence_.Concatenate(_621_v7, _621_v7))
				_ = _rhs98
				var _rhs99 bool = (_this).F12()
				_ = _rhs99
				var _rhs100 bool = (_this).F12()
				_ = _rhs100
				var _lhs81 *GlobalState = globalState
				_ = _lhs81
				var _lhs82 *GlobalState = globalState
				_ = _lhs82
				var _lhs83 _dafny.Array = _617_v5
				_ = _lhs83
				var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_617_v5), 0))
				_ = _lhs84
				_616_v4 = _rhs97
				_lhs81.F3 = _rhs98
				_lhs82.F4 = _rhs99
				(_lhs83).ArraySet1(_rhs100, (_lhs84).Int())
				var _622_v8 _dafny.CodePoint
				_ = _622_v8
				_622_v8 = _dafny.CodePoint('r')
				var _623_v9 D1
				_ = _623_v9
				_623_v9 = Companion_D1_.Create_DC1_((Companion_D1_.Create_DC1_(_dafny.CodePoint('j'))).Dtor_cf1())
				_622_v8 = (_623_v9).Dtor_cf1()
				(globalState).F3 = (_617_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_617_v5), 0))).Int()).(bool)
				(globalState).F3 = (_609_v0).Select((Companion_Default___.SafeIndex(_610_i0, _dafny.IntOfUint32((_609_v0).Cardinality()))).Uint32()).(bool)
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_617_v5), 0))
				_ = _index109
				(_617_v5).ArraySet1((_this).F12(), (_index109).Int())
			} else {
				(globalState).F4 = ((_this).F12()) && ((func() bool {
					if (_this).F12() {
						return (_this).F12()
					}
					return (_this).F12()
				})())
				var _624_v10 _dafny.MultiSet
				_ = _624_v10
				_624_v10 = _dafny.MultiSetOf(_609_v0)
				(globalState).F3 = (_624_v10).IsProperSubsetOf((_624_v10).Difference(_624_v10))
				var _625_v11 _dafny.Array
				_ = _625_v11
				var _len0_14 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_14
				var _nw102 _dafny.Array
				_ = _nw102
				if _len0_14.Cmp(_dafny.Zero) == 0 {
					_nw102 = _dafny.NewArray(_len0_14)
				} else {
					var _init14 func(_dafny.Int) bool = func(_626_i4 _dafny.Int) bool {
						return (_this).F12()
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
				_625_v11 = _nw102
				_625_v11 = _625_v11
				var _627_v12 _dafny.Sequence
				_ = _627_v12
				_627_v12 = _dafny.SeqOf(_609_v0)
				_627_v12 = _627_v12
				var _628_v13 _dafny.Int
				_ = _628_v13
				_628_v13 = _dafny.IntOfInt64(658)
				_628_v13 = ((_this).F7()).Plus(Companion_Default___.SafeModuloInt((_this).F8(), (_this).F7()))
			}
			var _629_v14 _dafny.CodePoint
			_ = _629_v14
			_629_v14 = _dafny.CodePoint('g')
			_629_v14 = _629_v14
			(globalState).F4 = (_this).F12()
			var _630_v15 _dafny.Map
			_ = _630_v15
			_630_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), Companion_Default___.SafeModuloInt(_610_i0, (_dafny.MultiSetOf((_this).F12())).Cardinality()))
			_630_v15 = _630_v15
		}
		var _631_v16 _dafny.MultiSet
		_ = _631_v16
		_631_v16 = _dafny.MultiSetOf((_this).F8())
		var _632_v17 _dafny.Array
		_ = _632_v17
		var _633_v18 bool
		_ = _633_v18
		var _out2 _dafny.Array
		_ = _out2
		var _out3 bool
		_ = _out3
		_out2, _out3 = Companion_Default___.M0((_this).F7(), (_this).Fm1(_dafny.SetOf((_this).F12(), (_this).F12()), _dafny.IntOfInt64(-562), (_this).F7(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(893))).Cardinality(), (_this).F8())).Cardinality(), globalState), (_631_v16).IsSubsetOf(_631_v16), globalState)
		_632_v17 = _out2
		_633_v18 = _out3
		var _634_v19 _dafny.CodePoint
		_ = _634_v19
		_634_v19 = _dafny.CodePoint('b')
		var _635_v20 _dafny.Map
		_ = _635_v20
		_635_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm12(globalState), _634_v19)
		_635_v20 = _635_v20
		var _636_i5 _dafny.Int
		_ = _636_i5
		_636_i5 = _dafny.Zero
		{
			for (_this).F12() {
				{
					if (_636_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_636_i5 = (_636_i5).Plus(_dafny.One)
					var _637_v21 _dafny.Sequence
					_ = _637_v21
					_637_v21 = _dafny.UnicodeSeqOfUtf8Bytes("xronwep")
					var _rhs101 _dafny.Sequence = p0
					_ = _rhs101
					var _rhs102 bool = (_this).F12()
					_ = _rhs102
					_637_v21 = _rhs101
					_633_v18 = _rhs102
					_633_v18 = (_this).F12()
					var _638_v22 _dafny.Int
					_ = _638_v22
					_638_v22 = _dafny.IntOfInt64(771)
					var _639_v23 _dafny.Array
					_ = _639_v23
					var _nw103 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
					_ = _nw103
					_639_v23 = _nw103
					var _640_v24 D1
					_ = _640_v24
					_640_v24 = Companion_D1_.Create_DC1_(_634_v19)
					var _641_v25 _dafny.MultiSet
					_ = _641_v25
					_641_v25 = _dafny.MultiSetOf(_640_v24)
					var _642_v26 _dafny.Sequence
					_ = _642_v26
					_642_v26 = _dafny.SeqOf(_641_v25)
					var _rhs103 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(643), _dafny.IntOfInt64(663)))
					_ = _rhs103
					var _rhs104 _dafny.Array = _639_v23
					_ = _rhs104
					var _rhs105 _dafny.CodePoint = _634_v19
					_ = _rhs105
					var _rhs106 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p0, _637_v21)
					_ = _rhs106
					var _rhs107 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(763))).Uint32(), func(coer46 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
						return func(arg46 _dafny.Int) interface{} {
							return coer46(arg46)
						}
					}((func(_643_v24 D1) func(_dafny.Int) _dafny.MultiSet {
						return func(_644_i6 _dafny.Int) _dafny.MultiSet {
							return _dafny.MultiSetFromSeq(_dafny.SeqOf(_643_v24, _643_v24))
						}
					})(_640_v24))), _642_v26)
					_ = _rhs107
					_638_v22 = _rhs103
					_639_v23 = _rhs104
					_634_v19 = _rhs105
					_637_v21 = _rhs106
					_642_v26 = _rhs107
					var _645_v27 _dafny.Map
					_ = _645_v27
					_645_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_633_v18, _633_v18)
					var _646_v28 _dafny.Map
					_ = _646_v28
					_646_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_645_v27).Merge(_645_v27), _639_v23)
					_646_v28 = (_646_v28).Update((_645_v27).Merge(_645_v27), _639_v23)
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _647_v29 _dafny.MultiSet
		_ = _647_v29
		_647_v29 = _dafny.MultiSetOf(false)
		var _648_v30 _dafny.Sequence
		_ = _648_v30
		_648_v30 = _dafny.SeqOf((_647_v29).Cardinality())
		var _hi5 _dafny.Int = ((_this).F7()).Times((_this).F7())
		_ = _hi5
		for _649_i7 := (_648_v30).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_648_v30).Cardinality()))).Uint32()).(_dafny.Int); _649_i7.Cmp(_hi5) < 0; _649_i7 = _649_i7.Plus(_dafny.One) {
			var _650_v31 _dafny.Array
			_ = _650_v31
			var _651_v32 bool
			_ = _651_v32
			var _out4 _dafny.Array
			_ = _out4
			var _out5 bool
			_ = _out5
			_out4, _out5 = Companion_Default___.M0(((_this).F7()).Plus((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("vvjdbcbj"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-435))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}((func(_652_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_653_i8 _dafny.Int) _dafny.CodePoint {
					return _652_v19
				}
			})(_634_v19))))).Cardinality())), _633_v18, (_633_v18) == (true), globalState)
			_650_v31 = _out4
			_651_v32 = _out5
			(globalState).F3 = (_651_v32)
			_609_v0 = _dafny.Companion_Sequence_.Concatenate(_609_v0, _dafny.Companion_Sequence_.Concatenate(_609_v0, _609_v0))
			var _rhs108 _dafny.Sequence = _648_v30
			_ = _rhs108
			var _rhs109 bool = (_this).F12()
			_ = _rhs109
			var _lhs85 *GlobalState = globalState
			_ = _lhs85
			_648_v30 = _rhs108
			_lhs85.F3 = _rhs109
		}
		var _654_v33 _dafny.Sequence
		_ = _654_v33
		_654_v33 = _dafny.SeqOf((func() _dafny.Sequence {
			if _633_v18 {
				return _609_v0
			}
			return _609_v0
		})())
		_609_v0 = (_654_v33).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_this).F8(), (_this).F7()), _dafny.IntOfUint32((_654_v33).Cardinality()))).Uint32()).(_dafny.Sequence)
	}
}
func (_this *C3) M2(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _655_v0 _dafny.Sequence
		_ = _655_v0
		_655_v0 = _dafny.UnicodeSeqOfUtf8Bytes("dwwkgew")
		r1 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(267), _dafny.IntOfUint32((_655_v0).Cardinality()))
		var _656_v1 _dafny.CodePoint
		_ = _656_v1
		_656_v1 = _dafny.CodePoint('m')
		var _657_v2 _dafny.Array
		_ = _657_v2
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_15
		var _nw104 _dafny.Array
		_ = _nw104
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw104 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) bool = (func(_658_p2 _dafny.Int) func(_dafny.Int) bool {
				return func(_659_i0 _dafny.Int) bool {
					return ((_dafny.Zero).Minus(_658_p2)).Cmp((_this).F8()) > 0
				}
			})(p2)
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw104 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw104).ArraySet1(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw104).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_657_v2 = _nw104
		var _rhs110 _dafny.CodePoint = _dafny.CodePoint('y')
		_ = _rhs110
		var _rhs111 _dafny.Array = _657_v2
		_ = _rhs111
		var _rhs112 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_655_v0).Cardinality()))
		_ = _rhs112
		var _rhs113 _dafny.Int = _dafny.IntOfInt64(255)
		_ = _rhs113
		_656_v1 = _rhs110
		_657_v2 = _rhs111
		r1 = _rhs112
		r1 = _rhs113
		r2 = p3
		var _660_v3 _dafny.Sequence
		_ = _660_v3
		_660_v3 = _dafny.SeqOf((_this).F12(), true, p1)
		var _661_v4 _dafny.MultiSet
		_ = _661_v4
		_661_v4 = _dafny.MultiSetOf(!(p1), (_660_v3).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_660_v3).Cardinality()))).Uint32()).(bool))
		var _662_v5 D2
		_ = _662_v5
		_662_v5 = Companion_D2_.Create_DC4_(_661_v4)
		var _pat_let_tv11 = _661_v4
		_ = _pat_let_tv11
		_661_v4 = (func(_pat_let7_0 D2) D2 {
			return func(_663_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let8_0 _dafny.MultiSet) D2 {
					return func(_664_dt__update_hcf3_h0 _dafny.MultiSet) D2 {
						return Companion_D2_.Create_DC4_(_664_dt__update_hcf3_h0)
					}(_pat_let8_0)
				}(_pat_let_tv11)
			}(_pat_let7_0)
		}(_662_v5)).Dtor_cf3()
		var _hi6 _dafny.Int = (_dafny.Zero).Minus((_this).F7())
		_ = _hi6
		for _665_i1 := p2; _665_i1.Cmp(_hi6) < 0; _665_i1 = _665_i1.Plus(_dafny.One) {
			r1 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(-574)), _dafny.IntOfInt64(444)))
			var _666_v6 bool
			_ = _666_v6
			_666_v6 = (_this).F12()
			var _source15 bool = _666_v6
			_ = _source15
			var _667___mcc_h0 bool = _source15
			_ = _667___mcc_h0
			var _668_cf0 bool = _667___mcc_h0
			_ = _668_cf0
			r1 = p0
			r1 = (_this).F8()
			var _669_v7 _dafny.Map
			_ = _669_v7
			_669_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_this).Fm13(_655_v0, globalState))
			r1 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_669_v7).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1), _660_v3)).Cardinality()))))
			var _670_v8 T2
			_ = _670_v8
			var _nw105 *C2 = New_C2_()
			_ = _nw105
			_nw105.Ctor__(p2, p2)
			_670_v8 = _nw105
			_670_v8 = _670_v8
			var _671_v9 D5
			_ = _671_v9
			_671_v9 = Companion_D5_.Create_DC12_(!((_this).Fm12(globalState)))
			var _672_v10 _dafny.MultiSet
			_ = _672_v10
			_672_v10 = _dafny.MultiSetOf(_dafny.IntOfInt64(-703), p0, _dafny.IntOfUint32((_655_v0).Cardinality()))
			var _673_v11 _dafny.Sequence
			_ = _673_v11
			_673_v11 = _dafny.SeqOf(_665_i1)
			var _rhs114 bool = ((_this).F7()).Cmp(_dafny.IntOfUint32((Companion_Default___.Fm30(p3, globalState)).Cardinality())) < 0
			_ = _rhs114
			var _rhs115 bool = !((p3) || ((_this).F12()))
			_ = _rhs115
			var _rhs116 D5 = _671_v9
			_ = _rhs116
			var _rhs117 _dafny.MultiSet = (_672_v10).Difference((_672_v10).Union(_dafny.MultiSetFromSeq(_673_v11)))
			_ = _rhs117
			var _lhs86 *GlobalState = globalState
			_ = _lhs86
			r2 = _rhs114
			_lhs86.F3 = _rhs115
			_671_v9 = _rhs116
			_672_v10 = _rhs117
			(globalState).F4 = !((func() bool {
				if p1 {
					return false
				}
				return !(p1)
			})())
		}
		var _674_v12 _dafny.Array
		_ = _674_v12
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_16
		var _nw106 _dafny.Array
		_ = _nw106
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw106 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) _dafny.Int = (func(_675_v3 _dafny.Sequence, _676_p1 bool) func(_dafny.Int) _dafny.Int {
				return func(_677_i3 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_677_i3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_675_v3, (Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_675_v3).Cardinality()))).Uint32(), _676_p1)).Cardinality()))
				}
			})(_660_v3, p1)
			_ = _init16
			var _element0_16 = _init16(_dafny.Zero)
			_ = _element0_16
			_nw106 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
			(_nw106).ArraySet1(_element0_16, 0)
			var _nativeLen0_16 = (_len0_16).Int()
			_ = _nativeLen0_16
			for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
				(_nw106).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
			}
		}
		_674_v12 = _nw106
		for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_674_v12), 0))); ; {
			_guard_loop_3, _ok35 := _iter35()
			if !_ok35 {
				break
			}
			var _678_i2 _dafny.Int
			_678_i2 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_678_i2).Sign() != -1) && ((_678_i2).Cmp(_dafny.ArrayLen((_674_v12), 0)) < 0)) {
				(_674_v12).ArraySet1(Companion_Default___.SafeModuloInt(_678_i2, (func() _dafny.Int {
					if (_this).F12() {
						return p2
					}
					return (_dafny.Zero).Minus((_this).F8())
				})()), (_678_i2).Int())
			}
		}
		r0 = _660_v3
		r1 = (p2).Times((_this).F7())
		r2 = p3
		return r0, r1, r2
	}
}
func (_this *C3) F12() bool {
	{
		return _this._f12
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f7  _dafny.Int
	_f8  _dafny.Int
	_f11 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.Zero
	_this._f11 = false
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

func (_this *C4) F7() _dafny.Int {
	return _this._f7
}
func (_this *C4) F8() _dafny.Int {
	return _this._f8
}
func (_this *C4) Ctor__(f11 bool, f7 _dafny.Int, f8 _dafny.Int) {
	{
		(_this)._f11 = f11
		(_this)._f7 = f7
		(_this)._f8 = f8
	}
}
func (_this *C4) Fm0(p0 bool, globalState *GlobalState) _dafny.Set {
	{
		return ((_dafny.SetOf(_dafny.SetOf((_this).F11()))).Union(_dafny.SetOf(_dafny.SetOf((_this).F11(), (_this).F11()), _dafny.SetOf((_this).F11(), false, false, (_this).F11())))).Union((_dafny.SetOf(_dafny.SetOf((_this).F11()))).Union(func() _dafny.Set {
			var _coll32 = _dafny.NewBuilder()
			_ = _coll32
			for _iter36 := _dafny.Iterate((func() _dafny.Set {
				var _coll33 = _dafny.NewBuilder()
				_ = _coll33
				for _iter37 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SetOf((_this).F11()), _dafny.SetOf((_this).F11(), (_this).F11(), (_this).F11()))).Elements()); ; {
					_compr_33, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _679_v0 _dafny.Set
					_679_v0 = interface{}(_compr_33).(_dafny.Set)
					if (_dafny.MultiSetOf(_dafny.SetOf((_this).F11()), _dafny.SetOf((_this).F11(), (_this).F11(), (_this).F11()))).Contains(_679_v0) {
						_coll33.Add(_679_v0)
					}
				}
				return _coll33.ToSet()
			}()).Elements()); ; {
				_compr_32, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _680_v1 _dafny.Set
				_680_v1 = interface{}(_compr_32).(_dafny.Set)
				if (func() _dafny.Set {
					var _coll34 = _dafny.NewBuilder()
					_ = _coll34
					for _iter38 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SetOf((_this).F11()), _dafny.SetOf((_this).F11(), (_this).F11(), (_this).F11()))).Elements()); ; {
						_compr_34, _ok38 := _iter38()
						if !_ok38 {
							break
						}
						var _681_v0 _dafny.Set
						_681_v0 = interface{}(_compr_34).(_dafny.Set)
						if (_dafny.MultiSetOf(_dafny.SetOf((_this).F11()), _dafny.SetOf((_this).F11(), (_this).F11(), (_this).F11()))).Contains(_681_v0) {
							_coll34.Add(_681_v0)
						}
					}
					return _coll34.ToSet()
				}()).Contains(_680_v1) {
					_coll32.Add(_680_v1)
				}
			}
			return _coll32.ToSet()
		}()))
	}
}
func (_this *C4) Fm1(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(480))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg48 _dafny.Int) interface{} {
				return coer48(arg48)
			}
		}(func(_682_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		}))).Cardinality())).Cmp((_this).F7()) < 0
	}
}
func (_this *C4) Fm11(globalState *GlobalState) _dafny.Int {
	{
		if (_this).F11() {
			return (_this).F8()
		} else {
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F7(), (_this).F7()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(702))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}(func(_683_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(453)
			})))).Cardinality())
		}
	}
}
func (_this *C4) M5(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (bool, bool, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		r1 = p0
		var _684_v0 *C2
		_ = _684_v0
		var _nw107 *C2 = New_C2_()
		_ = _nw107
		_nw107.Ctor__(_dafny.IntOfInt64(537), _dafny.IntOfInt64(-183))
		_684_v0 = _nw107
		var _685_v1 *C2
		_ = _685_v1
		var _nw108 *C2 = New_C2_()
		_ = _nw108
		_nw108.Ctor__(p2, Companion_Default___.SafeDivisionInt((_this).F7(), (_this).F7()))
		_685_v1 = _nw108
		var _686_i0 _dafny.Int
		_ = _686_i0
		_686_i0 = _dafny.Zero
		{
			for (func() bool {
				if (_this).F11() {
					return p0
				}
				return p0
			})() {
				{
					if (_686_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_686_i0 = (_686_i0).Plus(_dafny.One)
					var _687_v2 _dafny.Int
					_ = _687_v2
					_687_v2 = _dafny.IntOfInt64(241)
					_687_v2 = p2
					var _688_v3 *C1
					_ = _688_v3
					var _nw109 *C1 = New_C1_()
					_ = _nw109
					_nw109.Ctor__(Companion_Default___.SafeModuloInt((_this).F7(), _dafny.IntOfInt64(83)), Companion_Default___.SafeModuloInt((_this).F7(), Companion_Default___.Fm16(true, globalState)))
					_688_v3 = _nw109
					var _689_v4 _dafny.Array
					_ = _689_v4
					var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
					_ = _nw110
					_689_v4 = _nw110
					_689_v4 = _689_v4
					var _690_v5 _dafny.Map
					_ = _690_v5
					_690_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
					var _691_v6 _dafny.Map
					_ = _691_v6
					_691_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_690_v5, _dafny.SeqOf(p0, (_this).F11()))
					_687_v2 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p2), Companion_Default___.SafeDivisionInt(_687_v2, (_691_v6).Cardinality()))
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _source16 D3 = Companion_Default___.Fm26(_dafny.MultiSetOf(p2), (_this).F11(), globalState)
		_ = _source16
		if _source16.Is_DC7() {
			var _692___mcc_h0 bool = _source16.Get_().(D3_DC7).Cf10
			_ = _692___mcc_h0
			var _693_cf10 bool = _692___mcc_h0
			_ = _693_cf10
			var _694_v7 _dafny.Int
			_ = _694_v7
			_694_v7 = _dafny.IntOfInt64(184)
			var _695_v8 _dafny.Sequence
			_ = _695_v8
			_695_v8 = _dafny.SeqOf((_this).F11(), false)
			var _696_v9 _dafny.Sequence
			_ = _696_v9
			_696_v9 = _dafny.SeqOf(_695_v8, _695_v8)
			_694_v7 = (Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F8()), _694_v7)).Times(_dafny.IntOfUint32(((_696_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lhbhuxmc")).Cardinality()), _dafny.IntOfUint32((_696_v9).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
			var _697_v10 _dafny.MultiSet
			_ = _697_v10
			_697_v10 = _dafny.MultiSetOf((_694_v7).Times(p3))
			var _698_v11 _dafny.Sequence
			_ = _698_v11
			_698_v11 = _dafny.SeqOf(p1)
			_694_v7 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_697_v10).Contains(p2) {
					return (_697_v10).Multiplicity(p2)
				}
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_698_v11, _698_v11)).Cardinality())
			})())
			_698_v11 = _698_v11
			var _699_v12 _dafny.Sequence
			_ = _699_v12
			_699_v12 = _dafny.UnicodeSeqOfUtf8Bytes("md")
			var _700_v13 _dafny.Array
			_ = _700_v13
			var _nwElement0_18 _dafny.Int = (_this).F8()
			_ = _nwElement0_18
			var _nw111 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(10))
			_ = _nw111
			(_nw111).ArraySet1(_nwElement0_18, 0)
			(_nw111).ArraySet1(p1, 1)
			(_nw111).ArraySet1((_this).F8(), 2)
			(_nw111).ArraySet1(_dafny.IntOfInt64(-318), 3)
			(_nw111).ArraySet1(p2, 4)
			(_nw111).ArraySet1(_694_v7, 5)
			(_nw111).ArraySet1(_dafny.IntOfUint32((_699_v12).Cardinality()), 6)
			(_nw111).ArraySet1((_698_v11).Select((Companion_Default___.SafeIndex(_694_v7, _dafny.IntOfUint32((_698_v11).Cardinality()))).Uint32()).(_dafny.Int), 7)
			(_nw111).ArraySet1((_this).F8(), 8)
			(_nw111).ArraySet1(p3, 9)
			_700_v13 = _nw111
			var _701_v14 _dafny.Array
			_ = _701_v14
			var _nwElement0_19 _dafny.Array = _700_v13
			_ = _nwElement0_19
			var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.One)
			_ = _nw112
			(_nw112).ArraySet1(_nwElement0_19, 0)
			_701_v14 = _nw112
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_701_v14), 0))
			_ = _index110
			(_701_v14).ArraySet1(_700_v13, (_index110).Int())
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_701_v14), 0))
			_ = _index111
			(_701_v14).ArraySet1(_700_v13, (_index111).Int())
		} else {
			var _702___mcc_h1 _dafny.Set = _source16.Get_().(D3_DC6).Cf9
			_ = _702___mcc_h1
			var _703_cf9 _dafny.Set = _702___mcc_h1
			_ = _703_cf9
			var _704_v15 _dafny.Sequence
			_ = _704_v15
			_704_v15 = _dafny.UnicodeSeqOfUtf8Bytes("u")
			var _705_v16 _dafny.Set
			_ = _705_v16
			_705_v16 = _dafny.SetOf(_704_v15)
			var _706_v17 _dafny.Array
			_ = _706_v17
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_17
			var _nw113 _dafny.Array
			_ = _nw113
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw113 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) bool = (func(_707_p0 bool) func(_dafny.Int) bool {
					return func(_708_i1 _dafny.Int) bool {
						return _707_p0
					}
				})(p0)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw113 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw113).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw113).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_706_v17 = _nw113
			var _709_v18 _dafny.Map
			_ = _709_v18
			_709_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v15, _706_v17)
			var _710_v19 _dafny.Int
			_ = _710_v19
			_710_v19 = _dafny.IntOfInt64(-109)
			var _711_v20 _dafny.Sequence
			_ = _711_v20
			_711_v20 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ekpn"), _706_v17), _709_v18)
			var _712_v21 _dafny.Set
			_ = _712_v21
			_712_v21 = _dafny.SetOf(p0)
			var _713_v22 _dafny.MultiSet
			_ = _713_v22
			_713_v22 = _dafny.MultiSetOf((_this).F11())
			var _rhs118 _dafny.Set = (_705_v16).Difference(_705_v16)
			_ = _rhs118
			var _rhs119 _dafny.Map = (_711_v20).Select((Companion_Default___.SafeIndex((_712_v21).Cardinality(), _dafny.IntOfUint32((_711_v20).Cardinality()))).Uint32()).(_dafny.Map)
			_ = _rhs119
			var _rhs120 _dafny.Int = ((_713_v22).Intersection(_713_v22)).Cardinality()
			_ = _rhs120
			_705_v16 = _rhs118
			_709_v18 = _rhs119
			_710_v19 = _rhs120
			r3 = !((_this).F11())
			var _714_v23 _dafny.Sequence
			_ = _714_v23
			_714_v23 = _dafny.SeqOf(p0, p0)
			var _715_v24 _dafny.Sequence
			_ = _715_v24
			_715_v24 = _dafny.SeqOf(_dafny.IntOfInt64(80), (_this).F7(), p1)
			(globalState).F4 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_714_v23, _714_v23), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_715_v24).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_714_v23, _714_v23)).Cardinality()))).Uint32(), (_this).F11()), _714_v23)
			_714_v23 = _dafny.Companion_Sequence_.Update(_714_v23, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.IntOfUint32((_714_v23).Cardinality()))).Uint32(), p0)
		}
		var _716_v25 _dafny.Int
		_ = _716_v25
		_716_v25 = _dafny.IntOfInt64(321)
		_716_v25 = p2
		r0 = p0
		r1 = (_this).F11()
		r2 = !(p0) || ((_this).F11())
		r3 = !((_this).F11())
		return r0, r1, r2, r3
	}
}
func (_this *C4) F11() bool {
	{
		return _this._f11
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f7  _dafny.Int
	_f8  _dafny.Int
	_f9  _dafny.Sequence
	_f10 _dafny.Int
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.Zero
	_this._f9 = _dafny.EmptySeq
	_this._f10 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C5{}
var _ T1 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F7() _dafny.Int {
	return _this._f7
}
func (_this *C5) F8() _dafny.Int {
	return _this._f8
}
func (_this *C5) Ctor__(f9 _dafny.Sequence, f10 _dafny.Int, f7 _dafny.Int, f8 _dafny.Int) {
	{
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f7 = f7
		(_this)._f8 = f8
	}
}
func (_this *C5) Fm0(p0 bool, globalState *GlobalState) _dafny.Set {
	{
		return ((_dafny.SetOf(_dafny.SetOf(false, false))).Difference(_dafny.SetOf(_dafny.SetOf(true, true), _dafny.SetOf(false, true)))).Union(_dafny.SetOf(_dafny.SetOf(false), _dafny.SetOf(true, false, !(true), false, true)))
	}
}
func (_this *C5) Fm1(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return (false) && (!_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F7())).Cardinality()), (_this).F10(), (_this).F8(), (_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F10(), (_this).F10()))).Cardinality(), _dafny.IntOfUint32(((_this).F9()).Cardinality())), _dafny.SetOf((_this).F10())), _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(64)))))
	}
}
func (_this *C5) Fm6(p0 _dafny.CodePoint, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F8()
	}
}
func (_this *C5) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F10()).Times((_this).F7())
	}
}
func (_this *C5) M4(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _717_v0 _dafny.Array
		_ = _717_v0
		var _len0_18 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_18
		var _nw114 _dafny.Array
		_ = _nw114
		if _len0_18.Cmp(_dafny.Zero) == 0 {
			_nw114 = _dafny.NewArray(_len0_18)
		} else {
			var _init18 func(_dafny.Int) _dafny.Int = func(_718_i0 _dafny.Int) _dafny.Int {
				return (_718_i0).Times((_this).F7())
			}
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
		_717_v0 = _nw114
		var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
		_ = _index112
		(_717_v0).ArraySet1((_this).F7(), (_index112).Int())
		var _719_v1 bool
		_ = _719_v1
		_719_v1 = false
		var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
		_ = _index113
		(_717_v0).ArraySet1((_this).Fm6(_dafny.CodePoint('a'), Companion_Default___.Fm8(_719_v1, _dafny.SeqOf((_this).F8(), p0, p0), _719_v1, globalState), _719_v1, globalState), (_index113).Int())
		var _720_v2 _dafny.Sequence
		_ = _720_v2
		_720_v2 = _dafny.SeqOf(_719_v1, _719_v1, _719_v1)
		var _721_v3 _dafny.Map
		_ = _721_v3
		_721_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _719_v1)
		var _722_v4 _dafny.CodePoint
		_ = _722_v4
		_722_v4 = _dafny.CodePoint('g')
		var _723_v5 _dafny.Map
		_ = _723_v5
		_723_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32(((_this).F9()).Cardinality()))
		var _724_v6 bool
		_ = _724_v6
		_724_v6 = _719_v1
		var _725_v7 _dafny.Set
		_ = _725_v7
		_725_v7 = _dafny.SetOf((_this).F8(), _dafny.IntOfUint32(((_this).F9()).Cardinality()))
		var _726_v8 _dafny.Map
		_ = _726_v8
		_726_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v1, false)
		var _727_v9 _dafny.MultiSet
		_ = _727_v9
		_727_v9 = _dafny.MultiSetOf(Companion_Default___.Fm9(_719_v1, globalState))
		var _728_v10 _dafny.Set
		_ = _728_v10
		_728_v10 = _dafny.SetOf(_719_v1)
		var _729_v11 _dafny.MultiSet
		_ = _729_v11
		_729_v11 = _dafny.MultiSetOf(_719_v1, _719_v1)
		var _730_v12 _dafny.Sequence
		_ = _730_v12
		_730_v12 = _dafny.SeqOf((_this).F7())
		var _731_v13 _dafny.Map
		_ = _731_v13
		_731_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_730_v12).Cardinality())))).Update(p0, Companion_Default___.Abs((_this).F7()))).Update(p0, Companion_Default___.Abs((_this).F8())), _719_v1)
		var _732_v14 _dafny.Array
		_ = _732_v14
		var _nwElement0_20 bool = _719_v1
		_ = _nwElement0_20
		var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(27))
		_ = _nw115
		(_nw115).ArraySet1(_nwElement0_20, 0)
		(_nw115).ArraySet1(_719_v1, 1)
		(_nw115).ArraySet1((_720_v2).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_720_v2).Cardinality()))).Uint32()).(bool), 2)
		(_nw115).ArraySet1((func() bool {
			if (_721_v3).Contains(_dafny.UnicodeSeqOfUtf8Bytes("twaaoc")) {
				return (_721_v3).Get(_dafny.UnicodeSeqOfUtf8Bytes("twaaoc")).(bool)
			}
			return _719_v1
		})(), 3)
		(_nw115).ArraySet1(_719_v1, 4)
		(_nw115).ArraySet1(_719_v1, 5)
		(_nw115).ArraySet1(_719_v1, 6)
		(_nw115).ArraySet1(_719_v1, 7)
		(_nw115).ArraySet1(_719_v1, 8)
		(_nw115).ArraySet1(((_dafny.Zero).Minus((_this).Fm6(_722_v4, (_this).F9(), _719_v1, globalState))).Cmp((_dafny.Zero).Minus((_dafny.MultiSetOf(_719_v1)).Cardinality())) < 0, 9)
		(_nw115).ArraySet1(!(_719_v1), 10)
		(_nw115).ArraySet1(_719_v1, 11)
		(_nw115).ArraySet1(!(_723_v5).Contains(_dafny.IntOfUint32((_720_v2).Cardinality())), 12)
		(_nw115).ArraySet1((_724_v6), 13)
		(_nw115).ArraySet1(_719_v1, 14)
		(_nw115).ArraySet1(_719_v1, 15)
		(_nw115).ArraySet1((_725_v7).IsSubsetOf(_725_v7), 16)
		(_nw115).ArraySet1(!(_727_v9).Contains(_726_v8), 17)
		(_nw115).ArraySet1(_719_v1, 18)
		(_nw115).ArraySet1(_719_v1, 19)
		(_nw115).ArraySet1((_719_v1) || (_719_v1), 20)
		(_nw115).ArraySet1(!((_719_v1) && (_719_v1)), 21)
		(_nw115).ArraySet1((_728_v10).IsDisjointFrom(_728_v10), 22)
		(_nw115).ArraySet1((_729_v11).IsDisjointFrom(_dafny.MultiSetOf(_719_v1, _719_v1)), 23)
		(_nw115).ArraySet1(false, 24)
		(_nw115).ArraySet1(_719_v1, 25)
		(_nw115).ArraySet1((func() bool {
			if (_731_v13).Contains(_dafny.MultiSetOf((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), p0, (_this).F7())) {
				return (_731_v13).Get(_dafny.MultiSetOf((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), p0, (_this).F7())).(bool)
			}
			return !(_719_v1)
		})(), 26)
		_732_v14 = _nw115
		var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))
		_ = _index114
		(_732_v14).ArraySet1((func() bool {
			if _719_v1 {
				return _719_v1
			}
			return !(_719_v1)
		})(), (_index114).Int())
		var _733_v15 _dafny.MultiSet
		_ = _733_v15
		_733_v15 = _dafny.MultiSetOf((_this).F8())
		var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
		_ = _index115
		var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))
		_ = _index116
		var _rhs121 _dafny.Int = (_this).Fm7(Companion_Default___.SafeModuloInt(((_733_v15).Update((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), Companion_Default___.Abs((_this).F7()))).Cardinality(), p0), globalState)
		_ = _rhs121
		var _rhs122 bool = (_720_v2).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_720_v2).Cardinality()))).Uint32()).(bool)
		_ = _rhs122
		var _lhs87 _dafny.Array = _717_v0
		_ = _lhs87
		var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
		_ = _lhs88
		var _lhs89 _dafny.Array = _732_v14
		_ = _lhs89
		var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))
		_ = _lhs90
		(_lhs87).ArraySet1(_rhs121, (_lhs88).Int())
		(_lhs89).ArraySet1(_rhs122, (_lhs90).Int())
		var _734_v16 _dafny.Sequence
		_ = _734_v16
		_734_v16 = _dafny.SeqOf(_717_v0, _717_v0)
		var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
		_ = _index117
		(_717_v0).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_734_v16, _734_v16), (Companion_Default___.SafeIndex(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_732_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))).Int()).(bool), (_732_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))).Int()).(bool))).Cardinality()).Plus((_this).F8()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_734_v16, _734_v16)).Cardinality()))).Uint32(), _717_v0)).Cardinality()), (_index117).Int())
		if _719_v1 {
			var _735_v17 _dafny.Map
			_ = _735_v17
			_735_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_722_v4, _dafny.Companion_Sequence_.Update(Companion_Default___.Fm10((_this).F9(), (_732_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))).Int()).(bool), (_dafny.Zero).Minus(p0), (_this).F7(), globalState), (Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((Companion_Default___.Fm10((_this).F9(), (_732_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))).Int()).(bool), (_dafny.Zero).Minus(p0), (_this).F7(), globalState)).Cardinality()))).Uint32(), _719_v1))
			var _736_v18 _dafny.Sequence
			_ = _736_v18
			_736_v18 = _dafny.SeqOf(_724_v6, _724_v6)
			_735_v17 = (_735_v17).Update(_722_v4, _736_v18)
			_719_v1 = (func() bool {
				if (_726_v8).Contains((_732_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))).Int()).(bool)) {
					return (_726_v8).Get((_732_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))).Int()).(bool)).(bool)
				}
				return !((_732_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))).Int()).(bool))
			})()
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))
			_ = _index118
			(_732_v14).ArraySet1(_719_v1, (_index118).Int())
			var _737_v19 _dafny.Map
			_ = _737_v19
			_737_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v1, _732_v14)
			var _738_v20 _dafny.Array
			_ = _738_v20
			var _nwElement0_21 _dafny.Array = _732_v14
			_ = _nwElement0_21
			var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(15))
			_ = _nw116
			(_nw116).ArraySet1(_nwElement0_21, 0)
			(_nw116).ArraySet1(_732_v14, 1)
			(_nw116).ArraySet1(_732_v14, 2)
			(_nw116).ArraySet1(_732_v14, 3)
			(_nw116).ArraySet1((func() _dafny.Array {
				if (_737_v19).Contains(_719_v1) {
					return (_737_v19).Get(_719_v1).(_dafny.Array)
				}
				return _732_v14
			})(), 4)
			(_nw116).ArraySet1(_732_v14, 5)
			(_nw116).ArraySet1(_732_v14, 6)
			(_nw116).ArraySet1(_732_v14, 7)
			(_nw116).ArraySet1(_732_v14, 8)
			(_nw116).ArraySet1(_732_v14, 9)
			(_nw116).ArraySet1(_732_v14, 10)
			(_nw116).ArraySet1(_732_v14, 11)
			(_nw116).ArraySet1((func() _dafny.Array {
				if (_737_v19).Contains(true) {
					return (_737_v19).Get(true).(_dafny.Array)
				}
				return _732_v14
			})(), 12)
			(_nw116).ArraySet1(_732_v14, 13)
			(_nw116).ArraySet1(_732_v14, 14)
			_738_v20 = _nw116
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_738_v20), 0))
			_ = _index119
			(_738_v20).ArraySet1(_732_v14, (_index119).Int())
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_738_v20), 0))
			_ = _index120
			var _nw117 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
			_ = _nw117
			(_738_v20).ArraySet1(_nw117, (_index120).Int())
			_732_v14 = _dafny.ArrayCastTo((_738_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_738_v20), 0))).Int()))
		} else {
			var _739_v22 _dafny.Set
			_ = _739_v22
			_739_v22 = _dafny.SetOf((_this).F9())
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))
			_ = _index121
			var _rhs123 _dafny.Int = (_this).F7()
			_ = _rhs123
			var _rhs124 bool = (_739_v22).IsProperSubsetOf(func() _dafny.Set {
				var _coll35 = _dafny.NewBuilder()
				_ = _coll35
				for _iter39 := _dafny.Iterate((_dafny.SeqOf(_dafny.Companion_Sequence_.Update((_this).F9(), (Companion_Default___.SafeIndex((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_this).F9()).Cardinality()))).Uint32(), _dafny.CodePoint('f')), (_this).F9())).Elements()); ; {
					_compr_35, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _740_v21 _dafny.Sequence
					_740_v21 = interface{}(_compr_35).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.Companion_Sequence_.Update((_this).F9(), (Companion_Default___.SafeIndex((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_this).F9()).Cardinality()))).Uint32(), _dafny.CodePoint('f')), (_this).F9()), _740_v21) {
						_coll35.Add(_740_v21)
					}
				}
				return _coll35.ToSet()
			}())
			_ = _rhs124
			var _lhs91 _dafny.Array = _732_v14
			_ = _lhs91
			var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))
			_ = _lhs92
			r0 = _rhs123
			(_lhs91).ArraySet1(_rhs124, (_lhs92).Int())
			var _741_v23 *C3
			_ = _741_v23
			var _nw118 *C3 = New_C3_()
			_ = _nw118
			_nw118.Ctor__(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("xs"), (_this).F9()), (_this).F7(), (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int))
			_741_v23 = _nw118
			var _742_v24 *C1
			_ = _742_v24
			var _nw119 *C1 = New_C1_()
			_ = _nw119
			_nw119.Ctor__(((_this).F7()).Plus((_this).F7()), ((_dafny.Zero).Minus(_dafny.IntOfUint32((_730_v12).Cardinality()))).Minus((_this).F7()))
			_742_v24 = _nw119
			if _719_v1 {
				_733_v15 = _dafny.MultiSetOf(_dafny.IntOfInt64(-599), (_this).F8(), (_this).Fm7((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), globalState))
				var _743_v25 *C0
				_ = _743_v25
				var _nw120 *C0 = New_C0_()
				_ = _nw120
				_nw120.Ctor__(_717_v0, _dafny.IntOfInt64(991), (_this).F10())
				_743_v25 = _nw120
				var _744_v26 _dafny.Sequence
				_ = _744_v26
				_744_v26 = _dafny.SeqOf(_743_v25)
				var _745_v27 _dafny.Set
				_ = _745_v27
				_745_v27 = _dafny.SetOf(_744_v26)
				(globalState).F3 = (Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_745_v27).Cardinality()), p0)).Cmp(Companion_Default___.SafeModuloInt(Companion_Default___.Fm16(_719_v1, globalState), (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int))) >= 0
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
				_ = _index122
				(_717_v0).ArraySet1(Companion_Default___.SafeModuloInt((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(965)), (_index122).Int())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
				_ = _index123
				(_717_v0).ArraySet1(_dafny.IntOfInt64(602), (_index123).Int())
				var _746_v28 *C4
				_ = _746_v28
				var _nw121 *C4 = New_C4_()
				_ = _nw121
				_nw121.Ctor__((_741_v23).F12(), (_this).F10(), p0)
				_746_v28 = _nw121
			} else {
				_730_v12 = _dafny.Companion_Sequence_.Concatenate(_730_v12, _730_v12)
				var _747_v29 *C1
				_ = _747_v29
				var _nw122 *C1 = New_C1_()
				_ = _nw122
				_nw122.Ctor__((_this).F8(), (_dafny.Zero).Minus((_this).Fm7((_this).F8(), globalState)))
				_747_v29 = _nw122
				var _748_v30 _dafny.Map
				_ = _748_v30
				_748_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_741_v23).F12(), _dafny.IntOfUint32((_720_v2).Cardinality()))
				(globalState).F4 = !((_748_v30).Update((_741_v23).F12(), (_this).F10())).Equals(_748_v30)
				(globalState).F4 = (_732_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))).Int()).(bool)
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
				_ = _index124
				(_717_v0).ArraySet1((_this).F10(), (_index124).Int())
			}
			var _749_v31 _dafny.Map
			_ = _749_v31
			_749_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_733_v15).Union(_733_v15), _dafny.UnicodeSeqOfUtf8Bytes("fvhdv"))
			_749_v31 = (_749_v31).Update(_dafny.MultiSetFromSeq(_730_v12), _dafny.Companion_Sequence_.Concatenate((_this).F9(), (_this).F9()))
		}
		var _750_v32 *C4
		_ = _750_v32
		var _nw123 *C4 = New_C4_()
		_ = _nw123
		_nw123.Ctor__((_732_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))).Int()).(bool), (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), (_this).F10())
		_750_v32 = _nw123
		var _751_v33 _dafny.Sequence
		_ = _751_v33
		_751_v33 = _dafny.SeqOf(_750_v32, _750_v32)
		_751_v33 = _751_v33
		var _hi7 _dafny.Int = (_this).F7()
		_ = _hi7
		for _752_i1 := p0; _752_i1.Cmp(_hi7) < 0; _752_i1 = _752_i1.Plus(_dafny.One) {
			var _753_v34 _dafny.Sequence
			_ = _753_v34
			_753_v34 = _dafny.UnicodeSeqOfUtf8Bytes("tjdnlna")
			var _754_v35 _dafny.Array
			_ = _754_v35
			var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
			_ = _nw124
			_754_v35 = _nw124
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_754_v35), 0))
			_ = _index125
			(_754_v35).ArraySet1(_732_v14, (_index125).Int())
			var _755_v36 _dafny.Array
			_ = _755_v36
			var _nwElement0_22 _dafny.CodePoint = _dafny.CodePoint('f')
			_ = _nwElement0_22
			var _nw125 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(2))
			_ = _nw125
			(_nw125).ArraySet1CodePoint(_nwElement0_22, 0)
			(_nw125).ArraySet1CodePoint(_722_v4, 1)
			_755_v36 = _nw125
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_755_v36), 0))
			_ = _index126
			(_755_v36).ArraySet1CodePoint(_722_v4, (_index126).Int())
			var _756_v37 _dafny.Map
			_ = _756_v37
			_756_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p0, (_733_v15).Cardinality()), (_732_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))).Int()).(bool))
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
			_ = _index127
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_754_v35), 0))
			_ = _index128
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_755_v36), 0))
			_ = _index129
			var _rhs125 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("olimw")
			_ = _rhs125
			var _rhs126 _dafny.Int = ((_756_v37).Merge(_756_v37)).Cardinality()
			_ = _rhs126
			var _rhs127 bool = ((_this).F8()).Cmp(Companion_Default___.SafeModuloInt((_this).F8(), (_this).F8())) < 0
			_ = _rhs127
			var _rhs128 _dafny.Array = _732_v14
			_ = _rhs128
			var _rhs129 _dafny.CodePoint = _722_v4
			_ = _rhs129
			var _lhs93 _dafny.Array = _717_v0
			_ = _lhs93
			var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
			_ = _lhs94
			var _lhs95 *GlobalState = globalState
			_ = _lhs95
			var _lhs96 _dafny.Array = _754_v35
			_ = _lhs96
			var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_754_v35), 0))
			_ = _lhs97
			var _lhs98 _dafny.Array = _755_v36
			_ = _lhs98
			var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_755_v36), 0))
			_ = _lhs99
			_753_v34 = _rhs125
			(_lhs93).ArraySet1(_rhs126, (_lhs94).Int())
			_lhs95.F3 = _rhs127
			(_lhs96).ArraySet1(_rhs128, (_lhs97).Int())
			(_lhs98).ArraySet1CodePoint(_rhs129, (_lhs99).Int())
			var _757_v38 D3
			_ = _757_v38
			_757_v38 = Companion_D3_.Create_DC7_((_750_v32).F11())
			var _source17 D3 = _757_v38
			_ = _source17
			if _source17.Is_DC7() {
				var _758___mcc_h0 bool = _source17.Get_().(D3_DC7).Cf10
				_ = _758___mcc_h0
				var _759_cf10 bool = _758___mcc_h0
				_ = _759_cf10
				var _760_v39 _dafny.Set
				_ = _760_v39
				_760_v39 = _dafny.SetOf(_723_v5, _723_v5, _723_v5, _723_v5, _723_v5)
				var _761_v41 _dafny.Array
				_ = _761_v41
				var _nwElement0_23 _dafny.Set = Companion_Default___.Fm31((_this).F8(), _dafny.MultiSetFromSeq(Companion_Default___.Fm30(_759_cf10, globalState)), globalState)
				_ = _nwElement0_23
				var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(15))
				_ = _nw126
				(_nw126).ArraySet1(_nwElement0_23, 0)
				(_nw126).ArraySet1(_760_v39, 1)
				(_nw126).ArraySet1(func() _dafny.Set {
					var _coll36 = _dafny.NewBuilder()
					_ = _coll36
					for _iter40 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(386))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg50 _dafny.Int) interface{} {
							return coer50(arg50)
						}
					}((func(_762_v5 _dafny.Map, _763_i1 _dafny.Int, _764_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
						return func(_765_i2 _dafny.Int) _dafny.Map {
							return (_762_v5).Update(_dafny.IntOfUint32((_dafny.SeqOf(_763_i1)).Cardinality()), _764_p0)
						}
					})(_723_v5, _752_i1, p0)))).Elements()); ; {
						_compr_36, _ok40 := _iter40()
						if !_ok40 {
							break
						}
						var _766_v40 _dafny.Map
						_766_v40 = interface{}(_compr_36).(_dafny.Map)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(386))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg51 _dafny.Int) interface{} {
								return coer51(arg51)
							}
						}((func(_767_v5 _dafny.Map, _768_i1 _dafny.Int, _769_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
							return func(_765_i2 _dafny.Int) _dafny.Map {
								return (_767_v5).Update(_dafny.IntOfUint32((_dafny.SeqOf(_768_i1)).Cardinality()), _769_p0)
							}
						})(_723_v5, _752_i1, p0))), _766_v40) {
							_coll36.Add(_766_v40)
						}
					}
					return _coll36.ToSet()
				}(), 2)
				(_nw126).ArraySet1(_760_v39, 3)
				(_nw126).ArraySet1(_760_v39, 4)
				(_nw126).ArraySet1(_760_v39, 5)
				(_nw126).ArraySet1(_760_v39, 6)
				(_nw126).ArraySet1((_760_v39).Intersection(_760_v39), 7)
				(_nw126).ArraySet1(_760_v39, 8)
				(_nw126).ArraySet1((Companion_Default___.Fm31((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), _729_v11, globalState)).Difference(_760_v39), 9)
				(_nw126).ArraySet1(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_dafny.Zero).Minus((_this).F10()))), 10)
				(_nw126).ArraySet1(_760_v39, 11)
				(_nw126).ArraySet1(_760_v39, 12)
				(_nw126).ArraySet1(_760_v39, 13)
				(_nw126).ArraySet1(Companion_Default___.Fm31((_this).F7(), _dafny.MultiSetFromSeq(_720_v2), globalState), 14)
				_761_v41 = _nw126
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_761_v41), 0))
				_ = _index130
				(_761_v41).ArraySet1(_760_v39, (_index130).Int())
				var _770_v42 _dafny.Sequence
				_ = _770_v42
				_770_v42 = _dafny.SeqOf((_760_v39).Intersection(_760_v39))
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_761_v41), 0))
				_ = _index131
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
				_ = _index132
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_755_v36), 0))
				_ = _index133
				var _rhs130 _dafny.Sequence = (_this).F9()
				_ = _rhs130
				var _rhs131 bool = ((_this).F8()).Cmp((_this).F10()) < 0
				_ = _rhs131
				var _rhs132 _dafny.Set = (_770_v42).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F10()), _dafny.IntOfUint32((_770_v42).Cardinality()))).Uint32()).(_dafny.Set)
				_ = _rhs132
				var _rhs133 _dafny.Int = ((_this).F7()).Plus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cniy")).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).F9(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F8()), _dafny.IntOfUint32(((_this).F9()).Cardinality()))).Uint32(), (_755_v36).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_755_v36), 0))).Int()))).Cardinality())))
				_ = _rhs133
				var _rhs134 _dafny.CodePoint = _722_v4
				_ = _rhs134
				var _lhs100 _dafny.Array = _761_v41
				_ = _lhs100
				var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_761_v41), 0))
				_ = _lhs101
				var _lhs102 _dafny.Array = _717_v0
				_ = _lhs102
				var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
				_ = _lhs103
				var _lhs104 _dafny.Array = _755_v36
				_ = _lhs104
				var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_755_v36), 0))
				_ = _lhs105
				_753_v34 = _rhs130
				_719_v1 = _rhs131
				(_lhs100).ArraySet1(_rhs132, (_lhs101).Int())
				(_lhs102).ArraySet1(_rhs133, (_lhs103).Int())
				(_lhs104).ArraySet1CodePoint(_rhs134, (_lhs105).Int())
				_717_v0 = _717_v0
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))
				_ = _index134
				var _rhs135 bool = true
				_ = _rhs135
				var _rhs136 _dafny.Array = _717_v0
				_ = _rhs136
				var _rhs137 _dafny.Sequence = _dafny.SeqOf((_730_v12).Select((Companion_Default___.SafeIndex((_733_v15).Cardinality(), _dafny.IntOfUint32((_730_v12).Cardinality()))).Uint32()).(_dafny.Int))
				_ = _rhs137
				var _rhs138 bool = ((_750_v32).F11()) == ((_732_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))).Int()).(bool))
				_ = _rhs138
				var _rhs139 _dafny.MultiSet = _727_v9
				_ = _rhs139
				var _lhs106 _dafny.Array = _732_v14
				_ = _lhs106
				var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))
				_ = _lhs107
				(_lhs106).ArraySet1(_rhs135, (_lhs107).Int())
				_717_v0 = _rhs136
				_730_v12 = _rhs137
				_759_cf10 = _rhs138
				_727_v9 = _rhs139
				_721_v3 = (_721_v3).Update((_this).F9(), (_732_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_732_v14), 0))).Int()).(bool))
			} else {
				var _771___mcc_h1 _dafny.Set = _source17.Get_().(D3_DC6).Cf9
				_ = _771___mcc_h1
				var _772_cf9 _dafny.Set = _771___mcc_h1
				_ = _772_cf9
				(globalState).F4 = _719_v1
				var _773_v43 _dafny.Map
				_ = _773_v43
				_773_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_725_v7, _dafny.Companion_Sequence_.Concatenate((_this).F9(), _753_v34))
				_753_v34 = (func() _dafny.Sequence {
					if (_773_v43).Contains((_772_cf9).Union(_772_cf9)) {
						return (_773_v43).Get((_772_cf9).Union(_772_cf9)).(_dafny.Sequence)
					}
					return _753_v34
				})()
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))
				_ = _index135
				(_717_v0).ArraySet1(_752_i1, (_index135).Int())
				(globalState).F4 = (_750_v32).F11()
			}
			var _774_v44 D1
			_ = _774_v44
			_774_v44 = Companion_D1_.Create_DC1_(_dafny.CodePoint('u'))
			var _775_v45 D1
			_ = _775_v45
			_775_v45 = Companion_D1_.Create_DC3_(_774_v44)
			var _776_v46 bool
			_ = _776_v46
			var _777_v47 bool
			_ = _777_v47
			var _778_v48 bool
			_ = _778_v48
			var _779_v49 bool
			_ = _779_v49
			var _out6 bool
			_ = _out6
			var _out7 bool
			_ = _out7
			var _out8 bool
			_ = _out8
			var _out9 bool
			_ = _out9
			_out6, _out7, _out8, _out9 = (_750_v32).M5((_750_v32).F11(), _dafny.IntOfUint32((_730_v12).Cardinality()), _dafny.IntOfInt64(677), (_dafny.SetOf(_775_v45)).Cardinality(), globalState)
			_776_v46 = _out6
			_777_v47 = _out7
			_778_v48 = _out8
			_779_v49 = _out9
			r0 = Companion_Default___.SafeDivisionInt(((_726_v8).Merge(_726_v8)).Cardinality(), (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(993), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int))
		}
		r0 = (_this).F8()
		return r0
	}
}
func (_this *C5) F9() _dafny.Sequence {
	{
		return _this._f9
	}
}
func (_this *C5) F10() _dafny.Int {
	{
		return _this._f10
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f7 _dafny.Int
	_f8 _dafny.Int
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C6{}
var _ T1 = &C6{}
var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F7() _dafny.Int {
	return _this._f7
}
func (_this *C6) F8() _dafny.Int {
	return _this._f8
}
func (_this *C6) Ctor__(f7 _dafny.Int, f8 _dafny.Int) {
	{
		(_this)._f7 = f7
		(_this)._f8 = f8
	}
}
func (_this *C6) Fm2(p0 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lysobld"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(732))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg52 _dafny.Int) interface{} {
				return coer52(arg52)
			}
		}(func(_780_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))), _dafny.UnicodeSeqOfUtf8Bytes("exbusrxu"))
	}
}
func (_this *C6) Fm0(p0 bool, globalState *GlobalState) _dafny.Set {
	{
		return (_dafny.SetOf(_dafny.SetOf(false, !(!(true)), !(false)))).Intersection((_dafny.SetOf(_dafny.SetOf(true, false))).Difference(_dafny.SetOf(_dafny.SetOf(false), _dafny.SetOf(false, true, false, true, true))))
	}
}
func (_this *C6) Fm1(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_dafny.IntOfInt64(970)).Cmp((_this).F7()) <= 0
	}
}
func (_this *C6) Fm3(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		if !(false) {
			return _dafny.IntOfInt64(-377)
		} else {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sv")).Cardinality())
		}
	}
}
func (_this *C6) Fm4(p0 bool, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F7()
	}
}
func (_this *C6) M1(p0 _dafny.Sequence, globalState *GlobalState) {
	{
		var _781_v0 _dafny.Int
		_ = _781_v0
		_781_v0 = _dafny.IntOfInt64(-479)
		var _782_v1 _dafny.Sequence
		_ = _782_v1
		_782_v1 = _dafny.UnicodeSeqOfUtf8Bytes("hndygtt")
		var _783_v2 bool
		_ = _783_v2
		_783_v2 = true
		var _784_v3 _dafny.MultiSet
		_ = _784_v3
		_784_v3 = _dafny.MultiSetOf(_783_v2)
		var _rhs140 bool = _783_v2
		_ = _rhs140
		var _rhs141 _dafny.Int = (func() _dafny.Int {
			if _783_v2 {
				return ((_this).F8()).Times((func() _dafny.Int {
					if (_784_v3).Contains(_783_v2) {
						return (_784_v3).Multiplicity(_783_v2)
					}
					return (_this).F8()
				})())
			}
			return ((_this).F7()).Plus((_this).F8())
		})()
		_ = _rhs141
		var _rhs142 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_782_v1, (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_this).F7(), _781_v0), _dafny.IntOfUint32((_782_v1).Cardinality()))).Uint32(), Companion_Default___.Fm5(globalState))
		_ = _rhs142
		var _lhs108 *GlobalState = globalState
		_ = _lhs108
		_lhs108.F3 = _rhs140
		_781_v0 = _rhs141
		_782_v1 = _rhs142
		var _785_v4 _dafny.Array
		_ = _785_v4
		var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw127
		_785_v4 = _nw127
		var _786_v5 _dafny.MultiSet
		_ = _786_v5
		_786_v5 = _dafny.MultiSetOf(_dafny.IntOfUint32((_782_v1).Cardinality()))
		var _787_v6 _dafny.Map
		_ = _787_v6
		_787_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_783_v2, _786_v5)
		var _788_v7 _dafny.Sequence
		_ = _788_v7
		_788_v7 = _dafny.SeqOf((_dafny.Zero).Minus(((func() _dafny.MultiSet {
			if (_787_v6).Contains(_783_v2) {
				return (_787_v6).Get(_783_v2).(_dafny.MultiSet)
			}
			return _786_v5
		})()).Cardinality()), (_this).F7())
		var _789_v8 T1
		_ = _789_v8
		var _nw128 *C0 = New_C0_()
		_ = _nw128
		_nw128.Ctor__(_785_v4, (_this).F8(), Companion_Default___.SafeModuloInt((_this).F8(), (_788_v7).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_788_v7).Cardinality()))).Uint32()).(_dafny.Int)))
		_789_v8 = _nw128
		_789_v8 = _789_v8
		var _790_v9 _dafny.Sequence
		_ = _790_v9
		_790_v9 = _dafny.SeqOf((_786_v5).Update(_781_v0, Companion_Default___.Abs(_781_v0)))
		var _791_i0 _dafny.Int
		_ = _791_i0
		_791_i0 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.IsProperPrefixOf(_790_v9, _dafny.SeqOf(_786_v5, _786_v5)) {
				{
					if (_791_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_791_i0 = (_791_i0).Plus(_dafny.One)
					if (func() bool {
						if _783_v2 {
							return _dafny.Companion_Sequence_.Equal(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(938))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg53 _dafny.Int) interface{} {
									return coer53(arg53)
								}
							}(func(_792_i1 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('e')
							})))
						}
						return (func() bool {
							if _783_v2 {
								return _783_v2
							}
							return _783_v2
						})()
					})() {
						var _793_v10 _dafny.CodePoint
						_ = _793_v10
						_793_v10 = _dafny.CodePoint('r')
						_782_v1 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_782_v1, _782_v1), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_789_v8).F7(), (_this).F7()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_782_v1, _782_v1)).Cardinality()))).Uint32(), _793_v10)
						var _794_v11 _dafny.Array
						_ = _794_v11
						var _len0_19 _dafny.Int = _dafny.IntOfInt64(26)
						_ = _len0_19
						var _nw129 _dafny.Array
						_ = _nw129
						if _len0_19.Cmp(_dafny.Zero) == 0 {
							_nw129 = _dafny.NewArray(_len0_19)
						} else {
							var _init19 func(_dafny.Int) bool = (func(_795_p0 _dafny.Sequence, _796_v1 _dafny.Sequence) func(_dafny.Int) bool {
								return func(_797_i2 _dafny.Int) bool {
									return _dafny.Companion_Sequence_.Equal(_795_p0, _796_v1)
								}
							})(p0, _782_v1)
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
						_794_v11 = _nw129
						var _798_v12 D5
						_ = _798_v12
						_798_v12 = Companion_D5_.Create_DC13_(true, _793_v10)
						var _799_v13 D6
						_ = _799_v13
						_799_v13 = Companion_D6_.Create_DC15_(_794_v11)
						var _rhs143 _dafny.Array = (_799_v13).Dtor_cf23()
						_ = _rhs143
						var _rhs144 D5 = Companion_D5_.Create_DC13_(_783_v2, _793_v10)
						_ = _rhs144
						_794_v11 = _rhs143
						_798_v12 = _rhs144
						var _800_v14 _dafny.Array
						_ = _800_v14
						var _nw130 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(9))
						_ = _nw130
						_800_v14 = _nw130
						var _801_v15 _dafny.Map
						_ = _801_v15
						_801_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_783_v2, _800_v14)
						_800_v14 = (func() _dafny.Array {
							if (_801_v15).Contains(_783_v2) {
								return (_801_v15).Get(_783_v2).(_dafny.Array)
							}
							return _800_v14
						})()
						var _802_v16 _dafny.Set
						_ = _802_v16
						_802_v16 = _dafny.SetOf(_793_v10, _793_v10)
						_802_v16 = (_802_v16).Union(_802_v16)
						var _803_v17 _dafny.Set
						_ = _803_v17
						_803_v17 = _dafny.SetOf(Companion_Default___.Fm8(_783_v2, _788_v7, _783_v2, globalState), _dafny.Companion_Sequence_.Concatenate(p0, p0))
						var _804_v18 _dafny.Map
						_ = _804_v18
						_804_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-863))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg54 _dafny.Int) interface{} {
								return coer54(arg54)
							}
						}((func(_805_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_806_i3 _dafny.Int) _dafny.CodePoint {
								return _805_v10
							}
						})(_793_v10))), _dafny.SeqOf(_783_v2))
						_803_v17 = ((_803_v17).Difference(_dafny.SetOf(_782_v1))).Union(func() _dafny.Set {
							var _coll37 = _dafny.NewBuilder()
							_ = _coll37
							for _iter41 := _dafny.Iterate((_804_v18).Keys().Elements()); ; {
								_compr_37, _ok41 := _iter41()
								if !_ok41 {
									break
								}
								var _807_v19 _dafny.Sequence
								_807_v19 = interface{}(_compr_37).(_dafny.Sequence)
								if (_804_v18).Contains(_807_v19) {
									_coll37.Add(_807_v19)
								}
							}
							return _coll37.ToSet()
						}())
					} else {
						var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_785_v4), 0))
						_ = _index136
						(_785_v4).ArraySet1((func() _dafny.Int {
							if (_784_v3).Contains(!(_783_v2)) {
								return (_784_v3).Multiplicity(!(_783_v2))
							}
							return (_this).F7()
						})(), (_index136).Int())
						var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_785_v4), 0))
						_ = _index137
						(_785_v4).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm20(_783_v2, _783_v2, _783_v2, globalState), (_789_v8).F8()), (_index137).Int())
						var _808_v20 D7
						_ = _808_v20
						_808_v20 = Companion_D7_.Create_DC17_(p0)
						var _809_v21 _dafny.Sequence
						_ = _809_v21
						_809_v21 = _dafny.SeqOf(_782_v1, Companion_Default___.Fm8(_783_v2, _788_v7, _783_v2, globalState))
						var _810_v22 _dafny.Set
						_ = _810_v22
						_810_v22 = _dafny.SetOf(_782_v1, _dafny.UnicodeSeqOfUtf8Bytes("j"))
						var _811_v23 _dafny.Array
						_ = _811_v23
						var _nwElement0_24 _dafny.Sequence = _782_v1
						_ = _nwElement0_24
						var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(21))
						_ = _nw131
						(_nw131).ArraySet1(_nwElement0_24, 0)
						(_nw131).ArraySet1(_782_v1, 1)
						(_nw131).ArraySet1(_782_v1, 2)
						(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_782_v1, _782_v1), 3)
						(_nw131).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(362))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg55 _dafny.Int) interface{} {
								return coer55(arg55)
							}
						}(func(_812_i4 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('b')
						})), 4)
						(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_782_v1, p0), 5)
						(_nw131).ArraySet1((_808_v20).Dtor_cf26(), 6)
						(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(10))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg56 _dafny.Int) interface{} {
								return coer56(arg56)
							}
						}(func(_813_i5 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('l')
						}))), 7)
						(_nw131).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(805))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg57 _dafny.Int) interface{} {
								return coer57(arg57)
							}
						}(func(_814_i6 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('u')
						})), 8)
						(_nw131).ArraySet1(_782_v1, 9)
						(_nw131).ArraySet1(p0, 10)
						(_nw131).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(519))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg58 _dafny.Int) interface{} {
								return coer58(arg58)
							}
						}(func(_815_i7 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('o')
						})), 11)
						(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_808_v20).Dtor_cf26(), _dafny.UnicodeSeqOfUtf8Bytes("klcneosa")), 12)
						(_nw131).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vbsipip"), 13)
						(_nw131).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("wudxpkey"), 14)
						(_nw131).ArraySet1(p0, 15)
						(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_782_v1, (_809_v21).Select((Companion_Default___.SafeIndex((_810_v22).Cardinality(), _dafny.IntOfUint32((_809_v21).Cardinality()))).Uint32()).(_dafny.Sequence)), 16)
						(_nw131).ArraySet1(_782_v1, 17)
						(_nw131).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_782_v1, _782_v1), 18)
						(_nw131).ArraySet1(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_786_v5).Cardinality(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), Companion_Default___.Fm5(globalState)), 19)
						(_nw131).ArraySet1(_782_v1, 20)
						_811_v23 = _nw131
						var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_811_v23), 0))
						_ = _index138
						(_811_v23).ArraySet1(_782_v1, (_index138).Int())
						var _816_v24 _dafny.Sequence
						_ = _816_v24
						_816_v24 = _dafny.SeqOf(_783_v2)
						var _817_v25 _dafny.Map
						_ = _817_v25
						_817_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_781_v0, (Companion_D5_.Create_DC12_(_783_v2)).Dtor_cf19())
						var _818_v26 _dafny.Map
						_ = _818_v26
						_818_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_816_v24).Cardinality()), (_817_v25).Cardinality())
						var _819_v27 _dafny.CodePoint
						_ = _819_v27
						_819_v27 = _dafny.CodePoint('j')
						var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_811_v23), 0))
						_ = _index139
						(_811_v23).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_782_v1, (Companion_Default___.SafeIndex((_818_v26).Cardinality(), _dafny.IntOfUint32((_782_v1).Cardinality()))).Uint32(), _819_v27), _dafny.UnicodeSeqOfUtf8Bytes("vhe")), (_index139).Int())
						var _820_v28 _dafny.Array
						_ = _820_v28
						var _nw132 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
						_ = _nw132
						_820_v28 = _nw132
						var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_820_v28), 0))
						_ = _index140
						(_820_v28).ArraySet1(_783_v2, (_index140).Int())
						var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_820_v28), 0))
						_ = _index141
						(_820_v28).ArraySet1((func() bool {
							if !(!(_783_v2) || (_783_v2)) {
								return _783_v2
							}
							return _783_v2
						})(), (_index141).Int())
						var _821_v29 _dafny.Array
						_ = _821_v29
						var _len0_20 _dafny.Int = _dafny.IntOfInt64(9)
						_ = _len0_20
						var _nw133 _dafny.Array
						_ = _nw133
						if _len0_20.Cmp(_dafny.Zero) == 0 {
							_nw133 = _dafny.NewArray(_len0_20)
						} else {
							var _init20 func(_dafny.Int) _dafny.Sequence = (func(_822_v24 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_823_i8 _dafny.Int) _dafny.Sequence {
									return _dafny.SeqOf(_822_v24)
								}
							})(_816_v24)
							_ = _init20
							var _element0_20 = _init20(_dafny.Zero)
							_ = _element0_20
							_nw133 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
							(_nw133).ArraySet1(_element0_20, 0)
							var _nativeLen0_20 = (_len0_20).Int()
							_ = _nativeLen0_20
							for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
								(_nw133).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
							}
						}
						_821_v29 = _nw133
						var _824_v30 _dafny.Sequence
						_ = _824_v30
						_824_v30 = _dafny.SeqOf(_816_v24, _816_v24)
						var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_821_v29), 0))
						_ = _index142
						(_821_v29).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_824_v30, _824_v30), (_index142).Int())
						var _825_v31 D5
						_ = _825_v31
						_825_v31 = Companion_D5_.Create_DC13_((_820_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_820_v28), 0))).Int()).(bool), _819_v27)
						var _826_v32 _dafny.Map
						_ = _826_v32
						_826_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_825_v31, _819_v27)
						var _827_v33 D7
						_ = _827_v33
						_827_v33 = Companion_D7_.Create_DC20_(_819_v27, ((func() _dafny.Map {
							if true {
								return _826_v32
							}
							return _826_v32
						})()).Cardinality())
						var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_785_v4), 0))
						_ = _index143
						var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_821_v29), 0))
						_ = _index144
						var _rhs145 _dafny.Int = (_789_v8).F7()
						_ = _rhs145
						var _rhs146 _dafny.Sequence = _824_v30
						_ = _rhs146
						var _rhs147 bool = _783_v2
						_ = _rhs147
						var _rhs148 D7 = Companion_D7_.Create_DC20_(_819_v27, (_dafny.Zero).Minus((_dafny.Zero).Minus((_789_v8).F8())))
						_ = _rhs148
						var _rhs149 bool = !((_820_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_820_v28), 0))).Int()).(bool))
						_ = _rhs149
						var _lhs109 _dafny.Array = _785_v4
						_ = _lhs109
						var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_785_v4), 0))
						_ = _lhs110
						var _lhs111 _dafny.Array = _821_v29
						_ = _lhs111
						var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_821_v29), 0))
						_ = _lhs112
						var _lhs113 *GlobalState = globalState
						_ = _lhs113
						var _lhs114 *GlobalState = globalState
						_ = _lhs114
						(_lhs109).ArraySet1(_rhs145, (_lhs110).Int())
						(_lhs111).ArraySet1(_rhs146, (_lhs112).Int())
						_lhs113.F3 = _rhs147
						_827_v33 = _rhs148
						_lhs114.F4 = _rhs149
						_781_v0 = (_789_v8).F7()
					}
					var _828_v34 _dafny.CodePoint
					_ = _828_v34
					_828_v34 = _dafny.CodePoint('k')
					var _829_v35 D5
					_ = _829_v35
					_829_v35 = Companion_D5_.Create_DC13_(_783_v2, _828_v34)
					var _source18 D5 = _829_v35
					_ = _source18
					if _source18.Is_DC12() {
						var _830___mcc_h0 bool = _source18.Get_().(D5_DC12).Cf19
						_ = _830___mcc_h0
						var _831_cf19 bool = _830___mcc_h0
						_ = _831_cf19
						var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_785_v4), 0))
						_ = _index145
						(_785_v4).ArraySet1((_dafny.Zero).Minus((_789_v8).F7()), (_index145).Int())
						var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_785_v4), 0))
						_ = _index146
						(_785_v4).ArraySet1((_this).F8(), (_index146).Int())
						_781_v0 = (_789_v8).F7()
						var _832_v36 _dafny.Map
						_ = _832_v36
						_832_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_831_cf19, (_785_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_785_v4), 0))).Int()).(_dafny.Int))
						var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_785_v4), 0))
						_ = _index147
						(_785_v4).ArraySet1(Companion_Default___.SafeDivisionInt((_832_v36).Cardinality(), (_785_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_785_v4), 0))).Int()).(_dafny.Int)), (_index147).Int())
						var _833_v37 _dafny.Array
						_ = _833_v37
						var _len0_21 _dafny.Int = _dafny.IntOfInt64(19)
						_ = _len0_21
						var _nw134 _dafny.Array
						_ = _nw134
						if _len0_21.Cmp(_dafny.Zero) == 0 {
							_nw134 = _dafny.NewArray(_len0_21)
						} else {
							var _init21 func(_dafny.Int) bool = (func(_834_v2 bool) func(_dafny.Int) bool {
								return func(_835_i9 _dafny.Int) bool {
									return _834_v2
								}
							})(_783_v2)
							_ = _init21
							var _element0_21 = _init21(_dafny.Zero)
							_ = _element0_21
							_nw134 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
							(_nw134).ArraySet1(_element0_21, 0)
							var _nativeLen0_21 = (_len0_21).Int()
							_ = _nativeLen0_21
							for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
								(_nw134).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
							}
						}
						_833_v37 = _nw134
						var _836_v38 _dafny.Sequence
						_ = _836_v38
						_836_v38 = _dafny.SeqOf(_833_v37)
						var _837_v39 D4
						_ = _837_v39
						_837_v39 = Companion_D4_.Create_DC9_((_this).F8(), _836_v38, (_dafny.Zero).Minus((_this).F8()), (_dafny.Zero).Minus((_789_v8).F8()))
						var _pat_let_tv12 = _781_v0
						_ = _pat_let_tv12
						var _pat_let_tv13 = _781_v0
						_ = _pat_let_tv13
						var _838_v40 _dafny.Map
						_ = _838_v40
						_838_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let9_0 D4) D4 {
							return func(_839_dt__update__tmp_h0 D4) D4 {
								return func(_pat_let10_0 _dafny.Int) D4 {
									return func(_840_dt__update_hcf15_h0 _dafny.Int) D4 {
										return func(_pat_let11_0 _dafny.Int) D4 {
											return func(_841_dt__update_hcf12_h0 _dafny.Int) D4 {
												return Companion_D4_.Create_DC9_(_841_dt__update_hcf12_h0, (_839_dt__update__tmp_h0).Dtor_cf13(), (_839_dt__update__tmp_h0).Dtor_cf14(), _840_dt__update_hcf15_h0)
											}(_pat_let11_0)
										}(_pat_let_tv13)
									}(_pat_let10_0)
								}(_pat_let_tv12)
							}(_pat_let9_0)
						}(_837_v39), (_789_v8).F7())
						var _842_v41 _dafny.Map
						_ = _842_v41
						_842_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _831_cf19)
						_838_v40 = (_838_v40).Update(_837_v39, (_842_v41).Cardinality())
					} else if _source18.Is_DC13() {
						var _843___mcc_h1 bool = _source18.Get_().(D5_DC13).Cf20
						_ = _843___mcc_h1
						var _844___mcc_h2 _dafny.CodePoint = _source18.Get_().(D5_DC13).Cf21
						_ = _844___mcc_h2
						var _845_cf21 _dafny.CodePoint = _844___mcc_h2
						_ = _845_cf21
						var _846_cf20 bool = _843___mcc_h1
						_ = _846_cf20
						var _847_v42 _dafny.Set
						_ = _847_v42
						_847_v42 = _dafny.SetOf(!(_846_cf20))
						var _848_v43 _dafny.Set
						_ = _848_v43
						_848_v43 = _dafny.SetOf(_dafny.SetOf(!(_783_v2)), _847_v42, (_847_v42).Difference(_847_v42))
						_848_v43 = (_dafny.SetOf(_847_v42, _847_v42, _847_v42, _dafny.SetOf(_846_cf20))).Union(_848_v43)
						var _849_v44 _dafny.Array
						_ = _849_v44
						var _nwElement0_25 bool = _846_cf20
						_ = _nwElement0_25
						var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(15))
						_ = _nw135
						(_nw135).ArraySet1(_nwElement0_25, 0)
						(_nw135).ArraySet1(_846_cf20, 1)
						(_nw135).ArraySet1(_783_v2, 2)
						(_nw135).ArraySet1(_783_v2, 3)
						(_nw135).ArraySet1(_846_cf20, 4)
						(_nw135).ArraySet1(_846_cf20, 5)
						(_nw135).ArraySet1(true, 6)
						(_nw135).ArraySet1(true, 7)
						(_nw135).ArraySet1(_846_cf20, 8)
						(_nw135).ArraySet1(_783_v2, 9)
						(_nw135).ArraySet1(_846_cf20, 10)
						(_nw135).ArraySet1(true, 11)
						(_nw135).ArraySet1(_846_cf20, 12)
						(_nw135).ArraySet1(true, 13)
						(_nw135).ArraySet1(_846_cf20, 14)
						_849_v44 = _nw135
						var _850_v45 _dafny.MultiSet
						_ = _850_v45
						_850_v45 = _dafny.MultiSetOf(_849_v44)
						(globalState).F3 = (_850_v45).IsSubsetOf((func() _dafny.MultiSet {
							if _846_cf20 {
								return _850_v45
							}
							return _850_v45
						})())
						var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_785_v4), 0))
						_ = _index148
						(_785_v4).ArraySet1((_this).F8(), (_index148).Int())
						var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_785_v4), 0))
						_ = _index149
						(_785_v4).ArraySet1((_dafny.Zero).Minus(((_dafny.Zero).Minus((_784_v3).Cardinality())).Times((_789_v8).F8())), (_index149).Int())
						var _851_v46 _dafny.Array
						_ = _851_v46
						var _nwElement0_26 _dafny.Array = _849_v44
						_ = _nwElement0_26
						var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(17))
						_ = _nw136
						(_nw136).ArraySet1(_nwElement0_26, 0)
						(_nw136).ArraySet1(_849_v44, 1)
						(_nw136).ArraySet1(_849_v44, 2)
						(_nw136).ArraySet1(_849_v44, 3)
						(_nw136).ArraySet1(_849_v44, 4)
						(_nw136).ArraySet1(_849_v44, 5)
						(_nw136).ArraySet1(_849_v44, 6)
						(_nw136).ArraySet1(_849_v44, 7)
						(_nw136).ArraySet1(_849_v44, 8)
						(_nw136).ArraySet1(_849_v44, 9)
						(_nw136).ArraySet1(_849_v44, 10)
						(_nw136).ArraySet1(_849_v44, 11)
						(_nw136).ArraySet1(_849_v44, 12)
						(_nw136).ArraySet1(_849_v44, 13)
						(_nw136).ArraySet1(_849_v44, 14)
						(_nw136).ArraySet1(_849_v44, 15)
						(_nw136).ArraySet1(_849_v44, 16)
						_851_v46 = _nw136
						var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_851_v46), 0))
						_ = _index150
						(_851_v46).ArraySet1(_849_v44, (_index150).Int())
						var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_851_v46), 0))
						_ = _index151
						(_851_v46).ArraySet1((func() _dafny.Array {
							if (_847_v42).IsProperSubsetOf(_dafny.SetOf(_846_cf20, _783_v2, _783_v2, _783_v2, _846_cf20)) {
								return _849_v44
							}
							return _849_v44
						})(), (_index151).Int())
					} else if _source18.Is_DC11() {
						var _852___mcc_h3 _dafny.Map = _source18.Get_().(D5_DC11).Cf18
						_ = _852___mcc_h3
						var _853_cf18 _dafny.Map = _852___mcc_h3
						_ = _853_cf18
						var _854_v47 *C5
						_ = _854_v47
						var _nw137 *C5 = New_C5_()
						_ = _nw137
						_nw137.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("gk"), (_789_v8).F8(), (_dafny.Zero).Minus((_this).Fm3(_783_v2, _783_v2, _783_v2, _783_v2, globalState)), (_789_v8).F7())
						_854_v47 = _nw137
						_854_v47 = _854_v47
						var _855_v48 _dafny.Sequence
						_ = _855_v48
						_855_v48 = _dafny.SeqOf(_788_v7)
						_855_v48 = _855_v48
						var _856_v49 T2
						_ = _856_v49
						var _nw138 *C3 = New_C3_()
						_ = _nw138
						_nw138.Ctor__(_783_v2, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((p0).Cardinality()), (_789_v8).F7()), (_dafny.Zero).Minus((_789_v8).F8()))
						_856_v49 = _nw138
						var _nw139 *C2 = New_C2_()
						_ = _nw139
						_nw139.Ctor__(_dafny.IntOfInt64(922), (_dafny.Zero).Minus((_789_v8).F8()))
						_856_v49 = _nw139
						var _857_v50 _dafny.Int
						_ = _857_v50
						var _858_v51 _dafny.Int
						_ = _858_v51
						var _859_v52 bool
						_ = _859_v52
						var _out10 _dafny.Int
						_ = _out10
						var _out11 _dafny.Int
						_ = _out11
						var _out12 bool
						_ = _out12
						_out10, _out11, _out12 = (_this).M3(((_this).F7()).Cmp((_856_v49).F8()) > 0, _781_v0, (func() _dafny.Array {
							if _783_v2 {
								return _785_v4
							}
							return _785_v4
						})(), (_this).Fm4(_783_v2, _783_v2, globalState), globalState)
						_857_v50 = _out10
						_858_v51 = _out11
						_859_v52 = _out12
					} else {
						var _860___mcc_h4 D5 = _source18.Get_().(D5_DC14).Cf22
						_ = _860___mcc_h4
						var _861_cf22 D5 = _860___mcc_h4
						_ = _861_cf22
						var _862_v53 D1
						_ = _862_v53
						_862_v53 = Companion_D1_.Create_DC2_()
						var _863_v54 _dafny.Sequence
						_ = _863_v54
						_863_v54 = _dafny.SeqOf(_862_v53)
						_863_v54 = _863_v54
						var _864_v55 _dafny.Map
						_ = _864_v55
						_864_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_828_v34, (_789_v8).F8())
						_781_v0 = (_864_v55).Cardinality()
						_781_v0 = _781_v0
						_781_v0 = _dafny.IntOfUint32((p0).Cardinality())
					}
					var _865_v56 _dafny.Map
					_ = _865_v56
					_865_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_781_v0, _783_v2)
					var _866_v57 _dafny.Map
					_ = _866_v57
					_866_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_865_v56, _785_v4)
					var _867_v58 _dafny.Map
					_ = _867_v58
					_867_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v4, _828_v34)
					var _868_v59 _dafny.Map
					_ = _868_v59
					_868_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_828_v34, _785_v4)
					var _869_v60 _dafny.Array
					_ = _869_v60
					var _nwElement0_27 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
						if (_866_v57).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_788_v7).Cardinality()), _783_v2)) {
							return (_866_v57).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_788_v7).Cardinality()), _783_v2)).(_dafny.Array)
						}
						return _785_v4
					})(), _828_v34)
					_ = _nwElement0_27
					var _nw140 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(14))
					_ = _nw140
					(_nw140).ArraySet1(_nwElement0_27, 0)
					(_nw140).ArraySet1(_867_v58, 1)
					(_nw140).ArraySet1(_867_v58, 2)
					(_nw140).ArraySet1((_867_v58).Merge(_867_v58), 3)
					(_nw140).ArraySet1(_867_v58, 4)
					(_nw140).ArraySet1(_867_v58, 5)
					(_nw140).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v4, _828_v34), 6)
					(_nw140).ArraySet1((_867_v58).Update((func() _dafny.Array {
						if (_868_v59).Contains(_828_v34) {
							return (_868_v59).Get(_828_v34).(_dafny.Array)
						}
						return _785_v4
					})(), _828_v34), 7)
					(_nw140).ArraySet1((_867_v58).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v4, _828_v34)), 8)
					(_nw140).ArraySet1(_867_v58, 9)
					(_nw140).ArraySet1((func() _dafny.Map {
						if _783_v2 {
							return _867_v58
						}
						return _867_v58
					})(), 10)
					(_nw140).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v4, _828_v34)).Merge(_867_v58), 11)
					(_nw140).ArraySet1(_867_v58, 12)
					(_nw140).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v4, _828_v34), 13)
					_869_v60 = _nw140
					var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_869_v60), 0))
					_ = _index152
					(_869_v60).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v4, _828_v34), (_index152).Int())
					var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_869_v60), 0))
					_ = _index153
					(_869_v60).ArraySet1(_867_v58, (_index153).Int())
					var _870_v61 _dafny.Sequence
					_ = _870_v61
					_870_v61 = _dafny.SeqOf(_785_v4, _785_v4)
					var _871_v62 _dafny.Array
					_ = _871_v62
					var _nwElement0_28 _dafny.Array = _785_v4
					_ = _nwElement0_28
					var _nw141 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(17))
					_ = _nw141
					(_nw141).ArraySet1(_nwElement0_28, 0)
					(_nw141).ArraySet1(_785_v4, 1)
					(_nw141).ArraySet1(_785_v4, 2)
					(_nw141).ArraySet1(_785_v4, 3)
					(_nw141).ArraySet1(_785_v4, 4)
					(_nw141).ArraySet1(_785_v4, 5)
					(_nw141).ArraySet1(_785_v4, 6)
					(_nw141).ArraySet1(_785_v4, 7)
					(_nw141).ArraySet1(_785_v4, 8)
					(_nw141).ArraySet1(_785_v4, 9)
					(_nw141).ArraySet1(_785_v4, 10)
					(_nw141).ArraySet1((_870_v61).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_870_v61).Cardinality()))).Uint32()).(_dafny.Array), 11)
					(_nw141).ArraySet1(_785_v4, 12)
					(_nw141).ArraySet1(_785_v4, 13)
					(_nw141).ArraySet1(_785_v4, 14)
					(_nw141).ArraySet1(_785_v4, 15)
					(_nw141).ArraySet1(_785_v4, 16)
					_871_v62 = _nw141
					var _872_v63 D4
					_ = _872_v63
					_872_v63 = Companion_D4_.Create_DC8_(_871_v62)
					var _source19 D4 = _872_v63
					_ = _source19
					if _source19.Is_DC9() {
						var _873___mcc_h5 _dafny.Int = _source19.Get_().(D4_DC9).Cf12
						_ = _873___mcc_h5
						var _874___mcc_h6 _dafny.Sequence = _source19.Get_().(D4_DC9).Cf13
						_ = _874___mcc_h6
						var _875___mcc_h7 _dafny.Int = _source19.Get_().(D4_DC9).Cf14
						_ = _875___mcc_h7
						var _876___mcc_h8 _dafny.Int = _source19.Get_().(D4_DC9).Cf15
						_ = _876___mcc_h8
						var _877_cf15 _dafny.Int = _876___mcc_h8
						_ = _877_cf15
						var _878_cf14 _dafny.Int = _875___mcc_h7
						_ = _878_cf14
						var _879_cf13 _dafny.Sequence = _874___mcc_h6
						_ = _879_cf13
						var _880_cf12 _dafny.Int = _873___mcc_h5
						_ = _880_cf12
						var _881_v64 _dafny.Sequence
						_ = _881_v64
						_881_v64 = _dafny.SeqOf(_782_v1, _dafny.SeqOf(_828_v34, _828_v34), _782_v1)
						_878_cf14 = Companion_Default___.SafeModuloInt(((_789_v8).F7()).Times((_784_v3).Cardinality()), ((func() _dafny.Int {
							if (_784_v3).Contains(_783_v2) {
								return (_784_v3).Multiplicity(_783_v2)
							}
							return (_789_v8).F8()
						})()).Plus(_dafny.IntOfUint32((_881_v64).Cardinality())))
						var _882_v65 _dafny.Map
						_ = _882_v65
						_882_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_783_v2, _dafny.IntOfInt64(647))
						var _883_v66 _dafny.Set
						_ = _883_v66
						_883_v66 = _dafny.SetOf((_789_v8).F8(), Companion_Default___.SafeDivisionInt((_789_v8).F7(), _dafny.IntOfInt64(901)), (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
							if (_882_v65).Contains(_783_v2) {
								return (_882_v65).Get(_783_v2).(_dafny.Int)
							}
							return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hcam")).Cardinality())
						})())), _dafny.IntOfInt64(330))
						_883_v66 = ((_883_v66).Intersection(_dafny.SetOf((_789_v8).F7()))).Difference(_883_v66)
						var _884_v67 _dafny.Array
						_ = _884_v67
						var _nw142 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
						_ = _nw142
						_884_v67 = _nw142
						_884_v67 = (func() _dafny.Array {
							if (_783_v2) == (_783_v2) {
								return _884_v67
							}
							return _884_v67
						})()
						var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_884_v67), 0))
						_ = _index154
						(_884_v67).ArraySet1(_783_v2, (_index154).Int())
						var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_884_v67), 0))
						_ = _index155
						(_884_v67).ArraySet1(!(_783_v2) || (_783_v2), (_index155).Int())
					} else if _source19.Is_DC10() {
						var _885___mcc_h9 _dafny.Int = _source19.Get_().(D4_DC10).Cf16
						_ = _885___mcc_h9
						var _886___mcc_h10 bool = _source19.Get_().(D4_DC10).Cf17
						_ = _886___mcc_h10
						var _887_cf17 bool = _886___mcc_h10
						_ = _887_cf17
						var _888_cf16 _dafny.Int = _885___mcc_h9
						_ = _888_cf16
						var _889_v68 *C5
						_ = _889_v68
						var _nw143 *C5 = New_C5_()
						_ = _nw143
						_nw143.Ctor__(_782_v1, _dafny.IntOfInt64(386), (_786_v5).Cardinality(), (_789_v8).F7())
						_889_v68 = _nw143
						var _890_v69 _dafny.Map
						_ = _890_v69
						_890_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_888_cf16, _889_v68)
						var _891_v70 _dafny.Map
						_ = _891_v70
						_891_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _890_v69)
						_891_v70 = (_891_v70).Update((_789_v8).F7(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _889_v68)).Update(_dafny.IntOfInt64(358), _889_v68))
						var _892_v71 _dafny.Map
						_ = _892_v71
						_892_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_783_v2), (_889_v68).F9())
						var _893_v72 _dafny.Map
						_ = _893_v72
						_893_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _783_v2)
						var _894_v73 D7
						_ = _894_v73
						_894_v73 = Companion_D7_.Create_DC20_(_828_v34, (_889_v68).F10())
						var _895_v74 _dafny.Sequence
						_ = _895_v74
						_895_v74 = _dafny.SeqOf(_894_v73)
						var _896_v75 *C0
						_ = _896_v75
						var _nw144 *C0 = New_C0_()
						_ = _nw144
						_nw144.Ctor__(_785_v4, _dafny.IntOfInt64(741), (_dafny.MultiSetOf(_887_cf17)).Cardinality())
						_896_v75 = _nw144
						var _897_v76 _dafny.Map
						_ = _897_v76
						_897_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v4, _896_v75)
						var _898_v77 _dafny.Array
						_ = _898_v77
						var _nwElement0_29 bool = _783_v2
						_ = _nwElement0_29
						var _nw145 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(26))
						_ = _nw145
						(_nw145).ArraySet1(_nwElement0_29, 0)
						(_nw145).ArraySet1(_783_v2, 1)
						(_nw145).ArraySet1(!(_dafny.Companion_Sequence_.IsPrefixOf((_889_v68).F9(), (func() _dafny.Sequence {
							if (_892_v71).Contains(_887_cf17) {
								return (_892_v71).Get(_887_cf17).(_dafny.Sequence)
							}
							return _dafny.UnicodeSeqOfUtf8Bytes("ejyl")
						})())), 2)
						(_nw145).ArraySet1(_783_v2, 3)
						(_nw145).ArraySet1(false, 4)
						(_nw145).ArraySet1(((_this).F8()).Cmp(_dafny.IntOfInt64(-50)) >= 0, 5)
						(_nw145).ArraySet1((Companion_D5_.Create_DC12_(_783_v2)).Dtor_cf19(), 6)
						(_nw145).ArraySet1((_887_cf17) && (_887_cf17), 7)
						(_nw145).ArraySet1((_783_v2) == (_783_v2), 8)
						(_nw145).ArraySet1((_893_v72).Contains(false), 9)
						(_nw145).ArraySet1(false, 10)
						(_nw145).ArraySet1(!(!_dafny.Companion_Sequence_.Contains(_782_v1, _dafny.CodePoint('q'))), 11)
						(_nw145).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_895_v74, _895_v74), 12)
						(_nw145).ArraySet1(_783_v2, 13)
						(_nw145).ArraySet1((func() bool {
							if (_865_v56).Contains((_789_v8).F7()) {
								return (_865_v56).Get((_789_v8).F7()).(bool)
							}
							return (func() bool {
								if (_865_v56).Contains(_dafny.IntOfInt64(-911)) {
									return (_865_v56).Get(_dafny.IntOfInt64(-911)).(bool)
								}
								return _783_v2
							})()
						})(), 14)
						(_nw145).ArraySet1(_887_cf17, 15)
						(_nw145).ArraySet1(_783_v2, 16)
						(_nw145).ArraySet1(_783_v2, 17)
						(_nw145).ArraySet1((Companion_Default___.Fm32(globalState)).Dtor_cf20(), 18)
						(_nw145).ArraySet1(_887_cf17, 19)
						(_nw145).ArraySet1(!(_887_cf17), 20)
						(_nw145).ArraySet1((_dafny.IntOfInt64(391)).Cmp((_789_v8).F7()) >= 0, 21)
						(_nw145).ArraySet1(_887_cf17, 22)
						(_nw145).ArraySet1((_897_v76).Contains(_785_v4), 23)
						(_nw145).ArraySet1(_783_v2, 24)
						(_nw145).ArraySet1(_887_cf17, 25)
						_898_v77 = _nw145
						_898_v77 = _898_v77
						(globalState).F4 = _783_v2
						var _899_v78 _dafny.Set
						_ = _899_v78
						_899_v78 = _dafny.SetOf(_782_v1)
						var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(432), _dafny.ArrayLen((_898_v77), 0))
						_ = _index156
						(_898_v77).ArraySet1((_899_v78).Contains(_dafny.UnicodeSeqOfUtf8Bytes("rhndgwb")), (_index156).Int())
						var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_785_v4), 0))
						_ = _index157
						(_785_v4).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((func() _dafny.Int {
							if (_786_v5).Contains((_789_v8).F8()) {
								return (_786_v5).Multiplicity((_789_v8).F8())
							}
							return _dafny.IntOfInt64(718)
						})()), _888_cf16), (_index157).Int())
						var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(432), _dafny.ArrayLen((_898_v77), 0))
						_ = _index158
						var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_785_v4), 0))
						_ = _index159
						var _rhs150 bool = !((Companion_Default___.SafeDivisionInt((_789_v8).F7(), (_789_v8).F8())).Cmp((_789_v8).F8()) < 0)
						_ = _rhs150
						var _rhs151 _dafny.CodePoint = Companion_Default___.Fm5(globalState)
						_ = _rhs151
						var _rhs152 _dafny.Int = (_789_v8).F8()
						_ = _rhs152
						var _rhs153 _dafny.Int = (_789_v8).F8()
						_ = _rhs153
						var _rhs154 bool = (_888_cf16).Cmp((_789_v8).F7()) == 0
						_ = _rhs154
						var _lhs115 _dafny.Array = _898_v77
						_ = _lhs115
						var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(432), _dafny.ArrayLen((_898_v77), 0))
						_ = _lhs116
						var _lhs117 _dafny.Array = _785_v4
						_ = _lhs117
						var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_785_v4), 0))
						_ = _lhs118
						var _lhs119 *GlobalState = globalState
						_ = _lhs119
						(_lhs115).ArraySet1(_rhs150, (_lhs116).Int())
						_828_v34 = _rhs151
						_888_cf16 = _rhs152
						(_lhs117).ArraySet1(_rhs153, (_lhs118).Int())
						_lhs119.F4 = _rhs154
					} else {
						var _900___mcc_h11 _dafny.Array = _source19.Get_().(D4_DC8).Cf11
						_ = _900___mcc_h11
						var _901_cf11 _dafny.Array = _900___mcc_h11
						_ = _901_cf11
						var _902_v79 _dafny.Set
						_ = _902_v79
						_902_v79 = _dafny.SetOf((_789_v8).F8())
						var _903_v80 _dafny.Map
						_ = _903_v80
						_903_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_782_v1).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_782_v1).Cardinality()))).Uint32()).(_dafny.CodePoint), _902_v79)
						_902_v79 = (func() _dafny.Set {
							if (_903_v80).Contains(_828_v34) {
								return (_903_v80).Get(_828_v34).(_dafny.Set)
							}
							return _902_v79
						})()
						_781_v0 = ((_789_v8).F7()).Minus((_789_v8).F7())
						var _904_v81 _dafny.MultiSet
						_ = _904_v81
						_904_v81 = _dafny.MultiSetOf(_828_v34)
						(globalState).F4 = (func() bool {
							if _783_v2 {
								return _783_v2
							}
							return !(_904_v81).Contains(_828_v34)
						})()
						var _905_v82 _dafny.Array
						_ = _905_v82
						var _nw146 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
						_ = _nw146
						_905_v82 = _nw146
						var _906_v83 _dafny.Array
						_ = _906_v83
						var _nwElement0_30 _dafny.Array = _905_v82
						_ = _nwElement0_30
						var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(4))
						_ = _nw147
						(_nw147).ArraySet1(_nwElement0_30, 0)
						(_nw147).ArraySet1((func() _dafny.Array {
							if _783_v2 {
								return _905_v82
							}
							return _905_v82
						})(), 1)
						(_nw147).ArraySet1(_905_v82, 2)
						(_nw147).ArraySet1(_905_v82, 3)
						_906_v83 = _nw147
						var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_906_v83), 0))
						_ = _index160
						(_906_v83).ArraySet1(_905_v82, (_index160).Int())
						var _907_v84 _dafny.Map
						_ = _907_v84
						_907_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_783_v2), _783_v2)
						var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_906_v83), 0))
						_ = _index161
						var _rhs155 _dafny.Array = _785_v4
						_ = _rhs155
						var _rhs156 _dafny.Int = _dafny.IntOfInt64(151)
						_ = _rhs156
						var _rhs157 _dafny.Array = _905_v82
						_ = _rhs157
						var _rhs158 _dafny.Int = _781_v0
						_ = _rhs158
						var _rhs159 bool = (func() bool {
							if (_907_v84).Contains(_783_v2) {
								return (_907_v84).Get(_783_v2).(bool)
							}
							return _783_v2
						})()
						_ = _rhs159
						var _lhs120 _dafny.Array = _906_v83
						_ = _lhs120
						var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_906_v83), 0))
						_ = _lhs121
						var _lhs122 *GlobalState = globalState
						_ = _lhs122
						_785_v4 = _rhs155
						_781_v0 = _rhs156
						(_lhs120).ArraySet1(_rhs157, (_lhs121).Int())
						_781_v0 = _rhs158
						_lhs122.F4 = _rhs159
					}
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _908_v85 _dafny.CodePoint
		_ = _908_v85
		_908_v85 = _dafny.CodePoint('s')
		var _909_v86 _dafny.Map
		_ = _909_v86
		_909_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(276))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg59 _dafny.Int) interface{} {
				return coer59(arg59)
			}
		}((func(_910_v2 bool) func(_dafny.Int) _dafny.Int {
			return func(_911_i10 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_910_v2, _910_v2)).Cardinality()
			}
		})(_783_v2))), _788_v7), _908_v85)
		_909_v86 = (_909_v86).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-981))).Uint32(), func(coer60 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg60 _dafny.Int) interface{} {
				return coer60(arg60)
			}
		}((func(_912_v8 T1) func(_dafny.Int) _dafny.Int {
			return func(_913_i11 _dafny.Int) _dafny.Int {
				return (_912_v8).F7()
			}
		})(_789_v8))), _908_v85)
		var _914_v87 _dafny.Map
		_ = _914_v87
		_914_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_781_v0, (_789_v8).F8())
		_914_v87 = (_914_v87).Update((_this).F8(), (_dafny.IntOfInt64(-460)).Times((_789_v8).F7()))
		var _915_v88 _dafny.Map
		_ = _915_v88
		_915_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_783_v2, _914_v87)
		var _916_v89 _dafny.Sequence
		_ = _916_v89
		_916_v89 = _dafny.SeqOf(false, _783_v2, _783_v2)
		_915_v88 = (_915_v88).Update((_916_v89).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_916_v89).Cardinality()))).Uint32()).(bool), (_914_v87).Merge(_914_v87))
	}
}
func (_this *C6) M2(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		if p3 {
			(globalState).F4 = true
			var _917_v0 _dafny.Sequence
			_ = _917_v0
			_917_v0 = _dafny.UnicodeSeqOfUtf8Bytes("bjrtfrpd")
			_917_v0 = _dafny.Companion_Sequence_.Concatenate(_917_v0, _917_v0)
			(globalState).F4 = ((_this).F7()).Cmp(p0) != 0
			var _918_v1 _dafny.CodePoint
			_ = _918_v1
			_918_v1 = _dafny.CodePoint('q')
			var _919_v2 D5
			_ = _919_v2
			_919_v2 = Companion_D5_.Create_DC14_(Companion_D5_.Create_DC13_(p3, _918_v1))
			_919_v2 = (func() D5 {
				if p3 {
					return _919_v2
				}
				return _919_v2
			})()
			var _920_v3 _dafny.Array
			_ = _920_v3
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_22
			var _nw148 _dafny.Array
			_ = _nw148
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw148 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Int = func(_921_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_921_i0, (_this).F7())
				}
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
			_920_v3 = _nw148
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_920_v3), 0))
			_ = _index162
			(_920_v3).ArraySet1(p2, (_index162).Int())
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_920_v3), 0))
			_ = _index163
			(_920_v3).ArraySet1((_this).F7(), (_index163).Int())
		} else {
			var _922_v4 _dafny.MultiSet
			_ = _922_v4
			_922_v4 = _dafny.MultiSetOf(p3)
			var _923_v5 _dafny.Map
			_ = _923_v5
			_923_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_922_v4, p2)
			var _924_v6 _dafny.MultiSet
			_ = _924_v6
			_924_v6 = _dafny.MultiSetOf(_dafny.IntOfInt64(-253))
			var _925_v7 _dafny.Sequence
			_ = _925_v7
			_925_v7 = _dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('h'))
			var _926_v8 _dafny.Sequence
			_ = _926_v8
			_926_v8 = _dafny.SeqOf(Companion_Default___.Fm24(p3, (_dafny.Zero).Minus((_922_v4).Cardinality()), _925_v7, globalState))
			var _927_v9 D6
			_ = _927_v9
			_927_v9 = Companion_D6_.Create_DC16_((_923_v5).Update(_922_v4, (_924_v6).Cardinality()), _926_v8)
			var _source20 D6 = _927_v9
			_ = _source20
			if _source20.Is_DC16() {
				var _928___mcc_h0 _dafny.Map = _source20.Get_().(D6_DC16).Cf24
				_ = _928___mcc_h0
				var _929___mcc_h1 _dafny.Sequence = _source20.Get_().(D6_DC16).Cf25
				_ = _929___mcc_h1
				var _930_cf25 _dafny.Sequence = _929___mcc_h1
				_ = _930_cf25
				var _931_cf24 _dafny.Map = _928___mcc_h0
				_ = _931_cf24
				var _932_v10 _dafny.CodePoint
				_ = _932_v10
				_932_v10 = _dafny.CodePoint('o')
				var _933_v11 D7
				_ = _933_v11
				_933_v11 = Companion_D7_.Create_DC20_(_932_v10, (_this).F7())
				r1 = (_933_v11).Dtor_cf33()
				var _934_v12 T1
				_ = _934_v12
				var _nw149 *C4 = New_C4_()
				_ = _nw149
				_nw149.Ctor__(p1, p0, (_this).F7())
				_934_v12 = _nw149
				var _935_v13 _dafny.MultiSet
				_ = _935_v13
				_935_v13 = _dafny.MultiSetOf(_934_v12)
				var _936_v14 _dafny.Map
				_ = _936_v14
				_936_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
				var _937_v15 _dafny.Map
				_ = _937_v15
				_937_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_934_v12).F8(), p0)
				var _938_v16 _dafny.Map
				_ = _938_v16
				_938_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_925_v7, (func() _dafny.Int {
					if (_922_v4).Contains(p1) {
						return (_922_v4).Multiplicity(p1)
					}
					return p2
				})())
				var _939_v17 _dafny.Sequence
				_ = _939_v17
				_939_v17 = _dafny.SeqOf(_dafny.IntOfUint32((_930_cf25).Cardinality()), (func() _dafny.Int {
					if (_938_v16).Contains(_925_v7) {
						return (_938_v16).Get(_925_v7).(_dafny.Int)
					}
					return (_934_v12).F7()
				})(), (_934_v12).F7(), (_934_v12).F7())
				var _940_v18 _dafny.Set
				_ = _940_v18
				_940_v18 = _dafny.SetOf(true, p3, p3)
				var _941_v19 _dafny.Array
				_ = _941_v19
				var _nwElement0_31 _dafny.Int = _dafny.IntOfInt64(886)
				_ = _nwElement0_31
				var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(18))
				_ = _nw150
				(_nw150).ArraySet1(_nwElement0_31, 0)
				(_nw150).ArraySet1((_this).F8(), 1)
				(_nw150).ArraySet1(Companion_Default___.SafeDivisionInt(p2, (_dafny.Zero).Minus(p2)), 2)
				(_nw150).ArraySet1(p0, 3)
				(_nw150).ArraySet1((_this).F7(), 4)
				(_nw150).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yiww")).Cardinality()), 5)
				(_nw150).ArraySet1((_935_v13).Cardinality(), 6)
				(_nw150).ArraySet1((_this).F7(), 7)
				(_nw150).ArraySet1(((_934_v12).F7()).Minus(_dafny.IntOfInt64(636)), 8)
				(_nw150).ArraySet1((_this).F8(), 9)
				(_nw150).ArraySet1((_934_v12).F7(), 10)
				(_nw150).ArraySet1((_this).F7(), 11)
				(_nw150).ArraySet1((_934_v12).F8(), 12)
				(_nw150).ArraySet1(p0, 13)
				(_nw150).ArraySet1(((_936_v14).Cardinality()).Plus(p0), 14)
				(_nw150).ArraySet1(((func() _dafny.Int {
					if (_922_v4).Contains(!(p1)) {
						return (_922_v4).Multiplicity(!(p1))
					}
					return (func() _dafny.Int {
						if (_937_v15).Contains((_934_v12).F7()) {
							return (_937_v15).Get((_934_v12).F7()).(_dafny.Int)
						}
						return (_this).F8()
					})()
				})()).Plus((_939_v17).Select((Companion_Default___.SafeIndex((_940_v18).Cardinality(), _dafny.IntOfUint32((_939_v17).Cardinality()))).Uint32()).(_dafny.Int)), 15)
				(_nw150).ArraySet1((func() _dafny.Int {
					if p1 {
						return (_this).F8()
					}
					return _dafny.IntOfInt64(-589)
				})(), 16)
				(_nw150).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_930_cf25).Cardinality()), _dafny.IntOfInt64(755)), 17)
				_941_v19 = _nw150
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_941_v19), 0))
				_ = _index164
				(_941_v19).ArraySet1(Companion_Default___.SafeDivisionInt(p0, p2), (_index164).Int())
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_941_v19), 0))
				_ = _index165
				(_941_v19).ArraySet1((_this).F8(), (_index165).Int())
				var _942_v20 _dafny.Int
				_ = _942_v20
				var _943_v21 _dafny.Int
				_ = _943_v21
				var _944_v22 bool
				_ = _944_v22
				var _out13 _dafny.Int
				_ = _out13
				var _out14 _dafny.Int
				_ = _out14
				var _out15 bool
				_ = _out15
				_out13, _out14, _out15 = (_this).M3(p3, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(689), _dafny.IntOfUint32((_925_v7).Cardinality()))).Cardinality(), _941_v19, (_this).F8(), globalState)
				_942_v20 = _out13
				_943_v21 = _out14
				_944_v22 = _out15
				var _945_v23 _dafny.Int
				_ = _945_v23
				var _946_v24 _dafny.Int
				_ = _946_v24
				var _947_v25 bool
				_ = _947_v25
				var _out16 _dafny.Int
				_ = _out16
				var _out17 _dafny.Int
				_ = _out17
				var _out18 bool
				_ = _out18
				_out16, _out17, _out18 = (_this).M3(false, ((_this).F8()).Minus((_934_v12).F7()), _941_v19, (_dafny.Zero).Minus((_this).F7()), globalState)
				_945_v23 = _out16
				_946_v24 = _out17
				_947_v25 = _out18
			} else {
				var _948___mcc_h2 _dafny.Array = _source20.Get_().(D6_DC15).Cf23
				_ = _948___mcc_h2
				var _949_cf23 _dafny.Array = _948___mcc_h2
				_ = _949_cf23
				r1 = (_this).F7()
				var _950_v26 _dafny.CodePoint
				_ = _950_v26
				_950_v26 = _dafny.CodePoint('e')
				var _951_v27 D7
				_ = _951_v27
				_951_v27 = Companion_D7_.Create_DC20_((func() _dafny.CodePoint {
					if p3 {
						return _950_v26
					}
					return _950_v26
				})(), (_this).F8())
				var _952_v28 D5
				_ = _952_v28
				_952_v28 = Companion_D5_.Create_DC13_(p1, _950_v26)
				var _953_v29 _dafny.Array
				_ = _953_v29
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_23
				var _nw151 _dafny.Array
				_ = _nw151
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw151 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Int = (func(_954_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_955_i1 _dafny.Int) _dafny.Int {
							return (_955_i1).Times(_954_p0)
						}
					})(p0)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw151 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw151).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw151).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_953_v29 = _nw151
				var _956_v30 _dafny.Map
				_ = _956_v30
				_956_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_952_v28, _953_v29)
				var _957_v31 _dafny.Sequence
				_ = _957_v31
				_957_v31 = _dafny.SeqOf(_956_v30)
				var _rhs160 bool = ((_957_v31).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_957_v31).Cardinality()))).Uint32()).(_dafny.Map)).Equals(((_956_v30).Update(_952_v28, _953_v29)).Update(_952_v28, _953_v29))
				_ = _rhs160
				var _rhs161 bool = (func() bool {
					if p3 {
						return p1
					}
					return p1
				})()
				_ = _rhs161
				var _rhs162 D7 = Companion_D7_.Create_DC20_(_950_v26, (_dafny.Zero).Minus(p0))
				_ = _rhs162
				var _lhs123 *GlobalState = globalState
				_ = _lhs123
				var _lhs124 *GlobalState = globalState
				_ = _lhs124
				_lhs123.F3 = _rhs160
				_lhs124.F3 = _rhs161
				_951_v27 = _rhs162
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_949_cf23), 0))
				_ = _index166
				(_949_cf23).ArraySet1(p3, (_index166).Int())
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_949_cf23), 0))
				_ = _index167
				(_949_cf23).ArraySet1(!(p3), (_index167).Int())
				r1 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(873))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}((func(_958_v26 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_959_i2 _dafny.Int) _dafny.CodePoint {
						return _958_v26
					}
				})(_950_v26))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(19))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}(func(_960_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('i')
				})))).Cardinality())).Times(_dafny.IntOfInt64(-398))
			}
			var _961_v32 _dafny.Map
			_ = _961_v32
			_961_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfInt64(662))
			var _962_v33 D5
			_ = _962_v33
			_962_v33 = Companion_D5_.Create_DC11_(_961_v32)
			var _source21 D5 = _962_v33
			_ = _source21
			if _source21.Is_DC12() {
				var _963___mcc_h3 bool = _source21.Get_().(D5_DC12).Cf19
				_ = _963___mcc_h3
				var _964_cf19 bool = _963___mcc_h3
				_ = _964_cf19
				var _965_v34 _dafny.Sequence
				_ = _965_v34
				_965_v34 = _dafny.SeqOf(_927_v9)
				var _966_v35 D3
				_ = _966_v35
				_966_v35 = Companion_D3_.Create_DC6_(_dafny.SetOf(_dafny.IntOfInt64(-311), p0, (_this).F7(), (_this).F8(), (_this).F8()))
				var _967_v36 _dafny.Set
				_ = _967_v36
				_967_v36 = _dafny.SetOf(p2, (_this).F8())
				var _pat_let_tv14 = _967_v36
				_ = _pat_let_tv14
				var _968_v37 _dafny.Map
				_ = _968_v37
				_968_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_965_v34).Cardinality()), (_this).F8()), func(_pat_let12_0 D3) D3 {
					return func(_969_dt__update__tmp_h0 D3) D3 {
						return func(_pat_let13_0 _dafny.Set) D3 {
							return func(_970_dt__update_hcf9_h0 _dafny.Set) D3 {
								return Companion_D3_.Create_DC6_(_970_dt__update_hcf9_h0)
							}(_pat_let13_0)
						}(_pat_let_tv14)
					}(_pat_let12_0)
				}(_966_v35))
				_968_v37 = (_968_v37).Update(((_dafny.Zero).Minus(p0)).Times(p0), _966_v35)
				_925_v7 = _925_v7
				(globalState).F4 = p1
				r1 = _dafny.IntOfInt64(-788)
			} else if _source21.Is_DC13() {
				var _971___mcc_h4 bool = _source21.Get_().(D5_DC13).Cf20
				_ = _971___mcc_h4
				var _972___mcc_h5 _dafny.CodePoint = _source21.Get_().(D5_DC13).Cf21
				_ = _972___mcc_h5
				var _973_cf21 _dafny.CodePoint = _972___mcc_h5
				_ = _973_cf21
				var _974_cf20 bool = _971___mcc_h4
				_ = _974_cf20
				var _975_v38 _dafny.Array
				_ = _975_v38
				var _nwElement0_32 bool = ((_this).F7()).Cmp((_this).F7()) < 0
				_ = _nwElement0_32
				var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(8))
				_ = _nw152
				(_nw152).ArraySet1(_nwElement0_32, 0)
				(_nw152).ArraySet1((func() bool {
					if p3 {
						return p3
					}
					return _974_cf20
				})(), 1)
				(_nw152).ArraySet1(p1, 2)
				(_nw152).ArraySet1(true, 3)
				(_nw152).ArraySet1(_974_cf20, 4)
				(_nw152).ArraySet1(_974_cf20, 5)
				(_nw152).ArraySet1(_974_cf20, 6)
				(_nw152).ArraySet1(_974_cf20, 7)
				_975_v38 = _nw152
				var _rhs163 _dafny.Int = p0
				_ = _rhs163
				var _rhs164 _dafny.Array = _975_v38
				_ = _rhs164
				var _rhs165 _dafny.Int = (p2).Minus(Companion_Default___.SafeDivisionInt(p2, (_this).Fm3(_974_cf20, p1, p3, true, globalState)))
				_ = _rhs165
				r1 = _rhs163
				_975_v38 = _rhs164
				r1 = _rhs165
				var _976_v39 _dafny.Sequence
				_ = _976_v39
				_976_v39 = _dafny.SeqOf(_975_v38)
				var _977_v40 _dafny.Array
				_ = _977_v40
				var _nw153 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
				_ = _nw153
				_977_v40 = _nw153
				var _978_v41 _dafny.Sequence
				_ = _978_v41
				_978_v41 = _dafny.SeqOf((_this).F8())
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_977_v40), 0))
				_ = _index168
				(_977_v40).ArraySet1(_978_v41, (_index168).Int())
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_977_v40), 0))
				_ = _index169
				var _rhs166 _dafny.Sequence = (func() _dafny.Sequence {
					if p1 {
						return _976_v39
					}
					return _976_v39
				})()
				_ = _rhs166
				var _rhs167 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_978_v41, _dafny.SeqOf((_this).F7())), _dafny.Companion_Sequence_.Concatenate(_978_v41, _978_v41))
				_ = _rhs167
				var _rhs168 _dafny.Sequence = _925_v7
				_ = _rhs168
				var _rhs169 _dafny.CodePoint = _973_cf21
				_ = _rhs169
				var _lhs125 _dafny.Array = _977_v40
				_ = _lhs125
				var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_977_v40), 0))
				_ = _lhs126
				_976_v39 = _rhs166
				(_lhs125).ArraySet1(_rhs167, (_lhs126).Int())
				_925_v7 = _rhs168
				_973_cf21 = _rhs169
				var _979_v42 *C1
				_ = _979_v42
				var _nw154 *C1 = New_C1_()
				_ = _nw154
				_nw154.Ctor__((_this).F7(), (_dafny.Zero).Minus(p0))
				_979_v42 = _nw154
				var _980_v43 _dafny.Map
				_ = _980_v43
				_980_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, !(p3))
				var _981_v44 _dafny.Set
				_ = _981_v44
				_981_v44 = _dafny.SetOf((_this).F7(), p2, p0, (_dafny.Zero).Minus((_980_v43).Cardinality()))
				var _982_v45 _dafny.Map
				_ = _982_v45
				_982_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_979_v42, (_981_v44).Cardinality())
				_982_v45 = (_982_v45).Update(_979_v42, p0)
				var _983_v46 _dafny.Array
				_ = _983_v46
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_24
				var _nw155 _dafny.Array
				_ = _nw155
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw155 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.Int = func(_984_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_984_i4, (_this).F8())
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
				_983_v46 = _nw155
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_983_v46), 0))
				_ = _index170
				(_983_v46).ArraySet1(Companion_Default___.SafeDivisionInt(p2, (_this).F7()), (_index170).Int())
				var _985_v47 _dafny.Map
				_ = _985_v47
				_985_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _dafny.IntOfUint32((_925_v7).Cardinality()))
				var _986_v48 _dafny.Map
				_ = _986_v48
				_986_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _985_v47)
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_983_v46), 0))
				_ = _index171
				(_983_v46).ArraySet1((Companion_Default___.SafeModuloInt(p0, ((func() _dafny.Map {
					if (_986_v48).Contains(p3) {
						return (_986_v48).Get(p3).(_dafny.Map)
					}
					return _985_v47
				})()).Cardinality())).Plus((_this).F8()), (_index171).Int())
			} else if _source21.Is_DC11() {
				var _987___mcc_h6 _dafny.Map = _source21.Get_().(D5_DC11).Cf18
				_ = _987___mcc_h6
				var _988_cf18 _dafny.Map = _987___mcc_h6
				_ = _988_cf18
				(globalState).F3 = p3
				var _989_v49 _dafny.CodePoint
				_ = _989_v49
				_989_v49 = _dafny.CodePoint('j')
				var _990_v50 _dafny.MultiSet
				_ = _990_v50
				_990_v50 = _dafny.MultiSetOf(_989_v49)
				var _991_v51 _dafny.Sequence
				_ = _991_v51
				_991_v51 = _dafny.SeqOf((_990_v50).Cardinality(), p0, p2)
				var _992_v52 *C5
				_ = _992_v52
				var _nw156 *C5 = New_C5_()
				_ = _nw156
				_nw156.Ctor__(_925_v7, (_this).F7(), (_991_v51).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_991_v51).Cardinality()))).Uint32()).(_dafny.Int), (_this).F8())
				_992_v52 = _nw156
				r1 = (_this).F7()
				r1 = p2
			} else {
				var _993___mcc_h7 D5 = _source21.Get_().(D5_DC14).Cf22
				_ = _993___mcc_h7
				var _994_cf22 D5 = _993___mcc_h7
				_ = _994_cf22
				_961_v32 = (_961_v32).Update(true, (_dafny.Zero).Minus((_this).F7()))
				r1 = ((_this).F8()).Plus((func() _dafny.Int {
					if p3 {
						return (_this).F8()
					}
					return (_this).F8()
				})())
				var _995_v53 _dafny.Array
				_ = _995_v53
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_25
				var _nw157 _dafny.Array
				_ = _nw157
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw157 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) _dafny.Int = func(_996_i5 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_996_i5, _dafny.IntOfInt64(648))
					}
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw157 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw157).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw157).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_995_v53 = _nw157
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_995_v53), 0))
				_ = _index172
				(_995_v53).ArraySet1(_dafny.IntOfUint32((_925_v7).Cardinality()), (_index172).Int())
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_995_v53), 0))
				_ = _index173
				(_995_v53).ArraySet1(p0, (_index173).Int())
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_995_v53), 0))
				_ = _index174
				(_995_v53).ArraySet1(_dafny.IntOfInt64(758), (_index174).Int())
			}
			var _997_v54 D1
			_ = _997_v54
			_997_v54 = Companion_D1_.Create_DC3_(Companion_Default___.Fm23((_this).F8(), globalState))
			var _source22 D1 = _997_v54
			_ = _source22
			if _source22.Is_DC2() {
				var _998_v55 _dafny.Array
				_ = _998_v55
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_26
				var _nw158 _dafny.Array
				_ = _nw158
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw158 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) bool = (func(_999_p1 bool) func(_dafny.Int) bool {
						return func(_1000_i6 _dafny.Int) bool {
							return _999_p1
						}
					})(p1)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw158 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw158).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw158).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_998_v55 = _nw158
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_998_v55), 0))
				_ = _index175
				(_998_v55).ArraySet1(p1, (_index175).Int())
				var _1001_v56 _dafny.Array
				_ = _1001_v56
				var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
				_ = _nw159
				_1001_v56 = _nw159
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1001_v56), 0))
				_ = _index176
				(_1001_v56).ArraySet1(p2, (_index176).Int())
				var _1002_v57 *C1
				_ = _1002_v57
				var _nw160 *C1 = New_C1_()
				_ = _nw160
				_nw160.Ctor__((_this).F7(), (_this).F8())
				_1002_v57 = _nw160
				var _1003_v58 _dafny.Sequence
				_ = _1003_v58
				_1003_v58 = _dafny.SeqOf(_1002_v57)
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_998_v55), 0))
				_ = _index177
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1001_v56), 0))
				_ = _index178
				var _rhs170 bool = (p1) && (!_dafny.Companion_Sequence_.Equal(_1003_v58, _1003_v58))
				_ = _rhs170
				var _rhs171 bool = (p0).Cmp(_dafny.IntOfUint32((_926_v8).Cardinality())) < 0
				_ = _rhs171
				var _rhs172 _dafny.Int = Companion_Default___.Fm16(!(p1), globalState)
				_ = _rhs172
				var _rhs173 _dafny.Int = (p2).Plus(_dafny.IntOfInt64(738))
				_ = _rhs173
				var _lhs127 _dafny.Array = _998_v55
				_ = _lhs127
				var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_998_v55), 0))
				_ = _lhs128
				var _lhs129 *GlobalState = globalState
				_ = _lhs129
				var _lhs130 _dafny.Array = _1001_v56
				_ = _lhs130
				var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1001_v56), 0))
				_ = _lhs131
				(_lhs127).ArraySet1(_rhs170, (_lhs128).Int())
				_lhs129.F4 = _rhs171
				r1 = _rhs172
				(_lhs130).ArraySet1(_rhs173, (_lhs131).Int())
				var _1004_v59 _dafny.Map
				_ = _1004_v59
				_1004_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Int {
					if (_922_v4).Contains((_998_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_998_v55), 0))).Int()).(bool)) {
						return (_922_v4).Multiplicity((_998_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_998_v55), 0))).Int()).(bool))
					}
					return (_this).F8()
				})()).Minus((_this).F7()), p0)
				_1004_v59 = (_1004_v59).Update(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(594), _dafny.IntOfInt64(147)), p2)
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1001_v56), 0))
				_ = _index179
				(_1001_v56).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_925_v7, (func() _dafny.Sequence {
					if (_998_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_998_v55), 0))).Int()).(bool) {
						return _925_v7
					}
					return _925_v7
				})())).Cardinality()), (_index179).Int())
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1001_v56), 0))
				_ = _index180
				(_1001_v56).ArraySet1((p2).Times((_1001_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1001_v56), 0))).Int()).(_dafny.Int)), (_index180).Int())
			} else if _source22.Is_DC1() {
				var _1005___mcc_h8 _dafny.CodePoint = _source22.Get_().(D1_DC1).Cf1
				_ = _1005___mcc_h8
				var _1006_cf1 _dafny.CodePoint = _1005___mcc_h8
				_ = _1006_cf1
				var _1007_v60 _dafny.Map
				_ = _1007_v60
				_1007_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), p1)
				var _1008_v61 _dafny.Map
				_ = _1008_v61
				_1008_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)).Cardinality())
				var _1009_v62 _dafny.Sequence
				_ = _1009_v62
				_1009_v62 = _dafny.SeqOf(_1008_v61, _1008_v61)
				var _1010_v63 _dafny.Map
				_ = _1010_v63
				_1010_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_925_v7, p3)
				var _1011_v64 _dafny.Array
				_ = _1011_v64
				var _nwElement0_33 _dafny.Int = (_this).F8()
				_ = _nwElement0_33
				var _nw161 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(27))
				_ = _nw161
				(_nw161).ArraySet1(_nwElement0_33, 0)
				(_nw161).ArraySet1((_1007_v60).Cardinality(), 1)
				(_nw161).ArraySet1(Companion_Default___.Fm16(p1, globalState), 2)
				(_nw161).ArraySet1(Companion_Default___.Fm20(p3, p1, p1, globalState), 3)
				(_nw161).ArraySet1((_this).F7(), 4)
				(_nw161).ArraySet1((_dafny.Zero).Minus((_this).F7()), 5)
				(_nw161).ArraySet1((_this).F7(), 6)
				(_nw161).ArraySet1(p0, 7)
				(_nw161).ArraySet1(_dafny.IntOfUint32((_925_v7).Cardinality()), 8)
				(_nw161).ArraySet1(_dafny.IntOfUint32((_1009_v62).Cardinality()), 9)
				(_nw161).ArraySet1(_dafny.IntOfUint32((_926_v8).Cardinality()), 10)
				(_nw161).ArraySet1((_dafny.Zero).Minus((_this).F8()), 11)
				(_nw161).ArraySet1(_dafny.IntOfInt64(500), 12)
				(_nw161).ArraySet1(p0, 13)
				(_nw161).ArraySet1(p2, 14)
				(_nw161).ArraySet1(_dafny.IntOfInt64(532), 15)
				(_nw161).ArraySet1(p2, 16)
				(_nw161).ArraySet1(p2, 17)
				(_nw161).ArraySet1(_dafny.IntOfUint32((_925_v7).Cardinality()), 18)
				(_nw161).ArraySet1(p0, 19)
				(_nw161).ArraySet1((_this).F7(), 20)
				(_nw161).ArraySet1((_this).F8(), 21)
				(_nw161).ArraySet1((_this).F8(), 22)
				(_nw161).ArraySet1((_1010_v63).Cardinality(), 23)
				(_nw161).ArraySet1(_dafny.IntOfInt64(274), 24)
				(_nw161).ArraySet1(p2, 25)
				(_nw161).ArraySet1(p0, 26)
				_1011_v64 = _nw161
				var _1012_v65 _dafny.Sequence
				_ = _1012_v65
				_1012_v65 = _dafny.SeqOf(_1011_v64, _1011_v64, _1011_v64, _1011_v64, _1011_v64)
				var _1013_v66 _dafny.Array
				_ = _1013_v66
				var _nwElement0_34 _dafny.Array = _1011_v64
				_ = _nwElement0_34
				var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(10))
				_ = _nw162
				(_nw162).ArraySet1(_nwElement0_34, 0)
				(_nw162).ArraySet1(_1011_v64, 1)
				(_nw162).ArraySet1(_1011_v64, 2)
				(_nw162).ArraySet1(_1011_v64, 3)
				(_nw162).ArraySet1(_1011_v64, 4)
				(_nw162).ArraySet1(_1011_v64, 5)
				(_nw162).ArraySet1((_1012_v65).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_1012_v65).Cardinality()))).Uint32()).(_dafny.Array), 6)
				(_nw162).ArraySet1(_1011_v64, 7)
				(_nw162).ArraySet1(_1011_v64, 8)
				(_nw162).ArraySet1(_1011_v64, 9)
				_1013_v66 = _nw162
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_1013_v66), 0))
				_ = _index181
				(_1013_v66).ArraySet1(_1011_v64, (_index181).Int())
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_1013_v66), 0))
				_ = _index182
				(_1013_v66).ArraySet1(_1011_v64, (_index182).Int())
				var _1014_v67 _dafny.Set
				_ = _1014_v67
				_1014_v67 = _dafny.SetOf(_925_v7)
				(globalState).F4 = ((_1014_v67).Difference(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("hwlofrqbi")))).IsSubsetOf((func() _dafny.Set {
					if (_926_v8).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_926_v8).Cardinality()))).Uint32()).(bool) {
						return _1014_v67
					}
					return _1014_v67
				})())
				var _1015_v68 _dafny.Map
				_ = _1015_v68
				_1015_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _925_v7)
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_1011_v64), 0))
				_ = _index183
				(_1011_v64).ArraySet1((_1015_v68).Cardinality(), (_index183).Int())
				var _1016_v69 _dafny.Sequence
				_ = _1016_v69
				_1016_v69 = _dafny.SeqOf(p0)
				var _1017_v70 _dafny.Map
				_ = _1017_v70
				_1017_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1006_cf1, _925_v7)
				var _1018_v71 _dafny.Map
				_ = _1018_v71
				_1018_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (func() _dafny.Array {
					if p3 {
						return _1013_v66
					}
					return _1013_v66
				})())
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_1011_v64), 0))
				_ = _index184
				var _rhs174 _dafny.Int = ((_dafny.Zero).Minus((_this).F8())).Plus((_this).F7())
				_ = _rhs174
				var _rhs175 _dafny.Sequence = _1016_v69
				_ = _rhs175
				var _rhs176 bool = (func() bool {
					if !_dafny.Companion_Sequence_.Contains(_925_v7, _1006_cf1) {
						return p1
					}
					return (func() bool {
						if (_1007_v60).Contains((_this).F8()) {
							return (_1007_v60).Get((_this).F8()).(bool)
						}
						return !(false)
					})()
				})()
				_ = _rhs176
				var _rhs177 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F7(), (_1017_v70).Cardinality())
				_ = _rhs177
				var _rhs178 _dafny.Array = (func() _dafny.Array {
					if (_1018_v71).Contains(true) {
						return (_1018_v71).Get(true).(_dafny.Array)
					}
					return _1013_v66
				})()
				_ = _rhs178
				var _lhs132 _dafny.Array = _1011_v64
				_ = _lhs132
				var _lhs133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_1011_v64), 0))
				_ = _lhs133
				var _lhs134 *GlobalState = globalState
				_ = _lhs134
				(_lhs132).ArraySet1(_rhs174, (_lhs133).Int())
				_1016_v69 = _rhs175
				_lhs134.F3 = _rhs176
				r1 = _rhs177
				_1013_v66 = _rhs178
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_1011_v64), 0))
				_ = _index185
				(_1011_v64).ArraySet1(((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F7()), (_this).F8()))).Times(_dafny.IntOfInt64(646)), (_index185).Int())
			} else {
				var _1019___mcc_h9 D1 = _source22.Get_().(D1_DC3).Cf2
				_ = _1019___mcc_h9
				var _1020_cf2 D1 = _1019___mcc_h9
				_ = _1020_cf2
				var _1021_v72 _dafny.Set
				_ = _1021_v72
				_1021_v72 = _dafny.SetOf(p2)
				_1021_v72 = _1021_v72
				var _1022_v73 T0
				_ = _1022_v73
				var _nw163 *C5 = New_C5_()
				_ = _nw163
				_nw163.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("svx"), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(532))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}(func(_1023_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				}))).Cardinality()), _dafny.IntOfInt64(135), (_this).F8())
				_1022_v73 = _nw163
				var _1024_v74 _dafny.Map
				_ = _1024_v74
				_1024_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_925_v7, _1022_v73)
				var _1025_v75 _dafny.Sequence
				_ = _1025_v75
				_1025_v75 = _dafny.SeqOf((func() T0 {
					if (_1024_v74).Contains((_this).Fm2(p3, globalState)) {
						return (_1024_v74).Get((_this).Fm2(p3, globalState)).(T0)
					}
					return _1022_v73
				})(), _1022_v73)
				var _1026_v76 *C2
				_ = _1026_v76
				var _nw164 *C2 = New_C2_()
				_ = _nw164
				_nw164.Ctor__(_dafny.IntOfUint32((_1025_v75).Cardinality()), Companion_Default___.Fm20(p3, false, p1, globalState))
				_1026_v76 = _nw164
				_926_v8 = _926_v8
				var _1027_v77 _dafny.Array
				_ = _1027_v77
				var _nw165 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
				_ = _nw165
				_1027_v77 = _nw165
				var _1028_v78 _dafny.Array
				_ = _1028_v78
				var _nw166 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
				_ = _nw166
				_1028_v78 = _nw166
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_1027_v77), 0))
				_ = _index186
				(_1027_v77).ArraySet1(_1028_v78, (_index186).Int())
				var _1029_v79 _dafny.Array
				_ = _1029_v79
				var _nw167 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
				_ = _nw167
				_1029_v79 = _nw167
				var _1030_v80 D7
				_ = _1030_v80
				_1030_v80 = Companion_D7_.Create_DC19_(p0, p1, (_this).F7())
				var _1031_v81 _dafny.Map
				_ = _1031_v81
				_1031_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1022_v73).F8()).Minus(p2), ((_1030_v80).Dtor_cf31()).Cmp(p2) <= 0)
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_1027_v77), 0))
				_ = _index187
				var _rhs179 _dafny.Array = _1028_v78
				_ = _rhs179
				var _rhs180 _dafny.Sequence = _925_v7
				_ = _rhs180
				var _rhs181 _dafny.Array = _1029_v79
				_ = _rhs181
				var _rhs182 bool = p3
				_ = _rhs182
				var _rhs183 _dafny.Map = _1031_v81
				_ = _rhs183
				var _lhs135 _dafny.Array = _1027_v77
				_ = _lhs135
				var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_1027_v77), 0))
				_ = _lhs136
				var _lhs137 *GlobalState = globalState
				_ = _lhs137
				(_lhs135).ArraySet1(_rhs179, (_lhs136).Int())
				_925_v7 = _rhs180
				_1029_v79 = _rhs181
				_lhs137.F3 = _rhs182
				_1031_v81 = _rhs183
			}
			r1 = p2
			var _1032_v82 _dafny.Array
			_ = _1032_v82
			var _nw168 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw168
			_1032_v82 = _nw168
			var _1033_v83 _dafny.Array
			_ = _1033_v83
			var _nw169 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
			_ = _nw169
			_1033_v83 = _nw169
			var _1034_v84 _dafny.Map
			_ = _1034_v84
			_1034_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1033_v83)
			var _1035_v85 _dafny.CodePoint
			_ = _1035_v85
			_1035_v85 = _dafny.CodePoint('w')
			var _1036_v86 _dafny.Array
			_ = _1036_v86
			var _nwElement0_35 _dafny.Int = (func() _dafny.Int {
				if (_924_v6).Contains(p0) {
					return (_924_v6).Multiplicity(p0)
				}
				return (_this).F8()
			})()
			_ = _nwElement0_35
			var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(14))
			_ = _nw170
			(_nw170).ArraySet1(_nwElement0_35, 0)
			(_nw170).ArraySet1((_this).F7(), 1)
			(_nw170).ArraySet1(p0, 2)
			(_nw170).ArraySet1((_1034_v84).Cardinality(), 3)
			(_nw170).ArraySet1((_this).F7(), 4)
			(_nw170).ArraySet1((_dafny.Zero).Minus(p0), 5)
			(_nw170).ArraySet1(p2, 6)
			(_nw170).ArraySet1((_this).Fm4(true, true, globalState), 7)
			(_nw170).ArraySet1(p2, 8)
			(_nw170).ArraySet1(Companion_Default___.Fm20(p1, p1, p3, globalState), 9)
			(_nw170).ArraySet1((_this).F7(), 10)
			(_nw170).ArraySet1(_dafny.IntOfInt64(965), 11)
			(_nw170).ArraySet1(((_dafny.MultiSetOf(_1035_v85)).Update(_1035_v85, Companion_Default___.Abs(_dafny.IntOfUint32((_925_v7).Cardinality())))).Cardinality(), 12)
			(_nw170).ArraySet1(p0, 13)
			_1036_v86 = _nw170
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_1032_v82), 0))
			_ = _index188
			(_1032_v82).ArraySet1(_1036_v86, (_index188).Int())
			var _1037_v87 D2
			_ = _1037_v87
			_1037_v87 = Companion_D2_.Create_DC4_(_922_v4)
			var _1038_v88 _dafny.MultiSet
			_ = _1038_v88
			_1038_v88 = _dafny.MultiSetOf(_1037_v87, _1037_v87, _1037_v87)
			var _1039_v89 _dafny.Map
			_ = _1039_v89
			_1039_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), p3)
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_1032_v82), 0))
			_ = _index189
			var _rhs184 _dafny.Int = (func() _dafny.Int {
				if p1 {
					return _dafny.IntOfInt64(-197)
				}
				return Companion_Default___.SafeDivisionInt((_this).F8(), (_1039_v89).Cardinality())
			})()
			_ = _rhs184
			var _rhs185 _dafny.Array = _1036_v86
			_ = _rhs185
			var _rhs186 _dafny.MultiSet = (_dafny.MultiSetOf(_1037_v87)).Union(_1038_v88)
			_ = _rhs186
			var _rhs187 _dafny.Int = ((_this).F8()).Plus((_this).F7())
			_ = _rhs187
			var _lhs138 _dafny.Array = _1032_v82
			_ = _lhs138
			var _lhs139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_1032_v82), 0))
			_ = _lhs139
			r1 = _rhs184
			(_lhs138).ArraySet1(_rhs185, (_lhs139).Int())
			_1038_v88 = _rhs186
			r1 = _rhs187
		}
		if p3 {
			var _1040_v90 _dafny.CodePoint
			_ = _1040_v90
			_1040_v90 = _dafny.CodePoint('h')
			_1040_v90 = _1040_v90
			r1 = (func() _dafny.Int {
				if !(true) {
					return (_this).F7()
				}
				return (func() _dafny.Int {
					if p1 {
						return p0
					}
					return (_this).F7()
				})()
			})()
			var _1041_v91 _dafny.Array
			_ = _1041_v91
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_27
			var _nw171 _dafny.Array
			_ = _nw171
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw171 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Int = func(_1042_i8 _dafny.Int) _dafny.Int {
					return (_1042_i8).Times(_dafny.IntOfInt64(-317))
				}
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw171 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw171).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw171).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1041_v91 = _nw171
			var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1041_v91), 0))
			_ = _index190
			(_1041_v91).ArraySet1(_dafny.IntOfInt64(629), (_index190).Int())
			var _1043_v92 _dafny.MultiSet
			_ = _1043_v92
			_1043_v92 = _dafny.MultiSetOf(p2, p2)
			var _1044_v94 _dafny.Set
			_ = _1044_v94
			_1044_v94 = _dafny.SetOf(Companion_Default___.Fm29(p3, _1043_v92, globalState), (func() _dafny.Set {
				var _coll38 = _dafny.NewBuilder()
				_ = _coll38
				for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(224), _dafny.IntOfInt64(478))); ; {
					_compr_38, _ok42 := _iter42()
					if !_ok42 {
						break
					}
					var _1045_v93 _dafny.Int
					_1045_v93 = interface{}(_compr_38).(_dafny.Int)
					if ((_dafny.IntOfInt64(224)).Cmp(_1045_v93) <= 0) && ((_1045_v93).Cmp(_dafny.IntOfInt64(478)) < 0) {
						_coll38.Add((_1045_v93).Plus(_dafny.IntOfInt64(701)))
					}
				}
				return _coll38.ToSet()
			}()).Intersection(_dafny.SetOf(_dafny.IntOfInt64(-460), (_this).F8(), (_dafny.Zero).Minus(p2))))
			var _1046_v95 _dafny.Sequence
			_ = _1046_v95
			_1046_v95 = _dafny.UnicodeSeqOfUtf8Bytes("ctt")
			var _1047_v96 _dafny.Set
			_ = _1047_v96
			_1047_v96 = _dafny.SetOf(_dafny.IntOfInt64(788))
			var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1041_v91), 0))
			_ = _index191
			var _rhs188 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F7(), _dafny.IntOfUint32((_1046_v95).Cardinality())))
			_ = _rhs188
			var _rhs189 _dafny.Set = ((_1044_v94).Difference(_dafny.SetOf(_1047_v96))).Difference(_1044_v94)
			_ = _rhs189
			var _rhs190 _dafny.Int = p0
			_ = _rhs190
			var _lhs140 _dafny.Array = _1041_v91
			_ = _lhs140
			var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1041_v91), 0))
			_ = _lhs141
			(_lhs140).ArraySet1(_rhs188, (_lhs141).Int())
			_1044_v94 = _rhs189
			r1 = _rhs190
			r1 = ((_this).Fm3(p3, p1, p3, false, globalState)).Times((_dafny.IntOfInt64(-226)).Times((_1041_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1041_v91), 0))).Int()).(_dafny.Int)))
			var _1048_v97 _dafny.Sequence
			_ = _1048_v97
			_1048_v97 = _dafny.SeqOf(p1, (Companion_D5_.Create_DC13_(Companion_Default___.Fm24(p1, (_this).F7(), _1046_v95, globalState), _1040_v90)).Dtor_cf20(), !(p3))
			var _1049_v98 T2
			_ = _1049_v98
			var _nw172 *C2 = New_C2_()
			_ = _nw172
			_nw172.Ctor__((_this).F8(), p2)
			_1049_v98 = _nw172
			var _1050_v99 _dafny.Map
			_ = _1050_v99
			_1050_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1047_v96).Cardinality(), _1049_v98)
			var _1051_v100 _dafny.Array
			_ = _1051_v100
			var _nwElement0_36 bool = p1
			_ = _nwElement0_36
			var _nw173 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(4))
			_ = _nw173
			(_nw173).ArraySet1(_nwElement0_36, 0)
			(_nw173).ArraySet1((_1048_v97).Select((Companion_Default___.SafeIndex((_1050_v99).Cardinality(), _dafny.IntOfUint32((_1048_v97).Cardinality()))).Uint32()).(bool), 1)
			(_nw173).ArraySet1(p3, 2)
			(_nw173).ArraySet1(p3, 3)
			_1051_v100 = _nw173
			var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_1051_v100), 0))
			_ = _index192
			(_1051_v100).ArraySet1(p1, (_index192).Int())
			var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_1051_v100), 0))
			_ = _index193
			(_1051_v100).ArraySet1((_1048_v97).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.IntOfUint32((_1048_v97).Cardinality()))).Uint32()).(bool), (_index193).Int())
		} else {
			var _1052_v101 _dafny.Map
			_ = _1052_v101
			_1052_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
			r1 = (((_1052_v101).Update(p3, p1)).Cardinality()).Times((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ymjehvl")).Cardinality())).Plus((_this).F7()))
			var _1053_v102 _dafny.Sequence
			_ = _1053_v102
			_1053_v102 = _dafny.UnicodeSeqOfUtf8Bytes("uaehms")
			var _1054_v103 _dafny.Sequence
			_ = _1054_v103
			_1054_v103 = _dafny.SeqOf(_1053_v102, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(274))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg64 _dafny.Int) interface{} {
					return coer64(arg64)
				}
			}(func(_1055_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			})), _1053_v102)
			if !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_1053_v102, _1053_v102), _dafny.Companion_Sequence_.Concatenate((_1054_v103).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_1054_v103).Cardinality()))).Uint32()).(_dafny.Sequence), _1053_v102)) {
				r1 = (_this).F8()
				var _1056_v104 _dafny.Sequence
				_ = _1056_v104
				_1056_v104 = _dafny.SeqOf(p1)
				var _1057_v105 _dafny.CodePoint
				_ = _1057_v105
				_1057_v105 = _dafny.CodePoint('i')
				var _1058_v106 _dafny.Sequence
				_ = _1058_v106
				_1058_v106 = _dafny.SeqOf((_this).F8())
				var _1059_v107 _dafny.Array
				_ = _1059_v107
				var _nwElement0_37 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1056_v104, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1056_v104).Cardinality()))).Uint32(), p1)).Cardinality())
				_ = _nwElement0_37
				var _nw174 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(10))
				_ = _nw174
				(_nw174).ArraySet1(_nwElement0_37, 0)
				(_nw174).ArraySet1((Companion_D7_.Create_DC20_(_1057_v105, _dafny.IntOfUint32((_1054_v103).Cardinality()))).Dtor_cf33(), 1)
				(_nw174).ArraySet1((_this).F7(), 2)
				(_nw174).ArraySet1((p2).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), p1)).Cardinality()), 3)
				(_nw174).ArraySet1((func() _dafny.Int {
					if p3 {
						return p2
					}
					return (_this).F7()
				})(), 4)
				(_nw174).ArraySet1((_dafny.IntOfUint32((_1058_v106).Cardinality())).Plus((_dafny.Zero).Minus((_this).F8())), 5)
				(_nw174).ArraySet1(p2, 6)
				(_nw174).ArraySet1((_dafny.IntOfInt64(545)).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-723))), 7)
				(_nw174).ArraySet1(p2, 8)
				(_nw174).ArraySet1((_this).F7(), 9)
				_1059_v107 = _nw174
				var _1060_v108 _dafny.Map
				_ = _1060_v108
				_1060_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1057_v105, p3)
				var _1061_v109 *C5
				_ = _1061_v109
				var _nw175 *C5 = New_C5_()
				_ = _nw175
				_nw175.Ctor__(_1053_v102, (_this).F7(), (_this).F7(), (_dafny.Zero).Minus((_1060_v108).Cardinality()))
				_1061_v109 = _nw175
				var _1062_v110 _dafny.Map
				_ = _1062_v110
				_1062_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _1061_v109)
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_1059_v107), 0))
				_ = _index194
				(_1059_v107).ArraySet1(((_this).F8()).Minus(((_1062_v110).Update((_this).F8(), _1061_v109)).Cardinality()), (_index194).Int())
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_1059_v107), 0))
				_ = _index195
				var _rhs191 _dafny.Int = (_this).F7()
				_ = _rhs191
				var _rhs192 _dafny.Int = (_this).F8()
				_ = _rhs192
				var _lhs142 _dafny.Array = _1059_v107
				_ = _lhs142
				var _lhs143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_1059_v107), 0))
				_ = _lhs143
				(_lhs142).ArraySet1(_rhs191, (_lhs143).Int())
				r1 = _rhs192
				_1057_v105 = _1057_v105
				var _1063_v111 _dafny.Map
				_ = _1063_v111
				_1063_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), (_this).F8())
				var _1064_v113 _dafny.Set
				_ = _1064_v113
				_1064_v113 = _dafny.SetOf(Companion_Default___.Fm5(globalState))
				var _1065_v114 _dafny.Sequence
				_ = _1065_v114
				_1065_v114 = _dafny.SeqOf((_dafny.SetOf(_1057_v105)).Union(func() _dafny.Set {
					var _coll39 = _dafny.NewBuilder()
					_ = _coll39
					for _iter43 := _dafny.Iterate((_1063_v111).Keys().Elements()); ; {
						_compr_39, _ok43 := _iter43()
						if !_ok43 {
							break
						}
						var _1066_v112 _dafny.CodePoint
						_1066_v112 = interface{}(_compr_39).(_dafny.CodePoint)
						if (_1063_v111).Contains(_1066_v112) {
							_coll39.Add(_1066_v112)
						}
					}
					return _coll39.ToSet()
				}()), _1064_v113)
				_1065_v114 = _dafny.Companion_Sequence_.Concatenate(_1065_v114, _dafny.Companion_Sequence_.Concatenate(_1065_v114, _1065_v114))
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(17), _dafny.ArrayLen((_1059_v107), 0))
				_ = _index196
				(_1059_v107).ArraySet1(((_this).F7()).Plus((_this).F7()), (_index196).Int())
			} else {
				var _1067_v115 _dafny.Sequence
				_ = _1067_v115
				_1067_v115 = _dafny.SeqOf(p1)
				var _1068_v116 D2
				_ = _1068_v116
				_1068_v116 = Companion_D2_.Create_DC5_(p1, (_this).F8(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1067_v115, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1067_v115).Cardinality()), _dafny.IntOfUint32((_1067_v115).Cardinality()))).Uint32(), p1)).Cardinality()), (_this).F7(), Companion_Default___.Fm24(p3, _dafny.IntOfInt64(998), _1053_v102, globalState))
				var _1069_v117 _dafny.Set
				_ = _1069_v117
				_1069_v117 = _dafny.SetOf(p3)
				var _1070_v118 _dafny.Array
				_ = _1070_v118
				var _nwElement0_38 bool = !(p3)
				_ = _nwElement0_38
				var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(17))
				_ = _nw176
				(_nw176).ArraySet1(_nwElement0_38, 0)
				(_nw176).ArraySet1(p1, 1)
				(_nw176).ArraySet1(!(!((func() bool {
					if (_1068_v116).Dtor_cf4() {
						return p3
					}
					return (_this).Fm1(_1069_v117, (_this).F7(), _dafny.IntOfInt64(-719), (_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), globalState)
				})())), 2)
				(_nw176).ArraySet1(p3, 3)
				(_nw176).ArraySet1(p3, 4)
				(_nw176).ArraySet1(p3, 5)
				(_nw176).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1053_v102, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(279))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}(func(_1071_i10 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('b')
				}))), 6)
				(_nw176).ArraySet1(p1, 7)
				(_nw176).ArraySet1(p1, 8)
				(_nw176).ArraySet1(!(p3) || (true), 9)
				(_nw176).ArraySet1(p3, 10)
				(_nw176).ArraySet1(p3, 11)
				(_nw176).ArraySet1(p1, 12)
				(_nw176).ArraySet1(!(p1) || (true), 13)
				(_nw176).ArraySet1(p3, 14)
				(_nw176).ArraySet1(p3, 15)
				(_nw176).ArraySet1(p3, 16)
				_1070_v118 = _nw176
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_1070_v118), 0))
				_ = _index197
				(_1070_v118).ArraySet1(true, (_index197).Int())
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_1070_v118), 0))
				_ = _index198
				(_1070_v118).ArraySet1(((_this).F8()).Cmp(((_this).F8()).Plus(p2)) != 0, (_index198).Int())
				var _1072_v119 _dafny.Set
				_ = _1072_v119
				_1072_v119 = _dafny.SetOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_1052_v101).Cardinality())), p2)
				r1 = (p0).Times((_1072_v119).Cardinality())
				var _1073_v120 _dafny.Map
				_ = _1073_v120
				_1073_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F7())
				var _1074_v121 *C1
				_ = _1074_v121
				var _nw177 *C1 = New_C1_()
				_ = _nw177
				_nw177.Ctor__((func() _dafny.Int {
					if (_1073_v120).Contains(p0) {
						return (_1073_v120).Get(p0).(_dafny.Int)
					}
					return (_this).F7()
				})(), (func() _dafny.Int {
					if !(p3) {
						return (_this).F7()
					}
					return (_this).Fm3(p1, true, (_1070_v118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_1070_v118), 0))).Int()).(bool), p1, globalState)
				})())
				_1074_v121 = _nw177
				var _1075_v122 D5
				_ = _1075_v122
				_1075_v122 = Companion_D5_.Create_DC13_(p1, _dafny.CodePoint('q'))
				var _1076_v123 _dafny.Map
				_ = _1076_v123
				_1076_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(Companion_Default___.Fm16(p3, globalState), p0), _dafny.SeqOf(_1075_v122, _1075_v122, Companion_Default___.Fm32(globalState), Companion_Default___.Fm32(globalState)))
				var _1077_v124 _dafny.MultiSet
				_ = _1077_v124
				_1077_v124 = _dafny.MultiSetOf(p3)
				var _1078_v125 _dafny.MultiSet
				_ = _1078_v125
				_1078_v125 = _dafny.MultiSetOf((func() _dafny.Int {
					if (_1077_v124).Contains(p3) {
						return (_1077_v124).Multiplicity(p3)
					}
					return (_this).F7()
				})())
				var _1079_v126 _dafny.Sequence
				_ = _1079_v126
				_1079_v126 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt((_1078_v125).Cardinality(), _dafny.IntOfInt64(827)), (_dafny.Zero).Minus((_dafny.IntOfInt64(-451)).Plus(p0)), p0, Companion_Default___.SafeModuloInt((_this).F7(), p2))
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_1070_v118), 0))
				_ = _index199
				var _rhs193 bool = (func() bool {
					if p1 {
						return false
					}
					return p1
				})()
				_ = _rhs193
				var _rhs194 _dafny.Sequence = _1053_v102
				_ = _rhs194
				var _rhs195 _dafny.Map = _1076_v123
				_ = _rhs195
				var _rhs196 bool = (_1070_v118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_1070_v118), 0))).Int()).(bool)
				_ = _rhs196
				var _rhs197 _dafny.Sequence = _1079_v126
				_ = _rhs197
				var _lhs144 *GlobalState = globalState
				_ = _lhs144
				var _lhs145 _dafny.Array = _1070_v118
				_ = _lhs145
				var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_1070_v118), 0))
				_ = _lhs146
				_lhs144.F4 = _rhs193
				_1053_v102 = _rhs194
				_1076_v123 = _rhs195
				(_lhs145).ArraySet1(_rhs196, (_lhs146).Int())
				_1079_v126 = _rhs197
				r1 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1053_v102, _1053_v102)).Cardinality())).Plus((func() _dafny.Int {
					if (_1078_v125).Contains(Companion_Default___.Fm20(p3, (_1070_v118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_1070_v118), 0))).Int()).(bool), p1, globalState)) {
						return (_1078_v125).Multiplicity(Companion_Default___.Fm20(p3, (_1070_v118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.ArrayLen((_1070_v118), 0))).Int()).(bool), p1, globalState))
					}
					return _dafny.IntOfUint32((_1067_v115).Cardinality())
				})())
			}
			var _1080_v127 *C3
			_ = _1080_v127
			var _nw178 *C3 = New_C3_()
			_ = _nw178
			_nw178.Ctor__(!(p3), (_this).F7(), p2)
			_1080_v127 = _nw178
			var _1081_v128 _dafny.Array
			_ = _1081_v128
			var _nw179 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
			_ = _nw179
			_1081_v128 = _nw179
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1081_v128), 0))
			_ = _index200
			(_1081_v128).ArraySet1(_1053_v102, (_index200).Int())
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1081_v128), 0))
			_ = _index201
			(_1081_v128).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1053_v102, _1053_v102), (_index201).Int())
			var _1082_v129 _dafny.Set
			_ = _1082_v129
			_1082_v129 = _dafny.SetOf((_1080_v127).F12())
			var _1083_v130 _dafny.Sequence
			_ = _1083_v130
			_1083_v130 = _dafny.SeqOf(p1, p3)
			var _1084_v131 _dafny.Array
			_ = _1084_v131
			var _nwElement0_39 _dafny.Int = _dafny.IntOfUint32((_1053_v102).Cardinality())
			_ = _nwElement0_39
			var _nw180 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(18))
			_ = _nw180
			(_nw180).ArraySet1(_nwElement0_39, 0)
			(_nw180).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F8(), p2)), 1)
			(_nw180).ArraySet1(p0, 2)
			(_nw180).ArraySet1(p2, 3)
			(_nw180).ArraySet1(p2, 4)
			(_nw180).ArraySet1((_dafny.Zero).Minus((_this).F8()), 5)
			(_nw180).ArraySet1((_this).F7(), 6)
			(_nw180).ArraySet1((_this).F8(), 7)
			(_nw180).ArraySet1((_1082_v129).Cardinality(), 8)
			(_nw180).ArraySet1(_dafny.IntOfUint32((_1083_v130).Cardinality()), 9)
			(_nw180).ArraySet1(_dafny.IntOfInt64(-703), 10)
			(_nw180).ArraySet1(Companion_Default___.SafeModuloInt(p2, (_1082_v129).Cardinality()), 11)
			(_nw180).ArraySet1(_dafny.IntOfUint32(((_1080_v127).Fm2(true, globalState)).Cardinality()), 12)
			(_nw180).ArraySet1(p0, 13)
			(_nw180).ArraySet1(((_this).F7()).Plus(p2), 14)
			(_nw180).ArraySet1((_this).F8(), 15)
			(_nw180).ArraySet1((_dafny.IntOfInt64(125)).Minus((_this).F8()), 16)
			(_nw180).ArraySet1(p2, 17)
			_1084_v131 = _nw180
			_1084_v131 = _1084_v131
		}
		if (p0).Cmp(p2) >= 0 {
			var _1085_v132 _dafny.Sequence
			_ = _1085_v132
			_1085_v132 = _dafny.SeqOf(p0)
			var _1086_v133 _dafny.Sequence
			_ = _1086_v133
			_1086_v133 = _dafny.SeqOf(_1085_v132, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(774))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg66 _dafny.Int) interface{} {
					return coer66(arg66)
				}
			}(func(_1087_i11 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(431)
			})))
			var _1088_v134 _dafny.Map
			_ = _1088_v134
			_1088_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _dafny.IntOfUint32((_1086_v133).Cardinality()))
			_1088_v134 = (_1088_v134).Update(_dafny.IntOfInt64(-267), ((_this).F8()).Plus(p0))
			var _1089_v135 _dafny.MultiSet
			_ = _1089_v135
			_1089_v135 = _dafny.MultiSetOf(p1, p3, p3, !(p3), !(p1))
			var _1090_v136 _dafny.Map
			_ = _1090_v136
			_1090_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1089_v135, p1)
			_1090_v136 = (_1090_v136).Update(_1089_v135, ((_this).F8()).Cmp((_this).F7()) < 0)
			var _1091_v137 _dafny.Map
			_ = _1091_v137
			_1091_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), p3)
			_1091_v137 = (_1091_v137).Update((_dafny.Zero).Minus((_this).F7()), true)
			(globalState).F4 = p3
			(globalState).F3 = (p0).Cmp((_this).F7()) > 0
		} else {
			var _1092_v138 _dafny.Sequence
			_ = _1092_v138
			_1092_v138 = _dafny.SeqOf(p2)
			var _1093_v139 _dafny.Sequence
			_ = _1093_v139
			_1093_v139 = _dafny.UnicodeSeqOfUtf8Bytes("fxgvl")
			var _1094_v140 _dafny.Map
			_ = _1094_v140
			_1094_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1092_v138).Cardinality()), _dafny.IntOfUint32((_1093_v139).Cardinality()))
			var _1095_v141 _dafny.Set
			_ = _1095_v141
			_1095_v141 = _dafny.SetOf(p1, p1, p1, p3, p3)
			var _1096_v142 _dafny.MultiSet
			_ = _1096_v142
			_1096_v142 = _dafny.MultiSetOf(_1095_v141)
			var _1097_v143 _dafny.Map
			_ = _1097_v143
			_1097_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)
			var _1098_v144 _dafny.MultiSet
			_ = _1098_v144
			_1098_v144 = _dafny.MultiSetOf(p2)
			var _1099_v145 _dafny.MultiSet
			_ = _1099_v145
			_1099_v145 = _dafny.MultiSetOf(_1098_v144, _1098_v144)
			var _1100_v146 _dafny.Array
			_ = _1100_v146
			var _nwElement0_40 _dafny.Int = (_this).F7()
			_ = _nwElement0_40
			var _nw181 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(9))
			_ = _nw181
			(_nw181).ArraySet1(_nwElement0_40, 0)
			(_nw181).ArraySet1((_dafny.Zero).Minus((_this).F8()), 1)
			(_nw181).ArraySet1((_this).F7(), 2)
			(_nw181).ArraySet1(_dafny.IntOfInt64(34), 3)
			(_nw181).ArraySet1((func() _dafny.Int {
				if (_1099_v145).Contains(_1098_v144) {
					return (_1099_v145).Multiplicity(_1098_v144)
				}
				return p0
			})(), 4)
			(_nw181).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !(p3))).Cardinality(), 5)
			(_nw181).ArraySet1(p0, 6)
			(_nw181).ArraySet1(_dafny.IntOfUint32((_1093_v139).Cardinality()), 7)
			(_nw181).ArraySet1(p2, 8)
			_1100_v146 = _nw181
			var _1101_v147 _dafny.Map
			_ = _1101_v147
			_1101_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1100_v146, p3)
			var _1102_v148 _dafny.Array
			_ = _1102_v148
			var _nwElement0_41 _dafny.Int = p2
			_ = _nwElement0_41
			var _nw182 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(26))
			_ = _nw182
			(_nw182).ArraySet1(_nwElement0_41, 0)
			(_nw182).ArraySet1((func() _dafny.Int {
				if (_1094_v140).Contains(p0) {
					return (_1094_v140).Get(p0).(_dafny.Int)
				}
				return p0
			})(), 1)
			(_nw182).ArraySet1((_this).F8(), 2)
			(_nw182).ArraySet1((_this).F8(), 3)
			(_nw182).ArraySet1(_dafny.IntOfInt64(201), 4)
			(_nw182).ArraySet1(p2, 5)
			(_nw182).ArraySet1((func() _dafny.Int {
				if (_1096_v142).Contains(_1095_v141) {
					return (_1096_v142).Multiplicity(_1095_v141)
				}
				return (_1097_v143).Cardinality()
			})(), 6)
			(_nw182).ArraySet1(p2, 7)
			(_nw182).ArraySet1(_dafny.IntOfInt64(-434), 8)
			(_nw182).ArraySet1((_1092_v138).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1092_v138).Cardinality()))).Uint32()).(_dafny.Int), 9)
			(_nw182).ArraySet1((_this).Fm3(p3, p1, p1, (func() bool {
				if (_1097_v143).Contains(p3) {
					return (_1097_v143).Get(p3).(bool)
				}
				return p1
			})(), globalState), 10)
			(_nw182).ArraySet1(p2, 11)
			(_nw182).ArraySet1((_this).F7(), 12)
			(_nw182).ArraySet1(p2, 13)
			(_nw182).ArraySet1(_dafny.IntOfUint32((_1092_v138).Cardinality()), 14)
			(_nw182).ArraySet1((_this).F7(), 15)
			(_nw182).ArraySet1((_this).F7(), 16)
			(_nw182).ArraySet1(_dafny.IntOfUint32((_1092_v138).Cardinality()), 17)
			(_nw182).ArraySet1((_dafny.Zero).Minus((_1101_v147).Cardinality()), 18)
			(_nw182).ArraySet1(p0, 19)
			(_nw182).ArraySet1(p0, 20)
			(_nw182).ArraySet1(p2, 21)
			(_nw182).ArraySet1(p2, 22)
			(_nw182).ArraySet1(Companion_Default___.Fm20(p3, true, false, globalState), 23)
			(_nw182).ArraySet1(p2, 24)
			(_nw182).ArraySet1(p0, 25)
			_1102_v148 = _nw182
			var _1103_v149 *C0
			_ = _1103_v149
			var _nw183 *C0 = New_C0_()
			_ = _nw183
			_nw183.Ctor__(_1100_v146, (_this).F7(), (_this).F8())
			_1103_v149 = _nw183
			var _1104_v150 _dafny.Sequence
			_ = _1104_v150
			_1104_v150 = _dafny.SeqOf(p1)
			r0 = _dafny.Companion_Sequence_.Concatenate(_1104_v150, _1104_v150)
			var _1105_v151 _dafny.Map
			_ = _1105_v151
			_1105_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F7()).Cmp((_this).F7()) < 0, (_this).F7())
			_1105_v151 = (_1105_v151).Update(p3, (p2).Minus((_this).F8()))
			r1 = p2
			var _arr14 _dafny.Array = _1103_v149.F13
			_ = _arr14
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_1103_v149.F13), 0))
			_ = _index202
			(_arr14).ArraySet1(p2, (_index202).Int())
			var _arr15 _dafny.Array = _1103_v149.F13
			_ = _arr15
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_1103_v149.F13), 0))
			_ = _index203
			(_arr15).ArraySet1(p2, (_index203).Int())
		}
		if p1 {
			var _1106_v152 _dafny.CodePoint
			_ = _1106_v152
			_1106_v152 = _dafny.CodePoint('e')
			var _1107_v153 _dafny.Sequence
			_ = _1107_v153
			_1107_v153 = _dafny.UnicodeSeqOfUtf8Bytes("vlqsggw")
			var _1108_v154 _dafny.Sequence
			_ = _1108_v154
			_1108_v154 = _dafny.SeqOf((_1107_v153).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1107_v153).Cardinality()))).Uint32()).(_dafny.CodePoint), _1106_v152, Companion_Default___.Fm5(globalState))
			_1106_v152 = (func() _dafny.CodePoint {
				if p1 {
					return _1106_v152
				}
				return (_1108_v154).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1108_v154).Cardinality()))).Uint32()).(_dafny.CodePoint)
			})()
			var _1109_v155 _dafny.Array
			_ = _1109_v155
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_28
			var _nw184 _dafny.Array
			_ = _nw184
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw184 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) _dafny.Int = func(_1110_i12 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1110_i12, (_this).F8())
				}
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw184 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw184).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw184).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1109_v155 = _nw184
			var _1111_v156 _dafny.Map
			_ = _1111_v156
			_1111_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1109_v155, p2)
			_1111_v156 = (_1111_v156).Update(_1109_v155, (func() _dafny.Int {
				if !(p3) {
					return (_this).F8()
				}
				return p0
			})())
			r1 = p0
			if p3 {
				var _1112_v157 *C2
				_ = _1112_v157
				var _nw185 *C2 = New_C2_()
				_ = _nw185
				_nw185.Ctor__((_dafny.Zero).Minus(_dafny.IntOfUint32((_1108_v154).Cardinality())), (_this).F8())
				_1112_v157 = _nw185
				(globalState).F3 = p1
				_1109_v155 = _1109_v155
				var _1113_v158 _dafny.Sequence
				_ = _1113_v158
				_1113_v158 = _dafny.SeqOf(p1)
				var _1114_v159 _dafny.Sequence
				_ = _1114_v159
				_1114_v159 = _dafny.SeqOf(_1113_v158)
				var _1115_v160 _dafny.Sequence
				_ = _1115_v160
				_1115_v160 = _dafny.SeqOf((_this).F8(), p0, _dafny.IntOfUint32(((_1114_v159).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1114_v159).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
				var _1116_v161 _dafny.MultiSet
				_ = _1116_v161
				_1116_v161 = _dafny.MultiSetOf(p0, (_this).F8(), (_dafny.Zero).Minus(p0))
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_1109_v155), 0))
				_ = _index204
				(_1109_v155).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ikupdc")).Cardinality())).Times((func() _dafny.Int {
					if (_1116_v161).Contains((_dafny.Zero).Minus(p2)) {
						return (_1116_v161).Multiplicity((_dafny.Zero).Minus(p2))
					}
					return _dafny.IntOfUint32((_1115_v160).Cardinality())
				})())), (_index204).Int())
				var _1117_v162 _dafny.Map
				_ = _1117_v162
				_1117_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1113_v158).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_1115_v160, _1115_v160))
				var _1118_v163 _dafny.Map
				_ = _1118_v163
				_1118_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p3)
				var _1119_v164 _dafny.MultiSet
				_ = _1119_v164
				_1119_v164 = _dafny.MultiSetOf((_1114_v159).Select((Companion_Default___.SafeIndex((_1118_v163).Cardinality(), _dafny.IntOfUint32((_1114_v159).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_1109_v155), 0))
				_ = _index205
				var _rhs198 _dafny.Int = (_this).F7()
				_ = _rhs198
				var _rhs199 _dafny.Sequence = (func() _dafny.Sequence {
					if (_1117_v162).Contains(p0) {
						return (_1117_v162).Get(p0).(_dafny.Sequence)
					}
					return _1115_v160
				})()
				_ = _rhs199
				var _rhs200 _dafny.Int = (((_dafny.Zero).Minus((_1119_v164).Cardinality())).Plus(p2)).Minus((_this).F7())
				_ = _rhs200
				var _lhs147 _dafny.Array = _1109_v155
				_ = _lhs147
				var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_1109_v155), 0))
				_ = _lhs148
				r1 = _rhs198
				_1115_v160 = _rhs199
				(_lhs147).ArraySet1(_rhs200, (_lhs148).Int())
				(globalState).F4 = !(p1)
			} else {
				var _1120_v165 _dafny.Set
				_ = _1120_v165
				_1120_v165 = _dafny.SetOf(p3, p1, p1)
				var _1121_v166 _dafny.Set
				_ = _1121_v166
				_1121_v166 = _dafny.SetOf(p3, p1, (_this).Fm1(_1120_v165, p2, _dafny.IntOfInt64(-90), _dafny.IntOfUint32((_1107_v153).Cardinality()), globalState), !(p1), p3)
				var _1122_v167 D1
				_ = _1122_v167
				_1122_v167 = Companion_D1_.Create_DC2_()
				var _1123_v168 D1
				_ = _1123_v168
				_1123_v168 = Companion_D1_.Create_DC3_(_1122_v167)
				var _1124_v169 D1
				_ = _1124_v169
				_1124_v169 = Companion_D1_.Create_DC3_(_1123_v168)
				var _1125_v170 D1
				_ = _1125_v170
				_1125_v170 = Companion_D1_.Create_DC3_(_1124_v169)
				var _1126_v171 _dafny.MultiSet
				_ = _1126_v171
				_1126_v171 = _dafny.MultiSetOf(Companion_D1_.Create_DC3_(_1124_v169), _1125_v170)
				var _1127_v172 _dafny.Array
				_ = _1127_v172
				var _nwElement0_42 bool = !(p1)
				_ = _nwElement0_42
				var _nw186 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(15))
				_ = _nw186
				(_nw186).ArraySet1(_nwElement0_42, 0)
				(_nw186).ArraySet1(p3, 1)
				(_nw186).ArraySet1(p1, 2)
				(_nw186).ArraySet1(p3, 3)
				(_nw186).ArraySet1((_this).Fm1(_1120_v165, p2, (func() _dafny.Int {
					if (_1126_v171).Contains(_1125_v170) {
						return (_1126_v171).Multiplicity(_1125_v170)
					}
					return p0
				})(), (_this).F7(), globalState), 4)
				(_nw186).ArraySet1(p1, 5)
				(_nw186).ArraySet1(p3, 6)
				(_nw186).ArraySet1(p1, 7)
				(_nw186).ArraySet1(p3, 8)
				(_nw186).ArraySet1(p3, 9)
				(_nw186).ArraySet1(p3, 10)
				(_nw186).ArraySet1(p3, 11)
				(_nw186).ArraySet1(p3, 12)
				(_nw186).ArraySet1(p1, 13)
				(_nw186).ArraySet1(p1, 14)
				_1127_v172 = _nw186
				var _1128_v173 D6
				_ = _1128_v173
				_1128_v173 = Companion_D6_.Create_DC15_(_1127_v172)
				var _rhs201 D6 = _1128_v173
				_ = _rhs201
				var _rhs202 _dafny.Array = (func() _dafny.Array {
					if p1 {
						return _1127_v172
					}
					return _1127_v172
				})()
				_ = _rhs202
				var _rhs203 _dafny.Int = (_this).Fm4(p3, p3, globalState)
				_ = _rhs203
				var _rhs204 _dafny.Int = (_this).F7()
				_ = _rhs204
				_1128_v173 = _rhs201
				_1127_v172 = _rhs202
				r1 = _rhs203
				r1 = _rhs204
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_1127_v172), 0))
				_ = _index206
				(_1127_v172).ArraySet1(false, (_index206).Int())
				var _1129_v174 bool
				_ = _1129_v174
				_1129_v174 = p3
				var _1130_v175 _dafny.Array
				_ = _1130_v175
				var _nwElement0_43 bool = p1
				_ = _nwElement0_43
				var _nw187 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(16))
				_ = _nw187
				(_nw187).ArraySet1(_nwElement0_43, 0)
				(_nw187).ArraySet1(_1129_v174, 1)
				(_nw187).ArraySet1(p1, 2)
				(_nw187).ArraySet1(p1, 3)
				(_nw187).ArraySet1(_1129_v174, 4)
				(_nw187).ArraySet1(_1129_v174, 5)
				(_nw187).ArraySet1(_1129_v174, 6)
				(_nw187).ArraySet1(Companion_Default___.Fm33(globalState), 7)
				(_nw187).ArraySet1(Companion_Default___.Fm33(globalState), 8)
				(_nw187).ArraySet1(_1129_v174, 9)
				(_nw187).ArraySet1(_1129_v174, 10)
				(_nw187).ArraySet1(_1129_v174, 11)
				(_nw187).ArraySet1(_1129_v174, 12)
				(_nw187).ArraySet1(_1129_v174, 13)
				(_nw187).ArraySet1(_1129_v174, 14)
				(_nw187).ArraySet1(_1129_v174, 15)
				_1130_v175 = _nw187
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_1127_v172), 0))
				_ = _index207
				var _rhs205 bool = !(p1) || (((_this).Fm1(_1121_v166, (_this).F8(), (_this).F8(), (_this).F8(), globalState)) || (p3))
				_ = _rhs205
				var _rhs206 _dafny.Array = _1130_v175
				_ = _rhs206
				var _lhs149 _dafny.Array = _1127_v172
				_ = _lhs149
				var _lhs150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_1127_v172), 0))
				_ = _lhs150
				(_lhs149).ArraySet1(_rhs205, (_lhs150).Int())
				_1130_v175 = _rhs206
				var _1131_v176 _dafny.Map
				_ = _1131_v176
				_1131_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)
				var _1132_v177 _dafny.Sequence
				_ = _1132_v177
				_1132_v177 = _dafny.SeqOf(p3)
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_1127_v172), 0))
				_ = _index208
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_1127_v172), 0))
				_ = _index209
				var _rhs207 bool = ((func() _dafny.Int {
					if (_1131_v176).Contains(p3) {
						return (_1131_v176).Get(p3).(_dafny.Int)
					}
					return p0
				})()).Cmp(Companion_Default___.SafeDivisionInt(p2, (_this).F8())) >= 0
				_ = _rhs207
				var _rhs208 bool = (_1132_v177).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1132_v177).Cardinality()))).Uint32()).(bool)
				_ = _rhs208
				var _lhs151 _dafny.Array = _1127_v172
				_ = _lhs151
				var _lhs152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_1127_v172), 0))
				_ = _lhs152
				var _lhs153 _dafny.Array = _1127_v172
				_ = _lhs153
				var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_1127_v172), 0))
				_ = _lhs154
				(_lhs151).ArraySet1(_rhs207, (_lhs152).Int())
				(_lhs153).ArraySet1(_rhs208, (_lhs154).Int())
				var _1133_v178 _dafny.Sequence
				_ = _1133_v178
				_1133_v178 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_1108_v154, (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1108_v154).Cardinality()))).Uint32(), _1106_v152), _1108_v154, _dafny.UnicodeSeqOfUtf8Bytes("bp"), _1107_v153, _1107_v153)
				var _1134_v179 _dafny.Map
				_ = _1134_v179
				_1134_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_1107_v153, (_1133_v178).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1133_v178).Cardinality()))).Uint32()).(_dafny.Sequence)), _1127_v172)
				var _1135_v180 _dafny.MultiSet
				_ = _1135_v180
				_1135_v180 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_1108_v154, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1108_v154).Cardinality()))).Uint32(), _1106_v152))
				var _1136_v181 _dafny.Set
				_ = _1136_v181
				_1136_v181 = _dafny.SetOf((func() _dafny.Array {
					if (_1134_v179).Contains(_1135_v180) {
						return (_1134_v179).Get(_1135_v180).(_dafny.Array)
					}
					return _1127_v172
				})())
				var _1137_v182 D5
				_ = _1137_v182
				_1137_v182 = Companion_D5_.Create_DC11_(_1131_v176)
				var _rhs209 _dafny.Set = (_1136_v181).Union(_1136_v181)
				_ = _rhs209
				var _rhs210 D5 = _1137_v182
				_ = _rhs210
				var _rhs211 bool = (_1132_v177).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(205))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}((func(_1138_v152 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1139_i13 _dafny.Int) _dafny.CodePoint {
						return _1138_v152
					}
				})(_1106_v152))), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(205))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg68 _dafny.Int) interface{} {
						return coer68(arg68)
					}
				}((func(_1140_v152 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1141_i13 _dafny.Int) _dafny.CodePoint {
						return _1140_v152
					}
				})(_1106_v152)))).Cardinality()))).Uint32(), _dafny.CodePoint('o'))).Cardinality()), _dafny.IntOfUint32((_1132_v177).Cardinality()))).Uint32()).(bool)
				_ = _rhs211
				var _rhs212 _dafny.Int = p2
				_ = _rhs212
				var _lhs155 *GlobalState = globalState
				_ = _lhs155
				_1136_v181 = _rhs209
				_1137_v182 = _rhs210
				_lhs155.F4 = _rhs211
				r1 = _rhs212
				_1132_v177 = _1132_v177
			}
			r1 = p2
		} else {
			(globalState).F4 = p3
			(globalState).F3 = (((_this).F8()).Cmp(_dafny.IntOfInt64(-20)) > 0) == (p1)
			r1 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(193), p0)
			var _1142_v183 _dafny.Array
			_ = _1142_v183
			var _nw188 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw188
			_1142_v183 = _nw188
			var _1143_v184 D6
			_ = _1143_v184
			_1143_v184 = Companion_D6_.Create_DC15_(_1142_v183)
			_1142_v183 = (_1143_v184).Dtor_cf23()
			var _1144_v185 _dafny.Sequence
			_ = _1144_v185
			_1144_v185 = _dafny.SeqOf(p1, p3, true)
			var _1145_v186 D6
			_ = _1145_v186
			_1145_v186 = Companion_D6_.Create_DC16_((Companion_Default___.Fm35(p2, globalState)).Dtor_cf24(), _1144_v185)
			var _1146_v187 _dafny.Map
			_ = _1146_v187
			_1146_v187 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(p1), p0)
			var _1147_v188 _dafny.Set
			_ = _1147_v188
			_1147_v188 = _dafny.SetOf(Companion_D6_.Create_DC16_(_1146_v187, _1144_v185))
			var _1148_v189 _dafny.Sequence
			_ = _1148_v189
			_1148_v189 = _dafny.SeqOf(Companion_Default___.Fm34(p1, p3, globalState), (_dafny.SetOf(_1145_v186, _1145_v186)).Intersection(_dafny.SetOf(_1145_v186, _1145_v186)), _1147_v188, _1147_v188, (_1147_v188).Union(_1147_v188))
			_1148_v189 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(705))).Uint32(), func(coer69 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
				return func(arg69 _dafny.Int) interface{} {
					return coer69(arg69)
				}
			}((func(_1149_v188 _dafny.Set) func(_dafny.Int) _dafny.Set {
				return func(_1150_i14 _dafny.Int) _dafny.Set {
					return _1149_v188
				}
			})(_1147_v188)))
		}
		var _1151_v190 _dafny.Array
		_ = _1151_v190
		var _nw189 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
		_ = _nw189
		_1151_v190 = _nw189
		var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1151_v190), 0))
		_ = _index210
		(_1151_v190).ArraySet1(p3, (_index210).Int())
		var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1151_v190), 0))
		_ = _index211
		(_1151_v190).ArraySet1(p3, (_index211).Int())
		r1 = p2
		var _1152_v191 _dafny.Sequence
		_ = _1152_v191
		_1152_v191 = _dafny.SeqOf(p3)
		r0 = _1152_v191
		var _1153_v192 D5
		_ = _1153_v192
		_1153_v192 = Companion_D5_.Create_DC11_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1151_v190).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1151_v190), 0))).Int()).(bool), p2)).Update(p1, (_this).F8()))
		var _pat_let_tv15 = p0
		_ = _pat_let_tv15
		var _pat_let_tv16 = _1151_v190
		_ = _pat_let_tv16
		var _pat_let_tv17 = _1151_v190
		_ = _pat_let_tv17
		var _pat_let_tv18 = p0
		_ = _pat_let_tv18
		var _pat_let_tv19 = p3
		_ = _pat_let_tv19
		r1 = (_dafny.Zero).Minus(func(_source23 D5) _dafny.Int {
			if _source23.Is_DC12() {
				var _1154___mcc_h10 bool = _source23.Get_().(D5_DC12).Cf19
				_ = _1154___mcc_h10
				var _1155_cf19 bool = _1154___mcc_h10
				_ = _1155_cf19
				return (_dafny.Zero).Minus(_pat_let_tv15)
			} else if _source23.Is_DC13() {
				var _1156___mcc_h11 bool = _source23.Get_().(D5_DC13).Cf20
				_ = _1156___mcc_h11
				var _1157___mcc_h12 _dafny.CodePoint = _source23.Get_().(D5_DC13).Cf21
				_ = _1157___mcc_h12
				var _1158_cf21 _dafny.CodePoint = _1157___mcc_h12
				_ = _1158_cf21
				var _1159_cf20 bool = _1156___mcc_h11
				_ = _1159_cf20
				return _dafny.IntOfInt64(862)
			} else if _source23.Is_DC11() {
				var _1160___mcc_h13 _dafny.Map = _source23.Get_().(D5_DC11).Cf18
				_ = _1160___mcc_h13
				var _1161_cf18 _dafny.Map = _1160___mcc_h13
				_ = _1161_cf18
				return (_this).F8()
			} else {
				var _1162___mcc_h14 D5 = _source23.Get_().(D5_DC14).Cf22
				_ = _1162___mcc_h14
				var _1163_cf22 D5 = _1162___mcc_h14
				_ = _1163_cf22
				return Companion_Default___.SafeModuloInt((Companion_D2_.Create_DC5_((Companion_D3_.Create_DC7_((_pat_let_tv17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_pat_let_tv16), 0))).Int()).(bool))).Dtor_cf10(), (_this).F7(), (_this).F8(), _pat_let_tv18, (_pat_let_tv19))).Dtor_cf7(), _dafny.IntOfInt64(-174))
			}
		}(_1153_v192))
		r2 = !(p3)
		return r0, r1, r2
	}
}
func (_this *C6) M3(p0 bool, p1 _dafny.Int, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _1164_v0 *C2
		_ = _1164_v0
		var _nw190 *C2 = New_C2_()
		_ = _nw190
		_nw190.Ctor__((_this).F8(), _dafny.IntOfInt64(687))
		_1164_v0 = _nw190
		var _1165_v1 _dafny.Set
		_ = _1165_v1
		_1165_v1 = _dafny.SetOf(_1164_v0, _1164_v0, _1164_v0, _1164_v0, _1164_v0)
		var _1166_v2 _dafny.Map
		_ = _1166_v2
		_1166_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_1165_v1).Intersection(_1165_v1))
		_1165_v1 = (func() _dafny.Set {
			if (_1166_v2).Contains(p1) {
				return (_1166_v2).Get(p1).(_dafny.Set)
			}
			return _1165_v1
		})()
		var _1167_v3 _dafny.Sequence
		_ = _1167_v3
		_1167_v3 = _dafny.SeqOf((_this).F8())
		r0 = (_1167_v3).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1167_v3).Cardinality()))).Uint32()).(_dafny.Int)
		var _1168_v4 T1
		_ = _1168_v4
		var _nw191 *C0 = New_C0_()
		_ = _nw191
		_nw191.Ctor__(p2, (_this).F8(), (Companion_Default___.Fm20(p0, p0, p0, globalState)).Plus((_this).F8()))
		_1168_v4 = _nw191
		_1168_v4 = _1168_v4
		var _1169_v5 D4
		_ = _1169_v5
		_1169_v5 = Companion_D4_.Create_DC10_((_dafny.Zero).Minus((_this).F7()), ((_this).F8()).Cmp(p3) > 0)
		var _1170_v6 _dafny.Array
		_ = _1170_v6
		var _len0_29 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_29
		var _nw192 _dafny.Array
		_ = _nw192
		if _len0_29.Cmp(_dafny.Zero) == 0 {
			_nw192 = _dafny.NewArray(_len0_29)
		} else {
			var _init29 func(_dafny.Int) _dafny.MultiSet = (func(_1171_p0 bool) func(_dafny.Int) _dafny.MultiSet {
				return func(_1172_i0 _dafny.Int) _dafny.MultiSet {
					return (_dafny.MultiSetOf(_1171_p0)).Difference(_dafny.MultiSetOf(_1171_p0, _1171_p0, _1171_p0, _1171_p0, _1171_p0))
				}
			})(p0)
			_ = _init29
			var _element0_29 = _init29(_dafny.Zero)
			_ = _element0_29
			_nw192 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
			(_nw192).ArraySet1(_element0_29, 0)
			var _nativeLen0_29 = (_len0_29).Int()
			_ = _nativeLen0_29
			for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
				(_nw192).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
			}
		}
		_1170_v6 = _nw192
		var _rhs213 D4 = _1169_v5
		_ = _rhs213
		var _rhs214 _dafny.Array = _1170_v6
		_ = _rhs214
		_1169_v5 = _rhs213
		_1170_v6 = _rhs214
		var _hi8 _dafny.Int = _dafny.IntOfInt64(-776)
		_ = _hi8
		for _1173_i1 := _dafny.IntOfInt64(307); _1173_i1.Cmp(_hi8) < 0; _1173_i1 = _1173_i1.Plus(_dafny.One) {
			var _1174_v7 D7
			_ = _1174_v7
			_1174_v7 = Companion_D7_.Create_DC19_(p1, p0, (_1168_v4).F7())
			var _1175_v8 _dafny.Array
			_ = _1175_v8
			var _len0_30 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_30
			var _nw193 _dafny.Array
			_ = _nw193
			if _len0_30.Cmp(_dafny.Zero) == 0 {
				_nw193 = _dafny.NewArray(_len0_30)
			} else {
				var _init30 func(_dafny.Int) bool = (func(_1176_p0 bool) func(_dafny.Int) bool {
					return func(_1177_i2 _dafny.Int) bool {
						return _1176_p0
					}
				})(p0)
				_ = _init30
				var _element0_30 = _init30(_dafny.Zero)
				_ = _element0_30
				_nw193 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
				(_nw193).ArraySet1(_element0_30, 0)
				var _nativeLen0_30 = (_len0_30).Int()
				_ = _nativeLen0_30
				for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
					(_nw193).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
				}
			}
			_1175_v8 = _nw193
			var _1178_v9 _dafny.MultiSet
			_ = _1178_v9
			_1178_v9 = _dafny.MultiSetOf(_1175_v8, _1175_v8, _1175_v8, _1175_v8)
			var _1179_v10 _dafny.MultiSet
			_ = _1179_v10
			_1179_v10 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1168_v4, p0)).Cardinality())
			var _1180_v11 *C4
			_ = _1180_v11
			var _nw194 *C4 = New_C4_()
			_ = _nw194
			_nw194.Ctor__((_1179_v10).IsProperSubsetOf(_dafny.MultiSetOf((_1174_v7).Dtor_cf29(), (_1178_v9).Cardinality())), (_1168_v4).F8(), (_1168_v4).F8())
			_1180_v11 = _nw194
			var _1181_v12 _dafny.CodePoint
			_ = _1181_v12
			_1181_v12 = _dafny.CodePoint('y')
			_1181_v12 = _1181_v12
			var _1182_v13 _dafny.Sequence
			_ = _1182_v13
			_1182_v13 = _dafny.SeqOf(p0, p0)
			var _1183_v14 _dafny.Sequence
			_ = _1183_v14
			_1183_v14 = _dafny.SeqOf(_1182_v13)
			var _1184_v15 *C3
			_ = _1184_v15
			var _nw195 *C3 = New_C3_()
			_ = _nw195
			_nw195.Ctor__((_1180_v11).F11(), (_1168_v4).F7(), (_this).F7())
			_1184_v15 = _nw195
			var _1185_v16 _dafny.Set
			_ = _1185_v16
			_1185_v16 = _dafny.SetOf(_1184_v15)
			var _1186_v17 *C3
			_ = _1186_v17
			var _nw196 *C3 = New_C3_()
			_ = _nw196
			_nw196.Ctor__(!(_dafny.Companion_Sequence_.Equal(_1182_v13, (_1183_v14).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1183_v14).Cardinality()))).Uint32()).(_dafny.Sequence))), _dafny.IntOfInt64(17), ((_1185_v16).Cardinality()).Plus(p1))
			_1186_v17 = _nw196
			var _1187_v18 _dafny.Sequence
			_ = _1187_v18
			_1187_v18 = _dafny.UnicodeSeqOfUtf8Bytes("eu")
			var _1188_v19 D8
			_ = _1188_v19
			_1188_v19 = Companion_D8_.Create_DC21_(_1184_v15)
			var _rhs215 *C3 = (_1188_v19).Dtor_cf34()
			_ = _rhs215
			var _rhs216 _dafny.Int = (_1168_v4).F7()
			_ = _rhs216
			var _rhs217 _dafny.Sequence = _dafny.Companion_Sequence_.Update((_1186_v17).Fm2((_1186_v17).F12(), globalState), (Companion_Default___.SafeIndex((_1168_v4).F8(), _dafny.IntOfUint32(((_1186_v17).Fm2((_1186_v17).F12(), globalState)).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
				if (_1184_v15).F12() {
					return _1181_v12
				}
				return _1181_v12
			})())
			_ = _rhs217
			_1186_v17 = _rhs215
			r0 = _rhs216
			_1187_v18 = _rhs217
			r2 = (_1184_v15).F12()
		}
		var _1189_v20 _dafny.CodePoint
		_ = _1189_v20
		_1189_v20 = _dafny.CodePoint('a')
		var _1190_v21 _dafny.Set
		_ = _1190_v21
		_1190_v21 = _dafny.SetOf(false)
		var _1191_v22 _dafny.Map
		_ = _1191_v22
		_1191_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _1190_v21)
		var _1192_v23 _dafny.Map
		_ = _1192_v23
		_1192_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1168_v4).F7(), (func() _dafny.Set {
			if (_1191_v22).Contains(true) {
				return (_1191_v22).Get(true).(_dafny.Set)
			}
			return _1190_v21
		})())
		var _1193_v24 _dafny.Sequence
		_ = _1193_v24
		_1193_v24 = _dafny.SeqOf(p0, p0)
		var _1194_v25 _dafny.Map
		_ = _1194_v25
		_1194_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1189_v20, (func() _dafny.Set {
			if (_1192_v23).Contains(_dafny.IntOfUint32((_1193_v24).Cardinality())) {
				return (_1192_v23).Get(_dafny.IntOfUint32((_1193_v24).Cardinality())).(_dafny.Set)
			}
			return _dafny.SetOf(!(p0))
		})())
		_1194_v25 = (_1194_v25).Update(_1189_v20, _1190_v21)
		r0 = _dafny.IntOfInt64(872)
		r1 = (_1168_v4).F7()
		r2 = p0
		return r0, r1, r2
	}
}

// End of class C6
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
