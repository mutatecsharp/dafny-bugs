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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.CodePoint, globalState *GlobalState) bool {
	return !_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(916))).Uint32(), func(coer0 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetFromSeq(_dafny.SeqOf(true))
	})), (_dafny.MultiSetOf(false)).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false, false, false, false))))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('m')
}
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) D0 {
	var _source0 _dafny.MultiSet = (func() _dafny.MultiSet {
		if !(!(!(false))) {
			return _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(709))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_1_i0 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus((_dafny.SetOf((_dafny.SetOf(_dafny.IntOfInt64(-390))).Cardinality())).Cardinality())
			})), _dafny.SeqOf((func() _dafny.Map {
				var _coll0 = _dafny.NewMapBuilder()
				_ = _coll0
				for _iter0 := _dafny.Iterate((func() _dafny.Set {
					var _coll1 = _dafny.NewBuilder()
					_ = _coll1
					for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_dafny.SeqOf(false)), true)).Keys().Elements()); ; {
						_compr_1, _ok1 := _iter1()
						if !_ok1 {
							break
						}
						var _2_v1 D3
						_2_v1 = interface{}(_compr_1).(D3)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_dafny.SeqOf(false)), true)).Contains(_2_v1) {
							_coll1.Add(_2_v1)
						}
					}
					return _coll1.ToSet()
				}()).Elements()); ; {
					_compr_0, _ok0 := _iter0()
					if !_ok0 {
						break
					}
					var _3_v0 D3
					_3_v0 = interface{}(_compr_0).(D3)
					if (func() _dafny.Set {
						var _coll2 = _dafny.NewBuilder()
						_ = _coll2
						for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_dafny.SeqOf(false)), true)).Keys().Elements()); ; {
							_compr_2, _ok2 := _iter2()
							if !_ok2 {
								break
							}
							var _4_v1 D3
							_4_v1 = interface{}(_compr_2).(D3)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_dafny.SeqOf(false)), true)).Contains(_4_v1) {
								_coll2.Add(_4_v1)
							}
						}
						return _coll2.ToSet()
					}()).Contains(_3_v0) {
						_coll0.Add(_3_v0, true)
					}
				}
				return _coll0.ToMap()
			}()).Cardinality(), _dafny.IntOfInt64(120)), _dafny.SeqOf(_dafny.IntOfInt64(796)), _dafny.SeqOf(_dafny.IntOfInt64(126)), _dafny.SeqOf(_dafny.IntOfInt64(285), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ngfivc")).Cardinality())))
		}
		return _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(-535), _dafny.IntOfInt64(-862)), _dafny.SeqOf(_dafny.IntOfInt64(569))))
	})()
	_ = _source0
	var _5___mcc_h0 _dafny.MultiSet = _source0
	_ = _5___mcc_h0
	var _6_cf88 _dafny.MultiSet = _5___mcc_h0
	_ = _6_cf88
	return Companion_D0_.Create_DC1_()
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((_dafny.MultiSetOf((_dafny.IntOfInt64(-560)).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-401))).Cardinality()))).Cardinality())))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(710))).Uint32(), func(coer2 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_7_i0 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('m'))).Cardinality()), _dafny.IntOfInt64(-29), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-691), (_dafny.MultiSetOf(true, !(true))).Cardinality())).Cardinality(), _dafny.IntOfInt64(-9))
	})), (func() _dafny.Sequence {
		if true {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(314))).Uint32(), func(coer3 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_8_i1 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(true, false, false, true))).Cardinality()))
			}))
		}
		return _dafny.SeqOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(967), _dafny.IntOfInt64(448))).Cardinality(), (func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(366), _dafny.IntOfInt64(228))); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _9_v0 _dafny.Int
				_9_v0 = interface{}(_compr_3).(_dafny.Int)
				if ((_dafny.IntOfInt64(366)).Cmp(_9_v0) <= 0) && ((_9_v0).Cmp(_dafny.IntOfInt64(228)) < 0) {
					_coll3.Add((_9_v0).Minus(((Companion_D16_.Create_DC33_(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(734))).Cardinality(), _dafny.MultiSetOf(false), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("i")).Cardinality())), _dafny.SetOf(true))).Dtor_cf54()).Cardinality()), _dafny.CodePoint('y'))
				}
			}
			return _coll3.ToMap()
		}()).Cardinality())).Cardinality()))
	})())
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(603))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(157)))).Merge((func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SetOf(Companion_D0_.Create_DC1_())).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _10_v0 D0
			_10_v0 = interface{}(_compr_4).(D0)
			if (_dafny.SetOf(Companion_D0_.Create_DC1_())).Contains(_10_v0) {
				_coll4.Add(_10_v0, _dafny.IntOfInt64(558))
			}
		}
		return _coll4.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), (_dafny.MultiSetOf(_dafny.CodePoint('c'), _dafny.CodePoint('e'), _dafny.CodePoint('u'))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 bool, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(-606), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(167))).Cardinality(), _dafny.IntOfInt64(841), _dafny.IntOfInt64(909))).Cardinality()))).Cardinality(), (_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(464), (Companion_D21_.Create_DC43_(true, false, _dafny.SeqOf(!(false)))).Dtor_cf70())).Cardinality()).Times(_dafny.IntOfInt64(159))))
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, p1 bool, p2 bool, p3 D2, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("eh"), _dafny.UnicodeSeqOfUtf8Bytes("rgorcc")), _dafny.UnicodeSeqOfUtf8Bytes("oxshrry"))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('n'))).Cardinality()), _dafny.IntOfInt64(28)), (_dafny.SetOf(true)).Intersection(_dafny.SetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(false, false, false, false))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 bool, globalState *GlobalState) D3 {
	var _source1 D6 = Companion_D6_.Create_DC15_(_dafny.MultiSetOf(_dafny.SeqOf(false)))
	_ = _source1
	if _source1.Is_DC16() {
		var _11___mcc_h0 bool = _source1.Get_().(D6_DC16).Cf20
		_ = _11___mcc_h0
		var _12___mcc_h1 bool = _source1.Get_().(D6_DC16).Cf21
		_ = _12___mcc_h1
		var _13___mcc_h2 _dafny.Sequence = _source1.Get_().(D6_DC16).Cf22
		_ = _13___mcc_h2
		var _14_cf22 _dafny.Sequence = _13___mcc_h2
		_ = _14_cf22
		var _15_cf21 bool = _12___mcc_h1
		_ = _15_cf21
		var _16_cf20 bool = _11___mcc_h0
		_ = _16_cf20
		return Companion_D3_.Create_DC8_(_dafny.SeqOf(false, _16_cf20))
	} else {
		var _17___mcc_h3 _dafny.MultiSet = _source1.Get_().(D6_DC15).Cf19
		_ = _17___mcc_h3
		var _18_cf19 _dafny.MultiSet = _17___mcc_h3
		_ = _18_cf19
		if false {
			return Companion_D3_.Create_DC8_(_dafny.SeqOf(true))
		} else {
			return Companion_D3_.Create_DC8_(_dafny.SeqOf(!(false), true))
		}
	}
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 _dafny.MultiSet, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	var _source2 D19 = Companion_D19_.Create_DC38_(false, !(true))
	_ = _source2
	if _source2.Is_DC38() {
		var _19___mcc_h0 bool = _source2.Get_().(D19_DC38).Cf63
		_ = _19___mcc_h0
		var _20___mcc_h1 bool = _source2.Get_().(D19_DC38).Cf64
		_ = _20___mcc_h1
		var _21_cf64 bool = _20___mcc_h1
		_ = _21_cf64
		var _22_cf63 bool = _19___mcc_h0
		_ = _22_cf63
		return _dafny.MultiSetOf(_21_cf64, _21_cf64, _22_cf63, _21_cf64)
	} else if _source2.Is_DC37() {
		var _23___mcc_h2 _dafny.Map = _source2.Get_().(D19_DC37).Cf62
		_ = _23___mcc_h2
		var _24_cf62 _dafny.Map = _23___mcc_h2
		_ = _24_cf62
		return (Companion_D16_.Create_DC33_(true, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(250))).Cardinality())).Cardinality()), _dafny.MultiSetOf(true), _dafny.IntOfInt64(527), _dafny.SetOf(!(true)))).Dtor_cf54()
	} else {
		var _25___mcc_h3 D19 = _source2.Get_().(D19_DC39).Cf65
		_ = _25___mcc_h3
		var _26_cf65 D19 = _25___mcc_h3
		_ = _26_cf65
		return (_dafny.MultiSetOf(false, true)).Union(_dafny.MultiSetOf(true, true))
	}
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_dafny.MultiSetOf(true)).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("n"))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(855), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xt")).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(290))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_27_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	})))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())), _dafny.UnicodeSeqOfUtf8Bytes("pmrrdi"))))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(false)
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.SeqOf(!(!(true))), _dafny.SeqOf(true))).Union(_dafny.MultiSetOf(_dafny.SeqOf(false)))).Difference(_dafny.MultiSetOf(_dafny.SeqOf(true)))
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D0_.Create_DC1_(), Companion_D0_.Create_DC1_()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D0_.Create_DC1_(), Companion_D0_.Create_DC1_(), Companion_D0_.Create_DC1_(), Companion_D0_.Create_DC1_()), _dafny.SeqOf(Companion_D0_.Create_DC1_())))
}
func (_static *CompanionStruct_Default___) Fm27(p0 bool, p1 _dafny.Int, p2 _dafny.MultiSet, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.SeqOf(false))).Union((_dafny.MultiSetOf(_dafny.SeqOf(false, !(true), !(true), false), _dafny.SeqOf(false, true, true))).Difference(_dafny.MultiSetOf(_dafny.SeqOf(true))))
}
func (_static *CompanionStruct_Default___) Fm28(globalState *GlobalState) _dafny.Map {
	if (_dafny.IntOfInt64(958)).Cmp(_dafny.IntOfInt64(397)) == 0 {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-268))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(388))
	}
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(452), _dafny.IntOfInt64(863))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(true, false), _dafny.SeqOf(true, true, false, !(true), !(true)), _dafny.SeqOf(true, true, false))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _28_v0 _dafny.Sequence
			_28_v0 = interface{}(_compr_5).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.SeqOf(true, false), _dafny.SeqOf(true, true, false, !(true), !(true)), _dafny.SeqOf(true, true, false))).Contains(_28_v0) {
				_coll5.Add(_28_v0, !(true))
			}
		}
		return _coll5.ToMap()
	}()).Cardinality(), _dafny.IntOfInt64(-354)))
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(812)), _dafny.SeqOf(_dafny.IntOfInt64(-968), _dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(573))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_29_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()), (func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(122), _dafny.IntOfInt64(-15))); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _30_v0 _dafny.Int
				_30_v0 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(122)).Cmp(_30_v0) <= 0) && ((_30_v0).Cmp(_dafny.IntOfInt64(-15)) < 0) {
					_coll6.Add((_30_v0).Times(_dafny.IntOfInt64(-977)), true)
				}
			}
			return _coll6.ToMap()
		}()).Cardinality())
	})))
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.IntOfInt64(-505), (_dafny.SetOf(_dafny.IntOfInt64(-506))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf((_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.UnicodeSeqOfUtf8Bytes("vyojj"))).Cardinality())).Cardinality())).Intersection(_dafny.SetOf(_dafny.IntOfInt64(864)))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC15_((_dafny.MultiSetOf(_dafny.SeqOf(true, true))).Union(_dafny.MultiSetOf(_dafny.SeqOf(false, true))))
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate((_dafny.SeqOf(_dafny.MultiSetOf((_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfInt64(236)))).Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _31_v0 _dafny.MultiSet
			_31_v0 = interface{}(_compr_7).(_dafny.MultiSet)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.MultiSetOf((_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfInt64(236))), _31_v0) {
				_coll7.Add(_31_v0, _dafny.UnicodeSeqOfUtf8Bytes("e"))
			}
		}
		return _coll7.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(299), _dafny.IntOfInt64(935), _dafny.IntOfInt64(-557), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("chaced")).Cardinality()), _dafny.IntOfInt64(934))).Cardinality(), _dafny.IntOfInt64(375))), _dafny.UnicodeSeqOfUtf8Bytes("jnbkik")))).Merge((func() _dafny.Map {
		if false {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(343))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}(func(_32_i0 _dafny.Int) _dafny.Int {
				return (_dafny.MultiSetOf(true)).Cardinality()
			}))), _dafny.UnicodeSeqOfUtf8Bytes("wdpr"))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(25))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_33_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
		})), _dafny.IntOfInt64(-972))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(331))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_34_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		})))
	})())
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tghobrq")).Cardinality()))).Intersection(_dafny.SetOf(_dafny.IntOfInt64(-815)))).Intersection((_dafny.SetOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.SetOf((_dafny.SetOf(false, true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality()))).Cardinality())), _dafny.IntOfInt64(186))).Union(_dafny.SetOf(_dafny.IntOfInt64(858), _dafny.IntOfInt64(-526), (Companion_D11_.Create_DC24_(!(true), true, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tfoksstqw")).Cardinality())))).Dtor_cf38(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(653))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_35_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm36(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(35)), _dafny.SeqOf((_dafny.MultiSetOf(!(!(false)), true)).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(918))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_36_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(903)
	})))
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("gcdnnimh")
}
func (_static *CompanionStruct_Default___) Fm38(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) D0 {
	if (_dafny.IntOfInt64(472)).Cmp((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ky")).Cardinality()))).Cardinality()) < 0 {
		return Companion_D0_.Create_DC0_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ys")).Cardinality()))).Cardinality())))
	} else if !(true) {
		return Companion_D0_.Create_DC0_(_dafny.IntOfInt64(621))
	} else {
		return Companion_D0_.Create_DC0_(_dafny.IntOfInt64(820))
	}
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Sequence, p1 D12, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("tjnq"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(462))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_37_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), _dafny.CodePoint('d'))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm40(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('y'), _dafny.CodePoint('a'))).Cardinality())))).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()), _dafny.IntOfInt64(-193))).Difference(func() _dafny.Set {
		var _coll8 = _dafny.NewBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(763), _dafny.IntOfInt64(69))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _38_v0 _dafny.Int
			_38_v0 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(763)).Cmp(_38_v0) <= 0) && ((_38_v0).Cmp(_dafny.IntOfInt64(69)) < 0) {
				_coll8.Add((_38_v0).Minus(_dafny.IntOfInt64(256)))
			}
		}
		return _coll8.ToSet()
	}())).Intersection((func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(288), _dafny.IntOfInt64(859))); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _39_v1 _dafny.Int
				_39_v1 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(288)).Cmp(_39_v1) <= 0) && ((_39_v1).Cmp(_dafny.IntOfInt64(859)) < 0) {
					_coll10.Add((_39_v1).Plus(_dafny.IntOfInt64(596)), (_dafny.SetOf(false, false, true)).Cardinality())
				}
			}
			return _coll10.ToMap()
		}()).Keys().Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _40_v2 _dafny.Int
			_40_v2 = interface{}(_compr_9).(_dafny.Int)
			if (func() _dafny.Map {
				var _coll11 = _dafny.NewMapBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(288), _dafny.IntOfInt64(859))); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _39_v1 _dafny.Int
					_39_v1 = interface{}(_compr_11).(_dafny.Int)
					if ((_dafny.IntOfInt64(288)).Cmp(_39_v1) <= 0) && ((_39_v1).Cmp(_dafny.IntOfInt64(859)) < 0) {
						_coll11.Add((_39_v1).Plus(_dafny.IntOfInt64(596)), (_dafny.SetOf(false, false, true)).Cardinality())
					}
				}
				return _coll11.ToMap()
			}()).Contains(_40_v2) {
				_coll9.Add((_40_v2).Plus(_dafny.IntOfInt64(-500)))
			}
		}
		return _coll9.ToSet()
	}()).Union(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, true, false)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.CodePoint, p1 _dafny.Set, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(459), (_dafny.Zero).Minus(_dafny.IntOfInt64(-980)))).Merge(func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(581), _dafny.IntOfInt64(197))); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _41_v0 _dafny.Int
			_41_v0 = interface{}(_compr_12).(_dafny.Int)
			if ((_dafny.IntOfInt64(581)).Cmp(_41_v0) <= 0) && ((_41_v0).Cmp(_dafny.IntOfInt64(197)) < 0) {
				_coll12.Add(Companion_Default___.SafeDivisionInt(_41_v0, _dafny.IntOfInt64(501)), (_dafny.MultiSetFromSeq(_dafny.SeqOf(!(!(!(!(!(false))))), false))).Cardinality())
			}
		}
		return _coll12.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm43(globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC16_(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("a"), _dafny.UnicodeSeqOfUtf8Bytes("exmbhmdo")), !(true) || (true), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-346)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("onku")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm44(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(445))
}
func (_static *CompanionStruct_Default___) Fm45(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(true, false), _dafny.MultiSetOf(true, false, false, true, true), _dafny.MultiSetOf(false)), (func() _dafny.Sequence {
		if !(true) {
			return _dafny.SeqOf(_dafny.MultiSetOf(false), _dafny.MultiSetOf(true))
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(824))).Uint32(), func(coer12 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_42_i0 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetOf(!(false))
		}))
	})())
}
func (_static *CompanionStruct_Default___) Fm46(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	if true {
		return _dafny.SetOf(false)
	} else {
		return (Companion_D16_.Create_DC33_(false, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(false)).Cardinality())).Cardinality()), _dafny.MultiSetOf(false, true), _dafny.IntOfInt64(873), _dafny.SetOf(!(true), true))).Dtor_cf56()
	}
}
func (_static *CompanionStruct_Default___) Fm47(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), (func() _dafny.Sequence {
		if true {
			return _dafny.SeqOf(false, true)
		}
		return _dafny.SeqOf(!(true))
	})()))
}
func (_static *CompanionStruct_Default___) Fm49(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(191)), _dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(737), _dafny.IntOfInt64(-741))).Cardinality(), _dafny.IntOfInt64(344))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-224))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_43_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(606), (Companion_D6_.Create_DC16_(true, false, _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(576), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())).Cardinality()))).Dtor_cf20())).Cardinality())).Cardinality()))
	}))), _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm50(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	var _source3 D1 = Companion_D1_.Create_DC3_(!(!(true)))
	_ = _source3
	if _source3.Is_DC4() {
		var _44___mcc_h0 _dafny.Int = _source3.Get_().(D1_DC4).Cf3
		_ = _44___mcc_h0
		var _45___mcc_h1 _dafny.Int = _source3.Get_().(D1_DC4).Cf4
		_ = _45___mcc_h1
		var _46_cf4 _dafny.Int = _45___mcc_h1
		_ = _46_cf4
		var _47_cf3 _dafny.Int = _44___mcc_h0
		_ = _47_cf3
		return _dafny.UnicodeSeqOfUtf8Bytes("dyg")
	} else {
		var _48___mcc_h2 bool = _source3.Get_().(D1_DC3).Cf2
		_ = _48___mcc_h2
		var _49_cf2 bool = _48___mcc_h2
		_ = _49_cf2
		return _dafny.UnicodeSeqOfUtf8Bytes("dstoi")
	}
}
func (_static *CompanionStruct_Default___) Fm51(p0 _dafny.Map, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(891), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(372))).Cardinality()), false))
}
func (_static *CompanionStruct_Default___) Fm52(p0 _dafny.Int, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC8_((func() _dafny.Sequence {
		if false {
			return _dafny.SeqOf((Companion_D16_.Create_DC33_(false, _dafny.IntOfInt64(-609), _dafny.MultiSetFromSeq(_dafny.SeqOf(true, true, false, false)), _dafny.IntOfInt64(-454), _dafny.SetOf(false))).Dtor_cf52())
		}
		return _dafny.SeqOf(false)
	})())
}
func (_static *CompanionStruct_Default___) Fm53(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true, false)).Difference((_dafny.SetOf(false, false)).Intersection(_dafny.SetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm54(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(155))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_50_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false))
	}))))
}
func (_static *CompanionStruct_Default___) Fm55(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) D1 {
	var _source4 D21 = Companion_D21_.Create_DC43_(false, true, _dafny.SeqOf(!(!(true))))
	_ = _source4
	if _source4.Is_DC43() {
		var _51___mcc_h0 bool = _source4.Get_().(D21_DC43).Cf69
		_ = _51___mcc_h0
		var _52___mcc_h1 bool = _source4.Get_().(D21_DC43).Cf70
		_ = _52___mcc_h1
		var _53___mcc_h2 _dafny.Sequence = _source4.Get_().(D21_DC43).Cf71
		_ = _53___mcc_h2
		var _54_cf71 _dafny.Sequence = _53___mcc_h2
		_ = _54_cf71
		var _55_cf70 bool = _52___mcc_h1
		_ = _55_cf70
		var _56_cf69 bool = _51___mcc_h0
		_ = _56_cf69
		return Companion_D1_.Create_DC3_(false)
	} else if _source4.Is_DC44() {
		var _57___mcc_h3 bool = _source4.Get_().(D21_DC44).Cf72
		_ = _57___mcc_h3
		var _58_cf72 bool = _57___mcc_h3
		_ = _58_cf72
		return Companion_D1_.Create_DC3_(false)
	} else {
		var _59___mcc_h4 _dafny.Array = _source4.Get_().(D21_DC42).Cf68
		_ = _59___mcc_h4
		var _60_cf68 _dafny.Array = _59___mcc_h4
		_ = _60_cf68
		return Companion_D1_.Create_DC3_(true)
	}
}
func (_static *CompanionStruct_Default___) Fm56(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	if (false) && (!(true)) {
		return _dafny.SeqOf(false)
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(!(false)))
	}
}
func (_static *CompanionStruct_Default___) Fm57(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(189))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_61_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(681)))).Cardinality())
	})), _dafny.SeqOf(_dafny.IntOfInt64(-469))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-661))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_62_i1 _dafny.Int) _dafny.Int {
		return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), false))).Cardinality()
	})))
}
func (_static *CompanionStruct_Default___) Fm58(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D12_.Create_DC27_(_dafny.SeqOf(_dafny.SeqOf(true))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D12_.Create_DC27_(_dafny.SeqOf(_dafny.SeqOf(true, false)))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D12_.Create_DC27_(_dafny.SeqOf(_dafny.SeqOf(false), _dafny.SeqOf(true)))))
}
func (_static *CompanionStruct_Default___) Fm59(globalState *GlobalState) _dafny.Map {
	var _source5 D21 = Companion_D21_.Create_DC43_(true, false, _dafny.SeqOf(false, false))
	_ = _source5
	if _source5.Is_DC43() {
		var _63___mcc_h0 bool = _source5.Get_().(D21_DC43).Cf69
		_ = _63___mcc_h0
		var _64___mcc_h1 bool = _source5.Get_().(D21_DC43).Cf70
		_ = _64___mcc_h1
		var _65___mcc_h2 _dafny.Sequence = _source5.Get_().(D21_DC43).Cf71
		_ = _65___mcc_h2
		var _66_cf71 _dafny.Sequence = _65___mcc_h2
		_ = _66_cf71
		var _67_cf70 bool = _64___mcc_h1
		_ = _67_cf70
		var _68_cf69 bool = _63___mcc_h0
		_ = _68_cf69
		return func() _dafny.Map {
			var _coll13 = _dafny.NewMapBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_dafny.SeqOf(Companion_D10_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_cf69, _dafny.IntOfInt64(804))))).Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _69_v0 D10
				_69_v0 = interface{}(_compr_13).(D10)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D10_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_cf69, _dafny.IntOfInt64(804)))), _69_v0) {
					_coll13.Add(_69_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sg")).Cardinality()))
				}
			}
			return _coll13.ToMap()
		}()
	} else if _source5.Is_DC44() {
		var _70___mcc_h3 bool = _source5.Get_().(D21_DC44).Cf72
		_ = _70___mcc_h3
		var _71_cf72 bool = _70___mcc_h3
		_ = _71_cf72
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D10_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_cf72, _dafny.IntOfInt64(18))), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("c")).Cardinality()))))
	} else {
		var _72___mcc_h4 _dafny.Array = _source5.Get_().(D21_DC42).Cf68
		_ = _72___mcc_h4
		var _73_cf68 _dafny.Array = _72___mcc_h4
		_ = _73_cf68
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D10_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(2))).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(776))).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D10_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(335))), (_dafny.SetOf(_dafny.IntOfInt64(377), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(762), _dafny.IntOfInt64(-921))).Cardinality())))).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm60(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false), _dafny.SeqOf(false)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(909), _dafny.IntOfInt64(368))).Cardinality(), !(!(!(false))))).Merge(func() _dafny.Map {
		var _coll14 = _dafny.NewMapBuilder()
		_ = _coll14
		for _iter14 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(896))).Elements()); ; {
			_compr_14, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			var _74_v0 _dafny.Int
			_74_v0 = interface{}(_compr_14).(_dafny.Int)
			if (_dafny.MultiSetOf(_dafny.IntOfInt64(896))).Contains(_74_v0) {
				_coll14.Add((_74_v0).Minus(_dafny.IntOfInt64(75)), true)
			}
		}
		return _coll14.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm61(p0 _dafny.Set, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(true, !(true), false, true, false)), _dafny.SeqOf(_dafny.SeqOf(!(!(false)), true, true, false, !(false))))
}
func (_static *CompanionStruct_Default___) Fm62(globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC12_((func() _dafny.CodePoint {
		if !(true) {
			return _dafny.CodePoint('p')
		}
		return _dafny.CodePoint('a')
	})())
}
func (_static *CompanionStruct_Default___) Fm63(globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC13_(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hxnhw")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aog")).Cardinality())))).Cardinality(), (false) == (false), _dafny.IntOfInt64(850), _dafny.SeqOf(!(false)))
}
func (_static *CompanionStruct_Default___) Fm64(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.CodePoint('e'))).Union(_dafny.SetOf(_dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('x')))).Union((func() _dafny.Set {
		var _coll15 = _dafny.NewBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.CodePoint('f'))).Elements()); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _75_v0 _dafny.CodePoint
			_75_v0 = interface{}(_compr_15).(_dafny.CodePoint)
			if (_dafny.MultiSetOf(_dafny.CodePoint('f'))).Contains(_75_v0) {
				_coll15.Add(_75_v0)
			}
		}
		return _coll15.ToSet()
	}()).Difference(_dafny.SetOf(_dafny.CodePoint('a'))))
}
func (_static *CompanionStruct_Default___) Fm65(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('f'))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('r')))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('q'))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.CodePoint('j'))))
}
func (_static *CompanionStruct_Default___) Fm66(p0 _dafny.Int, globalState *GlobalState) D16 {
	return Companion_D16_.Create_DC33_(!(true) || (false), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gk")).Cardinality()), _dafny.IntOfInt64(925)), ((Companion_D16_.Create_DC33_(false, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.MultiSetOf(!(false)), _dafny.IntOfInt64(249), _dafny.SetOf(false))).Dtor_cf54()).Intersection(_dafny.MultiSetOf(true)), _dafny.IntOfInt64(-152), (_dafny.SetOf(false)).Difference(_dafny.SetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm67(p0 bool, globalState *GlobalState) _dafny.Map {
	if true {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(876), _dafny.IntOfInt64(861)), !(!(true)))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(720), _dafny.IntOfInt64(548))).Cardinality(), (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("rss"), _dafny.UnicodeSeqOfUtf8Bytes("s"))).Cardinality()), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("udygyb")).Cardinality()), _dafny.IntOfInt64(-885)), !(true)))
	}
}
func (_static *CompanionStruct_Default___) Fm68(globalState *GlobalState) D12 {
	return Companion_D12_.Create_DC27_((func() _dafny.Sequence {
		if false {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(223))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}(func(_76_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(true)
			}))
		}
		return _dafny.SeqOf(_dafny.SeqOf(false))
	})())
}
func (_static *CompanionStruct_Default___) Fm69(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-3)))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(-982)))).Cardinality(), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("igvytirrh"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(614))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_77_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))))
}
func (_static *CompanionStruct_Default___) Fm70(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(7))).Uint32(), func(coer19 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}(func(_78_i0 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetOf(_dafny.CodePoint('d'))
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(866))).Uint32(), func(coer20 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg20 _dafny.Int) interface{} {
			return coer20(arg20)
		}
	}(func(_79_i1 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetOf(_dafny.CodePoint('j'), _dafny.CodePoint('q'))
	}))), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('k')), _dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('c'), _dafny.CodePoint('f'), _dafny.CodePoint('b')), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('q'))), _dafny.MultiSetOf(_dafny.CodePoint('b'), _dafny.CodePoint('c')), _dafny.MultiSetOf(_dafny.CodePoint('h'))))
}
func (_static *CompanionStruct_Default___) Fm71(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false))
}
func (_static *CompanionStruct_Default___) Fm72(p0 _dafny.Map, p1 _dafny.MultiSet, globalState *GlobalState) _dafny.MultiSet {
	var _source6 _dafny.Set = _dafny.SetOf(_dafny.IntOfInt64(350), _dafny.IntOfUint32((_dafny.SeqOf(true, !(true), true)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(846))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}(func(_80_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))).Cardinality()))
	_ = _source6
	var _81___mcc_h0 _dafny.Set = _source6
	_ = _81___mcc_h0
	var _82_cf23 _dafny.Set = _81___mcc_h0
	_ = _82_cf23
	return _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(510))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg22 _dafny.Int) interface{} {
			return coer22(arg22)
		}
	}(func(_83_i1 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(_dafny.IntOfInt64(180))
	})))
}
func (_static *CompanionStruct_Default___) Fm73(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.SeqOf((func() _dafny.Map {
		var _coll16 = _dafny.NewMapBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(544), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qkivckp")).Cardinality()))).Keys().Elements()); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _84_v0 _dafny.Int
			_84_v0 = interface{}(_compr_16).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(544), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qkivckp")).Cardinality()))).Contains(_84_v0) {
				_coll16.Add(Companion_Default___.SafeModuloInt(_84_v0, _dafny.IntOfInt64(688)), false)
			}
		}
		return _coll16.ToMap()
	}()).Cardinality()), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true), !(false))).Cardinality()))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(732)))))).Union(_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(519))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}(func(_85_i0 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(602), _dafny.IntOfInt64(-700))).Cardinality()
	}))))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Map, p1 bool, p2 _dafny.MultiSet, p3 bool, globalState *GlobalState) (bool, _dafny.Int) {
	var r0 bool = false
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	if true {
		var _86_v0 _dafny.Array
		_ = _86_v0
		var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
		_ = _nw0
		_86_v0 = _nw0
		var _87_v1 _dafny.Map
		_ = _87_v1
		_87_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _86_v0)
		_86_v0 = (func() _dafny.Array {
			if (_87_v1).Contains(p3) {
				return (_87_v1).Get(p3).(_dafny.Array)
			}
			return _86_v0
		})()
		var _88_v3 _dafny.Int
		_ = _88_v3
		_88_v3 = _dafny.IntOfInt64(50)
		var _89_v4 _dafny.Sequence
		_ = _89_v4
		_89_v4 = _dafny.SeqOf(_88_v3, _88_v3, _88_v3)
		var _90_v5 _dafny.Set
		_ = _90_v5
		_90_v5 = _dafny.SetOf(_88_v3)
		var _91_v6 _dafny.MultiSet
		_ = _91_v6
		_91_v6 = _dafny.MultiSetOf(p3)
		var _92_v7 _dafny.Set
		_ = _92_v7
		_92_v7 = _dafny.SetOf(p3, p1, p3, p1, p3)
		var _93_v8 _dafny.Map
		_ = _93_v8
		_93_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
		var _94_v9 _dafny.Sequence
		_ = _94_v9
		_94_v9 = _dafny.SeqOf(_93_v8, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p3), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1), _93_v8)
		var _95_v10 _dafny.Array
		_ = _95_v10
		var _nwElement0_0 bool = p3
		_ = _nwElement0_0
		var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(19))
		_ = _nw1
		(_nw1).ArraySet1(_nwElement0_0, 0)
		(_nw1).ArraySet1(p3, 1)
		(_nw1).ArraySet1(p1, 2)
		(_nw1).ArraySet1(((func() _dafny.Map {
			var _coll17 = _dafny.NewMapBuilder()
			_ = _coll17
			for _iter17 := _dafny.Iterate((_89_v4).Elements()); ; {
				_compr_17, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _96_v2 _dafny.Int
				_96_v2 = interface{}(_compr_17).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_89_v4, _96_v2) {
					_coll17.Add(Companion_Default___.SafeDivisionInt(_96_v2, _dafny.IntOfInt64(88)), (_dafny.Zero).Minus(_88_v3))
				}
			}
			return _coll17.ToMap()
		}()).Cardinality()).Cmp(_dafny.IntOfInt64(803)) != 0, 3)
		(_nw1).ArraySet1(p1, 4)
		(_nw1).ArraySet1(p1, 5)
		(_nw1).ArraySet1((_90_v5).IsProperSubsetOf(_90_v5), 6)
		(_nw1).ArraySet1(p1, 7)
		(_nw1).ArraySet1(p1, 8)
		(_nw1).ArraySet1((_88_v3).Cmp(_88_v3) <= 0, 9)
		(_nw1).ArraySet1((_91_v6).IsDisjointFrom(_91_v6), 10)
		(_nw1).ArraySet1((_92_v7).IsProperSubsetOf(Companion_Default___.Fm53(_88_v3, _94_v9, globalState)), 11)
		(_nw1).ArraySet1(p1, 12)
		(_nw1).ArraySet1(p1, 13)
		(_nw1).ArraySet1(p1, 14)
		(_nw1).ArraySet1(p1, 15)
		(_nw1).ArraySet1(!(p1) || (p3), 16)
		(_nw1).ArraySet1(p3, 17)
		(_nw1).ArraySet1(p3, 18)
		_95_v10 = _nw1
		var _97_v11 _dafny.Sequence
		_ = _97_v11
		_97_v11 = _dafny.SeqOf(true)
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_95_v10), 0))
		_ = _index0
		(_95_v10).ArraySet1((_97_v11).Select((Companion_Default___.SafeIndex(_88_v3, _dafny.IntOfUint32((_97_v11).Cardinality()))).Uint32()).(bool), (_index0).Int())
		var _98_v12 _dafny.Sequence
		_ = _98_v12
		_98_v12 = _dafny.UnicodeSeqOfUtf8Bytes("fku")
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_95_v10), 0))
		_ = _index1
		(_95_v10).ArraySet1((_90_v5).IsProperSubsetOf(_dafny.SetOf(_dafny.IntOfUint32((_98_v12).Cardinality()), _88_v3)), (_index1).Int())
		var _99_v13 _dafny.Array
		_ = _99_v13
		var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
		_ = _nw2
		_99_v13 = _nw2
		var _100_v14 *C4
		_ = _100_v14
		var _nw3 *C4 = New_C4_()
		_ = _nw3
		_nw3.Ctor__(_99_v13, (_dafny.Zero).Minus((_dafny.Zero).Minus(_88_v3)), (_95_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_95_v10), 0))).Int()).(bool))
		_100_v14 = _nw3
		_100_v14 = _100_v14
		var _101_v15 _dafny.Array
		_ = _101_v15
		var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
		_ = _nw4
		_101_v15 = _nw4
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_101_v15), 0))
		_ = _index2
		(_101_v15).ArraySet1(((_91_v6).Cardinality()).Minus(_88_v3), (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_101_v15), 0))
		_ = _index3
		(_101_v15).ArraySet1((((func() _dafny.Int {
			if (p2).Contains((_92_v7).Cardinality()) {
				return (p2).Multiplicity((_92_v7).Cardinality())
			}
			return _dafny.IntOfUint32((_98_v12).Cardinality())
		})()).Plus(_88_v3)).Plus(_88_v3), (_index3).Int())
		var _102_v16 *C2
		_ = _102_v16
		var _nw5 *C2 = New_C2_()
		_ = _nw5
		_nw5.Ctor__(_98_v12, (_97_v11).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_101_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_101_v15), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_97_v11).Cardinality()))).Uint32()).(bool))
		_102_v16 = _nw5
		var _103_v17 D27
		_ = _103_v17
		_103_v17 = Companion_D27_.Create_DC56_(_102_v16)
		_102_v16 = (_103_v17).Dtor_cf95()
	} else {
		var _104_v18 _dafny.CodePoint
		_ = _104_v18
		_104_v18 = _dafny.CodePoint('r')
		var _105_v19 _dafny.Array
		_ = _105_v19
		var _nwElement0_1 bool = p1
		_ = _nwElement0_1
		var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(8))
		_ = _nw6
		(_nw6).ArraySet1(_nwElement0_1, 0)
		(_nw6).ArraySet1(Companion_Default___.Fm0(_104_v18, globalState), 1)
		(_nw6).ArraySet1(p3, 2)
		(_nw6).ArraySet1(p3, 3)
		(_nw6).ArraySet1(p1, 4)
		(_nw6).ArraySet1(false, 5)
		(_nw6).ArraySet1(p3, 6)
		(_nw6).ArraySet1(p1, 7)
		_105_v19 = _nw6
		var _106_v20 _dafny.Map
		_ = _106_v20
		_106_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _105_v19)
		var _107_v21 _dafny.Array
		_ = _107_v21
		var _nwElement0_2 _dafny.Array = _105_v19
		_ = _nwElement0_2
		var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(22))
		_ = _nw7
		(_nw7).ArraySet1(_nwElement0_2, 0)
		(_nw7).ArraySet1(_105_v19, 1)
		(_nw7).ArraySet1(_105_v19, 2)
		(_nw7).ArraySet1(_105_v19, 3)
		(_nw7).ArraySet1(_105_v19, 4)
		(_nw7).ArraySet1(_105_v19, 5)
		(_nw7).ArraySet1(_105_v19, 6)
		(_nw7).ArraySet1(_105_v19, 7)
		(_nw7).ArraySet1(_105_v19, 8)
		(_nw7).ArraySet1(_105_v19, 9)
		(_nw7).ArraySet1(_105_v19, 10)
		(_nw7).ArraySet1(_105_v19, 11)
		(_nw7).ArraySet1(_105_v19, 12)
		(_nw7).ArraySet1(_105_v19, 13)
		(_nw7).ArraySet1(_105_v19, 14)
		(_nw7).ArraySet1(_105_v19, 15)
		(_nw7).ArraySet1(_105_v19, 16)
		(_nw7).ArraySet1(_105_v19, 17)
		(_nw7).ArraySet1(_105_v19, 18)
		(_nw7).ArraySet1(_105_v19, 19)
		(_nw7).ArraySet1((func() _dafny.Array {
			if (_106_v20).Contains(p3) {
				return (_106_v20).Get(p3).(_dafny.Array)
			}
			return _105_v19
		})(), 20)
		(_nw7).ArraySet1(_105_v19, 21)
		_107_v21 = _nw7
		var _108_v22 _dafny.Int
		_ = _108_v22
		_108_v22 = _dafny.IntOfInt64(852)
		var _109_v23 T1
		_ = _109_v23
		var _nw8 *C4 = New_C4_()
		_ = _nw8
		_nw8.Ctor__(_107_v21, _108_v22, p1)
		_109_v23 = _nw8
		var _110_v24 _dafny.Set
		_ = _110_v24
		_110_v24 = _dafny.SetOf(_109_v23)
		var _111_v25 D12
		_ = _111_v25
		_111_v25 = Companion_D12_.Create_DC28_((_110_v24).Cardinality(), _108_v22, (_109_v23).F4(), (_109_v23).F4(), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tbi")).Cardinality()), (_109_v23).F6()))
		_111_v25 = _111_v25
		var _112_v26 _dafny.Sequence
		_ = _112_v26
		_112_v26 = _dafny.UnicodeSeqOfUtf8Bytes("filnlnc")
		_112_v26 = _112_v26
		var _113_v27 _dafny.Sequence
		_ = _113_v27
		_113_v27 = _dafny.SeqOf((_109_v23).F6(), _108_v22)
		var _114_v28 _dafny.MultiSet
		_ = _114_v28
		_114_v28 = _dafny.MultiSetOf(p3)
		var _115_v29 _dafny.Array
		_ = _115_v29
		var _nwElement0_3 _dafny.Sequence = _113_v27
		_ = _nwElement0_3
		var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(22))
		_ = _nw9
		(_nw9).ArraySet1(_nwElement0_3, 0)
		(_nw9).ArraySet1(_113_v27, 1)
		(_nw9).ArraySet1(_113_v27, 2)
		(_nw9).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_112_v26).Cardinality())), 3)
		(_nw9).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_113_v27, _113_v27), 4)
		(_nw9).ArraySet1(_dafny.Companion_Sequence_.Update(_113_v27, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_112_v26).Cardinality()), _dafny.IntOfUint32((_113_v27).Cardinality()))).Uint32(), (_109_v23).F6()), 5)
		(_nw9).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_113_v27, _113_v27), 6)
		(_nw9).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_113_v27, _113_v27), 7)
		(_nw9).ArraySet1(_113_v27, 8)
		(_nw9).ArraySet1(_113_v27, 9)
		(_nw9).ArraySet1(_dafny.SeqOf((func() _dafny.Int {
			if (_114_v28).Contains(p1) {
				return (_114_v28).Multiplicity(p1)
			}
			return _108_v22
		})()), 10)
		(_nw9).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_113_v27, _dafny.SeqOf((_113_v27).Select((Companion_Default___.SafeIndex((_109_v23).F6(), _dafny.IntOfUint32((_113_v27).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_112_v26).Cardinality()), _108_v22, (func() _dafny.Int {
			if (_114_v28).Contains(false) {
				return (_114_v28).Multiplicity(false)
			}
			return (p2).Cardinality()
		})(), _108_v22)), 11)
		(_nw9).ArraySet1(_dafny.SeqOf(_108_v22, _dafny.IntOfUint32((_112_v26).Cardinality())), 12)
		(_nw9).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_109_v23).F6()), _113_v27), 13)
		(_nw9).ArraySet1(_113_v27, 14)
		(_nw9).ArraySet1(_113_v27, 15)
		(_nw9).ArraySet1(_113_v27, 16)
		(_nw9).ArraySet1(_113_v27, 17)
		(_nw9).ArraySet1(_113_v27, 18)
		(_nw9).ArraySet1(_113_v27, 19)
		(_nw9).ArraySet1(_113_v27, 20)
		(_nw9).ArraySet1(Companion_Default___.Fm57((_109_v23).F6(), (_109_v23).F6(), globalState), 21)
		_115_v29 = _nw9
		var _116_v30 _dafny.Map
		_ = _116_v30
		_116_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _108_v22)
		var _117_v31 D10
		_ = _117_v31
		_117_v31 = Companion_D10_.Create_DC21_((_116_v30).Update(p3, (_109_v23).F6()))
		var _118_v32 _dafny.Sequence
		_ = _118_v32
		_118_v32 = _dafny.SeqOf(_117_v31)
		var _119_v33 _dafny.Sequence
		_ = _119_v33
		_119_v33 = _dafny.SeqOf(Companion_Default___.Fm0(_104_v18, globalState), p3)
		var _rhs0 bool = !((_119_v33).Select((Companion_Default___.SafeIndex(_108_v22, _dafny.IntOfUint32((_119_v33).Cardinality()))).Uint32()).(bool))
		_ = _rhs0
		var _rhs1 _dafny.Array = _115_v29
		_ = _rhs1
		var _rhs2 _dafny.Sequence = (func() _dafny.Sequence {
			if false {
				return _118_v32
			}
			return _dafny.Companion_Sequence_.Update(_118_v32, (Companion_Default___.SafeIndex((_109_v23).F6(), _dafny.IntOfUint32((_118_v32).Cardinality()))).Uint32(), _117_v31)
		})()
		_ = _rhs2
		var _rhs3 _dafny.Int = (_109_v23).Fm7(globalState)
		_ = _rhs3
		r0 = _rhs0
		_115_v29 = _rhs1
		_118_v32 = _rhs2
		r1 = _rhs3
		var _120_v34 D16
		_ = _120_v34
		_120_v34 = Companion_D16_.Create_DC33_((_109_v23).F4(), (_109_v23).Fm9(p3, p3, (_109_v23).F6(), globalState), _114_v28, _dafny.IntOfUint32((_112_v26).Cardinality()), _dafny.SetOf((_109_v23).F4(), p1))
		r0 = !(!(((_120_v34).Dtor_cf54()).IsSubsetOf(_114_v28)) || ((_120_v34).Dtor_cf52()))
		var _121_v35 _dafny.Map
		_ = _121_v35
		_121_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v22, p3)
		var _122_v36 _dafny.Map
		_ = _122_v36
		_122_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_104_v18, _119_v33)
		var _123_v37 *C1
		_ = _123_v37
		var _nw10 *C1 = New_C1_()
		_ = _nw10
		_nw10.Ctor__((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(185), (func() _dafny.Int {
			if (_114_v28).Contains((func() bool {
				if (_121_v35).Contains(_dafny.IntOfInt64(884)) {
					return (_121_v35).Get(_dafny.IntOfInt64(884)).(bool)
				}
				return p1
			})()) {
				return (_114_v28).Multiplicity((func() bool {
					if (_121_v35).Contains(_dafny.IntOfInt64(884)) {
						return (_121_v35).Get(_dafny.IntOfInt64(884)).(bool)
					}
					return p1
				})())
			}
			return _108_v22
		})())), _108_v22, (func() _dafny.Array {
			if p1 {
				return _107_v21
			}
			return _109_v23.F5()
		})(), (_122_v36).Cardinality(), (p2).IsProperSubsetOf(p2))
		_123_v37 = _nw10
		var _rhs4 *C1 = _123_v37
		_ = _rhs4
		var _rhs5 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_113_v27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(476))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}((func(_124_v23 T1) func(_dafny.Int) _dafny.Int {
			return func(_125_i0 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus((_124_v23).F6())
			}
		})(_109_v23))))
		_ = _rhs5
		var _rhs6 _dafny.Array = _105_v19
		_ = _rhs6
		_123_v37 = _rhs4
		_113_v27 = _rhs5
		_105_v19 = _rhs6
	}
	var _126_v38 _dafny.Int
	_ = _126_v38
	_126_v38 = _dafny.IntOfInt64(197)
	var _hi0 _dafny.Int = Companion_Default___.Fm3(_126_v38, true, globalState)
	_ = _hi0
	for _127_i1 := _126_v38; _127_i1.Cmp(_hi0) < 0; _127_i1 = _127_i1.Plus(_dafny.One) {
		_126_v38 = Companion_Default___.SafeModuloInt(_126_v38, _126_v38)
		var _128_v39 _dafny.Sequence
		_ = _128_v39
		_128_v39 = _dafny.SeqOf(_126_v38)
		var _129_v40 _dafny.Sequence
		_ = _129_v40
		_129_v40 = _dafny.UnicodeSeqOfUtf8Bytes("etegojdr")
		r0 = (func() bool {
			if false {
				return p3
			}
			return !_dafny.Companion_Sequence_.Equal(_128_v39, _dafny.Companion_Sequence_.Update(_128_v39, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_129_v40).Cardinality()), _dafny.IntOfUint32((_128_v39).Cardinality()))).Uint32(), _127_i1))
		})()
		var _130_v41 *C2
		_ = _130_v41
		var _nw11 *C2 = New_C2_()
		_ = _nw11
		_nw11.Ctor__(_dafny.Companion_Sequence_.Concatenate(_129_v40, _129_v40), p3)
		_130_v41 = _nw11
		_130_v41 = (func() *C2 {
			if p1 {
				return _130_v41
			}
			return _130_v41
		})()
		var _131_v42 _dafny.Sequence
		_ = _131_v42
		_131_v42 = _dafny.SeqOf(!(p1))
		var _132_v43 _dafny.Set
		_ = _132_v43
		_132_v43 = _dafny.SetOf(_131_v42, _131_v42, _131_v42, _131_v42)
		var _133_v44 _dafny.Sequence
		_ = _133_v44
		_133_v44 = _dafny.SeqOf(_131_v42, _131_v42)
		var _134_v46 _dafny.Array
		_ = _134_v46
		var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
		_ = _nw12
		_134_v46 = _nw12
		var _135_v47 *C9
		_ = _135_v47
		var _nw13 *C9 = New_C9_()
		_ = _nw13
		_nw13.Ctor__(_127_i1, (_132_v43).Difference(func() _dafny.Set {
			var _coll18 = _dafny.NewBuilder()
			_ = _coll18
			for _iter18 := _dafny.Iterate((_133_v44).Elements()); ; {
				_compr_18, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _136_v45 _dafny.Sequence
				_136_v45 = interface{}(_compr_18).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_133_v44, _136_v45) {
					_coll18.Add(_136_v45)
				}
			}
			return _coll18.ToSet()
		}()), _134_v46, _127_i1, p3)
		_135_v47 = _nw13
	}
	var _137_v48 _dafny.Array
	_ = _137_v48
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(4)
	_ = _len0_0
	var _nw14 _dafny.Array
	_ = _nw14
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw14 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Int = func(_138_i2 _dafny.Int) _dafny.Int {
			return (_138_i2).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _dafny.IntOfInt64(33))).Cardinality())
		}
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
	_137_v48 = _nw14
	var _139_v49 _dafny.Sequence
	_ = _139_v49
	_139_v49 = _dafny.SeqOf(_137_v48, _137_v48, _137_v48, _137_v48)
	_139_v49 = _dafny.Companion_Sequence_.Concatenate(_139_v49, _139_v49)
	_126_v38 = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm3(_126_v38, p1, globalState), _126_v38)
	if p1 {
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))
		_ = _index4
		(_137_v48).ArraySet1(_dafny.IntOfInt64(946), (_index4).Int())
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))
		_ = _index5
		(_137_v48).ArraySet1(_dafny.IntOfInt64(-463), (_index5).Int())
		var _140_v50 _dafny.Array
		_ = _140_v50
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_1
		var _nw15 _dafny.Array
		_ = _nw15
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw15 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.CodePoint = func(_141_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			}
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw15 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw15).ArraySet1CodePoint(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw15).ArraySet1CodePoint(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_140_v50 = _nw15
		var _142_v51 _dafny.CodePoint
		_ = _142_v51
		_142_v51 = _dafny.CodePoint('j')
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_140_v50), 0))
		_ = _index6
		(_140_v50).ArraySet1CodePoint(_142_v51, (_index6).Int())
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_140_v50), 0))
		_ = _index7
		var _rhs7 _dafny.CodePoint = _142_v51
		_ = _rhs7
		var _rhs8 bool = !(p1) || (Companion_Default___.Fm0(_142_v51, globalState))
		_ = _rhs8
		var _rhs9 _dafny.Int = (_137_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))).Int()).(_dafny.Int)
		_ = _rhs9
		var _rhs10 bool = !(p3) || (!(!(false)))
		_ = _rhs10
		var _lhs0 _dafny.Array = _140_v50
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_140_v50), 0))
		_ = _lhs1
		(_lhs0).ArraySet1CodePoint(_rhs7, (_lhs1).Int())
		r0 = _rhs8
		r1 = _rhs9
		r0 = _rhs10
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_140_v50), 0))
		_ = _index8
		(_140_v50).ArraySet1CodePoint(_dafny.CodePoint('b'), (_index8).Int())
		var _143_v52 _dafny.Sequence
		_ = _143_v52
		_143_v52 = _dafny.SeqOf(p1)
		var _144_v53 _dafny.Sequence
		_ = _144_v53
		_144_v53 = _dafny.UnicodeSeqOfUtf8Bytes("bjpsw")
		var _145_v54 _dafny.Set
		_ = _145_v54
		_145_v54 = _dafny.SetOf(_143_v52, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(false, p3), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_144_v53).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, p3)).Cardinality()))).Uint32(), p3))
		var _146_v55 D11
		_ = _146_v55
		_146_v55 = Companion_D11_.Create_DC24_(p1, p3, (_137_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))).Int()).(_dafny.Int))
		var _147_v56 D1
		_ = _147_v56
		_147_v56 = Companion_D1_.Create_DC3_(p1)
		var _148_v57 _dafny.Array
		_ = _148_v57
		var _nwElement0_4 bool = p3
		_ = _nwElement0_4
		var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(27))
		_ = _nw16
		(_nw16).ArraySet1(_nwElement0_4, 0)
		(_nw16).ArraySet1(p3, 1)
		(_nw16).ArraySet1(p3, 2)
		(_nw16).ArraySet1(p1, 3)
		(_nw16).ArraySet1(true, 4)
		(_nw16).ArraySet1((_146_v55).Dtor_cf36(), 5)
		(_nw16).ArraySet1(p1, 6)
		(_nw16).ArraySet1(p3, 7)
		(_nw16).ArraySet1(p3, 8)
		(_nw16).ArraySet1((_147_v56).Dtor_cf2(), 9)
		(_nw16).ArraySet1(!(p3), 10)
		(_nw16).ArraySet1(p3, 11)
		(_nw16).ArraySet1(p1, 12)
		(_nw16).ArraySet1(p3, 13)
		(_nw16).ArraySet1(p3, 14)
		(_nw16).ArraySet1(p3, 15)
		(_nw16).ArraySet1(p3, 16)
		(_nw16).ArraySet1(p1, 17)
		(_nw16).ArraySet1(p1, 18)
		(_nw16).ArraySet1(!(p1), 19)
		(_nw16).ArraySet1(p3, 20)
		(_nw16).ArraySet1(false, 21)
		(_nw16).ArraySet1(p1, 22)
		(_nw16).ArraySet1(p3, 23)
		(_nw16).ArraySet1(!(p1), 24)
		(_nw16).ArraySet1(true, 25)
		(_nw16).ArraySet1(p1, 26)
		_148_v57 = _nw16
		var _149_v58 _dafny.Array
		_ = _149_v58
		var _nwElement0_5 _dafny.Array = _148_v57
		_ = _nwElement0_5
		var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(8))
		_ = _nw17
		(_nw17).ArraySet1(_nwElement0_5, 0)
		(_nw17).ArraySet1(_148_v57, 1)
		(_nw17).ArraySet1(_148_v57, 2)
		(_nw17).ArraySet1(_148_v57, 3)
		(_nw17).ArraySet1(_148_v57, 4)
		(_nw17).ArraySet1(_148_v57, 5)
		(_nw17).ArraySet1(_148_v57, 6)
		(_nw17).ArraySet1(_148_v57, 7)
		_149_v58 = _nw17
		var _150_v59 _dafny.Sequence
		_ = _150_v59
		_150_v59 = _dafny.SeqOf(_149_v58)
		var _151_v60 _dafny.Set
		_ = _151_v60
		_151_v60 = _dafny.SetOf(_126_v38)
		var _152_v61 D12
		_ = _152_v61
		_152_v61 = Companion_D12_.Create_DC28_((_151_v60).Cardinality(), _126_v38, p3, p3, _dafny.IntOfInt64(516))
		var _pat_let_tv0 = _126_v38
		_ = _pat_let_tv0
		var _153_v62 *C9
		_ = _153_v62
		var _nw18 *C9 = New_C9_()
		_ = _nw18
		_nw18.Ctor__((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(544))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}((func(_154_v50 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
			return func(_155_i4 _dafny.Int) _dafny.CodePoint {
				return (_154_v50).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_154_v50), 0))).Int())
			}
		})(_140_v50)))).Cardinality())).Plus((_137_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))).Int()).(_dafny.Int)), _145_v54, (_150_v59).Select((Companion_Default___.SafeIndex(_126_v38, _dafny.IntOfUint32((_150_v59).Cardinality()))).Uint32()).(_dafny.Array), (_137_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))).Int()).(_dafny.Int), (func(_pat_let0_0 D12) D12 {
			return func(_156_dt__update__tmp_h0 D12) D12 {
				return func(_pat_let1_0 _dafny.Int) D12 {
					return func(_157_dt__update_hcf44_h0 _dafny.Int) D12 {
						return func(_pat_let2_0 bool) D12 {
							return func(_158_dt__update_hcf45_h0 bool) D12 {
								return Companion_D12_.Create_DC28_((_156_dt__update__tmp_h0).Dtor_cf43(), _157_dt__update_hcf44_h0, _158_dt__update_hcf45_h0, (_156_dt__update__tmp_h0).Dtor_cf46(), (_156_dt__update__tmp_h0).Dtor_cf47())
							}(_pat_let2_0)
						}(true)
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(_152_v61)).Dtor_cf46())
		_153_v62 = _nw18
		var _159_v63 _dafny.Sequence
		_ = _159_v63
		_159_v63 = _dafny.SeqOf(_153_v62.F7)
		var _160_v64 _dafny.Sequence
		_ = _160_v64
		_160_v64 = _dafny.SeqOf((_137_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))).Int()).(_dafny.Int), _153_v62.F7, (_dafny.Zero).Minus(_153_v62.F7), _153_v62.F7, (_159_v63).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.IntOfUint32((_159_v63).Cardinality()))).Uint32()).(_dafny.Int))
		var _161_v65 D9
		_ = _161_v65
		_161_v65 = Companion_D9_.Create_DC20_(p3, (_140_v50).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_140_v50), 0))).Int()), _153_v62.F7)
		var _162_v66 _dafny.Array
		_ = _162_v66
		var _nwElement0_6 _dafny.Int = _126_v38
		_ = _nwElement0_6
		var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(13))
		_ = _nw19
		(_nw19).ArraySet1(_nwElement0_6, 0)
		(_nw19).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(975), false)).Cardinality(), _153_v62.F7), 1)
		(_nw19).ArraySet1(_153_v62.F7, 2)
		(_nw19).ArraySet1(_126_v38, 3)
		(_nw19).ArraySet1(_153_v62.F7, 4)
		(_nw19).ArraySet1((_dafny.Zero).Minus(_126_v38), 5)
		(_nw19).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_159_v63, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(729))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}((func(_163_v48 _dafny.Array) func(_dafny.Int) _dafny.Int {
			return func(_164_i5 _dafny.Int) _dafny.Int {
				return (_163_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_163_v48), 0))).Int()).(_dafny.Int)
			}
		})(_137_v48))))).Cardinality()), 6)
		(_nw19).ArraySet1((_137_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))).Int()).(_dafny.Int), 7)
		(_nw19).ArraySet1(Companion_Default___.Fm3((_161_v65).Dtor_cf28(), p3, globalState), 8)
		(_nw19).ArraySet1(_153_v62.F7, 9)
		(_nw19).ArraySet1((_137_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))).Int()).(_dafny.Int), 10)
		(_nw19).ArraySet1((func() _dafny.Int {
			if p1 {
				return _126_v38
			}
			return _dafny.IntOfInt64(715)
		})(), 11)
		(_nw19).ArraySet1(_153_v62.F7, 12)
		_162_v66 = _nw19
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))
		_ = _index9
		var _rhs11 *C9 = _153_v62
		_ = _rhs11
		var _rhs12 _dafny.Int = Companion_Default___.Fm3(_126_v38, p3, globalState)
		_ = _rhs12
		var _rhs13 _dafny.Array = _162_v66
		_ = _rhs13
		var _rhs14 bool = p1
		_ = _rhs14
		var _lhs2 _dafny.Array = _137_v48
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))
		_ = _lhs3
		_153_v62 = _rhs11
		(_lhs2).ArraySet1(_rhs12, (_lhs3).Int())
		_162_v66 = _rhs13
		r0 = _rhs14
		var _165_v67 *C6
		_ = _165_v67
		var _nw20 *C6 = New_C6_()
		_ = _nw20
		_nw20.Ctor__(p3, p3, _149_v58, _dafny.IntOfInt64(-356))
		_165_v67 = _nw20
		var _166_v68 _dafny.Set
		_ = _166_v68
		_166_v68 = _dafny.SetOf(_165_v67)
		var _167_v69 _dafny.Map
		_ = _167_v69
		_167_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v68, p1)
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))
		_ = _index10
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))
		_ = _index11
		var _rhs15 _dafny.Int = (_126_v38).Times(Companion_Default___.Fm3(_126_v38, false, globalState))
		_ = _rhs15
		var _rhs16 _dafny.Int = (_126_v38).Minus((_137_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))).Int()).(_dafny.Int))
		_ = _rhs16
		var _rhs17 _dafny.Int = ((_137_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))).Int()).(_dafny.Int)).Plus(Companion_Default___.SafeModuloInt(_153_v62.F7, _153_v62.F7))
		_ = _rhs17
		var _rhs18 _dafny.Int = ((func() _dafny.Int {
			if !(p1) {
				return (_167_v69).Cardinality()
			}
			return (_dafny.Zero).Minus((_137_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))).Int()).(_dafny.Int))
		})()).Plus((_126_v38).Plus(_153_v62.F7))
		_ = _rhs18
		var _lhs4 _dafny.Array = _137_v48
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))
		_ = _lhs5
		var _lhs6 _dafny.Array = _137_v48
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_137_v48), 0))
		_ = _lhs7
		(_lhs4).ArraySet1(_rhs15, (_lhs5).Int())
		_126_v38 = _rhs16
		(_lhs6).ArraySet1(_rhs17, (_lhs7).Int())
		r1 = _rhs18
	} else {
		var _168_v70 _dafny.CodePoint
		_ = _168_v70
		_168_v70 = _dafny.CodePoint('g')
		var _169_v71 _dafny.Map
		_ = _169_v71
		_169_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v70, _126_v38)
		var _170_v72 _dafny.Sequence
		_ = _170_v72
		_170_v72 = _dafny.SeqOf(_126_v38, _126_v38, (func() _dafny.Int {
			if (_169_v71).Contains(_168_v70) {
				return (_169_v71).Get(_168_v70).(_dafny.Int)
			}
			return _126_v38
		})(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(540))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}((func(_171_v38 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_172_i6 _dafny.Int) _dafny.Int {
				return _171_v38
			}
		})(_126_v38)))).Cardinality()))
		var _173_v73 D1
		_ = _173_v73
		_173_v73 = Companion_D1_.Create_DC4_(_126_v38, _126_v38)
		var _174_v74 _dafny.Map
		_ = _174_v74
		_174_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _126_v38)
		var _175_v75 _dafny.Set
		_ = _175_v75
		_175_v75 = _dafny.SetOf(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_170_v72).Cardinality()), _126_v38), _126_v38, (_173_v73).Dtor_cf4(), (func() _dafny.Int {
			if (_174_v74).Contains(p1) {
				return (_174_v74).Get(p1).(_dafny.Int)
			}
			return _126_v38
		})())
		_175_v75 = ((func() _dafny.Set {
			if !(p1) {
				return _175_v75
			}
			return _175_v75
		})()).Difference((_175_v75).Intersection(_175_v75))
		var _176_v76 _dafny.Map
		_ = _176_v76
		_176_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		var _177_v77 _dafny.MultiSet
		_ = _177_v77
		_177_v77 = _dafny.MultiSetOf(p3, p3)
		var _178_v78 D11
		_ = _178_v78
		_178_v78 = Companion_D11_.Create_DC24_(Companion_Default___.Fm0(_168_v70, globalState), (func() bool {
			if (_176_v76).Contains(false) {
				return (_176_v76).Get(false).(bool)
			}
			return p3
		})(), (_177_v77).Cardinality())
		var _179_v79 _dafny.MultiSet
		_ = _179_v79
		_179_v79 = _dafny.MultiSetOf(_178_v78)
		var _180_v80 _dafny.Array
		_ = _180_v80
		var _nwElement0_7 bool = p3
		_ = _nwElement0_7
		var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(26))
		_ = _nw21
		(_nw21).ArraySet1(_nwElement0_7, 0)
		(_nw21).ArraySet1(p3, 1)
		(_nw21).ArraySet1(p1, 2)
		(_nw21).ArraySet1(p1, 3)
		(_nw21).ArraySet1(p1, 4)
		(_nw21).ArraySet1(p1, 5)
		(_nw21).ArraySet1(p3, 6)
		(_nw21).ArraySet1(Companion_Default___.Fm0(_168_v70, globalState), 7)
		(_nw21).ArraySet1(!(false), 8)
		(_nw21).ArraySet1(p3, 9)
		(_nw21).ArraySet1(p3, 10)
		(_nw21).ArraySet1(p3, 11)
		(_nw21).ArraySet1(p3, 12)
		(_nw21).ArraySet1(!(p1), 13)
		(_nw21).ArraySet1(p1, 14)
		(_nw21).ArraySet1(!(p3), 15)
		(_nw21).ArraySet1(true, 16)
		(_nw21).ArraySet1(p3, 17)
		(_nw21).ArraySet1(false, 18)
		(_nw21).ArraySet1(p1, 19)
		(_nw21).ArraySet1(p3, 20)
		(_nw21).ArraySet1(p3, 21)
		(_nw21).ArraySet1(p1, 22)
		(_nw21).ArraySet1(p1, 23)
		(_nw21).ArraySet1(p3, 24)
		(_nw21).ArraySet1(p1, 25)
		_180_v80 = _nw21
		var _181_v81 _dafny.Array
		_ = _181_v81
		var _nwElement0_8 _dafny.Array = _180_v80
		_ = _nwElement0_8
		var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(26))
		_ = _nw22
		(_nw22).ArraySet1(_nwElement0_8, 0)
		(_nw22).ArraySet1(_180_v80, 1)
		(_nw22).ArraySet1(_180_v80, 2)
		(_nw22).ArraySet1(_180_v80, 3)
		(_nw22).ArraySet1(_180_v80, 4)
		(_nw22).ArraySet1(_180_v80, 5)
		(_nw22).ArraySet1(_180_v80, 6)
		(_nw22).ArraySet1(_180_v80, 7)
		(_nw22).ArraySet1(_180_v80, 8)
		(_nw22).ArraySet1(_180_v80, 9)
		(_nw22).ArraySet1(_180_v80, 10)
		(_nw22).ArraySet1(_180_v80, 11)
		(_nw22).ArraySet1(_180_v80, 12)
		(_nw22).ArraySet1(_180_v80, 13)
		(_nw22).ArraySet1(_180_v80, 14)
		(_nw22).ArraySet1(_180_v80, 15)
		(_nw22).ArraySet1(_180_v80, 16)
		(_nw22).ArraySet1(_180_v80, 17)
		(_nw22).ArraySet1(_180_v80, 18)
		(_nw22).ArraySet1(_180_v80, 19)
		(_nw22).ArraySet1(_180_v80, 20)
		(_nw22).ArraySet1(_180_v80, 21)
		(_nw22).ArraySet1(_180_v80, 22)
		(_nw22).ArraySet1(_180_v80, 23)
		(_nw22).ArraySet1(_180_v80, 24)
		(_nw22).ArraySet1(_180_v80, 25)
		_181_v81 = _nw22
		var _182_v82 *C5
		_ = _182_v82
		var _nw23 *C5 = New_C5_()
		_ = _nw23
		_nw23.Ctor__((p3) == (p3), (_179_v79).Cardinality(), _181_v81, _126_v38, p3)
		_182_v82 = _nw23
		_182_v82 = _182_v82
		var _183_v83 _dafny.Array
		_ = _183_v83
		var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
		_ = _nw24
		_183_v83 = _nw24
		var _184_v84 _dafny.Map
		_ = _184_v84
		_184_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v83, (_182_v82).F13())
		_184_v84 = (_184_v84).Update(_183_v83, p3)
		var _185_v85 _dafny.Sequence
		_ = _185_v85
		_185_v85 = _dafny.SeqOf(p1)
		var _186_v86 _dafny.Set
		_ = _186_v86
		_186_v86 = _dafny.SetOf(_185_v85)
		var _187_v87 *C9
		_ = _187_v87
		var _nw25 *C9 = New_C9_()
		_ = _nw25
		_nw25.Ctor__((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-613))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}((func(_188_v70 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_189_i7 _dafny.Int) _dafny.CodePoint {
				return _188_v70
			}
		})(_168_v70)))).Cardinality())), _186_v86, _181_v81, _182_v82.F14, p1)
		_187_v87 = _nw25
		var _190_v88 _dafny.Map
		_ = _190_v88
		_190_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_187_v87, p3)
		_190_v88 = (_190_v88).Update((func() *C9 {
			if (_182_v82).F13() {
				return _187_v87
			}
			return _187_v87
		})(), p1)
		var _191_v89 _dafny.Sequence
		_ = _191_v89
		_191_v89 = _dafny.UnicodeSeqOfUtf8Bytes("xu")
		var _192_v90 T0
		_ = _192_v90
		var _nw26 *C2 = New_C2_()
		_ = _nw26
		_nw26.Ctor__(_191_v89, p1)
		_192_v90 = _nw26
		var _193_v91 T0
		_ = _193_v91
		_193_v91 = _192_v90
		_193_v91 = _193_v91
	}
	if p3 {
		var _194_v92 _dafny.CodePoint
		_ = _194_v92
		_194_v92 = _dafny.CodePoint('k')
		var _195_v93 _dafny.Sequence
		_ = _195_v93
		_195_v93 = _dafny.UnicodeSeqOfUtf8Bytes("hvok")
		var _196_v94 D9
		_ = _196_v94
		_196_v94 = Companion_D9_.Create_DC19_(_195_v93)
		var _197_v95 T0
		_ = _197_v95
		var _nw27 *C2 = New_C2_()
		_ = _nw27
		_nw27.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-339))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_198_i8 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})), _dafny.Companion_Sequence_.Contains((_196_v94).Dtor_cf25(), _194_v92))
		_197_v95 = _nw27
		var _rhs19 T0 = _197_v95
		_ = _rhs19
		var _rhs20 _dafny.Int = (_dafny.Zero).Minus(_126_v38)
		_ = _rhs20
		_197_v95 = _rhs19
		_126_v38 = _rhs20
		var _199_v96 _dafny.MultiSet
		_ = _199_v96
		_199_v96 = _dafny.MultiSetOf((_197_v95).F4(), p3)
		var _200_v97 _dafny.Map
		_ = _200_v97
		_200_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p1)
		var _201_v98 _dafny.Sequence
		_ = _201_v98
		_201_v98 = _dafny.SeqOf(_200_v97, _200_v97)
		var _source7 D16 = Companion_D16_.Create_DC33_((_197_v95).F4(), _126_v38, (_199_v96).Difference(_199_v96), _126_v38, Companion_Default___.Fm53(_126_v38, _201_v98, globalState))
		_ = _source7
		if _source7.Is_DC33() {
			var _202___mcc_h0 bool = _source7.Get_().(D16_DC33).Cf52
			_ = _202___mcc_h0
			var _203___mcc_h1 _dafny.Int = _source7.Get_().(D16_DC33).Cf53
			_ = _203___mcc_h1
			var _204___mcc_h2 _dafny.MultiSet = _source7.Get_().(D16_DC33).Cf54
			_ = _204___mcc_h2
			var _205___mcc_h3 _dafny.Int = _source7.Get_().(D16_DC33).Cf55
			_ = _205___mcc_h3
			var _206___mcc_h4 _dafny.Set = _source7.Get_().(D16_DC33).Cf56
			_ = _206___mcc_h4
			var _207_cf56 _dafny.Set = _206___mcc_h4
			_ = _207_cf56
			var _208_cf55 _dafny.Int = _205___mcc_h3
			_ = _208_cf55
			var _209_cf54 _dafny.MultiSet = _204___mcc_h2
			_ = _209_cf54
			var _210_cf53 _dafny.Int = _203___mcc_h1
			_ = _210_cf53
			var _211_cf52 bool = _202___mcc_h0
			_ = _211_cf52
			var _212_v99 _dafny.Array
			_ = _212_v99
			var _nw28 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(20))
			_ = _nw28
			_212_v99 = _nw28
			var _213_v100 _dafny.Array
			_ = _213_v100
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
			_ = _nw29
			_213_v100 = _nw29
			var _214_v101 _dafny.Sequence
			_ = _214_v101
			_214_v101 = _dafny.SeqOf(_210_cf53)
			var _215_v102 *C4
			_ = _215_v102
			var _nw30 *C4 = New_C4_()
			_ = _nw30
			_nw30.Ctor__(_213_v100, (_214_v101).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_214_v101).Cardinality()), _dafny.IntOfUint32((_214_v101).Cardinality()))).Uint32()).(_dafny.Int), p3)
			_215_v102 = _nw30
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_212_v99), 0))
			_ = _index12
			(_212_v99).ArraySet1(_215_v102, (_index12).Int())
			var _216_v103 _dafny.Sequence
			_ = _216_v103
			_216_v103 = _dafny.SeqOf(!(p3))
			var _217_v106 _dafny.MultiSet
			_ = _217_v106
			_217_v106 = _dafny.MultiSetOf(func() _dafny.Map {
				var _coll19 = _dafny.NewMapBuilder()
				_ = _coll19
				for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(438), _dafny.IntOfInt64(391))); ; {
					_compr_19, _ok19 := _iter19()
					if !_ok19 {
						break
					}
					var _218_v104 _dafny.Int
					_218_v104 = interface{}(_compr_19).(_dafny.Int)
					if ((_dafny.IntOfInt64(438)).Cmp(_218_v104) <= 0) && ((_218_v104).Cmp(_dafny.IntOfInt64(391)) < 0) {
						_coll19.Add(Companion_Default___.SafeDivisionInt(_218_v104, (_200_v97).Cardinality()), p1)
					}
				}
				return _coll19.ToMap()
			}(), func() _dafny.Map {
				var _coll20 = _dafny.NewMapBuilder()
				_ = _coll20
				for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(355), _dafny.IntOfInt64(267))); ; {
					_compr_20, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _219_v105 _dafny.Int
					_219_v105 = interface{}(_compr_20).(_dafny.Int)
					if ((_dafny.IntOfInt64(355)).Cmp(_219_v105) <= 0) && ((_219_v105).Cmp(_dafny.IntOfInt64(267)) < 0) {
						_coll20.Add((_219_v105).Times(_126_v38), p1)
					}
				}
				return _coll20.ToMap()
			}())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_212_v99), 0))
			_ = _index13
			var _rhs21 _dafny.Int = _210_cf53
			_ = _rhs21
			var _rhs22 bool = (_216_v103).Select((Companion_Default___.SafeIndex((_217_v106).Cardinality(), _dafny.IntOfUint32((_216_v103).Cardinality()))).Uint32()).(bool)
			_ = _rhs22
			var _rhs23 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_215_v102).Fm8(_210_cf53, _211_cf52, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(263))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_220_v92 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_221_i9 _dafny.Int) _dafny.CodePoint {
						return _220_v92
					}
				})(_194_v92))), _214_v101, globalState) {
					return _194_v92
				}
				return _dafny.CodePoint('p')
			})()
			_ = _rhs23
			var _rhs24 *C4 = _215_v102
			_ = _rhs24
			var _lhs8 _dafny.Array = _212_v99
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_212_v99), 0))
			_ = _lhs9
			r1 = _rhs21
			r0 = _rhs22
			_194_v92 = _rhs23
			(_lhs8).ArraySet1(_rhs24, (_lhs9).Int())
			var _222_v107 _dafny.Array
			_ = _222_v107
			var _nw31 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
			_ = _nw31
			_222_v107 = _nw31
			var _223_v108 _dafny.Set
			_ = _223_v108
			_223_v108 = _dafny.SetOf(_216_v103, _216_v103, _dafny.SeqOf((_197_v95).F4()), _216_v103, _216_v103)
			var _224_v109 *C9
			_ = _224_v109
			var _nw32 *C9 = New_C9_()
			_ = _nw32
			_nw32.Ctor__(_210_cf53, _223_v108, _213_v100, _208_cf55, _211_cf52)
			_224_v109 = _nw32
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_222_v107), 0))
			_ = _index14
			(_222_v107).ArraySet1(_224_v109, (_index14).Int())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_222_v107), 0))
			_ = _index15
			(_222_v107).ArraySet1(_224_v109, (_index15).Int())
			_211_cf52 = !(false)
			_211_cf52 = ((_dafny.Zero).Minus((_215_v102).Fm7(globalState))).Cmp(_210_cf53) < 0
		} else {
			var _225___mcc_h5 _dafny.Array = _source7.Get_().(D16_DC32).Cf51
			_ = _225___mcc_h5
			var _226_cf51 _dafny.Array = _225___mcc_h5
			_ = _226_cf51
			_194_v92 = (func() _dafny.CodePoint {
				if (func() bool {
					if p1 {
						return (_197_v95).F4()
					}
					return (_197_v95).F4()
				})() {
					return _194_v92
				}
				return Companion_Default___.Fm1(_126_v38, (_197_v95).F4(), globalState)
			})()
			var _227_v110 _dafny.Array
			_ = _227_v110
			var _nw33 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw33
			_227_v110 = _nw33
			var _228_v111 _dafny.Sequence
			_ = _228_v111
			_228_v111 = _dafny.SeqOf((_dafny.SetOf(_227_v110, _227_v110, _227_v110, _227_v110, _227_v110)).Cardinality())
			r0 = ((_228_v111).Select((Companion_Default___.SafeIndex(_126_v38, _dafny.IntOfUint32((_228_v111).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((_dafny.Zero).Minus((func() _dafny.Int {
				if (_197_v95).F4() {
					return Companion_Default___.Fm3(_126_v38, (_197_v95).F4(), globalState)
				}
				return _126_v38
			})())) > 0
			var _229_v112 _dafny.Array
			_ = _229_v112
			var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(19))
			_ = _nw34
			_229_v112 = _nw34
			var _230_v113 _dafny.Array
			_ = _230_v113
			var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw35
			_230_v113 = _nw35
			var _231_v114 _dafny.Array
			_ = _231_v114
			_231_v114 = _230_v113
			var _232_v115 _dafny.Set
			_ = _232_v115
			_232_v115 = _dafny.SetOf(_231_v114, _231_v114)
			var _233_v116 _dafny.Set
			_ = _233_v116
			_233_v116 = _dafny.SetOf(_dafny.SetOf(_231_v114), _232_v115)
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_229_v112), 0))
			_ = _index16
			(_229_v112).ArraySet1(_233_v116, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_229_v112), 0))
			_ = _index17
			(_229_v112).ArraySet1(_dafny.SetOf(_232_v115), (_index17).Int())
			r1 = (_dafny.Zero).Minus(_126_v38)
		}
		var _234_v117 _dafny.Map
		_ = _234_v117
		_234_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v38, _137_v48)
		_234_v117 = (_234_v117).Update(_126_v38, _137_v48)
		_126_v38 = _126_v38
		r0 = p1
	} else {
		r0 = (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(863))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}(func(_235_i10 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('l')
		}))).Cardinality())).Cmp(_126_v38) >= 0
		var _236_v118 _dafny.Array
		_ = _236_v118
		var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.One)
		_ = _nw36
		_236_v118 = _nw36
		var _237_v119 D22
		_ = _237_v119
		_237_v119 = Companion_D22_.Create_DC45_(_236_v118)
		_237_v119 = Companion_D22_.Create_DC45_(_236_v118)
		var _238_v120 _dafny.Map
		_ = _238_v120
		_238_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v38, p3)
		var _239_v122 _dafny.Array
		_ = _239_v122
		var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
		_ = _nw37
		_239_v122 = _nw37
		var _240_v123 _dafny.Sequence
		_ = _240_v123
		_240_v123 = _dafny.SeqOf(_126_v38, _126_v38)
		var _241_v124 *C1
		_ = _241_v124
		var _nw38 *C1 = New_C1_()
		_ = _nw38
		_nw38.Ctor__((func() _dafny.Set {
			var _coll21 = _dafny.NewBuilder()
			_ = _coll21
			for _iter21 := _dafny.Iterate((_238_v120).Keys().Elements()); ; {
				_compr_21, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _242_v121 _dafny.Int
				_242_v121 = interface{}(_compr_21).(_dafny.Int)
				if (_238_v120).Contains(_242_v121) {
					_coll21.Add((_242_v121).Plus((_dafny.SetOf(true, true)).Cardinality()))
				}
			}
			return _coll21.ToSet()
		}()).Cardinality(), _126_v38, _239_v122, _dafny.IntOfUint32((_240_v123).Cardinality()), p1)
		_241_v124 = _nw38
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_137_v48), 0))
		_ = _index18
		(_137_v48).ArraySet1(_126_v38, (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_137_v48), 0))
		_ = _index19
		(_137_v48).ArraySet1(_126_v38, (_index19).Int())
		_126_v38 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(623))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg32 _dafny.Int) interface{} {
				return coer32(arg32)
			}
		}(func(_243_i11 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))).Cardinality())
	}
	var _244_v125 _dafny.Sequence
	_ = _244_v125
	_244_v125 = _dafny.UnicodeSeqOfUtf8Bytes("rjhatt")
	var _245_v126 _dafny.Array
	_ = _245_v126
	var _nw39 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
	_ = _nw39
	_245_v126 = _nw39
	var _246_v127 _dafny.Sequence
	_ = _246_v127
	_246_v127 = _dafny.SeqOf(_245_v126, _245_v126)
	var _247_v128 D20
	_ = _247_v128
	_247_v128 = Companion_D20_.Create_DC41_(_126_v38)
	var _248_v129 D10
	_ = _248_v129
	_248_v129 = Companion_D10_.Create_DC22_(_244_v125, p1, p3, (_246_v127).Select((Companion_Default___.SafeIndex(_126_v38, _dafny.IntOfUint32((_246_v127).Cardinality()))).Uint32()).(_dafny.Array), (_247_v128).Dtor_cf67())
	var _249_v130 _dafny.Map
	_ = _249_v130
	_249_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v38, p3)
	var _250_v131 _dafny.Map
	_ = _250_v131
	_250_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
	var _251_v132 _dafny.Array
	_ = _251_v132
	var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
	_ = _nw40
	_251_v132 = _nw40
	var _252_v133 *C6
	_ = _252_v133
	var _nw41 *C6 = New_C6_()
	_ = _nw41
	_nw41.Ctor__((func() bool {
		if (_249_v130).Contains(_126_v38) {
			return (_249_v130).Get(_126_v38).(bool)
		}
		return p1
	})(), !((func() bool {
		if (_250_v131).Contains(p1) {
			return (_250_v131).Get(p1).(bool)
		}
		return p1
	})()), _251_v132, Companion_Default___.Fm3(_126_v38, false, globalState))
	_252_v133 = _nw41
	var _253_v134 _dafny.Map
	_ = _253_v134
	_253_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_252_v133, _126_v38)
	r0 = (_dafny.IntOfUint32(((_248_v129).Dtor_cf30()).Cardinality())).Cmp((_253_v134).Cardinality()) >= 0
	var _254_v135 _dafny.CodePoint
	_ = _254_v135
	_254_v135 = _dafny.CodePoint('t')
	r1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ewmfdfes"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(714))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg33 _dafny.Int) interface{} {
			return coer33(arg33)
		}
	}(func(_255_i12 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ewmfdfes")).Cardinality()))).Uint32(), _254_v135), _dafny.UnicodeSeqOfUtf8Bytes("nv"))).Cardinality())
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _256_v0 bool
	_ = _256_v0
	_256_v0 = false
	var _257_v1 _dafny.MultiSet
	_ = _257_v1
	_257_v1 = _dafny.MultiSetOf(_256_v0, _256_v0, _256_v0)
	var _258_globalState *GlobalState
	_ = _258_globalState
	var _nw42 *GlobalState = New_GlobalState_()
	_ = _nw42
	_nw42.Ctor__(_257_v1, _dafny.IntOfInt64(-993), false, _dafny.IntOfInt64(862))
	_258_globalState = _nw42
	var _259_v2 _dafny.Int
	_ = _259_v2
	_259_v2 = _dafny.IntOfInt64(598)
	var _260_v3 _dafny.Array
	_ = _260_v3
	var _nwElement0_9 _dafny.Int = _259_v2
	_ = _nwElement0_9
	var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(14))
	_ = _nw43
	(_nw43).ArraySet1(_nwElement0_9, 0)
	(_nw43).ArraySet1(_259_v2, 1)
	(_nw43).ArraySet1(_259_v2, 2)
	(_nw43).ArraySet1((func() _dafny.Int {
		if _256_v0 {
			return _259_v2
		}
		return _259_v2
	})(), 3)
	(_nw43).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(532))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg34 _dafny.Int) interface{} {
			return coer34(arg34)
		}
	}(func(_261_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))).Cardinality()), _dafny.IntOfInt64(334)), 4)
	(_nw43).ArraySet1(_dafny.IntOfInt64(-930), 5)
	(_nw43).ArraySet1(_259_v2, 6)
	(_nw43).ArraySet1(_259_v2, 7)
	(_nw43).ArraySet1(_259_v2, 8)
	(_nw43).ArraySet1(((Companion_D0_.Create_DC2_(_259_v2)).Dtor_cf1()).Minus((_dafny.Zero).Minus(_259_v2)), 9)
	(_nw43).ArraySet1(_259_v2, 10)
	(_nw43).ArraySet1((_259_v2).Minus(_259_v2), 11)
	(_nw43).ArraySet1(((_dafny.Zero).Minus((_dafny.Zero).Minus(_259_v2))).Times(_dafny.IntOfInt64(743)), 12)
	(_nw43).ArraySet1(_259_v2, 13)
	_260_v3 = _nw43
	_260_v3 = _260_v3
	var _262_v4 _dafny.CodePoint
	_ = _262_v4
	_262_v4 = _dafny.CodePoint('j')
	_262_v4 = _262_v4
	var _263_v5 _dafny.Map
	_ = _263_v5
	_263_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_259_v2).Times(_259_v2), _256_v0)
	_263_v5 = (_263_v5).Update((_259_v2).Minus(_259_v2), _256_v0)
	var _264_v6 _dafny.Set
	_ = _264_v6
	_264_v6 = _dafny.SetOf(_259_v2, _259_v2)
	var _265_v7 _dafny.Set
	_ = _265_v7
	_265_v7 = _dafny.SetOf(_264_v6)
	var _266_v8 _dafny.Sequence
	_ = _266_v8
	_266_v8 = _dafny.UnicodeSeqOfUtf8Bytes("smw")
	var _267_v9 _dafny.Map
	_ = _267_v9
	_267_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_265_v7, (_dafny.IntOfUint32((_266_v8).Cardinality())).Cmp(_259_v2) != 0)
	_256_v0 = (func() bool {
		if (_267_v9).Contains(_265_v7) {
			return (_267_v9).Get(_265_v7).(bool)
		}
		return _256_v0
	})()
	var _268_v10 _dafny.Array
	_ = _268_v10
	var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(2))
	_ = _nw44
	_268_v10 = _nw44
	for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_268_v10), 0))); ; {
		_guard_loop_0, _ok22 := _iter22()
		if !_ok22 {
			break
		}
		var _269_i1 _dafny.Int
		_269_i1 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_269_i1).Sign() != -1) && ((_269_i1).Cmp(_dafny.ArrayLen((_268_v10), 0)) < 0)) {
			(_268_v10).ArraySet1(((_dafny.SetOf(_256_v0)).Difference(_dafny.SetOf(_256_v0, false))).Difference((_dafny.SetOf(_256_v0, true, _256_v0, false)).Intersection(_dafny.SetOf(_256_v0, _256_v0, (Companion_D1_.Create_DC3_(_256_v0)).Dtor_cf2(), _256_v0, _256_v0))), (_269_i1).Int())
		}
	}
	_256_v0 = !(_256_v0) || (_256_v0)
	var _270_v11 _dafny.Array
	_ = _270_v11
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(6)
	_ = _len0_2
	var _nw45 _dafny.Array
	_ = _nw45
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw45 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) bool = (func(_271_v0 bool) func(_dafny.Int) bool {
			return func(_272_i2 _dafny.Int) bool {
				return _271_v0
			}
		})(_256_v0)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw45 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw45).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw45).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_270_v11 = _nw45
	var _273_v12 _dafny.Sequence
	_ = _273_v12
	_273_v12 = _dafny.SeqOf(_270_v11, _270_v11)
	var _274_v13 _dafny.Map
	_ = _274_v13
	_274_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_256_v0), _262_v4)
	var _275_v14 _dafny.Sequence
	_ = _275_v14
	_275_v14 = _dafny.SeqOf(_259_v2)
	var _276_v15 _dafny.Array
	_ = _276_v15
	var _nwElement0_10 bool = !(false)
	_ = _nwElement0_10
	var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(21))
	_ = _nw46
	(_nw46).ArraySet1(_nwElement0_10, 0)
	(_nw46).ArraySet1(Companion_Default___.Fm0(Companion_Default___.Fm1((_dafny.Zero).Minus(_259_v2), _256_v0, _258_globalState), _258_globalState), 1)
	(_nw46).ArraySet1(_256_v0, 2)
	(_nw46).ArraySet1(_256_v0, 3)
	(_nw46).ArraySet1(!(Companion_Default___.Fm0(_dafny.CodePoint('e'), _258_globalState)), 4)
	(_nw46).ArraySet1(_256_v0, 5)
	(_nw46).ArraySet1(_256_v0, 6)
	(_nw46).ArraySet1(!(_256_v0), 7)
	(_nw46).ArraySet1(_256_v0, 8)
	(_nw46).ArraySet1(!(_256_v0), 9)
	(_nw46).ArraySet1(_256_v0, 10)
	(_nw46).ArraySet1(_dafny.Companion_Sequence_.Contains(_273_v12, _270_v11), 11)
	(_nw46).ArraySet1(_256_v0, 12)
	(_nw46).ArraySet1(_256_v0, 13)
	(_nw46).ArraySet1(_256_v0, 14)
	(_nw46).ArraySet1((_259_v2).Cmp((func() _dafny.Int {
		if (_257_v1).Contains(_256_v0) {
			return (_257_v1).Multiplicity(_256_v0)
		}
		return _dafny.IntOfUint32((_266_v8).Cardinality())
	})()) != 0, 15)
	(_nw46).ArraySet1((func() bool {
		if _256_v0 {
			return _256_v0
		}
		return !(_256_v0)
	})(), 16)
	(_nw46).ArraySet1(_256_v0, 17)
	(_nw46).ArraySet1(_256_v0, 18)
	(_nw46).ArraySet1(!(Companion_Default___.Fm0((func() _dafny.CodePoint {
		if (_274_v13).Contains(_256_v0) {
			return (_274_v13).Get(_256_v0).(_dafny.CodePoint)
		}
		return _262_v4
	})(), _258_globalState)), 19)
	(_nw46).ArraySet1((_dafny.IntOfUint32((_266_v8).Cardinality())).Cmp(_dafny.IntOfUint32((_275_v14).Cardinality())) < 0, 20)
	_276_v15 = _nw46
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))
	_ = _index20
	(_276_v15).ArraySet1(false, (_index20).Int())
	var _277_v16 D0
	_ = _277_v16
	_277_v16 = Companion_D0_.Create_DC1_()
	var _278_v17 _dafny.Sequence
	_ = _278_v17
	_278_v17 = _dafny.SeqOf(_277_v16, Companion_Default___.Fm2(_258_globalState))
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))
	_ = _index21
	var _rhs25 bool = !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(Companion_D0_.Create_DC1_()), _dafny.Companion_Sequence_.Concatenate(_278_v17, _278_v17)))
	_ = _rhs25
	var _rhs26 bool = Companion_Default___.Fm0(_262_v4, _258_globalState)
	_ = _rhs26
	var _lhs10 _dafny.Array = _276_v15
	_ = _lhs10
	var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))
	_ = _lhs11
	_256_v0 = _rhs25
	(_lhs10).ArraySet1(_rhs26, (_lhs11).Int())
	var _279_v18 _dafny.MultiSet
	_ = _279_v18
	_279_v18 = _dafny.MultiSetOf(_259_v2)
	var _280_v19 _dafny.Map
	_ = _280_v19
	_280_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_279_v18, _266_v8)
	var _281_v20 bool
	_ = _281_v20
	var _282_v21 _dafny.Int
	_ = _282_v21
	var _out0 bool
	_ = _out0
	var _out1 _dafny.Int
	_ = _out1
	_out0, _out1 = Companion_Default___.M0(_280_v19, (func() bool {
		if (_276_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))).Int()).(bool) {
			return !(_256_v0)
		}
		return (_276_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))).Int()).(bool)
	})(), (_279_v18).Update(_259_v2, Companion_Default___.Abs(_259_v2)), (_276_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))).Int()).(bool), _258_globalState)
	_281_v20 = _out0
	_282_v21 = _out1
	var _283_v22 _dafny.Array
	_ = _283_v22
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(24)
	_ = _len0_3
	var _nw47 _dafny.Array
	_ = _nw47
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw47 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) _dafny.Sequence = (func(_284_v14 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_285_i3 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_284_v14, _284_v14)
			}
		})(_275_v14)
		_ = _init3
		var _element0_3 = _init3(_dafny.Zero)
		_ = _element0_3
		_nw47 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
		(_nw47).ArraySet1(_element0_3, 0)
		var _nativeLen0_3 = (_len0_3).Int()
		_ = _nativeLen0_3
		for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
			(_nw47).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
		}
	}
	_283_v22 = _nw47
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_283_v22), 0))
	_ = _index22
	(_283_v22).ArraySet1(_275_v14, (_index22).Int())
	var _286_v23 D0
	_ = _286_v23
	_286_v23 = Companion_D0_.Create_DC2_(_259_v2)
	var _pat_let_tv1 = _277_v16
	_ = _pat_let_tv1
	var _pat_let_tv2 = _277_v16
	_ = _pat_let_tv2
	var _pat_let_tv3 = _277_v16
	_ = _pat_let_tv3
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))
	_ = _index23
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_283_v22), 0))
	_ = _index24
	var _rhs27 D0 = func(_source8 D0) D0 {
		if _source8.Is_DC1() {
			return _pat_let_tv1
		} else if _source8.Is_DC2() {
			var _287___mcc_h0 _dafny.Int = _source8.Get_().(D0_DC2).Cf1
			_ = _287___mcc_h0
			var _288_cf1 _dafny.Int = _287___mcc_h0
			_ = _288_cf1
			return _pat_let_tv2
		} else {
			var _289___mcc_h1 _dafny.Int = _source8.Get_().(D0_DC0).Cf0
			_ = _289___mcc_h1
			var _290_cf0 _dafny.Int = _289___mcc_h1
			_ = _290_cf0
			return _pat_let_tv3
		}
	}(_286_v23)
	_ = _rhs27
	var _rhs28 bool = true
	_ = _rhs28
	var _rhs29 bool = Companion_Default___.Fm0(_262_v4, _258_globalState)
	_ = _rhs29
	var _rhs30 _dafny.Sequence = _275_v14
	_ = _rhs30
	var _rhs31 _dafny.Int = _259_v2
	_ = _rhs31
	var _lhs12 _dafny.Array = _276_v15
	_ = _lhs12
	var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))
	_ = _lhs13
	var _lhs14 _dafny.Array = _283_v22
	_ = _lhs14
	var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_283_v22), 0))
	_ = _lhs15
	_277_v16 = _rhs27
	(_lhs12).ArraySet1(_rhs28, (_lhs13).Int())
	_281_v20 = _rhs29
	(_lhs14).ArraySet1(_rhs30, (_lhs15).Int())
	_259_v2 = _rhs31
	_286_v23 = _286_v23
	if (_276_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))).Int()).(bool) {
		var _291_v24 _dafny.Map
		_ = _291_v24
		_291_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_259_v2, _260_v3)
		_260_v3 = (func() _dafny.Array {
			if (_291_v24).Contains(_259_v2) {
				return (_291_v24).Get(_259_v2).(_dafny.Array)
			}
			return _260_v3
		})()
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_260_v3), 0))
		_ = _index25
		(_260_v3).ArraySet1((_282_v21).Times(Companion_Default___.Fm3(_259_v2, (_276_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))).Int()).(bool), _258_globalState)), (_index25).Int())
		var _292_v25 _dafny.Sequence
		_ = _292_v25
		_292_v25 = _dafny.SeqOf(_279_v18, (_279_v18).Union(_dafny.MultiSetOf(_259_v2, _dafny.IntOfInt64(673))))
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_260_v3), 0))
		_ = _index26
		var _rhs32 bool = !(_256_v0)
		_ = _rhs32
		var _rhs33 _dafny.Int = Companion_Default___.SafeDivisionInt(_282_v21, _259_v2)
		_ = _rhs33
		var _rhs34 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_292_v25, Companion_Default___.Fm4(false, true, _258_globalState))
		_ = _rhs34
		var _rhs35 D0 = _286_v23
		_ = _rhs35
		var _lhs16 _dafny.Array = _260_v3
		_ = _lhs16
		var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_260_v3), 0))
		_ = _lhs17
		_281_v20 = _rhs32
		(_lhs16).ArraySet1(_rhs33, (_lhs17).Int())
		_292_v25 = _rhs34
		_286_v23 = _rhs35
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_260_v3), 0))
		_ = _index27
		(_260_v3).ArraySet1(_dafny.IntOfInt64(902), (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))
		_ = _index28
		(_276_v15).ArraySet1(_281_v20, (_index28).Int())
		var _293_v26 D2
		_ = _293_v26
		_293_v26 = Companion_D2_.Create_DC5_(_276_v15)
		_270_v11 = (_293_v26).Dtor_cf5()
	} else {
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_260_v3), 0))
		_ = _index29
		(_260_v3).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if _281_v20 {
				return _266_v8
			}
			return _266_v8
		})()).Cardinality()), (_index29).Int())
		var _294_v27 _dafny.Array
		_ = _294_v27
		var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
		_ = _nw48
		_294_v27 = _nw48
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_294_v27), 0))
		_ = _index30
		(_294_v27).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("nwenxedn"), (_index30).Int())
		var _295_v28 _dafny.Sequence
		_ = _295_v28
		_295_v28 = _dafny.SeqOf(_266_v8, _266_v8, _dafny.UnicodeSeqOfUtf8Bytes("qystgqh"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(635))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}((func(_296_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_297_i4 _dafny.Int) _dafny.CodePoint {
				return _296_v4
			}
		})(_262_v4))), _dafny.Companion_Sequence_.Concatenate(_266_v8, _266_v8))
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_260_v3), 0))
		_ = _index31
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_294_v27), 0))
		_ = _index32
		var _rhs36 _dafny.Int = _dafny.IntOfUint32((_295_v28).Cardinality())
		_ = _rhs36
		var _rhs37 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(912))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg36 _dafny.Int) interface{} {
				return coer36(arg36)
			}
		}(func(_298_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		}))
		_ = _rhs37
		var _lhs18 _dafny.Array = _260_v3
		_ = _lhs18
		var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_260_v3), 0))
		_ = _lhs19
		var _lhs20 _dafny.Array = _294_v27
		_ = _lhs20
		var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_294_v27), 0))
		_ = _lhs21
		(_lhs18).ArraySet1(_rhs36, (_lhs19).Int())
		(_lhs20).ArraySet1(_rhs37, (_lhs21).Int())
		var _299_v30 _dafny.Sequence
		_ = _299_v30
		_299_v30 = _dafny.SeqOf(_279_v18)
		var _300_v31 _dafny.Map
		_ = _300_v31
		_300_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _280_v19)
		var _301_v32 bool
		_ = _301_v32
		var _302_v33 _dafny.Int
		_ = _302_v33
		var _out2 bool
		_ = _out2
		var _out3 _dafny.Int
		_ = _out3
		_out2, _out3 = Companion_Default___.M0((func() _dafny.Map {
			var _coll22 = _dafny.NewMapBuilder()
			_ = _coll22
			for _iter23 := _dafny.Iterate((_299_v30).Elements()); ; {
				_compr_22, _ok23 := _iter23()
				if !_ok23 {
					break
				}
				var _303_v29 _dafny.MultiSet
				_303_v29 = interface{}(_compr_22).(_dafny.MultiSet)
				if _dafny.Companion_Sequence_.Contains(_299_v30, _303_v29) {
					_coll22.Add(_303_v29, _266_v8)
				}
			}
			return _coll22.ToMap()
		}()).Merge((func() _dafny.Map {
			if (_300_v31).Contains(_281_v20) {
				return (_300_v31).Get(_281_v20).(_dafny.Map)
			}
			return _280_v19
		})()), _256_v0, _279_v18, !(_256_v0), _258_globalState)
		_301_v32 = _out2
		_302_v33 = _out3
		_256_v0 = (_276_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))).Int()).(bool)
		var _304_v34 _dafny.Set
		_ = _304_v34
		_304_v34 = _dafny.SetOf(_266_v8)
		var _305_v35 _dafny.Map
		_ = _305_v35
		_305_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_256_v0, _301_v32)
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_260_v3), 0))
		_ = _index33
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_260_v3), 0))
		_ = _index34
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_260_v3), 0))
		_ = _index35
		var _rhs38 _dafny.Int = Companion_Default___.Fm3(((_304_v34).Cardinality()).Plus(_302_v33), !((_276_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))).Int()).(bool)), _258_globalState)
		_ = _rhs38
		var _rhs39 _dafny.MultiSet = _279_v18
		_ = _rhs39
		var _rhs40 _dafny.Int = _282_v21
		_ = _rhs40
		var _rhs41 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_259_v2, (_dafny.Zero).Minus((func() _dafny.Int {
			if _281_v20 {
				return _302_v33
			}
			return (_305_v35).Cardinality()
		})())))
		_ = _rhs41
		var _rhs42 _dafny.Int = _302_v33
		_ = _rhs42
		var _lhs22 _dafny.Array = _260_v3
		_ = _lhs22
		var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_260_v3), 0))
		_ = _lhs23
		var _lhs24 _dafny.Array = _260_v3
		_ = _lhs24
		var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_260_v3), 0))
		_ = _lhs25
		var _lhs26 _dafny.Array = _260_v3
		_ = _lhs26
		var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_260_v3), 0))
		_ = _lhs27
		_282_v21 = _rhs38
		_279_v18 = _rhs39
		(_lhs22).ArraySet1(_rhs40, (_lhs23).Int())
		(_lhs24).ArraySet1(_rhs41, (_lhs25).Int())
		(_lhs26).ArraySet1(_rhs42, (_lhs27).Int())
		_256_v0 = !(_301_v32) || (_301_v32)
	}
	var _306_v36 _dafny.Map
	_ = _306_v36
	_306_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_277_v16, _282_v21)
	var _307_v40 _dafny.Set
	_ = _307_v40
	_307_v40 = _dafny.SetOf(Companion_Default___.Fm0(_262_v4, _258_globalState))
	var _308_v41 _dafny.Array
	_ = _308_v41
	var _nwElement0_11 _dafny.Map = _306_v36
	_ = _nwElement0_11
	var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(12))
	_ = _nw49
	(_nw49).ArraySet1(_nwElement0_11, 0)
	(_nw49).ArraySet1((_306_v36).Merge(_306_v36), 1)
	(_nw49).ArraySet1(_306_v36, 2)
	(_nw49).ArraySet1(_306_v36, 3)
	(_nw49).ArraySet1(func() _dafny.Map {
		var _coll23 = _dafny.NewMapBuilder()
		_ = _coll23
		for _iter24 := _dafny.Iterate((_306_v36).Keys().Elements()); ; {
			_compr_23, _ok24 := _iter24()
			if !_ok24 {
				break
			}
			var _309_v37 D0
			_309_v37 = interface{}(_compr_23).(D0)
			if (_306_v36).Contains(_309_v37) {
				_coll23.Add(_309_v37, (_275_v14).Select((Companion_Default___.SafeIndex(_259_v2, _dafny.IntOfUint32((_275_v14).Cardinality()))).Uint32()).(_dafny.Int))
			}
		}
		return _coll23.ToMap()
	}(), 4)
	(_nw49).ArraySet1(_306_v36, 5)
	(_nw49).ArraySet1(_306_v36, 6)
	(_nw49).ArraySet1(_306_v36, 7)
	(_nw49).ArraySet1((func() _dafny.Map {
		var _coll24 = _dafny.NewMapBuilder()
		_ = _coll24
		for _iter25 := _dafny.Iterate((_dafny.SeqOf(_277_v16, _277_v16)).Elements()); ; {
			_compr_24, _ok25 := _iter25()
			if !_ok25 {
				break
			}
			var _310_v38 D0
			_310_v38 = interface{}(_compr_24).(D0)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_277_v16, _277_v16), _310_v38) {
				_coll24.Add(_310_v38, (func() _dafny.Map {
					var _coll25 = _dafny.NewMapBuilder()
					_ = _coll25
					for _iter26 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D2_.Create_DC7_(_259_v2)))).Elements()); ; {
						_compr_25, _ok26 := _iter26()
						if !_ok26 {
							break
						}
						var _311_v39 D2
						_311_v39 = interface{}(_compr_25).(D2)
						if (_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D2_.Create_DC7_(_259_v2)))).Contains(_311_v39) {
							_coll25.Add(_311_v39, _282_v21)
						}
					}
					return _coll25.ToMap()
				}()).Cardinality())
			}
		}
		return _coll24.ToMap()
	}()).Update(Companion_D0_.Create_DC1_(), Companion_Default___.Fm3((_307_v40).Cardinality(), _256_v0, _258_globalState)), 8)
	(_nw49).ArraySet1(_306_v36, 9)
	(_nw49).ArraySet1(Companion_Default___.Fm5(_282_v21, (_276_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))).Int()).(bool), _258_globalState), 10)
	(_nw49).ArraySet1(_306_v36, 11)
	_308_v41 = _nw49
	_308_v41 = (func() _dafny.Array {
		if !(_256_v0) {
			return _308_v41
		}
		return _308_v41
	})()
	var _312_v42 _dafny.MultiSet
	_ = _312_v42
	_312_v42 = _dafny.MultiSetOf(_262_v4)
	var _313_v43 bool
	_ = _313_v43
	var _314_v44 _dafny.Int
	_ = _314_v44
	var _out4 bool
	_ = _out4
	var _out5 _dafny.Int
	_ = _out5
	_out4, _out5 = Companion_Default___.M0(_280_v19, false, _279_v18, (_312_v42).Contains(_dafny.CodePoint('n')), _258_globalState)
	_313_v43 = _out4
	_314_v44 = _out5
	var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))
	_ = _index36
	(_260_v3).ArraySet1(_314_v44, (_index36).Int())
	var _315_v45 _dafny.Map
	_ = _315_v45
	_315_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-750))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg37 _dafny.Int) interface{} {
			return coer37(arg37)
		}
	}((func(_316_v21 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_317_i6 _dafny.Int) _dafny.Int {
			return _316_v21
		}
	})(_282_v21))))
	var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))
	_ = _index37
	(_260_v3).ArraySet1((_dafny.Zero).Minus((((func() _dafny.Map {
		if _281_v20 {
			return _315_v45
		}
		return _315_v45
	})()).Cardinality()).Times(_259_v2)), (_index37).Int())
	var _hi1 _dafny.Int = _282_v21
	_ = _hi1
	for _318_i7 := _314_v44; _318_i7.Cmp(_hi1) < 0; _318_i7 = _318_i7.Plus(_dafny.One) {
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))
		_ = _index38
		(_260_v3).ArraySet1((_259_v2).Plus((_260_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))).Int()).(_dafny.Int)), (_index38).Int())
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))
		_ = _index39
		(_260_v3).ArraySet1(_259_v2, (_index39).Int())
		var _319_v46 D19
		_ = _319_v46
		_319_v46 = Companion_D19_.Create_DC38_(_281_v20, _313_v43)
		var _320_v47 _dafny.Sequence
		_ = _320_v47
		_320_v47 = _dafny.SeqOf(false, false)
		var _321_v48 _dafny.Set
		_ = _321_v48
		_321_v48 = _dafny.SetOf(_320_v47, _320_v47, _320_v47)
		var _322_v49 _dafny.Map
		_ = _322_v49
		_322_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_256_v0, _276_v15)
		var _323_v50 D10
		_ = _323_v50
		_323_v50 = Companion_D10_.Create_DC22_(_266_v8, _256_v0, _281_v20, _276_v15, _dafny.IntOfInt64(867))
		var _324_v51 _dafny.Array
		_ = _324_v51
		var _nwElement0_12 _dafny.Array = _270_v11
		_ = _nwElement0_12
		var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(18))
		_ = _nw50
		(_nw50).ArraySet1(_nwElement0_12, 0)
		(_nw50).ArraySet1(_270_v11, 1)
		(_nw50).ArraySet1(_276_v15, 2)
		(_nw50).ArraySet1(_276_v15, 3)
		(_nw50).ArraySet1(_276_v15, 4)
		(_nw50).ArraySet1(_270_v11, 5)
		(_nw50).ArraySet1(_270_v11, 6)
		(_nw50).ArraySet1(_276_v15, 7)
		(_nw50).ArraySet1(_270_v11, 8)
		(_nw50).ArraySet1(_270_v11, 9)
		(_nw50).ArraySet1(_276_v15, 10)
		(_nw50).ArraySet1(_270_v11, 11)
		(_nw50).ArraySet1((func() _dafny.Array {
			if (_322_v49).Contains(_256_v0) {
				return (_322_v49).Get(_256_v0).(_dafny.Array)
			}
			return _270_v11
		})(), 12)
		(_nw50).ArraySet1(_276_v15, 13)
		(_nw50).ArraySet1(_270_v11, 14)
		(_nw50).ArraySet1(_276_v15, 15)
		(_nw50).ArraySet1(_270_v11, 16)
		(_nw50).ArraySet1((_323_v50).Dtor_cf33(), 17)
		_324_v51 = _nw50
		var _325_v52 *C9
		_ = _325_v52
		var _nw51 *C9 = New_C9_()
		_ = _nw51
		_nw51.Ctor__(Companion_Default___.Fm3(_318_i7, (_319_v46).Dtor_cf64(), _258_globalState), _321_v48, _324_v51, _259_v2, Companion_Default___.Fm0(_262_v4, _258_globalState))
		_325_v52 = _nw51
		_260_v3 = _260_v3
	}
	if !(!(_281_v20)) {
		var _326_v53 _dafny.Array
		_ = _326_v53
		var _nwElement0_13 _dafny.Array = _270_v11
		_ = _nwElement0_13
		var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(10))
		_ = _nw52
		(_nw52).ArraySet1(_nwElement0_13, 0)
		(_nw52).ArraySet1(_270_v11, 1)
		(_nw52).ArraySet1(_276_v15, 2)
		(_nw52).ArraySet1(_276_v15, 3)
		(_nw52).ArraySet1(_270_v11, 4)
		(_nw52).ArraySet1(_270_v11, 5)
		(_nw52).ArraySet1(_270_v11, 6)
		(_nw52).ArraySet1(_270_v11, 7)
		(_nw52).ArraySet1(_276_v15, 8)
		(_nw52).ArraySet1(_276_v15, 9)
		_326_v53 = _nw52
		var _327_v54 *C5
		_ = _327_v54
		var _nw53 *C5 = New_C5_()
		_ = _nw53
		_nw53.Ctor__(_281_v20, _259_v2, _326_v53, _dafny.IntOfUint32((_266_v8).Cardinality()), _313_v43)
		_327_v54 = _nw53
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))
		_ = _index40
		(_260_v3).ArraySet1(_dafny.IntOfUint32(((Companion_D26_.Create_DC53_(_dafny.SeqOf(_327_v54, _327_v54, _327_v54))).Dtor_cf89()).Cardinality()), (_index40).Int())
		var _328_v55 _dafny.Set
		_ = _328_v55
		_328_v55 = _dafny.SetOf(_262_v4, _262_v4)
		var _329_v56 D6
		_ = _329_v56
		_329_v56 = Companion_D6_.Create_DC16_((_327_v54).F13(), (func() bool {
			if (_263_v5).Contains((_328_v55).Cardinality()) {
				return (_263_v5).Get((_328_v55).Cardinality()).(bool)
			}
			return _313_v43
		})(), (_283_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_283_v22), 0))).Int()).(_dafny.Sequence))
		var _330_v57 _dafny.Sequence
		_ = _330_v57
		_330_v57 = _dafny.SeqOf(true, (_327_v54).Fm8(_282_v21, _256_v0, _266_v8, (_329_v56).Dtor_cf22(), _258_globalState))
		_262_v4 = (_266_v8).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_330_v57).Cardinality()), _259_v2), _dafny.IntOfUint32((_266_v8).Cardinality()))).Uint32()).(_dafny.CodePoint)
		var _331_v59 _dafny.Sequence
		_ = _331_v59
		_331_v59 = _dafny.SeqOf(_326_v53)
		var _332_v60 T1
		_ = _332_v60
		var _nw54 *C9 = New_C9_()
		_ = _nw54
		_nw54.Ctor__(_259_v2, func() _dafny.Set {
			var _coll26 = _dafny.NewBuilder()
			_ = _coll26
			for _iter27 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-88))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg38 _dafny.Int) interface{} {
					return coer38(arg38)
				}
			}((func(_333_v57 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_334_i8 _dafny.Int) _dafny.Sequence {
					return _333_v57
				}
			})(_330_v57)))).Elements()); ; {
				_compr_26, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _335_v58 _dafny.Sequence
				_335_v58 = interface{}(_compr_26).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-88))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_336_v57 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_334_i8 _dafny.Int) _dafny.Sequence {
						return _336_v57
					}
				})(_330_v57))), _335_v58) {
					_coll26.Add(_335_v58)
				}
			}
			return _coll26.ToSet()
		}(), (_331_v59).Select((Companion_Default___.SafeIndex(_259_v2, _dafny.IntOfUint32((_331_v59).Cardinality()))).Uint32()).(_dafny.Array), _314_v44, (_327_v54).F13())
		_332_v60 = _nw54
		var _337_v61 _dafny.Set
		_ = _337_v61
		_337_v61 = _dafny.SetOf(_332_v60)
		_337_v61 = _337_v61
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))
		_ = _index41
		(_260_v3).ArraySet1((_332_v60).Fm9((_327_v54).F13(), (_327_v54).Fm8((_260_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))).Int()).(_dafny.Int), _313_v43, _266_v8, _dafny.SeqOf(_327_v54.F14, (_260_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))).Int()).(_dafny.Int)), _258_globalState), _282_v21, _258_globalState), (_index41).Int())
		_313_v43 = (_327_v54).F13()
	} else {
		var _338_v63 _dafny.MultiSet
		_ = _338_v63
		_338_v63 = _dafny.MultiSetOf(_279_v18)
		var _339_v64 bool
		_ = _339_v64
		var _340_v65 _dafny.Int
		_ = _340_v65
		var _out6 bool
		_ = _out6
		var _out7 _dafny.Int
		_ = _out7
		_out6, _out7 = Companion_Default___.M0(func() _dafny.Map {
			var _coll27 = _dafny.NewMapBuilder()
			_ = _coll27
			for _iter28 := _dafny.Iterate((_338_v63).Elements()); ; {
				_compr_27, _ok28 := _iter28()
				if !_ok28 {
					break
				}
				var _341_v62 _dafny.MultiSet
				_341_v62 = interface{}(_compr_27).(_dafny.MultiSet)
				if (_338_v63).Contains(_341_v62) {
					_coll27.Add(_341_v62, _266_v8)
				}
			}
			return _coll27.ToMap()
		}(), _281_v20, _dafny.MultiSetFromSeq(_275_v14), (_281_v20) || (_256_v0), _258_globalState)
		_339_v64 = _out6
		_340_v65 = _out7
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))
		_ = _index42
		(_260_v3).ArraySet1(((_279_v18).Cardinality()).Times((_264_v6).Cardinality()), (_index42).Int())
		var _342_v66 D16
		_ = _342_v66
		_342_v66 = Companion_D16_.Create_DC33_(_313_v43, _259_v2, _257_v1, _314_v44, _307_v40)
		if ((_dafny.IntOfInt64(980)).Cmp((_260_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))).Int()).(_dafny.Int)) != 0) && ((_342_v66).Dtor_cf52()) {
			var _343_v67 bool
			_ = _343_v67
			var _344_v68 _dafny.Int
			_ = _344_v68
			var _out8 bool
			_ = _out8
			var _out9 _dafny.Int
			_ = _out9
			_out8, _out9 = Companion_Default___.M0(_280_v19, _313_v43, _279_v18, _281_v20, _258_globalState)
			_343_v67 = _out8
			_344_v68 = _out9
			_343_v67 = _339_v64
			var _345_v69 _dafny.Map
			_ = _345_v69
			_345_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(320), _279_v18)
			var _346_v70 _dafny.Sequence
			_ = _346_v70
			_346_v70 = _dafny.SeqOf(_256_v0, !(_345_v69).Equals(_345_v69))
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))
			_ = _index43
			(_276_v15).ArraySet1((_346_v70).Select((Companion_Default___.SafeIndex(_259_v2, _dafny.IntOfUint32((_346_v70).Cardinality()))).Uint32()).(bool), (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))
			_ = _index44
			(_276_v15).ArraySet1(_281_v20, (_index44).Int())
			_256_v0 = (func() bool {
				if (func() bool {
					if _339_v64 {
						return _256_v0
					}
					return _339_v64
				})() {
					return _256_v0
				}
				return (_276_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))).Int()).(bool)
			})()
		} else {
			var _347_v71 _dafny.Sequence
			_ = _347_v71
			_347_v71 = _dafny.SeqOf(_339_v64, _313_v43)
			_256_v0 = (_347_v71).Select((Companion_Default___.SafeIndex((_260_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_347_v71).Cardinality()))).Uint32()).(bool)
			_266_v8 = _dafny.UnicodeSeqOfUtf8Bytes("becsmvx")
			var _348_v72 bool
			_ = _348_v72
			var _349_v73 _dafny.Int
			_ = _349_v73
			var _out10 bool
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out10, _out11 = Companion_Default___.M0(_280_v19, (_259_v2).Cmp(_259_v2) <= 0, _279_v18, _313_v43, _258_globalState)
			_348_v72 = _out10
			_349_v73 = _out11
			var _350_v74 D21
			_ = _350_v74
			_350_v74 = Companion_D21_.Create_DC44_(false)
			var _pat_let_tv4 = _313_v43
			_ = _pat_let_tv4
			_350_v74 = func(_pat_let3_0 D21) D21 {
				return func(_351_dt__update__tmp_h0 D21) D21 {
					return func(_pat_let4_0 bool) D21 {
						return func(_352_dt__update_hcf72_h0 bool) D21 {
							return Companion_D21_.Create_DC44_(_352_dt__update_hcf72_h0)
						}(_pat_let4_0)
					}(_pat_let_tv4)
				}(_pat_let3_0)
			}(_350_v74)
			_348_v72 = !((_279_v18).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(-977), _dafny.IntOfInt64(547), (func() _dafny.Int {
				if (_257_v1).Contains(_281_v20) {
					return (_257_v1).Multiplicity(_281_v20)
				}
				return _314_v44
			})()))).Contains(Companion_Default___.Fm3((_260_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))).Int()).(_dafny.Int), _348_v72, _258_globalState))
		}
		var _353_v75 _dafny.MultiSet
		_ = _353_v75
		_353_v75 = _dafny.MultiSetOf(_dafny.SeqOf(_259_v2, _340_v65), _275_v14)
		var _354_v76 _dafny.MultiSet
		_ = _354_v76
		_354_v76 = _353_v75
		var _355_v77 _dafny.Sequence
		_ = _355_v77
		_355_v77 = _dafny.SeqOf(_281_v20)
		var _356_v78 _dafny.Sequence
		_ = _356_v78
		_356_v78 = _dafny.SeqOf(_355_v77)
		var _357_v79 D12
		_ = _357_v79
		_357_v79 = Companion_D12_.Create_DC27_(_356_v78)
		var _358_v80 _dafny.Map
		_ = _358_v80
		_358_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_266_v8, _357_v79)
		var _359_v81 _dafny.Sequence
		_ = _359_v81
		_359_v81 = _dafny.SeqOf(_358_v80, _358_v80)
		var _360_v82 _dafny.Array
		_ = _360_v82
		var _nwElement0_14 _dafny.MultiSet = _354_v76
		_ = _nwElement0_14
		var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(25))
		_ = _nw55
		(_nw55).ArraySet1(_nwElement0_14, 0)
		(_nw55).ArraySet1(_354_v76, 1)
		(_nw55).ArraySet1(_353_v75, 2)
		(_nw55).ArraySet1(_354_v76, 3)
		(_nw55).ArraySet1(_354_v76, 4)
		(_nw55).ArraySet1(_353_v75, 5)
		(_nw55).ArraySet1(Companion_Default___.Fm72((_359_v81).Select((Companion_Default___.SafeIndex(_314_v44, _dafny.IntOfUint32((_359_v81).Cardinality()))).Uint32()).(_dafny.Map), _279_v18, _258_globalState), 6)
		(_nw55).ArraySet1(_354_v76, 7)
		(_nw55).ArraySet1(_354_v76, 8)
		(_nw55).ArraySet1(Companion_Default___.Fm73(_258_globalState), 9)
		(_nw55).ArraySet1(_353_v75, 10)
		(_nw55).ArraySet1(_354_v76, 11)
		(_nw55).ArraySet1(_354_v76, 12)
		(_nw55).ArraySet1(_354_v76, 13)
		(_nw55).ArraySet1(_354_v76, 14)
		(_nw55).ArraySet1(_354_v76, 15)
		(_nw55).ArraySet1(_354_v76, 16)
		(_nw55).ArraySet1(_354_v76, 17)
		(_nw55).ArraySet1(_354_v76, 18)
		(_nw55).ArraySet1(_354_v76, 19)
		(_nw55).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(790))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}((func(_361_v14 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_362_i9 _dafny.Int) _dafny.Sequence {
				return _361_v14
			}
		})(_275_v14)))), 20)
		(_nw55).ArraySet1(_354_v76, 21)
		(_nw55).ArraySet1(_354_v76, 22)
		(_nw55).ArraySet1(_354_v76, 23)
		(_nw55).ArraySet1((_353_v75).Update((_283_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_283_v22), 0))).Int()).(_dafny.Sequence), Companion_Default___.Abs((_260_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_260_v3), 0))).Int()).(_dafny.Int))), 24)
		_360_v82 = _nw55
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_360_v82), 0))
		_ = _index45
		(_360_v82).ArraySet1(_354_v76, (_index45).Int())
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_360_v82), 0))
		_ = _index46
		(_360_v82).ArraySet1(_354_v76, (_index46).Int())
		var _rhs43 _dafny.CodePoint = _262_v4
		_ = _rhs43
		var _rhs44 bool = !((_276_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_276_v15), 0))).Int()).(bool))
		_ = _rhs44
		var _rhs45 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_266_v8, _266_v8)).Cardinality()), _282_v21)
		_ = _rhs45
		_262_v4 = _rhs43
		_281_v20 = _rhs44
		_259_v2 = _rhs45
	}
	_dafny.Print(_256_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_257_v1).Equals(_dafny.MultiSetOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_258_globalState).F0()).Equals(_dafny.MultiSetOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_259_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v3).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_262_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_263_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(357604), false).UpdateUnsafe(_dafny.Zero, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_264_v6).Equals(_dafny.SetOf(_dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v7).Equals(_dafny.SetOf(_dafny.SetOf(_dafny.IntOfInt64(598)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_266_v8.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_267_v9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.SetOf(_dafny.IntOfInt64(598))), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_268_v10).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_268_v10).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v11).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_273_v12).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_v13).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_275_v14, _dafny.SeqOf(_dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v15).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_278_v17, _dafny.SeqOf(Companion_D0_.Create_DC1_(), Companion_D0_.Create_DC1_())))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_279_v18).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_280_v19).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(598)), _dafny.UnicodeSeqOfUtf8Bytes("smw"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_281_v20)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_282_v21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_283_v22).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v23).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v36).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_307_v40).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_308_v41).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_308_v41).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_308_v41).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_308_v41).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_308_v41).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(598))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_308_v41).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_308_v41).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_308_v41).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_308_v41).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_308_v41).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_308_v41).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_308_v41).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(), _dafny.IntOfInt64(10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_312_v42).Equals(_dafny.MultiSetOf(_dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_313_v43)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_314_v44)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_315_v45).Cardinality())
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_() D0 {
	return D0{D0_DC1{}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf1 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf1 _dafny.Int) D0 {
	return D0{D0_DC2{Cf1}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
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
	return Companion_D0_.Create_DC1_()
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf1
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
			return "D0.DC1"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf1) + ")"
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
			_, ok := other.Get_().(D0_DC1)
			return ok
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0
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

type D1_DC4 struct {
	Cf3 _dafny.Int
	Cf4 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf3 _dafny.Int, Cf4 _dafny.Int) D1 {
	return D1{D1_DC4{Cf3, Cf4}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf2 bool
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf2 bool) D1 {
	return D1{D1_DC3{Cf2}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.Zero, _dafny.Zero)
}

func (_this D1) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) Dtor_cf2() bool {
	return _this.Get_().(D1_DC3).Cf2
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
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
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4.Cmp(data2.Cf4) == 0
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf2 == data2.Cf2
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
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_() D2 {
	return D2{D2_DC6{}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf6 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf6 _dafny.Int) D2 {
	return D2{D2_DC7{Cf6}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC5 struct {
	Cf5 _dafny.Array
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf5 _dafny.Array) D2 {
	return D2{D2_DC5{Cf5}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_()
}

func (_this D2) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf6
}

func (_this D2) Dtor_cf5() _dafny.Array {
	return _this.Get_().(D2_DC5).Cf5
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf6) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf5) + ")"
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
			_, ok := other.Get_().(D2_DC6)
			return ok
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf5 == data2.Cf5
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

type D3_DC9 struct {
	Cf8  bool
	Cf9  _dafny.Int
	Cf10 bool
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf8 bool, Cf9 _dafny.Int, Cf10 bool) D3 {
	return D3{D3_DC9{Cf8, Cf9, Cf10}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC10 struct {
	Cf11 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf11 _dafny.Int) D3 {
	return D3{D3_DC10{Cf11}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC8 struct {
	Cf7 _dafny.Sequence
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf7 _dafny.Sequence) D3 {
	return D3{D3_DC8{Cf7}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC11 struct {
	Cf12 D3
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf12 D3) D3 {
	return D3{D3_DC11{Cf12}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(false, _dafny.Zero, false)
}

func (_this D3) Dtor_cf8() bool {
	return _this.Get_().(D3_DC9).Cf8
}

func (_this D3) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf9
}

func (_this D3) Dtor_cf10() bool {
	return _this.Get_().(D3_DC9).Cf10
}

func (_this D3) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf11
}

func (_this D3) Dtor_cf7() _dafny.Sequence {
	return _this.Get_().(D3_DC8).Cf7
}

func (_this D3) Dtor_cf12() D3 {
	return _this.Get_().(D3_DC11).Cf12
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf7) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf7.Equals(data2.Cf7)
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
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

type D4_DC13 struct {
	Cf14 _dafny.Int
	Cf15 bool
	Cf16 _dafny.Int
	Cf17 _dafny.Sequence
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf14 _dafny.Int, Cf15 bool, Cf16 _dafny.Int, Cf17 _dafny.Sequence) D4 {
	return D4{D4_DC13{Cf14, Cf15, Cf16, Cf17}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC12 struct {
	Cf13 _dafny.CodePoint
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf13 _dafny.CodePoint) D4 {
	return D4{D4_DC12{Cf13}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_(_dafny.Zero, false, _dafny.Zero, _dafny.EmptySeq)
}

func (_this D4) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf14
}

func (_this D4) Dtor_cf15() bool {
	return _this.Get_().(D4_DC13).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf16
}

func (_this D4) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D4_DC13).Cf17
}

func (_this D4) Dtor_cf13() _dafny.CodePoint {
	return _this.Get_().(D4_DC12).Cf13
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15 == data2.Cf15 && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17.Equals(data2.Cf17)
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf13 == data2.Cf13
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
	Cf18 _dafny.MultiSet
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf18 _dafny.MultiSet) D5 {
	return D5{D5_DC14{Cf18}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D5) Dtor_cf18() _dafny.MultiSet {
	return _this.Get_().(D5_DC14).Cf18
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf18) + ")"
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
			return ok && data1.Cf18.Equals(data2.Cf18)
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
	Cf20 bool
	Cf21 bool
	Cf22 _dafny.Sequence
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf20 bool, Cf21 bool, Cf22 _dafny.Sequence) D6 {
	return D6{D6_DC16{Cf20, Cf21, Cf22}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC15 struct {
	Cf19 _dafny.MultiSet
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf19 _dafny.MultiSet) D6 {
	return D6{D6_DC15{Cf19}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC16_(false, false, _dafny.EmptySeq)
}

func (_this D6) Dtor_cf20() bool {
	return _this.Get_().(D6_DC16).Cf20
}

func (_this D6) Dtor_cf21() bool {
	return _this.Get_().(D6_DC16).Cf21
}

func (_this D6) Dtor_cf22() _dafny.Sequence {
	return _this.Get_().(D6_DC16).Cf22
}

func (_this D6) Dtor_cf19() _dafny.MultiSet {
	return _this.Get_().(D6_DC15).Cf19
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf19) + ")"
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
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21 == data2.Cf21 && data1.Cf22.Equals(data2.Cf22)
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf19.Equals(data2.Cf19)
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
	Cf23 _dafny.Set
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf23 _dafny.Set) D7 {
	return D7{D7_DC17{Cf23}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D7) Dtor_cf23() _dafny.Set {
	return _this.Get_().(D7_DC17).Cf23
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf23) + ")"
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
			return ok && data1.Cf23.Equals(data2.Cf23)
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

type D8_DC18 struct {
	Cf24 _dafny.Map
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf24 _dafny.Map) D8 {
	return D8{D8_DC18{Cf24}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D8) Dtor_cf24() _dafny.Map {
	return _this.Get_().(D8_DC18).Cf24
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC18:
		{
			return "D8.DC18" + "(" + _dafny.String(data.Cf24) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC18:
		{
			data2, ok := other.Get_().(D8_DC18)
			return ok && data1.Cf24.Equals(data2.Cf24)
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

type D9_DC20 struct {
	Cf26 bool
	Cf27 _dafny.CodePoint
	Cf28 _dafny.Int
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf26 bool, Cf27 _dafny.CodePoint, Cf28 _dafny.Int) D9 {
	return D9{D9_DC20{Cf26, Cf27, Cf28}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

type D9_DC19 struct {
	Cf25 _dafny.Sequence
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_(Cf25 _dafny.Sequence) D9 {
	return D9{D9_DC19{Cf25}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC20_(false, _dafny.CodePoint('D'), _dafny.Zero)
}

func (_this D9) Dtor_cf26() bool {
	return _this.Get_().(D9_DC20).Cf26
}

func (_this D9) Dtor_cf27() _dafny.CodePoint {
	return _this.Get_().(D9_DC20).Cf27
}

func (_this D9) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D9_DC20).Cf28
}

func (_this D9) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D9_DC19).Cf25
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC20:
		{
			return "D9.DC20" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D9_DC19:
		{
			return "D9.DC19" + "(" + data.Cf25.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC20:
		{
			data2, ok := other.Get_().(D9_DC20)
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27 && data1.Cf28.Cmp(data2.Cf28) == 0
		}
	case D9_DC19:
		{
			data2, ok := other.Get_().(D9_DC19)
			return ok && data1.Cf25.Equals(data2.Cf25)
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
	Cf30 _dafny.Sequence
	Cf31 bool
	Cf32 bool
	Cf33 _dafny.Array
	Cf34 _dafny.Int
}

func (D10_DC22) isD10() {}

func (CompanionStruct_D10_) Create_DC22_(Cf30 _dafny.Sequence, Cf31 bool, Cf32 bool, Cf33 _dafny.Array, Cf34 _dafny.Int) D10 {
	return D10{D10_DC22{Cf30, Cf31, Cf32, Cf33, Cf34}}
}

func (_this D10) Is_DC22() bool {
	_, ok := _this.Get_().(D10_DC22)
	return ok
}

type D10_DC21 struct {
	Cf29 _dafny.Map
}

func (D10_DC21) isD10() {}

func (CompanionStruct_D10_) Create_DC21_(Cf29 _dafny.Map) D10 {
	return D10{D10_DC21{Cf29}}
}

func (_this D10) Is_DC21() bool {
	_, ok := _this.Get_().(D10_DC21)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC22_(_dafny.EmptySeq, false, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D10) Dtor_cf30() _dafny.Sequence {
	return _this.Get_().(D10_DC22).Cf30
}

func (_this D10) Dtor_cf31() bool {
	return _this.Get_().(D10_DC22).Cf31
}

func (_this D10) Dtor_cf32() bool {
	return _this.Get_().(D10_DC22).Cf32
}

func (_this D10) Dtor_cf33() _dafny.Array {
	return _this.Get_().(D10_DC22).Cf33
}

func (_this D10) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D10_DC22).Cf34
}

func (_this D10) Dtor_cf29() _dafny.Map {
	return _this.Get_().(D10_DC21).Cf29
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC22:
		{
			return "D10.DC22" + "(" + data.Cf30.VerbatimString(true) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D10_DC21:
		{
			return "D10.DC21" + "(" + _dafny.String(data.Cf29) + ")"
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
			return ok && data1.Cf30.Equals(data2.Cf30) && data1.Cf31 == data2.Cf31 && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33 && data1.Cf34.Cmp(data2.Cf34) == 0
		}
	case D10_DC21:
		{
			data2, ok := other.Get_().(D10_DC21)
			return ok && data1.Cf29.Equals(data2.Cf29)
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
	Cf36 bool
	Cf37 bool
	Cf38 _dafny.Int
}

func (D11_DC24) isD11() {}

func (CompanionStruct_D11_) Create_DC24_(Cf36 bool, Cf37 bool, Cf38 _dafny.Int) D11 {
	return D11{D11_DC24{Cf36, Cf37, Cf38}}
}

func (_this D11) Is_DC24() bool {
	_, ok := _this.Get_().(D11_DC24)
	return ok
}

type D11_DC25 struct {
	Cf39 bool
	Cf40 _dafny.Int
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_(Cf39 bool, Cf40 _dafny.Int) D11 {
	return D11{D11_DC25{Cf39, Cf40}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

type D11_DC26 struct {
	Cf41 bool
}

func (D11_DC26) isD11() {}

func (CompanionStruct_D11_) Create_DC26_(Cf41 bool) D11 {
	return D11{D11_DC26{Cf41}}
}

func (_this D11) Is_DC26() bool {
	_, ok := _this.Get_().(D11_DC26)
	return ok
}

type D11_DC23 struct {
	Cf35 _dafny.Array
}

func (D11_DC23) isD11() {}

func (CompanionStruct_D11_) Create_DC23_(Cf35 _dafny.Array) D11 {
	return D11{D11_DC23{Cf35}}
}

func (_this D11) Is_DC23() bool {
	_, ok := _this.Get_().(D11_DC23)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC24_(false, false, _dafny.Zero)
}

func (_this D11) Dtor_cf36() bool {
	return _this.Get_().(D11_DC24).Cf36
}

func (_this D11) Dtor_cf37() bool {
	return _this.Get_().(D11_DC24).Cf37
}

func (_this D11) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D11_DC24).Cf38
}

func (_this D11) Dtor_cf39() bool {
	return _this.Get_().(D11_DC25).Cf39
}

func (_this D11) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D11_DC25).Cf40
}

func (_this D11) Dtor_cf41() bool {
	return _this.Get_().(D11_DC26).Cf41
}

func (_this D11) Dtor_cf35() _dafny.Array {
	return _this.Get_().(D11_DC23).Cf35
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC24:
		{
			return "D11.DC24" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D11_DC25:
		{
			return "D11.DC25" + "(" + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ")"
		}
	case D11_DC26:
		{
			return "D11.DC26" + "(" + _dafny.String(data.Cf41) + ")"
		}
	case D11_DC23:
		{
			return "D11.DC23" + "(" + _dafny.String(data.Cf35) + ")"
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
			return ok && data1.Cf36 == data2.Cf36 && data1.Cf37 == data2.Cf37 && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D11_DC25:
		{
			data2, ok := other.Get_().(D11_DC25)
			return ok && data1.Cf39 == data2.Cf39 && data1.Cf40.Cmp(data2.Cf40) == 0
		}
	case D11_DC26:
		{
			data2, ok := other.Get_().(D11_DC26)
			return ok && data1.Cf41 == data2.Cf41
		}
	case D11_DC23:
		{
			data2, ok := other.Get_().(D11_DC23)
			return ok && data1.Cf35 == data2.Cf35
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
	Cf43 _dafny.Int
	Cf44 _dafny.Int
	Cf45 bool
	Cf46 bool
	Cf47 _dafny.Int
}

func (D12_DC28) isD12() {}

func (CompanionStruct_D12_) Create_DC28_(Cf43 _dafny.Int, Cf44 _dafny.Int, Cf45 bool, Cf46 bool, Cf47 _dafny.Int) D12 {
	return D12{D12_DC28{Cf43, Cf44, Cf45, Cf46, Cf47}}
}

func (_this D12) Is_DC28() bool {
	_, ok := _this.Get_().(D12_DC28)
	return ok
}

type D12_DC27 struct {
	Cf42 _dafny.Sequence
}

func (D12_DC27) isD12() {}

func (CompanionStruct_D12_) Create_DC27_(Cf42 _dafny.Sequence) D12 {
	return D12{D12_DC27{Cf42}}
}

func (_this D12) Is_DC27() bool {
	_, ok := _this.Get_().(D12_DC27)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC28_(_dafny.Zero, _dafny.Zero, false, false, _dafny.Zero)
}

func (_this D12) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D12_DC28).Cf43
}

func (_this D12) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D12_DC28).Cf44
}

func (_this D12) Dtor_cf45() bool {
	return _this.Get_().(D12_DC28).Cf45
}

func (_this D12) Dtor_cf46() bool {
	return _this.Get_().(D12_DC28).Cf46
}

func (_this D12) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D12_DC28).Cf47
}

func (_this D12) Dtor_cf42() _dafny.Sequence {
	return _this.Get_().(D12_DC27).Cf42
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC28:
		{
			return "D12.DC28" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ")"
		}
	case D12_DC27:
		{
			return "D12.DC27" + "(" + _dafny.String(data.Cf42) + ")"
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
			return ok && data1.Cf43.Cmp(data2.Cf43) == 0 && data1.Cf44.Cmp(data2.Cf44) == 0 && data1.Cf45 == data2.Cf45 && data1.Cf46 == data2.Cf46 && data1.Cf47.Cmp(data2.Cf47) == 0
		}
	case D12_DC27:
		{
			data2, ok := other.Get_().(D12_DC27)
			return ok && data1.Cf42.Equals(data2.Cf42)
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
	Cf48 _dafny.Set
}

func (D13_DC29) isD13() {}

func (CompanionStruct_D13_) Create_DC29_(Cf48 _dafny.Set) D13 {
	return D13{D13_DC29{Cf48}}
}

func (_this D13) Is_DC29() bool {
	_, ok := _this.Get_().(D13_DC29)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D13) Dtor_cf48() _dafny.Set {
	return _this.Get_().(D13_DC29).Cf48
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC29:
		{
			return "D13.DC29" + "(" + _dafny.String(data.Cf48) + ")"
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
			return ok && data1.Cf48.Equals(data2.Cf48)
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
	Cf49 _dafny.Map
}

func (D14_DC30) isD14() {}

func (CompanionStruct_D14_) Create_DC30_(Cf49 _dafny.Map) D14 {
	return D14{D14_DC30{Cf49}}
}

func (_this D14) Is_DC30() bool {
	_, ok := _this.Get_().(D14_DC30)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D14) Dtor_cf49() _dafny.Map {
	return _this.Get_().(D14_DC30).Cf49
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC30:
		{
			return "D14.DC30" + "(" + _dafny.String(data.Cf49) + ")"
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
			return ok && data1.Cf49.Equals(data2.Cf49)
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

type D15_DC31 struct {
	Cf50 _dafny.Array
}

func (D15_DC31) isD15() {}

func (CompanionStruct_D15_) Create_DC31_(Cf50 _dafny.Array) D15 {
	return D15{D15_DC31{Cf50}}
}

func (_this D15) Is_DC31() bool {
	_, ok := _this.Get_().(D15_DC31)
	return ok
}

func (CompanionStruct_D15_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D15) Dtor_cf50() _dafny.Array {
	return _this.Get_().(D15_DC31).Cf50
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC31:
		{
			return "D15.DC31" + "(" + _dafny.String(data.Cf50) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC31:
		{
			data2, ok := other.Get_().(D15_DC31)
			return ok && data1.Cf50 == data2.Cf50
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

type D16_DC33 struct {
	Cf52 bool
	Cf53 _dafny.Int
	Cf54 _dafny.MultiSet
	Cf55 _dafny.Int
	Cf56 _dafny.Set
}

func (D16_DC33) isD16() {}

func (CompanionStruct_D16_) Create_DC33_(Cf52 bool, Cf53 _dafny.Int, Cf54 _dafny.MultiSet, Cf55 _dafny.Int, Cf56 _dafny.Set) D16 {
	return D16{D16_DC33{Cf52, Cf53, Cf54, Cf55, Cf56}}
}

func (_this D16) Is_DC33() bool {
	_, ok := _this.Get_().(D16_DC33)
	return ok
}

type D16_DC32 struct {
	Cf51 _dafny.Array
}

func (D16_DC32) isD16() {}

func (CompanionStruct_D16_) Create_DC32_(Cf51 _dafny.Array) D16 {
	return D16{D16_DC32{Cf51}}
}

func (_this D16) Is_DC32() bool {
	_, ok := _this.Get_().(D16_DC32)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC33_(false, _dafny.Zero, _dafny.EmptyMultiSet, _dafny.Zero, _dafny.EmptySet)
}

func (_this D16) Dtor_cf52() bool {
	return _this.Get_().(D16_DC33).Cf52
}

func (_this D16) Dtor_cf53() _dafny.Int {
	return _this.Get_().(D16_DC33).Cf53
}

func (_this D16) Dtor_cf54() _dafny.MultiSet {
	return _this.Get_().(D16_DC33).Cf54
}

func (_this D16) Dtor_cf55() _dafny.Int {
	return _this.Get_().(D16_DC33).Cf55
}

func (_this D16) Dtor_cf56() _dafny.Set {
	return _this.Get_().(D16_DC33).Cf56
}

func (_this D16) Dtor_cf51() _dafny.Array {
	return _this.Get_().(D16_DC32).Cf51
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC33:
		{
			return "D16.DC33" + "(" + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ")"
		}
	case D16_DC32:
		{
			return "D16.DC32" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC33:
		{
			data2, ok := other.Get_().(D16_DC33)
			return ok && data1.Cf52 == data2.Cf52 && data1.Cf53.Cmp(data2.Cf53) == 0 && data1.Cf54.Equals(data2.Cf54) && data1.Cf55.Cmp(data2.Cf55) == 0 && data1.Cf56.Equals(data2.Cf56)
		}
	case D16_DC32:
		{
			data2, ok := other.Get_().(D16_DC32)
			return ok && data1.Cf51 == data2.Cf51
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

type D17_DC34 struct {
	Cf57 _dafny.Sequence
}

func (D17_DC34) isD17() {}

func (CompanionStruct_D17_) Create_DC34_(Cf57 _dafny.Sequence) D17 {
	return D17{D17_DC34{Cf57}}
}

func (_this D17) Is_DC34() bool {
	_, ok := _this.Get_().(D17_DC34)
	return ok
}

func (CompanionStruct_D17_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D17) Dtor_cf57() _dafny.Sequence {
	return _this.Get_().(D17_DC34).Cf57
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC34:
		{
			return "D17.DC34" + "(" + _dafny.String(data.Cf57) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC34:
		{
			data2, ok := other.Get_().(D17_DC34)
			return ok && data1.Cf57.Equals(data2.Cf57)
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

type D18_DC36 struct {
	Cf59 _dafny.Int
	Cf60 _dafny.Int
	Cf61 _dafny.Int
}

func (D18_DC36) isD18() {}

func (CompanionStruct_D18_) Create_DC36_(Cf59 _dafny.Int, Cf60 _dafny.Int, Cf61 _dafny.Int) D18 {
	return D18{D18_DC36{Cf59, Cf60, Cf61}}
}

func (_this D18) Is_DC36() bool {
	_, ok := _this.Get_().(D18_DC36)
	return ok
}

type D18_DC35 struct {
	Cf58 _dafny.Sequence
}

func (D18_DC35) isD18() {}

func (CompanionStruct_D18_) Create_DC35_(Cf58 _dafny.Sequence) D18 {
	return D18{D18_DC35{Cf58}}
}

func (_this D18) Is_DC35() bool {
	_, ok := _this.Get_().(D18_DC35)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC36_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D18) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D18_DC36).Cf59
}

func (_this D18) Dtor_cf60() _dafny.Int {
	return _this.Get_().(D18_DC36).Cf60
}

func (_this D18) Dtor_cf61() _dafny.Int {
	return _this.Get_().(D18_DC36).Cf61
}

func (_this D18) Dtor_cf58() _dafny.Sequence {
	return _this.Get_().(D18_DC35).Cf58
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC36:
		{
			return "D18.DC36" + "(" + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ")"
		}
	case D18_DC35:
		{
			return "D18.DC35" + "(" + _dafny.String(data.Cf58) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC36:
		{
			data2, ok := other.Get_().(D18_DC36)
			return ok && data1.Cf59.Cmp(data2.Cf59) == 0 && data1.Cf60.Cmp(data2.Cf60) == 0 && data1.Cf61.Cmp(data2.Cf61) == 0
		}
	case D18_DC35:
		{
			data2, ok := other.Get_().(D18_DC35)
			return ok && data1.Cf58.Equals(data2.Cf58)
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

type D19_DC38 struct {
	Cf63 bool
	Cf64 bool
}

func (D19_DC38) isD19() {}

func (CompanionStruct_D19_) Create_DC38_(Cf63 bool, Cf64 bool) D19 {
	return D19{D19_DC38{Cf63, Cf64}}
}

func (_this D19) Is_DC38() bool {
	_, ok := _this.Get_().(D19_DC38)
	return ok
}

type D19_DC37 struct {
	Cf62 _dafny.Map
}

func (D19_DC37) isD19() {}

func (CompanionStruct_D19_) Create_DC37_(Cf62 _dafny.Map) D19 {
	return D19{D19_DC37{Cf62}}
}

func (_this D19) Is_DC37() bool {
	_, ok := _this.Get_().(D19_DC37)
	return ok
}

type D19_DC39 struct {
	Cf65 D19
}

func (D19_DC39) isD19() {}

func (CompanionStruct_D19_) Create_DC39_(Cf65 D19) D19 {
	return D19{D19_DC39{Cf65}}
}

func (_this D19) Is_DC39() bool {
	_, ok := _this.Get_().(D19_DC39)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC38_(false, false)
}

func (_this D19) Dtor_cf63() bool {
	return _this.Get_().(D19_DC38).Cf63
}

func (_this D19) Dtor_cf64() bool {
	return _this.Get_().(D19_DC38).Cf64
}

func (_this D19) Dtor_cf62() _dafny.Map {
	return _this.Get_().(D19_DC37).Cf62
}

func (_this D19) Dtor_cf65() D19 {
	return _this.Get_().(D19_DC39).Cf65
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC38:
		{
			return "D19.DC38" + "(" + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ")"
		}
	case D19_DC37:
		{
			return "D19.DC37" + "(" + _dafny.String(data.Cf62) + ")"
		}
	case D19_DC39:
		{
			return "D19.DC39" + "(" + _dafny.String(data.Cf65) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC38:
		{
			data2, ok := other.Get_().(D19_DC38)
			return ok && data1.Cf63 == data2.Cf63 && data1.Cf64 == data2.Cf64
		}
	case D19_DC37:
		{
			data2, ok := other.Get_().(D19_DC37)
			return ok && data1.Cf62.Equals(data2.Cf62)
		}
	case D19_DC39:
		{
			data2, ok := other.Get_().(D19_DC39)
			return ok && data1.Cf65.Equals(data2.Cf65)
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

type D20_DC41 struct {
	Cf67 _dafny.Int
}

func (D20_DC41) isD20() {}

func (CompanionStruct_D20_) Create_DC41_(Cf67 _dafny.Int) D20 {
	return D20{D20_DC41{Cf67}}
}

func (_this D20) Is_DC41() bool {
	_, ok := _this.Get_().(D20_DC41)
	return ok
}

type D20_DC40 struct {
	Cf66 _dafny.Map
}

func (D20_DC40) isD20() {}

func (CompanionStruct_D20_) Create_DC40_(Cf66 _dafny.Map) D20 {
	return D20{D20_DC40{Cf66}}
}

func (_this D20) Is_DC40() bool {
	_, ok := _this.Get_().(D20_DC40)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC41_(_dafny.Zero)
}

func (_this D20) Dtor_cf67() _dafny.Int {
	return _this.Get_().(D20_DC41).Cf67
}

func (_this D20) Dtor_cf66() _dafny.Map {
	return _this.Get_().(D20_DC40).Cf66
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC41:
		{
			return "D20.DC41" + "(" + _dafny.String(data.Cf67) + ")"
		}
	case D20_DC40:
		{
			return "D20.DC40" + "(" + _dafny.String(data.Cf66) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC41:
		{
			data2, ok := other.Get_().(D20_DC41)
			return ok && data1.Cf67.Cmp(data2.Cf67) == 0
		}
	case D20_DC40:
		{
			data2, ok := other.Get_().(D20_DC40)
			return ok && data1.Cf66.Equals(data2.Cf66)
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

type D21_DC43 struct {
	Cf69 bool
	Cf70 bool
	Cf71 _dafny.Sequence
}

func (D21_DC43) isD21() {}

func (CompanionStruct_D21_) Create_DC43_(Cf69 bool, Cf70 bool, Cf71 _dafny.Sequence) D21 {
	return D21{D21_DC43{Cf69, Cf70, Cf71}}
}

func (_this D21) Is_DC43() bool {
	_, ok := _this.Get_().(D21_DC43)
	return ok
}

type D21_DC44 struct {
	Cf72 bool
}

func (D21_DC44) isD21() {}

func (CompanionStruct_D21_) Create_DC44_(Cf72 bool) D21 {
	return D21{D21_DC44{Cf72}}
}

func (_this D21) Is_DC44() bool {
	_, ok := _this.Get_().(D21_DC44)
	return ok
}

type D21_DC42 struct {
	Cf68 _dafny.Array
}

func (D21_DC42) isD21() {}

func (CompanionStruct_D21_) Create_DC42_(Cf68 _dafny.Array) D21 {
	return D21{D21_DC42{Cf68}}
}

func (_this D21) Is_DC42() bool {
	_, ok := _this.Get_().(D21_DC42)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC43_(false, false, _dafny.EmptySeq)
}

func (_this D21) Dtor_cf69() bool {
	return _this.Get_().(D21_DC43).Cf69
}

func (_this D21) Dtor_cf70() bool {
	return _this.Get_().(D21_DC43).Cf70
}

func (_this D21) Dtor_cf71() _dafny.Sequence {
	return _this.Get_().(D21_DC43).Cf71
}

func (_this D21) Dtor_cf72() bool {
	return _this.Get_().(D21_DC44).Cf72
}

func (_this D21) Dtor_cf68() _dafny.Array {
	return _this.Get_().(D21_DC42).Cf68
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC43:
		{
			return "D21.DC43" + "(" + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ")"
		}
	case D21_DC44:
		{
			return "D21.DC44" + "(" + _dafny.String(data.Cf72) + ")"
		}
	case D21_DC42:
		{
			return "D21.DC42" + "(" + _dafny.String(data.Cf68) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC43:
		{
			data2, ok := other.Get_().(D21_DC43)
			return ok && data1.Cf69 == data2.Cf69 && data1.Cf70 == data2.Cf70 && data1.Cf71.Equals(data2.Cf71)
		}
	case D21_DC44:
		{
			data2, ok := other.Get_().(D21_DC44)
			return ok && data1.Cf72 == data2.Cf72
		}
	case D21_DC42:
		{
			data2, ok := other.Get_().(D21_DC42)
			return ok && data1.Cf68 == data2.Cf68
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

type D22_DC46 struct {
	Cf74 _dafny.Int
	Cf75 _dafny.Int
	Cf76 bool
}

func (D22_DC46) isD22() {}

func (CompanionStruct_D22_) Create_DC46_(Cf74 _dafny.Int, Cf75 _dafny.Int, Cf76 bool) D22 {
	return D22{D22_DC46{Cf74, Cf75, Cf76}}
}

func (_this D22) Is_DC46() bool {
	_, ok := _this.Get_().(D22_DC46)
	return ok
}

type D22_DC47 struct {
	Cf77 _dafny.Int
	Cf78 bool
	Cf79 bool
	Cf80 _dafny.Int
	Cf81 _dafny.Int
}

func (D22_DC47) isD22() {}

func (CompanionStruct_D22_) Create_DC47_(Cf77 _dafny.Int, Cf78 bool, Cf79 bool, Cf80 _dafny.Int, Cf81 _dafny.Int) D22 {
	return D22{D22_DC47{Cf77, Cf78, Cf79, Cf80, Cf81}}
}

func (_this D22) Is_DC47() bool {
	_, ok := _this.Get_().(D22_DC47)
	return ok
}

type D22_DC45 struct {
	Cf73 _dafny.Array
}

func (D22_DC45) isD22() {}

func (CompanionStruct_D22_) Create_DC45_(Cf73 _dafny.Array) D22 {
	return D22{D22_DC45{Cf73}}
}

func (_this D22) Is_DC45() bool {
	_, ok := _this.Get_().(D22_DC45)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC46_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D22) Dtor_cf74() _dafny.Int {
	return _this.Get_().(D22_DC46).Cf74
}

func (_this D22) Dtor_cf75() _dafny.Int {
	return _this.Get_().(D22_DC46).Cf75
}

func (_this D22) Dtor_cf76() bool {
	return _this.Get_().(D22_DC46).Cf76
}

func (_this D22) Dtor_cf77() _dafny.Int {
	return _this.Get_().(D22_DC47).Cf77
}

func (_this D22) Dtor_cf78() bool {
	return _this.Get_().(D22_DC47).Cf78
}

func (_this D22) Dtor_cf79() bool {
	return _this.Get_().(D22_DC47).Cf79
}

func (_this D22) Dtor_cf80() _dafny.Int {
	return _this.Get_().(D22_DC47).Cf80
}

func (_this D22) Dtor_cf81() _dafny.Int {
	return _this.Get_().(D22_DC47).Cf81
}

func (_this D22) Dtor_cf73() _dafny.Array {
	return _this.Get_().(D22_DC45).Cf73
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC46:
		{
			return "D22.DC46" + "(" + _dafny.String(data.Cf74) + ", " + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ")"
		}
	case D22_DC47:
		{
			return "D22.DC47" + "(" + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ", " + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ")"
		}
	case D22_DC45:
		{
			return "D22.DC45" + "(" + _dafny.String(data.Cf73) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC46:
		{
			data2, ok := other.Get_().(D22_DC46)
			return ok && data1.Cf74.Cmp(data2.Cf74) == 0 && data1.Cf75.Cmp(data2.Cf75) == 0 && data1.Cf76 == data2.Cf76
		}
	case D22_DC47:
		{
			data2, ok := other.Get_().(D22_DC47)
			return ok && data1.Cf77.Cmp(data2.Cf77) == 0 && data1.Cf78 == data2.Cf78 && data1.Cf79 == data2.Cf79 && data1.Cf80.Cmp(data2.Cf80) == 0 && data1.Cf81.Cmp(data2.Cf81) == 0
		}
	case D22_DC45:
		{
			data2, ok := other.Get_().(D22_DC45)
			return ok && data1.Cf73 == data2.Cf73
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

type D23_DC49 struct {
	Cf83 _dafny.Int
	Cf84 _dafny.Int
	Cf85 _dafny.Int
}

func (D23_DC49) isD23() {}

func (CompanionStruct_D23_) Create_DC49_(Cf83 _dafny.Int, Cf84 _dafny.Int, Cf85 _dafny.Int) D23 {
	return D23{D23_DC49{Cf83, Cf84, Cf85}}
}

func (_this D23) Is_DC49() bool {
	_, ok := _this.Get_().(D23_DC49)
	return ok
}

type D23_DC48 struct {
	Cf82 *C0
}

func (D23_DC48) isD23() {}

func (CompanionStruct_D23_) Create_DC48_(Cf82 *C0) D23 {
	return D23{D23_DC48{Cf82}}
}

func (_this D23) Is_DC48() bool {
	_, ok := _this.Get_().(D23_DC48)
	return ok
}

type D23_DC50 struct {
	Cf86 D23
}

func (D23_DC50) isD23() {}

func (CompanionStruct_D23_) Create_DC50_(Cf86 D23) D23 {
	return D23{D23_DC50{Cf86}}
}

func (_this D23) Is_DC50() bool {
	_, ok := _this.Get_().(D23_DC50)
	return ok
}

func (CompanionStruct_D23_) Default() D23 {
	return Companion_D23_.Create_DC49_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D23) Dtor_cf83() _dafny.Int {
	return _this.Get_().(D23_DC49).Cf83
}

func (_this D23) Dtor_cf84() _dafny.Int {
	return _this.Get_().(D23_DC49).Cf84
}

func (_this D23) Dtor_cf85() _dafny.Int {
	return _this.Get_().(D23_DC49).Cf85
}

func (_this D23) Dtor_cf82() *C0 {
	return _this.Get_().(D23_DC48).Cf82
}

func (_this D23) Dtor_cf86() D23 {
	return _this.Get_().(D23_DC50).Cf86
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC49:
		{
			return "D23.DC49" + "(" + _dafny.String(data.Cf83) + ", " + _dafny.String(data.Cf84) + ", " + _dafny.String(data.Cf85) + ")"
		}
	case D23_DC48:
		{
			return "D23.DC48" + "(" + _dafny.String(data.Cf82) + ")"
		}
	case D23_DC50:
		{
			return "D23.DC50" + "(" + _dafny.String(data.Cf86) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC49:
		{
			data2, ok := other.Get_().(D23_DC49)
			return ok && data1.Cf83.Cmp(data2.Cf83) == 0 && data1.Cf84.Cmp(data2.Cf84) == 0 && data1.Cf85.Cmp(data2.Cf85) == 0
		}
	case D23_DC48:
		{
			data2, ok := other.Get_().(D23_DC48)
			return ok && data1.Cf82 == data2.Cf82
		}
	case D23_DC50:
		{
			data2, ok := other.Get_().(D23_DC50)
			return ok && data1.Cf86.Equals(data2.Cf86)
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

type D24_DC51 struct {
	Cf87 T0
}

func (D24_DC51) isD24() {}

func (CompanionStruct_D24_) Create_DC51_(Cf87 T0) D24 {
	return D24{D24_DC51{Cf87}}
}

func (_this D24) Is_DC51() bool {
	_, ok := _this.Get_().(D24_DC51)
	return ok
}

func (CompanionStruct_D24_) Default() T0 {
	return (T0)(nil)
}

func (_this D24) Dtor_cf87() T0 {
	return _this.Get_().(D24_DC51).Cf87
}

func (_this D24) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D24_DC51:
		{
			return "D24.DC51" + "(" + _dafny.String(data.Cf87) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D24) Equals(other D24) bool {
	switch data1 := _this.Get_().(type) {
	case D24_DC51:
		{
			data2, ok := other.Get_().(D24_DC51)
			return ok && _dafny.AreEqual(data1.Cf87, data2.Cf87)
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

type D25_DC52 struct {
	Cf88 _dafny.MultiSet
}

func (D25_DC52) isD25() {}

func (CompanionStruct_D25_) Create_DC52_(Cf88 _dafny.MultiSet) D25 {
	return D25{D25_DC52{Cf88}}
}

func (_this D25) Is_DC52() bool {
	_, ok := _this.Get_().(D25_DC52)
	return ok
}

func (CompanionStruct_D25_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D25) Dtor_cf88() _dafny.MultiSet {
	return _this.Get_().(D25_DC52).Cf88
}

func (_this D25) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D25_DC52:
		{
			return "D25.DC52" + "(" + _dafny.String(data.Cf88) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D25) Equals(other D25) bool {
	switch data1 := _this.Get_().(type) {
	case D25_DC52:
		{
			data2, ok := other.Get_().(D25_DC52)
			return ok && data1.Cf88.Equals(data2.Cf88)
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

type D26_DC54 struct {
	Cf90 _dafny.Int
	Cf91 _dafny.Int
	Cf92 bool
	Cf93 _dafny.Int
}

func (D26_DC54) isD26() {}

func (CompanionStruct_D26_) Create_DC54_(Cf90 _dafny.Int, Cf91 _dafny.Int, Cf92 bool, Cf93 _dafny.Int) D26 {
	return D26{D26_DC54{Cf90, Cf91, Cf92, Cf93}}
}

func (_this D26) Is_DC54() bool {
	_, ok := _this.Get_().(D26_DC54)
	return ok
}

type D26_DC53 struct {
	Cf89 _dafny.Sequence
}

func (D26_DC53) isD26() {}

func (CompanionStruct_D26_) Create_DC53_(Cf89 _dafny.Sequence) D26 {
	return D26{D26_DC53{Cf89}}
}

func (_this D26) Is_DC53() bool {
	_, ok := _this.Get_().(D26_DC53)
	return ok
}

type D26_DC55 struct {
	Cf94 D26
}

func (D26_DC55) isD26() {}

func (CompanionStruct_D26_) Create_DC55_(Cf94 D26) D26 {
	return D26{D26_DC55{Cf94}}
}

func (_this D26) Is_DC55() bool {
	_, ok := _this.Get_().(D26_DC55)
	return ok
}

func (CompanionStruct_D26_) Default() D26 {
	return Companion_D26_.Create_DC54_(_dafny.Zero, _dafny.Zero, false, _dafny.Zero)
}

func (_this D26) Dtor_cf90() _dafny.Int {
	return _this.Get_().(D26_DC54).Cf90
}

func (_this D26) Dtor_cf91() _dafny.Int {
	return _this.Get_().(D26_DC54).Cf91
}

func (_this D26) Dtor_cf92() bool {
	return _this.Get_().(D26_DC54).Cf92
}

func (_this D26) Dtor_cf93() _dafny.Int {
	return _this.Get_().(D26_DC54).Cf93
}

func (_this D26) Dtor_cf89() _dafny.Sequence {
	return _this.Get_().(D26_DC53).Cf89
}

func (_this D26) Dtor_cf94() D26 {
	return _this.Get_().(D26_DC55).Cf94
}

func (_this D26) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D26_DC54:
		{
			return "D26.DC54" + "(" + _dafny.String(data.Cf90) + ", " + _dafny.String(data.Cf91) + ", " + _dafny.String(data.Cf92) + ", " + _dafny.String(data.Cf93) + ")"
		}
	case D26_DC53:
		{
			return "D26.DC53" + "(" + _dafny.String(data.Cf89) + ")"
		}
	case D26_DC55:
		{
			return "D26.DC55" + "(" + _dafny.String(data.Cf94) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D26) Equals(other D26) bool {
	switch data1 := _this.Get_().(type) {
	case D26_DC54:
		{
			data2, ok := other.Get_().(D26_DC54)
			return ok && data1.Cf90.Cmp(data2.Cf90) == 0 && data1.Cf91.Cmp(data2.Cf91) == 0 && data1.Cf92 == data2.Cf92 && data1.Cf93.Cmp(data2.Cf93) == 0
		}
	case D26_DC53:
		{
			data2, ok := other.Get_().(D26_DC53)
			return ok && data1.Cf89.Equals(data2.Cf89)
		}
	case D26_DC55:
		{
			data2, ok := other.Get_().(D26_DC55)
			return ok && data1.Cf94.Equals(data2.Cf94)
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

type D27_DC57 struct {
	Cf96 _dafny.Sequence
}

func (D27_DC57) isD27() {}

func (CompanionStruct_D27_) Create_DC57_(Cf96 _dafny.Sequence) D27 {
	return D27{D27_DC57{Cf96}}
}

func (_this D27) Is_DC57() bool {
	_, ok := _this.Get_().(D27_DC57)
	return ok
}

type D27_DC56 struct {
	Cf95 *C2
}

func (D27_DC56) isD27() {}

func (CompanionStruct_D27_) Create_DC56_(Cf95 *C2) D27 {
	return D27{D27_DC56{Cf95}}
}

func (_this D27) Is_DC56() bool {
	_, ok := _this.Get_().(D27_DC56)
	return ok
}

func (CompanionStruct_D27_) Default() D27 {
	return Companion_D27_.Create_DC57_(_dafny.EmptySeq)
}

func (_this D27) Dtor_cf96() _dafny.Sequence {
	return _this.Get_().(D27_DC57).Cf96
}

func (_this D27) Dtor_cf95() *C2 {
	return _this.Get_().(D27_DC56).Cf95
}

func (_this D27) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D27_DC57:
		{
			return "D27.DC57" + "(" + _dafny.String(data.Cf96) + ")"
		}
	case D27_DC56:
		{
			return "D27.DC56" + "(" + _dafny.String(data.Cf95) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D27) Equals(other D27) bool {
	switch data1 := _this.Get_().(type) {
	case D27_DC57:
		{
			data2, ok := other.Get_().(D27_DC57)
			return ok && data1.Cf96.Equals(data2.Cf96)
		}
	case D27_DC56:
		{
			data2, ok := other.Get_().(D27_DC56)
			return ok && data1.Cf95 == data2.Cf95
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

type D28_DC58 struct {
	Cf97 _dafny.Map
}

func (D28_DC58) isD28() {}

func (CompanionStruct_D28_) Create_DC58_(Cf97 _dafny.Map) D28 {
	return D28{D28_DC58{Cf97}}
}

func (_this D28) Is_DC58() bool {
	_, ok := _this.Get_().(D28_DC58)
	return ok
}

func (CompanionStruct_D28_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D28) Dtor_cf97() _dafny.Map {
	return _this.Get_().(D28_DC58).Cf97
}

func (_this D28) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D28_DC58:
		{
			return "D28.DC58" + "(" + _dafny.String(data.Cf97) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D28) Equals(other D28) bool {
	switch data1 := _this.Get_().(type) {
	case D28_DC58:
		{
			data2, ok := other.Get_().(D28_DC58)
			return ok && data1.Cf97.Equals(data2.Cf97)
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
	Fm6(p0 D0, p1 D0, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet
	Fm7(globalState *GlobalState) _dafny.Int
	M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Map
	M2(p0 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Array)
	F4() bool
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
	Fm6(p0 D0, p1 D0, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet
	Fm7(globalState *GlobalState) _dafny.Int
	M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Map
	M2(p0 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Array)
	F4() bool
	F5() _dafny.Array
	F5_set_(value _dafny.Array)
	Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) bool
	Fm9(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int
	F6() _dafny.Int
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
	_f0 _dafny.MultiSet
	_f1 _dafny.Int
	_f2 bool
	_f3 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this._f0 = _dafny.EmptyMultiSet
	_this._f1 = _dafny.Zero
	_this._f2 = false
	_this._f3 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.MultiSet, f1 _dafny.Int, f2 bool, f3 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
	}
}
func (_this *GlobalState) F0() _dafny.MultiSet {
	{
		return _this._f0
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

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f12 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f12 = _dafny.Zero
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

func (_this *C0) Ctor__(f12 _dafny.Int) {
	{
		(_this)._f12 = f12
	}
}
func (_this *C0) Fm17(p0 bool, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C0) F12() _dafny.Int {
	{
		return _this._f12
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f5  _dafny.Array
	_f6  _dafny.Int
	_f4  bool
	F16  _dafny.Int
	_f15 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.Zero
	_this._f4 = false
	_this.F16 = _dafny.Zero
	_this._f15 = _dafny.Zero
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

func (_this *C1) F5() _dafny.Array {
	return _this._f5
}
func (_this *C1) F5_set_(value _dafny.Array) {
	_this._f5 = value
}
func (_this *C1) F6() _dafny.Int {
	return _this._f6
}
func (_this *C1) F4() bool {
	return _this._f4
}
func (_this *C1) Ctor__(f15 _dafny.Int, f16 _dafny.Int, f5 _dafny.Array, f6 _dafny.Int, f4 bool) {
	{
		(_this)._f15 = f15
		(_this).F16 = f16
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f4 = f4
	}
}
func (_this *C1) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F4()
	}
}
func (_this *C1) Fm9(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt((_this).F15(), (_dafny.Zero).Minus((_this).F6()))
	}
}
func (_this *C1) Fm6(p0 D0, p1 D0, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		if ((_this).F4()) || (true) {
			return (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(43))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}(func(_363_i0 _dafny.Int) _dafny.Int {
				return (_this).F6()
			})))).Intersection(_dafny.MultiSetOf((_this).F6(), _dafny.IntOfInt64(469)))
		} else {
			return (_dafny.MultiSetOf(_this.F16)).Union(_dafny.MultiSetOf((_this).F6()))
		}
	}
}
func (_this *C1) Fm7(globalState *GlobalState) _dafny.Int {
	{
		var _source9 D1 = Companion_D1_.Create_DC3_(!((_this).F4()))
		_ = _source9
		if _source9.Is_DC4() {
			var _364___mcc_h0 _dafny.Int = _source9.Get_().(D1_DC4).Cf3
			_ = _364___mcc_h0
			var _365___mcc_h1 _dafny.Int = _source9.Get_().(D1_DC4).Cf4
			_ = _365___mcc_h1
			var _366_cf4 _dafny.Int = _365___mcc_h1
			_ = _366_cf4
			var _367_cf3 _dafny.Int = _364___mcc_h0
			_ = _367_cf3
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iyylxfgh")).Cardinality())
		} else {
			var _368___mcc_h2 bool = _source9.Get_().(D1_DC3).Cf2
			_ = _368___mcc_h2
			var _369_cf2 bool = _368___mcc_h2
			_ = _369_cf2
			return _dafny.IntOfUint32(((func() _dafny.Sequence {
				if true {
					return _dafny.SeqOf(_369_cf2, !(_369_cf2))
				}
				return _dafny.SeqOf(_369_cf2, _369_cf2)
			})()).Cardinality())
		}
	}
}
func (_this *C1) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _370_v0 _dafny.Set
		_ = _370_v0
		_370_v0 = _dafny.SetOf((_this).F15())
		var _371_v1 _dafny.MultiSet
		_ = _371_v1
		_371_v1 = _dafny.MultiSetOf(_this.F16, (_370_v0).Cardinality(), _this.F16)
		var _372_v2 D3
		_ = _372_v2
		_372_v2 = Companion_D3_.Create_DC9_((_this).F4(), (_this).F6(), !((_this).F4()))
		var _source10 D3 = (func() D3 {
			if (_371_v1).IsProperSubsetOf(_dafny.MultiSetOf((_this).F6())) {
				return _372_v2
			}
			return func(_pat_let5_0 D3) D3 {
				return func(_373_dt__update__tmp_h0 D3) D3 {
					return func(_pat_let6_0 bool) D3 {
						return func(_374_dt__update_hcf8_h0 bool) D3 {
							return Companion_D3_.Create_DC9_(_374_dt__update_hcf8_h0, (_373_dt__update__tmp_h0).Dtor_cf9(), (_373_dt__update__tmp_h0).Dtor_cf10())
						}(_pat_let6_0)
					}(false)
				}(_pat_let5_0)
			}(_372_v2)
		})()
		_ = _source10
		if _source10.Is_DC9() {
			var _375___mcc_h0 bool = _source10.Get_().(D3_DC9).Cf8
			_ = _375___mcc_h0
			var _376___mcc_h1 _dafny.Int = _source10.Get_().(D3_DC9).Cf9
			_ = _376___mcc_h1
			var _377___mcc_h2 bool = _source10.Get_().(D3_DC9).Cf10
			_ = _377___mcc_h2
			var _378_cf10 bool = _377___mcc_h2
			_ = _378_cf10
			var _379_cf9 _dafny.Int = _376___mcc_h1
			_ = _379_cf9
			var _380_cf8 bool = _375___mcc_h0
			_ = _380_cf8
			var _381_v3 _dafny.CodePoint
			_ = _381_v3
			_381_v3 = _dafny.CodePoint('k')
			if (_dafny.IntOfInt64(-828)).Cmp((_dafny.IntOfUint32((Companion_Default___.Fm37(_380_cf8, !(_378_cf10), p0, Companion_Default___.Fm0(_381_v3, globalState), globalState)).Cardinality())).Minus(_379_cf9)) <= 0 {
				var _382_v4 _dafny.Array
				_ = _382_v4
				var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
				_ = _nw56
				_382_v4 = _nw56
				var _383_v5 _dafny.Array
				_ = _383_v5
				var _len0_4 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_4
				var _nw57 _dafny.Array
				_ = _nw57
				if _len0_4.Cmp(_dafny.Zero) == 0 {
					_nw57 = _dafny.NewArray(_len0_4)
				} else {
					var _init4 func(_dafny.Int) _dafny.Int = (func(_384_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_385_i0 _dafny.Int) _dafny.Int {
							return (_385_i0).Times(_384_p0)
						}
					})(p0)
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
				_383_v5 = _nw57
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(500), _dafny.ArrayLen((_382_v4), 0))
				_ = _index47
				(_382_v4).ArraySet1(_383_v5, (_index47).Int())
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(500), _dafny.ArrayLen((_382_v4), 0))
				_ = _index48
				(_382_v4).ArraySet1(_383_v5, (_index48).Int())
				var _386_v6 _dafny.Sequence
				_ = _386_v6
				_386_v6 = _dafny.UnicodeSeqOfUtf8Bytes("tqnumrpc")
				_386_v6 = _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(435))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_387_v6 _dafny.Sequence, _388_cf9 _dafny.Int, _389_p0 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
					return func(_390_i1 _dafny.Int) _dafny.CodePoint {
						return (_387_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((Companion_D6_.Create_DC16_((_this).F4(), (_this).F4(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(566))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg43 _dafny.Int) interface{} {
								return coer43(arg43)
							}
						}((func(_391_cf9 _dafny.Int, _392_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_393_i2 _dafny.Int) _dafny.Int {
								return (_dafny.SetOf(_dafny.IntOfInt64(732), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_391_cf9, _392_p0)).Cardinality())).Cardinality()
							}
						})(_388_cf9, _389_p0))))).Dtor_cf22()).Cardinality()), _dafny.IntOfUint32((_387_v6).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_386_v6, _379_cf9, p0))), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(435))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}((func(_394_v6 _dafny.Sequence, _395_cf9 _dafny.Int, _396_p0 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
					return func(_397_i1 _dafny.Int) _dafny.CodePoint {
						return (_394_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((Companion_D6_.Create_DC16_((_this).F4(), (_this).F4(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(566))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg45 _dafny.Int) interface{} {
								return coer45(arg45)
							}
						}((func(_398_cf9 _dafny.Int, _399_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_400_i2 _dafny.Int) _dafny.Int {
								return (_dafny.SetOf(_dafny.IntOfInt64(732), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_398_cf9, _399_p0)).Cardinality())).Cardinality()
							}
						})(_395_cf9, _396_p0))))).Dtor_cf22()).Cardinality()), _dafny.IntOfUint32((_394_v6).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_386_v6, _379_cf9, p0)))).Cardinality()))).Uint32(), _381_v3)
				var _401_v7 _dafny.Sequence
				_ = _401_v7
				_401_v7 = _dafny.SeqOf(_379_cf9)
				var _402_v8 _dafny.Map
				_ = _402_v8
				_402_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm8((_this).Fm7(globalState), (_this).F4(), _386_v6, _401_v7, globalState), _dafny.IntOfUint32((_401_v7).Cardinality()))
				r0 = _402_v8
				_380_cf8 = _378_cf10
				var _403_v9 _dafny.Sequence
				_ = _403_v9
				_403_v9 = _dafny.SeqOf((_this).F4())
				_403_v9 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_380_cf8), _403_v9)
			} else {
				var _404_v11 _dafny.Sequence
				_ = _404_v11
				_404_v11 = _dafny.SeqOf((_this).F4(), _380_cf8)
				var _405_v12 _dafny.Map
				_ = _405_v12
				_405_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll28 = _dafny.NewMapBuilder()
					_ = _coll28
					for _iter29 := _dafny.Iterate((_dafny.MultiSetOf((_this).F15())).Elements()); ; {
						_compr_28, _ok29 := _iter29()
						if !_ok29 {
							break
						}
						var _406_v10 _dafny.Int
						_406_v10 = interface{}(_compr_28).(_dafny.Int)
						if (_dafny.MultiSetOf((_this).F15())).Contains(_406_v10) {
							_coll28.Add(Companion_Default___.SafeModuloInt(_406_v10, p0), (_this).F4())
						}
					}
					return _coll28.ToMap()
				}(), Companion_Default___.Fm3(_379_cf9, (_404_v11).Select((Companion_Default___.SafeIndex(_379_cf9, _dafny.IntOfUint32((_404_v11).Cardinality()))).Uint32()).(bool), globalState))
				_405_v12 = (_405_v12).Merge(_405_v12)
				var _407_v13 *C0
				_ = _407_v13
				var _nw58 *C0 = New_C0_()
				_ = _nw58
				_nw58.Ctor__(_379_cf9)
				_407_v13 = _nw58
				(_this).F16 = Companion_Default___.SafeDivisionInt((_this).F15(), (_407_v13).F12())
				var _408_v14 _dafny.Sequence
				_ = _408_v14
				_408_v14 = _dafny.SeqOf((_dafny.Zero).Minus(p0), _dafny.IntOfInt64(886))
				(_this).F16 = ((_408_v14).Select((Companion_Default___.SafeIndex(_379_cf9, _dafny.IntOfUint32((_408_v14).Cardinality()))).Uint32()).(_dafny.Int)).Minus(_dafny.IntOfInt64(760))
				var _409_v15 _dafny.MultiSet
				_ = _409_v15
				_409_v15 = _dafny.MultiSetOf(!(!(_380_cf8)))
				var _410_v16 _dafny.Sequence
				_ = _410_v16
				_410_v16 = _dafny.UnicodeSeqOfUtf8Bytes("tplcf")
				(_this).F16 = (_379_cf9).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm37(_378_cf10, (_this).F4(), (func() _dafny.Int {
					if (_409_v15).Contains((_this).F4()) {
						return (_409_v15).Multiplicity((_this).F4())
					}
					return _dafny.IntOfInt64(150)
				})(), (_this).F4(), globalState), _dafny.Companion_Sequence_.Update(_410_v16, (Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_410_v16).Cardinality()))).Uint32(), _381_v3))).Cardinality()))
			}
			_380_cf8 = _380_cf8
			(_this).F5_set_(_this.F5())
			if _378_cf10 {
				_380_cf8 = !(_380_cf8)
				(_this).F16 = Companion_Default___.SafeModuloInt(p0, (_this).F6())
				_379_cf9 = (_379_cf9).Minus(_dafny.IntOfInt64(905))
				var _411_v17 *C0
				_ = _411_v17
				var _nw59 *C0 = New_C0_()
				_ = _nw59
				_nw59.Ctor__((_379_cf9).Minus(_379_cf9))
				_411_v17 = _nw59
				var _nw60 *C0 = New_C0_()
				_ = _nw60
				_nw60.Ctor__(_dafny.IntOfInt64(921))
				_411_v17 = _nw60
				var _412_v18 _dafny.Sequence
				_ = _412_v18
				_412_v18 = _dafny.SeqOf((_this).F4(), true, !(_378_cf10))
				var _413_v19 _dafny.Sequence
				_ = _413_v19
				_413_v19 = _dafny.SeqOf(((_this).F6()).Times(_dafny.IntOfUint32(((Companion_D12_.Create_DC27_(_dafny.SeqOf(_412_v18))).Dtor_cf42()).Cardinality())))
				_413_v19 = _413_v19
			} else {
				(_this).F16 = (_dafny.Zero).Minus(_379_cf9)
				var _414_v20 _dafny.Array
				_ = _414_v20
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_5
				var _nw61 _dafny.Array
				_ = _nw61
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw61 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) _dafny.Sequence = (func(_415_cf9 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_416_i3 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_415_cf9, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lyl")).Cardinality()))
						}
					})(_379_cf9)
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw61 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw61).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw61).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_414_v20 = _nw61
				var _417_v21 D6
				_ = _417_v21
				_417_v21 = Companion_D6_.Create_DC16_(_378_cf10, _378_cf10, _dafny.SeqOf(p0))
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_414_v20), 0))
				_ = _index49
				(_414_v20).ArraySet1((_417_v21).Dtor_cf22(), (_index49).Int())
				var _418_v22 _dafny.Sequence
				_ = _418_v22
				_418_v22 = _dafny.SeqOf((_this).F6())
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_414_v20), 0))
				_ = _index50
				(_414_v20).ArraySet1(_418_v22, (_index50).Int())
				var _419_v23 _dafny.Map
				_ = _419_v23
				_419_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _this.F16)
				var _420_v24 _dafny.Sequence
				_ = _420_v24
				_420_v24 = _dafny.UnicodeSeqOfUtf8Bytes("tlcjkdars")
				_378_cf10 = (func() bool {
					if (_this).Fm8((func() _dafny.Int {
						if (_419_v23).Contains(_380_cf8) {
							return (_419_v23).Get(_380_cf8).(_dafny.Int)
						}
						return _dafny.IntOfUint32((_420_v24).Cardinality())
					})(), true, _420_v24, _418_v22, globalState) {
						return _378_cf10
					}
					return !((_this).F4())
				})()
				var _421_v25 _dafny.Array
				_ = _421_v25
				var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw62
				_421_v25 = _nw62
				_421_v25 = _421_v25
				_381_v3 = (func() _dafny.CodePoint {
					if (_this).F4() {
						return _381_v3
					}
					return _381_v3
				})()
			}
		} else if _source10.Is_DC10() {
			var _422___mcc_h3 _dafny.Int = _source10.Get_().(D3_DC10).Cf11
			_ = _422___mcc_h3
			var _423_cf11 _dafny.Int = _422___mcc_h3
			_ = _423_cf11
			if (_this).F4() {
				var _424_v26 *C0
				_ = _424_v26
				var _nw63 *C0 = New_C0_()
				_ = _nw63
				_nw63.Ctor__((_this).Fm7(globalState))
				_424_v26 = _nw63
				var _425_v27 _dafny.Sequence
				_ = _425_v27
				_425_v27 = _dafny.UnicodeSeqOfUtf8Bytes("nfgsjgori")
				var _426_v28 D9
				_ = _426_v28
				_426_v28 = Companion_D9_.Create_DC19_(_425_v27)
				var _pat_let_tv5 = _425_v27
				_ = _pat_let_tv5
				var _427_v29 _dafny.MultiSet
				_ = _427_v29
				_427_v29 = _dafny.MultiSetOf(func(_pat_let7_0 D9) D9 {
					return func(_428_dt__update__tmp_h1 D9) D9 {
						return func(_pat_let8_0 _dafny.Sequence) D9 {
							return func(_429_dt__update_hcf25_h0 _dafny.Sequence) D9 {
								return Companion_D9_.Create_DC19_(_429_dt__update_hcf25_h0)
							}(_pat_let8_0)
						}(_pat_let_tv5)
					}(_pat_let7_0)
				}(_426_v28), _426_v28)
				var _430_v30 _dafny.Set
				_ = _430_v30
				_430_v30 = _dafny.SetOf(true, (_this).F4())
				var _431_v31 _dafny.Sequence
				_ = _431_v31
				_431_v31 = _dafny.SeqOf((_430_v30).Cardinality())
				var _432_v32 _dafny.Map
				_ = _432_v32
				_432_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_427_v29, _431_v31)
				var _433_v34 _dafny.Set
				_ = _433_v34
				_433_v34 = _dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_426_v28)), _dafny.MultiSetOf(_426_v28), _427_v29, _427_v29)
				var _434_v35 _dafny.Sequence
				_ = _434_v35
				_434_v35 = _dafny.SeqOf((_this).F4())
				var _435_v36 _dafny.Array
				_ = _435_v36
				var _nw64 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw64
				_435_v36 = _nw64
				var _436_v37 _dafny.Map
				_ = _436_v37
				_436_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_435_v36, _423_cf11)
				var _437_v38 _dafny.Array
				_ = _437_v38
				var _nwElement0_15 _dafny.Int = _this.F16
				_ = _nwElement0_15
				var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(14))
				_ = _nw65
				(_nw65).ArraySet1(_nwElement0_15, 0)
				(_nw65).ArraySet1(((_432_v32).Merge(func() _dafny.Map {
					var _coll29 = _dafny.NewMapBuilder()
					_ = _coll29
					for _iter30 := _dafny.Iterate((_433_v34).Elements()); ; {
						_compr_29, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _438_v33 _dafny.MultiSet
						_438_v33 = interface{}(_compr_29).(_dafny.MultiSet)
						if (_433_v34).Contains(_438_v33) {
							_coll29.Add(_438_v33, _dafny.SeqOf((_424_v26).F12(), _dafny.IntOfInt64(-615)))
						}
					}
					return _coll29.ToMap()
				}())).Cardinality(), 1)
				(_nw65).ArraySet1(p0, 2)
				(_nw65).ArraySet1(_423_cf11, 3)
				(_nw65).ArraySet1((_this).F6(), 4)
				(_nw65).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_434_v35).Cardinality()), p0), 5)
				(_nw65).ArraySet1((_this).F6(), 6)
				(_nw65).ArraySet1((func() _dafny.Int {
					if (_this).Fm8((_this).F6(), (_434_v35).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm3(_dafny.IntOfInt64(130), false, globalState), _dafny.IntOfUint32((_434_v35).Cardinality()))).Uint32()).(bool), _425_v27, _431_v31, globalState) {
						return (_this).F6()
					}
					return _dafny.IntOfUint32((_431_v31).Cardinality())
				})(), 7)
				(_nw65).ArraySet1(_this.F16, 8)
				(_nw65).ArraySet1(p0, 9)
				(_nw65).ArraySet1(_this.F16, 10)
				(_nw65).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if (_436_v37).Contains(_435_v36) {
						return (_436_v37).Get(_435_v36).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_434_v35).Cardinality())
				})()), 11)
				(_nw65).ArraySet1((_424_v26).F12(), 12)
				(_nw65).ArraySet1(Companion_Default___.SafeDivisionInt(_423_cf11, p0), 13)
				_437_v38 = _nw65
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_437_v38), 0))
				_ = _index51
				(_437_v38).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-342), Companion_Default___.Fm3((_this).F6(), (_this).F4(), globalState)), (_index51).Int())
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_437_v38), 0))
				_ = _index52
				(_437_v38).ArraySet1(Companion_Default___.Fm3(((_this).F6()).Minus((_424_v26).F12()), (_this).F4(), globalState), (_index52).Int())
				var _439_v39 _dafny.Sequence
				_ = _439_v39
				_439_v39 = _dafny.SeqOf(_430_v30)
				var _440_v40 _dafny.Map
				_ = _440_v40
				_440_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_434_v35, _434_v35), _439_v39)
				var _441_v41 _dafny.Map
				_ = _441_v41
				_441_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F15()), _430_v30)
				_440_v40 = (_440_v40).Update(_434_v35, _dafny.SeqOf((func() _dafny.Set {
					if (_441_v41).Contains(_dafny.IntOfInt64(372)) {
						return (_441_v41).Get(_dafny.IntOfInt64(372)).(_dafny.Set)
					}
					return _430_v30
				})()))
				var _442_v42 bool
				_ = _442_v42
				_442_v42 = true
				var _443_v43 _dafny.MultiSet
				_ = _443_v43
				_443_v43 = _dafny.MultiSetOf(_442_v42, _442_v42, (_430_v30).IsProperSubsetOf(_dafny.SetOf((_this).F4(), (_this).F4(), _442_v42, (_this).Fm8((_437_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_437_v38), 0))).Int()).(_dafny.Int), (_this).F4(), _425_v27, _431_v31, globalState), _442_v42)))
				var _rhs46 _dafny.Int = (_424_v26).F12()
				_ = _rhs46
				var _rhs47 bool = false
				_ = _rhs47
				var _rhs48 _dafny.Set = func() _dafny.Set {
					var _coll30 = _dafny.NewBuilder()
					_ = _coll30
					for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(681), _dafny.IntOfInt64(893))); ; {
						_compr_30, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _444_v44 _dafny.Int
						_444_v44 = interface{}(_compr_30).(_dafny.Int)
						if ((_dafny.IntOfInt64(681)).Cmp(_444_v44) <= 0) && ((_444_v44).Cmp(_dafny.IntOfInt64(893)) < 0) {
							_coll30.Add((_444_v44).Times((_424_v26).F12()))
						}
					}
					return _coll30.ToSet()
				}()
				_ = _rhs48
				var _rhs49 _dafny.MultiSet = (_dafny.MultiSetOf((_this).F4(), false)).Update(!((_this).F4()) || (false), Companion_Default___.Abs((_dafny.IntOfUint32((_425_v27).Cardinality())).Plus((_424_v26).F12())))
				_ = _rhs49
				var _lhs28 *C1 = _this
				_ = _lhs28
				_lhs28.F16 = _rhs46
				_442_v42 = _rhs47
				_370_v0 = _rhs48
				_443_v43 = _rhs49
				var _445_v45 _dafny.Set
				_ = _445_v45
				_445_v45 = _370_v0
				var _446_v46 _dafny.Set
				_ = _446_v46
				_446_v46 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(710))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}(func(_447_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				})))
				var _448_v47 _dafny.Map
				_ = _448_v47
				_448_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_445_v45, (_446_v46).IsDisjointFrom(_446_v46))
				_448_v47 = (_448_v47).Update(_445_v45, (_this).F4())
			} else {
				var _449_v48 _dafny.Sequence
				_ = _449_v48
				_449_v48 = _dafny.UnicodeSeqOfUtf8Bytes("qrngovhln")
				_449_v48 = _dafny.Companion_Sequence_.Concatenate(_449_v48, _449_v48)
				var _450_v49 _dafny.MultiSet
				_ = _450_v49
				_450_v49 = _dafny.MultiSetOf(_449_v48, _449_v48, _449_v48)
				_423_cf11 = Companion_Default___.SafeModuloInt(Companion_Default___.Fm3((_this).F6(), (_this).F4(), globalState), (func() _dafny.Int {
					if (_450_v49).Contains(_449_v48) {
						return (_450_v49).Multiplicity(_449_v48)
					}
					return (_this).F15()
				})())
				var _451_v50 _dafny.Array
				_ = _451_v50
				var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
				_ = _nw66
				_451_v50 = _nw66
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_451_v50), 0))
				_ = _index53
				(_451_v50).ArraySet1((_dafny.Zero).Minus((_370_v0).Cardinality()), (_index53).Int())
				var _452_v51 _dafny.Set
				_ = _452_v51
				_452_v51 = _dafny.SetOf((_this).F4())
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_451_v50), 0))
				_ = _index54
				(_451_v50).ArraySet1(Companion_Default___.SafeModuloInt((_this).Fm7(globalState), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _452_v51)).Cardinality())), (_index54).Int())
				_423_cf11 = (_451_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_451_v50), 0))).Int()).(_dafny.Int)
				var _453_v52 _dafny.Array
				_ = _453_v52
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_6
				var _nw67 _dafny.Array
				_ = _nw67
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw67 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) _dafny.MultiSet = func(_454_i5 _dafny.Int) _dafny.MultiSet {
						return (_dafny.MultiSetOf(!((_this).F4()))).Intersection(_dafny.MultiSetOf((_this).F4(), (_this).F4()))
					}
					_ = _init6
					var _element0_6 = _init6(_dafny.Zero)
					_ = _element0_6
					_nw67 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
					(_nw67).ArraySet1(_element0_6, 0)
					var _nativeLen0_6 = (_len0_6).Int()
					_ = _nativeLen0_6
					for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
						(_nw67).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
					}
				}
				_453_v52 = _nw67
				var _455_v53 _dafny.MultiSet
				_ = _455_v53
				_455_v53 = _dafny.MultiSetOf((_this).F4())
				var _456_v54 _dafny.Map
				_ = _456_v54
				_456_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _455_v53)
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_453_v52), 0))
				_ = _index55
				(_453_v52).ArraySet1((func() _dafny.MultiSet {
					if !((_this).F4()) {
						return _dafny.MultiSetOf((_this).F4())
					}
					return (func() _dafny.MultiSet {
						if (_456_v54).Contains((_this).F4()) {
							return (_456_v54).Get((_this).F4()).(_dafny.MultiSet)
						}
						return _455_v53
					})()
				})(), (_index55).Int())
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_453_v52), 0))
				_ = _index56
				(_453_v52).ArraySet1((_455_v53).Difference((_dafny.MultiSetOf((_this).F4())).Difference(_455_v53)), (_index56).Int())
			}
			var _457_v55 _dafny.Sequence
			_ = _457_v55
			_457_v55 = _dafny.UnicodeSeqOfUtf8Bytes("dohmu")
			var _458_v56 _dafny.Map
			_ = _458_v56
			_458_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_457_v55, (_this).F6())
			var _459_v57 _dafny.Sequence
			_ = _459_v57
			_459_v57 = _dafny.SeqOf(_458_v56)
			var _460_v58 _dafny.Array
			_ = _460_v58
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw68
			_460_v58 = _nw68
			var _461_v59 _dafny.Int
			_ = _461_v59
			var _out12 _dafny.Int
			_ = _out12
			_out12 = (_this).M11((_this).F4(), !((_this).F4()), ((_459_v57).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_459_v57).Cardinality()))).Uint32()).(_dafny.Map)).Update(_457_v55, _423_cf11), _460_v58, globalState)
			_461_v59 = _out12
			var _462_v60 _dafny.Array
			_ = _462_v60
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_7
			var _nw69 _dafny.Array
			_ = _nw69
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw69 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Sequence = (func(_463_v55 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_464_i6 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_463_v55, _463_v55)
					}
				})(_457_v55)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw69 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw69).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw69).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_462_v60 = _nw69
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_462_v60), 0))
			_ = _index57
			(_462_v60).ArraySet1(_457_v55, (_index57).Int())
			var _465_v61 _dafny.Map
			_ = _465_v61
			_465_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _423_cf11)
			var _466_v62 _dafny.Map
			_ = _466_v62
			_466_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-87), _465_v61)
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_462_v60), 0))
			_ = _index58
			(_462_v60).ArraySet1(Companion_Default___.Fm37((_this).F4(), (_this).F4(), p0, !((func() _dafny.Map {
				if (_466_v62).Contains(p0) {
					return (_466_v62).Get(p0).(_dafny.Map)
				}
				return (_465_v61).Update((_this).Fm9((_this).F4(), true, _423_cf11, globalState), (_this).F15())
			})()).Contains((_this).F15()), globalState), (_index58).Int())
			var _467_v63 _dafny.Array
			_ = _467_v63
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_8
			var _nw70 _dafny.Array
			_ = _nw70
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw70 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.MultiSet = func(_468_i7 _dafny.Int) _dafny.MultiSet {
					return (_dafny.MultiSetOf(Companion_D4_.Create_DC12_(_dafny.CodePoint('f')))).Intersection(_dafny.MultiSetOf(Companion_D4_.Create_DC12_(_dafny.CodePoint('l'))))
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw70 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw70).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw70).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_467_v63 = _nw70
			_467_v63 = _467_v63
		} else if _source10.Is_DC8() {
			var _469___mcc_h4 _dafny.Sequence = _source10.Get_().(D3_DC8).Cf7
			_ = _469___mcc_h4
			var _470_cf7 _dafny.Sequence = _469___mcc_h4
			_ = _470_cf7
			var _471_v64 _dafny.Array
			_ = _471_v64
			var _nw71 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw71
			_471_v64 = _nw71
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_471_v64), 0))
			_ = _index59
			(_471_v64).ArraySet1((_this.F16).Cmp(_dafny.IntOfInt64(548)) > 0, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_471_v64), 0))
			_ = _index60
			(_471_v64).ArraySet1((_this).F4(), (_index60).Int())
			var _472_v65 _dafny.Sequence
			_ = _472_v65
			_472_v65 = _dafny.UnicodeSeqOfUtf8Bytes("dv")
			_472_v65 = _dafny.UnicodeSeqOfUtf8Bytes("qssuchryd")
			var _473_v66 *C0
			_ = _473_v66
			var _nw72 *C0 = New_C0_()
			_ = _nw72
			_nw72.Ctor__(p0)
			_473_v66 = _nw72
			var _474_v67 _dafny.Array
			_ = _474_v67
			var _nw73 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw73
			_474_v67 = _nw73
			var _475_v68 _dafny.Set
			_ = _475_v68
			_475_v68 = _dafny.SetOf(_474_v67, _474_v67)
			var _476_v69 _dafny.Set
			_ = _476_v69
			_476_v69 = _475_v68
			var _477_v70 _dafny.Map
			_ = _477_v70
			_477_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_476_v69))
			var _478_v71 _dafny.Sequence
			_ = _478_v71
			_478_v71 = _dafny.SeqOf(_475_v68)
			_477_v70 = (_477_v70).Update(!((_this).F4()), (func() _dafny.Set {
				if (_this).F4() {
					return _475_v68
				}
				return (_478_v71).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_478_v71).Cardinality()))).Uint32()).(_dafny.Set)
			})())
		} else {
			var _479___mcc_h5 D3 = _source10.Get_().(D3_DC11).Cf12
			_ = _479___mcc_h5
			var _480_cf12 D3 = _479___mcc_h5
			_ = _480_cf12
			var _source11 D11 = Companion_D11_.Create_DC24_(!((_this).F4()), (_this).F4(), Companion_Default___.SafeModuloInt(_this.F16, _dafny.IntOfInt64(-625)))
			_ = _source11
			if _source11.Is_DC24() {
				var _481___mcc_h6 bool = _source11.Get_().(D11_DC24).Cf36
				_ = _481___mcc_h6
				var _482___mcc_h7 bool = _source11.Get_().(D11_DC24).Cf37
				_ = _482___mcc_h7
				var _483___mcc_h8 _dafny.Int = _source11.Get_().(D11_DC24).Cf38
				_ = _483___mcc_h8
				var _484_cf38 _dafny.Int = _483___mcc_h8
				_ = _484_cf38
				var _485_cf37 bool = _482___mcc_h7
				_ = _485_cf37
				var _486_cf36 bool = _481___mcc_h6
				_ = _486_cf36
				(_this).F16 = _484_cf38
				_486_cf36 = _486_cf36
				var _487_v72 _dafny.Array
				_ = _487_v72
				var _nw74 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw74
				_487_v72 = _nw74
				var _488_v73 D2
				_ = _488_v73
				_488_v73 = Companion_D2_.Create_DC5_(_487_v72)
				_488_v73 = _488_v73
				var _489_v74 _dafny.Array
				_ = _489_v74
				var _nwElement0_16 _dafny.Int = _dafny.IntOfInt64(445)
				_ = _nwElement0_16
				var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(2))
				_ = _nw75
				(_nw75).ArraySet1(_nwElement0_16, 0)
				(_nw75).ArraySet1((_this).F15(), 1)
				_489_v74 = _nw75
				var _490_v75 _dafny.Sequence
				_ = _490_v75
				_490_v75 = _dafny.SeqOf(_489_v74)
				_490_v75 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_489_v74), _dafny.Companion_Sequence_.Update(_490_v75, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_490_v75).Cardinality()))).Uint32(), _489_v74)), _dafny.Companion_Sequence_.Concatenate(_490_v75, _490_v75))
			} else if _source11.Is_DC25() {
				var _491___mcc_h9 bool = _source11.Get_().(D11_DC25).Cf39
				_ = _491___mcc_h9
				var _492___mcc_h10 _dafny.Int = _source11.Get_().(D11_DC25).Cf40
				_ = _492___mcc_h10
				var _493_cf40 _dafny.Int = _492___mcc_h10
				_ = _493_cf40
				var _494_cf39 bool = _491___mcc_h9
				_ = _494_cf39
				var _495_v76 _dafny.CodePoint
				_ = _495_v76
				_495_v76 = _dafny.CodePoint('b')
				var _496_v77 _dafny.Map
				_ = _496_v77
				_496_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("odwhwaejb"), (_dafny.MultiSetOf(_494_cf39)).Cardinality())
				var _497_v79 _dafny.Sequence
				_ = _497_v79
				_497_v79 = _dafny.UnicodeSeqOfUtf8Bytes("dqtrskur")
				var _498_v80 _dafny.Set
				_ = _498_v80
				_498_v80 = _dafny.SetOf(_497_v79)
				var _499_v81 _dafny.Array
				_ = _499_v81
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_9
				var _nw76 _dafny.Array
				_ = _nw76
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw76 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) bool = func(_500_i9 _dafny.Int) bool {
						return !((_this).F4())
					}
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw76 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw76).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw76).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_499_v81 = _nw76
				var _501_v82 _dafny.Int
				_ = _501_v82
				var _out13 _dafny.Int
				_ = _out13
				_out13 = (_this).M11(_dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm37(_494_cf39, (_this).F4(), (_this).F6(), false, globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(906))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}(func(_502_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				}))), Companion_Default___.Fm0(_495_v76, globalState), (_496_v77).Merge(func() _dafny.Map {
					var _coll31 = _dafny.NewMapBuilder()
					_ = _coll31
					for _iter32 := _dafny.Iterate((_498_v80).Elements()); ; {
						_compr_31, _ok32 := _iter32()
						if !_ok32 {
							break
						}
						var _503_v78 _dafny.Sequence
						_503_v78 = interface{}(_compr_31).(_dafny.Sequence)
						if (_498_v80).Contains(_503_v78) {
							_coll31.Add(_503_v78, (_dafny.SetOf((_this).F4(), _494_cf39, _494_cf39)).Cardinality())
						}
					}
					return _coll31.ToMap()
				}()), _499_v81, globalState)
				_501_v82 = _out13
				_497_v79 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(812))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_504_v76 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_505_i10 _dafny.Int) _dafny.CodePoint {
						return _504_v76
					}
				})(_495_v76)))
				var _506_v83 _dafny.Sequence
				_ = _506_v83
				_506_v83 = _dafny.SeqOf(_this.F16)
				var _507_v84 _dafny.Map
				_ = _507_v84
				_507_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_506_v83), _497_v79)
				var _508_v85 bool
				_ = _508_v85
				var _509_v86 _dafny.Int
				_ = _509_v86
				var _out14 bool
				_ = _out14
				var _out15 _dafny.Int
				_ = _out15
				_out14, _out15 = Companion_Default___.M0(_507_v84, _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("lcxrsang"), _497_v79), (_371_v1).Intersection(_371_v1), !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_497_v79, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_497_v79).Cardinality()))).Uint32(), _495_v76), _495_v76), globalState)
				_508_v85 = _out14
				_509_v86 = _out15
				var _510_v87 _dafny.Sequence
				_ = _510_v87
				_510_v87 = _dafny.SeqOf(_494_cf39)
				var _511_v88 _dafny.Map
				_ = _511_v88
				_511_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_509_v86, _495_v76)
				var _512_v89 _dafny.Map
				_ = _512_v89
				_512_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_509_v86, _493_cf40)
				var _513_v90 _dafny.Map
				_ = _513_v90
				_513_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_512_v89).Cardinality(), _493_cf40)
				var _514_v91 _dafny.Sequence
				_ = _514_v91
				_514_v91 = _dafny.SeqOf(Companion_Default___.Fm40((_this).F4(), _508_v85, _509_v86, globalState), _510_v87, _510_v87)
				_508_v85 = !(((_371_v1).Intersection(_dafny.MultiSetFromSeq(Companion_Default___.Fm39(_497_v79, Companion_D12_.Create_DC27_(_514_v91), _494_cf39, globalState)))).IsProperSubsetOf((_371_v1).Difference((_this).Fm6(Companion_Default___.Fm38((_this).F15(), p0, _508_v85, _dafny.IntOfUint32((_510_v87).Cardinality()), globalState), Companion_D0_.Create_DC0_((_this).F15()), _511_v88, (_513_v90).Cardinality(), globalState))))
			} else if _source11.Is_DC26() {
				var _515___mcc_h11 bool = _source11.Get_().(D11_DC26).Cf41
				_ = _515___mcc_h11
				var _516_cf41 bool = _515___mcc_h11
				_ = _516_cf41
				(_this).F16 = Companion_Default___.SafeDivisionInt((_this).F6(), (_this).F15())
				var _517_v92 _dafny.Array
				_ = _517_v92
				var _nw77 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw77
				_517_v92 = _nw77
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_517_v92), 0))
				_ = _index61
				(_517_v92).ArraySet1((_this.F16).Cmp(p0) <= 0, (_index61).Int())
				var _518_v93 _dafny.Sequence
				_ = _518_v93
				_518_v93 = _dafny.SeqOf(_dafny.IntOfInt64(106), (_this).F15())
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_517_v92), 0))
				_ = _index62
				(_517_v92).ArraySet1(_dafny.Companion_Sequence_.Equal(_518_v93, _518_v93), (_index62).Int())
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_517_v92), 0))
				_ = _index63
				(_517_v92).ArraySet1(((_517_v92).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_517_v92), 0))).Int()).(bool)) == ((_517_v92).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_517_v92), 0))).Int()).(bool)), (_index63).Int())
				var _519_v94 D0
				_ = _519_v94
				_519_v94 = Companion_D0_.Create_DC0_((_this).F15())
				var _520_v95 _dafny.Sequence
				_ = _520_v95
				_520_v95 = _dafny.UnicodeSeqOfUtf8Bytes("r")
				var _521_v96 _dafny.Map
				_ = _521_v96
				_521_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _520_v95)
				var _522_v97 _dafny.Sequence
				_ = _522_v97
				_522_v97 = _dafny.SeqOf(_520_v95, (func() _dafny.Sequence {
					if (_521_v96).Contains(_dafny.IntOfInt64(-706)) {
						return (_521_v96).Get(_dafny.IntOfInt64(-706)).(_dafny.Sequence)
					}
					return _520_v95
				})())
				var _523_v98 _dafny.CodePoint
				_ = _523_v98
				_523_v98 = _dafny.CodePoint('y')
				var _pat_let_tv6 = p0
				_ = _pat_let_tv6
				var _pat_let_tv7 = p0
				_ = _pat_let_tv7
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_517_v92), 0))
				_ = _index64
				var _rhs50 D0 = func(_pat_let9_0 D0) D0 {
					return func(_524_dt__update__tmp_h2 D0) D0 {
						return func(_pat_let10_0 _dafny.Int) D0 {
							return func(_525_dt__update_hcf0_h0 _dafny.Int) D0 {
								return Companion_D0_.Create_DC0_(_525_dt__update_hcf0_h0)
							}(_pat_let10_0)
						}((_pat_let_tv6).Plus(_pat_let_tv7))
					}(_pat_let9_0)
				}(_519_v94)
				_ = _rhs50
				var _rhs51 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_522_v97, _522_v97)
				_ = _rhs51
				var _rhs52 bool = (_this).F4()
				_ = _rhs52
				var _rhs53 _dafny.CodePoint = _523_v98
				_ = _rhs53
				var _lhs29 _dafny.Array = _517_v92
				_ = _lhs29
				var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_517_v92), 0))
				_ = _lhs30
				_519_v94 = _rhs50
				_522_v97 = _rhs51
				(_lhs29).ArraySet1(_rhs52, (_lhs30).Int())
				_523_v98 = _rhs53
			} else {
				var _526___mcc_h12 _dafny.Array = _source11.Get_().(D11_DC23).Cf35
				_ = _526___mcc_h12
				var _527_cf35 _dafny.Array = _526___mcc_h12
				_ = _527_cf35
				var _528_v99 _dafny.Array
				_ = _528_v99
				var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
				_ = _nw78
				_528_v99 = _nw78
				var _529_v100 _dafny.Sequence
				_ = _529_v100
				_529_v100 = _dafny.UnicodeSeqOfUtf8Bytes("yfnegvsi")
				var _530_v101 _dafny.Array
				_ = _530_v101
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_10
				var _nw79 _dafny.Array
				_ = _nw79
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw79 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Int = func(_531_i11 _dafny.Int) _dafny.Int {
						return (_531_i11).Minus(_this.F16)
					}
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw79 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw79).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw79).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_530_v101 = _nw79
				var _532_v102 _dafny.Map
				_ = _532_v102
				_532_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_529_v100, _530_v101)
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_528_v99), 0))
				_ = _index65
				(_528_v99).ArraySet1((_532_v102).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(252))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}(func(_533_i12 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				})), _530_v101), (_index65).Int())
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_528_v99), 0))
				_ = _index66
				(_528_v99).ArraySet1(_532_v102, (_index66).Int())
				var _534_v103 _dafny.Sequence
				_ = _534_v103
				_534_v103 = _dafny.SeqOf(_this.F16, _dafny.IntOfInt64(357))
				_529_v100 = Companion_Default___.Fm37((_this).F4(), (p0).Cmp(_dafny.IntOfUint32((_534_v103).Cardinality())) != 0, (_dafny.Zero).Minus((_this).F6()), (_this).F4(), globalState)
				var _535_v104 _dafny.Array
				_ = _535_v104
				var _nw80 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
				_ = _nw80
				_535_v104 = _nw80
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_535_v104), 0))
				_ = _index67
				(_535_v104).ArraySet1((_this).F4(), (_index67).Int())
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_535_v104), 0))
				_ = _index68
				(_535_v104).ArraySet1(((_this).F4()) && ((_this).F4()), (_index68).Int())
				var _536_v105 _dafny.Map
				_ = _536_v105
				_536_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3((_this).F15(), (_this).F4(), globalState), (_535_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_535_v104), 0))).Int()).(bool))
				var _537_v106 _dafny.Map
				_ = _537_v106
				_537_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_535_v104, (_this).F15())
				var _538_v107 _dafny.CodePoint
				_ = _538_v107
				_538_v107 = _dafny.CodePoint('f')
				var _539_v108 _dafny.Map
				_ = _539_v108
				_539_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_538_v107, (_this).F4())
				var _540_v109 D10
				_ = _540_v109
				_540_v109 = Companion_D10_.Create_DC22_(_529_v100, (_this).F4(), (func() bool {
					if (_539_v108).Contains(_538_v107) {
						return (_539_v108).Get(_538_v107).(bool)
					}
					return !((_this).F4())
				})(), _535_v104, p0)
				_536_v105 = (_536_v105).Update((func() _dafny.Int {
					if (_537_v106).Contains(_535_v104) {
						return (_537_v106).Get(_535_v104).(_dafny.Int)
					}
					return (func() _dafny.Int {
						if (_371_v1).Contains(_this.F16) {
							return (_371_v1).Multiplicity(_this.F16)
						}
						return (_this).F15()
					})()
				})(), (_540_v109).Dtor_cf32())
			}
			var _541_v110 _dafny.Array
			_ = _541_v110
			var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw81
			_541_v110 = _nw81
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_541_v110), 0))
			_ = _index69
			(_541_v110).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(739), (_this).F15()), (_index69).Int())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_541_v110), 0))
			_ = _index70
			(_541_v110).ArraySet1((_this).F6(), (_index70).Int())
			(_this).F16 = ((_this).F15()).Plus((_dafny.Zero).Minus(p0))
			var _542_v111 *C0
			_ = _542_v111
			var _nw82 *C0 = New_C0_()
			_ = _nw82
			_nw82.Ctor__((_this).F15())
			_542_v111 = _nw82
		}
		var _543_v112 _dafny.Sequence
		_ = _543_v112
		_543_v112 = _dafny.UnicodeSeqOfUtf8Bytes("jmhajedk")
		var _544_v113 _dafny.Set
		_ = _544_v113
		_544_v113 = _dafny.SetOf(_543_v112)
		var _545_i13 _dafny.Int
		_ = _545_i13
		_545_i13 = _dafny.Zero
		{
			for !((_544_v113).IsProperSubsetOf((_544_v113).Intersection(_544_v113))) {
				{
					if (_545_i13).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_545_i13 = (_545_i13).Plus(_dafny.One)
					(_this).F16 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_543_v112, _543_v112)).Cardinality())
					var _546_v114 *C0
					_ = _546_v114
					var _nw83 *C0 = New_C0_()
					_ = _nw83
					_nw83.Ctor__(_this.F16)
					_546_v114 = _nw83
					var _547_v115 _dafny.CodePoint
					_ = _547_v115
					_547_v115 = _dafny.CodePoint('n')
					var _548_v116 _dafny.Map
					_ = _548_v116
					_548_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _547_v115)
					var _549_v117 _dafny.Sequence
					_ = _549_v117
					_549_v117 = _dafny.SeqOf(((_this).F15()).Plus((_548_v116).Cardinality()), (_546_v114).F12(), (_this).F15())
					var _550_v118 _dafny.Array
					_ = _550_v118
					var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
					_ = _nw84
					_550_v118 = _nw84
					var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_550_v118), 0))
					_ = _index71
					(_550_v118).ArraySet1((_this).F15(), (_index71).Int())
					var _551_v119 _dafny.Set
					_ = _551_v119
					_551_v119 = _dafny.SetOf((_this).F4(), (_this).F4())
					var _552_v120 _dafny.Map
					_ = _552_v120
					_552_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16, _551_v119)
					var _553_v121 _dafny.Map
					_ = _553_v121
					_553_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), (_this).F4())
					var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_550_v118), 0))
					_ = _index72
					var _rhs54 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_549_v117, _dafny.SeqOf((_552_v120).Cardinality(), _dafny.IntOfInt64(798)))
					_ = _rhs54
					var _rhs55 _dafny.Int = (_546_v114).F12()
					_ = _rhs55
					var _rhs56 _dafny.Int = ((_this).F6()).Times((_553_v121).Cardinality())
					_ = _rhs56
					var _lhs31 *C1 = _this
					_ = _lhs31
					var _lhs32 _dafny.Array = _550_v118
					_ = _lhs32
					var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_550_v118), 0))
					_ = _lhs33
					_549_v117 = _rhs54
					_lhs31.F16 = _rhs55
					(_lhs32).ArraySet1(_rhs56, (_lhs33).Int())
					var _554_v122 _dafny.Array
					_ = _554_v122
					var _nw85 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
					_ = _nw85
					_554_v122 = _nw85
					var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_554_v122), 0))
					_ = _index73
					(_554_v122).ArraySet1((func() bool {
						if (_this).F4() {
							return true
						}
						return (_this).F4()
					})(), (_index73).Int())
					var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_550_v118), 0))
					_ = _index74
					var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_554_v122), 0))
					_ = _index75
					var _rhs57 _dafny.Int = (_dafny.Zero).Minus(((_this).F6()).Times((_550_v118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_550_v118), 0))).Int()).(_dafny.Int)))
					_ = _rhs57
					var _rhs58 bool = (_this).F4()
					_ = _rhs58
					var _lhs34 _dafny.Array = _550_v118
					_ = _lhs34
					var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_550_v118), 0))
					_ = _lhs35
					var _lhs36 _dafny.Array = _554_v122
					_ = _lhs36
					var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_554_v122), 0))
					_ = _lhs37
					(_lhs34).ArraySet1(_rhs57, (_lhs35).Int())
					(_lhs36).ArraySet1(_rhs58, (_lhs37).Int())
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		var _555_v123 _dafny.CodePoint
		_ = _555_v123
		_555_v123 = _dafny.CodePoint('a')
		var _556_v124 _dafny.Map
		_ = _556_v124
		_556_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), (_this).F4())
		var _557_v125 _dafny.MultiSet
		_ = _557_v125
		_557_v125 = _dafny.MultiSetOf((_this).F4())
		var _558_v126 D12
		_ = _558_v126
		_558_v126 = Companion_D12_.Create_DC28_(_dafny.IntOfUint32((Companion_Default___.Fm40((_this).F4(), (func() bool {
			if (_556_v124).Contains((func() _dafny.Int {
				if (_557_v125).Contains((_this).F4()) {
					return (_557_v125).Multiplicity((_this).F4())
				}
				return (_this).F6()
			})()) {
				return (_556_v124).Get((func() _dafny.Int {
					if (_557_v125).Contains((_this).F4()) {
						return (_557_v125).Multiplicity((_this).F4())
					}
					return (_this).F6()
				})()).(bool)
			}
			return (_this).F4()
		})(), _this.F16, globalState)).Cardinality()), (_dafny.Zero).Minus((_this).F6()), (_this).F4(), (_this).F4(), (_this).F15())
		var _pat_let_tv8 = _555_v123
		_ = _pat_let_tv8
		var _pat_let_tv9 = _555_v123
		_ = _pat_let_tv9
		_555_v123 = func(_source12 D12) _dafny.CodePoint {
			if _source12.Is_DC28() {
				var _559___mcc_h13 _dafny.Int = _source12.Get_().(D12_DC28).Cf43
				_ = _559___mcc_h13
				var _560___mcc_h14 _dafny.Int = _source12.Get_().(D12_DC28).Cf44
				_ = _560___mcc_h14
				var _561___mcc_h15 bool = _source12.Get_().(D12_DC28).Cf45
				_ = _561___mcc_h15
				var _562___mcc_h16 bool = _source12.Get_().(D12_DC28).Cf46
				_ = _562___mcc_h16
				var _563___mcc_h17 _dafny.Int = _source12.Get_().(D12_DC28).Cf47
				_ = _563___mcc_h17
				var _564_cf47 _dafny.Int = _563___mcc_h17
				_ = _564_cf47
				var _565_cf46 bool = _562___mcc_h16
				_ = _565_cf46
				var _566_cf45 bool = _561___mcc_h15
				_ = _566_cf45
				var _567_cf44 _dafny.Int = _560___mcc_h14
				_ = _567_cf44
				var _568_cf43 _dafny.Int = _559___mcc_h13
				_ = _568_cf43
				return _pat_let_tv8
			} else {
				var _569___mcc_h18 _dafny.Sequence = _source12.Get_().(D12_DC27).Cf42
				_ = _569___mcc_h18
				var _570_cf42 _dafny.Sequence = _569___mcc_h18
				_ = _570_cf42
				return _pat_let_tv9
			}
		}(_558_v126)
		var _571_v127 _dafny.Sequence
		_ = _571_v127
		_571_v127 = _dafny.SeqOf((_this).F15())
		var _572_v128 _dafny.Set
		_ = _572_v128
		_572_v128 = _dafny.SetOf(Companion_D6_.Create_DC16_((_this).F4(), (_this).F4(), _571_v127), Companion_Default___.Fm43(globalState))
		var _573_v129 _dafny.Array
		_ = _573_v129
		var _nwElement0_17 bool = (_this).F4()
		_ = _nwElement0_17
		var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(20))
		_ = _nw86
		(_nw86).ArraySet1(_nwElement0_17, 0)
		(_nw86).ArraySet1(!((func() bool {
			if (_556_v124).Contains((_this).F6()) {
				return (_556_v124).Get((_this).F6()).(bool)
			}
			return (_this).F4()
		})()), 1)
		(_nw86).ArraySet1(!(false), 2)
		(_nw86).ArraySet1((p0).Cmp((_dafny.Zero).Minus((_this).F15())) < 0, 3)
		(_nw86).ArraySet1((_this).F4(), 4)
		(_nw86).ArraySet1((_this).F4(), 5)
		(_nw86).ArraySet1((_this).F4(), 6)
		(_nw86).ArraySet1(!((_this).Fm8((_this).F15(), (_this).F4(), _543_v112, _571_v127, globalState)), 7)
		(_nw86).ArraySet1((_this).F4(), 8)
		(_nw86).ArraySet1((p0).Cmp((_this).F15()) <= 0, 9)
		(_nw86).ArraySet1(false, 10)
		(_nw86).ArraySet1((_this).F4(), 11)
		(_nw86).ArraySet1((Companion_Default___.Fm41(_dafny.IntOfUint32((_543_v112).Cardinality()), _555_v123, (_this).F4(), globalState)).IsSubsetOf(Companion_Default___.Fm41(_dafny.IntOfInt64(440), _555_v123, (_this).F4(), globalState)), 12)
		(_nw86).ArraySet1(Companion_Default___.Fm0(_555_v123, globalState), 13)
		(_nw86).ArraySet1((_this).F4(), 14)
		(_nw86).ArraySet1(false, 15)
		(_nw86).ArraySet1(!((_this).F4()) || ((_this).F4()), 16)
		(_nw86).ArraySet1(true, 17)
		(_nw86).ArraySet1((Companion_Default___.Fm42(_555_v123, _572_v128, globalState)).Contains(p0), 18)
		(_nw86).ArraySet1((_this).F4(), 19)
		_573_v129 = _nw86
		for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_573_v129), 0))); ; {
			_guard_loop_1, _ok33 := _iter33()
			if !_ok33 {
				break
			}
			var _574_i14 _dafny.Int
			_574_i14 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_574_i14).Sign() != -1) && ((_574_i14).Cmp(_dafny.ArrayLen((_573_v129), 0)) < 0)) {
				(_573_v129).ArraySet1((_this).F4(), (_574_i14).Int())
			}
		}
		var _575_i15 _dafny.Int
		_ = _575_i15
		_575_i15 = _dafny.Zero
		{
			for (_this).F4() {
				{
					if (_575_i15).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_575_i15 = (_575_i15).Plus(_dafny.One)
					(_this).F16 = (_dafny.Zero).Minus((_this.F16).Times(p0))
					var _576_v130 _dafny.Sequence
					_ = _576_v130
					_576_v130 = _dafny.SeqOf(false, (_this).F4(), (_this).F4())
					var _577_v131 D4
					_ = _577_v131
					_577_v131 = Companion_D4_.Create_DC13_(_this.F16, false, p0, (Companion_D4_.Create_DC13_(_dafny.IntOfInt64(362), true, _this.F16, _576_v130)).Dtor_cf17())
					var _578_v132 _dafny.Map
					_ = _578_v132
					_578_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_577_v131, (_this).F4())
					_578_v132 = (_578_v132).Update(_577_v131, (_this).F4())
					var _579_v133 *C0
					_ = _579_v133
					var _nw87 *C0 = New_C0_()
					_ = _nw87
					_nw87.Ctor__(_dafny.IntOfInt64(753))
					_579_v133 = _nw87
					_579_v133 = _579_v133
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _580_v134 bool
		_ = _580_v134
		_580_v134 = false
		var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_573_v129), 0))
		_ = _index76
		(_573_v129).ArraySet1(true, (_index76).Int())
		var _581_v135 _dafny.Array
		_ = _581_v135
		var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
		_ = _nw88
		_581_v135 = _nw88
		var _582_v136 _dafny.Set
		_ = _582_v136
		_582_v136 = _dafny.SetOf(_581_v135)
		var _583_v137 _dafny.Array
		_ = _583_v137
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_11
		var _nw89 _dafny.Array
		_ = _nw89
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw89 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) D0 = func(_584_i16 _dafny.Int) D0 {
				return Companion_D0_.Create_DC1_()
			}
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw89 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw89).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw89).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_583_v137 = _nw89
		var _585_v138 _dafny.Map
		_ = _585_v138
		_585_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_582_v136, _583_v137)
		var _586_v139 _dafny.Sequence
		_ = _586_v139
		_586_v139 = _dafny.SeqOf(_585_v138)
		var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_581_v135), 0))
		_ = _index77
		(_581_v135).ArraySet1(((_586_v139).Select((Companion_Default___.SafeIndex((_this).F15(), _dafny.IntOfUint32((_586_v139).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), (_index77).Int())
		var _587_v140 D11
		_ = _587_v140
		_587_v140 = Companion_D11_.Create_DC25_(_580_v134, p0)
		var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_573_v129), 0))
		_ = _index78
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_581_v135), 0))
		_ = _index79
		var _rhs59 _dafny.Int = (_587_v140).Dtor_cf40()
		_ = _rhs59
		var _rhs60 bool = (_580_v134) == (_580_v134)
		_ = _rhs60
		var _rhs61 bool = _580_v134
		_ = _rhs61
		var _rhs62 _dafny.Int = p0
		_ = _rhs62
		var _rhs63 bool = (_this).F4()
		_ = _rhs63
		var _lhs38 *C1 = _this
		_ = _lhs38
		var _lhs39 _dafny.Array = _573_v129
		_ = _lhs39
		var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_573_v129), 0))
		_ = _lhs40
		var _lhs41 _dafny.Array = _581_v135
		_ = _lhs41
		var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_581_v135), 0))
		_ = _lhs42
		_lhs38.F16 = _rhs59
		_580_v134 = _rhs60
		(_lhs39).ArraySet1(_rhs61, (_lhs40).Int())
		(_lhs41).ArraySet1(_rhs62, (_lhs42).Int())
		_580_v134 = _rhs63
		var _588_v141 _dafny.Map
		_ = _588_v141
		_588_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_555_v123, globalState), p0)
		r0 = _588_v141
		return r0
	}
}
func (_this *C1) M2(p0 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Array) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _589_v0 _dafny.Array
		_ = _589_v0
		var _nw90 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
		_ = _nw90
		_589_v0 = _nw90
		for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_589_v0), 0))); ; {
			_guard_loop_2, _ok34 := _iter34()
			if !_ok34 {
				break
			}
			var _590_i0 _dafny.Int
			_590_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_590_i0).Sign() != -1) && ((_590_i0).Cmp(_dafny.ArrayLen((_589_v0), 0)) < 0)) {
				(_589_v0).ArraySet1(!((_this).F4()) || ((_this).F4()), (_590_i0).Int())
			}
		}
		var _591_i1 _dafny.Int
		_ = _591_i1
		_591_i1 = _dafny.Zero
		{
			for true {
				{
					if (_591_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_591_i1 = (_591_i1).Plus(_dafny.One)
					var _592_v1 _dafny.Map
					_ = _592_v1
					_592_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), true)
					_592_v1 = (_592_v1).Merge(_592_v1)
					var _593_v2 _dafny.MultiSet
					_ = _593_v2
					_593_v2 = _dafny.MultiSetOf(p0)
					var _594_v3 _dafny.Map
					_ = _594_v3
					_594_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_this).F4())).IsDisjointFrom(_593_v2), (_this).F6())
					_594_v3 = (Companion_Default___.Fm44(globalState)).Merge((_594_v3).Merge(_594_v3))
					var _595_v4 _dafny.Sequence
					_ = _595_v4
					_595_v4 = _dafny.SeqOf(!((_this).F4()))
					(_this).F16 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _dafny.IntOfUint32((_595_v4).Cardinality()))).Cardinality()
					var _596_v5 *C0
					_ = _596_v5
					var _nw91 *C0 = New_C0_()
					_ = _nw91
					_nw91.Ctor__((_this).F6())
					_596_v5 = _nw91
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _597_v6 _dafny.Sequence
		_ = _597_v6
		_597_v6 = _dafny.SeqOf(p0)
		var _598_v7 D0
		_ = _598_v7
		_598_v7 = Companion_D0_.Create_DC2_(_dafny.IntOfUint32((_597_v6).Cardinality()))
		var _source13 D0 = _598_v7
		_ = _source13
		if _source13.Is_DC1() {
			var _599_v8 bool
			_ = _599_v8
			_599_v8 = false
			var _600_v9 _dafny.CodePoint
			_ = _600_v9
			_600_v9 = _dafny.CodePoint('n')
			_599_v8 = Companion_Default___.Fm0(_600_v9, globalState)
			var _601_v10 _dafny.Set
			_ = _601_v10
			_601_v10 = _dafny.SetOf((_this).F4(), _599_v8, _599_v8, false)
			if ((_601_v10).Intersection(_601_v10)).Equals(_601_v10) {
				var _602_v11 _dafny.Sequence
				_ = _602_v11
				_602_v11 = _dafny.SeqOf(_589_v0)
				var _603_v12 _dafny.Array
				_ = _603_v12
				var _nw92 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw92
				_603_v12 = _nw92
				var _604_v13 _dafny.Map
				_ = _604_v13
				_604_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !((_this).F4()))
				var _605_v14 _dafny.MultiSet
				_ = _605_v14
				_605_v14 = _dafny.MultiSetOf(_600_v9, _600_v9, _600_v9)
				var _606_v15 _dafny.Sequence
				_ = _606_v15
				_606_v15 = _dafny.SeqOf((_this).F15())
				var _607_v16 _dafny.Array
				_ = _607_v16
				var _nwElement0_18 _dafny.Int = (_this).F6()
				_ = _nwElement0_18
				var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(26))
				_ = _nw93
				(_nw93).ArraySet1(_nwElement0_18, 0)
				(_nw93).ArraySet1(_this.F16, 1)
				(_nw93).ArraySet1(Companion_Default___.Fm3((_this).F6(), p0, globalState), 2)
				(_nw93).ArraySet1(_this.F16, 3)
				(_nw93).ArraySet1(_this.F16, 4)
				(_nw93).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_603_v12, (_this).Fm7(globalState))).Cardinality(), 5)
				(_nw93).ArraySet1((_this).F6(), 6)
				(_nw93).ArraySet1((_this).F15(), 7)
				(_nw93).ArraySet1((_this).F6(), 8)
				(_nw93).ArraySet1((_604_v13).Cardinality(), 9)
				(_nw93).ArraySet1((_this).F6(), 10)
				(_nw93).ArraySet1((_this).F15(), 11)
				(_nw93).ArraySet1(_dafny.IntOfUint32((_597_v6).Cardinality()), 12)
				(_nw93).ArraySet1((_dafny.Zero).Minus((_this).F15()), 13)
				(_nw93).ArraySet1(_this.F16, 14)
				(_nw93).ArraySet1((_this).F15(), 15)
				(_nw93).ArraySet1((_this).F6(), 16)
				(_nw93).ArraySet1(_this.F16, 17)
				(_nw93).ArraySet1((_this).F15(), 18)
				(_nw93).ArraySet1((_this).F15(), 19)
				(_nw93).ArraySet1((_605_v14).Cardinality(), 20)
				(_nw93).ArraySet1(_dafny.IntOfUint32((_606_v15).Cardinality()), 21)
				(_nw93).ArraySet1(_this.F16, 22)
				(_nw93).ArraySet1(_this.F16, 23)
				(_nw93).ArraySet1(_dafny.IntOfInt64(-935), 24)
				(_nw93).ArraySet1(_dafny.IntOfInt64(642), 25)
				_607_v16 = _nw93
				var _608_v17 _dafny.MultiSet
				_ = _608_v17
				_608_v17 = _dafny.MultiSetOf(_607_v16, _607_v16)
				var _609_v18 _dafny.Map
				_ = _609_v18
				_609_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_589_v0), _602_v11)), ((func() _dafny.Int {
					if (_608_v17).Contains(_603_v12) {
						return (_608_v17).Multiplicity(_603_v12)
					}
					return (_606_v15).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_606_v15).Cardinality()))).Uint32()).(_dafny.Int)
				})()).Plus(_dafny.IntOfInt64(907)))
				var _610_v19 _dafny.Map
				_ = _610_v19
				_610_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(155), _this.F16)
				_609_v18 = (_609_v18).Update((_this).F4(), (func() _dafny.Int {
					if (_610_v19).Contains(_this.F16) {
						return (_610_v19).Get(_this.F16).(_dafny.Int)
					}
					return _this.F16
				})())
				var _611_v20 D3
				_ = _611_v20
				_611_v20 = Companion_D3_.Create_DC8_(Companion_Default___.Fm40(p0, _599_v8, _dafny.IntOfUint32((Companion_Default___.Fm37(_599_v8, (_this).F4(), (_this).F6(), (_this).F4(), globalState)).Cardinality()), globalState))
				var _612_v21 D3
				_ = _612_v21
				_612_v21 = Companion_D3_.Create_DC11_(_611_v20)
				_612_v21 = _612_v21
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_589_v0), 0))
				_ = _index80
				(_589_v0).ArraySet1(p0, (_index80).Int())
				var _613_v22 _dafny.Sequence
				_ = _613_v22
				_613_v22 = _dafny.UnicodeSeqOfUtf8Bytes("shvveca")
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_589_v0), 0))
				_ = _index81
				var _rhs64 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_613_v22, _613_v22)).Cardinality())))
				_ = _rhs64
				var _rhs65 _dafny.Int = Companion_Default___.Fm3((_this).F15(), _599_v8, globalState)
				_ = _rhs65
				var _rhs66 bool = _599_v8
				_ = _rhs66
				var _rhs67 bool = _599_v8
				_ = _rhs67
				var _rhs68 bool = _599_v8
				_ = _rhs68
				var _lhs43 *C1 = _this
				_ = _lhs43
				var _lhs44 *C1 = _this
				_ = _lhs44
				var _lhs45 _dafny.Array = _589_v0
				_ = _lhs45
				var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_589_v0), 0))
				_ = _lhs46
				_lhs43.F16 = _rhs64
				_lhs44.F16 = _rhs65
				_599_v8 = _rhs66
				(_lhs45).ArraySet1(_rhs67, (_lhs46).Int())
				_599_v8 = _rhs68
				var _614_v23 _dafny.Array
				_ = _614_v23
				var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(5))
				_ = _nw94
				_614_v23 = _nw94
				var _615_v24 _dafny.Map
				_ = _615_v24
				_615_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_600_v9, (_this).F4())
				var _616_v25 _dafny.Sequence
				_ = _616_v25
				_616_v25 = _dafny.SeqOf(_614_v23, _614_v23)
				var _617_v26 _dafny.Array
				_ = _617_v26
				var _nwElement0_19 _dafny.Array = _614_v23
				_ = _nwElement0_19
				var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(20))
				_ = _nw95
				(_nw95).ArraySet1(_nwElement0_19, 0)
				(_nw95).ArraySet1(_614_v23, 1)
				(_nw95).ArraySet1((func() _dafny.Array {
					if (func() bool {
						if (_615_v24).Contains(_600_v9) {
							return (_615_v24).Get(_600_v9).(bool)
						}
						return (_589_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_589_v0), 0))).Int()).(bool)
					})() {
						return _614_v23
					}
					return _614_v23
				})(), 2)
				(_nw95).ArraySet1(_614_v23, 3)
				(_nw95).ArraySet1(_614_v23, 4)
				(_nw95).ArraySet1(_614_v23, 5)
				(_nw95).ArraySet1(_614_v23, 6)
				(_nw95).ArraySet1(_614_v23, 7)
				(_nw95).ArraySet1(_614_v23, 8)
				(_nw95).ArraySet1((_616_v25).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_616_v25).Cardinality()))).Uint32()).(_dafny.Array), 9)
				(_nw95).ArraySet1(_614_v23, 10)
				(_nw95).ArraySet1(_614_v23, 11)
				(_nw95).ArraySet1(_614_v23, 12)
				(_nw95).ArraySet1(_614_v23, 13)
				(_nw95).ArraySet1(_614_v23, 14)
				(_nw95).ArraySet1(_614_v23, 15)
				(_nw95).ArraySet1(_614_v23, 16)
				(_nw95).ArraySet1(_614_v23, 17)
				(_nw95).ArraySet1(_614_v23, 18)
				(_nw95).ArraySet1(_614_v23, 19)
				_617_v26 = _nw95
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_617_v26), 0))
				_ = _index82
				(_617_v26).ArraySet1(_614_v23, (_index82).Int())
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_617_v26), 0))
				_ = _index83
				(_617_v26).ArraySet1(_614_v23, (_index83).Int())
				_613_v22 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_613_v22, _613_v22), _613_v22)
			} else {
				var _618_v27 _dafny.Map
				_ = _618_v27
				_618_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F4())
				var _619_v28 _dafny.Map
				_ = _619_v28
				_619_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _618_v27)
				_619_v28 = (_619_v28).Update(p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(871), _599_v8))
				var _620_v29 _dafny.Map
				_ = _620_v29
				_620_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16, (_this).F15())
				(_this).F16 = (func() _dafny.Int {
					if (_620_v29).Contains((_this).F15()) {
						return (_620_v29).Get((_this).F15()).(_dafny.Int)
					}
					return _this.F16
				})()
				(_this).F16 = (_this).F6()
				var _621_v30 _dafny.MultiSet
				_ = _621_v30
				_621_v30 = _dafny.MultiSetOf((_this).F15(), _dafny.IntOfInt64(-639))
				var _622_v31 _dafny.Array
				_ = _622_v31
				var _nw96 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw96
				_622_v31 = _nw96
				var _623_v32 _dafny.Sequence
				_ = _623_v32
				_623_v32 = _dafny.SeqOf(_622_v31, _622_v31, _622_v31)
				var _rhs69 _dafny.Int = (_this.F16).Plus((_this).F6())
				_ = _rhs69
				var _rhs70 _dafny.MultiSet = _621_v30
				_ = _rhs70
				var _rhs71 bool = (_this).F4()
				_ = _rhs71
				var _rhs72 _dafny.Sequence = _623_v32
				_ = _rhs72
				var _lhs47 *C1 = _this
				_ = _lhs47
				_lhs47.F16 = _rhs69
				_621_v30 = _rhs70
				_599_v8 = _rhs71
				_623_v32 = _rhs72
				(_this).F16 = _this.F16
			}
			(_this).F16 = (_this.F16).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(591))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}((func(_624_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_625_i2 _dafny.Int) _dafny.CodePoint {
					return _624_v9
				}
			})(_600_v9)))).Cardinality()))
			_589_v0 = _589_v0
		} else if _source13.Is_DC2() {
			var _626___mcc_h0 _dafny.Int = _source13.Get_().(D0_DC2).Cf1
			_ = _626___mcc_h0
			var _627_cf1 _dafny.Int = _626___mcc_h0
			_ = _627_cf1
			var _628_v33 bool
			_ = _628_v33
			_628_v33 = false
			var _629_v34 _dafny.Map
			_ = _629_v34
			_629_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _589_v0)
			var _630_v35 _dafny.Sequence
			_ = _630_v35
			_630_v35 = _dafny.SeqOf(_589_v0, (func() _dafny.Array {
				if (_629_v34).Contains((_this).F4()) {
					return (_629_v34).Get((_this).F4()).(_dafny.Array)
				}
				return _589_v0
			})())
			var _631_v36 _dafny.Sequence
			_ = _631_v36
			_631_v36 = _dafny.UnicodeSeqOfUtf8Bytes("gx")
			var _632_v37 _dafny.Map
			_ = _632_v37
			_632_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_630_v35).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("htbf")).Cardinality()), _dafny.IntOfUint32((_630_v35).Cardinality()))).Uint32()).(_dafny.Array), Companion_Default___.Fm3(_dafny.IntOfUint32((_631_v36).Cardinality()), p0, globalState))
			var _633_v38 _dafny.Map
			_ = _633_v38
			_633_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_627_cf1, (_this).F15())
			var _634_v39 _dafny.Set
			_ = _634_v39
			_634_v39 = _dafny.SetOf(_633_v38, _633_v38)
			var _635_v40 _dafny.Map
			_ = _635_v40
			_635_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_628_v33, p0)
			var _636_v41 _dafny.Map
			_ = _636_v41
			_636_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _627_cf1)
			var _637_v42 _dafny.CodePoint
			_ = _637_v42
			_637_v42 = _dafny.CodePoint('r')
			var _638_v43 _dafny.Set
			_ = _638_v43
			_638_v43 = _dafny.SetOf(_637_v42, _637_v42)
			var _639_v44 _dafny.Array
			_ = _639_v44
			var _nwElement0_20 _dafny.Int = _this.F16
			_ = _nwElement0_20
			var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(24))
			_ = _nw97
			(_nw97).ArraySet1(_nwElement0_20, 0)
			(_nw97).ArraySet1((_635_v40).Cardinality(), 1)
			(_nw97).ArraySet1(_this.F16, 2)
			(_nw97).ArraySet1(_dafny.IntOfInt64(377), 3)
			(_nw97).ArraySet1(_627_cf1, 4)
			(_nw97).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_this).F6(), _627_cf1, (_this).F6(), (_629_v34).Cardinality(), _dafny.IntOfUint32((_631_v36).Cardinality()))).Cardinality()), 5)
			(_nw97).ArraySet1((_this).F15(), 6)
			(_nw97).ArraySet1(_dafny.IntOfInt64(495), 7)
			(_nw97).ArraySet1(_this.F16, 8)
			(_nw97).ArraySet1((_dafny.MultiSetOf(_this.F16, _this.F16)).Cardinality(), 9)
			(_nw97).ArraySet1((func() _dafny.Int {
				if (_636_v41).Contains((_this).F4()) {
					return (_636_v41).Get((_this).F4()).(_dafny.Int)
				}
				return _this.F16
			})(), 10)
			(_nw97).ArraySet1((_this).F6(), 11)
			(_nw97).ArraySet1(_627_cf1, 12)
			(_nw97).ArraySet1((_this).F6(), 13)
			(_nw97).ArraySet1(_this.F16, 14)
			(_nw97).ArraySet1((_638_v43).Cardinality(), 15)
			(_nw97).ArraySet1(_this.F16, 16)
			(_nw97).ArraySet1(_this.F16, 17)
			(_nw97).ArraySet1(_627_cf1, 18)
			(_nw97).ArraySet1(_627_cf1, 19)
			(_nw97).ArraySet1(_dafny.IntOfUint32((_631_v36).Cardinality()), 20)
			(_nw97).ArraySet1((_636_v41).Cardinality(), 21)
			(_nw97).ArraySet1(_dafny.IntOfUint32((_631_v36).Cardinality()), 22)
			(_nw97).ArraySet1(_dafny.IntOfInt64(180), 23)
			_639_v44 = _nw97
			var _640_v45 _dafny.Sequence
			_ = _640_v45
			_640_v45 = _dafny.SeqOf(_639_v44, _639_v44, _639_v44, _639_v44, _639_v44)
			var _641_v46 _dafny.Map
			_ = _641_v46
			_641_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_631_v36, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(62))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}((func(_642_v38 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_643_i3 _dafny.Int) _dafny.Map {
					return _642_v38
				}
			})(_633_v38)))).Cardinality()))
			var _644_v47 D1
			_ = _644_v47
			_644_v47 = Companion_D1_.Create_DC3_(p0)
			var _645_v48 _dafny.Map
			_ = _645_v48
			_645_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_644_v47, _627_cf1)
			var _646_v49 _dafny.Array
			_ = _646_v49
			var _nwElement0_21 _dafny.Int = Companion_Default___.SafeModuloInt(_this.F16, _dafny.IntOfUint32((_631_v36).Cardinality()))
			_ = _nwElement0_21
			var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(26))
			_ = _nw98
			(_nw98).ArraySet1(_nwElement0_21, 0)
			(_nw98).ArraySet1((_this).F6(), 1)
			(_nw98).ArraySet1((_dafny.Zero).Minus((_this).F15()), 2)
			(_nw98).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_597_v6).Cardinality()), Companion_Default___.Fm3((_634_v39).Cardinality(), _628_v33, globalState)), 3)
			(_nw98).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_627_cf1), _this.F16), 4)
			(_nw98).ArraySet1(_dafny.IntOfInt64(813), 5)
			(_nw98).ArraySet1(_dafny.IntOfUint32((_640_v45).Cardinality()), 6)
			(_nw98).ArraySet1((_this).F6(), 7)
			(_nw98).ArraySet1(_this.F16, 8)
			(_nw98).ArraySet1(_dafny.IntOfInt64(206), 9)
			(_nw98).ArraySet1(((_this).F15()).Times((_this).F6()), 10)
			(_nw98).ArraySet1(_dafny.IntOfInt64(-867), 11)
			(_nw98).ArraySet1(_this.F16, 12)
			(_nw98).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F15(), (_dafny.Zero).Minus((_dafny.Zero).Minus(_627_cf1))), 13)
			(_nw98).ArraySet1(Companion_Default___.SafeDivisionInt(_627_cf1, _dafny.IntOfInt64(915)), 14)
			(_nw98).ArraySet1(_627_cf1, 15)
			(_nw98).ArraySet1((_this).Fm9(p0, p0, (func() _dafny.Int {
				if (_641_v46).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(886))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_647_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_648_i4 _dafny.Int) _dafny.CodePoint {
						return _647_v42
					}
				})(_637_v42)))) {
					return (_641_v46).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(886))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg53 _dafny.Int) interface{} {
							return coer53(arg53)
						}
					}((func(_649_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_650_i4 _dafny.Int) _dafny.CodePoint {
							return _649_v42
						}
					})(_637_v42)))).(_dafny.Int)
				}
				return _this.F16
			})(), globalState), 16)
			(_nw98).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_597_v6).Cardinality()), _this.F16), 17)
			(_nw98).ArraySet1(_627_cf1, 18)
			(_nw98).ArraySet1((_this).F15(), 19)
			(_nw98).ArraySet1((_645_v48).Cardinality(), 20)
			(_nw98).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eadjey")).Cardinality()), 21)
			(_nw98).ArraySet1(_627_cf1, 22)
			(_nw98).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_597_v6).Cardinality()), _dafny.IntOfUint32((_631_v36).Cardinality())), 23)
			(_nw98).ArraySet1((_this).F15(), 24)
			(_nw98).ArraySet1((func() _dafny.Int {
				if (_636_v41).Contains(_628_v33) {
					return (_636_v41).Get(_628_v33).(_dafny.Int)
				}
				return _dafny.IntOfInt64(-222)
			})(), 25)
			_646_v49 = _nw98
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_639_v44), 0))
			_ = _index84
			(_639_v44).ArraySet1(_627_cf1, (_index84).Int())
			var _651_v50 _dafny.Sequence
			_ = _651_v50
			_651_v50 = _dafny.SeqOf(_dafny.IntOfInt64(676), _627_cf1)
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_639_v44), 0))
			_ = _index85
			var _rhs73 bool = !(((_this).F4()) == (p0))
			_ = _rhs73
			var _rhs74 bool = (_this).F4()
			_ = _rhs74
			var _rhs75 _dafny.Map = _632_v37
			_ = _rhs75
			var _rhs76 _dafny.Int = ((func() _dafny.Int {
				if (_636_v41).Contains(_628_v33) {
					return (_636_v41).Get(_628_v33).(_dafny.Int)
				}
				return _dafny.IntOfInt64(94)
			})()).Plus(Companion_Default___.SafeDivisionInt(_627_cf1, _dafny.IntOfUint32((_651_v50).Cardinality())))
			_ = _rhs76
			var _rhs77 _dafny.Sequence = _597_v6
			_ = _rhs77
			var _lhs48 _dafny.Array = _639_v44
			_ = _lhs48
			var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_639_v44), 0))
			_ = _lhs49
			_628_v33 = _rhs73
			_628_v33 = _rhs74
			_632_v37 = _rhs75
			(_lhs48).ArraySet1(_rhs76, (_lhs49).Int())
			_597_v6 = _rhs77
			var _rhs78 bool = _628_v33
			_ = _rhs78
			var _rhs79 _dafny.Sequence = _631_v36
			_ = _rhs79
			_628_v33 = _rhs78
			_631_v36 = _rhs79
			var _652_v51 _dafny.Map
			_ = _652_v51
			_652_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_630_v35, (_this).F4())
			var _653_v52 _dafny.Int
			_ = _653_v52
			var _out16 _dafny.Int
			_ = _out16
			_out16 = (_this).M11((func() bool {
				if (_652_v51).Contains(_630_v35) {
					return (_652_v51).Get(_630_v35).(bool)
				}
				return _628_v33
			})(), _dafny.Companion_Sequence_.IsProperPrefixOf(_651_v50, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-844))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg54 _dafny.Int) interface{} {
					return coer54(arg54)
				}
			}(func(_654_i5 _dafny.Int) _dafny.Int {
				return (_this).F6()
			}))), (_641_v46).Merge(_641_v46), _589_v0, globalState)
			_653_v52 = _out16
			if p0 {
				_628_v33 = (_597_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_631_v36).Cardinality()), _dafny.IntOfUint32((_597_v6).Cardinality()))).Uint32()).(bool)
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_589_v0), 0))
				_ = _index86
				(_589_v0).ArraySet1(_628_v33, (_index86).Int())
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_589_v0), 0))
				_ = _index87
				(_589_v0).ArraySet1(_628_v33, (_index87).Int())
				_628_v33 = (_this).F4()
				_635_v40 = (_635_v40).Update(p0, Companion_Default___.Fm0(_637_v42, globalState))
				_653_v52 = (_dafny.IntOfInt64(628)).Times((_dafny.Zero).Minus(((_dafny.Zero).Minus((_633_v38).Cardinality())).Times((_639_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_639_v44), 0))).Int()).(_dafny.Int))))
			} else {
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_639_v44), 0))
				_ = _index88
				(_639_v44).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_627_cf1, (_this).F15()), (_651_v50).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_651_v50).Cardinality()))).Uint32()).(_dafny.Int)), (_index88).Int())
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_589_v0), 0))
				_ = _index89
				(_589_v0).ArraySet1((_597_v6).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_597_v6).Cardinality()))).Uint32()).(bool), (_index89).Int())
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_589_v0), 0))
				_ = _index90
				(_589_v0).ArraySet1((func() bool {
					if Companion_Default___.Fm0(_637_v42, globalState) {
						return !(_628_v33)
					}
					return !(true)
				})(), (_index90).Int())
				var _655_v53 _dafny.MultiSet
				_ = _655_v53
				_655_v53 = _dafny.MultiSetOf(_628_v33, _628_v33)
				var _656_v54 D11
				_ = _656_v54
				_656_v54 = Companion_D11_.Create_DC25_((_589_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_589_v0), 0))).Int()).(bool), _this.F16)
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_589_v0), 0))
				_ = _index91
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_639_v44), 0))
				_ = _index92
				var _rhs80 bool = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F4(), p0), _dafny.SeqOf(true, (_this).F4())), (_dafny.MultiSetFromSeq(_597_v6)).IsProperSubsetOf(_655_v53))
				_ = _rhs80
				var _rhs81 _dafny.Int = ((_this).F15()).Times((_656_v54).Dtor_cf40())
				_ = _rhs81
				var _rhs82 _dafny.CodePoint = _637_v42
				_ = _rhs82
				var _lhs50 _dafny.Array = _589_v0
				_ = _lhs50
				var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_589_v0), 0))
				_ = _lhs51
				var _lhs52 _dafny.Array = _639_v44
				_ = _lhs52
				var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_639_v44), 0))
				_ = _lhs53
				(_lhs50).ArraySet1(_rhs80, (_lhs51).Int())
				(_lhs52).ArraySet1(_rhs81, (_lhs53).Int())
				r0 = _rhs82
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_589_v0), 0))
				_ = _index93
				var _rhs83 _dafny.Int = (_dafny.Zero).Minus(((_633_v38).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_639_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_639_v44), 0))).Int()).(_dafny.Int), _627_cf1))).Cardinality())
				_ = _rhs83
				var _rhs84 _dafny.Int = (_639_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_639_v44), 0))).Int()).(_dafny.Int)
				_ = _rhs84
				var _rhs85 _dafny.Array = _646_v49
				_ = _rhs85
				var _rhs86 bool = !(!(_628_v33))
				_ = _rhs86
				var _lhs54 _dafny.Array = _589_v0
				_ = _lhs54
				var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_589_v0), 0))
				_ = _lhs55
				_653_v52 = _rhs83
				_653_v52 = _rhs84
				_646_v49 = _rhs85
				(_lhs54).ArraySet1(_rhs86, (_lhs55).Int())
				_628_v33 = (_589_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_589_v0), 0))).Int()).(bool)
			}
		} else {
			var _657___mcc_h1 _dafny.Int = _source13.Get_().(D0_DC0).Cf0
			_ = _657___mcc_h1
			var _658_cf0 _dafny.Int = _657___mcc_h1
			_ = _658_cf0
			var _659_v55 *C0
			_ = _659_v55
			var _nw99 *C0 = New_C0_()
			_ = _nw99
			_nw99.Ctor__((_this.F16).Plus(_658_cf0))
			_659_v55 = _nw99
			var _660_v56 bool
			_ = _660_v56
			_660_v56 = true
			var _661_v57 _dafny.Array
			_ = _661_v57
			var _nw100 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw100
			_661_v57 = _nw100
			var _662_v58 _dafny.CodePoint
			_ = _662_v58
			_662_v58 = _dafny.CodePoint('c')
			var _663_v59 _dafny.Map
			_ = _663_v59
			_663_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_659_v55, p0)
			var _664_v60 _dafny.Sequence
			_ = _664_v60
			_664_v60 = _dafny.SeqOf(_662_v58)
			var _665_v61 _dafny.MultiSet
			_ = _665_v61
			_665_v61 = _dafny.MultiSetOf(_589_v0, _589_v0)
			var _666_v62 _dafny.Sequence
			_ = _666_v62
			_666_v62 = _dafny.SeqOf((_665_v61).Cardinality())
			var _667_v63 _dafny.Sequence
			_ = _667_v63
			_667_v63 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_661_v57, _662_v58)).Cardinality(), (_this).F6(), Companion_Default___.Fm3((_this).F15(), (_this).Fm8((_dafny.Zero).Minus(_658_cf0), !((func() bool {
				if (_663_v59).Contains(_659_v55) {
					return (_663_v59).Get(_659_v55).(bool)
				}
				return p0
			})()), _664_v60, _666_v62, globalState), globalState))
			var _668_v64 _dafny.Array
			_ = _668_v64
			var _nwElement0_22 _dafny.Int = (_dafny.Zero).Minus((_659_v55).F12())
			_ = _nwElement0_22
			var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(4))
			_ = _nw101
			(_nw101).ArraySet1(_nwElement0_22, 0)
			(_nw101).ArraySet1((_this).F15(), 1)
			(_nw101).ArraySet1((_667_v63).Select((Companion_Default___.SafeIndex(_658_cf0, _dafny.IntOfUint32((_667_v63).Cardinality()))).Uint32()).(_dafny.Int), 2)
			(_nw101).ArraySet1(_658_cf0, 3)
			_668_v64 = _nw101
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_661_v57), 0))
			_ = _index94
			(_661_v57).ArraySet1((_659_v55).F12(), (_index94).Int())
			var _669_v65 _dafny.MultiSet
			_ = _669_v65
			_669_v65 = _dafny.MultiSetOf(p0, _660_v56)
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_661_v57), 0))
			_ = _index95
			var _rhs87 bool = (func() bool {
				if _660_v56 {
					return (_669_v65).IsSubsetOf(_669_v65)
				}
				return (_this).F4()
			})()
			_ = _rhs87
			var _rhs88 _dafny.Int = (func() _dafny.Int {
				if p0 {
					return _this.F16
				}
				return _this.F16
			})()
			_ = _rhs88
			var _rhs89 _dafny.Int = _658_cf0
			_ = _rhs89
			var _lhs56 _dafny.Array = _661_v57
			_ = _lhs56
			var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_661_v57), 0))
			_ = _lhs57
			_660_v56 = _rhs87
			_658_cf0 = _rhs88
			(_lhs56).ArraySet1(_rhs89, (_lhs57).Int())
			var _670_v66 _dafny.Map
			_ = _670_v66
			_670_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_658_cf0, (_661_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_661_v57), 0))).Int()).(_dafny.Int))
			var _671_v67 _dafny.MultiSet
			_ = _671_v67
			_671_v67 = _dafny.MultiSetOf(_658_cf0, _dafny.IntOfInt64(-806), _dafny.IntOfUint32((_597_v6).Cardinality()), (func() _dafny.Int {
				if (_670_v66).Contains(_this.F16) {
					return (_670_v66).Get(_this.F16).(_dafny.Int)
				}
				return (_659_v55).F12()
			})())
			var _672_v68 _dafny.MultiSet
			_ = _672_v68
			_672_v68 = _671_v67
			var _673_v69 _dafny.Map
			_ = _673_v69
			_673_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_672_v68, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_661_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_661_v57), 0))).Int()).(_dafny.Int), (_this).F15())).Merge(_670_v66))
			_673_v69 = _673_v69
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_661_v57), 0))
			_ = _index96
			(_661_v57).ArraySet1(((_661_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_661_v57), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfUint32((_666_v62).Cardinality())), (_index96).Int())
		}
		var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_589_v0), 0))
		_ = _index97
		(_589_v0).ArraySet1((_this).F4(), (_index97).Int())
		var _674_v70 _dafny.Sequence
		_ = _674_v70
		_674_v70 = _dafny.UnicodeSeqOfUtf8Bytes("uehfebb")
		var _675_v71 _dafny.CodePoint
		_ = _675_v71
		_675_v71 = _dafny.CodePoint('v')
		var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_589_v0), 0))
		_ = _index98
		(_589_v0).ArraySet1((func() bool {
			if _dafny.Companion_Sequence_.IsProperPrefixOf(_674_v70, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(272))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg55 _dafny.Int) interface{} {
					return coer55(arg55)
				}
			}(func(_676_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}))) {
				return !_dafny.Companion_Sequence_.Contains(_674_v70, _675_v71)
			}
			return (_this).F4()
		})(), (_index98).Int())
		var _677_v73 _dafny.MultiSet
		_ = _677_v73
		_677_v73 = _dafny.MultiSetOf((_this).F6())
		var _hi2 _dafny.Int = (_this.F16).Minus(_this.F16)
		_ = _hi2
		for _678_i7 := (_dafny.Zero).Minus((func() _dafny.Map {
			var _coll33 = _dafny.NewMapBuilder()
			_ = _coll33
			for _iter36 := _dafny.Iterate((_677_v73).Elements()); ; {
				_compr_33, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _692_v72 _dafny.Int
				_692_v72 = interface{}(_compr_33).(_dafny.Int)
				if (_677_v73).Contains(_692_v72) {
					_coll33.Add((_692_v72).Plus((_dafny.Zero).Minus((_677_v73).Cardinality())), (func() _dafny.Int {
						if (_677_v73).Contains((_this).F6()) {
							return (_677_v73).Multiplicity((_this).F6())
						}
						return (_this).F6()
					})())
				}
			}
			return _coll33.ToMap()
		}()).Cardinality()); _678_i7.Cmp(_hi2) < 0; _678_i7 = _678_i7.Plus(_dafny.One) {
			var _679_v75 _dafny.Array
			_ = _679_v75
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_12
			var _nw102 _dafny.Array
			_ = _nw102
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw102 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Int = func(_680_i8 _dafny.Int) _dafny.Int {
					return (_680_i8).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(904)), func() _dafny.Set {
						var _coll32 = _dafny.NewBuilder()
						_ = _coll32
						for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(739), _dafny.IntOfInt64(930))); ; {
							_compr_32, _ok35 := _iter35()
							if !_ok35 {
								break
							}
							var _681_v74 _dafny.Int
							_681_v74 = interface{}(_compr_32).(_dafny.Int)
							if ((_dafny.IntOfInt64(739)).Cmp(_681_v74) <= 0) && ((_681_v74).Cmp(_dafny.IntOfInt64(930)) < 0) {
								_coll32.Add(Companion_Default___.SafeDivisionInt(_681_v74, _this.F16))
							}
						}
						return _coll32.ToSet()
					}(), _dafny.SetOf((_this).F6(), (_this).F15()))).Cardinality()))
				}
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw102 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw102).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw102).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_679_v75 = _nw102
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_679_v75), 0))
			_ = _index99
			(_679_v75).ArraySet1((_this).Fm9(p0, (_this).F4(), _678_i7, globalState), (_index99).Int())
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_679_v75), 0))
			_ = _index100
			(_679_v75).ArraySet1((_this).F15(), (_index100).Int())
			var _682_v76 *C0
			_ = _682_v76
			var _nw103 *C0 = New_C0_()
			_ = _nw103
			_nw103.Ctor__((_679_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_679_v75), 0))).Int()).(_dafny.Int))
			_682_v76 = _nw103
			var _683_v77 _dafny.MultiSet
			_ = _683_v77
			_683_v77 = _dafny.MultiSetOf(p0, (_589_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_589_v0), 0))).Int()).(bool))
			var _684_v78 _dafny.Sequence
			_ = _684_v78
			_684_v78 = _dafny.SeqOf(_683_v77, (_dafny.MultiSetOf(true, (_this).F4(), (_this).F4(), (_589_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_589_v0), 0))).Int()).(bool), (_this).F4())).Intersection(_683_v77))
			_684_v78 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_684_v78, Companion_Default___.Fm45(p0, (_this).F15(), (_589_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_589_v0), 0))).Int()).(bool), (_682_v76).F12(), globalState)), _dafny.SeqOf(_683_v77))
			var _685_v79 D1
			_ = _685_v79
			_685_v79 = Companion_D1_.Create_DC3_((_this).Fm8(_dafny.IntOfInt64(779), (_589_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_589_v0), 0))).Int()).(bool), _674_v70, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(520))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg56 _dafny.Int) interface{} {
					return coer56(arg56)
				}
			}((func(_686_v75 _dafny.Array) func(_dafny.Int) _dafny.Int {
				return func(_687_i9 _dafny.Int) _dafny.Int {
					return (_686_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_686_v75), 0))).Int()).(_dafny.Int)
				}
			})(_679_v75))), globalState))
			var _688_v80 _dafny.Array
			_ = _688_v80
			var _nw104 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw104
			_688_v80 = _nw104
			var _pat_let_tv10 = p0
			_ = _pat_let_tv10
			var _689_v81 _dafny.Map
			_ = _689_v81
			_689_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func(_pat_let11_0 D1) D1 {
				return func(_690_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let12_0 bool) D1 {
						return func(_691_dt__update_hcf2_h0 bool) D1 {
							return Companion_D1_.Create_DC3_(_691_dt__update_hcf2_h0)
						}(_pat_let12_0)
					}(_pat_let_tv10)
				}(_pat_let11_0)
			}(_685_v79)).Dtor_cf2(), _688_v80)
			_689_v81 = (_689_v81).Update(!(p0), _688_v80)
		}
		var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_589_v0), 0))
		_ = _index101
		(_589_v0).ArraySet1((_this).F4(), (_index101).Int())
		r0 = _675_v71
		var _693_v82 _dafny.Array
		_ = _693_v82
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_13
		var _nw105 _dafny.Array
		_ = _nw105
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw105 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) D1 = func(_694_i10 _dafny.Int) D1 {
				return Companion_D1_.Create_DC3_((_this).F4())
			}
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw105 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw105).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw105).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_693_v82 = _nw105
		r1 = _693_v82
		return r0, r1
	}
}
func (_this *C1) M11(p0 bool, p1 bool, p2 _dafny.Map, p3 _dafny.Array, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _695_v0 _dafny.Sequence
		_ = _695_v0
		_695_v0 = _dafny.UnicodeSeqOfUtf8Bytes("h")
		_695_v0 = _695_v0
		var _696_v1 bool
		_ = _696_v1
		_696_v1 = false
		var _arr0 _dafny.Array = _this.F5()
		_ = _arr0
		var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_this.F5()), 0))
		_ = _index102
		(_arr0).ArraySet1(p3, (_index102).Int())
		var _697_v2 _dafny.Sequence
		_ = _697_v2
		_697_v2 = _dafny.SeqOf(_696_v1)
		var _698_v3 D1
		_ = _698_v3
		_698_v3 = Companion_D1_.Create_DC4_((_this).F15(), _this.F16)
		var _pat_let_tv11 = p1
		_ = _pat_let_tv11
		var _pat_let_tv12 = p0
		_ = _pat_let_tv12
		var _arr1 _dafny.Array = _this.F5()
		_ = _arr1
		var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_this.F5()), 0))
		_ = _index103
		var _rhs90 bool = (_697_v2).Select((Companion_Default___.SafeIndex(((_this).F15()).Minus((_this).F6()), _dafny.IntOfUint32((_697_v2).Cardinality()))).Uint32()).(bool)
		_ = _rhs90
		var _rhs91 _dafny.Array = p3
		_ = _rhs91
		var _rhs92 bool = p0
		_ = _rhs92
		var _rhs93 bool = func(_source14 D1) bool {
			if _source14.Is_DC4() {
				var _699___mcc_h0 _dafny.Int = _source14.Get_().(D1_DC4).Cf3
				_ = _699___mcc_h0
				var _700___mcc_h1 _dafny.Int = _source14.Get_().(D1_DC4).Cf4
				_ = _700___mcc_h1
				var _701_cf4 _dafny.Int = _700___mcc_h1
				_ = _701_cf4
				var _702_cf3 _dafny.Int = _699___mcc_h0
				_ = _702_cf3
				return _pat_let_tv11
			} else {
				var _703___mcc_h2 bool = _source14.Get_().(D1_DC3).Cf2
				_ = _703___mcc_h2
				var _704_cf2 bool = _703___mcc_h2
				_ = _704_cf2
				return _pat_let_tv12
			}
		}(func(_pat_let13_0 D1) D1 {
			return func(_705_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let14_0 _dafny.Int) D1 {
					return func(_706_dt__update_hcf3_h0 _dafny.Int) D1 {
						return Companion_D1_.Create_DC4_(_706_dt__update_hcf3_h0, (_705_dt__update__tmp_h0).Dtor_cf4())
					}(_pat_let14_0)
				}(_dafny.IntOfInt64(-948))
			}(_pat_let13_0)
		}(_698_v3))
		_ = _rhs93
		var _rhs94 bool = false
		_ = _rhs94
		var _lhs58 _dafny.Array = _this.F5()
		_ = _lhs58
		var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_this.F5()), 0))
		_ = _lhs59
		_696_v1 = _rhs90
		(_lhs58).ArraySet1(_rhs91, (_lhs59).Int())
		_696_v1 = _rhs92
		_696_v1 = _rhs93
		_696_v1 = _rhs94
		var _707_v4 _dafny.CodePoint
		_ = _707_v4
		_707_v4 = _dafny.CodePoint('d')
		var _708_v5 _dafny.Map
		_ = _708_v5
		_708_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_707_v4, p0)
		var _709_i0 _dafny.Int
		_ = _709_i0
		_709_i0 = _dafny.Zero
		{
			for (func() bool {
				if !_dafny.Companion_Sequence_.Equal(_695_v0, _695_v0) {
					return ((_708_v5).Cardinality()).Cmp((_this).F6()) > 0
				}
				return p1
			})() {
				{
					if (_709_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_709_i0 = (_709_i0).Plus(_dafny.One)
					var _710_v6 _dafny.Array
					_ = _710_v6
					var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
					_ = _nw106
					_710_v6 = _nw106
					_710_v6 = _710_v6
					var _711_v7 _dafny.Map
					_ = _711_v7
					_711_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_dafny.Zero).Minus((_this).F6()))
					var _712_v8 *C0
					_ = _712_v8
					var _nw107 *C0 = New_C0_()
					_ = _nw107
					_nw107.Ctor__((func() _dafny.Int {
						if (_this).F4() {
							return (_711_v7).Cardinality()
						}
						return (_this).F6()
					})())
					_712_v8 = _nw107
					_696_v1 = _dafny.Companion_Sequence_.Equal(Companion_Default___.Fm40((_this).F4(), p1, (_712_v8).F12(), globalState), _dafny.SeqOf(_696_v1))
					var _713_v9 _dafny.Set
					_ = _713_v9
					_713_v9 = _dafny.SetOf((_696_v1) || ((_this).F4()), p0)
					_713_v9 = _713_v9
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		if false {
			var _714_v10 _dafny.Set
			_ = _714_v10
			_714_v10 = _dafny.SetOf(p3)
			var _715_v11 _dafny.Array
			_ = _715_v11
			var _nw108 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
			_ = _nw108
			_715_v11 = _nw108
			var _716_v12 _dafny.Array
			_ = _716_v12
			var _nwElement0_23 _dafny.Array = _715_v11
			_ = _nwElement0_23
			var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(12))
			_ = _nw109
			(_nw109).ArraySet1(_nwElement0_23, 0)
			(_nw109).ArraySet1(_715_v11, 1)
			(_nw109).ArraySet1(_715_v11, 2)
			(_nw109).ArraySet1(_715_v11, 3)
			(_nw109).ArraySet1(_715_v11, 4)
			(_nw109).ArraySet1(_715_v11, 5)
			(_nw109).ArraySet1(_715_v11, 6)
			(_nw109).ArraySet1(_715_v11, 7)
			(_nw109).ArraySet1(_715_v11, 8)
			(_nw109).ArraySet1(_715_v11, 9)
			(_nw109).ArraySet1(_715_v11, 10)
			(_nw109).ArraySet1(_715_v11, 11)
			_716_v12 = _nw109
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_716_v12), 0))
			_ = _index104
			(_716_v12).ArraySet1(_715_v11, (_index104).Int())
			var _717_v13 _dafny.Map
			_ = _717_v13
			_717_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(_this.F16)).Cardinality()), (_this).Fm9((_this).F4(), p1, (_this).F15(), globalState)), (_this).F15())
			var _718_v14 _dafny.Sequence
			_ = _718_v14
			_718_v14 = _dafny.SeqOf(_715_v11, _715_v11, _715_v11, _715_v11)
			var _719_v15 _dafny.MultiSet
			_ = _719_v15
			_719_v15 = _dafny.MultiSetOf(p0)
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_716_v12), 0))
			_ = _index105
			var _rhs95 _dafny.Set = _714_v10
			_ = _rhs95
			var _rhs96 _dafny.Array = (_718_v14).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_719_v15).Contains(!(false)) {
					return (_719_v15).Multiplicity(!(false))
				}
				return (_this).F15()
			})(), _dafny.IntOfUint32((_718_v14).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs96
			var _rhs97 _dafny.Map = (_717_v13).Merge(_717_v13)
			_ = _rhs97
			var _lhs60 _dafny.Array = _716_v12
			_ = _lhs60
			var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_716_v12), 0))
			_ = _lhs61
			_714_v10 = _rhs95
			(_lhs60).ArraySet1(_rhs96, (_lhs61).Int())
			_717_v13 = _rhs97
			(_this).F16 = (_this).Fm7(globalState)
			if (_this).F4() {
				var _720_v16 _dafny.Array
				_ = _720_v16
				var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw110
				_720_v16 = _nw110
				var _721_v17 _dafny.Map
				_ = _721_v17
				_721_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_695_v0).Cardinality()), _720_v16)
				_721_v17 = _721_v17
				var _722_v18 _dafny.MultiSet
				_ = _722_v18
				_722_v18 = _dafny.MultiSetOf((_this).F15())
				_722_v18 = ((_722_v18).Intersection(_dafny.MultiSetOf((_this).F15()))).Update(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_695_v0).Cardinality())), (_this).F15()), Companion_Default___.Abs((_this).F15()))
				var _arr2 _dafny.Array = _dafny.ArrayCastTo((_this.F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_this.F5()), 0))).Int()))
				_ = _arr2
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_dafny.ArrayCastTo((_this.F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_this.F5()), 0))).Int()))), 0))
				_ = _index106
				(_arr2).ArraySet1((_this).F4(), (_index106).Int())
				var _723_v19 _dafny.Map
				_ = _723_v19
				_723_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _695_v0)
				var _724_v20 _dafny.Sequence
				_ = _724_v20
				_724_v20 = _dafny.SeqOf((_dafny.Zero).Minus(_this.F16))
				var _725_v21 _dafny.Map
				_ = _725_v21
				_725_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfUint32((_724_v20).Cardinality())))
				var _726_v22 D12
				_ = _726_v22
				_726_v22 = Companion_D12_.Create_DC28_((_this).F6(), _this.F16, _696_v1, !(_696_v1), (func() _dafny.Int {
					if (_725_v21).Contains((_this).F4()) {
						return (_725_v21).Get((_this).F4()).(_dafny.Int)
					}
					return _dafny.IntOfInt64(146)
				})())
				var _arr3 _dafny.Array = _dafny.ArrayCastTo((_this.F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_this.F5()), 0))).Int()))
				_ = _arr3
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_dafny.ArrayCastTo((_this.F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_this.F5()), 0))).Int()))), 0))
				_ = _index107
				var _rhs98 _dafny.Int = (_723_v19).Cardinality()
				_ = _rhs98
				var _rhs99 _dafny.Int = (_726_v22).Dtor_cf44()
				_ = _rhs99
				var _rhs100 bool = p0
				_ = _rhs100
				var _lhs62 *C1 = _this
				_ = _lhs62
				var _lhs63 _dafny.Array = _dafny.ArrayCastTo((_this.F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_this.F5()), 0))).Int()))
				_ = _lhs63
				var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_dafny.ArrayCastTo((_this.F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_this.F5()), 0))).Int()))), 0))
				_ = _lhs64
				_lhs62.F16 = _rhs98
				r0 = _rhs99
				(_lhs63).ArraySet1(_rhs100, (_lhs64).Int())
				var _727_v23 _dafny.Sequence
				_ = _727_v23
				_727_v23 = _dafny.SeqOf(_695_v0, _695_v0)
				_696_v1 = !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_727_v23, _727_v23), _727_v23))
				var _arr4 _dafny.Array = _dafny.ArrayCastTo((_this.F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_this.F5()), 0))).Int()))
				_ = _arr4
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_dafny.ArrayCastTo((_this.F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_this.F5()), 0))).Int()))), 0))
				_ = _index108
				(_arr4).ArraySet1(!((_this).F4()), (_index108).Int())
			} else {
				r0 = ((_this).F15()).Times(_this.F16)
				var _728_v24 D0
				_ = _728_v24
				_728_v24 = Companion_D0_.Create_DC2_((_dafny.Zero).Minus((_this).F6()))
				_728_v24 = _728_v24
				(_this).F16 = _this.F16
				var _729_v25 _dafny.Map
				_ = _729_v25
				_729_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_697_v2, p0)
				var _730_v26 _dafny.Set
				_ = _730_v26
				_730_v26 = _dafny.SetOf((_this).F15())
				_729_v25 = (_729_v25).Update(_697_v2, (_730_v26).IsSubsetOf(_730_v26))
				r0 = _dafny.IntOfUint32((_695_v0).Cardinality())
			}
			var _731_v27 D0
			_ = _731_v27
			_731_v27 = Companion_D0_.Create_DC0_((_this).F6())
			var _732_v28 *C0
			_ = _732_v28
			var _nw111 *C0 = New_C0_()
			_ = _nw111
			_nw111.Ctor__(_dafny.IntOfUint32((_697_v2).Cardinality()))
			_732_v28 = _nw111
			var _733_v29 _dafny.Map
			_ = _733_v29
			_733_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_732_v28, _696_v1)
			var _734_v30 _dafny.Map
			_ = _734_v30
			_734_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _733_v29)
			var _735_v31 _dafny.Sequence
			_ = _735_v31
			_735_v31 = _dafny.SeqOf(_dafny.IntOfInt64(-159), (_dafny.Zero).Minus(_this.F16))
			var _rhs101 D0 = _731_v27
			_ = _rhs101
			var _rhs102 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_732_v28, (_this).F4())).Update(_732_v28, !((_this).F4()))).Merge((func() _dafny.Map {
				if (_734_v30).Contains((_this).F15()) {
					return (_734_v30).Get((_this).F15()).(_dafny.Map)
				}
				return _733_v29
			})())
			_ = _rhs102
			var _rhs103 _dafny.Int = (_this).F6()
			_ = _rhs103
			var _rhs104 bool = (_this).Fm8(((_this).F15()).Times((_732_v28).F12()), ((_735_v31).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_735_v31).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((_732_v28).F12()) <= 0, _695_v0, _735_v31, globalState)
			_ = _rhs104
			var _rhs105 bool = (p1) && (p0)
			_ = _rhs105
			_731_v27 = _rhs101
			_733_v29 = _rhs102
			r0 = _rhs103
			_696_v1 = _rhs104
			_696_v1 = _rhs105
			var _736_v32 _dafny.Map
			_ = _736_v32
			_736_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_732_v28).F12()), (func() _dafny.Int {
				if (_719_v15).Contains(!(_696_v1)) {
					return (_719_v15).Multiplicity(!(_696_v1))
				}
				return (_this).Fm7(globalState)
			})())
			var _737_v33 _dafny.Map
			_ = _737_v33
			_737_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16, _732_v28)
			var _738_v34 _dafny.Map
			_ = _738_v34
			_738_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_737_v33).Equals(_737_v33), _736_v32)
			_736_v32 = (func() _dafny.Map {
				if (_738_v34).Contains(p0) {
					return (_738_v34).Get(p0).(_dafny.Map)
				}
				return _736_v32
			})()
		} else {
			(_this).F16 = _this.F16
			var _739_v35 *C0
			_ = _739_v35
			var _nw112 *C0 = New_C0_()
			_ = _nw112
			_nw112.Ctor__(_dafny.IntOfInt64(520))
			_739_v35 = _nw112
			var _740_v36 _dafny.MultiSet
			_ = _740_v36
			_740_v36 = _dafny.MultiSetOf(p1)
			var _741_v37 D3
			_ = _741_v37
			_741_v37 = Companion_D3_.Create_DC10_((_740_v36).Cardinality())
			var _742_v38 _dafny.Map
			_ = _742_v38
			_742_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16, ((_741_v37).Dtor_cf11()).Cmp((_739_v35).F12()) == 0)
			var _743_v39 _dafny.Sequence
			_ = _743_v39
			_743_v39 = _dafny.SeqOf(_this.F16)
			var _744_v40 _dafny.Sequence
			_ = _744_v40
			_744_v40 = _dafny.SeqOf(_dafny.SeqOf((_739_v35).F12()), _743_v39, _743_v39, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-982))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg57 _dafny.Int) interface{} {
					return coer57(arg57)
				}
			}((func(_745_v35 *C0) func(_dafny.Int) _dafny.Int {
				return func(_746_i1 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus((_745_v35).F12())
				}
			})(_739_v35))), _743_v39)
			_742_v38 = (_742_v38).Update((_this).F6(), !_dafny.Companion_Sequence_.Equal(_744_v40, _744_v40))
			var _747_v41 _dafny.Array
			_ = _747_v41
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_14
			var _nw113 _dafny.Array
			_ = _nw113
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw113 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Int = func(_748_i3 _dafny.Int) _dafny.Int {
					return (_748_i3).Minus((_this).F6())
				}
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw113 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw113).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw113).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_747_v41 = _nw113
			var _749_v42 _dafny.Map
			_ = _749_v42
			_749_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wirva"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-993))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg58 _dafny.Int) interface{} {
					return coer58(arg58)
				}
			}(func(_750_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			}))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _747_v41))
			var _751_v43 _dafny.Map
			_ = _751_v43
			_751_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _747_v41)
			_749_v42 = (_749_v42).Update(_695_v0, (_751_v43).Merge(_751_v43))
			_696_v1 = (_this).F4()
		}
		var _752_v44 _dafny.Map
		_ = _752_v44
		_752_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F16)
		r0 = (_this.F16).Plus(Companion_Default___.SafeDivisionInt((_752_v44).Cardinality(), _dafny.IntOfInt64(-690)))
		var _753_v45 _dafny.Map
		_ = _753_v45
		_753_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_707_v4, (_this).F15())
		if (_753_v45).Equals((_753_v45).Merge((_753_v45).Update(_707_v4, _dafny.IntOfInt64(220)))) {
			(_this).F16 = Companion_Default___.SafeModuloInt(_this.F16, (_this).F15())
			var _754_v46 _dafny.Sequence
			_ = _754_v46
			_754_v46 = _dafny.SeqOf((_dafny.IntOfUint32((_695_v0).Cardinality())).Times(_dafny.IntOfInt64(668)), (_this.F16).Times((_this).F15()))
			var _755_v47 D12
			_ = _755_v47
			_755_v47 = Companion_D12_.Create_DC27_(_dafny.SeqOf(_697_v2))
			_754_v46 = Companion_Default___.Fm39(_695_v0, _755_v47, p1, globalState)
			var _756_v48 _dafny.MultiSet
			_ = _756_v48
			_756_v48 = _dafny.MultiSetOf(!(false), p0)
			var _757_v49 _dafny.MultiSet
			_ = _757_v49
			_757_v49 = _dafny.MultiSetOf(_this.F16, _this.F16, (_756_v48).Cardinality())
			var _758_v50 D9
			_ = _758_v50
			_758_v50 = Companion_D9_.Create_DC19_(_dafny.UnicodeSeqOfUtf8Bytes("imf"))
			var _759_v51 _dafny.Map
			_ = _759_v51
			_759_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_757_v49, (_758_v50).Dtor_cf25())
			var _760_v52 _dafny.Set
			_ = _760_v52
			_760_v52 = _dafny.SetOf(_this.F16)
			var _761_v53 bool
			_ = _761_v53
			var _762_v54 _dafny.Int
			_ = _762_v54
			var _out17 bool
			_ = _out17
			var _out18 _dafny.Int
			_ = _out18
			_out17, _out18 = Companion_Default___.M0(_759_v51, (_760_v52).IsSubsetOf(_dafny.SetOf((_this).F6())), _757_v49, _696_v1, globalState)
			_761_v53 = _out17
			_762_v54 = _out18
			var _763_v55 _dafny.Map
			_ = _763_v55
			_763_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_754_v46).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_695_v0).Cardinality()), _dafny.IntOfUint32((_754_v46).Cardinality()))).Uint32()).(_dafny.Int))
			var _764_v56 _dafny.Array
			_ = _764_v56
			var _nwElement0_24 _dafny.Int = (_dafny.IntOfUint32((_695_v0).Cardinality())).Plus((_dafny.Zero).Minus((_this).F6()))
			_ = _nwElement0_24
			var _nw114 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(6))
			_ = _nw114
			(_nw114).ArraySet1(_nwElement0_24, 0)
			(_nw114).ArraySet1(_dafny.IntOfInt64(398), 1)
			(_nw114).ArraySet1(_dafny.IntOfInt64(62), 2)
			(_nw114).ArraySet1(Companion_Default___.SafeModuloInt(_this.F16, (_dafny.Zero).Minus((_756_v48).Cardinality())), 3)
			(_nw114).ArraySet1((func() _dafny.Int {
				if (_763_v55).Contains((_this).F6()) {
					return (_763_v55).Get((_this).F6()).(_dafny.Int)
				}
				return _this.F16
			})(), 4)
			(_nw114).ArraySet1((_762_v54).Times(_this.F16), 5)
			_764_v56 = _nw114
			var _rhs106 _dafny.Array = _764_v56
			_ = _rhs106
			var _rhs107 bool = p1
			_ = _rhs107
			var _rhs108 _dafny.Int = (_this.F16).Plus((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(311))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg59 _dafny.Int) interface{} {
					return coer59(arg59)
				}
			}(func(_765_i4 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("csd")).Cardinality())
			}))).Cardinality())).Plus((Companion_Default___.Fm46(_707_v4, globalState)).Cardinality()))
			_ = _rhs108
			var _rhs109 _dafny.Int = (func() _dafny.Int {
				if (_756_v48).Contains(_696_v1) {
					return (_756_v48).Multiplicity(_696_v1)
				}
				return (_this).F6()
			})()
			_ = _rhs109
			var _lhs65 *C1 = _this
			_ = _lhs65
			var _lhs66 *C1 = _this
			_ = _lhs66
			_764_v56 = _rhs106
			_761_v53 = _rhs107
			_lhs65.F16 = _rhs108
			_lhs66.F16 = _rhs109
			if (_this).F4() {
				var _766_v57 bool
				_ = _766_v57
				var _767_v58 _dafny.Int
				_ = _767_v58
				var _out19 bool
				_ = _out19
				var _out20 _dafny.Int
				_ = _out20
				_out19, _out20 = Companion_Default___.M0((_759_v51).Update(_757_v49, _695_v0), p0, _757_v49, p1, globalState)
				_766_v57 = _out19
				_767_v58 = _out20
				var _768_v60 _dafny.Set
				_ = _768_v60
				_768_v60 = _dafny.SetOf(_757_v49)
				var _769_v61 bool
				_ = _769_v61
				var _770_v62 _dafny.Int
				_ = _770_v62
				var _out21 bool
				_ = _out21
				var _out22 _dafny.Int
				_ = _out22
				_out21, _out22 = Companion_Default___.M0(func() _dafny.Map {
					var _coll34 = _dafny.NewMapBuilder()
					_ = _coll34
					for _iter37 := _dafny.Iterate((_768_v60).Elements()); ; {
						_compr_34, _ok37 := _iter37()
						if !_ok37 {
							break
						}
						var _771_v59 _dafny.MultiSet
						_771_v59 = interface{}(_compr_34).(_dafny.MultiSet)
						if (_768_v60).Contains(_771_v59) {
							_coll34.Add(_771_v59, _dafny.UnicodeSeqOfUtf8Bytes("lgs"))
						}
					}
					return _coll34.ToMap()
				}(), ((_752_v44).Cardinality()).Cmp(_dafny.IntOfInt64(580)) >= 0, (_757_v49).Union(_757_v49), _696_v1, globalState)
				_769_v61 = _out21
				_770_v62 = _out22
				_696_v1 = p1
				_769_v61 = p1
				r0 = _762_v54
			} else {
				var _772_v63 _dafny.Map
				_ = _772_v63
				_772_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), p1)
				_752_v44 = (_752_v44).Update((_this).F4(), ((_this).F15()).Times((_this).Fm9(false, p1, (_772_v63).Cardinality(), globalState)))
				var _773_v64 _dafny.Map
				_ = _773_v64
				_773_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_761_v53, p0)
				var _774_v65 _dafny.Sequence
				_ = _774_v65
				_774_v65 = _dafny.SeqOf(_757_v49)
				var _775_v66 _dafny.Map
				_ = _775_v66
				_775_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_773_v64).Merge(_773_v64), ((_774_v65).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F16), _dafny.IntOfUint32((_774_v65).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality())
				_775_v66 = _775_v66
				var _776_v67 *C0
				_ = _776_v67
				var _nw115 *C0 = New_C0_()
				_ = _nw115
				_nw115.Ctor__((_this.F16).Plus(_dafny.IntOfInt64(843)))
				_776_v67 = _nw115
				_756_v48 = ((_756_v48).Union(_756_v48)).Union((_756_v48).Union(Companion_Default___.Fm47((_this).F6(), globalState)))
				_696_v1 = _761_v53
			}
		} else {
			_696_v1 = (p0) || (p1)
			_707_v4 = _707_v4
			var _777_v68 _dafny.Array
			_ = _777_v68
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_15
			var _nw116 _dafny.Array
			_ = _nw116
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw116 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Sequence = func(_778_i5 _dafny.Int) _dafny.Sequence {
					return (func() _dafny.Sequence {
						if (_this).F4() {
							return _dafny.SeqOf((_dafny.MultiSetOf((_this).F6(), _dafny.IntOfInt64(-594), _this.F16)).Cardinality(), (_dafny.Zero).Minus((_this).F15()))
						}
						return _dafny.SeqOf(_this.F16)
					})()
				}
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw116 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw116).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw116).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_777_v68 = _nw116
			var _779_v69 _dafny.Map
			_ = _779_v69
			_779_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _696_v1)
			var _780_v70 _dafny.Sequence
			_ = _780_v70
			_780_v70 = _dafny.SeqOf((_779_v69).Cardinality())
			var _781_v71 _dafny.Sequence
			_ = _781_v71
			_781_v71 = _dafny.SeqOf((_this).F15(), _dafny.IntOfUint32((_780_v70).Cardinality()), (_this).F6())
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_777_v68), 0))
			_ = _index109
			(_777_v68).ArraySet1(_781_v71, (_index109).Int())
			var _782_v72 _dafny.Map
			_ = _782_v72
			_782_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_707_v4, globalState), _707_v4)
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_777_v68), 0))
			_ = _index110
			(_777_v68).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F15()), _781_v71), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F15()), _781_v71)).Cardinality()))).Uint32(), (_782_v72).Cardinality()), _dafny.SeqOf(_this.F16, (_this).F15())), (_index110).Int())
			var _783_v73 _dafny.Map
			_ = _783_v73
			_783_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_695_v0, _695_v0), (_this).F6())
			_783_v73 = (_783_v73).Update(_695_v0, (_this.F16).Plus(_dafny.IntOfInt64(424)))
			_696_v1 = p0
		}
		r0 = _this.F16
		return r0
	}
}
func (_this *C1) F15() _dafny.Int {
	{
		return _this._f15
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f4  bool
	_f17 _dafny.Sequence
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f4 = false
	_this._f17 = _dafny.EmptySeq
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

func (_this *C2) F4() bool {
	return _this._f4
}
func (_this *C2) Ctor__(f17 _dafny.Sequence, f4 bool) {
	{
		(_this)._f17 = f17
		(_this)._f4 = f4
	}
}
func (_this *C2) Fm6(p0 D0, p1 D0, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetOf(_dafny.IntOfInt64(486)))
	}
}
func (_this *C2) Fm7(globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(831)
	}
}
func (_this *C2) Fm48(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return (Companion_D1_.Create_DC3_((_this).F4())).Dtor_cf2()
	}
}
func (_this *C2) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _hi3 _dafny.Int = p0
		_ = _hi3
		for _784_i0 := (_this).Fm7(globalState); _784_i0.Cmp(_hi3) < 0; _784_i0 = _784_i0.Plus(_dafny.One) {
			var _785_v0 _dafny.Array
			_ = _785_v0
			var _nw117 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw117
			_785_v0 = _nw117
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_16
			var _nw118 _dafny.Array
			_ = _nw118
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw118 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) bool = func(_786_i1 _dafny.Int) bool {
					return (_this).F4()
				}
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw118 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw118).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw118).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_785_v0 = _nw118
			var _787_v1 _dafny.Int
			_ = _787_v1
			_787_v1 = _dafny.IntOfInt64(805)
			_787_v1 = _dafny.IntOfInt64(-78)
			var _788_v2 _dafny.Set
			_ = _788_v2
			_788_v2 = _dafny.SetOf((_this).F4())
			var _789_v3 _dafny.Int
			_ = _789_v3
			var _790_v4 _dafny.Int
			_ = _790_v4
			var _out23 _dafny.Int
			_ = _out23
			var _out24 _dafny.Int
			_ = _out24
			_out23, _out24 = (_this).M12(_787_v1, (_788_v2).Difference(_788_v2), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_787_v1, (_788_v2).Cardinality()), (_this).F4(), globalState)
			_789_v3 = _out23
			_790_v4 = _out24
			if (_this).F4() {
				var _791_v5 bool
				_ = _791_v5
				_791_v5 = true
				_791_v5 = (_788_v2).IsProperSubsetOf(_dafny.SetOf((_this).F4(), (_this).F4(), (_this).F4()))
				var _792_v6 _dafny.Set
				_ = _792_v6
				_792_v6 = _dafny.SetOf(_784_i0, _789_v3)
				var _793_v7 _dafny.MultiSet
				_ = _793_v7
				_793_v7 = _dafny.MultiSetOf(false, _791_v5, _791_v5, _791_v5)
				var _794_v8 _dafny.CodePoint
				_ = _794_v8
				_794_v8 = _dafny.CodePoint('p')
				var _795_v9 _dafny.Set
				_ = _795_v9
				_795_v9 = _dafny.SetOf(_792_v6, (_792_v6).Intersection(_dafny.SetOf(_784_i0, (_793_v7).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).F17(), (Companion_Default___.SafeIndex(_790_v4, _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32(), _794_v8)).Cardinality()), ((_793_v7).Update(_791_v5, Companion_Default___.Abs((_dafny.Zero).Minus(p0)))).Cardinality())))
				var _796_v10 D11
				_ = _796_v10
				_796_v10 = Companion_D11_.Create_DC26_(_791_v5)
				var _797_v11 _dafny.Map
				_ = _797_v11
				_797_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_796_v10, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hjgsgbnw"), (_this).F17()))
				var _798_v12 _dafny.Array
				_ = _798_v12
				var _nw119 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(16))
				_ = _nw119
				_798_v12 = _nw119
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_798_v12), 0))
				_ = _index111
				(_798_v12).ArraySet1(Companion_D2_.Create_DC6_(), (_index111).Int())
				var _799_v14 _dafny.MultiSet
				_ = _799_v14
				_799_v14 = _dafny.MultiSetOf(_796_v10)
				var _800_v15 D2
				_ = _800_v15
				_800_v15 = Companion_D2_.Create_DC6_()
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_798_v12), 0))
				_ = _index112
				var _rhs110 _dafny.Set = (func() _dafny.Set {
					if false {
						return _795_v9
					}
					return (_795_v9).Difference(_795_v9)
				})()
				_ = _rhs110
				var _rhs111 _dafny.Map = ((_797_v11).Merge(_797_v11)).Merge((func() _dafny.Map {
					var _coll35 = _dafny.NewMapBuilder()
					_ = _coll35
					for _iter38 := _dafny.Iterate((_799_v14).Elements()); ; {
						_compr_35, _ok38 := _iter38()
						if !_ok38 {
							break
						}
						var _801_v13 D11
						_801_v13 = interface{}(_compr_35).(D11)
						if (_799_v14).Contains(_801_v13) {
							_coll35.Add(_801_v13, (_this).F17())
						}
					}
					return _coll35.ToMap()
				}()).Merge(_797_v11))
				_ = _rhs111
				var _rhs112 D2 = _800_v15
				_ = _rhs112
				var _rhs113 bool = (!((_this).F4())) && ((_787_v1).Cmp(_787_v1) >= 0)
				_ = _rhs113
				var _rhs114 bool = (_this).F4()
				_ = _rhs114
				var _lhs67 _dafny.Array = _798_v12
				_ = _lhs67
				var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_798_v12), 0))
				_ = _lhs68
				_795_v9 = _rhs110
				_797_v11 = _rhs111
				(_lhs67).ArraySet1(_rhs112, (_lhs68).Int())
				_791_v5 = _rhs113
				_791_v5 = _rhs114
				var _802_v16 *C0
				_ = _802_v16
				var _nw120 *C0 = New_C0_()
				_ = _nw120
				_nw120.Ctor__(_dafny.IntOfInt64(-980))
				_802_v16 = _nw120
				_791_v5 = !((func() bool {
					if _791_v5 {
						return (_this).F4()
					}
					return _791_v5
				})())
				var _803_v17 D2
				_ = _803_v17
				_803_v17 = Companion_D2_.Create_DC5_(_785_v0)
				var _804_v18 _dafny.Map
				_ = _804_v18
				_804_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_791_v5) || ((_this).F4()), (_803_v17).Dtor_cf5())
				_804_v18 = (_804_v18).Update(!(_791_v5) || (_791_v5), _785_v0)
			} else {
				var _805_v19 _dafny.MultiSet
				_ = _805_v19
				_805_v19 = _dafny.MultiSetOf(_784_i0)
				_805_v19 = (_805_v19).Intersection(_805_v19)
				var _806_v20 _dafny.Array
				_ = _806_v20
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_17
				var _nw121 _dafny.Array
				_ = _nw121
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw121 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.Int = (func(_807_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_808_i2 _dafny.Int) _dafny.Int {
							return (_808_i2).Minus(_807_v4)
						}
					})(_790_v4)
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
				_806_v20 = _nw121
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_806_v20), 0))
				_ = _index113
				(_806_v20).ArraySet1(_784_i0, (_index113).Int())
				var _809_v21 _dafny.Sequence
				_ = _809_v21
				_809_v21 = _dafny.SeqOf(_789_v3)
				var _810_v22 _dafny.Sequence
				_ = _810_v22
				_810_v22 = _dafny.SeqOf(_809_v21)
				var _811_v23 bool
				_ = _811_v23
				_811_v23 = false
				var _812_v24 _dafny.Sequence
				_ = _812_v24
				_812_v24 = _dafny.SeqOf((_this).F4(), _811_v23)
				var _813_v25 _dafny.Sequence
				_ = _813_v25
				_813_v25 = _dafny.SeqOf((_this).F17(), (_this).F17())
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_806_v20), 0))
				_ = _index114
				var _rhs115 _dafny.Int = (_this).Fm7(globalState)
				_ = _rhs115
				var _rhs116 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm49((_this).F17(), _812_v24, _789_v3, globalState), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_813_v25).Select((Companion_Default___.SafeIndex(_789_v3, _dafny.IntOfUint32((_813_v25).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm49((_this).F17(), _812_v24, _789_v3, globalState)).Cardinality()))).Uint32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(345))).Uint32(), func(coer60 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}((func(_814_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_815_i3 _dafny.Int) _dafny.Int {
						return _814_i0
					}
				})(_784_i0)))), _dafny.SeqOf(_809_v21))
				_ = _rhs116
				var _rhs117 bool = true
				_ = _rhs117
				var _rhs118 _dafny.Int = (_dafny.Zero).Minus((_789_v3).Plus((_dafny.Zero).Minus(_787_v1)))
				_ = _rhs118
				var _rhs119 bool = (_this).F4()
				_ = _rhs119
				var _lhs69 _dafny.Array = _806_v20
				_ = _lhs69
				var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_806_v20), 0))
				_ = _lhs70
				(_lhs69).ArraySet1(_rhs115, (_lhs70).Int())
				_810_v22 = _rhs116
				_811_v23 = _rhs117
				_790_v4 = _rhs118
				_811_v23 = _rhs119
				var _816_v26 _dafny.MultiSet
				_ = _816_v26
				_816_v26 = _dafny.MultiSetOf(_811_v23, _811_v23, false)
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(217), _dafny.ArrayLen((_806_v20), 0))
				_ = _index115
				(_806_v20).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if (_816_v26).Contains(_811_v23) {
						return (_816_v26).Multiplicity(_811_v23)
					}
					return _790_v4
				})()), (_index115).Int())
				_788_v2 = _dafny.SetOf(!((_this).F4()))
				var _817_v27 _dafny.Array
				_ = _817_v27
				var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.One)
				_ = _nw122
				_817_v27 = _nw122
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_817_v27), 0))
				_ = _index116
				(_817_v27).ArraySet1CodePoint(_dafny.CodePoint('l'), (_index116).Int())
				var _818_v28 _dafny.CodePoint
				_ = _818_v28
				_818_v28 = _dafny.CodePoint('c')
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_817_v27), 0))
				_ = _index117
				(_817_v27).ArraySet1CodePoint(_818_v28, (_index117).Int())
			}
		}
		var _819_v29 _dafny.Array
		_ = _819_v29
		var _nw123 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
		_ = _nw123
		_819_v29 = _nw123
		var _820_v30 D10
		_ = _820_v30
		_820_v30 = Companion_D10_.Create_DC22_(_dafny.UnicodeSeqOfUtf8Bytes("jdyfjcy"), ((_dafny.Zero).Minus(p0)).Cmp(_dafny.IntOfInt64(122)) == 0, !((_this).F4()), _819_v29, (_dafny.Zero).Minus(p0))
		var _source15 D10 = _820_v30
		_ = _source15
		if _source15.Is_DC22() {
			var _821___mcc_h0 _dafny.Sequence = _source15.Get_().(D10_DC22).Cf30
			_ = _821___mcc_h0
			var _822___mcc_h1 bool = _source15.Get_().(D10_DC22).Cf31
			_ = _822___mcc_h1
			var _823___mcc_h2 bool = _source15.Get_().(D10_DC22).Cf32
			_ = _823___mcc_h2
			var _824___mcc_h3 _dafny.Array = _source15.Get_().(D10_DC22).Cf33
			_ = _824___mcc_h3
			var _825___mcc_h4 _dafny.Int = _source15.Get_().(D10_DC22).Cf34
			_ = _825___mcc_h4
			var _826_cf34 _dafny.Int = _825___mcc_h4
			_ = _826_cf34
			var _827_cf33 _dafny.Array = _824___mcc_h3
			_ = _827_cf33
			var _828_cf32 bool = _823___mcc_h2
			_ = _828_cf32
			var _829_cf31 bool = _822___mcc_h1
			_ = _829_cf31
			var _830_cf30 _dafny.Sequence = _821___mcc_h0
			_ = _830_cf30
			_826_cf34 = (func() _dafny.Int {
				if _829_cf31 {
					return Companion_Default___.SafeDivisionInt(_826_cf34, p0)
				}
				return _dafny.IntOfInt64(569)
			})()
			var _831_v31 _dafny.Map
			_ = _831_v31
			_831_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_this).F4() {
					return _828_cf32
				}
				return (_this).F4()
			})(), _826_cf34)
			var _832_v32 _dafny.Sequence
			_ = _832_v32
			_832_v32 = _dafny.SeqOf(p0)
			var _833_v33 _dafny.Set
			_ = _833_v33
			_833_v33 = _dafny.SetOf(_828_cf32, true)
			var _834_v34 _dafny.Array
			_ = _834_v34
			var _nwElement0_25 _dafny.Int = p0
			_ = _nwElement0_25
			var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(9))
			_ = _nw124
			(_nw124).ArraySet1(_nwElement0_25, 0)
			(_nw124).ArraySet1(_826_cf34, 1)
			(_nw124).ArraySet1(_dafny.IntOfUint32((_832_v32).Cardinality()), 2)
			(_nw124).ArraySet1(_dafny.IntOfInt64(156), 3)
			(_nw124).ArraySet1(p0, 4)
			(_nw124).ArraySet1(p0, 5)
			(_nw124).ArraySet1(_dafny.IntOfUint32((_830_cf30).Cardinality()), 6)
			(_nw124).ArraySet1((_this).Fm7(globalState), 7)
			(_nw124).ArraySet1((_833_v33).Cardinality(), 8)
			_834_v34 = _nw124
			var _835_v35 _dafny.Map
			_ = _835_v35
			_835_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_834_v34, p0)
			_831_v31 = (_831_v31).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_834_v34, (_this).Fm7(globalState))).Equals((_835_v35).Update(_834_v34, _dafny.IntOfInt64(-360))), p0)
			var _836_v36 D2
			_ = _836_v36
			_836_v36 = Companion_D2_.Create_DC7_((_826_cf34).Plus(p0))
			var _rhs120 _dafny.Int = p0
			_ = _rhs120
			var _rhs121 bool = _828_cf32
			_ = _rhs121
			var _rhs122 D2 = _836_v36
			_ = _rhs122
			_826_cf34 = _rhs120
			_828_cf32 = _rhs121
			_836_v36 = _rhs122
			var _837_v37 *C0
			_ = _837_v37
			var _nw125 *C0 = New_C0_()
			_ = _nw125
			_nw125.Ctor__(p0)
			_837_v37 = _nw125
		} else {
			var _838___mcc_h5 _dafny.Map = _source15.Get_().(D10_DC21).Cf29
			_ = _838___mcc_h5
			var _839_cf29 _dafny.Map = _838___mcc_h5
			_ = _839_cf29
			var _840_v38 D2
			_ = _840_v38
			_840_v38 = Companion_D2_.Create_DC6_()
			var _source16 D2 = _840_v38
			_ = _source16
			if _source16.Is_DC6() {
				var _841_v39 _dafny.CodePoint
				_ = _841_v39
				_841_v39 = _dafny.CodePoint('w')
				var _842_v40 D9
				_ = _842_v40
				_842_v40 = Companion_D9_.Create_DC20_((_this).F4(), _841_v39, _dafny.IntOfUint32(((_this).F17()).Cardinality()))
				var _843_v41 _dafny.Sequence
				_ = _843_v41
				_843_v41 = _dafny.SeqOf(_842_v40)
				var _844_v42 _dafny.Sequence
				_ = _844_v42
				_844_v42 = _843_v41
				_843_v41 = (_843_v41)
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_819_v29), 0))
				_ = _index118
				(_819_v29).ArraySet1((func() bool {
					if (_this).F4() {
						return (_this).F4()
					}
					return (_this).F4()
				})(), (_index118).Int())
				var _845_v43 _dafny.Array
				_ = _845_v43
				var _nw126 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw126
				_845_v43 = _nw126
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_845_v43), 0))
				_ = _index119
				(_845_v43).ArraySet1(Companion_Default___.SafeDivisionInt(p0, p0), (_index119).Int())
				var _846_v44 _dafny.Array
				_ = _846_v44
				var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw127
				_846_v44 = _nw127
				var _847_v45 bool
				_ = _847_v45
				_847_v45 = true
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_819_v29), 0))
				_ = _index120
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_845_v43), 0))
				_ = _index121
				var _rhs123 bool = (_this).F4()
				_ = _rhs123
				var _rhs124 _dafny.Int = p0
				_ = _rhs124
				var _rhs125 _dafny.Array = _846_v44
				_ = _rhs125
				var _rhs126 bool = !(_847_v45)
				_ = _rhs126
				var _lhs71 _dafny.Array = _819_v29
				_ = _lhs71
				var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_819_v29), 0))
				_ = _lhs72
				var _lhs73 _dafny.Array = _845_v43
				_ = _lhs73
				var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_845_v43), 0))
				_ = _lhs74
				(_lhs71).ArraySet1(_rhs123, (_lhs72).Int())
				(_lhs73).ArraySet1(_rhs124, (_lhs74).Int())
				_846_v44 = _rhs125
				_847_v45 = _rhs126
				var _848_v46 _dafny.Array
				_ = _848_v46
				var _nw128 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw128
				_848_v46 = _nw128
				_848_v46 = _848_v46
				var _849_v47 _dafny.Sequence
				_ = _849_v47
				_849_v47 = _dafny.SeqOf((_819_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_819_v29), 0))).Int()).(bool))
				var _850_v48 _dafny.Sequence
				_ = _850_v48
				_850_v48 = _dafny.SeqOf(_dafny.IntOfUint32((_849_v47).Cardinality()), p0)
				var _851_v49 _dafny.MultiSet
				_ = _851_v49
				_851_v49 = _dafny.MultiSetOf(false, (_819_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_819_v29), 0))).Int()).(bool))
				var _852_v50 _dafny.Map
				_ = _852_v50
				_852_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_850_v48).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F17()).Cardinality()), _dafny.IntOfUint32((_850_v48).Cardinality()))).Uint32()).(_dafny.Int), (func() _dafny.Int {
					if (_819_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_819_v29), 0))).Int()).(bool) {
						return (_dafny.Zero).Minus(p0)
					}
					return (_851_v49).Cardinality()
				})())
				var _853_v51 _dafny.Set
				_ = _853_v51
				_853_v51 = _dafny.SetOf((_845_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_845_v43), 0))).Int()).(_dafny.Int))
				var _854_v52 _dafny.MultiSet
				_ = _854_v52
				_854_v52 = _dafny.MultiSetOf((_853_v51).Cardinality(), p0)
				_852_v50 = (_852_v50).Update((_845_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_845_v43), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
					if _847_v45 {
						return _dafny.IntOfInt64(460)
					}
					return (_854_v52).Cardinality()
				})())
			} else if _source16.Is_DC7() {
				var _855___mcc_h6 _dafny.Int = _source16.Get_().(D2_DC7).Cf6
				_ = _855___mcc_h6
				var _856_cf6 _dafny.Int = _855___mcc_h6
				_ = _856_cf6
				var _857_v53 _dafny.Sequence
				_ = _857_v53
				_857_v53 = _dafny.SeqOf(_dafny.SetOf((_this).F4(), (_this).F4(), !(true)))
				_857_v53 = _857_v53
				var _858_v54 _dafny.Sequence
				_ = _858_v54
				_858_v54 = _dafny.SeqOf((_this).F4())
				var _859_v55 _dafny.Sequence
				_ = _859_v55
				_859_v55 = _dafny.SeqOf((func() bool {
					if (_this).F4() {
						return (_858_v54).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_856_cf6), _dafny.IntOfUint32((_858_v54).Cardinality()))).Uint32()).(bool)
					}
					return false
				})())
				_859_v55 = _858_v54
				var _860_v56 _dafny.MultiSet
				_ = _860_v56
				_860_v56 = _dafny.MultiSetOf(p0, p0)
				var _861_v57 _dafny.Map
				_ = _861_v57
				_861_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
					if (_this).Fm48((_this).F4(), _856_cf6, p0, globalState) {
						return (_this).F17()
					}
					return Companion_Default___.Fm50(_dafny.IntOfInt64(892), p0, (_860_v56).Cardinality(), (_this).F4(), globalState)
				})(), p0)
				var _862_v58 _dafny.Sequence
				_ = _862_v58
				_862_v58 = _dafny.SeqOf(p0, p0)
				var _863_v59 _dafny.Map
				_ = _863_v59
				_863_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_862_v58).Cardinality()), (_this).F17())
				var _864_v60 _dafny.Sequence
				_ = _864_v60
				_864_v60 = _dafny.SeqOf((_863_v59).Cardinality())
				var _865_v61 _dafny.Sequence
				_ = _865_v61
				_865_v61 = _dafny.SeqOf(_dafny.IntOfInt64(325), _856_cf6, p0, (_864_v60).Select((Companion_Default___.SafeIndex(_856_cf6, _dafny.IntOfUint32((_864_v60).Cardinality()))).Uint32()).(_dafny.Int))
				_861_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(228))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}(func(_866_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				})), ((_865_v61).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_865_v61).Cardinality()))).Uint32()).(_dafny.Int)).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(761))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}(func(_867_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('x')
				}))).Cardinality())))
				var _868_v62 bool
				_ = _868_v62
				_868_v62 = false
				var _869_v63 _dafny.Array
				_ = _869_v63
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_18
				var _nw129 _dafny.Array
				_ = _nw129
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw129 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Int = (func(_870_cf6 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_871_i6 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_871_i6, _870_cf6)
						}
					})(_856_cf6)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw129 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw129).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw129).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_869_v63 = _nw129
				var _872_v64 _dafny.Map
				_ = _872_v64
				_872_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_869_v63), 0))
				_ = _index122
				(_869_v63).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf(_856_cf6)).Cardinality())).Times((func() _dafny.Int {
					if (_872_v64).Contains(p0) {
						return (_872_v64).Get(p0).(_dafny.Int)
					}
					return _856_cf6
				})()), (_index122).Int())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_869_v63), 0))
				_ = _index123
				var _rhs127 bool = _868_v62
				_ = _rhs127
				var _rhs128 _dafny.Int = (func() _dafny.Int {
					if (_872_v64).Contains(_856_cf6) {
						return (_872_v64).Get(_856_cf6).(_dafny.Int)
					}
					return p0
				})()
				_ = _rhs128
				var _rhs129 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("hasb"), (_this).F17())
				_ = _rhs129
				var _lhs75 _dafny.Array = _869_v63
				_ = _lhs75
				var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_869_v63), 0))
				_ = _lhs76
				_868_v62 = _rhs127
				(_lhs75).ArraySet1(_rhs128, (_lhs76).Int())
				_868_v62 = _rhs129
			} else {
				var _873___mcc_h7 _dafny.Array = _source16.Get_().(D2_DC5).Cf5
				_ = _873___mcc_h7
				var _874_cf5 _dafny.Array = _873___mcc_h7
				_ = _874_cf5
				var _875_v65 _dafny.Array
				_ = _875_v65
				var _nw130 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(10))
				_ = _nw130
				_875_v65 = _nw130
				var _876_v66 _dafny.Map
				_ = _876_v66
				_876_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F4())
				var _877_v67 _dafny.Set
				_ = _877_v67
				_877_v67 = _dafny.SetOf(_876_v66)
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_875_v65), 0))
				_ = _index124
				(_875_v65).ArraySet1((func() _dafny.Set {
					if (_this).F4() {
						return _877_v67
					}
					return _877_v67
				})(), (_index124).Int())
				var _878_v68 _dafny.MultiSet
				_ = _878_v68
				_878_v68 = _dafny.MultiSetOf((_this).F4())
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_875_v65), 0))
				_ = _index125
				(_875_v65).ArraySet1(func() _dafny.Set {
					var _coll36 = _dafny.NewBuilder()
					_ = _coll36
					for _iter39 := _dafny.Iterate(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_876_v66, _dafny.MultiSetOf(false))).Update(_876_v66, _878_v68)).Keys().Elements()); ; {
						_compr_36, _ok39 := _iter39()
						if !_ok39 {
							break
						}
						var _879_v69 _dafny.Map
						_879_v69 = interface{}(_compr_36).(_dafny.Map)
						if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_876_v66, _dafny.MultiSetOf(false))).Update(_876_v66, _878_v68)).Contains(_879_v69) {
							_coll36.Add(_879_v69)
						}
					}
					return _coll36.ToSet()
				}(), (_index125).Int())
				var _880_v70 _dafny.Map
				_ = _880_v70
				_880_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _881_v71 _dafny.Set
				_ = _881_v71
				_881_v71 = _dafny.SetOf((_this).F4(), (_this).F4())
				var _882_v72 _dafny.Array
				_ = _882_v72
				var _nwElement0_26 _dafny.Int = ((_880_v70).Update(p0, p0)).Cardinality()
				_ = _nwElement0_26
				var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(16))
				_ = _nw131
				(_nw131).ArraySet1(_nwElement0_26, 0)
				(_nw131).ArraySet1(p0, 1)
				(_nw131).ArraySet1(_dafny.IntOfInt64(245), 2)
				(_nw131).ArraySet1(p0, 3)
				(_nw131).ArraySet1((_881_v71).Cardinality(), 4)
				(_nw131).ArraySet1(p0, 5)
				(_nw131).ArraySet1(p0, 6)
				(_nw131).ArraySet1(p0, 7)
				(_nw131).ArraySet1(p0, 8)
				(_nw131).ArraySet1(p0, 9)
				(_nw131).ArraySet1((_839_cf29).Cardinality(), 10)
				(_nw131).ArraySet1(p0, 11)
				(_nw131).ArraySet1(p0, 12)
				(_nw131).ArraySet1(p0, 13)
				(_nw131).ArraySet1(_dafny.IntOfUint32(((_this).F17()).Cardinality()), 14)
				(_nw131).ArraySet1(((_880_v70).Merge(_880_v70)).Cardinality(), 15)
				_882_v72 = _nw131
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_882_v72), 0))
				_ = _index126
				(_882_v72).ArraySet1(p0, (_index126).Int())
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_882_v72), 0))
				_ = _index127
				(_882_v72).ArraySet1(_dafny.IntOfInt64(844), (_index127).Int())
				var _883_v73 *C0
				_ = _883_v73
				var _nw132 *C0 = New_C0_()
				_ = _nw132
				_nw132.Ctor__(((_882_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_882_v72), 0))).Int()).(_dafny.Int)).Times((_882_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_882_v72), 0))).Int()).(_dafny.Int)))
				_883_v73 = _nw132
				var _884_v74 _dafny.Map
				_ = _884_v74
				_884_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(74), _881_v71)
				var _885_v75 _dafny.Sequence
				_ = _885_v75
				_885_v75 = _dafny.SeqOf((_883_v73).F12(), (_884_v74).Cardinality())
				var _886_v76 _dafny.Sequence
				_ = _886_v76
				_886_v76 = _dafny.SeqOf(((_dafny.Zero).Minus(_dafny.IntOfUint32((_885_v75).Cardinality()))).Minus(p0), (_883_v73).F12())
				_886_v76 = _886_v76
			}
			var _887_v77 _dafny.Sequence
			_ = _887_v77
			_887_v77 = _dafny.SeqOf((_this).F4())
			var _888_v78 _dafny.Sequence
			_ = _888_v78
			_888_v78 = _dafny.SeqOf(_887_v77, _887_v77, _dafny.SeqOf((_this).F4(), (_this).F4()))
			_887_v77 = (Companion_D4_.Create_DC13_(p0, true, (_dafny.Zero).Minus(_dafny.IntOfUint32(((_this).F17()).Cardinality())), (_888_v78).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_888_v78).Cardinality()))).Uint32()).(_dafny.Sequence))).Dtor_cf17()
			var _889_v79 _dafny.Array
			_ = _889_v79
			var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw133
			_889_v79 = _nw133
			var _890_v80 _dafny.Map
			_ = _890_v80
			_890_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_889_v79, (_this).F4())
			_890_v80 = (_890_v80).Update(_889_v79, (p0).Cmp(p0) >= 0)
			if ((p0).Cmp(p0) <= 0) || ((p0).Cmp(p0) >= 0) {
				var _891_v81 _dafny.Array
				_ = _891_v81
				var _nwElement0_27 _dafny.Array = _819_v29
				_ = _nwElement0_27
				var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(4))
				_ = _nw134
				(_nw134).ArraySet1(_nwElement0_27, 0)
				(_nw134).ArraySet1(_819_v29, 1)
				(_nw134).ArraySet1(_819_v29, 2)
				(_nw134).ArraySet1(_819_v29, 3)
				_891_v81 = _nw134
				var _892_v82 *C1
				_ = _892_v82
				var _nw135 *C1 = New_C1_()
				_ = _nw135
				_nw135.Ctor__(p0, p0, _891_v81, (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bw")).Cardinality())).Times(_dafny.IntOfInt64(206)), (_this).F4())
				_892_v82 = _nw135
				var _893_v83 _dafny.Sequence
				_ = _893_v83
				_893_v83 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mxhimf")).Cardinality()))
				var _894_v84 _dafny.Map
				_ = _894_v84
				_894_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_893_v83, !((_this).F4()) || ((_this).F4()))
				var _895_v86 _dafny.Sequence
				_ = _895_v86
				_895_v86 = _dafny.SeqOf(_893_v83, _893_v83)
				_894_v84 = ((func() _dafny.Map {
					var _coll37 = _dafny.NewMapBuilder()
					_ = _coll37
					for _iter40 := _dafny.Iterate((_895_v86).Elements()); ; {
						_compr_37, _ok40 := _iter40()
						if !_ok40 {
							break
						}
						var _896_v85 _dafny.Sequence
						_896_v85 = interface{}(_compr_37).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_895_v86, _896_v85) {
							_coll37.Add(_896_v85, (_this).F4())
						}
					}
					return _coll37.ToMap()
				}()).Merge(_894_v84)).Merge((_894_v84).Merge(_894_v84))
				var _897_v87 *C0
				_ = _897_v87
				var _nw136 *C0 = New_C0_()
				_ = _nw136
				_nw136.Ctor__(p0)
				_897_v87 = _nw136
				var _898_v88 bool
				_ = _898_v88
				_898_v88 = false
				var _899_v89 _dafny.Array
				_ = _899_v89
				var _nwElement0_28 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_this).F17(), (_this).F17())
				_ = _nwElement0_28
				var _nw137 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(4))
				_ = _nw137
				(_nw137).ArraySet1(_nwElement0_28, 0)
				(_nw137).ArraySet1((_this).F17(), 1)
				(_nw137).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F17(), (_this).F17()), 2)
				(_nw137).ArraySet1((_this).F17(), 3)
				_899_v89 = _nw137
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_899_v89), 0))
				_ = _index128
				(_899_v89).ArraySet1((_this).F17(), (_index128).Int())
				var _900_v90 _dafny.CodePoint
				_ = _900_v90
				_900_v90 = _dafny.CodePoint('e')
				var _901_v91 _dafny.Map
				_ = _901_v91
				_901_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_900_v90, _819_v29)
				var _902_v92 _dafny.MultiSet
				_ = _902_v92
				_902_v92 = _dafny.MultiSetOf(_819_v29, (func() _dafny.Array {
					if (_901_v91).Contains(_dafny.CodePoint('o')) {
						return (_901_v91).Get(_dafny.CodePoint('o')).(_dafny.Array)
					}
					return _819_v29
				})(), _819_v29, _819_v29, (func() _dafny.Array {
					if _898_v88 {
						return _819_v29
					}
					return _819_v29
				})())
				var _903_v93 _dafny.Map
				_ = _903_v93
				_903_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_892_v82).F15(), _dafny.IntOfUint32(((_this).F17()).Cardinality()))
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_899_v89), 0))
				_ = _index129
				var _rhs130 _dafny.Int = (func() _dafny.Int {
					if (_902_v92).Contains(_819_v29) {
						return (_902_v92).Multiplicity(_819_v29)
					}
					return ((func() _dafny.Int {
						if (_903_v93).Contains((_dafny.Zero).Minus(_892_v82.F16)) {
							return (_903_v93).Get((_dafny.Zero).Minus(_892_v82.F16)).(_dafny.Int)
						}
						return (_893_v83).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_893_v83).Cardinality()))).Uint32()).(_dafny.Int)
					})()).Minus((_892_v82).F15())
				})()
				_ = _rhs130
				var _rhs131 bool = ((_897_v87).F12()).Cmp((_897_v87).F12()) != 0
				_ = _rhs131
				var _rhs132 bool = _898_v88
				_ = _rhs132
				var _rhs133 bool = _898_v88
				_ = _rhs133
				var _rhs134 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-804))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}((func(_904_v90 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_905_i7 _dafny.Int) _dafny.CodePoint {
						return _904_v90
					}
				})(_900_v90)))
				_ = _rhs134
				var _lhs77 *C1 = _892_v82
				_ = _lhs77
				var _lhs78 _dafny.Array = _899_v89
				_ = _lhs78
				var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_899_v89), 0))
				_ = _lhs79
				_lhs77.F16 = _rhs130
				_898_v88 = _rhs131
				_898_v88 = _rhs132
				_898_v88 = _rhs133
				(_lhs78).ArraySet1(_rhs134, (_lhs79).Int())
				_900_v90 = _dafny.CodePoint('x')
			} else {
				var _906_v94 _dafny.Int
				_ = _906_v94
				_906_v94 = _dafny.IntOfInt64(297)
				_906_v94 = p0
				var _907_v95 bool
				_ = _907_v95
				_907_v95 = false
				var _908_v96 D1
				_ = _908_v96
				_908_v96 = Companion_D1_.Create_DC4_(p0, p0)
				var _909_v97 _dafny.Array
				_ = _909_v97
				var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
				_ = _nw138
				_909_v97 = _nw138
				var _910_v98 *C1
				_ = _910_v98
				var _nw139 *C1 = New_C1_()
				_ = _nw139
				_nw139.Ctor__(p0, p0, _909_v97, _906_v94, (_this).F4())
				_910_v98 = _nw139
				var _911_v99 _dafny.Sequence
				_ = _911_v99
				_911_v99 = _dafny.SeqOf(_910_v98)
				var _912_v100 _dafny.Set
				_ = _912_v100
				_912_v100 = _dafny.SetOf((_this).F17(), (_this).F17())
				var _913_v101 _dafny.Map
				_ = _913_v101
				_913_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_912_v100, _911_v99)
				var _914_v102 _dafny.Map
				_ = _914_v102
				_914_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _907_v95)
				var _915_v103 _dafny.Set
				_ = _915_v103
				_915_v103 = _dafny.SetOf((_this).F4(), (func() bool {
					if (_914_v102).Contains(_907_v95) {
						return (_914_v102).Get(_907_v95).(bool)
					}
					return true
				})())
				var _916_v104 _dafny.CodePoint
				_ = _916_v104
				_916_v104 = _dafny.CodePoint('e')
				var _917_v105 _dafny.Map
				_ = _917_v105
				_917_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_916_v104, _915_v103)
				var _rhs135 _dafny.Int = p0
				_ = _rhs135
				var _rhs136 bool = !_dafny.Companion_Sequence_.Equal(_911_v99, (func() _dafny.Sequence {
					if (_913_v101).Contains(_912_v100) {
						return (_913_v101).Get(_912_v100).(_dafny.Sequence)
					}
					return _911_v99
				})())
				_ = _rhs136
				var _rhs137 D1 = _908_v96
				_ = _rhs137
				var _rhs138 bool = ((_915_v103).Union((func() _dafny.Set {
					if (_917_v105).Contains(_916_v104) {
						return (_917_v105).Get(_916_v104).(_dafny.Set)
					}
					return _dafny.SetOf(!(true))
				})())).IsSubsetOf(_dafny.SetOf(_907_v95, (_this).F4()))
				_ = _rhs138
				var _rhs139 _dafny.Int = p0
				_ = _rhs139
				_906_v94 = _rhs135
				_907_v95 = _rhs136
				_908_v96 = _rhs137
				_907_v95 = _rhs138
				_906_v94 = _rhs139
				_907_v95 = false
				var _918_v106 _dafny.Sequence
				_ = _918_v106
				_918_v106 = _dafny.SeqOf(_889_v79)
				var _919_v107 D2
				_ = _919_v107
				_919_v107 = Companion_D2_.Create_DC5_(_819_v29)
				_907_v95 = ((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_918_v106).Cardinality()), _906_v94))).Cmp(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_919_v107)).Cardinality())))) <= 0
				var _920_v108 _dafny.Map
				_ = _920_v108
				_920_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-14), (_910_v98).F15())
				var _921_v109 _dafny.Int
				_ = _921_v109
				var _922_v110 _dafny.Int
				_ = _922_v110
				var _out25 _dafny.Int
				_ = _out25
				var _out26 _dafny.Int
				_ = _out26
				_out25, _out26 = (_this).M12(p0, _915_v103, (_920_v108).Merge(_920_v108), (_this).F4(), globalState)
				_921_v109 = _out25
				_922_v110 = _out26
			}
		}
		var _923_v111 *C0
		_ = _923_v111
		var _nw140 *C0 = New_C0_()
		_ = _nw140
		_nw140.Ctor__(p0)
		_923_v111 = _nw140
		var _924_i8 _dafny.Int
		_ = _924_i8
		_924_i8 = _dafny.Zero
		{
			for ((_this).F4()) && ((_this).F4()) {
				{
					if (_924_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_924_i8 = (_924_i8).Plus(_dafny.One)
					var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_819_v29), 0))
					_ = _index130
					(_819_v29).ArraySet1((_this).F4(), (_index130).Int())
					var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_819_v29), 0))
					_ = _index131
					(_819_v29).ArraySet1((_this).F4(), (_index131).Int())
					var _925_v112 _dafny.Map
					_ = _925_v112
					_925_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !((_this).F4()))
					var _926_v113 _dafny.Map
					_ = _926_v113
					_926_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_819_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_819_v29), 0))).Int()).(bool))
					var _927_v114 _dafny.Set
					_ = _927_v114
					_927_v114 = _dafny.SetOf((_this).F4())
					_925_v112 = Companion_Default___.Fm51(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
						if (_926_v113).Contains(Companion_Default___.Fm0(_dafny.CodePoint('c'), globalState)) {
							return (_926_v113).Get(Companion_Default___.Fm0(_dafny.CodePoint('c'), globalState)).(bool)
						}
						return false
					})(), _927_v114), true, globalState)
					var _928_v115 _dafny.Sequence
					_ = _928_v115
					_928_v115 = _dafny.SeqOf(_dafny.Companion_Sequence_.IsPrefixOf((_this).F17(), (_this).F17()), (func() bool {
						if (_926_v113).Contains((_this).F4()) {
							return (_926_v113).Get((_this).F4()).(bool)
						}
						return (_819_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_819_v29), 0))).Int()).(bool)
					})(), (_819_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_819_v29), 0))).Int()).(bool), (_this).F4())
					var _929_v116 D11
					_ = _929_v116
					_929_v116 = Companion_D11_.Create_DC26_((_819_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_819_v29), 0))).Int()).(bool))
					var _930_v117 _dafny.MultiSet
					_ = _930_v117
					_930_v117 = _dafny.MultiSetOf((_this).F4())
					var _931_v118 D16
					_ = _931_v118
					_931_v118 = Companion_D16_.Create_DC33_((_929_v116).Dtor_cf41(), (_923_v111).F12(), _930_v117, p0, _927_v114)
					var _932_v119 D4
					_ = _932_v119
					_932_v119 = Companion_D4_.Create_DC13_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_931_v118), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqOf(_931_v118)).Cardinality()))).Uint32(), _931_v118)).Cardinality()), (_819_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_819_v29), 0))).Int()).(bool), p0, _928_v115)
					var _pat_let_tv13 = p0
					_ = _pat_let_tv13
					var _pat_let_tv14 = _928_v115
					_ = _pat_let_tv14
					_928_v115 = _dafny.Companion_Sequence_.Concatenate((func(_pat_let15_0 D4) D4 {
						return func(_933_dt__update__tmp_h1 D4) D4 {
							return func(_pat_let16_0 _dafny.Int) D4 {
								return func(_934_dt__update_hcf14_h0 _dafny.Int) D4 {
									return func(_pat_let17_0 _dafny.Sequence) D4 {
										return func(_935_dt__update_hcf17_h0 _dafny.Sequence) D4 {
											return Companion_D4_.Create_DC13_(_934_dt__update_hcf14_h0, (_933_dt__update__tmp_h1).Dtor_cf15(), (_933_dt__update__tmp_h1).Dtor_cf16(), _935_dt__update_hcf17_h0)
										}(_pat_let17_0)
									}(_pat_let_tv14)
								}(_pat_let16_0)
							}(_pat_let_tv13)
						}(_pat_let15_0)
					}(_932_v119)).Dtor_cf17(), _928_v115)
					var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_819_v29), 0))
					_ = _index132
					var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_819_v29), 0))
					_ = _index133
					var _rhs140 bool = ((func() _dafny.Set {
						if (_this).F4() {
							return _927_v114
						}
						return _927_v114
					})()).IsSubsetOf(_927_v114)
					_ = _rhs140
					var _rhs141 bool = !((_this).F4())
					_ = _rhs141
					var _lhs80 _dafny.Array = _819_v29
					_ = _lhs80
					var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_819_v29), 0))
					_ = _lhs81
					var _lhs82 _dafny.Array = _819_v29
					_ = _lhs82
					var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_819_v29), 0))
					_ = _lhs83
					(_lhs80).ArraySet1(_rhs140, (_lhs81).Int())
					(_lhs82).ArraySet1(_rhs141, (_lhs83).Int())
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _936_v120 _dafny.Int
		_ = _936_v120
		_936_v120 = _dafny.IntOfInt64(906)
		_936_v120 = Companion_Default___.SafeModuloInt(_936_v120, (_923_v111).F12())
		_936_v120 = (_dafny.IntOfUint32(((_this).F17()).Cardinality())).Plus(_936_v120)
		var _937_v121 _dafny.Map
		_ = _937_v121
		_937_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_this).F4() {
				return (_923_v111).Fm17((_this).F4(), globalState)
			}
			return (_this).F4()
		})(), ((_923_v111).F12()).Minus(_dafny.IntOfInt64(518)))
		r0 = _937_v121
		return r0
	}
}
func (_this *C2) M2(p0 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Array) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _938_v0 _dafny.Sequence
		_ = _938_v0
		_938_v0 = _dafny.SeqOf((_this).F4(), (_this).F4())
		var _939_v1 D3
		_ = _939_v1
		_939_v1 = Companion_D3_.Create_DC8_(_938_v0)
		var _940_v2 _dafny.Array
		_ = _940_v2
		var _nwElement0_29 D3 = _939_v1
		_ = _nwElement0_29
		var _nw141 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(26))
		_ = _nw141
		(_nw141).ArraySet1(_nwElement0_29, 0)
		(_nw141).ArraySet1(Companion_D3_.Create_DC8_(_938_v0), 1)
		(_nw141).ArraySet1(_939_v1, 2)
		(_nw141).ArraySet1(_939_v1, 3)
		(_nw141).ArraySet1(Companion_D3_.Create_DC8_(_938_v0), 4)
		(_nw141).ArraySet1(Companion_D3_.Create_DC8_(_938_v0), 5)
		(_nw141).ArraySet1(_939_v1, 6)
		(_nw141).ArraySet1(_939_v1, 7)
		(_nw141).ArraySet1(_939_v1, 8)
		(_nw141).ArraySet1(_939_v1, 9)
		(_nw141).ArraySet1(Companion_D3_.Create_DC8_(_938_v0), 10)
		(_nw141).ArraySet1(_939_v1, 11)
		(_nw141).ArraySet1(Companion_Default___.Fm52(_dafny.IntOfInt64(-402), globalState), 12)
		(_nw141).ArraySet1(_939_v1, 13)
		(_nw141).ArraySet1(_939_v1, 14)
		(_nw141).ArraySet1(_939_v1, 15)
		(_nw141).ArraySet1(_939_v1, 16)
		(_nw141).ArraySet1(_939_v1, 17)
		(_nw141).ArraySet1(_939_v1, 18)
		(_nw141).ArraySet1(_939_v1, 19)
		(_nw141).ArraySet1(_939_v1, 20)
		(_nw141).ArraySet1(_939_v1, 21)
		(_nw141).ArraySet1(Companion_D3_.Create_DC8_(_938_v0), 22)
		(_nw141).ArraySet1(Companion_D3_.Create_DC8_(_938_v0), 23)
		(_nw141).ArraySet1(_939_v1, 24)
		(_nw141).ArraySet1(_939_v1, 25)
		_940_v2 = _nw141
		var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_940_v2), 0))
		_ = _index134
		(_940_v2).ArraySet1(_939_v1, (_index134).Int())
		var _pat_let_tv15 = _938_v0
		_ = _pat_let_tv15
		var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_940_v2), 0))
		_ = _index135
		(_940_v2).ArraySet1(func(_pat_let18_0 D3) D3 {
			return func(_941_dt__update__tmp_h0 D3) D3 {
				return func(_pat_let19_0 _dafny.Sequence) D3 {
					return func(_942_dt__update_hcf7_h0 _dafny.Sequence) D3 {
						return Companion_D3_.Create_DC8_(_942_dt__update_hcf7_h0)
					}(_pat_let19_0)
				}(_pat_let_tv15)
			}(_pat_let18_0)
		}(_939_v1), (_index135).Int())
		var _943_v3 _dafny.Int
		_ = _943_v3
		_943_v3 = _dafny.IntOfInt64(187)
		var _944_v4 _dafny.Array
		_ = _944_v4
		var _nwElement0_30 _dafny.Int = Companion_Default___.SafeModuloInt(_943_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_943_v3, _943_v3), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.IntOfUint32((_dafny.SeqOf(_943_v3, _943_v3)).Cardinality()))).Uint32(), _943_v3)).Cardinality()))
		_ = _nwElement0_30
		var _nw142 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(9))
		_ = _nw142
		(_nw142).ArraySet1(_nwElement0_30, 0)
		(_nw142).ArraySet1(_943_v3, 1)
		(_nw142).ArraySet1(_943_v3, 2)
		(_nw142).ArraySet1(_943_v3, 3)
		(_nw142).ArraySet1(_943_v3, 4)
		(_nw142).ArraySet1((_943_v3).Plus(_dafny.IntOfInt64(577)), 5)
		(_nw142).ArraySet1(Companion_Default___.Fm3(_943_v3, (_this).F4(), globalState), 6)
		(_nw142).ArraySet1(_943_v3, 7)
		(_nw142).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32(((_this).F17()).Cardinality())), 8)
		_944_v4 = _nw142
		for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_944_v4), 0))); ; {
			_guard_loop_3, _ok41 := _iter41()
			if !_ok41 {
				break
			}
			var _945_i0 _dafny.Int
			_945_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_945_i0).Sign() != -1) && ((_945_i0).Cmp(_dafny.ArrayLen((_944_v4), 0)) < 0)) {
				(_944_v4).ArraySet1(Companion_Default___.SafeModuloInt(_945_i0, (_dafny.Zero).Minus((_943_v3).Minus(_943_v3))), (_945_i0).Int())
			}
		}
		var _946_v5 _dafny.Array
		_ = _946_v5
		var _nwElement0_31 bool = (_this).F4()
		_ = _nwElement0_31
		var _nw143 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.One)
		_ = _nw143
		(_nw143).ArraySet1(_nwElement0_31, 0)
		_946_v5 = _nw143
		var _947_v6 _dafny.Map
		_ = _947_v6
		_947_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _946_v5)
		var _948_v7 _dafny.Map
		_ = _948_v7
		_948_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((_947_v6).Cardinality()).Times(_943_v3)), (_dafny.Zero).Minus(_943_v3))
		_948_v7 = (_948_v7).Update(_943_v3, _943_v3)
		var _949_v8 _dafny.MultiSet
		_ = _949_v8
		_949_v8 = _dafny.MultiSetOf(!((_this).F4()), p0)
		var _950_v11 _dafny.CodePoint
		_ = _950_v11
		_950_v11 = _dafny.CodePoint('n')
		var _951_v12 _dafny.Array
		_ = _951_v12
		var _nwElement0_32 _dafny.Array = _946_v5
		_ = _nwElement0_32
		var _nw144 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(13))
		_ = _nw144
		(_nw144).ArraySet1(_nwElement0_32, 0)
		(_nw144).ArraySet1(_946_v5, 1)
		(_nw144).ArraySet1(_946_v5, 2)
		(_nw144).ArraySet1(_946_v5, 3)
		(_nw144).ArraySet1(_946_v5, 4)
		(_nw144).ArraySet1(_946_v5, 5)
		(_nw144).ArraySet1(_946_v5, 6)
		(_nw144).ArraySet1(_946_v5, 7)
		(_nw144).ArraySet1(_946_v5, 8)
		(_nw144).ArraySet1(_946_v5, 9)
		(_nw144).ArraySet1(_946_v5, 10)
		(_nw144).ArraySet1(_946_v5, 11)
		(_nw144).ArraySet1(_946_v5, 12)
		_951_v12 = _nw144
		var _952_v14 *C1
		_ = _952_v14
		var _nw145 *C1 = New_C1_()
		_ = _nw145
		_nw145.Ctor__(_dafny.IntOfUint32((Companion_Default___.Fm50((func() _dafny.Int {
			if (_949_v8).Contains((_this).F4()) {
				return (_949_v8).Multiplicity((_this).F4())
			}
			return _943_v3
		})(), (func() _dafny.Map {
			var _coll38 = _dafny.NewMapBuilder()
			_ = _coll38
			for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(321), _dafny.IntOfInt64(377))); ; {
				_compr_38, _ok42 := _iter42()
				if !_ok42 {
					break
				}
				var _953_v9 _dafny.Int
				_953_v9 = interface{}(_compr_38).(_dafny.Int)
				if ((_dafny.IntOfInt64(321)).Cmp(_953_v9) <= 0) && ((_953_v9).Cmp(_dafny.IntOfInt64(377)) < 0) {
					_coll38.Add(Companion_Default___.SafeDivisionInt(_953_v9, (func() _dafny.Map {
						var _coll39 = _dafny.NewMapBuilder()
						_ = _coll39
						for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(72), _dafny.IntOfInt64(693))); ; {
							_compr_39, _ok43 := _iter43()
							if !_ok43 {
								break
							}
							var _954_v10 _dafny.Int
							_954_v10 = interface{}(_compr_39).(_dafny.Int)
							if ((_dafny.IntOfInt64(72)).Cmp(_954_v10) <= 0) && ((_954_v10).Cmp(_dafny.IntOfInt64(693)) < 0) {
								_coll39.Add((_954_v10).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cii")).Cardinality())), (_this).F4())
							}
						}
						return _coll39.ToMap()
					}()).Cardinality()), _dafny.IntOfUint32(((_this).F17()).Cardinality()))
				}
			}
			return _coll38.ToMap()
		}()).Cardinality(), _943_v3, Companion_Default___.Fm0(_950_v11, globalState), globalState)).Cardinality()), _943_v3, _951_v12, Companion_Default___.Fm3((func() _dafny.Map {
			var _coll40 = _dafny.NewMapBuilder()
			_ = _coll40
			for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(412), _dafny.IntOfInt64(-676))); ; {
				_compr_40, _ok44 := _iter44()
				if !_ok44 {
					break
				}
				var _955_v13 _dafny.Int
				_955_v13 = interface{}(_compr_40).(_dafny.Int)
				if ((_dafny.IntOfInt64(412)).Cmp(_955_v13) <= 0) && ((_955_v13).Cmp(_dafny.IntOfInt64(-676)) < 0) {
					_coll40.Add((_955_v13).Times(_943_v3), p0)
				}
			}
			return _coll40.ToMap()
		}()).Cardinality(), p0, globalState), p0)
		_952_v14 = _nw145
		var _956_v15 _dafny.Set
		_ = _956_v15
		_956_v15 = _dafny.SetOf(p0, p0, p0)
		_956_v15 = Companion_Default___.Fm53(Companion_Default___.SafeModuloInt(_943_v3, _943_v3), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm54(_952_v14.F16, globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(644))).Uint32(), func(coer64 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg64 _dafny.Int) interface{} {
				return coer64(arg64)
			}
		}((func(_957_p0 bool) func(_dafny.Int) _dafny.Map {
			return func(_958_i1 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _957_p0)
			}
		})(p0)))), globalState)
		var _959_i2 _dafny.Int
		_ = _959_i2
		_959_i2 = _dafny.Zero
		{
			for !(((func() _dafny.Int {
				if p0 {
					return _943_v3
				}
				return (_952_v14).F15()
			})()).Cmp((_952_v14.F16).Times((_952_v14).F15())) < 0) {
				{
					if (_959_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_959_i2 = (_959_i2).Plus(_dafny.One)
					var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_946_v5), 0))
					_ = _index136
					(_946_v5).ArraySet1(((_952_v14).F15()).Cmp(_dafny.IntOfInt64(21)) < 0, (_index136).Int())
					var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_946_v5), 0))
					_ = _index137
					(_946_v5).ArraySet1(Companion_Default___.Fm0(_950_v11, globalState), (_index137).Int())
					(_952_v14).F16 = Companion_Default___.Fm3((_this).Fm7(globalState), p0, globalState)
					var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_946_v5), 0))
					_ = _index138
					(_946_v5).ArraySet1((_this).F4(), (_index138).Int())
					var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_946_v5), 0))
					_ = _index139
					(_946_v5).ArraySet1((_dafny.IntOfUint32(((_this).F17()).Cardinality())).Cmp(_943_v3) <= 0, (_index139).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		r0 = _950_v11
		var _960_v16 D1
		_ = _960_v16
		_960_v16 = Companion_D1_.Create_DC3_(p0)
		var _nwElement0_33 D1 = _960_v16
		_ = _nwElement0_33
		var _nw146 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(12))
		_ = _nw146
		(_nw146).ArraySet1(_nwElement0_33, 0)
		(_nw146).ArraySet1(_960_v16, 1)
		(_nw146).ArraySet1(_960_v16, 2)
		(_nw146).ArraySet1(_960_v16, 3)
		(_nw146).ArraySet1(_960_v16, 4)
		(_nw146).ArraySet1(_960_v16, 5)
		(_nw146).ArraySet1(_960_v16, 6)
		(_nw146).ArraySet1(_960_v16, 7)
		(_nw146).ArraySet1(_960_v16, 8)
		(_nw146).ArraySet1(func(_pat_let20_0 D1) D1 {
			return func(_961_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let21_0 bool) D1 {
					return func(_962_dt__update_hcf2_h0 bool) D1 {
						return Companion_D1_.Create_DC3_(_962_dt__update_hcf2_h0)
					}(_pat_let21_0)
				}((_this).F4())
			}(_pat_let20_0)
		}(_960_v16), 9)
		(_nw146).ArraySet1(Companion_Default___.Fm55(!(p0), (_this).F4(), _dafny.IntOfInt64(547), globalState), 10)
		(_nw146).ArraySet1(Companion_D1_.Create_DC3_((_this).F4()), 11)
		r1 = _nw146
		return r0, r1
	}
}
func (_this *C2) M12(p0 _dafny.Int, p1 _dafny.Set, p2 _dafny.Map, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _963_v0 _dafny.Array
		_ = _963_v0
		var _nw147 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
		_ = _nw147
		_963_v0 = _nw147
		var _hi4 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(_963_v0, _963_v0, _963_v0, _963_v0)).Cardinality())
		_ = _hi4
		for _964_i0 := ((p1).Intersection(_dafny.SetOf((_this).F4(), (_this).F4()))).Cardinality(); _964_i0.Cmp(_hi4) < 0; _964_i0 = _964_i0.Plus(_dafny.One) {
			var _965_v1 *C0
			_ = _965_v1
			var _nw148 *C0 = New_C0_()
			_ = _nw148
			_nw148.Ctor__(_dafny.IntOfInt64(271))
			_965_v1 = _nw148
			var _966_v2 _dafny.Map
			_ = _966_v2
			_966_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), p0)
			var _967_v3 _dafny.Sequence
			_ = _967_v3
			_967_v3 = _dafny.SeqOf(_964_i0)
			var _968_v4 D9
			_ = _968_v4
			_968_v4 = Companion_D9_.Create_DC19_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-645))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg65 _dafny.Int) interface{} {
					return coer65(arg65)
				}
			}(func(_969_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			})))
			var _970_v5 _dafny.Array
			_ = _970_v5
			var _nwElement0_34 _dafny.Int = p0
			_ = _nwElement0_34
			var _nw149 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(21))
			_ = _nw149
			(_nw149).ArraySet1(_nwElement0_34, 0)
			(_nw149).ArraySet1(_964_i0, 1)
			(_nw149).ArraySet1(_964_i0, 2)
			(_nw149).ArraySet1(_dafny.IntOfInt64(896), 3)
			(_nw149).ArraySet1((_965_v1).F12(), 4)
			(_nw149).ArraySet1((_dafny.MultiSetFromSeq(_967_v3)).Cardinality(), 5)
			(_nw149).ArraySet1((_965_v1).F12(), 6)
			(_nw149).ArraySet1(_964_i0, 7)
			(_nw149).ArraySet1(p0, 8)
			(_nw149).ArraySet1(_964_i0, 9)
			(_nw149).ArraySet1(_964_i0, 10)
			(_nw149).ArraySet1(p0, 11)
			(_nw149).ArraySet1(_dafny.IntOfUint32(((_this).F17()).Cardinality()), 12)
			(_nw149).ArraySet1(_dafny.IntOfUint32(((_this).F17()).Cardinality()), 13)
			(_nw149).ArraySet1((_965_v1).F12(), 14)
			(_nw149).ArraySet1(_dafny.IntOfUint32(((_968_v4).Dtor_cf25()).Cardinality()), 15)
			(_nw149).ArraySet1(_dafny.IntOfInt64(-648), 16)
			(_nw149).ArraySet1(_964_i0, 17)
			(_nw149).ArraySet1(p0, 18)
			(_nw149).ArraySet1(_964_i0, 19)
			(_nw149).ArraySet1(p0, 20)
			_970_v5 = _nw149
			var _971_v6 _dafny.Sequence
			_ = _971_v6
			_971_v6 = _dafny.SeqOf(_970_v5)
			var _972_v7 _dafny.Map
			_ = _972_v7
			_972_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_966_v2, _971_v6)
			var _973_v8 _dafny.Sequence
			_ = _973_v8
			_973_v8 = _dafny.SeqOf(_971_v6)
			_972_v7 = (_972_v7).Update(_966_v2, (_973_v8).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_973_v8).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _974_v9 _dafny.Array
			_ = _974_v9
			var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw150
			_974_v9 = _nw150
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_974_v9), 0))
			_ = _index140
			(_974_v9).ArraySet1(_963_v0, (_index140).Int())
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_963_v0), 0))
			_ = _index141
			(_963_v0).ArraySet1(false, (_index141).Int())
			var _975_v10 _dafny.CodePoint
			_ = _975_v10
			_975_v10 = _dafny.CodePoint('c')
			var _976_v11 T1
			_ = _976_v11
			var _nw151 *C1 = New_C1_()
			_ = _nw151
			_nw151.Ctor__(_964_i0, p0, _974_v9, (_this).Fm7(globalState), p3)
			_976_v11 = _nw151
			var _977_v12 _dafny.Map
			_ = _977_v12
			_977_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_976_v11, (_976_v11).F4())
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_963_v0), 0))
			_ = _index142
			(_963_v0).ArraySet1((func() bool {
				if (_977_v12).Contains(_976_v11) {
					return (_977_v12).Get(_976_v11).(bool)
				}
				return (_976_v11).F4()
			})(), (_index142).Int())
			var _978_v13 _dafny.Sequence
			_ = _978_v13
			_978_v13 = _dafny.SeqOf((_976_v11).F4(), (_976_v11).F4())
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_974_v9), 0))
			_ = _index143
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_963_v0), 0))
			_ = _index144
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_963_v0), 0))
			_ = _index145
			var _rhs142 _dafny.Array = _963_v0
			_ = _rhs142
			var _rhs143 _dafny.Int = (_this).Fm7(globalState)
			_ = _rhs143
			var _rhs144 bool = !((func() bool {
				if (_976_v11).F4() {
					return !(false)
				}
				return (_976_v11).F4()
			})())
			_ = _rhs144
			var _rhs145 _dafny.CodePoint = _dafny.CodePoint('k')
			_ = _rhs145
			var _rhs146 bool = (func() bool {
				if (_978_v13).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_978_v13).Cardinality()))).Uint32()).(bool) {
					return !(p3) || ((_this).F4())
				}
				return p3
			})()
			_ = _rhs146
			var _lhs84 _dafny.Array = _974_v9
			_ = _lhs84
			var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_974_v9), 0))
			_ = _lhs85
			var _lhs86 _dafny.Array = _963_v0
			_ = _lhs86
			var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_963_v0), 0))
			_ = _lhs87
			var _lhs88 _dafny.Array = _963_v0
			_ = _lhs88
			var _lhs89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_963_v0), 0))
			_ = _lhs89
			(_lhs84).ArraySet1(_rhs142, (_lhs85).Int())
			r0 = _rhs143
			(_lhs86).ArraySet1(_rhs144, (_lhs87).Int())
			_975_v10 = _rhs145
			(_lhs88).ArraySet1(_rhs146, (_lhs89).Int())
			var _979_v14 _dafny.Map
			_ = _979_v14
			_979_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_976_v11).F4())
			var _980_v15 _dafny.Sequence
			_ = _980_v15
			_980_v15 = _dafny.SeqOf(_979_v14, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_979_v14).Cardinality(), p3)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_965_v1).F12(), (_976_v11).F4())), _979_v14)
			_980_v15 = _980_v15
		}
		var _981_v16 _dafny.CodePoint
		_ = _981_v16
		_981_v16 = _dafny.CodePoint('l')
		var _982_i2 _dafny.Int
		_ = _982_i2
		_982_i2 = _dafny.Zero
		{
			for Companion_Default___.Fm0(_981_v16, globalState) {
				{
					if (_982_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_982_i2 = (_982_i2).Plus(_dafny.One)
					var _983_v17 D10
					_ = _983_v17
					_983_v17 = Companion_D10_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), p0))
					var _984_v18 _dafny.Sequence
					_ = _984_v18
					_984_v18 = _dafny.SeqOf(_983_v17, _983_v17)
					var _985_v19 _dafny.MultiSet
					_ = _985_v19
					_985_v19 = _dafny.MultiSetOf((_dafny.IntOfInt64(390)).Minus(p0), p0, ((_dafny.Zero).Minus(_dafny.IntOfUint32((_984_v18).Cardinality()))).Minus((p2).Cardinality()))
					r0 = (func() _dafny.Int {
						if (_985_v19).Contains((_dafny.Zero).Minus((p0).Times((_dafny.Zero).Minus(p0)))) {
							return (_985_v19).Multiplicity((_dafny.Zero).Minus((p0).Times((_dafny.Zero).Minus(p0))))
						}
						return p0
					})()
					var _986_v20 _dafny.Sequence
					_ = _986_v20
					_986_v20 = _dafny.UnicodeSeqOfUtf8Bytes("ulsxdir")
					_986_v20 = _dafny.Companion_Sequence_.Concatenate(_986_v20, (func() _dafny.Sequence {
						if !((_this).F4()) {
							return _986_v20
						}
						return (_this).F17()
					})())
					var _987_v21 _dafny.MultiSet
					_ = _987_v21
					_987_v21 = _dafny.MultiSetOf(p3)
					var _988_v22 _dafny.Array
					_ = _988_v22
					var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
					_ = _nw152
					_988_v22 = _nw152
					var _989_v23 *C1
					_ = _989_v23
					var _nw153 *C1 = New_C1_()
					_ = _nw153
					_nw153.Ctor__(p0, (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.SeqOf(_981_v16, _981_v16, _981_v16)).Cardinality())).Plus((func() _dafny.Int {
						if (_987_v21).Contains(true) {
							return (_987_v21).Multiplicity(true)
						}
						return p0
					})())), (func() _dafny.Array {
						if p3 {
							return _988_v22
						}
						return _988_v22
					})(), p0, (p0).Cmp(_dafny.IntOfInt64(493)) != 0)
					_989_v23 = _nw153
					r0 = ((_989_v23).Fm9((_this).F4(), (_this).F4(), (_989_v23).F15(), globalState)).Times((_989_v23).F15())
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		r0 = (_dafny.Zero).Minus(((p0).Minus(p0)).Minus(Companion_Default___.SafeModuloInt(p0, p0)))
		var _990_v24 _dafny.Array
		_ = _990_v24
		var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw154
		_990_v24 = _nw154
		var _991_v25 _dafny.Map
		_ = _991_v25
		_991_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus((_this).Fm7(globalState))).Plus(p0), _990_v24)
		_990_v24 = (func() _dafny.Array {
			if (_991_v25).Contains(_dafny.IntOfInt64(718)) {
				return (_991_v25).Get(_dafny.IntOfInt64(718)).(_dafny.Array)
			}
			return _990_v24
		})()
		var _992_v26 D2
		_ = _992_v26
		_992_v26 = Companion_D2_.Create_DC7_((func() _dafny.Int {
			if (_this).F4() {
				return p0
			}
			return p0
		})())
		var _source17 D2 = _992_v26
		_ = _source17
		if _source17.Is_DC6() {
			var _993_v27 *C0
			_ = _993_v27
			var _nw155 *C0 = New_C0_()
			_ = _nw155
			_nw155.Ctor__(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfInt64(-298)), p0))
			_993_v27 = _nw155
			var _994_v28 bool
			_ = _994_v28
			_994_v28 = true
			_994_v28 = _994_v28
			r0 = p0
			r1 = (_993_v27).F12()
		} else if _source17.Is_DC7() {
			var _995___mcc_h0 _dafny.Int = _source17.Get_().(D2_DC7).Cf6
			_ = _995___mcc_h0
			var _996_cf6 _dafny.Int = _995___mcc_h0
			_ = _996_cf6
			r1 = p0
			var _997_v29 bool
			_ = _997_v29
			_997_v29 = false
			_997_v29 = (p0).Cmp(_dafny.IntOfUint32(((_this).F17()).Cardinality())) != 0
			var _998_v30 _dafny.MultiSet
			_ = _998_v30
			_998_v30 = _dafny.MultiSetOf((_dafny.Zero).Minus(p0))
			var _999_v31 _dafny.Map
			_ = _999_v31
			_999_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_998_v30, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-304))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg66 _dafny.Int) interface{} {
					return coer66(arg66)
				}
			}((func(_1000_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1001_i3 _dafny.Int) _dafny.CodePoint {
					return _1000_v16
				}
			})(_981_v16))))
			var _1002_v32 _dafny.Sequence
			_ = _1002_v32
			_1002_v32 = _dafny.SeqOf(p3)
			var _1003_v33 _dafny.Sequence
			_ = _1003_v33
			_1003_v33 = _dafny.SeqOf(Companion_Default___.Fm56((_this).F4(), _996_cf6, p3, globalState), _1002_v32, _1002_v32)
			var _1004_v34 _dafny.MultiSet
			_ = _1004_v34
			_1004_v34 = _dafny.MultiSetOf(_1002_v32)
			var _1005_v35 bool
			_ = _1005_v35
			var _1006_v36 _dafny.Int
			_ = _1006_v36
			var _out27 bool
			_ = _out27
			var _out28 _dafny.Int
			_ = _out28
			_out27, _out28 = Companion_Default___.M0(_999_v31, (_dafny.MultiSetFromSeq(_1003_v33)).IsSubsetOf(_1004_v34), _998_v30, ((p2).Cardinality()).Cmp(p0) == 0, globalState)
			_1005_v35 = _out27
			_1006_v36 = _out28
			var _1007_v37 _dafny.Array
			_ = _1007_v37
			var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
			_ = _nw156
			_1007_v37 = _nw156
			var _1008_v38 *C1
			_ = _1008_v38
			var _nw157 *C1 = New_C1_()
			_ = _nw157
			_nw157.Ctor__(p0, _dafny.IntOfInt64(143), _1007_v37, _1006_v36, !(false))
			_1008_v38 = _nw157
			var _1009_v39 _dafny.MultiSet
			_ = _1009_v39
			_1009_v39 = _dafny.MultiSetOf((_this).F4(), p3)
			var _1010_v40 D6
			_ = _1010_v40
			_1010_v40 = Companion_D6_.Create_DC16_(false, p3, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_996_cf6), (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1005_v35, _1008_v38)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_996_cf6)).Cardinality()))).Uint32(), (_1009_v39).Cardinality()))
			var _1011_v41 _dafny.Map
			_ = _1011_v41
			_1011_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(_996_cf6, (_this).F4(), globalState), _1010_v40)
			var _1012_v42 _dafny.Set
			_ = _1012_v42
			_1012_v42 = _dafny.SetOf(_963_v0)
			_1011_v41 = (_1011_v41).Update(((_1012_v42).Union(_1012_v42)).Cardinality(), _1010_v40)
		} else {
			var _1013___mcc_h1 _dafny.Array = _source17.Get_().(D2_DC5).Cf5
			_ = _1013___mcc_h1
			var _1014_cf5 _dafny.Array = _1013___mcc_h1
			_ = _1014_cf5
			var _1015_v43 _dafny.MultiSet
			_ = _1015_v43
			_1015_v43 = _dafny.MultiSetOf(true)
			var _1016_v44 _dafny.Sequence
			_ = _1016_v44
			_1016_v44 = _dafny.SeqOf(p0, (func() _dafny.Int {
				if (_1015_v43).Contains(p3) {
					return (_1015_v43).Multiplicity(p3)
				}
				return (_dafny.Zero).Minus(p0)
			})(), p0)
			var _rhs147 _dafny.Array = _963_v0
			_ = _rhs147
			var _rhs148 _dafny.Int = p0
			_ = _rhs148
			var _rhs149 _dafny.Int = (_1016_v44).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(p0, p0), _dafny.IntOfUint32((_1016_v44).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs149
			_1014_cf5 = _rhs147
			r0 = _rhs148
			r1 = _rhs149
			var _1017_v45 bool
			_ = _1017_v45
			_1017_v45 = true
			_1017_v45 = !((_this).F4())
			var _1018_v46 _dafny.Map
			_ = _1018_v46
			_1018_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_990_v24, !((_this).F4()))
			var _1019_v47 _dafny.Map
			_ = _1019_v47
			_1019_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1017_v45, _1018_v46)
			_1019_v47 = (_1019_v47).Update(p3, (func() _dafny.Map {
				if _1017_v45 {
					return _1018_v46
				}
				return _1018_v46
			})())
			var _1020_v48 _dafny.Map
			_ = _1020_v48
			_1020_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_this).F17(), _dafny.UnicodeSeqOfUtf8Bytes("x")), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F17(), _dafny.UnicodeSeqOfUtf8Bytes("x"))).Cardinality()))).Uint32(), _981_v16), (_1015_v43).IsSubsetOf(_dafny.MultiSetOf(p3)))
			_1020_v48 = (_1020_v48).Update((func() _dafny.Sequence {
				if false {
					return _dafny.UnicodeSeqOfUtf8Bytes("jpo")
				}
				return _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("qlx"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qlx")).Cardinality()))).Uint32(), _981_v16)
			})(), (_this).F4())
		}
		if ((_this).F4()) && (p3) {
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_963_v0), 0))
			_ = _index146
			(_963_v0).ArraySet1(p3, (_index146).Int())
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_963_v0), 0))
			_ = _index147
			(_963_v0).ArraySet1((_this).F4(), (_index147).Int())
			var _1021_v49 _dafny.Sequence
			_ = _1021_v49
			_1021_v49 = _dafny.SeqOf(p0, p0)
			var _1022_v50 _dafny.Sequence
			_ = _1022_v50
			_1022_v50 = _dafny.SeqOf((func() _dafny.Int {
				if (p2).Contains((_1021_v49).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_963_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_963_v0), 0))).Int()).(bool), p3)).Cardinality())), _dafny.IntOfUint32((_1021_v49).Cardinality()))).Uint32()).(_dafny.Int)) {
					return (p2).Get((_1021_v49).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_963_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_963_v0), 0))).Int()).(bool), p3)).Cardinality())), _dafny.IntOfUint32((_1021_v49).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Int)
				}
				return p0
			})(), p0)
			_1022_v50 = _dafny.Companion_Sequence_.Update(_1021_v49, (Companion_Default___.SafeIndex((_this).Fm7(globalState), _dafny.IntOfUint32((_1021_v49).Cardinality()))).Uint32(), p0)
			var _1023_v51 _dafny.Array
			_ = _1023_v51
			var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
			_ = _nw158
			_1023_v51 = _nw158
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_1023_v51), 0))
			_ = _index148
			(_1023_v51).ArraySet1((_this).F17(), (_index148).Int())
			var _1024_v52 _dafny.Sequence
			_ = _1024_v52
			_1024_v52 = _dafny.SeqOf(!((_963_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_963_v0), 0))).Int()).(bool)))
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_1023_v51), 0))
			_ = _index149
			(_1023_v51).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_1024_v52).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1024_v52).Cardinality()))).Uint32()).(bool) {
					return (_this).F17()
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("dehlggigm")
			})(), (_this).F17()), (_index149).Int())
			var _1025_v53 _dafny.Array
			_ = _1025_v53
			var _nwElement0_35 _dafny.CodePoint = _981_v16
			_ = _nwElement0_35
			var _nw159 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(15))
			_ = _nw159
			(_nw159).ArraySet1CodePoint(_nwElement0_35, 0)
			(_nw159).ArraySet1CodePoint(_981_v16, 1)
			(_nw159).ArraySet1CodePoint(_981_v16, 2)
			(_nw159).ArraySet1CodePoint(_981_v16, 3)
			(_nw159).ArraySet1CodePoint(_dafny.CodePoint('e'), 4)
			(_nw159).ArraySet1CodePoint(_dafny.CodePoint('s'), 5)
			(_nw159).ArraySet1CodePoint(_981_v16, 6)
			(_nw159).ArraySet1CodePoint(_981_v16, 7)
			(_nw159).ArraySet1CodePoint(((_1023_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_1023_v51), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_1023_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_1023_v51), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint), 8)
			(_nw159).ArraySet1CodePoint(_dafny.CodePoint('g'), 9)
			(_nw159).ArraySet1CodePoint(Companion_Default___.Fm1(p0, p3, globalState), 10)
			(_nw159).ArraySet1CodePoint(_981_v16, 11)
			(_nw159).ArraySet1CodePoint(_981_v16, 12)
			(_nw159).ArraySet1CodePoint(Companion_Default___.Fm1(p0, (_963_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_963_v0), 0))).Int()).(bool), globalState), 13)
			(_nw159).ArraySet1CodePoint(_981_v16, 14)
			_1025_v53 = _nw159
			_1025_v53 = _1025_v53
			var _1026_v54 _dafny.Array
			_ = _1026_v54
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_19
			var _nw160 _dafny.Array
			_ = _nw160
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw160 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) bool = (func(_1027_p3 bool) func(_dafny.Int) bool {
					return func(_1028_i4 _dafny.Int) bool {
						return _1027_p3
					}
				})(p3)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw160 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw160).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw160).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_1026_v54 = _nw160
			var _1029_v55 _dafny.Map
			_ = _1029_v55
			_1029_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfInt64(135)), _1026_v54)
			_1029_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1026_v54)
		} else {
			var _1030_v56 bool
			_ = _1030_v56
			_1030_v56 = false
			var _1031_v57 _dafny.Sequence
			_ = _1031_v57
			_1031_v57 = _dafny.SeqOf((_this).F4())
			var _1032_v58 _dafny.Map
			_ = _1032_v58
			_1032_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
			var _1033_v59 _dafny.Sequence
			_ = _1033_v59
			_1033_v59 = _dafny.SeqOf((_1032_v58).Cardinality())
			var _1034_v60 _dafny.Map
			_ = _1034_v60
			_1034_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1030_v56, _dafny.IntOfUint32((_1033_v59).Cardinality()))
			var _1035_v61 _dafny.MultiSet
			_ = _1035_v61
			_1035_v61 = _dafny.MultiSetOf((_1031_v57).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1034_v60).Contains(p3) {
					return (_1034_v60).Get(p3).(_dafny.Int)
				}
				return p0
			})(), _dafny.IntOfUint32((_1031_v57).Cardinality()))).Uint32()).(bool), p3)
			var _1036_v62 _dafny.Array
			_ = _1036_v62
			var _nwElement0_36 _dafny.MultiSet = _1035_v61
			_ = _nwElement0_36
			var _nw161 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(2))
			_ = _nw161
			(_nw161).ArraySet1(_nwElement0_36, 0)
			(_nw161).ArraySet1(_1035_v61, 1)
			_1036_v62 = _nw161
			var _1037_v63 _dafny.Array
			_ = _1037_v63
			var _nwElement0_37 _dafny.Array = _1036_v62
			_ = _nwElement0_37
			var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(7))
			_ = _nw162
			(_nw162).ArraySet1(_nwElement0_37, 0)
			(_nw162).ArraySet1(_1036_v62, 1)
			(_nw162).ArraySet1(_1036_v62, 2)
			(_nw162).ArraySet1(_1036_v62, 3)
			(_nw162).ArraySet1(_1036_v62, 4)
			(_nw162).ArraySet1((func() _dafny.Array {
				if (_this).F4() {
					return _1036_v62
				}
				return _1036_v62
			})(), 5)
			(_nw162).ArraySet1(_1036_v62, 6)
			_1037_v63 = _nw162
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1037_v63), 0))
			_ = _index150
			(_1037_v63).ArraySet1(_1036_v62, (_index150).Int())
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1037_v63), 0))
			_ = _index151
			var _rhs150 _dafny.Int = (_dafny.Zero).Minus(p0)
			_ = _rhs150
			var _rhs151 _dafny.Array = _963_v0
			_ = _rhs151
			var _rhs152 bool = !((_this).F4())
			_ = _rhs152
			var _rhs153 _dafny.Array = _1036_v62
			_ = _rhs153
			var _lhs90 _dafny.Array = _1037_v63
			_ = _lhs90
			var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1037_v63), 0))
			_ = _lhs91
			r1 = _rhs150
			_963_v0 = _rhs151
			_1030_v56 = _rhs152
			(_lhs90).ArraySet1(_rhs153, (_lhs91).Int())
			var _1038_v64 _dafny.Array
			_ = _1038_v64
			var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
			_ = _nw163
			_1038_v64 = _nw163
			var _1039_v65 *C1
			_ = _1039_v65
			var _nw164 *C1 = New_C1_()
			_ = _nw164
			_nw164.Ctor__((_1033_v59).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1033_v59).Cardinality()))).Uint32()).(_dafny.Int), Companion_Default___.SafeModuloInt(p0, p0), _1038_v64, Companion_Default___.SafeDivisionInt((_dafny.MultiSetOf(_dafny.IntOfUint32((_1031_v57).Cardinality()), _dafny.IntOfInt64(9))).Cardinality(), p0), p3)
			_1039_v65 = _nw164
			_1030_v56 = _1030_v56
			_1030_v56 = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_1033_v59, (Companion_Default___.SafeIndex((_1039_v65).F15(), _dafny.IntOfUint32((_1033_v59).Cardinality()))).Uint32(), p0), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1033_v59, Companion_Default___.Fm57((_1039_v65).F15(), (_dafny.Zero).Minus(_1039_v65.F16), globalState)), (Companion_Default___.SafeIndex(_1039_v65.F16, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1033_v59, Companion_Default___.Fm57((_1039_v65).F15(), (_dafny.Zero).Minus(_1039_v65.F16), globalState))).Cardinality()))).Uint32(), _1039_v65.F16))
			_1030_v56 = (_1039_v65.F16).Cmp(_1039_v65.F16) <= 0
		}
		r0 = _dafny.IntOfInt64(-259)
		r1 = (_dafny.Zero).Minus(p0)
		return r0, r1
	}
}
func (_this *C2) F17() _dafny.Sequence {
	{
		return _this._f17
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f5 _dafny.Array
	_f6 _dafny.Int
	_f4 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.Zero
	_this._f4 = false
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

func (_this *C3) F5() _dafny.Array {
	return _this._f5
}
func (_this *C3) F5_set_(value _dafny.Array) {
	_this._f5 = value
}
func (_this *C3) F6() _dafny.Int {
	return _this._f6
}
func (_this *C3) F4() bool {
	return _this._f4
}
func (_this *C3) Ctor__(f5 _dafny.Array, f6 _dafny.Int, f4 bool) {
	{
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f4 = f4
	}
}
func (_this *C3) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return !(((_this).F6()).Cmp((_dafny.Zero).Minus((_this).F6())) != 0) || ((_this).F4())
	}
}
func (_this *C3) Fm9(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(476)
	}
}
func (_this *C3) Fm6(p0 D0, p1 D0, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetOf((_this).F6(), (_this).F6())).Difference(_dafny.MultiSetOf((_this).F6(), (_this).F6()))
	}
}
func (_this *C3) Fm7(globalState *GlobalState) _dafny.Int {
	{
		return (_this).F6()
	}
}
func (_this *C3) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _1040_v0 _dafny.Array
		_ = _1040_v0
		var _len0_20 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_20
		var _nw165 _dafny.Array
		_ = _nw165
		if _len0_20.Cmp(_dafny.Zero) == 0 {
			_nw165 = _dafny.NewArray(_len0_20)
		} else {
			var _init20 func(_dafny.Int) _dafny.Int = (func(_1041_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1042_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1042_i1, _1041_p0)
				}
			})(p0)
			_ = _init20
			var _element0_20 = _init20(_dafny.Zero)
			_ = _element0_20
			_nw165 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
			(_nw165).ArraySet1(_element0_20, 0)
			var _nativeLen0_20 = (_len0_20).Int()
			_ = _nativeLen0_20
			for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
				(_nw165).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
			}
		}
		_1040_v0 = _nw165
		for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1040_v0), 0))); ; {
			_guard_loop_4, _ok45 := _iter45()
			if !_ok45 {
				break
			}
			var _1043_i0 _dafny.Int
			_1043_i0 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1043_i0).Sign() != -1) && ((_1043_i0).Cmp(_dafny.ArrayLen((_1040_v0), 0)) < 0)) {
				(_1040_v0).ArraySet1((_1043_i0).Minus(p0), (_1043_i0).Int())
			}
		}
		var _1044_v1 _dafny.Sequence
		_ = _1044_v1
		_1044_v1 = _dafny.SeqOf(false)
		var _1045_v2 _dafny.MultiSet
		_ = _1045_v2
		_1045_v2 = _dafny.MultiSetOf(_1044_v1)
		var _1046_v3 D6
		_ = _1046_v3
		_1046_v3 = Companion_D6_.Create_DC15_(_1045_v2)
		_1046_v3 = Companion_Default___.Fm33((_this).F4(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dqdyv")).Cardinality())), (func() bool {
			if !((_this).F4()) {
				return false
			}
			return !((_this).F4())
		})(), globalState)
		var _1047_v4 _dafny.Sequence
		_ = _1047_v4
		_1047_v4 = _dafny.SeqOf(p0)
		var _1048_v5 _dafny.Set
		_ = _1048_v5
		_1048_v5 = _dafny.SetOf(p0, _dafny.IntOfUint32((_1047_v4).Cardinality()))
		var _hi5 _dafny.Int = (p0).Minus((_1048_v5).Cardinality())
		_ = _hi5
		for _1049_i2 := p0; _1049_i2.Cmp(_hi5) < 0; _1049_i2 = _1049_i2.Plus(_dafny.One) {
			var _1050_v6 _dafny.Sequence
			_ = _1050_v6
			_1050_v6 = _dafny.UnicodeSeqOfUtf8Bytes("sj")
			var _1051_v7 _dafny.Map
			_ = _1051_v7
			_1051_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(871))).Uint32(), func(coer67 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg67 _dafny.Int) interface{} {
					return coer67(arg67)
				}
			}((func(_1052_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1053_i3 _dafny.Int) _dafny.Int {
					return _1052_p0
				}
			})(p0)))), _1050_v6)
			var _1054_v8 _dafny.MultiSet
			_ = _1054_v8
			_1054_v8 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1050_v6).Cardinality()))
			var _1055_v9 bool
			_ = _1055_v9
			var _1056_v10 _dafny.Int
			_ = _1056_v10
			var _out29 bool
			_ = _out29
			var _out30 _dafny.Int
			_ = _out30
			_out29, _out30 = Companion_Default___.M0((Companion_Default___.Fm34(_1049_i2, p0, !((_this).F4()), (_this).F4(), globalState)).Merge(_1051_v7), (_this).F4(), _1054_v8, (_this).F4(), globalState)
			_1055_v9 = _out29
			_1056_v10 = _out30
			_1056_v10 = ((_this).F6()).Times((func() _dafny.Int {
				if !((_this).F4()) {
					return p0
				}
				return (func() _dafny.Int {
					if (_1054_v8).Contains(_1049_i2) {
						return (_1054_v8).Multiplicity(_1049_i2)
					}
					return p0
				})()
			})())
			var _1057_v11 *C0
			_ = _1057_v11
			var _nw166 *C0 = New_C0_()
			_ = _nw166
			_nw166.Ctor__(_dafny.IntOfInt64(887))
			_1057_v11 = _nw166
			var _1058_v12 _dafny.Set
			_ = _1058_v12
			_1058_v12 = Companion_Default___.Fm35(_1049_i2, globalState)
			var _source18 _dafny.Set = _1058_v12
			_ = _source18
			var _1059___mcc_h0 _dafny.Set = _source18
			_ = _1059___mcc_h0
			var _1060_cf23 _dafny.Set = _1059___mcc_h0
			_ = _1060_cf23
			var _1061_v13 _dafny.Map
			_ = _1061_v13
			_1061_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1057_v11).F12(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(832))).Uint32(), func(coer68 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg68 _dafny.Int) interface{} {
					return coer68(arg68)
				}
			}((func(_1062_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1063_i5 _dafny.Int) _dafny.Int {
					return _1062_p0
				}
			})(p0))))
			var _1064_v14 _dafny.Array
			_ = _1064_v14
			var _nwElement0_38 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(870))).Uint32(), func(coer69 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg69 _dafny.Int) interface{} {
					return coer69(arg69)
				}
			}((func(_1065_v11 *C0) func(_dafny.Int) _dafny.Int {
				return func(_1066_i4 _dafny.Int) _dafny.Int {
					return (_1065_v11).F12()
				}
			})(_1057_v11)))
			_ = _nwElement0_38
			var _nw167 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(28))
			_ = _nw167
			(_nw167).ArraySet1(_nwElement0_38, 0)
			(_nw167).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1047_v4, _1047_v4), 1)
			(_nw167).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1047_v4, _1047_v4), 2)
			(_nw167).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1047_v4, _dafny.Companion_Sequence_.Update(_1047_v4, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.IntOfUint32((_1047_v4).Cardinality()))).Uint32(), _1056_v10)), 3)
			(_nw167).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_1061_v13).Contains((_1057_v11).F12()) {
					return (_1061_v13).Get((_1057_v11).F12()).(_dafny.Sequence)
				}
				return _1047_v4
			})(), _1047_v4), 4)
			(_nw167).ArraySet1(_1047_v4, 5)
			(_nw167).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1047_v4, _dafny.SeqOf(_1049_i2, (_this).Fm7(globalState))), 6)
			(_nw167).ArraySet1(_1047_v4, 7)
			(_nw167).ArraySet1(_1047_v4, 8)
			(_nw167).ArraySet1((func() _dafny.Sequence {
				if (_this).F4() {
					return _1047_v4
				}
				return _dafny.SeqOf(_1049_i2)
			})(), 9)
			(_nw167).ArraySet1(_1047_v4, 10)
			(_nw167).ArraySet1(_1047_v4, 11)
			(_nw167).ArraySet1(_1047_v4, 12)
			(_nw167).ArraySet1(_1047_v4, 13)
			(_nw167).ArraySet1(_1047_v4, 14)
			(_nw167).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(386))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg70 _dafny.Int) interface{} {
					return coer70(arg70)
				}
			}((func(_1067_v10 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1068_i6 _dafny.Int) _dafny.Int {
					return _1067_v10
				}
			})(_1056_v10))), 15)
			(_nw167).ArraySet1(_dafny.SeqOf((_dafny.Zero).Minus((_1057_v11).F12())), 16)
			(_nw167).ArraySet1(_1047_v4, 17)
			(_nw167).ArraySet1(Companion_Default___.Fm36(true, globalState), 18)
			(_nw167).ArraySet1(_dafny.SeqOf((_1054_v8).Cardinality(), (_dafny.Zero).Minus(_1056_v10)), 19)
			(_nw167).ArraySet1(_1047_v4, 20)
			(_nw167).ArraySet1(_dafny.SeqOf(_1049_i2, (_1057_v11).F12(), _dafny.IntOfUint32((_1044_v1).Cardinality()), _1056_v10), 21)
			(_nw167).ArraySet1(_1047_v4, 22)
			(_nw167).ArraySet1(_dafny.Companion_Sequence_.Update(_1047_v4, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1049_i2), _dafny.IntOfUint32((_1047_v4).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1050_v6).Cardinality())), 23)
			(_nw167).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(633))).Uint32(), func(coer71 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg71 _dafny.Int) interface{} {
					return coer71(arg71)
				}
			}(func(_1069_i7 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf((_this).F4(), false)).Cardinality())
			})), 24)
			(_nw167).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(528))).Uint32(), func(coer72 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg72 _dafny.Int) interface{} {
					return coer72(arg72)
				}
			}((func(_1070_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1071_i8 _dafny.Int) _dafny.Int {
					return _1070_p0
				}
			})(p0))), _dafny.SeqOf(_dafny.IntOfInt64(-466))), 25)
			(_nw167).ArraySet1(Companion_Default___.Fm36((_this).Fm8((_dafny.MultiSetOf(_1055_v9, (_this).F4(), _1055_v9)).Cardinality(), false, _1050_v6, _1047_v4, globalState), globalState), 26)
			(_nw167).ArraySet1(_1047_v4, 27)
			_1064_v14 = _nw167
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_1064_v14), 0))
			_ = _index152
			(_1064_v14).ArraySet1(_1047_v4, (_index152).Int())
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_1064_v14), 0))
			_ = _index153
			(_1064_v14).ArraySet1(_1047_v4, (_index153).Int())
			_1056_v10 = (_1057_v11).F12()
			var _1072_v15 _dafny.Sequence
			_ = _1072_v15
			_1072_v15 = _dafny.SeqOf(_1050_v6)
			var _1073_v16 _dafny.Sequence
			_ = _1073_v16
			_1073_v16 = _dafny.SeqOf(_1050_v6, (_1072_v15).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1072_v15).Cardinality()))).Uint32()).(_dafny.Sequence), _1050_v6)
			var _1074_v17 D3
			_ = _1074_v17
			_1074_v17 = Companion_D3_.Create_DC10_(_dafny.IntOfUint32(((_1073_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_1056_v10)).Cardinality()), _dafny.IntOfUint32((_1073_v16).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
			_1074_v17 = _1074_v17
			_1055_v9 = _1055_v9
		}
		var _1075_v18 _dafny.Array
		_ = _1075_v18
		var _nw168 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
		_ = _nw168
		_1075_v18 = _nw168
		for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1075_v18), 0))); ; {
			_guard_loop_5, _ok46 := _iter46()
			if !_ok46 {
				break
			}
			var _1076_i9 _dafny.Int
			_1076_i9 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_1076_i9).Sign() != -1) && ((_1076_i9).Cmp(_dafny.ArrayLen((_1075_v18), 0)) < 0)) {
				(_1075_v18).ArraySet1(_1044_v1, (_1076_i9).Int())
			}
		}
		var _1077_v19 bool
		_ = _1077_v19
		_1077_v19 = true
		_1077_v19 = (_this).F4()
		var _source19 D0 = Companion_D0_.Create_DC2_(_dafny.IntOfInt64(-988))
		_ = _source19
		if _source19.Is_DC1() {
			var _1078_v20 _dafny.Map
			_ = _1078_v20
			_1078_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(p0, _1077_v19, globalState), p0)
			var _1079_v21 _dafny.Sequence
			_ = _1079_v21
			_1079_v21 = _dafny.UnicodeSeqOfUtf8Bytes("vtke")
			_1078_v20 = (_1078_v20).Update(_dafny.IntOfUint32((_1079_v21).Cardinality()), Companion_Default___.SafeDivisionInt((_1048_v5).Cardinality(), (_dafny.Zero).Minus(p0)))
			_1077_v19 = (_this).F4()
			var _1080_v22 _dafny.Int
			_ = _1080_v22
			_1080_v22 = _dafny.IntOfInt64(-943)
			var _1081_v23 _dafny.Set
			_ = _1081_v23
			_1081_v23 = _dafny.SetOf(_1040_v0, _1040_v0)
			_1080_v22 = (_1081_v23).Cardinality()
			_1077_v19 = (_this).F4()
		} else if _source19.Is_DC2() {
			var _1082___mcc_h1 _dafny.Int = _source19.Get_().(D0_DC2).Cf1
			_ = _1082___mcc_h1
			var _1083_cf1 _dafny.Int = _1082___mcc_h1
			_ = _1083_cf1
			var _1084_v24 _dafny.Map
			_ = _1084_v24
			_1084_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1048_v5, p0)
			_1083_cf1 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_1084_v24).Contains(_1048_v5) {
					return (_1084_v24).Get(_1048_v5).(_dafny.Int)
				}
				return _1083_cf1
			})(), (_this).F6())
			var _1085_v25 _dafny.Map
			_ = _1085_v25
			_1085_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F6())
			var _1086_v26 D10
			_ = _1086_v26
			_1086_v26 = Companion_D10_.Create_DC21_((func() _dafny.Map {
				if !((_this).F4()) {
					return _1085_v25
				}
				return _1085_v25
			})())
			var _source20 D10 = _1086_v26
			_ = _source20
			if _source20.Is_DC22() {
				var _1087___mcc_h3 _dafny.Sequence = _source20.Get_().(D10_DC22).Cf30
				_ = _1087___mcc_h3
				var _1088___mcc_h4 bool = _source20.Get_().(D10_DC22).Cf31
				_ = _1088___mcc_h4
				var _1089___mcc_h5 bool = _source20.Get_().(D10_DC22).Cf32
				_ = _1089___mcc_h5
				var _1090___mcc_h6 _dafny.Array = _source20.Get_().(D10_DC22).Cf33
				_ = _1090___mcc_h6
				var _1091___mcc_h7 _dafny.Int = _source20.Get_().(D10_DC22).Cf34
				_ = _1091___mcc_h7
				var _1092_cf34 _dafny.Int = _1091___mcc_h7
				_ = _1092_cf34
				var _1093_cf33 _dafny.Array = _1090___mcc_h6
				_ = _1093_cf33
				var _1094_cf32 bool = _1089___mcc_h5
				_ = _1094_cf32
				var _1095_cf31 bool = _1088___mcc_h4
				_ = _1095_cf31
				var _1096_cf30 _dafny.Sequence = _1087___mcc_h3
				_ = _1096_cf30
				var _1097_v27 _dafny.CodePoint
				_ = _1097_v27
				_1097_v27 = _dafny.CodePoint('t')
				var _1098_v28 bool
				_ = _1098_v28
				var _out31 bool
				_ = _out31
				_out31 = (_this).M9(Companion_Default___.Fm0(_1097_v27, globalState), globalState)
				_1098_v28 = _out31
				_1083_cf1 = ((_1083_cf1).Times(_1083_cf1)).Plus(_dafny.IntOfInt64(-560))
				var _1099_v29 _dafny.Array
				_ = _1099_v29
				var _nw169 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(23))
				_ = _nw169
				_1099_v29 = _nw169
				var _1100_v30 _dafny.Map
				_ = _1100_v30
				_1100_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1099_v29, _1077_v19)
				var _1101_v31 D11
				_ = _1101_v31
				_1101_v31 = Companion_D11_.Create_DC23_(_1099_v29)
				_1100_v30 = (_1100_v30).Update((_1101_v31).Dtor_cf35(), (_this).F4())
				var _1102_v32 _dafny.MultiSet
				_ = _1102_v32
				_1102_v32 = _dafny.MultiSetOf(p0)
				_1077_v19 = (func() bool {
					if _1077_v19 {
						return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf((_this).F6()), _1047_v4)
					}
					return (_1102_v32).IsSubsetOf(_1102_v32)
				})()
			} else {
				var _1103___mcc_h8 _dafny.Map = _source20.Get_().(D10_DC21).Cf29
				_ = _1103___mcc_h8
				var _1104_cf29 _dafny.Map = _1103___mcc_h8
				_ = _1104_cf29
				_1045_v2 = _1045_v2
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_1040_v0), 0))
				_ = _index154
				(_1040_v0).ArraySet1((_this).F6(), (_index154).Int())
				var _1105_v33 _dafny.Set
				_ = _1105_v33
				_1105_v33 = _dafny.SetOf(_1077_v19)
				var _1106_v34 _dafny.Map
				_ = _1106_v34
				_1106_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1105_v33)
				var _1107_v35 _dafny.CodePoint
				_ = _1107_v35
				_1107_v35 = _dafny.CodePoint('k')
				var _1108_v36 *C0
				_ = _1108_v36
				var _nw170 *C0 = New_C0_()
				_ = _nw170
				_nw170.Ctor__((_this).F6())
				_1108_v36 = _nw170
				var _1109_v37 _dafny.MultiSet
				_ = _1109_v37
				_1109_v37 = _dafny.MultiSetOf(p0, (_dafny.Zero).Minus((p0).Times((_this).F6())), (_1083_cf1).Times(_dafny.IntOfInt64(386)), p0, (_this).Fm9((_this).F4(), (_this).F4(), (_1108_v36).F12(), globalState))
				var _1110_v38 _dafny.Sequence
				_ = _1110_v38
				_1110_v38 = _dafny.UnicodeSeqOfUtf8Bytes("wcnfyfu")
				var _1111_v39 _dafny.MultiSet
				_ = _1111_v39
				_1111_v39 = _dafny.MultiSetOf(_1109_v37)
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_1040_v0), 0))
				_ = _index155
				var _rhs154 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
					if (_1106_v34).Contains((_this).F6()) {
						return (_1106_v34).Get((_this).F6()).(_dafny.Set)
					}
					return _1105_v33
				})(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1107_v35, _1108_v36))).Cardinality()
				_ = _rhs154
				var _rhs155 _dafny.Int = (func() _dafny.Int {
					if (_1109_v37).Contains(_dafny.IntOfUint32((_1110_v38).Cardinality())) {
						return (_1109_v37).Multiplicity(_dafny.IntOfUint32((_1110_v38).Cardinality()))
					}
					return _1083_cf1
				})()
				_ = _rhs155
				var _rhs156 _dafny.Int = (func() _dafny.Int {
					if (_1111_v39).Contains((_1109_v37).Intersection(_1109_v37)) {
						return (_1111_v39).Multiplicity((_1109_v37).Intersection(_1109_v37))
					}
					return p0
				})()
				_ = _rhs156
				var _rhs157 _dafny.Int = _1083_cf1
				_ = _rhs157
				var _lhs92 _dafny.Array = _1040_v0
				_ = _lhs92
				var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_1040_v0), 0))
				_ = _lhs93
				_1083_cf1 = _rhs154
				_1083_cf1 = _rhs155
				_1083_cf1 = _rhs156
				(_lhs92).ArraySet1(_rhs157, (_lhs93).Int())
				var _1112_v40 _dafny.Map
				_ = _1112_v40
				_1112_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-194), _dafny.IntOfInt64(341)), _dafny.UnicodeSeqOfUtf8Bytes("mou"))
				_1112_v40 = (_1112_v40).Update((_1040_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_1040_v0), 0))).Int()).(_dafny.Int), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(70))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}((func(_1113_v35 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1114_i10 _dafny.Int) _dafny.CodePoint {
						return _1113_v35
					}
				})(_1107_v35))))
				var _1115_v41 _dafny.Map
				_ = _1115_v41
				_1115_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1107_v35)
				var _1116_v42 _dafny.Map
				_ = _1116_v42
				_1116_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1077_v19, (_this).F4())
				var _1117_v43 _dafny.Map
				_ = _1117_v43
				_1117_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1083_cf1, (_1116_v42).Cardinality())
				var _1118_v44 _dafny.Map
				_ = _1118_v44
				_1118_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), false)
				_1115_v41 = (_1115_v41).Update(((_1040_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_1040_v0), 0))).Int()).(_dafny.Int)).Cmp((func() _dafny.Int {
					if (_1117_v43).Contains(_1083_cf1) {
						return (_1117_v43).Get(_1083_cf1).(_dafny.Int)
					}
					return (_1118_v44).Cardinality()
				})()) <= 0, _1107_v35)
			}
			_1077_v19 = _1077_v19
			var _source21 D6 = _1046_v3
			_ = _source21
			if _source21.Is_DC16() {
				var _1119___mcc_h9 bool = _source21.Get_().(D6_DC16).Cf20
				_ = _1119___mcc_h9
				var _1120___mcc_h10 bool = _source21.Get_().(D6_DC16).Cf21
				_ = _1120___mcc_h10
				var _1121___mcc_h11 _dafny.Sequence = _source21.Get_().(D6_DC16).Cf22
				_ = _1121___mcc_h11
				var _1122_cf22 _dafny.Sequence = _1121___mcc_h11
				_ = _1122_cf22
				var _1123_cf21 bool = _1120___mcc_h10
				_ = _1123_cf21
				var _1124_cf20 bool = _1119___mcc_h9
				_ = _1124_cf20
				var _1125_v45 T1
				_ = _1125_v45
				var _nw171 *C1 = New_C1_()
				_ = _nw171
				_nw171.Ctor__((_dafny.Zero).Minus(_1083_cf1), (_this).F6(), _this.F5(), _dafny.IntOfInt64(490), false)
				_1125_v45 = _nw171
				var _1126_v46 _dafny.Map
				_ = _1126_v46
				_1126_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1125_v45, (_1125_v45).F4())
				var _1127_v47 _dafny.Map
				_ = _1127_v47
				_1127_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(p0, _1123_cf21, globalState), _1126_v46)
				_1127_v47 = _1127_v47
				var _1128_v48 _dafny.CodePoint
				_ = _1128_v48
				_1128_v48 = _dafny.CodePoint('r')
				var _1129_v49 _dafny.Map
				_ = _1129_v49
				_1129_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1128_v48, _1128_v48)
				var _1130_v50 _dafny.MultiSet
				_ = _1130_v50
				_1130_v50 = _dafny.MultiSetOf(_1129_v49)
				var _1131_v51 _dafny.Set
				_ = _1131_v51
				_1131_v51 = _dafny.SetOf(_1077_v19, false)
				var _1132_v52 _dafny.MultiSet
				_ = _1132_v52
				_1132_v52 = _dafny.MultiSetOf((_this).F6())
				_1124_cf20 = ((_1132_v52).Union(_1132_v52)).IsSubsetOf(_dafny.MultiSetOf((_1130_v50).Cardinality(), (_1131_v51).Cardinality(), p0, _1083_cf1))
				var _1133_v53 _dafny.MultiSet
				_ = _1133_v53
				_1133_v53 = _dafny.MultiSetOf(_1077_v19)
				_1044_v1 = (Companion_D4_.Create_DC13_((_1125_v45).F6(), _1077_v19, (_1133_v53).Cardinality(), _1044_v1)).Dtor_cf17()
				var _1134_v54 _dafny.Sequence
				_ = _1134_v54
				_1134_v54 = _dafny.UnicodeSeqOfUtf8Bytes("wfi")
				var _1135_v55 _dafny.Map
				_ = _1135_v55
				_1135_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1132_v52, _1134_v54)
				var _1136_v56 bool
				_ = _1136_v56
				var _1137_v57 _dafny.Int
				_ = _1137_v57
				var _out32 bool
				_ = _out32
				var _out33 _dafny.Int
				_ = _out33
				_out32, _out33 = Companion_Default___.M0(_1135_v55, _dafny.Companion_Sequence_.IsPrefixOf(_1122_cf22, _1047_v4), _1132_v52, (_1125_v45).F4(), globalState)
				_1136_v56 = _out32
				_1137_v57 = _out33
			} else {
				var _1138___mcc_h12 _dafny.MultiSet = _source21.Get_().(D6_DC15).Cf19
				_ = _1138___mcc_h12
				var _1139_cf19 _dafny.MultiSet = _1138___mcc_h12
				_ = _1139_cf19
				_1083_cf1 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).Fm9(_1077_v19, true, (_this).F6(), globalState), (func() _dafny.Int {
					if (_1085_v25).Contains((_this).F4()) {
						return (_1085_v25).Get((_this).F4()).(_dafny.Int)
					}
					return p0
				})()))
				var _1140_v58 _dafny.Array
				_ = _1140_v58
				var _nwElement0_39 bool = _1077_v19
				_ = _nwElement0_39
				var _nw172 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(2))
				_ = _nw172
				(_nw172).ArraySet1(_nwElement0_39, 0)
				(_nw172).ArraySet1((_this).F4(), 1)
				_1140_v58 = _nw172
				var _1141_v59 _dafny.Map
				_ = _1141_v59
				_1141_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1140_v58, _1083_cf1)
				var _1142_v60 _dafny.Array
				_ = _1142_v60
				_1142_v60 = _this.F5()
				var _1143_v61 *C1
				_ = _1143_v61
				var _nw173 *C1 = New_C1_()
				_ = _nw173
				_nw173.Ctor__((_this).F6(), (_this).F6(), (_1142_v60), (_this).F6(), _1077_v19)
				_1143_v61 = _nw173
				var _1144_v62 _dafny.Map
				_ = _1144_v62
				_1144_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1140_v58, _1143_v61)
				var _1145_v63 _dafny.Map
				_ = _1145_v63
				_1145_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1083_cf1, (_1144_v62).Cardinality())
				var _1146_v64 T1
				_ = _1146_v64
				var _nw174 *C1 = New_C1_()
				_ = _nw174
				_nw174.Ctor__((_1145_v63).Cardinality(), (_this).F6(), _this.F5(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1047_v4).Cardinality())), (_this).F4())
				_1146_v64 = _nw174
				var _1147_v65 _dafny.Map
				_ = _1147_v65
				_1147_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1146_v64, _1140_v58)
				_1141_v59 = (_1141_v59).Update((func() _dafny.Array {
					if (_1147_v65).Contains(_1146_v64) {
						return (_1147_v65).Get(_1146_v64).(_dafny.Array)
					}
					return _1140_v58
				})(), _1143_v61.F16)
				var _1148_v66 _dafny.MultiSet
				_ = _1148_v66
				_1148_v66 = _dafny.MultiSetOf(false, !(_1077_v19))
				var _1149_v67 D12
				_ = _1149_v67
				_1149_v67 = Companion_D12_.Create_DC28_(p0, (_1148_v66).Cardinality(), false, (_1077_v19) || (_1077_v19), (_this).Fm9(_1077_v19, _1077_v19, _1143_v61.F16, globalState))
				_1149_v67 = _1149_v67
				var _1150_v68 _dafny.Map
				_ = _1150_v68
				_1150_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1143_v61.F16, _dafny.CodePoint('m'))
				(_1143_v61).F16 = Companion_Default___.SafeModuloInt(p0, (_1150_v68).Cardinality())
			}
		} else {
			var _1151___mcc_h2 _dafny.Int = _source19.Get_().(D0_DC0).Cf0
			_ = _1151___mcc_h2
			var _1152_cf0 _dafny.Int = _1151___mcc_h2
			_ = _1152_cf0
			var _1153_v69 _dafny.Sequence
			_ = _1153_v69
			_1153_v69 = _dafny.UnicodeSeqOfUtf8Bytes("qlh")
			var _1154_v70 _dafny.Map
			_ = _1154_v70
			_1154_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(Companion_Default___.Fm39(_1153_v69, Companion_D12_.Create_DC27_(_dafny.SeqOf(_dafny.SeqOf((_this).F4()))), Companion_Default___.Fm0(_dafny.CodePoint('r'), globalState), globalState)), _1153_v69)
			var _1155_v71 bool
			_ = _1155_v71
			var _1156_v72 _dafny.Int
			_ = _1156_v72
			var _out34 bool
			_ = _out34
			var _out35 _dafny.Int
			_ = _out35
			_out34, _out35 = Companion_Default___.M0((_1154_v70).Merge(_1154_v70), (_this).F4(), _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1047_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(68))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg74 _dafny.Int) interface{} {
					return coer74(arg74)
				}
			}((func(_1157_v69 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_1158_i11 _dafny.Int) _dafny.Int {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1157_v69, _dafny.CodePoint('i'))).Cardinality()
				}
			})(_1153_v69))))), !(_1077_v19), globalState)
			_1155_v71 = _out34
			_1156_v72 = _out35
			if false {
				_1155_v71 = _1077_v19
				_1040_v0 = _1040_v0
				var _1159_v74 *C1
				_ = _1159_v74
				var _nw175 *C1 = New_C1_()
				_ = _nw175
				_nw175.Ctor__((_dafny.Zero).Minus(_1156_v72), (_1047_v4).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1047_v4).Cardinality()))).Uint32()).(_dafny.Int), _this.F5(), ((_1048_v5).Union(func() _dafny.Set {
					var _coll41 = _dafny.NewBuilder()
					_ = _coll41
					for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(779), _dafny.IntOfInt64(-489))); ; {
						_compr_41, _ok47 := _iter47()
						if !_ok47 {
							break
						}
						var _1160_v73 _dafny.Int
						_1160_v73 = interface{}(_compr_41).(_dafny.Int)
						if ((_dafny.IntOfInt64(779)).Cmp(_1160_v73) <= 0) && ((_1160_v73).Cmp(_dafny.IntOfInt64(-489)) < 0) {
							_coll41.Add((_1160_v73).Minus(p0))
						}
					}
					return _coll41.ToSet()
				}())).Cardinality(), (_1156_v72).Cmp(_1156_v72) > 0)
				_1159_v74 = _nw175
				_1077_v19 = (_this).F4()
				var _1161_v75 _dafny.Array
				_ = _1161_v75
				var _nw176 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw176
				_1161_v75 = _nw176
				_1161_v75 = _1161_v75
			} else {
				_1152_cf0 = _1156_v72
				var _1162_v76 _dafny.Map
				_ = _1162_v76
				_1162_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1044_v1).Cardinality()), (_this).F6())
				var _1163_v77 _dafny.MultiSet
				_ = _1163_v77
				_1163_v77 = _dafny.MultiSetOf(_dafny.IntOfInt64(226))
				var _1164_v78 _dafny.Sequence
				_ = _1164_v78
				_1164_v78 = _dafny.SeqOf(_1163_v77)
				_1155_v71 = (_1044_v1).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_1162_v76).Cardinality(), ((_1164_v78).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.IntOfUint32((_1164_v78).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()), _dafny.IntOfUint32((_1044_v1).Cardinality()))).Uint32()).(bool)
				var _1165_v79 _dafny.Map
				_ = _1165_v79
				_1165_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1152_cf0).Cmp(_1152_cf0) < 0, _1155_v71)
				_1077_v19 = (func() bool {
					if (_1165_v79).Contains((_this).F4()) {
						return (_1165_v79).Get((_this).F4()).(bool)
					}
					return _1077_v19
				})()
				var _1166_v80 _dafny.CodePoint
				_ = _1166_v80
				_1166_v80 = _dafny.CodePoint('v')
				_1166_v80 = _1166_v80
				var _1167_v81 _dafny.Array
				_ = _1167_v81
				var _nw177 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
				_ = _nw177
				_1167_v81 = _nw177
				_1167_v81 = (Companion_D2_.Create_DC5_(_1167_v81)).Dtor_cf5()
			}
			var _1168_v82 _dafny.CodePoint
			_ = _1168_v82
			_1168_v82 = _dafny.CodePoint('l')
			_1168_v82 = _1168_v82
			var _1169_v85 _dafny.Set
			_ = _1169_v85
			_1169_v85 = _dafny.SetOf((_this).F4(), _1155_v71, Companion_Default___.Fm0(_1168_v82, globalState))
			var _1170_v86 _dafny.Array
			_ = _1170_v86
			var _nwElement0_40 _dafny.Set = _1048_v5
			_ = _nwElement0_40
			var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(6))
			_ = _nw178
			(_nw178).ArraySet1(_nwElement0_40, 0)
			(_nw178).ArraySet1(func() _dafny.Set {
				var _coll42 = _dafny.NewBuilder()
				_ = _coll42
				for _iter48 := _dafny.Iterate((_1047_v4).Elements()); ; {
					_compr_42, _ok48 := _iter48()
					if !_ok48 {
						break
					}
					var _1171_v83 _dafny.Int
					_1171_v83 = interface{}(_compr_42).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_1047_v4, _1171_v83) {
						_coll42.Add(Companion_Default___.SafeDivisionInt(_1171_v83, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-932))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg75 _dafny.Int) interface{} {
								return coer75(arg75)
							}
						}(func(_1172_i12 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('v')
						}))).Cardinality())))
					}
				}
				return _coll42.ToSet()
			}(), 1)
			(_nw178).ArraySet1(_1048_v5, 2)
			(_nw178).ArraySet1((func() _dafny.Set {
				if _1155_v71 {
					return _1048_v5
				}
				return func() _dafny.Set {
					var _coll43 = _dafny.NewBuilder()
					_ = _coll43
					for _iter49 := _dafny.Iterate((_1047_v4).Elements()); ; {
						_compr_43, _ok49 := _iter49()
						if !_ok49 {
							break
						}
						var _1173_v84 _dafny.Int
						_1173_v84 = interface{}(_compr_43).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1047_v4, _1173_v84) {
							_coll43.Add(Companion_Default___.SafeDivisionInt(_1173_v84, (_dafny.Zero).Minus(_dafny.IntOfInt64(-223))))
						}
					}
					return _coll43.ToSet()
				}()
			})(), 3)
			(_nw178).ArraySet1(_1048_v5, 4)
			(_nw178).ArraySet1(Companion_Default___.Fm35((_1169_v85).Cardinality(), globalState), 5)
			_1170_v86 = _nw178
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_1170_v86), 0))
			_ = _index156
			(_1170_v86).ArraySet1(_1048_v5, (_index156).Int())
			var _1174_v87 _dafny.MultiSet
			_ = _1174_v87
			_1174_v87 = _dafny.MultiSetOf(p0)
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_1170_v86), 0))
			_ = _index157
			(_1170_v86).ArraySet1((_1048_v5).Intersection(_dafny.SetOf((func() _dafny.Int {
				if (_1174_v87).Contains((_this).F6()) {
					return (_1174_v87).Multiplicity((_this).F6())
				}
				return (_this).F6()
			})(), p0, (_this).F6(), _1152_cf0, p0)), (_index157).Int())
		}
		var _1175_v88 _dafny.Map
		_ = _1175_v88
		_1175_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1077_v19, (_this).F6())
		r0 = _1175_v88
		return r0
	}
}
func (_this *C3) M2(p0 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Array) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _1176_v0 _dafny.MultiSet
		_ = _1176_v0
		_1176_v0 = _dafny.MultiSetOf((_this).F6())
		var _1177_v1 _dafny.Array
		_ = _1177_v1
		var _nwElement0_41 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
			if (_1176_v0).Contains((_this).F6()) {
				return (_1176_v0).Multiplicity((_this).F6())
			}
			return (_this).F6()
		})())
		_ = _nwElement0_41
		var _nw179 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(3))
		_ = _nw179
		(_nw179).ArraySet1(_nwElement0_41, 0)
		(_nw179).ArraySet1((_this).F6(), 1)
		(_nw179).ArraySet1((_dafny.Zero).Minus(((_this).F6()).Times((_this).Fm7(globalState))), 2)
		_1177_v1 = _nw179
		var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
		_ = _index158
		(_1177_v1).ArraySet1((_this).F6(), (_index158).Int())
		var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
		_ = _index159
		(_1177_v1).ArraySet1((_this).F6(), (_index159).Int())
		var _1178_v2 _dafny.Map
		_ = _1178_v2
		_1178_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int), _1177_v1)
		var _1179_v3 _dafny.Map
		_ = _1179_v3
		_1179_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(442), (func() _dafny.Array {
			if (_1178_v2).Contains((_this).F6()) {
				return (_1178_v2).Get((_this).F6()).(_dafny.Array)
			}
			return _1177_v1
		})())
		var _1180_v4 _dafny.Sequence
		_ = _1180_v4
		_1180_v4 = _dafny.UnicodeSeqOfUtf8Bytes("ivk")
		var _1181_v5 _dafny.Sequence
		_ = _1181_v5
		_1181_v5 = _dafny.SeqOf(_1177_v1, _1177_v1, _1177_v1, _1177_v1)
		var _1182_v6 _dafny.Array
		_ = _1182_v6
		var _nwElement0_42 _dafny.Array = _1177_v1
		_ = _nwElement0_42
		var _nw180 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(24))
		_ = _nw180
		(_nw180).ArraySet1(_nwElement0_42, 0)
		(_nw180).ArraySet1(_1177_v1, 1)
		(_nw180).ArraySet1(_1177_v1, 2)
		(_nw180).ArraySet1(_1177_v1, 3)
		(_nw180).ArraySet1(_1177_v1, 4)
		(_nw180).ArraySet1(_1177_v1, 5)
		(_nw180).ArraySet1(_1177_v1, 6)
		(_nw180).ArraySet1(_1177_v1, 7)
		(_nw180).ArraySet1(_1177_v1, 8)
		(_nw180).ArraySet1(_1177_v1, 9)
		(_nw180).ArraySet1(_1177_v1, 10)
		(_nw180).ArraySet1(_1177_v1, 11)
		(_nw180).ArraySet1((func() _dafny.Array {
			if !((_this).F4()) {
				return _1177_v1
			}
			return _1177_v1
		})(), 12)
		(_nw180).ArraySet1(_1177_v1, 13)
		(_nw180).ArraySet1(_1177_v1, 14)
		(_nw180).ArraySet1((func() _dafny.Array {
			if (_this).F4() {
				return _1177_v1
			}
			return _1177_v1
		})(), 15)
		(_nw180).ArraySet1((func() _dafny.Array {
			if (_1179_v3).Contains(_dafny.IntOfUint32((_1180_v4).Cardinality())) {
				return (_1179_v3).Get(_dafny.IntOfUint32((_1180_v4).Cardinality())).(_dafny.Array)
			}
			return _1177_v1
		})(), 16)
		(_nw180).ArraySet1(_1177_v1, 17)
		(_nw180).ArraySet1(_1177_v1, 18)
		(_nw180).ArraySet1(_1177_v1, 19)
		(_nw180).ArraySet1(_1177_v1, 20)
		(_nw180).ArraySet1(_1177_v1, 21)
		(_nw180).ArraySet1((_1181_v5).Select((Companion_Default___.SafeIndex((_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1181_v5).Cardinality()))).Uint32()).(_dafny.Array), 22)
		(_nw180).ArraySet1(_1177_v1, 23)
		_1182_v6 = _nw180
		var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_1182_v6), 0))
		_ = _index160
		(_1182_v6).ArraySet1(_1177_v1, (_index160).Int())
		var _1183_v7 D2
		_ = _1183_v7
		_1183_v7 = Companion_D2_.Create_DC7_(_dafny.IntOfUint32((Companion_Default___.Fm37(p0, (_this).F4(), (_this).F6(), p0, globalState)).Cardinality()))
		var _1184_v8 _dafny.MultiSet
		_ = _1184_v8
		_1184_v8 = _dafny.MultiSetOf(_1183_v7, _1183_v7)
		var _1185_v9 _dafny.Set
		_ = _1185_v9
		_1185_v9 = _dafny.SetOf(p0, p0)
		var _1186_v10 _dafny.Map
		_ = _1186_v10
		_1186_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1184_v8).Cardinality()).Cmp((_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int)) == 0, ((_1185_v9).Cardinality()).Cmp((_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int)) != 0)
		var _1187_v11 bool
		_ = _1187_v11
		_1187_v11 = true
		var _1188_v12 D16
		_ = _1188_v12
		_1188_v12 = Companion_D16_.Create_DC32_(_1177_v1)
		var _1189_v13 _dafny.Map
		_ = _1189_v13
		_1189_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1177_v1, (_1188_v12).Dtor_cf51())
		var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_1182_v6), 0))
		_ = _index161
		var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
		_ = _index162
		var _rhs158 _dafny.Array = (func() _dafny.Array {
			if (_1189_v13).Contains(_1177_v1) {
				return (_1189_v13).Get(_1177_v1).(_dafny.Array)
			}
			return _1177_v1
		})()
		_ = _rhs158
		var _rhs159 _dafny.Int = (_this).F6()
		_ = _rhs159
		var _rhs160 _dafny.Map = (_1186_v10).Merge((_1186_v10).Merge(_1186_v10))
		_ = _rhs160
		var _rhs161 bool = !((_this).F4()) || ((_this).F4())
		_ = _rhs161
		var _lhs94 _dafny.Array = _1182_v6
		_ = _lhs94
		var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_1182_v6), 0))
		_ = _lhs95
		var _lhs96 _dafny.Array = _1177_v1
		_ = _lhs96
		var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
		_ = _lhs97
		(_lhs94).ArraySet1(_rhs158, (_lhs95).Int())
		(_lhs96).ArraySet1(_rhs159, (_lhs97).Int())
		_1186_v10 = _rhs160
		_1187_v11 = _rhs161
		var _1190_v14 _dafny.Array
		_ = _1190_v14
		var _nw181 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
		_ = _nw181
		_1190_v14 = _nw181
		var _1191_v15 _dafny.Sequence
		_ = _1191_v15
		_1191_v15 = _dafny.SeqOf((_this).F4())
		var _1192_v16 _dafny.Map
		_ = _1192_v16
		_1192_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int))
		var _1193_v17 _dafny.Sequence
		_ = _1193_v17
		_1193_v17 = _dafny.SeqOf(_1192_v16, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus((_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int))), _1192_v16)
		var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
		_ = _index163
		var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
		_ = _index164
		var _rhs162 _dafny.Int = (_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int)
		_ = _rhs162
		var _rhs163 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_this).F4() {
				return _dafny.Companion_Sequence_.Concatenate(_1191_v15, _1191_v15)
			}
			return _dafny.Companion_Sequence_.Concatenate(_1191_v15, _1191_v15)
		})()).Cardinality())
		_ = _rhs163
		var _rhs164 bool = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_1193_v17, _1193_v17), _1192_v16)
		_ = _rhs164
		var _rhs165 bool = _1187_v11
		_ = _rhs165
		var _rhs166 _dafny.Array = _1190_v14
		_ = _rhs166
		var _lhs98 _dafny.Array = _1177_v1
		_ = _lhs98
		var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
		_ = _lhs99
		var _lhs100 _dafny.Array = _1177_v1
		_ = _lhs100
		var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
		_ = _lhs101
		(_lhs98).ArraySet1(_rhs162, (_lhs99).Int())
		(_lhs100).ArraySet1(_rhs163, (_lhs101).Int())
		_1187_v11 = _rhs164
		_1187_v11 = _rhs165
		_1190_v14 = _rhs166
		var _1194_v18 T0
		_ = _1194_v18
		var _nw182 *C2 = New_C2_()
		_ = _nw182
		_nw182.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(587))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg76 _dafny.Int) interface{} {
				return coer76(arg76)
			}
		}(func(_1195_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		})), p0)
		_1194_v18 = _nw182
		var _1196_v19 _dafny.Map
		_ = _1196_v19
		_1196_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1194_v18).F4(), _1194_v18)
		var _1197_v20 _dafny.Map
		_ = _1197_v20
		_1197_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1194_v18)
		var _1198_v21 _dafny.Sequence
		_ = _1198_v21
		_1198_v21 = _dafny.SeqOf(_1194_v18, (func() T0 {
			if (_1196_v19).Contains(p0) {
				return (_1196_v19).Get(p0).(T0)
			}
			return (func() T0 {
				if (_1197_v20).Contains(_dafny.IntOfInt64(-228)) {
					return (_1197_v20).Get(_dafny.IntOfInt64(-228)).(T0)
				}
				return _1194_v18
			})()
		})(), _1194_v18, _1194_v18, _1194_v18)
		var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
		_ = _index165
		(_1177_v1).ArraySet1((_dafny.IntOfUint32((_1198_v21).Cardinality())).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-701))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg77 _dafny.Int) interface{} {
				return coer77(arg77)
			}
		}(func(_1199_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		}))).Cardinality())), (_index165).Int())
		var _1200_v22 _dafny.Array
		_ = _1200_v22
		var _len0_21 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_21
		var _nw183 _dafny.Array
		_ = _nw183
		if _len0_21.Cmp(_dafny.Zero) == 0 {
			_nw183 = _dafny.NewArray(_len0_21)
		} else {
			var _init21 func(_dafny.Int) _dafny.Sequence = func(_1201_i2 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("q"), _dafny.UnicodeSeqOfUtf8Bytes("fag"))
			}
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
		_1200_v22 = _nw183
		var _1202_v23 _dafny.Sequence
		_ = _1202_v23
		_1202_v23 = _dafny.SeqOf(_1180_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(819))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg78 _dafny.Int) interface{} {
				return coer78(arg78)
			}
		}(func(_1203_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		})))
		var _1204_v24 _dafny.CodePoint
		_ = _1204_v24
		_1204_v24 = _dafny.CodePoint('f')
		var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1200_v22), 0))
		_ = _index166
		(_1200_v22).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_1202_v23).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1202_v23).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32(((_1202_v23).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1202_v23).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _1204_v24), _dafny.UnicodeSeqOfUtf8Bytes("nusipv")), (_index166).Int())
		var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1200_v22), 0))
		_ = _index167
		(_1200_v22).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.Zero)).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg79 _dafny.Int) interface{} {
				return coer79(arg79)
			}
		}(func(_1205_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		})), (_index167).Int())
		if Companion_Default___.Fm0(_1204_v24, globalState) {
			var _1206_v25 _dafny.Array
			_ = _1206_v25
			var _nw184 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(21))
			_ = _nw184
			_1206_v25 = _nw184
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_1190_v14), 0))
			_ = _index168
			(_1190_v14).ArraySet1(_1206_v25, (_index168).Int())
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_1190_v14), 0))
			_ = _index169
			(_1190_v14).ArraySet1(_1206_v25, (_index169).Int())
			_1187_v11 = !(p0)
			var _1207_v26 D11
			_ = _1207_v26
			_1207_v26 = Companion_D11_.Create_DC26_((_1194_v18).F4())
			var _1208_v27 _dafny.Map
			_ = _1208_v27
			_1208_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1207_v26, (_1194_v18).F4())
			var _1209_v28 _dafny.Sequence
			_ = _1209_v28
			_1209_v28 = _dafny.SeqOf(_1208_v27)
			var _1210_v29 D18
			_ = _1210_v29
			_1210_v29 = Companion_D18_.Create_DC35_(_1209_v28)
			var _1211_v30 _dafny.Sequence
			_ = _1211_v30
			_1211_v30 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(451))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg80 _dafny.Int) interface{} {
					return coer80(arg80)
				}
			}((func(_1212_v24 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1213_i5 _dafny.Int) _dafny.CodePoint {
					return _1212_v24
				}
			})(_1204_v24)))).Cardinality()), (_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int))
			var _1214_v31 _dafny.Sequence
			_ = _1214_v31
			_1214_v31 = _dafny.SeqOf((_1211_v30).Select((Companion_Default___.SafeIndex((_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1211_v30).Cardinality()))).Uint32()).(_dafny.Int))
			var _1215_v32 _dafny.Map
			_ = _1215_v32
			_1215_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1211_v30).Cardinality()), _1186_v10)
			var _1216_v33 _dafny.Sequence
			_ = _1216_v33
			_1216_v33 = _dafny.SeqOf(_1215_v32, _1215_v32)
			var _1217_v34 *C1
			_ = _1217_v34
			var _nw185 *C1 = New_C1_()
			_ = _nw185
			_nw185.Ctor__((_dafny.Zero).Minus(_dafny.IntOfUint32(((_1210_v29).Dtor_cf58()).Cardinality())), ((_1216_v33).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_1187_v11)).Cardinality(), _dafny.IntOfUint32((_1216_v33).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _this.F5(), ((_1214_v31).Select((Companion_Default___.SafeIndex((_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1214_v31).Cardinality()))).Uint32()).(_dafny.Int)).Minus((_1194_v18).Fm7(globalState)), true)
			_1217_v34 = _nw185
			var _1218_v35 _dafny.Set
			_ = _1218_v35
			_1218_v35 = _dafny.SetOf(_dafny.MultiSetFromSeq(_1191_v15))
			var _1219_v36 _dafny.Array
			_ = _1219_v36
			var _nwElement0_43 bool = Companion_Default___.Fm0(_1204_v24, globalState)
			_ = _nwElement0_43
			var _nw186 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(27))
			_ = _nw186
			(_nw186).ArraySet1(_nwElement0_43, 0)
			(_nw186).ArraySet1(!(((_this).F6()).Cmp((_1217_v34).F15()) <= 0), 1)
			(_nw186).ArraySet1((_1218_v35).IsSubsetOf(_1218_v35), 2)
			(_nw186).ArraySet1((_this).F4(), 3)
			(_nw186).ArraySet1(false, 4)
			(_nw186).ArraySet1(_1187_v11, 5)
			(_nw186).ArraySet1((_this).F4(), 6)
			(_nw186).ArraySet1((_this).F4(), 7)
			(_nw186).ArraySet1(true, 8)
			(_nw186).ArraySet1((_1194_v18).F4(), 9)
			(_nw186).ArraySet1((_this).F4(), 10)
			(_nw186).ArraySet1(p0, 11)
			(_nw186).ArraySet1(_1187_v11, 12)
			(_nw186).ArraySet1((_1185_v9).IsSubsetOf(_1185_v9), 13)
			(_nw186).ArraySet1((_1194_v18).F4(), 14)
			(_nw186).ArraySet1((_this).F4(), 15)
			(_nw186).ArraySet1(p0, 16)
			(_nw186).ArraySet1(((_1194_v18).F4()) || (p0), 17)
			(_nw186).ArraySet1(p0, 18)
			(_nw186).ArraySet1((_1194_v18).F4(), 19)
			(_nw186).ArraySet1((_this).F4(), 20)
			(_nw186).ArraySet1(p0, 21)
			(_nw186).ArraySet1(p0, 22)
			(_nw186).ArraySet1(_1187_v11, 23)
			(_nw186).ArraySet1(true, 24)
			(_nw186).ArraySet1((_1194_v18).F4(), 25)
			(_nw186).ArraySet1(p0, 26)
			_1219_v36 = _nw186
			var _1220_v37 _dafny.Set
			_ = _1220_v37
			_1220_v37 = _dafny.SetOf((_1200_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1200_v22), 0))).Int()).(_dafny.Sequence))
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_1219_v36), 0))
			_ = _index170
			(_1219_v36).ArraySet1((_1220_v37).IsProperSubsetOf(_1220_v37), (_index170).Int())
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
			_ = _index171
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_1219_v36), 0))
			_ = _index172
			var _rhs167 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F6(), ((_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int)).Times(_1217_v34.F16))
			_ = _rhs167
			var _rhs168 bool = p0
			_ = _rhs168
			var _rhs169 bool = (_1194_v18).F4()
			_ = _rhs169
			var _lhs102 _dafny.Array = _1177_v1
			_ = _lhs102
			var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
			_ = _lhs103
			var _lhs104 _dafny.Array = _1219_v36
			_ = _lhs104
			var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_1219_v36), 0))
			_ = _lhs105
			(_lhs102).ArraySet1(_rhs167, (_lhs103).Int())
			_1187_v11 = _rhs168
			(_lhs104).ArraySet1(_rhs169, (_lhs105).Int())
			_1187_v11 = !((Companion_D12_.Create_DC28_(Companion_Default___.Fm3(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1191_v15, (Companion_Default___.SafeIndex((_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1191_v15).Cardinality()))).Uint32(), !(p0)), (Companion_Default___.SafeIndex(_1217_v34.F16, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1191_v15, (Companion_Default___.SafeIndex((_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1191_v15).Cardinality()))).Uint32(), !(p0))).Cardinality()))).Uint32(), (_1219_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_1219_v36), 0))).Int()).(bool))).Cardinality()), (_this).F4(), globalState), _1217_v34.F16, (_1194_v18).F4(), (_1219_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_1219_v36), 0))).Int()).(bool), (_this).F6())).Dtor_cf46())
		} else {
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
			_ = _index173
			(_1177_v1).ArraySet1((_this).F6(), (_index173).Int())
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))
			_ = _index174
			(_1177_v1).ArraySet1((_this).F6(), (_index174).Int())
			_1187_v11 = true
			var _1221_v38 *C2
			_ = _1221_v38
			var _nw187 *C2 = New_C2_()
			_ = _nw187
			_nw187.Ctor__((_1200_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1200_v22), 0))).Int()).(_dafny.Sequence), _1187_v11)
			_1221_v38 = _nw187
			var _1222_v39 _dafny.MultiSet
			_ = _1222_v39
			_1222_v39 = _dafny.MultiSetOf(_1221_v38, _1221_v38, _1221_v38)
			var _1223_v40 _dafny.Array
			_ = _1223_v40
			var _nwElement0_44 bool = !((_1222_v39).Update(_1221_v38, Companion_Default___.Abs((_this).F6()))).Contains(_1221_v38)
			_ = _nwElement0_44
			var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(10))
			_ = _nw188
			(_nw188).ArraySet1(_nwElement0_44, 0)
			(_nw188).ArraySet1(true, 1)
			(_nw188).ArraySet1(Companion_Default___.Fm0(_1204_v24, globalState), 2)
			(_nw188).ArraySet1((func() bool {
				if (_1194_v18).F4() {
					return _1187_v11
				}
				return true
			})(), 3)
			(_nw188).ArraySet1(!(_1185_v9).Equals(_1185_v9), 4)
			(_nw188).ArraySet1((_1176_v0).Equals(_1176_v0), 5)
			(_nw188).ArraySet1(((_this).F6()).Cmp((_this).F6()) > 0, 6)
			(_nw188).ArraySet1(_1187_v11, 7)
			(_nw188).ArraySet1((_this).F4(), 8)
			(_nw188).ArraySet1((_1194_v18).F4(), 9)
			_1223_v40 = _nw188
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_1223_v40), 0))
			_ = _index175
			(_1223_v40).ArraySet1(_1187_v11, (_index175).Int())
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_1223_v40), 0))
			_ = _index176
			(_1223_v40).ArraySet1((_1194_v18).F4(), (_index176).Int())
			var _1224_v41 *C1
			_ = _1224_v41
			var _nw189 *C1 = New_C1_()
			_ = _nw189
			_nw189.Ctor__((_1177_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1177_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(806), _this.F5(), ((_this).F6()).Times((_this).F6()), _1187_v11)
			_1224_v41 = _nw189
		}
		r0 = _1204_v24
		var _1225_v42 _dafny.Array
		_ = _1225_v42
		var _nw190 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(23))
		_ = _nw190
		_1225_v42 = _nw190
		r1 = _1225_v42
		return r0, r1
	}
}
func (_this *C3) M9(p0 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		r0 = !(p0)
		var _1226_v0 _dafny.Array
		_ = _1226_v0
		var _len0_22 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_22
		var _nw191 _dafny.Array
		_ = _nw191
		if _len0_22.Cmp(_dafny.Zero) == 0 {
			_nw191 = _dafny.NewArray(_len0_22)
		} else {
			var _init22 func(_dafny.Int) _dafny.Sequence = func(_1227_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(773))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg81 _dafny.Int) interface{} {
						return coer81(arg81)
					}
				}(func(_1228_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				})), _dafny.UnicodeSeqOfUtf8Bytes("uqvrar"))
			}
			_ = _init22
			var _element0_22 = _init22(_dafny.Zero)
			_ = _element0_22
			_nw191 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
			(_nw191).ArraySet1(_element0_22, 0)
			var _nativeLen0_22 = (_len0_22).Int()
			_ = _nativeLen0_22
			for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
				(_nw191).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
			}
		}
		_1226_v0 = _nw191
		var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1226_v0), 0))
		_ = _index177
		(_1226_v0).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("t"), (_index177).Int())
		var _1229_v1 _dafny.Sequence
		_ = _1229_v1
		_1229_v1 = _dafny.UnicodeSeqOfUtf8Bytes("w")
		var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1226_v0), 0))
		_ = _index178
		(_1226_v0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm37(!((_this).F4()), (_this).F4(), _dafny.IntOfUint32((_1229_v1).Cardinality()), (_this).F4(), globalState), _1229_v1), (_index178).Int())
		var _1230_v2 _dafny.CodePoint
		_ = _1230_v2
		_1230_v2 = _dafny.CodePoint('v')
		var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1226_v0), 0))
		_ = _index179
		(_1226_v0).ArraySet1(_dafny.Companion_Sequence_.Update(_1229_v1, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F6()), _dafny.IntOfUint32((_1229_v1).Cardinality()))).Uint32(), _1230_v2), (_index179).Int())
		var _1231_v3 _dafny.Sequence
		_ = _1231_v3
		_1231_v3 = _dafny.SeqOf((_this).F6(), (_this).F6(), (_this).F6())
		var _1232_v4 *C1
		_ = _1232_v4
		var _nw192 *C1 = New_C1_()
		_ = _nw192
		_nw192.Ctor__((_dafny.Zero).Minus((_this).F6()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1231_v3, _1231_v3)).Cardinality()), _this.F5(), (_this).F6(), !((_this).F4()))
		_1232_v4 = _nw192
		if !((_this).F4()) {
			var _1233_v5 _dafny.Array
			_ = _1233_v5
			var _len0_23 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_23
			var _nw193 _dafny.Array
			_ = _nw193
			if _len0_23.Cmp(_dafny.Zero) == 0 {
				_nw193 = _dafny.NewArray(_len0_23)
			} else {
				var _init23 func(_dafny.Int) bool = func(_1234_i2 _dafny.Int) bool {
					return (_this).F4()
				}
				_ = _init23
				var _element0_23 = _init23(_dafny.Zero)
				_ = _element0_23
				_nw193 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
				(_nw193).ArraySet1(_element0_23, 0)
				var _nativeLen0_23 = (_len0_23).Int()
				_ = _nativeLen0_23
				for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
					(_nw193).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
				}
			}
			_1233_v5 = _nw193
			var _1235_v6 _dafny.Map
			_ = _1235_v6
			_1235_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1233_v5)
			var _1236_v7 _dafny.Sequence
			_ = _1236_v7
			_1236_v7 = _dafny.SeqOf(_1233_v5, _1233_v5, _1233_v5)
			_1233_v5 = (func() _dafny.Array {
				if (_1235_v6).Contains((_this).F4()) {
					return (_1235_v6).Get((_this).F4()).(_dafny.Array)
				}
				return (_1236_v7).Select((Companion_Default___.SafeIndex((_1232_v4).F15(), _dafny.IntOfUint32((_1236_v7).Cardinality()))).Uint32()).(_dafny.Array)
			})()
			var _1237_v8 *C2
			_ = _1237_v8
			var _nw194 *C2 = New_C2_()
			_ = _nw194
			_nw194.Ctor__(_1229_v1, p0)
			_1237_v8 = _nw194
			(_1232_v4).F16 = (_1232_v4).F15()
			var _1238_v9 _dafny.Map
			_ = _1238_v9
			_1238_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1232_v4).F15(), (_this).F4())
			_1238_v9 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1232_v4).F15(), (_this).F4())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1232_v4).F15(), p0))
			var _1239_v10 D10
			_ = _1239_v10
			_1239_v10 = Companion_D10_.Create_DC22_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(791))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg82 _dafny.Int) interface{} {
					return coer82(arg82)
				}
			}((func(_1240_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1241_i3 _dafny.Int) _dafny.CodePoint {
					return _1240_v2
				}
			})(_1230_v2))), false, (_this).F4(), _1233_v5, _dafny.IntOfInt64(184))
			var _1242_v11 D1
			_ = _1242_v11
			_1242_v11 = Companion_D1_.Create_DC4_((_1239_v10).Dtor_cf34(), (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lroyicek")).Cardinality())).Minus(_dafny.IntOfInt64(455)))
			_1242_v11 = _1242_v11
		} else {
			r0 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("vtossvsa"), _dafny.Companion_Sequence_.Concatenate((_1226_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1226_v0), 0))).Int()).(_dafny.Sequence), (_1226_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1226_v0), 0))).Int()).(_dafny.Sequence)))
			r0 = ((func() _dafny.Int {
				if (_this).F4() {
					return _1232_v4.F16
				}
				return (_this).F6()
			})()).Cmp(_dafny.IntOfUint32(((_1226_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1226_v0), 0))).Int()).(_dafny.Sequence)).Cardinality())) >= 0
			var _1243_v12 _dafny.Array
			_ = _1243_v12
			var _nwElement0_45 _dafny.CodePoint = _1230_v2
			_ = _nwElement0_45
			var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(2))
			_ = _nw195
			(_nw195).ArraySet1CodePoint(_nwElement0_45, 0)
			(_nw195).ArraySet1CodePoint(_1230_v2, 1)
			_1243_v12 = _nw195
			var _1244_v13 _dafny.Map
			_ = _1244_v13
			_1244_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1243_v12)
			var _1245_v14 _dafny.Sequence
			_ = _1245_v14
			_1245_v14 = _dafny.SeqOf(_1243_v12)
			var _1246_v15 _dafny.Array
			_ = _1246_v15
			var _nwElement0_46 _dafny.Array = _1243_v12
			_ = _nwElement0_46
			var _nw196 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(8))
			_ = _nw196
			(_nw196).ArraySet1(_nwElement0_46, 0)
			(_nw196).ArraySet1(_1243_v12, 1)
			(_nw196).ArraySet1(_1243_v12, 2)
			(_nw196).ArraySet1(_1243_v12, 3)
			(_nw196).ArraySet1(_1243_v12, 4)
			(_nw196).ArraySet1((func() _dafny.Array {
				if (_1244_v13).Contains(p0) {
					return (_1244_v13).Get(p0).(_dafny.Array)
				}
				return (_1245_v14).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1232_v4).F15()), _dafny.IntOfUint32((_1245_v14).Cardinality()))).Uint32()).(_dafny.Array)
			})(), 5)
			(_nw196).ArraySet1(_1243_v12, 6)
			(_nw196).ArraySet1(_1243_v12, 7)
			_1246_v15 = _nw196
			var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1246_v15), 0))
			_ = _index180
			(_1246_v15).ArraySet1(_1243_v12, (_index180).Int())
			var _1247_v16 _dafny.MultiSet
			_ = _1247_v16
			_1247_v16 = _dafny.MultiSetOf(_1232_v4.F16)
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1226_v0), 0))
			_ = _index181
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1246_v15), 0))
			_ = _index182
			var _rhs170 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bsbsmmac"), _dafny.UnicodeSeqOfUtf8Bytes("ps")), (_1226_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1226_v0), 0))).Int()).(_dafny.Sequence))
			_ = _rhs170
			var _rhs171 _dafny.Int = (func() _dafny.Int {
				if (_1247_v16).Contains(_1232_v4.F16) {
					return (_1247_v16).Multiplicity(_1232_v4.F16)
				}
				return (_1232_v4).F15()
			})()
			_ = _rhs171
			var _rhs172 _dafny.Array = _1243_v12
			_ = _rhs172
			var _lhs106 _dafny.Array = _1226_v0
			_ = _lhs106
			var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1226_v0), 0))
			_ = _lhs107
			var _lhs108 *C1 = _1232_v4
			_ = _lhs108
			var _lhs109 _dafny.Array = _1246_v15
			_ = _lhs109
			var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1246_v15), 0))
			_ = _lhs110
			(_lhs106).ArraySet1(_rhs170, (_lhs107).Int())
			_lhs108.F16 = _rhs171
			(_lhs109).ArraySet1(_rhs172, (_lhs110).Int())
			var _1248_v18 _dafny.Array
			_ = _1248_v18
			var _len0_24 _dafny.Int = _dafny.One
			_ = _len0_24
			var _nw197 _dafny.Array
			_ = _nw197
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw197 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) _dafny.Set = (func(_1249_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
					return func(_1250_i4 _dafny.Int) _dafny.Set {
						return func() _dafny.Set {
							var _coll44 = _dafny.NewBuilder()
							_ = _coll44
							for _iter50 := _dafny.Iterate((_1249_v3).Elements()); ; {
								_compr_44, _ok50 := _iter50()
								if !_ok50 {
									break
								}
								var _1251_v17 _dafny.Int
								_1251_v17 = interface{}(_compr_44).(_dafny.Int)
								if _dafny.Companion_Sequence_.Contains(_1249_v3, _1251_v17) {
									_coll44.Add((_1251_v17).Plus(_dafny.IntOfInt64(-142)))
								}
							}
							return _coll44.ToSet()
						}()
					}
				})(_1231_v3)
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw197 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw197).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw197).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_1248_v18 = _nw197
			var _1252_v19 _dafny.Array
			_ = _1252_v19
			var _nwElement0_47 _dafny.Array = _1248_v18
			_ = _nwElement0_47
			var _nw198 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(4))
			_ = _nw198
			(_nw198).ArraySet1(_nwElement0_47, 0)
			(_nw198).ArraySet1((func() _dafny.Array {
				if p0 {
					return _1248_v18
				}
				return _1248_v18
			})(), 1)
			(_nw198).ArraySet1((func() _dafny.Array {
				if p0 {
					return _1248_v18
				}
				return _1248_v18
			})(), 2)
			(_nw198).ArraySet1(_1248_v18, 3)
			_1252_v19 = _nw198
			var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_1252_v19), 0))
			_ = _index183
			(_1252_v19).ArraySet1(_1248_v18, (_index183).Int())
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_1252_v19), 0))
			_ = _index184
			(_1252_v19).ArraySet1((func() _dafny.Array {
				if (_this).F4() {
					return _1248_v18
				}
				return _1248_v18
			})(), (_index184).Int())
			if true {
				var _1253_v20 _dafny.Array
				_ = _1253_v20
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_25
				var _nw199 _dafny.Array
				_ = _nw199
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw199 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) _dafny.Map = (func(_1254_v16 _dafny.MultiSet) func(_dafny.Int) _dafny.Map {
						return func(_1255_i5 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_1254_v16).Cardinality())
						}
					})(_1247_v16)
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
				_1253_v20 = _nw199
				var _1256_v21 _dafny.Map
				_ = _1256_v21
				_1256_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _1232_v4.F16)
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1253_v20), 0))
				_ = _index185
				(_1253_v20).ArraySet1(_1256_v21, (_index185).Int())
				var _1257_v22 _dafny.Map
				_ = _1257_v22
				_1257_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F4())
				var _1258_v23 _dafny.Map
				_ = _1258_v23
				_1258_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (func() bool {
					if (_1257_v22).Contains((_this).F4()) {
						return (_1257_v22).Get((_this).F4()).(bool)
					}
					return p0
				})())
				var _1259_v24 _dafny.Sequence
				_ = _1259_v24
				_1259_v24 = _dafny.SeqOf(_1257_v22, _1258_v23, _1257_v22)
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1246_v15), 0))
				_ = _index186
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1253_v20), 0))
				_ = _index187
				var _rhs173 bool = _dafny.Companion_Sequence_.Equal(_1259_v24, _1259_v24)
				_ = _rhs173
				var _rhs174 _dafny.Array = _1243_v12
				_ = _rhs174
				var _rhs175 _dafny.Map = _1256_v21
				_ = _rhs175
				var _rhs176 _dafny.MultiSet = _1247_v16
				_ = _rhs176
				var _lhs111 _dafny.Array = _1246_v15
				_ = _lhs111
				var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1246_v15), 0))
				_ = _lhs112
				var _lhs113 _dafny.Array = _1253_v20
				_ = _lhs113
				var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1253_v20), 0))
				_ = _lhs114
				r0 = _rhs173
				(_lhs111).ArraySet1(_rhs174, (_lhs112).Int())
				(_lhs113).ArraySet1(_rhs175, (_lhs114).Int())
				_1247_v16 = _rhs176
				var _1260_v25 _dafny.Array
				_ = _1260_v25
				var _nw200 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(22))
				_ = _nw200
				_1260_v25 = _nw200
				var _1261_v26 _dafny.MultiSet
				_ = _1261_v26
				_1261_v26 = _dafny.MultiSetOf(p0)
				var _1262_v27 _dafny.Map
				_ = _1262_v27
				_1262_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1261_v26).Cardinality(), _dafny.IntOfInt64(548))
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((_1260_v25), 0))
				_ = _index188
				(_1260_v25).ArraySet1(_1262_v27, (_index188).Int())
				var _1263_v28 _dafny.Set
				_ = _1263_v28
				_1263_v28 = _dafny.SetOf(p0)
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((_1260_v25), 0))
				_ = _index189
				(_1260_v25).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1263_v28).Cardinality(), _dafny.IntOfInt64(580)), (_index189).Int())
				r0 = (_this).F4()
				(_1232_v4).F16 = _1232_v4.F16
				(_1232_v4).F16 = (_1232_v4.F16).Plus(((_1232_v4).F15()).Times(_1232_v4.F16))
			} else {
				var _1264_v29 _dafny.Map
				_ = _1264_v29
				_1264_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(4))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}((func(_1265_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1266_i6 _dafny.Int) _dafny.CodePoint {
						return _1265_v2
					}
				})(_1230_v2))), _dafny.IntOfInt64(307))
				var _1267_v30 _dafny.Map
				_ = _1267_v30
				_1267_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _1264_v29)
				var _1268_v31 _dafny.Array
				_ = _1268_v31
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_26
				var _nw201 _dafny.Array
				_ = _nw201
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw201 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) bool = func(_1269_i7 _dafny.Int) bool {
						return true
					}
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
				_1268_v31 = _nw201
				var _1270_v32 _dafny.Int
				_ = _1270_v32
				var _out36 _dafny.Int
				_ = _out36
				_out36 = (_1232_v4).M11(p0, p0, (func() _dafny.Map {
					if (_1267_v30).Contains(p0) {
						return (_1267_v30).Get(p0).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1226_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1226_v0), 0))).Int()).(_dafny.Sequence), _1232_v4.F16)
				})(), _1268_v31, globalState)
				_1270_v32 = _out36
				var _1271_v33 _dafny.Map
				_ = _1271_v33
				_1271_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1232_v4).F15(), p0)
				var _1272_v34 _dafny.Sequence
				_ = _1272_v34
				_1272_v34 = _dafny.SeqOf((func() bool {
					if (_1271_v33).Contains((_1232_v4).F15()) {
						return (_1271_v33).Get((_1232_v4).F15()).(bool)
					}
					return false
				})())
				var _1273_v35 _dafny.Sequence
				_ = _1273_v35
				_1273_v35 = _dafny.SeqOf(_1272_v34, _1272_v34)
				var _1274_v36 D12
				_ = _1274_v36
				_1274_v36 = Companion_D12_.Create_DC27_(_1273_v35)
				var _pat_let_tv16 = _1273_v35
				_ = _pat_let_tv16
				var _pat_let_tv17 = _1229_v1
				_ = _pat_let_tv17
				var _pat_let_tv18 = _1273_v35
				_ = _pat_let_tv18
				var _pat_let_tv19 = _1273_v35
				_ = _pat_let_tv19
				var _pat_let_tv20 = _1232_v4
				_ = _pat_let_tv20
				var _pat_let_tv21 = _1273_v35
				_ = _pat_let_tv21
				var _pat_let_tv22 = _1273_v35
				_ = _pat_let_tv22
				_1274_v36 = func(_pat_let22_0 D12) D12 {
					return func(_1277_dt__update__tmp_h1 D12) D12 {
						return func(_pat_let25_0 _dafny.Sequence) D12 {
							return func(_1278_dt__update_hcf42_h1 _dafny.Sequence) D12 {
								return Companion_D12_.Create_DC27_(_1278_dt__update_hcf42_h1)
							}(_pat_let25_0)
						}(_pat_let_tv22)
					}(_pat_let22_0)
				}(func(_pat_let23_0 D12) D12 {
					return func(_1275_dt__update__tmp_h0 D12) D12 {
						return func(_pat_let24_0 _dafny.Sequence) D12 {
							return func(_1276_dt__update_hcf42_h0 _dafny.Sequence) D12 {
								return Companion_D12_.Create_DC27_(_1276_dt__update_hcf42_h0)
							}(_pat_let24_0)
						}(_dafny.Companion_Sequence_.Update(_pat_let_tv16, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_pat_let_tv17).Cardinality()), _dafny.IntOfUint32((_pat_let_tv18).Cardinality()))).Uint32(), (_pat_let_tv19).Select((Companion_Default___.SafeIndex(_pat_let_tv20.F16, _dafny.IntOfUint32((_pat_let_tv21).Cardinality()))).Uint32()).(_dafny.Sequence)))
					}(_pat_let23_0)
				}(_1274_v36))
				var _arr5 _dafny.Array = _this.F5()
				_ = _arr5
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_this.F5()), 0))
				_ = _index190
				(_arr5).ArraySet1(_1268_v31, (_index190).Int())
				var _arr6 _dafny.Array = _this.F5()
				_ = _arr6
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_this.F5()), 0))
				_ = _index191
				(_arr6).ArraySet1(_1268_v31, (_index191).Int())
				var _1279_v37 _dafny.Map
				_ = _1279_v37
				_1279_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _1268_v31)
				var _1280_v38 _dafny.Map
				_ = _1280_v38
				_1280_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1230_v2, (func() _dafny.Array {
					if (_1279_v37).Contains((_this).F4()) {
						return (_1279_v37).Get((_this).F4()).(_dafny.Array)
					}
					return _dafny.ArrayCastTo((_this.F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_this.F5()), 0))).Int()))
				})())
				var _1281_v39 D19
				_ = _1281_v39
				_1281_v39 = Companion_D19_.Create_DC37_(_1280_v38)
				var _rhs177 bool = (_this).F4()
				_ = _rhs177
				var _rhs178 _dafny.Map = (_1281_v39).Dtor_cf62()
				_ = _rhs178
				r0 = _rhs177
				_1280_v38 = _rhs178
				var _1282_v40 _dafny.CodePoint
				_ = _1282_v40
				var _1283_v41 _dafny.Array
				_ = _1283_v41
				var _out37 _dafny.CodePoint
				_ = _out37
				var _out38 _dafny.Array
				_ = _out38
				_out37, _out38 = (_1232_v4).M2(p0, globalState)
				_1282_v40 = _out37
				_1283_v41 = _out38
			}
		}
		var _1284_v42 _dafny.Map
		_ = _1284_v42
		_1284_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _1285_v43 _dafny.Map
		_ = _1285_v43
		_1285_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1284_v42, _1230_v2)
		var _1286_v44 _dafny.Map
		_ = _1286_v44
		_1286_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.UnicodeSeqOfUtf8Bytes("mcrpadd"))
		var _1287_v45 bool
		_ = _1287_v45
		var _1288_v46 _dafny.Int
		_ = _1288_v46
		var _1289_v47 _dafny.Int
		_ = _1289_v47
		var _1290_v48 bool
		_ = _1290_v48
		var _out39 bool
		_ = _out39
		var _out40 _dafny.Int
		_ = _out40
		var _out41 _dafny.Int
		_ = _out41
		var _out42 bool
		_ = _out42
		_out39, _out40, _out41, _out42 = (_this).M10(_dafny.MultiSetOf((_1285_v43).Cardinality(), (_this).F6(), _1232_v4.F16), _dafny.Companion_Sequence_.IsProperPrefixOf(_1229_v1, (func() _dafny.Sequence {
			if (_1286_v44).Contains(_1232_v4.F16) {
				return (_1286_v44).Get(_1232_v4.F16).(_dafny.Sequence)
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("h")
		})()), globalState)
		_1287_v45 = _out39
		_1288_v46 = _out40
		_1289_v47 = _out41
		_1290_v48 = _out42
		r0 = p0
		return r0
	}
}
func (_this *C3) M10(p0 _dafny.MultiSet, p1 bool, globalState *GlobalState) (bool, _dafny.Int, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 bool = false
		_ = r3
		var _1291_v0 _dafny.Array
		_ = _1291_v0
		var _nw202 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw202
		_1291_v0 = _nw202
		for _iter51 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1291_v0), 0))); ; {
			_guard_loop_6, _ok51 := _iter51()
			if !_ok51 {
				break
			}
			var _1292_i0 _dafny.Int
			_1292_i0 = interface{}(_guard_loop_6).(_dafny.Int)
			if (true) && (((_1292_i0).Sign() != -1) && ((_1292_i0).Cmp(_dafny.ArrayLen((_1291_v0), 0)) < 0)) {
				(_1291_v0).ArraySet1(Companion_Default___.SafeModuloInt(_1292_i0, (_this).F6()), (_1292_i0).Int())
			}
		}
		var _hi6 _dafny.Int = _dafny.IntOfInt64(-365)
		_ = _hi6
		for _1293_i1 := (_this).F6(); _1293_i1.Cmp(_hi6) < 0; _1293_i1 = _1293_i1.Plus(_dafny.One) {
			var _1294_v1 D16
			_ = _1294_v1
			_1294_v1 = Companion_D16_.Create_DC32_(_1291_v0)
			var _source22 D16 = _1294_v1
			_ = _source22
			if _source22.Is_DC33() {
				var _1295___mcc_h0 bool = _source22.Get_().(D16_DC33).Cf52
				_ = _1295___mcc_h0
				var _1296___mcc_h1 _dafny.Int = _source22.Get_().(D16_DC33).Cf53
				_ = _1296___mcc_h1
				var _1297___mcc_h2 _dafny.MultiSet = _source22.Get_().(D16_DC33).Cf54
				_ = _1297___mcc_h2
				var _1298___mcc_h3 _dafny.Int = _source22.Get_().(D16_DC33).Cf55
				_ = _1298___mcc_h3
				var _1299___mcc_h4 _dafny.Set = _source22.Get_().(D16_DC33).Cf56
				_ = _1299___mcc_h4
				var _1300_cf56 _dafny.Set = _1299___mcc_h4
				_ = _1300_cf56
				var _1301_cf55 _dafny.Int = _1298___mcc_h3
				_ = _1301_cf55
				var _1302_cf54 _dafny.MultiSet = _1297___mcc_h2
				_ = _1302_cf54
				var _1303_cf53 _dafny.Int = _1296___mcc_h1
				_ = _1303_cf53
				var _1304_cf52 bool = _1295___mcc_h0
				_ = _1304_cf52
				_1304_cf52 = !(true)
				var _1305_v2 _dafny.Sequence
				_ = _1305_v2
				_1305_v2 = _dafny.SeqOf(_1304_cf52, false)
				var _1306_v3 D4
				_ = _1306_v3
				_1306_v3 = Companion_D4_.Create_DC13_((_this).F6(), _1304_cf52, _1293_i1, _1305_v2)
				r0 = ((func() D4 {
					if p1 {
						return _1306_v3
					}
					return _1306_v3
				})()).Dtor_cf15()
				var _1307_v4 _dafny.Array
				_ = _1307_v4
				var _nw203 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw203
				_1307_v4 = _nw203
				_1307_v4 = _1307_v4
				r2 = (_dafny.Zero).Minus(_1293_i1)
			} else {
				var _1308___mcc_h5 _dafny.Array = _source22.Get_().(D16_DC32).Cf51
				_ = _1308___mcc_h5
				var _1309_cf51 _dafny.Array = _1308___mcc_h5
				_ = _1309_cf51
				var _1310_v5 _dafny.Array
				_ = _1310_v5
				var _nw204 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
				_ = _nw204
				_1310_v5 = _nw204
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_1310_v5), 0))
				_ = _index192
				(_1310_v5).ArraySet1((_this).F4(), (_index192).Int())
				var _1311_v6 D0
				_ = _1311_v6
				_1311_v6 = Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1293_i1)).Cardinality())
				var _1312_v7 _dafny.Set
				_ = _1312_v7
				_1312_v7 = _dafny.SetOf(_1311_v6)
				var _1313_v8 _dafny.MultiSet
				_ = _1313_v8
				_1313_v8 = _dafny.MultiSetOf((_this).F4())
				var _1314_v10 _dafny.Set
				_ = _1314_v10
				_1314_v10 = _dafny.SetOf(_1312_v7, _1312_v7, func() _dafny.Set {
					var _coll45 = _dafny.NewBuilder()
					_ = _coll45
					for _iter52 := _dafny.Iterate((_dafny.SeqOf(Companion_Default___.Fm38((((_1313_v8).Update((_this).F4(), Companion_Default___.Abs(_1293_i1))).Update((_this).F4(), Companion_Default___.Abs(_1293_i1))).Cardinality(), _1293_i1, p1, (_this).F6(), globalState))).Elements()); ; {
						_compr_45, _ok52 := _iter52()
						if !_ok52 {
							break
						}
						var _1315_v9 D0
						_1315_v9 = interface{}(_compr_45).(D0)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_Default___.Fm38((((_1313_v8).Update((_this).F4(), Companion_Default___.Abs(_1293_i1))).Update((_this).F4(), Companion_Default___.Abs(_1293_i1))).Cardinality(), _1293_i1, p1, (_this).F6(), globalState)), _1315_v9) {
							_coll45.Add(_1315_v9)
						}
					}
					return _coll45.ToSet()
				}())
				var _1316_v11 _dafny.MultiSet
				_ = _1316_v11
				_1316_v11 = _dafny.MultiSetOf(((_dafny.SetOf(_1312_v7)).Difference(_1314_v10)).Cardinality())
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_1310_v5), 0))
				_ = _index193
				var _rhs179 bool = !((_this).F4()) || ((_this).F4())
				_ = _rhs179
				var _rhs180 _dafny.MultiSet = _1316_v11
				_ = _rhs180
				var _rhs181 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("skjggxj")).Cardinality()))).Cardinality())
				_ = _rhs181
				var _lhs115 _dafny.Array = _1310_v5
				_ = _lhs115
				var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_1310_v5), 0))
				_ = _lhs116
				(_lhs115).ArraySet1(_rhs179, (_lhs116).Int())
				_1316_v11 = _rhs180
				r1 = _rhs181
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_1310_v5), 0))
				_ = _index194
				(_1310_v5).ArraySet1(((_1310_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_1310_v5), 0))).Int()).(bool)) && (((_1310_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_1310_v5), 0))).Int()).(bool)) == (false)), (_index194).Int())
				var _1317_v12 _dafny.Sequence
				_ = _1317_v12
				_1317_v12 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F6()), (_this).F6())
				r2 = _dafny.IntOfUint32((_1317_v12).Cardinality())
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_1310_v5), 0))
				_ = _index195
				(_1310_v5).ArraySet1(p1, (_index195).Int())
			}
			var _1318_v13 _dafny.Array
			_ = _1318_v13
			var _nw205 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw205
			_1318_v13 = _nw205
			var _1319_v14 D11
			_ = _1319_v14
			_1319_v14 = Companion_D11_.Create_DC26_((_this).F4())
			var _rhs182 _dafny.Array = _1318_v13
			_ = _rhs182
			var _rhs183 D11 = _1319_v14
			_ = _rhs183
			_1318_v13 = _rhs182
			_1319_v14 = _rhs183
			var _1320_v15 _dafny.Map
			_ = _1320_v15
			_1320_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p1)
			var _1321_v16 *C1
			_ = _1321_v16
			var _nw206 *C1 = New_C1_()
			_ = _nw206
			_nw206.Ctor__(((_1320_v15).Cardinality()).Times(_1293_i1), _1293_i1, _this.F5(), _1293_i1, !((_this).F4()) || (!(false)))
			_1321_v16 = _nw206
			var _1322_v17 _dafny.CodePoint
			_ = _1322_v17
			_1322_v17 = _dafny.CodePoint('b')
			r3 = Companion_Default___.Fm0(_1322_v17, globalState)
		}
		if (_this).F4() {
			var _1323_v18 D3
			_ = _1323_v18
			_1323_v18 = Companion_D3_.Create_DC10_(((_this).F6()).Times((_this).F6()))
			var _source23 D3 = _1323_v18
			_ = _source23
			if _source23.Is_DC9() {
				var _1324___mcc_h6 bool = _source23.Get_().(D3_DC9).Cf8
				_ = _1324___mcc_h6
				var _1325___mcc_h7 _dafny.Int = _source23.Get_().(D3_DC9).Cf9
				_ = _1325___mcc_h7
				var _1326___mcc_h8 bool = _source23.Get_().(D3_DC9).Cf10
				_ = _1326___mcc_h8
				var _1327_cf10 bool = _1326___mcc_h8
				_ = _1327_cf10
				var _1328_cf9 _dafny.Int = _1325___mcc_h7
				_ = _1328_cf9
				var _1329_cf8 bool = _1324___mcc_h6
				_ = _1329_cf8
				r2 = (_this).F6()
				var _1330_v19 _dafny.Sequence
				_ = _1330_v19
				_1330_v19 = _dafny.UnicodeSeqOfUtf8Bytes("uig")
				var _1331_v20 _dafny.Map
				_ = _1331_v20
				_1331_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1330_v19)
				var _1332_v21 D0
				_ = _1332_v21
				_1332_v21 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(767))
				var _1333_v22 _dafny.CodePoint
				_ = _1333_v22
				_1333_v22 = _dafny.CodePoint('t')
				var _1334_v23 _dafny.Set
				_ = _1334_v23
				_1334_v23 = _dafny.SetOf(_1330_v19)
				var _1335_v24 bool
				_ = _1335_v24
				var _1336_v25 _dafny.Int
				_ = _1336_v25
				var _out43 bool
				_ = _out43
				var _out44 _dafny.Int
				_ = _out44
				_out43, _out44 = Companion_Default___.M0(_1331_v20, _1329_cf8, (_this).Fm6(Companion_D0_.Create_DC0_((p0).Cardinality()), _1332_v21, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1328_cf9, _1333_v22), (_1334_v23).Cardinality(), globalState), _1327_cf10, globalState)
				_1335_v24 = _out43
				_1336_v25 = _out44
				var _1337_v26 _dafny.Array
				_ = _1337_v26
				var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
				_ = _nw207
				_1337_v26 = _nw207
				var _1338_v27 _dafny.Map
				_ = _1338_v27
				_1338_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1323_v18, Companion_Default___.Fm58(globalState))
				var _1339_v28 _dafny.Sequence
				_ = _1339_v28
				_1339_v28 = _dafny.SeqOf((_this).F6())
				var _1340_v29 _dafny.Sequence
				_ = _1340_v29
				_1340_v29 = _dafny.SeqOf((_this).Fm8(_1336_v25, false, _1330_v19, _1339_v28, globalState))
				var _1341_v30 _dafny.Sequence
				_ = _1341_v30
				_1341_v30 = _dafny.SeqOf(_1340_v29)
				var _1342_v31 D12
				_ = _1342_v31
				_1342_v31 = Companion_D12_.Create_DC27_(_dafny.Companion_Sequence_.Update(_1341_v30, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1340_v29).Cardinality()), _dafny.IntOfUint32((_1341_v30).Cardinality()))).Uint32(), _1340_v29))
				var _1343_v32 _dafny.Map
				_ = _1343_v32
				_1343_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _1342_v31)
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_1337_v26), 0))
				_ = _index196
				(_1337_v26).ArraySet1((func() _dafny.Map {
					if (_1338_v27).Contains(_1323_v18) {
						return (_1338_v27).Get(_1323_v18).(_dafny.Map)
					}
					return _1343_v32
				})(), (_index196).Int())
				var _1344_v33 D20
				_ = _1344_v33
				_1344_v33 = Companion_D20_.Create_DC40_(_1343_v32)
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_1337_v26), 0))
				_ = _index197
				(_1337_v26).ArraySet1((_1343_v32).Merge((_1343_v32).Merge((Companion_D20_.Create_DC40_((_1344_v33).Dtor_cf66())).Dtor_cf66())), (_index197).Int())
				var _1345_v34 _dafny.Map
				_ = _1345_v34
				_1345_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1336_v25, _1327_cf10)
				var _1346_v35 *C1
				_ = _1346_v35
				var _nw208 *C1 = New_C1_()
				_ = _nw208
				_nw208.Ctor__(((_1345_v34).Cardinality()).Plus(_1328_cf9), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1330_v19).Cardinality()), _1336_v25), _this.F5(), (_this).F6(), _1335_v24)
				_1346_v35 = _nw208
			} else if _source23.Is_DC10() {
				var _1347___mcc_h9 _dafny.Int = _source23.Get_().(D3_DC10).Cf11
				_ = _1347___mcc_h9
				var _1348_cf11 _dafny.Int = _1347___mcc_h9
				_ = _1348_cf11
				r1 = _dafny.IntOfInt64(491)
				var _1349_v36 _dafny.Sequence
				_ = _1349_v36
				_1349_v36 = _dafny.UnicodeSeqOfUtf8Bytes("fvlmlhsge")
				r2 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1349_v36, _1349_v36)).Cardinality())
				var _1350_v37 *C1
				_ = _1350_v37
				var _nw209 *C1 = New_C1_()
				_ = _nw209
				_nw209.Ctor__((_this).F6(), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(753), (_this).F6()), _this.F5(), Companion_Default___.SafeModuloInt((_this).F6(), (_this).F6()), (_this).F4())
				_1350_v37 = _nw209
				_1291_v0 = _1291_v0
			} else if _source23.Is_DC8() {
				var _1351___mcc_h10 _dafny.Sequence = _source23.Get_().(D3_DC8).Cf7
				_ = _1351___mcc_h10
				var _1352_cf7 _dafny.Sequence = _1351___mcc_h10
				_ = _1352_cf7
				var _1353_v38 _dafny.CodePoint
				_ = _1353_v38
				_1353_v38 = _dafny.CodePoint('e')
				var _1354_v39 _dafny.Map
				_ = _1354_v39
				_1354_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1353_v38, false)
				r0 = !(p1) || ((func() bool {
					if (_1354_v39).Contains(_1353_v38) {
						return (_1354_v39).Get(_1353_v38).(bool)
					}
					return !((_this).F4())
				})())
				var _1355_v40 _dafny.Sequence
				_ = _1355_v40
				_1355_v40 = _dafny.UnicodeSeqOfUtf8Bytes("honftv")
				r2 = ((_this).F6()).Minus((_this).Fm9((_this).F4(), (_this).F4(), _dafny.IntOfUint32((_1355_v40).Cardinality()), globalState))
				var _1356_v41 _dafny.Array
				_ = _1356_v41
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_27
				var _nw210 _dafny.Array
				_ = _nw210
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw210 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) bool = func(_1357_i2 _dafny.Int) bool {
						return true
					}
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw210 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw210).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw210).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_1356_v41 = _nw210
				var _1358_v42 _dafny.Map
				_ = _1358_v42
				_1358_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _1356_v41)
				var _1359_v43 D19
				_ = _1359_v43
				_1359_v43 = Companion_D19_.Create_DC37_(_1358_v42)
				var _1360_v44 _dafny.Map
				_ = _1360_v44
				_1360_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F4()) || (true), _1359_v43)
				_1360_v44 = (_1360_v44).Update((_this).F4(), _1359_v43)
				var _1361_v45 *C2
				_ = _1361_v45
				var _nw211 *C2 = New_C2_()
				_ = _nw211
				_nw211.Ctor__(_1355_v40, (_this).F4())
				_1361_v45 = _nw211
				var _1362_v46 _dafny.Sequence
				_ = _1362_v46
				_1362_v46 = _dafny.SeqOf(_1361_v45, _1361_v45, _1361_v45)
				var _1363_v47 _dafny.Sequence
				_ = _1363_v47
				_1363_v47 = _dafny.SeqOf(_1362_v46, _dafny.SeqOf(_1361_v45))
				r3 = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_1363_v47, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1363_v47).Cardinality()))).Uint32(), _1362_v46), _1363_v47)
			} else {
				var _1364___mcc_h11 D3 = _source23.Get_().(D3_DC11).Cf12
				_ = _1364___mcc_h11
				var _1365_cf12 D3 = _1364___mcc_h11
				_ = _1365_cf12
				var _1366_v48 _dafny.Sequence
				_ = _1366_v48
				_1366_v48 = _dafny.SeqOf((_this).F6(), (_this).F6())
				_1366_v48 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(904))).Uint32(), func(coer84 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}((func(_1367_p1 bool) func(_dafny.Int) _dafny.Int {
					return func(_1368_i3 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1367_p1), _dafny.SeqOf((_this).F4(), (_this).F4()))).Cardinality())
					}
				})(p1)))
				var _1369_v49 _dafny.Sequence
				_ = _1369_v49
				_1369_v49 = _dafny.UnicodeSeqOfUtf8Bytes("bwcahcko")
				var _1370_v50 _dafny.CodePoint
				_ = _1370_v50
				_1370_v50 = _dafny.CodePoint('h')
				var _1371_v51 D9
				_ = _1371_v51
				_1371_v51 = Companion_D9_.Create_DC19_(_1369_v49)
				_1369_v49 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_this).F4() {
						return _dafny.UnicodeSeqOfUtf8Bytes("ylnncpc")
					}
					return _1369_v49
				})(), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_this).F4() {
						return _dafny.UnicodeSeqOfUtf8Bytes("ylnncpc")
					}
					return _1369_v49
				})()).Cardinality()))).Uint32(), _1370_v50), (_1371_v51).Dtor_cf25())
				var _1372_v52 D12
				_ = _1372_v52
				_1372_v52 = Companion_D12_.Create_DC28_((_this).F6(), (_dafny.Zero).Minus((_this).F6()), (_this).F4(), p1, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_this).F4() {
						return _dafny.UnicodeSeqOfUtf8Bytes("iunuhhphg")
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(651))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg85 _dafny.Int) interface{} {
							return coer85(arg85)
						}
					}((func(_1373_v50 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1374_i4 _dafny.Int) _dafny.CodePoint {
							return _1373_v50
						}
					})(_1370_v50)))
				})()).Cardinality()))
				var _rhs184 D12 = _1372_v52
				_ = _rhs184
				var _rhs185 bool = (true) || ((_this).F4())
				_ = _rhs185
				var _rhs186 _dafny.Sequence = _1369_v49
				_ = _rhs186
				_1372_v52 = _rhs184
				r0 = _rhs185
				_1369_v49 = _rhs186
				r3 = p1
			}
			var _1375_v53 _dafny.Sequence
			_ = _1375_v53
			_1375_v53 = _dafny.UnicodeSeqOfUtf8Bytes("ukokag")
			r1 = _dafny.IntOfUint32((Companion_Default___.Fm37(p1, !(p1), _dafny.IntOfUint32((_1375_v53).Cardinality()), (_this).Fm8((_this).F6(), p1, _1375_v53, _dafny.SeqOf(_dafny.IntOfUint32((_1375_v53).Cardinality()), (_this).F6()), globalState), globalState)).Cardinality())
			r2 = (_dafny.Zero).Minus((_this).F6())
			var _1376_v54 _dafny.Set
			_ = _1376_v54
			_1376_v54 = _dafny.SetOf((_this).F4(), p1)
			var _1377_v55 _dafny.Sequence
			_ = _1377_v55
			_1377_v55 = _dafny.SeqOf((_1376_v54).IsSubsetOf(_1376_v54))
			r0 = (_1377_v55).Select((Companion_Default___.SafeIndex(((_this).F6()).Minus(_dafny.IntOfInt64(779)), _dafny.IntOfUint32((_1377_v55).Cardinality()))).Uint32()).(bool)
			var _1378_v56 _dafny.CodePoint
			_ = _1378_v56
			_1378_v56 = _dafny.CodePoint('i')
			_1378_v56 = _1378_v56
		} else {
			r1 = (_this).F6()
			var _1379_v57 _dafny.Sequence
			_ = _1379_v57
			_1379_v57 = _dafny.UnicodeSeqOfUtf8Bytes("kcdy")
			var _1380_v58 _dafny.Map
			_ = _1380_v58
			_1380_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _1379_v57)
			var _1381_v59 _dafny.Map
			_ = _1381_v59
			_1381_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1379_v57, _1380_v58)
			var _1382_v60 _dafny.Map
			_ = _1382_v60
			_1382_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1380_v58)
			_1381_v59 = (_1381_v59).Update(_1379_v57, ((func() _dafny.Map {
				if (_1382_v60).Contains(_dafny.IntOfUint32((_1379_v57).Cardinality())) {
					return (_1382_v60).Get(_dafny.IntOfUint32((_1379_v57).Cardinality())).(_dafny.Map)
				}
				return _1380_v58
			})()).Merge(_1380_v58))
			var _1383_v61 _dafny.Sequence
			_ = _1383_v61
			_1383_v61 = _dafny.SeqOf((_this).F6())
			r3 = (_this).Fm8((_this).F6(), (p1) && ((_this).F4()), _1379_v57, _1383_v61, globalState)
			var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_1291_v0), 0))
			_ = _index198
			(_1291_v0).ArraySet1((_this).F6(), (_index198).Int())
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_1291_v0), 0))
			_ = _index199
			(_1291_v0).ArraySet1((_this).F6(), (_index199).Int())
			var _1384_v62 _dafny.Array
			_ = _1384_v62
			var _nw212 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(14))
			_ = _nw212
			_1384_v62 = _nw212
			var _1385_v63 _dafny.Set
			_ = _1385_v63
			_1385_v63 = _dafny.SetOf((_this).F6(), _dafny.IntOfUint32((_1379_v57).Cardinality()))
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1384_v62), 0))
			_ = _index200
			(_1384_v62).ArraySet1(_1385_v63, (_index200).Int())
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1384_v62), 0))
			_ = _index201
			(_1384_v62).ArraySet1(_1385_v63, (_index201).Int())
		}
		var _1386_v64 _dafny.Sequence
		_ = _1386_v64
		_1386_v64 = _dafny.UnicodeSeqOfUtf8Bytes("iktm")
		var _hi7 _dafny.Int = (_this).F6()
		_ = _hi7
		for _1387_i5 := ((_this).F6()).Minus(_dafny.IntOfUint32((_1386_v64).Cardinality())); _1387_i5.Cmp(_hi7) < 0; _1387_i5 = _1387_i5.Plus(_dafny.One) {
			var _1388_v65 _dafny.Sequence
			_ = _1388_v65
			_1388_v65 = _dafny.SeqOf((_this).F6())
			r2 = (_1388_v65).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1388_v65).Cardinality()))).Uint32()).(_dafny.Int)
			var _1389_v66 _dafny.MultiSet
			_ = _1389_v66
			_1389_v66 = _dafny.MultiSetOf((_this).F4(), !((_this).F4()))
			var _pat_let_tv23 = _1386_v64
			_ = _pat_let_tv23
			var _source24 D12 = func(_pat_let26_0 D12) D12 {
				return func(_1392_dt__update__tmp_h1 D12) D12 {
					return func(_pat_let29_0 _dafny.Int) D12 {
						return func(_1393_dt__update_hcf47_h0 _dafny.Int) D12 {
							return Companion_D12_.Create_DC28_((_1392_dt__update__tmp_h1).Dtor_cf43(), (_1392_dt__update__tmp_h1).Dtor_cf44(), (_1392_dt__update__tmp_h1).Dtor_cf45(), (_1392_dt__update__tmp_h1).Dtor_cf46(), _1393_dt__update_hcf47_h0)
						}(_pat_let29_0)
					}(_dafny.IntOfUint32((_pat_let_tv23).Cardinality()))
				}(_pat_let26_0)
			}(func(_pat_let27_0 D12) D12 {
				return func(_1390_dt__update__tmp_h0 D12) D12 {
					return func(_pat_let28_0 _dafny.Int) D12 {
						return func(_1391_dt__update_hcf43_h0 _dafny.Int) D12 {
							return Companion_D12_.Create_DC28_(_1391_dt__update_hcf43_h0, (_1390_dt__update__tmp_h0).Dtor_cf44(), (_1390_dt__update__tmp_h0).Dtor_cf45(), (_1390_dt__update__tmp_h0).Dtor_cf46(), (_1390_dt__update__tmp_h0).Dtor_cf47())
						}(_pat_let28_0)
					}((_this).F6())
				}(_pat_let27_0)
			}(Companion_D12_.Create_DC28_((_this).F6(), ((p0).Update(_1387_i5, Companion_Default___.Abs((_1389_v66).Cardinality()))).Cardinality(), p1, true, _1387_i5)))
			_ = _source24
			if _source24.Is_DC28() {
				var _1394___mcc_h12 _dafny.Int = _source24.Get_().(D12_DC28).Cf43
				_ = _1394___mcc_h12
				var _1395___mcc_h13 _dafny.Int = _source24.Get_().(D12_DC28).Cf44
				_ = _1395___mcc_h13
				var _1396___mcc_h14 bool = _source24.Get_().(D12_DC28).Cf45
				_ = _1396___mcc_h14
				var _1397___mcc_h15 bool = _source24.Get_().(D12_DC28).Cf46
				_ = _1397___mcc_h15
				var _1398___mcc_h16 _dafny.Int = _source24.Get_().(D12_DC28).Cf47
				_ = _1398___mcc_h16
				var _1399_cf47 _dafny.Int = _1398___mcc_h16
				_ = _1399_cf47
				var _1400_cf46 bool = _1397___mcc_h15
				_ = _1400_cf46
				var _1401_cf45 bool = _1396___mcc_h14
				_ = _1401_cf45
				var _1402_cf44 _dafny.Int = _1395___mcc_h13
				_ = _1402_cf44
				var _1403_cf43 _dafny.Int = _1394___mcc_h12
				_ = _1403_cf43
				var _1404_v67 D0
				_ = _1404_v67
				_1404_v67 = Companion_D0_.Create_DC0_((_dafny.IntOfInt64(879)).Minus(_1399_cf47))
				_1404_v67 = _1404_v67
				_1386_v64 = _1386_v64
				var _1405_v68 _dafny.Set
				_ = _1405_v68
				_1405_v68 = _dafny.SetOf((_this).F6(), _1399_cf47)
				_1400_cf46 = (_1387_i5).Cmp((_dafny.Zero).Minus(((_1405_v68).Cardinality()).Times((_1388_v65).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1405_v68).Cardinality()), _dafny.IntOfUint32((_1388_v65).Cardinality()))).Uint32()).(_dafny.Int)))) == 0
				var _1406_v69 _dafny.Map
				_ = _1406_v69
				_1406_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1399_cf47, p1)
				var _1407_v71 _dafny.MultiSet
				_ = _1407_v71
				_1407_v71 = _dafny.MultiSetOf((func() _dafny.Map {
					if _1401_cf45 {
						return _1406_v69
					}
					return _1406_v69
				})(), func() _dafny.Map {
					var _coll46 = _dafny.NewMapBuilder()
					_ = _coll46
					for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(416), _dafny.IntOfInt64(-448))); ; {
						_compr_46, _ok53 := _iter53()
						if !_ok53 {
							break
						}
						var _1408_v70 _dafny.Int
						_1408_v70 = interface{}(_compr_46).(_dafny.Int)
						if ((_dafny.IntOfInt64(416)).Cmp(_1408_v70) <= 0) && ((_1408_v70).Cmp(_dafny.IntOfInt64(-448)) < 0) {
							_coll46.Add(Companion_Default___.SafeModuloInt(_1408_v70, (_this).F6()), _1400_cf46)
						}
					}
					return _coll46.ToMap()
				}(), _1406_v69)
				_1407_v71 = _1407_v71
			} else {
				var _1409___mcc_h17 _dafny.Sequence = _source24.Get_().(D12_DC27).Cf42
				_ = _1409___mcc_h17
				var _1410_cf42 _dafny.Sequence = _1409___mcc_h17
				_ = _1410_cf42
				r3 = !((_dafny.IntOfUint32((_1386_v64).Cardinality())).Cmp((_this).F6()) <= 0) || ((_this).F4())
				r1 = _1387_i5
				var _1411_v72 _dafny.Array
				_ = _1411_v72
				var _nwElement0_48 bool = (func() bool {
					if !(!(p1)) {
						return !(p1)
					}
					return (_this).F4()
				})()
				_ = _nwElement0_48
				var _nw213 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(3))
				_ = _nw213
				(_nw213).ArraySet1(_nwElement0_48, 0)
				(_nw213).ArraySet1((_this).F4(), 1)
				(_nw213).ArraySet1((_this).F4(), 2)
				_1411_v72 = _nw213
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_1411_v72), 0))
				_ = _index202
				(_1411_v72).ArraySet1(p1, (_index202).Int())
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_1411_v72), 0))
				_ = _index203
				var _rhs187 bool = p1
				_ = _rhs187
				var _rhs188 bool = p1
				_ = _rhs188
				var _lhs117 _dafny.Array = _1411_v72
				_ = _lhs117
				var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_1411_v72), 0))
				_ = _lhs118
				r0 = _rhs187
				(_lhs117).ArraySet1(_rhs188, (_lhs118).Int())
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_1411_v72), 0))
				_ = _index204
				(_1411_v72).ArraySet1((_1411_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_1411_v72), 0))).Int()).(bool), (_index204).Int())
			}
			var _1412_v73 _dafny.Map
			_ = _1412_v73
			_1412_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _dafny.IntOfUint32((_1386_v64).Cardinality()))
			var _1413_v74 D10
			_ = _1413_v74
			_1413_v74 = Companion_D10_.Create_DC21_(_1412_v73)
			var _1414_v75 _dafny.Map
			_ = _1414_v75
			_1414_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1413_v74, _dafny.IntOfUint32((_1386_v64).Cardinality()))
			_1414_v75 = Companion_Default___.Fm59(globalState)
			r1 = (_dafny.Zero).Minus((_this).F6())
		}
		var _1415_v76 _dafny.Sequence
		_ = _1415_v76
		_1415_v76 = _dafny.SeqOf((_this).F6())
		_1415_v76 = (func() _dafny.Sequence {
			if p1 {
				return _dafny.SeqOf(_dafny.IntOfInt64(-887), (_this).F6(), (_this).F6(), (_this).F6(), (_this).F6())
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(899))).Uint32(), func(coer86 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg86 _dafny.Int) interface{} {
					return coer86(arg86)
				}
			}(func(_1416_i6 _dafny.Int) _dafny.Int {
				return (_this).F6()
			}))
		})()
		var _1417_v77 _dafny.Map
		_ = _1417_v77
		_1417_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), ((_this).F6()).Minus((_this).F6()))
		_1417_v77 = (Companion_Default___.Fm44(globalState)).Merge((_1417_v77).Merge(_1417_v77))
		r0 = true
		r1 = (_this).F6()
		r2 = (_this).F6()
		r3 = !((_this).F4())
		return r0, r1, r2, r3
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f5 _dafny.Array
	_f6 _dafny.Int
	_f4 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.Zero
	_this._f4 = false
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

func (_this *C4) F5() _dafny.Array {
	return _this._f5
}
func (_this *C4) F5_set_(value _dafny.Array) {
	_this._f5 = value
}
func (_this *C4) F6() _dafny.Int {
	return _this._f6
}
func (_this *C4) F4() bool {
	return _this._f4
}
func (_this *C4) Ctor__(f5 _dafny.Array, f6 _dafny.Int, f4 bool) {
	{
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f4 = f4
	}
}
func (_this *C4) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return !((_this).F4())
	}
}
func (_this *C4) Fm9(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F6()
	}
}
func (_this *C4) Fm6(p0 D0, p1 D0, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return _dafny.MultiSetOf(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(117), _dafny.IntOfInt64(842)), (func() _dafny.Int {
			if (_this).F4() {
				return (_this).F6()
			}
			return (_this).F6()
		})(), (_this).F6(), Companion_Default___.SafeDivisionInt((_this).F6(), (_this).F6()), (_this).F6())
	}
}
func (_this *C4) Fm7(globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((func(_source25 D9) _dafny.Sequence {
			if _source25.Is_DC20() {
				var _1418___mcc_h0 bool = _source25.Get_().(D9_DC20).Cf26
				_ = _1418___mcc_h0
				var _1419___mcc_h1 _dafny.CodePoint = _source25.Get_().(D9_DC20).Cf27
				_ = _1419___mcc_h1
				var _1420___mcc_h2 _dafny.Int = _source25.Get_().(D9_DC20).Cf28
				_ = _1420___mcc_h2
				var _1421_cf28 _dafny.Int = _1420___mcc_h2
				_ = _1421_cf28
				var _1422_cf27 _dafny.CodePoint = _1419___mcc_h1
				_ = _1422_cf27
				var _1423_cf26 bool = _1418___mcc_h0
				_ = _1423_cf26
				return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uaimtxsem"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(744))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg87 _dafny.Int) interface{} {
						return coer87(arg87)
					}
				}((func(_1424_cf27 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1425_i0 _dafny.Int) _dafny.CodePoint {
						return _1424_cf27
					}
				})(_1422_cf27))))
			} else {
				var _1426___mcc_h3 _dafny.Sequence = _source25.Get_().(D9_DC19).Cf25
				_ = _1426___mcc_h3
				var _1427_cf25 _dafny.Sequence = _1426___mcc_h3
				_ = _1427_cf25
				return _1427_cf25
			}
		}(Companion_D9_.Create_DC19_(_dafny.UnicodeSeqOfUtf8Bytes("sxfde")))).Cardinality())
	}
}
func (_this *C4) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _1428_v0 _dafny.Sequence
		_ = _1428_v0
		_1428_v0 = _dafny.UnicodeSeqOfUtf8Bytes("tv")
		var _1429_v1 _dafny.MultiSet
		_ = _1429_v1
		_1429_v1 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F4())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_1428_v0)).Cardinality()))
		var _1430_v2 bool
		_ = _1430_v2
		var _1431_v3 _dafny.Int
		_ = _1431_v3
		var _out45 bool
		_ = _out45
		var _out46 _dafny.Int
		_ = _out46
		_out45, _out46 = Companion_Default___.M0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1429_v1, _1428_v0), !((func() bool {
			if (_this).F4() {
				return (_this).F4()
			}
			return (_this).F4()
		})()), _1429_v1, (_this).F4(), globalState)
		_1430_v2 = _out45
		_1431_v3 = _out46
		var _1432_v4 _dafny.Array
		_ = _1432_v4
		var _len0_28 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_28
		var _nw214 _dafny.Array
		_ = _nw214
		if _len0_28.Cmp(_dafny.Zero) == 0 {
			_nw214 = _dafny.NewArray(_len0_28)
		} else {
			var _init28 func(_dafny.Int) _dafny.CodePoint = func(_1433_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			}
			_ = _init28
			var _element0_28 = _init28(_dafny.Zero)
			_ = _element0_28
			_nw214 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
			(_nw214).ArraySet1CodePoint(_element0_28, 0)
			var _nativeLen0_28 = (_len0_28).Int()
			_ = _nativeLen0_28
			for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
				(_nw214).ArraySet1CodePoint(_init28(_dafny.IntOf(_i0_28)), _i0_28)
			}
		}
		_1432_v4 = _nw214
		var _1434_v5 _dafny.CodePoint
		_ = _1434_v5
		_1434_v5 = _dafny.CodePoint('i')
		var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1432_v4), 0))
		_ = _index205
		(_1432_v4).ArraySet1CodePoint(_1434_v5, (_index205).Int())
		var _1435_v6 _dafny.Sequence
		_ = _1435_v6
		_1435_v6 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt(_1431_v3, (_this).F6()))
		var _1436_v7 _dafny.Map
		_ = _1436_v7
		_1436_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1430_v2, (_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm30(_1435_v6, true, globalState)).Cardinality())))
		var _1437_v8 _dafny.Map
		_ = _1437_v8
		_1437_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1430_v2, (_this).F4())
		var _1438_v9 _dafny.Map
		_ = _1438_v9
		_1438_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1437_v8, (_dafny.SetOf((_this).F4(), _1430_v2)).Cardinality())
		var _1439_v11 _dafny.Array
		_ = _1439_v11
		var _nwElement0_49 _dafny.Int = _dafny.IntOfInt64(-811)
		_ = _nwElement0_49
		var _nw215 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(22))
		_ = _nw215
		(_nw215).ArraySet1(_nwElement0_49, 0)
		(_nw215).ArraySet1(p0, 1)
		(_nw215).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-476))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg88 _dafny.Int) interface{} {
				return coer88(arg88)
			}
		}((func(_1440_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1441_i1 _dafny.Int) _dafny.CodePoint {
				return _1440_v5
			}
		})(_1434_v5)))).Cardinality()), 2)
		(_nw215).ArraySet1((_this).F6(), 3)
		(_nw215).ArraySet1((_this).F6(), 4)
		(_nw215).ArraySet1((_this).F6(), 5)
		(_nw215).ArraySet1((_1436_v7).Cardinality(), 6)
		(_nw215).ArraySet1((_this).F6(), 7)
		(_nw215).ArraySet1((_this).F6(), 8)
		(_nw215).ArraySet1(_1431_v3, 9)
		(_nw215).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p0, p0, (_dafny.Zero).Minus(_1431_v3), (_this).F6())).Cardinality()), 10)
		(_nw215).ArraySet1((_this).F6(), 11)
		(_nw215).ArraySet1(p0, 12)
		(_nw215).ArraySet1((_this).F6(), 13)
		(_nw215).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ejh")).Cardinality()), 14)
		(_nw215).ArraySet1(Companion_Default___.Fm3(_1431_v3, (_this).F4(), globalState), 15)
		(_nw215).ArraySet1(_1431_v3, 16)
		(_nw215).ArraySet1(_1431_v3, 17)
		(_nw215).ArraySet1((func() _dafny.Int {
			if (_1438_v9).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1430_v2, (_this).F4())) {
				return (_1438_v9).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1430_v2, (_this).F4())).(_dafny.Int)
			}
			return (func() _dafny.Map {
				var _coll47 = _dafny.NewMapBuilder()
				_ = _coll47
				for _iter54 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(374), _dafny.IntOfInt64(306))); ; {
					_compr_47, _ok54 := _iter54()
					if !_ok54 {
						break
					}
					var _1442_v10 _dafny.Int
					_1442_v10 = interface{}(_compr_47).(_dafny.Int)
					if ((_dafny.IntOfInt64(374)).Cmp(_1442_v10) <= 0) && ((_1442_v10).Cmp(_dafny.IntOfInt64(306)) < 0) {
						_coll47.Add((_1442_v10).Minus((_this).F6()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_1435_v6)).Cardinality(), (_this).F4())).Cardinality())
					}
				}
				return _coll47.ToMap()
			}()).Cardinality()
		})(), 18)
		(_nw215).ArraySet1(p0, 19)
		(_nw215).ArraySet1(_1431_v3, 20)
		(_nw215).ArraySet1((_this).F6(), 21)
		_1439_v11 = _nw215
		var _1443_v12 _dafny.Sequence
		_ = _1443_v12
		_1443_v12 = _dafny.SeqOf(_1439_v11)
		var _1444_v13 _dafny.MultiSet
		_ = _1444_v13
		_1444_v13 = _dafny.MultiSetOf(_1439_v11, (_1443_v12).Select((Companion_Default___.SafeIndex(_1431_v3, _dafny.IntOfUint32((_1443_v12).Cardinality()))).Uint32()).(_dafny.Array))
		var _1445_v14 _dafny.MultiSet
		_ = _1445_v14
		_1445_v14 = _dafny.MultiSetOf(_1444_v13, (_1444_v13).Difference((_1444_v13).Update(_1439_v11, Companion_Default___.Abs(p0))))
		var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1432_v4), 0))
		_ = _index206
		var _rhs189 _dafny.CodePoint = _1434_v5
		_ = _rhs189
		var _rhs190 _dafny.Int = (_this).F6()
		_ = _rhs190
		var _rhs191 _dafny.Int = (func() _dafny.Int {
			if (_1445_v14).Contains((_1444_v13).Union(_1444_v13)) {
				return (_1445_v14).Multiplicity((_1444_v13).Union(_1444_v13))
			}
			return (func() _dafny.Int {
				if _1430_v2 {
					return _1431_v3
				}
				return p0
			})()
		})()
		_ = _rhs191
		var _rhs192 _dafny.Sequence = _1435_v6
		_ = _rhs192
		var _lhs119 _dafny.Array = _1432_v4
		_ = _lhs119
		var _lhs120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1432_v4), 0))
		_ = _lhs120
		(_lhs119).ArraySet1CodePoint(_rhs189, (_lhs120).Int())
		_1431_v3 = _rhs190
		_1431_v3 = _rhs191
		_1435_v6 = _rhs192
		var _1446_v15 *C0
		_ = _1446_v15
		var _nw216 *C0 = New_C0_()
		_ = _nw216
		_nw216.Ctor__(Companion_Default___.SafeDivisionInt(_1431_v3, _dafny.IntOfUint32((_1435_v6).Cardinality())))
		_1446_v15 = _nw216
		var _1447_v16 _dafny.Map
		_ = _1447_v16
		_1447_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1434_v5, _1430_v2)
		_1447_v16 = (_1447_v16).Update((_1432_v4).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1432_v4), 0))).Int()), (_this).F4())
		var _hi8 _dafny.Int = (_dafny.Zero).Minus(((_1446_v15).F12()).Minus(_dafny.IntOfInt64(281)))
		_ = _hi8
		for _1448_i2 := _dafny.IntOfInt64(-659); _1448_i2.Cmp(_hi8) < 0; _1448_i2 = _1448_i2.Plus(_dafny.One) {
			var _1449_v17 *C0
			_ = _1449_v17
			var _nw217 *C0 = New_C0_()
			_ = _nw217
			_nw217.Ctor__(_1431_v3)
			_1449_v17 = _nw217
			var _1450_v18 _dafny.Map
			_ = _1450_v18
			_1450_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jl")).Cardinality()), Companion_Default___.Fm31(_1448_i2, _1431_v3, _dafny.IntOfUint32((_1428_v0).Cardinality()), (_this).F4(), globalState))
			var _1451_v19 _dafny.Map
			_ = _1451_v19
			_1451_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1431_v3, (_this).F4())
			var _1452_v20 _dafny.Set
			_ = _1452_v20
			_1452_v20 = _dafny.SetOf((_1429_v1).Cardinality(), (_1436_v7).Cardinality(), (Companion_D0_.Create_DC0_((_this).F6())).Dtor_cf0(), (_1451_v19).Cardinality(), _dafny.IntOfUint32((_1435_v6).Cardinality()))
			var _1453_v23 _dafny.Array
			_ = _1453_v23
			var _nwElement0_50 _dafny.Set = (func() _dafny.Set {
				if (_1450_v18).Contains((_1446_v15).F12()) {
					return (_1450_v18).Get((_1446_v15).F12()).(_dafny.Set)
				}
				return _1452_v20
			})()
			_ = _nwElement0_50
			var _nw218 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(8))
			_ = _nw218
			(_nw218).ArraySet1(_nwElement0_50, 0)
			(_nw218).ArraySet1(func() _dafny.Set {
				var _coll48 = _dafny.NewBuilder()
				_ = _coll48
				for _iter55 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(380), _dafny.IntOfInt64(684))); ; {
					_compr_48, _ok55 := _iter55()
					if !_ok55 {
						break
					}
					var _1454_v21 _dafny.Int
					_1454_v21 = interface{}(_compr_48).(_dafny.Int)
					if ((_dafny.IntOfInt64(380)).Cmp(_1454_v21) <= 0) && ((_1454_v21).Cmp(_dafny.IntOfInt64(684)) < 0) {
						_coll48.Add((_1454_v21).Minus((_1446_v15).F12()))
					}
				}
				return _coll48.ToSet()
			}(), 1)
			(_nw218).ArraySet1(_dafny.SetOf((_this).F6(), (_1446_v15).F12(), (_1449_v17).F12(), (_1429_v1).Cardinality()), 2)
			(_nw218).ArraySet1((Companion_Default___.Fm32(_dafny.IntOfInt64(558), _1430_v2, globalState)), 3)
			(_nw218).ArraySet1(_1452_v20, 4)
			(_nw218).ArraySet1(_dafny.SetOf(_1448_i2), 5)
			(_nw218).ArraySet1(_1452_v20, 6)
			(_nw218).ArraySet1(func() _dafny.Set {
				var _coll49 = _dafny.NewBuilder()
				_ = _coll49
				for _iter56 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(788), _dafny.IntOfInt64(145))); ; {
					_compr_49, _ok56 := _iter56()
					if !_ok56 {
						break
					}
					var _1455_v22 _dafny.Int
					_1455_v22 = interface{}(_compr_49).(_dafny.Int)
					if ((_dafny.IntOfInt64(788)).Cmp(_1455_v22) <= 0) && ((_1455_v22).Cmp(_dafny.IntOfInt64(145)) < 0) {
						_coll49.Add((_1455_v22).Minus((_1449_v17).F12()))
					}
				}
				return _coll49.ToSet()
			}(), 7)
			_1453_v23 = _nw218
			var _1456_v24 _dafny.Set
			_ = _1456_v24
			_1456_v24 = _dafny.SetOf((_1449_v17).Fm17((_this).F4(), globalState), !(!((_this).F4())), _1430_v2, (_this).F4(), !((_this).F4()))
			var _1457_v25 _dafny.Map
			_ = _1457_v25
			_1457_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1456_v24, _1430_v2)
			var _1458_v26 _dafny.Map
			_ = _1458_v26
			_1458_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if _1430_v2 {
					return _1453_v23
				}
				return _1453_v23
			})(), (_1457_v25).Cardinality())
			_1458_v26 = (_1458_v26).Update(_1453_v23, Companion_Default___.SafeModuloInt(p0, (_1446_v15).F12()))
			var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_1439_v11), 0))
			_ = _index207
			(_1439_v11).ArraySet1((_1436_v7).Cardinality(), (_index207).Int())
			var _1459_v27 _dafny.Sequence
			_ = _1459_v27
			_1459_v27 = _dafny.SeqOf((_this).F4(), _1430_v2, _1430_v2, _1430_v2)
			var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((_1439_v11), 0))
			_ = _index208
			(_1439_v11).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1459_v27, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1459_v27).Cardinality()), _dafny.IntOfUint32((_1459_v27).Cardinality()))).Uint32(), (_this).F4()), _1459_v27)).Cardinality()), (_index208).Int())
			var _1460_v28 _dafny.MultiSet
			_ = _1460_v28
			_1460_v28 = _dafny.MultiSetOf((_1446_v15).Fm17((_this).F4(), globalState), (_1446_v15).Fm17(true, globalState), _1430_v2)
			_1460_v28 = _1460_v28
		}
		var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1432_v4), 0))
		_ = _index209
		(_1432_v4).ArraySet1CodePoint((_1432_v4).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1432_v4), 0))).Int()), (_index209).Int())
		r0 = _1436_v7
		return r0
	}
}
func (_this *C4) M2(p0 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Array) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _1461_v0 _dafny.Map
		_ = _1461_v0
		_1461_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus((_this).F6())).Times((_this).F6()), (_this).Fm7(globalState))
		var _1462_v1 _dafny.Array
		_ = _1462_v1
		var _nw219 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
		_ = _nw219
		_1462_v1 = _nw219
		var _1463_v2 _dafny.Map
		_ = _1463_v2
		_1463_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1462_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_dafny.Zero).Minus((_this).F6())))
		_1461_v0 = (func() _dafny.Map {
			if (_1463_v2).Contains(_1462_v1) {
				return (_1463_v2).Get(_1462_v1).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
		})()
		var _1464_v3 _dafny.Sequence
		_ = _1464_v3
		_1464_v3 = _dafny.UnicodeSeqOfUtf8Bytes("tjatmipct")
		var _1465_v4 _dafny.Sequence
		_ = _1465_v4
		_1465_v4 = _dafny.SeqOf(_dafny.IntOfUint32((_1464_v3).Cardinality()), (_this).F6())
		_1465_v4 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(355))).Uint32(), func(coer89 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg89 _dafny.Int) interface{} {
				return coer89(arg89)
			}
		}(func(_1466_i0 _dafny.Int) _dafny.Int {
			return (_this).F6()
		}))
		var _hi9 _dafny.Int = _dafny.IntOfUint32((_1464_v3).Cardinality())
		_ = _hi9
		for _1467_i1 := (_this).F6(); _1467_i1.Cmp(_hi9) < 0; _1467_i1 = _1467_i1.Plus(_dafny.One) {
			var _1468_v5 _dafny.Sequence
			_ = _1468_v5
			_1468_v5 = _dafny.SeqOf(p0)
			var _1469_v6 D4
			_ = _1469_v6
			_1469_v6 = Companion_D4_.Create_DC13_(Companion_Default___.SafeModuloInt((_1461_v0).Cardinality(), _1467_i1), (_this).F4(), _dafny.IntOfInt64(-262), _1468_v5)
			var _source26 D4 = _1469_v6
			_ = _source26
			if _source26.Is_DC13() {
				var _1470___mcc_h0 _dafny.Int = _source26.Get_().(D4_DC13).Cf14
				_ = _1470___mcc_h0
				var _1471___mcc_h1 bool = _source26.Get_().(D4_DC13).Cf15
				_ = _1471___mcc_h1
				var _1472___mcc_h2 _dafny.Int = _source26.Get_().(D4_DC13).Cf16
				_ = _1472___mcc_h2
				var _1473___mcc_h3 _dafny.Sequence = _source26.Get_().(D4_DC13).Cf17
				_ = _1473___mcc_h3
				var _1474_cf17 _dafny.Sequence = _1473___mcc_h3
				_ = _1474_cf17
				var _1475_cf16 _dafny.Int = _1472___mcc_h2
				_ = _1475_cf16
				var _1476_cf15 bool = _1471___mcc_h1
				_ = _1476_cf15
				var _1477_cf14 _dafny.Int = _1470___mcc_h0
				_ = _1477_cf14
				var _1478_v7 _dafny.CodePoint
				_ = _1478_v7
				_1478_v7 = _dafny.CodePoint('y')
				_1475_cf16 = (func() _dafny.Int {
					if !_dafny.Companion_Sequence_.Contains(_1464_v3, _1478_v7) {
						return _1477_cf14
					}
					return _1477_cf14
				})()
				var _1479_v8 *C0
				_ = _1479_v8
				var _nw220 *C0 = New_C0_()
				_ = _nw220
				_nw220.Ctor__((_this).F6())
				_1479_v8 = _nw220
				var _1480_v9 _dafny.Map
				_ = _1480_v9
				_1480_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1479_v8, _1467_i1)
				var _1481_v10 _dafny.Map
				_ = _1481_v10
				_1481_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1476_cf15, _1464_v3)
				var _1482_v11 _dafny.Array
				_ = _1482_v11
				var _nwElement0_51 bool = (_this).F4()
				_ = _nwElement0_51
				var _nw221 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(22))
				_ = _nw221
				(_nw221).ArraySet1(_nwElement0_51, 0)
				(_nw221).ArraySet1((_this).F4(), 1)
				(_nw221).ArraySet1(p0, 2)
				(_nw221).ArraySet1(p0, 3)
				(_nw221).ArraySet1((_this).F4(), 4)
				(_nw221).ArraySet1((_this).F4(), 5)
				(_nw221).ArraySet1(_1476_cf15, 6)
				(_nw221).ArraySet1(p0, 7)
				(_nw221).ArraySet1(p0, 8)
				(_nw221).ArraySet1((_this).F4(), 9)
				(_nw221).ArraySet1(!(!(_1476_cf15)), 10)
				(_nw221).ArraySet1((_this).F4(), 11)
				(_nw221).ArraySet1((_this).F4(), 12)
				(_nw221).ArraySet1(_1476_cf15, 13)
				(_nw221).ArraySet1(!(p0), 14)
				(_nw221).ArraySet1(p0, 15)
				(_nw221).ArraySet1((_this).F4(), 16)
				(_nw221).ArraySet1(true, 17)
				(_nw221).ArraySet1(_1476_cf15, 18)
				(_nw221).ArraySet1((_this).Fm8(_1475_cf16, _1476_cf15, (func() _dafny.Sequence {
					if (_1481_v10).Contains(p0) {
						return (_1481_v10).Get(p0).(_dafny.Sequence)
					}
					return _1464_v3
				})(), _1465_v4, globalState), 19)
				(_nw221).ArraySet1(_1476_cf15, 20)
				(_nw221).ArraySet1((_this).F4(), 21)
				_1482_v11 = _nw221
				_1461_v0 = (_1461_v0).Update(((_1480_v9).Cardinality()).Minus(_dafny.IntOfInt64(450)), (Companion_D10_.Create_DC22_(_1464_v3, p0, false, _1482_v11, (_1479_v8).F12())).Dtor_cf34())
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1462_v1), 0))
				_ = _index210
				(_1462_v1).ArraySet1(_dafny.IntOfInt64(784), (_index210).Int())
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_1462_v1), 0))
				_ = _index211
				(_1462_v1).ArraySet1((_1467_i1).Times(_dafny.IntOfInt64(925)), (_index211).Int())
				var _1483_v12 D0
				_ = _1483_v12
				_1483_v12 = Companion_D0_.Create_DC2_(_1475_cf16)
				var _1484_v13 _dafny.MultiSet
				_ = _1484_v13
				_1484_v13 = _dafny.MultiSetOf(_1478_v7)
				var _1485_v14 T1
				_ = _1485_v14
				var _nw222 *C3 = New_C3_()
				_ = _nw222
				_nw222.Ctor__(_this.F5(), (_1484_v13).Cardinality(), Companion_Default___.Fm0(_dafny.CodePoint('f'), globalState))
				_1485_v14 = _nw222
				var _1486_v15 _dafny.Sequence
				_ = _1486_v15
				_1486_v15 = _dafny.SeqOf(_1485_v14, _1485_v14, _1485_v14, _1485_v14, _1485_v14)
				var _pat_let_tv24 = _1468_v5
				_ = _pat_let_tv24
				var _pat_let_tv25 = _1475_cf16
				_ = _pat_let_tv25
				var _pat_let_tv26 = _1475_cf16
				_ = _pat_let_tv26
				var _pat_let_tv27 = _1462_v1
				_ = _pat_let_tv27
				var _pat_let_tv28 = _1462_v1
				_ = _pat_let_tv28
				var _1487_v16 _dafny.Array
				_ = _1487_v16
				var _nwElement0_52 D0 = _1483_v12
				_ = _nwElement0_52
				var _nw223 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(21))
				_ = _nw223
				(_nw223).ArraySet1(_nwElement0_52, 0)
				(_nw223).ArraySet1(_1483_v12, 1)
				(_nw223).ArraySet1(func(_pat_let30_0 D0) D0 {
					return func(_1488_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let31_0 _dafny.Int) D0 {
							return func(_1489_dt__update_hcf1_h0 _dafny.Int) D0 {
								return Companion_D0_.Create_DC2_(_1489_dt__update_hcf1_h0)
							}(_pat_let31_0)
						}(_dafny.IntOfUint32((_pat_let_tv24).Cardinality()))
					}(_pat_let30_0)
				}(Companion_D0_.Create_DC2_(_1475_cf16)), 2)
				(_nw223).ArraySet1(_1483_v12, 3)
				(_nw223).ArraySet1(_1483_v12, 4)
				(_nw223).ArraySet1(_1483_v12, 5)
				(_nw223).ArraySet1(_1483_v12, 6)
				(_nw223).ArraySet1(func(_pat_let32_0 D0) D0 {
					return func(_1490_dt__update__tmp_h1 D0) D0 {
						return func(_pat_let33_0 _dafny.Int) D0 {
							return func(_1491_dt__update_hcf1_h1 _dafny.Int) D0 {
								return Companion_D0_.Create_DC2_(_1491_dt__update_hcf1_h1)
							}(_pat_let33_0)
						}(_pat_let_tv25)
					}(_pat_let32_0)
				}(Companion_D0_.Create_DC2_(_dafny.IntOfInt64(291))), 7)
				(_nw223).ArraySet1(Companion_D0_.Create_DC2_(_dafny.IntOfInt64(431)), 8)
				(_nw223).ArraySet1(func(_pat_let34_0 D0) D0 {
					return func(_1492_dt__update__tmp_h2 D0) D0 {
						return func(_pat_let35_0 _dafny.Int) D0 {
							return func(_1493_dt__update_hcf1_h2 _dafny.Int) D0 {
								return Companion_D0_.Create_DC2_(_1493_dt__update_hcf1_h2)
							}(_pat_let35_0)
						}(_pat_let_tv26)
					}(_pat_let34_0)
				}(_1483_v12), 9)
				(_nw223).ArraySet1(_1483_v12, 10)
				(_nw223).ArraySet1(_1483_v12, 11)
				(_nw223).ArraySet1(func(_pat_let36_0 D0) D0 {
					return func(_1494_dt__update__tmp_h3 D0) D0 {
						return func(_pat_let37_0 _dafny.Int) D0 {
							return func(_1495_dt__update_hcf1_h3 _dafny.Int) D0 {
								return Companion_D0_.Create_DC2_(_1495_dt__update_hcf1_h3)
							}(_pat_let37_0)
						}((_pat_let_tv28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_pat_let_tv27), 0))).Int()).(_dafny.Int))
					}(_pat_let36_0)
				}(_1483_v12), 12)
				(_nw223).ArraySet1(_1483_v12, 13)
				(_nw223).ArraySet1(_1483_v12, 14)
				(_nw223).ArraySet1(Companion_D0_.Create_DC2_(_dafny.IntOfUint32((_1486_v15).Cardinality())), 15)
				(_nw223).ArraySet1(_1483_v12, 16)
				(_nw223).ArraySet1(Companion_D0_.Create_DC2_(_1467_i1), 17)
				(_nw223).ArraySet1(_1483_v12, 18)
				(_nw223).ArraySet1(_1483_v12, 19)
				(_nw223).ArraySet1(Companion_D0_.Create_DC2_((_1479_v8).F12()), 20)
				_1487_v16 = _nw223
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_29
				var _nw224 _dafny.Array
				_ = _nw224
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw224 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) D0 = (func(_1496_v12 D0) func(_dafny.Int) D0 {
						return func(_1497_i2 _dafny.Int) D0 {
							return _1496_v12
						}
					})(_1483_v12)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw224 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw224).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw224).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1487_v16 = _nw224
			} else {
				var _1498___mcc_h4 _dafny.CodePoint = _source26.Get_().(D4_DC12).Cf13
				_ = _1498___mcc_h4
				var _1499_cf13 _dafny.CodePoint = _1498___mcc_h4
				_ = _1499_cf13
				var _1500_v17 _dafny.Map
				_ = _1500_v17
				_1500_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1462_v1)
				_1500_v17 = (_1500_v17).Update(_1467_i1, _1462_v1)
				var _1501_v18 _dafny.Int
				_ = _1501_v18
				_1501_v18 = _dafny.IntOfInt64(277)
				var _1502_v20 _dafny.Map
				_ = _1502_v20
				_1502_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1467_i1), _1499_cf13)
				var _1503_v21 _dafny.Set
				_ = _1503_v21
				_1503_v21 = _dafny.SetOf(_1502_v20)
				_1501_v18 = Companion_Default___.SafeDivisionInt((func() _dafny.Map {
					var _coll50 = _dafny.NewMapBuilder()
					_ = _coll50
					for _iter57 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(447), _dafny.IntOfInt64(-282))); ; {
						_compr_50, _ok57 := _iter57()
						if !_ok57 {
							break
						}
						var _1504_v19 _dafny.Int
						_1504_v19 = interface{}(_compr_50).(_dafny.Int)
						if ((_dafny.IntOfInt64(447)).Cmp(_1504_v19) <= 0) && ((_1504_v19).Cmp(_dafny.IntOfInt64(-282)) < 0) {
							_coll50.Add(Companion_Default___.SafeDivisionInt(_1504_v19, _dafny.IntOfInt64(592)), _1467_i1)
						}
					}
					return _coll50.ToMap()
				}()).Cardinality(), Companion_Default___.SafeDivisionInt((_1503_v21).Cardinality(), (_this).F6()))
				var _1505_v22 _dafny.Array
				_ = _1505_v22
				var _nw225 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
				_ = _nw225
				_1505_v22 = _nw225
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1505_v22), 0))
				_ = _index212
				(_1505_v22).ArraySet1(_1465_v4, (_index212).Int())
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1505_v22), 0))
				_ = _index213
				(_1505_v22).ArraySet1(_dafny.SeqOf(_1501_v18), (_index213).Int())
				var _1506_v23 _dafny.Array
				_ = _1506_v23
				var _nw226 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(21))
				_ = _nw226
				_1506_v23 = _nw226
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1506_v23), 0))
				_ = _index214
				(_1506_v23).ArraySet1CodePoint(_1499_cf13, (_index214).Int())
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1506_v23), 0))
				_ = _index215
				(_1506_v23).ArraySet1CodePoint(_1499_cf13, (_index215).Int())
			}
			var _1507_v24 bool
			_ = _1507_v24
			_1507_v24 = true
			var _1508_v25 _dafny.MultiSet
			_ = _1508_v25
			_1508_v25 = _dafny.MultiSetOf(_1467_i1, (_this).F6())
			var _1509_v26 _dafny.Set
			_ = _1509_v26
			_1509_v26 = _dafny.SetOf(p0, (_this).F4())
			var _1510_v27 _dafny.Sequence
			_ = _1510_v27
			_1510_v27 = _dafny.SeqOf(_1465_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(980))).Uint32(), func(coer90 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg90 _dafny.Int) interface{} {
					return coer90(arg90)
				}
			}((func(_1511_i1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1512_i3 _dafny.Int) _dafny.Int {
					return _1511_i1
				}
			})(_1467_i1))))
			_1507_v24 = (_this).Fm8((func() _dafny.Int {
				if (_1508_v25).Contains((_1509_v26).Cardinality()) {
					return (_1508_v25).Multiplicity((_1509_v26).Cardinality())
				}
				return _dafny.IntOfInt64(622)
			})(), _dafny.Companion_Sequence_.Equal(_1510_v27, _dafny.SeqOf(_1465_v4)), _1464_v3, _dafny.Companion_Sequence_.Concatenate(_1465_v4, _1465_v4), globalState)
			var _1513_v28 D11
			_ = _1513_v28
			_1513_v28 = Companion_D11_.Create_DC24_((_this).F4(), (_this).F4(), _dafny.IntOfUint32((_1464_v3).Cardinality()))
			var _1514_v29 D20
			_ = _1514_v29
			_1514_v29 = Companion_D20_.Create_DC41_((func() _dafny.Int {
				if _1507_v24 {
					return _1467_i1
				}
				return (_1513_v28).Dtor_cf38()
			})())
			_1514_v29 = Companion_D20_.Create_DC41_(_dafny.IntOfInt64(130))
			var _1515_v30 _dafny.Int
			_ = _1515_v30
			_1515_v30 = _dafny.IntOfInt64(174)
			var _1516_v31 _dafny.Set
			_ = _1516_v31
			_1516_v31 = _dafny.SetOf(_1467_i1)
			var _1517_v32 _dafny.Set
			_ = _1517_v32
			_1517_v32 = _1516_v31
			var _1518_v33 _dafny.Set
			_ = _1518_v33
			_1518_v33 = _dafny.SetOf(_1517_v32)
			_1515_v30 = ((_dafny.Zero).Minus(((_dafny.Zero).Minus(_1515_v30)).Minus((_this).F6()))).Plus((func() _dafny.Int {
				if (_this).F4() {
					return _1515_v30
				}
				return (_1518_v33).Cardinality()
			})())
		}
		var _1519_v34 bool
		_ = _1519_v34
		_1519_v34 = false
		var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1462_v1), 0))
		_ = _index216
		(_1462_v1).ArraySet1(((_this).F6()).Times((_this).F6()), (_index216).Int())
		var _1520_v35 _dafny.Int
		_ = _1520_v35
		_1520_v35 = _dafny.IntOfInt64(978)
		var _1521_v36 _dafny.Map
		_ = _1521_v36
		_1521_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1465_v4).Select((Companion_Default___.SafeIndex(_1520_v35, _dafny.IntOfUint32((_1465_v4).Cardinality()))).Uint32()).(_dafny.Int), _1519_v34)
		var _1522_v37 _dafny.Map
		_ = _1522_v37
		_1522_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_this).F4()), _1521_v36)
		var _1523_v38 _dafny.Sequence
		_ = _1523_v38
		_1523_v38 = _dafny.SeqOf(_dafny.MultiSetOf((_this).F6()))
		var _1524_v39 _dafny.MultiSet
		_ = _1524_v39
		_1524_v39 = _dafny.MultiSetOf(p0)
		var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1462_v1), 0))
		_ = _index217
		var _rhs193 bool = !(_1522_v37).Equals(Companion_Default___.Fm60(_dafny.IntOfUint32((_1523_v38).Cardinality()), (_this).F6(), _1520_v35, globalState))
		_ = _rhs193
		var _rhs194 _dafny.Int = _1520_v35
		_ = _rhs194
		var _rhs195 _dafny.Int = Companion_Default___.SafeDivisionInt((_1524_v39).Cardinality(), _1520_v35)
		_ = _rhs195
		var _rhs196 bool = !(!(_1519_v34))
		_ = _rhs196
		var _rhs197 bool = (_this).F4()
		_ = _rhs197
		var _lhs121 _dafny.Array = _1462_v1
		_ = _lhs121
		var _lhs122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1462_v1), 0))
		_ = _lhs122
		_1519_v34 = _rhs193
		(_lhs121).ArraySet1(_rhs194, (_lhs122).Int())
		_1520_v35 = _rhs195
		_1519_v34 = _rhs196
		_1519_v34 = _rhs197
		var _1525_v40 _dafny.Array
		_ = _1525_v40
		var _len0_30 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_30
		var _nw227 _dafny.Array
		_ = _nw227
		if _len0_30.Cmp(_dafny.Zero) == 0 {
			_nw227 = _dafny.NewArray(_len0_30)
		} else {
			var _init30 func(_dafny.Int) D12 = (func(_1526_p0 bool) func(_dafny.Int) D12 {
				return func(_1527_i4 _dafny.Int) D12 {
					return Companion_D12_.Create_DC27_(_dafny.SeqOf(_dafny.SeqOf(false), _dafny.SeqOf(_1526_p0)))
				}
			})(p0)
			_ = _init30
			var _element0_30 = _init30(_dafny.Zero)
			_ = _element0_30
			_nw227 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
			(_nw227).ArraySet1(_element0_30, 0)
			var _nativeLen0_30 = (_len0_30).Int()
			_ = _nativeLen0_30
			for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
				(_nw227).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
			}
		}
		_1525_v40 = _nw227
		var _1528_v41 _dafny.Set
		_ = _1528_v41
		_1528_v41 = _dafny.SetOf((_this).F4(), (_this).F4())
		var _1529_v42 D12
		_ = _1529_v42
		_1529_v42 = Companion_D12_.Create_DC27_(Companion_Default___.Fm61(_1528_v41, globalState))
		var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_1525_v40), 0))
		_ = _index218
		(_1525_v40).ArraySet1(_1529_v42, (_index218).Int())
		var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_1525_v40), 0))
		_ = _index219
		var _rhs198 _dafny.Int = _1520_v35
		_ = _rhs198
		var _rhs199 D12 = _1529_v42
		_ = _rhs199
		var _lhs123 _dafny.Array = _1525_v40
		_ = _lhs123
		var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_1525_v40), 0))
		_ = _lhs124
		_1520_v35 = _rhs198
		(_lhs123).ArraySet1(_rhs199, (_lhs124).Int())
		var _1530_v43 *C1
		_ = _1530_v43
		var _nw228 *C1 = New_C1_()
		_ = _nw228
		_nw228.Ctor__(_1520_v35, Companion_Default___.SafeModuloInt((_1462_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1462_v1), 0))).Int()).(_dafny.Int), _1520_v35), _this.F5(), (_this).F6(), (_this).Fm8((_this).F6(), (_this).F4(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(566))).Uint32(), func(coer91 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg91 _dafny.Int) interface{} {
				return coer91(arg91)
			}
		}(func(_1531_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-532))).Uint32(), func(coer92 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg92 _dafny.Int) interface{} {
				return coer92(arg92)
			}
		}((func(_1532_v35 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1533_i6 _dafny.Int) _dafny.Int {
				return _1532_v35
			}
		})(_1520_v35))), globalState))
		_1530_v43 = _nw228
		var _1534_v44 _dafny.CodePoint
		_ = _1534_v44
		_1534_v44 = _dafny.CodePoint('f')
		var _1535_v45 D9
		_ = _1535_v45
		_1535_v45 = Companion_D9_.Create_DC20_((_this).F4(), _1534_v44, _1530_v43.F16)
		var _1536_v46 _dafny.Sequence
		_ = _1536_v46
		_1536_v46 = _dafny.SeqOf(_1535_v45)
		var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1462_v1), 0))
		_ = _index220
		var _rhs200 *C1 = (func() *C1 {
			if ((func() _dafny.Int {
				if (_1461_v0).Contains((_1462_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1462_v1), 0))).Int()).(_dafny.Int)) {
					return (_1461_v0).Get((_1462_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1462_v1), 0))).Int()).(_dafny.Int)).(_dafny.Int)
				}
				return (_1462_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1462_v1), 0))).Int()).(_dafny.Int)
			})()).Cmp((_1462_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1462_v1), 0))).Int()).(_dafny.Int)) <= 0 {
				return _1530_v43
			}
			return _1530_v43
		})()
		_ = _rhs200
		var _rhs201 bool = p0
		_ = _rhs201
		var _rhs202 bool = p0
		_ = _rhs202
		var _rhs203 _dafny.Int = (_dafny.Zero).Minus(func(_source27 _dafny.Sequence) _dafny.Int {
			var _1537___mcc_h5 _dafny.Sequence = _source27
			_ = _1537___mcc_h5
			var _1538_cf57 _dafny.Sequence = _1537___mcc_h5
			_ = _1538_cf57
			return _dafny.IntOfInt64(911)
		}(_1536_v46))
		_ = _rhs203
		var _lhs125 _dafny.Array = _1462_v1
		_ = _lhs125
		var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1462_v1), 0))
		_ = _lhs126
		_1530_v43 = _rhs200
		_1519_v34 = _rhs201
		_1519_v34 = _rhs202
		(_lhs125).ArraySet1(_rhs203, (_lhs126).Int())
		r0 = _1534_v44
		var _1539_v47 _dafny.Array
		_ = _1539_v47
		var _nw229 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(2))
		_ = _nw229
		_1539_v47 = _nw229
		r1 = _1539_v47
		return r0, r1
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f5  _dafny.Array
	_f6  _dafny.Int
	_f4  bool
	F14  _dafny.Int
	_f13 bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.Zero
	_this._f4 = false
	_this.F14 = _dafny.Zero
	_this._f13 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F5() _dafny.Array {
	return _this._f5
}
func (_this *C5) F5_set_(value _dafny.Array) {
	_this._f5 = value
}
func (_this *C5) F6() _dafny.Int {
	return _this._f6
}
func (_this *C5) F4() bool {
	return _this._f4
}
func (_this *C5) Ctor__(f13 bool, f14 _dafny.Int, f5 _dafny.Array, f6 _dafny.Int, f4 bool) {
	{
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f4 = f4
	}
}
func (_this *C5) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		if (_this).F13() {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F4())).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13()))
		} else {
			return !(_dafny.MultiSetOf(_this.F14)).Contains((_dafny.Zero).Minus(_this.F14))
		}
	}
}
func (_this *C5) Fm9(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(30)
	}
}
func (_this *C5) Fm6(p0 D0, p1 D0, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetOf((_this).F6(), _dafny.IntOfInt64(309), (_this).F6())).Difference(_dafny.MultiSetOf((_dafny.Zero).Minus(_this.F14)))
	}
}
func (_this *C5) Fm7(globalState *GlobalState) _dafny.Int {
	{
		var _source28 D3 = Companion_D3_.Create_DC9_((_this).F13(), (_this).F6(), (_this).F4())
		_ = _source28
		if _source28.Is_DC9() {
			var _1540___mcc_h0 bool = _source28.Get_().(D3_DC9).Cf8
			_ = _1540___mcc_h0
			var _1541___mcc_h1 _dafny.Int = _source28.Get_().(D3_DC9).Cf9
			_ = _1541___mcc_h1
			var _1542___mcc_h2 bool = _source28.Get_().(D3_DC9).Cf10
			_ = _1542___mcc_h2
			var _1543_cf10 bool = _1542___mcc_h2
			_ = _1543_cf10
			var _1544_cf9 _dafny.Int = _1541___mcc_h1
			_ = _1544_cf9
			var _1545_cf8 bool = _1540___mcc_h0
			_ = _1545_cf8
			return (_this).F6()
		} else if _source28.Is_DC10() {
			var _1546___mcc_h3 _dafny.Int = _source28.Get_().(D3_DC10).Cf11
			_ = _1546___mcc_h3
			var _1547_cf11 _dafny.Int = _1546___mcc_h3
			_ = _1547_cf11
			return (_this.F14).Plus(_this.F14)
		} else if _source28.Is_DC8() {
			var _1548___mcc_h4 _dafny.Sequence = _source28.Get_().(D3_DC8).Cf7
			_ = _1548___mcc_h4
			var _1549_cf7 _dafny.Sequence = _1548___mcc_h4
			_ = _1549_cf7
			return (_dafny.IntOfInt64(64)).Minus(_dafny.IntOfUint32(((Companion_D9_.Create_DC19_(_dafny.UnicodeSeqOfUtf8Bytes("fafdlayl"))).Dtor_cf25()).Cardinality()))
		} else {
			var _1550___mcc_h5 D3 = _source28.Get_().(D3_DC11).Cf12
			_ = _1550___mcc_h5
			var _1551_cf12 D3 = _1550___mcc_h5
			_ = _1551_cf12
			return (_this).F6()
		}
	}
}
func (_this *C5) Fm26(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(_this.F14)
	}
}
func (_this *C5) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _1552_v0 _dafny.MultiSet
		_ = _1552_v0
		_1552_v0 = _dafny.MultiSetOf((_this).F13(), (_this).F4(), (_this).F4(), (_this).F13(), (_this).F13())
		var _1553_v1 _dafny.Map
		_ = _1553_v1
		_1553_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC15_(Companion_Default___.Fm27((_this).F13(), _dafny.IntOfInt64(35), _1552_v0, (_this).F13(), globalState)), ((_1552_v0).Intersection(_1552_v0)).Cardinality())
		_1553_v1 = _1553_v1
		(_this).F14 = (_this).F6()
		var _1554_v2 _dafny.Set
		_ = _1554_v2
		_1554_v2 = _dafny.SetOf((_this).F4())
		var _1555_v3 _dafny.Sequence
		_ = _1555_v3
		_1555_v3 = _dafny.SeqOf(_1554_v2)
		var _1556_v4 _dafny.Map
		_ = _1556_v4
		_1556_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _dafny.IntOfUint32((_1555_v3).Cardinality()))
		var _1557_v5 _dafny.CodePoint
		_ = _1557_v5
		_1557_v5 = _dafny.CodePoint('a')
		var _1558_v6 D9
		_ = _1558_v6
		_1558_v6 = Companion_D9_.Create_DC20_((_this).F13(), _1557_v5, (_this).F6())
		var _1559_v7 *C0
		_ = _1559_v7
		var _nw230 *C0 = New_C0_()
		_ = _nw230
		_nw230.Ctor__((func() _dafny.Int {
			if (_1556_v4).Contains((_1558_v6).Dtor_cf26()) {
				return (_1556_v4).Get((_1558_v6).Dtor_cf26()).(_dafny.Int)
			}
			return p0
		})())
		_1559_v7 = _nw230
		var _1560_v8 _dafny.Sequence
		_ = _1560_v8
		_1560_v8 = _dafny.SeqOf((_this).F13())
		var _1561_v9 _dafny.Sequence
		_ = _1561_v9
		_1561_v9 = _dafny.SeqOf(_dafny.SeqOf((_this).F4(), (_this).F4(), (_this).F13(), (_this).F4(), (_this).F4()))
		var _1562_v11 _dafny.Sequence
		_ = _1562_v11
		_1562_v11 = _dafny.UnicodeSeqOfUtf8Bytes("trc")
		var _1563_v12 _dafny.Array
		_ = _1563_v12
		var _len0_31 _dafny.Int = _dafny.One
		_ = _len0_31
		var _nw231 _dafny.Array
		_ = _nw231
		if _len0_31.Cmp(_dafny.Zero) == 0 {
			_nw231 = _dafny.NewArray(_len0_31)
		} else {
			var _init31 func(_dafny.Int) _dafny.Int = func(_1564_i0 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_1564_i0, (_this).F6())
			}
			_ = _init31
			var _element0_31 = _init31(_dafny.Zero)
			_ = _element0_31
			_nw231 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
			(_nw231).ArraySet1(_element0_31, 0)
			var _nativeLen0_31 = (_len0_31).Int()
			_ = _nativeLen0_31
			for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
				(_nw231).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
			}
		}
		_1563_v12 = _nw231
		var _1565_v13 _dafny.Map
		_ = _1565_v13
		_1565_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _1563_v12)
		var _1566_v14 _dafny.Array
		_ = _1566_v14
		var _nwElement0_53 _dafny.Int = _this.F14
		_ = _nwElement0_53
		var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(20))
		_ = _nw232
		(_nw232).ArraySet1(_nwElement0_53, 0)
		(_nw232).ArraySet1((_dafny.Zero).Minus(_this.F14), 1)
		(_nw232).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(59), (_1559_v7).F12(), (_1559_v7).F12())).Cardinality()), 2)
		(_nw232).ArraySet1((_this.F14).Minus(_this.F14), 3)
		(_nw232).ArraySet1((_1559_v7).F12(), 4)
		(_nw232).ArraySet1((_this).F6(), 5)
		(_nw232).ArraySet1(p0, 6)
		(_nw232).ArraySet1((_this).F6(), 7)
		(_nw232).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_1560_v8)).Cardinality(), (_1559_v7).F12())).Cardinality(), 8)
		(_nw232).ArraySet1((_this).F6(), 9)
		(_nw232).ArraySet1(_this.F14, 10)
		(_nw232).ArraySet1((_1559_v7).F12(), 11)
		(_nw232).ArraySet1(_dafny.IntOfUint32((_1561_v9).Cardinality()), 12)
		(_nw232).ArraySet1((func() _dafny.Set {
			var _coll51 = _dafny.NewBuilder()
			_ = _coll51
			for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(253), _dafny.IntOfInt64(-291))); ; {
				_compr_51, _ok58 := _iter58()
				if !_ok58 {
					break
				}
				var _1567_v10 _dafny.Int
				_1567_v10 = interface{}(_compr_51).(_dafny.Int)
				if ((_dafny.IntOfInt64(253)).Cmp(_1567_v10) <= 0) && ((_1567_v10).Cmp(_dafny.IntOfInt64(-291)) < 0) {
					_coll51.Add((_1567_v10).Times((_1559_v7).F12()))
				}
			}
			return _coll51.ToSet()
		}()).Cardinality(), 13)
		(_nw232).ArraySet1(((_this).F6()).Minus(_dafny.IntOfUint32((_1562_v11).Cardinality())), 14)
		(_nw232).ArraySet1((_1565_v13).Cardinality(), 15)
		(_nw232).ArraySet1((_this).F6(), 16)
		(_nw232).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(438), _this.F14), 17)
		(_nw232).ArraySet1((_this).F6(), 18)
		(_nw232).ArraySet1((_1559_v7).F12(), 19)
		_1566_v14 = _nw232
		var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_1563_v12), 0))
		_ = _index221
		(_1563_v12).ArraySet1(_this.F14, (_index221).Int())
		var _1568_v15 D10
		_ = _1568_v15
		_1568_v15 = Companion_D10_.Create_DC21_(Companion_Default___.Fm28(globalState))
		var _1569_v16 _dafny.Sequence
		_ = _1569_v16
		_1569_v16 = _dafny.SeqOf(_1555_v3, _1555_v3)
		var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_1563_v12), 0))
		_ = _index222
		var _rhs204 _dafny.Int = (func() _dafny.Int {
			if (_this).F13() {
				return (_1559_v7).F12()
			}
			return (_1559_v7).F12()
		})()
		_ = _rhs204
		var _rhs205 _dafny.Map = (_1568_v15).Dtor_cf29()
		_ = _rhs205
		var _rhs206 _dafny.Sequence = (func() _dafny.Sequence {
			if (_this.F14).Cmp(_dafny.IntOfInt64(521)) != 0 {
				return _dafny.SeqOf(_1554_v2)
			}
			return (_1569_v16).Select((Companion_Default___.SafeIndex((_1559_v7).F12(), _dafny.IntOfUint32((_1569_v16).Cardinality()))).Uint32()).(_dafny.Sequence)
		})()
		_ = _rhs206
		var _lhs127 _dafny.Array = _1563_v12
		_ = _lhs127
		var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_1563_v12), 0))
		_ = _lhs128
		(_lhs127).ArraySet1(_rhs204, (_lhs128).Int())
		_1556_v4 = _rhs205
		_1555_v3 = _rhs206
		var _1570_v17 _dafny.Map
		_ = _1570_v17
		_1570_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1557_v5, _1557_v5)
		_1570_v17 = (_1570_v17).Update(_dafny.CodePoint('w'), _1557_v5)
		var _1571_v18 _dafny.MultiSet
		_ = _1571_v18
		_1571_v18 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1562_v11).Cardinality()), (_1554_v2).Cardinality())
		var _1572_v19 _dafny.Map
		_ = _1572_v19
		_1572_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1571_v18, _dafny.UnicodeSeqOfUtf8Bytes("ydutpac"))
		var _1573_v20 bool
		_ = _1573_v20
		var _1574_v21 _dafny.Int
		_ = _1574_v21
		var _out47 bool
		_ = _out47
		var _out48 _dafny.Int
		_ = _out48
		_out47, _out48 = Companion_Default___.M0((_1572_v19).Merge(_1572_v19), false, _1571_v18, (_this).F4(), globalState)
		_1573_v20 = _out47
		_1574_v21 = _out48
		r0 = _1556_v4
		return r0
	}
}
func (_this *C5) M2(p0 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Array) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _1575_v1 _dafny.Sequence
		_ = _1575_v1
		_1575_v1 = _dafny.UnicodeSeqOfUtf8Bytes("rptj")
		var _1576_v2 _dafny.MultiSet
		_ = _1576_v2
		_1576_v2 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1575_v1).Cardinality()))
		var _1577_v3 _dafny.Sequence
		_ = _1577_v3
		_1577_v3 = _dafny.SeqOf(_this.F14, (_this).F6(), _this.F14, _this.F14, (_this).F6())
		var _1578_v4 _dafny.Sequence
		_ = _1578_v4
		_1578_v4 = _dafny.SeqOf(_1576_v2, _dafny.MultiSetFromSeq(_1577_v3))
		var _1579_v5 bool
		_ = _1579_v5
		var _1580_v6 _dafny.Int
		_ = _1580_v6
		var _out49 bool
		_ = _out49
		var _out50 _dafny.Int
		_ = _out50
		_out49, _out50 = Companion_Default___.M0(func() _dafny.Map {
			var _coll52 = _dafny.NewMapBuilder()
			_ = _coll52
			for _iter59 := _dafny.Iterate((_1578_v4).Elements()); ; {
				_compr_52, _ok59 := _iter59()
				if !_ok59 {
					break
				}
				var _1581_v0 _dafny.MultiSet
				_1581_v0 = interface{}(_compr_52).(_dafny.MultiSet)
				if _dafny.Companion_Sequence_.Contains(_1578_v4, _1581_v0) {
					_coll52.Add(_1581_v0, _1575_v1)
				}
			}
			return _coll52.ToMap()
		}(), (_this).F4(), _1576_v2, ((_this).F6()).Cmp((_this).F6()) != 0, globalState)
		_1579_v5 = _out49
		_1580_v6 = _out50
		_1575_v1 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(592))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg93 _dafny.Int) interface{} {
				return coer93(arg93)
			}
		}(func(_1582_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		}))
		var _1583_v7 *C0
		_ = _1583_v7
		var _nw233 *C0 = New_C0_()
		_ = _nw233
		_nw233.Ctor__(_dafny.IntOfUint32((_dafny.SeqOf((_this).F4(), !(_1579_v5), !((_this).F13()))).Cardinality()))
		_1583_v7 = _nw233
		var _1584_v8 _dafny.CodePoint
		_ = _1584_v8
		_1584_v8 = _dafny.CodePoint('p')
		var _1585_v9 _dafny.Map
		_ = _1585_v9
		_1585_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14, Companion_Default___.Fm0(_1584_v8, globalState))
		var _1586_v10 _dafny.MultiSet
		_ = _1586_v10
		_1586_v10 = _dafny.MultiSetOf(_1584_v8)
		_1585_v9 = (_1585_v9).Update((_dafny.Zero).Minus((_1586_v10).Cardinality()), (_1580_v6).Cmp((_dafny.Zero).Minus(_this.F14)) == 0)
		var _hi10 _dafny.Int = _dafny.IntOfInt64(496)
		_ = _hi10
		for _1587_i1 := ((_this).F6()).Minus((_this).F6()); _1587_i1.Cmp(_hi10) < 0; _1587_i1 = _1587_i1.Plus(_dafny.One) {
			if (false) && (!(p0)) {
				var _1588_v11 _dafny.Array
				_ = _1588_v11
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_32
				var _nw234 _dafny.Array
				_ = _nw234
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw234 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) bool = func(_1589_i2 _dafny.Int) bool {
						return (_this).F4()
					}
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw234 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw234).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw234).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1588_v11 = _nw234
				var _1590_v12 _dafny.Map
				_ = _1590_v12
				_1590_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v11, p0)
				var _1591_v13 *C0
				_ = _1591_v13
				var _nw235 *C0 = New_C0_()
				_ = _nw235
				_nw235.Ctor__((_this).Fm26((_1583_v7).F12(), (_this).Fm8((_this).F6(), true, _1575_v1, _1577_v3, globalState), (_1590_v12).Cardinality(), globalState))
				_1591_v13 = _nw235
				var _1592_v14 _dafny.Map
				_ = _1592_v14
				_1592_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F6())
				var _1593_v15 _dafny.Map
				_ = _1593_v15
				_1593_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1592_v14).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(218))).Uint32(), func(coer94 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg94 _dafny.Int) interface{} {
						return coer94(arg94)
					}
				}((func(_1594_v13 *C0) func(_dafny.Int) _dafny.Int {
					return func(_1595_i3 _dafny.Int) _dafny.Int {
						return (_1594_v13).F12()
					}
				})(_1591_v13)))).Cardinality()))
				_1593_v15 = (_1593_v15).Update((func() _dafny.Map {
					var _coll53 = _dafny.NewMapBuilder()
					_ = _coll53
					for _iter60 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(444), _dafny.IntOfInt64(515))); ; {
						_compr_53, _ok60 := _iter60()
						if !_ok60 {
							break
						}
						var _1596_v16 _dafny.Int
						_1596_v16 = interface{}(_compr_53).(_dafny.Int)
						if ((_dafny.IntOfInt64(444)).Cmp(_1596_v16) <= 0) && ((_1596_v16).Cmp(_dafny.IntOfInt64(515)) < 0) {
							_coll53.Add((_1596_v16).Times((_1583_v7).F12()), (_dafny.MultiSetOf(_dafny.IntOfInt64(610), _this.F14)).Cardinality())
						}
					}
					return _coll53.ToMap()
				}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(82))).Uint32(), func(coer95 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg95 _dafny.Int) interface{} {
						return coer95(arg95)
					}
				}((func(_1597_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1598_i4 _dafny.Int) _dafny.Sequence {
						return _1597_v3
					}
				})(_1577_v3))), (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfInt64(541))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(82))).Uint32(), func(coer96 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg96 _dafny.Int) interface{} {
						return coer96(arg96)
					}
				}((func(_1599_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1600_i4 _dafny.Int) _dafny.Sequence {
						return _1599_v3
					}
				})(_1577_v3)))).Cardinality()))).Uint32(), _1577_v3)).Cardinality()))).Cardinality()))
				var _1601_v17 _dafny.Array
				_ = _1601_v17
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_33
				var _nw236 _dafny.Array
				_ = _nw236
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw236 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) _dafny.Int = (func(_1602_v7 *C0) func(_dafny.Int) _dafny.Int {
						return func(_1603_i5 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1603_i5, (_1602_v7).F12())
						}
					})(_1583_v7)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw236 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw236).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw236).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1601_v17 = _nw236
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_1601_v17), 0))
				_ = _index223
				(_1601_v17).ArraySet1((_dafny.IntOfInt64(40)).Times(_this.F14), (_index223).Int())
				var _1604_v18 _dafny.MultiSet
				_ = _1604_v18
				_1604_v18 = _dafny.MultiSetOf((_this).F4(), p0, (_this).F4())
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_1601_v17), 0))
				_ = _index224
				(_1601_v17).ArraySet1(Companion_Default___.SafeDivisionInt((_this.F14).Minus((_1604_v18).Cardinality()), (_this).F6()), (_index224).Int())
				_1579_v5 = (_this).F4()
				_1580_v6 = (_1583_v7).F12()
			} else {
				var _1605_v19 _dafny.Map
				_ = _1605_v19
				_1605_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F14).Minus((_1583_v7).F12()), (_dafny.IntOfInt64(-598)).Minus((_1583_v7).F12()))
				_1605_v19 = Companion_Default___.Fm29(globalState)
				_1580_v6 = (_1583_v7).F12()
				r0 = _1584_v8
				_1584_v8 = _1584_v8
				var _rhs207 _dafny.Int = _this.F14
				_ = _rhs207
				var _rhs208 _dafny.Int = ((_1583_v7).F12()).Times(_dafny.IntOfUint32((_1575_v1).Cardinality()))
				_ = _rhs208
				var _rhs209 _dafny.Int = (_this).F6()
				_ = _rhs209
				var _rhs210 _dafny.CodePoint = _1584_v8
				_ = _rhs210
				var _rhs211 _dafny.Int = (_1583_v7).F12()
				_ = _rhs211
				var _lhs129 *C5 = _this
				_ = _lhs129
				var _lhs130 *C5 = _this
				_ = _lhs130
				var _lhs131 *C5 = _this
				_ = _lhs131
				_1580_v6 = _rhs207
				_lhs129.F14 = _rhs208
				_lhs130.F14 = _rhs209
				r0 = _rhs210
				_lhs131.F14 = _rhs211
			}
			if _1579_v5 {
				(_this).F14 = _1587_i1
				(_this).F14 = (_1587_i1).Times(_1580_v6)
				_1580_v6 = (_1583_v7).F12()
				_1579_v5 = !((_this).F13())
				var _1606_v20 _dafny.Array
				_ = _1606_v20
				var _nw237 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
				_ = _nw237
				_1606_v20 = _nw237
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_1606_v20), 0))
				_ = _index225
				(_1606_v20).ArraySet1(_1575_v1, (_index225).Int())
				var _1607_v21 _dafny.Sequence
				_ = _1607_v21
				_1607_v21 = _dafny.SeqOf(_1575_v1)
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_1606_v20), 0))
				_ = _index226
				(_1606_v20).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_1607_v21).Select((Companion_Default___.SafeIndex((_1583_v7).F12(), _dafny.IntOfUint32((_1607_v21).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate((_1607_v21).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1577_v3).Cardinality()), _dafny.IntOfUint32((_1607_v21).Cardinality()))).Uint32()).(_dafny.Sequence), _1575_v1)), (_index226).Int())
			} else {
				var _1608_v22 T0
				_ = _1608_v22
				var _nw238 *C2 = New_C2_()
				_ = _nw238
				_nw238.Ctor__(_1575_v1, p0)
				_1608_v22 = _nw238
				var _1609_v23 _dafny.Map
				_ = _1609_v23
				_1609_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1579_v5, _1608_v22)
				_1609_v23 = (_1609_v23).Update(!((_this).F13()) || ((_this).F13()), _1608_v22)
				_1579_v5 = !(p0)
				var _1610_v24 D4
				_ = _1610_v24
				_1610_v24 = Companion_D4_.Create_DC12_(_1584_v8)
				var _1611_v25 _dafny.Array
				_ = _1611_v25
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_34
				var _nw239 _dafny.Array
				_ = _nw239
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw239 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) _dafny.CodePoint = func(_1612_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('d')
					}
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw239 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw239).ArraySet1CodePoint(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw239).ArraySet1CodePoint(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1611_v25 = _nw239
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_1611_v25), 0))
				_ = _index227
				(_1611_v25).ArraySet1CodePoint(_1584_v8, (_index227).Int())
				var _1613_v26 _dafny.Sequence
				_ = _1613_v26
				_1613_v26 = _dafny.SeqOf(!(_1579_v5))
				var _1614_v27 _dafny.Array
				_ = _1614_v27
				var _nwElement0_54 bool = !((_this).F4())
				_ = _nwElement0_54
				var _nw240 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(11))
				_ = _nw240
				(_nw240).ArraySet1(_nwElement0_54, 0)
				(_nw240).ArraySet1((_this).F4(), 1)
				(_nw240).ArraySet1((_this).F4(), 2)
				(_nw240).ArraySet1((_this).F13(), 3)
				(_nw240).ArraySet1(false, 4)
				(_nw240).ArraySet1((_this).F13(), 5)
				(_nw240).ArraySet1(false, 6)
				(_nw240).ArraySet1(!(p0) || (true), 7)
				(_nw240).ArraySet1(!((_this).F4()), 8)
				(_nw240).ArraySet1(_dafny.Companion_Sequence_.Equal(_1613_v26, _1613_v26), 9)
				(_nw240).ArraySet1(Companion_Default___.Fm0(_1584_v8, globalState), 10)
				_1614_v27 = _nw240
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_1614_v27), 0))
				_ = _index228
				(_1614_v27).ArraySet1((_this).Fm8(_this.F14, p0, _1575_v1, _dafny.SeqOf(_this.F14), globalState), (_index228).Int())
				var _1615_v28 D10
				_ = _1615_v28
				_1615_v28 = Companion_D10_.Create_DC22_(_1575_v1, p0, (_this).F13(), _1614_v27, _1587_i1)
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_1611_v25), 0))
				_ = _index229
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_1614_v27), 0))
				_ = _index230
				var _rhs212 D4 = Companion_Default___.Fm62(globalState)
				_ = _rhs212
				var _rhs213 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (_1615_v28).Dtor_cf31() {
						return _1584_v8
					}
					return _dafny.CodePoint('d')
				})()
				_ = _rhs213
				var _rhs214 bool = (p0) == ((_1587_i1).Cmp((_1583_v7).F12()) == 0)
				_ = _rhs214
				var _rhs215 _dafny.Array = _1614_v27
				_ = _rhs215
				var _lhs132 _dafny.Array = _1611_v25
				_ = _lhs132
				var _lhs133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_1611_v25), 0))
				_ = _lhs133
				var _lhs134 _dafny.Array = _1614_v27
				_ = _lhs134
				var _lhs135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_1614_v27), 0))
				_ = _lhs135
				_1610_v24 = _rhs212
				(_lhs132).ArraySet1CodePoint(_rhs213, (_lhs133).Int())
				(_lhs134).ArraySet1(_rhs214, (_lhs135).Int())
				_1614_v27 = _rhs215
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_1614_v27), 0))
				_ = _index231
				var _rhs216 bool = p0
				_ = _rhs216
				var _rhs217 bool = (_1580_v6).Cmp((func() _dafny.Int {
					if _1579_v5 {
						return _dafny.IntOfInt64(909)
					}
					return _dafny.IntOfUint32((_1575_v1).Cardinality())
				})()) >= 0
				_ = _rhs217
				var _lhs136 _dafny.Array = _1614_v27
				_ = _lhs136
				var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_1614_v27), 0))
				_ = _lhs137
				(_lhs136).ArraySet1(_rhs216, (_lhs137).Int())
				_1579_v5 = _rhs217
				_1579_v5 = p0
			}
			var _1616_v29 _dafny.Array
			_ = _1616_v29
			var _nwElement0_55 bool = (_this).F13()
			_ = _nwElement0_55
			var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(4))
			_ = _nw241
			(_nw241).ArraySet1(_nwElement0_55, 0)
			(_nw241).ArraySet1(_1579_v5, 1)
			(_nw241).ArraySet1((_dafny.IntOfInt64(49)).Cmp(_1580_v6) < 0, 2)
			(_nw241).ArraySet1((_this).F4(), 3)
			_1616_v29 = _nw241
			_1616_v29 = _1616_v29
			var _1617_v30 _dafny.Set
			_ = _1617_v30
			_1617_v30 = _dafny.SetOf(p0, false)
			(_this).F14 = ((Companion_D16_.Create_DC33_(_1579_v5, _dafny.IntOfUint32((_1575_v1).Cardinality()), _dafny.MultiSetOf(p0, true), _1580_v6, _1617_v30)).Dtor_cf54()).Cardinality()
		}
		var _1618_v31 _dafny.Sequence
		_ = _1618_v31
		_1618_v31 = _dafny.SeqOf(p0, _1579_v5, _1579_v5)
		var _1619_v32 _dafny.Sequence
		_ = _1619_v32
		_1619_v32 = _dafny.SeqOf(_1618_v31)
		var _1620_v33 D12
		_ = _1620_v33
		_1620_v33 = Companion_D12_.Create_DC27_(_1619_v32)
		var _source29 D12 = _1620_v33
		_ = _source29
		if _source29.Is_DC28() {
			var _1621___mcc_h0 _dafny.Int = _source29.Get_().(D12_DC28).Cf43
			_ = _1621___mcc_h0
			var _1622___mcc_h1 _dafny.Int = _source29.Get_().(D12_DC28).Cf44
			_ = _1622___mcc_h1
			var _1623___mcc_h2 bool = _source29.Get_().(D12_DC28).Cf45
			_ = _1623___mcc_h2
			var _1624___mcc_h3 bool = _source29.Get_().(D12_DC28).Cf46
			_ = _1624___mcc_h3
			var _1625___mcc_h4 _dafny.Int = _source29.Get_().(D12_DC28).Cf47
			_ = _1625___mcc_h4
			var _1626_cf47 _dafny.Int = _1625___mcc_h4
			_ = _1626_cf47
			var _1627_cf46 bool = _1624___mcc_h3
			_ = _1627_cf46
			var _1628_cf45 bool = _1623___mcc_h2
			_ = _1628_cf45
			var _1629_cf44 _dafny.Int = _1622___mcc_h1
			_ = _1629_cf44
			var _1630_cf43 _dafny.Int = _1621___mcc_h0
			_ = _1630_cf43
			var _1631_v34 _dafny.Array
			_ = _1631_v34
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_35
			var _nw242 _dafny.Array
			_ = _nw242
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw242 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) _dafny.Map = func(_1632_i7 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), false)
				}
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw242 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw242).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw242).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1631_v34 = _nw242
			_1631_v34 = _1631_v34
			if Companion_Default___.Fm0(_1584_v8, globalState) {
				var _1633_v35 _dafny.Set
				_ = _1633_v35
				_1633_v35 = _dafny.SetOf((_this).F13(), _1627_cf46)
				var _1634_v36 _dafny.Map
				_ = _1634_v36
				_1634_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1633_v35).Intersection(_1633_v35), (_this).F4())
				_1634_v36 = (_1634_v36).Update((_1633_v35).Union(_dafny.SetOf((_this).F4())), !((_this).F4()) || (_1579_v5))
				_1579_v5 = (_this).F13()
				var _1635_v37 *C3
				_ = _1635_v37
				var _nw243 *C3 = New_C3_()
				_ = _nw243
				_nw243.Ctor__(_this.F5(), (_dafny.Zero).Minus(_this.F14), p0)
				_1635_v37 = _nw243
				var _1636_v38 _dafny.Map
				_ = _1636_v38
				_1636_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1635_v37, (_this).F13())
				_1636_v38 = (_1636_v38).Update(_1635_v37, _1627_cf46)
				_1584_v8 = _1584_v8
				_1575_v1 = _dafny.Companion_Sequence_.Concatenate(_1575_v1, _1575_v1)
			} else {
				_1629_cf44 = _dafny.IntOfInt64(483)
				var _1637_v39 _dafny.Array
				_ = _1637_v39
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_36
				var _nw244 _dafny.Array
				_ = _nw244
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw244 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Int = func(_1638_i8 _dafny.Int) _dafny.Int {
						return (_1638_i8).Times(_dafny.IntOfInt64(146))
					}
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw244 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw244).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw244).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1637_v39 = _nw244
				var _1639_v40 D16
				_ = _1639_v40
				_1639_v40 = Companion_D16_.Create_DC32_(_1637_v39)
				_1637_v39 = (_1639_v40).Dtor_cf51()
				_1579_v5 = _1628_cf45
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_1637_v39), 0))
				_ = _index232
				(_1637_v39).ArraySet1((_1583_v7).F12(), (_index232).Int())
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_1637_v39), 0))
				_ = _index233
				(_1637_v39).ArraySet1(_dafny.IntOfInt64(342), (_index233).Int())
				_1626_cf47 = (_1583_v7).F12()
			}
			var _1640_v41 _dafny.Set
			_ = _1640_v41
			_1640_v41 = _dafny.SetOf((_this).Fm7(globalState))
			var _1641_v42 _dafny.Map
			_ = _1641_v42
			_1641_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1628_cf45, _1628_cf45)
			var _1642_v43 _dafny.Map
			_ = _1642_v43
			_1642_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1577_v3).Cardinality()), (_1641_v42).Cardinality())
			var _rhs218 bool = ((_dafny.SetOf((func() _dafny.Int {
				if (_1642_v43).Contains((_this).F6()) {
					return (_1642_v43).Get((_this).F6()).(_dafny.Int)
				}
				return (_1583_v7).F12()
			})(), (_1583_v7).F12())).Difference(func() _dafny.Set {
				var _coll54 = _dafny.NewBuilder()
				_ = _coll54
				for _iter61 := _dafny.Iterate((_1576_v2).Elements()); ; {
					_compr_54, _ok61 := _iter61()
					if !_ok61 {
						break
					}
					var _1643_v44 _dafny.Int
					_1643_v44 = interface{}(_compr_54).(_dafny.Int)
					if (_1576_v2).Contains(_1643_v44) {
						_coll54.Add((_1643_v44).Plus((_dafny.SetOf(true, true, false, false)).Cardinality()))
					}
				}
				return _coll54.ToSet()
			}())).IsSubsetOf(_1640_v41)
			_ = _rhs218
			var _rhs219 _dafny.Int = ((_1629_cf44).Times((_this).F6())).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("reowmtl")).Cardinality()))
			_ = _rhs219
			_1579_v5 = _rhs218
			_1630_cf43 = _rhs219
			var _1644_v45 *C2
			_ = _1644_v45
			var _nw245 *C2 = New_C2_()
			_ = _nw245
			_nw245.Ctor__(_1575_v1, _1628_cf45)
			_1644_v45 = _nw245
		} else {
			var _1645___mcc_h5 _dafny.Sequence = _source29.Get_().(D12_DC27).Cf42
			_ = _1645___mcc_h5
			var _1646_cf42 _dafny.Sequence = _1645___mcc_h5
			_ = _1646_cf42
			var _1647_v46 D9
			_ = _1647_v46
			_1647_v46 = Companion_D9_.Create_DC19_(_1575_v1)
			var _1648_v47 _dafny.Array
			_ = _1648_v47
			var _nwElement0_56 D9 = Companion_D9_.Create_DC19_(Companion_Default___.Fm37((_this).F13(), false, _dafny.IntOfUint32((_1575_v1).Cardinality()), p0, globalState))
			_ = _nwElement0_56
			var _nw246 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(3))
			_ = _nw246
			(_nw246).ArraySet1(_nwElement0_56, 0)
			(_nw246).ArraySet1(_1647_v46, 1)
			(_nw246).ArraySet1(_1647_v46, 2)
			_1648_v47 = _nw246
			var _1649_v48 _dafny.Set
			_ = _1649_v48
			_1649_v48 = _dafny.SetOf(_1579_v5)
			var _pat_let_tv29 = _1580_v6
			_ = _pat_let_tv29
			var _pat_let_tv30 = _1649_v48
			_ = _pat_let_tv30
			var _pat_let_tv31 = globalState
			_ = _pat_let_tv31
			var _pat_let_tv32 = _1575_v1
			_ = _pat_let_tv32
			var _pat_let_tv33 = _1575_v1
			_ = _pat_let_tv33
			var _pat_let_tv34 = _1575_v1
			_ = _pat_let_tv34
			var _nwElement0_57 D9 = Companion_D9_.Create_DC19_(_1575_v1)
			_ = _nwElement0_57
			var _nw247 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(20))
			_ = _nw247
			(_nw247).ArraySet1(_nwElement0_57, 0)
			(_nw247).ArraySet1(_1647_v46, 1)
			(_nw247).ArraySet1(Companion_D9_.Create_DC19_(_1575_v1), 2)
			(_nw247).ArraySet1(func(_pat_let38_0 D9) D9 {
				return func(_1650_dt__update__tmp_h0 D9) D9 {
					return func(_pat_let39_0 _dafny.Sequence) D9 {
						return func(_1651_dt__update_hcf25_h0 _dafny.Sequence) D9 {
							return Companion_D9_.Create_DC19_(_1651_dt__update_hcf25_h0)
						}(_pat_let39_0)
					}(Companion_Default___.Fm50(_this.F14, _pat_let_tv29, (_pat_let_tv30).Cardinality(), !(!(!(false))), _pat_let_tv31))
				}(_pat_let38_0)
			}(_1647_v46), 3)
			(_nw247).ArraySet1(_1647_v46, 4)
			(_nw247).ArraySet1(Companion_D9_.Create_DC19_(_dafny.UnicodeSeqOfUtf8Bytes("ruvt")), 5)
			(_nw247).ArraySet1(_1647_v46, 6)
			(_nw247).ArraySet1(Companion_D9_.Create_DC19_(_1575_v1), 7)
			(_nw247).ArraySet1(_1647_v46, 8)
			(_nw247).ArraySet1((func() D9 {
				if p0 {
					return func(_pat_let40_0 D9) D9 {
						return func(_1654_dt__update__tmp_h2 D9) D9 {
							return func(_pat_let43_0 _dafny.Sequence) D9 {
								return func(_1655_dt__update_hcf25_h2 _dafny.Sequence) D9 {
									return Companion_D9_.Create_DC19_(_1655_dt__update_hcf25_h2)
								}(_pat_let43_0)
							}(_pat_let_tv33)
						}(_pat_let40_0)
					}(func(_pat_let41_0 D9) D9 {
						return func(_1652_dt__update__tmp_h1 D9) D9 {
							return func(_pat_let42_0 _dafny.Sequence) D9 {
								return func(_1653_dt__update_hcf25_h1 _dafny.Sequence) D9 {
									return Companion_D9_.Create_DC19_(_1653_dt__update_hcf25_h1)
								}(_pat_let42_0)
							}(_pat_let_tv32)
						}(_pat_let41_0)
					}(_1647_v46))
				}
				return _1647_v46
			})(), 9)
			(_nw247).ArraySet1(_1647_v46, 10)
			(_nw247).ArraySet1(_1647_v46, 11)
			(_nw247).ArraySet1(_1647_v46, 12)
			(_nw247).ArraySet1((func() D9 {
				if (_this).F13() {
					return Companion_D9_.Create_DC19_(_dafny.UnicodeSeqOfUtf8Bytes("khiwbj"))
				}
				return _1647_v46
			})(), 13)
			(_nw247).ArraySet1(Companion_D9_.Create_DC19_(_dafny.UnicodeSeqOfUtf8Bytes("qmqtcd")), 14)
			(_nw247).ArraySet1(_1647_v46, 15)
			(_nw247).ArraySet1(_1647_v46, 16)
			(_nw247).ArraySet1(Companion_D9_.Create_DC19_(_1575_v1), 17)
			(_nw247).ArraySet1(func(_pat_let44_0 D9) D9 {
				return func(_1656_dt__update__tmp_h3 D9) D9 {
					return func(_pat_let45_0 _dafny.Sequence) D9 {
						return func(_1657_dt__update_hcf25_h3 _dafny.Sequence) D9 {
							return Companion_D9_.Create_DC19_(_1657_dt__update_hcf25_h3)
						}(_pat_let45_0)
					}(_pat_let_tv34)
				}(_pat_let44_0)
			}(_1647_v46), 18)
			(_nw247).ArraySet1(_1647_v46, 19)
			_1648_v47 = _nw247
			var _1658_v49 _dafny.Array
			_ = _1658_v49
			var _nwElement0_58 bool = !((_this).F13())
			_ = _nwElement0_58
			var _nw248 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(9))
			_ = _nw248
			(_nw248).ArraySet1(_nwElement0_58, 0)
			(_nw248).ArraySet1(_1579_v5, 1)
			(_nw248).ArraySet1((Companion_Default___.Fm53(_1580_v6, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(22))).Uint32(), func(coer97 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg97 _dafny.Int) interface{} {
					return coer97(arg97)
				}
			}((func(_1659_v5 bool) func(_dafny.Int) _dafny.Map {
				return func(_1660_i9 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1659_v5, (_this).F13())
				}
			})(_1579_v5))), (Companion_Default___.SafeIndex((_1583_v7).F12(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(22))).Uint32(), func(coer98 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg98 _dafny.Int) interface{} {
					return coer98(arg98)
				}
			}((func(_1661_v5 bool) func(_dafny.Int) _dafny.Map {
				return func(_1662_i9 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1661_v5, (_this).F13())
				}
			})(_1579_v5)))).Cardinality()))).Uint32(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _1579_v5)).Update(p0, (_this).F13())), globalState)).IsSubsetOf(_1649_v48), 2)
			(_nw248).ArraySet1((func() bool {
				if p0 {
					return (_this).F4()
				}
				return Companion_Default___.Fm0(_1584_v8, globalState)
			})(), 3)
			(_nw248).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1618_v31, _1618_v31), 4)
			(_nw248).ArraySet1((_this).F13(), 5)
			(_nw248).ArraySet1((_this).F4(), 6)
			(_nw248).ArraySet1((_this).F13(), 7)
			(_nw248).ArraySet1(false, 8)
			_1658_v49 = _nw248
			var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_1658_v49), 0))
			_ = _index234
			(_1658_v49).ArraySet1(true, (_index234).Int())
			var _1663_v50 _dafny.Array
			_ = _1663_v50
			var _nw249 _dafny.Array = _dafny.NewArrayWithValue(Companion_D18_.Default(), _dafny.One)
			_ = _nw249
			_1663_v50 = _nw249
			var _1664_v51 D11
			_ = _1664_v51
			_1664_v51 = Companion_D11_.Create_DC26_(_1579_v5)
			var _1665_v52 _dafny.Map
			_ = _1665_v52
			_1665_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1664_v51, true)
			var _1666_v53 _dafny.Sequence
			_ = _1666_v53
			_1666_v53 = _dafny.SeqOf(_1665_v52)
			var _1667_v54 D18
			_ = _1667_v54
			_1667_v54 = Companion_D18_.Create_DC35_(_1666_v53)
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_1663_v50), 0))
			_ = _index235
			(_1663_v50).ArraySet1(_1667_v54, (_index235).Int())
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_1658_v49), 0))
			_ = _index236
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_1663_v50), 0))
			_ = _index237
			var _rhs220 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm37(p0, (_this).F13(), _1580_v6, (_this).F13(), globalState), _1575_v1), _1575_v1)
			_ = _rhs220
			var _rhs221 D18 = _1667_v54
			_ = _rhs221
			var _lhs138 _dafny.Array = _1658_v49
			_ = _lhs138
			var _lhs139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_1658_v49), 0))
			_ = _lhs139
			var _lhs140 _dafny.Array = _1663_v50
			_ = _lhs140
			var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_1663_v50), 0))
			_ = _lhs141
			(_lhs138).ArraySet1(_rhs220, (_lhs139).Int())
			(_lhs140).ArraySet1(_rhs221, (_lhs141).Int())
			var _1668_v55 D12
			_ = _1668_v55
			_1668_v55 = Companion_D12_.Create_DC28_((_this).F6(), (_1583_v7).F12(), p0, true, _1580_v6)
			var _rhs222 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F6(), (func() _dafny.Int {
				if !(_1579_v5) {
					return _1580_v6
				}
				return _1580_v6
			})())
			_ = _rhs222
			var _rhs223 bool = ((_1649_v48).Intersection(_1649_v48)).IsSubsetOf(_dafny.SetOf((_this).F13(), true, (_1668_v55).Dtor_cf45(), (_this).F4(), false))
			_ = _rhs223
			var _lhs142 *C5 = _this
			_ = _lhs142
			_lhs142.F14 = _rhs222
			_1579_v5 = _rhs223
			var _1669_v56 *C1
			_ = _1669_v56
			var _nw250 *C1 = New_C1_()
			_ = _nw250
			_nw250.Ctor__(_this.F14, _dafny.IntOfInt64(425), (func() _dafny.Array {
				if (_this).F13() {
					return _this.F5()
				}
				return _this.F5()
			})(), _this.F14, !(p0))
			_1669_v56 = _nw250
		}
		var _pat_let_tv35 = _1584_v8
		_ = _pat_let_tv35
		var _pat_let_tv36 = _1584_v8
		_ = _pat_let_tv36
		var _pat_let_tv37 = _1584_v8
		_ = _pat_let_tv37
		var _pat_let_tv38 = _1575_v1
		_ = _pat_let_tv38
		r0 = func(_source30 D11) _dafny.CodePoint {
			if _source30.Is_DC24() {
				var _1670___mcc_h6 bool = _source30.Get_().(D11_DC24).Cf36
				_ = _1670___mcc_h6
				var _1671___mcc_h7 bool = _source30.Get_().(D11_DC24).Cf37
				_ = _1671___mcc_h7
				var _1672___mcc_h8 _dafny.Int = _source30.Get_().(D11_DC24).Cf38
				_ = _1672___mcc_h8
				var _1673_cf38 _dafny.Int = _1672___mcc_h8
				_ = _1673_cf38
				var _1674_cf37 bool = _1671___mcc_h7
				_ = _1674_cf37
				var _1675_cf36 bool = _1670___mcc_h6
				_ = _1675_cf36
				return _dafny.CodePoint('w')
			} else if _source30.Is_DC25() {
				var _1676___mcc_h9 bool = _source30.Get_().(D11_DC25).Cf39
				_ = _1676___mcc_h9
				var _1677___mcc_h10 _dafny.Int = _source30.Get_().(D11_DC25).Cf40
				_ = _1677___mcc_h10
				var _1678_cf40 _dafny.Int = _1677___mcc_h10
				_ = _1678_cf40
				var _1679_cf39 bool = _1676___mcc_h9
				_ = _1679_cf39
				return _pat_let_tv35
			} else if _source30.Is_DC26() {
				var _1680___mcc_h11 bool = _source30.Get_().(D11_DC26).Cf41
				_ = _1680___mcc_h11
				var _1681_cf41 bool = _1680___mcc_h11
				_ = _1681_cf41
				return _pat_let_tv36
			} else {
				var _1682___mcc_h12 _dafny.Array = _source30.Get_().(D11_DC23).Cf35
				_ = _1682___mcc_h12
				var _1683_cf35 _dafny.Array = _1682___mcc_h12
				_ = _1683_cf35
				return _pat_let_tv37
			}
		}(func(_pat_let46_0 D11) D11 {
			return func(_1684_dt__update__tmp_h4 D11) D11 {
				return func(_pat_let47_0 _dafny.Int) D11 {
					return func(_1685_dt__update_hcf38_h0 _dafny.Int) D11 {
						return Companion_D11_.Create_DC24_((_1684_dt__update__tmp_h4).Dtor_cf36(), (_1684_dt__update__tmp_h4).Dtor_cf37(), _1685_dt__update_hcf38_h0)
					}(_pat_let47_0)
				}(_dafny.IntOfUint32((_pat_let_tv38).Cardinality()))
			}(_pat_let46_0)
		}(Companion_D11_.Create_DC24_(p0, !((_this).F4()), (_1583_v7).F12())))
		var _1686_v57 _dafny.Array
		_ = _1686_v57
		var _nw251 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(19))
		_ = _nw251
		_1686_v57 = _nw251
		r1 = _1686_v57
		return r0, r1
	}
}
func (_this *C5) F13() bool {
	{
		return _this._f13
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f5  _dafny.Array
	_f6  _dafny.Int
	_f4  bool
	_f11 bool
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.Zero
	_this._f4 = false
	_this._f11 = false
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C6{}
var _ T1 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F5() _dafny.Array {
	return _this._f5
}
func (_this *C6) F5_set_(value _dafny.Array) {
	_this._f5 = value
}
func (_this *C6) F6() _dafny.Int {
	return _this._f6
}
func (_this *C6) F4() bool {
	return _this._f4
}
func (_this *C6) Ctor__(f11 bool, f4 bool, f5 _dafny.Array, f6 _dafny.Int) {
	{
		(_this)._f11 = f11
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
	}
}
func (_this *C6) Fm6(p0 D0, p1 D0, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F6()))).Intersection(_dafny.MultiSetOf((_this).F6()))
	}
}
func (_this *C6) Fm7(globalState *GlobalState) _dafny.Int {
	{
		return (((_this).F6()).Times((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F6())).Cardinality(), (_this).F6())).Cardinality())).Times(Companion_Default___.SafeDivisionInt((_this).F6(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.UnicodeSeqOfUtf8Bytes("hkuc"))).Cardinality()))
	}
}
func (_this *C6) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return !(((_this).F11()) && (((_this).F6()).Cmp((_dafny.MultiSetOf((_this).F4(), (_this).F4())).Cardinality()) < 0))
	}
}
func (_this *C6) Fm9(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F6()
	}
}
func (_this *C6) Fm15(globalState *GlobalState) _dafny.Int {
	{
		var _source31 D0 = Companion_D0_.Create_DC1_()
		_ = _source31
		if _source31.Is_DC1() {
			return (_dafny.Zero).Minus((_this).F6())
		} else if _source31.Is_DC2() {
			var _1687___mcc_h0 _dafny.Int = _source31.Get_().(D0_DC2).Cf1
			_ = _1687___mcc_h0
			var _1688_cf1 _dafny.Int = _1687___mcc_h0
			_ = _1688_cf1
			return Companion_Default___.SafeDivisionInt(_1688_cf1, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F11())).Cardinality()))).Cardinality()))
		} else {
			var _1689___mcc_h1 _dafny.Int = _source31.Get_().(D0_DC0).Cf0
			_ = _1689___mcc_h1
			var _1690_cf0 _dafny.Int = _1689___mcc_h1
			_ = _1690_cf0
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F4(), (_this).F11()), (Companion_D3_.Create_DC8_(_dafny.SeqOf(false))).Dtor_cf7())).Cardinality())
		}
	}
}
func (_this *C6) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _1691_v0 _dafny.Sequence
		_ = _1691_v0
		_1691_v0 = _dafny.UnicodeSeqOfUtf8Bytes("gteox")
		var _1692_v1 _dafny.Map
		_ = _1692_v1
		_1692_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F11()), _1691_v0)
		var _1693_v2 D2
		_ = _1693_v2
		_1693_v2 = Companion_D2_.Create_DC6_()
		var _1694_v3 D1
		_ = _1694_v3
		_1694_v3 = Companion_D1_.Create_DC3_((_this).F11())
		var _1695_v4 _dafny.Array
		_ = _1695_v4
		var _nwElement0_59 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F11()), _1691_v0)
		_ = _nwElement0_59
		var _nw252 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(17))
		_ = _nw252
		(_nw252).ArraySet1(_nwElement0_59, 0)
		(_nw252).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1691_v0), 1)
		(_nw252).ArraySet1(_1692_v1, 2)
		(_nw252).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm16((_this).F4(), (_this).F4(), (_this).F11(), _1693_v2, globalState), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((Companion_Default___.Fm16((_this).F4(), (_this).F4(), (_this).F11(), _1693_v2, globalState)).Cardinality()))).Uint32(), _dafny.CodePoint('l'))), 3)
		(_nw252).ArraySet1(_1692_v1, 4)
		(_nw252).ArraySet1(_1692_v1, 5)
		(_nw252).ArraySet1(_1692_v1, 6)
		(_nw252).ArraySet1(_1692_v1, 7)
		(_nw252).ArraySet1((func() _dafny.Map {
			if (_1694_v3).Dtor_cf2() {
				return _1692_v1
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _dafny.UnicodeSeqOfUtf8Bytes("jwkwygjor"))
		})(), 8)
		(_nw252).ArraySet1((_1692_v1).Merge(_1692_v1), 9)
		(_nw252).ArraySet1((_1692_v1).Merge(_1692_v1), 10)
		(_nw252).ArraySet1(_1692_v1, 11)
		(_nw252).ArraySet1(_1692_v1, 12)
		(_nw252).ArraySet1(_1692_v1, 13)
		(_nw252).ArraySet1(_1692_v1, 14)
		(_nw252).ArraySet1(_1692_v1, 15)
		(_nw252).ArraySet1(_1692_v1, 16)
		_1695_v4 = _nw252
		var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(480), _dafny.ArrayLen((_1695_v4), 0))
		_ = _index238
		(_1695_v4).ArraySet1((_1692_v1).Merge(_1692_v1), (_index238).Int())
		var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(480), _dafny.ArrayLen((_1695_v4), 0))
		_ = _index239
		(_1695_v4).ArraySet1(_1692_v1, (_index239).Int())
		var _hi11 _dafny.Int = (_this).Fm7(globalState)
		_ = _hi11
		for _1696_i0 := p0; _1696_i0.Cmp(_hi11) < 0; _1696_i0 = _1696_i0.Plus(_dafny.One) {
			var _1697_v5 _dafny.Array
			_ = _1697_v5
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_37
			var _nw253 _dafny.Array
			_ = _nw253
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw253 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) _dafny.Int = (func(_1698_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1699_i1 _dafny.Int) _dafny.Int {
						return (_1699_i1).Minus(_1698_i0)
					}
				})(_1696_i0)
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw253 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw253).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw253).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1697_v5 = _nw253
			var _1700_v6 _dafny.Map
			_ = _1700_v6
			_1700_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F4())
			var _1701_v7 _dafny.Map
			_ = _1701_v7
			_1701_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1697_v5, _1700_v6)
			var _1702_v8 _dafny.Map
			_ = _1702_v8
			_1702_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _1696_i0)
			_1701_v7 = (_1701_v7).Update(_1697_v5, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_1702_v8).Contains((_this).F11()) {
					return (_1702_v8).Get((_this).F11()).(_dafny.Int)
				}
				return (_this).F6()
			})(), (_this).F4()))
			var _1703_v9 _dafny.MultiSet
			_ = _1703_v9
			_1703_v9 = _dafny.MultiSetOf(_1696_i0)
			if (_1703_v9).IsDisjointFrom(_1703_v9) {
				var _1704_v10 _dafny.Array
				_ = _1704_v10
				var _nw254 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
				_ = _nw254
				_1704_v10 = _nw254
				var _1705_v11 _dafny.Sequence
				_ = _1705_v11
				_1705_v11 = _dafny.SeqOf(_1697_v5)
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1704_v10), 0))
				_ = _index240
				(_1704_v10).ArraySet1(_1705_v11, (_index240).Int())
				var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_1704_v10), 0))
				_ = _index241
				(_1704_v10).ArraySet1(_1705_v11, (_index241).Int())
				var _1706_v12 _dafny.Int
				_ = _1706_v12
				_1706_v12 = _dafny.IntOfInt64(85)
				var _1707_v13 bool
				_ = _1707_v13
				_1707_v13 = false
				var _rhs224 _dafny.Int = (p0).Minus((_dafny.Zero).Minus((_this).F6()))
				_ = _rhs224
				var _rhs225 bool = !((_this).F4())
				_ = _rhs225
				_1706_v12 = _rhs224
				_1707_v13 = _rhs225
				var _1708_v14 _dafny.Array
				_ = _1708_v14
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_38
				var _nw255 _dafny.Array
				_ = _nw255
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw255 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) D0 = (func(_1709_v12 _dafny.Int) func(_dafny.Int) D0 {
						return func(_1710_i2 _dafny.Int) D0 {
							return Companion_D0_.Create_DC0_(_1709_v12)
						}
					})(_1706_v12)
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw255 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw255).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw255).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1708_v14 = _nw255
				_1708_v14 = _1708_v14
				var _1711_v15 _dafny.Set
				_ = _1711_v15
				_1711_v15 = _dafny.SetOf(_dafny.IntOfInt64(-882))
				_1706_v12 = ((_1711_v15).Cardinality()).Plus(p0)
				var _1712_v16 _dafny.Map
				_ = _1712_v16
				_1712_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1697_v5)
				var _1713_v17 _dafny.Map
				_ = _1713_v17
				_1713_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1712_v16, (func() _dafny.Int {
					if (_this).F4() {
						return (func() _dafny.Int {
							if (_1702_v8).Contains(false) {
								return (_1702_v8).Get(false).(_dafny.Int)
							}
							return (_this).Fm9(!(_1707_v13), (_this).F11(), p0, globalState)
						})()
					}
					return _1696_i0
				})())
				_1713_v17 = (_1713_v17).Update((_1712_v16).Update((_this).F4(), _1697_v5), _1696_i0)
			} else {
				var _1714_v18 *C0
				_ = _1714_v18
				var _nw256 *C0 = New_C0_()
				_ = _nw256
				_nw256.Ctor__(_1696_i0)
				_1714_v18 = _nw256
				var _1715_v19 _dafny.Int
				_ = _1715_v19
				_1715_v19 = _dafny.IntOfInt64(759)
				var _1716_v20 _dafny.Map
				_ = _1716_v20
				_1716_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(410))
				var _1717_v21 _dafny.Map
				_ = _1717_v21
				_1717_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1716_v20)
				_1715_v19 = (_1717_v21).Cardinality()
				_1715_v19 = ((_1714_v18).F12()).Times((_dafny.Zero).Minus(_1715_v19))
				var _1718_v22 bool
				_ = _1718_v22
				_1718_v22 = true
				var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_1697_v5), 0))
				_ = _index242
				(_1697_v5).ArraySet1((Companion_Default___.Fm3(_1715_v19, _1718_v22, globalState)).Plus(_1715_v19), (_index242).Int())
				var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_1697_v5), 0))
				_ = _index243
				(_1697_v5).ArraySet1((func() _dafny.Int {
					if (_1702_v8).Contains((_this).F4()) {
						return (_1702_v8).Get((_this).F4()).(_dafny.Int)
					}
					return _1696_i0
				})(), (_index243).Int())
				var _1719_v23 D2
				_ = _1719_v23
				_1719_v23 = Companion_D2_.Create_DC7_((_1714_v18).F12())
				var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_1697_v5), 0))
				_ = _index244
				var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_1697_v5), 0))
				_ = _index245
				var _rhs226 bool = _1718_v22
				_ = _rhs226
				var _rhs227 _dafny.Int = (_dafny.Zero).Minus(((_1714_v18).F12()).Minus((_dafny.IntOfUint32((Companion_Default___.Fm16((_this).F4(), _1718_v22, (_this).F11(), _1693_v2, globalState)).Cardinality())).Plus(p0)))
				_ = _rhs227
				var _rhs228 _dafny.Int = (p0).Times((_1719_v23).Dtor_cf6())
				_ = _rhs228
				var _rhs229 _dafny.Int = (_this).F6()
				_ = _rhs229
				var _rhs230 bool = (_1694_v3).Dtor_cf2()
				_ = _rhs230
				var _lhs143 _dafny.Array = _1697_v5
				_ = _lhs143
				var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_1697_v5), 0))
				_ = _lhs144
				var _lhs145 _dafny.Array = _1697_v5
				_ = _lhs145
				var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_1697_v5), 0))
				_ = _lhs146
				_1718_v22 = _rhs226
				(_lhs143).ArraySet1(_rhs227, (_lhs144).Int())
				(_lhs145).ArraySet1(_rhs228, (_lhs146).Int())
				_1715_v19 = _rhs229
				_1718_v22 = _rhs230
				_1715_v19 = (_1703_v9).Cardinality()
			}
			var _1720_v24 _dafny.Set
			_ = _1720_v24
			_1720_v24 = _dafny.SetOf((_this).F4())
			var _1721_v25 _dafny.Map
			_ = _1721_v25
			_1721_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1720_v24, (_this).F11())
			var _1722_v26 _dafny.Set
			_ = _1722_v26
			_1722_v26 = _dafny.SetOf((_1721_v25).Cardinality(), p0, p0, p0)
			_1722_v26 = _1722_v26
			var _1723_v27 _dafny.Int
			_ = _1723_v27
			_1723_v27 = _dafny.IntOfInt64(-271)
			_1723_v27 = _1723_v27
		}
		var _1724_i3 _dafny.Int
		_ = _1724_i3
		_1724_i3 = _dafny.Zero
		{
			for (_this).F11() {
				{
					if (_1724_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1724_i3 = (_1724_i3).Plus(_dafny.One)
					var _1725_v28 _dafny.Array
					_ = _1725_v28
					var _nw257 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
					_ = _nw257
					_1725_v28 = _nw257
					_1725_v28 = _1725_v28
					var _1726_v29 bool
					_ = _1726_v29
					_1726_v29 = false
					var _1727_v30 _dafny.Sequence
					_ = _1727_v30
					_1727_v30 = _dafny.SeqOf(_dafny.IntOfInt64(583))
					_1726_v29 = _dafny.Companion_Sequence_.Equal(_1727_v30, _1727_v30)
					_1726_v29 = (_this).F11()
					var _1728_v31 _dafny.Int
					_ = _1728_v31
					_1728_v31 = _dafny.IntOfInt64(-981)
					var _rhs231 bool = (_this).F11()
					_ = _rhs231
					var _rhs232 _dafny.Int = _dafny.IntOfInt64(759)
					_ = _rhs232
					_1726_v29 = _rhs231
					_1728_v31 = _rhs232
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1729_v32 _dafny.MultiSet
		_ = _1729_v32
		_1729_v32 = _dafny.MultiSetOf((_this).F11())
		var _1730_v33 _dafny.Map
		_ = _1730_v33
		_1730_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F6()), (_this).F4())
		if (((_this).F6()).Plus((_this).F6())).Cmp((p0).Times((func() _dafny.Int {
			if (_1729_v32).Contains((_this).F4()) {
				return (_1729_v32).Multiplicity((_this).F4())
			}
			return (_1730_v33).Cardinality()
		})())) == 0 {
			var _1731_v34 _dafny.Sequence
			_ = _1731_v34
			_1731_v34 = _dafny.SeqOf(_dafny.IntOfInt64(327), p0)
			var _1732_v35 *C0
			_ = _1732_v35
			var _nw258 *C0 = New_C0_()
			_ = _nw258
			_nw258.Ctor__(_dafny.IntOfUint32((_1731_v34).Cardinality()))
			_1732_v35 = _nw258
			var _1733_v36 _dafny.Sequence
			_ = _1733_v36
			_1733_v36 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("wfhdckno"))
			var _1734_v37 *C0
			_ = _1734_v37
			var _nw259 *C0 = New_C0_()
			_ = _nw259
			_nw259.Ctor__(_dafny.IntOfUint32((_1733_v36).Cardinality()))
			_1734_v37 = _nw259
			var _1735_v38 _dafny.Array
			_ = _1735_v38
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_39
			var _nw260 _dafny.Array
			_ = _nw260
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw260 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.Sequence = (func(_1736_v0 _dafny.Sequence, _1737_v34 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1738_i4 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1736_v0, _1736_v0), (Companion_Default___.SafeIndex((_1737_v34).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.IntOfUint32((_1737_v34).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1736_v0, _1736_v0)).Cardinality()))).Uint32(), _dafny.CodePoint('j'))
					}
				})(_1691_v0, _1731_v34)
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw260 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw260).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw260).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1735_v38 = _nw260
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_1735_v38), 0))
			_ = _index246
			(_1735_v38).ArraySet1(_1691_v0, (_index246).Int())
			var _1739_v39 _dafny.CodePoint
			_ = _1739_v39
			_1739_v39 = _dafny.CodePoint('s')
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_1735_v38), 0))
			_ = _index247
			(_1735_v38).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1691_v0, _dafny.Companion_Sequence_.Concatenate(_1691_v0, _dafny.Companion_Sequence_.Update(_1691_v0, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1691_v0).Cardinality()))).Uint32(), _1739_v39))), (_index247).Int())
			var _1740_v40 _dafny.Set
			_ = _1740_v40
			_1740_v40 = _dafny.SetOf(p0)
			var _1741_v41 _dafny.Array
			_ = _1741_v41
			var _nw261 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw261
			_1741_v41 = _nw261
			var _1742_v42 _dafny.Map
			_ = _1742_v42
			_1742_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1740_v40).Intersection(_1740_v40)).Cardinality(), _1741_v41)
			_1742_v42 = (_1742_v42).Update(_dafny.IntOfInt64(-824), _1741_v41)
			var _1743_v43 _dafny.Map
			_ = _1743_v43
			_1743_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1691_v0, (_1734_v37).F12())
			var _1744_v44 _dafny.Array
			_ = _1744_v44
			var _nwElement0_60 bool = (_this).F4()
			_ = _nwElement0_60
			var _nw262 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(6))
			_ = _nw262
			(_nw262).ArraySet1(_nwElement0_60, 0)
			(_nw262).ArraySet1((_this).F11(), 1)
			(_nw262).ArraySet1((_this).F11(), 2)
			(_nw262).ArraySet1((_this).F4(), 3)
			(_nw262).ArraySet1((_this).F4(), 4)
			(_nw262).ArraySet1((_this).F11(), 5)
			_1744_v44 = _nw262
			var _1745_v45 _dafny.Map
			_ = _1745_v45
			_1745_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1744_v44, (_this).F11())
			_1743_v43 = (_1743_v43).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(863))).Uint32(), func(coer99 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg99 _dafny.Int) interface{} {
					return coer99(arg99)
				}
			}((func(_1746_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1747_i5 _dafny.Int) _dafny.CodePoint {
					return _1746_v39
				}
			})(_1739_v39))), (Companion_Default___.Fm18(_1691_v0, (_1745_v45).Cardinality(), globalState)).Cardinality())
		} else {
			var _1748_v46 _dafny.Int
			_ = _1748_v46
			_1748_v46 = _dafny.IntOfInt64(602)
			var _1749_v47 _dafny.CodePoint
			_ = _1749_v47
			_1749_v47 = _dafny.CodePoint('j')
			var _rhs233 _dafny.Int = p0
			_ = _rhs233
			var _rhs234 _dafny.Int = p0
			_ = _rhs234
			var _rhs235 _dafny.CodePoint = _1749_v47
			_ = _rhs235
			_1748_v46 = _rhs233
			_1748_v46 = _rhs234
			_1749_v47 = _rhs235
			if (_this).F4() {
				_1748_v46 = p0
				_1748_v46 = (func() _dafny.Int {
					if (_1748_v46).Cmp(p0) > 0 {
						return _dafny.IntOfUint32((Companion_Default___.Fm16((_this).F11(), (_this).F4(), (Companion_D3_.Create_DC9_((_this).F4(), _1748_v46, (_this).F11())).Dtor_cf10(), _1693_v2, globalState)).Cardinality())
					}
					return _1748_v46
				})()
				var _1750_v48 _dafny.Set
				_ = _1750_v48
				_1750_v48 = _dafny.SetOf(_dafny.IntOfInt64(274), p0, _dafny.IntOfUint32((_1691_v0).Cardinality()))
				_1750_v48 = (_1750_v48).Union(_dafny.SetOf(p0))
				var _1751_v49 _dafny.Array
				_ = _1751_v49
				var _len0_40 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_40
				var _nw263 _dafny.Array
				_ = _nw263
				if _len0_40.Cmp(_dafny.Zero) == 0 {
					_nw263 = _dafny.NewArray(_len0_40)
				} else {
					var _init40 func(_dafny.Int) _dafny.Int = (func(_1752_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1753_i6 _dafny.Int) _dafny.Int {
							return (_1753_i6).Times(_1752_p0)
						}
					})(p0)
					_ = _init40
					var _element0_40 = _init40(_dafny.Zero)
					_ = _element0_40
					_nw263 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
					(_nw263).ArraySet1(_element0_40, 0)
					var _nativeLen0_40 = (_len0_40).Int()
					_ = _nativeLen0_40
					for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
						(_nw263).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
					}
				}
				_1751_v49 = _nw263
				var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_1751_v49), 0))
				_ = _index248
				(_1751_v49).ArraySet1((func() _dafny.Int {
					if (_this).F4() {
						return (_this).F6()
					}
					return p0
				})(), (_index248).Int())
				var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_1751_v49), 0))
				_ = _index249
				(_1751_v49).ArraySet1(_dafny.IntOfInt64(348), (_index249).Int())
				var _1754_v50 *C0
				_ = _1754_v50
				var _nw264 *C0 = New_C0_()
				_ = _nw264
				_nw264.Ctor__((func() _dafny.Int {
					if (_this).F11() {
						return _dafny.IntOfInt64(-268)
					}
					return p0
				})())
				_1754_v50 = _nw264
			} else {
				var _1755_v51 D4
				_ = _1755_v51
				_1755_v51 = Companion_D4_.Create_DC12_(_1749_v47)
				_1749_v47 = (_1755_v51).Dtor_cf13()
				var _1756_v52 _dafny.Sequence
				_ = _1756_v52
				_1756_v52 = _dafny.SeqOf((_dafny.IntOfUint32((Companion_Default___.Fm16((_this).F11(), Companion_Default___.Fm0(_1749_v47, globalState), (_this).F11(), _1693_v2, globalState)).Cardinality())).Times((_this).Fm15(globalState)))
				var _1757_v53 _dafny.Sequence
				_ = _1757_v53
				_1757_v53 = _dafny.SeqOf((_this).F4())
				var _1758_v54 _dafny.Array
				_ = _1758_v54
				var _nw265 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw265
				_1758_v54 = _nw265
				var _1759_v55 _dafny.Array
				_ = _1759_v55
				var _nwElement0_61 _dafny.Sequence = _1757_v53
				_ = _nwElement0_61
				var _nw266 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(19))
				_ = _nw266
				(_nw266).ArraySet1(_nwElement0_61, 0)
				(_nw266).ArraySet1(_dafny.SeqOf((_this).F4()), 1)
				(_nw266).ArraySet1(Companion_Default___.Fm19((_this).F11(), Companion_Default___.Fm3((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1749_v47, _1758_v54)).Cardinality()), (_this).F11(), globalState), globalState), 2)
				(_nw266).ArraySet1(_1757_v53, 3)
				(_nw266).ArraySet1(_1757_v53, 4)
				(_nw266).ArraySet1(_1757_v53, 5)
				(_nw266).ArraySet1(_1757_v53, 6)
				(_nw266).ArraySet1(_1757_v53, 7)
				(_nw266).ArraySet1((func() _dafny.Sequence {
					if (_this).F4() {
						return _1757_v53
					}
					return _1757_v53
				})(), 8)
				(_nw266).ArraySet1(_1757_v53, 9)
				(_nw266).ArraySet1(_1757_v53, 10)
				(_nw266).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1757_v53, (Companion_Default___.SafeIndex(_1748_v46, _dafny.IntOfUint32((_1757_v53).Cardinality()))).Uint32(), true), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1757_v53, (Companion_Default___.SafeIndex(_1748_v46, _dafny.IntOfUint32((_1757_v53).Cardinality()))).Uint32(), true)).Cardinality()))).Uint32(), false), 11)
				(_nw266).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1757_v53, _1757_v53), 12)
				(_nw266).ArraySet1(_1757_v53, 13)
				(_nw266).ArraySet1(_1757_v53, 14)
				(_nw266).ArraySet1(_1757_v53, 15)
				(_nw266).ArraySet1(_1757_v53, 16)
				(_nw266).ArraySet1(_1757_v53, 17)
				(_nw266).ArraySet1(_1757_v53, 18)
				_1759_v55 = _nw266
				var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_1759_v55), 0))
				_ = _index250
				(_1759_v55).ArraySet1(_dafny.Companion_Sequence_.Concatenate((Companion_D4_.Create_DC13_((_this).F6(), (_this).F11(), _1748_v46, _1757_v53)).Dtor_cf17(), _1757_v53), (_index250).Int())
				var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_1759_v55), 0))
				_ = _index251
				var _rhs236 _dafny.Sequence = _1756_v52
				_ = _rhs236
				var _rhs237 _dafny.Int = p0
				_ = _rhs237
				var _rhs238 _dafny.Int = _dafny.IntOfInt64(-794)
				_ = _rhs238
				var _rhs239 _dafny.Sequence = _1757_v53
				_ = _rhs239
				var _lhs147 _dafny.Array = _1759_v55
				_ = _lhs147
				var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_1759_v55), 0))
				_ = _lhs148
				_1756_v52 = _rhs236
				_1748_v46 = _rhs237
				_1748_v46 = _rhs238
				(_lhs147).ArraySet1(_rhs239, (_lhs148).Int())
				var _1760_v56 D3
				_ = _1760_v56
				_1760_v56 = Companion_D3_.Create_DC10_(p0)
				var _1761_v57 D3
				_ = _1761_v57
				_1761_v57 = Companion_D3_.Create_DC11_(_1760_v56)
				var _1762_v58 _dafny.Array
				_ = _1762_v58
				var _nwElement0_62 bool = (_this).F11()
				_ = _nwElement0_62
				var _nw267 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(6))
				_ = _nw267
				(_nw267).ArraySet1(_nwElement0_62, 0)
				(_nw267).ArraySet1((_this).F4(), 1)
				(_nw267).ArraySet1((_this).F4(), 2)
				(_nw267).ArraySet1((_this).F4(), 3)
				(_nw267).ArraySet1((_this).F4(), 4)
				(_nw267).ArraySet1((_this).F4(), 5)
				_1762_v58 = _nw267
				var _1763_v59 _dafny.MultiSet
				_ = _1763_v59
				_1763_v59 = _dafny.MultiSetOf(_1762_v58)
				var _1764_v60 _dafny.Map
				_ = _1764_v60
				_1764_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1761_v57, (_1763_v59).Difference(_1763_v59))
				_1764_v60 = (_1764_v60).Update(_1761_v57, _1763_v59)
				var _1765_v61 _dafny.MultiSet
				_ = _1765_v61
				_1765_v61 = _dafny.MultiSetOf((func() _dafny.Int {
					if (_this).F4() {
						return _dafny.IntOfInt64(-133)
					}
					return p0
				})(), _1748_v46, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-336))).Uint32(), func(coer100 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}(func(_1766_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				}))).Cardinality()), p0)
				_1765_v61 = (_dafny.MultiSetFromSeq(_1756_v52))
				var _1767_v62 _dafny.Map
				_ = _1767_v62
				_1767_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _1768_v63 _dafny.Map
				_ = _1768_v63
				_1768_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1765_v61).Cardinality(), _1691_v0)
				var _1769_v64 _dafny.Sequence
				_ = _1769_v64
				_1769_v64 = _dafny.SeqOf(_1768_v63, _1768_v63)
				var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_1758_v54), 0))
				_ = _index252
				(_1758_v54).ArraySet1((func() _dafny.Int {
					if (_this).F4() {
						return Companion_Default___.Fm3((func() _dafny.Int {
							if (_1767_v62).Contains(_1748_v46) {
								return (_1767_v62).Get(_1748_v46).(_dafny.Int)
							}
							return _1748_v46
						})(), (_this).F4(), globalState)
					}
					return ((_1769_v64).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1691_v0).Cardinality()), _dafny.IntOfUint32((_1769_v64).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()
				})(), (_index252).Int())
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_1758_v54), 0))
				_ = _index253
				(_1758_v54).ArraySet1((func() _dafny.Int {
					if (_this).F11() {
						return p0
					}
					return _dafny.IntOfUint32(((Companion_D3_.Create_DC8_((Companion_Default___.Fm20((_this).F6(), (_this).F4(), globalState)).Dtor_cf7())).Dtor_cf7()).Cardinality())
				})(), (_index253).Int())
			}
			var _1770_v65 _dafny.Array
			_ = _1770_v65
			var _len0_41 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_41
			var _nw268 _dafny.Array
			_ = _nw268
			if _len0_41.Cmp(_dafny.Zero) == 0 {
				_nw268 = _dafny.NewArray(_len0_41)
			} else {
				var _init41 func(_dafny.Int) bool = (func(_1771_v46 _dafny.Int, _1772_v0 _dafny.Sequence, _1773_p0 _dafny.Int) func(_dafny.Int) bool {
					return func(_1774_i8 _dafny.Int) bool {
						return (_dafny.MultiSetOf(_1771_v46, _dafny.IntOfUint32((_1772_v0).Cardinality()))).IsProperSubsetOf(_dafny.MultiSetOf(_1773_p0, _1771_v46, _1771_v46))
					}
				})(_1748_v46, _1691_v0, p0)
				_ = _init41
				var _element0_41 = _init41(_dafny.Zero)
				_ = _element0_41
				_nw268 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
				(_nw268).ArraySet1(_element0_41, 0)
				var _nativeLen0_41 = (_len0_41).Int()
				_ = _nativeLen0_41
				for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
					(_nw268).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
				}
			}
			_1770_v65 = _nw268
			_1770_v65 = _1770_v65
			_1730_v33 = (_1730_v33).Update(_dafny.IntOfInt64(451), (_this).F4())
			var _1775_v66 bool
			_ = _1775_v66
			_1775_v66 = true
			var _1776_v67 _dafny.Map
			_ = _1776_v67
			_1776_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F6())
			_1775_v66 = (_dafny.MultiSetOf((_this).F6())).Contains((_1776_v67).Cardinality())
		}
		var _1777_v68 bool
		_ = _1777_v68
		_1777_v68 = true
		_1777_v68 = (_this).F4()
		var _1778_v69 _dafny.Int
		_ = _1778_v69
		_1778_v69 = _dafny.IntOfInt64(678)
		_1778_v69 = (_dafny.Zero).Minus((func() _dafny.Int {
			if !((_1777_v68) == (_1777_v68)) {
				return (_this).F6()
			}
			return p0
		})())
		var _1779_v70 _dafny.Map
		_ = _1779_v70
		_1779_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1777_v68, p0)
		r0 = _1779_v70
		return r0
	}
}
func (_this *C6) M2(p0 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Array) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _1780_v0 _dafny.Sequence
		_ = _1780_v0
		_1780_v0 = _dafny.SeqOf((_this).F11(), p0, (_this).F4())
		var _1781_v1 _dafny.Sequence
		_ = _1781_v1
		_1781_v1 = _dafny.SeqOf((_1780_v0).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1780_v0).Cardinality()))).Uint32()).(bool), p0, (_this).F4())
		var _1782_v2 D3
		_ = _1782_v2
		_1782_v2 = Companion_D3_.Create_DC8_(_dafny.Companion_Sequence_.Update(_1780_v0, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1780_v0).Cardinality()))).Uint32(), (_this).F11()))
		var _source32 D3 = _1782_v2
		_ = _source32
		if _source32.Is_DC9() {
			var _1783___mcc_h0 bool = _source32.Get_().(D3_DC9).Cf8
			_ = _1783___mcc_h0
			var _1784___mcc_h1 _dafny.Int = _source32.Get_().(D3_DC9).Cf9
			_ = _1784___mcc_h1
			var _1785___mcc_h2 bool = _source32.Get_().(D3_DC9).Cf10
			_ = _1785___mcc_h2
			var _1786_cf10 bool = _1785___mcc_h2
			_ = _1786_cf10
			var _1787_cf9 _dafny.Int = _1784___mcc_h1
			_ = _1787_cf9
			var _1788_cf8 bool = _1783___mcc_h0
			_ = _1788_cf8
			var _1789_v3 _dafny.Array
			_ = _1789_v3
			var _nw269 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw269
			_1789_v3 = _nw269
			var _1790_v4 _dafny.Array
			_ = _1790_v4
			var _len0_42 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_42
			var _nw270 _dafny.Array
			_ = _nw270
			if _len0_42.Cmp(_dafny.Zero) == 0 {
				_nw270 = _dafny.NewArray(_len0_42)
			} else {
				var _init42 func(_dafny.Int) bool = func(_1791_i0 _dafny.Int) bool {
					return !((_this).F11())
				}
				_ = _init42
				var _element0_42 = _init42(_dafny.Zero)
				_ = _element0_42
				_nw270 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
				(_nw270).ArraySet1(_element0_42, 0)
				var _nativeLen0_42 = (_len0_42).Int()
				_ = _nativeLen0_42
				for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
					(_nw270).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
				}
			}
			_1790_v4 = _nw270
			var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_1790_v4), 0))
			_ = _index254
			(_1790_v4).ArraySet1((_this).F11(), (_index254).Int())
			var _1792_v5 _dafny.Sequence
			_ = _1792_v5
			_1792_v5 = _dafny.SeqOf(_1789_v3, _1789_v3, _1789_v3, _1789_v3, _1789_v3)
			var _1793_v6 _dafny.MultiSet
			_ = _1793_v6
			_1793_v6 = _dafny.MultiSetOf(_1790_v4, _1790_v4)
			var _1794_v7 _dafny.Sequence
			_ = _1794_v7
			_1794_v7 = _dafny.UnicodeSeqOfUtf8Bytes("rmjgkq")
			var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_1790_v4), 0))
			_ = _index255
			var _rhs240 _dafny.Int = (_this).Fm9(_1786_cf10, (_this).F11(), _1787_cf9, globalState)
			_ = _rhs240
			var _rhs241 _dafny.Array = (_1792_v5).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1793_v6).Contains(_1790_v4) {
					return (_1793_v6).Multiplicity(_1790_v4)
				}
				return _dafny.IntOfUint32((_1794_v7).Cardinality())
			})(), _dafny.IntOfUint32((_1792_v5).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs241
			var _rhs242 _dafny.Int = (_dafny.Zero).Minus((_dafny.IntOfUint32((_1794_v7).Cardinality())).Plus(Companion_Default___.SafeModuloInt((_this).Fm9(true, (_this).F11(), _1787_cf9, globalState), (_this).F6())))
			_ = _rhs242
			var _rhs243 _dafny.Int = (func() _dafny.Int {
				if p0 {
					return _dafny.IntOfUint32((_1794_v7).Cardinality())
				}
				return _1787_cf9
			})()
			_ = _rhs243
			var _rhs244 bool = _1786_cf10
			_ = _rhs244
			var _lhs149 _dafny.Array = _1790_v4
			_ = _lhs149
			var _lhs150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_1790_v4), 0))
			_ = _lhs150
			_1787_cf9 = _rhs240
			_1789_v3 = _rhs241
			_1787_cf9 = _rhs242
			_1787_cf9 = _rhs243
			(_lhs149).ArraySet1(_rhs244, (_lhs150).Int())
			if (_this).F4() {
				_1786_cf10 = false
				var _arr7 _dafny.Array = _this.F5()
				_ = _arr7
				var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_this.F5()), 0))
				_ = _index256
				(_arr7).ArraySet1(_1790_v4, (_index256).Int())
				var _arr8 _dafny.Array = _this.F5()
				_ = _arr8
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_this.F5()), 0))
				_ = _index257
				(_arr8).ArraySet1((Companion_D2_.Create_DC5_(_1790_v4)).Dtor_cf5(), (_index257).Int())
				var _1795_v8 D3
				_ = _1795_v8
				_1795_v8 = Companion_D3_.Create_DC10_(_dafny.IntOfUint32((_1794_v7).Cardinality()))
				var _1796_v9 _dafny.Array
				_ = _1796_v9
				var _nwElement0_63 D3 = (func() D3 {
					if _1786_cf10 {
						return _1795_v8
					}
					return _1795_v8
				})()
				_ = _nwElement0_63
				var _nw271 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(25))
				_ = _nw271
				(_nw271).ArraySet1(_nwElement0_63, 0)
				(_nw271).ArraySet1(_1795_v8, 1)
				(_nw271).ArraySet1(_1795_v8, 2)
				(_nw271).ArraySet1(_1795_v8, 3)
				(_nw271).ArraySet1(_1795_v8, 4)
				(_nw271).ArraySet1(_1795_v8, 5)
				(_nw271).ArraySet1(_1795_v8, 6)
				(_nw271).ArraySet1(_1795_v8, 7)
				(_nw271).ArraySet1(_1795_v8, 8)
				(_nw271).ArraySet1(_1795_v8, 9)
				(_nw271).ArraySet1(_1795_v8, 10)
				(_nw271).ArraySet1(Companion_D3_.Create_DC10_(_1787_cf9), 11)
				(_nw271).ArraySet1(_1795_v8, 12)
				(_nw271).ArraySet1(_1795_v8, 13)
				(_nw271).ArraySet1(Companion_D3_.Create_DC10_((_this).F6()), 14)
				(_nw271).ArraySet1(_1795_v8, 15)
				(_nw271).ArraySet1(Companion_D3_.Create_DC10_((_this).F6()), 16)
				(_nw271).ArraySet1(_1795_v8, 17)
				(_nw271).ArraySet1((func() D3 {
					if _1786_cf10 {
						return _1795_v8
					}
					return _1795_v8
				})(), 18)
				(_nw271).ArraySet1(_1795_v8, 19)
				(_nw271).ArraySet1(_1795_v8, 20)
				(_nw271).ArraySet1(_1795_v8, 21)
				(_nw271).ArraySet1(_1795_v8, 22)
				(_nw271).ArraySet1(_1795_v8, 23)
				(_nw271).ArraySet1(_1795_v8, 24)
				_1796_v9 = _nw271
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1796_v9), 0))
				_ = _index258
				(_1796_v9).ArraySet1(_1795_v8, (_index258).Int())
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1796_v9), 0))
				_ = _index259
				(_1796_v9).ArraySet1(_1795_v8, (_index259).Int())
				_1787_cf9 = Companion_Default___.SafeModuloInt(_1787_cf9, (_this).F6())
				var _1797_v10 _dafny.Map
				_ = _1797_v10
				_1797_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_1790_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_1790_v4), 0))).Int()).(bool))
				var _1798_v11 _dafny.Sequence
				_ = _1798_v11
				_1798_v11 = _dafny.SeqOf((_1797_v10).Update(_1786_cf10, _1788_cf8))
				_1797_v10 = ((_1798_v11).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1798_v11).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_1797_v10)
			} else {
				var _1799_v12 _dafny.Map
				_ = _1799_v12
				_1799_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1781_v1).Select((Companion_Default___.SafeIndex(_1787_cf9, _dafny.IntOfUint32((_1781_v1).Cardinality()))).Uint32()).(bool)), (_this).F6())
				var _1800_v13 *C0
				_ = _1800_v13
				var _nw272 *C0 = New_C0_()
				_ = _nw272
				_nw272.Ctor__((func() _dafny.Int {
					if (_1799_v12).Contains(_1786_cf10) {
						return (_1799_v12).Get(_1786_cf10).(_dafny.Int)
					}
					return (_this).F6()
				})())
				_1800_v13 = _nw272
				var _1801_v14 D2
				_ = _1801_v14
				_1801_v14 = Companion_D2_.Create_DC7_(Companion_Default___.SafeModuloInt((_this).F6(), _dafny.IntOfInt64(433)))
				var _1802_v15 _dafny.Set
				_ = _1802_v15
				_1802_v15 = _dafny.SetOf((_1800_v13).F12())
				var _1803_v16 _dafny.MultiSet
				_ = _1803_v16
				_1803_v16 = _dafny.MultiSetOf(_1802_v15, _1802_v15, _1802_v15)
				var _1804_v17 _dafny.Map
				_ = _1804_v17
				_1804_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1800_v13, _1790_v4)
				var _1805_v18 _dafny.Map
				_ = _1805_v18
				_1805_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(376), (_1804_v17).Cardinality())
				_1801_v14 = (func() D2 {
					if !((_1803_v16).Update(_1802_v15, Companion_Default___.Abs((_1800_v13).F12()))).Contains(_dafny.SetOf((_this).F6())) {
						return Companion_D2_.Create_DC7_((_1805_v18).Cardinality())
					}
					return _1801_v14
				})()
				var _1806_v19 _dafny.CodePoint
				_ = _1806_v19
				_1806_v19 = _dafny.CodePoint('s')
				r0 = _1806_v19
				_1786_cf10 = (_this).F4()
				_1787_cf9 = _1787_cf9
			}
			var _1807_v20 *C0
			_ = _1807_v20
			var _nw273 *C0 = New_C0_()
			_ = _nw273
			_nw273.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1780_v0).Cardinality()), (_this).F6()))
			_1807_v20 = _nw273
			var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_1790_v4), 0))
			_ = _index260
			(_1790_v4).ArraySet1((_1807_v20).Fm17(p0, globalState), (_index260).Int())
		} else if _source32.Is_DC10() {
			var _1808___mcc_h3 _dafny.Int = _source32.Get_().(D3_DC10).Cf11
			_ = _1808___mcc_h3
			var _1809_cf11 _dafny.Int = _1808___mcc_h3
			_ = _1809_cf11
			var _1810_v21 _dafny.Map
			_ = _1810_v21
			_1810_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(137), _1809_cf11), (func() bool {
				if (_this).F11() {
					return (_this).F11()
				}
				return p0
			})())
			_1810_v21 = (_1810_v21).Update(_1809_cf11, (_this).F11())
			var _1811_v22 bool
			_ = _1811_v22
			_1811_v22 = false
			var _1812_v23 _dafny.Array
			_ = _1812_v23
			var _nw274 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw274
			_1812_v23 = _nw274
			var _1813_v24 _dafny.Sequence
			_ = _1813_v24
			_1813_v24 = _dafny.SeqOf(_1809_cf11)
			var _rhs245 bool = p0
			_ = _rhs245
			var _rhs246 _dafny.Int = _1809_cf11
			_ = _rhs246
			var _rhs247 bool = !_dafny.Companion_Sequence_.Equal(_1813_v24, _1813_v24)
			_ = _rhs247
			var _rhs248 _dafny.Array = _1812_v23
			_ = _rhs248
			_1811_v22 = _rhs245
			_1809_cf11 = _rhs246
			_1811_v22 = _rhs247
			_1812_v23 = _rhs248
			var _1814_v25 _dafny.Sequence
			_ = _1814_v25
			_1814_v25 = _dafny.UnicodeSeqOfUtf8Bytes("gan")
			var _1815_v26 _dafny.CodePoint
			_ = _1815_v26
			_1815_v26 = _dafny.CodePoint('j')
			_1814_v25 = (func() _dafny.Sequence {
				if (func() bool {
					if (_this).F4() {
						return Companion_Default___.Fm0(_1815_v26, globalState)
					}
					return !((_this).F4())
				})() {
					return _1814_v25
				}
				return _dafny.Companion_Sequence_.Concatenate(_1814_v25, _1814_v25)
			})()
			var _1816_v27 _dafny.Array
			_ = _1816_v27
			var _nw275 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
			_ = _nw275
			_1816_v27 = _nw275
			_1816_v27 = _1816_v27
		} else if _source32.Is_DC8() {
			var _1817___mcc_h4 _dafny.Sequence = _source32.Get_().(D3_DC8).Cf7
			_ = _1817___mcc_h4
			var _1818_cf7 _dafny.Sequence = _1817___mcc_h4
			_ = _1818_cf7
			var _1819_v28 _dafny.Array
			_ = _1819_v28
			var _nw276 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw276
			_1819_v28 = _nw276
			var _1820_v29 D1
			_ = _1820_v29
			_1820_v29 = Companion_D1_.Create_DC4_((_dafny.Zero).Minus((_this).F6()), (_this).F6())
			var _1821_v30 _dafny.Sequence
			_ = _1821_v30
			_1821_v30 = _dafny.SeqOf((_1820_v29).Dtor_cf3())
			var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1819_v28), 0))
			_ = _index261
			(_1819_v28).ArraySet1((_1821_v30).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1821_v30).Cardinality()))).Uint32()).(_dafny.Int), (_index261).Int())
			var _1822_v31 _dafny.Map
			_ = _1822_v31
			_1822_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _dafny.IntOfInt64(476))
			var _1823_v32 _dafny.MultiSet
			_ = _1823_v32
			_1823_v32 = _dafny.MultiSetOf(_1822_v31, _1822_v31)
			var _1824_v33 _dafny.CodePoint
			_ = _1824_v33
			_1824_v33 = _dafny.CodePoint('v')
			var _1825_v34 _dafny.Sequence
			_ = _1825_v34
			_1825_v34 = _dafny.SeqOf(_1824_v33)
			var _1826_v35 _dafny.MultiSet
			_ = _1826_v35
			_1826_v35 = _dafny.MultiSetOf(p0, (_this).Fm8(_dafny.IntOfInt64(45), (_this).F11(), _1825_v34, _1821_v30, globalState), p0, (_this).F11(), (_this).F4())
			var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1819_v28), 0))
			_ = _index262
			(_1819_v28).ArraySet1(((Companion_Default___.Fm21((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F6())), _1823_v32, _dafny.IntOfInt64(193), (_this).F6(), globalState)).Difference(_1826_v35)).Cardinality(), (_index262).Int())
			var _1827_v36 *C0
			_ = _1827_v36
			var _nw277 *C0 = New_C0_()
			_ = _nw277
			_nw277.Ctor__((_dafny.Zero).Minus(((_1819_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1819_v28), 0))).Int()).(_dafny.Int)).Minus((_1819_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1819_v28), 0))).Int()).(_dafny.Int))))
			_1827_v36 = _nw277
			var _1828_v37 _dafny.Array
			_ = _1828_v37
			var _nw278 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw278
			_1828_v37 = _nw278
			var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1828_v37), 0))
			_ = _index263
			(_1828_v37).ArraySet1(((_this).F4()) || ((_this).F11()), (_index263).Int())
			var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1828_v37), 0))
			_ = _index264
			(_1828_v37).ArraySet1(!((_1780_v0).Select((Companion_Default___.SafeIndex((_dafny.IntOfInt64(175)).Plus((_1819_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1819_v28), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_1780_v0).Cardinality()))).Uint32()).(bool)), (_index264).Int())
			var _1829_v38 D4
			_ = _1829_v38
			_1829_v38 = Companion_D4_.Create_DC13_((_1819_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1819_v28), 0))).Int()).(_dafny.Int), (_this).F11(), _dafny.IntOfUint32((_1825_v34).Cardinality()), _1781_v1)
			var _1830_v39 _dafny.Set
			_ = _1830_v39
			_1830_v39 = _dafny.SetOf((_1829_v38).Dtor_cf15(), p0, (_this).F4())
			var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1828_v37), 0))
			_ = _index265
			var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1819_v28), 0))
			_ = _index266
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1819_v28), 0))
			_ = _index267
			var _rhs249 bool = (_1781_v1).Select((Companion_Default___.SafeIndex(((_1830_v39).Union(_1830_v39)).Cardinality(), _dafny.IntOfUint32((_1781_v1).Cardinality()))).Uint32()).(bool)
			_ = _rhs249
			var _rhs250 _dafny.Int = (func() _dafny.Int {
				if (func() bool {
					if p0 {
						return (_this).F4()
					}
					return p0
				})() {
					return (_this).F6()
				}
				return (_1827_v36).F12()
			})()
			_ = _rhs250
			var _rhs251 _dafny.Int = (_1821_v30).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1821_v30).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs251
			var _lhs151 _dafny.Array = _1828_v37
			_ = _lhs151
			var _lhs152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1828_v37), 0))
			_ = _lhs152
			var _lhs153 _dafny.Array = _1819_v28
			_ = _lhs153
			var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1819_v28), 0))
			_ = _lhs154
			var _lhs155 _dafny.Array = _1819_v28
			_ = _lhs155
			var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1819_v28), 0))
			_ = _lhs156
			(_lhs151).ArraySet1(_rhs249, (_lhs152).Int())
			(_lhs153).ArraySet1(_rhs250, (_lhs154).Int())
			(_lhs155).ArraySet1(_rhs251, (_lhs156).Int())
		} else {
			var _1831___mcc_h5 D3 = _source32.Get_().(D3_DC11).Cf12
			_ = _1831___mcc_h5
			var _1832_cf12 D3 = _1831___mcc_h5
			_ = _1832_cf12
			var _1833_v40 _dafny.Array
			_ = _1833_v40
			var _len0_43 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_43
			var _nw279 _dafny.Array
			_ = _nw279
			if _len0_43.Cmp(_dafny.Zero) == 0 {
				_nw279 = _dafny.NewArray(_len0_43)
			} else {
				var _init43 func(_dafny.Int) _dafny.Int = func(_1834_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_1834_i1, (_this).F6())
				}
				_ = _init43
				var _element0_43 = _init43(_dafny.Zero)
				_ = _element0_43
				_nw279 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
				(_nw279).ArraySet1(_element0_43, 0)
				var _nativeLen0_43 = (_len0_43).Int()
				_ = _nativeLen0_43
				for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
					(_nw279).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
				}
			}
			_1833_v40 = _nw279
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))
			_ = _index268
			(_1833_v40).ArraySet1(_dafny.IntOfInt64(257), (_index268).Int())
			var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))
			_ = _index269
			(_1833_v40).ArraySet1((_this).F6(), (_index269).Int())
			var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))
			_ = _index270
			(_1833_v40).ArraySet1((_this).F6(), (_index270).Int())
			var _1835_v41 bool
			_ = _1835_v41
			_1835_v41 = false
			var _1836_v42 _dafny.Array
			_ = _1836_v42
			var _nw280 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw280
			_1836_v42 = _nw280
			var _1837_v43 _dafny.MultiSet
			_ = _1837_v43
			_1837_v43 = _dafny.MultiSetOf(p0, !(_1835_v41), (_this).F4())
			var _1838_v44 _dafny.Map
			_ = _1838_v44
			_1838_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1837_v43)
			var _1839_v45 _dafny.MultiSet
			_ = _1839_v45
			_1839_v45 = _dafny.MultiSetOf(_1837_v43, _dafny.MultiSetFromSeq(_1781_v1), (func() _dafny.MultiSet {
				if (_1838_v44).Contains(_1835_v41) {
					return (_1838_v44).Get(_1835_v41).(_dafny.MultiSet)
				}
				return _1837_v43
			})(), _1837_v43)
			var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1836_v42), 0))
			_ = _index271
			(_1836_v42).ArraySet1((_1839_v45).IsSubsetOf(_1839_v45), (_index271).Int())
			var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))
			_ = _index272
			var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1836_v42), 0))
			_ = _index273
			var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))
			_ = _index274
			var _rhs252 bool = true
			_ = _rhs252
			var _rhs253 _dafny.Int = (_1833_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))).Int()).(_dafny.Int)
			_ = _rhs253
			var _rhs254 bool = (_this).F11()
			_ = _rhs254
			var _rhs255 _dafny.Int = Companion_Default___.SafeModuloInt((_1833_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
				if (_this).F11() {
					return (_this).F6()
				}
				return (_1833_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))).Int()).(_dafny.Int)
			})())
			_ = _rhs255
			var _lhs157 _dafny.Array = _1833_v40
			_ = _lhs157
			var _lhs158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))
			_ = _lhs158
			var _lhs159 _dafny.Array = _1836_v42
			_ = _lhs159
			var _lhs160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1836_v42), 0))
			_ = _lhs160
			var _lhs161 _dafny.Array = _1833_v40
			_ = _lhs161
			var _lhs162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))
			_ = _lhs162
			_1835_v41 = _rhs252
			(_lhs157).ArraySet1(_rhs253, (_lhs158).Int())
			(_lhs159).ArraySet1(_rhs254, (_lhs160).Int())
			(_lhs161).ArraySet1(_rhs255, (_lhs162).Int())
			if true {
				var _1840_v46 D0
				_ = _1840_v46
				_1840_v46 = Companion_D0_.Create_DC2_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SetOf((_this).F6(), (_this).F6()))).Cardinality()))
				var _1841_v47 _dafny.MultiSet
				_ = _1841_v47
				_1841_v47 = _dafny.MultiSetOf((_1833_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))).Int()).(_dafny.Int), (_1840_v46).Dtor_cf1())
				var _1842_v48 D3
				_ = _1842_v48
				_1842_v48 = Companion_D3_.Create_DC9_((_this).F11(), (_1841_v47).Cardinality(), (_this).F11())
				var _1843_v49 _dafny.Sequence
				_ = _1843_v49
				_1843_v49 = _dafny.SeqOf((_1842_v48).Dtor_cf9(), (_1833_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))).Int()).(_dafny.Int))
				var _1844_v50 bool
				_ = _1844_v50
				var _1845_v51 _dafny.Int
				_ = _1845_v51
				var _out51 bool
				_ = _out51
				var _out52 _dafny.Int
				_ = _out52
				_out51, _out52 = Companion_Default___.M0(Companion_Default___.Fm22(globalState), (p0) == ((_this).F4()), (_dafny.MultiSetFromSeq(_1843_v49)).Intersection(_1841_v47), _1835_v41, globalState)
				_1844_v50 = _out51
				_1845_v51 = _out52
				var _1846_v52 _dafny.Map
				_ = _1846_v52
				_1846_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1833_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))).Int()).(_dafny.Int)).Plus(_1845_v51), _dafny.IntOfInt64(200))
				_1846_v52 = (_1846_v52).Update(Companion_Default___.Fm3(_1845_v51, _1835_v41, globalState), (_1843_v49).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-562), _dafny.IntOfUint32((_1843_v49).Cardinality()))).Uint32()).(_dafny.Int))
				var _1847_v53 _dafny.Sequence
				_ = _1847_v53
				_1847_v53 = _dafny.UnicodeSeqOfUtf8Bytes("owcvyfr")
				var _1848_v54 _dafny.Map
				_ = _1848_v54
				_1848_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_1847_v53).Cardinality())).Times(_1845_v51), _1780_v0)
				_1848_v54 = (_1848_v54).Update(_1845_v51, _dafny.SeqOf((_1836_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1836_v42), 0))).Int()).(bool), (_this).F11(), (_this).F11(), p0))
				var _1849_v55 _dafny.CodePoint
				_ = _1849_v55
				_1849_v55 = _dafny.CodePoint('d')
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1836_v42), 0))
				_ = _index275
				(_1836_v42).ArraySet1(Companion_Default___.Fm0(_1849_v55, globalState), (_index275).Int())
				var _1850_v56 _dafny.Set
				_ = _1850_v56
				_1850_v56 = _dafny.SetOf((_1836_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1836_v42), 0))).Int()).(bool))
				var _1851_v57 _dafny.Sequence
				_ = _1851_v57
				_1851_v57 = _dafny.SeqOf(_1833_v40, _1833_v40)
				var _rhs256 _dafny.Set = Companion_Default___.Fm23(_1845_v51, (_this).F6(), globalState)
				_ = _rhs256
				var _rhs257 _dafny.Sequence = _1851_v57
				_ = _rhs257
				var _rhs258 bool = _1835_v41
				_ = _rhs258
				_1850_v56 = _rhs256
				_1851_v57 = _rhs257
				_1844_v50 = _rhs258
			} else {
				_1835_v41 = !((_this).F11())
				var _1852_v58 _dafny.Map
				_ = _1852_v58
				_1852_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1837_v43).Cardinality(), true)
				var _1853_v59 _dafny.MultiSet
				_ = _1853_v59
				_1853_v59 = _dafny.MultiSetOf(_1852_v58)
				var _1854_v60 _dafny.Sequence
				_ = _1854_v60
				_1854_v60 = _dafny.SeqOf(_1852_v58, _1852_v58)
				var _1855_v61 _dafny.Sequence
				_ = _1855_v61
				_1855_v61 = _dafny.UnicodeSeqOfUtf8Bytes("wgpjpb")
				var _1856_v62 _dafny.Map
				_ = _1856_v62
				_1856_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_1853_v59).Union(_dafny.MultiSetOf(_1852_v58, _1852_v58, _1852_v58, (_1854_v60).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1855_v61).Cardinality()), _dafny.IntOfUint32((_1854_v60).Cardinality()))).Uint32()).(_dafny.Map), _1852_v58)))
				_1856_v62 = (_1856_v62).Update((_1833_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))).Int()).(_dafny.Int), _1853_v59)
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_1833_v40), 0))
				_ = _index276
				(_1833_v40).ArraySet1(((_this).F6()).Minus((_this).F6()), (_index276).Int())
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1836_v42), 0))
				_ = _index277
				(_1836_v42).ArraySet1((_dafny.IntOfInt64(803)).Cmp(((_this).Fm9(p0, p0, (_this).Fm15(globalState), globalState)).Times((func() _dafny.Int {
					if (_1837_v43).Contains(p0) {
						return (_1837_v43).Multiplicity(p0)
					}
					return (_this).F6()
				})())) >= 0, (_index277).Int())
				var _1857_v63 _dafny.CodePoint
				_ = _1857_v63
				_1857_v63 = _dafny.CodePoint('e')
				r0 = _1857_v63
			}
		}
		var _1858_i2 _dafny.Int
		_ = _1858_i2
		_1858_i2 = _dafny.Zero
		{
			for (_this).F4() {
				{
					if (_1858_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1858_i2 = (_1858_i2).Plus(_dafny.One)
					var _1859_v64 _dafny.Array
					_ = _1859_v64
					var _len0_44 _dafny.Int = _dafny.IntOfInt64(21)
					_ = _len0_44
					var _nw281 _dafny.Array
					_ = _nw281
					if _len0_44.Cmp(_dafny.Zero) == 0 {
						_nw281 = _dafny.NewArray(_len0_44)
					} else {
						var _init44 func(_dafny.Int) bool = (func(_1860_p0 bool) func(_dafny.Int) bool {
							return func(_1861_i3 _dafny.Int) bool {
								return _1860_p0
							}
						})(p0)
						_ = _init44
						var _element0_44 = _init44(_dafny.Zero)
						_ = _element0_44
						_nw281 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
						(_nw281).ArraySet1(_element0_44, 0)
						var _nativeLen0_44 = (_len0_44).Int()
						_ = _nativeLen0_44
						for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
							(_nw281).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
						}
					}
					_1859_v64 = _nw281
					_1859_v64 = _1859_v64
					var _1862_v65 bool
					_ = _1862_v65
					_1862_v65 = true
					_1862_v65 = (_this).F4()
					var _1863_v66 _dafny.MultiSet
					_ = _1863_v66
					_1863_v66 = _dafny.MultiSetOf(_dafny.IntOfInt64(126), (_this).F6())
					var _1864_v67 D0
					_ = _1864_v67
					_1864_v67 = Companion_D0_.Create_DC2_((_1863_v66).Cardinality())
					var _1865_v68 _dafny.Map
					_ = _1865_v68
					_1865_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_1864_v67).Dtor_cf1())
					_1865_v68 = (_1865_v68).Update((_this).F4(), ((_this).F6()).Plus((_this).F6()))
					var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1859_v64), 0))
					_ = _index278
					(_1859_v64).ArraySet1(!(p0), (_index278).Int())
					var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1859_v64), 0))
					_ = _index279
					(_1859_v64).ArraySet1(p0, (_index279).Int())
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _1866_v69 _dafny.Map
		_ = _1866_v69
		_1866_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
		_1866_v69 = (_1866_v69).Merge((_1866_v69).Merge(_1866_v69))
		_1780_v0 = _dafny.Companion_Sequence_.Concatenate(_1780_v0, _1781_v1)
		var _hi12 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F6(), (_this).F6())
		_ = _hi12
		for _1867_i4 := (_this).F6(); _1867_i4.Cmp(_hi12) < 0; _1867_i4 = _1867_i4.Plus(_dafny.One) {
			var _1868_v70 _dafny.Int
			_ = _1868_v70
			_1868_v70 = _dafny.IntOfInt64(577)
			_1868_v70 = (_this).F6()
			var _1869_v71 D2
			_ = _1869_v71
			_1869_v71 = Companion_D2_.Create_DC7_((_this).F6())
			var _source33 D2 = _1869_v71
			_ = _source33
			if _source33.Is_DC6() {
				var _1870_v72 bool
				_ = _1870_v72
				var _1871_v73 bool
				_ = _1871_v73
				var _1872_v74 D1
				_ = _1872_v74
				var _out53 bool
				_ = _out53
				var _out54 bool
				_ = _out54
				var _out55 D1
				_ = _out55
				_out53, _out54, _out55 = (_this).M8(Companion_Default___.SafeModuloInt((_this).F6(), _1868_v70), _1868_v70, globalState)
				_1870_v72 = _out53
				_1871_v73 = _out54
				_1872_v74 = _out55
				var _1873_v75 _dafny.CodePoint
				_ = _1873_v75
				_1873_v75 = _dafny.CodePoint('j')
				var _1874_v76 _dafny.Sequence
				_ = _1874_v76
				_1874_v76 = _dafny.UnicodeSeqOfUtf8Bytes("ftxvovkv")
				var _1875_v77 D0
				_ = _1875_v77
				_1875_v77 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_1874_v76).Cardinality()))
				var _1876_v78 _dafny.Map
				_ = _1876_v78
				_1876_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1875_v77, _1870_v72)
				var _1877_v79 _dafny.Map
				_ = _1877_v79
				_1877_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_1873_v75, globalState), _1876_v78)
				var _1878_v80 _dafny.Array
				_ = _1878_v80
				var _nwElement0_64 _dafny.Int = _1868_v70
				_ = _nwElement0_64
				var _nw282 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(5))
				_ = _nw282
				(_nw282).ArraySet1(_nwElement0_64, 0)
				(_nw282).ArraySet1((_dafny.Zero).Minus(((func() _dafny.Map {
					if (_1877_v79).Contains(!(p0)) {
						return (_1877_v79).Get(!(p0)).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_1868_v70), (_this).F4())
				})()).Cardinality()), 1)
				(_nw282).ArraySet1((_this).F6(), 2)
				(_nw282).ArraySet1(_dafny.IntOfInt64(-352), 3)
				(_nw282).ArraySet1((_this).F6(), 4)
				_1878_v80 = _nw282
				var _1879_v81 _dafny.Set
				_ = _1879_v81
				_1879_v81 = _dafny.SetOf(_1878_v80)
				_1870_v72 = (_1879_v81).IsDisjointFrom(_1879_v81)
				_1866_v69 = (_1866_v69).Update((_this).F6(), (_this).F6())
				var _1880_v82 D0
				_ = _1880_v82
				_1880_v82 = Companion_D0_.Create_DC2_(_1868_v70)
				_1880_v82 = _1880_v82
			} else if _source33.Is_DC7() {
				var _1881___mcc_h6 _dafny.Int = _source33.Get_().(D2_DC7).Cf6
				_ = _1881___mcc_h6
				var _1882_cf6 _dafny.Int = _1881___mcc_h6
				_ = _1882_cf6
				var _1883_v83 _dafny.Array
				_ = _1883_v83
				var _len0_45 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_45
				var _nw283 _dafny.Array
				_ = _nw283
				if _len0_45.Cmp(_dafny.Zero) == 0 {
					_nw283 = _dafny.NewArray(_len0_45)
				} else {
					var _init45 func(_dafny.Int) bool = (func(_1884_p0 bool) func(_dafny.Int) bool {
						return func(_1885_i5 _dafny.Int) bool {
							return _1884_p0
						}
					})(p0)
					_ = _init45
					var _element0_45 = _init45(_dafny.Zero)
					_ = _element0_45
					_nw283 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
					(_nw283).ArraySet1(_element0_45, 0)
					var _nativeLen0_45 = (_len0_45).Int()
					_ = _nativeLen0_45
					for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
						(_nw283).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
					}
				}
				_1883_v83 = _nw283
				var _1886_v84 _dafny.CodePoint
				_ = _1886_v84
				_1886_v84 = _dafny.CodePoint('w')
				var _1887_v85 _dafny.Map
				_ = _1887_v85
				_1887_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1886_v84, _1867_i4)
				var _1888_v87 _dafny.Sequence
				_ = _1888_v87
				_1888_v87 = _dafny.SeqOf(_1886_v84, _1886_v84)
				var _1889_v88 _dafny.Sequence
				_ = _1889_v88
				_1889_v88 = _dafny.SeqOf(func() _dafny.Map {
					var _coll55 = _dafny.NewMapBuilder()
					_ = _coll55
					for _iter62 := _dafny.Iterate((_1888_v87).Elements()); ; {
						_compr_55, _ok62 := _iter62()
						if !_ok62 {
							break
						}
						var _1890_v86 _dafny.CodePoint
						_1890_v86 = interface{}(_compr_55).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_1888_v87, _1890_v86) {
							_coll55.Add(_1890_v86, _1868_v70)
						}
					}
					return _coll55.ToMap()
				}())
				var _1891_v89 _dafny.Map
				_ = _1891_v89
				_1891_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1883_v83), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1887_v85, _1887_v85), _dafny.Companion_Sequence_.Update(_1889_v88, (Companion_Default___.SafeIndex(_1882_cf6, _dafny.IntOfUint32((_1889_v88).Cardinality()))).Uint32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(Companion_Default___.Fm3(_1882_cf6, (_this).F11(), globalState), (_this).F11(), globalState), (_dafny.Zero).Minus((_this).F6())))))
				var _1892_v90 _dafny.Sequence
				_ = _1892_v90
				_1892_v90 = _dafny.SeqOf(_1883_v83, _1883_v83, _1883_v83)
				var _1893_v91 _dafny.Sequence
				_ = _1893_v91
				_1893_v91 = _dafny.SeqOf(_1883_v83, _1883_v83, (_1892_v90).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-553), _dafny.IntOfUint32((_1892_v90).Cardinality()))).Uint32()).(_dafny.Array), (Companion_D2_.Create_DC5_(_1883_v83)).Dtor_cf5(), _1883_v83)
				_1891_v89 = (_1891_v89).Update(_1892_v90, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(707))).Uint32(), func(coer101 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg101 _dafny.Int) interface{} {
						return coer101(arg101)
					}
				}((func(_1894_v85 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_1895_i6 _dafny.Int) _dafny.Map {
						return _1894_v85
					}
				})(_1887_v85))))
				var _1896_v92 _dafny.Map
				_ = _1896_v92
				_1896_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1888_v87, (_this).F6())
				_1896_v92 = (_1896_v92).Update(_dafny.UnicodeSeqOfUtf8Bytes("qpjli"), _1882_cf6)
				var _1897_v93 _dafny.Map
				_ = _1897_v93
				_1897_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				r0 = Companion_Default___.Fm1(((_1897_v93).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))).Cardinality(), (_this).F4(), globalState)
				var _1898_v94 _dafny.Array
				_ = _1898_v94
				var _nw284 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
				_ = _nw284
				_1898_v94 = _nw284
				var _1899_v95 _dafny.Array
				_ = _1899_v95
				var _len0_46 _dafny.Int = _dafny.One
				_ = _len0_46
				var _nw285 _dafny.Array
				_ = _nw285
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw285 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) _dafny.Int = func(_1900_i7 _dafny.Int) _dafny.Int {
						return (_1900_i7).Minus(_dafny.IntOfInt64(964))
					}
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw285 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw285).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw285).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1899_v95 = _nw285
				var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1898_v94), 0))
				_ = _index280
				(_1898_v94).ArraySet1(_1899_v95, (_index280).Int())
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1898_v94), 0))
				_ = _index281
				(_1898_v94).ArraySet1(_1899_v95, (_index281).Int())
			} else {
				var _1901___mcc_h7 _dafny.Array = _source33.Get_().(D2_DC5).Cf5
				_ = _1901___mcc_h7
				var _1902_cf5 _dafny.Array = _1901___mcc_h7
				_ = _1902_cf5
				var _1903_v96 *C0
				_ = _1903_v96
				var _nw286 *C0 = New_C0_()
				_ = _nw286
				_nw286.Ctor__((_this).F6())
				_1903_v96 = _nw286
				var _1904_v97 _dafny.MultiSet
				_ = _1904_v97
				_1904_v97 = _dafny.MultiSetOf((_1903_v96).F12(), (_1903_v96).F12(), (_this).F6(), _1868_v70)
				var _1905_v98 _dafny.Map
				_ = _1905_v98
				_1905_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1904_v97, _dafny.UnicodeSeqOfUtf8Bytes("synxuae"))
				var _1906_v99 bool
				_ = _1906_v99
				var _1907_v100 _dafny.Int
				_ = _1907_v100
				var _out56 bool
				_ = _out56
				var _out57 _dafny.Int
				_ = _out57
				_out56, _out57 = Companion_Default___.M0(_1905_v98, (_this).F11(), _dafny.MultiSetOf(_1868_v70), (_this).F11(), globalState)
				_1906_v99 = _out56
				_1907_v100 = _out57
				var _1908_v101 _dafny.Sequence
				_ = _1908_v101
				_1908_v101 = _dafny.UnicodeSeqOfUtf8Bytes("adljl")
				var _1909_v102 _dafny.Set
				_ = _1909_v102
				_1909_v102 = _dafny.SetOf((func() _dafny.Sequence {
					if (_1780_v0).Select((Companion_Default___.SafeIndex(_1868_v70, _dafny.IntOfUint32((_1780_v0).Cardinality()))).Uint32()).(bool) {
						return _1908_v101
					}
					return _1908_v101
				})())
				_1909_v102 = _1909_v102
				var _1910_v103 _dafny.Map
				_ = _1910_v103
				_1910_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(p0)), (_dafny.Zero).Minus(_1868_v70))
				var _1911_v104 _dafny.Sequence
				_ = _1911_v104
				_1911_v104 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1906_v99, (_1903_v96).F12()), _1910_v103, _1910_v103, _1910_v103)
				var _1912_v105 *C0
				_ = _1912_v105
				var _nw287 *C0 = New_C0_()
				_ = _nw287
				_nw287.Ctor__(_dafny.IntOfUint32((_1911_v104).Cardinality()))
				_1912_v105 = _nw287
			}
			var _1913_v106 _dafny.Sequence
			_ = _1913_v106
			_1913_v106 = _dafny.UnicodeSeqOfUtf8Bytes("tasx")
			var _1914_v107 _dafny.Array
			_ = _1914_v107
			var _nw288 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw288
			_1914_v107 = _nw288
			var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_1914_v107), 0))
			_ = _index282
			(_1914_v107).ArraySet1(_dafny.IntOfInt64(415), (_index282).Int())
			var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_1914_v107), 0))
			_ = _index283
			var _rhs259 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("xohxe")
			_ = _rhs259
			var _rhs260 _dafny.Int = _dafny.IntOfUint32((_1913_v106).Cardinality())
			_ = _rhs260
			var _lhs163 _dafny.Array = _1914_v107
			_ = _lhs163
			var _lhs164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_1914_v107), 0))
			_ = _lhs164
			_1913_v106 = _rhs259
			(_lhs163).ArraySet1(_rhs260, (_lhs164).Int())
			var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_1914_v107), 0))
			_ = _index284
			(_1914_v107).ArraySet1(_1868_v70, (_index284).Int())
		}
		var _1915_v108 _dafny.CodePoint
		_ = _1915_v108
		_1915_v108 = _dafny.CodePoint('h')
		if Companion_Default___.Fm0(_1915_v108, globalState) {
			var _1916_v109 _dafny.Int
			_ = _1916_v109
			_1916_v109 = _dafny.IntOfInt64(639)
			var _1917_v110 _dafny.MultiSet
			_ = _1917_v110
			_1917_v110 = _dafny.MultiSetOf((_this).F4())
			_1916_v109 = ((_this).F6()).Times(_dafny.IntOfUint32((_dafny.SeqOf(_1916_v109, (_1917_v110).Cardinality(), (_this).F6())).Cardinality()))
			var _1918_v111 bool
			_ = _1918_v111
			_1918_v111 = false
			var _1919_v112 _dafny.Sequence
			_ = _1919_v112
			_1919_v112 = _dafny.UnicodeSeqOfUtf8Bytes("s")
			var _1920_v113 _dafny.Map
			_ = _1920_v113
			_1920_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1919_v112).Cardinality()), _1918_v111)
			_1918_v111 = (func() bool {
				if (_1920_v113).Contains(_1916_v109) {
					return (_1920_v113).Get(_1916_v109).(bool)
				}
				return (_this).F11()
			})()
			var _1921_v114 _dafny.Map
			_ = _1921_v114
			_1921_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1780_v0, _1918_v111)
			_1921_v114 = (_1921_v114).Update(_dafny.SeqOf(p0), p0)
			var _1922_v115 _dafny.MultiSet
			_ = _1922_v115
			_1922_v115 = _dafny.MultiSetOf(Companion_Default___.Fm19(p0, _1916_v109, globalState))
			var _1923_v116 D1
			_ = _1923_v116
			_1923_v116 = Companion_D1_.Create_DC3_(_1918_v111)
			var _1924_v117 _dafny.Sequence
			_ = _1924_v117
			_1924_v117 = _dafny.SeqOf(_1780_v0, _1780_v0)
			var _1925_v118 _dafny.Array
			_ = _1925_v118
			var _nwElement0_65 _dafny.MultiSet = _1922_v115
			_ = _nwElement0_65
			var _nw289 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(22))
			_ = _nw289
			(_nw289).ArraySet1(_nwElement0_65, 0)
			(_nw289).ArraySet1(_dafny.MultiSetOf(_1781_v1, _1780_v0, _dafny.SeqOf((_this).F11()), _1780_v0), 1)
			(_nw289).ArraySet1(_1922_v115, 2)
			(_nw289).ArraySet1(_1922_v115, 3)
			(_nw289).ArraySet1(_dafny.MultiSetOf(_1781_v1), 4)
			(_nw289).ArraySet1(_1922_v115, 5)
			(_nw289).ArraySet1(_1922_v115, 6)
			(_nw289).ArraySet1(_dafny.MultiSetOf(_1781_v1, _dafny.SeqOf((_this).F11(), true), _1780_v0, _1780_v0), 7)
			(_nw289).ArraySet1((func() _dafny.MultiSet {
				if (_1923_v116).Dtor_cf2() {
					return _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(146))).Uint32(), func(coer102 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg102 _dafny.Int) interface{} {
							return coer102(arg102)
						}
					}((func(_1926_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1927_i8 _dafny.Int) _dafny.Sequence {
							return _1926_v0
						}
					})(_1780_v0))))
				}
				return Companion_Default___.Fm24(_1916_v109, _1915_v108, _1916_v109, globalState)
			})(), 8)
			(_nw289).ArraySet1(_1922_v115, 9)
			(_nw289).ArraySet1((_1922_v115).Intersection(_dafny.MultiSetFromSeq(_1924_v117)), 10)
			(_nw289).ArraySet1(_1922_v115, 11)
			(_nw289).ArraySet1((_1922_v115).Intersection(_1922_v115), 12)
			(_nw289).ArraySet1(_1922_v115, 13)
			(_nw289).ArraySet1((_1922_v115).Intersection(_1922_v115), 14)
			(_nw289).ArraySet1((_dafny.MultiSetOf(_dafny.SeqOf((_this).F4()))).Intersection(_1922_v115), 15)
			(_nw289).ArraySet1(_1922_v115, 16)
			(_nw289).ArraySet1(_1922_v115, 17)
			(_nw289).ArraySet1(_1922_v115, 18)
			(_nw289).ArraySet1(_1922_v115, 19)
			(_nw289).ArraySet1(_1922_v115, 20)
			(_nw289).ArraySet1(_1922_v115, 21)
			_1925_v118 = _nw289
			var _1928_v119 _dafny.Map
			_ = _1928_v119
			_1928_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1918_v111, _dafny.SeqOf(_1781_v1))
			var _1929_v120 D6
			_ = _1929_v120
			_1929_v120 = Companion_D6_.Create_DC15_(_1922_v115)
			var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen((_1925_v118), 0))
			_ = _index285
			(_1925_v118).ArraySet1((_dafny.MultiSetFromSeq((func() _dafny.Sequence {
				if (_1928_v119).Contains(p0) {
					return (_1928_v119).Get(p0).(_dafny.Sequence)
				}
				return _1924_v117
			})())).Union((_1929_v120).Dtor_cf19()), (_index285).Int())
			var _1930_v121 _dafny.Sequence
			_ = _1930_v121
			_1930_v121 = _dafny.SeqOf(_1916_v109, (_dafny.Zero).Minus(_1916_v109))
			var _1931_v122 _dafny.MultiSet
			_ = _1931_v122
			_1931_v122 = _dafny.MultiSetOf(_1930_v121)
			var _1932_v123 _dafny.Array
			_ = _1932_v123
			var _nw290 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw290
			_1932_v123 = _nw290
			var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen((_1925_v118), 0))
			_ = _index286
			var _rhs261 _dafny.MultiSet = (_1922_v115).Update(_dafny.SeqOf(false), Companion_Default___.Abs((_this).F6()))
			_ = _rhs261
			var _rhs262 _dafny.MultiSet = _1931_v122
			_ = _rhs262
			var _rhs263 _dafny.Array = _1932_v123
			_ = _rhs263
			var _rhs264 bool = (_this).F11()
			_ = _rhs264
			var _lhs165 _dafny.Array = _1925_v118
			_ = _lhs165
			var _lhs166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen((_1925_v118), 0))
			_ = _lhs166
			(_lhs165).ArraySet1(_rhs261, (_lhs166).Int())
			_1931_v122 = _rhs262
			_1932_v123 = _rhs263
			_1918_v111 = _rhs264
			_1916_v109 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(313), _1916_v109)
		} else {
			var _1933_v124 D3
			_ = _1933_v124
			_1933_v124 = Companion_D3_.Create_DC10_((_this).F6())
			_1933_v124 = Companion_D3_.Create_DC10_((_this).F6())
			var _1934_v125 _dafny.Int
			_ = _1934_v125
			_1934_v125 = _dafny.IntOfInt64(216)
			_1934_v125 = (func() _dafny.Int {
				if p0 {
					return (func() _dafny.Int {
						if (_1866_v69).Contains(_dafny.IntOfInt64(-103)) {
							return (_1866_v69).Get(_dafny.IntOfInt64(-103)).(_dafny.Int)
						}
						return _1934_v125
					})()
				}
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1780_v0, (Companion_Default___.SafeIndex(_1934_v125, _dafny.IntOfUint32((_1780_v0).Cardinality()))).Uint32(), Companion_Default___.Fm0(_1915_v108, globalState)), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1780_v0, (Companion_Default___.SafeIndex(_1934_v125, _dafny.IntOfUint32((_1780_v0).Cardinality()))).Uint32(), Companion_Default___.Fm0(_1915_v108, globalState))).Cardinality()))).Uint32(), true)).Cardinality())
			})()
			var _source34 D3 = Companion_D3_.Create_DC8_(_1780_v0)
			_ = _source34
			if _source34.Is_DC9() {
				var _1935___mcc_h8 bool = _source34.Get_().(D3_DC9).Cf8
				_ = _1935___mcc_h8
				var _1936___mcc_h9 _dafny.Int = _source34.Get_().(D3_DC9).Cf9
				_ = _1936___mcc_h9
				var _1937___mcc_h10 bool = _source34.Get_().(D3_DC9).Cf10
				_ = _1937___mcc_h10
				var _1938_cf10 bool = _1937___mcc_h10
				_ = _1938_cf10
				var _1939_cf9 _dafny.Int = _1936___mcc_h9
				_ = _1939_cf9
				var _1940_cf8 bool = _1935___mcc_h8
				_ = _1940_cf8
				var _1941_v126 _dafny.MultiSet
				_ = _1941_v126
				_1941_v126 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(676))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg103 _dafny.Int) interface{} {
						return coer103(arg103)
					}
				}((func(_1942_v108 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1943_i9 _dafny.Int) _dafny.CodePoint {
						return _1942_v108
					}
				})(_1915_v108)))).Cardinality()), (_dafny.SetOf((_this).F11(), !((_this).F4()), _1940_cf8)).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm25(_1939_cf9, (_this).F6(), (_this).F4(), (_this).F11(), globalState)).Cardinality()), _1934_v125)
				var _1944_v127 _dafny.Sequence
				_ = _1944_v127
				_1944_v127 = _dafny.SeqOf(_1941_v126)
				var _1945_v128 _dafny.Sequence
				_ = _1945_v128
				_1945_v128 = _dafny.SeqOf(_1939_cf9, (_this).F6(), _1934_v125)
				_1941_v126 = (_1944_v127).Select((Companion_Default___.SafeIndex(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1934_v125, _1945_v128)).Update((_this).F6(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(807))).Uint32(), func(coer104 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg104 _dafny.Int) interface{} {
						return coer104(arg104)
					}
				}((func(_1946_cf9 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1947_i10 _dafny.Int) _dafny.Int {
						return _1946_cf9
					}
				})(_1939_cf9))))).Cardinality(), _dafny.IntOfUint32((_1944_v127).Cardinality()))).Uint32()).(_dafny.MultiSet)
				var _1948_v129 _dafny.Set
				_ = _1948_v129
				_1948_v129 = _dafny.SetOf(_1934_v125)
				_1939_cf9 = (_this).Fm9((_this).F11(), _1940_cf8, (_1948_v129).Cardinality(), globalState)
				var _1949_v130 _dafny.Array
				_ = _1949_v130
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_47
				var _nw291 _dafny.Array
				_ = _nw291
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw291 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) _dafny.Int = func(_1950_i11 _dafny.Int) _dafny.Int {
						return (_1950_i11).Times((_this).F6())
					}
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw291 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw291).ArraySet1(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw291).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_1949_v130 = _nw291
				var _1951_v131 _dafny.Array
				_ = _1951_v131
				var _len0_48 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_48
				var _nw292 _dafny.Array
				_ = _nw292
				if _len0_48.Cmp(_dafny.Zero) == 0 {
					_nw292 = _dafny.NewArray(_len0_48)
				} else {
					var _init48 func(_dafny.Int) D0 = func(_1952_i12 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_((_this).F6())
					}
					_ = _init48
					var _element0_48 = _init48(_dafny.Zero)
					_ = _element0_48
					_nw292 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
					(_nw292).ArraySet1(_element0_48, 0)
					var _nativeLen0_48 = (_len0_48).Int()
					_ = _nativeLen0_48
					for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
						(_nw292).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
					}
				}
				_1951_v131 = _nw292
				var _1953_v132 D0
				_ = _1953_v132
				_1953_v132 = Companion_D0_.Create_DC0_(_1939_cf9)
				var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_1951_v131), 0))
				_ = _index287
				(_1951_v131).ArraySet1(_1953_v132, (_index287).Int())
				var _1954_v133 _dafny.Array
				_ = _1954_v133
				var _nw293 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw293
				_1954_v133 = _nw293
				var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_1951_v131), 0))
				_ = _index288
				var _rhs265 D3 = _1933_v124
				_ = _rhs265
				var _rhs266 _dafny.Array = _1949_v130
				_ = _rhs266
				var _rhs267 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(533), _1939_cf9))).Plus((func() _dafny.Int {
					if _1938_cf10 {
						return _1939_cf9
					}
					return _dafny.IntOfInt64(-614)
				})())))
				_ = _rhs267
				var _rhs268 D0 = _1953_v132
				_ = _rhs268
				var _rhs269 _dafny.Array = _1954_v133
				_ = _rhs269
				var _lhs167 _dafny.Array = _1951_v131
				_ = _lhs167
				var _lhs168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_1951_v131), 0))
				_ = _lhs168
				_1933_v124 = _rhs265
				_1949_v130 = _rhs266
				_1939_cf9 = _rhs267
				(_lhs167).ArraySet1(_rhs268, (_lhs168).Int())
				_1954_v133 = _rhs269
				_1934_v125 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_1939_cf9))
			} else if _source34.Is_DC10() {
				var _1955___mcc_h11 _dafny.Int = _source34.Get_().(D3_DC10).Cf11
				_ = _1955___mcc_h11
				var _1956_cf11 _dafny.Int = _1955___mcc_h11
				_ = _1956_cf11
				var _1957_v134 _dafny.Array
				_ = _1957_v134
				var _len0_49 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_49
				var _nw294 _dafny.Array
				_ = _nw294
				if _len0_49.Cmp(_dafny.Zero) == 0 {
					_nw294 = _dafny.NewArray(_len0_49)
				} else {
					var _init49 func(_dafny.Int) _dafny.Int = func(_1958_i13 _dafny.Int) _dafny.Int {
						return (_1958_i13).Plus((_this).F6())
					}
					_ = _init49
					var _element0_49 = _init49(_dafny.Zero)
					_ = _element0_49
					_nw294 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
					(_nw294).ArraySet1(_element0_49, 0)
					var _nativeLen0_49 = (_len0_49).Int()
					_ = _nativeLen0_49
					for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
						(_nw294).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
					}
				}
				_1957_v134 = _nw294
				var _1959_v135 _dafny.Map
				_ = _1959_v135
				_1959_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1957_v134, p0)
				_1959_v135 = (((_1959_v135).Update(_1957_v134, (_this).F4())).Update(_1957_v134, (_this).F4()))
				var _1960_v136 bool
				_ = _1960_v136
				_1960_v136 = true
				var _rhs270 _dafny.Int = (_this).F6()
				_ = _rhs270
				var _rhs271 _dafny.Int = ((_this).F6()).Plus(_1956_cf11)
				_ = _rhs271
				var _rhs272 bool = !_dafny.Companion_Sequence_.Equal((Companion_D3_.Create_DC8_(_1780_v0)).Dtor_cf7(), _dafny.SeqOf(p0, p0, (_this).F11(), _1960_v136))
				_ = _rhs272
				_1956_cf11 = _rhs270
				_1956_cf11 = _rhs271
				_1960_v136 = _rhs272
				var _1961_v137 _dafny.Sequence
				_ = _1961_v137
				_1961_v137 = _dafny.UnicodeSeqOfUtf8Bytes("fjydewlp")
				var _1962_v138 T1
				_ = _1962_v138
				var _nw295 *C3 = New_C3_()
				_ = _nw295
				_nw295.Ctor__(_this.F5(), _dafny.IntOfUint32((_1961_v137).Cardinality()), false)
				_1962_v138 = _nw295
				var _1963_v139 _dafny.Map
				_ = _1963_v139
				_1963_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1960_v136, _1962_v138)
				var _1964_v140 _dafny.Map
				_ = _1964_v140
				_1964_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1962_v138)
				_1963_v139 = (_1963_v139).Update((func() bool {
					if p0 {
						return (_this).F11()
					}
					return (_1962_v138).F4()
				})(), (func() T1 {
					if (_1964_v140).Contains((_1962_v138).F6()) {
						return (_1964_v140).Get((_1962_v138).F6()).(T1)
					}
					return _1962_v138
				})())
				_1934_v125 = _1934_v125
			} else if _source34.Is_DC8() {
				var _1965___mcc_h12 _dafny.Sequence = _source34.Get_().(D3_DC8).Cf7
				_ = _1965___mcc_h12
				var _1966_cf7 _dafny.Sequence = _1965___mcc_h12
				_ = _1966_cf7
				var _1967_v141 _dafny.Array
				_ = _1967_v141
				var _nw296 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw296
				_1967_v141 = _nw296
				var _1968_v142 _dafny.Sequence
				_ = _1968_v142
				_1968_v142 = _dafny.SeqOf(_1967_v141)
				var _1969_v143 _dafny.Sequence
				_ = _1969_v143
				_1969_v143 = _dafny.UnicodeSeqOfUtf8Bytes("sxv")
				var _1970_v144 D4
				_ = _1970_v144
				_1970_v144 = Companion_D4_.Create_DC13_((_dafny.Zero).Minus(_1934_v125), !_dafny.Companion_Sequence_.Contains(_1968_v142, _1967_v141), Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_1969_v143).Cardinality())), (_this).F6()), _1780_v0)
				var _1971_v145 bool
				_ = _1971_v145
				_1971_v145 = false
				var _1972_v146 _dafny.Sequence
				_ = _1972_v146
				_1972_v146 = _dafny.SeqOf((_this).F6(), _1934_v125)
				var _rhs273 _dafny.Int = ((_this).F6()).Plus((_this).F6())
				_ = _rhs273
				var _rhs274 D4 = Companion_Default___.Fm63(globalState)
				_ = _rhs274
				var _rhs275 _dafny.Int = (_1972_v146).Select((Companion_Default___.SafeIndex(_1934_v125, _dafny.IntOfUint32((_1972_v146).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs275
				var _rhs276 bool = (_this).F4()
				_ = _rhs276
				_1934_v125 = _rhs273
				_1970_v144 = _rhs274
				_1934_v125 = _rhs275
				_1971_v145 = _rhs276
				var _1973_v147 _dafny.MultiSet
				_ = _1973_v147
				_1973_v147 = _dafny.MultiSetOf((_this).F11())
				_1971_v145 = (!((_this).F11())) == ((_1973_v147).IsProperSubsetOf(_1973_v147))
				var _1974_v148 _dafny.Map
				_ = _1974_v148
				_1974_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1915_v108, _dafny.Companion_Sequence_.Concatenate(_1969_v143, _1969_v143))
				_1974_v148 = _1974_v148
				var _1975_v150 _dafny.Array
				_ = _1975_v150
				var _len0_50 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_50
				var _nw297 _dafny.Array
				_ = _nw297
				if _len0_50.Cmp(_dafny.Zero) == 0 {
					_nw297 = _dafny.NewArray(_len0_50)
				} else {
					var _init50 func(_dafny.Int) _dafny.MultiSet = (func(_1976_v125 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
						return func(_1977_i14 _dafny.Int) _dafny.MultiSet {
							return _dafny.MultiSetOf(func() _dafny.Set {
								var _coll56 = _dafny.NewBuilder()
								_ = _coll56
								for _iter63 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(528), _dafny.IntOfInt64(26))); ; {
									_compr_56, _ok63 := _iter63()
									if !_ok63 {
										break
									}
									var _1978_v149 _dafny.Int
									_1978_v149 = interface{}(_compr_56).(_dafny.Int)
									if ((_dafny.IntOfInt64(528)).Cmp(_1978_v149) <= 0) && ((_1978_v149).Cmp(_dafny.IntOfInt64(26)) < 0) {
										_coll56.Add(Companion_Default___.SafeDivisionInt(_1978_v149, _1976_v125))
									}
								}
								return _coll56.ToSet()
							}(), _dafny.SetOf((_this).F6()))
						}
					})(_1934_v125)
					_ = _init50
					var _element0_50 = _init50(_dafny.Zero)
					_ = _element0_50
					_nw297 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
					(_nw297).ArraySet1(_element0_50, 0)
					var _nativeLen0_50 = (_len0_50).Int()
					_ = _nativeLen0_50
					for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
						(_nw297).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
					}
				}
				_1975_v150 = _nw297
				var _1979_v151 _dafny.Map
				_ = _1979_v151
				_1979_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-561), (Companion_D3_.Create_DC9_(_1971_v145, (_this).Fm15(globalState), (_this).F11())).Dtor_cf8())
				var _1980_v152 _dafny.Set
				_ = _1980_v152
				_1980_v152 = _dafny.SetOf((_1979_v151).Cardinality())
				var _1981_v153 _dafny.MultiSet
				_ = _1981_v153
				_1981_v153 = _dafny.MultiSetOf(_1980_v152, _1980_v152)
				var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_1975_v150), 0))
				_ = _index289
				(_1975_v150).ArraySet1(_1981_v153, (_index289).Int())
				var _1982_v154 _dafny.Sequence
				_ = _1982_v154
				_1982_v154 = _dafny.SeqOf(_1980_v152, _1980_v152, _1980_v152)
				var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_1975_v150), 0))
				_ = _index290
				(_1975_v150).ArraySet1((_1981_v153).Update(((_1982_v154).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1969_v143).Cardinality()), _dafny.IntOfUint32((_1982_v154).Cardinality()))).Uint32()).(_dafny.Set)).Difference(_1980_v152), Companion_Default___.Abs((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1972_v146).Cardinality()), _1934_v125)))), (_index290).Int())
			} else {
				var _1983___mcc_h13 D3 = _source34.Get_().(D3_DC11).Cf12
				_ = _1983___mcc_h13
				var _1984_cf12 D3 = _1983___mcc_h13
				_ = _1984_cf12
				_1934_v125 = (_this).Fm15(globalState)
				var _1985_v155 _dafny.Map
				_ = _1985_v155
				_1985_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _dafny.IntOfInt64(725))
				var _1986_v156 _dafny.MultiSet
				_ = _1986_v156
				_1986_v156 = _dafny.MultiSetOf(_1985_v155, Companion_Default___.Fm44(globalState), _1985_v155, _1985_v155, _1985_v155)
				_1986_v156 = (_1986_v156).Union((_1986_v156).Union((_dafny.MultiSetOf(_1985_v155)).Update(_1985_v155, Companion_Default___.Abs(Companion_Default___.Fm3(_1934_v125, p0, globalState)))))
				_1915_v108 = _1915_v108
				_1985_v155 = _1985_v155
			}
			_1934_v125 = (func() _dafny.Int {
				if p0 {
					return (_dafny.Zero).Minus((_this).F6())
				}
				return ((_this).F6()).Plus((_this).F6())
			})()
			var _1987_v157 bool
			_ = _1987_v157
			_1987_v157 = true
			_1987_v157 = p0
		}
		var _1988_v158 _dafny.Sequence
		_ = _1988_v158
		_1988_v158 = _dafny.UnicodeSeqOfUtf8Bytes("ehhyianjb")
		var _1989_v159 *C2
		_ = _1989_v159
		var _nw298 *C2 = New_C2_()
		_ = _nw298
		_nw298.Ctor__(_1988_v158, p0)
		_1989_v159 = _nw298
		var _1990_v160 _dafny.Map
		_ = _1990_v160
		_1990_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(692), _1989_v159)
		r0 = Companion_Default___.Fm1(Companion_Default___.Fm3((_this).F6(), p0, globalState), ((_this).Fm7(globalState)).Cmp((_1990_v160).Cardinality()) <= 0, globalState)
		var _1991_v161 D1
		_ = _1991_v161
		_1991_v161 = Companion_D1_.Create_DC3_((_this).F4())
		var _1992_v162 _dafny.Map
		_ = _1992_v162
		_1992_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), p0)
		var _1993_v163 _dafny.Array
		_ = _1993_v163
		var _nwElement0_66 D1 = _1991_v161
		_ = _nwElement0_66
		var _nw299 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(29))
		_ = _nw299
		(_nw299).ArraySet1(_nwElement0_66, 0)
		(_nw299).ArraySet1(_1991_v161, 1)
		(_nw299).ArraySet1(_1991_v161, 2)
		(_nw299).ArraySet1(Companion_Default___.Fm55((func() bool {
			if (_1992_v162).Contains(p0) {
				return (_1992_v162).Get(p0).(bool)
			}
			return (_this).F11()
		})(), (_this).F4(), (_this).F6(), globalState), 3)
		(_nw299).ArraySet1(_1991_v161, 4)
		(_nw299).ArraySet1(_1991_v161, 5)
		(_nw299).ArraySet1(_1991_v161, 6)
		(_nw299).ArraySet1(_1991_v161, 7)
		(_nw299).ArraySet1(_1991_v161, 8)
		(_nw299).ArraySet1(_1991_v161, 9)
		(_nw299).ArraySet1(_1991_v161, 10)
		(_nw299).ArraySet1(_1991_v161, 11)
		(_nw299).ArraySet1(_1991_v161, 12)
		(_nw299).ArraySet1(Companion_D1_.Create_DC3_(Companion_Default___.Fm0(_1915_v108, globalState)), 13)
		(_nw299).ArraySet1(_1991_v161, 14)
		(_nw299).ArraySet1(_1991_v161, 15)
		(_nw299).ArraySet1(Companion_D1_.Create_DC3_((_this).F11()), 16)
		(_nw299).ArraySet1(_1991_v161, 17)
		(_nw299).ArraySet1(_1991_v161, 18)
		(_nw299).ArraySet1(_1991_v161, 19)
		(_nw299).ArraySet1(_1991_v161, 20)
		(_nw299).ArraySet1(_1991_v161, 21)
		(_nw299).ArraySet1(_1991_v161, 22)
		(_nw299).ArraySet1(_1991_v161, 23)
		(_nw299).ArraySet1(_1991_v161, 24)
		(_nw299).ArraySet1(_1991_v161, 25)
		(_nw299).ArraySet1(_1991_v161, 26)
		(_nw299).ArraySet1(_1991_v161, 27)
		(_nw299).ArraySet1(_1991_v161, 28)
		_1993_v163 = _nw299
		var _1994_v164 _dafny.Map
		_ = _1994_v164
		_1994_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1993_v163)
		var _1995_v165 _dafny.Map
		_ = _1995_v165
		_1995_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F6())
		var _1996_v166 _dafny.Map
		_ = _1996_v166
		_1996_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1995_v165, (_this).F4())
		r1 = (func() _dafny.Array {
			if (_1994_v164).Contains((func() bool {
				if (_1996_v166).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F6())) {
					return (_1996_v166).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F6())).(bool)
				}
				return (_this).F4()
			})()) {
				return (_1994_v164).Get((func() bool {
					if (_1996_v166).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F6())) {
						return (_1996_v166).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F6())).(bool)
					}
					return (_this).F4()
				})()).(_dafny.Array)
			}
			return _1993_v163
		})()
		return r0, r1
	}
}
func (_this *C6) M7(globalState *GlobalState) {
	{
		var _1997_v0 bool
		_ = _1997_v0
		_1997_v0 = true
		_1997_v0 = (_this).F11()
		var _1998_v1 _dafny.Int
		_ = _1998_v1
		_1998_v1 = _dafny.IntOfInt64(843)
		var _1999_v2 _dafny.Sequence
		_ = _1999_v2
		_1999_v2 = _dafny.SeqOf((_this).F6())
		var _2000_v3 _dafny.Sequence
		_ = _2000_v3
		_2000_v3 = _dafny.UnicodeSeqOfUtf8Bytes("gnnf")
		var _2001_v4 _dafny.Sequence
		_ = _2001_v4
		_2001_v4 = _dafny.SeqOf(true)
		var _2002_v5 _dafny.Sequence
		_ = _2002_v5
		_2002_v5 = _dafny.SeqOf(_2001_v4)
		var _2003_v6 D12
		_ = _2003_v6
		_2003_v6 = Companion_D12_.Create_DC27_(_2002_v5)
		var _2004_v7 _dafny.Sequence
		_ = _2004_v7
		_2004_v7 = _dafny.SeqOf(_1999_v2, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm39(_2000_v3, _2003_v6, (_this).F11(), globalState), _dafny.SeqOf(_1998_v1, (_this).F6())))
		_1998_v1 = _dafny.IntOfUint32(((_2004_v7).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_this).F4() {
				return (_this).F6()
			}
			return (_this).F6()
		})(), _dafny.IntOfUint32((_2004_v7).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
		var _2005_v8 _dafny.Array
		_ = _2005_v8
		var _nw300 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw300
		_2005_v8 = _nw300
		_2005_v8 = _2005_v8
		var _2006_v9 *C4
		_ = _2006_v9
		var _nw301 *C4 = New_C4_()
		_ = _nw301
		_nw301.Ctor__(_this.F5(), _1998_v1, (_this).F4())
		_2006_v9 = _nw301
		var _rhs277 _dafny.Int = Companion_Default___.SafeDivisionInt(_1998_v1, (_this).F6())
		_ = _rhs277
		var _rhs278 bool = (_dafny.IntOfUint32(((func() _dafny.Sequence {
			if _1997_v0 {
				return _2001_v4
			}
			return _dafny.Companion_Sequence_.Update(_2001_v4, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2001_v4).Cardinality()))).Uint32(), (_this).F11())
		})()).Cardinality())).Cmp(_1998_v1) != 0
		_ = _rhs278
		_1998_v1 = _rhs277
		_1997_v0 = _rhs278
		var _2007_v10 _dafny.Map
		_ = _2007_v10
		_2007_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1998_v1)
		var _2008_v11 *C1
		_ = _2008_v11
		var _nw302 *C1 = New_C1_()
		_ = _nw302
		_nw302.Ctor__(_1998_v1, ((func() _dafny.Int {
			if (_2007_v10).Contains(_1998_v1) {
				return (_2007_v10).Get(_1998_v1).(_dafny.Int)
			}
			return _1998_v1
		})()).Times(_dafny.IntOfUint32((_2000_v3).Cardinality())), _this.F5(), _1998_v1, !(_1997_v0))
		_2008_v11 = _nw302
	}
}
func (_this *C6) M8(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) (bool, bool, D1) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 D1 = Companion_D1_.Default()
		_ = r2
		var _2009_v0 _dafny.CodePoint
		_ = _2009_v0
		_2009_v0 = _dafny.CodePoint('a')
		var _2010_v1 _dafny.Array
		_ = _2010_v1
		var _nw303 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw303
		_2010_v1 = _nw303
		var _2011_v2 _dafny.Map
		_ = _2011_v2
		_2011_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2009_v0, _2010_v1)
		var _2012_v3 D19
		_ = _2012_v3
		_2012_v3 = Companion_D19_.Create_DC39_(Companion_D19_.Create_DC37_(_2011_v2))
		var _2013_v4 D19
		_ = _2013_v4
		_2013_v4 = Companion_D19_.Create_DC37_((_2011_v2).Update(_2009_v0, _2010_v1))
		_2012_v3 = Companion_D19_.Create_DC39_(_2013_v4)
		if (_this).F11() {
			var _2014_v5 _dafny.Int
			_ = _2014_v5
			_2014_v5 = _dafny.IntOfInt64(594)
			_2014_v5 = (p1).Plus(Companion_Default___.Fm3((_this).F6(), (_this).F11(), globalState))
			var _2015_v6 _dafny.Set
			_ = _2015_v6
			_2015_v6 = _dafny.SetOf(!((_this).F4()))
			_2015_v6 = _2015_v6
			r1 = ((_this).F11()) == ((func() bool {
				if (_this).F4() {
					return (_this).F11()
				}
				return (_this).F11()
			})())
			_2014_v5 = (_2014_v5).Plus((_this).F6())
			r0 = (_this).F4()
		} else {
			var _arr9 _dafny.Array = _this.F5()
			_ = _arr9
			var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_this.F5()), 0))
			_ = _index291
			(_arr9).ArraySet1(_2010_v1, (_index291).Int())
			var _arr10 _dafny.Array = _this.F5()
			_ = _arr10
			var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_this.F5()), 0))
			_ = _index292
			(_arr10).ArraySet1(_2010_v1, (_index292).Int())
			var _2016_v7 _dafny.Sequence
			_ = _2016_v7
			_2016_v7 = _dafny.SeqOf((_this).F11(), (_this).F4(), true, (_this).F4())
			var _2017_v8 _dafny.Sequence
			_ = _2017_v8
			_2017_v8 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_2016_v7, _2016_v7), _dafny.Companion_Sequence_.Concatenate(_2016_v7, _2016_v7), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F11()), _2016_v7))
			_2017_v8 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(108))).Uint32(), func(coer105 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}((func(_2018_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_2019_i0 _dafny.Int) _dafny.Sequence {
					return _2018_v7
				}
			})(_2016_v7)))
			var _2020_v9 _dafny.MultiSet
			_ = _2020_v9
			_2020_v9 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("brp"))
			var _2021_v10 _dafny.MultiSet
			_ = _2021_v10
			_2021_v10 = _dafny.MultiSetOf(_dafny.IntOfInt64(255), (func() _dafny.Int {
				if (_2020_v9).Contains(_dafny.UnicodeSeqOfUtf8Bytes("cufybjlks")) {
					return (_2020_v9).Multiplicity(_dafny.UnicodeSeqOfUtf8Bytes("cufybjlks"))
				}
				return p0
			})())
			var _2022_v11 D0
			_ = _2022_v11
			_2022_v11 = Companion_D0_.Create_DC0_((_this).F6())
			var _2023_v12 _dafny.Map
			_ = _2023_v12
			_2023_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2009_v0)
			var _2024_v13 _dafny.Map
			_ = _2024_v13
			_2024_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F4())
			var _2025_v14 bool
			_ = _2025_v14
			var _2026_v15 _dafny.Int
			_ = _2026_v15
			var _out58 bool
			_ = _out58
			var _out59 _dafny.Int
			_ = _out59
			_out58, _out59 = Companion_Default___.M0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2021_v10, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(577))).Uint32(), func(coer106 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg106 _dafny.Int) interface{} {
					return coer106(arg106)
				}
			}((func(_2027_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2028_i1 _dafny.Int) _dafny.CodePoint {
					return _2027_v0
				}
			})(_2009_v0)))), (_this).F4(), ((_this).Fm6(_2022_v11, _2022_v11, _2023_v12, p1, globalState)).Union(_2021_v10), !(false) || ((func() bool {
				if (_2024_v13).Contains((_this).F11()) {
					return (_2024_v13).Get((_this).F11()).(bool)
				}
				return (_this).F4()
			})()), globalState)
			_2025_v14 = _out58
			_2026_v15 = _out59
			var _2029_v16 _dafny.Sequence
			_ = _2029_v16
			_2029_v16 = _dafny.UnicodeSeqOfUtf8Bytes("wwdadof")
			var _2030_v17 _dafny.Map
			_ = _2030_v17
			_2030_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2026_v15, _2029_v16)
			var _2031_v18 D2
			_ = _2031_v18
			_2031_v18 = Companion_D2_.Create_DC6_()
			_2030_v17 = (_2030_v17).Update(Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfInt64(442)), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm16(_2025_v14, (_this).F11(), false, _2031_v18, globalState), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((Companion_Default___.Fm16(_2025_v14, (_this).F11(), false, _2031_v18, globalState)).Cardinality()))).Uint32(), _2009_v0), _2029_v16))
			var _2032_v19 _dafny.Array
			_ = _2032_v19
			var _nwElement0_67 _dafny.Array = _dafny.ArrayCastTo((_this.F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_this.F5()), 0))).Int()))
			_ = _nwElement0_67
			var _nw304 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(2))
			_ = _nw304
			(_nw304).ArraySet1(_nwElement0_67, 0)
			(_nw304).ArraySet1(_2010_v1, 1)
			_2032_v19 = _nw304
			var _2033_v20 _dafny.Array
			_ = _2033_v20
			_2033_v20 = _2032_v19
			var _2034_v21 *C5
			_ = _2034_v21
			var _nw305 *C5 = New_C5_()
			_ = _nw305
			_nw305.Ctor__((_this).F4(), (_dafny.Zero).Minus(p0), (_2033_v20), (_dafny.IntOfUint32((_2029_v16).Cardinality())).Times(p0), _2025_v14)
			_2034_v21 = _nw305
		}
		var _2035_v22 _dafny.Int
		_ = _2035_v22
		_2035_v22 = _dafny.IntOfInt64(861)
		_2035_v22 = (_2035_v22).Times((p1).Minus(p1))
		if (_this).F11() {
			var _2036_v23 _dafny.Sequence
			_ = _2036_v23
			_2036_v23 = _dafny.UnicodeSeqOfUtf8Bytes("senka")
			var _2037_v24 _dafny.Map
			_ = _2037_v24
			_2037_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F5(), _dafny.Companion_Sequence_.Equal(_2036_v23, _2036_v23))
			_2037_v24 = (_2037_v24).Update(_this.F5(), (_this).F4())
			var _2038_v25 _dafny.Set
			_ = _2038_v25
			_2038_v25 = _dafny.SetOf((_this).F4(), !((_this).F11()), (_this).F11())
			_2038_v25 = (_2038_v25).Intersection(_2038_v25)
			r0 = (_this).F4()
			(_this).M7(globalState)
			_2009_v0 = _2009_v0
		} else {
			var _2039_v26 D1
			_ = _2039_v26
			_2039_v26 = Companion_D1_.Create_DC3_(!(true))
			r2 = _2039_v26
			var _2040_v27 _dafny.Map
			_ = _2040_v27
			_2040_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F4()), (_this).F4())
			var _2041_v28 _dafny.Map
			_ = _2041_v28
			_2041_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (func() bool {
				if (_2040_v27).Contains((_this).F4()) {
					return (_2040_v27).Get((_this).F4()).(bool)
				}
				return !((_this).F11())
			})())
			_2041_v28 = (_2041_v28).Update(_2035_v22, (_this).F11())
			r0 = (_this).F11()
			var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_2010_v1), 0))
			_ = _index293
			(_2010_v1).ArraySet1((_this).F4(), (_index293).Int())
			var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_2010_v1), 0))
			_ = _index294
			(_2010_v1).ArraySet1((_this).F11(), (_index294).Int())
			var _2042_v29 _dafny.Sequence
			_ = _2042_v29
			_2042_v29 = _dafny.UnicodeSeqOfUtf8Bytes("pgj")
			var _2043_v30 *C2
			_ = _2043_v30
			var _nw306 *C2 = New_C2_()
			_ = _nw306
			_nw306.Ctor__(_2042_v29, (_2010_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_2010_v1), 0))).Int()).(bool))
			_2043_v30 = _nw306
		}
		var _2044_v31 _dafny.Sequence
		_ = _2044_v31
		_2044_v31 = _dafny.UnicodeSeqOfUtf8Bytes("a")
		var _2045_v32 T0
		_ = _2045_v32
		var _nw307 *C2 = New_C2_()
		_ = _nw307
		_nw307.Ctor__(_2044_v31, (_this).F11())
		_2045_v32 = _nw307
		_2045_v32 = _2045_v32
		var _2046_v33 _dafny.Map
		_ = _2046_v33
		_2046_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_this).F11() {
				return (_this).F11()
			}
			return (_this).F4()
		})(), (_this).F11())
		_2046_v33 = _2046_v33
		r0 = (_this).F4()
		r1 = ((_this).F6()).Cmp(((_this).F6()).Plus(_2035_v22)) >= 0
		var _2047_v34 D1
		_ = _2047_v34
		_2047_v34 = Companion_D1_.Create_DC3_((_this).F11())
		r2 = _2047_v34
		return r0, r1, r2
	}
}
func (_this *C6) F11() bool {
	{
		return _this._f11
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f5  _dafny.Array
	_f6  _dafny.Int
	_f4  bool
	_f10 _dafny.Array
	_f9  _dafny.Map
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.Zero
	_this._f4 = false
	_this._f10 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f9 = _dafny.EmptyMap
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

func (_this *C7) F5() _dafny.Array {
	return _this._f5
}
func (_this *C7) F5_set_(value _dafny.Array) {
	_this._f5 = value
}
func (_this *C7) F6() _dafny.Int {
	return _this._f6
}
func (_this *C7) F4() bool {
	return _this._f4
}
func (_this *C7) Ctor__(f9 _dafny.Map, f10 _dafny.Array, f5 _dafny.Array, f6 _dafny.Int, f4 bool) {
	{
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f4 = f4
	}
}
func (_this *C7) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		if (func() bool {
			if (_this).F4() {
				return (_this).F4()
			}
			return (_this).F4()
		})() {
			return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(799))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg107 _dafny.Int) interface{} {
					return coer107(arg107)
				}
			}(func(_2048_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			})), _dafny.SeqOf(_dafny.CodePoint('h')))
		} else {
			return ((_this).F6()).Cmp((_dafny.Zero).Minus((_this).F6())) == 0
		}
	}
}
func (_this *C7) Fm9(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F6()
	}
}
func (_this *C7) Fm6(p0 D0, p1 D0, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return ((func() _dafny.MultiSet {
			if false {
				return _dafny.MultiSetOf((_this).F6())
			}
			return _dafny.MultiSetOf((_this).F6())
		})()).Difference((_dafny.MultiSetOf((_this).F6())).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(472))))
	}
}
func (_this *C7) Fm7(globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F6(), (_this).F6()))).Minus(_dafny.IntOfInt64(677))
	}
}
func (_this *C7) Fm13(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F4()), _dafny.SeqOf(!((_this).F4()), (_this).F4(), (_this).F4()))
	}
}
func (_this *C7) Fm14(p0 _dafny.Map, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return ((_this).F6()).Cmp(_dafny.IntOfInt64(664)) == 0
	}
}
func (_this *C7) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _2049_v0 *C1
		_ = _2049_v0
		var _nw308 *C1 = New_C1_()
		_ = _nw308
		_nw308.Ctor__((_this).F6(), p0, _this.F5(), p0, (_this).F4())
		_2049_v0 = _nw308
		var _2050_v1 _dafny.Sequence
		_ = _2050_v1
		_2050_v1 = _dafny.SeqOf((_this).F4(), (_this).F4(), true, (_this).F4(), !((_this).F4()))
		var _source35 D4 = Companion_D4_.Create_DC13_(_dafny.IntOfInt64(194), !((_this).F4()), ((_this).F6()).Minus(_2049_v0.F16), _2050_v1)
		_ = _source35
		if _source35.Is_DC13() {
			var _2051___mcc_h0 _dafny.Int = _source35.Get_().(D4_DC13).Cf14
			_ = _2051___mcc_h0
			var _2052___mcc_h1 bool = _source35.Get_().(D4_DC13).Cf15
			_ = _2052___mcc_h1
			var _2053___mcc_h2 _dafny.Int = _source35.Get_().(D4_DC13).Cf16
			_ = _2053___mcc_h2
			var _2054___mcc_h3 _dafny.Sequence = _source35.Get_().(D4_DC13).Cf17
			_ = _2054___mcc_h3
			var _2055_cf17 _dafny.Sequence = _2054___mcc_h3
			_ = _2055_cf17
			var _2056_cf16 _dafny.Int = _2053___mcc_h2
			_ = _2056_cf16
			var _2057_cf15 bool = _2052___mcc_h1
			_ = _2057_cf15
			var _2058_cf14 _dafny.Int = _2051___mcc_h0
			_ = _2058_cf14
			var _2059_v2 _dafny.Set
			_ = _2059_v2
			_2059_v2 = _dafny.SetOf((_dafny.Zero).Minus((Companion_Default___.Fm64(p0, (_this).F6(), _2057_cf15, globalState)).Cardinality()))
			var _2060_v3 D9
			_ = _2060_v3
			_2060_v3 = Companion_D9_.Create_DC20_(_2057_cf15, _dafny.CodePoint('n'), _2049_v0.F16)
			var _2061_v4 *C6
			_ = _2061_v4
			var _nw309 *C6 = New_C6_()
			_ = _nw309
			_nw309.Ctor__(((_2059_v2).Cardinality()).Cmp(_dafny.IntOfInt64(272)) < 0, (_this).F4(), _this.F5(), (_2060_v3).Dtor_cf28())
			_2061_v4 = _nw309
			_2057_cf15 = true
			_2057_cf15 = !(!((_2061_v4).F11()))
			(_2049_v0).F16 = ((_2049_v0).F15()).Minus(_dafny.IntOfInt64(-913))
		} else {
			var _2062___mcc_h4 _dafny.CodePoint = _source35.Get_().(D4_DC12).Cf13
			_ = _2062___mcc_h4
			var _2063_cf13 _dafny.CodePoint = _2062___mcc_h4
			_ = _2063_cf13
			var _2064_v5 _dafny.Array
			_ = _2064_v5
			var _nw310 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw310
			_2064_v5 = _nw310
			var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_2064_v5), 0))
			_ = _index295
			(_2064_v5).ArraySet1(_dafny.IntOfUint32((_2050_v1).Cardinality()), (_index295).Int())
			var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_2064_v5), 0))
			_ = _index296
			(_2064_v5).ArraySet1((_2049_v0).F15(), (_index296).Int())
			var _2065_v6 _dafny.Array
			_ = _2065_v6
			var _len0_51 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_51
			var _nw311 _dafny.Array
			_ = _nw311
			if _len0_51.Cmp(_dafny.Zero) == 0 {
				_nw311 = _dafny.NewArray(_len0_51)
			} else {
				var _init51 func(_dafny.Int) _dafny.MultiSet = (func(_2066_v0 *C1, _2067_p0 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
					return func(_2068_i0 _dafny.Int) _dafny.MultiSet {
						return (_dafny.MultiSetOf(_2066_v0.F16)).Union(_dafny.MultiSetOf(_2067_p0))
					}
				})(_2049_v0, p0)
				_ = _init51
				var _element0_51 = _init51(_dafny.Zero)
				_ = _element0_51
				_nw311 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
				(_nw311).ArraySet1(_element0_51, 0)
				var _nativeLen0_51 = (_len0_51).Int()
				_ = _nativeLen0_51
				for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
					(_nw311).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
				}
			}
			_2065_v6 = _nw311
			var _2069_v7 _dafny.MultiSet
			_ = _2069_v7
			_2069_v7 = _dafny.MultiSetOf(p0)
			var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_2065_v6), 0))
			_ = _index297
			(_2065_v6).ArraySet1((func() _dafny.MultiSet {
				if (_this).F4() {
					return _2069_v7
				}
				return _2069_v7
			})(), (_index297).Int())
			var _2070_v8 bool
			_ = _2070_v8
			_2070_v8 = true
			var _2071_v9 _dafny.Map
			_ = _2071_v9
			_2071_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2049_v0).F15(), true)
			var _2072_v10 *C3
			_ = _2072_v10
			var _nw312 *C3 = New_C3_()
			_ = _nw312
			_nw312.Ctor__(_this.F5(), (_2049_v0).F15(), (func() bool {
				if (_this).Fm14(_2071_v9, true, _2049_v0.F16, false, globalState) {
					return (_this).F4()
				}
				return _2070_v8
			})())
			_2072_v10 = _nw312
			var _2073_v11 _dafny.Map
			_ = _2073_v11
			_2073_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2070_v8, _2070_v8)
			var _2074_v12 _dafny.MultiSet
			_ = _2074_v12
			_2074_v12 = _dafny.MultiSetOf(_2070_v8, (func() bool {
				if (_2073_v11).Contains((_this).F4()) {
					return (_2073_v11).Get((_this).F4()).(bool)
				}
				return _2070_v8
			})())
			var _2075_v13 _dafny.Sequence
			_ = _2075_v13
			_2075_v13 = _dafny.UnicodeSeqOfUtf8Bytes("sdatbjvc")
			var _2076_v14 *C2
			_ = _2076_v14
			var _nw313 *C2 = New_C2_()
			_ = _nw313
			_nw313.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("xyw"), (_this).F4())
			_2076_v14 = _nw313
			var _2077_v15 _dafny.Map
			_ = _2077_v15
			_2077_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2049_v0.F16, _2076_v14)
			var _2078_v16 D16
			_ = _2078_v16
			_2078_v16 = Companion_D16_.Create_DC33_(false, (_this).F6(), _2074_v12, _dafny.IntOfUint32((_2075_v13).Cardinality()), Companion_Default___.Fm23((_2077_v15).Cardinality(), _2049_v0.F16, globalState))
			var _2079_v17 D9
			_ = _2079_v17
			_2079_v17 = Companion_D9_.Create_DC20_((_2078_v16).Dtor_cf52(), _2063_cf13, _2049_v0.F16)
			var _pat_let_tv39 = _2073_v11
			_ = _pat_let_tv39
			var _2080_v18 _dafny.Array
			_ = _2080_v18
			var _nwElement0_68 _dafny.CodePoint = _2063_cf13
			_ = _nwElement0_68
			var _nw314 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(6))
			_ = _nw314
			(_nw314).ArraySet1CodePoint(_nwElement0_68, 0)
			(_nw314).ArraySet1CodePoint(_2063_cf13, 1)
			(_nw314).ArraySet1CodePoint((func(_pat_let48_0 D9) D9 {
				return func(_2081_dt__update__tmp_h0 D9) D9 {
					return func(_pat_let49_0 bool) D9 {
						return func(_2082_dt__update_hcf26_h0 bool) D9 {
							return func(_pat_let50_0 _dafny.Int) D9 {
								return func(_2083_dt__update_hcf28_h0 _dafny.Int) D9 {
									return Companion_D9_.Create_DC20_(_2082_dt__update_hcf26_h0, (_2081_dt__update__tmp_h0).Dtor_cf27(), _2083_dt__update_hcf28_h0)
								}(_pat_let50_0)
							}((_pat_let_tv39).Cardinality())
						}(_pat_let49_0)
					}((_this).F4())
				}(_pat_let48_0)
			}(_2079_v17)).Dtor_cf27(), 2)
			(_nw314).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_this).F4() {
					return _2063_cf13
				}
				return _2063_cf13
			})(), 3)
			(_nw314).ArraySet1CodePoint(_2063_cf13, 4)
			(_nw314).ArraySet1CodePoint(_2063_cf13, 5)
			_2080_v18 = _nw314
			var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_2065_v6), 0))
			_ = _index298
			var _rhs279 _dafny.MultiSet = _2069_v7
			_ = _rhs279
			var _rhs280 bool = (((_this).F4()) || (true)) && (((_2049_v0).F15()).Cmp(_2049_v0.F16) != 0)
			_ = _rhs280
			var _rhs281 *C3 = _2072_v10
			_ = _rhs281
			var _rhs282 _dafny.Array = _2080_v18
			_ = _rhs282
			var _lhs169 _dafny.Array = _2065_v6
			_ = _lhs169
			var _lhs170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_2065_v6), 0))
			_ = _lhs170
			(_lhs169).ArraySet1(_rhs279, (_lhs170).Int())
			_2070_v8 = _rhs280
			_2072_v10 = _rhs281
			_2080_v18 = _rhs282
			(_2049_v0).F16 = _2049_v0.F16
			var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_2064_v5), 0))
			_ = _index299
			var _rhs283 _dafny.Int = (Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(989), (_2049_v0).F15())).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cfu")).Cardinality()))
			_ = _rhs283
			var _rhs284 _dafny.Int = _dafny.IntOfInt64(196)
			_ = _rhs284
			var _rhs285 _dafny.Int = (_dafny.Zero).Minus(_2049_v0.F16)
			_ = _rhs285
			var _lhs171 *C1 = _2049_v0
			_ = _lhs171
			var _lhs172 *C1 = _2049_v0
			_ = _lhs172
			var _lhs173 _dafny.Array = _2064_v5
			_ = _lhs173
			var _lhs174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_2064_v5), 0))
			_ = _lhs174
			_lhs171.F16 = _rhs283
			_lhs172.F16 = _rhs284
			(_lhs173).ArraySet1(_rhs285, (_lhs174).Int())
		}
		var _2084_v19 _dafny.Map
		_ = _2084_v19
		_2084_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2049_v0.F16)
		var _2085_v20 _dafny.Map
		_ = _2085_v20
		_2085_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _2084_v19)
		var _2086_v21 _dafny.Sequence
		_ = _2086_v21
		_2086_v21 = _dafny.SeqOf(_2084_v19, (func() _dafny.Map {
			if (_2085_v20).Contains(Companion_Default___.Fm0(_dafny.CodePoint('a'), globalState)) {
				return (_2085_v20).Get(Companion_Default___.Fm0(_dafny.CodePoint('a'), globalState)).(_dafny.Map)
			}
			return _2084_v19
		})())
		var _2087_v22 _dafny.MultiSet
		_ = _2087_v22
		_2087_v22 = _dafny.MultiSetOf(((_2086_v21).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_2086_v21).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _2049_v0.F16, p0)
		var _2088_v23 _dafny.Map
		_ = _2088_v23
		_2088_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2087_v22, _dafny.UnicodeSeqOfUtf8Bytes("rcxfakrum"))
		var _2089_v24 bool
		_ = _2089_v24
		var _2090_v25 _dafny.Int
		_ = _2090_v25
		var _out60 bool
		_ = _out60
		var _out61 _dafny.Int
		_ = _out61
		_out60, _out61 = Companion_Default___.M0(_2088_v23, (_this).F4(), _2087_v22, (_this).F4(), globalState)
		_2089_v24 = _out60
		_2090_v25 = _out61
		var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen(((_this).F10()), 0))
		_ = _index300
		((_this).F10()).ArraySet1(_2089_v24, (_index300).Int())
		var _2091_v26 D11
		_ = _2091_v26
		_2091_v26 = Companion_D11_.Create_DC26_(!((_this).F4()))
		var _2092_v27 _dafny.Set
		_ = _2092_v27
		_2092_v27 = _dafny.SetOf(_2091_v26, _2091_v26, _2091_v26, _2091_v26)
		var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen(((_this).F10()), 0))
		_ = _index301
		((_this).F10()).ArraySet1((_2090_v25).Cmp(((_2092_v27).Intersection(_2092_v27)).Cardinality()) != 0, (_index301).Int())
		var _2093_v28 _dafny.Sequence
		_ = _2093_v28
		_2093_v28 = _dafny.UnicodeSeqOfUtf8Bytes("lvmi")
		var _2094_v29 _dafny.Map
		_ = _2094_v29
		_2094_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))
		var _2095_v30 _dafny.Sequence
		_ = _2095_v30
		_2095_v30 = _dafny.SeqOf((_this).F6(), _dafny.IntOfUint32((_2050_v1).Cardinality()))
		var _2096_v31 _dafny.Sequence
		_ = _2096_v31
		_2096_v31 = _dafny.SeqOf(_2049_v0.F16, _dafny.IntOfUint32((_2095_v30).Cardinality()))
		var _2097_v32 _dafny.Array
		_ = _2097_v32
		var _nwElement0_69 _dafny.Int = (_this).Fm7(globalState)
		_ = _nwElement0_69
		var _nw315 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(14))
		_ = _nw315
		(_nw315).ArraySet1(_nwElement0_69, 0)
		(_nw315).ArraySet1((_2049_v0).F15(), 1)
		(_nw315).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_2093_v28).Cardinality()), (_2049_v0).F15()), 2)
		(_nw315).ArraySet1((_2094_v29).Cardinality(), 3)
		(_nw315).ArraySet1((_this).F6(), 4)
		(_nw315).ArraySet1(Companion_Default___.SafeModuloInt(_2090_v25, _dafny.IntOfUint32((_2050_v1).Cardinality())), 5)
		(_nw315).ArraySet1(((_this).F6()).Times(p0), 6)
		(_nw315).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(34), _dafny.IntOfInt64(583)), 7)
		(_nw315).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_2096_v31).Cardinality()), _dafny.IntOfInt64(967)), 8)
		(_nw315).ArraySet1(_2049_v0.F16, 9)
		(_nw315).ArraySet1((_2049_v0).F15(), 10)
		(_nw315).ArraySet1(_2049_v0.F16, 11)
		(_nw315).ArraySet1(_2090_v25, 12)
		(_nw315).ArraySet1((_2049_v0).F15(), 13)
		_2097_v32 = _nw315
		var _2098_v33 _dafny.Map
		_ = _2098_v33
		_2098_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(30), ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))
		var _2099_v34 _dafny.Array
		_ = _2099_v34
		var _nw316 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
		_ = _nw316
		_2099_v34 = _nw316
		var _2100_v35 D10
		_ = _2100_v35
		_2100_v35 = Companion_D10_.Create_DC22_(_dafny.UnicodeSeqOfUtf8Bytes("tsfhsrurn"), _2089_v24, (_this).F4(), _2099_v34, (_this).F6())
		var _2101_v36 _dafny.MultiSet
		_ = _2101_v36
		_2101_v36 = _dafny.MultiSetOf((_this).F4(), _2089_v24)
		var _2102_v37 *C4
		_ = _2102_v37
		var _nw317 *C4 = New_C4_()
		_ = _nw317
		_nw317.Ctor__(_this.F5(), _2090_v25, _2089_v24)
		_2102_v37 = _nw317
		var _2103_v38 _dafny.Map
		_ = _2103_v38
		_2103_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2102_v37, _2093_v28)
		var _2104_v39 _dafny.Set
		_ = _2104_v39
		_2104_v39 = _dafny.SetOf((_this).F4(), _2089_v24)
		var _2105_v40 D16
		_ = _2105_v40
		_2105_v40 = Companion_D16_.Create_DC33_((func() bool {
			if (_2098_v33).Contains(_2049_v0.F16) {
				return (_2098_v33).Get(_2049_v0.F16).(bool)
			}
			return ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool)
		})(), (_2102_v37).Fm7(globalState), _dafny.MultiSetOf(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool)), (_2087_v22).Cardinality(), _2104_v39)
		var _2106_v41 D16
		_ = _2106_v41
		_2106_v41 = Companion_D16_.Create_DC33_((_this).Fm14(_2098_v33, ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), (_2049_v0).F15(), (_2100_v35).Dtor_cf32(), globalState), (_dafny.Zero).Minus(((func() _dafny.Int {
			if (_2084_v19).Contains(_2049_v0.F16) {
				return (_2084_v19).Get(_2049_v0.F16).(_dafny.Int)
			}
			return (_2087_v22).Cardinality()
		})()).Plus(_dafny.IntOfInt64(431))), ((_2101_v36).Update(_2089_v24, Companion_Default___.Abs((_2103_v38).Cardinality()))).Intersection(_2101_v36), (_dafny.Zero).Minus((_2049_v0.F16).Times(_2049_v0.F16)), _dafny.SetOf(_2089_v24, _2089_v24, true, !(_2089_v24), !((_2105_v40).Dtor_cf52())))
		var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen(((_this).F10()), 0))
		_ = _index302
		var _rhs286 _dafny.Int = _dafny.IntOfInt64(154)
		_ = _rhs286
		var _rhs287 _dafny.Int = _dafny.IntOfInt64(842)
		_ = _rhs287
		var _rhs288 _dafny.Array = _2097_v32
		_ = _rhs288
		var _rhs289 D16 = _2105_v40
		_ = _rhs289
		var _rhs290 bool = ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool)
		_ = _rhs290
		var _lhs175 *C1 = _2049_v0
		_ = _lhs175
		var _lhs176 _dafny.Array = (_this).F10()
		_ = _lhs176
		var _lhs177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen(((_this).F10()), 0))
		_ = _lhs177
		_lhs175.F16 = _rhs286
		_2090_v25 = _rhs287
		_2097_v32 = _rhs288
		_2105_v40 = _rhs289
		(_lhs176).ArraySet1(_rhs290, (_lhs177).Int())
		var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen(((_this).F10()), 0))
		_ = _index303
		var _rhs291 _dafny.Int = p0
		_ = _rhs291
		var _rhs292 bool = (func() bool {
			if (_2090_v25).Cmp(_2049_v0.F16) == 0 {
				return false
			}
			return false
		})()
		_ = _rhs292
		var _lhs178 *C1 = _2049_v0
		_ = _lhs178
		var _lhs179 _dafny.Array = (_this).F10()
		_ = _lhs179
		var _lhs180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen(((_this).F10()), 0))
		_ = _lhs180
		_lhs178.F16 = _rhs291
		(_lhs179).ArraySet1(_rhs292, (_lhs180).Int())
		var _2107_v42 _dafny.Map
		_ = _2107_v42
		_2107_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), p0)
		r0 = _2107_v42
		return r0
	}
}
func (_this *C7) M2(p0 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Array) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _2108_v0 _dafny.Int
		_ = _2108_v0
		_2108_v0 = _dafny.IntOfInt64(-830)
		var _2109_v1 _dafny.Sequence
		_ = _2109_v1
		_2109_v1 = _dafny.UnicodeSeqOfUtf8Bytes("q")
		var _2110_v2 *C3
		_ = _2110_v2
		var _nw318 *C3 = New_C3_()
		_ = _nw318
		_nw318.Ctor__(_this.F5(), _dafny.IntOfUint32((_2109_v1).Cardinality()), p0)
		_2110_v2 = _nw318
		var _2111_v3 _dafny.Map
		_ = _2111_v3
		_2111_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2110_v2, false)
		_2108_v0 = (Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2111_v3, (_this).F6())).Cardinality(), _2108_v0)).Times((_this).F6())
		var _2112_i0 _dafny.Int
		_ = _2112_i0
		_2112_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_2112_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_2112_i0 = (_2112_i0).Plus(_dafny.One)
					var _2113_v4 _dafny.Array
					_ = _2113_v4
					var _nw319 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
					_ = _nw319
					_2113_v4 = _nw319
					var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_2113_v4), 0))
					_ = _index304
					(_2113_v4).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(578))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg108 _dafny.Int) interface{} {
							return coer108(arg108)
						}
					}(func(_2114_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('o')
					})), (_index304).Int())
					var _2115_v5 bool
					_ = _2115_v5
					_2115_v5 = false
					var _2116_v6 _dafny.Array
					_ = _2116_v6
					var _len0_52 _dafny.Int = _dafny.IntOfInt64(13)
					_ = _len0_52
					var _nw320 _dafny.Array
					_ = _nw320
					if _len0_52.Cmp(_dafny.Zero) == 0 {
						_nw320 = _dafny.NewArray(_len0_52)
					} else {
						var _init52 func(_dafny.Int) _dafny.CodePoint = func(_2117_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('d')
						}
						_ = _init52
						var _element0_52 = _init52(_dafny.Zero)
						_ = _element0_52
						_nw320 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
						(_nw320).ArraySet1CodePoint(_element0_52, 0)
						var _nativeLen0_52 = (_len0_52).Int()
						_ = _nativeLen0_52
						for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
							(_nw320).ArraySet1CodePoint(_init52(_dafny.IntOf(_i0_52)), _i0_52)
						}
					}
					_2116_v6 = _nw320
					var _2118_v7 _dafny.CodePoint
					_ = _2118_v7
					_2118_v7 = _dafny.CodePoint('h')
					var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_2116_v6), 0))
					_ = _index305
					(_2116_v6).ArraySet1CodePoint(_2118_v7, (_index305).Int())
					var _2119_v8 _dafny.Set
					_ = _2119_v8
					_2119_v8 = _dafny.SetOf(_dafny.IntOfInt64(474), (_this).F6(), (_this).F6(), (_dafny.Zero).Minus((_this).Fm7(globalState)))
					var _2120_v9 _dafny.Sequence
					_ = _2120_v9
					_2120_v9 = _dafny.SeqOf(!((_this).F4()))
					var _2121_v10 _dafny.Sequence
					_ = _2121_v10
					_2121_v10 = _dafny.SeqOf(_2109_v1)
					var _2122_v11 _dafny.Map
					_ = _2122_v11
					_2122_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _2121_v10)
					var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_2113_v4), 0))
					_ = _index306
					var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_2116_v6), 0))
					_ = _index307
					var _rhs293 _dafny.Sequence = _2109_v1
					_ = _rhs293
					var _rhs294 bool = _dafny.Companion_Sequence_.Contains(_2120_v9, true)
					_ = _rhs294
					var _rhs295 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_2122_v11).Contains((_this).F10()) {
							return (_2122_v11).Get((_this).F10()).(_dafny.Sequence)
						}
						return _dafny.Companion_Sequence_.Concatenate(_2121_v10, _2121_v10)
					})()).Cardinality())
					_ = _rhs295
					var _rhs296 _dafny.CodePoint = _2118_v7
					_ = _rhs296
					var _rhs297 _dafny.Set = _2119_v8
					_ = _rhs297
					var _lhs181 _dafny.Array = _2113_v4
					_ = _lhs181
					var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_2113_v4), 0))
					_ = _lhs182
					var _lhs183 _dafny.Array = _2116_v6
					_ = _lhs183
					var _lhs184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_2116_v6), 0))
					_ = _lhs184
					(_lhs181).ArraySet1(_rhs293, (_lhs182).Int())
					_2115_v5 = _rhs294
					_2108_v0 = _rhs295
					(_lhs183).ArraySet1CodePoint(_rhs296, (_lhs184).Int())
					_2119_v8 = _rhs297
					_2108_v0 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(404), _dafny.IntOfUint32((_2109_v1).Cardinality()))
					var _2123_v12 _dafny.Map
					_ = _2123_v12
					_2123_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2118_v7)
					var _2124_v13 _dafny.Sequence
					_ = _2124_v13
					_2124_v13 = _dafny.SeqOf((Companion_Default___.Fm65(_2108_v0, _dafny.IntOfInt64(625), _dafny.IntOfInt64(209), globalState)).Update(true, (_2116_v6).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_2116_v6), 0))).Int())), _2123_v12)
					_2108_v0 = (_dafny.Zero).Minus((((_2123_v12).Update(_2115_v5, (_2116_v6).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_2116_v6), 0))).Int()))).Merge((_2123_v12).Merge((_2124_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2109_v1).Cardinality()), _dafny.IntOfUint32((_2124_v13).Cardinality()))).Uint32()).(_dafny.Map)))).Cardinality())
					_2115_v5 = p0
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _2125_v14 T0
		_ = _2125_v14
		var _nw321 *C6 = New_C6_()
		_ = _nw321
		_nw321.Ctor__((_this).F4(), p0, _this.F5(), _2108_v0)
		_2125_v14 = _nw321
		var _2126_v15 _dafny.Map
		_ = _2126_v15
		_2126_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _2125_v14)
		var _2127_v16 _dafny.Map
		_ = _2127_v16
		_2127_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
		var _2128_v17 D3
		_ = _2128_v17
		_2128_v17 = Companion_D3_.Create_DC9_(false, ((_2126_v15).Cardinality()).Plus((_2127_v16).Cardinality()), (_2125_v14).F4())
		var _source36 D3 = _2128_v17
		_ = _source36
		if _source36.Is_DC9() {
			var _2129___mcc_h0 bool = _source36.Get_().(D3_DC9).Cf8
			_ = _2129___mcc_h0
			var _2130___mcc_h1 _dafny.Int = _source36.Get_().(D3_DC9).Cf9
			_ = _2130___mcc_h1
			var _2131___mcc_h2 bool = _source36.Get_().(D3_DC9).Cf10
			_ = _2131___mcc_h2
			var _2132_cf10 bool = _2131___mcc_h2
			_ = _2132_cf10
			var _2133_cf9 _dafny.Int = _2130___mcc_h1
			_ = _2133_cf9
			var _2134_cf8 bool = _2129___mcc_h0
			_ = _2134_cf8
			var _2135_v18 _dafny.Set
			_ = _2135_v18
			_2135_v18 = _dafny.SetOf((_2125_v14).F4(), p0, (_this).F4())
			var _2136_v19 _dafny.Array
			_ = _2136_v19
			var _len0_53 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_53
			var _nw322 _dafny.Array
			_ = _nw322
			if _len0_53.Cmp(_dafny.Zero) == 0 {
				_nw322 = _dafny.NewArray(_len0_53)
			} else {
				var _init53 func(_dafny.Int) _dafny.Int = (func(_2137_cf9 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2138_i3 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_2138_i3, _2137_cf9)
					}
				})(_2133_cf9)
				_ = _init53
				var _element0_53 = _init53(_dafny.Zero)
				_ = _element0_53
				_nw322 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
				(_nw322).ArraySet1(_element0_53, 0)
				var _nativeLen0_53 = (_len0_53).Int()
				_ = _nativeLen0_53
				for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
					(_nw322).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
				}
			}
			_2136_v19 = _nw322
			var _2139_v20 _dafny.Map
			_ = _2139_v20
			_2139_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2136_v19, _2135_v18)
			var _2140_v21 _dafny.CodePoint
			_ = _2140_v21
			_2140_v21 = _dafny.CodePoint('x')
			var _2141_v22 _dafny.Map
			_ = _2141_v22
			_2141_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_2125_v14).F4())
			var _2142_v23 _dafny.Array
			_ = _2142_v23
			var _nwElement0_70 _dafny.Set = _2135_v18
			_ = _nwElement0_70
			var _nw323 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(25))
			_ = _nw323
			(_nw323).ArraySet1(_nwElement0_70, 0)
			(_nw323).ArraySet1(_dafny.SetOf(true), 1)
			(_nw323).ArraySet1(_2135_v18, 2)
			(_nw323).ArraySet1(_2135_v18, 3)
			(_nw323).ArraySet1((_dafny.SetOf(_2132_cf10, p0, (_2125_v14).F4())).Difference(_2135_v18), 4)
			(_nw323).ArraySet1((func() _dafny.Set {
				if _2134_cf8 {
					return _2135_v18
				}
				return _2135_v18
			})(), 5)
			(_nw323).ArraySet1(_2135_v18, 6)
			(_nw323).ArraySet1(_2135_v18, 7)
			(_nw323).ArraySet1(_dafny.SetOf(_2134_cf8), 8)
			(_nw323).ArraySet1(_2135_v18, 9)
			(_nw323).ArraySet1(_2135_v18, 10)
			(_nw323).ArraySet1(_2135_v18, 11)
			(_nw323).ArraySet1(_dafny.SetOf(p0, (_2125_v14).F4(), (_2125_v14).F4(), _2132_cf10, (_2125_v14).F4()), 12)
			(_nw323).ArraySet1(_2135_v18, 13)
			(_nw323).ArraySet1(_2135_v18, 14)
			(_nw323).ArraySet1(_2135_v18, 15)
			(_nw323).ArraySet1(_2135_v18, 16)
			(_nw323).ArraySet1((_2135_v18).Difference(_2135_v18), 17)
			(_nw323).ArraySet1(_2135_v18, 18)
			(_nw323).ArraySet1(_2135_v18, 19)
			(_nw323).ArraySet1(((func() _dafny.Set {
				if (_2139_v20).Contains(_2136_v19) {
					return (_2139_v20).Get(_2136_v19).(_dafny.Set)
				}
				return _2135_v18
			})()).Union(Companion_Default___.Fm46(_2140_v21, globalState)), 20)
			(_nw323).ArraySet1(_2135_v18, 21)
			(_nw323).ArraySet1((_2135_v18).Intersection(_2135_v18), 22)
			(_nw323).ArraySet1(_2135_v18, 23)
			(_nw323).ArraySet1((_2135_v18).Union(Companion_Default___.Fm53((_this).F6(), _dafny.SeqOf(_2141_v22, _2141_v22), globalState)), 24)
			_2142_v23 = _nw323
			var _2143_v24 _dafny.MultiSet
			_ = _2143_v24
			_2143_v24 = _dafny.MultiSetOf((_2125_v14).F4())
			var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_2142_v23), 0))
			_ = _index308
			(_2142_v23).ArraySet1(Companion_Default___.Fm53(_2133_cf9, Companion_Default___.Fm54((_2143_v24).Cardinality(), globalState), globalState), (_index308).Int())
			var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_2142_v23), 0))
			_ = _index309
			(_2142_v23).ArraySet1(_2135_v18, (_index309).Int())
			var _2144_v25 _dafny.Array
			_ = _2144_v25
			_2144_v25 = _this.F5()
			var _2145_v26 *C5
			_ = _2145_v26
			var _nw324 *C5 = New_C5_()
			_ = _nw324
			_nw324.Ctor__(_2132_cf10, (_this).F6(), (_2144_v25), _2133_cf9, (_2125_v14).F4())
			_2145_v26 = _nw324
			if (!((_2145_v26).F13())) || ((_2108_v0).Cmp(_2108_v0) < 0) {
				var _2146_v27 _dafny.Array
				_ = _2146_v27
				var _nw325 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(25))
				_ = _nw325
				_2146_v27 = _nw325
				var _2147_v28 _dafny.Array
				_ = _2147_v28
				var _nw326 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
				_ = _nw326
				_2147_v28 = _nw326
				var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_2136_v19), 0))
				_ = _index310
				(_2136_v19).ArraySet1(Companion_Default___.Fm3(_2108_v0, _2134_cf8, globalState), (_index310).Int())
				var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_2136_v19), 0))
				_ = _index311
				var _rhs298 _dafny.Array = _2146_v27
				_ = _rhs298
				var _rhs299 _dafny.Array = _2147_v28
				_ = _rhs299
				var _rhs300 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_2108_v0), _dafny.IntOfInt64(61)))
				_ = _rhs300
				var _lhs185 _dafny.Array = _2136_v19
				_ = _lhs185
				var _lhs186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_2136_v19), 0))
				_ = _lhs186
				_2146_v27 = _rhs298
				_2147_v28 = _rhs299
				(_lhs185).ArraySet1(_rhs300, (_lhs186).Int())
				var _2148_v29 _dafny.MultiSet
				_ = _2148_v29
				_2148_v29 = _dafny.MultiSetOf((_this).F10())
				_2148_v29 = _dafny.MultiSetOf((_this).F10())
				_2132_cf10 = ((_2108_v0).Plus((_this).F6())).Cmp(((_dafny.Zero).Minus((_dafny.Zero).Minus((_2136_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_2136_v19), 0))).Int()).(_dafny.Int)))).Times(_2145_v26.F14)) < 0
				var _2149_v30 _dafny.Sequence
				_ = _2149_v30
				_2149_v30 = _dafny.SeqOf((_dafny.Zero).Minus(_2145_v26.F14))
				_2108_v0 = (_2149_v30).Select((Companion_Default___.SafeIndex(_2145_v26.F14, _dafny.IntOfUint32((_2149_v30).Cardinality()))).Uint32()).(_dafny.Int)
				var _2150_v31 *C4
				_ = _2150_v31
				var _nw327 *C4 = New_C4_()
				_ = _nw327
				_nw327.Ctor__(_this.F5(), _2145_v26.F14, _2134_cf8)
				_2150_v31 = _nw327
			} else {
				var _2151_v32 _dafny.Map
				_ = _2151_v32
				_2151_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2108_v0)
				_2127_v16 = (_2127_v16).Update(((_this).F6()).Plus((_2151_v32).Cardinality()), _2133_cf9)
				_2108_v0 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_2145_v26.F14), _2145_v26.F14)
				_2108_v0 = _dafny.IntOfUint32((_2109_v1).Cardinality())
				var _2152_v33 D2
				_ = _2152_v33
				_2152_v33 = Companion_D2_.Create_DC5_((_this).F10())
				_2152_v33 = _2152_v33
				var _2153_v34 _dafny.Map
				_ = _2153_v34
				_2153_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2145_v26.F14, _2132_cf10)
				_2153_v34 = (_2153_v34).Update(((_this).F6()).Minus((_this).F6()), !((_2125_v14).F4()))
			}
			_2132_cf10 = !((_2125_v14).F4())
		} else if _source36.Is_DC10() {
			var _2154___mcc_h3 _dafny.Int = _source36.Get_().(D3_DC10).Cf11
			_ = _2154___mcc_h3
			var _2155_cf11 _dafny.Int = _2154___mcc_h3
			_ = _2155_cf11
			var _2156_v35 _dafny.Map
			_ = _2156_v35
			_2156_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F4())).Cardinality())), (_this).F10())
			_2156_v35 = (_2156_v35).Update((_this).F6(), (_this).F10())
			var _2157_v36 D12
			_ = _2157_v36
			_2157_v36 = Companion_D12_.Create_DC28_(_2108_v0, _2155_cf11, false, (_this).F4(), _2155_cf11)
			if (_2157_v36).Dtor_cf46() {
				var _2158_v37 _dafny.Map
				_ = _2158_v37
				_2158_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2125_v14).F4(), _2109_v1)
				var _2159_v39 _dafny.Array
				_ = _2159_v39
				var _nwElement0_71 _dafny.Int = _dafny.IntOfInt64(154)
				_ = _nwElement0_71
				var _nw328 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(7))
				_ = _nw328
				(_nw328).ArraySet1(_nwElement0_71, 0)
				(_nw328).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_2109_v1).Cardinality()), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _2108_v0)).Update(p0, _2108_v0)).Cardinality()), 1)
				(_nw328).ArraySet1((_this).Fm7(globalState), 2)
				(_nw328).ArraySet1(Companion_Default___.SafeModuloInt((func() _dafny.Set {
					var _coll57 = _dafny.NewBuilder()
					_ = _coll57
					for _iter64 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(587), _dafny.IntOfInt64(174))); ; {
						_compr_57, _ok64 := _iter64()
						if !_ok64 {
							break
						}
						var _2160_v38 _dafny.Int
						_2160_v38 = interface{}(_compr_57).(_dafny.Int)
						if ((_dafny.IntOfInt64(587)).Cmp(_2160_v38) <= 0) && ((_2160_v38).Cmp(_dafny.IntOfInt64(174)) < 0) {
							_coll57.Add((_2160_v38).Minus((_this).F6()))
						}
					}
					return _coll57.ToSet()
				}()).Cardinality(), _2155_cf11), 3)
				(_nw328).ArraySet1((_this).F6(), 4)
				(_nw328).ArraySet1((_dafny.MultiSetOf((_2125_v14).F4(), (_this).F4(), p0)).Cardinality(), 5)
				(_nw328).ArraySet1((_this).F6(), 6)
				_2159_v39 = _nw328
				var _2161_v40 _dafny.Sequence
				_ = _2161_v40
				_2161_v40 = _dafny.SeqOf(p0)
				var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2159_v39), 0))
				_ = _index312
				(_2159_v39).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2161_v40, _2161_v40)).Cardinality()), (_index312).Int())
				var _2162_v41 _dafny.Array
				_ = _2162_v41
				var _nw329 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw329
				_2162_v41 = _nw329
				var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_2162_v41), 0))
				_ = _index313
				(_2162_v41).ArraySet1(_2159_v39, (_index313).Int())
				var _2163_v42 bool
				_ = _2163_v42
				_2163_v42 = false
				var _2164_v43 _dafny.CodePoint
				_ = _2164_v43
				_2164_v43 = _dafny.CodePoint('h')
				var _2165_v44 _dafny.Map
				_ = _2165_v44
				_2165_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2125_v14).F4(), _2108_v0)
				var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2159_v39), 0))
				_ = _index314
				var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_2162_v41), 0))
				_ = _index315
				var _rhs301 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _2109_v1)).Merge(((_this).F9()).Merge(_2158_v37))
				_ = _rhs301
				var _rhs302 _dafny.Int = _2108_v0
				_ = _rhs302
				var _rhs303 _dafny.Array = _2159_v39
				_ = _rhs303
				var _rhs304 bool = Companion_Default___.Fm0(_2164_v43, globalState)
				_ = _rhs304
				var _rhs305 bool = ((_dafny.Zero).Minus((_dafny.Zero).Minus(((_2165_v44).Merge(_2165_v44)).Cardinality()))).Cmp(_dafny.IntOfInt64(-513)) == 0
				_ = _rhs305
				var _lhs187 _dafny.Array = _2159_v39
				_ = _lhs187
				var _lhs188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2159_v39), 0))
				_ = _lhs188
				var _lhs189 _dafny.Array = _2162_v41
				_ = _lhs189
				var _lhs190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_2162_v41), 0))
				_ = _lhs190
				_2158_v37 = _rhs301
				(_lhs187).ArraySet1(_rhs302, (_lhs188).Int())
				(_lhs189).ArraySet1(_rhs303, (_lhs190).Int())
				_2163_v42 = _rhs304
				_2163_v42 = _rhs305
				var _2166_v45 _dafny.Set
				_ = _2166_v45
				_2166_v45 = _dafny.SetOf(!((_2125_v14).F4()), !((_2125_v14).F4()))
				var _2167_v46 D16
				_ = _2167_v46
				_2167_v46 = Companion_D16_.Create_DC33_(false, _2155_cf11, _dafny.MultiSetOf(false), (_this).F6(), _2166_v45)
				_2163_v42 = (_dafny.MultiSetOf((_2125_v14).F4())).IsDisjointFrom((_2167_v46).Dtor_cf54())
				var _2168_v47 _dafny.Array
				_ = _2168_v47
				var _nw330 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(23))
				_ = _nw330
				_2168_v47 = _nw330
				var _2169_v48 _dafny.MultiSet
				_ = _2169_v48
				_2169_v48 = _dafny.MultiSetOf((_2159_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2159_v39), 0))).Int()).(_dafny.Int), (_2165_v44).Cardinality(), _2155_cf11)
				var _2170_v49 _dafny.Array
				_ = _2170_v49
				var _nwElement0_72 _dafny.CodePoint = _2164_v43
				_ = _nwElement0_72
				var _nw331 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(12))
				_ = _nw331
				(_nw331).ArraySet1CodePoint(_nwElement0_72, 0)
				(_nw331).ArraySet1CodePoint(Companion_Default___.Fm1((_2159_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2159_v39), 0))).Int()).(_dafny.Int), (_2125_v14).F4(), globalState), 1)
				(_nw331).ArraySet1CodePoint(_dafny.CodePoint('c'), 2)
				(_nw331).ArraySet1CodePoint(_2164_v43, 3)
				(_nw331).ArraySet1CodePoint(_2164_v43, 4)
				(_nw331).ArraySet1CodePoint(_2164_v43, 5)
				(_nw331).ArraySet1CodePoint(_2164_v43, 6)
				(_nw331).ArraySet1CodePoint(_dafny.CodePoint('y'), 7)
				(_nw331).ArraySet1CodePoint(_2164_v43, 8)
				(_nw331).ArraySet1CodePoint(_2164_v43, 9)
				(_nw331).ArraySet1CodePoint(_dafny.CodePoint('l'), 10)
				(_nw331).ArraySet1CodePoint(_2164_v43, 11)
				_2170_v49 = _nw331
				var _2171_v50 _dafny.Map
				_ = _2171_v50
				_2171_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_2169_v48).Contains(_2155_cf11) {
						return (_2169_v48).Multiplicity(_2155_cf11)
					}
					return _dafny.IntOfInt64(879)
				})(), _2170_v49)
				var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_2168_v47), 0))
				_ = _index316
				(_2168_v47).ArraySet1(_2171_v50, (_index316).Int())
				var _2172_v51 D11
				_ = _2172_v51
				_2172_v51 = Companion_D11_.Create_DC25_(true, _2155_cf11)
				var _2173_v52 _dafny.Set
				_ = _2173_v52
				_2173_v52 = _dafny.SetOf(_2108_v0)
				var _2174_v53 _dafny.Sequence
				_ = _2174_v53
				_2174_v53 = _dafny.SeqOf((_2172_v51).Dtor_cf40(), _dafny.IntOfUint32((_dafny.SeqOf(_2173_v52)).Cardinality()))
				var _2175_v54 _dafny.Map
				_ = _2175_v54
				_2175_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2174_v53).Cardinality()), (_2125_v14).F4())
				var _2176_v55 _dafny.Set
				_ = _2176_v55
				_2176_v55 = _dafny.SetOf(_2175_v54)
				var _2177_v56 _dafny.MultiSet
				_ = _2177_v56
				_2177_v56 = _dafny.MultiSetOf((_2125_v14).F4(), !(_2176_v55).Equals(_2176_v55), _2163_v42)
				var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_2168_v47), 0))
				_ = _index317
				var _rhs306 bool = (p0) || ((_2125_v14).F4())
				_ = _rhs306
				var _rhs307 bool = _2163_v42
				_ = _rhs307
				var _rhs308 _dafny.Int = (_2177_v56).Cardinality()
				_ = _rhs308
				var _rhs309 _dafny.Map = (func() _dafny.Map {
					if p0 {
						return (_2171_v50).Merge(_2171_v50)
					}
					return (_2171_v50).Merge(_2171_v50)
				})()
				_ = _rhs309
				var _lhs191 _dafny.Array = _2168_v47
				_ = _lhs191
				var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_2168_v47), 0))
				_ = _lhs192
				_2163_v42 = _rhs306
				_2163_v42 = _rhs307
				_2108_v0 = _rhs308
				(_lhs191).ArraySet1(_rhs309, (_lhs192).Int())
				var _2178_v57 *C1
				_ = _2178_v57
				var _nw332 *C1 = New_C1_()
				_ = _nw332
				_nw332.Ctor__((_2159_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2159_v39), 0))).Int()).(_dafny.Int), (_2159_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2159_v39), 0))).Int()).(_dafny.Int), _this.F5(), (func() _dafny.Int {
					if (_2165_v44).Contains(p0) {
						return (_2165_v44).Get(p0).(_dafny.Int)
					}
					return (_this).F6()
				})(), (_2125_v14).F4())
				_2178_v57 = _nw332
				var _2179_v58 _dafny.Sequence
				_ = _2179_v58
				_2179_v58 = _dafny.SeqOf((func() *C1 {
					if (_this).F4() {
						return _2178_v57
					}
					return _2178_v57
				})(), _2178_v57, _2178_v57, _2178_v57, _2178_v57)
				var _2180_v59 _dafny.Sequence
				_ = _2180_v59
				_2180_v59 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_2179_v58, _dafny.Companion_Sequence_.Update(_2179_v58, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("y")).Cardinality()), _dafny.IntOfUint32((_2179_v58).Cardinality()))).Uint32(), _2178_v57)))
				_2179_v58 = (_2180_v59).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.IntOfUint32((_2180_v59).Cardinality()))).Uint32()).(_dafny.Sequence)
				var _2181_v60 D9
				_ = _2181_v60
				_2181_v60 = Companion_D9_.Create_DC20_((_2125_v14).F4(), _2164_v43, Companion_Default___.SafeDivisionInt((_this).F6(), (_this).F6()))
				var _2182_v61 _dafny.Map
				_ = _2182_v61
				_2182_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2164_v43, (_this).F10())
				var _2183_v62 D19
				_ = _2183_v62
				_2183_v62 = Companion_D19_.Create_DC37_(_2182_v61)
				var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2159_v39), 0))
				_ = _index318
				var _rhs310 _dafny.Int = ((func() _dafny.Set {
					var _coll58 = _dafny.NewBuilder()
					_ = _coll58
					for _iter65 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(375), _dafny.IntOfInt64(806))); ; {
						_compr_58, _ok65 := _iter65()
						if !_ok65 {
							break
						}
						var _2184_v63 _dafny.Int
						_2184_v63 = interface{}(_compr_58).(_dafny.Int)
						if ((_dafny.IntOfInt64(375)).Cmp(_2184_v63) <= 0) && ((_2184_v63).Cmp(_dafny.IntOfInt64(806)) < 0) {
							_coll58.Add((_2184_v63).Times((_2178_v57).F15()))
						}
					}
					return _coll58.ToSet()
				}()).Cardinality()).Times(_2108_v0)
				_ = _rhs310
				var _rhs311 D9 = _2181_v60
				_ = _rhs311
				var _rhs312 bool = (_2125_v14).F4()
				_ = _rhs312
				var _rhs313 D19 = _2183_v62
				_ = _rhs313
				var _lhs193 _dafny.Array = _2159_v39
				_ = _lhs193
				var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2159_v39), 0))
				_ = _lhs194
				(_lhs193).ArraySet1(_rhs310, (_lhs194).Int())
				_2181_v60 = _rhs311
				_2163_v42 = _rhs312
				_2183_v62 = _rhs313
			} else {
				_2155_cf11 = _2108_v0
				var _2185_v64 bool
				_ = _2185_v64
				_2185_v64 = false
				var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index319
				((_this).F10()).ArraySet1(!((_2125_v14).F4()), (_index319).Int())
				var _2186_v65 _dafny.Array
				_ = _2186_v65
				var _nw333 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
				_ = _nw333
				_2186_v65 = _nw333
				var _2187_v66 _dafny.Map
				_ = _2187_v66
				_2187_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2108_v0, _2186_v65)
				var _2188_v67 _dafny.Set
				_ = _2188_v67
				_2188_v67 = _dafny.SetOf(false)
				var _2189_v68 _dafny.Sequence
				_ = _2189_v68
				_2189_v68 = _dafny.SeqOf(_2188_v67, _2188_v67, _2188_v67, _2188_v67)
				var _2190_v69 _dafny.Map
				_ = _2190_v69
				_2190_v69 = _2127_v16
				var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index320
				var _rhs314 bool = (_2108_v0).Cmp((_2187_v66).Cardinality()) >= 0
				_ = _rhs314
				var _rhs315 bool = !(!(_dafny.MultiSetOf(_2155_cf11)).Equals((_dafny.MultiSetOf((_this).F6(), _dafny.IntOfInt64(762))).Update((_this).F6(), Companion_Default___.Abs(_dafny.IntOfUint32((_2189_v68).Cardinality())))))
				_ = _rhs315
				var _rhs316 bool = true
				_ = _rhs316
				var _rhs317 bool = !(!(!(_2127_v16).Equals((_2190_v69))) || ((_2125_v14).F4()))
				_ = _rhs317
				var _lhs195 _dafny.Array = (_this).F10()
				_ = _lhs195
				var _lhs196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs196
				_2185_v64 = _rhs314
				_2185_v64 = _rhs315
				(_lhs195).ArraySet1(_rhs316, (_lhs196).Int())
				_2185_v64 = _rhs317
				_2108_v0 = _2108_v0
				var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index321
				((_this).F10()).ArraySet1(!(p0) || (((_2125_v14).F4()) || (p0)), (_index321).Int())
				var _2191_v70 _dafny.Map
				_ = _2191_v70
				_2191_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))
				var _2192_v71 _dafny.Sequence
				_ = _2192_v71
				_2192_v71 = _dafny.SeqOf(_2155_cf11)
				var _2193_v72 _dafny.Map
				_ = _2193_v72
				_2193_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC16_(true, (func() bool {
					if (_2191_v70).Contains((_2125_v14).F4()) {
						return (_2191_v70).Get((_2125_v14).F4()).(bool)
					}
					return (_2125_v14).F4()
				})(), _2192_v71), !(_2185_v64))
				_2193_v72 = _2193_v72
			}
			var _2194_v73 _dafny.Sequence
			_ = _2194_v73
			_2194_v73 = _dafny.SeqOf(true)
			var _2195_v74 _dafny.Set
			_ = _2195_v74
			_2195_v74 = _dafny.SetOf(_2108_v0, _dafny.IntOfUint32((_2194_v73).Cardinality()), _2108_v0)
			if (func() bool {
				if !((_2125_v14).F4()) {
					return (_2195_v74).IsProperSubsetOf(_2195_v74)
				}
				return p0
			})() {
				var _2196_v75 bool
				_ = _2196_v75
				_2196_v75 = false
				var _2197_v76 _dafny.MultiSet
				_ = _2197_v76
				_2197_v76 = _dafny.MultiSetOf(_dafny.IntOfUint32(((Companion_D10_.Create_DC22_(Companion_Default___.Fm50(_2108_v0, _2108_v0, _2155_cf11, (_2125_v14).F4(), globalState), p0, (_this).F4(), (_this).F10(), (_this).F6())).Dtor_cf30()).Cardinality()), _dafny.IntOfInt64(-163), _2155_cf11)
				_2196_v75 = (((_2197_v76).Update(_2155_cf11, Companion_Default___.Abs((_this).F6()))).Cardinality()).Cmp(_2108_v0) == 0
				_2155_cf11 = (_this).F6()
				_2155_cf11 = (((_this).F6()).Plus((_this).F6())).Minus(_dafny.IntOfUint32((_2109_v1).Cardinality()))
				_2155_cf11 = (func() _dafny.Int {
					if (Companion_Default___.Fm66(_2108_v0, globalState)).Dtor_cf52() {
						return (_2110_v2).Fm7(globalState)
					}
					return _dafny.IntOfUint32((_2194_v73).Cardinality())
				})()
				var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index322
				((_this).F10()).ArraySet1(p0, (_index322).Int())
				var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index323
				((_this).F10()).ArraySet1((_2108_v0).Cmp(_2108_v0) == 0, (_index323).Int())
			} else {
				var _2198_v77 bool
				_ = _2198_v77
				_2198_v77 = true
				var _2199_v78 D11
				_ = _2199_v78
				_2199_v78 = Companion_D11_.Create_DC25_(_2198_v77, (_this).F6())
				_2198_v77 = (_2199_v78).Dtor_cf39()
				var _2200_v79 _dafny.Map
				_ = _2200_v79
				_2200_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_2155_cf11, (_this).F6()), (_2125_v14).F4())
				var _2201_v80 _dafny.Sequence
				_ = _2201_v80
				_2201_v80 = _dafny.SeqOf((_dafny.MultiSetOf((_this).F6())).Cardinality(), _2155_cf11)
				var _2202_v81 *C1
				_ = _2202_v81
				var _nw334 *C1 = New_C1_()
				_ = _nw334
				_nw334.Ctor__(((_2200_v79).Update(_2201_v80, (_2125_v14).F4())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(400))).Uint32(), func(coer109 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg109 _dafny.Int) interface{} {
						return coer109(arg109)
					}
				}((func(_2203_cf11 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2204_i4 _dafny.Int) _dafny.Int {
						return _2203_cf11
					}
				})(_2155_cf11)))).Cardinality()), _this.F5(), _2155_cf11, (_2125_v14).F4())
				_2202_v81 = _nw334
				var _2205_v82 _dafny.Map
				_ = _2205_v82
				_2205_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2198_v77, (_dafny.Zero).Minus((_this).F6()))
				(_2202_v81).F16 = (_2202_v81.F16).Plus(Companion_Default___.SafeModuloInt((_2205_v82).Cardinality(), (_2125_v14).Fm7(globalState)))
				var _2206_v83 _dafny.MultiSet
				_ = _2206_v83
				_2206_v83 = _dafny.MultiSetOf(_2202_v81.F16, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2202_v81.F16, _2198_v77)).Cardinality(), (_2202_v81).F15())
				var _2207_v84 _dafny.Sequence
				_ = _2207_v84
				_2207_v84 = _dafny.SeqOf(_2194_v73)
				var _2208_v85 _dafny.MultiSet
				_ = _2208_v85
				_2208_v85 = _dafny.MultiSetOf(((_dafny.Zero).Minus((_dafny.Zero).Minus(_2202_v81.F16))).Times(_2155_cf11), Companion_Default___.SafeDivisionInt((_this).F6(), _2108_v0), (func() _dafny.Int {
					if (_2206_v83).Contains(_dafny.IntOfInt64(220)) {
						return (_2206_v83).Multiplicity(_dafny.IntOfInt64(220))
					}
					return (func() _dafny.Int {
						if (_2205_v82).Contains((_this).F4()) {
							return (_2205_v82).Get((_this).F4()).(_dafny.Int)
						}
						return (Companion_D4_.Create_DC13_((_2127_v16).Cardinality(), _2198_v77, _2108_v0, (_2207_v84).Select((Companion_Default___.SafeIndex(_2202_v81.F16, _dafny.IntOfUint32((_2207_v84).Cardinality()))).Uint32()).(_dafny.Sequence))).Dtor_cf16()
					})()
				})(), (_this).F6())
				var _2209_v86 _dafny.Map
				_ = _2209_v86
				_2209_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2202_v81.F16, (_this).F4())
				_2208_v85 = ((_dafny.MultiSetFromSeq(_2201_v80)).Update(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tfiwvdm")).Cardinality()), Companion_Default___.Abs((_2209_v86).Cardinality()))).Update(_2108_v0, Companion_Default___.Abs((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F6()))))
				var _2210_v87 _dafny.Array
				_ = _2210_v87
				var _nw335 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
				_ = _nw335
				_2210_v87 = _nw335
				var _2211_v88 D10
				_ = _2211_v88
				_2211_v88 = Companion_D10_.Create_DC21_(_2205_v82)
				var _2212_v89 _dafny.Map
				_ = _2212_v89
				_2212_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2211_v88, _2108_v0)
				var _2213_v90 _dafny.Sequence
				_ = _2213_v90
				_2213_v90 = _dafny.SeqOf(_2212_v89)
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen((_2210_v87), 0))
				_ = _index324
				(_2210_v87).ArraySet1(_2213_v90, (_index324).Int())
				var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index325
				((_this).F10()).ArraySet1((_2125_v14).F4(), (_index325).Int())
				var _2214_v91 _dafny.Set
				_ = _2214_v91
				_2214_v91 = _dafny.SetOf(_2109_v1, _dafny.UnicodeSeqOfUtf8Bytes("ugmeom"), _2109_v1)
				var _2215_v92 _dafny.Map
				_ = _2215_v92
				_2215_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_2109_v1, _dafny.UnicodeSeqOfUtf8Bytes("q")), ((_dafny.MultiSetOf((_this).F6(), _2202_v81.F16, _2155_cf11, _2155_cf11, (_dafny.Zero).Minus((_2214_v91).Cardinality()))).Cardinality()).Times(_dafny.IntOfUint32((_2201_v80).Cardinality())))
				var _2216_v93 _dafny.CodePoint
				_ = _2216_v93
				_2216_v93 = _dafny.CodePoint('w')
				var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen((_2210_v87), 0))
				_ = _index326
				var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index327
				var _rhs318 bool = p0
				_ = _rhs318
				var _rhs319 _dafny.Sequence = _2213_v90
				_ = _rhs319
				var _rhs320 bool = (_2125_v14).F4()
				_ = _rhs320
				var _rhs321 _dafny.CodePoint = _2216_v93
				_ = _rhs321
				var _rhs322 _dafny.Map = (_2215_v92).Update(_2109_v1, _2202_v81.F16)
				_ = _rhs322
				var _lhs197 _dafny.Array = _2210_v87
				_ = _lhs197
				var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen((_2210_v87), 0))
				_ = _lhs198
				var _lhs199 _dafny.Array = (_this).F10()
				_ = _lhs199
				var _lhs200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs200
				_2198_v77 = _rhs318
				(_lhs197).ArraySet1(_rhs319, (_lhs198).Int())
				(_lhs199).ArraySet1(_rhs320, (_lhs200).Int())
				r0 = _rhs321
				_2215_v92 = _rhs322
			}
			var _arr11 _dafny.Array = _this.F5()
			_ = _arr11
			var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_this.F5()), 0))
			_ = _index328
			(_arr11).ArraySet1((_this).F10(), (_index328).Int())
			var _2217_v94 _dafny.Sequence
			_ = _2217_v94
			_2217_v94 = _dafny.SeqOf((_this).F10(), (_this).F10(), (func() _dafny.Array {
				if (_2125_v14).F4() {
					return (_this).F10()
				}
				return (_this).F10()
			})(), (_this).F10(), (func() _dafny.Array {
				if (_2156_v35).Contains(_2108_v0) {
					return (_2156_v35).Get(_2108_v0).(_dafny.Array)
				}
				return (_this).F10()
			})())
			var _arr12 _dafny.Array = _this.F5()
			_ = _arr12
			var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_this.F5()), 0))
			_ = _index329
			(_arr12).ArraySet1((_2217_v94).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_2108_v0, _dafny.IntOfInt64(122)), _dafny.IntOfUint32((_2217_v94).Cardinality()))).Uint32()).(_dafny.Array), (_index329).Int())
		} else if _source36.Is_DC8() {
			var _2218___mcc_h4 _dafny.Sequence = _source36.Get_().(D3_DC8).Cf7
			_ = _2218___mcc_h4
			var _2219_cf7 _dafny.Sequence = _2218___mcc_h4
			_ = _2219_cf7
			var _2220_v95 _dafny.Map
			_ = _2220_v95
			_2220_v95 = Companion_Default___.Fm29(globalState)
			var _2221_v96 _dafny.Sequence
			_ = _2221_v96
			_2221_v96 = _dafny.SeqOf(_2219_cf7)
			var _2222_v97 D20
			_ = _2222_v97
			_2222_v97 = Companion_D20_.Create_DC40_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), Companion_D12_.Create_DC27_(_2221_v96)))
			var _2223_v98 _dafny.Map
			_ = _2223_v98
			_2223_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2219_cf7).Cardinality()), _2222_v97)
			var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index330
			((_this).F10()).ArraySet1((_this).F4(), (_index330).Int())
			var _2224_v99 _dafny.MultiSet
			_ = _2224_v99
			_2224_v99 = _dafny.MultiSetOf(_2219_cf7, _dafny.SeqOf(p0, (_this).F4(), p0))
			var _2225_v100 D6
			_ = _2225_v100
			_2225_v100 = Companion_D6_.Create_DC15_((_2224_v99).Update(_dafny.Companion_Sequence_.Update(_2219_cf7, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-524), _dafny.IntOfUint32((_2219_cf7).Cardinality()))).Uint32(), (_2125_v14).F4()), Companion_Default___.Abs(_2108_v0)))
			var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index331
			var _rhs323 _dafny.Int = (_this).F6()
			_ = _rhs323
			var _rhs324 _dafny.Map = _2220_v95
			_ = _rhs324
			var _rhs325 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _2222_v97)
			_ = _rhs325
			var _rhs326 bool = p0
			_ = _rhs326
			var _rhs327 D6 = _2225_v100
			_ = _rhs327
			var _lhs201 _dafny.Array = (_this).F10()
			_ = _lhs201
			var _lhs202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _lhs202
			_2108_v0 = _rhs323
			_2220_v95 = _rhs324
			_2223_v98 = _rhs325
			(_lhs201).ArraySet1(_rhs326, (_lhs202).Int())
			_2225_v100 = _rhs327
			var _2226_v101 _dafny.Array
			_ = _2226_v101
			var _len0_54 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_54
			var _nw336 _dafny.Array
			_ = _nw336
			if _len0_54.Cmp(_dafny.Zero) == 0 {
				_nw336 = _dafny.NewArray(_len0_54)
			} else {
				var _init54 func(_dafny.Int) _dafny.Sequence = (func(_2227_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_2228_i5 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2227_v1, _2227_v1), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2227_v1, _2227_v1)).Cardinality()))).Uint32(), _dafny.CodePoint('s'))
					}
				})(_2109_v1)
				_ = _init54
				var _element0_54 = _init54(_dafny.Zero)
				_ = _element0_54
				_nw336 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
				(_nw336).ArraySet1(_element0_54, 0)
				var _nativeLen0_54 = (_len0_54).Int()
				_ = _nativeLen0_54
				for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
					(_nw336).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
				}
			}
			_2226_v101 = _nw336
			var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_2226_v101), 0))
			_ = _index332
			(_2226_v101).ArraySet1(_2109_v1, (_index332).Int())
			var _2229_v102 _dafny.Sequence
			_ = _2229_v102
			_2229_v102 = _dafny.SeqOf((_this).F6())
			var _2230_v103 _dafny.CodePoint
			_ = _2230_v103
			_2230_v103 = _dafny.CodePoint('o')
			var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_2226_v101), 0))
			_ = _index333
			(_2226_v101).ArraySet1(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm37((_this).F4(), (func() bool {
				if (_2125_v14).F4() {
					return (_2125_v14).F4()
				}
				return (_2110_v2).Fm8(_2108_v0, true, _2109_v1, _2229_v102, globalState)
			})(), (_this).F6(), (_2125_v14).F4(), globalState), (Companion_Default___.SafeIndex(_2108_v0, _dafny.IntOfUint32((Companion_Default___.Fm37((_this).F4(), (func() bool {
				if (_2125_v14).F4() {
					return (_2125_v14).F4()
				}
				return (_2110_v2).Fm8(_2108_v0, true, _2109_v1, _2229_v102, globalState)
			})(), (_this).F6(), (_2125_v14).F4(), globalState)).Cardinality()))).Uint32(), _2230_v103), (_index333).Int())
			if p0 {
				var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index334
				((_this).F10()).ArraySet1((!((_2219_cf7).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.IntOfUint32((_2219_cf7).Cardinality()))).Uint32()).(bool)) || ((_2125_v14).F4())) && ((_this).F4()), (_index334).Int())
				var _2231_v104 _dafny.Array
				_ = _2231_v104
				var _len0_55 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_55
				var _nw337 _dafny.Array
				_ = _nw337
				if _len0_55.Cmp(_dafny.Zero) == 0 {
					_nw337 = _dafny.NewArray(_len0_55)
				} else {
					var _init55 func(_dafny.Int) bool = func(_2232_i6 _dafny.Int) bool {
						return (Companion_D1_.Create_DC3_(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))).Dtor_cf2()
					}
					_ = _init55
					var _element0_55 = _init55(_dafny.Zero)
					_ = _element0_55
					_nw337 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
					(_nw337).ArraySet1(_element0_55, 0)
					var _nativeLen0_55 = (_len0_55).Int()
					_ = _nativeLen0_55
					for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
						(_nw337).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
					}
				}
				_2231_v104 = _nw337
				_2231_v104 = (func() _dafny.Array {
					if _dafny.Companion_Sequence_.IsPrefixOf(_2229_v102, _2229_v102) {
						return _2231_v104
					}
					return _2231_v104
				})()
				_2108_v0 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt((_this).F6(), _dafny.IntOfUint32((_2219_cf7).Cardinality())), _dafny.IntOfInt64(846))
				var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index335
				((_this).F10()).ArraySet1(!(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool)), (_index335).Int())
				_2108_v0 = _2108_v0
			} else {
				var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index336
				var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_2226_v101), 0))
				_ = _index337
				var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index338
				var _rhs328 bool = !(!((_2125_v14).F4()) || (true))
				_ = _rhs328
				var _rhs329 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_2226_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_2226_v101), 0))).Int()).(_dafny.Sequence), _2109_v1)
				_ = _rhs329
				var _rhs330 bool = (func() bool {
					if (_2125_v14).F4() {
						return (_2108_v0).Cmp(_2108_v0) <= 0
					}
					return ((_this).F6()).Cmp(_2108_v0) >= 0
				})()
				_ = _rhs330
				var _rhs331 _dafny.Int = _2108_v0
				_ = _rhs331
				var _lhs203 _dafny.Array = (_this).F10()
				_ = _lhs203
				var _lhs204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs204
				var _lhs205 _dafny.Array = _2226_v101
				_ = _lhs205
				var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_2226_v101), 0))
				_ = _lhs206
				var _lhs207 _dafny.Array = (_this).F10()
				_ = _lhs207
				var _lhs208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs208
				(_lhs203).ArraySet1(_rhs328, (_lhs204).Int())
				(_lhs205).ArraySet1(_rhs329, (_lhs206).Int())
				(_lhs207).ArraySet1(_rhs330, (_lhs208).Int())
				_2108_v0 = _rhs331
				var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index339
				((_this).F10()).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_2125_v14).F4())).Contains((Companion_D3_.Create_DC9_(Companion_Default___.Fm0(_2230_v103, globalState), (_this).F6(), ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))).Dtor_cf8()), (_index339).Int())
				var _2233_v105 D12
				_ = _2233_v105
				_2233_v105 = Companion_D12_.Create_DC27_(_dafny.Companion_Sequence_.Update(_2221_v96, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2221_v96).Cardinality()))).Uint32(), _2219_cf7))
				var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index340
				var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index341
				var _rhs332 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm39((_2226_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_2226_v101), 0))).Int()).(_dafny.Sequence), _2233_v105, (_2125_v14).F4(), globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2108_v0), _2229_v102))
				_ = _rhs332
				var _rhs333 _dafny.Int = (_2229_v102).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2229_v102).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs333
				var _rhs334 bool = (_2125_v14).F4()
				_ = _rhs334
				var _rhs335 bool = !(false)
				_ = _rhs335
				var _lhs209 _dafny.Array = (_this).F10()
				_ = _lhs209
				var _lhs210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs210
				var _lhs211 _dafny.Array = (_this).F10()
				_ = _lhs211
				var _lhs212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs212
				_2229_v102 = _rhs332
				_2108_v0 = _rhs333
				(_lhs209).ArraySet1(_rhs334, (_lhs210).Int())
				(_lhs211).ArraySet1(_rhs335, (_lhs212).Int())
				var _2234_v106 T1
				_ = _2234_v106
				var _nw338 *C6 = New_C6_()
				_ = _nw338
				_nw338.Ctor__(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), (_this).F4(), _this.F5(), (_this).F6())
				_2234_v106 = _nw338
				var _2235_v107 _dafny.Map
				_ = _2235_v107
				_2235_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2234_v106, _2109_v1)
				var _2236_v108 _dafny.Map
				_ = _2236_v108
				_2236_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2125_v14).F4(), (_this).F6())
				var _2237_v109 _dafny.MultiSet
				_ = _2237_v109
				_2237_v109 = _dafny.MultiSetOf((_2234_v106).F6(), (func() _dafny.Int {
					if (_2236_v108).Contains(!((_2110_v2).Fm8(_dafny.IntOfUint32((_2219_cf7).Cardinality()), (_2234_v106).F4(), _2109_v1, _2229_v102, globalState))) {
						return (_2236_v108).Get(!((_2110_v2).Fm8(_dafny.IntOfUint32((_2219_cf7).Cardinality()), (_2234_v106).F4(), _2109_v1, _2229_v102, globalState))).(_dafny.Int)
					}
					return (_this).F6()
				})())
				var _2238_v110 _dafny.Map
				_ = _2238_v110
				_2238_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2234_v106).F6(), _2109_v1)
				var _2239_v111 _dafny.Map
				_ = _2239_v111
				_2239_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), (_2125_v14).F4())
				var _2240_v112 _dafny.MultiSet
				_ = _2240_v112
				_2240_v112 = _dafny.MultiSetOf(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))
				var _2241_v113 _dafny.Array
				_ = _2241_v113
				var _nwElement0_73 _dafny.Int = (_2110_v2).Fm7(globalState)
				_ = _nwElement0_73
				var _nw339 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(26))
				_ = _nw339
				(_nw339).ArraySet1(_nwElement0_73, 0)
				(_nw339).ArraySet1(_2108_v0, 1)
				(_nw339).ArraySet1((_this).Fm9(true, false, (_this).F6(), globalState), 2)
				(_nw339).ArraySet1(((_2127_v16).Merge(_2127_v16)).Cardinality(), 3)
				(_nw339).ArraySet1(_2108_v0, 4)
				(_nw339).ArraySet1(_2108_v0, 5)
				(_nw339).ArraySet1(((_2235_v107).Cardinality()).Times(_dafny.IntOfUint32((_2219_cf7).Cardinality())), 6)
				(_nw339).ArraySet1(_2108_v0, 7)
				(_nw339).ArraySet1(((_this).F6()).Plus(_2108_v0), 8)
				(_nw339).ArraySet1((_2234_v106).F6(), 9)
				(_nw339).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2229_v102, _2229_v102)).Cardinality()), 10)
				(_nw339).ArraySet1(((_2237_v109).Difference(_2237_v109)).Cardinality(), 11)
				(_nw339).ArraySet1((_2234_v106).F6(), 12)
				(_nw339).ArraySet1((_this).F6(), 13)
				(_nw339).ArraySet1((func() _dafny.Int {
					if (_2234_v106).F4() {
						return (_2238_v110).Cardinality()
					}
					return (_this).F6()
				})(), 14)
				(_nw339).ArraySet1(_2108_v0, 15)
				(_nw339).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), 16)
				(_nw339).ArraySet1((_this).F6(), 17)
				(_nw339).ArraySet1((_2234_v106).F6(), 18)
				(_nw339).ArraySet1((_this).F6(), 19)
				(_nw339).ArraySet1(((_dafny.SetOf(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))).Cardinality()).Times((_this).F6()), 20)
				(_nw339).ArraySet1(_2108_v0, 21)
				(_nw339).ArraySet1(_dafny.IntOfUint32((_2109_v1).Cardinality()), 22)
				(_nw339).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_2226_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_2226_v101), 0))).Int()).(_dafny.Sequence), (_2226_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_2226_v101), 0))).Int()).(_dafny.Sequence))).Cardinality()), 23)
				(_nw339).ArraySet1((_2239_v111).Cardinality(), 24)
				(_nw339).ArraySet1((_2240_v112).Cardinality(), 25)
				_2241_v113 = _nw339
				var _2242_v114 _dafny.Map
				_ = _2242_v114
				_2242_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2127_v16, (_2125_v14).F4())
				var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_2241_v113), 0))
				_ = _index342
				(_2241_v113).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_2125_v14).F4() {
						return _dafny.SeqOf((_2242_v114).Cardinality())
					}
					return _2229_v102
				})()).Cardinality()), (_index342).Int())
				var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_2241_v113), 0))
				_ = _index343
				(_2241_v113).ArraySet1(Companion_Default___.SafeDivisionInt(_2108_v0, (_this).F6()), (_index343).Int())
				var _2243_v115 _dafny.Array
				_ = _2243_v115
				var _len0_56 _dafny.Int = _dafny.One
				_ = _len0_56
				var _nw340 _dafny.Array
				_ = _nw340
				if _len0_56.Cmp(_dafny.Zero) == 0 {
					_nw340 = _dafny.NewArray(_len0_56)
				} else {
					var _init56 func(_dafny.Int) bool = (func(_2244_v106 T1) func(_dafny.Int) bool {
						return func(_2245_i7 _dafny.Int) bool {
							return (_2244_v106).F4()
						}
					})(_2234_v106)
					_ = _init56
					var _element0_56 = _init56(_dafny.Zero)
					_ = _element0_56
					_nw340 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
					(_nw340).ArraySet1(_element0_56, 0)
					var _nativeLen0_56 = (_len0_56).Int()
					_ = _nativeLen0_56
					for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
						(_nw340).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
					}
				}
				_2243_v115 = _nw340
				var _2246_v116 _dafny.Map
				_ = _2246_v116
				_2246_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2230_v103, _2243_v115)
				var _2247_v117 D19
				_ = _2247_v117
				_2247_v117 = Companion_D19_.Create_DC37_(_2246_v116)
				var _2248_v118 D19
				_ = _2248_v118
				_2248_v118 = Companion_D19_.Create_DC39_(_2247_v117)
				var _2249_v119 _dafny.Map
				_ = _2249_v119
				_2249_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2248_v118, (_dafny.Zero).Minus((_this).F6()))
				var _2250_v120 _dafny.Map
				_ = _2250_v120
				_2250_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2230_v103, (_2249_v119).Cardinality())
				_2250_v120 = (_2250_v120).Update(_2230_v103, _dafny.IntOfInt64(235))
			}
			var _2251_v121 T1
			_ = _2251_v121
			var _nw341 *C1 = New_C1_()
			_ = _nw341
			_nw341.Ctor__(_dafny.IntOfInt64(669), (_this).F6(), _this.F5(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kv")).Cardinality()), (_2125_v14).F4())
			_2251_v121 = _nw341
			var _2252_v122 _dafny.Sequence
			_ = _2252_v122
			_2252_v122 = _dafny.SeqOf(Companion_Default___.Fm57(_2108_v0, _2108_v0, globalState), _2229_v102, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(875))).Uint32(), func(coer110 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg110 _dafny.Int) interface{} {
					return coer110(arg110)
				}
			}(func(_2253_i8 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(246)
			})), _2229_v102)
			var _nw342 *C5 = New_C5_()
			_ = _nw342
			_nw342.Ctor__((_2251_v121).F4(), Companion_Default___.Fm3(Companion_Default___.Fm3((_this).F6(), true, globalState), ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), globalState), _2251_v121.F5(), _dafny.IntOfUint32((_2252_v122).Cardinality()), p0)
			_2251_v121 = _nw342
		} else {
			var _2254___mcc_h5 D3 = _source36.Get_().(D3_DC11).Cf12
			_ = _2254___mcc_h5
			var _2255_cf12 D3 = _2254___mcc_h5
			_ = _2255_cf12
			var _2256_v123 _dafny.Map
			_ = _2256_v123
			_2256_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2109_v1, _dafny.CodePoint('a'))
			var _2257_v124 _dafny.MultiSet
			_ = _2257_v124
			_2257_v124 = _dafny.MultiSetOf((_2125_v14).F4(), true, (_2125_v14).F4())
			var _2258_v125 _dafny.Map
			_ = _2258_v125
			_2258_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC6_(), (func() _dafny.Int {
				if (_2127_v16).Contains((_2256_v123).Cardinality()) {
					return (_2127_v16).Get((_2256_v123).Cardinality()).(_dafny.Int)
				}
				return (func() _dafny.Int {
					if (_2257_v124).Contains((_2125_v14).F4()) {
						return (_2257_v124).Multiplicity((_2125_v14).F4())
					}
					return (_2110_v2).Fm9(true, false, _dafny.IntOfInt64(632), globalState)
				})()
			})())
			var _2259_v126 D2
			_ = _2259_v126
			_2259_v126 = Companion_D2_.Create_DC6_()
			_2258_v125 = (_2258_v125).Update(_2259_v126, Companion_Default___.SafeModuloInt((_this).F6(), (_this).F6()))
			var _2260_v127 _dafny.Array
			_ = _2260_v127
			var _len0_57 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_57
			var _nw343 _dafny.Array
			_ = _nw343
			if _len0_57.Cmp(_dafny.Zero) == 0 {
				_nw343 = _dafny.NewArray(_len0_57)
			} else {
				var _init57 func(_dafny.Int) _dafny.Int = (func(_2261_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2262_i9 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_2262_i9, _2261_v0)
					}
				})(_2108_v0)
				_ = _init57
				var _element0_57 = _init57(_dafny.Zero)
				_ = _element0_57
				_nw343 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
				(_nw343).ArraySet1(_element0_57, 0)
				var _nativeLen0_57 = (_len0_57).Int()
				_ = _nativeLen0_57
				for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
					(_nw343).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
				}
			}
			_2260_v127 = _nw343
			var _2263_v128 D16
			_ = _2263_v128
			_2263_v128 = Companion_D16_.Create_DC32_(_2260_v127)
			var _source37 D16 = (func() D16 {
				if p0 {
					return _2263_v128
				}
				return _2263_v128
			})()
			_ = _source37
			if _source37.Is_DC33() {
				var _2264___mcc_h6 bool = _source37.Get_().(D16_DC33).Cf52
				_ = _2264___mcc_h6
				var _2265___mcc_h7 _dafny.Int = _source37.Get_().(D16_DC33).Cf53
				_ = _2265___mcc_h7
				var _2266___mcc_h8 _dafny.MultiSet = _source37.Get_().(D16_DC33).Cf54
				_ = _2266___mcc_h8
				var _2267___mcc_h9 _dafny.Int = _source37.Get_().(D16_DC33).Cf55
				_ = _2267___mcc_h9
				var _2268___mcc_h10 _dafny.Set = _source37.Get_().(D16_DC33).Cf56
				_ = _2268___mcc_h10
				var _2269_cf56 _dafny.Set = _2268___mcc_h10
				_ = _2269_cf56
				var _2270_cf55 _dafny.Int = _2267___mcc_h9
				_ = _2270_cf55
				var _2271_cf54 _dafny.MultiSet = _2266___mcc_h8
				_ = _2271_cf54
				var _2272_cf53 _dafny.Int = _2265___mcc_h7
				_ = _2272_cf53
				var _2273_cf52 bool = _2264___mcc_h6
				_ = _2273_cf52
				var _2274_v129 _dafny.Map
				_ = _2274_v129
				_2274_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2125_v14).F4(), _2270_cf55)
				_2108_v0 = Companion_Default___.SafeModuloInt((_this).F6(), ((_2274_v129).Cardinality()).Plus(_dafny.IntOfUint32((_2109_v1).Cardinality())))
				var _2275_v130 _dafny.CodePoint
				_ = _2275_v130
				_2275_v130 = _dafny.CodePoint('f')
				r0 = _2275_v130
				var _2276_v131 D4
				_ = _2276_v131
				_2276_v131 = Companion_D4_.Create_DC12_(_dafny.CodePoint('p'))
				_2276_v131 = _2276_v131
				var _2277_v132 _dafny.Array
				_ = _2277_v132
				var _len0_58 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_58
				var _nw344 _dafny.Array
				_ = _nw344
				if _len0_58.Cmp(_dafny.Zero) == 0 {
					_nw344 = _dafny.NewArray(_len0_58)
				} else {
					var _init58 func(_dafny.Int) _dafny.Set = (func(_2278_v0 _dafny.Int, _2279_cf55 _dafny.Int, _2280_v130 _dafny.CodePoint) func(_dafny.Int) _dafny.Set {
						return func(_2281_i10 _dafny.Int) _dafny.Set {
							return (_dafny.SetOf(_dafny.SeqOf((_this).F6(), _2278_v0, (_this).F6()), _dafny.SeqOf(_2279_cf55, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dv")).Cardinality()), (_dafny.SetOf(_2280_v130, _2280_v130)).Cardinality()), _dafny.SeqOf(_2278_v0))).Difference(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(741))))
						}
					})(_2108_v0, _2270_cf55, _2275_v130)
					_ = _init58
					var _element0_58 = _init58(_dafny.Zero)
					_ = _element0_58
					_nw344 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
					(_nw344).ArraySet1(_element0_58, 0)
					var _nativeLen0_58 = (_len0_58).Int()
					_ = _nativeLen0_58
					for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
						(_nw344).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
					}
				}
				_2277_v132 = _nw344
				var _2282_v133 _dafny.Sequence
				_ = _2282_v133
				_2282_v133 = _dafny.SeqOf((_dafny.Zero).Minus(_2270_cf55))
				var _2283_v134 _dafny.MultiSet
				_ = _2283_v134
				_2283_v134 = _dafny.MultiSetOf(_2125_v14)
				var _2284_v135 _dafny.Set
				_ = _2284_v135
				_2284_v135 = _dafny.SetOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tm")).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(283))).Uint32(), func(coer111 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg111 _dafny.Int) interface{} {
						return coer111(arg111)
					}
				}((func(_2285_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2286_i11 _dafny.Int) _dafny.Int {
						return _2285_v0
					}
				})(_2108_v0))), _2282_v133, _2282_v133, _dafny.SeqOf(_dafny.IntOfInt64(982), (_2283_v134).Cardinality()))
				var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_2277_v132), 0))
				_ = _index344
				(_2277_v132).ArraySet1(_2284_v135, (_index344).Int())
				var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index345
				((_this).F10()).ArraySet1(_2273_cf52, (_index345).Int())
				var _2287_v136 D6
				_ = _2287_v136
				_2287_v136 = Companion_D6_.Create_DC16_((_2125_v14).F4(), (_this).F4(), _2282_v133)
				var _2288_v137 _dafny.Set
				_ = _2288_v137
				_2288_v137 = _dafny.SetOf(_2287_v136)
				var _2289_v138 _dafny.Map
				_ = _2289_v138
				_2289_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm42(_2275_v130, _2288_v137, globalState), _2275_v130)
				var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_2260_v127), 0))
				_ = _index346
				(_2260_v127).ArraySet1(_2108_v0, (_index346).Int())
				var _2290_v139 _dafny.Sequence
				_ = _2290_v139
				_2290_v139 = _dafny.SeqOf(_2273_cf52, _2273_cf52)
				var _2291_v140 _dafny.Sequence
				_ = _2291_v140
				_2291_v140 = _dafny.SeqOf(_2290_v139, _2290_v139, _2290_v139)
				var _2292_v141 D12
				_ = _2292_v141
				_2292_v141 = Companion_D12_.Create_DC27_(_2291_v140)
				var _2293_v142 _dafny.Map
				_ = _2293_v142
				_2293_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2125_v14).F4(), (_this).F4())
				var _2294_v143 _dafny.Map
				_ = _2294_v143
				_2294_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (func() bool {
					if (_2293_v142).Contains(p0) {
						return (_2293_v142).Get(p0).(bool)
					}
					return (_2125_v14).F4()
				})())
				var _2295_v144 D9
				_ = _2295_v144
				_2295_v144 = Companion_D9_.Create_DC19_(_dafny.UnicodeSeqOfUtf8Bytes("mskr"))
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_2277_v132), 0))
				_ = _index347
				var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index348
				var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_2260_v127), 0))
				_ = _index349
				var _rhs336 _dafny.Set = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_2282_v133, Companion_Default___.Fm39(_2109_v1, _2292_v141, (func() bool {
					if (_2294_v143).Contains(_dafny.IntOfUint32((_2282_v133).Cardinality())) {
						return (_2294_v143).Get(_dafny.IntOfUint32((_2282_v133).Cardinality())).(bool)
					}
					return p0
				})(), globalState)))
				_ = _rhs336
				var _rhs337 bool = (_this).F4()
				_ = _rhs337
				var _rhs338 _dafny.Map = (_2289_v138).Merge(_2289_v138)
				_ = _rhs338
				var _rhs339 _dafny.Int = _dafny.IntOfUint32(((_2295_v144).Dtor_cf25()).Cardinality())
				_ = _rhs339
				var _lhs213 _dafny.Array = _2277_v132
				_ = _lhs213
				var _lhs214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_2277_v132), 0))
				_ = _lhs214
				var _lhs215 _dafny.Array = (_this).F10()
				_ = _lhs215
				var _lhs216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs216
				var _lhs217 _dafny.Array = _2260_v127
				_ = _lhs217
				var _lhs218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_2260_v127), 0))
				_ = _lhs218
				(_lhs213).ArraySet1(_rhs336, (_lhs214).Int())
				(_lhs215).ArraySet1(_rhs337, (_lhs216).Int())
				_2289_v138 = _rhs338
				(_lhs217).ArraySet1(_rhs339, (_lhs218).Int())
			} else {
				var _2296___mcc_h11 _dafny.Array = _source37.Get_().(D16_DC32).Cf51
				_ = _2296___mcc_h11
				var _2297_cf51 _dafny.Array = _2296___mcc_h11
				_ = _2297_cf51
				var _2298_v145 D3
				_ = _2298_v145
				_2298_v145 = Companion_D3_.Create_DC10_((_this).F6())
				_2108_v0 = Companion_Default___.SafeModuloInt((_2298_v145).Dtor_cf11(), _dafny.IntOfInt64(765))
				var _2299_v146 _dafny.Sequence
				_ = _2299_v146
				_2299_v146 = _dafny.SeqOf(false, (_2110_v2).Fm8(_2108_v0, (_this).F4(), _2109_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(133))).Uint32(), func(coer112 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg112 _dafny.Int) interface{} {
						return coer112(arg112)
					}
				}(func(_2300_i12 _dafny.Int) _dafny.Int {
					return (_this).F6()
				})), globalState))
				var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index350
				((_this).F10()).ArraySet1((_2299_v146).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2299_v146).Cardinality()))).Uint32()).(bool), (_index350).Int())
				var _2301_v147 _dafny.Array
				_ = _2301_v147
				var _len0_59 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_59
				var _nw345 _dafny.Array
				_ = _nw345
				if _len0_59.Cmp(_dafny.Zero) == 0 {
					_nw345 = _dafny.NewArray(_len0_59)
				} else {
					var _init59 func(_dafny.Int) D0 = func(_2302_i13 _dafny.Int) D0 {
						return Companion_D0_.Create_DC1_()
					}
					_ = _init59
					var _element0_59 = _init59(_dafny.Zero)
					_ = _element0_59
					_nw345 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
					(_nw345).ArraySet1(_element0_59, 0)
					var _nativeLen0_59 = (_len0_59).Int()
					_ = _nativeLen0_59
					for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
						(_nw345).ArraySet1(_init59(_dafny.IntOf(_i0_59)), _i0_59)
					}
				}
				_2301_v147 = _nw345
				var _2303_v148 bool
				_ = _2303_v148
				_2303_v148 = true
				var _2304_v149 _dafny.CodePoint
				_ = _2304_v149
				_2304_v149 = _dafny.CodePoint('p')
				var _2305_v150 _dafny.Array
				_ = _2305_v150
				var _nwElement0_74 _dafny.Sequence = _2109_v1
				_ = _nwElement0_74
				var _nw346 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(10))
				_ = _nw346
				(_nw346).ArraySet1(_nwElement0_74, 0)
				(_nw346).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("xkjwa"), 1)
				(_nw346).ArraySet1(_2109_v1, 2)
				(_nw346).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tc"), _2109_v1), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2125_v14).F4(), _2304_v149)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tc"), _2109_v1)).Cardinality()))).Uint32(), _dafny.CodePoint('c')), 3)
				(_nw346).ArraySet1(_2109_v1, 4)
				(_nw346).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vtrm"), 5)
				(_nw346).ArraySet1(_2109_v1, 6)
				(_nw346).ArraySet1((func() _dafny.Sequence {
					if (_2125_v14).F4() {
						return _2109_v1
					}
					return _2109_v1
				})(), 7)
				(_nw346).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(471))).Uint32(), func(coer113 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg113 _dafny.Int) interface{} {
						return coer113(arg113)
					}
				}(func(_2306_i14 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				})), 8)
				(_nw346).ArraySet1(_2109_v1, 9)
				_2305_v150 = _nw346
				var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(38), _dafny.ArrayLen((_2305_v150), 0))
				_ = _index351
				(_2305_v150).ArraySet1((func() _dafny.Sequence {
					if (_this).F4() {
						return _2109_v1
					}
					return _2109_v1
				})(), (_index351).Int())
				var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index352
				var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(38), _dafny.ArrayLen((_2305_v150), 0))
				_ = _index353
				var _rhs340 bool = (_2125_v14).F4()
				_ = _rhs340
				var _rhs341 _dafny.Array = (Companion_D21_.Create_DC42_(_2301_v147)).Dtor_cf68()
				_ = _rhs341
				var _rhs342 bool = (_2125_v14).F4()
				_ = _rhs342
				var _rhs343 _dafny.Int = _2108_v0
				_ = _rhs343
				var _rhs344 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2109_v1, _2109_v1)
				_ = _rhs344
				var _lhs219 _dafny.Array = (_this).F10()
				_ = _lhs219
				var _lhs220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs220
				var _lhs221 _dafny.Array = _2305_v150
				_ = _lhs221
				var _lhs222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(38), _dafny.ArrayLen((_2305_v150), 0))
				_ = _lhs222
				(_lhs219).ArraySet1(_rhs340, (_lhs220).Int())
				_2301_v147 = _rhs341
				_2303_v148 = _rhs342
				_2108_v0 = _rhs343
				(_lhs221).ArraySet1(_rhs344, (_lhs222).Int())
				var _2307_v151 _dafny.Map
				_ = _2307_v151
				_2307_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))
				var _2308_v152 _dafny.Set
				_ = _2308_v152
				_2308_v152 = _dafny.SetOf((_this).F6())
				var _2309_v153 _dafny.Map
				_ = _2309_v153
				_2309_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2307_v151, (_2308_v152).Cardinality())
				var _rhs345 bool = Companion_Default___.Fm0(_dafny.CodePoint('o'), globalState)
				_ = _rhs345
				var _rhs346 bool = !((_2309_v153).Equals(_2309_v153))
				_ = _rhs346
				_2303_v148 = _rhs345
				_2303_v148 = _rhs346
				_2108_v0 = _2108_v0
			}
			var _2310_v154 _dafny.Map
			_ = _2310_v154
			_2310_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2109_v1, _dafny.IntOfUint32((_2109_v1).Cardinality()))
			var _2311_v155 _dafny.Map
			_ = _2311_v155
			_2311_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2310_v154, _2108_v0)
			_2311_v155 = (_2311_v155).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2109_v1, (_dafny.Zero).Minus((_this).F6())), (_this).F6())
			var _2312_v156 _dafny.Sequence
			_ = _2312_v156
			_2312_v156 = _dafny.SeqOf(_2108_v0)
			var _2313_v157 _dafny.Set
			_ = _2313_v157
			_2313_v157 = _dafny.SetOf(p0)
			_2108_v0 = (_2312_v156).Select((Companion_Default___.SafeIndex((_2313_v157).Cardinality(), _dafny.IntOfUint32((_2312_v156).Cardinality()))).Uint32()).(_dafny.Int)
		}
		var _2314_v158 bool
		_ = _2314_v158
		_2314_v158 = true
		_2314_v158 = (_this).F4()
		var _2315_v159 _dafny.CodePoint
		_ = _2315_v159
		_2315_v159 = _dafny.CodePoint('k')
		r0 = _2315_v159
		_2108_v0 = _2108_v0
		r0 = _2315_v159
		var _2316_v160 _dafny.Array
		_ = _2316_v160
		var _nw347 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(19))
		_ = _nw347
		_2316_v160 = _nw347
		r1 = _2316_v160
		return r0, r1
	}
}
func (_this *C7) M6(globalState *GlobalState) (D1, _dafny.Sequence) {
	{
		var r0 D1 = Companion_D1_.Default()
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var _2317_i0 _dafny.Int
		_ = _2317_i0
		_2317_i0 = _dafny.Zero
		{
			for (_this).F4() {
				{
					if (_2317_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_2317_i0 = (_2317_i0).Plus(_dafny.One)
					var _2318_v0 _dafny.Array
					_ = _2318_v0
					var _nw348 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
					_ = _nw348
					_2318_v0 = _nw348
					var _2319_v1 _dafny.Map
					_ = _2319_v1
					_2319_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_this).F4())
					var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_2318_v0), 0))
					_ = _index354
					(_2318_v0).ArraySet1((_2319_v1).Merge(_2319_v1), (_index354).Int())
					var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_2318_v0), 0))
					_ = _index355
					(_2318_v0).ArraySet1((_2319_v1).Merge(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_this).F4())).Update((_this).F10(), (_this).F4())).Merge(_2319_v1)), (_index355).Int())
					var _2320_v2 _dafny.Array
					_ = _2320_v2
					var _nw349 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
					_ = _nw349
					_2320_v2 = _nw349
					var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_2320_v2), 0))
					_ = _index356
					(_2320_v2).ArraySet1((_this).F6(), (_index356).Int())
					var _2321_v3 _dafny.CodePoint
					_ = _2321_v3
					_2321_v3 = _dafny.CodePoint('w')
					var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F10()), 0))
					_ = _index357
					((_this).F10()).ArraySet1(Companion_Default___.Fm0(_2321_v3, globalState), (_index357).Int())
					var _2322_v4 bool
					_ = _2322_v4
					_2322_v4 = true
					var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_2320_v2), 0))
					_ = _index358
					var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F10()), 0))
					_ = _index359
					var _rhs347 _dafny.Int = (_this).F6()
					_ = _rhs347
					var _rhs348 bool = (_this).F4()
					_ = _rhs348
					var _rhs349 bool = _2322_v4
					_ = _rhs349
					var _lhs223 _dafny.Array = _2320_v2
					_ = _lhs223
					var _lhs224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_2320_v2), 0))
					_ = _lhs224
					var _lhs225 _dafny.Array = (_this).F10()
					_ = _lhs225
					var _lhs226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F10()), 0))
					_ = _lhs226
					(_lhs223).ArraySet1(_rhs347, (_lhs224).Int())
					(_lhs225).ArraySet1(_rhs348, (_lhs226).Int())
					_2322_v4 = _rhs349
					var _2323_v5 _dafny.Array
					_ = _2323_v5
					var _len0_60 _dafny.Int = _dafny.IntOfInt64(13)
					_ = _len0_60
					var _nw350 _dafny.Array
					_ = _nw350
					if _len0_60.Cmp(_dafny.Zero) == 0 {
						_nw350 = _dafny.NewArray(_len0_60)
					} else {
						var _init60 func(_dafny.Int) _dafny.Map = (func(_2324_v2 _dafny.Array) func(_dafny.Int) _dafny.Map {
							return func(_2325_i1 _dafny.Int) _dafny.Map {
								return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_2324_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_2324_v2), 0))).Int()).(_dafny.Int))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())))
							}
						})(_2320_v2)
						_ = _init60
						var _element0_60 = _init60(_dafny.Zero)
						_ = _element0_60
						_nw350 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
						(_nw350).ArraySet1(_element0_60, 0)
						var _nativeLen0_60 = (_len0_60).Int()
						_ = _nativeLen0_60
						for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
							(_nw350).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
						}
					}
					_2323_v5 = _nw350
					var _2326_v6 _dafny.Sequence
					_ = _2326_v6
					_2326_v6 = _dafny.SeqOf((Companion_D1_.Create_DC4_((_2320_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_2320_v2), 0))).Int()).(_dafny.Int), (_2320_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_2320_v2), 0))).Int()).(_dafny.Int))).Dtor_cf3())
					var _2327_v7 _dafny.Map
					_ = _2327_v7
					_2327_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2326_v6).Cardinality()), (_this).F6())
					var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_2323_v5), 0))
					_ = _index360
					(_2323_v5).ArraySet1(_2327_v7, (_index360).Int())
					var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_2323_v5), 0))
					_ = _index361
					(_2323_v5).ArraySet1(_2327_v7, (_index361).Int())
					var _2328_v8 *C6
					_ = _2328_v8
					var _nw351 *C6 = New_C6_()
					_ = _nw351
					_nw351.Ctor__(Companion_Default___.Fm0(_2321_v3, globalState), false, _this.F5(), Companion_Default___.SafeModuloInt((_this).F6(), (_this).F6()))
					_2328_v8 = _nw351
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _hi13 _dafny.Int = (_this).F6()
		_ = _hi13
		for _2329_i2 := (_this).F6(); _2329_i2.Cmp(_hi13) < 0; _2329_i2 = _2329_i2.Plus(_dafny.One) {
			var _index362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index362
			((_this).F10()).ArraySet1((false) || (true), (_index362).Int())
			var _index363 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index363
			((_this).F10()).ArraySet1(((_this).F6()).Cmp(_2329_i2) >= 0, (_index363).Int())
			var _index364 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index364
			((_this).F10()).ArraySet1(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), (_index364).Int())
			var _2330_v9 _dafny.Array
			_ = _2330_v9
			var _len0_61 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_61
			var _nw352 _dafny.Array
			_ = _nw352
			if _len0_61.Cmp(_dafny.Zero) == 0 {
				_nw352 = _dafny.NewArray(_len0_61)
			} else {
				var _init61 func(_dafny.Int) _dafny.Map = func(_2331_i3 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))
				}
				_ = _init61
				var _element0_61 = _init61(_dafny.Zero)
				_ = _element0_61
				_nw352 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
				(_nw352).ArraySet1(_element0_61, 0)
				var _nativeLen0_61 = (_len0_61).Int()
				_ = _nativeLen0_61
				for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
					(_nw352).ArraySet1(_init61(_dafny.IntOf(_i0_61)), _i0_61)
				}
			}
			_2330_v9 = _nw352
			var _2332_v10 D11
			_ = _2332_v10
			_2332_v10 = Companion_D11_.Create_DC23_(_2330_v9)
			var _2333_v11 _dafny.Map
			_ = _2333_v11
			_2333_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _2332_v10)
			var _2334_v12 _dafny.Int
			_ = _2334_v12
			_2334_v12 = _dafny.IntOfInt64(472)
			var _rhs350 _dafny.Map = _2333_v11
			_ = _rhs350
			var _rhs351 _dafny.Int = _2329_i2
			_ = _rhs351
			_2333_v11 = _rhs350
			_2334_v12 = _rhs351
			var _index365 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index365
			var _index366 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index366
			var _rhs352 bool = !(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))
			_ = _rhs352
			var _rhs353 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_2329_i2).Times(_dafny.IntOfInt64(697))))
			_ = _rhs353
			var _rhs354 _dafny.Int = ((_dafny.Zero).Minus(_2334_v12)).Minus((_this).Fm7(globalState))
			_ = _rhs354
			var _rhs355 bool = false
			_ = _rhs355
			var _lhs227 _dafny.Array = (_this).F10()
			_ = _lhs227
			var _lhs228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _lhs228
			var _lhs229 _dafny.Array = (_this).F10()
			_ = _lhs229
			var _lhs230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _lhs230
			(_lhs227).ArraySet1(_rhs352, (_lhs228).Int())
			_2334_v12 = _rhs353
			_2334_v12 = _rhs354
			(_lhs229).ArraySet1(_rhs355, (_lhs230).Int())
		}
		var _2335_v13 bool
		_ = _2335_v13
		_2335_v13 = false
		_2335_v13 = _2335_v13
		var _index367 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))
		_ = _index367
		((_this).F10()).ArraySet1((_2335_v13) || ((_this).F4()), (_index367).Int())
		var _index368 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))
		_ = _index368
		((_this).F10()).ArraySet1(_2335_v13, (_index368).Int())
		var _2336_v14 _dafny.Sequence
		_ = _2336_v14
		_2336_v14 = _dafny.UnicodeSeqOfUtf8Bytes("ptmrel")
		var _2337_v15 _dafny.Sequence
		_ = _2337_v15
		_2337_v15 = _dafny.SeqOf(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if ((_this).F9()).Contains((_this).F4()) {
				return ((_this).F9()).Get((_this).F4()).(_dafny.Sequence)
			}
			return _2336_v14
		})()).Cardinality()), (_this).F6(), (_this).F6(), (_this).F6())
		var _2338_v16 _dafny.Sequence
		_ = _2338_v16
		_2338_v16 = _dafny.SeqOf(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))
		var _2339_v17 _dafny.MultiSet
		_ = _2339_v17
		_2339_v17 = _dafny.MultiSetOf(_2338_v16)
		var _2340_v18 _dafny.Map
		_ = _2340_v18
		_2340_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D6_.Create_DC15_(_2339_v17))
		if ((_2337_v15).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2337_v15).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((_2340_v18).Cardinality()) > 0 {
			var _2341_v19 _dafny.Set
			_ = _2341_v19
			_2341_v19 = _dafny.SetOf((_this).F6())
			var _2342_v20 _dafny.Map
			_ = _2342_v20
			_2342_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), (_2341_v19).Cardinality())
			var _2343_v21 D10
			_ = _2343_v21
			_2343_v21 = Companion_D10_.Create_DC21_(_2342_v20)
			var _source38 D10 = _2343_v21
			_ = _source38
			if _source38.Is_DC22() {
				var _2344___mcc_h0 _dafny.Sequence = _source38.Get_().(D10_DC22).Cf30
				_ = _2344___mcc_h0
				var _2345___mcc_h1 bool = _source38.Get_().(D10_DC22).Cf31
				_ = _2345___mcc_h1
				var _2346___mcc_h2 bool = _source38.Get_().(D10_DC22).Cf32
				_ = _2346___mcc_h2
				var _2347___mcc_h3 _dafny.Array = _source38.Get_().(D10_DC22).Cf33
				_ = _2347___mcc_h3
				var _2348___mcc_h4 _dafny.Int = _source38.Get_().(D10_DC22).Cf34
				_ = _2348___mcc_h4
				var _2349_cf34 _dafny.Int = _2348___mcc_h4
				_ = _2349_cf34
				var _2350_cf33 _dafny.Array = _2347___mcc_h3
				_ = _2350_cf33
				var _2351_cf32 bool = _2346___mcc_h2
				_ = _2351_cf32
				var _2352_cf31 bool = _2345___mcc_h1
				_ = _2352_cf31
				var _2353_cf30 _dafny.Sequence = _2344___mcc_h0
				_ = _2353_cf30
				_2353_cf30 = _dafny.Companion_Sequence_.Concatenate(_2353_cf30, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mlwyh"), _2353_cf30))
				var _2354_v22 _dafny.Array
				_ = _2354_v22
				var _len0_62 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_62
				var _nw353 _dafny.Array
				_ = _nw353
				if _len0_62.Cmp(_dafny.Zero) == 0 {
					_nw353 = _dafny.NewArray(_len0_62)
				} else {
					var _init62 func(_dafny.Int) _dafny.Int = func(_2355_i4 _dafny.Int) _dafny.Int {
						return (_2355_i4).Minus(_dafny.IntOfInt64(199))
					}
					_ = _init62
					var _element0_62 = _init62(_dafny.Zero)
					_ = _element0_62
					_nw353 = _dafny.NewArrayFromExample(_element0_62, nil, _len0_62)
					(_nw353).ArraySet1(_element0_62, 0)
					var _nativeLen0_62 = (_len0_62).Int()
					_ = _nativeLen0_62
					for _i0_62 := 1; _i0_62 < _nativeLen0_62; _i0_62++ {
						(_nw353).ArraySet1(_init62(_dafny.IntOf(_i0_62)), _i0_62)
					}
				}
				_2354_v22 = _nw353
				var _2356_v23 _dafny.Map
				_ = _2356_v23
				_2356_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2337_v15).Cardinality()), (_this).F6())
				var _index369 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_2354_v22), 0))
				_ = _index369
				(_2354_v22).ArraySet1((func() _dafny.Int {
					if (_2356_v23).Contains(_dafny.IntOfUint32((_2336_v14).Cardinality())) {
						return (_2356_v23).Get(_dafny.IntOfUint32((_2336_v14).Cardinality())).(_dafny.Int)
					}
					return _2349_cf34
				})(), (_index369).Int())
				var _index370 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_2354_v22), 0))
				_ = _index370
				(_2354_v22).ArraySet1((_this).F6(), (_index370).Int())
				var _2357_v24 D11
				_ = _2357_v24
				_2357_v24 = Companion_D11_.Create_DC25_(_2352_cf31, (_this).F6())
				var _pat_let_tv40 = _2335_v13
				_ = _pat_let_tv40
				var _2358_v25 *C5
				_ = _2358_v25
				var _nw354 *C5 = New_C5_()
				_ = _nw354
				_nw354.Ctor__((func(_pat_let51_0 D11) D11 {
					return func(_2359_dt__update__tmp_h0 D11) D11 {
						return func(_pat_let52_0 bool) D11 {
							return func(_2360_dt__update_hcf39_h0 bool) D11 {
								return Companion_D11_.Create_DC25_(_2360_dt__update_hcf39_h0, (_2359_dt__update__tmp_h0).Dtor_cf40())
							}(_pat_let52_0)
						}(_pat_let_tv40)
					}(_pat_let51_0)
				}(_2357_v24)).Dtor_cf39(), (_dafny.Zero).Minus((_this).F6()), _this.F5(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("romg")).Cardinality()), ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))
				_2358_v25 = _nw354
				var _2361_v26 _dafny.Array
				_ = _2361_v26
				var _len0_63 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_63
				var _nw355 _dafny.Array
				_ = _nw355
				if _len0_63.Cmp(_dafny.Zero) == 0 {
					_nw355 = _dafny.NewArray(_len0_63)
				} else {
					var _init63 func(_dafny.Int) _dafny.MultiSet = (func(_2362_cf32 bool) func(_dafny.Int) _dafny.MultiSet {
						return func(_2363_i5 _dafny.Int) _dafny.MultiSet {
							return _dafny.MultiSetOf(_2362_cf32)
						}
					})(_2351_cf32)
					_ = _init63
					var _element0_63 = _init63(_dafny.Zero)
					_ = _element0_63
					_nw355 = _dafny.NewArrayFromExample(_element0_63, nil, _len0_63)
					(_nw355).ArraySet1(_element0_63, 0)
					var _nativeLen0_63 = (_len0_63).Int()
					_ = _nativeLen0_63
					for _i0_63 := 1; _i0_63 < _nativeLen0_63; _i0_63++ {
						(_nw355).ArraySet1(_init63(_dafny.IntOf(_i0_63)), _i0_63)
					}
				}
				_2361_v26 = _nw355
				_2361_v26 = _2361_v26
			} else {
				var _2364___mcc_h5 _dafny.Map = _source38.Get_().(D10_DC21).Cf29
				_ = _2364___mcc_h5
				var _2365_cf29 _dafny.Map = _2364___mcc_h5
				_ = _2365_cf29
				var _2366_v27 _dafny.Int
				_ = _2366_v27
				_2366_v27 = _dafny.IntOfInt64(371)
				_2366_v27 = _2366_v27
				var _2367_v28 _dafny.MultiSet
				_ = _2367_v28
				_2367_v28 = _dafny.MultiSetOf(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), (_this).F4(), (_this).F4(), _2335_v13)
				var _2368_v29 _dafny.MultiSet
				_ = _2368_v29
				_2368_v29 = _dafny.MultiSetOf(_2366_v27, (_2367_v28).Cardinality())
				var _2369_v30 _dafny.Map
				_ = _2369_v30
				_2369_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2368_v29, _2336_v14)
				var _2370_v31 bool
				_ = _2370_v31
				var _2371_v32 _dafny.Int
				_ = _2371_v32
				var _out62 bool
				_ = _out62
				var _out63 _dafny.Int
				_ = _out63
				_out62, _out63 = Companion_Default___.M0(_2369_v30, ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), _2368_v29, (_dafny.IntOfInt64(253)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2337_v15, (Companion_Default___.SafeIndex(_2366_v27, _dafny.IntOfUint32((_2337_v15).Cardinality()))).Uint32(), _2366_v27)).Cardinality())) > 0, globalState)
				_2370_v31 = _out62
				_2371_v32 = _out63
				_2366_v27 = (_this).F6()
				var _index371 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index371
				((_this).F10()).ArraySet1((_this).F4(), (_index371).Int())
			}
			if _2335_v13 {
				var _2372_v33 _dafny.MultiSet
				_ = _2372_v33
				_2372_v33 = _dafny.MultiSetOf((_this).F6())
				var _2373_v34 _dafny.Map
				_ = _2373_v34
				_2373_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2372_v33, _2336_v14)
				var _2374_v35 bool
				_ = _2374_v35
				var _2375_v36 _dafny.Int
				_ = _2375_v36
				var _out64 bool
				_ = _out64
				var _out65 _dafny.Int
				_ = _out65
				_out64, _out65 = Companion_Default___.M0(_2373_v34, true, _dafny.MultiSetFromSeq(_2337_v15), _2335_v13, globalState)
				_2374_v35 = _out64
				_2375_v36 = _out65
				var _2376_v37 _dafny.Array
				_ = _2376_v37
				var _len0_64 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_64
				var _nw356 _dafny.Array
				_ = _nw356
				if _len0_64.Cmp(_dafny.Zero) == 0 {
					_nw356 = _dafny.NewArray(_len0_64)
				} else {
					var _init64 func(_dafny.Int) _dafny.Set = (func(_2377_v14 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
						return func(_2378_i6 _dafny.Int) _dafny.Set {
							return (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(929))).Uint32(), func(coer114 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg114 _dafny.Int) interface{} {
									return coer114(arg114)
								}
							}(func(_2379_i7 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('a')
							})))).Union(_dafny.SetOf(_2377_v14, _dafny.UnicodeSeqOfUtf8Bytes("mj")))
						}
					})(_2336_v14)
					_ = _init64
					var _element0_64 = _init64(_dafny.Zero)
					_ = _element0_64
					_nw356 = _dafny.NewArrayFromExample(_element0_64, nil, _len0_64)
					(_nw356).ArraySet1(_element0_64, 0)
					var _nativeLen0_64 = (_len0_64).Int()
					_ = _nativeLen0_64
					for _i0_64 := 1; _i0_64 < _nativeLen0_64; _i0_64++ {
						(_nw356).ArraySet1(_init64(_dafny.IntOf(_i0_64)), _i0_64)
					}
				}
				_2376_v37 = _nw356
				var _2380_v38 _dafny.Set
				_ = _2380_v38
				_2380_v38 = _dafny.SetOf(_2336_v14)
				var _index372 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_2376_v37), 0))
				_ = _index372
				(_2376_v37).ArraySet1(_2380_v38, (_index372).Int())
				var _index373 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_2376_v37), 0))
				_ = _index373
				(_2376_v37).ArraySet1(_2380_v38, (_index373).Int())
				_2375_v36 = _2375_v36
				var _2381_v39 *C4
				_ = _2381_v39
				var _nw357 *C4 = New_C4_()
				_ = _nw357
				_nw357.Ctor__(_this.F5(), (_2337_v15).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2337_v15).Cardinality()))).Uint32()).(_dafny.Int), (_this).F4())
				_2381_v39 = _nw357
				var _2382_v40 D4
				_ = _2382_v40
				_2382_v40 = Companion_D4_.Create_DC13_(_2375_v36, _2374_v35, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wdwcp")).Cardinality()), _2338_v16)
				var _2383_v41 _dafny.CodePoint
				_ = _2383_v41
				_2383_v41 = _dafny.CodePoint('i')
				var _rhs356 _dafny.Int = ((_this).F6()).Minus((_2382_v40).Dtor_cf14())
				_ = _rhs356
				var _rhs357 bool = _2335_v13
				_ = _rhs357
				var _rhs358 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_2336_v14, (Companion_Default___.SafeIndex(_2375_v36, _dafny.IntOfUint32((_2336_v14).Cardinality()))).Uint32(), _2383_v41), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_2336_v14, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2336_v14).Cardinality()))).Uint32(), _2383_v41), _2336_v14))).Cardinality())
				_ = _rhs358
				var _rhs359 _dafny.Int = _2375_v36
				_ = _rhs359
				var _rhs360 _dafny.Int = (_2375_v36).Plus(_dafny.IntOfInt64(439))
				_ = _rhs360
				_2375_v36 = _rhs356
				_2374_v35 = _rhs357
				_2375_v36 = _rhs358
				_2375_v36 = _rhs359
				_2375_v36 = _rhs360
			} else {
				var _2384_v42 _dafny.Array
				_ = _2384_v42
				var _len0_65 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_65
				var _nw358 _dafny.Array
				_ = _nw358
				if _len0_65.Cmp(_dafny.Zero) == 0 {
					_nw358 = _dafny.NewArray(_len0_65)
				} else {
					var _init65 func(_dafny.Int) _dafny.Int = func(_2385_i8 _dafny.Int) _dafny.Int {
						return (_2385_i8).Minus((_this).F6())
					}
					_ = _init65
					var _element0_65 = _init65(_dafny.Zero)
					_ = _element0_65
					_nw358 = _dafny.NewArrayFromExample(_element0_65, nil, _len0_65)
					(_nw358).ArraySet1(_element0_65, 0)
					var _nativeLen0_65 = (_len0_65).Int()
					_ = _nativeLen0_65
					for _i0_65 := 1; _i0_65 < _nativeLen0_65; _i0_65++ {
						(_nw358).ArraySet1(_init65(_dafny.IntOf(_i0_65)), _i0_65)
					}
				}
				_2384_v42 = _nw358
				var _index374 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_2384_v42), 0))
				_ = _index374
				(_2384_v42).ArraySet1(((_this).F6()).Minus((_this).F6()), (_index374).Int())
				var _index375 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_2384_v42), 0))
				_ = _index375
				(_2384_v42).ArraySet1((_this).F6(), (_index375).Int())
				var _index376 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_2384_v42), 0))
				_ = _index376
				(_2384_v42).ArraySet1(((_2384_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_2384_v42), 0))).Int()).(_dafny.Int)).Times((_2384_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_2384_v42), 0))).Int()).(_dafny.Int)), (_index376).Int())
				var _2386_v44 _dafny.MultiSet
				_ = _2386_v44
				_2386_v44 = _dafny.MultiSetFromSeq(_2337_v15)
				var _2387_v45 _dafny.Map
				_ = _2387_v45
				_2387_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2386_v44, ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))
				var _2388_v46 _dafny.Map
				_ = _2388_v46
				_2388_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					var _coll59 = _dafny.NewMapBuilder()
					_ = _coll59
					for _iter66 := _dafny.Iterate((_2387_v45).Keys().Elements()); ; {
						_compr_59, _ok66 := _iter66()
						if !_ok66 {
							break
						}
						var _2389_v43 _dafny.MultiSet
						_2389_v43 = interface{}(_compr_59).(_dafny.MultiSet)
						if (_2387_v45).Contains(_2389_v43) {
							_coll59.Add(_2389_v43, _dafny.IntOfUint32((_2336_v14).Cardinality()))
						}
					}
					return _coll59.ToMap()
				}()).Cardinality(), (_this).F4())
				_2388_v46 = _2388_v46
				_2335_v13 = _2335_v13
				var _index377 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_2384_v42), 0))
				_ = _index377
				(_2384_v42).ArraySet1((_this).F6(), (_index377).Int())
			}
			_2337_v15 = _dafny.SeqOf(Companion_Default___.SafeModuloInt((_this).F6(), (_2342_v20).Cardinality()), (_this).F6())
			var _2390_v47 _dafny.Int
			_ = _2390_v47
			_2390_v47 = _dafny.IntOfInt64(355)
			_2390_v47 = (func() _dafny.Int {
				if _2335_v13 {
					return (_this).F6()
				}
				return _2390_v47
			})()
			var _index378 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index378
			((_this).F10()).ArraySet1(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), (_index378).Int())
		} else {
			var _2391_v48 _dafny.Array
			_ = _2391_v48
			var _nw359 _dafny.Array = _dafny.NewArrayWithValue(Companion_D21_.Default(), _dafny.One)
			_ = _nw359
			_2391_v48 = _nw359
			var _2392_v49 _dafny.Array
			_ = _2392_v49
			var _len0_66 _dafny.Int = _dafny.One
			_ = _len0_66
			var _nw360 _dafny.Array
			_ = _nw360
			if _len0_66.Cmp(_dafny.Zero) == 0 {
				_nw360 = _dafny.NewArray(_len0_66)
			} else {
				var _init66 func(_dafny.Int) D0 = func(_2393_i9 _dafny.Int) D0 {
					return Companion_D0_.Create_DC1_()
				}
				_ = _init66
				var _element0_66 = _init66(_dafny.Zero)
				_ = _element0_66
				_nw360 = _dafny.NewArrayFromExample(_element0_66, nil, _len0_66)
				(_nw360).ArraySet1(_element0_66, 0)
				var _nativeLen0_66 = (_len0_66).Int()
				_ = _nativeLen0_66
				for _i0_66 := 1; _i0_66 < _nativeLen0_66; _i0_66++ {
					(_nw360).ArraySet1(_init66(_dafny.IntOf(_i0_66)), _i0_66)
				}
			}
			_2392_v49 = _nw360
			var _2394_v50 D21
			_ = _2394_v50
			_2394_v50 = Companion_D21_.Create_DC42_(_2392_v49)
			var _pat_let_tv41 = _2392_v49
			_ = _pat_let_tv41
			var _index379 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_2391_v48), 0))
			_ = _index379
			(_2391_v48).ArraySet1(func(_pat_let53_0 D21) D21 {
				return func(_2395_dt__update__tmp_h1 D21) D21 {
					return func(_pat_let54_0 _dafny.Array) D21 {
						return func(_2396_dt__update_hcf68_h0 _dafny.Array) D21 {
							return Companion_D21_.Create_DC42_(_2396_dt__update_hcf68_h0)
						}(_pat_let54_0)
					}(_pat_let_tv41)
				}(_pat_let53_0)
			}(_2394_v50), (_index379).Int())
			var _index380 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_2391_v48), 0))
			_ = _index380
			(_2391_v48).ArraySet1(_2394_v50, (_index380).Int())
			var _2397_v51 _dafny.CodePoint
			_ = _2397_v51
			_2397_v51 = _dafny.CodePoint('a')
			_2397_v51 = _2397_v51
			var _2398_v52 _dafny.Set
			_ = _2398_v52
			_2398_v52 = _dafny.SetOf(true)
			var _index381 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index381
			((_this).F10()).ArraySet1((_2398_v52).IsSubsetOf(_2398_v52), (_index381).Int())
			var _2399_v53 _dafny.Sequence
			_ = _2399_v53
			_2399_v53 = _dafny.SeqOf(_2337_v15, _2337_v15, _2337_v15)
			var _2400_v54 _dafny.MultiSet
			_ = _2400_v54
			_2400_v54 = _dafny.MultiSetOf((_this).F6(), (_this).F6(), _dafny.IntOfUint32((_2399_v53).Cardinality()))
			var _2401_v55 _dafny.Map
			_ = _2401_v55
			_2401_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2400_v54).Cardinality(), ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool))
			var _index382 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index382
			((_this).F10()).ArraySet1((_2335_v13) == ((_this).Fm14(_2401_v55, !((_this).F4()), (_this).F6(), ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), globalState)), (_index382).Int())
			var _2402_v56 _dafny.Array
			_ = _2402_v56
			var _len0_67 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_67
			var _nw361 _dafny.Array
			_ = _nw361
			if _len0_67.Cmp(_dafny.Zero) == 0 {
				_nw361 = _dafny.NewArray(_len0_67)
			} else {
				var _init67 func(_dafny.Int) _dafny.Int = func(_2403_i10 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_2403_i10, (_this).F6())
				}
				_ = _init67
				var _element0_67 = _init67(_dafny.Zero)
				_ = _element0_67
				_nw361 = _dafny.NewArrayFromExample(_element0_67, nil, _len0_67)
				(_nw361).ArraySet1(_element0_67, 0)
				var _nativeLen0_67 = (_len0_67).Int()
				_ = _nativeLen0_67
				for _i0_67 := 1; _i0_67 < _nativeLen0_67; _i0_67++ {
					(_nw361).ArraySet1(_init67(_dafny.IntOf(_i0_67)), _i0_67)
				}
			}
			_2402_v56 = _nw361
			var _index383 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_2402_v56), 0))
			_ = _index383
			(_2402_v56).ArraySet1(((_this).F6()).Plus(Companion_Default___.Fm3((_this).F6(), true, globalState)), (_index383).Int())
			var _index384 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_2402_v56), 0))
			_ = _index384
			(_2402_v56).ArraySet1((_this).F6(), (_index384).Int())
		}
		var _2404_v57 _dafny.CodePoint
		_ = _2404_v57
		_2404_v57 = _dafny.CodePoint('n')
		if Companion_Default___.Fm0(_2404_v57, globalState) {
			var _index385 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index385
			((_this).F10()).ArraySet1(_2335_v13, (_index385).Int())
			var _2405_v58 *C6
			_ = _2405_v58
			var _nw362 *C6 = New_C6_()
			_ = _nw362
			_nw362.Ctor__((_this).F4(), _2335_v13, _this.F5(), (func() _dafny.Int {
				if false {
					return (_this).F6()
				}
				return (_this).F6()
			})())
			_2405_v58 = _nw362
			var _2406_v59 _dafny.Int
			_ = _2406_v59
			_2406_v59 = _dafny.IntOfInt64(-792)
			var _index386 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index386
			var _rhs361 *C6 = _2405_v58
			_ = _rhs361
			var _rhs362 bool = ((_this).Fm9((_2405_v58).F11(), ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), _dafny.IntOfUint32((_2336_v14).Cardinality()), globalState)).Cmp(_2406_v59) == 0
			_ = _rhs362
			var _rhs363 _dafny.Int = _2406_v59
			_ = _rhs363
			var _lhs231 _dafny.Array = (_this).F10()
			_ = _lhs231
			var _lhs232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _lhs232
			_2405_v58 = _rhs361
			(_lhs231).ArraySet1(_rhs362, (_lhs232).Int())
			_2406_v59 = _rhs363
			_2406_v59 = (_this).F6()
			_2406_v59 = _2406_v59
			var _2407_v60 _dafny.Array
			_ = _2407_v60
			var _len0_68 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_68
			var _nw363 _dafny.Array
			_ = _nw363
			if _len0_68.Cmp(_dafny.Zero) == 0 {
				_nw363 = _dafny.NewArray(_len0_68)
			} else {
				var _init68 func(_dafny.Int) _dafny.Int = (func(_2408_v59 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2409_i11 _dafny.Int) _dafny.Int {
						return (_2409_i11).Times(_2408_v59)
					}
				})(_2406_v59)
				_ = _init68
				var _element0_68 = _init68(_dafny.Zero)
				_ = _element0_68
				_nw363 = _dafny.NewArrayFromExample(_element0_68, nil, _len0_68)
				(_nw363).ArraySet1(_element0_68, 0)
				var _nativeLen0_68 = (_len0_68).Int()
				_ = _nativeLen0_68
				for _i0_68 := 1; _i0_68 < _nativeLen0_68; _i0_68++ {
					(_nw363).ArraySet1(_init68(_dafny.IntOf(_i0_68)), _i0_68)
				}
			}
			_2407_v60 = _nw363
			var _index387 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_2407_v60), 0))
			_ = _index387
			(_2407_v60).ArraySet1(Companion_Default___.SafeDivisionInt(_2406_v59, (_this).F6()), (_index387).Int())
			var _2410_v61 _dafny.MultiSet
			_ = _2410_v61
			_2410_v61 = _dafny.MultiSetOf((_dafny.Zero).Minus(_2406_v59))
			var _2411_v62 _dafny.Map
			_ = _2411_v62
			_2411_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2410_v61, _2335_v13)
			var _2412_v63 _dafny.MultiSet
			_ = _2412_v63
			_2412_v63 = _2410_v61
			var _2413_v64 _dafny.Map
			_ = _2413_v64
			_2413_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2406_v59, (_2412_v63))
			var _2414_v65 _dafny.Map
			_ = _2414_v65
			_2414_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), (func() bool {
				if (_2411_v62).Contains((func() _dafny.MultiSet {
					if (_2413_v64).Contains(_2406_v59) {
						return (_2413_v64).Get(_2406_v59).(_dafny.MultiSet)
					}
					return _2410_v61
				})()) {
					return (_2411_v62).Get((func() _dafny.MultiSet {
						if (_2413_v64).Contains(_2406_v59) {
							return (_2413_v64).Get(_2406_v59).(_dafny.MultiSet)
						}
						return _2410_v61
					})()).(bool)
				}
				return (_2405_v58).F11()
			})())
			var _index388 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_2407_v60), 0))
			_ = _index388
			var _rhs364 _dafny.Int = (_dafny.Zero).Minus((_2406_v59).Minus((_dafny.IntOfUint32((_2337_v15).Cardinality())).Times(_2406_v59)))
			_ = _rhs364
			var _rhs365 bool = (func() bool {
				if (_2414_v65).Contains((_2405_v58).F11()) {
					return (_2414_v65).Get((_2405_v58).F11()).(bool)
				}
				return (_2405_v58).F11()
			})()
			_ = _rhs365
			var _rhs366 _dafny.Int = (_this).Fm7(globalState)
			_ = _rhs366
			var _rhs367 *C6 = _2405_v58
			_ = _rhs367
			var _lhs233 _dafny.Array = _2407_v60
			_ = _lhs233
			var _lhs234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_2407_v60), 0))
			_ = _lhs234
			(_lhs233).ArraySet1(_rhs364, (_lhs234).Int())
			_2335_v13 = _rhs365
			_2406_v59 = _rhs366
			_2405_v58 = _rhs367
		} else {
			var _2415_v66 _dafny.Int
			_ = _2415_v66
			_2415_v66 = _dafny.IntOfInt64(313)
			_2415_v66 = Companion_Default___.SafeDivisionInt(_2415_v66, ((_this).F6()).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gpm")).Cardinality())))
			var _2416_v67 _dafny.Array
			_ = _2416_v67
			var _nw364 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw364
			_2416_v67 = _nw364
			_2416_v67 = _2416_v67
			var _2417_v68 D9
			_ = _2417_v68
			_2417_v68 = Companion_D9_.Create_DC19_(_dafny.UnicodeSeqOfUtf8Bytes("p"))
			var _2418_v69 _dafny.Map
			_ = _2418_v69
			_2418_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2417_v68, (_this).F4())
			var _2419_v70 _dafny.Map
			_ = _2419_v70
			_2419_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2418_v69).Merge(_2418_v69), _2404_v57)
			_2419_v70 = (_2419_v70).Update(_2418_v69, _2404_v57)
			_2415_v66 = (_dafny.Zero).Minus((_2415_v66).Plus(_2415_v66))
			var _index389 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index389
			((_this).F10()).ArraySet1((((_this).F6()).Cmp((_this).F6()) == 0) && (!(!_dafny.Companion_Sequence_.Equal(_2337_v15, _2337_v15))), (_index389).Int())
		}
		var _2420_v71 _dafny.Set
		_ = _2420_v71
		_2420_v71 = _dafny.SetOf((_this).F6())
		var _2421_v72 _dafny.Map
		_ = _2421_v72
		_2421_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F6()), !(!(_2335_v13)))
		var _2422_v73 _dafny.Array
		_ = _2422_v73
		var _nwElement0_75 _dafny.Int = (_this).F6()
		_ = _nwElement0_75
		var _nw365 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(28))
		_ = _nw365
		(_nw365).ArraySet1(_nwElement0_75, 0)
		(_nw365).ArraySet1((_this).F6(), 1)
		(_nw365).ArraySet1((_this).F6(), 2)
		(_nw365).ArraySet1((_2420_v71).Cardinality(), 3)
		(_nw365).ArraySet1((_this).F6(), 4)
		(_nw365).ArraySet1((_this).F6(), 5)
		(_nw365).ArraySet1((_this).F6(), 6)
		(_nw365).ArraySet1(_dafny.IntOfInt64(617), 7)
		(_nw365).ArraySet1((_this).F6(), 8)
		(_nw365).ArraySet1(_dafny.IntOfInt64(844), 9)
		(_nw365).ArraySet1((_this).F6(), 10)
		(_nw365).ArraySet1((_this).F6(), 11)
		(_nw365).ArraySet1(_dafny.IntOfInt64(-957), 12)
		(_nw365).ArraySet1((_this).F6(), 13)
		(_nw365).ArraySet1(_dafny.IntOfInt64(860), 14)
		(_nw365).ArraySet1(_dafny.IntOfUint32((_2336_v14).Cardinality()), 15)
		(_nw365).ArraySet1((_dafny.Zero).Minus((_this).F6()), 16)
		(_nw365).ArraySet1((_this).F6(), 17)
		(_nw365).ArraySet1((_this).F6(), 18)
		(_nw365).ArraySet1((_this).F6(), 19)
		(_nw365).ArraySet1(_dafny.IntOfUint32((_2336_v14).Cardinality()), 20)
		(_nw365).ArraySet1((_this).F6(), 21)
		(_nw365).ArraySet1((_this).F6(), 22)
		(_nw365).ArraySet1((_this).F6(), 23)
		(_nw365).ArraySet1((_this).F6(), 24)
		(_nw365).ArraySet1((_this).F6(), 25)
		(_nw365).ArraySet1((_2421_v72).Cardinality(), 26)
		(_nw365).ArraySet1(_dafny.IntOfInt64(682), 27)
		_2422_v73 = _nw365
		var _2423_v74 _dafny.Map
		_ = _2423_v74
		_2423_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2422_v73, (_this).F6())
		var _pat_let_tv42 = _2423_v74
		_ = _pat_let_tv42
		r0 = func(_pat_let55_0 D1) D1 {
			return func(_2426_dt__update__tmp_h2 D1) D1 {
				return func(_pat_let56_0 bool) D1 {
					return func(_2427_dt__update_hcf2_h0 bool) D1 {
						return Companion_D1_.Create_DC3_(_2427_dt__update_hcf2_h0)
					}(_pat_let56_0)
				}(((_this).F6()).Cmp((_pat_let_tv42).Cardinality()) >= 0)
			}(_pat_let55_0)
		}(Companion_D1_.Create_DC3_((_this).Fm8((_this).F6(), _2335_v13, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-493))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg115 _dafny.Int) interface{} {
				return coer115(arg115)
			}
		}((func(_2424_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2425_i12 _dafny.Int) _dafny.CodePoint {
				return _2424_v57
			}
		})(_2404_v57))), _2337_v15, globalState)))
		r1 = _2336_v14
		return r0, r1
	}
}
func (_this *C7) F10() _dafny.Array {
	{
		return _this._f10
	}
}
func (_this *C7) F9() _dafny.Map {
	{
		return _this._f9
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
func (_this *C8) Fm11(globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("prxjdttbg"), _dafny.UnicodeSeqOfUtf8Bytes("or"))
	}
}
func (_this *C8) Fm12(p0 bool, p1 D0, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(157))).Uint32(), func(coer116 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg116 _dafny.Int) interface{} {
				return coer116(arg116)
			}
		}(func(_2428_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-632)
		})), _dafny.SeqOf(_dafny.IntOfInt64(-337), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(854), _dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality())))).Cardinality())
	}
}
func (_this *C8) M4(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) {
	{
		var _2429_v0 D0
		_ = _2429_v0
		_2429_v0 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(986))
		var _2430_i0 _dafny.Int
		_ = _2430_i0
		_2430_i0 = _dafny.Zero
		{
			for !((_this).Fm12(p2, _2429_v0, p1, p2, globalState)) {
				{
					if (_2430_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_2430_i0 = (_2430_i0).Plus(_dafny.One)
					var _2431_v1 _dafny.CodePoint
					_ = _2431_v1
					_2431_v1 = _dafny.CodePoint('h')
					_2431_v1 = _2431_v1
					var _2432_v2 _dafny.Sequence
					_ = _2432_v2
					_2432_v2 = _dafny.UnicodeSeqOfUtf8Bytes("mhlv")
					_2432_v2 = _dafny.Companion_Sequence_.Concatenate(_2432_v2, _2432_v2)
					var _2433_v3 bool
					_ = _2433_v3
					_2433_v3 = true
					var _2434_v4 _dafny.Array
					_ = _2434_v4
					var _nw366 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
					_ = _nw366
					_2434_v4 = _nw366
					var _rhs368 bool = p2
					_ = _rhs368
					var _rhs369 _dafny.Array = _2434_v4
					_ = _rhs369
					_2433_v3 = _rhs368
					_2434_v4 = _rhs369
					var _2435_v5 _dafny.Int
					_ = _2435_v5
					_2435_v5 = _dafny.IntOfInt64(337)
					var _rhs370 bool = ((_dafny.Zero).Minus(p0)).Cmp(_dafny.IntOfUint32((_2432_v2).Cardinality())) >= 0
					_ = _rhs370
					var _rhs371 _dafny.Int = _dafny.IntOfUint32((_2432_v2).Cardinality())
					_ = _rhs371
					_2433_v3 = _rhs370
					_2435_v5 = _rhs371
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _2436_v6 _dafny.Int
		_ = _2436_v6
		_2436_v6 = _dafny.IntOfInt64(631)
		_2436_v6 = (p1).Times(p0)
		var _2437_v7 _dafny.Array
		_ = _2437_v7
		var _nw367 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
		_ = _nw367
		_2437_v7 = _nw367
		var _index390 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_2437_v7), 0))
		_ = _index390
		(_2437_v7).ArraySet1((func() _dafny.Int {
			if true {
				return _dafny.IntOfInt64(194)
			}
			return p1
		})(), (_index390).Int())
		var _2438_v8 D11
		_ = _2438_v8
		_2438_v8 = Companion_D11_.Create_DC26_(p2)
		var _2439_v9 _dafny.Map
		_ = _2439_v9
		_2439_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		var _2440_v10 _dafny.Sequence
		_ = _2440_v10
		_2440_v10 = _dafny.UnicodeSeqOfUtf8Bytes("drtvj")
		var _2441_v11 D12
		_ = _2441_v11
		_2441_v11 = Companion_D12_.Create_DC28_(p0, _dafny.IntOfUint32((_2440_v10).Cardinality()), p2, !(p2), p0)
		var _2442_v12 _dafny.CodePoint
		_ = _2442_v12
		_2442_v12 = _dafny.CodePoint('a')
		var _2443_v13 _dafny.Sequence
		_ = _2443_v13
		_2443_v13 = _dafny.SeqOf(p2)
		var _2444_v14 D21
		_ = _2444_v14
		_2444_v14 = Companion_D21_.Create_DC43_(p2, false, _2443_v13)
		var _2445_v15 _dafny.Sequence
		_ = _2445_v15
		_2445_v15 = _dafny.SeqOf(p2, (_2444_v14).Dtor_cf70(), p2)
		var _2446_v16 _dafny.Array
		_ = _2446_v16
		var _nwElement0_76 bool = p2
		_ = _nwElement0_76
		var _nw368 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(21))
		_ = _nw368
		(_nw368).ArraySet1(_nwElement0_76, 0)
		(_nw368).ArraySet1(p2, 1)
		(_nw368).ArraySet1(true, 2)
		(_nw368).ArraySet1(p2, 3)
		(_nw368).ArraySet1(p2, 4)
		(_nw368).ArraySet1((_2438_v8).Dtor_cf41(), 5)
		(_nw368).ArraySet1(p2, 6)
		(_nw368).ArraySet1(p2, 7)
		(_nw368).ArraySet1(p2, 8)
		(_nw368).ArraySet1(p2, 9)
		(_nw368).ArraySet1(p2, 10)
		(_nw368).ArraySet1(p2, 11)
		(_nw368).ArraySet1(p2, 12)
		(_nw368).ArraySet1(p2, 13)
		(_nw368).ArraySet1(p2, 14)
		(_nw368).ArraySet1(p2, 15)
		(_nw368).ArraySet1((func() bool {
			if (_2439_v9).Contains(p2) {
				return (_2439_v9).Get(p2).(bool)
			}
			return true
		})(), 16)
		(_nw368).ArraySet1(p2, 17)
		(_nw368).ArraySet1((Companion_D9_.Create_DC20_((_2441_v11).Dtor_cf46(), _2442_v12, _2436_v6)).Dtor_cf26(), 18)
		(_nw368).ArraySet1(!((_2443_v13).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2443_v13).Cardinality()))).Uint32()).(bool)), 19)
		(_nw368).ArraySet1(p2, 20)
		_2446_v16 = _nw368
		var _2447_v17 _dafny.Array
		_ = _2447_v17
		var _nwElement0_77 _dafny.Array = _2446_v16
		_ = _nwElement0_77
		var _nw369 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(28))
		_ = _nw369
		(_nw369).ArraySet1(_nwElement0_77, 0)
		(_nw369).ArraySet1(_2446_v16, 1)
		(_nw369).ArraySet1(_2446_v16, 2)
		(_nw369).ArraySet1(_2446_v16, 3)
		(_nw369).ArraySet1(_2446_v16, 4)
		(_nw369).ArraySet1(_2446_v16, 5)
		(_nw369).ArraySet1(_2446_v16, 6)
		(_nw369).ArraySet1(_2446_v16, 7)
		(_nw369).ArraySet1(_2446_v16, 8)
		(_nw369).ArraySet1(_2446_v16, 9)
		(_nw369).ArraySet1(_2446_v16, 10)
		(_nw369).ArraySet1(_2446_v16, 11)
		(_nw369).ArraySet1(_2446_v16, 12)
		(_nw369).ArraySet1(_2446_v16, 13)
		(_nw369).ArraySet1(_2446_v16, 14)
		(_nw369).ArraySet1(_2446_v16, 15)
		(_nw369).ArraySet1(_2446_v16, 16)
		(_nw369).ArraySet1(_2446_v16, 17)
		(_nw369).ArraySet1(_2446_v16, 18)
		(_nw369).ArraySet1(_2446_v16, 19)
		(_nw369).ArraySet1(_2446_v16, 20)
		(_nw369).ArraySet1(_2446_v16, 21)
		(_nw369).ArraySet1(_2446_v16, 22)
		(_nw369).ArraySet1(_2446_v16, 23)
		(_nw369).ArraySet1(_2446_v16, 24)
		(_nw369).ArraySet1(_2446_v16, 25)
		(_nw369).ArraySet1(_2446_v16, 26)
		(_nw369).ArraySet1(_2446_v16, 27)
		_2447_v17 = _nw369
		var _2448_v18 _dafny.MultiSet
		_ = _2448_v18
		_2448_v18 = _dafny.MultiSetOf(_dafny.IntOfInt64(-502))
		var _2449_v19 T1
		_ = _2449_v19
		var _nw370 *C3 = New_C3_()
		_ = _nw370
		_nw370.Ctor__(_2447_v17, (_2436_v6).Times(p0), (_2448_v18).IsProperSubsetOf(_2448_v18))
		_2449_v19 = _nw370
		var _arr13 _dafny.Array = _2449_v19.F5()
		_ = _arr13
		var _index391 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_2449_v19.F5()), 0))
		_ = _index391
		(_arr13).ArraySet1(_2446_v16, (_index391).Int())
		var _2450_v20 _dafny.Map
		_ = _2450_v20
		_2450_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2436_v6)
		var _2451_v21 _dafny.MultiSet
		_ = _2451_v21
		_2451_v21 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_2445_v15, (Companion_Default___.SafeIndex((_2450_v20).Cardinality(), _dafny.IntOfUint32((_2445_v15).Cardinality()))).Uint32(), (_2449_v19).F4()))
		var _2452_v22 D6
		_ = _2452_v22
		_2452_v22 = Companion_D6_.Create_DC15_(_2451_v21)
		var _pat_let_tv43 = _2442_v12
		_ = _pat_let_tv43
		var _pat_let_tv44 = _2440_v10
		_ = _pat_let_tv44
		var _pat_let_tv45 = _2436_v6
		_ = _pat_let_tv45
		var _pat_let_tv46 = _2440_v10
		_ = _pat_let_tv46
		var _index392 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_2437_v7), 0))
		_ = _index392
		var _arr14 _dafny.Array = _2449_v19.F5()
		_ = _arr14
		var _index393 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_2449_v19.F5()), 0))
		_ = _index393
		var _rhs372 _dafny.Int = p0
		_ = _rhs372
		var _rhs373 T1 = _2449_v19
		_ = _rhs373
		var _rhs374 _dafny.CodePoint = func(_source39 D6) _dafny.CodePoint {
			if _source39.Is_DC16() {
				var _2453___mcc_h0 bool = _source39.Get_().(D6_DC16).Cf20
				_ = _2453___mcc_h0
				var _2454___mcc_h1 bool = _source39.Get_().(D6_DC16).Cf21
				_ = _2454___mcc_h1
				var _2455___mcc_h2 _dafny.Sequence = _source39.Get_().(D6_DC16).Cf22
				_ = _2455___mcc_h2
				var _2456_cf22 _dafny.Sequence = _2455___mcc_h2
				_ = _2456_cf22
				var _2457_cf21 bool = _2454___mcc_h1
				_ = _2457_cf21
				var _2458_cf20 bool = _2453___mcc_h0
				_ = _2458_cf20
				return _pat_let_tv43
			} else {
				var _2459___mcc_h3 _dafny.MultiSet = _source39.Get_().(D6_DC15).Cf19
				_ = _2459___mcc_h3
				var _2460_cf19 _dafny.MultiSet = _2459___mcc_h3
				_ = _2460_cf19
				return (_pat_let_tv44).Select((Companion_Default___.SafeIndex(_pat_let_tv45, _dafny.IntOfUint32((_pat_let_tv46).Cardinality()))).Uint32()).(_dafny.CodePoint)
			}
		}(_2452_v22)
		_ = _rhs374
		var _rhs375 _dafny.Array = _2446_v16
		_ = _rhs375
		var _lhs235 _dafny.Array = _2437_v7
		_ = _lhs235
		var _lhs236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_2437_v7), 0))
		_ = _lhs236
		var _lhs237 _dafny.Array = _2449_v19.F5()
		_ = _lhs237
		var _lhs238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_2449_v19.F5()), 0))
		_ = _lhs238
		(_lhs235).ArraySet1(_rhs372, (_lhs236).Int())
		_2449_v19 = _rhs373
		_2442_v12 = _rhs374
		(_lhs237).ArraySet1(_rhs375, (_lhs238).Int())
		var _2461_i1 _dafny.Int
		_ = _2461_i1
		_2461_i1 = _dafny.Zero
		{
			for (_2449_v19).F4() {
				{
					if (_2461_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_2461_i1 = (_2461_i1).Plus(_dafny.One)
					_2437_v7 = _2437_v7
					var _2462_v23 _dafny.Array
					_ = _2462_v23
					var _len0_69 _dafny.Int = _dafny.IntOfInt64(23)
					_ = _len0_69
					var _nw371 _dafny.Array
					_ = _nw371
					if _len0_69.Cmp(_dafny.Zero) == 0 {
						_nw371 = _dafny.NewArray(_len0_69)
					} else {
						var _init69 func(_dafny.Int) _dafny.Sequence = (func(_2463_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
							return func(_2464_i2 _dafny.Int) _dafny.Sequence {
								return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(205))).Uint32(), func(coer117 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg117 _dafny.Int) interface{} {
										return coer117(arg117)
									}
								}((func(_2465_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_2466_i3 _dafny.Int) _dafny.CodePoint {
										return _2465_v12
									}
								})(_2463_v12)))
							}
						})(_2442_v12)
						_ = _init69
						var _element0_69 = _init69(_dafny.Zero)
						_ = _element0_69
						_nw371 = _dafny.NewArrayFromExample(_element0_69, nil, _len0_69)
						(_nw371).ArraySet1(_element0_69, 0)
						var _nativeLen0_69 = (_len0_69).Int()
						_ = _nativeLen0_69
						for _i0_69 := 1; _i0_69 < _nativeLen0_69; _i0_69++ {
							(_nw371).ArraySet1(_init69(_dafny.IntOf(_i0_69)), _i0_69)
						}
					}
					_2462_v23 = _nw371
					var _2467_v24 D1
					_ = _2467_v24
					_2467_v24 = Companion_D1_.Create_DC4_((_2437_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_2437_v7), 0))).Int()).(_dafny.Int), _2436_v6)
					var _2468_v25 _dafny.Sequence
					_ = _2468_v25
					_2468_v25 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(Companion_Default___.Fm47((_2449_v19).F6(), globalState))).Cardinality()))
					var _2469_v26 *C1
					_ = _2469_v26
					var _nw372 *C1 = New_C1_()
					_ = _nw372
					_nw372.Ctor__((_2437_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_2437_v7), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-616), _2449_v19.F5(), _dafny.IntOfUint32((_2443_v13).Cardinality()), p2)
					_2469_v26 = _nw372
					var _index394 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_2462_v23), 0))
					_ = _index394
					(_2462_v23).ArraySet1(Companion_Default___.Fm50((_2467_v24).Dtor_cf3(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2468_v25, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2468_v25).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.SeqOf(p1, (_2449_v19).F6())).Cardinality()))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_2442_v12, globalState), _2469_v26)).Cardinality(), (_2449_v19).F4(), globalState), (_index394).Int())
					var _index395 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_2462_v23), 0))
					_ = _index395
					(_2462_v23).ArraySet1(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
						if (_2449_v19).F4() {
							return _dafny.Companion_Sequence_.Concatenate(_2440_v10, _2440_v10)
						}
						return _2440_v10
					})(), (Companion_Default___.SafeIndex((_dafny.IntOfUint32((_2440_v10).Cardinality())).Plus(_dafny.IntOfUint32((_2440_v10).Cardinality())), _dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_2449_v19).F4() {
							return _dafny.Companion_Sequence_.Concatenate(_2440_v10, _2440_v10)
						}
						return _2440_v10
					})()).Cardinality()))).Uint32(), _2442_v12), (_index395).Int())
					var _2470_v27 _dafny.Map
					_ = _2470_v27
					_2470_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), p0)
					_2470_v27 = (_2470_v27).Update(_2442_v12, p1)
					_2447_v17 = _2449_v19.F5()
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
		var _2471_v28 _dafny.MultiSet
		_ = _2471_v28
		_2471_v28 = _2448_v18
		var _2472_v29 _dafny.Sequence
		_ = _2472_v29
		_2472_v29 = _dafny.SeqOf(_2448_v18, _2471_v28, _2471_v28)
		var _2473_v30 _dafny.Set
		_ = _2473_v30
		_2473_v30 = _dafny.SetOf(_2472_v29)
		_2473_v30 = _2473_v30
		var _2474_v31 bool
		_ = _2474_v31
		_2474_v31 = false
		_2474_v31 = (_2449_v19).F4()
	}
}
func (_this *C8) M5(p0 _dafny.Int, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _2475_v0 bool
		_ = _2475_v0
		_2475_v0 = false
		if _2475_v0 {
			r0 = !(_2475_v0)
			var _2476_v1 _dafny.Set
			_ = _2476_v1
			_2476_v1 = _dafny.SetOf(_2475_v0, _2475_v0)
			r1 = (_2476_v1).IsSubsetOf(_2476_v1)
			var _2477_v2 _dafny.Map
			_ = _2477_v2
			_2477_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (p0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(259))).Cardinality())))
			_2477_v2 = (_2477_v2).Update(p0, p0)
			r1 = (p0).Cmp(p0) <= 0
			var _2478_v3 _dafny.Array
			_ = _2478_v3
			var _nw373 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw373
			_2478_v3 = _nw373
			var _2479_v4 *C3
			_ = _2479_v4
			var _nw374 *C3 = New_C3_()
			_ = _nw374
			_nw374.Ctor__(_2478_v3, p0, _2475_v0)
			_2479_v4 = _nw374
		} else {
			var _2480_v5 _dafny.Sequence
			_ = _2480_v5
			_2480_v5 = _dafny.UnicodeSeqOfUtf8Bytes("h")
			_2475_v0 = ((p0).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_2480_v5).Cardinality())))).Cmp(_dafny.IntOfInt64(793)) == 0
			if !(_2475_v0) {
				var _2481_v6 D2
				_ = _2481_v6
				_2481_v6 = Companion_D2_.Create_DC6_()
				_2481_v6 = _2481_v6
				var _2482_v7 _dafny.Sequence
				_ = _2482_v7
				_2482_v7 = _dafny.SeqOf(_dafny.IntOfInt64(270), (_dafny.Zero).Minus(p0))
				var _2483_v8 _dafny.Sequence
				_ = _2483_v8
				_2483_v8 = _dafny.SeqOf(_2475_v0, _2475_v0)
				var _2484_v9 _dafny.MultiSet
				_ = _2484_v9
				_2484_v9 = _dafny.MultiSetOf((_dafny.MultiSetFromSeq(_2483_v8)).Cardinality())
				var _2485_v10 _dafny.Array
				_ = _2485_v10
				var _nwElement0_78 _dafny.Int = p0
				_ = _nwElement0_78
				var _nw375 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_78, nil, _dafny.IntOfInt64(23))
				_ = _nw375
				(_nw375).ArraySet1(_nwElement0_78, 0)
				(_nw375).ArraySet1(_dafny.IntOfUint32((_2480_v5).Cardinality()), 1)
				(_nw375).ArraySet1(p0, 2)
				(_nw375).ArraySet1(_dafny.IntOfUint32((_2482_v7).Cardinality()), 3)
				(_nw375).ArraySet1(_dafny.IntOfInt64(-91), 4)
				(_nw375).ArraySet1(p0, 5)
				(_nw375).ArraySet1(p0, 6)
				(_nw375).ArraySet1(p0, 7)
				(_nw375).ArraySet1((_2484_v9).Cardinality(), 8)
				(_nw375).ArraySet1(p0, 9)
				(_nw375).ArraySet1(p0, 10)
				(_nw375).ArraySet1(p0, 11)
				(_nw375).ArraySet1(_dafny.IntOfInt64(-870), 12)
				(_nw375).ArraySet1(p0, 13)
				(_nw375).ArraySet1(p0, 14)
				(_nw375).ArraySet1(p0, 15)
				(_nw375).ArraySet1(p0, 16)
				(_nw375).ArraySet1(p0, 17)
				(_nw375).ArraySet1(_dafny.IntOfInt64(951), 18)
				(_nw375).ArraySet1(_dafny.IntOfInt64(424), 19)
				(_nw375).ArraySet1(p0, 20)
				(_nw375).ArraySet1(p0, 21)
				(_nw375).ArraySet1(p0, 22)
				_2485_v10 = _nw375
				var _2486_v11 _dafny.Set
				_ = _2486_v11
				_2486_v11 = _dafny.SetOf(_2485_v10, _2485_v10, _2485_v10)
				var _2487_v12 _dafny.Set
				_ = _2487_v12
				_2487_v12 = _2486_v11
				var _2488_v13 _dafny.Set
				_ = _2488_v13
				_2488_v13 = _dafny.SetOf(_2486_v11, _2487_v12)
				_2475_v0 = (_2488_v13).IsProperSubsetOf(_2488_v13)
				var _2489_v14 _dafny.Map
				_ = _2489_v14
				_2489_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Times(p0), _2475_v0)
				_2489_v14 = (_2489_v14).Update((p0).Minus((_2482_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-179), _dafny.IntOfUint32((_2482_v7).Cardinality()))).Uint32()).(_dafny.Int)), _2475_v0)
				var _2490_v15 _dafny.CodePoint
				_ = _2490_v15
				_2490_v15 = _dafny.CodePoint('d')
				var _2491_v16 _dafny.Map
				_ = _2491_v16
				_2491_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2475_v0, _2490_v15)
				var _2492_v17 _dafny.Map
				_ = _2492_v17
				_2492_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2475_v0, (_2491_v16).Cardinality())
				var _2493_v18 D10
				_ = _2493_v18
				_2493_v18 = Companion_D10_.Create_DC21_(_2492_v17)
				var _2494_v19 _dafny.Set
				_ = _2494_v19
				_2494_v19 = _dafny.SetOf(p0, _dafny.IntOfInt64(85))
				var _2495_v22 _dafny.Array
				_ = _2495_v22
				var _nwElement0_79 _dafny.Set = (_2494_v19).Intersection(_2494_v19)
				_ = _nwElement0_79
				var _nw376 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_79, nil, _dafny.IntOfInt64(19))
				_ = _nw376
				(_nw376).ArraySet1(_nwElement0_79, 0)
				(_nw376).ArraySet1((_2494_v19).Difference(_2494_v19), 1)
				(_nw376).ArraySet1(Companion_Default___.Fm31((_2489_v14).Cardinality(), p0, p0, false, globalState), 2)
				(_nw376).ArraySet1(_2494_v19, 3)
				(_nw376).ArraySet1(_dafny.SetOf(_dafny.IntOfUint32((_2480_v5).Cardinality())), 4)
				(_nw376).ArraySet1(_2494_v19, 5)
				(_nw376).ArraySet1((_dafny.SetOf(_dafny.IntOfInt64(738), p0)).Difference(_2494_v19), 6)
				(_nw376).ArraySet1(_2494_v19, 7)
				(_nw376).ArraySet1(_2494_v19, 8)
				(_nw376).ArraySet1(_2494_v19, 9)
				(_nw376).ArraySet1(_2494_v19, 10)
				(_nw376).ArraySet1(func() _dafny.Set {
					var _coll60 = _dafny.NewBuilder()
					_ = _coll60
					for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(141), _dafny.IntOfInt64(982))); ; {
						_compr_60, _ok67 := _iter67()
						if !_ok67 {
							break
						}
						var _2496_v20 _dafny.Int
						_2496_v20 = interface{}(_compr_60).(_dafny.Int)
						if ((_dafny.IntOfInt64(141)).Cmp(_2496_v20) <= 0) && ((_2496_v20).Cmp(_dafny.IntOfInt64(982)) < 0) {
							_coll60.Add((_2496_v20).Minus(p0))
						}
					}
					return _coll60.ToSet()
				}(), 11)
				(_nw376).ArraySet1(_dafny.SetOf(_dafny.IntOfInt64(230)), 12)
				(_nw376).ArraySet1(_2494_v19, 13)
				(_nw376).ArraySet1(_2494_v19, 14)
				(_nw376).ArraySet1((_dafny.SetOf(p0)).Intersection(_2494_v19), 15)
				(_nw376).ArraySet1(func() _dafny.Set {
					var _coll61 = _dafny.NewBuilder()
					_ = _coll61
					for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-129), _dafny.IntOfInt64(732))); ; {
						_compr_61, _ok68 := _iter68()
						if !_ok68 {
							break
						}
						var _2497_v21 _dafny.Int
						_2497_v21 = interface{}(_compr_61).(_dafny.Int)
						if ((_dafny.IntOfInt64(-129)).Cmp(_2497_v21) <= 0) && ((_2497_v21).Cmp(_dafny.IntOfInt64(732)) < 0) {
							_coll61.Add(Companion_Default___.SafeDivisionInt(_2497_v21, p0))
						}
					}
					return _coll61.ToSet()
				}(), 16)
				(_nw376).ArraySet1(_2494_v19, 17)
				(_nw376).ArraySet1(_2494_v19, 18)
				_2495_v22 = _nw376
				var _index396 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_2495_v22), 0))
				_ = _index396
				(_2495_v22).ArraySet1((_2494_v19).Difference(_2494_v19), (_index396).Int())
				var _2498_v23 _dafny.Sequence
				_ = _2498_v23
				_2498_v23 = _dafny.SeqOf(_2480_v5)
				var _2499_v24 _dafny.MultiSet
				_ = _2499_v24
				_2499_v24 = _dafny.MultiSetOf(_2475_v0, false, !(_2475_v0))
				var _2500_v25 _dafny.Set
				_ = _2500_v25
				_2500_v25 = _dafny.SetOf(_2475_v0)
				var _2501_v26 D16
				_ = _2501_v26
				_2501_v26 = Companion_D16_.Create_DC33_(_2475_v0, _dafny.IntOfUint32((_2480_v5).Cardinality()), _2499_v24, p0, _2500_v25)
				var _2502_v27 _dafny.Set
				_ = _2502_v27
				_2502_v27 = _dafny.SetOf(p0, p0, p0, p0, p0)
				var _index397 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_2495_v22), 0))
				_ = _index397
				var _rhs376 bool = !(false) || (!_dafny.Companion_Sequence_.Contains((_2498_v23).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2498_v23).Cardinality()))).Uint32()).(_dafny.Sequence), _2490_v15))
				_ = _rhs376
				var _rhs377 D10 = Companion_D10_.Create_DC21_(_2492_v17)
				_ = _rhs377
				var _rhs378 bool = (_2501_v26).Dtor_cf52()
				_ = _rhs378
				var _rhs379 _dafny.Set = (_2502_v27)
				_ = _rhs379
				var _lhs239 _dafny.Array = _2495_v22
				_ = _lhs239
				var _lhs240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_2495_v22), 0))
				_ = _lhs240
				_2475_v0 = _rhs376
				_2493_v18 = _rhs377
				r1 = _rhs378
				(_lhs239).ArraySet1(_rhs379, (_lhs240).Int())
				var _2503_v28 _dafny.Int
				_ = _2503_v28
				_2503_v28 = _dafny.IntOfInt64(113)
				_2503_v28 = (_dafny.Zero).Minus(p0)
			} else {
				var _2504_v29 _dafny.Int
				_ = _2504_v29
				_2504_v29 = _dafny.IntOfInt64(868)
				_2504_v29 = (_dafny.Zero).Minus(p0)
				r0 = _2475_v0
				var _2505_v30 _dafny.MultiSet
				_ = _2505_v30
				_2505_v30 = _dafny.MultiSetOf(_2504_v29)
				var _2506_v31 _dafny.Sequence
				_ = _2506_v31
				_2506_v31 = _dafny.SeqOf(_2505_v30)
				_2504_v29 = _dafny.IntOfUint32((_2506_v31).Cardinality())
				var _2507_v32 _dafny.Array
				_ = _2507_v32
				var _nw377 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw377
				_2507_v32 = _nw377
				var _2508_v33 _dafny.Map
				_ = _2508_v33
				_2508_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _2509_v34 _dafny.Map
				_ = _2509_v34
				_2509_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2508_v33, _2475_v0)
				var _2510_v35 *C4
				_ = _2510_v35
				var _nw378 *C4 = New_C4_()
				_ = _nw378
				_nw378.Ctor__(_2507_v32, p0, (Companion_Default___.Fm67(_2475_v0, globalState)).Equals(_2509_v34))
				_2510_v35 = _nw378
				var _2511_v36 _dafny.Array
				_ = _2511_v36
				var _len0_70 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_70
				var _nw379 _dafny.Array
				_ = _nw379
				if _len0_70.Cmp(_dafny.Zero) == 0 {
					_nw379 = _dafny.NewArray(_len0_70)
				} else {
					var _init70 func(_dafny.Int) bool = (func(_2512_v0 bool) func(_dafny.Int) bool {
						return func(_2513_i0 _dafny.Int) bool {
							return (_dafny.SetOf(_2512_v0, _2512_v0)).IsSubsetOf(_dafny.SetOf(_2512_v0, _2512_v0))
						}
					})(_2475_v0)
					_ = _init70
					var _element0_70 = _init70(_dafny.Zero)
					_ = _element0_70
					_nw379 = _dafny.NewArrayFromExample(_element0_70, nil, _len0_70)
					(_nw379).ArraySet1(_element0_70, 0)
					var _nativeLen0_70 = (_len0_70).Int()
					_ = _nativeLen0_70
					for _i0_70 := 1; _i0_70 < _nativeLen0_70; _i0_70++ {
						(_nw379).ArraySet1(_init70(_dafny.IntOf(_i0_70)), _i0_70)
					}
				}
				_2511_v36 = _nw379
				var _2514_v37 D21
				_ = _2514_v37
				_2514_v37 = Companion_D21_.Create_DC44_(_2475_v0)
				var _pat_let_tv47 = _2475_v0
				_ = _pat_let_tv47
				var _index398 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(521), _dafny.ArrayLen((_2511_v36), 0))
				_ = _index398
				(_2511_v36).ArraySet1((func(_pat_let57_0 D21) D21 {
					return func(_2515_dt__update__tmp_h1 D21) D21 {
						return func(_pat_let58_0 bool) D21 {
							return func(_2516_dt__update_hcf72_h0 bool) D21 {
								return Companion_D21_.Create_DC44_(_2516_dt__update_hcf72_h0)
							}(_pat_let58_0)
						}(_pat_let_tv47)
					}(_pat_let57_0)
				}(_2514_v37)).Dtor_cf72(), (_index398).Int())
				var _index399 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(521), _dafny.ArrayLen((_2511_v36), 0))
				_ = _index399
				(_2511_v36).ArraySet1(true, (_index399).Int())
			}
			_2480_v5 = _2480_v5
			var _2517_v38 _dafny.Array
			_ = _2517_v38
			var _len0_71 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_71
			var _nw380 _dafny.Array
			_ = _nw380
			if _len0_71.Cmp(_dafny.Zero) == 0 {
				_nw380 = _dafny.NewArray(_len0_71)
			} else {
				var _init71 func(_dafny.Int) _dafny.Int = (func(_2518_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2519_i1 _dafny.Int) _dafny.Int {
						return (_2519_i1).Plus(_2518_p0)
					}
				})(p0)
				_ = _init71
				var _element0_71 = _init71(_dafny.Zero)
				_ = _element0_71
				_nw380 = _dafny.NewArrayFromExample(_element0_71, nil, _len0_71)
				(_nw380).ArraySet1(_element0_71, 0)
				var _nativeLen0_71 = (_len0_71).Int()
				_ = _nativeLen0_71
				for _i0_71 := 1; _i0_71 < _nativeLen0_71; _i0_71++ {
					(_nw380).ArraySet1(_init71(_dafny.IntOf(_i0_71)), _i0_71)
				}
			}
			_2517_v38 = _nw380
			var _index400 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_2517_v38), 0))
			_ = _index400
			(_2517_v38).ArraySet1(_dafny.IntOfInt64(647), (_index400).Int())
			var _2520_v39 _dafny.MultiSet
			_ = _2520_v39
			_2520_v39 = _dafny.MultiSetOf(!(_2475_v0))
			var _index401 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_2517_v38), 0))
			_ = _index401
			(_2517_v38).ArraySet1(((_2520_v39).Cardinality()).Times(p0), (_index401).Int())
			var _2521_v40 _dafny.Sequence
			_ = _2521_v40
			_2521_v40 = _dafny.SeqOf(_2475_v0, _2475_v0, _2475_v0, _2475_v0)
			var _2522_v41 D4
			_ = _2522_v41
			_2522_v41 = Companion_D4_.Create_DC13_(_dafny.IntOfInt64(137), _2475_v0, _dafny.IntOfUint32((_2480_v5).Cardinality()), _2521_v40)
			var _2523_v42 _dafny.MultiSet
			_ = _2523_v42
			_2523_v42 = _dafny.MultiSetOf((_2517_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_2517_v38), 0))).Int()).(_dafny.Int))
			var _index402 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_2517_v38), 0))
			_ = _index402
			(_2517_v38).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_2522_v41).Dtor_cf17(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((func() _dafny.Int {
				if (_2523_v42).Contains(p0) {
					return (_2523_v42).Multiplicity(p0)
				}
				return p0
			})()), _dafny.IntOfUint32(((_2522_v41).Dtor_cf17()).Cardinality()))).Uint32(), _2475_v0), _2521_v40), _dafny.Companion_Sequence_.Concatenate(_2521_v40, _2521_v40))).Cardinality()), (_index402).Int())
		}
		var _2524_v43 _dafny.Set
		_ = _2524_v43
		_2524_v43 = _dafny.SetOf(false)
		_2475_v0 = ((_2524_v43).Difference(_2524_v43)).IsDisjointFrom((_2524_v43).Difference(_2524_v43))
		if _2475_v0 {
			var _2525_v44 _dafny.CodePoint
			_ = _2525_v44
			_2525_v44 = _dafny.CodePoint('p')
			var _2526_v45 D4
			_ = _2526_v45
			_2526_v45 = Companion_D4_.Create_DC12_(_2525_v44)
			var _pat_let_tv48 = _2475_v0
			_ = _pat_let_tv48
			var _pat_let_tv49 = _2525_v44
			_ = _pat_let_tv49
			var _pat_let_tv50 = _2525_v44
			_ = _pat_let_tv50
			var _source40 D4 = func(_pat_let59_0 D4) D4 {
				return func(_2527_dt__update__tmp_h2 D4) D4 {
					return func(_pat_let60_0 _dafny.CodePoint) D4 {
						return func(_2528_dt__update_hcf13_h0 _dafny.CodePoint) D4 {
							return Companion_D4_.Create_DC12_(_2528_dt__update_hcf13_h0)
						}(_pat_let60_0)
					}((func() _dafny.CodePoint {
						if _pat_let_tv48 {
							return _pat_let_tv49
						}
						return _pat_let_tv50
					})())
				}(_pat_let59_0)
			}(_2526_v45)
			_ = _source40
			if _source40.Is_DC13() {
				var _2529___mcc_h0 _dafny.Int = _source40.Get_().(D4_DC13).Cf14
				_ = _2529___mcc_h0
				var _2530___mcc_h1 bool = _source40.Get_().(D4_DC13).Cf15
				_ = _2530___mcc_h1
				var _2531___mcc_h2 _dafny.Int = _source40.Get_().(D4_DC13).Cf16
				_ = _2531___mcc_h2
				var _2532___mcc_h3 _dafny.Sequence = _source40.Get_().(D4_DC13).Cf17
				_ = _2532___mcc_h3
				var _2533_cf17 _dafny.Sequence = _2532___mcc_h3
				_ = _2533_cf17
				var _2534_cf16 _dafny.Int = _2531___mcc_h2
				_ = _2534_cf16
				var _2535_cf15 bool = _2530___mcc_h1
				_ = _2535_cf15
				var _2536_cf14 _dafny.Int = _2529___mcc_h0
				_ = _2536_cf14
				_2534_cf16 = Companion_Default___.Fm3(_2536_cf14, _2535_cf15, globalState)
				var _2537_v46 *C2
				_ = _2537_v46
				var _nw381 *C2 = New_C2_()
				_ = _nw381
				_nw381.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("bgj"), _2475_v0)
				_2537_v46 = _nw381
				var _2538_v47 _dafny.Sequence
				_ = _2538_v47
				_2538_v47 = _dafny.SeqOf((_2534_cf16).Plus(p0), _dafny.IntOfInt64(-328))
				_2538_v47 = _dafny.Companion_Sequence_.Concatenate(_2538_v47, _2538_v47)
				r0 = _2535_cf15
			} else {
				var _2539___mcc_h4 _dafny.CodePoint = _source40.Get_().(D4_DC12).Cf13
				_ = _2539___mcc_h4
				var _2540_cf13 _dafny.CodePoint = _2539___mcc_h4
				_ = _2540_cf13
				var _2541_v48 _dafny.Array
				_ = _2541_v48
				var _nw382 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
				_ = _nw382
				_2541_v48 = _nw382
				var _2542_v49 *C1
				_ = _2542_v49
				var _nw383 *C1 = New_C1_()
				_ = _nw383
				_nw383.Ctor__(p0, p0, _2541_v48, p0, _2475_v0)
				_2542_v49 = _nw383
				_2542_v49 = _2542_v49
				(_2542_v49).F16 = (_2542_v49).Fm7(globalState)
				var _2543_v50 _dafny.Map
				_ = _2543_v50
				_2543_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2475_v0, _2475_v0)
				var _2544_v51 _dafny.MultiSet
				_ = _2544_v51
				_2544_v51 = _dafny.MultiSetOf(_2475_v0)
				_2543_v50 = (_2543_v50).Update(true, (_dafny.MultiSetFromSeq(_dafny.SeqOf(_2475_v0))).Equals(_2544_v51))
				var _2545_v52 _dafny.MultiSet
				_ = _2545_v52
				_2545_v52 = _dafny.MultiSetOf(_2525_v44, _2525_v44, _2540_cf13, _2525_v44)
				_2545_v52 = (func() _dafny.MultiSet {
					if _2475_v0 {
						return _2545_v52
					}
					return _dafny.MultiSetOf(_2540_cf13)
				})()
			}
			var _2546_v53 _dafny.Array
			_ = _2546_v53
			var _nw384 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(5))
			_ = _nw384
			_2546_v53 = _nw384
			var _index403 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_2546_v53), 0))
			_ = _index403
			(_2546_v53).ArraySet1CodePoint(_2525_v44, (_index403).Int())
			var _index404 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_2546_v53), 0))
			_ = _index404
			(_2546_v53).ArraySet1CodePoint(_2525_v44, (_index404).Int())
			var _2547_v54 _dafny.Int
			_ = _2547_v54
			_2547_v54 = _dafny.IntOfInt64(-581)
			_2547_v54 = _2547_v54
			var _2548_v55 _dafny.Sequence
			_ = _2548_v55
			_2548_v55 = _dafny.UnicodeSeqOfUtf8Bytes("pbuouf")
			var _2549_v56 _dafny.MultiSet
			_ = _2549_v56
			_2549_v56 = _dafny.MultiSetOf(_2548_v55, _2548_v55)
			var _2550_v57 D11
			_ = _2550_v57
			_2550_v57 = Companion_D11_.Create_DC25_(((_dafny.Zero).Minus(p0)).Cmp(p0) == 0, Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus((_2549_v56).Cardinality())), (_dafny.Zero).Minus(p0)))
			_2550_v57 = _2550_v57
			var _2551_v58 _dafny.Array
			_ = _2551_v58
			var _nwElement0_80 _dafny.Int = (_dafny.Zero).Minus(_2547_v54)
			_ = _nwElement0_80
			var _nw385 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_80, nil, _dafny.IntOfInt64(4))
			_ = _nw385
			(_nw385).ArraySet1(_nwElement0_80, 0)
			(_nw385).ArraySet1(p0, 1)
			(_nw385).ArraySet1(_2547_v54, 2)
			(_nw385).ArraySet1(_2547_v54, 3)
			_2551_v58 = _nw385
			var _2552_v59 D16
			_ = _2552_v59
			_2552_v59 = Companion_D16_.Create_DC32_(_2551_v58)
			var _source41 D16 = _2552_v59
			_ = _source41
			if _source41.Is_DC33() {
				var _2553___mcc_h5 bool = _source41.Get_().(D16_DC33).Cf52
				_ = _2553___mcc_h5
				var _2554___mcc_h6 _dafny.Int = _source41.Get_().(D16_DC33).Cf53
				_ = _2554___mcc_h6
				var _2555___mcc_h7 _dafny.MultiSet = _source41.Get_().(D16_DC33).Cf54
				_ = _2555___mcc_h7
				var _2556___mcc_h8 _dafny.Int = _source41.Get_().(D16_DC33).Cf55
				_ = _2556___mcc_h8
				var _2557___mcc_h9 _dafny.Set = _source41.Get_().(D16_DC33).Cf56
				_ = _2557___mcc_h9
				var _2558_cf56 _dafny.Set = _2557___mcc_h9
				_ = _2558_cf56
				var _2559_cf55 _dafny.Int = _2556___mcc_h8
				_ = _2559_cf55
				var _2560_cf54 _dafny.MultiSet = _2555___mcc_h7
				_ = _2560_cf54
				var _2561_cf53 _dafny.Int = _2554___mcc_h6
				_ = _2561_cf53
				var _2562_cf52 bool = _2553___mcc_h5
				_ = _2562_cf52
				r0 = !((func() bool {
					if _2562_cf52 {
						return (func() bool {
							if _2475_v0 {
								return _2475_v0
							}
							return true
						})()
					}
					return _2475_v0
				})())
				var _2563_v60 D11
				_ = _2563_v60
				_2563_v60 = Companion_D11_.Create_DC26_(_2475_v0)
				var _2564_v61 _dafny.Map
				_ = _2564_v61
				_2564_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2475_v0, _2475_v0)
				var _2565_v62 _dafny.Map
				_ = _2565_v62
				_2565_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2563_v60, (func() bool {
					if (_2564_v61).Contains(_2475_v0) {
						return (_2564_v61).Get(_2475_v0).(bool)
					}
					return _2562_cf52
				})())
				var _2566_v63 _dafny.Sequence
				_ = _2566_v63
				_2566_v63 = _dafny.SeqOf(_2565_v62)
				var _2567_v64 D18
				_ = _2567_v64
				_2567_v64 = Companion_D18_.Create_DC35_(_2566_v63)
				var _2568_v65 _dafny.Array
				_ = _2568_v65
				var _nw386 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw386
				_2568_v65 = _nw386
				var _2569_v66 _dafny.Map
				_ = _2569_v66
				_2569_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2525_v44, _2475_v0)
				var _2570_v67 _dafny.Array
				_ = _2570_v67
				var _nw387 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
				_ = _nw387
				_2570_v67 = _nw387
				var _2571_v68 T1
				_ = _2571_v68
				var _nw388 *C6 = New_C6_()
				_ = _nw388
				_nw388.Ctor__(_2475_v0, (func() bool {
					if (_2569_v66).Contains(_dafny.CodePoint('g')) {
						return (_2569_v66).Get(_dafny.CodePoint('g')).(bool)
					}
					return _2475_v0
				})(), _2570_v67, _2559_cf55)
				_2571_v68 = _nw388
				var _2572_v69 _dafny.Array
				_ = _2572_v69
				var _nwElement0_81 T1 = _2571_v68
				_ = _nwElement0_81
				var _nw389 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_81, nil, _dafny.IntOfInt64(27))
				_ = _nw389
				(_nw389).ArraySet1(_nwElement0_81, 0)
				(_nw389).ArraySet1(_2571_v68, 1)
				(_nw389).ArraySet1(_2571_v68, 2)
				(_nw389).ArraySet1(_2571_v68, 3)
				(_nw389).ArraySet1(_2571_v68, 4)
				(_nw389).ArraySet1(_2571_v68, 5)
				(_nw389).ArraySet1(_2571_v68, 6)
				(_nw389).ArraySet1(_2571_v68, 7)
				(_nw389).ArraySet1(_2571_v68, 8)
				(_nw389).ArraySet1(_2571_v68, 9)
				(_nw389).ArraySet1(_2571_v68, 10)
				(_nw389).ArraySet1(_2571_v68, 11)
				(_nw389).ArraySet1(_2571_v68, 12)
				(_nw389).ArraySet1(_2571_v68, 13)
				(_nw389).ArraySet1(_2571_v68, 14)
				(_nw389).ArraySet1(_2571_v68, 15)
				(_nw389).ArraySet1(_2571_v68, 16)
				(_nw389).ArraySet1(_2571_v68, 17)
				(_nw389).ArraySet1(_2571_v68, 18)
				(_nw389).ArraySet1(_2571_v68, 19)
				(_nw389).ArraySet1(_2571_v68, 20)
				(_nw389).ArraySet1(_2571_v68, 21)
				(_nw389).ArraySet1(_2571_v68, 22)
				(_nw389).ArraySet1(_2571_v68, 23)
				(_nw389).ArraySet1(_2571_v68, 24)
				(_nw389).ArraySet1(_2571_v68, 25)
				(_nw389).ArraySet1(_2571_v68, 26)
				_2572_v69 = _nw389
				var _2573_v70 _dafny.MultiSet
				_ = _2573_v70
				_2573_v70 = _dafny.MultiSetOf(_2572_v69, _2572_v69)
				var _2574_v71 _dafny.Map
				_ = _2574_v71
				_2574_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.MultiSet {
					if _2562_cf52 {
						return _2573_v70
					}
					return _2573_v70
				})(), _2525_v44)
				var _index405 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_2546_v53), 0))
				_ = _index405
				var _rhs380 _dafny.CodePoint = Companion_Default___.Fm1(_dafny.IntOfUint32((_2548_v55).Cardinality()), _2475_v0, globalState)
				_ = _rhs380
				var _rhs381 bool = _2562_cf52
				_ = _rhs381
				var _rhs382 D18 = _2567_v64
				_ = _rhs382
				var _rhs383 _dafny.Array = _2568_v65
				_ = _rhs383
				var _rhs384 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (_2574_v71).Contains(_dafny.MultiSetOf(_2572_v69)) {
						return (_2574_v71).Get(_dafny.MultiSetOf(_2572_v69)).(_dafny.CodePoint)
					}
					return _2525_v44
				})()
				_ = _rhs384
				var _lhs241 _dafny.Array = _2546_v53
				_ = _lhs241
				var _lhs242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_2546_v53), 0))
				_ = _lhs242
				(_lhs241).ArraySet1CodePoint(_rhs380, (_lhs242).Int())
				_2562_cf52 = _rhs381
				_2567_v64 = _rhs382
				_2568_v65 = _rhs383
				_2525_v44 = _rhs384
				var _2575_v72 _dafny.Sequence
				_ = _2575_v72
				_2575_v72 = _dafny.SeqOf(_2558_cf56)
				var _2576_v73 _dafny.Map
				_ = _2576_v73
				_2576_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_2559_cf55), _dafny.IntOfInt64(-139))
				var _2577_v74 _dafny.Sequence
				_ = _2577_v74
				_2577_v74 = _dafny.SeqOf(_2562_cf52, (_2571_v68).F4())
				var _2578_v75 D9
				_ = _2578_v75
				_2578_v75 = Companion_D9_.Create_DC19_(_2548_v55)
				var _pat_let_tv51 = _2548_v55
				_ = _pat_let_tv51
				var _2579_v76 _dafny.Array
				_ = _2579_v76
				var _nwElement0_82 bool = _2562_cf52
				_ = _nwElement0_82
				var _nw390 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_82, nil, _dafny.IntOfInt64(29))
				_ = _nw390
				(_nw390).ArraySet1(_nwElement0_82, 0)
				(_nw390).ArraySet1(_2475_v0, 1)
				(_nw390).ArraySet1(_2562_cf52, 2)
				(_nw390).ArraySet1(true, 3)
				(_nw390).ArraySet1(!(_2562_cf52), 4)
				(_nw390).ArraySet1((_2571_v68).F4(), 5)
				(_nw390).ArraySet1((_2562_cf52) && (false), 6)
				(_nw390).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(614))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg118 _dafny.Int) interface{} {
						return coer118(arg118)
					}
				}((func(_2580_v53 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_2581_i2 _dafny.Int) _dafny.CodePoint {
						return (_2580_v53).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_2580_v53), 0))).Int())
					}
				})(_2546_v53)))).Cardinality())).Cmp(_2547_v54) < 0, 7)
				(_nw390).ArraySet1(_2562_cf52, 8)
				(_nw390).ArraySet1(_2562_cf52, 9)
				(_nw390).ArraySet1(_2475_v0, 10)
				(_nw390).ArraySet1((_2547_v54).Cmp((_dafny.Zero).Minus(p0)) > 0, 11)
				(_nw390).ArraySet1(((_2575_v72).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_2576_v73).Contains((_2571_v68).F6()) {
						return (_2576_v73).Get((_2571_v68).F6()).(_dafny.Int)
					}
					return _2547_v54
				})(), _dafny.IntOfUint32((_2575_v72).Cardinality()))).Uint32()).(_dafny.Set)).Contains(!(_2475_v0)), 12)
				(_nw390).ArraySet1(_2475_v0, 13)
				(_nw390).ArraySet1(Companion_Default___.Fm0(_2525_v44, globalState), 14)
				(_nw390).ArraySet1((_2577_v74).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2577_v74).Cardinality()))).Uint32()).(bool), 15)
				(_nw390).ArraySet1(_2475_v0, 16)
				(_nw390).ArraySet1(false, 17)
				(_nw390).ArraySet1(_2562_cf52, 18)
				(_nw390).ArraySet1((_2571_v68).F4(), 19)
				(_nw390).ArraySet1(_2475_v0, 20)
				(_nw390).ArraySet1(_2562_cf52, 21)
				(_nw390).ArraySet1(_2562_cf52, 22)
				(_nw390).ArraySet1(_2475_v0, 23)
				(_nw390).ArraySet1((_2571_v68).F4(), 24)
				(_nw390).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(895))).Uint32(), func(coer119 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg119 _dafny.Int) interface{} {
						return coer119(arg119)
					}
				}((func(_2582_v53 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_2583_i3 _dafny.Int) _dafny.CodePoint {
						return (_2582_v53).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_2582_v53), 0))).Int())
					}
				})(_2546_v53))), (func(_pat_let61_0 D9) D9 {
					return func(_2584_dt__update__tmp_h3 D9) D9 {
						return func(_pat_let62_0 _dafny.Sequence) D9 {
							return func(_2585_dt__update_hcf25_h0 _dafny.Sequence) D9 {
								return Companion_D9_.Create_DC19_(_2585_dt__update_hcf25_h0)
							}(_pat_let62_0)
						}(_pat_let_tv51)
					}(_pat_let61_0)
				}(_2578_v75)).Dtor_cf25()), 25)
				(_nw390).ArraySet1(false, 26)
				(_nw390).ArraySet1((_2571_v68).F4(), 27)
				(_nw390).ArraySet1(!((_2571_v68).F4()) || (_2562_cf52), 28)
				_2579_v76 = _nw390
				var _index406 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_2579_v76), 0))
				_ = _index406
				(_2579_v76).ArraySet1(_2562_cf52, (_index406).Int())
				var _2586_v77 _dafny.Map
				_ = _2586_v77
				_2586_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2576_v73, (_2571_v68).F4())
				var _2587_v78 _dafny.Map
				_ = _2587_v78
				_2587_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2571_v68).F6(), _2551_v58)
				var _2588_v79 _dafny.Map
				_ = _2588_v79
				_2588_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3((_2586_v77).Cardinality(), _2475_v0, globalState), (func() _dafny.Array {
					if (_2587_v78).Contains(_2561_cf53) {
						return (_2587_v78).Get(_2561_cf53).(_dafny.Array)
					}
					return _2551_v58
				})())
				var _2589_v80 D20
				_ = _2589_v80
				_2589_v80 = Companion_D20_.Create_DC41_(_2547_v54)
				var _2590_v81 _dafny.MultiSet
				_ = _2590_v81
				_2590_v81 = _dafny.MultiSetOf((_2571_v68).F6())
				var _2591_v82 _dafny.Sequence
				_ = _2591_v82
				_2591_v82 = _dafny.SeqOf((_2571_v68).F6(), _2559_cf55)
				var _2592_v83 _dafny.MultiSet
				_ = _2592_v83
				_2592_v83 = _dafny.MultiSetOf((_2590_v81).Cardinality(), _dafny.IntOfUint32((_2591_v82).Cardinality()), _2547_v54, _2561_cf53, _dafny.IntOfInt64(943))
				var _2593_v84 *C1
				_ = _2593_v84
				var _nw391 *C1 = New_C1_()
				_ = _nw391
				_nw391.Ctor__(_dafny.IntOfInt64(925), p0, _2570_v67, _2561_cf53, !(_2475_v0))
				_2593_v84 = _nw391
				var _2594_v85 _dafny.Map
				_ = _2594_v85
				_2594_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(625), _dafny.SeqOf(((_2590_v81).Update(_2559_cf55, Companion_Default___.Abs((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2593_v84, (_2560_cf54).Cardinality())).Cardinality()))).Cardinality()))
				var _index407 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_2579_v76), 0))
				_ = _index407
				var _rhs385 bool = (_2571_v68).Fm8((_2589_v80).Dtor_cf67(), _2562_cf52, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(69))).Uint32(), func(coer120 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg120 _dafny.Int) interface{} {
						return coer120(arg120)
					}
				}((func(_2595_v53 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_2596_i4 _dafny.Int) _dafny.CodePoint {
						return (_2595_v53).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_2595_v53), 0))).Int())
					}
				})(_2546_v53))), (func() _dafny.Sequence {
					if (_2594_v85).Contains(_dafny.IntOfInt64(852)) {
						return (_2594_v85).Get(_dafny.IntOfInt64(852)).(_dafny.Sequence)
					}
					return _2591_v82
				})(), globalState)
				_ = _rhs385
				var _rhs386 bool = (_2571_v68).F4()
				_ = _rhs386
				var _rhs387 bool = !((_2571_v68).F4())
				_ = _rhs387
				var _rhs388 _dafny.Map = (_2588_v79).Merge(_2588_v79)
				_ = _rhs388
				var _lhs243 _dafny.Array = _2579_v76
				_ = _lhs243
				var _lhs244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_2579_v76), 0))
				_ = _lhs244
				(_lhs243).ArraySet1(_rhs385, (_lhs244).Int())
				_2562_cf52 = _rhs386
				r0 = _rhs387
				_2588_v79 = _rhs388
				var _2597_v86 _dafny.Array
				_ = _2597_v86
				var _nw392 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
				_ = _nw392
				_2597_v86 = _nw392
				var _2598_v87 _dafny.Set
				_ = _2598_v87
				_2598_v87 = _dafny.SetOf(_2579_v76, _2579_v76)
				var _2599_v88 D22
				_ = _2599_v88
				_2599_v88 = Companion_D22_.Create_DC45_(_2597_v86)
				var _2600_v89 _dafny.Map
				_ = _2600_v89
				_2600_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2598_v87, (_2599_v88).Dtor_cf73())
				_2597_v86 = (func() _dafny.Array {
					if (_2600_v89).Contains(_2598_v87) {
						return (_2600_v89).Get(_2598_v87).(_dafny.Array)
					}
					return _2597_v86
				})()
			} else {
				var _2601___mcc_h10 _dafny.Array = _source41.Get_().(D16_DC32).Cf51
				_ = _2601___mcc_h10
				var _2602_cf51 _dafny.Array = _2601___mcc_h10
				_ = _2602_cf51
				var _2603_v90 _dafny.Array
				_ = _2603_v90
				var _nw393 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw393
				_2603_v90 = _nw393
				var _2604_v91 _dafny.Sequence
				_ = _2604_v91
				_2604_v91 = _dafny.SeqOf(_2547_v54, p0, p0, _dafny.IntOfInt64(13))
				var _2605_v92 _dafny.Map
				_ = _2605_v92
				_2605_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_2603_v90, _2603_v90, _2603_v90, _2603_v90, _2603_v90), (_dafny.MultiSetFromSeq(_2604_v91)).Cardinality())
				var _2606_v93 _dafny.Map
				_ = _2606_v93
				_2606_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_2605_v92).Contains(_dafny.SeqOf(_2603_v90, _2603_v90, _2603_v90, _2603_v90)) {
						return (_2605_v92).Get(_dafny.SeqOf(_2603_v90, _2603_v90, _2603_v90, _2603_v90)).(_dafny.Int)
					}
					return _2547_v54
				})(), Companion_Default___.Fm40(false, true, p0, globalState))
				_2606_v93 = (_2606_v93).Update(_dafny.IntOfInt64(192), Companion_Default___.Fm19(_2475_v0, p0, globalState))
				var _index408 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_2603_v90), 0))
				_ = _index408
				(_2603_v90).ArraySet1((_this).Fm11(globalState), (_index408).Int())
				var _index409 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_2603_v90), 0))
				_ = _index409
				(_2603_v90).ArraySet1(_2475_v0, (_index409).Int())
				var _2607_v94 _dafny.Array
				_ = _2607_v94
				var _len0_72 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_72
				var _nw394 _dafny.Array
				_ = _nw394
				if _len0_72.Cmp(_dafny.Zero) == 0 {
					_nw394 = _dafny.NewArray(_len0_72)
				} else {
					var _init72 func(_dafny.Int) _dafny.Sequence = (func(_2608_v44 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
						return func(_2609_i5 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(370))).Uint32(), func(coer121 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg121 _dafny.Int) interface{} {
									return coer121(arg121)
								}
							}((func(_2610_v44 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_2611_i6 _dafny.Int) _dafny.CodePoint {
									return _2610_v44
								}
							})(_2608_v44)))
						}
					})(_2525_v44)
					_ = _init72
					var _element0_72 = _init72(_dafny.Zero)
					_ = _element0_72
					_nw394 = _dafny.NewArrayFromExample(_element0_72, nil, _len0_72)
					(_nw394).ArraySet1(_element0_72, 0)
					var _nativeLen0_72 = (_len0_72).Int()
					_ = _nativeLen0_72
					for _i0_72 := 1; _i0_72 < _nativeLen0_72; _i0_72++ {
						(_nw394).ArraySet1(_init72(_dafny.IntOf(_i0_72)), _i0_72)
					}
				}
				_2607_v94 = _nw394
				var _2612_v95 _dafny.Array
				_ = _2612_v95
				var _len0_73 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_73
				var _nw395 _dafny.Array
				_ = _nw395
				if _len0_73.Cmp(_dafny.Zero) == 0 {
					_nw395 = _dafny.NewArray(_len0_73)
				} else {
					var _init73 func(_dafny.Int) _dafny.Map = (func(_2613_p0 _dafny.Int, _2614_v54 _dafny.Int, _2615_v90 _dafny.Array) func(_dafny.Int) _dafny.Map {
						return func(_2616_i7 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2613_p0, (_dafny.Zero).Minus(_2614_v54))).Cardinality(), (_2615_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_2615_v90), 0))).Int()).(bool))
						}
					})(p0, _2547_v54, _2603_v90)
					_ = _init73
					var _element0_73 = _init73(_dafny.Zero)
					_ = _element0_73
					_nw395 = _dafny.NewArrayFromExample(_element0_73, nil, _len0_73)
					(_nw395).ArraySet1(_element0_73, 0)
					var _nativeLen0_73 = (_len0_73).Int()
					_ = _nativeLen0_73
					for _i0_73 := 1; _i0_73 < _nativeLen0_73; _i0_73++ {
						(_nw395).ArraySet1(_init73(_dafny.IntOf(_i0_73)), _i0_73)
					}
				}
				_2612_v95 = _nw395
				var _rhs389 _dafny.Array = _2607_v94
				_ = _rhs389
				var _rhs390 _dafny.Array = _2612_v95
				_ = _rhs390
				_2607_v94 = _rhs389
				_2612_v95 = _rhs390
				var _2617_v96 _dafny.Sequence
				_ = _2617_v96
				_2617_v96 = _dafny.SeqOf(_2548_v55, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(525))).Uint32(), func(coer122 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg122 _dafny.Int) interface{} {
						return coer122(arg122)
					}
				}((func(_2618_v44 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2619_i8 _dafny.Int) _dafny.CodePoint {
						return _2618_v44
					}
				})(_2525_v44))), _2548_v55, _2548_v55)
				var _2620_v97 D12
				_ = _2620_v97
				_2620_v97 = Companion_D12_.Create_DC27_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(195))).Uint32(), func(coer123 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg123 _dafny.Int) interface{} {
						return coer123(arg123)
					}
				}((func(_2621_v0 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_2622_i9 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_2621_v0)
					}
				})(_2475_v0))))
				var _2623_v98 _dafny.Set
				_ = _2623_v98
				_2623_v98 = _dafny.SetOf(_2547_v54, p0)
				var _2624_v99 _dafny.Map
				_ = _2624_v99
				_2624_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_2603_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_2603_v90), 0))).Int()).(bool)), _2623_v98)
				var _index410 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_2603_v90), 0))
				_ = _index410
				var _rhs391 bool = (p0).Cmp(p0) >= 0
				_ = _rhs391
				var _rhs392 _dafny.Sequence = Companion_Default___.Fm39(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_2548_v55, (Companion_Default___.SafeIndex(_2547_v54, _dafny.IntOfUint32((_2548_v55).Cardinality()))).Uint32(), (_2546_v53).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_2546_v53), 0))).Int())), (_2617_v96).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2617_v96).Cardinality()))).Uint32()).(_dafny.Sequence)), _2620_v97, _2475_v0, globalState)
				_ = _rhs392
				var _rhs393 bool = (_this).Fm11(globalState)
				_ = _rhs393
				var _rhs394 _dafny.Int = (((func() _dafny.Set {
					if (_2624_v99).Contains(false) {
						return (_2624_v99).Get(false).(_dafny.Set)
					}
					return _2623_v98
				})()).Cardinality()).Plus(_2547_v54)
				_ = _rhs394
				var _rhs395 _dafny.Int = _2547_v54
				_ = _rhs395
				var _lhs245 _dafny.Array = _2603_v90
				_ = _lhs245
				var _lhs246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_2603_v90), 0))
				_ = _lhs246
				r1 = _rhs391
				_2604_v91 = _rhs392
				(_lhs245).ArraySet1(_rhs393, (_lhs246).Int())
				_2547_v54 = _rhs394
				_2547_v54 = _rhs395
			}
		} else {
			var _2625_v100 _dafny.MultiSet
			_ = _2625_v100
			_2625_v100 = _dafny.MultiSetOf(_2475_v0)
			r0 = ((_2625_v100).Difference(_2625_v100)).IsDisjointFrom(_dafny.MultiSetOf(_2475_v0, _2475_v0, _2475_v0))
			var _2626_v101 _dafny.Array
			_ = _2626_v101
			var _len0_74 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_74
			var _nw396 _dafny.Array
			_ = _nw396
			if _len0_74.Cmp(_dafny.Zero) == 0 {
				_nw396 = _dafny.NewArray(_len0_74)
			} else {
				var _init74 func(_dafny.Int) bool = (func(_2627_v0 bool) func(_dafny.Int) bool {
					return func(_2628_i10 _dafny.Int) bool {
						return (_2627_v0) == (_2627_v0)
					}
				})(_2475_v0)
				_ = _init74
				var _element0_74 = _init74(_dafny.Zero)
				_ = _element0_74
				_nw396 = _dafny.NewArrayFromExample(_element0_74, nil, _len0_74)
				(_nw396).ArraySet1(_element0_74, 0)
				var _nativeLen0_74 = (_len0_74).Int()
				_ = _nativeLen0_74
				for _i0_74 := 1; _i0_74 < _nativeLen0_74; _i0_74++ {
					(_nw396).ArraySet1(_init74(_dafny.IntOf(_i0_74)), _i0_74)
				}
			}
			_2626_v101 = _nw396
			var _index411 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_2626_v101), 0))
			_ = _index411
			(_2626_v101).ArraySet1(_2475_v0, (_index411).Int())
			var _2629_v102 _dafny.Map
			_ = _2629_v102
			_2629_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2475_v0, p0)
			var _index412 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_2626_v101), 0))
			_ = _index412
			(_2626_v101).ArraySet1(((func() _dafny.Int {
				if (_2629_v102).Contains(_2475_v0) {
					return (_2629_v102).Get(_2475_v0).(_dafny.Int)
				}
				return p0
			})()).Cmp(Companion_Default___.Fm3((_2524_v43).Cardinality(), _2475_v0, globalState)) == 0, (_index412).Int())
			var _2630_v103 _dafny.Array
			_ = _2630_v103
			var _len0_75 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_75
			var _nw397 _dafny.Array
			_ = _nw397
			if _len0_75.Cmp(_dafny.Zero) == 0 {
				_nw397 = _dafny.NewArray(_len0_75)
			} else {
				var _init75 func(_dafny.Int) _dafny.Int = (func(_2631_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2632_i11 _dafny.Int) _dafny.Int {
						return (_2632_i11).Minus(_2631_p0)
					}
				})(p0)
				_ = _init75
				var _element0_75 = _init75(_dafny.Zero)
				_ = _element0_75
				_nw397 = _dafny.NewArrayFromExample(_element0_75, nil, _len0_75)
				(_nw397).ArraySet1(_element0_75, 0)
				var _nativeLen0_75 = (_len0_75).Int()
				_ = _nativeLen0_75
				for _i0_75 := 1; _i0_75 < _nativeLen0_75; _i0_75++ {
					(_nw397).ArraySet1(_init75(_dafny.IntOf(_i0_75)), _i0_75)
				}
			}
			_2630_v103 = _nw397
			var _index413 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2630_v103), 0))
			_ = _index413
			(_2630_v103).ArraySet1((_dafny.Zero).Minus(p0), (_index413).Int())
			var _index414 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2630_v103), 0))
			_ = _index414
			(_2630_v103).ArraySet1(p0, (_index414).Int())
			var _2633_v104 _dafny.Array
			_ = _2633_v104
			var _nw398 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
			_ = _nw398
			_2633_v104 = _nw398
			var _2634_v105 D22
			_ = _2634_v105
			_2634_v105 = Companion_D22_.Create_DC45_(_2633_v104)
			var _pat_let_tv52 = _2634_v105
			_ = _pat_let_tv52
			var _2635_v106 _dafny.Array
			_ = _2635_v106
			var _nwElement0_83 D22 = _2634_v105
			_ = _nwElement0_83
			var _nw399 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_83, nil, _dafny.IntOfInt64(4))
			_ = _nw399
			(_nw399).ArraySet1(_nwElement0_83, 0)
			(_nw399).ArraySet1(_2634_v105, 1)
			(_nw399).ArraySet1(func(_pat_let63_0 D22) D22 {
				return func(_2636_dt__update__tmp_h4 D22) D22 {
					return func(_pat_let64_0 _dafny.Array) D22 {
						return func(_2637_dt__update_hcf73_h0 _dafny.Array) D22 {
							return Companion_D22_.Create_DC45_(_2637_dt__update_hcf73_h0)
						}(_pat_let64_0)
					}((_pat_let_tv52).Dtor_cf73())
				}(_pat_let63_0)
			}(_2634_v105), 2)
			(_nw399).ArraySet1(_2634_v105, 3)
			_2635_v106 = _nw399
			var _index415 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_2635_v106), 0))
			_ = _index415
			(_2635_v106).ArraySet1(_2634_v105, (_index415).Int())
			var _index416 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_2635_v106), 0))
			_ = _index416
			(_2635_v106).ArraySet1(_2634_v105, (_index416).Int())
			var _2638_v107 *C0
			_ = _2638_v107
			var _nw400 *C0 = New_C0_()
			_ = _nw400
			_nw400.Ctor__((_2630_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2630_v103), 0))).Int()).(_dafny.Int))
			_2638_v107 = _nw400
			var _2639_v108 _dafny.Sequence
			_ = _2639_v108
			_2639_v108 = _dafny.SeqOf(_2638_v107, _2638_v107)
			var _2640_v109 _dafny.Map
			_ = _2640_v109
			_2640_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2475_v0, p0)).Cardinality(), _2638_v107)
			var _2641_v110 D23
			_ = _2641_v110
			_2641_v110 = Companion_D23_.Create_DC48_(_2638_v107)
			var _2642_v111 _dafny.Array
			_ = _2642_v111
			var _nwElement0_84 *C0 = _2638_v107
			_ = _nwElement0_84
			var _nw401 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_84, nil, _dafny.IntOfInt64(24))
			_ = _nw401
			(_nw401).ArraySet1(_nwElement0_84, 0)
			(_nw401).ArraySet1(_2638_v107, 1)
			(_nw401).ArraySet1(_2638_v107, 2)
			(_nw401).ArraySet1(_2638_v107, 3)
			(_nw401).ArraySet1(_2638_v107, 4)
			(_nw401).ArraySet1(_2638_v107, 5)
			(_nw401).ArraySet1(_2638_v107, 6)
			(_nw401).ArraySet1(_2638_v107, 7)
			(_nw401).ArraySet1(_2638_v107, 8)
			(_nw401).ArraySet1(_2638_v107, 9)
			(_nw401).ArraySet1(_2638_v107, 10)
			(_nw401).ArraySet1(_2638_v107, 11)
			(_nw401).ArraySet1((_2639_v108).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_2638_v107).F12(), p0)).Cardinality()), _dafny.IntOfUint32((_2639_v108).Cardinality()))).Uint32()).(*C0), 12)
			(_nw401).ArraySet1(_2638_v107, 13)
			(_nw401).ArraySet1(_2638_v107, 14)
			(_nw401).ArraySet1((func() *C0 {
				if (_2640_v109).Contains(p0) {
					return (_2640_v109).Get(p0).(*C0)
				}
				return _2638_v107
			})(), 15)
			(_nw401).ArraySet1(_2638_v107, 16)
			(_nw401).ArraySet1((_2641_v110).Dtor_cf82(), 17)
			(_nw401).ArraySet1(_2638_v107, 18)
			(_nw401).ArraySet1(_2638_v107, 19)
			(_nw401).ArraySet1(_2638_v107, 20)
			(_nw401).ArraySet1(_2638_v107, 21)
			(_nw401).ArraySet1(_2638_v107, 22)
			(_nw401).ArraySet1((_2641_v110).Dtor_cf82(), 23)
			_2642_v111 = _nw401
			_2642_v111 = _2642_v111
		}
		var _2643_v112 _dafny.Array
		_ = _2643_v112
		var _len0_76 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_76
		var _nw402 _dafny.Array
		_ = _nw402
		if _len0_76.Cmp(_dafny.Zero) == 0 {
			_nw402 = _dafny.NewArray(_len0_76)
		} else {
			var _init76 func(_dafny.Int) _dafny.Int = (func(_2644_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2645_i13 _dafny.Int) _dafny.Int {
					return (_2645_i13).Times(_2644_p0)
				}
			})(p0)
			_ = _init76
			var _element0_76 = _init76(_dafny.Zero)
			_ = _element0_76
			_nw402 = _dafny.NewArrayFromExample(_element0_76, nil, _len0_76)
			(_nw402).ArraySet1(_element0_76, 0)
			var _nativeLen0_76 = (_len0_76).Int()
			_ = _nativeLen0_76
			for _i0_76 := 1; _i0_76 < _nativeLen0_76; _i0_76++ {
				(_nw402).ArraySet1(_init76(_dafny.IntOf(_i0_76)), _i0_76)
			}
		}
		_2643_v112 = _nw402
		for _iter69 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2643_v112), 0))); ; {
			_guard_loop_7, _ok69 := _iter69()
			if !_ok69 {
				break
			}
			var _2646_i12 _dafny.Int
			_2646_i12 = interface{}(_guard_loop_7).(_dafny.Int)
			if (true) && (((_2646_i12).Sign() != -1) && ((_2646_i12).Cmp(_dafny.ArrayLen((_2643_v112), 0)) < 0)) {
				(_2643_v112).ArraySet1(Companion_Default___.SafeModuloInt(_2646_i12, (p0).Plus(p0)), (_2646_i12).Int())
			}
		}
		var _2647_v113 _dafny.Array
		_ = _2647_v113
		var _nw403 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
		_ = _nw403
		_2647_v113 = _nw403
		var _2648_v114 _dafny.Array
		_ = _2648_v114
		var _len0_77 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_77
		var _nw404 _dafny.Array
		_ = _nw404
		if _len0_77.Cmp(_dafny.Zero) == 0 {
			_nw404 = _dafny.NewArray(_len0_77)
		} else {
			var _init77 func(_dafny.Int) bool = (func(_2649_v0 bool) func(_dafny.Int) bool {
				return func(_2650_i14 _dafny.Int) bool {
					return _2649_v0
				}
			})(_2475_v0)
			_ = _init77
			var _element0_77 = _init77(_dafny.Zero)
			_ = _element0_77
			_nw404 = _dafny.NewArrayFromExample(_element0_77, nil, _len0_77)
			(_nw404).ArraySet1(_element0_77, 0)
			var _nativeLen0_77 = (_len0_77).Int()
			_ = _nativeLen0_77
			for _i0_77 := 1; _i0_77 < _nativeLen0_77; _i0_77++ {
				(_nw404).ArraySet1(_init77(_dafny.IntOf(_i0_77)), _i0_77)
			}
		}
		_2648_v114 = _nw404
		var _2651_v115 _dafny.Array
		_ = _2651_v115
		var _nwElement0_85 _dafny.Array = _2648_v114
		_ = _nwElement0_85
		var _nw405 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_85, nil, _dafny.IntOfInt64(23))
		_ = _nw405
		(_nw405).ArraySet1(_nwElement0_85, 0)
		(_nw405).ArraySet1(_2648_v114, 1)
		(_nw405).ArraySet1(_2648_v114, 2)
		(_nw405).ArraySet1(_2648_v114, 3)
		(_nw405).ArraySet1(_2648_v114, 4)
		(_nw405).ArraySet1(_2648_v114, 5)
		(_nw405).ArraySet1(_2648_v114, 6)
		(_nw405).ArraySet1(_2648_v114, 7)
		(_nw405).ArraySet1(_2648_v114, 8)
		(_nw405).ArraySet1(_2648_v114, 9)
		(_nw405).ArraySet1(_2648_v114, 10)
		(_nw405).ArraySet1(_2648_v114, 11)
		(_nw405).ArraySet1(_2648_v114, 12)
		(_nw405).ArraySet1(_2648_v114, 13)
		(_nw405).ArraySet1(_2648_v114, 14)
		(_nw405).ArraySet1(_2648_v114, 15)
		(_nw405).ArraySet1(_2648_v114, 16)
		(_nw405).ArraySet1(_2648_v114, 17)
		(_nw405).ArraySet1(_2648_v114, 18)
		(_nw405).ArraySet1(_2648_v114, 19)
		(_nw405).ArraySet1(_2648_v114, 20)
		(_nw405).ArraySet1(_2648_v114, 21)
		(_nw405).ArraySet1(_2648_v114, 22)
		_2651_v115 = _nw405
		var _2652_v116 _dafny.Sequence
		_ = _2652_v116
		_2652_v116 = _dafny.UnicodeSeqOfUtf8Bytes("clpbitvi")
		var _2653_v117 _dafny.Sequence
		_ = _2653_v117
		_2653_v117 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("vvtigi"), _2652_v116)
		var _2654_v118 *C3
		_ = _2654_v118
		var _nw406 *C3 = New_C3_()
		_ = _nw406
		_nw406.Ctor__(_2651_v115, _dafny.IntOfUint32((_2653_v117).Cardinality()), _2475_v0)
		_2654_v118 = _nw406
		var _index417 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_2647_v113), 0))
		_ = _index417
		(_2647_v113).ArraySet1(_2654_v118, (_index417).Int())
		var _2655_v119 _dafny.Map
		_ = _2655_v119
		_2655_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2475_v0, p0)
		var _2656_v120 *C0
		_ = _2656_v120
		var _nw407 *C0 = New_C0_()
		_ = _nw407
		_nw407.Ctor__((_dafny.Zero).Minus((_2655_v119).Cardinality()))
		_2656_v120 = _nw407
		var _index418 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_2647_v113), 0))
		_ = _index418
		var _rhs396 bool = _2475_v0
		_ = _rhs396
		var _rhs397 *C3 = _2654_v118
		_ = _rhs397
		var _rhs398 *C0 = _2656_v120
		_ = _rhs398
		var _lhs247 _dafny.Array = _2647_v113
		_ = _lhs247
		var _lhs248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_2647_v113), 0))
		_ = _lhs248
		r1 = _rhs396
		(_lhs247).ArraySet1(_rhs397, (_lhs248).Int())
		_2656_v120 = _rhs398
		var _2657_v121 _dafny.Int
		_ = _2657_v121
		_2657_v121 = _dafny.IntOfInt64(-191)
		_2657_v121 = ((_dafny.Zero).Minus((_2656_v120).F12())).Minus((_2656_v120).F12())
		var _2658_v122 D22
		_ = _2658_v122
		_2658_v122 = Companion_D22_.Create_DC47_(_dafny.IntOfUint32((_2652_v116).Cardinality()), !(!(_2475_v0)), _2475_v0, (_2656_v120).F12(), (_dafny.Zero).Minus(p0))
		r0 = (_2658_v122).Dtor_cf79()
		r1 = _2475_v0
		return r0, r1
	}
}

// End of class C8

// Definition of class C9
type C9 struct {
	_f5 _dafny.Array
	_f6 _dafny.Int
	_f4 bool
	F7  _dafny.Int
	_f8 _dafny.Set
}

func New_C9_() *C9 {
	_this := C9{}

	_this._f5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = _dafny.Zero
	_this._f4 = false
	_this.F7 = _dafny.Zero
	_this._f8 = _dafny.EmptySet
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C9{}
var _ T0 = &C9{}
var _ _dafny.TraitOffspring = &C9{}

func (_this *C9) F5() _dafny.Array {
	return _this._f5
}
func (_this *C9) F5_set_(value _dafny.Array) {
	_this._f5 = value
}
func (_this *C9) F6() _dafny.Int {
	return _this._f6
}
func (_this *C9) F4() bool {
	return _this._f4
}
func (_this *C9) Ctor__(f7 _dafny.Int, f8 _dafny.Set, f5 _dafny.Array, f6 _dafny.Int, f4 bool) {
	{
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f4 = f4
	}
}
func (_this *C9) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F4()
	}
}
func (_this *C9) Fm9(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var _source42 D1 = (func() D1 {
			if (_this).F4() {
				return Companion_D1_.Create_DC4_(_dafny.IntOfUint32((_dafny.SeqOf(_this.F7)).Cardinality()), _dafny.IntOfInt64(912))
			}
			return Companion_D1_.Create_DC4_((_this).F6(), _dafny.IntOfInt64(392))
		})()
		_ = _source42
		if _source42.Is_DC4() {
			var _2659___mcc_h0 _dafny.Int = _source42.Get_().(D1_DC4).Cf3
			_ = _2659___mcc_h0
			var _2660___mcc_h1 _dafny.Int = _source42.Get_().(D1_DC4).Cf4
			_ = _2660___mcc_h1
			var _2661_cf4 _dafny.Int = _2660___mcc_h1
			_ = _2661_cf4
			var _2662_cf3 _dafny.Int = _2659___mcc_h0
			_ = _2662_cf3
			return (_dafny.Zero).Minus(_2662_cf3)
		} else {
			var _2663___mcc_h2 bool = _source42.Get_().(D1_DC3).Cf2
			_ = _2663___mcc_h2
			var _2664_cf2 bool = _2663___mcc_h2
			_ = _2664_cf2
			return (_this).F6()
		}
	}
}
func (_this *C9) Fm6(p0 D0, p1 D0, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F4()))).Cardinality(), (_this).F6(), _this.F7, _dafny.IntOfInt64(-280))).Union((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fku")).Cardinality()))).Intersection(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, (_this).F4())).Cardinality())))
	}
}
func (_this *C9) Fm7(globalState *GlobalState) _dafny.Int {
	{
		return _this.F7
	}
}
func (_this *C9) M1(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _2665_v0 _dafny.Array
		_ = _2665_v0
		var _len0_78 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_78
		var _nw408 _dafny.Array
		_ = _nw408
		if _len0_78.Cmp(_dafny.Zero) == 0 {
			_nw408 = _dafny.NewArray(_len0_78)
		} else {
			var _init78 func(_dafny.Int) bool = func(_2666_i1 _dafny.Int) bool {
				return (_this).F4()
			}
			_ = _init78
			var _element0_78 = _init78(_dafny.Zero)
			_ = _element0_78
			_nw408 = _dafny.NewArrayFromExample(_element0_78, nil, _len0_78)
			(_nw408).ArraySet1(_element0_78, 0)
			var _nativeLen0_78 = (_len0_78).Int()
			_ = _nativeLen0_78
			for _i0_78 := 1; _i0_78 < _nativeLen0_78; _i0_78++ {
				(_nw408).ArraySet1(_init78(_dafny.IntOf(_i0_78)), _i0_78)
			}
		}
		_2665_v0 = _nw408
		for _iter70 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2665_v0), 0))); ; {
			_guard_loop_8, _ok70 := _iter70()
			if !_ok70 {
				break
			}
			var _2667_i0 _dafny.Int
			_2667_i0 = interface{}(_guard_loop_8).(_dafny.Int)
			if (true) && (((_2667_i0).Sign() != -1) && ((_2667_i0).Cmp(_dafny.ArrayLen((_2665_v0), 0)) < 0)) {
				(_2665_v0).ArraySet1(((_this).F6()).Cmp(p0) >= 0, (_2667_i0).Int())
			}
		}
		var _hi14 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_this.F7, _this.F7))
		_ = _hi14
		for _2668_i2 := Companion_Default___.SafeModuloInt(_this.F7, _this.F7); _2668_i2.Cmp(_hi14) < 0; _2668_i2 = _2668_i2.Plus(_dafny.One) {
			var _2669_v1 D1
			_ = _2669_v1
			_2669_v1 = Companion_D1_.Create_DC4_((_this).F6(), p0)
			var _2670_v2 _dafny.Map
			_ = _2670_v2
			_2670_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F4()), (func() D1 {
				if false {
					return _2669_v1
				}
				return _2669_v1
			})())
			var _2671_v3 _dafny.Sequence
			_ = _2671_v3
			_2671_v3 = _dafny.SeqOf(_dafny.CodePoint('f'))
			var _2672_v4 _dafny.Map
			_ = _2672_v4
			_2672_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _2671_v3)
			var _2673_v5 _dafny.CodePoint
			_ = _2673_v5
			_2673_v5 = _dafny.CodePoint('h')
			var _2674_v6 _dafny.Map
			_ = _2674_v6
			_2674_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2671_v3, (Companion_Default___.SafeIndex(_2668_i2, _dafny.IntOfUint32((_2671_v3).Cardinality()))).Uint32(), _2673_v5)).Cardinality()))
			_2670_v2 = (_2670_v2).Update((_this).F4(), Companion_Default___.Fm10((_this).F4(), false, _2672_v4, (_2674_v6).Cardinality(), globalState))
			var _2675_v7 bool
			_ = _2675_v7
			_2675_v7 = true
			_2675_v7 = (_this).F4()
			var _2676_v8 *C0
			_ = _2676_v8
			var _nw409 *C0 = New_C0_()
			_ = _nw409
			_nw409.Ctor__((_2668_i2).Plus((_dafny.Zero).Minus((_this).F6())))
			_2676_v8 = _nw409
			if _2675_v7 {
				var _2677_v9 _dafny.MultiSet
				_ = _2677_v9
				_2677_v9 = _dafny.MultiSetOf((_2676_v8).F12(), (_dafny.MultiSetOf((_2674_v6).Cardinality(), _dafny.IntOfInt64(886))).Cardinality())
				var _2678_v10 _dafny.Map
				_ = _2678_v10
				_2678_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2677_v9).Intersection(_2677_v9), _dafny.IntOfInt64(451))
				var _2679_v11 D0
				_ = _2679_v11
				_2679_v11 = Companion_D0_.Create_DC0_(_2668_i2)
				var _2680_v12 _dafny.Map
				_ = _2680_v12
				_2680_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _2673_v5)
				_2678_v10 = (_2678_v10).Update((_dafny.MultiSetOf(p0, (_dafny.Zero).Minus((_2676_v8).F12()), _this.F7)).Union((_this).Fm6(_2679_v11, _2679_v11, _2680_v12, _dafny.IntOfInt64(251), globalState)), _this.F7)
				var _rhs399 bool = (_this).F4()
				_ = _rhs399
				var _rhs400 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(_2668_i2))).Cardinality())
				_ = _rhs400
				var _lhs249 *C9 = _this
				_ = _lhs249
				_2675_v7 = _rhs399
				_lhs249.F7 = _rhs400
				_2675_v7 = _2675_v7
				var _2681_v13 _dafny.Array
				_ = _2681_v13
				var _len0_79 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_79
				var _nw410 _dafny.Array
				_ = _nw410
				if _len0_79.Cmp(_dafny.Zero) == 0 {
					_nw410 = _dafny.NewArray(_len0_79)
				} else {
					var _init79 func(_dafny.Int) _dafny.Sequence = (func(_2682_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_2683_i3 _dafny.Int) _dafny.Sequence {
							return _2682_v3
						}
					})(_2671_v3)
					_ = _init79
					var _element0_79 = _init79(_dafny.Zero)
					_ = _element0_79
					_nw410 = _dafny.NewArrayFromExample(_element0_79, nil, _len0_79)
					(_nw410).ArraySet1(_element0_79, 0)
					var _nativeLen0_79 = (_len0_79).Int()
					_ = _nativeLen0_79
					for _i0_79 := 1; _i0_79 < _nativeLen0_79; _i0_79++ {
						(_nw410).ArraySet1(_init79(_dafny.IntOf(_i0_79)), _i0_79)
					}
				}
				_2681_v13 = _nw410
				var _index419 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_2681_v13), 0))
				_ = _index419
				(_2681_v13).ArraySet1(Companion_Default___.Fm16(_2675_v7, (_this).F4(), !(_2675_v7), Companion_D2_.Create_DC6_(), globalState), (_index419).Int())
				var _2684_v14 _dafny.MultiSet
				_ = _2684_v14
				_2684_v14 = _dafny.MultiSetOf(_2679_v11, _2679_v11)
				var _2685_v15 _dafny.Array
				_ = _2685_v15
				var _nwElement0_86 _dafny.MultiSet = _2684_v14
				_ = _nwElement0_86
				var _nw411 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_86, nil, _dafny.IntOfInt64(2))
				_ = _nw411
				(_nw411).ArraySet1(_nwElement0_86, 0)
				(_nw411).ArraySet1(_2684_v14, 1)
				_2685_v15 = _nw411
				var _index420 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_2685_v15), 0))
				_ = _index420
				(_2685_v15).ArraySet1(_dafny.MultiSetOf(_2679_v11), (_index420).Int())
				var _2686_v16 _dafny.Sequence
				_ = _2686_v16
				_2686_v16 = _dafny.SeqOf(func(_pat_let65_0 D0) D0 {
					return func(_2687_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let66_0 _dafny.Int) D0 {
							return func(_2688_dt__update_hcf0_h0 _dafny.Int) D0 {
								return Companion_D0_.Create_DC0_(_2688_dt__update_hcf0_h0)
							}(_pat_let66_0)
						}((_dafny.Zero).Minus(_2668_i2))
					}(_pat_let65_0)
				}(_2679_v11))
				var _index421 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_2681_v13), 0))
				_ = _index421
				var _index422 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_2685_v15), 0))
				_ = _index422
				var _rhs401 _dafny.Sequence = _2671_v3
				_ = _rhs401
				var _rhs402 _dafny.MultiSet = (_2684_v14).Difference(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_2686_v16, _2686_v16)))
				_ = _rhs402
				var _lhs250 _dafny.Array = _2681_v13
				_ = _lhs250
				var _lhs251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_2681_v13), 0))
				_ = _lhs251
				var _lhs252 _dafny.Array = _2685_v15
				_ = _lhs252
				var _lhs253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_2685_v15), 0))
				_ = _lhs253
				(_lhs250).ArraySet1(_rhs401, (_lhs251).Int())
				(_lhs252).ArraySet1(_rhs402, (_lhs253).Int())
				var _2689_v17 D12
				_ = _2689_v17
				_2689_v17 = Companion_D12_.Create_DC28_(p0, _this.F7, (_this).F4(), true, p0)
				var _2690_v18 _dafny.Sequence
				_ = _2690_v18
				_2690_v18 = _dafny.SeqOf((_2689_v17).Dtor_cf46(), (_this).F4(), (_2675_v7) == (_2675_v7))
				var _2691_v19 _dafny.Sequence
				_ = _2691_v19
				_2691_v19 = _dafny.SeqOf((_this).F6())
				_2690_v18 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_2690_v18, (Companion_Default___.SafeIndex(_2668_i2, _dafny.IntOfUint32((_2690_v18).Cardinality()))).Uint32(), (_2690_v18).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2691_v19).Cardinality()), _dafny.IntOfUint32((_2690_v18).Cardinality()))).Uint32()).(bool)), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_this).F6(), _this.F7), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2690_v18, (Companion_Default___.SafeIndex(_2668_i2, _dafny.IntOfUint32((_2690_v18).Cardinality()))).Uint32(), (_2690_v18).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2691_v19).Cardinality()), _dafny.IntOfUint32((_2690_v18).Cardinality()))).Uint32()).(bool))).Cardinality()))).Uint32(), !(true))
			} else {
				var _2692_v20 _dafny.Map
				_ = _2692_v20
				_2692_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2671_v3, Companion_Default___.Fm1(_this.F7, _2675_v7, globalState))
				_2692_v20 = (_2692_v20).Update(_dafny.UnicodeSeqOfUtf8Bytes("yrkp"), _2673_v5)
				var _2693_v21 *C6
				_ = _2693_v21
				var _nw412 *C6 = New_C6_()
				_ = _nw412
				_nw412.Ctor__((_this).F4(), (_this).F4(), (func() _dafny.Array {
					if _2675_v7 {
						return _this.F5()
					}
					return _this.F5()
				})(), (_2676_v8).F12())
				_2693_v21 = _nw412
				var _2694_v22 _dafny.Set
				_ = _2694_v22
				_2694_v22 = _dafny.SetOf(!(_2675_v7))
				_2675_v7 = !(_2694_v22).Equals(Companion_Default___.Fm46(_dafny.CodePoint('t'), globalState))
				var _2695_v23 D3
				_ = _2695_v23
				_2695_v23 = Companion_D3_.Create_DC9_(true, p0, (_this).F4())
				var _2696_v24 _dafny.Map
				_ = _2696_v24
				_2696_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F6()), (_2695_v23).Dtor_cf9())
				var _2697_v25 _dafny.Sequence
				_ = _2697_v25
				_2697_v25 = _dafny.SeqOf(_2675_v7, (_this).F4())
				var _2698_v26 _dafny.Sequence
				_ = _2698_v26
				_2698_v26 = _dafny.SeqOf(_dafny.IntOfUint32((_2697_v25).Cardinality()), (_this).F6())
				var _2699_v27 D6
				_ = _2699_v27
				_2699_v27 = Companion_D6_.Create_DC16_(false, (_2693_v21).F11(), _2698_v26)
				var _2700_v28 _dafny.Set
				_ = _2700_v28
				_2700_v28 = _dafny.SetOf(_2699_v27, _2699_v27, _2699_v27)
				var _rhs403 _dafny.Map = (Companion_Default___.Fm42(_2673_v5, _2700_v28, globalState)).Merge(_2696_v24)
				_ = _rhs403
				var _rhs404 bool = (_2668_i2).Cmp((_2676_v8).F12()) < 0
				_ = _rhs404
				var _rhs405 bool = ((_2676_v8).F12()).Cmp((_this).Fm7(globalState)) != 0
				_ = _rhs405
				var _rhs406 _dafny.Int = (p0).Minus(_dafny.IntOfInt64(345))
				_ = _rhs406
				var _rhs407 _dafny.Int = p0
				_ = _rhs407
				var _lhs254 *C9 = _this
				_ = _lhs254
				var _lhs255 *C9 = _this
				_ = _lhs255
				_2696_v24 = _rhs403
				_2675_v7 = _rhs404
				_2675_v7 = _rhs405
				_lhs254.F7 = _rhs406
				_lhs255.F7 = _rhs407
				_2675_v7 = ((_2676_v8).F12()).Cmp(((_this).F6()).Minus(_this.F7)) != 0
			}
		}
		var _2701_i4 _dafny.Int
		_ = _2701_i4
		_2701_i4 = _dafny.Zero
		{
			for (_this).F4() {
				{
					if (_2701_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L13
					}
					_2701_i4 = (_2701_i4).Plus(_dafny.One)
					(_this).F7 = (_dafny.Zero).Minus((_this).F6())
					var _2702_v29 _dafny.Sequence
					_ = _2702_v29
					_2702_v29 = _dafny.UnicodeSeqOfUtf8Bytes("cnnryc")
					var _2703_v30 T0
					_ = _2703_v30
					var _nw413 *C2 = New_C2_()
					_ = _nw413
					_nw413.Ctor__(_2702_v29, (_this).F4())
					_2703_v30 = _nw413
					var _2704_v31 _dafny.Sequence
					_ = _2704_v31
					_2704_v31 = _dafny.SeqOf(_2703_v30, _2703_v30, _2703_v30)
					var _2705_v32 T0
					_ = _2705_v32
					_2705_v32 = _2703_v30
					var _2706_v33 _dafny.Array
					_ = _2706_v33
					var _nwElement0_87 T0 = _2703_v30
					_ = _nwElement0_87
					var _nw414 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_87, nil, _dafny.IntOfInt64(18))
					_ = _nw414
					(_nw414).ArraySet1(_nwElement0_87, 0)
					(_nw414).ArraySet1((_2704_v31).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2704_v31).Cardinality()))).Uint32()).(T0), 1)
					(_nw414).ArraySet1((_2704_v31).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2702_v29).Cardinality()), _dafny.IntOfUint32((_2704_v31).Cardinality()))).Uint32()).(T0), 2)
					(_nw414).ArraySet1((func() T0 {
						if (_this).F4() {
							return _2703_v30
						}
						return _2703_v30
					})(), 3)
					(_nw414).ArraySet1((_2705_v32), 4)
					(_nw414).ArraySet1(_2703_v30, 5)
					(_nw414).ArraySet1((func() T0 {
						if (_this).F4() {
							return _2703_v30
						}
						return _2703_v30
					})(), 6)
					(_nw414).ArraySet1(_2703_v30, 7)
					(_nw414).ArraySet1(_2703_v30, 8)
					(_nw414).ArraySet1(_2703_v30, 9)
					(_nw414).ArraySet1(_2703_v30, 10)
					(_nw414).ArraySet1(_2703_v30, 11)
					(_nw414).ArraySet1(_2703_v30, 12)
					(_nw414).ArraySet1(_2703_v30, 13)
					(_nw414).ArraySet1(_2703_v30, 14)
					(_nw414).ArraySet1(_2703_v30, 15)
					(_nw414).ArraySet1(_2703_v30, 16)
					(_nw414).ArraySet1(_2703_v30, 17)
					_2706_v33 = _nw414
					var _index423 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_2706_v33), 0))
					_ = _index423
					(_2706_v33).ArraySet1(_2703_v30, (_index423).Int())
					var _index424 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_2706_v33), 0))
					_ = _index424
					(_2706_v33).ArraySet1(_2703_v30, (_index424).Int())
					var _source43 D12 = Companion_Default___.Fm68(globalState)
					_ = _source43
					if _source43.Is_DC28() {
						var _2707___mcc_h0 _dafny.Int = _source43.Get_().(D12_DC28).Cf43
						_ = _2707___mcc_h0
						var _2708___mcc_h1 _dafny.Int = _source43.Get_().(D12_DC28).Cf44
						_ = _2708___mcc_h1
						var _2709___mcc_h2 bool = _source43.Get_().(D12_DC28).Cf45
						_ = _2709___mcc_h2
						var _2710___mcc_h3 bool = _source43.Get_().(D12_DC28).Cf46
						_ = _2710___mcc_h3
						var _2711___mcc_h4 _dafny.Int = _source43.Get_().(D12_DC28).Cf47
						_ = _2711___mcc_h4
						var _2712_cf47 _dafny.Int = _2711___mcc_h4
						_ = _2712_cf47
						var _2713_cf46 bool = _2710___mcc_h3
						_ = _2713_cf46
						var _2714_cf45 bool = _2709___mcc_h2
						_ = _2714_cf45
						var _2715_cf44 _dafny.Int = _2708___mcc_h1
						_ = _2715_cf44
						var _2716_cf43 _dafny.Int = _2707___mcc_h0
						_ = _2716_cf43
						_2716_cf43 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2702_v29, _2702_v29)).Cardinality())
						var _2717_v34 _dafny.Map
						_ = _2717_v34
						_2717_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2703_v30).F4(), !(_2713_cf46))
						_2717_v34 = _2717_v34
						var _2718_v35 _dafny.Array
						_ = _2718_v35
						var _len0_80 _dafny.Int = _dafny.IntOfInt64(28)
						_ = _len0_80
						var _nw415 _dafny.Array
						_ = _nw415
						if _len0_80.Cmp(_dafny.Zero) == 0 {
							_nw415 = _dafny.NewArray(_len0_80)
						} else {
							var _init80 func(_dafny.Int) _dafny.Set = (func(_2719_cf46 bool, _2720_v30 T0, _2721_cf45 bool) func(_dafny.Int) _dafny.Set {
								return func(_2722_i5 _dafny.Int) _dafny.Set {
									return _dafny.SetOf((_dafny.SetOf(_2719_cf46, (_2720_v30).F4(), !((_2720_v30).F4()), (_this).F4(), _2721_cf45)).Cardinality())
								}
							})(_2713_cf46, _2703_v30, _2714_cf45)
							_ = _init80
							var _element0_80 = _init80(_dafny.Zero)
							_ = _element0_80
							_nw415 = _dafny.NewArrayFromExample(_element0_80, nil, _len0_80)
							(_nw415).ArraySet1(_element0_80, 0)
							var _nativeLen0_80 = (_len0_80).Int()
							_ = _nativeLen0_80
							for _i0_80 := 1; _i0_80 < _nativeLen0_80; _i0_80++ {
								(_nw415).ArraySet1(_init80(_dafny.IntOf(_i0_80)), _i0_80)
							}
						}
						_2718_v35 = _nw415
						var _index425 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2718_v35), 0))
						_ = _index425
						(_2718_v35).ArraySet1(_dafny.SetOf(_2715_cf44), (_index425).Int())
						var _index426 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2718_v35), 0))
						_ = _index426
						(_2718_v35).ArraySet1((_dafny.SetOf(p0, p0, _dafny.IntOfUint32((_dafny.SeqOf(_2713_cf46)).Cardinality()), _2716_cf43, (_this).Fm7(globalState))).Intersection(_dafny.SetOf((_this).F6())), (_index426).Int())
						var _rhs408 bool = _2714_cf45
						_ = _rhs408
						var _rhs409 bool = ((_2718_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2718_v35), 0))).Int()).(_dafny.Set)).IsProperSubsetOf(((_2718_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2718_v35), 0))).Int()).(_dafny.Set)).Intersection((_2718_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2718_v35), 0))).Int()).(_dafny.Set)))
						_ = _rhs409
						_2714_cf45 = _rhs408
						_2713_cf46 = _rhs409
					} else {
						var _2723___mcc_h5 _dafny.Sequence = _source43.Get_().(D12_DC27).Cf42
						_ = _2723___mcc_h5
						var _2724_cf42 _dafny.Sequence = _2723___mcc_h5
						_ = _2724_cf42
						var _2725_v36 _dafny.Sequence
						_ = _2725_v36
						_2725_v36 = _dafny.SeqOf((_this).F4(), (_2703_v30).F4(), (_2703_v30).F4())
						_2725_v36 = _dafny.Companion_Sequence_.Update(_2725_v36, (Companion_Default___.SafeIndex(_this.F7, _dafny.IntOfUint32((_2725_v36).Cardinality()))).Uint32(), (_this).F4())
						var _2726_v37 bool
						_ = _2726_v37
						_2726_v37 = false
						_2726_v37 = _2726_v37
						var _index427 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_2665_v0), 0))
						_ = _index427
						(_2665_v0).ArraySet1(_2726_v37, (_index427).Int())
						var _2727_v38 _dafny.MultiSet
						_ = _2727_v38
						_2727_v38 = _dafny.MultiSetOf((_this).F4(), (_2703_v30).F4(), true)
						var _index428 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_2665_v0), 0))
						_ = _index428
						(_2665_v0).ArraySet1((_dafny.MultiSetOf(_2726_v37)).IsSubsetOf(_2727_v38), (_index428).Int())
						(_this).F7 = p0
					}
					var _2728_v39 _dafny.Map
					_ = _2728_v39
					_2728_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2703_v30, (_2703_v30).F4())
					var _2729_v40 D12
					_ = _2729_v40
					_2729_v40 = Companion_D12_.Create_DC28_(_this.F7, _dafny.IntOfUint32((_2702_v29).Cardinality()), (_2703_v30).F4(), true, _dafny.IntOfUint32((_2702_v29).Cardinality()))
					var _2730_v41 _dafny.Set
					_ = _2730_v41
					_2730_v41 = _dafny.SetOf((func() bool {
						if (_2728_v39).Contains(Companion_T0_.CastTo_((_2706_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_2706_v33), 0))).Int()))) {
							return (_2728_v39).Get(Companion_T0_.CastTo_((_2706_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_2706_v33), 0))).Int()))).(bool)
						}
						return (_2703_v30).F4()
					})(), false, (_2729_v40).Dtor_cf45(), (_this).F4(), (_this).F4())
					var _2731_v42 _dafny.Map
					_ = _2731_v42
					_2731_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2703_v30).F4(), (_2703_v30).F4())
					var _2732_v43 _dafny.Sequence
					_ = _2732_v43
					_2732_v43 = _dafny.SeqOf(_2731_v42, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_2703_v30).F4()))
					var _2733_v44 _dafny.Sequence
					_ = _2733_v44
					_2733_v44 = _dafny.SeqOf(_2730_v41, Companion_Default___.Fm53((_this).F6(), _2732_v43, globalState))
					(_this).F7 = Companion_Default___.SafeDivisionInt(((_2733_v44).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2733_v44).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), (_dafny.IntOfInt64(308)).Plus((_this).F6()))
					goto C13
				}
			C13:
			}
			goto L13
		}
	L13:
		var _2734_v45 _dafny.Array
		_ = _2734_v45
		_2734_v45 = _this.F5()
		var _2735_v46 bool
		_ = _2735_v46
		_2735_v46 = false
		var _2736_v47 _dafny.Set
		_ = _2736_v47
		_2736_v47 = _dafny.SetOf(_this.F7, _this.F7)
		var _2737_v48 _dafny.Sequence
		_ = _2737_v48
		_2737_v48 = _dafny.SeqOf(false, (_2736_v47).IsProperSubsetOf(_2736_v47), _2735_v46, (_2735_v46) == (_2735_v46))
		var _2738_v49 _dafny.MultiSet
		_ = _2738_v49
		_2738_v49 = _dafny.MultiSetOf(p0, (_this).F6())
		var _rhs410 _dafny.Array = _2734_v45
		_ = _rhs410
		var _rhs411 _dafny.Int = _dafny.IntOfUint32((_2737_v48).Cardinality())
		_ = _rhs411
		var _rhs412 bool = (_2738_v49).Equals((_2738_v49).Difference(_2738_v49))
		_ = _rhs412
		var _lhs256 *C9 = _this
		_ = _lhs256
		_2734_v45 = _rhs410
		_lhs256.F7 = _rhs411
		_2735_v46 = _rhs412
		var _2739_v50 *C0
		_ = _2739_v50
		var _nw416 *C0 = New_C0_()
		_ = _nw416
		_nw416.Ctor__(p0)
		_2739_v50 = _nw416
		var _2740_v51 _dafny.CodePoint
		_ = _2740_v51
		_2740_v51 = _dafny.CodePoint('d')
		var _2741_v52 _dafny.Sequence
		_ = _2741_v52
		_2741_v52 = _dafny.UnicodeSeqOfUtf8Bytes("nom")
		var _2742_v53 D3
		_ = _2742_v53
		_2742_v53 = Companion_D3_.Create_DC9_((_this).F4(), p0, true)
		var _index429 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_2665_v0), 0))
		_ = _index429
		(_2665_v0).ArraySet1((_2742_v53).Dtor_cf10(), (_index429).Int())
		var _2743_v54 _dafny.MultiSet
		_ = _2743_v54
		_2743_v54 = _dafny.MultiSetOf(_2738_v49)
		var _2744_v55 _dafny.Map
		_ = _2744_v55
		_2744_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2740_v51, _2743_v54)
		var _index430 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_2665_v0), 0))
		_ = _index430
		var _rhs413 _dafny.Int = (func() _dafny.Int {
			if _2735_v46 {
				return _this.F7
			}
			return _dafny.IntOfInt64(-737)
		})()
		_ = _rhs413
		var _rhs414 _dafny.CodePoint = _2740_v51
		_ = _rhs414
		var _rhs415 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("dxypm")
		_ = _rhs415
		var _rhs416 bool = ((func() _dafny.MultiSet {
			if (_2744_v55).Contains(_2740_v51) {
				return (_2744_v55).Get(_2740_v51).(_dafny.MultiSet)
			}
			return _2743_v54
		})()).IsProperSubsetOf(_2743_v54)
		_ = _rhs416
		var _rhs417 bool = _2735_v46
		_ = _rhs417
		var _lhs257 *C9 = _this
		_ = _lhs257
		var _lhs258 _dafny.Array = _2665_v0
		_ = _lhs258
		var _lhs259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_2665_v0), 0))
		_ = _lhs259
		_lhs257.F7 = _rhs413
		_2740_v51 = _rhs414
		_2741_v52 = _rhs415
		_2735_v46 = _rhs416
		(_lhs258).ArraySet1(_rhs417, (_lhs259).Int())
		var _2745_v56 _dafny.Map
		_ = _2745_v56
		_2745_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2665_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_2665_v0), 0))).Int()).(bool), (_2739_v50).F12())
		r0 = (_2745_v56).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_2665_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_2665_v0), 0))).Int()).(bool)), (_2739_v50).F12()))
		return r0
	}
}
func (_this *C9) M2(p0 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Array) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _2746_v0 *C4
		_ = _2746_v0
		var _nw417 *C4 = New_C4_()
		_ = _nw417
		_nw417.Ctor__(_this.F5(), Companion_Default___.SafeDivisionInt((_this).Fm7(globalState), _this.F7), !(p0))
		_2746_v0 = _nw417
		var _2747_i0 _dafny.Int
		_ = _2747_i0
		_2747_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_2747_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L14
					}
					_2747_i0 = (_2747_i0).Plus(_dafny.One)
					var _2748_v1 *C8
					_ = _2748_v1
					var _nw418 *C8 = New_C8_()
					_ = _nw418
					_nw418.Ctor__()
					_2748_v1 = _nw418
					var _2749_v2 _dafny.Sequence
					_ = _2749_v2
					_2749_v2 = _dafny.UnicodeSeqOfUtf8Bytes("bj")
					var _2750_v3 _dafny.Map
					_ = _2750_v3
					_2750_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, _2749_v2)
					_2750_v3 = Companion_Default___.Fm69(globalState)
					var _2751_v4 _dafny.Array
					_ = _2751_v4
					var _len0_81 _dafny.Int = _dafny.IntOfInt64(22)
					_ = _len0_81
					var _nw419 _dafny.Array
					_ = _nw419
					if _len0_81.Cmp(_dafny.Zero) == 0 {
						_nw419 = _dafny.NewArray(_len0_81)
					} else {
						var _init81 func(_dafny.Int) _dafny.Int = func(_2752_i1 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_2752_i1, _dafny.IntOfUint32((_dafny.SeqOf((_this).F4(), false)).Cardinality()))
						}
						_ = _init81
						var _element0_81 = _init81(_dafny.Zero)
						_ = _element0_81
						_nw419 = _dafny.NewArrayFromExample(_element0_81, nil, _len0_81)
						(_nw419).ArraySet1(_element0_81, 0)
						var _nativeLen0_81 = (_len0_81).Int()
						_ = _nativeLen0_81
						for _i0_81 := 1; _i0_81 < _nativeLen0_81; _i0_81++ {
							(_nw419).ArraySet1(_init81(_dafny.IntOf(_i0_81)), _i0_81)
						}
					}
					_2751_v4 = _nw419
					_2751_v4 = _2751_v4
					(_this).F7 = (_this).F6()
					goto C14
				}
			C14:
			}
			goto L14
		}
	L14:
		var _2753_v5 _dafny.Array
		_ = _2753_v5
		var _nw420 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
		_ = _nw420
		_2753_v5 = _nw420
		_2753_v5 = _2753_v5
		if p0 {
			var _index431 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_2753_v5), 0))
			_ = _index431
			(_2753_v5).ArraySet1((_this).F6(), (_index431).Int())
			var _index432 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_2753_v5), 0))
			_ = _index432
			(_2753_v5).ArraySet1(Companion_Default___.SafeModuloInt((_this).F6(), (_this).F6()), (_index432).Int())
			var _2754_v6 bool
			_ = _2754_v6
			_2754_v6 = true
			var _2755_v7 _dafny.Set
			_ = _2755_v7
			_2755_v7 = _dafny.SetOf(p0)
			var _2756_v8 _dafny.CodePoint
			_ = _2756_v8
			_2756_v8 = _dafny.CodePoint('w')
			_2754_v6 = ((Companion_Default___.Fm46(_2756_v8, globalState)).Intersection(_2755_v7)).IsProperSubsetOf((_2755_v7).Intersection(_2755_v7))
			var _2757_v9 *C3
			_ = _2757_v9
			var _nw421 *C3 = New_C3_()
			_ = _nw421
			_nw421.Ctor__(_this.F5(), (_2753_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_2753_v5), 0))).Int()).(_dafny.Int), _2754_v6)
			_2757_v9 = _nw421
			var _2758_v10 _dafny.Array
			_ = _2758_v10
			var _nw422 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
			_ = _nw422
			_2758_v10 = _nw422
			var _index433 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_2758_v10), 0))
			_ = _index433
			(_2758_v10).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(767))).Uint32(), func(coer124 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg124 _dafny.Int) interface{} {
					return coer124(arg124)
				}
			}((func(_2759_v5 _dafny.Array) func(_dafny.Int) _dafny.Int {
				return func(_2760_i2 _dafny.Int) _dafny.Int {
					return (_2759_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_2759_v5), 0))).Int()).(_dafny.Int)
				}
			})(_2753_v5))), (_index433).Int())
			var _2761_v11 _dafny.Sequence
			_ = _2761_v11
			_2761_v11 = _dafny.SeqOf(_this.F7)
			var _2762_v12 _dafny.Sequence
			_ = _2762_v12
			_2762_v12 = _dafny.UnicodeSeqOfUtf8Bytes("xondmryd")
			var _2763_v13 _dafny.Sequence
			_ = _2763_v13
			_2763_v13 = _dafny.SeqOf(p0, (_this).F4(), (_this).F4())
			var _2764_v14 D12
			_ = _2764_v14
			_2764_v14 = Companion_D12_.Create_DC27_(_dafny.SeqOf(_dafny.Companion_Sequence_.Update(_2763_v13, (Companion_Default___.SafeIndex((_2753_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_2753_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2763_v13).Cardinality()))).Uint32(), _2754_v6)))
			var _2765_v15 _dafny.Sequence
			_ = _2765_v15
			_2765_v15 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_2761_v11, _dafny.SeqOf((_this).F6())), _dafny.Companion_Sequence_.Concatenate(_2761_v11, Companion_Default___.Fm39(_2762_v12, _2764_v14, (_this).F4(), globalState)))
			var _2766_v16 _dafny.MultiSet
			_ = _2766_v16
			_2766_v16 = _dafny.MultiSetOf((_this).F6(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cprbau")).Cardinality()))
			var _index434 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_2758_v10), 0))
			_ = _index434
			(_2758_v10).ArraySet1((_2765_v15).Select((Companion_Default___.SafeIndex((_2766_v16).Cardinality(), _dafny.IntOfUint32((_2765_v15).Cardinality()))).Uint32()).(_dafny.Sequence), (_index434).Int())
			var _2767_v17 _dafny.MultiSet
			_ = _2767_v17
			_2767_v17 = _dafny.MultiSetOf(_2762_v12)
			var _2768_v18 _dafny.Map
			_ = _2768_v18
			_2768_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_2753_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_2753_v5), 0))).Int()).(_dafny.Int))
			_2767_v17 = (_2767_v17).Difference((_2767_v17).Update(_2762_v12, Companion_Default___.Abs((_2768_v18).Cardinality())))
		} else {
			var _2769_v19 _dafny.Sequence
			_ = _2769_v19
			_2769_v19 = _dafny.SeqOf((_this).F6(), Companion_Default___.SafeModuloInt(_this.F7, _this.F7), _this.F7)
			_2769_v19 = _dafny.Companion_Sequence_.Concatenate(_2769_v19, _dafny.Companion_Sequence_.Concatenate(_2769_v19, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(484))).Uint32(), func(coer125 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg125 _dafny.Int) interface{} {
					return coer125(arg125)
				}
			}(func(_2770_i3 _dafny.Int) _dafny.Int {
				return (_this).F6()
			}))))
			var _2771_v20 _dafny.CodePoint
			_ = _2771_v20
			_2771_v20 = _dafny.CodePoint('f')
			var _2772_v21 *C5
			_ = _2772_v21
			var _nw423 *C5 = New_C5_()
			_ = _nw423
			_nw423.Ctor__(!(!(p0)), _this.F7, _this.F5(), (_this).F6(), Companion_Default___.Fm0(_2771_v20, globalState))
			_2772_v21 = _nw423
			var _2773_v22 bool
			_ = _2773_v22
			_2773_v22 = true
			var _2774_v23 _dafny.Set
			_ = _2774_v23
			_2774_v23 = _dafny.SetOf(p0)
			var _index435 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_2753_v5), 0))
			_ = _index435
			(_2753_v5).ArraySet1(Companion_Default___.SafeDivisionInt(_2772_v21.F14, (_dafny.Zero).Minus((_2774_v23).Cardinality())), (_index435).Int())
			var _2775_v24 _dafny.Set
			_ = _2775_v24
			_2775_v24 = _dafny.SetOf((_this).F6(), _dafny.IntOfInt64(18))
			var _2776_v25 _dafny.MultiSet
			_ = _2776_v25
			_2776_v25 = _dafny.MultiSetOf((_2772_v21).F13(), (_2772_v21).F13())
			var _2777_v26 _dafny.Map
			_ = _2777_v26
			_2777_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2772_v21).F13(), _dafny.MultiSetOf(p0))
			var _index436 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_2753_v5), 0))
			_ = _index436
			var _rhs418 bool = (_2775_v24).IsSubsetOf((_2775_v24).Difference(_2775_v24))
			_ = _rhs418
			var _rhs419 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F6(), ((_this).F6()).Times(_this.F7))
			_ = _rhs419
			var _rhs420 bool = (_dafny.MultiSetOf(false, p0)).IsSubsetOf(_2776_v25)
			_ = _rhs420
			var _rhs421 _dafny.Int = (((_this).F6()).Times(_this.F7)).Times(((_this).F6()).Minus((_2772_v21).Fm7(globalState)))
			_ = _rhs421
			var _rhs422 bool = !(_2777_v26).Contains(!((_2772_v21).F13()) || (p0))
			_ = _rhs422
			var _lhs260 *C5 = _2772_v21
			_ = _lhs260
			var _lhs261 _dafny.Array = _2753_v5
			_ = _lhs261
			var _lhs262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_2753_v5), 0))
			_ = _lhs262
			_2773_v22 = _rhs418
			_lhs260.F14 = _rhs419
			_2773_v22 = _rhs420
			(_lhs261).ArraySet1(_rhs421, (_lhs262).Int())
			_2773_v22 = _rhs422
			_2773_v22 = p0
			var _index437 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_2753_v5), 0))
			_ = _index437
			(_2753_v5).ArraySet1(_dafny.IntOfInt64(45), (_index437).Int())
		}
		var _2778_v27 _dafny.Sequence
		_ = _2778_v27
		_2778_v27 = _dafny.UnicodeSeqOfUtf8Bytes("packlnnxk")
		_2778_v27 = _2778_v27
		var _2779_v28 _dafny.Sequence
		_ = _2779_v28
		_2779_v28 = _dafny.SeqOf((_this).F4())
		var _2780_v29 _dafny.CodePoint
		_ = _2780_v29
		_2780_v29 = _dafny.CodePoint('y')
		var _2781_v30 _dafny.MultiSet
		_ = _2781_v30
		_2781_v30 = _dafny.MultiSetOf(_2780_v29)
		var _2782_v31 _dafny.Sequence
		_ = _2782_v31
		_2782_v31 = _dafny.SeqOf(_2781_v30, _2781_v30)
		if _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm70((_this).F4(), _dafny.IntOfInt64(736), _2779_v28, p0, globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2781_v30), _2782_v31)) {
			(_this).F7 = _this.F7
			var _2783_v32 _dafny.Map
			_ = _2783_v32
			_2783_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, (_this).F4())
			var _rhs423 _dafny.Int = _this.F7
			_ = _rhs423
			var _rhs424 _dafny.Int = (func() _dafny.Int {
				if !((func() bool {
					if (_2783_v32).Contains(_dafny.IntOfUint32((_2778_v27).Cardinality())) {
						return (_2783_v32).Get(_dafny.IntOfUint32((_2778_v27).Cardinality())).(bool)
					}
					return (_this).F4()
				})()) || (p0) {
					return _this.F7
				}
				return (_this).F6()
			})()
			_ = _rhs424
			var _rhs425 _dafny.Int = _this.F7
			_ = _rhs425
			var _rhs426 _dafny.Sequence = _2779_v28
			_ = _rhs426
			var _lhs263 *C9 = _this
			_ = _lhs263
			var _lhs264 *C9 = _this
			_ = _lhs264
			var _lhs265 *C9 = _this
			_ = _lhs265
			_lhs263.F7 = _rhs423
			_lhs264.F7 = _rhs424
			_lhs265.F7 = _rhs425
			_2779_v28 = _rhs426
			var _2784_v33 bool
			_ = _2784_v33
			_2784_v33 = true
			_2784_v33 = !(_2784_v33)
			var _2785_v34 _dafny.Sequence
			_ = _2785_v34
			_2785_v34 = _dafny.SeqOf(_this.F7)
			var _index438 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(227), _dafny.ArrayLen((_2753_v5), 0))
			_ = _index438
			(_2753_v5).ArraySet1((_dafny.IntOfUint32((_2785_v34).Cardinality())).Times((_this).F6()), (_index438).Int())
			var _index439 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_2753_v5), 0))
			_ = _index439
			(_2753_v5).ArraySet1((func() _dafny.Int {
				if (_this).Fm8(_this.F7, _2784_v33, _2778_v27, _dafny.Companion_Sequence_.Update(_2785_v34, (Companion_Default___.SafeIndex(_this.F7, _dafny.IntOfUint32((_2785_v34).Cardinality()))).Uint32(), _this.F7), globalState) {
					return _dafny.IntOfUint32((_2779_v28).Cardinality())
				}
				return Companion_Default___.Fm3((_this).F6(), false, globalState)
			})(), (_index439).Int())
			var _2786_v35 _dafny.Map
			_ = _2786_v35
			_2786_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
			var _2787_v36 _dafny.MultiSet
			_ = _2787_v36
			_2787_v36 = _dafny.MultiSetOf((func() bool {
				if (_2786_v35).Contains(_2784_v33) {
					return (_2786_v35).Get(_2784_v33).(bool)
				}
				return (_2779_v28).Select((Companion_Default___.SafeIndex(_this.F7, _dafny.IntOfUint32((_2779_v28).Cardinality()))).Uint32()).(bool)
			})(), _2784_v33, !(_2784_v33))
			var _index440 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(227), _dafny.ArrayLen((_2753_v5), 0))
			_ = _index440
			var _index441 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_2753_v5), 0))
			_ = _index441
			var _rhs427 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F6(), (func() _dafny.Int {
				if (_2787_v36).Contains((_2746_v0).Fm8(_this.F7, false, _2778_v27, _2785_v34, globalState)) {
					return (_2787_v36).Multiplicity((_2746_v0).Fm8(_this.F7, false, _2778_v27, _2785_v34, globalState))
				}
				return _dafny.IntOfUint32((_dafny.SeqOf(p0, p0)).Cardinality())
			})())))
			_ = _rhs427
			var _rhs428 _dafny.Int = _this.F7
			_ = _rhs428
			var _rhs429 _dafny.Int = Companion_Default___.Fm3((_this).F6(), _2784_v33, globalState)
			_ = _rhs429
			var _rhs430 bool = (_2779_v28).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2779_v28).Cardinality()))).Uint32()).(bool)
			_ = _rhs430
			var _rhs431 bool = p0
			_ = _rhs431
			var _lhs266 _dafny.Array = _2753_v5
			_ = _lhs266
			var _lhs267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(227), _dafny.ArrayLen((_2753_v5), 0))
			_ = _lhs267
			var _lhs268 _dafny.Array = _2753_v5
			_ = _lhs268
			var _lhs269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_2753_v5), 0))
			_ = _lhs269
			var _lhs270 *C9 = _this
			_ = _lhs270
			(_lhs266).ArraySet1(_rhs427, (_lhs267).Int())
			(_lhs268).ArraySet1(_rhs428, (_lhs269).Int())
			_lhs270.F7 = _rhs429
			_2784_v33 = _rhs430
			_2784_v33 = _rhs431
			var _2788_v37 _dafny.Sequence
			_ = _2788_v37
			_2788_v37 = _dafny.SeqOf(_2778_v27)
			_2784_v33 = (_2746_v0).Fm8((_this).F6(), (_this).F4(), (_2788_v37).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2788_v37).Cardinality()))).Uint32()).(_dafny.Sequence), _2785_v34, globalState)
		} else {
			var _2789_v38 _dafny.Array
			_ = _2789_v38
			var _nw424 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
			_ = _nw424
			_2789_v38 = _nw424
			var _2790_v39 D2
			_ = _2790_v39
			_2790_v39 = Companion_D2_.Create_DC5_(_2789_v38)
			_2790_v39 = _2790_v39
			_2789_v38 = _2789_v38
			var _2791_v40 _dafny.Array
			_ = _2791_v40
			var _nw425 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
			_ = _nw425
			_2791_v40 = _nw425
			var _2792_v41 _dafny.Sequence
			_ = _2792_v41
			_2792_v41 = _dafny.SeqOf(_2753_v5, _2753_v5)
			var _2793_v42 _dafny.Map
			_ = _2793_v42
			_2793_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2792_v41).Select((Companion_Default___.SafeIndex(_this.F7, _dafny.IntOfUint32((_2792_v41).Cardinality()))).Uint32()).(_dafny.Array), _2791_v40)
			_2791_v40 = (func() _dafny.Array {
				if (_2793_v42).Contains((_2792_v41).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2792_v41).Cardinality()))).Uint32()).(_dafny.Array)) {
					return (_2793_v42).Get((_2792_v41).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_2792_v41).Cardinality()))).Uint32()).(_dafny.Array)).(_dafny.Array)
				}
				return _2791_v40
			})()
			var _2794_v43 _dafny.Map
			_ = _2794_v43
			_2794_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F6()), _this.F7)
			var _2795_v44 _dafny.Sequence
			_ = _2795_v44
			_2795_v44 = _dafny.SeqOf((_this).F6())
			(_this).F7 = Companion_Default___.SafeDivisionInt((_dafny.IntOfInt64(480)).Plus((func() _dafny.Int {
				if (_2794_v43).Contains((_2795_v44).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.IntOfUint32((_2795_v44).Cardinality()))).Uint32()).(_dafny.Int)) {
					return (_2794_v43).Get((_2795_v44).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.IntOfUint32((_2795_v44).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_2779_v28).Cardinality())
			})()), _dafny.IntOfInt64(444))
			var _2796_v45 _dafny.Map
			_ = _2796_v45
			_2796_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2753_v5, _2789_v38)
			(_this).F7 = (_2796_v45).Cardinality()
		}
		r0 = _2780_v29
		var _nw426 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(9))
		_ = _nw426
		r1 = _nw426
		return r0, r1
	}
}
func (_this *C9) M3(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _2797_v0 _dafny.Sequence
		_ = _2797_v0
		_2797_v0 = _dafny.SeqOf(true)
		var _2798_v1 _dafny.Sequence
		_ = _2798_v1
		_2798_v1 = _dafny.SeqOf(_2797_v0)
		var _2799_v2 D12
		_ = _2799_v2
		_2799_v2 = Companion_D12_.Create_DC27_(_2798_v1)
		var _2800_v3 _dafny.Sequence
		_ = _2800_v3
		_2800_v3 = _dafny.SeqOf((Companion_Default___.Fm71(globalState)).Cardinality(), _dafny.IntOfInt64(953))
		var _2801_v4 D6
		_ = _2801_v4
		_2801_v4 = Companion_D6_.Create_DC16_((_this).F4(), !((_this).F4()) || ((_this).F4()), _dafny.Companion_Sequence_.Concatenate(_2800_v3, _2800_v3))
		var _2802_v5 _dafny.Sequence
		_ = _2802_v5
		_2802_v5 = _dafny.UnicodeSeqOfUtf8Bytes("wweol")
		var _2803_v6 _dafny.Map
		_ = _2803_v6
		_2803_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F4())
		var _2804_v7 _dafny.Set
		_ = _2804_v7
		_2804_v7 = _dafny.SetOf(false)
		var _2805_v8 *C5
		_ = _2805_v8
		var _nw427 *C5 = New_C5_()
		_ = _nw427
		_nw427.Ctor__((_this).F4(), (_2804_v7).Cardinality(), _this.F5(), (_this).F6(), !(Companion_Default___.Fm0(_dafny.CodePoint('s'), globalState)))
		_2805_v8 = _nw427
		var _2806_v9 _dafny.Sequence
		_ = _2806_v9
		_2806_v9 = _dafny.SeqOf(_2805_v8)
		var _2807_v10 _dafny.Set
		_ = _2807_v10
		_2807_v10 = _dafny.SetOf(_2805_v8, (_2806_v9).Select((Companion_Default___.SafeIndex(_this.F7, _dafny.IntOfUint32((_2806_v9).Cardinality()))).Uint32()).(*C5))
		var _2808_v11 _dafny.MultiSet
		_ = _2808_v11
		_2808_v11 = _dafny.MultiSetOf(!((_2797_v0).Select((Companion_Default___.SafeIndex(_2805_v8.F14, _dafny.IntOfUint32((_2797_v0).Cardinality()))).Uint32()).(bool)), (_this).F4(), (_this).F4())
		var _2809_v12 _dafny.Array
		_ = _2809_v12
		var _nwElement0_88 _dafny.Int = (_this).F6()
		_ = _nwElement0_88
		var _nw428 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_88, nil, _dafny.IntOfInt64(15))
		_ = _nw428
		(_nw428).ArraySet1(_nwElement0_88, 0)
		(_nw428).ArraySet1(((_2803_v6).Merge(_2803_v6)).Cardinality(), 1)
		(_nw428).ArraySet1((_dafny.SetOf((_this).F4(), (_this).F4(), (_this).F4(), (_this).F4())).Cardinality(), 2)
		(_nw428).ArraySet1(p0, 3)
		(_nw428).ArraySet1(Companion_Default___.SafeModuloInt(p0, (_2807_v10).Cardinality()), 4)
		(_nw428).ArraySet1((_this).F6(), 5)
		(_nw428).ArraySet1(p0, 6)
		(_nw428).ArraySet1((_2805_v8).Fm9((_2805_v8).F13(), (_2805_v8).F13(), p1, globalState), 7)
		(_nw428).ArraySet1(p1, 8)
		(_nw428).ArraySet1((_2805_v8).Fm7(globalState), 9)
		(_nw428).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
			if (_2808_v11).Contains(false) {
				return (_2808_v11).Multiplicity(false)
			}
			return p0
		})()), 10)
		(_nw428).ArraySet1((p0).Minus(_this.F7), 11)
		(_nw428).ArraySet1((_dafny.IntOfUint32((_2800_v3).Cardinality())).Plus(p0), 12)
		(_nw428).ArraySet1((_2804_v7).Cardinality(), 13)
		(_nw428).ArraySet1(p0, 14)
		_2809_v12 = _nw428
		var _index442 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))
		_ = _index442
		(_2809_v12).ArraySet1((Companion_D4_.Create_DC13_(_dafny.IntOfUint32((_2802_v5).Cardinality()), (_this).F4(), _this.F7, _2797_v0)).Dtor_cf14(), (_index442).Int())
		var _2810_v13 _dafny.CodePoint
		_ = _2810_v13
		_2810_v13 = _dafny.CodePoint('r')
		var _2811_v14 D9
		_ = _2811_v14
		_2811_v14 = Companion_D9_.Create_DC19_(_dafny.UnicodeSeqOfUtf8Bytes("dks"))
		var _index443 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))
		_ = _index443
		var _rhs432 D12 = _2799_v2
		_ = _rhs432
		var _rhs433 D6 = Companion_D6_.Create_DC16_((_2805_v8).F13(), Companion_Default___.Fm0(Companion_Default___.Fm1(_dafny.IntOfUint32((_2802_v5).Cardinality()), (_this).F4(), globalState), globalState), _dafny.Companion_Sequence_.Update(_2800_v3, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2800_v3).Cardinality()))).Uint32(), p1))
		_ = _rhs433
		var _rhs434 _dafny.Sequence = ((func() D9 {
			if (_2797_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-8), _dafny.IntOfUint32((_2797_v0).Cardinality()))).Uint32()).(bool) {
				return _2811_v14
			}
			return _2811_v14
		})()).Dtor_cf25()
		_ = _rhs434
		var _rhs435 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lvpib")).Cardinality())
		_ = _rhs435
		var _rhs436 _dafny.CodePoint = _2810_v13
		_ = _rhs436
		var _lhs271 _dafny.Array = _2809_v12
		_ = _lhs271
		var _lhs272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))
		_ = _lhs272
		_2799_v2 = _rhs432
		_2801_v4 = _rhs433
		_2802_v5 = _rhs434
		(_lhs271).ArraySet1(_rhs435, (_lhs272).Int())
		_2810_v13 = _rhs436
		if false {
			var _2812_v15 bool
			_ = _2812_v15
			_2812_v15 = false
			var _2813_v16 _dafny.MultiSet
			_ = _2813_v16
			_2813_v16 = _dafny.MultiSetOf(_dafny.IntOfUint32((_2797_v0).Cardinality()))
			_2812_v15 = (_2813_v16).IsSubsetOf((_2813_v16).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F7))))
			var _index444 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))
			_ = _index444
			(_2809_v12).ArraySet1((_dafny.Zero).Minus(p0), (_index444).Int())
			var _index445 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))
			_ = _index445
			(_2809_v12).ArraySet1(_this.F7, (_index445).Int())
			var _index446 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))
			_ = _index446
			(_2809_v12).ArraySet1(_2805_v8.F14, (_index446).Int())
			var _2814_v17 _dafny.Array
			_ = _2814_v17
			var _nw429 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(17))
			_ = _nw429
			_2814_v17 = _nw429
			var _2815_v18 _dafny.Map
			_ = _2815_v18
			_2815_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _2805_v8.F14)
			var _2816_v19 _dafny.Set
			_ = _2816_v19
			_2816_v19 = _dafny.SetOf(_2815_v18)
			var _index447 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_2814_v17), 0))
			_ = _index447
			(_2814_v17).ArraySet1(_2816_v19, (_index447).Int())
			var _index448 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_2814_v17), 0))
			_ = _index448
			var _rhs437 _dafny.Set = (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2812_v15, _2805_v8.F14), _2815_v18)).Union(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), p0), _2815_v18, _2815_v18, _2815_v18))
			_ = _rhs437
			var _rhs438 _dafny.Int = (((_this).F6()).Minus((_dafny.Zero).Minus(_this.F7))).Minus(p0)
			_ = _rhs438
			var _rhs439 bool = (_2805_v8).F13()
			_ = _rhs439
			var _lhs273 _dafny.Array = _2814_v17
			_ = _lhs273
			var _lhs274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_2814_v17), 0))
			_ = _lhs274
			var _lhs275 *C5 = _2805_v8
			_ = _lhs275
			(_lhs273).ArraySet1(_rhs437, (_lhs274).Int())
			_lhs275.F14 = _rhs438
			_2812_v15 = _rhs439
		} else {
			var _2817_v20 bool
			_ = _2817_v20
			_2817_v20 = true
			var _2818_v21 _dafny.Array
			_ = _2818_v21
			var _nw430 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw430
			_2818_v21 = _nw430
			var _2819_v22 _dafny.Map
			_ = _2819_v22
			_2819_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2818_v21, true)
			_2817_v20 = (func() bool {
				if (_2819_v22).Contains(_2818_v21) {
					return (_2819_v22).Get(_2818_v21).(bool)
				}
				return (_2805_v8).F13()
			})()
			_2817_v20 = _2817_v20
			var _2820_v23 *C5
			_ = _2820_v23
			var _nw431 *C5 = New_C5_()
			_ = _nw431
			_nw431.Ctor__(((_2805_v8).F13()) || (_2817_v20), _this.F7, _this.F5(), (_this).F6(), ((_2805_v8).F13()) && ((_2805_v8).F13()))
			_2820_v23 = _nw431
			_2817_v20 = !((_2820_v23).F13())
			var _2821_v24 _dafny.MultiSet
			_ = _2821_v24
			_2821_v24 = _dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(685)), _2800_v3, _2800_v3, _2800_v3, Companion_Default___.Fm39(_2802_v5, _2799_v2, (_2805_v8).F13(), globalState))
			var _2822_v25 _dafny.MultiSet
			_ = _2822_v25
			_2822_v25 = _dafny.MultiSetOf(_2800_v3)
			_2821_v24 = (_2822_v25)
		}
		var _index449 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))
		_ = _index449
		var _rhs440 _dafny.Int = (_2809_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))).Int()).(_dafny.Int)
		_ = _rhs440
		var _rhs441 _dafny.Array = _2809_v12
		_ = _rhs441
		var _lhs276 _dafny.Array = _2809_v12
		_ = _lhs276
		var _lhs277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))
		_ = _lhs277
		(_lhs276).ArraySet1(_rhs440, (_lhs277).Int())
		_2809_v12 = _rhs441
		var _2823_v26 _dafny.Map
		_ = _2823_v26
		_2823_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F4())
		var _2824_v27 _dafny.MultiSet
		_ = _2824_v27
		_2824_v27 = _dafny.MultiSetOf(_2809_v12)
		var _2825_v28 _dafny.Array
		_ = _2825_v28
		var _len0_82 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_82
		var _nw432 _dafny.Array
		_ = _nw432
		if _len0_82.Cmp(_dafny.Zero) == 0 {
			_nw432 = _dafny.NewArray(_len0_82)
		} else {
			var _init82 func(_dafny.Int) bool = (func(_2826_v8 *C5) func(_dafny.Int) bool {
				return func(_2827_i0 _dafny.Int) bool {
					return (_2826_v8).F13()
				}
			})(_2805_v8)
			_ = _init82
			var _element0_82 = _init82(_dafny.Zero)
			_ = _element0_82
			_nw432 = _dafny.NewArrayFromExample(_element0_82, nil, _len0_82)
			(_nw432).ArraySet1(_element0_82, 0)
			var _nativeLen0_82 = (_len0_82).Int()
			_ = _nativeLen0_82
			for _i0_82 := 1; _i0_82 < _nativeLen0_82; _i0_82++ {
				(_nw432).ArraySet1(_init82(_dafny.IntOf(_i0_82)), _i0_82)
			}
		}
		_2825_v28 = _nw432
		var _2828_v29 _dafny.Map
		_ = _2828_v29
		_2828_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2802_v5, (Companion_Default___.SafeIndex((_2824_v27).Cardinality(), _dafny.IntOfUint32((_2802_v5).Cardinality()))).Uint32(), p2)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_2825_v28)).Cardinality()))
		var _2829_v30 _dafny.Map
		_ = _2829_v30
		_2829_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_2828_v29).Cardinality())
		var _2830_v31 _dafny.Array
		_ = _2830_v31
		var _nwElement0_89 bool = (_2805_v8).F13()
		_ = _nwElement0_89
		var _nw433 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_89, nil, _dafny.IntOfInt64(16))
		_ = _nw433
		(_nw433).ArraySet1(_nwElement0_89, 0)
		(_nw433).ArraySet1((_2805_v8).F13(), 1)
		(_nw433).ArraySet1((_2805_v8).Fm8(p1, (_2805_v8).F13(), _2802_v5, _2800_v3, globalState), 2)
		(_nw433).ArraySet1(!((_this).F4()), 3)
		(_nw433).ArraySet1((_this).F4(), 4)
		(_nw433).ArraySet1(((_2805_v8).F13()) == ((_this).F4()), 5)
		(_nw433).ArraySet1(!(Companion_Default___.Fm0(p2, globalState)), 6)
		(_nw433).ArraySet1((_this).F4(), 7)
		(_nw433).ArraySet1((_2805_v8).F13(), 8)
		(_nw433).ArraySet1(((_this).F4()) || ((Companion_D4_.Create_DC13_((_2809_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))).Int()).(_dafny.Int), (_2805_v8).F13(), p0, _dafny.Companion_Sequence_.Update(_2797_v0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2797_v0).Cardinality()), _dafny.IntOfUint32((_2797_v0).Cardinality()))).Uint32(), (_2805_v8).F13()))).Dtor_cf15()), 9)
		(_nw433).ArraySet1((_2805_v8).F13(), 10)
		(_nw433).ArraySet1(!(_2808_v11).Equals(_dafny.MultiSetOf((func() bool {
			if (_2823_v26).Contains((_2829_v30).Cardinality()) {
				return (_2823_v26).Get((_2829_v30).Cardinality()).(bool)
			}
			return (_2805_v8).F13()
		})(), false)), 11)
		(_nw433).ArraySet1(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-420))).Uint32(), func(coer126 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg126 _dafny.Int) interface{} {
				return coer126(arg126)
			}
		}((func(_2831_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2832_i1 _dafny.Int) _dafny.CodePoint {
				return _2831_v13
			}
		})(_2810_v13))), _2802_v5)), 12)
		(_nw433).ArraySet1((_2805_v8).F13(), 13)
		(_nw433).ArraySet1((_this).F4(), 14)
		(_nw433).ArraySet1((_this).F4(), 15)
		_2830_v31 = _nw433
		_2825_v28 = _2825_v28
		var _2833_v32 bool
		_ = _2833_v32
		_2833_v32 = true
		_2833_v32 = false
		if (_this).F4() {
			(_2805_v8).F14 = Companion_Default___.SafeDivisionInt(_2805_v8.F14, (_dafny.Zero).Minus((_2809_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))).Int()).(_dafny.Int)))
			_2833_v32 = ((_2805_v8).F13()) || ((_dafny.MultiSetOf(_dafny.IntOfUint32((_2797_v0).Cardinality()), _this.F7)).IsDisjointFrom(_dafny.MultiSetOf(p0)))
			var _index450 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_2825_v28), 0))
			_ = _index450
			(_2825_v28).ArraySet1((_2805_v8).F13(), (_index450).Int())
			var _index451 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_2825_v28), 0))
			_ = _index451
			(_2825_v28).ArraySet1((_2805_v8).F13(), (_index451).Int())
			var _2834_v33 *C0
			_ = _2834_v33
			var _nw434 *C0 = New_C0_()
			_ = _nw434
			_nw434.Ctor__(_2805_v8.F14)
			_2834_v33 = _nw434
			var _index452 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))
			_ = _index452
			(_2809_v12).ArraySet1((_2834_v33).F12(), (_index452).Int())
		} else {
			(_2805_v8).F14 = (_2809_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))).Int()).(_dafny.Int)
			_2803_v6 = (_2803_v6).Update((_this).F4(), (_2805_v8).F13())
			(_2805_v8).F14 = (_2805_v8.F14).Times(_dafny.IntOfInt64(248))
			var _index453 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_2809_v12), 0))
			_ = _index453
			(_2809_v12).ArraySet1(p0, (_index453).Int())
			(_2805_v8).F14 = (_this).F6()
		}
		r0 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, (_this).F6())).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2805_v8.F14, _dafny.IntOfUint32((_2802_v5).Cardinality()))).Cardinality(), p0)).Merge(_2828_v29)
		return r0
	}
}
func (_this *C9) F8() _dafny.Set {
	{
		return _this._f8
	}
}

// End of class C9
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
